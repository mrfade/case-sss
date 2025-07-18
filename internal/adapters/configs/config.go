package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App   *App
		HTTP  *HTTP
		DB    *DB
		Redis *Redis
		JWT   *JWT
	}

	App struct {
		Env           string
		ApiBaseURL    string
		WebBaseURL    string
		WebBaseOrigin string
	}

	HTTP struct {
		AllowedOrigins string
		Host           string
		Port           string
	}

	DB struct {
		Connection string
		Host       string
		User       string
		Password   string
		DbName     string
		Port       string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
	}

	JWT struct {
		SecretToken string
	}
)

type ConfigManager struct {
	Container *Container
}

func getString(key string) string {
	return os.Getenv(key)
}

func NewConfigManager() (*ConfigManager, error) {
	manager := &ConfigManager{
		Container: &Container{
			App:   &App{},
			HTTP:  &HTTP{},
			DB:    &DB{},
			Redis: &Redis{},
			JWT:   &JWT{},
		},
	}

	err := manager.LoadConfigs()
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (C *ConfigManager) LoadConfigs() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	// Load configs from environment variables
	C.Container.App.Env = getString("APP_ENV")
	C.Container.App.ApiBaseURL = getString("API_BASE_URL")

	// HTTP Configs
	C.Container.HTTP.AllowedOrigins = getString("ALLOWED_ORIGINS")
	C.Container.HTTP.Host = getString("HTTP_HOST")
	C.Container.HTTP.Port = getString("HTTP_PORT")

	// Database Secrets
	C.Container.DB.User = getString("DB_USER")
	C.Container.DB.Password = getString("DB_PASS")
	C.Container.DB.Host = getString("DB_HOST")
	C.Container.DB.Port = getString("DB_PORT")
	C.Container.DB.DbName = getString("DB_NAME")

	// Redis Secrets
	C.Container.Redis.Password = getString("REDIS_PASS")
	C.Container.Redis.Host = getString("REDIS_HOST")
	C.Container.Redis.Port = getString("REDIS_PORT")

	// Auth Secrets
	C.Container.JWT.SecretToken = getString("JWT_SECRET_TOKEN")

	return nil
}

func (manager *ConfigManager) IsDevelopment() bool {
	return manager.Container.App.Env == "DEVELOPMENT"
}

func (manager *ConfigManager) IsProduction() bool {
	return manager.Container.App.Env == "PRODUCTION"
}
