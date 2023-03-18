package r_authentication

import "SigmatechTechnicalTest-Golang/api/models"

func (r AuthenticationRepo) GetUserByNIKRepo(nik string) (res models.User, err error) {
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
			FROM users u
			WHERE u.nik = ?
				AND u.deleted_at IS NULL
			ORDER BY u.created_at DESC
			LIMIT 1
		`, nik).Scan(
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
