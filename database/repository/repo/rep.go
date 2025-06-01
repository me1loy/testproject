package repo

import (
	"context"

	"server/database/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Repository struct {
	Client Client
}

type Repository1 interface {
	CreateUser(ctx context.Context, u models.User) (models.User, error)
	GetProjects(ctx context.Context) ([]models.Project, error)
}
