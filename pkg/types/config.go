package types

import "fmt"

// EnvConfig содержит настройки из переменных окружения
type EnvConfig struct {
	ApiPort string
}

// YAMLConfig содержит настройки из YAML файла
type YAMLConfig struct {
	Configuration struct {
		Backend struct {
			ApiPort string `yaml:"api_port"`
		} `yaml:"backend"`
		Database struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
			SSLMode  string `yaml:"sslmode"`
		} `yaml:"database"`
	} `yaml:"configuration"`
}

// GetDatabaseDSN возвращает строку подключения к PostgreSQL
func (c *YAMLConfig) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Configuration.Database.Host,
		c.Configuration.Database.Port,
		c.Configuration.Database.User,
		c.Configuration.Database.Password,
		c.Configuration.Database.DBName,
		c.Configuration.Database.SSLMode,
	)
}