package mmr

import (
	_b "errors"
	_a "fmt"
	_ge "io"

	_aa "bitbucket.org/shenghui0779/gopdf/common"
	_gea "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_f "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
)

type Decoder struct {
	_aga, _gf int
	_gg       *runData
	_cdd      []*code
	_gbc      []*code
	_ea       []*code
}

func (_cde *Decoder) detectAndSkipEOL() error {
	for {
		_bgb, _bea := _cde._gg.uncompressGetCode(_cde._ea)
		if _bea != nil {
			return _bea
		}
		if _bgb != nil && _bgb._d == EOL {
			_cde._gg._bbbe += _bgb._e
		} else {
			return nil
		}
	}
}

type runData struct {
	_dff  *_gea.SubstreamReader
	_bbbe int
	_bda  int
	_fba  int
	_cgd  []byte
	_afe  int
	_de   int
}

const (
	EOF  = -3
	_bgg = -2
	EOL  = -1
	_aab = 8
	_adg = (1 << _aab) - 1
	_ae  = 5
	_gb  = (1 << _ae) - 1
)

func (_ebg *runData) fillBuffer(_bae int) error {
	_ebg._afe = _bae
	_, _ecg := _ebg._dff.Seek(int64(_bae), _ge.SeekStart)
	if _ecg != nil {
		if _ecg == _ge.EOF {
			_aa.Log.Debug("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046")
			_ebg._de = -1
		} else {
			return _ecg
		}
	}
	if _ecg == nil {
		_ebg._de, _ecg = _ebg._dff.Read(_ebg._cgd)
		if _ecg != nil {
			if _ecg == _ge.EOF {
				_aa.Log.Trace("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046")
				_ebg._de = -1
			} else {
				return _ecg
			}
		}
	}
	if _ebg._de > -1 && _ebg._de < 3 {
		for _ebg._de < 3 {
			_fcd, _gbb := _ebg._dff.ReadByte()
			if _gbb != nil {
				if _gbb == _ge.EOF {
					_ebg._cgd[_ebg._de] = 0
				} else {
					return _gbb
				}
			} else {
				_ebg._cgd[_ebg._de] = _fcd & 0xFF
			}
			_ebg._de++
		}
	}
	_ebg._de -= 3
	if _ebg._de < 0 {
		_ebg._cgd = make([]byte, len(_ebg._cgd))
		_ebg._de = len(_ebg._cgd) - 3
	}
	return nil
}
func (_bbb *Decoder) initTables() (_fcc error) {
	if _bbb._cdd == nil {
		_bbb._cdd, _fcc = _bbb.createLittleEndianTable(_bf)
		if _fcc != nil {
			return
		}
		_bbb._gbc, _fcc = _bbb.createLittleEndianTable(_bcg)
		if _fcc != nil {
			return
		}
		_bbb._ea, _fcc = _bbb.createLittleEndianTable(_bd)
		if _fcc != nil {
			return
		}
	}
	return nil
}
func (_gega *Decoder) uncompress1d(_bec *runData, _egf []int, _cc int) (int, error) {
	var (
		_cbb  = true
		_ff   int
		_fcca *code
		_bbe  int
		_gbg  error
	)
_agdd:
	for _ff < _cc {
	_fbd:
		for {
			if _cbb {
				_fcca, _gbg = _bec.uncompressGetCode(_gega._cdd)
				if _gbg != nil {
					return 0, _gbg
				}
			} else {
				_fcca, _gbg = _bec.uncompressGetCode(_gega._gbc)
				if _gbg != nil {
					return 0, _gbg
				}
			}
			_bec._bbbe += _fcca._e
			if _fcca._d < 0 {
				break _agdd
			}
			_ff += _fcca._d
			if _fcca._d < 64 {
				_cbb = !_cbb
				_egf[_bbe] = _ff
				_bbe++
				break _fbd
			}
		}
	}
	if _egf[_bbe] != _cc {
		_egf[_bbe] = _cc
	}
	_gag := EOL
	if _fcca != nil && _fcca._d != EOL {
		_gag = _bbe
	}
	return _gag, nil
}
func (_fbee *runData) uncompressGetCode(_ccb []*code) (*code, error) {
	return _fbee.uncompressGetCodeLittleEndian(_ccb)
}
func _aaf(_bc, _ad int) int {
	if _bc > _ad {
		return _ad
	}
	return _bc
}
func _fge(_fd [3]int) *code  { return &code{_e: _fd[0], _eg: _fd[1], _d: _fd[2]} }
func (_fcf *runData) align() { _fcf._bbbe = ((_fcf._bbbe + 7) >> 3) << 3 }
func (_ee *Decoder) fillBitmap(_ed *_f.Bitmap, _gdd int, _geg []int, _fcg int) error {
	var _cg byte
	_fa := 0
	_dfd := _ed.GetByteIndex(_fa, _gdd)
	for _ec := 0; _ec < _fcg; _ec++ {
		_edc := byte(1)
		_cb := _geg[_ec]
		if (_ec & 1) == 0 {
			_edc = 0
		}
		for _fa < _cb {
			_cg = (_cg << 1) | _edc
			_fa++
			if (_fa & 7) == 0 {
				if _agf := _ed.SetByte(_dfd, _cg); _agf != nil {
					return _agf
				}
				_dfd++
				_cg = 0
			}
		}
	}
	if (_fa & 7) != 0 {
		_cg <<= uint(8 - (_fa & 7))
		if _bgc := _ed.SetByte(_dfd, _cg); _bgc != nil {
			return _bgc
		}
	}
	return nil
}

