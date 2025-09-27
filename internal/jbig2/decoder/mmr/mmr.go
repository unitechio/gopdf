package mmr

import (
	_d "errors"
	_fd "fmt"
	_c "io"

	_fc "unitechio/gopdf/gopdf/common"
	_b "unitechio/gopdf/gopdf/internal/bitwise"
	_bf "unitechio/gopdf/gopdf/internal/jbig2/bitmap"
)

const (
	EOF  = -3
	_dgc = -2
	EOL  = -1
	_dde = 8
	_bg  = (1 << _dde) - 1
	_gd  = 5
	_ad  = (1 << _gd) - 1
)

func (_abf *Decoder) createLittleEndianTable(_cbb [][3]int) ([]*code, error) {
	_fed := make([]*code, _bg+1)
	for _ec := 0; _ec < len(_cbb); _ec++ {
		_ffb := _cf(_cbb[_ec])
		if _ffb._bfb <= _dde {
			_fba := _dde - _ffb._bfb
			_feda := _ffb._cb << uint(_fba)
			for _acf := (1 << uint(_fba)) - 1; _acf >= 0; _acf-- {
				_df := _feda | _acf
				_fed[_df] = _ffb
			}
		} else {
			_cc := _ffb._cb >> uint(_ffb._bfb-_dde)
			if _fed[_cc] == nil {
				_afa := _cf([3]int{})
				_afa._a = make([]*code, _ad+1)
				_fed[_cc] = _afa
			}
			if _ffb._bfb <= _dde+_gd {
				_gbd := _dde + _gd - _ffb._bfb
				_aad := (_ffb._cb << uint(_gbd)) & _ad
				_fed[_cc]._fa = true
				for _eb := (1 << uint(_gbd)) - 1; _eb >= 0; _eb-- {
					_fed[_cc]._a[_aad|_eb] = _ffb
				}
			} else {
				return nil, _d.New("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072")
			}
		}
	}
	return _fed, nil
}
func _cf(_da [3]int) *code { return &code{_bfb: _da[0], _cb: _da[1], _dd: _da[2]} }

type Decoder struct {
	_gdg, _bgf int
	_eg        *runData
	_fab       []*code
	_gb        []*code
	_ce        []*code
}

func (_acb *runData) uncompressGetNextCodeLittleEndian() (int, error) {
	_bea := _acb._afc - _acb._cbd
	if _bea < 0 || _bea > 24 {
		_ecgd := (_acb._afc >> 3) - _acb._ceg
		if _ecgd >= _acb._ddgg {
			_ecgd += _acb._ceg
			if _faf := _acb.fillBuffer(_ecgd); _faf != nil {
				return 0, _faf
			}
			_ecgd -= _acb._ceg
		}
		_fcc := (uint32(_acb._ddg[_ecgd]&0xFF) << 16) | (uint32(_acb._ddg[_ecgd+1]&0xFF) << 8) | (uint32(_acb._ddg[_ecgd+2] & 0xFF))
		_efaf := uint32(_acb._afc & 7)
		_fcc <<= _efaf
		_acb._abe = int(_fcc)
	} else {
		_bbc := _acb._cbd & 7
		_ccb := 7 - _bbc
		if _bea <= _ccb {
			_acb._abe <<= uint(_bea)
		} else {
			_faga := (_acb._cbd >> 3) + 3 - _acb._ceg
			if _faga >= _acb._ddgg {
				_faga += _acb._ceg
				if _aggb := _acb.fillBuffer(_faga); _aggb != nil {
					return 0, _aggb
				}
				_faga -= _acb._ceg
			}
			_bbc = 8 - _bbc
			for {
				_acb._abe <<= uint(_bbc)
				_acb._abe |= int(uint(_acb._ddg[_faga]) & 0xFF)
				_bea -= _bbc
				_faga++
				_bbc = 8
				if !(_bea >= 8) {
					break
				}
			}
			_acb._abe <<= uint(_bea)
		}
	}
	_acb._cbd = _acb._afc
	return _acb._abe, nil
}

