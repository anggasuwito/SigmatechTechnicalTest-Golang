package r_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"database/sql"
)

func (r TransactionRepo) TXCreateTransactionHistory(tx *sql.Tx, req models.TransactionBuyRequest) (err error) {
	stmt, err := tx.
		Prepare(`
			INSERT INTO transaction_history
			(user_id,transaction_setting_id,company_asset_id)
				VALUES
			(?,?,?)
		`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(req.UserID, req.TransactionSettingID, req.CompanyAssetID)
	return err
}
