package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(c *gin.Context) {
	// ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	log.Print("upgrade:", err)
	// 	return
	// }
	// defer ws.Close()
	// for {
	// 	mt, message, err := ws.ReadMessage()
	// 	if err != nil {
	// 		log.Println("read:", err)
	// 		break
	// 	}
	// 	log.Printf("recv: %s", message)
	// 	ch := make(chan string)
	// 	go k8s.LogPod2(context.TODO(), ch)
	// 	for m := range ch {
	// 		err = ws.WriteMessage(mt, []byte(m))
	// 		if err != nil {
	// 			fmt.Printf(err.Error())
	// 		}
	// 	}
	// }
	// fmt.Printf("end of handler")

}
