package usecase

import (
	"SigmatechTechnicalTest-Golang/api/repository"
	"SigmatechTechnicalTest-Golang/api/usecase/uc_authentication"
	"database/sql"
)

type UseCase struct {
	AuthenticationUC uc_authentication.AuthenticationUCInterface
}

func NewUseCase(db *sql.DB) UseCase {
	repo := repository.NewRepository(db)
	return UseCase{
		AuthenticationUC: uc_authentication.NewAuthenticationUC(repo),
	}
}
