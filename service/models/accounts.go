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
    First_name    	string            `json:"first_name" validate:"required,min=2,max=100"`
    Last_name     	string            `json:"last_name" validate:"required,min=2,max=100"`
    Password      	string            `json:"password" validate:"required,min=6""`
    Email         	string            `json:"email" validate:"email,required"`
    Phone         	string            `json:"phone" validate:"required"`
}

type UserCreate struct {
    First_name    	string            `json:"first_name" validate:"required,min=2,max=100"`
    Last_name     	string            `json:"last_name" validate:"required,min=2,max=100"`
    Password      	string            `json:"password" validate:"required,min=6""`
    Email         	string            `json:"email" validate:"email,required"`
    Phone         	string            `json:"phone" validate:"required"`
    Created_at    	time.Time          `json:"created_at"`
    Updated_at    	time.Time          `json:"updated_at"`
}

type UserUpdate struct {
    First_name    	string            `json:"first_name" validate:"min=2,max=100"`
    Last_name     	string            `json:"last_name" validate:"min=2,max=100"`
    Password      	string            `json:"password" validate:"min=6""`
    Email         	string            `json:"email" validate:"email,required"`
    Phone         	string            `json:"phone" validate:"required"`
    Created_at    	time.Time          `json:"created_at"`
    Updated_at    	time.Time          `json:"updated_at"`
}

type UserDB struct {
	First_name    	string            `json:"first_name" validate:"required,min=2,max=100"`
    Last_name     	string            `json:"last_name" validate:"required,min=2,max=100"`
    Password      	string            `json:"password" validate:"required,min=6""`
    Email         	string            `json:"email" validate:"email,required"`
    Phone         	string            `json:"phone" validate:"required"`
	User_id       	string             `json:"user_id"`   
    Created_at    	time.Time          `json:"created_at"`
    Updated_at    	time.Time          `json:"updated_at"`    
	ID            	primitive.ObjectID `bson:"_id"`
}

type UserLogin struct {
	Password      	string            `json:"password" validate:"required,min=6""`
    Email         	string            `json:"email" validate:"email,required"`
}