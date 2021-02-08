package main

import (
	"github.com/gin-gonic/gin"
	Kernel "github.com/star-inc/kaguya_kernel"
	TalkService "github.com/star-inc/kaguya_kernel/service/talk"
)

func main() {
	router := gin.Default()

	router.GET("/talk", func(c *gin.Context) {
		handler := Kernel.Run(TalkService.NewServiceInterface())
		err := handler.HandleRequest(c.Writer, c.Request)
		if err != nil {
			panic(err)
		}
	})

	err := router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
