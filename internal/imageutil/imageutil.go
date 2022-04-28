package imageutil

import (
	_be "encoding/binary"
	_a "errors"
	_ea "fmt"
	_c "image"
	_ag "image/color"
	_f "image/draw"
	_e "math"

	_ca "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
)

func (_geeg *Monochrome) getBit(_cgg, _aggc int) uint8 {
	return _geeg.Data[_cgg+(_aggc>>3)] >> uint(7-(_aggc&7)) & 1
}
func (_ggac *Gray2) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtGray2BPC(x, y, _ggac.BytesPerLine, _ggac.Data, _ggac.Decode)
}
func _fae(_dge _ag.NRGBA64) _ag.NRGBA {
	return _ag.NRGBA{R: uint8(_dge.R >> 8), G: uint8(_dge.G >> 8), B: uint8(_dge.B >> 8), A: uint8(_dge.A >> 8)}
}

type Gray8 struct{ ImageBase }

func _efde(_cdf _ag.CMYK) _ag.NRGBA {
	_gga, _ebg, _acgc := _ag.CMYKToRGB(_cdf.C, _cdf.M, _cdf.Y, _cdf.K)
	return _ag.NRGBA{R: _gga, G: _ebg, B: _acgc, A: 0xff}
}
func (_cebca *Gray2) GrayAt(x, y int) _ag.Gray {
	_ggff, _ := ColorAtGray2BPC(x, y, _cebca.BytesPerLine, _cebca.Data, _cebca.Decode)
	return _ggff
}
func (_bfae *ImageBase) newAlpha() {
	_abgc := BytesPerLine(_bfae.Width, _bfae.BitsPerComponent, 1)
	_bfae.Alpha = make([]byte, _bfae.Height*_abgc)
}
func ColorAtGray16BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_ag.Gray16, error) {
	_ebga := (y*bytesPerLine/2 + x) * 2
	if _ebga+1 >= len(data) {
		return _ag.Gray16{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_ddae := uint16(data[_ebga])<<8 | uint16(data[_ebga+1])
	if len(decode) == 2 {
		_ddae = uint16(uint64(LinearInterpolate(float64(_ddae), 0, 65535, decode[0], decode[1])))
	}
	return _ag.Gray16{Y: _ddae}, nil
}
func _bgge(_bffa *Monochrome, _gce, _edfc, _abdf, _beae int, _fbee RasterOperator, _cafa *Monochrome, _fgfe, _cefe int) error {
	if _bffa == nil {
		return _a.New("\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _fbee == PixDst {
		return nil
	}
	switch _fbee {
	case PixClr, PixSet, PixNotDst:
		_ecdb(_bffa, _gce, _edfc, _abdf, _beae, _fbee)
		return nil
	}
	if _cafa == nil {
		_ca.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _a.New("\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _fbeb := _bbd(_bffa, _gce, _edfc, _abdf, _beae, _fbee, _cafa, _fgfe, _cefe); _fbeb != nil {
		return _fbeb
	}
	return nil
}
func (_gabb *Gray8) ColorModel() _ag.Model  { return _ag.GrayModel }
func (_fceg *RGBA32) ColorModel() _ag.Model { return _ag.NRGBAModel }
func _aaff(_abafa []byte, _fdcde Image) error {
	_aedg := true
	for _feea := 0; _feea < len(_abafa); _feea++ {
		if _abafa[_feea] != 0xff {
			_aedg = false
			break
		}
	}
	if _aedg {
		switch _bbbdce := _fdcde.(type) {
		case *NRGBA32:
			_bbbdce.Alpha = nil
		case *NRGBA64:
			_bbbdce.Alpha = nil
		default:
			return _ea.Errorf("i\u006ete\u0072n\u0061l\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020i\u006d\u0061\u0067\u0065\u0020s\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070e\u0020\u002a\u004eRGB\u0041\u0033\u0032\u0020\u006f\u0072 \u002a\u004e\u0052\u0047\u0042\u0041\u0036\u0034\u0020\u0062\u0075\u0074 \u0069s\u003a\u0020\u0025\u0054", _fdcde)
		}
	}
	return nil
}
func _bed(_dag, _afd int, _bfa []byte) *Monochrome {
	_cb := _gea(_dag, _afd)
	_cb.Data = _bfa
	return _cb
}

var _ Gray = &Monochrome{}

func _ddfg(_fcca int, _ggcag int) error {
	return _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", _fcca, _ggcag)
}
func (_affb *CMYK32) ColorModel() _ag.Model { return _ag.CMYKModel }
func (_gbbe *Gray16) Validate() error {
	if len(_gbbe.Data) != _gbbe.Height*_gbbe.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_ccab *Gray2) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _ccab.Width, Y: _ccab.Height}}
}

type RasterOperator int

func (_aaec *Monochrome) At(x, y int) _ag.Color { _dbcg, _ := _aaec.ColorAt(x, y); return _dbcg }
func (_begef *NRGBA32) ColorModel() _ag.Model   { return _ag.NRGBAModel }

type ColorConverter interface {
	Convert(_eff _c.Image) (Image, error)
}

func _da(_dc *Monochrome, _bgc int, _bb []uint) (*Monochrome, error) {
	_bgd := _bgc * _dc.Width
	_ce := _bgc * _dc.Height
	_fb := _gea(_bgd, _ce)
	for _fdb, _df := range _bb {
		var _gg error
		switch _df {
		case 2:
			_gg = _fbd(_fb, _dc)
		case 4:
			_gg = _ee(_fb, _dc)
		case 8:
			_gg = _gge(_fb, _dc)
		}
		if _gg != nil {
			return nil, _gg
		}
		if _fdb != len(_bb)-1 {
			_dc = _fb.copy()
		}
	}
	return _fb, nil
}
func (_gdgc *NRGBA16) setNRGBA(_efgcc, _edff, _dceea int, _edfac _ag.NRGBA) {
	if _efgcc*3%2 == 0 {
		_gdgc.Data[_dceea] = (_edfac.R>>4)<<4 | (_edfac.G >> 4)
		_gdgc.Data[_dceea+1] = (_edfac.B>>4)<<4 | (_gdgc.Data[_dceea+1] & 0xf)
	} else {
		_gdgc.Data[_dceea] = (_gdgc.Data[_dceea] & 0xf0) | (_edfac.R >> 4)
		_gdgc.Data[_dceea+1] = (_edfac.G>>4)<<4 | (_edfac.B >> 4)
	}
	if _gdgc.Alpha != nil {
		_edbb := _edff * BytesPerLine(_gdgc.Width, 4, 1)
		if _edbb < len(_gdgc.Alpha) {
			if _efgcc%2 == 0 {
				_gdgc.Alpha[_edbb] = (_edfac.A>>uint(4))<<uint(4) | (_gdgc.Alpha[_dceea] & 0xf)
			} else {
				_gdgc.Alpha[_edbb] = (_gdgc.Alpha[_edbb] & 0xf0) | (_edfac.A >> uint(4))
			}
		}
	}
}
func _ecdb(_efdf *Monochrome, _aeda, _eccf, _bdba, _cbde int, _fegd RasterOperator) {
	if _aeda < 0 {
		_bdba += _aeda
		_aeda = 0
	}
	_ccbc := _aeda + _bdba - _efdf.Width
	if _ccbc > 0 {
		_bdba -= _ccbc
	}
	if _eccf < 0 {
		_cbde += _eccf
		_eccf = 0
	}
	_ebcda := _eccf + _cbde - _efdf.Height
	if _ebcda > 0 {
		_cbde -= _ebcda
	}
	if _bdba <= 0 || _cbde <= 0 {
		return
	}
	if (_aeda & 7) == 0 {
		_edfgf(_efdf, _aeda, _eccf, _bdba, _cbde, _fegd)
	} else {
		_gfdf(_efdf, _aeda, _eccf, _bdba, _cbde, _fegd)
	}
}
func (_eagff *ImageBase) getByte(_gaaf int) (byte, error) {
	if _gaaf > len(_eagff.Data)-1 || _gaaf < 0 {
		return 0, _ea.Errorf("\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", _gaaf)
	}
	return _eagff.Data[_gaaf], nil
}
func _aeeb(_afab NRGBA, _cdda RGBA, _cdcg _c.Rectangle) {
	for _faaa := 0; _faaa < _cdcg.Max.X; _faaa++ {
		for _egdc := 0; _egdc < _cdcg.Max.Y; _egdc++ {
			_edac := _afab.NRGBAAt(_faaa, _egdc)
			_cdda.SetRGBA(_faaa, _egdc, _ede(_edac))
		}
	}
}
func ColorAtCMYK(x, y, width int, data []byte, decode []float64) (_ag.CMYK, error) {
	_ccd := 4 * (y*width + x)
	if _ccd+3 >= len(data) {
		return _ag.CMYK{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	C := data[_ccd] & 0xff
	M := data[_ccd+1] & 0xff
	Y := data[_ccd+2] & 0xff
	K := data[_ccd+3] & 0xff
	if len(decode) == 8 {
		C = uint8(uint32(LinearInterpolate(float64(C), 0, 255, decode[0], decode[1])) & 0xff)
		M = uint8(uint32(LinearInterpolate(float64(M), 0, 255, decode[2], decode[3])) & 0xff)
		Y = uint8(uint32(LinearInterpolate(float64(Y), 0, 255, decode[4], decode[5])) & 0xff)
		K = uint8(uint32(LinearInterpolate(float64(K), 0, 255, decode[6], decode[7])) & 0xff)
	}
	return _ag.CMYK{C: C, M: M, Y: Y, K: K}, nil
}

var (
	_bcdf = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_face = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func _gcagc(_afdc _c.Image, _defg Image, _dgadd _c.Rectangle) {
	if _cdc, _bdde := _afdc.(SMasker); _bdde && _cdc.HasAlpha() {
		_defg.(SMasker).MakeAlpha()
	}
	_eceg(_afdc, _defg, _dgadd)
}
func (_aaegf *Gray16) Histogram() (_bfc [256]int) {
	for _fad := 0; _fad < _aaegf.Width; _fad++ {
		for _gbge := 0; _gbge < _aaegf.Height; _gbge++ {
			_bfc[_aaegf.GrayAt(_fad, _gbge).Y]++
		}
	}
	return _bfc
}
func (_dae *Gray2) Copy() Image { return &Gray2{ImageBase: _dae.copy()} }
func _adce(_afcf _ag.NRGBA64) _ag.RGBA {
	_gfbc, _faa, _ebag, _cebb := _afcf.RGBA()
	return _ag.RGBA{R: uint8(_gfbc >> 8), G: uint8(_faa >> 8), B: uint8(_ebag >> 8), A: uint8(_cebb >> 8)}
}
func (_ebf *CMYK32) Copy() Image { return &CMYK32{ImageBase: _ebf.copy()} }

var _ Image = &RGBA32{}

func ColorAtGray8BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_ag.Gray, error) {
	_ccec := y*bytesPerLine + x
	if _ccec >= len(data) {
		return _ag.Gray{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_eddgb := data[_ccec]
	if len(decode) == 2 {
		_eddgb = uint8(uint32(LinearInterpolate(float64(_eddgb), 0, 255, decode[0], decode[1])) & 0xff)
	}
	return _ag.Gray{Y: _eddgb}, nil
}
func (_dabd *Monochrome) copy() *Monochrome {
	_ccda := _gea(_dabd.Width, _dabd.Height)
	_ccda.ModelThreshold = _dabd.ModelThreshold
	_ccda.Data = make([]byte, len(_dabd.Data))
	copy(_ccda.Data, _dabd.Data)
	if len(_dabd.Decode) != 0 {
		_ccda.Decode = make([]float64, len(_dabd.Decode))
		copy(_ccda.Decode, _dabd.Decode)
	}
	if len(_dabd.Alpha) != 0 {
		_ccda.Alpha = make([]byte, len(_dabd.Alpha))
		copy(_ccda.Alpha, _dabd.Alpha)
	}
	return _ccda
}
func (_egddd *CMYK32) Set(x, y int, c _ag.Color) {
	_bbb := 4 * (y*_egddd.Width + x)
	if _bbb+3 >= len(_egddd.Data) {
		return
	}
	_bda := _ag.CMYKModel.Convert(c).(_ag.CMYK)
	_egddd.Data[_bbb] = _bda.C
	_egddd.Data[_bbb+1] = _bda.M
	_egddd.Data[_bbb+2] = _bda.Y
	_egddd.Data[_bbb+3] = _bda.K
}
func _gfg() (_aef [256]uint16) {
	for _ecf := 0; _ecf < 256; _ecf++ {
		if _ecf&0x01 != 0 {
			_aef[_ecf] |= 0x3
		}
		if _ecf&0x02 != 0 {
			_aef[_ecf] |= 0xc
		}
		if _ecf&0x04 != 0 {
			_aef[_ecf] |= 0x30
		}
		if _ecf&0x08 != 0 {
			_aef[_ecf] |= 0xc0
		}
		if _ecf&0x10 != 0 {
			_aef[_ecf] |= 0x300
		}
		if _ecf&0x20 != 0 {
			_aef[_ecf] |= 0xc00
		}
		if _ecf&0x40 != 0 {
			_aef[_ecf] |= 0x3000
		}
		if _ecf&0x80 != 0 {
			_aef[_ecf] |= 0xc000
		}
	}
	return _aef
}
func _gcdb(_affc _c.Image) (Image, error) {
	if _beb, _ege := _affc.(*CMYK32); _ege {
		return _beb.Copy(), nil
	}
	_dfba := _affc.Bounds()
	_dfed, _cebc := NewImage(_dfba.Max.X, _dfba.Max.Y, 8, 4, nil, nil, nil)
	if _cebc != nil {
		return nil, _cebc
	}
	switch _ffd := _affc.(type) {
	case CMYK:
		_ffg(_ffd, _dfed.(CMYK), _dfba)
	case Gray:
		_caab(_ffd, _dfed.(CMYK), _dfba)
	case NRGBA:
		_bbf(_ffd, _dfed.(CMYK), _dfba)
	case RGBA:
		_gfce(_ffd, _dfed.(CMYK), _dfba)
	default:
		_eceg(_affc, _dfed, _dfba)
	}
	return _dfed, nil
}
func (_edffe *NRGBA64) ColorModel() _ag.Model { return _ag.NRGBA64Model }
func (_gaef *NRGBA16) At(x, y int) _ag.Color {
	_ffgc, _ := _gaef.ColorAt(x, y)
	return _ffgc
}
func _dggc(_aeaeg Gray, _fecdg NRGBA, _dead _c.Rectangle) {
	for _geefe := 0; _geefe < _dead.Max.X; _geefe++ {
		for _cdag := 0; _cdag < _dead.Max.Y; _cdag++ {
			_cgaa := _aeaeg.GrayAt(_geefe, _cdag)
			_fecdg.SetNRGBA(_geefe, _cdag, _abgg(_cgaa))
		}
	}
}

type colorConverter struct {
	_dcc func(_gdf _c.Image) (Image, error)
}

func _dce(_eee _c.Image) (Image, error) {
	if _gbbb, _afga := _eee.(*Monochrome); _afga {
		return _gbbb, nil
	}
	_abc := _eee.Bounds()
	var _fefb Gray
	switch _efaa := _eee.(type) {
	case Gray:
		_fefb = _efaa
	case NRGBA:
		_fefb = &Gray8{ImageBase: NewImageBase(_abc.Max.X, _abc.Max.Y, 8, 1, nil, nil, nil)}
		_cgee(_fefb, _efaa, _abc)
	case nrgba64:
		_fefb = &Gray8{ImageBase: NewImageBase(_abc.Max.X, _abc.Max.Y, 8, 1, nil, nil, nil)}
		_ccb(_fefb, _efaa, _abc)
	default:
		_fbde, _gabf := GrayConverter.Convert(_eee)
		if _gabf != nil {
			return nil, _gabf
		}
		_fefb = _fbde.(Gray)
	}
	_bcdb, _dgc := NewImage(_abc.Max.X, _abc.Max.Y, 1, 1, nil, nil, nil)
	if _dgc != nil {
		return nil, _dgc
	}
	_gef := _bcdb.(*Monochrome)
	_ccdc := AutoThresholdTriangle(GrayHistogram(_fefb))
	for _gcb := 0; _gcb < _abc.Max.X; _gcb++ {
		for _bcdd := 0; _bcdd < _abc.Max.Y; _bcdd++ {
			_gcaf := _dagc(_fefb.GrayAt(_gcb, _bcdd), monochromeModel(_ccdc))
			_gef.SetGray(_gcb, _bcdd, _gcaf)
		}
	}
	return _bcdb, nil
}
func _caab(_fce Gray, _fdc CMYK, _cgeb _c.Rectangle) {
	for _gfba := 0; _gfba < _cgeb.Max.X; _gfba++ {
		for _caee := 0; _caee < _cgeb.Max.Y; _caee++ {
			_bccb := _fce.GrayAt(_gfba, _caee)
			_fdc.SetCMYK(_gfba, _caee, _aabe(_bccb))
		}
	}
}
func _gba(_ecgd, _fg *Monochrome, _eecd []byte, _cef int) (_eed error) {
	var (
		_gac, _gff, _gbb, _ceeb, _abg, _acg, _gfa, _ega int
		_gag, _ebef                                     uint32
		_bgcg, _fbf                                     byte
		_edba                                           uint16
	)
	_aff := make([]byte, 4)
	_fec := make([]byte, 4)
	for _gbb = 0; _gbb < _ecgd.Height-1; _gbb, _ceeb = _gbb+2, _ceeb+1 {
		_gac = _gbb * _ecgd.BytesPerLine
		_gff = _ceeb * _fg.BytesPerLine
		for _abg, _acg = 0, 0; _abg < _cef; _abg, _acg = _abg+4, _acg+1 {
			for _gfa = 0; _gfa < 4; _gfa++ {
				_ega = _gac + _abg + _gfa
				if _ega <= len(_ecgd.Data)-1 && _ega < _gac+_ecgd.BytesPerLine {
					_aff[_gfa] = _ecgd.Data[_ega]
				} else {
					_aff[_gfa] = 0x00
				}
				_ega = _gac + _ecgd.BytesPerLine + _abg + _gfa
				if _ega <= len(_ecgd.Data)-1 && _ega < _gac+(2*_ecgd.BytesPerLine) {
					_fec[_gfa] = _ecgd.Data[_ega]
				} else {
					_fec[_gfa] = 0x00
				}
			}
			_gag = _be.BigEndian.Uint32(_aff)
			_ebef = _be.BigEndian.Uint32(_fec)
			_ebef |= _gag
			_ebef |= _ebef << 1
			_ebef &= 0xaaaaaaaa
			_gag = _ebef | (_ebef << 7)
			_bgcg = byte(_gag >> 24)
			_fbf = byte((_gag >> 8) & 0xff)
			_ega = _gff + _acg
			if _ega+1 == len(_fg.Data)-1 || _ega+1 >= _gff+_fg.BytesPerLine {
				_fg.Data[_ega] = _eecd[_bgcg]
			} else {
				_edba = (uint16(_eecd[_bgcg]) << 8) | uint16(_eecd[_fbf])
				if _eed = _fg.setTwoBytes(_ega, _edba); _eed != nil {
					return _ea.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _ega)
				}
				_acg++
			}
		}
	}
	return nil
}
func (_feb *Gray8) Histogram() (_gcfd [256]int) {
	for _ebcd := 0; _ebcd < len(_feb.Data); _ebcd++ {
		_gcfd[_feb.Data[_ebcd]]++
	}
	return _gcfd
}
func (_cbeed *Monochrome) getBitAt(_fgbe, _aabg int) bool {
	_eeda := _aabg*_cbeed.BytesPerLine + (_fgbe >> 3)
	_deaf := _fgbe & 0x07
	_gbefd := uint(7 - _deaf)
	if _eeda > len(_cbeed.Data)-1 {
		return false
	}
	if (_cbeed.Data[_eeda]>>_gbefd)&0x01 >= 1 {
		return true
	}
	return false
}
func _gcga(_ceea _ag.RGBA) _ag.NRGBA {
	switch _ceea.A {
	case 0xff:
		return _ag.NRGBA{R: _ceea.R, G: _ceea.G, B: _ceea.B, A: 0xff}
	case 0x00:
		return _ag.NRGBA{}
	default:
		_aec, _dfa, _eda, _dccb := _ceea.RGBA()
		_aec = (_aec * 0xffff) / _dccb
		_dfa = (_dfa * 0xffff) / _dccb
		_eda = (_eda * 0xffff) / _dccb
		return _ag.NRGBA{R: uint8(_aec >> 8), G: uint8(_dfa >> 8), B: uint8(_eda >> 8), A: uint8(_dccb >> 8)}
	}
}
func _fedg(_age *_c.NYCbCrA, _eabd RGBA, _cde _c.Rectangle) {
	for _eecg := 0; _eecg < _cde.Max.X; _eecg++ {
		for _cfgb := 0; _cfgb < _cde.Max.Y; _cfgb++ {
			_fcfd := _age.NYCbCrAAt(_eecg, _cfgb)
			_eabd.SetRGBA(_eecg, _cfgb, _agc(_fcfd))
		}
	}
}
func (_adbe *NRGBA16) Set(x, y int, c _ag.Color) {
	_fdag := y*_adbe.BytesPerLine + x*3/2
	if _fdag+1 >= len(_adbe.Data) {
		return
	}
	_dfeb := NRGBA16Model.Convert(c).(_ag.NRGBA)
	_adbe.setNRGBA(x, y, _fdag, _dfeb)
}
func ColorAtGray4BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_ag.Gray, error) {
	_baff := y*bytesPerLine + x>>1
	if _baff >= len(data) {
		return _ag.Gray{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_bfac := data[_baff] >> uint(4-(x&1)*4) & 0xf
	if len(decode) == 2 {
		_bfac = uint8(uint32(LinearInterpolate(float64(_bfac), 0, 15, decode[0], decode[1])) & 0xf)
	}
	return _ag.Gray{Y: _bfac * 17 & 0xff}, nil
}
func _agc(_efg _ag.NYCbCrA) _ag.RGBA {
	_efdg, _edbag, _fbbg, _cdgd := _dga(_efg).RGBA()
	return _ag.RGBA{R: uint8(_efdg >> 8), G: uint8(_edbag >> 8), B: uint8(_fbbg >> 8), A: uint8(_cdgd >> 8)}
}
func (_dda *CMYK32) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _dda.Width, Y: _dda.Height}}
}
func _aab(_afa _ag.Gray) _ag.RGBA { return _ag.RGBA{R: _afa.Y, G: _afa.Y, B: _afa.Y, A: 0xff} }
func _ffg(_efa, _ece CMYK, _ddad _c.Rectangle) {
	for _eaf := 0; _eaf < _ddad.Max.X; _eaf++ {
		for _ded := 0; _ded < _ddad.Max.Y; _ded++ {
			_ece.SetCMYK(_eaf, _ded, _efa.CMYKAt(_eaf, _ded))
		}
	}
}
func (_dgbaa *Monochrome) ReduceBinary(factor float64) (*Monochrome, error) {
	_fea := _gefd(uint(factor))
	if !IsPowerOf2(uint(factor)) {
		_fea++
	}
	_dcfd := make([]int, _fea)
	for _dcac := range _dcfd {
		_dcfd[_dcac] = 4
	}
	_dcb, _eegcc := _beed(_dgbaa, _dcfd...)
	if _eegcc != nil {
		return nil, _eegcc
	}
	return _dcb, nil
}
func (_bgaa *Monochrome) setBit(_bgad, _cac int) { _bgaa.Data[_bgad+(_cac>>3)] |= 0x80 >> uint(_cac&7) }
func _ffed(_bbg _ag.CMYK) _ag.Gray {
	_gebg, _gde, _agg := _ag.CMYKToRGB(_bbg.C, _bbg.M, _bbg.Y, _bbg.K)
	_adcf := (19595*uint32(_gebg) + 38470*uint32(_gde) + 7471*uint32(_agg) + 1<<7) >> 16
	return _ag.Gray{Y: uint8(_adcf)}
}

var _ Image = &Gray8{}

func (_dfd *Gray4) setGray(_cddb int, _eadf int, _eadc _ag.Gray) {
	_aea := _eadf * _dfd.BytesPerLine
	_egde := _aea + (_cddb >> 1)
	if _egde >= len(_dfd.Data) {
		return
	}
	_egb := _eadc.Y >> 4
	_dfd.Data[_egde] = (_dfd.Data[_egde] & (^(0xf0 >> uint(4*(_cddb&1))))) | (_egb << uint(4-4*(_cddb&1)))
}
func _beab(_eacc NRGBA, _bcb Gray, _gfbce _c.Rectangle) {
	for _dec := 0; _dec < _gfbce.Max.X; _dec++ {
		for _ebbc := 0; _ebbc < _gfbce.Max.Y; _ebbc++ {
			_fegc := _gbgf(_eacc.NRGBAAt(_dec, _ebbc))
			_bcb.SetGray(_dec, _ebbc, _fegc)
		}
	}
}
func _ddc(_eaag, _dbdc, _cead byte) byte { return (_eaag &^ (_cead)) | (_dbdc & _cead) }
func _ddga(_gefbd *_c.NYCbCrA, _aagf NRGBA, _faeb _c.Rectangle) {
	for _baca := 0; _baca < _faeb.Max.X; _baca++ {
		for _caca := 0; _caca < _faeb.Max.Y; _caca++ {
			_dcfg := _gefbd.NYCbCrAAt(_baca, _caca)
			_aagf.SetNRGBA(_baca, _caca, _dga(_dcfg))
		}
	}
}
func _cgee(_geef Gray, _fda NRGBA, _dcda _c.Rectangle) {
	for _afca := 0; _afca < _dcda.Max.X; _afca++ {
		for _bggc := 0; _bggc < _dcda.Max.Y; _bggc++ {
			_cfgc := _aaaa(_fda.NRGBAAt(_afca, _bggc))
			_geef.SetGray(_afca, _bggc, _cfgc)
		}
	}
}
func (_faab *Gray4) SetGray(x, y int, g _ag.Gray) {
	if x >= _faab.Width || y >= _faab.Height {
		return
	}
	g = _gagf(g)
	_faab.setGray(x, y, g)
}

var (
	Gray2Model   = _ag.ModelFunc(_cfbg)
	Gray4Model   = _ag.ModelFunc(_eccc)
	NRGBA16Model = _ag.ModelFunc(_gfee)
)

func (_eegb *Gray2) Histogram() (_eeaa [256]int) {
	for _eege := 0; _eege < _eegb.Width; _eege++ {
		for _ebb := 0; _ebb < _eegb.Height; _ebb++ {
			_eeaa[_eegb.GrayAt(_eege, _ebb).Y]++
		}
	}
	return _eeaa
}
func init() { _aeag() }
func ConverterFunc(converterFunc func(_dfbag _c.Image) (Image, error)) ColorConverter {
	return colorConverter{_dcc: converterFunc}
}
func ColorAtNRGBA16(x, y, width, bytesPerLine int, data, alpha []byte, decode []float64) (_ag.NRGBA, error) {
	_dgaa := y*bytesPerLine + x*3/2
	if _dgaa+1 >= len(data) {
		return _ag.NRGBA{}, _ddfg(x, y)
	}
	const (
		_bbgg  = 0xf
		_fcgce = uint8(0xff)
	)
	_fdcd := _fcgce
	if alpha != nil {
		_gfda := y * BytesPerLine(width, 4, 1)
		if _gfda < len(alpha) {
			if x%2 == 0 {
				_fdcd = (alpha[_gfda] >> uint(4)) & _bbgg
			} else {
				_fdcd = alpha[_gfda] & _bbgg
			}
			_fdcd |= _fdcd << 4
		}
	}
	var _bfce, _cfgcb, _cbea uint8
	if x*3%2 == 0 {
		_bfce = (data[_dgaa] >> uint(4)) & _bbgg
		_cfgcb = data[_dgaa] & _bbgg
		_cbea = (data[_dgaa+1] >> uint(4)) & _bbgg
	} else {
		_bfce = data[_dgaa] & _bbgg
		_cfgcb = (data[_dgaa+1] >> uint(4)) & _bbgg
		_cbea = data[_dgaa+1] & _bbgg
	}
	if len(decode) == 6 {
		_bfce = uint8(uint32(LinearInterpolate(float64(_bfce), 0, 15, decode[0], decode[1])) & 0xf)
		_cfgcb = uint8(uint32(LinearInterpolate(float64(_cfgcb), 0, 15, decode[2], decode[3])) & 0xf)
		_cbea = uint8(uint32(LinearInterpolate(float64(_cbea), 0, 15, decode[4], decode[5])) & 0xf)
	}
	return _ag.NRGBA{R: (_bfce << 4) | (_bfce & 0xf), G: (_cfgcb << 4) | (_cfgcb & 0xf), B: (_cbea << 4) | (_cbea & 0xf), A: _fdcd}, nil
}

type NRGBA interface {
	NRGBAAt(_gccbf, _bbbdc int) _ag.NRGBA
	SetNRGBA(_bdbb, _ggg int, _efbe _ag.NRGBA)
}

var _ Image = &NRGBA32{}

func (_bgfbec *NRGBA32) Copy() Image { return &NRGBA32{ImageBase: _bgfbec.copy()} }
func (_abdb *ImageBase) MakeAlpha()  { _abdb.newAlpha() }
func _feef(_edga _c.Image) (Image, error) {
	if _bcac, _dfc := _edga.(*Gray2); _dfc {
		return _bcac.Copy(), nil
	}
	_deb := _edga.Bounds()
	_bfgb, _ffdb := NewImage(_deb.Max.X, _deb.Max.Y, 2, 1, nil, nil, nil)
	if _ffdb != nil {
		return nil, _ffdb
	}
	_ffdf(_edga, _bfgb, _deb)
	return _bfgb, nil
}

type ImageBase struct {
	Width, Height                     int
	BitsPerComponent, ColorComponents int
	Data, Alpha                       []byte
	Decode                            []float64
	BytesPerLine                      int
}

func (_dgae monochromeModel) Convert(c _ag.Color) _ag.Color {
	_cgf := _ag.GrayModel.Convert(c).(_ag.Gray)
	return _dagc(_cgf, _dgae)
}
func _ggf(_caec _ag.NRGBA64) _ag.Gray {
	var _aba _ag.NRGBA64
	if _caec == _aba {
		return _ag.Gray{Y: 0xff}
	}
	_cbe, _dabg, _bcce, _ := _caec.RGBA()
	_acc := (19595*_cbe + 38470*_dabg + 7471*_bcce + 1<<15) >> 24
	return _ag.Gray{Y: uint8(_acc)}
}
func (_gec *Monochrome) clearBit(_aeee, _bdc int) { _gec.Data[_aeee] &= ^(0x80 >> uint(_bdc&7)) }

type NRGBA64 struct{ ImageBase }

func (_ggge *NRGBA64) Base() *ImageBase { return &_ggge.ImageBase }
func (_cfga *Gray8) At(x, y int) _ag.Color {
	_efaad, _ := _cfga.ColorAt(x, y)
	return _efaad
}
func _aeea(_gfge _ag.NRGBA) _ag.NRGBA {
	_gfge.R = _gfge.R>>4 | (_gfge.R>>4)<<4
	_gfge.G = _gfge.G>>4 | (_gfge.G>>4)<<4
	_gfge.B = _gfge.B>>4 | (_gfge.B>>4)<<4
	return _gfge
}
func (_dabb *Monochrome) InverseData() error {
	return _dabb.RasterOperation(0, 0, _dabb.Width, _dabb.Height, PixNotDst, nil, 0, 0)
}
func MonochromeModel(threshold uint8) _ag.Model { return monochromeModel(threshold) }
func (_acff *Monochrome) SetGray(x, y int, g _ag.Gray) {
	_eagb := y*_acff.BytesPerLine + x>>3
	if _eagb > len(_acff.Data)-1 {
		return
	}
	g = _dagc(g, monochromeModel(_acff.ModelThreshold))
	_acff.setGray(x, g, _eagb)
}
func _baeg(_dcff _c.Image) (Image, error) {
	if _bgfe, _caeb := _dcff.(*Gray16); _caeb {
		return _bgfe.Copy(), nil
	}
	_eagfg := _dcff.Bounds()
	_gfac, _aacde := NewImage(_eagfg.Max.X, _eagfg.Max.Y, 16, 1, nil, nil, nil)
	if _aacde != nil {
		return nil, _aacde
	}
	_ffdf(_dcff, _gfac, _eagfg)
	return _gfac, nil
}
func (_deg *Gray8) Copy() Image { return &Gray8{ImageBase: _deg.copy()} }
func (_dbbc *Monochrome) ResolveDecode() error {
	if len(_dbbc.Decode) != 2 {
		return nil
	}
	if _dbbc.Decode[0] == 1 && _dbbc.Decode[1] == 0 {
		if _gacd := _dbbc.InverseData(); _gacd != nil {
			return _gacd
		}
		_dbbc.Decode = nil
	}
	return nil
}
func (_ccbb *NRGBA64) NRGBA64At(x, y int) _ag.NRGBA64 {
	_cdff, _ := ColorAtNRGBA64(x, y, _ccbb.Width, _ccbb.Data, _ccbb.Alpha, _ccbb.Decode)
	return _cdff
}
func AutoThresholdTriangle(histogram [256]int) uint8 {
	var _egee, _fdfef, _beeb, _fdge int
	for _fggb := 0; _fggb < len(histogram); _fggb++ {
		if histogram[_fggb] > 0 {
			_egee = _fggb
			break
		}
	}
	if _egee > 0 {
		_egee--
	}
	for _gfbag := 255; _gfbag > 0; _gfbag-- {
		if histogram[_gfbag] > 0 {
			_fdge = _gfbag
			break
		}
	}
	if _fdge < 255 {
		_fdge++
	}
	for _fdcf := 0; _fdcf < 256; _fdcf++ {
		if histogram[_fdcf] > _fdfef {
			_beeb = _fdcf
			_fdfef = histogram[_fdcf]
		}
	}
	var _acgd bool
	if (_beeb - _egee) < (_fdge - _beeb) {
		_acgd = true
		var _fbeg int
		_aggf := 255
		for _fbeg < _aggf {
			_fgbg := histogram[_fbeg]
			histogram[_fbeg] = histogram[_aggf]
			histogram[_aggf] = _fgbg
			_fbeg++
			_aggf--
		}
		_egee = 255 - _fdge
		_beeb = 255 - _beeb
	}
	if _egee == _beeb {
		return uint8(_egee)
	}
	_ddgd := float64(histogram[_beeb])
	_cfdg := float64(_egee - _beeb)
	_fdea := _e.Sqrt(_ddgd*_ddgd + _cfdg*_cfdg)
	_ddgd /= _fdea
	_cfdg /= _fdea
	_fdea = _ddgd*float64(_egee) + _cfdg*float64(histogram[_egee])
	_ddec := _egee
	var _gaagc float64
	for _ggaf := _egee + 1; _ggaf <= _beeb; _ggaf++ {
		_ffea := _ddgd*float64(_ggaf) + _cfdg*float64(histogram[_ggaf]) - _fdea
		if _ffea > _gaagc {
			_ddec = _ggaf
			_gaagc = _ffea
		}
	}
	_ddec--
	if _acgd {
		var _fcfa int
		_gdfg := 255
		for _fcfa < _gdfg {
			_begg := histogram[_fcfa]
			histogram[_fcfa] = histogram[_gdfg]
			histogram[_gdfg] = _begg
			_fcfa++
			_gdfg--
		}
		return uint8(255 - _ddec)
	}
	return uint8(_ddec)
}

type monochromeThresholdConverter struct{ Threshold uint8 }

func (_eegcag *NRGBA16) Base() *ImageBase { return &_eegcag.ImageBase }
func _bcdc(_fadg CMYK, _babf Gray, _deae _c.Rectangle) {
	for _dedc := 0; _dedc < _deae.Max.X; _dedc++ {
		for _geede := 0; _geede < _deae.Max.Y; _geede++ {
			_gdac := _ffed(_fadg.CMYKAt(_dedc, _geede))
			_babf.SetGray(_dedc, _geede, _gdac)
		}
	}
}
func (_eaae *Monochrome) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _bgge(_eaae, dx, dy, dw, dh, op, src, sx, sy)
}
func (_gbeg *CMYK32) At(x, y int) _ag.Color { _adc, _ := _gbeg.ColorAt(x, y); return _adc }
func ColorAtNRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_ag.NRGBA, error) {
	_cceb := y*width + x
	_aeba := 3 * _cceb
	if _aeba+2 >= len(data) {
		return _ag.NRGBA{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_dbabe := uint8(0xff)
	if alpha != nil && len(alpha) > _cceb {
		_dbabe = alpha[_cceb]
	}
	_dbeg, _egac, _gbcg := data[_aeba], data[_aeba+1], data[_aeba+2]
	if len(decode) == 6 {
		_dbeg = uint8(uint32(LinearInterpolate(float64(_dbeg), 0, 255, decode[0], decode[1])) & 0xff)
		_egac = uint8(uint32(LinearInterpolate(float64(_egac), 0, 255, decode[2], decode[3])) & 0xff)
		_gbcg = uint8(uint32(LinearInterpolate(float64(_gbcg), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _ag.NRGBA{R: _dbeg, G: _egac, B: _gbcg, A: _dbabe}, nil
}
func _ede(_ggeg _ag.NRGBA) _ag.RGBA {
	_gcc, _edda, _acfc, _ecege := _ggeg.RGBA()
	return _ag.RGBA{R: uint8(_gcc >> 8), G: uint8(_edda >> 8), B: uint8(_acfc >> 8), A: uint8(_ecege >> 8)}
}
func _bbcb(_ggfg uint8) bool {
	if _ggfg == 0 || _ggfg == 255 {
		return true
	}
	return false
}
func AddDataPadding(width, height, bitsPerComponent, colorComponents int, data []byte) ([]byte, error) {
	_dfea := BytesPerLine(width, bitsPerComponent, colorComponents)
	if _dfea == width*colorComponents*bitsPerComponent/8 {
		return data, nil
	}
	_fbda := width * colorComponents * bitsPerComponent
	_fbea := _dfea * 8
	_gaag := 8 - (_fbea - _fbda)
	_eddag := _d.NewReader(data)
	_dbcf := _dfea - 1
	_bccc := make([]byte, _dbcf)
	_cfc := make([]byte, height*_dfea)
	_aeae := _d.NewWriterMSB(_cfc)
	var _ecde uint64
	var _babb error
	for _eddae := 0; _eddae < height; _eddae++ {
		_, _babb = _eddag.Read(_bccc)
		if _babb != nil {
			return nil, _babb
		}
		_, _babb = _aeae.Write(_bccc)
		if _babb != nil {
			return nil, _babb
		}
		_ecde, _babb = _eddag.ReadBits(byte(_gaag))
		if _babb != nil {
			return nil, _babb
		}
		_, _babb = _aeae.WriteBits(_ecde, _gaag)
		if _babb != nil {
			return nil, _babb
		}
		_aeae.FinishByte()
	}
	return _cfc, nil
}
func ScaleAlphaToMonochrome(data []byte, width, height int) ([]byte, error) {
	_bg := BytesPerLine(width, 8, 1)
	if len(data) < _bg*height {
		return nil, nil
	}
	_bc := &Gray8{NewImageBase(width, height, 8, 1, data, nil, nil)}
	_cd, _ef := MonochromeConverter.Convert(_bc)
	if _ef != nil {
		return nil, _ef
	}
	return _cd.Base().Data, nil
}
func _aedd(_bce _c.Image) (Image, error) {
	if _ffcg, _fcab := _bce.(*RGBA32); _fcab {
		return _ffcg.Copy(), nil
	}
	_ffcdb, _bcge, _bdbe := _bdfe(_bce, 1)
	_gbcd := &RGBA32{ImageBase: NewImageBase(_ffcdb.Max.X, _ffcdb.Max.Y, 8, 3, nil, _bdbe, nil)}
	_debg(_bce, _gbcd, _ffcdb)
	if len(_bdbe) != 0 && !_bcge {
		if _eadff := _aaff(_bdbe, _gbcd); _eadff != nil {
			return nil, _eadff
		}
	}
	return _gbcd, nil
}
func (_gcca *Gray8) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtGray8BPC(x, y, _gcca.BytesPerLine, _gcca.Data, _gcca.Decode)
}

var (
	MonochromeConverter = ConverterFunc(_dce)
	Gray2Converter      = ConverterFunc(_feef)
	Gray4Converter      = ConverterFunc(_fdcb)
	GrayConverter       = ConverterFunc(_ceae)
	Gray16Converter     = ConverterFunc(_baeg)
	NRGBA16Converter    = ConverterFunc(_ecfe)
	NRGBAConverter      = ConverterFunc(_bac)
	NRGBA64Converter    = ConverterFunc(_gcff)
	RGBAConverter       = ConverterFunc(_aedd)
	CMYKConverter       = ConverterFunc(_gcdb)
)

func ColorAtGray2BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_ag.Gray, error) {
	_ggecf := y*bytesPerLine + x>>2
	if _ggecf >= len(data) {
		return _ag.Gray{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_cbdb := data[_ggecf] >> uint(6-(x&3)*2) & 3
	if len(decode) == 2 {
		_cbdb = uint8(uint32(LinearInterpolate(float64(_cbdb), 0, 3.0, decode[0], decode[1])) & 3)
	}
	return _ag.Gray{Y: _cbdb * 85}, nil
}
func (_gfe *Gray4) Base() *ImageBase { return &_gfe.ImageBase }
func _aabe(_cdd _ag.Gray) _ag.CMYK   { return _ag.CMYK{K: 0xff - _cdd.Y} }
func (_aeeg *NRGBA16) Copy() Image   { return &NRGBA16{ImageBase: _aeeg.copy()} }
func (_gcfda *ImageBase) setEightPartlyBytes(_dfag, _bege int, _dgcbg uint64) (_gdad error) {
	var (
		_ged  byte
		_gagg int
	)
	for _caebg := 1; _caebg <= _bege; _caebg++ {
		_gagg = 64 - _caebg*8
		_ged = byte(_dgcbg >> uint(_gagg) & 0xff)
		if _gdad = _gcfda.setByte(_dfag+_caebg-1, _ged); _gdad != nil {
			return _gdad
		}
	}
	_dbca := _gcfda.BytesPerLine*8 - _gcfda.Width
	if _dbca == 0 {
		return nil
	}
	_gagg -= 8
	_ged = byte(_dgcbg>>uint(_gagg)&0xff) << uint(_dbca)
	if _gdad = _gcfda.setByte(_dfag+_bege, _ged); _gdad != nil {
		return _gdad
	}
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

func _edfgf(_bbbd *Monochrome, _dcffb, _ebdeg int, _ggcb, _cfeba int, _cgeeg RasterOperator) {
	var (
		_affa          int
		_eaad          byte
		_fcebb, _egaba int
		_affe          int
	)
	_gaaa := _ggcb >> 3
	_geca := _ggcb & 7
	if _geca > 0 {
		_eaad = _bcdf[_geca]
	}
	_affa = _bbbd.BytesPerLine*_ebdeg + (_dcffb >> 3)
	switch _cgeeg {
	case PixClr:
		for _fcebb = 0; _fcebb < _cfeba; _fcebb++ {
			_affe = _affa + _fcebb*_bbbd.BytesPerLine
			for _egaba = 0; _egaba < _gaaa; _egaba++ {
				_bbbd.Data[_affe] = 0x0
				_affe++
			}
			if _geca > 0 {
				_bbbd.Data[_affe] = _ddc(_bbbd.Data[_affe], 0x0, _eaad)
			}
		}
	case PixSet:
		for _fcebb = 0; _fcebb < _cfeba; _fcebb++ {
			_affe = _affa + _fcebb*_bbbd.BytesPerLine
			for _egaba = 0; _egaba < _gaaa; _egaba++ {
				_bbbd.Data[_affe] = 0xff
				_affe++
			}
			if _geca > 0 {
				_bbbd.Data[_affe] = _ddc(_bbbd.Data[_affe], 0xff, _eaad)
			}
		}
	case PixNotDst:
		for _fcebb = 0; _fcebb < _cfeba; _fcebb++ {
			_affe = _affa + _fcebb*_bbbd.BytesPerLine
			for _egaba = 0; _egaba < _gaaa; _egaba++ {
				_bbbd.Data[_affe] = ^_bbbd.Data[_affe]
				_affe++
			}
			if _geca > 0 {
				_bbbd.Data[_affe] = _ddc(_bbbd.Data[_affe], ^_bbbd.Data[_affe], _eaad)
			}
		}
	}
}
func GrayHistogram(g Gray) (_egeb [256]int) {
	switch _ggdf := g.(type) {
	case Histogramer:
		return _ggdf.Histogram()
	case _c.Image:
		_bcdgf := _ggdf.Bounds()
		for _bfef := 0; _bfef < _bcdgf.Max.X; _bfef++ {
			for _adef := 0; _adef < _bcdgf.Max.Y; _adef++ {
				_egeb[g.GrayAt(_bfef, _adef).Y]++
			}
		}
		return _egeb
	default:
		return [256]int{}
	}
}

var _ _c.Image = &Gray8{}

type RGBA interface {
	RGBAAt(_agfb, _caebed int) _ag.RGBA
	SetRGBA(_dcbe, _dcea int, _facg _ag.RGBA)
}

func (_gdd *Gray8) Validate() error {
	if len(_gdd.Data) != _gdd.Height*_gdd.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_gggf *NRGBA16) ColorModel() _ag.Model { return NRGBA16Model }
func (_dcfc *Gray4) Set(x, y int, c _ag.Color) {
	if x >= _dcfc.Width || y >= _dcfc.Height {
		return
	}
	_dbgb := Gray4Model.Convert(c).(_ag.Gray)
	_dcfc.setGray(x, y, _dbgb)
}
func (_gadd *Gray4) GrayAt(x, y int) _ag.Gray {
	_eegg, _ := ColorAtGray4BPC(x, y, _gadd.BytesPerLine, _gadd.Data, _gadd.Decode)
	return _eegg
}
func (_eac *Gray4) Copy() Image { return &Gray4{ImageBase: _eac.copy()} }
func (_adcg *Monochrome) Scale(scale float64) (*Monochrome, error) {
	var _gcfc bool
	_fgfd := scale
	if scale < 1 {
		_fgfd = 1 / scale
		_gcfc = true
	}
	_bcf := NextPowerOf2(uint(_fgfd))
	if InDelta(float64(_bcf), _fgfd, 0.001) {
		if _gcfc {
			return _adcg.ReduceBinary(_fgfd)
		}
		return _adcg.ExpandBinary(int(_bcf))
	}
	_cbfe := int(_e.RoundToEven(float64(_adcg.Width) * scale))
	_cbfd := int(_e.RoundToEven(float64(_adcg.Height) * scale))
	return _adcg.ScaleLow(_cbfe, _cbfd)
}
func (_ebad *Gray8) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _ebad.Width, Y: _ebad.Height}}
}
func IsPowerOf2(n uint) bool { return n > 0 && (n&(n-1)) == 0 }
func (_gacc *Gray16) GrayAt(x, y int) _ag.Gray {
	_ecdag, _ := _gacc.ColorAt(x, y)
	return _ag.Gray{Y: uint8(_ecdag.(_ag.Gray16).Y >> 8)}
}
func (_fdgc *Gray16) SetGray(x, y int, g _ag.Gray) {
	_aefcb := (y*_fdgc.BytesPerLine/2 + x) * 2
	if _aefcb+1 >= len(_fdgc.Data) {
		return
	}
	_fdgc.Data[_aefcb] = g.Y
	_fdgc.Data[_aefcb+1] = g.Y
}
func (_dbgd *Monochrome) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtGray1BPC(x, y, _dbgd.BytesPerLine, _dbgd.Data, _dbgd.Decode)
}
func _dcd(_fab *Monochrome, _agb int, _agd []byte) (_dbc *Monochrome, _dca error) {
	const _aed = "\u0072\u0065d\u0075\u0063\u0065R\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079"
	if _fab == nil {
		return nil, _a.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _agb < 1 || _agb > 4 {
		return nil, _a.New("\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _fab.Height <= 1 {
		return nil, _a.New("\u0073\u006f\u0075rc\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020m\u0075s\u0074 \u0062e\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027")
	}
	_dbc = _gea(_fab.Width/2, _fab.Height/2)
	if _agd == nil {
		_agd = _gfd()
	}
	_edb := _bbe(_fab.BytesPerLine, 2*_dbc.BytesPerLine)
	switch _agb {
	case 1:
		_dca = _gba(_fab, _dbc, _agd, _edb)
	case 2:
		_dca = _edg(_fab, _dbc, _agd, _edb)
	case 3:
		_dca = _cbc(_fab, _dbc, _agd, _edb)
	case 4:
		_dca = _efd(_fab, _dbc, _agd, _edb)
	}
	if _dca != nil {
		return nil, _dca
	}
	return _dbc, nil
}
func (_ffdc *Monochrome) setGrayBit(_eeag, _bad int) { _ffdc.Data[_eeag] |= 0x80 >> uint(_bad&7) }
func (_gdgb *Gray16) Copy() Image                    { return &Gray16{ImageBase: _gdgb.copy()} }
func _edg(_dfb, _gbe *Monochrome, _dgb []byte, _ceb int) (_fba error) {
	var (
		_ada, _fbg, _cgd, _aee, _beg, _gfc, _eag, _acf int
		_baf, _dcdg, _de, _geab                        uint32
		_eagf, _fee                                    byte
		_cca                                           uint16
	)
	_ccc := make([]byte, 4)
	_bae := make([]byte, 4)
	for _cgd = 0; _cgd < _dfb.Height-1; _cgd, _aee = _cgd+2, _aee+1 {
		_ada = _cgd * _dfb.BytesPerLine
		_fbg = _aee * _gbe.BytesPerLine
		for _beg, _gfc = 0, 0; _beg < _ceb; _beg, _gfc = _beg+4, _gfc+1 {
			for _eag = 0; _eag < 4; _eag++ {
				_acf = _ada + _beg + _eag
				if _acf <= len(_dfb.Data)-1 && _acf < _ada+_dfb.BytesPerLine {
					_ccc[_eag] = _dfb.Data[_acf]
				} else {
					_ccc[_eag] = 0x00
				}
				_acf = _ada + _dfb.BytesPerLine + _beg + _eag
				if _acf <= len(_dfb.Data)-1 && _acf < _ada+(2*_dfb.BytesPerLine) {
					_bae[_eag] = _dfb.Data[_acf]
				} else {
					_bae[_eag] = 0x00
				}
			}
			_baf = _be.BigEndian.Uint32(_ccc)
			_dcdg = _be.BigEndian.Uint32(_bae)
			_de = _baf & _dcdg
			_de |= _de << 1
			_geab = _baf | _dcdg
			_geab &= _geab << 1
			_dcdg = _de | _geab
			_dcdg &= 0xaaaaaaaa
			_baf = _dcdg | (_dcdg << 7)
			_eagf = byte(_baf >> 24)
			_fee = byte((_baf >> 8) & 0xff)
			_acf = _fbg + _gfc
			if _acf+1 == len(_gbe.Data)-1 || _acf+1 >= _fbg+_gbe.BytesPerLine {
				if _fba = _gbe.setByte(_acf, _dgb[_eagf]); _fba != nil {
					return _ea.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _acf)
				}
			} else {
				_cca = (uint16(_dgb[_eagf]) << 8) | uint16(_dgb[_fee])
				if _fba = _gbe.setTwoBytes(_acf, _cca); _fba != nil {
					return _ea.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _acf)
				}
				_gfc++
			}
		}
	}
	return nil
}
func _bdfe(_bcfd _c.Image, _cfd int) (_c.Rectangle, bool, []byte) {
	_cfff := _bcfd.Bounds()
	var (
		_eggd bool
		_bgb  []byte
	)
	switch _eged := _bcfd.(type) {
	case SMasker:
		_eggd = _eged.HasAlpha()
	case NRGBA, RGBA, *_c.RGBA64, nrgba64, *_c.NYCbCrA:
		_bgb = make([]byte, _cfff.Max.X*_cfff.Max.Y*_cfd)
	case *_c.Paletted:
		var _fcaf bool
		for _, _dcbag := range _eged.Palette {
			_fbcc, _dgcf, _bgdd, _ecegef := _dcbag.RGBA()
			if _fbcc == 0 && _dgcf == 0 && _bgdd == 0 && _ecegef != 0 {
				_fcaf = true
				break
			}
		}
		if _fcaf {
			_bgb = make([]byte, _cfff.Max.X*_cfff.Max.Y*_cfd)
		}
	}
	return _cfff, _eggd, _bgb
}
func (_bdege *Monochrome) ExpandBinary(factor int) (*Monochrome, error) {
	if !IsPowerOf2(uint(factor)) {
		return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0065\u0078\u0070\u0061\u006e\u0064\u0020b\u0069n\u0061\u0072\u0079\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", factor)
	}
	return _efb(_bdege, factor)
}
func ColorAtGray1BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_ag.Gray, error) {
	_cce := y*bytesPerLine + x>>3
	if _cce >= len(data) {
		return _ag.Gray{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_caed := data[_cce] >> uint(7-(x&7)) & 1
	if len(decode) == 2 {
		_caed = uint8(LinearInterpolate(float64(_caed), 0.0, 1.0, decode[0], decode[1])) & 1
	}
	return _ag.Gray{Y: _caed * 255}, nil
}

type Gray16 struct{ ImageBase }

var _ Gray = &Gray16{}
var (
	_gd  = _gfg()
	_gcg = _gcd()
	_gdb = _gee()
)

func (_cdgc *RGBA32) Validate() error {
	if len(_cdgc.Data) != 3*_cdgc.Width*_cdgc.Height {
		return _a.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func _fbd(_ced, _bd *Monochrome) (_ed error) {
	_bde := _bd.BytesPerLine
	_efc := _ced.BytesPerLine
	var (
		_cdb                       byte
		_aa                        uint16
		_ggc, _gca, _cec, _dd, _cc int
	)
	for _cec = 0; _cec < _bd.Height; _cec++ {
		_ggc = _cec * _bde
		_gca = 2 * _cec * _efc
		for _dd = 0; _dd < _bde; _dd++ {
			_cdb = _bd.Data[_ggc+_dd]
			_aa = _gd[_cdb]
			_cc = _gca + _dd*2
			if _ced.BytesPerLine != _bd.BytesPerLine*2 && (_dd+1)*2 > _ced.BytesPerLine {
				_ed = _ced.setByte(_cc, byte(_aa>>8))
			} else {
				_ed = _ced.setTwoBytes(_cc, _aa)
			}
			if _ed != nil {
				return _ed
			}
		}
		for _dd = 0; _dd < _efc; _dd++ {
			_cc = _gca + _efc + _dd
			_cdb = _ced.Data[_gca+_dd]
			if _ed = _ced.setByte(_cc, _cdb); _ed != nil {
				return _ed
			}
		}
	}
	return nil
}
func _gbgf(_cfe _ag.NRGBA) _ag.Gray {
	_fgga, _dgba, _bgfb, _ := _cfe.RGBA()
	_bgfbe := (19595*_fgga + 38470*_dgba + 7471*_bgfb + 1<<15) >> 24
	return _ag.Gray{Y: uint8(_bgfbe)}
}

type CMYK interface {
	CMYKAt(_dbb, _dabc int) _ag.CMYK
	SetCMYK(_ccga, _abe int, _dbcb _ag.CMYK)
}

func _ecfe(_ebdegf _c.Image) (Image, error) {
	if _daed, _bdbba := _ebdegf.(*NRGBA16); _bdbba {
		return _daed.Copy(), nil
	}
	_ebbd := _ebdegf.Bounds()
	_eccfe, _gbbg := NewImage(_ebbd.Max.X, _ebbd.Max.Y, 4, 3, nil, nil, nil)
	if _gbbg != nil {
		return nil, _gbbg
	}
	_cfaf(_ebdegf, _eccfe, _ebbd)
	return _eccfe, nil
}

var _ _c.Image = &RGBA32{}

func (_fafd *RGBA32) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtRGBA32(x, y, _fafd.Width, _fafd.Data, _fafd.Alpha, _fafd.Decode)
}
func _ffdf(_bag _c.Image, _ebfb Image, _cbfgg _c.Rectangle) {
	switch _fgddc := _bag.(type) {
	case Gray:
		_bcaa(_fgddc, _ebfb.(Gray), _cbfgg)
	case NRGBA:
		_beab(_fgddc, _ebfb.(Gray), _cbfgg)
	case CMYK:
		_bcdc(_fgddc, _ebfb.(Gray), _cbfgg)
	case RGBA:
		_efgcb(_fgddc, _ebfb.(Gray), _cbfgg)
	default:
		_eceg(_bag, _ebfb.(Image), _cbfgg)
	}
}
func (_cbbb *RGBA32) Base() *ImageBase { return &_cbbb.ImageBase }
func (_dcacd *Gray8) SetGray(x, y int, g _ag.Gray) {
	_ebde := y*_dcacd.BytesPerLine + x
	if _ebde > len(_dcacd.Data)-1 {
		return
	}
	_dcacd.Data[_ebde] = g.Y
}
func _dagc(_dbdd _ag.Gray, _bffe monochromeModel) _ag.Gray {
	if _dbdd.Y > uint8(_bffe) {
		return _ag.Gray{Y: _e.MaxUint8}
	}
	return _ag.Gray{}
}
func (_efag *NRGBA64) setNRGBA64(_adgc int, _eace _ag.NRGBA64, _bddg int) {
	_efag.Data[_adgc] = uint8(_eace.R >> 8)
	_efag.Data[_adgc+1] = uint8(_eace.R & 0xff)
	_efag.Data[_adgc+2] = uint8(_eace.G >> 8)
	_efag.Data[_adgc+3] = uint8(_eace.G & 0xff)
	_efag.Data[_adgc+4] = uint8(_eace.B >> 8)
	_efag.Data[_adgc+5] = uint8(_eace.B & 0xff)
	if _bddg+1 < len(_efag.Alpha) {
		_efag.Alpha[_bddg] = uint8(_eace.A >> 8)
		_efag.Alpha[_bddg+1] = uint8(_eace.A & 0xff)
	}
}
func (_ffcd *NRGBA32) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _ffcd.Width, Y: _ffcd.Height}}
}
func (_ecdg *ImageBase) setEightBytes(_decd int, _aaag uint64) error {
	_edbgd := _ecdg.BytesPerLine - (_decd % _ecdg.BytesPerLine)
	if _ecdg.BytesPerLine != _ecdg.Width>>3 {
		_edbgd--
	}
	if _edbgd >= 8 {
		return _ecdg.setEightFullBytes(_decd, _aaag)
	}
	return _ecdg.setEightPartlyBytes(_decd, _edbgd, _aaag)
}
func _fbe(_fbc _ag.CMYK) _ag.RGBA {
	_cbdg, _dgg, _dbd := _ag.CMYKToRGB(_fbc.C, _fbc.M, _fbc.Y, _fbc.K)
	return _ag.RGBA{R: _cbdg, G: _dgg, B: _dbd, A: 0xff}
}

type monochromeModel uint8

func _edbe(_edgf *_c.Gray16, _fedc uint8) *_c.Gray {
	_fcegb := _edgf.Bounds()
	_bgda := _c.NewGray(_fcegb)
	for _fgdc := 0; _fgdc < _fcegb.Dx(); _fgdc++ {
		for _aaga := 0; _aaga < _fcegb.Dy(); _aaga++ {
			_gcef := _edgf.Gray16At(_fgdc, _aaga)
			_bgda.SetGray(_fgdc, _aaga, _ag.Gray{Y: _aabf(uint8(_gcef.Y/256), _fedc)})
		}
	}
	return _bgda
}
func LinearInterpolate(x, xmin, xmax, ymin, ymax float64) float64 {
	if _e.Abs(xmax-xmin) < 0.000001 {
		return ymin
	}
	_fff := ymin + (x-xmin)*(ymax-ymin)/(xmax-xmin)
	return _fff
}
func (_gacb *NRGBA64) SetNRGBA64(x, y int, c _ag.NRGBA64) {
	_dfbc := (y*_gacb.Width + x) * 2
	_ggceb := _dfbc * 3
	if _ggceb+5 >= len(_gacb.Data) {
		return
	}
	_gacb.setNRGBA64(_ggceb, c, _dfbc)
}
func (_ceeaf *Monochrome) ColorModel() _ag.Model { return MonochromeModel(_ceeaf.ModelThreshold) }

var ErrInvalidImage = _a.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")

func _fdfb(_bfde nrgba64, _ddbc RGBA, _gdaae _c.Rectangle) {
	for _adbd := 0; _adbd < _gdaae.Max.X; _adbd++ {
		for _dbddd := 0; _dbddd < _gdaae.Max.Y; _dbddd++ {
			_bcddg := _bfde.NRGBA64At(_adbd, _dbddd)
			_ddbc.SetRGBA(_adbd, _dbddd, _adce(_bcddg))
		}
	}
}
func _fdcb(_geed _c.Image) (Image, error) {
	if _cfeb, _dbef := _geed.(*Gray4); _dbef {
		return _cfeb.Copy(), nil
	}
	_dcfcc := _geed.Bounds()
	_ccee, _fbfb := NewImage(_dcfcc.Max.X, _dcfcc.Max.Y, 4, 1, nil, nil, nil)
	if _fbfb != nil {
		return nil, _fbfb
	}
	_ffdf(_geed, _ccee, _dcfcc)
	return _ccee, nil
}
func _fbdc(_fggg *Monochrome, _fbdg, _fafg, _fedd, _aade int, _bffec RasterOperator, _bfe *Monochrome, _dbgda, _dbda int) error {
	var (
		_bdgb        bool
		_cafd        bool
		_caeee       int
		_bdcb        int
		_eagcg       int
		_bebf        bool
		_acee        byte
		_gabed       int
		_gecef       int
		_aacc        int
		_ffgb, _aggg int
	)
	_bcaaf := 8 - (_fbdg & 7)
	_efea := _face[_bcaaf]
	_bffaf := _fggg.BytesPerLine*_fafg + (_fbdg >> 3)
	_dgbe := _bfe.BytesPerLine*_dbda + (_dbgda >> 3)
	if _fedd < _bcaaf {
		_bdgb = true
		_efea &= _bcdf[8-_bcaaf+_fedd]
	}
	if !_bdgb {
		_caeee = (_fedd - _bcaaf) >> 3
		if _caeee > 0 {
			_cafd = true
			_bdcb = _bffaf + 1
			_eagcg = _dgbe + 1
		}
	}
	_gabed = (_fbdg + _fedd) & 7
	if !(_bdgb || _gabed == 0) {
		_bebf = true
		_acee = _bcdf[_gabed]
		_gecef = _bffaf + 1 + _caeee
		_aacc = _dgbe + 1 + _caeee
	}
	switch _bffec {
	case PixSrc:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], _bfe.Data[_dgbe], _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] = _bfe.Data[_eagcg+_aggg]
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], _bfe.Data[_aacc], _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixNotSrc:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], ^_bfe.Data[_dgbe], _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] = ^_bfe.Data[_eagcg+_aggg]
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], ^_bfe.Data[_aacc], _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixSrcOrDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], _bfe.Data[_dgbe]|_fggg.Data[_bffaf], _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] |= _bfe.Data[_eagcg+_aggg]
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], _bfe.Data[_aacc]|_fggg.Data[_gecef], _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixSrcAndDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], _bfe.Data[_dgbe]&_fggg.Data[_bffaf], _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] &= _bfe.Data[_eagcg+_aggg]
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], _bfe.Data[_aacc]&_fggg.Data[_gecef], _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixSrcXorDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], _bfe.Data[_dgbe]^_fggg.Data[_bffaf], _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] ^= _bfe.Data[_eagcg+_aggg]
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], _bfe.Data[_aacc]^_fggg.Data[_gecef], _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], ^(_bfe.Data[_dgbe])|_fggg.Data[_bffaf], _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] |= ^(_bfe.Data[_eagcg+_aggg])
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], ^(_bfe.Data[_aacc])|_fggg.Data[_gecef], _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], ^(_bfe.Data[_dgbe])&_fggg.Data[_bffaf], _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] &= ^_bfe.Data[_eagcg+_aggg]
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], ^(_bfe.Data[_aacc])&_fggg.Data[_gecef], _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], _bfe.Data[_dgbe]|^(_fggg.Data[_bffaf]), _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] = _bfe.Data[_eagcg+_aggg] | ^(_fggg.Data[_bdcb+_aggg])
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], _bfe.Data[_aacc]|^(_fggg.Data[_gecef]), _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], _bfe.Data[_dgbe]&^(_fggg.Data[_bffaf]), _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] = _bfe.Data[_eagcg+_aggg] &^ (_fggg.Data[_bdcb+_aggg])
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], _bfe.Data[_aacc]&^(_fggg.Data[_gecef]), _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], ^(_bfe.Data[_dgbe] | _fggg.Data[_bffaf]), _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] = ^(_bfe.Data[_eagcg+_aggg] | _fggg.Data[_bdcb+_aggg])
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], ^(_bfe.Data[_aacc] | _fggg.Data[_gecef]), _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], ^(_bfe.Data[_dgbe] & _fggg.Data[_bffaf]), _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] = ^(_bfe.Data[_eagcg+_aggg] & _fggg.Data[_bdcb+_aggg])
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], ^(_bfe.Data[_aacc] & _fggg.Data[_gecef]), _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		for _ffgb = 0; _ffgb < _aade; _ffgb++ {
			_fggg.Data[_bffaf] = _ddc(_fggg.Data[_bffaf], ^(_bfe.Data[_dgbe] ^ _fggg.Data[_bffaf]), _efea)
			_bffaf += _fggg.BytesPerLine
			_dgbe += _bfe.BytesPerLine
		}
		if _cafd {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				for _aggg = 0; _aggg < _caeee; _aggg++ {
					_fggg.Data[_bdcb+_aggg] = ^(_bfe.Data[_eagcg+_aggg] ^ _fggg.Data[_bdcb+_aggg])
				}
				_bdcb += _fggg.BytesPerLine
				_eagcg += _bfe.BytesPerLine
			}
		}
		if _bebf {
			for _ffgb = 0; _ffgb < _aade; _ffgb++ {
				_fggg.Data[_gecef] = _ddc(_fggg.Data[_gecef], ^(_bfe.Data[_aacc] ^ _fggg.Data[_gecef]), _acee)
				_gecef += _fggg.BytesPerLine
				_aacc += _bfe.BytesPerLine
			}
		}
	default:
		_ca.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _bffec)
		return _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_aacb *Monochrome) Copy() Image {
	return &Monochrome{ImageBase: _aacb.ImageBase.copy(), ModelThreshold: _aacb.ModelThreshold}
}
func _cbc(_fbfa, _cba *Monochrome, _gbee []byte, _aeb int) (_cbd error) {
	var (
		_cge, _gaa, _daaa, _ddf, _affg, _dab, _dee, _cccd int
		_ceg, _bfb, _fcf, _cad                            uint32
		_ggee, _aaa                                       byte
		_eca                                              uint16
	)
	_dfgg := make([]byte, 4)
	_eddb := make([]byte, 4)
	for _daaa = 0; _daaa < _fbfa.Height-1; _daaa, _ddf = _daaa+2, _ddf+1 {
		_cge = _daaa * _fbfa.BytesPerLine
		_gaa = _ddf * _cba.BytesPerLine
		for _affg, _dab = 0, 0; _affg < _aeb; _affg, _dab = _affg+4, _dab+1 {
			for _dee = 0; _dee < 4; _dee++ {
				_cccd = _cge + _affg + _dee
				if _cccd <= len(_fbfa.Data)-1 && _cccd < _cge+_fbfa.BytesPerLine {
					_dfgg[_dee] = _fbfa.Data[_cccd]
				} else {
					_dfgg[_dee] = 0x00
				}
				_cccd = _cge + _fbfa.BytesPerLine + _affg + _dee
				if _cccd <= len(_fbfa.Data)-1 && _cccd < _cge+(2*_fbfa.BytesPerLine) {
					_eddb[_dee] = _fbfa.Data[_cccd]
				} else {
					_eddb[_dee] = 0x00
				}
			}
			_ceg = _be.BigEndian.Uint32(_dfgg)
			_bfb = _be.BigEndian.Uint32(_eddb)
			_fcf = _ceg & _bfb
			_fcf |= _fcf << 1
			_cad = _ceg | _bfb
			_cad &= _cad << 1
			_bfb = _fcf & _cad
			_bfb &= 0xaaaaaaaa
			_ceg = _bfb | (_bfb << 7)
			_ggee = byte(_ceg >> 24)
			_aaa = byte((_ceg >> 8) & 0xff)
			_cccd = _gaa + _dab
			if _cccd+1 == len(_cba.Data)-1 || _cccd+1 >= _gaa+_cba.BytesPerLine {
				if _cbd = _cba.setByte(_cccd, _gbee[_ggee]); _cbd != nil {
					return _ea.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _cccd)
				}
			} else {
				_eca = (uint16(_gbee[_ggee]) << 8) | uint16(_gbee[_aaa])
				if _cbd = _cba.setTwoBytes(_cccd, _eca); _cbd != nil {
					return _ea.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _cccd)
				}
				_dab++
			}
		}
	}
	return nil
}
func _gfd() (_ecc []byte) {
	_ecc = make([]byte, 256)
	for _aad := 0; _aad < 256; _aad++ {
		_gdbd := byte(_aad)
		_ecc[_gdbd] = (_gdbd & 0x01) | ((_gdbd & 0x04) >> 1) | ((_gdbd & 0x10) >> 2) | ((_gdbd & 0x40) >> 3) | ((_gdbd & 0x02) << 3) | ((_gdbd & 0x08) << 2) | ((_gdbd & 0x20) << 1) | (_gdbd & 0x80)
	}
	return _ecc
}
func _efgcb(_ecfd RGBA, _eacg Gray, _ggd _c.Rectangle) {
	for _eefa := 0; _eefa < _ggd.Max.X; _eefa++ {
		for _gdba := 0; _gdba < _ggd.Max.Y; _gdba++ {
			_eead := _dbg(_ecfd.RGBAAt(_eefa, _gdba))
			_eacg.SetGray(_eefa, _gdba, _eead)
		}
	}
}
func InDelta(expected, current, delta float64) bool {
	_gbce := expected - current
	if _gbce <= -delta || _gbce >= delta {
		return false
	}
	return true
}
func _gega(_egcd *Monochrome, _fgdaf, _dbbd, _acffa, _dgfc int, _dbcfd RasterOperator, _gefb *Monochrome, _egef, _dage int) error {
	var (
		_bebg        bool
		_ffeg        bool
		_fccg        byte
		_bcaac       int
		_dccd        int
		_cbda        int
		_cgdd        int
		_edge        bool
		_dcffa       int
		_fcgc        int
		_cbfgga      int
		_aadf        bool
		_dbea        byte
		_caebe       int
		_abf         int
		_egab        int
		_gbdd        byte
		_cdgg        int
		_dcfcg       int
		_ccdd        uint
		_dbf         uint
		_gceb        byte
		_aaea        shift
		_bbce        bool
		_dcfab       bool
		_bbfc, _bbag int
	)
	if _egef&7 != 0 {
		_dcfcg = 8 - (_egef & 7)
	}
	if _fgdaf&7 != 0 {
		_dccd = 8 - (_fgdaf & 7)
	}
	if _dcfcg == 0 && _dccd == 0 {
		_gceb = _face[0]
	} else {
		if _dccd > _dcfcg {
			_ccdd = uint(_dccd - _dcfcg)
		} else {
			_ccdd = uint(8 - (_dcfcg - _dccd))
		}
		_dbf = 8 - _ccdd
		_gceb = _face[_ccdd]
	}
	if (_fgdaf & 7) != 0 {
		_bebg = true
		_bcaac = 8 - (_fgdaf & 7)
		_fccg = _face[_bcaac]
		_cbda = _egcd.BytesPerLine*_dbbd + (_fgdaf >> 3)
		_cgdd = _gefb.BytesPerLine*_dage + (_egef >> 3)
		_cdgg = 8 - (_egef & 7)
		if _bcaac > _cdgg {
			_aaea = _egff
			if _acffa >= _dcfcg {
				_bbce = true
			}
		} else {
			_aaea = _fdga
		}
	}
	if _acffa < _bcaac {
		_ffeg = true
		_fccg &= _bcdf[8-_bcaac+_acffa]
	}
	if !_ffeg {
		_dcffa = (_acffa - _bcaac) >> 3
		if _dcffa != 0 {
			_edge = true
			_fcgc = _egcd.BytesPerLine*_dbbd + ((_fgdaf + _dccd) >> 3)
			_cbfgga = _gefb.BytesPerLine*_dage + ((_egef + _dccd) >> 3)
		}
	}
	_caebe = (_fgdaf + _acffa) & 7
	if !(_ffeg || _caebe == 0) {
		_aadf = true
		_dbea = _bcdf[_caebe]
		_abf = _egcd.BytesPerLine*_dbbd + ((_fgdaf + _dccd) >> 3) + _dcffa
		_egab = _gefb.BytesPerLine*_dage + ((_egef + _dccd) >> 3) + _dcffa
		if _caebe > int(_dbf) {
			_dcfab = true
		}
	}
	switch _dbcfd {
	case PixSrc:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], _gbdd, _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] = _gbdd
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], _gbdd, _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixNotSrc:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], ^_gbdd, _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] = ^_gbdd
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], ^_gbdd, _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixSrcOrDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], _gbdd|_egcd.Data[_cbda], _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] |= _gbdd
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], _gbdd|_egcd.Data[_abf], _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixSrcAndDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], _gbdd&_egcd.Data[_cbda], _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] &= _gbdd
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], _gbdd&_egcd.Data[_abf], _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixSrcXorDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], _gbdd^_egcd.Data[_cbda], _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] ^= _gbdd
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], _gbdd^_egcd.Data[_abf], _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], ^_gbdd|_egcd.Data[_cbda], _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] |= ^_gbdd
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], ^_gbdd|_egcd.Data[_abf], _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], ^_gbdd&_egcd.Data[_cbda], _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] &= ^_gbdd
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], ^_gbdd&_egcd.Data[_abf], _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], _gbdd|^_egcd.Data[_cbda], _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] = _gbdd | ^_egcd.Data[_fcgc+_bbag]
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], _gbdd|^_egcd.Data[_abf], _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], _gbdd&^_egcd.Data[_cbda], _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] = _gbdd &^ _egcd.Data[_fcgc+_bbag]
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], _gbdd&^_egcd.Data[_abf], _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], ^(_gbdd | _egcd.Data[_cbda]), _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] = ^(_gbdd | _egcd.Data[_fcgc+_bbag])
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], ^(_gbdd | _egcd.Data[_abf]), _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], ^(_gbdd & _egcd.Data[_cbda]), _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] = ^(_gbdd & _egcd.Data[_fcgc+_bbag])
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], ^(_gbdd & _egcd.Data[_abf]), _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		if _bebg {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				if _aaea == _egff {
					_gbdd = _gefb.Data[_cgdd] << _ccdd
					if _bbce {
						_gbdd = _ddc(_gbdd, _gefb.Data[_cgdd+1]>>_dbf, _gceb)
					}
				} else {
					_gbdd = _gefb.Data[_cgdd] >> _dbf
				}
				_egcd.Data[_cbda] = _ddc(_egcd.Data[_cbda], ^(_gbdd ^ _egcd.Data[_cbda]), _fccg)
				_cbda += _egcd.BytesPerLine
				_cgdd += _gefb.BytesPerLine
			}
		}
		if _edge {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				for _bbag = 0; _bbag < _dcffa; _bbag++ {
					_gbdd = _ddc(_gefb.Data[_cbfgga+_bbag]<<_ccdd, _gefb.Data[_cbfgga+_bbag+1]>>_dbf, _gceb)
					_egcd.Data[_fcgc+_bbag] = ^(_gbdd ^ _egcd.Data[_fcgc+_bbag])
				}
				_fcgc += _egcd.BytesPerLine
				_cbfgga += _gefb.BytesPerLine
			}
		}
		if _aadf {
			for _bbfc = 0; _bbfc < _dgfc; _bbfc++ {
				_gbdd = _gefb.Data[_egab] << _ccdd
				if _dcfab {
					_gbdd = _ddc(_gbdd, _gefb.Data[_egab+1]>>_dbf, _gceb)
				}
				_egcd.Data[_abf] = _ddc(_egcd.Data[_abf], ^(_gbdd ^ _egcd.Data[_abf]), _dbea)
				_abf += _egcd.BytesPerLine
				_egab += _gefb.BytesPerLine
			}
		}
	default:
		_ca.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _dbcfd)
		return _a.New("\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}
func (_fgd *CMYK32) Base() *ImageBase { return &_fgd.ImageBase }
func (_abef *Gray16) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtGray16BPC(x, y, _abef.BytesPerLine, _abef.Data, _abef.Decode)
}

