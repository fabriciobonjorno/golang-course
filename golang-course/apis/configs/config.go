package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

// conf holds the application configuration
type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  int
	TokenAuth     *jwtauth.JWTAuth // JWT authentication instance
}

// LoadConfig loads configuration from file and environment variables, use viper package to manage configs
func LoadConfig(path string) (*conf, error) {
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

	return cfg, nil
}
