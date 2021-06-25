package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

type SocketHandler struct {
}

func NewSocketHandler() SocketHandler {
	return SocketHandler{}
}

var webSocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (socketHandler SocketHandler) Setup(c *gin.Context) {

	//Upgrade get request to webSocket protocol
	ws, err := webSocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	//Read data in ws
	mt, message, err := ws.ReadMessage()
	if err != nil {
		log.Println("error read message")
		log.Fatal(err)
	}

	//Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}

		message = []byte(string(message) + " " + strconv.Itoa(count))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("error write message: " + err.Error())
		}
		time.Sleep(1 * time.Second)

	}
}