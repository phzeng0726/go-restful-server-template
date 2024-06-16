package logger

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

type LoggerManager interface {
	InitLogger() error
	GetLogger() *zap.Logger
	InfoWithElapsedTime(ctx context.Context, action string, startTime time.Time, optionalFields ...map[string]interface{})
	ErrorWithElapsedTime(ctx context.Context, action string, startTime time.Time, err error, optionalFields ...map[string]interface{})
}

type Manager struct {
	AppName       string
	Env           string
	LogFolderPath string
}

func NewManager(appName, env, logFolderPath string) (*Manager, error) {
	return &Manager{
		AppName:       appName,
		Env:           env,
		LogFolderPath: logFolderPath,
	}, nil
}

func (m *Manager) generateLogFilePath() string {

	envMap := map[string]string{
		"development":        "dev",
		"docker-development": "docker_dev",
		"test":               "test",
		"production":         "prod",
	}
	env, ok := envMap[m.Env]
	if !ok {
		env = ""
	}

	date := time.Now().Format("20060102")
	logFileName := fmt.Sprintf("%s_%s_%s.log", m.AppName, env, date)
	logFilePath := fmt.Sprintf("%s/%s", m.LogFolderPath, logFileName)

	// 資料夾不在的時候會自己創
	if err := os.MkdirAll(m.LogFolderPath, 0755); err != nil {
		fmt.Printf("failed to create log folder: %v\n", err)
	}

	return logFilePath
}

func (m *Manager) InitLogger() error {
	// 設定 Encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 設定 Writers
	logFilePath := m.generateLogFilePath()

	file, _ := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	consoleSyncer := zapcore.AddSync(os.Stdout)
	fileSyncer := zapcore.AddSync(file)

	// 設定 Core
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // 使用JSON
			consoleSyncer,                         // 輸出到console
			zapcore.InfoLevel,                     // 日誌級別
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // 使用JSON
			fileSyncer,                            // 輸出到檔案
			zapcore.InfoLevel,                     // 日誌級別
		),
	)

	logger := zap.New(core)
	logger.Info("Logger initialized", zap.String("log_file", logFilePath))

	globalLogger = logger
	return nil
}

func (m *Manager) GetLogger() *zap.Logger {
	return globalLogger
}

// 用來記錄log用
func (m *Manager) getUserIdByCtx(ctx context.Context) (string, error) {
	userId := ctx.Value("userId")

	userIdStr, ok := userId.(string)
	if !ok && userId != nil {
		return "", errors.New("failed to convert value to string")
	}

	return userIdStr, nil
}

func (m *Manager) logWithElapsedTime(ctx context.Context, startTime time.Time, optionalFields ...map[string]interface{}) []zapcore.Field {
	fields := make(map[string]interface{})
	if len(optionalFields) > 0 {
		for k, v := range optionalFields[0] {
			fields[k] = v
		}
	}

	requester, err := m.getUserIdByCtx(ctx)
	if err != nil {
		globalLogger.Error(err.Error())
	}

	fields["requester"] = requester
	fields["elapsed_time"] = time.Since(startTime)
	var zapFields []zap.Field
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	return zapFields
}

func (m *Manager) InfoWithElapsedTime(ctx context.Context, action string, startTime time.Time, optionalFields ...map[string]interface{}) {
	zapFields := m.logWithElapsedTime(ctx, startTime, optionalFields...)
	globalLogger.Info(action, zapFields...)
}

func (m *Manager) ErrorWithElapsedTime(ctx context.Context, action string, startTime time.Time, err error, optionalFields ...map[string]interface{}) {
	zapFields := m.logWithElapsedTime(ctx, startTime, optionalFields...)
	zapFields = append(zapFields, zap.Any("err", err))
	globalLogger.Error(action, zapFields...)
}
