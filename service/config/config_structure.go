package config

type Config struct {
	Project struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
	} `yaml:"project"`

	Server struct {
		Host string `yaml: "host"`
		Port int `yaml:"port"`
	} `yaml:"server"`

	MongoDb struct {
		Host     string `yaml:"host"`
		Port     int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Type     string `yaml:"type"`
	} `yaml:"mongodb"`
}


