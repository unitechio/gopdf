package ccittfax

import (
	_a "errors"
	_ge "io"
	_e "math"

	_gf "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
)

func _gec(_dgd []byte, _faa int, _fgda code) ([]byte, int) {
	_cdd := 0
	for _cdd < _fgda.BitsWritten {
		_deb := _faa / 8
		_afb := _faa % 8
		if _deb >= len(_dgd) {
			_dgd = append(_dgd, 0)
		}
		_gbfe := 8 - _afb
		_aaf := _fgda.BitsWritten - _cdd
		if _gbfe > _aaf {
			_gbfe = _aaf
		}
		if _cdd < 8 {
			_dgd[_deb] = _dgd[_deb] | byte(_fgda.Code>>uint(8+_afb-_cdd))&_ecd[8-_gbfe-_afb]
		} else {
			_dgd[_deb] = _dgd[_deb] | (byte(_fgda.Code<<uint(_cdd-8))&_ecd[8-_gbfe])>>uint(_afb)
		}
		_faa += _gbfe
		_cdd += _gbfe
	}
	return _dgd, _faa
}
func init() {
	_gee = &treeNode{_cag: true, _bcbd: _ab}
	_ae = &treeNode{_bcbd: _ed, _gbe: _gee}
	_ae._gbfd = _ae
	_b = &tree{_fcgf: &treeNode{}}
	if _ag := _b.fillWithNode(12, 0, _ae); _ag != nil {
		panic(_ag.Error())
	}
	if _ea := _b.fillWithNode(12, 1, _gee); _ea != nil {
		panic(_ea.Error())
	}
	_d = &tree{_fcgf: &treeNode{}}
	for _fd := 0; _fd < len(_ecf); _fd++ {
		for _fe := 0; _fe < len(_ecf[_fd]); _fe++ {
			if _bb := _d.fill(_fd+2, int(_ecf[_fd][_fe]), int(_fg[_fd][_fe])); _bb != nil {
				panic(_bb.Error())
			}
		}
	}
	if _ba := _d.fillWithNode(12, 0, _ae); _ba != nil {
		panic(_ba.Error())
	}
	if _af := _d.fillWithNode(12, 1, _gee); _af != nil {
		panic(_af.Error())
	}
	_c = &tree{_fcgf: &treeNode{}}
	for _ce := 0; _ce < len(_gg); _ce++ {
		for _gbc := 0; _gbc < len(_gg[_ce]); _gbc++ {
			if _fdc := _c.fill(_ce+4, int(_gg[_ce][_gbc]), int(_bgg[_ce][_gbc])); _fdc != nil {
				panic(_fdc.Error())
			}
		}
	}
	if _aee := _c.fillWithNode(12, 0, _ae); _aee != nil {
		panic(_aee.Error())
	}
	if _ff := _c.fillWithNode(12, 1, _gee); _ff != nil {
		panic(_ff.Error())
	}
	_ee = &tree{_fcgf: &treeNode{}}
	if _ac := _ee.fill(4, 1, _f); _ac != nil {
		panic(_ac.Error())
	}
	if _bg := _ee.fill(3, 1, _gb); _bg != nil {
		panic(_bg.Error())
	}
	if _bbb := _ee.fill(1, 1, 0); _bbb != nil {
		panic(_bbb.Error())
	}
	if _eaa := _ee.fill(3, 3, 1); _eaa != nil {
		panic(_eaa.Error())
	}
	if _geb := _ee.fill(6, 3, 2); _geb != nil {
		panic(_geb.Error())
	}
	if _afe := _ee.fill(7, 3, 3); _afe != nil {
		panic(_afe.Error())
	}
	if _fc := _ee.fill(3, 2, -1); _fc != nil {
		panic(_fc.Error())
	}
	if _edg := _ee.fill(6, 2, -2); _edg != nil {
		panic(_edg.Error())
	}
	if _cf := _ee.fill(7, 2, -3); _cf != nil {
		panic(_cf.Error())
	}
}

var _gg = [...][]uint16{{0x7, 0x8, 0xb, 0xc, 0xe, 0xf}, {0x12, 0x13, 0x14, 0x1b, 0x7, 0x8}, {0x17, 0x18, 0x2a, 0x2b, 0x3, 0x34, 0x35, 0x7, 0x8}, {0x13, 0x17, 0x18, 0x24, 0x27, 0x28, 0x2b, 0x3, 0x37, 0x4, 0x8, 0xc}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1a, 0x1b, 0x2, 0x24, 0x25, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x3, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x4, 0x4a, 0x4b, 0x5, 0x52, 0x53, 0x54, 0x55, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x67, 0x68, 0xa, 0xb}, {0x98, 0x99, 0x9a, 0x9b, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb}, {}, {0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f}}

func (_afd *Decoder) decodeRowType2() error {
	if _afd._gfb {
		_afd._cd.Align()
	}
	if _eac := _afd.decode1D(); _eac != nil {
		return _eac
	}
	return nil
}

const (
	_ tiffType = iota
	_gef
	_cfd
	_eb
)

func (_afa *Decoder) decodeRowType6() error {
	if _afa._gfb {
		_afa._cd.Align()
	}
	if _afa._de {
		_afa._cd.Mark()
		_cb, _baf := _afa.tryFetchEOL()
		if _baf != nil {
			return _baf
		}
		if _cb {
			_cb, _baf = _afa.tryFetchEOL()
			if _baf != nil {
				return _baf
			}
			if _cb {
				return _ge.EOF
			}
		}
		_afa._cd.Reset()
	}
	return _afa.decode2D()
}
func (_cge *Encoder) appendEncodedRow(_def, _ebag []byte, _ebb int) []byte {
	if len(_def) > 0 && _ebb != 0 && !_cge.EncodedByteAlign {
		_def[len(_def)-1] = _def[len(_def)-1] | _ebag[0]
		_def = append(_def, _ebag[1:]...)
	} else {
		_def = append(_def, _ebag...)
	}
	return _def
}

type code struct {
	Code        uint16
	BitsWritten int
}

