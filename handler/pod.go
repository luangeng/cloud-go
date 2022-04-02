package handler

import (
	model "cloud/model"
	k8s "cloud/vender/k8s"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ListPod(c *gin.Context) {
	list := k8s.ListPod("default")
	c.JSON(200, list)
}

func CreatePod(c *gin.Context) {
	var param model.Pod
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	k8s.CreatePod(param)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeletePod(c *gin.Context) {
	k8s.DeletePod("default", "test")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func LogPod(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	var mt int
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		var message []byte
		var i = 0
		for {
			mt, message, err = ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
			}
			i++
			fmt.Println("read")
			if string(message) == "re" {
				i--
			}
			if string(message) == "close" || i > 60 {
				cancel()
				break
			}
			time.Sleep(time.Second)
		}
	}()

	ch := make(chan string)
	go k8s.LogPodFollow(ctx, ch)
	for m := range ch {
		err = ws.WriteMessage(mt, []byte(m))
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

	fmt.Printf("end of handler")
}

func ExecPod(c *gin.Context) {
	msg, err := k8s.ExecPodOnce(nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, msg)
}

func ExecPodWs(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	handler := &WSStreamHandler{conn: ws}
	handler.cond = sync.NewCond(handler)

	go handler.Run()
	err = k8s.ExecPod(handler, handler, handler)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type WSStreamHandler struct {
	buff []byte
	cond *sync.Cond

	sync.Mutex

	conn *websocket.Conn
}

func (h *WSStreamHandler) Run() {
	for {
		_, msg, err := h.conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		h.Lock()
		h.buff = append(h.buff, msg...)
		h.cond.Signal()
		h.Unlock()
		if err != nil {
			panic(err)
		}
	}
}

func (h *WSStreamHandler) Read(b []byte) (size int, err error) {
	h.Lock()
	for len(h.buff) == 0 {
		h.cond.Wait()
	}
	size = copy(b, h.buff)
	h.buff = h.buff[size:]
	h.Unlock()
	return
}

func (h *WSStreamHandler) Write(b []byte) (size int, err error) {
	size = len(b)
	fmt.Println(string(b))
	err = h.conn.WriteMessage(websocket.TextMessage, b)
	return
}
