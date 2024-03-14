package imageutil

import (
	_efc "encoding/binary"
	_d "errors"
	_c "fmt"
	_a "image"
	_g "image/color"
	_f "image/draw"
	_ef "math"

	_af "bitbucket.org/shenghui0779/gopdf/common"
	_ab "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
)

func (_fdeb *Gray8) Base() *ImageBase { return &_fdeb.ImageBase }
func _afge(_dfgg, _ceef RGBA, _ebac _a.Rectangle) {
	for _cade := 0; _cade < _ebac.Max.X; _cade++ {
		for _bbfcb := 0; _bbfcb < _ebac.Max.Y; _bbfcb++ {
			_ceef.SetRGBA(_cade, _bbfcb, _dfgg.RGBAAt(_cade, _bbfcb))
		}
	}
}

var _ _a.Image = &Gray8{}

func _acc() (_bcf [256]uint32) {
	for _fa := 0; _fa < 256; _fa++ {
		if _fa&0x01 != 0 {
			_bcf[_fa] |= 0xf
		}
		if _fa&0x02 != 0 {
			_bcf[_fa] |= 0xf0
		}
		if _fa&0x04 != 0 {
			_bcf[_fa] |= 0xf00
		}
		if _fa&0x08 != 0 {
			_bcf[_fa] |= 0xf000
		}
		if _fa&0x10 != 0 {
			_bcf[_fa] |= 0xf0000
		}
		if _fa&0x20 != 0 {
			_bcf[_fa] |= 0xf00000
		}
		if _fa&0x40 != 0 {
			_bcf[_fa] |= 0xf000000
		}
		if _fa&0x80 != 0 {
			_bcf[_fa] |= 0xf0000000
		}
	}
	return _bcf
}

func (_eda *Gray8) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _eda.Width, Y: _eda.Height}}
}

func (_eadf monochromeModel) Convert(c _g.Color) _g.Color {
	_ddaf := _g.GrayModel.Convert(c).(_g.Gray)
	return _ggd(_ddaf, _eadf)
}

type SMasker interface {
	HasAlpha() bool
	GetAlpha() []byte
	MakeAlpha()
}

func (_cbe *CMYK32) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtCMYK(x, y, _cbe.Width, _cbe.Data, _cbe.Decode)
}

var _ Image = &Gray2{}

func _ead(_adfg _g.CMYK) _g.NRGBA {
	_gdg, _adc, _ggfd := _g.CMYKToRGB(_adfg.C, _adfg.M, _adfg.Y, _adfg.K)
	return _g.NRGBA{R: _gdg, G: _adc, B: _ggfd, A: 0xff}
}
func (_gbc *Monochrome) setGrayBit(_fegg, _bggf int) { _gbc.Data[_fegg] |= 0x80 >> uint(_bggf&7) }
func (_bdga *RGBA32) Copy() Image                    { return &RGBA32{ImageBase: _bdga.copy()} }
func GetConverter(bitsPerComponent, colorComponents int) (ColorConverter, error) {
	switch colorComponents {
	case 1:
		switch bitsPerComponent {
		case 1:
			return MonochromeConverter, nil
		case 2:
			return Gray2Converter, nil
		case 4:
			return Gray4Converter, nil
		case 8:
			return GrayConverter, nil
		case 16:
			return Gray16Converter, nil
		}
	case 3:
		switch bitsPerComponent {
		case 4:
			return NRGBA16Converter, nil
		case 8:
			return NRGBAConverter, nil
		case 16:
			return NRGBA64Converter, nil
		}
	case 4:
		return CMYKConverter, nil
	}
	return nil, _c.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0043o\u006e\u0076\u0065\u0072\u0074\u0065\u0072\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072\u0073\u002e\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003a\u0020\u0025\u0064\u002c\u0020\u0043\u006f\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u003a \u0025\u0064", bitsPerComponent, colorComponents)
}

func (_bcea *Gray4) GrayAt(x, y int) _g.Gray {
	_bea, _ := ColorAtGray4BPC(x, y, _bcea.BytesPerLine, _bcea.Data, _bcea.Decode)
	return _bea
}
func (_bgca *NRGBA32) Base() *ImageBase { return &_bgca.ImageBase }
func (_bebbf *NRGBA64) setNRGBA64(_eca int, _dfbca _g.NRGBA64, _fege int) {
	_bebbf.Data[_eca] = uint8(_dfbca.R >> 8)
	_bebbf.Data[_eca+1] = uint8(_dfbca.R & 0xff)
	_bebbf.Data[_eca+2] = uint8(_dfbca.G >> 8)
	_bebbf.Data[_eca+3] = uint8(_dfbca.G & 0xff)
	_bebbf.Data[_eca+4] = uint8(_dfbca.B >> 8)
	_bebbf.Data[_eca+5] = uint8(_dfbca.B & 0xff)
	if _fege+1 < len(_bebbf.Alpha) {
		_bebbf.Alpha[_fege] = uint8(_dfbca.A >> 8)
		_bebbf.Alpha[_fege+1] = uint8(_dfbca.A & 0xff)
	}
}
func (_aebc *ImageBase) GetAlpha() []byte { return _aebc.Alpha }
func _gcba(_gddf *Monochrome, _aagfd, _efgf int, _debc, _fbcge int, _eece RasterOperator, _bfbg *Monochrome, _fcbed, _fbba int) error {
	var _edaa, _gfgb, _feag, _cfff int
	if _aagfd < 0 {
		_fcbed -= _aagfd
		_debc += _aagfd
		_aagfd = 0
	}
	if _fcbed < 0 {
		_aagfd -= _fcbed
		_debc += _fcbed
		_fcbed = 0
	}
	_edaa = _aagfd + _debc - _gddf.Width
	if _edaa > 0 {
		_debc -= _edaa
	}
	_gfgb = _fcbed + _debc - _bfbg.Width
	if _gfgb > 0 {
		_debc -= _gfgb
	}
	if _efgf < 0 {
		_fbba -= _efgf
		_fbcge += _efgf
		_efgf = 0
	}
	if _fbba < 0 {
		_efgf -= _fbba
		_fbcge += _fbba
		_fbba = 0
	}
	_feag = _efgf + _fbcge - _gddf.Height
	if _feag > 0 {
		_fbcge -= _feag
	}
	_cfff = _fbba + _fbcge - _bfbg.Height
	if _cfff > 0 {
		_fbcge -= _cfff
	}
	if _debc <= 0 || _fbcge <= 0 {
		return nil
	}
	var _egfccg error
	switch {
	case _aagfd&7 == 0 && _fcbed&7 == 0:
		_egfccg = _baac(_gddf, _aagfd, _efgf, _debc, _fbcge, _eece, _bfbg, _fcbed, _fbba)
	case _aagfd&7 == _fcbed&7:
		_egfccg = _baaba(_gddf, _aagfd, _efgf, _debc, _fbcge, _eece, _bfbg, _fcbed, _fbba)
	default:
		_egfccg = _dbbb(_gddf, _aagfd, _efgf, _debc, _fbcge, _eece, _bfbg, _fcbed, _fbba)
	}
	if _egfccg != nil {
		return _egfccg
	}
	return nil
}

func _ggfe() {
	for _bgcc := 0; _bgcc < 256; _bgcc++ {
		_gfbe[_bgcc] = uint8(_bgcc&0x1) + (uint8(_bgcc>>1) & 0x1) + (uint8(_bgcc>>2) & 0x1) + (uint8(_bgcc>>3) & 0x1) + (uint8(_bgcc>>4) & 0x1) + (uint8(_bgcc>>5) & 0x1) + (uint8(_bgcc>>6) & 0x1) + (uint8(_bgcc>>7) & 0x1)
	}
}

func (_ebfc *ImageBase) HasAlpha() bool {
	if _ebfc.Alpha == nil {
		return false
	}
	for _eafd := range _ebfc.Alpha {
		if _ebfc.Alpha[_eafd] != 0xff {
			return true
		}
	}
	return false
}

func (_aeea *Gray16) At(x, y int) _g.Color {
	_cead, _ := _aeea.ColorAt(x, y)
	return _cead
}

func _fcbe(_bbbe Gray, _cde nrgba64, _baga _a.Rectangle) {
	for _dcdf := 0; _dcdf < _baga.Max.X; _dcdf++ {
		for _adab := 0; _adab < _baga.Max.Y; _adab++ {
			_bbdc := _gad(_cde.NRGBA64At(_dcdf, _adab))
			_bbbe.SetGray(_dcdf, _adab, _bbdc)
		}
	}
}
func _eaedg(_adfc _g.Gray) _g.NRGBA { return _g.NRGBA{R: _adfc.Y, G: _adfc.Y, B: _adfc.Y, A: 0xff} }
func (_bbbdc *Gray16) Set(x, y int, c _g.Color) {
	_eedb := (y*_bbbdc.BytesPerLine/2 + x) * 2
	if _eedb+1 >= len(_bbbdc.Data) {
		return
	}
	_gcfc := _g.Gray16Model.Convert(c).(_g.Gray16)
	_bbbdc.Data[_eedb], _bbbdc.Data[_eedb+1] = uint8(_gcfc.Y>>8), uint8(_gcfc.Y&0xff)
}

func ColorAtGray16BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray16, error) {
	_gffe := (y*bytesPerLine/2 + x) * 2
	if _gffe+1 >= len(data) {
		return _g.Gray16{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_ccgc := uint16(data[_gffe])<<8 | uint16(data[_gffe+1])
	if len(decode) == 2 {
		_ccgc = uint16(uint64(LinearInterpolate(float64(_ccgc), 0, 65535, decode[0], decode[1])))
	}
	return _g.Gray16{Y: _ccgc}, nil
}

func (_bbfg *NRGBA16) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtNRGBA16(x, y, _bbfg.Width, _bbfg.BytesPerLine, _bbfg.Data, _bbfg.Alpha, _bbfg.Decode)
}

type ColorConverter interface {
	Convert(_cff _a.Image) (Image, error)
}
type Monochrome struct {
	ImageBase
	ModelThreshold uint8
}

func _aef() (_efcb []byte) {
	_efcb = make([]byte, 256)
	for _cee := 0; _cee < 256; _cee++ {
		_cea := byte(_cee)
		_efcb[_cea] = (_cea & 0x01) | ((_cea & 0x04) >> 1) | ((_cea & 0x10) >> 2) | ((_cea & 0x40) >> 3) | ((_cea & 0x02) << 3) | ((_cea & 0x08) << 2) | ((_cea & 0x20) << 1) | (_cea & 0x80)
	}
	return _efcb
}

func _gcd() (_ffb [256]uint16) {
	for _ac := 0; _ac < 256; _ac++ {
		if _ac&0x01 != 0 {
			_ffb[_ac] |= 0x3
		}
		if _ac&0x02 != 0 {
			_ffb[_ac] |= 0xc
		}
		if _ac&0x04 != 0 {
			_ffb[_ac] |= 0x30
		}
		if _ac&0x08 != 0 {
			_ffb[_ac] |= 0xc0
		}
		if _ac&0x10 != 0 {
			_ffb[_ac] |= 0x300
		}
		if _ac&0x20 != 0 {
			_ffb[_ac] |= 0xc00
		}
		if _ac&0x40 != 0 {
			_ffb[_ac] |= 0x3000
		}
		if _ac&0x80 != 0 {
			_ffb[_ac] |= 0xc000
		}
	}
	return _ffb
}

func ColorAtGray8BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_gfac := y*bytesPerLine + x
	if _gfac >= len(data) {
		return _g.Gray{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_egbgf := data[_gfac]
	if len(decode) == 2 {
		_egbgf = uint8(uint32(LinearInterpolate(float64(_egbgf), 0, 255, decode[0], decode[1])) & 0xff)
	}
	return _g.Gray{Y: _egbgf}, nil
}

func (_ddd *Monochrome) Scale(scale float64) (*Monochrome, error) {
	var _ffec bool
	_gffc := scale
	if scale < 1 {
		_gffc = 1 / scale
		_ffec = true
	}
	_gdf := NextPowerOf2(uint(_gffc))
	if InDelta(float64(_gdf), _gffc, 0.001) {
		if _ffec {
			return _ddd.ReduceBinary(_gffc)
		}
		return _ddd.ExpandBinary(int(_gdf))
	}
	_dfff := int(_ef.RoundToEven(float64(_ddd.Width) * scale))
	_fecf := int(_ef.RoundToEven(float64(_ddd.Height) * scale))
	return _ddd.ScaleLow(_dfff, _fecf)
}

const (
	_aegbf shift = iota
	_daae
)

func (_ebcd *Monochrome) getBit(_fbdc, _ebef int) uint8 {
	return _ebcd.Data[_fbdc+(_ebef>>3)] >> uint(7-(_ebef&7)) & 1
}

func _bcd(_ecdc _g.NRGBA64) _g.RGBA {
	_efcc, _cgfe, _babc, _aeg := _ecdc.RGBA()
	return _g.RGBA{R: uint8(_efcc >> 8), G: uint8(_cgfe >> 8), B: uint8(_babc >> 8), A: uint8(_aeg >> 8)}
}

func _befe(_abdf _g.CMYK) _g.Gray {
	_aeec, _eefb, _fbd := _g.CMYKToRGB(_abdf.C, _abdf.M, _abdf.Y, _abdf.K)
	_cfee := (19595*uint32(_aeec) + 38470*uint32(_eefb) + 7471*uint32(_fbd) + 1<<7) >> 16
	return _g.Gray{Y: uint8(_cfee)}
}

var _ Image = &Monochrome{}

func (_gdeg *RGBA32) SetRGBA(x, y int, c _g.RGBA) {
	_gdbf := y*_gdeg.Width + x
	_gaege := 3 * _gdbf
	if _gaege+2 >= len(_gdeg.Data) {
		return
	}
	_gdeg.setRGBA(_gdbf, c)
}
func (_afbe *Gray4) ColorModel() _g.Model { return Gray4Model }

var (
	MonochromeConverter = ConverterFunc(_egfc)
	Gray2Converter      = ConverterFunc(_deb)
	Gray4Converter      = ConverterFunc(_eegf)
	GrayConverter       = ConverterFunc(_dgca)
	Gray16Converter     = ConverterFunc(_cgb)
	NRGBA16Converter    = ConverterFunc(_afdbf)
	NRGBAConverter      = ConverterFunc(_gbea)
	NRGBA64Converter    = ConverterFunc(_fadd)
	RGBAConverter       = ConverterFunc(_eeggg)
	CMYKConverter       = ConverterFunc(_ceg)
)

func IsGrayImgBlackAndWhite(i *_a.Gray) bool { return _abbe(i) }
func (_bfeb *Monochrome) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _agbc(_bfeb, dx, dy, dw, dh, op, src, sx, sy)
}

var _ Gray = &Monochrome{}

func _fdd(_acg *Monochrome, _ebe ...int) (_aec *Monochrome, _egf error) {
	if _acg == nil {
		return nil, _d.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_ebe) == 0 {
		return nil, _d.New("\u0074h\u0065\u0072e\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020o\u006e\u0065\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006f\u0066 \u0072\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e")
	}
	_dfa := _aef()
	_aec = _acg
	for _, _bga := range _ebe {
		if _bga <= 0 {
			break
		}
		_aec, _egf = _dee(_aec, _bga, _dfa)
		if _egf != nil {
			return nil, _egf
		}
	}
	return _aec, nil
}

type monochromeModel uint8

func (_cag *Gray4) Base() *ImageBase { return &_cag.ImageBase }
func _dgdd(_egc _g.RGBA) _g.Gray {
	_aedd := (19595*uint32(_egc.R) + 38470*uint32(_egc.G) + 7471*uint32(_egc.B) + 1<<7) >> 16
	return _g.Gray{Y: uint8(_aedd)}
}

func NextPowerOf2(n uint) uint {
	if IsPowerOf2(n) {
		return n
	}
	return 1 << (_gfga(n) + 1)
}

func (_bgaa *Gray4) Validate() error {
	if len(_bgaa.Data) != _bgaa.Height*_bgaa.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

func _eged(_geed _g.NYCbCrA) _g.RGBA {
	_gfa, _bedb, _gfd, _bge := _dcgc(_geed).RGBA()
	return _g.RGBA{R: uint8(_gfa >> 8), G: uint8(_bedb >> 8), B: uint8(_gfd >> 8), A: uint8(_bge >> 8)}
}

func (_cgga *NRGBA64) NRGBA64At(x, y int) _g.NRGBA64 {
	_fef, _ := ColorAtNRGBA64(x, y, _cgga.Width, _cgga.Data, _cgga.Alpha, _cgga.Decode)
	return _fef
}

func (_acca *Gray2) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray2BPC(x, y, _acca.BytesPerLine, _acca.Data, _acca.Decode)
}

func (_cfd *Gray4) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray4BPC(x, y, _cfd.BytesPerLine, _cfd.Data, _cfd.Decode)
}

func _facd(_gafdg _g.NRGBA) _g.NRGBA {
	_gafdg.R = _gafdg.R>>4 | (_gafdg.R>>4)<<4
	_gafdg.G = _gafdg.G>>4 | (_gafdg.G>>4)<<4
	_gafdg.B = _gafdg.B>>4 | (_gafdg.B>>4)<<4
	return _gafdg
}

type shift int

func _abadbf(_bgfcd nrgba64, _fgeb RGBA, _acfd _a.Rectangle) {
	for _fcadg := 0; _fcadg < _acfd.Max.X; _fcadg++ {
		for _ddbbb := 0; _ddbbb < _acfd.Max.Y; _ddbbb++ {
			_cdeff := _bgfcd.NRGBA64At(_fcadg, _ddbbb)
			_fgeb.SetRGBA(_fcadg, _ddbbb, _bcd(_cdeff))
		}
	}
}

