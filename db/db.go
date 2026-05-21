package db

import (
	"context"
	"wae/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql() error {
	db, err := gorm.Open(mysql.Open(config.Config.ServerConfig.MysqlDSN), &gorm.Config{})
	Db = db
	return err
}

var RedisClient *redis.Client

func InitRedis() error {
	opt, err := redis.ParseURL(config.Config.ServerConfig.RedisDSN)
	if err != nil {
		return err
	}
	redisClient := redis.NewClient(opt)
	RedisClient = redisClient
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}
