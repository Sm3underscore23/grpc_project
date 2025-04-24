package config

import (
	"fmt"
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

type dbConfig struct {
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbSSL      string `yaml:"db_sslmode"`
}

type config struct {
	Host     string   `yaml:"host"`
	Port     string   `yaml:"port"`
	DbConfig dbConfig `yaml:"db"`
}

func MainConfigInit(path string) (config, error) {
	var mainConfig config

	if _, err := os.Stat(path); err != nil {
		return config{}, err
	}

	rowConfig, err := os.ReadFile(path)
	if err != nil {
		return config{}, err
	}

	err = yaml.Unmarshal(rowConfig, &mainConfig)
	if err != nil {
		return config{}, err
	}

	return mainConfig, nil
}

func (cfg *config) DbConfigLoad() string {

	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.DbConfig.DbHost,
		cfg.DbConfig.DbPort,
		cfg.DbConfig.DbName,
		cfg.DbConfig.DbUser,
		cfg.DbConfig.DbPassword,
		cfg.DbConfig.DbSSL,
	)
}

func (cfg *config) ServerConfigLoader() string {

	return net.JoinHostPort(cfg.Host, cfg.Port)
}
