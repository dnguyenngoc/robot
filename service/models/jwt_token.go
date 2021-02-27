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
