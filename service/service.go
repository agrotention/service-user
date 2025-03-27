package service

import (
	pb_user "github.com/agrotention/pb-user"
	"github.com/agrotention/service-user/repo"
)

type ServiceUser struct {
	repo *repo.RepoUser
	pb_user.UnimplementedServiceUserServer
}

func NewService(repo *repo.RepoUser) *ServiceUser {
	return &ServiceUser{repo: repo}
}
