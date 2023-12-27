package sessions

import (
	"context"
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/sigchat/sc-users/pkg/repository/sessions"
	"time"
)

type Interactor struct {
	repository sessions.Repository
}

func NewInteractor(repository sessions.Repository) *Interactor {
	return &Interactor{repository: repository}
}

func (in *Interactor) GetOrCreateSession(ctx context.Context, userID int) (*model.Session, error) {
	if existingSession, err := in.repository.GetSessionByUserID(ctx, userID); err == nil {
		if !existingSession.ExpiresAt.Before(time.Now()) {
			return existingSession, nil
		}
	}

	return in.repository.CreateSession(ctx, userID)
}
