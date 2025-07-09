package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"user-doc-collaborative-editor/pkg/logger"
)

var DB *sql.DB

func InitMySQL(dsn string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"action": "InitMySQL",
			"error":  err.Error(),
		}).Error("Failed to open MySQL database")
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		logger.WithFields(logrus.Fields{
			"action": "InitMySQL",
			"error":  err.Error(),
		}).Error("Failed to ping MySQL database")
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	logger.WithFields(logrus.Fields{
		"action": "InitMySQL",
	}).Info("Successfully connected to the database")
	log.Println("Successfully connected to the database")
	return DB, nil
}
