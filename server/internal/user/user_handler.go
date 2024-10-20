package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{
Service
}

// returns a pointer to handler constructor 
func NewHandler(s Service) *Handler{
	return &Handler{
		Service:s,
	}
}
// returns a handler

//method
func(h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err:= c.ShouldBindJSON(&u); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"ShouldBindJSON error":err.Error()})
	}
	res,err:=h.Service.CreateUser(c.Request.Context(),&u)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"CreateUser error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,res)
}


