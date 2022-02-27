package config

import "time"

var REDIS_ADDR string
var MONGO_URL string
var DB_NAME string = "test"
var STATIC_CACHE_TIME time.Duration = 30 * time.Second

func Init(conf_type string) {
	switch conf_type {
	case "docker":
		REDIS_ADDR = "redis:6379"
		MONGO_URL = "mongodb://db:27017/?connect=direct"

	case "shell":
		REDIS_ADDR = "localhost:6379"
		MONGO_URL = "mongodb://localhost:27017/?connect=direct"
	}

}
