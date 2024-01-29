package main

import (
	"context"
	maincfg "currency-rate/config"
	"currency-rate/internal/app"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
)

const configPath = "./config/config.yaml"

func main() {

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(yaml.NewEncoder()))))
	if err != nil {
		logrus.Fatal(err)
	}

	if err = cfg.Load(file.NewSource(file.WithPath(configPath))); err != nil {
		logrus.Fatal(err)
	}

	content := &maincfg.Config{}

	if err = cfg.Scan(content); err != nil {
		logrus.Fatal(err)
	}
	content.Port = cfg.Get("port").String("8080")
	content.Address = cfg.Get("address").String("localhost")

	ctx := context.TODO()
	app.Run(ctx, content)
}
