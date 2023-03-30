package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-clean-architecture/db"
	"go-clean-architecture/routers"
	"net/http"
)

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, DisableColors: false, FullTimestamp: true})
	gin.ForceConsoleColor()
	r := gin.Default()
	DB := db.Init()
	routers.Setup(DB, r)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
