package uuid

import (
	_c "crypto/rand"
	_g "encoding/hex"
	_e "io"
)

var _eg = _c.Reader

func _fg(_ag []byte, _ab UUID) {
	_g.Encode(_ag, _ab[:4])
	_ag[8] = '-'
	_g.Encode(_ag[9:13], _ab[4:6])
	_ag[13] = '-'
	_g.Encode(_ag[14:18], _ab[6:8])
	_ag[18] = '-'
	_g.Encode(_ag[19:23], _ab[8:10])
	_ag[23] = '-'
	_g.Encode(_ag[24:], _ab[10:])
}

var Nil = _da

func (_cc UUID) String() string { var _ce [36]byte; _fg(_ce[:], _cc); return string(_ce[:]) }

type UUID [16]byte

var _da UUID

func MustUUID() UUID {
	uuid, _de := NewUUID()
	if _de != nil {
		panic(_de)
	}
	return uuid
}
func NewUUID() (UUID, error) {
	var uuid UUID
	_, _eb := _e.ReadFull(_eg, uuid[:])
	if _eb != nil {
		return _da, _eb
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid, nil
}
