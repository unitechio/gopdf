package arithmetic

import (
	_c "bytes"
	_f "io"

	_ca "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_de "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_bbfe *Encoder) renormalize() {
	for {
		_bbfe._gb <<= 1
		_bbfe._fb <<= 1
		_bbfe._df--
		if _bbfe._df == 0 {
			_bbfe.byteOut()
		}
		if (_bbfe._gb & 0x8000) != 0 {
			break
		}
	}
}
func (_geg *Encoder) flush() {
	_geg.setBits()
	_geg._fb <<= _geg._df
	_geg.byteOut()
	_geg._fb <<= _geg._df
	_geg.byteOut()
	_geg.emit()
	if _geg._af != 0xff {
		_geg._fef++
		_geg._af = 0xff
		_geg.emit()
	}
	_geg._fef++
	_geg._af = 0xac
	_geg._fef++
	_geg.emit()
}

var _ _f.WriterTo = &Encoder{}

func _ag(_fce int) *codingContext {
	return &codingContext{_fe: make([]byte, _fce), _gc: make([]byte, _fce)}
}
func (_egce *Encoder) byteOut() {
	if _egce._af == 0xff {
		_egce.rBlock()
		return
	}
	if _egce._fb < 0x8000000 {
		_egce.lBlock()
		return
	}
	_egce._af++
	if _egce._af != 0xff {
		_egce.lBlock()
		return
	}
	_egce._fb &= 0x7ffffff
	_egce.rBlock()
}
func (_bff *Encoder) Refine(iTemp, iTarget *_d.Bitmap, ox, oy int) error {
	for _eg := 0; _eg < iTarget.Height; _eg++ {
		var _deb int
		_febb := _eg + oy
		var (
			_ec, _bbf, _bd, _gbe, _cgc uint16
			_fa, _gec, _ce, _acb, _gga byte
		)
		if _febb >= 1 && (_febb-1) < iTemp.Height {
			_fa = iTemp.Data[(_febb-1)*iTemp.RowStride]
		}
		if _febb >= 0 && _febb < iTemp.Height {
			_gec = iTemp.Data[_febb*iTemp.RowStride]
		}
		if _febb >= -1 && _febb+1 < iTemp.Height {
			_ce = iTemp.Data[(_febb+1)*iTemp.RowStride]
		}
		if _eg >= 1 {
			_acb = iTarget.Data[(_eg-1)*iTarget.RowStride]
		}
		_gga = iTarget.Data[_eg*iTarget.RowStride]
		_egc := uint(6 + ox)
		_ec = uint16(_fa >> _egc)
		_bbf = uint16(_gec >> _egc)
		_bd = uint16(_ce >> _egc)
		_gbe = uint16(_acb >> 6)
		_ebf := uint(2 - ox)
		_fa <<= _ebf
		_gec <<= _ebf
		_ce <<= _ebf
		_acb <<= 2
		for _deb = 0; _deb < iTarget.Width; _deb++ {
			_dec := (_ec << 10) | (_bbf << 7) | (_bd << 4) | (_gbe << 1) | _cgc
			_ced := _gga >> 7
			_cb := _bff.encodeBit(_bff._bcb, uint32(_dec), _ced)
			if _cb != nil {
				return _cb
			}
			_ec <<= 1
			_bbf <<= 1
			_bd <<= 1
			_gbe <<= 1
			_ec |= uint16(_fa >> 7)
			_bbf |= uint16(_gec >> 7)
			_bd |= uint16(_ce >> 7)
			_gbe |= uint16(_acb >> 7)
			_cgc = uint16(_ced)
			_dca := _deb % 8
			_aeg := _deb/8 + 1
			if _dca == 5+ox {
				_fa, _gec, _ce = 0, 0, 0
				if _aeg < iTemp.RowStride && _febb >= 1 && (_febb-1) < iTemp.Height {
					_fa = iTemp.Data[(_febb-1)*iTemp.RowStride+_aeg]
				}
				if _aeg < iTemp.RowStride && _febb >= 0 && _febb < iTemp.Height {
					_gec = iTemp.Data[_febb*iTemp.RowStride+_aeg]
				}
				if _aeg < iTemp.RowStride && _febb >= -1 && (_febb+1) < iTemp.Height {
					_ce = iTemp.Data[(_febb+1)*iTemp.RowStride+_aeg]
				}
			} else {
				_fa <<= 1
				_gec <<= 1
				_ce <<= 1
			}
			if _dca == 5 && _eg >= 1 {
				_acb = 0
				if _aeg < iTarget.RowStride {
					_acb = iTarget.Data[(_eg-1)*iTarget.RowStride+_aeg]
				}
			} else {
				_acb <<= 1
			}
			if _dca == 7 {
				_gga = 0
				if _aeg < iTarget.RowStride {
					_gga = iTarget.Data[_eg*iTarget.RowStride+_aeg]
				}
			} else {
				_gga <<= 1
			}
			_ec &= 7
			_bbf &= 7
			_bd &= 7
			_gbe &= 7
		}
	}
	return nil
}

