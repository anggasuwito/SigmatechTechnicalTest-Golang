package models

import "time"

type TransactionSetting struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Tenor     int       `json:"tenor"`
	Limit     int       `json:"limit"`
	Interest  int       `json:"interest"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
