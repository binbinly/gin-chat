package app

import (
	"os"
	"path/filepath"
)

const (
	// EnvLocal 本地环境
	EnvLocal = "local"
	// EnvDev 开发环境
	EnvDev = "dev"
	// EnvTest 测试环境
	EnvTest = "test"
	// EnvProd 生产环境
	EnvProd = "prod"
)

// RootDir 运行根目录
func RootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}

// IsProd 是否生产环境
func IsProd() bool {
	return Conf.Env == EnvProd
}

// IsTest 是否测试环境
func IsTest() bool {
	return Conf.Env == EnvTest
}

// IsDev 是否为开发环境
func IsDev() bool {
	return Conf.Env == EnvDev
}

// IsLocal 是否本地环境
func IsLocal() bool {
	return Conf.Env == EnvLocal
}