func init() {
	_be = make(map[int]code)
	_be[0] = code{Code: 13<<8 | 3<<6, BitsWritten: 10}
	_be[1] = code{Code: 2 << (5 + 8), BitsWritten: 3}
	_be[2] = code{Code: 3 << (6 + 8), BitsWritten: 2}
	_be[3] = code{Code: 2 << (6 + 8), BitsWritten: 2}
	_be[4] = code{Code: 3 << (5 + 8), BitsWritten: 3}
	_be[5] = code{Code: 3 << (4 + 8), BitsWritten: 4}
	_be[6] = code{Code: 2 << (4 + 8), BitsWritten: 4}
	_be[7] = code{Code: 3 << (3 + 8), BitsWritten: 5}
	_be[8] = code{Code: 5 << (2 + 8), BitsWritten: 6}
	_be[9] = code{Code: 4 << (2 + 8), BitsWritten: 6}
	_be[10] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_be[11] = code{Code: 5 << (1 + 8), BitsWritten: 7}
	_be[12] = code{Code: 7 << (1 + 8), BitsWritten: 7}
	_be[13] = code{Code: 4 << 8, BitsWritten: 8}
	_be[14] = code{Code: 7 << 8, BitsWritten: 8}
	_be[15] = code{Code: 12 << 8, BitsWritten: 9}
	_be[16] = code{Code: 5<<8 | 3<<6, BitsWritten: 10}
	_be[17] = code{Code: 6 << 8, BitsWritten: 10}
	_be[18] = code{Code: 2 << 8, BitsWritten: 10}
	_be[19] = code{Code: 12<<8 | 7<<5, BitsWritten: 11}
	_be[20] = code{Code: 13 << 8, BitsWritten: 11}
	_be[21] = code{Code: 13<<8 | 4<<5, BitsWritten: 11}
	_be[22] = code{Code: 6<<8 | 7<<5, BitsWritten: 11}
	_be[23] = code{Code: 5 << 8, BitsWritten: 11}
	_be[24] = code{Code: 2<<8 | 7<<5, BitsWritten: 11}
	_be[25] = code{Code: 3 << 8, BitsWritten: 11}
	_be[26] = code{Code: 12<<8 | 10<<4, BitsWritten: 12}
	_be[27] = code{Code: 12<<8 | 11<<4, BitsWritten: 12}
	_be[28] = code{Code: 12<<8 | 12<<4, BitsWritten: 12}
	_be[29] = code{Code: 12<<8 | 13<<4, BitsWritten: 12}
	_be[30] = code{Code: 6<<8 | 8<<4, BitsWritten: 12}
	_be[31] = code{Code: 6<<8 | 9<<4, BitsWritten: 12}
	_be[32] = code{Code: 6<<8 | 10<<4, BitsWritten: 12}
	_be[33] = code{Code: 6<<8 | 11<<4, BitsWritten: 12}
	_be[34] = code{Code: 13<<8 | 2<<4, BitsWritten: 12}
	_be[35] = code{Code: 13<<8 | 3<<4, BitsWritten: 12}
	_be[36] = code{Code: 13<<8 | 4<<4, BitsWritten: 12}
	_be[37] = code{Code: 13<<8 | 5<<4, BitsWritten: 12}
	_be[38] = code{Code: 13<<8 | 6<<4, BitsWritten: 12}
	_be[39] = code{Code: 13<<8 | 7<<4, BitsWritten: 12}
	_be[40] = code{Code: 6<<8 | 12<<4, BitsWritten: 12}
	_be[41] = code{Code: 6<<8 | 13<<4, BitsWritten: 12}
	_be[42] = code{Code: 13<<8 | 10<<4, BitsWritten: 12}
	_be[43] = code{Code: 13<<8 | 11<<4, BitsWritten: 12}
	_be[44] = code{Code: 5<<8 | 4<<4, BitsWritten: 12}
	_be[45] = code{Code: 5<<8 | 5<<4, BitsWritten: 12}
	_be[46] = code{Code: 5<<8 | 6<<4, BitsWritten: 12}
	_be[47] = code{Code: 5<<8 | 7<<4, BitsWritten: 12}
	_be[48] = code{Code: 6<<8 | 4<<4, BitsWritten: 12}
	_be[49] = code{Code: 6<<8 | 5<<4, BitsWritten: 12}
	_be[50] = code{Code: 5<<8 | 2<<4, BitsWritten: 12}
	_be[51] = code{Code: 5<<8 | 3<<4, BitsWritten: 12}
	_be[52] = code{Code: 2<<8 | 4<<4, BitsWritten: 12}
	_be[53] = code{Code: 3<<8 | 7<<4, BitsWritten: 12}
	_be[54] = code{Code: 3<<8 | 8<<4, BitsWritten: 12}
	_be[55] = code{Code: 2<<8 | 7<<4, BitsWritten: 12}
	_be[56] = code{Code: 2<<8 | 8<<4, BitsWritten: 12}
	_be[57] = code{Code: 5<<8 | 8<<4, BitsWritten: 12}
	_be[58] = code{Code: 5<<8 | 9<<4, BitsWritten: 12}
	_be[59] = code{Code: 2<<8 | 11<<4, BitsWritten: 12}
	_be[60] = code{Code: 2<<8 | 12<<4, BitsWritten: 12}
	_be[61] = code{Code: 5<<8 | 10<<4, BitsWritten: 12}
	_be[62] = code{Code: 6<<8 | 6<<4, BitsWritten: 12}
	_be[63] = code{Code: 6<<8 | 7<<4, BitsWritten: 12}
	_ad = make(map[int]code)
	_ad[0] = code{Code: 53 << 8, BitsWritten: 8}
	_ad[1] = code{Code: 7 << (2 + 8), BitsWritten: 6}
	_ad[2] = code{Code: 7 << (4 + 8), BitsWritten: 4}
	_ad[3] = code{Code: 8 << (4 + 8), BitsWritten: 4}
	_ad[4] = code{Code: 11 << (4 + 8), BitsWritten: 4}
	_ad[5] = code{Code: 12 << (4 + 8), BitsWritten: 4}
	_ad[6] = code{Code: 14 << (4 + 8), BitsWritten: 4}
	_ad[7] = code{Code: 15 << (4 + 8), BitsWritten: 4}
	_ad[8] = code{Code: 19 << (3 + 8), BitsWritten: 5}
	_ad[9] = code{Code: 20 << (3 + 8), BitsWritten: 5}
	_ad[10] = code{Code: 7 << (3 + 8), BitsWritten: 5}
	_ad[11] = code{Code: 8 << (3 + 8), BitsWritten: 5}
	_ad[12] = code{Code: 8 << (2 + 8), BitsWritten: 6}
	_ad[13] = code{Code: 3 << (2 + 8), BitsWritten: 6}
	_ad[14] = code{Code: 52 << (2 + 8), BitsWritten: 6}
	_ad[15] = code{Code: 53 << (2 + 8), BitsWritten: 6}
	_ad[16] = code{Code: 42 << (2 + 8), BitsWritten: 6}
	_ad[17] = code{Code: 43 << (2 + 8), BitsWritten: 6}
	_ad[18] = code{Code: 39 << (1 + 8), BitsWritten: 7}
	_ad[19] = code{Code: 12 << (1 + 8), BitsWritten: 7}
	_ad[20] = code{Code: 8 << (1 + 8), BitsWritten: 7}
	_ad[21] = code{Code: 23 << (1 + 8), BitsWritten: 7}
	_ad[22] = code{Code: 3 << (1 + 8), BitsWritten: 7}
	_ad[23] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_ad[24] = code{Code: 40 << (1 + 8), BitsWritten: 7}
	_ad[25] = code{Code: 43 << (1 + 8), BitsWritten: 7}
	_ad[26] = code{Code: 19 << (1 + 8), BitsWritten: 7}
	_ad[27] = code{Code: 36 << (1 + 8), BitsWritten: 7}
	_ad[28] = code{Code: 24 << (1 + 8), BitsWritten: 7}
	_ad[29] = code{Code: 2 << 8, BitsWritten: 8}
	_ad[30] = code{Code: 3 << 8, BitsWritten: 8}
	_ad[31] = code{Code: 26 << 8, BitsWritten: 8}
	_ad[32] = code{Code: 27 << 8, BitsWritten: 8}
	_ad[33] = code{Code: 18 << 8, BitsWritten: 8}
	_ad[34] = code{Code: 19 << 8, BitsWritten: 8}
	_ad[35] = code{Code: 20 << 8, BitsWritten: 8}
	_ad[36] = code{Code: 21 << 8, BitsWritten: 8}
	_ad[37] = code{Code: 22 << 8, BitsWritten: 8}
	_ad[38] = code{Code: 23 << 8, BitsWritten: 8}
	_ad[39] = code{Code: 40 << 8, BitsWritten: 8}
	_ad[40] = code{Code: 41 << 8, BitsWritten: 8}
	_ad[41] = code{Code: 42 << 8, BitsWritten: 8}
	_ad[42] = code{Code: 43 << 8, BitsWritten: 8}
	_ad[43] = code{Code: 44 << 8, BitsWritten: 8}
	_ad[44] = code{Code: 45 << 8, BitsWritten: 8}
	_ad[45] = code{Code: 4 << 8, BitsWritten: 8}
	_ad[46] = code{Code: 5 << 8, BitsWritten: 8}
	_ad[47] = code{Code: 10 << 8, BitsWritten: 8}
	_ad[48] = code{Code: 11 << 8, BitsWritten: 8}
	_ad[49] = code{Code: 82 << 8, BitsWritten: 8}
	_ad[50] = code{Code: 83 << 8, BitsWritten: 8}
	_ad[51] = code{Code: 84 << 8, BitsWritten: 8}
	_ad[52] = code{Code: 85 << 8, BitsWritten: 8}
	_ad[53] = code{Code: 36 << 8, BitsWritten: 8}
	_ad[54] = code{Code: 37 << 8, BitsWritten: 8}
	_ad[55] = code{Code: 88 << 8, BitsWritten: 8}
	_ad[56] = code{Code: 89 << 8, BitsWritten: 8}
	_ad[57] = code{Code: 90 << 8, BitsWritten: 8}
	_ad[58] = code{Code: 91 << 8, BitsWritten: 8}
	_ad[59] = code{Code: 74 << 8, BitsWritten: 8}
	_ad[60] = code{Code: 75 << 8, BitsWritten: 8}
	_ad[61] = code{Code: 50 << 8, BitsWritten: 8}
	_ad[62] = code{Code: 51 << 8, BitsWritten: 8}
	_ad[63] = code{Code: 52 << 8, BitsWritten: 8}
	_cfe = make(map[int]code)
	_cfe[64] = code{Code: 3<<8 | 3<<6, BitsWritten: 10}
	_cfe[128] = code{Code: 12<<8 | 8<<4, BitsWritten: 12}
	_cfe[192] = code{Code: 12<<8 | 9<<4, BitsWritten: 12}
	_cfe[256] = code{Code: 5<<8 | 11<<4, BitsWritten: 12}
	_cfe[320] = code{Code: 3<<8 | 3<<4, BitsWritten: 12}
	_cfe[384] = code{Code: 3<<8 | 4<<4, BitsWritten: 12}
	_cfe[448] = code{Code: 3<<8 | 5<<4, BitsWritten: 12}
	_cfe[512] = code{Code: 3<<8 | 12<<3, BitsWritten: 13}
	_cfe[576] = code{Code: 3<<8 | 13<<3, BitsWritten: 13}
	_cfe[640] = code{Code: 2<<8 | 10<<3, BitsWritten: 13}
	_cfe[704] = code{Code: 2<<8 | 11<<3, BitsWritten: 13}
	_cfe[768] = code{Code: 2<<8 | 12<<3, BitsWritten: 13}
	_cfe[832] = code{Code: 2<<8 | 13<<3, BitsWritten: 13}
	_cfe[896] = code{Code: 3<<8 | 18<<3, BitsWritten: 13}
	_cfe[960] = code{Code: 3<<8 | 19<<3, BitsWritten: 13}
	_cfe[1024] = code{Code: 3<<8 | 20<<3, BitsWritten: 13}
	_cfe[1088] = code{Code: 3<<8 | 21<<3, BitsWritten: 13}
	_cfe[1152] = code{Code: 3<<8 | 22<<3, BitsWritten: 13}
	_cfe[1216] = code{Code: 119 << 3, BitsWritten: 13}
	_cfe[1280] = code{Code: 2<<8 | 18<<3, BitsWritten: 13}
	_cfe[1344] = code{Code: 2<<8 | 19<<3, BitsWritten: 13}
	_cfe[1408] = code{Code: 2<<8 | 20<<3, BitsWritten: 13}
	_cfe[1472] = code{Code: 2<<8 | 21<<3, BitsWritten: 13}
	_cfe[1536] = code{Code: 2<<8 | 26<<3, BitsWritten: 13}
	_cfe[1600] = code{Code: 2<<8 | 27<<3, BitsWritten: 13}
	_cfe[1664] = code{Code: 3<<8 | 4<<3, BitsWritten: 13}
	_cfe[1728] = code{Code: 3<<8 | 5<<3, BitsWritten: 13}
	_fde = make(map[int]code)
	_fde[64] = code{Code: 27 << (3 + 8), BitsWritten: 5}
	_fde[128] = code{Code: 18 << (3 + 8), BitsWritten: 5}
	_fde[192] = code{Code: 23 << (2 + 8), BitsWritten: 6}
	_fde[256] = code{Code: 55 << (1 + 8), BitsWritten: 7}
	_fde[320] = code{Code: 54 << 8, BitsWritten: 8}
	_fde[384] = code{Code: 55 << 8, BitsWritten: 8}
	_fde[448] = code{Code: 100 << 8, BitsWritten: 8}
	_fde[512] = code{Code: 101 << 8, BitsWritten: 8}
	_fde[576] = code{Code: 104 << 8, BitsWritten: 8}
	_fde[640] = code{Code: 103 << 8, BitsWritten: 8}
	_fde[704] = code{Code: 102 << 8, BitsWritten: 9}
	_fde[768] = code{Code: 102<<8 | 1<<7, BitsWritten: 9}
	_fde[832] = code{Code: 105 << 8, BitsWritten: 9}
	_fde[896] = code{Code: 105<<8 | 1<<7, BitsWritten: 9}
	_fde[960] = code{Code: 106 << 8, BitsWritten: 9}
	_fde[1024] = code{Code: 106<<8 | 1<<7, BitsWritten: 9}
	_fde[1088] = code{Code: 107 << 8, BitsWritten: 9}
	_fde[1152] = code{Code: 107<<8 | 1<<7, BitsWritten: 9}
	_fde[1216] = code{Code: 108 << 8, BitsWritten: 9}
	_fde[1280] = code{Code: 108<<8 | 1<<7, BitsWritten: 9}
	_fde[1344] = code{Code: 109 << 8, BitsWritten: 9}
	_fde[1408] = code{Code: 109<<8 | 1<<7, BitsWritten: 9}
	_fde[1472] = code{Code: 76 << 8, BitsWritten: 9}
	_fde[1536] = code{Code: 76<<8 | 1<<7, BitsWritten: 9}
	_fde[1600] = code{Code: 77 << 8, BitsWritten: 9}
	_fde[1664] = code{Code: 24 << (2 + 8), BitsWritten: 6}
	_fde[1728] = code{Code: 77<<8 | 1<<7, BitsWritten: 9}
	_edgc = make(map[int]code)
	_edgc[1792] = code{Code: 1 << 8, BitsWritten: 11}
	_edgc[1856] = code{Code: 1<<8 | 4<<5, BitsWritten: 11}
	_edgc[1920] = code{Code: 1<<8 | 5<<5, BitsWritten: 11}
	_edgc[1984] = code{Code: 1<<8 | 2<<4, BitsWritten: 12}
	_edgc[2048] = code{Code: 1<<8 | 3<<4, BitsWritten: 12}
	_edgc[2112] = code{Code: 1<<8 | 4<<4, BitsWritten: 12}
	_edgc[2176] = code{Code: 1<<8 | 5<<4, BitsWritten: 12}
	_edgc[2240] = code{Code: 1<<8 | 6<<4, BitsWritten: 12}
	_edgc[2304] = code{Code: 1<<8 | 7<<4, BitsWritten: 12}
	_edgc[2368] = code{Code: 1<<8 | 12<<4, BitsWritten: 12}
	_edgc[2432] = code{Code: 1<<8 | 13<<4, BitsWritten: 12}
	_edgc[2496] = code{Code: 1<<8 | 14<<4, BitsWritten: 12}
	_edgc[2560] = code{Code: 1<<8 | 15<<4, BitsWritten: 12}
	_ecd = make(map[int]byte)
	_ecd[0] = 0xFF
	_ecd[1] = 0xFE
	_ecd[2] = 0xFC
	_ecd[3] = 0xF8
	_ecd[4] = 0xF0
	_ecd[5] = 0xE0
	_ecd[6] = 0xC0
	_ecd[7] = 0x80
	_ecd[8] = 0x00
}
func _bgc(_cdg, _gdf []byte, _baa int) int {
	_fcd := _dgg(_gdf, _baa)
	if _fcd < len(_gdf) && (_baa == -1 && _gdf[_fcd] == _bdec || _baa >= 0 && _baa < len(_cdg) && _cdg[_baa] == _gdf[_fcd] || _baa >= len(_cdg) && _cdg[_baa-1] != _gdf[_fcd]) {
		_fcd = _dgg(_gdf, _fcd)
	}
	return _fcd
}
func (_ffa *Encoder) encodeG4(_ddc [][]byte) []byte {
	_babc := make([][]byte, len(_ddc))
	copy(_babc, _ddc)
	_babc = _egf(_babc)
	var _babg []byte
	var _cgc int
	for _daa := 1; _daa < len(_babc); _daa++ {
		if _ffa.Rows > 0 && !_ffa.EndOfBlock && _daa == (_ffa.Rows+1) {
			break
		}
		var _acg []byte
		var _beb, _ggb, _aeb int
		_gcaf := _cgc
		_ggc := -1
		for _ggc < len(_babc[_daa]) {
			_beb = _dgg(_babc[_daa], _ggc)
			_ggb = _bgc(_babc[_daa], _babc[_daa-1], _ggc)
			_aeb = _dgg(_babc[_daa-1], _ggb)
			if _aeb < _beb {
				_acg, _gcaf = _gec(_acg, _gcaf, _aa)
				_ggc = _aeb
			} else {
				if _e.Abs(float64(_ggb-_beb)) > 3 {
					_acg, _gcaf, _ggc = _dfe(_babc[_daa], _acg, _gcaf, _ggc, _beb)
				} else {
					_acg, _gcaf = _dabd(_acg, _gcaf, _beb, _ggb)
					_ggc = _beb
				}
			}
		}
		_babg = _ffa.appendEncodedRow(_babg, _acg, _cgc)
		if _ffa.EncodedByteAlign {
			_gcaf = 0
		}
		_cgc = _gcaf % 8
	}
	if _ffa.EndOfBlock {
		_ged, _ := _dcf(_cgc)
		_babg = _ffa.appendEncodedRow(_babg, _ged, _cgc)
	}
	return _babg
}
func _gfa(_acf []byte, _eeg int, _gfc code) ([]byte, int) {
	_abe := true
	var _cdeee []byte
	_cdeee, _eeg = _gec(nil, _eeg, _gfc)
	_cdb := 0
	var _eec int
	for _cdb < len(_acf) {
		_eec, _cdb = _gab(_acf, _abe, _cdb)
		_cdeee, _eeg = _caad(_cdeee, _eeg, _eec, _abe)
		_abe = !_abe
	}
	return _cdeee, _eeg % 8
}

