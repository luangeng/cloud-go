package handler

import (
	"bytes"
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

	var stdin bytes.Buffer
	var stdout, stderr bytes.Buffer
	go k8s.ExecPod(&stdin, stdout, stderr)
	// if err != nil {
	// 	log.Print("err:", err)
	// }
	fmt.Println("end of execpod")
	// ch := make(chan string)
	// ctx, cancel := context.WithCancel(context.Background())
	for {
		mt, msg, err := ws.ReadMessage()
		fmt.Println(string(msg))
		stdin.Write(msg)
		b, _ := stdout.ReadBytes('\n')
		fmt.Println(b)
		err = ws.WriteMessage(mt, b)
		if err != nil {
			fmt.Printf(err.Error())
		}
		time.Sleep(time.Second)
	}
}
