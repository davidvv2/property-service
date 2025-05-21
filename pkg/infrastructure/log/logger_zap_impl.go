package log

import (
	"fmt"
	"os"

	"property-service/pkg/configs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Enforce the implementation of the Connector interface.
var _ Logger = (*LoggerZapImpl)(nil)

// loggerZapImpl : struct containing all necessary objects for logging.
type LoggerZapImpl struct {
	zapLogger *zap.Logger
	config    *configs.BackendStruct
}

// New : This function will initialise zap logging.
func NewZapImpl(configs *configs.BackendStruct) LoggerZapImpl {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	consoleConfig := zap.NewDevelopmentEncoderConfig()

	consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
	consoleConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	zaps := zap.New(core,
		zap.AddCallerSkip(1),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	return LoggerZapImpl{
		zapLogger: zaps,
		config:    configs,
	}
}

// Info : logs a info message.
func (l LoggerZapImpl) Info(format string, obj ...interface{}) {
	l.zapLogger.Info(fmt.Sprintf(format, obj...))
}

// Error : logs a error message.
func (l LoggerZapImpl) Error(format string, obj ...interface{}) {
	l.zapLogger.Error(fmt.Sprintf(format, obj...))
}

// Panic : logs a Panic message.
func (l LoggerZapImpl) Panic(format string, obj ...interface{}) {
	l.zapLogger.Panic(fmt.Sprintf(format, obj...))
}

// Fatal : logs a debug message.
func (l LoggerZapImpl) Fatal(format string, obj ...interface{}) {
	l.zapLogger.Fatal(fmt.Sprintf(format, obj...))
}

// Debug : logs a debug message.
func (l LoggerZapImpl) Debug(format string, obj ...interface{}) {
	if l.config.Environment == "development" {
		l.zapLogger.Debug(fmt.Sprintf(format, obj...))
	}
}

// Info logs an info message.
func (l LoggerZapImpl) InfoWithFields(message string, fields Fields) {
	l.zapLogger.Info(message, toZapFields(fields)...)
}

// Error logs an error message.
func (l LoggerZapImpl) ErrorWithFields(message string, fields Fields) {
	l.zapLogger.Error(message, toZapFields(fields)...)
}

// Panic logs a panic message.
func (l LoggerZapImpl) PanicWithFields(message string, fields Fields) {
	l.zapLogger.Panic(message, toZapFields(fields)...)
}

// Fatal logs a fatal message.
func (l LoggerZapImpl) FatalWithFields(message string, fields Fields) {
	l.zapLogger.Fatal(message, toZapFields(fields)...)
}

// Debug logs a debug message only in development.
func (l LoggerZapImpl) DebugWithFields(message string, fields Fields) {
	if l.config.Environment == "development" {
		l.zapLogger.Debug(message, toZapFields(fields)...)
	}
}
