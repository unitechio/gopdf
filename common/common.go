package common

import (
	_g "fmt"
	_e "io"
	_gb "os"
	_c "path/filepath"
	_f "runtime"
	_bg "time"
)

// Error does nothing for dummy logger.
func (DummyLogger) Error(format string, args ...interface{}) {}

var Log Logger = DummyLogger{}

const _ae = 10

// Debug logs debug message.
func (_gg WriterLogger) Debug(format string, args ...interface{}) {
	if _gg.LogLevel >= LogLevelDebug {
		_ged := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_gg.logToWriter(_gg.Output, _ged, format, args...)
	}
}
func (_bd WriterLogger) logToWriter(_cde _e.Writer, _bgg string, _cbg string, _egg ...interface{}) {
	_bea(_cde, _bgg, _cbg, _egg)
}

// Warning logs warning message.
func (_ef ConsoleLogger) Warning(format string, args ...interface{}) {
	if _ef.LogLevel >= LogLevelWarning {
		_fdc := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_ef.output(_gb.Stdout, _fdc, format, args...)
	}
}

const _fg = 15

// Debug does nothing for dummy logger.
func (DummyLogger) Debug(format string, args ...interface{}) {}

// Debug logs debug message.
func (_fa ConsoleLogger) Debug(format string, args ...interface{}) {
	if _fa.LogLevel >= LogLevelDebug {
		_ca := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_fa.output(_gb.Stdout, _ca, format, args...)
	}
}

const _fddf = "\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034"

// NewConsoleLogger creates new console logger.
func NewConsoleLogger(logLevel LogLevel) *ConsoleLogger { return &ConsoleLogger{LogLevel: logLevel} }

// Trace logs trace message.
func (_ggf WriterLogger) Trace(format string, args ...interface{}) {
	if _ggf.LogLevel >= LogLevelTrace {
		_ab := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_ggf.logToWriter(_ggf.Output, _ab, format, args...)
	}
}

// Info logs info message.
func (_ba ConsoleLogger) Info(format string, args ...interface{}) {
	if _ba.LogLevel >= LogLevelInfo {
		_ceb := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_ba.output(_gb.Stdout, _ceb, format, args...)
	}
}

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ac WriterLogger) IsLogLevel(level LogLevel) bool { return _ac.LogLevel >= level }

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger(logLevel LogLevel, writer _e.Writer) *WriterLogger {
	_ec := WriterLogger{Output: writer, LogLevel: logLevel}
	return &_ec
}

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger(logger Logger) { Log = logger }

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat(t _bg.Time) string { return t.Format(_fddf) + "\u0020\u0055\u0054\u0043" }

// Trace logs trace message.
func (_eb ConsoleLogger) Trace(format string, args ...interface{}) {
	if _eb.LogLevel >= LogLevelTrace {
		_egc := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_eb.output(_gb.Stdout, _egc, format, args...)
	}
}

const _bc = 30

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fd ConsoleLogger) IsLogLevel(level LogLevel) bool { return _fd.LogLevel >= level }

// Error logs error message.
func (_be WriterLogger) Error(format string, args ...interface{}) {
	if _be.LogLevel >= LogLevelError {
		_af := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_be.logToWriter(_be.Output, _af, format, args...)
	}
}

// Logger is the interface used for logging in the unipdf package.
type Logger interface {
	Error(_bga string, _ce ...interface{})
	Warning(_ge string, _cea ...interface{})
	Notice(_ed string, _bf ...interface{})
	Info(_fb string, _ga ...interface{})
	Debug(_d string, _a ...interface{})
	Trace(_cee string, _eg ...interface{})
	IsLogLevel(_ee LogLevel) bool
}

// Info logs info message.
func (_cd WriterLogger) Info(format string, args ...interface{}) {
	if _cd.LogLevel >= LogLevelInfo {
		_acaf := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_cd.logToWriter(_cd.Output, _acaf, format, args...)
	}
}

// LogLevel is the verbosity level for logging.
type LogLevel int

// Warning logs warning message.
func (_caa WriterLogger) Warning(format string, args ...interface{}) {
	if _caa.LogLevel >= LogLevelWarning {
		_gf := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_caa.logToWriter(_caa.Output, _gf, format, args...)
	}
}

// IsLogLevel returns true from dummy logger.
func (DummyLogger) IsLogLevel(level LogLevel) bool { return true }

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct {
	LogLevel LogLevel
	Output   _e.Writer
}

// Error logs error message.
func (_dd ConsoleLogger) Error(format string, args ...interface{}) {
	if _dd.LogLevel >= LogLevelError {
		_eee := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_dd.output(_gb.Stdout, _eee, format, args...)
	}
}

const _fcd = 2022
const Version = "\u0033\u002e\u0034\u0030\u002e\u0030"

// Notice logs notice message.
func (_gac ConsoleLogger) Notice(format string, args ...interface{}) {
	if _gac.LogLevel >= LogLevelNotice {
		_bge := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_gac.output(_gb.Stdout, _bge, format, args...)
	}
}

// Info does nothing for dummy logger.
func (DummyLogger) Info(format string, args ...interface{}) {}
func (_cg ConsoleLogger) output(_gec _e.Writer, _fbg string, _dg string, _ceg ...interface{}) {
	_bea(_gec, _fbg, _dg, _ceg...)
}

const _efc = 28

func _bea(_gc _e.Writer, _edg string, _df string, _fc ...interface{}) {
	_, _fdd, _aff, _gcg := _f.Caller(3)
	if !_gcg {
		_fdd = "\u003f\u003f\u003f"
		_aff = 0
	} else {
		_fdd = _c.Base(_fdd)
	}
	_cgf := _g.Sprintf("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ", _edg, _fdd, _aff) + _df + "\u000a"
	_g.Fprintf(_gc, _cgf, _fc...)
}

// Trace does nothing for dummy logger.
func (DummyLogger) Trace(format string, args ...interface{}) {}

// Notice logs notice message.
func (_cb WriterLogger) Notice(format string, args ...interface{}) {
	if _cb.LogLevel >= LogLevelNotice {
		_aca := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_cb.logToWriter(_cb.Output, _aca, format, args...)
	}
}

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{ LogLevel LogLevel }

var ReleasedAt = _bg.Date(_fcd, _ae, _efc, _fg, _bc, 0, 0, _bg.UTC)

const (
	LogLevelTrace   LogLevel = 5
	LogLevelDebug   LogLevel = 4
	LogLevelInfo    LogLevel = 3
	LogLevelNotice  LogLevel = 2
	LogLevelWarning LogLevel = 1
	LogLevelError   LogLevel = 0
)

// Notice does nothing for dummy logger.
func (DummyLogger) Notice(format string, args ...interface{}) {}

// DummyLogger does nothing.
type DummyLogger struct{}

// Warning does nothing for dummy logger.
func (DummyLogger) Warning(format string, args ...interface{}) {}
