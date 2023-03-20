package r_transaction

import "SigmatechTechnicalTest-Golang/api/models"

func (r TransactionRepo) GetTransactionSettingRepo(userID string, tenor int) (res models.TransactionSetting, err error) {
	err = r.DB.
		QueryRow(`
			SELECT 
				ts.id,
				ts.user_id,
				ts.tenor,
				ts.max_limit,
				ts.interest,
				ts.created_at,
				ts.updated_at,
				ts.deleted_at
			FROM transaction_setting ts
			WHERE ts.tenor = ?
			  	AND ts.user_id = ?
				AND ts.deleted_at IS NULL
			ORDER BY ts.created_at DESC
			LIMIT 1
		`, tenor, userID).Scan(
		&res.ID,
		&res.UserID,
		&res.Tenor,
		&res.MaxLimit,
		&res.Interest,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	return res, err
}
