package mygoProject

import (
	"encoding/gob"
	"os"
)

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type UserInput struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func ginTest() {
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

var wg sync.WaitGroup

func main() {

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go demo(i)
	}
	//阻塞,直到WaitGroup队列中所有任务执行结束时自动解除阻塞
	fmt.Println("开始阻塞")
	wg.Wait()
	fmt.Println("任务执行结束,解除阻塞")
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"nick", 18}
	file, _ := os.Create("person.gob")
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err := encoder.Encode(person)
	if err != nil {
		fmt.Sprintf("encode fail")
	}
}
func demo(index int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("第%d次执行，i的值是：%d\n", index, i)
	}
	wg.Done()
}
