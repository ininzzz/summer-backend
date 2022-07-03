package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func GenerateToken(UserID int64) (string, error) {
	now := time.Now()
	expire := now.Add(12 * time.Hour)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}).SignedString([]byte("golang"))
	if err != nil {
		logrus.Errorf("[GeneratrToken] err: %v", err.Error())
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
