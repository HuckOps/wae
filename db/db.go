package db

import (
	"context"
	"wae/config"
	"wae/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitMysql() error {
	db, err := gorm.Open(mysql.Open(config.Config.ServerConfig.MysqlDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	Db = db
	err = db.AutoMigrate(&model.Service{})
	return err
}

var RedisClient *redis.Client

func InitRedis(ctx context.Context) error {
	opt, err := redis.ParseURL(config.Config.ServerConfig.RedisDSN)
	if err != nil {
		return err
	}
	redisClient := redis.NewClient(opt)
	RedisClient = redisClient
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}
