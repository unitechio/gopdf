package endian

import (
	_g "encoding/binary"
	_e "unsafe"
)

var (
	ByteOrder _g.ByteOrder
	_gf       bool
)

func init() {
	const _b = int(_e.Sizeof(0))
	_ea := 1
	_fc := (*[_b]byte)(_e.Pointer(&_ea))
	if _fc[0] == 0 {
		_gf = true
		ByteOrder = _g.BigEndian
	} else {
		ByteOrder = _g.LittleEndian
	}
}
func IsBig() bool    { return _gf }
func IsLittle() bool { return !_gf }
