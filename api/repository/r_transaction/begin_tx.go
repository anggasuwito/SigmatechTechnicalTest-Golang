package r_transaction

import "database/sql"

func (r TransactionRepo) BeginTX() (*sql.Tx, error) {
	return r.DB.Begin()
}
