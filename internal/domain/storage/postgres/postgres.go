package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"orders/internal/config"
	"orders/internal/domain/entity"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(cfg config.Postgres) (*Postgres, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection string: %w", err)
	}

	// Устанавливаем параметры пула
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Проверяем подключение
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.Println("Successfully connected to database")

	return &Postgres{
		Pool: pool,
	}, nil
}

	
func (db *Postgres) SaveUser(user entity.User) error{
	query := "INSERT INTO users(username, password_hash) VALUES($1,$2);"


	_, err := db.Pool.Exec(context.Background(), query, user.Username, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

// GetUserById(id int64) entity.User

// DeleteUserById(id int64) error

// UpdateUserById(id int64) error
