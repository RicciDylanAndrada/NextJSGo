package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct{

	hub *Hub
}

// constructor for hub
func NewHandler( h *Hub) *Handler{
	return &Handler{
		hub:h,
	}
}



type CreateRoomReq struct{
	ID string `json:"id"`
	Name string `json:"name"`

}

// reciever method for createRoom
func (h *Handler) CreateRoom(c *gin.Context){
var req CreateRoomReq
if err:=c.ShouldBindJSON(&req);err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{"error:":err.Error()})
	return 
}
h.hub.Rooms[req.ID]=&Room{
	ID:req.ID,
	Name:req.Name,
	Clients: make(map[string]*Client),
}
// store room in mem
c.JSON(http.StatusOK,req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http) bool{
		// origin := r.Header.Get("Origin")

		// return origin =="http://localhost:3000"
		return true
	},
}
func (h *Handler) JoinRoom(c *gin.Context){
conn,err:=upgrader.Upgrade(c.Writer,c.Request,nil)

if err!=nil{
	 c.JSON(http.StatusBadRequest,gin.H{"error joining room":err.Error{}})
	 return 
}

roomID:=c.Param("roomid")
clientID:=c.Param("userId")
username:=c.Param("username")
cl:=&Client{
	Conn: conn,
	Message: make(chan *Message,10),
	ID: clientID,
	RoomID: roomID,
	Username: username,
}
m:=&Message{
	Content:  username+ "new user joined the room",
	RoomID: roomID,
	Username: username,
	
}
// send client info to register channel

h.hub.Register<-cl


// register user to room
//broad cast


// send message to broadcast channel
h.hub.Broadcast<-m

//write message
//read message

}
