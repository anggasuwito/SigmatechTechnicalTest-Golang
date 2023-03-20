package r_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
)

func (r TransactionRepo) GetListOrderRepo(userID string) (res []models.TransactionList, err error) {
	rows, err := r.DB.
		Query(`
			SELECT 
				th.id,
				ca.otr_price,
				c.company_fee,
				ts.tenor,
				ts.interest,
				ca.asset_name
			FROM transaction_history th
			LEFT JOIN transaction_setting ts ON ts.id = th.transaction_setting_id
			LEFT JOIN company_asset ca ON ca.id = th.company_asset_id
			LEFT JOIN company c ON c.id = ca.company_id
			WHERE th.user_id = ?
				AND th.deleted_at IS NULL
			ORDER BY th.created_at DESC
		`, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.TransactionList
		if err = rows.Scan(
			&item.ID,
			&item.OTRPrice,
			&item.Fee,
			&item.Tenor,
			&item.Interest,
			&item.AssetName,
		); err != nil {
			return res, err
		}

		res = append(res, item)
	}
	return res, err
}
