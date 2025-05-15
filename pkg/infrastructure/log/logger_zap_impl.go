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
	Sugger *zap.Logger
	config *configs.BackendStruct
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
		Sugger: zaps,
		config: configs,
	}
}

// Info : logs a info message.
func (l LoggerZapImpl) Info(format string, obj ...interface{}) {
	l.Sugger.Info(fmt.Sprintf(format, obj...))
}

// Error : logs a error message.
func (l LoggerZapImpl) Error(format string, obj ...interface{}) {
	l.Sugger.Error(fmt.Sprintf(format, obj...))
}

// Panic : logs a Panic message.
func (l LoggerZapImpl) Panic(format string, obj ...interface{}) {
	l.Sugger.Panic(fmt.Sprintf(format, obj...))
}

// Fatal : logs a debug message.
func (l LoggerZapImpl) Fatal(format string, obj ...interface{}) {
	l.Sugger.Fatal(fmt.Sprintf(format, obj...))
}

// Debug : logs a debug message.
func (l LoggerZapImpl) Debug(format string, obj ...interface{}) {
	if l.config.Environment == "development" {
		l.Sugger.Debug(fmt.Sprintf(format, obj...))
	}
}
