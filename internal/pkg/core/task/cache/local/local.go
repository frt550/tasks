package local

import (
	"context"
	"sync"
	cachePkg "tasks/internal/pkg/core/task/cache"
	"tasks/internal/pkg/core/task/models"
	"time"

	"github.com/pkg/errors"
)

const poolSize = 10
const shortDuration = 10 * time.Millisecond

var (
	ErrTaskNotExists = errors.New("Task does not exist")
	ErrTaskExists    = errors.New("Task already exists")
)

func New() cachePkg.Interface {
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

func (c *cache) List() []models.Task {
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
	return result
}

func (c *cache) Add(task models.Task) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
		cancel()
	}()

	if _, ok := c.data[task.Id]; ok {
		return errors.Wrapf(ErrTaskExists, "task-id: [%d]", task.Id)
	}
	c.data[task.Id] = task
	return nil
}

func (c *cache) Update(task models.Task) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
		cancel()
	}()

	if _, ok := c.data[task.Id]; !ok {
		return errors.Wrapf(ErrTaskNotExists, "task-id: [%d]", task.Id)
	}
	c.data[task.Id] = task
	return nil
}

func (c *cache) Delete(id uint) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
		cancel()
	}()

	if _, ok := c.data[id]; !ok {
		return errors.Wrapf(ErrTaskNotExists, "task-id: [%d]", id)
	}
	delete(c.data, id)
	return nil
}

func (c *cache) Get(id uint) (models.Task, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	_, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
		cancel()
	}()

	if _, ok := c.data[id]; !ok {
		return models.Task{}, errors.Wrapf(ErrTaskNotExists, "task-id: [%d]", id)
	}
	return c.data[id], nil
}
