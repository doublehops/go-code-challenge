package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"codechallenge.local/internal/pkg/handlers/device"
	"codechallenge.local/internal/pkg/types/api"
)

var (
	dh *device.Handler = device.New()
)

func main() {

	r := gin.Default()

	r.POST("/isgood", func(ctx *gin.Context) {

		var device api.DeviceCheckDetails
		if err := ctx.ShouldBindJSON(&device); err != nil {
			ctx.JSON(422, gin.H{"error": fmt.Sprintf("Unable to process request. %s", err.Error())})
		}

		err := dh.CheckDevice(&device)
		if err != nil {
			ctx.JSON(422, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"puppy": "true"})
	})

	r.Run()
}