var _ Image = &NRGBA32{}

func _egba(_dfe, _eag *Monochrome, _dcg []byte, _ccb int) (_gcb error) {
	var (
		_cad, _begc, _cbd, _dda, _bd, _dgd, _eea, _ffa int
		_ded, _efce                                    uint32
		_bba, _aab                                     byte
		_cd                                            uint16
	)
	_fdf := make([]byte, 4)
	_ffc := make([]byte, 4)
	for _cbd = 0; _cbd < _dfe.Height-1; _cbd, _dda = _cbd+2, _dda+1 {
		_cad = _cbd * _dfe.BytesPerLine
		_begc = _dda * _eag.BytesPerLine
		for _bd, _dgd = 0, 0; _bd < _ccb; _bd, _dgd = _bd+4, _dgd+1 {
			for _eea = 0; _eea < 4; _eea++ {
				_ffa = _cad + _bd + _eea
				if _ffa <= len(_dfe.Data)-1 && _ffa < _cad+_dfe.BytesPerLine {
					_fdf[_eea] = _dfe.Data[_ffa]
				} else {
					_fdf[_eea] = 0x00
				}
				_ffa = _cad + _dfe.BytesPerLine + _bd + _eea
				if _ffa <= len(_dfe.Data)-1 && _ffa < _cad+(2*_dfe.BytesPerLine) {
					_ffc[_eea] = _dfe.Data[_ffa]
				} else {
					_ffc[_eea] = 0x00
				}
			}
			_ded = _efc.BigEndian.Uint32(_fdf)
			_efce = _efc.BigEndian.Uint32(_ffc)
			_efce |= _ded
			_efce |= _efce << 1
			_efce &= 0xaaaaaaaa
			_ded = _efce | (_efce << 7)
			_bba = byte(_ded >> 24)
			_aab = byte((_ded >> 8) & 0xff)
			_ffa = _begc + _dgd
			if _ffa+1 == len(_eag.Data)-1 || _ffa+1 >= _begc+_eag.BytesPerLine {
				_eag.Data[_ffa] = _dcg[_bba]
			} else {
				_cd = (uint16(_dcg[_bba]) << 8) | uint16(_dcg[_aab])
				if _gcb = _eag.setTwoBytes(_ffa, _cd); _gcb != nil {
					return _c.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _ffa)
				}
				_dgd++
			}
		}
	}
	return nil
}

