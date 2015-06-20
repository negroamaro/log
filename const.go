////////////////////////////////////////////////////////////////////////////////
// Copyright 2015 Negroamaro. All rights reserved.                            //
////////////////////////////////////////////////////////////////////////////////

package log

// OS environment key.
const (
	configFilename = "${NEGROAMARO_LOG_CONFIG_FILE}"
)

// Level is log level type.
type Level int

// log level constant
const (
	LvlTrace Level = iota
	LvlDebug
	LvlInfo
	LvlWarn
	LvlError
	LvlFatal
	// internal use only.
	lvlUnknown
)

// String is string converter for Level type.
func (l Level) String() string {
	switch l {
	case LvlTrace:
		return "TRACE"
	case LvlDebug:
		return "DEBUG"
	case LvlInfo:
		return "INFO"
	case LvlWarn:
		return "WARN"
	case LvlError:
		return "ERROR"
	case LvlFatal:
		return "FATAL"
	default:
		return "unknown"
	}
}

// EOF
