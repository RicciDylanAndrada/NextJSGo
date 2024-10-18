package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string ,error){
	hashP,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",fmt.Errorf("failed to hash: %w",err)
	}

	return string(hashP),nil
} 
func CheckPassword(password string,hashP string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashP) ,[]byte(password))
}