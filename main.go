package main

import (
	"github.com/joho/godotenv"
	"log"
	"github.com/gin-gonic/gin"
	"lidongheDemo/routers"

	"lidongheDemo/models"
	"os"
	"strconv"
	"lidongheDemo/cron"
	"lidongheDemo/websocket"
	"github.com/lestrrat/go-file-rotatelogs"
	"time"
)

func main() {


	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	logWriterFormat := "./var/log/gin.%Y%m%d.log"
	logWriter, err := rotatelogs.New(
		logWriterFormat,
		rotatelogs.WithLinkName("var/log/access_log.log"),
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	gin.DefaultWriter = logWriter


	engine := gin.Default()


	routers.InitRouter(engine)
	models.ConnectDB()
	bool,_:=strconv.ParseBool(os.Getenv("Redis"))//是否开启redis
	if bool  {
		//判断是否启动Redis
		models.InitRedis()
	}
	wbool,_:=strconv.ParseBool(os.Getenv("WebSocket"))//是否开启websocket
	if wbool {
		go websocket.StartWebsocket(os.Getenv("WebSocketaddrPort"))
	}

	go cron.InitCron() //启动定时器
	engine.Run(":"+os.Getenv("ServicePort")) // listen and serve on 0.0.0.0:8080
}