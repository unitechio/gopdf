package errors

import (
	_a "fmt"

	_b "golang.org/x/xerrors"
)

func (_e *processError) Error() string {
	var _bd string
	if _e._dc != "" {
		_bd = _e._dc
	}
	_bd += "\u0050r\u006f\u0063\u0065\u0073\u0073\u003a " + _e._g
	if _e._ga != "" {
		_bd += "\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020" + _e._ga
	}
	if _e._ac != nil {
		_bd += "\u002e\u0020" + _e._ac.Error()
	}
	return _bd
}
func Wrap(err error, processName, message string) error {
	if _gb, _c := err.(*processError); _c {
		_gb._dc = ""
	}
	_ba := _dcf(message, processName)
	_ba._ac = err
	return _ba
}
func Wrapf(err error, processName, message string, arguments ...interface{}) error {
	if _ec, _dcfc := err.(*processError); _dcfc {
		_ec._dc = ""
	}
	_ff := _dcf(_a.Sprintf(message, arguments...), processName)
	_ff._ac = err
	return _ff
}
func Errorf(processName, message string, arguments ...interface{}) error {
	return _dcf(_a.Sprintf(message, arguments...), processName)
}
func Error(processName, message string) error { return _dcf(message, processName) }
func (_db *processError) Unwrap() error       { return _db._ac }

type processError struct {
	_dc string
	_g  string
	_ga string
	_ac error
}

func _dcf(_de, _bf string) *processError {
	return &processError{_dc: "\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d", _ga: _de, _g: _bf}
}

var _ _b.Wrapper = (*processError)(nil)
