package main

import (
	"time"

	"github.com/chiachan163/cc-orm/example/logic/model"
	"github.com/chiachan163/cc-orm/model/mongo"
	"github.com/chiachan163/cc-orm/model/mysql"
	"github.com/chiachan163/cc-orm/model/redis"
	"github.com/henrylee2cn/cfgo"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/goutil"
)

type config struct {
	Mysql       mysql.Config  `yaml:"mysql"`
	Mongo       mongo.Config  `yaml:"mongo"`
	Redis       redis.Config  `yaml:"redis"`
	CacheExpire time.Duration `yaml:"cache_expire"`
	LogLevel    string        `yaml:"log_level"`
}

func (c *config) Reload(bind cfgo.BindFunc) error {
	err := bind()
	if err != nil {
		return err
	}
	if c.CacheExpire == 0 {
		c.CacheExpire = time.Hour * 24
	}
	if len(c.LogLevel) == 0 {
		c.LogLevel = "TRACE"
	}
	erpc.SetLoggerLevel(c.LogLevel)
	var (
		mysqlConfig *mysql.Config
		mongoConfig *mongo.Config
		redisConfig = &c.Redis
	)
	if len(c.Mysql.Host) > 0 {
		mysqlConfig = &c.Mysql
	}
	if len(c.Mongo.Addrs) > 0 {
		mongoConfig = &c.Mongo
	}
	err = model.Init(mysqlConfig, mongoConfig, redisConfig, c.CacheExpire)
	if err != nil {
		erpc.Errorf("%v", err)
	}
	return nil
}

var cfg = &config{
	Redis:       *redis.NewConfig(),
	CacheExpire: time.Hour * 24,
	LogLevel:    "TRACE",
}

func init() {
	goutil.WritePidFile()
}
