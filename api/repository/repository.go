package repository

import (
	"SigmatechTechnicalTest-Golang/api/repository/r_authentication"
	"SigmatechTechnicalTest-Golang/api/repository/r_transaction"
	"database/sql"
)

type Repository struct {
	AuthenticationRepo r_authentication.AuthenticationRepoInterface
	TransactionRepo    r_transaction.TransactionRepoInterface
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		AuthenticationRepo: r_authentication.NewAuthenticationRepo(db),
		TransactionRepo:    r_transaction.NewTransactionRepo(db),
	}
}
