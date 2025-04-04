package configs

import (
	rlog "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var (
	gloablConfig GlobalConfig
	once         sync.Once
)

type GlobalConfig struct {
	AppConfig AppConf `yaml:"app" mapstructure:"app"`
	LogConfig LogConf `yaml:"log" mapstructure:"log"`
	DbConfig  DbConf  `yaml:"db" mapstructure:"db"`
}

type AppConf struct {
	AppName string `yaml:"app_name" mapstructure:"app_name"`
	Version string `yaml:"version" mapstructure:"version"`
	Port    int    `yaml:"port" mapstructure:"port"`
	RunMod  string `yaml:"run_mod" mapstructure:"run_mod"`
}

type LogConf struct {
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"`
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`
	SaveDays   uint   `yaml:"save_days" mapstructure:"save_days"`
	Level      string `yaml:"level" mapstructure:"level"`
}

type DbConf struct {
	Host        string `yaml:"host" mapstructure:"host"`
	Port        string `yaml:"port" mapstructure:"port"`
	PassWord    string `yaml:"password" mapstructure:"password"`
	User        string `yaml:"user" mapstructure:"user"`
	DbName      string `yaml:"dbname" mapstructure:"dbname"`
	MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
	MaxIdleTime int    `yaml:"max_idle_time" mapstructure:"max_idle_time"`
}

func GetGlobalConfig() *GlobalConfig {
	once.Do(readConf)
	return &gloablConfig
}

func readConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file err" + err.Error())
	}

	err = viper.Unmarshal(&gloablConfig)
	if err != nil {
		panic("unmarshal config file err")
	}
}

func InitGlobalConfig() {
	config := GetGlobalConfig()
	level, err := log.ParseLevel(config.LogConfig.Level)
	if err != nil {
		panic("parse log level err")
	}
	log.SetFormatter(&logFormatter{
		log.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	})
	log.SetReportCaller(true)
	log.SetLevel(level)
	switch gloablConfig.LogConfig.LogPattern {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "file":
		logger, err := rlog.New(
			config.LogConfig.LogPath+".%Y%m%d",
			rlog.WithRotationTime(time.Hour*24),
			rlog.WithRotationCount(config.LogConfig.SaveDays),
		)
		if err != nil {
			panic("log conf err")
		}
		log.SetOutput(logger)
	default:
		panic("log init err")
	}
}
