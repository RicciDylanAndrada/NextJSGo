package user

import (
	"context"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const(
	secretKey="secret"
)

// service struct with timeout and repo as data model
type service struct{
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service{
	return &service{
		repository,
		time.Duration(2)* time.Second,

	}
}

// reciever function ( method) for service struct
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes,error){
ctx,cancel:=context.WithTimeout(c, s.timeout)
defer cancel()

hashP,err:=util.HashPassword(req.Password)
if err!=nil{
	return nil,err
}

u:= &User{
	Username: req.Username,
	Email: req.Email,
	Password:hashP,
}
r,err:=s.Repository.CreateUser(ctx,u)
// error handling 
if(err!=nil){
	return nil,err
}

// appoint new CreateUserRes pointer to res 
res:=&CreateUserRes{
	ID:strconv.Itoa(int(r.ID)),
	Username:r.Username,
	Email:r.Email,
}
return res,nil
}

type MyJwtClaims struct{
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims

}




func (s *service) Login(c context.Context,req *LoginUserReq)(*LoginUserRes,error){
	signedKey := goDotEnvVariable("SIGNED_KEY")
	ctx,cancel := context.WithTimeout(c,s.timeout)
	defer cancel()

	u,err:=s.Repository.GetUserByEmail(ctx,req.Email)
	if err!=nil{
		return &LoginUserRes{},err
	}
	err = util.CheckPassword(req.Password,u.Password)
	if err!=nil{
	return &LoginUserRes{},err
	}
	var idString=strconv.Itoa(int(u.ID))

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,MyJwtClaims{
		ID:idString,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:idString,
			ExpiresAt:jwt.NewNumericDate((time.Now().Add(24*time.Hour))),

		},

	})
	ss,err:=token.SignedString([]byte(signedKey))
	if err!=nil{
		return &LoginUserRes{},err
	}
	return &LoginUserRes{accessToken: ss,Username: u.Username,ID:idString},nil

}