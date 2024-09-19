package main

import (
	"encoding/gob"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
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

type Person struct {
	Name string
	Age  int
}

func waitGroup() {

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go demo(i)
	}
	//阻塞,直到WaitGroup队列中所有任务执行结束时自动解除阻塞
	fmt.Println("开始阻塞")
	wg.Wait()
	fmt.Println("任务执行结束,解除阻塞")

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

var DB *gorm.DB

func init() {
	// 连接数据库
	//配置MySQL连接参数
	username := "root2"     //账号
	password := "w3323656"  //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "tm_platform" //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
}

// 保存数据的函数

func main() {

	strategyForUpdate := StrategyForUpdate{ClusterId: "ree", StrategyId: 1231, Result: 1, Path: "/tesss", CreateTime: time.Now()}
	//strategyForUpdate.id = 2
	//strategyForUpdate.clusterId = "rwerwe"
	//strategyForUpdate.strategyId = 1423232
	//strategyForUpdate.result = 0
	//strategyForUpdate.path = "/test"
	fmt.Println(strategyForUpdate)
	err := DB.Create(&strategyForUpdate).Error
	if err != nil {
		fmt.Println(err.Error())

	}
}

type StrategyForUpdate struct {
	Id           int64     `gorm:"column:id"`
	StrategyId   int64     `gorm:"column:strategy_id"`
	StrategyType string    `gorm:"column:strategy_type"`
	ClusterId    string    `gorm:"column:cluster_id"`
	Path         string    `gorm:"column:path"`
	Result       int32     `gorm:"column:result"`
	CreateTime   time.Time `json:"column:create_time"`
}

func (this *StrategyForUpdate) TableName() string {
	return "strategy_for_update"
}
