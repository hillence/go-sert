package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startHTTPServer() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Telegram bot is running",
		})
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	addr := fmt.Sprintf("0.0.0.0:%s", AppPort)
	log.Printf("HTTP server listening on %s", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf("http server stopped: %v", err)
	}
}
