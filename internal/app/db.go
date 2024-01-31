package app

import (
	"currency-rate/config"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

func NewConnectDB(cfg config.DB) *sql.DB {
	dbConfig := &cfg
	dbClient, err := ConnectPG(dbConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	if err = dbClient.Ping(); err != nil {
		logrus.Fatal(err)
	}
	return dbClient
}

func ConnectPG(c *config.DB) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.Username, c.Password, c.Host, c.Port, c.DBName))
	if err != nil {
		return nil, err
	}
	return db, err
}
