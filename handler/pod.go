package handler

import (
	model "cloud/model"
	k8s "cloud/vender/k8s"
	"context"
	"fmt"
	"log"
	"net/http"
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
		for {
			mt, message, err = ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
			}
			fmt.Println("read")
			if string(message) == "close" {
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

}
