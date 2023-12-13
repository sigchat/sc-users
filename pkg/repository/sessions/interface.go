package sessions

import (
	"context"
	"github.com/sigchat/sc-users/pkg/domain/model"
)

type Repository interface {
	CreateSession(ctx context.Context, userID int) (*model.Session, error)
	GetSessions(ctx context.Context) []model.Session
	GetSessionByUserID(ctx context.Context, userID int) (*model.Session, error)
	DeleteSessionByUserID(ctx context.Context, userID int) error
}
