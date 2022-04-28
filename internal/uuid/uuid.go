package uuid

import (
	_f "crypto/rand"
	_a "encoding/hex"
	_fd "io"
)

var Nil = _cbc
var _cbc UUID

func (_d UUID) String() string { var _cg [36]byte; _e(_cg[:], _d); return string(_cg[:]) }

var _ae = _f.Reader

func NewUUID() (UUID, error) {
	var uuid UUID
	_, _cd := _fd.ReadFull(_ae, uuid[:])
	if _cd != nil {
		return _cbc, _cd
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid, nil
}
func MustUUID() UUID {
	uuid, _cb := NewUUID()
	if _cb != nil {
		panic(_cb)
	}
	return uuid
}
func _e(_ec []byte, _cf UUID) {
	_a.Encode(_ec, _cf[:4])
	_ec[8] = '-'
	_a.Encode(_ec[9:13], _cf[4:6])
	_ec[13] = '-'
	_a.Encode(_ec[14:18], _cf[6:8])
	_ec[18] = '-'
	_a.Encode(_ec[19:23], _cf[8:10])
	_ec[23] = '-'
	_a.Encode(_ec[24:], _cf[10:])
}

type UUID [16]byte