var _ _c.Image = &Gray16{}

func NewImageBase(width int, height int, bitsPerComponent int, colorComponents int, data []byte, alpha []byte, decode []float64) ImageBase {
	_daeb := ImageBase{Width: width, Height: height, BitsPerComponent: bitsPerComponent, ColorComponents: colorComponents, Data: data, Alpha: alpha, Decode: decode, BytesPerLine: BytesPerLine(width, bitsPerComponent, colorComponents)}
	if data == nil {
		_daeb.Data = make([]byte, height*_daeb.BytesPerLine)
	}
	return _daeb
}
func (_faea *Gray4) Validate() error {
	if len(_faea.Data) != _faea.Height*_faea.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func _efb(_g *Monochrome, _fd int) (*Monochrome, error) {
	if _g == nil {
		return nil, _a.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _fd == 1 {
		return _g.copy(), nil
	}
	if !IsPowerOf2(uint(_fd)) {
		return nil, _ea.Errorf("\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006ci\u0064 \u0065x\u0070a\u006e\u0064\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _fd)
	}
	_gc := _edc(_fd)
	return _da(_g, _fd, _gc)
}
func (_dbba *ImageBase) Pix() []byte { return _dbba.Data }
func _cbag(_cab _ag.RGBA) _ag.CMYK {
	_fed, _dad, _cddd, _bdee := _ag.RGBToCMYK(_cab.R, _cab.G, _cab.B)
	return _ag.CMYK{C: _fed, M: _dad, Y: _cddd, K: _bdee}
}
func _adb(_ccf *Monochrome, _ba, _fde int) (*Monochrome, error) {
	if _ccf == nil {
		return nil, _a.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ba <= 0 || _fde <= 0 {
		return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _ba == _fde {
		if _ba == 1 {
			return _ccf.copy(), nil
		}
		if _ba == 2 || _ba == 4 || _ba == 8 {
			_gcf, _ffc := _efb(_ccf, _ba)
			if _ffc != nil {
				return nil, _ffc
			}
			return _gcf, nil
		}
	}
	_ffb := _ba * _ccf.Width
	_db := _fde * _ccf.Height
	_dfg := _gea(_ffb, _db)
	_bge := _dfg.BytesPerLine
	var (
		_dfe, _bee, _daa, _aaf, _bf int
		_bcd                        byte
		_bgee                       error
	)
	for _bee = 0; _bee < _ccf.Height; _bee++ {
		_dfe = _fde * _bee * _bge
		for _daa = 0; _daa < _ccf.Width; _daa++ {
			if _fa := _ccf.getBitAt(_daa, _bee); _fa {
				_bf = _ba * _daa
				for _aaf = 0; _aaf < _ba; _aaf++ {
					_dfg.setIndexedBit(_dfe*8 + _bf + _aaf)
				}
			}
		}
		for _aaf = 1; _aaf < _fde; _aaf++ {
			_dg := _dfe + _aaf*_bge
			for _eece := 0; _eece < _bge; _eece++ {
				if _bcd, _bgee = _dfg.getByte(_dfe + _eece); _bgee != nil {
					return nil, _bgee
				}
				if _bgee = _dfg.setByte(_dg+_eece, _bcd); _bgee != nil {
					return nil, _bgee
				}
			}
		}
	}
	return _dfg, nil
}
func NewImage(width, height, bitsPerComponent, colorComponents int, data, alpha []byte, decode []float64) (Image, error) {
	_gabe := NewImageBase(width, height, bitsPerComponent, colorComponents, data, alpha, decode)
	var _bede Image
	switch colorComponents {
	case 1:
		switch bitsPerComponent {
		case 1:
			_bede = &Monochrome{ImageBase: _gabe, ModelThreshold: 0x0f}
		case 2:
			_bede = &Gray2{ImageBase: _gabe}
		case 4:
			_bede = &Gray4{ImageBase: _gabe}
		case 8:
			_bede = &Gray8{ImageBase: _gabe}
		case 16:
			_bede = &Gray16{ImageBase: _gabe}
		}
	case 3:
		switch bitsPerComponent {
		case 4:
			_bede = &NRGBA16{ImageBase: _gabe}
		case 8:
			_bede = &NRGBA32{ImageBase: _gabe}
		case 16:
			_bede = &NRGBA64{ImageBase: _gabe}
		}
	case 4:
		_bede = &CMYK32{ImageBase: _gabe}
	}
	if _bede == nil {
		return nil, ErrInvalidImage
	}
	return _bede, nil
}
func _ee(_dde, _cee *Monochrome) (_eb error) {
	_gf := _cee.BytesPerLine
	_cae := _dde.BytesPerLine
	_ab := _cee.BytesPerLine*4 - _dde.BytesPerLine
	var (
		_ae, _eba                            byte
		_fe                                  uint32
		_ec, _ge, _efe, _bcg, _ad, _gb, _bga int
	)
	for _efe = 0; _efe < _cee.Height; _efe++ {
		_ec = _efe * _gf
		_ge = 4 * _efe * _cae
		for _bcg = 0; _bcg < _gf; _bcg++ {
			_ae = _cee.Data[_ec+_bcg]
			_fe = _gcg[_ae]
			_gb = _ge + _bcg*4
			if _ab != 0 && (_bcg+1)*4 > _dde.BytesPerLine {
				for _ad = _ab; _ad > 0; _ad-- {
					_eba = byte((_fe >> uint(_ad*8)) & 0xff)
					_bga = _gb + (_ab - _ad)
					if _eb = _dde.setByte(_bga, _eba); _eb != nil {
						return _eb
					}
				}
			} else if _eb = _dde.setFourBytes(_gb, _fe); _eb != nil {
				return _eb
			}
			if _eb = _dde.setFourBytes(_ge+_bcg*4, _gcg[_cee.Data[_ec+_bcg]]); _eb != nil {
				return _eb
			}
		}
		for _ad = 1; _ad < 4; _ad++ {
			for _bcg = 0; _bcg < _cae; _bcg++ {
				if _eb = _dde.setByte(_ge+_ad*_cae+_bcg, _dde.Data[_ge+_bcg]); _eb != nil {
					return _eb
				}
			}
		}
	}
	return nil
}
func (_ebec *NRGBA32) Validate() error {
	if len(_ebec.Data) != 3*_ebec.Width*_ebec.Height {
		return _a.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func (_dffd *NRGBA16) Validate() error {
	if len(_dffd.Data) != 3*_dffd.Width*_dffd.Height/2 {
		return _a.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func (_bgfc *NRGBA16) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtNRGBA16(x, y, _bgfc.Width, _bgfc.BytesPerLine, _bgfc.Data, _bgfc.Alpha, _bgfc.Decode)
}
func (_ecda *Monochrome) IsUnpadded() bool { return (_ecda.Width * _ecda.Height) == len(_ecda.Data) }
func (_gcebf *RGBA32) Copy() Image         { return &RGBA32{ImageBase: _gcebf.copy()} }
func (_aceb *RGBA32) RGBAAt(x, y int) _ag.RGBA {
	_afbf, _ := ColorAtRGBA32(x, y, _aceb.Width, _aceb.Data, _aceb.Alpha, _aceb.Decode)
	return _afbf
}
func ImgToGray(i _c.Image) *_c.Gray {
	if _dagf, _ebdc := i.(*_c.Gray); _ebdc {
		return _dagf
	}
	_eefg := i.Bounds()
	_bfcd := _c.NewGray(_eefg)
	for _defd := 0; _defd < _eefg.Max.X; _defd++ {
		for _aedge := 0; _aedge < _eefg.Max.Y; _aedge++ {
			_ceee := i.At(_defd, _aedge)
			_bfcd.Set(_defd, _aedge, _ceee)
		}
	}
	return _bfcd
}
func (_fdec *ImageBase) HasAlpha() bool {
	if _fdec.Alpha == nil {
		return false
	}
	for _abeg := range _fdec.Alpha {
		if _fdec.Alpha[_abeg] != 0xff {
			return true
		}
	}
	return false
}
func (_eeagb *NRGBA32) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtNRGBA32(x, y, _eeagb.Width, _eeagb.Data, _eeagb.Alpha, _eeagb.Decode)
}
func (_cecb *ImageBase) setFourBytes(_dagb int, _dbgf uint32) error {
	if _dagb+3 > len(_cecb.Data)-1 {
		return _ea.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _dagb)
	}
	_cecb.Data[_dagb] = byte((_dbgf & 0xff000000) >> 24)
	_cecb.Data[_dagb+1] = byte((_dbgf & 0xff0000) >> 16)
	_cecb.Data[_dagb+2] = byte((_dbgf & 0xff00) >> 8)
	_cecb.Data[_dagb+3] = byte(_dbgf & 0xff)
	return nil
}

var _ Image = &Gray16{}

func (_aacdd *Gray2) Validate() error {
	if len(_aacdd.Data) != _aacdd.Height*_aacdd.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func _begcf(_gdgg RGBA, _aegbe NRGBA, _gffc _c.Rectangle) {
	for _adbg := 0; _adbg < _gffc.Max.X; _adbg++ {
		for _cgcd := 0; _cgcd < _gffc.Max.Y; _cgcd++ {
			_egea := _gdgg.RGBAAt(_adbg, _cgcd)
			_aegbe.SetNRGBA(_adbg, _cgcd, _gcga(_egea))
		}
	}
}
func _gcd() (_cda [256]uint32) {
	for _eef := 0; _eef < 256; _eef++ {
		if _eef&0x01 != 0 {
			_cda[_eef] |= 0xf
		}
		if _eef&0x02 != 0 {
			_cda[_eef] |= 0xf0
		}
		if _eef&0x04 != 0 {
			_cda[_eef] |= 0xf00
		}
		if _eef&0x08 != 0 {
			_cda[_eef] |= 0xf000
		}
		if _eef&0x10 != 0 {
			_cda[_eef] |= 0xf0000
		}
		if _eef&0x20 != 0 {
			_cda[_eef] |= 0xf00000
		}
		if _eef&0x40 != 0 {
			_cda[_eef] |= 0xf000000
		}
		if _eef&0x80 != 0 {
			_cda[_eef] |= 0xf0000000
		}
	}
	return _cda
}

type Image interface {
	_f.Image
	Base() *ImageBase
	Copy() Image
	Pix() []byte
	ColorAt(_cede, _bggcd int) (_ag.Color, error)
	Validate() error
}

func (_aaeae *NRGBA32) setRGBA(_cggb int, _aaaeg _ag.NRGBA) {
	_daf := 3 * _cggb
	_aaeae.Data[_daf] = _aaaeg.R
	_aaeae.Data[_daf+1] = _aaaeg.G
	_aaeae.Data[_daf+2] = _aaaeg.B
	if _cggb < len(_aaeae.Alpha) {
		_aaeae.Alpha[_cggb] = _aaaeg.A
	}
}
func (_cged *RGBA32) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _cged.Width, Y: _cged.Height}}
}
func (_dbce *Gray4) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _dbce.Width, Y: _dbce.Height}}
}

type shift int

func _gdaa(_fcgf *Monochrome, _dbab, _ddea, _caeg, _addae int, _dgbbe RasterOperator, _geegc *Monochrome, _gegc, _cace int) error {
	var (
		_bbga        byte
		_fdbc        int
		_adg         int
		_egdb, _eaba int
		_gfgd, _aaaf int
	)
	_aegg := _caeg >> 3
	_cdge := _caeg & 7
	if _cdge > 0 {
		_bbga = _bcdf[_cdge]
	}
	_fdbc = _geegc.BytesPerLine*_cace + (_gegc >> 3)
	_adg = _fcgf.BytesPerLine*_ddea + (_dbab >> 3)
	switch _dgbbe {
	case PixSrc:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] = _geegc.Data[_egdb]
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], _geegc.Data[_egdb], _bbga)
			}
		}
	case PixNotSrc:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] = ^(_geegc.Data[_egdb])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], ^_geegc.Data[_egdb], _bbga)
			}
		}
	case PixSrcOrDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] |= _geegc.Data[_egdb]
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], _geegc.Data[_egdb]|_fcgf.Data[_eaba], _bbga)
			}
		}
	case PixSrcAndDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] &= _geegc.Data[_egdb]
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], _geegc.Data[_egdb]&_fcgf.Data[_eaba], _bbga)
			}
		}
	case PixSrcXorDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] ^= _geegc.Data[_egdb]
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], _geegc.Data[_egdb]^_fcgf.Data[_eaba], _bbga)
			}
		}
	case PixNotSrcOrDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] |= ^(_geegc.Data[_egdb])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], ^(_geegc.Data[_egdb])|_fcgf.Data[_eaba], _bbga)
			}
		}
	case PixNotSrcAndDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] &= ^(_geegc.Data[_egdb])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], ^(_geegc.Data[_egdb])&_fcgf.Data[_eaba], _bbga)
			}
		}
	case PixSrcOrNotDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] = _geegc.Data[_egdb] | ^(_fcgf.Data[_eaba])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], _geegc.Data[_egdb]|^(_fcgf.Data[_eaba]), _bbga)
			}
		}
	case PixSrcAndNotDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] = _geegc.Data[_egdb] &^ (_fcgf.Data[_eaba])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], _geegc.Data[_egdb]&^(_fcgf.Data[_eaba]), _bbga)
			}
		}
	case PixNotPixSrcOrDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] = ^(_geegc.Data[_egdb] | _fcgf.Data[_eaba])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], ^(_geegc.Data[_egdb] | _fcgf.Data[_eaba]), _bbga)
			}
		}
	case PixNotPixSrcAndDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] = ^(_geegc.Data[_egdb] & _fcgf.Data[_eaba])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], ^(_geegc.Data[_egdb] & _fcgf.Data[_eaba]), _bbga)
			}
		}
	case PixNotPixSrcXorDst:
		for _gfgd = 0; _gfgd < _addae; _gfgd++ {
			_egdb = _fdbc + _gfgd*_geegc.BytesPerLine
			_eaba = _adg + _gfgd*_fcgf.BytesPerLine
			for _aaaf = 0; _aaaf < _aegg; _aaaf++ {
				_fcgf.Data[_eaba] = ^(_geegc.Data[_egdb] ^ _fcgf.Data[_eaba])
				_eaba++
				_egdb++
			}
			if _cdge > 0 {
				_fcgf.Data[_eaba] = _ddc(_fcgf.Data[_eaba], ^(_geegc.Data[_egdb] ^ _fcgf.Data[_eaba]), _bbga)
			}
		}
	default:
		_ca.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _dgbbe)
		return _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func _ceae(_addaf _c.Image) (Image, error) {
	if _fcc, _gccb := _addaf.(*Gray8); _gccb {
		return _fcc.Copy(), nil
	}
	_gece := _addaf.Bounds()
	_abgd, _egc := NewImage(_gece.Max.X, _gece.Max.Y, 8, 1, nil, nil, nil)
	if _egc != nil {
		return nil, _egc
	}
	_ffdf(_addaf, _abgd, _gece)
	return _abgd, nil
}
func (_egbg *Gray16) At(x, y int) _ag.Color { _gccg, _ := _egbg.ColorAt(x, y); return _gccg }
func (_bcfg *Monochrome) GrayAt(x, y int) _ag.Gray {
	_eafb, _ := ColorAtGray1BPC(x, y, _bcfg.BytesPerLine, _bcfg.Data, _bcfg.Decode)
	return _eafb
}
func (_bcga *Gray16) Set(x, y int, c _ag.Color) {
	_cfac := (y*_bcga.BytesPerLine/2 + x) * 2
	if _cfac+1 >= len(_bcga.Data) {
		return
	}
	_dff := _ag.Gray16Model.Convert(c).(_ag.Gray16)
	_bcga.Data[_cfac], _bcga.Data[_cfac+1] = uint8(_dff.Y>>8), uint8(_dff.Y&0xff)
}
func (_aaed *NRGBA32) Base() *ImageBase     { return &_aaed.ImageBase }
func (_ceff *Gray4) ColorModel() _ag.Model  { return Gray4Model }
func (_dade *RGBA32) At(x, y int) _ag.Color { _dbbdf, _ := _dade.ColorAt(x, y); return _dbbdf }
func (_cffa *Gray16) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _cffa.Width, Y: _cffa.Height}}
}

