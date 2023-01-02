package utils

import (
	"github.com/go-chi/jwtauth/v5"
)

type Token struct {
	Auth *jwtauth.JWTAuth
}

var token *Token = nil

func GetToken(secret string) *Token {
	if token == nil {
		token = &Token{
			Auth: jwtauth.New("HS256", []byte(secret), nil),
		}
	}
	return token
}

func (t *Token) MakeToken(userID string, roleName string) (string, error) {
	_, tokenString, err := t.Auth.Encode(map[string]interface{}{"userID": userID, "roleName": roleName})
	return tokenString, err
}
