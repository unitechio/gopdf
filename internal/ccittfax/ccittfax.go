package ccittfax

import (
	_g "errors"
	_a "io"
	_gd "math"

	_c "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
)

var (
	_ag  *treeNode
	_ff  *treeNode
	_cf  *tree
	_e   *tree
	_gg  *tree
	_fe  *tree
	_ca  = -2000
	_cfb = -1000
	_ee  = -3000
	_caf = -4000
)

func init() {
	_ag = &treeNode{_bce: true, _bfe: _ca}
	_ff = &treeNode{_bfe: _cfb, _age: _ag}
	_ff._adff = _ff
	_gg = &tree{_fbff: &treeNode{}}
	if _eb := _gg.fillWithNode(12, 0, _ff); _eb != nil {
		panic(_eb.Error())
	}
	if _b := _gg.fillWithNode(12, 1, _ag); _b != nil {
		panic(_b.Error())
	}
	_cf = &tree{_fbff: &treeNode{}}
	for _ggd := 0; _ggd < len(_bb); _ggd++ {
		for _cb := 0; _cb < len(_bb[_ggd]); _cb++ {
			if _fd := _cf.fill(_ggd+2, int(_bb[_ggd][_cb]), int(_aa[_ggd][_cb])); _fd != nil {
				panic(_fd.Error())
			}
		}
	}
	if _cff := _cf.fillWithNode(12, 0, _ff); _cff != nil {
		panic(_cff.Error())
	}
	if _cab := _cf.fillWithNode(12, 1, _ag); _cab != nil {
		panic(_cab.Error())
	}
	_e = &tree{_fbff: &treeNode{}}
	for _d := 0; _d < len(_dg); _d++ {
		for _fc := 0; _fc < len(_dg[_d]); _fc++ {
			if _fa := _e.fill(_d+4, int(_dg[_d][_fc]), int(_gdg[_d][_fc])); _fa != nil {
				panic(_fa.Error())
			}
		}
	}
	if _cd := _e.fillWithNode(12, 0, _ff); _cd != nil {
		panic(_cd.Error())
	}
	if _fdc := _e.fillWithNode(12, 1, _ag); _fdc != nil {
		panic(_fdc.Error())
	}
	_fe = &tree{_fbff: &treeNode{}}
	if _ga := _fe.fill(4, 1, _ee); _ga != nil {
		panic(_ga.Error())
	}
	if _gf := _fe.fill(3, 1, _caf); _gf != nil {
		panic(_gf.Error())
	}
	if _fad := _fe.fill(1, 1, 0); _fad != nil {
		panic(_fad.Error())
	}
	if _fab := _fe.fill(3, 3, 1); _fab != nil {
		panic(_fab.Error())
	}
	if _dc := _fe.fill(6, 3, 2); _dc != nil {
		panic(_dc.Error())
	}
	if _cae := _fe.fill(7, 3, 3); _cae != nil {
		panic(_cae.Error())
	}
	if _fdf := _fe.fill(3, 2, -1); _fdf != nil {
		panic(_fdf.Error())
	}
	if _ec := _fe.fill(6, 2, -2); _ec != nil {
		panic(_ec.Error())
	}
	if _ad := _fe.fill(7, 2, -3); _ad != nil {
		panic(_ad.Error())
	}
}
func _ebd(_ffcb, _cba []byte, _gac int) int {
	_gbgf := _ecfd(_cba, _gac)
	if _gbgf < len(_cba) && (_gac == -1 && _cba[_gbgf] == _ceb || _gac >= 0 && _gac < len(_ffcb) && _ffcb[_gac] == _cba[_gbgf] || _gac >= len(_ffcb) && _ffcb[_gac-1] != _cba[_gbgf]) {
		_gbgf = _ecfd(_cba, _gbgf)
	}
	return _gbgf
}

var (
	_gff = _g.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043")
	_dbd = _g.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
)

