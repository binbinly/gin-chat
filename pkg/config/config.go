package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var conf *Config

// Config conf struct.
type Config struct {
	env       string
	configDir string
	fileType  string //yaml, json, toml, default is yaml
	val       map[string]*viper.Viper
	mu        sync.Mutex
}

// New create a config instance.
func New(cfgDir string, opts ...Option) *Config {
	// must set config dir
	if cfgDir == "" {
		panic("config dir is not set")
	}
	c := Config{
		configDir: cfgDir,
		fileType:  fileTypeYaml,
		val:       make(map[string]*viper.Viper),
	}
	for _, opt := range opts {
		opt(&c)
	}

	conf = &c

	return &c
}

// Load alias for config func.
func Load(filename string, val any) error { return conf.Load(filename, val) }

// Load scan data to struct.
func (c *Config) Load(filename string, val any) error {
	v, err := c.LoadWithType(filename)
	if err != nil {
		return err
	}

	if err = v.Unmarshal(&val); err != nil {
		return err
	}
	return nil
}

// LoadWithType load conf by file type.
func LoadWithType(filename string) (*viper.Viper, error) {
	return conf.LoadWithType(filename)
}

// LoadWithType load conf by file type.
func (c *Config) LoadWithType(filename string) (v *viper.Viper, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.val[filename]
	if ok {
		return v, nil
	}

	v, err = c.load(filename)
	if err != nil {
		return nil, err
	}
	c.val[filename] = v
	return v, nil
}

// Load file.
func (c *Config) load(filename string) (*viper.Viper, error) {
	env := GetEnv("APP_ENV", "")
	if c.env != "" {
		env = c.env
	}
	path := filepath.Join(c.configDir, env)

	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.SetConfigType(c.fileType)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	log.Println("Using config file:", v.ConfigFileUsed())

	return v, nil
}

// GetEnv get value from env.
func GetEnv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}
