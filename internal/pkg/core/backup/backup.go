package backup

import (
	"context"
	"encoding/json"
	"tasks/internal/pkg/core/backup/models"
	"tasks/internal/pkg/core/backup/repository"
	transportPkg "tasks/internal/pkg/core/backup/transport"
	producerPkg "tasks/internal/pkg/core/producer"
	timePkg "tasks/internal/pkg/core/time"
	"tasks/pkg/contract/kafka"

	"github.com/google/uuid"

	"github.com/patrickmn/go-cache"
)

type Interface interface {
	Backup(ctx context.Context) (*models.Backup, error)
	AsyncBackup(ctx context.Context, requestId string) (*models.AsyncBackup, error)
}

func New(repository repository.Interface, cache *cache.Cache, producer producerPkg.Interface) Interface {
	return &core{
		repository: repository,
		cache:      cache,
		producer:   producer,
	}
}

type core struct {
	repository repository.Interface
	cache      *cache.Cache
	producer   producerPkg.Interface
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
		CreatedAt: timePkg.NowUTCFormatted(),
	}
	if err := b.repository.Insert(ctx, backup); err != nil {
		return nil, err
	}
	return backup, nil
}

// AsyncBackup
// task:
// consume message from task-all-request topic
// publish to task-all-response topic
//
// backup:
// publish message to topic task-all-request
// consume message from task-all-response topic
func (b *core) AsyncBackup(ctx context.Context, requestId string) (*models.AsyncBackup, error) {
	// for new request, generate request id and publish it
	if requestId == "" {
		newRequestId := uuid.New().String()
		b.cache.Set(newRequestId, nil, cache.NoExpiration)

		// publish message to topic task-all-request
		err := b.producer.Publish(kafka.TopicTaskAllRequest, kafka.TaskAllRequestMessage{RequestId: newRequestId})
		if err != nil {
			return nil, err
		}
		return &models.AsyncBackup{
			RequestId: newRequestId,
			State:     "registered",
			Backup:    nil,
		}, nil
	}

	jsonData, found := b.cache.Get(requestId)
	// request id is not found
	if found == false {
		return &models.AsyncBackup{
			RequestId: requestId,
			State:     "unknown/expired request id",
			Backup:    nil,
		}, nil
	}

	// message is not processed yet
	if jsonData == nil {
		return &models.AsyncBackup{
			RequestId: requestId,
			State:     "waiting",
			Backup:    nil,
		}, nil
	}

	// message is processed
	backup := &models.Backup{
		Data:      jsonData.(string),
		CreatedAt: timePkg.NowUTCFormatted(),
	}
	if err := b.repository.Insert(ctx, backup); err != nil {
		return nil, err
	}

	return &models.AsyncBackup{
		RequestId: requestId,
		State:     "completed",
		Backup:    backup,
	}, nil
}
