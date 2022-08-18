package repository

import (
	"context"
	"tasks/internal/pkg/core/backup/models"
)

type Interface interface {
	Insert(ctx context.Context, task *models.Backup) error
}
