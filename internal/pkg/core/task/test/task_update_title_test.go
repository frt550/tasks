//go:build integration

package test

import (
	"context"
)

func (s *ServiceSuite) TestUpdateTitleMustUpdateTitle() {
	task, err1 := s.core.Create(context.Background(), "New task")
	expected := "Updated task"

	// Act
	task, err2 := s.core.UpdateTitle(context.Background(), task.Id, expected)
	dbTask, err3 := s.repository.FindOneById(context.Background(), task.Id)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal(expected, dbTask.Title)
}
