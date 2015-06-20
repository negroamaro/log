////////////////////////////////////////////////////////////////////////////////
// Copyright 2015 Negroamaro. All rights reserved.                            //
////////////////////////////////////////////////////////////////////////////////

package log

var internal = newNOPLogger(LvlTrace)

// register makes a logger available.
func register(l iLogger) {
	// checking nil.
	if l == nil {
		panic("log: register logger is nil")
	}
	// checking multiple times initialization.
	switch internal.(type) {
	case *nopLogger:
		internal = l
	default:
		panic("log: register was called many times.")
	}
}

// IsTraceEnabled : Check whether this logger is enabled for the TRACE Level.
func IsTraceEnabled() bool {
	return internal.IsTraceEnabled()
}

// IsDebugEnabled : Check whether this logger is enabled for the DEBUG Level.
func IsDebugEnabled() bool {
	return internal.IsDebugEnabled()
}

// IsInfoEnabled : Check whether this logger is enabled for the INFO Level.
func IsInfoEnabled() bool {
	return internal.IsInfoEnabled()
}

// IsWarnEnabled : Check whether this logger is enabled for the WARN Level.
func IsWarnEnabled() bool {
	return internal.IsWarnEnabled()
}

// IsErrorEnabled : Check whether this logger is enabled for the ERROR Level.
func IsErrorEnabled() bool {
	return internal.IsErrorEnabled()
}

// IsFatalEnabled : Check whether this logger is enabled for the FATAL Level.
func IsFatalEnabled() bool {
	return internal.IsFatalEnabled()
}

// Trace : Log a message at the TRACE level according to the specified format and argument.
func Trace(format string, params ...interface{}) {
	internal.Trace(format, params...)
}

// Debug : Log a message at the DEBUG level according to the specified format and argument.
func Debug(format string, params ...interface{}) {
	internal.Debug(format, params...)
}

// Info : Log a message at the INFO level according to the specified format and argument.
func Info(format string, params ...interface{}) {
	internal.Info(format, params...)
}

// Warn : Log a message at the WARN level according to the specified format and argument.
func Warn(format string, params ...interface{}) {
	internal.Warn(format, params...)
}

// Error : Log a message at the ERROR level according to the specified format and argument.
func Error(format string, params ...interface{}) {
	internal.Error(format, params...)
}

// Fatal : Log a message at the FATAL level according to the specified format and argument.
func Fatal(format string, params ...interface{}) {
	internal.Fatal(format, params...)
}

// SetLevel : Changing log level dynamically.
func SetLevel(level Level) error {
	return internal.SetLevel(level)
}

// EOF
