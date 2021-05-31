package infrastructure

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	// migrate v4 database mysql
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//Migrations -> Migration Struct
type Migrations struct {
	logger Logger
	env    Env
}

//NewMigrations -> return new Migrations struct
func NewMigrations(logger Logger, env Env) Migrations {
	return Migrations{
		logger: logger,
		env:    env,
	}
}

//Migrate -> migrates all table
func (m Migrations) Migrate() {
	m.logger.Zap.Info("Migrating schemas...........")

	username := m.env.DBUsername
	password := m.env.DBPassword
	host := m.env.DBHost
	port := m.env.DBPort
	dbname := m.env.DBName
	environment := m.env.Environment

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

	migrations, err := migrate.New("file://migration/", "mysql://"+url)
	if err != nil {
		m.logger.Zap.Error("Error [migration File init] ::", err.Error())
		panic(err)
	}

	m.logger.Zap.Info("---- Running Migration -----")
	err = migrations.Steps(1000)
	if err != nil {
		m.logger.Zap.Error("Error in migration: ", err.Error())
	}
}
