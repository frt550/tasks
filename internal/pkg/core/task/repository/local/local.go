package local

import (
	"context"
	"sync"
	"tasks/internal/pkg/core/counter"
	errPkg "tasks/internal/pkg/core/error"
	"tasks/internal/pkg/core/task/models"
	storagePkg "tasks/internal/pkg/core/task/repository"
	"time"

	"github.com/pkg/errors"
)

const poolSize = 10
const shortDuration = 10 * time.Millisecond

func New() storagePkg.Interface {
	return &cache{
		mu:     sync.RWMutex{},
		data:   map[uint]models.Task{},
		poolCh: make(chan struct{}, poolSize),
	}
}

type cache struct {
	mu     sync.RWMutex
	data   map[uint]models.Task
	poolCh chan struct{}
}

func (c *cache) FindAll(_ context.Context, _, _ uint64) ([]models.Task, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
	}()

	result := make([]models.Task, 0, len(c.data))
	for _, value := range c.data {
		result = append(result, value)
	}
	return result, nil
}

func (c *cache) Insert(_ context.Context, task *models.Task) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
		cancel()
	}()

	task.Id = counter.GetId()
	if _, ok := c.data[task.Id]; ok {
		return errors.Wrapf(errPkg.DomainError, "Sorry, task #%d is already exists", task.Id)
	}
	c.data[task.Id] = *task
	return nil
}

func (c *cache) Update(_ context.Context, task *models.Task) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
		cancel()
	}()

	if _, ok := c.data[task.Id]; !ok {
		return errors.Wrapf(errPkg.DomainError, "Sorry, task #%d is not found", task.Id)
	}
	c.data[task.Id] = *task
	return nil
}

func (c *cache) DeleteById(_ context.Context, id uint) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
		cancel()
	}()

	if _, ok := c.data[id]; !ok {
		return errors.Wrapf(errPkg.DomainError, "Sorry, task #%d is not found", id)
	}
	delete(c.data, id)
	return nil
}

func (c *cache) FindOneById(_ context.Context, id uint) (models.Task, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
		cancel()
	}()

	if _, ok := c.data[id]; !ok {
		return models.Task{}, errors.Wrapf(errPkg.DomainError, "Sorry, task #%d is not found", id)
	}
	return c.data[id], nil
}
