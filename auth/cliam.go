package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenKey  = "Authorization"
	ClaimsKey = "Cliams"
)

type Cliams struct {
	roleBase
	jwt.StandardClaims
	Salt string `json:"salt"`
}

func GenerateJWT(expires int, secret, issuer string, rBase roleBase, salt string) (string, error) {
	now := time.Now()
	expire := now.Add(time.Duration(expires) * time.Second)
	claims := Cliams{
		roleBase: rBase,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    issuer,
		},
		Salt: salt,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func ParseJwtWithSalt(signed, token string, salt string) (*Cliams, error) {
	claims, err := ParseJwt(signed, token)
	if err != nil {
		return nil, err
	}
	if claims.Salt != salt {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, err
}

func ParseJwt(signed, token string) (*Cliams, error) {
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

func IntoCtx(ctx context.Context, cliams *Cliams) context.Context {
	data, _ := json.Marshal(cliams)
	return context.WithValue(ctx, ClaimsKey, string(data))
}

func FromCtx(ctx context.Context) (*Cliams, error) {
	val := ctx.Value(ClaimsKey)
	src, ok := val.(string)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	c := Cliams{}
	err := json.Unmarshal([]byte(src), &c)
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	return &c, nil
}
