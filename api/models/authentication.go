package models

import "database/sql"

type LoginRequest struct {
	NIK string `json:"nik"`
}

type User struct {
	ID            sql.NullString `json:"id"`
	NIK           sql.NullString `json:"nik"`
	FullName      sql.NullString `json:"full_name"`
	LegalName     sql.NullString `json:"legal_name"`
	BirthPlace    sql.NullString `json:"birth_place"`
	BirthDate     sql.NullString `json:"birth_date"`
	Salary        sql.NullInt64  `json:"salary"`
	IdentityPhoto sql.NullString `json:"identity_photo"`
	SelfiePhoto   sql.NullString `json:"selfie_photo"`
	CreatedAt     sql.NullString `json:"created_at"`
	UpdatedAt     sql.NullString `json:"updated_at"`
	DeletedAt     sql.NullString `json:"deleted_at"`
}
