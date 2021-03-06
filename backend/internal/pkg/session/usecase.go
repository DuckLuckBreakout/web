package session

import "github.com/DuckLuckBreakout/web/backend/internal/pkg/models"

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock github.com/DuckLuckBreakout/web/backend/internal/pkg/session UseCase

type UseCase interface {
	GetUserIdBySession(sessionCookieValue string) (uint64, error)
	CreateNewSession(userId uint64) (*models.Session, error)
	DestroySession(sessionCookieValue string) error
}