func (_gfbad *RGBA32) Validate() error {
	if len(_gfbad.Data) != 3*_gfbad.Width*_gfbad.Height {
		return _d.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

var _ _a.Image = &Gray16{}

func _egfc(_efb _a.Image) (Image, error) {
	if _gfgc, _gcea := _efb.(*Monochrome); _gcea {
		return _gfgc, nil
	}
	_gff := _efb.Bounds()
	var _cafb Gray
	switch _afb := _efb.(type) {
	case Gray:
		_cafb = _afb
	case NRGBA:
		_cafb = &Gray8{ImageBase: NewImageBase(_gff.Max.X, _gff.Max.Y, 8, 1, nil, nil, nil)}
		_bcdg(_cafb, _afb, _gff)
	case nrgba64:
		_cafb = &Gray8{ImageBase: NewImageBase(_gff.Max.X, _gff.Max.Y, 8, 1, nil, nil, nil)}
		_fcbe(_cafb, _afb, _gff)
	default:
		_bbadf, _babb := GrayConverter.Convert(_efb)
		if _babb != nil {
			return nil, _babb
		}
		_cafb = _bbadf.(Gray)
	}
	_acab, _ebc := NewImage(_gff.Max.X, _gff.Max.Y, 1, 1, nil, nil, nil)
	if _ebc != nil {
		return nil, _ebc
	}
	_gdbc := _acab.(*Monochrome)
	_daag := AutoThresholdTriangle(GrayHistogram(_cafb))
	for _ggfg := 0; _ggfg < _gff.Max.X; _ggfg++ {
		for _afdf := 0; _afdf < _gff.Max.Y; _afdf++ {
			_gcdc := _ggd(_cafb.GrayAt(_ggfg, _afdf), monochromeModel(_daag))
			_gdbc.SetGray(_ggfg, _afdf, _gcdc)
		}
	}
	return _acab, nil
}

func _ffed(_cggb, _cegg Gray, _cgab _a.Rectangle) {
	for _cfda := 0; _cfda < _cgab.Max.X; _cfda++ {
		for _fecc := 0; _fecc < _cgab.Max.Y; _fecc++ {
			_cegg.SetGray(_cfda, _fecc, _cggb.GrayAt(_cfda, _fecc))
		}
	}
}

func _edg(_fdc _g.CMYK) _g.RGBA {
	_ddgf, _bad, _dfaf := _g.CMYKToRGB(_fdc.C, _fdc.M, _fdc.Y, _fdc.K)
	return _g.RGBA{R: _ddgf, G: _bad, B: _dfaf, A: 0xff}
}
func (_ffbg *NRGBA16) Base() *ImageBase { return &_ffbg.ImageBase }
func _bega(_gadg _g.Color) _g.Color {
	_gfef := _g.NRGBAModel.Convert(_gadg).(_g.NRGBA)
	return _facd(_gfef)
}

var _ Image = &RGBA32{}

func _caa(_dcgf _g.NRGBA64) _g.NRGBA {
	return _g.NRGBA{R: uint8(_dcgf.R >> 8), G: uint8(_dcgf.G >> 8), B: uint8(_dcgf.B >> 8), A: uint8(_dcgf.A >> 8)}
}

func _eegf(_bcgg _a.Image) (Image, error) {
	if _dedc, _bbaa := _bcgg.(*Gray4); _bbaa {
		return _dedc.Copy(), nil
	}
	_bagg := _bcgg.Bounds()
	_afdc, _bfed := NewImage(_bagg.Max.X, _bagg.Max.Y, 4, 1, nil, nil, nil)
	if _bfed != nil {
		return nil, _bfed
	}
	_feda(_bcgg, _afdc, _bagg)
	return _afdc, nil
}
func _decfa(_abb, _aeeag, _ggfcg byte) byte { return (_abb &^ (_ggfcg)) | (_aeeag & _ggfcg) }
func init()                                 { _ggfe() }
func (_dffg *Gray16) Base() *ImageBase      { return &_dffg.ImageBase }
func (_dfcdf *NRGBA64) Set(x, y int, c _g.Color) {
	_abe := (y*_dfcdf.Width + x) * 2
	_aecb := _abe * 3
	if _aecb+5 >= len(_dfcdf.Data) {
		return
	}
	_fadf := _g.NRGBA64Model.Convert(c).(_g.NRGBA64)
	_dfcdf.setNRGBA64(_aecb, _fadf, _abe)
}

func _egd(_gaf *Monochrome, _cc, _ebg int) (*Monochrome, error) {
	if _gaf == nil {
		return nil, _d.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _cc <= 0 || _ebg <= 0 {
		return nil, _d.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _cc == _ebg {
		if _cc == 1 {
			return _gaf.copy(), nil
		}
		if _cc == 2 || _cc == 4 || _cc == 8 {
			_gge, _egb := _dg(_gaf, _cc)
			if _egb != nil {
				return nil, _egb
			}
			return _gge, nil
		}
	}
	_bcc := _cc * _gaf.Width
	_dgg := _ebg * _gaf.Height
	_ed := _bce(_bcc, _dgg)
	_dgcf := _ed.BytesPerLine
	var (
		_febe, _feg, _efda, _bfea, _gbd int
		_cec                            byte
		_aac                            error
	)
	for _feg = 0; _feg < _gaf.Height; _feg++ {
		_febe = _ebg * _feg * _dgcf
		for _efda = 0; _efda < _gaf.Width; _efda++ {
			if _fg := _gaf.getBitAt(_efda, _feg); _fg {
				_gbd = _cc * _efda
				for _bfea = 0; _bfea < _cc; _bfea++ {
					_ed.setIndexedBit(_febe*8 + _gbd + _bfea)
				}
			}
		}
		for _bfea = 1; _bfea < _ebg; _bfea++ {
			_fd := _febe + _bfea*_dgcf
			for _fgd := 0; _fgd < _dgcf; _fgd++ {
				if _cec, _aac = _ed.getByte(_febe + _fgd); _aac != nil {
					return nil, _aac
				}
				if _aac = _ed.setByte(_fd+_fgd, _cec); _aac != nil {
					return nil, _aac
				}
			}
		}
	}
	return _ed, nil
}

func (_cabbe *Gray16) Validate() error {
	if len(_cabbe.Data) != _cabbe.Height*_cabbe.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_beee *Monochrome) IsUnpadded() bool { return (_beee.Width * _beee.Height) == len(_beee.Data) }
func _agg(_bcg _g.RGBA) _g.NRGBA {
	switch _bcg.A {
	case 0xff:
		return _g.NRGBA{R: _bcg.R, G: _bcg.G, B: _bcg.B, A: 0xff}
	case 0x00:
		return _g.NRGBA{}
	default:
		_ffee, _ede, _feae, _gef := _bcg.RGBA()
		_ffee = (_ffee * 0xffff) / _gef
		_ede = (_ede * 0xffff) / _gef
		_feae = (_feae * 0xffff) / _gef
		return _g.NRGBA{R: uint8(_ffee >> 8), G: uint8(_ede >> 8), B: uint8(_feae >> 8), A: uint8(_gef >> 8)}
	}
}

func _deb(_dagd _a.Image) (Image, error) {
	if _ceag, _dbf := _dagd.(*Gray2); _dbf {
		return _ceag.Copy(), nil
	}
	_ccdf := _dagd.Bounds()
	_bebb, _bffg := NewImage(_ccdf.Max.X, _ccdf.Max.Y, 2, 1, nil, nil, nil)
	if _bffg != nil {
		return nil, _bffg
	}
	_feda(_dagd, _bebb, _ccdf)
	return _bebb, nil
}

func _dgca(_ecee _a.Image) (Image, error) {
	if _eec, _ddf := _ecee.(*Gray8); _ddf {
		return _eec.Copy(), nil
	}
	_cgfg := _ecee.Bounds()
	_egce, _bcdf := NewImage(_cgfg.Max.X, _cgfg.Max.Y, 8, 1, nil, nil, nil)
	if _bcdf != nil {
		return nil, _bcdf
	}
	_feda(_ecee, _egce, _cgfg)
	return _egce, nil
}

func (_cba *CMYK32) Set(x, y int, c _g.Color) {
	_ggff := 4 * (y*_cba.Width + x)
	if _ggff+3 >= len(_cba.Data) {
		return
	}
	_cfef := _g.CMYKModel.Convert(c).(_g.CMYK)
	_cba.Data[_ggff] = _cfef.C
	_cba.Data[_ggff+1] = _cfef.M
	_cba.Data[_ggff+2] = _cfef.Y
	_cba.Data[_ggff+3] = _cfef.K
}

func RasterOperation(dest *Monochrome, dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _agbc(dest, dx, dy, dw, dh, op, src, sx, sy)
}

func GrayHistogram(g Gray) (_cgdc [256]int) {
	switch _cbbb := g.(type) {
	case Histogramer:
		return _cbbb.Histogram()
	case _a.Image:
		_cdad := _cbbb.Bounds()
		for _geac := 0; _geac < _cdad.Max.X; _geac++ {
			for _dgddg := 0; _dgddg < _cdad.Max.Y; _dgddg++ {
				_cgdc[g.GrayAt(_geac, _dgddg).Y]++
			}
		}
		return _cgdc
	default:
		return [256]int{}
	}
}
func (_cgd *Monochrome) Base() *ImageBase { return &_cgd.ImageBase }
func BytesPerLine(width, bitsPerComponent, colorComponents int) int {
	return ((width*bitsPerComponent)*colorComponents + 7) >> 3
}

func (_bcbc *NRGBA16) Set(x, y int, c _g.Color) {
	_edb := y*_bcbc.BytesPerLine + x*3/2
	if _edb+1 >= len(_bcbc.Data) {
		return
	}
	_cbdba := NRGBA16Model.Convert(c).(_g.NRGBA)
	_bcbc.setNRGBA(x, y, _edb, _cbdba)
}

func _bgf(_bdfc NRGBA, _fegb Gray, _fgc _a.Rectangle) {
	for _dbbc := 0; _dbbc < _fgc.Max.X; _dbbc++ {
		for _ddfe := 0; _ddfe < _fgc.Max.Y; _ddfe++ {
			_eccb := _deec(_bdfc.NRGBAAt(_dbbc, _ddfe))
			_fegb.SetGray(_dbbc, _ddfe, _eccb)
		}
	}
}

func (_cga *Gray8) GrayAt(x, y int) _g.Gray {
	_cgag, _ := ColorAtGray8BPC(x, y, _cga.BytesPerLine, _cga.Data, _cga.Decode)
	return _cgag
}

func (_efgfa *NRGBA32) At(x, y int) _g.Color {
	_egfbb, _ := _efgfa.ColorAt(x, y)
	return _egfbb
}

func (_cfg *ImageBase) setByte(_aaac int, _cabgg byte) error {
	if _aaac > len(_cfg.Data)-1 {
		return _d.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_cfg.Data[_aaac] = _cabgg
	return nil
}

func _ggd(_eff _g.Gray, _ebf monochromeModel) _g.Gray {
	if _eff.Y > uint8(_ebf) {
		return _g.Gray{Y: _ef.MaxUint8}
	}
	return _g.Gray{}
}

func (_dbdf *Gray8) Histogram() (_gcbb [256]int) {
	for _aagb := 0; _aagb < len(_dbdf.Data); _aagb++ {
		_gcbb[_dbdf.Data[_aagb]]++
	}
	return _gcbb
}

var _gfbe [256]uint8

func (_dbae *NRGBA64) ColorModel() _g.Model { return _g.NRGBA64Model }
func ColorAtNRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_g.NRGBA, error) {
	_bbc := y*width + x
	_becd := 3 * _bbc
	if _becd+2 >= len(data) {
		return _g.NRGBA{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_bbbc := uint8(0xff)
	if alpha != nil && len(alpha) > _bbc {
		_bbbc = alpha[_bbc]
	}
	_ddfed, _fbfc, _babf := data[_becd], data[_becd+1], data[_becd+2]
	if len(decode) == 6 {
		_ddfed = uint8(uint32(LinearInterpolate(float64(_ddfed), 0, 255, decode[0], decode[1])) & 0xff)
		_fbfc = uint8(uint32(LinearInterpolate(float64(_fbfc), 0, 255, decode[2], decode[3])) & 0xff)
		_babf = uint8(uint32(LinearInterpolate(float64(_babf), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _g.NRGBA{R: _ddfed, G: _fbfc, B: _babf, A: _bbbc}, nil
}
func (_bfa *CMYK32) Copy() Image { return &CMYK32{ImageBase: _bfa.copy()} }
func FromGoImage(i _a.Image) (Image, error) {
	switch _bbfc := i.(type) {
	case Image:
		return _bbfc.Copy(), nil
	case Gray:
		return GrayConverter.Convert(i)
	case *_a.Gray16:
		return Gray16Converter.Convert(i)
	case CMYK:
		return CMYKConverter.Convert(i)
	case *_a.NRGBA64:
		return NRGBA64Converter.Convert(i)
	default:
		return NRGBAConverter.Convert(i)
	}
}

func _gb(_afd, _aee *Monochrome) (_df error) {
	_adf := _aee.BytesPerLine
	_gag := _afd.BytesPerLine
	var (
		_aff                      byte
		_bff                      uint16
		_gae, _be, _ba, _aa, _gga int
	)
	for _ba = 0; _ba < _aee.Height; _ba++ {
		_gae = _ba * _adf
		_be = 2 * _ba * _gag
		for _aa = 0; _aa < _adf; _aa++ {
			_aff = _aee.Data[_gae+_aa]
			_bff = _ea[_aff]
			_gga = _be + _aa*2
			if _afd.BytesPerLine != _aee.BytesPerLine*2 && (_aa+1)*2 > _afd.BytesPerLine {
				_df = _afd.setByte(_gga, byte(_bff>>8))
			} else {
				_df = _afd.setTwoBytes(_gga, _bff)
			}
			if _df != nil {
				return _df
			}
		}
		for _aa = 0; _aa < _gag; _aa++ {
			_gga = _be + _gag + _aa
			_aff = _afd.Data[_be+_aa]
			if _df = _afd.setByte(_gga, _aff); _df != nil {
				return _df
			}
		}
	}
	return nil
}

func (_dafe *Gray8) Set(x, y int, c _g.Color) {
	_dbab := y*_dafe.BytesPerLine + x
	if _dbab > len(_dafe.Data)-1 {
		return
	}
	_ffde := _g.GrayModel.Convert(c)
	_dafe.Data[_dbab] = _ffde.(_g.Gray).Y
}

func (_cgec *NRGBA64) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtNRGBA64(x, y, _cgec.Width, _cgec.Data, _cgec.Alpha, _cgec.Decode)
}

func InDelta(expected, current, delta float64) bool {
	_ccbcf := expected - current
	if _ccbcf <= -delta || _ccbcf >= delta {
		return false
	}
	return true
}

type Gray4 struct{ ImageBase }

func (_geec *Gray2) Set(x, y int, c _g.Color) {
	if x >= _geec.Width || y >= _geec.Height {
		return
	}
	_cbg := Gray2Model.Convert(c).(_g.Gray)
	_bfbd := y * _geec.BytesPerLine
	_agdf := _bfbd + (x >> 2)
	_dcfc := _cbg.Y >> 6
	_geec.Data[_agdf] = (_geec.Data[_agdf] & (^(0xc0 >> uint(2*((x)&3))))) | (_dcfc << uint(6-2*(x&3)))
}

func (_ggge *Monochrome) GrayAt(x, y int) _g.Gray {
	_gafb, _ := ColorAtGray1BPC(x, y, _ggge.BytesPerLine, _ggge.Data, _ggge.Decode)
	return _gafb
}

func ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor int, data, alpha []byte, decode []float64) (_g.Color, error) {
	switch bitsPerColor {
	case 4:
		return ColorAtNRGBA16(x, y, width, bytesPerLine, data, alpha, decode)
	case 8:
		return ColorAtNRGBA32(x, y, width, data, alpha, decode)
	case 16:
		return ColorAtNRGBA64(x, y, width, data, alpha, decode)
	default:
		return nil, _c.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0067\u0062\u0020b\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0061\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}

type NRGBA64 struct{ ImageBase }

func (_cfba *NRGBA64) SetNRGBA64(x, y int, c _g.NRGBA64) {
	_edgee := (y*_cfba.Width + x) * 2
	_fgagc := _edgee * 3
	if _fgagc+5 >= len(_cfba.Data) {
		return
	}
	_cfba.setNRGBA64(_fgagc, c, _edgee)
}
func (_ebfb *ImageBase) Pix() []byte { return _ebfb.Data }
func ColorAtGray1BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_cdbg := y*bytesPerLine + x>>3
	if _cdbg >= len(data) {
		return _g.Gray{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_ceca := data[_cdbg] >> uint(7-(x&7)) & 1
	if len(decode) == 2 {
		_ceca = uint8(LinearInterpolate(float64(_ceca), 0.0, 1.0, decode[0], decode[1])) & 1
	}
	return _g.Gray{Y: _ceca * 255}, nil
}

func _feda(_ffce _a.Image, _fafa Image, _gcfg _a.Rectangle) {
	switch _begd := _ffce.(type) {
	case Gray:
		_ffed(_begd, _fafa.(Gray), _gcfg)
	case NRGBA:
		_bgf(_begd, _fafa.(Gray), _gcfg)
	case CMYK:
		_deff(_begd, _fafa.(Gray), _gcfg)
	case RGBA:
		_efgd(_begd, _fafa.(Gray), _gcfg)
	default:
		_bgg(_ffce, _fafa, _gcfg)
	}
}

var _ _a.Image = &Gray2{}

func (_acbc *Gray16) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray16BPC(x, y, _acbc.BytesPerLine, _acbc.Data, _acbc.Decode)
}

func (_acdae *Gray4) Histogram() (_baag [256]int) {
	for _gdd := 0; _gdd < _acdae.Width; _gdd++ {
		for _bcbd := 0; _bcbd < _acdae.Height; _bcbd++ {
			_baag[_acdae.GrayAt(_gdd, _bcbd).Y]++
		}
	}
	return _baag
}

func (_aba *Gray8) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray8BPC(x, y, _aba.BytesPerLine, _aba.Data, _aba.Decode)
}

func _cb(_bef, _cf *Monochrome) (_adb error) {
	_ca := _cf.BytesPerLine
	_ce := _bef.BytesPerLine
	_fb := _cf.BytesPerLine*4 - _bef.BytesPerLine
	var (
		_bg, _gcf                               byte
		_ggag                                   uint32
		_afe, _afc, _geb, _ega, _afa, _efd, _gd int
	)
	for _geb = 0; _geb < _cf.Height; _geb++ {
		_afe = _geb * _ca
		_afc = 4 * _geb * _ce
		for _ega = 0; _ega < _ca; _ega++ {
			_bg = _cf.Data[_afe+_ega]
			_ggag = _cca[_bg]
			_efd = _afc + _ega*4
			if _fb != 0 && (_ega+1)*4 > _bef.BytesPerLine {
				for _afa = _fb; _afa > 0; _afa-- {
					_gcf = byte((_ggag >> uint(_afa*8)) & 0xff)
					_gd = _efd + (_fb - _afa)
					if _adb = _bef.setByte(_gd, _gcf); _adb != nil {
						return _adb
					}
				}
			} else if _adb = _bef.setFourBytes(_efd, _ggag); _adb != nil {
				return _adb
			}
			if _adb = _bef.setFourBytes(_afc+_ega*4, _cca[_cf.Data[_afe+_ega]]); _adb != nil {
				return _adb
			}
		}
		for _afa = 1; _afa < 4; _afa++ {
			for _ega = 0; _ega < _ce; _ega++ {
				if _adb = _bef.setByte(_afc+_afa*_ce+_ega, _bef.Data[_afc+_ega]); _adb != nil {
					return _adb
				}
			}
		}
	}
	return nil
}

var (
	_ Gray     = &Gray16{}
	_ _a.Image = &Monochrome{}
)

func _ge(_dd *Monochrome, _gg int, _bf []uint) (*Monochrome, error) {
	_bfe := _gg * _dd.Width
	_eg := _gg * _dd.Height
	_db := _bce(_bfe, _eg)
	for _ae, _de := range _bf {
		var _ga error
		switch _de {
		case 2:
			_ga = _gb(_db, _dd)
		case 4:
			_ga = _cb(_db, _dd)
		case 8:
			_ga = _fbf(_db, _dd)
		}
		if _ga != nil {
			return nil, _ga
		}
		if _ae != len(_bf)-1 {
			_dd = _db.copy()
		}
	}
	return _db, nil
}

type ImageBase struct {
	Width, Height                     int
	BitsPerComponent, ColorComponents int
	Data, Alpha                       []byte
	Decode                            []float64
	BytesPerLine                      int
}

var _ _a.Image = &Gray4{}

func (_cdba *Monochrome) At(x, y int) _g.Color {
	_gdgb, _ := _cdba.ColorAt(x, y)
	return _gdgb
}

func (_fgfg *NRGBA32) Validate() error {
	if len(_fgfg.Data) != 3*_fgfg.Width*_fgfg.Height {
		return _d.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

func _bed(_cfbf _g.NRGBA) _g.CMYK {
	_aefd, _fff, _gdaf, _ := _cfbf.RGBA()
	_dff, _dfdb, _gbad, _beb := _g.RGBToCMYK(uint8(_aefd>>8), uint8(_fff>>8), uint8(_gdaf>>8))
	return _g.CMYK{C: _dff, M: _dfdb, Y: _gbad, K: _beb}
}

func (_bfg *Gray2) SetGray(x, y int, gray _g.Gray) {
	_fggd := _cabb(gray)
	_ebge := y * _bfg.BytesPerLine
	_fdef := _ebge + (x >> 2)
	if _fdef >= len(_bfg.Data) {
		return
	}
	_eage := _fggd.Y >> 6
	_bfg.Data[_fdef] = (_bfg.Data[_fdef] & (^(0xc0 >> uint(2*((x)&3))))) | (_eage << uint(6-2*(x&3)))
}

func _fadd(_acdf _a.Image) (Image, error) {
	if _fagf, _fcaf := _acdf.(*NRGBA64); _fcaf {
		return _fagf.Copy(), nil
	}
	_gaee, _aad, _aefc := _gfae(_acdf, 2)
	_adff, _geda := NewImage(_gaee.Max.X, _gaee.Max.Y, 16, 3, nil, _aefc, nil)
	if _geda != nil {
		return nil, _geda
	}
	_bdaa(_acdf, _adff, _gaee)
	if len(_aefc) != 0 && !_aad {
		if _ebefg := _bbac(_aefc, _adff); _ebefg != nil {
			return nil, _ebefg
		}
	}
	return _adff, nil
}
func (_eggf *Gray2) ColorModel() _g.Model { return Gray2Model }
func ImgToGray(i _a.Image) *_a.Gray {
	if _efgfb, _cacd := i.(*_a.Gray); _cacd {
		return _efgfb
	}
	_edcd := i.Bounds()
	_ebdf := _a.NewGray(_edcd)
	for _eead := 0; _eead < _edcd.Max.X; _eead++ {
		for _eabd := 0; _eabd < _edcd.Max.Y; _eabd++ {
			_dece := i.At(_eead, _eabd)
			_ebdf.Set(_eead, _eabd, _dece)
		}
	}
	return _ebdf
}

func (_deffa *ImageBase) getByte(_dcef int) (byte, error) {
	if _dcef > len(_deffa.Data)-1 || _dcef < 0 {
		return 0, _c.Errorf("\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", _dcef)
	}
	return _deffa.Data[_dcef], nil
}

func (_fcd *Monochrome) Copy() Image {
	return &Monochrome{ImageBase: _fcd.ImageBase.copy(), ModelThreshold: _fcd.ModelThreshold}
}

func ScaleAlphaToMonochrome(data []byte, width, height int) ([]byte, error) {
	_fc := BytesPerLine(width, 8, 1)
	if len(data) < _fc*height {
		return nil, nil
	}
	_fe := &Gray8{NewImageBase(width, height, 8, 1, data, nil, nil)}
	_gc, _eb := MonochromeConverter.Convert(_fe)
	if _eb != nil {
		return nil, _eb
	}
	return _gc.Base().Data, nil
}

func (_adg *Gray16) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _adg.Width, Y: _adg.Height}}
}

func _efef(_bbdb *_a.NYCbCrA, _edgec RGBA, _cbdd _a.Rectangle) {
	for _fffb := 0; _fffb < _cbdd.Max.X; _fffb++ {
		for _dceg := 0; _dceg < _cbdd.Max.Y; _dceg++ {
			_adbd := _bbdb.NYCbCrAAt(_fffb, _dceg)
			_edgec.SetRGBA(_fffb, _dceg, _eged(_adbd))
		}
	}
}

type RGBA interface {
	RGBAAt(_gdcd, _cgde int) _g.RGBA
	SetRGBA(_geg, _aaeg int, _bfbc _g.RGBA)
}

func _bdge(_gbbf RGBA, _dcfd NRGBA, _ebefa _a.Rectangle) {
	for _dbca := 0; _dbca < _ebefa.Max.X; _dbca++ {
		for _abadb := 0; _abadb < _ebefa.Max.Y; _abadb++ {
			_agfg := _gbbf.RGBAAt(_dbca, _abadb)
			_dcfd.SetNRGBA(_dbca, _abadb, _agg(_agfg))
		}
	}
}
func (_edab *NRGBA16) ColorModel() _g.Model { return NRGBA16Model }

var _ _a.Image = &RGBA32{}

type NRGBA16 struct{ ImageBase }

func (_eagb *Gray8) Copy() Image           { return &Gray8{ImageBase: _eagb.copy()} }
func (_gbcc *RGBA32) ColorModel() _g.Model { return _g.NRGBAModel }
func _fgf(_gf int) []uint {
	var _ccf []uint
	_fbb := _gf
	_bee := _fbb / 8
	if _bee != 0 {
		for _fab := 0; _fab < _bee; _fab++ {
			_ccf = append(_ccf, 8)
		}
		_gfg := _fbb % 8
		_fbb = 0
		if _gfg != 0 {
			_fbb = _gfg
		}
	}
	_gbf := _fbb / 4
	if _gbf != 0 {
		for _dga := 0; _dga < _gbf; _dga++ {
			_ccf = append(_ccf, 4)
		}
		_ccd := _fbb % 4
		_fbb = 0
		if _ccd != 0 {
			_fbb = _ccd
		}
	}
	_ggb := _fbb / 2
	if _ggb != 0 {
		for _beg := 0; _beg < _ggb; _beg++ {
			_ccf = append(_ccf, 2)
		}
	}
	return _ccf
}

func (_dbee *RGBA32) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _dbee.Width, Y: _dbee.Height}}
}

func (_eega *Gray16) SetGray(x, y int, g _g.Gray) {
	_dadd := (y*_eega.BytesPerLine/2 + x) * 2
	if _dadd+1 >= len(_eega.Data) {
		return
	}
	_eega.Data[_dadd] = g.Y
	_eega.Data[_dadd+1] = g.Y
}

type Gray16 struct{ ImageBase }

func _dbbb(_eaeb *Monochrome, _cbgd, _cece, _agdde, _abgbe int, _dedb RasterOperator, _geef *Monochrome, _acada, _dccg int) error {
	var (
		_age         bool
		_cdeac       bool
		_bdda        byte
		_eaec        int
		_cage        int
		_bdca        int
		_fcf         int
		_gbbb        bool
		_gceb        int
		_edggg       int
		_aea         int
		_bbbde       bool
		_dagf        byte
		_ddbd        int
		_gbaa        int
		_gabb        int
		_eadc        byte
		_bdbe        int
		_dfba        int
		_cbb         uint
		_bafd        uint
		_efga        byte
		_cgbf        shift
		_bafdf       bool
		_dacd        bool
		_ddcf, _cdac int
	)
	if _acada&7 != 0 {
		_dfba = 8 - (_acada & 7)
	}
	if _cbgd&7 != 0 {
		_cage = 8 - (_cbgd & 7)
	}
	if _dfba == 0 && _cage == 0 {
		_efga = _ffagf[0]
	} else {
		if _cage > _dfba {
			_cbb = uint(_cage - _dfba)
		} else {
			_cbb = uint(8 - (_dfba - _cage))
		}
		_bafd = 8 - _cbb
		_efga = _ffagf[_cbb]
	}
	if (_cbgd & 7) != 0 {
		_age = true
		_eaec = 8 - (_cbgd & 7)
		_bdda = _ffagf[_eaec]
		_bdca = _eaeb.BytesPerLine*_cece + (_cbgd >> 3)
		_fcf = _geef.BytesPerLine*_dccg + (_acada >> 3)
		_bdbe = 8 - (_acada & 7)
		if _eaec > _bdbe {
			_cgbf = _aegbf
			if _agdde >= _dfba {
				_bafdf = true
			}
		} else {
			_cgbf = _daae
		}
	}
	if _agdde < _eaec {
		_cdeac = true
		_bdda &= _bafab[8-_eaec+_agdde]
	}
	if !_cdeac {
		_gceb = (_agdde - _eaec) >> 3
		if _gceb != 0 {
			_gbbb = true
			_edggg = _eaeb.BytesPerLine*_cece + ((_cbgd + _cage) >> 3)
			_aea = _geef.BytesPerLine*_dccg + ((_acada + _cage) >> 3)
		}
	}
	_ddbd = (_cbgd + _agdde) & 7
	if !(_cdeac || _ddbd == 0) {
		_bbbde = true
		_dagf = _bafab[_ddbd]
		_gbaa = _eaeb.BytesPerLine*_cece + ((_cbgd + _cage) >> 3) + _gceb
		_gabb = _geef.BytesPerLine*_dccg + ((_acada + _cage) >> 3) + _gceb
		if _ddbd > int(_bafd) {
			_dacd = true
		}
	}
	switch _dedb {
	case PixSrc:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], _eadc, _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] = _eadc
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], _eadc, _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixNotSrc:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], ^_eadc, _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] = ^_eadc
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], ^_eadc, _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixSrcOrDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], _eadc|_eaeb.Data[_bdca], _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] |= _eadc
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], _eadc|_eaeb.Data[_gbaa], _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixSrcAndDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], _eadc&_eaeb.Data[_bdca], _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] &= _eadc
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], _eadc&_eaeb.Data[_gbaa], _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixSrcXorDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], _eadc^_eaeb.Data[_bdca], _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] ^= _eadc
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], _eadc^_eaeb.Data[_gbaa], _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], ^_eadc|_eaeb.Data[_bdca], _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] |= ^_eadc
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], ^_eadc|_eaeb.Data[_gbaa], _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], ^_eadc&_eaeb.Data[_bdca], _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] &= ^_eadc
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], ^_eadc&_eaeb.Data[_gbaa], _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], _eadc|^_eaeb.Data[_bdca], _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] = _eadc | ^_eaeb.Data[_edggg+_cdac]
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], _eadc|^_eaeb.Data[_gbaa], _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], _eadc&^_eaeb.Data[_bdca], _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] = _eadc &^ _eaeb.Data[_edggg+_cdac]
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], _eadc&^_eaeb.Data[_gbaa], _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], ^(_eadc | _eaeb.Data[_bdca]), _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] = ^(_eadc | _eaeb.Data[_edggg+_cdac])
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], ^(_eadc | _eaeb.Data[_gbaa]), _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], ^(_eadc & _eaeb.Data[_bdca]), _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] = ^(_eadc & _eaeb.Data[_edggg+_cdac])
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], ^(_eadc & _eaeb.Data[_gbaa]), _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		if _age {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				if _cgbf == _aegbf {
					_eadc = _geef.Data[_fcf] << _cbb
					if _bafdf {
						_eadc = _decfa(_eadc, _geef.Data[_fcf+1]>>_bafd, _efga)
					}
				} else {
					_eadc = _geef.Data[_fcf] >> _bafd
				}
				_eaeb.Data[_bdca] = _decfa(_eaeb.Data[_bdca], ^(_eadc ^ _eaeb.Data[_bdca]), _bdda)
				_bdca += _eaeb.BytesPerLine
				_fcf += _geef.BytesPerLine
			}
		}
		if _gbbb {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				for _cdac = 0; _cdac < _gceb; _cdac++ {
					_eadc = _decfa(_geef.Data[_aea+_cdac]<<_cbb, _geef.Data[_aea+_cdac+1]>>_bafd, _efga)
					_eaeb.Data[_edggg+_cdac] = ^(_eadc ^ _eaeb.Data[_edggg+_cdac])
				}
				_edggg += _eaeb.BytesPerLine
				_aea += _geef.BytesPerLine
			}
		}
		if _bbbde {
			for _ddcf = 0; _ddcf < _abgbe; _ddcf++ {
				_eadc = _geef.Data[_gabb] << _cbb
				if _dacd {
					_eadc = _decfa(_eadc, _geef.Data[_gabb+1]>>_bafd, _efga)
				}
				_eaeb.Data[_gbaa] = _decfa(_eaeb.Data[_gbaa], ^(_eadc ^ _eaeb.Data[_gbaa]), _dagf)
				_gbaa += _eaeb.BytesPerLine
				_gabb += _geef.BytesPerLine
			}
		}
	default:
		_af.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _dedb)
		return _d.New("\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}

func (_dag *Monochrome) ReduceBinary(factor float64) (*Monochrome, error) {
	_gafg := _gfga(uint(factor))
	if !IsPowerOf2(uint(factor)) {
		_gafg++
	}
	_ffef := make([]int, _gafg)
	for _cbef := range _ffef {
		_ffef[_cbef] = 4
	}
	_ggee, _cdb := _fdd(_dag, _ffef...)
	if _cdb != nil {
		return nil, _cdb
	}
	return _ggee, nil
}

func AddDataPadding(width, height, bitsPerComponent, colorComponents int, data []byte) ([]byte, error) {
	_bgfa := BytesPerLine(width, bitsPerComponent, colorComponents)
	if _bgfa == width*colorComponents*bitsPerComponent/8 {
		return data, nil
	}
	_fecfb := width * colorComponents * bitsPerComponent
	_gfe := _bgfa * 8
	_afdb := 8 - (_gfe - _fecfb)
	_aafd := _ab.NewReader(data)
	_fag := _bgfa - 1
	_dfcb := make([]byte, _fag)
	_gbef := make([]byte, height*_bgfa)
	_bcgb := _ab.NewWriterMSB(_gbef)
	var _dbeb uint64
	var _dbgb error
	for _dage := 0; _dage < height; _dage++ {
		_, _dbgb = _aafd.Read(_dfcb)
		if _dbgb != nil {
			return nil, _dbgb
		}
		_, _dbgb = _bcgb.Write(_dfcb)
		if _dbgb != nil {
			return nil, _dbgb
		}
		_dbeb, _dbgb = _aafd.ReadBits(byte(_afdb))
		if _dbgb != nil {
			return nil, _dbgb
		}
		_, _dbgb = _bcgb.WriteBits(_dbeb, _afdb)
		if _dbgb != nil {
			return nil, _dbgb
		}
		_bcgb.FinishByte()
	}
	return _gbef, nil
}

func (_cggc *NRGBA32) setRGBA(_cgfed int, _bde _g.NRGBA) {
	_ffedd := 3 * _cgfed
	_cggc.Data[_ffedd] = _bde.R
	_cggc.Data[_ffedd+1] = _bde.G
	_cggc.Data[_ffedd+2] = _bde.B
	if _cgfed < len(_cggc.Alpha) {
		_cggc.Alpha[_cgfed] = _bde.A
	}
}
func (_agge *Monochrome) ColorModel() _g.Model { return MonochromeModel(_agge.ModelThreshold) }
func _baaba(_afgg *Monochrome, _agcb, _ggbb, _efea, _cdea int, _bcfd RasterOperator, _fcag *Monochrome, _bfda, _eegaf int) error {
	var (
		_feeb        bool
		_bafg        bool
		_egabc       int
		_caca        int
		_eecg        int
		_aaba        bool
		_cgcc        byte
		_ebeg        int
		_gbde        int
		_baacg       int
		_ecff, _agdd int
	)
	_bgeg := 8 - (_agcb & 7)
	_cbcf := _ffagf[_bgeg]
	_dggb := _afgg.BytesPerLine*_ggbb + (_agcb >> 3)
	_ccab := _fcag.BytesPerLine*_eegaf + (_bfda >> 3)
	if _efea < _bgeg {
		_feeb = true
		_cbcf &= _bafab[8-_bgeg+_efea]
	}
	if !_feeb {
		_egabc = (_efea - _bgeg) >> 3
		if _egabc > 0 {
			_bafg = true
			_caca = _dggb + 1
			_eecg = _ccab + 1
		}
	}
	_ebeg = (_agcb + _efea) & 7
	if !(_feeb || _ebeg == 0) {
		_aaba = true
		_cgcc = _bafab[_ebeg]
		_gbde = _dggb + 1 + _egabc
		_baacg = _ccab + 1 + _egabc
	}
	switch _bcfd {
	case PixSrc:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], _fcag.Data[_ccab], _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] = _fcag.Data[_eecg+_agdd]
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], _fcag.Data[_baacg], _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixNotSrc:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], ^_fcag.Data[_ccab], _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] = ^_fcag.Data[_eecg+_agdd]
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], ^_fcag.Data[_baacg], _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixSrcOrDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], _fcag.Data[_ccab]|_afgg.Data[_dggb], _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] |= _fcag.Data[_eecg+_agdd]
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], _fcag.Data[_baacg]|_afgg.Data[_gbde], _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixSrcAndDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], _fcag.Data[_ccab]&_afgg.Data[_dggb], _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] &= _fcag.Data[_eecg+_agdd]
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], _fcag.Data[_baacg]&_afgg.Data[_gbde], _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixSrcXorDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], _fcag.Data[_ccab]^_afgg.Data[_dggb], _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] ^= _fcag.Data[_eecg+_agdd]
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], _fcag.Data[_baacg]^_afgg.Data[_gbde], _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], ^(_fcag.Data[_ccab])|_afgg.Data[_dggb], _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] |= ^(_fcag.Data[_eecg+_agdd])
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], ^(_fcag.Data[_baacg])|_afgg.Data[_gbde], _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], ^(_fcag.Data[_ccab])&_afgg.Data[_dggb], _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] &= ^_fcag.Data[_eecg+_agdd]
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], ^(_fcag.Data[_baacg])&_afgg.Data[_gbde], _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], _fcag.Data[_ccab]|^(_afgg.Data[_dggb]), _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] = _fcag.Data[_eecg+_agdd] | ^(_afgg.Data[_caca+_agdd])
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], _fcag.Data[_baacg]|^(_afgg.Data[_gbde]), _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], _fcag.Data[_ccab]&^(_afgg.Data[_dggb]), _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] = _fcag.Data[_eecg+_agdd] &^ (_afgg.Data[_caca+_agdd])
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], _fcag.Data[_baacg]&^(_afgg.Data[_gbde]), _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], ^(_fcag.Data[_ccab] | _afgg.Data[_dggb]), _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] = ^(_fcag.Data[_eecg+_agdd] | _afgg.Data[_caca+_agdd])
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], ^(_fcag.Data[_baacg] | _afgg.Data[_gbde]), _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], ^(_fcag.Data[_ccab] & _afgg.Data[_dggb]), _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] = ^(_fcag.Data[_eecg+_agdd] & _afgg.Data[_caca+_agdd])
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], ^(_fcag.Data[_baacg] & _afgg.Data[_gbde]), _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		for _ecff = 0; _ecff < _cdea; _ecff++ {
			_afgg.Data[_dggb] = _decfa(_afgg.Data[_dggb], ^(_fcag.Data[_ccab] ^ _afgg.Data[_dggb]), _cbcf)
			_dggb += _afgg.BytesPerLine
			_ccab += _fcag.BytesPerLine
		}
		if _bafg {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				for _agdd = 0; _agdd < _egabc; _agdd++ {
					_afgg.Data[_caca+_agdd] = ^(_fcag.Data[_eecg+_agdd] ^ _afgg.Data[_caca+_agdd])
				}
				_caca += _afgg.BytesPerLine
				_eecg += _fcag.BytesPerLine
			}
		}
		if _aaba {
			for _ecff = 0; _ecff < _cdea; _ecff++ {
				_afgg.Data[_gbde] = _decfa(_afgg.Data[_gbde], ^(_fcag.Data[_baacg] ^ _afgg.Data[_gbde]), _cgcc)
				_gbde += _afgg.BytesPerLine
				_baacg += _fcag.BytesPerLine
			}
		}
	default:
		_af.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _bcfd)
		return _d.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}

