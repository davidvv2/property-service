package log

// Logger is an interface used to log messages to the terminal with optional structured context.
//
// The basic logging methods (Info, Error, Panic, Fatal, Debug) are deprecated in favor of structured logging methods,
// which accept additional context in the form of key/value pairs for better clarity and filtering.
type Logger interface {
	Info(format string, obj ...any)
	Error(format string, obj ...any)
	Panic(format string, obj ...any)
	Fatal(format string, obj ...any)
	Debug(format string, obj ...any)

	// Structured logging methods accept a message and a Fields object which contains key/value pairs providing additional context.
	InfoWithFields(message string, fields Fields)
	ErrorWithFields(message string, fields Fields)
	PanicWithFields(message string, fields Fields)
	FatalWithFields(message string, fields Fields)
	DebugWithFields(message string, fields Fields)
}
