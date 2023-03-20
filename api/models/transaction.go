package models

import (
	"database/sql"
)

type TransactionBuyRequest struct {
	UserID         string `json:"user_id"`
	Tenor          int    `json:"tenor"`
	CompanyAssetID string `json:"company_asset_id"`
}

type TransactionSetting struct {
	ID        sql.NullString `json:"id"`
	UserID    sql.NullString `json:"user_id"`
	Tenor     sql.NullInt64  `json:"tenor"`
	MaxLimit  sql.NullInt64  `json:"max_limit"`
	Interest  sql.NullInt64  `json:"interest"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

type CompanyAsset struct {
	ID                sql.NullString `json:"id"`
	CompanyID         sql.NullString `json:"company_id"`
	AssetName         sql.NullString `json:"asset_name"`
	OtrPrice          sql.NullInt64  `json:"otr_price"`
	StockAvailability sql.NullInt64  `json:"stock_availability"`
	CreatedAt         sql.NullString `json:"created_at"`
	UpdatedAt         sql.NullString `json:"updated_at"`
	DeletedAt         sql.NullString `json:"deleted_at"`
}
