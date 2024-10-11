package main

import (
	"fmt"
	"log"
	"server/db"
)

func main(){
	_,err:=db.NewDataBase()
	if(err!=nil){
		log.Fatalf("could not initialize db connection : %s", err)
	}else{
		fmt.Println("could  initialize db connection");

	}
}