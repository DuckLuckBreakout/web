package notification

import "github.com/DuckLuckBreakout/web/backend/internal/pkg/models"

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock github.com/DuckLuckBreakout/web/backend/internal/pkg/notification UseCase

type UseCase interface {
	SubscribeUser(userId uint64, credentials *models.NotificationCredentials) error
	UnsubscribeUser(userId uint64, endpoint string) error
}
