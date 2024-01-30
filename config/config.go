package config

type Config struct {
	ServerAddress string `yaml:"serverAddress"` //TODO struct
	Address       string `yaml:"address"`       //TODO struct
	DB            `yaml:"database"`
}
type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}