var _ _a.Image = &NRGBA16{}

func (_agaa *NRGBA16) At(x, y int) _g.Color { _efcca, _ := _agaa.ColorAt(x, y); return _efcca }
func _bbfcbc(_ggeea uint8) bool {
	if _ggeea == 0 || _ggeea == 255 {
		return true
	}
	return false
}

func NewImage(width, height, bitsPerComponent, colorComponents int, data, alpha []byte, decode []float64) (Image, error) {
	_beeb := NewImageBase(width, height, bitsPerComponent, colorComponents, data, alpha, decode)
	var _adbf Image
	switch colorComponents {
	case 1:
		switch bitsPerComponent {
		case 1:
			_adbf = &Monochrome{ImageBase: _beeb, ModelThreshold: 0x0f}
		case 2:
			_adbf = &Gray2{ImageBase: _beeb}
		case 4:
			_adbf = &Gray4{ImageBase: _beeb}
		case 8:
			_adbf = &Gray8{ImageBase: _beeb}
		case 16:
			_adbf = &Gray16{ImageBase: _beeb}
		}
	case 3:
		switch bitsPerComponent {
		case 4:
			_adbf = &NRGBA16{ImageBase: _beeb}
		case 8:
			_adbf = &NRGBA32{ImageBase: _beeb}
		case 16:
			_adbf = &NRGBA64{ImageBase: _beeb}
		}
	case 4:
		_adbf = &CMYK32{ImageBase: _beeb}
	}
	if _adbf == nil {
		return nil, ErrInvalidImage
	}
	return _adbf, nil
}
func MonochromeModel(threshold uint8) _g.Model { return monochromeModel(threshold) }

var _ Gray = &Gray2{}

func (_fbc *Monochrome) Histogram() (_ccde [256]int) {
	for _, _dec := range _fbc.Data {
		_ccde[0xff] += int(_gfbe[_fbc.Data[_dec]])
	}
	return _ccde
}

func _eagg(_ddfg *Monochrome, _fbcc, _cfbg int, _bgbf, _ggbf int, _cgge RasterOperator) {
	var (
		_bfag        int
		_fgb         byte
		_cfgg, _fdbf int
		_cge         int
	)
	_fcab := _bgbf >> 3
	_eeee := _bgbf & 7
	if _eeee > 0 {
		_fgb = _bafab[_eeee]
	}
	_bfag = _ddfg.BytesPerLine*_cfbg + (_fbcc >> 3)
	switch _cgge {
	case PixClr:
		for _cfgg = 0; _cfgg < _ggbf; _cfgg++ {
			_cge = _bfag + _cfgg*_ddfg.BytesPerLine
			for _fdbf = 0; _fdbf < _fcab; _fdbf++ {
				_ddfg.Data[_cge] = 0x0
				_cge++
			}
			if _eeee > 0 {
				_ddfg.Data[_cge] = _decfa(_ddfg.Data[_cge], 0x0, _fgb)
			}
		}
	case PixSet:
		for _cfgg = 0; _cfgg < _ggbf; _cfgg++ {
			_cge = _bfag + _cfgg*_ddfg.BytesPerLine
			for _fdbf = 0; _fdbf < _fcab; _fdbf++ {
				_ddfg.Data[_cge] = 0xff
				_cge++
			}
			if _eeee > 0 {
				_ddfg.Data[_cge] = _decfa(_ddfg.Data[_cge], 0xff, _fgb)
			}
		}
	case PixNotDst:
		for _cfgg = 0; _cfgg < _ggbf; _cfgg++ {
			_cge = _bfag + _cfgg*_ddfg.BytesPerLine
			for _fdbf = 0; _fdbf < _fcab; _fdbf++ {
				_ddfg.Data[_cge] = ^_ddfg.Data[_cge]
				_cge++
			}
			if _eeee > 0 {
				_ddfg.Data[_cge] = _decfa(_ddfg.Data[_cge], ^_ddfg.Data[_cge], _fgb)
			}
		}
	}
}

func _bce(_cce, _gbe int) *Monochrome {
	return &Monochrome{ImageBase: NewImageBase(_cce, _gbe, 1, 1, nil, nil, nil), ModelThreshold: 0x0f}
}

func _bfec(_dedcg *Monochrome, _cdgd, _ddcfc, _eagf, _dbc int, _ddde RasterOperator) {
	if _cdgd < 0 {
		_eagf += _cdgd
		_cdgd = 0
	}
	_fgfe := _cdgd + _eagf - _dedcg.Width
	if _fgfe > 0 {
		_eagf -= _fgfe
	}
	if _ddcfc < 0 {
		_dbc += _ddcfc
		_ddcfc = 0
	}
	_feggf := _ddcfc + _dbc - _dedcg.Height
	if _feggf > 0 {
		_dbc -= _feggf
	}
	if _eagf <= 0 || _dbc <= 0 {
		return
	}
	if (_cdgd & 7) == 0 {
		_eagg(_dedcg, _cdgd, _ddcfc, _eagf, _dbc, _ddde)
	} else {
		_gcgb(_dedcg, _cdgd, _ddcfc, _eagf, _dbc, _ddde)
	}
}

func _bdb(_aecf, _bfee *Monochrome, _bfb []byte, _bcba int) (_befb error) {
	var (
		_dab, _bbf, _gcc, _dge, _cac, _bag, _aed, _caf int
		_gee, _ccaf, _edd, _daa                        uint32
		_ag, _fga                                      byte
		_egbc                                          uint16
	)
	_gfb := make([]byte, 4)
	_gce := make([]byte, 4)
	for _gcc = 0; _gcc < _aecf.Height-1; _gcc, _dge = _gcc+2, _dge+1 {
		_dab = _gcc * _aecf.BytesPerLine
		_bbf = _dge * _bfee.BytesPerLine
		for _cac, _bag = 0, 0; _cac < _bcba; _cac, _bag = _cac+4, _bag+1 {
			for _aed = 0; _aed < 4; _aed++ {
				_caf = _dab + _cac + _aed
				if _caf <= len(_aecf.Data)-1 && _caf < _dab+_aecf.BytesPerLine {
					_gfb[_aed] = _aecf.Data[_caf]
				} else {
					_gfb[_aed] = 0x00
				}
				_caf = _dab + _aecf.BytesPerLine + _cac + _aed
				if _caf <= len(_aecf.Data)-1 && _caf < _dab+(2*_aecf.BytesPerLine) {
					_gce[_aed] = _aecf.Data[_caf]
				} else {
					_gce[_aed] = 0x00
				}
			}
			_gee = _efc.BigEndian.Uint32(_gfb)
			_ccaf = _efc.BigEndian.Uint32(_gce)
			_edd = _gee & _ccaf
			_edd |= _edd << 1
			_daa = _gee | _ccaf
			_daa &= _daa << 1
			_ccaf = _edd & _daa
			_ccaf &= 0xaaaaaaaa
			_gee = _ccaf | (_ccaf << 7)
			_ag = byte(_gee >> 24)
			_fga = byte((_gee >> 8) & 0xff)
			_caf = _bbf + _bag
			if _caf+1 == len(_bfee.Data)-1 || _caf+1 >= _bbf+_bfee.BytesPerLine {
				if _befb = _bfee.setByte(_caf, _bfb[_ag]); _befb != nil {
					return _c.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _caf)
				}
			} else {
				_egbc = (uint16(_bfb[_ag]) << 8) | uint16(_bfb[_fga])
				if _befb = _bfee.setTwoBytes(_caf, _egbc); _befb != nil {
					return _c.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _caf)
				}
				_bag++
			}
		}
	}
	return nil
}

type CMYK32 struct{ ImageBase }

func (_gabf *Gray16) Histogram() (_dfdc [256]int) {
	for _afbeg := 0; _afbeg < _gabf.Width; _afbeg++ {
		for _gdeb := 0; _gdeb < _gabf.Height; _gdeb++ {
			_dfdc[_gabf.GrayAt(_afbeg, _gdeb).Y]++
		}
	}
	return _dfdc
}
func (_bgef *NRGBA64) Copy() Image { return &NRGBA64{ImageBase: _bgef.copy()} }
func (_ceaa *Monochrome) getBitAt(_ggeg, _fcbf int) bool {
	_acda := _fcbf*_ceaa.BytesPerLine + (_ggeg >> 3)
	_aga := _ggeg & 0x07
	_gdaff := uint(7 - _aga)
	if _acda > len(_ceaa.Data)-1 {
		return false
	}
	if (_ceaa.Data[_acda]>>_gdaff)&0x01 >= 1 {
		return true
	}
	return false
}

func (_fbg *CMYK32) SetCMYK(x, y int, c _g.CMYK) {
	_fgge := 4 * (y*_fbg.Width + x)
	if _fgge+3 >= len(_fbg.Data) {
		return
	}
	_fbg.Data[_fgge] = c.C
	_fbg.Data[_fgge+1] = c.M
	_fbg.Data[_fgge+2] = c.Y
	_fbg.Data[_fgge+3] = c.K
}
func (_dgeb *CMYK32) At(x, y int) _g.Color { _dcd, _ := _dgeb.ColorAt(x, y); return _dcd }
func _dbebe(_fdgbg CMYK, _gcdf RGBA, _dbgdd _a.Rectangle) {
	for _cddbc := 0; _cddbc < _dbgdd.Max.X; _cddbc++ {
		for _defea := 0; _defea < _dbgdd.Max.Y; _defea++ {
			_eded := _fdgbg.CMYKAt(_cddbc, _defea)
			_gcdf.SetRGBA(_cddbc, _defea, _edg(_eded))
		}
	}
}

var _ NRGBA = &NRGBA32{}

