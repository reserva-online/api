package authentication

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const userIdClaimKey = "user_id"
const emailClaimKey = "email"

func CreateToken(userid int) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims[userIdClaimKey] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetTokenUser(tokenString string) (int, error) {
	claims, err := GetFormattedToken(tokenString)
	if err != nil {
		return 0, err
	}
	var floatUserId float64
	for key, val := range claims {
		if userIdClaimKey == key {
			floatUserId = val.(float64)
		}
	}
	return int(floatUserId), nil
}

func GetFormattedToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func CreateRecoverToken(email string) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims[emailClaimKey] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetEmailRecoverToken(tokenString string) (string, error) {
	claims, err := GetFormattedToken(tokenString)
	if err != nil {
		return "", err
	}
	var email string
	for key, val := range claims {
		if emailClaimKey == key {
			email = val.(string)
		}
	}
	return email, nil
}
