package log

// Logger : logging object used to log messages to the terminal.
type Logger interface {
	/*Info : logs a info message*/
	Info(format string, obj ...any)
	/*Error : logs a error message*/
	Error(format string, obj ...any)
	/*Panic : logs a Panic message*/
	Panic(format string, obj ...any)
	/*Fatal : logs a debug message*/
	Fatal(format string, obj ...any)
	/*Debug : logs a debug message*/
	Debug(format string, obj ...any)
}
