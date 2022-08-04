package task

import (
	cachePkg "tasks/internal/pkg/core/task/cache"
	localCachePkg "tasks/internal/pkg/core/task/cache/local"
	"tasks/internal/pkg/core/task/models"
)

type Interface interface {
	Create(task models.Task) error
	Update(task models.Task) error
	Delete(id uint) error
	List() []models.Task
	Get(id uint) (models.Task, error)
}

func New() Interface {
	return &core{
		cache: localCachePkg.New(),
	}
}

type core struct {
	cache cachePkg.Interface
}

func (c *core) Create(task models.Task) error {
	return c.cache.Add(task)
}

func (c *core) Update(task models.Task) error {
	return c.cache.Update(task)
}

func (c *core) Delete(id uint) error {
	return c.cache.Delete(id)
}

func (c *core) List() []models.Task {
	return c.cache.List()
}

func (c *core) Get(id uint) (models.Task, error) {
	return c.cache.Get(id)
}
