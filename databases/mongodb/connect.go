package mongodb

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"microservice/logs"
)

var (
	mongoDBInstance *mongo.Client
	mongoOnce       sync.Once
)

func Connect() (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	mongoOnce.Do(func() {
		maxRetries := 5
		retryCount := 0

		uri := os.Getenv("MONGODB_URI")

		if uri == "" {
			msg := "missing required environment variable: MONGODB_URI"
			err = fmt.Errorf(msg)
			logs.Error(msg)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Set initial timeout
		defer cancel()

		for retryCount < maxRetries {
			serverAPI := options.ServerAPI(options.ServerAPIVersion1)
			opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

			client, err = mongo.Connect(ctx, opts)
			if err != nil {
				msg := fmt.Sprintf("Failed to connect to MongoDB, retrying... (attempt %d/%d): %v", retryCount+1, maxRetries, err)
				logs.Warn(msg)
				retryCount++
				time.Sleep(5 * time.Second)
				continue
			}

			// Ping the database to ensure connection is alive
			err = client.Ping(ctx, nil)
			if err != nil {
				msg := fmt.Sprintf("Failed to ping MongoDB, retrying... (attempt %d/%d): %v", retryCount+1, maxRetries, err)
				logs.Warn(msg)
				retryCount++
				time.Sleep(5 * time.Second)
				continue
			}

			break // connected and ping succeeded
		}

		if retryCount >= maxRetries {
			msg := fmt.Sprintf("failed to connect to MongoDB after %d retries", maxRetries)
			logs.Error(msg)
			err = fmt.Errorf(msg)
			return
		}

		mongoDBInstance = client
		logs.Info("Successfully connected to MongoDB")
	})
	return client, err
}

func Close() error {
	if mongoDBInstance != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := mongoDBInstance.Disconnect(ctx)
		if err != nil {
			return fmt.Errorf("failed to disconnect MongoDB: %w", err)
		}
		logs.Info("MongoDB connection closed gracefully")
	}
	return nil
}
