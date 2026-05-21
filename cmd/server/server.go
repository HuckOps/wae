package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"wae/config"
	"wae/db"
	"wae/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger(zap.DebugLevel.String())

	if err := config.LoadConfig("/home/huck/wae/config.yaml"); err != nil {
		log.Fatalf("Can't read config file. pls check.")
	}

	if err := db.InitMysql(); err != nil {
		log.Fatalf("Can't init mysql. pls check.")
	}

	if err := db.InitRedis(); err != nil {
		log.Fatalf("Can't init redis. pls check.")
	}

	r := gin.Default()

	// Start API Server
	go func() {
		if err := r.Run(config.Config.ServerConfig.ApiListenAddr); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