var _ NRGBA = &NRGBA32{}

func _bbd(_accb *Monochrome, _ebae, _gbed int, _acbg, _gaf int, _gfbb RasterOperator, _aag *Monochrome, _ebabf, _gcag int) error {
	var _bba, _abcc, _gdaf, _cag int
	if _ebae < 0 {
		_ebabf -= _ebae
		_acbg += _ebae
		_ebae = 0
	}
	if _ebabf < 0 {
		_ebae -= _ebabf
		_acbg += _ebabf
		_ebabf = 0
	}
	_bba = _ebae + _acbg - _accb.Width
	if _bba > 0 {
		_acbg -= _bba
	}
	_abcc = _ebabf + _acbg - _aag.Width
	if _abcc > 0 {
		_acbg -= _abcc
	}
	if _gbed < 0 {
		_gcag -= _gbed
		_gaf += _gbed
		_gbed = 0
	}
	if _gcag < 0 {
		_gbed -= _gcag
		_gaf += _gcag
		_gcag = 0
	}
	_gdaf = _gbed + _gaf - _accb.Height
	if _gdaf > 0 {
		_gaf -= _gdaf
	}
	_cag = _gcag + _gaf - _aag.Height
	if _cag > 0 {
		_gaf -= _cag
	}
	if _acbg <= 0 || _gaf <= 0 {
		return nil
	}
	var _fca error
	switch {
	case _ebae&7 == 0 && _ebabf&7 == 0:
		_fca = _gdaa(_accb, _ebae, _gbed, _acbg, _gaf, _gfbb, _aag, _ebabf, _gcag)
	case _ebae&7 == _ebabf&7:
		_fca = _fbdc(_accb, _ebae, _gbed, _acbg, _gaf, _gfbb, _aag, _ebabf, _gcag)
	default:
		_fca = _gega(_accb, _ebae, _gbed, _acbg, _gaf, _gfbb, _aag, _ebabf, _gcag)
	}
	if _fca != nil {
		return _fca
	}
	return nil
}
func (_ceac *CMYK32) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtCMYK(x, y, _ceac.Width, _ceac.Data, _ceac.Decode)
}
func (_adda *Gray4) Histogram() (_fgda [256]int) {
	for _aeec := 0; _aeec < _adda.Width; _aeec++ {
		for _cabb := 0; _cabb < _adda.Height; _cabb++ {
			_fgda[_adda.GrayAt(_aeec, _cabb).Y]++
		}
	}
	return _fgda
}
func FromGoImage(i _c.Image) (Image, error) {
	switch _adf := i.(type) {
	case Image:
		return _adf.Copy(), nil
	case Gray:
		return GrayConverter.Convert(i)
	case *_c.Gray16:
		return Gray16Converter.Convert(i)
	case CMYK:
		return CMYKConverter.Convert(i)
	case *_c.NRGBA64:
		return NRGBA64Converter.Convert(i)
	default:
		return NRGBAConverter.Convert(i)
	}
}
func (_bafe *NRGBA64) Validate() error {
	if len(_bafe.Data) != 3*2*_bafe.Width*_bafe.Height {
		return _a.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

const (
	_egff shift = iota
	_fdga
)

func _gee() (_cg [256]uint64) {
	for _cgb := 0; _cgb < 256; _cgb++ {
		if _cgb&0x01 != 0 {
			_cg[_cgb] |= 0xff
		}
		if _cgb&0x02 != 0 {
			_cg[_cgb] |= 0xff00
		}
		if _cgb&0x04 != 0 {
			_cg[_cgb] |= 0xff0000
		}
		if _cgb&0x08 != 0 {
			_cg[_cgb] |= 0xff000000
		}
		if _cgb&0x10 != 0 {
			_cg[_cgb] |= 0xff00000000
		}
		if _cgb&0x20 != 0 {
			_cg[_cgb] |= 0xff0000000000
		}
		if _cgb&0x40 != 0 {
			_cg[_cgb] |= 0xff000000000000
		}
		if _cgb&0x80 != 0 {
			_cg[_cgb] |= 0xff00000000000000
		}
	}
	return _cg
}
func (_cebg *ImageBase) GetAlpha() []byte { return _cebg.Alpha }
func (_fadga *NRGBA64) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _fadga.Width, Y: _fadga.Height}}
}
func _agdc(_fdg _ag.NRGBA) _ag.CMYK {
	_agbc, _ddada, _gbd, _ := _fdg.RGBA()
	_fbff, _aedb, _dfec, _ccdb := _ag.RGBToCMYK(uint8(_agbc>>8), uint8(_ddada>>8), uint8(_gbd>>8))
	return _ag.CMYK{C: _fbff, M: _aedb, Y: _dfec, K: _ccdb}
}
func _dedd(_ggde, _bgec NRGBA, _aaef _c.Rectangle) {
	for _bedc := 0; _bedc < _aaef.Max.X; _bedc++ {
		for _gbbd := 0; _gbbd < _aaef.Max.Y; _gbbd++ {
			_bgec.SetNRGBA(_bedc, _gbbd, _ggde.NRGBAAt(_bedc, _gbbd))
		}
	}
}

