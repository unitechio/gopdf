package imageutil

import (
	_cd "encoding/binary"
	_cb "errors"
	_a "fmt"
	_f "image"
	_g "image/color"
	_cg "image/draw"
	_e "math"

	_ag "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
)

func (_dcdfg *NRGBA16) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _dcdfg.Width, Y: _dcdfg.Height}}
}
func _edcf(_gacg _f.Image) (Image, error) {
	if _dfaa, _daae := _gacg.(*NRGBA32); _daae {
		return _dfaa.Copy(), nil
	}
	_feaf, _dfbf, _cbfad := _ddbfg(_gacg, 1)
	_bafaa, _cgga := NewImage(_feaf.Max.X, _feaf.Max.Y, 8, 3, nil, _cbfad, nil)
	if _cgga != nil {
		return nil, _cgga
	}
	_decfg(_gacg, _bafaa, _feaf)
	if len(_cbfad) != 0 && !_dfbf {
		if _bgc := _cfd(_cbfad, _bafaa); _bgc != nil {
			return nil, _bgc
		}
	}
	return _bafaa, nil
}
func (_feag *Gray4) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _feag.Width, Y: _feag.Height}}
}
func _debgb(_agcb int, _cabf int) int {
	if _agcb < _cabf {
		return _agcb
	}
	return _cabf
}
func ConverterFunc(converterFunc func(_fdcb _f.Image) (Image, error)) ColorConverter {
	return colorConverter{_fcg: converterFunc}
}
func (_ffdd *Gray4) GrayAt(x, y int) _g.Gray {
	_dcec, _ := ColorAtGray4BPC(x, y, _ffdd.BytesPerLine, _ffdd.Data, _ffdd.Decode)
	return _dcec
}

var (
	Gray2Model   = _g.ModelFunc(_cggc)
	Gray4Model   = _g.ModelFunc(_ccb)
	NRGBA16Model = _g.ModelFunc(_agfb)
)
var _ Gray = &Monochrome{}

func (_bcfe *CMYK32) At(x, y int) _g.Color { _fcb, _ := _bcfe.ColorAt(x, y); return _fcb }
func (_bffa *Gray8) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _bffa.Width, Y: _bffa.Height}}
}
func (_dee *NRGBA16) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtNRGBA16(x, y, _dee.Width, _dee.BytesPerLine, _dee.Data, _dee.Alpha, _dee.Decode)
}
func BytesPerLine(width, bitsPerComponent, colorComponents int) int {
	return ((width*bitsPerComponent)*colorComponents + 7) >> 3
}
func (_aba *CMYK32) Copy() Image { return &CMYK32{ImageBase: _aba.copy()} }
func _bbb(_fdeb, _aab *Monochrome, _aad []byte, _fgg int) (_gca error) {
	var (
		_ggg, _fgef, _egd, _bfa, _deg, _dce, _gdd, _cca int
		_ecgb, _fce, _bggf, _edf                        uint32
		_efc, _cab                                      byte
		_dcd                                            uint16
	)
	_fdf := make([]byte, 4)
	_fag := make([]byte, 4)
	for _egd = 0; _egd < _fdeb.Height-1; _egd, _bfa = _egd+2, _bfa+1 {
		_ggg = _egd * _fdeb.BytesPerLine
		_fgef = _bfa * _aab.BytesPerLine
		for _deg, _dce = 0, 0; _deg < _fgg; _deg, _dce = _deg+4, _dce+1 {
			for _gdd = 0; _gdd < 4; _gdd++ {
				_cca = _ggg + _deg + _gdd
				if _cca <= len(_fdeb.Data)-1 && _cca < _ggg+_fdeb.BytesPerLine {
					_fdf[_gdd] = _fdeb.Data[_cca]
				} else {
					_fdf[_gdd] = 0x00
				}
				_cca = _ggg + _fdeb.BytesPerLine + _deg + _gdd
				if _cca <= len(_fdeb.Data)-1 && _cca < _ggg+(2*_fdeb.BytesPerLine) {
					_fag[_gdd] = _fdeb.Data[_cca]
				} else {
					_fag[_gdd] = 0x00
				}
			}
			_ecgb = _cd.BigEndian.Uint32(_fdf)
			_fce = _cd.BigEndian.Uint32(_fag)
			_bggf = _ecgb & _fce
			_bggf |= _bggf << 1
			_edf = _ecgb | _fce
			_edf &= _edf << 1
			_fce = _bggf | _edf
			_fce &= 0xaaaaaaaa
			_ecgb = _fce | (_fce << 7)
			_efc = byte(_ecgb >> 24)
			_cab = byte((_ecgb >> 8) & 0xff)
			_cca = _fgef + _dce
			if _cca+1 == len(_aab.Data)-1 || _cca+1 >= _fgef+_aab.BytesPerLine {
				if _gca = _aab.setByte(_cca, _aad[_efc]); _gca != nil {
					return _a.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _cca)
				}
			} else {
				_dcd = (uint16(_aad[_efc]) << 8) | uint16(_aad[_cab])
				if _gca = _aab.setTwoBytes(_cca, _dcd); _gca != nil {
					return _a.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _cca)
				}
				_dce++
			}
		}
	}
	return nil
}
func (_gdfc *ImageBase) newAlpha() {
	_effdf := BytesPerLine(_gdfc.Width, _gdfc.BitsPerComponent, 1)
	_gdfc.Alpha = make([]byte, _gdfc.Height*_effdf)
}
func _baa(_fba _g.RGBA) _g.CMYK {
	_dgba, _bcfa, _fcbd, _dec := _g.RGBToCMYK(_fba.R, _fba.G, _fba.B)
	return _g.CMYK{C: _dgba, M: _bcfa, Y: _fcbd, K: _dec}
}
func _ecag(_dbca uint) uint {
	var _agcd uint
	for _dbca != 0 {
		_dbca >>= 1
		_agcd++
	}
	return _agcd - 1
}
func _caf() (_gefb []byte) {
	_gefb = make([]byte, 256)
	for _gfcc := 0; _gfcc < 256; _gfcc++ {
		_cggf := byte(_gfcc)
		_gefb[_cggf] = (_cggf & 0x01) | ((_cggf & 0x04) >> 1) | ((_cggf & 0x10) >> 2) | ((_cggf & 0x40) >> 3) | ((_cggf & 0x02) << 3) | ((_cggf & 0x08) << 2) | ((_cggf & 0x20) << 1) | (_cggf & 0x80)
	}
	return _gefb
}
func _fdbdc(_cfae _g.Gray) _g.Gray {
	_cfae.Y >>= 4
	_cfae.Y |= _cfae.Y << 4
	return _cfae
}
func _cadeee(_fcae *_f.Gray) bool {
	for _feae := 0; _feae < len(_fcae.Pix); _feae++ {
		if !_addaf(_fcae.Pix[_feae]) {
			return false
		}
	}
	return true
}
func (_cbga *RGBA32) SetRGBA(x, y int, c _g.RGBA) {
	_eeeg := y*_cbga.Width + x
	_cffa := 3 * _eeeg
	if _cffa+2 >= len(_cbga.Data) {
		return
	}
	_cbga.setRGBA(_eeeg, c)
}
func (_dgc *Gray16) Copy() Image            { return &Gray16{ImageBase: _dgc.copy()} }
func (_gaaa *NRGBA64) At(x, y int) _g.Color { _gagc, _ := _gaaa.ColorAt(x, y); return _gagc }

var _ _f.Image = &Gray2{}

