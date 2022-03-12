package utils

import (
	"log"

	"github.com/google/uuid"
	"gitlab.com/golang-team-template/monolith/pkg/jwt"
)

//GetUserIDFromToken ...
func GetUserIDFromToken(accessToken string) (uuid.UUID, error) {

	claims, err := jwt.ExtractClaims(accessToken, []byte(conf.JWTSecretKey))
	if err != nil {
		log.Println("could not extract claims:", err)
		return uuid.Nil, err
	}
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}
