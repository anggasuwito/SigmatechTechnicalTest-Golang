package uc_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"SigmatechTechnicalTest-Golang/api/repository"
)

type TransactionUC struct {
	Repo repository.Repository
}

type TransactionUCInterface interface {
	BuyUC(req models.TransactionBuyRequest) (res string, err error)
}

func NewTransactionUC(repo repository.Repository) TransactionUCInterface {
	return &TransactionUC{Repo: repo}
}
