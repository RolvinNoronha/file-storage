package models

import "github.com/golang-jwt/jwt/v5"

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	UserId    string `json:"userId"`
	JwtClaims jwt.MapClaims
}
