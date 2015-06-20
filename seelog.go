// +build negroamaro_log_is_seelog

////////////////////////////////////////////////////////////////////////////////
// Copyright 2015 Negroamaro. All rights reserved.                            //
////////////////////////////////////////////////////////////////////////////////

package log

import (
	"encoding/xml"
	"github.com/cihub/seelog"
	"io/ioutil"
	"os"
	"reflect"
)

const (
	prefix = ""
	indent = "  "
)

var defaultConfig = []byte(`
<seelog type="sync" minlevel="trace" maxlevel="critical">
  <outputs>
    <console formatid="default"/>
  </outputs>
  <formats>
    <format id="default" format="[%Date(2006/01/02 15:04:05.000)][%LEV] %Msg%n"/>
  </formats>
</seelog>`)

////////////////////////////////////////////////////////////////////////////////
// seelog configuration scheme.
////////////////////////////////////////////////////////////////////////////////

// config : <seelog> element type
type config struct {
	// attributes
	Type          string `xml:"type,attr,omitempty"`
	Asyncinterval int64  `xml:"asyncinterval,attr,omitempty"`
	MinLevel      string `xml:"minlevel,attr,omitempty"`
	MaxLevel      string `xml:"maxlevel,attr,omitempty"`
	// child elements
	Exceptions *exceptions `xml:"exceptions,omitempty"`
	Outputs    *outputs    `xml:"outputs,omitempty"`
	Formats    *formats    `xml:"formats,omitempty"`
}

// exceptions : <exceptions> element type
type exceptions struct {
	// child elements
	Exception []*exception `xml:"exception,omitempty"`
}

// exception : <exception> element type
type exception struct {
	// attributes
	FuncPattern string `xml:"funcpattern,attr,omitempty"`
	MinLevel    string `xml:"minlevel,attr,omitempty"`
	FilePattern string `xml:"filepattern,attr,omitempty"`
}

// outputs : <outputs> element type
type outputs struct {
	// attributes
	FormatID string `xml:"formatid,attr,omitempty"`
	// child elements
	Console     *console     `xml:"exception,omitempty"`
	Splitter    []*splitter  `xml:"splitter,omitempty"`
	RollingFile *rollingFile `xml:"rollingfile,omitempty"`
	Buffered    *buffered    `xml:"buffered,omitempty"`
	Filter      *filter      `xml:"filter,omitempty"`
}

type output struct {
	// attributes
	FormatID string `xml:"formatid,attr,omitempty"`
	Name     string `xml:"name,attr,omitempty"`
}

// file : <file> element type
type file struct {
	// attributes
	Path string `xml:"path,attr,omitempty"`
}

// console : <console> element type
type console struct {
	*output
}

// splitter : <splitter> element type
type splitter struct {
	*output
	// child elements
	File []*file `xml:"file,omitempty"`
}

// rollingFile : <rollingfile> element type
type rollingFile struct {
	*output
	// attributes
	Type        string `xml:"type,attr,omitempty"`
	Filename    string `xml:"filename,attr,omitempty"`
	MaxSize     int64  `xml:"maxsize,attr,omitempty"`
	MaxRolls    int64  `xml:"maxrolls,attr,omitempty"`
	DatePattern string `xml:"datepattern,attr,omitempty"`
}

// buffered : <buffered> element type
type buffered struct {
	*output
	// attributes
	Size        int64 `xml:"size,attr,omitempty"`
	FlushPeriod int64 `xml:"flushperiod,attr,omitempty"`
	// child elements
	File *file `xml:"file,omitempty"`
}

// filter : <filter> element type
type filter struct {
	// attributes
	Levels string `xml:"levels,attr,omitempty"`
	// child elements
	File *file `xml:"file,omitempty"`
	// TODO skip unused child element <smtp>, <conn>, ...
}

// formats : <formats> element type
type formats struct {
	// child elements
	Format []*format `xml:"format,omitempty"`
}

// Format : <format> element type
type format struct {
	// attributes
	ID     string `xml:"id,attr,omitempty"`
	Format string `xml:"format,attr,omitempty"`
}

func marshal(c config) ([]byte, error) {
	// set root element name
	obj := struct {
		config
		XMLName struct{} `xml:"seelog"`
	}{config: c}
	//return xml.Marshal(obj)
	return xml.MarshalIndent(obj, prefix, indent)
}

func unmarshal(bytes []byte, c *config) error {
	return xml.Unmarshal(bytes, c)
}

////////////////////////////////////////////////////////////////////////////////
// seelog logger implementation.
////////////////////////////////////////////////////////////////////////////////