func (_df tiffType) String() string {
	switch _df {
	case _ab:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065"
	case _db:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034"
	case _gdf:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036"
	default:
		return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d"
	}
}
func (_ecb *treeNode) walk(_ebga bool) *treeNode {
	if _ebga {
		return _ecb._age
	}
	return _ecb._adff
}
func (_dae *Encoder) encodeG4(_fbg [][]byte) []byte {
	_cad := make([][]byte, len(_fbg))
	copy(_cad, _fbg)
	_cad = _fbd(_cad)
	var _abe []byte
	var _fag int
	for _fga := 1; _fga < len(_cad); _fga++ {
		if _dae.Rows > 0 && !_dae.EndOfBlock && _fga == (_dae.Rows+1) {
			break
		}
		var _aee []byte
		var _ccf, _def, _cbgf int
		_aaf := _fag
		_fcb := -1
		for _fcb < len(_cad[_fga]) {
			_ccf = _ecfd(_cad[_fga], _fcb)
			_def = _ebd(_cad[_fga], _cad[_fga-1], _fcb)
			_cbgf = _ecfd(_cad[_fga-1], _def)
			if _cbgf < _ccf {
				_aee, _aaf = _bgb(_aee, _aaf, _dge)
				_fcb = _cbgf
			} else {
				if _gd.Abs(float64(_def-_ccf)) > 3 {
					_aee, _aaf, _fcb = _eada(_cad[_fga], _aee, _aaf, _fcb, _ccf)
				} else {
					_aee, _aaf = _acab(_aee, _aaf, _ccf, _def)
					_fcb = _ccf
				}
			}
		}
		_abe = _dae.appendEncodedRow(_abe, _aee, _fag)
		if _dae.EncodedByteAlign {
			_aaf = 0
		}
		_fag = _aaf % 8
	}
	if _dae.EndOfBlock {
		_ddg, _ := _agdf(_fag)
		_abe = _dae.appendEncodedRow(_abe, _ddg, _fag)
	}
	return _abe
}
func (_gag *Decoder) Read(in []byte) (int, error) {
	if _gag._eg != nil {
		return 0, _gag._eg
	}
	_ba := len(in)
	var (
		_fbc int
		_gcg int
	)
	for _ba != 0 {
		if _gag._bbc >= _gag._bf {
			if _af := _gag.fetch(); _af != nil {
				_gag._eg = _af
				return 0, _af
			}
		}
		if _gag._bf == -1 {
			return _fbc, _a.EOF
		}
		switch {
		case _ba <= _gag._bf-_gag._bbc:
			_cdd := _gag._gfd[_gag._bbc : _gag._bbc+_ba]
			for _, _eag := range _cdd {
				if !_gag._acc {
					_eag = ^_eag
				}
				in[_gcg] = _eag
				_gcg++
			}
			_fbc += len(_cdd)
			_gag._bbc += len(_cdd)
			return _fbc, nil
		default:
			_gba := _gag._gfd[_gag._bbc:]
			for _, _fae := range _gba {
				if !_gag._acc {
					_fae = ^_fae
				}
				in[_gcg] = _fae
				_gcg++
			}
			_fbc += len(_gba)
			_gag._bbc += len(_gba)
			_ba -= len(_gba)
		}
	}
	return _fbc, nil
}
func init() {
	_ggb = make(map[int]code)
	_ggb[0] = code{Code: 13<<8 | 3<<6, BitsWritten: 10}
	_ggb[1] = code{Code: 2 << (5 + 8), BitsWritten: 3}
	_ggb[2] = code{Code: 3 << (6 + 8), BitsWritten: 2}
	_ggb[3] = code{Code: 2 << (6 + 8), BitsWritten: 2}
	_ggb[4] = code{Code: 3 << (5 + 8), BitsWritten: 3}
	_ggb[5] = code{Code: 3 << (4 + 8), BitsWritten: 4}
	_ggb[6] = code{Code: 2 << (4 + 8), BitsWritten: 4}
	_ggb[7] = code{Code: 3 << (3 + 8), BitsWritten: 5}
	_ggb[8] = code{Code: 5 << (2 + 8), BitsWritten: 6}
	_ggb[9] = code{Code: 4 << (2 + 8), BitsWritten: 6}
	_ggb[10] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_ggb[11] = code{Code: 5 << (1 + 8), BitsWritten: 7}
	_ggb[12] = code{Code: 7 << (1 + 8), BitsWritten: 7}
	_ggb[13] = code{Code: 4 << 8, BitsWritten: 8}
	_ggb[14] = code{Code: 7 << 8, BitsWritten: 8}
	_ggb[15] = code{Code: 12 << 8, BitsWritten: 9}
	_ggb[16] = code{Code: 5<<8 | 3<<6, BitsWritten: 10}
	_ggb[17] = code{Code: 6 << 8, BitsWritten: 10}
	_ggb[18] = code{Code: 2 << 8, BitsWritten: 10}
	_ggb[19] = code{Code: 12<<8 | 7<<5, BitsWritten: 11}
	_ggb[20] = code{Code: 13 << 8, BitsWritten: 11}
	_ggb[21] = code{Code: 13<<8 | 4<<5, BitsWritten: 11}
	_ggb[22] = code{Code: 6<<8 | 7<<5, BitsWritten: 11}
	_ggb[23] = code{Code: 5 << 8, BitsWritten: 11}
	_ggb[24] = code{Code: 2<<8 | 7<<5, BitsWritten: 11}
	_ggb[25] = code{Code: 3 << 8, BitsWritten: 11}
	_ggb[26] = code{Code: 12<<8 | 10<<4, BitsWritten: 12}
	_ggb[27] = code{Code: 12<<8 | 11<<4, BitsWritten: 12}
	_ggb[28] = code{Code: 12<<8 | 12<<4, BitsWritten: 12}
	_ggb[29] = code{Code: 12<<8 | 13<<4, BitsWritten: 12}
	_ggb[30] = code{Code: 6<<8 | 8<<4, BitsWritten: 12}
	_ggb[31] = code{Code: 6<<8 | 9<<4, BitsWritten: 12}
	_ggb[32] = code{Code: 6<<8 | 10<<4, BitsWritten: 12}
	_ggb[33] = code{Code: 6<<8 | 11<<4, BitsWritten: 12}
	_ggb[34] = code{Code: 13<<8 | 2<<4, BitsWritten: 12}
	_ggb[35] = code{Code: 13<<8 | 3<<4, BitsWritten: 12}
	_ggb[36] = code{Code: 13<<8 | 4<<4, BitsWritten: 12}
	_ggb[37] = code{Code: 13<<8 | 5<<4, BitsWritten: 12}
	_ggb[38] = code{Code: 13<<8 | 6<<4, BitsWritten: 12}
	_ggb[39] = code{Code: 13<<8 | 7<<4, BitsWritten: 12}
	_ggb[40] = code{Code: 6<<8 | 12<<4, BitsWritten: 12}
	_ggb[41] = code{Code: 6<<8 | 13<<4, BitsWritten: 12}
	_ggb[42] = code{Code: 13<<8 | 10<<4, BitsWritten: 12}
	_ggb[43] = code{Code: 13<<8 | 11<<4, BitsWritten: 12}
	_ggb[44] = code{Code: 5<<8 | 4<<4, BitsWritten: 12}
	_ggb[45] = code{Code: 5<<8 | 5<<4, BitsWritten: 12}
	_ggb[46] = code{Code: 5<<8 | 6<<4, BitsWritten: 12}
	_ggb[47] = code{Code: 5<<8 | 7<<4, BitsWritten: 12}
	_ggb[48] = code{Code: 6<<8 | 4<<4, BitsWritten: 12}
	_ggb[49] = code{Code: 6<<8 | 5<<4, BitsWritten: 12}
	_ggb[50] = code{Code: 5<<8 | 2<<4, BitsWritten: 12}
	_ggb[51] = code{Code: 5<<8 | 3<<4, BitsWritten: 12}
	_ggb[52] = code{Code: 2<<8 | 4<<4, BitsWritten: 12}
	_ggb[53] = code{Code: 3<<8 | 7<<4, BitsWritten: 12}
	_ggb[54] = code{Code: 3<<8 | 8<<4, BitsWritten: 12}
	_ggb[55] = code{Code: 2<<8 | 7<<4, BitsWritten: 12}
	_ggb[56] = code{Code: 2<<8 | 8<<4, BitsWritten: 12}
	_ggb[57] = code{Code: 5<<8 | 8<<4, BitsWritten: 12}
	_ggb[58] = code{Code: 5<<8 | 9<<4, BitsWritten: 12}
	_ggb[59] = code{Code: 2<<8 | 11<<4, BitsWritten: 12}
	_ggb[60] = code{Code: 2<<8 | 12<<4, BitsWritten: 12}
	_ggb[61] = code{Code: 5<<8 | 10<<4, BitsWritten: 12}
	_ggb[62] = code{Code: 6<<8 | 6<<4, BitsWritten: 12}
	_ggb[63] = code{Code: 6<<8 | 7<<4, BitsWritten: 12}
	_ac = make(map[int]code)
	_ac[0] = code{Code: 53 << 8, BitsWritten: 8}
	_ac[1] = code{Code: 7 << (2 + 8), BitsWritten: 6}
	_ac[2] = code{Code: 7 << (4 + 8), BitsWritten: 4}
	_ac[3] = code{Code: 8 << (4 + 8), BitsWritten: 4}
	_ac[4] = code{Code: 11 << (4 + 8), BitsWritten: 4}
	_ac[5] = code{Code: 12 << (4 + 8), BitsWritten: 4}
	_ac[6] = code{Code: 14 << (4 + 8), BitsWritten: 4}
	_ac[7] = code{Code: 15 << (4 + 8), BitsWritten: 4}
	_ac[8] = code{Code: 19 << (3 + 8), BitsWritten: 5}
	_ac[9] = code{Code: 20 << (3 + 8), BitsWritten: 5}
	_ac[10] = code{Code: 7 << (3 + 8), BitsWritten: 5}
	_ac[11] = code{Code: 8 << (3 + 8), BitsWritten: 5}
	_ac[12] = code{Code: 8 << (2 + 8), BitsWritten: 6}
	_ac[13] = code{Code: 3 << (2 + 8), BitsWritten: 6}
	_ac[14] = code{Code: 52 << (2 + 8), BitsWritten: 6}
	_ac[15] = code{Code: 53 << (2 + 8), BitsWritten: 6}
	_ac[16] = code{Code: 42 << (2 + 8), BitsWritten: 6}
	_ac[17] = code{Code: 43 << (2 + 8), BitsWritten: 6}
	_ac[18] = code{Code: 39 << (1 + 8), BitsWritten: 7}
	_ac[19] = code{Code: 12 << (1 + 8), BitsWritten: 7}
	_ac[20] = code{Code: 8 << (1 + 8), BitsWritten: 7}
	_ac[21] = code{Code: 23 << (1 + 8), BitsWritten: 7}
	_ac[22] = code{Code: 3 << (1 + 8), BitsWritten: 7}
	_ac[23] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_ac[24] = code{Code: 40 << (1 + 8), BitsWritten: 7}
	_ac[25] = code{Code: 43 << (1 + 8), BitsWritten: 7}
	_ac[26] = code{Code: 19 << (1 + 8), BitsWritten: 7}
	_ac[27] = code{Code: 36 << (1 + 8), BitsWritten: 7}
	_ac[28] = code{Code: 24 << (1 + 8), BitsWritten: 7}
	_ac[29] = code{Code: 2 << 8, BitsWritten: 8}
	_ac[30] = code{Code: 3 << 8, BitsWritten: 8}
	_ac[31] = code{Code: 26 << 8, BitsWritten: 8}
	_ac[32] = code{Code: 27 << 8, BitsWritten: 8}
	_ac[33] = code{Code: 18 << 8, BitsWritten: 8}
	_ac[34] = code{Code: 19 << 8, BitsWritten: 8}
	_ac[35] = code{Code: 20 << 8, BitsWritten: 8}
	_ac[36] = code{Code: 21 << 8, BitsWritten: 8}
	_ac[37] = code{Code: 22 << 8, BitsWritten: 8}
	_ac[38] = code{Code: 23 << 8, BitsWritten: 8}
	_ac[39] = code{Code: 40 << 8, BitsWritten: 8}
	_ac[40] = code{Code: 41 << 8, BitsWritten: 8}
	_ac[41] = code{Code: 42 << 8, BitsWritten: 8}
	_ac[42] = code{Code: 43 << 8, BitsWritten: 8}
	_ac[43] = code{Code: 44 << 8, BitsWritten: 8}
	_ac[44] = code{Code: 45 << 8, BitsWritten: 8}
	_ac[45] = code{Code: 4 << 8, BitsWritten: 8}
	_ac[46] = code{Code: 5 << 8, BitsWritten: 8}
	_ac[47] = code{Code: 10 << 8, BitsWritten: 8}
	_ac[48] = code{Code: 11 << 8, BitsWritten: 8}
	_ac[49] = code{Code: 82 << 8, BitsWritten: 8}
	_ac[50] = code{Code: 83 << 8, BitsWritten: 8}
	_ac[51] = code{Code: 84 << 8, BitsWritten: 8}
	_ac[52] = code{Code: 85 << 8, BitsWritten: 8}
	_ac[53] = code{Code: 36 << 8, BitsWritten: 8}
	_ac[54] = code{Code: 37 << 8, BitsWritten: 8}
	_ac[55] = code{Code: 88 << 8, BitsWritten: 8}
	_ac[56] = code{Code: 89 << 8, BitsWritten: 8}
	_ac[57] = code{Code: 90 << 8, BitsWritten: 8}
	_ac[58] = code{Code: 91 << 8, BitsWritten: 8}
	_ac[59] = code{Code: 74 << 8, BitsWritten: 8}
	_ac[60] = code{Code: 75 << 8, BitsWritten: 8}
	_ac[61] = code{Code: 50 << 8, BitsWritten: 8}
	_ac[62] = code{Code: 51 << 8, BitsWritten: 8}
	_ac[63] = code{Code: 52 << 8, BitsWritten: 8}
	_ada = make(map[int]code)
	_ada[64] = code{Code: 3<<8 | 3<<6, BitsWritten: 10}
	_ada[128] = code{Code: 12<<8 | 8<<4, BitsWritten: 12}
	_ada[192] = code{Code: 12<<8 | 9<<4, BitsWritten: 12}
	_ada[256] = code{Code: 5<<8 | 11<<4, BitsWritten: 12}
	_ada[320] = code{Code: 3<<8 | 3<<4, BitsWritten: 12}
	_ada[384] = code{Code: 3<<8 | 4<<4, BitsWritten: 12}
	_ada[448] = code{Code: 3<<8 | 5<<4, BitsWritten: 12}
	_ada[512] = code{Code: 3<<8 | 12<<3, BitsWritten: 13}
	_ada[576] = code{Code: 3<<8 | 13<<3, BitsWritten: 13}
	_ada[640] = code{Code: 2<<8 | 10<<3, BitsWritten: 13}
	_ada[704] = code{Code: 2<<8 | 11<<3, BitsWritten: 13}
	_ada[768] = code{Code: 2<<8 | 12<<3, BitsWritten: 13}
	_ada[832] = code{Code: 2<<8 | 13<<3, BitsWritten: 13}
	_ada[896] = code{Code: 3<<8 | 18<<3, BitsWritten: 13}
	_ada[960] = code{Code: 3<<8 | 19<<3, BitsWritten: 13}
	_ada[1024] = code{Code: 3<<8 | 20<<3, BitsWritten: 13}
	_ada[1088] = code{Code: 3<<8 | 21<<3, BitsWritten: 13}
	_ada[1152] = code{Code: 3<<8 | 22<<3, BitsWritten: 13}
	_ada[1216] = code{Code: 119 << 3, BitsWritten: 13}
	_ada[1280] = code{Code: 2<<8 | 18<<3, BitsWritten: 13}
	_ada[1344] = code{Code: 2<<8 | 19<<3, BitsWritten: 13}
	_ada[1408] = code{Code: 2<<8 | 20<<3, BitsWritten: 13}
	_ada[1472] = code{Code: 2<<8 | 21<<3, BitsWritten: 13}
	_ada[1536] = code{Code: 2<<8 | 26<<3, BitsWritten: 13}
	_ada[1600] = code{Code: 2<<8 | 27<<3, BitsWritten: 13}
	_ada[1664] = code{Code: 3<<8 | 4<<3, BitsWritten: 13}
	_ada[1728] = code{Code: 3<<8 | 5<<3, BitsWritten: 13}
	_fg = make(map[int]code)
	_fg[64] = code{Code: 27 << (3 + 8), BitsWritten: 5}
	_fg[128] = code{Code: 18 << (3 + 8), BitsWritten: 5}
	_fg[192] = code{Code: 23 << (2 + 8), BitsWritten: 6}
	_fg[256] = code{Code: 55 << (1 + 8), BitsWritten: 7}
	_fg[320] = code{Code: 54 << 8, BitsWritten: 8}
	_fg[384] = code{Code: 55 << 8, BitsWritten: 8}
	_fg[448] = code{Code: 100 << 8, BitsWritten: 8}
	_fg[512] = code{Code: 101 << 8, BitsWritten: 8}
	_fg[576] = code{Code: 104 << 8, BitsWritten: 8}
	_fg[640] = code{Code: 103 << 8, BitsWritten: 8}
	_fg[704] = code{Code: 102 << 8, BitsWritten: 9}
	_fg[768] = code{Code: 102<<8 | 1<<7, BitsWritten: 9}
	_fg[832] = code{Code: 105 << 8, BitsWritten: 9}
	_fg[896] = code{Code: 105<<8 | 1<<7, BitsWritten: 9}
	_fg[960] = code{Code: 106 << 8, BitsWritten: 9}
	_fg[1024] = code{Code: 106<<8 | 1<<7, BitsWritten: 9}
	_fg[1088] = code{Code: 107 << 8, BitsWritten: 9}
	_fg[1152] = code{Code: 107<<8 | 1<<7, BitsWritten: 9}
	_fg[1216] = code{Code: 108 << 8, BitsWritten: 9}
	_fg[1280] = code{Code: 108<<8 | 1<<7, BitsWritten: 9}
	_fg[1344] = code{Code: 109 << 8, BitsWritten: 9}
	_fg[1408] = code{Code: 109<<8 | 1<<7, BitsWritten: 9}
	_fg[1472] = code{Code: 76 << 8, BitsWritten: 9}
	_fg[1536] = code{Code: 76<<8 | 1<<7, BitsWritten: 9}
	_fg[1600] = code{Code: 77 << 8, BitsWritten: 9}
	_fg[1664] = code{Code: 24 << (2 + 8), BitsWritten: 6}
	_fg[1728] = code{Code: 77<<8 | 1<<7, BitsWritten: 9}
	_ggg = make(map[int]code)
	_ggg[1792] = code{Code: 1 << 8, BitsWritten: 11}
	_ggg[1856] = code{Code: 1<<8 | 4<<5, BitsWritten: 11}
	_ggg[1920] = code{Code: 1<<8 | 5<<5, BitsWritten: 11}
	_ggg[1984] = code{Code: 1<<8 | 2<<4, BitsWritten: 12}
	_ggg[2048] = code{Code: 1<<8 | 3<<4, BitsWritten: 12}
	_ggg[2112] = code{Code: 1<<8 | 4<<4, BitsWritten: 12}
	_ggg[2176] = code{Code: 1<<8 | 5<<4, BitsWritten: 12}
	_ggg[2240] = code{Code: 1<<8 | 6<<4, BitsWritten: 12}
	_ggg[2304] = code{Code: 1<<8 | 7<<4, BitsWritten: 12}
	_ggg[2368] = code{Code: 1<<8 | 12<<4, BitsWritten: 12}
	_ggg[2432] = code{Code: 1<<8 | 13<<4, BitsWritten: 12}
	_ggg[2496] = code{Code: 1<<8 | 14<<4, BitsWritten: 12}
	_ggg[2560] = code{Code: 1<<8 | 15<<4, BitsWritten: 12}
	_gda = make(map[int]byte)
	_gda[0] = 0xFF
	_gda[1] = 0xFE
	_gda[2] = 0xFC
	_gda[3] = 0xF8
	_gda[4] = 0xF0
	_gda[5] = 0xE0
	_gda[6] = 0xC0
	_gda[7] = 0x80
	_gda[8] = 0x00
}
func _cdc(_gbe []byte, _bfd int, _cbe code) ([]byte, int) {
	_dag := true
	var _dagc []byte
	_dagc, _bfd = _bgb(nil, _bfd, _cbe)
	_ccd := 0
	var _gfg int
	for _ccd < len(_gbe) {
		_gfg, _ccd = _edgb(_gbe, _dag, _ccd)
		_dagc, _bfd = _gfc(_dagc, _bfd, _gfg, _dag)
		_dag = !_dag
	}
	return _dagc, _bfd % 8
}
func _bgb(_dfb []byte, _ega int, _aaeg code) ([]byte, int) {
	_ecf := 0
	for _ecf < _aaeg.BitsWritten {
		_efcg := _ega / 8
		_beag := _ega % 8
		if _efcg >= len(_dfb) {
			_dfb = append(_dfb, 0)
		}
		_adf := 8 - _beag
		_affa := _aaeg.BitsWritten - _ecf
		if _adf > _affa {
			_adf = _affa
		}
		if _ecf < 8 {
			_dfb[_efcg] = _dfb[_efcg] | byte(_aaeg.Code>>uint(8+_beag-_ecf))&_gda[8-_adf-_beag]
		} else {
			_dfb[_efcg] = _dfb[_efcg] | (byte(_aaeg.Code<<uint(_ecf-8))&_gda[8-_adf])>>uint(_beag)
		}
		_ega += _adf
		_ecf += _adf
	}
	return _dfb, _ega
}
func (_aec *Decoder) tryFetchRTC2D() (_ffcg error) {
	_aec._ef.Mark()
	var _dac bool
	for _fcaf := 0; _fcaf < 5; _fcaf++ {
		_dac, _ffcg = _aec.tryFetchEOL1()
		if _ffcg != nil {
			if _g.Is(_ffcg, _a.EOF) {
				if _fcaf == 0 {
					break
				}
				return _gff
			}
		}
		if _dac {
			continue
		}
		if _fcaf > 0 {
			return _gff
		}
		break
	}
	if _dac {
		return _a.EOF
	}
	_aec._ef.Reset()
	return _ffcg
}
func (_ecd *Decoder) decoderRowType41D() error {
	if _ecd._fde {
		_ecd._ef.Align()
	}
	_ecd._ef.Mark()
	var (
		_fca bool
		_edg error
	)
	if _ecd._ce {
		_fca, _edg = _ecd.tryFetchEOL()
		if _edg != nil {
			return _edg
		}
		if !_fca {
			return _dbd
		}
	} else {
		_fca, _edg = _ecd.looseFetchEOL()
		if _edg != nil {
			return _edg
		}
	}
	if !_fca {
		_ecd._ef.Reset()
	}
	if _fca && _ecd._ea {
		_ecd._ef.Mark()
		for _aab := 0; _aab < 5; _aab++ {
			_fca, _edg = _ecd.tryFetchEOL()
			if _edg != nil {
				if _g.Is(_edg, _a.EOF) {
					if _aab == 0 {
						break
					}
					return _gff
				}
			}
			if _fca {
				continue
			}
			if _aab > 0 {
				return _gff
			}
			break
		}
		if _fca {
			return _a.EOF
		}
		_ecd._ef.Reset()
	}
	if _edg = _ecd.decode1D(); _edg != nil {
		return _edg
	}
	return nil
}

