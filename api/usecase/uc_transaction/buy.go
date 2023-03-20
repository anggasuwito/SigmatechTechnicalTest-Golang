package uc_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"errors"
	"fmt"
	"strings"
)

func (uc TransactionUC) BuyUC(req models.TransactionBuyRequest) (res string, err error) {
	_, err = uc.Repo.AuthenticationRepo.GetUserByIDRepo(req.UserID)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		return res, errors.New("user not found")
	} else if err != nil {
		return res, err
	}

	transactionSetting, err := uc.Repo.TransactionRepo.GetTransactionSettingRepo(req.UserID, req.Tenor)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		return res, fmt.Errorf("you are not allow to request with %v tenor", req.Tenor)
	} else if err != nil {
		return res, err
	}
	req.TransactionSettingID = transactionSetting.ID.String

	tx, _ := uc.Repo.TransactionRepo.BeginTX()
	companyAsset, err := uc.Repo.TransactionRepo.TXGetAvailCompanyAsset(tx, req.CompanyAssetID)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		tx.Rollback()
		return res, errors.New("company asset not found")
	} else if err != nil {
		tx.Rollback()
		return res, err
	} else if companyAsset.StockAvailability.Int64 == 0 {
		tx.Rollback()
		return res, errors.New("this asset is not availability in this moment")
	}

	err = uc.Repo.TransactionRepo.TXUpdateCompanyAsset(tx, req.CompanyAssetID)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	err = uc.Repo.TransactionRepo.TXCreateTransactionHistory(tx, req)
	if err != nil {
		tx.Rollback()
		return res, err
	}
	tx.Commit()
	return res, err
}
