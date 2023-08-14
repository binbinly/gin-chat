package app

import (
	"log"
	"time"

	"gin-chat/pkg/config"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// DB 数据库全局变量
var DB *gorm.DB

type DBConfig struct {
	Default orm.Config
}

// InitDB init dbs
func InitDB() *gorm.DB {
	var cfg = &DBConfig{}
	if err := loadDBConf(cfg); err != nil {
		log.Fatalf("load db conf err: %v", err)
	}

	DB = orm.NewDB(&cfg.Default)
	return DB
}

// InitBasicDB init basic db
func InitBasicDB(driver, dsn string) *gorm.DB {
	DB = orm.NewDB(&orm.Config{
		Driver: driver,
		Dsn:    dsn,
	})
	return DB
}

// loadDBConf load dbs config
func loadDBConf(cfg *DBConfig) error {
	if err := config.Load("database", cfg, func(v *viper.Viper) {
		v.SetDefault("default", map[string]any{
			"Driver":          "mysql",
			"Host":            "127.0.0.1",
			"Port":            3306,
			"User":            "root",
			"Password":        "root",
			"Database":        "chat",
			"Debug":           true,
			"MaxIdleConn":     10,
			"MaxOpenConn":     100,
			"ConnMaxLifeTime": 100 * time.Second,
		})
		v.BindEnv("default.driver", "CHAT_DB_DRIVER")
		v.BindEnv("default.dsn", "CHAT_DB_DSN")
		v.BindEnv("default.host", "CHAT_DB_HOST")
		v.BindEnv("default.port", "CHAT_DB_PORT")
		v.BindEnv("default.user", "CHAT_DB_USER")
		v.BindEnv("default.password", "CHAT_DB_PASSWORD")
		v.BindEnv("default.database", "CHAT_DB_DATABASE")
	}); err != nil {
		return err
	}

	return nil
}
