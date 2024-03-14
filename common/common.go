package common

import (
	_e "fmt"
	_fe "io"
	_fc "os"
	_g "path/filepath"
	_f "runtime"
	_c "time"
)

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_eb ConsoleLogger) IsLogLevel(level LogLevel) bool { return _eb.LogLevel >= level }

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger(logLevel LogLevel, writer _fe.Writer) *WriterLogger {
	_cg := WriterLogger{Output: writer, LogLevel: logLevel}
	return &_cg
}

const Version = "\u0033\u002e\u0035\u0035\u002e\u0030"

var Log Logger = DummyLogger{}

// Error logs error message.
func (_fg ConsoleLogger) Error(format string, args ...interface{}) {
	if _fg.LogLevel >= LogLevelError {
		_ba := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_fg.output(_fc.Stdout, _ba, format, args...)
	}
}

// Warning does nothing for dummy logger.
func (DummyLogger) Warning(format string, args ...interface{}) {}

// NewConsoleLogger creates new console logger.
func NewConsoleLogger(logLevel LogLevel) *ConsoleLogger { return &ConsoleLogger{LogLevel: logLevel} }

const (
	LogLevelTrace   LogLevel = 5
	LogLevelDebug   LogLevel = 4
	LogLevelInfo    LogLevel = 3
	LogLevelNotice  LogLevel = 2
	LogLevelWarning LogLevel = 1
	LogLevelError   LogLevel = 0
)

// Notice logs notice message.
func (_ed WriterLogger) Notice(format string, args ...interface{}) {
	if _ed.LogLevel >= LogLevelNotice {
		_aed := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_ed.logToWriter(_ed.Output, _aed, format, args...)
	}
}

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{ LogLevel LogLevel }

// Trace logs trace message.
func (_fca ConsoleLogger) Trace(format string, args ...interface{}) {
	if _fca.LogLevel >= LogLevelTrace {
		_age := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_fca.output(_fc.Stdout, _age, format, args...)
	}
}

// Notice logs notice message.
func (_ga ConsoleLogger) Notice(format string, args ...interface{}) {
	if _ga.LogLevel >= LogLevelNotice {
		_ae := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_ga.output(_fc.Stdout, _ae, format, args...)
	}
}

// Error logs error message.
func (_df WriterLogger) Error(format string, args ...interface{}) {
	if _df.LogLevel >= LogLevelError {
		_cf := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_df.logToWriter(_df.Output, _cf, format, args...)
	}
}

// Warning logs warning message.
func (_fea ConsoleLogger) Warning(format string, args ...interface{}) {
	if _fea.LogLevel >= LogLevelWarning {
		_dc := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_fea.output(_fc.Stdout, _dc, format, args...)
	}
}

// Trace logs trace message.
func (_da WriterLogger) Trace(format string, args ...interface{}) {
	if _da.LogLevel >= LogLevelTrace {
		_efb := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_da.logToWriter(_da.Output, _efb, format, args...)
	}
}

func _fd(_dgc _fe.Writer, _gca string, _de string, _ca ...interface{}) {
	_, _gabb, _ebe, _bg := _f.Caller(3)
	if !_bg {
		_gabb = "\u003f\u003f\u003f"
		_ebe = 0
	} else {
		_gabb = _g.Base(_gabb)
	}
	_bd := _e.Sprintf("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ", _gca, _gabb, _ebe) + _de + "\u000a"
	_e.Fprintf(_dgc, _bd, _ca...)
}

// Info logs info message.
func (_bab WriterLogger) Info(format string, args ...interface{}) {
	if _bab.LogLevel >= LogLevelInfo {
		_ged := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_bab.logToWriter(_bab.Output, _ged, format, args...)
	}
}

var ReleasedAt = _c.Date(_bda, _fa, _dd, _gcd, _eba, 0, 0, _c.UTC)

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat(t _c.Time) string { return t.Format(_ad) + "\u0020\u0055\u0054\u0043" }

func (_ccc ConsoleLogger) output(_gb _fe.Writer, _aga string, _eg string, _ff ...interface{}) {
	_fd(_gb, _aga, _eg, _ff...)
}

const _fa = 2

// Debug does nothing for dummy logger.
func (DummyLogger) Debug(format string, args ...interface{}) {}

const (
	_gcd = 15
	_dd  = 12
)

// Debug logs debug message.
func (_ccf WriterLogger) Debug(format string, args ...interface{}) {
	if _ccf.LogLevel >= LogLevelDebug {
		_efe := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_ccf.logToWriter(_ccf.Output, _efe, format, args...)
	}
}

const _eba = 30

// DummyLogger does nothing.
type DummyLogger struct{}

// Info does nothing for dummy logger.
func (DummyLogger) Info(format string, args ...interface{}) {}

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct {
	LogLevel LogLevel
	Output   _fe.Writer
}

// Logger is the interface used for logging in the unipdf package.
type Logger interface {
	Error(_ec string, _ea ...interface{})
	Warning(_ef string, _ac ...interface{})
	Notice(_cc string, _d ...interface{})
	Info(_b string, _bf ...interface{})
	Debug(_aa string, _gg ...interface{})
	Trace(_cb string, _ag ...interface{})
	IsLogLevel(_cd LogLevel) bool
}

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ace WriterLogger) IsLogLevel(level LogLevel) bool { return _ace.LogLevel >= level }

// Notice does nothing for dummy logger.
func (DummyLogger) Notice(format string, args ...interface{}) {}

// LogLevel is the verbosity level for logging.
type LogLevel int

// Info logs info message.
func (_feb ConsoleLogger) Info(format string, args ...interface{}) {
	if _feb.LogLevel >= LogLevelInfo {
		_gc := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_feb.output(_fc.Stdout, _gc, format, args...)
	}
}

// IsLogLevel returns true from dummy logger.
func (DummyLogger) IsLogLevel(level LogLevel) bool { return true }

// Error does nothing for dummy logger.
func (DummyLogger) Error(format string, args ...interface{}) {}

const _ad = "\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034"

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger(logger Logger) { Log = logger }

func (_fga WriterLogger) logToWriter(_af _fe.Writer, _cgd string, _cfg string, _gaa ...interface{}) {
	_fd(_af, _cgd, _cfg, _gaa)
}

// Warning logs warning message.
func (_dg WriterLogger) Warning(format string, args ...interface{}) {
	if _dg.LogLevel >= LogLevelWarning {
		_ge := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_dg.logToWriter(_dg.Output, _ge, format, args...)
	}
}

const _bda = 2024

// Debug logs debug message.
func (_bc ConsoleLogger) Debug(format string, args ...interface{}) {
	if _bc.LogLevel >= LogLevelDebug {
		_dcg := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_bc.output(_fc.Stdout, _dcg, format, args...)
	}
}

// Trace does nothing for dummy logger.
func (DummyLogger) Trace(format string, args ...interface{}) {}
