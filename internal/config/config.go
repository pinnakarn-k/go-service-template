package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App         AppConfig         `mapstructure:"app"`
	Server      ServerConfig      `mapstructure:"server"`
	CORS        CORSConfig        `mapstructure:"cors"`
	HTTPClient  HTTPClientConfig  `mapstructure:"http_client"`
	Auth        AuthConfig        `mapstructure:"auth"`
	Integration IntegrationConfig `mapstructure:"integration"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
}

type ServerConfig struct {
	Port        int `mapstructure:"port"`
	BodyLimitMB int `mapstructure:"body_limit_mb"`
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

type IntegrationConfig struct {
	GetProducts IntegrationConfigItem `mapstructure:"get_products"`
	GetCarts    IntegrationConfigItem `mapstructure:"get_carts"`
	GetPosts    IntegrationConfigItem `mapstructure:"get_posts"`
	GetPostByID IntegrationConfigItem `mapstructure:"get_post_by_id"`
}

type IntegrationConfigItem struct {
	URL         string `mapstructure:"url"`
	Application string `mapstructure:"application"`
	Requester   string `mapstructure:"requester"`
	Key         string `mapstructure:"key"`
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
