package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Repository содержит все репозитории
type Repository struct {
	User UserRepository
	// Здесь можно добавить другие репозитории
	// Product ProductRepository
	// Order OrderRepository
}

// NewRepository создает новый экземпляр Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		// Инициализация других репозиториев
	}
}

// InitDB инициализирует подключение к PostgreSQL
func InitDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	// Проверка подключения
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logrus.Info("Successfully connected to PostgreSQL database")
	return db, nil
}
