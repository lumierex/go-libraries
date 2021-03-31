package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	w "samples/ws/internal/ws"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	r := gin.Default()

	manager := w.NewClientManager()
	go manager.Run()

	r.GET("/ws", func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return
		}
		id := ctx.Query("uid")
		defer conn.Close()

		c := &w.Client{
			Manager: manager,
			Conn:    conn,
			Send:    nil,
			Id:      id,
		}

		manager.Register <- c

		go c.Write()
		go c.Read()

		//for {
		//	msgType, msg, err := ws.ReadMessage()
		//	fmt.Println("receive: message: ", string(msg))
		//	if err != nil {
		//		break
		//	}
		//
		//	time.Sleep(time.Second * 2)
		//	err = ws.WriteMessage(msgType, []byte("hello world"))
		//	if err != nil {
		//		break
		//	}
		//}
	})

	r.Run(":8989")
}
