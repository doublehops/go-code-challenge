package main

import (
	"github.com/gin-gonic/gin"

	"codechallenge.local/internal/pkg/handlers/device"
)

var (
	dh *device.Handler = device.New()
)

func main() {

	r := gin.Default()

	r.POST("/isgood", func(ctx *gin.Context) {

		err := dh.CheckDevice(ctx)
		if err != nil {
			ctx.JSON(422, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"puppy": true})
	})

	r.Run()
}