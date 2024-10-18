package user

import (
	"context"
	"strconv"
	"time"
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

// hash 

u:= &User{
	Username: req.Username,
	Email: req.Email,
	Password:req.Password,
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