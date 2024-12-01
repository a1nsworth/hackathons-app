package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Host string `env:"SERVER_HOST" env-required:"true"`
	Port string `env:"SERVER_PORT" env-required:"true"`
}

func (c ServerConfig) GetURL() string {
	return fmt.Sprintf(":%s", c.Port)
}

type DBConfig struct {
	// Dialect  string `env:"DB_DIALECT" env-required:"true"`
	Port     string `yaml:"DB_HOST" env:"DB_PORT" env-required:"true"`
	Name     string `env:"DB_NAME" env-required:"true"`
	Host     string `env:"DB_HOST" env-required:"true"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

func (db DBConfig) GetDsn() string {
	return "host=" + db.Host + " port=" + db.Port + " user=" + db.User + " password=" + db.Password + " dbname=" + db.Name
}

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

func GetConfig(path string) (config Config) {
	err := cleanenv.ReadConfig(path, &config)
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}
	return
}