func ColorAtNRGBA64(x, y, width int, data, alpha []byte, decode []float64) (_g.NRGBA64, error) {
	_ecgb := (y*width + x) * 2
	_gaged := _ecgb * 3
	if _gaged+5 >= len(data) {
		return _g.NRGBA64{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	const _eeed = 0xffff
	_badg := uint16(_eeed)
	if alpha != nil && len(alpha) > _ecgb+1 {
		_badg = uint16(alpha[_ecgb])<<8 | uint16(alpha[_ecgb+1])
	}
	_eeaf := uint16(data[_gaged])<<8 | uint16(data[_gaged+1])
	_agdgf := uint16(data[_gaged+2])<<8 | uint16(data[_gaged+3])
	_bdgc := uint16(data[_gaged+4])<<8 | uint16(data[_gaged+5])
	if len(decode) == 6 {
		_eeaf = uint16(uint64(LinearInterpolate(float64(_eeaf), 0, 65535, decode[0], decode[1])) & _eeed)
		_agdgf = uint16(uint64(LinearInterpolate(float64(_agdgf), 0, 65535, decode[2], decode[3])) & _eeed)
		_bdgc = uint16(uint64(LinearInterpolate(float64(_bdgc), 0, 65535, decode[4], decode[5])) & _eeed)
	}
	return _g.NRGBA64{R: _eeaf, G: _agdgf, B: _bdgc, A: _badg}, nil
}

var (
	_bafab = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_ffagf = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func (_aaaff *NRGBA64) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _aaaff.Width, Y: _aaaff.Height}}
}

func _eacg(_dgcfb _g.Color) _g.Color {
	_def := _g.GrayModel.Convert(_dgcfb).(_g.Gray)
	return _eefa(_def)
}

func _egcg(_gbb _g.NRGBA) _g.RGBA {
	_faf, _gbada, _bbad, _dfc := _gbb.RGBA()
	return _g.RGBA{R: uint8(_faf >> 8), G: uint8(_gbada >> 8), B: uint8(_bbad >> 8), A: uint8(_dfc >> 8)}
}

func _dee(_fec *Monochrome, _gdb int, _eeg []byte) (_cef *Monochrome, _dcc error) {
	const _ebd = "\u0072\u0065d\u0075\u0063\u0065R\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079"
	if _fec == nil {
		return nil, _d.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _gdb < 1 || _gdb > 4 {
		return nil, _d.New("\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _fec.Height <= 1 {
		return nil, _d.New("\u0073\u006f\u0075rc\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020m\u0075s\u0074 \u0062e\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027")
	}
	_cef = _bce(_fec.Width/2, _fec.Height/2)
	if _eeg == nil {
		_eeg = _aef()
	}
	_afab := _gaff(_fec.BytesPerLine, 2*_cef.BytesPerLine)
	switch _gdb {
	case 1:
		_dcc = _egba(_fec, _cef, _eeg, _afab)
	case 2:
		_dcc = _ccg(_fec, _cef, _eeg, _afab)
	case 3:
		_dcc = _bdb(_fec, _cef, _eeg, _afab)
	case 4:
		_dcc = _gaeg(_fec, _cef, _eeg, _afab)
	}
	if _dcc != nil {
		return nil, _dcc
	}
	return _cef, nil
}

type RasterOperator int

func (_gdgg *Gray4) Copy() Image { return &Gray4{ImageBase: _gdgg.copy()} }
func ConverterFunc(converterFunc func(_dgf _a.Image) (Image, error)) ColorConverter {
	return colorConverter{_dabg: converterFunc}
}

type Gray2 struct{ ImageBase }

func (_egdc *ImageBase) setEightPartlyBytes(_adda, _dfdba int, _bebc uint64) (_bedc error) {
	var (
		_bddf byte
		_fgag int
	)
	for _aeeca := 1; _aeeca <= _dfdba; _aeeca++ {
		_fgag = 64 - _aeeca*8
		_bddf = byte(_bebc >> uint(_fgag) & 0xff)
		if _bedc = _egdc.setByte(_adda+_aeeca-1, _bddf); _bedc != nil {
			return _bedc
		}
	}
	_dbe := _egdc.BytesPerLine*8 - _egdc.Width
	if _dbe == 0 {
		return nil
	}
	_fgag -= 8
	_bddf = byte(_bebc>>uint(_fgag)&0xff) << uint(_dbe)
	if _bedc = _egdc.setByte(_adda+_dfdba, _bddf); _bedc != nil {
		return _bedc
	}
	return nil
}

func (_cecb *Monochrome) Set(x, y int, c _g.Color) {
	_fdgb := y*_cecb.BytesPerLine + x>>3
	if _fdgb > len(_cecb.Data)-1 {
		return
	}
	_gaeca := _cecb.ColorModel().Convert(c).(_g.Gray)
	_cecb.setGray(x, _gaeca, _fdgb)
}

type colorConverter struct {
	_dabg func(_bgc _a.Image) (Image, error)
}

func (_fed *Monochrome) copy() *Monochrome {
	_efdg := _bce(_fed.Width, _fed.Height)
	_efdg.ModelThreshold = _fed.ModelThreshold
	_efdg.Data = make([]byte, len(_fed.Data))
	copy(_efdg.Data, _fed.Data)
	if len(_fed.Decode) != 0 {
		_efdg.Decode = make([]float64, len(_fed.Decode))
		copy(_efdg.Decode, _fed.Decode)
	}
	if len(_fed.Alpha) != 0 {
		_efdg.Alpha = make([]byte, len(_fed.Alpha))
		copy(_efdg.Alpha, _fed.Alpha)
	}
	return _efdg
}

func AutoThresholdTriangle(histogram [256]int) uint8 {
	var _abbc, _gdafe, _agddd, _ebfa int
	for _cdebc := 0; _cdebc < len(histogram); _cdebc++ {
		if histogram[_cdebc] > 0 {
			_abbc = _cdebc
			break
		}
	}
	if _abbc > 0 {
		_abbc--
	}
	for _adfd := 255; _adfd > 0; _adfd-- {
		if histogram[_adfd] > 0 {
			_ebfa = _adfd
			break
		}
	}
	if _ebfa < 255 {
		_ebfa++
	}
	for _gcgae := 0; _gcgae < 256; _gcgae++ {
		if histogram[_gcgae] > _gdafe {
			_agddd = _gcgae
			_gdafe = histogram[_gcgae]
		}
	}
	var _bdfg bool
	if (_agddd - _abbc) < (_ebfa - _agddd) {
		_bdfg = true
		var _bgad int
		_ddda := 255
		for _bgad < _ddda {
			_afed := histogram[_bgad]
			histogram[_bgad] = histogram[_ddda]
			histogram[_ddda] = _afed
			_bgad++
			_ddda--
		}
		_abbc = 255 - _ebfa
		_agddd = 255 - _agddd
	}
	if _abbc == _agddd {
		return uint8(_abbc)
	}
	_cceg := float64(histogram[_agddd])
	_eccbb := float64(_abbc - _agddd)
	_ecag := _ef.Sqrt(_cceg*_cceg + _eccbb*_eccbb)
	_cceg /= _ecag
	_eccbb /= _ecag
	_ecag = _cceg*float64(_abbc) + _eccbb*float64(histogram[_abbc])
	_degg := _abbc
	var _cbeac float64
	for _affc := _abbc + 1; _affc <= _agddd; _affc++ {
		_cadca := _cceg*float64(_affc) + _eccbb*float64(histogram[_affc]) - _ecag
		if _cadca > _cbeac {
			_degg = _affc
			_cbeac = _cadca
		}
	}
	_degg--
	if _bdfg {
		var _bfeag int
		_eefg := 255
		for _bfeag < _eefg {
			_gcda := histogram[_bfeag]
			histogram[_bfeag] = histogram[_eefg]
			histogram[_eefg] = _gcda
			_bfeag++
			_eefg--
		}
		return uint8(255 - _degg)
	}
	return uint8(_degg)
}

func (_bcfdg *NRGBA16) SetNRGBA(x, y int, c _g.NRGBA) {
	_edge := y*_bcfdg.BytesPerLine + x*3/2
	if _edge+1 >= len(_bcfdg.Data) {
		return
	}
	c = _facd(c)
	_bcfdg.setNRGBA(x, y, _edge, c)
}

func _cgb(_baec _a.Image) (Image, error) {
	if _cddc, _ebde := _baec.(*Gray16); _ebde {
		return _cddc.Copy(), nil
	}
	_aae := _baec.Bounds()
	_dafg, _edc := NewImage(_aae.Max.X, _aae.Max.Y, 16, 1, nil, nil, nil)
	if _edc != nil {
		return nil, _edc
	}
	_feda(_baec, _dafg, _aae)
	return _dafg, nil
}

func (_acad *Monochrome) InverseData() error {
	return _acad.RasterOperation(0, 0, _acad.Width, _acad.Height, PixNotDst, nil, 0, 0)
}

func _dc() (_ade [256]uint64) {
	for _eae := 0; _eae < 256; _eae++ {
		if _eae&0x01 != 0 {
			_ade[_eae] |= 0xff
		}
		if _eae&0x02 != 0 {
			_ade[_eae] |= 0xff00
		}
		if _eae&0x04 != 0 {
			_ade[_eae] |= 0xff0000
		}
		if _eae&0x08 != 0 {
			_ade[_eae] |= 0xff000000
		}
		if _eae&0x10 != 0 {
			_ade[_eae] |= 0xff00000000
		}
		if _eae&0x20 != 0 {
			_ade[_eae] |= 0xff0000000000
		}
		if _eae&0x40 != 0 {
			_ade[_eae] |= 0xff000000000000
		}
		if _eae&0x80 != 0 {
			_ade[_eae] |= 0xff00000000000000
		}
	}
	return _ade
}

func (_fad *CMYK32) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _fad.Width, Y: _fad.Height}}
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

func _ceg(_bdd _a.Image) (Image, error) {
	if _febd, _fge := _bdd.(*CMYK32); _fge {
		return _febd.Copy(), nil
	}
	_bab := _bdd.Bounds()
	_bgb, _ddac := NewImage(_bab.Max.X, _bab.Max.Y, 8, 4, nil, nil, nil)
	if _ddac != nil {
		return nil, _ddac
	}
	switch _fdg := _bdd.(type) {
	case CMYK:
		_cdg(_fdg, _bgb.(CMYK), _bab)
	case Gray:
		_cfbc(_fdg, _bgb.(CMYK), _bab)
	case NRGBA:
		_cgf(_fdg, _bgb.(CMYK), _bab)
	case RGBA:
		_aca(_fdg, _bgb.(CMYK), _bab)
	default:
		_bgg(_bdd, _bgb, _bab)
	}
	return _bgb, nil
}

func _cabb(_dce _g.Gray) _g.Gray {
	_ddga := _dce.Y >> 6
	_ddga |= _ddga << 2
	_dce.Y = _ddga | _ddga<<4
	return _dce
}
func (_gcg *Monochrome) setBit(_bbbd, _gdc int) { _gcg.Data[_bbbd+(_gdc>>3)] |= 0x80 >> uint(_gdc&7) }
func _gbea(_dcfcd _a.Image) (Image, error) {
	if _cbag, _ebbe := _dcfcd.(*NRGBA32); _ebbe {
		return _cbag.Copy(), nil
	}
	_eagea, _agdg, _dbede := _gfae(_dcfcd, 1)
	_caag, _acfc := NewImage(_eagea.Max.X, _eagea.Max.Y, 8, 3, nil, _dbede, nil)
	if _acfc != nil {
		return nil, _acfc
	}
	_aefb(_dcfcd, _caag, _eagea)
	if len(_dbede) != 0 && !_agdg {
		if _fgcd := _bbac(_dbede, _caag); _fgcd != nil {
			return nil, _fgcd
		}
	}
	return _caag, nil
}
func (_aabd *Gray8) ColorModel() _g.Model { return _g.GrayModel }
func ImgToBinary(i _a.Image, threshold uint8) *_a.Gray {
	switch _cdc := i.(type) {
	case *_a.Gray:
		if _abbe(_cdc) {
			return _cdc
		}
		return _ffbcg(_cdc, threshold)
	case *_a.Gray16:
		return _dfef(_cdc, threshold)
	default:
		return _eaac(_cdc, threshold)
	}
}

type Gray interface {
	GrayAt(_aabe, _dbd int) _g.Gray
	SetGray(_dcgb, _fafe int, _edef _g.Gray)
}

func ColorAt(x, y, width, bitsPerColor, colorComponents, bytesPerLine int, data, alpha []byte, decode []float64) (_g.Color, error) {
	switch colorComponents {
	case 1:
		return ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine, data, decode)
	case 3:
		return ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor, data, alpha, decode)
	case 4:
		return ColorAtCMYK(x, y, width, data, decode)
	default:
		return nil, _c.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063o\u006c\u006f\u0072\u0020\u0063\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", colorComponents)
	}
}

func (_aag *Monochrome) SetGray(x, y int, g _g.Gray) {
	_dbbg := y*_aag.BytesPerLine + x>>3
	if _dbbg > len(_aag.Data)-1 {
		return
	}
	g = _ggd(g, monochromeModel(_aag.ModelThreshold))
	_aag.setGray(x, g, _dbbg)
}

func _cfbc(_dcbd Gray, _ece CMYK, _cfa _a.Rectangle) {
	for _cbfd := 0; _cbfd < _cfa.Max.X; _cbfd++ {
		for _dae := 0; _dae < _cfa.Max.Y; _dae++ {
			_ceeg := _dcbd.GrayAt(_cbfd, _dae)
			_ece.SetCMYK(_cbfd, _dae, _dgga(_ceeg))
		}
	}
}

type RGBA32 struct{ ImageBase }

func _aefb(_abea _a.Image, _ecdg Image, _fcae _a.Rectangle) {
	if _bfc, _dfffg := _abea.(SMasker); _dfffg && _bfc.HasAlpha() {
		_ecdg.(SMasker).MakeAlpha()
	}
	switch _eadce := _abea.(type) {
	case Gray:
		_dcag(_eadce, _ecdg.(NRGBA), _fcae)
	case NRGBA:
		_deedb(_eadce, _ecdg.(NRGBA), _fcae)
	case *_a.NYCbCrA:
		_ggfge(_eadce, _ecdg.(NRGBA), _fcae)
	case CMYK:
		_gcbaf(_eadce, _ecdg.(NRGBA), _fcae)
	case RGBA:
		_bdge(_eadce, _ecdg.(NRGBA), _fcae)
	case nrgba64:
		_eacgf(_eadce, _ecdg.(NRGBA), _fcae)
	default:
		_bgg(_abea, _ecdg, _fcae)
	}
}

func _bdaa(_geca _a.Image, _gdbg Image, _bcgf _a.Rectangle) {
	if _dcfa, _fcg := _geca.(SMasker); _fcg && _dcfa.HasAlpha() {
		_gdbg.(SMasker).MakeAlpha()
	}
	_bgg(_geca, _gdbg, _bcgf)
}

func _baac(_gfba *Monochrome, _ffgd, _edec, _aacd, _dfcd int, _eegg RasterOperator, _ebgf *Monochrome, _dbgd, _gfgd int) error {
	var (
		_gfdd         byte
		_bdced        int
		_ffbc         int
		_eggc, _feed  int
		_eefaf, _bdff int
	)
	_baab := _aacd >> 3
	_gfge := _aacd & 7
	if _gfge > 0 {
		_gfdd = _bafab[_gfge]
	}
	_bdced = _ebgf.BytesPerLine*_gfgd + (_dbgd >> 3)
	_ffbc = _gfba.BytesPerLine*_edec + (_ffgd >> 3)
	switch _eegg {
	case PixSrc:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] = _ebgf.Data[_eggc]
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], _ebgf.Data[_eggc], _gfdd)
			}
		}
	case PixNotSrc:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] = ^(_ebgf.Data[_eggc])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], ^_ebgf.Data[_eggc], _gfdd)
			}
		}
	case PixSrcOrDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] |= _ebgf.Data[_eggc]
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], _ebgf.Data[_eggc]|_gfba.Data[_feed], _gfdd)
			}
		}
	case PixSrcAndDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] &= _ebgf.Data[_eggc]
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], _ebgf.Data[_eggc]&_gfba.Data[_feed], _gfdd)
			}
		}
	case PixSrcXorDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] ^= _ebgf.Data[_eggc]
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], _ebgf.Data[_eggc]^_gfba.Data[_feed], _gfdd)
			}
		}
	case PixNotSrcOrDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] |= ^(_ebgf.Data[_eggc])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], ^(_ebgf.Data[_eggc])|_gfba.Data[_feed], _gfdd)
			}
		}
	case PixNotSrcAndDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] &= ^(_ebgf.Data[_eggc])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], ^(_ebgf.Data[_eggc])&_gfba.Data[_feed], _gfdd)
			}
		}
	case PixSrcOrNotDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] = _ebgf.Data[_eggc] | ^(_gfba.Data[_feed])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], _ebgf.Data[_eggc]|^(_gfba.Data[_feed]), _gfdd)
			}
		}
	case PixSrcAndNotDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] = _ebgf.Data[_eggc] &^ (_gfba.Data[_feed])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], _ebgf.Data[_eggc]&^(_gfba.Data[_feed]), _gfdd)
			}
		}
	case PixNotPixSrcOrDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] = ^(_ebgf.Data[_eggc] | _gfba.Data[_feed])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], ^(_ebgf.Data[_eggc] | _gfba.Data[_feed]), _gfdd)
			}
		}
	case PixNotPixSrcAndDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] = ^(_ebgf.Data[_eggc] & _gfba.Data[_feed])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], ^(_ebgf.Data[_eggc] & _gfba.Data[_feed]), _gfdd)
			}
		}
	case PixNotPixSrcXorDst:
		for _eefaf = 0; _eefaf < _dfcd; _eefaf++ {
			_eggc = _bdced + _eefaf*_ebgf.BytesPerLine
			_feed = _ffbc + _eefaf*_gfba.BytesPerLine
			for _bdff = 0; _bdff < _baab; _bdff++ {
				_gfba.Data[_feed] = ^(_ebgf.Data[_eggc] ^ _gfba.Data[_feed])
				_feed++
				_eggc++
			}
			if _gfge > 0 {
				_gfba.Data[_feed] = _decfa(_gfba.Data[_feed], ^(_ebgf.Data[_eggc] ^ _gfba.Data[_feed]), _gfdd)
			}
		}
	default:
		_af.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _eegg)
		return _d.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}

