package api

import (
	"context"
	backupPkg "tasks/internal/pkg/core/backup"
	backupModelPkg "tasks/internal/pkg/core/backup/models"
	pb "tasks/pkg/api/backup"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (i *implementation) AsyncBackupCreate(ctx context.Context, in *pb.AsyncBackupCreateRequest) (*pb.AsyncBackupCreateResponse, error) {
	if backup, err := i.backup.AsyncBackup(ctx, in.RequestId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return createAsyncBackupResponse(backup), nil
	}
}

func createBackupResponse(backup *backupModelPkg.Backup) *pb.BackupResponse {
	if backup == nil {
		return nil
	}
	return &pb.BackupResponse{
		Id:        backup.Id,
		Data:      backup.Data,
		CreatedAt: backup.CreatedAt,
	}
}

func createAsyncBackupResponse(asyncBackup *backupModelPkg.AsyncBackup) *pb.AsyncBackupCreateResponse {
	return &pb.AsyncBackupCreateResponse{
		RequestId: asyncBackup.RequestId,
		State:     asyncBackup.State,
		Backup:    createBackupResponse(asyncBackup.Backup),
	}
}
