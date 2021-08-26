package utils

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"postbar/datamodels"
	err2 "postbar/err"
)

func init() {
	err := godotenv.Load("../envs/dbs.env")
	err2.Reciteerr(&err)
}

func GetMySqlENV() *datamodels.MySQLENV {
	return &datamodels.MySQLENV{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
	}
}

func GetMongoENV() (string, string, error) {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	if len(host) == 0 || len(port) == 0 {
		errs := errors.New("error while getting environment")
		err2.Reciteerr(&errs)
		return "", "", errs
	}
	return host, port, nil
}
