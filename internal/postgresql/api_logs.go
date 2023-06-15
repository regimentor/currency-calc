package postgresql

import (
	"context"
	"fmt"
	"log"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	"github.com/regimentor/currency-calc/internal"
)

type ApiLogStorage struct {
	connection *pgxpool.Pool
}

func NewApiLogStorage(connection *pgxpool.Pool) *ApiLogStorage {
	return &ApiLogStorage{connection: connection}
}

func (a *ApiLogStorage) Create(ctx context.Context, apiLog *internal.CreateApiLogsDto) error {
	log.Printf("ApiLogStorage.Create apiLog: %v", apiLog)

	query := `
		insert into api_logs (user_id, request_time, request_type)
		values ($1, $2, $3) returning id;	
	`

	year, month, day := apiLog.RequestTime.Date()
	dateStr := fmt.Sprintf("%d-%d-%d", year, month, day)

	row := a.connection.QueryRow(ctx, query,
		apiLog.UserId, dateStr, apiLog.RequestType)

	var id int64
	err := row.Scan(id)

	if err != nil {
		return fmt.Errorf("create api log due err: %v", err)
	}

	return nil
}
