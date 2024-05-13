package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInput struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		//获取url路径参数 请求为http://localhost:8080/hello/100，用c.param
		//获取query，用c.Query
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi" + name,
		})
	})
	r.POST("/users", func(context *gin.Context) {
		var input UserInput
		//获取body中的单个值
		name := context.PostForm("name")
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		context.JSON(http.StatusOK, gin.H{
			"name":     input.Name,
			"age":      input.Age,
			"bodyName": name,
		})
	})
	r.Run()
}
