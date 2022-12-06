package arithmetic

import (
	_da "bytes"
	_b "io"

	_be "bitbucket.org/shenghui0779/gopdf/common"
	_dd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_dc "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_feb *Encoder) renormalize() {
	for {
		_feb._egb <<= 1
		_feb._c <<= 1
		_feb._bg--
		if _feb._bg == 0 {
			_feb.byteOut()
		}
		if (_feb._egb & 0x8000) != 0 {
			break
		}
	}
}
func (_gfcg *Encoder) encodeOOB(_ead Class) error {
	_eae := _gfcg._gg[_ead]
	_ffc := _gfcg.encodeBit(_eae, 1, 1)
	if _ffc != nil {
		return _ffc
	}
	_ffc = _gfcg.encodeBit(_eae, 3, 0)
	if _ffc != nil {
		return _ffc
	}
	_ffc = _gfcg.encodeBit(_eae, 6, 0)
	if _ffc != nil {
		return _ffc
	}
	_ffc = _gfcg.encodeBit(_eae, 12, 0)
	if _ffc != nil {
		return _ffc
	}
	return nil
}
func (_ddf *Encoder) Final() { _ddf.flush() }

const _aeb = 0x9b25

func (_abb *Encoder) lBlock() {
	if _abb._aea >= 0 {
		_abb.emit()
	}
	_abb._aea++
	_abb._bdb = uint8(_abb._c >> 19)
	_abb._c &= 0x7ffff
	_abb._bg = 8
}
func (_cef *Encoder) setBits() {
	_cga := _cef._c + uint32(_cef._egb)
	_cef._c |= 0xffff
	if _cef._c >= _cga {
		_cef._c -= 0x8000
	}
}
func (_fdf *Encoder) Flush() {
	_fdf._fa = 0
	_fdf._ec = nil
	_fdf._aea = -1
}
func (_ef *Encoder) Init() {
	_ef._bdbd = _fc(_gbg)
	_ef._egb = 0x8000
	_ef._c = 0
	_ef._bg = 12
	_ef._aea = -1
	_ef._bdb = 0
	_ef._fa = 0
	_ef._dcf = make([]byte, _ebea)
	for _efc := 0; _efc < len(_ef._gg); _efc++ {
		_ef._gg[_efc] = _fc(512)
	}
	_ef._dae = nil
}
func (_fcgg *Encoder) flush() {
	_fcgg.setBits()
	_fcgg._c <<= _fcgg._bg
	_fcgg.byteOut()
	_fcgg._c <<= _fcgg._bg
	_fcgg.byteOut()
	_fcgg.emit()
	if _fcgg._bdb != 0xff {
		_fcgg._aea++
		_fcgg._bdb = 0xff
		_fcgg.emit()
	}
	_fcgg._aea++
	_fcgg._bdb = 0xac
	_fcgg._aea++
	_fcgg.emit()
}
func (_fee *Encoder) code1(_aa *codingContext, _gga uint32, _acg uint16, _egbd byte) {
	if _aa.mps(_gga) == 1 {
		_fee.codeMPS(_aa, _gga, _acg, _egbd)
	} else {
		_fee.codeLPS(_aa, _gga, _acg, _egbd)
	}
}
func (_gbf *Encoder) encodeInteger(_ecde Class, _gfc int) error {
	const _aga = "E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072"
	if _gfc > 2000000000 || _gfc < -2000000000 {
		return _dc.Errorf(_aga, "\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", _gfc)
	}
	_fba := _gbf._gg[_ecde]
	_bgdg := uint32(1)
	var _bae int
	for ; ; _bae++ {
		if _ae[_bae]._bb <= _gfc && _ae[_bae]._a >= _gfc {
			break
		}
	}
	if _gfc < 0 {
		_gfc = -_gfc
	}
	_gfc -= int(_ae[_bae]._ad)
	_agg := _ae[_bae]._dce
	for _bcd := uint8(0); _bcd < _ae[_bae]._g; _bcd++ {
		_gfcd := _agg & 1
		if _dcfb := _gbf.encodeBit(_fba, _bgdg, _gfcd); _dcfb != nil {
			return _dc.Wrap(_dcfb, _aga, "")
		}
		_agg >>= 1
		if _bgdg&0x100 > 0 {
			_bgdg = (((_bgdg << 1) | uint32(_gfcd)) & 0x1ff) | 0x100
		} else {
			_bgdg = (_bgdg << 1) | uint32(_gfcd)
		}
	}
	_gfc <<= 32 - _ae[_bae]._dcb
	for _ada := uint8(0); _ada < _ae[_bae]._dcb; _ada++ {
		_ccgb := uint8((uint32(_gfc) & 0x80000000) >> 31)
		if _dcdc := _gbf.encodeBit(_fba, _bgdg, _ccgb); _dcdc != nil {
			return _dc.Wrap(_dcdc, _aga, "\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064")
		}
		_gfc <<= 1
		if _bgdg&0x100 != 0 {
			_bgdg = (((_bgdg << 1) | uint32(_ccgb)) & 0x1ff) | 0x100
		} else {
			_bgdg = (_bgdg << 1) | uint32(_ccgb)
		}
	}
	return nil
}
func (_f *codingContext) mps(_bf uint32) int { return int(_f._bd[_bf]) }
func (_ega *Encoder) byteOut() {
	if _ega._bdb == 0xff {
		_ega.rBlock()
		return
	}
	if _ega._c < 0x8000000 {
		_ega.lBlock()
		return
	}
	_ega._bdb++
	if _ega._bdb != 0xff {
		_ega.lBlock()
		return
	}
	_ega._c &= 0x7ffffff
	_ega.rBlock()
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
func (_cea *Encoder) Reset() {
	_cea._egb = 0x8000
	_cea._c = 0
	_cea._bg = 12
	_cea._aea = -1
	_cea._bdb = 0
	_cea._dae = nil
	_cea._bdbd = _fc(_gbg)
}

var _eaf = []state{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}

type state struct {
	_dff         uint16
	_gecd, _gbfc uint8
	_cca         uint8
}

func (_eg *codingContext) flipMps(_fd uint32) { _eg._bd[_fd] = 1 - _eg._bd[_fd] }
func (_dg *Encoder) EncodeBitmap(bm *_dd.Bitmap, duplicateLineRemoval bool) error {
	_be.Log.Trace("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s", bm.Width, bm.Height, bm)
	var (
		_dee, _ac       uint8
		_bc, _gc, _efg  uint16
		_cf, _fe, _efgf byte
		_gd, _aeae, _cb int
		_gge, _bdd      []byte
	)
	for _gf := 0; _gf < bm.Height; _gf++ {
		_cf, _fe = 0, 0
		if _gf >= 2 {
			_cf = bm.Data[(_gf-2)*bm.RowStride]
		}
		if _gf >= 1 {
			_fe = bm.Data[(_gf-1)*bm.RowStride]
			if duplicateLineRemoval {
				_aeae = _gf * bm.RowStride
				_gge = bm.Data[_aeae : _aeae+bm.RowStride]
				_cb = (_gf - 1) * bm.RowStride
				_bdd = bm.Data[_cb : _cb+bm.RowStride]
				if _da.Equal(_gge, _bdd) {
					_ac = _dee ^ 1
					_dee = 1
				} else {
					_ac = _dee
					_dee = 0
				}
			}
		}
		if duplicateLineRemoval {
			if _ce := _dg.encodeBit(_dg._bdbd, _aeb, _ac); _ce != nil {
				return _ce
			}
			if _dee != 0 {
				continue
			}
		}
		_efgf = bm.Data[_gf*bm.RowStride]
		_bc = uint16(_cf >> 5)
		_gc = uint16(_fe >> 4)
		_cf <<= 3
		_fe <<= 4
		_efg = 0
		for _gd = 0; _gd < bm.Width; _gd++ {
			_eb := uint32(_bc<<11 | _gc<<4 | _efg)
			_bgc := (_efgf & 0x80) >> 7
			_bge := _dg.encodeBit(_dg._bdbd, _eb, _bgc)
			if _bge != nil {
				return _bge
			}
			_bc <<= 1
			_gc <<= 1
			_efg <<= 1
			_bc |= uint16((_cf & 0x80) >> 7)
			_gc |= uint16((_fe & 0x80) >> 7)
			_efg |= uint16(_bgc)
			_ge := _gd % 8
			_eff := _gd/8 + 1
			if _ge == 4 && _gf >= 2 {
				_cf = 0
				if _eff < bm.RowStride {
					_cf = bm.Data[(_gf-2)*bm.RowStride+_eff]
				}
			} else {
				_cf <<= 1
			}
			if _ge == 3 && _gf >= 1 {
				_fe = 0
				if _eff < bm.RowStride {
					_fe = bm.Data[(_gf-1)*bm.RowStride+_eff]
				}
			} else {
				_fe <<= 1
			}
			if _ge == 7 {
				_efgf = 0
				if _eff < bm.RowStride {
					_efgf = bm.Data[_gf*bm.RowStride+_eff]
				}
			} else {
				_efgf <<= 1
			}
			_bc &= 31
			_gc &= 127
			_efg &= 15
		}
	}
	return nil
}
func _fc(_bef int) *codingContext {
	return &codingContext{_adf: make([]byte, _bef), _bd: make([]byte, _bef)}
}

type intEncRangeS struct {
	_bb, _a  int
	_dce, _g uint8
	_ad      uint16
	_dcb     uint8
}

func New() *Encoder {
	_de := &Encoder{}
	_de.Init()
	return _de
}
func (_bfa *Encoder) WriteTo(w _b.Writer) (int64, error) {
	const _cd = "\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f"
	var _fca int64
	for _bddb, _gec := range _bfa._ec {
		_daeg, _ba := w.Write(_gec)
		if _ba != nil {
			return 0, _dc.Wrapf(_ba, _cd, "\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b", _bddb)
		}
		_fca += int64(_daeg)
	}
	_bfa._dcf = _bfa._dcf[:_bfa._fa]
	_edc, _ecd := w.Write(_bfa._dcf)
	if _ecd != nil {
		return 0, _dc.Wrap(_ecd, _cd, "\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073")
	}
	_fca += int64(_edc)
	return _fca, nil
}

type Encoder struct {
	_c        uint32
	_egb      uint16
	_bg, _bdb uint8
	_aea      int
	_df       int
	_ec       [][]byte
	_dcf      []byte
	_fa       int
	_bdbd     *codingContext
	_gg       [13]*codingContext
	_dae      *codingContext
}

func (_cgf *Encoder) codeMPS(_cce *codingContext, _ag uint32, _dfac uint16, _bccd byte) {
	_cgf._egb -= _dfac
	if _cgf._egb&0x8000 != 0 {
		_cgf._c += uint32(_dfac)
		return
	}
	if _cgf._egb < _dfac {
		_cgf._egb = _dfac
	} else {
		_cgf._c += uint32(_dfac)
	}
	_cce._adf[_ag] = _eaf[_bccd]._gecd
	_cgf.renormalize()
}

type codingContext struct {
	_adf []byte
	_bd  []byte
}

var _ _b.WriterTo = &Encoder{}

func (_ddd *Encoder) DataSize() int { return _ddd.dataSize() }

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

var _ae = []intEncRangeS{{0, 3, 0, 2, 0, 2}, {-1, -1, 9, 4, 0, 0}, {-3, -2, 5, 3, 2, 1}, {4, 19, 2, 3, 4, 4}, {-19, -4, 3, 3, 4, 4}, {20, 83, 6, 4, 20, 6}, {-83, -20, 7, 4, 20, 6}, {84, 339, 14, 5, 84, 8}, {-339, -84, 15, 5, 84, 8}, {340, 4435, 30, 6, 340, 12}, {-4435, -340, 31, 6, 340, 12}, {4436, 2000000000, 62, 6, 4436, 32}, {-2000000000, -4436, 63, 6, 4436, 32}}

func (_cdc *Encoder) encodeIAID(_ff, _ggf int) error {
	if _cdc._dae == nil {
		_cdc._dae = _fc(1 << uint(_ff))
	}
	_dga := uint32(1<<uint32(_ff+1)) - 1
	_ggf <<= uint(32 - _ff)
	_ab := uint32(1)
	for _cggg := 0; _cggg < _ff; _cggg++ {
		_ee := _ab & _dga
		_acf := uint8((uint32(_ggf) & 0x80000000) >> 31)
		if _ceb := _cdc.encodeBit(_cdc._dae, _ee, _acf); _ceb != nil {
			return _ceb
		}
		_ab = (_ab << 1) | uint32(_acf)
		_ggf <<= 1
	}
	return nil
}
func (_add *Encoder) Refine(iTemp, iTarget *_dd.Bitmap, ox, oy int) error {
	for _ebe := 0; _ebe < iTarget.Height; _ebe++ {
		var _db int
		_bbe := _ebe + oy
		var (
			_dcd, _af, _dec, _ed, _cc      uint16
			_fcd, _aebe, _cbf, _fea, _efce byte
		)
		if _bbe >= 1 && (_bbe-1) < iTemp.Height {
			_fcd = iTemp.Data[(_bbe-1)*iTemp.RowStride]
		}
		if _bbe >= 0 && _bbe < iTemp.Height {
			_aebe = iTemp.Data[_bbe*iTemp.RowStride]
		}
		if _bbe >= -1 && _bbe+1 < iTemp.Height {
			_cbf = iTemp.Data[(_bbe+1)*iTemp.RowStride]
		}
		if _ebe >= 1 {
			_fea = iTarget.Data[(_ebe-1)*iTarget.RowStride]
		}
		_efce = iTarget.Data[_ebe*iTarget.RowStride]
		_dcea := uint(6 + ox)
		_dcd = uint16(_fcd >> _dcea)
		_af = uint16(_aebe >> _dcea)
		_dec = uint16(_cbf >> _dcea)
		_ed = uint16(_fea >> 6)
		_ccf := uint(2 - ox)
		_fcd <<= _ccf
		_aebe <<= _ccf
		_cbf <<= _ccf
		_fea <<= 2
		for _db = 0; _db < iTarget.Width; _db++ {
			_geb := (_dcd << 10) | (_af << 7) | (_dec << 4) | (_ed << 1) | _cc
			_dbg := _efce >> 7
			_gb := _add.encodeBit(_add._bdbd, uint32(_geb), _dbg)
			if _gb != nil {
				return _gb
			}
			_dcd <<= 1
			_af <<= 1
			_dec <<= 1
			_ed <<= 1
			_dcd |= uint16(_fcd >> 7)
			_af |= uint16(_aebe >> 7)
			_dec |= uint16(_cbf >> 7)
			_ed |= uint16(_fea >> 7)
			_cc = uint16(_dbg)
			_bdc := _db % 8
			_bcc := _db/8 + 1
			if _bdc == 5+ox {
				_fcd, _aebe, _cbf = 0, 0, 0
				if _bcc < iTemp.RowStride && _bbe >= 1 && (_bbe-1) < iTemp.Height {
					_fcd = iTemp.Data[(_bbe-1)*iTemp.RowStride+_bcc]
				}
				if _bcc < iTemp.RowStride && _bbe >= 0 && _bbe < iTemp.Height {
					_aebe = iTemp.Data[_bbe*iTemp.RowStride+_bcc]
				}
				if _bcc < iTemp.RowStride && _bbe >= -1 && (_bbe+1) < iTemp.Height {
					_cbf = iTemp.Data[(_bbe+1)*iTemp.RowStride+_bcc]
				}
			} else {
				_fcd <<= 1
				_aebe <<= 1
				_cbf <<= 1
			}
			if _bdc == 5 && _ebe >= 1 {
				_fea = 0
				if _bcc < iTarget.RowStride {
					_fea = iTarget.Data[(_ebe-1)*iTarget.RowStride+_bcc]
				}
			} else {
				_fea <<= 1
			}
			if _bdc == 7 {
				_efce = 0
				if _bcc < iTarget.RowStride {
					_efce = iTarget.Data[_ebe*iTarget.RowStride+_bcc]
				}
			} else {
				_efce <<= 1
			}
			_dcd &= 7
			_af &= 7
			_dec &= 7
			_ed &= 7
		}
	}
	return nil
}
func (_ggeb *Encoder) encodeBit(_cfc *codingContext, _ccg uint32, _eca uint8) error {
	const _cbd = "\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074"
	_ggeb._df++
	if _ccg >= uint32(len(_cfc._adf)) {
		return _dc.Errorf(_cbd, "\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027", _ccg)
	}
	_cbb := _cfc._adf[_ccg]
	_eab := _cfc.mps(_ccg)
	_ggaa := _eaf[_cbb]._dff
	_be.Log.Trace("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064", _ggeb._df, _eca, _cbb, _eab, _ggaa, _ggeb._egb, _ggeb._c, _ggeb._bg, _ggeb._bdb, _ggeb._aea)
	if _eca == 0 {
		_ggeb.code0(_cfc, _ccg, _ggaa, _cbb)
	} else {
		_ggeb.code1(_cfc, _ccg, _ggaa, _cbb)
	}
	return nil
}
func (_ca *Encoder) codeLPS(_fb *codingContext, _bdg uint32, _fcg uint16, _ea byte) {
	_ca._egb -= _fcg
	if _ca._egb < _fcg {
		_ca._c += uint32(_fcg)
	} else {
		_ca._egb = _fcg
	}
	if _eaf[_ea]._cca == 1 {
		_fb.flipMps(_bdg)
	}
	_fb._adf[_bdg] = _eaf[_ea]._gbfc
	_ca.renormalize()
}
func (_gdd *Encoder) dataSize() int { return _ebea*len(_gdd._ec) + _gdd._fa }
func (_dbe *Encoder) rBlock() {
	if _dbe._aea >= 0 {
		_dbe.emit()
	}
	_dbe._aea++
	_dbe._bdb = uint8(_dbe._c >> 20)
	_dbe._c &= 0xfffff
	_dbe._bg = 7
}

const (
	_gbg  = 65536
	_ebea = 20 * 1024
)

func (_cg *Encoder) EncodeInteger(proc Class, value int) (_fg error) {
	_be.Log.Trace("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027", value, proc)
	if _fg = _cg.encodeInteger(proc, value); _fg != nil {
		return _dc.Wrap(_fg, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072", "")
	}
	return nil
}
func (_gcc *Encoder) EncodeIAID(symbolCodeLength, value int) (_bce error) {
	_be.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", symbolCodeLength, value)
	if _bce = _gcc.encodeIAID(symbolCodeLength, value); _bce != nil {
		return _dc.Wrap(_bce, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044", "")
	}
	return nil
}
func (_bda *Encoder) emit() {
	if _bda._fa == _ebea {
		_bda._ec = append(_bda._ec, _bda._dcf)
		_bda._dcf = make([]byte, _ebea)
		_bda._fa = 0
	}
	_bda._dcf[_bda._fa] = _bda._bdb
	_bda._fa++
}

type Class int

func (_cgg *Encoder) EncodeOOB(proc Class) (_ebf error) {
	_be.Log.Trace("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027", proc)
	if _ebf = _cgg.encodeOOB(proc); _ebf != nil {
		return _dc.Wrap(_ebf, "\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB", "")
	}
	return nil
}
func (_afc *Encoder) code0(_dbf *codingContext, _afe uint32, _ece uint16, _dfa byte) {
	if _dbf.mps(_afe) == 0 {
		_afc.codeMPS(_dbf, _afe, _ece, _dfa)
	} else {
		_afc.codeLPS(_dbf, _afe, _ece, _dfa)
	}
}