var _ RGBA = &RGBA32{}

type NRGBA16 struct{ ImageBase }

var _ _c.Image = &Monochrome{}

func ColorAtNRGBA64(x, y, width int, data, alpha []byte, decode []float64) (_ag.NRGBA64, error) {
	_cacc := (y*width + x) * 2
	_cafdd := _cacc * 3
	if _cafdd+5 >= len(data) {
		return _ag.NRGBA64{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	const _efcg = 0xffff
	_gbde := uint16(_efcg)
	if alpha != nil && len(alpha) > _cacc+1 {
		_gbde = uint16(alpha[_cacc])<<8 | uint16(alpha[_cacc+1])
	}
	_bbgge := uint16(data[_cafdd])<<8 | uint16(data[_cafdd+1])
	_ggb := uint16(data[_cafdd+2])<<8 | uint16(data[_cafdd+3])
	_agcc := uint16(data[_cafdd+4])<<8 | uint16(data[_cafdd+5])
	if len(decode) == 6 {
		_bbgge = uint16(uint64(LinearInterpolate(float64(_bbgge), 0, 65535, decode[0], decode[1])) & _efcg)
		_ggb = uint16(uint64(LinearInterpolate(float64(_ggb), 0, 65535, decode[2], decode[3])) & _efcg)
		_agcc = uint16(uint64(LinearInterpolate(float64(_agcc), 0, 65535, decode[4], decode[5])) & _efcg)
	}
	return _ag.NRGBA64{R: _bbgge, G: _ggb, B: _agcc, A: _gbde}, nil
}

var _ Image = &NRGBA64{}

func (_aaad *RGBA32) setRGBA(_cgae int, _dffdg _ag.RGBA) {
	_gcbe := 3 * _cgae
	_aaad.Data[_gcbe] = _dffdg.R
	_aaad.Data[_gcbe+1] = _dffdg.G
	_aaad.Data[_gcbe+2] = _dffdg.B
	if _cgae < len(_aaad.Alpha) {
		_aaad.Alpha[_cgae] = _dffdg.A
	}
}
func _ccb(_eadg Gray, _cbf nrgba64, _fdfc _c.Rectangle) {
	for _aadg := 0; _aadg < _fdfc.Max.X; _aadg++ {
		for _gbcf := 0; _gbcf < _fdfc.Max.Y; _gbcf++ {
			_bfg := _ggf(_cbf.NRGBA64At(_aadg, _gbcf))
			_eadg.SetGray(_aadg, _gbcf, _bfg)
		}
	}
}
func ImgToBinary(i _c.Image, threshold uint8) *_c.Gray {
	switch _cgefc := i.(type) {
	case *_c.Gray:
		if _ecca(_cgefc) {
			return _cgefc
		}
		return _cebbe(_cgefc, threshold)
	case *_c.Gray16:
		return _edbe(_cgefc, threshold)
	default:
		return _gaed(_cgefc, threshold)
	}
}
func (_badf *Gray2) ColorModel() _ag.Model { return Gray2Model }
func _gaed(_fabbf _c.Image, _cefb uint8) *_c.Gray {
	_bgab := _fabbf.Bounds()
	_caga := _c.NewGray(_bgab)
	var (
		_bbgaa _ag.Color
		_bfba  _ag.Gray
	)
	for _gcbf := 0; _gcbf < _bgab.Max.X; _gcbf++ {
		for _dbeb := 0; _dbeb < _bgab.Max.Y; _dbeb++ {
			_bbgaa = _fabbf.At(_gcbf, _dbeb)
			_caga.Set(_gcbf, _dbeb, _bbgaa)
			_bfba = _caga.GrayAt(_gcbf, _dbeb)
			_caga.SetGray(_gcbf, _dbeb, _ag.Gray{Y: _aabf(_bfba.Y, _cefb)})
		}
	}
	return _caga
}
func _gfdf(_gecb *Monochrome, _dgd, _bgcb int, _egg, _ccae int, _fbdf RasterOperator) {
	var (
		_bggf bool
		_ddg  bool
		_bdea int
		_bcdg int
		_fag  int
		_dcbb int
		_bggb bool
		_edgb byte
	)
	_gcec := 8 - (_dgd & 7)
	_eegca := _face[_gcec]
	_ccbe := _gecb.BytesPerLine*_bgcb + (_dgd >> 3)
	if _egg < _gcec {
		_bggf = true
		_eegca &= _bcdf[8-_gcec+_egg]
	}
	if !_bggf {
		_bdea = (_egg - _gcec) >> 3
		if _bdea != 0 {
			_ddg = true
			_bcdg = _ccbe + 1
		}
	}
	_fag = (_dgd + _egg) & 7
	if !(_bggf || _fag == 0) {
		_bggb = true
		_edgb = _bcdf[_fag]
		_dcbb = _ccbe + 1 + _bdea
	}
	var _bfga, _ffef int
	switch _fbdf {
	case PixClr:
		for _bfga = 0; _bfga < _ccae; _bfga++ {
			_gecb.Data[_ccbe] = _ddc(_gecb.Data[_ccbe], 0x0, _eegca)
			_ccbe += _gecb.BytesPerLine
		}
		if _ddg {
			for _bfga = 0; _bfga < _ccae; _bfga++ {
				for _ffef = 0; _ffef < _bdea; _ffef++ {
					_gecb.Data[_bcdg+_ffef] = 0x0
				}
				_bcdg += _gecb.BytesPerLine
			}
		}
		if _bggb {
			for _bfga = 0; _bfga < _ccae; _bfga++ {
				_gecb.Data[_dcbb] = _ddc(_gecb.Data[_dcbb], 0x0, _edgb)
				_dcbb += _gecb.BytesPerLine
			}
		}
	case PixSet:
		for _bfga = 0; _bfga < _ccae; _bfga++ {
			_gecb.Data[_ccbe] = _ddc(_gecb.Data[_ccbe], 0xff, _eegca)
			_ccbe += _gecb.BytesPerLine
		}
		if _ddg {
			for _bfga = 0; _bfga < _ccae; _bfga++ {
				for _ffef = 0; _ffef < _bdea; _ffef++ {
					_gecb.Data[_bcdg+_ffef] = 0xff
				}
				_bcdg += _gecb.BytesPerLine
			}
		}
		if _bggb {
			for _bfga = 0; _bfga < _ccae; _bfga++ {
				_gecb.Data[_dcbb] = _ddc(_gecb.Data[_dcbb], 0xff, _edgb)
				_dcbb += _gecb.BytesPerLine
			}
		}
	case PixNotDst:
		for _bfga = 0; _bfga < _ccae; _bfga++ {
			_gecb.Data[_ccbe] = _ddc(_gecb.Data[_ccbe], ^_gecb.Data[_ccbe], _eegca)
			_ccbe += _gecb.BytesPerLine
		}
		if _ddg {
			for _bfga = 0; _bfga < _ccae; _bfga++ {
				for _ffef = 0; _ffef < _bdea; _ffef++ {
					_gecb.Data[_bcdg+_ffef] = ^(_gecb.Data[_bcdg+_ffef])
				}
				_bcdg += _gecb.BytesPerLine
			}
		}
		if _bggb {
			for _bfga = 0; _bfga < _ccae; _bfga++ {
				_gecb.Data[_dcbb] = _ddc(_gecb.Data[_dcbb], ^_gecb.Data[_dcbb], _edgb)
				_dcbb += _gecb.BytesPerLine
			}
		}
	}
}
func _gbdec(_badc CMYK, _faca NRGBA, _bdag _c.Rectangle) {
	for _cceg := 0; _cceg < _bdag.Max.X; _cceg++ {
		for _ddac := 0; _ddac < _bdag.Max.Y; _ddac++ {
			_cfad := _badc.CMYKAt(_cceg, _ddac)
			_faca.SetNRGBA(_cceg, _ddac, _efde(_cfad))
		}
	}
}
func (_geg *Monochrome) AddPadding() (_fgf error) {
	if _gfgg := ((_geg.Width * _geg.Height) + 7) >> 3; len(_geg.Data) < _gfgg {
		return _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(_geg.Data), _gfgg)
	}
	_bdeg := _geg.Width % 8
	if _bdeg == 0 {
		return nil
	}
	_bdb := _geg.Width / 8
	_cadd := _d.NewReader(_geg.Data)
	_cbb := make([]byte, _geg.Height*_geg.BytesPerLine)
	_dfef := _d.NewWriterMSB(_cbb)
	_ggeea := make([]byte, _bdb)
	var (
		_cbee int
		_bff  uint64
	)
	for _cbee = 0; _cbee < _geg.Height; _cbee++ {
		if _, _fgf = _cadd.Read(_ggeea); _fgf != nil {
			return _fgf
		}
		if _, _fgf = _dfef.Write(_ggeea); _fgf != nil {
			return _fgf
		}
		if _bff, _fgf = _cadd.ReadBits(byte(_bdeg)); _fgf != nil {
			return _fgf
		}
		if _fgf = _dfef.WriteByte(byte(_bff) << uint(8-_bdeg)); _fgf != nil {
			return _fgf
		}
	}
	_geg.Data = _dfef.Data()
	return nil
}
func _def(_dcabc _ag.Gray) _ag.Gray {
	_feg := _dcabc.Y >> 6
	_feg |= _feg << 2
	_dcabc.Y = _feg | _feg<<4
	return _dcabc
}

var _ _c.Image = &NRGBA16{}

func (_gabfe *Monochrome) Set(x, y int, c _ag.Color) {
	_cdae := y*_gabfe.BytesPerLine + x>>3
	if _cdae > len(_gabfe.Data)-1 {
		return
	}
	_edfd := _gabfe.ColorModel().Convert(c).(_ag.Gray)
	_gabfe.setGray(x, _edfd, _cdae)
}

var _ Image = &Gray2{}

func _aabf(_ggda, _agee uint8) uint8 {
	if _ggda < _agee {
		return 255
	}
	return 0
}
func (_eab *Gray2) Base() *ImageBase      { return &_eab.ImageBase }
func (_bdd *Gray2) At(x, y int) _ag.Color { _cfa, _ := _bdd.ColorAt(x, y); return _cfa }
func _gefd(_gcdc uint) uint {
	var _aca uint
	for _gcdc != 0 {
		_gcdc >>= 1
		_aca++
	}
	return _aca - 1
}
func (_fcee *monochromeThresholdConverter) Convert(img _c.Image) (Image, error) {
	if _acb, _fbgc := img.(*Monochrome); _fbgc {
		return _acb.Copy(), nil
	}
	_ccgad := img.Bounds()
	_gda, _fbfda := NewImage(_ccgad.Max.X, _ccgad.Max.Y, 1, 1, nil, nil, nil)
	if _fbfda != nil {
		return nil, _fbfda
	}
	_gda.(*Monochrome).ModelThreshold = _fcee.Threshold
	for _cdbe := 0; _cdbe < _ccgad.Max.X; _cdbe++ {
		for _fbfdf := 0; _fbfdf < _ccgad.Max.Y; _fbfdf++ {
			_fceb := img.At(_cdbe, _fbfdf)
			_gda.Set(_cdbe, _fbfdf, _fceb)
		}
	}
	return _gda, nil
}
func (_cga colorConverter) Convert(src _c.Image) (Image, error) { return _cga._dcc(src) }
func NextPowerOf2(n uint) uint {
	if IsPowerOf2(n) {
		return n
	}
	return 1 << (_gefd(n) + 1)
}

var _ Gray = &Gray2{}

func _eceg(_ebd _c.Image, _ecab Image, _aac _c.Rectangle) {
	for _befge := 0; _befge < _aac.Max.X; _befge++ {
		for _dcab := 0; _dcab < _aac.Max.Y; _dcab++ {
			_gbef := _ebd.At(_befge, _dcab)
			_ecab.Set(_befge, _dcab, _gbef)
		}
	}
}

var _ _c.Image = &NRGBA32{}

func _aaaa(_ebdf _ag.NRGBA) _ag.Gray {
	var _fbdd _ag.NRGBA
	if _ebdf == _fbdd {
		return _ag.Gray{Y: 0xff}
	}
	_fgdd, _dcdc, _aaeg, _ := _ebdf.RGBA()
	_afg := (19595*_fgdd + 38470*_dcdc + 7471*_aaeg + 1<<15) >> 24
	return _ag.Gray{Y: uint8(_afg)}
}
func _bcaa(_abca, _adad Gray, _ffdg _c.Rectangle) {
	for _aabb := 0; _aabb < _ffdg.Max.X; _aabb++ {
		for _dac := 0; _dac < _ffdg.Max.Y; _dac++ {
			_adad.SetGray(_aabb, _dac, _abca.GrayAt(_aabb, _dac))
		}
	}
}
func _dbg(_fdbe _ag.RGBA) _ag.Gray {
	_fabe := (19595*uint32(_fdbe.R) + 38470*uint32(_fdbe.G) + 7471*uint32(_fdbe.B) + 1<<7) >> 16
	return _ag.Gray{Y: uint8(_fabe)}
}
func MonochromeThresholdConverter(threshold uint8) ColorConverter {
	return &monochromeThresholdConverter{Threshold: threshold}
}
func _efd(_edbg, _eddg *Monochrome, _gad []byte, _aeg int) (_gcge error) {
	var (
		_ffba, _faba, _edf, _edfa, _eeg, _gfb, _cfg, _egd int
		_egdd, _cgef                                      uint32
		_fdf, _ccg                                        byte
		_dcf                                              uint16
	)
	_cea := make([]byte, 4)
	_fcg := make([]byte, 4)
	for _edf = 0; _edf < _edbg.Height-1; _edf, _edfa = _edf+2, _edfa+1 {
		_ffba = _edf * _edbg.BytesPerLine
		_faba = _edfa * _eddg.BytesPerLine
		for _eeg, _gfb = 0, 0; _eeg < _aeg; _eeg, _gfb = _eeg+4, _gfb+1 {
			for _cfg = 0; _cfg < 4; _cfg++ {
				_egd = _ffba + _eeg + _cfg
				if _egd <= len(_edbg.Data)-1 && _egd < _ffba+_edbg.BytesPerLine {
					_cea[_cfg] = _edbg.Data[_egd]
				} else {
					_cea[_cfg] = 0x00
				}
				_egd = _ffba + _edbg.BytesPerLine + _eeg + _cfg
				if _egd <= len(_edbg.Data)-1 && _egd < _ffba+(2*_edbg.BytesPerLine) {
					_fcg[_cfg] = _edbg.Data[_egd]
				} else {
					_fcg[_cfg] = 0x00
				}
			}
			_egdd = _be.BigEndian.Uint32(_cea)
			_cgef = _be.BigEndian.Uint32(_fcg)
			_cgef &= _egdd
			_cgef &= _cgef << 1
			_cgef &= 0xaaaaaaaa
			_egdd = _cgef | (_cgef << 7)
			_fdf = byte(_egdd >> 24)
			_ccg = byte((_egdd >> 8) & 0xff)
			_egd = _faba + _gfb
			if _egd+1 == len(_eddg.Data)-1 || _egd+1 >= _faba+_eddg.BytesPerLine {
				_eddg.Data[_egd] = _gad[_fdf]
				if _gcge = _eddg.setByte(_egd, _gad[_fdf]); _gcge != nil {
					return _ea.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _egd)
				}
			} else {
				_dcf = (uint16(_gad[_fdf]) << 8) | uint16(_gad[_ccg])
				if _gcge = _eddg.setTwoBytes(_egd, _dcf); _gcge != nil {
					return _ea.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _egd)
				}
				_gfb++
			}
		}
	}
	return nil
}
func ColorAt(x, y, width, bitsPerColor, colorComponents, bytesPerLine int, data, alpha []byte, decode []float64) (_ag.Color, error) {
	switch colorComponents {
	case 1:
		return ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine, data, decode)
	case 3:
		return ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor, data, alpha, decode)
	case 4:
		return ColorAtCMYK(x, y, width, data, decode)
	default:
		return nil, _ea.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063o\u006c\u006f\u0072\u0020\u0063\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", colorComponents)
	}
}
func _gcff(_gdc _c.Image) (Image, error) {
	if _afacd, _gffg := _gdc.(*NRGBA64); _gffg {
		return _afacd.Copy(), nil
	}
	_ddbb, _eeage, _fecb := _bdfe(_gdc, 2)
	_afbe, _acfe := NewImage(_ddbb.Max.X, _ddbb.Max.Y, 16, 3, nil, _fecb, nil)
	if _acfe != nil {
		return nil, _acfe
	}
	_gcagc(_gdc, _afbe, _ddbb)
	if len(_fecb) != 0 && !_eeage {
		if _egfa := _aaff(_fecb, _afbe); _egfa != nil {
			return nil, _egfa
		}
	}
	return _afbe, nil
}
func (_egfe *RGBA32) SetRGBA(x, y int, c _ag.RGBA) {
	_gdcf := y*_egfe.Width + x
	_dgec := 3 * _gdcf
	if _dgec+2 >= len(_egfe.Data) {
		return
	}
	_egfe.setRGBA(_gdcf, c)
}
func (_ace *Gray16) Base() *ImageBase { return &_ace.ImageBase }
func (_agf *CMYK32) Validate() error {
	if len(_agf.Data) != 4*_agf.Width*_agf.Height {
		return _a.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

type Monochrome struct {
	ImageBase
	ModelThreshold uint8
}

func (_cff *Monochrome) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _cff.Width, Y: _cff.Height}}
}
func (_cbbg *ImageBase) setEightFullBytes(_begd int, _efbb uint64) error {
	if _begd+7 > len(_cbbg.Data)-1 {
		return _a.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_cbbg.Data[_begd] = byte((_efbb & 0xff00000000000000) >> 56)
	_cbbg.Data[_begd+1] = byte((_efbb & 0xff000000000000) >> 48)
	_cbbg.Data[_begd+2] = byte((_efbb & 0xff0000000000) >> 40)
	_cbbg.Data[_begd+3] = byte((_efbb & 0xff00000000) >> 32)
	_cbbg.Data[_begd+4] = byte((_efbb & 0xff000000) >> 24)
	_cbbg.Data[_begd+5] = byte((_efbb & 0xff0000) >> 16)
	_cbbg.Data[_begd+6] = byte((_efbb & 0xff00) >> 8)
	_cbbg.Data[_begd+7] = byte(_efbb & 0xff)
	return nil
}
func (_gdg *Monochrome) ScaleLow(width, height int) (*Monochrome, error) {
	if width < 0 || height < 0 {
		return nil, _a.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064t\u0068\u0020\u0061\u006e\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_dcdd := _gea(width, height)
	_dabbg := make([]int, height)
	_aacd := make([]int, width)
	_bfbd := float64(_gdg.Width) / float64(width)
	_dgaf := float64(_gdg.Height) / float64(height)
	for _fcgd := 0; _fcgd < height; _fcgd++ {
		_dabbg[_fcgd] = int(_e.Min(_dgaf*float64(_fcgd)+0.5, float64(_gdg.Height-1)))
	}
	for _aaae := 0; _aaae < width; _aaae++ {
		_aacd[_aaae] = int(_e.Min(_bfbd*float64(_aaae)+0.5, float64(_gdg.Width-1)))
	}
	_ggec := -1
	_eeef := byte(0)
	for _eded := 0; _eded < height; _eded++ {
		_dea := _dabbg[_eded] * _gdg.BytesPerLine
		_fabb := _eded * _dcdd.BytesPerLine
		for _aefc := 0; _aefc < width; _aefc++ {
			_feaf := _aacd[_aefc]
			if _feaf != _ggec {
				_eeef = _gdg.getBit(_dea, _feaf)
				if _eeef != 0 {
					_dcdd.setBit(_fabb, _aefc)
				}
				_ggec = _feaf
			} else {
				if _eeef != 0 {
					_dcdd.setBit(_fabb, _aefc)
				}
			}
		}
	}
	return _dcdd, nil
}
func (_eedf *CMYK32) SetCMYK(x, y int, c _ag.CMYK) {
	_fef := 4 * (y*_eedf.Width + x)
	if _fef+3 >= len(_eedf.Data) {
		return
	}
	_eedf.Data[_fef] = c.C
	_eedf.Data[_fef+1] = c.M
	_eedf.Data[_fef+2] = c.Y
	_eedf.Data[_fef+3] = c.K
}
func _eccc(_abaf _ag.Color) _ag.Color {
	_bcagf := _ag.GrayModel.Convert(_abaf).(_ag.Gray)
	return _gagf(_bcagf)
}
func _eecb(_ebcg CMYK, _fbeba RGBA, _daga _c.Rectangle) {
	for _afda := 0; _afda < _daga.Max.X; _afda++ {
		for _ggege := 0; _ggege < _daga.Max.Y; _ggege++ {
			_cbcab := _ebcg.CMYKAt(_afda, _ggege)
			_fbeba.SetRGBA(_afda, _ggege, _fbe(_cbcab))
		}
	}
}
func (_fgea *NRGBA64) At(x, y int) _ag.Color { _cfba, _ := _fgea.ColorAt(x, y); return _cfba }
func (_cbed *Gray8) GrayAt(x, y int) _ag.Gray {
	_eecf, _ := ColorAtGray8BPC(x, y, _cbed.BytesPerLine, _cbed.Data, _cbed.Decode)
	return _eecf
}

type Gray2 struct{ ImageBase }

var _ Image = &Monochrome{}

func (_bcde *NRGBA32) SetNRGBA(x, y int, c _ag.NRGBA) {
	_gfeg := y*_bcde.Width + x
	_dedg := 3 * _gfeg
	if _dedg+2 >= len(_bcde.Data) {
		return
	}
	_bcde.setRGBA(_gfeg, c)
}

type SMasker interface {
	HasAlpha() bool
	GetAlpha() []byte
	MakeAlpha()
}

func (_dfbf *Gray8) Base() *ImageBase { return &_dfbf.ImageBase }
func RasterOperation(dest *Monochrome, dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _bgge(dest, dx, dy, dw, dh, op, src, sx, sy)
}
func (_dcbg *NRGBA64) Set(x, y int, c _ag.Color) {
	_cgc := (y*_dcbg.Width + x) * 2
	_cdbeb := _cgc * 3
	if _cdbeb+5 >= len(_dcbg.Data) {
		return
	}
	_fgff := _ag.NRGBA64Model.Convert(c).(_ag.NRGBA64)
	_dcbg.setNRGBA64(_cdbeb, _fgff, _cgc)
}
func (_abdc *Monochrome) setGray(_fgb int, _cbfg _ag.Gray, _dadd int) {
	if _cbfg.Y == 0 {
		_abdc.clearBit(_dadd, _fgb)
	} else {
		_abdc.setGrayBit(_dadd, _fgb)
	}
}

var _ Image = &CMYK32{}

type NRGBA32 struct{ ImageBase }
type CMYK32 struct{ ImageBase }

func BytesPerLine(width, bitsPerComponent, colorComponents int) int {
	return ((width*bitsPerComponent)*colorComponents + 7) >> 3
}
func (_afac *NRGBA16) Bounds() _c.Rectangle {
	return _c.Rectangle{Max: _c.Point{X: _afac.Width, Y: _afac.Height}}
}
func (_accg *Monochrome) Validate() error {
	if len(_accg.Data) != _accg.Height*_accg.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

var _ Image = &Gray4{}

func (_efge *Monochrome) Histogram() (_ccgb [256]int) {
	for _, _cbff := range _efge.Data {
		_ccgb[0xff] += int(_ggffa[_efge.Data[_cbff]])
	}
	return _ccgb
}
func ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor int, data, alpha []byte, decode []float64) (_ag.Color, error) {
	switch bitsPerColor {
	case 4:
		return ColorAtNRGBA16(x, y, width, bytesPerLine, data, alpha, decode)
	case 8:
		return ColorAtNRGBA32(x, y, width, data, alpha, decode)
	case 16:
		return ColorAtNRGBA64(x, y, width, data, alpha, decode)
	default:
		return nil, _ea.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0067\u0062\u0020b\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0061\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}

var _ Image = &NRGBA16{}
var _ Gray = &Gray8{}

func _cfaf(_eabac _c.Image, _ecgf Image, _eaef _c.Rectangle) {
	if _cbcd, _dfdd := _eabac.(SMasker); _dfdd && _cbcd.HasAlpha() {
		_ecgf.(SMasker).MakeAlpha()
	}
	switch _gbab := _eabac.(type) {
	case Gray:
		_dggc(_gbab, _ecgf.(NRGBA), _eaef)
	case NRGBA:
		_dedd(_gbab, _ecgf.(NRGBA), _eaef)
	case *_c.NYCbCrA:
		_ddga(_gbab, _ecgf.(NRGBA), _eaef)
	case CMYK:
		_gbdec(_gbab, _ecgf.(NRGBA), _eaef)
	case RGBA:
		_begcf(_gbab, _ecgf.(NRGBA), _eaef)
	case nrgba64:
		_ggdc(_gbab, _ecgf.(NRGBA), _eaef)
	default:
		_eceg(_eabac, _ecgf, _eaef)
	}
}
func _gfee(_cdgec _ag.Color) _ag.Color {
	_gabec := _ag.NRGBAModel.Convert(_cdgec).(_ag.NRGBA)
	return _aeea(_gabec)
}

type Gray4 struct{ ImageBase }

func ColorAtRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_ag.RGBA, error) {
	_acga := y*width + x
	_fbdb := 3 * _acga
	if _fbdb+2 >= len(data) {
		return _ag.RGBA{}, _ea.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_bedcg := uint8(0xff)
	if alpha != nil && len(alpha) > _acga {
		_bedcg = alpha[_acga]
	}
	_bgaf, _cbdgc, _efeg := data[_fbdb], data[_fbdb+1], data[_fbdb+2]
	if len(decode) == 6 {
		_bgaf = uint8(uint32(LinearInterpolate(float64(_bgaf), 0, 255, decode[0], decode[1])) & 0xff)
		_cbdgc = uint8(uint32(LinearInterpolate(float64(_cbdgc), 0, 255, decode[2], decode[3])) & 0xff)
		_efeg = uint8(uint32(LinearInterpolate(float64(_efeg), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _ag.RGBA{R: _bgaf, G: _cbdgc, B: _efeg, A: _bedcg}, nil
}
func IsGrayImgBlackAndWhite(i *_c.Gray) bool { return _ecca(i) }
func (_acfa *ImageBase) copy() ImageBase {
	_cefd := *_acfa
	_cefd.Data = make([]byte, len(_acfa.Data))
	copy(_cefd.Data, _acfa.Data)
	return _cefd
}
func (_dfac *RGBA32) Set(x, y int, c _ag.Color) {
	_deab := y*_dfac.Width + x
	_gcbc := 3 * _deab
	if _gcbc+2 >= len(_dfac.Data) {
		return
	}
	_aaedg := _ag.RGBAModel.Convert(c).(_ag.RGBA)
	_dfac.setRGBA(_deab, _aaedg)
}

var _ Gray = &Gray4{}

func (_afe *NRGBA32) NRGBAAt(x, y int) _ag.NRGBA {
	_bfgc, _ := ColorAtNRGBA32(x, y, _afe.Width, _afe.Data, _afe.Alpha, _afe.Decode)
	return _bfgc
}
func (_cbfed *NRGBA16) NRGBAAt(x, y int) _ag.NRGBA {
	_dcaba, _ := ColorAtNRGBA16(x, y, _cbfed.Width, _cbfed.BytesPerLine, _cbfed.Data, _cbfed.Alpha, _cbfed.Decode)
	return _dcaba
}

var _ _c.Image = &Gray2{}

func (_efgb *Monochrome) setIndexedBit(_gaab int) { _efgb.Data[(_gaab >> 3)] |= 0x80 >> uint(_gaab&7) }
func _abgg(_eegc _ag.Gray) _ag.NRGBA              { return _ag.NRGBA{R: _eegc.Y, G: _eegc.Y, B: _eegc.Y, A: 0xff} }
func _debg(_bgba _c.Image, _ccdf Image, _beeg _c.Rectangle) {
	if _fbce, _deeb := _bgba.(SMasker); _deeb && _fbce.HasAlpha() {
		_ccdf.(SMasker).MakeAlpha()
	}
	switch _bdab := _bgba.(type) {
	case Gray:
		_fgc(_bdab, _ccdf.(RGBA), _beeg)
	case NRGBA:
		_aeeb(_bdab, _ccdf.(RGBA), _beeg)
	case *_c.NYCbCrA:
		_fedg(_bdab, _ccdf.(RGBA), _beeg)
	case CMYK:
		_eecb(_bdab, _ccdf.(RGBA), _beeg)
	case RGBA:
		_fbfc(_bdab, _ccdf.(RGBA), _beeg)
	case nrgba64:
		_fdfb(_bdab, _ccdf.(RGBA), _beeg)
	default:
		_eceg(_bgba, _ccdf, _beeg)
	}
}
func _ecca(_fabad *_c.Gray) bool {
	for _aegbc := 0; _aegbc < len(_fabad.Pix); _aegbc++ {
		if !_bbcb(_fabad.Pix[_aegbc]) {
			return false
		}
	}
	return true
}
func (_cgdg *ImageBase) setTwoBytes(_cbg int, _fac uint16) error {
	if _cbg+1 > len(_cgdg.Data)-1 {
		return _a.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_cgdg.Data[_cbg] = byte((_fac & 0xff00) >> 8)
	_cgdg.Data[_cbg+1] = byte(_fac & 0xff)
	return nil
}
func ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine int, data []byte, decode []float64) (_ag.Color, error) {
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
		return nil, _ea.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0067\u0072\u0061\u0079\u0020\u0073c\u0061\u006c\u0065\u0020\u0062\u0069\u0074s\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020a\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}
func _cebbe(_ggga *_c.Gray, _ebac uint8) *_c.Gray {
	_abb := _ggga.Bounds()
	_efaaa := _c.NewGray(_abb)
	for _gbda := 0; _gbda < _abb.Dx(); _gbda++ {
		for _gace := 0; _gace < _abb.Dy(); _gace++ {
			_debe := _ggga.GrayAt(_gbda, _gace)
			_efaaa.SetGray(_gbda, _gace, _ag.Gray{Y: _aabf(_debe.Y, _ebac)})
		}
	}
	return _efaaa
}
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
	return nil, _ea.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0043o\u006e\u0076\u0065\u0072\u0074\u0065\u0072\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072\u0073\u002e\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003a\u0020\u0025\u0064\u002c\u0020\u0043\u006f\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u003a \u0025\u0064", bitsPerComponent, colorComponents)
}

type Histogramer interface{ Histogram() [256]int }

var _ NRGBA = &NRGBA16{}

func (_gfceg *Gray2) Set(x, y int, c _ag.Color) {
	if x >= _gfceg.Width || y >= _gfceg.Height {
		return
	}
	_dgcb := Gray2Model.Convert(c).(_ag.Gray)
	_bdeda := y * _gfceg.BytesPerLine
	_aace := _bdeda + (x >> 2)
	_dgad := _dgcb.Y >> 6
	_gfceg.Data[_aace] = (_gfceg.Data[_aace] & (^(0xc0 >> uint(2*((x)&3))))) | (_dgad << uint(6-2*(x&3)))
}
func _gea(_fc, _afc int) *Monochrome {
	return &Monochrome{ImageBase: NewImageBase(_fc, _afc, 1, 1, nil, nil, nil), ModelThreshold: 0x0f}
}
func _ggdc(_eaff nrgba64, _dfgc NRGBA, _bfgg _c.Rectangle) {
	for _efgg := 0; _efgg < _bfgg.Max.X; _efgg++ {
		for _gcac := 0; _gcac < _bfgg.Max.Y; _gcac++ {
			_fefc := _eaff.NRGBA64At(_efgg, _gcac)
			_dfgc.SetNRGBA(_efgg, _gcac, _fae(_fefc))
		}
	}
}
func (_fbfef *NRGBA32) At(x, y int) _ag.Color {
	_eae, _ := _fbfef.ColorAt(x, y)
	return _eae
}
func (_gaac *ImageBase) setByte(_ddb int, _ade byte) error {
	if _ddb > len(_gaac.Data)-1 {
		return _a.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_gaac.Data[_ddb] = _ade
	return nil
}
func _fgc(_afbd Gray, _ccfa RGBA, _aabee _c.Rectangle) {
	for _bega := 0; _bega < _aabee.Max.X; _bega++ {
		for _bfd := 0; _bfd < _aabee.Max.Y; _bfd++ {
			_cbcb := _afbd.GrayAt(_bega, _bfd)
			_ccfa.SetRGBA(_bega, _bfd, _aab(_cbcb))
		}
	}
}
func _fbfc(_adgb, _ecbb RGBA, _cfbf _c.Rectangle) {
	for _bfbe := 0; _bfbe < _cfbf.Max.X; _bfbe++ {
		for _cebf := 0; _cebf < _cfbf.Max.Y; _cebf++ {
			_ecbb.SetRGBA(_bfbe, _cebf, _adgb.RGBAAt(_bfbe, _cebf))
		}
	}
}

type RGBA32 struct{ ImageBase }

func (_dgbb *Gray8) Set(x, y int, c _ag.Color) {
	_ecbf := y*_dgbb.BytesPerLine + x
	if _ecbf > len(_dgbb.Data)-1 {
		return
	}
	_cgde := _ag.GrayModel.Convert(c)
	_dgbb.Data[_ecbf] = _cgde.(_ag.Gray).Y
}
func _bbf(_caa NRGBA, _cedc CMYK, _babe _c.Rectangle) {
	for _gae := 0; _gae < _babe.Max.X; _gae++ {
		for _fbgd := 0; _fbgd < _babe.Max.Y; _fbgd++ {
			_fge := _caa.NRGBAAt(_gae, _fbgd)
			_cedc.SetCMYK(_gae, _fbgd, _agdc(_fge))
		}
	}
}
func (_eedc *NRGBA64) Copy() Image { return &NRGBA64{ImageBase: _eedc.copy()} }
func (_gfff *NRGBA64) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtNRGBA64(x, y, _gfff.Width, _gfff.Data, _gfff.Alpha, _gfff.Decode)
}
func _dga(_bea _ag.NYCbCrA) _ag.NRGBA {
	_bded := int32(_bea.Y) * 0x10101
	_ffgda := int32(_bea.Cb) - 128
	_cdg := int32(_bea.Cr) - 128
	_eedb := _bded + 91881*_cdg
	if uint32(_eedb)&0xff000000 == 0 {
		_eedb >>= 8
	} else {
		_eedb = ^(_eedb >> 31) & 0xffff
	}
	_bdg := _bded - 22554*_ffgda - 46802*_cdg
	if uint32(_bdg)&0xff000000 == 0 {
		_bdg >>= 8
	} else {
		_bdg = ^(_bdg >> 31) & 0xffff
	}
	_aecf := _bded + 116130*_ffgda
	if uint32(_aecf)&0xff000000 == 0 {
		_aecf >>= 8
	} else {
		_aecf = ^(_aecf >> 31) & 0xffff
	}
	return _ag.NRGBA{R: uint8(_eedb >> 8), G: uint8(_bdg >> 8), B: uint8(_aecf >> 8), A: _bea.A}
}

var _ggffa [256]uint8

func _beed(_add *Monochrome, _ecg ...int) (_bca *Monochrome, _gbf error) {
	if _add == nil {
		return nil, _a.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_ecg) == 0 {
		return nil, _a.New("\u0074h\u0065\u0072e\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020o\u006e\u0065\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006f\u0066 \u0072\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e")
	}
	_ead := _gfd()
	_bca = _add
	for _, _gab := range _ecg {
		if _gab <= 0 {
			break
		}
		_bca, _gbf = _dcd(_bca, _gab, _ead)
		if _gbf != nil {
			return nil, _gbf
		}
	}
	return _bca, nil
}
func (_caef *NRGBA32) Set(x, y int, c _ag.Color) {
	_fbdag := y*_caef.Width + x
	_gagb := 3 * _fbdag
	if _gagb+2 >= len(_caef.Data) {
		return
	}
	_dfecd := _ag.NRGBAModel.Convert(c).(_ag.NRGBA)
	_caef.setRGBA(_fbdag, _dfecd)
}
func (_edde *NRGBA16) SetNRGBA(x, y int, c _ag.NRGBA) {
	_cecc := y*_edde.BytesPerLine + x*3/2
	if _cecc+1 >= len(_edde.Data) {
		return
	}
	c = _aeea(c)
	_edde.setNRGBA(x, y, _cecc, c)
}

var _ _c.Image = &NRGBA64{}

func (_eea *Monochrome) Base() *ImageBase { return &_eea.ImageBase }
func (_cdad *CMYK32) CMYKAt(x, y int) _ag.CMYK {
	_befg, _ := ColorAtCMYK(x, y, _cdad.Width, _cdad.Data, _cdad.Decode)
	return _befg
}
func _bbe(_dcfa int, _abad int) int {
	if _dcfa < _abad {
		return _dcfa
	}
	return _abad
}
func (_dba *Gray16) ColorModel() _ag.Model { return _ag.Gray16Model }
func (_dbbg *Gray4) At(x, y int) _ag.Color {
	_dgcbd, _ := _dbbg.ColorAt(x, y)
	return _dgcbd
}
func _cfbg(_ggce _ag.Color) _ag.Color {
	_cebe := _ag.GrayModel.Convert(_ggce).(_ag.Gray)
	return _def(_cebe)
}

var _ _c.Image = &Gray4{}

func (_dcae *Gray4) ColorAt(x, y int) (_ag.Color, error) {
	return ColorAtGray4BPC(x, y, _dcae.BytesPerLine, _dcae.Data, _dcae.Decode)
}
func _gfce(_ecb RGBA, _eagc CMYK, _geb _c.Rectangle) {
	for _ebc := 0; _ebc < _geb.Max.X; _ebc++ {
		for _gbg := 0; _gbg < _geb.Max.Y; _gbg++ {
			_bgf := _ecb.RGBAAt(_ebc, _gbg)
			_eagc.SetCMYK(_ebc, _gbg, _cbag(_bgf))
		}
	}
}
func _gagf(_cbca _ag.Gray) _ag.Gray {
	_cbca.Y >>= 4
	_cbca.Y |= _cbca.Y << 4
	return _cbca
}

type Gray interface {
	GrayAt(_abd, _egf int) _ag.Gray
	SetGray(_dccc, _edfg int, _cbagc _ag.Gray)
}
type nrgba64 interface {
	NRGBA64At(_bbde, _aagb int) _ag.NRGBA64
	SetNRGBA64(_ddfgd, _aga int, _aegb _ag.NRGBA64)
}

func _gge(_ga, _ebab *Monochrome) (_aae error) {
	_bef := _ebab.BytesPerLine
	_fbb := _ga.BytesPerLine
	var _ff, _ac, _eec, _ceca, _bcc int
	for _eec = 0; _eec < _ebab.Height; _eec++ {
		_ff = _eec * _bef
		_ac = 8 * _eec * _fbb
		for _ceca = 0; _ceca < _bef; _ceca++ {
			if _aae = _ga.setEightBytes(_ac+_ceca*8, _gdb[_ebab.Data[_ff+_ceca]]); _aae != nil {
				return _aae
			}
		}
		for _bcc = 1; _bcc < 8; _bcc++ {
			for _ceca = 0; _ceca < _fbb; _ceca++ {
				if _aae = _ga.setByte(_ac+_bcc*_fbb+_ceca, _ga.Data[_ac+_ceca]); _aae != nil {
					return _aae
				}
			}
		}
	}
	return nil
}
func _bac(_abfa _c.Image) (Image, error) {
	if _caedf, _eafc := _abfa.(*NRGBA32); _eafc {
		return _caedf.Copy(), nil
	}
	_gffd, _eggf, _fbae := _bdfe(_abfa, 1)
	_daeg, _gbec := NewImage(_gffd.Max.X, _gffd.Max.Y, 8, 3, nil, _fbae, nil)
	if _gbec != nil {
		return nil, _gbec
	}
	_cfaf(_abfa, _daeg, _gffd)
	if len(_fbae) != 0 && !_eggf {
		if _fdfe := _aaff(_fbae, _daeg); _fdfe != nil {
			return nil, _fdfe
		}
	}
	return _daeg, nil
}
func (_fdgb *Gray2) SetGray(x, y int, gray _ag.Gray) {
	_dcba := _def(gray)
	_dgf := y * _fdgb.BytesPerLine
	_bdeb := _dgf + (x >> 2)
	if _bdeb >= len(_fdgb.Data) {
		return
	}
	_dcee := _dcba.Y >> 6
	_fdgb.Data[_bdeb] = (_fdgb.Data[_bdeb] & (^(0xc0 >> uint(2*((x)&3))))) | (_dcee << uint(6-2*(x&3)))
}
func _aeag() {
	for _eagcb := 0; _eagcb < 256; _eagcb++ {
		_ggffa[_eagcb] = uint8(_eagcb&0x1) + (uint8(_eagcb>>1) & 0x1) + (uint8(_eagcb>>2) & 0x1) + (uint8(_eagcb>>3) & 0x1) + (uint8(_eagcb>>4) & 0x1) + (uint8(_eagcb>>5) & 0x1) + (uint8(_eagcb>>6) & 0x1) + (uint8(_eagcb>>7) & 0x1)
	}
}
func _edc(_af int) []uint {
	var _cf []uint
	_cfb := _af
	_eg := _cfb / 8
	if _eg != 0 {
		for _ffe := 0; _ffe < _eg; _ffe++ {
			_cf = append(_cf, 8)
		}
		_bdf := _cfb % 8
		_cfb = 0
		if _bdf != 0 {
			_cfb = _bdf
		}
	}
	_cdbc := _cfb / 4
	if _cdbc != 0 {
		for _ebe := 0; _ebe < _cdbc; _ebe++ {
			_cf = append(_cf, 4)
		}
		_edd := _cfb % 4
		_cfb = 0
		if _edd != 0 {
			_cfb = _edd
		}
	}
	_gbc := _cfb / 2
	if _gbc != 0 {
		for _bab := 0; _bab < _gbc; _bab++ {
			_cf = append(_cf, 2)
		}
	}
	return _cf
}
