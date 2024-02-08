package common

import (
	_c "fmt"
	_cb "io"
	_eg "os"
	_ca "path/filepath"
	_ec "runtime"
	_e "time"
)

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct {
	LogLevel LogLevel
	Output   _cb.Writer
}

var ReleasedAt = _e.Date(_ee, _ce, _bge, _gb, _bd, 0, 0, _e.UTC)

// Error does nothing for dummy logger.
func (DummyLogger) Error(format string, args ...interface{}) {}

// Notice logs notice message.
func (_caf WriterLogger) Notice(format string, args ...interface{}) {
	if _caf.LogLevel >= LogLevelNotice {
		_bf := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_caf.logToWriter(_caf.Output, _bf, format, args...)
	}
}

// Debug logs debug message.
func (_abd ConsoleLogger) Debug(format string, args ...interface{}) {
	if _abd.LogLevel >= LogLevelDebug {
		_gd := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_abd.output(_eg.Stdout, _gd, format, args...)
	}
}

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_gc WriterLogger) IsLogLevel(level LogLevel) bool { return _gc.LogLevel >= level }

// DummyLogger does nothing.
type DummyLogger struct{}

// LogLevel is the verbosity level for logging.
type LogLevel int

const _ee = 2024

// Debug does nothing for dummy logger.
func (DummyLogger) Debug(format string, args ...interface{}) {}

// IsLogLevel returns true from dummy logger.
func (DummyLogger) IsLogLevel(level LogLevel) bool { return true }

// Info logs info message.
func (_ffa WriterLogger) Info(format string, args ...interface{}) {
	if _ffa.LogLevel >= LogLevelInfo {
		_fc := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_ffa.logToWriter(_ffa.Output, _fc, format, args...)
	}
}

// NewConsoleLogger creates new console logger.
func NewConsoleLogger(logLevel LogLevel) *ConsoleLogger { return &ConsoleLogger{LogLevel: logLevel} }

// Trace logs trace message.
func (_ebf ConsoleLogger) Trace(format string, args ...interface{}) {
	if _ebf.LogLevel >= LogLevelTrace {
		_bc := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_ebf.output(_eg.Stdout, _bc, format, args...)
	}
}

// Trace does nothing for dummy logger.
func (DummyLogger) Trace(format string, args ...interface{}) {}
func (_afc WriterLogger) logToWriter(_faa _cb.Writer, _afbb string, _fae string, _aeb ...interface{}) {
	_de(_faa, _afbb, _fae, _aeb)
}

// Warning logs warning message.
func (_baeb WriterLogger) Warning(format string, args ...interface{}) {
	if _baeb.LogLevel >= LogLevelWarning {
		_afb := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_baeb.logToWriter(_baeb.Output, _afb, format, args...)
	}
}

// Trace logs trace message.
func (_fed WriterLogger) Trace(format string, args ...interface{}) {
	if _fed.LogLevel >= LogLevelTrace {
		_egdc := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_fed.logToWriter(_fed.Output, _egdc, format, args...)
	}
}

const _bge = 22

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat(t _e.Time) string { return t.Format(_bab) + "\u0020\u0055\u0054\u0043" }

// Notice logs notice message.
func (_be ConsoleLogger) Notice(format string, args ...interface{}) {
	if _be.LogLevel >= LogLevelNotice {
		_ac := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_be.output(_eg.Stdout, _ac, format, args...)
	}
}

// Logger is the interface used for logging in the unipdf package.
type Logger interface {
	Error(_f string, _cc ...interface{})
	Warning(_g string, _fd ...interface{})
	Notice(_b string, _eb ...interface{})
	Info(_a string, _ba ...interface{})
	Debug(_ff string, _ad ...interface{})
	Trace(_da string, _ga ...interface{})
	IsLogLevel(_bg LogLevel) bool
}

// Warning does nothing for dummy logger.
func (DummyLogger) Warning(format string, args ...interface{}) {}

// Warning logs warning message.
func (_cd ConsoleLogger) Warning(format string, args ...interface{}) {
	if _cd.LogLevel >= LogLevelWarning {
		_ab := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_cd.output(_eg.Stdout, _ab, format, args...)
	}
}

const _gb = 15

// Debug logs debug message.
func (_gf WriterLogger) Debug(format string, args ...interface{}) {
	if _gf.LogLevel >= LogLevelDebug {
		_gae := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_gf.logToWriter(_gf.Output, _gae, format, args...)
	}
}

// Notice does nothing for dummy logger.
func (DummyLogger) Notice(format string, args ...interface{}) {}

// Error logs error message.
func (_fe ConsoleLogger) Error(format string, args ...interface{}) {
	if _fe.LogLevel >= LogLevelError {
		_dd := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_fe.output(_eg.Stdout, _dd, format, args...)
	}
}

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger(logger Logger) { Log = logger }
func (_ag ConsoleLogger) output(_af _cb.Writer, _ea string, _bae string, _ddb ...interface{}) {
	_de(_af, _ea, _bae, _ddb...)
}

const Version = "\u0033\u002e\u0035\u0034\u002e\u0030"

// Info logs info message.
func (_cg ConsoleLogger) Info(format string, args ...interface{}) {
	if _cg.LogLevel >= LogLevelInfo {
		_dg := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_cg.output(_eg.Stdout, _dg, format, args...)
	}
}

// Error logs error message.
func (_fa WriterLogger) Error(format string, args ...interface{}) {
	if _fa.LogLevel >= LogLevelError {
		_ae := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_fa.logToWriter(_fa.Output, _ae, format, args...)
	}
}

const (
	LogLevelTrace   LogLevel = 5
	LogLevelDebug   LogLevel = 4
	LogLevelInfo    LogLevel = 3
	LogLevelNotice  LogLevel = 2
	LogLevelWarning LogLevel = 1
	LogLevelError   LogLevel = 0
)
const _bd = 30

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{ LogLevel LogLevel }

var Log Logger = DummyLogger{}

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_egd ConsoleLogger) IsLogLevel(level LogLevel) bool { return _egd.LogLevel >= level }

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger(logLevel LogLevel, writer _cb.Writer) *WriterLogger {
	_cab := WriterLogger{Output: writer, LogLevel: logLevel}
	return &_cab
}

const _ce = 1

func _de(_abg _cb.Writer, _ebfa string, _gfe string, _abc ...interface{}) {
	_, _ege, _bb, _egb := _ec.Caller(3)
	if !_egb {
		_ege = "\u003f\u003f\u003f"
		_bb = 0
	} else {
		_ege = _ca.Base(_ege)
	}
	_bea := _c.Sprintf("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ", _ebfa, _ege, _bb) + _gfe + "\u000a"
	_c.Fprintf(_abg, _bea, _abc...)
}

// Info does nothing for dummy logger.
func (DummyLogger) Info(format string, args ...interface{}) {}

const _bab = "\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034"
