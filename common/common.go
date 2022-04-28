package common

import (
	_f "fmt"
	_fe "io"
	_df "os"
	_db "path/filepath"
	_d "runtime"
	_cb "time"
)

func (_ga ConsoleLogger) output(_fa _fe.Writer, _ea string, _ac string, _dcce ...interface{}) {
	_dfd(_fa, _ea, _ac, _dcce...)
}

// Debug logs debug message.
func (_ab ConsoleLogger) Debug(format string, args ...interface{}) {
	if _ab.LogLevel >= LogLevelDebug {
		_eee := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_ab.output(_df.Stdout, _eee, format, args...)
	}
}

// Warning does nothing for dummy logger.
func (DummyLogger) Warning(format string, args ...interface{}) {}

// LogLevel is the verbosity level for logging.
type LogLevel int

// Notice does nothing for dummy logger.
func (DummyLogger) Notice(format string, args ...interface{}) {}

// Error does nothing for dummy logger.
func (DummyLogger) Error(format string, args ...interface{}) {}

const _ccf = "\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034"

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{ LogLevel LogLevel }

const (
	LogLevelTrace   LogLevel = 5
	LogLevelDebug   LogLevel = 4
	LogLevelInfo    LogLevel = 3
	LogLevelNotice  LogLevel = 2
	LogLevelWarning LogLevel = 1
	LogLevelError   LogLevel = 0
)

// Trace does nothing for dummy logger.
func (DummyLogger) Trace(format string, args ...interface{}) {}

// Warning logs warning message.
func (_ccb WriterLogger) Warning(format string, args ...interface{}) {
	if _ccb.LogLevel >= LogLevelWarning {
		_cfa := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_ccb.logToWriter(_ccb.Output, _cfa, format, args...)
	}
}

// Trace logs trace message.
func (_eb ConsoleLogger) Trace(format string, args ...interface{}) {
	if _eb.LogLevel >= LogLevelTrace {
		_dcc := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_eb.output(_df.Stdout, _dcc, format, args...)
	}
}

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct {
	LogLevel LogLevel
	Output   _fe.Writer
}

// Error logs error message.
func (_ba WriterLogger) Error(format string, args ...interface{}) {
	if _ba.LogLevel >= LogLevelError {
		_cf := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_ba.logToWriter(_ba.Output, _cf, format, args...)
	}
}

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat(t _cb.Time) string { return t.Format(_ccf) + "\u0020\u0055\u0054\u0043" }

const _ege = 26

var Log Logger = DummyLogger{}

// Notice logs notice message.
func (_bcd ConsoleLogger) Notice(format string, args ...interface{}) {
	if _bcd.LogLevel >= LogLevelNotice {
		_efe := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_bcd.output(_df.Stdout, _efe, format, args...)
	}
}

// Debug logs debug message.
func (_fd WriterLogger) Debug(format string, args ...interface{}) {
	if _fd.LogLevel >= LogLevelDebug {
		_bcf := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_fd.logToWriter(_fd.Output, _bcf, format, args...)
	}
}

// Info logs info message.
func (_dd ConsoleLogger) Info(format string, args ...interface{}) {
	if _dd.LogLevel >= LogLevelInfo {
		_ddg := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_dd.output(_df.Stdout, _ddg, format, args...)
	}
}

// DummyLogger does nothing.
type DummyLogger struct{}

func _dfd(_da _fe.Writer, _gb string, _ff string, _gf ...interface{}) {
	_, _gbe, _efd, _dfe := _d.Caller(3)
	if !_dfe {
		_gbe = "\u003f\u003f\u003f"
		_efd = 0
	} else {
		_gbe = _db.Base(_gbe)
	}
	_gff := _f.Sprintf("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ", _gb, _gbe, _efd) + _ff + "\u000a"
	_f.Fprintf(_da, _gff, _gf...)
}

// Error logs error message.
func (_ef ConsoleLogger) Error(format string, args ...interface{}) {
	if _ef.LogLevel >= LogLevelError {
		_bgb := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_ef.output(_df.Stdout, _bgb, format, args...)
	}
}

// NewConsoleLogger creates new console logger.
func NewConsoleLogger(logLevel LogLevel) *ConsoleLogger { return &ConsoleLogger{LogLevel: logLevel} }

// Notice logs notice message.
func (_cad WriterLogger) Notice(format string, args ...interface{}) {
	if _cad.LogLevel >= LogLevelNotice {
		_bgd := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_cad.logToWriter(_cad.Output, _bgd, format, args...)
	}
}

const Version = "\u0033\u002e\u0033\u0034\u002e\u0030"

// Debug does nothing for dummy logger.
func (DummyLogger) Debug(format string, args ...interface{}) {}

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fg WriterLogger) IsLogLevel(level LogLevel) bool { return _fg.LogLevel >= level }

// Trace logs trace message.
func (_fdf WriterLogger) Trace(format string, args ...interface{}) {
	if _fdf.LogLevel >= LogLevelTrace {
		_eab := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_fdf.logToWriter(_fdf.Output, _eab, format, args...)
	}
}

const _ge = 15
const _ggg = 4

func (_gg WriterLogger) logToWriter(_bcdg _fe.Writer, _efa string, _gab string, _fdc ...interface{}) {
	_dfd(_bcdg, _efa, _gab, _fdc)
}

// IsLogLevel returns true from dummy logger.
func (DummyLogger) IsLogLevel(level LogLevel) bool { return true }

const _ccbd = 2022

// Info does nothing for dummy logger.
func (DummyLogger) Info(format string, args ...interface{}) {}

// Warning logs warning message.
func (_cba ConsoleLogger) Warning(format string, args ...interface{}) {
	if _cba.LogLevel >= LogLevelWarning {
		_bc := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_cba.output(_df.Stdout, _bc, format, args...)
	}
}

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger(logLevel LogLevel, writer _fe.Writer) *WriterLogger {
	_fc := WriterLogger{Output: writer, LogLevel: logLevel}
	return &_fc
}

var ReleasedAt = _cb.Date(_ccbd, _ggg, _ege, _ge, _fgb, 0, 0, _cb.UTC)

// SetLogger sets 'logger' to be used by the unidoc gopdf library.
func SetLogger(logger Logger) { Log = logger }

const _fgb = 30

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_cg ConsoleLogger) IsLogLevel(level LogLevel) bool { return _cg.LogLevel >= level }

// Info logs info message.
func (_fcg WriterLogger) Info(format string, args ...interface{}) {
	if _fcg.LogLevel >= LogLevelInfo {
		_eag := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_fcg.logToWriter(_fcg.Output, _eag, format, args...)
	}
}

// Logger is the interface used for logging in the gopdf package.
type Logger interface {
	Error(_b string, _cd ...interface{})
	Warning(_bf string, _e ...interface{})
	Notice(_cc string, _a ...interface{})
	Info(_aa string, _dc ...interface{})
	Debug(_eg string, _g ...interface{})
	Trace(_ca string, _bg ...interface{})
	IsLogLevel(_ee LogLevel) bool
}
