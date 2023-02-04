/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2023/2/4 20:21
 * @Project_Name : scaffold
 * @File : settings.go
 * @Software :GoLand
 */

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level     string `mapstructure:"level"`
	Filename  string `mapstructure:"filename"`
	MaxSize   int    `mapstructure:"max_size"`
	MaxAge    int    `mapstructure:"max_age"`
	MaxBackup int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	DB          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	PoolSize    int    `mapstructure:"pool_size"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal failed:", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper.Unmarshal failed:", err)
		}
	})
	fmt.Println("Config init success")
	return
}
