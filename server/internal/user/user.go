// interacts with DB and runs query
package user

import "context"
type User struct{
	ID int64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`

}

type Repository interface {
	//returns behavior 
	CreateUser(ctx context.Context,user *User) (*User,error)
}