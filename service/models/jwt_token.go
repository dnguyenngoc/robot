package models

import jwt "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

type Token struct {
	AccessToken 	string      `json:"access_token" validate:"required"`
	FreshToken	string      `json:"fresh_token" validate:"required"`
}