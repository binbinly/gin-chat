package config

const (
	fileTypeYaml = "yaml"
	fileTypeJson = "json"
	fileTypeToml = "toml"
)

type Option func(*Config)

func WithFileTypeYaml() Option {
	return func(c *Config) {
		c.fileType = fileTypeYaml
	}
}

func WithFileTypeJson() Option {
	return func(c *Config) {
		c.fileType = fileTypeJson
	}
}

func WithFileTypeToml() Option {
	return func(c *Config) {
		c.fileType = fileTypeToml
	}
}

func WithEnv(name string) Option {
	return func(c *Config) {
		c.env = name
	}
}

func WithConfigDir(dir string) Option {
	return func(c *Config) {
		c.configDir = dir
	}
}

func WithEnvPrefix(prefix string) Option {
	return func(c *Config) {
		c.envPrefix = prefix
	}
}
