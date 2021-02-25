package settings

import (
	"fmt"
	"github.com/spf13/viper"
)

func ViperReadEnvPath(path string, fileName string, fileType string) {
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
	var configuration Config
	
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	viper.WriteConfigAs("./final_env.json")
}