// package initializer.
func init() {

	// initialize seelog default logger.
	cfg := defaultConfig
	defaultLogger, _ := seelog.LoggerFromConfigAsBytes(cfg)
	seelog.ReplaceLogger(defaultLogger)

	defer func() {
		var obj config
		e := unmarshal(cfg, &obj)
		if e != nil {
			// unreachable, because seelog.LoggerFromConfigAsBytes(config) passed.
		}
		// register seelog logger implementation.
		register(newSeelogLogger(obj))
	}()

	// get logger configuration filename from os environment.
	filename := os.ExpandEnv(configFilename)
	// filename checking.
	_, e := os.Stat(filename)
	if e != nil {
		seelog.Warnf("invalid '%s' environment value. (value:'%s', error:'%s')", configFilename, filename, e)
		return
	}
	// raed configuration file.
	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		seelog.Warnf("read configuration file failed. (error:'%s')", e)
		return
	}
	// create seelog logger.
	logger, e := seelog.LoggerFromConfigAsBytes(bytes)
	if e != nil {
		seelog.Warnf("can't initialize seelog. (error:'%s')", e)
		return
	}
	// replace seelog default logger.
	seelog.ReplaceLogger(logger)
	// update logger config
	cfg = bytes
}

// seelog logger type.
type seelogLogger struct {
	config config
	level  []bool
}

// internal constructor.
func newSeelogLogger(config config) iLogger {
	// get seelog internal log level field.
	unusedLevels := reflect.ValueOf(seelog.Current).Elem().FieldByName("unusedLevels")
	level := make([]bool, unusedLevels.Len())
	for i := 0; i < len(level); i++ {
		level[i] = !unusedLevels.Index(i).Bool()
	}
	return &seelogLogger{config, level}
}

// seelog logger implementation.
func (l *seelogLogger) IsTraceEnabled() bool {
	return l.level[LvlTrace]
}

// seelog logger implementation.
func (l *seelogLogger) IsDebugEnabled() bool {
	return l.level[LvlDebug]
}

// seelog logger implementation.
func (l *seelogLogger) IsInfoEnabled() bool {
	return l.level[LvlInfo]
}

// seelog logger implementation.
func (l *seelogLogger) IsWarnEnabled() bool {
	return l.level[LvlWarn]
}

// seelog logger implementation.
func (l *seelogLogger) IsErrorEnabled() bool {
	return l.level[LvlError]
}

// seelog logger implementation.
func (l *seelogLogger) IsFatalEnabled() bool {
	return l.level[LvlFatal]
}

// seelog logger implementation.
func (l *seelogLogger) Trace(format string, params ...interface{}) {
	seelog.Tracef(format, params...)
}

// seelog logger implementation.
func (l *seelogLogger) Debug(format string, params ...interface{}) {
	seelog.Debugf(format, params...)
}

// seelog logger implementation.
func (l *seelogLogger) Info(format string, params ...interface{}) {
	seelog.Infof(format, params...)
}

// seelog logger implementation.
func (l *seelogLogger) Warn(format string, params ...interface{}) {
	seelog.Warnf(format, params...)
}

// seelog logger implementation.
func (l *seelogLogger) Error(format string, params ...interface{}) {
	seelog.Errorf(format, params...)
}

// seelog logger implementation.
func (l *seelogLogger) Fatal(format string, params ...interface{}) {
	seelog.Criticalf(format, params...)
}

// seelog logger implementation.
func (l *seelogLogger) SetLevel(level Level) error {
	// update seelog logger
	currentLevel := l.config.MinLevel
	l.config.MinLevel = toString(level)
	bytes, e := marshal(l.config)
	if e != nil {
		l.config.MinLevel = currentLevel
		return e
	}
	logger, e := seelog.LoggerFromConfigAsBytes(bytes)
	if e != nil {
		l.config.MinLevel = currentLevel
		return e
	}
	seelog.ReplaceLogger(logger)
	// update internal logger level
	for i := int(LvlTrace); i <= int(LvlFatal); i++ {
		l.level[i] = (i >= int(level))
	}
	return nil
}

func toString(l Level) string {
	switch l {
	case LvlTrace:
		return "trace"
	case LvlDebug:
		return "debug"
	case LvlInfo:
		return "info"
	case LvlWarn:
		return "warn"
	case LvlError:
		return "error"
	case LvlFatal:
		return "critical"
	default:
		return "unknown"
	}
}

func fromString(l string) Level {
	switch l {
	case "trace":
		return LvlTrace
	case "debug":
		return LvlDebug
	case "info":
		return LvlInfo
	case "warn":
		return LvlWarn
	case "error":
		return LvlError
	case "critical":
		return LvlFatal
	default:
		return lvlUnknown
	}
}

// EOF
