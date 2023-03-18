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
}
