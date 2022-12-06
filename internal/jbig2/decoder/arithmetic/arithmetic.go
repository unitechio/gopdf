package arithmetic

import (
	_c "fmt"
	_dd "io"
	_ddb "strings"

	_da "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_cb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
)

type Decoder struct {
	ContextSize          []uint32
	ReferedToContextSize []uint32
	_a                   _b.StreamReader
	_e                   uint8
	_ea                  uint64
	_ca                  uint32
	_ad                  int64
	_ec                  int32
	_eb                  int32
	_ee                  int64
}

func NewStats(contextSize int32, index int32) *DecoderStats {
	return &DecoderStats{_ga: index, _daa: contextSize, _eec: make([]byte, contextSize), _dfd: make([]byte, contextSize)}
}
func New(r _b.StreamReader) (*Decoder, error) {
	_ed := &Decoder{_a: r, ContextSize: []uint32{16, 13, 10, 10}, ReferedToContextSize: []uint32{13, 10}}
	if _dg := _ed.init(); _dg != nil {
		return nil, _dg
	}
	return _ed, nil
}

type DecoderStats struct {
	_ga  int32
	_daa int32
	_eec []byte
	_dfd []byte
}

func (_edc *DecoderStats) String() string {
	_ebc := &_ddb.Builder{}
	_ebc.WriteString(_c.Sprintf("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a", len(_edc._eec)))
	for _be, _cec := range _edc._eec {
		if _cec != 0 {
			_ebc.WriteString(_c.Sprintf("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a", _be, _cec))
		}
	}
	return _ebc.String()
}

var (
	_cf = [][4]uint32{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0}, {0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0}, {0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0}, {0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0}, {0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0}, {0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0}, {0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0}, {0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0}, {0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0}, {0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0}, {0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0}, {0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0}, {0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0}, {0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0}, {0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0}, {0x5601, 46, 46, 0}}
)

