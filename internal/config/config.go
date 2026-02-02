package config

import (
	"fmt"
	"log"
	"os"

	"github.com/potterbl/story-backend/pkg/types"
	"gopkg.in/yaml.v3"
)

//func InitEnvConfig() *types.EnvConfig {
//	cfg := &types.EnvConfig{
//		ApiPort: getEnv("API_PORT", "8080"),
//	}
//	return cfg
//}

func InitYAMLConfig() *types.YAMLConfig {
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg *types.YAMLConfig

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
