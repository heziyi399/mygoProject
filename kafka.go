package mygoProject

import (
	"fmt"
)

func Producer() {
	//byte("hello")：这部分是类型转换操作，它告诉 Go 编译器将字符串 "hello" 转换成一个字节切片。
	//p := kafka.NewBytes([]byte("hello"))
	bytes := []byte("hello")
	fmt.Print(string(bytes))

}
