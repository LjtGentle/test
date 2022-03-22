package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Req struct {
	Name string `json:"name" form:"name" binding:"required" `
	Address time.Time `json:"address" form:"address"`
}

func main () {
	r := gin.Default()
	r.GET("/time_test", func(c *gin.Context) {
		var req Req
		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(500,gin.H{
				"error":fmt.Sprintf("参数错误,err=%+v",err),
			})
			return
		}
		c.JSON(200,gin.H{
			"data":req,
		})
	})
	r.Run(":8080")
}