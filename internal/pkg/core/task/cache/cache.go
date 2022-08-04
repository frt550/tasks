package cache

import "tasks/internal/pkg/core/task/models"

type Interface interface {
	List() []models.Task
	Add(task models.Task) error
	Update(task models.Task) error
	Delete(id uint) error
	Get(id uint) (models.Task, error)
}
