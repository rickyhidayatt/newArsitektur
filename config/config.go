package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var configEnv map[string]string

func init() {
	content, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &configEnv)
	if err != nil {
		log.Fatal("invalid config file")
	}
}

var defaultConfig = map[string]string{
	"HTTP_PORT":       ":8080",
	"DATABASE_HOST":   "localhost",
	"DATABASE_PORT":   "5432",
	"DATABASE_NAME":   "xpora_users",
	"DATABASE_USER":   "postgres",
	"DATABASE_PASS":   "password",
	"DATABASE_SSL":    "disable",
	"TEMP_UPLOAD_DIR": "/temp",
}

func GetEnv(key string) string {
	value, ok := configEnv[key]
	if !ok {
		defaultValue, ok := defaultConfig[key]
		if !ok {
			return ""
		}
		return defaultValue
	}
	return value
}
