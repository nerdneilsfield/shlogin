package logger

import (
	"go.uber.org/zap"
	_ "go.uber.org/zap" // 防止被 go mod tidy 移除
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	structLogger  *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

var logger = NewLogger(false, false, "")

func NewLogger(verbose bool, saveToFile bool, logFilePath string) *Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var structLogger *zap.Logger
	var sugaredLogger *zap.SugaredLogger
	if verbose && saveToFile {
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(logFilePath),
			zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return level >= zapcore.DebugLevel
			}),
		)
		structLogger = zap.New(core)

		config := zap.Config{
			Level:         zap.NewAtomicLevelAt(zapcore.DebugLevel),
			Development:   true,
			Encoding:      "json",
			EncoderConfig: encoderConfig,
			OutputPaths:   []string{"stdout"},
		}
		temp_logger, _ := config.Build()
		sugaredLogger = temp_logger.Sugar()
	} else {
		structLogger = nil
		config := zap.Config{
			Level:         zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Development:   false,
			Encoding:      "json",
			EncoderConfig: encoderConfig,
			OutputPaths:   []string{"stdout"},
		}
		temp_logger, _ := config.Build()
		sugaredLogger = temp_logger.Sugar()
	}

	logger_ := &Logger{
		structLogger:  structLogger,
		sugaredLogger: sugaredLogger,
	}
	defer logger_.structLogger.Sync()
	defer logger_.sugaredLogger.Sync()
	return logger_
}

func SetLogger(logger_ *Logger) {
	logger = logger_
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.structLogger != nil {
		l.structLogger.Info(msg, args...)
	}
	l.sugaredLogger.Infof(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.structLogger != nil {
		l.structLogger.Error(msg, args...)
	}
	l.sugaredLogger.Errorf(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	if l.structLogger != nil {
		l.structLogger.Fatal(msg, args...)
	}
	l.sugaredLogger.Fatalf(msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.structLogger != nil {
		l.structLogger.Debug(msg, args...)
	}
	l.sugaredLogger.Debugf(msg, args...)
}

func (l *Logger) Panic(msg string, args ...interface{}) {
	if l.structLogger != nil {
		l.structLogger.Panic(msg, args...)
	}
	l.sugaredLogger.Panicf(msg, args...)
}
