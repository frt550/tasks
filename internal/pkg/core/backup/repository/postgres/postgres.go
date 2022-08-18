package postgres

import (
	"context"
	"fmt"
	"tasks/internal/pkg/core/backup/models"
	"tasks/internal/pkg/core/pool"

	"github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v4/pgxpool"

	repositoryPkg "tasks/internal/pkg/core/backup/repository"
)

type repository struct {
	pool *pgxpool.Pool
}

func New() repositoryPkg.Interface {
	return &repository{pool.GetInstance()}
}

func (r *repository) Insert(ctx context.Context, backup *models.Backup) error {
	sql, args, err := squirrel.
		Insert("backup").
		Columns("data, created_at").
		Values(backup.Data, backup.CreatedAt).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("Repository.Insert: to sql: %w", err)
	}

	row := r.pool.QueryRow(ctx, sql, args...)
	if err := row.Scan(&backup.Id); err != nil {
		return fmt.Errorf("Repository.Insert: exec: %w", err)
	}
	return nil
}
