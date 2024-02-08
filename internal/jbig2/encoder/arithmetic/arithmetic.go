package arithmetic

import (
	_f "bytes"
	_fa "io"

	_da "bitbucket.org/shenghui0779/gopdf/common"
	_e "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_b "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

type intEncRangeS struct {
	_ba, _db int
	_bb, _eb uint8
	_fc      uint16
	_fb      uint8
}

func (_fcf *Encoder) encodeBit(_dbf *codingContext, _cfee uint32, _aa uint8) error {
	const _gaf = "\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074"
	_fcf._dfa++
	if _cfee >= uint32(len(_dbf._fce)) {
		return _b.Errorf(_gaf, "\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027", _cfee)
	}
	_aef := _dbf._fce[_cfee]
	_cddb := _dbf.mps(_cfee)
	_fbdf := _dee[_aef]._bda
	_da.Log.Trace("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064", _fcf._dfa, _aa, _aef, _cddb, _fbdf, _fcf._dd, _fcf._cb, _fcf._eae, _fcf._fd, _fcf._cc)
	if _aa == 0 {
		_fcf.code0(_dbf, _cfee, _fbdf, _aef)
	} else {
		_fcf.code1(_dbf, _cfee, _fbdf, _aef)
	}
	return nil
}
func _a(_gg int) *codingContext {
	return &codingContext{_fce: make([]byte, _gg), _dac: make([]byte, _gg)}
}

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

func (_eg *Encoder) codeLPS(_bbd *codingContext, _dgd uint32, _bbdf uint16, _cdc byte) {
	_eg._dd -= _bbdf
	if _eg._dd < _bbdf {
		_eg._cb += uint32(_bbdf)
	} else {
		_eg._dd = _bbdf
	}
	if _dee[_cdc]._agca == 1 {
		_bbd.flipMps(_dgd)
	}
	_bbd._fce[_dgd] = _dee[_cdc]._gba
	_eg.renormalize()
}
func (_ebc *Encoder) EncodeBitmap(bm *_e.Bitmap, duplicateLineRemoval bool) error {
	_da.Log.Trace("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s", bm.Width, bm.Height, bm)
	var (
		_gc, _fbd       uint8
		_dcb, _bad, _de uint16
		_ecf, _ga, _af  byte
		_gd, _ad, _ab   int
		_ae, _gf        []byte
	)
	for _cdg := 0; _cdg < bm.Height; _cdg++ {
		_ecf, _ga = 0, 0
		if _cdg >= 2 {
			_ecf = bm.Data[(_cdg-2)*bm.RowStride]
		}
		if _cdg >= 1 {
			_ga = bm.Data[(_cdg-1)*bm.RowStride]
			if duplicateLineRemoval {
				_ad = _cdg * bm.RowStride
				_ae = bm.Data[_ad : _ad+bm.RowStride]
				_ab = (_cdg - 1) * bm.RowStride
				_gf = bm.Data[_ab : _ab+bm.RowStride]
				if _f.Equal(_ae, _gf) {
					_fbd = _gc ^ 1
					_gc = 1
				} else {
					_fbd = _gc
					_gc = 0
				}
			}
		}
		if duplicateLineRemoval {
			if _gcd := _ebc.encodeBit(_ebc._dag, _cg, _fbd); _gcd != nil {
				return _gcd
			}
			if _gc != 0 {
				continue
			}
		}
		_af = bm.Data[_cdg*bm.RowStride]
		_dcb = uint16(_ecf >> 5)
		_bad = uint16(_ga >> 4)
		_ecf <<= 3
		_ga <<= 4
		_de = 0
		for _gd = 0; _gd < bm.Width; _gd++ {
			_gfd := uint32(_dcb<<11 | _bad<<4 | _de)
			_afc := (_af & 0x80) >> 7
			_cdd := _ebc.encodeBit(_ebc._dag, _gfd, _afc)
			if _cdd != nil {
				return _cdd
			}
			_dcb <<= 1
			_bad <<= 1
			_de <<= 1
			_dcb |= uint16((_ecf & 0x80) >> 7)
			_bad |= uint16((_ga & 0x80) >> 7)
			_de |= uint16(_afc)
			_ag := _gd % 8
			_afd := _gd/8 + 1
			if _ag == 4 && _cdg >= 2 {
				_ecf = 0
				if _afd < bm.RowStride {
					_ecf = bm.Data[(_cdg-2)*bm.RowStride+_afd]
				}
			} else {
				_ecf <<= 1
			}
			if _ag == 3 && _cdg >= 1 {
				_ga = 0
				if _afd < bm.RowStride {
					_ga = bm.Data[(_cdg-1)*bm.RowStride+_afd]
				}
			} else {
				_ga <<= 1
			}
			if _ag == 7 {
				_af = 0
				if _afd < bm.RowStride {
					_af = bm.Data[_cdg*bm.RowStride+_afd]
				}
			} else {
				_af <<= 1
			}
			_dcb &= 31
			_bad &= 127
			_de &= 15
		}
	}
	return nil
}
func (_dgf *Encoder) code1(_adc *codingContext, _bgc uint32, _cfg uint16, _adf byte) {
	if _adc.mps(_bgc) == 1 {
		_dgf.codeMPS(_adc, _bgc, _cfg, _adf)
	} else {
		_dgf.codeLPS(_adc, _bgc, _cfg, _adf)
	}
}

const (
	_ffc = 65536
	_cgg = 20 * 1024
)

var _ _fa.WriterTo = &Encoder{}

func (_dg *Encoder) Flush() { _dg._gb = 0; _dg._dda = nil; _dg._cc = -1 }
func (_fgd *Encoder) setBits() {
	_abcc := _fgd._cb + uint32(_fgd._dd)
	_fgd._cb |= 0xffff
	if _fgd._cb >= _abcc {
		_fgd._cb -= 0x8000
	}
}
func (_efa *Encoder) dataSize() int { return _cgg*len(_efa._dda) + _efa._gb }
func (_dba *Encoder) lBlock() {
	if _dba._cc >= 0 {
		_dba.emit()
	}
	_dba._cc++
	_dba._fd = uint8(_dba._cb >> 19)
	_dba._cb &= 0x7ffff
	_dba._eae = 8
}
func (_ebg *Encoder) Init() {
	_ebg._dag = _a(_ffc)
	_ebg._dd = 0x8000
	_ebg._cb = 0
	_ebg._eae = 12
	_ebg._cc = -1
	_ebg._fd = 0
	_ebg._gb = 0
	_ebg._ec = make([]byte, _cgg)
	for _ef := 0; _ef < len(_ebg._be); _ef++ {
		_ebg._be[_ef] = _a(512)
	}
	_ebg._cd = nil
}
func (_cffc *Encoder) encodeOOB(_abgf Class) error {
	_ffd := _cffc._be[_abgf]
	_eegf := _cffc.encodeBit(_ffd, 1, 1)
	if _eegf != nil {
		return _eegf
	}
	_eegf = _cffc.encodeBit(_ffd, 3, 0)
	if _eegf != nil {
		return _eegf
	}
	_eegf = _cffc.encodeBit(_ffd, 6, 0)
	if _eegf != nil {
		return _eegf
	}
	_eegf = _cffc.encodeBit(_ffd, 12, 0)
	if _eegf != nil {
		return _eegf
	}
	return nil
}
func (_bag *Encoder) EncodeOOB(proc Class) (_cf error) {
	_da.Log.Trace("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027", proc)
	if _cf = _bag.encodeOOB(proc); _cf != nil {
		return _b.Wrap(_cf, "\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB", "")
	}
	return nil
}
func (_df *codingContext) flipMps(_ea uint32) { _df._dac[_ea] = 1 - _df._dac[_ea] }

type codingContext struct {
	_fce []byte
	_dac []byte
}

func (_bbf *Encoder) renormalize() {
	for {
		_bbf._dd <<= 1
		_bbf._cb <<= 1
		_bbf._eae--
		if _bbf._eae == 0 {
			_bbf.byteOut()
		}
		if (_bbf._dd & 0x8000) != 0 {
			break
		}
	}
}
func (_ff *Encoder) Final() { _ff.flush() }
func (_cda *Encoder) WriteTo(w _fa.Writer) (int64, error) {
	const _dfd = "\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f"
	var _aga int64
	for _gde, _dae := range _cda._dda {
		_abc, _eeg := w.Write(_dae)
		if _eeg != nil {
			return 0, _b.Wrapf(_eeg, _dfd, "\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b", _gde)
		}
		_aga += int64(_abc)
	}
	_cda._ec = _cda._ec[:_cda._gb]
	_cfe, _ggd := w.Write(_cda._ec)
	if _ggd != nil {
		return 0, _b.Wrap(_ggd, _dfd, "\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073")
	}
	_aga += int64(_cfe)
	return _aga, nil
}
func (_fg *Encoder) EncodeIAID(symbolCodeLength, value int) (_ccf error) {
	_da.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", symbolCodeLength, value)
	if _ccf = _fg.encodeIAID(symbolCodeLength, value); _ccf != nil {
		return _b.Wrap(_ccf, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044", "")
	}
	return nil
}
func (_bga *Encoder) emit() {
	if _bga._gb == _cgg {
		_bga._dda = append(_bga._dda, _bga._ec)
		_bga._ec = make([]byte, _cgg)
		_bga._gb = 0
	}
	_bga._ec[_bga._gb] = _bga._fd
	_bga._gb++
}
func (_efc *Encoder) code0(_eba *codingContext, _dcd uint32, _bcd uint16, _ceg byte) {
	if _eba.mps(_dcd) == 0 {
		_efc.codeMPS(_eba, _dcd, _bcd, _ceg)
	} else {
		_efc.codeLPS(_eba, _dcd, _bcd, _ceg)
	}
}

var _dee = []state{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}

func New() *Encoder { _ge := &Encoder{}; _ge.Init(); return _ge }

type Encoder struct {
	_cb       uint32
	_dd       uint16
	_eae, _fd uint8
	_cc       int
	_dfa      int
	_dda      [][]byte
	_ec       []byte
	_gb       int
	_dag      *codingContext
	_be       [13]*codingContext
	_cd       *codingContext
}

func (_daa *Encoder) encodeIAID(_cag, _fbe int) error {
	if _daa._cd == nil {
		_daa._cd = _a(1 << uint(_cag))
	}
	_ecg := uint32(1<<uint32(_cag+1)) - 1
	_fbe <<= uint(32 - _cag)
	_cbc := uint32(1)
	for _gea := 0; _gea < _cag; _gea++ {
		_bf := _cbc & _ecg
		_aaa := uint8((uint32(_fbe) & 0x80000000) >> 31)
		if _dcdf := _daa.encodeBit(_daa._cd, _bf, _aaa); _dcdf != nil {
			return _dcdf
		}
		_cbc = (_cbc << 1) | uint32(_aaa)
		_fbe <<= 1
	}
	return nil
}
func (_agg *Encoder) Refine(iTemp, iTarget *_e.Bitmap, ox, oy int) error {
	for _bed := 0; _bed < iTarget.Height; _bed++ {
		var _gcda int
		_cff := _bed + oy
		var (
			_ebe, _ee, _aea, _gdb, _dge uint16
			_eca, _dgg, _ce, _ac, _agc  byte
		)
		if _cff >= 1 && (_cff-1) < iTemp.Height {
			_eca = iTemp.Data[(_cff-1)*iTemp.RowStride]
		}
		if _cff >= 0 && _cff < iTemp.Height {
			_dgg = iTemp.Data[_cff*iTemp.RowStride]
		}
		if _cff >= -1 && _cff+1 < iTemp.Height {
			_ce = iTemp.Data[(_cff+1)*iTemp.RowStride]
		}
		if _bed >= 1 {
			_ac = iTarget.Data[(_bed-1)*iTarget.RowStride]
		}
		_agc = iTarget.Data[_bed*iTarget.RowStride]
		_dgc := uint(6 + ox)
		_ebe = uint16(_eca >> _dgc)
		_ee = uint16(_dgg >> _dgc)
		_aea = uint16(_ce >> _dgc)
		_gdb = uint16(_ac >> 6)
		_ddf := uint(2 - ox)
		_eca <<= _ddf
		_dgg <<= _ddf
		_ce <<= _ddf
		_ac <<= 2
		for _gcda = 0; _gcda < iTarget.Width; _gcda++ {
			_eac := (_ebe << 10) | (_ee << 7) | (_aea << 4) | (_gdb << 1) | _dge
			_fbb := _agc >> 7
			_aeaf := _agg.encodeBit(_agg._dag, uint32(_eac), _fbb)
			if _aeaf != nil {
				return _aeaf
			}
			_ebe <<= 1
			_ee <<= 1
			_aea <<= 1
			_gdb <<= 1
			_ebe |= uint16(_eca >> 7)
			_ee |= uint16(_dgg >> 7)
			_aea |= uint16(_ce >> 7)
			_gdb |= uint16(_ac >> 7)
			_dge = uint16(_fbb)
			_ebgd := _gcda % 8
			_ca := _gcda/8 + 1
			if _ebgd == 5+ox {
				_eca, _dgg, _ce = 0, 0, 0
				if _ca < iTemp.RowStride && _cff >= 1 && (_cff-1) < iTemp.Height {
					_eca = iTemp.Data[(_cff-1)*iTemp.RowStride+_ca]
				}
				if _ca < iTemp.RowStride && _cff >= 0 && _cff < iTemp.Height {
					_dgg = iTemp.Data[_cff*iTemp.RowStride+_ca]
				}
				if _ca < iTemp.RowStride && _cff >= -1 && (_cff+1) < iTemp.Height {
					_ce = iTemp.Data[(_cff+1)*iTemp.RowStride+_ca]
				}
			} else {
				_eca <<= 1
				_dgg <<= 1
				_ce <<= 1
			}
			if _ebgd == 5 && _bed >= 1 {
				_ac = 0
				if _ca < iTarget.RowStride {
					_ac = iTarget.Data[(_bed-1)*iTarget.RowStride+_ca]
				}
			} else {
				_ac <<= 1
			}
			if _ebgd == 7 {
				_agc = 0
				if _ca < iTarget.RowStride {
					_agc = iTarget.Data[_bed*iTarget.RowStride+_ca]
				}
			} else {
				_agc <<= 1
			}
			_ebe &= 7
			_ee &= 7
			_aea &= 7
			_gdb &= 7
		}
	}
	return nil
}

type Class int

func (_aaae *Encoder) rBlock() {
	if _aaae._cc >= 0 {
		_aaae.emit()
	}
	_aaae._cc++
	_aaae._fd = uint8(_aaae._cb >> 20)
	_aaae._cb &= 0xfffff
	_aaae._eae = 7
}

type state struct {
	_bda       uint16
	_gdg, _gba uint8
	_agca      uint8
}

func (_cec *Encoder) encodeInteger(_fdb Class, _aefe int) error {
	const _dfe = "E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072"
	if _aefe > 2000000000 || _aefe < -2000000000 {
		return _b.Errorf(_dfe, "\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027", _aefe)
	}
	_fde := _cec._be[_fdb]
	_aec := uint32(1)
	var _fda int
	for ; ; _fda++ {
		if _bg[_fda]._ba <= _aefe && _bg[_fda]._db >= _aefe {
			break
		}
	}
	if _aefe < 0 {
		_aefe = -_aefe
	}
	_aefe -= int(_bg[_fda]._fc)
	_bac := _bg[_fda]._bb
	for _fae := uint8(0); _fae < _bg[_fda]._eb; _fae++ {
		_aca := _bac & 1
		if _eab := _cec.encodeBit(_fde, _aec, _aca); _eab != nil {
			return _b.Wrap(_eab, _dfe, "")
		}
		_bac >>= 1
		if _aec&0x100 > 0 {
			_aec = (((_aec << 1) | uint32(_aca)) & 0x1ff) | 0x100
		} else {
			_aec = (_aec << 1) | uint32(_aca)
		}
	}
	_aefe <<= 32 - _bg[_fda]._fb
	for _abg := uint8(0); _abg < _bg[_fda]._fb; _abg++ {
		_gad := uint8((uint32(_aefe) & 0x80000000) >> 31)
		if _fdba := _cec.encodeBit(_fde, _aec, _gad); _fdba != nil {
			return _b.Wrap(_fdba, _dfe, "\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064")
		}
		_aefe <<= 1
		if _aec&0x100 != 0 {
			_aec = (((_aec << 1) | uint32(_gad)) & 0x1ff) | 0x100
		} else {
			_aec = (_aec << 1) | uint32(_gad)
		}
	}
	return nil
}
func (_dc *Encoder) DataSize() int { return _dc.dataSize() }
func (_cac *Encoder) codeMPS(_bd *codingContext, _cba uint32, _ffa uint16, _fgb byte) {
	_cac._dd -= _ffa
	if _cac._dd&0x8000 != 0 {
		_cac._cb += uint32(_ffa)
		return
	}
	if _cac._dd < _ffa {
		_cac._dd = _ffa
	} else {
		_cac._cb += uint32(_ffa)
	}
	_bd._fce[_cba] = _dee[_fgb]._gdg
	_cac.renormalize()
}
func (_ebda *Encoder) flush() {
	_ebda.setBits()
	_ebda._cb <<= _ebda._eae
	_ebda.byteOut()
	_ebda._cb <<= _ebda._eae
	_ebda.byteOut()
	_ebda.emit()
	if _ebda._fd != 0xff {
		_ebda._cc++
		_ebda._fd = 0xff
		_ebda.emit()
	}
	_ebda._cc++
	_ebda._fd = 0xac
	_ebda._cc++
	_ebda.emit()
}

var _bg = []intEncRangeS{{0, 3, 0, 2, 0, 2}, {-1, -1, 9, 4, 0, 0}, {-3, -2, 5, 3, 2, 1}, {4, 19, 2, 3, 4, 4}, {-19, -4, 3, 3, 4, 4}, {20, 83, 6, 4, 20, 6}, {-83, -20, 7, 4, 20, 6}, {84, 339, 14, 5, 84, 8}, {-339, -84, 15, 5, 84, 8}, {340, 4435, 30, 6, 340, 12}, {-4435, -340, 31, 6, 340, 12}, {4436, 2000000000, 62, 6, 4436, 32}, {-2000000000, -4436, 63, 6, 4436, 32}}

func (_c *codingContext) mps(_g uint32) int { return int(_c._dac[_g]) }
func (_ecab *Encoder) Reset() {
	_ecab._dd = 0x8000
	_ecab._cb = 0
	_ecab._eae = 12
	_ecab._cc = -1
	_ecab._fd = 0
	_ecab._cd = nil
	_ecab._dag = _a(_ffc)
}
func (_ebd *Encoder) EncodeInteger(proc Class, value int) (_faa error) {
	_da.Log.Trace("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027", value, proc)
	if _faa = _ebd.encodeInteger(proc, value); _faa != nil {
		return _b.Wrap(_faa, "\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072", "")
	}
	return nil
}
func (_bc Class) String() string {
	switch _bc {
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

const _cg = 0x9b25

func (_efd *Encoder) byteOut() {
	if _efd._fd == 0xff {
		_efd.rBlock()
		return
	}
	if _efd._cb < 0x8000000 {
		_efd.lBlock()
		return
	}
	_efd._fd++
	if _efd._fd != 0xff {
		_efd.lBlock()
		return
	}
	_efd._cb &= 0x7ffffff
	_efd.rBlock()
}
