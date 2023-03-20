package r_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"database/sql"
)

type TransactionRepo struct {
	DB *sql.DB
}

type TransactionRepoInterface interface {
	BeginTX() (*sql.Tx, error)
	GetListOrderRepo(userID string) (res []models.TransactionList, err error)
	GetTransactionSettingRepo(userID string, tenor int) (res models.TransactionSetting, err error)
	TXGetAvailCompanyAsset(tx *sql.Tx, id string) (res models.CompanyAsset, err error)
	TXUpdateCompanyAsset(tx *sql.Tx, id string) (err error)
	TXCreateTransactionHistory(tx *sql.Tx, req models.TransactionBuyRequest) (err error)
}

func NewTransactionRepo(db *sql.DB) TransactionRepoInterface {
	return &TransactionRepo{DB: db}
}
