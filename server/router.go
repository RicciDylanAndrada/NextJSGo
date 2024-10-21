package router

import (
	"server/internal/user"
	"server/internal/user/ws"

	"github.com/gin-gonic/gin"
)


var r*gin.Engine
func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler){
r=gin.Default()

r.POST("/signup",userHandler.CreateUser)
r.POST("/login",userHandler.Login)
r.POST("/logout",userHandler.Logout)


// ws endpoints
r.POST("/ws/createRoom",wsHandler.CreateRoom)

}
func Start(addr string) error{
	// starts on address:port
	return r.Run(addr)
}