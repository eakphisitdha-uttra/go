package mssql

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	"microservice/logs"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	sqlDBOnce     sync.Once
	sqlDBInstance *sql.DB
)

func Connect() (*sql.DB, error) {
	var err error
	var db *sql.DB
	sqlDBOnce.Do(func() {
		maxRetries := 5
		retryCount := 0

		DB_HOST := os.Getenv("MS_DB_HOST")
		DB_USERNAME := os.Getenv("MS_DB_USERNAME")
		DB_PASSWORD := os.Getenv("MS_DB_PASSWORD")
		DB_DATABASE := os.Getenv("MS_DB_DATABASE")
		DB_PORT := os.Getenv("MS_DB_PORT")

		requiredEnvVars := map[string]string{
			"MS_DB_HOST":     DB_HOST,
			"MS_DB_USERNAME": DB_USERNAME,
			"MS_DB_PASSWORD": DB_PASSWORD,
			"MS_DB_DATABASE": DB_DATABASE,
			"MS_DB_PORT":     DB_PORT,
		}

		missingVars := make([]string, 0)
		for k, v := range requiredEnvVars {
			if v == "" {
				missingVars = append(missingVars, k)
			}
		}

		if len(missingVars) > 0 {
			msg := fmt.Sprintf("missing required environment variables: %v", missingVars)
			err = fmt.Errorf(msg)
			logs.Error(msg)
			return
		}

		for retryCount < maxRetries {
			dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
				DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)

			db, err = sql.Open("sqlserver", dsn)
			if err != nil {
				msg := fmt.Sprintf("Failed to open database connection, retrying... (attempt %d/%d): %v", retryCount+1, maxRetries, err)
				logs.Warn(msg)
				retryCount++
				time.Sleep(5 * time.Second)
				continue
			}

			// Ping the database to ensure connection is alive
			err = db.Ping()
			if err != nil {
				msg := fmt.Sprintf("Failed to ping database, retrying... (attempt %d/%d): %v", retryCount+1, maxRetries, err)
				logs.Warn(msg)
				retryCount++
				time.Sleep(5 * time.Second)
				if err := db.Close(); err != nil {
					logs.Error(fmt.Sprintf("Error closing db connection: %v", err))
				}
				continue
			}

			break //connected and ping succeeded
		}

		if retryCount >= maxRetries {
			msg := fmt.Sprintf("failed to connect to database after %d retries", maxRetries)
			logs.Error(msg)
			err = fmt.Errorf(msg)
			return
		}

		sqlDBInstance = db
		logs.Info(fmt.Sprintf("Successfully connected to Database: host=%s, dbname=%s", DB_HOST, DB_DATABASE))
	})
	return db, err
}

func Close() error {
	if sqlDBInstance != nil {
		if err := sqlDBInstance.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %w", err)
		}
		logs.Info("Database connection closed gracefully")
	}
	return nil
}
