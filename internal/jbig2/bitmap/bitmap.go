package bitmap

import (
	_a "encoding/binary"
	_g "image"
	_fc "math"
	_cc "sort"
	_ef "strings"

	_fce "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_fg "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_fa "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_c "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func _gbfb(_defd, _bfcf *Bitmap, _fafe, _aedd int) (*Bitmap, error) {
	const _gabf = "\u006fp\u0065\u006e\u0042\u0072\u0069\u0063k"
	if _bfcf == nil {
		return nil, _c.Error(_gabf, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _fafe < 1 && _aedd < 1 {
		return nil, _c.Error(_gabf, "\u0068\u0053\u0069\u007ae \u003c\u0020\u0031\u0020\u0026\u0026\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u003c \u0031")
	}
	if _fafe == 1 && _aedd == 1 {
		return _bfcf.Copy(), nil
	}
	if _fafe == 1 || _aedd == 1 {
		var _bfa error
		_ecca := SelCreateBrick(_aedd, _fafe, _aedd/2, _fafe/2, SelHit)
		_defd, _bfa = _ddgge(_defd, _bfcf, _ecca)
		if _bfa != nil {
			return nil, _c.Wrap(_bfa, _gabf, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _defd, nil
	}
	_bfgf := SelCreateBrick(1, _fafe, 0, _fafe/2, SelHit)
	_gccd := SelCreateBrick(_aedd, 1, _aedd/2, 0, SelHit)
	_bfee, _bfcfc := _gcae(nil, _bfcf, _bfgf)
	if _bfcfc != nil {
		return nil, _c.Wrap(_bfcfc, _gabf, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_defd, _bfcfc = _gcae(_defd, _bfee, _gccd)
	if _bfcfc != nil {
		return nil, _c.Wrap(_bfcfc, _gabf, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	_, _bfcfc = _bag(_bfee, _defd, _bfgf)
	if _bfcfc != nil {
		return nil, _c.Wrap(_bfcfc, _gabf, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_, _bfcfc = _bag(_defd, _bfee, _gccd)
	if _bfcfc != nil {
		return nil, _c.Wrap(_bfcfc, _gabf, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _defd, nil
}

type (
	Getter        interface{ GetBitmap() *Bitmap }
	ClassedPoints struct {
		*Points
		_fa.IntSlice
		_bdgab func(_beag, _acaa int) bool
	}
)

func (_afdag *Bitmap) ClipRectangle(box *_g.Rectangle) (_bbc *Bitmap, _cabg *_g.Rectangle, _ddd error) {
	const _bab = "\u0043\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, nil, _c.Error(_bab, "\u0062o\u0078 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	_aee, _dc := _afdag.Width, _afdag.Height
	_dfad := _g.Rect(0, 0, _aee, _dc)
	if !box.Overlaps(_dfad) {
		return nil, nil, _c.Error(_bab, "b\u006f\u0078\u0020\u0064oe\u0073n\u0027\u0074\u0020\u006f\u0076e\u0072\u006c\u0061\u0070\u0020\u0062")
	}
	_baag := box.Intersect(_dfad)
	_eec, _decee := _baag.Min.X, _baag.Min.Y
	_bgf, _ffa := _baag.Dx(), _baag.Dy()
	_bbc = New(_bgf, _ffa)
	_bbc.Text = _afdag.Text
	if _ddd = _bbc.RasterOperation(0, 0, _bgf, _ffa, PixSrc, _afdag, _eec, _decee); _ddd != nil {
		return nil, nil, _c.Wrap(_ddd, _bab, "\u0050\u0069\u0078\u0053\u0072\u0063\u0020\u0074\u006f\u0020\u0063\u006ci\u0070\u0070\u0065\u0064")
	}
	_cabg = &_baag
	return _bbc, _cabg, nil
}

func (_bce *Bitmap) resizeImageData(_ddc *Bitmap) error {
	if _ddc == nil {
		return _c.Error("\u0072e\u0073i\u007a\u0065\u0049\u006d\u0061\u0067\u0065\u0044\u0061\u0074\u0061", "\u0073r\u0063 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _bce.SizesEqual(_ddc) {
		return nil
	}
	_bce.Data = make([]byte, len(_ddc.Data))
	_bce.Width = _ddc.Width
	_bce.Height = _ddc.Height
	_bce.RowStride = _ddc.RowStride
	return nil
}

const (
	CmbOpOr CombinationOperator = iota
	CmbOpAnd
	CmbOpXor
	CmbOpXNor
	CmbOpReplace
	CmbOpNot
)

func (_gaac *Boxes) Add(box *_g.Rectangle) error {
	if _gaac == nil {
		return _c.Error("\u0042o\u0078\u0065\u0073\u002e\u0041\u0064d", "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	*_gaac = append(*_gaac, box)
	return nil
}

func _bg() (_eba [256]uint32) {
	for _bcc := 0; _bcc < 256; _bcc++ {
		if _bcc&0x01 != 0 {
			_eba[_bcc] |= 0xf
		}
		if _bcc&0x02 != 0 {
			_eba[_bcc] |= 0xf0
		}
		if _bcc&0x04 != 0 {
			_eba[_bcc] |= 0xf00
		}
		if _bcc&0x08 != 0 {
			_eba[_bcc] |= 0xf000
		}
		if _bcc&0x10 != 0 {
			_eba[_bcc] |= 0xf0000
		}
		if _bcc&0x20 != 0 {
			_eba[_bcc] |= 0xf00000
		}
		if _bcc&0x40 != 0 {
			_eba[_bcc] |= 0xf000000
		}
		if _bcc&0x80 != 0 {
			_eba[_bcc] |= 0xf0000000
		}
	}
	return _eba
}

func _daf(_geede, _gfdd *Bitmap, _daga, _bbbg int) (_gafgg error) {
	const _acec = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0038"
	var (
		_aeef, _bfge, _cgcf, _cfef                                int
		_fddga, _dfge, _ffbcd, _edfc, _aedge, _fccc, _eabc, _debd byte
	)
	for _aeef = 0; _aeef < _daga; _aeef++ {
		_cgcf = _aeef * _geede.RowStride
		_cfef = _aeef * _gfdd.RowStride
		for _bfge = 0; _bfge < _bbbg; _bfge++ {
			if _fddga, _gafgg = _geede.GetByte(_cgcf + _bfge); _gafgg != nil {
				return _c.Wrap(_gafgg, _acec, "\u0067e\u0074 \u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _dfge, _gafgg = _gfdd.GetByte(_cfef + _bfge); _gafgg != nil {
				return _c.Wrap(_gafgg, _acec, "\u0067\u0065\u0074\u0020\u006d\u0061\u0073\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _aeef > 0 {
				if _ffbcd, _gafgg = _geede.GetByte(_cgcf - _geede.RowStride + _bfge); _gafgg != nil {
					return _c.Wrap(_gafgg, _acec, "\u0069\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
				}
				_fddga |= _ffbcd | (_ffbcd << 1) | (_ffbcd >> 1)
				if _bfge > 0 {
					if _debd, _gafgg = _geede.GetByte(_cgcf - _geede.RowStride + _bfge - 1); _gafgg != nil {
						return _c.Wrap(_gafgg, _acec, "\u0069\u0020\u003e\u00200 \u0026\u0026\u0020\u006a\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
					}
					_fddga |= _debd << 7
				}
				if _bfge < _bbbg-1 {
					if _debd, _gafgg = _geede.GetByte(_cgcf - _geede.RowStride + _bfge + 1); _gafgg != nil {
						return _c.Wrap(_gafgg, _acec, "\u006a\u0020<\u0020\u0077\u0070l\u0020\u002d\u0020\u0031\u0020\u0062\u0079\u0074\u0065")
					}
					_fddga |= _debd >> 7
				}
			}
			if _bfge > 0 {
				if _edfc, _gafgg = _geede.GetByte(_cgcf + _bfge - 1); _gafgg != nil {
					return _c.Wrap(_gafgg, _acec, "\u006a\u0020\u003e \u0030")
				}
				_fddga |= _edfc << 7
			}
			_fddga &= _dfge
			if _fddga == 0 || ^_fddga == 0 {
				if _gafgg = _geede.SetByte(_cgcf+_bfge, _fddga); _gafgg != nil {
					return _c.Wrap(_gafgg, _acec, "\u0073e\u0074t\u0069\u006e\u0067\u0020\u0065m\u0070\u0074y\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_eabc = _fddga
				_fddga = (_fddga | (_fddga >> 1) | (_fddga << 1)) & _dfge
				if (_fddga ^ _eabc) == 0 {
					if _gafgg = _geede.SetByte(_cgcf+_bfge, _fddga); _gafgg != nil {
						return _c.Wrap(_gafgg, _acec, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _aeef = _daga - 1; _aeef >= 0; _aeef-- {
		_cgcf = _aeef * _geede.RowStride
		_cfef = _aeef * _gfdd.RowStride
		for _bfge = _bbbg - 1; _bfge >= 0; _bfge-- {
			if _fddga, _gafgg = _geede.GetByte(_cgcf + _bfge); _gafgg != nil {
				return _c.Wrap(_gafgg, _acec, "\u0072\u0065\u0076er\u0073\u0065\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _dfge, _gafgg = _gfdd.GetByte(_cfef + _bfge); _gafgg != nil {
				return _c.Wrap(_gafgg, _acec, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _aeef < _daga-1 {
				if _aedge, _gafgg = _geede.GetByte(_cgcf + _geede.RowStride + _bfge); _gafgg != nil {
					return _c.Wrap(_gafgg, _acec, "\u0069\u0020\u003c\u0020h\u0020\u002d\u0020\u0031\u0020\u002d\u003e\u0020\u0067\u0065t\u0020s\u006f\u0075\u0072\u0063\u0065\u0020\u0062y\u0074\u0065")
				}
				_fddga |= _aedge | (_aedge << 1) | _aedge>>1
				if _bfge > 0 {
					if _debd, _gafgg = _geede.GetByte(_cgcf + _geede.RowStride + _bfge - 1); _gafgg != nil {
						return _c.Wrap(_gafgg, _acec, "\u0069\u0020\u003c h\u002d\u0031\u0020\u0026\u0020\u006a\u0020\u003e\u00200\u0020-\u003e \u0067e\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
					}
					_fddga |= _debd << 7
				}
				if _bfge < _bbbg-1 {
					if _debd, _gafgg = _geede.GetByte(_cgcf + _geede.RowStride + _bfge + 1); _gafgg != nil {
						return _c.Wrap(_gafgg, _acec, "\u0069\u0020\u003c\u0020\u0068\u002d\u0031\u0020\u0026\u0026\u0020\u006a\u0020\u003c\u0077\u0070\u006c\u002d\u0031\u0020\u002d\u003e\u0020\u0067e\u0074\u0020\u0073\u006f\u0075r\u0063\u0065 \u0062\u0079\u0074\u0065")
					}
					_fddga |= _debd >> 7
				}
			}
			if _bfge < _bbbg-1 {
				if _fccc, _gafgg = _geede.GetByte(_cgcf + _bfge + 1); _gafgg != nil {
					return _c.Wrap(_gafgg, _acec, "\u006a\u0020<\u0020\u0077\u0070\u006c\u0020\u002d\u0031\u0020\u002d\u003e\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020by\u0074\u0065")
				}
				_fddga |= _fccc >> 7
			}
			_fddga &= _dfge
			if _fddga == 0 || (^_fddga) == 0 {
				if _gafgg = _geede.SetByte(_cgcf+_bfge, _fddga); _gafgg != nil {
					return _c.Wrap(_gafgg, _acec, "\u0073e\u0074 \u006d\u0061\u0073\u006b\u0065\u0064\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_eabc = _fddga
				_fddga = (_fddga | (_fddga >> 1) | (_fddga << 1)) & _dfge
				if (_fddga ^ _eabc) == 0 {
					if _gafgg = _geede.SetByte(_cgcf+_bfge, _fddga); _gafgg != nil {
						return _c.Wrap(_gafgg, _acec, "r\u0065\u0076\u0065\u0072se\u0020s\u0065\u0074\u0020\u0070\u0072e\u0076\u0020\u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	return nil
}

type SizeComparison int

func (_gcbf *Bitmaps) GroupByHeight() (*BitmapsArray, error) {
	const _dad = "\u0047\u0072\u006f\u0075\u0070\u0042\u0079\u0048\u0065\u0069\u0067\u0068\u0074"
	if len(_gcbf.Values) == 0 {
		return nil, _c.Error(_dad, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_cdcd := &BitmapsArray{}
	_gcbf.SortByHeight()
	_dfec := -1
	_gbab := -1
	for _gcdaa := 0; _gcdaa < len(_gcbf.Values); _gcdaa++ {
		_bgcg := _gcbf.Values[_gcdaa].Height
		if _bgcg > _dfec {
			_dfec = _bgcg
			_gbab++
			_cdcd.Values = append(_cdcd.Values, &Bitmaps{})
		}
		_cdcd.Values[_gbab].AddBitmap(_gcbf.Values[_gcdaa])
	}
	return _cdcd, nil
}
func (_aadcg *ClassedPoints) SortByX() { _aadcg._bdgab = _aadcg.xSortFunction(); _cc.Sort(_aadcg) }
func _edff(_ffd, _bdd *Bitmap, _fdge int, _baf []byte, _gfbcd int) (_be error) {
	const _bge = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0034"
	var (
		_eca, _agg, _feg, _ddb, _gdg, _gfae, _ced, _dece int
		_bdc, _faf                                       uint32
		_gac, _eeb                                       byte
		_ebf                                             uint16
	)
	_egbc := make([]byte, 4)
	_cgc := make([]byte, 4)
	for _feg = 0; _feg < _ffd.Height-1; _feg, _ddb = _feg+2, _ddb+1 {
		_eca = _feg * _ffd.RowStride
		_agg = _ddb * _bdd.RowStride
		for _gdg, _gfae = 0, 0; _gdg < _gfbcd; _gdg, _gfae = _gdg+4, _gfae+1 {
			for _ced = 0; _ced < 4; _ced++ {
				_dece = _eca + _gdg + _ced
				if _dece <= len(_ffd.Data)-1 && _dece < _eca+_ffd.RowStride {
					_egbc[_ced] = _ffd.Data[_dece]
				} else {
					_egbc[_ced] = 0x00
				}
				_dece = _eca + _ffd.RowStride + _gdg + _ced
				if _dece <= len(_ffd.Data)-1 && _dece < _eca+(2*_ffd.RowStride) {
					_cgc[_ced] = _ffd.Data[_dece]
				} else {
					_cgc[_ced] = 0x00
				}
			}
			_bdc = _a.BigEndian.Uint32(_egbc)
			_faf = _a.BigEndian.Uint32(_cgc)
			_faf &= _bdc
			_faf &= _faf << 1
			_faf &= 0xaaaaaaaa
			_bdc = _faf | (_faf << 7)
			_gac = byte(_bdc >> 24)
			_eeb = byte((_bdc >> 8) & 0xff)
			_dece = _agg + _gfae
			if _dece+1 == len(_bdd.Data)-1 || _dece+1 >= _agg+_bdd.RowStride {
				_bdd.Data[_dece] = _baf[_gac]
				if _be = _bdd.SetByte(_dece, _baf[_gac]); _be != nil {
					return _c.Wrapf(_be, _bge, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _dece)
				}
			} else {
				_ebf = (uint16(_baf[_gac]) << 8) | uint16(_baf[_eeb])
				if _be = _bdd.setTwoBytes(_dece, _ebf); _be != nil {
					return _c.Wrapf(_be, _bge, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _dece)
				}
				_gfae++
			}
		}
	}
	return nil
}

func (_aeed Points) XSorter() func(_ecbb, _fcbga int) bool {
	return func(_ffe, _aea int) bool { return _aeed[_ffe].X < _aeed[_aea].X }
}
func (_ffagc *ClassedPoints) Len() int              { return _ffagc.IntSlice.Size() }
func (_aafc *ClassedPoints) YAtIndex(i int) float32 { return (*_aafc.Points)[_aafc.IntSlice[i]].Y }
func _bbegg(_dggee, _cacee *Bitmap, _adga *Selection) (*Bitmap, error) {
	const _bga = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0032"
	var _gadd, _aaba int
	if _cacee == nil {
		return nil, _c.Error(_bga, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _adga == nil {
		return nil, _c.Error(_bga, "\u0073e\u006c \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_gadd = _adga.Width
	_aaba = _adga.Height
	if _gadd == 0 || _aaba == 0 {
		return nil, _c.Error(_bga, "\u0073\u0065\u006c\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _dggee == nil {
		return _cacee.createTemplate(), nil
	}
	if _dgced := _dggee.resizeImageData(_cacee); _dgced != nil {
		return nil, _dgced
	}
	return _dggee, nil
}

func init() {
	for _fee := 0; _fee < 256; _fee++ {
		_cee[_fee] = uint8(_fee&0x1) + (uint8(_fee>>1) & 0x1) + (uint8(_fee>>2) & 0x1) + (uint8(_fee>>3) & 0x1) + (uint8(_fee>>4) & 0x1) + (uint8(_fee>>5) & 0x1) + (uint8(_fee>>6) & 0x1) + (uint8(_fee>>7) & 0x1)
	}
}
func (_cege *ClassedPoints) Less(i, j int) bool { return _cege._bdgab(i, j) }
func (_egdfd *Bitmaps) Size() int               { return len(_egdfd.Values) }
func RankHausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH, area1, area3 int, rank float32, tab8 []int) (_fgec bool, _cdef error) {
	const _fbeg = "\u0052\u0061\u006ek\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_bfdd, _adgc := p1.Width, p1.Height
	_ecdc, _cfcf := p3.Width, p3.Height
	if _fa.Abs(_bfdd-_ecdc) > maxDiffW {
		return false, nil
	}
	if _fa.Abs(_adgc-_cfcf) > maxDiffH {
		return false, nil
	}
	_afcb := int(float32(area1)*(1.0-rank) + 0.5)
	_becb := int(float32(area3)*(1.0-rank) + 0.5)
	var _ddgbc, _fged int
	if delX >= 0 {
		_ddgbc = int(delX + 0.5)
	} else {
		_ddgbc = int(delX - 0.5)
	}
	if delY >= 0 {
		_fged = int(delY + 0.5)
	} else {
		_fged = int(delY - 0.5)
	}
	_bgfb := p1.CreateTemplate()
	if _cdef = _bgfb.RasterOperation(0, 0, _bfdd, _adgc, PixSrc, p1, 0, 0); _cdef != nil {
		return false, _c.Wrap(_cdef, _fbeg, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _cdef = _bgfb.RasterOperation(_ddgbc, _fged, _bfdd, _adgc, PixNotSrcAndDst, p4, 0, 0); _cdef != nil {
		return false, _c.Wrap(_cdef, _fbeg, "\u0074 \u0026\u0020\u0021\u0070\u0034")
	}
	_fgec, _cdef = _bgfb.ThresholdPixelSum(_afcb, tab8)
	if _cdef != nil {
		return false, _c.Wrap(_cdef, _fbeg, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0031")
	}
	if _fgec {
		return false, nil
	}
	if _cdef = _bgfb.RasterOperation(_ddgbc, _fged, _ecdc, _cfcf, PixSrc, p3, 0, 0); _cdef != nil {
		return false, _c.Wrap(_cdef, _fbeg, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _cdef = _bgfb.RasterOperation(0, 0, _ecdc, _cfcf, PixNotSrcAndDst, p2, 0, 0); _cdef != nil {
		return false, _c.Wrap(_cdef, _fbeg, "\u0074 \u0026\u0020\u0021\u0070\u0032")
	}
	_fgec, _cdef = _bgfb.ThresholdPixelSum(_becb, tab8)
	if _cdef != nil {
		return false, _c.Wrap(_cdef, _fbeg, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0033")
	}
	return !_fgec, nil
}

func _aeg(_cgf, _dgc int) *Bitmap {
	return &Bitmap{Width: _cgf, Height: _dgc, RowStride: (_cgf + 7) >> 3}
}

func (_abf *Bitmap) nextOnPixel(_cabgf, _egac int) (_cabe _g.Point, _fcde bool, _ffca error) {
	const _cdaa = "n\u0065\u0078\u0074\u004f\u006e\u0050\u0069\u0078\u0065\u006c"
	_cabe, _fcde, _ffca = _abf.nextOnPixelLow(_abf.Width, _abf.Height, _abf.RowStride, _cabgf, _egac)
	if _ffca != nil {
		return _cabe, false, _c.Wrap(_ffca, _cdaa, "")
	}
	return _cabe, _fcde, nil
}

func (_gec *Bitmap) GetChocolateData() []byte {
	if _gec.Color == Vanilla {
		_gec.inverseData()
	}
	return _gec.Data
}
func (_ggfbd *byWidth) Less(i, j int) bool { return _ggfbd.Values[i].Width < _ggfbd.Values[j].Width }
func (_ffac *Bitmap) Equals(s *Bitmap) bool {
	if len(_ffac.Data) != len(s.Data) || _ffac.Width != s.Width || _ffac.Height != s.Height {
		return false
	}
	for _ecaa := 0; _ecaa < _ffac.Height; _ecaa++ {
		_egg := _ecaa * _ffac.RowStride
		for _ccd := 0; _ccd < _ffac.RowStride; _ccd++ {
			if _ffac.Data[_egg+_ccd] != s.Data[_egg+_ccd] {
				return false
			}
		}
	}
	return true
}
func (_cebca *Bitmaps) SortByWidth() { _ebaa := (*byWidth)(_cebca); _cc.Sort(_ebaa) }
func (_dag *Bitmap) SetPixel(x, y int, pixel byte) error {
	_bac := _dag.GetByteIndex(x, y)
	if _bac > len(_dag.Data)-1 {
		return _c.Errorf("\u0053\u0065\u0074\u0050\u0069\u0078\u0065\u006c", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", _bac)
	}
	_efg := _dag.GetBitOffset(x)
	_gafd := uint(7 - _efg)
	_aeb := _dag.Data[_bac]
	var _bddb byte
	if pixel == 1 {
		_bddb = _aeb | (pixel & 0x01 << _gafd)
	} else {
		_bddb = _aeb &^ (1 << _gafd)
	}
	_dag.Data[_bac] = _bddb
	return nil
}

func (_abg *Bitmap) GetVanillaData() []byte {
	if _abg.Color == Chocolate {
		_abg.inverseData()
	}
	return _abg.Data
}
func (_bdde *ClassedPoints) SortByY() { _bdde._bdgab = _bdde.ySortFunction(); _cc.Sort(_bdde) }
func (_gafg Points) YSorter() func(_fbb, _aabb int) bool {
	return func(_gdfc, _fbag int) bool { return _gafg[_gdfc].Y < _gafg[_fbag].Y }
}

func _cfde(_baad, _efbca *Bitmap, _addfa, _gace int) (*Bitmap, error) {
	const _abfg = "d\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063\u006b"
	if _efbca == nil {
		_fce.Log.Debug("\u0064\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063k\u0020\u0073\u006f\u0075\u0072\u0063\u0065 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
		return nil, _c.Error(_abfg, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _addfa < 1 || _gace < 1 {
		return nil, _c.Error(_abfg, "\u0068\u0053\u007a\u0069\u0065 \u0061\u006e\u0064\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u0061\u0072\u0065 \u006e\u006f\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _addfa == 1 && _gace == 1 {
		_agaf, _ebdad := _gbfag(_baad, _efbca)
		if _ebdad != nil {
			return nil, _c.Wrap(_ebdad, _abfg, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _agaf, nil
	}
	if _addfa == 1 || _gace == 1 {
		_ffae := SelCreateBrick(_gace, _addfa, _gace/2, _addfa/2, SelHit)
		_eddg, _ddgba := _bag(_baad, _efbca, _ffae)
		if _ddgba != nil {
			return nil, _c.Wrap(_ddgba, _abfg, "\u0068s\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _eddg, nil
	}
	_cfaa := SelCreateBrick(1, _addfa, 0, _addfa/2, SelHit)
	_dfgg := SelCreateBrick(_gace, 1, _gace/2, 0, SelHit)
	_cace, _aaab := _bag(nil, _efbca, _cfaa)
	if _aaab != nil {
		return nil, _c.Wrap(_aaab, _abfg, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_baad, _aaab = _bag(_baad, _cace, _dfgg)
	if _aaab != nil {
		return nil, _c.Wrap(_aaab, _abfg, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _baad, nil
}

func (_dabb CombinationOperator) String() string {
	var _edad string
	switch _dabb {
	case CmbOpOr:
		_edad = "\u004f\u0052"
	case CmbOpAnd:
		_edad = "\u0041\u004e\u0044"
	case CmbOpXor:
		_edad = "\u0058\u004f\u0052"
	case CmbOpXNor:
		_edad = "\u0058\u004e\u004f\u0052"
	case CmbOpReplace:
		_edad = "\u0052E\u0050\u004c\u0041\u0043\u0045"
	case CmbOpNot:
		_edad = "\u004e\u004f\u0054"
	}
	return _edad
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

func _fcbb(_ebcd, _ebce, _eggdb byte) byte { return (_ebcd &^ (_eggdb)) | (_ebce & _eggdb) }
func _gbfc(_cfcb *Bitmap, _aadb ...MorphProcess) (_afcga *Bitmap, _efca error) {
	const _ggca = "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065"
	if _cfcb == nil {
		return nil, _c.Error(_ggca, "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065 \u0073\u006f\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if len(_aadb) == 0 {
		return nil, _c.Error(_ggca, "m\u006f\u0072\u0070\u0068\u0053\u0065q\u0075\u0065\u006e\u0063\u0065\u002c \u0073\u0065\u0071\u0075\u0065\u006e\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _efca = _adcca(_aadb...); _efca != nil {
		return nil, _c.Wrap(_efca, _ggca, "")
	}
	var _cagc, _aebd, _bdbd int
	_afcga = _cfcb.Copy()
	for _, _gdca := range _aadb {
		switch _gdca.Operation {
		case MopDilation:
			_cagc, _aebd = _gdca.getWidthHeight()
			_afcga, _efca = DilateBrick(nil, _afcga, _cagc, _aebd)
			if _efca != nil {
				return nil, _c.Wrap(_efca, _ggca, "")
			}
		case MopErosion:
			_cagc, _aebd = _gdca.getWidthHeight()
			_afcga, _efca = _dge(nil, _afcga, _cagc, _aebd)
			if _efca != nil {
				return nil, _c.Wrap(_efca, _ggca, "")
			}
		case MopOpening:
			_cagc, _aebd = _gdca.getWidthHeight()
			_afcga, _efca = _gbfb(nil, _afcga, _cagc, _aebd)
			if _efca != nil {
				return nil, _c.Wrap(_efca, _ggca, "")
			}
		case MopClosing:
			_cagc, _aebd = _gdca.getWidthHeight()
			_afcga, _efca = _ecbe(nil, _afcga, _cagc, _aebd)
			if _efca != nil {
				return nil, _c.Wrap(_efca, _ggca, "")
			}
		case MopRankBinaryReduction:
			_afcga, _efca = _fga(_afcga, _gdca.Arguments...)
			if _efca != nil {
				return nil, _c.Wrap(_efca, _ggca, "")
			}
		case MopReplicativeBinaryExpansion:
			_afcga, _efca = _ebbf(_afcga, _gdca.Arguments[0])
			if _efca != nil {
				return nil, _c.Wrap(_efca, _ggca, "")
			}
		case MopAddBorder:
			_bdbd = _gdca.Arguments[0]
			_afcga, _efca = _afcga.AddBorder(_bdbd, 0)
			if _efca != nil {
				return nil, _c.Wrap(_efca, _ggca, "")
			}
		default:
			return nil, _c.Error(_ggca, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u006fr\u0070\u0068\u004f\u0070\u0065\u0072\u0061ti\u006f\u006e\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u006f t\u0068\u0065 \u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065")
		}
	}
	if _bdbd > 0 {
		_afcga, _efca = _afcga.RemoveBorder(_bdbd)
		if _efca != nil {
			return nil, _c.Wrap(_efca, _ggca, "\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u003e\u0020\u0030")
		}
	}
	return _afcga, nil
}

func (_gbfa *Bitmap) GetPixel(x, y int) bool {
	_ebfg := _gbfa.GetByteIndex(x, y)
	_aeea := _gbfa.GetBitOffset(x)
	_fffc := uint(7 - _aeea)
	if _ebfg > len(_gbfa.Data)-1 {
		_fce.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0072\u0061\u006e\u0067\u0065\u002e \u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0027\u0025\u0064'\u002c\u0020\u0062m\u003a\u0020\u0027\u0025\u0073\u0027", x, y, _gbfa)
		return false
	}
	if (_gbfa.Data[_ebfg]>>_fffc)&0x01 >= 1 {
		return true
	}
	return false
}

var _bbcde = [5]int{1, 2, 3, 0, 4}

func _aecg(_adaf *Bitmap, _dbbe *Bitmap, _efba *Selection, _egga **Bitmap) (*Bitmap, error) {
	const _gcef = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0031"
	if _dbbe == nil {
		return nil, _c.Error(_gcef, "\u004d\u006f\u0072\u0070\u0068\u0041\u0072\u0067\u0073\u0031\u0020'\u0073\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066i\u006e\u0065\u0064")
	}
	if _efba == nil {
		return nil, _c.Error(_gcef, "\u004d\u006f\u0072\u0068p\u0041\u0072\u0067\u0073\u0031\u0020\u0027\u0073\u0065\u006c'\u0020n\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064")
	}
	_agbd, _fabf := _efba.Height, _efba.Width
	if _agbd == 0 || _fabf == 0 {
		return nil, _c.Error(_gcef, "\u0073\u0065\u006c\u0065ct\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _adaf == nil {
		_adaf = _dbbe.createTemplate()
		*_egga = _dbbe
		return _adaf, nil
	}
	_adaf.Width = _dbbe.Width
	_adaf.Height = _dbbe.Height
	_adaf.RowStride = _dbbe.RowStride
	_adaf.Color = _dbbe.Color
	_adaf.Data = make([]byte, _dbbe.RowStride*_dbbe.Height)
	if _adaf == _dbbe {
		*_egga = _dbbe.Copy()
	} else {
		*_egga = _dbbe
	}
	return _adaf, nil
}

func (_bccfa MorphProcess) verify(_bdfc int, _cge, _gfeb *int) error {
	const _gbcb = "\u004d\u006f\u0072\u0070hP\u0072\u006f\u0063\u0065\u0073\u0073\u002e\u0076\u0065\u0072\u0069\u0066\u0079"
	switch _bccfa.Operation {
	case MopDilation, MopErosion, MopOpening, MopClosing:
		if len(_bccfa.Arguments) != 2 {
			return _c.Error(_gbcb, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0064\u0027\u002c\u0020\u0027\u0065\u0027\u002c \u0027\u006f\u0027\u002c\u0020\u0027\u0063\u0027\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0032\u0020\u0061r\u0067\u0075\u006d\u0065\u006et\u0073")
		}
		_fede, _fggf := _bccfa.getWidthHeight()
		if _fede <= 0 || _fggf <= 0 {
			return _c.Error(_gbcb, "O\u0070er\u0061t\u0069o\u006e\u003a\u0020\u0027\u0064'\u002c\u0020\u0027e\u0027\u002c\u0020\u0027\u006f'\u002c\u0020\u0027c\u0027\u0020\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0062\u006f\u0074h w\u0069\u0064\u0074\u0068\u0020\u0061n\u0064\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0074\u006f\u0020b\u0065 \u003e\u003d\u0020\u0030")
		}
	case MopRankBinaryReduction:
		_gegf := len(_bccfa.Arguments)
		*_cge += _gegf
		if _gegf < 1 || _gegf > 4 {
			return _c.Error(_gbcb, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0072\u0027\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020\u006d\u006fs\u0074\u0020\u0034\u0020\u0061\u0072g\u0075\u006d\u0065n\u0074\u0073")
		}
		for _cece := 0; _cece < _gegf; _cece++ {
			if _bccfa.Arguments[_cece] < 1 || _bccfa.Arguments[_cece] > 4 {
				return _c.Error(_gbcb, "\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0052\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 \u00280\u002c\u0020\u0034\u003e")
			}
		}
	case MopReplicativeBinaryExpansion:
		if len(_bccfa.Arguments) == 0 {
			return _c.Error(_gbcb, "\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074i\u0076\u0065\u0042in\u0061\u0072\u0079\u0045\u0078\u0070a\u006e\u0073\u0069\u006f\u006e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020o\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_ebaf := _bccfa.Arguments[0]
		if _ebaf != 2 && _ebaf != 4 && _ebaf != 8 {
			return _c.Error(_gbcb, "R\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u0076\u0065\u0042\u0069\u006e\u0061\u0072\u0079\u0045\u0078\u0070\u0061\u006e\u0073\u0069\u006f\u006e\u0020m\u0075s\u0074\u0020\u0062\u0065 \u006f\u0066 \u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
		}
		*_cge -= _bbcde[_ebaf/4]
	case MopAddBorder:
		if len(_bccfa.Arguments) == 0 {
			return _c.Error(_gbcb, "\u0041\u0064\u0064B\u006f\u0072\u0064\u0065r\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_acbg := _bccfa.Arguments[0]
		if _bdfc > 0 {
			return _c.Error(_gbcb, "\u0041\u0064\u0064\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020f\u0069\u0072\u0073\u0074\u0020\u006d\u006f\u0072\u0070\u0068\u0020\u0070\u0072o\u0063\u0065\u0073\u0073")
		}
		if _acbg < 1 {
			return _c.Error(_gbcb, "\u0041\u0064\u0064\u0042o\u0072\u0064\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065 \u006co\u0077\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
		}
		*_gfeb = _acbg
	}
	return nil
}
func Centroid(bm *Bitmap, centTab, sumTab []int) (Point, error) { return bm.centroid(centTab, sumTab) }
func _cede(_agggfd *Bitmap, _cefe, _caabg int, _ffagb, _ceecf int, _ddad RasterOperator, _fgfe *Bitmap, _ceff, _aaaa int) error {
	var _geggc, _cfdd, _gagf, _egacd int
	if _cefe < 0 {
		_ceff -= _cefe
		_ffagb += _cefe
		_cefe = 0
	}
	if _ceff < 0 {
		_cefe -= _ceff
		_ffagb += _ceff
		_ceff = 0
	}
	_geggc = _cefe + _ffagb - _agggfd.Width
	if _geggc > 0 {
		_ffagb -= _geggc
	}
	_cfdd = _ceff + _ffagb - _fgfe.Width
	if _cfdd > 0 {
		_ffagb -= _cfdd
	}
	if _caabg < 0 {
		_aaaa -= _caabg
		_ceecf += _caabg
		_caabg = 0
	}
	if _aaaa < 0 {
		_caabg -= _aaaa
		_ceecf += _aaaa
		_aaaa = 0
	}
	_gagf = _caabg + _ceecf - _agggfd.Height
	if _gagf > 0 {
		_ceecf -= _gagf
	}
	_egacd = _aaaa + _ceecf - _fgfe.Height
	if _egacd > 0 {
		_ceecf -= _egacd
	}
	if _ffagb <= 0 || _ceecf <= 0 {
		return nil
	}
	var _dffd error
	switch {
	case _cefe&7 == 0 && _ceff&7 == 0:
		_dffd = _gcda(_agggfd, _cefe, _caabg, _ffagb, _ceecf, _ddad, _fgfe, _ceff, _aaaa)
	case _cefe&7 == _ceff&7:
		_dffd = _agfd(_agggfd, _cefe, _caabg, _ffagb, _ceecf, _ddad, _fgfe, _ceff, _aaaa)
	default:
		_dffd = _bcde(_agggfd, _cefe, _caabg, _ffagb, _ceecf, _ddad, _fgfe, _ceff, _aaaa)
	}
	if _dffd != nil {
		return _c.Wrap(_dffd, "r\u0061\u0073\u0074\u0065\u0072\u004f\u0070\u004c\u006f\u0077", "")
	}
	return nil
}

func (_cccfb *Boxes) selectWithIndicator(_edef *_fa.NumSlice) (_egdeg *Boxes, _ffcg error) {
	const _gaaa = "\u0042o\u0078\u0065\u0073\u002es\u0065\u006c\u0065\u0063\u0074W\u0069t\u0068I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _cccfb == nil {
		return nil, _c.Error(_gaaa, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _edef == nil {
		return nil, _c.Error(_gaaa, "\u0027\u006ea\u0027\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_edef) != len(*_cccfb) {
		return nil, _c.Error(_gaaa, "\u0062\u006f\u0078\u0065\u0073\u0020\u0027\u0062\u0027\u0020\u0068\u0061\u0073\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020s\u0069\u007a\u0065\u0020\u0074h\u0061\u006e \u0027\u006e\u0061\u0027")
	}
	var _acge, _dgcg int
	for _bega := 0; _bega < len(*_edef); _bega++ {
		if _acge, _ffcg = _edef.GetInt(_bega); _ffcg != nil {
			return nil, _c.Wrap(_ffcg, _gaaa, "\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u0067\u0020c\u006f\u0075\u006e\u0074")
		}
		if _acge == 1 {
			_dgcg++
		}
	}
	if _dgcg == len(*_cccfb) {
		return _cccfb, nil
	}
	_eeefc := Boxes{}
	for _eddb := 0; _eddb < len(*_edef); _eddb++ {
		_acge = int((*_edef)[_eddb])
		if _acge == 0 {
			continue
		}
		_eeefc = append(_eeefc, (*_cccfb)[_eddb])
	}
	_egdeg = &_eeefc
	return _egdeg, nil
}

func (_ebad Points) Get(i int) (Point, error) {
	if i > len(_ebad)-1 {
		return Point{}, _c.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _ebad[i], nil
}

var (
	_fceec = _ga()
	_fdag  = _bg()
	_aebb  = _ceg()
)

func (_gef *Points) Add(pt *Points) error {
	const _eggc = "\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0041\u0064\u0064"
	if _gef == nil {
		return _c.Error(_eggc, "\u0070o\u0069n\u0074\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if pt == nil {
		return _c.Error(_eggc, "a\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	*_gef = append(*_gef, *pt...)
	return nil
}

func _acef(_edeg int) int {
	if _edeg < 0 {
		return -_edeg
	}
	return _edeg
}

const (
	Vanilla Color = iota
	Chocolate
)

func (_bgg *Bitmap) String() string {
	_aec := "\u000a"
	for _ddgg := 0; _ddgg < _bgg.Height; _ddgg++ {
		var _eeeg string
		for _fgaf := 0; _fgaf < _bgg.Width; _fgaf++ {
			_fgg := _bgg.GetPixel(_fgaf, _ddgg)
			if _fgg {
				_eeeg += "\u0031"
			} else {
				_eeeg += "\u0030"
			}
		}
		_aec += _eeeg + "\u000a"
	}
	return _aec
}

func (_cefc *Bitmap) countPixels() int {
	var (
		_eef  int
		_dfca uint8
		_ggfb byte
		_bff  int
	)
	_bbe := _cefc.RowStride
	_daa := uint(_cefc.Width & 0x07)
	if _daa != 0 {
		_dfca = uint8((0xff << (8 - _daa)) & 0xff)
		_bbe--
	}
	for _gab := 0; _gab < _cefc.Height; _gab++ {
		for _bff = 0; _bff < _bbe; _bff++ {
			_ggfb = _cefc.Data[_gab*_cefc.RowStride+_bff]
			_eef += int(_cee[_ggfb])
		}
		if _daa != 0 {
			_eef += int(_cee[_cefc.Data[_gab*_cefc.RowStride+_bff]&_dfca])
		}
	}
	return _eef
}

func (_ebd *Bitmap) removeBorderGeneral(_fcaf, _gddf, _bbag, _geb int) (*Bitmap, error) {
	const _acce = "\u0072\u0065\u006d\u006fve\u0042\u006f\u0072\u0064\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _fcaf < 0 || _gddf < 0 || _bbag < 0 || _geb < 0 {
		return nil, _c.Error(_acce, "\u006e\u0065g\u0061\u0074\u0069\u0076\u0065\u0020\u0062\u0072\u006f\u0064\u0065\u0072\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u0076\u0061lu\u0065\u0073")
	}
	_dgb, _eebc := _ebd.Width, _ebd.Height
	_bdf := _dgb - _fcaf - _gddf
	_baae := _eebc - _bbag - _geb
	if _bdf <= 0 {
		return nil, _c.Errorf(_acce, "w\u0069\u0064\u0074\u0068: \u0025d\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u003e\u0020\u0030", _bdf)
	}
	if _baae <= 0 {
		return nil, _c.Errorf(_acce, "\u0068\u0065\u0069\u0067ht\u003a\u0020\u0025\u0064\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030", _baae)
	}
	_gbeb := New(_bdf, _baae)
	_gbeb.Color = _ebd.Color
	_bdga := _gbeb.RasterOperation(0, 0, _bdf, _baae, PixSrc, _ebd, _fcaf, _bbag)
	if _bdga != nil {
		return nil, _c.Wrap(_bdga, _acce, "")
	}
	return _gbeb, nil
}

const (
	ComponentConn Component = iota
	ComponentCharacters
	ComponentWords
)

type SizeSelection int

func _bdae(_cedg, _cbeea, _ffag *Bitmap) (*Bitmap, error) {
	const _edea = "\u0073\u0075\u0062\u0074\u0072\u0061\u0063\u0074"
	if _cbeea == nil {
		return nil, _c.Error(_edea, "'\u0073\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _ffag == nil {
		return nil, _c.Error(_edea, "'\u0073\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	var _fcc error
	switch {
	case _cedg == _cbeea:
		if _fcc = _cedg.RasterOperation(0, 0, _cbeea.Width, _cbeea.Height, PixNotSrcAndDst, _ffag, 0, 0); _fcc != nil {
			return nil, _c.Wrap(_fcc, _edea, "\u0064 \u003d\u003d\u0020\u0073\u0031")
		}
	case _cedg == _ffag:
		if _fcc = _cedg.RasterOperation(0, 0, _cbeea.Width, _cbeea.Height, PixNotSrcAndDst, _cbeea, 0, 0); _fcc != nil {
			return nil, _c.Wrap(_fcc, _edea, "\u0064 \u003d\u003d\u0020\u0073\u0032")
		}
	default:
		_cedg, _fcc = _gbfag(_cedg, _cbeea)
		if _fcc != nil {
			return nil, _c.Wrap(_fcc, _edea, "")
		}
		if _fcc = _cedg.RasterOperation(0, 0, _cbeea.Width, _cbeea.Height, PixNotSrcAndDst, _ffag, 0, 0); _fcc != nil {
			return nil, _c.Wrap(_fcc, _edea, "\u0064e\u0066\u0061\u0075\u006c\u0074")
		}
	}
	return _cedg, nil
}

func (_dfeg *ClassedPoints) validateIntSlice() error {
	const _dbbb = "\u0076\u0061l\u0069\u0064\u0061t\u0065\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065"
	for _, _bbbb := range _dfeg.IntSlice {
		if _bbbb >= (_dfeg.Points.Size()) {
			return _c.Errorf(_dbbb, "c\u006c\u0061\u0073\u0073\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006ci\u0064 \u0069\u006e\u0064\u0065x\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0064", _bbbb, _dfeg.Points.Size())
		}
	}
	return nil
}
func (_cdfc *BitmapsArray) AddBox(box *_g.Rectangle) { _cdfc.Boxes = append(_cdfc.Boxes, box) }
func (_edec Points) GetGeometry(i int) (_dgaf, _ecfd float32, _fdfd error) {
	if i > len(_edec)-1 {
		return 0, 0, _c.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_bceb := _edec[i]
	return _bceb.X, _bceb.Y, nil
}

func (_cgeg *Bitmaps) CountPixels() *_fa.NumSlice {
	_aeff := &_fa.NumSlice{}
	for _, _dgda := range _cgeg.Values {
		_aeff.AddInt(_dgda.CountPixels())
	}
	return _aeff
}

func _fcac(_bae, _ggbc int) int {
	if _bae < _ggbc {
		return _bae
	}
	return _ggbc
}

func (_eae *Bitmap) SetByte(index int, v byte) error {
	if index > len(_eae.Data)-1 || index < 0 {
		return _c.Errorf("\u0053e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", index)
	}
	_eae.Data[index] = v
	return nil
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

type byHeight Bitmaps

type BitmapsArray struct {
	Values []*Bitmaps
	Boxes  []*_g.Rectangle
}

func (_eaag *ClassedPoints) GroupByY() ([]*ClassedPoints, error) {
	const _dfab = "\u0043\u006c\u0061\u0073se\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0072\u006f\u0075\u0070\u0042y\u0059"
	if _gbeg := _eaag.validateIntSlice(); _gbeg != nil {
		return nil, _c.Wrap(_gbeg, _dfab, "")
	}
	if _eaag.IntSlice.Size() == 0 {
		return nil, _c.Error(_dfab, "\u004e\u006f\u0020\u0063la\u0073\u0073\u0065\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_eaag.SortByY()
	var (
		_beed []*ClassedPoints
		_agca int
	)
	_gbgd := -1
	var _bcae *ClassedPoints
	for _acde := 0; _acde < len(_eaag.IntSlice); _acde++ {
		_agca = int(_eaag.YAtIndex(_acde))
		if _agca != _gbgd {
			_bcae = &ClassedPoints{Points: _eaag.Points}
			_gbgd = _agca
			_beed = append(_beed, _bcae)
		}
		_bcae.IntSlice = append(_bcae.IntSlice, _eaag.IntSlice[_acde])
	}
	for _, _bcfd := range _beed {
		_bcfd.SortByX()
	}
	return _beed, nil
}

func _gada(_bafa *Bitmap, _faff, _gegg, _ebab, _cbg int, _ebfac RasterOperator, _ecbba *Bitmap, _fcbe, _bfac int) error {
	const _bdfad = "\u0072a\u0073t\u0065\u0072\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e"
	if _bafa == nil {
		return _c.Error(_bdfad, "\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _ebfac == PixDst {
		return nil
	}
	switch _ebfac {
	case PixClr, PixSet, PixNotDst:
		_cgba(_bafa, _faff, _gegg, _ebab, _cbg, _ebfac)
		return nil
	}
	if _ecbba == nil {
		_fce.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _c.Error(_bdfad, "\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _ebbc := _cede(_bafa, _faff, _gegg, _ebab, _cbg, _ebfac, _ecbba, _fcbe, _bfac); _ebbc != nil {
		return _c.Wrap(_ebbc, _bdfad, "")
	}
	return nil
}

func (_ebc *Bitmap) GetUnpaddedData() ([]byte, error) {
	_dcg := uint(_ebc.Width & 0x07)
	if _dcg == 0 {
		return _ebc.Data, nil
	}
	_bgfg := _ebc.Width * _ebc.Height
	if _bgfg%8 != 0 {
		_bgfg >>= 3
		_bgfg++
	} else {
		_bgfg >>= 3
	}
	_cbe := make([]byte, _bgfg)
	_gbde := _b.NewWriterMSB(_cbe)
	const _eff = "\u0047e\u0074U\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	for _gad := 0; _gad < _ebc.Height; _gad++ {
		for _cca := 0; _cca < _ebc.RowStride; _cca++ {
			_ace := _ebc.Data[_gad*_ebc.RowStride+_cca]
			if _cca != _ebc.RowStride-1 {
				_acdf := _gbde.WriteByte(_ace)
				if _acdf != nil {
					return nil, _c.Wrap(_acdf, _eff, "")
				}
				continue
			}
			for _gecf := uint(0); _gecf < _dcg; _gecf++ {
				_ebeb := _gbde.WriteBit(int(_ace >> (7 - _gecf) & 0x01))
				if _ebeb != nil {
					return nil, _c.Wrap(_ebeb, _eff, "")
				}
			}
		}
	}
	return _cbe, nil
}

func (_dbdb *Bitmap) AddBorder(borderSize, val int) (*Bitmap, error) {
	if borderSize == 0 {
		return _dbdb.Copy(), nil
	}
	_baa, _fcf := _dbdb.addBorderGeneral(borderSize, borderSize, borderSize, borderSize, val)
	if _fcf != nil {
		return nil, _c.Wrap(_fcf, "\u0041d\u0064\u0042\u006f\u0072\u0064\u0065r", "")
	}
	return _baa, nil
}

func _caec(_ggde *Bitmap) (_fcdgc *Bitmap, _gde int, _gbdga error) {
	const _gdb = "\u0042i\u0074\u006d\u0061\u0070.\u0077\u006f\u0072\u0064\u004da\u0073k\u0042y\u0044\u0069\u006c\u0061\u0074\u0069\u006fn"
	if _ggde == nil {
		return nil, 0, _c.Errorf(_gdb, "\u0027\u0073\u0027\u0020bi\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	var _agdd, _ffaa *Bitmap
	if _agdd, _gbdga = _gbfag(nil, _ggde); _gbdga != nil {
		return nil, 0, _c.Wrap(_gbdga, _gdb, "\u0063\u006f\u0070\u0079\u0020\u0027\u0073\u0027")
	}
	var (
		_cbdd        [13]int
		_fdaf, _addf int
	)
	_dff := 12
	_adgd := _fa.NewNumSlice(_dff + 1)
	_gca := _fa.NewNumSlice(_dff + 1)
	var _bgga *Boxes
	for _efgg := 0; _efgg <= _dff; _efgg++ {
		if _efgg == 0 {
			if _ffaa, _gbdga = _gbfag(nil, _agdd); _gbdga != nil {
				return nil, 0, _c.Wrap(_gbdga, _gdb, "\u0066i\u0072\u0073\u0074\u0020\u0062\u006d2")
			}
		} else {
			if _ffaa, _gbdga = _gbfc(_agdd, MorphProcess{Operation: MopDilation, Arguments: []int{2, 1}}); _gbdga != nil {
				return nil, 0, _c.Wrap(_gbdga, _gdb, "\u0064\u0069\u006ca\u0074\u0069\u006f\u006e\u0020\u0062\u006d\u0032")
			}
		}
		if _bgga, _gbdga = _ffaa.connComponentsBB(4); _gbdga != nil {
			return nil, 0, _c.Wrap(_gbdga, _gdb, "")
		}
		_cbdd[_efgg] = len(*_bgga)
		_adgd.AddInt(_cbdd[_efgg])
		switch _efgg {
		case 0:
			_fdaf = _cbdd[0]
		default:
			_addf = _cbdd[_efgg-1] - _cbdd[_efgg]
			_gca.AddInt(_addf)
		}
		_agdd = _ffaa
	}
	_cbfa := true
	_fffgg := 2
	var _gacd, _eecg int
	for _eeee := 1; _eeee < len(*_gca); _eeee++ {
		if _gacd, _gbdga = _adgd.GetInt(_eeee); _gbdga != nil {
			return nil, 0, _c.Wrap(_gbdga, _gdb, "\u0043\u0068\u0065\u0063ki\u006e\u0067\u0020\u0062\u0065\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0069o\u006e")
		}
		if _cbfa && _gacd < int(0.3*float32(_fdaf)) {
			_fffgg = _eeee + 1
			_cbfa = false
		}
		if _addf, _gbdga = _gca.GetInt(_eeee); _gbdga != nil {
			return nil, 0, _c.Wrap(_gbdga, _gdb, "\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006ea\u0044\u0069\u0066\u0066")
		}
		if _addf > _eecg {
			_eecg = _addf
		}
	}
	_dfgb := _ggde.XResolution
	if _dfgb == 0 {
		_dfgb = 150
	}
	if _dfgb > 110 {
		_fffgg++
	}
	if _fffgg < 2 {
		_fce.Log.Trace("J\u0042\u0049\u0047\u0032\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0042\u0065\u0073\u0074 \u0074\u006f\u0020\u006d\u0069\u006e\u0069\u006d\u0075\u006d a\u006c\u006c\u006fw\u0061b\u006c\u0065")
		_fffgg = 2
	}
	_gde = _fffgg + 1
	if _fcdgc, _gbdga = _gcecf(nil, _ggde, _fffgg+1, 1); _gbdga != nil {
		return nil, 0, _c.Wrap(_gbdga, _gdb, "\u0067\u0065\u0074\u0074in\u0067\u0020\u006d\u0061\u0073\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	return _fcdgc, _gde, nil
}

func (_bbda *Bitmap) AddBorderGeneral(left, right, top, bot int, val int) (*Bitmap, error) {
	return _bbda.addBorderGeneral(left, right, top, bot, val)
}

type CombinationOperator int

func _gcae(_bbeg, _cebg *Bitmap, _agee *Selection) (*Bitmap, error) {
	const _aafd = "\u0065\u0072\u006fd\u0065"
	var (
		_gbcda error
		_dbba  *Bitmap
	)
	_bbeg, _gbcda = _aecg(_bbeg, _cebg, _agee, &_dbba)
	if _gbcda != nil {
		return nil, _c.Wrap(_gbcda, _aafd, "")
	}
	if _gbcda = _bbeg.setAll(); _gbcda != nil {
		return nil, _c.Wrap(_gbcda, _aafd, "")
	}
	var _dffg SelectionValue
	for _ecba := 0; _ecba < _agee.Height; _ecba++ {
		for _gdfd := 0; _gdfd < _agee.Width; _gdfd++ {
			_dffg = _agee.Data[_ecba][_gdfd]
			if _dffg == SelHit {
				_gbcda = _gada(_bbeg, _agee.Cx-_gdfd, _agee.Cy-_ecba, _cebg.Width, _cebg.Height, PixSrcAndDst, _dbba, 0, 0)
				if _gbcda != nil {
					return nil, _c.Wrap(_gbcda, _aafd, "")
				}
			}
		}
	}
	if MorphBC == SymmetricMorphBC {
		return _bbeg, nil
	}
	_cedc, _befd, _dfdee, _gfeba := _agee.findMaxTranslations()
	if _cedc > 0 {
		if _gbcda = _bbeg.RasterOperation(0, 0, _cedc, _cebg.Height, PixClr, nil, 0, 0); _gbcda != nil {
			return nil, _c.Wrap(_gbcda, _aafd, "\u0078\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _dfdee > 0 {
		if _gbcda = _bbeg.RasterOperation(_cebg.Width-_dfdee, 0, _dfdee, _cebg.Height, PixClr, nil, 0, 0); _gbcda != nil {
			return nil, _c.Wrap(_gbcda, _aafd, "\u0078\u006e\u0020\u003e\u0020\u0030")
		}
	}
	if _befd > 0 {
		if _gbcda = _bbeg.RasterOperation(0, 0, _cebg.Width, _befd, PixClr, nil, 0, 0); _gbcda != nil {
			return nil, _c.Wrap(_gbcda, _aafd, "\u0079\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _gfeba > 0 {
		if _gbcda = _bbeg.RasterOperation(0, _cebg.Height-_gfeba, _cebg.Width, _gfeba, PixClr, nil, 0, 0); _gbcda != nil {
			return nil, _c.Wrap(_gbcda, _aafd, "\u0079\u006e\u0020\u003e\u0020\u0030")
		}
	}
	return _bbeg, nil
}

func _dbgc(_fcgg, _bbae *Bitmap, _acff, _bdbf, _deaa, _fcea, _gcf, _bdegc, _fbdd, _gge int, _bed CombinationOperator) error {
	var _adgf int
	_gcec := func() { _adgf++; _deaa += _bbae.RowStride; _fcea += _fcgg.RowStride; _gcf += _fcgg.RowStride }
	for _adgf = _acff; _adgf < _bdbf; _gcec() {
		var _ebfa uint16
		_fcfg := _deaa
		for _ggfd := _fcea; _ggfd <= _gcf; _ggfd++ {
			_bdgb, _cedge := _bbae.GetByte(_fcfg)
			if _cedge != nil {
				return _cedge
			}
			_cgd, _cedge := _fcgg.GetByte(_ggfd)
			if _cedge != nil {
				return _cedge
			}
			_ebfa = (_ebfa | uint16(_cgd)) << uint(_gge)
			_cgd = byte(_ebfa >> 8)
			if _ggfd == _gcf {
				_cgd = _aabd(uint(_bdegc), _cgd)
			}
			if _cedge = _bbae.SetByte(_fcfg, _fbfd(_bdgb, _cgd, _bed)); _cedge != nil {
				return _cedge
			}
			_fcfg++
			_ebfa <<= uint(_fbdd)
		}
	}
	return nil
}
func (_dfabb *Selection) setOrigin(_ffdf, _gdgb int) { _dfabb.Cy, _dfabb.Cx = _ffdf, _gdgb }
func (_bcge *ClassedPoints) GetIntYByClass(i int) (int, error) {
	const _ffec = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _bcge.IntSlice.Size() {
		return 0, _c.Errorf(_ffec, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_bcge.YAtIndex(i)), nil
}
func DilateBrick(d, s *Bitmap, hSize, vSize int) (*Bitmap, error)  { return _cfde(d, s, hSize, vSize) }
func Dilate(d *Bitmap, s *Bitmap, sel *Selection) (*Bitmap, error) { return _bag(d, s, sel) }
func _adcca(_cefa ...MorphProcess) (_fdgg error) {
	const _eceg = "v\u0065r\u0069\u0066\u0079\u004d\u006f\u0072\u0070\u0068P\u0072\u006f\u0063\u0065ss\u0065\u0073"
	var _egfd, _agbe int
	for _dcgg, _gabd := range _cefa {
		if _fdgg = _gabd.verify(_dcgg, &_egfd, &_agbe); _fdgg != nil {
			return _c.Wrap(_fdgg, _eceg, "")
		}
	}
	if _agbe != 0 && _egfd != 0 {
		return _c.Error(_eceg, "\u004d\u006f\u0072\u0070\u0068\u0020\u0073\u0065\u0071\u0075\u0065n\u0063\u0065\u0020\u002d\u0020\u0062\u006f\u0072d\u0065r\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u0065\u0074\u0020\u0072\u0065\u0064u\u0063\u0074\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0030")
	}
	return nil
}

func (_cdeg *Bitmap) connComponentsBitmapsBB(_dfde *Bitmaps, _feb int) (_gecee *Boxes, _fgag error) {
	const _dfg = "\u0063\u006f\u006enC\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0069\u0074\u006d\u0061\u0070\u0073\u0042\u0042"
	if _feb != 4 && _feb != 8 {
		return nil, _c.Error(_dfg, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _dfde == nil {
		return nil, _c.Error(_dfg, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if len(_dfde.Values) > 0 {
		return nil, _c.Error(_dfg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u006fn\u002d\u0065\u006d\u0070\u0074\u0079\u0020\u0042\u0069\u0074m\u0061\u0070\u0073")
	}
	if _cdeg.Zero() {
		return &Boxes{}, nil
	}
	var _caab, _egbb, _bddg, _dab *Bitmap
	_cdeg.setPadBits(0)
	if _caab, _fgag = _gbfag(nil, _cdeg); _fgag != nil {
		return nil, _c.Wrap(_fgag, _dfg, "\u0062\u006d\u0031")
	}
	if _egbb, _fgag = _gbfag(nil, _cdeg); _fgag != nil {
		return nil, _c.Wrap(_fgag, _dfg, "\u0062\u006d\u0032")
	}
	_bcg := &_fa.Stack{}
	_bcg.Aux = &_fa.Stack{}
	_gecee = &Boxes{}
	var (
		_fccg, _acfa int
		_fegg        _g.Point
		_fcgdc       bool
		_dgd         *_g.Rectangle
	)
	for {
		if _fegg, _fcgdc, _fgag = _caab.nextOnPixel(_fccg, _acfa); _fgag != nil {
			return nil, _c.Wrap(_fgag, _dfg, "")
		}
		if !_fcgdc {
			break
		}
		if _dgd, _fgag = _fdee(_caab, _bcg, _fegg.X, _fegg.Y, _feb); _fgag != nil {
			return nil, _c.Wrap(_fgag, _dfg, "")
		}
		if _fgag = _gecee.Add(_dgd); _fgag != nil {
			return nil, _c.Wrap(_fgag, _dfg, "")
		}
		if _bddg, _fgag = _caab.clipRectangle(_dgd, nil); _fgag != nil {
			return nil, _c.Wrap(_fgag, _dfg, "\u0062\u006d\u0033")
		}
		if _dab, _fgag = _egbb.clipRectangle(_dgd, nil); _fgag != nil {
			return nil, _c.Wrap(_fgag, _dfg, "\u0062\u006d\u0034")
		}
		if _, _fgag = _geg(_bddg, _bddg, _dab); _fgag != nil {
			return nil, _c.Wrap(_fgag, _dfg, "\u0062m\u0033\u0020\u005e\u0020\u0062\u006d4")
		}
		if _fgag = _egbb.RasterOperation(_dgd.Min.X, _dgd.Min.Y, _dgd.Dx(), _dgd.Dy(), PixSrcXorDst, _bddg, 0, 0); _fgag != nil {
			return nil, _c.Wrap(_fgag, _dfg, "\u0062\u006d\u0032\u0020\u002d\u0058\u004f\u0052\u002d>\u0020\u0062\u006d\u0033")
		}
		_dfde.AddBitmap(_bddg)
		_fccg = _fegg.X
		_acfa = _fegg.Y
	}
	_dfde.Boxes = *_gecee
	return _gecee, nil
}
func MakePixelSumTab8() []int { return _ccda() }
func (_ecc *Bitmap) setTwoBytes(_fafd int, _fcbd uint16) error {
	if _fafd+1 > len(_ecc.Data)-1 {
		return _c.Errorf("s\u0065\u0074\u0054\u0077\u006f\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _fafd)
	}
	_ecc.Data[_fafd] = byte((_fcbd & 0xff00) >> 8)
	_ecc.Data[_fafd+1] = byte(_fcbd & 0xff)
	return nil
}

type MorphProcess struct {
	Operation MorphOperation
	Arguments []int
}

func _fcdga(_edce, _dfb int) int {
	if _edce > _dfb {
		return _edce
	}
	return _dfb
}

func (_fggd MorphProcess) getWidthHeight() (_defbc, _fcee int) {
	return _fggd.Arguments[0], _fggd.Arguments[1]
}

func CorrelationScore(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_eag float64, _bacdg error) {
	const _baec = "\u0063\u006fr\u0072\u0065\u006ca\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065"
	if bm1 == nil || bm2 == nil {
		return 0, _c.Error(_baec, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0062\u0069\u0074ma\u0070\u0073")
	}
	if tab == nil {
		return 0, _c.Error(_baec, "\u0027\u0074\u0061\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if area1 <= 0 || area2 <= 0 {
		return 0, _c.Error(_baec, "\u0061\u0072\u0065\u0061s\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067r\u0065a\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
	}
	_ffg, _bcfcd := bm1.Width, bm1.Height
	_dfbb, _cfgd := bm2.Width, bm2.Height
	_abab := _acef(_ffg - _dfbb)
	if _abab > maxDiffW {
		return 0, nil
	}
	_edae := _acef(_bcfcd - _cfgd)
	if _edae > maxDiffH {
		return 0, nil
	}
	var _bcga, _dddc int
	if delX >= 0 {
		_bcga = int(delX + 0.5)
	} else {
		_bcga = int(delX - 0.5)
	}
	if delY >= 0 {
		_dddc = int(delY + 0.5)
	} else {
		_dddc = int(delY - 0.5)
	}
	_bfcd := _fcdga(_dddc, 0)
	_caaa := _fcac(_cfgd+_dddc, _bcfcd)
	_cega := bm1.RowStride * _bfcd
	_cfc := bm2.RowStride * (_bfcd - _dddc)
	_fcbg := _fcdga(_bcga, 0)
	_dgag := _fcac(_dfbb+_bcga, _ffg)
	_bddc := bm2.RowStride
	var _ade, _adfae int
	if _bcga >= 8 {
		_ade = _bcga >> 3
		_cega += _ade
		_fcbg -= _ade << 3
		_dgag -= _ade << 3
		_bcga &= 7
	} else if _bcga <= -8 {
		_adfae = -((_bcga + 7) >> 3)
		_cfc += _adfae
		_bddc -= _adfae
		_bcga += _adfae << 3
	}
	if _fcbg >= _dgag || _bfcd >= _caaa {
		return 0, nil
	}
	_afcg := (_dgag + 7) >> 3
	var (
		_eefc, _defbg, _gdea byte
		_accd, _edbf, _cgdc  int
	)
	switch {
	case _bcga == 0:
		for _cgdc = _bfcd; _cgdc < _caaa; _cgdc, _cega, _cfc = _cgdc+1, _cega+bm1.RowStride, _cfc+bm2.RowStride {
			for _edbf = 0; _edbf < _afcg; _edbf++ {
				_gdea = bm1.Data[_cega+_edbf] & bm2.Data[_cfc+_edbf]
				_accd += tab[_gdea]
			}
		}
	case _bcga > 0:
		if _bddc < _afcg {
			for _cgdc = _bfcd; _cgdc < _caaa; _cgdc, _cega, _cfc = _cgdc+1, _cega+bm1.RowStride, _cfc+bm2.RowStride {
				_eefc, _defbg = bm1.Data[_cega], bm2.Data[_cfc]>>uint(_bcga)
				_gdea = _eefc & _defbg
				_accd += tab[_gdea]
				for _edbf = 1; _edbf < _bddc; _edbf++ {
					_eefc, _defbg = bm1.Data[_cega+_edbf], (bm2.Data[_cfc+_edbf]>>uint(_bcga))|(bm2.Data[_cfc+_edbf-1]<<uint(8-_bcga))
					_gdea = _eefc & _defbg
					_accd += tab[_gdea]
				}
				_eefc = bm1.Data[_cega+_edbf]
				_defbg = bm2.Data[_cfc+_edbf-1] << uint(8-_bcga)
				_gdea = _eefc & _defbg
				_accd += tab[_gdea]
			}
		} else {
			for _cgdc = _bfcd; _cgdc < _caaa; _cgdc, _cega, _cfc = _cgdc+1, _cega+bm1.RowStride, _cfc+bm2.RowStride {
				_eefc, _defbg = bm1.Data[_cega], bm2.Data[_cfc]>>uint(_bcga)
				_gdea = _eefc & _defbg
				_accd += tab[_gdea]
				for _edbf = 1; _edbf < _afcg; _edbf++ {
					_eefc = bm1.Data[_cega+_edbf]
					_defbg = (bm2.Data[_cfc+_edbf] >> uint(_bcga)) | (bm2.Data[_cfc+_edbf-1] << uint(8-_bcga))
					_gdea = _eefc & _defbg
					_accd += tab[_gdea]
				}
			}
		}
	default:
		if _afcg < _bddc {
			for _cgdc = _bfcd; _cgdc < _caaa; _cgdc, _cega, _cfc = _cgdc+1, _cega+bm1.RowStride, _cfc+bm2.RowStride {
				for _edbf = 0; _edbf < _afcg; _edbf++ {
					_eefc = bm1.Data[_cega+_edbf]
					_defbg = bm2.Data[_cfc+_edbf] << uint(-_bcga)
					_defbg |= bm2.Data[_cfc+_edbf+1] >> uint(8+_bcga)
					_gdea = _eefc & _defbg
					_accd += tab[_gdea]
				}
			}
		} else {
			for _cgdc = _bfcd; _cgdc < _caaa; _cgdc, _cega, _cfc = _cgdc+1, _cega+bm1.RowStride, _cfc+bm2.RowStride {
				for _edbf = 0; _edbf < _afcg-1; _edbf++ {
					_eefc = bm1.Data[_cega+_edbf]
					_defbg = bm2.Data[_cfc+_edbf] << uint(-_bcga)
					_defbg |= bm2.Data[_cfc+_edbf+1] >> uint(8+_bcga)
					_gdea = _eefc & _defbg
					_accd += tab[_gdea]
				}
				_eefc = bm1.Data[_cega+_edbf]
				_defbg = bm2.Data[_cfc+_edbf] << uint(-_bcga)
				_gdea = _eefc & _defbg
				_accd += tab[_gdea]
			}
		}
	}
	_eag = float64(_accd) * float64(_accd) / (float64(area1) * float64(area2))
	return _eag, nil
}

func _gbfag(_bbaa, _gffb *Bitmap) (*Bitmap, error) {
	if _gffb == nil {
		return nil, _c.Error("\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _gffb == _bbaa {
		return _bbaa, nil
	}
	if _bbaa == nil {
		_bbaa = _gffb.createTemplate()
		copy(_bbaa.Data, _gffb.Data)
		return _bbaa, nil
	}
	_aeba := _bbaa.resizeImageData(_gffb)
	if _aeba != nil {
		return nil, _c.Wrap(_aeba, "\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "")
	}
	_bbaa.Text = _gffb.Text
	copy(_bbaa.Data, _gffb.Data)
	return _bbaa, nil
}

func (_edca *Bitmap) SizesEqual(s *Bitmap) bool {
	if _edca == s {
		return true
	}
	if _edca.Width != s.Width || _edca.Height != s.Height {
		return false
	}
	return true
}

func _ceg() (_cfd [256]uint64) {
	for _decc := 0; _decc < 256; _decc++ {
		if _decc&0x01 != 0 {
			_cfd[_decc] |= 0xff
		}
		if _decc&0x02 != 0 {
			_cfd[_decc] |= 0xff00
		}
		if _decc&0x04 != 0 {
			_cfd[_decc] |= 0xff0000
		}
		if _decc&0x08 != 0 {
			_cfd[_decc] |= 0xff000000
		}
		if _decc&0x10 != 0 {
			_cfd[_decc] |= 0xff00000000
		}
		if _decc&0x20 != 0 {
			_cfd[_decc] |= 0xff0000000000
		}
		if _decc&0x40 != 0 {
			_cfd[_decc] |= 0xff000000000000
		}
		if _decc&0x80 != 0 {
			_cfd[_decc] |= 0xff00000000000000
		}
	}
	return _cfd
}

func (_bgcb *Bitmaps) SelectByIndexes(idx []int) (*Bitmaps, error) {
	const _cacg = "B\u0069\u0074\u006d\u0061\u0070\u0073.\u0053\u006f\u0072\u0074\u0049\u006e\u0064\u0065\u0078e\u0073\u0042\u0079H\u0065i\u0067\u0068\u0074"
	_bffd, _ceecfd := _bgcb.selectByIndexes(idx)
	if _ceecfd != nil {
		return nil, _c.Wrap(_ceecfd, _cacg, "")
	}
	return _bffd, nil
}
func (_ggcc *Bitmaps) AddBitmap(bm *Bitmap) { _ggcc.Values = append(_ggcc.Values, bm) }
func (_dcd *Bitmap) inverseData() {
	if _bad := _dcd.RasterOperation(0, 0, _dcd.Width, _dcd.Height, PixNotDst, nil, 0, 0); _bad != nil {
		_fce.Log.Debug("\u0049n\u0076\u0065\u0072\u0073e\u0020\u0064\u0061\u0074\u0061 \u0066a\u0069l\u0065\u0064\u003a\u0020\u0027\u0025\u0076'", _bad)
	}
	if _dcd.Color == Chocolate {
		_dcd.Color = Vanilla
	} else {
		_dcd.Color = Chocolate
	}
}

func (_dbdba *Bitmap) thresholdPixelSum(_fcgb int) bool {
	var (
		_gfgf  int
		_aga   uint8
		_aegca byte
		_bdb   int
	)
	_gba := _dbdba.RowStride
	_agag := uint(_dbdba.Width & 0x07)
	if _agag != 0 {
		_aga = uint8((0xff << (8 - _agag)) & 0xff)
		_gba--
	}
	for _bbef := 0; _bbef < _dbdba.Height; _bbef++ {
		for _bdb = 0; _bdb < _gba; _bdb++ {
			_aegca = _dbdba.Data[_bbef*_dbdba.RowStride+_bdb]
			_gfgf += int(_cee[_aegca])
		}
		if _agag != 0 {
			_aegca = _dbdba.Data[_bbef*_dbdba.RowStride+_bdb] & _aga
			_gfgf += int(_cee[_aegca])
		}
		if _gfgf > _fcgb {
			return true
		}
	}
	return false
}

type Point struct{ X, Y float32 }

func _cgba(_aebe *Bitmap, _ebfgf, _dbec, _gedc, _deaeb int, _bfgb RasterOperator) {
	if _ebfgf < 0 {
		_gedc += _ebfgf
		_ebfgf = 0
	}
	_bfbe := _ebfgf + _gedc - _aebe.Width
	if _bfbe > 0 {
		_gedc -= _bfbe
	}
	if _dbec < 0 {
		_deaeb += _dbec
		_dbec = 0
	}
	_gbea := _dbec + _deaeb - _aebe.Height
	if _gbea > 0 {
		_deaeb -= _gbea
	}
	if _gedc <= 0 || _deaeb <= 0 {
		return
	}
	if (_ebfgf & 7) == 0 {
		_cgde(_aebe, _ebfgf, _dbec, _gedc, _deaeb, _bfgb)
	} else {
		_gfbcb(_aebe, _ebfgf, _dbec, _gedc, _deaeb, _bfgb)
	}
}

func (_bbbd *Bitmap) addPadBits() (_ccdb error) {
	const _cfea = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0061\u0064\u0064\u0050\u0061d\u0042\u0069\u0074\u0073"
	_bfd := _bbbd.Width % 8
	if _bfd == 0 {
		return nil
	}
	_deceef := _bbbd.Width / 8
	_agc := _b.NewReader(_bbbd.Data)
	_cbee := make([]byte, _bbbd.Height*_bbbd.RowStride)
	_aaa := _b.NewWriterMSB(_cbee)
	_cddd := make([]byte, _deceef)
	var (
		_aaad int
		_faaf uint64
	)
	for _aaad = 0; _aaad < _bbbd.Height; _aaad++ {
		if _, _ccdb = _agc.Read(_cddd); _ccdb != nil {
			return _c.Wrap(_ccdb, _cfea, "\u0066u\u006c\u006c\u0020\u0062\u0079\u0074e")
		}
		if _, _ccdb = _aaa.Write(_cddd); _ccdb != nil {
			return _c.Wrap(_ccdb, _cfea, "\u0066\u0075\u006c\u006c\u0020\u0062\u0079\u0074\u0065\u0073")
		}
		if _faaf, _ccdb = _agc.ReadBits(byte(_bfd)); _ccdb != nil {
			return _c.Wrap(_ccdb, _cfea, "\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0069\u0074\u0073")
		}
		if _ccdb = _aaa.WriteByte(byte(_faaf) << uint(8-_bfd)); _ccdb != nil {
			return _c.Wrap(_ccdb, _cfea, "\u006ca\u0073\u0074\u0020\u0062\u0079\u0074e")
		}
	}
	_bbbd.Data = _aaa.Data()
	return nil
}

func (_cdb *Bitmap) setEightPartlyBytes(_bbdf, _fdcc int, _ggda uint64) (_gabc error) {
	var (
		_ffdd byte
		_abfa int
	)
	const _fdb = "\u0073\u0065\u0074\u0045ig\u0068\u0074\u0050\u0061\u0072\u0074\u006c\u0079\u0042\u0079\u0074\u0065\u0073"
	for _gfge := 1; _gfge <= _fdcc; _gfge++ {
		_abfa = 64 - _gfge*8
		_ffdd = byte(_ggda >> uint(_abfa) & 0xff)
		_fce.Log.Trace("\u0074\u0065\u006d\u0070\u003a\u0020\u0025\u0030\u0038\u0062\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a %\u0064,\u0020\u0069\u0064\u0078\u003a\u0020\u0025\u0064\u002c\u0020\u0066\u0075l\u006c\u0042\u0079\u0074\u0065\u0073\u004e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c \u0073\u0068\u0069\u0066\u0074\u003a\u0020\u0025\u0064", _ffdd, _bbdf, _bbdf+_gfge-1, _fdcc, _abfa)
		if _gabc = _cdb.SetByte(_bbdf+_gfge-1, _ffdd); _gabc != nil {
			return _c.Wrap(_gabc, _fdb, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
		}
	}
	_addb := _cdb.RowStride*8 - _cdb.Width
	if _addb == 0 {
		return nil
	}
	_abfa -= 8
	_ffdd = byte(_ggda>>uint(_abfa)&0xff) << uint(_addb)
	if _gabc = _cdb.SetByte(_bbdf+_fdcc, _ffdd); _gabc != nil {
		return _c.Wrap(_gabc, _fdb, "\u0070\u0061\u0064\u0064\u0065\u0064")
	}
	return nil
}

type BoundaryCondition int

func Centroids(bms []*Bitmap) (*Points, error) {
	_bcag := make([]Point, len(bms))
	_egce := _cadad()
	_cggb := _ccda()
	var _acgc error
	for _begcd, _bggb := range bms {
		_bcag[_begcd], _acgc = _bggb.centroid(_egce, _cggb)
		if _acgc != nil {
			return nil, _acgc
		}
	}
	_fcga := Points(_bcag)
	return &_fcga, nil
}

func _ccda() []int {
	_bbge := make([]int, 256)
	for _cfaaf := 0; _cfaaf <= 0xff; _cfaaf++ {
		_ecffe := byte(_cfaaf)
		_bbge[_ecffe] = int(_ecffe&0x1) + (int(_ecffe>>1) & 0x1) + (int(_ecffe>>2) & 0x1) + (int(_ecffe>>3) & 0x1) + (int(_ecffe>>4) & 0x1) + (int(_ecffe>>5) & 0x1) + (int(_ecffe>>6) & 0x1) + (int(_ecffe>>7) & 0x1)
	}
	return _bbge
}

func _ga() (_ggc [256]uint16) {
	for _dec := 0; _dec < 256; _dec++ {
		if _dec&0x01 != 0 {
			_ggc[_dec] |= 0x3
		}
		if _dec&0x02 != 0 {
			_ggc[_dec] |= 0xc
		}
		if _dec&0x04 != 0 {
			_ggc[_dec] |= 0x30
		}
		if _dec&0x08 != 0 {
			_ggc[_dec] |= 0xc0
		}
		if _dec&0x10 != 0 {
			_ggc[_dec] |= 0x300
		}
		if _dec&0x20 != 0 {
			_ggc[_dec] |= 0xc00
		}
		if _dec&0x40 != 0 {
			_ggc[_dec] |= 0x3000
		}
		if _dec&0x80 != 0 {
			_ggc[_dec] |= 0xc000
		}
	}
	return _ggc
}
func (_bbea *byHeight) Less(i, j int) bool { return _bbea.Values[i].Height < _bbea.Values[j].Height }
func (_gag *Bitmap) InverseData()          { _gag.inverseData() }
func (_decd *Bitmap) ThresholdPixelSum(thresh int, tab8 []int) (_aca bool, _gcdb error) {
	const _bec = "\u0042i\u0074\u006d\u0061\u0070\u002e\u0054\u0068\u0072\u0065\u0073\u0068o\u006c\u0064\u0050\u0069\u0078\u0065\u006c\u0053\u0075\u006d"
	if tab8 == nil {
		tab8 = _ccda()
	}
	_beb := _decd.Width >> 3
	_efcd := _decd.Width & 7
	_aebg := byte(0xff << uint(8-_efcd))
	var (
		_gbba, _gece, _acag, _dgf int
		_faca                     byte
	)
	for _gbba = 0; _gbba < _decd.Height; _gbba++ {
		_acag = _decd.RowStride * _gbba
		for _gece = 0; _gece < _beb; _gece++ {
			_faca, _gcdb = _decd.GetByte(_acag + _gece)
			if _gcdb != nil {
				return false, _c.Wrap(_gcdb, _bec, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
			}
			_dgf += tab8[_faca]
		}
		if _efcd != 0 {
			_faca, _gcdb = _decd.GetByte(_acag + _gece)
			if _gcdb != nil {
				return false, _c.Wrap(_gcdb, _bec, "p\u0061\u0072\u0074\u0069\u0061\u006c\u0042\u0079\u0074\u0065")
			}
			_faca &= _aebg
			_dgf += tab8[_faca]
		}
		if _dgf > thresh {
			return true, nil
		}
	}
	return _aca, nil
}

func (_fagd *BitmapsArray) GetBox(i int) (*_g.Rectangle, error) {
	const _fagda = "\u0042\u0069\u0074\u006dap\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u006f\u0078"
	if _fagd == nil {
		return nil, _c.Error(_fagda, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_fagd.Boxes)-1 {
		return nil, _c.Errorf(_fagda, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _fagd.Boxes[i], nil
}

func _faeg(_gae, _bcac *Bitmap, _gceb int, _eab []byte, _fag int) (_afda error) {
	const _edf = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0033"
	var (
		_fgc, _bcca, _eeef, _ab, _ece, _dga, _fdg, _edg int
		_abb, _cdc, _ddac, _caf                         uint32
		_adc, _fgda                                     byte
		_fff                                            uint16
	)
	_gbf := make([]byte, 4)
	_efc := make([]byte, 4)
	for _eeef = 0; _eeef < _gae.Height-1; _eeef, _ab = _eeef+2, _ab+1 {
		_fgc = _eeef * _gae.RowStride
		_bcca = _ab * _bcac.RowStride
		for _ece, _dga = 0, 0; _ece < _fag; _ece, _dga = _ece+4, _dga+1 {
			for _fdg = 0; _fdg < 4; _fdg++ {
				_edg = _fgc + _ece + _fdg
				if _edg <= len(_gae.Data)-1 && _edg < _fgc+_gae.RowStride {
					_gbf[_fdg] = _gae.Data[_edg]
				} else {
					_gbf[_fdg] = 0x00
				}
				_edg = _fgc + _gae.RowStride + _ece + _fdg
				if _edg <= len(_gae.Data)-1 && _edg < _fgc+(2*_gae.RowStride) {
					_efc[_fdg] = _gae.Data[_edg]
				} else {
					_efc[_fdg] = 0x00
				}
			}
			_abb = _a.BigEndian.Uint32(_gbf)
			_cdc = _a.BigEndian.Uint32(_efc)
			_ddac = _abb & _cdc
			_ddac |= _ddac << 1
			_caf = _abb | _cdc
			_caf &= _caf << 1
			_cdc = _ddac & _caf
			_cdc &= 0xaaaaaaaa
			_abb = _cdc | (_cdc << 7)
			_adc = byte(_abb >> 24)
			_fgda = byte((_abb >> 8) & 0xff)
			_edg = _bcca + _dga
			if _edg+1 == len(_bcac.Data)-1 || _edg+1 >= _bcca+_bcac.RowStride {
				if _afda = _bcac.SetByte(_edg, _eab[_adc]); _afda != nil {
					return _c.Wrapf(_afda, _edf, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _edg)
				}
			} else {
				_fff = (uint16(_eab[_adc]) << 8) | uint16(_eab[_fgda])
				if _afda = _bcac.setTwoBytes(_edg, _fff); _afda != nil {
					return _c.Wrapf(_afda, _edf, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _edg)
				}
				_dga++
			}
		}
	}
	return nil
}

func (_dfac *Bitmap) addBorderGeneral(_aae, _ged, _gaa, _bdg int, _cfe int) (*Bitmap, error) {
	const _egcb = "\u0061\u0064d\u0042\u006f\u0072d\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _aae < 0 || _ged < 0 || _gaa < 0 || _bdg < 0 {
		return nil, _c.Error(_egcb, "n\u0065\u0067\u0061\u0074iv\u0065 \u0062\u006f\u0072\u0064\u0065r\u0020\u0061\u0064\u0064\u0065\u0064")
	}
	_bgfge, _adg := _dfac.Width, _dfac.Height
	_cbd := _bgfge + _aae + _ged
	_bbb := _adg + _gaa + _bdg
	_bdeb := New(_cbd, _bbb)
	_bdeb.Color = _dfac.Color
	_ebb := PixClr
	if _cfe > 0 {
		_ebb = PixSet
	}
	_faa := _bdeb.RasterOperation(0, 0, _aae, _bbb, _ebb, nil, 0, 0)
	if _faa != nil {
		return nil, _c.Wrap(_faa, _egcb, "\u006c\u0065\u0066\u0074")
	}
	_faa = _bdeb.RasterOperation(_cbd-_ged, 0, _ged, _bbb, _ebb, nil, 0, 0)
	if _faa != nil {
		return nil, _c.Wrap(_faa, _egcb, "\u0072\u0069\u0067h\u0074")
	}
	_faa = _bdeb.RasterOperation(0, 0, _cbd, _gaa, _ebb, nil, 0, 0)
	if _faa != nil {
		return nil, _c.Wrap(_faa, _egcb, "\u0074\u006f\u0070")
	}
	_faa = _bdeb.RasterOperation(0, _bbb-_bdg, _cbd, _bdg, _ebb, nil, 0, 0)
	if _faa != nil {
		return nil, _c.Wrap(_faa, _egcb, "\u0062\u006f\u0074\u0074\u006f\u006d")
	}
	_faa = _bdeb.RasterOperation(_aae, _gaa, _bgfge, _adg, PixSrc, _dfac, 0, 0)
	if _faa != nil {
		return nil, _c.Wrap(_faa, _egcb, "\u0063\u006f\u0070\u0079")
	}
	return _bdeb, nil
}

func RasterOperation(dest *Bitmap, dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _gada(dest, dx, dy, dw, dh, op, src, sx, sy)
}

func (_bdac *Bitmap) nextOnPixelLow(_ege, _fgcb, _bba, _bgbc, _gdf int) (_fgf _g.Point, _ddbf bool, _fcae error) {
	const _bdgg = "B\u0069\u0074\u006d\u0061p.\u006ee\u0078\u0074\u004f\u006e\u0050i\u0078\u0065\u006c\u004c\u006f\u0077"
	var (
		_fegf int
		_gafb byte
	)
	_ede := _gdf * _bba
	_gdd := _ede + (_bgbc / 8)
	if _gafb, _fcae = _bdac.GetByte(_gdd); _fcae != nil {
		return _fgf, false, _c.Wrap(_fcae, _bdgg, "\u0078\u0053\u0074\u0061\u0072\u0074\u0020\u0061\u006e\u0064 \u0079\u0053\u0074\u0061\u0072\u0074\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	if _gafb != 0 {
		_gaag := _bgbc - (_bgbc % 8) + 7
		for _fegf = _bgbc; _fegf <= _gaag && _fegf < _ege; _fegf++ {
			if _bdac.GetPixel(_fegf, _gdf) {
				_fgf.X = _fegf
				_fgf.Y = _gdf
				return _fgf, true, nil
			}
		}
	}
	_cfga := (_bgbc / 8) + 1
	_fegf = 8 * _cfga
	var _bebf int
	for _gdd = _ede + _cfga; _fegf < _ege; _gdd, _fegf = _gdd+1, _fegf+8 {
		if _gafb, _fcae = _bdac.GetByte(_gdd); _fcae != nil {
			return _fgf, false, _c.Wrap(_fcae, _bdgg, "r\u0065\u0073\u0074\u0020of\u0020t\u0068\u0065\u0020\u006c\u0069n\u0065\u0020\u0062\u0079\u0074\u0065")
		}
		if _gafb == 0 {
			continue
		}
		for _bebf = 0; _bebf < 8 && _fegf < _ege; _bebf, _fegf = _bebf+1, _fegf+1 {
			if _bdac.GetPixel(_fegf, _gdf) {
				_fgf.X = _fegf
				_fgf.Y = _gdf
				return _fgf, true, nil
			}
		}
	}
	for _ccg := _gdf + 1; _ccg < _fgcb; _ccg++ {
		_ede = _ccg * _bba
		for _gdd, _fegf = _ede, 0; _fegf < _ege; _gdd, _fegf = _gdd+1, _fegf+8 {
			if _gafb, _fcae = _bdac.GetByte(_gdd); _fcae != nil {
				return _fgf, false, _c.Wrap(_fcae, _bdgg, "\u0066o\u006cl\u006f\u0077\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u0073")
			}
			if _gafb == 0 {
				continue
			}
			for _bebf = 0; _bebf < 8 && _fegf < _ege; _bebf, _fegf = _bebf+1, _fegf+1 {
				if _bdac.GetPixel(_fegf, _ccg) {
					_fgf.X = _fegf
					_fgf.Y = _ccg
					return _fgf, true, nil
				}
			}
		}
	}
	return _fgf, false, nil
}

func New(width, height int) *Bitmap {
	_gbcd := _aeg(width, height)
	_gbcd.Data = make([]byte, height*_gbcd.RowStride)
	return _gbcd
}

func (_bfga *BitmapsArray) GetBitmaps(i int) (*Bitmaps, error) {
	const _bcaeg = "\u0042\u0069\u0074ma\u0070\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u0069\u0074\u006d\u0061\u0070\u0073"
	if _bfga == nil {
		return nil, _c.Error(_bcaeg, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_bfga.Values)-1 {
		return nil, _c.Errorf(_bcaeg, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _bfga.Values[i], nil
}

func _dge(_abfd, _ebef *Bitmap, _gaea, _cedgb int) (*Bitmap, error) {
	const _ebge = "\u0065\u0072\u006f\u0064\u0065\u0042\u0072\u0069\u0063\u006b"
	if _ebef == nil {
		return nil, _c.Error(_ebge, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _gaea < 1 || _cedgb < 1 {
		return nil, _c.Error(_ebge, "\u0068\u0073\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069\u007a\u0065\u0020\u0061\u0072e\u0020\u006e\u006f\u0074\u0020\u0067\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006fr\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _gaea == 1 && _cedgb == 1 {
		_cbff, _caabc := _gbfag(_abfd, _ebef)
		if _caabc != nil {
			return nil, _c.Wrap(_caabc, _ebge, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _cbff, nil
	}
	if _gaea == 1 || _cedgb == 1 {
		_eea := SelCreateBrick(_cedgb, _gaea, _cedgb/2, _gaea/2, SelHit)
		_edceb, _babbd := _gcae(_abfd, _ebef, _eea)
		if _babbd != nil {
			return nil, _c.Wrap(_babbd, _ebge, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _edceb, nil
	}
	_gaeae := SelCreateBrick(1, _gaea, 0, _gaea/2, SelHit)
	_fddg := SelCreateBrick(_cedgb, 1, _cedgb/2, 0, SelHit)
	_gda, _gbaf := _gcae(nil, _ebef, _gaeae)
	if _gbaf != nil {
		return nil, _c.Wrap(_gbaf, _ebge, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_abfd, _gbaf = _gcae(_abfd, _gda, _fddg)
	if _gbaf != nil {
		return nil, _c.Wrap(_gbaf, _ebge, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _abfd, nil
}

func ClipBoxToRectangle(box *_g.Rectangle, wi, hi int) (_cgcc *_g.Rectangle, _bacd error) {
	const _gcea = "\u0043l\u0069p\u0042\u006f\u0078\u0054\u006fR\u0065\u0063t\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, _c.Error(_gcea, "\u0027\u0062\u006f\u0078\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if box.Min.X >= wi || box.Min.Y >= hi || box.Max.X <= 0 || box.Max.Y <= 0 {
		return nil, _c.Error(_gcea, "\u0027\u0062\u006fx'\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065")
	}
	_gcb := *box
	_cgcc = &_gcb
	if _cgcc.Min.X < 0 {
		_cgcc.Max.X += _cgcc.Min.X
		_cgcc.Min.X = 0
	}
	if _cgcc.Min.Y < 0 {
		_cgcc.Max.Y += _cgcc.Min.Y
		_cgcc.Min.Y = 0
	}
	if _cgcc.Max.X > wi {
		_cgcc.Max.X = wi
	}
	if _cgcc.Max.Y > hi {
		_cgcc.Max.Y = hi
	}
	return _cgcc, nil
}

func _ddgge(_feff, _gfdb *Bitmap, _bcdg *Selection) (*Bitmap, error) {
	const _gbbb = "\u006f\u0070\u0065\u006e"
	var _afaa error
	_feff, _afaa = _bbegg(_feff, _gfdb, _bcdg)
	if _afaa != nil {
		return nil, _c.Wrap(_afaa, _gbbb, "")
	}
	_ddeb, _afaa := _gcae(nil, _gfdb, _bcdg)
	if _afaa != nil {
		return nil, _c.Wrap(_afaa, _gbbb, "")
	}
	_, _afaa = _bag(_feff, _ddeb, _bcdg)
	if _afaa != nil {
		return nil, _c.Wrap(_afaa, _gbbb, "")
	}
	return _feff, nil
}

func (_ecf *Boxes) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_adae *Boxes, _bgggc error) {
	const _fcded = "\u0042o\u0078e\u0073\u002e\u0053\u0065\u006ce\u0063\u0074B\u0079\u0053\u0069\u007a\u0065"
	if _ecf == nil {
		return nil, _c.Error(_fcded, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_ecf) == 0 {
		return _ecf, nil
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _c.Errorf(_fcded, "\u0069\u006e\u0076al\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE:
	default:
		return nil, _c.Errorf(_fcded, "i\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020t\u0079\u0070\u0065:\u0020'\u0025\u0064\u0027", tp)
	}
	_bgec := _ecf.makeSizeIndicator(width, height, tp, relation)
	_cag, _bgggc := _ecf.selectWithIndicator(_bgec)
	if _bgggc != nil {
		return nil, _c.Wrap(_bgggc, _fcded, "")
	}
	return _cag, nil
}
func (_bged *byWidth) Len() int { return len(_bged.Values) }
func _ebbf(_acbf *Bitmap, _gdac int) (*Bitmap, error) {
	const _gaab = "\u0065x\u0070a\u006e\u0064\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0065"
	if _acbf == nil {
		return nil, _c.Error(_gaab, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _gdac <= 0 {
		return nil, _c.Error(_gaab, "i\u006e\u0076\u0061\u006cid\u0020f\u0061\u0063\u0074\u006f\u0072 \u002d\u0020\u003c\u003d\u0020\u0030")
	}
	if _gdac == 1 {
		_gfca, _agfa := _gbfag(nil, _acbf)
		if _agfa != nil {
			return nil, _c.Wrap(_agfa, _gaab, "\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u003d\u0020\u0031")
		}
		return _gfca, nil
	}
	_bbdab, _eefcb := _edb(_acbf, _gdac, _gdac)
	if _eefcb != nil {
		return nil, _c.Wrap(_eefcb, _gaab, "")
	}
	return _bbdab, nil
}

func _gcda(_bbbec *Bitmap, _cgbgd, _cebgb, _bbcdg, _gabe int, _dgde RasterOperator, _gadb *Bitmap, _aedg, _gdba int) error {
	var (
		_gdbd          byte
		_afce          int
		_ccgd          int
		_dgee, _gaca   int
		_bdgaa, _bgfga int
	)
	_fdcd := _bbcdg >> 3
	_gga := _bbcdg & 7
	if _gga > 0 {
		_gdbd = _abfb[_gga]
	}
	_afce = _gadb.RowStride*_gdba + (_aedg >> 3)
	_ccgd = _bbbec.RowStride*_cebgb + (_cgbgd >> 3)
	switch _dgde {
	case PixSrc:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] = _gadb.Data[_dgee]
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], _gadb.Data[_dgee], _gdbd)
			}
		}
	case PixNotSrc:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] = ^(_gadb.Data[_dgee])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], ^_gadb.Data[_dgee], _gdbd)
			}
		}
	case PixSrcOrDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] |= _gadb.Data[_dgee]
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], _gadb.Data[_dgee]|_bbbec.Data[_gaca], _gdbd)
			}
		}
	case PixSrcAndDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] &= _gadb.Data[_dgee]
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], _gadb.Data[_dgee]&_bbbec.Data[_gaca], _gdbd)
			}
		}
	case PixSrcXorDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] ^= _gadb.Data[_dgee]
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], _gadb.Data[_dgee]^_bbbec.Data[_gaca], _gdbd)
			}
		}
	case PixNotSrcOrDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] |= ^(_gadb.Data[_dgee])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], ^(_gadb.Data[_dgee])|_bbbec.Data[_gaca], _gdbd)
			}
		}
	case PixNotSrcAndDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] &= ^(_gadb.Data[_dgee])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], ^(_gadb.Data[_dgee])&_bbbec.Data[_gaca], _gdbd)
			}
		}
	case PixSrcOrNotDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] = _gadb.Data[_dgee] | ^(_bbbec.Data[_gaca])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], _gadb.Data[_dgee]|^(_bbbec.Data[_gaca]), _gdbd)
			}
		}
	case PixSrcAndNotDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] = _gadb.Data[_dgee] &^ (_bbbec.Data[_gaca])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], _gadb.Data[_dgee]&^(_bbbec.Data[_gaca]), _gdbd)
			}
		}
	case PixNotPixSrcOrDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] = ^(_gadb.Data[_dgee] | _bbbec.Data[_gaca])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], ^(_gadb.Data[_dgee] | _bbbec.Data[_gaca]), _gdbd)
			}
		}
	case PixNotPixSrcAndDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] = ^(_gadb.Data[_dgee] & _bbbec.Data[_gaca])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], ^(_gadb.Data[_dgee] & _bbbec.Data[_gaca]), _gdbd)
			}
		}
	case PixNotPixSrcXorDst:
		for _bdgaa = 0; _bdgaa < _gabe; _bdgaa++ {
			_dgee = _afce + _bdgaa*_gadb.RowStride
			_gaca = _ccgd + _bdgaa*_bbbec.RowStride
			for _bgfga = 0; _bgfga < _fdcd; _bgfga++ {
				_bbbec.Data[_gaca] = ^(_gadb.Data[_dgee] ^ _bbbec.Data[_gaca])
				_gaca++
				_dgee++
			}
			if _gga > 0 {
				_bbbec.Data[_gaca] = _fcbb(_bbbec.Data[_gaca], ^(_gadb.Data[_dgee] ^ _bbbec.Data[_gaca]), _gdbd)
			}
		}
	default:
		_fce.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _dgde)
		return _c.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004co\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}

func _edb(_bfc *Bitmap, _ea, _aag int) (*Bitmap, error) {
	const _egd = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0052\u0065p\u006c\u0069\u0063\u0061\u0074\u0065"
	if _bfc == nil {
		return nil, _c.Error(_egd, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ea <= 0 || _aag <= 0 {
		return nil, _c.Error(_egd, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _ea == _aag {
		if _ea == 1 {
			_cad, _dda := _gbfag(nil, _bfc)
			if _dda != nil {
				return nil, _c.Wrap(_dda, _egd, "\u0078\u0046\u0061\u0063\u0074\u0020\u003d\u003d\u0020y\u0046\u0061\u0063\u0074")
			}
			return _cad, nil
		}
		if _ea == 2 || _ea == 4 || _ea == 8 {
			_fda, _bcf := _acb(_bfc, _ea)
			if _bcf != nil {
				return nil, _c.Wrap(_bcf, _egd, "\u0078\u0046a\u0063\u0074\u0020i\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
			}
			return _fda, nil
		}
	}
	_ggg := _ea * _bfc.Width
	_gd := _aag * _bfc.Height
	_ggf := New(_ggg, _gd)
	_edd := _ggf.RowStride
	var (
		_fde, _ecg, _dea, _fcd, _fae int
		_cd                          byte
		_ce                          error
	)
	for _ecg = 0; _ecg < _bfc.Height; _ecg++ {
		_fde = _aag * _ecg * _edd
		for _dea = 0; _dea < _bfc.Width; _dea++ {
			if _ff := _bfc.GetPixel(_dea, _ecg); _ff {
				_fae = _ea * _dea
				for _fcd = 0; _fcd < _ea; _fcd++ {
					_ggf.setBit(_fde*8 + _fae + _fcd)
				}
			}
		}
		for _fcd = 1; _fcd < _aag; _fcd++ {
			_dg := _fde + _fcd*_edd
			for _efb := 0; _efb < _edd; _efb++ {
				if _cd, _ce = _ggf.GetByte(_fde + _efb); _ce != nil {
					return nil, _c.Wrapf(_ce, _egd, "\u0072\u0065\u0070\u006cic\u0061\u0074\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0027\u0025d\u0027", _fcd)
				}
				if _ce = _ggf.SetByte(_dg+_efb, _cd); _ce != nil {
					return nil, _c.Wrap(_ce, _egd, "\u0053\u0065\u0074\u0074in\u0067\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
				}
			}
		}
	}
	return _ggf, nil
}
func (_dcf *Bitmap) CountPixels() int { return _dcf.countPixels() }
func _cabed(_aefc, _fbge *Bitmap, _eac, _bgdb, _cac, _eeegf, _aefad, _ecgg, _gfec, _afdd int, _gbag CombinationOperator, _efa int) error {
	var _egda int
	_aaeg := func() {
		_egda++
		_cac += _fbge.RowStride
		_eeegf += _aefc.RowStride
		_aefad += _aefc.RowStride
	}
	for _egda = _eac; _egda < _bgdb; _aaeg() {
		var _gggc uint16
		_daea := _cac
		for _adcc := _eeegf; _adcc <= _aefad; _adcc++ {
			_cgg, _fegfc := _fbge.GetByte(_daea)
			if _fegfc != nil {
				return _fegfc
			}
			_afgf, _fegfc := _aefc.GetByte(_adcc)
			if _fegfc != nil {
				return _fegfc
			}
			_gggc = (_gggc | (uint16(_afgf) & 0xff)) << uint(_afdd)
			_afgf = byte(_gggc >> 8)
			if _fegfc = _fbge.SetByte(_daea, _fbfd(_cgg, _afgf, _gbag)); _fegfc != nil {
				return _fegfc
			}
			_daea++
			_gggc <<= uint(_gfec)
			if _adcc == _aefad {
				_afgf = byte(_gggc >> (8 - uint8(_afdd)))
				if _efa != 0 {
					_afgf = _aabd(uint(8+_ecgg), _afgf)
				}
				_cgg, _fegfc = _fbge.GetByte(_daea)
				if _fegfc != nil {
					return _fegfc
				}
				if _fegfc = _fbge.SetByte(_daea, _fbfd(_cgg, _afgf, _gbag)); _fegfc != nil {
					return _fegfc
				}
			}
		}
	}
	return nil
}

type Color int

func (_gabfd *Bitmaps) makeSizeIndicator(_eeba, _befag int, _agfc LocationFilter, _egcbb SizeComparison) (_gbge *_fa.NumSlice, _gbcc error) {
	const _cdca = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u006d\u0061\u006b\u0065S\u0069z\u0065I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _gabfd == nil {
		return nil, _c.Error(_cdca, "\u0062\u0069\u0074ma\u0070\u0073\u0020\u0027\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch _agfc {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _c.Errorf(_cdca, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", _agfc)
	}
	switch _egcbb {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _c.Errorf(_cdca, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", _egcbb)
	}
	_gbge = &_fa.NumSlice{}
	var (
		_fdfca, _bebg, _geeac int
		_faeff                *Bitmap
	)
	for _, _faeff = range _gabfd.Values {
		_fdfca = 0
		_bebg, _geeac = _faeff.Width, _faeff.Height
		switch _agfc {
		case LocSelectWidth:
			if (_egcbb == SizeSelectIfLT && _bebg < _eeba) || (_egcbb == SizeSelectIfGT && _bebg > _eeba) || (_egcbb == SizeSelectIfLTE && _bebg <= _eeba) || (_egcbb == SizeSelectIfGTE && _bebg >= _eeba) || (_egcbb == SizeSelectIfEQ && _bebg == _eeba) {
				_fdfca = 1
			}
		case LocSelectHeight:
			if (_egcbb == SizeSelectIfLT && _geeac < _befag) || (_egcbb == SizeSelectIfGT && _geeac > _befag) || (_egcbb == SizeSelectIfLTE && _geeac <= _befag) || (_egcbb == SizeSelectIfGTE && _geeac >= _befag) || (_egcbb == SizeSelectIfEQ && _geeac == _befag) {
				_fdfca = 1
			}
		case LocSelectIfEither:
			if (_egcbb == SizeSelectIfLT && (_bebg < _eeba || _geeac < _befag)) || (_egcbb == SizeSelectIfGT && (_bebg > _eeba || _geeac > _befag)) || (_egcbb == SizeSelectIfLTE && (_bebg <= _eeba || _geeac <= _befag)) || (_egcbb == SizeSelectIfGTE && (_bebg >= _eeba || _geeac >= _befag)) || (_egcbb == SizeSelectIfEQ && (_bebg == _eeba || _geeac == _befag)) {
				_fdfca = 1
			}
		case LocSelectIfBoth:
			if (_egcbb == SizeSelectIfLT && (_bebg < _eeba && _geeac < _befag)) || (_egcbb == SizeSelectIfGT && (_bebg > _eeba && _geeac > _befag)) || (_egcbb == SizeSelectIfLTE && (_bebg <= _eeba && _geeac <= _befag)) || (_egcbb == SizeSelectIfGTE && (_bebg >= _eeba && _geeac >= _befag)) || (_egcbb == SizeSelectIfEQ && (_bebg == _eeba && _geeac == _befag)) {
				_fdfca = 1
			}
		}
		_gbge.AddInt(_fdfca)
	}
	return _gbge, nil
}

func _gadgd(_aabc *_fa.Stack) (_cacad *fillSegment, _cbdfc error) {
	const _cgcg = "\u0070\u006f\u0070\u0046\u0069\u006c\u006c\u0053\u0065g\u006d\u0065\u006e\u0074"
	if _aabc == nil {
		return nil, _c.Error(_cgcg, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _aabc.Aux == nil {
		return nil, _c.Error(_cgcg, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	_gfee, _ebde := _aabc.Pop()
	if !_ebde {
		return nil, nil
	}
	_eedf, _ebde := _gfee.(*fillSegment)
	if !_ebde {
		return nil, _c.Error(_cgcg, "\u0073\u0074\u0061ck\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020c\u006fn\u0074a\u0069n\u0020\u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_cacad = &fillSegment{_eedf._eddd, _eedf._decge, _eedf._dfef + _eedf._bbgg, _eedf._bbgg}
	_aabc.Aux.Push(_eedf)
	return _cacad, nil
}

func Blit(src *Bitmap, dst *Bitmap, x, y int, op CombinationOperator) error {
	var _aefa, _bfcce int
	_gee := src.RowStride - 1
	if x < 0 {
		_bfcce = -x
		x = 0
	} else if x+src.Width > dst.Width {
		_gee -= src.Width + x - dst.Width
	}
	if y < 0 {
		_aefa = -y
		y = 0
		_bfcce += src.RowStride
		_gee += src.RowStride
	} else if y+src.Height > dst.Height {
		_aefa = src.Height + y - dst.Height
	}
	var (
		_caeb int
		_fbf  error
	)
	_gege := x & 0x07
	_dcb := 8 - _gege
	_aeee := src.Width & 0x07
	_bfde := _dcb - _aeee
	_ddfa := _dcb&0x07 != 0
	_badf := src.Width <= ((_gee-_bfcce)<<3)+_dcb
	_cedgd := dst.GetByteIndex(x, y)
	_fdbf := _aefa + dst.Height
	if src.Height > _fdbf {
		_caeb = _fdbf
	} else {
		_caeb = src.Height
	}
	switch {
	case !_ddfa:
		_fbf = _bccg(src, dst, _aefa, _caeb, _cedgd, _bfcce, _gee, op)
	case _badf:
		_fbf = _dbgc(src, dst, _aefa, _caeb, _cedgd, _bfcce, _gee, _bfde, _gege, _dcb, op)
	default:
		_fbf = _cabed(src, dst, _aefa, _caeb, _cedgd, _bfcce, _gee, _bfde, _gege, _dcb, op, _aeee)
	}
	return _fbf
}

func (_bgeb *Bitmaps) HeightSorter() func(_abee, _adfc int) bool {
	return func(_gbefb, _adca int) bool {
		_gegd := _bgeb.Values[_gbefb].Height < _bgeb.Values[_adca].Height
		_fce.Log.Debug("H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0076\u0020<\u0020\u0025\u0076\u0020= \u0025\u0076", _bgeb.Values[_gbefb].Height, _bgeb.Values[_adca].Height, _gegd)
		return _gegd
	}
}

func (_fbcb Points) GetIntX(i int) (int, error) {
	if i >= len(_fbcb) {
		return 0, _c.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0058", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_fbcb[i].X), nil
}

func (_ecfdd *ClassedPoints) Swap(i, j int) {
	_ecfdd.IntSlice[i], _ecfdd.IntSlice[j] = _ecfdd.IntSlice[j], _ecfdd.IntSlice[i]
}
func (_deaf *Bitmap) Equivalent(s *Bitmap) bool { return _deaf.equivalent(s) }

type MorphOperation int

func NewWithUnpaddedData(width, height int, data []byte) (*Bitmap, error) {
	const _bdcb = "\u004e\u0065\u0077\u0057it\u0068\u0055\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	_afg := _aeg(width, height)
	_afg.Data = data
	if _bbf := ((width * height) + 7) >> 3; len(data) < _bbf {
		return nil, _c.Errorf(_bdcb, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(data), _bbf)
	}
	if _gdcc := _afg.addPadBits(); _gdcc != nil {
		return nil, _c.Wrap(_gdcc, _bdcb, "")
	}
	return _afg, nil
}

type Component int

func (_afcbc *Bitmaps) String() string {
	_ddab := _ef.Builder{}
	for _, _abccb := range _afcbc.Values {
		_ddab.WriteString(_abccb.String())
		_ddab.WriteRune('\n')
	}
	return _ddab.String()
}

const (
	SelDontCare SelectionValue = iota
	SelHit
	SelMiss
)
const _fceea = 5000

func CorrelationScoreThresholded(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab, downcount []int, scoreThreshold float32) (bool, error) {
	const _fbfb = "C\u006f\u0072\u0072\u0065\u006c\u0061t\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054h\u0072\u0065\u0073h\u006fl\u0064\u0065\u0064"
	if bm1 == nil {
		return false, _c.Error(_fbfb, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d1\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if bm2 == nil {
		return false, _c.Error(_fbfb, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d2\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if area1 <= 0 || area2 <= 0 {
		return false, _c.Error(_fbfb, "c\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0053\u0063\u006f\u0072\u0065\u0054\u0068re\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u002d\u0020\u0061\u0072\u0065\u0061s \u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u003e\u0020\u0030")
	}
	if downcount == nil {
		return false, _c.Error(_fbfb, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u006f\u0020\u0027\u0064\u006f\u0077\u006e\u0063\u006f\u0075\u006e\u0074\u0027")
	}
	if tab == nil {
		return false, _c.Error(_fbfb, "p\u0072\u006f\u0076\u0069de\u0064 \u006e\u0069\u006c\u0020\u0027s\u0075\u006d\u0074\u0061\u0062\u0027")
	}
	_bgdf, _cdbe := bm1.Width, bm1.Height
	_adag, _geda := bm2.Width, bm2.Height
	if _fa.Abs(_bgdf-_adag) > maxDiffW {
		return false, nil
	}
	if _fa.Abs(_cdbe-_geda) > maxDiffH {
		return false, nil
	}
	_accb := int(delX + _fa.Sign(delX)*0.5)
	_cbbf := int(delY + _fa.Sign(delY)*0.5)
	_bfcg := int(_fc.Ceil(_fc.Sqrt(float64(scoreThreshold) * float64(area1) * float64(area2))))
	_gbbe := bm2.RowStride
	_edbe := _fcdga(_cbbf, 0)
	_geed := _fcac(_geda+_cbbf, _cdbe)
	_fedc := bm1.RowStride * _edbe
	_degd := bm2.RowStride * (_edbe - _cbbf)
	var _cbef int
	if _geed <= _cdbe {
		_cbef = downcount[_geed-1]
	}
	_abc := _fcdga(_accb, 0)
	_ecbg := _fcac(_adag+_accb, _bgdf)
	var _geea, _fgage int
	if _accb >= 8 {
		_geea = _accb >> 3
		_fedc += _geea
		_abc -= _geea << 3
		_ecbg -= _geea << 3
		_accb &= 7
	} else if _accb <= -8 {
		_fgage = -((_accb + 7) >> 3)
		_degd += _fgage
		_gbbe -= _fgage
		_accb += _fgage << 3
	}
	var (
		_eafe, _aacg, _eeff int
		_edegf, _fbe, _fafa byte
	)
	if _abc >= _ecbg || _edbe >= _geed {
		return false, nil
	}
	_begg := (_ecbg + 7) >> 3
	switch {
	case _accb == 0:
		for _aacg = _edbe; _aacg < _geed; _aacg, _fedc, _degd = _aacg+1, _fedc+bm1.RowStride, _degd+bm2.RowStride {
			for _eeff = 0; _eeff < _begg; _eeff++ {
				_edegf = bm1.Data[_fedc+_eeff] & bm2.Data[_degd+_eeff]
				_eafe += tab[_edegf]
			}
			if _eafe >= _bfcg {
				return true, nil
			}
			if _gaee := _eafe + downcount[_aacg] - _cbef; _gaee < _bfcg {
				return false, nil
			}
		}
	case _accb > 0 && _gbbe < _begg:
		for _aacg = _edbe; _aacg < _geed; _aacg, _fedc, _degd = _aacg+1, _fedc+bm1.RowStride, _degd+bm2.RowStride {
			_fbe = bm1.Data[_fedc]
			_fafa = bm2.Data[_degd] >> uint(_accb)
			_edegf = _fbe & _fafa
			_eafe += tab[_edegf]
			for _eeff = 1; _eeff < _gbbe; _eeff++ {
				_fbe = bm1.Data[_fedc+_eeff]
				_fafa = bm2.Data[_degd+_eeff]>>uint(_accb) | bm2.Data[_degd+_eeff-1]<<uint(8-_accb)
				_edegf = _fbe & _fafa
				_eafe += tab[_edegf]
			}
			_fbe = bm1.Data[_fedc+_eeff]
			_fafa = bm2.Data[_degd+_eeff-1] << uint(8-_accb)
			_edegf = _fbe & _fafa
			_eafe += tab[_edegf]
			if _eafe >= _bfcg {
				return true, nil
			} else if _eafe+downcount[_aacg]-_cbef < _bfcg {
				return false, nil
			}
		}
	case _accb > 0 && _gbbe >= _begg:
		for _aacg = _edbe; _aacg < _geed; _aacg, _fedc, _degd = _aacg+1, _fedc+bm1.RowStride, _degd+bm2.RowStride {
			_fbe = bm1.Data[_fedc]
			_fafa = bm2.Data[_degd] >> uint(_accb)
			_edegf = _fbe & _fafa
			_eafe += tab[_edegf]
			for _eeff = 1; _eeff < _begg; _eeff++ {
				_fbe = bm1.Data[_fedc+_eeff]
				_fafa = bm2.Data[_degd+_eeff] >> uint(_accb)
				_fafa |= bm2.Data[_degd+_eeff-1] << uint(8-_accb)
				_edegf = _fbe & _fafa
				_eafe += tab[_edegf]
			}
			if _eafe >= _bfcg {
				return true, nil
			} else if _eafe+downcount[_aacg]-_cbef < _bfcg {
				return false, nil
			}
		}
	case _begg < _gbbe:
		for _aacg = _edbe; _aacg < _geed; _aacg, _fedc, _degd = _aacg+1, _fedc+bm1.RowStride, _degd+bm2.RowStride {
			for _eeff = 0; _eeff < _begg; _eeff++ {
				_fbe = bm1.Data[_fedc+_eeff]
				_fafa = bm2.Data[_degd+_eeff] << uint(-_accb)
				_fafa |= bm2.Data[_degd+_eeff+1] >> uint(8+_accb)
				_edegf = _fbe & _fafa
				_eafe += tab[_edegf]
			}
			if _eafe >= _bfcg {
				return true, nil
			} else if _bddgf := _eafe + downcount[_aacg] - _cbef; _bddgf < _bfcg {
				return false, nil
			}
		}
	case _gbbe >= _begg:
		for _aacg = _edbe; _aacg < _geed; _aacg, _fedc, _degd = _aacg+1, _fedc+bm1.RowStride, _degd+bm2.RowStride {
			for _eeff = 0; _eeff < _begg; _eeff++ {
				_fbe = bm1.Data[_fedc+_eeff]
				_fafa = bm2.Data[_degd+_eeff] << uint(-_accb)
				_fafa |= bm2.Data[_degd+_eeff+1] >> uint(8+_accb)
				_edegf = _fbe & _fafa
				_eafe += tab[_edegf]
			}
			_fbe = bm1.Data[_fedc+_eeff]
			_fafa = bm2.Data[_degd+_eeff] << uint(-_accb)
			_edegf = _fbe & _fafa
			_eafe += tab[_edegf]
			if _eafe >= _bfcg {
				return true, nil
			} else if _eafe+downcount[_aacg]-_cbef < _bfcg {
				return false, nil
			}
		}
	}
	_cgbga := float32(_eafe) * float32(_eafe) / (float32(area1) * float32(area2))
	if _cgbga >= scoreThreshold {
		_fce.Log.Trace("\u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064\u0020\u003c\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0020\u0025\u0064\u0020\u0062\u0075\u0074\u0020\u0073c\u006f\u0072\u0065\u0020\u0025\u0066\u0020\u003e\u003d\u0020\u0073\u0063\u006fr\u0065\u0054\u0068\u0072\u0065\u0073h\u006f\u006c\u0064 \u0025\u0066", _eafe, _bfcg, _cgbga, scoreThreshold)
	}
	return false, nil
}

type Points []Point

func (_caegd *byWidth) Swap(i, j int) {
	_caegd.Values[i], _caegd.Values[j] = _caegd.Values[j], _caegd.Values[i]
	if _caegd.Boxes != nil {
		_caegd.Boxes[i], _caegd.Boxes[j] = _caegd.Boxes[j], _caegd.Boxes[i]
	}
}

type LocationFilter int

func (_agce *Bitmaps) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_gefc *Bitmaps, _aabdg error) {
	const _dade = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0053\u0065\u006ce\u0063\u0074\u0042\u0079Si\u007a\u0065"
	if _agce == nil {
		return nil, _c.Error(_dade, "\u0027\u0062\u0027 B\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _c.Errorf(_dade, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _c.Errorf(_dade, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", relation)
	}
	_bfaf, _aabdg := _agce.makeSizeIndicator(width, height, tp, relation)
	if _aabdg != nil {
		return nil, _c.Wrap(_aabdg, _dade, "")
	}
	_gefc, _aabdg = _agce.selectByIndicator(_bfaf)
	if _aabdg != nil {
		return nil, _c.Wrap(_aabdg, _dade, "")
	}
	return _gefc, nil
}

var _ _cc.Interface = &ClassedPoints{}

func _aabd(_afec uint, _agb byte) byte { return _agb >> _afec << _afec }
func (_bfdb *Bitmap) ConnComponents(bms *Bitmaps, connectivity int) (_gaaaf *Boxes, _deec error) {
	const _deaab = "B\u0069\u0074\u006d\u0061p.\u0043o\u006e\u006e\u0043\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0073"
	if _bfdb == nil {
		return nil, _c.Error(_deaab, "\u0070r\u006f\u0076\u0069\u0064e\u0064\u0020\u0065\u006d\u0070t\u0079 \u0027b\u0027\u0020\u0062\u0069\u0074\u006d\u0061p")
	}
	if connectivity != 4 && connectivity != 8 {
		return nil, _c.Error(_deaab, "\u0063\u006f\u006ene\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006e\u006f\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038")
	}
	if bms == nil {
		if _gaaaf, _deec = _bfdb.connComponentsBB(connectivity); _deec != nil {
			return nil, _c.Wrap(_deec, _deaab, "")
		}
	} else {
		if _gaaaf, _deec = _bfdb.connComponentsBitmapsBB(bms, connectivity); _deec != nil {
			return nil, _c.Wrap(_deec, _deaab, "")
		}
	}
	return _gaaaf, nil
}

var (
	_abfb  = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_fgcfb = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func _gcgg(_adeg, _dbc *Bitmap, _ceec *Selection) (*Bitmap, error) {
	const _fgba = "c\u006c\u006f\u0073\u0065\u0042\u0069\u0074\u006d\u0061\u0070"
	var _eacac error
	if _adeg, _eacac = _bbegg(_adeg, _dbc, _ceec); _eacac != nil {
		return nil, _eacac
	}
	_eabgd, _eacac := _bag(nil, _dbc, _ceec)
	if _eacac != nil {
		return nil, _c.Wrap(_eacac, _fgba, "")
	}
	if _, _eacac = _gcae(_adeg, _eabgd, _ceec); _eacac != nil {
		return nil, _c.Wrap(_eacac, _fgba, "")
	}
	return _adeg, nil
}

func (_ddfba Points) GetIntY(i int) (int, error) {
	if i >= len(_ddfba) {
		return 0, _c.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0059", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_ddfba[i].Y), nil
}

func NewWithData(width, height int, data []byte) (*Bitmap, error) {
	const _cdd = "N\u0065\u0077\u0057\u0069\u0074\u0068\u0044\u0061\u0074\u0061"
	_aab := _aeg(width, height)
	_aab.Data = data
	if len(data) < height*_aab.RowStride {
		return nil, _c.Errorf(_cdd, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u003a \u0025\u0064\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u003a\u0020\u0025\u0064", len(data), height*_aab.RowStride)
	}
	return _aab, nil
}

func _bdegg(_fece *Bitmap, _bgab *_fa.Stack, _gagd, _afdb int) (_afecc *_g.Rectangle, _bagc error) {
	const _gdff = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _fece == nil {
		return nil, _c.Error(_gdff, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _bgab == nil {
		return nil, _c.Error(_gdff, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_cgad, _gagc := _fece.Width, _fece.Height
	_bgaa := _cgad - 1
	_bece := _gagc - 1
	if _gagd < 0 || _gagd > _bgaa || _afdb < 0 || _afdb > _bece || !_fece.GetPixel(_gagd, _afdb) {
		return nil, nil
	}
	var _egfb *_g.Rectangle
	_egfb, _bagc = Rect(100000, 100000, 0, 0)
	if _bagc != nil {
		return nil, _c.Wrap(_bagc, _gdff, "")
	}
	if _bagc = _dcge(_bgab, _gagd, _gagd, _afdb, 1, _bece, _egfb); _bagc != nil {
		return nil, _c.Wrap(_bagc, _gdff, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _bagc = _dcge(_bgab, _gagd, _gagd, _afdb+1, -1, _bece, _egfb); _bagc != nil {
		return nil, _c.Wrap(_bagc, _gdff, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_egfb.Min.X, _egfb.Max.X = _gagd, _gagd
	_egfb.Min.Y, _egfb.Max.Y = _afdb, _afdb
	var (
		_badd *fillSegment
		_afbf int
	)
	for _bgab.Len() > 0 {
		if _badd, _bagc = _gadgd(_bgab); _bagc != nil {
			return nil, _c.Wrap(_bagc, _gdff, "")
		}
		_afdb = _badd._dfef
		for _gagd = _badd._eddd; _gagd >= 0 && _fece.GetPixel(_gagd, _afdb); _gagd-- {
			if _bagc = _fece.SetPixel(_gagd, _afdb, 0); _bagc != nil {
				return nil, _c.Wrap(_bagc, _gdff, "")
			}
		}
		if _gagd >= _badd._eddd {
			for _gagd++; _gagd <= _badd._decge && _gagd <= _bgaa && !_fece.GetPixel(_gagd, _afdb); _gagd++ {
			}
			_afbf = _gagd
			if !(_gagd <= _badd._decge && _gagd <= _bgaa) {
				continue
			}
		} else {
			_afbf = _gagd + 1
			if _afbf < _badd._eddd-1 {
				if _bagc = _dcge(_bgab, _afbf, _badd._eddd-1, _badd._dfef, -_badd._bbgg, _bece, _egfb); _bagc != nil {
					return nil, _c.Wrap(_bagc, _gdff, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
				}
			}
			_gagd = _badd._eddd + 1
		}
		for {
			for ; _gagd <= _bgaa && _fece.GetPixel(_gagd, _afdb); _gagd++ {
				if _bagc = _fece.SetPixel(_gagd, _afdb, 0); _bagc != nil {
					return nil, _c.Wrap(_bagc, _gdff, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _bagc = _dcge(_bgab, _afbf, _gagd-1, _badd._dfef, _badd._bbgg, _bece, _egfb); _bagc != nil {
				return nil, _c.Wrap(_bagc, _gdff, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _gagd > _badd._decge+1 {
				if _bagc = _dcge(_bgab, _badd._decge+1, _gagd-1, _badd._dfef, -_badd._bbgg, _bece, _egfb); _bagc != nil {
					return nil, _c.Wrap(_bagc, _gdff, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _gagd++; _gagd <= _badd._decge && _gagd <= _bgaa && !_fece.GetPixel(_gagd, _afdb); _gagd++ {
			}
			_afbf = _gagd
			if !(_gagd <= _badd._decge && _gagd <= _bgaa) {
				break
			}
		}
	}
	_egfb.Max.X++
	_egfb.Max.Y++
	return _egfb, nil
}

func _afb(_gce, _da *Bitmap, _dbf int, _acd []byte, _gbd int) (_fac error) {
	const _dbag = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0031"
	var (
		_bca, _ad, _bcff, _bbd, _eegg, _fdac, _fgbd, _fdab int
		_dfc, _fe                                          uint32
		_afa, _gbb                                         byte
		_gbc                                               uint16
	)
	_cea := make([]byte, 4)
	_cfa := make([]byte, 4)
	for _bcff = 0; _bcff < _gce.Height-1; _bcff, _bbd = _bcff+2, _bbd+1 {
		_bca = _bcff * _gce.RowStride
		_ad = _bbd * _da.RowStride
		for _eegg, _fdac = 0, 0; _eegg < _gbd; _eegg, _fdac = _eegg+4, _fdac+1 {
			for _fgbd = 0; _fgbd < 4; _fgbd++ {
				_fdab = _bca + _eegg + _fgbd
				if _fdab <= len(_gce.Data)-1 && _fdab < _bca+_gce.RowStride {
					_cea[_fgbd] = _gce.Data[_fdab]
				} else {
					_cea[_fgbd] = 0x00
				}
				_fdab = _bca + _gce.RowStride + _eegg + _fgbd
				if _fdab <= len(_gce.Data)-1 && _fdab < _bca+(2*_gce.RowStride) {
					_cfa[_fgbd] = _gce.Data[_fdab]
				} else {
					_cfa[_fgbd] = 0x00
				}
			}
			_dfc = _a.BigEndian.Uint32(_cea)
			_fe = _a.BigEndian.Uint32(_cfa)
			_fe |= _dfc
			_fe |= _fe << 1
			_fe &= 0xaaaaaaaa
			_dfc = _fe | (_fe << 7)
			_afa = byte(_dfc >> 24)
			_gbb = byte((_dfc >> 8) & 0xff)
			_fdab = _ad + _fdac
			if _fdab+1 == len(_da.Data)-1 || _fdab+1 >= _ad+_da.RowStride {
				_da.Data[_fdab] = _acd[_afa]
			} else {
				_gbc = (uint16(_acd[_afa]) << 8) | uint16(_acd[_gbb])
				if _fac = _da.setTwoBytes(_fdab, _gbc); _fac != nil {
					return _c.Wrapf(_fac, _dbag, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fdab)
				}
				_fdac++
			}
		}
	}
	return nil
}

type fillSegment struct {
	_eddd  int
	_decge int
	_dfef  int
	_bbgg  int
}

var (
	_dbdg *Bitmap
	_daff *Bitmap
)

func _fga(_eeec *Bitmap, _acg ...int) (_eeg *Bitmap, _gfa error) {
	const _fcdg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0043\u0061\u0073\u0063\u0061\u0064\u0065"
	if _eeec == nil {
		return nil, _c.Error(_fcdg, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_acg) == 0 || len(_acg) > 4 {
		return nil, _c.Error(_fcdg, "t\u0068\u0065\u0072\u0065\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0074\u0020\u006cea\u0073\u0074\u0020\u006fn\u0065\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020mo\u0073\u0074 \u0034\u0020\u006c\u0065\u0076\u0065\u006c\u0073")
	}
	if _acg[0] <= 0 {
		_fce.Log.Debug("\u006c\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030 \u002d\u0020\u006e\u006f\u0020\u0072\u0065\u0064\u0075\u0063t\u0069\u006f\u006e")
		_eeg, _gfa = _gbfag(nil, _eeec)
		if _gfa != nil {
			return nil, _c.Wrap(_gfa, _fcdg, "l\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030")
		}
		return _eeg, nil
	}
	_aaf := _fcgc()
	_eeg = _eeec
	for _gdc, _dfa := range _acg {
		if _dfa <= 0 {
			break
		}
		_eeg, _gfa = _afd(_eeg, _dfa, _aaf)
		if _gfa != nil {
			return nil, _c.Wrapf(_gfa, _fcdg, "\u006c\u0065\u0076\u0065\u006c\u0025\u0064\u0020\u0072\u0065\u0064\u0075c\u0074\u0069\u006f\u006e", _gdc)
		}
	}
	return _eeg, nil
}

var MorphBC BoundaryCondition

func _agfd(_fgcf *Bitmap, _edge, _deba, _dgfb, _eeda int, _aagd RasterOperator, _cdgc *Bitmap, _bafce, _dgac int) error {
	var (
		_efcc        bool
		_dcfc        bool
		_bcecd       int
		_gfecf       int
		_ddae        int
		_egab        bool
		_aebdc       byte
		_bccd        int
		_fgfc        int
		_fcgbg       int
		_efde, _dgfd int
	)
	_fccb := 8 - (_edge & 7)
	_dcdc := _fgcfb[_fccb]
	_bcgc := _fgcf.RowStride*_deba + (_edge >> 3)
	_dgbb := _cdgc.RowStride*_dgac + (_bafce >> 3)
	if _dgfb < _fccb {
		_efcc = true
		_dcdc &= _abfb[8-_fccb+_dgfb]
	}
	if !_efcc {
		_bcecd = (_dgfb - _fccb) >> 3
		if _bcecd > 0 {
			_dcfc = true
			_gfecf = _bcgc + 1
			_ddae = _dgbb + 1
		}
	}
	_bccd = (_edge + _dgfb) & 7
	if !(_efcc || _bccd == 0) {
		_egab = true
		_aebdc = _abfb[_bccd]
		_fgfc = _bcgc + 1 + _bcecd
		_fcgbg = _dgbb + 1 + _bcecd
	}
	switch _aagd {
	case PixSrc:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], _cdgc.Data[_dgbb], _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] = _cdgc.Data[_ddae+_dgfd]
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], _cdgc.Data[_fcgbg], _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixNotSrc:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], ^_cdgc.Data[_dgbb], _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] = ^_cdgc.Data[_ddae+_dgfd]
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], ^_cdgc.Data[_fcgbg], _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixSrcOrDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], _cdgc.Data[_dgbb]|_fgcf.Data[_bcgc], _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] |= _cdgc.Data[_ddae+_dgfd]
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], _cdgc.Data[_fcgbg]|_fgcf.Data[_fgfc], _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixSrcAndDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], _cdgc.Data[_dgbb]&_fgcf.Data[_bcgc], _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] &= _cdgc.Data[_ddae+_dgfd]
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], _cdgc.Data[_fcgbg]&_fgcf.Data[_fgfc], _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixSrcXorDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], _cdgc.Data[_dgbb]^_fgcf.Data[_bcgc], _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] ^= _cdgc.Data[_ddae+_dgfd]
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], _cdgc.Data[_fcgbg]^_fgcf.Data[_fgfc], _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixNotSrcOrDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], ^(_cdgc.Data[_dgbb])|_fgcf.Data[_bcgc], _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] |= ^(_cdgc.Data[_ddae+_dgfd])
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], ^(_cdgc.Data[_fcgbg])|_fgcf.Data[_fgfc], _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixNotSrcAndDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], ^(_cdgc.Data[_dgbb])&_fgcf.Data[_bcgc], _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] &= ^_cdgc.Data[_ddae+_dgfd]
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], ^(_cdgc.Data[_fcgbg])&_fgcf.Data[_fgfc], _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixSrcOrNotDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], _cdgc.Data[_dgbb]|^(_fgcf.Data[_bcgc]), _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] = _cdgc.Data[_ddae+_dgfd] | ^(_fgcf.Data[_gfecf+_dgfd])
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], _cdgc.Data[_fcgbg]|^(_fgcf.Data[_fgfc]), _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixSrcAndNotDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], _cdgc.Data[_dgbb]&^(_fgcf.Data[_bcgc]), _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] = _cdgc.Data[_ddae+_dgfd] &^ (_fgcf.Data[_gfecf+_dgfd])
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], _cdgc.Data[_fcgbg]&^(_fgcf.Data[_fgfc]), _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], ^(_cdgc.Data[_dgbb] | _fgcf.Data[_bcgc]), _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] = ^(_cdgc.Data[_ddae+_dgfd] | _fgcf.Data[_gfecf+_dgfd])
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], ^(_cdgc.Data[_fcgbg] | _fgcf.Data[_fgfc]), _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], ^(_cdgc.Data[_dgbb] & _fgcf.Data[_bcgc]), _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] = ^(_cdgc.Data[_ddae+_dgfd] & _fgcf.Data[_gfecf+_dgfd])
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], ^(_cdgc.Data[_fcgbg] & _fgcf.Data[_fgfc]), _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		for _efde = 0; _efde < _eeda; _efde++ {
			_fgcf.Data[_bcgc] = _fcbb(_fgcf.Data[_bcgc], ^(_cdgc.Data[_dgbb] ^ _fgcf.Data[_bcgc]), _dcdc)
			_bcgc += _fgcf.RowStride
			_dgbb += _cdgc.RowStride
		}
		if _dcfc {
			for _efde = 0; _efde < _eeda; _efde++ {
				for _dgfd = 0; _dgfd < _bcecd; _dgfd++ {
					_fgcf.Data[_gfecf+_dgfd] = ^(_cdgc.Data[_ddae+_dgfd] ^ _fgcf.Data[_gfecf+_dgfd])
				}
				_gfecf += _fgcf.RowStride
				_ddae += _cdgc.RowStride
			}
		}
		if _egab {
			for _efde = 0; _efde < _eeda; _efde++ {
				_fgcf.Data[_fgfc] = _fcbb(_fgcf.Data[_fgfc], ^(_cdgc.Data[_fcgbg] ^ _fgcf.Data[_fgfc]), _aebdc)
				_fgfc += _fgcf.RowStride
				_fcgbg += _cdgc.RowStride
			}
		}
	default:
		_fce.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _aagd)
		return _c.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0056\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004c\u006f\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}

func (_ecad *Bitmap) GetComponents(components Component, maxWidth, maxHeight int) (_fcfc *Bitmaps, _efcf *Boxes, _dgge error) {
	const _bfdc = "B\u0069t\u006d\u0061\u0070\u002e\u0047\u0065\u0074\u0043o\u006d\u0070\u006f\u006een\u0074\u0073"
	if _ecad == nil {
		return nil, nil, _c.Error(_bfdc, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0042\u0069\u0074\u006da\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064\u002e")
	}
	switch components {
	case ComponentConn, ComponentCharacters, ComponentWords:
	default:
		return nil, nil, _c.Error(_bfdc, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074s\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072")
	}
	if _ecad.Zero() {
		_efcf = &Boxes{}
		_fcfc = &Bitmaps{}
		return _fcfc, _efcf, nil
	}
	switch components {
	case ComponentConn:
		_fcfc = &Bitmaps{}
		if _efcf, _dgge = _ecad.ConnComponents(_fcfc, 8); _dgge != nil {
			return nil, nil, _c.Wrap(_dgge, _bfdc, "\u006e\u006f \u0070\u0072\u0065p\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentCharacters:
		_acee, _gcdbg := MorphSequence(_ecad, MorphProcess{Operation: MopClosing, Arguments: []int{1, 6}})
		if _gcdbg != nil {
			return nil, nil, _c.Wrap(_gcdbg, _bfdc, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _fce.Log.IsLogLevel(_fce.LogLevelTrace) {
			_fce.Log.Trace("\u0043o\u006d\u0070o\u006e\u0065\u006e\u0074C\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074ma\u0070\u0020\u0061f\u0074\u0065r\u0020\u0063\u006c\u006f\u0073\u0069n\u0067\u003a \u0025\u0073", _acee.String())
		}
		_bcdb := &Bitmaps{}
		_efcf, _gcdbg = _acee.ConnComponents(_bcdb, 8)
		if _gcdbg != nil {
			return nil, nil, _c.Wrap(_gcdbg, _bfdc, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _fce.Log.IsLogLevel(_fce.LogLevelTrace) {
			_fce.Log.Trace("\u0043\u006f\u006d\u0070\u006f\u006ee\u006e\u0074\u0043\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0020a\u0066\u0074\u0065\u0072\u0020\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076i\u0074y\u003a\u0020\u0025\u0073", _bcdb.String())
		}
		if _fcfc, _gcdbg = _bcdb.ClipToBitmap(_ecad); _gcdbg != nil {
			return nil, nil, _c.Wrap(_gcdbg, _bfdc, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentWords:
		_eeeb := 1
		var _dcbb *Bitmap
		switch {
		case _ecad.XResolution <= 200:
			_dcbb = _ecad
		case _ecad.XResolution <= 400:
			_eeeb = 2
			_dcbb, _dgge = _fga(_ecad, 1, 0, 0, 0)
			if _dgge != nil {
				return nil, nil, _c.Wrap(_dgge, _bfdc, "w\u006f\u0072\u0064\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0020\u002d \u0078\u0072\u0065s\u003c=\u0034\u0030\u0030")
			}
		default:
			_eeeb = 4
			_dcbb, _dgge = _fga(_ecad, 1, 1, 0, 0)
			if _dgge != nil {
				return nil, nil, _c.Wrap(_dgge, _bfdc, "\u0077\u006f\u0072\u0064 \u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073 \u002d \u0078\u0072\u0065\u0073\u0020\u003e\u00204\u0030\u0030")
			}
		}
		_ddbfcf, _, _eceb := _caec(_dcbb)
		if _eceb != nil {
			return nil, nil, _c.Wrap(_eceb, _bfdc, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_fcef, _eceb := _ebbf(_ddbfcf, _eeeb)
		if _eceb != nil {
			return nil, nil, _c.Wrap(_eceb, _bfdc, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_gfd := &Bitmaps{}
		if _efcf, _eceb = _fcef.ConnComponents(_gfd, 4); _eceb != nil {
			return nil, nil, _c.Wrap(_eceb, _bfdc, "\u0077\u006f\u0072\u0064\u0020\u0070r\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u002c\u0020\u0063\u006f\u006en\u0065\u0063\u0074\u0020\u0065\u0078\u0070a\u006e\u0064\u0065\u0064")
		}
		if _fcfc, _eceb = _gfd.ClipToBitmap(_ecad); _eceb != nil {
			return nil, nil, _c.Wrap(_eceb, _bfdc, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
	}
	_fcfc, _dgge = _fcfc.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _dgge != nil {
		return nil, nil, _c.Wrap(_dgge, _bfdc, "")
	}
	_efcf, _dgge = _efcf.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _dgge != nil {
		return nil, nil, _c.Wrap(_dgge, _bfdc, "")
	}
	return _fcfc, _efcf, nil
}

func NewClassedPoints(points *Points, classes _fa.IntSlice) (*ClassedPoints, error) {
	const _gedf = "\u004e\u0065w\u0043\u006c\u0061s\u0073\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073"
	if points == nil {
		return nil, _c.Error(_gedf, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if classes == nil {
		return nil, _c.Error(_gedf, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0063\u006c\u0061ss\u0065\u0073")
	}
	_gdfcf := &ClassedPoints{Points: points, IntSlice: classes}
	if _fbagc := _gdfcf.validateIntSlice(); _fbagc != nil {
		return nil, _c.Wrap(_fbagc, _gedf, "")
	}
	return _gdfcf, nil
}

func (_aeeac *Bitmaps) GroupByWidth() (*BitmapsArray, error) {
	const _ffcf = "\u0047\u0072\u006fu\u0070\u0042\u0079\u0057\u0069\u0064\u0074\u0068"
	if len(_aeeac.Values) == 0 {
		return nil, _c.Error(_ffcf, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_dbfb := &BitmapsArray{}
	_aeeac.SortByWidth()
	_bdacd := -1
	_dabd := -1
	for _ggac := 0; _ggac < len(_aeeac.Values); _ggac++ {
		_eeebe := _aeeac.Values[_ggac].Width
		if _eeebe > _bdacd {
			_bdacd = _eeebe
			_dabd++
			_dbfb.Values = append(_dbfb.Values, &Bitmaps{})
		}
		_dbfb.Values[_dabd].AddBitmap(_aeeac.Values[_ggac])
	}
	return _dbfb, nil
}

func _fbfd(_ebdd, _fadae byte, _ddbfc CombinationOperator) byte {
	switch _ddbfc {
	case CmbOpOr:
		return _fadae | _ebdd
	case CmbOpAnd:
		return _fadae & _ebdd
	case CmbOpXor:
		return _fadae ^ _ebdd
	case CmbOpXNor:
		return ^(_fadae ^ _ebdd)
	case CmbOpNot:
		return ^(_fadae)
	default:
		return _fadae
	}
}

const (
	_aaddc shift = iota
	_dded
)

func (_caff *Bitmap) setEightBytes(_fed int, _dacg uint64) error {
	_bbff := _caff.RowStride - (_fed % _caff.RowStride)
	if _caff.RowStride != _caff.Width>>3 {
		_bbff--
	}
	if _bbff >= 8 {
		return _caff.setEightFullBytes(_fed, _dacg)
	}
	return _caff.setEightPartlyBytes(_fed, _bbff, _dacg)
}

func (_dgfc *Bitmap) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _gada(_dgfc, dx, dy, dw, dh, op, src, sx, sy)
}
func (_ceeg *byHeight) Len() int { return len(_ceeg.Values) }

var _cee [256]uint8

func _gb(_cfg, _af *Bitmap) (_eb error) {
	const _ag = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0034"
	_aff := _af.RowStride
	_gfb := _cfg.RowStride
	_cg := _af.RowStride*4 - _cfg.RowStride
	var (
		_ca, _efd                            byte
		_ge                                  uint32
		_cae, _eg, _bb, _bf, _cbc, _gfg, _db int
	)
	for _bb = 0; _bb < _af.Height; _bb++ {
		_cae = _bb * _aff
		_eg = 4 * _bb * _gfb
		for _bf = 0; _bf < _aff; _bf++ {
			_ca = _af.Data[_cae+_bf]
			_ge = _fdag[_ca]
			_gfg = _eg + _bf*4
			if _cg != 0 && (_bf+1)*4 > _cfg.RowStride {
				for _cbc = _cg; _cbc > 0; _cbc-- {
					_efd = byte((_ge >> uint(_cbc*8)) & 0xff)
					_db = _gfg + (_cg - _cbc)
					if _eb = _cfg.SetByte(_db, _efd); _eb != nil {
						return _c.Wrapf(_eb, _ag, "D\u0069\u0066\u0066\u0065\u0072\u0065n\u0074\u0020\u0072\u006f\u0077\u0073\u0074\u0072\u0069d\u0065\u0073\u002e \u004b:\u0020\u0025\u0064", _cbc)
					}
				}
			} else if _eb = _cfg.setFourBytes(_gfg, _ge); _eb != nil {
				return _c.Wrap(_eb, _ag, "")
			}
			if _eb = _cfg.setFourBytes(_eg+_bf*4, _fdag[_af.Data[_cae+_bf]]); _eb != nil {
				return _c.Wrap(_eb, _ag, "")
			}
		}
		for _cbc = 1; _cbc < 4; _cbc++ {
			for _bf = 0; _bf < _gfb; _bf++ {
				if _eb = _cfg.SetByte(_eg+_cbc*_gfb+_bf, _cfg.Data[_eg+_bf]); _eb != nil {
					return _c.Wrapf(_eb, _ag, "\u0063\u006f\u0070\u0079\u0020\u0027\u0071\u0075\u0061\u0064\u0072\u0061\u0062l\u0065\u0027\u0020\u006c\u0069\u006ee\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0079\u0074\u0065\u003a \u0027\u0025\u0064\u0027", _cbc, _bf)
				}
			}
		}
	}
	return nil
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

func _dcge(_befa *_fa.Stack, _addg, _bfab, _eccg, _gcde, _geac int, _bgdbb *_g.Rectangle) (_adcb error) {
	const _gfba = "\u0070\u0075\u0073\u0068\u0046\u0069\u006c\u006c\u0053\u0065\u0067m\u0065\u006e\u0074\u0042\u006f\u0075\u006e\u0064\u0069\u006eg\u0042\u006f\u0078"
	if _befa == nil {
		return _c.Error(_gfba, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _bgdbb == nil {
		return _c.Error(_gfba, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0069\u006da\u0067e\u002e\u0052\u0065\u0063\u0074\u0061\u006eg\u006c\u0065")
	}
	_bgdbb.Min.X = _fa.Min(_bgdbb.Min.X, _addg)
	_bgdbb.Max.X = _fa.Max(_bgdbb.Max.X, _bfab)
	_bgdbb.Min.Y = _fa.Min(_bgdbb.Min.Y, _eccg)
	_bgdbb.Max.Y = _fa.Max(_bgdbb.Max.Y, _eccg)
	if !(_eccg+_gcde >= 0 && _eccg+_gcde <= _geac) {
		return nil
	}
	if _befa.Aux == nil {
		return _c.Error(_gfba, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	var _eabga *fillSegment
	_faad, _dca := _befa.Aux.Pop()
	if _dca {
		if _eabga, _dca = _faad.(*fillSegment); !_dca {
			return _c.Error(_gfba, "a\u0075\u0078\u0053\u0074\u0061\u0063k\u0020\u0064\u0061\u0074\u0061\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061 \u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065n\u0074")
		}
	} else {
		_eabga = &fillSegment{}
	}
	_eabga._eddd = _addg
	_eabga._decge = _bfab
	_eabga._dfef = _eccg
	_eabga._bbgg = _gcde
	_befa.Push(_eabga)
	return nil
}

func (_cdce *Bitmap) RemoveBorder(borderSize int) (*Bitmap, error) {
	if borderSize == 0 {
		return _cdce.Copy(), nil
	}
	_bcfc, _cbcg := _cdce.removeBorderGeneral(borderSize, borderSize, borderSize, borderSize)
	if _cbcg != nil {
		return nil, _c.Wrap(_cbcg, "\u0052\u0065\u006do\u0076\u0065\u0042\u006f\u0072\u0064\u0065\u0072", "")
	}
	return _bcfc, nil
}

func (_fbd *Bitmap) GetByte(index int) (byte, error) {
	if index > len(_fbd.Data)-1 || index < 0 {
		return 0, _c.Errorf("\u0047e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _fbd.Data[index], nil
}

func (_gegea *byHeight) Swap(i, j int) {
	_gegea.Values[i], _gegea.Values[j] = _gegea.Values[j], _gegea.Values[i]
	if _gegea.Boxes != nil {
		_gegea.Boxes[i], _gegea.Boxes[j] = _gegea.Boxes[j], _gegea.Boxes[i]
	}
}

type byWidth Bitmaps

func (_deae *Bitmap) connComponentsBB(_egcbd int) (_gebg *Boxes, _fdca error) {
	const _cgag = "\u0042\u0069\u0074ma\u0070\u002e\u0063\u006f\u006e\u006e\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0042"
	if _egcbd != 4 && _egcbd != 8 {
		return nil, _c.Error(_cgag, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _deae.Zero() {
		return &Boxes{}, nil
	}
	_deae.setPadBits(0)
	_fgfg, _fdca := _gbfag(nil, _deae)
	if _fdca != nil {
		return nil, _c.Wrap(_fdca, _cgag, "\u0062\u006d\u0031")
	}
	_efef := &_fa.Stack{}
	_efef.Aux = &_fa.Stack{}
	_gebg = &Boxes{}
	var (
		_eaae, _dfda int
		_dfcad       _g.Point
		_daee        bool
		_bede        *_g.Rectangle
	)
	for {
		if _dfcad, _daee, _fdca = _fgfg.nextOnPixel(_dfda, _eaae); _fdca != nil {
			return nil, _c.Wrap(_fdca, _cgag, "")
		}
		if !_daee {
			break
		}
		if _bede, _fdca = _fdee(_fgfg, _efef, _dfcad.X, _dfcad.Y, _egcbd); _fdca != nil {
			return nil, _c.Wrap(_fdca, _cgag, "")
		}
		if _fdca = _gebg.Add(_bede); _fdca != nil {
			return nil, _c.Wrap(_fdca, _cgag, "")
		}
		_dfda = _dfcad.X
		_eaae = _dfcad.Y
	}
	return _gebg, nil
}

func _afd(_acc *Bitmap, _cab int, _aad []byte) (_fgd *Bitmap, _fb error) {
	const _ecb = "\u0072\u0065\u0064\u0075\u0063\u0065\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0032"
	if _acc == nil {
		return nil, _c.Error(_ecb, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _cab < 1 || _cab > 4 {
		return nil, _c.Error(_ecb, "\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _acc.Height <= 1 {
		return nil, _c.Errorf(_ecb, "\u0073o\u0075\u0072c\u0065\u0020\u0068e\u0069\u0067\u0068\u0074\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061t\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027\u0020-\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", _acc.Height)
	}
	_fgd = New(_acc.Width/2, _acc.Height/2)
	if _aad == nil {
		_aad = _fcgc()
	}
	_ba := _fcac(_acc.RowStride, 2*_fgd.RowStride)
	switch _cab {
	case 1:
		_fb = _afb(_acc, _fgd, _cab, _aad, _ba)
	case 2:
		_fb = _cef(_acc, _fgd, _cab, _aad, _ba)
	case 3:
		_fb = _faeg(_acc, _fgd, _cab, _aad, _ba)
	case 4:
		_fb = _edff(_acc, _fgd, _cab, _aad, _ba)
	}
	if _fb != nil {
		return nil, _fb
	}
	return _fgd, nil
}
func (_aac *Bitmap) CreateTemplate() *Bitmap { return _aac.createTemplate() }
func _acdd(_fadab *Bitmap, _ebgfb *Bitmap, _dabc int) (_gaacb error) {
	const _agdc = "\u0073\u0065\u0065\u0064\u0066\u0069\u006c\u006c\u0042\u0069\u006e\u0061r\u0079\u004c\u006f\u0077"
	_bfcgg := _fcac(_fadab.Height, _ebgfb.Height)
	_eaab := _fcac(_fadab.RowStride, _ebgfb.RowStride)
	switch _dabc {
	case 4:
		_gaacb = _abdb(_fadab, _ebgfb, _bfcgg, _eaab)
	case 8:
		_gaacb = _daf(_fadab, _ebgfb, _bfcgg, _eaab)
	default:
		return _c.Errorf(_agdc, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0034\u0020\u006fr\u0020\u0038\u0020\u002d\u0020i\u0073\u003a \u0027\u0025\u0064\u0027", _dabc)
	}
	if _gaacb != nil {
		return _c.Wrap(_gaacb, _agdc, "")
	}
	return nil
}

func (_ccga *Bitmaps) GetBox(i int) (*_g.Rectangle, error) {
	const _cdda = "\u0047\u0065\u0074\u0042\u006f\u0078"
	if _ccga == nil {
		return nil, _c.Error(_cdda, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074\u006d\u0061\u0070s\u0027")
	}
	if i > len(_ccga.Boxes)-1 {
		return nil, _c.Errorf(_cdda, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _ccga.Boxes[i], nil
}

func _gafgd(_egbbc, _bbbaf int, _ffdfe string) *Selection {
	_eacg := &Selection{Height: _egbbc, Width: _bbbaf, Name: _ffdfe}
	_eacg.Data = make([][]SelectionValue, _egbbc)
	for _ccabf := 0; _ccabf < _egbbc; _ccabf++ {
		_eacg.Data[_ccabf] = make([]SelectionValue, _bbbaf)
	}
	return _eacg
}

func (_gfag *ClassedPoints) xSortFunction() func(_daaf int, _bafc int) bool {
	return func(_bfbf, _acbc int) bool { return _gfag.XAtIndex(_bfbf) < _gfag.XAtIndex(_acbc) }
}

func _gfbcb(_dgfbd *Bitmap, _ebdc, _gade int, _bbaee, _ffge int, _gddc RasterOperator) {
	var (
		_bfdec bool
		_gadfe bool
		_dfdb  int
		_cdbd  int
		_gbfd  int
		_caega int
		_bcdab bool
		_adgg  byte
	)
	_eafb := 8 - (_ebdc & 7)
	_eace := _fgcfb[_eafb]
	_eggae := _dgfbd.RowStride*_gade + (_ebdc >> 3)
	if _bbaee < _eafb {
		_bfdec = true
		_eace &= _abfb[8-_eafb+_bbaee]
	}
	if !_bfdec {
		_dfdb = (_bbaee - _eafb) >> 3
		if _dfdb != 0 {
			_gadfe = true
			_cdbd = _eggae + 1
		}
	}
	_gbfd = (_ebdc + _bbaee) & 7
	if !(_bfdec || _gbfd == 0) {
		_bcdab = true
		_adgg = _abfb[_gbfd]
		_caega = _eggae + 1 + _dfdb
	}
	var _caca, _cgef int
	switch _gddc {
	case PixClr:
		for _caca = 0; _caca < _ffge; _caca++ {
			_dgfbd.Data[_eggae] = _fcbb(_dgfbd.Data[_eggae], 0x0, _eace)
			_eggae += _dgfbd.RowStride
		}
		if _gadfe {
			for _caca = 0; _caca < _ffge; _caca++ {
				for _cgef = 0; _cgef < _dfdb; _cgef++ {
					_dgfbd.Data[_cdbd+_cgef] = 0x0
				}
				_cdbd += _dgfbd.RowStride
			}
		}
		if _bcdab {
			for _caca = 0; _caca < _ffge; _caca++ {
				_dgfbd.Data[_caega] = _fcbb(_dgfbd.Data[_caega], 0x0, _adgg)
				_caega += _dgfbd.RowStride
			}
		}
	case PixSet:
		for _caca = 0; _caca < _ffge; _caca++ {
			_dgfbd.Data[_eggae] = _fcbb(_dgfbd.Data[_eggae], 0xff, _eace)
			_eggae += _dgfbd.RowStride
		}
		if _gadfe {
			for _caca = 0; _caca < _ffge; _caca++ {
				for _cgef = 0; _cgef < _dfdb; _cgef++ {
					_dgfbd.Data[_cdbd+_cgef] = 0xff
				}
				_cdbd += _dgfbd.RowStride
			}
		}
		if _bcdab {
			for _caca = 0; _caca < _ffge; _caca++ {
				_dgfbd.Data[_caega] = _fcbb(_dgfbd.Data[_caega], 0xff, _adgg)
				_caega += _dgfbd.RowStride
			}
		}
	case PixNotDst:
		for _caca = 0; _caca < _ffge; _caca++ {
			_dgfbd.Data[_eggae] = _fcbb(_dgfbd.Data[_eggae], ^_dgfbd.Data[_eggae], _eace)
			_eggae += _dgfbd.RowStride
		}
		if _gadfe {
			for _caca = 0; _caca < _ffge; _caca++ {
				for _cgef = 0; _cgef < _dfdb; _cgef++ {
					_dgfbd.Data[_cdbd+_cgef] = ^(_dgfbd.Data[_cdbd+_cgef])
				}
				_cdbd += _dgfbd.RowStride
			}
		}
		if _bcdab {
			for _caca = 0; _caca < _ffge; _caca++ {
				_dgfbd.Data[_caega] = _fcbb(_dgfbd.Data[_caega], ^_dgfbd.Data[_caega], _adgg)
				_caega += _dgfbd.RowStride
			}
		}
	}
}

func (_ffba *Bitmaps) selectByIndicator(_gabff *_fa.NumSlice) (_bgaae *Bitmaps, _ggfea error) {
	const _agde = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u0073\u0065\u006c\u0065c\u0074B\u0079I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _ffba == nil {
		return nil, _c.Error(_agde, "\u0027\u0062\u0027 b\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _gabff == nil {
		return nil, _c.Error(_agde, "'\u006e\u0061\u0027\u0020\u0069\u006ed\u0069\u0063\u0061\u0074\u006f\u0072\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	if len(_ffba.Values) == 0 {
		return _ffba, nil
	}
	if len(*_gabff) != len(_ffba.Values) {
		return nil, _c.Errorf(_agde, "\u006ea\u0020\u006ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064,\u0020\u0069\u0073\u0020\u0064\u0069\u0066\u0066\u0065re\u006e\u0074\u0020t\u0068\u0061n\u0020\u0062\u0069\u0074\u006d\u0061p\u0073\u003a \u0025\u0064", len(*_gabff), len(_ffba.Values))
	}
	var _ggec, _cdbf, _gfbb int
	for _cdbf = 0; _cdbf < len(*_gabff); _cdbf++ {
		if _ggec, _ggfea = _gabff.GetInt(_cdbf); _ggfea != nil {
			return nil, _c.Wrap(_ggfea, _agde, "f\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		if _ggec == 1 {
			_gfbb++
		}
	}
	if _gfbb == len(_ffba.Values) {
		return _ffba, nil
	}
	_bgaae = &Bitmaps{}
	_dgbdf := len(_ffba.Values) == len(_ffba.Boxes)
	for _cdbf = 0; _cdbf < len(*_gabff); _cdbf++ {
		if _ggec = int((*_gabff)[_cdbf]); _ggec == 0 {
			continue
		}
		_bgaae.Values = append(_bgaae.Values, _ffba.Values[_cdbf])
		if _dgbdf {
			_bgaae.Boxes = append(_bgaae.Boxes, _ffba.Boxes[_cdbf])
		}
	}
	return _bgaae, nil
}

func _cadad() []int {
	_bdaef := make([]int, 256)
	_bdaef[0] = 0
	_bdaef[1] = 7
	var _afcde int
	for _afcde = 2; _afcde < 4; _afcde++ {
		_bdaef[_afcde] = _bdaef[_afcde-2] + 6
	}
	for _afcde = 4; _afcde < 8; _afcde++ {
		_bdaef[_afcde] = _bdaef[_afcde-4] + 5
	}
	for _afcde = 8; _afcde < 16; _afcde++ {
		_bdaef[_afcde] = _bdaef[_afcde-8] + 4
	}
	for _afcde = 16; _afcde < 32; _afcde++ {
		_bdaef[_afcde] = _bdaef[_afcde-16] + 3
	}
	for _afcde = 32; _afcde < 64; _afcde++ {
		_bdaef[_afcde] = _bdaef[_afcde-32] + 2
	}
	for _afcde = 64; _afcde < 128; _afcde++ {
		_bdaef[_afcde] = _bdaef[_afcde-64] + 1
	}
	for _afcde = 128; _afcde < 256; _afcde++ {
		_bdaef[_afcde] = _bdaef[_afcde-128]
	}
	return _bdaef
}

func (_fdf *Bitmap) setAll() error {
	_ddf := _gada(_fdf, 0, 0, _fdf.Width, _fdf.Height, PixSet, nil, 0, 0)
	if _ddf != nil {
		return _c.Wrap(_ddf, "\u0073\u0065\u0074\u0041\u006c\u006c", "")
	}
	return nil
}

func (_ddge *Bitmap) centroid(_egcbde, _ecff []int) (Point, error) {
	_baab := Point{}
	_ddge.setPadBits(0)
	if len(_egcbde) == 0 {
		_egcbde = _cadad()
	}
	if len(_ecff) == 0 {
		_ecff = _ccda()
	}
	var _afcd, _ccf, _gcg, _efbf, _cdad, _agae int
	var _dfgf byte
	for _cdad = 0; _cdad < _ddge.Height; _cdad++ {
		_degg := _ddge.RowStride * _cdad
		_efbf = 0
		for _agae = 0; _agae < _ddge.RowStride; _agae++ {
			_dfgf = _ddge.Data[_degg+_agae]
			if _dfgf != 0 {
				_efbf += _ecff[_dfgf]
				_afcd += _egcbde[_dfgf] + _agae*8*_ecff[_dfgf]
			}
		}
		_gcg += _efbf
		_ccf += _efbf * _cdad
	}
	if _gcg != 0 {
		_baab.X = float32(_afcd) / float32(_gcg)
		_baab.Y = float32(_ccf) / float32(_gcg)
	}
	return _baab, nil
}

type shift int

func _cgde(_gbae *Bitmap, _gddee, _aggdd int, _ffbc, _ebbfe int, _ccab RasterOperator) {
	var (
		_babdb      int
		_gfbca      byte
		_gea, _fbaa int
		_dfed       int
	)
	_fbba := _ffbc >> 3
	_ddee := _ffbc & 7
	if _ddee > 0 {
		_gfbca = _abfb[_ddee]
	}
	_babdb = _gbae.RowStride*_aggdd + (_gddee >> 3)
	switch _ccab {
	case PixClr:
		for _gea = 0; _gea < _ebbfe; _gea++ {
			_dfed = _babdb + _gea*_gbae.RowStride
			for _fbaa = 0; _fbaa < _fbba; _fbaa++ {
				_gbae.Data[_dfed] = 0x0
				_dfed++
			}
			if _ddee > 0 {
				_gbae.Data[_dfed] = _fcbb(_gbae.Data[_dfed], 0x0, _gfbca)
			}
		}
	case PixSet:
		for _gea = 0; _gea < _ebbfe; _gea++ {
			_dfed = _babdb + _gea*_gbae.RowStride
			for _fbaa = 0; _fbaa < _fbba; _fbaa++ {
				_gbae.Data[_dfed] = 0xff
				_dfed++
			}
			if _ddee > 0 {
				_gbae.Data[_dfed] = _fcbb(_gbae.Data[_dfed], 0xff, _gfbca)
			}
		}
	case PixNotDst:
		for _gea = 0; _gea < _ebbfe; _gea++ {
			_dfed = _babdb + _gea*_gbae.RowStride
			for _fbaa = 0; _fbaa < _fbba; _fbaa++ {
				_gbae.Data[_dfed] = ^_gbae.Data[_dfed]
				_dfed++
			}
			if _ddee > 0 {
				_gbae.Data[_dfed] = _fcbb(_gbae.Data[_dfed], ^_gbae.Data[_dfed], _gfbca)
			}
		}
	}
}

type Boxes []*_g.Rectangle

const (
	_ SizeSelection = iota
	SizeSelectByWidth
	SizeSelectByHeight
	SizeSelectByMaxDimension
	SizeSelectByArea
	SizeSelectByPerimeter
)

type SelectionValue int

func SelCreateBrick(h, w int, cy, cx int, tp SelectionValue) *Selection {
	_agbeb := _gafgd(h, w, "")
	_agbeb.setOrigin(cy, cx)
	var _gdbb, _bgfd int
	for _gdbb = 0; _gdbb < h; _gdbb++ {
		for _bgfd = 0; _bgfd < w; _bgfd++ {
			_agbeb.Data[_gdbb][_bgfd] = tp
		}
	}
	return _agbeb
}

func (_dacgb *ClassedPoints) ySortFunction() func(_begce int, _fbcg int) bool {
	return func(_bgc, _bdgf int) bool { return _dacgb.YAtIndex(_bgc) < _dacgb.YAtIndex(_bdgf) }
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

func (_fcgf *Bitmaps) AddBox(box *_g.Rectangle) { _fcgf.Boxes = append(_fcgf.Boxes, box) }
func (_ageb *Boxes) makeSizeIndicator(_bgda, _baca int, _fgef LocationFilter, _cccf SizeComparison) *_fa.NumSlice {
	_gbdf := &_fa.NumSlice{}
	var _gggg, _bbfe, _ddag int
	for _, _gbgb := range *_ageb {
		_gggg = 0
		_bbfe, _ddag = _gbgb.Dx(), _gbgb.Dy()
		switch _fgef {
		case LocSelectWidth:
			if (_cccf == SizeSelectIfLT && _bbfe < _bgda) || (_cccf == SizeSelectIfGT && _bbfe > _bgda) || (_cccf == SizeSelectIfLTE && _bbfe <= _bgda) || (_cccf == SizeSelectIfGTE && _bbfe >= _bgda) {
				_gggg = 1
			}
		case LocSelectHeight:
			if (_cccf == SizeSelectIfLT && _ddag < _baca) || (_cccf == SizeSelectIfGT && _ddag > _baca) || (_cccf == SizeSelectIfLTE && _ddag <= _baca) || (_cccf == SizeSelectIfGTE && _ddag >= _baca) {
				_gggg = 1
			}
		case LocSelectIfEither:
			if (_cccf == SizeSelectIfLT && (_ddag < _baca || _bbfe < _bgda)) || (_cccf == SizeSelectIfGT && (_ddag > _baca || _bbfe > _bgda)) || (_cccf == SizeSelectIfLTE && (_ddag <= _baca || _bbfe <= _bgda)) || (_cccf == SizeSelectIfGTE && (_ddag >= _baca || _bbfe >= _bgda)) {
				_gggg = 1
			}
		case LocSelectIfBoth:
			if (_cccf == SizeSelectIfLT && (_ddag < _baca && _bbfe < _bgda)) || (_cccf == SizeSelectIfGT && (_ddag > _baca && _bbfe > _bgda)) || (_cccf == SizeSelectIfLTE && (_ddag <= _baca && _bbfe <= _bgda)) || (_cccf == SizeSelectIfGTE && (_ddag >= _baca && _bbfe >= _bgda)) {
				_gggg = 1
			}
		}
		_gbdf.AddInt(_gggg)
	}
	return _gbdf
}

func (_fcec *Bitmaps) WidthSorter() func(_bace, _cagg int) bool {
	return func(_bdcdg, _dfcc int) bool { return _fcec.Values[_bdcdg].Width < _fcec.Values[_dfcc].Width }
}

func (_eggcf *Bitmaps) ClipToBitmap(s *Bitmap) (*Bitmaps, error) {
	const _deee = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0043\u006c\u0069p\u0054\u006f\u0042\u0069tm\u0061\u0070"
	if _eggcf == nil {
		return nil, _c.Error(_deee, "\u0042\u0069\u0074\u006dap\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if s == nil {
		return nil, _c.Error(_deee, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	_agffd := len(_eggcf.Values)
	_abcc := &Bitmaps{Values: make([]*Bitmap, _agffd), Boxes: make([]*_g.Rectangle, _agffd)}
	var (
		_cebe, _adagg *Bitmap
		_fceg         *_g.Rectangle
		_egdf         error
	)
	for _ebga := 0; _ebga < _agffd; _ebga++ {
		if _cebe, _egdf = _eggcf.GetBitmap(_ebga); _egdf != nil {
			return nil, _c.Wrap(_egdf, _deee, "")
		}
		if _fceg, _egdf = _eggcf.GetBox(_ebga); _egdf != nil {
			return nil, _c.Wrap(_egdf, _deee, "")
		}
		if _adagg, _egdf = s.clipRectangle(_fceg, nil); _egdf != nil {
			return nil, _c.Wrap(_egdf, _deee, "")
		}
		if _adagg, _egdf = _adagg.And(_cebe); _egdf != nil {
			return nil, _c.Wrap(_egdf, _deee, "")
		}
		_abcc.Values[_ebga] = _adagg
		_abcc.Boxes[_ebga] = _fceg
	}
	return _abcc, nil
}

const (
	AsymmetricMorphBC BoundaryCondition = iota
	SymmetricMorphBC
)

func _fdee(_dgcc *Bitmap, _aaef *_fa.Stack, _gccb, _aceg, _bcdd int) (_cegd *_g.Rectangle, _affg error) {
	const _gcga = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _dgcc == nil {
		return nil, _c.Error(_gcga, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _aaef == nil {
		return nil, _c.Error(_gcga, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	switch _bcdd {
	case 4:
		if _cegd, _affg = _bdegg(_dgcc, _aaef, _gccb, _aceg); _affg != nil {
			return nil, _c.Wrap(_affg, _gcga, "")
		}
		return _cegd, nil
	case 8:
		if _cegd, _affg = _caaf(_dgcc, _aaef, _gccb, _aceg); _affg != nil {
			return nil, _c.Wrap(_affg, _gcga, "")
		}
		return _cegd, nil
	default:
		return nil, _c.Errorf(_gcga, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u0069\u0073 \u006eo\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038\u003a\u0020\u0027\u0025\u0064\u0027", _bcdd)
	}
}

func (_fdd *Bitmap) clipRectangle(_cec, _efga *_g.Rectangle) (_eega *Bitmap, _gcc error) {
	const _dage = "\u0063\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if _cec == nil {
		return nil, _c.Error(_dage, "\u0070r\u006fv\u0069\u0064\u0065\u0064\u0020n\u0069\u006c \u0027\u0062\u006f\u0078\u0027")
	}
	_bdeg, _bef := _fdd.Width, _fdd.Height
	_efdd, _gcc := ClipBoxToRectangle(_cec, _bdeg, _bef)
	if _gcc != nil {
		_fce.Log.Warning("\u0027\u0062ox\u0027\u0020\u0064o\u0065\u0073\u006e\u0027t o\u0076er\u006c\u0061\u0070\u0020\u0062\u0069\u0074ma\u0070\u0020\u0027\u0062\u0027\u003a\u0020%\u0076", _gcc)
		return nil, nil
	}
	_bggg, _ddgb := _efdd.Min.X, _efdd.Min.Y
	_aebgc, _bfg := _efdd.Max.X-_efdd.Min.X, _efdd.Max.Y-_efdd.Min.Y
	_eega = New(_aebgc, _bfg)
	_eega.Text = _fdd.Text
	if _gcc = _eega.RasterOperation(0, 0, _aebgc, _bfg, PixSrc, _fdd, _bggg, _ddgb); _gcc != nil {
		return nil, _c.Wrap(_gcc, _dage, "")
	}
	if _efga != nil {
		*_efga = *_efdd
	}
	return _eega, nil
}

func (_eabg *Bitmap) setEightFullBytes(_cgbg int, _bbg uint64) error {
	if _cgbg+7 > len(_eabg.Data)-1 {
		return _c.Error("\u0073\u0065\u0074\u0045\u0069\u0067\u0068\u0074\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_eabg.Data[_cgbg] = byte((_bbg & 0xff00000000000000) >> 56)
	_eabg.Data[_cgbg+1] = byte((_bbg & 0xff000000000000) >> 48)
	_eabg.Data[_cgbg+2] = byte((_bbg & 0xff0000000000) >> 40)
	_eabg.Data[_cgbg+3] = byte((_bbg & 0xff00000000) >> 32)
	_eabg.Data[_cgbg+4] = byte((_bbg & 0xff000000) >> 24)
	_eabg.Data[_cgbg+5] = byte((_bbg & 0xff0000) >> 16)
	_eabg.Data[_cgbg+6] = byte((_bbg & 0xff00) >> 8)
	_eabg.Data[_cgbg+7] = byte(_bbg & 0xff)
	return nil
}
func (_dcggf *BitmapsArray) AddBitmaps(bm *Bitmaps) { _dcggf.Values = append(_dcggf.Values, bm) }
func _ecbe(_fgecc, _ddec *Bitmap, _gceg, _edefg int) (*Bitmap, error) {
	const _cgbb = "\u0063\u006c\u006f\u0073\u0065\u0053\u0061\u0066\u0065B\u0072\u0069\u0063\u006b"
	if _ddec == nil {
		return nil, _c.Error(_cgbb, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _gceg < 1 || _edefg < 1 {
		return nil, _c.Error(_cgbb, "\u0068s\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _gceg == 1 && _edefg == 1 {
		return _gbfag(_fgecc, _ddec)
	}
	if MorphBC == SymmetricMorphBC {
		_bdec, _bccb := _gcecf(_fgecc, _ddec, _gceg, _edefg)
		if _bccb != nil {
			return nil, _c.Wrap(_bccb, _cgbb, "\u0053\u0079m\u006d\u0065\u0074r\u0069\u0063\u004d\u006f\u0072\u0070\u0068\u0042\u0043")
		}
		return _bdec, nil
	}
	_gbee := _fcdga(_gceg/2, _edefg/2)
	_aega := 8 * ((_gbee + 7) / 8)
	_aadd, _gcgc := _ddec.AddBorder(_aega, 0)
	if _gcgc != nil {
		return nil, _c.Wrapf(_gcgc, _cgbb, "\u0042\u006f\u0072\u0064\u0065\u0072\u0053\u0069\u007ae\u003a\u0020\u0025\u0064", _aega)
	}
	var _agcf, _dbb *Bitmap
	if _gceg == 1 || _edefg == 1 {
		_eafd := SelCreateBrick(_edefg, _gceg, _edefg/2, _gceg/2, SelHit)
		_agcf, _gcgc = _gcgg(nil, _aadd, _eafd)
		if _gcgc != nil {
			return nil, _c.Wrap(_gcgc, _cgbb, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
	} else {
		_egfae := SelCreateBrick(1, _gceg, 0, _gceg/2, SelHit)
		_cada, _adbd := _bag(nil, _aadd, _egfae)
		if _adbd != nil {
			return nil, _c.Wrap(_adbd, _cgbb, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0064\u0069\u006c\u0061t\u0065")
		}
		_gbcg := SelCreateBrick(_edefg, 1, _edefg/2, 0, SelHit)
		_agcf, _adbd = _bag(nil, _cada, _gbcg)
		if _adbd != nil {
			return nil, _c.Wrap(_adbd, _cgbb, "\u0072\u0065\u0067ul\u0061\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
		}
		if _, _adbd = _gcae(_cada, _agcf, _egfae); _adbd != nil {
			return nil, _c.Wrap(_adbd, _cgbb, "r\u0065\u0067\u0075\u006car\u0020-\u0020\u0066\u0069\u0072\u0073t\u0020\u0065\u0072\u006f\u0064\u0065")
		}
		if _, _adbd = _gcae(_agcf, _cada, _gbcg); _adbd != nil {
			return nil, _c.Wrap(_adbd, _cgbb, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0065\u0072\u006fd\u0065")
		}
	}
	if _dbb, _gcgc = _agcf.RemoveBorder(_aega); _gcgc != nil {
		return nil, _c.Wrap(_gcgc, _cgbb, "\u0072e\u0067\u0075\u006c\u0061\u0072")
	}
	if _fgecc == nil {
		return _dbb, nil
	}
	if _, _gcgc = _gbfag(_fgecc, _dbb); _gcgc != nil {
		return nil, _gcgc
	}
	return _fgecc, nil
}
func (_ebdaf *Bitmaps) SortByHeight() { _bfebc := (*byHeight)(_ebdaf); _cc.Sort(_bfebc) }

type Bitmaps struct {
	Values []*Bitmap
	Boxes  []*_g.Rectangle
}

func (_ccb *Bitmap) ToImage() _g.Image {
	_cga, _bdda := _fg.NewImage(_ccb.Width, _ccb.Height, 1, 1, _ccb.Data, nil, nil)
	if _bdda != nil {
		_fce.Log.Error("\u0043\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020j\u0062\u0069\u0067\u0032\u002e\u0042\u0069\u0074m\u0061p\u0020\u0074\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002e\u0049\u006d\u0061\u0067e\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _bdda)
	}
	return _cga
}

func MorphSequence(src *Bitmap, sequence ...MorphProcess) (*Bitmap, error) {
	return _gbfc(src, sequence...)
}

func _dgcd(_cccc, _gcdg, _dffda *Bitmap, _caga int) (*Bitmap, error) {
	const _ebed = "\u0073\u0065\u0065\u0064\u0046\u0069\u006c\u006c\u0042i\u006e\u0061\u0072\u0079"
	if _gcdg == nil {
		return nil, _c.Error(_ebed, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _dffda == nil {
		return nil, _c.Error(_ebed, "'\u006da\u0073\u006b\u0027\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _caga != 4 && _caga != 8 {
		return nil, _c.Error(_ebed, "\u0063\u006f\u006en\u0065\u0063\u0074\u0069v\u0069\u0074\u0079\u0020\u006e\u006f\u0074 \u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0034\u002c\u0038\u007d")
	}
	var _addc error
	_cccc, _addc = _gbfag(_cccc, _gcdg)
	if _addc != nil {
		return nil, _c.Wrap(_addc, _ebed, "\u0063o\u0070y\u0020\u0073\u006f\u0075\u0072c\u0065\u0020t\u006f\u0020\u0027\u0064\u0027")
	}
	_fegd := _gcdg.createTemplate()
	_dffda.setPadBits(0)
	for _dggd := 0; _dggd < _fceea; _dggd++ {
		_fegd, _addc = _gbfag(_fegd, _cccc)
		if _addc != nil {
			return nil, _c.Wrapf(_addc, _ebed, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _dggd)
		}
		if _addc = _acdd(_cccc, _dffda, _caga); _addc != nil {
			return nil, _c.Wrapf(_addc, _ebed, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _dggd)
		}
		if _fegd.Equals(_cccc) {
			break
		}
	}
	return _cccc, nil
}

func (_gegga *Bitmaps) selectByIndexes(_abea []int) (*Bitmaps, error) {
	_dbgcb := &Bitmaps{}
	for _, _bffg := range _abea {
		_fafbd, _dfdf := _gegga.GetBitmap(_bffg)
		if _dfdf != nil {
			return nil, _c.Wrap(_dfdf, "\u0073e\u006ce\u0063\u0074\u0042\u0079\u0049\u006e\u0064\u0065\u0078\u0065\u0073", "")
		}
		_dbgcb.AddBitmap(_fafbd)
	}
	return _dbgcb, nil
}

func _ae(_gfbc *Bitmap, _ee *Bitmap, _bc int) (_eee error) {
	const _def = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0050\u006fw\u0065\u0072\u0032\u004c\u006f\u0077"
	switch _bc {
	case 2:
		_eee = _ec(_gfbc, _ee)
	case 4:
		_eee = _gb(_gfbc, _ee)
	case 8:
		_eee = _ebg(_gfbc, _ee)
	default:
		return _c.Error(_def, "\u0065\u0078p\u0061\u006e\u0073\u0069o\u006e\u0020f\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020r\u0061\u006e\u0067\u0065")
	}
	if _eee != nil {
		_eee = _c.Wrap(_eee, _def, "")
	}
	return _eee
}
func (_ceca *Bitmap) setBit(_gdge int)    { _ceca.Data[(_gdge >> 3)] |= 0x80 >> uint(_gdge&7) }
func MakePixelCentroidTab8() []int        { return _cadad() }
func (_abd *Bitmap) SetPadBits(value int) { _abd.setPadBits(value) }
func (_aabdf *Bitmaps) GetBitmap(i int) (*Bitmap, error) {
	const _gcad = "\u0047e\u0074\u0042\u0069\u0074\u006d\u0061p"
	if _aabdf == nil {
		return nil, _c.Error(_gcad, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if i > len(_aabdf.Values)-1 {
		return nil, _c.Errorf(_gcad, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _aabdf.Values[i], nil
}

func (_gfc *Bitmap) And(s *Bitmap) (_bfcc *Bitmap, _ecbd error) {
	const _bcb = "\u0042\u0069\u0074\u006d\u0061\u0070\u002e\u0041\u006e\u0064"
	if _gfc == nil {
		return nil, _c.Error(_bcb, "\u0027b\u0069t\u006d\u0061\u0070\u0020\u0027b\u0027\u0020i\u0073\u0020\u006e\u0069\u006c")
	}
	if s == nil {
		return nil, _c.Error(_bcb, "\u0062\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if !_gfc.SizesEqual(s) {
		_fce.Log.Debug("\u0025\u0073\u0020-\u0020\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0069\u007a\u0065 \u0077\u0069\u0074\u0068\u0020\u0027\u0062\u0027", _bcb)
	}
	if _bfcc, _ecbd = _gbfag(_bfcc, _gfc); _ecbd != nil {
		return nil, _c.Wrap(_ecbd, _bcb, "\u0063\u0061\u006e't\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _ecbd = _bfcc.RasterOperation(0, 0, _bfcc.Width, _bfcc.Height, PixSrcAndDst, s, 0, 0); _ecbd != nil {
		return nil, _c.Wrap(_ecbd, _bcb, "")
	}
	return _bfcc, nil
}

func _cef(_bde, _edba *Bitmap, _ada int, _dbd []byte, _faee int) (_ggdc error) {
	const _aed = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0032"
	var (
		_faec, _agf, _ddg, _egc, _eaa, _adf, _cbf, _bcfa int
		_aedb, _bcd, _facd, _ggdcf                       uint32
		_ceb, _gbe                                       byte
		_cda                                             uint16
	)
	_gcd := make([]byte, 4)
	_fad := make([]byte, 4)
	for _ddg = 0; _ddg < _bde.Height-1; _ddg, _egc = _ddg+2, _egc+1 {
		_faec = _ddg * _bde.RowStride
		_agf = _egc * _edba.RowStride
		for _eaa, _adf = 0, 0; _eaa < _faee; _eaa, _adf = _eaa+4, _adf+1 {
			for _cbf = 0; _cbf < 4; _cbf++ {
				_bcfa = _faec + _eaa + _cbf
				if _bcfa <= len(_bde.Data)-1 && _bcfa < _faec+_bde.RowStride {
					_gcd[_cbf] = _bde.Data[_bcfa]
				} else {
					_gcd[_cbf] = 0x00
				}
				_bcfa = _faec + _bde.RowStride + _eaa + _cbf
				if _bcfa <= len(_bde.Data)-1 && _bcfa < _faec+(2*_bde.RowStride) {
					_fad[_cbf] = _bde.Data[_bcfa]
				} else {
					_fad[_cbf] = 0x00
				}
			}
			_aedb = _a.BigEndian.Uint32(_gcd)
			_bcd = _a.BigEndian.Uint32(_fad)
			_facd = _aedb & _bcd
			_facd |= _facd << 1
			_ggdcf = _aedb | _bcd
			_ggdcf &= _ggdcf << 1
			_bcd = _facd | _ggdcf
			_bcd &= 0xaaaaaaaa
			_aedb = _bcd | (_bcd << 7)
			_ceb = byte(_aedb >> 24)
			_gbe = byte((_aedb >> 8) & 0xff)
			_bcfa = _agf + _adf
			if _bcfa+1 == len(_edba.Data)-1 || _bcfa+1 >= _agf+_edba.RowStride {
				if _ggdc = _edba.SetByte(_bcfa, _dbd[_ceb]); _ggdc != nil {
					return _c.Wrapf(_ggdc, _aed, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _bcfa)
				}
			} else {
				_cda = (uint16(_dbd[_ceb]) << 8) | uint16(_dbd[_gbe])
				if _ggdc = _edba.setTwoBytes(_bcfa, _cda); _ggdc != nil {
					return _c.Wrapf(_ggdc, _aed, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _bcfa)
				}
				_adf++
			}
		}
	}
	return nil
}
func (_bcec *ClassedPoints) XAtIndex(i int) float32 { return (*_bcec.Points)[_bcec.IntSlice[i]].X }
func (_bgb *Bitmap) GetByteIndex(x, y int) int      { return y*_bgb.RowStride + (x >> 3) }
func (_agaa *Selection) findMaxTranslations() (_aabe, _eaed, _dgfa, _ffcc int) {
	for _gffbe := 0; _gffbe < _agaa.Height; _gffbe++ {
		for _gdbda := 0; _gdbda < _agaa.Width; _gdbda++ {
			if _agaa.Data[_gffbe][_gdbda] == SelHit {
				_aabe = _fcdga(_aabe, _agaa.Cx-_gdbda)
				_eaed = _fcdga(_eaed, _agaa.Cy-_gffbe)
				_dgfa = _fcdga(_dgfa, _gdbda-_agaa.Cx)
				_ffcc = _fcdga(_ffcc, _gffbe-_agaa.Cy)
			}
		}
	}
	return _aabe, _eaed, _dgfa, _ffcc
}

func (_cebc *Bitmap) Zero() bool {
	_fffg := _cebc.Width / 8
	_fafb := _cebc.Width & 7
	var _bgfc byte
	if _fafb != 0 {
		_bgfc = byte(0xff << uint(8-_fafb))
	}
	var _adad, _bgd, _afba int
	for _bgd = 0; _bgd < _cebc.Height; _bgd++ {
		_adad = _cebc.RowStride * _bgd
		for _afba = 0; _afba < _fffg; _afba, _adad = _afba+1, _adad+1 {
			if _cebc.Data[_adad] != 0 {
				return false
			}
		}
		if _fafb > 0 {
			if _cebc.Data[_adad]&_bgfc != 0 {
				return false
			}
		}
	}
	return true
}

func (_fca *Bitmap) clearAll() error {
	return _fca.RasterOperation(0, 0, _fca.Width, _fca.Height, PixClr, nil, 0, 0)
}

func _ebg(_edc, _dbg *Bitmap) (_age error) {
	const _dba = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0038"
	_ac := _dbg.RowStride
	_egb := _edc.RowStride
	var _df, _gfe, _fgb, _cgb, _bfe int
	for _fgb = 0; _fgb < _dbg.Height; _fgb++ {
		_df = _fgb * _ac
		_gfe = 8 * _fgb * _egb
		for _cgb = 0; _cgb < _ac; _cgb++ {
			if _age = _edc.setEightBytes(_gfe+_cgb*8, _aebb[_dbg.Data[_df+_cgb]]); _age != nil {
				return _c.Wrap(_age, _dba, "")
			}
		}
		for _bfe = 1; _bfe < 8; _bfe++ {
			for _cgb = 0; _cgb < _egb; _cgb++ {
				if _age = _edc.SetByte(_gfe+_bfe*_egb+_cgb, _edc.Data[_gfe+_cgb]); _age != nil {
					return _c.Wrap(_age, _dba, "")
				}
			}
		}
	}
	return nil
}

func (_ggdd *Bitmap) SetDefaultPixel() {
	for _aef := range _ggdd.Data {
		_ggdd.Data[_aef] = byte(0xff)
	}
}

func _caaf(_cba *Bitmap, _bgde *_fa.Stack, _gefbe, _gbfdf int) (_daed *_g.Rectangle, _ecab error) {
	const _acgad = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _cba == nil {
		return nil, _c.Error(_acgad, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _bgde == nil {
		return nil, _c.Error(_acgad, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_abbg, _cbac := _cba.Width, _cba.Height
	_ggeb := _abbg - 1
	_gffg := _cbac - 1
	if _gefbe < 0 || _gefbe > _ggeb || _gbfdf < 0 || _gbfdf > _gffg || !_cba.GetPixel(_gefbe, _gbfdf) {
		return nil, nil
	}
	_afdaf := _g.Rect(100000, 100000, 0, 0)
	if _ecab = _dcge(_bgde, _gefbe, _gefbe, _gbfdf, 1, _gffg, &_afdaf); _ecab != nil {
		return nil, _c.Wrap(_ecab, _acgad, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _ecab = _dcge(_bgde, _gefbe, _gefbe, _gbfdf+1, -1, _gffg, &_afdaf); _ecab != nil {
		return nil, _c.Wrap(_ecab, _acgad, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_afdaf.Min.X, _afdaf.Max.X = _gefbe, _gefbe
	_afdaf.Min.Y, _afdaf.Max.Y = _gbfdf, _gbfdf
	var (
		_adde *fillSegment
		_abgb int
	)
	for _bgde.Len() > 0 {
		if _adde, _ecab = _gadgd(_bgde); _ecab != nil {
			return nil, _c.Wrap(_ecab, _acgad, "")
		}
		_gbfdf = _adde._dfef
		for _gefbe = _adde._eddd - 1; _gefbe >= 0 && _cba.GetPixel(_gefbe, _gbfdf); _gefbe-- {
			if _ecab = _cba.SetPixel(_gefbe, _gbfdf, 0); _ecab != nil {
				return nil, _c.Wrap(_ecab, _acgad, "\u0031s\u0074\u0020\u0073\u0065\u0074")
			}
		}
		if _gefbe >= _adde._eddd-1 {
			for {
				for _gefbe++; _gefbe <= _adde._decge+1 && _gefbe <= _ggeb && !_cba.GetPixel(_gefbe, _gbfdf); _gefbe++ {
				}
				_abgb = _gefbe
				if !(_gefbe <= _adde._decge+1 && _gefbe <= _ggeb) {
					break
				}
				for ; _gefbe <= _ggeb && _cba.GetPixel(_gefbe, _gbfdf); _gefbe++ {
					if _ecab = _cba.SetPixel(_gefbe, _gbfdf, 0); _ecab != nil {
						return nil, _c.Wrap(_ecab, _acgad, "\u0032n\u0064\u0020\u0073\u0065\u0074")
					}
				}
				if _ecab = _dcge(_bgde, _abgb, _gefbe-1, _adde._dfef, _adde._bbgg, _gffg, &_afdaf); _ecab != nil {
					return nil, _c.Wrap(_ecab, _acgad, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
				}
				if _gefbe > _adde._decge {
					if _ecab = _dcge(_bgde, _adde._decge+1, _gefbe-1, _adde._dfef, -_adde._bbgg, _gffg, &_afdaf); _ecab != nil {
						return nil, _c.Wrap(_ecab, _acgad, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
					}
				}
			}
			continue
		}
		_abgb = _gefbe + 1
		if _abgb < _adde._eddd {
			if _ecab = _dcge(_bgde, _abgb, _adde._eddd-1, _adde._dfef, -_adde._bbgg, _gffg, &_afdaf); _ecab != nil {
				return nil, _c.Wrap(_ecab, _acgad, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
			}
		}
		_gefbe = _adde._eddd
		for {
			for ; _gefbe <= _ggeb && _cba.GetPixel(_gefbe, _gbfdf); _gefbe++ {
				if _ecab = _cba.SetPixel(_gefbe, _gbfdf, 0); _ecab != nil {
					return nil, _c.Wrap(_ecab, _acgad, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _ecab = _dcge(_bgde, _abgb, _gefbe-1, _adde._dfef, _adde._bbgg, _gffg, &_afdaf); _ecab != nil {
				return nil, _c.Wrap(_ecab, _acgad, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _gefbe > _adde._decge {
				if _ecab = _dcge(_bgde, _adde._decge+1, _gefbe-1, _adde._dfef, -_adde._bbgg, _gffg, &_afdaf); _ecab != nil {
					return nil, _c.Wrap(_ecab, _acgad, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _gefbe++; _gefbe <= _adde._decge+1 && _gefbe <= _ggeb && !_cba.GetPixel(_gefbe, _gbfdf); _gefbe++ {
			}
			_abgb = _gefbe
			if !(_gefbe <= _adde._decge+1 && _gefbe <= _ggeb) {
				break
			}
		}
	}
	_afdaf.Max.X++
	_afdaf.Max.Y++
	return &_afdaf, nil
}

func _abdb(_aadcge, _fbbd *Bitmap, _dcbbf, _bbba int) (_gbgg error) {
	const _fdfc = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0034"
	var (
		_gcee, _ffece, _cdfb, _accf                         int
		_ddacd, _cfec, _ffgd, _gcaed, _decdd, _bcffc, _ccgb byte
	)
	for _gcee = 0; _gcee < _dcbbf; _gcee++ {
		_cdfb = _gcee * _aadcge.RowStride
		_accf = _gcee * _fbbd.RowStride
		for _ffece = 0; _ffece < _bbba; _ffece++ {
			_ddacd, _gbgg = _aadcge.GetByte(_cdfb + _ffece)
			if _gbgg != nil {
				return _c.Wrap(_gbgg, _fdfc, "\u0066i\u0072\u0073\u0074\u0020\u0067\u0065t")
			}
			_cfec, _gbgg = _fbbd.GetByte(_accf + _ffece)
			if _gbgg != nil {
				return _c.Wrap(_gbgg, _fdfc, "\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0067\u0065\u0074")
			}
			if _gcee > 0 {
				_ffgd, _gbgg = _aadcge.GetByte(_cdfb - _aadcge.RowStride + _ffece)
				if _gbgg != nil {
					return _c.Wrap(_gbgg, _fdfc, "\u0069\u0020\u003e \u0030")
				}
				_ddacd |= _ffgd
			}
			if _ffece > 0 {
				_gcaed, _gbgg = _aadcge.GetByte(_cdfb + _ffece - 1)
				if _gbgg != nil {
					return _c.Wrap(_gbgg, _fdfc, "\u006a\u0020\u003e \u0030")
				}
				_ddacd |= _gcaed << 7
			}
			_ddacd &= _cfec
			if _ddacd == 0 || (^_ddacd) == 0 {
				if _gbgg = _aadcge.SetByte(_cdfb+_ffece, _ddacd); _gbgg != nil {
					return _c.Wrap(_gbgg, _fdfc, "b\u0074\u0020\u003d\u003d 0\u0020|\u007c\u0020\u0028\u005e\u0062t\u0029\u0020\u003d\u003d\u0020\u0030")
				}
				continue
			}
			for {
				_ccgb = _ddacd
				_ddacd = (_ddacd | (_ddacd >> 1) | (_ddacd << 1)) & _cfec
				if (_ddacd ^ _ccgb) == 0 {
					if _gbgg = _aadcge.SetByte(_cdfb+_ffece, _ddacd); _gbgg != nil {
						return _c.Wrap(_gbgg, _fdfc, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _gcee = _dcbbf - 1; _gcee >= 0; _gcee-- {
		_cdfb = _gcee * _aadcge.RowStride
		_accf = _gcee * _fbbd.RowStride
		for _ffece = _bbba - 1; _ffece >= 0; _ffece-- {
			if _ddacd, _gbgg = _aadcge.GetByte(_cdfb + _ffece); _gbgg != nil {
				return _c.Wrap(_gbgg, _fdfc, "\u0072\u0065\u0076\u0065\u0072\u0073\u0065\u0020\u0066\u0069\u0072\u0073t\u0020\u0067\u0065\u0074")
			}
			if _cfec, _gbgg = _fbbd.GetByte(_accf + _ffece); _gbgg != nil {
				return _c.Wrap(_gbgg, _fdfc, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _gcee < _dcbbf-1 {
				if _decdd, _gbgg = _aadcge.GetByte(_cdfb + _aadcge.RowStride + _ffece); _gbgg != nil {
					return _c.Wrap(_gbgg, _fdfc, "\u0072\u0065v\u0065\u0072\u0073e\u0020\u0069\u0020\u003c\u0020\u0068\u0020\u002d\u0031")
				}
				_ddacd |= _decdd
			}
			if _ffece < _bbba-1 {
				if _bcffc, _gbgg = _aadcge.GetByte(_cdfb + _ffece + 1); _gbgg != nil {
					return _c.Wrap(_gbgg, _fdfc, "\u0072\u0065\u0076\u0065rs\u0065\u0020\u006a\u0020\u003c\u0020\u0077\u0070\u006c\u0020\u002d\u0020\u0031")
				}
				_ddacd |= _bcffc >> 7
			}
			_ddacd &= _cfec
			if _ddacd == 0 || (^_ddacd) == 0 {
				if _gbgg = _aadcge.SetByte(_cdfb+_ffece, _ddacd); _gbgg != nil {
					return _c.Wrap(_gbgg, _fdfc, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006d\u0061\u0073k\u0065\u0064\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061i\u006c\u0065\u0064")
				}
				continue
			}
			for {
				_ccgb = _ddacd
				_ddacd = (_ddacd | (_ddacd >> 1) | (_ddacd << 1)) & _cfec
				if (_ddacd ^ _ccgb) == 0 {
					if _gbgg = _aadcge.SetByte(_cdfb+_ffece, _ddacd); _gbgg != nil {
						return _c.Wrap(_gbgg, _fdfc, "\u0072e\u0076\u0065\u0072\u0073e\u0020\u0073\u0065\u0074\u0074i\u006eg\u0020p\u0072\u0065\u0076\u0020\u0062\u0079\u0074e")
					}
					break
				}
			}
		}
	}
	return nil
}
func Copy(d, s *Bitmap) (*Bitmap, error) { return _gbfag(d, s) }
func _cbcgg(_cfgf, _gfac *Bitmap, _daeg CombinationOperator) *Bitmap {
	_bccf := New(_cfgf.Width, _cfgf.Height)
	for _cdf := 0; _cdf < len(_bccf.Data); _cdf++ {
		_bccf.Data[_cdf] = _fbfd(_cfgf.Data[_cdf], _gfac.Data[_cdf], _daeg)
	}
	return _bccf
}

func _acb(_efeg *Bitmap, _ccc int) (*Bitmap, error) {
	const _gfef = "\u0065x\u0070a\u006e\u0064\u0042\u0069\u006ea\u0072\u0079P\u006f\u0077\u0065\u0072\u0032"
	if _efeg == nil {
		return nil, _c.Error(_gfef, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ccc == 1 {
		return _gbfag(nil, _efeg)
	}
	if _ccc != 2 && _ccc != 4 && _ccc != 8 {
		return nil, _c.Error(_gfef, "\u0066\u0061\u0063t\u006f\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069n\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_de := _ccc * _efeg.Width
	_ebe := _ccc * _efeg.Height
	_eda := New(_de, _ebe)
	var _fd error
	switch _ccc {
	case 2:
		_fd = _ec(_eda, _efeg)
	case 4:
		_fd = _gb(_eda, _efeg)
	case 8:
		_fd = _ebg(_eda, _efeg)
	}
	if _fd != nil {
		return nil, _c.Wrap(_fd, _gfef, "")
	}
	return _eda, nil
}

func CorrelationScoreSimple(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_efdde float64, _adcd error) {
	const _aaff = "\u0043\u006f\u0072\u0072el\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0053\u0069\u006d\u0070l\u0065"
	if bm1 == nil || bm2 == nil {
		return _efdde, _c.Error(_aaff, "n\u0069l\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0073 \u0070\u0072\u006f\u0076id\u0065\u0064")
	}
	if tab == nil {
		return _efdde, _c.Error(_aaff, "\u0074\u0061\u0062\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if area1 == 0 || area2 == 0 {
		return _efdde, _c.Error(_aaff, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0061\u0072e\u0061\u0073\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u0020\u0030")
	}
	_dgce, _cccb := bm1.Width, bm1.Height
	_cgfb, _bcfca := bm2.Width, bm2.Height
	if _acef(_dgce-_cgfb) > maxDiffW {
		return 0, nil
	}
	if _acef(_cccb-_bcfca) > maxDiffH {
		return 0, nil
	}
	var _agff, _eagf int
	if delX >= 0 {
		_agff = int(delX + 0.5)
	} else {
		_agff = int(delX - 0.5)
	}
	if delY >= 0 {
		_eagf = int(delY + 0.5)
	} else {
		_eagf = int(delY - 0.5)
	}
	_bfb := bm1.createTemplate()
	if _adcd = _bfb.RasterOperation(_agff, _eagf, _cgfb, _bcfca, PixSrc, bm2, 0, 0); _adcd != nil {
		return _efdde, _c.Wrap(_adcd, _aaff, "\u0062m\u0032 \u0074\u006f\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _adcd = _bfb.RasterOperation(0, 0, _dgce, _cccb, PixSrcAndDst, bm1, 0, 0); _adcd != nil {
		return _efdde, _c.Wrap(_adcd, _aaff, "b\u006d\u0031\u0020\u0061\u006e\u0064\u0020\u0062\u006d\u0054")
	}
	_bfdf := _bfb.countPixels()
	_efdde = float64(_bfdf) * float64(_bfdf) / (float64(area1) * float64(area2))
	return _efdde, nil
}

var _ggcd = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x78, 0x27, 0xC2, 0x27, 0x91, 0x00, 0x22, 0x48, 0x21, 0x03, 0x24, 0x91, 0x00, 0x22, 0x48, 0x21, 0x02, 0xA4, 0x95, 0x00, 0x22, 0x48, 0x21, 0x02, 0x64, 0x9B, 0x00, 0x3C, 0x78, 0x21, 0x02, 0x27, 0x91, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func Rect(x, y, w, h int) (*_g.Rectangle, error) {
	const _cade = "b\u0069\u0074\u006d\u0061\u0070\u002e\u0052\u0065\u0063\u0074"
	if x < 0 {
		w += x
		x = 0
		if w <= 0 {
			return nil, _c.Errorf(_cade, "x\u003a\u0027\u0025\u0064\u0027\u0020<\u0020\u0030\u0020\u0061\u006e\u0064\u0020\u0077\u003a \u0027\u0025\u0064'\u0020<\u003d\u0020\u0030", x, w)
		}
	}
	if y < 0 {
		h += y
		y = 0
		if h <= 0 {
			return nil, _c.Error(_cade, "\u0079\u0020\u003c 0\u0020\u0061\u006e\u0064\u0020\u0062\u006f\u0078\u0020\u006f\u0066\u0066\u0020\u002b\u0071\u0075\u0061\u0064")
		}
	}
	_fgfa := _g.Rect(x, y, x+w, y+h)
	return &_fgfa, nil
}

const (
	_ SizeComparison = iota
	SizeSelectIfLT
	SizeSelectIfGT
	SizeSelectIfLTE
	SizeSelectIfGTE
	SizeSelectIfEQ
)

func (_gaf *Bitmap) GetBitOffset(x int) int { return x & 0x07 }
func (_aaac *Bitmap) setFourBytes(_eed int, _edfa uint32) error {
	if _eed+3 > len(_aaac.Data)-1 {
		return _c.Errorf("\u0073\u0065\u0074F\u006f\u0075\u0072\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _eed)
	}
	_aaac.Data[_eed] = byte((_edfa & 0xff000000) >> 24)
	_aaac.Data[_eed+1] = byte((_edfa & 0xff0000) >> 16)
	_aaac.Data[_eed+2] = byte((_edfa & 0xff00) >> 8)
	_aaac.Data[_eed+3] = byte(_edfa & 0xff)
	return nil
}

func (_gced *Bitmap) RemoveBorderGeneral(left, right, top, bot int) (*Bitmap, error) {
	return _gced.removeBorderGeneral(left, right, top, bot)
}

func _bcde(_dagea *Bitmap, _efaa, _facb, _geeag, _geeg int, _gefe RasterOperator, _febc *Bitmap, _fbce, _dfeeb int) error {
	var (
		_aggd      bool
		_febb      bool
		_bdgbe     byte
		_ggce      int
		_ccae      int
		_fgefg     int
		_ecec      int
		_aebf      bool
		_bfgc      int
		_bcda      int
		_efcad     int
		_bcbg      bool
		_begf      byte
		_deeb      int
		_afbb      int
		_ggggf     int
		_efgab     byte
		_fea       int
		_eged      int
		_efce      uint
		_gbad      uint
		_ggga      byte
		_fgfgc     shift
		_abdg      bool
		_bcad      bool
		_ffb, _dbe int
	)
	if _fbce&7 != 0 {
		_eged = 8 - (_fbce & 7)
	}
	if _efaa&7 != 0 {
		_ccae = 8 - (_efaa & 7)
	}
	if _eged == 0 && _ccae == 0 {
		_ggga = _fgcfb[0]
	} else {
		if _ccae > _eged {
			_efce = uint(_ccae - _eged)
		} else {
			_efce = uint(8 - (_eged - _ccae))
		}
		_gbad = 8 - _efce
		_ggga = _fgcfb[_efce]
	}
	if (_efaa & 7) != 0 {
		_aggd = true
		_ggce = 8 - (_efaa & 7)
		_bdgbe = _fgcfb[_ggce]
		_fgefg = _dagea.RowStride*_facb + (_efaa >> 3)
		_ecec = _febc.RowStride*_dfeeb + (_fbce >> 3)
		_fea = 8 - (_fbce & 7)
		if _ggce > _fea {
			_fgfgc = _aaddc
			if _geeag >= _eged {
				_abdg = true
			}
		} else {
			_fgfgc = _dded
		}
	}
	if _geeag < _ggce {
		_febb = true
		_bdgbe &= _abfb[8-_ggce+_geeag]
	}
	if !_febb {
		_bfgc = (_geeag - _ggce) >> 3
		if _bfgc != 0 {
			_aebf = true
			_bcda = _dagea.RowStride*_facb + ((_efaa + _ccae) >> 3)
			_efcad = _febc.RowStride*_dfeeb + ((_fbce + _ccae) >> 3)
		}
	}
	_deeb = (_efaa + _geeag) & 7
	if !(_febb || _deeb == 0) {
		_bcbg = true
		_begf = _abfb[_deeb]
		_afbb = _dagea.RowStride*_facb + ((_efaa + _ccae) >> 3) + _bfgc
		_ggggf = _febc.RowStride*_dfeeb + ((_fbce + _ccae) >> 3) + _bfgc
		if _deeb > int(_gbad) {
			_bcad = true
		}
	}
	switch _gefe {
	case PixSrc:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], _efgab, _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] = _efgab
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], _efgab, _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixNotSrc:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], ^_efgab, _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] = ^_efgab
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], ^_efgab, _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixSrcOrDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], _efgab|_dagea.Data[_fgefg], _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] |= _efgab
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], _efgab|_dagea.Data[_afbb], _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixSrcAndDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], _efgab&_dagea.Data[_fgefg], _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] &= _efgab
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], _efgab&_dagea.Data[_afbb], _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixSrcXorDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], _efgab^_dagea.Data[_fgefg], _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] ^= _efgab
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], _efgab^_dagea.Data[_afbb], _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixNotSrcOrDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], ^_efgab|_dagea.Data[_fgefg], _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] |= ^_efgab
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], ^_efgab|_dagea.Data[_afbb], _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixNotSrcAndDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], ^_efgab&_dagea.Data[_fgefg], _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] &= ^_efgab
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], ^_efgab&_dagea.Data[_afbb], _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixSrcOrNotDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], _efgab|^_dagea.Data[_fgefg], _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] = _efgab | ^_dagea.Data[_bcda+_dbe]
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], _efgab|^_dagea.Data[_afbb], _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixSrcAndNotDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], _efgab&^_dagea.Data[_fgefg], _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] = _efgab &^ _dagea.Data[_bcda+_dbe]
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], _efgab&^_dagea.Data[_afbb], _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], ^(_efgab | _dagea.Data[_fgefg]), _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] = ^(_efgab | _dagea.Data[_bcda+_dbe])
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], ^(_efgab | _dagea.Data[_afbb]), _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], ^(_efgab & _dagea.Data[_fgefg]), _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] = ^(_efgab & _dagea.Data[_bcda+_dbe])
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], ^(_efgab & _dagea.Data[_afbb]), _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		if _aggd {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				if _fgfgc == _aaddc {
					_efgab = _febc.Data[_ecec] << _efce
					if _abdg {
						_efgab = _fcbb(_efgab, _febc.Data[_ecec+1]>>_gbad, _ggga)
					}
				} else {
					_efgab = _febc.Data[_ecec] >> _gbad
				}
				_dagea.Data[_fgefg] = _fcbb(_dagea.Data[_fgefg], ^(_efgab ^ _dagea.Data[_fgefg]), _bdgbe)
				_fgefg += _dagea.RowStride
				_ecec += _febc.RowStride
			}
		}
		if _aebf {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				for _dbe = 0; _dbe < _bfgc; _dbe++ {
					_efgab = _fcbb(_febc.Data[_efcad+_dbe]<<_efce, _febc.Data[_efcad+_dbe+1]>>_gbad, _ggga)
					_dagea.Data[_bcda+_dbe] = ^(_efgab ^ _dagea.Data[_bcda+_dbe])
				}
				_bcda += _dagea.RowStride
				_efcad += _febc.RowStride
			}
		}
		if _bcbg {
			for _ffb = 0; _ffb < _geeg; _ffb++ {
				_efgab = _febc.Data[_ggggf] << _efce
				if _bcad {
					_efgab = _fcbb(_efgab, _febc.Data[_ggggf+1]>>_gbad, _ggga)
				}
				_dagea.Data[_afbb] = _fcbb(_dagea.Data[_afbb], ^(_efgab ^ _dagea.Data[_afbb]), _begf)
				_afbb += _dagea.RowStride
				_ggggf += _febc.RowStride
			}
		}
	default:
		_fce.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _gefe)
		return _c.Error("\u0072a\u0073t\u0065\u0072\u004f\u0070\u0047e\u006e\u0065r\u0061\u006c\u004c\u006f\u0077", "\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}

func init() {
	const _bffe = "\u0062\u0069\u0074\u006dap\u0073\u002e\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0061\u0074\u0069o\u006e"
	_dbdg = New(50, 40)
	var _fbgc error
	_dbdg, _fbgc = _dbdg.AddBorder(2, 1)
	if _fbgc != nil {
		panic(_c.Wrap(_fbgc, _bffe, "f\u0072\u0061\u006d\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
	_daff, _fbgc = NewWithData(50, 22, _ggcd)
	if _fbgc != nil {
		panic(_c.Wrap(_fbgc, _bffe, "i\u006d\u0061\u0067\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
}

func _bccg(_cff, _dacb *Bitmap, _bee, _ccaf, _dfe, _babb, _cbcb int, _cffb CombinationOperator) error {
	var _edfe int
	_acf := func() { _edfe++; _dfe += _dacb.RowStride; _babb += _cff.RowStride; _cbcb += _cff.RowStride }
	for _edfe = _bee; _edfe < _ccaf; _acf() {
		_aggg := _dfe
		for _cdg := _babb; _cdg <= _cbcb; _cdg++ {
			_gggd, _begc := _dacb.GetByte(_aggg)
			if _begc != nil {
				return _begc
			}
			_eeca, _begc := _cff.GetByte(_cdg)
			if _begc != nil {
				return _begc
			}
			if _begc = _dacb.SetByte(_aggg, _fbfd(_gggd, _eeca, _cffb)); _begc != nil {
				return _begc
			}
			_aggg++
		}
	}
	return nil
}

func (_fdec *Bitmap) createTemplate() *Bitmap {
	return &Bitmap{Width: _fdec.Width, Height: _fdec.Height, RowStride: _fdec.RowStride, Color: _fdec.Color, Text: _fdec.Text, BitmapNumber: _fdec.BitmapNumber, Special: _fdec.Special, Data: make([]byte, len(_fdec.Data))}
}

func _gcecf(_cfdg, _decdf *Bitmap, _gadc, _egag int) (*Bitmap, error) {
	const _agggf = "\u0063\u006c\u006f\u0073\u0065\u0042\u0072\u0069\u0063\u006b"
	if _decdf == nil {
		return nil, _c.Error(_agggf, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _gadc < 1 || _egag < 1 {
		return nil, _c.Error(_agggf, "\u0068S\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0053\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _gadc == 1 && _egag == 1 {
		return _decdf.Copy(), nil
	}
	if _gadc == 1 || _egag == 1 {
		_eaec := SelCreateBrick(_egag, _gadc, _egag/2, _gadc/2, SelHit)
		var _bdfa error
		_cfdg, _bdfa = _gcgg(_cfdg, _decdf, _eaec)
		if _bdfa != nil {
			return nil, _c.Wrap(_bdfa, _agggf, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _cfdg, nil
	}
	_gbgc := SelCreateBrick(1, _gadc, 0, _gadc/2, SelHit)
	_dbfd := SelCreateBrick(_egag, 1, _egag/2, 0, SelHit)
	_fcad, _ebgfe := _bag(nil, _decdf, _gbgc)
	if _ebgfe != nil {
		return nil, _c.Wrap(_ebgfe, _agggf, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _cfdg, _ebgfe = _bag(_cfdg, _fcad, _dbfd); _ebgfe != nil {
		return nil, _c.Wrap(_ebgfe, _agggf, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _, _ebgfe = _gcae(_fcad, _cfdg, _gbgc); _ebgfe != nil {
		return nil, _c.Wrap(_ebgfe, _agggf, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	if _, _ebgfe = _gcae(_cfdg, _fcad, _dbfd); _ebgfe != nil {
		return nil, _c.Wrap(_ebgfe, _agggf, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _cfdg, nil
}

func HausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH int) (bool, error) {
	const _cegag = "\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_ddfb, _afddc := p1.Width, p1.Height
	_daba, _faef := p3.Width, p3.Height
	if _fa.Abs(_ddfb-_daba) > maxDiffW {
		return false, nil
	}
	if _fa.Abs(_afddc-_faef) > maxDiffH {
		return false, nil
	}
	_bbcd := int(delX + _fa.Sign(delX)*0.5)
	_accg := int(delY + _fa.Sign(delY)*0.5)
	var _fbea error
	_eggd := p1.CreateTemplate()
	if _fbea = _eggd.RasterOperation(0, 0, _ddfb, _afddc, PixSrc, p1, 0, 0); _fbea != nil {
		return false, _c.Wrap(_fbea, _cegag, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _fbea = _eggd.RasterOperation(_bbcd, _accg, _ddfb, _afddc, PixNotSrcAndDst, p4, 0, 0); _fbea != nil {
		return false, _c.Wrap(_fbea, _cegag, "\u0021p\u0034\u0020\u0026\u0020\u0074")
	}
	if _eggd.Zero() {
		return false, nil
	}
	if _fbea = _eggd.RasterOperation(_bbcd, _accg, _daba, _faef, PixSrc, p3, 0, 0); _fbea != nil {
		return false, _c.Wrap(_fbea, _cegag, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _fbea = _eggd.RasterOperation(0, 0, _daba, _faef, PixNotSrcAndDst, p2, 0, 0); _fbea != nil {
		return false, _c.Wrap(_fbea, _cegag, "\u0021p\u0032\u0020\u0026\u0020\u0074")
	}
	return _eggd.Zero(), nil
}

type Selection struct {
	Height, Width int
	Cx, Cy        int
	Name          string
	Data          [][]SelectionValue
}

func _geg(_efbb, _efdg, _dfd *Bitmap) (*Bitmap, error) {
	const _dcfe = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0078\u006f\u0072"
	if _efdg == nil {
		return nil, _c.Error(_dcfe, "'\u0062\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _dfd == nil {
		return nil, _c.Error(_dcfe, "'\u0062\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _efbb == _dfd {
		return nil, _c.Error(_dcfe, "'\u0064\u0027\u0020\u003d\u003d\u0020\u0027\u0062\u0032\u0027")
	}
	if !_efdg.SizesEqual(_dfd) {
		_fce.Log.Debug("\u0025s\u0020\u002d \u0042\u0069\u0074\u006da\u0070\u0020\u0027b\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074 e\u0071\u0075\u0061l\u0020\u0073i\u007a\u0065\u0020\u0077\u0069\u0074h\u0020\u0027b\u0032\u0027", _dcfe)
	}
	var _ccdd error
	if _efbb, _ccdd = _gbfag(_efbb, _efdg); _ccdd != nil {
		return nil, _c.Wrap(_ccdd, _dcfe, "\u0063\u0061n\u0027\u0074\u0020c\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027")
	}
	if _ccdd = _efbb.RasterOperation(0, 0, _efbb.Width, _efbb.Height, PixSrcXorDst, _dfd, 0, 0); _ccdd != nil {
		return nil, _c.Wrap(_ccdd, _dcfe, "")
	}
	return _efbb, nil
}
func (_dfee Points) Size() int { return len(_dfee) }
func _abgd(_fef, _gbac *Bitmap, _dgga, _bgbb, _baeb uint, _cafb, _aeeee int, _fbfda bool, _baee, _aebaf int) error {
	for _bgdac := _cafb; _bgdac < _aeeee; _bgdac++ {
		if _baee+1 < len(_fef.Data) {
			_bgggf := _bgdac+1 == _aeeee
			_gaeg, _aade := _fef.GetByte(_baee)
			if _aade != nil {
				return _aade
			}
			_baee++
			_gaeg <<= _dgga
			_aeeef, _aade := _fef.GetByte(_baee)
			if _aade != nil {
				return _aade
			}
			_aeeef >>= _bgbb
			_caeg := _gaeg | _aeeef
			if _bgggf && !_fbfda {
				_caeg = _aabd(_baeb, _caeg)
			}
			_aade = _gbac.SetByte(_aebaf, _caeg)
			if _aade != nil {
				return _aade
			}
			_aebaf++
			if _bgggf && _fbfda {
				_gbdg, _fcgbe := _fef.GetByte(_baee)
				if _fcgbe != nil {
					return _fcgbe
				}
				_gbdg <<= _dgga
				_caeg = _aabd(_baeb, _gbdg)
				if _fcgbe = _gbac.SetByte(_aebaf, _caeg); _fcgbe != nil {
					return _fcgbe
				}
			}
			continue
		}
		_bbcc, _efae := _fef.GetByte(_baee)
		if _efae != nil {
			_fce.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0061\u0074\u003a\u0020%\u0064\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0073", _baee, _efae)
			return _efae
		}
		_bbcc <<= _dgga
		_baee++
		_efae = _gbac.SetByte(_aebaf, _bbcc)
		if _efae != nil {
			return _efae
		}
		_aebaf++
	}
	return nil
}

func (_gff *Bitmap) Copy() *Bitmap {
	_egf := make([]byte, len(_gff.Data))
	copy(_egf, _gff.Data)
	return &Bitmap{Width: _gff.Width, Height: _gff.Height, RowStride: _gff.RowStride, Data: _egf, Color: _gff.Color, Text: _gff.Text, BitmapNumber: _gff.BitmapNumber, Special: _gff.Special}
}
func (_gefb *Points) AddPoint(x, y float32) { *_gefb = append(*_gefb, Point{x, y}) }
func (_gadg *Boxes) Get(i int) (*_g.Rectangle, error) {
	const _egde = "\u0042o\u0078\u0065\u0073\u002e\u0047\u0065t"
	if _gadg == nil {
		return nil, _c.Error(_egde, "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if i > len(*_gadg)-1 {
		return nil, _c.Errorf(_egde, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return (*_gadg)[i], nil
}

func _ec(_d, _gg *Bitmap) (_fcg error) {
	const _bd = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0032"
	_gf := _gg.RowStride
	_cf := _d.RowStride
	var (
		_efe                      byte
		_aa                       uint16
		_gc, _ggd, _dd, _ed, _ggb int
	)
	for _dd = 0; _dd < _gg.Height; _dd++ {
		_gc = _dd * _gf
		_ggd = 2 * _dd * _cf
		for _ed = 0; _ed < _gf; _ed++ {
			_efe = _gg.Data[_gc+_ed]
			_aa = _fceec[_efe]
			_ggb = _ggd + _ed*2
			if _d.RowStride != _gg.RowStride*2 && (_ed+1)*2 > _d.RowStride {
				_fcg = _d.SetByte(_ggb, byte(_aa>>8))
			} else {
				_fcg = _d.setTwoBytes(_ggb, _aa)
			}
			if _fcg != nil {
				return _c.Wrap(_fcg, _bd, "")
			}
		}
		for _ed = 0; _ed < _cf; _ed++ {
			_ggb = _ggd + _cf + _ed
			_efe = _d.Data[_ggd+_ed]
			if _fcg = _d.SetByte(_ggb, _efe); _fcg != nil {
				return _c.Wrapf(_fcg, _bd, "c\u006f\u0070\u0079\u0020\u0064\u006fu\u0062\u006c\u0065\u0064\u0020\u006ci\u006e\u0065\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0042\u0079\u0074\u0065\u003a\u0020\u0027\u0025d\u0027", _ggd+_ed, _ggd+_cf+_ed)
			}
		}
	}
	return nil
}

func Extract(roi _g.Rectangle, src *Bitmap) (*Bitmap, error) {
	_fedg := New(roi.Dx(), roi.Dy())
	_dgba := roi.Min.X & 0x07
	_fec := 8 - _dgba
	_bfeb := uint(8 - _fedg.Width&0x07)
	_adgb := src.GetByteIndex(roi.Min.X, roi.Min.Y)
	_gdde := src.GetByteIndex(roi.Max.X-1, roi.Min.Y)
	_cdcf := _fedg.RowStride == _gdde+1-_adgb
	var _bea int
	for _gede := roi.Min.Y; _gede < roi.Max.Y; _gede++ {
		_efbc := _adgb
		_ffcab := _bea
		switch {
		case _adgb == _gdde:
			_fab, _eaca := src.GetByte(_efbc)
			if _eaca != nil {
				return nil, _eaca
			}
			_fab <<= uint(_dgba)
			_eaca = _fedg.SetByte(_ffcab, _aabd(_bfeb, _fab))
			if _eaca != nil {
				return nil, _eaca
			}
		case _dgba == 0:
			for _gaga := _adgb; _gaga <= _gdde; _gaga++ {
				_faeb, _acae := src.GetByte(_efbc)
				if _acae != nil {
					return nil, _acae
				}
				_efbc++
				if _gaga == _gdde && _cdcf {
					_faeb = _aabd(_bfeb, _faeb)
				}
				_acae = _fedg.SetByte(_ffcab, _faeb)
				if _acae != nil {
					return nil, _acae
				}
				_ffcab++
			}
		default:
			_deb := _abgd(src, _fedg, uint(_dgba), uint(_fec), _bfeb, _adgb, _gdde, _cdcf, _efbc, _ffcab)
			if _deb != nil {
				return nil, _deb
			}
		}
		_adgb += src.RowStride
		_gdde += src.RowStride
		_bea += _fedg.RowStride
	}
	return _fedg, nil
}

type RasterOperator int

func CombineBytes(oldByte, newByte byte, op CombinationOperator) byte {
	return _fbfd(oldByte, newByte, op)
}

func (_dcgb *ClassedPoints) GetIntXByClass(i int) (int, error) {
	const _ebefc = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _dcgb.IntSlice.Size() {
		return 0, _c.Errorf(_ebefc, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_dcgb.XAtIndex(i)), nil
}

func (_ebda *Bitmap) setPadBits(_gbg int) {
	_cgcb := 8 - _ebda.Width%8
	if _cgcb == 8 {
		return
	}
	_cbdf := _ebda.Width / 8
	_gecc := _fgcfb[_cgcb]
	if _gbg == 0 {
		_gecc ^= _gecc
	}
	var _dac int
	for _dee := 0; _dee < _ebda.Height; _dee++ {
		_dac = _dee*_ebda.RowStride + _cbdf
		if _gbg == 0 {
			_ebda.Data[_dac] &= _gecc
		} else {
			_ebda.Data[_dac] |= _gecc
		}
	}
}

func (_aegc *Bitmap) equivalent(_fbg *Bitmap) bool {
	if _aegc == _fbg {
		return true
	}
	if !_aegc.SizesEqual(_fbg) {
		return false
	}
	_fada := _cbcgg(_aegc, _fbg, CmbOpXor)
	_dgg := _aegc.countPixels()
	_fdde := int(0.25 * float32(_dgg))
	if _fada.thresholdPixelSum(_fdde) {
		return false
	}
	var (
		_egfa [9][9]int
		_ceag [18][9]int
		_aadc [9][18]int
		_agd  int
		_caba int
	)
	_defb := 9
	_abe := _aegc.Height / _defb
	_bdca := _aegc.Width / _defb
	_eaf, _afc := _abe/2, _bdca/2
	if _abe < _bdca {
		_eaf = _bdca / 2
		_afc = _abe / 2
	}
	_ggfe := float64(_eaf) * float64(_afc) * _fc.Pi
	_eaeg := int(float64(_abe*_bdca/2) * 0.9)
	_cde := int(float64(_bdca*_abe/2) * 0.9)
	for _fge := 0; _fge < _defb; _fge++ {
		_dae := _bdca*_fge + _agd
		var _beg int
		if _fge == _defb-1 {
			_agd = 0
			_beg = _aegc.Width
		} else {
			_beg = _dae + _bdca
			if ((_aegc.Width - _agd) % _defb) > 0 {
				_agd++
				_beg++
			}
		}
		for _bccc := 0; _bccc < _defb; _bccc++ {
			_afgg := _abe*_bccc + _caba
			var _acga int
			if _bccc == _defb-1 {
				_caba = 0
				_acga = _aegc.Height
			} else {
				_acga = _afgg + _abe
				if (_aegc.Height-_caba)%_defb > 0 {
					_caba++
					_acga++
				}
			}
			var _affe, _fdae, _bgea, _bda int
			_cefcb := (_dae + _beg) / 2
			_dde := (_afgg + _acga) / 2
			for _fffcg := _dae; _fffcg < _beg; _fffcg++ {
				for _fdc := _afgg; _fdc < _acga; _fdc++ {
					if _fada.GetPixel(_fffcg, _fdc) {
						if _fffcg < _cefcb {
							_affe++
						} else {
							_fdae++
						}
						if _fdc < _dde {
							_bda++
						} else {
							_bgea++
						}
					}
				}
			}
			_egfa[_fge][_bccc] = _affe + _fdae
			_ceag[_fge*2][_bccc] = _affe
			_ceag[_fge*2+1][_bccc] = _fdae
			_aadc[_fge][_bccc*2] = _bda
			_aadc[_fge][_bccc*2+1] = _bgea
		}
	}
	for _bbbe := 0; _bbbe < _defb*2-1; _bbbe++ {
		for _cbfd := 0; _cbfd < (_defb - 1); _cbfd++ {
			var _adb int
			for _cbb := 0; _cbb < 2; _cbb++ {
				for _ega := 0; _ega < 2; _ega++ {
					_adb += _ceag[_bbbe+_cbb][_cbfd+_ega]
				}
			}
			if _adb > _cde {
				return false
			}
		}
	}
	for _afe := 0; _afe < (_defb - 1); _afe++ {
		for _ffc := 0; _ffc < ((_defb * 2) - 1); _ffc++ {
			var _adfa int
			for _ceba := 0; _ceba < 2; _ceba++ {
				for _decg := 0; _decg < 2; _decg++ {
					_adfa += _aadc[_afe+_ceba][_ffc+_decg]
				}
			}
			if _adfa > _eaeg {
				return false
			}
		}
	}
	for _fcb := 0; _fcb < (_defb - 2); _fcb++ {
		for _ebgf := 0; _ebgf < (_defb - 2); _ebgf++ {
			var _aba, _dgff int
			for _fcgd := 0; _fcgd < 3; _fcgd++ {
				for _fcag := 0; _fcag < 3; _fcag++ {
					if _fcgd == _fcag {
						_aba += _egfa[_fcb+_fcgd][_ebgf+_fcag]
					}
					if (2 - _fcgd) == _fcag {
						_dgff += _egfa[_fcb+_fcgd][_ebgf+_fcag]
					}
				}
			}
			if _aba > _cde || _dgff > _cde {
				return false
			}
		}
	}
	for _ddgf := 0; _ddgf < (_defb - 1); _ddgf++ {
		for _ecd := 0; _ecd < (_defb - 1); _ecd++ {
			var _babd int
			for _bgff := 0; _bgff < 2; _bgff++ {
				for _aagg := 0; _aagg < 2; _aagg++ {
					_babd += _egfa[_ddgf+_bgff][_ecd+_aagg]
				}
			}
			if float64(_babd) > _ggfe {
				return false
			}
		}
	}
	return true
}

func _bag(_bfceb *Bitmap, _ccafg *Bitmap, _fgbaa *Selection) (*Bitmap, error) {
	var (
		_cfgb *Bitmap
		_befe error
	)
	_bfceb, _befe = _aecg(_bfceb, _ccafg, _fgbaa, &_cfgb)
	if _befe != nil {
		return nil, _befe
	}
	if _befe = _bfceb.clearAll(); _befe != nil {
		return nil, _befe
	}
	var _bddbf SelectionValue
	for _fbc := 0; _fbc < _fgbaa.Height; _fbc++ {
		for _bdaf := 0; _bdaf < _fgbaa.Width; _bdaf++ {
			_bddbf = _fgbaa.Data[_fbc][_bdaf]
			if _bddbf == SelHit {
				if _befe = _bfceb.RasterOperation(_bdaf-_fgbaa.Cx, _fbc-_fgbaa.Cy, _ccafg.Width, _ccafg.Height, PixSrcOrDst, _cfgb, 0, 0); _befe != nil {
					return nil, _befe
				}
			}
		}
	}
	return _bfceb, nil
}

func _fcgc() (_aedf []byte) {
	_aedf = make([]byte, 256)
	for _caa := 0; _caa < 256; _caa++ {
		_adfd := byte(_caa)
		_aedf[_adfd] = (_adfd & 0x01) | ((_adfd & 0x04) >> 1) | ((_adfd & 0x10) >> 2) | ((_adfd & 0x40) >> 3) | ((_adfd & 0x02) << 3) | ((_adfd & 0x08) << 2) | ((_adfd & 0x20) << 1) | (_adfd & 0x80)
	}
	return _aedf
}
