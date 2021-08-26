package utils

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"postbar/datamodels"
	err2 "postbar/err"
	"strings"
)

func getCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	return path
}

var (
	tmpEnvPah  = getCurrentPath() + "envs/dbs.env"
	ActualPath = "../envs/dbs.env"
)

func init() {
	log.Println("Path", getCurrentPath())
	err := godotenv.Load(tmpEnvPah)
	err2.ReciteErr(&err)
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
		errs := errors.New("error while getting environment variables")
		err2.ReciteErr(&errs)
		return "", "", errs
	}
	return host, port, nil
}
