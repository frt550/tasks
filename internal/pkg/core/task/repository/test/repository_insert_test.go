//go:build integration

package test

import (
	"context"
	"tasks/internal/pkg/core/task/models"
	timePkg "tasks/internal/pkg/core/time"
)

func (s *RepositorySuite) TestInsertMustInsertTask() {
	expected := &models.Task{
		Title:     "New task 1",
		CreatedAt: timePkg.NowUTCFormatted(),
	}

	// Act
	err1 := s.repository.Insert(context.Background(), expected)
	task, err2 := s.repository.FindOneById(context.Background(), expected.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Equal(expected, task)
}

func (s *RepositorySuite) TestInsertMustReturnErrorIfCreatedAtIsNotSet() {
	// Act
	err := s.repository.Insert(context.Background(), &models.Task{
		Title: "New task 1",
	})

	// Assert
	s.Require().NotNil(err)
}
