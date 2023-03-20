package usecase

import (
	"SigmatechTechnicalTest-Golang/api/repository"
	"SigmatechTechnicalTest-Golang/api/usecase/uc_authentication"
	"SigmatechTechnicalTest-Golang/api/usecase/uc_transaction"
	"database/sql"
)

type UseCase struct {
	AuthenticationUC uc_authentication.AuthenticationUCInterface
	TransactionUC    uc_transaction.TransactionUCInterface
}

func NewUseCase(db *sql.DB) UseCase {
	repo := repository.NewRepository(db)
	return UseCase{
		AuthenticationUC: uc_authentication.NewAuthenticationUC(repo),
		TransactionUC:    uc_transaction.NewTransactionUC(repo),
	}
}
