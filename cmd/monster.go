package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"monster-base-backend/internal/service"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	e := gin.Default()
	startService := service.NewStartService(e)
	startService.Start()
	e.Run(":9900")
}
