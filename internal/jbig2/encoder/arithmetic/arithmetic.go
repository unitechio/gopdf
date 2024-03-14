package arithmetic

import (
	_f "bytes"
	_fa "io"

	_bc "bitbucket.org/shenghui0779/gopdf/common"
	_ba "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_fe "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

type Encoder struct {
	_bg      uint32
	_gd      uint16
	_ed, _eb uint8
	_cb      int
	_fc      int
	_bd      [][]byte
	_fag     []byte
	_cf      int
	_cd      *codingContext
	_cfa     [13]*codingContext
	_fdd     *codingContext
}

func (_ff *Encoder) EncodeOOB(proc Class) (_fdad error) {
	_bc.Log.Trace("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027", proc)
	if _fdad = _ff.encodeOOB(proc); _fdad != nil {
		return _fe.Wrap(_fdad, "\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB", "")
	}
	return nil
}

var _ _fa.WriterTo = &Encoder{}

func (_ffg *Encoder) encodeInteger(_afe Class, _cff int) error {
	const _bagf = "E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072"
	if _cff > 2000000000 || _cff < -2000000000 {
		return _fe.Errorf(_bagf, "\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", _cff)
	}
	_dccd := _ffg._cfa[_afe]
	_ce := uint32(1)
	var _gecb int
	for ; ; _gecb++ {
		if _ag[_gecb]._c <= _cff && _ag[_gecb]._a >= _cff {
			break
		}
	}
	if _cff < 0 {
		_cff = -_cff
	}
	_cff -= int(_ag[_gecb]._fd)
	_ceb := _ag[_gecb]._fb
	for _ef := uint8(0); _ef < _ag[_gecb]._fac; _ef++ {
		_gbgb := _ceb & 1
		if _fddgb := _ffg.encodeBit(_dccd, _ce, _gbgb); _fddgb != nil {
			return _fe.Wrap(_fddgb, _bagf, "")
		}
		_ceb >>= 1
		if _ce&0x100 > 0 {
			_ce = (((_ce << 1) | uint32(_gbgb)) & 0x1ff) | 0x100
		} else {
			_ce = (_ce << 1) | uint32(_gbgb)
		}
	}
	_cff <<= 32 - _ag[_gecb]._ac
	for _badg := uint8(0); _badg < _ag[_gecb]._ac; _badg++ {
		_de := uint8((uint32(_cff) & 0x80000000) >> 31)
		if _fddc := _ffg.encodeBit(_dccd, _ce, _de); _fddc != nil {
			return _fe.Wrap(_fddc, _bagf, "\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064")
		}
		_cff <<= 1
		if _ce&0x100 != 0 {
			_ce = (((_ce << 1) | uint32(_de)) & 0x1ff) | 0x100
		} else {
			_ce = (_ce << 1) | uint32(_de)
		}
	}
	return nil
}

func (_baf *Encoder) renormalize() {
	for {
		_baf._gd <<= 1
		_baf._bg <<= 1
		_baf._ed--
		if _baf._ed == 0 {
			_baf.byteOut()
		}
		if (_baf._gd & 0x8000) != 0 {
			break
		}
	}
}

type codingContext struct {
	_d  []byte
	_ca []byte
}

func (_dfd *Encoder) lBlock() {
	if _dfd._cb >= 0 {
		_dfd.emit()
	}
	_dfd._cb++
	_dfd._eb = uint8(_dfd._bg >> 19)
	_dfd._bg &= 0x7ffff
	_dfd._ed = 8
}

func (_ab *Encoder) Init() {
	_ab._cd = _ec(_gae)
	_ab._gd = 0x8000
	_ab._bg = 0
	_ab._ed = 12
	_ab._cb = -1
	_ab._eb = 0
	_ab._cf = 0
	_ab._fag = make([]byte, _adae)
	for _cc := 0; _cc < len(_ab._cfa); _cc++ {
		_ab._cfa[_cc] = _ec(512)
	}
	_ab._fdd = nil
}

var _ag = []intEncRangeS{{0, 3, 0, 2, 0, 2}, {-1, -1, 9, 4, 0, 0}, {-3, -2, 5, 3, 2, 1}, {4, 19, 2, 3, 4, 4}, {-19, -4, 3, 3, 4, 4}, {20, 83, 6, 4, 20, 6}, {-83, -20, 7, 4, 20, 6}, {84, 339, 14, 5, 84, 8}, {-339, -84, 15, 5, 84, 8}, {340, 4435, 30, 6, 340, 12}, {-4435, -340, 31, 6, 340, 12}, {4436, 2000000000, 62, 6, 4436, 32}, {-2000000000, -4436, 63, 6, 4436, 32}}

func (_ebbd *Encoder) EncodeIAID(symbolCodeLength, value int) (_edc error) {
	_bc.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", symbolCodeLength, value)
	if _edc = _ebbd.encodeIAID(symbolCodeLength, value); _edc != nil {
		return _fe.Wrap(_edc, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044", "")
	}
	return nil
}

const _ebb = 0x9b25

func (_ffd *Encoder) setBits() {
	_daa := _ffd._bg + uint32(_ffd._gd)
	_ffd._bg |= 0xffff
	if _ffd._bg >= _daa {
		_ffd._bg -= 0x8000
	}
}

type intEncRangeS struct {
	_c, _a    int
	_fb, _fac uint8
	_fd       uint16
	_ac       uint8
}

func (_cde *Encoder) EncodeInteger(proc Class, value int) (_bad error) {
	_bc.Log.Trace("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027", value, proc)
	if _bad = _cde.encodeInteger(proc, value); _bad != nil {
		return _fe.Wrap(_bad, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072", "")
	}
	return nil
}

func (_fgb *Encoder) Reset() {
	_fgb._gd = 0x8000
	_fgb._bg = 0
	_fgb._ed = 12
	_fgb._cb = -1
	_fgb._eb = 0
	_fgb._fdd = nil
	_fgb._cd = _ec(_gae)
}

func (_dge *Encoder) rBlock() {
	if _dge._cb >= 0 {
		_dge.emit()
	}
	_dge._cb++
	_dge._eb = uint8(_dge._bg >> 20)
	_dge._bg &= 0xfffff
	_dge._ed = 7
}

func (_agd *Encoder) EncodeBitmap(bm *_ba.Bitmap, duplicateLineRemoval bool) error {
	_bc.Log.Trace("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s", bm.Width, bm.Height, bm)
	var (
		_gcd, _abb      uint8
		_gcc, _dd, _bf  uint16
		_dc, _fea, _gg  byte
		_da, _ecd, _bfg int
		_ddd, _fg       []byte
	)
	for _fed := 0; _fed < bm.Height; _fed++ {
		_dc, _fea = 0, 0
		if _fed >= 2 {
			_dc = bm.Data[(_fed-2)*bm.RowStride]
		}
		if _fed >= 1 {
			_fea = bm.Data[(_fed-1)*bm.RowStride]
			if duplicateLineRemoval {
				_ecd = _fed * bm.RowStride
				_ddd = bm.Data[_ecd : _ecd+bm.RowStride]
				_bfg = (_fed - 1) * bm.RowStride
				_fg = bm.Data[_bfg : _bfg+bm.RowStride]
				if _f.Equal(_ddd, _fg) {
					_abb = _gcd ^ 1
					_gcd = 1
				} else {
					_abb = _gcd
					_gcd = 0
				}
			}
		}
		if duplicateLineRemoval {
			if _fda := _agd.encodeBit(_agd._cd, _ebb, _abb); _fda != nil {
				return _fda
			}
			if _gcd != 0 {
				continue
			}
		}
		_gg = bm.Data[_fed*bm.RowStride]
		_gcc = uint16(_dc >> 5)
		_dd = uint16(_fea >> 4)
		_dc <<= 3
		_fea <<= 4
		_bf = 0
		for _da = 0; _da < bm.Width; _da++ {
			_eg := uint32(_gcc<<11 | _dd<<4 | _bf)
			_acg := (_gg & 0x80) >> 7
			_ee := _agd.encodeBit(_agd._cd, _eg, _acg)
			if _ee != nil {
				return _ee
			}
			_gcc <<= 1
			_dd <<= 1
			_bf <<= 1
			_gcc |= uint16((_dc & 0x80) >> 7)
			_dd |= uint16((_fea & 0x80) >> 7)
			_bf |= uint16(_acg)
			_gcdd := _da % 8
			_bcb := _da/8 + 1
			if _gcdd == 4 && _fed >= 2 {
				_dc = 0
				if _bcb < bm.RowStride {
					_dc = bm.Data[(_fed-2)*bm.RowStride+_bcb]
				}
			} else {
				_dc <<= 1
			}
			if _gcdd == 3 && _fed >= 1 {
				_fea = 0
				if _bcb < bm.RowStride {
					_fea = bm.Data[(_fed-1)*bm.RowStride+_bcb]
				}
			} else {
				_fea <<= 1
			}
			if _gcdd == 7 {
				_gg = 0
				if _bcb < bm.RowStride {
					_gg = bm.Data[_fed*bm.RowStride+_bcb]
				}
			} else {
				_gg <<= 1
			}
			_gcc &= 31
			_dd &= 127
			_bf &= 15
		}
	}
	return nil
}

func (_cgb *Encoder) codeLPS(_cgbg *codingContext, _gce uint32, _fddg uint16, _acgc byte) {
	_cgb._gd -= _fddg
	if _cgb._gd < _fddg {
		_cgb._bg += uint32(_fddg)
	} else {
		_cgb._gd = _fddg
	}
	if _ffdd[_acgc]._ead == 1 {
		_cgbg.flipMps(_gce)
	}
	_cgbg._d[_gce] = _ffdd[_acgc]._acc
	_cgb.renormalize()
}

func (_fga *Encoder) Refine(iTemp, iTarget *_ba.Bitmap, ox, oy int) error {
	for _gf := 0; _gf < iTarget.Height; _gf++ {
		var _bge int
		_bgb := _gf + oy
		var (
			_gfe, _fbb, _aeb, _cbg, _df uint16
			_dg, _age, _dbf, _abbc, _gb byte
		)
		if _bgb >= 1 && (_bgb-1) < iTemp.Height {
			_dg = iTemp.Data[(_bgb-1)*iTemp.RowStride]
		}
		if _bgb >= 0 && _bgb < iTemp.Height {
			_age = iTemp.Data[_bgb*iTemp.RowStride]
		}
		if _bgb >= -1 && _bgb+1 < iTemp.Height {
			_dbf = iTemp.Data[(_bgb+1)*iTemp.RowStride]
		}
		if _gf >= 1 {
			_abbc = iTarget.Data[(_gf-1)*iTarget.RowStride]
		}
		_gb = iTarget.Data[_gf*iTarget.RowStride]
		_cgc := uint(6 + ox)
		_gfe = uint16(_dg >> _cgc)
		_fbb = uint16(_age >> _cgc)
		_aeb = uint16(_dbf >> _cgc)
		_cbg = uint16(_abbc >> 6)
		_ccf := uint(2 - ox)
		_dg <<= _ccf
		_age <<= _ccf
		_dbf <<= _ccf
		_abbc <<= 2
		for _bge = 0; _bge < iTarget.Width; _bge++ {
			_dcd := (_gfe << 10) | (_fbb << 7) | (_aeb << 4) | (_cbg << 1) | _df
			_gba := _gb >> 7
			_ad := _fga.encodeBit(_fga._cd, uint32(_dcd), _gba)
			if _ad != nil {
				return _ad
			}
			_gfe <<= 1
			_fbb <<= 1
			_aeb <<= 1
			_cbg <<= 1
			_gfe |= uint16(_dg >> 7)
			_fbb |= uint16(_age >> 7)
			_aeb |= uint16(_dbf >> 7)
			_cbg |= uint16(_abbc >> 7)
			_df = uint16(_gba)
			_edf := _bge % 8
			_fgd := _bge/8 + 1
			if _edf == 5+ox {
				_dg, _age, _dbf = 0, 0, 0
				if _fgd < iTemp.RowStride && _bgb >= 1 && (_bgb-1) < iTemp.Height {
					_dg = iTemp.Data[(_bgb-1)*iTemp.RowStride+_fgd]
				}
				if _fgd < iTemp.RowStride && _bgb >= 0 && _bgb < iTemp.Height {
					_age = iTemp.Data[_bgb*iTemp.RowStride+_fgd]
				}
				if _fgd < iTemp.RowStride && _bgb >= -1 && (_bgb+1) < iTemp.Height {
					_dbf = iTemp.Data[(_bgb+1)*iTemp.RowStride+_fgd]
				}
			} else {
				_dg <<= 1
				_age <<= 1
				_dbf <<= 1
			}
			if _edf == 5 && _gf >= 1 {
				_abbc = 0
				if _fgd < iTarget.RowStride {
					_abbc = iTarget.Data[(_gf-1)*iTarget.RowStride+_fgd]
				}
			} else {
				_abbc <<= 1
			}
			if _edf == 7 {
				_gb = 0
				if _fgd < iTarget.RowStride {
					_gb = iTarget.Data[_gf*iTarget.RowStride+_fgd]
				}
			} else {
				_gb <<= 1
			}
			_gfe &= 7
			_fbb &= 7
			_aeb &= 7
			_cbg &= 7
		}
	}
	return nil
}

func _ec(_fec int) *codingContext {
	return &codingContext{_d: make([]byte, _fec), _ca: make([]byte, _fec)}
}

func (_gec *Encoder) byteOut() {
	if _gec._eb == 0xff {
		_gec.rBlock()
		return
	}
	if _gec._bg < 0x8000000 {
		_gec.lBlock()
		return
	}
	_gec._eb++
	if _gec._eb != 0xff {
		_gec.lBlock()
		return
	}
	_gec._bg &= 0x7ffffff
	_gec.rBlock()
}

var _ffdd = []state{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}

func (_dba *Encoder) dataSize() int { return _adae*len(_dba._bd) + _dba._cf }
func (_db *Encoder) DataSize() int  { return _db.dataSize() }
func (_adf *Encoder) encodeIAID(_gdde, _eca int) error {
	if _adf._fdd == nil {
		_adf._fdd = _ec(1 << uint(_gdde))
	}
	_fcf := uint32(1<<uint32(_gdde+1)) - 1
	_eca <<= uint(32 - _gdde)
	_ceg := uint32(1)
	for _fbf := 0; _fbf < _gdde; _fbf++ {
		_bbc := _ceg & _fcf
		_agb := uint8((uint32(_eca) & 0x80000000) >> 31)
		if _bbb := _adf.encodeBit(_adf._fdd, _bbc, _agb); _bbb != nil {
			return _bbb
		}
		_ceg = (_ceg << 1) | uint32(_agb)
		_eca <<= 1
	}
	return nil
}
func (_cg *Encoder) Final() { _cg.flush() }

type Class int

func (_ebf *Encoder) encodeOOB(_ada Class) error {
	_ccg := _ebf._cfa[_ada]
	_edg := _ebf.encodeBit(_ccg, 1, 1)
	if _edg != nil {
		return _edg
	}
	_edg = _ebf.encodeBit(_ccg, 3, 0)
	if _edg != nil {
		return _edg
	}
	_edg = _ebf.encodeBit(_ccg, 6, 0)
	if _edg != nil {
		return _edg
	}
	_edg = _ebf.encodeBit(_ccg, 12, 0)
	if _edg != nil {
		return _edg
	}
	return nil
}

func (_edb *Encoder) code1(_bbf *codingContext, _ea uint32, _aa uint16, _bfgb byte) {
	if _bbf.mps(_ea) == 1 {
		_edb.codeMPS(_bbf, _ea, _aa, _bfgb)
	} else {
		_edb.codeLPS(_bbf, _ea, _aa, _bfgb)
	}
}
func (_gc *codingContext) flipMps(_ae uint32) { _gc._ca[_ae] = 1 - _gc._ca[_ae] }
func (_bb *Encoder) WriteTo(w _fa.Writer) (int64, error) {
	const _cga = "\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f"
	var _fagf int64
	for _dbg, _dcc := range _bb._bd {
		_af, _aebd := w.Write(_dcc)
		if _aebd != nil {
			return 0, _fe.Wrapf(_aebd, _cga, "\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b", _dbg)
		}
		_fagf += int64(_af)
	}
	_bb._fag = _bb._fag[:_bb._cf]
	_bag, _egd := w.Write(_bb._fag)
	if _egd != nil {
		return 0, _fe.Wrap(_egd, _cga, "\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073")
	}
	_fagf += int64(_bag)
	return _fagf, nil
}

func (_cfaf *Encoder) flush() {
	_cfaf.setBits()
	_cfaf._bg <<= _cfaf._ed
	_cfaf.byteOut()
	_cfaf._bg <<= _cfaf._ed
	_cfaf.byteOut()
	_cfaf.emit()
	if _cfaf._eb != 0xff {
		_cfaf._cb++
		_cfaf._eb = 0xff
		_cfaf.emit()
	}
	_cfaf._cb++
	_cfaf._eb = 0xac
	_cfaf._cb++
	_cfaf.emit()
}

func (_e Class) String() string {
	switch _e {
	case IAAI:
		return "\u0049\u0041\u0041\u0049"
	case IADH:
		return "\u0049\u0041\u0044\u0048"
	case IADS:
		return "\u0049\u0041\u0044\u0053"
	case IADT:
		return "\u0049\u0041\u0044\u0054"
	case IADW:
		return "\u0049\u0041\u0044\u0057"
	case IAEX:
		return "\u0049\u0041\u0045\u0058"
	case IAFS:
		return "\u0049\u0041\u0046\u0053"
	case IAIT:
		return "\u0049\u0041\u0049\u0054"
	case IARDH:
		return "\u0049\u0041\u0052D\u0048"
	case IARDW:
		return "\u0049\u0041\u0052D\u0057"
	case IARDX:
		return "\u0049\u0041\u0052D\u0058"
	case IARDY:
		return "\u0049\u0041\u0052D\u0059"
	case IARI:
		return "\u0049\u0041\u0052\u0049"
	default:
		return "\u0055N\u004b\u004e\u004f\u0057\u004e"
	}
}

type state struct {
	_dgd        uint16
	_dcee, _acc uint8
	_ead        uint8
}

func (_gdd *Encoder) code0(_ddbf *codingContext, _ddc uint32, _ga uint16, _faga byte) {
	if _ddbf.mps(_ddc) == 0 {
		_gdd.codeMPS(_ddbf, _ddc, _ga, _faga)
	} else {
		_gdd.codeLPS(_ddbf, _ddc, _ga, _faga)
	}
}
func (_g *codingContext) mps(_caa uint32) int { return int(_g._ca[_caa]) }

const (
	IAAI Class = iota
	IADH
	IADS
	IADT
	IADW
	IAEX
	IAFS
	IAIT
	IARDH
	IARDW
	IARDX
	IARDY
	IARI
)

func New() *Encoder { _ge := &Encoder{}; _ge.Init(); return _ge }
func (_fbe *Encoder) codeMPS(_eec *codingContext, _fgf uint32, _agdg uint16, _be byte) {
	_fbe._gd -= _agdg
	if _fbe._gd&0x8000 != 0 {
		_fbe._bg += uint32(_agdg)
		return
	}
	if _fbe._gd < _agdg {
		_fbe._gd = _agdg
	} else {
		_fbe._bg += uint32(_agdg)
	}
	_eec._d[_fgf] = _ffdd[_be]._dcee
	_fbe.renormalize()
}

func (_dce *Encoder) encodeBit(_gbe *codingContext, _abd uint32, _bdc uint8) error {
	const _fae = "\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074"
	_dce._fc++
	if _abd >= uint32(len(_gbe._d)) {
		return _fe.Errorf(_fae, "\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027", _abd)
	}
	_gfeb := _gbe._d[_abd]
	_ffb := _gbe.mps(_abd)
	_gbg := _ffdd[_gfeb]._dgd
	_bc.Log.Trace("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064", _dce._fc, _bdc, _gfeb, _ffb, _gbg, _dce._gd, _dce._bg, _dce._ed, _dce._eb, _dce._cb)
	if _bdc == 0 {
		_dce.code0(_gbe, _abd, _gbg, _gfeb)
	} else {
		_dce.code1(_gbe, _abd, _gbg, _gfeb)
	}
	return nil
}

func (_cac *Encoder) emit() {
	if _cac._cf == _adae {
		_cac._bd = append(_cac._bd, _cac._fag)
		_cac._fag = make([]byte, _adae)
		_cac._cf = 0
	}
	_cac._fag[_cac._cf] = _cac._eb
	_cac._cf++
}

const (
	_gae  = 65536
	_adae = 20 * 1024
)

func (_ddb *Encoder) Flush() { _ddb._cf = 0; _ddb._bd = nil; _ddb._cb = -1 }
