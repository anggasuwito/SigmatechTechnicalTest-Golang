package models

import "time"

type LoginRequest struct {
	NIK string `json:"nik"`
}

type User struct {
	ID            string    `json:"id"`
	NIK           string    `json:"nik"`
	FullName      string    `json:"full_name"`
	LegalName     string    `json:"legal_name"`
	BirthPlace    string    `json:"birth_place"`
	BirthDate     string    `json:"birth_date"`
	Salary        int       `json:"salary"`
	IdentityPhoto string    `json:"identity_photo"`
	SelfiePhoto   string    `json:"selfie_photo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
