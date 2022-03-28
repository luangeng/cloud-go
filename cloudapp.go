package main

import (
	// "cloud/handler"
	"net/http"

	handler "cloud/handler"

	client "cloud/vender/k8s"

	"github.com/gin-gonic/gin"
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

	client.Init()

	r := gin.New()
	r.Use(gin.Logger())

	cloud := r.Group("/cloud")
	{
		g1 := cloud.Group("/node")
		{
			g1.GET("/s", handler.ListNode)
			g1.GET("/d", handler.ListNodeDetail)
		}
		g6 := cloud.Group("/ns")
		{
			g6.GET("/", handler.ListNs)
			g6.POST("/", handler.CreateNs)
		}
		g2 := cloud.Group("/pod")
		{
			g2.GET("/", vvv)
		}
		g3 := cloud.Group("/service")
		{

			g3.GET("/", handler.ListService)
			g3.POST("/", handler.CreateService)
			g3.DELETE("/", handler.DeleteService)
		}
		g4 := cloud.Group("/pvc")
		{
			g4.GET("/s", handler.ListPv)
			g4.GET("/d", handler.ListPvDetail)
			g4.POST("/", handler.CreatePvc)
			g4.DELETE("/", handler.DeletePvc)
		}
		g5 := cloud.Group("/deploy")
		{
			g5.GET("/s", handler.ListDeploy)
			g5.GET("/d", handler.ListDeployDetail)
			g5.POST("/", handler.CreateDeploy)
			g5.DELETE("/", handler.DeleteDeploy)
		}

	}

	r.GET("/test", handler.Test)

	r.GET("/ok", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Static("/web", "web")

	r.Run()

}

func vvv(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