func ColorAtNRGBA64(x, y, width int, data, alpha []byte, decode []float64) (_g.NRGBA64, error) {
	_bafc := (y*width + x) * 2
	_fbce := _bafc * 3
	if _fbce+5 >= len(data) {
		return _g.NRGBA64{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	const _fcgb = 0xffff
	_fbdf := uint16(_fcgb)
	if alpha != nil && len(alpha) > _bafc+1 {
		_fbdf = uint16(alpha[_bafc])<<8 | uint16(alpha[_bafc+1])
	}
	_aegd := uint16(data[_fbce])<<8 | uint16(data[_fbce+1])
	_feaaf := uint16(data[_fbce+2])<<8 | uint16(data[_fbce+3])
	_fegg := uint16(data[_fbce+4])<<8 | uint16(data[_fbce+5])
	if len(decode) == 6 {
		_aegd = uint16(uint64(LinearInterpolate(float64(_aegd), 0, 65535, decode[0], decode[1])) & _fcgb)
		_feaaf = uint16(uint64(LinearInterpolate(float64(_feaaf), 0, 65535, decode[2], decode[3])) & _fcgb)
		_fegg = uint16(uint64(LinearInterpolate(float64(_fegg), 0, 65535, decode[4], decode[5])) & _fcgb)
	}
	return _g.NRGBA64{R: _aegd, G: _feaaf, B: _fegg, A: _fbdf}, nil
}
func _gggf(_fcgd _f.Image) (Image, error) {
	if _cgfc, _cbdce := _fcgd.(*Monochrome); _cbdce {
		return _cgfc, nil
	}
	_cbfc := _fcgd.Bounds()
	var _dgd Gray
	switch _bfcd := _fcgd.(type) {
	case Gray:
		_dgd = _bfcd
	case NRGBA:
		_dgd = &Gray8{ImageBase: NewImageBase(_cbfc.Max.X, _cbfc.Max.Y, 8, 1, nil, nil, nil)}
		_adb(_dgd, _bfcd, _cbfc)
	case nrgba64:
		_dgd = &Gray8{ImageBase: NewImageBase(_cbfc.Max.X, _cbfc.Max.Y, 8, 1, nil, nil, nil)}
		_egace(_dgd, _bfcd, _cbfc)
	default:
		_baad, _fdb := GrayConverter.Convert(_fcgd)
		if _fdb != nil {
			return nil, _fdb
		}
		_dgd = _baad.(Gray)
	}
	_bdg, _edfc := NewImage(_cbfc.Max.X, _cbfc.Max.Y, 1, 1, nil, nil, nil)
	if _edfc != nil {
		return nil, _edfc
	}
	_agg := _bdg.(*Monochrome)
	_fccf := AutoThresholdTriangle(GrayHistogram(_dgd))
	for _bggg := 0; _bggg < _cbfc.Max.X; _bggg++ {
		for _bbbg := 0; _bbbg < _cbfc.Max.Y; _bbbg++ {
			_afcg := _fda(_dgd.GrayAt(_bggg, _bbbg), monochromeModel(_fccf))
			_agg.SetGray(_bggg, _bbbg, _afcg)
		}
	}
	return _bdg, nil
}
func (_cfeae *ImageBase) GetAlpha() []byte { return _cfeae.Alpha }
func _ebc(_fge *Monochrome, _ded ...int) (_gdg *Monochrome, _ecb error) {
	if _fge == nil {
		return nil, _cb.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_ded) == 0 {
		return nil, _cb.New("\u0074h\u0065\u0072e\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020o\u006e\u0065\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006f\u0066 \u0072\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e")
	}
	_aae := _caf()
	_gdg = _fge
	for _, _eggb := range _ded {
		if _eggb <= 0 {
			break
		}
		_gdg, _ecb = _dbab(_gdg, _eggb, _aae)
		if _ecb != nil {
			return nil, _ecb
		}
	}
	return _gdg, nil
}
func _acdaa(_cbfce Gray, _cfbca NRGBA, _dadb _f.Rectangle) {
	for _bccbd := 0; _bccbd < _dadb.Max.X; _bccbd++ {
		for _gfbg := 0; _gfbg < _dadb.Max.Y; _gfbg++ {
			_dfcbc := _cbfce.GrayAt(_bccbd, _gfbg)
			_cfbca.SetNRGBA(_bccbd, _gfbg, _bacd(_dfcbc))
		}
	}
}
func (_fadbg *Gray4) Validate() error {
	if len(_fadbg.Data) != _fadbg.Height*_fadbg.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_cfgb *Monochrome) Histogram() (_eeea [256]int) {
	for _, _affc := range _cfgb.Data {
		_eeea[0xff] += int(_bcde[_cfgb.Data[_affc]])
	}
	return _eeea
}
func _gbeb(_egde _f.Image, _deaa uint8) *_f.Gray {
	_dbccg := _egde.Bounds()
	_fbea := _f.NewGray(_dbccg)
	var (
		_bbbe  _g.Color
		_dedad _g.Gray
	)
	for _agbe := 0; _agbe < _dbccg.Max.X; _agbe++ {
		for _edbbdg := 0; _edbbdg < _dbccg.Max.Y; _edbbdg++ {
			_bbbe = _egde.At(_agbe, _edbbdg)
			_fbea.Set(_agbe, _edbbdg, _bbbe)
			_dedad = _fbea.GrayAt(_agbe, _edbbdg)
			_fbea.SetGray(_agbe, _edbbdg, _g.Gray{Y: _afdb(_dedad.Y, _deaa)})
		}
	}
	return _fbea
}
func FromGoImage(i _f.Image) (Image, error) {
	switch _eag := i.(type) {
	case Image:
		return _eag.Copy(), nil
	case Gray:
		return GrayConverter.Convert(i)
	case *_f.Gray16:
		return Gray16Converter.Convert(i)
	case CMYK:
		return CMYKConverter.Convert(i)
	case *_f.NRGBA64:
		return NRGBA64Converter.Convert(i)
	default:
		return NRGBAConverter.Convert(i)
	}
}
func ColorAtGray1BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_gfed := y*bytesPerLine + x>>3
	if _gfed >= len(data) {
		return _g.Gray{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_bddf := data[_gfed] >> uint(7-(x&7)) & 1
	if len(decode) == 2 {
		_bddf = uint8(LinearInterpolate(float64(_bddf), 0.0, 1.0, decode[0], decode[1])) & 1
	}
	return _g.Gray{Y: _bddf * 255}, nil
}
func (_cdcf *NRGBA64) SetNRGBA64(x, y int, c _g.NRGBA64) {
	_fdgfc := (y*_cdcf.Width + x) * 2
	_edbf := _fdgfc * 3
	if _edbf+5 >= len(_cdcf.Data) {
		return
	}
	_cdcf.setNRGBA64(_edbf, c, _fdgfc)
}

var _bcde [256]uint8

func _addaf(_gfdc uint8) bool {
	if _gfdc == 0 || _gfdc == 255 {
		return true
	}
	return false
}
func _fecf(_faag _f.Image) (Image, error) {
	if _cegc, _cege := _faag.(*Gray16); _cege {
		return _cegc.Copy(), nil
	}
	_cea := _faag.Bounds()
	_cbdg, _acbg := NewImage(_cea.Max.X, _cea.Max.Y, 16, 1, nil, nil, nil)
	if _acbg != nil {
		return nil, _acbg
	}
	_gceff(_faag, _cbdg, _cea)
	return _cbdg, nil
}

var _ Image = &Gray8{}

type Image interface {
	_cg.Image
	Base() *ImageBase
	Copy() Image
	Pix() []byte
	ColorAt(_aggc, _dff int) (_g.Color, error)
	Validate() error
}

func (_gfcg *Monochrome) GrayAt(x, y int) _g.Gray {
	_cbgd, _ := ColorAtGray1BPC(x, y, _gfcg.BytesPerLine, _gfcg.Data, _gfcg.Decode)
	return _cbgd
}
func _cgb(_gcef RGBA, _feb CMYK, _dbd _f.Rectangle) {
	for _eba := 0; _eba < _dbd.Max.X; _eba++ {
		for _fad := 0; _fad < _dbd.Max.Y; _fad++ {
			_dab := _gcef.RGBAAt(_eba, _fad)
			_feb.SetCMYK(_eba, _fad, _baa(_dab))
		}
	}
}
func _dedc(_daeg CMYK, _fbag RGBA, _egacc _f.Rectangle) {
	for _ggga := 0; _ggga < _egacc.Max.X; _ggga++ {
		for _geccb := 0; _geccb < _egacc.Max.Y; _geccb++ {
			_fbdef := _daeg.CMYKAt(_ggga, _geccb)
			_fbag.SetRGBA(_ggga, _geccb, _fbd(_fbdef))
		}
	}
}
func ImgToBinary(i _f.Image, threshold uint8) *_f.Gray {
	switch _dfdcc := i.(type) {
	case *_f.Gray:
		if _cadeee(_dfdcc) {
			return _dfdcc
		}
		return _bacda(_dfdcc, threshold)
	case *_f.Gray16:
		return _defac(_dfdcc, threshold)
	default:
		return _gbeb(_dfdcc, threshold)
	}
}
func _cc(_fegc, _dfc int, _bbe []byte) *Monochrome {
	_dbabc := _bd(_fegc, _dfc)
	_dbabc.Data = _bbe
	return _dbabc
}
func (_gfcf *Gray16) SetGray(x, y int, g _g.Gray) {
	_cgca := (y*_gfcf.BytesPerLine/2 + x) * 2
	if _cgca+1 >= len(_gfcf.Data) {
		return
	}
	_gfcf.Data[_cgca] = g.Y
	_gfcf.Data[_cgca+1] = g.Y
}
func (_fagd *RGBA32) Copy() Image        { return &RGBA32{ImageBase: _fagd.copy()} }
func (_dbac *ImageBase) Pix() []byte     { return _dbac.Data }
func (_bdaec *NRGBA16) Base() *ImageBase { return &_bdaec.ImageBase }

type Gray interface {
	GrayAt(_acaf, _ddbb int) _g.Gray
	SetGray(_agee, _gcefc int, _cggfg _g.Gray)
}

func (_dbdf *Gray8) ColorModel() _g.Model { return _g.GrayModel }
func (_afa *CMYK32) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtCMYK(x, y, _afa.Width, _afa.Data, _afa.Decode)
}
func _ecdg(_geee *Monochrome, _ggdg, _ffdc, _dbee, _baaa int, _cfceb RasterOperator, _dbbbd *Monochrome, _ffcd, _aacd int) error {
	var (
		_fbf         bool
		_cgaa        bool
		_cffea       byte
		_agb         int
		_fbbe        int
		_bbaa        int
		_bcbf        int
		_eaf         bool
		_fgfca       int
		_cfef        int
		_gcde        int
		_aaag        bool
		_bggff       byte
		_ggfd        int
		_geaae       int
		_ffcc        int
		_cadee       byte
		_cge         int
		_bfcc        int
		_cfedd       uint
		_agbc        uint
		_gace        byte
		_dead        shift
		_fage        bool
		_bagc        bool
		_dffg, _aece int
	)
	if _ffcd&7 != 0 {
		_bfcc = 8 - (_ffcd & 7)
	}
	if _ggdg&7 != 0 {
		_fbbe = 8 - (_ggdg & 7)
	}
	if _bfcc == 0 && _fbbe == 0 {
		_gace = _ebgf[0]
	} else {
		if _fbbe > _bfcc {
			_cfedd = uint(_fbbe - _bfcc)
		} else {
			_cfedd = uint(8 - (_bfcc - _fbbe))
		}
		_agbc = 8 - _cfedd
		_gace = _ebgf[_cfedd]
	}
	if (_ggdg & 7) != 0 {
		_fbf = true
		_agb = 8 - (_ggdg & 7)
		_cffea = _ebgf[_agb]
		_bbaa = _geee.BytesPerLine*_ffdc + (_ggdg >> 3)
		_bcbf = _dbbbd.BytesPerLine*_aacd + (_ffcd >> 3)
		_cge = 8 - (_ffcd & 7)
		if _agb > _cge {
			_dead = _gafc
			if _dbee >= _bfcc {
				_fage = true
			}
		} else {
			_dead = _cfaee
		}
	}
	if _dbee < _agb {
		_cgaa = true
		_cffea &= _dcdfd[8-_agb+_dbee]
	}
	if !_cgaa {
		_fgfca = (_dbee - _agb) >> 3
		if _fgfca != 0 {
			_eaf = true
			_cfef = _geee.BytesPerLine*_ffdc + ((_ggdg + _fbbe) >> 3)
			_gcde = _dbbbd.BytesPerLine*_aacd + ((_ffcd + _fbbe) >> 3)
		}
	}
	_ggfd = (_ggdg + _dbee) & 7
	if !(_cgaa || _ggfd == 0) {
		_aaag = true
		_bggff = _dcdfd[_ggfd]
		_geaae = _geee.BytesPerLine*_ffdc + ((_ggdg + _fbbe) >> 3) + _fgfca
		_ffcc = _dbbbd.BytesPerLine*_aacd + ((_ffcd + _fbbe) >> 3) + _fgfca
		if _ggfd > int(_agbc) {
			_bagc = true
		}
	}
	switch _cfceb {
	case PixSrc:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], _cadee, _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] = _cadee
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], _cadee, _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixNotSrc:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], ^_cadee, _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] = ^_cadee
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], ^_cadee, _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixSrcOrDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], _cadee|_geee.Data[_bbaa], _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] |= _cadee
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], _cadee|_geee.Data[_geaae], _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixSrcAndDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], _cadee&_geee.Data[_bbaa], _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] &= _cadee
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], _cadee&_geee.Data[_geaae], _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixSrcXorDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], _cadee^_geee.Data[_bbaa], _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] ^= _cadee
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], _cadee^_geee.Data[_geaae], _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], ^_cadee|_geee.Data[_bbaa], _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] |= ^_cadee
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], ^_cadee|_geee.Data[_geaae], _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], ^_cadee&_geee.Data[_bbaa], _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] &= ^_cadee
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], ^_cadee&_geee.Data[_geaae], _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], _cadee|^_geee.Data[_bbaa], _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] = _cadee | ^_geee.Data[_cfef+_aece]
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], _cadee|^_geee.Data[_geaae], _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], _cadee&^_geee.Data[_bbaa], _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] = _cadee &^ _geee.Data[_cfef+_aece]
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], _cadee&^_geee.Data[_geaae], _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], ^(_cadee | _geee.Data[_bbaa]), _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] = ^(_cadee | _geee.Data[_cfef+_aece])
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], ^(_cadee | _geee.Data[_geaae]), _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], ^(_cadee & _geee.Data[_bbaa]), _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] = ^(_cadee & _geee.Data[_cfef+_aece])
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], ^(_cadee & _geee.Data[_geaae]), _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		if _fbf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				if _dead == _gafc {
					_cadee = _dbbbd.Data[_bcbf] << _cfedd
					if _fage {
						_cadee = _dacg(_cadee, _dbbbd.Data[_bcbf+1]>>_agbc, _gace)
					}
				} else {
					_cadee = _dbbbd.Data[_bcbf] >> _agbc
				}
				_geee.Data[_bbaa] = _dacg(_geee.Data[_bbaa], ^(_cadee ^ _geee.Data[_bbaa]), _cffea)
				_bbaa += _geee.BytesPerLine
				_bcbf += _dbbbd.BytesPerLine
			}
		}
		if _eaf {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				for _aece = 0; _aece < _fgfca; _aece++ {
					_cadee = _dacg(_dbbbd.Data[_gcde+_aece]<<_cfedd, _dbbbd.Data[_gcde+_aece+1]>>_agbc, _gace)
					_geee.Data[_cfef+_aece] = ^(_cadee ^ _geee.Data[_cfef+_aece])
				}
				_cfef += _geee.BytesPerLine
				_gcde += _dbbbd.BytesPerLine
			}
		}
		if _aaag {
			for _dffg = 0; _dffg < _baaa; _dffg++ {
				_cadee = _dbbbd.Data[_ffcc] << _cfedd
				if _bagc {
					_cadee = _dacg(_cadee, _dbbbd.Data[_ffcc+1]>>_agbc, _gace)
				}
				_geee.Data[_geaae] = _dacg(_geee.Data[_geaae], ^(_cadee ^ _geee.Data[_geaae]), _bggff)
				_geaae += _geee.BytesPerLine
				_ffcc += _dbbbd.BytesPerLine
			}
		}
	default:
		_ag.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _cfceb)
		return _cb.New("\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}

type colorConverter struct {
	_fcg func(_efcb _f.Image) (Image, error)
}

func _fefe(_fegbff _f.Image) (Image, error) {
	if _caac, _dbfeg := _fegbff.(*NRGBA64); _dbfeg {
		return _caac.Copy(), nil
	}
	_ecga, _egdca, _fbec := _ddbfg(_fegbff, 2)
	_dadg, _fcgec := NewImage(_ecga.Max.X, _ecga.Max.Y, 16, 3, nil, _fbec, nil)
	if _fcgec != nil {
		return nil, _fcgec
	}
	_ageec(_fegbff, _dadg, _ecga)
	if len(_fbec) != 0 && !_egdca {
		if _bgbf := _cfd(_fbec, _dadg); _bgbf != nil {
			return nil, _bgbf
		}
	}
	return _dadg, nil
}
func IsPowerOf2(n uint) bool { return n > 0 && (n&(n-1)) == 0 }
func _baab(_dfac RGBA, _cfcea NRGBA, _dbfc _f.Rectangle) {
	for _cgeba := 0; _cgeba < _dbfc.Max.X; _cgeba++ {
		for _adee := 0; _adee < _dbfc.Max.Y; _adee++ {
			_dfeb := _dfac.RGBAAt(_cgeba, _adee)
			_cfcea.SetNRGBA(_cgeba, _adee, _bdda(_dfeb))
		}
	}
}

var _ Image = &Gray4{}

func (_gabe *ImageBase) setByte(_fded int, _cac byte) error {
	if _fded > len(_gabe.Data)-1 {
		return _cb.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_gabe.Data[_fded] = _cac
	return nil
}
func (_fgf *CMYK32) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _fgf.Width, Y: _fgf.Height}}
}
func (_fbcg *Monochrome) setGrayBit(_feedb, _gda int) { _fbcg.Data[_feedb] |= 0x80 >> uint(_gda&7) }
func _dfbd(_dbff _f.Image) (Image, error) {
	if _adfa, _bfca := _dbff.(*Gray4); _bfca {
		return _adfa.Copy(), nil
	}
	_cec := _dbff.Bounds()
	_afca, _bddb := NewImage(_cec.Max.X, _cec.Max.Y, 4, 1, nil, nil, nil)
	if _bddb != nil {
		return nil, _bddb
	}
	_gceff(_dbff, _afca, _cec)
	return _afca, nil
}
func _dbf(_cgf, _faf *Monochrome, _ddd []byte, _ade int) (_gcg error) {
	var (
		_adc, _gce, _aee, _dcc, _ebcf, _bab, _efa, _bef int
		_ebbf, _aeb                                     uint32
		_fdc, _gdgg                                     byte
		_ecc                                            uint16
	)
	_afb := make([]byte, 4)
	_bdd := make([]byte, 4)
	for _aee = 0; _aee < _cgf.Height-1; _aee, _dcc = _aee+2, _dcc+1 {
		_adc = _aee * _cgf.BytesPerLine
		_gce = _dcc * _faf.BytesPerLine
		for _ebcf, _bab = 0, 0; _ebcf < _ade; _ebcf, _bab = _ebcf+4, _bab+1 {
			for _efa = 0; _efa < 4; _efa++ {
				_bef = _adc + _ebcf + _efa
				if _bef <= len(_cgf.Data)-1 && _bef < _adc+_cgf.BytesPerLine {
					_afb[_efa] = _cgf.Data[_bef]
				} else {
					_afb[_efa] = 0x00
				}
				_bef = _adc + _cgf.BytesPerLine + _ebcf + _efa
				if _bef <= len(_cgf.Data)-1 && _bef < _adc+(2*_cgf.BytesPerLine) {
					_bdd[_efa] = _cgf.Data[_bef]
				} else {
					_bdd[_efa] = 0x00
				}
			}
			_ebbf = _cd.BigEndian.Uint32(_afb)
			_aeb = _cd.BigEndian.Uint32(_bdd)
			_aeb |= _ebbf
			_aeb |= _aeb << 1
			_aeb &= 0xaaaaaaaa
			_ebbf = _aeb | (_aeb << 7)
			_fdc = byte(_ebbf >> 24)
			_gdgg = byte((_ebbf >> 8) & 0xff)
			_bef = _gce + _bab
			if _bef+1 == len(_faf.Data)-1 || _bef+1 >= _gce+_faf.BytesPerLine {
				_faf.Data[_bef] = _ddd[_fdc]
			} else {
				_ecc = (uint16(_ddd[_fdc]) << 8) | uint16(_ddd[_gdgg])
				if _gcg = _faf.setTwoBytes(_bef, _ecc); _gcg != nil {
					return _a.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _bef)
				}
				_bab++
			}
		}
	}
	return nil
}
func _ecgg() (_afc [256]uint64) {
	for _bgg := 0; _bgg < 256; _bgg++ {
		if _bgg&0x01 != 0 {
			_afc[_bgg] |= 0xff
		}
		if _bgg&0x02 != 0 {
			_afc[_bgg] |= 0xff00
		}
		if _bgg&0x04 != 0 {
			_afc[_bgg] |= 0xff0000
		}
		if _bgg&0x08 != 0 {
			_afc[_bgg] |= 0xff000000
		}
		if _bgg&0x10 != 0 {
			_afc[_bgg] |= 0xff00000000
		}
		if _bgg&0x20 != 0 {
			_afc[_bgg] |= 0xff0000000000
		}
		if _bgg&0x40 != 0 {
			_afc[_bgg] |= 0xff000000000000
		}
		if _bgg&0x80 != 0 {
			_afc[_bgg] |= 0xff00000000000000
		}
	}
	return _afc
}

type NRGBA64 struct{ ImageBase }

func _gcd(_cfc, _dfg Gray, _efgc _f.Rectangle) {
	for _gfcff := 0; _gfcff < _efgc.Max.X; _gfcff++ {
		for _gfeg := 0; _gfeg < _efgc.Max.Y; _gfeg++ {
			_dfg.SetGray(_gfcff, _gfeg, _cfc.GrayAt(_gfcff, _gfeg))
		}
	}
}
func _decfg(_bdad _f.Image, _efed Image, _gadb _f.Rectangle) {
	if _cbaf, _aced := _bdad.(SMasker); _aced && _cbaf.HasAlpha() {
		_efed.(SMasker).MakeAlpha()
	}
	switch _fgad := _bdad.(type) {
	case Gray:
		_acdaa(_fgad, _efed.(NRGBA), _gadb)
	case NRGBA:
		_fcda(_fgad, _efed.(NRGBA), _gadb)
	case *_f.NYCbCrA:
		_gcgc(_fgad, _efed.(NRGBA), _gadb)
	case CMYK:
		_ecab(_fgad, _efed.(NRGBA), _gadb)
	case RGBA:
		_baab(_fgad, _efed.(NRGBA), _gadb)
	case nrgba64:
		_fcbf(_fgad, _efed.(NRGBA), _gadb)
	default:
		_gfe(_bdad, _efed, _gadb)
	}
}
func (_gecg *NRGBA16) setNRGBA(_eagc, _ccde, _fdcad int, _dgea _g.NRGBA) {
	if _eagc*3%2 == 0 {
		_gecg.Data[_fdcad] = (_dgea.R>>4)<<4 | (_dgea.G >> 4)
		_gecg.Data[_fdcad+1] = (_dgea.B>>4)<<4 | (_gecg.Data[_fdcad+1] & 0xf)
	} else {
		_gecg.Data[_fdcad] = (_gecg.Data[_fdcad] & 0xf0) | (_dgea.R >> 4)
		_gecg.Data[_fdcad+1] = (_dgea.G>>4)<<4 | (_dgea.B >> 4)
	}
	if _gecg.Alpha != nil {
		_deda := _ccde * BytesPerLine(_gecg.Width, 4, 1)
		if _deda < len(_gecg.Alpha) {
			if _eagc%2 == 0 {
				_gecg.Alpha[_deda] = (_dgea.A>>uint(4))<<uint(4) | (_gecg.Alpha[_fdcad] & 0xf)
			} else {
				_gecg.Alpha[_deda] = (_gecg.Alpha[_deda] & 0xf0) | (_dgea.A >> uint(4))
			}
		}
	}
}

type Gray4 struct{ ImageBase }

func init() { _gea() }
func _bd(_feg, _bdb int) *Monochrome {
	return &Monochrome{ImageBase: NewImageBase(_feg, _bdb, 1, 1, nil, nil, nil), ModelThreshold: 0x0f}
}
func _cad(_eccf, _dad *Monochrome, _egac []byte, _gac int) (_adde error) {
	var (
		_fca, _dfcb, _cbff, _cfg, _bgf, _bda, _egf, _acf int
		_efdag, _gdgc                                    uint32
		_egag, _eccg                                     byte
		_ecba                                            uint16
	)
	_ffgf := make([]byte, 4)
	_eea := make([]byte, 4)
	for _cbff = 0; _cbff < _eccf.Height-1; _cbff, _cfg = _cbff+2, _cfg+1 {
		_fca = _cbff * _eccf.BytesPerLine
		_dfcb = _cfg * _dad.BytesPerLine
		for _bgf, _bda = 0, 0; _bgf < _gac; _bgf, _bda = _bgf+4, _bda+1 {
			for _egf = 0; _egf < 4; _egf++ {
				_acf = _fca + _bgf + _egf
				if _acf <= len(_eccf.Data)-1 && _acf < _fca+_eccf.BytesPerLine {
					_ffgf[_egf] = _eccf.Data[_acf]
				} else {
					_ffgf[_egf] = 0x00
				}
				_acf = _fca + _eccf.BytesPerLine + _bgf + _egf
				if _acf <= len(_eccf.Data)-1 && _acf < _fca+(2*_eccf.BytesPerLine) {
					_eea[_egf] = _eccf.Data[_acf]
				} else {
					_eea[_egf] = 0x00
				}
			}
			_efdag = _cd.BigEndian.Uint32(_ffgf)
			_gdgc = _cd.BigEndian.Uint32(_eea)
			_gdgc &= _efdag
			_gdgc &= _gdgc << 1
			_gdgc &= 0xaaaaaaaa
			_efdag = _gdgc | (_gdgc << 7)
			_egag = byte(_efdag >> 24)
			_eccg = byte((_efdag >> 8) & 0xff)
			_acf = _dfcb + _bda
			if _acf+1 == len(_dad.Data)-1 || _acf+1 >= _dfcb+_dad.BytesPerLine {
				_dad.Data[_acf] = _egac[_egag]
				if _adde = _dad.setByte(_acf, _egac[_egag]); _adde != nil {
					return _a.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _acf)
				}
			} else {
				_ecba = (uint16(_egac[_egag]) << 8) | uint16(_egac[_eccg])
				if _adde = _dad.setTwoBytes(_acf, _ecba); _adde != nil {
					return _a.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _acf)
				}
				_bda++
			}
		}
	}
	return nil
}
func (_fcd *Monochrome) Set(x, y int, c _g.Color) {
	_gag := y*_fcd.BytesPerLine + x>>3
	if _gag > len(_fcd.Data)-1 {
		return
	}
	_gefc := _fcd.ColorModel().Convert(c).(_g.Gray)
	_fcd.setGray(x, _gefc, _gag)
}

