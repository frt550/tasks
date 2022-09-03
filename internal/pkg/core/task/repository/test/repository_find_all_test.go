//go:build integration

package test

import (
	"context"
	"tasks/internal/pkg/core/task/models"
	timePkg "tasks/internal/pkg/core/time"
)

func (s *RepositorySuite) TestFindAllMustReturnAllTasks() {
	err1 := s.repository.Insert(context.Background(), &models.Task{
		Title:     "New task 1",
		CreatedAt: timePkg.NowUTCFormatted(),
	})
	err2 := s.repository.Insert(context.Background(), &models.Task{
		Title:     "New task 2",
		CreatedAt: timePkg.NowUTCFormatted(),
	})

	// Act
	tasks, err3 := s.repository.FindAll(context.Background(), 0, 0)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal(2, len(tasks))
}

func (s *RepositorySuite) TestFindAllMustReturnAllTasksUsingLimit() {
	err1 := s.repository.Insert(context.Background(), &models.Task{
		Title:     "New task 1",
		CreatedAt: timePkg.NowUTCFormatted(),
	})
	err2 := s.repository.Insert(context.Background(), &models.Task{
		Title:     "New task 2",
		CreatedAt: timePkg.NowUTCFormatted(),
	})

	// Act
	tasks, err3 := s.repository.FindAll(context.Background(), 1, 0)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal(1, len(tasks))
}

func (s *RepositorySuite) TestFindAllMustReturnAllTasksUsingOffset() {
	err1 := s.repository.Insert(context.Background(), &models.Task{
		Title:     "New task 1",
		CreatedAt: timePkg.NowUTCFormatted(),
	})
	err2 := s.repository.Insert(context.Background(), &models.Task{
		Title:     "New task 2",
		CreatedAt: timePkg.NowUTCFormatted(),
	})
	tasks, err3 := s.repository.FindAll(context.Background(), 0, 0)

	// Act
	tasks, err4 := s.repository.FindAll(context.Background(), 0, uint64(len(tasks)-1))

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Nil(err4)
	s.Require().Equal(1, len(tasks))
}