var (
	_bdec byte = 1
	_edf  byte = 0
)

type tree struct{ _fcgf *treeNode }

func (_bbc *Decoder) decode1D() error {
	var (
		_ebg int
		_ddg error
	)
	_ebgd := true
	_bbc._bfca = 0
	for {
		var _bdd int
		if _ebgd {
			_bdd, _ddg = _bbc.decodeRun(_c)
		} else {
			_bdd, _ddg = _bbc.decodeRun(_d)
		}
		if _ddg != nil {
			return _ddg
		}
		_ebg += _bdd
		_bbc._bac[_bbc._bfca] = _ebg
		_bbc._bfca++
		_ebgd = !_ebgd
		if _ebg >= _bbc._aeg {
			break
		}
	}
	return nil
}
func _caad(_cgd []byte, _cgcc int, _babb int, _bbbda bool) ([]byte, int) {
	var (
		_baeg code
		_gff  bool
	)
	for !_gff {
		_baeg, _babb, _gff = _cdbd(_babb, _bbbda)
		_cgd, _cgcc = _gec(_cgd, _cgcc, _baeg)
	}
	return _cgd, _cgcc
}
func (_dcg *Decoder) decodeRowType4() error {
	if !_dcg._fdcd {
		return _dcg.decoderRowType41D()
	}
	if _dcg._gfb {
		_dcg._cd.Align()
	}
	_dcg._cd.Mark()
	_ffc, _cffd := _dcg.tryFetchEOL()
	if _cffd != nil {
		return _cffd
	}
	if !_ffc && _dcg._ccf {
		_dcg._bggf++
		if _dcg._bggf > _dcg._gd {
			return _fa
		}
		_dcg._cd.Reset()
	}
	if !_ffc {
		_dcg._cd.Reset()
	}
	_abd, _cffd := _dcg._cd.ReadBool()
	if _cffd != nil {
		return _cffd
	}
	if _abd {
		if _ffc && _dcg._de {
			if _cffd = _dcg.tryFetchRTC2D(); _cffd != nil {
				return _cffd
			}
		}
		_cffd = _dcg.decode1D()
	} else {
		_cffd = _dcg.decode2D()
	}
	if _cffd != nil {
		return _cffd
	}
	return nil
}
func (_bfb *Decoder) Read(in []byte) (int, error) {
	if _bfb._aad != nil {
		return 0, _bfb._aad
	}
	_bggd := len(in)
	var (
		_bcf int
		_caf int
	)
	for _bggd != 0 {
		if _bfb._geg >= _bfb._bc {
			if _dbc := _bfb.fetch(); _dbc != nil {
				_bfb._aad = _dbc
				return 0, _dbc
			}
		}
		if _bfb._bc == -1 {
			return _bcf, _ge.EOF
		}
		switch {
		case _bggd <= _bfb._bc-_bfb._geg:
			_fce := _bfb._dgc[_bfb._geg : _bfb._geg+_bggd]
			for _, _deg := range _fce {
				if !_bfb._bbf {
					_deg = ^_deg
				}
				in[_caf] = _deg
				_caf++
			}
			_bcf += len(_fce)
			_bfb._geg += len(_fce)
			return _bcf, nil
		default:
			_dgf := _bfb._dgc[_bfb._geg:]
			for _, _gge := range _dgf {
				if !_bfb._bbf {
					_gge = ^_gge
				}
				in[_caf] = _gge
				_caf++
			}
			_bcf += len(_dgf)
			_bfb._geg += len(_dgf)
			_bggd -= len(_dgf)
		}
	}
	return _bcf, nil
}

