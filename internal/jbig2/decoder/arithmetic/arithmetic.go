package arithmetic

import (
	_ce "fmt"
	_f "io"
	_g "strings"

	_gc "unitechio/gopdf/gopdf/common"
	_d "unitechio/gopdf/gopdf/internal/bitwise"
	_b "unitechio/gopdf/gopdf/internal/jbig2/internal"
)

func (_bf *Decoder) readByte() error {
	if _bf._dg.AbsolutePosition() > _bf._gd {
		if _, _gab := _bf._dg.Seek(-1, _f.SeekCurrent); _gab != nil {
			return _gab
		}
	}
	_bc, _cg := _bf._dg.ReadByte()
	if _cg != nil {
		return _cg
	}
	_bf._fd = _bc
	if _bf._fd == 0xFF {
		_cgd, _cede := _bf._dg.ReadByte()
		if _cede != nil {
			return _cede
		}
		if _cgd > 0x8F {
			_bf._gg += 0xFF00
			_bf._ef = 8
			if _, _ea := _bf._dg.Seek(-2, _f.SeekCurrent); _ea != nil {
				return _ea
			}
		} else {
			_bf._gg += uint64(_cgd) << 9
			_bf._ef = 7
		}
	} else {
		_bc, _cg = _bf._dg.ReadByte()
		if _cg != nil {
			return _cg
		}
		_bf._fd = _bc
		_bf._gg += uint64(_bf._fd) << 8
		_bf._ef = 8
	}
	_bf._gg &= 0xFFFFFFFFFF
	return nil
}

func (_fg *Decoder) mpsExchange(_gfa *DecoderStats, _fee int32) int {
	_bec := _gfa._bdb[_gfa._cd]
	if _fg._eg < _e[_fee][0] {
		if _e[_fee][3] == 1 {
			_gfa.toggleMps()
		}
		_gfa.setEntry(int(_e[_fee][2]))
		return int(1 - _bec)
	}
	_gfa.setEntry(int(_e[_fee][1]))
	return int(_bec)
}

func (_fea *DecoderStats) Copy() *DecoderStats {
	_bfd := &DecoderStats{_bd: _fea._bd, _fgf: make([]byte, _fea._bd)}
	copy(_bfd._fgf, _fea._fgf)
	return _bfd
}

type Decoder struct {
	ContextSize          []uint32
	ReferedToContextSize []uint32
	_dg                  *_d.Reader
	_fd                  uint8
	_gg                  uint64
	_eg                  uint32
	_be                  int64
	_ef                  int32
	_cf                  int32
	_gd                  int64
}

var _e = [][4]uint32{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}

func (_ca *Decoder) lpsExchange(_age *DecoderStats, _aa int32, _dge uint32) int {
	_cfe := _age.getMps()
	if _ca._eg < _dge {
		_age.setEntry(int(_e[_aa][1]))
		_ca._eg = _dge
		return int(_cfe)
	}
	if _e[_aa][3] == 1 {
		_age.toggleMps()
	}
	_age.setEntry(int(_e[_aa][2]))
	_ca._eg = _dge
	return int(1 - _cfe)
}

func (_df *Decoder) init() error {
	_df._gd = _df._dg.AbsolutePosition()
	_db, _cba := _df._dg.ReadByte()
	if _cba != nil {
		_gc.Log.Debug("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076", _cba)
		return _cba
	}
	_df._fd = _db
	_df._gg = uint64(_db) << 16
	if _cba = _df.readByte(); _cba != nil {
		return _cba
	}
	_df._gg <<= 7
	_df._ef -= 7
	_df._eg = 0x8000
	_df._cf++
	return nil
}

func (_gge *Decoder) renormalize() error {
	for {
		if _gge._ef == 0 {
			if _ggd := _gge.readByte(); _ggd != nil {
				return _ggd
			}
		}
		_gge._eg <<= 1
		_gge._gg <<= 1
		_gge._ef--
		if (_gge._eg & 0x8000) != 0 {
			break
		}
	}
	_gge._gg &= 0xffffffff
	return nil
}

func New(r *_d.Reader) (*Decoder, error) {
	_a := &Decoder{_dg: r, ContextSize: []uint32{16, 13, 10, 10}, ReferedToContextSize: []uint32{13, 10}}
	if _cb := _a.init(); _cb != nil {
		return nil, _cb
	}
	return _a, nil
}
func (_deb *DecoderStats) SetIndex(index int32) { _deb._cd = index }
func (_aag *DecoderStats) toggleMps()           { _aag._bdb[_aag._cd] ^= 1 }
func (_cae *DecoderStats) cx() byte             { return _cae._fgf[_cae._cd] }
func NewStats(contextSize int32, index int32) *DecoderStats {
	return &DecoderStats{_cd: index, _bd: contextSize, _fgf: make([]byte, contextSize), _bdb: make([]byte, contextSize)}
}

