package task

import (
	"context"
	"errors"
	"fmt"
	errPkg "tasks/internal/pkg/core/error"
	"tasks/internal/pkg/core/task/models"
	"tasks/internal/pkg/core/task/repository"
	timePkg "tasks/internal/pkg/core/time"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("must create new instance", func(t *testing.T) {
		mockRepository := &repository.MockInterface{}
		core := New(mockRepository)
		require.NotNil(t, core)
	})
}

func TestCreate(t *testing.T) {
	t.Run("must trim title", func(t *testing.T) {
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("Insert", mock.Anything, mock.Anything).
			Return(nil).
			Once()
		core := &core{
			repository: mockRepository,
		}
		expected := "New task"
		title := "\t\n\v\f\r " + expected + "\t\n\v\f\r "

		// Act
		task, err := core.Create(context.Background(), title)

		// Assert
		require.Nil(t, err)
		require.Equal(t, expected, task.Title)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error if title is empty", func(t *testing.T) {
		mockRepository := &repository.MockInterface{}
		core := &core{
			repository: mockRepository,
		}
		expected := "Title cannot be empty: "
		title := "\t\n\v\f\r "

		// Act
		task, err := core.Create(context.Background(), title)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.ErrorIs(t, err, errPkg.DomainError)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error from repository.Insert()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("Insert", mock.Anything, mock.Anything).
			Return(errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Create(context.Background(), "New task")

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return created task", func(t *testing.T) {
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("Insert", mock.Anything, mock.Anything).
			Return(nil).
			Once()
		core := &core{
			repository: mockRepository,
		}
		title := "New task"
		expected := &models.Task{
			Id:          0,
			Title:       title,
			IsCompleted: false,
			CreatedAt:   timePkg.NowUTCFormatted(),
			CompletedAt: "",
		}

		// Act
		task, err := core.Create(context.Background(), title)

		// Assert
		require.Nil(t, err)
		require.Equal(t, expected, task)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("must trim title", func(t *testing.T) {
		mockTask := &models.Task{
			Title: "New task",
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(mockTask, nil).
			Once()
		mockRepository.
			On("Update", mock.Anything, mock.Anything).
			Return(nil).
			Once()
		core := &core{
			repository: mockRepository,
		}
		expected := "Updated task"
		title := "\t\n\v\f\r " + expected + "\t\n\v\f\r "

		// Act
		task, err := core.UpdateTitle(context.Background(), mockTask.Id, title)

		// Assert
		require.Nil(t, err)
		require.Equal(t, expected, task.Title)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error if title is empty", func(t *testing.T) {
		mockRepository := &repository.MockInterface{}
		core := &core{
			repository: mockRepository,
		}
		expected := "Title cannot be empty: "
		title := "\t\n\v\f\r "

		// Act
		task, err := core.UpdateTitle(context.Background(), 1, title)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.ErrorIs(t, err, errPkg.DomainError)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error from repository.FindOneById()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(nil, errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.UpdateTitle(context.Background(), 1, "New task")

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error if task is completed", func(t *testing.T) {
		mockTask := &models.Task{
			Id:          1,
			Title:       "New task",
			IsCompleted: true,
			CreatedAt:   timePkg.NowUTCFormatted(),
			CompletedAt: timePkg.NowUTCFormatted(),
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(mockTask, nil).
			Once()
		core := &core{
			repository: mockRepository,
		}
		expected := "Completed task cannot be updated: "

		// Act
		task, err := core.UpdateTitle(context.Background(), 1, "New task")

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.ErrorIs(t, err, errPkg.DomainError)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error from repository.Update()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(&models.Task{}, nil).
			Once()
		mockRepository.
			On("Update", mock.Anything, mock.Anything).
			Return(errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.UpdateTitle(context.Background(), 1, "New task")

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return updated task", func(t *testing.T) {
		mockTask := &models.Task{
			Id:    1,
			Title: "New task",
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(mockTask, nil).
			Once()
		mockRepository.
			On("Update", mock.Anything, mock.Anything).
			Return(nil).
			Once()
		core := &core{
			repository: mockRepository,
		}
		expected := "Updated task"

		// Act
		task, err := core.UpdateTitle(context.Background(), mockTask.Id, expected)

		// Assert
		require.Nil(t, err)
		require.Equal(t, expected, task.Title)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})
}

func TestDelete(t *testing.T) {
	t.Run("must return error from repository.FindOneById()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(nil, errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Delete(context.Background(), 1)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error from repository.DeleteById()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(&models.Task{}, nil).
			Once()
		mockRepository.
			On("DeleteById", mock.Anything, mock.Anything).
			Return(errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Delete(context.Background(), 1)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return deleted task", func(t *testing.T) {
		mockTask := &models.Task{
			Id:    1,
			Title: "New task",
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(mockTask, nil).
			Once()
		mockRepository.
			On("DeleteById", mock.Anything, mock.Anything).
			Return(nil).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Delete(context.Background(), mockTask.Id)

		// Assert
		require.Nil(t, err)
		require.Equal(t, mockTask, task)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})
}

func TestComplete(t *testing.T) {
	t.Run("must return error from repository.FindOneById()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(nil, errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Complete(context.Background(), 1)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error if task is already completed", func(t *testing.T) {
		mockTask := &models.Task{
			Id:          1,
			Title:       "New task",
			IsCompleted: true,
			CompletedAt: timePkg.NowUTCFormatted(),
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(mockTask, nil).
			Once()
		core := &core{
			repository: mockRepository,
		}
		expected := fmt.Sprintf("Task %v is already completed: ", mockTask.String())

		// Act
		task, err := core.Complete(context.Background(), mockTask.Id)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.ErrorIs(t, err, errPkg.DomainError)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return error from repository.Update()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(&models.Task{}, nil).
			Once()
		mockRepository.
			On("Update", mock.Anything, mock.Anything).
			Return(errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Complete(context.Background(), 1)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return completed task", func(t *testing.T) {
		mockTask := &models.Task{
			Id:          1,
			Title:       "New task",
			IsCompleted: false,
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(mockTask, nil).
			Once()
		mockRepository.
			On("Update", mock.Anything, mock.Anything).
			Return(nil).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Complete(context.Background(), mockTask.Id)

		// Assert
		require.Nil(t, err)
		require.True(t, task.IsCompleted)
		require.Equal(t, timePkg.NowUTCFormatted(), task.CompletedAt)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})
}

func TestAll(t *testing.T) {
	t.Run("must return error from repository.FindAll()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindAll", mock.Anything, mock.Anything, mock.Anything).
			Return([]*models.Task{}, errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		tasks, err := core.All(context.Background(), 0, 0)

		// Assert
		require.Equal(t, []*models.Task{}, tasks)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return all tasks", func(t *testing.T) {
		expected := []*models.Task{
			{Id: 1},
			{Id: 2},
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindAll", mock.Anything, mock.Anything, mock.Anything).
			Return(expected, nil).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		tasks, err := core.All(context.Background(), 0, 0)

		// Assert
		require.Equal(t, expected, tasks)
		require.Nil(t, err)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must pass limit and offset to repository", func(t *testing.T) {
		var limit uint64 = 100
		var offset uint64 = 1000
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindAll", mock.Anything, limit, offset).
			Return([]*models.Task{}, nil).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		tasks, err := core.All(context.Background(), limit, offset)

		// Assert
		require.Equal(t, []*models.Task{}, tasks)
		require.Nil(t, err)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})
}

func TestGet(t *testing.T) {
	t.Run("must return error from repository.FindOneById()", func(t *testing.T) {
		expected := "error from repository"
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(nil, errors.New(expected)).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Get(context.Background(), 1)

		// Assert
		require.Nil(t, task)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
		mock.AssertExpectationsForObjects(t, mockRepository)
	})

	t.Run("must return task", func(t *testing.T) {
		mockTask := &models.Task{
			Id:    1,
			Title: "New task",
		}
		mockRepository := &repository.MockInterface{}
		mockRepository.
			On("FindOneById", mock.Anything, mock.Anything).
			Return(mockTask, nil).
			Once()
		core := &core{
			repository: mockRepository,
		}

		// Act
		task, err := core.Get(context.Background(), mockTask.Id)

		// Assert
		require.Nil(t, err)
		require.Equal(t, mockTask, task)
		mock.AssertExpectationsForObjects(t, mockRepository)
	})
}
