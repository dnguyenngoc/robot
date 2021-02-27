package models

import jwt "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Email      string
	FirstName string
	LastName  string
	UId        string
	UserType  string
	jwt.StandardClaims
}

type Token struct {
	AccessToken 	string      `json:"access_token" validate:"required"`
	FreshToken	string      `json:"fresh_token" validate:"required"`
}