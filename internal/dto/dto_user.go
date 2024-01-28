package dto

import "github.com/dgrijalva/jwt-go"

type CreateUser struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type JwtUserInfo struct {
	Username string
	Role     string
}

type JwtClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