const (
	_dc mmrCode = iota
	_cf
	_dg
	_dd
	_cd
	_bga
	_af
	_ag
	_dgd
	_cfd
	_ga
)

func (_ffg *runData) uncompressGetNextCodeLittleEndian() (int, error) {
	_bgca := _ffg._bbbe - _ffg._bda
	if _bgca < 0 || _bgca > 24 {
		_aecd := (_ffg._bbbe >> 3) - _ffg._afe
		if _aecd >= _ffg._de {
			_aecd += _ffg._afe
			if _feg := _ffg.fillBuffer(_aecd); _feg != nil {
				return 0, _feg
			}
			_aecd -= _ffg._afe
		}
		_fbg := (uint32(_ffg._cgd[_aecd]&0xFF) << 16) | (uint32(_ffg._cgd[_aecd+1]&0xFF) << 8) | (uint32(_ffg._cgd[_aecd+2] & 0xFF))
		_bcf := uint32(_ffg._bbbe & 7)
		_fbg <<= _bcf
		_ffg._fba = int(_fbg)
	} else {
		_acf := _ffg._bda & 7
		_bggc := 7 - _acf
		if _bgca <= _bggc {
			_ffg._fba <<= uint(_bgca)
		} else {
			_afb := (_ffg._bda >> 3) + 3 - _ffg._afe
			if _afb >= _ffg._de {
				_afb += _ffg._afe
				if _agg := _ffg.fillBuffer(_afb); _agg != nil {
					return 0, _agg
				}
				_afb -= _ffg._afe
			}
			_acf = 8 - _acf
			for {
				_ffg._fba <<= uint(_acf)
				_ffg._fba |= int(uint(_ffg._cgd[_afb]) & 0xFF)
				_bgca -= _acf
				_afb++
				_acf = 8
				if !(_bgca >= 8) {
					break
				}
			}
			_ffg._fba <<= uint(_bgca)
		}
	}
	_ffg._bda = _ffg._bbbe
	return _ffg._fba, nil
}
func (_cgf *runData) uncompressGetCodeLittleEndian(_adc []*code) (*code, error) {
	_ggc, _cab := _cgf.uncompressGetNextCodeLittleEndian()
	if _cab != nil {
		_aa.Log.Debug("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076", _cab)
		return nil, _cab
	}
	_ggc &= 0xffffff
	_fab := _ggc >> (_ede - _aab)
	_fcfe := _adc[_fab]
	if _fcfe != nil && _fcfe._bg {
		_fab = (_ggc >> (_ede - _aab - _ae)) & _gb
		_fcfe = _fcfe._fg[_fab]
	}
	return _fcfe, nil
}
func New(r _gea.StreamReader, width, height int, dataOffset, dataLength int64) (*Decoder, error) {
	_bdb := &Decoder{_aga: width, _gf: height}
	_gd, _gdf := _gea.NewSubstreamReader(r, uint64(dataOffset), uint64(dataLength))
	if _gdf != nil {
		return nil, _gdf
	}
	_cfdd, _gdf := _bfe(_gd)
	if _gdf != nil {
		return nil, _gdf
	}
	_bdb._gg = _cfdd
	if _be := _bdb.initTables(); _be != nil {
		return nil, _be
	}
	return _bdb, nil
}

const (
	_bacc int  = 1024 << 7
	_cgc  int  = 3
	_ede  uint = 24
)

