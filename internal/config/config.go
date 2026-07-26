package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App             AppConfig             `mapstructure:"app"`
	Server          ServerConfig          `mapstructure:"server"`
	CORS            CORSConfig            `mapstructure:"cors"`
	HTTPClient      HTTPClientConfig      `mapstructure:"http_client"`
	Auth            AuthConfig            `mapstructure:"auth"`
	DummyJSON       DummyJSONConfig       `mapstructure:"dummyjson"`
	JSONPlaceholder JSONPlaceholderConfig `mapstructure:"jsonplaceholder"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
}

type ServerConfig struct {
	Port      int `mapstructure:"port"`
	BodyLimit int `mapstructure:"body_limit"`
}

type CORSConfig struct {
	AllowOrigins string `mapstructure:"allow_origins"`
	AllowMethods string `mapstructure:"allow_methods"`
	AllowHeaders string `mapstructure:"allow_headers"`
}

type HTTPClientConfig struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

type AuthConfig struct {
	CookieName string `mapstructure:"cookie_name"`
	JWTSecret  string `mapstructure:"jwt_secret"`
}

type DummyJSONConfig struct {
	BaseURL string `mapstructure:"base_url"`
}

type JSONPlaceholderConfig struct {
	BaseURL string `mapstructure:"base_url"`
}

func Load() (*Config, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("usage: go run . <dev|uat|production>")
	}

	env := os.Args[1]

	viper.SetConfigName(env)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}
