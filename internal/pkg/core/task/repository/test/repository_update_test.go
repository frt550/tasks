//go:build integration

package test

import (
	"context"
	"tasks/internal/pkg/core/task/models"
	timePkg "tasks/internal/pkg/core/time"
)

func (s *RepositorySuite) TestUpdateMustUpdateTask() {
	task := &models.Task{
		Title:     "New task 1",
		CreatedAt: timePkg.NowUTCFormatted(),
	}
	err1 := s.repository.Insert(context.Background(), task)

	// Act
	task.Title = "Updated"
	err2 := s.repository.Update(context.Background(), task)
	task, err3 := s.repository.FindOneById(context.Background(), task.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal("Updated", task.Title)
}