func _db(_fgg, _c int) int {
	if _fgg < _c {
		return _c
	}
	return _fgg
}
func (_ac *code) String() string {
	return _a.Sprintf("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _ac._e, _ac._eg, _ac._d)
}
func (_fb *Decoder) UncompressMMR() (_fe *_f.Bitmap, _fda error) {
	_fe = _f.New(_fb._aga, _fb._gf)
	_ab := make([]int, _fe.Width+5)
	_fbe := make([]int, _fe.Width+5)
	_fbe[0] = _fe.Width
	_ba := 1
	var _bb int
	for _afc := 0; _afc < _fe.Height; _afc++ {
		_bb, _fda = _fb.uncompress2d(_fb._gg, _fbe, _ba, _ab, _fe.Width)
		if _fda != nil {
			return nil, _fda
		}
		if _bb == EOF {
			break
		}
		if _bb > 0 {
			_fda = _fb.fillBitmap(_fe, _afc, _ab, _bb)
			if _fda != nil {
				return nil, _fda
			}
		}
		_fbe, _ab = _ab, _fbe
		_ba = _bb
	}
	if _fda = _fb.detectAndSkipEOL(); _fda != nil {
		return nil, _fda
	}
	_fb._gg.align()
	return _fe, nil
}

type code struct {
	_e  int
	_eg int
	_d  int
	_fg []*code
	_bg bool
}

var (
	_bd  = [][3]int{{4, 0x1, int(_dc)}, {3, 0x1, int(_cf)}, {1, 0x1, int(_dg)}, {3, 0x3, int(_dd)}, {6, 0x3, int(_cd)}, {7, 0x3, int(_bga)}, {3, 0x2, int(_af)}, {6, 0x2, int(_ag)}, {7, 0x2, int(_dgd)}, {10, 0xf, int(_cfd)}, {12, 0xf, int(_ga)}, {12, 0x1, int(EOL)}}
	_bf  = [][3]int{{4, 0x07, 2}, {4, 0x08, 3}, {4, 0x0B, 4}, {4, 0x0C, 5}, {4, 0x0E, 6}, {4, 0x0F, 7}, {5, 0x12, 128}, {5, 0x13, 8}, {5, 0x14, 9}, {5, 0x1B, 64}, {5, 0x07, 10}, {5, 0x08, 11}, {6, 0x17, 192}, {6, 0x18, 1664}, {6, 0x2A, 16}, {6, 0x2B, 17}, {6, 0x03, 13}, {6, 0x34, 14}, {6, 0x35, 15}, {6, 0x07, 1}, {6, 0x08, 12}, {7, 0x13, 26}, {7, 0x17, 21}, {7, 0x18, 28}, {7, 0x24, 27}, {7, 0x27, 18}, {7, 0x28, 24}, {7, 0x2B, 25}, {7, 0x03, 22}, {7, 0x37, 256}, {7, 0x04, 23}, {7, 0x08, 20}, {7, 0xC, 19}, {8, 0x12, 33}, {8, 0x13, 34}, {8, 0x14, 35}, {8, 0x15, 36}, {8, 0x16, 37}, {8, 0x17, 38}, {8, 0x1A, 31}, {8, 0x1B, 32}, {8, 0x02, 29}, {8, 0x24, 53}, {8, 0x25, 54}, {8, 0x28, 39}, {8, 0x29, 40}, {8, 0x2A, 41}, {8, 0x2B, 42}, {8, 0x2C, 43}, {8, 0x2D, 44}, {8, 0x03, 30}, {8, 0x32, 61}, {8, 0x33, 62}, {8, 0x34, 63}, {8, 0x35, 0}, {8, 0x36, 320}, {8, 0x37, 384}, {8, 0x04, 45}, {8, 0x4A, 59}, {8, 0x4B, 60}, {8, 0x5, 46}, {8, 0x52, 49}, {8, 0x53, 50}, {8, 0x54, 51}, {8, 0x55, 52}, {8, 0x58, 55}, {8, 0x59, 56}, {8, 0x5A, 57}, {8, 0x5B, 58}, {8, 0x64, 448}, {8, 0x65, 512}, {8, 0x67, 640}, {8, 0x68, 576}, {8, 0x0A, 47}, {8, 0x0B, 48}, {9, 0x01, _bgg}, {9, 0x98, 1472}, {9, 0x99, 1536}, {9, 0x9A, 1600}, {9, 0x9B, 1728}, {9, 0xCC, 704}, {9, 0xCD, 768}, {9, 0xD2, 832}, {9, 0xD3, 896}, {9, 0xD4, 960}, {9, 0xD5, 1024}, {9, 0xD6, 1088}, {9, 0xD7, 1152}, {9, 0xD8, 1216}, {9, 0xD9, 1280}, {9, 0xDA, 1344}, {9, 0xDB, 1408}, {10, 0x01, _bgg}, {11, 0x01, _bgg}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}}
	_bcg = [][3]int{{2, 0x02, 3}, {2, 0x03, 2}, {3, 0x02, 1}, {3, 0x03, 4}, {4, 0x02, 6}, {4, 0x03, 5}, {5, 0x03, 7}, {6, 0x04, 9}, {6, 0x05, 8}, {7, 0x04, 10}, {7, 0x05, 11}, {7, 0x07, 12}, {8, 0x04, 13}, {8, 0x07, 14}, {9, 0x01, _bgg}, {9, 0x18, 15}, {10, 0x01, _bgg}, {10, 0x17, 16}, {10, 0x18, 17}, {10, 0x37, 0}, {10, 0x08, 18}, {10, 0x0F, 64}, {11, 0x01, _bgg}, {11, 0x17, 24}, {11, 0x18, 25}, {11, 0x28, 23}, {11, 0x37, 22}, {11, 0x67, 19}, {11, 0x68, 20}, {11, 0x6C, 21}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}, {12, 0x24, 52}, {12, 0x27, 55}, {12, 0x28, 56}, {12, 0x2B, 59}, {12, 0x2C, 60}, {12, 0x33, 320}, {12, 0x34, 384}, {12, 0x35, 448}, {12, 0x37, 53}, {12, 0x38, 54}, {12, 0x52, 50}, {12, 0x53, 51}, {12, 0x54, 44}, {12, 0x55, 45}, {12, 0x56, 46}, {12, 0x57, 47}, {12, 0x58, 57}, {12, 0x59, 58}, {12, 0x5A, 61}, {12, 0x5B, 256}, {12, 0x64, 48}, {12, 0x65, 49}, {12, 0x66, 62}, {12, 0x67, 63}, {12, 0x68, 30}, {12, 0x69, 31}, {12, 0x6A, 32}, {12, 0x6B, 33}, {12, 0x6C, 40}, {12, 0x6D, 41}, {12, 0xC8, 128}, {12, 0xC9, 192}, {12, 0xCA, 26}, {12, 0xCB, 27}, {12, 0xCC, 28}, {12, 0xCD, 29}, {12, 0xD2, 34}, {12, 0xD3, 35}, {12, 0xD4, 36}, {12, 0xD5, 37}, {12, 0xD6, 38}, {12, 0xD7, 39}, {12, 0xDA, 42}, {12, 0xDB, 43}, {13, 0x4A, 640}, {13, 0x4B, 704}, {13, 0x4C, 768}, {13, 0x4D, 832}, {13, 0x52, 1280}, {13, 0x53, 1344}, {13, 0x54, 1408}, {13, 0x55, 1472}, {13, 0x5A, 1536}, {13, 0x5B, 1600}, {13, 0x64, 1664}, {13, 0x65, 1728}, {13, 0x6C, 512}, {13, 0x6D, 576}, {13, 0x72, 896}, {13, 0x73, 960}, {13, 0x74, 1024}, {13, 0x75, 1088}, {13, 0x76, 1152}, {13, 0x77, 1216}}
)

