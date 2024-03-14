package errors

import (
	_g "fmt"

	_e "golang.org/x/xerrors"
)

func Errorf(processName, message string, arguments ...interface{}) error {
	return _da(_g.Sprintf(message, arguments...), processName)
}

func (_ae *processError) Error() string {
	var _cf string
	if _ae._c != "" {
		_cf = _ae._c
	}
	_cf += "\u0050r\u006f\u0063\u0065\u0073\u0073\u003a " + _ae._a
	if _ae._f != "" {
		_cf += "\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020" + _ae._f
	}
	if _ae._ed != nil {
		_cf += "\u002e\u0020" + _ae._ed.Error()
	}
	return _cf
}

var _ _e.Wrapper = (*processError)(nil)

type processError struct {
	_c  string
	_a  string
	_f  string
	_ed error
}

func Wrap(err error, processName, message string) error {
	if _bd, _ab := err.(*processError); _ab {
		_bd._c = ""
	}
	_ea := _da(message, processName)
	_ea._ed = err
	return _ea
}
func Error(processName, message string) error { return _da(message, processName) }
func Wrapf(err error, processName, message string, arguments ...interface{}) error {
	if _aa, _ce := err.(*processError); _ce {
		_aa._c = ""
	}
	_cc := _da(_g.Sprintf(message, arguments...), processName)
	_cc._ed = err
	return _cc
}
func (_gf *processError) Unwrap() error { return _gf._ed }
func _da(_df, _ag string) *processError {
	return &processError{_c: "\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d", _f: _df, _a: _ag}
}
