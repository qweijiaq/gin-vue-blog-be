package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
	"server/config"
	"server/global"
)

const ConfigFile = "settings_copy.yaml"

// InitConf 读取 YAML 文件的配置
func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile) // 读取 YAML 文件
	if err != nil {
		panic(fmt.Errorf("读取 YAML 文件发生错误: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c) // 反序列化
	if err != nil {
		log.Fatalf("反序列化 YAML 文件发生错误: %v", err)
	}
	log.Println("配置文件初始化加载成功")
	global.Config = c
}

// SetYaml 配置 YAML 文件
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
