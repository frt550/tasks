package transport

import (
	"context"
	"tasks/internal/config"
	"tasks/internal/pkg/core/task/models"
	pb "tasks/pkg/api/task"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Interface interface {
	All(ctx context.Context) ([]models.Task, error)
}

func New() (Interface, error) {
	conn, err := grpc.Dial(config.Config.Task.Grpc.ClientTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewAdminClient(conn)
	return &transport{client}, nil
}

type transport struct {
	client pb.AdminClient
}

func (t *transport) All(ctx context.Context) ([]models.Task, error) {
	taskAllResponse, err := t.client.TaskAll(ctx, &pb.TaskAllRequest{})
	if err != nil {
		return make([]models.Task, 0), err
	}
	result := make([]models.Task, 0, len(taskAllResponse.Tasks))
	for _, taskResponse := range taskAllResponse.Tasks {
		task, err := createTask(taskResponse)
		if err != nil {
			return make([]models.Task, 0), err
		}
		result = append(result, task)
	}
	return result, nil
}

func createTask(taskResponse *pb.TaskResponse) (models.Task, error) {
	return models.Task{
		Id:          taskResponse.Id,
		Title:       taskResponse.Title,
		IsCompleted: taskResponse.IsCompleted,
		CreatedAt:   taskResponse.CreatedAt,
		CompletedAt: taskResponse.CompletedAt,
	}, nil
}
