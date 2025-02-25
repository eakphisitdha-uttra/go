package logs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	logErrorOnce  sync.Once
	logErrorMutex sync.Mutex
	logErrorFile  *os.File
	LogError      *zap.Logger

	logInfoOnce  sync.Once
	logInfoMutex sync.Mutex
	logInfoFile  *os.File
	LogInfo      *zap.Logger

	logDebugOnce  sync.Once
	logDebugMutex sync.Mutex
	logDebugFile  *os.File
	LogDebug      *zap.Logger
)

func InitLoggerError() {
	logErrorOnce.Do(func() {
		LogError = initLogger("STORAGES_FOLDER_ERROR_PATH", &logErrorFile)
	})
	deleteOldLogs("STORAGES_FOLDER_ERROR_PATH", "LOG_ERROR_EXP")
}

func InitLoggerInfo() {
	logInfoOnce.Do(func() {
		LogInfo = initLogger("STORAGES_FOLDER_INFO_PATH", &logInfoFile)
	})
	deleteOldLogs("STORAGES_FOLDER_INFO_PATH", "LOG_INFO_EXP")
}

func InitLoggerDebug() {
	logDebugOnce.Do(func() {
		LogDebug = initLogger("STORAGES_FOLDER_DEBUG_PATH", &logDebugFile)
	})
	deleteOldLogs("STORAGES_FOLDER_DEBUG_PATH", "LOG_DEBUG_EXP")
}

func initLogger(envFolder string, file **os.File) *zap.Logger {
	storagesFolder := os.Getenv(envFolder)
	if storagesFolder == "" {
		panic(fmt.Sprintf("Environment variable %s is not set", envFolder))
	}
	if _, err := os.Stat(storagesFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(storagesFolder, 0755); err != nil {
			panic(err)
		}
	}

	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s/%s.log", storagesFolder, date)
	var err error
	*file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.StacktraceKey = ""
	config.OutputPaths = []string{filename}
	config.ErrorOutputPaths = []string{filename}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func deleteOldLogs(folderEnv, expEnv string) {
	storagesFolder := os.Getenv(folderEnv)
	if storagesFolder == "" {
		fmt.Printf("Environment variable %s is not set\n", folderEnv)
		return
	}

	expDays, err := strconv.Atoi(os.Getenv(expEnv))
	if err != nil {
		fmt.Printf("Invalid %s value, defaulting to 7 days\n", expEnv)
		expDays = 7
	}

	filepath.Walk(storagesFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Failed to access %s: %v\n", path, err)
			return filepath.SkipDir
		}
		if info.Mode().IsRegular() && time.Since(info.ModTime()).Hours() > float64(expDays*24) {
			if err := os.Remove(path); err != nil {
				fmt.Printf("Failed to remove %s: %v\n", path, err)
			}
		}
		return nil
	})
}

func CloseLogError() {
	closeLogger(logErrorFile, LogError)
}

func CloseLogInfo() {
	closeLogger(logInfoFile, LogInfo)
}

func CloseLogDebug() {
	closeLogger(logDebugFile, LogDebug)
}

func closeLogger(file *os.File, logger *zap.Logger) {
	if file != nil {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing log file: %v\n", err)
		}
	}
	if logger != nil {
		_ = logger.Sync()
	}
}

func logMessage(logger *zap.Logger, mutex *sync.Mutex, level string, message interface{}, fields ...zap.Field) {
	if logger == nil {
		fmt.Println("Logger is not initialized")
		return
	}
	mutex.Lock()
	defer mutex.Unlock()

	msg := fmt.Sprintf("%v", message)
	switch level {
	case "info":
		logger.Info(msg, fields...)
	case "debug":
		logger.Debug(msg, fields...)
	case "warn":
		logger.Warn(msg, fields...)
	case "error":
		logger.Error(msg, fields...)
	}
}

func Info(message interface{}, fields ...zap.Field) {
	logMessage(LogInfo, &logInfoMutex, "info", message, fields...)
}

func Debug(message interface{}, fields ...zap.Field) {
	logMessage(LogDebug, &logDebugMutex, "debug", message, fields...)
}

func Warn(message interface{}, fields ...zap.Field) {
	logMessage(LogError, &logErrorMutex, "warn", message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	logMessage(LogError, &logErrorMutex, "error", message, fields...)
}
