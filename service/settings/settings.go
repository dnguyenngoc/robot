/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package settings

import (
	"log"
	"github.com/spf13/viper"
)

// Global variable
var Env Config


// define config
type Config struct {
	Project struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"project"`

	Server struct {
		Host 	 string `yaml: "host"`
		Port 	 int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int `yaml:"port"`
		User 	 string `yaml:"user"`
		Pass 	 string `yaml:"pass"`
		Name 	 string `yaml:"name"`
		Type     string `yaml:"type"`
	} `yaml:"database"`
}

func InitialViperWriteCofig(path string, fileName string, fileType string) {
	/*
        It is a basic viber to make application can make final_file environment for .yaml, .toml ...
        Just only run one time when starting server.
    */

	// Set the file name of the configurations file
	viper.SetConfigName(fileName)

	// Set the path to look for the configurations file
	viper.AddConfigPath(path)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Set config file type (ex: yml)
	viper.SetConfigType(fileType)

	// Handle read config failed
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Unable read Flie Config, %v", err)
	}

	// handle write file to structure config
	if err := viper.Unmarshal(&Env); err != nil {
		log.Panicf("Unable to decode into struct, %v", err)
	}
	return
}
