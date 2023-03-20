package r_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"database/sql"
)

func (r TransactionRepo) TXGetAvailCompanyAsset(tx *sql.Tx, id string) (res models.CompanyAsset, err error) {
	err = tx.
		QueryRow(`
			SELECT 
					ca.id,
					ca.company_id,
					ca.asset_name,
					ca.otr_price,
					ca.stock_availability,
					ca.created_at,
					ca.updated_at,
					ca.deleted_at
			FROM company_asset ca
			LEFT JOIN company c ON c.id = ca.company_id
			WHERE ca.id = ?
				AND ca.deleted_at IS NULL
			ORDER BY ca.created_at DESC
			LIMIT 1
			FOR UPDATE
		`, id).Scan(
		&res.ID,
		&res.CompanyID,
		&res.AssetName,
		&res.OtrPrice,
		&res.StockAvailability,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	return res, err
}
