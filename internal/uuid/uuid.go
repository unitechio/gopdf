package uuid

import (
	_g "crypto/rand"
	_ac "encoding/hex"
	_a "io"
)

var _gdg UUID

func _ga(_fef []byte, _c UUID) {
	_ac.Encode(_fef, _c[:4])
	_fef[8] = '-'
	_ac.Encode(_fef[9:13], _c[4:6])
	_fef[13] = '-'
	_ac.Encode(_fef[14:18], _c[6:8])
	_fef[18] = '-'
	_ac.Encode(_fef[19:23], _c[8:10])
	_fef[23] = '-'
	_ac.Encode(_fef[24:], _c[10:])
}

var Nil = _gdg
var _ge = _g.Reader

func (_gb UUID) String() string { var _fe [36]byte; _ga(_fe[:], _gb); return string(_fe[:]) }
func MustUUID() UUID {
	uuid, _fc := NewUUID()
	if _fc != nil {
		panic(_fc)
	}
	return uuid
}
func NewUUID() (UUID, error) {
	var uuid UUID
	_, _f := _a.ReadFull(_ge, uuid[:])
	if _f != nil {
		return _gdg, _f
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid, nil
}

type UUID [16]byte
