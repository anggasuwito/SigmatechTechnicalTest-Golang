package middleware

import (
	"SigmatechTechnicalTest-Golang/helper/constants"
	"SigmatechTechnicalTest-Golang/helper/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func VerifyToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.GetHeader("Authorization")
		bearer := strings.Split(header, "Bearer ")
		if len(bearer) <= 1 {
			response.JSONUnauthorized(response.Response{Ctx: context, Error: errors.New("missing authentication header")})
			return
		} else {
			token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
			})

			if err != nil {
				response.JSONUnauthorized(response.Response{Ctx: context, Error: err})
				return
			} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userID, _ := claims[constants.UserIDKey].(string)
				context.Set(constants.UserIDKey, userID)
			}
		}
	}
}