var (
	_gee *treeNode
	_ae  *treeNode
	_d   *tree
	_c   *tree
	_b   *tree
	_ee  *tree
	_ab  = -2000
	_ed  = -1000
	_f   = -3000
	_gb  = -4000
)

func (_fdg *Decoder) decode2D() error {
	_fdg._bd = _fdg._bfca
	_fdg._bac, _fdg._df = _fdg._df, _fdg._bac
	_dga := true
	var (
		_eg   bool
		_fefe int
		_aba  error
	)
	_fdg._bfca = 0
_eag:
	for _fefe < _fdg._aeg {
		_bacb := _ee._fcgf
		for {
			_eg, _aba = _fdg._cd.ReadBool()
			if _aba != nil {
				return _aba
			}
			_bacb = _bacb.walk(_eg)
			if _bacb == nil {
				continue _eag
			}
			if !_bacb._cag {
				continue
			}
			switch _bacb._bcbd {
			case _gb:
				var _ecce int
				if _dga {
					_ecce, _aba = _fdg.decodeRun(_c)
				} else {
					_ecce, _aba = _fdg.decodeRun(_d)
				}
				if _aba != nil {
					return _aba
				}
				_fefe += _ecce
				_fdg._bac[_fdg._bfca] = _fefe
				_fdg._bfca++
				if _dga {
					_ecce, _aba = _fdg.decodeRun(_d)
				} else {
					_ecce, _aba = _fdg.decodeRun(_c)
				}
				if _aba != nil {
					return _aba
				}
				_fefe += _ecce
				_fdg._bac[_fdg._bfca] = _fefe
				_fdg._bfca++
			case _f:
				_eba := _fdg.getNextChangingElement(_fefe, _dga) + 1
				if _eba >= _fdg._bd {
					_fefe = _fdg._aeg
				} else {
					_fefe = _fdg._df[_eba]
				}
			default:
				_ebaf := _fdg.getNextChangingElement(_fefe, _dga)
				if _ebaf >= _fdg._bd || _ebaf == -1 {
					_fefe = _fdg._aeg + _bacb._bcbd
				} else {
					_fefe = _fdg._df[_ebaf] + _bacb._bcbd
				}
				_fdg._bac[_fdg._bfca] = _fefe
				_fdg._bfca++
				_dga = !_dga
			}
			continue _eag
		}
	}
	return nil
}

