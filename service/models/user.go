/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
    FirstName    	*string            `json:"first_name" validate:"required,min=2,max=100"`
    LastName     	*string            `json:"last_name" validate:"required,min=2,max=100"`
    Password      	*string            `json:"password" validate:"required,min=6""`
    Email         	*string            `json:"email" validate:"email,required"`
    Phone         	*string            `json:"phone" validate:"required"`
}

type UserCreate struct {
    User
    CreatedAt    	time.Time          `json:"created_at"`
    UpdatedAt    	time.Time          `json:"updated_at"`
}

type UserUpdate struct {
    FirstName    	*string            `json:"first_name" validate:"min=2,max=100"`
    LastName     	*string            `json:"last_name" validate:"min=2,max=100"`
    Password      	*string            `json:"password" validate:"min=6""`
    Email         	*string            `json:"email" validate:"email,required"`
    Phone         	*string            `json:"phone" validate:"required"`
    CreatedAt    	time.Time          `json:"created_at"`
    UpdatedAt    	time.Time          `json:"updated_at"`
}

type UserDB struct {
	UserCreate
	UserId       	string             `json:"user_id"`       
	ID            	primitive.ObjectID `bson:"_id"`
}