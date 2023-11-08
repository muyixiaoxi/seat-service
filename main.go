package main

import (
	"seat-service/initialization"
	"seat-service/router"
)

func main() {
	//初始化viper
	initialization.InitViper()
	//初始化zap
	initialization.InitZap()
	//初始化mysql
	initialization.InitMysql()
	//初始化redis
	initialization.InitRedis()

	router.Router()
}
