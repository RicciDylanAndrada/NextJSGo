// interacts with DB and runs query
package user

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

type User struct{
	ID int64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`

}

type CreateUserReq struct{
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type CreateUserRes struct{
	ID string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
}
type LoginUserReq struct{
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type LoginUserRes struct{
	accessToken string 
	Username string `json:"username" db:"username"`
	ID string `json:"id" db:"id"`

}

type Repository interface {
	//returns behavior 
	CreateUser(ctx context.Context,user *User) (*User,error)
	GetUserByEmail(ctx context.Context, email string) ( *User ,error)
}

type Service interface{
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes,error)
	Login(c context.Context,req *LoginUserReq)(*LoginUserRes,error)
}