func (_ba *Decoder) DecodeInt(stats *DecoderStats) (int32, error) {
	var (
		_efd, _bee    int32
		_ff, _ge, _ag int
		_gcb          error
	)
	if stats == nil {
		stats = NewStats(512, 1)
	}
	_ba._be = 1
	_ge, _gcb = _ba.decodeIntBit(stats)
	if _gcb != nil {
		return 0, _gcb
	}
	_ff, _gcb = _ba.decodeIntBit(stats)
	if _gcb != nil {
		return 0, _gcb
	}
	if _ff == 1 {
		_ff, _gcb = _ba.decodeIntBit(stats)
		if _gcb != nil {
			return 0, _gcb
		}
		if _ff == 1 {
			_ff, _gcb = _ba.decodeIntBit(stats)
			if _gcb != nil {
				return 0, _gcb
			}
			if _ff == 1 {
				_ff, _gcb = _ba.decodeIntBit(stats)
				if _gcb != nil {
					return 0, _gcb
				}
				if _ff == 1 {
					_ff, _gcb = _ba.decodeIntBit(stats)
					if _gcb != nil {
						return 0, _gcb
					}
					if _ff == 1 {
						_ag = 32
						_bee = 4436
					} else {
						_ag = 12
						_bee = 340
					}
				} else {
					_ag = 8
					_bee = 84
				}
			} else {
				_ag = 6
				_bee = 20
			}
		} else {
			_ag = 4
			_bee = 4
		}
	} else {
		_ag = 2
		_bee = 0
	}
	for _gdc := 0; _gdc < _ag; _gdc++ {
		_ff, _gcb = _ba.decodeIntBit(stats)
		if _gcb != nil {
			return 0, _gcb
		}
		_efd = (_efd << 1) | int32(_ff)
	}
	_efd += _bee
	if _ge == 0 {
		return _efd, nil
	} else if _ge == 1 && _efd > 0 {
		return -_efd, nil
	}
	return 0, _b.ErrOOB
}

func (_dbd *Decoder) decodeIntBit(_gcbe *DecoderStats) (int, error) {
	_gcbe.SetIndex(int32(_dbd._be))
	_bbf, _gca := _dbd.DecodeBit(_gcbe)
	if _gca != nil {
		_gc.Log.Debug("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _gca)
		return _bbf, _gca
	}
	if _dbd._be < 256 {
		_dbd._be = ((_dbd._be << uint64(1)) | int64(_bbf)) & 0x1ff
	} else {
		_dbd._be = (((_dbd._be<<uint64(1) | int64(_bbf)) & 511) | 256) & 0x1ff
	}
	return _bbf, nil
}

func (_dd *DecoderStats) Reset() {
	for _dgeg := 0; _dgeg < len(_dd._fgf); _dgeg++ {
		_dd._fgf[_dgeg] = 0
		_dd._bdb[_dgeg] = 0
	}
}
func (_ee *DecoderStats) setEntry(_dfb int) { _caf := byte(_dfb & 0x7f); _ee._fgf[_ee._cd] = _caf }
func (_ffb *DecoderStats) String() string {
	_cfb := &_g.Builder{}
	_cfb.WriteString(_ce.Sprintf("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a", len(_ffb._fgf)))
	for _dc, _bbfb := range _ffb._fgf {
		if _bbfb != 0 {
			_cfb.WriteString(_ce.Sprintf("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a", _dc, _bbfb))
		}
	}
	return _cfb.String()
}
func (_fbe *DecoderStats) getMps() byte { return _fbe._bdb[_fbe._cd] }

type DecoderStats struct {
	_cd  int32
	_bd  int32
	_fgf []byte
	_bdb []byte
}

func (_gad *Decoder) DecodeIAID(codeLen uint64, stats *DecoderStats) (int64, error) {
	_gad._be = 1
	var _fdf uint64
	for _fdf = 0; _fdf < codeLen; _fdf++ {
		stats.SetIndex(int32(_gad._be))
		_de, _bb := _gad.DecodeBit(stats)
		if _bb != nil {
			return 0, _bb
		}
		_gad._be = (_gad._be << 1) | int64(_de)
	}
	_fb := _gad._be - (1 << codeLen)
	return _fb, nil
}

func (_fe *Decoder) DecodeBit(stats *DecoderStats) (int, error) {
	var (
		_fdb int
		_fa  = _e[stats.cx()][0]
		_ced = int32(stats.cx())
	)
	defer func() { _fe._cf++ }()
	_fe._eg -= _fa
	if (_fe._gg >> 16) < uint64(_fa) {
		_fdb = _fe.lpsExchange(stats, _ced, _fa)
		if _ga := _fe.renormalize(); _ga != nil {
			return 0, _ga
		}
	} else {
		_fe._gg -= uint64(_fa) << 16
		if (_fe._eg & 0x8000) == 0 {
			_fdb = _fe.mpsExchange(stats, _ced)
			if _efb := _fe.renormalize(); _efb != nil {
				return 0, _efb
			}
		} else {
			_fdb = int(stats.getMps())
		}
	}
	return _fdb, nil
}

func (_aab *DecoderStats) Overwrite(dNew *DecoderStats) {
	for _ad := 0; _ad < len(_aab._fgf); _ad++ {
		_aab._fgf[_ad] = dNew._fgf[_ad]
		_aab._bdb[_ad] = dNew._bdb[_ad]
	}
}
