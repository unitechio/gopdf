package arithmetic

import (
	_fcf "fmt"
	_fc "io"
	_e "strings"

	_g "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_ef "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
)

func (_dd *DecoderStats) SetIndex(index int32) { _dd._fec = index }
func (_efg *Decoder) DecodeBit(stats *DecoderStats) (int, error) {
	var (
		_ee int
		_fg = _gf[stats.cx()][0]
		_cg = int32(stats.cx())
	)
	defer func() { _efg._fe++ }()
	_efg._gg -= _fg
	if (_efg._gd >> 16) < uint64(_fg) {
		_ee = _efg.lpsExchange(stats, _cg, _fg)
		if _dg := _efg.renormalize(); _dg != nil {
			return 0, _dg
		}
	} else {
		_efg._gd -= uint64(_fg) << 16
		if (_efg._gg & 0x8000) == 0 {
			_ee = _efg.mpsExchange(stats, _cg)
			if _b := _efg.renormalize(); _b != nil {
				return 0, _b
			}
		} else {
			_ee = int(stats.getMps())
		}
	}
	return _ee, nil
}
func (_bcf *Decoder) readByte() error {
	if _bcf._c.StreamPosition() > _bcf._ga {
		if _, _gdf := _bcf._c.Seek(-1, _fc.SeekCurrent); _gdf != nil {
			return _gdf
		}
	}
	_cgc, _bcfb := _bcf._c.ReadByte()
	if _bcfb != nil {
		return _bcfb
	}
	_bcf._ed = _cgc
	if _bcf._ed == 0xFF {
		_ag, _ggc := _bcf._c.ReadByte()
		if _ggc != nil {
			return _ggc
		}
		if _ag > 0x8F {
			_bcf._gd += 0xFF00
			_bcf._fd = 8
			if _, _fa := _bcf._c.Seek(-2, _fc.SeekCurrent); _fa != nil {
				return _fa
			}
		} else {
			_bcf._gd += uint64(_ag) << 9
			_bcf._fd = 7
		}
	} else {
		_cgc, _bcfb = _bcf._c.ReadByte()
		if _bcfb != nil {
			return _bcfb
		}
		_bcf._ed = _cgc
		_bcf._gd += uint64(_bcf._ed) << 8
		_bcf._fd = 8
	}
	_bcf._gd &= 0xFFFFFFFFFF
	return nil
}
func (_cff *Decoder) DecodeIAID(codeLen uint64, stats *DecoderStats) (int64, error) {
	_cff._gc = 1
	var _dgd uint64
	for _dgd = 0; _dgd < codeLen; _dgd++ {
		stats.SetIndex(int32(_cff._gc))
		_bc, _de := _cff.DecodeBit(stats)
		if _de != nil {
			return 0, _de
		}
		_cff._gc = (_cff._gc << 1) | int64(_bc)
	}
	_ce := _cff._gc - (1 << codeLen)
	return _ce, nil
}
func (_fcg *DecoderStats) setEntry(_bff int) {
	_gdfa := byte(_bff & 0x7f)
	_fcg._gcda[_fcg._fec] = _gdfa
}
func (_ea *Decoder) lpsExchange(_aag *DecoderStats, _af int32, _faa uint32) int {
	_agf := _aag.getMps()
	if _ea._gg < _faa {
		_aag.setEntry(int(_gf[_af][1]))
		_ea._gg = _faa
		return int(_agf)
	}
	if _gf[_af][3] == 1 {
		_aag.toggleMps()
	}
	_aag.setEntry(int(_gf[_af][2]))
	_ea._gg = _faa
	return int(1 - _agf)
}
func (_be *Decoder) init() error {
	_be._ga = _be._c.StreamPosition()
	_bef, _ab := _be._c.ReadByte()
	if _ab != nil {
		_g.Log.Debug("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076", _ab)
		return _ab
	}
	_be._ed = _bef
	_be._gd = uint64(_bef) << 16
	if _ab = _be.readByte(); _ab != nil {
		return _ab
	}
	_be._gd <<= 7
	_be._fd -= 7
	_be._gg = 0x8000
	_be._fe++
	return nil
}

var (
	_gf = [][4]uint32{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}
)

