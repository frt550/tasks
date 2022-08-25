package test

import (
	"context"
	errPkg "tasks/internal/pkg/core/error"
	"tasks/internal/pkg/core/task/models"
	timePkg "tasks/internal/pkg/core/time"
)

func (s *RepositorySuite) TestFindOneByIdMustFindTask() {
	expected := &models.Task{
		Title:     "New task 1",
		CreatedAt: timePkg.NowUTCFormatted(),
	}
	err1 := s.repository.Insert(context.Background(), expected)

	// Act
	task, err2 := s.repository.FindOneById(context.Background(), expected.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Equal(expected.Id, task.Id)
}
func (s *RepositorySuite) TestFindOneByIdMustReturnErrorIfTaskIsNotFound() {
	// Act
	task, err := s.repository.FindOneById(context.Background(), 0)

	// Assert
	s.Require().Nil(task)
	s.Require().Error(err)
	s.Require().ErrorIs(err, errPkg.DomainError)
}
