package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database modal
type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env Env, zapLogger Logger) Database {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName
	environment := env.Environment

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	if environment != "local" {
		url = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			dbname,
		)
	}
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		zapLogger.Zap.Info("Url: ", url)
		zapLogger.Zap.Panic(err)
	}

	zapLogger.Zap.Info("Database connection established")

	return Database{
		DB: db,
	}
}
