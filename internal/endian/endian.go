package endian

import (
	_e "encoding/binary"
	_g "unsafe"
)

func IsBig() bool { return _d }
func init() {
	const _c = int(_g.Sizeof(0))
	_f := 1
	_ff := (*[_c]byte)(_g.Pointer(&_f))
	if _ff[0] == 0 {
		_d = true
		ByteOrder = _e.BigEndian
	} else {
		ByteOrder = _e.LittleEndian
	}
}

var (
	ByteOrder _e.ByteOrder
	_d        bool
)

func IsLittle() bool { return !_d }
