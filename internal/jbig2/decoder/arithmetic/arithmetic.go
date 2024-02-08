package arithmetic

import (
	_fe "fmt"
	_f "io"
	_cg "strings"

	_e "bitbucket.org/shenghui0779/gopdf/common"
	_a "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_ab "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
)

func (_dd *Decoder) renormalize() error {
	for {
		if _dd._abb == 0 {
			if _fb := _dd.readByte(); _fb != nil {
				return _fb
			}
		}
		_dd._ed <<= 1
		_dd._cf <<= 1
		_dd._abb--
		if (_dd._ed & 0x8000) != 0 {
			break
		}
	}
	_dd._cf &= 0xffffffff
	return nil
}
func NewStats(contextSize int32, index int32) *DecoderStats {
	return &DecoderStats{_aba: index, _gd: contextSize, _ea: make([]byte, contextSize), _acb: make([]byte, contextSize)}
}
func (_aab *Decoder) decodeIntBit(_dc *DecoderStats) (int, error) {
	_dc.SetIndex(int32(_aab._d))
	_egd, _dcg := _aab.DecodeBit(_dc)
	if _dcg != nil {
		_e.Log.Debug("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _dcg)
		return _egd, _dcg
	}
	if _aab._d < 256 {
		_aab._d = ((_aab._d << uint64(1)) | int64(_egd)) & 0x1ff
	} else {
		_aab._d = (((_aab._d<<uint64(1) | int64(_egd)) & 511) | 256) & 0x1ff
	}
	return _egd, nil
}
func (_aed *Decoder) mpsExchange(_dec *DecoderStats, _fcd int32) int {
	_be := _dec._acb[_dec._aba]
	if _aed._ed < _ff[_fcd][0] {
		if _ff[_fcd][3] == 1 {
			_dec.toggleMps()
		}
		_dec.setEntry(int(_ff[_fcd][2]))
		return int(1 - _be)
	}
	_dec.setEntry(int(_ff[_fcd][1]))
	return int(_be)
}
func (_aea *DecoderStats) getMps() byte         { return _aea._acb[_aea._aba] }
func (_ebd *DecoderStats) SetIndex(index int32) { _ebd._aba = index }
func (_gc *Decoder) init() error {
	_gc._db = _gc._ffa.AbsolutePosition()
	_cae, _ec := _gc._ffa.ReadByte()
	if _ec != nil {
		_e.Log.Debug("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076", _ec)
		return _ec
	}
	_gc._ef = _cae
	_gc._cf = uint64(_cae) << 16
	if _ec = _gc.readByte(); _ec != nil {
		return _ec
	}
	_gc._cf <<= 7
	_gc._abb -= 7
	_gc._ed = 0x8000
	_gc._g++
	return nil
}
func (_gf *DecoderStats) Reset() {
	for _ge := 0; _ge < len(_gf._ea); _ge++ {
		_gf._ea[_ge] = 0
		_gf._acb[_ge] = 0
	}
}
func New(r *_a.Reader) (*Decoder, error) {
	_eb := &Decoder{_ffa: r, ContextSize: []uint32{16, 13, 10, 10}, ReferedToContextSize: []uint32{13, 10}}
	if _ag := _eb.init(); _ag != nil {
		return nil, _ag
	}
	return _eb, nil
}
func (_eaa *DecoderStats) Overwrite(dNew *DecoderStats) {
	for _gg := 0; _gg < len(_eaa._ea); _gg++ {
		_eaa._ea[_gg] = dNew._ea[_gg]
		_eaa._acb[_gg] = dNew._acb[_gg]
	}
}

type Decoder struct {
	ContextSize          []uint32
	ReferedToContextSize []uint32
	_ffa                 *_a.Reader
	_ef                  uint8
	_cf                  uint64
	_ed                  uint32
	_d                   int64
	_abb                 int32
	_g                   int32
	_db                  int64
}

func (_bda *DecoderStats) toggleMps() { _bda._acb[_bda._aba] ^= 1 }
func (_gga *DecoderStats) String() string {
	_ebdf := &_cg.Builder{}
	_ebdf.WriteString(_fe.Sprintf("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a", len(_gga._ea)))
	for _acd, _fbd := range _gga._ea {
		if _fbd != 0 {
			_ebdf.WriteString(_fe.Sprintf("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a", _acd, _fbd))
		}
	}
	return _ebdf.String()
}
func (_ggf *DecoderStats) cx() byte { return _ggf._ea[_ggf._aba] }
func (_dbd *Decoder) readByte() error {
	if _dbd._ffa.AbsolutePosition() > _dbd._db {
		if _, _de := _dbd._ffa.Seek(-1, _f.SeekCurrent); _de != nil {
			return _de
		}
	}
	_fc, _ac := _dbd._ffa.ReadByte()
	if _ac != nil {
		return _ac
	}
	_dbd._ef = _fc
	if _dbd._ef == 0xFF {
		_aad, _dbb := _dbd._ffa.ReadByte()
		if _dbb != nil {
			return _dbb
		}
		if _aad > 0x8F {
			_dbd._cf += 0xFF00
			_dbd._abb = 8
			if _, _dg := _dbd._ffa.Seek(-2, _f.SeekCurrent); _dg != nil {
				return _dg
			}
		} else {
			_dbd._cf += uint64(_aad) << 9
			_dbd._abb = 7
		}
	} else {
		_fc, _ac = _dbd._ffa.ReadByte()
		if _ac != nil {
			return _ac
		}
		_dbd._ef = _fc
		_dbd._cf += uint64(_dbd._ef) << 8
		_dbd._abb = 8
	}
	_dbd._cf &= 0xFFFFFFFFFF
	return nil
}

type DecoderStats struct {
	_aba int32
	_gd  int32
	_ea  []byte
	_acb []byte
}

var (
	_ff = [][4]uint32{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}
)

func (_aegf *Decoder) lpsExchange(_ee *DecoderStats, _cb int32, _fca uint32) int {
	_bd := _ee.getMps()
	if _aegf._ed < _fca {
		_ee.setEntry(int(_ff[_cb][1]))
		_aegf._ed = _fca
		return int(_bd)
	}
	if _ff[_cb][3] == 1 {
		_ee.toggleMps()
	}
	_ee.setEntry(int(_ff[_cb][2]))
	_aegf._ed = _fca
	return int(1 - _bd)
}
func (_bf *Decoder) DecodeInt(stats *DecoderStats) (int32, error) {
	var (
		_aa, _df       int32
		_eda, _ce, _eg int
		_dbf           error
	)
	if stats == nil {
		stats = NewStats(512, 1)
	}
	_bf._d = 1
	_ce, _dbf = _bf.decodeIntBit(stats)
	if _dbf != nil {
		return 0, _dbf
	}
	_eda, _dbf = _bf.decodeIntBit(stats)
	if _dbf != nil {
		return 0, _dbf
	}
	if _eda == 1 {
		_eda, _dbf = _bf.decodeIntBit(stats)
		if _dbf != nil {
			return 0, _dbf
		}
		if _eda == 1 {
			_eda, _dbf = _bf.decodeIntBit(stats)
			if _dbf != nil {
				return 0, _dbf
			}
			if _eda == 1 {
				_eda, _dbf = _bf.decodeIntBit(stats)
				if _dbf != nil {
					return 0, _dbf
				}
				if _eda == 1 {
					_eda, _dbf = _bf.decodeIntBit(stats)
					if _dbf != nil {
						return 0, _dbf
					}
					if _eda == 1 {
						_eg = 32
						_df = 4436
					} else {
						_eg = 12
						_df = 340
					}
				} else {
					_eg = 8
					_df = 84
				}
			} else {
				_eg = 6
				_df = 20
			}
		} else {
			_eg = 4
			_df = 4
		}
	} else {
		_eg = 2
		_df = 0
	}
	for _fa := 0; _fa < _eg; _fa++ {
		_eda, _dbf = _bf.decodeIntBit(stats)
		if _dbf != nil {
			return 0, _dbf
		}
		_aa = (_aa << 1) | int32(_eda)
	}
	_aa += _df
	if _ce == 0 {
		return _aa, nil
	} else if _ce == 1 && _aa > 0 {
		return -_aa, nil
	}
	return 0, _ab.ErrOOB
}
func (_bb *DecoderStats) Copy() *DecoderStats {
	_ebf := &DecoderStats{_gd: _bb._gd, _ea: make([]byte, _bb._gd)}
	copy(_ebf._ea, _bb._ea)
	return _ebf
}
func (_edf *DecoderStats) setEntry(_bbf int) {
	_fgb := byte(_bbf & 0x7f)
	_edf._ea[_edf._aba] = _fgb
}
func (_cab *Decoder) DecodeIAID(codeLen uint64, stats *DecoderStats) (int64, error) {
	_cab._d = 1
	var _dbe uint64
	for _dbe = 0; _dbe < codeLen; _dbe++ {
		stats.SetIndex(int32(_cab._d))
		_aef, _aee := _cab.DecodeBit(stats)
		if _aee != nil {
			return 0, _aee
		}
		_cab._d = (_cab._d << 1) | int64(_aef)
	}
	_fdd := _cab._d - (1 << codeLen)
	return _fdd, nil
}
func (_fd *Decoder) DecodeBit(stats *DecoderStats) (int, error) {
	var (
		_fg int
		_b  = _ff[stats.cx()][0]
		_bg = int32(stats.cx())
	)
	defer func() { _fd._g++ }()
	_fd._ed -= _b
	if (_fd._cf >> 16) < uint64(_b) {
		_fg = _fd.lpsExchange(stats, _bg, _b)
		if _ca := _fd.renormalize(); _ca != nil {
			return 0, _ca
		}
	} else {
		_fd._cf -= uint64(_b) << 16
		if (_fd._ed & 0x8000) == 0 {
			_fg = _fd.mpsExchange(stats, _bg)
			if _ae := _fd.renormalize(); _ae != nil {
				return 0, _ae
			}
		} else {
			_fg = int(stats.getMps())
		}
	}
	return _fg, nil
}
