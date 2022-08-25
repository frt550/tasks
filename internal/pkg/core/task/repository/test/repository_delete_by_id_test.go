package test

import (
	"context"
	"tasks/internal/pkg/core/task/models"
	timePkg "tasks/internal/pkg/core/time"
)

func (s *RepositorySuite) TestDeleteMustDeleteTask() {
	task := &models.Task{
		Title:     "New task 1",
		CreatedAt: timePkg.NowUTCFormatted(),
	}
	err1 := s.repository.Insert(context.Background(), task)

	// Act
	err2 := s.repository.DeleteById(context.Background(), task.Id)
	task, err3 := s.repository.FindOneById(context.Background(), task.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().NotNil(err3)
	s.Require().Nil(task)
}
