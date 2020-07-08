package redis

import (
	"gin-api/config/yaml"
	"github.com/gomodule/redigo/redis"
	"log"
)

var Client redis.Conn

func init() {

	config := yaml.Conf().Redis

	Client, _ = redis.Dial("tcp", config.Host+":"+config.Port)

	if config.Password != "" {
		Client.Do("AUTH", config.Password)
	}

	_, err := Client.Do("SELECT", config.Database)

	if err == nil {
		log.Println("Redis Connection Succeeded")
	} else {
		log.Println("Redis Connection Error : " + err.Error())
	}

}
