package test

import (
	"context"
	pb "tasks/pkg/api/task"
)

func (s *ApiSuite) TestGetMustReturnResponse() {
	task, err1 := s.core.Create(context.Background(), "New task")
	req := &pb.TaskGetRequest{
		Id: task.Id,
	}
	// Act
	res, err2 := s.api.TaskGet(context.Background(), req)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Equal(&pb.TaskResponse{
		Id:          req.Id,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		CompletedAt: task.CompletedAt,
	}, res)
}

func (s *ApiSuite) TestGetMustReturnError() {
	req := &pb.TaskGetRequest{
		Id: 0,
	}
	// Act
	res, err := s.api.TaskGet(context.Background(), req)

	// Assert
	s.Require().Nil(res)
	s.Require().NotNil(err)
	s.Require().Equal("rpc error: code = Internal desc = Sorry, task #0 is not found: ", err.Error())
}