var _ Image = &Gray4{}

func (_gafd *CMYK32) CMYKAt(x, y int) _g.CMYK {
	_ddbb, _ := ColorAtCMYK(x, y, _gafd.Width, _gafd.Data, _gafd.Decode)
	return _ddbb
}

func ColorAtGray4BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_bfaf := y*bytesPerLine + x>>1
	if _bfaf >= len(data) {
		return _g.Gray{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_daee := data[_bfaf] >> uint(4-(x&1)*4) & 0xf
	if len(decode) == 2 {
		_daee = uint8(uint32(LinearInterpolate(float64(_daee), 0, 15, decode[0], decode[1])) & 0xf)
	}
	return _g.Gray{Y: _daee * 17 & 0xff}, nil
}

func (_edeg *NRGBA16) NRGBAAt(x, y int) _g.NRGBA {
	_babbc, _ := ColorAtNRGBA16(x, y, _edeg.Width, _edeg.BytesPerLine, _edeg.Data, _edeg.Alpha, _edeg.Decode)
	return _babbc
}

type Histogramer interface{ Histogram() [256]int }

func (_bgac *Gray2) At(x, y int) _g.Color { _eeb, _ := _bgac.ColorAt(x, y); return _eeb }

var ErrInvalidImage = _d.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")

func (_cdbac *Gray4) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _cdbac.Width, Y: _cdbac.Height}}
}
func (_gbfd *RGBA32) At(x, y int) _g.Color { _cffb, _ := _gbfd.ColorAt(x, y); return _cffb }
func (_accf *NRGBA64) Validate() error {
	if len(_accf.Data) != 3*2*_accf.Width*_accf.Height {
		return _d.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

func (_gaeb *NRGBA32) NRGBAAt(x, y int) _g.NRGBA {
	_dfea, _ := ColorAtNRGBA32(x, y, _gaeb.Width, _gaeb.Data, _gaeb.Alpha, _gaeb.Decode)
	return _dfea
}

func ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine int, data []byte, decode []float64) (_g.Color, error) {
	switch bitsPerColor {
	case 1:
		return ColorAtGray1BPC(x, y, bytesPerLine, data, decode)
	case 2:
		return ColorAtGray2BPC(x, y, bytesPerLine, data, decode)
	case 4:
		return ColorAtGray4BPC(x, y, bytesPerLine, data, decode)
	case 8:
		return ColorAtGray8BPC(x, y, bytesPerLine, data, decode)
	case 16:
		return ColorAtGray16BPC(x, y, bytesPerLine, data, decode)
	default:
		return nil, _c.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0067\u0072\u0061\u0079\u0020\u0073c\u0061\u006c\u0065\u0020\u0062\u0069\u0074s\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020a\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}

type Image interface {
	_f.Image
	Base() *ImageBase
	Copy() Image
	Pix() []byte
	ColorAt(_abfe, _eedbb int) (_g.Color, error)
	Validate() error
}

func (_bafa *Gray8) SetGray(x, y int, g _g.Gray) {
	_gfdae := y*_bafa.BytesPerLine + x
	if _gfdae > len(_bafa.Data)-1 {
		return
	}
	_bafa.Data[_gfdae] = g.Y
}

func _ffbcg(_cbac *_a.Gray, _ebbb uint8) *_a.Gray {
	_deggg := _cbac.Bounds()
	_fefb := _a.NewGray(_deggg)
	for _efdb := 0; _efdb < _deggg.Dx(); _efdb++ {
		for _afca := 0; _afca < _deggg.Dy(); _afca++ {
			_eecb := _cbac.GrayAt(_efdb, _afca)
			_fefb.SetGray(_efdb, _afca, _g.Gray{Y: _ggeb(_eecb.Y, _ebbb)})
		}
	}
	return _fefb
}

func (_ggg *CMYK32) Validate() error {
	if len(_ggg.Data) != 4*_ggg.Width*_ggg.Height {
		return _d.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

var (
	Gray2Model   = _g.ModelFunc(_baf)
	Gray4Model   = _g.ModelFunc(_eacg)
	NRGBA16Model = _g.ModelFunc(_bega)
)

func (_ged colorConverter) Convert(src _a.Image) (Image, error) { return _ged._dabg(src) }
func MonochromeThresholdConverter(threshold uint8) ColorConverter {
	return &monochromeThresholdConverter{Threshold: threshold}
}

func (_bbfb *Monochrome) ScaleLow(width, height int) (*Monochrome, error) {
	if width < 0 || height < 0 {
		return nil, _d.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064t\u0068\u0020\u0061\u006e\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_bec := _bce(width, height)
	_add := make([]int, height)
	_gab := make([]int, width)
	_eafc := float64(_bbfb.Width) / float64(width)
	_cdf := float64(_bbfb.Height) / float64(height)
	for _ddbe := 0; _ddbe < height; _ddbe++ {
		_add[_ddbe] = int(_ef.Min(_cdf*float64(_ddbe)+0.5, float64(_bbfb.Height-1)))
	}
	for _cabg := 0; _cabg < width; _cabg++ {
		_gab[_cabg] = int(_ef.Min(_eafc*float64(_cabg)+0.5, float64(_bbfb.Width-1)))
	}
	_agf := -1
	_deed := byte(0)
	for _cegc := 0; _cegc < height; _cegc++ {
		_cdef := _add[_cegc] * _bbfb.BytesPerLine
		_edfb := _cegc * _bec.BytesPerLine
		for _bada := 0; _bada < width; _bada++ {
			_ggaf := _gab[_bada]
			if _ggaf != _agf {
				_deed = _bbfb.getBit(_cdef, _ggaf)
				if _deed != 0 {
					_bec.setBit(_edfb, _bada)
				}
				_agf = _ggaf
			} else {
				if _deed != 0 {
					_bec.setBit(_edfb, _bada)
				}
			}
		}
	}
	return _bec, nil
}

func (_abc *Monochrome) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _abc.Width, Y: _abc.Height}}
}

func (_agcba *NRGBA32) SetNRGBA(x, y int, c _g.NRGBA) {
	_eggb := y*_agcba.Width + x
	_baecb := 3 * _eggb
	if _baecb+2 >= len(_agcba.Data) {
		return
	}
	_agcba.setRGBA(_eggb, c)
}

var _ Gray = &Gray8{}

func _dg(_dgc *Monochrome, _b int) (*Monochrome, error) {
	if _dgc == nil {
		return nil, _d.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _b == 1 {
		return _dgc.copy(), nil
	}
	if !IsPowerOf2(uint(_b)) {
		return nil, _c.Errorf("\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006ci\u0064 \u0065x\u0070a\u006e\u0064\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _b)
	}
	_ad := _fgf(_b)
	return _ge(_dgc, _b, _ad)
}
func (_fdff *NRGBA32) ColorModel() _g.Model { return _g.NRGBAModel }
func _dfd(_efe _g.NRGBA) _g.Gray {
	var _gba _g.NRGBA
	if _efe == _gba {
		return _g.Gray{Y: 0xff}
	}
	_ggea, _bdg, _fcb, _ := _efe.RGBA()
	_gaec := (19595*_ggea + 38470*_bdg + 7471*_fcb + 1<<15) >> 24
	return _g.Gray{Y: uint8(_gaec)}
}

func _deff(_abgb CMYK, _fac Gray, _gcfga _a.Rectangle) {
	for _acegd := 0; _acegd < _gcfga.Max.X; _acegd++ {
		for _babe := 0; _babe < _gcfga.Max.Y; _babe++ {
			_beea := _befe(_abgb.CMYKAt(_acegd, _babe))
			_fac.SetGray(_acegd, _babe, _beea)
		}
	}
}

func (_egfe *Monochrome) ResolveDecode() error {
	if len(_egfe.Decode) != 2 {
		return nil
	}
	if _egfe.Decode[0] == 1 && _egfe.Decode[1] == 0 {
		if _agc := _egfe.InverseData(); _agc != nil {
			return _agc
		}
		_egfe.Decode = nil
	}
	return nil
}

func _aca(_ege RGBA, _gda CMYK, _bgag _a.Rectangle) {
	for _bae := 0; _bae < _bgag.Max.X; _bae++ {
		for _eee := 0; _eee < _bgag.Max.Y; _eee++ {
			_aceb := _ege.RGBAAt(_bae, _eee)
			_gda.SetCMYK(_bae, _eee, _eab(_aceb))
		}
	}
}

func _fbf(_bb, _efcg *Monochrome) (_bc error) {
	_bcb := _efcg.BytesPerLine
	_fea := _bb.BytesPerLine
	var _ff, _ffd, _feb, _ee, _gbg int
	for _feb = 0; _feb < _efcg.Height; _feb++ {
		_ff = _feb * _bcb
		_ffd = 8 * _feb * _fea
		for _ee = 0; _ee < _bcb; _ee++ {
			if _bc = _bb.setEightBytes(_ffd+_ee*8, _cfb[_efcg.Data[_ff+_ee]]); _bc != nil {
				return _bc
			}
		}
		for _gbg = 1; _gbg < 8; _gbg++ {
			for _ee = 0; _ee < _fea; _ee++ {
				if _bc = _bb.setByte(_ffd+_gbg*_fea+_ee, _bb.Data[_ffd+_ee]); _bc != nil {
					return _bc
				}
			}
		}
	}
	return nil
}

func _deedb(_ggab, _aebeb NRGBA, _dacgf _a.Rectangle) {
	for _eegc := 0; _eegc < _dacgf.Max.X; _eegc++ {
		for _ddab := 0; _ddab < _dacgf.Max.Y; _ddab++ {
			_aebeb.SetNRGBA(_eegc, _ddab, _ggab.NRGBAAt(_eegc, _ddab))
		}
	}
}

func (_fffd *ImageBase) setEightBytes(_ggaa int, _ffg uint64) error {
	_bgeb := _fffd.BytesPerLine - (_ggaa % _fffd.BytesPerLine)
	if _fffd.BytesPerLine != _fffd.Width>>3 {
		_bgeb--
	}
	if _bgeb >= 8 {
		return _fffd.setEightFullBytes(_ggaa, _ffg)
	}
	return _fffd.setEightPartlyBytes(_ggaa, _bgeb, _ffg)
}

func _gcbaf(_daac CMYK, _ggef NRGBA, _cgca _a.Rectangle) {
	for _bcgc := 0; _bcgc < _cgca.Max.X; _bcgc++ {
		for _fgdb := 0; _fgdb < _cgca.Max.Y; _fgdb++ {
			_eaa := _daac.CMYKAt(_bcgc, _fgdb)
			_ggef.SetNRGBA(_bcgc, _fgdb, _ead(_eaa))
		}
	}
}

func (_deab *Monochrome) ExpandBinary(factor int) (*Monochrome, error) {
	if !IsPowerOf2(uint(factor)) {
		return nil, _c.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0065\u0078\u0070\u0061\u006e\u0064\u0020b\u0069n\u0061\u0072\u0079\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", factor)
	}
	return _dg(_deab, factor)
}

func _ddb(_da, _bcca int, _ec []byte) *Monochrome {
	_ccdg := _bce(_da, _bcca)
	_ccdg.Data = _ec
	return _ccdg
}

var _ Image = &NRGBA64{}

func (_gfdc *RGBA32) RGBAAt(x, y int) _g.RGBA {
	_cbea, _ := ColorAtRGBA32(x, y, _gfdc.Width, _gfdc.Data, _gfdc.Alpha, _gfdc.Decode)
	return _cbea
}

func _bbac(_ceggb []byte, _aeddg Image) error {
	_abbg := true
	for _bbade := 0; _bbade < len(_ceggb); _bbade++ {
		if _ceggb[_bbade] != 0xff {
			_abbg = false
			break
		}
	}
	if _abbg {
		switch _addg := _aeddg.(type) {
		case *NRGBA32:
			_addg.Alpha = nil
		case *NRGBA64:
			_addg.Alpha = nil
		default:
			return _c.Errorf("i\u006ete\u0072n\u0061l\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020i\u006d\u0061\u0067\u0065\u0020s\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070e\u0020\u002a\u004eRGB\u0041\u0033\u0032\u0020\u006f\u0072 \u002a\u004e\u0052\u0047\u0042\u0041\u0036\u0034\u0020\u0062\u0075\u0074 \u0069s\u003a\u0020\u0025\u0054", _aeddg)
		}
	}
	return nil
}

func (_fdcf *Monochrome) setGray(_gcbg int, _egeb _g.Gray, _ecb int) {
	if _egeb.Y == 0 {
		_fdcf.clearBit(_ecb, _gcbg)
	} else {
		_fdcf.setGrayBit(_ecb, _gcbg)
	}
}
func IsPowerOf2(n uint) bool { return n > 0 && (n&(n-1)) == 0 }
func _caed(_gcfa int, _egag int) error {
	return _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", _gcfa, _egag)
}

func (_dfaa *ImageBase) setFourBytes(_gcee int, _bfdd uint32) error {
	if _gcee+3 > len(_dfaa.Data)-1 {
		return _c.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _gcee)
	}
	_dfaa.Data[_gcee] = byte((_bfdd & 0xff000000) >> 24)
	_dfaa.Data[_gcee+1] = byte((_bfdd & 0xff0000) >> 16)
	_dfaa.Data[_gcee+2] = byte((_bfdd & 0xff00) >> 8)
	_dfaa.Data[_gcee+3] = byte(_bfdd & 0xff)
	return nil
}

func _bcdg(_agd Gray, _faa NRGBA, _agb _a.Rectangle) {
	for _aeca := 0; _aeca < _agb.Max.X; _aeca++ {
		for _gfaf := 0; _gfaf < _agb.Max.Y; _gfaf++ {
			_dbg := _dfd(_faa.NRGBAAt(_aeca, _gfaf))
			_agd.SetGray(_aeca, _gfaf, _dbg)
		}
	}
}

func _ecdd(_edbd NRGBA, _gcbf RGBA, _fbefe _a.Rectangle) {
	for _edfcd := 0; _edfcd < _fbefe.Max.X; _edfcd++ {
		for _defe := 0; _defe < _fbefe.Max.Y; _defe++ {
			_dgdc := _edbd.NRGBAAt(_edfcd, _defe)
			_gcbf.SetRGBA(_edfcd, _defe, _egcg(_dgdc))
		}
	}
}

func _eeggg(_fbdf _a.Image) (Image, error) {
	if _fbccc, _cafc := _fbdf.(*RGBA32); _cafc {
		return _fbccc.Copy(), nil
	}
	_bbbgfc, _cbfdf, _egfd := _gfae(_fbdf, 1)
	_dabe := &RGBA32{ImageBase: NewImageBase(_bbbgfc.Max.X, _bbbgfc.Max.Y, 8, 3, nil, _egfd, nil)}
	_ccfb(_fbdf, _dabe, _bbbgfc)
	if len(_egfd) != 0 && !_cbfdf {
		if _gcgab := _bbac(_egfd, _dabe); _gcgab != nil {
			return nil, _gcgab
		}
	}
	return _dabe, nil
}

func (_fbfg *Monochrome) AddPadding() (_bcab error) {
	if _eddf := ((_fbfg.Width * _fbfg.Height) + 7) >> 3; len(_fbfg.Data) < _eddf {
		return _c.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(_fbfg.Data), _eddf)
	}
	_dgag := _fbfg.Width % 8
	if _dgag == 0 {
		return nil
	}
	_acac := _fbfg.Width / 8
	_eed := _ab.NewReader(_fbfg.Data)
	_dgfd := make([]byte, _fbfg.Height*_fbfg.BytesPerLine)
	_cadc := _ab.NewWriterMSB(_dgfd)
	_fbbg := make([]byte, _acac)
	var (
		_bceb int
		_ecc  uint64
	)
	for _bceb = 0; _bceb < _fbfg.Height; _bceb++ {
		if _, _bcab = _eed.Read(_fbbg); _bcab != nil {
			return _bcab
		}
		if _, _bcab = _cadc.Write(_fbbg); _bcab != nil {
			return _bcab
		}
		if _ecc, _bcab = _eed.ReadBits(byte(_dgag)); _bcab != nil {
			return _bcab
		}
		if _bcab = _cadc.WriteByte(byte(_ecc) << uint(8-_dgag)); _bcab != nil {
			return _bcab
		}
	}
	_fbfg.Data = _cadc.Data()
	return nil
}

func _cgf(_edf NRGBA, _ecd CMYK, _ffe _a.Rectangle) {
	for _afg := 0; _afg < _ffe.Max.X; _afg++ {
		for _dcdb := 0; _dcdb < _ffe.Max.Y; _dcdb++ {
			_fca := _edf.NRGBAAt(_afg, _dcdb)
			_ecd.SetCMYK(_afg, _dcdb, _bed(_fca))
		}
	}
}

func (_ebed *NRGBA16) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _ebed.Width, Y: _ebed.Height}}
}

