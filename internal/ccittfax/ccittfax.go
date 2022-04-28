package ccittfax

import (
	_g "errors"
	_e "io"
	_b "math"

	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
)

func _cgc(_edad []byte, _ggeg int) ([]byte, int) { return _acbf(_edad, _ggeg, _fde) }
func init() {
	_bg = &treeNode{_dfb: true, _fffag: _ab}
	_ff = &treeNode{_fffag: _ag, _efefd: _bg}
	_ff._ebf = _ff
	_a = &tree{_fgea: &treeNode{}}
	if _gb := _a.fillWithNode(12, 0, _ff); _gb != nil {
		panic(_gb.Error())
	}
	if _gfa := _a.fillWithNode(12, 1, _bg); _gfa != nil {
		panic(_gfa.Error())
	}
	_gf = &tree{_fgea: &treeNode{}}
	for _fg := 0; _fg < len(_dde); _fg++ {
		for _c := 0; _c < len(_dde[_fg]); _c++ {
			if _cc := _gf.fill(_fg+2, int(_dde[_fg][_c]), int(_ge[_fg][_c])); _cc != nil {
				panic(_cc.Error())
			}
		}
	}
	if _eb := _gf.fillWithNode(12, 0, _ff); _eb != nil {
		panic(_eb.Error())
	}
	if _ea := _gf.fillWithNode(12, 1, _bg); _ea != nil {
		panic(_ea.Error())
	}
	_da = &tree{_fgea: &treeNode{}}
	for _aa := 0; _aa < len(_bge); _aa++ {
		for _bc := 0; _bc < len(_bge[_aa]); _bc++ {
			if _ca := _da.fill(_aa+4, int(_bge[_aa][_bc]), int(_cf[_aa][_bc])); _ca != nil {
				panic(_ca.Error())
			}
		}
	}
	if _dd := _da.fillWithNode(12, 0, _ff); _dd != nil {
		panic(_dd.Error())
	}
	if _fe := _da.fillWithNode(12, 1, _bg); _fe != nil {
		panic(_fe.Error())
	}
	_eg = &tree{_fgea: &treeNode{}}
	if _fa := _eg.fill(4, 1, _de); _fa != nil {
		panic(_fa.Error())
	}
	if _ege := _eg.fill(3, 1, _ef); _ege != nil {
		panic(_ege.Error())
	}
	if _bf := _eg.fill(1, 1, 0); _bf != nil {
		panic(_bf.Error())
	}
	if _bgc := _eg.fill(3, 3, 1); _bgc != nil {
		panic(_bgc.Error())
	}
	if _fd := _eg.fill(6, 3, 2); _fd != nil {
		panic(_fd.Error())
	}
	if _ebg := _eg.fill(7, 3, 3); _ebg != nil {
		panic(_ebg.Error())
	}
	if _df := _eg.fill(3, 2, -1); _df != nil {
		panic(_df.Error())
	}
	if _db := _eg.fill(6, 2, -2); _db != nil {
		panic(_db.Error())
	}
	if _af := _eg.fill(7, 2, -3); _af != nil {
		panic(_af.Error())
	}
}
func (_bfe *tree) fillWithNode(_cfe, _bcea int, _bdgb *treeNode) error {
	_ccff := _bfe._fgea
	for _bfdf := 0; _bfdf < _cfe; _bfdf++ {
		_abc := uint(_cfe - 1 - _bfdf)
		_bbe := ((_bcea >> _abc) & 1) != 0
		_cae := _ccff.walk(_bbe)
		if _cae != nil {
			if _cae._dfb {
				return _g.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_ccff = _cae
			continue
		}
		if _bfdf == _cfe-1 {
			_cae = _bdgb
		} else {
			_cae = &treeNode{}
		}
		if _bcea == 0 {
			_cae._gfda = true
		}
		_ccff.set(_bbe, _cae)
		_ccff = _cae
	}
	return nil
}

var _dde = [...][]uint16{{0x2, 0x3}, {0x2, 0x3}, {0x2, 0x3}, {0x3}, {0x4, 0x5}, {0x4, 0x5, 0x7}, {0x4, 0x7}, {0x18}, {0x17, 0x18, 0x37, 0x8, 0xf}, {0x17, 0x18, 0x28, 0x37, 0x67, 0x68, 0x6c, 0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f, 0x24, 0x27, 0x28, 0x2b, 0x2c, 0x33, 0x34, 0x35, 0x37, 0x38, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xda, 0xdb}, {0x4a, 0x4b, 0x4c, 0x4d, 0x52, 0x53, 0x54, 0x55, 0x5a, 0x5b, 0x64, 0x65, 0x6c, 0x6d, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77}}

func NewDecoder(data []byte, options DecodeOptions) (*Decoder, error) {
	_cac := &Decoder{_dag: _f.NewReader(data), _afc: options.Columns, _gcd: options.Rows, _bdf: options.DamagedRowsBeforeError, _afg: make([]byte, (options.Columns+7)/8), _dfea: make([]int, options.Columns+2), _gac: make([]int, options.Columns+2), _eac: options.EncodedByteAligned, _faf: options.BlackIsOne, _gg: options.EndOfLine, _bcb: options.EndOfBlock}
	switch {
	case options.K == 0:
		_cac._cafe = _ba
		if len(data) < 20 {
			return nil, _g.New("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m")
		}
		_ac := data[:20]
		if _ac[0] != 0 || (_ac[1]>>4 != 1 && _ac[1] != 1) {
			_cac._cafe = _eeg
			_aeb := (uint16(_ac[0])<<8 + uint16(_ac[1]&0xff)) >> 4
			for _daa := 12; _daa < 160; _daa++ {
				_aeb = (_aeb << 1) + uint16((_ac[_daa/8]>>uint16(7-(_daa%8)))&0x01)
				if _aeb&0xfff == 1 {
					_cac._cafe = _ba
					break
				}
			}
		}
	case options.K < 0:
		_cac._cafe = _fb
	case options.K > 0:
		_cac._cafe = _ba
		_cac._abg = true
	}
	switch _cac._cafe {
	case _eeg, _ba, _fb:
	default:
		return nil, _g.New("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065")
	}
	return _cac, nil
}
func (_gde *Decoder) decodeG32D() error {
	_gde._fbc = _gde._bdc
	_gde._gac, _gde._dfea = _gde._dfea, _gde._gac
	_be := true
	var (
		_gff bool
		_eec int
		_deb error
	)
	_gde._bdc = 0
_fga:
	for _eec < _gde._afc {
		_eab := _eg._fgea
		for {
			_gff, _deb = _gde._dag.ReadBool()
			if _deb != nil {
				return _deb
			}
			_eab = _eab.walk(_gff)
			if _eab == nil {
				continue _fga
			}
			if !_eab._dfb {
				continue
			}
			switch _eab._fffag {
			case _ef:
				var _ecff int
				if _be {
					_ecff, _deb = _gde.decodeRun(_da)
				} else {
					_ecff, _deb = _gde.decodeRun(_gf)
				}
				if _deb != nil {
					return _deb
				}
				_eec += _ecff
				_gde._gac[_gde._bdc] = _eec
				_gde._bdc++
				if _be {
					_ecff, _deb = _gde.decodeRun(_gf)
				} else {
					_ecff, _deb = _gde.decodeRun(_da)
				}
				if _deb != nil {
					return _deb
				}
				_eec += _ecff
				_gde._gac[_gde._bdc] = _eec
				_gde._bdc++
			case _de:
				_gab := _gde.getNextChangingElement(_eec, _be) + 1
				if _gab >= _gde._fbc {
					_eec = _gde._afc
				} else {
					_eec = _gde._dfea[_gab]
				}
			default:
				_edd := _gde.getNextChangingElement(_eec, _be)
				if _edd >= _gde._fbc || _edd == -1 {
					_eec = _gde._afc + _eab._fffag
				} else {
					_eec = _gde._dfea[_edd] + _eab._fffag
				}
				_gde._gac[_gde._bdc] = _eec
				_gde._bdc++
				_be = !_be
			}
			continue _fga
		}
	}
	return nil
}

var _cf = [...][]uint16{{2, 3, 4, 5, 6, 7}, {128, 8, 9, 64, 10, 11}, {192, 1664, 16, 17, 13, 14, 15, 1, 12}, {26, 21, 28, 27, 18, 24, 25, 22, 256, 23, 20, 19}, {33, 34, 35, 36, 37, 38, 31, 32, 29, 53, 54, 39, 40, 41, 42, 43, 44, 30, 61, 62, 63, 0, 320, 384, 45, 59, 60, 46, 49, 50, 51, 52, 55, 56, 57, 58, 448, 512, 640, 576, 47, 48}, {1472, 1536, 1600, 1728, 704, 768, 832, 896, 960, 1024, 1088, 1152, 1216, 1280, 1344, 1408}, {}, {1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560}}

const (
	_ tiffType = iota
	_eeg
	_ba
	_fb
)

func init() {
	_bb = make(map[int]code)
	_bb[0] = code{Code: 13<<8 | 3<<6, BitsWritten: 10}
	_bb[1] = code{Code: 2 << (5 + 8), BitsWritten: 3}
	_bb[2] = code{Code: 3 << (6 + 8), BitsWritten: 2}
	_bb[3] = code{Code: 2 << (6 + 8), BitsWritten: 2}
	_bb[4] = code{Code: 3 << (5 + 8), BitsWritten: 3}
	_bb[5] = code{Code: 3 << (4 + 8), BitsWritten: 4}
	_bb[6] = code{Code: 2 << (4 + 8), BitsWritten: 4}
	_bb[7] = code{Code: 3 << (3 + 8), BitsWritten: 5}
	_bb[8] = code{Code: 5 << (2 + 8), BitsWritten: 6}
	_bb[9] = code{Code: 4 << (2 + 8), BitsWritten: 6}
	_bb[10] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_bb[11] = code{Code: 5 << (1 + 8), BitsWritten: 7}
	_bb[12] = code{Code: 7 << (1 + 8), BitsWritten: 7}
	_bb[13] = code{Code: 4 << 8, BitsWritten: 8}
	_bb[14] = code{Code: 7 << 8, BitsWritten: 8}
	_bb[15] = code{Code: 12 << 8, BitsWritten: 9}
	_bb[16] = code{Code: 5<<8 | 3<<6, BitsWritten: 10}
	_bb[17] = code{Code: 6 << 8, BitsWritten: 10}
	_bb[18] = code{Code: 2 << 8, BitsWritten: 10}
	_bb[19] = code{Code: 12<<8 | 7<<5, BitsWritten: 11}
	_bb[20] = code{Code: 13 << 8, BitsWritten: 11}
	_bb[21] = code{Code: 13<<8 | 4<<5, BitsWritten: 11}
	_bb[22] = code{Code: 6<<8 | 7<<5, BitsWritten: 11}
	_bb[23] = code{Code: 5 << 8, BitsWritten: 11}
	_bb[24] = code{Code: 2<<8 | 7<<5, BitsWritten: 11}
	_bb[25] = code{Code: 3 << 8, BitsWritten: 11}
	_bb[26] = code{Code: 12<<8 | 10<<4, BitsWritten: 12}
	_bb[27] = code{Code: 12<<8 | 11<<4, BitsWritten: 12}
	_bb[28] = code{Code: 12<<8 | 12<<4, BitsWritten: 12}
	_bb[29] = code{Code: 12<<8 | 13<<4, BitsWritten: 12}
	_bb[30] = code{Code: 6<<8 | 8<<4, BitsWritten: 12}
	_bb[31] = code{Code: 6<<8 | 9<<4, BitsWritten: 12}
	_bb[32] = code{Code: 6<<8 | 10<<4, BitsWritten: 12}
	_bb[33] = code{Code: 6<<8 | 11<<4, BitsWritten: 12}
	_bb[34] = code{Code: 13<<8 | 2<<4, BitsWritten: 12}
	_bb[35] = code{Code: 13<<8 | 3<<4, BitsWritten: 12}
	_bb[36] = code{Code: 13<<8 | 4<<4, BitsWritten: 12}
	_bb[37] = code{Code: 13<<8 | 5<<4, BitsWritten: 12}
	_bb[38] = code{Code: 13<<8 | 6<<4, BitsWritten: 12}
	_bb[39] = code{Code: 13<<8 | 7<<4, BitsWritten: 12}
	_bb[40] = code{Code: 6<<8 | 12<<4, BitsWritten: 12}
	_bb[41] = code{Code: 6<<8 | 13<<4, BitsWritten: 12}
	_bb[42] = code{Code: 13<<8 | 10<<4, BitsWritten: 12}
	_bb[43] = code{Code: 13<<8 | 11<<4, BitsWritten: 12}
	_bb[44] = code{Code: 5<<8 | 4<<4, BitsWritten: 12}
	_bb[45] = code{Code: 5<<8 | 5<<4, BitsWritten: 12}
	_bb[46] = code{Code: 5<<8 | 6<<4, BitsWritten: 12}
	_bb[47] = code{Code: 5<<8 | 7<<4, BitsWritten: 12}
	_bb[48] = code{Code: 6<<8 | 4<<4, BitsWritten: 12}
	_bb[49] = code{Code: 6<<8 | 5<<4, BitsWritten: 12}
	_bb[50] = code{Code: 5<<8 | 2<<4, BitsWritten: 12}
	_bb[51] = code{Code: 5<<8 | 3<<4, BitsWritten: 12}
	_bb[52] = code{Code: 2<<8 | 4<<4, BitsWritten: 12}
	_bb[53] = code{Code: 3<<8 | 7<<4, BitsWritten: 12}
	_bb[54] = code{Code: 3<<8 | 8<<4, BitsWritten: 12}
	_bb[55] = code{Code: 2<<8 | 7<<4, BitsWritten: 12}
	_bb[56] = code{Code: 2<<8 | 8<<4, BitsWritten: 12}
	_bb[57] = code{Code: 5<<8 | 8<<4, BitsWritten: 12}
	_bb[58] = code{Code: 5<<8 | 9<<4, BitsWritten: 12}
	_bb[59] = code{Code: 2<<8 | 11<<4, BitsWritten: 12}
	_bb[60] = code{Code: 2<<8 | 12<<4, BitsWritten: 12}
	_bb[61] = code{Code: 5<<8 | 10<<4, BitsWritten: 12}
	_bb[62] = code{Code: 6<<8 | 6<<4, BitsWritten: 12}
	_bb[63] = code{Code: 6<<8 | 7<<4, BitsWritten: 12}
	_gbc = make(map[int]code)
	_gbc[0] = code{Code: 53 << 8, BitsWritten: 8}
	_gbc[1] = code{Code: 7 << (2 + 8), BitsWritten: 6}
	_gbc[2] = code{Code: 7 << (4 + 8), BitsWritten: 4}
	_gbc[3] = code{Code: 8 << (4 + 8), BitsWritten: 4}
	_gbc[4] = code{Code: 11 << (4 + 8), BitsWritten: 4}
	_gbc[5] = code{Code: 12 << (4 + 8), BitsWritten: 4}
	_gbc[6] = code{Code: 14 << (4 + 8), BitsWritten: 4}
	_gbc[7] = code{Code: 15 << (4 + 8), BitsWritten: 4}
	_gbc[8] = code{Code: 19 << (3 + 8), BitsWritten: 5}
	_gbc[9] = code{Code: 20 << (3 + 8), BitsWritten: 5}
	_gbc[10] = code{Code: 7 << (3 + 8), BitsWritten: 5}
	_gbc[11] = code{Code: 8 << (3 + 8), BitsWritten: 5}
	_gbc[12] = code{Code: 8 << (2 + 8), BitsWritten: 6}
	_gbc[13] = code{Code: 3 << (2 + 8), BitsWritten: 6}
	_gbc[14] = code{Code: 52 << (2 + 8), BitsWritten: 6}
	_gbc[15] = code{Code: 53 << (2 + 8), BitsWritten: 6}
	_gbc[16] = code{Code: 42 << (2 + 8), BitsWritten: 6}
	_gbc[17] = code{Code: 43 << (2 + 8), BitsWritten: 6}
	_gbc[18] = code{Code: 39 << (1 + 8), BitsWritten: 7}
	_gbc[19] = code{Code: 12 << (1 + 8), BitsWritten: 7}
	_gbc[20] = code{Code: 8 << (1 + 8), BitsWritten: 7}
	_gbc[21] = code{Code: 23 << (1 + 8), BitsWritten: 7}
	_gbc[22] = code{Code: 3 << (1 + 8), BitsWritten: 7}
	_gbc[23] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_gbc[24] = code{Code: 40 << (1 + 8), BitsWritten: 7}
	_gbc[25] = code{Code: 43 << (1 + 8), BitsWritten: 7}
	_gbc[26] = code{Code: 19 << (1 + 8), BitsWritten: 7}
	_gbc[27] = code{Code: 36 << (1 + 8), BitsWritten: 7}
	_gbc[28] = code{Code: 24 << (1 + 8), BitsWritten: 7}
	_gbc[29] = code{Code: 2 << 8, BitsWritten: 8}
	_gbc[30] = code{Code: 3 << 8, BitsWritten: 8}
	_gbc[31] = code{Code: 26 << 8, BitsWritten: 8}
	_gbc[32] = code{Code: 27 << 8, BitsWritten: 8}
	_gbc[33] = code{Code: 18 << 8, BitsWritten: 8}
	_gbc[34] = code{Code: 19 << 8, BitsWritten: 8}
	_gbc[35] = code{Code: 20 << 8, BitsWritten: 8}
	_gbc[36] = code{Code: 21 << 8, BitsWritten: 8}
	_gbc[37] = code{Code: 22 << 8, BitsWritten: 8}
	_gbc[38] = code{Code: 23 << 8, BitsWritten: 8}
	_gbc[39] = code{Code: 40 << 8, BitsWritten: 8}
	_gbc[40] = code{Code: 41 << 8, BitsWritten: 8}
	_gbc[41] = code{Code: 42 << 8, BitsWritten: 8}
	_gbc[42] = code{Code: 43 << 8, BitsWritten: 8}
	_gbc[43] = code{Code: 44 << 8, BitsWritten: 8}
	_gbc[44] = code{Code: 45 << 8, BitsWritten: 8}
	_gbc[45] = code{Code: 4 << 8, BitsWritten: 8}
	_gbc[46] = code{Code: 5 << 8, BitsWritten: 8}
	_gbc[47] = code{Code: 10 << 8, BitsWritten: 8}
	_gbc[48] = code{Code: 11 << 8, BitsWritten: 8}
	_gbc[49] = code{Code: 82 << 8, BitsWritten: 8}
	_gbc[50] = code{Code: 83 << 8, BitsWritten: 8}
	_gbc[51] = code{Code: 84 << 8, BitsWritten: 8}
	_gbc[52] = code{Code: 85 << 8, BitsWritten: 8}
	_gbc[53] = code{Code: 36 << 8, BitsWritten: 8}
	_gbc[54] = code{Code: 37 << 8, BitsWritten: 8}
	_gbc[55] = code{Code: 88 << 8, BitsWritten: 8}
	_gbc[56] = code{Code: 89 << 8, BitsWritten: 8}
	_gbc[57] = code{Code: 90 << 8, BitsWritten: 8}
	_gbc[58] = code{Code: 91 << 8, BitsWritten: 8}
	_gbc[59] = code{Code: 74 << 8, BitsWritten: 8}
	_gbc[60] = code{Code: 75 << 8, BitsWritten: 8}
	_gbc[61] = code{Code: 50 << 8, BitsWritten: 8}
	_gbc[62] = code{Code: 51 << 8, BitsWritten: 8}
	_gbc[63] = code{Code: 52 << 8, BitsWritten: 8}
	_fc = make(map[int]code)
	_fc[64] = code{Code: 3<<8 | 3<<6, BitsWritten: 10}
	_fc[128] = code{Code: 12<<8 | 8<<4, BitsWritten: 12}
	_fc[192] = code{Code: 12<<8 | 9<<4, BitsWritten: 12}
	_fc[256] = code{Code: 5<<8 | 11<<4, BitsWritten: 12}
	_fc[320] = code{Code: 3<<8 | 3<<4, BitsWritten: 12}
	_fc[384] = code{Code: 3<<8 | 4<<4, BitsWritten: 12}
	_fc[448] = code{Code: 3<<8 | 5<<4, BitsWritten: 12}
	_fc[512] = code{Code: 3<<8 | 12<<3, BitsWritten: 13}
	_fc[576] = code{Code: 3<<8 | 13<<3, BitsWritten: 13}
	_fc[640] = code{Code: 2<<8 | 10<<3, BitsWritten: 13}
	_fc[704] = code{Code: 2<<8 | 11<<3, BitsWritten: 13}
	_fc[768] = code{Code: 2<<8 | 12<<3, BitsWritten: 13}
	_fc[832] = code{Code: 2<<8 | 13<<3, BitsWritten: 13}
	_fc[896] = code{Code: 3<<8 | 18<<3, BitsWritten: 13}
	_fc[960] = code{Code: 3<<8 | 19<<3, BitsWritten: 13}
	_fc[1024] = code{Code: 3<<8 | 20<<3, BitsWritten: 13}
	_fc[1088] = code{Code: 3<<8 | 21<<3, BitsWritten: 13}
	_fc[1152] = code{Code: 3<<8 | 22<<3, BitsWritten: 13}
	_fc[1216] = code{Code: 119 << 3, BitsWritten: 13}
	_fc[1280] = code{Code: 2<<8 | 18<<3, BitsWritten: 13}
	_fc[1344] = code{Code: 2<<8 | 19<<3, BitsWritten: 13}
	_fc[1408] = code{Code: 2<<8 | 20<<3, BitsWritten: 13}
	_fc[1472] = code{Code: 2<<8 | 21<<3, BitsWritten: 13}
	_fc[1536] = code{Code: 2<<8 | 26<<3, BitsWritten: 13}
	_fc[1600] = code{Code: 2<<8 | 27<<3, BitsWritten: 13}
	_fc[1664] = code{Code: 3<<8 | 4<<3, BitsWritten: 13}
	_fc[1728] = code{Code: 3<<8 | 5<<3, BitsWritten: 13}
	_fdg = make(map[int]code)
	_fdg[64] = code{Code: 27 << (3 + 8), BitsWritten: 5}
	_fdg[128] = code{Code: 18 << (3 + 8), BitsWritten: 5}
	_fdg[192] = code{Code: 23 << (2 + 8), BitsWritten: 6}
	_fdg[256] = code{Code: 55 << (1 + 8), BitsWritten: 7}
	_fdg[320] = code{Code: 54 << 8, BitsWritten: 8}
	_fdg[384] = code{Code: 55 << 8, BitsWritten: 8}
	_fdg[448] = code{Code: 100 << 8, BitsWritten: 8}
	_fdg[512] = code{Code: 101 << 8, BitsWritten: 8}
	_fdg[576] = code{Code: 104 << 8, BitsWritten: 8}
	_fdg[640] = code{Code: 103 << 8, BitsWritten: 8}
	_fdg[704] = code{Code: 102 << 8, BitsWritten: 9}
	_fdg[768] = code{Code: 102<<8 | 1<<7, BitsWritten: 9}
	_fdg[832] = code{Code: 105 << 8, BitsWritten: 9}
	_fdg[896] = code{Code: 105<<8 | 1<<7, BitsWritten: 9}
	_fdg[960] = code{Code: 106 << 8, BitsWritten: 9}
	_fdg[1024] = code{Code: 106<<8 | 1<<7, BitsWritten: 9}
	_fdg[1088] = code{Code: 107 << 8, BitsWritten: 9}
	_fdg[1152] = code{Code: 107<<8 | 1<<7, BitsWritten: 9}
	_fdg[1216] = code{Code: 108 << 8, BitsWritten: 9}
	_fdg[1280] = code{Code: 108<<8 | 1<<7, BitsWritten: 9}
	_fdg[1344] = code{Code: 109 << 8, BitsWritten: 9}
	_fdg[1408] = code{Code: 109<<8 | 1<<7, BitsWritten: 9}
	_fdg[1472] = code{Code: 76 << 8, BitsWritten: 9}
	_fdg[1536] = code{Code: 76<<8 | 1<<7, BitsWritten: 9}
	_fdg[1600] = code{Code: 77 << 8, BitsWritten: 9}
	_fdg[1664] = code{Code: 24 << (2 + 8), BitsWritten: 6}
	_fdg[1728] = code{Code: 77<<8 | 1<<7, BitsWritten: 9}
	_ffd = make(map[int]code)
	_ffd[1792] = code{Code: 1 << 8, BitsWritten: 11}
	_ffd[1856] = code{Code: 1<<8 | 4<<5, BitsWritten: 11}
	_ffd[1920] = code{Code: 1<<8 | 5<<5, BitsWritten: 11}
	_ffd[1984] = code{Code: 1<<8 | 2<<4, BitsWritten: 12}
	_ffd[2048] = code{Code: 1<<8 | 3<<4, BitsWritten: 12}
	_ffd[2112] = code{Code: 1<<8 | 4<<4, BitsWritten: 12}
	_ffd[2176] = code{Code: 1<<8 | 5<<4, BitsWritten: 12}
	_ffd[2240] = code{Code: 1<<8 | 6<<4, BitsWritten: 12}
	_ffd[2304] = code{Code: 1<<8 | 7<<4, BitsWritten: 12}
	_ffd[2368] = code{Code: 1<<8 | 12<<4, BitsWritten: 12}
	_ffd[2432] = code{Code: 1<<8 | 13<<4, BitsWritten: 12}
	_ffd[2496] = code{Code: 1<<8 | 14<<4, BitsWritten: 12}
	_ffd[2560] = code{Code: 1<<8 | 15<<4, BitsWritten: 12}
	_ebb = make(map[int]byte)
	_ebb[0] = 0xFF
	_ebb[1] = 0xFE
	_ebb[2] = 0xFC
	_ebb[3] = 0xF8
	_ebb[4] = 0xF0
	_ebb[5] = 0xE0
	_ebb[6] = 0xC0
	_ebb[7] = 0x80
	_ebb[8] = 0x00
}
func (_gee *Decoder) tryFetchRTC2D() (_fgef error) {
	_gee._dag.Mark()
	var _ecg bool
	for _fcdd := 0; _fcdd < 5; _fcdd++ {
		_ecg, _fgef = _gee.tryFetchEOL1()
		if _fgef != nil {
			if _g.Is(_fgef, _e.EOF) {
				if _fcdd == 0 {
					break
				}
				return _bfd
			}
		}
		if _ecg {
			continue
		}
		if _fcdd > 0 {
			return _bfd
		}
		break
	}
	if _ecg {
		return _e.EOF
	}
	_gee._dag.Reset()
	return _fgef
}
func _eff(_cge, _bac int) code {
	var _dfd code
	switch _bac - _cge {
	case -1:
		_dfd = _bd
	case -2:
		_dfd = _bfb
	case -3:
		_dfd = _ga
	case 0:
		_dfd = _cb
	case 1:
		_dfd = _ce
	case 2:
		_dfd = _ee
	case 3:
		_dfd = _abf
	}
	return _dfd
}
func (_ad *Decoder) fetch() error {
	if _ad._dfe == -1 {
		return nil
	}
	if _ad._ecf < _ad._dfe {
		return nil
	}
	_ad._dfe = 0
	_fff := _ad.decodeRow()
	if _fff != nil {
		if !_g.Is(_fff, _e.EOF) {
			return _fff
		}
		if _ad._dfe != 0 {
			return _fff
		}
		_ad._dfe = -1
	}
	_ad._ecf = 0
	return nil
}
func (_ffbc *treeNode) walk(_efff bool) *treeNode {
	if _efff {
		return _ffbc._efefd
	}
	return _ffbc._ebf
}

var (
	_bb  map[int]code
	_gbc map[int]code
	_fc  map[int]code
	_fdg map[int]code
	_ffd map[int]code
	_ebb map[int]byte
	_ec  = code{Code: 1 << 4, BitsWritten: 12}
	_caf = code{Code: 3 << 3, BitsWritten: 13}
	_gbb = code{Code: 2 << 3, BitsWritten: 13}
	_fde = code{Code: 1 << 12, BitsWritten: 4}
	_age = code{Code: 1 << 13, BitsWritten: 3}
	_cb  = code{Code: 1 << 15, BitsWritten: 1}
	_bd  = code{Code: 3 << 13, BitsWritten: 3}
	_bfb = code{Code: 3 << 10, BitsWritten: 6}
	_ga  = code{Code: 3 << 9, BitsWritten: 7}
	_ce  = code{Code: 2 << 13, BitsWritten: 3}
	_ee  = code{Code: 2 << 10, BitsWritten: 6}
	_abf = code{Code: 2 << 9, BitsWritten: 7}
)

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

func (_gd *Decoder) decodeRowType4() error {
	if !_gd._abg {
		return _gd.decoderRowType41D()
	}
	if _gd._eac {
		_gd._dag.Align()
	}
	_gd._dag.Mark()
	_bfg, _gcdb := _gd.tryFetchEOL()
	if _gcdb != nil {
		return _gcdb
	}
	if !_bfg && _gd._gg {
		_gd._ae++
		if _gd._ae > _gd._bdf {
			return _aac
		}
		_gd._dag.Reset()
	}
	if !_bfg {
		_gd._dag.Reset()
	}
	_dgb, _gcdb := _gd._dag.ReadBool()
	if _gcdb != nil {
		return _gcdb
	}
	if _dgb {
		if _bfg && _gd._bcb {
			if _gcdb = _gd.tryFetchRTC2D(); _gcdb != nil {
				return _gcdb
			}
		}
		_gcdb = _gd.decode1D()
	} else {
		_gcdb = _gd.decode2D()
	}
	if _gcdb != nil {
		return _gcdb
	}
	return nil
}
func _bddg(_gbg []byte, _cfa int, _cba code) ([]byte, int) {
	_cdce := true
	var _eaad []byte
	_eaad, _cfa = _acbf(nil, _cfa, _cba)
	_bcc := 0
	var _cgd int
	for _bcc < len(_gbg) {
		_cgd, _bcc = _fag(_gbg, _cdce, _bcc)
		_eaad, _cfa = _bfae(_eaad, _cfa, _cgd, _cdce)
		_cdce = !_cdce
	}
	return _eaad, _cfa % 8
}
func (_fgg *Decoder) tryFetchEOL() (bool, error) {
	_bbd, _dbcf := _fgg._dag.ReadBits(12)
	if _dbcf != nil {
		return false, _dbcf
	}
	return _bbd == 0x1, nil
}
func (_dad *Decoder) decodeRow() (_baa error) {
	if !_dad._bcb && _dad._gcd > 0 && _dad._gcd == _dad._bfc {
		return _e.EOF
	}
	switch _dad._cafe {
	case _eeg:
		_baa = _dad.decodeRowType2()
	case _ba:
		_baa = _dad.decodeRowType4()
	case _fb:
		_baa = _dad.decodeRowType6()
	}
	if _baa != nil {
		return _baa
	}
	_cfc := 0
	_dge := true
	_dad._bda = 0
	for _ebgb := 0; _ebgb < _dad._bdc; _ebgb++ {
		_fbee := _dad._afc
		if _ebgb != _dad._bdc {
			_fbee = _dad._gac[_ebgb]
		}
		if _fbee > _dad._afc {
			_fbee = _dad._afc
		}
		_fffa := _cfc / 8
		for _cfc%8 != 0 && _fbee-_cfc > 0 {
			var _dc byte
			if !_dge {
				_dc = 1 << uint(7-(_cfc%8))
			}
			_dad._afg[_fffa] |= _dc
			_cfc++
		}
		if _cfc%8 == 0 {
			_fffa = _cfc / 8
			var _bgce byte
			if !_dge {
				_bgce = 0xff
			}
			for _fbee-_cfc > 7 {
				_dad._afg[_fffa] = _bgce
				_cfc += 8
				_fffa++
			}
		}
		for _fbee-_cfc > 0 {
			if _cfc%8 == 0 {
				_dad._afg[_fffa] = 0
			}
			var _bbc byte
			if !_dge {
				_bbc = 1 << uint(7-(_cfc%8))
			}
			_dad._afg[_fffa] |= _bbc
			_cfc++
		}
		_dge = !_dge
	}
	if _cfc != _dad._afc {
		return _g.New("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068")
	}
	_dad._dfe = (_cfc + 7) / 8
	_dad._bfc++
	return nil
}
func (_cec *Decoder) decodeRowType6() error {
	if _cec._eac {
		_cec._dag.Align()
	}
	if _cec._bcb {
		_cec._dag.Mark()
		_ffe, _efd := _cec.tryFetchEOL()
		if _efd != nil {
			return _efd
		}
		if _ffe {
			_ffe, _efd = _cec.tryFetchEOL()
			if _efd != nil {
				return _efd
			}
			if _ffe {
				return _e.EOF
			}
		}
		_cec._dag.Reset()
	}
	return _cec.decode2D()
}
func (_aacf *Decoder) decodeRowType2() error {
	if _aacf._eac {
		_aacf._dag.Align()
	}
	if _gaa := _aacf.decode1D(); _gaa != nil {
		return _gaa
	}
	return nil
}
func _gce(_bddgb, _bdfe []byte, _abfg int) int {
	_ccc := _bfcg(_bdfe, _abfg)
	if _ccc < len(_bdfe) && (_abfg == -1 && _bdfe[_ccc] == _ged || _abfg >= 0 && _abfg < len(_bddgb) && _bddgb[_abfg] == _bdfe[_ccc] || _abfg >= len(_bddgb) && _bddgb[_abfg-1] != _bdfe[_ccc]) {
		_ccc = _bfcg(_bdfe, _ccc)
	}
	return _ccc
}
func _bccc(_gced []byte, _fgae, _eeca, _cfdd int) ([]byte, int) {
	_aace := _eff(_eeca, _cfdd)
	_gced, _fgae = _acbf(_gced, _fgae, _aace)
	return _gced, _fgae
}
func _dgbac(_dab [][]byte) [][]byte {
	_feb := make([]byte, len(_dab[0]))
	for _efg := range _feb {
		_feb[_efg] = _ged
	}
	_dab = append(_dab, []byte{})
	for _bgcd := len(_dab) - 1; _bgcd > 0; _bgcd-- {
		_dab[_bgcd] = _dab[_bgcd-1]
	}
	_dab[0] = _feb
	return _dab
}
func (_fge *Decoder) tryFetchEOL1() (bool, error) {
	_acg, _bag := _fge._dag.ReadBits(13)
	if _bag != nil {
		return false, _bag
	}
	return _acg == 0x3, nil
}
func _acbf(_eda []byte, _dbec int, _fgd code) ([]byte, int) {
	_aebf := 0
	for _aebf < _fgd.BitsWritten {
		_ddef := _dbec / 8
		_fgb := _dbec % 8
		if _ddef >= len(_eda) {
			_eda = append(_eda, 0)
		}
		_eba := 8 - _fgb
		_dae := _fgd.BitsWritten - _aebf
		if _eba > _dae {
			_eba = _dae
		}
		if _aebf < 8 {
			_eda[_ddef] = _eda[_ddef] | byte(_fgd.Code>>uint(8+_fgb-_aebf))&_ebb[8-_eba-_fgb]
		} else {
			_eda[_ddef] = _eda[_ddef] | (byte(_fgd.Code<<uint(_aebf-8))&_ebb[8-_eba])>>uint(_fgb)
		}
		_dbec += _eba
		_aebf += _eba
	}
	return _eda, _dbec
}
func (_dfcg *Decoder) looseFetchEOL() (bool, error) {
	_ega, _fcd := _dfcg._dag.ReadBits(12)
	if _fcd != nil {
		return false, _fcd
	}
	switch _ega {
	case 0x1:
		return true, nil
	case 0x0:
		for {
			_ggce, _ced := _dfcg._dag.ReadBool()
			if _ced != nil {
				return false, _ced
			}
			if _ggce {
				return true, nil
			}
		}
	default:
		return false, nil
	}
}
func (_bgg *Encoder) encodeG32D(_cab [][]byte) []byte {
	var _fdb []byte
	var _gfd int
	for _ecd := 0; _ecd < len(_cab); _ecd += _bgg.K {
		if _bgg.Rows > 0 && !_bgg.EndOfBlock && _ecd == _bgg.Rows {
			break
		}
		_fab, _bdcc := _bddg(_cab[_ecd], _gfd, _caf)
		_fdb = _bgg.appendEncodedRow(_fdb, _fab, _gfd)
		if _bgg.EncodedByteAlign {
			_bdcc = 0
		}
		_gfd = _bdcc
		for _cfce := _ecd + 1; _cfce < (_ecd+_bgg.K) && _cfce < len(_cab); _cfce++ {
			if _bgg.Rows > 0 && !_bgg.EndOfBlock && _cfce == _bgg.Rows {
				break
			}
			_cgg, _bcf := _acbf(nil, _gfd, _gbb)
			var _bfgf, _ccee, _afe int
			_fabd := -1
			for _fabd < len(_cab[_cfce]) {
				_bfgf = _bfcg(_cab[_cfce], _fabd)
				_ccee = _gce(_cab[_cfce], _cab[_cfce-1], _fabd)
				_afe = _bfcg(_cab[_cfce-1], _ccee)
				if _afe < _bfgf {
					_cgg, _bcf = _cgc(_cgg, _bcf)
					_fabd = _afe
				} else {
					if _b.Abs(float64(_ccee-_bfgf)) > 3 {
						_cgg, _bcf, _fabd = _agg(_cab[_cfce], _cgg, _bcf, _fabd, _bfgf)
					} else {
						_cgg, _bcf = _bccc(_cgg, _bcf, _bfgf, _ccee)
						_fabd = _bfgf
					}
				}
			}
			_fdb = _bgg.appendEncodedRow(_fdb, _cgg, _gfd)
			if _bgg.EncodedByteAlign {
				_bcf = 0
			}
			_gfd = _bcf % 8
		}
	}
	if _bgg.EndOfBlock {
		_edf, _ := _edfe(_gfd)
		_fdb = _bgg.appendEncodedRow(_fdb, _edf, _gfd)
	}
	return _fdb
}

var (
	_bfd = _g.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043")
	_aac = _g.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
)

func (_ddg *treeNode) set(_dgcc bool, _cdef *treeNode) {
	if !_dgcc {
		_ddg._ebf = _cdef
	} else {
		_ddg._efefd = _cdef
	}
}
func _fdf(_ccf int, _faa bool) (code, int, bool) {
	if _ccf < 64 {
		if _faa {
			return _gbc[_ccf], 0, true
		}
		return _bb[_ccf], 0, true
	}
	_bae := _ccf / 64
	if _bae > 40 {
		return _ffd[2560], _ccf - 2560, false
	}
	if _bae > 27 {
		return _ffd[_bae*64], _ccf - _bae*64, false
	}
	if _faa {
		return _fdg[_bae*64], _ccf - _bae*64, false
	}
	return _fc[_bae*64], _ccf - _bae*64, false
}

type treeNode struct {
	_ebf   *treeNode
	_efefd *treeNode
	_fffag int
	_gfda  bool
	_dfb   bool
}

func _edfe(_dca int) ([]byte, int) {
	var _gfe []byte
	for _gfag := 0; _gfag < 6; _gfag++ {
		_gfe, _dca = _acbf(_gfe, _dca, _caf)
	}
	return _gfe, _dca % 8
}
func (_cbd *Decoder) decode1D() error {
	var (
		_dcc  int
		_cacc error
	)
	_aad := true
	_cbd._bdc = 0
	for {
		var _cg int
		if _aad {
			_cg, _cacc = _cbd.decodeRun(_da)
		} else {
			_cg, _cacc = _cbd.decodeRun(_gf)
		}
		if _cacc != nil {
			return _cacc
		}
		_dcc += _cg
		_cbd._gac[_cbd._bdc] = _dcc
		_cbd._bdc++
		_aad = !_aad
		if _dcc >= _cbd._afc {
			break
		}
	}
	return nil
}

var (
	_ged byte = 1
	_ddb byte = 0
)

func (_afcd *Encoder) Encode(pixels [][]byte) []byte {
	if _afcd.BlackIs1 {
		_ged = 0
		_ddb = 1
	} else {
		_ged = 1
		_ddb = 0
	}
	if _afcd.K == 0 {
		return _afcd.encodeG31D(pixels)
	}
	if _afcd.K > 0 {
		return _afcd.encodeG32D(pixels)
	}
	if _afcd.K < 0 {
		return _afcd.encodeG4(pixels)
	}
	return nil
}
func _adf(_cceg, _ddeb []byte, _ecfb int, _cad bool) int {
	_cfcg := _bfcg(_ddeb, _ecfb)
	if _cfcg < len(_ddeb) && (_ecfb == -1 && _ddeb[_cfcg] == _ged || _ecfb >= 0 && _ecfb < len(_cceg) && _cceg[_ecfb] == _ddeb[_cfcg] || _ecfb >= len(_cceg) && _cad && _ddeb[_cfcg] == _ged || _ecfb >= len(_cceg) && !_cad && _ddeb[_cfcg] == _ddb) {
		_cfcg = _bfcg(_ddeb, _cfcg)
	}
	return _cfcg
}
func _bfae(_eabe []byte, _geb int, _aeg int, _bab bool) ([]byte, int) {
	var (
		_dccfc code
		_ggb   bool
	)
	for !_ggb {
		_dccfc, _aeg, _ggb = _fdf(_aeg, _bab)
		_eabe, _geb = _acbf(_eabe, _geb, _dccfc)
	}
	return _eabe, _geb
}

var _bge = [...][]uint16{{0x7, 0x8, 0xb, 0xc, 0xe, 0xf}, {0x12, 0x13, 0x14, 0x1b, 0x7, 0x8}, {0x17, 0x18, 0x2a, 0x2b, 0x3, 0x34, 0x35, 0x7, 0x8}, {0x13, 0x17, 0x18, 0x24, 0x27, 0x28, 0x2b, 0x3, 0x37, 0x4, 0x8, 0xc}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1a, 0x1b, 0x2, 0x24, 0x25, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x3, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x4, 0x4a, 0x4b, 0x5, 0x52, 0x53, 0x54, 0x55, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x67, 0x68, 0xa, 0xb}, {0x98, 0x99, 0x9a, 0x9b, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb}, {}, {0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f}}

