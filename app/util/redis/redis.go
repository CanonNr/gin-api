package redis

import (
	"gin-api/app/util/yaml"
	"github.com/gomodule/redigo/redis"
	"log"
)

func Client() (redis.Conn, error) {

	config := yaml.Conf().Redis

	Client, _ := redis.Dial("tcp", config.Host+":"+config.Port)

	Client.Do("AUTH", config.Password)

	_, err := Client.Do("SELECT", config.Database)

	if err == nil {
		log.Println("Redis Connection Succeeded")
	} else {
		log.Fatalln("Redis Connection Error : " + err.Error())
	}
	return Client, err
}
