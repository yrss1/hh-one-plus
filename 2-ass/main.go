package main

import (
	"2-ass/handlers"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	go func() {
		for {
			handlers.InsertData()
			time.Sleep(10 * time.Minute)
		}
	}()
	handlers.InsertData()

	rtr := gin.Default()
	api := rtr.Group("/api")
	{
		api.GET("/coins", handlers.GetCoins)
		api.GET("/coins/:id", handlers.GetCoin)
	}
	rtr.Run("localhost:8088")
}
