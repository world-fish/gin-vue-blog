package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

// InitCore 读取 YAML 配置文件并初始化配置。
// 该函数负责读取名为 "settings.yaml" 的配置文件，并将其内容解析到 config.Config 结构体中。
func InitCore() {
	// 定义配置文件的路径和名称
	const ConfigFile = "settings.yaml"

	// 创建一个指向 config.Config 结构体的指针，用于存储解析后的配置
	c := &config.Config{}

	// 使用 ioutil.ReadFile 读取 YAML 文件的内容
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		// 如果读取文件时发生错误，程序会终止并输出错误信息
		panic(fmt.Errorf("get yamlConf error:%v", err))
	}

	// 将读取的 YAML 文件内容解析到 Config 结构体中
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		// 如果解析 YAML 文件内容时发生错误，使用 log.Fatalf 打印错误信息并终止程序
		log.Fatalf("config Init Unmarshal: %v", err)
	}

	// 打印日志信息，表示配置文件成功加载并初始化
	log.Println("config yamlFile load Init success")

	global.Config = c
}
