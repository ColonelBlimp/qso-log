//go:build windows

package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"qso-log/backend/utils"
)

var (
	errorFormatInit = "db.Init: %w"
)

var dbConn *sql.DB

// Init initializes the database.
func Init(databaseDir, driverName, dataSourceName string) error {
	if dbConn != nil {
		return nil
	}

	err := utils.PathExists(databaseDir)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(databaseDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf(errorFormatInit, err)
		}
	}

	if !errors.Is(err, os.ErrExist) {
		return fmt.Errorf(errorFormatInit, err)
	}

	dbConn, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return fmt.Errorf(errorFormatInit, err)
	}

	err = dbConn.Ping()
	if err != nil {
		return fmt.Errorf(errorFormatInit, err)
	}

	return nil
}
