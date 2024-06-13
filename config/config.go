package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	sid   string
	App   *App
	Redis *Redis
}
type Server struct {
	Http      int `json:"http" yaml:"http"`
	Websocket int `json:"websocket" yaml:"websocket"`
	Tcp       int `json:"tcp" yaml:"tcp"`
}

func New(filename string) *Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var conf Config
	if yaml.Unmarshal(content, &conf) != nil {
		panic(fmt.Sprintf("解析config.yml读取错误：%v", err))
	}
	return &conf
}