var _ NRGBA = &NRGBA32{}
var _ Image = &Gray2{}

func (_dfgg *NRGBA32) ColorModel() _g.Model { return _g.NRGBAModel }
func (_ebg *Monochrome) IsUnpadded() bool   { return (_ebg.Width * _ebg.Height) == len(_ebg.Data) }

type CMYK32 struct{ ImageBase }

func (_addba *Gray2) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray2BPC(x, y, _addba.BytesPerLine, _addba.Data, _addba.Decode)
}
func _fgaff(_adda Gray, _babba RGBA, _fgfd _f.Rectangle) {
	for _bffb := 0; _bffb < _fgfd.Max.X; _bffb++ {
		for _adaga := 0; _adaga < _fgfd.Max.Y; _adaga++ {
			_dgdd := _adda.GrayAt(_bffb, _adaga)
			_babba.SetRGBA(_bffb, _adaga, _afee(_dgdd))
		}
	}
}
func (_aaccg *NRGBA16) Copy() Image { return &NRGBA16{ImageBase: _aaccg.copy()} }
func (_fgb *Gray2) Histogram() (_bcgb [256]int) {
	for _fcce := 0; _fcce < _fgb.Width; _fcce++ {
		for _fff := 0; _fff < _fgb.Height; _fff++ {
			_bcgb[_fgb.GrayAt(_fcce, _fff).Y]++
		}
	}
	return _bcgb
}
func _cfd(_dcgg []byte, _aaeg Image) error {
	_aadec := true
	for _dfbdf := 0; _dfbdf < len(_dcgg); _dfbdf++ {
		if _dcgg[_dfbdf] != 0xff {
			_aadec = false
			break
		}
	}
	if _aadec {
		switch _cdfdb := _aaeg.(type) {
		case *NRGBA32:
			_cdfdb.Alpha = nil
		case *NRGBA64:
			_cdfdb.Alpha = nil
		default:
			return _a.Errorf("i\u006ete\u0072n\u0061l\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020i\u006d\u0061\u0067\u0065\u0020s\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070e\u0020\u002a\u004eRGB\u0041\u0033\u0032\u0020\u006f\u0072 \u002a\u004e\u0052\u0047\u0042\u0041\u0036\u0034\u0020\u0062\u0075\u0074 \u0069s\u003a\u0020\u0025\u0054", _aaeg)
		}
	}
	return nil
}
func _bccbb(_cfeg _f.Image) (Image, error) {
	if _dfe, _beab := _cfeg.(*Gray8); _beab {
		return _dfe.Copy(), nil
	}
	_deff := _cfeg.Bounds()
	_gedg, _fgd := NewImage(_deff.Max.X, _deff.Max.Y, 8, 1, nil, nil, nil)
	if _fgd != nil {
		return nil, _fgd
	}
	_gceff(_cfeg, _gedg, _deff)
	return _gedg, nil
}
func _agfb(_cdddd _g.Color) _g.Color {
	_efdf := _g.NRGBAModel.Convert(_cdddd).(_g.NRGBA)
	return _dbgg(_efdf)
}
func (_afdc *Gray8) GrayAt(x, y int) _g.Gray {
	_efce, _ := ColorAtGray8BPC(x, y, _afdc.BytesPerLine, _afdc.Data, _afdc.Decode)
	return _efce
}
func _ecab(_bfd CMYK, _aedc NRGBA, _adgf _f.Rectangle) {
	for _cgbg := 0; _cgbg < _adgf.Max.X; _cgbg++ {
		for _aded := 0; _aded < _adgf.Max.Y; _aded++ {
			_cgeg := _bfd.CMYKAt(_cgbg, _aded)
			_aedc.SetNRGBA(_cgbg, _aded, _fcefd(_cgeg))
		}
	}
}
func _fcef(_efbfa NRGBA, _gfg CMYK, _fcf _f.Rectangle) {
	for _fdfd := 0; _fdfd < _fcf.Max.X; _fdfd++ {
		for _ecd := 0; _ecd < _fcf.Max.Y; _ecd++ {
			_bdac := _efbfa.NRGBAAt(_fdfd, _ecd)
			_gfg.SetCMYK(_fdfd, _ecd, _cgd(_bdac))
		}
	}
}
func (_cddd *Gray4) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray4BPC(x, y, _cddd.BytesPerLine, _cddd.Data, _cddd.Decode)
}
func (_abec *Gray8) Histogram() (_gff [256]int) {
	for _cggb := 0; _cggb < len(_abec.Data); _cggb++ {
		_gff[_abec.Data[_cggb]]++
	}
	return _gff
}
func (_fgeb *Gray16) Histogram() (_gab [256]int) {
	for _fgbf := 0; _fgbf < _fgeb.Width; _fgbf++ {
		for _feff := 0; _feff < _fgeb.Height; _feff++ {
			_gab[_fgeb.GrayAt(_fgbf, _feff).Y]++
		}
	}
	return _gab
}
func (_abf *CMYK32) Validate() error {
	if len(_abf.Data) != 4*_abf.Width*_abf.Height {
		return _cb.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func AddDataPadding(width, height, bitsPerComponent, colorComponents int, data []byte) ([]byte, error) {
	_gfef := BytesPerLine(width, bitsPerComponent, colorComponents)
	if _gfef == width*colorComponents*bitsPerComponent/8 {
		return data, nil
	}
	_eab := width * colorComponents * bitsPerComponent
	_cbbg := _gfef * 8
	_gbea := 8 - (_cbbg - _eab)
	_caga := _b.NewReader(data)
	_ceed := _gfef - 1
	_gdggd := make([]byte, _ceed)
	_gcab := make([]byte, height*_gfef)
	_gefbf := _b.NewWriterMSB(_gcab)
	var _acea uint64
	var _edab error
	for _adfg := 0; _adfg < height; _adfg++ {
		_, _edab = _caga.Read(_gdggd)
		if _edab != nil {
			return nil, _edab
		}
		_, _edab = _gefbf.Write(_gdggd)
		if _edab != nil {
			return nil, _edab
		}
		_acea, _edab = _caga.ReadBits(byte(_gbea))
		if _edab != nil {
			return nil, _edab
		}
		_, _edab = _gefbf.WriteBits(_acea, _gbea)
		if _edab != nil {
			return nil, _edab
		}
		_gefbf.FinishByte()
	}
	return _gcab, nil
}
func (_agef *Monochrome) Scale(scale float64) (*Monochrome, error) {
	var _ebca bool
	_cdae := scale
	if scale < 1 {
		_cdae = 1 / scale
		_ebca = true
	}
	_gfca := NextPowerOf2(uint(_cdae))
	if InDelta(float64(_gfca), _cdae, 0.001) {
		if _ebca {
			return _agef.ReduceBinary(_cdae)
		}
		return _agef.ExpandBinary(int(_gfca))
	}
	_fbdd := int(_e.RoundToEven(float64(_agef.Width) * scale))
	_adf := int(_e.RoundToEven(float64(_agef.Height) * scale))
	return _agef.ScaleLow(_fbdd, _adf)
}
func (_ecfaa *NRGBA32) Base() *ImageBase { return &_ecfaa.ImageBase }
func (_feea *Monochrome) ReduceBinary(factor float64) (*Monochrome, error) {
	_feed := _ecag(uint(factor))
	if !IsPowerOf2(uint(factor)) {
		_feed++
	}
	_aaa := make([]int, _feed)
	for _gagg := range _aaa {
		_aaa[_gagg] = 4
	}
	_ebbe, _bgfc := _ebc(_feea, _aaa...)
	if _bgfc != nil {
		return nil, _bgfc
	}
	return _ebbe, nil
}
func ScaleAlphaToMonochrome(data []byte, width, height int) ([]byte, error) {
	_ce := BytesPerLine(width, 8, 1)
	if len(data) < _ce*height {
		return nil, nil
	}
	_ge := &Gray8{NewImageBase(width, height, 8, 1, data, nil, nil)}
	_ec, _d := MonochromeConverter.Convert(_ge)
	if _d != nil {
		return nil, _d
	}
	return _ec.Base().Data, nil
}
func _gfe(_dcbb _f.Image, _ecca Image, _dde _f.Rectangle) {
	for _cag := 0; _cag < _dde.Max.X; _cag++ {
		for _aacb := 0; _aacb < _dde.Max.Y; _aacb++ {
			_ecgc := _dcbb.At(_cag, _aacb)
			_ecca.Set(_cag, _aacb, _ecgc)
		}
	}
}
func (_cdfd *Monochrome) Copy() Image {
	return &Monochrome{ImageBase: _cdfd.ImageBase.copy(), ModelThreshold: _cdfd.ModelThreshold}
}
func (_aecc *NRGBA32) SetNRGBA(x, y int, c _g.NRGBA) {
	_cadcb := y*_aecc.Width + x
	_bddg := 3 * _cadcb
	if _bddg+2 >= len(_aecc.Data) {
		return
	}
	_aecc.setRGBA(_cadcb, c)
}
func (_fea *Monochrome) At(x, y int) _g.Color {
	_fbde, _ := _fea.ColorAt(x, y)
	return _fbde
}
func (_gdeeg *Gray4) Base() *ImageBase   { return &_gdeeg.ImageBase }
func (_ffe *Gray2) At(x, y int) _g.Color { _bff, _ := _ffe.ColorAt(x, y); return _bff }

var _ Image = &Gray16{}
var (
	_dcdfd = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_ebgf  = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func LinearInterpolate(x, xmin, xmax, ymin, ymax float64) float64 {
	if _e.Abs(xmax-xmin) < 0.000001 {
		return ymin
	}
	_fdef := ymin + (x-xmin)*(ymax-ymin)/(xmax-xmin)
	return _fdef
}

type ImageBase struct {
	Width, Height                     int
	BitsPerComponent, ColorComponents int
	Data, Alpha                       []byte
	Decode                            []float64
	BytesPerLine                      int
}

var _ Image = &NRGBA16{}
var _ Image = &RGBA32{}

func (_cccb *Gray4) ColorModel() _g.Model                        { return Gray4Model }
func (_cadg colorConverter) Convert(src _f.Image) (Image, error) { return _cadg._fcg(src) }
func _ca(_dcb *Monochrome, _ebb, _gf int) (*Monochrome, error) {
	if _dcb == nil {
		return nil, _cb.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _ebb <= 0 || _gf <= 0 {
		return nil, _cb.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _ebb == _gf {
		if _ebb == 1 {
			return _dcb.copy(), nil
		}
		if _ebb == 2 || _ebb == 4 || _ebb == 8 {
			_fc, _dda := _gb(_dcb, _ebb)
			if _dda != nil {
				return nil, _dda
			}
			return _fc, nil
		}
	}
	_efd := _ebb * _dcb.Width
	_dbc := _gf * _dcb.Height
	_fb := _bd(_efd, _dbc)
	_egb := _fb.BytesPerLine
	var (
		_ae, _dba, _aeg, _af, _be int
		_efda                     byte
		_ggd                      error
	)
	for _dba = 0; _dba < _dcb.Height; _dba++ {
		_ae = _gf * _dba * _egb
		for _aeg = 0; _aeg < _dcb.Width; _aeg++ {
			if _cbg := _dcb.getBitAt(_aeg, _dba); _cbg {
				_be = _ebb * _aeg
				for _af = 0; _af < _ebb; _af++ {
					_fb.setIndexedBit(_ae*8 + _be + _af)
				}
			}
		}
		for _af = 1; _af < _gf; _af++ {
			_abd := _ae + _af*_egb
			for _ega := 0; _ega < _egb; _ega++ {
				if _efda, _ggd = _fb.getByte(_ae + _ega); _ggd != nil {
					return nil, _ggd
				}
				if _ggd = _fb.setByte(_abd+_ega, _efda); _ggd != nil {
					return nil, _ggd
				}
			}
		}
	}
	return _fb, nil
}

type shift int

func _ccb(_geg _g.Color) _g.Color         { _gbd := _g.GrayModel.Convert(_geg).(_g.Gray); return _fdbdc(_gbd) }
func (_degb *Gray8) At(x, y int) _g.Color { _cbe, _ := _degb.ColorAt(x, y); return _cbe }

type NRGBA32 struct{ ImageBase }

func _acge(_gcbbc nrgba64, _gdfbf RGBA, _bca _f.Rectangle) {
	for _eccfd := 0; _eccfd < _bca.Max.X; _eccfd++ {
		for _ebgcc := 0; _ebgcc < _bca.Max.Y; _ebgcc++ {
			_cccc := _gcbbc.NRGBA64At(_eccfd, _ebgcc)
			_gdfbf.SetRGBA(_eccfd, _ebgcc, _gdee(_cccc))
		}
	}
}
func _cgge(_ecbe Gray, _ada CMYK, _gad _f.Rectangle) {
	for _fef := 0; _fef < _gad.Max.X; _fef++ {
		for _ebd := 0; _ebd < _gad.Max.Y; _ebd++ {
			_fagc := _ecbe.GrayAt(_fef, _ebd)
			_ada.SetCMYK(_fef, _ebd, _ddg(_fagc))
		}
	}
}
func (_afbc *CMYK32) Set(x, y int, c _g.Color) {
	_defc := 4 * (y*_afbc.Width + x)
	if _defc+3 >= len(_afbc.Data) {
		return
	}
	_ggag := _g.CMYKModel.Convert(c).(_g.CMYK)
	_afbc.Data[_defc] = _ggag.C
	_afbc.Data[_defc+1] = _ggag.M
	_afbc.Data[_defc+2] = _ggag.Y
	_afbc.Data[_defc+3] = _ggag.K
}
func (_ecge *Monochrome) clearBit(_cffde, _dbcca int) {
	_ecge.Data[_cffde] &= ^(0x80 >> uint(_dbcca&7))
}

type RGBA interface {
	RGBAAt(_dacc, _ecfa int) _g.RGBA
	SetRGBA(_dgef, _dffe int, _eaaf _g.RGBA)
}

func _edc(_agae NRGBA, _daaa Gray, _debg _f.Rectangle) {
	for _bbbbg := 0; _bbbbg < _debg.Max.X; _bbbbg++ {
		for _agga := 0; _agga < _debg.Max.Y; _agga++ {
			_faea := _eca(_agae.NRGBAAt(_bbbbg, _agga))
			_daaa.SetGray(_bbbbg, _agga, _faea)
		}
	}
}
func (_adga *Gray2) Base() *ImageBase { return &_adga.ImageBase }
func (_cffe *Gray16) Set(x, y int, c _g.Color) {
	_daag := (y*_cffe.BytesPerLine/2 + x) * 2
	if _daag+1 >= len(_cffe.Data) {
		return
	}
	_gfec := _g.Gray16Model.Convert(c).(_g.Gray16)
	_cffe.Data[_daag], _cffe.Data[_daag+1] = uint8(_gfec.Y>>8), uint8(_gfec.Y&0xff)
}

var ErrInvalidImage = _cb.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")

func (_agdg *RGBA32) Set(x, y int, c _g.Color) {
	_efdc := y*_agdg.Width + x
	_effc := 3 * _efdc
	if _effc+2 >= len(_agdg.Data) {
		return
	}
	_gdfcb := _g.RGBAModel.Convert(c).(_g.RGBA)
	_agdg.setRGBA(_efdc, _gdfcb)
}
func (_dca *CMYK32) ColorModel() _g.Model { return _g.CMYKModel }
func (_eeab *NRGBA32) setRGBA(_ddcf int, _eagcb _g.NRGBA) {
	_faefa := 3 * _ddcf
	_eeab.Data[_faefa] = _eagcb.R
	_eeab.Data[_faefa+1] = _eagcb.G
	_eeab.Data[_faefa+2] = _eagcb.B
	if _ddcf < len(_eeab.Alpha) {
		_eeab.Alpha[_ddcf] = _eagcb.A
	}
}
func (_beda *NRGBA16) NRGBAAt(x, y int) _g.NRGBA {
	_gfefe, _ := ColorAtNRGBA16(x, y, _beda.Width, _beda.BytesPerLine, _beda.Data, _beda.Alpha, _beda.Decode)
	return _gfefe
}
func _cf() (_dcf [256]uint32) {
	for _ggba := 0; _ggba < 256; _ggba++ {
		if _ggba&0x01 != 0 {
			_dcf[_ggba] |= 0xf
		}
		if _ggba&0x02 != 0 {
			_dcf[_ggba] |= 0xf0
		}
		if _ggba&0x04 != 0 {
			_dcf[_ggba] |= 0xf00
		}
		if _ggba&0x08 != 0 {
			_dcf[_ggba] |= 0xf000
		}
		if _ggba&0x10 != 0 {
			_dcf[_ggba] |= 0xf0000
		}
		if _ggba&0x20 != 0 {
			_dcf[_ggba] |= 0xf00000
		}
		if _ggba&0x40 != 0 {
			_dcf[_ggba] |= 0xf000000
		}
		if _ggba&0x80 != 0 {
			_dcf[_ggba] |= 0xf0000000
		}
	}
	return _dcf
}
func (_fafd *RGBA32) At(x, y int) _g.Color { _aefe, _ := _fafd.ColorAt(x, y); return _aefe }
func (_cba *Monochrome) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _cba.Width, Y: _cba.Height}}
}
func ColorAtGray4BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_bbdf := y*bytesPerLine + x>>1
	if _bbdf >= len(data) {
		return _g.Gray{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_eac := data[_bbdf] >> uint(4-(x&1)*4) & 0xf
	if len(decode) == 2 {
		_eac = uint8(uint32(LinearInterpolate(float64(_eac), 0, 15, decode[0], decode[1])) & 0xf)
	}
	return _g.Gray{Y: _eac * 17 & 0xff}, nil
}
func _ed(_ab *Monochrome, _ced int, _gd []uint) (*Monochrome, error) {
	_ef := _ced * _ab.Width
	_eb := _ced * _ab.Height
	_geb := _bd(_ef, _eb)
	for _ff, _ea := range _gd {
		var _cda error
		switch _ea {
		case 2:
			_cda = _cgg(_geb, _ab)
		case 4:
			_cda = _dgb(_geb, _ab)
		case 8:
			_cda = _cdc(_geb, _ab)
		}
		if _cda != nil {
			return nil, _cda
		}
		if _ff != len(_gd)-1 {
			_ab = _geb.copy()
		}
	}
	return _geb, nil
}
func (_ffb *ImageBase) copy() ImageBase {
	_baadb := *_ffb
	_baadb.Data = make([]byte, len(_ffb.Data))
	copy(_baadb.Data, _ffb.Data)
	return _baadb
}
func _ddg(_dfa _g.Gray) _g.CMYK { return _g.CMYK{K: 0xff - _dfa.Y} }

var (
	_ffg = _aac()
	_dbg = _cf()
	_cbf = _ecgg()
)

func (_efac *NRGBA64) Base() *ImageBase { return &_efac.ImageBase }

var (
	MonochromeConverter = ConverterFunc(_gggf)
	Gray2Converter      = ConverterFunc(_bfeb)
	Gray4Converter      = ConverterFunc(_dfbd)
	GrayConverter       = ConverterFunc(_bccbb)
	Gray16Converter     = ConverterFunc(_fecf)
	NRGBA16Converter    = ConverterFunc(_cbgb)
	NRGBAConverter      = ConverterFunc(_edcf)
	NRGBA64Converter    = ConverterFunc(_fefe)
	RGBAConverter       = ConverterFunc(_cabd)
	CMYKConverter       = ConverterFunc(_fec)
)

func (_bdaf *NRGBA32) Copy() Image { return &NRGBA32{ImageBase: _bdaf.copy()} }
func _edd(_aeeb, _edgb CMYK, _bcd _f.Rectangle) {
	for _acb := 0; _acb < _bcd.Max.X; _acb++ {
		for _bac := 0; _bac < _bcd.Max.Y; _bac++ {
			_edgb.SetCMYK(_acb, _bac, _aeeb.CMYKAt(_acb, _bac))
		}
	}
}
func _fda(_aea _g.Gray, _ebde monochromeModel) _g.Gray {
	if _aea.Y > uint8(_ebde) {
		return _g.Gray{Y: _e.MaxUint8}
	}
	return _g.Gray{}
}
func (_fagff *NRGBA64) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtNRGBA64(x, y, _fagff.Width, _fagff.Data, _fagff.Alpha, _fagff.Decode)
}
func MonochromeThresholdConverter(threshold uint8) ColorConverter {
	return &monochromeThresholdConverter{Threshold: threshold}
}
func _fcefd(_bfc _g.CMYK) _g.NRGBA {
	_gdb, _ggbg, _bba := _g.CMYKToRGB(_bfc.C, _bfc.M, _bfc.Y, _bfc.K)
	return _g.NRGBA{R: _gdb, G: _ggbg, B: _bba, A: 0xff}
}
func (_gdfd *Monochrome) Base() *ImageBase { return &_gdfd.ImageBase }
func (_agea *Monochrome) getBitAt(_fegb, _ccfff int) bool {
	_cffd := _ccfff*_agea.BytesPerLine + (_fegb >> 3)
	_fagf := _fegb & 0x07
	_adafa := uint(7 - _fagf)
	if _cffd > len(_agea.Data)-1 {
		return false
	}
	if (_agea.Data[_cffd]>>_adafa)&0x01 >= 1 {
		return true
	}
	return false
}
func ColorAtRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_g.RGBA, error) {
	_eceg := y*width + x
	_ccfa := 3 * _eceg
	if _ccfa+2 >= len(data) {
		return _g.RGBA{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_cbffa := uint8(0xff)
	if alpha != nil && len(alpha) > _eceg {
		_cbffa = alpha[_eceg]
	}
	_gaacg, _bcbc, _cfcg := data[_ccfa], data[_ccfa+1], data[_ccfa+2]
	if len(decode) == 6 {
		_gaacg = uint8(uint32(LinearInterpolate(float64(_gaacg), 0, 255, decode[0], decode[1])) & 0xff)
		_bcbc = uint8(uint32(LinearInterpolate(float64(_bcbc), 0, 255, decode[2], decode[3])) & 0xff)
		_cfcg = uint8(uint32(LinearInterpolate(float64(_cfcg), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _g.RGBA{R: _gaacg, G: _bcbc, B: _cfcg, A: _cbffa}, nil
}
func _dgbb(_daff RGBA, _fbddg Gray, _ffddd _f.Rectangle) {
	for _efae := 0; _efae < _ffddd.Max.X; _efae++ {
		for _adad := 0; _adad < _ffddd.Max.Y; _adad++ {
			_ecggc := _gddg(_daff.RGBAAt(_efae, _adad))
			_fbddg.SetGray(_efae, _adad, _ecggc)
		}
	}
}
func (_caae *NRGBA16) ColorModel() _g.Model { return NRGBA16Model }
func (_ebaa *monochromeThresholdConverter) Convert(img _f.Image) (Image, error) {
	if _cde, _abfe := img.(*Monochrome); _abfe {
		return _cde.Copy(), nil
	}
	_ecf := img.Bounds()
	_dcfg, _fbb := NewImage(_ecf.Max.X, _ecf.Max.Y, 1, 1, nil, nil, nil)
	if _fbb != nil {
		return nil, _fbb
	}
	_dcfg.(*Monochrome).ModelThreshold = _ebaa.Threshold
	for _gfgd := 0; _gfgd < _ecf.Max.X; _gfgd++ {
		for _aff := 0; _aff < _ecf.Max.Y; _aff++ {
			_cdb := img.At(_gfgd, _aff)
			_dcfg.Set(_gfgd, _aff, _cdb)
		}
	}
	return _dcfg, nil
}
func _dbab(_dgg *Monochrome, _bfb int, _fa []byte) (_fe *Monochrome, _ece error) {
	const _dbaa = "\u0072\u0065d\u0075\u0063\u0065R\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079"
	if _dgg == nil {
		return nil, _cb.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _bfb < 1 || _bfb > 4 {
		return nil, _cb.New("\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _dgg.Height <= 1 {
		return nil, _cb.New("\u0073\u006f\u0075rc\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020m\u0075s\u0074 \u0062e\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027")
	}
	_fe = _bd(_dgg.Width/2, _dgg.Height/2)
	if _fa == nil {
		_fa = _caf()
	}
	_dfb := _debgb(_dgg.BytesPerLine, 2*_fe.BytesPerLine)
	switch _bfb {
	case 1:
		_ece = _dbf(_dgg, _fe, _fa, _dfb)
	case 2:
		_ece = _bbb(_dgg, _fe, _fa, _dfb)
	case 3:
		_ece = _cbd(_dgg, _fe, _fa, _dfb)
	case 4:
		_ece = _cad(_dgg, _fe, _fa, _dfb)
	}
	if _ece != nil {
		return nil, _ece
	}
	return _fe, nil
}
func _cgd(_aacc _g.NRGBA) _g.CMYK {
	_cgc, _fga, _ccd, _ := _aacc.RGBA()
	_fafg, _afcb, _adg, _cdf := _g.RGBToCMYK(uint8(_cgc>>8), uint8(_fga>>8), uint8(_ccd>>8))
	return _g.CMYK{C: _fafg, M: _afcb, Y: _adg, K: _cdf}
}
func (_bdbc *ImageBase) setEightBytes(_ggbad int, _fgfa uint64) error {
	_ffbd := _bdbc.BytesPerLine - (_ggbad % _bdbc.BytesPerLine)
	if _bdbc.BytesPerLine != _bdbc.Width>>3 {
		_ffbd--
	}
	if _ffbd >= 8 {
		return _bdbc.setEightFullBytes(_ggbad, _fgfa)
	}
	return _bdbc.setEightPartlyBytes(_ggbad, _ffbd, _fgfa)
}

type Histogramer interface{ Histogram() [256]int }
type monochromeThresholdConverter struct{ Threshold uint8 }

func (_bffc *Gray4) setGray(_bggga int, _bed int, _faad _g.Gray) {
	_acdd := _bed * _bffc.BytesPerLine
	_ceba := _acdd + (_bggga >> 1)
	if _ceba >= len(_bffc.Data) {
		return
	}
	_eccd := _faad.Y >> 4
	_bffc.Data[_ceba] = (_bffc.Data[_ceba] & (^(0xf0 >> uint(4*(_bggga&1))))) | (_eccd << uint(4-4*(_bggga&1)))
}
func _dgb(_egg, _abc *Monochrome) (_dc error) {
	_ad := _abc.BytesPerLine
	_da := _egg.BytesPerLine
	_cdd := _abc.BytesPerLine*4 - _egg.BytesPerLine
	var (
		_gde, _ee                             byte
		_cddf                                 uint32
		_df, _fdg, _dea, _gef, _bc, _bf, _deb int
	)
	for _dea = 0; _dea < _abc.Height; _dea++ {
		_df = _dea * _ad
		_fdg = 4 * _dea * _da
		for _gef = 0; _gef < _ad; _gef++ {
			_gde = _abc.Data[_df+_gef]
			_cddf = _dbg[_gde]
			_bf = _fdg + _gef*4
			if _cdd != 0 && (_gef+1)*4 > _egg.BytesPerLine {
				for _bc = _cdd; _bc > 0; _bc-- {
					_ee = byte((_cddf >> uint(_bc*8)) & 0xff)
					_deb = _bf + (_cdd - _bc)
					if _dc = _egg.setByte(_deb, _ee); _dc != nil {
						return _dc
					}
				}
			} else if _dc = _egg.setFourBytes(_bf, _cddf); _dc != nil {
				return _dc
			}
			if _dc = _egg.setFourBytes(_fdg+_gef*4, _dbg[_abc.Data[_df+_gef]]); _dc != nil {
				return _dc
			}
		}
		for _bc = 1; _bc < 4; _bc++ {
			for _gef = 0; _gef < _da; _gef++ {
				if _dc = _egg.setByte(_fdg+_bc*_da+_gef, _egg.Data[_fdg+_gef]); _dc != nil {
					return _dc
				}
			}
		}
	}
	return nil
}
func (_caca *RGBA32) Validate() error {
	if len(_caca.Data) != 3*_caca.Width*_caca.Height {
		return _cb.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

const (
	_gafc shift = iota
	_cfaee
)

func (_bggee *Gray16) Base() *ImageBase { return &_bggee.ImageBase }
func (_gba *Monochrome) copy() *Monochrome {
	_dbbb := _bd(_gba.Width, _gba.Height)
	_dbbb.ModelThreshold = _gba.ModelThreshold
	_dbbb.Data = make([]byte, len(_gba.Data))
	copy(_dbbb.Data, _gba.Data)
	if len(_gba.Decode) != 0 {
		_dbbb.Decode = make([]float64, len(_gba.Decode))
		copy(_dbbb.Decode, _gba.Decode)
	}
	if len(_gba.Alpha) != 0 {
		_dbbb.Alpha = make([]byte, len(_gba.Alpha))
		copy(_dbbb.Alpha, _gba.Alpha)
	}
	return _dbbb
}
func _gb(_age *Monochrome, _bg int) (*Monochrome, error) {
	if _age == nil {
		return nil, _cb.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _bg == 1 {
		return _age.copy(), nil
	}
	if !IsPowerOf2(uint(_bg)) {
		return nil, _a.Errorf("\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006ci\u0064 \u0065x\u0070a\u006e\u0064\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _bg)
	}
	_dd := _def(_bg)
	return _ed(_age, _bg, _dd)
}
func (_cga *Gray8) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray8BPC(x, y, _cga.BytesPerLine, _cga.Data, _cga.Decode)
}
func ColorAtGray2BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_cfed := y*bytesPerLine + x>>2
	if _cfed >= len(data) {
		return _g.Gray{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_adfb := data[_cfed] >> uint(6-(x&3)*2) & 3
	if len(decode) == 2 {
		_adfb = uint8(uint32(LinearInterpolate(float64(_adfb), 0, 3.0, decode[0], decode[1])) & 3)
	}
	return _g.Gray{Y: _adfb * 85}, nil
}
func (_bdacb *NRGBA64) setNRGBA64(_dfaae int, _cbac _g.NRGBA64, _bfae int) {
	_bdacb.Data[_dfaae] = uint8(_cbac.R >> 8)
	_bdacb.Data[_dfaae+1] = uint8(_cbac.R & 0xff)
	_bdacb.Data[_dfaae+2] = uint8(_cbac.G >> 8)
	_bdacb.Data[_dfaae+3] = uint8(_cbac.G & 0xff)
	_bdacb.Data[_dfaae+4] = uint8(_cbac.B >> 8)
	_bdacb.Data[_dfaae+5] = uint8(_cbac.B & 0xff)
	if _bfae+1 < len(_bdacb.Alpha) {
		_bdacb.Alpha[_bfae] = uint8(_cbac.A >> 8)
		_bdacb.Alpha[_bfae+1] = uint8(_cbac.A & 0xff)
	}
}

type RasterOperator int

func _deffe(_aeda *_f.NYCbCrA, _bfac RGBA, _eaad _f.Rectangle) {
	for _beb := 0; _beb < _eaad.Max.X; _beb++ {
		for _fegge := 0; _fegge < _eaad.Max.Y; _fegge++ {
			_bfbg := _aeda.NYCbCrAAt(_beb, _fegge)
			_bfac.SetRGBA(_beb, _fegge, _fdebb(_bfbg))
		}
	}
}
func _gea() {
	for _dbfb := 0; _dbfb < 256; _dbfb++ {
		_bcde[_dbfb] = uint8(_dbfb&0x1) + (uint8(_dbfb>>1) & 0x1) + (uint8(_dbfb>>2) & 0x1) + (uint8(_dbfb>>3) & 0x1) + (uint8(_dbfb>>4) & 0x1) + (uint8(_dbfb>>5) & 0x1) + (uint8(_dbfb>>6) & 0x1) + (uint8(_dbfb>>7) & 0x1)
	}
}

var _ _f.Image = &NRGBA16{}

func (_egdc *NRGBA64) Validate() error {
	if len(_egdc.Data) != 3*2*_egdc.Width*_egdc.Height {
		return _cb.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

type CMYK interface {
	CMYKAt(_gaf, _ffd int) _g.CMYK
	SetCMYK(_dga, _fbc int, _aga _g.CMYK)
}

func _aac() (_cee [256]uint16) {
	for _edg := 0; _edg < 256; _edg++ {
		if _edg&0x01 != 0 {
			_cee[_edg] |= 0x3
		}
		if _edg&0x02 != 0 {
			_cee[_edg] |= 0xc
		}
		if _edg&0x04 != 0 {
			_cee[_edg] |= 0x30
		}
		if _edg&0x08 != 0 {
			_cee[_edg] |= 0xc0
		}
		if _edg&0x10 != 0 {
			_cee[_edg] |= 0x300
		}
		if _edg&0x20 != 0 {
			_cee[_edg] |= 0xc00
		}
		if _edg&0x40 != 0 {
			_cee[_edg] |= 0x3000
		}
		if _edg&0x80 != 0 {
			_cee[_edg] |= 0xc000
		}
	}
	return _cee
}

var _ Image = &Monochrome{}

func _fee(_ddb _g.NRGBA) _g.Gray {
	var _gggd _g.NRGBA
	if _ddb == _gggd {
		return _g.Gray{Y: 0xff}
	}
	_eec, _cabc, _baca, _ := _ddb.RGBA()
	_dbb := (19595*_eec + 38470*_cabc + 7471*_baca + 1<<15) >> 24
	return _g.Gray{Y: uint8(_dbb)}
}

var _ Gray = &Gray8{}

func ColorAtNRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_g.NRGBA, error) {
	_gedcc := y*width + x
	_fccfa := 3 * _gedcc
	if _fccfa+2 >= len(data) {
		return _g.NRGBA{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_cgfe := uint8(0xff)
	if alpha != nil && len(alpha) > _gedcc {
		_cgfe = alpha[_gedcc]
	}
	_gbaf, _fegbf, _bce := data[_fccfa], data[_fccfa+1], data[_fccfa+2]
	if len(decode) == 6 {
		_gbaf = uint8(uint32(LinearInterpolate(float64(_gbaf), 0, 255, decode[0], decode[1])) & 0xff)
		_fegbf = uint8(uint32(LinearInterpolate(float64(_fegbf), 0, 255, decode[2], decode[3])) & 0xff)
		_bce = uint8(uint32(LinearInterpolate(float64(_bce), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _g.NRGBA{R: _gbaf, G: _fegbf, B: _bce, A: _cgfe}, nil
}
func (_cfce *Monochrome) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _dcaa(_cfce, dx, dy, dw, dh, op, src, sx, sy)
}
func _bcgf(_fgae *Monochrome, _dgbc, _ggcf int, _ddbfe, _eafg int, _cafb RasterOperator) {
	var (
		_aadb        int
		_cdbfa       byte
		_gabf, _fgeg int
		_dfdf        int
	)
	_agf := _ddbfe >> 3
	_cbge := _ddbfe & 7
	if _cbge > 0 {
		_cdbfa = _dcdfd[_cbge]
	}
	_aadb = _fgae.BytesPerLine*_ggcf + (_dgbc >> 3)
	switch _cafb {
	case PixClr:
		for _gabf = 0; _gabf < _eafg; _gabf++ {
			_dfdf = _aadb + _gabf*_fgae.BytesPerLine
			for _fgeg = 0; _fgeg < _agf; _fgeg++ {
				_fgae.Data[_dfdf] = 0x0
				_dfdf++
			}
			if _cbge > 0 {
				_fgae.Data[_dfdf] = _dacg(_fgae.Data[_dfdf], 0x0, _cdbfa)
			}
		}
	case PixSet:
		for _gabf = 0; _gabf < _eafg; _gabf++ {
			_dfdf = _aadb + _gabf*_fgae.BytesPerLine
			for _fgeg = 0; _fgeg < _agf; _fgeg++ {
				_fgae.Data[_dfdf] = 0xff
				_dfdf++
			}
			if _cbge > 0 {
				_fgae.Data[_dfdf] = _dacg(_fgae.Data[_dfdf], 0xff, _cdbfa)
			}
		}
	case PixNotDst:
		for _gabf = 0; _gabf < _eafg; _gabf++ {
			_dfdf = _aadb + _gabf*_fgae.BytesPerLine
			for _fgeg = 0; _fgeg < _agf; _fgeg++ {
				_fgae.Data[_dfdf] = ^_fgae.Data[_dfdf]
				_dfdf++
			}
			if _cbge > 0 {
				_fgae.Data[_dfdf] = _dacg(_fgae.Data[_dfdf], ^_fgae.Data[_dfdf], _cdbfa)
			}
		}
	}
}
func _afdb(_fgdb, _afgg uint8) uint8 {
	if _fgdb < _afgg {
		return 255
	}
	return 0
}
func (_eddf *Gray8) Base() *ImageBase { return &_eddf.ImageBase }
func (_acfcd *ImageBase) MakeAlpha()  { _acfcd.newAlpha() }
func _begd(_ffce *Monochrome, _beca, _cgfcf, _addg, _fedc int, _geeg RasterOperator) {
	if _beca < 0 {
		_addg += _beca
		_beca = 0
	}
	_defa := _beca + _addg - _ffce.Width
	if _defa > 0 {
		_addg -= _defa
	}
	if _cgfcf < 0 {
		_fedc += _cgfcf
		_cgfcf = 0
	}
	_debbb := _cgfcf + _fedc - _ffce.Height
	if _debbb > 0 {
		_fedc -= _debbb
	}
	if _addg <= 0 || _fedc <= 0 {
		return
	}
	if (_beca & 7) == 0 {
		_bcgf(_ffce, _beca, _cgfcf, _addg, _fedc, _geeg)
	} else {
		_adfe(_ffce, _beca, _cgfcf, _addg, _fedc, _geeg)
	}
}
func (_egad *CMYK32) CMYKAt(x, y int) _g.CMYK {
	_gbec, _ := ColorAtCMYK(x, y, _egad.Width, _egad.Data, _egad.Decode)
	return _gbec
}
func (_ged *Gray2) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _ged.Width, Y: _ged.Height}}
}
func _fbd(_eecc _g.CMYK) _g.RGBA {
	_dcg, _cae, _bfe := _g.CMYKToRGB(_eecc.C, _eecc.M, _eecc.Y, _eecc.K)
	return _g.RGBA{R: _dcg, G: _cae, B: _bfe, A: 0xff}
}
func _ageec(_fcfc _f.Image, _ggbb Image, _agbcg _f.Rectangle) {
	if _affb, _fbed := _fcfc.(SMasker); _fbed && _affb.HasAlpha() {
		_ggbb.(SMasker).MakeAlpha()
	}
	_gfe(_fcfc, _ggbb, _agbcg)
}
func _cgg(_aa, _gg *Monochrome) (_fd error) {
	_de := _gg.BytesPerLine
	_dg := _aa.BytesPerLine
	var (
		_eg                          byte
		_ga                          uint16
		_ggb, _efg, _ac, _efb, _cdad int
	)
	for _ac = 0; _ac < _gg.Height; _ac++ {
		_ggb = _ac * _de
		_efg = 2 * _ac * _dg
		for _efb = 0; _efb < _de; _efb++ {
			_eg = _gg.Data[_ggb+_efb]
			_ga = _ffg[_eg]
			_cdad = _efg + _efb*2
			if _aa.BytesPerLine != _gg.BytesPerLine*2 && (_efb+1)*2 > _aa.BytesPerLine {
				_fd = _aa.setByte(_cdad, byte(_ga>>8))
			} else {
				_fd = _aa.setTwoBytes(_cdad, _ga)
			}
			if _fd != nil {
				return _fd
			}
		}
		for _efb = 0; _efb < _dg; _efb++ {
			_cdad = _efg + _dg + _efb
			_eg = _aa.Data[_efg+_efb]
			if _fd = _aa.setByte(_cdad, _eg); _fd != nil {
				return _fd
			}
		}
	}
	return nil
}
func (_abddf *NRGBA32) NRGBAAt(x, y int) _g.NRGBA {
	_efab, _ := ColorAtNRGBA32(x, y, _abddf.Width, _abddf.Data, _abddf.Alpha, _abddf.Decode)
	return _efab
}
func _cggc(_fefc _g.Color) _g.Color {
	_cgged := _g.GrayModel.Convert(_fefc).(_g.Gray)
	return _aeeg(_cgged)
}
func (_gefce *Gray2) Set(x, y int, c _g.Color) {
	if x >= _gefce.Width || y >= _gefce.Height {
		return
	}
	_eefd := Gray2Model.Convert(c).(_g.Gray)
	_decf := y * _gefce.BytesPerLine
	_ede := _decf + (x >> 2)
	_dfbe := _eefd.Y >> 6
	_gefce.Data[_ede] = (_gefce.Data[_ede] & (^(0xc0 >> uint(2*((x)&3))))) | (_dfbe << uint(6-2*(x&3)))
}

var _ Image = &NRGBA64{}

func (_fac *CMYK32) Base() *ImageBase { return &_fac.ImageBase }
func _fcbf(_deeb nrgba64, _gcabd NRGBA, _dbfd _f.Rectangle) {
	for _ecad := 0; _ecad < _dbfd.Max.X; _ecad++ {
		for _bgbe := 0; _bgbe < _dbfd.Max.Y; _bgbe++ {
			_aeeeb := _deeb.NRGBA64At(_ecad, _bgbe)
			_gcabd.SetNRGBA(_ecad, _bgbe, _aca(_aeeeb))
		}
	}
}
func (_cdcc *Monochrome) setIndexedBit(_bcgcc int) {
	_cdcc.Data[(_bcgcc >> 3)] |= 0x80 >> uint(_bcgcc&7)
}
func _adb(_bbd Gray, _fedb NRGBA, _dae _f.Rectangle) {
	for _fafga := 0; _fafga < _dae.Max.X; _fafga++ {
		for _dgbe := 0; _dgbe < _dae.Max.Y; _dgbe++ {
			_afdg := _fee(_fedb.NRGBAAt(_fafga, _dgbe))
			_bbd.SetGray(_fafga, _dgbe, _afdg)
		}
	}
}
func _gfff(_gcgf *Monochrome, _gffe, _ecgf int, _begc, _cgae int, _aec RasterOperator, _ccab *Monochrome, _dfba, _beaef int) error {
	var _aeee, _edcb, _edcc, _geef int
	if _gffe < 0 {
		_dfba -= _gffe
		_begc += _gffe
		_gffe = 0
	}
	if _dfba < 0 {
		_gffe -= _dfba
		_begc += _dfba
		_dfba = 0
	}
	_aeee = _gffe + _begc - _gcgf.Width
	if _aeee > 0 {
		_begc -= _aeee
	}
	_edcb = _dfba + _begc - _ccab.Width
	if _edcb > 0 {
		_begc -= _edcb
	}
	if _ecgf < 0 {
		_beaef -= _ecgf
		_cgae += _ecgf
		_ecgf = 0
	}
	if _beaef < 0 {
		_ecgf -= _beaef
		_cgae += _beaef
		_beaef = 0
	}
	_edcc = _ecgf + _cgae - _gcgf.Height
	if _edcc > 0 {
		_cgae -= _edcc
	}
	_geef = _beaef + _cgae - _ccab.Height
	if _geef > 0 {
		_cgae -= _geef
	}
	if _begc <= 0 || _cgae <= 0 {
		return nil
	}
	var _gaac error
	switch {
	case _gffe&7 == 0 && _dfba&7 == 0:
		_gaac = _bafab(_gcgf, _gffe, _ecgf, _begc, _cgae, _aec, _ccab, _dfba, _beaef)
	case _gffe&7 == _dfba&7:
		_gaac = _ddgc(_gcgf, _gffe, _ecgf, _begc, _cgae, _aec, _ccab, _dfba, _beaef)
	default:
		_gaac = _ecdg(_gcgf, _gffe, _ecgf, _begc, _cgae, _aec, _ccab, _dfba, _beaef)
	}
	if _gaac != nil {
		return _gaac
	}
	return nil
}
func IsGrayImgBlackAndWhite(i *_f.Gray) bool { return _cadeee(i) }
func _aeeg(_dafd _g.Gray) _g.Gray {
	_aed := _dafd.Y >> 6
	_aed |= _aed << 2
	_dafd.Y = _aed | _aed<<4
	return _dafd
}
func (_dbec *ImageBase) setTwoBytes(_afaf int, _aeef uint16) error {
	if _afaf+1 > len(_dbec.Data)-1 {
		return _cb.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_dbec.Data[_afaf] = byte((_aeef & 0xff00) >> 8)
	_dbec.Data[_afaf+1] = byte(_aeef & 0xff)
	return nil
}
func _ddbfg(_fegcg _f.Image, _fagg int) (_f.Rectangle, bool, []byte) {
	_cfee := _fegcg.Bounds()
	var (
		_aabf bool
		_ffga []byte
	)
	switch _adge := _fegcg.(type) {
	case SMasker:
		_aabf = _adge.HasAlpha()
	case NRGBA, RGBA, *_f.RGBA64, nrgba64, *_f.NYCbCrA:
		_ffga = make([]byte, _cfee.Max.X*_cfee.Max.Y*_fagg)
	case *_f.Paletted:
		var _gfad bool
		for _, _dagg := range _adge.Palette {
			_dafg, _ffbdc, _dbgc, _ebgc := _dagg.RGBA()
			if _dafg == 0 && _ffbdc == 0 && _dbgc == 0 && _ebgc != 0 {
				_gfad = true
				break
			}
		}
		if _gfad {
			_ffga = make([]byte, _cfee.Max.X*_cfee.Max.Y*_fagg)
		}
	}
	return _cfee, _aabf, _ffga
}
func (_egdg *Gray4) Set(x, y int, c _g.Color) {
	if x >= _egdg.Width || y >= _egdg.Height {
		return
	}
	_cfgg := Gray4Model.Convert(c).(_g.Gray)
	_egdg.setGray(x, y, _cfgg)
}
func (_ceg *Monochrome) setGray(_dfbg int, _cage _g.Gray, _gdffd int) {
	if _cage.Y == 0 {
		_ceg.clearBit(_gdffd, _dfbg)
	} else {
		_ceg.setGrayBit(_gdffd, _dfbg)
	}
}
func _fae(_cce _g.NRGBA64) _g.Gray {
	var _bagb _g.NRGBA64
	if _cce == _bagb {
		return _g.Gray{Y: 0xff}
	}
	_deag, _dgbg, _faga, _ := _cce.RGBA()
	_gcb := (19595*_deag + 38470*_dgbg + 7471*_faga + 1<<15) >> 24
	return _g.Gray{Y: uint8(_gcb)}
}
func _defac(_aebf *_f.Gray16, _cfbe uint8) *_f.Gray {
	_ccdg := _aebf.Bounds()
	_gdec := _f.NewGray(_ccdg)
	for _dfgge := 0; _dfgge < _ccdg.Dx(); _dfgge++ {
		for _cbee := 0; _cbee < _ccdg.Dy(); _cbee++ {
			_ceab := _aebf.Gray16At(_dfgge, _cbee)
			_gdec.SetGray(_dfgge, _cbee, _g.Gray{Y: _afdb(uint8(_ceab.Y/256), _cfbe)})
		}
	}
	return _gdec
}
func (_cef *Gray8) Set(x, y int, c _g.Color) {
	_adce := y*_cef.BytesPerLine + x
	if _adce > len(_cef.Data)-1 {
		return
	}
	_dcbe := _g.GrayModel.Convert(c)
	_cef.Data[_adce] = _dcbe.(_g.Gray).Y
}
func (_abdd *Monochrome) ColorModel() _g.Model { return MonochromeModel(_abdd.ModelThreshold) }
func (_gcggd *NRGBA32) Set(x, y int, c _g.Color) {
	_gfb := y*_gcggd.Width + x
	_gbc := 3 * _gfb
	if _gbc+2 >= len(_gcggd.Data) {
		return
	}
	_bbdd := _g.NRGBAModel.Convert(c).(_g.NRGBA)
	_gcggd.setRGBA(_gfb, _bbdd)
}
func _egec(_eefg CMYK, _effd Gray, _gee _f.Rectangle) {
	for _fdge := 0; _fdge < _gee.Max.X; _fdge++ {
		for _afea := 0; _afea < _gee.Max.Y; _afea++ {
			_cccd := _ggc(_eefg.CMYKAt(_fdge, _afea))
			_effd.SetGray(_fdge, _afea, _cccd)
		}
	}
}
func (_gccg *Monochrome) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray1BPC(x, y, _gccg.BytesPerLine, _gccg.Data, _gccg.Decode)
}
func _fcda(_bfeaf, _gfccb NRGBA, _ffbg _f.Rectangle) {
	for _baaff := 0; _baaff < _ffbg.Max.X; _baaff++ {
		for _babb := 0; _babb < _ffbg.Max.Y; _babb++ {
			_gfccb.SetNRGBA(_baaff, _babb, _bfeaf.NRGBAAt(_baaff, _babb))
		}
	}
}

type NRGBA interface {
	NRGBAAt(_adcf, _fgcf int) _g.NRGBA
	SetNRGBA(_becf, _aecef int, _acda _g.NRGBA)
}

func _ddgc(_decdg *Monochrome, _gfcad, _fbac, _bcge, _gcac int, _egfag RasterOperator, _fdgae *Monochrome, _gbda, _acg int) error {
	var (
		_acbgf     bool
		_ebcdc     bool
		_aggae     int
		_fgaf      int
		_cagbb     int
		_fffe      bool
		_gdag      byte
		_eacc      int
		_aaf       int
		_ecdc      int
		_cbc, _aef int
	)
	_eaab := 8 - (_gfcad & 7)
	_bdgc := _ebgf[_eaab]
	_dbfa := _decdg.BytesPerLine*_fbac + (_gfcad >> 3)
	_fecfd := _fdgae.BytesPerLine*_acg + (_gbda >> 3)
	if _bcge < _eaab {
		_acbgf = true
		_bdgc &= _dcdfd[8-_eaab+_bcge]
	}
	if !_acbgf {
		_aggae = (_bcge - _eaab) >> 3
		if _aggae > 0 {
			_ebcdc = true
			_fgaf = _dbfa + 1
			_cagbb = _fecfd + 1
		}
	}
	_eacc = (_gfcad + _bcge) & 7
	if !(_acbgf || _eacc == 0) {
		_fffe = true
		_gdag = _dcdfd[_eacc]
		_aaf = _dbfa + 1 + _aggae
		_ecdc = _fecfd + 1 + _aggae
	}
	switch _egfag {
	case PixSrc:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], _fdgae.Data[_fecfd], _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] = _fdgae.Data[_cagbb+_aef]
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], _fdgae.Data[_ecdc], _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixNotSrc:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], ^_fdgae.Data[_fecfd], _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] = ^_fdgae.Data[_cagbb+_aef]
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], ^_fdgae.Data[_ecdc], _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixSrcOrDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], _fdgae.Data[_fecfd]|_decdg.Data[_dbfa], _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] |= _fdgae.Data[_cagbb+_aef]
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], _fdgae.Data[_ecdc]|_decdg.Data[_aaf], _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixSrcAndDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], _fdgae.Data[_fecfd]&_decdg.Data[_dbfa], _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] &= _fdgae.Data[_cagbb+_aef]
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], _fdgae.Data[_ecdc]&_decdg.Data[_aaf], _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixSrcXorDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], _fdgae.Data[_fecfd]^_decdg.Data[_dbfa], _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] ^= _fdgae.Data[_cagbb+_aef]
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], _fdgae.Data[_ecdc]^_decdg.Data[_aaf], _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], ^(_fdgae.Data[_fecfd])|_decdg.Data[_dbfa], _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] |= ^(_fdgae.Data[_cagbb+_aef])
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], ^(_fdgae.Data[_ecdc])|_decdg.Data[_aaf], _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], ^(_fdgae.Data[_fecfd])&_decdg.Data[_dbfa], _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] &= ^_fdgae.Data[_cagbb+_aef]
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], ^(_fdgae.Data[_ecdc])&_decdg.Data[_aaf], _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], _fdgae.Data[_fecfd]|^(_decdg.Data[_dbfa]), _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] = _fdgae.Data[_cagbb+_aef] | ^(_decdg.Data[_fgaf+_aef])
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], _fdgae.Data[_ecdc]|^(_decdg.Data[_aaf]), _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], _fdgae.Data[_fecfd]&^(_decdg.Data[_dbfa]), _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] = _fdgae.Data[_cagbb+_aef] &^ (_decdg.Data[_fgaf+_aef])
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], _fdgae.Data[_ecdc]&^(_decdg.Data[_aaf]), _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], ^(_fdgae.Data[_fecfd] | _decdg.Data[_dbfa]), _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] = ^(_fdgae.Data[_cagbb+_aef] | _decdg.Data[_fgaf+_aef])
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], ^(_fdgae.Data[_ecdc] | _decdg.Data[_aaf]), _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], ^(_fdgae.Data[_fecfd] & _decdg.Data[_dbfa]), _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] = ^(_fdgae.Data[_cagbb+_aef] & _decdg.Data[_fgaf+_aef])
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], ^(_fdgae.Data[_ecdc] & _decdg.Data[_aaf]), _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		for _cbc = 0; _cbc < _gcac; _cbc++ {
			_decdg.Data[_dbfa] = _dacg(_decdg.Data[_dbfa], ^(_fdgae.Data[_fecfd] ^ _decdg.Data[_dbfa]), _bdgc)
			_dbfa += _decdg.BytesPerLine
			_fecfd += _fdgae.BytesPerLine
		}
		if _ebcdc {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				for _aef = 0; _aef < _aggae; _aef++ {
					_decdg.Data[_fgaf+_aef] = ^(_fdgae.Data[_cagbb+_aef] ^ _decdg.Data[_fgaf+_aef])
				}
				_fgaf += _decdg.BytesPerLine
				_cagbb += _fdgae.BytesPerLine
			}
		}
		if _fffe {
			for _cbc = 0; _cbc < _gcac; _cbc++ {
				_decdg.Data[_aaf] = _dacg(_decdg.Data[_aaf], ^(_fdgae.Data[_ecdc] ^ _decdg.Data[_aaf]), _gdag)
				_aaf += _decdg.BytesPerLine
				_ecdc += _fdgae.BytesPerLine
			}
		}
	default:
		_ag.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _egfag)
		return _cb.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_agdc *NRGBA32) At(x, y int) _g.Color { _deed, _ := _agdc.ColorAt(x, y); return _deed }
