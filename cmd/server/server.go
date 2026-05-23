package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"wae/config"
	"wae/db"
	"wae/pkg/logger"
	"wae/pkg/oidc"
	"wae/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	var configPath string = "./config.yaml"

	flag.StringVar(&configPath, "c", "./config.yaml", "Server config yaml path(default: ./config.yaml).")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger.InitLogger(zap.DebugLevel.String())

	if err := config.LoadConfig(configPath); err != nil {
		logger.Logger.Fatal("Can't read config file. pls check.")
	}

	if err := db.InitMysql(); err != nil {
		logger.Logger.Fatal("Can't init mysql. pls check.")
	}

	if err := db.InitRedis(ctx); err != nil {
		logger.Logger.Fatal("Can't init redis. pls check.")
	}

	if err := oidc.InitOIDC(ctx); err != nil {
		logger.Logger.Fatal("Can't init oidc. pls check.")
	}

	r := gin.Default()

	router.RegisterRouter(r)

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
