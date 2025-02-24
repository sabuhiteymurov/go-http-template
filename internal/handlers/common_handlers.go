package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go-net_http-project/db/services"
	"go-net_http-project/internal/helpers"
)

type CommonHandlers struct {
	DB        *pgxpool.Pool
	DbService *services.DatabaseService
	Validator *helpers.CustomValidator
}

func NewCommonHandlers(db *pgxpool.Pool) *CommonHandlers {
	return &CommonHandlers{
		DB:        db,
		DbService: services.NewDatabaseService(db),
	}
}
