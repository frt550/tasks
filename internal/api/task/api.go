package api

import (
	"context"
	taskPkg "tasks/internal/pkg/core/task"
	taskModelPkg "tasks/internal/pkg/core/task/models"
	pb "tasks/pkg/api/task"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (i *implementation) TaskCreate(ctx context.Context, in *pb.TaskCreateRequest) (*pb.TaskResponse, error) {
	if task, err := i.task.Create(ctx, in.GetTitle()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return createTaskResponse(task), nil
	}
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskResponse, error) {
	task, err := i.task.Get(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return createTaskResponse(task), nil
	}
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskResponse, error) {
	if task, err := i.task.UpdateTitle(ctx, in.GetId(), in.GetTitle()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return createTaskResponse(task), nil
	}
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskResponse, error) {
	if task, err := i.task.Delete(ctx, in.GetId()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return createTaskResponse(task), nil
	}
}

func (i *implementation) TaskComplete(ctx context.Context, in *pb.TaskCompleteRequest) (*pb.TaskResponse, error) {
	if task, err := i.task.Complete(ctx, in.GetId()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return createTaskResponse(task), nil
	}
}

func (i *implementation) TaskAll(ctx context.Context, in *pb.TaskAllRequest) (*pb.TaskAllResponse, error) {
	tasks, err := i.task.All(ctx, in.Limit, in.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	taskResponses := make([]*pb.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		taskResponses = append(taskResponses, createTaskResponse(task))
	}
	return &pb.TaskAllResponse{
		Tasks: taskResponses,
	}, nil
}

func createTaskResponse(task *taskModelPkg.Task) *pb.TaskResponse {
	return &pb.TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		CompletedAt: task.CompletedAt,
	}
}
