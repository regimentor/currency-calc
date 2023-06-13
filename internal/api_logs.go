package internal

import "time"

const (
	GET_ALL     = 1
	GET_BY_BASE = 2
)

type ApiLogs struct {
	ID          UserId `json:"id"`
	RequestType int    `json:"request_type"`
	RequestTime string `json:"request_date"`
	UserId      int64  `json:"user_id"`
}

type CreateApiLogsDto struct {
	RequestType int       `json:"request_type"`
	RequestTime time.Time `json:"request_time"`
	UserId      UserId    `json:"user_id"`
}

type ApiLogsStorage interface {
	Create(apiLogs *CreateApiLogsDto) error
}

type ApiLogsRepository struct {
	storage ApiLogsStorage
}

func NewApiLogsRepository(storage ApiLogsStorage) *ApiLogsRepository {
	return &ApiLogsRepository{storage: storage}
}

func (r *ApiLogsRepository) Create(apiLogs *CreateApiLogsDto) error {
	return r.storage.Create(apiLogs)
}
