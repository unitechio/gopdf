package errors

import (
	_c "fmt"

	_fd "golang.org/x/xerrors"
)

func _db(_ed, _be string) *processError {
	return &processError{_b: "\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d", _eg: _ed, _ee: _be}
}
func Error(processName, message string) error { return _db(message, processName) }
func Wrap(err error, processName, message string) error {
	if _fa, _a := err.(*processError); _a {
		_fa._b = ""
	}
	_da := _db(message, processName)
	_da._d = err
	return _da
}
func (_cc *processError) Error() string {
	var _g string
	if _cc._b != "" {
		_g = _cc._b
	}
	_g += "\u0050r\u006f\u0063\u0065\u0073\u0073\u003a " + _cc._ee
	if _cc._eg != "" {
		_g += "\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020" + _cc._eg
	}
	if _cc._d != nil {
		_g += "\u002e\u0020" + _cc._d.Error()
	}
	return _g
}

var _ _fd.Wrapper = (*processError)(nil)

func Errorf(processName, message string, arguments ...interface{}) error {
	return _db(_c.Sprintf(message, arguments...), processName)
}

type processError struct {
	_b  string
	_ee string
	_eg string
	_d  error
}

func Wrapf(err error, processName, message string, arguments ...interface{}) error {
	if _ae, _de := err.(*processError); _de {
		_ae._b = ""
	}
	_ec := _db(_c.Sprintf(message, arguments...), processName)
	_ec._d = err
	return _ec
}
func (_dc *processError) Unwrap() error { return _dc._d }
