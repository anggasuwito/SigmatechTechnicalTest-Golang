package uc_authentication

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"SigmatechTechnicalTest-Golang/helper/constants"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func (uc AuthenticationUC) LoginUC(req models.LoginRequest) (res string, err error) {
	user, err := uc.Repo.AuthenticationRepo.GetUserByNIKRepo(req.NIK)
	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		return res, errors.New("user not found")
	} else if err != nil {
		return res, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constants.UserIDKey: user.ID.String,
		"exp":               jwt.NewNumericDate(time.Now().Add(10 * time.Hour)),
	})

	res, err = token.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	return res, err
}
