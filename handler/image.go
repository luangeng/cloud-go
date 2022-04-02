package handler

import "github.com/gin-gonic/gin"

func Uplaod(c *gin.Context) {
	header, err := c.FormFile("upload-key")
	if err != nil {
		panic(err)
	}
	dst := header.Filename
	err = c.SaveUploadedFile(header, dst)
	if err != nil {
		panic(err)
	}
}
