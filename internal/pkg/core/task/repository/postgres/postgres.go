package postgres

import (
	"context"
	"fmt"
	errPkg "tasks/internal/pkg/core/error"
	"tasks/internal/pkg/core/pool"
	"tasks/internal/pkg/core/task/models"
	storagePkg "tasks/internal/pkg/core/task/repository"

	"github.com/jackc/pgx/v4"

	"github.com/pkg/errors"

	"github.com/georgysavva/scany/pgxscan"

	"github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	pool *pgxpool.Pool
}

func New() storagePkg.Interface {
	return &repository{pool.GetInstance()}
}

func (r *repository) FindAll(ctx context.Context, limit, offset uint64) ([]models.Task, error) {
	selectBuilder := squirrel.
		Select("id, title, is_completed, created_at, completed_at").
		From("task").
		PlaceholderFormat(squirrel.Dollar).OrderBy()
	if limit > 0 {
		selectBuilder = selectBuilder.Limit(limit)
	}
	if offset > 0 {
		selectBuilder = selectBuilder.Offset(offset).OrderBy("created_at ASC")
	}
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return make([]models.Task, 0), fmt.Errorf("Repository.FindAll: to sql: %w", err)
	}

	var result []models.Task
	err = pgxscan.Select(ctx, r.pool, &result, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("Repository.FindAll: select: %w", err)
	}
	return result, nil
}

func (r *repository) Insert(ctx context.Context, task *models.Task) error {
	sql, args, err := squirrel.
		Insert("task").
		Columns("title, is_completed, created_at, completed_at").
		Values(task.Title, task.IsCompleted, task.CreatedAt, task.CompletedAt).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("Repository.Insert: to sql: %w", err)
	}

	row := r.pool.QueryRow(ctx, sql, args...)
	if err := row.Scan(&task.Id); err != nil {
		return fmt.Errorf("Repository.Insert: exec: %w", err)
	}
	return nil
}

func (r *repository) Update(ctx context.Context, task *models.Task) error {
	sql, args, err := squirrel.
		Update("task").
		SetMap(squirrel.Eq{
			"title":        task.Title,
			"is_completed": task.IsCompleted,
			"created_at":   task.CreatedAt,
			"completed_at": task.CompletedAt,
		}).
		Where(squirrel.Eq{"id": task.Id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("Repository.UpdateTitle: to sql: %w", err)
	}

	if ct, err := r.pool.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("Repository.UpdateTitle: exec: %w", err)
	} else {
		if ct.RowsAffected() == 0 {
			return errors.Wrapf(errPkg.DomainError, "Sorry, task #%d is not found", task.Id)
		} else {
			return nil
		}
	}
}

func (r *repository) DeleteById(ctx context.Context, id uint) error {
	sql, args, err := squirrel.
		Delete("task").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("Repository.DeleteById: to sql: %w", err)
	}

	_, err = r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Repository.DeleteById: exec: %w", err)
	}
	return nil
}

func (r *repository) FindOneById(ctx context.Context, id uint) (models.Task, error) {
	sql, args, err := squirrel.
		Select("id, title, is_completed, created_at, completed_at").
		From("task").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Task{}, fmt.Errorf("Repository.FindOneById: to sql: %w", err)
	}

	var result models.Task
	if err := pgxscan.Get(ctx, r.pool, &result, sql, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Task{}, errors.Wrapf(errPkg.DomainError, "Sorry, task #%d is not found", id)
		} else {
			return models.Task{}, fmt.Errorf("Repository.FindOneById: select: %w", err)
		}
	}

	return result, nil
}
