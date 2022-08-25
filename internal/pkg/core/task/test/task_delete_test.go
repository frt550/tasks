package test

import (
	"context"
)

func (s *ServiceSuite) TestDeleteMustDeleteTask() {
	task, err1 := s.core.Create(context.Background(), "New task")

	// Act
	task, err2 := s.core.Delete(context.Background(), task.Id)
	dbTask, err3 := s.repository.FindOneById(context.Background(), task.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().NotNil(err3)
	s.Require().Nil(dbTask)
}
