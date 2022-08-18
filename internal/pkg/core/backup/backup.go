package backup

import (
	"context"
	"encoding/json"
	"tasks/internal/pkg/core/backup/models"
	"tasks/internal/pkg/core/backup/repository"
	"tasks/internal/pkg/core/backup/repository/postgres"
	transportPkg "tasks/internal/pkg/core/backup/transport"
	"time"
)

type Interface interface {
	Backup(ctx context.Context) (*models.Backup, error)
}

func New() Interface {
	return &core{
		repository: postgres.New(),
	}
}

type core struct {
	repository repository.Interface
}

func (b *core) Backup(ctx context.Context) (*models.Backup, error) {
	transport, err := transportPkg.New()
	if err != nil {
		return nil, err
	}
	tasks, err := transport.All(ctx)
	if err != nil {
		return nil, err
	}
	j, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}
	backup := &models.Backup{
		Data:      string(j),
		CreatedAt: time.Now(),
	}
	if err := b.repository.Insert(ctx, backup); err != nil {
		return nil, err
	}
	return backup, nil
}
