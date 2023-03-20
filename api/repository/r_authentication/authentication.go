package r_authentication

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"database/sql"
)

type AuthenticationRepo struct {
	DB *sql.DB
}

type AuthenticationRepoInterface interface {
	GetUserByNIKRepo(nik string) (res models.User, err error)
	GetUserByIDRepo(id string) (res models.User, err error)
}

func NewAuthenticationRepo(db *sql.DB) AuthenticationRepoInterface {
	return &AuthenticationRepo{DB: db}
}
