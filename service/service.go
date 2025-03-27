package service

import (
	"context"

	pb "github.com/agrotention/pb-user"
	"github.com/agrotention/service-user/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceUser struct {
	repo *repo.RepoUser
	pb.UnimplementedServiceUserServer
}

func NewService(repo *repo.RepoUser) *ServiceUser {
	return &ServiceUser{repo: repo}
}

func (*ServiceUser) CreateUser(ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (*ServiceUser) GetPrivateDetailUser(ctx context.Context,
	req *pb.DetailUserRequest,
) (*pb.PrivateDetailUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateDetailUser not implemented")
}

func (*ServiceUser) GetPublicDetailUser(ctx context.Context,
	req *pb.DetailUserRequest,
) (*pb.PublicDetailUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicDetailUser not implemented")
}

func (*ServiceUser) GetPasswordHash(ctx context.Context,
	req *pb.GetPasswordHashRequest,
) (*pb.GetPasswordHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordHash not implemented")
}

func (*ServiceUser) GetListUser(ctx context.Context,
	req *pb.ListUserRequest,
) (*pb.ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListUser not implemented")
}

func (*ServiceUser) GetListDeletedUser(ctx context.Context,
	req *pb.ListUserRequest,
) (*pb.ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListDeletedUser not implemented")
}

func (*ServiceUser) UpdatePublicDetailUser(ctx context.Context,
	req *pb.UpdatePublicDetailUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublicDetailUser not implemented")
}

func (*ServiceUser) UpdateCredentialUser(ctx context.Context,
	req *pb.UpdateCredentialUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCredentialUser not implemented")
}

func (*ServiceUser) DisableUser(ctx context.Context,
	req *pb.DisableUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableUser not implemented")
}

func (*ServiceUser) EnableUser(ctx context.Context,
	req *pb.EnableUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableUser not implemented")
}

func (*ServiceUser) DeleteUser(ctx context.Context,
	req *pb.DeleteUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
