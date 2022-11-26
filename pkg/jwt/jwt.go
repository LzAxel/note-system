package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	Secret   string
	TokenTTL time.Duration
}

func NewJWTManager(secret string, tokenTTL time.Duration) *JWTManager {
	return &JWTManager{Secret: secret, TokenTTL: tokenTTL}
}

func (m *JWTManager) NewJWT(accountId int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Subject:   strconv.Itoa(accountId),
			ExpiresAt: time.Now().Add(m.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	)
	signedToken, err := token.SignedString([]byte(m.Secret))
	if err != nil {
		return ""
	}

	return signedToken
}

func (m *JWTManager) Parse(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid sign method")
		}

		return []byte(m.Secret), nil
	})
	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}

func (m *JWTManager) Claims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return claims, errors.New("failed to parse claims")
	}

	return claims, nil
}
