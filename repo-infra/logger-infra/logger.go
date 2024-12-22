package logger_infra

import (
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
	"os"

	log_golang "log"
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

func NewLogger(lc fx.Lifecycle, config *config_dto.Config) (*zap.SugaredLogger, error) {
	// To check or create log file
	appLogFile, err := os.OpenFile(config.Log.AppLogFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("NewLogger: error opening log file: %s", err.Error())
		return nil, err
	}

	// Registering the lumberjack sink
	lumberjackLogFile := fmt.Sprintf("lumberjack:%s", config.Log.AppLogFile)
	err = zap.RegisterSink("lumberjack", func(url *url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &lumberjack.Logger{
				Filename:   url.Path,
				MaxSize:    1024, // megabytes
				MaxBackups: 5,
				Compress:   true, // disabled by default
			},
		}, nil
	})

	// Configuring the logger
	var cfg zap.Config

	if config.Debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	cfg.Sampling = nil           // disabling sampling to prevent dropping of messages, it can be enabled if required
	cfg.DisableStacktrace = true //overwriting the default output paths and ErrorOutputPaths from stderr to appLogFile
	cfg.OutputPaths = []string{lumberjackLogFile}
	cfg.ErrorOutputPaths = []string{lumberjackLogFile}

	// Building the logger
	logger, err := cfg.Build()
	if err != nil {
		fmt.Printf("NewLogger: error config Build: %v", err)
		return nil, err
	}
	logger.Info("NewLogger: logger construction succeeded")
	log := logger.Sugar()
	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(context context.Context) error {
			log.Info("NewLogger: Flushing log data")
			return StopLogger(log)
		},
	})
	log_golang.SetOutput(appLogFile)
	os.Stderr = appLogFile
	os.Stdout = appLogFile
	return log, nil
}

func StopLogger(logger *zap.SugaredLogger) error {
	return logger.Sync()
}
