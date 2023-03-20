package models

import (
	"database/sql"
)

type TransactionBuyRequest struct {
	UserID               string `json:"user_id"`
	Tenor                int    `json:"tenor"`
	CompanyAssetID       string `json:"company_asset_id"`
	TransactionSettingID string `json:"transaction_setting_id"`
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

type TransactionList struct {
	ID        sql.NullString `json:"id"`
	OTRPrice  sql.NullInt64  `json:"otr_price"`
	Fee       sql.NullInt64  `json:"fee"`
	Tenor     sql.NullInt64  `json:"tenor"`
	Interest  sql.NullInt64  `json:"interest"`
	AssetName sql.NullString `json:"asset_name"`
}

type TransactionListResponse struct {
	TransactionID string `json:"transaction_id"`
	OTRPrice      int64  `json:"otr_price"`
	AdminFee      int64  `json:"admin_fee"`
	Tenor         int64  `json:"tenor"`
	Interest      int64  `json:"interest"`
	AssetName     string `json:"asset_name"`
	TotalPrice    int64  `json:"total_price"`
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
