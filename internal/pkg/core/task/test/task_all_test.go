//go:build integration

package test

import (
	"context"
)

func (s *ServiceSuite) TestAllMustReturnAllTasks() {
	_, err1 := s.core.Create(context.Background(), "New task 1")
	_, err2 := s.core.Create(context.Background(), "New task 2")

	// Act
	tasks, err3 := s.core.All(context.Background(), 0, 0)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal(2, len(tasks))
}
func (s *ServiceSuite) TestAllMustReturnAllTasksUsingLimit() {
	_, err1 := s.core.Create(context.Background(), "New task 1")
	_, err2 := s.core.Create(context.Background(), "New task 2")

	// Act
	tasks, err3 := s.core.All(context.Background(), 1, 0)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().True(len(tasks) == 1)
}

func (s *ServiceSuite) TestAllMustReturnAllTasksUsingOffset() {
	_, err1 := s.core.Create(context.Background(), "New task 1")
	_, err2 := s.core.Create(context.Background(), "New task 2")

	// Act
	tasks, err3 := s.core.All(context.Background(), 0, 1)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal(1, len(tasks))
}