func (_dfa *Decoder) init() error {
	_dfa._ee = _dfa._a.StreamPosition()
	_bg, _edd := _dfa._a.ReadByte()
	if _edd != nil {
		_da.Log.Debug("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076", _edd)
		return _edd
	}
	_dfa._e = _bg
	_dfa._ea = uint64(_bg) << 16
	if _edd = _dfa.readByte(); _edd != nil {
		return _edd
	}
	_dfa._ea <<= 7
	_dfa._ec -= 7
	_dfa._ca = 0x8000
	_dfa._eb++
	return nil
}
func (_ag *Decoder) DecodeBit(stats *DecoderStats) (int, error) {
	var (
		_af int
		_f  = _cf[stats.cx()][0]
		_g  = int32(stats.cx())
	)
	defer func() { _ag._eb++ }()
	_ag._ca -= _f
	if (_ag._ea >> 16) < uint64(_f) {
		_af = _ag.lpsExchange(stats, _g, _f)
		if _cg := _ag.renormalize(); _cg != nil {
			return 0, _cg
		}
	} else {
		_ag._ea -= uint64(_f) << 16
		if (_ag._ca & 0x8000) == 0 {
			_af = _ag.mpsExchange(stats, _g)
			if _eeb := _ag.renormalize(); _eeb != nil {
				return 0, _eeb
			}
		} else {
			_af = int(stats.getMps())
		}
	}
	return _af, nil
}
func (_ff *Decoder) readByte() error {
	if _ff._a.StreamPosition() > _ff._ee {
		if _, _adf := _ff._a.Seek(-1, _dd.SeekCurrent); _adf != nil {
			return _adf
		}
	}
	_fdb, _ebf := _ff._a.ReadByte()
	if _ebf != nil {
		return _ebf
	}
	_ff._e = _fdb
	if _ff._e == 0xFF {
		_bf, _bga := _ff._a.ReadByte()
		if _bga != nil {
			return _bga
		}
		if _bf > 0x8F {
			_ff._ea += 0xFF00
			_ff._ec = 8
			if _, _bbb := _ff._a.Seek(-2, _dd.SeekCurrent); _bbb != nil {
				return _bbb
			}
		} else {
			_ff._ea += uint64(_bf) << 9
			_ff._ec = 7
		}
	} else {
		_fdb, _ebf = _ff._a.ReadByte()
		if _ebf != nil {
			return _ebf
		}
		_ff._e = _fdb
		_ff._ea += uint64(_ff._e) << 8
		_ff._ec = 8
	}
	_ff._ea &= 0xFFFFFFFFFF
	return nil
}
func (_caf *Decoder) mpsExchange(_dcc *DecoderStats, _fec int32) int {
	_dcb := _dcc._dfd[_dcc._ga]
	if _caf._ca < _cf[_fec][0] {
		if _cf[_fec][3] == 1 {
			_dcc.toggleMps()
		}
		_dcc.setEntry(int(_cf[_fec][2]))
		return int(1 - _dcb)
	}
	_dcc.setEntry(int(_cf[_fec][1]))
	return int(_dcb)
}
func (_gb *Decoder) renormalize() error {
	for {
		if _gb._ec == 0 {
			if _ab := _gb.readByte(); _ab != nil {
				return _ab
			}
		}
		_gb._ca <<= 1
		_gb._ea <<= 1
		_gb._ec--
		if (_gb._ca & 0x8000) != 0 {
			break
		}
	}
	_gb._ea &= 0xffffffff
	return nil
}
func (_ffe *Decoder) lpsExchange(_dff *DecoderStats, _ffd int32, _cae uint32) int {
	_bgc := _dff.getMps()
	if _ffe._ca < _cae {
		_dff.setEntry(int(_cf[_ffd][1]))
		_ffe._ca = _cae
		return int(_bgc)
	}
	if _cf[_ffd][3] == 1 {
		_dff.toggleMps()
	}
	_dff.setEntry(int(_cf[_ffd][2]))
	_ffe._ca = _cae
	return int(1 - _bgc)
}
func (_cfd *DecoderStats) getMps() byte { return _cfd._dfd[_cfd._ga] }
func (_afe *DecoderStats) Reset() {
	for _ecb := 0; _ecb < len(_afe._eec); _ecb++ {
		_afe._eec[_ecb] = 0
		_afe._dfd[_ecb] = 0
	}
}
func (_fd *Decoder) DecodeIAID(codeLen uint64, stats *DecoderStats) (int64, error) {
	_fd._ad = 1
	var _dad uint64
	for _dad = 0; _dad < codeLen; _dad++ {
		stats.SetIndex(int32(_fd._ad))
		_fb, _ce := _fd.DecodeBit(stats)
		if _ce != nil {
			return 0, _ce
		}
		_fd._ad = (_fd._ad << 1) | int64(_fb)
	}
	_ba := _fd._ad - (1 << codeLen)
	return _ba, nil
}
func (_cafg *DecoderStats) Overwrite(dNew *DecoderStats) {
	for _gd := 0; _gd < len(_cafg._eec); _gd++ {
		_cafg._eec[_gd] = dNew._eec[_gd]
		_cafg._dfd[_gd] = dNew._dfd[_gd]
	}
}
func (_ge *Decoder) DecodeInt(stats *DecoderStats) (int32, error) {
	var (
		_gg, _ef       int32
		_geb, _bb, _df int
		_gee           error
	)
	if stats == nil {
		stats = NewStats(512, 1)
	}
	_ge._ad = 1
	_bb, _gee = _ge.decodeIntBit(stats)
	if _gee != nil {
		return 0, _gee
	}
	_geb, _gee = _ge.decodeIntBit(stats)
	if _gee != nil {
		return 0, _gee
	}
	if _geb == 1 {
		_geb, _gee = _ge.decodeIntBit(stats)
		if _gee != nil {
			return 0, _gee
		}
		if _geb == 1 {
			_geb, _gee = _ge.decodeIntBit(stats)
			if _gee != nil {
				return 0, _gee
			}
			if _geb == 1 {
				_geb, _gee = _ge.decodeIntBit(stats)
				if _gee != nil {
					return 0, _gee
				}
				if _geb == 1 {
					_geb, _gee = _ge.decodeIntBit(stats)
					if _gee != nil {
						return 0, _gee
					}
					if _geb == 1 {
						_df = 32
						_ef = 4436
					} else {
						_df = 12
						_ef = 340
					}
				} else {
					_df = 8
					_ef = 84
				}
			} else {
				_df = 6
				_ef = 20
			}
		} else {
			_df = 4
			_ef = 4
		}
	} else {
		_df = 2
		_ef = 0
	}
	for _fe := 0; _fe < _df; _fe++ {
		_geb, _gee = _ge.decodeIntBit(stats)
		if _gee != nil {
			return 0, _gee
		}
		_gg = (_gg << 1) | int32(_geb)
	}
	_gg += _ef
	if _bb == 0 {
		return _gg, nil
	} else if _bb == 1 && _gg > 0 {
		return -_gg, nil
	}
	return 0, _cb.ErrOOB
}
func (_bd *DecoderStats) SetIndex(index int32) { _bd._ga = index }
func (_ecg *DecoderStats) toggleMps()          { _ecg._dfd[_ecg._ga] ^= 1 }
func (_dac *DecoderStats) cx() byte            { return _dac._eec[_dac._ga] }
func (_ae *DecoderStats) Copy() *DecoderStats {
	_bc := &DecoderStats{_daa: _ae._daa, _eec: make([]byte, _ae._daa)}
	for _ged := 0; _ged < len(_ae._eec); _ged++ {
		_bc._eec[_ged] = _ae._eec[_ged]
	}
	return _bc
}
func (_bac *Decoder) decodeIntBit(_dc *DecoderStats) (int, error) {
	_dc.SetIndex(int32(_bac._ad))
	_dda, _eg := _bac.DecodeBit(_dc)
	if _eg != nil {
		_da.Log.Debug("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _eg)
		return _dda, _eg
	}
	if _bac._ad < 256 {
		_bac._ad = ((_bac._ad << uint64(1)) | int64(_dda)) & 0x1ff
	} else {
		_bac._ad = (((_bac._ad<<uint64(1) | int64(_dda)) & 511) | 256) & 0x1ff
	}
	return _dda, nil
}
func (_fed *DecoderStats) setEntry(_efb int) {
	_gbe := byte(_efb & 0x7f)
	_fed._eec[_fed._ga] = _gbe
}
