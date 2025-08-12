package databases

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"github.com/adibhauzan/crons/configs"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitSqlServerConnection() (*gorm.DB, error) {
	DbUser := configs.MSSQLDBUser
	DbPassword := configs.MSSQLSAPassword
	DbHost := configs.MSSQLDBHost
	DbPort := configs.MSSQLDBPort
	DbName := configs.MSSQLDBName

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		url.QueryEscape(DbUser), url.QueryEscape(DbPassword),
		url.QueryEscape(DbHost),
		DbPort,
		url.QueryEscape(DbName),
	)

	var level logger.LogLevel
	switch configs.MODE {
	case "production":
		level = logger.Error
	case "staging":
		level = logger.Warn
	case "development":
		level = logger.Info
	case "local":
		level = logger.Info
	default:
		level = logger.Info
	}

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  level,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		FullSaveAssociations: true,
		PrepareStmt:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQL Server: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Minute * 1)

	return db, nil
}
