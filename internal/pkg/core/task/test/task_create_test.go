package test

import (
	"context"
)

func (s *ServiceSuite) TestCreateMustCreateTask() {
	// Act
	expected, err1 := s.core.Create(context.Background(), "New task")
	dbTask, err2 := s.repository.FindOneById(context.Background(), expected.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Equal(expected, dbTask)
}
