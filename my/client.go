package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type ApplicationJobHistory struct {
	gorm.Model
	Id         int64  `json:"id" gorm:"column:id;primary_key"`
	JobId      int64  `json:"job_id" gorm:"column:job_id"`
	ExecCode   int    `json:"exec_code" gorm:"column:exec_code"`
	Stdout     string `json:"stdout" gorm:"column:stdout"`
	Stderr     string `json:"stderr" gorm:"coOlumn:stderr"`
	CreateTime string `json:"createtime" gorm:"column:createtime"`
}
type DetectorServiceModule struct {
	RedisClient          *redis.Client //用于监听
	RedisClientForMc     *redis.Client //用于接口取数
	redisClientForReport *redis.Client //用于上报
}

func main() {
	// 连接到服务端建立的tcp连接
	dsn := "root2:w3323656@tcp(127.0.0.1:3306)/tm_platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&ApplicationJobHistory{})
	fmt.Print(time.Now().Format("2006:01:02 15:04:05"))
	detect := &DetectorServiceModule{
		RedisClient:          redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"}),
		RedisClientForMc:     redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"}),
		redisClientForReport: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"}),
	}
	detect.RedisClient.Set("ret", "eter", 0)

}
