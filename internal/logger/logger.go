package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.SugaredLogger
	zapLogger *zap.Logger
}

func NewLogger(level string) (*Logger, error) {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		return nil, fmt.Errorf("invalid log level: %s", level)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/logs.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     7, // days
		Compress:   true,
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleDebugging, zapLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, zapLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), consoleErrors, zapLevel),
	)

	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger := &Logger{SugaredLogger: zapLogger.Sugar(), zapLogger: zapLogger}

	return logger, nil
}

func (l *Logger) Close() error {
	return l.zapLogger.Sync()
}
