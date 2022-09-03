//go:build integration

package test

import (
	"context"
	pb "tasks/pkg/api/task"
)

func (s *ApiSuite) TestAllMustReturnResponse() {
	task1, err1 := s.core.Create(context.Background(), "New task")
	task2, err2 := s.core.Create(context.Background(), "New task")
	req := &pb.TaskAllRequest{}
	expected := &pb.TaskAllResponse{
		Tasks: []*pb.TaskResponse{
			{
				Id:          task1.Id,
				Title:       task1.Title,
				IsCompleted: task1.IsCompleted,
				CreatedAt:   task1.CreatedAt,
				CompletedAt: task1.CompletedAt,
			},
			{
				Id:          task2.Id,
				Title:       task2.Title,
				IsCompleted: task2.IsCompleted,
				CreatedAt:   task2.CreatedAt,
				CompletedAt: task2.CompletedAt,
			},
		},
	}

	// Act
	res, err3 := s.api.TaskAll(context.Background(), req)

	// Assert
	s.Require().Nil(err1)
	s.Require().Nil(err2)
	s.Require().Nil(err3)
	s.Require().Equal(expected, res)
}
