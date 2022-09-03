//go:build integration

package test

import (
	"context"
	pb "tasks/pkg/api/task"
)

func (s *ApiSuite) TestUpdateMustReturnResponse() {
	task, err1 := s.core.Create(context.Background(), "New task")
	req := &pb.TaskUpdateRequest{
		Id:    task.Id,
		Title: "Updated",
	}
	// Act
	res, err2 := s.api.TaskUpdate(context.Background(), req)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Equal(&pb.TaskResponse{
		Id:          req.Id,
		Title:       req.Title,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		CompletedAt: task.CompletedAt,
	}, res)
}

func (s *ApiSuite) TestUpdateMustReturnError() {
	req := &pb.TaskUpdateRequest{
		Id:    0,
		Title: "Updated",
	}
	// Act
	res, err := s.api.TaskUpdate(context.Background(), req)

	// Assert
	s.Require().Nil(res)
	s.Require().NotNil(err)
	s.Require().Equal("rpc error: code = Internal desc = Sorry, task #0 is not found: ", err.Error())
}
