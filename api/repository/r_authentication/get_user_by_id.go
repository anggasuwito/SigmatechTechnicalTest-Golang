package r_authentication

import "SigmatechTechnicalTest-Golang/api/models"

func (r AuthenticationRepo) GetUserByIDRepo(id string) (res models.User, err error) {
	err = r.DB.
		QueryRow(`
			SELECT 
				u.id,
				u.nik,
				u.full_name,
				u.legal_name,
				u.birth_place,
				u.birth_date,
				u.salary,
				u.identity_photo,
				u.selfie_photo,
				u.created_at,
				u.updated_at,
				u.deleted_at
			FROM user u
			WHERE u.id = ?
				AND u.deleted_at IS NULL
			ORDER BY u.created_at DESC
			LIMIT 1
		`, id).Scan(
		&res.ID,
		&res.NIK,
		&res.FullName,
		&res.LegalName,
		&res.BirthPlace,
		&res.BirthDate,
		&res.Salary,
		&res.IdentityPhoto,
		&res.SelfiePhoto,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	return res, err
}
