package handlers

import (
	"github.com/gin-gonic/gin"
	"mydockerweb/docker"
	"strconv"
	"strings"
	"time"
)

func Info(c *gin.Context) {
	info := docker.Info()
	c.JSON(200, info)
}

type Container struct {
	ID   string
	Image string
	Name string
	Time string
	State string
	Status string
}

func ListContainer(c *gin.Context) {
	containers := docker.ListContainer()
	var cons = make([]Container, len(containers))
	for i, con := range containers {
		var name =con.Names[0]
		cons[i] = Container{con.ID[:12], con.Image,name, ptime(con.Created), con.State, con.Status}
	}
	c.JSON(200, cons)
}

func ContainerLog(c *gin.Context){
	id:=c.Param("id")
	docker.PrintLog(id, c.Writer)
}

func Inspect(c *gin.Context){
	id:=c.Param("id")
	var a = docker.Inspect(id)
	c.JSON(200, a)
}

func StartContainer(c *gin.Context) {

}

func StopContainer(c *gin.Context) {

}

func DeleteContainer(c *gin.Context) {

}

func CreateContainer(c *gin.Context) {

}

type Image struct {
	ID   string
	Name string
	Tag string
	Time string
	Size string
}

func ListImage(c *gin.Context) {
	images := docker.ListImage()
	var imgs = make([]Image, len(images))
	for i, img := range images {
		var name ="<none>"
		if img.RepoDigests!=nil && len(img.RepoDigests)>0 {
			a := strings.Split(img.RepoDigests[0], "@")
			name = a[0]
		}
		var tag = "<none>"
		if img.RepoTags!=nil && len(img.RepoTags)>0 {
			b := strings.Split(img.RepoTags[0], ":")
			tag = b[1]
		}
		imgs[i] = Image{"",name , tag, ptime(img.Created), size(img.Size)}
	}
	c.JSON(200, imgs)
}

func DeleteImage(c *gin.Context) {

}

func PullImage(c *gin.Context) {

}

func ptime(s int64)string{
	return time.Unix(s, 0).Format("2006-01-02 15:04:05")
}
func size(s int64) string {
	var ss = s / 1024
	if ss < 1024 {
		return strconv.Itoa(int(ss)) + "KB"
	} else {
		return strconv.Itoa(int(ss/1024)) + "MB"
	}
}