func (_fc *Decoder) createLittleEndianTable(_df [][3]int) ([]*code, error) {
	_gda := make([]*code, _adg+1)
	for _bfb := 0; _bfb < len(_df); _bfb++ {
		_cfg := _fge(_df[_bfb])
		if _cfg._e <= _aab {
			_acc := _aab - _cfg._e
			_gaf := _cfg._eg << uint(_acc)
			for _aec := (1 << uint(_acc)) - 1; _aec >= 0; _aec-- {
				_bggb := _gaf | _aec
				_gda[_bggb] = _cfg
			}
		} else {
			_ca := _cfg._eg >> uint(_cfg._e-_aab)
			if _gda[_ca] == nil {
				var _agd = _fge([3]int{})
				_agd._fg = make([]*code, _gb+1)
				_gda[_ca] = _agd
			}
			if _cfg._e <= _aab+_ae {
				_bdc := _aab + _ae - _cfg._e
				_gfg := (_cfg._eg << uint(_bdc)) & _gb
				_gda[_ca]._bg = true
				for _bdg := (1 << uint(_bdc)) - 1; _bdg >= 0; _bdg-- {
					_gda[_ca]._fg[_gfg|_bdg] = _cfg
				}
			} else {
				return nil, _b.New("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072")
			}
		}
	}
	return _gda, nil
}

type mmrCode int

