package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DuckLuckBreakout/ozonBackend/internal/pkg/notification"
	"github.com/DuckLuckBreakout/ozonBackend/internal/server/errors"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	conn *redis.Client
}

type DtoUserId struct {
	Id uint64 `json:"id"`
}

type DtoNotificationKeys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type DtoSubscribes struct {
	Credentials map[string]*DtoNotificationKeys `json:"subscribes" valid:"notnull"`
}

func NewSessionRedisRepository(conn *redis.Client) notification.Repository {
	return &RedisRepository{
		conn: conn,
	}
}

func (r *RedisRepository) getNewKey(value uint64) string {
	return fmt.Sprintf("notification:%d", value)
}

func (r *RedisRepository) AddSubscribeUser(userId *DtoUserId, subscribes *DtoSubscribes) error {
	key := r.getNewKey(userId.Id)

	data, err := json.Marshal(subscribes)
	if err != nil {
		return errors.ErrCanNotMarshal
	}

	err = r.conn.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		return errors.ErrDBInternalError
	}

	return nil
}

func (r *RedisRepository) SelectCredentialsByUserId(userId *DtoUserId) (*DtoSubscribes, error) {
	subscribes := &DtoSubscribes{}
	key := r.getNewKey(userId.Id)

	data, err := r.conn.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, errors.ErrSessionNotFound
	}

	if err = json.Unmarshal(data, subscribes); err != nil {
		return nil, errors.ErrCanNotUnmarshal
	}

	return subscribes, nil
}

func (r *RedisRepository) DeleteSubscribeUser(userId *DtoUserId) error {
	key := r.getNewKey(userId.Id)

	err := r.conn.Del(context.Background(), key).Err()
	if err != nil {
		return errors.ErrDBInternalError
	}

	return nil
}