func (_fgee *Gray2) Validate() error {
	if len(_fgee.Data) != _fgee.Height*_fgee.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func _dgga(_acb _g.Gray) _g.CMYK { return _g.CMYK{K: 0xff - _acb.Y} }
func (_cgabf *RGBA32) Set(x, y int, c _g.Color) {
	_daba := y*_cgabf.Width + x
	_caea := 3 * _daba
	if _caea+2 >= len(_cgabf.Data) {
		return
	}
	_bdbf := _g.RGBAModel.Convert(c).(_g.RGBA)
	_cgabf.setRGBA(_daba, _bdbf)
}

type NRGBA32 struct{ ImageBase }

func (_ggad *Gray2) Copy() Image { return &Gray2{ImageBase: _ggad.copy()} }
func (_aagf *Gray8) Validate() error {
	if len(_aagf.Data) != _aagf.Height*_aagf.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

func (_fdbc *ImageBase) newAlpha() {
	_dde := BytesPerLine(_fdbc.Width, _fdbc.BitsPerComponent, 1)
	_fdbc.Alpha = make([]byte, _fdbc.Height*_dde)
}
func (_fdgd *NRGBA64) Base() *ImageBase { return &_fdgd.ImageBase }
func _efgd(_cffc RGBA, _efdgf Gray, _aaf _a.Rectangle) {
	for _abf := 0; _abf < _aaf.Max.X; _abf++ {
		for _eacb := 0; _eacb < _aaf.Max.Y; _eacb++ {
			_ceb := _dgdd(_cffc.RGBAAt(_abf, _eacb))
			_efdgf.SetGray(_abf, _eacb, _ceb)
		}
	}
}

func (_gbeb *NRGBA32) Set(x, y int, c _g.Color) {
	_egfb := y*_gbeb.Width + x
	_ccdb := 3 * _egfb
	if _ccdb+2 >= len(_gbeb.Data) {
		return
	}
	_adbfd := _g.NRGBAModel.Convert(c).(_g.NRGBA)
	_gbeb.setRGBA(_egfb, _adbfd)
}
func (_gcga *NRGBA32) Copy() Image { return &NRGBA32{ImageBase: _gcga.copy()} }
func _dabb(_afaa Gray, _eege RGBA, _cged _a.Rectangle) {
	for _egaf := 0; _egaf < _cged.Max.X; _egaf++ {
		for _agcf := 0; _agcf < _cged.Max.Y; _agcf++ {
			_bceg := _afaa.GrayAt(_egaf, _agcf)
			_eege.SetRGBA(_egaf, _agcf, _efgb(_bceg))
		}
	}
}

var _ Image = &Gray8{}

type Gray8 struct{ ImageBase }

var _ _a.Image = &NRGBA32{}

func _eacgf(_eagd nrgba64, _fdfgd NRGBA, _fabcd _a.Rectangle) {
	for _cdgg := 0; _cdgg < _fabcd.Max.X; _cdgg++ {
		for _fdec := 0; _fdec < _fabcd.Max.Y; _fdec++ {
			_abef := _eagd.NRGBA64At(_cdgg, _fdec)
			_fdfgd.SetNRGBA(_cdgg, _fdec, _caa(_abef))
		}
	}
}

func _eab(_cda _g.RGBA) _g.CMYK {
	_cefg, _aaa, _fabg, _fcbc := _g.RGBToCMYK(_cda.R, _cda.G, _cda.B)
	return _g.CMYK{C: _cefg, M: _aaa, Y: _fabg, K: _fcbc}
}

var _ _a.Image = &NRGBA64{}

func _afdbf(_bbba _a.Image) (Image, error) {
	if _dggc, _egbd := _bbba.(*NRGBA16); _egbd {
		return _dggc.Copy(), nil
	}
	_gagf := _bbba.Bounds()
	_ccea, _abfa := NewImage(_gagf.Max.X, _gagf.Max.Y, 4, 3, nil, nil, nil)
	if _abfa != nil {
		return nil, _abfa
	}
	_aefb(_bbba, _ccea, _gagf)
	return _ccea, nil
}
func (_dfce *NRGBA64) At(x, y int) _g.Color { _abad, _ := _dfce.ColorAt(x, y); return _abad }
func (_gaa *ImageBase) copy() ImageBase {
	_eccbe := *_gaa
	_eccbe.Data = make([]byte, len(_gaa.Data))
	copy(_eccbe.Data, _gaa.Data)
	return _eccbe
}

func (_dbaec *RGBA32) setRGBA(_aebb int, _gedga _g.RGBA) {
	_caee := 3 * _aebb
	_dbaec.Data[_caee] = _gedga.R
	_dbaec.Data[_caee+1] = _gedga.G
	_dbaec.Data[_caee+2] = _gedga.B
	if _aebb < len(_dbaec.Alpha) {
		_dbaec.Alpha[_aebb] = _gedga.A
	}
}

func (_ebb *Gray2) Histogram() (_egab [256]int) {
	for _bdcf := 0; _bdcf < _ebb.Width; _bdcf++ {
		for _gead := 0; _gead < _ebb.Height; _gead++ {
			_egab[_ebb.GrayAt(_bdcf, _gead).Y]++
		}
	}
	return _egab
}

func ColorAtNRGBA16(x, y, width, bytesPerLine int, data, alpha []byte, decode []float64) (_g.NRGBA, error) {
	_afbeb := y*bytesPerLine + x*3/2
	if _afbeb+1 >= len(data) {
		return _g.NRGBA{}, _caed(x, y)
	}
	const (
		_febec = 0xf
		_aaegc = uint8(0xff)
	)
	_cdeb := _aaegc
	if alpha != nil {
		_fded := y * BytesPerLine(width, 4, 1)
		if _fded < len(alpha) {
			if x%2 == 0 {
				_cdeb = (alpha[_fded] >> uint(4)) & _febec
			} else {
				_cdeb = alpha[_fded] & _febec
			}
			_cdeb |= _cdeb << 4
		}
	}
	var _bbadb, _eedc, _facg uint8
	if x*3%2 == 0 {
		_bbadb = (data[_afbeb] >> uint(4)) & _febec
		_eedc = data[_afbeb] & _febec
		_facg = (data[_afbeb+1] >> uint(4)) & _febec
	} else {
		_bbadb = data[_afbeb] & _febec
		_eedc = (data[_afbeb+1] >> uint(4)) & _febec
		_facg = data[_afbeb+1] & _febec
	}
	if len(decode) == 6 {
		_bbadb = uint8(uint32(LinearInterpolate(float64(_bbadb), 0, 15, decode[0], decode[1])) & 0xf)
		_eedc = uint8(uint32(LinearInterpolate(float64(_eedc), 0, 15, decode[2], decode[3])) & 0xf)
		_facg = uint8(uint32(LinearInterpolate(float64(_facg), 0, 15, decode[4], decode[5])) & 0xf)
	}
	return _g.NRGBA{R: (_bbadb << 4) | (_bbadb & 0xf), G: (_eedc << 4) | (_eedc & 0xf), B: (_facg << 4) | (_facg & 0xf), A: _cdeb}, nil
}

func _gaeg(_dfb, _cab *Monochrome, _bdc []byte, _bca int) (_abd error) {
	var (
		_eddc, _gccf, _cabe, _efg, _dca, _ace, _fde, _gac int
		_eac, _dac                                        uint32
		_ecg, _eaf                                        byte
		_aedb                                             uint16
	)
	_bage := make([]byte, 4)
	_fgg := make([]byte, 4)
	for _cabe = 0; _cabe < _dfb.Height-1; _cabe, _efg = _cabe+2, _efg+1 {
		_eddc = _cabe * _dfb.BytesPerLine
		_gccf = _efg * _cab.BytesPerLine
		for _dca, _ace = 0, 0; _dca < _bca; _dca, _ace = _dca+4, _ace+1 {
			for _fde = 0; _fde < 4; _fde++ {
				_gac = _eddc + _dca + _fde
				if _gac <= len(_dfb.Data)-1 && _gac < _eddc+_dfb.BytesPerLine {
					_bage[_fde] = _dfb.Data[_gac]
				} else {
					_bage[_fde] = 0x00
				}
				_gac = _eddc + _dfb.BytesPerLine + _dca + _fde
				if _gac <= len(_dfb.Data)-1 && _gac < _eddc+(2*_dfb.BytesPerLine) {
					_fgg[_fde] = _dfb.Data[_gac]
				} else {
					_fgg[_fde] = 0x00
				}
			}
			_eac = _efc.BigEndian.Uint32(_bage)
			_dac = _efc.BigEndian.Uint32(_fgg)
			_dac &= _eac
			_dac &= _dac << 1
			_dac &= 0xaaaaaaaa
			_eac = _dac | (_dac << 7)
			_ecg = byte(_eac >> 24)
			_eaf = byte((_eac >> 8) & 0xff)
			_gac = _gccf + _ace
			if _gac+1 == len(_cab.Data)-1 || _gac+1 >= _gccf+_cab.BytesPerLine {
				_cab.Data[_gac] = _bdc[_ecg]
				if _abd = _cab.setByte(_gac, _bdc[_ecg]); _abd != nil {
					return _c.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _gac)
				}
			} else {
				_aedb = (uint16(_bdc[_ecg]) << 8) | uint16(_bdc[_eaf])
				if _abd = _cab.setTwoBytes(_gac, _aedb); _abd != nil {
					return _c.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _gac)
				}
				_ace++
			}
		}
	}
	return nil
}

func _gad(_cfbd _g.NRGBA64) _g.Gray {
	var _dccb _g.NRGBA64
	if _cfbd == _dccb {
		return _g.Gray{Y: 0xff}
	}
	_dcf, _egg, _aebe, _ := _cfbd.RGBA()
	_cgg := (19595*_dcf + 38470*_egg + 7471*_aebe + 1<<15) >> 24
	return _g.Gray{Y: uint8(_cgg)}
}

func (_afcf *Monochrome) Validate() error {
	if len(_afcf.Data) != _afcf.Height*_afcf.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

func (_dbbgg *Gray4) setGray(_egbg int, _bda int, _eeff _g.Gray) {
	_bdbc := _bda * _dbbgg.BytesPerLine
	_ebeb := _bdbc + (_egbg >> 1)
	if _ebeb >= len(_dbbgg.Data) {
		return
	}
	_fcec := _eeff.Y >> 4
	_dbbgg.Data[_ebeb] = (_dbbgg.Data[_ebeb] & (^(0xf0 >> uint(4*(_egbg&1))))) | (_fcec << uint(4-4*(_egbg&1)))
}

var _ Image = &NRGBA16{}

func (_cabf *Gray16) ColorModel() _g.Model { return _g.Gray16Model }

type CMYK interface {
	CMYKAt(_dgaa, _eef int) _g.CMYK
	SetCMYK(_fbe, _gde int, _acd _g.CMYK)
}

func (_agde *NRGBA32) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _agde.Width, Y: _agde.Height}}
}

func _ggeb(_cadg, _dggab uint8) uint8 {
	if _cadg < _dggab {
		return 255
	}
	return 0
}

var (
	_ea  = _gcd()
	_cca = _acc()
	_cfb = _dc()
)

