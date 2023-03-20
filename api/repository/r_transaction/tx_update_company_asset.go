package r_transaction

import (
	"database/sql"
)

func (r TransactionRepo) TXUpdateCompanyAsset(tx *sql.Tx, id string) (err error) {
	_, err = tx.
		Exec(`
			UPDATE
			    company_asset
			SET
			    stock_availability = stock_availability - 1
			WHERE id = ?
				AND deleted_at IS NULL
		`, id)
	return err
}
