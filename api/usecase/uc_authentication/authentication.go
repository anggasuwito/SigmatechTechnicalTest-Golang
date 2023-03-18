package uc_authentication

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"SigmatechTechnicalTest-Golang/api/repository"
)

type AuthenticationUC struct {
	Repo repository.Repository
}

type AuthenticationUCInterface interface {
	LoginUC(req models.LoginRequest) (res string, err error)
}

func NewAuthenticationUC(repo repository.Repository) AuthenticationUCInterface {
	return &AuthenticationUC{Repo: repo}
}