func (_bccg *DecoderStats) getMps() byte { return _bccg._bcc[_bccg._fec] }
func (_eb *Decoder) DecodeInt(stats *DecoderStats) (int32, error) {
	var (
		_gfc, _edg      int32
		_bb, _bbd, _fef int
		_ge             error
	)
	if stats == nil {
		stats = NewStats(512, 1)
	}
	_eb._gc = 1
	_bbd, _ge = _eb.decodeIntBit(stats)
	if _ge != nil {
		return 0, _ge
	}
	_bb, _ge = _eb.decodeIntBit(stats)
	if _ge != nil {
		return 0, _ge
	}
	if _bb == 1 {
		_bb, _ge = _eb.decodeIntBit(stats)
		if _ge != nil {
			return 0, _ge
		}
		if _bb == 1 {
			_bb, _ge = _eb.decodeIntBit(stats)
			if _ge != nil {
				return 0, _ge
			}
			if _bb == 1 {
				_bb, _ge = _eb.decodeIntBit(stats)
				if _ge != nil {
					return 0, _ge
				}
				if _bb == 1 {
					_bb, _ge = _eb.decodeIntBit(stats)
					if _ge != nil {
						return 0, _ge
					}
					if _bb == 1 {
						_fef = 32
						_edg = 4436
					} else {
						_fef = 12
						_edg = 340
					}
				} else {
					_fef = 8
					_edg = 84
				}
			} else {
				_fef = 6
				_edg = 20
			}
		} else {
			_fef = 4
			_edg = 4
		}
	} else {
		_fef = 2
		_edg = 0
	}
	for _cf := 0; _cf < _fef; _cf++ {
		_bb, _ge = _eb.decodeIntBit(stats)
		if _ge != nil {
			return 0, _ge
		}
		_gfc = (_gfc << 1) | int32(_bb)
	}
	_gfc += _edg
	if _bbd == 0 {
		return _gfc, nil
	} else if _bbd == 1 && _gfc > 0 {
		return -_gfc, nil
	}
	return 0, _ef.ErrOOB
}
func (_gge *Decoder) mpsExchange(_dc *DecoderStats, _ad int32) int {
	_gcd := _dc._bcc[_dc._fec]
	if _gge._gg < _gf[_ad][0] {
		if _gf[_ad][3] == 1 {
			_dc.toggleMps()
		}
		_dc.setEntry(int(_gf[_ad][2]))
		return int(1 - _gcd)
	}
	_dc.setEntry(int(_gf[_ad][1]))
	return int(_gcd)
}
func (_cd *DecoderStats) cx() byte { return _cd._gcda[_cd._fec] }
func (_gea *DecoderStats) String() string {
	_cbf := &_e.Builder{}
	_cbf.WriteString(_fcf.Sprintf("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a", len(_gea._gcda)))
	for _afd, _fdf := range _gea._gcda {
		if _fdf != 0 {
			_cbf.WriteString(_fcf.Sprintf("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a", _afd, _fdf))
		}
	}
	return _cbf.String()
}
func NewStats(contextSize int32, index int32) *DecoderStats {
	return &DecoderStats{_fec: index, _df: contextSize, _gcda: make([]byte, contextSize), _bcc: make([]byte, contextSize)}
}
func (_abf *Decoder) decodeIntBit(_aa *DecoderStats) (int, error) {
	_aa.SetIndex(int32(_abf._gc))
	_gff, _bf := _abf.DecodeBit(_aa)
	if _bf != nil {
		_g.Log.Debug("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _bf)
		return _gff, _bf
	}
	if _abf._gc < 256 {
		_abf._gc = ((_abf._gc << uint64(1)) | int64(_gff)) & 0x1ff
	} else {
		_abf._gc = (((_abf._gc<<uint64(1) | int64(_gff)) & 511) | 256) & 0x1ff
	}
	return _gff, nil
}
func (_gadb *DecoderStats) toggleMps() { _gadb._bcc[_gadb._fec] ^= 1 }

type DecoderStats struct {
	_fec  int32
	_df   int32
	_gcda []byte
	_bcc  []byte
}

func (_bccc *DecoderStats) Reset() {
	for _ca := 0; _ca < len(_bccc._gcda); _ca++ {
		_bccc._gcda[_ca] = 0
		_bccc._bcc[_ca] = 0
	}
}

type Decoder struct {
	ContextSize          []uint32
	ReferedToContextSize []uint32
	_c                   _d.StreamReader
	_ed                  uint8
	_gd                  uint64
	_gg                  uint32
	_gc                  int64
	_fd                  int32
	_fe                  int32
	_ga                  int64
}

func (_abd *DecoderStats) Overwrite(dNew *DecoderStats) {
	for _bcca := 0; _bcca < len(_abd._gcda); _bcca++ {
		_abd._gcda[_bcca] = dNew._gcda[_bcca]
		_abd._bcc[_bcca] = dNew._bcc[_bcca]
	}
}
func New(r _d.StreamReader) (*Decoder, error) {
	_efb := &Decoder{_c: r, ContextSize: []uint32{16, 13, 10, 10}, ReferedToContextSize: []uint32{13, 10}}
	if _a := _efb.init(); _a != nil {
		return nil, _a
	}
	return _efb, nil
}
func (_cbg *DecoderStats) Copy() *DecoderStats {
	_dfe := &DecoderStats{_df: _cbg._df, _gcda: make([]byte, _cbg._df)}
	for _eef := 0; _eef < len(_cbg._gcda); _eef++ {
		_dfe._gcda[_eef] = _cbg._gcda[_eef]
	}
	return _dfe
}
func (_gfa *Decoder) renormalize() error {
	for {
		if _gfa._fd == 0 {
			if _cb := _gfa.readByte(); _cb != nil {
				return _cb
			}
		}
		_gfa._gg <<= 1
		_gfa._gd <<= 1
		_gfa._fd--
		if (_gfa._gg & 0x8000) != 0 {
			break
		}
	}
	_gfa._gd &= 0xffffffff
	return nil
}
