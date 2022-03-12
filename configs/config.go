package configs

import (
	"errors"

	"github.com/spf13/viper"
)

//Config ...
type Config struct {
	Environment               string `json:"environment"`
	HTTPPort                  string `json:"http_port"`
	LogLevel                  string `json:"log_level"`
	CasbinConfigPath          string
	MiddlewareRolesPath       string
	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
	PostgresDatabase          string
	PostgresHost              string
	PostgresPort              uint16
	PostgresUser              string
	PostgresPassword          string
	RedisHost 				  string
	RedisPort                 int
	SendgridEmail			  string
}

//Load loads the configuration from environment variables
func Load() Config {
	var config Config

	v := viper.New()
	v.AutomaticEnv()

	v.SetDefault("ENVIRONMENT", "development")
	v.SetDefault("LOG_LEVEL", "debug")
	v.SetDefault("HTTP_PORT", ":8000")
	v.SetDefault("CASBIN_CONFIG_PATH", "./configs/rbac_model.conf")
	v.SetDefault("MIDDLEWARE_ROLES_PATH", "./configs/models.csv")
	v.SetDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)
	v.SetDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)
	v.SetDefault("POSTGRES_HOST", "localhost")
	v.SetDefault("POSTGRES_PORT", 5432)
	v.SetDefault("POSTGRES_USER", "muhammad")
	v.SetDefault("POSTGRES_PASSWORD", "12345")
	v.SetDefault("POSTGRES_DB", "mono_db")
	v.SetDefault("REDIS_HOST","localhost")
	v.SetDefault("REDIS_PORT", 6379)

	config.Environment = v.GetString("ENVIRONMENT")
	config.HTTPPort = v.GetString("HTTP_PORT")
	config.LogLevel = v.GetString("LOG_LEVEL")
	config.CasbinConfigPath = v.GetString("CASBIN_CONFIG_PATH")
	config.MiddlewareRolesPath = v.GetString("MIDDLEWARE_ROLES_PATH")
	config.JWTSecretKey = v.GetString("JWT_SECRET_KEY")
	config.JWTRefreshKey = v.GetString("JWT_REFRESH_KEY")
	config.JWTSecretKeyExpireMinutes = v.GetInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	config.JWTRefreshKeyExpireHours = v.GetInt("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")
	config.PostgresDatabase = v.GetString("POSTGRES_DB")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPassword = v.GetString("POSTGRES_PASSWORD")
	config.PostgresHost = v.GetString("POSTGRES_HOST")
	config.PostgresPort = (uint16)(v.GetUint("POSTGRES_PORT"))
	config.RedisHost = v.GetString("REDIS_HOST")
	config.RedisPort = v.GetInt("REDIS_PORT")
	return config
}

//Validate validates the configuration
func (c *Config) Validate() error {
	if c.HTTPPort == "" {
		return errors.New("http_port required")
	}
	return nil
}
