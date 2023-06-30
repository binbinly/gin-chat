package mysql

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

var cfg = &Config{}

type Config struct {
	Default orm.Config
}

// NewDB new mysql db
func NewDB() *gorm.DB {
	if err := loadConf(); err != nil {
		log.Fatalf("load orm conf err: %v", err)
	}

	DB = orm.NewMySQL(&cfg.Default)
	return DB
}

// NewBasicDB new mysql db
func NewBasicDB(host, user, pwd, name string) *gorm.DB {
	DB = orm.NewBasicMySQL(host, user, pwd, name)
	return DB
}

// loadConf load db config
func loadConf() error {
	if err := config.Load("database", cfg, func(v *viper.Viper) {
		v.SetDefault("default", map[string]any{
			"Addr":            "127.0.0.1:3306",
			"User":            "root",
			"Password":        "root",
			"Database":        "chat",
			"Debug":           true,
			"MaxIdleConn":     10,
			"MaxOpenConn":     100,
			"ConnMaxLifeTime": 100 * time.Second,
		})
		v.BindEnv("default.addr", "CHAT_MYSQL_ADDR")
		v.BindEnv("default.user", "CHAT_MYSQL_USER")
		v.BindEnv("default.password", "CHAT_MYSQL_PASSWORD")
		v.BindEnv("default.database", "CHAT_MYSQL_DATABASE")
	}); err != nil {
		return err
	}

	return nil
}
