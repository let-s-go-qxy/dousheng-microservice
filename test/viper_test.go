package test

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

type Config struct {
	Database Database
}

type Database struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type Mysql struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	Addr        string `yaml:"addr"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	DbVerify    int    `yaml:"dbVerify"`
	DbUserLike  int    `yaml:"dbUserLike"`
	DbVideoLike int    `yaml:"dbVideoLike"`
}

func ParseYaml() {
	path, _ := os.Getwd()
	var configViperConfig = viper.New()
	configViperConfig.SetConfigName("config")
	configViperConfig.AddConfigPath(path + "/../conf/")
	configViperConfig.SetConfigType("yaml")
	//读取配置文件内容
	if err := configViperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	var c Config
	if err := configViperConfig.Unmarshal(&c); err != nil {
		panic(err)
	}
	fmt.Println(c.Database.Redis)
	fmt.Println(c.Database.Mysql)
}

func TestViper(t *testing.T) {
	ParseYaml()
}
