package postgresql

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	"clean_architecture/logs"

	_ "github.com/lib/pq"
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

		DB_HOST := os.Getenv("PG_DB_HOST")
		DB_USERNAME := os.Getenv("PG_DB_USERNAME")
		DB_PASSWORD := os.Getenv("PG_DB_PASSWORD")
		DB_DATABASE := os.Getenv("PG_DB_DATABASE")
		DB_PORT := os.Getenv("PG_DB_PORT")

		requiredEnvVars := map[string]string{
			"PG_DB_HOST":     DB_HOST,
			"PG_DB_USERNAME": DB_USERNAME,
			"PG_DB_PASSWORD": DB_PASSWORD,
			"PG_DB_DATABASE": DB_DATABASE,
			"PG_DB_PORT":     DB_PORT,
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
			dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
				DB_HOST, DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_PORT)

			db, err = sql.Open("postgres", dsn)
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
