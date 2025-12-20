package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type EnvConfig struct {
	ApiPort string
}

type YAMLConfig struct {
	Configuration struct {
		Backend struct {
			ApiPort string `yaml:"api_port"`
		} `yaml:"backend"`
	} `yaml:"configuration"`
}

//func InitEnvConfig() *EnvConfig {
//	cfg := &EnvConfig{
//		ApiPort: getEnv("API_PORT", "8080"),
//	}
//	return cfg
//}

func InitYAMLConfig() *YAMLConfig {
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg *YAMLConfig

	if err := yaml.Unmarshal(f, &cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)

	return cfg
}

//func getEnv(key, defaultValue string) string {
//	if value := os.Getenv(key); value != "" {
//		return value
//	}
//	return defaultValue
//}
