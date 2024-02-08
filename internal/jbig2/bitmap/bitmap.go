package bitmap

import (
	_db "encoding/binary"
	_aa "image"
	_ea "math"
	_ff "sort"
	_c "strings"
	_f "testing"

	_ca "bitbucket.org/shenghui0779/gopdf/common"
	_ce "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_g "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_ac "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_e "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_d "github.com/stretchr/testify/require"
)

func _acbb(_bacfg *Bitmap, _ddag, _dbgb int, _fbafd, _gcdg int, _efda RasterOperator) {
	var (
		_ceded        int
		_dddae        byte
		_dbafc, _adef int
		_gfegc        int
	)
	_bbbg := _fbafd >> 3
	_caaba := _fbafd & 7
	if _caaba > 0 {
		_dddae = _acbba[_caaba]
	}
	_ceded = _bacfg.RowStride*_dbgb + (_ddag >> 3)
	switch _efda {
	case PixClr:
		for _dbafc = 0; _dbafc < _gcdg; _dbafc++ {
			_gfegc = _ceded + _dbafc*_bacfg.RowStride
			for _adef = 0; _adef < _bbbg; _adef++ {
				_bacfg.Data[_gfegc] = 0x0
				_gfegc++
			}
			if _caaba > 0 {
				_bacfg.Data[_gfegc] = _bbgf(_bacfg.Data[_gfegc], 0x0, _dddae)
			}
		}
	case PixSet:
		for _dbafc = 0; _dbafc < _gcdg; _dbafc++ {
			_gfegc = _ceded + _dbafc*_bacfg.RowStride
			for _adef = 0; _adef < _bbbg; _adef++ {
				_bacfg.Data[_gfegc] = 0xff
				_gfegc++
			}
			if _caaba > 0 {
				_bacfg.Data[_gfegc] = _bbgf(_bacfg.Data[_gfegc], 0xff, _dddae)
			}
		}
	case PixNotDst:
		for _dbafc = 0; _dbafc < _gcdg; _dbafc++ {
			_gfegc = _ceded + _dbafc*_bacfg.RowStride
			for _adef = 0; _adef < _bbbg; _adef++ {
				_bacfg.Data[_gfegc] = ^_bacfg.Data[_gfegc]
				_gfegc++
			}
			if _caaba > 0 {
				_bacfg.Data[_gfegc] = _bbgf(_bacfg.Data[_gfegc], ^_bacfg.Data[_gfegc], _dddae)
			}
		}
	}
}

type RasterOperator int

func TstWordBitmap(t *_f.T, scale ...int) *Bitmap {
	_ggab := 1
	if len(scale) > 0 {
		_ggab = scale[0]
	}
	_cabf := 3
	_aadgg := 9 + 7 + 15 + 2*_cabf
	_bfcca := 5 + _cabf + 5
	_cffbd := New(_aadgg*_ggab, _bfcca*_ggab)
	_cagf := &Bitmaps{}
	var _afgfa *int
	_cabf *= _ggab
	_bfaea := 0
	_afgfa = &_bfaea
	_bbfg := 0
	_ecac := TstDSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, _cabf)
	_ecac = TstISymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstTSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, _cabf)
	_ecac = TstNSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstWSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 0)
	*_afgfa = 0
	_bbfg = 5*_ggab + _cabf
	_ecac = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstRSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, _cabf)
	_ecac = TstNSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstESymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstVSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstESymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 1*_ggab)
	_ecac = TstRSymbol(t, scale...)
	TstAddSymbol(t, _cagf, _ecac, _afgfa, _bbfg, 0)
	TstWriteSymbols(t, _cagf, _cffbd)
	return _cffbd
}
func TstTSymbol(t *_f.T, scale ...int) *Bitmap {
	_ddde, _fdbc := NewWithData(5, 5, []byte{0xF8, 0x20, 0x20, 0x20, 0x20})
	_d.NoError(t, _fdbc)
	return TstGetScaledSymbol(t, _ddde, scale...)
}
func _cfaba(_ecgfb *Bitmap, _ffgf, _daec, _gfeg, _fgefd int, _eafcf RasterOperator, _gaebc *Bitmap, _efdf, _cgfd int) error {
	var (
		_efadf         bool
		_bcaae         bool
		_bafe          int
		_afgca         int
		_gadbe         int
		_cfgdg         bool
		_facf          byte
		_ceef          int
		_ceaa          int
		_eef           int
		_cfgde, _edfac int
	)
	_aacd := 8 - (_ffgf & 7)
	_gfff := _fbdg[_aacd]
	_ecdb := _ecgfb.RowStride*_daec + (_ffgf >> 3)
	_cgag := _gaebc.RowStride*_cgfd + (_efdf >> 3)
	if _gfeg < _aacd {
		_efadf = true
		_gfff &= _acbba[8-_aacd+_gfeg]
	}
	if !_efadf {
		_bafe = (_gfeg - _aacd) >> 3
		if _bafe > 0 {
			_bcaae = true
			_afgca = _ecdb + 1
			_gadbe = _cgag + 1
		}
	}
	_ceef = (_ffgf + _gfeg) & 7
	if !(_efadf || _ceef == 0) {
		_cfgdg = true
		_facf = _acbba[_ceef]
		_ceaa = _ecdb + 1 + _bafe
		_eef = _cgag + 1 + _bafe
	}
	switch _eafcf {
	case PixSrc:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], _gaebc.Data[_cgag], _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] = _gaebc.Data[_gadbe+_edfac]
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], _gaebc.Data[_eef], _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixNotSrc:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], ^_gaebc.Data[_cgag], _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] = ^_gaebc.Data[_gadbe+_edfac]
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], ^_gaebc.Data[_eef], _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixSrcOrDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], _gaebc.Data[_cgag]|_ecgfb.Data[_ecdb], _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] |= _gaebc.Data[_gadbe+_edfac]
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], _gaebc.Data[_eef]|_ecgfb.Data[_ceaa], _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixSrcAndDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], _gaebc.Data[_cgag]&_ecgfb.Data[_ecdb], _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] &= _gaebc.Data[_gadbe+_edfac]
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], _gaebc.Data[_eef]&_ecgfb.Data[_ceaa], _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixSrcXorDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], _gaebc.Data[_cgag]^_ecgfb.Data[_ecdb], _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] ^= _gaebc.Data[_gadbe+_edfac]
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], _gaebc.Data[_eef]^_ecgfb.Data[_ceaa], _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixNotSrcOrDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], ^(_gaebc.Data[_cgag])|_ecgfb.Data[_ecdb], _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] |= ^(_gaebc.Data[_gadbe+_edfac])
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], ^(_gaebc.Data[_eef])|_ecgfb.Data[_ceaa], _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixNotSrcAndDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], ^(_gaebc.Data[_cgag])&_ecgfb.Data[_ecdb], _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] &= ^_gaebc.Data[_gadbe+_edfac]
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], ^(_gaebc.Data[_eef])&_ecgfb.Data[_ceaa], _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixSrcOrNotDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], _gaebc.Data[_cgag]|^(_ecgfb.Data[_ecdb]), _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] = _gaebc.Data[_gadbe+_edfac] | ^(_ecgfb.Data[_afgca+_edfac])
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], _gaebc.Data[_eef]|^(_ecgfb.Data[_ceaa]), _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixSrcAndNotDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], _gaebc.Data[_cgag]&^(_ecgfb.Data[_ecdb]), _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] = _gaebc.Data[_gadbe+_edfac] &^ (_ecgfb.Data[_afgca+_edfac])
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], _gaebc.Data[_eef]&^(_ecgfb.Data[_ceaa]), _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], ^(_gaebc.Data[_cgag] | _ecgfb.Data[_ecdb]), _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] = ^(_gaebc.Data[_gadbe+_edfac] | _ecgfb.Data[_afgca+_edfac])
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], ^(_gaebc.Data[_eef] | _ecgfb.Data[_ceaa]), _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], ^(_gaebc.Data[_cgag] & _ecgfb.Data[_ecdb]), _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] = ^(_gaebc.Data[_gadbe+_edfac] & _ecgfb.Data[_afgca+_edfac])
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], ^(_gaebc.Data[_eef] & _ecgfb.Data[_ceaa]), _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
			_ecgfb.Data[_ecdb] = _bbgf(_ecgfb.Data[_ecdb], ^(_gaebc.Data[_cgag] ^ _ecgfb.Data[_ecdb]), _gfff)
			_ecdb += _ecgfb.RowStride
			_cgag += _gaebc.RowStride
		}
		if _bcaae {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				for _edfac = 0; _edfac < _bafe; _edfac++ {
					_ecgfb.Data[_afgca+_edfac] = ^(_gaebc.Data[_gadbe+_edfac] ^ _ecgfb.Data[_afgca+_edfac])
				}
				_afgca += _ecgfb.RowStride
				_gadbe += _gaebc.RowStride
			}
		}
		if _cfgdg {
			for _cfgde = 0; _cfgde < _fgefd; _cfgde++ {
				_ecgfb.Data[_ceaa] = _bbgf(_ecgfb.Data[_ceaa], ^(_gaebc.Data[_eef] ^ _ecgfb.Data[_ceaa]), _facf)
				_ceaa += _ecgfb.RowStride
				_eef += _gaebc.RowStride
			}
		}
	default:
		_ca.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _eafcf)
		return _e.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0056\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004c\u006f\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_fcb *Bitmap) Copy() *Bitmap {
	_ffcc := make([]byte, len(_fcb.Data))
	copy(_ffcc, _fcb.Data)
	return &Bitmap{Width: _fcb.Width, Height: _fcb.Height, RowStride: _fcb.RowStride, Data: _ffcc, Color: _fcb.Color, Text: _fcb.Text, BitmapNumber: _fcb.BitmapNumber, Special: _fcb.Special}
}
func _beaad(_cbdf *Bitmap, _efga *_ac.Stack, _aecg, _bgdb, _bgcc int) (_bebf *_aa.Rectangle, _efea error) {
	const _dffd = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _cbdf == nil {
		return nil, _e.Error(_dffd, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _efga == nil {
		return nil, _e.Error(_dffd, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	switch _bgcc {
	case 4:
		if _bebf, _efea = _ddgf(_cbdf, _efga, _aecg, _bgdb); _efea != nil {
			return nil, _e.Wrap(_efea, _dffd, "")
		}
		return _bebf, nil
	case 8:
		if _bebf, _efea = _egbea(_cbdf, _efga, _aecg, _bgdb); _efea != nil {
			return nil, _e.Wrap(_efea, _dffd, "")
		}
		return _bebf, nil
	default:
		return nil, _e.Errorf(_dffd, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u0069\u0073 \u006eo\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038\u003a\u0020\u0027\u0025\u0064\u0027", _bgcc)
	}
}
func (_feac *Bitmap) Zero() bool {
	_edbb := _feac.Width / 8
	_fbgd := _feac.Width & 7
	var _fabc byte
	if _fbgd != 0 {
		_fabc = byte(0xff << uint(8-_fbgd))
	}
	var _ecf, _eab, _aafb int
	for _eab = 0; _eab < _feac.Height; _eab++ {
		_ecf = _feac.RowStride * _eab
		for _aafb = 0; _aafb < _edbb; _aafb, _ecf = _aafb+1, _ecf+1 {
			if _feac.Data[_ecf] != 0 {
				return false
			}
		}
		if _fbgd > 0 {
			if _feac.Data[_ecf]&_fabc != 0 {
				return false
			}
		}
	}
	return true
}
func (_fbbc Points) Get(i int) (Point, error) {
	if i > len(_fbbc)-1 {
		return Point{}, _e.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _fbbc[i], nil
}
func (_dgac CombinationOperator) String() string {
	var _ffgbc string
	switch _dgac {
	case CmbOpOr:
		_ffgbc = "\u004f\u0052"
	case CmbOpAnd:
		_ffgbc = "\u0041\u004e\u0044"
	case CmbOpXor:
		_ffgbc = "\u0058\u004f\u0052"
	case CmbOpXNor:
		_ffgbc = "\u0058\u004e\u004f\u0052"
	case CmbOpReplace:
		_ffgbc = "\u0052E\u0050\u004c\u0041\u0043\u0045"
	case CmbOpNot:
		_ffgbc = "\u004e\u004f\u0054"
	}
	return _ffgbc
}
func (_ecbfc *byHeight) Len() int { return len(_ecbfc.Values) }
func (_bdaf MorphProcess) verify(_fafa int, _bgab, _ecgfe *int) error {
	const _dadb = "\u004d\u006f\u0072\u0070hP\u0072\u006f\u0063\u0065\u0073\u0073\u002e\u0076\u0065\u0072\u0069\u0066\u0079"
	switch _bdaf.Operation {
	case MopDilation, MopErosion, MopOpening, MopClosing:
		if len(_bdaf.Arguments) != 2 {
			return _e.Error(_dadb, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0064\u0027\u002c\u0020\u0027\u0065\u0027\u002c \u0027\u006f\u0027\u002c\u0020\u0027\u0063\u0027\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0032\u0020\u0061r\u0067\u0075\u006d\u0065\u006et\u0073")
		}
		_bccb, _cfega := _bdaf.getWidthHeight()
		if _bccb <= 0 || _cfega <= 0 {
			return _e.Error(_dadb, "O\u0070er\u0061t\u0069o\u006e\u003a\u0020\u0027\u0064'\u002c\u0020\u0027e\u0027\u002c\u0020\u0027\u006f'\u002c\u0020\u0027c\u0027\u0020\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0062\u006f\u0074h w\u0069\u0064\u0074\u0068\u0020\u0061n\u0064\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0074\u006f\u0020b\u0065 \u003e\u003d\u0020\u0030")
		}
	case MopRankBinaryReduction:
		_eged := len(_bdaf.Arguments)
		*_bgab += _eged
		if _eged < 1 || _eged > 4 {
			return _e.Error(_dadb, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0072\u0027\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020\u006d\u006fs\u0074\u0020\u0034\u0020\u0061\u0072g\u0075\u006d\u0065n\u0074\u0073")
		}
		for _egba := 0; _egba < _eged; _egba++ {
			if _bdaf.Arguments[_egba] < 1 || _bdaf.Arguments[_egba] > 4 {
				return _e.Error(_dadb, "\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0052\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 \u00280\u002c\u0020\u0034\u003e")
			}
		}
	case MopReplicativeBinaryExpansion:
		if len(_bdaf.Arguments) == 0 {
			return _e.Error(_dadb, "\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074i\u0076\u0065\u0042in\u0061\u0072\u0079\u0045\u0078\u0070a\u006e\u0073\u0069\u006f\u006e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020o\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_efcf := _bdaf.Arguments[0]
		if _efcf != 2 && _efcf != 4 && _efcf != 8 {
			return _e.Error(_dadb, "R\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u0076\u0065\u0042\u0069\u006e\u0061\u0072\u0079\u0045\u0078\u0070\u0061\u006e\u0073\u0069\u006f\u006e\u0020m\u0075s\u0074\u0020\u0062\u0065 \u006f\u0066 \u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
		}
		*_bgab -= _fdde[_efcf/4]
	case MopAddBorder:
		if len(_bdaf.Arguments) == 0 {
			return _e.Error(_dadb, "\u0041\u0064\u0064B\u006f\u0072\u0064\u0065r\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_gfbg := _bdaf.Arguments[0]
		if _fafa > 0 {
			return _e.Error(_dadb, "\u0041\u0064\u0064\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020f\u0069\u0072\u0073\u0074\u0020\u006d\u006f\u0072\u0070\u0068\u0020\u0070\u0072o\u0063\u0065\u0073\u0073")
		}
		if _gfbg < 1 {
			return _e.Error(_dadb, "\u0041\u0064\u0064\u0042o\u0072\u0064\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065 \u006co\u0077\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
		}
		*_ecgfe = _gfbg
	}
	return nil
}
func (_efd *Bitmap) SetPixel(x, y int, pixel byte) error {
	_bda := _efd.GetByteIndex(x, y)
	if _bda > len(_efd.Data)-1 {
		return _e.Errorf("\u0053\u0065\u0074\u0050\u0069\u0078\u0065\u006c", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", _bda)
	}
	_cdcb := _efd.GetBitOffset(x)
	_bfgd := uint(7 - _cdcb)
	_gfg := _efd.Data[_bda]
	var _abgf byte
	if pixel == 1 {
		_abgf = _gfg | (pixel & 0x01 << _bfgd)
	} else {
		_abgf = _gfg &^ (1 << _bfgd)
	}
	_efd.Data[_bda] = _abgf
	return nil
}
func DilateBrick(d, s *Bitmap, hSize, vSize int) (*Bitmap, error) { return _dgdg(d, s, hSize, vSize) }
func _bcc() (_egb []byte) {
	_egb = make([]byte, 256)
	for _adc := 0; _adc < 256; _adc++ {
		_gbc := byte(_adc)
		_egb[_gbc] = (_gbc & 0x01) | ((_gbc & 0x04) >> 1) | ((_gbc & 0x10) >> 2) | ((_gbc & 0x40) >> 3) | ((_gbc & 0x02) << 3) | ((_gbc & 0x08) << 2) | ((_gbc & 0x20) << 1) | (_gbc & 0x80)
	}
	return _egb
}
func _cfgc(_eegg, _feee *Bitmap, _fdbg *Selection) (*Bitmap, error) {
	const _acdg = "\u0065\u0072\u006fd\u0065"
	var (
		_gbcf error
		_bada *Bitmap
	)
	_eegg, _gbcf = _begeb(_eegg, _feee, _fdbg, &_bada)
	if _gbcf != nil {
		return nil, _e.Wrap(_gbcf, _acdg, "")
	}
	if _gbcf = _eegg.setAll(); _gbcf != nil {
		return nil, _e.Wrap(_gbcf, _acdg, "")
	}
	var _aebg SelectionValue
	for _agd := 0; _agd < _fdbg.Height; _agd++ {
		for _ffccd := 0; _ffccd < _fdbg.Width; _ffccd++ {
			_aebg = _fdbg.Data[_agd][_ffccd]
			if _aebg == SelHit {
				_gbcf = _gegg(_eegg, _fdbg.Cx-_ffccd, _fdbg.Cy-_agd, _feee.Width, _feee.Height, PixSrcAndDst, _bada, 0, 0)
				if _gbcf != nil {
					return nil, _e.Wrap(_gbcf, _acdg, "")
				}
			}
		}
	}
	if MorphBC == SymmetricMorphBC {
		return _eegg, nil
	}
	_dacba, _afdd, _eedaa, _gbca := _fdbg.findMaxTranslations()
	if _dacba > 0 {
		if _gbcf = _eegg.RasterOperation(0, 0, _dacba, _feee.Height, PixClr, nil, 0, 0); _gbcf != nil {
			return nil, _e.Wrap(_gbcf, _acdg, "\u0078\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _eedaa > 0 {
		if _gbcf = _eegg.RasterOperation(_feee.Width-_eedaa, 0, _eedaa, _feee.Height, PixClr, nil, 0, 0); _gbcf != nil {
			return nil, _e.Wrap(_gbcf, _acdg, "\u0078\u006e\u0020\u003e\u0020\u0030")
		}
	}
	if _afdd > 0 {
		if _gbcf = _eegg.RasterOperation(0, 0, _feee.Width, _afdd, PixClr, nil, 0, 0); _gbcf != nil {
			return nil, _e.Wrap(_gbcf, _acdg, "\u0079\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _gbca > 0 {
		if _gbcf = _eegg.RasterOperation(0, _feee.Height-_gbca, _feee.Width, _gbca, PixClr, nil, 0, 0); _gbcf != nil {
			return nil, _e.Wrap(_gbcf, _acdg, "\u0079\u006e\u0020\u003e\u0020\u0030")
		}
	}
	return _eegg, nil
}
func (_bdfb *ClassedPoints) XAtIndex(i int) float32 { return (*_bdfb.Points)[_bdfb.IntSlice[i]].X }
func init() {
	for _edb := 0; _edb < 256; _edb++ {
		_fde[_edb] = uint8(_edb&0x1) + (uint8(_edb>>1) & 0x1) + (uint8(_edb>>2) & 0x1) + (uint8(_edb>>3) & 0x1) + (uint8(_edb>>4) & 0x1) + (uint8(_edb>>5) & 0x1) + (uint8(_edb>>6) & 0x1) + (uint8(_edb>>7) & 0x1)
	}
}
func (_dfdeb *byWidth) Len() int { return len(_dfdeb.Values) }
func _gcg(_aaaa, _acfac int) int {
	if _aaaa < _acfac {
		return _aaaa
	}
	return _acfac
}
func NewWithUnpaddedData(width, height int, data []byte) (*Bitmap, error) {
	const _egf = "\u004e\u0065\u0077\u0057it\u0068\u0055\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	_dcf := _eced(width, height)
	_dcf.Data = data
	if _adeag := ((width * height) + 7) >> 3; len(data) < _adeag {
		return nil, _e.Errorf(_egf, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(data), _adeag)
	}
	if _gac := _dcf.addPadBits(); _gac != nil {
		return nil, _e.Wrap(_gac, _egf, "")
	}
	return _dcf, nil
}
func CombineBytes(oldByte, newByte byte, op CombinationOperator) byte {
	return _eeabd(oldByte, newByte, op)
}
func _cdf() (_gca [256]uint32) {
	for _bf := 0; _bf < 256; _bf++ {
		if _bf&0x01 != 0 {
			_gca[_bf] |= 0xf
		}
		if _bf&0x02 != 0 {
			_gca[_bf] |= 0xf0
		}
		if _bf&0x04 != 0 {
			_gca[_bf] |= 0xf00
		}
		if _bf&0x08 != 0 {
			_gca[_bf] |= 0xf000
		}
		if _bf&0x10 != 0 {
			_gca[_bf] |= 0xf0000
		}
		if _bf&0x20 != 0 {
			_gca[_bf] |= 0xf00000
		}
		if _bf&0x40 != 0 {
			_gca[_bf] |= 0xf000000
		}
		if _bf&0x80 != 0 {
			_gca[_bf] |= 0xf0000000
		}
	}
	return _gca
}
func _fbc() (_fbg [256]uint64) {
	for _bee := 0; _bee < 256; _bee++ {
		if _bee&0x01 != 0 {
			_fbg[_bee] |= 0xff
		}
		if _bee&0x02 != 0 {
			_fbg[_bee] |= 0xff00
		}
		if _bee&0x04 != 0 {
			_fbg[_bee] |= 0xff0000
		}
		if _bee&0x08 != 0 {
			_fbg[_bee] |= 0xff000000
		}
		if _bee&0x10 != 0 {
			_fbg[_bee] |= 0xff00000000
		}
		if _bee&0x20 != 0 {
			_fbg[_bee] |= 0xff0000000000
		}
		if _bee&0x40 != 0 {
			_fbg[_bee] |= 0xff000000000000
		}
		if _bee&0x80 != 0 {
			_fbg[_bee] |= 0xff00000000000000
		}
	}
	return _fbg
}
func (_ggfgb *Bitmap) setBit(_cabg int) { _ggfgb.Data[(_cabg >> 3)] |= 0x80 >> uint(_cabg&7) }
func (_gccba *ClassedPoints) GetIntYByClass(i int) (int, error) {
	const _cbfd = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _gccba.IntSlice.Size() {
		return 0, _e.Errorf(_cbfd, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_gccba.YAtIndex(i)), nil
}
func (_bdfbc *Bitmap) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _gegg(_bdfbc, dx, dy, dw, dh, op, src, sx, sy)
}
func _adac(_dfaad *Bitmap, _edfc, _bceda, _cedg, _fabdf int, _fdbd RasterOperator, _gafbc *Bitmap, _dcgf, _acabd int) error {
	var (
		_gdfga       byte
		_aaag        int
		_bedf        int
		_afea, _adge int
		_dddc, _cggd int
	)
	_agca := _cedg >> 3
	_adfa := _cedg & 7
	if _adfa > 0 {
		_gdfga = _acbba[_adfa]
	}
	_aaag = _gafbc.RowStride*_acabd + (_dcgf >> 3)
	_bedf = _dfaad.RowStride*_bceda + (_edfc >> 3)
	switch _fdbd {
	case PixSrc:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] = _gafbc.Data[_afea]
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], _gafbc.Data[_afea], _gdfga)
			}
		}
	case PixNotSrc:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] = ^(_gafbc.Data[_afea])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], ^_gafbc.Data[_afea], _gdfga)
			}
		}
	case PixSrcOrDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] |= _gafbc.Data[_afea]
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], _gafbc.Data[_afea]|_dfaad.Data[_adge], _gdfga)
			}
		}
	case PixSrcAndDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] &= _gafbc.Data[_afea]
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], _gafbc.Data[_afea]&_dfaad.Data[_adge], _gdfga)
			}
		}
	case PixSrcXorDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] ^= _gafbc.Data[_afea]
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], _gafbc.Data[_afea]^_dfaad.Data[_adge], _gdfga)
			}
		}
	case PixNotSrcOrDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] |= ^(_gafbc.Data[_afea])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], ^(_gafbc.Data[_afea])|_dfaad.Data[_adge], _gdfga)
			}
		}
	case PixNotSrcAndDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] &= ^(_gafbc.Data[_afea])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], ^(_gafbc.Data[_afea])&_dfaad.Data[_adge], _gdfga)
			}
		}
	case PixSrcOrNotDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] = _gafbc.Data[_afea] | ^(_dfaad.Data[_adge])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], _gafbc.Data[_afea]|^(_dfaad.Data[_adge]), _gdfga)
			}
		}
	case PixSrcAndNotDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] = _gafbc.Data[_afea] &^ (_dfaad.Data[_adge])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], _gafbc.Data[_afea]&^(_dfaad.Data[_adge]), _gdfga)
			}
		}
	case PixNotPixSrcOrDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] = ^(_gafbc.Data[_afea] | _dfaad.Data[_adge])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], ^(_gafbc.Data[_afea] | _dfaad.Data[_adge]), _gdfga)
			}
		}
	case PixNotPixSrcAndDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] = ^(_gafbc.Data[_afea] & _dfaad.Data[_adge])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], ^(_gafbc.Data[_afea] & _dfaad.Data[_adge]), _gdfga)
			}
		}
	case PixNotPixSrcXorDst:
		for _dddc = 0; _dddc < _fabdf; _dddc++ {
			_afea = _aaag + _dddc*_gafbc.RowStride
			_adge = _bedf + _dddc*_dfaad.RowStride
			for _cggd = 0; _cggd < _agca; _cggd++ {
				_dfaad.Data[_adge] = ^(_gafbc.Data[_afea] ^ _dfaad.Data[_adge])
				_adge++
				_afea++
			}
			if _adfa > 0 {
				_dfaad.Data[_adge] = _bbgf(_dfaad.Data[_adge], ^(_gafbc.Data[_afea] ^ _dfaad.Data[_adge]), _gdfga)
			}
		}
	default:
		_ca.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _fdbd)
		return _e.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004co\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_baae *Points) AddPoint(x, y float32) { *_baae = append(*_baae, Point{x, y}) }
func HausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH int) (bool, error) {
	const _fgcf = "\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_abdg, _ege := p1.Width, p1.Height
	_fef, _bcfeg := p3.Width, p3.Height
	if _ac.Abs(_abdg-_fef) > maxDiffW {
		return false, nil
	}
	if _ac.Abs(_ege-_bcfeg) > maxDiffH {
		return false, nil
	}
	_aee := int(delX + _ac.Sign(delX)*0.5)
	_aafe := int(delY + _ac.Sign(delY)*0.5)
	var _afbda error
	_gbe := p1.CreateTemplate()
	if _afbda = _gbe.RasterOperation(0, 0, _abdg, _ege, PixSrc, p1, 0, 0); _afbda != nil {
		return false, _e.Wrap(_afbda, _fgcf, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _afbda = _gbe.RasterOperation(_aee, _aafe, _abdg, _ege, PixNotSrcAndDst, p4, 0, 0); _afbda != nil {
		return false, _e.Wrap(_afbda, _fgcf, "\u0021p\u0034\u0020\u0026\u0020\u0074")
	}
	if _gbe.Zero() {
		return false, nil
	}
	if _afbda = _gbe.RasterOperation(_aee, _aafe, _fef, _bcfeg, PixSrc, p3, 0, 0); _afbda != nil {
		return false, _e.Wrap(_afbda, _fgcf, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _afbda = _gbe.RasterOperation(0, 0, _fef, _bcfeg, PixNotSrcAndDst, p2, 0, 0); _afbda != nil {
		return false, _e.Wrap(_afbda, _fgcf, "\u0021p\u0032\u0020\u0026\u0020\u0074")
	}
	return _gbe.Zero(), nil
}
func (_fadc *Bitmap) SetDefaultPixel() {
	for _bcf := range _fadc.Data {
		_fadc.Data[_bcf] = byte(0xff)
	}
}

var MorphBC BoundaryCondition
var (
	_acbba = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_fbdg  = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func (_aafga Points) GetIntY(i int) (int, error) {
	if i >= len(_aafga) {
		return 0, _e.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0059", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_aafga[i].Y), nil
}
func (_edddc *Bitmap) ConnComponents(bms *Bitmaps, connectivity int) (_efbd *Boxes, _gadc error) {
	const _gbgbd = "B\u0069\u0074\u006d\u0061p.\u0043o\u006e\u006e\u0043\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0073"
	if _edddc == nil {
		return nil, _e.Error(_gbgbd, "\u0070r\u006f\u0076\u0069\u0064e\u0064\u0020\u0065\u006d\u0070t\u0079 \u0027b\u0027\u0020\u0062\u0069\u0074\u006d\u0061p")
	}
	if connectivity != 4 && connectivity != 8 {
		return nil, _e.Error(_gbgbd, "\u0063\u006f\u006ene\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006e\u006f\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038")
	}
	if bms == nil {
		if _efbd, _gadc = _edddc.connComponentsBB(connectivity); _gadc != nil {
			return nil, _e.Wrap(_gadc, _gbgbd, "")
		}
	} else {
		if _efbd, _gadc = _edddc.connComponentsBitmapsBB(bms, connectivity); _gadc != nil {
			return nil, _e.Wrap(_gadc, _gbgbd, "")
		}
	}
	return _efbd, nil
}
func (_feeg *ClassedPoints) SortByX() { _feeg._acab = _feeg.xSortFunction(); _ff.Sort(_feeg) }
func (_abdfb Points) GetGeometry(i int) (_gbdc, _becce float32, _gfcef error) {
	if i > len(_abdfb)-1 {
		return 0, 0, _e.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_bfe := _abdfb[i]
	return _bfe.X, _bfe.Y, nil
}

const _decfc = 5000

func _caagg(_gccb *Bitmap, _cfab *Bitmap, _fcdf *Selection) (*Bitmap, error) {
	var (
		_cefgd *Bitmap
		_aebdc error
	)
	_gccb, _aebdc = _begeb(_gccb, _cfab, _fcdf, &_cefgd)
	if _aebdc != nil {
		return nil, _aebdc
	}
	if _aebdc = _gccb.clearAll(); _aebdc != nil {
		return nil, _aebdc
	}
	var _gba SelectionValue
	for _bedac := 0; _bedac < _fcdf.Height; _bedac++ {
		for _acfe := 0; _acfe < _fcdf.Width; _acfe++ {
			_gba = _fcdf.Data[_bedac][_acfe]
			if _gba == SelHit {
				if _aebdc = _gccb.RasterOperation(_acfe-_fcdf.Cx, _bedac-_fcdf.Cy, _cfab.Width, _cfab.Height, PixSrcOrDst, _cefgd, 0, 0); _aebdc != nil {
					return nil, _aebdc
				}
			}
		}
	}
	return _gccb, nil
}
func (_gcee Points) YSorter() func(_cfaa, _gadd int) bool {
	return func(_adga, _adcb int) bool { return _gcee[_adga].Y < _gcee[_adcb].Y }
}
func (_ebbbb *byWidth) Less(i, j int) bool { return _ebbbb.Values[i].Width < _ebbbb.Values[j].Width }
func (_babd *Bitmaps) ClipToBitmap(s *Bitmap) (*Bitmaps, error) {
	const _cfbcf = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0043\u006c\u0069p\u0054\u006f\u0042\u0069tm\u0061\u0070"
	if _babd == nil {
		return nil, _e.Error(_cfbcf, "\u0042\u0069\u0074\u006dap\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if s == nil {
		return nil, _e.Error(_cfbcf, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	_fdccd := len(_babd.Values)
	_cagg := &Bitmaps{Values: make([]*Bitmap, _fdccd), Boxes: make([]*_aa.Rectangle, _fdccd)}
	var (
		_fedda, _dfgf *Bitmap
		_gfef         *_aa.Rectangle
		_bbee         error
	)
	for _bdafd := 0; _bdafd < _fdccd; _bdafd++ {
		if _fedda, _bbee = _babd.GetBitmap(_bdafd); _bbee != nil {
			return nil, _e.Wrap(_bbee, _cfbcf, "")
		}
		if _gfef, _bbee = _babd.GetBox(_bdafd); _bbee != nil {
			return nil, _e.Wrap(_bbee, _cfbcf, "")
		}
		if _dfgf, _bbee = s.clipRectangle(_gfef, nil); _bbee != nil {
			return nil, _e.Wrap(_bbee, _cfbcf, "")
		}
		if _dfgf, _bbee = _dfgf.And(_fedda); _bbee != nil {
			return nil, _e.Wrap(_bbee, _cfbcf, "")
		}
		_cagg.Values[_bdafd] = _dfgf
		_cagg.Boxes[_bdafd] = _gfef
	}
	return _cagg, nil
}
func _gfgf(_fedd, _faeb *Bitmap, _ecef, _fbcg int) (*Bitmap, error) {
	const _bbfa = "\u006fp\u0065\u006e\u0042\u0072\u0069\u0063k"
	if _faeb == nil {
		return nil, _e.Error(_bbfa, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _ecef < 1 && _fbcg < 1 {
		return nil, _e.Error(_bbfa, "\u0068\u0053\u0069\u007ae \u003c\u0020\u0031\u0020\u0026\u0026\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u003c \u0031")
	}
	if _ecef == 1 && _fbcg == 1 {
		return _faeb.Copy(), nil
	}
	if _ecef == 1 || _fbcg == 1 {
		var _eede error
		_bdafc := SelCreateBrick(_fbcg, _ecef, _fbcg/2, _ecef/2, SelHit)
		_fedd, _eede = _cbdc(_fedd, _faeb, _bdafc)
		if _eede != nil {
			return nil, _e.Wrap(_eede, _bbfa, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _fedd, nil
	}
	_gfbe := SelCreateBrick(1, _ecef, 0, _ecef/2, SelHit)
	_bfga := SelCreateBrick(_fbcg, 1, _fbcg/2, 0, SelHit)
	_dffg, _caga := _cfgc(nil, _faeb, _gfbe)
	if _caga != nil {
		return nil, _e.Wrap(_caga, _bbfa, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_fedd, _caga = _cfgc(_fedd, _dffg, _bfga)
	if _caga != nil {
		return nil, _e.Wrap(_caga, _bbfa, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	_, _caga = _caagg(_dffg, _fedd, _gfbe)
	if _caga != nil {
		return nil, _e.Wrap(_caga, _bbfa, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_, _caga = _caagg(_fedd, _dffg, _bfga)
	if _caga != nil {
		return nil, _e.Wrap(_caga, _bbfa, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _fedd, nil
}
func (_adba *ClassedPoints) Swap(i, j int) {
	_adba.IntSlice[i], _adba.IntSlice[j] = _adba.IntSlice[j], _adba.IntSlice[i]
}
func (_dcc *Bitmap) setAll() error {
	_egg := _gegg(_dcc, 0, 0, _dcc.Width, _dcc.Height, PixSet, nil, 0, 0)
	if _egg != nil {
		return _e.Wrap(_egg, "\u0073\u0065\u0074\u0041\u006c\u006c", "")
	}
	return nil
}
func (_ebbf *Boxes) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_bga *Boxes, _afa error) {
	const _gcae = "\u0042o\u0078e\u0073\u002e\u0053\u0065\u006ce\u0063\u0074B\u0079\u0053\u0069\u007a\u0065"
	if _ebbf == nil {
		return nil, _e.Error(_gcae, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_ebbf) == 0 {
		return _ebbf, nil
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _e.Errorf(_gcae, "\u0069\u006e\u0076al\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE:
	default:
		return nil, _e.Errorf(_gcae, "i\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020t\u0079\u0070\u0065:\u0020'\u0025\u0064\u0027", tp)
	}
	_cba := _ebbf.makeSizeIndicator(width, height, tp, relation)
	_gfb, _afa := _ebbf.selectWithIndicator(_cba)
	if _afa != nil {
		return nil, _e.Wrap(_afa, _gcae, "")
	}
	return _gfb, nil
}
func (_agfc *Bitmaps) String() string {
	_bfce := _c.Builder{}
	for _, _bdaa := range _agfc.Values {
		_bfce.WriteString(_bdaa.String())
		_bfce.WriteRune('\n')
	}
	return _bfce.String()
}
func (_cge *Bitmap) SetByte(index int, v byte) error {
	if index > len(_cge.Data)-1 || index < 0 {
		return _e.Errorf("\u0053e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", index)
	}
	_cge.Data[index] = v
	return nil
}

var (
	_cbff *Bitmap
	_gabd *Bitmap
)

func _aefd(_eac, _fffc int) int {
	if _eac > _fffc {
		return _eac
	}
	return _fffc
}
func ClipBoxToRectangle(box *_aa.Rectangle, wi, hi int) (_eada *_aa.Rectangle, _eabf error) {
	const _fee = "\u0043l\u0069p\u0042\u006f\u0078\u0054\u006fR\u0065\u0063t\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, _e.Error(_fee, "\u0027\u0062\u006f\u0078\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if box.Min.X >= wi || box.Min.Y >= hi || box.Max.X <= 0 || box.Max.Y <= 0 {
		return nil, _e.Error(_fee, "\u0027\u0062\u006fx'\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065")
	}
	_aadc := *box
	_eada = &_aadc
	if _eada.Min.X < 0 {
		_eada.Max.X += _eada.Min.X
		_eada.Min.X = 0
	}
	if _eada.Min.Y < 0 {
		_eada.Max.Y += _eada.Min.Y
		_eada.Min.Y = 0
	}
	if _eada.Max.X > wi {
		_eada.Max.X = wi
	}
	if _eada.Max.Y > hi {
		_eada.Max.Y = hi
	}
	return _eada, nil
}

const (
	ComponentConn Component = iota
	ComponentCharacters
	ComponentWords
)

func _gegg(_dadg *Bitmap, _gagg, _ebe, _cgdb, _eda int, _fcdcd RasterOperator, _aagca *Bitmap, _abcb, _cacd int) error {
	const _ggbfb = "\u0072a\u0073t\u0065\u0072\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e"
	if _dadg == nil {
		return _e.Error(_ggbfb, "\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _fcdcd == PixDst {
		return nil
	}
	switch _fcdcd {
	case PixClr, PixSet, PixNotDst:
		_edcf(_dadg, _gagg, _ebe, _cgdb, _eda, _fcdcd)
		return nil
	}
	if _aagca == nil {
		_ca.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _e.Error(_ggbfb, "\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _gfcae := _fgedg(_dadg, _gagg, _ebe, _cgdb, _eda, _fcdcd, _aagca, _abcb, _cacd); _gfcae != nil {
		return _e.Wrap(_gfcae, _ggbfb, "")
	}
	return nil
}
func (_bebe *Bitmap) String() string {
	var _geac = "\u000a"
	for _feg := 0; _feg < _bebe.Height; _feg++ {
		var _dfe string
		for _cae := 0; _cae < _bebe.Width; _cae++ {
			_aec := _bebe.GetPixel(_cae, _feg)
			if _aec {
				_dfe += "\u0031"
			} else {
				_dfe += "\u0030"
			}
		}
		_geac += _dfe + "\u000a"
	}
	return _geac
}
func (_dfbd *Bitmaps) GroupByWidth() (*BitmapsArray, error) {
	const _bfffa = "\u0047\u0072\u006fu\u0070\u0042\u0079\u0057\u0069\u0064\u0074\u0068"
	if len(_dfbd.Values) == 0 {
		return nil, _e.Error(_bfffa, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_fcdbc := &BitmapsArray{}
	_dfbd.SortByWidth()
	_aeec := -1
	_ddbc := -1
	for _gcagg := 0; _gcagg < len(_dfbd.Values); _gcagg++ {
		_bcbg := _dfbd.Values[_gcagg].Width
		if _bcbg > _aeec {
			_aeec = _bcbg
			_ddbc++
			_fcdbc.Values = append(_fcdbc.Values, &Bitmaps{})
		}
		_fcdbc.Values[_ddbc].AddBitmap(_dfbd.Values[_gcagg])
	}
	return _fcdbc, nil
}
func (_bcaa *Bitmap) GetComponents(components Component, maxWidth, maxHeight int) (_gfdf *Bitmaps, _cbfbg *Boxes, _gfge error) {
	const _accb = "B\u0069t\u006d\u0061\u0070\u002e\u0047\u0065\u0074\u0043o\u006d\u0070\u006f\u006een\u0074\u0073"
	if _bcaa == nil {
		return nil, nil, _e.Error(_accb, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0042\u0069\u0074\u006da\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064\u002e")
	}
	switch components {
	case ComponentConn, ComponentCharacters, ComponentWords:
	default:
		return nil, nil, _e.Error(_accb, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074s\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072")
	}
	if _bcaa.Zero() {
		_cbfbg = &Boxes{}
		_gfdf = &Bitmaps{}
		return _gfdf, _cbfbg, nil
	}
	switch components {
	case ComponentConn:
		_gfdf = &Bitmaps{}
		if _cbfbg, _gfge = _bcaa.ConnComponents(_gfdf, 8); _gfge != nil {
			return nil, nil, _e.Wrap(_gfge, _accb, "\u006e\u006f \u0070\u0072\u0065p\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentCharacters:
		_fgaa, _bgda := MorphSequence(_bcaa, MorphProcess{Operation: MopClosing, Arguments: []int{1, 6}})
		if _bgda != nil {
			return nil, nil, _e.Wrap(_bgda, _accb, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _ca.Log.IsLogLevel(_ca.LogLevelTrace) {
			_ca.Log.Trace("\u0043o\u006d\u0070o\u006e\u0065\u006e\u0074C\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074ma\u0070\u0020\u0061f\u0074\u0065r\u0020\u0063\u006c\u006f\u0073\u0069n\u0067\u003a \u0025\u0073", _fgaa.String())
		}
		_dacd := &Bitmaps{}
		_cbfbg, _bgda = _fgaa.ConnComponents(_dacd, 8)
		if _bgda != nil {
			return nil, nil, _e.Wrap(_bgda, _accb, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _ca.Log.IsLogLevel(_ca.LogLevelTrace) {
			_ca.Log.Trace("\u0043\u006f\u006d\u0070\u006f\u006ee\u006e\u0074\u0043\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0020a\u0066\u0074\u0065\u0072\u0020\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076i\u0074y\u003a\u0020\u0025\u0073", _dacd.String())
		}
		if _gfdf, _bgda = _dacd.ClipToBitmap(_bcaa); _bgda != nil {
			return nil, nil, _e.Wrap(_bgda, _accb, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentWords:
		_efgb := 1
		var _efbf *Bitmap
		switch {
		case _bcaa.XResolution <= 200:
			_efbf = _bcaa
		case _bcaa.XResolution <= 400:
			_efgb = 2
			_efbf, _gfge = _fbb(_bcaa, 1, 0, 0, 0)
			if _gfge != nil {
				return nil, nil, _e.Wrap(_gfge, _accb, "w\u006f\u0072\u0064\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0020\u002d \u0078\u0072\u0065s\u003c=\u0034\u0030\u0030")
			}
		default:
			_efgb = 4
			_efbf, _gfge = _fbb(_bcaa, 1, 1, 0, 0)
			if _gfge != nil {
				return nil, nil, _e.Wrap(_gfge, _accb, "\u0077\u006f\u0072\u0064 \u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073 \u002d \u0078\u0072\u0065\u0073\u0020\u003e\u00204\u0030\u0030")
			}
		}
		_cea, _, _bbc := _afgc(_efbf)
		if _bbc != nil {
			return nil, nil, _e.Wrap(_bbc, _accb, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_ddf, _bbc := _cbfg(_cea, _efgb)
		if _bbc != nil {
			return nil, nil, _e.Wrap(_bbc, _accb, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_gfca := &Bitmaps{}
		if _cbfbg, _bbc = _ddf.ConnComponents(_gfca, 4); _bbc != nil {
			return nil, nil, _e.Wrap(_bbc, _accb, "\u0077\u006f\u0072\u0064\u0020\u0070r\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u002c\u0020\u0063\u006f\u006en\u0065\u0063\u0074\u0020\u0065\u0078\u0070a\u006e\u0064\u0065\u0064")
		}
		if _gfdf, _bbc = _gfca.ClipToBitmap(_bcaa); _bbc != nil {
			return nil, nil, _e.Wrap(_bbc, _accb, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
	}
	_gfdf, _gfge = _gfdf.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _gfge != nil {
		return nil, nil, _e.Wrap(_gfge, _accb, "")
	}
	_cbfbg, _gfge = _cbfbg.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _gfge != nil {
		return nil, nil, _e.Wrap(_gfge, _accb, "")
	}
	return _gfdf, _cbfbg, nil
}
func _bbfda(_fcfd *Bitmap, _gfffa, _aaab int, _ddacd, _gcfd int, _fgag RasterOperator) {
	var (
		_bgbf   bool
		_dfbebe bool
		_cgfe   int
		_eefb   int
		_dbgd   int
		_acgf   int
		_ebabg  bool
		_cbgg   byte
	)
	_cfebe := 8 - (_gfffa & 7)
	_fbba := _fbdg[_cfebe]
	_gcaf := _fcfd.RowStride*_aaab + (_gfffa >> 3)
	if _ddacd < _cfebe {
		_bgbf = true
		_fbba &= _acbba[8-_cfebe+_ddacd]
	}
	if !_bgbf {
		_cgfe = (_ddacd - _cfebe) >> 3
		if _cgfe != 0 {
			_dfbebe = true
			_eefb = _gcaf + 1
		}
	}
	_dbgd = (_gfffa + _ddacd) & 7
	if !(_bgbf || _dbgd == 0) {
		_ebabg = true
		_cbgg = _acbba[_dbgd]
		_acgf = _gcaf + 1 + _cgfe
	}
	var _cdce, _baff int
	switch _fgag {
	case PixClr:
		for _cdce = 0; _cdce < _gcfd; _cdce++ {
			_fcfd.Data[_gcaf] = _bbgf(_fcfd.Data[_gcaf], 0x0, _fbba)
			_gcaf += _fcfd.RowStride
		}
		if _dfbebe {
			for _cdce = 0; _cdce < _gcfd; _cdce++ {
				for _baff = 0; _baff < _cgfe; _baff++ {
					_fcfd.Data[_eefb+_baff] = 0x0
				}
				_eefb += _fcfd.RowStride
			}
		}
		if _ebabg {
			for _cdce = 0; _cdce < _gcfd; _cdce++ {
				_fcfd.Data[_acgf] = _bbgf(_fcfd.Data[_acgf], 0x0, _cbgg)
				_acgf += _fcfd.RowStride
			}
		}
	case PixSet:
		for _cdce = 0; _cdce < _gcfd; _cdce++ {
			_fcfd.Data[_gcaf] = _bbgf(_fcfd.Data[_gcaf], 0xff, _fbba)
			_gcaf += _fcfd.RowStride
		}
		if _dfbebe {
			for _cdce = 0; _cdce < _gcfd; _cdce++ {
				for _baff = 0; _baff < _cgfe; _baff++ {
					_fcfd.Data[_eefb+_baff] = 0xff
				}
				_eefb += _fcfd.RowStride
			}
		}
		if _ebabg {
			for _cdce = 0; _cdce < _gcfd; _cdce++ {
				_fcfd.Data[_acgf] = _bbgf(_fcfd.Data[_acgf], 0xff, _cbgg)
				_acgf += _fcfd.RowStride
			}
		}
	case PixNotDst:
		for _cdce = 0; _cdce < _gcfd; _cdce++ {
			_fcfd.Data[_gcaf] = _bbgf(_fcfd.Data[_gcaf], ^_fcfd.Data[_gcaf], _fbba)
			_gcaf += _fcfd.RowStride
		}
		if _dfbebe {
			for _cdce = 0; _cdce < _gcfd; _cdce++ {
				for _baff = 0; _baff < _cgfe; _baff++ {
					_fcfd.Data[_eefb+_baff] = ^(_fcfd.Data[_eefb+_baff])
				}
				_eefb += _fcfd.RowStride
			}
		}
		if _ebabg {
			for _cdce = 0; _cdce < _gcfd; _cdce++ {
				_fcfd.Data[_acgf] = _bbgf(_fcfd.Data[_acgf], ^_fcfd.Data[_acgf], _cbgg)
				_acgf += _fcfd.RowStride
			}
		}
	}
}
func _bac(_bdeb, _bebb *Bitmap, _fcgb, _gdda, _cfbd, _facb, _dac, _gga, _bba, _bffc int, _bbde CombinationOperator) error {
	var _dfeg int
	_cbfb := func() { _dfeg++; _cfbd += _bebb.RowStride; _facb += _bdeb.RowStride; _dac += _bdeb.RowStride }
	for _dfeg = _fcgb; _dfeg < _gdda; _cbfb() {
		var _eeaf uint16
		_gbgf := _cfbd
		for _bgg := _facb; _bgg <= _dac; _bgg++ {
			_debc, _adab := _bebb.GetByte(_gbgf)
			if _adab != nil {
				return _adab
			}
			_afd, _adab := _bdeb.GetByte(_bgg)
			if _adab != nil {
				return _adab
			}
			_eeaf = (_eeaf | uint16(_afd)) << uint(_bffc)
			_afd = byte(_eeaf >> 8)
			if _bgg == _dac {
				_afd = _fcgbg(uint(_gga), _afd)
			}
			if _adab = _bebb.SetByte(_gbgf, _eeabd(_debc, _afd, _bbde)); _adab != nil {
				return _adab
			}
			_gbgf++
			_eeaf <<= uint(_bba)
		}
	}
	return nil
}
func _egdf() []int {
	_bcce := make([]int, 256)
	_bcce[0] = 0
	_bcce[1] = 7
	var _afad int
	for _afad = 2; _afad < 4; _afad++ {
		_bcce[_afad] = _bcce[_afad-2] + 6
	}
	for _afad = 4; _afad < 8; _afad++ {
		_bcce[_afad] = _bcce[_afad-4] + 5
	}
	for _afad = 8; _afad < 16; _afad++ {
		_bcce[_afad] = _bcce[_afad-8] + 4
	}
	for _afad = 16; _afad < 32; _afad++ {
		_bcce[_afad] = _bcce[_afad-16] + 3
	}
	for _afad = 32; _afad < 64; _afad++ {
		_bcce[_afad] = _bcce[_afad-32] + 2
	}
	for _afad = 64; _afad < 128; _afad++ {
		_bcce[_afad] = _bcce[_afad-64] + 1
	}
	for _afad = 128; _afad < 256; _afad++ {
		_bcce[_afad] = _bcce[_afad-128]
	}
	return _bcce
}
func NewWithData(width, height int, data []byte) (*Bitmap, error) {
	const _age = "N\u0065\u0077\u0057\u0069\u0074\u0068\u0044\u0061\u0074\u0061"
	_gdc := _eced(width, height)
	_gdc.Data = data
	if len(data) < height*_gdc.RowStride {
		return nil, _e.Errorf(_age, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u003a \u0025\u0064\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u003a\u0020\u0025\u0064", len(data), height*_gdc.RowStride)
	}
	return _gdc, nil
}
func _adgf(_dbf *Bitmap, _gab, _ge int) (*Bitmap, error) {
	const _cg = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0052\u0065p\u006c\u0069\u0063\u0061\u0074\u0065"
	if _dbf == nil {
		return nil, _e.Error(_cg, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _gab <= 0 || _ge <= 0 {
		return nil, _e.Error(_cg, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _gab == _ge {
		if _gab == 1 {
			_bef, _dfa := _cdcf(nil, _dbf)
			if _dfa != nil {
				return nil, _e.Wrap(_dfa, _cg, "\u0078\u0046\u0061\u0063\u0074\u0020\u003d\u003d\u0020y\u0046\u0061\u0063\u0074")
			}
			return _bef, nil
		}
		if _gab == 2 || _gab == 4 || _gab == 8 {
			_ebb, _bcb := _eec(_dbf, _gab)
			if _bcb != nil {
				return nil, _e.Wrap(_bcb, _cg, "\u0078\u0046a\u0063\u0074\u0020i\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
			}
			return _ebb, nil
		}
	}
	_ed := _gab * _dbf.Width
	_aaac := _ge * _dbf.Height
	_cag := New(_ed, _aaac)
	_faa := _cag.RowStride
	var (
		_ecg, _cegg, _cef, _af, _caag int
		_gb                           byte
		_fda                          error
	)
	for _cegg = 0; _cegg < _dbf.Height; _cegg++ {
		_ecg = _ge * _cegg * _faa
		for _cef = 0; _cef < _dbf.Width; _cef++ {
			if _efa := _dbf.GetPixel(_cef, _cegg); _efa {
				_caag = _gab * _cef
				for _af = 0; _af < _gab; _af++ {
					_cag.setBit(_ecg*8 + _caag + _af)
				}
			}
		}
		for _af = 1; _af < _ge; _af++ {
			_cd := _ecg + _af*_faa
			for _gcd := 0; _gcd < _faa; _gcd++ {
				if _gb, _fda = _cag.GetByte(_ecg + _gcd); _fda != nil {
					return nil, _e.Wrapf(_fda, _cg, "\u0072\u0065\u0070\u006cic\u0061\u0074\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0027\u0025d\u0027", _af)
				}
				if _fda = _cag.SetByte(_cd+_gcd, _gb); _fda != nil {
					return nil, _e.Wrap(_fda, _cg, "\u0053\u0065\u0074\u0074in\u0067\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
				}
			}
		}
	}
	return _cag, nil
}
func _fafg(_aeg, _eeb *Bitmap, _adbc, _egce, _bede uint, _cdbg, _eaaab int, _gee bool, _dfbe, _aged int) error {
	for _gfcg := _cdbg; _gfcg < _eaaab; _gfcg++ {
		if _dfbe+1 < len(_aeg.Data) {
			_aegb := _gfcg+1 == _eaaab
			_bfdc, _ggbd := _aeg.GetByte(_dfbe)
			if _ggbd != nil {
				return _ggbd
			}
			_dfbe++
			_bfdc <<= _adbc
			_ffgb, _ggbd := _aeg.GetByte(_dfbe)
			if _ggbd != nil {
				return _ggbd
			}
			_ffgb >>= _egce
			_dcde := _bfdc | _ffgb
			if _aegb && !_gee {
				_dcde = _fcgbg(_bede, _dcde)
			}
			_ggbd = _eeb.SetByte(_aged, _dcde)
			if _ggbd != nil {
				return _ggbd
			}
			_aged++
			if _aegb && _gee {
				_fadb, _fcc := _aeg.GetByte(_dfbe)
				if _fcc != nil {
					return _fcc
				}
				_fadb <<= _adbc
				_dcde = _fcgbg(_bede, _fadb)
				if _fcc = _eeb.SetByte(_aged, _dcde); _fcc != nil {
					return _fcc
				}
			}
			continue
		}
		_bgdg, _aefb := _aeg.GetByte(_dfbe)
		if _aefb != nil {
			_ca.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0061\u0074\u003a\u0020%\u0064\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0073", _dfbe, _aefb)
			return _aefb
		}
		_bgdg <<= _adbc
		_dfbe++
		_aefb = _eeb.SetByte(_aged, _bgdg)
		if _aefb != nil {
			return _aefb
		}
		_aged++
	}
	return nil
}
func _bbgf(_eadcg, _defce, _cgce byte) byte { return (_eadcg &^ (_cgce)) | (_defce & _cgce) }
func _fafaf(_dggge *Bitmap, _bab *Bitmap, _bcee int) (_gcce error) {
	const _ebgf = "\u0073\u0065\u0065\u0064\u0066\u0069\u006c\u006c\u0042\u0069\u006e\u0061r\u0079\u004c\u006f\u0077"
	_caaga := _gcg(_dggge.Height, _bab.Height)
	_fbdf := _gcg(_dggge.RowStride, _bab.RowStride)
	switch _bcee {
	case 4:
		_gcce = _bgge(_dggge, _bab, _caaga, _fbdf)
	case 8:
		_gcce = _ggee(_dggge, _bab, _caaga, _fbdf)
	default:
		return _e.Errorf(_ebgf, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0034\u0020\u006fr\u0020\u0038\u0020\u002d\u0020i\u0073\u003a \u0027\u0025\u0064\u0027", _bcee)
	}
	if _gcce != nil {
		return _e.Wrap(_gcce, _ebgf, "")
	}
	return nil
}
func (_fecbf *BitmapsArray) GetBitmaps(i int) (*Bitmaps, error) {
	const _ebbfd = "\u0042\u0069\u0074ma\u0070\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u0069\u0074\u006d\u0061\u0070\u0073"
	if _fecbf == nil {
		return nil, _e.Error(_ebbfd, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_fecbf.Values)-1 {
		return nil, _e.Errorf(_ebbfd, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _fecbf.Values[i], nil
}
func (_dfdg *BitmapsArray) AddBox(box *_aa.Rectangle) { _dfdg.Boxes = append(_dfdg.Boxes, box) }
func (_abgc *Bitmap) Equals(s *Bitmap) bool {
	if len(_abgc.Data) != len(s.Data) || _abgc.Width != s.Width || _abgc.Height != s.Height {
		return false
	}
	for _efbg := 0; _efbg < _abgc.Height; _efbg++ {
		_dfb := _efbg * _abgc.RowStride
		for _agg := 0; _agg < _abgc.RowStride; _agg++ {
			if _abgc.Data[_dfb+_agg] != s.Data[_dfb+_agg] {
				return false
			}
		}
	}
	return true
}
func TstOSymbol(t *_f.T, scale ...int) *Bitmap {
	_bgabb, _ffgg := NewWithData(4, 5, []byte{0xF0, 0x90, 0x90, 0x90, 0xF0})
	_d.NoError(t, _ffgg)
	return TstGetScaledSymbol(t, _bgabb, scale...)
}
func (_gdcee *ClassedPoints) SortByY()       { _gdcee._acab = _gdcee.ySortFunction(); _ff.Sort(_gdcee) }
func (_begf *Bitmap) GetBitOffset(x int) int { return x & 0x07 }
func (_caab Points) GetIntX(i int) (int, error) {
	if i >= len(_caab) {
		return 0, _e.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0058", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_caab[i].X), nil
}
func (_ggfg *Bitmap) resizeImageData(_gfde *Bitmap) error {
	if _gfde == nil {
		return _e.Error("\u0072e\u0073i\u007a\u0065\u0049\u006d\u0061\u0067\u0065\u0044\u0061\u0074\u0061", "\u0073r\u0063 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ggfg.SizesEqual(_gfde) {
		return nil
	}
	_ggfg.Data = make([]byte, len(_gfde.Data))
	_ggfg.Width = _gfde.Width
	_ggfg.Height = _gfde.Height
	_ggfg.RowStride = _gfde.RowStride
	return nil
}
func TstImageBitmapData() []byte { return _gabd.Data }
func RankHausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH, area1, area3 int, rank float32, tab8 []int) (_dgdc bool, _ecdg error) {
	const _bcea = "\u0052\u0061\u006ek\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_aega, _dfbf := p1.Width, p1.Height
	_geff, _cefg := p3.Width, p3.Height
	if _ac.Abs(_aega-_geff) > maxDiffW {
		return false, nil
	}
	if _ac.Abs(_dfbf-_cefg) > maxDiffH {
		return false, nil
	}
	_fegc := int(float32(area1)*(1.0-rank) + 0.5)
	_cafd := int(float32(area3)*(1.0-rank) + 0.5)
	var _ddfc, _ddce int
	if delX >= 0 {
		_ddfc = int(delX + 0.5)
	} else {
		_ddfc = int(delX - 0.5)
	}
	if delY >= 0 {
		_ddce = int(delY + 0.5)
	} else {
		_ddce = int(delY - 0.5)
	}
	_egaa := p1.CreateTemplate()
	if _ecdg = _egaa.RasterOperation(0, 0, _aega, _dfbf, PixSrc, p1, 0, 0); _ecdg != nil {
		return false, _e.Wrap(_ecdg, _bcea, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _ecdg = _egaa.RasterOperation(_ddfc, _ddce, _aega, _dfbf, PixNotSrcAndDst, p4, 0, 0); _ecdg != nil {
		return false, _e.Wrap(_ecdg, _bcea, "\u0074 \u0026\u0020\u0021\u0070\u0034")
	}
	_dgdc, _ecdg = _egaa.ThresholdPixelSum(_fegc, tab8)
	if _ecdg != nil {
		return false, _e.Wrap(_ecdg, _bcea, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0031")
	}
	if _dgdc {
		return false, nil
	}
	if _ecdg = _egaa.RasterOperation(_ddfc, _ddce, _geff, _cefg, PixSrc, p3, 0, 0); _ecdg != nil {
		return false, _e.Wrap(_ecdg, _bcea, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _ecdg = _egaa.RasterOperation(0, 0, _geff, _cefg, PixNotSrcAndDst, p2, 0, 0); _ecdg != nil {
		return false, _e.Wrap(_ecdg, _bcea, "\u0074 \u0026\u0020\u0021\u0070\u0032")
	}
	_dgdc, _ecdg = _egaa.ThresholdPixelSum(_cafd, tab8)
	if _ecdg != nil {
		return false, _e.Wrap(_ecdg, _bcea, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0033")
	}
	return !_dgdc, nil
}
func _fbb(_edc *Bitmap, _dbb ...int) (_afc *Bitmap, _fbce error) {
	const _ecgg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0043\u0061\u0073\u0063\u0061\u0064\u0065"
	if _edc == nil {
		return nil, _e.Error(_ecgg, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_dbb) == 0 || len(_dbb) > 4 {
		return nil, _e.Error(_ecgg, "t\u0068\u0065\u0072\u0065\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0074\u0020\u006cea\u0073\u0074\u0020\u006fn\u0065\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020mo\u0073\u0074 \u0034\u0020\u006c\u0065\u0076\u0065\u006c\u0073")
	}
	if _dbb[0] <= 0 {
		_ca.Log.Debug("\u006c\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030 \u002d\u0020\u006e\u006f\u0020\u0072\u0065\u0064\u0075\u0063t\u0069\u006f\u006e")
		_afc, _fbce = _cdcf(nil, _edc)
		if _fbce != nil {
			return nil, _e.Wrap(_fbce, _ecgg, "l\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030")
		}
		return _afc, nil
	}
	_dd := _bcc()
	_afc = _edc
	for _cce, _fff := range _dbb {
		if _fff <= 0 {
			break
		}
		_afc, _fbce = _dcdc(_afc, _fff, _dd)
		if _fbce != nil {
			return nil, _e.Wrapf(_fbce, _ecgg, "\u006c\u0065\u0076\u0065\u006c\u0025\u0064\u0020\u0072\u0065\u0064\u0075c\u0074\u0069\u006f\u006e", _cce)
		}
	}
	return _afc, nil
}

type CombinationOperator int

func _afgc(_cafb *Bitmap) (_eddda *Bitmap, _bfaf int, _daee error) {
	const _cecc = "\u0042i\u0074\u006d\u0061\u0070.\u0077\u006f\u0072\u0064\u004da\u0073k\u0042y\u0044\u0069\u006c\u0061\u0074\u0069\u006fn"
	if _cafb == nil {
		return nil, 0, _e.Errorf(_cecc, "\u0027\u0073\u0027\u0020bi\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	var _aea, _abdf *Bitmap
	if _aea, _daee = _cdcf(nil, _cafb); _daee != nil {
		return nil, 0, _e.Wrap(_daee, _cecc, "\u0063\u006f\u0070\u0079\u0020\u0027\u0073\u0027")
	}
	var (
		_gfce        [13]int
		_fgfe, _bffb int
	)
	_cbeec := 12
	_ggbf := _ac.NewNumSlice(_cbeec + 1)
	_gbfd := _ac.NewNumSlice(_cbeec + 1)
	var _bfcd *Boxes
	for _gacc := 0; _gacc <= _cbeec; _gacc++ {
		if _gacc == 0 {
			if _abdf, _daee = _cdcf(nil, _aea); _daee != nil {
				return nil, 0, _e.Wrap(_daee, _cecc, "\u0066i\u0072\u0073\u0074\u0020\u0062\u006d2")
			}
		} else {
			if _abdf, _daee = _ecbgb(_aea, MorphProcess{Operation: MopDilation, Arguments: []int{2, 1}}); _daee != nil {
				return nil, 0, _e.Wrap(_daee, _cecc, "\u0064\u0069\u006ca\u0074\u0069\u006f\u006e\u0020\u0062\u006d\u0032")
			}
		}
		if _bfcd, _daee = _abdf.connComponentsBB(4); _daee != nil {
			return nil, 0, _e.Wrap(_daee, _cecc, "")
		}
		_gfce[_gacc] = len(*_bfcd)
		_ggbf.AddInt(_gfce[_gacc])
		switch _gacc {
		case 0:
			_fgfe = _gfce[0]
		default:
			_bffb = _gfce[_gacc-1] - _gfce[_gacc]
			_gbfd.AddInt(_bffb)
		}
		_aea = _abdf
	}
	_fdb := true
	_egdg := 2
	var _ebbc, _dgda int
	for _aca := 1; _aca < len(*_gbfd); _aca++ {
		if _ebbc, _daee = _ggbf.GetInt(_aca); _daee != nil {
			return nil, 0, _e.Wrap(_daee, _cecc, "\u0043\u0068\u0065\u0063ki\u006e\u0067\u0020\u0062\u0065\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0069o\u006e")
		}
		if _fdb && _ebbc < int(0.3*float32(_fgfe)) {
			_egdg = _aca + 1
			_fdb = false
		}
		if _bffb, _daee = _gbfd.GetInt(_aca); _daee != nil {
			return nil, 0, _e.Wrap(_daee, _cecc, "\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006ea\u0044\u0069\u0066\u0066")
		}
		if _bffb > _dgda {
			_dgda = _bffb
		}
	}
	_aaacd := _cafb.XResolution
	if _aaacd == 0 {
		_aaacd = 150
	}
	if _aaacd > 110 {
		_egdg++
	}
	if _egdg < 2 {
		_ca.Log.Trace("J\u0042\u0049\u0047\u0032\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0042\u0065\u0073\u0074 \u0074\u006f\u0020\u006d\u0069\u006e\u0069\u006d\u0075\u006d a\u006c\u006c\u006fw\u0061b\u006c\u0065")
		_egdg = 2
	}
	_bfaf = _egdg + 1
	if _eddda, _daee = _cdfd(nil, _cafb, _egdg+1, 1); _daee != nil {
		return nil, 0, _e.Wrap(_daee, _cecc, "\u0067\u0065\u0074\u0074in\u0067\u0020\u006d\u0061\u0073\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	return _eddda, _bfaf, nil
}
func NewClassedPoints(points *Points, classes _ac.IntSlice) (*ClassedPoints, error) {
	const _cbeed = "\u004e\u0065w\u0043\u006c\u0061s\u0073\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073"
	if points == nil {
		return nil, _e.Error(_cbeed, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if classes == nil {
		return nil, _e.Error(_cbeed, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0063\u006c\u0061ss\u0065\u0073")
	}
	_fece := &ClassedPoints{Points: points, IntSlice: classes}
	if _deecg := _fece.validateIntSlice(); _deecg != nil {
		return nil, _e.Wrap(_deecg, _cbeed, "")
	}
	return _fece, nil
}
func _gedd(_daea, _bedbd *Bitmap, _gfbgg *Selection) (*Bitmap, error) {
	const _egbe = "c\u006c\u006f\u0073\u0065\u0042\u0069\u0074\u006d\u0061\u0070"
	var _gddd error
	if _daea, _gddd = _gebfg(_daea, _bedbd, _gfbgg); _gddd != nil {
		return nil, _gddd
	}
	_cgaa, _gddd := _caagg(nil, _bedbd, _gfbgg)
	if _gddd != nil {
		return nil, _e.Wrap(_gddd, _egbe, "")
	}
	if _, _gddd = _cfgc(_daea, _cgaa, _gfbgg); _gddd != nil {
		return nil, _e.Wrap(_gddd, _egbe, "")
	}
	return _daea, nil
}
func _dfbeb() []int {
	_gdab := make([]int, 256)
	for _eebg := 0; _eebg <= 0xff; _eebg++ {
		_bge := byte(_eebg)
		_gdab[_bge] = int(_bge&0x1) + (int(_bge>>1) & 0x1) + (int(_bge>>2) & 0x1) + (int(_bge>>3) & 0x1) + (int(_bge>>4) & 0x1) + (int(_bge>>5) & 0x1) + (int(_bge>>6) & 0x1) + (int(_bge>>7) & 0x1)
	}
	return _gdab
}
func (_adgg *Bitmap) GetUnpaddedData() ([]byte, error) {
	_dgg := uint(_adgg.Width & 0x07)
	if _dgg == 0 {
		return _adgg.Data, nil
	}
	_cde := _adgg.Width * _adgg.Height
	if _cde%8 != 0 {
		_cde >>= 3
		_cde++
	} else {
		_cde >>= 3
	}
	_bdf := make([]byte, _cde)
	_fbeb := _ce.NewWriterMSB(_bdf)
	const _afe = "\u0047e\u0074U\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	for _dfd := 0; _dfd < _adgg.Height; _dfd++ {
		for _ggdd := 0; _ggdd < _adgg.RowStride; _ggdd++ {
			_egc := _adgg.Data[_dfd*_adgg.RowStride+_ggdd]
			if _ggdd != _adgg.RowStride-1 {
				_fgdg := _fbeb.WriteByte(_egc)
				if _fgdg != nil {
					return nil, _e.Wrap(_fgdg, _afe, "")
				}
				continue
			}
			for _aed := uint(0); _aed < _dgg; _aed++ {
				_bce := _fbeb.WriteBit(int(_egc >> (7 - _aed) & 0x01))
				if _bce != nil {
					return nil, _e.Wrap(_bce, _afe, "")
				}
			}
		}
	}
	return _bdf, nil
}

type SizeComparison int

func _fcd(_cdc, _gfd *Bitmap, _fbd int, _cfd []byte, _aafg int) (_ebc error) {
	const _cgg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0034"
	var (
		_fea, _dgf, _adea, _cfg, _fcf, _cegb, _gcc, _badc int
		_ggg, _cgdc                                       uint32
		_bcbb, _ecd                                       byte
		_ddd                                              uint16
	)
	_bd := make([]byte, 4)
	_bcd := make([]byte, 4)
	for _adea = 0; _adea < _cdc.Height-1; _adea, _cfg = _adea+2, _cfg+1 {
		_fea = _adea * _cdc.RowStride
		_dgf = _cfg * _gfd.RowStride
		for _fcf, _cegb = 0, 0; _fcf < _aafg; _fcf, _cegb = _fcf+4, _cegb+1 {
			for _gcc = 0; _gcc < 4; _gcc++ {
				_badc = _fea + _fcf + _gcc
				if _badc <= len(_cdc.Data)-1 && _badc < _fea+_cdc.RowStride {
					_bd[_gcc] = _cdc.Data[_badc]
				} else {
					_bd[_gcc] = 0x00
				}
				_badc = _fea + _cdc.RowStride + _fcf + _gcc
				if _badc <= len(_cdc.Data)-1 && _badc < _fea+(2*_cdc.RowStride) {
					_bcd[_gcc] = _cdc.Data[_badc]
				} else {
					_bcd[_gcc] = 0x00
				}
			}
			_ggg = _db.BigEndian.Uint32(_bd)
			_cgdc = _db.BigEndian.Uint32(_bcd)
			_cgdc &= _ggg
			_cgdc &= _cgdc << 1
			_cgdc &= 0xaaaaaaaa
			_ggg = _cgdc | (_cgdc << 7)
			_bcbb = byte(_ggg >> 24)
			_ecd = byte((_ggg >> 8) & 0xff)
			_badc = _dgf + _cegb
			if _badc+1 == len(_gfd.Data)-1 || _badc+1 >= _dgf+_gfd.RowStride {
				_gfd.Data[_badc] = _cfd[_bcbb]
				if _ebc = _gfd.SetByte(_badc, _cfd[_bcbb]); _ebc != nil {
					return _e.Wrapf(_ebc, _cgg, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _badc)
				}
			} else {
				_ddd = (uint16(_cfd[_bcbb]) << 8) | uint16(_cfd[_ecd])
				if _ebc = _gfd.setTwoBytes(_badc, _ddd); _ebc != nil {
					return _e.Wrapf(_ebc, _cgg, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _badc)
				}
				_cegb++
			}
		}
	}
	return nil
}
func (_bbce *Bitmaps) AddBitmap(bm *Bitmap) { _bbce.Values = append(_bbce.Values, bm) }

type Points []Point

func (_dcce *ClassedPoints) YAtIndex(i int) float32 { return (*_dcce.Points)[_dcce.IntSlice[i]].Y }
func (_dfgd *Bitmap) addBorderGeneral(_bafc, _cdea, _faf, _bfc int, _fbeg int) (*Bitmap, error) {
	const _bcg = "\u0061\u0064d\u0042\u006f\u0072d\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _bafc < 0 || _cdea < 0 || _faf < 0 || _bfc < 0 {
		return nil, _e.Error(_bcg, "n\u0065\u0067\u0061\u0074iv\u0065 \u0062\u006f\u0072\u0064\u0065r\u0020\u0061\u0064\u0064\u0065\u0064")
	}
	_feb, _gcb := _dfgd.Width, _dfgd.Height
	_agec := _feb + _bafc + _cdea
	_cfe := _gcb + _faf + _bfc
	_bfbb := New(_agec, _cfe)
	_bfbb.Color = _dfgd.Color
	_dbd := PixClr
	if _fbeg > 0 {
		_dbd = PixSet
	}
	_cab := _bfbb.RasterOperation(0, 0, _bafc, _cfe, _dbd, nil, 0, 0)
	if _cab != nil {
		return nil, _e.Wrap(_cab, _bcg, "\u006c\u0065\u0066\u0074")
	}
	_cab = _bfbb.RasterOperation(_agec-_cdea, 0, _cdea, _cfe, _dbd, nil, 0, 0)
	if _cab != nil {
		return nil, _e.Wrap(_cab, _bcg, "\u0072\u0069\u0067h\u0074")
	}
	_cab = _bfbb.RasterOperation(0, 0, _agec, _faf, _dbd, nil, 0, 0)
	if _cab != nil {
		return nil, _e.Wrap(_cab, _bcg, "\u0074\u006f\u0070")
	}
	_cab = _bfbb.RasterOperation(0, _cfe-_bfc, _agec, _bfc, _dbd, nil, 0, 0)
	if _cab != nil {
		return nil, _e.Wrap(_cab, _bcg, "\u0062\u006f\u0074\u0074\u006f\u006d")
	}
	_cab = _bfbb.RasterOperation(_bafc, _faf, _feb, _gcb, PixSrc, _dfgd, 0, 0)
	if _cab != nil {
		return nil, _e.Wrap(_cab, _bcg, "\u0063\u006f\u0070\u0079")
	}
	return _bfbb, nil
}
func _gdbca(_fgdga *Bitmap, _bdcc, _dacc, _egge, _bbfc int, _aeag RasterOperator, _bdee *Bitmap, _gacb, _bcdb int) error {
	var (
		_agac         bool
		_dcaec        bool
		_fbfedd       byte
		_deea         int
		_fcggb        int
		_fdad         int
		_edaa         int
		_ceede        bool
		_ccd          int
		_dcbe         int
		_afece        int
		_bbbf         bool
		_adedb        byte
		_bbcg         int
		_edba         int
		_bbddd        int
		_fgab         byte
		_cafab        int
		_bcfba        int
		_deba         uint
		_agbca        uint
		_gcff         byte
		_adfd         shift
		_fafd         bool
		_cdgf         bool
		_fbde, _fcbaf int
	)
	if _gacb&7 != 0 {
		_bcfba = 8 - (_gacb & 7)
	}
	if _bdcc&7 != 0 {
		_fcggb = 8 - (_bdcc & 7)
	}
	if _bcfba == 0 && _fcggb == 0 {
		_gcff = _fbdg[0]
	} else {
		if _fcggb > _bcfba {
			_deba = uint(_fcggb - _bcfba)
		} else {
			_deba = uint(8 - (_bcfba - _fcggb))
		}
		_agbca = 8 - _deba
		_gcff = _fbdg[_deba]
	}
	if (_bdcc & 7) != 0 {
		_agac = true
		_deea = 8 - (_bdcc & 7)
		_fbfedd = _fbdg[_deea]
		_fdad = _fgdga.RowStride*_dacc + (_bdcc >> 3)
		_edaa = _bdee.RowStride*_bcdb + (_gacb >> 3)
		_cafab = 8 - (_gacb & 7)
		if _deea > _cafab {
			_adfd = _dgfe
			if _egge >= _bcfba {
				_fafd = true
			}
		} else {
			_adfd = _ecde
		}
	}
	if _egge < _deea {
		_dcaec = true
		_fbfedd &= _acbba[8-_deea+_egge]
	}
	if !_dcaec {
		_ccd = (_egge - _deea) >> 3
		if _ccd != 0 {
			_ceede = true
			_dcbe = _fgdga.RowStride*_dacc + ((_bdcc + _fcggb) >> 3)
			_afece = _bdee.RowStride*_bcdb + ((_gacb + _fcggb) >> 3)
		}
	}
	_bbcg = (_bdcc + _egge) & 7
	if !(_dcaec || _bbcg == 0) {
		_bbbf = true
		_adedb = _acbba[_bbcg]
		_edba = _fgdga.RowStride*_dacc + ((_bdcc + _fcggb) >> 3) + _ccd
		_bbddd = _bdee.RowStride*_bcdb + ((_gacb + _fcggb) >> 3) + _ccd
		if _bbcg > int(_agbca) {
			_cdgf = true
		}
	}
	switch _aeag {
	case PixSrc:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], _fgab, _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] = _fgab
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], _fgab, _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixNotSrc:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], ^_fgab, _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] = ^_fgab
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], ^_fgab, _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixSrcOrDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], _fgab|_fgdga.Data[_fdad], _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] |= _fgab
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], _fgab|_fgdga.Data[_edba], _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixSrcAndDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], _fgab&_fgdga.Data[_fdad], _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] &= _fgab
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], _fgab&_fgdga.Data[_edba], _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixSrcXorDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], _fgab^_fgdga.Data[_fdad], _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] ^= _fgab
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], _fgab^_fgdga.Data[_edba], _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixNotSrcOrDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], ^_fgab|_fgdga.Data[_fdad], _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] |= ^_fgab
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], ^_fgab|_fgdga.Data[_edba], _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixNotSrcAndDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], ^_fgab&_fgdga.Data[_fdad], _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] &= ^_fgab
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], ^_fgab&_fgdga.Data[_edba], _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixSrcOrNotDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], _fgab|^_fgdga.Data[_fdad], _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] = _fgab | ^_fgdga.Data[_dcbe+_fcbaf]
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], _fgab|^_fgdga.Data[_edba], _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixSrcAndNotDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], _fgab&^_fgdga.Data[_fdad], _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] = _fgab &^ _fgdga.Data[_dcbe+_fcbaf]
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], _fgab&^_fgdga.Data[_edba], _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], ^(_fgab | _fgdga.Data[_fdad]), _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] = ^(_fgab | _fgdga.Data[_dcbe+_fcbaf])
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], ^(_fgab | _fgdga.Data[_edba]), _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], ^(_fgab & _fgdga.Data[_fdad]), _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] = ^(_fgab & _fgdga.Data[_dcbe+_fcbaf])
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], ^(_fgab & _fgdga.Data[_edba]), _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		if _agac {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				if _adfd == _dgfe {
					_fgab = _bdee.Data[_edaa] << _deba
					if _fafd {
						_fgab = _bbgf(_fgab, _bdee.Data[_edaa+1]>>_agbca, _gcff)
					}
				} else {
					_fgab = _bdee.Data[_edaa] >> _agbca
				}
				_fgdga.Data[_fdad] = _bbgf(_fgdga.Data[_fdad], ^(_fgab ^ _fgdga.Data[_fdad]), _fbfedd)
				_fdad += _fgdga.RowStride
				_edaa += _bdee.RowStride
			}
		}
		if _ceede {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				for _fcbaf = 0; _fcbaf < _ccd; _fcbaf++ {
					_fgab = _bbgf(_bdee.Data[_afece+_fcbaf]<<_deba, _bdee.Data[_afece+_fcbaf+1]>>_agbca, _gcff)
					_fgdga.Data[_dcbe+_fcbaf] = ^(_fgab ^ _fgdga.Data[_dcbe+_fcbaf])
				}
				_dcbe += _fgdga.RowStride
				_afece += _bdee.RowStride
			}
		}
		if _bbbf {
			for _fbde = 0; _fbde < _bbfc; _fbde++ {
				_fgab = _bdee.Data[_bbddd] << _deba
				if _cdgf {
					_fgab = _bbgf(_fgab, _bdee.Data[_bbddd+1]>>_agbca, _gcff)
				}
				_fgdga.Data[_edba] = _bbgf(_fgdga.Data[_edba], ^(_fgab ^ _fgdga.Data[_edba]), _adedb)
				_edba += _fgdga.RowStride
				_bbddd += _bdee.RowStride
			}
		}
	default:
		_ca.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _aeag)
		return _e.Error("\u0072a\u0073t\u0065\u0072\u004f\u0070\u0047e\u006e\u0065r\u0061\u006c\u004c\u006f\u0077", "\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}
func (_bde *Bitmap) setEightBytes(_ffgcf int, _agff uint64) error {
	_gag := _bde.RowStride - (_ffgcf % _bde.RowStride)
	if _bde.RowStride != _bde.Width>>3 {
		_gag--
	}
	if _gag >= 8 {
		return _bde.setEightFullBytes(_ffgcf, _agff)
	}
	return _bde.setEightPartlyBytes(_ffgcf, _gag, _agff)
}
func (_dbcg *Bitmap) setPadBits(_dccf int) {
	_cbee := 8 - _dbcg.Width%8
	if _cbee == 8 {
		return
	}
	_faab := _dbcg.Width / 8
	_cbeb := _fbdg[_cbee]
	if _dccf == 0 {
		_cbeb ^= _cbeb
	}
	var _gdca int
	for _bced := 0; _bced < _dbcg.Height; _bced++ {
		_gdca = _bced*_dbcg.RowStride + _faab
		if _dccf == 0 {
			_dbcg.Data[_gdca] &= _cbeb
		} else {
			_dbcg.Data[_gdca] |= _cbeb
		}
	}
}
func (_bgadb *Bitmap) centroid(_dge, _addb []int) (Point, error) {
	_aaggeb := Point{}
	_bgadb.setPadBits(0)
	if len(_dge) == 0 {
		_dge = _egdf()
	}
	if len(_addb) == 0 {
		_addb = _dfbeb()
	}
	var _bcca, _gecb, _fgdfc, _aaec, _ggec, _fdg int
	var _fccg byte
	for _ggec = 0; _ggec < _bgadb.Height; _ggec++ {
		_cecg := _bgadb.RowStride * _ggec
		_aaec = 0
		for _fdg = 0; _fdg < _bgadb.RowStride; _fdg++ {
			_fccg = _bgadb.Data[_cecg+_fdg]
			if _fccg != 0 {
				_aaec += _addb[_fccg]
				_bcca += _dge[_fccg] + _fdg*8*_addb[_fccg]
			}
		}
		_fgdfc += _aaec
		_gecb += _aaec * _ggec
	}
	if _fgdfc != 0 {
		_aaggeb.X = float32(_bcca) / float32(_fgdfc)
		_aaggeb.Y = float32(_gecb) / float32(_fgdfc)
	}
	return _aaggeb, nil
}
func (_dgde *byHeight) Swap(i, j int) {
	_dgde.Values[i], _dgde.Values[j] = _dgde.Values[j], _dgde.Values[i]
	if _dgde.Boxes != nil {
		_dgde.Boxes[i], _dgde.Boxes[j] = _dgde.Boxes[j], _dgde.Boxes[i]
	}
}
func (_cafc *Bitmap) clearAll() error {
	return _cafc.RasterOperation(0, 0, _cafc.Width, _cafc.Height, PixClr, nil, 0, 0)
}
func (_gfab *ClassedPoints) xSortFunction() func(_fgcd int, _eegge int) bool {
	return func(_ebbcg, _bcgc int) bool { return _gfab.XAtIndex(_ebbcg) < _gfab.XAtIndex(_bcgc) }
}
func (_bedcd *Bitmap) setFourBytes(_acfa int, _ecbf uint32) error {
	if _acfa+3 > len(_bedcd.Data)-1 {
		return _e.Errorf("\u0073\u0065\u0074F\u006f\u0075\u0072\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _acfa)
	}
	_bedcd.Data[_acfa] = byte((_ecbf & 0xff000000) >> 24)
	_bedcd.Data[_acfa+1] = byte((_ecbf & 0xff0000) >> 16)
	_bedcd.Data[_acfa+2] = byte((_ecbf & 0xff00) >> 8)
	_bedcd.Data[_acfa+3] = byte(_ecbf & 0xff)
	return nil
}
func TstISymbol(t *_f.T, scale ...int) *Bitmap {
	_aeef, _gaab := NewWithData(1, 5, []byte{0x80, 0x80, 0x80, 0x80, 0x80})
	_d.NoError(t, _gaab)
	return TstGetScaledSymbol(t, _aeef, scale...)
}
func (_dgdb MorphProcess) getWidthHeight() (_fdag, _gfe int) {
	return _dgdb.Arguments[0], _dgdb.Arguments[1]
}

const (
	MopDilation MorphOperation = iota
	MopErosion
	MopOpening
	MopClosing
	MopRankBinaryReduction
	MopReplicativeBinaryExpansion
	MopAddBorder
)

func (_ged *Boxes) makeSizeIndicator(_ddge, _bafb int, _fada LocationFilter, _gfdg SizeComparison) *_ac.NumSlice {
	_bgad := &_ac.NumSlice{}
	var _fgef, _agbb, _ccca int
	for _, _cdeb := range *_ged {
		_fgef = 0
		_agbb, _ccca = _cdeb.Dx(), _cdeb.Dy()
		switch _fada {
		case LocSelectWidth:
			if (_gfdg == SizeSelectIfLT && _agbb < _ddge) || (_gfdg == SizeSelectIfGT && _agbb > _ddge) || (_gfdg == SizeSelectIfLTE && _agbb <= _ddge) || (_gfdg == SizeSelectIfGTE && _agbb >= _ddge) {
				_fgef = 1
			}
		case LocSelectHeight:
			if (_gfdg == SizeSelectIfLT && _ccca < _bafb) || (_gfdg == SizeSelectIfGT && _ccca > _bafb) || (_gfdg == SizeSelectIfLTE && _ccca <= _bafb) || (_gfdg == SizeSelectIfGTE && _ccca >= _bafb) {
				_fgef = 1
			}
		case LocSelectIfEither:
			if (_gfdg == SizeSelectIfLT && (_ccca < _bafb || _agbb < _ddge)) || (_gfdg == SizeSelectIfGT && (_ccca > _bafb || _agbb > _ddge)) || (_gfdg == SizeSelectIfLTE && (_ccca <= _bafb || _agbb <= _ddge)) || (_gfdg == SizeSelectIfGTE && (_ccca >= _bafb || _agbb >= _ddge)) {
				_fgef = 1
			}
		case LocSelectIfBoth:
			if (_gfdg == SizeSelectIfLT && (_ccca < _bafb && _agbb < _ddge)) || (_gfdg == SizeSelectIfGT && (_ccca > _bafb && _agbb > _ddge)) || (_gfdg == SizeSelectIfLTE && (_ccca <= _bafb && _agbb <= _ddge)) || (_gfdg == SizeSelectIfGTE && (_ccca >= _bafb && _agbb >= _ddge)) {
				_fgef = 1
			}
		}
		_bgad.AddInt(_fgef)
	}
	return _bgad
}
func TstAddSymbol(t *_f.T, bms *Bitmaps, sym *Bitmap, x *int, y int, space int) {
	bms.AddBitmap(sym)
	_aegg := _aa.Rect(*x, y, *x+sym.Width, y+sym.Height)
	bms.AddBox(&_aegg)
	*x += sym.Width + space
}

var _fdde = [5]int{1, 2, 3, 0, 4}

func (_gaff *Boxes) Add(box *_aa.Rectangle) error {
	if _gaff == nil {
		return _e.Error("\u0042o\u0078\u0065\u0073\u002e\u0041\u0064d", "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	*_gaff = append(*_gaff, box)
	return nil
}
func (_gbgfb *Bitmaps) WidthSorter() func(_gcbbd, _eadb int) bool {
	return func(_agfb, _cbad int) bool { return _gbgfb.Values[_agfb].Width < _gbgfb.Values[_cbad].Width }
}
func _dgdg(_fdcf, _aeae *Bitmap, _gdbb, _agee int) (*Bitmap, error) {
	const _eedb = "d\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063\u006b"
	if _aeae == nil {
		_ca.Log.Debug("\u0064\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063k\u0020\u0073\u006f\u0075\u0072\u0063\u0065 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
		return nil, _e.Error(_eedb, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _gdbb < 1 || _agee < 1 {
		return nil, _e.Error(_eedb, "\u0068\u0053\u007a\u0069\u0065 \u0061\u006e\u0064\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u0061\u0072\u0065 \u006e\u006f\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _gdbb == 1 && _agee == 1 {
		_cdebf, _ebbbf := _cdcf(_fdcf, _aeae)
		if _ebbbf != nil {
			return nil, _e.Wrap(_ebbbf, _eedb, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _cdebf, nil
	}
	if _gdbb == 1 || _agee == 1 {
		_cefgg := SelCreateBrick(_agee, _gdbb, _agee/2, _gdbb/2, SelHit)
		_cfeb, _bacc := _caagg(_fdcf, _aeae, _cefgg)
		if _bacc != nil {
			return nil, _e.Wrap(_bacc, _eedb, "\u0068s\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _cfeb, nil
	}
	_ecbg := SelCreateBrick(1, _gdbb, 0, _gdbb/2, SelHit)
	_dbef := SelCreateBrick(_agee, 1, _agee/2, 0, SelHit)
	_gadb, _bcaaf := _caagg(nil, _aeae, _ecbg)
	if _bcaaf != nil {
		return nil, _e.Wrap(_bcaaf, _eedb, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_fdcf, _bcaaf = _caagg(_fdcf, _gadb, _dbef)
	if _bcaaf != nil {
		return nil, _e.Wrap(_bcaaf, _eedb, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _fdcf, nil
}
func (_agebd *Selection) findMaxTranslations() (_efbgc, _edfad, _dcbc, _cdfb int) {
	for _gdee := 0; _gdee < _agebd.Height; _gdee++ {
		for _gddg := 0; _gddg < _agebd.Width; _gddg++ {
			if _agebd.Data[_gdee][_gddg] == SelHit {
				_efbgc = _aefd(_efbgc, _agebd.Cx-_gddg)
				_edfad = _aefd(_edfad, _agebd.Cy-_gdee)
				_dcbc = _aefd(_dcbc, _gddg-_agebd.Cx)
				_cdfb = _aefd(_cdfb, _gdee-_agebd.Cy)
			}
		}
	}
	return _efbgc, _edfad, _dcbc, _cdfb
}
func _ggdga(_fcdbe, _feadd int, _ggfc string) *Selection {
	_afbf := &Selection{Height: _fcdbe, Width: _feadd, Name: _ggfc}
	_afbf.Data = make([][]SelectionValue, _fcdbe)
	for _aaecf := 0; _aaecf < _fcdbe; _aaecf++ {
		_afbf.Data[_aaecf] = make([]SelectionValue, _feadd)
	}
	return _afbf
}

type SizeSelection int

func Extract(roi _aa.Rectangle, src *Bitmap) (*Bitmap, error) {
	_cadc := New(roi.Dx(), roi.Dy())
	_fege := roi.Min.X & 0x07
	_abf := 8 - _fege
	_dbce := uint(8 - _cadc.Width&0x07)
	_abdcf := src.GetByteIndex(roi.Min.X, roi.Min.Y)
	_ega := src.GetByteIndex(roi.Max.X-1, roi.Min.Y)
	_cac := _cadc.RowStride == _ega+1-_abdcf
	var _dda int
	for _eeed := roi.Min.Y; _eeed < roi.Max.Y; _eeed++ {
		_gacdc := _abdcf
		_gfcb := _dda
		switch {
		case _abdcf == _ega:
			_afgf, _dgabb := src.GetByte(_gacdc)
			if _dgabb != nil {
				return nil, _dgabb
			}
			_afgf <<= uint(_fege)
			_dgabb = _cadc.SetByte(_gfcb, _fcgbg(_dbce, _afgf))
			if _dgabb != nil {
				return nil, _dgabb
			}
		case _fege == 0:
			for _gdbc := _abdcf; _gdbc <= _ega; _gdbc++ {
				_dbbc, _fged := src.GetByte(_gacdc)
				if _fged != nil {
					return nil, _fged
				}
				_gacdc++
				if _gdbc == _ega && _cac {
					_dbbc = _fcgbg(_dbce, _dbbc)
				}
				_fged = _cadc.SetByte(_gfcb, _dbbc)
				if _fged != nil {
					return nil, _fged
				}
				_gfcb++
			}
		default:
			_fbfed := _fafg(src, _cadc, uint(_fege), uint(_abf), _dbce, _abdcf, _ega, _cac, _gacdc, _gfcb)
			if _fbfed != nil {
				return nil, _fbfed
			}
		}
		_abdcf += src.RowStride
		_ega += src.RowStride
		_dda += _cadc.RowStride
	}
	return _cadc, nil
}
func (_beaa *Bitmap) setEightFullBytes(_dggc int, _gfac uint64) error {
	if _dggc+7 > len(_beaa.Data)-1 {
		return _e.Error("\u0073\u0065\u0074\u0045\u0069\u0067\u0068\u0074\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_beaa.Data[_dggc] = byte((_gfac & 0xff00000000000000) >> 56)
	_beaa.Data[_dggc+1] = byte((_gfac & 0xff000000000000) >> 48)
	_beaa.Data[_dggc+2] = byte((_gfac & 0xff0000000000) >> 40)
	_beaa.Data[_dggc+3] = byte((_gfac & 0xff00000000) >> 32)
	_beaa.Data[_dggc+4] = byte((_gfac & 0xff000000) >> 24)
	_beaa.Data[_dggc+5] = byte((_gfac & 0xff0000) >> 16)
	_beaa.Data[_dggc+6] = byte((_gfac & 0xff00) >> 8)
	_beaa.Data[_dggc+7] = byte(_gfac & 0xff)
	return nil
}
func (_afce *Bitmap) clipRectangle(_fgb, _dgb *_aa.Rectangle) (_bedg *Bitmap, _cced error) {
	const _gafe = "\u0063\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if _fgb == nil {
		return nil, _e.Error(_gafe, "\u0070r\u006fv\u0069\u0064\u0065\u0064\u0020n\u0069\u006c \u0027\u0062\u006f\u0078\u0027")
	}
	_aaggd, _dgbb := _afce.Width, _afce.Height
	_acfd, _cced := ClipBoxToRectangle(_fgb, _aaggd, _dgbb)
	if _cced != nil {
		_ca.Log.Warning("\u0027\u0062ox\u0027\u0020\u0064o\u0065\u0073\u006e\u0027t o\u0076er\u006c\u0061\u0070\u0020\u0062\u0069\u0074ma\u0070\u0020\u0027\u0062\u0027\u003a\u0020%\u0076", _cced)
		return nil, nil
	}
	_bag, _egca := _acfd.Min.X, _acfd.Min.Y
	_fcba, _aagc := _acfd.Max.X-_acfd.Min.X, _acfd.Max.Y-_acfd.Min.Y
	_bedg = New(_fcba, _aagc)
	_bedg.Text = _afce.Text
	if _cced = _bedg.RasterOperation(0, 0, _fcba, _aagc, PixSrc, _afce, _bag, _egca); _cced != nil {
		return nil, _e.Wrap(_cced, _gafe, "")
	}
	if _dgb != nil {
		*_dgb = *_acfd
	}
	return _bedg, nil
}
func (_fce *Bitmap) Equivalent(s *Bitmap) bool { return _fce.equivalent(s) }
func _aaa(_dab, _dc *Bitmap) (_adg error) {
	const _gaf = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0038"
	_dcd := _dc.RowStride
	_fc := _dab.RowStride
	var _be, _eba, _aef, _ef, _fag int
	for _aef = 0; _aef < _dc.Height; _aef++ {
		_be = _aef * _dcd
		_eba = 8 * _aef * _fc
		for _ef = 0; _ef < _dcd; _ef++ {
			if _adg = _dab.setEightBytes(_eba+_ef*8, _cede[_dc.Data[_be+_ef]]); _adg != nil {
				return _e.Wrap(_adg, _gaf, "")
			}
		}
		for _fag = 1; _fag < 8; _fag++ {
			for _ef = 0; _ef < _fc; _ef++ {
				if _adg = _dab.SetByte(_eba+_fag*_fc+_ef, _dab.Data[_eba+_ef]); _adg != nil {
					return _e.Wrap(_adg, _gaf, "")
				}
			}
		}
	}
	return nil
}
func (_abgb Points) XSorter() func(_egfc, _fade int) bool {
	return func(_gcdb, _cade int) bool { return _abgb[_gcdb].X < _abgb[_cade].X }
}
func _ad(_b, _da *Bitmap) (_ceg error) {
	const _caa = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0032"
	_ade := _da.RowStride
	_cb := _b.RowStride
	var (
		_df                      byte
		_acf                     uint16
		_bc, _eb, _cc, _fb, _cbf int
	)
	for _cc = 0; _cc < _da.Height; _cc++ {
		_bc = _cc * _ade
		_eb = 2 * _cc * _cb
		for _fb = 0; _fb < _ade; _fb++ {
			_df = _da.Data[_bc+_fb]
			_acf = _dfag[_df]
			_cbf = _eb + _fb*2
			if _b.RowStride != _da.RowStride*2 && (_fb+1)*2 > _b.RowStride {
				_ceg = _b.SetByte(_cbf, byte(_acf>>8))
			} else {
				_ceg = _b.setTwoBytes(_cbf, _acf)
			}
			if _ceg != nil {
				return _e.Wrap(_ceg, _caa, "")
			}
		}
		for _fb = 0; _fb < _cb; _fb++ {
			_cbf = _eb + _cb + _fb
			_df = _b.Data[_eb+_fb]
			if _ceg = _b.SetByte(_cbf, _df); _ceg != nil {
				return _e.Wrapf(_ceg, _caa, "c\u006f\u0070\u0079\u0020\u0064\u006fu\u0062\u006c\u0065\u0064\u0020\u006ci\u006e\u0065\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0042\u0079\u0074\u0065\u003a\u0020\u0027\u0025d\u0027", _eb+_fb, _eb+_cb+_fb)
			}
		}
	}
	return nil
}
func (_afcf *Bitmap) removeBorderGeneral(_dfdb, _bca, _gad, _bgb int) (*Bitmap, error) {
	const _ffgc = "\u0072\u0065\u006d\u006fve\u0042\u006f\u0072\u0064\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _dfdb < 0 || _bca < 0 || _gad < 0 || _bgb < 0 {
		return nil, _e.Error(_ffgc, "\u006e\u0065g\u0061\u0074\u0069\u0076\u0065\u0020\u0062\u0072\u006f\u0064\u0065\u0072\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u0076\u0061lu\u0065\u0073")
	}
	_cdb, _agge := _afcf.Width, _afcf.Height
	_adfc := _cdb - _dfdb - _bca
	_febf := _agge - _gad - _bgb
	if _adfc <= 0 {
		return nil, _e.Errorf(_ffgc, "w\u0069\u0064\u0074\u0068: \u0025d\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u003e\u0020\u0030", _adfc)
	}
	if _febf <= 0 {
		return nil, _e.Errorf(_ffgc, "\u0068\u0065\u0069\u0067ht\u003a\u0020\u0025\u0064\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030", _febf)
	}
	_ceed := New(_adfc, _febf)
	_ceed.Color = _afcf.Color
	_cffb := _ceed.RasterOperation(0, 0, _adfc, _febf, PixSrc, _afcf, _dfdb, _gad)
	if _cffb != nil {
		return nil, _e.Wrap(_cffb, _ffgc, "")
	}
	return _ceed, nil
}
func MakePixelSumTab8() []int { return _dfbeb() }
func _ddgf(_bgade *Bitmap, _gfdc *_ac.Stack, _dbcf, _gdcf int) (_gfed *_aa.Rectangle, _dgdf error) {
	const _ddgd = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _bgade == nil {
		return nil, _e.Error(_ddgd, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _gfdc == nil {
		return nil, _e.Error(_ddgd, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_aaeg, _bbfcb := _bgade.Width, _bgade.Height
	_bfbe := _aaeg - 1
	_ddfg := _bbfcb - 1
	if _dbcf < 0 || _dbcf > _bfbe || _gdcf < 0 || _gdcf > _ddfg || !_bgade.GetPixel(_dbcf, _gdcf) {
		return nil, nil
	}
	var _cgcf *_aa.Rectangle
	_cgcf, _dgdf = Rect(100000, 100000, 0, 0)
	if _dgdf != nil {
		return nil, _e.Wrap(_dgdf, _ddgd, "")
	}
	if _dgdf = _agadg(_gfdc, _dbcf, _dbcf, _gdcf, 1, _ddfg, _cgcf); _dgdf != nil {
		return nil, _e.Wrap(_dgdf, _ddgd, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _dgdf = _agadg(_gfdc, _dbcf, _dbcf, _gdcf+1, -1, _ddfg, _cgcf); _dgdf != nil {
		return nil, _e.Wrap(_dgdf, _ddgd, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_cgcf.Min.X, _cgcf.Max.X = _dbcf, _dbcf
	_cgcf.Min.Y, _cgcf.Max.Y = _gdcf, _gdcf
	var (
		_gcbd *fillSegment
		_ebbe int
	)
	for _gfdc.Len() > 0 {
		if _gcbd, _dgdf = _cfee(_gfdc); _dgdf != nil {
			return nil, _e.Wrap(_dgdf, _ddgd, "")
		}
		_gdcf = _gcbd._ffeg
		for _dbcf = _gcbd._abbc; _dbcf >= 0 && _bgade.GetPixel(_dbcf, _gdcf); _dbcf-- {
			if _dgdf = _bgade.SetPixel(_dbcf, _gdcf, 0); _dgdf != nil {
				return nil, _e.Wrap(_dgdf, _ddgd, "")
			}
		}
		if _dbcf >= _gcbd._abbc {
			for _dbcf++; _dbcf <= _gcbd._bcaf && _dbcf <= _bfbe && !_bgade.GetPixel(_dbcf, _gdcf); _dbcf++ {
			}
			_ebbe = _dbcf
			if !(_dbcf <= _gcbd._bcaf && _dbcf <= _bfbe) {
				continue
			}
		} else {
			_ebbe = _dbcf + 1
			if _ebbe < _gcbd._abbc-1 {
				if _dgdf = _agadg(_gfdc, _ebbe, _gcbd._abbc-1, _gcbd._ffeg, -_gcbd._bedfb, _ddfg, _cgcf); _dgdf != nil {
					return nil, _e.Wrap(_dgdf, _ddgd, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
				}
			}
			_dbcf = _gcbd._abbc + 1
		}
		for {
			for ; _dbcf <= _bfbe && _bgade.GetPixel(_dbcf, _gdcf); _dbcf++ {
				if _dgdf = _bgade.SetPixel(_dbcf, _gdcf, 0); _dgdf != nil {
					return nil, _e.Wrap(_dgdf, _ddgd, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _dgdf = _agadg(_gfdc, _ebbe, _dbcf-1, _gcbd._ffeg, _gcbd._bedfb, _ddfg, _cgcf); _dgdf != nil {
				return nil, _e.Wrap(_dgdf, _ddgd, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _dbcf > _gcbd._bcaf+1 {
				if _dgdf = _agadg(_gfdc, _gcbd._bcaf+1, _dbcf-1, _gcbd._ffeg, -_gcbd._bedfb, _ddfg, _cgcf); _dgdf != nil {
					return nil, _e.Wrap(_dgdf, _ddgd, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _dbcf++; _dbcf <= _gcbd._bcaf && _dbcf <= _bfbe && !_bgade.GetPixel(_dbcf, _gdcf); _dbcf++ {
			}
			_ebbe = _dbcf
			if !(_dbcf <= _gcbd._bcaf && _dbcf <= _bfbe) {
				break
			}
		}
	}
	_cgcf.Max.X++
	_cgcf.Max.Y++
	return _cgcf, nil
}
func (_gfee *ClassedPoints) Len() int { return _gfee.IntSlice.Size() }
func (_agbe *Bitmap) equivalent(_ebae *Bitmap) bool {
	if _agbe == _ebae {
		return true
	}
	if !_agbe.SizesEqual(_ebae) {
		return false
	}
	_cafa := _ffb(_agbe, _ebae, CmbOpXor)
	_adf := _agbe.countPixels()
	_gdd := int(0.25 * float32(_adf))
	if _cafa.thresholdPixelSum(_gdd) {
		return false
	}
	var (
		_bedb [9][9]int
		_cda  [18][9]int
		_dec  [9][18]int
		_gde  int
		_eee  int
	)
	_abc := 9
	_dcgd := _agbe.Height / _abc
	_cfb := _agbe.Width / _abc
	_adb, _fbfe := _dcgd/2, _cfb/2
	if _dcgd < _cfb {
		_adb = _cfb / 2
		_fbfe = _dcgd / 2
	}
	_edcg := float64(_adb) * float64(_fbfe) * _ea.Pi
	_becb := int(float64(_dcgd*_cfb/2) * 0.9)
	_gfaf := int(float64(_cfb*_dcgd/2) * 0.9)
	for _cbbag := 0; _cbbag < _abc; _cbbag++ {
		_eeab := _cfb*_cbbag + _gde
		var _gaee int
		if _cbbag == _abc-1 {
			_gde = 0
			_gaee = _agbe.Width
		} else {
			_gaee = _eeab + _cfb
			if ((_agbe.Width - _gde) % _abc) > 0 {
				_gde++
				_gaee++
			}
		}
		for _cfbf := 0; _cfbf < _abc; _cfbf++ {
			_cfgg := _dcgd*_cfbf + _eee
			var _efg int
			if _cfbf == _abc-1 {
				_eee = 0
				_efg = _agbe.Height
			} else {
				_efg = _cfgg + _dcgd
				if (_agbe.Height-_eee)%_abc > 0 {
					_eee++
					_efg++
				}
			}
			var _gdec, _ebg, _gcf, _cbg int
			_ceb := (_eeab + _gaee) / 2
			_aab := (_cfgg + _efg) / 2
			for _dbgc := _eeab; _dbgc < _gaee; _dbgc++ {
				for _dgbf := _cfgg; _dgbf < _efg; _dgbf++ {
					if _cafa.GetPixel(_dbgc, _dgbf) {
						if _dbgc < _ceb {
							_gdec++
						} else {
							_ebg++
						}
						if _dgbf < _aab {
							_cbg++
						} else {
							_gcf++
						}
					}
				}
			}
			_bedb[_cbbag][_cfbf] = _gdec + _ebg
			_cda[_cbbag*2][_cfbf] = _gdec
			_cda[_cbbag*2+1][_cfbf] = _ebg
			_dec[_cbbag][_cfbf*2] = _cbg
			_dec[_cbbag][_cfbf*2+1] = _gcf
		}
	}
	for _dgd := 0; _dgd < _abc*2-1; _dgd++ {
		for _adaf := 0; _adaf < (_abc - 1); _adaf++ {
			var _dggb int
			for _cee := 0; _cee < 2; _cee++ {
				for _bbbc := 0; _bbbc < 2; _bbbc++ {
					_dggb += _cda[_dgd+_cee][_adaf+_bbbc]
				}
			}
			if _dggb > _gfaf {
				return false
			}
		}
	}
	for _geae := 0; _geae < (_abc - 1); _geae++ {
		for _cfbc := 0; _cfbc < ((_abc * 2) - 1); _cfbc++ {
			var _aafa int
			for _ead := 0; _ead < 2; _ead++ {
				for _ffd := 0; _ffd < 2; _ffd++ {
					_aafa += _dec[_geae+_ead][_cfbc+_ffd]
				}
			}
			if _aafa > _becb {
				return false
			}
		}
	}
	for _cdeg := 0; _cdeg < (_abc - 2); _cdeg++ {
		for _bgd := 0; _bgd < (_abc - 2); _bgd++ {
			var _deec, _badb int
			for _ebd := 0; _ebd < 3; _ebd++ {
				for _geb := 0; _geb < 3; _geb++ {
					if _ebd == _geb {
						_deec += _bedb[_cdeg+_ebd][_bgd+_geb]
					}
					if (2 - _ebd) == _geb {
						_badb += _bedb[_cdeg+_ebd][_bgd+_geb]
					}
				}
			}
			if _deec > _gfaf || _badb > _gfaf {
				return false
			}
		}
	}
	for _cdee := 0; _cdee < (_abc - 1); _cdee++ {
		for _geg := 0; _geg < (_abc - 1); _geg++ {
			var _daabe int
			for _cbcc := 0; _cbcc < 2; _cbcc++ {
				for _eeeg := 0; _eeeg < 2; _eeeg++ {
					_daabe += _bedb[_cdee+_cbcc][_geg+_eeeg]
				}
			}
			if float64(_daabe) > _edcg {
				return false
			}
		}
	}
	return true
}

type fillSegment struct {
	_abbc  int
	_bcaf  int
	_ffeg  int
	_bedfb int
}
type Bitmap struct {
	Width, Height            int
	BitmapNumber             int
	RowStride                int
	Data                     []byte
	Color                    Color
	Special                  int
	Text                     string
	XResolution, YResolution int
}

func (_dcga *Bitmap) thresholdPixelSum(_beee int) bool {
	var (
		_fcge int
		_gbcg uint8
		_cccg byte
		_caba int
	)
	_dcgc := _dcga.RowStride
	_gdfa := uint(_dcga.Width & 0x07)
	if _gdfa != 0 {
		_gbcg = uint8((0xff << (8 - _gdfa)) & 0xff)
		_dcgc--
	}
	for _cgc := 0; _cgc < _dcga.Height; _cgc++ {
		for _caba = 0; _caba < _dcgc; _caba++ {
			_cccg = _dcga.Data[_cgc*_dcga.RowStride+_caba]
			_fcge += int(_fde[_cccg])
		}
		if _gdfa != 0 {
			_cccg = _dcga.Data[_cgc*_dcga.RowStride+_caba] & _gbcg
			_fcge += int(_fde[_cccg])
		}
		if _fcge > _beee {
			return true
		}
	}
	return false
}

type Component int

func (_egd *Bitmap) connComponentsBB(_cec int) (_faec *Boxes, _dcdge error) {
	const _cddd = "\u0042\u0069\u0074ma\u0070\u002e\u0063\u006f\u006e\u006e\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0042"
	if _cec != 4 && _cec != 8 {
		return nil, _e.Error(_cddd, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _egd.Zero() {
		return &Boxes{}, nil
	}
	_egd.setPadBits(0)
	_cbd, _dcdge := _cdcf(nil, _egd)
	if _dcdge != nil {
		return nil, _e.Wrap(_dcdge, _cddd, "\u0062\u006d\u0031")
	}
	_ggeg := &_ac.Stack{}
	_ggeg.Aux = &_ac.Stack{}
	_faec = &Boxes{}
	var (
		_dfed, _cga int
		_ffde       _aa.Point
		_adcf       bool
		_afbd       *_aa.Rectangle
	)
	for {
		if _ffde, _adcf, _dcdge = _cbd.nextOnPixel(_cga, _dfed); _dcdge != nil {
			return nil, _e.Wrap(_dcdge, _cddd, "")
		}
		if !_adcf {
			break
		}
		if _afbd, _dcdge = _beaad(_cbd, _ggeg, _ffde.X, _ffde.Y, _cec); _dcdge != nil {
			return nil, _e.Wrap(_dcdge, _cddd, "")
		}
		if _dcdge = _faec.Add(_afbd); _dcdge != nil {
			return nil, _e.Wrap(_dcdge, _cddd, "")
		}
		_cga = _ffde.X
		_dfed = _ffde.Y
	}
	return _faec, nil
}
func (_fcbg *byHeight) Less(i, j int) bool { return _fcbg.Values[i].Height < _fcbg.Values[j].Height }
func (_fac *Bitmap) ToImage() _aa.Image {
	_fead, _ccc := _g.NewImage(_fac.Width, _fac.Height, 1, 1, _fac.Data, nil, nil)
	if _ccc != nil {
		_ca.Log.Error("\u0043\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020j\u0062\u0069\u0067\u0032\u002e\u0042\u0069\u0074m\u0061p\u0020\u0074\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002e\u0049\u006d\u0061\u0067e\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ccc)
	}
	return _fead
}
func (_caaf *Bitmap) inverseData() {
	if _bdfd := _caaf.RasterOperation(0, 0, _caaf.Width, _caaf.Height, PixNotDst, nil, 0, 0); _bdfd != nil {
		_ca.Log.Debug("\u0049n\u0076\u0065\u0072\u0073e\u0020\u0064\u0061\u0074\u0061 \u0066a\u0069l\u0065\u0064\u003a\u0020\u0027\u0025\u0076'", _bdfd)
	}
	if _caaf.Color == Chocolate {
		_caaf.Color = Vanilla
	} else {
		_caaf.Color = Chocolate
	}
}
func (_dbga *Bitmap) GetVanillaData() []byte {
	if _dbga.Color == Chocolate {
		_dbga.inverseData()
	}
	return _dbga.Data
}

type BoundaryCondition int

func _eec(_ggdf *Bitmap, _beg int) (*Bitmap, error) {
	const _eaf = "\u0065x\u0070a\u006e\u0064\u0042\u0069\u006ea\u0072\u0079P\u006f\u0077\u0065\u0072\u0032"
	if _ggdf == nil {
		return nil, _e.Error(_eaf, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _beg == 1 {
		return _cdcf(nil, _ggdf)
	}
	if _beg != 2 && _beg != 4 && _beg != 8 {
		return nil, _e.Error(_eaf, "\u0066\u0061\u0063t\u006f\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069n\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_fdf := _beg * _ggdf.Width
	_acb := _beg * _ggdf.Height
	_ag := New(_fdf, _acb)
	var _dee error
	switch _beg {
	case 2:
		_dee = _ad(_ag, _ggdf)
	case 4:
		_dee = _gf(_ag, _ggdf)
	case 8:
		_dee = _aaa(_ag, _ggdf)
	}
	if _dee != nil {
		return nil, _e.Wrap(_dee, _eaf, "")
	}
	return _ag, nil
}
func (_ggcf *Bitmap) ThresholdPixelSum(thresh int, tab8 []int) (_aac bool, _fae error) {
	const _dbbg = "\u0042i\u0074\u006d\u0061\u0070\u002e\u0054\u0068\u0072\u0065\u0073\u0068o\u006c\u0064\u0050\u0069\u0078\u0065\u006c\u0053\u0075\u006d"
	if tab8 == nil {
		tab8 = _dfbeb()
	}
	_ebab := _ggcf.Width >> 3
	_bfb := _ggcf.Width & 7
	_dcdd := byte(0xff << uint(8-_bfb))
	var (
		_eea, _bbb, _ecgf, _aebd int
		_fbbf                    byte
	)
	for _eea = 0; _eea < _ggcf.Height; _eea++ {
		_ecgf = _ggcf.RowStride * _eea
		for _bbb = 0; _bbb < _ebab; _bbb++ {
			_fbbf, _fae = _ggcf.GetByte(_ecgf + _bbb)
			if _fae != nil {
				return false, _e.Wrap(_fae, _dbbg, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
			}
			_aebd += tab8[_fbbf]
		}
		if _bfb != 0 {
			_fbbf, _fae = _ggcf.GetByte(_ecgf + _bbb)
			if _fae != nil {
				return false, _e.Wrap(_fae, _dbbg, "p\u0061\u0072\u0074\u0069\u0061\u006c\u0042\u0079\u0074\u0065")
			}
			_fbbf &= _dcdd
			_aebd += tab8[_fbbf]
		}
		if _aebd > thresh {
			return true, nil
		}
	}
	return _aac, nil
}
func (_fbbg *Bitmap) And(s *Bitmap) (_bbdc *Bitmap, _fbbgf error) {
	const _fgcg = "\u0042\u0069\u0074\u006d\u0061\u0070\u002e\u0041\u006e\u0064"
	if _fbbg == nil {
		return nil, _e.Error(_fgcg, "\u0027b\u0069t\u006d\u0061\u0070\u0020\u0027b\u0027\u0020i\u0073\u0020\u006e\u0069\u006c")
	}
	if s == nil {
		return nil, _e.Error(_fgcg, "\u0062\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if !_fbbg.SizesEqual(s) {
		_ca.Log.Debug("\u0025\u0073\u0020-\u0020\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0069\u007a\u0065 \u0077\u0069\u0074\u0068\u0020\u0027\u0062\u0027", _fgcg)
	}
	if _bbdc, _fbbgf = _cdcf(_bbdc, _fbbg); _fbbgf != nil {
		return nil, _e.Wrap(_fbbgf, _fgcg, "\u0063\u0061\u006e't\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _fbbgf = _bbdc.RasterOperation(0, 0, _bbdc.Width, _bbdc.Height, PixSrcAndDst, s, 0, 0); _fbbgf != nil {
		return nil, _e.Wrap(_fbbgf, _fgcg, "")
	}
	return _bbdc, nil
}
func TstPSymbol(t *_f.T) *Bitmap {
	t.Helper()
	_cgdd := New(5, 8)
	_d.NoError(t, _cgdd.SetPixel(0, 0, 1))
	_d.NoError(t, _cgdd.SetPixel(1, 0, 1))
	_d.NoError(t, _cgdd.SetPixel(2, 0, 1))
	_d.NoError(t, _cgdd.SetPixel(3, 0, 1))
	_d.NoError(t, _cgdd.SetPixel(4, 1, 1))
	_d.NoError(t, _cgdd.SetPixel(0, 1, 1))
	_d.NoError(t, _cgdd.SetPixel(4, 2, 1))
	_d.NoError(t, _cgdd.SetPixel(0, 2, 1))
	_d.NoError(t, _cgdd.SetPixel(4, 3, 1))
	_d.NoError(t, _cgdd.SetPixel(0, 3, 1))
	_d.NoError(t, _cgdd.SetPixel(0, 4, 1))
	_d.NoError(t, _cgdd.SetPixel(1, 4, 1))
	_d.NoError(t, _cgdd.SetPixel(2, 4, 1))
	_d.NoError(t, _cgdd.SetPixel(3, 4, 1))
	_d.NoError(t, _cgdd.SetPixel(0, 5, 1))
	_d.NoError(t, _cgdd.SetPixel(0, 6, 1))
	_d.NoError(t, _cgdd.SetPixel(0, 7, 1))
	return _cgdd
}
func (_edfab *Bitmaps) HeightSorter() func(_ccgde, _cgfdb int) bool {
	return func(_ccfa, _ebbfb int) bool {
		_cgfdg := _edfab.Values[_ccfa].Height < _edfab.Values[_ebbfb].Height
		_ca.Log.Debug("H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0076\u0020<\u0020\u0025\u0076\u0020= \u0025\u0076", _edfab.Values[_ccfa].Height, _edfab.Values[_ebbfb].Height, _cgfdg)
		return _cgfdg
	}
}

var _fde [256]uint8

type shift int

func CorrelationScoreThresholded(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab, downcount []int, scoreThreshold float32) (bool, error) {
	const _cfeg = "C\u006f\u0072\u0072\u0065\u006c\u0061t\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054h\u0072\u0065\u0073h\u006fl\u0064\u0065\u0064"
	if bm1 == nil {
		return false, _e.Error(_cfeg, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d1\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if bm2 == nil {
		return false, _e.Error(_cfeg, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d2\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if area1 <= 0 || area2 <= 0 {
		return false, _e.Error(_cfeg, "c\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0053\u0063\u006f\u0072\u0065\u0054\u0068re\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u002d\u0020\u0061\u0072\u0065\u0061s \u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u003e\u0020\u0030")
	}
	if downcount == nil {
		return false, _e.Error(_cfeg, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u006f\u0020\u0027\u0064\u006f\u0077\u006e\u0063\u006f\u0075\u006e\u0074\u0027")
	}
	if tab == nil {
		return false, _e.Error(_cfeg, "p\u0072\u006f\u0076\u0069de\u0064 \u006e\u0069\u006c\u0020\u0027s\u0075\u006d\u0074\u0061\u0062\u0027")
	}
	_aded, _cdge := bm1.Width, bm1.Height
	_ffcce, _abfa := bm2.Width, bm2.Height
	if _ac.Abs(_aded-_ffcce) > maxDiffW {
		return false, nil
	}
	if _ac.Abs(_cdge-_abfa) > maxDiffH {
		return false, nil
	}
	_fdda := int(delX + _ac.Sign(delX)*0.5)
	_bgf := int(delY + _ac.Sign(delY)*0.5)
	_aagcf := int(_ea.Ceil(_ea.Sqrt(float64(scoreThreshold) * float64(area1) * float64(area2))))
	_ceae := bm2.RowStride
	_eagf := _aefd(_bgf, 0)
	_ecfa := _gcg(_abfa+_bgf, _cdge)
	_edbd := bm1.RowStride * _eagf
	_fcgc := bm2.RowStride * (_eagf - _bgf)
	var _gbga int
	if _ecfa <= _cdge {
		_gbga = downcount[_ecfa-1]
	}
	_fggd := _aefd(_fdda, 0)
	_bgba := _gcg(_ffcce+_fdda, _aded)
	var _baa, _dfaa int
	if _fdda >= 8 {
		_baa = _fdda >> 3
		_edbd += _baa
		_fggd -= _baa << 3
		_bgba -= _baa << 3
		_fdda &= 7
	} else if _fdda <= -8 {
		_dfaa = -((_fdda + 7) >> 3)
		_fcgc += _dfaa
		_ceae -= _dfaa
		_fdda += _dfaa << 3
	}
	var (
		_abac, _fdfb, _eead int
		_fbfd, _gdag, _eebd byte
	)
	if _fggd >= _bgba || _eagf >= _ecfa {
		return false, nil
	}
	_geca := (_bgba + 7) >> 3
	switch {
	case _fdda == 0:
		for _fdfb = _eagf; _fdfb < _ecfa; _fdfb, _edbd, _fcgc = _fdfb+1, _edbd+bm1.RowStride, _fcgc+bm2.RowStride {
			for _eead = 0; _eead < _geca; _eead++ {
				_fbfd = bm1.Data[_edbd+_eead] & bm2.Data[_fcgc+_eead]
				_abac += tab[_fbfd]
			}
			if _abac >= _aagcf {
				return true, nil
			}
			if _geefe := _abac + downcount[_fdfb] - _gbga; _geefe < _aagcf {
				return false, nil
			}
		}
	case _fdda > 0 && _ceae < _geca:
		for _fdfb = _eagf; _fdfb < _ecfa; _fdfb, _edbd, _fcgc = _fdfb+1, _edbd+bm1.RowStride, _fcgc+bm2.RowStride {
			_gdag = bm1.Data[_edbd]
			_eebd = bm2.Data[_fcgc] >> uint(_fdda)
			_fbfd = _gdag & _eebd
			_abac += tab[_fbfd]
			for _eead = 1; _eead < _ceae; _eead++ {
				_gdag = bm1.Data[_edbd+_eead]
				_eebd = bm2.Data[_fcgc+_eead]>>uint(_fdda) | bm2.Data[_fcgc+_eead-1]<<uint(8-_fdda)
				_fbfd = _gdag & _eebd
				_abac += tab[_fbfd]
			}
			_gdag = bm1.Data[_edbd+_eead]
			_eebd = bm2.Data[_fcgc+_eead-1] << uint(8-_fdda)
			_fbfd = _gdag & _eebd
			_abac += tab[_fbfd]
			if _abac >= _aagcf {
				return true, nil
			} else if _abac+downcount[_fdfb]-_gbga < _aagcf {
				return false, nil
			}
		}
	case _fdda > 0 && _ceae >= _geca:
		for _fdfb = _eagf; _fdfb < _ecfa; _fdfb, _edbd, _fcgc = _fdfb+1, _edbd+bm1.RowStride, _fcgc+bm2.RowStride {
			_gdag = bm1.Data[_edbd]
			_eebd = bm2.Data[_fcgc] >> uint(_fdda)
			_fbfd = _gdag & _eebd
			_abac += tab[_fbfd]
			for _eead = 1; _eead < _geca; _eead++ {
				_gdag = bm1.Data[_edbd+_eead]
				_eebd = bm2.Data[_fcgc+_eead] >> uint(_fdda)
				_eebd |= bm2.Data[_fcgc+_eead-1] << uint(8-_fdda)
				_fbfd = _gdag & _eebd
				_abac += tab[_fbfd]
			}
			if _abac >= _aagcf {
				return true, nil
			} else if _abac+downcount[_fdfb]-_gbga < _aagcf {
				return false, nil
			}
		}
	case _geca < _ceae:
		for _fdfb = _eagf; _fdfb < _ecfa; _fdfb, _edbd, _fcgc = _fdfb+1, _edbd+bm1.RowStride, _fcgc+bm2.RowStride {
			for _eead = 0; _eead < _geca; _eead++ {
				_gdag = bm1.Data[_edbd+_eead]
				_eebd = bm2.Data[_fcgc+_eead] << uint(-_fdda)
				_eebd |= bm2.Data[_fcgc+_eead+1] >> uint(8+_fdda)
				_fbfd = _gdag & _eebd
				_abac += tab[_fbfd]
			}
			if _abac >= _aagcf {
				return true, nil
			} else if _cfgf := _abac + downcount[_fdfb] - _gbga; _cfgf < _aagcf {
				return false, nil
			}
		}
	case _ceae >= _geca:
		for _fdfb = _eagf; _fdfb < _ecfa; _fdfb, _edbd, _fcgc = _fdfb+1, _edbd+bm1.RowStride, _fcgc+bm2.RowStride {
			for _eead = 0; _eead < _geca; _eead++ {
				_gdag = bm1.Data[_edbd+_eead]
				_eebd = bm2.Data[_fcgc+_eead] << uint(-_fdda)
				_eebd |= bm2.Data[_fcgc+_eead+1] >> uint(8+_fdda)
				_fbfd = _gdag & _eebd
				_abac += tab[_fbfd]
			}
			_gdag = bm1.Data[_edbd+_eead]
			_eebd = bm2.Data[_fcgc+_eead] << uint(-_fdda)
			_fbfd = _gdag & _eebd
			_abac += tab[_fbfd]
			if _abac >= _aagcf {
				return true, nil
			} else if _abac+downcount[_fdfb]-_gbga < _aagcf {
				return false, nil
			}
		}
	}
	_ggdg := float32(_abac) * float32(_abac) / (float32(area1) * float32(area2))
	if _ggdg >= scoreThreshold {
		_ca.Log.Trace("\u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064\u0020\u003c\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0020\u0025\u0064\u0020\u0062\u0075\u0074\u0020\u0073c\u006f\u0072\u0065\u0020\u0025\u0066\u0020\u003e\u003d\u0020\u0073\u0063\u006fr\u0065\u0054\u0068\u0072\u0065\u0073h\u006f\u006c\u0064 \u0025\u0066", _abac, _aagcf, _ggdg, scoreThreshold)
	}
	return false, nil
}
func _ddcd(_daed, _dgab *Bitmap, _gffa, _gebf, _bfa, _gacd, _ffce, _gdfe, _bfae, _dgbba int, _bbe CombinationOperator, _fcdb int) error {
	var _fadce int
	_fedg := func() {
		_fadce++
		_bfa += _dgab.RowStride
		_gacd += _daed.RowStride
		_ffce += _daed.RowStride
	}
	for _fadce = _gffa; _fadce < _gebf; _fedg() {
		var _fcgg uint16
		_eeda := _bfa
		for _dcfd := _gacd; _dcfd <= _ffce; _dcfd++ {
			_fffb, _dea := _dgab.GetByte(_eeda)
			if _dea != nil {
				return _dea
			}
			_dgdd, _dea := _daed.GetByte(_dcfd)
			if _dea != nil {
				return _dea
			}
			_fcgg = (_fcgg | (uint16(_dgdd) & 0xff)) << uint(_dgbba)
			_dgdd = byte(_fcgg >> 8)
			if _dea = _dgab.SetByte(_eeda, _eeabd(_fffb, _dgdd, _bbe)); _dea != nil {
				return _dea
			}
			_eeda++
			_fcgg <<= uint(_bfae)
			if _dcfd == _ffce {
				_dgdd = byte(_fcgg >> (8 - uint8(_dgbba)))
				if _fcdb != 0 {
					_dgdd = _fcgbg(uint(8+_gdfe), _dgdd)
				}
				_fffb, _dea = _dgab.GetByte(_eeda)
				if _dea != nil {
					return _dea
				}
				if _dea = _dgab.SetByte(_eeda, _eeabd(_fffb, _dgdd, _bbe)); _dea != nil {
					return _dea
				}
			}
		}
	}
	return nil
}
func CorrelationScoreSimple(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_cbaf float64, _aff error) {
	const _fdbb = "\u0043\u006f\u0072\u0072el\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0053\u0069\u006d\u0070l\u0065"
	if bm1 == nil || bm2 == nil {
		return _cbaf, _e.Error(_fdbb, "n\u0069l\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0073 \u0070\u0072\u006f\u0076id\u0065\u0064")
	}
	if tab == nil {
		return _cbaf, _e.Error(_fdbb, "\u0074\u0061\u0062\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if area1 == 0 || area2 == 0 {
		return _cbaf, _e.Error(_fdbb, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0061\u0072e\u0061\u0073\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u0020\u0030")
	}
	_fagb, _cegf := bm1.Width, bm1.Height
	_fcda, _dcae := bm2.Width, bm2.Height
	if _ggce(_fagb-_fcda) > maxDiffW {
		return 0, nil
	}
	if _ggce(_cegf-_dcae) > maxDiffH {
		return 0, nil
	}
	var _ffceg, _eaad int
	if delX >= 0 {
		_ffceg = int(delX + 0.5)
	} else {
		_ffceg = int(delX - 0.5)
	}
	if delY >= 0 {
		_eaad = int(delY + 0.5)
	} else {
		_eaad = int(delY - 0.5)
	}
	_gdce := bm1.createTemplate()
	if _aff = _gdce.RasterOperation(_ffceg, _eaad, _fcda, _dcae, PixSrc, bm2, 0, 0); _aff != nil {
		return _cbaf, _e.Wrap(_aff, _fdbb, "\u0062m\u0032 \u0074\u006f\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _aff = _gdce.RasterOperation(0, 0, _fagb, _cegf, PixSrcAndDst, bm1, 0, 0); _aff != nil {
		return _cbaf, _e.Wrap(_aff, _fdbb, "b\u006d\u0031\u0020\u0061\u006e\u0064\u0020\u0062\u006d\u0054")
	}
	_bbdb := _gdce.countPixels()
	_cbaf = float64(_bbdb) * float64(_bbdb) / (float64(area1) * float64(area2))
	return _cbaf, nil
}
func _fgedg(_eecd *Bitmap, _cdba, _bfgaf int, _cgeg, _gfad int, _cccb RasterOperator, _dceb *Bitmap, _egab, _bece int) error {
	var _adca, _bbgc, _abcg, _adbe int
	if _cdba < 0 {
		_egab -= _cdba
		_cgeg += _cdba
		_cdba = 0
	}
	if _egab < 0 {
		_cdba -= _egab
		_cgeg += _egab
		_egab = 0
	}
	_adca = _cdba + _cgeg - _eecd.Width
	if _adca > 0 {
		_cgeg -= _adca
	}
	_bbgc = _egab + _cgeg - _dceb.Width
	if _bbgc > 0 {
		_cgeg -= _bbgc
	}
	if _bfgaf < 0 {
		_bece -= _bfgaf
		_gfad += _bfgaf
		_bfgaf = 0
	}
	if _bece < 0 {
		_bfgaf -= _bece
		_gfad += _bece
		_bece = 0
	}
	_abcg = _bfgaf + _gfad - _eecd.Height
	if _abcg > 0 {
		_gfad -= _abcg
	}
	_adbe = _bece + _gfad - _dceb.Height
	if _adbe > 0 {
		_gfad -= _adbe
	}
	if _cgeg <= 0 || _gfad <= 0 {
		return nil
	}
	var _cfcd error
	switch {
	case _cdba&7 == 0 && _egab&7 == 0:
		_cfcd = _adac(_eecd, _cdba, _bfgaf, _cgeg, _gfad, _cccb, _dceb, _egab, _bece)
	case _cdba&7 == _egab&7:
		_cfcd = _cfaba(_eecd, _cdba, _bfgaf, _cgeg, _gfad, _cccb, _dceb, _egab, _bece)
	default:
		_cfcd = _gdbca(_eecd, _cdba, _bfgaf, _cgeg, _gfad, _cccb, _dceb, _egab, _bece)
	}
	if _cfcd != nil {
		return _e.Wrap(_cfcd, "r\u0061\u0073\u0074\u0065\u0072\u004f\u0070\u004c\u006f\u0077", "")
	}
	return nil
}
func Centroid(bm *Bitmap, centTab, sumTab []int) (Point, error) { return bm.centroid(centTab, sumTab) }
func (_gdae *ClassedPoints) GroupByY() ([]*ClassedPoints, error) {
	const _efff = "\u0043\u006c\u0061\u0073se\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0072\u006f\u0075\u0070\u0042y\u0059"
	if _fgeff := _gdae.validateIntSlice(); _fgeff != nil {
		return nil, _e.Wrap(_fgeff, _efff, "")
	}
	if _gdae.IntSlice.Size() == 0 {
		return nil, _e.Error(_efff, "\u004e\u006f\u0020\u0063la\u0073\u0073\u0065\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_gdae.SortByY()
	var (
		_cgec []*ClassedPoints
		_geeg int
	)
	_gcbb := -1
	var _gffab *ClassedPoints
	for _eeef := 0; _eeef < len(_gdae.IntSlice); _eeef++ {
		_geeg = int(_gdae.YAtIndex(_eeef))
		if _geeg != _gcbb {
			_gffab = &ClassedPoints{Points: _gdae.Points}
			_gcbb = _geeg
			_cgec = append(_cgec, _gffab)
		}
		_gffab.IntSlice = append(_gffab.IntSlice, _gdae.IntSlice[_eeef])
	}
	for _, _efe := range _cgec {
		_efe.SortByX()
	}
	return _cgec, nil
}
func _fdeg(_cdd, _cebf, _acc *Bitmap) (*Bitmap, error) {
	const _gdaa = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0078\u006f\u0072"
	if _cebf == nil {
		return nil, _e.Error(_gdaa, "'\u0062\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _acc == nil {
		return nil, _e.Error(_gdaa, "'\u0062\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _cdd == _acc {
		return nil, _e.Error(_gdaa, "'\u0064\u0027\u0020\u003d\u003d\u0020\u0027\u0062\u0032\u0027")
	}
	if !_cebf.SizesEqual(_acc) {
		_ca.Log.Debug("\u0025s\u0020\u002d \u0042\u0069\u0074\u006da\u0070\u0020\u0027b\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074 e\u0071\u0075\u0061l\u0020\u0073i\u007a\u0065\u0020\u0077\u0069\u0074h\u0020\u0027b\u0032\u0027", _gdaa)
	}
	var _ddg error
	if _cdd, _ddg = _cdcf(_cdd, _cebf); _ddg != nil {
		return nil, _e.Wrap(_ddg, _gdaa, "\u0063\u0061n\u0027\u0074\u0020c\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027")
	}
	if _ddg = _cdd.RasterOperation(0, 0, _cdd.Width, _cdd.Height, PixSrcXorDst, _acc, 0, 0); _ddg != nil {
		return nil, _e.Wrap(_ddg, _gdaa, "")
	}
	return _cdd, nil
}
func (_dcg *Bitmap) AddBorder(borderSize, val int) (*Bitmap, error) {
	if borderSize == 0 {
		return _dcg.Copy(), nil
	}
	_cgf, _fga := _dcg.addBorderGeneral(borderSize, borderSize, borderSize, borderSize, val)
	if _fga != nil {
		return nil, _e.Wrap(_fga, "\u0041d\u0064\u0042\u006f\u0072\u0064\u0065r", "")
	}
	return _cgf, nil
}

type byHeight Bitmaps

func _dcdc(_ab *Bitmap, _gbg int, _agf []byte) (_ece *Bitmap, _bad error) {
	const _aaf = "\u0072\u0065\u0064\u0075\u0063\u0065\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0032"
	if _ab == nil {
		return nil, _e.Error(_aaf, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _gbg < 1 || _gbg > 4 {
		return nil, _e.Error(_aaf, "\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _ab.Height <= 1 {
		return nil, _e.Errorf(_aaf, "\u0073o\u0075\u0072c\u0065\u0020\u0068e\u0069\u0067\u0068\u0074\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061t\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027\u0020-\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", _ab.Height)
	}
	_ece = New(_ab.Width/2, _ab.Height/2)
	if _agf == nil {
		_agf = _bcc()
	}
	_gd := _gcg(_ab.RowStride, 2*_ece.RowStride)
	switch _gbg {
	case 1:
		_bad = _aada(_ab, _ece, _gbg, _agf, _gd)
	case 2:
		_bad = _dba(_ab, _ece, _gbg, _agf, _gd)
	case 3:
		_bad = _fbac(_ab, _ece, _gbg, _agf, _gd)
	case 4:
		_bad = _fcd(_ab, _ece, _gbg, _agf, _gd)
	}
	if _bad != nil {
		return nil, _bad
	}
	return _ece, nil
}
func (_dbeb *Bitmap) nextOnPixelLow(_caec, _cedf, _fbed, _aedc, _gce int) (_dga _aa.Point, _gaef bool, _agbc error) {
	const _bcdd = "B\u0069\u0074\u006d\u0061p.\u006ee\u0078\u0074\u004f\u006e\u0050i\u0078\u0065\u006c\u004c\u006f\u0077"
	var (
		_cdgc int
		_gcbg byte
	)
	_ggcd := _gce * _fbed
	_dcbb := _ggcd + (_aedc / 8)
	if _gcbg, _agbc = _dbeb.GetByte(_dcbb); _agbc != nil {
		return _dga, false, _e.Wrap(_agbc, _bcdd, "\u0078\u0053\u0074\u0061\u0072\u0074\u0020\u0061\u006e\u0064 \u0079\u0053\u0074\u0061\u0072\u0074\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	if _gcbg != 0 {
		_cdgb := _aedc - (_aedc % 8) + 7
		for _cdgc = _aedc; _cdgc <= _cdgb && _cdgc < _caec; _cdgc++ {
			if _dbeb.GetPixel(_cdgc, _gce) {
				_dga.X = _cdgc
				_dga.Y = _gce
				return _dga, true, nil
			}
		}
	}
	_caecc := (_aedc / 8) + 1
	_cdgc = 8 * _caecc
	var _agae int
	for _dcbb = _ggcd + _caecc; _cdgc < _caec; _dcbb, _cdgc = _dcbb+1, _cdgc+8 {
		if _gcbg, _agbc = _dbeb.GetByte(_dcbb); _agbc != nil {
			return _dga, false, _e.Wrap(_agbc, _bcdd, "r\u0065\u0073\u0074\u0020of\u0020t\u0068\u0065\u0020\u006c\u0069n\u0065\u0020\u0062\u0079\u0074\u0065")
		}
		if _gcbg == 0 {
			continue
		}
		for _agae = 0; _agae < 8 && _cdgc < _caec; _agae, _cdgc = _agae+1, _cdgc+1 {
			if _dbeb.GetPixel(_cdgc, _gce) {
				_dga.X = _cdgc
				_dga.Y = _gce
				return _dga, true, nil
			}
		}
	}
	for _bdff := _gce + 1; _bdff < _cedf; _bdff++ {
		_ggcd = _bdff * _fbed
		for _dcbb, _cdgc = _ggcd, 0; _cdgc < _caec; _dcbb, _cdgc = _dcbb+1, _cdgc+8 {
			if _gcbg, _agbc = _dbeb.GetByte(_dcbb); _agbc != nil {
				return _dga, false, _e.Wrap(_agbc, _bcdd, "\u0066o\u006cl\u006f\u0077\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u0073")
			}
			if _gcbg == 0 {
				continue
			}
			for _agae = 0; _agae < 8 && _cdgc < _caec; _agae, _cdgc = _agae+1, _cdgc+1 {
				if _dbeb.GetPixel(_cdgc, _bdff) {
					_dga.X = _cdgc
					_dga.Y = _bdff
					return _dga, true, nil
				}
			}
		}
	}
	return _dga, false, nil
}
func (_bege *Bitmap) countPixels() int {
	var (
		_cgga int
		_ccg  uint8
		_cbfe byte
		_fbbb int
	)
	_dag := _bege.RowStride
	_cbbc := uint(_bege.Width & 0x07)
	if _cbbc != 0 {
		_ccg = uint8((0xff << (8 - _cbbc)) & 0xff)
		_dag--
	}
	for _gaed := 0; _gaed < _bege.Height; _gaed++ {
		for _fbbb = 0; _fbbb < _dag; _fbbb++ {
			_cbfe = _bege.Data[_gaed*_bege.RowStride+_fbbb]
			_cgga += int(_fde[_cbfe])
		}
		if _cbbc != 0 {
			_cgga += int(_fde[_bege.Data[_gaed*_bege.RowStride+_fbbb]&_ccg])
		}
	}
	return _cgga
}
func TstImageBitmap() *Bitmap { return _gabd.Copy() }
func Rect(x, y, w, h int) (*_aa.Rectangle, error) {
	const _fbaf = "b\u0069\u0074\u006d\u0061\u0070\u002e\u0052\u0065\u0063\u0074"
	if x < 0 {
		w += x
		x = 0
		if w <= 0 {
			return nil, _e.Errorf(_fbaf, "x\u003a\u0027\u0025\u0064\u0027\u0020<\u0020\u0030\u0020\u0061\u006e\u0064\u0020\u0077\u003a \u0027\u0025\u0064'\u0020<\u003d\u0020\u0030", x, w)
		}
	}
	if y < 0 {
		h += y
		y = 0
		if h <= 0 {
			return nil, _e.Error(_fbaf, "\u0079\u0020\u003c 0\u0020\u0061\u006e\u0064\u0020\u0062\u006f\u0078\u0020\u006f\u0066\u0066\u0020\u002b\u0071\u0075\u0061\u0064")
		}
	}
	_eaaa := _aa.Rect(x, y, x+w, y+h)
	return &_eaaa, nil
}

const (
	_ SizeComparison = iota
	SizeSelectIfLT
	SizeSelectIfGT
	SizeSelectIfLTE
	SizeSelectIfGTE
	SizeSelectIfEQ
)

func TstFrameBitmap() *Bitmap { return _cbff.Copy() }

type Boxes []*_aa.Rectangle
type ClassedPoints struct {
	*Points
	_ac.IntSlice
	_acab func(_abadf, _cgfg int) bool
}

func RasterOperation(dest *Bitmap, dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _gegg(dest, dx, dy, dw, dh, op, src, sx, sy)
}
func (_bedcg *Bitmap) connComponentsBitmapsBB(_cfda *Bitmaps, _cfggd int) (_ccf *Boxes, _egfd error) {
	const _gcba = "\u0063\u006f\u006enC\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0069\u0074\u006d\u0061\u0070\u0073\u0042\u0042"
	if _cfggd != 4 && _cfggd != 8 {
		return nil, _e.Error(_gcba, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _cfda == nil {
		return nil, _e.Error(_gcba, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if len(_cfda.Values) > 0 {
		return nil, _e.Error(_gcba, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u006fn\u002d\u0065\u006d\u0070\u0074\u0079\u0020\u0042\u0069\u0074m\u0061\u0070\u0073")
	}
	if _bedcg.Zero() {
		return &Boxes{}, nil
	}
	var (
		_bbfd, _geef, _ddfe, _eadd *Bitmap
	)
	_bedcg.setPadBits(0)
	if _bbfd, _egfd = _cdcf(nil, _bedcg); _egfd != nil {
		return nil, _e.Wrap(_egfd, _gcba, "\u0062\u006d\u0031")
	}
	if _geef, _egfd = _cdcf(nil, _bedcg); _egfd != nil {
		return nil, _e.Wrap(_egfd, _gcba, "\u0062\u006d\u0032")
	}
	_dca := &_ac.Stack{}
	_dca.Aux = &_ac.Stack{}
	_ccf = &Boxes{}
	var (
		_cdbc, _cgb int
		_gafb       _aa.Point
		_daca       bool
		_afec       *_aa.Rectangle
	)
	for {
		if _gafb, _daca, _egfd = _bbfd.nextOnPixel(_cdbc, _cgb); _egfd != nil {
			return nil, _e.Wrap(_egfd, _gcba, "")
		}
		if !_daca {
			break
		}
		if _afec, _egfd = _beaad(_bbfd, _dca, _gafb.X, _gafb.Y, _cfggd); _egfd != nil {
			return nil, _e.Wrap(_egfd, _gcba, "")
		}
		if _egfd = _ccf.Add(_afec); _egfd != nil {
			return nil, _e.Wrap(_egfd, _gcba, "")
		}
		if _ddfe, _egfd = _bbfd.clipRectangle(_afec, nil); _egfd != nil {
			return nil, _e.Wrap(_egfd, _gcba, "\u0062\u006d\u0033")
		}
		if _eadd, _egfd = _geef.clipRectangle(_afec, nil); _egfd != nil {
			return nil, _e.Wrap(_egfd, _gcba, "\u0062\u006d\u0034")
		}
		if _, _egfd = _fdeg(_ddfe, _ddfe, _eadd); _egfd != nil {
			return nil, _e.Wrap(_egfd, _gcba, "\u0062m\u0033\u0020\u005e\u0020\u0062\u006d4")
		}
		if _egfd = _geef.RasterOperation(_afec.Min.X, _afec.Min.Y, _afec.Dx(), _afec.Dy(), PixSrcXorDst, _ddfe, 0, 0); _egfd != nil {
			return nil, _e.Wrap(_egfd, _gcba, "\u0062\u006d\u0032\u0020\u002d\u0058\u004f\u0052\u002d>\u0020\u0062\u006d\u0033")
		}
		_cfda.AddBitmap(_ddfe)
		_cdbc = _gafb.X
		_cgb = _gafb.Y
	}
	_cfda.Boxes = *_ccf
	return _ccf, nil
}
func (_agfg *Bitmap) ClipRectangle(box *_aa.Rectangle) (_fad *Bitmap, _agb *_aa.Rectangle, _adec error) {
	const _beda = "\u0043\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, nil, _e.Error(_beda, "\u0062o\u0078 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	_ggbe, _dbc := _agfg.Width, _agfg.Height
	_cad := _aa.Rect(0, 0, _ggbe, _dbc)
	if !box.Overlaps(_cad) {
		return nil, nil, _e.Error(_beda, "b\u006f\u0078\u0020\u0064oe\u0073n\u0027\u0074\u0020\u006f\u0076e\u0072\u006c\u0061\u0070\u0020\u0062")
	}
	_gfa := box.Intersect(_cad)
	_dgfb, _cbe := _gfa.Min.X, _gfa.Min.Y
	_defc, _daab := _gfa.Dx(), _gfa.Dy()
	_fad = New(_defc, _daab)
	_fad.Text = _agfg.Text
	if _adec = _fad.RasterOperation(0, 0, _defc, _daab, PixSrc, _agfg, _dgfb, _cbe); _adec != nil {
		return nil, nil, _e.Wrap(_adec, _beda, "\u0050\u0069\u0078\u0053\u0072\u0063\u0020\u0074\u006f\u0020\u0063\u006ci\u0070\u0070\u0065\u0064")
	}
	_agb = &_gfa
	return _fad, _agb, nil
}

type MorphOperation int

func (_cecga *Bitmaps) selectByIndexes(_caded []int) (*Bitmaps, error) {
	_debg := &Bitmaps{}
	for _, _gbge := range _caded {
		_bdefb, _aabf := _cecga.GetBitmap(_gbge)
		if _aabf != nil {
			return nil, _e.Wrap(_aabf, "\u0073e\u006ce\u0063\u0074\u0042\u0079\u0049\u006e\u0064\u0065\u0078\u0065\u0073", "")
		}
		_debg.AddBitmap(_bdefb)
	}
	return _debg, nil
}
func MorphSequence(src *Bitmap, sequence ...MorphProcess) (*Bitmap, error) {
	return _ecbgb(src, sequence...)
}

type BitmapsArray struct {
	Values []*Bitmaps
	Boxes  []*_aa.Rectangle
}

func _egbea(_cccbf *Bitmap, _facbd *_ac.Stack, _eadce, _fefc int) (_ebfd *_aa.Rectangle, _dgbg error) {
	const _aaaf = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _cccbf == nil {
		return nil, _e.Error(_aaaf, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _facbd == nil {
		return nil, _e.Error(_aaaf, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_acca, _ecff := _cccbf.Width, _cccbf.Height
	_caabg := _acca - 1
	_gceb := _ecff - 1
	if _eadce < 0 || _eadce > _caabg || _fefc < 0 || _fefc > _gceb || !_cccbf.GetPixel(_eadce, _fefc) {
		return nil, nil
	}
	_ddeg := _aa.Rect(100000, 100000, 0, 0)
	if _dgbg = _agadg(_facbd, _eadce, _eadce, _fefc, 1, _gceb, &_ddeg); _dgbg != nil {
		return nil, _e.Wrap(_dgbg, _aaaf, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _dgbg = _agadg(_facbd, _eadce, _eadce, _fefc+1, -1, _gceb, &_ddeg); _dgbg != nil {
		return nil, _e.Wrap(_dgbg, _aaaf, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_ddeg.Min.X, _ddeg.Max.X = _eadce, _eadce
	_ddeg.Min.Y, _ddeg.Max.Y = _fefc, _fefc
	var (
		_fadd *fillSegment
		_aeab int
	)
	for _facbd.Len() > 0 {
		if _fadd, _dgbg = _cfee(_facbd); _dgbg != nil {
			return nil, _e.Wrap(_dgbg, _aaaf, "")
		}
		_fefc = _fadd._ffeg
		for _eadce = _fadd._abbc - 1; _eadce >= 0 && _cccbf.GetPixel(_eadce, _fefc); _eadce-- {
			if _dgbg = _cccbf.SetPixel(_eadce, _fefc, 0); _dgbg != nil {
				return nil, _e.Wrap(_dgbg, _aaaf, "\u0031s\u0074\u0020\u0073\u0065\u0074")
			}
		}
		if _eadce >= _fadd._abbc-1 {
			for {
				for _eadce++; _eadce <= _fadd._bcaf+1 && _eadce <= _caabg && !_cccbf.GetPixel(_eadce, _fefc); _eadce++ {
				}
				_aeab = _eadce
				if !(_eadce <= _fadd._bcaf+1 && _eadce <= _caabg) {
					break
				}
				for ; _eadce <= _caabg && _cccbf.GetPixel(_eadce, _fefc); _eadce++ {
					if _dgbg = _cccbf.SetPixel(_eadce, _fefc, 0); _dgbg != nil {
						return nil, _e.Wrap(_dgbg, _aaaf, "\u0032n\u0064\u0020\u0073\u0065\u0074")
					}
				}
				if _dgbg = _agadg(_facbd, _aeab, _eadce-1, _fadd._ffeg, _fadd._bedfb, _gceb, &_ddeg); _dgbg != nil {
					return nil, _e.Wrap(_dgbg, _aaaf, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
				}
				if _eadce > _fadd._bcaf {
					if _dgbg = _agadg(_facbd, _fadd._bcaf+1, _eadce-1, _fadd._ffeg, -_fadd._bedfb, _gceb, &_ddeg); _dgbg != nil {
						return nil, _e.Wrap(_dgbg, _aaaf, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
					}
				}
			}
			continue
		}
		_aeab = _eadce + 1
		if _aeab < _fadd._abbc {
			if _dgbg = _agadg(_facbd, _aeab, _fadd._abbc-1, _fadd._ffeg, -_fadd._bedfb, _gceb, &_ddeg); _dgbg != nil {
				return nil, _e.Wrap(_dgbg, _aaaf, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
			}
		}
		_eadce = _fadd._abbc
		for {
			for ; _eadce <= _caabg && _cccbf.GetPixel(_eadce, _fefc); _eadce++ {
				if _dgbg = _cccbf.SetPixel(_eadce, _fefc, 0); _dgbg != nil {
					return nil, _e.Wrap(_dgbg, _aaaf, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _dgbg = _agadg(_facbd, _aeab, _eadce-1, _fadd._ffeg, _fadd._bedfb, _gceb, &_ddeg); _dgbg != nil {
				return nil, _e.Wrap(_dgbg, _aaaf, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _eadce > _fadd._bcaf {
				if _dgbg = _agadg(_facbd, _fadd._bcaf+1, _eadce-1, _fadd._ffeg, -_fadd._bedfb, _gceb, &_ddeg); _dgbg != nil {
					return nil, _e.Wrap(_dgbg, _aaaf, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _eadce++; _eadce <= _fadd._bcaf+1 && _eadce <= _caabg && !_cccbf.GetPixel(_eadce, _fefc); _eadce++ {
			}
			_aeab = _eadce
			if !(_eadce <= _fadd._bcaf+1 && _eadce <= _caabg) {
				break
			}
		}
	}
	_ddeg.Max.X++
	_ddeg.Max.Y++
	return &_ddeg, nil
}
func (_gada Points) Size() int { return len(_gada) }
func _dba(_bfg, _bec *Bitmap, _fffd int, _eafc []byte, _gbgc int) (_cf error) {
	const _edd = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0032"
	var (
		_dfg, _bb, _bg, _cff, _abd, _daa, _dae, _def int
		_ggc, _gea, _cegd, _beb                      uint32
		_ggfe, _gec                                  byte
		_aba                                         uint16
	)
	_ffg := make([]byte, 4)
	_gbgd := make([]byte, 4)
	for _bg = 0; _bg < _bfg.Height-1; _bg, _cff = _bg+2, _cff+1 {
		_dfg = _bg * _bfg.RowStride
		_bb = _cff * _bec.RowStride
		for _abd, _daa = 0, 0; _abd < _gbgc; _abd, _daa = _abd+4, _daa+1 {
			for _dae = 0; _dae < 4; _dae++ {
				_def = _dfg + _abd + _dae
				if _def <= len(_bfg.Data)-1 && _def < _dfg+_bfg.RowStride {
					_ffg[_dae] = _bfg.Data[_def]
				} else {
					_ffg[_dae] = 0x00
				}
				_def = _dfg + _bfg.RowStride + _abd + _dae
				if _def <= len(_bfg.Data)-1 && _def < _dfg+(2*_bfg.RowStride) {
					_gbgd[_dae] = _bfg.Data[_def]
				} else {
					_gbgd[_dae] = 0x00
				}
			}
			_ggc = _db.BigEndian.Uint32(_ffg)
			_gea = _db.BigEndian.Uint32(_gbgd)
			_cegd = _ggc & _gea
			_cegd |= _cegd << 1
			_beb = _ggc | _gea
			_beb &= _beb << 1
			_gea = _cegd | _beb
			_gea &= 0xaaaaaaaa
			_ggc = _gea | (_gea << 7)
			_ggfe = byte(_ggc >> 24)
			_gec = byte((_ggc >> 8) & 0xff)
			_def = _bb + _daa
			if _def+1 == len(_bec.Data)-1 || _def+1 >= _bb+_bec.RowStride {
				if _cf = _bec.SetByte(_def, _eafc[_ggfe]); _cf != nil {
					return _e.Wrapf(_cf, _edd, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _def)
				}
			} else {
				_aba = (uint16(_eafc[_ggfe]) << 8) | uint16(_eafc[_gec])
				if _cf = _bec.setTwoBytes(_def, _aba); _cf != nil {
					return _e.Wrapf(_cf, _edd, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _def)
				}
				_daa++
			}
		}
	}
	return nil
}

const (
	_dgfe shift = iota
	_ecde
)

func _agadg(_gfbd *_ac.Stack, _babg, _acdd, _fgca, _cgfde, _cfegb int, _febd *_aa.Rectangle) (_bgfb error) {
	const _eaade = "\u0070\u0075\u0073\u0068\u0046\u0069\u006c\u006c\u0053\u0065\u0067m\u0065\u006e\u0074\u0042\u006f\u0075\u006e\u0064\u0069\u006eg\u0042\u006f\u0078"
	if _gfbd == nil {
		return _e.Error(_eaade, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _febd == nil {
		return _e.Error(_eaade, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0069\u006da\u0067e\u002e\u0052\u0065\u0063\u0074\u0061\u006eg\u006c\u0065")
	}
	_febd.Min.X = _ac.Min(_febd.Min.X, _babg)
	_febd.Max.X = _ac.Max(_febd.Max.X, _acdd)
	_febd.Min.Y = _ac.Min(_febd.Min.Y, _fgca)
	_febd.Max.Y = _ac.Max(_febd.Max.Y, _fgca)
	if !(_fgca+_cgfde >= 0 && _fgca+_cgfde <= _cfegb) {
		return nil
	}
	if _gfbd.Aux == nil {
		return _e.Error(_eaade, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	var _bbgg *fillSegment
	_bgdae, _aeff := _gfbd.Aux.Pop()
	if _aeff {
		if _bbgg, _aeff = _bgdae.(*fillSegment); !_aeff {
			return _e.Error(_eaade, "a\u0075\u0078\u0053\u0074\u0061\u0063k\u0020\u0064\u0061\u0074\u0061\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061 \u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065n\u0074")
		}
	} else {
		_bbgg = &fillSegment{}
	}
	_bbgg._abbc = _babg
	_bbgg._bcaf = _acdd
	_bbgg._ffeg = _fgca
	_bbgg._bedfb = _cgfde
	_gfbd.Push(_bbgg)
	return nil
}
func (_fdgbg *Bitmaps) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_eded *Bitmaps, _eaed error) {
	const _gdbe = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0053\u0065\u006ce\u0063\u0074\u0042\u0079Si\u007a\u0065"
	if _fdgbg == nil {
		return nil, _e.Error(_gdbe, "\u0027\u0062\u0027 B\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _e.Errorf(_gdbe, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _e.Errorf(_gdbe, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", relation)
	}
	_ccag, _eaed := _fdgbg.makeSizeIndicator(width, height, tp, relation)
	if _eaed != nil {
		return nil, _e.Wrap(_eaed, _gdbe, "")
	}
	_eded, _eaed = _fdgbg.selectByIndicator(_ccag)
	if _eaed != nil {
		return nil, _e.Wrap(_eaed, _gdbe, "")
	}
	return _eded, nil
}
func (_ecbc *Boxes) selectWithIndicator(_bcfe *_ac.NumSlice) (_agfe *Boxes, _fbbfb error) {
	const _eddd = "\u0042o\u0078\u0065\u0073\u002es\u0065\u006c\u0065\u0063\u0074W\u0069t\u0068I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _ecbc == nil {
		return nil, _e.Error(_eddd, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _bcfe == nil {
		return nil, _e.Error(_eddd, "\u0027\u006ea\u0027\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_bcfe) != len(*_ecbc) {
		return nil, _e.Error(_eddd, "\u0062\u006f\u0078\u0065\u0073\u0020\u0027\u0062\u0027\u0020\u0068\u0061\u0073\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020s\u0069\u007a\u0065\u0020\u0074h\u0061\u006e \u0027\u006e\u0061\u0027")
	}
	var _abcd, _dcfc int
	for _bfgc := 0; _bfgc < len(*_bcfe); _bfgc++ {
		if _abcd, _fbbfb = _bcfe.GetInt(_bfgc); _fbbfb != nil {
			return nil, _e.Wrap(_fbbfb, _eddd, "\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u0067\u0020c\u006f\u0075\u006e\u0074")
		}
		if _abcd == 1 {
			_dcfc++
		}
	}
	if _dcfc == len(*_ecbc) {
		return _ecbc, nil
	}
	_dad := Boxes{}
	for _deee := 0; _deee < len(*_bcfe); _deee++ {
		_abcd = int((*_bcfe)[_deee])
		if _abcd == 0 {
			continue
		}
		_dad = append(_dad, (*_ecbc)[_deee])
	}
	_agfe = &_dad
	return _agfe, nil
}
func (_gdg *Bitmap) GetPixel(x, y int) bool {
	_fab := _gdg.GetByteIndex(x, y)
	_becc := _gdg.GetBitOffset(x)
	_fbf := uint(7 - _becc)
	if _fab > len(_gdg.Data)-1 {
		_ca.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0072\u0061\u006e\u0067\u0065\u002e \u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0027\u0025\u0064'\u002c\u0020\u0062m\u003a\u0020\u0027\u0025\u0073\u0027", x, y, _gdg)
		return false
	}
	if (_gdg.Data[_fab]>>_fbf)&0x01 >= 1 {
		return true
	}
	return false
}
func (_aafgf *Bitmap) nextOnPixel(_fgg, _fbgdb int) (_dfde _aa.Point, _fddc bool, _ddc error) {
	const _befg = "n\u0065\u0078\u0074\u004f\u006e\u0050\u0069\u0078\u0065\u006c"
	_dfde, _fddc, _ddc = _aafgf.nextOnPixelLow(_aafgf.Width, _aafgf.Height, _aafgf.RowStride, _fgg, _fbgdb)
	if _ddc != nil {
		return _dfde, false, _e.Wrap(_ddc, _befg, "")
	}
	return _dfde, _fddc, nil
}

const (
	CmbOpOr CombinationOperator = iota
	CmbOpAnd
	CmbOpXor
	CmbOpXNor
	CmbOpReplace
	CmbOpNot
)

var _ _ff.Interface = &ClassedPoints{}

func _fcgbg(_dbaf uint, _fabd byte) byte { return _fabd >> _dbaf << _dbaf }
func TstFrameBitmapData() []byte         { return _cbff.Data }
func (_eeac *Points) Add(pt *Points) error {
	const _dgeb = "\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0041\u0064\u0064"
	if _eeac == nil {
		return _e.Error(_dgeb, "\u0070o\u0069n\u0074\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if pt == nil {
		return _e.Error(_dgeb, "a\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	*_eeac = append(*_eeac, *pt...)
	return nil
}
func (_afb *Bitmap) SetPadBits(value int) { _afb.setPadBits(value) }

type Color int

func _edcf(_ebfb *Bitmap, _bfde, _gaa, _gcgf, _afbed int, _fccf RasterOperator) {
	if _bfde < 0 {
		_gcgf += _bfde
		_bfde = 0
	}
	_dafb := _bfde + _gcgf - _ebfb.Width
	if _dafb > 0 {
		_gcgf -= _dafb
	}
	if _gaa < 0 {
		_afbed += _gaa
		_gaa = 0
	}
	_gfabb := _gaa + _afbed - _ebfb.Height
	if _gfabb > 0 {
		_afbed -= _gfabb
	}
	if _gcgf <= 0 || _afbed <= 0 {
		return
	}
	if (_bfde & 7) == 0 {
		_acbb(_ebfb, _bfde, _gaa, _gcgf, _afbed, _fccf)
	} else {
		_bbfda(_ebfb, _bfde, _gaa, _gcgf, _afbed, _fccf)
	}
}
func Dilate(d *Bitmap, s *Bitmap, sel *Selection) (*Bitmap, error) { return _caagg(d, s, sel) }
func _eeege(_gbec ...MorphProcess) (_dfba error) {
	const _fcdc = "v\u0065r\u0069\u0066\u0079\u004d\u006f\u0072\u0070\u0068P\u0072\u006f\u0063\u0065ss\u0065\u0073"
	var _eegf, _fdgb int
	for _cbeg, _aaca := range _gbec {
		if _dfba = _aaca.verify(_cbeg, &_eegf, &_fdgb); _dfba != nil {
			return _e.Wrap(_dfba, _fcdc, "")
		}
	}
	if _fdgb != 0 && _eegf != 0 {
		return _e.Error(_fcdc, "\u004d\u006f\u0072\u0070\u0068\u0020\u0073\u0065\u0071\u0075\u0065n\u0063\u0065\u0020\u002d\u0020\u0062\u006f\u0072d\u0065r\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u0065\u0074\u0020\u0072\u0065\u0064u\u0063\u0074\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0030")
	}
	return nil
}
func _cbdc(_cbeba, _acfg *Bitmap, _caed *Selection) (*Bitmap, error) {
	const _adggf = "\u006f\u0070\u0065\u006e"
	var _bdg error
	_cbeba, _bdg = _gebfg(_cbeba, _acfg, _caed)
	if _bdg != nil {
		return nil, _e.Wrap(_bdg, _adggf, "")
	}
	_cfbg, _bdg := _cfgc(nil, _acfg, _caed)
	if _bdg != nil {
		return nil, _e.Wrap(_bdg, _adggf, "")
	}
	_, _bdg = _caagg(_cbeba, _cfbg, _caed)
	if _bdg != nil {
		return nil, _e.Wrap(_bdg, _adggf, "")
	}
	return _cbeba, nil
}

type Point struct{ X, Y float32 }

func (_geagb *Bitmaps) SortByHeight() {
	_ebbbd := (*byHeight)(_geagb)
	_ff.Sort(_ebbbd)
}
func (_eff *Bitmap) AddBorderGeneral(left, right, top, bot int, val int) (*Bitmap, error) {
	return _eff.addBorderGeneral(left, right, top, bot, val)
}
func (_ffea *Bitmaps) makeSizeIndicator(_fcag, _ccfc int, _bcbbec LocationFilter, _ggaf SizeComparison) (_beedb *_ac.NumSlice, _egeb error) {
	const _dgebc = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u006d\u0061\u006b\u0065S\u0069z\u0065I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _ffea == nil {
		return nil, _e.Error(_dgebc, "\u0062\u0069\u0074ma\u0070\u0073\u0020\u0027\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch _bcbbec {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _e.Errorf(_dgebc, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", _bcbbec)
	}
	switch _ggaf {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _e.Errorf(_dgebc, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", _ggaf)
	}
	_beedb = &_ac.NumSlice{}
	var (
		_cccc, _ggba, _dddb int
		_gggc               *Bitmap
	)
	for _, _gggc = range _ffea.Values {
		_cccc = 0
		_ggba, _dddb = _gggc.Width, _gggc.Height
		switch _bcbbec {
		case LocSelectWidth:
			if (_ggaf == SizeSelectIfLT && _ggba < _fcag) || (_ggaf == SizeSelectIfGT && _ggba > _fcag) || (_ggaf == SizeSelectIfLTE && _ggba <= _fcag) || (_ggaf == SizeSelectIfGTE && _ggba >= _fcag) || (_ggaf == SizeSelectIfEQ && _ggba == _fcag) {
				_cccc = 1
			}
		case LocSelectHeight:
			if (_ggaf == SizeSelectIfLT && _dddb < _ccfc) || (_ggaf == SizeSelectIfGT && _dddb > _ccfc) || (_ggaf == SizeSelectIfLTE && _dddb <= _ccfc) || (_ggaf == SizeSelectIfGTE && _dddb >= _ccfc) || (_ggaf == SizeSelectIfEQ && _dddb == _ccfc) {
				_cccc = 1
			}
		case LocSelectIfEither:
			if (_ggaf == SizeSelectIfLT && (_ggba < _fcag || _dddb < _ccfc)) || (_ggaf == SizeSelectIfGT && (_ggba > _fcag || _dddb > _ccfc)) || (_ggaf == SizeSelectIfLTE && (_ggba <= _fcag || _dddb <= _ccfc)) || (_ggaf == SizeSelectIfGTE && (_ggba >= _fcag || _dddb >= _ccfc)) || (_ggaf == SizeSelectIfEQ && (_ggba == _fcag || _dddb == _ccfc)) {
				_cccc = 1
			}
		case LocSelectIfBoth:
			if (_ggaf == SizeSelectIfLT && (_ggba < _fcag && _dddb < _ccfc)) || (_ggaf == SizeSelectIfGT && (_ggba > _fcag && _dddb > _ccfc)) || (_ggaf == SizeSelectIfLTE && (_ggba <= _fcag && _dddb <= _ccfc)) || (_ggaf == SizeSelectIfGTE && (_ggba >= _fcag && _dddb >= _ccfc)) || (_ggaf == SizeSelectIfEQ && (_ggba == _fcag && _dddb == _ccfc)) {
				_cccc = 1
			}
		}
		_beedb.AddInt(_cccc)
	}
	return _beedb, nil
}
func TstRSymbol(t *_f.T, scale ...int) *Bitmap {
	_baba, _cceg := NewWithData(4, 5, []byte{0xF0, 0x90, 0xF0, 0xA0, 0x90})
	_d.NoError(t, _cceg)
	return TstGetScaledSymbol(t, _baba, scale...)
}
func _eeabd(_bffa, _gabe byte, _afba CombinationOperator) byte {
	switch _afba {
	case CmbOpOr:
		return _gabe | _bffa
	case CmbOpAnd:
		return _gabe & _bffa
	case CmbOpXor:
		return _gabe ^ _bffa
	case CmbOpXNor:
		return ^(_gabe ^ _bffa)
	case CmbOpNot:
		return ^(_gabe)
	default:
		return _gabe
	}
}
func (_cbafc *Bitmaps) CountPixels() *_ac.NumSlice {
	_fbbef := &_ac.NumSlice{}
	for _, _fggg := range _cbafc.Values {
		_fbbef.AddInt(_fggg.CountPixels())
	}
	return _fbbef
}
func Blit(src *Bitmap, dst *Bitmap, x, y int, op CombinationOperator) error {
	var _gdaf, _eag int
	_gbd := src.RowStride - 1
	if x < 0 {
		_eag = -x
		x = 0
	} else if x+src.Width > dst.Width {
		_gbd -= src.Width + x - dst.Width
	}
	if y < 0 {
		_gdaf = -y
		y = 0
		_eag += src.RowStride
		_gbd += src.RowStride
	} else if y+src.Height > dst.Height {
		_gdaf = src.Height + y - dst.Height
	}
	var (
		_feca int
		_abb  error
	)
	_cfbe := x & 0x07
	_fffdf := 8 - _cfbe
	_fffe := src.Width & 0x07
	_ceba := _fffdf - _fffe
	_decf := _fffdf&0x07 != 0
	_aedg := src.Width <= ((_gbd-_eag)<<3)+_fffdf
	_bbf := dst.GetByteIndex(x, y)
	_bfff := _gdaf + dst.Height
	if src.Height > _bfff {
		_feca = _bfff
	} else {
		_feca = src.Height
	}
	switch {
	case !_decf:
		_abb = _ceec(src, dst, _gdaf, _feca, _bbf, _eag, _gbd, op)
	case _aedg:
		_abb = _bac(src, dst, _gdaf, _feca, _bbf, _eag, _gbd, _ceba, _cfbe, _fffdf, op)
	default:
		_abb = _ddcd(src, dst, _gdaf, _feca, _bbf, _eag, _gbd, _ceba, _cfbe, _fffdf, op, _fffe)
	}
	return _abb
}
func (_cdcc *Bitmap) RemoveBorderGeneral(left, right, top, bot int) (*Bitmap, error) {
	return _cdcc.removeBorderGeneral(left, right, top, bot)
}
func (_ecgga *ClassedPoints) validateIntSlice() error {
	const _cfbgb = "\u0076\u0061l\u0069\u0064\u0061t\u0065\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065"
	for _, _ggdff := range _ecgga.IntSlice {
		if _ggdff >= (_ecgga.Points.Size()) {
			return _e.Errorf(_cfbgb, "c\u006c\u0061\u0073\u0073\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006ci\u0064 \u0069\u006e\u0064\u0065x\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0064", _ggdff, _ecgga.Points.Size())
		}
	}
	return nil
}
func (_bbdf *Bitmap) GetByte(index int) (byte, error) {
	if index > len(_bbdf.Data)-1 || index < 0 {
		return 0, _e.Errorf("\u0047e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _bbdf.Data[index], nil
}

type byWidth Bitmaps

func _eced(_gff, _eed int) *Bitmap {
	return &Bitmap{Width: _gff, Height: _eed, RowStride: (_gff + 7) >> 3}
}

const (
	SelDontCare SelectionValue = iota
	SelHit
	SelMiss
)
const (
	AsymmetricMorphBC BoundaryCondition = iota
	SymmetricMorphBC
)

func (_gaaf *Bitmaps) SelectByIndexes(idx []int) (*Bitmaps, error) {
	const _bcad = "B\u0069\u0074\u006d\u0061\u0070\u0073.\u0053\u006f\u0072\u0074\u0049\u006e\u0064\u0065\u0078e\u0073\u0042\u0079H\u0065i\u0067\u0068\u0074"
	_abdgb, _dcdb := _gaaf.selectByIndexes(idx)
	if _dcdb != nil {
		return nil, _e.Wrap(_dcdb, _bcad, "")
	}
	return _abdgb, nil
}
func (_gcadd *Bitmaps) AddBox(box *_aa.Rectangle) { _gcadd.Boxes = append(_gcadd.Boxes, box) }

type Bitmaps struct {
	Values []*Bitmap
	Boxes  []*_aa.Rectangle
}

func _gebfg(_ceece, _dega *Bitmap, _eedg *Selection) (*Bitmap, error) {
	const _decfb = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0032"
	var _bfcc, _geaa int
	if _dega == nil {
		return nil, _e.Error(_decfb, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _eedg == nil {
		return nil, _e.Error(_decfb, "\u0073e\u006c \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_bfcc = _eedg.Width
	_geaa = _eedg.Height
	if _bfcc == 0 || _geaa == 0 {
		return nil, _e.Error(_decfb, "\u0073\u0065\u006c\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _ceece == nil {
		return _dega.createTemplate(), nil
	}
	if _cbeea := _ceece.resizeImageData(_dega); _cbeea != nil {
		return nil, _cbeea
	}
	return _ceece, nil
}
func _ffb(_ageb, _gefb *Bitmap, _cefa CombinationOperator) *Bitmap {
	_afgfe := New(_ageb.Width, _ageb.Height)
	for _gedf := 0; _gedf < len(_afgfe.Data); _gedf++ {
		_afgfe.Data[_gedf] = _eeabd(_ageb.Data[_gedf], _gefb.Data[_gedf], _cefa)
	}
	return _afgfe
}
func TstNSymbol(t *_f.T, scale ...int) *Bitmap {
	_dgcg, _cdbaa := NewWithData(4, 5, []byte{0x90, 0xD0, 0xB0, 0x90, 0x90})
	_d.NoError(t, _cdbaa)
	return TstGetScaledSymbol(t, _dgcg, scale...)
}
func SelCreateBrick(h, w int, cy, cx int, tp SelectionValue) *Selection {
	_ffdb := _ggdga(h, w, "")
	_ffdb.setOrigin(cy, cx)
	var _aaad, _addbb int
	for _aaad = 0; _aaad < h; _aaad++ {
		for _addbb = 0; _addbb < w; _addbb++ {
			_ffdb.Data[_aaad][_addbb] = tp
		}
	}
	return _ffdb
}
func TstImageBitmapInverseData() []byte {
	_efab := _gabd.Copy()
	_efab.InverseData()
	return _efab.Data
}
func (_bbag *Bitmaps) GroupByHeight() (*BitmapsArray, error) {
	const _fcdcc = "\u0047\u0072\u006f\u0075\u0070\u0042\u0079\u0048\u0065\u0069\u0067\u0068\u0074"
	if len(_bbag.Values) == 0 {
		return nil, _e.Error(_fcdcc, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_bafg := &BitmapsArray{}
	_bbag.SortByHeight()
	_ebeg := -1
	_fecb := -1
	for _eegfb := 0; _eegfb < len(_bbag.Values); _eegfb++ {
		_cbdg := _bbag.Values[_eegfb].Height
		if _cbdg > _ebeg {
			_ebeg = _cbdg
			_fecb++
			_bafg.Values = append(_bafg.Values, &Bitmaps{})
		}
		_bafg.Values[_fecb].AddBitmap(_bbag.Values[_eegfb])
	}
	return _bafg, nil
}
func (_gfc *Bitmap) addPadBits() (_dbcd error) {
	const _fbge = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0061\u0064\u0064\u0050\u0061d\u0042\u0069\u0074\u0073"
	_effc := _gfc.Width % 8
	if _effc == 0 {
		return nil
	}
	_bedc := _gfc.Width / 8
	_fdac := _ce.NewReader(_gfc.Data)
	_cfgd := make([]byte, _gfc.Height*_gfc.RowStride)
	_cdccd := _ce.NewWriterMSB(_cfgd)
	_abdc := make([]byte, _bedc)
	var (
		_caf  int
		_afbe uint64
	)
	for _caf = 0; _caf < _gfc.Height; _caf++ {
		if _, _dbcd = _fdac.Read(_abdc); _dbcd != nil {
			return _e.Wrap(_dbcd, _fbge, "\u0066u\u006c\u006c\u0020\u0062\u0079\u0074e")
		}
		if _, _dbcd = _cdccd.Write(_abdc); _dbcd != nil {
			return _e.Wrap(_dbcd, _fbge, "\u0066\u0075\u006c\u006c\u0020\u0062\u0079\u0074\u0065\u0073")
		}
		if _afbe, _dbcd = _fdac.ReadBits(byte(_effc)); _dbcd != nil {
			return _e.Wrap(_dbcd, _fbge, "\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0069\u0074\u0073")
		}
		if _dbcd = _cdccd.WriteByte(byte(_afbe) << uint(8-_effc)); _dbcd != nil {
			return _e.Wrap(_dbcd, _fbge, "\u006ca\u0073\u0074\u0020\u0062\u0079\u0074e")
		}
	}
	_gfc.Data = _cdccd.Data()
	return nil
}
func _ada(_ggb *Bitmap, _aad *Bitmap, _ba int) (_fdc error) {
	const _bed = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0050\u006fw\u0065\u0072\u0032\u004c\u006f\u0077"
	switch _ba {
	case 2:
		_fdc = _ad(_ggb, _aad)
	case 4:
		_fdc = _gf(_ggb, _aad)
	case 8:
		_fdc = _aaa(_ggb, _aad)
	default:
		return _e.Error(_bed, "\u0065\u0078p\u0061\u006e\u0073\u0069o\u006e\u0020f\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020r\u0061\u006e\u0067\u0065")
	}
	if _fdc != nil {
		_fdc = _e.Wrap(_fdc, _bed, "")
	}
	return _fdc
}
func (_baad *ClassedPoints) GetIntXByClass(i int) (int, error) {
	const _bbef = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _baad.IntSlice.Size() {
		return 0, _e.Errorf(_bbef, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_baad.XAtIndex(i)), nil
}
func Copy(d, s *Bitmap) (*Bitmap, error) { return _cdcf(d, s) }
func _ceec(_fca, _efc *Bitmap, _eaba, _fabb, _dff, _aagge, _ffdd int, _gdfg CombinationOperator) error {
	var _bdc int
	_agc := func() {
		_bdc++
		_dff += _efc.RowStride
		_aagge += _fca.RowStride
		_ffdd += _fca.RowStride
	}
	for _bdc = _eaba; _bdc < _fabb; _agc() {
		_aedb := _dff
		for _gfcc := _aagge; _gfcc <= _ffdd; _gfcc++ {
			_eadc, _bgbg := _efc.GetByte(_aedb)
			if _bgbg != nil {
				return _bgbg
			}
			_cfa, _bgbg := _fca.GetByte(_gfcc)
			if _bgbg != nil {
				return _bgbg
			}
			if _bgbg = _efc.SetByte(_aedb, _eeabd(_eadc, _cfa, _gdfg)); _bgbg != nil {
				return _bgbg
			}
			_aedb++
		}
	}
	return nil
}
func TstGetScaledSymbol(t *_f.T, sm *Bitmap, scale ...int) *Bitmap {
	if len(scale) == 0 {
		return sm
	}
	if scale[0] == 1 {
		return sm
	}
	_befga, _eagg := MorphSequence(sm, MorphProcess{Operation: MopReplicativeBinaryExpansion, Arguments: scale})
	_d.NoError(t, _eagg)
	return _befga
}
func (_gdde *Bitmaps) SortByWidth() { _affg := (*byWidth)(_gdde); _ff.Sort(_affg) }

type Getter interface{ GetBitmap() *Bitmap }

func (_eadcf *Boxes) Get(i int) (*_aa.Rectangle, error) {
	const _cddb = "\u0042o\u0078\u0065\u0073\u002e\u0047\u0065t"
	if _eadcf == nil {
		return nil, _e.Error(_cddb, "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if i > len(*_eadcf)-1 {
		return nil, _e.Errorf(_cddb, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return (*_eadcf)[i], nil
}
func TstVSymbol(t *_f.T, scale ...int) *Bitmap {
	_cbae, _dgcb := NewWithData(5, 5, []byte{0x88, 0x88, 0x88, 0x50, 0x20})
	_d.NoError(t, _dgcb)
	return TstGetScaledSymbol(t, _cbae, scale...)
}

type LocationFilter int

func (_agcd *Bitmaps) GetBox(i int) (*_aa.Rectangle, error) {
	const _ded = "\u0047\u0065\u0074\u0042\u006f\u0078"
	if _agcd == nil {
		return nil, _e.Error(_ded, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074\u006d\u0061\u0070s\u0027")
	}
	if i > len(_agcd.Boxes)-1 {
		return nil, _e.Errorf(_ded, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _agcd.Boxes[i], nil
}
func _begeb(_dbafg *Bitmap, _ebdea *Bitmap, _dege *Selection, _gfcad **Bitmap) (*Bitmap, error) {
	const _aedgc = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0031"
	if _ebdea == nil {
		return nil, _e.Error(_aedgc, "\u004d\u006f\u0072\u0070\u0068\u0041\u0072\u0067\u0073\u0031\u0020'\u0073\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066i\u006e\u0065\u0064")
	}
	if _dege == nil {
		return nil, _e.Error(_aedgc, "\u004d\u006f\u0072\u0068p\u0041\u0072\u0067\u0073\u0031\u0020\u0027\u0073\u0065\u006c'\u0020n\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064")
	}
	_fbef, _gdbd := _dege.Height, _dege.Width
	if _fbef == 0 || _gdbd == 0 {
		return nil, _e.Error(_aedgc, "\u0073\u0065\u006c\u0065ct\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _dbafg == nil {
		_dbafg = _ebdea.createTemplate()
		*_gfcad = _ebdea
		return _dbafg, nil
	}
	_dbafg.Width = _ebdea.Width
	_dbafg.Height = _ebdea.Height
	_dbafg.RowStride = _ebdea.RowStride
	_dbafg.Color = _ebdea.Color
	_dbafg.Data = make([]byte, _ebdea.RowStride*_ebdea.Height)
	if _dbafg == _ebdea {
		*_gfcad = _ebdea.Copy()
	} else {
		*_gfcad = _ebdea
	}
	return _dbafg, nil
}
func (_eecf *Bitmaps) Size() int             { return len(_eecf.Values) }
func (_aag *Bitmap) CreateTemplate() *Bitmap { return _aag.createTemplate() }

var _bdgd = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x78, 0x27, 0xC2, 0x27, 0x91, 0x00, 0x22, 0x48, 0x21, 0x03, 0x24, 0x91, 0x00, 0x22, 0x48, 0x21, 0x02, 0xA4, 0x95, 0x00, 0x22, 0x48, 0x21, 0x02, 0x64, 0x9B, 0x00, 0x3C, 0x78, 0x21, 0x02, 0x27, 0x91, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func New(width, height int) *Bitmap {
	_bfdd := _eced(width, height)
	_bfdd.Data = make([]byte, height*_bfdd.RowStride)
	return _bfdd
}
func (_ced *Bitmap) GetByteIndex(x, y int) int     { return y*_ced.RowStride + (x >> 3) }
func (_cffe *BitmapsArray) AddBitmaps(bm *Bitmaps) { _cffe.Values = append(_cffe.Values, bm) }
func _cfee(_debe *_ac.Stack) (_bdcd *fillSegment, _dgcd error) {
	const _cabcg = "\u0070\u006f\u0070\u0046\u0069\u006c\u006c\u0053\u0065g\u006d\u0065\u006e\u0074"
	if _debe == nil {
		return nil, _e.Error(_cabcg, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _debe.Aux == nil {
		return nil, _e.Error(_cabcg, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	_fabe, _fced := _debe.Pop()
	if !_fced {
		return nil, nil
	}
	_aeac, _fced := _fabe.(*fillSegment)
	if !_fced {
		return nil, _e.Error(_cabcg, "\u0073\u0074\u0061ck\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020c\u006fn\u0074a\u0069n\u0020\u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_bdcd = &fillSegment{_aeac._abbc, _aeac._bcaf, _aeac._ffeg + _aeac._bedfb, _aeac._bedfb}
	_debe.Aux.Push(_aeac)
	return _bdcd, nil
}
func MakePixelCentroidTab8() []int { return _egdf() }
func TstASymbol(t *_f.T) *Bitmap {
	t.Helper()
	_ecc := New(6, 6)
	_d.NoError(t, _ecc.SetPixel(1, 0, 1))
	_d.NoError(t, _ecc.SetPixel(2, 0, 1))
	_d.NoError(t, _ecc.SetPixel(3, 0, 1))
	_d.NoError(t, _ecc.SetPixel(4, 0, 1))
	_d.NoError(t, _ecc.SetPixel(5, 1, 1))
	_d.NoError(t, _ecc.SetPixel(1, 2, 1))
	_d.NoError(t, _ecc.SetPixel(2, 2, 1))
	_d.NoError(t, _ecc.SetPixel(3, 2, 1))
	_d.NoError(t, _ecc.SetPixel(4, 2, 1))
	_d.NoError(t, _ecc.SetPixel(5, 2, 1))
	_d.NoError(t, _ecc.SetPixel(0, 3, 1))
	_d.NoError(t, _ecc.SetPixel(5, 3, 1))
	_d.NoError(t, _ecc.SetPixel(0, 4, 1))
	_d.NoError(t, _ecc.SetPixel(5, 4, 1))
	_d.NoError(t, _ecc.SetPixel(1, 5, 1))
	_d.NoError(t, _ecc.SetPixel(2, 5, 1))
	_d.NoError(t, _ecc.SetPixel(3, 5, 1))
	_d.NoError(t, _ecc.SetPixel(4, 5, 1))
	_d.NoError(t, _ecc.SetPixel(5, 5, 1))
	return _ecc
}
func (_cefb *Bitmap) setEightPartlyBytes(_dcdg, _bcgd int, _fed uint64) (_bffe error) {
	var (
		_gaga  byte
		_dbgad int
	)
	const _agbeb = "\u0073\u0065\u0074\u0045ig\u0068\u0074\u0050\u0061\u0072\u0074\u006c\u0079\u0042\u0079\u0074\u0065\u0073"
	for _cabc := 1; _cabc <= _bcgd; _cabc++ {
		_dbgad = 64 - _cabc*8
		_gaga = byte(_fed >> uint(_dbgad) & 0xff)
		_ca.Log.Trace("\u0074\u0065\u006d\u0070\u003a\u0020\u0025\u0030\u0038\u0062\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a %\u0064,\u0020\u0069\u0064\u0078\u003a\u0020\u0025\u0064\u002c\u0020\u0066\u0075l\u006c\u0042\u0079\u0074\u0065\u0073\u004e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c \u0073\u0068\u0069\u0066\u0074\u003a\u0020\u0025\u0064", _gaga, _dcdg, _dcdg+_cabc-1, _bcgd, _dbgad)
		if _bffe = _cefb.SetByte(_dcdg+_cabc-1, _gaga); _bffe != nil {
			return _e.Wrap(_bffe, _agbeb, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
		}
	}
	_eaa := _cefb.RowStride*8 - _cefb.Width
	if _eaa == 0 {
		return nil
	}
	_dbgad -= 8
	_gaga = byte(_fed>>uint(_dbgad)&0xff) << uint(_eaa)
	if _bffe = _cefb.SetByte(_dcdg+_bcgd, _gaga); _bffe != nil {
		return _e.Wrap(_bffe, _agbeb, "\u0070\u0061\u0064\u0064\u0065\u0064")
	}
	return nil
}

const (
	Vanilla Color = iota
	Chocolate
)

func _fbac(_efb, _gef *Bitmap, _gae int, _bff []byte, _bbd int) (_cbba error) {
	const _afg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0033"
	var (
		_dg, _ffc, _eg, _dgc, _fgc, _cfc, _fgd, _add int
		_fgf, _cdg, _edf, _efad                      uint32
		_acd, _gge                                   byte
		_cgd                                         uint16
	)
	_ffcd := make([]byte, 4)
	_ebbb := make([]byte, 4)
	for _eg = 0; _eg < _efb.Height-1; _eg, _dgc = _eg+2, _dgc+1 {
		_dg = _eg * _efb.RowStride
		_ffc = _dgc * _gef.RowStride
		for _fgc, _cfc = 0, 0; _fgc < _bbd; _fgc, _cfc = _fgc+4, _cfc+1 {
			for _fgd = 0; _fgd < 4; _fgd++ {
				_add = _dg + _fgc + _fgd
				if _add <= len(_efb.Data)-1 && _add < _dg+_efb.RowStride {
					_ffcd[_fgd] = _efb.Data[_add]
				} else {
					_ffcd[_fgd] = 0x00
				}
				_add = _dg + _efb.RowStride + _fgc + _fgd
				if _add <= len(_efb.Data)-1 && _add < _dg+(2*_efb.RowStride) {
					_ebbb[_fgd] = _efb.Data[_add]
				} else {
					_ebbb[_fgd] = 0x00
				}
			}
			_fgf = _db.BigEndian.Uint32(_ffcd)
			_cdg = _db.BigEndian.Uint32(_ebbb)
			_edf = _fgf & _cdg
			_edf |= _edf << 1
			_efad = _fgf | _cdg
			_efad &= _efad << 1
			_cdg = _edf & _efad
			_cdg &= 0xaaaaaaaa
			_fgf = _cdg | (_cdg << 7)
			_acd = byte(_fgf >> 24)
			_gge = byte((_fgf >> 8) & 0xff)
			_add = _ffc + _cfc
			if _add+1 == len(_gef.Data)-1 || _add+1 >= _ffc+_gef.RowStride {
				if _cbba = _gef.SetByte(_add, _bff[_acd]); _cbba != nil {
					return _e.Wrapf(_cbba, _afg, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _add)
				}
			} else {
				_cgd = (uint16(_bff[_acd]) << 8) | uint16(_bff[_gge])
				if _cbba = _gef.setTwoBytes(_add, _cgd); _cbba != nil {
					return _e.Wrapf(_cbba, _afg, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _add)
				}
				_cfc++
			}
		}
	}
	return nil
}

type SelectionValue int

func TstCSymbol(t *_f.T) *Bitmap {
	t.Helper()
	_dcfdd := New(6, 6)
	_d.NoError(t, _dcfdd.SetPixel(1, 0, 1))
	_d.NoError(t, _dcfdd.SetPixel(2, 0, 1))
	_d.NoError(t, _dcfdd.SetPixel(3, 0, 1))
	_d.NoError(t, _dcfdd.SetPixel(4, 0, 1))
	_d.NoError(t, _dcfdd.SetPixel(0, 1, 1))
	_d.NoError(t, _dcfdd.SetPixel(5, 1, 1))
	_d.NoError(t, _dcfdd.SetPixel(0, 2, 1))
	_d.NoError(t, _dcfdd.SetPixel(0, 3, 1))
	_d.NoError(t, _dcfdd.SetPixel(0, 4, 1))
	_d.NoError(t, _dcfdd.SetPixel(5, 4, 1))
	_d.NoError(t, _dcfdd.SetPixel(1, 5, 1))
	_d.NoError(t, _dcfdd.SetPixel(2, 5, 1))
	_d.NoError(t, _dcfdd.SetPixel(3, 5, 1))
	_d.NoError(t, _dcfdd.SetPixel(4, 5, 1))
	return _dcfdd
}
func _gf(_cbc, _ec *Bitmap) (_gc error) {
	const _fd = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0034"
	_ee := _ec.RowStride
	_de := _cbc.RowStride
	_gg := _ec.RowStride*4 - _cbc.RowStride
	var (
		_ga, _ggf                              byte
		_fa                                    uint32
		_fba, _ae, _fe, _deb, _dbg, _fec, _ggd int
	)
	for _fe = 0; _fe < _ec.Height; _fe++ {
		_fba = _fe * _ee
		_ae = 4 * _fe * _de
		for _deb = 0; _deb < _ee; _deb++ {
			_ga = _ec.Data[_fba+_deb]
			_fa = _bffba[_ga]
			_fec = _ae + _deb*4
			if _gg != 0 && (_deb+1)*4 > _cbc.RowStride {
				for _dbg = _gg; _dbg > 0; _dbg-- {
					_ggf = byte((_fa >> uint(_dbg*8)) & 0xff)
					_ggd = _fec + (_gg - _dbg)
					if _gc = _cbc.SetByte(_ggd, _ggf); _gc != nil {
						return _e.Wrapf(_gc, _fd, "D\u0069\u0066\u0066\u0065\u0072\u0065n\u0074\u0020\u0072\u006f\u0077\u0073\u0074\u0072\u0069d\u0065\u0073\u002e \u004b:\u0020\u0025\u0064", _dbg)
					}
				}
			} else if _gc = _cbc.setFourBytes(_fec, _fa); _gc != nil {
				return _e.Wrap(_gc, _fd, "")
			}
			if _gc = _cbc.setFourBytes(_ae+_deb*4, _bffba[_ec.Data[_fba+_deb]]); _gc != nil {
				return _e.Wrap(_gc, _fd, "")
			}
		}
		for _dbg = 1; _dbg < 4; _dbg++ {
			for _deb = 0; _deb < _de; _deb++ {
				if _gc = _cbc.SetByte(_ae+_dbg*_de+_deb, _cbc.Data[_ae+_deb]); _gc != nil {
					return _e.Wrapf(_gc, _fd, "\u0063\u006f\u0070\u0079\u0020\u0027\u0071\u0075\u0061\u0064\u0072\u0061\u0062l\u0065\u0027\u0020\u006c\u0069\u006ee\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0079\u0074\u0065\u003a \u0027\u0025\u0064\u0027", _dbg, _deb)
				}
			}
		}
	}
	return nil
}
func (_bfffe *ClassedPoints) ySortFunction() func(_affa int, _cedc int) bool {
	return func(_afaf, _gfbb int) bool { return _bfffe.YAtIndex(_afaf) < _bfffe.YAtIndex(_gfbb) }
}
func Centroids(bms []*Bitmap) (*Points, error) {
	_gded := make([]Point, len(bms))
	_gccg := _egdf()
	_caff := _dfbeb()
	var _dggg error
	for _gbcgd, _daeb := range bms {
		_gded[_gbcgd], _dggg = _daeb.centroid(_gccg, _caff)
		if _dggg != nil {
			return nil, _dggg
		}
	}
	_dcbf := Points(_gded)
	return &_dcbf, nil
}
func (_afcff *Bitmap) setTwoBytes(_cebc int, _ddb uint16) error {
	if _cebc+1 > len(_afcff.Data)-1 {
		return _e.Errorf("s\u0065\u0074\u0054\u0077\u006f\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _cebc)
	}
	_afcff.Data[_cebc] = byte((_ddb & 0xff00) >> 8)
	_afcff.Data[_cebc+1] = byte(_ddb & 0xff)
	return nil
}

type MorphProcess struct {
	Operation MorphOperation
	Arguments []int
}

const (
	_ LocationFilter = iota
	LocSelectWidth
	LocSelectHeight
	LocSelectXVal
	LocSelectYVal
	LocSelectIfEither
	LocSelectIfBoth
)

func (_dcceg *BitmapsArray) GetBox(i int) (*_aa.Rectangle, error) {
	const _gbcga = "\u0042\u0069\u0074\u006dap\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u006f\u0078"
	if _dcceg == nil {
		return nil, _e.Error(_gbcga, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_dcceg.Boxes)-1 {
		return nil, _e.Errorf(_gbcga, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _dcceg.Boxes[i], nil
}
func _cbfg(_fagg *Bitmap, _eddf int) (*Bitmap, error) {
	const _eaabg = "\u0065x\u0070a\u006e\u0064\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0065"
	if _fagg == nil {
		return nil, _e.Error(_eaabg, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _eddf <= 0 {
		return nil, _e.Error(_eaabg, "i\u006e\u0076\u0061\u006cid\u0020f\u0061\u0063\u0074\u006f\u0072 \u002d\u0020\u003c\u003d\u0020\u0030")
	}
	if _eddf == 1 {
		_afee, _aefa := _cdcf(nil, _fagg)
		if _aefa != nil {
			return nil, _e.Wrap(_aefa, _eaabg, "\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u003d\u0020\u0031")
		}
		return _afee, nil
	}
	_fgeb, _eeega := _adgf(_fagg, _eddf, _eddf)
	if _eeega != nil {
		return nil, _e.Wrap(_eeega, _eaabg, "")
	}
	return _fgeb, nil
}
func _aada(_cca, _ffe *Bitmap, _abg int, _cbb []byte, _agad int) (_aadad error) {
	const _befc = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0031"
	var (
		_gda, _gbb, _dbe, _aeb, _fg, _ggbc, _bfd, _fdd int
		_bea, _fbe                                     uint32
		_gbgb, _fddb                                   byte
		_deg                                           uint16
	)
	_edcd := make([]byte, 4)
	_dde := make([]byte, 4)
	for _dbe = 0; _dbe < _cca.Height-1; _dbe, _aeb = _dbe+2, _aeb+1 {
		_gda = _dbe * _cca.RowStride
		_gbb = _aeb * _ffe.RowStride
		for _fg, _ggbc = 0, 0; _fg < _agad; _fg, _ggbc = _fg+4, _ggbc+1 {
			for _bfd = 0; _bfd < 4; _bfd++ {
				_fdd = _gda + _fg + _bfd
				if _fdd <= len(_cca.Data)-1 && _fdd < _gda+_cca.RowStride {
					_edcd[_bfd] = _cca.Data[_fdd]
				} else {
					_edcd[_bfd] = 0x00
				}
				_fdd = _gda + _cca.RowStride + _fg + _bfd
				if _fdd <= len(_cca.Data)-1 && _fdd < _gda+(2*_cca.RowStride) {
					_dde[_bfd] = _cca.Data[_fdd]
				} else {
					_dde[_bfd] = 0x00
				}
			}
			_bea = _db.BigEndian.Uint32(_edcd)
			_fbe = _db.BigEndian.Uint32(_dde)
			_fbe |= _bea
			_fbe |= _fbe << 1
			_fbe &= 0xaaaaaaaa
			_bea = _fbe | (_fbe << 7)
			_gbgb = byte(_bea >> 24)
			_fddb = byte((_bea >> 8) & 0xff)
			_fdd = _gbb + _ggbc
			if _fdd+1 == len(_ffe.Data)-1 || _fdd+1 >= _gbb+_ffe.RowStride {
				_ffe.Data[_fdd] = _cbb[_gbgb]
			} else {
				_deg = (uint16(_cbb[_gbgb]) << 8) | uint16(_cbb[_fddb])
				if _aadad = _ffe.setTwoBytes(_fdd, _deg); _aadad != nil {
					return _e.Wrapf(_aadad, _befc, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fdd)
				}
				_ggbc++
			}
		}
	}
	return nil
}
func _ecbgb(_agea *Bitmap, _dfc ...MorphProcess) (_acad *Bitmap, _ddgb error) {
	const _ebf = "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065"
	if _agea == nil {
		return nil, _e.Error(_ebf, "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065 \u0073\u006f\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if len(_dfc) == 0 {
		return nil, _e.Error(_ebf, "m\u006f\u0072\u0070\u0068\u0053\u0065q\u0075\u0065\u006e\u0063\u0065\u002c \u0073\u0065\u0071\u0075\u0065\u006e\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _ddgb = _eeege(_dfc...); _ddgb != nil {
		return nil, _e.Wrap(_ddgb, _ebf, "")
	}
	var _gfcgd, _abe, _ffbf int
	_acad = _agea.Copy()
	for _, _agdg := range _dfc {
		switch _agdg.Operation {
		case MopDilation:
			_gfcgd, _abe = _agdg.getWidthHeight()
			_acad, _ddgb = DilateBrick(nil, _acad, _gfcgd, _abe)
			if _ddgb != nil {
				return nil, _e.Wrap(_ddgb, _ebf, "")
			}
		case MopErosion:
			_gfcgd, _abe = _agdg.getWidthHeight()
			_acad, _ddgb = _geag(nil, _acad, _gfcgd, _abe)
			if _ddgb != nil {
				return nil, _e.Wrap(_ddgb, _ebf, "")
			}
		case MopOpening:
			_gfcgd, _abe = _agdg.getWidthHeight()
			_acad, _ddgb = _gfgf(nil, _acad, _gfcgd, _abe)
			if _ddgb != nil {
				return nil, _e.Wrap(_ddgb, _ebf, "")
			}
		case MopClosing:
			_gfcgd, _abe = _agdg.getWidthHeight()
			_acad, _ddgb = _facd(nil, _acad, _gfcgd, _abe)
			if _ddgb != nil {
				return nil, _e.Wrap(_ddgb, _ebf, "")
			}
		case MopRankBinaryReduction:
			_acad, _ddgb = _fbb(_acad, _agdg.Arguments...)
			if _ddgb != nil {
				return nil, _e.Wrap(_ddgb, _ebf, "")
			}
		case MopReplicativeBinaryExpansion:
			_acad, _ddgb = _cbfg(_acad, _agdg.Arguments[0])
			if _ddgb != nil {
				return nil, _e.Wrap(_ddgb, _ebf, "")
			}
		case MopAddBorder:
			_ffbf = _agdg.Arguments[0]
			_acad, _ddgb = _acad.AddBorder(_ffbf, 0)
			if _ddgb != nil {
				return nil, _e.Wrap(_ddgb, _ebf, "")
			}
		default:
			return nil, _e.Error(_ebf, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u006fr\u0070\u0068\u004f\u0070\u0065\u0072\u0061ti\u006f\u006e\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u006f t\u0068\u0065 \u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065")
		}
	}
	if _ffbf > 0 {
		_acad, _ddgb = _acad.RemoveBorder(_ffbf)
		if _ddgb != nil {
			return nil, _e.Wrap(_ddgb, _ebf, "\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u003e\u0020\u0030")
		}
	}
	return _acad, nil
}

var (
	_dfag  = _aga()
	_bffba = _cdf()
	_cede  = _fbc()
)

func _aga() (_dcb [256]uint16) {
	for _fdfd := 0; _fdfd < 256; _fdfd++ {
		if _fdfd&0x01 != 0 {
			_dcb[_fdfd] |= 0x3
		}
		if _fdfd&0x02 != 0 {
			_dcb[_fdfd] |= 0xc
		}
		if _fdfd&0x04 != 0 {
			_dcb[_fdfd] |= 0x30
		}
		if _fdfd&0x08 != 0 {
			_dcb[_fdfd] |= 0xc0
		}
		if _fdfd&0x10 != 0 {
			_dcb[_fdfd] |= 0x300
		}
		if _fdfd&0x20 != 0 {
			_dcb[_fdfd] |= 0xc00
		}
		if _fdfd&0x40 != 0 {
			_dcb[_fdfd] |= 0x3000
		}
		if _fdfd&0x80 != 0 {
			_dcb[_fdfd] |= 0xc000
		}
	}
	return _dcb
}
func (_bcgf *Bitmaps) GetBitmap(i int) (*Bitmap, error) {
	const _gafgc = "\u0047e\u0074\u0042\u0069\u0074\u006d\u0061p"
	if _bcgf == nil {
		return nil, _e.Error(_gafgc, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if i > len(_bcgf.Values)-1 {
		return nil, _e.Errorf(_gafgc, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _bcgf.Values[i], nil
}
func _geag(_ddac, _afgd *Bitmap, _ccge, _fdcc int) (*Bitmap, error) {
	const _gdgd = "\u0065\u0072\u006f\u0064\u0065\u0042\u0072\u0069\u0063\u006b"
	if _afgd == nil {
		return nil, _e.Error(_gdgd, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ccge < 1 || _fdcc < 1 {
		return nil, _e.Error(_gdgd, "\u0068\u0073\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069\u007a\u0065\u0020\u0061\u0072e\u0020\u006e\u006f\u0074\u0020\u0067\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006fr\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _ccge == 1 && _fdcc == 1 {
		_accd, _aegaa := _cdcf(_ddac, _afgd)
		if _aegaa != nil {
			return nil, _e.Wrap(_aegaa, _gdgd, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _accd, nil
	}
	if _ccge == 1 || _fdcc == 1 {
		_eaab := SelCreateBrick(_fdcc, _ccge, _fdcc/2, _ccge/2, SelHit)
		_gggb, _caee := _cfgc(_ddac, _afgd, _eaab)
		if _caee != nil {
			return nil, _e.Wrap(_caee, _gdgd, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _gggb, nil
	}
	_bedgf := SelCreateBrick(1, _ccge, 0, _ccge/2, SelHit)
	_daf := SelCreateBrick(_fdcc, 1, _fdcc/2, 0, SelHit)
	_adbcf, _bbcda := _cfgc(nil, _afgd, _bedgf)
	if _bbcda != nil {
		return nil, _e.Wrap(_bbcda, _gdgd, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_ddac, _bbcda = _cfgc(_ddac, _adbcf, _daf)
	if _bbcda != nil {
		return nil, _e.Wrap(_bbcda, _gdgd, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _ddac, nil
}

const (
	_ SizeSelection = iota
	SizeSelectByWidth
	SizeSelectByHeight
	SizeSelectByMaxDimension
	SizeSelectByArea
	SizeSelectByPerimeter
)

type Selection struct {
	Height, Width int
	Cx, Cy        int
	Name          string
	Data          [][]SelectionValue
}

func (_fded *ClassedPoints) Less(i, j int) bool { return _fded._acab(i, j) }
func _cdfd(_afbc, _eeg *Bitmap, _cbca, _dacb int) (*Bitmap, error) {
	const _dbdg = "\u0063\u006c\u006f\u0073\u0065\u0042\u0072\u0069\u0063\u006b"
	if _eeg == nil {
		return nil, _e.Error(_dbdg, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _cbca < 1 || _dacb < 1 {
		return nil, _e.Error(_dbdg, "\u0068S\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0053\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _cbca == 1 && _dacb == 1 {
		return _eeg.Copy(), nil
	}
	if _cbca == 1 || _dacb == 1 {
		_beed := SelCreateBrick(_dacb, _cbca, _dacb/2, _cbca/2, SelHit)
		var _bdd error
		_afbc, _bdd = _gedd(_afbc, _eeg, _beed)
		if _bdd != nil {
			return nil, _e.Wrap(_bdd, _dbdg, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _afbc, nil
	}
	_dce := SelCreateBrick(1, _cbca, 0, _cbca/2, SelHit)
	_efca := SelCreateBrick(_dacb, 1, _dacb/2, 0, SelHit)
	_cgee, _ddab := _caagg(nil, _eeg, _dce)
	if _ddab != nil {
		return nil, _e.Wrap(_ddab, _dbdg, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _afbc, _ddab = _caagg(_afbc, _cgee, _efca); _ddab != nil {
		return nil, _e.Wrap(_ddab, _dbdg, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _, _ddab = _cfgc(_cgee, _afbc, _dce); _ddab != nil {
		return nil, _e.Wrap(_ddab, _dbdg, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	if _, _ddab = _cfgc(_afbc, _cgee, _efca); _ddab != nil {
		return nil, _e.Wrap(_ddab, _dbdg, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _afbc, nil
}
func (_ecb *Bitmap) SizesEqual(s *Bitmap) bool {
	if _ecb == s {
		return true
	}
	if _ecb.Width != s.Width || _ecb.Height != s.Height {
		return false
	}
	return true
}
func init() {
	const _aaaad = "\u0062\u0069\u0074\u006dap\u0073\u002e\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0061\u0074\u0069o\u006e"
	_cbff = New(50, 40)
	var _gdbcd error
	_cbff, _gdbcd = _cbff.AddBorder(2, 1)
	if _gdbcd != nil {
		panic(_e.Wrap(_gdbcd, _aaaad, "f\u0072\u0061\u006d\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
	_gabd, _gdbcd = NewWithData(50, 22, _bdgd)
	if _gdbcd != nil {
		panic(_e.Wrap(_gdbcd, _aaaad, "i\u006d\u0061\u0067\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
}
func TstWordBitmapWithSpaces(t *_f.T, scale ...int) *Bitmap {
	_cedgf := 1
	if len(scale) > 0 {
		_cedgf = scale[0]
	}
	_afaa := 3
	_afbaf := 9 + 7 + 15 + 2*_afaa + 2*_afaa
	_bgfg := 5 + _afaa + 5 + 2*_afaa
	_gfegd := New(_afbaf*_cedgf, _bgfg*_cedgf)
	_cbafg := &Bitmaps{}
	var _ccbag *int
	_afaa *= _cedgf
	_cbag := _afaa
	_ccbag = &_cbag
	_ccab := _afaa
	_bade := TstDSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, _afaa)
	_bade = TstISymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstTSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, _afaa)
	_bade = TstNSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstWSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 0)
	*_ccbag = _afaa
	_ccab = 5*_cedgf + _afaa
	_bade = TstOSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstRSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, _afaa)
	_bade = TstNSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstESymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstVSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstESymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 1*_cedgf)
	_bade = TstRSymbol(t, scale...)
	TstAddSymbol(t, _cbafg, _bade, _ccbag, _ccab, 0)
	TstWriteSymbols(t, _cbafg, _gfegd)
	return _gfegd
}
func (_fagc *Bitmap) RemoveBorder(borderSize int) (*Bitmap, error) {
	if borderSize == 0 {
		return _fagc.Copy(), nil
	}
	_baf, _aagg := _fagc.removeBorderGeneral(borderSize, borderSize, borderSize, borderSize)
	if _aagg != nil {
		return nil, _e.Wrap(_aagg, "\u0052\u0065\u006do\u0076\u0065\u0042\u006f\u0072\u0064\u0065\u0072", "")
	}
	return _baf, nil
}
func _ggee(_bcbd, _dcgfg *Bitmap, _ecgac, _ace int) (_gafg error) {
	const _dfbfe = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0038"
	var (
		_aagb, _eca, _agffg, _gafed                             int
		_edbc, _bgfe, _cbde, _bafbc, _bdcb, _ffa, _egfa, _cgaad byte
	)
	for _aagb = 0; _aagb < _ecgac; _aagb++ {
		_agffg = _aagb * _bcbd.RowStride
		_gafed = _aagb * _dcgfg.RowStride
		for _eca = 0; _eca < _ace; _eca++ {
			if _edbc, _gafg = _bcbd.GetByte(_agffg + _eca); _gafg != nil {
				return _e.Wrap(_gafg, _dfbfe, "\u0067e\u0074 \u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _bgfe, _gafg = _dcgfg.GetByte(_gafed + _eca); _gafg != nil {
				return _e.Wrap(_gafg, _dfbfe, "\u0067\u0065\u0074\u0020\u006d\u0061\u0073\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _aagb > 0 {
				if _cbde, _gafg = _bcbd.GetByte(_agffg - _bcbd.RowStride + _eca); _gafg != nil {
					return _e.Wrap(_gafg, _dfbfe, "\u0069\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
				}
				_edbc |= _cbde | (_cbde << 1) | (_cbde >> 1)
				if _eca > 0 {
					if _cgaad, _gafg = _bcbd.GetByte(_agffg - _bcbd.RowStride + _eca - 1); _gafg != nil {
						return _e.Wrap(_gafg, _dfbfe, "\u0069\u0020\u003e\u00200 \u0026\u0026\u0020\u006a\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
					}
					_edbc |= _cgaad << 7
				}
				if _eca < _ace-1 {
					if _cgaad, _gafg = _bcbd.GetByte(_agffg - _bcbd.RowStride + _eca + 1); _gafg != nil {
						return _e.Wrap(_gafg, _dfbfe, "\u006a\u0020<\u0020\u0077\u0070l\u0020\u002d\u0020\u0031\u0020\u0062\u0079\u0074\u0065")
					}
					_edbc |= _cgaad >> 7
				}
			}
			if _eca > 0 {
				if _bafbc, _gafg = _bcbd.GetByte(_agffg + _eca - 1); _gafg != nil {
					return _e.Wrap(_gafg, _dfbfe, "\u006a\u0020\u003e \u0030")
				}
				_edbc |= _bafbc << 7
			}
			_edbc &= _bgfe
			if _edbc == 0 || ^_edbc == 0 {
				if _gafg = _bcbd.SetByte(_agffg+_eca, _edbc); _gafg != nil {
					return _e.Wrap(_gafg, _dfbfe, "\u0073e\u0074t\u0069\u006e\u0067\u0020\u0065m\u0070\u0074y\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_egfa = _edbc
				_edbc = (_edbc | (_edbc >> 1) | (_edbc << 1)) & _bgfe
				if (_edbc ^ _egfa) == 0 {
					if _gafg = _bcbd.SetByte(_agffg+_eca, _edbc); _gafg != nil {
						return _e.Wrap(_gafg, _dfbfe, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _aagb = _ecgac - 1; _aagb >= 0; _aagb-- {
		_agffg = _aagb * _bcbd.RowStride
		_gafed = _aagb * _dcgfg.RowStride
		for _eca = _ace - 1; _eca >= 0; _eca-- {
			if _edbc, _gafg = _bcbd.GetByte(_agffg + _eca); _gafg != nil {
				return _e.Wrap(_gafg, _dfbfe, "\u0072\u0065\u0076er\u0073\u0065\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _bgfe, _gafg = _dcgfg.GetByte(_gafed + _eca); _gafg != nil {
				return _e.Wrap(_gafg, _dfbfe, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _aagb < _ecgac-1 {
				if _bdcb, _gafg = _bcbd.GetByte(_agffg + _bcbd.RowStride + _eca); _gafg != nil {
					return _e.Wrap(_gafg, _dfbfe, "\u0069\u0020\u003c\u0020h\u0020\u002d\u0020\u0031\u0020\u002d\u003e\u0020\u0067\u0065t\u0020s\u006f\u0075\u0072\u0063\u0065\u0020\u0062y\u0074\u0065")
				}
				_edbc |= _bdcb | (_bdcb << 1) | _bdcb>>1
				if _eca > 0 {
					if _cgaad, _gafg = _bcbd.GetByte(_agffg + _bcbd.RowStride + _eca - 1); _gafg != nil {
						return _e.Wrap(_gafg, _dfbfe, "\u0069\u0020\u003c h\u002d\u0031\u0020\u0026\u0020\u006a\u0020\u003e\u00200\u0020-\u003e \u0067e\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
					}
					_edbc |= _cgaad << 7
				}
				if _eca < _ace-1 {
					if _cgaad, _gafg = _bcbd.GetByte(_agffg + _bcbd.RowStride + _eca + 1); _gafg != nil {
						return _e.Wrap(_gafg, _dfbfe, "\u0069\u0020\u003c\u0020\u0068\u002d\u0031\u0020\u0026\u0026\u0020\u006a\u0020\u003c\u0077\u0070\u006c\u002d\u0031\u0020\u002d\u003e\u0020\u0067e\u0074\u0020\u0073\u006f\u0075r\u0063\u0065 \u0062\u0079\u0074\u0065")
					}
					_edbc |= _cgaad >> 7
				}
			}
			if _eca < _ace-1 {
				if _ffa, _gafg = _bcbd.GetByte(_agffg + _eca + 1); _gafg != nil {
					return _e.Wrap(_gafg, _dfbfe, "\u006a\u0020<\u0020\u0077\u0070\u006c\u0020\u002d\u0031\u0020\u002d\u003e\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020by\u0074\u0065")
				}
				_edbc |= _ffa >> 7
			}
			_edbc &= _bgfe
			if _edbc == 0 || (^_edbc) == 0 {
				if _gafg = _bcbd.SetByte(_agffg+_eca, _edbc); _gafg != nil {
					return _e.Wrap(_gafg, _dfbfe, "\u0073e\u0074 \u006d\u0061\u0073\u006b\u0065\u0064\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_egfa = _edbc
				_edbc = (_edbc | (_edbc >> 1) | (_edbc << 1)) & _bgfe
				if (_edbc ^ _egfa) == 0 {
					if _gafg = _bcbd.SetByte(_agffg+_eca, _edbc); _gafg != nil {
						return _e.Wrap(_gafg, _dfbfe, "r\u0065\u0076\u0065\u0072se\u0020s\u0065\u0074\u0020\u0070\u0072e\u0076\u0020\u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	return nil
}
func _bgge(_dcba, _fabf *Bitmap, _bgce, _bbda int) (_fadg error) {
	const _bae = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0034"
	var (
		_cfgfc, _beac, _fdgc, _bcbbe                      int
		_dgbbc, _decb, _fcac, _aefe, _bceb, _ageda, _agag byte
	)
	for _cfgfc = 0; _cfgfc < _bgce; _cfgfc++ {
		_fdgc = _cfgfc * _dcba.RowStride
		_bcbbe = _cfgfc * _fabf.RowStride
		for _beac = 0; _beac < _bbda; _beac++ {
			_dgbbc, _fadg = _dcba.GetByte(_fdgc + _beac)
			if _fadg != nil {
				return _e.Wrap(_fadg, _bae, "\u0066i\u0072\u0073\u0074\u0020\u0067\u0065t")
			}
			_decb, _fadg = _fabf.GetByte(_bcbbe + _beac)
			if _fadg != nil {
				return _e.Wrap(_fadg, _bae, "\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0067\u0065\u0074")
			}
			if _cfgfc > 0 {
				_fcac, _fadg = _dcba.GetByte(_fdgc - _dcba.RowStride + _beac)
				if _fadg != nil {
					return _e.Wrap(_fadg, _bae, "\u0069\u0020\u003e \u0030")
				}
				_dgbbc |= _fcac
			}
			if _beac > 0 {
				_aefe, _fadg = _dcba.GetByte(_fdgc + _beac - 1)
				if _fadg != nil {
					return _e.Wrap(_fadg, _bae, "\u006a\u0020\u003e \u0030")
				}
				_dgbbc |= _aefe << 7
			}
			_dgbbc &= _decb
			if _dgbbc == 0 || (^_dgbbc) == 0 {
				if _fadg = _dcba.SetByte(_fdgc+_beac, _dgbbc); _fadg != nil {
					return _e.Wrap(_fadg, _bae, "b\u0074\u0020\u003d\u003d 0\u0020|\u007c\u0020\u0028\u005e\u0062t\u0029\u0020\u003d\u003d\u0020\u0030")
				}
				continue
			}
			for {
				_agag = _dgbbc
				_dgbbc = (_dgbbc | (_dgbbc >> 1) | (_dgbbc << 1)) & _decb
				if (_dgbbc ^ _agag) == 0 {
					if _fadg = _dcba.SetByte(_fdgc+_beac, _dgbbc); _fadg != nil {
						return _e.Wrap(_fadg, _bae, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _cfgfc = _bgce - 1; _cfgfc >= 0; _cfgfc-- {
		_fdgc = _cfgfc * _dcba.RowStride
		_bcbbe = _cfgfc * _fabf.RowStride
		for _beac = _bbda - 1; _beac >= 0; _beac-- {
			if _dgbbc, _fadg = _dcba.GetByte(_fdgc + _beac); _fadg != nil {
				return _e.Wrap(_fadg, _bae, "\u0072\u0065\u0076\u0065\u0072\u0073\u0065\u0020\u0066\u0069\u0072\u0073t\u0020\u0067\u0065\u0074")
			}
			if _decb, _fadg = _fabf.GetByte(_bcbbe + _beac); _fadg != nil {
				return _e.Wrap(_fadg, _bae, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _cfgfc < _bgce-1 {
				if _bceb, _fadg = _dcba.GetByte(_fdgc + _dcba.RowStride + _beac); _fadg != nil {
					return _e.Wrap(_fadg, _bae, "\u0072\u0065v\u0065\u0072\u0073e\u0020\u0069\u0020\u003c\u0020\u0068\u0020\u002d\u0031")
				}
				_dgbbc |= _bceb
			}
			if _beac < _bbda-1 {
				if _ageda, _fadg = _dcba.GetByte(_fdgc + _beac + 1); _fadg != nil {
					return _e.Wrap(_fadg, _bae, "\u0072\u0065\u0076\u0065rs\u0065\u0020\u006a\u0020\u003c\u0020\u0077\u0070\u006c\u0020\u002d\u0020\u0031")
				}
				_dgbbc |= _ageda >> 7
			}
			_dgbbc &= _decb
			if _dgbbc == 0 || (^_dgbbc) == 0 {
				if _fadg = _dcba.SetByte(_fdgc+_beac, _dgbbc); _fadg != nil {
					return _e.Wrap(_fadg, _bae, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006d\u0061\u0073k\u0065\u0064\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061i\u006c\u0065\u0064")
				}
				continue
			}
			for {
				_agag = _dgbbc
				_dgbbc = (_dgbbc | (_dgbbc >> 1) | (_dgbbc << 1)) & _decb
				if (_dgbbc ^ _agag) == 0 {
					if _fadg = _dcba.SetByte(_fdgc+_beac, _dgbbc); _fadg != nil {
						return _e.Wrap(_fadg, _bae, "\u0072e\u0076\u0065\u0072\u0073e\u0020\u0073\u0065\u0074\u0074i\u006eg\u0020p\u0072\u0065\u0076\u0020\u0062\u0079\u0074e")
					}
					break
				}
			}
		}
	}
	return nil
}
func TstWSymbol(t *_f.T, scale ...int) *Bitmap {
	_dggd, _bffd := NewWithData(5, 5, []byte{0x88, 0x88, 0xA8, 0xD8, 0x88})
	_d.NoError(t, _bffd)
	return TstGetScaledSymbol(t, _dggd, scale...)
}
func (_fcg *Bitmap) createTemplate() *Bitmap {
	return &Bitmap{Width: _fcg.Width, Height: _fcg.Height, RowStride: _fcg.RowStride, Color: _fcg.Color, Text: _fcg.Text, BitmapNumber: _fcg.BitmapNumber, Special: _fcg.Special, Data: make([]byte, len(_fcg.Data))}
}
func TstDSymbol(t *_f.T, scale ...int) *Bitmap {
	_cggc, _agacf := NewWithData(4, 5, []byte{0xf0, 0x90, 0x90, 0x90, 0xE0})
	_d.NoError(t, _agacf)
	return TstGetScaledSymbol(t, _cggc, scale...)
}
func (_degf *Bitmap) GetChocolateData() []byte {
	if _degf.Color == Vanilla {
		_degf.inverseData()
	}
	return _degf.Data
}
func TstESymbol(t *_f.T, scale ...int) *Bitmap {
	_cfcg, _caeg := NewWithData(4, 5, []byte{0xF0, 0x80, 0xE0, 0x80, 0xF0})
	_d.NoError(t, _caeg)
	return TstGetScaledSymbol(t, _cfcg, scale...)
}
func TstWriteSymbols(t *_f.T, bms *Bitmaps, src *Bitmap) {
	for _adff := 0; _adff < bms.Size(); _adff++ {
		_bedcf := bms.Values[_adff]
		_eebb := bms.Boxes[_adff]
		_fagce := src.RasterOperation(_eebb.Min.X, _eebb.Min.Y, _bedcf.Width, _bedcf.Height, PixSrc, _bedcf, 0, 0)
		_d.NoError(t, _fagce)
	}
}
func _facd(_gedfe, _cefaa *Bitmap, _cebaf, _fbfb int) (*Bitmap, error) {
	const _fdfe = "\u0063\u006c\u006f\u0073\u0065\u0053\u0061\u0066\u0065B\u0072\u0069\u0063\u006b"
	if _cefaa == nil {
		return nil, _e.Error(_fdfe, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _cebaf < 1 || _fbfb < 1 {
		return nil, _e.Error(_fdfe, "\u0068s\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _cebaf == 1 && _fbfb == 1 {
		return _cdcf(_gedfe, _cefaa)
	}
	if MorphBC == SymmetricMorphBC {
		_bfdg, _cabb := _cdfd(_gedfe, _cefaa, _cebaf, _fbfb)
		if _cabb != nil {
			return nil, _e.Wrap(_cabb, _fdfe, "\u0053\u0079m\u006d\u0065\u0074r\u0069\u0063\u004d\u006f\u0072\u0070\u0068\u0042\u0043")
		}
		return _bfdg, nil
	}
	_dfab := _aefd(_cebaf/2, _fbfb/2)
	_dacda := 8 * ((_dfab + 7) / 8)
	_bacf, _eacg := _cefaa.AddBorder(_dacda, 0)
	if _eacg != nil {
		return nil, _e.Wrapf(_eacg, _fdfe, "\u0042\u006f\u0072\u0064\u0065\u0072\u0053\u0069\u007ae\u003a\u0020\u0025\u0064", _dacda)
	}
	var _fccb, _abba *Bitmap
	if _cebaf == 1 || _fbfb == 1 {
		_gbbc := SelCreateBrick(_fbfb, _cebaf, _fbfb/2, _cebaf/2, SelHit)
		_fccb, _eacg = _gedd(nil, _bacf, _gbbc)
		if _eacg != nil {
			return nil, _e.Wrap(_eacg, _fdfe, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
	} else {
		_dagf := SelCreateBrick(1, _cebaf, 0, _cebaf/2, SelHit)
		_aebb, _acg := _caagg(nil, _bacf, _dagf)
		if _acg != nil {
			return nil, _e.Wrap(_acg, _fdfe, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0064\u0069\u006c\u0061t\u0065")
		}
		_bbg := SelCreateBrick(_fbfb, 1, _fbfb/2, 0, SelHit)
		_fccb, _acg = _caagg(nil, _aebb, _bbg)
		if _acg != nil {
			return nil, _e.Wrap(_acg, _fdfe, "\u0072\u0065\u0067ul\u0061\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
		}
		if _, _acg = _cfgc(_aebb, _fccb, _dagf); _acg != nil {
			return nil, _e.Wrap(_acg, _fdfe, "r\u0065\u0067\u0075\u006car\u0020-\u0020\u0066\u0069\u0072\u0073t\u0020\u0065\u0072\u006f\u0064\u0065")
		}
		if _, _acg = _cfgc(_fccb, _aebb, _bbg); _acg != nil {
			return nil, _e.Wrap(_acg, _fdfe, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0065\u0072\u006fd\u0065")
		}
	}
	if _abba, _eacg = _fccb.RemoveBorder(_dacda); _eacg != nil {
		return nil, _e.Wrap(_eacg, _fdfe, "\u0072e\u0067\u0075\u006c\u0061\u0072")
	}
	if _gedfe == nil {
		return _abba, nil
	}
	if _, _eacg = _cdcf(_gedfe, _abba); _eacg != nil {
		return nil, _eacg
	}
	return _gedfe, nil
}

const (
	PixSrc             RasterOperator = 0xc
	PixDst             RasterOperator = 0xa
	PixNotSrc          RasterOperator = 0x3
	PixNotDst          RasterOperator = 0x5
	PixClr             RasterOperator = 0x0
	PixSet             RasterOperator = 0xf
	PixSrcOrDst        RasterOperator = 0xe
	PixSrcAndDst       RasterOperator = 0x8
	PixSrcXorDst       RasterOperator = 0x6
	PixNotSrcOrDst     RasterOperator = 0xb
	PixNotSrcAndDst    RasterOperator = 0x2
	PixSrcOrNotDst     RasterOperator = 0xd
	PixSrcAndNotDst    RasterOperator = 0x4
	PixNotPixSrcOrDst  RasterOperator = 0x1
	PixNotPixSrcAndDst RasterOperator = 0x7
	PixNotPixSrcXorDst RasterOperator = 0x9
	PixPaint                          = PixSrcOrDst
	PixSubtract                       = PixNotSrcAndDst
	PixMask                           = PixSrcAndDst
)

func _adcba(_acdc, _bbcf, _cbdcb *Bitmap, _dbda int) (*Bitmap, error) {
	const _bcda = "\u0073\u0065\u0065\u0064\u0046\u0069\u006c\u006c\u0042i\u006e\u0061\u0072\u0079"
	if _bbcf == nil {
		return nil, _e.Error(_bcda, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _cbdcb == nil {
		return nil, _e.Error(_bcda, "'\u006da\u0073\u006b\u0027\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _dbda != 4 && _dbda != 8 {
		return nil, _e.Error(_bcda, "\u0063\u006f\u006en\u0065\u0063\u0074\u0069v\u0069\u0074\u0079\u0020\u006e\u006f\u0074 \u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0034\u002c\u0038\u007d")
	}
	var _edddcd error
	_acdc, _edddcd = _cdcf(_acdc, _bbcf)
	if _edddcd != nil {
		return nil, _e.Wrap(_edddcd, _bcda, "\u0063o\u0070y\u0020\u0073\u006f\u0075\u0072c\u0065\u0020t\u006f\u0020\u0027\u0064\u0027")
	}
	_aggb := _bbcf.createTemplate()
	_cbdcb.setPadBits(0)
	for _bcdc := 0; _bcdc < _decfc; _bcdc++ {
		_aggb, _edddcd = _cdcf(_aggb, _acdc)
		if _edddcd != nil {
			return nil, _e.Wrapf(_edddcd, _bcda, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _bcdc)
		}
		if _edddcd = _fafaf(_acdc, _cbdcb, _dbda); _edddcd != nil {
			return nil, _e.Wrapf(_edddcd, _bcda, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _bcdc)
		}
		if _aggb.Equals(_acdc) {
			break
		}
	}
	return _acdc, nil
}
func _cdcf(_fdcd, _ccgd *Bitmap) (*Bitmap, error) {
	if _ccgd == nil {
		return nil, _e.Error("\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ccgd == _fdcd {
		return _fdcd, nil
	}
	if _fdcd == nil {
		_fdcd = _ccgd.createTemplate()
		copy(_fdcd.Data, _ccgd.Data)
		return _fdcd, nil
	}
	_feag := _fdcd.resizeImageData(_ccgd)
	if _feag != nil {
		return nil, _e.Wrap(_feag, "\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "")
	}
	_fdcd.Text = _ccgd.Text
	copy(_fdcd.Data, _ccgd.Data)
	return _fdcd, nil
}
func (_gbf *Bitmap) CountPixels() int { return _gbf.countPixels() }
func _dabc(_aafbf, _ebde, _fbbe *Bitmap) (*Bitmap, error) {
	const _ccb = "\u0073\u0075\u0062\u0074\u0072\u0061\u0063\u0074"
	if _ebde == nil {
		return nil, _e.Error(_ccb, "'\u0073\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _fbbe == nil {
		return nil, _e.Error(_ccb, "'\u0073\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	var _fgaf error
	switch {
	case _aafbf == _ebde:
		if _fgaf = _aafbf.RasterOperation(0, 0, _ebde.Width, _ebde.Height, PixNotSrcAndDst, _fbbe, 0, 0); _fgaf != nil {
			return nil, _e.Wrap(_fgaf, _ccb, "\u0064 \u003d\u003d\u0020\u0073\u0031")
		}
	case _aafbf == _fbbe:
		if _fgaf = _aafbf.RasterOperation(0, 0, _ebde.Width, _ebde.Height, PixNotSrcAndDst, _ebde, 0, 0); _fgaf != nil {
			return nil, _e.Wrap(_fgaf, _ccb, "\u0064 \u003d\u003d\u0020\u0073\u0032")
		}
	default:
		_aafbf, _fgaf = _cdcf(_aafbf, _ebde)
		if _fgaf != nil {
			return nil, _e.Wrap(_fgaf, _ccb, "")
		}
		if _fgaf = _aafbf.RasterOperation(0, 0, _ebde.Width, _ebde.Height, PixNotSrcAndDst, _fbbe, 0, 0); _fgaf != nil {
			return nil, _e.Wrap(_fgaf, _ccb, "\u0064e\u0066\u0061\u0075\u006c\u0074")
		}
	}
	return _aafbf, nil
}
func (_ccda *Selection) setOrigin(_dbcfc, _efadg int) { _ccda.Cy, _ccda.Cx = _dbcfc, _efadg }
func CorrelationScore(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_adcd float64, _egag error) {
	const _bbec = "\u0063\u006fr\u0072\u0065\u006ca\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065"
	if bm1 == nil || bm2 == nil {
		return 0, _e.Error(_bbec, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0062\u0069\u0074ma\u0070\u0073")
	}
	if tab == nil {
		return 0, _e.Error(_bbec, "\u0027\u0074\u0061\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if area1 <= 0 || area2 <= 0 {
		return 0, _e.Error(_bbec, "\u0061\u0072\u0065\u0061s\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067r\u0065a\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
	}
	_deae, _febb := bm1.Width, bm1.Height
	_dbfe, _fage := bm2.Width, bm2.Height
	_eae := _ggce(_deae - _dbfe)
	if _eae > maxDiffW {
		return 0, nil
	}
	_fgdf := _ggce(_febb - _fage)
	if _fgdf > maxDiffH {
		return 0, nil
	}
	var _aegd, _dbdf int
	if delX >= 0 {
		_aegd = int(delX + 0.5)
	} else {
		_aegd = int(delX - 0.5)
	}
	if delY >= 0 {
		_dbdf = int(delY + 0.5)
	} else {
		_dbdf = int(delY - 0.5)
	}
	_ccba := _aefd(_dbdf, 0)
	_cbbf := _gcg(_fage+_dbdf, _febb)
	_gdff := bm1.RowStride * _ccba
	_edfa := bm2.RowStride * (_ccba - _dbdf)
	_bgbc := _aefd(_aegd, 0)
	_abad := _gcg(_dbfe+_aegd, _deae)
	_fdba := bm2.RowStride
	var _gcag, _ddda int
	if _aegd >= 8 {
		_gcag = _aegd >> 3
		_gdff += _gcag
		_bgbc -= _gcag << 3
		_abad -= _gcag << 3
		_aegd &= 7
	} else if _aegd <= -8 {
		_ddda = -((_aegd + 7) >> 3)
		_edfa += _ddda
		_fdba -= _ddda
		_aegd += _ddda << 3
	}
	if _bgbc >= _abad || _ccba >= _cbbf {
		return 0, nil
	}
	_bbcd := (_abad + 7) >> 3
	var (
		_ede, _bcfb, _afcb  byte
		_cgge, _gcad, _geea int
	)
	switch {
	case _aegd == 0:
		for _geea = _ccba; _geea < _cbbf; _geea, _gdff, _edfa = _geea+1, _gdff+bm1.RowStride, _edfa+bm2.RowStride {
			for _gcad = 0; _gcad < _bbcd; _gcad++ {
				_afcb = bm1.Data[_gdff+_gcad] & bm2.Data[_edfa+_gcad]
				_cgge += tab[_afcb]
			}
		}
	case _aegd > 0:
		if _fdba < _bbcd {
			for _geea = _ccba; _geea < _cbbf; _geea, _gdff, _edfa = _geea+1, _gdff+bm1.RowStride, _edfa+bm2.RowStride {
				_ede, _bcfb = bm1.Data[_gdff], bm2.Data[_edfa]>>uint(_aegd)
				_afcb = _ede & _bcfb
				_cgge += tab[_afcb]
				for _gcad = 1; _gcad < _fdba; _gcad++ {
					_ede, _bcfb = bm1.Data[_gdff+_gcad], (bm2.Data[_edfa+_gcad]>>uint(_aegd))|(bm2.Data[_edfa+_gcad-1]<<uint(8-_aegd))
					_afcb = _ede & _bcfb
					_cgge += tab[_afcb]
				}
				_ede = bm1.Data[_gdff+_gcad]
				_bcfb = bm2.Data[_edfa+_gcad-1] << uint(8-_aegd)
				_afcb = _ede & _bcfb
				_cgge += tab[_afcb]
			}
		} else {
			for _geea = _ccba; _geea < _cbbf; _geea, _gdff, _edfa = _geea+1, _gdff+bm1.RowStride, _edfa+bm2.RowStride {
				_ede, _bcfb = bm1.Data[_gdff], bm2.Data[_edfa]>>uint(_aegd)
				_afcb = _ede & _bcfb
				_cgge += tab[_afcb]
				for _gcad = 1; _gcad < _bbcd; _gcad++ {
					_ede = bm1.Data[_gdff+_gcad]
					_bcfb = (bm2.Data[_edfa+_gcad] >> uint(_aegd)) | (bm2.Data[_edfa+_gcad-1] << uint(8-_aegd))
					_afcb = _ede & _bcfb
					_cgge += tab[_afcb]
				}
			}
		}
	default:
		if _bbcd < _fdba {
			for _geea = _ccba; _geea < _cbbf; _geea, _gdff, _edfa = _geea+1, _gdff+bm1.RowStride, _edfa+bm2.RowStride {
				for _gcad = 0; _gcad < _bbcd; _gcad++ {
					_ede = bm1.Data[_gdff+_gcad]
					_bcfb = bm2.Data[_edfa+_gcad] << uint(-_aegd)
					_bcfb |= bm2.Data[_edfa+_gcad+1] >> uint(8+_aegd)
					_afcb = _ede & _bcfb
					_cgge += tab[_afcb]
				}
			}
		} else {
			for _geea = _ccba; _geea < _cbbf; _geea, _gdff, _edfa = _geea+1, _gdff+bm1.RowStride, _edfa+bm2.RowStride {
				for _gcad = 0; _gcad < _bbcd-1; _gcad++ {
					_ede = bm1.Data[_gdff+_gcad]
					_bcfb = bm2.Data[_edfa+_gcad] << uint(-_aegd)
					_bcfb |= bm2.Data[_edfa+_gcad+1] >> uint(8+_aegd)
					_afcb = _ede & _bcfb
					_cgge += tab[_afcb]
				}
				_ede = bm1.Data[_gdff+_gcad]
				_bcfb = bm2.Data[_edfa+_gcad] << uint(-_aegd)
				_afcb = _ede & _bcfb
				_cgge += tab[_afcb]
			}
		}
	}
	_adcd = float64(_cgge) * float64(_cgge) / (float64(area1) * float64(area2))
	return _adcd, nil
}
func (_gefe *Bitmaps) selectByIndicator(_abgd *_ac.NumSlice) (_fabbc *Bitmaps, _faga error) {
	const _gccge = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u0073\u0065\u006c\u0065c\u0074B\u0079I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _gefe == nil {
		return nil, _e.Error(_gccge, "\u0027\u0062\u0027 b\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _abgd == nil {
		return nil, _e.Error(_gccge, "'\u006e\u0061\u0027\u0020\u0069\u006ed\u0069\u0063\u0061\u0074\u006f\u0072\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	if len(_gefe.Values) == 0 {
		return _gefe, nil
	}
	if len(*_abgd) != len(_gefe.Values) {
		return nil, _e.Errorf(_gccge, "\u006ea\u0020\u006ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064,\u0020\u0069\u0073\u0020\u0064\u0069\u0066\u0066\u0065re\u006e\u0074\u0020t\u0068\u0061n\u0020\u0062\u0069\u0074\u006d\u0061p\u0073\u003a \u0025\u0064", len(*_abgd), len(_gefe.Values))
	}
	var _badg, _bbba, _beedd int
	for _bbba = 0; _bbba < len(*_abgd); _bbba++ {
		if _badg, _faga = _abgd.GetInt(_bbba); _faga != nil {
			return nil, _e.Wrap(_faga, _gccge, "f\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		if _badg == 1 {
			_beedd++
		}
	}
	if _beedd == len(_gefe.Values) {
		return _gefe, nil
	}
	_fabbc = &Bitmaps{}
	_gfffe := len(_gefe.Values) == len(_gefe.Boxes)
	for _bbba = 0; _bbba < len(*_abgd); _bbba++ {
		if _badg = int((*_abgd)[_bbba]); _badg == 0 {
			continue
		}
		_fabbc.Values = append(_fabbc.Values, _gefe.Values[_bbba])
		if _gfffe {
			_fabbc.Boxes = append(_fabbc.Boxes, _gefe.Boxes[_bbba])
		}
	}
	return _fabbc, nil
}
func _ggce(_cafcf int) int {
	if _cafcf < 0 {
		return -_cafcf
	}
	return _cafcf
}
func (_dbfg *byWidth) Swap(i, j int) {
	_dbfg.Values[i], _dbfg.Values[j] = _dbfg.Values[j], _dbfg.Values[i]
	if _dbfg.Boxes != nil {
		_dbfg.Boxes[i], _dbfg.Boxes[j] = _dbfg.Boxes[j], _dbfg.Boxes[i]
	}
}
func (_gdf *Bitmap) InverseData() { _gdf.inverseData() }
