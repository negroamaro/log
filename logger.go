////////////////////////////////////////////////////////////////////////////////
// Copyright 2015 Negroamaro. All rights reserved.                            //
////////////////////////////////////////////////////////////////////////////////

package log

// iLogger is internal logger facade interface.
type iLogger interface {

	// Check whether this logger is enabled for the TRACE Level.
	IsTraceEnabled() bool

	// Check whether this logger is enabled for the DEBUG Level.
	IsDebugEnabled() bool

	// Check whether this logger is enabled for the INFO Level.
	IsInfoEnabled() bool

	// Check whether this logger is enabled for the WARN Level.
	IsWarnEnabled() bool

	// Check whether this logger is enabled for the ERROR Level.
	IsErrorEnabled() bool

	// Check whether this logger is enabled for the FATAL Level.
	IsFatalEnabled() bool

	// Log a message at the TRACE level according to the specified format and argument.
	Trace(format string, params ...interface{})

	// Log a message at the DEBUG level according to the specified format and argument.
	Debug(format string, params ...interface{})

	// Log a message at the INFO level according to the specified format and argument.
	Info(format string, params ...interface{})

	// Log a message at the WARN level according to the specified format and argument.
	Warn(format string, params ...interface{})

	// Log a message at the ERROR level according to the specified format and argument.
	Error(format string, params ...interface{})

	// Log a message at the FATAL level according to the specified format and argument.
	Fatal(format string, params ...interface{})

	// Changing log level dynamically.
	SetLevel(level Level) error
}

// EOF
