package main

import (
	"fmt"
	"log"
	router "server"
	"server/db"
	"server/internal/user"
	"server/internal/user/ws"
)

func main(){
	dbConn,err:=db.NewDataBase()
	if(err!=nil){
		log.Fatalf("could not initialize db connection : %s", err)
	}else{
		fmt.Println("could  initialize db connection");

	}
	userRep:=user.NewRepo(dbConn.GetDB())
	userSvc:=user.NewService((userRep))
	userHandler:=user.NewHandler(userSvc)

// hub initialization
	hub:=ws.NewHub()
	wsHandler:=ws.NewHandler(hub)

	go hub.Run()
	


	router.InitRouter(userHandler,wsHandler)
	router.Start("0.0.0.0:8080")
}