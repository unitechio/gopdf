package endian

import (
	_f "encoding/binary"
	_b "unsafe"
)

func IsLittle() bool { return !_ce }
func init() {
	const _fbd = int(_b.Sizeof(0))
	_cc := 1
	_g := (*[_fbd]byte)(_b.Pointer(&_cc))
	if _g[0] == 0 {
		_ce = true
		ByteOrder = _f.BigEndian
	} else {
		ByteOrder = _f.LittleEndian
	}
}

var (
	ByteOrder _f.ByteOrder
	_ce       bool
)

func IsBig() bool { return _ce }
