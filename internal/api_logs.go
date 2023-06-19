package internal

import (
	"context"
	"github.com/regimentor/currency-calc/internal/models"
	"time"
)

const (
	GetAll    = 1
	GetByBase = 2
)

type ApiLogs struct {
	ID          models.UserId `json:"id"`
	RequestType int           `json:"request_type"`
	RequestTime string        `json:"request_date"`
	UserId      int64         `json:"user_id"`
}

type CreateApiLogsDto struct {
	RequestType int           `json:"request_type"`
	RequestTime time.Time     `json:"request_time"`
	UserId      models.UserId `json:"user_id"`
}

type ApiLogsStorage interface {
	Create(ctx context.Context, apiLogs *CreateApiLogsDto) error
}

type ApiLogsRepository struct {
	storage ApiLogsStorage
}

func NewApiLogsRepository(storage ApiLogsStorage) *ApiLogsRepository {
	return &ApiLogsRepository{storage: storage}
}

func (r *ApiLogsRepository) Create(ctx context.Context, apiLogs *CreateApiLogsDto) error {
	return r.storage.Create(ctx, apiLogs)
}