func (_gc *Decoder) uncompress1d(_afg *runData, _afac []int, _fbb int) (int, error) {
	var (
		_dfe = true
		_fag int
		_fbe *code
		_gbf int
		_dbf error
	)
_ea:
	for _fag < _fbb {
	_fcae:
		for {
			if _dfe {
				_fbe, _dbf = _afg.uncompressGetCode(_gc._fab)
				if _dbf != nil {
					return 0, _dbf
				}
			} else {
				_fbe, _dbf = _afg.uncompressGetCode(_gc._gb)
				if _dbf != nil {
					return 0, _dbf
				}
			}
			_afg._afc += _fbe._bfb
			if _fbe._dd < 0 {
				break _ea
			}
			_fag += _fbe._dd
			if _fbe._dd < 64 {
				_dfe = !_dfe
				_afac[_gbf] = _fag
				_gbf++
				break _fcae
			}
		}
	}
	if _afac[_gbf] != _fbb {
		_afac[_gbf] = _fbb
	}
	_bef := EOL
	if _fbe != nil && _fbe._dd != EOL {
		_bef = _gbf
	}
	return _bef, nil
}

func (_ecc *Decoder) initTables() (_fg error) {
	if _ecc._fab == nil {
		_ecc._fab, _fg = _ecc.createLittleEndianTable(_ff)
		if _fg != nil {
			return
		}
		_ecc._gb, _fg = _ecc.createLittleEndianTable(_bga)
		if _fg != nil {
			return
		}
		_ecc._ce, _fg = _ecc.createLittleEndianTable(_ef)
		if _fg != nil {
			return
		}
	}
	return nil
}
func (_aed *runData) align() { _aed._afc = ((_aed._afc + 7) >> 3) << 3 }

type code struct {
	_bfb int
	_cb  int
	_dd  int
	_a   []*code
	_fa  bool
}

func New(r *_b.Reader, width, height int, dataOffset, dataLength int64) (*Decoder, error) {
	_feb := &Decoder{_gdg: width, _bgf: height}
	_fca, _bbe := r.NewPartialReader(int(dataOffset), int(dataLength), false)
	if _bbe != nil {
		return nil, _bbe
	}
	_aac, _bbe := _bdg(_fca)
	if _bbe != nil {
		return nil, _bbe
	}
	_, _bbe = r.Seek(_fca.RelativePosition(), _c.SeekCurrent)
	if _bbe != nil {
		return nil, _bbe
	}
	_feb._eg = _aac
	if _ade := _feb.initTables(); _ade != nil {
		return nil, _ade
	}
	return _feb, nil
}

