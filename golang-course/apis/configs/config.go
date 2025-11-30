package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// is possible use var cfg *conf ouside LoadConfig to avoid reloading config every time

// conf holds the application configuration
type conf struct {
	DBDriver      string           `mapstructure:"DB_DRIVER"`
	DBHost        string           `mapstructure:"DB_HOST"`
	DBPort        string           `mapstructure:"DB_PORT"`
	DBUser        string           `mapstructure:"DB_USER"`
	DBPassword    string           `mapstructure:"DB_PASSWORD"`
	DBName        string           `mapstructure:"DB_NAME"`
	WebServerPort string           `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string           `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int              `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth // JWT authentication instance
}

// LoadConfig loads configuration from file and environment variables, use viper package to manage configs
func LoadConfig(path string) (*conf, error) {
	// is possible use var cfg *conf to avoid reloading config every time
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// unmarshal config into conf struct var cfg conf on memory
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	// Initialize JWT authentication
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, err
}