func _bfcg(_efdb []byte, _cfd int) int {
	if _cfd >= len(_efdb) {
		return _cfd
	}
	if _cfd < -1 {
		_cfd = -1
	}
	var _fdbg byte
	if _cfd > -1 {
		_fdbg = _efdb[_cfd]
	} else {
		_fdbg = _ged
	}
	_fcc := _cfd + 1
	for _fcc < len(_efdb) {
		if _efdb[_fcc] != _fdbg {
			break
		}
		_fcc++
	}
	return _fcc
}
func _dcb(_dbd int) ([]byte, int) {
	var _eee []byte
	for _ffee := 0; _ffee < 6; _ffee++ {
		_eee, _dbd = _acbf(_eee, _dbd, _ec)
	}
	return _eee, _dbd % 8
}
func _eaa(_fae int) ([]byte, int) {
	var _aff []byte
	for _ecge := 0; _ecge < 2; _ecge++ {
		_aff, _fae = _acbf(_aff, _fae, _ec)
	}
	return _aff, _fae % 8
}
func (_ada *tree) fill(_fef, _fded, _bff int) error {
	_gae := _ada._fgea
	for _bef := 0; _bef < _fef; _bef++ {
		_gbd := _fef - 1 - _bef
		_fgc := ((_fded >> uint(_gbd)) & 1) != 0
		_cde := _gae.walk(_fgc)
		if _cde != nil {
			if _cde._dfb {
				return _g.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_gae = _cde
			continue
		}
		_cde = &treeNode{}
		if _bef == _fef-1 {
			_cde._fffag = _bff
			_cde._dfb = true
		}
		if _fded == 0 {
			_cde._gfda = true
		}
		_gae.set(_fgc, _cde)
		_gae = _cde
	}
	return nil
}
func (_cef *Encoder) encodeG31D(_fgf [][]byte) []byte {
	var _ecc []byte
	_gbf := 0
	for _bgea := range _fgf {
		if _cef.Rows > 0 && !_cef.EndOfBlock && _bgea == _cef.Rows {
			break
		}
		_acgc, _eag := _bddg(_fgf[_bgea], _gbf, _ec)
		_ecc = _cef.appendEncodedRow(_ecc, _acgc, _gbf)
		if _cef.EncodedByteAlign {
			_eag = 0
		}
		_gbf = _eag
	}
	if _cef.EndOfBlock {
		_ggf, _ := _dcb(_gbf)
		_ecc = _cef.appendEncodedRow(_ecc, _ggf, _gbf)
	}
	return _ecc
}
func (_efc *Decoder) getNextChangingElement(_ggd int, _ffdg bool) int {
	_gcc := 0
	if !_ffdg {
		_gcc = 1
	}
	_caff := int(uint32(_efc._bda)&0xFFFFFFFE) + _gcc
	if _caff > 2 {
		_caff -= 2
	}
	if _ggd == 0 {
		return _caff
	}
	for _eeaa := _caff; _eeaa < _efc._fbc; _eeaa += 2 {
		if _ggd < _efc._dfea[_eeaa] {
			_efc._bda = _eeaa
			return _eeaa
		}
	}
	return -1
}

type tiffType int

func (_bca *Encoder) encodeG4(_agf [][]byte) []byte {
	_gge := make([][]byte, len(_agf))
	copy(_gge, _agf)
	_gge = _dgbac(_gge)
	var _eeb []byte
	var _afb int
	for _acf := 1; _acf < len(_gge); _acf++ {
		if _bca.Rows > 0 && !_bca.EndOfBlock && _acf == (_bca.Rows+1) {
			break
		}
		var _ecb []byte
		var _bdfg, _acd, _dff int
		_ccg := _afb
		_cdb := -1
		for _cdb < len(_gge[_acf]) {
			_bdfg = _bfcg(_gge[_acf], _cdb)
			_acd = _gce(_gge[_acf], _gge[_acf-1], _cdb)
			_dff = _bfcg(_gge[_acf-1], _acd)
			if _dff < _bdfg {
				_ecb, _ccg = _acbf(_ecb, _ccg, _fde)
				_cdb = _dff
			} else {
				if _b.Abs(float64(_acd-_bdfg)) > 3 {
					_ecb, _ccg, _cdb = _agg(_gge[_acf], _ecb, _ccg, _cdb, _bdfg)
				} else {
					_ecb, _ccg = _bccc(_ecb, _ccg, _bdfg, _acd)
					_cdb = _bdfg
				}
			}
		}
		_eeb = _bca.appendEncodedRow(_eeb, _ecb, _afb)
		if _bca.EncodedByteAlign {
			_ccg = 0
		}
		_afb = _ccg % 8
	}
	if _bca.EndOfBlock {
		_efef, _ := _eaa(_afb)
		_eeb = _bca.appendEncodedRow(_eeb, _efef, _afb)
	}
	return _eeb
}

type tree struct{ _fgea *treeNode }

func (_acb *Decoder) Read(in []byte) (int, error) {
	if _acb._fdgd != nil {
		return 0, _acb._fdgd
	}
	_cd := len(in)
	var (
		_dfa  int
		_afgb int
	)
	for _cd != 0 {
		if _acb._ecf >= _acb._dfe {
			if _cdc := _acb.fetch(); _cdc != nil {
				_acb._fdgd = _cdc
				return 0, _cdc
			}
		}
		if _acb._dfe == -1 {
			return _dfa, _e.EOF
		}
		switch {
		case _cd <= _acb._dfe-_acb._ecf:
			_dfc := _acb._afg[_acb._ecf : _acb._ecf+_cd]
			for _, _ddeg := range _dfc {
				if !_acb._faf {
					_ddeg = ^_ddeg
				}
				in[_afgb] = _ddeg
				_afgb++
			}
			_dfa += len(_dfc)
			_acb._ecf += len(_dfc)
			return _dfa, nil
		default:
			_dg := _acb._afg[_acb._ecf:]
			for _, _bfab := range _dg {
				if !_acb._faf {
					_bfab = ^_bfab
				}
				in[_afgb] = _bfab
				_afgb++
			}
			_dfa += len(_dg)
			_acb._ecf += len(_dg)
			_cd -= len(_dg)
		}
	}
	return _dfa, nil
}

var (
	_bg *treeNode
	_ff *treeNode
	_gf *tree
	_da *tree
	_a  *tree
	_eg *tree
	_ab = -2000
	_ag = -1000
	_de = -3000
	_ef = -4000
)

type Decoder struct {
	_afc  int
	_gcd  int
	_bfc  int
	_afg  []byte
	_bdf  int
	_abg  bool
	_fed  bool
	_bfa  bool
	_faf  bool
	_gg   bool
	_bcb  bool
	_eac  bool
	_dfe  int
	_ecf  int
	_dfea []int
	_gac  []int
	_fbc  int
	_bdc  int
	_ae   int
	_bda  int
	_dag  *_f.Reader
	_cafe tiffType
	_fdgd error
}

func (_gga *Decoder) decode2D() error {
	_gga._fbc = _gga._bdc
	_gga._gac, _gga._dfea = _gga._dfea, _gga._gac
	_dbc := true
	var (
		_dgc  bool
		_bba  int
		_gfac error
	)
	_gga._bdc = 0
_ggc:
	for _bba < _gga._afc {
		_dccf := _eg._fgea
		for {
			_dgc, _gfac = _gga._dag.ReadBool()
			if _gfac != nil {
				return _gfac
			}
			_dccf = _dccf.walk(_dgc)
			if _dccf == nil {
				continue _ggc
			}
			if !_dccf._dfb {
				continue
			}
			switch _dccf._fffag {
			case _ef:
				var _bdd int
				if _dbc {
					_bdd, _gfac = _gga.decodeRun(_da)
				} else {
					_bdd, _gfac = _gga.decodeRun(_gf)
				}
				if _gfac != nil {
					return _gfac
				}
				_bba += _bdd
				_gga._gac[_gga._bdc] = _bba
				_gga._bdc++
				if _dbc {
					_bdd, _gfac = _gga.decodeRun(_gf)
				} else {
					_bdd, _gfac = _gga.decodeRun(_da)
				}
				if _gfac != nil {
					return _gfac
				}
				_bba += _bdd
				_gga._gac[_gga._bdc] = _bba
				_gga._bdc++
			case _de:
				_dbe := _gga.getNextChangingElement(_bba, _dbc) + 1
				if _dbe >= _gga._fbc {
					_bba = _gga._afc
				} else {
					_bba = _gga._dfea[_dbe]
				}
			default:
				_eea := _gga.getNextChangingElement(_bba, _dbc)
				if _eea >= _gga._fbc || _eea == -1 {
					_bba = _gga._afc + _dccf._fffag
				} else {
					_bba = _gga._dfea[_eea] + _dccf._fffag
				}
				_gga._gac[_gga._bdc] = _bba
				_gga._bdc++
				_dbc = !_dbc
			}
			continue _ggc
		}
	}
	return nil
}

type code struct {
	Code        uint16
	BitsWritten int
}

func _agg(_bfaa, _eeed []byte, _agea, _dgce, _dbg int) ([]byte, int, int) {
	_cbe := _bfcg(_bfaa, _dbg)
	_cdbe := _dgce >= 0 && _bfaa[_dgce] == _ged || _dgce == -1
	_eeed, _agea = _acbf(_eeed, _agea, _age)
	var _aef int
	if _dgce > -1 {
		_aef = _dbg - _dgce
	} else {
		_aef = _dbg - _dgce - 1
	}
	_eeed, _agea = _bfae(_eeed, _agea, _aef, _cdbe)
	_cdbe = !_cdbe
	_gfdf := _cbe - _dbg
	_eeed, _agea = _bfae(_eeed, _agea, _gfdf, _cdbe)
	_dgce = _cbe
	return _eeed, _agea, _dgce
}
func (_fdc tiffType) String() string {
	switch _fdc {
	case _eeg:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065"
	case _ba:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034"
	case _fb:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036"
	default:
		return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d"
	}
}

var _ge = [...][]uint16{{3, 2}, {1, 4}, {6, 5}, {7}, {9, 8}, {10, 11, 12}, {13, 14}, {15}, {16, 17, 0, 18, 64}, {24, 25, 23, 22, 19, 20, 21, 1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560, 52, 55, 56, 59, 60, 320, 384, 448, 53, 54, 50, 51, 44, 45, 46, 47, 57, 58, 61, 256, 48, 49, 62, 63, 30, 31, 32, 33, 40, 41, 128, 192, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 42, 43}, {640, 704, 768, 832, 1280, 1344, 1408, 1472, 1536, 1600, 1664, 1728, 512, 576, 896, 960, 1024, 1088, 1152, 1216}}

func _fag(_gdeg []byte, _fbb bool, _ded int) (int, int) {
	_ffeg := 0
	for _ded < len(_gdeg) {
		if _fbb {
			if _gdeg[_ded] != _ged {
				break
			}
		} else {
			if _gdeg[_ded] != _ddb {
				break
			}
		}
		_ffeg++
		_ded++
	}
	return _ffeg, _ded
}
func (_ed *Decoder) decoderRowType41D() error {
	if _ed._eac {
		_ed._dag.Align()
	}
	_ed._dag.Mark()
	var (
		_fcg bool
		_feg error
	)
	if _ed._gg {
		_fcg, _feg = _ed.tryFetchEOL()
		if _feg != nil {
			return _feg
		}
		if !_fcg {
			return _aac
		}
	} else {
		_fcg, _feg = _ed.looseFetchEOL()
		if _feg != nil {
			return _feg
		}
	}
	if !_fcg {
		_ed._dag.Reset()
	}
	if _fcg && _ed._bcb {
		_ed._dag.Mark()
		for _agd := 0; _agd < 5; _agd++ {
			_fcg, _feg = _ed.tryFetchEOL()
			if _feg != nil {
				if _g.Is(_feg, _e.EOF) {
					if _agd == 0 {
						break
					}
					return _bfd
				}
			}
			if _fcg {
				continue
			}
			if _agd > 0 {
				return _bfd
			}
			break
		}
		if _fcg {
			return _e.EOF
		}
		_ed._dag.Reset()
	}
	if _feg = _ed.decode1D(); _feg != nil {
		return _feg
	}
	return nil
}
func (_bdg *Encoder) appendEncodedRow(_acda, _cfcc []byte, _bce int) []byte {
	if len(_acda) > 0 && _bce != 0 && !_bdg.EncodedByteAlign {
		_acda[len(_acda)-1] = _acda[len(_acda)-1] | _cfcc[0]
		_acda = append(_acda, _cfcc[1:]...)
	} else {
		_acda = append(_acda, _cfcc...)
	}
	return _acda
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

func (_dgf *Decoder) decodeRun(_aca *tree) (int, error) {
	var _dgba int
	_fbcc := _aca._fgea
	for {
		_dged, _ffb := _dgf._dag.ReadBool()
		if _ffb != nil {
			return 0, _ffb
		}
		_fbcc = _fbcc.walk(_dged)
		if _fbcc == nil {
			return 0, _g.New("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _fbcc._dfb {
			_dgba += _fbcc._fffag
			switch {
			case _fbcc._fffag >= 64:
				_fbcc = _aca._fgea
			case _fbcc._fffag >= 0:
				return _dgba, nil
			default:
				return _dgf._afc, nil
			}
		}
	}
}
