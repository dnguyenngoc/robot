package settings

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Project struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
	} `yaml:"project"`

	Server struct {
		Host 	 string `yaml: "host"`
		Port 	 int `yaml:"port"`
	} `yaml:"server"`

	MongoDb struct {
		Host     string `yaml:"host"`
		Port     int `yaml:"port"`
		User 	 string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Type     string `yaml:"type"`
	} `yaml:"mongodb"`
}

func GetEnv() *Config {
	viper.SetConfigName("final_env.json")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	var configuration *Config
	if err := viper.Unmarshal(&configuration); err != nil {
		log.Panicf("error: %s", err)
	}
	return configuration
}