type Encoder struct {
	K                      int
	EndOfLine              bool
	EncodedByteAlign       bool
	Columns                int
	Rows                   int
	EndOfBlock             bool
	BlackIs1               bool
	DamagedRowsBeforeError int
}

func (_ebe *Decoder) getNextChangingElement(_eab int, _fbg bool) int {
	_bae := 0
	if !_fbg {
		_bae = 1
	}
	_aea := int(uint32(_ebe._gbde)&0xFFFFFFFE) + _bae
	if _aea > 2 {
		_aea -= 2
	}
	if _eab == 0 {
		return _aea
	}
	for _eff := _aea; _eff < _ebe._bd; _eff += 2 {
		if _eab < _ebe._df[_eff] {
			_ebe._gbde = _eff
			return _eff
		}
	}
	return -1
}
func _geab(_aff, _gedf []byte, _bfbb int, _aebe bool) int {
	_aegg := _dgg(_gedf, _bfbb)
	if _aegg < len(_gedf) && (_bfbb == -1 && _gedf[_aegg] == _bdec || _bfbb >= 0 && _bfbb < len(_aff) && _aff[_bfbb] == _gedf[_aegg] || _bfbb >= len(_aff) && _aebe && _gedf[_aegg] == _bdec || _bfbb >= len(_aff) && !_aebe && _gedf[_aegg] == _edf) {
		_aegg = _dgg(_gedf, _aegg)
	}
	return _aegg
}
func _dcf(_bad int) ([]byte, int) {
	var _ffe []byte
	for _ceg := 0; _ceg < 2; _ceg++ {
		_ffe, _bad = _gec(_ffe, _bad, _db)
	}
	return _ffe, _bad % 8
}
func (_effe *Encoder) Encode(pixels [][]byte) []byte {
	if _effe.BlackIs1 {
		_bdec = 0
		_edf = 1
	} else {
		_bdec = 1
		_edf = 0
	}
	if _effe.K == 0 {
		return _effe.encodeG31D(pixels)
	}
	if _effe.K > 0 {
		return _effe.encodeG32D(pixels)
	}
	if _effe.K < 0 {
		return _effe.encodeG4(pixels)
	}
	return nil
}
func _ggce(_ebd int) ([]byte, int) {
	var _fcg []byte
	for _gcc := 0; _gcc < 6; _gcc++ {
		_fcg, _ebd = _gec(_fcg, _ebd, _dg)
	}
	return _fcg, _ebd % 8
}
func _dfe(_effef, _acgb []byte, _edfb, _gggc, _bcd int) ([]byte, int, int) {
	_bdee := _dgg(_effef, _bcd)
	_egd := _gggc >= 0 && _effef[_gggc] == _bdec || _gggc == -1
	_acgb, _edfb = _gec(_acgb, _edfb, _bed)
	var _dcde int
	if _gggc > -1 {
		_dcde = _bcd - _gggc
	} else {
		_dcde = _bcd - _gggc - 1
	}
	_acgb, _edfb = _caad(_acgb, _edfb, _dcde, _egd)
	_egd = !_egd
	_dab := _bdee - _bcd
	_acgb, _edfb = _caad(_acgb, _edfb, _dab, _egd)
	_gggc = _bdee
	return _acgb, _edfb, _gggc
}
func NewDecoder(data []byte, options DecodeOptions) (*Decoder, error) {
	_ded := &Decoder{_cd: _gf.NewReader(data), _aeg: options.Columns, _da: options.Rows, _gd: options.DamagedRowsBeforeError, _dgc: make([]byte, (options.Columns+7)/8), _df: make([]int, options.Columns+2), _bac: make([]int, options.Columns+2), _gfb: options.EncodedByteAligned, _bbf: options.BlackIsOne, _ccf: options.EndOfLine, _de: options.EndOfBlock}
	switch {
	case options.K == 0:
		_ded._fff = _cfd
		if len(data) < 20 {
			return nil, _a.New("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m")
		}
		_bag := data[:20]
		if _bag[0] != 0 || (_bag[1]>>4 != 1 && _bag[1] != 1) {
			_ded._fff = _gef
			_edgce := (uint16(_bag[0])<<8 + uint16(_bag[1]&0xff)) >> 4
			for _ecc := 12; _ecc < 160; _ecc++ {
				_edgce = (_edgce << 1) + uint16((_bag[_ecc/8]>>uint16(7-(_ecc%8)))&0x01)
				if _edgce&0xfff == 1 {
					_ded._fff = _cfd
					break
				}
			}
		}
	case options.K < 0:
		_ded._fff = _eb
	case options.K > 0:
		_ded._fff = _cfd
		_ded._fdcd = true
	}
	switch _ded._fff {
	case _gef, _cfd, _eb:
	default:
		return nil, _a.New("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065")
	}
	return _ded, nil
}
func _ada(_beg []byte, _abc int) ([]byte, int) { return _gec(_beg, _abc, _aa) }
func _cdbd(_gea int, _ecca bool) (code, int, bool) {
	if _gea < 64 {
		if _ecca {
			return _ad[_gea], 0, true
		}
		return _be[_gea], 0, true
	}
	_edcc := _gea / 64
	if _edcc > 40 {
		return _edgc[2560], _gea - 2560, false
	}
	if _edcc > 27 {
		return _edgc[_edcc*64], _gea - _edcc*64, false
	}
	if _ecca {
		return _fde[_edcc*64], _gea - _edcc*64, false
	}
	return _cfe[_edcc*64], _gea - _edcc*64, false
}
func (_bef *Encoder) encodeG32D(_ggg [][]byte) []byte {
	var _dfa []byte
	var _efb int
	for _fdee := 0; _fdee < len(_ggg); _fdee += _bef.K {
		if _bef.Rows > 0 && !_bef.EndOfBlock && _fdee == _bef.Rows {
			break
		}
		_cac, _deee := _gfa(_ggg[_fdee], _efb, _dg)
		_dfa = _bef.appendEncodedRow(_dfa, _cac, _efb)
		if _bef.EncodedByteAlign {
			_deee = 0
		}
		_efb = _deee
		for _ade := _fdee + 1; _ade < (_fdee+_bef.K) && _ade < len(_ggg); _ade++ {
			if _bef.Rows > 0 && !_bef.EndOfBlock && _ade == _bef.Rows {
				break
			}
			_dfac, _fgg := _gec(nil, _efb, _bf)
			var _eaf, _fcb, _cfa int
			_dae := -1
			for _dae < len(_ggg[_ade]) {
				_eaf = _dgg(_ggg[_ade], _dae)
				_fcb = _bgc(_ggg[_ade], _ggg[_ade-1], _dae)
				_cfa = _dgg(_ggg[_ade-1], _fcb)
				if _cfa < _eaf {
					_dfac, _fgg = _ada(_dfac, _fgg)
					_dae = _cfa
				} else {
					if _e.Abs(float64(_fcb-_eaf)) > 3 {
						_dfac, _fgg, _dae = _dfe(_ggg[_ade], _dfac, _fgg, _dae, _eaf)
					} else {
						_dfac, _fgg = _dabd(_dfac, _fgg, _eaf, _fcb)
						_dae = _eaf
					}
				}
			}
			_dfa = _bef.appendEncodedRow(_dfa, _dfac, _efb)
			if _bef.EncodedByteAlign {
				_fgg = 0
			}
			_efb = _fgg % 8
		}
	}
	if _bef.EndOfBlock {
		_ece, _ := _ggce(_efb)
		_dfa = _bef.appendEncodedRow(_dfa, _ece, _efb)
	}
	return _dfa
}
func (_fdb *Decoder) tryFetchEOL1() (bool, error) {
	_edb, _aec := _fdb._cd.ReadBits(13)
	if _aec != nil {
		return false, _aec
	}
	return _edb == 0x3, nil
}

