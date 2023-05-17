package mysql

import (
	"log"

	"github.com/binbinly/pkg/storage/orm"
	"gorm.io/gorm"

	"gin-chat/pkg/config"
)

// DB 数据库全局变量
var DB *gorm.DB

// NewDB new mysql db
func NewDB() *gorm.DB {
	cfg, err := loadConf("")
	if err != nil {
		log.Fatalf("load orm conf err: %v", err)
	}

	DB = orm.NewMySQL(cfg)
	return DB
}

// NewBasicDB new mysql db
func NewBasicDB(host, user, pwd, name string) *gorm.DB {
	DB = orm.NewBasicMySQL(host, user, pwd, name)
	return DB
}

// loadConf load db config
func loadConf(name string) (*orm.Config, error) {
	if name == "" {
		name = "default"
	}
	v, err := config.LoadWithType("database")
	if err != nil {
		return nil, err
	}

	var cfg orm.Config
	if err = v.UnmarshalKey(name, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
