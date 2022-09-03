//go:build integration

package test

import (
	"context"
)

func (s *ServiceSuite) TestGetMustReturnTask() {
	task, err1 := s.core.Create(context.Background(), "New task")

	// Act
	dbTask, err2 := s.repository.FindOneById(context.Background(), task.Id)
	task, err3 := s.core.Get(context.Background(), task.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal(dbTask, task)
}