func (_ggbf *Gray8) Copy() Image            { return &Gray8{ImageBase: _ggbf.copy()} }
func (_bafba *NRGBA16) Validate() error {
	if len(_bafba.Data) != 3*_bafba.Width*_bafba.Height/2 {
		return _cb.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func _dacg(_gbf, _gdfg, _cggba byte) byte    { return (_gbf &^ (_cggba)) | (_gdfg & _cggba) }
func (_dfdfa *NRGBA16) At(x, y int) _g.Color { _baadg, _ := _dfdfa.ColorAt(x, y); return _baadg }
func NextPowerOf2(n uint) uint {
	if IsPowerOf2(n) {
		return n
	}
	return 1 << (_ecag(n) + 1)
}

var _ _f.Image = &RGBA32{}

func (_agcgb *Gray8) Validate() error {
	if len(_agcgb.Data) != _agcgb.Height*_agcgb.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func NewImageBase(width int, height int, bitsPerComponent int, colorComponents int, data []byte, alpha []byte, decode []float64) ImageBase {
	_gaae := ImageBase{Width: width, Height: height, BitsPerComponent: bitsPerComponent, ColorComponents: colorComponents, Data: data, Alpha: alpha, Decode: decode, BytesPerLine: BytesPerLine(width, bitsPerComponent, colorComponents)}
	if data == nil {
		_gaae.Data = make([]byte, height*_gaae.BytesPerLine)
	}
	return _gaae
}
func _deef(_cfedb, _ebeg RGBA, _dfdc _f.Rectangle) {
	for _cagcb := 0; _cagcb < _dfdc.Max.X; _cagcb++ {
		for _eefdg := 0; _eefdg < _dfdc.Max.Y; _eefdg++ {
			_ebeg.SetRGBA(_cagcb, _eefdg, _cfedb.RGBAAt(_cagcb, _eefdg))
		}
	}
}

type Monochrome struct {
	ImageBase
	ModelThreshold uint8
}

func (_cgcaf *NRGBA16) Set(x, y int, c _g.Color) {
	_cedb := y*_cgcaf.BytesPerLine + x*3/2
	if _cedb+1 >= len(_cgcaf.Data) {
		return
	}
	_eeac := NRGBA16Model.Convert(c).(_g.NRGBA)
	_cgcaf.setNRGBA(x, y, _cedb, _eeac)
}
func (_cgeb *NRGBA64) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _cgeb.Width, Y: _cgeb.Height}}
}

