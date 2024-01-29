package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	*gorm.DB
}

// New provides PostgresDB struct init
func New(host, port, name, user, password string) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	// product :=domain.Product{}
	// db.AutoMigrate(product)

	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) Close() {
	sqlDB, err := p.DB.DB()
	if err != nil {
		fmt.Printf("Error getting underlying DB: %v\n", err)
		return
	}

	sqlDB.Close()
}

func (p *PostgresDB) Error(err error) error {
	return err
}

func (p *PostgresDB) ErrSQLBuild(err error, message string) error {
	return fmt.Errorf("error during sql build, %s: %w", message, err)
}