func (_ddb *code) String() string {
	return _fd.Sprintf("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _ddb._bfb, _ddb._cb, _ddb._dd)
}

const (
	_ba mmrCode = iota
	_dg
	_g
	_bae
	_ag
	_ab
	_fe
	_bd
	_bb
	_bba
	_cd
)

func (_ga *Decoder) detectAndSkipEOL() error {
	for {
		_efa, _ddc := _ga._eg.uncompressGetCode(_ga._ce)
		if _ddc != nil {
			return _ddc
		}
		if _efa != nil && _efa._dd == EOL {
			_ga._eg._afc += _efa._bfb
		} else {
			return nil
		}
	}
}

func (_beg *Decoder) fillBitmap(_bdbb *_bf.Bitmap, _cde int, _de []int, _baf int) error {
	var _ecf byte
	_ccf := 0
	_dfg := _bdbb.GetByteIndex(_ccf, _cde)
	for _bff := 0; _bff < _baf; _bff++ {
		_ecg := byte(1)
		_bdc := _de[_bff]
		if (_bff & 1) == 0 {
			_ecg = 0
		}
		for _ccf < _bdc {
			_ecf = (_ecf << 1) | _ecg
			_ccf++
			if (_ccf & 7) == 0 {
				if _agb := _bdbb.SetByte(_dfg, _ecf); _agb != nil {
					return _agb
				}
				_dfg++
				_ecf = 0
			}
		}
	}
	if (_ccf & 7) != 0 {
		_ecf <<= uint(8 - (_ccf & 7))
		if _aab := _bdbb.SetByte(_dfg, _ecf); _aab != nil {
			return _aab
		}
	}
	return nil
}

func _fdb(_e, _aa int) int {
	if _e < _aa {
		return _aa
	}
	return _e
}

const (
	_dec int  = 1024 << 7
	_cac int  = 3
	_gaf uint = 24
)

func (_bgg *runData) uncompressGetCodeLittleEndian(_agg []*code) (*code, error) {
	_ebc, _efbe := _bgg.uncompressGetNextCodeLittleEndian()
	if _efbe != nil {
		_fc.Log.Debug("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076", _efbe)
		return nil, _efbe
	}
	_ebc &= 0xffffff
	_agf := _ebc >> (_gaf - _dde)
	_fea := _agg[_agf]
	if _fea != nil && _fea._fa {
		_agf = (_ebc >> (_gaf - _dde - _gd)) & _ad
		_fea = _fea._a[_agf]
	}
	return _fea, nil
}

func _bdg(_efb *_b.Reader) (*runData, error) {
	_fbg := &runData{_gaa: _efb, _afc: 0, _cbd: 1}
	_fcb := _db(_fdb(_cac, int(_efb.Length())), _dec)
	_fbg._ddg = make([]byte, _fcb)
	if _ae := _fbg.fillBuffer(0); _ae != nil {
		if _ae == _c.EOF {
			_fbg._ddg = make([]byte, 10)
			_fc.Log.Debug("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076", _ae)
		} else {
			return nil, _ae
		}
	}
	return _fbg, nil
}

func (_gf *Decoder) UncompressMMR() (_bdb *_bf.Bitmap, _fb error) {
	_bdb = _bf.New(_gf._gdg, _gf._bgf)
	_gg := make([]int, _bdb.Width+5)
	_gbg := make([]int, _bdb.Width+5)
	_gbg[0] = _bdb.Width
	_agc := 1
	var _be int
	for _af := 0; _af < _bdb.Height; _af++ {
		_be, _fb = _gf.uncompress2d(_gf._eg, _gbg, _agc, _gg, _bdb.Width)
		if _fb != nil {
			return nil, _fb
		}
		if _be == EOF {
			break
		}
		if _be > 0 {
			_fb = _gf.fillBitmap(_bdb, _af, _gg, _be)
			if _fb != nil {
				return nil, _fb
			}
		}
		_gbg, _gg = _gg, _gbg
		_agc = _be
	}
	if _fb = _gf.detectAndSkipEOL(); _fb != nil {
		return nil, _fb
	}
	_gf._eg.align()
	return _bdb, nil
}

func (_bab *runData) uncompressGetCode(_bfd []*code) (*code, error) {
	return _bab.uncompressGetCodeLittleEndian(_bfd)
}

func (_cfc *runData) fillBuffer(_adc int) error {
	_cfc._ceg = _adc
	_, _feg := _cfc._gaa.Seek(int64(_adc), _c.SeekStart)
	if _feg != nil {
		if _feg == _c.EOF {
			_fc.Log.Debug("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046")
			_cfc._ddgg = -1
		} else {
			return _feg
		}
	}
	if _feg == nil {
		_cfc._ddgg, _feg = _cfc._gaa.Read(_cfc._ddg)
		if _feg != nil {
			if _feg == _c.EOF {
				_fc.Log.Trace("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046")
				_cfc._ddgg = -1
			} else {
				return _feg
			}
		}
	}
	if _cfc._ddgg > -1 && _cfc._ddgg < 3 {
		for _cfc._ddgg < 3 {
			_efba, _ffbf := _cfc._gaa.ReadByte()
			if _ffbf != nil {
				if _ffbf == _c.EOF {
					_cfc._ddg[_cfc._ddgg] = 0
				} else {
					return _ffbf
				}
			} else {
				_cfc._ddg[_cfc._ddgg] = _efba & 0xFF
			}
			_cfc._ddgg++
		}
	}
	_cfc._ddgg -= 3
	if _cfc._ddgg < 0 {
		_cfc._ddg = make([]byte, len(_cfc._ddg))
		_cfc._ddgg = len(_cfc._ddg) - 3
	}
	return nil
}

func (_abd *Decoder) uncompress2d(_dgg *runData, _fdf []int, _ead int, _ggd []int, _ccc int) (int, error) {
	var (
		_bda  int
		_ebb  int
		_gab  int
		_bc   = true
		_agbe error
		_bgc  *code
	)
	_fdf[_ead] = _ccc
	_fdf[_ead+1] = _ccc
	_fdf[_ead+2] = _ccc + 1
	_fdf[_ead+3] = _ccc + 1
_fda:
	for _gab < _ccc {
		_bgc, _agbe = _dgg.uncompressGetCode(_abd._ce)
		if _agbe != nil {
			return EOL, nil
		}
		if _bgc == nil {
			_dgg._afc++
			break _fda
		}
		_dgg._afc += _bgc._bfb
		switch mmrCode(_bgc._dd) {
		case _g:
			_gab = _fdf[_bda]
		case _bae:
			_gab = _fdf[_bda] + 1
		case _fe:
			_gab = _fdf[_bda] - 1
		case _dg:
			for {
				var _dbd []*code
				if _bc {
					_dbd = _abd._fab
				} else {
					_dbd = _abd._gb
				}
				_bgc, _agbe = _dgg.uncompressGetCode(_dbd)
				if _agbe != nil {
					return 0, _agbe
				}
				if _bgc == nil {
					break _fda
				}
				_dgg._afc += _bgc._bfb
				if _bgc._dd < 64 {
					if _bgc._dd < 0 {
						_ggd[_ebb] = _gab
						_ebb++
						_bgc = nil
						break _fda
					}
					_gab += _bgc._dd
					_ggd[_ebb] = _gab
					_ebb++
					break
				}
				_gab += _bgc._dd
			}
			_add := _gab
		_dgbe:
			for {
				var _dfb []*code
				if !_bc {
					_dfb = _abd._fab
				} else {
					_dfb = _abd._gb
				}
				_bgc, _agbe = _dgg.uncompressGetCode(_dfb)
				if _agbe != nil {
					return 0, _agbe
				}
				if _bgc == nil {
					break _fda
				}
				_dgg._afc += _bgc._bfb
				if _bgc._dd < 64 {
					if _bgc._dd < 0 {
						_ggd[_ebb] = _gab
						_ebb++
						break _fda
					}
					_gab += _bgc._dd
					if _gab < _ccc || _gab != _add {
						_ggd[_ebb] = _gab
						_ebb++
					}
					break _dgbe
				}
				_gab += _bgc._dd
			}
			for _gab < _ccc && _fdf[_bda] <= _gab {
				_bda += 2
			}
			continue _fda
		case _ba:
			_bda++
			_gab = _fdf[_bda]
			_bda++
			continue _fda
		case _ag:
			_gab = _fdf[_bda] + 2
		case _bd:
			_gab = _fdf[_bda] - 2
		case _ab:
			_gab = _fdf[_bda] + 3
		case _bb:
			_gab = _fdf[_bda] - 3
		default:
			if _dgg._afc == 12 && _bgc._dd == EOL {
				_dgg._afc = 0
				if _, _agbe = _abd.uncompress1d(_dgg, _fdf, _ccc); _agbe != nil {
					return 0, _agbe
				}
				_dgg._afc++
				if _, _agbe = _abd.uncompress1d(_dgg, _ggd, _ccc); _agbe != nil {
					return 0, _agbe
				}
				_bge, _bed := _abd.uncompress1d(_dgg, _fdf, _ccc)
				if _bed != nil {
					return EOF, _bed
				}
				_dgg._afc++
				return _bge, nil
			}
			_gab = _ccc
			continue _fda
		}
		if _gab <= _ccc {
			_bc = !_bc
			_ggd[_ebb] = _gab
			_ebb++
			if _bda > 0 {
				_bda--
			} else {
				_bda++
			}
			for _gab < _ccc && _fdf[_bda] <= _gab {
				_bda += 2
			}
		}
	}
	if _ggd[_ebb] != _ccc {
		_ggd[_ebb] = _ccc
	}
	if _bgc == nil {
		return EOL, nil
	}
	return _ebb, nil
}

type (
	mmrCode int
	runData struct {
		_gaa  *_b.Reader
		_afc  int
		_cbd  int
		_abe  int
		_ddg  []byte
		_ceg  int
		_ddgg int
	}
)

var (
	_ef  = [][3]int{{4, 0x1, int(_ba)}, {3, 0x1, int(_dg)}, {1, 0x1, int(_g)}, {3, 0x3, int(_bae)}, {6, 0x3, int(_ag)}, {7, 0x3, int(_ab)}, {3, 0x2, int(_fe)}, {6, 0x2, int(_bd)}, {7, 0x2, int(_bb)}, {10, 0xf, int(_bba)}, {12, 0xf, int(_cd)}, {12, 0x1, int(EOL)}}
	_ff  = [][3]int{{4, 0x07, 2}, {4, 0x08, 3}, {4, 0x0B, 4}, {4, 0x0C, 5}, {4, 0x0E, 6}, {4, 0x0F, 7}, {5, 0x12, 128}, {5, 0x13, 8}, {5, 0x14, 9}, {5, 0x1B, 64}, {5, 0x07, 10}, {5, 0x08, 11}, {6, 0x17, 192}, {6, 0x18, 1664}, {6, 0x2A, 16}, {6, 0x2B, 17}, {6, 0x03, 13}, {6, 0x34, 14}, {6, 0x35, 15}, {6, 0x07, 1}, {6, 0x08, 12}, {7, 0x13, 26}, {7, 0x17, 21}, {7, 0x18, 28}, {7, 0x24, 27}, {7, 0x27, 18}, {7, 0x28, 24}, {7, 0x2B, 25}, {7, 0x03, 22}, {7, 0x37, 256}, {7, 0x04, 23}, {7, 0x08, 20}, {7, 0xC, 19}, {8, 0x12, 33}, {8, 0x13, 34}, {8, 0x14, 35}, {8, 0x15, 36}, {8, 0x16, 37}, {8, 0x17, 38}, {8, 0x1A, 31}, {8, 0x1B, 32}, {8, 0x02, 29}, {8, 0x24, 53}, {8, 0x25, 54}, {8, 0x28, 39}, {8, 0x29, 40}, {8, 0x2A, 41}, {8, 0x2B, 42}, {8, 0x2C, 43}, {8, 0x2D, 44}, {8, 0x03, 30}, {8, 0x32, 61}, {8, 0x33, 62}, {8, 0x34, 63}, {8, 0x35, 0}, {8, 0x36, 320}, {8, 0x37, 384}, {8, 0x04, 45}, {8, 0x4A, 59}, {8, 0x4B, 60}, {8, 0x5, 46}, {8, 0x52, 49}, {8, 0x53, 50}, {8, 0x54, 51}, {8, 0x55, 52}, {8, 0x58, 55}, {8, 0x59, 56}, {8, 0x5A, 57}, {8, 0x5B, 58}, {8, 0x64, 448}, {8, 0x65, 512}, {8, 0x67, 640}, {8, 0x68, 576}, {8, 0x0A, 47}, {8, 0x0B, 48}, {9, 0x01, _dgc}, {9, 0x98, 1472}, {9, 0x99, 1536}, {9, 0x9A, 1600}, {9, 0x9B, 1728}, {9, 0xCC, 704}, {9, 0xCD, 768}, {9, 0xD2, 832}, {9, 0xD3, 896}, {9, 0xD4, 960}, {9, 0xD5, 1024}, {9, 0xD6, 1088}, {9, 0xD7, 1152}, {9, 0xD8, 1216}, {9, 0xD9, 1280}, {9, 0xDA, 1344}, {9, 0xDB, 1408}, {10, 0x01, _dgc}, {11, 0x01, _dgc}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}}
	_bga = [][3]int{{2, 0x02, 3}, {2, 0x03, 2}, {3, 0x02, 1}, {3, 0x03, 4}, {4, 0x02, 6}, {4, 0x03, 5}, {5, 0x03, 7}, {6, 0x04, 9}, {6, 0x05, 8}, {7, 0x04, 10}, {7, 0x05, 11}, {7, 0x07, 12}, {8, 0x04, 13}, {8, 0x07, 14}, {9, 0x01, _dgc}, {9, 0x18, 15}, {10, 0x01, _dgc}, {10, 0x17, 16}, {10, 0x18, 17}, {10, 0x37, 0}, {10, 0x08, 18}, {10, 0x0F, 64}, {11, 0x01, _dgc}, {11, 0x17, 24}, {11, 0x18, 25}, {11, 0x28, 23}, {11, 0x37, 22}, {11, 0x67, 19}, {11, 0x68, 20}, {11, 0x6C, 21}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}, {12, 0x24, 52}, {12, 0x27, 55}, {12, 0x28, 56}, {12, 0x2B, 59}, {12, 0x2C, 60}, {12, 0x33, 320}, {12, 0x34, 384}, {12, 0x35, 448}, {12, 0x37, 53}, {12, 0x38, 54}, {12, 0x52, 50}, {12, 0x53, 51}, {12, 0x54, 44}, {12, 0x55, 45}, {12, 0x56, 46}, {12, 0x57, 47}, {12, 0x58, 57}, {12, 0x59, 58}, {12, 0x5A, 61}, {12, 0x5B, 256}, {12, 0x64, 48}, {12, 0x65, 49}, {12, 0x66, 62}, {12, 0x67, 63}, {12, 0x68, 30}, {12, 0x69, 31}, {12, 0x6A, 32}, {12, 0x6B, 33}, {12, 0x6C, 40}, {12, 0x6D, 41}, {12, 0xC8, 128}, {12, 0xC9, 192}, {12, 0xCA, 26}, {12, 0xCB, 27}, {12, 0xCC, 28}, {12, 0xCD, 29}, {12, 0xD2, 34}, {12, 0xD3, 35}, {12, 0xD4, 36}, {12, 0xD5, 37}, {12, 0xD6, 38}, {12, 0xD7, 39}, {12, 0xDA, 42}, {12, 0xDB, 43}, {13, 0x4A, 640}, {13, 0x4B, 704}, {13, 0x4C, 768}, {13, 0x4D, 832}, {13, 0x52, 1280}, {13, 0x53, 1344}, {13, 0x54, 1408}, {13, 0x55, 1472}, {13, 0x5A, 1536}, {13, 0x5B, 1600}, {13, 0x64, 1664}, {13, 0x65, 1728}, {13, 0x6C, 512}, {13, 0x6D, 576}, {13, 0x72, 896}, {13, 0x73, 960}, {13, 0x74, 1024}, {13, 0x75, 1088}, {13, 0x76, 1152}, {13, 0x77, 1216}}
)

func _db(_ca, _ac int) int {
	if _ca > _ac {
		return _ac
	}
	return _ca
}
