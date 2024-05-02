package main

import (
	"io"
	"net/http"
	"os"

	"github.com/fesTech-cloud/gin/controller"
	"github.com/fesTech-cloud/gin/middleware"
	"github.com/fesTech-cloud/gin/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setUpLogerOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setUpLogerOutput()
	server := gin.New()
	// server.Static("/css", "./templates/css")
	// server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())
	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	// GROUPING APIS
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "input validation error", "error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video created successfully"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// {
	// 	"title": "Cool",
	// 	"description": "Hello cool",
	// 	"url": "https://www.youtube.com/watch?v=sXJXLq1lN7U&list=RDsXJXLq1lN7U&start_radio=1",
	// 	"actors": 3,
	// 	"author": {
	// 		"firstname": "lil wayne",
	// 		"lastname": "T-pain",
	// 		"email": "lilwayne@tpain.com",
	// 		"age": 45
	// 	}
	// }

	server.Run(":8080")
}