var _ffa = []state{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}

func (_ae *Encoder) DataSize() int { return _ae.dataSize() }
func (_eac *Encoder) Reset() {
	_eac._gb = 0x8000
	_eac._fb = 0
	_eac._df = 12
	_eac._fef = -1
	_eac._af = 0
	_eac._gbb = nil
	_eac._bcb = _ag(_aff)
}

type Encoder struct {
	_fb      uint32
	_gb      uint16
	_df, _af uint8
	_fef     int
	_ac      int
	_cdb     [][]byte
	_bc      []byte
	_ef      int
	_bcb     *codingContext
	_ee      [13]*codingContext
	_gbb     *codingContext
}

func (_dcb *Encoder) Flush() { _dcb._ef = 0; _dcb._cdb = nil; _dcb._fef = -1 }
func (_ddb *Encoder) encodeOOB(_dfa Class) error {
	_bab := _ddb._ee[_dfa]
	_dbd := _ddb.encodeBit(_bab, 1, 1)
	if _dbd != nil {
		return _dbd
	}
	_dbd = _ddb.encodeBit(_bab, 3, 0)
	if _dbd != nil {
		return _dbd
	}
	_dbd = _ddb.encodeBit(_bab, 6, 0)
	if _dbd != nil {
		return _dbd
	}
	_dbd = _ddb.encodeBit(_bab, 12, 0)
	if _dbd != nil {
		return _dbd
	}
	return nil
}
func (_ed *Encoder) Final() { _ed.flush() }
func (_dd *Encoder) code0(_agb *codingContext, _bea uint32, _fg uint16, _dfe byte) {
	if _agb.mps(_bea) == 0 {
		_dd.codeMPS(_agb, _bea, _fg, _dfe)
	} else {
		_dd.codeLPS(_agb, _bea, _fg, _dfe)
	}
}
func (_aec *Encoder) code1(_ebfe *codingContext, _fbe uint32, _cbg uint16, _fca byte) {
	if _ebfe.mps(_fbe) == 1 {
		_aec.codeMPS(_ebfe, _fbe, _cbg, _fca)
	} else {
		_aec.codeLPS(_ebfe, _fbe, _cbg, _fca)
	}
}
func (_dbe *Encoder) WriteTo(w _f.Writer) (int64, error) {
	const _cae = "\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f"
	var _gffa int64
	for _aae, _dcd := range _dbe._cdb {
		_afc, _fab := w.Write(_dcd)
		if _fab != nil {
			return 0, _de.Wrapf(_fab, _cae, "\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b", _aae)
		}
		_gffa += int64(_afc)
	}
	_dbe._bc = _dbe._bc[:_dbe._ef]
	_afcd, _fcf := w.Write(_dbe._bc)
	if _fcf != nil {
		return 0, _de.Wrap(_fcf, _cae, "\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073")
	}
	_gffa += int64(_afcd)
	return _gffa, nil
}
func (_dgd *Encoder) EncodeIAID(symbolCodeLength, value int) (_ad error) {
	_ca.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", symbolCodeLength, value)
	if _ad = _dgd.encodeIAID(symbolCodeLength, value); _ad != nil {
		return _de.Wrap(_ad, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044", "")
	}
	return nil
}
func (_cd *codingContext) flipMps(_db uint32) { _cd._gc[_db] = 1 - _cd._gc[_db] }
func (_dbc *Encoder) Init() {
	_dbc._bcb = _ag(_aff)
	_dbc._gb = 0x8000
	_dbc._fb = 0
	_dbc._df = 12
	_dbc._fef = -1
	_dbc._af = 0
	_dbc._ef = 0
	_dbc._bc = make([]byte, _fbg)
	for _fec := 0; _fec < len(_dbc._ee); _fec++ {
		_dbc._ee[_fec] = _ag(512)
	}
	_dbc._gbb = nil
}
func (_bec *Encoder) rBlock() {
	if _bec._fef >= 0 {
		_bec.emit()
	}
	_bec._fef++
	_bec._af = uint8(_bec._fb >> 20)
	_bec._fb &= 0xfffff
	_bec._df = 7
}
func (_eef *Encoder) encodeIAID(_add, _abbc int) error {
	if _eef._gbb == nil {
		_eef._gbb = _ag(1 << uint(_add))
	}
	_afbf := uint32(1<<uint32(_add+1)) - 1
	_abbc <<= uint(32 - _add)
	_ggf := uint32(1)
	for _gbc := 0; _gbc < _add; _gbc++ {
		_fac := _ggf & _afbf
		_dea := uint8((uint32(_abbc) & 0x80000000) >> 31)
		if _fgb := _eef.encodeBit(_eef._gbb, _fac, _dea); _fgb != nil {
			return _fgb
		}
		_ggf = (_ggf << 1) | uint32(_dea)
		_abbc <<= 1
	}
	return nil
}
func (_dc Class) String() string {
	switch _dc {
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
func (_agf *Encoder) EncodeOOB(proc Class) (_bcf error) {
	_ca.Log.Trace("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027", proc)
	if _bcf = _agf.encodeOOB(proc); _bcf != nil {
		return _de.Wrap(_bcf, "\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB", "")
	}
	return nil
}
func (_ga *Encoder) dataSize() int { return _fbg*len(_ga._cdb) + _ga._ef }
func (_gef *Encoder) encodeInteger(_gac Class, _gfc int) error {
	const _cde = "E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072"
	if _gfc > 2000000000 || _gfc < -2000000000 {
		return _de.Errorf(_cde, "\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", _gfc)
	}
	_gcg := _gef._ee[_gac]
	_gdff := uint32(1)
	var _aef int
	for ; ; _aef++ {
		if _gfg[_aef]._a <= _gfc && _gfg[_aef]._cc >= _gfc {
			break
		}
	}
	if _gfc < 0 {
		_gfc = -_gfc
	}
	_gfc -= int(_gfg[_aef]._be)
	_edf := _gfg[_aef]._e
	for _gfb := uint8(0); _gfb < _gfg[_aef]._g; _gfb++ {
		_cabf := _edf & 1
		if _beac := _gef.encodeBit(_gcg, _gdff, _cabf); _beac != nil {
			return _de.Wrap(_beac, _cde, "")
		}
		_edf >>= 1
		if _gdff&0x100 > 0 {
			_gdff = (((_gdff << 1) | uint32(_cabf)) & 0x1ff) | 0x100
		} else {
			_gdff = (_gdff << 1) | uint32(_cabf)
		}
	}
	_gfc <<= 32 - _gfg[_aef]._gf
	for _eae := uint8(0); _eae < _gfg[_aef]._gf; _eae++ {
		_ff := uint8((uint32(_gfc) & 0x80000000) >> 31)
		if _acg := _gef.encodeBit(_gcg, _gdff, _ff); _acg != nil {
			return _de.Wrap(_acg, _cde, "\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064")
		}
		_gfc <<= 1
		if _gdff&0x100 != 0 {
			_gdff = (((_gdff << 1) | uint32(_ff)) & 0x1ff) | 0x100
		} else {
			_gdff = (_gdff << 1) | uint32(_ff)
		}
	}
	return nil
}

type Class int

func (_fda *Encoder) lBlock() {
	if _fda._fef >= 0 {
		_fda.emit()
	}
	_fda._fef++
	_fda._af = uint8(_fda._fb >> 19)
	_fda._fb &= 0x7ffff
	_fda._df = 8
}
func (_dcg *Encoder) EncodeBitmap(bm *_d.Bitmap, duplicateLineRemoval bool) error {
	_ca.Log.Trace("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s", bm.Width, bm.Height, bm)
	var (
		_ab, _bf        uint8
		_ccg, _ge, _gca uint16
		_eb, _dg, _gfe  byte
		_cdf, _aa, _agg int
		_gg, _agc       []byte
	)
	for _cg := 0; _cg < bm.Height; _cg++ {
		_eb, _dg = 0, 0
		if _cg >= 2 {
			_eb = bm.Data[(_cg-2)*bm.RowStride]
		}
		if _cg >= 1 {
			_dg = bm.Data[(_cg-1)*bm.RowStride]
			if duplicateLineRemoval {
				_aa = _cg * bm.RowStride
				_gg = bm.Data[_aa : _aa+bm.RowStride]
				_agg = (_cg - 1) * bm.RowStride
				_agc = bm.Data[_agg : _agg+bm.RowStride]
				if _c.Equal(_gg, _agc) {
					_bf = _ab ^ 1
					_ab = 1
				} else {
					_bf = _ab
					_ab = 0
				}
			}
		}
		if duplicateLineRemoval {
			if _afb := _dcg.encodeBit(_dcg._bcb, _caf, _bf); _afb != nil {
				return _afb
			}
			if _ab != 0 {
				continue
			}
		}
		_gfe = bm.Data[_cg*bm.RowStride]
		_ccg = uint16(_eb >> 5)
		_ge = uint16(_dg >> 4)
		_eb <<= 3
		_dg <<= 4
		_gca = 0
		for _cdf = 0; _cdf < bm.Width; _cdf++ {
			_gff := uint32(_ccg<<11 | _ge<<4 | _gca)
			_fd := (_gfe & 0x80) >> 7
			_aca := _dcg.encodeBit(_dcg._bcb, _gff, _fd)
			if _aca != nil {
				return _aca
			}
			_ccg <<= 1
			_ge <<= 1
			_gca <<= 1
			_ccg |= uint16((_eb & 0x80) >> 7)
			_ge |= uint16((_dg & 0x80) >> 7)
			_gca |= uint16(_fd)
			_feb := _cdf % 8
			_dff := _cdf/8 + 1
			if _feb == 4 && _cg >= 2 {
				_eb = 0
				if _dff < bm.RowStride {
					_eb = bm.Data[(_cg-2)*bm.RowStride+_dff]
				}
			} else {
				_eb <<= 1
			}
			if _feb == 3 && _cg >= 1 {
				_dg = 0
				if _dff < bm.RowStride {
					_dg = bm.Data[(_cg-1)*bm.RowStride+_dff]
				}
			} else {
				_dg <<= 1
			}
			if _feb == 7 {
				_gfe = 0
				if _dff < bm.RowStride {
					_gfe = bm.Data[_cg*bm.RowStride+_dff]
				}
			} else {
				_gfe <<= 1
			}
			_ccg &= 31
			_ge &= 127
			_gca &= 15
		}
	}
	return nil
}

var _gfg = []intEncRangeS{{0, 3, 0, 2, 0, 2}, {-1, -1, 9, 4, 0, 0}, {-3, -2, 5, 3, 2, 1}, {4, 19, 2, 3, 4, 4}, {-19, -4, 3, 3, 4, 4}, {20, 83, 6, 4, 20, 6}, {-83, -20, 7, 4, 20, 6}, {84, 339, 14, 5, 84, 8}, {-339, -84, 15, 5, 84, 8}, {340, 4435, 30, 6, 340, 12}, {-4435, -340, 31, 6, 340, 12}, {4436, 2000000000, 62, 6, 4436, 32}, {-2000000000, -4436, 63, 6, 4436, 32}}

type state struct {
	_gbg        uint16
	_gbeg, _ded uint8
	_bfc        uint8
}
type intEncRangeS struct {
	_a, _cc int
	_e, _g  uint8
	_be     uint16
	_gf     uint8
}

func New() *Encoder { _bbg := &Encoder{}; _bbg.Init(); return _bbg }
func (_ba *Encoder) EncodeInteger(proc Class, value int) (_ea error) {
	_ca.Log.Trace("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027", value, proc)
	if _ea = _ba.encodeInteger(proc, value); _ea != nil {
		return _de.Wrap(_ea, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072", "")
	}
	return nil
}
func (_cda *Encoder) setBits() {
	_fcg := _cda._fb + uint32(_cda._gb)
	_cda._fb |= 0xffff
	if _cda._fb >= _fcg {
		_cda._fb -= 0x8000
	}
}
func (_dcbf *Encoder) codeMPS(_gfeg *codingContext, _ecb uint32, _ged uint16, _egf byte) {
	_dcbf._gb -= _ged
	if _dcbf._gb&0x8000 != 0 {
		_dcbf._fb += uint32(_ged)
		return
	}
	if _dcbf._gb < _ged {
		_dcbf._gb = _ged
	} else {
		_dcbf._fb += uint32(_ged)
	}
	_gfeg._fe[_ecb] = _ffa[_egf]._gbeg
	_dcbf.renormalize()
}

const (
	_aff = 65536
	_fbg = 20 * 1024
)

func (_abc *Encoder) encodeBit(_cab *codingContext, _bg uint32, _cf uint8) error {
	const _dbg = "\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074"
	_abc._ac++
	if _bg >= uint32(len(_cab._fe)) {
		return _de.Errorf(_dbg, "\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027", _bg)
	}
	_baa := _cab._fe[_bg]
	_fea := _cab.mps(_bg)
	_caa := _ffa[_baa]._gbg
	_ca.Log.Trace("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064", _abc._ac, _cf, _baa, _fea, _caa, _abc._gb, _abc._fb, _abc._df, _abc._af, _abc._fef)
	if _cf == 0 {
		_abc.code0(_cab, _bg, _caa, _baa)
	} else {
		_abc.code1(_cab, _bg, _caa, _baa)
	}
	return nil
}

const _caf = 0x9b25
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

func (_acbb *Encoder) codeLPS(_gbec *codingContext, _gd uint32, _gde uint16, _cbe byte) {
	_acbb._gb -= _gde
	if _acbb._gb < _gde {
		_acbb._fb += uint32(_gde)
	} else {
		_acbb._gb = _gde
	}
	if _ffa[_cbe]._bfc == 1 {
		_gbec.flipMps(_gd)
	}
	_gbec._fe[_gd] = _ffa[_cbe]._ded
	_acbb.renormalize()
}
func (_fc *codingContext) mps(_bb uint32) int { return int(_fc._gc[_bb]) }

type codingContext struct {
	_fe []byte
	_gc []byte
}

func (_gdf *Encoder) emit() {
	if _gdf._ef == _fbg {
		_gdf._cdb = append(_gdf._cdb, _gdf._bc)
		_gdf._bc = make([]byte, _fbg)
		_gdf._ef = 0
	}
	_gdf._bc[_gdf._ef] = _gdf._af
	_gdf._ef++
}
