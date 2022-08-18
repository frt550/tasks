package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	backupPkg "tasks/internal/pkg/core/backup"
	backupModelPkg "tasks/internal/pkg/core/backup/models"
	pb "tasks/pkg/api/backup"
	"time"
)

func New(backup backupPkg.Interface) pb.AdminServer {
	return &implementation{
		backup: backup,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	backup backupPkg.Interface
}

func (i *implementation) BackupCreate(ctx context.Context, _ *emptypb.Empty) (*pb.BackupResponse, error) {
	if backup, err := i.backup.Backup(ctx); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return createBackupResponse(backup), nil
	}
}

func createBackupResponse(backup *backupModelPkg.Backup) *pb.BackupResponse {
	return &pb.BackupResponse{
		Id:        uint64(backup.Id),
		Data:      backup.Data,
		CreatedAt: backup.CreatedAt.Format(time.RFC850),
	}
}
