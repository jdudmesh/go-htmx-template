package boot

import (
	"context"
	"fmt"
	"os"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	CurrentDirectory string
	StaticFileDir    string `env:"STATIC_FILE_DIR,default=./static"`
	DatabaseURL      string `env:"DATABASE_URL,default=sqlite://file::memory:?cache=shared"`
	Server           struct {
		Port int `env:"PORT,default=8080"`
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		return nil, err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("getting current directory: %w", err)
	}
	cfg.CurrentDirectory = cwd

	return &cfg, nil
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}
