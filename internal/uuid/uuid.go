package uuid

import (
	_b "crypto/rand"
	_e "encoding/hex"
	_g "io"
)

func _ba(_ed []byte, _ad UUID) {
	_e.Encode(_ed, _ad[:4])
	_ed[8] = '-'
	_e.Encode(_ed[9:13], _ad[4:6])
	_ed[13] = '-'
	_e.Encode(_ed[14:18], _ad[6:8])
	_ed[18] = '-'
	_e.Encode(_ed[19:23], _ad[8:10])
	_ed[23] = '-'
	_e.Encode(_ed[24:], _ad[10:])
}

func NewUUID() (UUID, error) {
	var uuid UUID
	_, _dd := _g.ReadFull(_d, uuid[:])
	if _dd != nil {
		return _ab, _dd
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid, nil
}

var Nil = _ab

func (_bd UUID) String() string { var _eb [36]byte; _ba(_eb[:], _bd); return string(_eb[:]) }

var (
	_ab UUID
	_d  = _b.Reader
)

func MustUUID() UUID {
	uuid, _gb := NewUUID()
	if _gb != nil {
		panic(_gb)
	}
	return uuid
}

type UUID [16]byte
