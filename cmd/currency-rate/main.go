package main

import (
	"currency-rate/config"
	"currency-rate/internal/app"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	yamlcfg "gopkg.in/yaml.v3"
)

const configPath = "config/config.yaml"

func main() {

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	cfg := new(config.Config)
	file, err := os.ReadFile(configPath)
	if err != nil {
		l.Error(err)
	}

	if err = yamlcfg.Unmarshal(file, &cfg); err != nil {
		l.Fatal(err)
	}

	l.Fatal(app.Run(cfg))
}
