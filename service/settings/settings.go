/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package settings

import (
	"log"
	"github.com/spf13/viper"
	"fmt"
)

type Config struct {
	Project struct {
		Name        string `json:"name"`
		Environment string `json:"environment"`
	} `yaml:"project"`

	Server struct {
		Host 	 string `json: "host"`
		Port 	 int `json:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `json:"host"`
		Port     int `json:"port"`
		User 	 string `json:"user"`
		Pass string `json:"pass"`
		Database string `json:"database"`
		Type     string `json:"type"`
	} `yaml:"database"`
}


func GetEnv() (config Config) {
	/* 
		function load final_env inside variable Env
	*/
	viper.AddConfigPath(".")
	viper.SetConfigName("final_env")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
    if err != nil {
		log.Println("Error when read config File!")

    }
	viper.Unmarshal(&config)
    return
}


func InitialViperWriteCofig(path string, fileName string, fileType string) {
	/*
        It is a basic viber to make application can make final_file environment for .yaml, .toml ...
        Just only run one time when starting server.
    */

	//  Config structure
	var conf Config

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
		fmt.Printf("Error reading config file, %s", err)
	}

	// handle write file to structure config
	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	viper.WriteConfigAs("./final_env.json")
}