var _ _f.Image = &Gray16{}
var _ Image = &NRGBA32{}

func _adfe(_agbd *Monochrome, _caee, _gdfb int, _bdae, _fdec int, _fceg RasterOperator) {
	var (
		_gdef  bool
		_ebaf  bool
		_faee  int
		_bgdf  int
		_eadg  int
		_cebfd int
		_agac  bool
		_eaac  byte
	)
	_ggbe := 8 - (_caee & 7)
	_cadc := _ebgf[_ggbe]
	_cagc := _agbd.BytesPerLine*_gdfb + (_caee >> 3)
	if _bdae < _ggbe {
		_gdef = true
		_cadc &= _dcdfd[8-_ggbe+_bdae]
	}
	if !_gdef {
		_faee = (_bdae - _ggbe) >> 3
		if _faee != 0 {
			_ebaf = true
			_bgdf = _cagc + 1
		}
	}
	_eadg = (_caee + _bdae) & 7
	if !(_gdef || _eadg == 0) {
		_agac = true
		_eaac = _dcdfd[_eadg]
		_cebfd = _cagc + 1 + _faee
	}
	var _bfbe, _aefg int
	switch _fceg {
	case PixClr:
		for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
			_agbd.Data[_cagc] = _dacg(_agbd.Data[_cagc], 0x0, _cadc)
			_cagc += _agbd.BytesPerLine
		}
		if _ebaf {
			for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
				for _aefg = 0; _aefg < _faee; _aefg++ {
					_agbd.Data[_bgdf+_aefg] = 0x0
				}
				_bgdf += _agbd.BytesPerLine
			}
		}
		if _agac {
			for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
				_agbd.Data[_cebfd] = _dacg(_agbd.Data[_cebfd], 0x0, _eaac)
				_cebfd += _agbd.BytesPerLine
			}
		}
	case PixSet:
		for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
			_agbd.Data[_cagc] = _dacg(_agbd.Data[_cagc], 0xff, _cadc)
			_cagc += _agbd.BytesPerLine
		}
		if _ebaf {
			for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
				for _aefg = 0; _aefg < _faee; _aefg++ {
					_agbd.Data[_bgdf+_aefg] = 0xff
				}
				_bgdf += _agbd.BytesPerLine
			}
		}
		if _agac {
			for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
				_agbd.Data[_cebfd] = _dacg(_agbd.Data[_cebfd], 0xff, _eaac)
				_cebfd += _agbd.BytesPerLine
			}
		}
	case PixNotDst:
		for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
			_agbd.Data[_cagc] = _dacg(_agbd.Data[_cagc], ^_agbd.Data[_cagc], _cadc)
			_cagc += _agbd.BytesPerLine
		}
		if _ebaf {
			for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
				for _aefg = 0; _aefg < _faee; _aefg++ {
					_agbd.Data[_bgdf+_aefg] = ^(_agbd.Data[_bgdf+_aefg])
				}
				_bgdf += _agbd.BytesPerLine
			}
		}
		if _agac {
			for _bfbe = 0; _bfbe < _fdec; _bfbe++ {
				_agbd.Data[_cebfd] = _dacg(_agbd.Data[_cebfd], ^_agbd.Data[_cebfd], _eaac)
				_cebfd += _agbd.BytesPerLine
			}
		}
	}
}
func (_efea *Gray4) Histogram() (_gedc [256]int) {
	for _faff := 0; _faff < _efea.Width; _faff++ {
		for _cedg := 0; _cedg < _efea.Height; _cedg++ {
			_gedc[_efea.GrayAt(_faff, _cedg).Y]++
		}
	}
	return _gedc
}
func (_fcab *RGBA32) setRGBA(_begdc int, _cedd _g.RGBA) {
	_ecbb := 3 * _begdc
	_fcab.Data[_ecbb] = _cedd.R
	_fcab.Data[_ecbb+1] = _cedd.G
	_fcab.Data[_ecbb+2] = _cedd.B
	if _begdc < len(_fcab.Alpha) {
		_fcab.Alpha[_begdc] = _cedd.A
	}
}
func (_bae *NRGBA64) Set(x, y int, c _g.Color) {
	_fgcfc := (y*_bae.Width + x) * 2
	_bfcaf := _fgcfc * 3
	if _bfcaf+5 >= len(_bae.Data) {
		return
	}
	_fbbd := _g.NRGBA64Model.Convert(c).(_g.NRGBA64)
	_bae.setNRGBA64(_bfcaf, _fbbd, _fgcfc)
}
func _cabd(_bbed _f.Image) (Image, error) {
	if _egcg, _acdg := _bbed.(*RGBA32); _acdg {
		return _egcg.Copy(), nil
	}
	_edbbd, _dfef, _degbe := _ddbfg(_bbed, 1)
	_fgda := &RGBA32{ImageBase: NewImageBase(_edbbd.Max.X, _edbbd.Max.Y, 8, 3, nil, _degbe, nil)}
	_ecgd(_bbed, _fgda, _edbbd)
	if len(_degbe) != 0 && !_dfef {
		if _dbed := _cfd(_degbe, _fgda); _dbed != nil {
			return nil, _dbed
		}
	}
	return _fgda, nil
}
func _fec(_eef _f.Image) (Image, error) {
	if _accd, _bad := _eef.(*CMYK32); _bad {
		return _accd.Copy(), nil
	}
	_gdf := _eef.Bounds()
	_dag, _ege := NewImage(_gdf.Max.X, _gdf.Max.Y, 8, 4, nil, nil, nil)
	if _ege != nil {
		return nil, _ege
	}
	switch _bag := _eef.(type) {
	case CMYK:
		_edd(_bag, _dag.(CMYK), _gdf)
	case Gray:
		_cgge(_bag, _dag.(CMYK), _gdf)
	case NRGBA:
		_fcef(_bag, _dag.(CMYK), _gdf)
	case RGBA:
		_cgb(_bag, _dag.(CMYK), _gdf)
	default:
		_gfe(_eef, _dag, _gdf)
	}
	return _dag, nil
}
func (_acaa *NRGBA64) Copy() Image { return &NRGBA64{ImageBase: _acaa.copy()} }