var _aa = [...][]uint16{{3, 2}, {1, 4}, {6, 5}, {7}, {9, 8}, {10, 11, 12}, {13, 14}, {15}, {16, 17, 0, 18, 64}, {24, 25, 23, 22, 19, 20, 21, 1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560, 52, 55, 56, 59, 60, 320, 384, 448, 53, 54, 50, 51, 44, 45, 46, 47, 57, 58, 61, 256, 48, 49, 62, 63, 30, 31, 32, 33, 40, 41, 128, 192, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 42, 43}, {640, 704, 768, 832, 1280, 1344, 1408, 1472, 1536, 1600, 1664, 1728, 512, 576, 896, 960, 1024, 1088, 1152, 1216}}

func (_egd *Encoder) encodeG32D(_gaf [][]byte) []byte {
	var _feb []byte
	var _facf int
	for _bdc := 0; _bdc < len(_gaf); _bdc += _egd.K {
		if _egd.Rows > 0 && !_egd.EndOfBlock && _bdc == _egd.Rows {
			break
		}
		_cee, _dba := _cdc(_gaf[_bdc], _facf, _cbg)
		_feb = _egd.appendEncodedRow(_feb, _cee, _facf)
		if _egd.EncodedByteAlign {
			_dba = 0
		}
		_facf = _dba
		for _bea := _bdc + 1; _bea < (_bdc+_egd.K) && _bea < len(_gaf); _bea++ {
			if _egd.Rows > 0 && !_egd.EndOfBlock && _bea == _egd.Rows {
				break
			}
			_ggbe, _cfe := _bgb(nil, _facf, _de)
			var _fdd, _ffb, _bfc int
			_ffg := -1
			for _ffg < len(_gaf[_bea]) {
				_fdd = _ecfd(_gaf[_bea], _ffg)
				_ffb = _ebd(_gaf[_bea], _gaf[_bea-1], _ffg)
				_bfc = _ecfd(_gaf[_bea-1], _ffb)
				if _bfc < _fdd {
					_ggbe, _cfe = _ede(_ggbe, _cfe)
					_ffg = _bfc
				} else {
					if _gd.Abs(float64(_ffb-_fdd)) > 3 {
						_ggbe, _cfe, _ffg = _eada(_gaf[_bea], _ggbe, _cfe, _ffg, _fdd)
					} else {
						_ggbe, _cfe = _acab(_ggbe, _cfe, _fdd, _ffb)
						_ffg = _fdd
					}
				}
			}
			_feb = _egd.appendEncodedRow(_feb, _ggbe, _facf)
			if _egd.EncodedByteAlign {
				_cfe = 0
			}
			_facf = _cfe % 8
		}
	}
	if _egd.EndOfBlock {
		_geb, _ := _accf(_facf)
		_feb = _egd.appendEncodedRow(_feb, _geb, _facf)
	}
	return _feb
}
func (_bee *Decoder) decodeG32D() error {
	_bee._caed = _bee._gce
	_bee._gdac, _bee._ed = _bee._ed, _bee._gdac
	_bff := true
	var (
		_gae  bool
		_gdab int
		_cfd  error
	)
	_bee._gce = 0
_fef:
	for _gdab < _bee._fb {
		_dfc := _fe._fbff
		for {
			_gae, _cfd = _bee._ef.ReadBool()
			if _cfd != nil {
				return _cfd
			}
			_dfc = _dfc.walk(_gae)
			if _dfc == nil {
				continue _fef
			}
			if !_dfc._bce {
				continue
			}
			switch _dfc._bfe {
			case _caf:
				var _abd int
				if _bff {
					_abd, _cfd = _bee.decodeRun(_e)
				} else {
					_abd, _cfd = _bee.decodeRun(_cf)
				}
				if _cfd != nil {
					return _cfd
				}
				_gdab += _abd
				_bee._gdac[_bee._gce] = _gdab
				_bee._gce++
				if _bff {
					_abd, _cfd = _bee.decodeRun(_cf)
				} else {
					_abd, _cfd = _bee.decodeRun(_e)
				}
				if _cfd != nil {
					return _cfd
				}
				_gdab += _abd
				_bee._gdac[_bee._gce] = _gdab
				_bee._gce++
			case _ee:
				_bad := _bee.getNextChangingElement(_gdab, _bff) + 1
				if _bad >= _bee._caed {
					_gdab = _bee._fb
				} else {
					_gdab = _bee._ed[_bad]
				}
			default:
				_beef := _bee.getNextChangingElement(_gdab, _bff)
				if _beef >= _bee._caed || _beef == -1 {
					_gdab = _bee._fb + _dfc._bfe
				} else {
					_gdab = _bee._ed[_beef] + _dfc._bfe
				}
				_bee._gdac[_bee._gce] = _gdab
				_bee._gce++
				_bff = !_bff
			}
			continue _fef
		}
	}
	return nil
}

