//go:build integration

package test

import (
	"context"
	timePkg "tasks/internal/pkg/core/time"
	pb "tasks/pkg/api/task"
)

func (s *ApiSuite) TestCompleteMustReturnResponse() {
	task, err1 := s.core.Create(context.Background(), "New task")
	req := &pb.TaskCompleteRequest{
		Id: task.Id,
	}
	// Act
	res, err2 := s.api.TaskComplete(context.Background(), req)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Equal(&pb.TaskResponse{
		Id:          req.Id,
		Title:       task.Title,
		IsCompleted: true,
		CreatedAt:   task.CreatedAt,
		CompletedAt: timePkg.NowUTCFormatted(),
	}, res)
}

func (s *ApiSuite) TestCompleteMustReturnError() {
	req := &pb.TaskCompleteRequest{
		Id: 0,
	}
	// Act
	res, err := s.api.TaskComplete(context.Background(), req)

	// Assert
	s.Require().Nil(res)
	s.Require().NotNil(err)
	s.Require().Equal("rpc error: code = Internal desc = Sorry, task #0 is not found: ", err.Error())
}
