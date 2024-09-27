package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
