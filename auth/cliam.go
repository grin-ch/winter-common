package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenKey = "Authorization"
)

type RoleBase struct {
	Uid      int    `json:"uid"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex,omitempty"`
	Ip       string `json:"ip"`
}

type Cliams struct {
	RoleBase
	jwt.StandardClaims
}

func GenerateJWT(expires int, secret, issuer string, rBase RoleBase) (string, error) {
	now := time.Now()
	expire := now.Add(time.Duration(expires) * time.Second)
	claims := Cliams{
		RoleBase: rBase,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func ParseJWT(signed, token string) (*Cliams, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Cliams{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signed), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims == nil {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := tokenClaims.Claims.(*Cliams)
	if !ok || !tokenClaims.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}
	return claims, err
}
