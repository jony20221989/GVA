package middleware

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"server/config"
	"server/global"
)

func InitViper(path ...string) *viper.Viper {
	var configs string

	if len(path) == 0 {
		flag.StringVar(&configs, "c", "", "choose config file.")
		flag.Parse()
		if configs == "" { // 判断命令行参数是否为空
			switch gin.Mode() {
			case gin.DebugMode:
				configs = config.ConfigDebugFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config.ConfigDebugFile)
			case gin.ReleaseMode:
				configs = config.ConfigReleaseFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config.ConfigReleaseFile)
			case gin.TestMode:
				configs = config.ConfigTestFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config.ConfigTestFile)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", configs)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		configs = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", configs)
	}

	v := viper.New()
	//制定配置文件的路径
	v.SetConfigFile("config/" + configs)
	//制定配置文件的格式
	v.SetConfigType("yaml")
	// 读取配置信息
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//监听修改
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