var (
	_ceb byte = 1
	_afa byte = 0
)

func (_dd *Decoder) decodeRowType4() error {
	if !_dd._bg {
		return _dd.decoderRowType41D()
	}
	if _dd._fde {
		_dd._ef.Align()
	}
	_dd._ef.Mark()
	_cdg, _be := _dd.tryFetchEOL()
	if _be != nil {
		return _be
	}
	if !_cdg && _dd._ce {
		_dd._dgc++
		if _dd._dgc > _dd._cfca {
			return _dbd
		}
		_dd._ef.Reset()
	}
	if !_cdg {
		_dd._ef.Reset()
	}
	_agdc, _be := _dd._ef.ReadBool()
	if _be != nil {
		return _be
	}
	if _agdc {
		if _cdg && _dd._ea {
			if _be = _dd.tryFetchRTC2D(); _be != nil {
				return _be
			}
		}
		_be = _dd.decode1D()
	} else {
		_be = _dd.decode2D()
	}
	if _be != nil {
		return _be
	}
	return nil
}
func _ede(_adg []byte, _bed int) ([]byte, int) { return _bgb(_adg, _bed, _dge) }
func NewDecoder(data []byte, options DecodeOptions) (*Decoder, error) {
	_efa := &Decoder{_ef: _c.NewReader(data), _fb: options.Columns, _abg: options.Rows, _cfca: options.DamagedRowsBeforeError, _gfd: make([]byte, (options.Columns+7)/8), _ed: make([]int, options.Columns+2), _gdac: make([]int, options.Columns+2), _fde: options.EncodedByteAligned, _acc: options.BlackIsOne, _ce: options.EndOfLine, _ea: options.EndOfBlock}
	switch {
	case options.K == 0:
		_efa._aca = _db
		if len(data) < 20 {
			return nil, _g.New("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m")
		}
		_ced := data[:20]
		if _ced[0] != 0 || (_ced[1]>>4 != 1 && _ced[1] != 1) {
			_efa._aca = _ab
			_cdf := (uint16(_ced[0])<<8 + uint16(_ced[1]&0xff)) >> 4
			for _ceg := 12; _ceg < 160; _ceg++ {
				_cdf = (_cdf << 1) + uint16((_ced[_ceg/8]>>uint16(7-(_ceg%8)))&0x01)
				if _cdf&0xfff == 1 {
					_efa._aca = _db
					break
				}
			}
		}
	case options.K < 0:
		_efa._aca = _gdf
	case options.K > 0:
		_efa._aca = _db
		_efa._bg = true
	}
	switch _efa._aca {
	case _ab, _db, _gdf:
	default:
		return nil, _g.New("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065")
	}
	return _efa, nil
}
func _fggc(_bafd, _dee []byte, _faf int, _aaff bool) int {
	_fbf := _ecfd(_dee, _faf)
	if _fbf < len(_dee) && (_faf == -1 && _dee[_fbf] == _ceb || _faf >= 0 && _faf < len(_bafd) && _bafd[_faf] == _dee[_fbf] || _faf >= len(_bafd) && _aaff && _dee[_fbf] == _ceb || _faf >= len(_bafd) && !_aaff && _dee[_fbf] == _afa) {
		_fbf = _ecfd(_dee, _fbf)
	}
	return _fbf
}

