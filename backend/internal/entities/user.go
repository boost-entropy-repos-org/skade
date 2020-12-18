package entities

import (
	"time"
	"github.com/google/uuid"
)

//DbCreateUserParams are all params needed to create a new user
type DbCreateUserParams struct {
	Username		string	`json:"username" db:"username"`
	HashedPassword 	string	`json:"hashed_password" db:"hashed_password"`
	Email			string	`json:"email" db:"email"`
}

//User is the User Type for the entire skade application
type User struct {
	ID        		uuid.UUID     	`json:"id" db:"id"`
	Username  		string       	`json:"username" db:"username"`
	HashedPassword 	string			`json:"hashed_password" db:"hashed_password"`
	Email     		string       	`json:"email" db:"email"`
	CreatedAt 		time.Time 		`json:"created_at" db:"created_at"`
}