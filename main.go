package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"

	"dgg/cron"
	"dgg/module"
	"dgg/routers"
	"dgg/websocket"

	"os"
	"strconv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//logWriterFormat := "./var/log/gin.%Y%m%d.log"
	//logWriter, err := rotatelogs.New(
	//	logWriterFormat,
	//	rotatelogs.WithLinkName("var/log/access_log.log"),
	//	rotatelogs.WithMaxAge(24 * time.Hour),
	//	rotatelogs.WithRotationTime(time.Hour),
	//)
	//if err != nil {
	//	log.Printf("failed to create rotatelogs: %s", err)
	//	return
	//}
	//gin.DefaultWriter = logWriter

	engine := gin.Default()

	routers.InitRouter(engine)
	module.ConnectDB()

	bool, _ := strconv.ParseBool(os.Getenv("Redis")) //是否开启redis
	if bool {
		//判断是否启动Redis
		module.InitRedis()
	}

	wbool, _ := strconv.ParseBool(os.Getenv("WebSocket")) //是否开启websocket
	if wbool {
		go websocket.StartWebsocket(os.Getenv("WebSocketaddrPort"))
	}

	go cron.InitCron()                         //启动定时器
	engine.Run(":" + os.Getenv("ServicePort")) // listen and serve on 0.0.0.0:8080
}
