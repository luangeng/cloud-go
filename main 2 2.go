package main

import (
	"github.com/gin-gonic/gin"
	"mydockerweb/handlers"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	target := "/Main.html"
	http.Redirect(w, r, target, http.StatusFound)
}

//func Logger(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		startTime := time.Now()
//		h.ServeHTTP(w, r)
//		endTime := time.Since(startTime)
//		log.Printf("%s %d %v", r.URL, r.Method, endTime)
//	})
//}

func main() {

	r := gin.New()
	r.Use(gin.Logger())

	docker := r.Group("/docker")
	{
		docker.GET("/info", handlers.Info)
		docker.GET("/container", handlers.ListContainer)
		docker.POST("/container", handlers.CreateContainer)
		docker.DELETE("/container", handlers.DeleteContainer)
		docker.GET("/container/stop/:id", handlers.StopContainer)
		docker.GET("/container/start/:id", handlers.StartContainer)
		docker.GET("/container/log/:id", handlers.ContainerLog)
		docker.GET("/container/inspect/:id", handlers.Inspect)

		docker.GET("/image", handlers.ListImage)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Static("/web", "web")

	r.Run()

}