var _ecf = [...][]uint16{{0x2, 0x3}, {0x2, 0x3}, {0x2, 0x3}, {0x3}, {0x4, 0x5}, {0x4, 0x5, 0x7}, {0x4, 0x7}, {0x18}, {0x17, 0x18, 0x37, 0x8, 0xf}, {0x17, 0x18, 0x28, 0x37, 0x67, 0x68, 0x6c, 0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f, 0x24, 0x27, 0x28, 0x2b, 0x2c, 0x33, 0x34, 0x35, 0x37, 0x38, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xda, 0xdb}, {0x4a, 0x4b, 0x4c, 0x4d, 0x52, 0x53, 0x54, 0x55, 0x5a, 0x5b, 0x64, 0x65, 0x6c, 0x6d, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77}}
var _bgg = [...][]uint16{{2, 3, 4, 5, 6, 7}, {128, 8, 9, 64, 10, 11}, {192, 1664, 16, 17, 13, 14, 15, 1, 12}, {26, 21, 28, 27, 18, 24, 25, 22, 256, 23, 20, 19}, {33, 34, 35, 36, 37, 38, 31, 32, 29, 53, 54, 39, 40, 41, 42, 43, 44, 30, 61, 62, 63, 0, 320, 384, 45, 59, 60, 46, 49, 50, 51, 52, 55, 56, 57, 58, 448, 512, 640, 576, 47, 48}, {1472, 1536, 1600, 1728, 704, 768, 832, 896, 960, 1024, 1088, 1152, 1216, 1280, 1344, 1408}, {}, {1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560}}

func (_cfb *Decoder) decoderRowType41D() error {
	if _cfb._gfb {
		_cfb._cd.Align()
	}
	_cfb._cd.Mark()
	var (
		_dcc bool
		_fed error
	)
	if _cfb._ccf {
		_dcc, _fed = _cfb.tryFetchEOL()
		if _fed != nil {
			return _fed
		}
		if !_dcc {
			return _fa
		}
	} else {
		_dcc, _fed = _cfb.looseFetchEOL()
		if _fed != nil {
			return _fed
		}
	}
	if !_dcc {
		_cfb._cd.Reset()
	}
	if _dcc && _cfb._de {
		_cfb._cd.Mark()
		for _gcb := 0; _gcb < 5; _gcb++ {
			_dcc, _fed = _cfb.tryFetchEOL()
			if _fed != nil {
				if _a.Is(_fed, _ge.EOF) {
					if _gcb == 0 {
						break
					}
					return _gc
				}
			}
			if _dcc {
				continue
			}
			if _gcb > 0 {
				return _gc
			}
			break
		}
		if _dcc {
			return _ge.EOF
		}
		_cfb._cd.Reset()
	}
	if _fed = _cfb.decode1D(); _fed != nil {
		return _fed
	}
	return nil
}

