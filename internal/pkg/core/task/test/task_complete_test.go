//go:build integration

package test

import (
	"context"
)

func (s *ServiceSuite) TestCompleteMustCompleteTask() {
	task, err1 := s.core.Create(context.Background(), "New task")

	// Act
	task, err2 := s.core.Complete(context.Background(), task.Id)
	dbTask, err3 := s.repository.FindOneById(context.Background(), task.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().True(dbTask.IsCompleted)
}
