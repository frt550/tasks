package task

import (
	"fmt"
	"strings"
	errPkg "tasks/internal/pkg/core/error"
	"tasks/internal/pkg/core/task/models"
	"tasks/internal/pkg/core/task/repository"

	"github.com/pkg/errors"

	timePkg "tasks/internal/pkg/core/time"

	"golang.org/x/net/context"
)

type Interface interface {
	Create(ctx context.Context, title string) (*models.Task, error)
	UpdateTitle(ctx context.Context, id uint64, title string) (*models.Task, error)
	Delete(ctx context.Context, id uint64) (*models.Task, error)
	Complete(ctx context.Context, id uint64) (*models.Task, error)
	All(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Get(ctx context.Context, id uint64) (*models.Task, error)
}

func New(repository repository.Interface) Interface {
	return &core{
		repository: repository,
	}
}

type core struct {
	repository repository.Interface
}

func (c *core) Create(ctx context.Context, title string) (*models.Task, error) {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		return nil, errors.Wrap(errPkg.DomainError, "Title cannot be empty")
	}
	var task = &models.Task{
		Title:       title,
		IsCompleted: false,
		CreatedAt:   timePkg.NowUTCFormatted(),
		CompletedAt: "",
	}
	if err := c.repository.Insert(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (c *core) UpdateTitle(ctx context.Context, id uint64, title string) (*models.Task, error) {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		return nil, errors.Wrap(errPkg.DomainError, "Title cannot be empty")
	}

	task, err := c.repository.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	}

	if task.IsCompleted {
		return nil, errors.Wrap(errPkg.DomainError, "Completed task cannot be updated")
	}

	task.Title = title
	if err := c.repository.Update(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (c *core) Delete(ctx context.Context, id uint64) (*models.Task, error) {
	task, err := c.repository.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := c.repository.DeleteById(ctx, id); err != nil {
		return nil, err
	}
	return task, nil
}

func (c *core) Complete(ctx context.Context, id uint64) (*models.Task, error) {
	task, err := c.repository.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	}
	if task.IsCompleted {
		return nil, errors.Wrap(errPkg.DomainError, fmt.Sprintf("Task %v is already completed", task))
	}

	task.IsCompleted = true
	task.CompletedAt = timePkg.NowUTCFormatted()
	if err := c.repository.Update(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (c *core) All(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	return c.repository.FindAll(ctx, limit, offset)
}

func (c *core) Get(ctx context.Context, id uint64) (*models.Task, error) {
	if task, err := c.repository.FindOneById(ctx, id); err != nil {
		return nil, err
	} else {
		return task, nil
	}
}
