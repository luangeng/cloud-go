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
	c.JSON(200, Ok(nil))
}

func DeletePod(c *gin.Context) {
	k8s.DeletePod("default", "test")
	c.JSON(200, Ok(nil))
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

	handler := &WsStreamer{conn: ws}
	handler.cond = sync.NewCond(handler)

	go handler.Run()
	err = k8s.ExecPod(handler, handler, handler)
	if err != nil {
		fmt.Println(err.Error())
	}
}