type tiffType int
type DecodeOptions struct {
	Columns                int
	Rows                   int
	K                      int
	EncodedByteAligned     bool
	BlackIsOne             bool
	EndOfBlock             bool
	EndOfLine              bool
	DamagedRowsBeforeError int
}

func (_dee *Decoder) tryFetchEOL() (bool, error) {
	_cfc, _gca := _dee._cd.ReadBits(12)
	if _gca != nil {
		return false, _gca
	}
	return _cfc == 0x1, nil
}
func _egf(_baea [][]byte) [][]byte {
	_bgcb := make([]byte, len(_baea[0]))
	for _efbf := range _bgcb {
		_bgcb[_efbf] = _bdec
	}
	_baea = append(_baea, []byte{})
	for _bdfb := len(_baea) - 1; _bdfb > 0; _bdfb-- {
		_baea[_bdfb] = _baea[_bdfb-1]
	}
	_baea[0] = _bgcb
	return _baea
}
func (_edc *Decoder) decodeRow() (_adgd error) {
	if !_edc._de && _edc._da > 0 && _edc._da == _edc._agg {
		return _ge.EOF
	}
	switch _edc._fff {
	case _gef:
		_adgd = _edc.decodeRowType2()
	case _cfd:
		_adgd = _edc.decodeRowType4()
	case _eb:
		_adgd = _edc.decodeRowType6()
	}
	if _adgd != nil {
		return _adgd
	}
	_bdf := 0
	_afg := true
	_edc._gbde = 0
	for _bbbc := 0; _bbbc < _edc._bfca; _bbbc++ {
		_ggf := _edc._aeg
		if _bbbc != _edc._bfca {
			_ggf = _edc._bac[_bbbc]
		}
		if _ggf > _edc._aeg {
			_ggf = _edc._aeg
		}
		_eef := _bdf / 8
		for _bdf%8 != 0 && _ggf-_bdf > 0 {
			var _cee byte
			if !_afg {
				_cee = 1 << uint(7-(_bdf%8))
			}
			_edc._dgc[_eef] |= _cee
			_bdf++
		}
		if _bdf%8 == 0 {
			_eef = _bdf / 8
			var _ccg byte
			if !_afg {
				_ccg = 0xff
			}
			for _ggf-_bdf > 7 {
				_edc._dgc[_eef] = _ccg
				_bdf += 8
				_eef++
			}
		}
		for _ggf-_bdf > 0 {
			if _bdf%8 == 0 {
				_edc._dgc[_eef] = 0
			}
			var _ga byte
			if !_afg {
				_ga = 1 << uint(7-(_bdf%8))
			}
			_edc._dgc[_eef] |= _ga
			_bdf++
		}
		_afg = !_afg
	}
	if _bdf != _edc._aeg {
		return _a.New("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068")
	}
	_edc._bc = (_bdf + 7) / 8
	_edc._agg++
	return nil
}

type treeNode struct {
	_gbfd *treeNode
	_gbe  *treeNode
	_bcbd int
	_agbg bool
	_cag  bool
}

func _dgg(_dgba []byte, _gfaf int) int {
	if _gfaf >= len(_dgba) {
		return _gfaf
	}
	if _gfaf < -1 {
		_gfaf = -1
	}
	var _bga byte
	if _gfaf > -1 {
		_bga = _dgba[_gfaf]
	} else {
		_bga = _bdec
	}
	_egac := _gfaf + 1
	for _egac < len(_dgba) {
		if _dgba[_egac] != _bga {
			break
		}
		_egac++
	}
	return _egac
}
func (_gfg *Decoder) tryFetchRTC2D() (_eee error) {
	_gfg._cd.Mark()
	var _cdee bool
	for _ega := 0; _ega < 5; _ega++ {
		_cdee, _eee = _gfg.tryFetchEOL1()
		if _eee != nil {
			if _a.Is(_eee, _ge.EOF) {
				if _ega == 0 {
					break
				}
				return _gc
			}
		}
		if _cdee {
			continue
		}
		if _ega > 0 {
			return _gc
		}
		break
	}
	if _cdee {
		return _ge.EOF
	}
	_gfg._cd.Reset()
	return _eee
}
func _fdce(_cgg, _dea int) code {
	var _fdd code
	switch _dea - _cgg {
	case -1:
		_fdd = _adg
	case -2:
		_fdd = _bfc
	case -3:
		_fdd = _ceb
	case 0:
		_fdd = _cff
	case 1:
		_fdd = _cg
	case 2:
		_fdd = _dd
	case 3:
		_fdd = _bab
	}
	return _fdd
}
func (_fgc *Decoder) looseFetchEOL() (bool, error) {
	_ecdb, _gbg := _fgc._cd.ReadBits(12)
	if _gbg != nil {
		return false, _gbg
	}
	switch _ecdb {
	case 0x1:
		return true, nil
	case 0x0:
		for {
			_efc, _dbce := _fgc._cd.ReadBool()
			if _dbce != nil {
				return false, _dbce
			}
			if _efc {
				return true, nil
			}
		}
	default:
		return false, nil
	}
}

var _fg = [...][]uint16{{3, 2}, {1, 4}, {6, 5}, {7}, {9, 8}, {10, 11, 12}, {13, 14}, {15}, {16, 17, 0, 18, 64}, {24, 25, 23, 22, 19, 20, 21, 1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560, 52, 55, 56, 59, 60, 320, 384, 448, 53, 54, 50, 51, 44, 45, 46, 47, 57, 58, 61, 256, 48, 49, 62, 63, 30, 31, 32, 33, 40, 41, 128, 192, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 42, 43}, {640, 704, 768, 832, 1280, 1344, 1408, 1472, 1536, 1600, 1664, 1728, 512, 576, 896, 960, 1024, 1088, 1152, 1216}}

type Decoder struct {
	_aeg  int
	_da   int
	_agg  int
	_dgc  []byte
	_gd   int
	_fdcd bool
	_gbd  bool
	_ca   bool
	_bbf  bool
	_ccf  bool
	_de   bool
	_gfb  bool
	_bc   int
	_geg  int
	_df   []int
	_bac  []int
	_bd   int
	_bfca int
	_bggf int
	_gbde int
	_cd   *_gf.Reader
	_fff  tiffType
	_aad  error
}

func (_cec *treeNode) set(_ggge bool, _dfb *treeNode) {
	if !_ggge {
		_cec._gbfd = _dfb
	} else {
		_cec._gbe = _dfb
	}
}

var (
	_gc = _a.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043")
	_fa = _a.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
)

func (_fgd tiffType) String() string {
	switch _fgd {
	case _gef:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065"
	case _cfd:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034"
	case _eb:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036"
	default:
		return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d"
	}
}

