package main

import (
	"github.com/rogierlommers/logrus-redis-hook"
	log "github.com/sirupsen/logrus"
)

func init() {
	hook, err := logredis.NewHook("localhost", 6379, "my_redis_key", "v0", "my_app_name")
	if err == nil {
		log.AddHook(hook)
	} else {
		log.Error(err)
	}
}

func main() {
	// when hook is injected succesfully, logs will be send to redis server
	log.Info("just some info logging...")
}
