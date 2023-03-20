package uc_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"errors"
	"strings"
)

func (uc TransactionUC) ListOrderUC(userID string) (res []models.TransactionListResponse, err error) {
	_, err = uc.Repo.AuthenticationRepo.GetUserByIDRepo(userID)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		return res, errors.New("user not found")
	} else if err != nil {
		return res, err
	}

	data, err := uc.Repo.TransactionRepo.GetListOrderRepo(userID)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		return res, nil
	} else if err != nil {
		return res, err
	}

	for _, d := range data {
		interestPrice := d.OTRPrice.Int64 * d.Interest.Int64 / 100
		res = append(res, models.TransactionListResponse{
			TransactionID: d.ID.String,
			OTRPrice:      d.OTRPrice.Int64,
			AdminFee:      d.Fee.Int64,
			Tenor:         d.Tenor.Int64,
			Interest:      interestPrice,
			AssetName:     d.AssetName.String,
			TotalPrice:    d.OTRPrice.Int64 + d.Fee.Int64 + interestPrice,
		})
	}

	return res, err
}
