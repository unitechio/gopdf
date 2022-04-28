package errors

import (
	_g "fmt"

	_a "golang.org/x/xerrors"
)

func Error(processName, message string) error { return _ba(message, processName) }
func Errorf(processName, message string, arguments ...interface{}) error {
	return _ba(_g.Sprintf(message, arguments...), processName)
}
func (_eg *processError) Unwrap() error { return _eg._bc }

var _ _a.Wrapper = (*processError)(nil)

type processError struct {
	_da string
	_e  string
	_b  string
	_bc error
}

func Wrapf(err error, processName, message string, arguments ...interface{}) error {
	if _gc, _ec := err.(*processError); _ec {
		_gc._da = ""
	}
	_db := _ba(_g.Sprintf(message, arguments...), processName)
	_db._bc = err
	return _db
}
func _ba(_ad, _ge string) *processError {
	return &processError{_da: "\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d", _b: _ad, _e: _ge}
}
func Wrap(err error, processName, message string) error {
	if _bg, _ae := err.(*processError); _ae {
		_bg._da = ""
	}
	_bce := _ba(message, processName)
	_bce._bc = err
	return _bce
}
func (_be *processError) Error() string {
	var _c string
	if _be._da != "" {
		_c = _be._da
	}
	_c += "\u0050r\u006f\u0063\u0065\u0073\u0073\u003a " + _be._e
	if _be._b != "" {
		_c += "\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020" + _be._b
	}
	if _be._bc != nil {
		_c += "\u002e\u0020" + _be._bc.Error()
	}
	return _c
}