var _ NRGBA = &NRGBA16{}

func (_gdc *ImageBase) setFourBytes(_bcb int, _ebbd uint32) error {
	if _bcb+3 > len(_gdc.Data)-1 {
		return _a.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _bcb)
	}
	_gdc.Data[_bcb] = byte((_ebbd & 0xff000000) >> 24)
	_gdc.Data[_bcb+1] = byte((_ebbd & 0xff0000) >> 16)
	_gdc.Data[_bcb+2] = byte((_ebbd & 0xff00) >> 8)
	_gdc.Data[_bcb+3] = byte(_ebbd & 0xff)
	return nil
}

var _ _f.Image = &Gray8{}

func _fdebb(_dfcbg _g.NYCbCrA) _g.RGBA {
	_facc, _cagb, _dbge, _bdag := _gebbe(_dfcbg).RGBA()
	return _g.RGBA{R: uint8(_facc >> 8), G: uint8(_cagb >> 8), B: uint8(_dbge >> 8), A: uint8(_bdag >> 8)}
}

type SMasker interface {
	HasAlpha() bool
	GetAlpha() []byte
	MakeAlpha()
}

func (_aaba *CMYK32) SetCMYK(x, y int, c _g.CMYK) {
	_addd := 4 * (y*_aaba.Width + x)
	if _addd+3 >= len(_aaba.Data) {
		return
	}
	_aaba.Data[_addd] = c.C
	_aaba.Data[_addd+1] = c.M
	_aaba.Data[_addd+2] = c.Y
	_aaba.Data[_addd+3] = c.K
}
func _afee(_gdff _g.Gray) _g.RGBA { return _g.RGBA{R: _gdff.Y, G: _gdff.Y, B: _gdff.Y, A: 0xff} }
func (_bee *ImageBase) setEightFullBytes(_eebd int, _faef uint64) error {
	if _eebd+7 > len(_bee.Data)-1 {
		return _cb.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bee.Data[_eebd] = byte((_faef & 0xff00000000000000) >> 56)
	_bee.Data[_eebd+1] = byte((_faef & 0xff000000000000) >> 48)
	_bee.Data[_eebd+2] = byte((_faef & 0xff0000000000) >> 40)
	_bee.Data[_eebd+3] = byte((_faef & 0xff00000000) >> 32)
	_bee.Data[_eebd+4] = byte((_faef & 0xff000000) >> 24)
	_bee.Data[_eebd+5] = byte((_faef & 0xff0000) >> 16)
	_bee.Data[_eebd+6] = byte((_faef & 0xff00) >> 8)
	_bee.Data[_eebd+7] = byte(_faef & 0xff)
	return nil
}
func _ggc(_bbg _g.CMYK) _g.Gray {
	_cdg, _bdf, _eeb := _g.CMYKToRGB(_bbg.C, _bbg.M, _bbg.Y, _bbg.K)
	_cbfa := (19595*uint32(_cdg) + 38470*uint32(_bdf) + 7471*uint32(_eeb) + 1<<7) >> 16
	return _g.Gray{Y: uint8(_cbfa)}
}

var _ Gray = &Gray2{}

func (_ccff *Monochrome) InverseData() error {
	return _ccff.RasterOperation(0, 0, _ccff.Width, _ccff.Height, PixNotDst, nil, 0, 0)
}

var _ _f.Image = &NRGBA64{}
var _ _f.Image = &NRGBA32{}

func _egace(_fadd Gray, _gcge nrgba64, _daf _f.Rectangle) {
	for _dagc := 0; _dagc < _daf.Max.X; _dagc++ {
		for _ggdd := 0; _ggdd < _daf.Max.Y; _ggdd++ {
			_cfe := _fae(_gcge.NRGBA64At(_dagc, _ggdd))
			_fadd.SetGray(_dagc, _ggdd, _cfe)
		}
	}
}
func _gceff(_eeaa _f.Image, _cgda Image, _deac _f.Rectangle) {
	switch _aadd := _eeaa.(type) {
	case Gray:
		_gcd(_aadd, _cgda.(Gray), _deac)
	case NRGBA:
		_edc(_aadd, _cgda.(Gray), _deac)
	case CMYK:
		_egec(_aadd, _cgda.(Gray), _deac)
	case RGBA:
		_dgbb(_aadd, _cgda.(Gray), _deac)
	default:
		_gfe(_eeaa, _cgda.(Image), _deac)
	}
}
func (_debb *Gray16) GrayAt(x, y int) _g.Gray {
	_efgd, _ := _debb.ColorAt(x, y)
	return _g.Gray{Y: uint8(_efgd.(_g.Gray16).Y >> 8)}
}
func (_aeebf *NRGBA32) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtNRGBA32(x, y, _aeebf.Width, _aeebf.Data, _aeebf.Alpha, _aeebf.Decode)
}
func (_cgad *RGBA32) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _cgad.Width, Y: _cgad.Height}}
}
func _cdc(_efbf, _daa *Monochrome) (_ecg error) {
	_gga := _daa.BytesPerLine
	_gbe := _efbf.BytesPerLine
	var _add, _gbef, _edb, _eed, _db int
	for _edb = 0; _edb < _daa.Height; _edb++ {
		_add = _edb * _gga
		_gbef = 8 * _edb * _gbe
		for _eed = 0; _eed < _gga; _eed++ {
			if _ecg = _efbf.setEightBytes(_gbef+_eed*8, _cbf[_daa.Data[_add+_eed]]); _ecg != nil {
				return _ecg
			}
		}
		for _db = 1; _db < 8; _db++ {
			for _eed = 0; _eed < _gbe; _eed++ {
				if _ecg = _efbf.setByte(_gbef+_db*_gbe+_eed, _efbf.Data[_gbef+_eed]); _ecg != nil {
					return _ecg
				}
			}
		}
	}
	return nil
}
func (_ggf *Monochrome) ExpandBinary(factor int) (*Monochrome, error) {
	if !IsPowerOf2(uint(factor)) {
		return nil, _a.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0065\u0078\u0070\u0061\u006e\u0064\u0020b\u0069n\u0061\u0072\u0079\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", factor)
	}
	return _gb(_ggf, factor)
}
func _bacda(_gdda *_f.Gray, _fbgb uint8) *_f.Gray {
	_aaaeb := _gdda.Bounds()
	_fcff := _f.NewGray(_aaaeb)
	for _ecgag := 0; _ecgag < _aaaeb.Dx(); _ecgag++ {
		for _agfg := 0; _agfg < _aaaeb.Dy(); _agfg++ {
			_dcfa := _gdda.GrayAt(_ecgag, _agfg)
			_fcff.SetGray(_ecgag, _agfg, _g.Gray{Y: _afdb(_dcfa.Y, _fbgb)})
		}
	}
	return _fcff
}
func _cbgb(_feaa _f.Image) (Image, error) {
	if _aeec, _bgfe := _feaa.(*NRGBA16); _bgfe {
		return _aeec.Copy(), nil
	}
	_fgbc := _feaa.Bounds()
	_gdbg, _bfed := NewImage(_fgbc.Max.X, _fgbc.Max.Y, 4, 3, nil, nil, nil)
	if _bfed != nil {
		return nil, _bfed
	}
	_decfg(_feaa, _gdbg, _fgbc)
	return _gdbg, nil
}
func (_gagb *RGBA32) RGBAAt(x, y int) _g.RGBA {
	_adedd, _ := ColorAtRGBA32(x, y, _gagb.Width, _gagb.Data, _gagb.Alpha, _gagb.Decode)
	return _adedd
}
func MonochromeModel(threshold uint8) _g.Model { return monochromeModel(threshold) }

