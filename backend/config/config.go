package config

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type LLMConfig struct {
	BaseUrl string
	ApiKey  string
}

var DB *gorm.DB

func NewLLMConfig(baseUrl, apiKey string) *LLMConfig {
	return &LLMConfig{
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
	}
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("C:\\Users\\nju\\Desktop\\project\\GoProject\\mcpChat\\backend\\config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	log.Println(viper.GetString("llm.url"))
}

func InitDatabase() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.user")+":"+viper.GetString("mysql.password")+"@tcp(127.0.0.1:3306)/"+viper.GetString("mysql.database")+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	log.Println("db init success")
}
