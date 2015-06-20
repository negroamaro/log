////////////////////////////////////////////////////////////////////////////////
// Copyright 2015 Negroamaro. All rights reserved.                            //
////////////////////////////////////////////////////////////////////////////////

package log

// nop logger facade type.
type nopLogger struct {
	level Level
}

// internal constructor.
func newNOPLogger(level Level) iLogger {
	return &nopLogger{level}
}

// nop logger facade implementation.
func (l *nopLogger) IsTraceEnabled() bool {
	return l.level >= LvlTrace
}

// nop logger facade implementation.
func (l *nopLogger) IsDebugEnabled() bool {
	return l.level >= LvlDebug
}

// nop logger facade implementation.
func (l *nopLogger) IsInfoEnabled() bool {
	return l.level >= LvlInfo
}

// nop logger facade implementation.
func (l *nopLogger) IsWarnEnabled() bool {
	return l.level >= LvlWarn
}

// nop logger facade implementation.
func (l *nopLogger) IsErrorEnabled() bool {
	return l.level >= LvlError
}

// nop logger facade implementation.
func (l *nopLogger) IsFatalEnabled() bool {
	return l.level >= LvlFatal
}

// nop logger facade implementation.
func (l *nopLogger) Trace(format string, params ...interface{}) {
}

// nop logger facade implementation.
func (l *nopLogger) Debug(format string, params ...interface{}) {
}

// nop logger facade implementation.
func (l *nopLogger) Info(format string, params ...interface{}) {
}

// nop logger facade implementation.
func (l *nopLogger) Warn(format string, params ...interface{}) {
}

// nop logger facade implementation.
func (l *nopLogger) Error(format string, params ...interface{}) {
}

// nop logger facade implementation.
func (l *nopLogger) Fatal(format string, params ...interface{}) {
}

// nop logger facade implementation.
func (l *nopLogger) SetLevel(level Level) error {
	l.level = level
	return nil
}

// EOF
