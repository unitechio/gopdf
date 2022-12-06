package bitmap

import (
	_cc "encoding/binary"
	_aa "image"
	_ca "math"
	_a "sort"
	_dc "strings"
	_ba "testing"

	_gb "bitbucket.org/shenghui0779/gopdf/common"
	_c "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_ac "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_dd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_g "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_b "github.com/stretchr/testify/require"
)

func (_acbd *Selection) setOrigin(_acbba, _edcb int)              { _acbd.Cy, _acbd.Cx = _acbba, _edcb }
func DilateBrick(d, s *Bitmap, hSize, vSize int) (*Bitmap, error) { return _geaed(d, s, hSize, vSize) }
func (_dbdga *Bitmap) connComponentsBitmapsBB(_baefg *Bitmaps, _fcba int) (_fffdc *Boxes, _ecde error) {
	const _gegd = "\u0063\u006f\u006enC\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0069\u0074\u006d\u0061\u0070\u0073\u0042\u0042"
	if _fcba != 4 && _fcba != 8 {
		return nil, _g.Error(_gegd, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _baefg == nil {
		return nil, _g.Error(_gegd, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if len(_baefg.Values) > 0 {
		return nil, _g.Error(_gegd, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u006fn\u002d\u0065\u006d\u0070\u0074\u0079\u0020\u0042\u0069\u0074m\u0061\u0070\u0073")
	}
	if _dbdga.Zero() {
		return &Boxes{}, nil
	}
	var (
		_edef, _gdgf, _adcb, _baac *Bitmap
	)
	_dbdga.setPadBits(0)
	if _edef, _ecde = _feea(nil, _dbdga); _ecde != nil {
		return nil, _g.Wrap(_ecde, _gegd, "\u0062\u006d\u0031")
	}
	if _gdgf, _ecde = _feea(nil, _dbdga); _ecde != nil {
		return nil, _g.Wrap(_ecde, _gegd, "\u0062\u006d\u0032")
	}
	_cfcc := &_dd.Stack{}
	_cfcc.Aux = &_dd.Stack{}
	_fffdc = &Boxes{}
	var (
		_efd, _egdb int
		_cccf       _aa.Point
		_dgf        bool
		_gfaf       *_aa.Rectangle
	)
	for {
		if _cccf, _dgf, _ecde = _edef.nextOnPixel(_efd, _egdb); _ecde != nil {
			return nil, _g.Wrap(_ecde, _gegd, "")
		}
		if !_dgf {
			break
		}
		if _gfaf, _ecde = _gccg(_edef, _cfcc, _cccf.X, _cccf.Y, _fcba); _ecde != nil {
			return nil, _g.Wrap(_ecde, _gegd, "")
		}
		if _ecde = _fffdc.Add(_gfaf); _ecde != nil {
			return nil, _g.Wrap(_ecde, _gegd, "")
		}
		if _adcb, _ecde = _edef.clipRectangle(_gfaf, nil); _ecde != nil {
			return nil, _g.Wrap(_ecde, _gegd, "\u0062\u006d\u0033")
		}
		if _baac, _ecde = _gdgf.clipRectangle(_gfaf, nil); _ecde != nil {
			return nil, _g.Wrap(_ecde, _gegd, "\u0062\u006d\u0034")
		}
		if _, _ecde = _fcbc(_adcb, _adcb, _baac); _ecde != nil {
			return nil, _g.Wrap(_ecde, _gegd, "\u0062m\u0033\u0020\u005e\u0020\u0062\u006d4")
		}
		if _ecde = _gdgf.RasterOperation(_gfaf.Min.X, _gfaf.Min.Y, _gfaf.Dx(), _gfaf.Dy(), PixSrcXorDst, _adcb, 0, 0); _ecde != nil {
			return nil, _g.Wrap(_ecde, _gegd, "\u0062\u006d\u0032\u0020\u002d\u0058\u004f\u0052\u002d>\u0020\u0062\u006d\u0033")
		}
		_baefg.AddBitmap(_adcb)
		_efd = _cccf.X
		_egdb = _cccf.Y
	}
	_baefg.Boxes = *_fffdc
	return _fffdc, nil
}
func _aaad(_dbdg, _fbbd *Bitmap, _eeda, _cabaf, _bbcd uint, _badd, _gebf int, _cbba bool, _gfba, _efb int) error {
	for _efeg := _badd; _efeg < _gebf; _efeg++ {
		if _gfba+1 < len(_dbdg.Data) {
			_cdge := _efeg+1 == _gebf
			_gaac, _daff := _dbdg.GetByte(_gfba)
			if _daff != nil {
				return _daff
			}
			_gfba++
			_gaac <<= _eeda
			_fgc, _daff := _dbdg.GetByte(_gfba)
			if _daff != nil {
				return _daff
			}
			_fgc >>= _cabaf
			_decg := _gaac | _fgc
			if _cdge && !_cbba {
				_decg = _cddg(_bbcd, _decg)
			}
			_daff = _fbbd.SetByte(_efb, _decg)
			if _daff != nil {
				return _daff
			}
			_efb++
			if _cdge && _cbba {
				_agdc, _gbfe := _dbdg.GetByte(_gfba)
				if _gbfe != nil {
					return _gbfe
				}
				_agdc <<= _eeda
				_decg = _cddg(_bbcd, _agdc)
				if _gbfe = _fbbd.SetByte(_efb, _decg); _gbfe != nil {
					return _gbfe
				}
			}
			continue
		}
		_agfg, _feedc := _dbdg.GetByte(_gfba)
		if _feedc != nil {
			_gb.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0061\u0074\u003a\u0020%\u0064\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0073", _gfba, _feedc)
			return _feedc
		}
		_agfg <<= _eeda
		_gfba++
		_feedc = _fbbd.SetByte(_efb, _agfg)
		if _feedc != nil {
			return _feedc
		}
		_efb++
	}
	return nil
}
func (_daggc *ClassedPoints) SortByY() { _daggc._bacf = _daggc.ySortFunction(); _a.Sort(_daggc) }
func _cggf(_efdf *Bitmap, _aeea, _fgge int, _aaag, _aebb int, _daca RasterOperator, _eeea *Bitmap, _gde, _cbbdc int) error {
	var _ggfb, _geda, _aaf, _baag int
	if _aeea < 0 {
		_gde -= _aeea
		_aaag += _aeea
		_aeea = 0
	}
	if _gde < 0 {
		_aeea -= _gde
		_aaag += _gde
		_gde = 0
	}
	_ggfb = _aeea + _aaag - _efdf.Width
	if _ggfb > 0 {
		_aaag -= _ggfb
	}
	_geda = _gde + _aaag - _eeea.Width
	if _geda > 0 {
		_aaag -= _geda
	}
	if _fgge < 0 {
		_cbbdc -= _fgge
		_aebb += _fgge
		_fgge = 0
	}
	if _cbbdc < 0 {
		_fgge -= _cbbdc
		_aebb += _cbbdc
		_cbbdc = 0
	}
	_aaf = _fgge + _aebb - _efdf.Height
	if _aaf > 0 {
		_aebb -= _aaf
	}
	_baag = _cbbdc + _aebb - _eeea.Height
	if _baag > 0 {
		_aebb -= _baag
	}
	if _aaag <= 0 || _aebb <= 0 {
		return nil
	}
	var _becg error
	switch {
	case _aeea&7 == 0 && _gde&7 == 0:
		_becg = _fegb(_efdf, _aeea, _fgge, _aaag, _aebb, _daca, _eeea, _gde, _cbbdc)
	case _aeea&7 == _gde&7:
		_becg = _ggda(_efdf, _aeea, _fgge, _aaag, _aebb, _daca, _eeea, _gde, _cbbdc)
	default:
		_becg = _ccbfd(_efdf, _aeea, _fgge, _aaag, _aebb, _daca, _eeea, _gde, _cbbdc)
	}
	if _becg != nil {
		return _g.Wrap(_becg, "r\u0061\u0073\u0074\u0065\u0072\u004f\u0070\u004c\u006f\u0077", "")
	}
	return nil
}
func (_daed *Bitmap) setFourBytes(_eaf int, _dfg uint32) error {
	if _eaf+3 > len(_daed.Data)-1 {
		return _g.Errorf("\u0073\u0065\u0074F\u006f\u0075\u0072\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _eaf)
	}
	_daed.Data[_eaf] = byte((_dfg & 0xff000000) >> 24)
	_daed.Data[_eaf+1] = byte((_dfg & 0xff0000) >> 16)
	_daed.Data[_eaf+2] = byte((_dfg & 0xff00) >> 8)
	_daed.Data[_eaf+3] = byte(_dfg & 0xff)
	return nil
}
func (_bea *Bitmap) SizesEqual(s *Bitmap) bool {
	if _bea == s {
		return true
	}
	if _bea.Width != s.Width || _bea.Height != s.Height {
		return false
	}
	return true
}
func _cec(_ecg *Bitmap, _gff ...int) (_ccb *Bitmap, _fea error) {
	const _dad = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0043\u0061\u0073\u0063\u0061\u0064\u0065"
	if _ecg == nil {
		return nil, _g.Error(_dad, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_gff) == 0 || len(_gff) > 4 {
		return nil, _g.Error(_dad, "t\u0068\u0065\u0072\u0065\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0074\u0020\u006cea\u0073\u0074\u0020\u006fn\u0065\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020mo\u0073\u0074 \u0034\u0020\u006c\u0065\u0076\u0065\u006c\u0073")
	}
	if _gff[0] <= 0 {
		_gb.Log.Debug("\u006c\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030 \u002d\u0020\u006e\u006f\u0020\u0072\u0065\u0064\u0075\u0063t\u0069\u006f\u006e")
		_ccb, _fea = _feea(nil, _ecg)
		if _fea != nil {
			return nil, _g.Wrap(_fea, _dad, "l\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030")
		}
		return _ccb, nil
	}
	_fff := _bgba()
	_ccb = _ecg
	for _dcf, _cfd := range _gff {
		if _cfd <= 0 {
			break
		}
		_ccb, _fea = _aba(_ccb, _cfd, _fff)
		if _fea != nil {
			return nil, _g.Wrapf(_fea, _dad, "\u006c\u0065\u0076\u0065\u006c\u0025\u0064\u0020\u0072\u0065\u0064\u0075c\u0074\u0069\u006f\u006e", _dcf)
		}
	}
	return _ccb, nil
}
func _aee(_eaba *Bitmap, _bede, _cbd, _egfdc, _dafa int, _agea RasterOperator, _cdf *Bitmap, _abge, _ccg int) error {
	const _eadf = "\u0072a\u0073t\u0065\u0072\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e"
	if _eaba == nil {
		return _g.Error(_eadf, "\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _agea == PixDst {
		return nil
	}
	switch _agea {
	case PixClr, PixSet, PixNotDst:
		_fbfd(_eaba, _bede, _cbd, _egfdc, _dafa, _agea)
		return nil
	}
	if _cdf == nil {
		_gb.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _g.Error(_eadf, "\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _eeag := _cggf(_eaba, _bede, _cbd, _egfdc, _dafa, _agea, _cdf, _abge, _ccg); _eeag != nil {
		return _g.Wrap(_eeag, _eadf, "")
	}
	return nil
}
func (_dcccb *Bitmap) ConnComponents(bms *Bitmaps, connectivity int) (_agae *Boxes, _dbae error) {
	const _dccgb = "B\u0069\u0074\u006d\u0061p.\u0043o\u006e\u006e\u0043\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0073"
	if _dcccb == nil {
		return nil, _g.Error(_dccgb, "\u0070r\u006f\u0076\u0069\u0064e\u0064\u0020\u0065\u006d\u0070t\u0079 \u0027b\u0027\u0020\u0062\u0069\u0074\u006d\u0061p")
	}
	if connectivity != 4 && connectivity != 8 {
		return nil, _g.Error(_dccgb, "\u0063\u006f\u006ene\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006e\u006f\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038")
	}
	if bms == nil {
		if _agae, _dbae = _dcccb.connComponentsBB(connectivity); _dbae != nil {
			return nil, _g.Wrap(_dbae, _dccgb, "")
		}
	} else {
		if _agae, _dbae = _dcccb.connComponentsBitmapsBB(bms, connectivity); _dbae != nil {
			return nil, _g.Wrap(_dbae, _dccgb, "")
		}
	}
	return _agae, nil
}
func TstWordBitmapWithSpaces(t *_ba.T, scale ...int) *Bitmap {
	_cbfd := 1
	if len(scale) > 0 {
		_cbfd = scale[0]
	}
	_ebec := 3
	_bebgb := 9 + 7 + 15 + 2*_ebec + 2*_ebec
	_gccaa := 5 + _ebec + 5 + 2*_ebec
	_gage := New(_bebgb*_cbfd, _gccaa*_cbfd)
	_dfgfc := &Bitmaps{}
	var _bbcdg *int
	_ebec *= _cbfd
	_gaga := _ebec
	_bbcdg = &_gaga
	_dcceb := _ebec
	_eeefa := TstDSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstOSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, _ebec)
	_eeefa = TstISymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstTSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, _ebec)
	_eeefa = TstNSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstOSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstWSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 0)
	*_bbcdg = _ebec
	_dcceb = 5*_cbfd + _ebec
	_eeefa = TstOSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstRSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, _ebec)
	_eeefa = TstNSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstESymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstVSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstESymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 1*_cbfd)
	_eeefa = TstRSymbol(t, scale...)
	TstAddSymbol(t, _dfgfc, _eeefa, _bbcdg, _dcceb, 0)
	TstWriteSymbols(t, _dfgfc, _gage)
	return _gage
}
func _ge(_f, _be *Bitmap) (_ddc error) {
	const _ccd = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0032"
	_dcc := _be.RowStride
	_de := _f.RowStride
	var (
		_dee                     byte
		_fg                      uint16
		_bf, _e, _ae, _aec, _aab int
	)
	for _ae = 0; _ae < _be.Height; _ae++ {
		_bf = _ae * _dcc
		_e = 2 * _ae * _de
		for _aec = 0; _aec < _dcc; _aec++ {
			_dee = _be.Data[_bf+_aec]
			_fg = _eege[_dee]
			_aab = _e + _aec*2
			if _f.RowStride != _be.RowStride*2 && (_aec+1)*2 > _f.RowStride {
				_ddc = _f.SetByte(_aab, byte(_fg>>8))
			} else {
				_ddc = _f.setTwoBytes(_aab, _fg)
			}
			if _ddc != nil {
				return _g.Wrap(_ddc, _ccd, "")
			}
		}
		for _aec = 0; _aec < _de; _aec++ {
			_aab = _e + _de + _aec
			_dee = _f.Data[_e+_aec]
			if _ddc = _f.SetByte(_aab, _dee); _ddc != nil {
				return _g.Wrapf(_ddc, _ccd, "c\u006f\u0070\u0079\u0020\u0064\u006fu\u0062\u006c\u0065\u0064\u0020\u006ci\u006e\u0065\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0042\u0079\u0074\u0065\u003a\u0020\u0027\u0025d\u0027", _e+_aec, _e+_de+_aec)
			}
		}
	}
	return nil
}
func TstVSymbol(t *_ba.T, scale ...int) *Bitmap {
	_bcdc, _dfeae := NewWithData(5, 5, []byte{0x88, 0x88, 0x88, 0x50, 0x20})
	_b.NoError(t, _dfeae)
	return TstGetScaledSymbol(t, _bcdc, scale...)
}
func (_eda *Bitmap) Zero() bool {
	_fdgd := _eda.Width / 8
	_dafde := _eda.Width & 7
	var _ceab byte
	if _dafde != 0 {
		_ceab = byte(0xff << uint(8-_dafde))
	}
	var _efa, _gceg, _fbf int
	for _gceg = 0; _gceg < _eda.Height; _gceg++ {
		_efa = _eda.RowStride * _gceg
		for _fbf = 0; _fbf < _fdgd; _fbf, _efa = _fbf+1, _efa+1 {
			if _eda.Data[_efa] != 0 {
				return false
			}
		}
		if _dafde > 0 {
			if _eda.Data[_efa]&_ceab != 0 {
				return false
			}
		}
	}
	return true
}

const (
	AsymmetricMorphBC BoundaryCondition = iota
	SymmetricMorphBC
)

func _aage(_ffa, _acbgd int) int {
	if _ffa > _acbgd {
		return _ffa
	}
	return _acbgd
}
func init() {
	for _bad := 0; _bad < 256; _bad++ {
		_ebc[_bad] = uint8(_bad&0x1) + (uint8(_bad>>1) & 0x1) + (uint8(_bad>>2) & 0x1) + (uint8(_bad>>3) & 0x1) + (uint8(_bad>>4) & 0x1) + (uint8(_bad>>5) & 0x1) + (uint8(_bad>>6) & 0x1) + (uint8(_bad>>7) & 0x1)
	}
}
func (_dge *Bitmap) GetUnpaddedData() ([]byte, error) {
	_ecb := uint(_dge.Width & 0x07)
	if _ecb == 0 {
		return _dge.Data, nil
	}
	_dbde := _dge.Width * _dge.Height
	if _dbde%8 != 0 {
		_dbde >>= 3
		_dbde++
	} else {
		_dbde >>= 3
	}
	_gaff := make([]byte, _dbde)
	_fgd := _c.NewWriterMSB(_gaff)
	const _bab = "\u0047e\u0074U\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	for _gec := 0; _gec < _dge.Height; _gec++ {
		for _fdf := 0; _fdf < _dge.RowStride; _fdf++ {
			_becc := _dge.Data[_gec*_dge.RowStride+_fdf]
			if _fdf != _dge.RowStride-1 {
				_bfg := _fgd.WriteByte(_becc)
				if _bfg != nil {
					return nil, _g.Wrap(_bfg, _bab, "")
				}
				continue
			}
			for _dfd := uint(0); _dfd < _ecb; _dfd++ {
				_gdf := _fgd.WriteBit(int(_becc >> (7 - _dfd) & 0x01))
				if _gdf != nil {
					return nil, _g.Wrap(_gdf, _bab, "")
				}
			}
		}
	}
	return _gaff, nil
}
func (_gca *Bitmap) Equals(s *Bitmap) bool {
	if len(_gca.Data) != len(s.Data) || _gca.Width != s.Width || _gca.Height != s.Height {
		return false
	}
	for _bcgf := 0; _bcgf < _gca.Height; _bcgf++ {
		_cfg := _bcgf * _gca.RowStride
		for _bfef := 0; _bfef < _gca.RowStride; _bfef++ {
			if _gca.Data[_cfg+_bfef] != s.Data[_cfg+_bfef] {
				return false
			}
		}
	}
	return true
}
func (_eaca *Bitmap) InverseData()          { _eaca.inverseData() }
func (_fcgg *Bitmap) SetPadBits(value int)  { _fcgg.setPadBits(value) }
func (_efdg *Points) AddPoint(x, y float32) { *_efdg = append(*_efdg, Point{x, y}) }
func _fcbc(_fadf, _gcga, _gaba *Bitmap) (*Bitmap, error) {
	const _gggg = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0078\u006f\u0072"
	if _gcga == nil {
		return nil, _g.Error(_gggg, "'\u0062\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _gaba == nil {
		return nil, _g.Error(_gggg, "'\u0062\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _fadf == _gaba {
		return nil, _g.Error(_gggg, "'\u0064\u0027\u0020\u003d\u003d\u0020\u0027\u0062\u0032\u0027")
	}
	if !_gcga.SizesEqual(_gaba) {
		_gb.Log.Debug("\u0025s\u0020\u002d \u0042\u0069\u0074\u006da\u0070\u0020\u0027b\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074 e\u0071\u0075\u0061l\u0020\u0073i\u007a\u0065\u0020\u0077\u0069\u0074h\u0020\u0027b\u0032\u0027", _gggg)
	}
	var _daga error
	if _fadf, _daga = _feea(_fadf, _gcga); _daga != nil {
		return nil, _g.Wrap(_daga, _gggg, "\u0063\u0061n\u0027\u0074\u0020c\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027")
	}
	if _daga = _fadf.RasterOperation(0, 0, _fadf.Width, _fadf.Height, PixSrcXorDst, _gaba, 0, 0); _daga != nil {
		return nil, _g.Wrap(_daga, _gggg, "")
	}
	return _fadf, nil
}
func _bdga(_ecae *Bitmap, _defb *Bitmap, _edfaf int) (_bfcdc error) {
	const _gafa = "\u0073\u0065\u0065\u0064\u0066\u0069\u006c\u006c\u0042\u0069\u006e\u0061r\u0079\u004c\u006f\u0077"
	_fdeg := _efag(_ecae.Height, _defb.Height)
	_fddc := _efag(_ecae.RowStride, _defb.RowStride)
	switch _edfaf {
	case 4:
		_bfcdc = _aaba(_ecae, _defb, _fdeg, _fddc)
	case 8:
		_bfcdc = _cbdc(_ecae, _defb, _fdeg, _fddc)
	default:
		return _g.Errorf(_gafa, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0034\u0020\u006fr\u0020\u0038\u0020\u002d\u0020i\u0073\u003a \u0027\u0025\u0064\u0027", _edfaf)
	}
	if _bfcdc != nil {
		return _g.Wrap(_bfcdc, _gafa, "")
	}
	return nil
}

type MorphProcess struct {
	Operation MorphOperation
	Arguments []int
}

func _egcag(_fbbga *_dd.Stack, _eeab, _geab, _dfeec, _edcaf, _gcgee int, _fdbb *_aa.Rectangle) (_gcbfe error) {
	const _gcff = "\u0070\u0075\u0073\u0068\u0046\u0069\u006c\u006c\u0053\u0065\u0067m\u0065\u006e\u0074\u0042\u006f\u0075\u006e\u0064\u0069\u006eg\u0042\u006f\u0078"
	if _fbbga == nil {
		return _g.Error(_gcff, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _fdbb == nil {
		return _g.Error(_gcff, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0069\u006da\u0067e\u002e\u0052\u0065\u0063\u0074\u0061\u006eg\u006c\u0065")
	}
	_fdbb.Min.X = _dd.Min(_fdbb.Min.X, _eeab)
	_fdbb.Max.X = _dd.Max(_fdbb.Max.X, _geab)
	_fdbb.Min.Y = _dd.Min(_fdbb.Min.Y, _dfeec)
	_fdbb.Max.Y = _dd.Max(_fdbb.Max.Y, _dfeec)
	if !(_dfeec+_edcaf >= 0 && _dfeec+_edcaf <= _gcgee) {
		return nil
	}
	if _fbbga.Aux == nil {
		return _g.Error(_gcff, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	var _fedfe *fillSegment
	_eecg, _gcca := _fbbga.Aux.Pop()
	if _gcca {
		if _fedfe, _gcca = _eecg.(*fillSegment); !_gcca {
			return _g.Error(_gcff, "a\u0075\u0078\u0053\u0074\u0061\u0063k\u0020\u0064\u0061\u0074\u0061\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061 \u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065n\u0074")
		}
	} else {
		_fedfe = &fillSegment{}
	}
	_fedfe._dgfa = _eeab
	_fedfe._dade = _geab
	_fedfe._ecfb = _dfeec
	_fedfe._cgaee = _edcaf
	_fbbga.Push(_fedfe)
	return nil
}
func (_gdd *Bitmap) SetDefaultPixel() {
	for _ffeg := range _gdd.Data {
		_gdd.Data[_ffeg] = byte(0xff)
	}
}
func _cffff(_fdffa *Bitmap, _gecg int) (*Bitmap, error) {
	const _eccd = "\u0065x\u0070a\u006e\u0064\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0065"
	if _fdffa == nil {
		return nil, _g.Error(_eccd, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _gecg <= 0 {
		return nil, _g.Error(_eccd, "i\u006e\u0076\u0061\u006cid\u0020f\u0061\u0063\u0074\u006f\u0072 \u002d\u0020\u003c\u003d\u0020\u0030")
	}
	if _gecg == 1 {
		_gagg, _daac := _feea(nil, _fdffa)
		if _daac != nil {
			return nil, _g.Wrap(_daac, _eccd, "\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u003d\u0020\u0031")
		}
		return _gagg, nil
	}
	_eeddb, _ggae := _gac(_fdffa, _gecg, _gecg)
	if _ggae != nil {
		return nil, _g.Wrap(_ggae, _eccd, "")
	}
	return _eeddb, nil
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
const (
	SelDontCare SelectionValue = iota
	SelHit
	SelMiss
)

func (_dcga *ClassedPoints) GetIntYByClass(i int) (int, error) {
	const _dfbf = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _dcga.IntSlice.Size() {
		return 0, _g.Errorf(_dfbf, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_dcga.YAtIndex(i)), nil
}
func (_aef *Bitmaps) SortByHeight() { _abea := (*byHeight)(_aef); _a.Sort(_abea) }
func (_fgde *Bitmaps) String() string {
	_gcdf := _dc.Builder{}
	for _, _fbae := range _fgde.Values {
		_gcdf.WriteString(_fbae.String())
		_gcdf.WriteRune('\n')
	}
	return _gcdf.String()
}
func (_fecf *Bitmap) inverseData() {
	if _aecfd := _fecf.RasterOperation(0, 0, _fecf.Width, _fecf.Height, PixNotDst, nil, 0, 0); _aecfd != nil {
		_gb.Log.Debug("\u0049n\u0076\u0065\u0072\u0073e\u0020\u0064\u0061\u0074\u0061 \u0066a\u0069l\u0065\u0064\u003a\u0020\u0027\u0025\u0076'", _aecfd)
	}
	if _fecf.Color == Chocolate {
		_fecf.Color = Vanilla
	} else {
		_fecf.Color = Chocolate
	}
}
func _feea(_dgea, _efe *Bitmap) (*Bitmap, error) {
	if _efe == nil {
		return nil, _g.Error("\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _efe == _dgea {
		return _dgea, nil
	}
	if _dgea == nil {
		_dgea = _efe.createTemplate()
		copy(_dgea.Data, _efe.Data)
		return _dgea, nil
	}
	_dccce := _dgea.resizeImageData(_efe)
	if _dccce != nil {
		return nil, _g.Wrap(_dccce, "\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "")
	}
	_dgea.Text = _efe.Text
	copy(_dgea.Data, _efe.Data)
	return _dgea, nil
}
func MakePixelSumTab8() []int     { return _gdfg() }
func (_cfgfe *Bitmaps) Size() int { return len(_cfgfe.Values) }

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

func (_ffbdc *byWidth) Swap(i, j int) {
	_ffbdc.Values[i], _ffbdc.Values[j] = _ffbdc.Values[j], _ffbdc.Values[i]
	if _ffbdc.Boxes != nil {
		_ffbdc.Boxes[i], _ffbdc.Boxes[j] = _ffbdc.Boxes[j], _ffbdc.Boxes[i]
	}
}
func CorrelationScoreThresholded(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab, downcount []int, scoreThreshold float32) (bool, error) {
	const _egaeb = "C\u006f\u0072\u0072\u0065\u006c\u0061t\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054h\u0072\u0065\u0073h\u006fl\u0064\u0065\u0064"
	if bm1 == nil {
		return false, _g.Error(_egaeb, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d1\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if bm2 == nil {
		return false, _g.Error(_egaeb, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d2\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if area1 <= 0 || area2 <= 0 {
		return false, _g.Error(_egaeb, "c\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0053\u0063\u006f\u0072\u0065\u0054\u0068re\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u002d\u0020\u0061\u0072\u0065\u0061s \u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u003e\u0020\u0030")
	}
	if downcount == nil {
		return false, _g.Error(_egaeb, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u006f\u0020\u0027\u0064\u006f\u0077\u006e\u0063\u006f\u0075\u006e\u0074\u0027")
	}
	if tab == nil {
		return false, _g.Error(_egaeb, "p\u0072\u006f\u0076\u0069de\u0064 \u006e\u0069\u006c\u0020\u0027s\u0075\u006d\u0074\u0061\u0062\u0027")
	}
	_badb, _geff := bm1.Width, bm1.Height
	_fddb, _gcbfa := bm2.Width, bm2.Height
	if _dd.Abs(_badb-_fddb) > maxDiffW {
		return false, nil
	}
	if _dd.Abs(_geff-_gcbfa) > maxDiffH {
		return false, nil
	}
	_ddbb := int(delX + _dd.Sign(delX)*0.5)
	_geca := int(delY + _dd.Sign(delY)*0.5)
	_cece := int(_ca.Ceil(_ca.Sqrt(float64(scoreThreshold) * float64(area1) * float64(area2))))
	_agcb := bm2.RowStride
	_dbfg := _aage(_geca, 0)
	_cfac := _efag(_gcbfa+_geca, _geff)
	_fdgbd := bm1.RowStride * _dbfg
	_egg := bm2.RowStride * (_dbfg - _geca)
	var _cedd int
	if _cfac <= _geff {
		_cedd = downcount[_cfac-1]
	}
	_fcc := _aage(_ddbb, 0)
	_febdf := _efag(_fddb+_ddbb, _badb)
	var _bbdd, _bdgc int
	if _ddbb >= 8 {
		_bbdd = _ddbb >> 3
		_fdgbd += _bbdd
		_fcc -= _bbdd << 3
		_febdf -= _bbdd << 3
		_ddbb &= 7
	} else if _ddbb <= -8 {
		_bdgc = -((_ddbb + 7) >> 3)
		_egg += _bdgc
		_agcb -= _bdgc
		_ddbb += _bdgc << 3
	}
	var (
		_eea, _cegb, _abb    int
		_geffb, _bfac, _eacb byte
	)
	if _fcc >= _febdf || _dbfg >= _cfac {
		return false, nil
	}
	_dfbg := (_febdf + 7) >> 3
	switch {
	case _ddbb == 0:
		for _cegb = _dbfg; _cegb < _cfac; _cegb, _fdgbd, _egg = _cegb+1, _fdgbd+bm1.RowStride, _egg+bm2.RowStride {
			for _abb = 0; _abb < _dfbg; _abb++ {
				_geffb = bm1.Data[_fdgbd+_abb] & bm2.Data[_egg+_abb]
				_eea += tab[_geffb]
			}
			if _eea >= _cece {
				return true, nil
			}
			if _ccec := _eea + downcount[_cegb] - _cedd; _ccec < _cece {
				return false, nil
			}
		}
	case _ddbb > 0 && _agcb < _dfbg:
		for _cegb = _dbfg; _cegb < _cfac; _cegb, _fdgbd, _egg = _cegb+1, _fdgbd+bm1.RowStride, _egg+bm2.RowStride {
			_bfac = bm1.Data[_fdgbd]
			_eacb = bm2.Data[_egg] >> uint(_ddbb)
			_geffb = _bfac & _eacb
			_eea += tab[_geffb]
			for _abb = 1; _abb < _agcb; _abb++ {
				_bfac = bm1.Data[_fdgbd+_abb]
				_eacb = bm2.Data[_egg+_abb]>>uint(_ddbb) | bm2.Data[_egg+_abb-1]<<uint(8-_ddbb)
				_geffb = _bfac & _eacb
				_eea += tab[_geffb]
			}
			_bfac = bm1.Data[_fdgbd+_abb]
			_eacb = bm2.Data[_egg+_abb-1] << uint(8-_ddbb)
			_geffb = _bfac & _eacb
			_eea += tab[_geffb]
			if _eea >= _cece {
				return true, nil
			} else if _eea+downcount[_cegb]-_cedd < _cece {
				return false, nil
			}
		}
	case _ddbb > 0 && _agcb >= _dfbg:
		for _cegb = _dbfg; _cegb < _cfac; _cegb, _fdgbd, _egg = _cegb+1, _fdgbd+bm1.RowStride, _egg+bm2.RowStride {
			_bfac = bm1.Data[_fdgbd]
			_eacb = bm2.Data[_egg] >> uint(_ddbb)
			_geffb = _bfac & _eacb
			_eea += tab[_geffb]
			for _abb = 1; _abb < _dfbg; _abb++ {
				_bfac = bm1.Data[_fdgbd+_abb]
				_eacb = bm2.Data[_egg+_abb] >> uint(_ddbb)
				_eacb |= bm2.Data[_egg+_abb-1] << uint(8-_ddbb)
				_geffb = _bfac & _eacb
				_eea += tab[_geffb]
			}
			if _eea >= _cece {
				return true, nil
			} else if _eea+downcount[_cegb]-_cedd < _cece {
				return false, nil
			}
		}
	case _dfbg < _agcb:
		for _cegb = _dbfg; _cegb < _cfac; _cegb, _fdgbd, _egg = _cegb+1, _fdgbd+bm1.RowStride, _egg+bm2.RowStride {
			for _abb = 0; _abb < _dfbg; _abb++ {
				_bfac = bm1.Data[_fdgbd+_abb]
				_eacb = bm2.Data[_egg+_abb] << uint(-_ddbb)
				_eacb |= bm2.Data[_egg+_abb+1] >> uint(8+_ddbb)
				_geffb = _bfac & _eacb
				_eea += tab[_geffb]
			}
			if _eea >= _cece {
				return true, nil
			} else if _effc := _eea + downcount[_cegb] - _cedd; _effc < _cece {
				return false, nil
			}
		}
	case _agcb >= _dfbg:
		for _cegb = _dbfg; _cegb < _cfac; _cegb, _fdgbd, _egg = _cegb+1, _fdgbd+bm1.RowStride, _egg+bm2.RowStride {
			for _abb = 0; _abb < _dfbg; _abb++ {
				_bfac = bm1.Data[_fdgbd+_abb]
				_eacb = bm2.Data[_egg+_abb] << uint(-_ddbb)
				_eacb |= bm2.Data[_egg+_abb+1] >> uint(8+_ddbb)
				_geffb = _bfac & _eacb
				_eea += tab[_geffb]
			}
			_bfac = bm1.Data[_fdgbd+_abb]
			_eacb = bm2.Data[_egg+_abb] << uint(-_ddbb)
			_geffb = _bfac & _eacb
			_eea += tab[_geffb]
			if _eea >= _cece {
				return true, nil
			} else if _eea+downcount[_cegb]-_cedd < _cece {
				return false, nil
			}
		}
	}
	_cee := float32(_eea) * float32(_eea) / (float32(area1) * float32(area2))
	if _cee >= scoreThreshold {
		_gb.Log.Trace("\u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064\u0020\u003c\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0020\u0025\u0064\u0020\u0062\u0075\u0074\u0020\u0073c\u006f\u0072\u0065\u0020\u0025\u0066\u0020\u003e\u003d\u0020\u0073\u0063\u006fr\u0065\u0054\u0068\u0072\u0065\u0073h\u006f\u006c\u0064 \u0025\u0066", _eea, _cece, _cee, scoreThreshold)
	}
	return false, nil
}

const (
	_ SizeComparison = iota
	SizeSelectIfLT
	SizeSelectIfGT
	SizeSelectIfLTE
	SizeSelectIfGTE
	SizeSelectIfEQ
)

type RasterOperator int

func Centroids(bms []*Bitmap) (*Points, error) {
	_ffacb := make([]Point, len(bms))
	_gedd := _ceafb()
	_ddee := _gdfg()
	var _addf error
	for _bcgfe, _efed := range bms {
		_ffacb[_bcgfe], _addf = _efed.centroid(_gedd, _ddee)
		if _addf != nil {
			return nil, _addf
		}
	}
	_ggfd := Points(_ffacb)
	return &_ggfd, nil
}

var (
	_eege = _eca()
	_bdce = _dccg()
	_dfec = _baf()
)

func _gbbb(_gggf, _bbec *Bitmap, _dcgg, _afgf, _eadg, _edfb, _fecg, _ggge, _acef, _adeg int, _dcd CombinationOperator) error {
	var _adfc int
	_gbff := func() { _adfc++; _eadg += _bbec.RowStride; _edfb += _gggf.RowStride; _fecg += _gggf.RowStride }
	for _adfc = _dcgg; _adfc < _afgf; _gbff() {
		var _fcfe uint16
		_gcgd := _eadg
		for _efec := _edfb; _efec <= _fecg; _efec++ {
			_cgdd, _dfff := _bbec.GetByte(_gcgd)
			if _dfff != nil {
				return _dfff
			}
			_bcggc, _dfff := _gggf.GetByte(_efec)
			if _dfff != nil {
				return _dfff
			}
			_fcfe = (_fcfe | uint16(_bcggc)) << uint(_adeg)
			_bcggc = byte(_fcfe >> 8)
			if _efec == _fecg {
				_bcggc = _cddg(uint(_ggge), _bcggc)
			}
			if _dfff = _bbec.SetByte(_gcgd, _cfag(_cgdd, _bcggc, _dcd)); _dfff != nil {
				return _dfff
			}
			_gcgd++
			_fcfe <<= uint(_acef)
		}
	}
	return nil
}
func _geaed(_adbf, _cagc *Bitmap, _dbag, _abde int) (*Bitmap, error) {
	const _bggd = "d\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063\u006b"
	if _cagc == nil {
		_gb.Log.Debug("\u0064\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063k\u0020\u0073\u006f\u0075\u0072\u0063\u0065 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
		return nil, _g.Error(_bggd, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _dbag < 1 || _abde < 1 {
		return nil, _g.Error(_bggd, "\u0068\u0053\u007a\u0069\u0065 \u0061\u006e\u0064\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u0061\u0072\u0065 \u006e\u006f\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _dbag == 1 && _abde == 1 {
		_abce, _gbeg := _feea(_adbf, _cagc)
		if _gbeg != nil {
			return nil, _g.Wrap(_gbeg, _bggd, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _abce, nil
	}
	if _dbag == 1 || _abde == 1 {
		_cfece := SelCreateBrick(_abde, _dbag, _abde/2, _dbag/2, SelHit)
		_gfeae, _eefe := _ceed(_adbf, _cagc, _cfece)
		if _eefe != nil {
			return nil, _g.Wrap(_eefe, _bggd, "\u0068s\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _gfeae, nil
	}
	_cdcec := SelCreateBrick(1, _dbag, 0, _dbag/2, SelHit)
	_fcaf := SelCreateBrick(_abde, 1, _abde/2, 0, SelHit)
	_aadf, _eagd := _ceed(nil, _cagc, _cdcec)
	if _eagd != nil {
		return nil, _g.Wrap(_eagd, _bggd, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_adbf, _eagd = _ceed(_adbf, _aadf, _fcaf)
	if _eagd != nil {
		return nil, _g.Wrap(_eagd, _bggd, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _adbf, nil
}
func _ga(_gc, _gedc *Bitmap) (_dbb error) {
	const _ce = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0038"
	_ff := _gedc.RowStride
	_bgb := _gc.RowStride
	var _bd, _aed, _gedca, _gae, _gedg int
	for _gedca = 0; _gedca < _gedc.Height; _gedca++ {
		_bd = _gedca * _ff
		_aed = 8 * _gedca * _bgb
		for _gae = 0; _gae < _ff; _gae++ {
			if _dbb = _gc.setEightBytes(_aed+_gae*8, _dfec[_gedc.Data[_bd+_gae]]); _dbb != nil {
				return _g.Wrap(_dbb, _ce, "")
			}
		}
		for _gedg = 1; _gedg < 8; _gedg++ {
			for _gae = 0; _gae < _bgb; _gae++ {
				if _dbb = _gc.SetByte(_aed+_gedg*_bgb+_gae, _gc.Data[_aed+_gae]); _dbb != nil {
					return _g.Wrap(_dbb, _ce, "")
				}
			}
		}
	}
	return nil
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

func (_dbeg *Bitmaps) SelectByIndexes(idx []int) (*Bitmaps, error) {
	const _bcdec = "B\u0069\u0074\u006d\u0061\u0070\u0073.\u0053\u006f\u0072\u0074\u0049\u006e\u0064\u0065\u0078e\u0073\u0042\u0079H\u0065i\u0067\u0068\u0074"
	_fdcfc, _egfe := _dbeg.selectByIndexes(idx)
	if _egfe != nil {
		return nil, _g.Wrap(_egfe, _bcdec, "")
	}
	return _fdcfc, nil
}
func (_cbef *Bitmap) setEightPartlyBytes(_dccc, _bcgd int, _cbg uint64) (_bdfd error) {
	var (
		_ebe  byte
		_bcef int
	)
	const _cafc = "\u0073\u0065\u0074\u0045ig\u0068\u0074\u0050\u0061\u0072\u0074\u006c\u0079\u0042\u0079\u0074\u0065\u0073"
	for _faad := 1; _faad <= _bcgd; _faad++ {
		_bcef = 64 - _faad*8
		_ebe = byte(_cbg >> uint(_bcef) & 0xff)
		_gb.Log.Trace("\u0074\u0065\u006d\u0070\u003a\u0020\u0025\u0030\u0038\u0062\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a %\u0064,\u0020\u0069\u0064\u0078\u003a\u0020\u0025\u0064\u002c\u0020\u0066\u0075l\u006c\u0042\u0079\u0074\u0065\u0073\u004e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c \u0073\u0068\u0069\u0066\u0074\u003a\u0020\u0025\u0064", _ebe, _dccc, _dccc+_faad-1, _bcgd, _bcef)
		if _bdfd = _cbef.SetByte(_dccc+_faad-1, _ebe); _bdfd != nil {
			return _g.Wrap(_bdfd, _cafc, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
		}
	}
	_eedc := _cbef.RowStride*8 - _cbef.Width
	if _eedc == 0 {
		return nil
	}
	_bcef -= 8
	_ebe = byte(_cbg>>uint(_bcef)&0xff) << uint(_eedc)
	if _bdfd = _cbef.SetByte(_dccc+_bcgd, _ebe); _bdfd != nil {
		return _g.Wrap(_bdfd, _cafc, "\u0070\u0061\u0064\u0064\u0065\u0064")
	}
	return nil
}
func (_bceaf *Bitmaps) GetBox(i int) (*_aa.Rectangle, error) {
	const _eabb = "\u0047\u0065\u0074\u0042\u006f\u0078"
	if _bceaf == nil {
		return nil, _g.Error(_eabb, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074\u006d\u0061\u0070s\u0027")
	}
	if i > len(_bceaf.Boxes)-1 {
		return nil, _g.Errorf(_eabb, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _bceaf.Boxes[i], nil
}
func (_adaed *byWidth) Len() int { return len(_adaed.Values) }
func _gccg(_bgge *Bitmap, _ggedc *_dd.Stack, _ecfcd, _febe, _dccgbf int) (_cdggc *_aa.Rectangle, _cdgbb error) {
	const _ffae = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _bgge == nil {
		return nil, _g.Error(_ffae, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _ggedc == nil {
		return nil, _g.Error(_ffae, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	switch _dccgbf {
	case 4:
		if _cdggc, _cdgbb = _acdc(_bgge, _ggedc, _ecfcd, _febe); _cdgbb != nil {
			return nil, _g.Wrap(_cdgbb, _ffae, "")
		}
		return _cdggc, nil
	case 8:
		if _cdggc, _cdgbb = _ccgf(_bgge, _ggedc, _ecfcd, _febe); _cdgbb != nil {
			return nil, _g.Wrap(_cdgbb, _ffae, "")
		}
		return _cdggc, nil
	default:
		return nil, _g.Errorf(_ffae, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u0069\u0073 \u006eo\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038\u003a\u0020\u0027\u0025\u0064\u0027", _dccgbf)
	}
}
func _aaba(_fcggd, _cbfg *Bitmap, _aegb, _bbdc int) (_agfb error) {
	const _bceg = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0034"
	var (
		_ecaa, _abgg, _fdde, _ceabd                      int
		_fdcf, _gega, _cbac, _bfee, _bgaf, _ffgf, _acgab byte
	)
	for _ecaa = 0; _ecaa < _aegb; _ecaa++ {
		_fdde = _ecaa * _fcggd.RowStride
		_ceabd = _ecaa * _cbfg.RowStride
		for _abgg = 0; _abgg < _bbdc; _abgg++ {
			_fdcf, _agfb = _fcggd.GetByte(_fdde + _abgg)
			if _agfb != nil {
				return _g.Wrap(_agfb, _bceg, "\u0066i\u0072\u0073\u0074\u0020\u0067\u0065t")
			}
			_gega, _agfb = _cbfg.GetByte(_ceabd + _abgg)
			if _agfb != nil {
				return _g.Wrap(_agfb, _bceg, "\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0067\u0065\u0074")
			}
			if _ecaa > 0 {
				_cbac, _agfb = _fcggd.GetByte(_fdde - _fcggd.RowStride + _abgg)
				if _agfb != nil {
					return _g.Wrap(_agfb, _bceg, "\u0069\u0020\u003e \u0030")
				}
				_fdcf |= _cbac
			}
			if _abgg > 0 {
				_bfee, _agfb = _fcggd.GetByte(_fdde + _abgg - 1)
				if _agfb != nil {
					return _g.Wrap(_agfb, _bceg, "\u006a\u0020\u003e \u0030")
				}
				_fdcf |= _bfee << 7
			}
			_fdcf &= _gega
			if _fdcf == 0 || (^_fdcf) == 0 {
				if _agfb = _fcggd.SetByte(_fdde+_abgg, _fdcf); _agfb != nil {
					return _g.Wrap(_agfb, _bceg, "b\u0074\u0020\u003d\u003d 0\u0020|\u007c\u0020\u0028\u005e\u0062t\u0029\u0020\u003d\u003d\u0020\u0030")
				}
				continue
			}
			for {
				_acgab = _fdcf
				_fdcf = (_fdcf | (_fdcf >> 1) | (_fdcf << 1)) & _gega
				if (_fdcf ^ _acgab) == 0 {
					if _agfb = _fcggd.SetByte(_fdde+_abgg, _fdcf); _agfb != nil {
						return _g.Wrap(_agfb, _bceg, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _ecaa = _aegb - 1; _ecaa >= 0; _ecaa-- {
		_fdde = _ecaa * _fcggd.RowStride
		_ceabd = _ecaa * _cbfg.RowStride
		for _abgg = _bbdc - 1; _abgg >= 0; _abgg-- {
			if _fdcf, _agfb = _fcggd.GetByte(_fdde + _abgg); _agfb != nil {
				return _g.Wrap(_agfb, _bceg, "\u0072\u0065\u0076\u0065\u0072\u0073\u0065\u0020\u0066\u0069\u0072\u0073t\u0020\u0067\u0065\u0074")
			}
			if _gega, _agfb = _cbfg.GetByte(_ceabd + _abgg); _agfb != nil {
				return _g.Wrap(_agfb, _bceg, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _ecaa < _aegb-1 {
				if _bgaf, _agfb = _fcggd.GetByte(_fdde + _fcggd.RowStride + _abgg); _agfb != nil {
					return _g.Wrap(_agfb, _bceg, "\u0072\u0065v\u0065\u0072\u0073e\u0020\u0069\u0020\u003c\u0020\u0068\u0020\u002d\u0031")
				}
				_fdcf |= _bgaf
			}
			if _abgg < _bbdc-1 {
				if _ffgf, _agfb = _fcggd.GetByte(_fdde + _abgg + 1); _agfb != nil {
					return _g.Wrap(_agfb, _bceg, "\u0072\u0065\u0076\u0065rs\u0065\u0020\u006a\u0020\u003c\u0020\u0077\u0070\u006c\u0020\u002d\u0020\u0031")
				}
				_fdcf |= _ffgf >> 7
			}
			_fdcf &= _gega
			if _fdcf == 0 || (^_fdcf) == 0 {
				if _agfb = _fcggd.SetByte(_fdde+_abgg, _fdcf); _agfb != nil {
					return _g.Wrap(_agfb, _bceg, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006d\u0061\u0073k\u0065\u0064\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061i\u006c\u0065\u0064")
				}
				continue
			}
			for {
				_acgab = _fdcf
				_fdcf = (_fdcf | (_fdcf >> 1) | (_fdcf << 1)) & _gega
				if (_fdcf ^ _acgab) == 0 {
					if _agfb = _fcggd.SetByte(_fdde+_abgg, _fdcf); _agfb != nil {
						return _g.Wrap(_agfb, _bceg, "\u0072e\u0076\u0065\u0072\u0073e\u0020\u0073\u0065\u0074\u0074i\u006eg\u0020p\u0072\u0065\u0076\u0020\u0062\u0079\u0074e")
					}
					break
				}
			}
		}
	}
	return nil
}
func (_bccb *BitmapsArray) AddBox(box *_aa.Rectangle) { _bccb.Boxes = append(_bccb.Boxes, box) }
func (_fdaf *Bitmap) And(s *Bitmap) (_ggdg *Bitmap, _gbd error) {
	const _egca = "\u0042\u0069\u0074\u006d\u0061\u0070\u002e\u0041\u006e\u0064"
	if _fdaf == nil {
		return nil, _g.Error(_egca, "\u0027b\u0069t\u006d\u0061\u0070\u0020\u0027b\u0027\u0020i\u0073\u0020\u006e\u0069\u006c")
	}
	if s == nil {
		return nil, _g.Error(_egca, "\u0062\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if !_fdaf.SizesEqual(s) {
		_gb.Log.Debug("\u0025\u0073\u0020-\u0020\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0069\u007a\u0065 \u0077\u0069\u0074\u0068\u0020\u0027\u0062\u0027", _egca)
	}
	if _ggdg, _gbd = _feea(_ggdg, _fdaf); _gbd != nil {
		return nil, _g.Wrap(_gbd, _egca, "\u0063\u0061\u006e't\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _gbd = _ggdg.RasterOperation(0, 0, _ggdg.Width, _ggdg.Height, PixSrcAndDst, s, 0, 0); _gbd != nil {
		return nil, _g.Wrap(_gbd, _egca, "")
	}
	return _ggdg, nil
}
func _bgga(_ebaa, _fgab *Bitmap, _eadbb, _bbef int) (*Bitmap, error) {
	const _bdca = "\u0063\u006c\u006f\u0073\u0065\u0053\u0061\u0066\u0065B\u0072\u0069\u0063\u006b"
	if _fgab == nil {
		return nil, _g.Error(_bdca, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _eadbb < 1 || _bbef < 1 {
		return nil, _g.Error(_bdca, "\u0068s\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _eadbb == 1 && _bbef == 1 {
		return _feea(_ebaa, _fgab)
	}
	if MorphBC == SymmetricMorphBC {
		_gcab, _eebf := _gcbe(_ebaa, _fgab, _eadbb, _bbef)
		if _eebf != nil {
			return nil, _g.Wrap(_eebf, _bdca, "\u0053\u0079m\u006d\u0065\u0074r\u0069\u0063\u004d\u006f\u0072\u0070\u0068\u0042\u0043")
		}
		return _gcab, nil
	}
	_aceca := _aage(_eadbb/2, _bbef/2)
	_eec := 8 * ((_aceca + 7) / 8)
	_dcbb, _aecd := _fgab.AddBorder(_eec, 0)
	if _aecd != nil {
		return nil, _g.Wrapf(_aecd, _bdca, "\u0042\u006f\u0072\u0064\u0065\u0072\u0053\u0069\u007ae\u003a\u0020\u0025\u0064", _eec)
	}
	var _fbfb, _faadg *Bitmap
	if _eadbb == 1 || _bbef == 1 {
		_efcf := SelCreateBrick(_bbef, _eadbb, _bbef/2, _eadbb/2, SelHit)
		_fbfb, _aecd = _fgfg(nil, _dcbb, _efcf)
		if _aecd != nil {
			return nil, _g.Wrap(_aecd, _bdca, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
	} else {
		_cbccg := SelCreateBrick(1, _eadbb, 0, _eadbb/2, SelHit)
		_bgae, _efbf := _ceed(nil, _dcbb, _cbccg)
		if _efbf != nil {
			return nil, _g.Wrap(_efbf, _bdca, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0064\u0069\u006c\u0061t\u0065")
		}
		_ddd := SelCreateBrick(_bbef, 1, _bbef/2, 0, SelHit)
		_fbfb, _efbf = _ceed(nil, _bgae, _ddd)
		if _efbf != nil {
			return nil, _g.Wrap(_efbf, _bdca, "\u0072\u0065\u0067ul\u0061\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
		}
		if _, _efbf = _babe(_bgae, _fbfb, _cbccg); _efbf != nil {
			return nil, _g.Wrap(_efbf, _bdca, "r\u0065\u0067\u0075\u006car\u0020-\u0020\u0066\u0069\u0072\u0073t\u0020\u0065\u0072\u006f\u0064\u0065")
		}
		if _, _efbf = _babe(_fbfb, _bgae, _ddd); _efbf != nil {
			return nil, _g.Wrap(_efbf, _bdca, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0065\u0072\u006fd\u0065")
		}
	}
	if _faadg, _aecd = _fbfb.RemoveBorder(_eec); _aecd != nil {
		return nil, _g.Wrap(_aecd, _bdca, "\u0072e\u0067\u0075\u006c\u0061\u0072")
	}
	if _ebaa == nil {
		return _faadg, nil
	}
	if _, _aecd = _feea(_ebaa, _faadg); _aecd != nil {
		return nil, _aecd
	}
	return _ebaa, nil
}
func (_cag *Bitmap) nextOnPixelLow(_ecdb, _accd, _daebf, _bgef, _dbg int) (_fbgaf _aa.Point, _bedf bool, _deb error) {
	const _gaad = "B\u0069\u0074\u006d\u0061p.\u006ee\u0078\u0074\u004f\u006e\u0050i\u0078\u0065\u006c\u004c\u006f\u0077"
	var (
		_fefc int
		_gfc  byte
	)
	_fedf := _dbg * _daebf
	_fbgag := _fedf + (_bgef / 8)
	if _gfc, _deb = _cag.GetByte(_fbgag); _deb != nil {
		return _fbgaf, false, _g.Wrap(_deb, _gaad, "\u0078\u0053\u0074\u0061\u0072\u0074\u0020\u0061\u006e\u0064 \u0079\u0053\u0074\u0061\u0072\u0074\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	if _gfc != 0 {
		_ecaf := _bgef - (_bgef % 8) + 7
		for _fefc = _bgef; _fefc <= _ecaf && _fefc < _ecdb; _fefc++ {
			if _cag.GetPixel(_fefc, _dbg) {
				_fbgaf.X = _fefc
				_fbgaf.Y = _dbg
				return _fbgaf, true, nil
			}
		}
	}
	_dec := (_bgef / 8) + 1
	_fefc = 8 * _dec
	var _bfde int
	for _fbgag = _fedf + _dec; _fefc < _ecdb; _fbgag, _fefc = _fbgag+1, _fefc+8 {
		if _gfc, _deb = _cag.GetByte(_fbgag); _deb != nil {
			return _fbgaf, false, _g.Wrap(_deb, _gaad, "r\u0065\u0073\u0074\u0020of\u0020t\u0068\u0065\u0020\u006c\u0069n\u0065\u0020\u0062\u0079\u0074\u0065")
		}
		if _gfc == 0 {
			continue
		}
		for _bfde = 0; _bfde < 8 && _fefc < _ecdb; _bfde, _fefc = _bfde+1, _fefc+1 {
			if _cag.GetPixel(_fefc, _dbg) {
				_fbgaf.X = _fefc
				_fbgaf.Y = _dbg
				return _fbgaf, true, nil
			}
		}
	}
	for _ggga := _dbg + 1; _ggga < _accd; _ggga++ {
		_fedf = _ggga * _daebf
		for _fbgag, _fefc = _fedf, 0; _fefc < _ecdb; _fbgag, _fefc = _fbgag+1, _fefc+8 {
			if _gfc, _deb = _cag.GetByte(_fbgag); _deb != nil {
				return _fbgaf, false, _g.Wrap(_deb, _gaad, "\u0066o\u006cl\u006f\u0077\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u0073")
			}
			if _gfc == 0 {
				continue
			}
			for _bfde = 0; _bfde < 8 && _fefc < _ecdb; _bfde, _fefc = _bfde+1, _fefc+1 {
				if _cag.GetPixel(_fefc, _ggga) {
					_fbgaf.X = _fefc
					_fbgaf.Y = _ggga
					return _fbgaf, true, nil
				}
			}
		}
	}
	return _fbgaf, false, nil
}
func (_cdaa *Boxes) Add(box *_aa.Rectangle) error {
	if _cdaa == nil {
		return _g.Error("\u0042o\u0078\u0065\u0073\u002e\u0041\u0064d", "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	*_cdaa = append(*_cdaa, box)
	return nil
}
func (_cddf *Bitmap) RemoveBorderGeneral(left, right, top, bot int) (*Bitmap, error) {
	return _cddf.removeBorderGeneral(left, right, top, bot)
}

type SelectionValue int

func (_cbcff *Bitmap) centroid(_agaeg, _ebbf []int) (Point, error) {
	_bfgc := Point{}
	_cbcff.setPadBits(0)
	if len(_agaeg) == 0 {
		_agaeg = _ceafb()
	}
	if len(_ebbf) == 0 {
		_ebbf = _gdfg()
	}
	var _abga, _afgd, _bac, _ccdg, _fbff, _gddf int
	var _fgca byte
	for _fbff = 0; _fbff < _cbcff.Height; _fbff++ {
		_ggcfg := _cbcff.RowStride * _fbff
		_ccdg = 0
		for _gddf = 0; _gddf < _cbcff.RowStride; _gddf++ {
			_fgca = _cbcff.Data[_ggcfg+_gddf]
			if _fgca != 0 {
				_ccdg += _ebbf[_fgca]
				_abga += _agaeg[_fgca] + _gddf*8*_ebbf[_fgca]
			}
		}
		_bac += _ccdg
		_afgd += _ccdg * _fbff
	}
	if _bac != 0 {
		_bfgc.X = float32(_abga) / float32(_bac)
		_bfgc.Y = float32(_afgd) / float32(_bac)
	}
	return _bfgc, nil
}
func TstWordBitmap(t *_ba.T, scale ...int) *Bitmap {
	_bgaae := 1
	if len(scale) > 0 {
		_bgaae = scale[0]
	}
	_bafb := 3
	_gbbd := 9 + 7 + 15 + 2*_bafb
	_fbce := 5 + _bafb + 5
	_ggaea := New(_gbbd*_bgaae, _fbce*_bgaae)
	_dgbf := &Bitmaps{}
	var _fbcaa *int
	_bafb *= _bgaae
	_ceede := 0
	_fbcaa = &_ceede
	_bcfc := 0
	_efcb := TstDSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstOSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, _bafb)
	_efcb = TstISymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstTSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, _bafb)
	_efcb = TstNSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstOSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstWSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 0)
	*_fbcaa = 0
	_bcfc = 5*_bgaae + _bafb
	_efcb = TstOSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstRSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, _bafb)
	_efcb = TstNSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstESymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstVSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstESymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 1*_bgaae)
	_efcb = TstRSymbol(t, scale...)
	TstAddSymbol(t, _dgbf, _efcb, _fbcaa, _bcfc, 0)
	TstWriteSymbols(t, _dgbf, _ggaea)
	return _ggaea
}
func (_dba *Bitmap) CountPixels() int { return _dba.countPixels() }
func (_gged *Bitmap) GetPixel(x, y int) bool {
	_ede := _gged.GetByteIndex(x, y)
	_fed := _gged.GetBitOffset(x)
	_gba := uint(7 - _fed)
	if _ede > len(_gged.Data)-1 {
		_gb.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0072\u0061\u006e\u0067\u0065\u002e \u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0027\u0025\u0064'\u002c\u0020\u0062m\u003a\u0020\u0027\u0025\u0073\u0027", x, y, _gged)
		return false
	}
	if (_gged.Data[_ede]>>_gba)&0x01 >= 1 {
		return true
	}
	return false
}
func (_afbe *Bitmap) SetPixel(x, y int, pixel byte) error {
	_cdgg := _afbe.GetByteIndex(x, y)
	if _cdgg > len(_afbe.Data)-1 {
		return _g.Errorf("\u0053\u0065\u0074\u0050\u0069\u0078\u0065\u006c", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", _cdgg)
	}
	_gffc := _afbe.GetBitOffset(x)
	_gbb := uint(7 - _gffc)
	_ffde := _afbe.Data[_cdgg]
	var _acc byte
	if pixel == 1 {
		_acc = _ffde | (pixel & 0x01 << _gbb)
	} else {
		_acc = _ffde &^ (1 << _gbb)
	}
	_afbe.Data[_cdgg] = _acc
	return nil
}
func (_eeeb *Selection) findMaxTranslations() (_ddbba, _afcd, _fbfe, _fffc int) {
	for _bcbaa := 0; _bcbaa < _eeeb.Height; _bcbaa++ {
		for _edag := 0; _edag < _eeeb.Width; _edag++ {
			if _eeeb.Data[_bcbaa][_edag] == SelHit {
				_ddbba = _aage(_ddbba, _eeeb.Cx-_edag)
				_afcd = _aage(_afcd, _eeeb.Cy-_bcbaa)
				_fbfe = _aage(_fbfe, _edag-_eeeb.Cx)
				_fffc = _aage(_fffc, _bcbaa-_eeeb.Cy)
			}
		}
	}
	return _ddbba, _afcd, _fbfe, _fffc
}

var (
	_gbef  *Bitmap
	_fbbbf *Bitmap
)

func (_fcf *Bitmap) CreateTemplate() *Bitmap { return _fcf.createTemplate() }

var _ada = [5]int{1, 2, 3, 0, 4}

type Component int

func (_fdcbg *Bitmaps) HeightSorter() func(_cgaa, _cgad int) bool {
	return func(_ebdb, _egcfd int) bool {
		_bbfb := _fdcbg.Values[_ebdb].Height < _fdcbg.Values[_egcfd].Height
		_gb.Log.Debug("H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0076\u0020<\u0020\u0025\u0076\u0020= \u0025\u0076", _fdcbg.Values[_ebdb].Height, _fdcbg.Values[_egcfd].Height, _bbfb)
		return _bbfb
	}
}
func _acad(_aga, _cfeg *Bitmap, _edec CombinationOperator) *Bitmap {
	_egee := New(_aga.Width, _aga.Height)
	for _caba := 0; _caba < len(_egee.Data); _caba++ {
		_egee.Data[_caba] = _cfag(_aga.Data[_caba], _cfeg.Data[_caba], _edec)
	}
	return _egee
}
func TstISymbol(t *_ba.T, scale ...int) *Bitmap {
	_afbf, _efdga := NewWithData(1, 5, []byte{0x80, 0x80, 0x80, 0x80, 0x80})
	_b.NoError(t, _efdga)
	return TstGetScaledSymbol(t, _afbf, scale...)
}
func _babe(_cdgb, _faef *Bitmap, _fbag *Selection) (*Bitmap, error) {
	const _cbgg = "\u0065\u0072\u006fd\u0065"
	var (
		_feef error
		_fdac *Bitmap
	)
	_cdgb, _feef = _faaf(_cdgb, _faef, _fbag, &_fdac)
	if _feef != nil {
		return nil, _g.Wrap(_feef, _cbgg, "")
	}
	if _feef = _cdgb.setAll(); _feef != nil {
		return nil, _g.Wrap(_feef, _cbgg, "")
	}
	var _fceec SelectionValue
	for _ggegf := 0; _ggegf < _fbag.Height; _ggegf++ {
		for _egf := 0; _egf < _fbag.Width; _egf++ {
			_fceec = _fbag.Data[_ggegf][_egf]
			if _fceec == SelHit {
				_feef = _aee(_cdgb, _fbag.Cx-_egf, _fbag.Cy-_ggegf, _faef.Width, _faef.Height, PixSrcAndDst, _fdac, 0, 0)
				if _feef != nil {
					return nil, _g.Wrap(_feef, _cbgg, "")
				}
			}
		}
	}
	if MorphBC == SymmetricMorphBC {
		return _cdgb, nil
	}
	_gdgc, _bccg, _gcec, _fbgaff := _fbag.findMaxTranslations()
	if _gdgc > 0 {
		if _feef = _cdgb.RasterOperation(0, 0, _gdgc, _faef.Height, PixClr, nil, 0, 0); _feef != nil {
			return nil, _g.Wrap(_feef, _cbgg, "\u0078\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _gcec > 0 {
		if _feef = _cdgb.RasterOperation(_faef.Width-_gcec, 0, _gcec, _faef.Height, PixClr, nil, 0, 0); _feef != nil {
			return nil, _g.Wrap(_feef, _cbgg, "\u0078\u006e\u0020\u003e\u0020\u0030")
		}
	}
	if _bccg > 0 {
		if _feef = _cdgb.RasterOperation(0, 0, _faef.Width, _bccg, PixClr, nil, 0, 0); _feef != nil {
			return nil, _g.Wrap(_feef, _cbgg, "\u0079\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _fbgaff > 0 {
		if _feef = _cdgb.RasterOperation(0, _faef.Height-_fbgaff, _faef.Width, _fbgaff, PixClr, nil, 0, 0); _feef != nil {
			return nil, _g.Wrap(_feef, _cbgg, "\u0079\u006e\u0020\u003e\u0020\u0030")
		}
	}
	return _cdgb, nil
}
func (_ggg *Bitmap) Copy() *Bitmap {
	_gfbe := make([]byte, len(_ggg.Data))
	copy(_gfbe, _ggg.Data)
	return &Bitmap{Width: _ggg.Width, Height: _ggg.Height, RowStride: _ggg.RowStride, Data: _gfbe, Color: _ggg.Color, Text: _ggg.Text, BitmapNumber: _ggg.BitmapNumber, Special: _ggg.Special}
}
func (_eae *Bitmap) setPadBits(_fad int) {
	_degg := 8 - _eae.Width%8
	if _degg == 8 {
		return
	}
	_baef := _eae.Width / 8
	_aebe := _aeeg[_degg]
	if _fad == 0 {
		_aebe ^= _aebe
	}
	var _fgda int
	for _gggd := 0; _gggd < _eae.Height; _gggd++ {
		_fgda = _gggd*_eae.RowStride + _baef
		if _fad == 0 {
			_eae.Data[_fgda] &= _aebe
		} else {
			_eae.Data[_fgda] |= _aebe
		}
	}
}
func _efag(_ffb, _adbg int) int {
	if _ffb < _adbg {
		return _ffb
	}
	return _adbg
}
func _fegb(_aebg *Bitmap, _bebd, _babc, _bbdb, _gggde int, _cdff RasterOperator, _fbbgb *Bitmap, _fceda, _abad int) error {
	var (
		_gcfd        byte
		_eggb        int
		_dceb        int
		_gcad, _egcf int
		_bcga, _aaca int
	)
	_ddfg := _bbdb >> 3
	_bcbeb := _bbdb & 7
	if _bcbeb > 0 {
		_gcfd = _gbdb[_bcbeb]
	}
	_eggb = _fbbgb.RowStride*_abad + (_fceda >> 3)
	_dceb = _aebg.RowStride*_babc + (_bebd >> 3)
	switch _cdff {
	case PixSrc:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] = _fbbgb.Data[_gcad]
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], _fbbgb.Data[_gcad], _gcfd)
			}
		}
	case PixNotSrc:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] = ^(_fbbgb.Data[_gcad])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], ^_fbbgb.Data[_gcad], _gcfd)
			}
		}
	case PixSrcOrDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] |= _fbbgb.Data[_gcad]
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], _fbbgb.Data[_gcad]|_aebg.Data[_egcf], _gcfd)
			}
		}
	case PixSrcAndDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] &= _fbbgb.Data[_gcad]
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], _fbbgb.Data[_gcad]&_aebg.Data[_egcf], _gcfd)
			}
		}
	case PixSrcXorDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] ^= _fbbgb.Data[_gcad]
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], _fbbgb.Data[_gcad]^_aebg.Data[_egcf], _gcfd)
			}
		}
	case PixNotSrcOrDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] |= ^(_fbbgb.Data[_gcad])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], ^(_fbbgb.Data[_gcad])|_aebg.Data[_egcf], _gcfd)
			}
		}
	case PixNotSrcAndDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] &= ^(_fbbgb.Data[_gcad])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], ^(_fbbgb.Data[_gcad])&_aebg.Data[_egcf], _gcfd)
			}
		}
	case PixSrcOrNotDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] = _fbbgb.Data[_gcad] | ^(_aebg.Data[_egcf])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], _fbbgb.Data[_gcad]|^(_aebg.Data[_egcf]), _gcfd)
			}
		}
	case PixSrcAndNotDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] = _fbbgb.Data[_gcad] &^ (_aebg.Data[_egcf])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], _fbbgb.Data[_gcad]&^(_aebg.Data[_egcf]), _gcfd)
			}
		}
	case PixNotPixSrcOrDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] = ^(_fbbgb.Data[_gcad] | _aebg.Data[_egcf])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], ^(_fbbgb.Data[_gcad] | _aebg.Data[_egcf]), _gcfd)
			}
		}
	case PixNotPixSrcAndDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] = ^(_fbbgb.Data[_gcad] & _aebg.Data[_egcf])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], ^(_fbbgb.Data[_gcad] & _aebg.Data[_egcf]), _gcfd)
			}
		}
	case PixNotPixSrcXorDst:
		for _bcga = 0; _bcga < _gggde; _bcga++ {
			_gcad = _eggb + _bcga*_fbbgb.RowStride
			_egcf = _dceb + _bcga*_aebg.RowStride
			for _aaca = 0; _aaca < _ddfg; _aaca++ {
				_aebg.Data[_egcf] = ^(_fbbgb.Data[_gcad] ^ _aebg.Data[_egcf])
				_egcf++
				_gcad++
			}
			if _bcbeb > 0 {
				_aebg.Data[_egcf] = _aegg(_aebg.Data[_egcf], ^(_fbbgb.Data[_gcad] ^ _aebg.Data[_egcf]), _gcfd)
			}
		}
	default:
		_gb.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _cdff)
		return _g.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004co\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_cdab *Boxes) selectWithIndicator(_geed *_dd.NumSlice) (_gbe *Boxes, _ffbd error) {
	const _dgda = "\u0042o\u0078\u0065\u0073\u002es\u0065\u006c\u0065\u0063\u0074W\u0069t\u0068I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _cdab == nil {
		return nil, _g.Error(_dgda, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _geed == nil {
		return nil, _g.Error(_dgda, "\u0027\u006ea\u0027\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_geed) != len(*_cdab) {
		return nil, _g.Error(_dgda, "\u0062\u006f\u0078\u0065\u0073\u0020\u0027\u0062\u0027\u0020\u0068\u0061\u0073\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020s\u0069\u007a\u0065\u0020\u0074h\u0061\u006e \u0027\u006e\u0061\u0027")
	}
	var _cfgf, _fecff int
	for _eadb := 0; _eadb < len(*_geed); _eadb++ {
		if _cfgf, _ffbd = _geed.GetInt(_eadb); _ffbd != nil {
			return nil, _g.Wrap(_ffbd, _dgda, "\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u0067\u0020c\u006f\u0075\u006e\u0074")
		}
		if _cfgf == 1 {
			_fecff++
		}
	}
	if _fecff == len(*_cdab) {
		return _cdab, nil
	}
	_eaea := Boxes{}
	for _dgge := 0; _dgge < len(*_geed); _dgge++ {
		_cfgf = int((*_geed)[_dgge])
		if _cfgf == 0 {
			continue
		}
		_eaea = append(_eaea, (*_cdab)[_dgge])
	}
	_gbe = &_eaea
	return _gbe, nil
}
func Rect(x, y, w, h int) (*_aa.Rectangle, error) {
	const _gcbf = "b\u0069\u0074\u006d\u0061\u0070\u002e\u0052\u0065\u0063\u0074"
	if x < 0 {
		w += x
		x = 0
		if w <= 0 {
			return nil, _g.Errorf(_gcbf, "x\u003a\u0027\u0025\u0064\u0027\u0020<\u0020\u0030\u0020\u0061\u006e\u0064\u0020\u0077\u003a \u0027\u0025\u0064'\u0020<\u003d\u0020\u0030", x, w)
		}
	}
	if y < 0 {
		h += y
		y = 0
		if h <= 0 {
			return nil, _g.Error(_gcbf, "\u0079\u0020\u003c 0\u0020\u0061\u006e\u0064\u0020\u0062\u006f\u0078\u0020\u006f\u0066\u0066\u0020\u002b\u0071\u0075\u0061\u0064")
		}
	}
	_gbfa := _aa.Rect(x, y, x+w, y+h)
	return &_gbfa, nil
}
func TstGetScaledSymbol(t *_ba.T, sm *Bitmap, scale ...int) *Bitmap {
	if len(scale) == 0 {
		return sm
	}
	if scale[0] == 1 {
		return sm
	}
	_daae, _dfbad := MorphSequence(sm, MorphProcess{Operation: MopReplicativeBinaryExpansion, Arguments: scale})
	_b.NoError(t, _dfbad)
	return _daae
}
func _gdfg() []int {
	_fcfa := make([]int, 256)
	for _bfbe := 0; _bfbe <= 0xff; _bfbe++ {
		_fecd := byte(_bfbe)
		_fcfa[_fecd] = int(_fecd&0x1) + (int(_fecd>>1) & 0x1) + (int(_fecd>>2) & 0x1) + (int(_fecd>>3) & 0x1) + (int(_fecd>>4) & 0x1) + (int(_fecd>>5) & 0x1) + (int(_fecd>>6) & 0x1) + (int(_fecd>>7) & 0x1)
	}
	return _fcfa
}
func (_egda *Bitmap) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _aee(_egda, dx, dy, dw, dh, op, src, sx, sy)
}
func (_bcgb Points) XSorter() func(_efae, _fabf int) bool {
	return func(_fgdgad, _begbf int) bool { return _bcgb[_fgdgad].X < _bcgb[_begbf].X }
}
func (_dda *Bitmap) createTemplate() *Bitmap {
	return &Bitmap{Width: _dda.Width, Height: _dda.Height, RowStride: _dda.RowStride, Color: _dda.Color, Text: _dda.Text, BitmapNumber: _dda.BitmapNumber, Special: _dda.Special, Data: make([]byte, len(_dda.Data))}
}
func _gdb(_dab, _dcg *Bitmap, _afe int, _fab []byte, _bb int) (_dbbd error) {
	const _bcg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0031"
	var (
		_bcea, _dgc, _dag, _bdf, _bgcb, _aeb, _ebff, _ced int
		_afa, _cdd                                        uint32
		_fbg, _acb                                        byte
		_ed                                               uint16
	)
	_fbc := make([]byte, 4)
	_ebfff := make([]byte, 4)
	for _dag = 0; _dag < _dab.Height-1; _dag, _bdf = _dag+2, _bdf+1 {
		_bcea = _dag * _dab.RowStride
		_dgc = _bdf * _dcg.RowStride
		for _bgcb, _aeb = 0, 0; _bgcb < _bb; _bgcb, _aeb = _bgcb+4, _aeb+1 {
			for _ebff = 0; _ebff < 4; _ebff++ {
				_ced = _bcea + _bgcb + _ebff
				if _ced <= len(_dab.Data)-1 && _ced < _bcea+_dab.RowStride {
					_fbc[_ebff] = _dab.Data[_ced]
				} else {
					_fbc[_ebff] = 0x00
				}
				_ced = _bcea + _dab.RowStride + _bgcb + _ebff
				if _ced <= len(_dab.Data)-1 && _ced < _bcea+(2*_dab.RowStride) {
					_ebfff[_ebff] = _dab.Data[_ced]
				} else {
					_ebfff[_ebff] = 0x00
				}
			}
			_afa = _cc.BigEndian.Uint32(_fbc)
			_cdd = _cc.BigEndian.Uint32(_ebfff)
			_cdd |= _afa
			_cdd |= _cdd << 1
			_cdd &= 0xaaaaaaaa
			_afa = _cdd | (_cdd << 7)
			_fbg = byte(_afa >> 24)
			_acb = byte((_afa >> 8) & 0xff)
			_ced = _dgc + _aeb
			if _ced+1 == len(_dcg.Data)-1 || _ced+1 >= _dgc+_dcg.RowStride {
				_dcg.Data[_ced] = _fab[_fbg]
			} else {
				_ed = (uint16(_fab[_fbg]) << 8) | uint16(_fab[_acb])
				if _dbbd = _dcg.setTwoBytes(_ced, _ed); _dbbd != nil {
					return _g.Wrapf(_dbbd, _bcg, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _ced)
				}
				_aeb++
			}
		}
	}
	return nil
}
func HausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH int) (bool, error) {
	const _ffbe = "\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_efeca, _eedd := p1.Width, p1.Height
	_cbcg, _agab := p3.Width, p3.Height
	if _dd.Abs(_efeca-_cbcg) > maxDiffW {
		return false, nil
	}
	if _dd.Abs(_eedd-_agab) > maxDiffH {
		return false, nil
	}
	_bgce := int(delX + _dd.Sign(delX)*0.5)
	_dfgg := int(delY + _dd.Sign(delY)*0.5)
	var _baaa error
	_aabd := p1.CreateTemplate()
	if _baaa = _aabd.RasterOperation(0, 0, _efeca, _eedd, PixSrc, p1, 0, 0); _baaa != nil {
		return false, _g.Wrap(_baaa, _ffbe, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _baaa = _aabd.RasterOperation(_bgce, _dfgg, _efeca, _eedd, PixNotSrcAndDst, p4, 0, 0); _baaa != nil {
		return false, _g.Wrap(_baaa, _ffbe, "\u0021p\u0034\u0020\u0026\u0020\u0074")
	}
	if _aabd.Zero() {
		return false, nil
	}
	if _baaa = _aabd.RasterOperation(_bgce, _dfgg, _cbcg, _agab, PixSrc, p3, 0, 0); _baaa != nil {
		return false, _g.Wrap(_baaa, _ffbe, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _baaa = _aabd.RasterOperation(0, 0, _cbcg, _agab, PixNotSrcAndDst, p2, 0, 0); _baaa != nil {
		return false, _g.Wrap(_baaa, _ffbe, "\u0021p\u0032\u0020\u0026\u0020\u0074")
	}
	return _aabd.Zero(), nil
}
func (_efdb CombinationOperator) String() string {
	var _afab string
	switch _efdb {
	case CmbOpOr:
		_afab = "\u004f\u0052"
	case CmbOpAnd:
		_afab = "\u0041\u004e\u0044"
	case CmbOpXor:
		_afab = "\u0058\u004f\u0052"
	case CmbOpXNor:
		_afab = "\u0058\u004e\u004f\u0052"
	case CmbOpReplace:
		_afab = "\u0052E\u0050\u004c\u0041\u0043\u0045"
	case CmbOpNot:
		_afab = "\u004e\u004f\u0054"
	}
	return _afab
}

type CombinationOperator int

func _ccbfd(_ddfgc *Bitmap, _ecba, _eeaa, _fdgg, _ggcg int, _ebeaa RasterOperator, _ddbeg *Bitmap, _cgdc, _ffbf int) error {
	var (
		_defa        bool
		_fbaga       bool
		_decdg       byte
		_bbfd        int
		_dfed        int
		_fbcf        int
		_fbadd       int
		_aegeg       bool
		_dbdae       int
		_aadb        int
		_cacf        int
		_cega        bool
		_aagb        byte
		_dbga        int
		_fdcg        int
		_fcbb        int
		_cae         byte
		_bdeb        int
		_feag        int
		_cggfc       uint
		_cdfa        uint
		_cdcgf       byte
		_fbbe        shift
		_fafbe       bool
		_abba        bool
		_cgee, _eafc int
	)
	if _cgdc&7 != 0 {
		_feag = 8 - (_cgdc & 7)
	}
	if _ecba&7 != 0 {
		_dfed = 8 - (_ecba & 7)
	}
	if _feag == 0 && _dfed == 0 {
		_cdcgf = _aeeg[0]
	} else {
		if _dfed > _feag {
			_cggfc = uint(_dfed - _feag)
		} else {
			_cggfc = uint(8 - (_feag - _dfed))
		}
		_cdfa = 8 - _cggfc
		_cdcgf = _aeeg[_cggfc]
	}
	if (_ecba & 7) != 0 {
		_defa = true
		_bbfd = 8 - (_ecba & 7)
		_decdg = _aeeg[_bbfd]
		_fbcf = _ddfgc.RowStride*_eeaa + (_ecba >> 3)
		_fbadd = _ddbeg.RowStride*_ffbf + (_cgdc >> 3)
		_bdeb = 8 - (_cgdc & 7)
		if _bbfd > _bdeb {
			_fbbe = _eafd
			if _fdgg >= _feag {
				_fafbe = true
			}
		} else {
			_fbbe = _fbba
		}
	}
	if _fdgg < _bbfd {
		_fbaga = true
		_decdg &= _gbdb[8-_bbfd+_fdgg]
	}
	if !_fbaga {
		_dbdae = (_fdgg - _bbfd) >> 3
		if _dbdae != 0 {
			_aegeg = true
			_aadb = _ddfgc.RowStride*_eeaa + ((_ecba + _dfed) >> 3)
			_cacf = _ddbeg.RowStride*_ffbf + ((_cgdc + _dfed) >> 3)
		}
	}
	_dbga = (_ecba + _fdgg) & 7
	if !(_fbaga || _dbga == 0) {
		_cega = true
		_aagb = _gbdb[_dbga]
		_fdcg = _ddfgc.RowStride*_eeaa + ((_ecba + _dfed) >> 3) + _dbdae
		_fcbb = _ddbeg.RowStride*_ffbf + ((_cgdc + _dfed) >> 3) + _dbdae
		if _dbga > int(_cdfa) {
			_abba = true
		}
	}
	switch _ebeaa {
	case PixSrc:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], _cae, _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] = _cae
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], _cae, _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixNotSrc:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], ^_cae, _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] = ^_cae
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], ^_cae, _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixSrcOrDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], _cae|_ddfgc.Data[_fbcf], _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] |= _cae
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], _cae|_ddfgc.Data[_fdcg], _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixSrcAndDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], _cae&_ddfgc.Data[_fbcf], _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] &= _cae
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], _cae&_ddfgc.Data[_fdcg], _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixSrcXorDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], _cae^_ddfgc.Data[_fbcf], _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] ^= _cae
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], _cae^_ddfgc.Data[_fdcg], _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixNotSrcOrDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], ^_cae|_ddfgc.Data[_fbcf], _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] |= ^_cae
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], ^_cae|_ddfgc.Data[_fdcg], _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixNotSrcAndDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], ^_cae&_ddfgc.Data[_fbcf], _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] &= ^_cae
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], ^_cae&_ddfgc.Data[_fdcg], _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixSrcOrNotDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], _cae|^_ddfgc.Data[_fbcf], _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] = _cae | ^_ddfgc.Data[_aadb+_eafc]
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], _cae|^_ddfgc.Data[_fdcg], _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixSrcAndNotDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], _cae&^_ddfgc.Data[_fbcf], _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] = _cae &^ _ddfgc.Data[_aadb+_eafc]
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], _cae&^_ddfgc.Data[_fdcg], _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], ^(_cae | _ddfgc.Data[_fbcf]), _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] = ^(_cae | _ddfgc.Data[_aadb+_eafc])
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], ^(_cae | _ddfgc.Data[_fdcg]), _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], ^(_cae & _ddfgc.Data[_fbcf]), _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] = ^(_cae & _ddfgc.Data[_aadb+_eafc])
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], ^(_cae & _ddfgc.Data[_fdcg]), _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		if _defa {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				if _fbbe == _eafd {
					_cae = _ddbeg.Data[_fbadd] << _cggfc
					if _fafbe {
						_cae = _aegg(_cae, _ddbeg.Data[_fbadd+1]>>_cdfa, _cdcgf)
					}
				} else {
					_cae = _ddbeg.Data[_fbadd] >> _cdfa
				}
				_ddfgc.Data[_fbcf] = _aegg(_ddfgc.Data[_fbcf], ^(_cae ^ _ddfgc.Data[_fbcf]), _decdg)
				_fbcf += _ddfgc.RowStride
				_fbadd += _ddbeg.RowStride
			}
		}
		if _aegeg {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				for _eafc = 0; _eafc < _dbdae; _eafc++ {
					_cae = _aegg(_ddbeg.Data[_cacf+_eafc]<<_cggfc, _ddbeg.Data[_cacf+_eafc+1]>>_cdfa, _cdcgf)
					_ddfgc.Data[_aadb+_eafc] = ^(_cae ^ _ddfgc.Data[_aadb+_eafc])
				}
				_aadb += _ddfgc.RowStride
				_cacf += _ddbeg.RowStride
			}
		}
		if _cega {
			for _cgee = 0; _cgee < _ggcg; _cgee++ {
				_cae = _ddbeg.Data[_fcbb] << _cggfc
				if _abba {
					_cae = _aegg(_cae, _ddbeg.Data[_fcbb+1]>>_cdfa, _cdcgf)
				}
				_ddfgc.Data[_fdcg] = _aegg(_ddfgc.Data[_fdcg], ^(_cae ^ _ddfgc.Data[_fdcg]), _aagb)
				_fdcg += _ddfgc.RowStride
				_fcbb += _ddbeg.RowStride
			}
		}
	default:
		_gb.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _ebeaa)
		return _g.Error("\u0072a\u0073t\u0065\u0072\u004f\u0070\u0047e\u006e\u0065r\u0061\u006c\u004c\u006f\u0077", "\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}
func (_bcee MorphProcess) verify(_bedfg int, _fgeg, _badbg *int) error {
	const _bedb = "\u004d\u006f\u0072\u0070hP\u0072\u006f\u0063\u0065\u0073\u0073\u002e\u0076\u0065\u0072\u0069\u0066\u0079"
	switch _bcee.Operation {
	case MopDilation, MopErosion, MopOpening, MopClosing:
		if len(_bcee.Arguments) != 2 {
			return _g.Error(_bedb, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0064\u0027\u002c\u0020\u0027\u0065\u0027\u002c \u0027\u006f\u0027\u002c\u0020\u0027\u0063\u0027\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0032\u0020\u0061r\u0067\u0075\u006d\u0065\u006et\u0073")
		}
		_aeaa, _acadd := _bcee.getWidthHeight()
		if _aeaa <= 0 || _acadd <= 0 {
			return _g.Error(_bedb, "O\u0070er\u0061t\u0069o\u006e\u003a\u0020\u0027\u0064'\u002c\u0020\u0027e\u0027\u002c\u0020\u0027\u006f'\u002c\u0020\u0027c\u0027\u0020\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0062\u006f\u0074h w\u0069\u0064\u0074\u0068\u0020\u0061n\u0064\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0074\u006f\u0020b\u0065 \u003e\u003d\u0020\u0030")
		}
	case MopRankBinaryReduction:
		_ffed := len(_bcee.Arguments)
		*_fgeg += _ffed
		if _ffed < 1 || _ffed > 4 {
			return _g.Error(_bedb, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0072\u0027\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020\u006d\u006fs\u0074\u0020\u0034\u0020\u0061\u0072g\u0075\u006d\u0065n\u0074\u0073")
		}
		for _dafee := 0; _dafee < _ffed; _dafee++ {
			if _bcee.Arguments[_dafee] < 1 || _bcee.Arguments[_dafee] > 4 {
				return _g.Error(_bedb, "\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0052\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 \u00280\u002c\u0020\u0034\u003e")
			}
		}
	case MopReplicativeBinaryExpansion:
		if len(_bcee.Arguments) == 0 {
			return _g.Error(_bedb, "\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074i\u0076\u0065\u0042in\u0061\u0072\u0079\u0045\u0078\u0070a\u006e\u0073\u0069\u006f\u006e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020o\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_dfa := _bcee.Arguments[0]
		if _dfa != 2 && _dfa != 4 && _dfa != 8 {
			return _g.Error(_bedb, "R\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u0076\u0065\u0042\u0069\u006e\u0061\u0072\u0079\u0045\u0078\u0070\u0061\u006e\u0073\u0069\u006f\u006e\u0020m\u0075s\u0074\u0020\u0062\u0065 \u006f\u0066 \u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
		}
		*_fgeg -= _ada[_dfa/4]
	case MopAddBorder:
		if len(_bcee.Arguments) == 0 {
			return _g.Error(_bedb, "\u0041\u0064\u0064B\u006f\u0072\u0064\u0065r\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_ccf := _bcee.Arguments[0]
		if _bedfg > 0 {
			return _g.Error(_bedb, "\u0041\u0064\u0064\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020f\u0069\u0072\u0073\u0074\u0020\u006d\u006f\u0072\u0070\u0068\u0020\u0070\u0072o\u0063\u0065\u0073\u0073")
		}
		if _ccf < 1 {
			return _g.Error(_bedb, "\u0041\u0064\u0064\u0042o\u0072\u0064\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065 \u006co\u0077\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
		}
		*_badbg = _ccf
	}
	return nil
}
func CorrelationScoreSimple(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_abag float64, _bfcf error) {
	const _dgdac = "\u0043\u006f\u0072\u0072el\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0053\u0069\u006d\u0070l\u0065"
	if bm1 == nil || bm2 == nil {
		return _abag, _g.Error(_dgdac, "n\u0069l\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0073 \u0070\u0072\u006f\u0076id\u0065\u0064")
	}
	if tab == nil {
		return _abag, _g.Error(_dgdac, "\u0074\u0061\u0062\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if area1 == 0 || area2 == 0 {
		return _abag, _g.Error(_dgdac, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0061\u0072e\u0061\u0073\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u0020\u0030")
	}
	_degc, _fgg := bm1.Width, bm1.Height
	_bdee, _accc := bm2.Width, bm2.Height
	if _bfad(_degc-_bdee) > maxDiffW {
		return 0, nil
	}
	if _bfad(_fgg-_accc) > maxDiffH {
		return 0, nil
	}
	var _cbcc, _dbgd int
	if delX >= 0 {
		_cbcc = int(delX + 0.5)
	} else {
		_cbcc = int(delX - 0.5)
	}
	if delY >= 0 {
		_dbgd = int(delY + 0.5)
	} else {
		_dbgd = int(delY - 0.5)
	}
	_faga := bm1.createTemplate()
	if _bfcf = _faga.RasterOperation(_cbcc, _dbgd, _bdee, _accc, PixSrc, bm2, 0, 0); _bfcf != nil {
		return _abag, _g.Wrap(_bfcf, _dgdac, "\u0062m\u0032 \u0074\u006f\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _bfcf = _faga.RasterOperation(0, 0, _degc, _fgg, PixSrcAndDst, bm1, 0, 0); _bfcf != nil {
		return _abag, _g.Wrap(_bfcf, _dgdac, "b\u006d\u0031\u0020\u0061\u006e\u0064\u0020\u0062\u006d\u0054")
	}
	_edee := _faga.countPixels()
	_abag = float64(_edee) * float64(_edee) / (float64(area1) * float64(area2))
	return _abag, nil
}
func _cfag(_begb, _eeg byte, _eada CombinationOperator) byte {
	switch _eada {
	case CmbOpOr:
		return _eeg | _begb
	case CmbOpAnd:
		return _eeg & _begb
	case CmbOpXor:
		return _eeg ^ _begb
	case CmbOpXNor:
		return ^(_eeg ^ _begb)
	case CmbOpNot:
		return ^(_eeg)
	default:
		return _eeg
	}
}
func _cad(_ggeg, _bfc int) *Bitmap {
	return &Bitmap{Width: _ggeg, Height: _bfc, RowStride: (_ggeg + 7) >> 3}
}

type SizeSelection int

func (_cgge Points) GetIntY(i int) (int, error) {
	if i >= len(_cgge) {
		return 0, _g.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0059", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_cgge[i].Y), nil
}

type BitmapsArray struct {
	Values []*Bitmaps
	Boxes  []*_aa.Rectangle
}

func _fgfg(_bgab, _dgfd *Bitmap, _abafc *Selection) (*Bitmap, error) {
	const _cada = "c\u006c\u006f\u0073\u0065\u0042\u0069\u0074\u006d\u0061\u0070"
	var _egcg error
	if _bgab, _egcg = _cgccb(_bgab, _dgfd, _abafc); _egcg != nil {
		return nil, _egcg
	}
	_eeefe, _egcg := _ceed(nil, _dgfd, _abafc)
	if _egcg != nil {
		return nil, _g.Wrap(_egcg, _cada, "")
	}
	if _, _egcg = _babe(_bgab, _eeefe, _abafc); _egcg != nil {
		return nil, _g.Wrap(_egcg, _cada, "")
	}
	return _bgab, nil
}
func (_fee *Bitmap) addPadBits() (_dfc error) {
	const _deec = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0061\u0064\u0064\u0050\u0061d\u0042\u0069\u0074\u0073"
	_bafe := _fee.Width % 8
	if _bafe == 0 {
		return nil
	}
	_bbf := _fee.Width / 8
	_feee := _c.NewReader(_fee.Data)
	_feg := make([]byte, _fee.Height*_fee.RowStride)
	_gcgg := _c.NewWriterMSB(_feg)
	_bga := make([]byte, _bbf)
	var (
		_gef  int
		_dafe uint64
	)
	for _gef = 0; _gef < _fee.Height; _gef++ {
		if _, _dfc = _feee.Read(_bga); _dfc != nil {
			return _g.Wrap(_dfc, _deec, "\u0066u\u006c\u006c\u0020\u0062\u0079\u0074e")
		}
		if _, _dfc = _gcgg.Write(_bga); _dfc != nil {
			return _g.Wrap(_dfc, _deec, "\u0066\u0075\u006c\u006c\u0020\u0062\u0079\u0074\u0065\u0073")
		}
		if _dafe, _dfc = _feee.ReadBits(byte(_bafe)); _dfc != nil {
			return _g.Wrap(_dfc, _deec, "\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0069\u0074\u0073")
		}
		if _dfc = _gcgg.WriteByte(byte(_dafe) << uint(8-_bafe)); _dfc != nil {
			return _g.Wrap(_dfc, _deec, "\u006ca\u0073\u0074\u0020\u0062\u0079\u0074e")
		}
	}
	_fee.Data = _gcgg.Data()
	return nil
}
func _aegg(_cfgb, _aabf, _bgaee byte) byte { return (_cfgb &^ (_bgaee)) | (_aabf & _bgaee) }
func _gcbe(_dbac, _agac *Bitmap, _fded, _ecca int) (*Bitmap, error) {
	const _eebg = "\u0063\u006c\u006f\u0073\u0065\u0042\u0072\u0069\u0063\u006b"
	if _agac == nil {
		return nil, _g.Error(_eebg, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _fded < 1 || _ecca < 1 {
		return nil, _g.Error(_eebg, "\u0068S\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0053\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _fded == 1 && _ecca == 1 {
		return _agac.Copy(), nil
	}
	if _fded == 1 || _ecca == 1 {
		_fdga := SelCreateBrick(_ecca, _fded, _ecca/2, _fded/2, SelHit)
		var _geae error
		_dbac, _geae = _fgfg(_dbac, _agac, _fdga)
		if _geae != nil {
			return nil, _g.Wrap(_geae, _eebg, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _dbac, nil
	}
	_cbf := SelCreateBrick(1, _fded, 0, _fded/2, SelHit)
	_aace := SelCreateBrick(_ecca, 1, _ecca/2, 0, SelHit)
	_bfba, _cfec := _ceed(nil, _agac, _cbf)
	if _cfec != nil {
		return nil, _g.Wrap(_cfec, _eebg, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _dbac, _cfec = _ceed(_dbac, _bfba, _aace); _cfec != nil {
		return nil, _g.Wrap(_cfec, _eebg, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _, _cfec = _babe(_bfba, _dbac, _cbf); _cfec != nil {
		return nil, _g.Wrap(_cfec, _eebg, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	if _, _cfec = _babe(_dbac, _bfba, _aace); _cfec != nil {
		return nil, _g.Wrap(_cfec, _eebg, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _dbac, nil
}
func (_fca *Bitmap) String() string {
	var _aaef = "\u000a"
	for _egb := 0; _egb < _fca.Height; _egb++ {
		var _bba string
		for _caf := 0; _caf < _fca.Width; _caf++ {
			_dcce := _fca.GetPixel(_caf, _egb)
			if _dcce {
				_bba += "\u0031"
			} else {
				_bba += "\u0030"
			}
		}
		_aaef += _bba + "\u000a"
	}
	return _aaef
}
func _egcb(_ggedd *Bitmap) (_cecc *Bitmap, _bdc int, _gaaac error) {
	const _aacd = "\u0042i\u0074\u006d\u0061\u0070.\u0077\u006f\u0072\u0064\u004da\u0073k\u0042y\u0044\u0069\u006c\u0061\u0074\u0069\u006fn"
	if _ggedd == nil {
		return nil, 0, _g.Errorf(_aacd, "\u0027\u0073\u0027\u0020bi\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	var _bagfe, _acdf *Bitmap
	if _bagfe, _gaaac = _feea(nil, _ggedd); _gaaac != nil {
		return nil, 0, _g.Wrap(_gaaac, _aacd, "\u0063\u006f\u0070\u0079\u0020\u0027\u0073\u0027")
	}
	var (
		_dagg         [13]int
		_cbbd, _bbgee int
	)
	_eeef := 12
	_bdad := _dd.NewNumSlice(_eeef + 1)
	_dgeaa := _dd.NewNumSlice(_eeef + 1)
	var _dafeb *Boxes
	for _bcbd := 0; _bcbd <= _eeef; _bcbd++ {
		if _bcbd == 0 {
			if _acdf, _gaaac = _feea(nil, _bagfe); _gaaac != nil {
				return nil, 0, _g.Wrap(_gaaac, _aacd, "\u0066i\u0072\u0073\u0074\u0020\u0062\u006d2")
			}
		} else {
			if _acdf, _gaaac = _bbcc(_bagfe, MorphProcess{Operation: MopDilation, Arguments: []int{2, 1}}); _gaaac != nil {
				return nil, 0, _g.Wrap(_gaaac, _aacd, "\u0064\u0069\u006ca\u0074\u0069\u006f\u006e\u0020\u0062\u006d\u0032")
			}
		}
		if _dafeb, _gaaac = _acdf.connComponentsBB(4); _gaaac != nil {
			return nil, 0, _g.Wrap(_gaaac, _aacd, "")
		}
		_dagg[_bcbd] = len(*_dafeb)
		_bdad.AddInt(_dagg[_bcbd])
		switch _bcbd {
		case 0:
			_cbbd = _dagg[0]
		default:
			_bbgee = _dagg[_bcbd-1] - _dagg[_bcbd]
			_dgeaa.AddInt(_bbgee)
		}
		_bagfe = _acdf
	}
	_edc := true
	_afcf := 2
	var _gfcb, _bde int
	for _efagf := 1; _efagf < len(*_dgeaa); _efagf++ {
		if _gfcb, _gaaac = _bdad.GetInt(_efagf); _gaaac != nil {
			return nil, 0, _g.Wrap(_gaaac, _aacd, "\u0043\u0068\u0065\u0063ki\u006e\u0067\u0020\u0062\u0065\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0069o\u006e")
		}
		if _edc && _gfcb < int(0.3*float32(_cbbd)) {
			_afcf = _efagf + 1
			_edc = false
		}
		if _bbgee, _gaaac = _dgeaa.GetInt(_efagf); _gaaac != nil {
			return nil, 0, _g.Wrap(_gaaac, _aacd, "\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006ea\u0044\u0069\u0066\u0066")
		}
		if _bbgee > _bde {
			_bde = _bbgee
		}
	}
	_bgg := _ggedd.XResolution
	if _bgg == 0 {
		_bgg = 150
	}
	if _bgg > 110 {
		_afcf++
	}
	if _afcf < 2 {
		_gb.Log.Trace("J\u0042\u0049\u0047\u0032\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0042\u0065\u0073\u0074 \u0074\u006f\u0020\u006d\u0069\u006e\u0069\u006d\u0075\u006d a\u006c\u006c\u006fw\u0061b\u006c\u0065")
		_afcf = 2
	}
	_bdc = _afcf + 1
	if _cecc, _gaaac = _gcbe(nil, _ggedd, _afcf+1, 1); _gaaac != nil {
		return nil, 0, _g.Wrap(_gaaac, _aacd, "\u0067\u0065\u0074\u0074in\u0067\u0020\u006d\u0061\u0073\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	return _cecc, _bdc, nil
}
func _fgcf(_cffe *Bitmap, _baefc, _cegf int, _fbeg, _dabf int, _fcagg RasterOperator) {
	var (
		_aebga  bool
		_geaedf bool
		_bfgf   int
		_eabff  int
		_bedbg  int
		_ebdd   int
		_ggag   bool
		_dfaf   byte
	)
	_ebebc := 8 - (_baefc & 7)
	_facf := _aeeg[_ebebc]
	_eebb := _cffe.RowStride*_cegf + (_baefc >> 3)
	if _fbeg < _ebebc {
		_aebga = true
		_facf &= _gbdb[8-_ebebc+_fbeg]
	}
	if !_aebga {
		_bfgf = (_fbeg - _ebebc) >> 3
		if _bfgf != 0 {
			_geaedf = true
			_eabff = _eebb + 1
		}
	}
	_bedbg = (_baefc + _fbeg) & 7
	if !(_aebga || _bedbg == 0) {
		_ggag = true
		_dfaf = _gbdb[_bedbg]
		_ebdd = _eebb + 1 + _bfgf
	}
	var _bgdd, _cadcb int
	switch _fcagg {
	case PixClr:
		for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
			_cffe.Data[_eebb] = _aegg(_cffe.Data[_eebb], 0x0, _facf)
			_eebb += _cffe.RowStride
		}
		if _geaedf {
			for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
				for _cadcb = 0; _cadcb < _bfgf; _cadcb++ {
					_cffe.Data[_eabff+_cadcb] = 0x0
				}
				_eabff += _cffe.RowStride
			}
		}
		if _ggag {
			for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
				_cffe.Data[_ebdd] = _aegg(_cffe.Data[_ebdd], 0x0, _dfaf)
				_ebdd += _cffe.RowStride
			}
		}
	case PixSet:
		for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
			_cffe.Data[_eebb] = _aegg(_cffe.Data[_eebb], 0xff, _facf)
			_eebb += _cffe.RowStride
		}
		if _geaedf {
			for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
				for _cadcb = 0; _cadcb < _bfgf; _cadcb++ {
					_cffe.Data[_eabff+_cadcb] = 0xff
				}
				_eabff += _cffe.RowStride
			}
		}
		if _ggag {
			for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
				_cffe.Data[_ebdd] = _aegg(_cffe.Data[_ebdd], 0xff, _dfaf)
				_ebdd += _cffe.RowStride
			}
		}
	case PixNotDst:
		for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
			_cffe.Data[_eebb] = _aegg(_cffe.Data[_eebb], ^_cffe.Data[_eebb], _facf)
			_eebb += _cffe.RowStride
		}
		if _geaedf {
			for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
				for _cadcb = 0; _cadcb < _bfgf; _cadcb++ {
					_cffe.Data[_eabff+_cadcb] = ^(_cffe.Data[_eabff+_cadcb])
				}
				_eabff += _cffe.RowStride
			}
		}
		if _ggag {
			for _bgdd = 0; _bgdd < _dabf; _bgdd++ {
				_cffe.Data[_ebdd] = _aegg(_cffe.Data[_ebdd], ^_cffe.Data[_ebdd], _dfaf)
				_ebdd += _cffe.RowStride
			}
		}
	}
}

type shift int

func (_bfcb *ClassedPoints) GroupByY() ([]*ClassedPoints, error) {
	const _fggb = "\u0043\u006c\u0061\u0073se\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0072\u006f\u0075\u0070\u0042y\u0059"
	if _adbc := _bfcb.validateIntSlice(); _adbc != nil {
		return nil, _g.Wrap(_adbc, _fggb, "")
	}
	if _bfcb.IntSlice.Size() == 0 {
		return nil, _g.Error(_fggb, "\u004e\u006f\u0020\u0063la\u0073\u0073\u0065\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_bfcb.SortByY()
	var (
		_edac []*ClassedPoints
		_dfba int
	)
	_dbacb := -1
	var _dabbe *ClassedPoints
	for _gfac := 0; _gfac < len(_bfcb.IntSlice); _gfac++ {
		_dfba = int(_bfcb.YAtIndex(_gfac))
		if _dfba != _dbacb {
			_dabbe = &ClassedPoints{Points: _bfcb.Points}
			_dbacb = _dfba
			_edac = append(_edac, _dabbe)
		}
		_dabbe.IntSlice = append(_dabbe.IntSlice, _bfcb.IntSlice[_gfac])
	}
	for _, _ecdd := range _edac {
		_ecdd.SortByX()
	}
	return _edac, nil
}
func NewClassedPoints(points *Points, classes _dd.IntSlice) (*ClassedPoints, error) {
	const _dabg = "\u004e\u0065w\u0043\u006c\u0061s\u0073\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073"
	if points == nil {
		return nil, _g.Error(_dabg, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if classes == nil {
		return nil, _g.Error(_dabg, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0063\u006c\u0061ss\u0065\u0073")
	}
	_gabc := &ClassedPoints{Points: points, IntSlice: classes}
	if _dfde := _gabc.validateIntSlice(); _dfde != nil {
		return nil, _g.Wrap(_dfde, _dabg, "")
	}
	return _gabc, nil
}

const _afce = 5000

func TstDSymbol(t *_ba.T, scale ...int) *Bitmap {
	_effd, _bcbef := NewWithData(4, 5, []byte{0xf0, 0x90, 0x90, 0x90, 0xE0})
	_b.NoError(t, _bcbef)
	return TstGetScaledSymbol(t, _effd, scale...)
}
func _cbcd(_aaggg *_dd.Stack) (_abbg *fillSegment, _dagag error) {
	const _gagd = "\u0070\u006f\u0070\u0046\u0069\u006c\u006c\u0053\u0065g\u006d\u0065\u006e\u0074"
	if _aaggg == nil {
		return nil, _g.Error(_gagd, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _aaggg.Aux == nil {
		return nil, _g.Error(_gagd, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	_edde, _fdefc := _aaggg.Pop()
	if !_fdefc {
		return nil, nil
	}
	_fgfbg, _fdefc := _edde.(*fillSegment)
	if !_fdefc {
		return nil, _g.Error(_gagd, "\u0073\u0074\u0061ck\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020c\u006fn\u0074a\u0069n\u0020\u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_abbg = &fillSegment{_fgfbg._dgfa, _fgfbg._dade, _fgfbg._ecfb + _fgfbg._cgaee, _fgfbg._cgaee}
	_aaggg.Aux.Push(_fgfbg)
	return _abbg, nil
}

type fillSegment struct {
	_dgfa  int
	_dade  int
	_ecfb  int
	_cgaee int
}
type SizeComparison int

func (_fbed *ClassedPoints) validateIntSlice() error {
	const _fgcae = "\u0076\u0061l\u0069\u0064\u0061t\u0065\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065"
	for _, _decd := range _fbed.IntSlice {
		if _decd >= (_fbed.Points.Size()) {
			return _g.Errorf(_fgcae, "c\u006c\u0061\u0073\u0073\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006ci\u0064 \u0069\u006e\u0064\u0065x\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0064", _decd, _fbed.Points.Size())
		}
	}
	return nil
}
func _bgba() (_ddb []byte) {
	_ddb = make([]byte, 256)
	for _bdfe := 0; _bdfe < 256; _bdfe++ {
		_dadf := byte(_bdfe)
		_ddb[_dadf] = (_dadf & 0x01) | ((_dadf & 0x04) >> 1) | ((_dadf & 0x10) >> 2) | ((_dadf & 0x40) >> 3) | ((_dadf & 0x02) << 3) | ((_dadf & 0x08) << 2) | ((_dadf & 0x20) << 1) | (_dadf & 0x80)
	}
	return _ddb
}

var (
	_gbdb = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_aeeg = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func (_cgeb Points) Size() int { return len(_cgeb) }

type Boxes []*_aa.Rectangle

func _cgccb(_gcgcc, _ebcd *Bitmap, _dagae *Selection) (*Bitmap, error) {
	const _acgd = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0032"
	var _bffc, _bedaa int
	if _ebcd == nil {
		return nil, _g.Error(_acgd, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _dagae == nil {
		return nil, _g.Error(_acgd, "\u0073e\u006c \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_bffc = _dagae.Width
	_bedaa = _dagae.Height
	if _bffc == 0 || _bedaa == 0 {
		return nil, _g.Error(_acgd, "\u0073\u0065\u006c\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _gcgcc == nil {
		return _ebcd.createTemplate(), nil
	}
	if _gaca := _gcgcc.resizeImageData(_ebcd); _gaca != nil {
		return nil, _gaca
	}
	return _gcgcc, nil
}
func (_gfd *Bitmap) GetChocolateData() []byte {
	if _gfd.Color == Vanilla {
		_gfd.inverseData()
	}
	return _gfd.Data
}
func (_gafd *Bitmaps) selectByIndicator(_dbfc *_dd.NumSlice) (_edba *Bitmaps, _bbga error) {
	const _gaddf = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u0073\u0065\u006c\u0065c\u0074B\u0079I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _gafd == nil {
		return nil, _g.Error(_gaddf, "\u0027\u0062\u0027 b\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _dbfc == nil {
		return nil, _g.Error(_gaddf, "'\u006e\u0061\u0027\u0020\u0069\u006ed\u0069\u0063\u0061\u0074\u006f\u0072\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	if len(_gafd.Values) == 0 {
		return _gafd, nil
	}
	if len(*_dbfc) != len(_gafd.Values) {
		return nil, _g.Errorf(_gaddf, "\u006ea\u0020\u006ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064,\u0020\u0069\u0073\u0020\u0064\u0069\u0066\u0066\u0065re\u006e\u0074\u0020t\u0068\u0061n\u0020\u0062\u0069\u0074\u006d\u0061p\u0073\u003a \u0025\u0064", len(*_dbfc), len(_gafd.Values))
	}
	var _bbcbf, _eeega, _fgfaf int
	for _eeega = 0; _eeega < len(*_dbfc); _eeega++ {
		if _bbcbf, _bbga = _dbfc.GetInt(_eeega); _bbga != nil {
			return nil, _g.Wrap(_bbga, _gaddf, "f\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		if _bbcbf == 1 {
			_fgfaf++
		}
	}
	if _fgfaf == len(_gafd.Values) {
		return _gafd, nil
	}
	_edba = &Bitmaps{}
	_bgfb := len(_gafd.Values) == len(_gafd.Boxes)
	for _eeega = 0; _eeega < len(*_dbfc); _eeega++ {
		if _bbcbf = int((*_dbfc)[_eeega]); _bbcbf == 0 {
			continue
		}
		_edba.Values = append(_edba.Values, _gafd.Values[_eeega])
		if _bgfb {
			_edba.Boxes = append(_edba.Boxes, _gafd.Boxes[_eeega])
		}
	}
	return _edba, nil
}
func TstCSymbol(t *_ba.T) *Bitmap {
	t.Helper()
	_dbaec := New(6, 6)
	_b.NoError(t, _dbaec.SetPixel(1, 0, 1))
	_b.NoError(t, _dbaec.SetPixel(2, 0, 1))
	_b.NoError(t, _dbaec.SetPixel(3, 0, 1))
	_b.NoError(t, _dbaec.SetPixel(4, 0, 1))
	_b.NoError(t, _dbaec.SetPixel(0, 1, 1))
	_b.NoError(t, _dbaec.SetPixel(5, 1, 1))
	_b.NoError(t, _dbaec.SetPixel(0, 2, 1))
	_b.NoError(t, _dbaec.SetPixel(0, 3, 1))
	_b.NoError(t, _dbaec.SetPixel(0, 4, 1))
	_b.NoError(t, _dbaec.SetPixel(5, 4, 1))
	_b.NoError(t, _dbaec.SetPixel(1, 5, 1))
	_b.NoError(t, _dbaec.SetPixel(2, 5, 1))
	_b.NoError(t, _dbaec.SetPixel(3, 5, 1))
	_b.NoError(t, _dbaec.SetPixel(4, 5, 1))
	return _dbaec
}
func _bfad(_fbgg int) int {
	if _fbgg < 0 {
		return -_fbgg
	}
	return _fbgg
}
func (_ccbd *Points) Add(pt *Points) error {
	const _fcede = "\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0041\u0064\u0064"
	if _ccbd == nil {
		return _g.Error(_fcede, "\u0070o\u0069n\u0074\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if pt == nil {
		return _g.Error(_fcede, "a\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	*_ccbd = append(*_ccbd, *pt...)
	return nil
}
func (_adgg *Boxes) makeSizeIndicator(_gegg, _afeb int, _eeba LocationFilter, _ddaf SizeComparison) *_dd.NumSlice {
	_fadc := &_dd.NumSlice{}
	var _bceab, _ebea, _ddbc int
	for _, _bbce := range *_adgg {
		_bceab = 0
		_ebea, _ddbc = _bbce.Dx(), _bbce.Dy()
		switch _eeba {
		case LocSelectWidth:
			if (_ddaf == SizeSelectIfLT && _ebea < _gegg) || (_ddaf == SizeSelectIfGT && _ebea > _gegg) || (_ddaf == SizeSelectIfLTE && _ebea <= _gegg) || (_ddaf == SizeSelectIfGTE && _ebea >= _gegg) {
				_bceab = 1
			}
		case LocSelectHeight:
			if (_ddaf == SizeSelectIfLT && _ddbc < _afeb) || (_ddaf == SizeSelectIfGT && _ddbc > _afeb) || (_ddaf == SizeSelectIfLTE && _ddbc <= _afeb) || (_ddaf == SizeSelectIfGTE && _ddbc >= _afeb) {
				_bceab = 1
			}
		case LocSelectIfEither:
			if (_ddaf == SizeSelectIfLT && (_ddbc < _afeb || _ebea < _gegg)) || (_ddaf == SizeSelectIfGT && (_ddbc > _afeb || _ebea > _gegg)) || (_ddaf == SizeSelectIfLTE && (_ddbc <= _afeb || _ebea <= _gegg)) || (_ddaf == SizeSelectIfGTE && (_ddbc >= _afeb || _ebea >= _gegg)) {
				_bceab = 1
			}
		case LocSelectIfBoth:
			if (_ddaf == SizeSelectIfLT && (_ddbc < _afeb && _ebea < _gegg)) || (_ddaf == SizeSelectIfGT && (_ddbc > _afeb && _ebea > _gegg)) || (_ddaf == SizeSelectIfLTE && (_ddbc <= _afeb && _ebea <= _gegg)) || (_ddaf == SizeSelectIfGTE && (_ddbc >= _afeb && _ebea >= _gegg)) {
				_bceab = 1
			}
		}
		_fadc.AddInt(_bceab)
	}
	return _fadc
}
func _bbcc(_dbfbe *Bitmap, _eedb ...MorphProcess) (_gggfe *Bitmap, _fafc error) {
	const _bgff = "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065"
	if _dbfbe == nil {
		return nil, _g.Error(_bgff, "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065 \u0073\u006f\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if len(_eedb) == 0 {
		return nil, _g.Error(_bgff, "m\u006f\u0072\u0070\u0068\u0053\u0065q\u0075\u0065\u006e\u0063\u0065\u002c \u0073\u0065\u0071\u0075\u0065\u006e\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _fafc = _agabf(_eedb...); _fafc != nil {
		return nil, _g.Wrap(_fafc, _bgff, "")
	}
	var _bcbc, _gbg, _babg int
	_gggfe = _dbfbe.Copy()
	for _, _gedga := range _eedb {
		switch _gedga.Operation {
		case MopDilation:
			_bcbc, _gbg = _gedga.getWidthHeight()
			_gggfe, _fafc = DilateBrick(nil, _gggfe, _bcbc, _gbg)
			if _fafc != nil {
				return nil, _g.Wrap(_fafc, _bgff, "")
			}
		case MopErosion:
			_bcbc, _gbg = _gedga.getWidthHeight()
			_gggfe, _fafc = _ceba(nil, _gggfe, _bcbc, _gbg)
			if _fafc != nil {
				return nil, _g.Wrap(_fafc, _bgff, "")
			}
		case MopOpening:
			_bcbc, _gbg = _gedga.getWidthHeight()
			_gggfe, _fafc = _gbgf(nil, _gggfe, _bcbc, _gbg)
			if _fafc != nil {
				return nil, _g.Wrap(_fafc, _bgff, "")
			}
		case MopClosing:
			_bcbc, _gbg = _gedga.getWidthHeight()
			_gggfe, _fafc = _bgga(nil, _gggfe, _bcbc, _gbg)
			if _fafc != nil {
				return nil, _g.Wrap(_fafc, _bgff, "")
			}
		case MopRankBinaryReduction:
			_gggfe, _fafc = _cec(_gggfe, _gedga.Arguments...)
			if _fafc != nil {
				return nil, _g.Wrap(_fafc, _bgff, "")
			}
		case MopReplicativeBinaryExpansion:
			_gggfe, _fafc = _cffff(_gggfe, _gedga.Arguments[0])
			if _fafc != nil {
				return nil, _g.Wrap(_fafc, _bgff, "")
			}
		case MopAddBorder:
			_babg = _gedga.Arguments[0]
			_gggfe, _fafc = _gggfe.AddBorder(_babg, 0)
			if _fafc != nil {
				return nil, _g.Wrap(_fafc, _bgff, "")
			}
		default:
			return nil, _g.Error(_bgff, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u006fr\u0070\u0068\u004f\u0070\u0065\u0072\u0061ti\u006f\u006e\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u006f t\u0068\u0065 \u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065")
		}
	}
	if _babg > 0 {
		_gggfe, _fafc = _gggfe.RemoveBorder(_babg)
		if _fafc != nil {
			return nil, _g.Wrap(_fafc, _bgff, "\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u003e\u0020\u0030")
		}
	}
	return _gggfe, nil
}
func TstImageBitmap() *Bitmap { return _fbbbf.Copy() }
func (_ebafe *Bitmaps) makeSizeIndicator(_ebab, _eegf int, _dfdf LocationFilter, _dgcda SizeComparison) (_fegc *_dd.NumSlice, _bcdf error) {
	const _ecbdg = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u006d\u0061\u006b\u0065S\u0069z\u0065I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _ebafe == nil {
		return nil, _g.Error(_ecbdg, "\u0062\u0069\u0074ma\u0070\u0073\u0020\u0027\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch _dfdf {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _g.Errorf(_ecbdg, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", _dfdf)
	}
	switch _dgcda {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _g.Errorf(_ecbdg, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", _dgcda)
	}
	_fegc = &_dd.NumSlice{}
	var (
		_aece, _bdcf, _dacc int
		_bbeg               *Bitmap
	)
	for _, _bbeg = range _ebafe.Values {
		_aece = 0
		_bdcf, _dacc = _bbeg.Width, _bbeg.Height
		switch _dfdf {
		case LocSelectWidth:
			if (_dgcda == SizeSelectIfLT && _bdcf < _ebab) || (_dgcda == SizeSelectIfGT && _bdcf > _ebab) || (_dgcda == SizeSelectIfLTE && _bdcf <= _ebab) || (_dgcda == SizeSelectIfGTE && _bdcf >= _ebab) || (_dgcda == SizeSelectIfEQ && _bdcf == _ebab) {
				_aece = 1
			}
		case LocSelectHeight:
			if (_dgcda == SizeSelectIfLT && _dacc < _eegf) || (_dgcda == SizeSelectIfGT && _dacc > _eegf) || (_dgcda == SizeSelectIfLTE && _dacc <= _eegf) || (_dgcda == SizeSelectIfGTE && _dacc >= _eegf) || (_dgcda == SizeSelectIfEQ && _dacc == _eegf) {
				_aece = 1
			}
		case LocSelectIfEither:
			if (_dgcda == SizeSelectIfLT && (_bdcf < _ebab || _dacc < _eegf)) || (_dgcda == SizeSelectIfGT && (_bdcf > _ebab || _dacc > _eegf)) || (_dgcda == SizeSelectIfLTE && (_bdcf <= _ebab || _dacc <= _eegf)) || (_dgcda == SizeSelectIfGTE && (_bdcf >= _ebab || _dacc >= _eegf)) || (_dgcda == SizeSelectIfEQ && (_bdcf == _ebab || _dacc == _eegf)) {
				_aece = 1
			}
		case LocSelectIfBoth:
			if (_dgcda == SizeSelectIfLT && (_bdcf < _ebab && _dacc < _eegf)) || (_dgcda == SizeSelectIfGT && (_bdcf > _ebab && _dacc > _eegf)) || (_dgcda == SizeSelectIfLTE && (_bdcf <= _ebab && _dacc <= _eegf)) || (_dgcda == SizeSelectIfGTE && (_bdcf >= _ebab && _dacc >= _eegf)) || (_dgcda == SizeSelectIfEQ && (_bdcf == _ebab && _dacc == _eegf)) {
				_aece = 1
			}
		}
		_fegc.AddInt(_aece)
	}
	return _fegc, nil
}
func (_dfe *Bitmap) RemoveBorder(borderSize int) (*Bitmap, error) {
	if borderSize == 0 {
		return _dfe.Copy(), nil
	}
	_edf, _bff := _dfe.removeBorderGeneral(borderSize, borderSize, borderSize, borderSize)
	if _bff != nil {
		return nil, _g.Wrap(_bff, "\u0052\u0065\u006do\u0076\u0065\u0042\u006f\u0072\u0064\u0065\u0072", "")
	}
	return _edf, nil
}
func (_accg Points) YSorter() func(_gece, _baab int) bool {
	return func(_aaab, _ebgg int) bool { return _accg[_aaab].Y < _accg[_ebgg].Y }
}
func (_febd *Bitmap) ThresholdPixelSum(thresh int, tab8 []int) (_gcge bool, _gfea error) {
	const _fdff = "\u0042i\u0074\u006d\u0061\u0070\u002e\u0054\u0068\u0072\u0065\u0073\u0068o\u006c\u0064\u0050\u0069\u0078\u0065\u006c\u0053\u0075\u006d"
	if tab8 == nil {
		tab8 = _gdfg()
	}
	_cebc := _febd.Width >> 3
	_ead := _febd.Width & 7
	_dgg := byte(0xff << uint(8-_ead))
	var (
		_ceff, _eee, _gcbg, _afac int
		_cdb                      byte
	)
	for _ceff = 0; _ceff < _febd.Height; _ceff++ {
		_gcbg = _febd.RowStride * _ceff
		for _eee = 0; _eee < _cebc; _eee++ {
			_cdb, _gfea = _febd.GetByte(_gcbg + _eee)
			if _gfea != nil {
				return false, _g.Wrap(_gfea, _fdff, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
			}
			_afac += tab8[_cdb]
		}
		if _ead != 0 {
			_cdb, _gfea = _febd.GetByte(_gcbg + _eee)
			if _gfea != nil {
				return false, _g.Wrap(_gfea, _fdff, "p\u0061\u0072\u0074\u0069\u0061\u006c\u0042\u0079\u0074\u0065")
			}
			_cdb &= _dgg
			_afac += tab8[_cdb]
		}
		if _afac > thresh {
			return true, nil
		}
	}
	return _gcge, nil
}
func TstFrameBitmapData() []byte        { return _gbef.Data }
func _cddg(_dccea uint, _efc byte) byte { return _efc >> _dccea << _dccea }

const (
	Vanilla Color = iota
	Chocolate
)
const (
	ComponentConn Component = iota
	ComponentCharacters
	ComponentWords
)

func (_fdc *Bitmap) connComponentsBB(_dccd int) (_dca *Boxes, _addb error) {
	const _fcbg = "\u0042\u0069\u0074ma\u0070\u002e\u0063\u006f\u006e\u006e\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0042"
	if _dccd != 4 && _dccd != 8 {
		return nil, _g.Error(_fcbg, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _fdc.Zero() {
		return &Boxes{}, nil
	}
	_fdc.setPadBits(0)
	_dbcbg, _addb := _feea(nil, _fdc)
	if _addb != nil {
		return nil, _g.Wrap(_addb, _fcbg, "\u0062\u006d\u0031")
	}
	_eedca := &_dd.Stack{}
	_eedca.Aux = &_dd.Stack{}
	_dca = &Boxes{}
	var (
		_dded, _cbeg int
		_gfda        _aa.Point
		_cccc        bool
		_fgbf        *_aa.Rectangle
	)
	for {
		if _gfda, _cccc, _addb = _dbcbg.nextOnPixel(_cbeg, _dded); _addb != nil {
			return nil, _g.Wrap(_addb, _fcbg, "")
		}
		if !_cccc {
			break
		}
		if _fgbf, _addb = _gccg(_dbcbg, _eedca, _gfda.X, _gfda.Y, _dccd); _addb != nil {
			return nil, _g.Wrap(_addb, _fcbg, "")
		}
		if _addb = _dca.Add(_fgbf); _addb != nil {
			return nil, _g.Wrap(_addb, _fcbg, "")
		}
		_cbeg = _gfda.X
		_dded = _gfda.Y
	}
	return _dca, nil
}
func TstOSymbol(t *_ba.T, scale ...int) *Bitmap {
	_acdec, _bacc := NewWithData(4, 5, []byte{0xF0, 0x90, 0x90, 0x90, 0xF0})
	_b.NoError(t, _bacc)
	return TstGetScaledSymbol(t, _acdec, scale...)
}
func (_eabf *Bitmap) setTwoBytes(_cca int, _bgaa uint16) error {
	if _cca+1 > len(_eabf.Data)-1 {
		return _g.Errorf("s\u0065\u0074\u0054\u0077\u006f\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _cca)
	}
	_eabf.Data[_cca] = byte((_bgaa & 0xff00) >> 8)
	_eabf.Data[_cca+1] = byte(_bgaa & 0xff)
	return nil
}
func (_dfedb *Bitmaps) GroupByWidth() (*BitmapsArray, error) {
	const _bgfe = "\u0047\u0072\u006fu\u0070\u0042\u0079\u0057\u0069\u0064\u0074\u0068"
	if len(_dfedb.Values) == 0 {
		return nil, _g.Error(_bgfe, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_bbeee := &BitmapsArray{}
	_dfedb.SortByWidth()
	_gcac := -1
	_eegg := -1
	for _dafdb := 0; _dafdb < len(_dfedb.Values); _dafdb++ {
		_ebcb := _dfedb.Values[_dafdb].Width
		if _ebcb > _gcac {
			_gcac = _ebcb
			_eegg++
			_bbeee.Values = append(_bbeee.Values, &Bitmaps{})
		}
		_bbeee.Values[_eegg].AddBitmap(_dfedb.Values[_dafdb])
	}
	return _bbeee, nil
}
func (_dfea *Bitmaps) ClipToBitmap(s *Bitmap) (*Bitmaps, error) {
	const _fcdfb = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0043\u006c\u0069p\u0054\u006f\u0042\u0069tm\u0061\u0070"
	if _dfea == nil {
		return nil, _g.Error(_fcdfb, "\u0042\u0069\u0074\u006dap\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if s == nil {
		return nil, _g.Error(_fcdfb, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	_bdfa := len(_dfea.Values)
	_abfe := &Bitmaps{Values: make([]*Bitmap, _bdfa), Boxes: make([]*_aa.Rectangle, _bdfa)}
	var (
		_fbcfg, _gdca *Bitmap
		_efggb        *_aa.Rectangle
		_eeege        error
	)
	for _ccba := 0; _ccba < _bdfa; _ccba++ {
		if _fbcfg, _eeege = _dfea.GetBitmap(_ccba); _eeege != nil {
			return nil, _g.Wrap(_eeege, _fcdfb, "")
		}
		if _efggb, _eeege = _dfea.GetBox(_ccba); _eeege != nil {
			return nil, _g.Wrap(_eeege, _fcdfb, "")
		}
		if _gdca, _eeege = s.clipRectangle(_efggb, nil); _eeege != nil {
			return nil, _g.Wrap(_eeege, _fcdfb, "")
		}
		if _gdca, _eeege = _gdca.And(_fbcfg); _eeege != nil {
			return nil, _g.Wrap(_eeege, _fcdfb, "")
		}
		_abfe.Values[_ccba] = _gdca
		_abfe.Boxes[_ccba] = _efggb
	}
	return _abfe, nil
}

type Point struct{ X, Y float32 }

func _ccgf(_fgef *Bitmap, _acgg *_dd.Stack, _cage, _faab int) (_cbbdcd *_aa.Rectangle, _ffbed error) {
	const _aaga = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _fgef == nil {
		return nil, _g.Error(_aaga, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _acgg == nil {
		return nil, _g.Error(_aaga, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_defg, _fbbgbd := _fgef.Width, _fgef.Height
	_cagca := _defg - 1
	_edca := _fbbgbd - 1
	if _cage < 0 || _cage > _cagca || _faab < 0 || _faab > _edca || !_fgef.GetPixel(_cage, _faab) {
		return nil, nil
	}
	_gcd := _aa.Rect(100000, 100000, 0, 0)
	if _ffbed = _egcag(_acgg, _cage, _cage, _faab, 1, _edca, &_gcd); _ffbed != nil {
		return nil, _g.Wrap(_ffbed, _aaga, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _ffbed = _egcag(_acgg, _cage, _cage, _faab+1, -1, _edca, &_gcd); _ffbed != nil {
		return nil, _g.Wrap(_ffbed, _aaga, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_gcd.Min.X, _gcd.Max.X = _cage, _cage
	_gcd.Min.Y, _gcd.Max.Y = _faab, _faab
	var (
		_fadd *fillSegment
		_bdcb int
	)
	for _acgg.Len() > 0 {
		if _fadd, _ffbed = _cbcd(_acgg); _ffbed != nil {
			return nil, _g.Wrap(_ffbed, _aaga, "")
		}
		_faab = _fadd._ecfb
		for _cage = _fadd._dgfa - 1; _cage >= 0 && _fgef.GetPixel(_cage, _faab); _cage-- {
			if _ffbed = _fgef.SetPixel(_cage, _faab, 0); _ffbed != nil {
				return nil, _g.Wrap(_ffbed, _aaga, "\u0031s\u0074\u0020\u0073\u0065\u0074")
			}
		}
		if _cage >= _fadd._dgfa-1 {
			for {
				for _cage++; _cage <= _fadd._dade+1 && _cage <= _cagca && !_fgef.GetPixel(_cage, _faab); _cage++ {
				}
				_bdcb = _cage
				if !(_cage <= _fadd._dade+1 && _cage <= _cagca) {
					break
				}
				for ; _cage <= _cagca && _fgef.GetPixel(_cage, _faab); _cage++ {
					if _ffbed = _fgef.SetPixel(_cage, _faab, 0); _ffbed != nil {
						return nil, _g.Wrap(_ffbed, _aaga, "\u0032n\u0064\u0020\u0073\u0065\u0074")
					}
				}
				if _ffbed = _egcag(_acgg, _bdcb, _cage-1, _fadd._ecfb, _fadd._cgaee, _edca, &_gcd); _ffbed != nil {
					return nil, _g.Wrap(_ffbed, _aaga, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
				}
				if _cage > _fadd._dade {
					if _ffbed = _egcag(_acgg, _fadd._dade+1, _cage-1, _fadd._ecfb, -_fadd._cgaee, _edca, &_gcd); _ffbed != nil {
						return nil, _g.Wrap(_ffbed, _aaga, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
					}
				}
			}
			continue
		}
		_bdcb = _cage + 1
		if _bdcb < _fadd._dgfa {
			if _ffbed = _egcag(_acgg, _bdcb, _fadd._dgfa-1, _fadd._ecfb, -_fadd._cgaee, _edca, &_gcd); _ffbed != nil {
				return nil, _g.Wrap(_ffbed, _aaga, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
			}
		}
		_cage = _fadd._dgfa
		for {
			for ; _cage <= _cagca && _fgef.GetPixel(_cage, _faab); _cage++ {
				if _ffbed = _fgef.SetPixel(_cage, _faab, 0); _ffbed != nil {
					return nil, _g.Wrap(_ffbed, _aaga, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _ffbed = _egcag(_acgg, _bdcb, _cage-1, _fadd._ecfb, _fadd._cgaee, _edca, &_gcd); _ffbed != nil {
				return nil, _g.Wrap(_ffbed, _aaga, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _cage > _fadd._dade {
				if _ffbed = _egcag(_acgg, _fadd._dade+1, _cage-1, _fadd._ecfb, -_fadd._cgaee, _edca, &_gcd); _ffbed != nil {
					return nil, _g.Wrap(_ffbed, _aaga, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _cage++; _cage <= _fadd._dade+1 && _cage <= _cagca && !_fgef.GetPixel(_cage, _faab); _cage++ {
			}
			_bdcb = _cage
			if !(_cage <= _fadd._dade+1 && _cage <= _cagca) {
				break
			}
		}
	}
	_gcd.Max.X++
	_gcd.Max.Y++
	return &_gcd, nil
}
func (_agf *Bitmap) AddBorder(borderSize, val int) (*Bitmap, error) {
	if borderSize == 0 {
		return _agf.Copy(), nil
	}
	_abcf, _eff := _agf.addBorderGeneral(borderSize, borderSize, borderSize, borderSize, val)
	if _eff != nil {
		return nil, _g.Wrap(_eff, "\u0041d\u0064\u0042\u006f\u0072\u0064\u0065r", "")
	}
	return _abcf, nil
}
func (_bedff *ClassedPoints) Less(i, j int) bool { return _bedff._bacf(i, j) }
func (_adce *Bitmap) removeBorderGeneral(_dggf, _caad, _gcaf, _caaa int) (*Bitmap, error) {
	const _edd = "\u0072\u0065\u006d\u006fve\u0042\u006f\u0072\u0064\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _dggf < 0 || _caad < 0 || _gcaf < 0 || _caaa < 0 {
		return nil, _g.Error(_edd, "\u006e\u0065g\u0061\u0074\u0069\u0076\u0065\u0020\u0062\u0072\u006f\u0064\u0065\u0072\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u0076\u0061lu\u0065\u0073")
	}
	_cdae, _fdd := _adce.Width, _adce.Height
	_dgca := _cdae - _dggf - _caad
	_fgfda := _fdd - _gcaf - _caaa
	if _dgca <= 0 {
		return nil, _g.Errorf(_edd, "w\u0069\u0064\u0074\u0068: \u0025d\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u003e\u0020\u0030", _dgca)
	}
	if _fgfda <= 0 {
		return nil, _g.Errorf(_edd, "\u0068\u0065\u0069\u0067ht\u003a\u0020\u0025\u0064\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030", _fgfda)
	}
	_baa := New(_dgca, _fgfda)
	_baa.Color = _adce.Color
	_bccf := _baa.RasterOperation(0, 0, _dgca, _fgfda, PixSrc, _adce, _dggf, _gcaf)
	if _bccf != nil {
		return nil, _g.Wrap(_bccf, _edd, "")
	}
	return _baa, nil
}
func (_bgac *Bitmap) setAll() error {
	_bdg := _aee(_bgac, 0, 0, _bgac.Width, _bgac.Height, PixSet, nil, 0, 0)
	if _bdg != nil {
		return _g.Wrap(_bdg, "\u0073\u0065\u0074\u0041\u006c\u006c", "")
	}
	return nil
}
func _acdc(_abdf *Bitmap, _efbc *_dd.Stack, _aeba, _befe int) (_bccfa *_aa.Rectangle, _bffg error) {
	const _gdag = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _abdf == nil {
		return nil, _g.Error(_gdag, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _efbc == nil {
		return nil, _g.Error(_gdag, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_eddde, _gbbcd := _abdf.Width, _abdf.Height
	_cafdb := _eddde - 1
	_fefd := _gbbcd - 1
	if _aeba < 0 || _aeba > _cafdb || _befe < 0 || _befe > _fefd || !_abdf.GetPixel(_aeba, _befe) {
		return nil, nil
	}
	var _cbcfff *_aa.Rectangle
	_cbcfff, _bffg = Rect(100000, 100000, 0, 0)
	if _bffg != nil {
		return nil, _g.Wrap(_bffg, _gdag, "")
	}
	if _bffg = _egcag(_efbc, _aeba, _aeba, _befe, 1, _fefd, _cbcfff); _bffg != nil {
		return nil, _g.Wrap(_bffg, _gdag, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _bffg = _egcag(_efbc, _aeba, _aeba, _befe+1, -1, _fefd, _cbcfff); _bffg != nil {
		return nil, _g.Wrap(_bffg, _gdag, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_cbcfff.Min.X, _cbcfff.Max.X = _aeba, _aeba
	_cbcfff.Min.Y, _cbcfff.Max.Y = _befe, _befe
	var (
		_aagc  *fillSegment
		_cggca int
	)
	for _efbc.Len() > 0 {
		if _aagc, _bffg = _cbcd(_efbc); _bffg != nil {
			return nil, _g.Wrap(_bffg, _gdag, "")
		}
		_befe = _aagc._ecfb
		for _aeba = _aagc._dgfa; _aeba >= 0 && _abdf.GetPixel(_aeba, _befe); _aeba-- {
			if _bffg = _abdf.SetPixel(_aeba, _befe, 0); _bffg != nil {
				return nil, _g.Wrap(_bffg, _gdag, "")
			}
		}
		if _aeba >= _aagc._dgfa {
			for _aeba++; _aeba <= _aagc._dade && _aeba <= _cafdb && !_abdf.GetPixel(_aeba, _befe); _aeba++ {
			}
			_cggca = _aeba
			if !(_aeba <= _aagc._dade && _aeba <= _cafdb) {
				continue
			}
		} else {
			_cggca = _aeba + 1
			if _cggca < _aagc._dgfa-1 {
				if _bffg = _egcag(_efbc, _cggca, _aagc._dgfa-1, _aagc._ecfb, -_aagc._cgaee, _fefd, _cbcfff); _bffg != nil {
					return nil, _g.Wrap(_bffg, _gdag, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
				}
			}
			_aeba = _aagc._dgfa + 1
		}
		for {
			for ; _aeba <= _cafdb && _abdf.GetPixel(_aeba, _befe); _aeba++ {
				if _bffg = _abdf.SetPixel(_aeba, _befe, 0); _bffg != nil {
					return nil, _g.Wrap(_bffg, _gdag, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _bffg = _egcag(_efbc, _cggca, _aeba-1, _aagc._ecfb, _aagc._cgaee, _fefd, _cbcfff); _bffg != nil {
				return nil, _g.Wrap(_bffg, _gdag, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _aeba > _aagc._dade+1 {
				if _bffg = _egcag(_efbc, _aagc._dade+1, _aeba-1, _aagc._ecfb, -_aagc._cgaee, _fefd, _cbcfff); _bffg != nil {
					return nil, _g.Wrap(_bffg, _gdag, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _aeba++; _aeba <= _aagc._dade && _aeba <= _cafdb && !_abdf.GetPixel(_aeba, _befe); _aeba++ {
			}
			_cggca = _aeba
			if !(_aeba <= _aagc._dade && _aeba <= _cafdb) {
				break
			}
		}
	}
	_cbcfff.Max.X++
	_cbcfff.Max.Y++
	return _cbcfff, nil
}
func TstASymbol(t *_ba.T) *Bitmap {
	t.Helper()
	_aafb := New(6, 6)
	_b.NoError(t, _aafb.SetPixel(1, 0, 1))
	_b.NoError(t, _aafb.SetPixel(2, 0, 1))
	_b.NoError(t, _aafb.SetPixel(3, 0, 1))
	_b.NoError(t, _aafb.SetPixel(4, 0, 1))
	_b.NoError(t, _aafb.SetPixel(5, 1, 1))
	_b.NoError(t, _aafb.SetPixel(1, 2, 1))
	_b.NoError(t, _aafb.SetPixel(2, 2, 1))
	_b.NoError(t, _aafb.SetPixel(3, 2, 1))
	_b.NoError(t, _aafb.SetPixel(4, 2, 1))
	_b.NoError(t, _aafb.SetPixel(5, 2, 1))
	_b.NoError(t, _aafb.SetPixel(0, 3, 1))
	_b.NoError(t, _aafb.SetPixel(5, 3, 1))
	_b.NoError(t, _aafb.SetPixel(0, 4, 1))
	_b.NoError(t, _aafb.SetPixel(5, 4, 1))
	_b.NoError(t, _aafb.SetPixel(1, 5, 1))
	_b.NoError(t, _aafb.SetPixel(2, 5, 1))
	_b.NoError(t, _aafb.SetPixel(3, 5, 1))
	_b.NoError(t, _aafb.SetPixel(4, 5, 1))
	_b.NoError(t, _aafb.SetPixel(5, 5, 1))
	return _aafb
}
func TstWriteSymbols(t *_ba.T, bms *Bitmaps, src *Bitmap) {
	for _egcc := 0; _egcc < bms.Size(); _egcc++ {
		_bdbb := bms.Values[_egcc]
		_aeag := bms.Boxes[_egcc]
		_eeebd := src.RasterOperation(_aeag.Min.X, _aeag.Min.Y, _bdbb.Width, _bdbb.Height, PixSrc, _bdbb, 0, 0)
		_b.NoError(t, _eeebd)
	}
}
func (_decc *Bitmaps) AddBox(box *_aa.Rectangle) { _decc.Boxes = append(_decc.Boxes, box) }
func TstImageBitmapData() []byte                 { return _fbbbf.Data }
func _gbgf(_eabe, _cefd *Bitmap, _dbgge, _ebcg int) (*Bitmap, error) {
	const _dbacg = "\u006fp\u0065\u006e\u0042\u0072\u0069\u0063k"
	if _cefd == nil {
		return nil, _g.Error(_dbacg, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _dbgge < 1 && _ebcg < 1 {
		return nil, _g.Error(_dbacg, "\u0068\u0053\u0069\u007ae \u003c\u0020\u0031\u0020\u0026\u0026\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u003c \u0031")
	}
	if _dbgge == 1 && _ebcg == 1 {
		return _cefd.Copy(), nil
	}
	if _dbgge == 1 || _ebcg == 1 {
		var _ebeb error
		_eace := SelCreateBrick(_ebcg, _dbgge, _ebcg/2, _dbgge/2, SelHit)
		_eabe, _ebeb = _ggdf(_eabe, _cefd, _eace)
		if _ebeb != nil {
			return nil, _g.Wrap(_ebeb, _dbacg, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _eabe, nil
	}
	_gbgb := SelCreateBrick(1, _dbgge, 0, _dbgge/2, SelHit)
	_agbf := SelCreateBrick(_ebcg, 1, _ebcg/2, 0, SelHit)
	_fgfgf, _cgb := _babe(nil, _cefd, _gbgb)
	if _cgb != nil {
		return nil, _g.Wrap(_cgb, _dbacg, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_eabe, _cgb = _babe(_eabe, _fgfgf, _agbf)
	if _cgb != nil {
		return nil, _g.Wrap(_cgb, _dbacg, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	_, _cgb = _ceed(_fgfgf, _eabe, _gbgb)
	if _cgb != nil {
		return nil, _g.Wrap(_cgb, _dbacg, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_, _cgb = _ceed(_eabe, _fgfgf, _agbf)
	if _cgb != nil {
		return nil, _g.Wrap(_cgb, _dbacg, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _eabe, nil
}
func (_bbaa *Boxes) Get(i int) (*_aa.Rectangle, error) {
	const _bceb = "\u0042o\u0078\u0065\u0073\u002e\u0047\u0065t"
	if _bbaa == nil {
		return nil, _g.Error(_bceb, "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if i > len(*_bbaa)-1 {
		return nil, _g.Errorf(_bceb, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return (*_bbaa)[i], nil
}
func (_abfb *byHeight) Less(i, j int) bool { return _abfb.Values[i].Height < _abfb.Values[j].Height }
func (_adb *Bitmap) resizeImageData(_cdaeg *Bitmap) error {
	if _cdaeg == nil {
		return _g.Error("\u0072e\u0073i\u007a\u0065\u0049\u006d\u0061\u0067\u0065\u0044\u0061\u0074\u0061", "\u0073r\u0063 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _adb.SizesEqual(_cdaeg) {
		return nil
	}
	_adb.Data = make([]byte, len(_cdaeg.Data))
	_adb.Width = _cdaeg.Width
	_adb.Height = _cdaeg.Height
	_adb.RowStride = _cdaeg.RowStride
	return nil
}
func _bc(_gg, _db *Bitmap) (_gd error) {
	const _bg = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0034"
	_bgc := _db.RowStride
	_ged := _gg.RowStride
	_ab := _db.RowStride*4 - _gg.RowStride
	var (
		_aeg, _eb                               byte
		_cf                                     uint32
		_cff, _fgf, _dbc, _ee, _bcb, _dbf, _fgb int
	)
	for _dbc = 0; _dbc < _db.Height; _dbc++ {
		_cff = _dbc * _bgc
		_fgf = 4 * _dbc * _ged
		for _ee = 0; _ee < _bgc; _ee++ {
			_aeg = _db.Data[_cff+_ee]
			_cf = _bdce[_aeg]
			_dbf = _fgf + _ee*4
			if _ab != 0 && (_ee+1)*4 > _gg.RowStride {
				for _bcb = _ab; _bcb > 0; _bcb-- {
					_eb = byte((_cf >> uint(_bcb*8)) & 0xff)
					_fgb = _dbf + (_ab - _bcb)
					if _gd = _gg.SetByte(_fgb, _eb); _gd != nil {
						return _g.Wrapf(_gd, _bg, "D\u0069\u0066\u0066\u0065\u0072\u0065n\u0074\u0020\u0072\u006f\u0077\u0073\u0074\u0072\u0069d\u0065\u0073\u002e \u004b:\u0020\u0025\u0064", _bcb)
					}
				}
			} else if _gd = _gg.setFourBytes(_dbf, _cf); _gd != nil {
				return _g.Wrap(_gd, _bg, "")
			}
			if _gd = _gg.setFourBytes(_fgf+_ee*4, _bdce[_db.Data[_cff+_ee]]); _gd != nil {
				return _g.Wrap(_gd, _bg, "")
			}
		}
		for _bcb = 1; _bcb < 4; _bcb++ {
			for _ee = 0; _ee < _ged; _ee++ {
				if _gd = _gg.SetByte(_fgf+_bcb*_ged+_ee, _gg.Data[_fgf+_ee]); _gd != nil {
					return _g.Wrapf(_gd, _bg, "\u0063\u006f\u0070\u0079\u0020\u0027\u0071\u0075\u0061\u0064\u0072\u0061\u0062l\u0065\u0027\u0020\u006c\u0069\u006ee\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0079\u0074\u0065\u003a \u0027\u0025\u0064\u0027", _bcb, _ee)
				}
			}
		}
	}
	return nil
}
func SelCreateBrick(h, w int, cy, cx int, tp SelectionValue) *Selection {
	_abfa := _bgbc(h, w, "")
	_abfa.setOrigin(cy, cx)
	var _agdcg, _gbbg int
	for _agdcg = 0; _agdcg < h; _agdcg++ {
		for _gbbg = 0; _gbbg < w; _gbbg++ {
			_abfa.Data[_agdcg][_gbbg] = tp
		}
	}
	return _abfa
}
func (_fbb *Bitmap) GetBitOffset(x int) int { return x & 0x07 }
func CombineBytes(oldByte, newByte byte, op CombinationOperator) byte {
	return _cfag(oldByte, newByte, op)
}

const (
	CmbOpOr CombinationOperator = iota
	CmbOpAnd
	CmbOpXor
	CmbOpXNor
	CmbOpReplace
	CmbOpNot
)

func _aba(_bagg *Bitmap, _ddf int, _fcg []byte) (_cdc *Bitmap, _gcb error) {
	const _fga = "\u0072\u0065\u0064\u0075\u0063\u0065\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0032"
	if _bagg == nil {
		return nil, _g.Error(_fga, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _ddf < 1 || _ddf > 4 {
		return nil, _g.Error(_fga, "\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _bagg.Height <= 1 {
		return nil, _g.Errorf(_fga, "\u0073o\u0075\u0072c\u0065\u0020\u0068e\u0069\u0067\u0068\u0074\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061t\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027\u0020-\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", _bagg.Height)
	}
	_cdc = New(_bagg.Width/2, _bagg.Height/2)
	if _fcg == nil {
		_fcg = _bgba()
	}
	_eac := _efag(_bagg.RowStride, 2*_cdc.RowStride)
	switch _ddf {
	case 1:
		_gcb = _gdb(_bagg, _cdc, _ddf, _fcg, _eac)
	case 2:
		_gcb = _afg(_bagg, _cdc, _ddf, _fcg, _eac)
	case 3:
		_gcb = _bagb(_bagg, _cdc, _ddf, _fcg, _eac)
	case 4:
		_gcb = _fda(_bagg, _cdc, _ddf, _fcg, _eac)
	}
	if _gcb != nil {
		return nil, _gcb
	}
	return _cdc, nil
}
func (_ggf *Bitmap) addBorderGeneral(_cedf, _bbe, _fef, _cgcc int, _abcc int) (*Bitmap, error) {
	const _gga = "\u0061\u0064d\u0042\u006f\u0072d\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _cedf < 0 || _bbe < 0 || _fef < 0 || _cgcc < 0 {
		return nil, _g.Error(_gga, "n\u0065\u0067\u0061\u0074iv\u0065 \u0062\u006f\u0072\u0064\u0065r\u0020\u0061\u0064\u0064\u0065\u0064")
	}
	_faea, _gaa := _ggf.Width, _ggf.Height
	_eed := _faea + _cedf + _bbe
	_agc := _gaa + _fef + _cgcc
	_baga := New(_eed, _agc)
	_baga.Color = _ggf.Color
	_ceg := PixClr
	if _abcc > 0 {
		_ceg = PixSet
	}
	_agb := _baga.RasterOperation(0, 0, _cedf, _agc, _ceg, nil, 0, 0)
	if _agb != nil {
		return nil, _g.Wrap(_agb, _gga, "\u006c\u0065\u0066\u0074")
	}
	_agb = _baga.RasterOperation(_eed-_bbe, 0, _bbe, _agc, _ceg, nil, 0, 0)
	if _agb != nil {
		return nil, _g.Wrap(_agb, _gga, "\u0072\u0069\u0067h\u0074")
	}
	_agb = _baga.RasterOperation(0, 0, _eed, _fef, _ceg, nil, 0, 0)
	if _agb != nil {
		return nil, _g.Wrap(_agb, _gga, "\u0074\u006f\u0070")
	}
	_agb = _baga.RasterOperation(0, _agc-_cgcc, _eed, _cgcc, _ceg, nil, 0, 0)
	if _agb != nil {
		return nil, _g.Wrap(_agb, _gga, "\u0062\u006f\u0074\u0074\u006f\u006d")
	}
	_agb = _baga.RasterOperation(_cedf, _fef, _faea, _gaa, PixSrc, _ggf, 0, 0)
	if _agb != nil {
		return nil, _g.Wrap(_agb, _gga, "\u0063\u006f\u0070\u0079")
	}
	return _baga, nil
}
func _bgbc(_acfe, _bacg int, _fbgd string) *Selection {
	_bgad := &Selection{Height: _acfe, Width: _bacg, Name: _fbgd}
	_bgad.Data = make([][]SelectionValue, _acfe)
	for _aeae := 0; _aeae < _acfe; _aeae++ {
		_bgad.Data[_aeae] = make([]SelectionValue, _bacg)
	}
	return _bgad
}

type MorphOperation int

func NewWithUnpaddedData(width, height int, data []byte) (*Bitmap, error) {
	const _dbce = "\u004e\u0065\u0077\u0057it\u0068\u0055\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	_feb := _cad(width, height)
	_feb.Data = data
	if _ceb := ((width * height) + 7) >> 3; len(data) < _ceb {
		return nil, _g.Errorf(_dbce, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(data), _ceb)
	}
	if _ffe := _feb.addPadBits(); _ffe != nil {
		return nil, _g.Wrap(_ffe, _dbce, "")
	}
	return _feb, nil
}
func _dabb(_dbaf, _ggcf *Bitmap, _ebef, _aaefa, _ccda, _cgcb, _ccaag, _bgag, _bgf, _dbab int, _gcgc CombinationOperator, _dfee int) error {
	var _bgcd int
	_ecc := func() { _bgcd++; _ccda += _ggcf.RowStride; _cgcb += _dbaf.RowStride; _ccaag += _dbaf.RowStride }
	for _bgcd = _ebef; _bgcd < _aaefa; _ecc() {
		var _ffg uint16
		_fge := _ccda
		for _ceabb := _cgcb; _ceabb <= _ccaag; _ceabb++ {
			_eeb, _gadag := _ggcf.GetByte(_fge)
			if _gadag != nil {
				return _gadag
			}
			_dacf, _gadag := _dbaf.GetByte(_ceabb)
			if _gadag != nil {
				return _gadag
			}
			_ffg = (_ffg | (uint16(_dacf) & 0xff)) << uint(_dbab)
			_dacf = byte(_ffg >> 8)
			if _gadag = _ggcf.SetByte(_fge, _cfag(_eeb, _dacf, _gcgc)); _gadag != nil {
				return _gadag
			}
			_fge++
			_ffg <<= uint(_bgf)
			if _ceabb == _ccaag {
				_dacf = byte(_ffg >> (8 - uint8(_dbab)))
				if _dfee != 0 {
					_dacf = _cddg(uint(8+_bgag), _dacf)
				}
				_eeb, _gadag = _ggcf.GetByte(_fge)
				if _gadag != nil {
					return _gadag
				}
				if _gadag = _ggcf.SetByte(_fge, _cfag(_eeb, _dacf, _gcgc)); _gadag != nil {
					return _gadag
				}
			}
		}
	}
	return nil
}
func (_bbb *Boxes) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_acg *Boxes, _dce error) {
	const _gdac = "\u0042o\u0078e\u0073\u002e\u0053\u0065\u006ce\u0063\u0074B\u0079\u0053\u0069\u007a\u0065"
	if _bbb == nil {
		return nil, _g.Error(_gdac, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_bbb) == 0 {
		return _bbb, nil
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _g.Errorf(_gdac, "\u0069\u006e\u0076al\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE:
	default:
		return nil, _g.Errorf(_gdac, "i\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020t\u0079\u0070\u0065:\u0020'\u0025\u0064\u0027", tp)
	}
	_age := _bbb.makeSizeIndicator(width, height, tp, relation)
	_agcg, _dce := _bbb.selectWithIndicator(_age)
	if _dce != nil {
		return nil, _g.Wrap(_dce, _gdac, "")
	}
	return _agcg, nil
}

type Getter interface{ GetBitmap() *Bitmap }

func (_adbe *Bitmaps) CountPixels() *_dd.NumSlice {
	_adad := &_dd.NumSlice{}
	for _, _gadd := range _adbe.Values {
		_adad.AddInt(_gadd.CountPixels())
	}
	return _adad
}
func _ec(_gea *Bitmap, _ecd int) (*Bitmap, error) {
	const _bdb = "\u0065x\u0070a\u006e\u0064\u0042\u0069\u006ea\u0072\u0079P\u006f\u0077\u0065\u0072\u0032"
	if _gea == nil {
		return nil, _g.Error(_bdb, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ecd == 1 {
		return _feea(nil, _gea)
	}
	if _ecd != 2 && _ecd != 4 && _ecd != 8 {
		return nil, _g.Error(_bdb, "\u0066\u0061\u0063t\u006f\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069n\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_dg := _ecd * _gea.Width
	_cea := _ecd * _gea.Height
	_ad := New(_dg, _cea)
	var _ebb error
	switch _ecd {
	case 2:
		_ebb = _ge(_ad, _gea)
	case 4:
		_ebb = _bc(_ad, _gea)
	case 8:
		_ebb = _ga(_ad, _gea)
	}
	if _ebb != nil {
		return nil, _g.Wrap(_ebb, _bdb, "")
	}
	return _ad, nil
}
func (_adfef *byHeight) Len() int { return len(_adfef.Values) }
func (_bedg *Bitmap) countPixels() int {
	var (
		_bfb  int
		_gab  uint8
		_dfef byte
		_cfc  int
	)
	_gggb := _bedg.RowStride
	_abg := uint(_bedg.Width & 0x07)
	if _abg != 0 {
		_gab = uint8((0xff << (8 - _abg)) & 0xff)
		_gggb--
	}
	for _cgfec := 0; _cgfec < _bedg.Height; _cgfec++ {
		for _cfc = 0; _cfc < _gggb; _cfc++ {
			_dfef = _bedg.Data[_cgfec*_bedg.RowStride+_cfc]
			_bfb += int(_ebc[_dfef])
		}
		if _abg != 0 {
			_bfb += int(_ebc[_bedg.Data[_cgfec*_bedg.RowStride+_cfc]&_gab])
		}
	}
	return _bfb
}
func (_fgdaa *ClassedPoints) Swap(i, j int) {
	_fgdaa.IntSlice[i], _fgdaa.IntSlice[j] = _fgdaa.IntSlice[j], _fgdaa.IntSlice[i]
}
func (_accfg *Bitmap) setBit(_dgd int) { _accfg.Data[(_dgd >> 3)] |= 0x80 >> uint(_dgd&7) }
func TstAddSymbol(t *_ba.T, bms *Bitmaps, sym *Bitmap, x *int, y int, space int) {
	bms.AddBitmap(sym)
	_bcgec := _aa.Rect(*x, y, *x+sym.Width, y+sym.Height)
	bms.AddBox(&_bcgec)
	*x += sym.Width + space
}

type Color int

func (_bcge *ClassedPoints) Len() int { return _bcge.IntSlice.Size() }
func New(width, height int) *Bitmap {
	_ggd := _cad(width, height)
	_ggd.Data = make([]byte, height*_ggd.RowStride)
	return _ggd
}
func (_cfgfd Points) GetIntX(i int) (int, error) {
	if i >= len(_cfgfd) {
		return 0, _g.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0058", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_cfgfd[i].X), nil
}
func _ggda(_cead *Bitmap, _aecc, _ebbfg, _dagd, _bagd int, _adfe RasterOperator, _efagd *Bitmap, _dfgb, _ddbg int) error {
	var (
		_cebae       bool
		_acdb        bool
		_eddd        int
		_bedd        int
		_fbbb        int
		_edgf        bool
		_fgfbe       byte
		_feda        int
		_fcdf        int
		_ggcc        int
		_debb, _aabe int
	)
	_adec := 8 - (_aecc & 7)
	_ggdfg := _aeeg[_adec]
	_gefe := _cead.RowStride*_ebbfg + (_aecc >> 3)
	_fddg := _efagd.RowStride*_ddbg + (_dfgb >> 3)
	if _dagd < _adec {
		_cebae = true
		_ggdfg &= _gbdb[8-_adec+_dagd]
	}
	if !_cebae {
		_eddd = (_dagd - _adec) >> 3
		if _eddd > 0 {
			_acdb = true
			_bedd = _gefe + 1
			_fbbb = _fddg + 1
		}
	}
	_feda = (_aecc + _dagd) & 7
	if !(_cebae || _feda == 0) {
		_edgf = true
		_fgfbe = _gbdb[_feda]
		_fcdf = _gefe + 1 + _eddd
		_ggcc = _fddg + 1 + _eddd
	}
	switch _adfe {
	case PixSrc:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], _efagd.Data[_fddg], _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] = _efagd.Data[_fbbb+_aabe]
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], _efagd.Data[_ggcc], _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixNotSrc:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], ^_efagd.Data[_fddg], _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] = ^_efagd.Data[_fbbb+_aabe]
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], ^_efagd.Data[_ggcc], _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixSrcOrDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], _efagd.Data[_fddg]|_cead.Data[_gefe], _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] |= _efagd.Data[_fbbb+_aabe]
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], _efagd.Data[_ggcc]|_cead.Data[_fcdf], _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixSrcAndDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], _efagd.Data[_fddg]&_cead.Data[_gefe], _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] &= _efagd.Data[_fbbb+_aabe]
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], _efagd.Data[_ggcc]&_cead.Data[_fcdf], _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixSrcXorDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], _efagd.Data[_fddg]^_cead.Data[_gefe], _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] ^= _efagd.Data[_fbbb+_aabe]
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], _efagd.Data[_ggcc]^_cead.Data[_fcdf], _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixNotSrcOrDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], ^(_efagd.Data[_fddg])|_cead.Data[_gefe], _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] |= ^(_efagd.Data[_fbbb+_aabe])
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], ^(_efagd.Data[_ggcc])|_cead.Data[_fcdf], _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixNotSrcAndDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], ^(_efagd.Data[_fddg])&_cead.Data[_gefe], _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] &= ^_efagd.Data[_fbbb+_aabe]
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], ^(_efagd.Data[_ggcc])&_cead.Data[_fcdf], _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixSrcOrNotDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], _efagd.Data[_fddg]|^(_cead.Data[_gefe]), _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] = _efagd.Data[_fbbb+_aabe] | ^(_cead.Data[_bedd+_aabe])
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], _efagd.Data[_ggcc]|^(_cead.Data[_fcdf]), _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixSrcAndNotDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], _efagd.Data[_fddg]&^(_cead.Data[_gefe]), _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] = _efagd.Data[_fbbb+_aabe] &^ (_cead.Data[_bedd+_aabe])
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], _efagd.Data[_ggcc]&^(_cead.Data[_fcdf]), _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], ^(_efagd.Data[_fddg] | _cead.Data[_gefe]), _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] = ^(_efagd.Data[_fbbb+_aabe] | _cead.Data[_bedd+_aabe])
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], ^(_efagd.Data[_ggcc] | _cead.Data[_fcdf]), _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], ^(_efagd.Data[_fddg] & _cead.Data[_gefe]), _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] = ^(_efagd.Data[_fbbb+_aabe] & _cead.Data[_bedd+_aabe])
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], ^(_efagd.Data[_ggcc] & _cead.Data[_fcdf]), _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		for _debb = 0; _debb < _bagd; _debb++ {
			_cead.Data[_gefe] = _aegg(_cead.Data[_gefe], ^(_efagd.Data[_fddg] ^ _cead.Data[_gefe]), _ggdfg)
			_gefe += _cead.RowStride
			_fddg += _efagd.RowStride
		}
		if _acdb {
			for _debb = 0; _debb < _bagd; _debb++ {
				for _aabe = 0; _aabe < _eddd; _aabe++ {
					_cead.Data[_bedd+_aabe] = ^(_efagd.Data[_fbbb+_aabe] ^ _cead.Data[_bedd+_aabe])
				}
				_bedd += _cead.RowStride
				_fbbb += _efagd.RowStride
			}
		}
		if _edgf {
			for _debb = 0; _debb < _bagd; _debb++ {
				_cead.Data[_fcdf] = _aegg(_cead.Data[_fcdf], ^(_efagd.Data[_ggcc] ^ _cead.Data[_fcdf]), _fgfbe)
				_fcdf += _cead.RowStride
				_ggcc += _efagd.RowStride
			}
		}
	default:
		_gb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _adfe)
		return _g.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0056\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004c\u006f\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}

type LocationFilter int

func _afg(_deg, _egc *Bitmap, _dcca int, _fabd []byte, _bdd int) (_ace error) {
	const _aae = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0032"
	var (
		_dff, _bgbe, _cb, _ag, _afd, _adf, _bed, _fae int
		_fd, _aabb, _acbb, _afad                      uint32
		_cgc, _ffd                                    byte
		_gfb                                          uint16
	)
	_abc := make([]byte, 4)
	_dae := make([]byte, 4)
	for _cb = 0; _cb < _deg.Height-1; _cb, _ag = _cb+2, _ag+1 {
		_dff = _cb * _deg.RowStride
		_bgbe = _ag * _egc.RowStride
		for _afd, _adf = 0, 0; _afd < _bdd; _afd, _adf = _afd+4, _adf+1 {
			for _bed = 0; _bed < 4; _bed++ {
				_fae = _dff + _afd + _bed
				if _fae <= len(_deg.Data)-1 && _fae < _dff+_deg.RowStride {
					_abc[_bed] = _deg.Data[_fae]
				} else {
					_abc[_bed] = 0x00
				}
				_fae = _dff + _deg.RowStride + _afd + _bed
				if _fae <= len(_deg.Data)-1 && _fae < _dff+(2*_deg.RowStride) {
					_dae[_bed] = _deg.Data[_fae]
				} else {
					_dae[_bed] = 0x00
				}
			}
			_fd = _cc.BigEndian.Uint32(_abc)
			_aabb = _cc.BigEndian.Uint32(_dae)
			_acbb = _fd & _aabb
			_acbb |= _acbb << 1
			_afad = _fd | _aabb
			_afad &= _afad << 1
			_aabb = _acbb | _afad
			_aabb &= 0xaaaaaaaa
			_fd = _aabb | (_aabb << 7)
			_cgc = byte(_fd >> 24)
			_ffd = byte((_fd >> 8) & 0xff)
			_fae = _bgbe + _adf
			if _fae+1 == len(_egc.Data)-1 || _fae+1 >= _bgbe+_egc.RowStride {
				if _ace = _egc.SetByte(_fae, _fabd[_cgc]); _ace != nil {
					return _g.Wrapf(_ace, _aae, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _fae)
				}
			} else {
				_gfb = (uint16(_fabd[_cgc]) << 8) | uint16(_fabd[_ffd])
				if _ace = _egc.setTwoBytes(_fae, _gfb); _ace != nil {
					return _g.Wrapf(_ace, _aae, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fae)
				}
				_adf++
			}
		}
	}
	return nil
}
func (_ccef *Bitmaps) AddBitmap(bm *Bitmap) { _ccef.Values = append(_ccef.Values, bm) }
func (_eaed *Bitmaps) SortByWidth()         { _fcea := (*byWidth)(_eaed); _a.Sort(_fcea) }
func (_abcfg *Bitmap) GetComponents(components Component, maxWidth, maxHeight int) (_cdca *Bitmaps, _cadd *Boxes, _bbfc error) {
	const _dadg = "B\u0069t\u006d\u0061\u0070\u002e\u0047\u0065\u0074\u0043o\u006d\u0070\u006f\u006een\u0074\u0073"
	if _abcfg == nil {
		return nil, nil, _g.Error(_dadg, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0042\u0069\u0074\u006da\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064\u002e")
	}
	switch components {
	case ComponentConn, ComponentCharacters, ComponentWords:
	default:
		return nil, nil, _g.Error(_dadg, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074s\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072")
	}
	if _abcfg.Zero() {
		_cadd = &Boxes{}
		_cdca = &Bitmaps{}
		return _cdca, _cadd, nil
	}
	switch components {
	case ComponentConn:
		_cdca = &Bitmaps{}
		if _cadd, _bbfc = _abcfg.ConnComponents(_cdca, 8); _bbfc != nil {
			return nil, nil, _g.Wrap(_bbfc, _dadg, "\u006e\u006f \u0070\u0072\u0065p\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentCharacters:
		_cfae, _cegd := MorphSequence(_abcfg, MorphProcess{Operation: MopClosing, Arguments: []int{1, 6}})
		if _cegd != nil {
			return nil, nil, _g.Wrap(_cegd, _dadg, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _gb.Log.IsLogLevel(_gb.LogLevelTrace) {
			_gb.Log.Trace("\u0043o\u006d\u0070o\u006e\u0065\u006e\u0074C\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074ma\u0070\u0020\u0061f\u0074\u0065r\u0020\u0063\u006c\u006f\u0073\u0069n\u0067\u003a \u0025\u0073", _cfae.String())
		}
		_gaaa := &Bitmaps{}
		_cadd, _cegd = _cfae.ConnComponents(_gaaa, 8)
		if _cegd != nil {
			return nil, nil, _g.Wrap(_cegd, _dadg, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _gb.Log.IsLogLevel(_gb.LogLevelTrace) {
			_gb.Log.Trace("\u0043\u006f\u006d\u0070\u006f\u006ee\u006e\u0074\u0043\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0020a\u0066\u0074\u0065\u0072\u0020\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076i\u0074y\u003a\u0020\u0025\u0073", _gaaa.String())
		}
		if _cdca, _cegd = _gaaa.ClipToBitmap(_abcfg); _cegd != nil {
			return nil, nil, _g.Wrap(_cegd, _dadg, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentWords:
		_dbe := 1
		var _dbgg *Bitmap
		switch {
		case _abcfg.XResolution <= 200:
			_dbgg = _abcfg
		case _abcfg.XResolution <= 400:
			_dbe = 2
			_dbgg, _bbfc = _cec(_abcfg, 1, 0, 0, 0)
			if _bbfc != nil {
				return nil, nil, _g.Wrap(_bbfc, _dadg, "w\u006f\u0072\u0064\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0020\u002d \u0078\u0072\u0065s\u003c=\u0034\u0030\u0030")
			}
		default:
			_dbe = 4
			_dbgg, _bbfc = _cec(_abcfg, 1, 1, 0, 0)
			if _bbfc != nil {
				return nil, nil, _g.Wrap(_bbfc, _dadg, "\u0077\u006f\u0072\u0064 \u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073 \u002d \u0078\u0072\u0065\u0073\u0020\u003e\u00204\u0030\u0030")
			}
		}
		_ccbf, _, _gbbc := _egcb(_dbgg)
		if _gbbc != nil {
			return nil, nil, _g.Wrap(_gbbc, _dadg, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_gbcg, _gbbc := _cffff(_ccbf, _dbe)
		if _gbbc != nil {
			return nil, nil, _g.Wrap(_gbbc, _dadg, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_fefb := &Bitmaps{}
		if _cadd, _gbbc = _gbcg.ConnComponents(_fefb, 4); _gbbc != nil {
			return nil, nil, _g.Wrap(_gbbc, _dadg, "\u0077\u006f\u0072\u0064\u0020\u0070r\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u002c\u0020\u0063\u006f\u006en\u0065\u0063\u0074\u0020\u0065\u0078\u0070a\u006e\u0064\u0065\u0064")
		}
		if _cdca, _gbbc = _fefb.ClipToBitmap(_abcfg); _gbbc != nil {
			return nil, nil, _g.Wrap(_gbbc, _dadg, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
	}
	_cdca, _bbfc = _cdca.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _bbfc != nil {
		return nil, nil, _g.Wrap(_bbfc, _dadg, "")
	}
	_cadd, _bbfc = _cadd.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _bbfc != nil {
		return nil, nil, _g.Wrap(_bbfc, _dadg, "")
	}
	return _cdca, _cadd, nil
}
func _fgdg(_ebg, _bbee *Bitmap, _gfa, _facd, _fbbg, _feed, _aea int, _fedd CombinationOperator) error {
	var _dccab int
	_ecfc := func() { _dccab++; _fbbg += _bbee.RowStride; _feed += _ebg.RowStride; _aea += _ebg.RowStride }
	for _dccab = _gfa; _dccab < _facd; _ecfc() {
		_fcbca := _fbbg
		for _dfce := _feed; _dfce <= _aea; _dfce++ {
			_bgacf, _ebde := _bbee.GetByte(_fcbca)
			if _ebde != nil {
				return _ebde
			}
			_fbfg, _ebde := _ebg.GetByte(_dfce)
			if _ebde != nil {
				return _ebde
			}
			if _ebde = _bbee.SetByte(_fcbca, _cfag(_bgacf, _fbfg, _fedd)); _ebde != nil {
				return _ebde
			}
			_fcbca++
		}
	}
	return nil
}
func (_dbcf *BitmapsArray) GetBitmaps(i int) (*Bitmaps, error) {
	const _befed = "\u0042\u0069\u0074ma\u0070\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u0069\u0074\u006d\u0061\u0070\u0073"
	if _dbcf == nil {
		return nil, _g.Error(_befed, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_dbcf.Values)-1 {
		return nil, _g.Errorf(_befed, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _dbcf.Values[i], nil
}
func _baefb(_gcgac *Bitmap, _edfa, _cafa int, _cfacg, _gbab int, _dcfb RasterOperator) {
	var (
		_cde           int
		_cbeb          byte
		_abfc, _dccaba int
		_gfg           int
	)
	_afca := _cfacg >> 3
	_agda := _cfacg & 7
	if _agda > 0 {
		_cbeb = _gbdb[_agda]
	}
	_cde = _gcgac.RowStride*_cafa + (_edfa >> 3)
	switch _dcfb {
	case PixClr:
		for _abfc = 0; _abfc < _gbab; _abfc++ {
			_gfg = _cde + _abfc*_gcgac.RowStride
			for _dccaba = 0; _dccaba < _afca; _dccaba++ {
				_gcgac.Data[_gfg] = 0x0
				_gfg++
			}
			if _agda > 0 {
				_gcgac.Data[_gfg] = _aegg(_gcgac.Data[_gfg], 0x0, _cbeb)
			}
		}
	case PixSet:
		for _abfc = 0; _abfc < _gbab; _abfc++ {
			_gfg = _cde + _abfc*_gcgac.RowStride
			for _dccaba = 0; _dccaba < _afca; _dccaba++ {
				_gcgac.Data[_gfg] = 0xff
				_gfg++
			}
			if _agda > 0 {
				_gcgac.Data[_gfg] = _aegg(_gcgac.Data[_gfg], 0xff, _cbeb)
			}
		}
	case PixNotDst:
		for _abfc = 0; _abfc < _gbab; _abfc++ {
			_gfg = _cde + _abfc*_gcgac.RowStride
			for _dccaba = 0; _dccaba < _afca; _dccaba++ {
				_gcgac.Data[_gfg] = ^_gcgac.Data[_gfg]
				_gfg++
			}
			if _agda > 0 {
				_gcgac.Data[_gfg] = _aegg(_gcgac.Data[_gfg], ^_gcgac.Data[_gfg], _cbeb)
			}
		}
	}
}
func (_edda *Bitmaps) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_cebfe *Bitmaps, _gfae error) {
	const _bfcba = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0053\u0065\u006ce\u0063\u0074\u0042\u0079Si\u007a\u0065"
	if _edda == nil {
		return nil, _g.Error(_bfcba, "\u0027\u0062\u0027 B\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _g.Errorf(_bfcba, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _g.Errorf(_bfcba, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", relation)
	}
	_ccge, _gfae := _edda.makeSizeIndicator(width, height, tp, relation)
	if _gfae != nil {
		return nil, _g.Wrap(_gfae, _bfcba, "")
	}
	_cebfe, _gfae = _edda.selectByIndicator(_ccge)
	if _gfae != nil {
		return nil, _g.Wrap(_gfae, _bfcba, "")
	}
	return _cebfe, nil
}
func (_eccf *Bitmaps) WidthSorter() func(_edgg, _fdgbg int) bool {
	return func(_ccbe, _cbcfb int) bool { return _eccf.Values[_ccbe].Width < _eccf.Values[_cbcfb].Width }
}
func Blit(src *Bitmap, dst *Bitmap, x, y int, op CombinationOperator) error {
	var _ccaa, _faac int
	_abe := src.RowStride - 1
	if x < 0 {
		_faac = -x
		x = 0
	} else if x+src.Width > dst.Width {
		_abe -= src.Width + x - dst.Width
	}
	if y < 0 {
		_ccaa = -y
		y = 0
		_faac += src.RowStride
		_abe += src.RowStride
	} else if y+src.Height > dst.Height {
		_ccaa = src.Height + y - dst.Height
	}
	var (
		_ecad int
		_cfb  error
	)
	_cggd := x & 0x07
	_dfcc := 8 - _cggd
	_edb := src.Width & 0x07
	_gbdfg := _dfcc - _edb
	_ggcb := _dfcc&0x07 != 0
	_fefg := src.Width <= ((_abe-_faac)<<3)+_dfcc
	_gafg := dst.GetByteIndex(x, y)
	_cfdg := _ccaa + dst.Height
	if src.Height > _cfdg {
		_ecad = _cfdg
	} else {
		_ecad = src.Height
	}
	switch {
	case !_ggcb:
		_cfb = _fgdg(src, dst, _ccaa, _ecad, _gafg, _faac, _abe, op)
	case _fefg:
		_cfb = _gbbb(src, dst, _ccaa, _ecad, _gafg, _faac, _abe, _gbdfg, _cggd, _dfcc, op)
	default:
		_cfb = _dabb(src, dst, _ccaa, _ecad, _gafg, _faac, _abe, _gbdfg, _cggd, _dfcc, op, _edb)
	}
	return _cfb
}
func TstRSymbol(t *_ba.T, scale ...int) *Bitmap {
	_gceca, _dfeea := NewWithData(4, 5, []byte{0xF0, 0x90, 0xF0, 0xA0, 0x90})
	_b.NoError(t, _dfeea)
	return TstGetScaledSymbol(t, _gceca, scale...)
}
func Copy(d, s *Bitmap) (*Bitmap, error) { return _feea(d, s) }
func NewWithData(width, height int, data []byte) (*Bitmap, error) {
	const _fbe = "N\u0065\u0077\u0057\u0069\u0074\u0068\u0044\u0061\u0074\u0061"
	_dbd := _cad(width, height)
	_dbd.Data = data
	if len(data) < height*_dbd.RowStride {
		return nil, _g.Errorf(_fbe, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u003a \u0025\u0064\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u003a\u0020\u0025\u0064", len(data), height*_dbd.RowStride)
	}
	return _dbd, nil
}
func _bagb(_beg, _cbc *Bitmap, _bee int, _gce []byte, _ef int) (_aac error) {
	const _cbcf = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0033"
	var (
		_fgag, _adc, _bcgg, _aca, _bagf, _ffdf, _cbe, _dffg int
		_cgf, _dgcd, _cga, _cac                             uint32
		_dde, _afb                                          byte
		_aecf                                               uint16
	)
	_cbb := make([]byte, 4)
	_fde := make([]byte, 4)
	for _bcgg = 0; _bcgg < _beg.Height-1; _bcgg, _aca = _bcgg+2, _aca+1 {
		_fgag = _bcgg * _beg.RowStride
		_adc = _aca * _cbc.RowStride
		for _bagf, _ffdf = 0, 0; _bagf < _ef; _bagf, _ffdf = _bagf+4, _ffdf+1 {
			for _cbe = 0; _cbe < 4; _cbe++ {
				_dffg = _fgag + _bagf + _cbe
				if _dffg <= len(_beg.Data)-1 && _dffg < _fgag+_beg.RowStride {
					_cbb[_cbe] = _beg.Data[_dffg]
				} else {
					_cbb[_cbe] = 0x00
				}
				_dffg = _fgag + _beg.RowStride + _bagf + _cbe
				if _dffg <= len(_beg.Data)-1 && _dffg < _fgag+(2*_beg.RowStride) {
					_fde[_cbe] = _beg.Data[_dffg]
				} else {
					_fde[_cbe] = 0x00
				}
			}
			_cgf = _cc.BigEndian.Uint32(_cbb)
			_dgcd = _cc.BigEndian.Uint32(_fde)
			_cga = _cgf & _dgcd
			_cga |= _cga << 1
			_cac = _cgf | _dgcd
			_cac &= _cac << 1
			_dgcd = _cga & _cac
			_dgcd &= 0xaaaaaaaa
			_cgf = _dgcd | (_dgcd << 7)
			_dde = byte(_cgf >> 24)
			_afb = byte((_cgf >> 8) & 0xff)
			_dffg = _adc + _ffdf
			if _dffg+1 == len(_cbc.Data)-1 || _dffg+1 >= _adc+_cbc.RowStride {
				if _aac = _cbc.SetByte(_dffg, _gce[_dde]); _aac != nil {
					return _g.Wrapf(_aac, _cbcf, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _dffg)
				}
			} else {
				_aecf = (uint16(_gce[_dde]) << 8) | uint16(_gce[_afb])
				if _aac = _cbc.setTwoBytes(_dffg, _aecf); _aac != nil {
					return _g.Wrapf(_aac, _cbcf, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _dffg)
				}
				_ffdf++
			}
		}
	}
	return nil
}

var MorphBC BoundaryCondition

func (_fbbfd *ClassedPoints) XAtIndex(i int) float32 { return (*_fbbfd.Points)[_fbbfd.IntSlice[i]].X }
func (_cfca *Bitmap) nextOnPixel(_fece, _gada int) (_faa _aa.Point, _bcde bool, _abaa error) {
	const _gcbbb = "n\u0065\u0078\u0074\u004f\u006e\u0050\u0069\u0078\u0065\u006c"
	_faa, _bcde, _abaa = _cfca.nextOnPixelLow(_cfca.Width, _cfca.Height, _cfca.RowStride, _fece, _gada)
	if _abaa != nil {
		return _faa, false, _g.Wrap(_abaa, _gcbbb, "")
	}
	return _faa, _bcde, nil
}
func (_dfbfd *byWidth) Less(i, j int) bool { return _dfbfd.Values[i].Width < _dfbfd.Values[j].Width }

type ClassedPoints struct {
	*Points
	_dd.IntSlice
	_bacf func(_cafd, _eegec int) bool
}

func _ceed(_gbac *Bitmap, _eag *Bitmap, _fced *Selection) (*Bitmap, error) {
	var (
		_gcfa *Bitmap
		_bfge error
	)
	_gbac, _bfge = _faaf(_gbac, _eag, _fced, &_gcfa)
	if _bfge != nil {
		return nil, _bfge
	}
	if _bfge = _gbac.clearAll(); _bfge != nil {
		return nil, _bfge
	}
	var _gdgb SelectionValue
	for _cegbd := 0; _cegbd < _fced.Height; _cegbd++ {
		for _eefa := 0; _eefa < _fced.Width; _eefa++ {
			_gdgb = _fced.Data[_cegbd][_eefa]
			if _gdgb == SelHit {
				if _bfge = _gbac.RasterOperation(_eefa-_fced.Cx, _cegbd-_fced.Cy, _eag.Width, _eag.Height, PixSrcOrDst, _gcfa, 0, 0); _bfge != nil {
					return nil, _bfge
				}
			}
		}
	}
	return _gbac, nil
}
func _baf() (_fcb [256]uint64) {
	for _fb := 0; _fb < 256; _fb++ {
		if _fb&0x01 != 0 {
			_fcb[_fb] |= 0xff
		}
		if _fb&0x02 != 0 {
			_fcb[_fb] |= 0xff00
		}
		if _fb&0x04 != 0 {
			_fcb[_fb] |= 0xff0000
		}
		if _fb&0x08 != 0 {
			_fcb[_fb] |= 0xff000000
		}
		if _fb&0x10 != 0 {
			_fcb[_fb] |= 0xff00000000
		}
		if _fb&0x20 != 0 {
			_fcb[_fb] |= 0xff0000000000
		}
		if _fb&0x40 != 0 {
			_fcb[_fb] |= 0xff000000000000
		}
		if _fb&0x80 != 0 {
			_fcb[_fb] |= 0xff00000000000000
		}
	}
	return _fcb
}
func (_dfgf Points) Get(i int) (Point, error) {
	if i > len(_dfgf)-1 {
		return Point{}, _g.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _dfgf[i], nil
}
func TstImageBitmapInverseData() []byte {
	_ebed := _fbbbf.Copy()
	_ebed.InverseData()
	return _ebed.Data
}
func _dccg() (_cef [256]uint32) {
	for _acde := 0; _acde < 256; _acde++ {
		if _acde&0x01 != 0 {
			_cef[_acde] |= 0xf
		}
		if _acde&0x02 != 0 {
			_cef[_acde] |= 0xf0
		}
		if _acde&0x04 != 0 {
			_cef[_acde] |= 0xf00
		}
		if _acde&0x08 != 0 {
			_cef[_acde] |= 0xf000
		}
		if _acde&0x10 != 0 {
			_cef[_acde] |= 0xf0000
		}
		if _acde&0x20 != 0 {
			_cef[_acde] |= 0xf00000
		}
		if _acde&0x40 != 0 {
			_cef[_acde] |= 0xf000000
		}
		if _acde&0x80 != 0 {
			_cef[_acde] |= 0xf0000000
		}
	}
	return _cef
}
func (_gcbc *Bitmaps) GetBitmap(i int) (*Bitmap, error) {
	const _gadde = "\u0047e\u0074\u0042\u0069\u0074\u006d\u0061p"
	if _gcbc == nil {
		return nil, _g.Error(_gadde, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if i > len(_gcbc.Values)-1 {
		return nil, _g.Errorf(_gadde, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _gcbc.Values[i], nil
}
func _fafa(_cdfc, _gdc, _ddaa *Bitmap, _gfad int) (*Bitmap, error) {
	const _dgec = "\u0073\u0065\u0065\u0064\u0046\u0069\u006c\u006c\u0042i\u006e\u0061\u0072\u0079"
	if _gdc == nil {
		return nil, _g.Error(_dgec, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _ddaa == nil {
		return nil, _g.Error(_dgec, "'\u006da\u0073\u006b\u0027\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _gfad != 4 && _gfad != 8 {
		return nil, _g.Error(_dgec, "\u0063\u006f\u006en\u0065\u0063\u0074\u0069v\u0069\u0074\u0079\u0020\u006e\u006f\u0074 \u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0034\u002c\u0038\u007d")
	}
	var _beac error
	_cdfc, _beac = _feea(_cdfc, _gdc)
	if _beac != nil {
		return nil, _g.Wrap(_beac, _dgec, "\u0063o\u0070y\u0020\u0073\u006f\u0075\u0072c\u0065\u0020t\u006f\u0020\u0027\u0064\u0027")
	}
	_acaf := _gdc.createTemplate()
	_ddaa.setPadBits(0)
	for _cggc := 0; _cggc < _afce; _cggc++ {
		_acaf, _beac = _feea(_acaf, _cdfc)
		if _beac != nil {
			return nil, _g.Wrapf(_beac, _dgec, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _cggc)
		}
		if _beac = _bdga(_cdfc, _ddaa, _gfad); _beac != nil {
			return nil, _g.Wrapf(_beac, _dgec, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _cggc)
		}
		if _acaf.Equals(_cdfc) {
			break
		}
	}
	return _cdfc, nil
}
func _gac(_da *Bitmap, _acd, _fc int) (*Bitmap, error) {
	const _ebf = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0052\u0065p\u006c\u0069\u0063\u0061\u0074\u0065"
	if _da == nil {
		return nil, _g.Error(_ebf, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _acd <= 0 || _fc <= 0 {
		return nil, _g.Error(_ebf, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _acd == _fc {
		if _acd == 1 {
			_fgfa, _cg := _feea(nil, _da)
			if _cg != nil {
				return nil, _g.Wrap(_cg, _ebf, "\u0078\u0046\u0061\u0063\u0074\u0020\u003d\u003d\u0020y\u0046\u0061\u0063\u0074")
			}
			return _fgfa, nil
		}
		if _acd == 2 || _acd == 4 || _acd == 8 {
			_bge, _eg := _ec(_da, _acd)
			if _eg != nil {
				return nil, _g.Wrap(_eg, _ebf, "\u0078\u0046a\u0063\u0074\u0020i\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
			}
			return _bge, nil
		}
	}
	_cd := _acd * _da.Width
	_af := _fc * _da.Height
	_fcd := New(_cd, _af)
	_bce := _fcd.RowStride
	var (
		_fa, _ea, _bfa, _ddcb, _bec int
		_bfe                        byte
		_caa                        error
	)
	for _ea = 0; _ea < _da.Height; _ea++ {
		_fa = _fc * _ea * _bce
		for _bfa = 0; _bfa < _da.Width; _bfa++ {
			if _df := _da.GetPixel(_bfa, _ea); _df {
				_bec = _acd * _bfa
				for _ddcb = 0; _ddcb < _acd; _ddcb++ {
					_fcd.setBit(_fa*8 + _bec + _ddcb)
				}
			}
		}
		for _ddcb = 1; _ddcb < _fc; _ddcb++ {
			_cfff := _fa + _ddcb*_bce
			for _eba := 0; _eba < _bce; _eba++ {
				if _bfe, _caa = _fcd.GetByte(_fa + _eba); _caa != nil {
					return nil, _g.Wrapf(_caa, _ebf, "\u0072\u0065\u0070\u006cic\u0061\u0074\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0027\u0025d\u0027", _ddcb)
				}
				if _caa = _fcd.SetByte(_cfff+_eba, _bfe); _caa != nil {
					return nil, _g.Wrap(_caa, _ebf, "\u0053\u0065\u0074\u0074in\u0067\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
				}
			}
		}
	}
	return _fcd, nil
}
func (_cge *Bitmap) AddBorderGeneral(left, right, top, bot int, val int) (*Bitmap, error) {
	return _cge.addBorderGeneral(left, right, top, bot, val)
}

var _ebc [256]uint8

func _gge(_bgd *Bitmap, _gf *Bitmap, _bgdf int) (_bag error) {
	const _fe = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0050\u006fw\u0065\u0072\u0032\u004c\u006f\u0077"
	switch _bgdf {
	case 2:
		_bag = _ge(_bgd, _gf)
	case 4:
		_bag = _bc(_bgd, _gf)
	case 8:
		_bag = _ga(_bgd, _gf)
	default:
		return _g.Error(_fe, "\u0065\u0078p\u0061\u006e\u0073\u0069o\u006e\u0020f\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020r\u0061\u006e\u0067\u0065")
	}
	if _bag != nil {
		_bag = _g.Wrap(_bag, _fe, "")
	}
	return _bag
}
func ClipBoxToRectangle(box *_aa.Rectangle, wi, hi int) (_geb *_aa.Rectangle, _aeda error) {
	const _efg = "\u0043l\u0069p\u0042\u006f\u0078\u0054\u006fR\u0065\u0063t\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, _g.Error(_efg, "\u0027\u0062\u006f\u0078\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if box.Min.X >= wi || box.Min.Y >= hi || box.Max.X <= 0 || box.Max.Y <= 0 {
		return nil, _g.Error(_efg, "\u0027\u0062\u006fx'\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065")
	}
	_cgae := *box
	_geb = &_cgae
	if _geb.Min.X < 0 {
		_geb.Max.X += _geb.Min.X
		_geb.Min.X = 0
	}
	if _geb.Min.Y < 0 {
		_geb.Max.Y += _geb.Min.Y
		_geb.Min.Y = 0
	}
	if _geb.Max.X > wi {
		_geb.Max.X = wi
	}
	if _geb.Max.Y > hi {
		_geb.Max.Y = hi
	}
	return _geb, nil
}
func (_ebcce *Bitmaps) selectByIndexes(_feead []int) (*Bitmaps, error) {
	_gbbcde := &Bitmaps{}
	for _, _defe := range _feead {
		_gfcbe, _adfb := _ebcce.GetBitmap(_defe)
		if _adfb != nil {
			return nil, _g.Wrap(_adfb, "\u0073e\u006ce\u0063\u0074\u0042\u0079\u0049\u006e\u0064\u0065\u0078\u0065\u0073", "")
		}
		_gbbcde.AddBitmap(_gfcbe)
	}
	return _gbbcde, nil
}

const (
	_eafd shift = iota
	_fbba
)

func MakePixelCentroidTab8() []int             { return _ceafb() }
func (_aaa *Bitmap) Equivalent(s *Bitmap) bool { return _aaa.equivalent(s) }
func TstESymbol(t *_ba.T, scale ...int) *Bitmap {
	_dffge, _daeg := NewWithData(4, 5, []byte{0xF0, 0x80, 0xE0, 0x80, 0xF0})
	_b.NoError(t, _daeg)
	return TstGetScaledSymbol(t, _dffge, scale...)
}
func (_gbf *Bitmap) ToImage() _aa.Image {
	_dafd, _dgee := _ac.NewImage(_gbf.Width, _gbf.Height, 1, 1, _gbf.Data, nil, nil)
	if _dgee != nil {
		_gb.Log.Error("\u0043\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020j\u0062\u0069\u0067\u0032\u002e\u0042\u0069\u0074m\u0061p\u0020\u0074\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002e\u0049\u006d\u0061\u0067e\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _dgee)
	}
	return _dafd
}
func (_bbd *Bitmap) setEightFullBytes(_eaee int, _fcee uint64) error {
	if _eaee+7 > len(_bbd.Data)-1 {
		return _g.Error("\u0073\u0065\u0074\u0045\u0069\u0067\u0068\u0074\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bbd.Data[_eaee] = byte((_fcee & 0xff00000000000000) >> 56)
	_bbd.Data[_eaee+1] = byte((_fcee & 0xff000000000000) >> 48)
	_bbd.Data[_eaee+2] = byte((_fcee & 0xff0000000000) >> 40)
	_bbd.Data[_eaee+3] = byte((_fcee & 0xff00000000) >> 32)
	_bbd.Data[_eaee+4] = byte((_fcee & 0xff000000) >> 24)
	_bbd.Data[_eaee+5] = byte((_fcee & 0xff0000) >> 16)
	_bbd.Data[_eaee+6] = byte((_fcee & 0xff00) >> 8)
	_bbd.Data[_eaee+7] = byte(_fcee & 0xff)
	return nil
}
func (_acdcf *BitmapsArray) AddBitmaps(bm *Bitmaps) { _acdcf.Values = append(_acdcf.Values, bm) }
func TstWSymbol(t *_ba.T, scale ...int) *Bitmap {
	_dffd, _beafa := NewWithData(5, 5, []byte{0x88, 0x88, 0xA8, 0xD8, 0x88})
	_b.NoError(t, _beafa)
	return TstGetScaledSymbol(t, _dffd, scale...)
}

type Bitmaps struct {
	Values []*Bitmap
	Boxes  []*_aa.Rectangle
}
type byWidth Bitmaps

var _eggbe = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x78, 0x27, 0xC2, 0x27, 0x91, 0x00, 0x22, 0x48, 0x21, 0x03, 0x24, 0x91, 0x00, 0x22, 0x48, 0x21, 0x02, 0xA4, 0x95, 0x00, 0x22, 0x48, 0x21, 0x02, 0x64, 0x9B, 0x00, 0x3C, 0x78, 0x21, 0x02, 0x27, 0x91, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func (_ecbd *Bitmap) clipRectangle(_fdae, _gcbbc *_aa.Rectangle) (_bbfe *Bitmap, _fbga error) {
	const _aad = "\u0063\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if _fdae == nil {
		return nil, _g.Error(_aad, "\u0070r\u006fv\u0069\u0064\u0065\u0064\u0020n\u0069\u006c \u0027\u0062\u006f\u0078\u0027")
	}
	_beeb, _bffd := _ecbd.Width, _ecbd.Height
	_cecb, _fbga := ClipBoxToRectangle(_fdae, _beeb, _bffd)
	if _fbga != nil {
		_gb.Log.Warning("\u0027\u0062ox\u0027\u0020\u0064o\u0065\u0073\u006e\u0027t o\u0076er\u006c\u0061\u0070\u0020\u0062\u0069\u0074ma\u0070\u0020\u0027\u0062\u0027\u003a\u0020%\u0076", _fbga)
		return nil, nil
	}
	_dadd, _cda := _cecb.Min.X, _cecb.Min.Y
	_cgfe, _fdgb := _cecb.Max.X-_cecb.Min.X, _cecb.Max.Y-_cecb.Min.Y
	_bbfe = New(_cgfe, _fdgb)
	_bbfe.Text = _ecbd.Text
	if _fbga = _bbfe.RasterOperation(0, 0, _cgfe, _fdgb, PixSrc, _ecbd, _dadd, _cda); _fbga != nil {
		return nil, _g.Wrap(_fbga, _aad, "")
	}
	if _gcbbc != nil {
		*_gcbbc = *_cecb
	}
	return _bbfe, nil
}
func (_edg *ClassedPoints) xSortFunction() func(_agffg int, _bffcb int) bool {
	return func(_afde, _aecg int) bool { return _edg.XAtIndex(_afde) < _edg.XAtIndex(_aecg) }
}

type Selection struct {
	Height, Width int
	Cx, Cy        int
	Name          string
	Data          [][]SelectionValue
}

func (_bbc *Bitmap) ClipRectangle(box *_aa.Rectangle) (_cce *Bitmap, _def *_aa.Rectangle, _ggc error) {
	const _dbcb = "\u0043\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, nil, _g.Error(_dbcb, "\u0062o\u0078 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	_bcd, _cdg := _bbc.Width, _bbc.Height
	_fdg := _aa.Rect(0, 0, _bcd, _cdg)
	if !box.Overlaps(_fdg) {
		return nil, nil, _g.Error(_dbcb, "b\u006f\u0078\u0020\u0064oe\u0073n\u0027\u0074\u0020\u006f\u0076e\u0072\u006c\u0061\u0070\u0020\u0062")
	}
	_gbc := box.Intersect(_fdg)
	_aeca, _bae := _gbc.Min.X, _gbc.Min.Y
	_bfd, _gda := _gbc.Dx(), _gbc.Dy()
	_cce = New(_bfd, _gda)
	_cce.Text = _bbc.Text
	if _ggc = _cce.RasterOperation(0, 0, _bfd, _gda, PixSrc, _bbc, _aeca, _bae); _ggc != nil {
		return nil, nil, _g.Wrap(_ggc, _dbcb, "\u0050\u0069\u0078\u0053\u0072\u0063\u0020\u0074\u006f\u0020\u0063\u006ci\u0070\u0070\u0065\u0064")
	}
	_def = &_gbc
	return _cce, _def, nil
}

type Points []Point

func RasterOperation(dest *Bitmap, dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _aee(dest, dx, dy, dw, dh, op, src, sx, sy)
}
func CorrelationScore(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_efee float64, _adgb error) {
	const _acbf = "\u0063\u006fr\u0072\u0065\u006ca\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065"
	if bm1 == nil || bm2 == nil {
		return 0, _g.Error(_acbf, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0062\u0069\u0074ma\u0070\u0073")
	}
	if tab == nil {
		return 0, _g.Error(_acbf, "\u0027\u0074\u0061\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if area1 <= 0 || area2 <= 0 {
		return 0, _g.Error(_acbf, "\u0061\u0072\u0065\u0061s\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067r\u0065a\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
	}
	_ddg, _efab := bm1.Width, bm1.Height
	_aagg, _beb := bm2.Width, bm2.Height
	_ebda := _bfad(_ddg - _aagg)
	if _ebda > maxDiffW {
		return 0, nil
	}
	_ccce := _bfad(_efab - _beb)
	if _ccce > maxDiffH {
		return 0, nil
	}
	var _fdcb, _facdc int
	if delX >= 0 {
		_fdcb = int(delX + 0.5)
	} else {
		_fdcb = int(delX - 0.5)
	}
	if delY >= 0 {
		_facdc = int(delY + 0.5)
	} else {
		_facdc = int(delY - 0.5)
	}
	_fgce := _aage(_facdc, 0)
	_ccbb := _efag(_beb+_facdc, _efab)
	_cebf := bm1.RowStride * _fgce
	_bdab := bm2.RowStride * (_fgce - _facdc)
	_ccaae := _aage(_fdcb, 0)
	_bgegg := _efag(_aagg+_fdcb, _ddg)
	_fafb := bm2.RowStride
	var _aedb, _eedf int
	if _fdcb >= 8 {
		_aedb = _fdcb >> 3
		_cebf += _aedb
		_ccaae -= _aedb << 3
		_bgegg -= _aedb << 3
		_fdcb &= 7
	} else if _fdcb <= -8 {
		_eedf = -((_fdcb + 7) >> 3)
		_bdab += _eedf
		_fafb -= _eedf
		_fdcb += _eedf << 3
	}
	if _ccaae >= _bgegg || _fgce >= _ccbb {
		return 0, nil
	}
	_dcab := (_bgegg + 7) >> 3
	var (
		_fbad, _ebaf, _ecdf byte
		_ddfb, _acec, _abf  int
	)
	switch {
	case _fdcb == 0:
		for _abf = _fgce; _abf < _ccbb; _abf, _cebf, _bdab = _abf+1, _cebf+bm1.RowStride, _bdab+bm2.RowStride {
			for _acec = 0; _acec < _dcab; _acec++ {
				_ecdf = bm1.Data[_cebf+_acec] & bm2.Data[_bdab+_acec]
				_ddfb += tab[_ecdf]
			}
		}
	case _fdcb > 0:
		if _fafb < _dcab {
			for _abf = _fgce; _abf < _ccbb; _abf, _cebf, _bdab = _abf+1, _cebf+bm1.RowStride, _bdab+bm2.RowStride {
				_fbad, _ebaf = bm1.Data[_cebf], bm2.Data[_bdab]>>uint(_fdcb)
				_ecdf = _fbad & _ebaf
				_ddfb += tab[_ecdf]
				for _acec = 1; _acec < _fafb; _acec++ {
					_fbad, _ebaf = bm1.Data[_cebf+_acec], (bm2.Data[_bdab+_acec]>>uint(_fdcb))|(bm2.Data[_bdab+_acec-1]<<uint(8-_fdcb))
					_ecdf = _fbad & _ebaf
					_ddfb += tab[_ecdf]
				}
				_fbad = bm1.Data[_cebf+_acec]
				_ebaf = bm2.Data[_bdab+_acec-1] << uint(8-_fdcb)
				_ecdf = _fbad & _ebaf
				_ddfb += tab[_ecdf]
			}
		} else {
			for _abf = _fgce; _abf < _ccbb; _abf, _cebf, _bdab = _abf+1, _cebf+bm1.RowStride, _bdab+bm2.RowStride {
				_fbad, _ebaf = bm1.Data[_cebf], bm2.Data[_bdab]>>uint(_fdcb)
				_ecdf = _fbad & _ebaf
				_ddfb += tab[_ecdf]
				for _acec = 1; _acec < _dcab; _acec++ {
					_fbad = bm1.Data[_cebf+_acec]
					_ebaf = (bm2.Data[_bdab+_acec] >> uint(_fdcb)) | (bm2.Data[_bdab+_acec-1] << uint(8-_fdcb))
					_ecdf = _fbad & _ebaf
					_ddfb += tab[_ecdf]
				}
			}
		}
	default:
		if _dcab < _fafb {
			for _abf = _fgce; _abf < _ccbb; _abf, _cebf, _bdab = _abf+1, _cebf+bm1.RowStride, _bdab+bm2.RowStride {
				for _acec = 0; _acec < _dcab; _acec++ {
					_fbad = bm1.Data[_cebf+_acec]
					_ebaf = bm2.Data[_bdab+_acec] << uint(-_fdcb)
					_ebaf |= bm2.Data[_bdab+_acec+1] >> uint(8+_fdcb)
					_ecdf = _fbad & _ebaf
					_ddfb += tab[_ecdf]
				}
			}
		} else {
			for _abf = _fgce; _abf < _ccbb; _abf, _cebf, _bdab = _abf+1, _cebf+bm1.RowStride, _bdab+bm2.RowStride {
				for _acec = 0; _acec < _dcab-1; _acec++ {
					_fbad = bm1.Data[_cebf+_acec]
					_ebaf = bm2.Data[_bdab+_acec] << uint(-_fdcb)
					_ebaf |= bm2.Data[_bdab+_acec+1] >> uint(8+_fdcb)
					_ecdf = _fbad & _ebaf
					_ddfb += tab[_ecdf]
				}
				_fbad = bm1.Data[_cebf+_acec]
				_ebaf = bm2.Data[_bdab+_acec] << uint(-_fdcb)
				_ecdf = _fbad & _ebaf
				_ddfb += tab[_ecdf]
			}
		}
	}
	_efee = float64(_ddfb) * float64(_ddfb) / (float64(area1) * float64(area2))
	return _efee, nil
}
func _faaf(_edff *Bitmap, _abfd *Bitmap, _debc *Selection, _ggaeg **Bitmap) (*Bitmap, error) {
	const _bfcd = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0031"
	if _abfd == nil {
		return nil, _g.Error(_bfcd, "\u004d\u006f\u0072\u0070\u0068\u0041\u0072\u0067\u0073\u0031\u0020'\u0073\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066i\u006e\u0065\u0064")
	}
	if _debc == nil {
		return nil, _g.Error(_bfcd, "\u004d\u006f\u0072\u0068p\u0041\u0072\u0067\u0073\u0031\u0020\u0027\u0073\u0065\u006c'\u0020n\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064")
	}
	_cggg, _cagb := _debc.Height, _debc.Width
	if _cggg == 0 || _cagb == 0 {
		return nil, _g.Error(_bfcd, "\u0073\u0065\u006c\u0065ct\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _edff == nil {
		_edff = _abfd.createTemplate()
		*_ggaeg = _abfd
		return _edff, nil
	}
	_edff.Width = _abfd.Width
	_edff.Height = _abfd.Height
	_edff.RowStride = _abfd.RowStride
	_edff.Color = _abfd.Color
	_edff.Data = make([]byte, _abfd.RowStride*_abfd.Height)
	if _edff == _abfd {
		*_ggaeg = _abfd.Copy()
	} else {
		*_ggaeg = _abfd
	}
	return _edff, nil
}
func _ggdf(_eadbd, _bddc *Bitmap, _begd *Selection) (*Bitmap, error) {
	const _ebcc = "\u006f\u0070\u0065\u006e"
	var _gffd error
	_eadbd, _gffd = _cgccb(_eadbd, _bddc, _begd)
	if _gffd != nil {
		return nil, _g.Wrap(_gffd, _ebcc, "")
	}
	_fabc, _gffd := _babe(nil, _bddc, _begd)
	if _gffd != nil {
		return nil, _g.Wrap(_gffd, _ebcc, "")
	}
	_, _gffd = _ceed(_eadbd, _fabc, _begd)
	if _gffd != nil {
		return nil, _g.Wrap(_gffd, _ebcc, "")
	}
	return _eadbd, nil
}
func TstTSymbol(t *_ba.T, scale ...int) *Bitmap {
	_bfbc, _cefa := NewWithData(5, 5, []byte{0xF8, 0x20, 0x20, 0x20, 0x20})
	_b.NoError(t, _cefa)
	return TstGetScaledSymbol(t, _bfbc, scale...)
}
func (_bdceg Points) GetGeometry(i int) (_adag, _dea float32, _ffdegd error) {
	if i > len(_bdceg)-1 {
		return 0, 0, _g.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_ccfa := _bdceg[i]
	return _ccfa.X, _ccfa.Y, nil
}
func (_bcba *Bitmap) equivalent(_ccc *Bitmap) bool {
	if _bcba == _ccc {
		return true
	}
	if !_bcba.SizesEqual(_ccc) {
		return false
	}
	_bbg := _acad(_bcba, _ccc, CmbOpXor)
	_cab := _bcba.countPixels()
	_faee := int(0.25 * float32(_cab))
	if _bbg.thresholdPixelSum(_faee) {
		return false
	}
	var (
		_bdfeb [9][9]int
		_adg   [18][9]int
		_ceaf  [9][18]int
		_fdb   int
		_ega   int
	)
	_aag := 9
	_fec := _bcba.Height / _aag
	_fbca := _bcba.Width / _aag
	_ade, _cdcg := _fec/2, _fbca/2
	if _fec < _fbca {
		_ade = _fbca / 2
		_cdcg = _fec / 2
	}
	_cdce := float64(_ade) * float64(_cdcg) * _ca.Pi
	_fac := int(float64(_fec*_fbca/2) * 0.9)
	_bbge := int(float64(_fbca*_fec/2) * 0.9)
	for _gag := 0; _gag < _aag; _gag++ {
		_acbg := _fbca*_gag + _fdb
		var _cfa int
		if _gag == _aag-1 {
			_fdb = 0
			_cfa = _bcba.Width
		} else {
			_cfa = _acbg + _fbca
			if ((_bcba.Width - _fdb) % _aag) > 0 {
				_fdb++
				_cfa++
			}
		}
		for _fba := 0; _fba < _aag; _fba++ {
			_abgd := _fec*_fba + _ega
			var _geac int
			if _fba == _aag-1 {
				_ega = 0
				_geac = _bcba.Height
			} else {
				_geac = _abgd + _fec
				if (_bcba.Height-_ega)%_aag > 0 {
					_ega++
					_geac++
				}
			}
			var _aceg, _acab, _bcbe, _ege int
			_gfbeb := (_acbg + _cfa) / 2
			_eab := (_abgd + _geac) / 2
			for _fbd := _acbg; _fbd < _cfa; _fbd++ {
				for _fffd := _abgd; _fffd < _geac; _fffd++ {
					if _bbg.GetPixel(_fbd, _fffd) {
						if _fbd < _gfbeb {
							_aceg++
						} else {
							_acab++
						}
						if _fffd < _eab {
							_ege++
						} else {
							_bcbe++
						}
					}
				}
			}
			_bdfeb[_gag][_fba] = _aceg + _acab
			_adg[_gag*2][_fba] = _aceg
			_adg[_gag*2+1][_fba] = _acab
			_ceaf[_gag][_fba*2] = _ege
			_ceaf[_gag][_fba*2+1] = _bcbe
		}
	}
	for _cgg := 0; _cgg < _aag*2-1; _cgg++ {
		for _fcag := 0; _fcag < (_aag - 1); _fcag++ {
			var _cgd int
			for _fgfd := 0; _fgfd < 2; _fgfd++ {
				for _dcb := 0; _dcb < 2; _dcb++ {
					_cgd += _adg[_cgg+_fgfd][_fcag+_dcb]
				}
			}
			if _cgd > _bbge {
				return false
			}
		}
	}
	for _abd := 0; _abd < (_aag - 1); _abd++ {
		for _aebd := 0; _aebd < ((_aag * 2) - 1); _aebd++ {
			var _accf int
			for _ffdeg := 0; _ffdeg < 2; _ffdeg++ {
				for _feec := 0; _feec < 2; _feec++ {
					_accf += _ceaf[_abd+_ffdeg][_aebd+_feec]
				}
			}
			if _accf > _fac {
				return false
			}
		}
	}
	for _dbda := 0; _dbda < (_aag - 2); _dbda++ {
		for _daa := 0; _daa < (_aag - 2); _daa++ {
			var _dgb, _cfe int
			for _gadc := 0; _gadc < 3; _gadc++ {
				for _gdab := 0; _gdab < 3; _gdab++ {
					if _gadc == _gdab {
						_dgb += _bdfeb[_dbda+_gadc][_daa+_gdab]
					}
					if (2 - _gadc) == _gdab {
						_cfe += _bdfeb[_dbda+_gadc][_daa+_gdab]
					}
				}
			}
			if _dgb > _bbge || _cfe > _bbge {
				return false
			}
		}
	}
	for _fdab := 0; _fdab < (_aag - 1); _fdab++ {
		for _ffc := 0; _ffc < (_aag - 1); _ffc++ {
			var _bfdb int
			for _gfdd := 0; _gfdd < 2; _gfdd++ {
				for _cedc := 0; _cedc < 2; _cedc++ {
					_bfdb += _bdfeb[_fdab+_gfdd][_ffc+_cedc]
				}
			}
			if float64(_bfdb) > _cdce {
				return false
			}
		}
	}
	return true
}
func TstPSymbol(t *_ba.T) *Bitmap {
	t.Helper()
	_cgeg := New(5, 8)
	_b.NoError(t, _cgeg.SetPixel(0, 0, 1))
	_b.NoError(t, _cgeg.SetPixel(1, 0, 1))
	_b.NoError(t, _cgeg.SetPixel(2, 0, 1))
	_b.NoError(t, _cgeg.SetPixel(3, 0, 1))
	_b.NoError(t, _cgeg.SetPixel(4, 1, 1))
	_b.NoError(t, _cgeg.SetPixel(0, 1, 1))
	_b.NoError(t, _cgeg.SetPixel(4, 2, 1))
	_b.NoError(t, _cgeg.SetPixel(0, 2, 1))
	_b.NoError(t, _cgeg.SetPixel(4, 3, 1))
	_b.NoError(t, _cgeg.SetPixel(0, 3, 1))
	_b.NoError(t, _cgeg.SetPixel(0, 4, 1))
	_b.NoError(t, _cgeg.SetPixel(1, 4, 1))
	_b.NoError(t, _cgeg.SetPixel(2, 4, 1))
	_b.NoError(t, _cgeg.SetPixel(3, 4, 1))
	_b.NoError(t, _cgeg.SetPixel(0, 5, 1))
	_b.NoError(t, _cgeg.SetPixel(0, 6, 1))
	_b.NoError(t, _cgeg.SetPixel(0, 7, 1))
	return _cgeg
}
func (_geg *Bitmap) SetByte(index int, v byte) error {
	if index > len(_geg.Data)-1 || index < 0 {
		return _g.Errorf("\u0053e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", index)
	}
	_geg.Data[index] = v
	return nil
}
func _eca() (_gcf [256]uint16) {
	for _ecf := 0; _ecf < 256; _ecf++ {
		if _ecf&0x01 != 0 {
			_gcf[_ecf] |= 0x3
		}
		if _ecf&0x02 != 0 {
			_gcf[_ecf] |= 0xc
		}
		if _ecf&0x04 != 0 {
			_gcf[_ecf] |= 0x30
		}
		if _ecf&0x08 != 0 {
			_gcf[_ecf] |= 0xc0
		}
		if _ecf&0x10 != 0 {
			_gcf[_ecf] |= 0x300
		}
		if _ecf&0x20 != 0 {
			_gcf[_ecf] |= 0xc00
		}
		if _ecf&0x40 != 0 {
			_gcf[_ecf] |= 0x3000
		}
		if _ecf&0x80 != 0 {
			_gcf[_ecf] |= 0xc000
		}
	}
	return _gcf
}
func _ceafb() []int {
	_dabbg := make([]int, 256)
	_dabbg[0] = 0
	_dabbg[1] = 7
	var _bdaf int
	for _bdaf = 2; _bdaf < 4; _bdaf++ {
		_dabbg[_bdaf] = _dabbg[_bdaf-2] + 6
	}
	for _bdaf = 4; _bdaf < 8; _bdaf++ {
		_dabbg[_bdaf] = _dabbg[_bdaf-4] + 5
	}
	for _bdaf = 8; _bdaf < 16; _bdaf++ {
		_dabbg[_bdaf] = _dabbg[_bdaf-8] + 4
	}
	for _bdaf = 16; _bdaf < 32; _bdaf++ {
		_dabbg[_bdaf] = _dabbg[_bdaf-16] + 3
	}
	for _bdaf = 32; _bdaf < 64; _bdaf++ {
		_dabbg[_bdaf] = _dabbg[_bdaf-32] + 2
	}
	for _bdaf = 64; _bdaf < 128; _bdaf++ {
		_dabbg[_bdaf] = _dabbg[_bdaf-64] + 1
	}
	for _bdaf = 128; _bdaf < 256; _bdaf++ {
		_dabbg[_bdaf] = _dabbg[_bdaf-128]
	}
	return _dabbg
}
func Centroid(bm *Bitmap, centTab, sumTab []int) (Point, error) { return bm.centroid(centTab, sumTab) }

var _ _a.Interface = &ClassedPoints{}

type BoundaryCondition int

func (_gfgd *byHeight) Swap(i, j int) {
	_gfgd.Values[i], _gfgd.Values[j] = _gfgd.Values[j], _gfgd.Values[i]
	if _gfgd.Boxes != nil {
		_gfgd.Boxes[i], _gfgd.Boxes[j] = _gfgd.Boxes[j], _gfgd.Boxes[i]
	}
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

func TstNSymbol(t *_ba.T, scale ...int) *Bitmap {
	_ecbb, _daacc := NewWithData(4, 5, []byte{0x90, 0xD0, 0xB0, 0x90, 0x90})
	_b.NoError(t, _daacc)
	return TstGetScaledSymbol(t, _ecbb, scale...)
}
func (_bdcc *ClassedPoints) GetIntXByClass(i int) (int, error) {
	const _eadaf = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _bdcc.IntSlice.Size() {
		return 0, _g.Errorf(_eadaf, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_bdcc.XAtIndex(i)), nil
}
func _ceba(_efgg, _agff *Bitmap, _agfd, _ecec int) (*Bitmap, error) {
	const _gfce = "\u0065\u0072\u006f\u0064\u0065\u0042\u0072\u0069\u0063\u006b"
	if _agff == nil {
		return nil, _g.Error(_gfce, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _agfd < 1 || _ecec < 1 {
		return nil, _g.Error(_gfce, "\u0068\u0073\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069\u007a\u0065\u0020\u0061\u0072e\u0020\u006e\u006f\u0074\u0020\u0067\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006fr\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _agfd == 1 && _ecec == 1 {
		_fdef, _caca := _feea(_efgg, _agff)
		if _caca != nil {
			return nil, _g.Wrap(_caca, _gfce, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _fdef, nil
	}
	if _agfd == 1 || _ecec == 1 {
		_egfd := SelCreateBrick(_ecec, _agfd, _ecec/2, _agfd/2, SelHit)
		_gegb, _fddf := _babe(_efgg, _agff, _egfd)
		if _fddf != nil {
			return nil, _g.Wrap(_fddf, _gfce, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _gegb, nil
	}
	_addg := SelCreateBrick(1, _agfd, 0, _agfd/2, SelHit)
	_gdfa := SelCreateBrick(_ecec, 1, _ecec/2, 0, SelHit)
	_fgfb, _aceb := _babe(nil, _agff, _addg)
	if _aceb != nil {
		return nil, _g.Wrap(_aceb, _gfce, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_efgg, _aceb = _babe(_efgg, _fgfb, _gdfa)
	if _aceb != nil {
		return nil, _g.Wrap(_aceb, _gfce, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _efgg, nil
}
func Dilate(d *Bitmap, s *Bitmap, sel *Selection) (*Bitmap, error) { return _ceed(d, s, sel) }
func RankHausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH, area1, area3 int, rank float32, tab8 []int) (_eeeg bool, _abbd error) {
	const _cbccd = "\u0052\u0061\u006ek\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_cfed, _fafe := p1.Width, p1.Height
	_fdee, _aege := p3.Width, p3.Height
	if _dd.Abs(_cfed-_fdee) > maxDiffW {
		return false, nil
	}
	if _dd.Abs(_fafe-_aege) > maxDiffH {
		return false, nil
	}
	_efda := int(float32(area1)*(1.0-rank) + 0.5)
	_cba := int(float32(area3)*(1.0-rank) + 0.5)
	var _dggff, _ggcbc int
	if delX >= 0 {
		_dggff = int(delX + 0.5)
	} else {
		_dggff = int(delX - 0.5)
	}
	if delY >= 0 {
		_ggcbc = int(delY + 0.5)
	} else {
		_ggcbc = int(delY - 0.5)
	}
	_ebdc := p1.CreateTemplate()
	if _abbd = _ebdc.RasterOperation(0, 0, _cfed, _fafe, PixSrc, p1, 0, 0); _abbd != nil {
		return false, _g.Wrap(_abbd, _cbccd, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _abbd = _ebdc.RasterOperation(_dggff, _ggcbc, _cfed, _fafe, PixNotSrcAndDst, p4, 0, 0); _abbd != nil {
		return false, _g.Wrap(_abbd, _cbccd, "\u0074 \u0026\u0020\u0021\u0070\u0034")
	}
	_eeeg, _abbd = _ebdc.ThresholdPixelSum(_efda, tab8)
	if _abbd != nil {
		return false, _g.Wrap(_abbd, _cbccd, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0031")
	}
	if _eeeg {
		return false, nil
	}
	if _abbd = _ebdc.RasterOperation(_dggff, _ggcbc, _fdee, _aege, PixSrc, p3, 0, 0); _abbd != nil {
		return false, _g.Wrap(_abbd, _cbccd, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _abbd = _ebdc.RasterOperation(0, 0, _fdee, _aege, PixNotSrcAndDst, p2, 0, 0); _abbd != nil {
		return false, _g.Wrap(_abbd, _cbccd, "\u0074 \u0026\u0020\u0021\u0070\u0032")
	}
	_eeeg, _abbd = _ebdc.ThresholdPixelSum(_cba, tab8)
	if _abbd != nil {
		return false, _g.Wrap(_abbd, _cbccd, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0033")
	}
	return !_eeeg, nil
}
func init() {
	const _cgbd = "\u0062\u0069\u0074\u006dap\u0073\u002e\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0061\u0074\u0069o\u006e"
	_gbef = New(50, 40)
	var _cbdb error
	_gbef, _cbdb = _gbef.AddBorder(2, 1)
	if _cbdb != nil {
		panic(_g.Wrap(_cbdb, _cgbd, "f\u0072\u0061\u006d\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
	_fbbbf, _cbdb = NewWithData(50, 22, _eggbe)
	if _cbdb != nil {
		panic(_g.Wrap(_cbdb, _cgbd, "i\u006d\u0061\u0067\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
}
func (_cgdg MorphProcess) getWidthHeight() (_cccb, _acga int) {
	return _cgdg.Arguments[0], _cgdg.Arguments[1]
}
func _cbdc(_fcbe, _bcgeb *Bitmap, _gcfg, _ggdd int) (_dfafb error) {
	const _dbef = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0038"
	var (
		_ecea, _fadb, _edfba, _gebb                              int
		_ccdd, _cbgb, _fcca, _beee, _bcggd, _bddf, _bbcb, _dabgg byte
	)
	for _ecea = 0; _ecea < _gcfg; _ecea++ {
		_edfba = _ecea * _fcbe.RowStride
		_gebb = _ecea * _bcgeb.RowStride
		for _fadb = 0; _fadb < _ggdd; _fadb++ {
			if _ccdd, _dfafb = _fcbe.GetByte(_edfba + _fadb); _dfafb != nil {
				return _g.Wrap(_dfafb, _dbef, "\u0067e\u0074 \u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _cbgb, _dfafb = _bcgeb.GetByte(_gebb + _fadb); _dfafb != nil {
				return _g.Wrap(_dfafb, _dbef, "\u0067\u0065\u0074\u0020\u006d\u0061\u0073\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _ecea > 0 {
				if _fcca, _dfafb = _fcbe.GetByte(_edfba - _fcbe.RowStride + _fadb); _dfafb != nil {
					return _g.Wrap(_dfafb, _dbef, "\u0069\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
				}
				_ccdd |= _fcca | (_fcca << 1) | (_fcca >> 1)
				if _fadb > 0 {
					if _dabgg, _dfafb = _fcbe.GetByte(_edfba - _fcbe.RowStride + _fadb - 1); _dfafb != nil {
						return _g.Wrap(_dfafb, _dbef, "\u0069\u0020\u003e\u00200 \u0026\u0026\u0020\u006a\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
					}
					_ccdd |= _dabgg << 7
				}
				if _fadb < _ggdd-1 {
					if _dabgg, _dfafb = _fcbe.GetByte(_edfba - _fcbe.RowStride + _fadb + 1); _dfafb != nil {
						return _g.Wrap(_dfafb, _dbef, "\u006a\u0020<\u0020\u0077\u0070l\u0020\u002d\u0020\u0031\u0020\u0062\u0079\u0074\u0065")
					}
					_ccdd |= _dabgg >> 7
				}
			}
			if _fadb > 0 {
				if _beee, _dfafb = _fcbe.GetByte(_edfba + _fadb - 1); _dfafb != nil {
					return _g.Wrap(_dfafb, _dbef, "\u006a\u0020\u003e \u0030")
				}
				_ccdd |= _beee << 7
			}
			_ccdd &= _cbgb
			if _ccdd == 0 || ^_ccdd == 0 {
				if _dfafb = _fcbe.SetByte(_edfba+_fadb, _ccdd); _dfafb != nil {
					return _g.Wrap(_dfafb, _dbef, "\u0073e\u0074t\u0069\u006e\u0067\u0020\u0065m\u0070\u0074y\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_bbcb = _ccdd
				_ccdd = (_ccdd | (_ccdd >> 1) | (_ccdd << 1)) & _cbgb
				if (_ccdd ^ _bbcb) == 0 {
					if _dfafb = _fcbe.SetByte(_edfba+_fadb, _ccdd); _dfafb != nil {
						return _g.Wrap(_dfafb, _dbef, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _ecea = _gcfg - 1; _ecea >= 0; _ecea-- {
		_edfba = _ecea * _fcbe.RowStride
		_gebb = _ecea * _bcgeb.RowStride
		for _fadb = _ggdd - 1; _fadb >= 0; _fadb-- {
			if _ccdd, _dfafb = _fcbe.GetByte(_edfba + _fadb); _dfafb != nil {
				return _g.Wrap(_dfafb, _dbef, "\u0072\u0065\u0076er\u0073\u0065\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _cbgb, _dfafb = _bcgeb.GetByte(_gebb + _fadb); _dfafb != nil {
				return _g.Wrap(_dfafb, _dbef, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _ecea < _gcfg-1 {
				if _bcggd, _dfafb = _fcbe.GetByte(_edfba + _fcbe.RowStride + _fadb); _dfafb != nil {
					return _g.Wrap(_dfafb, _dbef, "\u0069\u0020\u003c\u0020h\u0020\u002d\u0020\u0031\u0020\u002d\u003e\u0020\u0067\u0065t\u0020s\u006f\u0075\u0072\u0063\u0065\u0020\u0062y\u0074\u0065")
				}
				_ccdd |= _bcggd | (_bcggd << 1) | _bcggd>>1
				if _fadb > 0 {
					if _dabgg, _dfafb = _fcbe.GetByte(_edfba + _fcbe.RowStride + _fadb - 1); _dfafb != nil {
						return _g.Wrap(_dfafb, _dbef, "\u0069\u0020\u003c h\u002d\u0031\u0020\u0026\u0020\u006a\u0020\u003e\u00200\u0020-\u003e \u0067e\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
					}
					_ccdd |= _dabgg << 7
				}
				if _fadb < _ggdd-1 {
					if _dabgg, _dfafb = _fcbe.GetByte(_edfba + _fcbe.RowStride + _fadb + 1); _dfafb != nil {
						return _g.Wrap(_dfafb, _dbef, "\u0069\u0020\u003c\u0020\u0068\u002d\u0031\u0020\u0026\u0026\u0020\u006a\u0020\u003c\u0077\u0070\u006c\u002d\u0031\u0020\u002d\u003e\u0020\u0067e\u0074\u0020\u0073\u006f\u0075r\u0063\u0065 \u0062\u0079\u0074\u0065")
					}
					_ccdd |= _dabgg >> 7
				}
			}
			if _fadb < _ggdd-1 {
				if _bddf, _dfafb = _fcbe.GetByte(_edfba + _fadb + 1); _dfafb != nil {
					return _g.Wrap(_dfafb, _dbef, "\u006a\u0020<\u0020\u0077\u0070\u006c\u0020\u002d\u0031\u0020\u002d\u003e\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020by\u0074\u0065")
				}
				_ccdd |= _bddf >> 7
			}
			_ccdd &= _cbgb
			if _ccdd == 0 || (^_ccdd) == 0 {
				if _dfafb = _fcbe.SetByte(_edfba+_fadb, _ccdd); _dfafb != nil {
					return _g.Wrap(_dfafb, _dbef, "\u0073e\u0074 \u006d\u0061\u0073\u006b\u0065\u0064\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_bbcb = _ccdd
				_ccdd = (_ccdd | (_ccdd >> 1) | (_ccdd << 1)) & _cbgb
				if (_ccdd ^ _bbcb) == 0 {
					if _dfafb = _fcbe.SetByte(_edfba+_fadb, _ccdd); _dfafb != nil {
						return _g.Wrap(_dfafb, _dbef, "r\u0065\u0076\u0065\u0072se\u0020s\u0065\u0074\u0020\u0070\u0072e\u0076\u0020\u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	return nil
}
func MorphSequence(src *Bitmap, sequence ...MorphProcess) (*Bitmap, error) {
	return _bbcc(src, sequence...)
}
func _agabf(_dfca ...MorphProcess) (_faed error) {
	const _cgfb = "v\u0065r\u0069\u0066\u0079\u004d\u006f\u0072\u0070\u0068P\u0072\u006f\u0063\u0065ss\u0065\u0073"
	var _ddbe, _bcf int
	for _cadc, _fcga := range _dfca {
		if _faed = _fcga.verify(_cadc, &_ddbe, &_bcf); _faed != nil {
			return _g.Wrap(_faed, _cgfb, "")
		}
	}
	if _bcf != 0 && _ddbe != 0 {
		return _g.Error(_cgfb, "\u004d\u006f\u0072\u0070\u0068\u0020\u0073\u0065\u0071\u0075\u0065n\u0063\u0065\u0020\u002d\u0020\u0062\u006f\u0072d\u0065r\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u0065\u0074\u0020\u0072\u0065\u0064u\u0063\u0074\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0030")
	}
	return nil
}
func TstFrameBitmap() *Bitmap { return _gbef.Copy() }
func (_bfbaa *BitmapsArray) GetBox(i int) (*_aa.Rectangle, error) {
	const _gdcg = "\u0042\u0069\u0074\u006dap\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u006f\u0078"
	if _bfbaa == nil {
		return nil, _g.Error(_gdcg, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_bfbaa.Boxes)-1 {
		return nil, _g.Errorf(_gdcg, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _bfbaa.Boxes[i], nil
}
func (_dac *Bitmap) GetVanillaData() []byte {
	if _dac.Color == Chocolate {
		_dac.inverseData()
	}
	return _dac.Data
}

const (
	_ SizeSelection = iota
	SizeSelectByWidth
	SizeSelectByHeight
	SizeSelectByMaxDimension
	SizeSelectByArea
	SizeSelectByPerimeter
)

func (_egd *Bitmap) setEightBytes(_add int, _fbbf uint64) error {
	_egae := _egd.RowStride - (_add % _egd.RowStride)
	if _egd.RowStride != _egd.Width>>3 {
		_egae--
	}
	if _egae >= 8 {
		return _egd.setEightFullBytes(_add, _fbbf)
	}
	return _egd.setEightPartlyBytes(_add, _egae, _fbbf)
}
func (_fag *Bitmap) GetByte(index int) (byte, error) {
	if index > len(_fag.Data)-1 || index < 0 {
		return 0, _g.Errorf("\u0047e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _fag.Data[index], nil
}
func (_gacd *Bitmap) GetByteIndex(x, y int) int { return y*_gacd.RowStride + (x >> 3) }
func Extract(roi _aa.Rectangle, src *Bitmap) (*Bitmap, error) {
	_dbfb := New(roi.Dx(), roi.Dy())
	_ffee := roi.Min.X & 0x07
	_faf := 8 - _ffee
	_gdg := uint(8 - _dbfb.Width&0x07)
	_gggbd := src.GetByteIndex(roi.Min.X, roi.Min.Y)
	_addd := src.GetByteIndex(roi.Max.X-1, roi.Min.Y)
	_ggef := _dbfb.RowStride == _addd+1-_gggbd
	var _ffbdb int
	for _ddfd := roi.Min.Y; _ddfd < roi.Max.Y; _ddfd++ {
		_ddae := _gggbd
		_effe := _ffbdb
		switch {
		case _gggbd == _addd:
			_daddf, _fgdga := src.GetByte(_ddae)
			if _fgdga != nil {
				return nil, _fgdga
			}
			_daddf <<= uint(_ffee)
			_fgdga = _dbfb.SetByte(_effe, _cddg(_gdg, _daddf))
			if _fgdga != nil {
				return nil, _fgdga
			}
		case _ffee == 0:
			for _gdbc := _gggbd; _gdbc <= _addd; _gdbc++ {
				_ffbb, _ddcf := src.GetByte(_ddae)
				if _ddcf != nil {
					return nil, _ddcf
				}
				_ddae++
				if _gdbc == _addd && _ggef {
					_ffbb = _cddg(_gdg, _ffbb)
				}
				_ddcf = _dbfb.SetByte(_effe, _ffbb)
				if _ddcf != nil {
					return nil, _ddcf
				}
				_effe++
			}
		default:
			_egbd := _aaad(src, _dbfb, uint(_ffee), uint(_faf), _gdg, _gggbd, _addd, _ggef, _ddae, _effe)
			if _egbd != nil {
				return nil, _egbd
			}
		}
		_gggbd += src.RowStride
		_addd += src.RowStride
		_ffbdb += _dbfb.RowStride
	}
	return _dbfb, nil
}
func _gcee(_bcgdd, _dada, _bef *Bitmap) (*Bitmap, error) {
	const _eef = "\u0073\u0075\u0062\u0074\u0072\u0061\u0063\u0074"
	if _dada == nil {
		return nil, _g.Error(_eef, "'\u0073\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _bef == nil {
		return nil, _g.Error(_eef, "'\u0073\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	var _dbbe error
	switch {
	case _bcgdd == _dada:
		if _dbbe = _bcgdd.RasterOperation(0, 0, _dada.Width, _dada.Height, PixNotSrcAndDst, _bef, 0, 0); _dbbe != nil {
			return nil, _g.Wrap(_dbbe, _eef, "\u0064 \u003d\u003d\u0020\u0073\u0031")
		}
	case _bcgdd == _bef:
		if _dbbe = _bcgdd.RasterOperation(0, 0, _dada.Width, _dada.Height, PixNotSrcAndDst, _dada, 0, 0); _dbbe != nil {
			return nil, _g.Wrap(_dbbe, _eef, "\u0064 \u003d\u003d\u0020\u0073\u0032")
		}
	default:
		_bcgdd, _dbbe = _feea(_bcgdd, _dada)
		if _dbbe != nil {
			return nil, _g.Wrap(_dbbe, _eef, "")
		}
		if _dbbe = _bcgdd.RasterOperation(0, 0, _dada.Width, _dada.Height, PixNotSrcAndDst, _bef, 0, 0); _dbbe != nil {
			return nil, _g.Wrap(_dbbe, _eef, "\u0064e\u0066\u0061\u0075\u006c\u0074")
		}
	}
	return _bcgdd, nil
}
func (_bbae *ClassedPoints) ySortFunction() func(_ffacd int, _gbcd int) bool {
	return func(_cade, _dfcf int) bool { return _bbae.YAtIndex(_cade) < _bbae.YAtIndex(_dfcf) }
}
func (_ffdfe *ClassedPoints) SortByX() { _ffdfe._bacf = _ffdfe.xSortFunction(); _a.Sort(_ffdfe) }
func _fda(_gaf, _ebd *Bitmap, _ceac int, _bgeg []byte, _ggb int) (_bcc error) {
	const _baggg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0034"
	var (
		_daf, _bca, _bagbe, _dga, _dfb, _gee, _gcbb, _fbgc int
		_gcg, _afc                                         uint32
		_gad, _gfe                                         byte
		_daeb                                              uint16
	)
	_agd := make([]byte, 4)
	_gfbf := make([]byte, 4)
	for _bagbe = 0; _bagbe < _gaf.Height-1; _bagbe, _dga = _bagbe+2, _dga+1 {
		_daf = _bagbe * _gaf.RowStride
		_bca = _dga * _ebd.RowStride
		for _dfb, _gee = 0, 0; _dfb < _ggb; _dfb, _gee = _dfb+4, _gee+1 {
			for _gcbb = 0; _gcbb < 4; _gcbb++ {
				_fbgc = _daf + _dfb + _gcbb
				if _fbgc <= len(_gaf.Data)-1 && _fbgc < _daf+_gaf.RowStride {
					_agd[_gcbb] = _gaf.Data[_fbgc]
				} else {
					_agd[_gcbb] = 0x00
				}
				_fbgc = _daf + _gaf.RowStride + _dfb + _gcbb
				if _fbgc <= len(_gaf.Data)-1 && _fbgc < _daf+(2*_gaf.RowStride) {
					_gfbf[_gcbb] = _gaf.Data[_fbgc]
				} else {
					_gfbf[_gcbb] = 0x00
				}
			}
			_gcg = _cc.BigEndian.Uint32(_agd)
			_afc = _cc.BigEndian.Uint32(_gfbf)
			_afc &= _gcg
			_afc &= _afc << 1
			_afc &= 0xaaaaaaaa
			_gcg = _afc | (_afc << 7)
			_gad = byte(_gcg >> 24)
			_gfe = byte((_gcg >> 8) & 0xff)
			_fbgc = _bca + _gee
			if _fbgc+1 == len(_ebd.Data)-1 || _fbgc+1 >= _bca+_ebd.RowStride {
				_ebd.Data[_fbgc] = _bgeg[_gad]
				if _bcc = _ebd.SetByte(_fbgc, _bgeg[_gad]); _bcc != nil {
					return _g.Wrapf(_bcc, _baggg, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _fbgc)
				}
			} else {
				_daeb = (uint16(_bgeg[_gad]) << 8) | uint16(_bgeg[_gfe])
				if _bcc = _ebd.setTwoBytes(_fbgc, _daeb); _bcc != nil {
					return _g.Wrapf(_bcc, _baggg, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fbgc)
				}
				_gee++
			}
		}
	}
	return nil
}
func (_beda *Bitmap) thresholdPixelSum(_ece int) bool {
	var (
		_gcc  int
		_acf  uint8
		_gbdf byte
		_dcfa int
	)
	_dbgc := _beda.RowStride
	_aade := uint(_beda.Width & 0x07)
	if _aade != 0 {
		_acf = uint8((0xff << (8 - _aade)) & 0xff)
		_dbgc--
	}
	for _beaf := 0; _beaf < _beda.Height; _beaf++ {
		for _dcfa = 0; _dcfa < _dbgc; _dcfa++ {
			_gbdf = _beda.Data[_beaf*_beda.RowStride+_dcfa]
			_gcc += int(_ebc[_gbdf])
		}
		if _aade != 0 {
			_gbdf = _beda.Data[_beaf*_beda.RowStride+_dcfa] & _acf
			_gcc += int(_ebc[_gbdf])
		}
		if _gcc > _ece {
			return true
		}
	}
	return false
}
func (_afade *Bitmap) clearAll() error {
	return _afade.RasterOperation(0, 0, _afade.Width, _afade.Height, PixClr, nil, 0, 0)
}
func _fbfd(_adbd *Bitmap, _bgdb, _cacfb, _dbge, _gagb int, _bebg RasterOperator) {
	if _bgdb < 0 {
		_dbge += _bgdb
		_bgdb = 0
	}
	_adcd := _bgdb + _dbge - _adbd.Width
	if _adcd > 0 {
		_dbge -= _adcd
	}
	if _cacfb < 0 {
		_gagb += _cacfb
		_cacfb = 0
	}
	_bdea := _cacfb + _gagb - _adbd.Height
	if _bdea > 0 {
		_gagb -= _bdea
	}
	if _dbge <= 0 || _gagb <= 0 {
		return
	}
	if (_bgdb & 7) == 0 {
		_baefb(_adbd, _bgdb, _cacfb, _dbge, _gagb, _bebg)
	} else {
		_fgcf(_adbd, _bgdb, _cacfb, _dbge, _gagb, _bebg)
	}
}
func (_bbcbd *Bitmaps) GroupByHeight() (*BitmapsArray, error) {
	const _dcfe = "\u0047\u0072\u006f\u0075\u0070\u0042\u0079\u0048\u0065\u0069\u0067\u0068\u0074"
	if len(_bbcbd.Values) == 0 {
		return nil, _g.Error(_dcfe, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_gggbc := &BitmapsArray{}
	_bbcbd.SortByHeight()
	_ccca := -1
	_cbacc := -1
	for _fgafb := 0; _fgafb < len(_bbcbd.Values); _fgafb++ {
		_ebccb := _bbcbd.Values[_fgafb].Height
		if _ebccb > _ccca {
			_ccca = _ebccb
			_cbacc++
			_gggbc.Values = append(_gggbc.Values, &Bitmaps{})
		}
		_gggbc.Values[_cbacc].AddBitmap(_bbcbd.Values[_fgafb])
	}
	return _gggbc, nil
}

type byHeight Bitmaps

func (_gdde *ClassedPoints) YAtIndex(i int) float32 { return (*_gdde.Points)[_gdde.IntSlice[i]].Y }
