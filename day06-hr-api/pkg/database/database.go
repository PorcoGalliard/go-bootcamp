package database

import (
	"fmt"
	configs "hrapi/internal/config"
	"hrapi/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

// InitDB initializes and returns Database instance (bukan set global).
// Panggil ini di main(), lalu inject ke dependencies.
func InitDB(cfg *configs.Config) (*Database, error) {
	dsn := generateDSN(cfg.Database)
	log.Printf("Connecting to database: %s@%s:%s/%s",
		cfg.Database.User,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)
	// Configure GORM logger based on environment
	gormConfig := &gorm.Config{}
	if cfg.Environment == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}
	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	if err := db.Exec("SET search_path TO hr").Error; err != nil {
		return nil, fmt.Errorf("Failed to set search path: %v ", err)
	}

	log.Printf("✅ Database connected successfully!")
	// Return embedded struct
	return &Database{DB: db}, nil
}

// generateDSN generates PostgreSQL connection string
func generateDSN(dbConfig configs.DatabaseConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s search_path=hr",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
		dbConfig.SSLMode,
		dbConfig.TimeZone)
}

// GetDB returns the database instance
func GetDB(db *Database) *gorm.DB {
	if db.DB == nil {
		log.Fatal("Database not initialized. Call InitDB first.")
	}
	return db.DB
}

// CloseDB sekarang ambil param Database (atau *gorm.DB jika embed tidak dipakai).
func CloseDB(db *Database) error {
	if db != nil {
		sqlDB, err := db.DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// AutoMigrate will execute given models, execute for all tables
func AutoMigrate(db *Database, models ...any) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	log.Printf("Running auto migration for %d models...", len(models))
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to auto migrate: %v", err)
	}
	log.Printf("✅ Auto migration completed successfully!")
	return nil
}
func InitAutoMigrate(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS hr")
	/* err := db.AutoMigrate(
	&models.Region{},
	)
	if err != nil {
	log.Fatal("Failed to migrate database:", err)
	} */
	// 1. Migrate Region dulu
	if err := db.AutoMigrate(&models.Region{}); err != nil {
		log.Fatal("Error migrating Region:", err)
	}
	// 1. Baru Migrate Country
	if err := db.AutoMigrate(&models.Country{}); err != nil {
		log.Fatal("Error migrating Country:", err)
	}
}
