package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"lottery_wechat/configs"
	"lottery_wechat/router"
)

func Init() {
	configs.InitGlobalConfig()
}

func main() {
	config := configs.GetGlobalConfig()
	fmt.Println(config)

	Init()
	log.Infof("这是我打印的Info日志！")

	r := router.SetRouter()
	if err := r.Run(fmt.Sprintf(":%d", config.AppConfig.Port)); err != nil {
		log.Errorf("server run err: %v", err)
	}

}