var _ Image = &CMYK32{}

type RGBA32 struct{ ImageBase }

func _gebbe(_dbe _g.NYCbCrA) _g.NRGBA {
	_dbcf := int32(_dbe.Y) * 0x10101
	_dbfe := int32(_dbe.Cb) - 128
	_bea := int32(_dbe.Cr) - 128
	_dbcc := _dbcf + 91881*_bea
	if uint32(_dbcc)&0xff000000 == 0 {
		_dbcc >>= 8
	} else {
		_dbcc = ^(_dbcc >> 31) & 0xffff
	}
	_fefa := _dbcf - 22554*_dbfe - 46802*_bea
	if uint32(_fefa)&0xff000000 == 0 {
		_fefa >>= 8
	} else {
		_fefa = ^(_fefa >> 31) & 0xffff
	}
	_eae := _dbcf + 116130*_dbfe
	if uint32(_eae)&0xff000000 == 0 {
		_eae >>= 8
	} else {
		_eae = ^(_eae >> 31) & 0xffff
	}
	return _g.NRGBA{R: uint8(_dbcc >> 8), G: uint8(_fefa >> 8), B: uint8(_eae >> 8), A: _dbe.A}
}
func (_fbge *NRGBA32) Validate() error {
	if len(_fbge.Data) != 3*_fbge.Width*_fbge.Height {
		return _cb.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

type Gray16 struct{ ImageBase }

func (_acag *Gray16) ColorModel() _g.Model { return _g.Gray16Model }
func (_eddd *Gray8) SetGray(x, y int, g _g.Gray) {
	_bgb := y*_eddd.BytesPerLine + x
	if _bgb > len(_eddd.Data)-1 {
		return
	}
	_eddd.Data[_bgb] = g.Y
}
func (_bfea *NRGBA64) ColorModel() _g.Model { return _g.NRGBA64Model }
func ImgToGray(i _f.Image) *_f.Gray {
	if _gcdf, _fece := i.(*_f.Gray); _fece {
		return _gcdf
	}
	_eaaga := i.Bounds()
	_gcce := _f.NewGray(_eaaga)
	for _cgdc := 0; _cgdc < _eaaga.Max.X; _cgdc++ {
		for _bcea := 0; _bcea < _eaaga.Max.Y; _bcea++ {
			_ffde := i.At(_cgdc, _bcea)
			_gcce.Set(_cgdc, _bcea, _ffde)
		}
	}
	return _gcce
}
func _dfaaee(_gdca NRGBA, _ddgf RGBA, _efdd _f.Rectangle) {
	for _agdd := 0; _agdd < _efdd.Max.X; _agdd++ {
		for _caab := 0; _caab < _efdd.Max.Y; _caab++ {
			_dffb := _gdca.NRGBAAt(_agdd, _caab)
			_ddgf.SetRGBA(_agdd, _caab, _dbcg(_dffb))
		}
	}
}
func _eca(_ffc _g.NRGBA) _g.Gray {
	_eegd, _aade, _cade, _ := _ffc.RGBA()
	_fcc := (19595*_eegd + 38470*_aade + 7471*_cade + 1<<15) >> 24
	return _g.Gray{Y: uint8(_fcc)}
}
func _bfeb(_faec _f.Image) (Image, error) {
	if _gcbb, _cdbd := _faec.(*Gray2); _cdbd {
		return _gcbb.Copy(), nil
	}
	_ccgc := _faec.Bounds()
	_gecb, _edbc := NewImage(_ccgc.Max.X, _ccgc.Max.Y, 2, 1, nil, nil, nil)
	if _edbc != nil {
		return nil, _edbc
	}
	_gceff(_faec, _gecb, _ccgc)
	return _gecb, nil
}
func (_eecg *Gray2) Validate() error {
	if len(_eecg.Data) != _eecg.Height*_eecg.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

var _ RGBA = &RGBA32{}

func ColorAt(x, y, width, bitsPerColor, colorComponents, bytesPerLine int, data, alpha []byte, decode []float64) (_g.Color, error) {
	switch colorComponents {
	case 1:
		return ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine, data, decode)
	case 3:
		return ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor, data, alpha, decode)
	case 4:
		return ColorAtCMYK(x, y, width, data, decode)
	default:
		return nil, _a.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063o\u006c\u006f\u0072\u0020\u0063\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", colorComponents)
	}
}
func _bafab(_cgaf *Monochrome, _aggg, _ccfb, _ddbf, _ebaaa int, _ffdg RasterOperator, _daac *Monochrome, _caa, _afg int) error {
	var (
		_afcc        byte
		_eggbd       int
		_ddda        int
		_geaa, _ecea int
		_ecac, _bfab int
	)
	_aede := _ddbf >> 3
	_dfdb := _ddbf & 7
	if _dfdb > 0 {
		_afcc = _dcdfd[_dfdb]
	}
	_eggbd = _daac.BytesPerLine*_afg + (_caa >> 3)
	_ddda = _cgaf.BytesPerLine*_ccfb + (_aggg >> 3)
	switch _ffdg {
	case PixSrc:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] = _daac.Data[_geaa]
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], _daac.Data[_geaa], _afcc)
			}
		}
	case PixNotSrc:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] = ^(_daac.Data[_geaa])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], ^_daac.Data[_geaa], _afcc)
			}
		}
	case PixSrcOrDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] |= _daac.Data[_geaa]
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], _daac.Data[_geaa]|_cgaf.Data[_ecea], _afcc)
			}
		}
	case PixSrcAndDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] &= _daac.Data[_geaa]
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], _daac.Data[_geaa]&_cgaf.Data[_ecea], _afcc)
			}
		}
	case PixSrcXorDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] ^= _daac.Data[_geaa]
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], _daac.Data[_geaa]^_cgaf.Data[_ecea], _afcc)
			}
		}
	case PixNotSrcOrDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] |= ^(_daac.Data[_geaa])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], ^(_daac.Data[_geaa])|_cgaf.Data[_ecea], _afcc)
			}
		}
	case PixNotSrcAndDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] &= ^(_daac.Data[_geaa])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], ^(_daac.Data[_geaa])&_cgaf.Data[_ecea], _afcc)
			}
		}
	case PixSrcOrNotDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] = _daac.Data[_geaa] | ^(_cgaf.Data[_ecea])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], _daac.Data[_geaa]|^(_cgaf.Data[_ecea]), _afcc)
			}
		}
	case PixSrcAndNotDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] = _daac.Data[_geaa] &^ (_cgaf.Data[_ecea])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], _daac.Data[_geaa]&^(_cgaf.Data[_ecea]), _afcc)
			}
		}
	case PixNotPixSrcOrDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] = ^(_daac.Data[_geaa] | _cgaf.Data[_ecea])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], ^(_daac.Data[_geaa] | _cgaf.Data[_ecea]), _afcc)
			}
		}
	case PixNotPixSrcAndDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] = ^(_daac.Data[_geaa] & _cgaf.Data[_ecea])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], ^(_daac.Data[_geaa] & _cgaf.Data[_ecea]), _afcc)
			}
		}
	case PixNotPixSrcXorDst:
		for _ecac = 0; _ecac < _ebaaa; _ecac++ {
			_geaa = _eggbd + _ecac*_daac.BytesPerLine
			_ecea = _ddda + _ecac*_cgaf.BytesPerLine
			for _bfab = 0; _bfab < _aede; _bfab++ {
				_cgaf.Data[_ecea] = ^(_daac.Data[_geaa] ^ _cgaf.Data[_ecea])
				_ecea++
				_geaa++
			}
			if _dfdb > 0 {
				_cgaf.Data[_ecea] = _dacg(_cgaf.Data[_ecea], ^(_daac.Data[_geaa] ^ _cgaf.Data[_ecea]), _afcc)
			}
		}
	default:
		_ag.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _ffdg)
		return _cb.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_dac monochromeModel) Convert(c _g.Color) _g.Color {
	_cfgc := _g.GrayModel.Convert(c).(_g.Gray)
	return _fda(_cfgc, _dac)
}
func (_gfgf *Monochrome) getBit(_cfa, _dcbba int) uint8 {
	return _gfgf.Data[_cfa+(_dcbba>>3)] >> uint(7-(_dcbba&7)) & 1
}
func ColorAtGray8BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray, error) {
	_bfg := y*bytesPerLine + x
	if _bfg >= len(data) {
		return _g.Gray{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_cdec := data[_bfg]
	if len(decode) == 2 {
		_cdec = uint8(uint32(LinearInterpolate(float64(_cdec), 0, 255, decode[0], decode[1])) & 0xff)
	}
	return _g.Gray{Y: _cdec}, nil
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
		return nil, _a.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0067\u0062\u0020b\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0061\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}

type ColorConverter interface {
	Convert(_acd _f.Image) (Image, error)
}

func (_ead *Monochrome) ResolveDecode() error {
	if len(_ead.Decode) != 2 {
		return nil
	}
	if _ead.Decode[0] == 1 && _ead.Decode[1] == 0 {
		if _eceb := _ead.InverseData(); _eceb != nil {
			return _eceb
		}
		_ead.Decode = nil
	}
	return nil
}
func (_ggddb *Gray16) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _ggddb.Width, Y: _ggddb.Height}}
}
func InDelta(expected, current, delta float64) bool {
	_ebda := expected - current
	if _ebda <= -delta || _ebda >= delta {
		return false
	}
	return true
}
func (_bcgg *Gray16) At(x, y int) _g.Color { _efdab, _ := _bcgg.ColorAt(x, y); return _efdab }

var _ Gray = &Gray16{}

func ColorAtGray16BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_g.Gray16, error) {
	_cbb := (y*bytesPerLine/2 + x) * 2
	if _cbb+1 >= len(data) {
		return _g.Gray16{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_dcge := uint16(data[_cbb])<<8 | uint16(data[_cbb+1])
	if len(decode) == 2 {
		_dcge = uint16(uint64(LinearInterpolate(float64(_dcge), 0, 65535, decode[0], decode[1])))
	}
	return _g.Gray16{Y: _dcge}, nil
}

var _ _f.Image = &Monochrome{}

func _cbd(_fab, _ccf *Monochrome, _edgg []byte, _acc int) (_afe error) {
	var (
		_ebcd, _cfb, _bcg, _gaa, _gcc, _ebe, _ccc, _ccg int
		_bfba, _bafa, _cbda, _bcf                       uint32
		_fed, _gfc                                      byte
		_gebb                                           uint16
	)
	_abe := make([]byte, 4)
	_cbdc := make([]byte, 4)
	for _bcg = 0; _bcg < _fab.Height-1; _bcg, _gaa = _bcg+2, _gaa+1 {
		_ebcd = _bcg * _fab.BytesPerLine
		_cfb = _gaa * _ccf.BytesPerLine
		for _gcc, _ebe = 0, 0; _gcc < _acc; _gcc, _ebe = _gcc+4, _ebe+1 {
			for _ccc = 0; _ccc < 4; _ccc++ {
				_ccg = _ebcd + _gcc + _ccc
				if _ccg <= len(_fab.Data)-1 && _ccg < _ebcd+_fab.BytesPerLine {
					_abe[_ccc] = _fab.Data[_ccg]
				} else {
					_abe[_ccc] = 0x00
				}
				_ccg = _ebcd + _fab.BytesPerLine + _gcc + _ccc
				if _ccg <= len(_fab.Data)-1 && _ccg < _ebcd+(2*_fab.BytesPerLine) {
					_cbdc[_ccc] = _fab.Data[_ccg]
				} else {
					_cbdc[_ccc] = 0x00
				}
			}
			_bfba = _cd.BigEndian.Uint32(_abe)
			_bafa = _cd.BigEndian.Uint32(_cbdc)
			_cbda = _bfba & _bafa
			_cbda |= _cbda << 1
			_bcf = _bfba | _bafa
			_bcf &= _bcf << 1
			_bafa = _cbda & _bcf
			_bafa &= 0xaaaaaaaa
			_bfba = _bafa | (_bafa << 7)
			_fed = byte(_bfba >> 24)
			_gfc = byte((_bfba >> 8) & 0xff)
			_ccg = _cfb + _ebe
			if _ccg+1 == len(_ccf.Data)-1 || _ccg+1 >= _cfb+_ccf.BytesPerLine {
				if _afe = _ccf.setByte(_ccg, _edgg[_fed]); _afe != nil {
					return _a.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _ccg)
				}
			} else {
				_gebb = (uint16(_edgg[_fed]) << 8) | uint16(_edgg[_gfc])
				if _afe = _ccf.setTwoBytes(_ccg, _gebb); _afe != nil {
					return _a.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _ccg)
				}
				_ebe++
			}
		}
	}
	return nil
}
func _gddg(_bafb _g.RGBA) _g.Gray {
	_fcge := (19595*uint32(_bafb.R) + 38470*uint32(_bafb.G) + 7471*uint32(_bafb.B) + 1<<7) >> 16
	return _g.Gray{Y: uint8(_fcge)}
}
func (_ecead *RGBA32) ColorModel() _g.Model { return _g.NRGBAModel }
func _aca(_ggcb _g.NRGBA64) _g.NRGBA {
	return _g.NRGBA{R: uint8(_ggcb.R >> 8), G: uint8(_ggcb.G >> 8), B: uint8(_ggcb.B >> 8), A: uint8(_ggcb.A >> 8)}
}
func _def(_fg int) []uint {
	var _bgge []uint
	_ceb := _fg
	_abda := _ceb / 8
	if _abda != 0 {
		for _bcc := 0; _bcc < _abda; _bcc++ {
			_bgge = append(_bgge, 8)
		}
		_fde := _ceb % 8
		_ceb = 0
		if _fde != 0 {
			_ceb = _fde
		}
	}
	_ba := _ceb / 4
	if _ba != 0 {
		for _bb := 0; _bb < _ba; _bb++ {
			_bgge = append(_bgge, 4)
		}
		_fdga := _ceb % 4
		_ceb = 0
		if _fdga != 0 {
			_ceb = _fdga
		}
	}
	_cebf := _ceb / 2
	if _cebf != 0 {
		for _gc := 0; _gc < _cebf; _gc++ {
			_bgge = append(_bgge, 2)
		}
	}
	return _bgge
}

type Gray2 struct{ ImageBase }

func _gdee(_bdfb _g.NRGBA64) _g.RGBA {
	_cff, _cgcg, _aaeb, _eff := _bdfb.RGBA()
	return _g.RGBA{R: uint8(_cff >> 8), G: uint8(_cgcg >> 8), B: uint8(_aaeb >> 8), A: uint8(_eff >> 8)}
}

var _ _f.Image = &Gray4{}

func (_egc *Gray16) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtGray16BPC(x, y, _egc.BytesPerLine, _egc.Data, _egc.Decode)
}
func (_fcdc *NRGBA32) Bounds() _f.Rectangle {
	return _f.Rectangle{Max: _f.Point{X: _fcdc.Width, Y: _fcdc.Height}}
}
func (_fbe *Gray2) ColorModel() _g.Model { return Gray2Model }
func (_agcg *Monochrome) SetGray(x, y int, g _g.Gray) {
	_afeb := y*_agcg.BytesPerLine + x>>3
	if _afeb > len(_agcg.Data)-1 {
		return
	}
	g = _fda(g, monochromeModel(_agcg.ModelThreshold))
	_agcg.setGray(x, g, _afeb)
}
func (_gfac *NRGBA64) NRGBA64At(x, y int) _g.NRGBA64 {
	_fdac, _ := ColorAtNRGBA64(x, y, _gfac.Width, _gfac.Data, _gfac.Alpha, _gfac.Decode)
	return _fdac
}
func (_fadc *ImageBase) setEightPartlyBytes(_bacaa, _bcdb int, _cbae uint64) (_caff error) {
	var (
		_fcag byte
		_cbab int
	)
	for _cdff := 1; _cdff <= _bcdb; _cdff++ {
		_cbab = 64 - _cdff*8
		_fcag = byte(_cbae >> uint(_cbab) & 0xff)
		if _caff = _fadc.setByte(_bacaa+_cdff-1, _fcag); _caff != nil {
			return _caff
		}
	}
	_dcdc := _fadc.BytesPerLine*8 - _fadc.Width
	if _dcdc == 0 {
		return nil
	}
	_cbab -= 8
	_fcag = byte(_cbae>>uint(_cbab)&0xff) << uint(_dcdc)
	if _caff = _fadc.setByte(_bacaa+_bcdb, _fcag); _caff != nil {
		return _caff
	}
	return nil
}