var (
	_be   map[int]code
	_ad   map[int]code
	_cfe  map[int]code
	_fde  map[int]code
	_edgc map[int]code
	_ecd  map[int]byte
	_db   = code{Code: 1 << 4, BitsWritten: 12}
	_dg   = code{Code: 3 << 3, BitsWritten: 13}
	_bf   = code{Code: 2 << 3, BitsWritten: 13}
	_aa   = code{Code: 1 << 12, BitsWritten: 4}
	_bed  = code{Code: 1 << 13, BitsWritten: 3}
	_cff  = code{Code: 1 << 15, BitsWritten: 1}
	_adg  = code{Code: 3 << 13, BitsWritten: 3}
	_bfc  = code{Code: 3 << 10, BitsWritten: 6}
	_ceb  = code{Code: 3 << 9, BitsWritten: 7}
	_cg   = code{Code: 2 << 13, BitsWritten: 3}
	_dd   = code{Code: 2 << 10, BitsWritten: 6}
	_bab  = code{Code: 2 << 9, BitsWritten: 7}
)

func _dabd(_aca []byte, _cga, _eda, _ggcc int) ([]byte, int) {
	_fba := _fdce(_eda, _ggcc)
	_aca, _cga = _gec(_aca, _cga, _fba)
	return _aca, _cga
}
func (_ffd *tree) fillWithNode(_aece, _fcbb int, _baeae *treeNode) error {
	_eccc := _ffd._fcgf
	for _adc := 0; _adc < _aece; _adc++ {
		_eage := uint(_aece - 1 - _adc)
		_bfg := ((_fcbb >> _eage) & 1) != 0
		_dad := _eccc.walk(_bfg)
		if _dad != nil {
			if _dad._cag {
				return _a.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_eccc = _dad
			continue
		}
		if _adc == _aece-1 {
			_dad = _baeae
		} else {
			_dad = &treeNode{}
		}
		if _fcbb == 0 {
			_dad._agbg = true
		}
		_eccc.set(_bfg, _dad)
		_eccc = _dad
	}
	return nil
}
func (_bdc *Decoder) decodeG32D() error {
	_bdc._bd = _bdc._bfca
	_bdc._bac, _bdc._df = _bdc._df, _bdc._bac
	_aega := true
	var (
		_cde bool
		_ecb int
		_ef  error
	)
	_bdc._bfca = 0
_fef:
	for _ecb < _bdc._aeg {
		_caa := _ee._fcgf
		for {
			_cde, _ef = _bdc._cd.ReadBool()
			if _ef != nil {
				return _ef
			}
			_caa = _caa.walk(_cde)
			if _caa == nil {
				continue _fef
			}
			if !_caa._cag {
				continue
			}
			switch _caa._bcbd {
			case _gb:
				var _dcd int
				if _aega {
					_dcd, _ef = _bdc.decodeRun(_c)
				} else {
					_dcd, _ef = _bdc.decodeRun(_d)
				}
				if _ef != nil {
					return _ef
				}
				_ecb += _dcd
				_bdc._bac[_bdc._bfca] = _ecb
				_bdc._bfca++
				if _aega {
					_dcd, _ef = _bdc.decodeRun(_d)
				} else {
					_dcd, _ef = _bdc.decodeRun(_c)
				}
				if _ef != nil {
					return _ef
				}
				_ecb += _dcd
				_bdc._bac[_bdc._bfca] = _ecb
				_bdc._bfca++
			case _f:
				_bce := _bdc.getNextChangingElement(_ecb, _aega) + 1
				if _bce >= _bdc._bd {
					_ecb = _bdc._aeg
				} else {
					_ecb = _bdc._df[_bce]
				}
			default:
				_cab := _bdc.getNextChangingElement(_ecb, _aega)
				if _cab >= _bdc._bd || _cab == -1 {
					_ecb = _bdc._aeg + _caa._bcbd
				} else {
					_ecb = _bdc._df[_cab] + _caa._bcbd
				}
				_bdc._bac[_bdc._bfca] = _ecb
				_bdc._bfca++
				_aega = !_aega
			}
			continue _fef
		}
	}
	return nil
}
func (_cda *treeNode) walk(_bec bool) *treeNode {
	if _bec {
		return _cda._gbe
	}
	return _cda._gbfd
}
func _gab(_afeb []byte, _agb bool, _dgb int) (int, int) {
	_bfd := 0
	for _dgb < len(_afeb) {
		if _agb {
			if _afeb[_dgb] != _bdec {
				break
			}
		} else {
			if _afeb[_dgb] != _edf {
				break
			}
		}
		_bfd++
		_dgb++
	}
	return _bfd, _dgb
}
func _gdc(_bbbd int) ([]byte, int) {
	var _cdc []byte
	for _ddgc := 0; _ddgc < 6; _ddgc++ {
		_cdc, _bbbd = _gec(_cdc, _bbbd, _db)
	}
	return _cdc, _bbbd % 8
}
func (_gfe *Decoder) fetch() error {
	if _gfe._bc == -1 {
		return nil
	}
	if _gfe._geg < _gfe._bc {
		return nil
	}
	_gfe._bc = 0
	_aef := _gfe.decodeRow()
	if _aef != nil {
		if !_a.Is(_aef, _ge.EOF) {
			return _aef
		}
		if _gfe._bc != 0 {
			return _aef
		}
		_gfe._bc = -1
	}
	_gfe._geg = 0
	return nil
}
func (_ccfe *Decoder) decodeRun(_acd *tree) (int, error) {
	var _bbd int
	_ebef := _acd._fcgf
	for {
		_aac, _bde := _ccfe._cd.ReadBool()
		if _bde != nil {
			return 0, _bde
		}
		_ebef = _ebef.walk(_aac)
		if _ebef == nil {
			return 0, _a.New("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _ebef._cag {
			_bbd += _ebef._bcbd
			switch {
			case _ebef._bcbd >= 64:
				_ebef = _acd._fcgf
			case _ebef._bcbd >= 0:
				return _bbd, nil
			default:
				return _ccfe._aeg, nil
			}
		}
	}
}
func (_edae *tree) fill(_bace, _eeb, _fffd int) error {
	_eafg := _edae._fcgf
	for _dggc := 0; _dggc < _bace; _dggc++ {
		_fgb := _bace - 1 - _dggc
		_ccff := ((_eeb >> uint(_fgb)) & 1) != 0
		_aefc := _eafg.walk(_ccff)
		if _aefc != nil {
			if _aefc._cag {
				return _a.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_eafg = _aefc
			continue
		}
		_aefc = &treeNode{}
		if _dggc == _bace-1 {
			_aefc._bcbd = _fffd
			_aefc._cag = true
		}
		if _eeb == 0 {
			_aefc._agbg = true
		}
		_eafg.set(_ccff, _aefc)
		_eafg = _aefc
	}
	return nil
}
func (_gbf *Encoder) encodeG31D(_afdd [][]byte) []byte {
	var _gcg []byte
	_fefa := 0
	for _dde := range _afdd {
		if _gbf.Rows > 0 && !_gbf.EndOfBlock && _dde == _gbf.Rows {
			break
		}
		_ffg, _bddd := _gfa(_afdd[_dde], _fefa, _db)
		_gcg = _gbf.appendEncodedRow(_gcg, _ffg, _fefa)
		if _gbf.EncodedByteAlign {
			_bddd = 0
		}
		_fefa = _bddd
	}
	if _gbf.EndOfBlock {
		_bcb, _ := _gdc(_fefa)
		_gcg = _gbf.appendEncodedRow(_gcg, _bcb, _fefa)
	}
	return _gcg
}
