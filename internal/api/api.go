package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/models"
	pb "tasks/pkg/api"
	"time"
)

func New(task taskPkg.Interface) pb.AdminServer {
	return &implementation{
		task: task,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	task taskPkg.Interface
}

func (i *implementation) TaskCreate(_ context.Context, in *pb.TaskCreateRequest) (*emptypb.Empty, error) {
	if err := i.task.Create(models.Task{
		Title: in.GetTitle(),
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (i *implementation) TaskRead(_ context.Context, in *pb.TaskReadRequest) (*pb.TaskReadResponse, error) {
	task, err := i.task.Get(uint(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := &pb.Task{
		Id:          uint64(task.Id),
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt.Format(time.RFC850),
		CompletedAt: task.CompletedAt.Format(time.RFC850),
	}
	return &pb.TaskReadResponse{
		Task: result,
	}, nil
}

func (i *implementation) TaskUpdate(_ context.Context, in *pb.TaskUpdateRequest) (*emptypb.Empty, error) {
	if err := i.task.Update(models.Task{
		Title: in.GetTitle(),
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (i *implementation) TaskDelete(_ context.Context, in *pb.TaskDeleteRequest) (*emptypb.Empty, error) {
	if err := i.task.Delete(uint(in.GetId())); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (i *implementation) TaskList(_ context.Context, _ *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	tasks := i.task.List()
	result := make([]*pb.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, &pb.Task{
			Id:          uint64(task.Id),
			Title:       task.Title,
			IsCompleted: task.IsCompleted,
			CreatedAt:   task.CreatedAt.Format(time.RFC850),
			CompletedAt: task.CompletedAt.Format(time.RFC850),
		})
	}
	return &pb.TaskListResponse{
		Tasks: result,
	}, nil
}