func (_eb *Decoder) uncompress2d(_egg *runData, _cff []int, _ef int, _gab []int, _adgc int) (int, error) {
	var (
		_acca int
		_fdc  int
		_gde  int
		_eff  = true
		_cbbd error
		_eef  *code
	)
	_cff[_ef] = _adgc
	_cff[_ef+1] = _adgc
	_cff[_ef+2] = _adgc + 1
	_cff[_ef+3] = _adgc + 1
_ega:
	for _gde < _adgc {
		_eef, _cbbd = _egg.uncompressGetCode(_eb._ea)
		if _cbbd != nil {
			return EOL, nil
		}
		if _eef == nil {
			_egg._bbbe++
			break _ega
		}
		_egg._bbbe += _eef._e
		switch mmrCode(_eef._d) {
		case _dg:
			_gde = _cff[_acca]
		case _dd:
			_gde = _cff[_acca] + 1
		case _af:
			_gde = _cff[_acca] - 1
		case _cf:
			for {
				var _becd []*code
				if _eff {
					_becd = _eb._cdd
				} else {
					_becd = _eb._gbc
				}
				_eef, _cbbd = _egg.uncompressGetCode(_becd)
				if _cbbd != nil {
					return 0, _cbbd
				}
				if _eef == nil {
					break _ega
				}
				_egg._bbbe += _eef._e
				if _eef._d < 64 {
					if _eef._d < 0 {
						_gab[_fdc] = _gde
						_fdc++
						_eef = nil
						break _ega
					}
					_gde += _eef._d
					_gab[_fdc] = _gde
					_fdc++
					break
				}
				_gde += _eef._d
			}
			_aea := _gde
		_cac:
			for {
				var _bac []*code
				if !_eff {
					_bac = _eb._cdd
				} else {
					_bac = _eb._gbc
				}
				_eef, _cbbd = _egg.uncompressGetCode(_bac)
				if _cbbd != nil {
					return 0, _cbbd
				}
				if _eef == nil {
					break _ega
				}
				_egg._bbbe += _eef._e
				if _eef._d < 64 {
					if _eef._d < 0 {
						_gab[_fdc] = _gde
						_fdc++
						break _ega
					}
					_gde += _eef._d
					if _gde < _adgc || _gde != _aea {
						_gab[_fdc] = _gde
						_fdc++
					}
					break _cac
				}
				_gde += _eef._d
			}
			for _gde < _adgc && _cff[_acca] <= _gde {
				_acca += 2
			}
			continue _ega
		case _dc:
			_acca++
			_gde = _cff[_acca]
			_acca++
			continue _ega
		case _cd:
			_gde = _cff[_acca] + 2
		case _ag:
			_gde = _cff[_acca] - 2
		case _bga:
			_gde = _cff[_acca] + 3
		case _dgd:
			_gde = _cff[_acca] - 3
		default:
			if _egg._bbbe == 12 && _eef._d == EOL {
				_egg._bbbe = 0
				if _, _cbbd = _eb.uncompress1d(_egg, _cff, _adgc); _cbbd != nil {
					return 0, _cbbd
				}
				_egg._bbbe++
				if _, _cbbd = _eb.uncompress1d(_egg, _gab, _adgc); _cbbd != nil {
					return 0, _cbbd
				}
				_dcc, _abd := _eb.uncompress1d(_egg, _cff, _adgc)
				if _abd != nil {
					return EOF, _abd
				}
				_egg._bbbe++
				return _dcc, nil
			}
			_gde = _adgc
			continue _ega
		}
		if _gde <= _adgc {
			_eff = !_eff
			_gab[_fdc] = _gde
			_fdc++
			if _acca > 0 {
				_acca--
			} else {
				_acca++
			}
			for _gde < _adgc && _cff[_acca] <= _gde {
				_acca += 2
			}
		}
	}
	if _gab[_fdc] != _adgc {
		_gab[_fdc] = _adgc
	}
	if _eef == nil {
		return EOL, nil
	}
	return _fdc, nil
}
func _bfe(_gge *_gea.SubstreamReader) (*runData, error) {
	_egb := &runData{_dff: _gge, _bbbe: 0, _bda: 1}
	_bfd := _aaf(_db(_cgc, int(_gge.Length())), _bacc)
	_egb._cgd = make([]byte, _bfd)
	if _dbf := _egb.fillBuffer(0); _dbf != nil {
		if _dbf == _ge.EOF {
			_egb._cgd = make([]byte, 10)
			_aa.Log.Debug("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076", _dbf)
		} else {
			return nil, _dbf
		}
	}
	return _egb, nil
}