func (_aefdb *Monochrome) clearBit(_ecf, _bccf int) { _aefdb.Data[_ecf] &= ^(0x80 >> uint(_bccf&7)) }
func ColorAtCMYK(x, y, width int, data []byte, decode []float64) (_g.CMYK, error) {
	_cg := 4 * (y*width + x)
	if _cg+3 >= len(data) {
		return _g.CMYK{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	C := data[_cg] & 0xff
	M := data[_cg+1] & 0xff
	Y := data[_cg+2] & 0xff
	K := data[_cg+3] & 0xff
	if len(decode) == 8 {
		C = uint8(uint32(LinearInterpolate(float64(C), 0, 255, decode[0], decode[1])) & 0xff)
		M = uint8(uint32(LinearInterpolate(float64(M), 0, 255, decode[2], decode[3])) & 0xff)
		Y = uint8(uint32(LinearInterpolate(float64(Y), 0, 255, decode[4], decode[5])) & 0xff)
		K = uint8(uint32(LinearInterpolate(float64(K), 0, 255, decode[6], decode[7])) & 0xff)
	}
	return _g.CMYK{C: C, M: M, Y: Y, K: K}, nil
}

func (_dgac *ImageBase) setEightFullBytes(_fba int, _dgecd uint64) error {
	if _fba+7 > len(_dgac.Data)-1 {
		return _d.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_dgac.Data[_fba] = byte((_dgecd & 0xff00000000000000) >> 56)
	_dgac.Data[_fba+1] = byte((_dgecd & 0xff000000000000) >> 48)
	_dgac.Data[_fba+2] = byte((_dgecd & 0xff0000000000) >> 40)
	_dgac.Data[_fba+3] = byte((_dgecd & 0xff00000000) >> 32)
	_dgac.Data[_fba+4] = byte((_dgecd & 0xff000000) >> 24)
	_dgac.Data[_fba+5] = byte((_dgecd & 0xff0000) >> 16)
	_dgac.Data[_fba+6] = byte((_dgecd & 0xff00) >> 8)
	_dgac.Data[_fba+7] = byte(_dgecd & 0xff)
	return nil
}

func ColorAtGray2BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_abgf := y*bytesPerLine + x>>2
	if _abgf >= len(data) {
		return _g.Gray{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_dbde := data[_abgf] >> uint(6-(x&3)*2) & 3
	if len(decode) == 2 {
		_dbde = uint8(uint32(LinearInterpolate(float64(_dbde), 0, 3.0, decode[0], decode[1])) & 3)
	}
	return _g.Gray{Y: _dbde * 85}, nil
}

var _ NRGBA = &NRGBA16{}

func _agbc(_edgg *Monochrome, _ddgfd, _ccba, _afba, _aafg int, _dfbc RasterOperator, _ddcd *Monochrome, _aaaf, _fbfge int) error {
	if _edgg == nil {
		return _d.New("\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _dfbc == PixDst {
		return nil
	}
	switch _dfbc {
	case PixClr, PixSet, PixNotDst:
		_bfec(_edgg, _ddgfd, _ccba, _afba, _aafg, _dfbc)
		return nil
	}
	if _ddcd == nil {
		_af.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _d.New("\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _fgfc := _gcba(_edgg, _ddgfd, _ccba, _afba, _aafg, _dfbc, _ddcd, _aaaf, _fbfge); _fgfc != nil {
		return _fgfc
	}
	return nil
}

func _ccfb(_cbfc _a.Image, _adfcd Image, _aefa _a.Rectangle) {
	if _aacc, _fagc := _cbfc.(SMasker); _fagc && _aacc.HasAlpha() {
		_adfcd.(SMasker).MakeAlpha()
	}
	switch _fgcdg := _cbfc.(type) {
	case Gray:
		_dabb(_fgcdg, _adfcd.(RGBA), _aefa)
	case NRGBA:
		_ecdd(_fgcdg, _adfcd.(RGBA), _aefa)
	case *_a.NYCbCrA:
		_efef(_fgcdg, _adfcd.(RGBA), _aefa)
	case CMYK:
		_dbebe(_fgcdg, _adfcd.(RGBA), _aefa)
	case RGBA:
		_afge(_fgcdg, _adfcd.(RGBA), _aefa)
	case nrgba64:
		_abadbf(_fgcdg, _adfcd.(RGBA), _aefa)
	default:
		_bgg(_cbfc, _adfcd, _aefa)
	}
}

var _ RGBA = &RGBA32{}

func (_adfga *monochromeThresholdConverter) Convert(img _a.Image) (Image, error) {
	if _deae, _cdgf := img.(*Monochrome); _cdgf {
		return _deae.Copy(), nil
	}
	_baae := img.Bounds()
	_acbf, _edfc := NewImage(_baae.Max.X, _baae.Max.Y, 1, 1, nil, nil, nil)
	if _edfc != nil {
		return nil, _edfc
	}
	_acbf.(*Monochrome).ModelThreshold = _adfga.Threshold
	for _ecda := 0; _ecda < _baae.Max.X; _ecda++ {
		for _aecae := 0; _aecae < _baae.Max.Y; _aecae++ {
			_gaegd := img.At(_ecda, _aecae)
			_acbf.Set(_ecda, _aecae, _gaegd)
		}
	}
	return _acbf, nil
}
func (_bcfa *CMYK32) Base() *ImageBase { return &_bcfa.ImageBase }
func (_eba *Monochrome) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray1BPC(x, y, _eba.BytesPerLine, _eba.Data, _eba.Decode)
}

func _dcag(_egac Gray, _fefd NRGBA, _afcfe _a.Rectangle) {
	for _afdfe := 0; _afdfe < _afcfe.Max.X; _afdfe++ {
		for _bcfaa := 0; _bcfaa < _afcfe.Max.Y; _bcfaa++ {
			_afec := _egac.GrayAt(_afdfe, _bcfaa)
			_fefd.SetNRGBA(_afdfe, _bcfaa, _eaedg(_afec))
		}
	}
}
func (_ceff *CMYK32) ColorModel() _g.Model { return _g.CMYKModel }
func _gfae(_bgd _a.Image, _eeba int) (_a.Rectangle, bool, []byte) {
	_gcgg := _bgd.Bounds()
	var (
		_adgf bool
		_debd []byte
	)
	switch _ccdfc := _bgd.(type) {
	case SMasker:
		_adgf = _ccdfc.HasAlpha()
	case NRGBA, RGBA, *_a.RGBA64, nrgba64, *_a.NYCbCrA:
		_debd = make([]byte, _gcgg.Max.X*_gcgg.Max.Y*_eeba)
	case *_a.Paletted:
		var _gcfgd bool
		for _, _eacf := range _ccdfc.Palette {
			_bgfc, _ggdc, _eddaa, _bac := _eacf.RGBA()
			if _bgfc == 0 && _ggdc == 0 && _eddaa == 0 && _bac != 0 {
				_gcfgd = true
				break
			}
		}
		if _gcfgd {
			_debd = make([]byte, _gcgg.Max.X*_gcgg.Max.Y*_eeba)
		}
	}
	return _gcgg, _adgf, _debd
}

func (_cfae *NRGBA16) Validate() error {
	if len(_cfae.Data) != 3*_cfae.Width*_cfae.Height/2 {
		return _d.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func (_gagg *Gray2) Base() *ImageBase { return &_gagg.ImageBase }
func _gaff(_eefc int, _edgc int) int {
	if _eefc < _edgc {
		return _eefc
	}
	return _edgc
}
func (_acdb *Gray16) Copy() Image { return &Gray16{ImageBase: _acdb.copy()} }
func (_edea *Gray4) At(x, y int) _g.Color {
	_bade, _ := _edea.ColorAt(x, y)
	return _bade
}

func _cdg(_dad, _bdce CMYK, _cbaa _a.Rectangle) {
	for _cefb := 0; _cefb < _cbaa.Max.X; _cefb++ {
		for _eaed := 0; _eaed < _cbaa.Max.Y; _eaed++ {
			_bdce.SetCMYK(_cefb, _eaed, _dad.CMYKAt(_cefb, _eaed))
		}
	}
}

func (_ebae *Gray2) Bounds() _a.Rectangle {
	return _a.Rectangle{Max: _a.Point{X: _ebae.Width, Y: _ebae.Height}}
}
func (_agfa *NRGBA16) Copy() Image { return &NRGBA16{ImageBase: _agfa.copy()} }

type nrgba64 interface {
	NRGBA64At(_aaad, _dgcg int) _g.NRGBA64
	SetNRGBA64(_geeg, _ceee int, _bfeef _g.NRGBA64)
}

func _dcgc(_cedf _g.NYCbCrA) _g.NRGBA {
	_baa := int32(_cedf.Y) * 0x10101
	_dea := int32(_cedf.Cb) - 128
	_acgd := int32(_cedf.Cr) - 128
	_gedg := _baa + 91881*_acgd
	if uint32(_gedg)&0xff000000 == 0 {
		_gedg >>= 8
	} else {
		_gedg = ^(_gedg >> 31) & 0xffff
	}
	_fbed := _baa - 22554*_dea - 46802*_acgd
	if uint32(_fbed)&0xff000000 == 0 {
		_fbed >>= 8
	} else {
		_fbed = ^(_fbed >> 31) & 0xffff
	}
	_ggbe := _baa + 116130*_dea
	if uint32(_ggbe)&0xff000000 == 0 {
		_ggbe >>= 8
	} else {
		_ggbe = ^(_ggbe >> 31) & 0xffff
	}
	return _g.NRGBA{R: uint8(_gedg >> 8), G: uint8(_fbed >> 8), B: uint8(_ggbe >> 8), A: _cedf.A}
}
func (_fbcg *Monochrome) setIndexedBit(_fgdf int) { _fbcg.Data[(_fgdf >> 3)] |= 0x80 >> uint(_fgdf&7) }
func _bgg(_edda _a.Image, _fee Image, _beec _a.Rectangle) {
	for _efa := 0; _efa < _beec.Max.X; _efa++ {
		for _fgec := 0; _fgec < _beec.Max.Y; _fgec++ {
			_ced := _edda.At(_efa, _fgec)
			_fee.Set(_efa, _fgec, _ced)
		}
	}
}

func NewImageBase(width int, height int, bitsPerComponent int, colorComponents int, data []byte, alpha []byte, decode []float64) ImageBase {
	_ccbc := ImageBase{Width: width, Height: height, BitsPerComponent: bitsPerComponent, ColorComponents: colorComponents, Data: data, Alpha: alpha, Decode: decode, BytesPerLine: BytesPerLine(width, bitsPerComponent, colorComponents)}
	if data == nil {
		_ccbc.Data = make([]byte, height*_ccbc.BytesPerLine)
	}
	return _ccbc
}

func (_bagee *ImageBase) setTwoBytes(_ebcdg int, _acgdb uint16) error {
	if _ebcdg+1 > len(_bagee.Data)-1 {
		return _d.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bagee.Data[_ebcdg] = byte((_acgdb & 0xff00) >> 8)
	_bagee.Data[_ebcdg+1] = byte(_acgdb & 0xff)
	return nil
}

type NRGBA interface {
	NRGBAAt(_beef, _bgbe int) _g.NRGBA
	SetNRGBA(_deg, _aedc int, _ecffe _g.NRGBA)
}

func _ccg(_ffcf, _fdb *Monochrome, _fegc []byte, _efdd int) (_ccff error) {
	var (
		_fabc, _dgdg, _bbb, _ggf, _ddg, _efdc, _aeb, _cae int
		_dcb, _cdd, _bbd, _fddf                           uint32
		_ffaf, _ada                                       byte
		_ffag                                             uint16
	)
	_cfe := make([]byte, 4)
	_abg := make([]byte, 4)
	for _bbb = 0; _bbb < _ffcf.Height-1; _bbb, _ggf = _bbb+2, _ggf+1 {
		_fabc = _bbb * _ffcf.BytesPerLine
		_dgdg = _ggf * _fdb.BytesPerLine
		for _ddg, _efdc = 0, 0; _ddg < _efdd; _ddg, _efdc = _ddg+4, _efdc+1 {
			for _aeb = 0; _aeb < 4; _aeb++ {
				_cae = _fabc + _ddg + _aeb
				if _cae <= len(_ffcf.Data)-1 && _cae < _fabc+_ffcf.BytesPerLine {
					_cfe[_aeb] = _ffcf.Data[_cae]
				} else {
					_cfe[_aeb] = 0x00
				}
				_cae = _fabc + _ffcf.BytesPerLine + _ddg + _aeb
				if _cae <= len(_ffcf.Data)-1 && _cae < _fabc+(2*_ffcf.BytesPerLine) {
					_abg[_aeb] = _ffcf.Data[_cae]
				} else {
					_abg[_aeb] = 0x00
				}
			}
			_dcb = _efc.BigEndian.Uint32(_cfe)
			_cdd = _efc.BigEndian.Uint32(_abg)
			_bbd = _dcb & _cdd
			_bbd |= _bbd << 1
			_fddf = _dcb | _cdd
			_fddf &= _fddf << 1
			_cdd = _bbd | _fddf
			_cdd &= 0xaaaaaaaa
			_dcb = _cdd | (_cdd << 7)
			_ffaf = byte(_dcb >> 24)
			_ada = byte((_dcb >> 8) & 0xff)
			_cae = _dgdg + _efdc
			if _cae+1 == len(_fdb.Data)-1 || _cae+1 >= _dgdg+_fdb.BytesPerLine {
				if _ccff = _fdb.setByte(_cae, _fegc[_ffaf]); _ccff != nil {
					return _c.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _cae)
				}
			} else {
				_ffag = (uint16(_fegc[_ffaf]) << 8) | uint16(_fegc[_ada])
				if _ccff = _fdb.setTwoBytes(_cae, _ffag); _ccff != nil {
					return _c.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _cae)
				}
				_efdc++
			}
		}
	}
	return nil
}

func _gfga(_cadd uint) uint {
	var _fggea uint
	for _cadd != 0 {
		_cadd >>= 1
		_fggea++
	}
	return _fggea - 1
}

var _ Image = &Gray16{}

func _baf(_bdf _g.Color) _g.Color { _acf := _g.GrayModel.Convert(_bdf).(_g.Gray); return _cabb(_acf) }
func _abbe(_bbg *_a.Gray) bool {
	for _ffaa := 0; _ffaa < len(_bbg.Pix); _ffaa++ {
		if !_bbfcbc(_bbg.Pix[_ffaa]) {
			return false
		}
	}
	return true
}

func ColorAtRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_g.RGBA, error) {
	_acba := y*width + x
	_dgacc := 3 * _acba
	if _dgacc+2 >= len(data) {
		return _g.RGBA{}, _c.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_efec := uint8(0xff)
	if alpha != nil && len(alpha) > _acba {
		_efec = alpha[_acba]
	}
	_eeggb, _dgfgb, _edbb := data[_dgacc], data[_dgacc+1], data[_dgacc+2]
	if len(decode) == 6 {
		_eeggb = uint8(uint32(LinearInterpolate(float64(_eeggb), 0, 255, decode[0], decode[1])) & 0xff)
		_dgfgb = uint8(uint32(LinearInterpolate(float64(_dgfgb), 0, 255, decode[2], decode[3])) & 0xff)
		_edbb = uint8(uint32(LinearInterpolate(float64(_edbb), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _g.RGBA{R: _eeggb, G: _dgfgb, B: _edbb, A: _efec}, nil
}

var _ Image = &CMYK32{}

func LinearInterpolate(x, xmin, xmax, ymin, ymax float64) float64 {
	if _ef.Abs(xmax-xmin) < 0.000001 {
		return ymin
	}
	_agac := ymin + (x-xmin)*(ymax-ymin)/(xmax-xmin)
	return _agac
}
func _efgb(_aceg _g.Gray) _g.RGBA { return _g.RGBA{R: _aceg.Y, G: _aceg.Y, B: _aceg.Y, A: 0xff} }
func _gcgb(_gecb *Monochrome, _cbdb, _cafg int, _babcc, _dacg int, _aead RasterOperator) {
	var (
		_bbbgf bool
		_fdebd bool
		_ccc   int
		_ddafc int
		_fdfg  int
		_gage  int
		_bead  bool
		_fbaa  byte
	)
	_efgae := 8 - (_cbdb & 7)
	_ffcg := _ffagf[_efgae]
	_bfgg := _gecb.BytesPerLine*_cafg + (_cbdb >> 3)
	if _babcc < _efgae {
		_bbbgf = true
		_ffcg &= _bafab[8-_efgae+_babcc]
	}
	if !_bbbgf {
		_ccc = (_babcc - _efgae) >> 3
		if _ccc != 0 {
			_fdebd = true
			_ddafc = _bfgg + 1
		}
	}
	_fdfg = (_cbdb + _babcc) & 7
	if !(_bbbgf || _fdfg == 0) {
		_bead = true
		_fbaa = _bafab[_fdfg]
		_gage = _bfgg + 1 + _ccc
	}
	var _edcb, _dedbb int
	switch _aead {
	case PixClr:
		for _edcb = 0; _edcb < _dacg; _edcb++ {
			_gecb.Data[_bfgg] = _decfa(_gecb.Data[_bfgg], 0x0, _ffcg)
			_bfgg += _gecb.BytesPerLine
		}
		if _fdebd {
			for _edcb = 0; _edcb < _dacg; _edcb++ {
				for _dedbb = 0; _dedbb < _ccc; _dedbb++ {
					_gecb.Data[_ddafc+_dedbb] = 0x0
				}
				_ddafc += _gecb.BytesPerLine
			}
		}
		if _bead {
			for _edcb = 0; _edcb < _dacg; _edcb++ {
				_gecb.Data[_gage] = _decfa(_gecb.Data[_gage], 0x0, _fbaa)
				_gage += _gecb.BytesPerLine
			}
		}
	case PixSet:
		for _edcb = 0; _edcb < _dacg; _edcb++ {
			_gecb.Data[_bfgg] = _decfa(_gecb.Data[_bfgg], 0xff, _ffcg)
			_bfgg += _gecb.BytesPerLine
		}
		if _fdebd {
			for _edcb = 0; _edcb < _dacg; _edcb++ {
				for _dedbb = 0; _dedbb < _ccc; _dedbb++ {
					_gecb.Data[_ddafc+_dedbb] = 0xff
				}
				_ddafc += _gecb.BytesPerLine
			}
		}
		if _bead {
			for _edcb = 0; _edcb < _dacg; _edcb++ {
				_gecb.Data[_gage] = _decfa(_gecb.Data[_gage], 0xff, _fbaa)
				_gage += _gecb.BytesPerLine
			}
		}
	case PixNotDst:
		for _edcb = 0; _edcb < _dacg; _edcb++ {
			_gecb.Data[_bfgg] = _decfa(_gecb.Data[_bfgg], ^_gecb.Data[_bfgg], _ffcg)
			_bfgg += _gecb.BytesPerLine
		}
		if _fdebd {
			for _edcb = 0; _edcb < _dacg; _edcb++ {
				for _dedbb = 0; _dedbb < _ccc; _dedbb++ {
					_gecb.Data[_ddafc+_dedbb] = ^(_gecb.Data[_ddafc+_dedbb])
				}
				_ddafc += _gecb.BytesPerLine
			}
		}
		if _bead {
			for _edcb = 0; _edcb < _dacg; _edcb++ {
				_gecb.Data[_gage] = _decfa(_gecb.Data[_gage], ^_gecb.Data[_gage], _fbaa)
				_gage += _gecb.BytesPerLine
			}
		}
	}
}

func (_ageb *NRGBA16) setNRGBA(_dbed, _becc, _cdab int, _ggeaa _g.NRGBA) {
	if _dbed*3%2 == 0 {
		_ageb.Data[_cdab] = (_ggeaa.R>>4)<<4 | (_ggeaa.G >> 4)
		_ageb.Data[_cdab+1] = (_ggeaa.B>>4)<<4 | (_ageb.Data[_cdab+1] & 0xf)
	} else {
		_ageb.Data[_cdab] = (_ageb.Data[_cdab] & 0xf0) | (_ggeaa.R >> 4)
		_ageb.Data[_cdab+1] = (_ggeaa.G>>4)<<4 | (_ggeaa.B >> 4)
	}
	if _ageb.Alpha != nil {
		_cacad := _becc * BytesPerLine(_ageb.Width, 4, 1)
		if _cacad < len(_ageb.Alpha) {
			if _dbed%2 == 0 {
				_ageb.Alpha[_cacad] = (_ggeaa.A>>uint(4))<<uint(4) | (_ageb.Alpha[_cdab] & 0xf)
			} else {
				_ageb.Alpha[_cacad] = (_ageb.Alpha[_cacad] & 0xf0) | (_ggeaa.A >> uint(4))
			}
		}
	}
}

func _dfef(_fccg *_a.Gray16, _gcdca uint8) *_a.Gray {
	_fbda := _fccg.Bounds()
	_dagg := _a.NewGray(_fbda)
	for _cfc := 0; _cfc < _fbda.Dx(); _cfc++ {
		for _afbf := 0; _afbf < _fbda.Dy(); _afbf++ {
			_adcb := _fccg.Gray16At(_cfc, _afbf)
			_dagg.SetGray(_cfc, _afbf, _g.Gray{Y: _ggeb(uint8(_adcb.Y/256), _gcdca)})
		}
	}
	return _dagg
}
func (_gbba *RGBA32) Base() *ImageBase { return &_gbba.ImageBase }

type monochromeThresholdConverter struct{ Threshold uint8 }

func (_aefe *NRGBA32) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtNRGBA32(x, y, _aefe.Width, _aefe.Data, _aefe.Alpha, _aefe.Decode)
}

func (_dgfg *Gray4) Set(x, y int, c _g.Color) {
	if x >= _dgfg.Width || y >= _dgfg.Height {
		return
	}
	_dcgcg := Gray4Model.Convert(c).(_g.Gray)
	_dgfg.setGray(x, y, _dcgcg)
}

var _ Gray = &Gray4{}

func (_fcad *Gray16) GrayAt(x, y int) _g.Gray {
	_dfcf, _ := _fcad.ColorAt(x, y)
	return _g.Gray{Y: uint8(_dfcf.(_g.Gray16).Y >> 8)}
}

func _deec(_dgec _g.NRGBA) _g.Gray {
	_bgce, _bbbg, _dbb, _ := _dgec.RGBA()
	_daef := (19595*_bgce + 38470*_bbbg + 7471*_dbb + 1<<15) >> 24
	return _g.Gray{Y: uint8(_daef)}
}

func _eefa(_bdfe _g.Gray) _g.Gray {
	_bdfe.Y >>= 4
	_bdfe.Y |= _bdfe.Y << 4
	return _bdfe
}
func (_fggb *ImageBase) MakeAlpha() { _fggb.newAlpha() }
func _eaac(_adcbf _a.Image, _fbgg uint8) *_a.Gray {
	_feec := _adcbf.Bounds()
	_geefg := _a.NewGray(_feec)
	var (
		_cecc _g.Color
		_gebf _g.Gray
	)
	for _bggg := 0; _bggg < _feec.Max.X; _bggg++ {
		for _cgba := 0; _cgba < _feec.Max.Y; _cgba++ {
			_cecc = _adcbf.At(_bggg, _cgba)
			_geefg.Set(_bggg, _cgba, _cecc)
			_gebf = _geefg.GrayAt(_bggg, _cgba)
			_geefg.SetGray(_bggg, _cgba, _g.Gray{Y: _ggeb(_gebf.Y, _fbgg)})
		}
	}
	return _geefg
}

func (_ggc *Gray2) GrayAt(x, y int) _g.Gray {
	_caec, _ := ColorAtGray2BPC(x, y, _ggc.BytesPerLine, _ggc.Data, _ggc.Decode)
	return _caec
}

func (_aeee *RGBA32) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtRGBA32(x, y, _aeee.Width, _aeee.Data, _aeee.Alpha, _aeee.Decode)
}

func _ggfge(_aeae *_a.NYCbCrA, _eecef NRGBA, _egaa _a.Rectangle) {
	for _cafd := 0; _cafd < _egaa.Max.X; _cafd++ {
		for _febc := 0; _febc < _egaa.Max.Y; _febc++ {
			_ebgc := _aeae.NYCbCrAAt(_cafd, _febc)
			_eecef.SetNRGBA(_cafd, _febc, _dcgc(_ebgc))
		}
	}
}
func (_ddc *Gray8) At(x, y int) _g.Color { _acbd, _ := _ddc.ColorAt(x, y); return _acbd }
func (_gfda *Gray4) SetGray(x, y int, g _g.Gray) {
	if x >= _gfda.Width || y >= _gfda.Height {
		return
	}
	g = _eefa(g)
	_gfda.setGray(x, y, g)
}