var _bb = [...][]uint16{{0x2, 0x3}, {0x2, 0x3}, {0x2, 0x3}, {0x3}, {0x4, 0x5}, {0x4, 0x5, 0x7}, {0x4, 0x7}, {0x18}, {0x17, 0x18, 0x37, 0x8, 0xf}, {0x17, 0x18, 0x28, 0x37, 0x67, 0x68, 0x6c, 0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f, 0x24, 0x27, 0x28, 0x2b, 0x2c, 0x33, 0x34, 0x35, 0x37, 0x38, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xda, 0xdb}, {0x4a, 0x4b, 0x4c, 0x4d, 0x52, 0x53, 0x54, 0x55, 0x5a, 0x5b, 0x64, 0x65, 0x6c, 0x6d, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77}}

func _ecfd(_dgb []byte, _aecg int) int {
	if _aecg >= len(_dgb) {
		return _aecg
	}
	if _aecg < -1 {
		_aecg = -1
	}
	var _gcd byte
	if _aecg > -1 {
		_gcd = _dgb[_aecg]
	} else {
		_gcd = _ceb
	}
	_bgfc := _aecg + 1
	for _bgfc < len(_dgb) {
		if _dgb[_bgfc] != _gcd {
			break
		}
		_bgfc++
	}
	return _bgfc
}
func (_bdee *Decoder) getNextChangingElement(_egf int, _cfbf bool) int {
	_cc := 0
	if !_cfbf {
		_cc = 1
	}
	_faa := int(uint32(_bdee._ebg)&0xFFFFFFFE) + _cc
	if _faa > 2 {
		_faa -= 2
	}
	if _egf == 0 {
		return _faa
	}
	for _cbgb := _faa; _cbgb < _bdee._caed; _cbgb += 2 {
		if _egf < _bdee._ed[_cbgb] {
			_bdee._ebg = _cbgb
			return _cbgb
		}
	}
	return -1
}

