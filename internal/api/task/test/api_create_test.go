//go:build integration

package test

import (
	"context"
	timePkg "tasks/internal/pkg/core/time"
	pb "tasks/pkg/api/task"
)

func (s *ApiSuite) TestCreateMustReturnResponse() {
	req := &pb.TaskCreateRequest{
		Title: "New task",
	}
	// Act
	res, err := s.api.TaskCreate(context.Background(), req)

	// Assert
	s.Require().Nil(err)
	s.Require().Greater(res.Id, uint64(0))
	s.Require().Equal(&pb.TaskResponse{
		Id:          res.Id,
		Title:       req.Title,
		IsCompleted: false,
		CreatedAt:   timePkg.NowUTCFormatted(),
		CompletedAt: "",
	}, res)
}