type nrgba64 interface {
	NRGBA64At(_cddfb, _gggdc int) _g.NRGBA64
	SetNRGBA64(_bgad, _dfea int, _afed _g.NRGBA64)
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
	return nil, _a.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0043o\u006e\u0076\u0065\u0072\u0074\u0065\u0072\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072\u0073\u002e\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003a\u0020\u0025\u0064\u002c\u0020\u0043\u006f\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u003a \u0025\u0064", bitsPerComponent, colorComponents)
}
func RasterOperation(dest *Monochrome, dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _dcaa(dest, dx, dy, dw, dh, op, src, sx, sy)
}

type monochromeModel uint8

func AutoThresholdTriangle(histogram [256]int) uint8 {
	var _fbee, _cfcga, _fedg, _dbeg int
	for _aabag := 0; _aabag < len(histogram); _aabag++ {
		if histogram[_aabag] > 0 {
			_fbee = _aabag
			break
		}
	}
	if _fbee > 0 {
		_fbee--
	}
	for _dgec := 255; _dgec > 0; _dgec-- {
		if histogram[_dgec] > 0 {
			_dbeg = _dgec
			break
		}
	}
	if _dbeg < 255 {
		_dbeg++
	}
	for _fdbe := 0; _fdbe < 256; _fdbe++ {
		if histogram[_fdbe] > _cfcga {
			_fedg = _fdbe
			_cfcga = histogram[_fdbe]
		}
	}
	var _cbgde bool
	if (_fedg - _fbee) < (_dbeg - _fedg) {
		_cbgde = true
		var _fdcd int
		_dggd := 255
		for _fdcd < _dggd {
			_fgde := histogram[_fdcd]
			histogram[_fdcd] = histogram[_dggd]
			histogram[_dggd] = _fgde
			_fdcd++
			_dggd--
		}
		_fbee = 255 - _dbeg
		_fedg = 255 - _fedg
	}
	if _fbee == _fedg {
		return uint8(_fbee)
	}
	_efcc := float64(histogram[_fedg])
	_bdgg := float64(_fbee - _fedg)
	_abaa := _e.Sqrt(_efcc*_efcc + _bdgg*_bdgg)
	_efcc /= _abaa
	_bdgg /= _abaa
	_abaa = _efcc*float64(_fbee) + _bdgg*float64(histogram[_fbee])
	_fead := _fbee
	var _ddad float64
	for _deca := _fbee + 1; _deca <= _fedg; _deca++ {
		_gfgb := _efcc*float64(_deca) + _bdgg*float64(histogram[_deca]) - _abaa
		if _gfgb > _ddad {
			_fead = _deca
			_ddad = _gfgb
		}
	}
	_fead--
	if _cbgde {
		var _bdbe int
		_dgeg := 255
		for _bdbe < _dgeg {
			_gebba := histogram[_bdbe]
			histogram[_bdbe] = histogram[_dgeg]
			histogram[_dgeg] = _gebba
			_bdbe++
			_dgeg--
		}
		return uint8(255 - _fead)
	}
	return uint8(_fead)
}
func (_eedf *Gray2) GrayAt(x, y int) _g.Gray {
	_gceb, _ := ColorAtGray2BPC(x, y, _eedf.BytesPerLine, _eedf.Data, _eedf.Decode)
	return _gceb
}
func _gcgc(_ccffb *_f.NYCbCrA, _gbefa NRGBA, _efaf _f.Rectangle) {
	for _aaaec := 0; _aaaec < _efaf.Max.X; _aaaec++ {
		for _bada := 0; _bada < _efaf.Max.Y; _bada++ {
			_debd := _ccffb.NYCbCrAAt(_aaaec, _bada)
			_gbefa.SetNRGBA(_aaaec, _bada, _gebbe(_debd))
		}
	}
}
func _cdfde(_fbgg int, _gcded int) error {
	return _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", _fbgg, _gcded)
}
func (_cdadf *Gray16) Validate() error {
	if len(_cdadf.Data) != _cdadf.Height*_cdadf.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

type NRGBA16 struct{ ImageBase }

func _bdda(_eggf _g.RGBA) _g.NRGBA {
	switch _eggf.A {
	case 0xff:
		return _g.NRGBA{R: _eggf.R, G: _eggf.G, B: _eggf.B, A: 0xff}
	case 0x00:
		return _g.NRGBA{}
	default:
		_gcee, _adaf, _fggd, _fgfc := _eggf.RGBA()
		_gcee = (_gcee * 0xffff) / _fgfc
		_adaf = (_adaf * 0xffff) / _fgfc
		_fggd = (_fggd * 0xffff) / _fgfc
		return _g.NRGBA{R: uint8(_gcee >> 8), G: uint8(_adaf >> 8), B: uint8(_fggd >> 8), A: uint8(_fgfc >> 8)}
	}
}
func _dbcg(_egfa _g.NRGBA) _g.RGBA {
	_abac, _eee, _efe, _fgc := _egfa.RGBA()
	return _g.RGBA{R: uint8(_abac >> 8), G: uint8(_eee >> 8), B: uint8(_efe >> 8), A: uint8(_fgc >> 8)}
}
func (_bcff *NRGBA16) SetNRGBA(x, y int, c _g.NRGBA) {
	_eddb := y*_bcff.BytesPerLine + x*3/2
	if _eddb+1 >= len(_bcff.Data) {
		return
	}
	c = _dbgg(c)
	_bcff.setNRGBA(x, y, _eddb, c)
}
func (_abg *Gray4) SetGray(x, y int, g _g.Gray) {
	if x >= _abg.Width || y >= _abg.Height {
		return
	}
	g = _fdbdc(g)
	_abg.setGray(x, y, g)
}
func (_cdcb *Monochrome) setBit(_cdgf, _bbba int) {
	_cdcb.Data[_cdgf+(_bbba>>3)] |= 0x80 >> uint(_bbba&7)
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
		return nil, _a.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0067\u0072\u0061\u0079\u0020\u0073c\u0061\u006c\u0065\u0020\u0062\u0069\u0074s\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020a\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}

var _ Gray = &Gray4{}

func _dbgg(_cefa _g.NRGBA) _g.NRGBA {
	_cefa.R = _cefa.R>>4 | (_cefa.R>>4)<<4
	_cefa.G = _cefa.G>>4 | (_cefa.G>>4)<<4
	_cefa.B = _cefa.B>>4 | (_cefa.B>>4)<<4
	return _cefa
}
func (_beg *Gray2) Copy() Image { return &Gray2{ImageBase: _beg.copy()} }
func (_cfea *ImageBase) HasAlpha() bool {
	if _cfea.Alpha == nil {
		return false
	}
	for _abfc := range _cfea.Alpha {
		if _cfea.Alpha[_abfc] != 0xff {
			return true
		}
	}
	return false
}
func (_babf *Monochrome) AddPadding() (_dfd error) {
	if _fcba := ((_babf.Width * _babf.Height) + 7) >> 3; len(_babf.Data) < _fcba {
		return _a.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(_babf.Data), _fcba)
	}
	_ace := _babf.Width % 8
	if _ace == 0 {
		return nil
	}
	_fdca := _babf.Width / 8
	_adgc := _b.NewReader(_babf.Data)
	_bgd := make([]byte, _babf.Height*_babf.BytesPerLine)
	_babc := _b.NewWriterMSB(_bgd)
	_ccec := make([]byte, _fdca)
	var (
		_fdgf int
		_gdge uint64
	)
	for _fdgf = 0; _fdgf < _babf.Height; _fdgf++ {
		if _, _dfd = _adgc.Read(_ccec); _dfd != nil {
			return _dfd
		}
		if _, _dfd = _babc.Write(_ccec); _dfd != nil {
			return _dfd
		}
		if _gdge, _dfd = _adgc.ReadBits(byte(_ace)); _dfd != nil {
			return _dfd
		}
		if _dfd = _babc.WriteByte(byte(_gdge) << uint(8-_ace)); _dfd != nil {
			return _dfd
		}
	}
	_babf.Data = _babc.Data()
	return nil
}
func (_fgaa *Gray4) At(x, y int) _g.Color { _bagf, _ := _fgaa.ColorAt(x, y); return _bagf }
func _ecgd(_eaag _f.Image, _eccab Image, _afce _f.Rectangle) {
	if _dafe, _bfga := _eaag.(SMasker); _bfga && _dafe.HasAlpha() {
		_eccab.(SMasker).MakeAlpha()
	}
	switch _gcgfb := _eaag.(type) {
	case Gray:
		_fgaff(_gcgfb, _eccab.(RGBA), _afce)
	case NRGBA:
		_dfaaee(_gcgfb, _eccab.(RGBA), _afce)
	case *_f.NYCbCrA:
		_deffe(_gcgfb, _eccab.(RGBA), _afce)
	case CMYK:
		_dedc(_gcgfb, _eccab.(RGBA), _afce)
	case RGBA:
		_deef(_gcgfb, _eccab.(RGBA), _afce)
	case nrgba64:
		_acge(_gcgfb, _eccab.(RGBA), _afce)
	default:
		_gfe(_eaag, _eccab, _afce)
	}
}
func _bacd(_gfd _g.Gray) _g.NRGBA { return _g.NRGBA{R: _gfd.Y, G: _gfd.Y, B: _gfd.Y, A: 0xff} }
func NewImage(width, height, bitsPerComponent, colorComponents int, data, alpha []byte, decode []float64) (Image, error) {
	_afeae := NewImageBase(width, height, bitsPerComponent, colorComponents, data, alpha, decode)
	var _ebag Image
	switch colorComponents {
	case 1:
		switch bitsPerComponent {
		case 1:
			_ebag = &Monochrome{ImageBase: _afeae, ModelThreshold: 0x0f}
		case 2:
			_ebag = &Gray2{ImageBase: _afeae}
		case 4:
			_ebag = &Gray4{ImageBase: _afeae}
		case 8:
			_ebag = &Gray8{ImageBase: _afeae}
		case 16:
			_ebag = &Gray16{ImageBase: _afeae}
		}
	case 3:
		switch bitsPerComponent {
		case 4:
			_ebag = &NRGBA16{ImageBase: _afeae}
		case 8:
			_ebag = &NRGBA32{ImageBase: _afeae}
		case 16:
			_ebag = &NRGBA64{ImageBase: _afeae}
		}
	case 4:
		_ebag = &CMYK32{ImageBase: _afeae}
	}
	if _ebag == nil {
		return nil, ErrInvalidImage
	}
	return _ebag, nil
}
func (_ddc *Gray2) SetGray(x, y int, gray _g.Gray) {
	_fadb := _aeeg(gray)
	_cbad := y * _ddc.BytesPerLine
	_fdaf := _cbad + (x >> 2)
	if _fdaf >= len(_ddc.Data) {
		return
	}
	_afdf := _fadb.Y >> 6
	_ddc.Data[_fdaf] = (_ddc.Data[_fdaf] & (^(0xc0 >> uint(2*((x)&3))))) | (_afdf << uint(6-2*(x&3)))
}
func ColorAtNRGBA16(x, y, width, bytesPerLine int, data, alpha []byte, decode []float64) (_g.NRGBA, error) {
	_faeee := y*bytesPerLine + x*3/2
	if _faeee+1 >= len(data) {
		return _g.NRGBA{}, _cdfde(x, y)
	}
	const (
		_eeba = 0xf
		_ggfg = uint8(0xff)
	)
	_gcggc := _ggfg
	if alpha != nil {
		_becd := y * BytesPerLine(width, 4, 1)
		if _becd < len(alpha) {
			if x%2 == 0 {
				_gcggc = (alpha[_becd] >> uint(4)) & _eeba
			} else {
				_gcggc = alpha[_becd] & _eeba
			}
			_gcggc |= _gcggc << 4
		}
	}
	var _cfcf, _ffge, _aecf uint8
	if x*3%2 == 0 {
		_cfcf = (data[_faeee] >> uint(4)) & _eeba
		_ffge = data[_faeee] & _eeba
		_aecf = (data[_faeee+1] >> uint(4)) & _eeba
	} else {
		_cfcf = data[_faeee] & _eeba
		_ffge = (data[_faeee+1] >> uint(4)) & _eeba
		_aecf = data[_faeee+1] & _eeba
	}
	if len(decode) == 6 {
		_cfcf = uint8(uint32(LinearInterpolate(float64(_cfcf), 0, 15, decode[0], decode[1])) & 0xf)
		_ffge = uint8(uint32(LinearInterpolate(float64(_ffge), 0, 15, decode[2], decode[3])) & 0xf)
		_aecf = uint8(uint32(LinearInterpolate(float64(_aecf), 0, 15, decode[4], decode[5])) & 0xf)
	}
	return _g.NRGBA{R: (_cfcf << 4) | (_cfcf & 0xf), G: (_ffge << 4) | (_ffge & 0xf), B: (_aecf << 4) | (_aecf & 0xf), A: _gcggc}, nil
}

type Gray8 struct{ ImageBase }

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

func _dcaa(_bbea *Monochrome, _egfb, _afcaa, _agd, _gbgd int, _ffa RasterOperator, _gbb *Monochrome, _bade, _ebdeb int) error {
	if _bbea == nil {
		return _cb.New("\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _ffa == PixDst {
		return nil
	}
	switch _ffa {
	case PixClr, PixSet, PixNotDst:
		_begd(_bbea, _egfb, _afcaa, _agd, _gbgd, _ffa)
		return nil
	}
	if _gbb == nil {
		_ag.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _cb.New("\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _fgbg := _gfff(_bbea, _egfb, _afcaa, _agd, _gbgd, _ffa, _gbb, _bade, _ebdeb); _fgbg != nil {
		return _fgbg
	}
	return nil
}
func ColorAtCMYK(x, y, width int, data []byte, decode []float64) (_g.CMYK, error) {
	_gcf := 4 * (y*width + x)
	if _gcf+3 >= len(data) {
		return _g.CMYK{}, _a.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	C := data[_gcf] & 0xff
	M := data[_gcf+1] & 0xff
	Y := data[_gcf+2] & 0xff
	K := data[_gcf+3] & 0xff
	if len(decode) == 8 {
		C = uint8(uint32(LinearInterpolate(float64(C), 0, 255, decode[0], decode[1])) & 0xff)
		M = uint8(uint32(LinearInterpolate(float64(M), 0, 255, decode[2], decode[3])) & 0xff)
		Y = uint8(uint32(LinearInterpolate(float64(Y), 0, 255, decode[4], decode[5])) & 0xff)
		K = uint8(uint32(LinearInterpolate(float64(K), 0, 255, decode[6], decode[7])) & 0xff)
	}
	return _g.CMYK{C: C, M: M, Y: Y, K: K}, nil
}
func (_gceec *RGBA32) Base() *ImageBase { return &_gceec.ImageBase }
func GrayHistogram(g Gray) (_dfeba [256]int) {
	switch _bfeg := g.(type) {
	case Histogramer:
		return _bfeg.Histogram()
	case _f.Image:
		_gbfe := _bfeg.Bounds()
		for _fbaf := 0; _fbaf < _gbfe.Max.X; _fbaf++ {
			for _eaee := 0; _eaee < _gbfe.Max.Y; _eaee++ {
				_dfeba[g.GrayAt(_fbaf, _eaee).Y]++
			}
		}
		return _dfeba
	default:
		return [256]int{}
	}
}
func (_fdee *ImageBase) getByte(_daaf int) (byte, error) {
	if _daaf > len(_fdee.Data)-1 || _daaf < 0 {
		return 0, _a.Errorf("\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", _daaf)
	}
	return _fdee.Data[_daaf], nil
}
func (_fdbd *Gray4) Copy() Image { return &Gray4{ImageBase: _fdbd.copy()} }
func (_acbe *Monochrome) ScaleLow(width, height int) (*Monochrome, error) {
	if width < 0 || height < 0 {
		return nil, _cb.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064t\u0068\u0020\u0061\u006e\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_eda := _bd(width, height)
	_bec := make([]int, height)
	_cggeb := make([]int, width)
	_gfa := float64(_acbe.Width) / float64(width)
	_agc := float64(_acbe.Height) / float64(height)
	for _decd := 0; _decd < height; _decd++ {
		_bec[_decd] = int(_e.Min(_agc*float64(_decd)+0.5, float64(_acbe.Height-1)))
	}
	for _dcdf := 0; _dcdf < width; _dcdf++ {
		_cggeb[_dcdf] = int(_e.Min(_gfa*float64(_dcdf)+0.5, float64(_acbe.Width-1)))
	}
	_ccca := -1
	_bcdf := byte(0)
	for _edbd := 0; _edbd < height; _edbd++ {
		_gec := _bec[_edbd] * _acbe.BytesPerLine
		_acdf := _edbd * _eda.BytesPerLine
		for _decb := 0; _decb < width; _decb++ {
			_adgg := _cggeb[_decb]
			if _adgg != _ccca {
				_bcdf = _acbe.getBit(_gec, _adgg)
				if _bcdf != 0 {
					_eda.setBit(_acdf, _decb)
				}
				_ccca = _adgg
			} else {
				if _bcdf != 0 {
					_eda.setBit(_acdf, _decb)
				}
			}
		}
	}
	return _eda, nil
}
func (_bggcg *RGBA32) ColorAt(x, y int) (_g.Color, error) {
	return ColorAtRGBA32(x, y, _bggcg.Width, _bggcg.Data, _bggcg.Alpha, _bggcg.Decode)
}
func (_bcgc *Monochrome) Validate() error {
	if len(_bcgc.Data) != _bcgc.Height*_bcgc.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
