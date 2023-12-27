package sessions

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sigchat/sc-http/pkg/transport/errors"
	"github.com/sigchat/sc-http/pkg/transport/server/auth"
	_ "github.com/sigchat/sc-http/pkg/transport/server/auth"
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/sigchat/sc-users/pkg/items"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

var _sessionsIDCnt = 1

type RAMDBRepository struct {
	items      items.List[*model.Session]
	m          sync.RWMutex
	privateKey []byte
}

func NewRAMDBRepository(privateKey []byte) *RAMDBRepository {
	return &RAMDBRepository{
		items:      items.List[*model.Session]{},
		m:          sync.RWMutex{},
		privateKey: privateKey,
	}
}

func (db *RAMDBRepository) CreateSession(ctx context.Context, userID int) (*model.Session, error) {
	expiresAt := time.Now().UTC().Add(time.Hour)

	c := auth.Claims{
		RegisteredClaims: &jwt.RegisteredClaims{},
		Extra: auth.ExtraClaims{
			UserID: userID,
		},
	}
	c.ID = uuid.Must(uuid.NewRandom()).String()
	c.Subject = "auth"
	c.Audience = []string{"*"}
	c.Issuer = "auth-service"
	// If dont subtract will be get error - Token used before issued
	c.IssuedAt = jwt.NewNumericDate(time.Now().UTC())
	c.ExpiresAt = jwt.NewNumericDate(expiresAt)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)

	privateKeyBytes, err := base64.StdEncoding.DecodeString(string(db.privateKey))
	if err != nil {
		slog.Warn("WARN decode private key string:", err)
		return nil, fmt.Errorf("base 64 decode: %v", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("generate key: %v", err)
	}

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("sign string: %v", err)
	}

	currentID := _sessionsIDCnt
	newSession := &model.Session{
		ID:          currentID,
		AccessToken: signedToken,
		ExpiresAt:   expiresAt,
	}

	_sessionsIDCnt++

	db.m.Lock()
	defer db.m.Unlock()

	db.items = append(db.items, newSession)
	return newSession, nil
}

func (db *RAMDBRepository) GetSessions(ctx context.Context) []model.Session {
	db.m.RLock()
	defer db.m.RUnlock()

	sessions := make([]model.Session, 0, len(db.items))
	db.items.Each(func(item *model.Session, index int) bool {
		sessions = append(sessions, *item)
		return false
	})
	return sessions
}

func (db *RAMDBRepository) GetSessionByUserID(ctx context.Context, userID int) (*model.Session, error) {
	db.m.RLock()
	defer db.m.RUnlock()

	filtered := db.items.Filter(func(item *model.Session, index int) bool {
		return item.ID == userID
	})

	if len(filtered) == 0 {
		return nil, errors.NewHttpError().
			WithCode(http.StatusNotFound).
			WithMessage(fmt.Sprintf("user with id=%d not found", userID))
	}

	return &*filtered[0], nil
}

func (db *RAMDBRepository) DeleteSessionByUserID(ctx context.Context, userID int) error {
	db.m.Lock()
	defer db.m.Unlock()

	sessionIndex := -1
	db.items.Each(func(item *model.Session, index int) bool {
		if item.ID == userID {
			sessionIndex = index
			return true
		}
		return false
	})

	if sessionIndex == -1 {
		return errors.NewHttpError().
			WithCode(http.StatusNotFound).
			WithMessage(fmt.Sprintf("user with id=%d not found", userID))
	}
	db.items = append(db.items[:sessionIndex], db.items[sessionIndex+1:]...)

	return nil
}