var (
	_ggb map[int]code
	_ac  map[int]code
	_ada map[int]code
	_fg  map[int]code
	_ggg map[int]code
	_gda map[int]byte
	_bbf = code{Code: 1 << 4, BitsWritten: 12}
	_cbg = code{Code: 3 << 3, BitsWritten: 13}
	_de  = code{Code: 2 << 3, BitsWritten: 13}
	_dge = code{Code: 1 << 12, BitsWritten: 4}
	_gc  = code{Code: 1 << 13, BitsWritten: 3}
	_fgf = code{Code: 1 << 15, BitsWritten: 1}
	_ffc = code{Code: 3 << 13, BitsWritten: 3}
	_cfc = code{Code: 3 << 10, BitsWritten: 6}
	_ae  = code{Code: 3 << 9, BitsWritten: 7}
	_fgc = code{Code: 2 << 13, BitsWritten: 3}
	_bd  = code{Code: 2 << 10, BitsWritten: 6}
	_acf = code{Code: 2 << 9, BitsWritten: 7}
)

func (_cddf *Decoder) decodeRow() (_agd error) {
	if !_cddf._ea && _cddf._abg > 0 && _cddf._abg == _cddf._gb {
		return _a.EOF
	}
	switch _cddf._aca {
	case _ab:
		_agd = _cddf.decodeRowType2()
	case _db:
		_agd = _cddf.decodeRowType4()
	case _gdf:
		_agd = _cddf.decodeRowType6()
	}
	if _agd != nil {
		return _agd
	}
	_bgf := 0
	_cde := true
	_cddf._ebg = 0
	for _gcc := 0; _gcc < _cddf._gce; _gcc++ {
		_ebge := _cddf._fb
		if _gcc != _cddf._gce {
			_ebge = _cddf._gdac[_gcc]
		}
		if _ebge > _cddf._fb {
			_ebge = _cddf._fb
		}
		_eef := _bgf / 8
		for _bgf%8 != 0 && _ebge-_bgf > 0 {
			var _ge byte
			if !_cde {
				_ge = 1 << uint(7-(_bgf%8))
			}
			_cddf._gfd[_eef] |= _ge
			_bgf++
		}
		if _bgf%8 == 0 {
			_eef = _bgf / 8
			var _fgg byte
			if !_cde {
				_fgg = 0xff
			}
			for _ebge-_bgf > 7 {
				_cddf._gfd[_eef] = _fgg
				_bgf += 8
				_eef++
			}
		}
		for _ebge-_bgf > 0 {
			if _bgf%8 == 0 {
				_cddf._gfd[_eef] = 0
			}
			var _deb byte
			if !_cde {
				_deb = 1 << uint(7-(_bgf%8))
			}
			_cddf._gfd[_eef] |= _deb
			_bgf++
		}
		_cde = !_cde
	}
	if _bgf != _cddf._fb {
		return _g.New("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068")
	}
	_cddf._bf = (_bgf + 7) / 8
	_cddf._gb++
	return nil
}
func _accf(_cebf int) ([]byte, int) {
	var _aff []byte
	for _acafd := 0; _acafd < 6; _acafd++ {
		_aff, _cebf = _bgb(_aff, _cebf, _cbg)
	}
	return _aff, _cebf % 8
}
func (_efb *Decoder) decode2D() error {
	_efb._caed = _efb._gce
	_efb._gdac, _efb._ed = _efb._ed, _efb._gdac
	_gbg := true
	var (
		_da   bool
		_cdfc int
		_bag  error
	)
	_efb._gce = 0
_cgc:
	for _cdfc < _efb._fb {
		_fac := _fe._fbff
		for {
			_da, _bag = _efb._ef.ReadBool()
			if _bag != nil {
				return _bag
			}
			_fac = _fac.walk(_da)
			if _fac == nil {
				continue _cgc
			}
			if !_fac._bce {
				continue
			}
			switch _fac._bfe {
			case _caf:
				var _gdb int
				if _gbg {
					_gdb, _bag = _efb.decodeRun(_e)
				} else {
					_gdb, _bag = _efb.decodeRun(_cf)
				}
				if _bag != nil {
					return _bag
				}
				_cdfc += _gdb
				_efb._gdac[_efb._gce] = _cdfc
				_efb._gce++
				if _gbg {
					_gdb, _bag = _efb.decodeRun(_cf)
				} else {
					_gdb, _bag = _efb.decodeRun(_e)
				}
				if _bag != nil {
					return _bag
				}
				_cdfc += _gdb
				_efb._gdac[_efb._gce] = _cdfc
				_efb._gce++
			case _ee:
				_eaa := _efb.getNextChangingElement(_cdfc, _gbg) + 1
				if _eaa >= _efb._caed {
					_cdfc = _efb._fb
				} else {
					_cdfc = _efb._ed[_eaa]
				}
			default:
				_cdeg := _efb.getNextChangingElement(_cdfc, _gbg)
				if _cdeg >= _efb._caed || _cdeg == -1 {
					_cdfc = _efb._fb + _fac._bfe
				} else {
					_cdfc = _efb._ed[_cdeg] + _fac._bfe
				}
				_efb._gdac[_efb._gce] = _cdfc
				_efb._gce++
				_gbg = !_gbg
			}
			continue _cgc
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
type code struct {
	Code        uint16
	BitsWritten int
}

func _cac(_ebb int, _cbdb bool) (code, int, bool) {
	if _ebb < 64 {
		if _cbdb {
			return _ac[_ebb], 0, true
		}
		return _ggb[_ebb], 0, true
	}
	_bdcc := _ebb / 64
	if _bdcc > 40 {
		return _ggg[2560], _ebb - 2560, false
	}
	if _bdcc > 27 {
		return _ggg[_bdcc*64], _ebb - _bdcc*64, false
	}
	if _cbdb {
		return _fg[_bdcc*64], _ebb - _bdcc*64, false
	}
	return _ada[_bdcc*64], _ebb - _bdcc*64, false
}

var _dg = [...][]uint16{{0x7, 0x8, 0xb, 0xc, 0xe, 0xf}, {0x12, 0x13, 0x14, 0x1b, 0x7, 0x8}, {0x17, 0x18, 0x2a, 0x2b, 0x3, 0x34, 0x35, 0x7, 0x8}, {0x13, 0x17, 0x18, 0x24, 0x27, 0x28, 0x2b, 0x3, 0x37, 0x4, 0x8, 0xc}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1a, 0x1b, 0x2, 0x24, 0x25, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x3, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x4, 0x4a, 0x4b, 0x5, 0x52, 0x53, 0x54, 0x55, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x67, 0x68, 0xa, 0xb}, {0x98, 0x99, 0x9a, 0x9b, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb}, {}, {0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f}}

func (_eda *Encoder) encodeG31D(_gdfb [][]byte) []byte {
	var _cecd []byte
	_efe := 0
	for _ddd := range _gdfb {
		if _eda.Rows > 0 && !_eda.EndOfBlock && _ddd == _eda.Rows {
			break
		}
		_gca, _bda := _cdc(_gdfb[_ddd], _efe, _bbf)
		_cecd = _eda.appendEncodedRow(_cecd, _gca, _efe)
		if _eda.EncodedByteAlign {
			_bda = 0
		}
		_efe = _bda
	}
	if _eda.EndOfBlock {
		_bac, _ := _cce(_efe)
		_cecd = _eda.appendEncodedRow(_cecd, _bac, _efe)
	}
	return _cecd
}
func (_cg *Decoder) decodeRowType2() error {
	if _cg._fde {
		_cg._ef.Align()
	}
	if _cdfg := _cg.decode1D(); _cdfg != nil {
		return _cdfg
	}
	return nil
}
func _gfc(_facff []byte, _gbb int, _cebd int, _bacf bool) ([]byte, int) {
	var (
		_cge code
		_bca bool
	)
	for !_bca {
		_cge, _cebd, _bca = _cac(_cebd, _bacf)
		_facff, _gbb = _bgb(_facff, _gbb, _cge)
	}
	return _facff, _gbb
}
func _eada(_cbc, _edged []byte, _dgce, _gagd, _gad int) ([]byte, int, int) {
	_dged := _ecfd(_cbc, _gad)
	_bacg := _gagd >= 0 && _cbc[_gagd] == _ceb || _gagd == -1
	_edged, _dgce = _bgb(_edged, _dgce, _gc)
	var _dcf int
	if _gagd > -1 {
		_dcf = _gad - _gagd
	} else {
		_dcf = _gad - _gagd - 1
	}
	_edged, _dgce = _gfc(_edged, _dgce, _dcf, _bacg)
	_bacg = !_bacg
	_efgg := _dged - _gad
	_edged, _dgce = _gfc(_edged, _dgce, _efgg, _bacg)
	_gagd = _dged
	return _edged, _dgce, _gagd
}
func (_beb *Encoder) Encode(pixels [][]byte) []byte {
	if _beb.BlackIs1 {
		_ceb = 0
		_afa = 1
	} else {
		_ceb = 1
		_afa = 0
	}
	if _beb.K == 0 {
		return _beb.encodeG31D(pixels)
	}
	if _beb.K > 0 {
		return _beb.encodeG32D(pixels)
	}
	if _beb.K < 0 {
		return _beb.encodeG4(pixels)
	}
	return nil
}
func _edgb(_bcb []byte, _cffg bool, _bbd int) (int, int) {
	_dgeb := 0
	for _bbd < len(_bcb) {
		if _cffg {
			if _bcb[_bbd] != _ceb {
				break
			}
		} else {
			if _bcb[_bbd] != _afa {
				break
			}
		}
		_dgeb++
		_bbd++
	}
	return _dgeb, _bbd
}

type treeNode struct {
	_adff *treeNode
	_age  *treeNode
	_bfe  int
	_fbgb bool
	_bce  bool
}
type tree struct{ _fbff *treeNode }

func (_beg *Decoder) tryFetchEOL() (bool, error) {
	_ccc, _acaf := _beg._ef.ReadBits(12)
	if _acaf != nil {
		return false, _acaf
	}
	return _ccc == 0x1, nil
}
func (_abf *Decoder) tryFetchEOL1() (bool, error) {
	_gdgb, _cca := _abf._ef.ReadBits(13)
	if _cca != nil {
		return false, _cca
	}
	return _gdgb == 0x3, nil
}
func (_dbb *Decoder) fetch() error {
	if _dbb._bf == -1 {
		return nil
	}
	if _dbb._bbc < _dbb._bf {
		return nil
	}
	_dbb._bf = 0
	_bba := _dbb.decodeRow()
	if _bba != nil {
		if !_g.Is(_bba, _a.EOF) {
			return _bba
		}
		if _dbb._bf != 0 {
			return _bba
		}
		_dbb._bf = -1
	}
	_dbb._bbc = 0
	return nil
}
func _fbd(_daed [][]byte) [][]byte {
	_dcb := make([]byte, len(_daed[0]))
	for _dfe := range _dcb {
		_dcb[_dfe] = _ceb
	}
	_daed = append(_daed, []byte{})
	for _aeg := len(_daed) - 1; _aeg > 0; _aeg-- {
		_daed[_aeg] = _daed[_aeg-1]
	}
	_daed[0] = _dcb
	return _daed
}
func _deef(_cdgd, _ccab int) code {
	var _dcc code
	switch _ccab - _cdgd {
	case -1:
		_dcc = _ffc
	case -2:
		_dcc = _cfc
	case -3:
		_dcc = _ae
	case 0:
		_dcc = _fgf
	case 1:
		_dcc = _fgc
	case 2:
		_dcc = _bd
	case 3:
		_dcc = _acf
	}
	return _dcc
}
func (_aeb *Decoder) decodeRun(_fed *tree) (int, error) {
	var _fefg int
	_bae := _fed._fbff
	for {
		_cecg, _efg := _aeb._ef.ReadBool()
		if _efg != nil {
			return 0, _efg
		}
		_bae = _bae.walk(_cecg)
		if _bae == nil {
			return 0, _g.New("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _bae._bce {
			_fefg += _bae._bfe
			switch {
			case _bae._bfe >= 64:
				_bae = _fed._fbff
			case _bae._bfe >= 0:
				return _fefg, nil
			default:
				return _aeb._fb, nil
			}
		}
	}
}
func (_edge *Decoder) looseFetchEOL() (bool, error) {
	_deg, _efc := _edge._ef.ReadBits(12)
	if _efc != nil {
		return false, _efc
	}
	switch _deg {
	case 0x1:
		return true, nil
	case 0x0:
		for {
			_bbe, _bc := _edge._ef.ReadBool()
			if _bc != nil {
				return false, _bc
			}
			if _bbe {
				return true, nil
			}
		}
	default:
		return false, nil
	}
}
func (_dbe *treeNode) set(_dbdb bool, _agea *treeNode) {
	if !_dbdb {
		_dbe._adff = _agea
	} else {
		_dbe._age = _agea
	}
}
func _acab(_dbad []byte, _gdgd, _edac, _fcc int) ([]byte, int) {
	_fge := _deef(_edac, _fcc)
	_dbad, _gdgd = _bgb(_dbad, _gdgd, _fge)
	return _dbad, _gdgd
}
func (_cea *Decoder) decodeRowType6() error {
	if _cea._fde {
		_cea._ef.Align()
	}
	if _cea._ea {
		_cea._ef.Mark()
		_gbf, _edc := _cea.tryFetchEOL()
		if _edc != nil {
			return _edc
		}
		if _gbf {
			_gbf, _edc = _cea.tryFetchEOL()
			if _edc != nil {
				return _edc
			}
			if _gbf {
				return _a.EOF
			}
		}
		_cea._ef.Reset()
	}
	return _cea.decode2D()
}

type tiffType int

var _gdg = [...][]uint16{{2, 3, 4, 5, 6, 7}, {128, 8, 9, 64, 10, 11}, {192, 1664, 16, 17, 13, 14, 15, 1, 12}, {26, 21, 28, 27, 18, 24, 25, 22, 256, 23, 20, 19}, {33, 34, 35, 36, 37, 38, 31, 32, 29, 53, 54, 39, 40, 41, 42, 43, 44, 30, 61, 62, 63, 0, 320, 384, 45, 59, 60, 46, 49, 50, 51, 52, 55, 56, 57, 58, 448, 512, 640, 576, 47, 48}, {1472, 1536, 1600, 1728, 704, 768, 832, 896, 960, 1024, 1088, 1152, 1216, 1280, 1344, 1408}, {}, {1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560}}

func _agdf(_cbd int) ([]byte, int) {
	var _dgcd []byte
	for _ecc := 0; _ecc < 2; _ecc++ {
		_dgcd, _cbd = _bgb(_dgcd, _cbd, _bbf)
	}
	return _dgcd, _cbd % 8
}
func (_cgce *tree) fillWithNode(_fcf, _bcg int, _gea *treeNode) error {
	_ccg := _cgce._fbff
	for _begb := 0; _begb < _fcf; _begb++ {
		_ffce := uint(_fcf - 1 - _begb)
		_edcg := ((_bcg >> _ffce) & 1) != 0
		_edgec := _ccg.walk(_edcg)
		if _edgec != nil {
			if _edgec._bce {
				return _g.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_ccg = _edgec
			continue
		}
		if _begb == _fcf-1 {
			_edgec = _gea
		} else {
			_edgec = &treeNode{}
		}
		if _bcg == 0 {
			_edgec._fbgb = true
		}
		_ccg.set(_edcg, _edgec)
		_ccg = _edgec
	}
	return nil
}
func (_decd *tree) fill(_cgd, _dcg, _cgg int) error {
	_eeg := _decd._fbff
	for _ggdd := 0; _ggdd < _cgd; _ggdd++ {
		_cbdd := _cgd - 1 - _ggdd
		_cafc := ((_dcg >> uint(_cbdd)) & 1) != 0
		_fdfe := _eeg.walk(_cafc)
		if _fdfe != nil {
			if _fdfe._bce {
				return _g.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_eeg = _fdfe
			continue
		}
		_fdfe = &treeNode{}
		if _ggdd == _cgd-1 {
			_fdfe._bfe = _cgg
			_fdfe._bce = true
		}
		if _dcg == 0 {
			_fdfe._fbgb = true
		}
		_eeg.set(_cafc, _fdfe)
		_eeg = _fdfe
	}
	return nil
}
func _cce(_gbc int) ([]byte, int) {
	var _fcbg []byte
	for _fgfc := 0; _fgfc < 6; _fgfc++ {
		_fcbg, _gbc = _bgb(_fcbg, _gbc, _bbf)
	}
	return _fcbg, _gbc % 8
}

const (
	_ tiffType = iota
	_ab
	_db
	_gdf
)

type Decoder struct {
	_fb   int
	_abg  int
	_gb   int
	_gfd  []byte
	_cfca int
	_bg   bool
	_dec  bool
	_bde  bool
	_acc  bool
	_ce   bool
	_ea   bool
	_fde  bool
	_bf   int
	_bbc  int
	_ed   []int
	_gdac []int
	_caed int
	_gce  int
	_dgc  int
	_ebg  int
	_ef   *_c.Reader
	_aca  tiffType
	_eg   error
}

func (_egff *Encoder) appendEncodedRow(_gdbc, _gef []byte, _gcae int) []byte {
	if len(_gdbc) > 0 && _gcae != 0 && !_egff.EncodedByteAlign {
		_gdbc[len(_gdbc)-1] = _gdbc[len(_gdbc)-1] | _gef[0]
		_gdbc = append(_gdbc, _gef[1:]...)
	} else {
		_gdbc = append(_gdbc, _gef...)
	}
	return _gdbc
}
func (_gcge *Decoder) decode1D() error {
	var (
		_cec  int
		_faef error
	)
	_caee := true
	_gcge._gce = 0
	for {
		var _baf int
		if _caee {
			_baf, _faef = _gcge.decodeRun(_e)
		} else {
			_baf, _faef = _gcge.decodeRun(_cf)
		}
		if _faef != nil {
			return _faef
		}
		_cec += _baf
		_gcge._gdac[_gcge._gce] = _cec
		_gcge._gce++
		_caee = !_caee
		if _cec >= _gcge._fb {
			break
		}
	}
	return nil
}
