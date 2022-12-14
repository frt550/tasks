//go:generate mockery --filename repository_mock.go --inpackage --name Interface

package repository

import (
	"context"
	"tasks/internal/pkg/core/task/models"
)

type Interface interface {
	FindAll(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Insert(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	DeleteById(ctx context.Context, id uint64) error
	FindOneById(ctx context.Context, id uint64) (*models.Task, error)
}
