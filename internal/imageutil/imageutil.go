package imageutil

import (
	_f "encoding/binary"
	_e "errors"
	_db "fmt"
	_ee "image"
	_d "image/color"
	_ed "image/draw"
	_g "math"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_eeg "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
)

func (_dcge *NRGBA32) Copy() Image     { return &NRGBA32{ImageBase: _dcge.copy()} }
func (_daef *Gray16) Base() *ImageBase { return &_daef.ImageBase }
func (_badb *monochromeThresholdConverter) Convert(img _ee.Image) (Image, error) {
	if _acfd, _bagf := img.(*Monochrome); _bagf {
		return _acfd.Copy(), nil
	}
	_dec := img.Bounds()
	_efgd, _beda := NewImage(_dec.Max.X, _dec.Max.Y, 1, 1, nil, nil, nil)
	if _beda != nil {
		return nil, _beda
	}
	_efgd.(*Monochrome).ModelThreshold = _badb.Threshold
	for _agf := 0; _agf < _dec.Max.X; _agf++ {
		for _cfgef := 0; _cfgef < _dec.Max.Y; _cfgef++ {
			_bccb := img.At(_agf, _cfgef)
			_efgd.Set(_agf, _cfgef, _bccb)
		}
	}
	return _efgd, nil
}
func (_fadb *Monochrome) Validate() error {
	if len(_fadb.Data) != _fadb.Height*_fadb.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_aagb *Gray2) GrayAt(x, y int) _d.Gray {
	_acc, _ := ColorAtGray2BPC(x, y, _aagb.BytesPerLine, _aagb.Data, _aagb.Decode)
	return _acc
}
func (_cagb *ImageBase) HasAlpha() bool {
	if _cagb.Alpha == nil {
		return false
	}
	for _gbce := range _cagb.Alpha {
		if _cagb.Alpha[_gbce] != 0xff {
			return true
		}
	}
	return false
}
func (_fbfa *Monochrome) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _fbfa.Width, Y: _fbfa.Height}}
}

var _ Image = &Gray2{}

func (_aecaf *Gray2) Histogram() (_aea [256]int) {
	for _gede := 0; _gede < _aecaf.Width; _gede++ {
		for _gccf := 0; _gccf < _aecaf.Height; _gccf++ {
			_aea[_aecaf.GrayAt(_gede, _gccf).Y]++
		}
	}
	return _aea
}
func _eeac(_gcecd int, _fgag int) error {
	return _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", _gcecd, _fgag)
}
func _cdcc(_gfbe _d.NRGBA64) _d.NRGBA {
	return _d.NRGBA{R: uint8(_gfbe.R >> 8), G: uint8(_gfbe.G >> 8), B: uint8(_gfbe.B >> 8), A: uint8(_gfbe.A >> 8)}
}
func _efaga(_cfbfd []byte, _dace Image) error {
	_ebea := true
	for _efda := 0; _efda < len(_cfbfd); _efda++ {
		if _cfbfd[_efda] != 0xff {
			_ebea = false
			break
		}
	}
	if _ebea {
		switch _febe := _dace.(type) {
		case *NRGBA32:
			_febe.Alpha = nil
		case *NRGBA64:
			_febe.Alpha = nil
		default:
			return _db.Errorf("i\u006ete\u0072n\u0061l\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020i\u006d\u0061\u0067\u0065\u0020s\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070e\u0020\u002a\u004eRGB\u0041\u0033\u0032\u0020\u006f\u0072 \u002a\u004e\u0052\u0047\u0042\u0041\u0036\u0034\u0020\u0062\u0075\u0074 \u0069s\u003a\u0020\u0025\u0054", _dace)
		}
	}
	return nil
}
func (_gaa *CMYK32) Set(x, y int, c _d.Color) {
	_gfcb := 4 * (y*_gaa.Width + x)
	if _gfcb+3 >= len(_gaa.Data) {
		return
	}
	_feab := _d.CMYKModel.Convert(c).(_d.CMYK)
	_gaa.Data[_gfcb] = _feab.C
	_gaa.Data[_gfcb+1] = _feab.M
	_gaa.Data[_gfcb+2] = _feab.Y
	_gaa.Data[_gfcb+3] = _feab.K
}

var _ _ee.Image = &Gray4{}

func _fafc(_fadf _d.CMYK) _d.Gray {
	_fbaf, _bcf, _cgb := _d.CMYKToRGB(_fadf.C, _fadf.M, _fadf.Y, _fadf.K)
	_gafb := (19595*uint32(_fbaf) + 38470*uint32(_bcf) + 7471*uint32(_cgb) + 1<<7) >> 16
	return _d.Gray{Y: uint8(_gafb)}
}
func (_bebg *Gray8) ColorModel() _d.Model { return _d.GrayModel }
func _efe() (_gbg [256]uint64) {
	for _ggf := 0; _ggf < 256; _ggf++ {
		if _ggf&0x01 != 0 {
			_gbg[_ggf] |= 0xff
		}
		if _ggf&0x02 != 0 {
			_gbg[_ggf] |= 0xff00
		}
		if _ggf&0x04 != 0 {
			_gbg[_ggf] |= 0xff0000
		}
		if _ggf&0x08 != 0 {
			_gbg[_ggf] |= 0xff000000
		}
		if _ggf&0x10 != 0 {
			_gbg[_ggf] |= 0xff00000000
		}
		if _ggf&0x20 != 0 {
			_gbg[_ggf] |= 0xff0000000000
		}
		if _ggf&0x40 != 0 {
			_gbg[_ggf] |= 0xff000000000000
		}
		if _ggf&0x80 != 0 {
			_gbg[_ggf] |= 0xff00000000000000
		}
	}
	return _gbg
}
func MonochromeThresholdConverter(threshold uint8) ColorConverter {
	return &monochromeThresholdConverter{Threshold: threshold}
}
func (_ggfg *Monochrome) At(x, y int) _d.Color { _dggb, _ := _ggfg.ColorAt(x, y); return _dggb }
func (_ebfbb *ImageBase) setEightBytes(_cgda int, _gagd uint64) error {
	_bfag := _ebfbb.BytesPerLine - (_cgda % _ebfbb.BytesPerLine)
	if _ebfbb.BytesPerLine != _ebfbb.Width>>3 {
		_bfag--
	}
	if _bfag >= 8 {
		return _ebfbb.setEightFullBytes(_cgda, _gagd)
	}
	return _ebfbb.setEightPartlyBytes(_cgda, _bfag, _gagd)
}
func (_fab *Gray16) Copy() Image      { return &Gray16{ImageBase: _fab.copy()} }
func (_adce *Gray8) Base() *ImageBase { return &_adce.ImageBase }
func (_egg *Gray4) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtGray4BPC(x, y, _egg.BytesPerLine, _egg.Data, _egg.Decode)
}
func _fbbg(_dbf Gray, _fbc nrgba64, _afg _ee.Rectangle) {
	for _abed := 0; _abed < _afg.Max.X; _abed++ {
		for _cce := 0; _cce < _afg.Max.Y; _cce++ {
			_geb := _aggf(_fbc.NRGBA64At(_abed, _cce))
			_dbf.SetGray(_abed, _cce, _geb)
		}
	}
}
func (_dgc monochromeModel) Convert(c _d.Color) _d.Color {
	_aec := _d.GrayModel.Convert(c).(_d.Gray)
	return _beg(_aec, _dgc)
}
func (_aae *NRGBA64) SetNRGBA64(x, y int, c _d.NRGBA64) {
	_acdf := (y*_aae.Width + x) * 2
	_fgbcd := _acdf * 3
	if _fgbcd+5 >= len(_aae.Data) {
		return
	}
	_aae.setNRGBA64(_fgbcd, c, _acdf)
}
func (_cedf *NRGBA64) Set(x, y int, c _d.Color) {
	_aade := (y*_cedf.Width + x) * 2
	_eggc := _aade * 3
	if _eggc+5 >= len(_cedf.Data) {
		return
	}
	_gbde := _d.NRGBA64Model.Convert(c).(_d.NRGBA64)
	_cedf.setNRGBA64(_eggc, _gbde, _aade)
}

var _bcdf [256]uint8

func (_egfa *NRGBA64) setNRGBA64(_beaa int, _eacf _d.NRGBA64, _edffc int) {
	_egfa.Data[_beaa] = uint8(_eacf.R >> 8)
	_egfa.Data[_beaa+1] = uint8(_eacf.R & 0xff)
	_egfa.Data[_beaa+2] = uint8(_eacf.G >> 8)
	_egfa.Data[_beaa+3] = uint8(_eacf.G & 0xff)
	_egfa.Data[_beaa+4] = uint8(_eacf.B >> 8)
	_egfa.Data[_beaa+5] = uint8(_eacf.B & 0xff)
	if _edffc+1 < len(_egfa.Alpha) {
		_egfa.Alpha[_edffc] = uint8(_eacf.A >> 8)
		_egfa.Alpha[_edffc+1] = uint8(_eacf.A & 0xff)
	}
}
func (_gfdf *Monochrome) Base() *ImageBase { return &_gfdf.ImageBase }
func _eea(_ce, _cb *Monochrome) (_bag error) {
	_cag := _cb.BytesPerLine
	_fe := _ce.BytesPerLine
	_ff := _cb.BytesPerLine*4 - _ce.BytesPerLine
	var (
		_cf, _gd                             byte
		_fc                                  uint32
		_gg, _caf, _bd, _acd, _fd, _fa, _gaf int
	)
	for _bd = 0; _bd < _cb.Height; _bd++ {
		_gg = _bd * _cag
		_caf = 4 * _bd * _fe
		for _acd = 0; _acd < _cag; _acd++ {
			_cf = _cb.Data[_gg+_acd]
			_fc = _acf[_cf]
			_fa = _caf + _acd*4
			if _ff != 0 && (_acd+1)*4 > _ce.BytesPerLine {
				for _fd = _ff; _fd > 0; _fd-- {
					_gd = byte((_fc >> uint(_fd*8)) & 0xff)
					_gaf = _fa + (_ff - _fd)
					if _bag = _ce.setByte(_gaf, _gd); _bag != nil {
						return _bag
					}
				}
			} else if _bag = _ce.setFourBytes(_fa, _fc); _bag != nil {
				return _bag
			}
			if _bag = _ce.setFourBytes(_caf+_acd*4, _acf[_cb.Data[_gg+_acd]]); _bag != nil {
				return _bag
			}
		}
		for _fd = 1; _fd < 4; _fd++ {
			for _acd = 0; _acd < _fe; _acd++ {
				if _bag = _ce.setByte(_caf+_fd*_fe+_acd, _ce.Data[_caf+_acd]); _bag != nil {
					return _bag
				}
			}
		}
	}
	return nil
}
func RasterOperation(dest *Monochrome, dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _gbcc(dest, dx, dy, dw, dh, op, src, sx, sy)
}
func (_acdd *Gray2) Validate() error {
	if len(_acdd.Data) != _acdd.Height*_acdd.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

var _ _ee.Image = &NRGBA64{}

func ColorAtNRGBA64(x, y, width int, data, alpha []byte, decode []float64) (_d.NRGBA64, error) {
	_dbac := (y*width + x) * 2
	_ecgc := _dbac * 3
	if _ecgc+5 >= len(data) {
		return _d.NRGBA64{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	const _gdfgg = 0xffff
	_bdge := uint16(_gdfgg)
	if alpha != nil && len(alpha) > _dbac+1 {
		_bdge = uint16(alpha[_dbac])<<8 | uint16(alpha[_dbac+1])
	}
	_egfd := uint16(data[_ecgc])<<8 | uint16(data[_ecgc+1])
	_dcea := uint16(data[_ecgc+2])<<8 | uint16(data[_ecgc+3])
	_daedd := uint16(data[_ecgc+4])<<8 | uint16(data[_ecgc+5])
	if len(decode) == 6 {
		_egfd = uint16(uint64(LinearInterpolate(float64(_egfd), 0, 65535, decode[0], decode[1])) & _gdfgg)
		_dcea = uint16(uint64(LinearInterpolate(float64(_dcea), 0, 65535, decode[2], decode[3])) & _gdfgg)
		_daedd = uint16(uint64(LinearInterpolate(float64(_daedd), 0, 65535, decode[4], decode[5])) & _gdfgg)
	}
	return _d.NRGBA64{R: _egfd, G: _dcea, B: _daedd, A: _bdge}, nil
}
func _gccd(_dcfeb RGBA, _eebc Gray, _cefc _ee.Rectangle) {
	for _cfea := 0; _cfea < _cefc.Max.X; _cfea++ {
		for _bdef := 0; _bdef < _cefc.Max.Y; _bdef++ {
			_cacc := _aced(_dcfeb.RGBAAt(_cfea, _bdef))
			_eebc.SetGray(_cfea, _bdef, _cacc)
		}
	}
}
func (_eade *Gray4) Set(x, y int, c _d.Color) {
	if x >= _eade.Width || y >= _eade.Height {
		return
	}
	_dcgg := Gray4Model.Convert(c).(_d.Gray)
	_eade.setGray(x, y, _dcgg)
}
func _badd(_gged _ee.Image) (Image, error) {
	if _fede, _cabg := _gged.(*Gray2); _cabg {
		return _fede.Copy(), nil
	}
	_dgfb := _gged.Bounds()
	_dgag, _aaad := NewImage(_dgfb.Max.X, _dgfb.Max.Y, 2, 1, nil, nil, nil)
	if _aaad != nil {
		return nil, _aaad
	}
	_gbgd(_gged, _dgag, _dgfb)
	return _dgag, nil
}

var _ Image = &Gray8{}

func (_fbcd *NRGBA32) setRGBA(_afdac int, _fce _d.NRGBA) {
	_dacf := 3 * _afdac
	_fbcd.Data[_dacf] = _fce.R
	_fbcd.Data[_dacf+1] = _fce.G
	_fbcd.Data[_dacf+2] = _fce.B
	if _afdac < len(_fbcd.Alpha) {
		_fbcd.Alpha[_afdac] = _fce.A
	}
}
func _gdea(_eggd CMYK, _fcfg RGBA, _cgdb _ee.Rectangle) {
	for _gbfcf := 0; _gbfcf < _cgdb.Max.X; _gbfcf++ {
		for _fecf := 0; _fecf < _cgdb.Max.Y; _fecf++ {
			_feae := _eggd.CMYKAt(_gbfcf, _fecf)
			_fcfg.SetRGBA(_gbfcf, _fecf, _adfe(_feae))
		}
	}
}
func _fega(_bad _ee.Image, _addaa Image, _gcd _ee.Rectangle) {
	for _dfg := 0; _dfg < _gcd.Max.X; _dfg++ {
		for _efec := 0; _efec < _gcd.Max.Y; _efec++ {
			_adcb := _bad.At(_dfg, _efec)
			_addaa.Set(_dfg, _efec, _adcb)
		}
	}
}
func (_gafd *Monochrome) ScaleLow(width, height int) (*Monochrome, error) {
	if width < 0 || height < 0 {
		return nil, _e.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064t\u0068\u0020\u0061\u006e\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_dgfe := _bed(width, height)
	_bdbg := make([]int, height)
	_fgfb := make([]int, width)
	_gfdg := float64(_gafd.Width) / float64(width)
	_agaa := float64(_gafd.Height) / float64(height)
	for _dfgd := 0; _dfgd < height; _dfgd++ {
		_bdbg[_dfgd] = int(_g.Min(_agaa*float64(_dfgd)+0.5, float64(_gafd.Height-1)))
	}
	for _fbe := 0; _fbe < width; _fbe++ {
		_fgfb[_fbe] = int(_g.Min(_gfdg*float64(_fbe)+0.5, float64(_gafd.Width-1)))
	}
	_cdd := -1
	_bgg := byte(0)
	for _ecga := 0; _ecga < height; _ecga++ {
		_bgb := _bdbg[_ecga] * _gafd.BytesPerLine
		_dgb := _ecga * _dgfe.BytesPerLine
		for _gff := 0; _gff < width; _gff++ {
			_daf := _fgfb[_gff]
			if _daf != _cdd {
				_bgg = _gafd.getBit(_bgb, _daf)
				if _bgg != 0 {
					_dgfe.setBit(_dgb, _gff)
				}
				_cdd = _daf
			} else {
				if _bgg != 0 {
					_dgfe.setBit(_dgb, _gff)
				}
			}
		}
	}
	return _dgfe, nil
}
func _bdgc(_dee uint) uint {
	var _eec uint
	for _dee != 0 {
		_dee >>= 1
		_eec++
	}
	return _eec - 1
}

type Image interface {
	_ed.Image
	Base() *ImageBase
	Copy() Image
	Pix() []byte
	ColorAt(_ccbdg, _aaag int) (_d.Color, error)
	Validate() error
}

func _gecf(_bfg _d.Gray) _d.NRGBA { return _d.NRGBA{R: _bfg.Y, G: _bfg.Y, B: _bfg.Y, A: 0xff} }
func _dfeg(_eda _d.CMYK) _d.NRGBA {
	_aag, _ccg, _gce := _d.CMYKToRGB(_eda.C, _eda.M, _eda.Y, _eda.K)
	return _d.NRGBA{R: _aag, G: _ccg, B: _gce, A: 0xff}
}
func (_abge *ImageBase) setByte(_geeef int, _ccdf byte) error {
	if _geeef > len(_abge.Data)-1 {
		return _e.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_abge.Data[_geeef] = _ccdf
	return nil
}
func (_acea *Monochrome) setIndexedBit(_agce int) { _acea.Data[(_agce >> 3)] |= 0x80 >> uint(_agce&7) }
func (_gdc *Gray2) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _gdc.Width, Y: _gdc.Height}}
}

var (
	Gray2Model   = _d.ModelFunc(_efef)
	Gray4Model   = _d.ModelFunc(_fdfa)
	NRGBA16Model = _d.ModelFunc(_fefe)
)

func (_egae *Monochrome) ColorModel() _d.Model { return MonochromeModel(_egae.ModelThreshold) }
func (_cfga *ImageBase) setEightPartlyBytes(_bfea, _aecf int, _bcegb uint64) (_faba error) {
	var (
		_cacb byte
		_babd int
	)
	for _gfga := 1; _gfga <= _aecf; _gfga++ {
		_babd = 64 - _gfga*8
		_cacb = byte(_bcegb >> uint(_babd) & 0xff)
		if _faba = _cfga.setByte(_bfea+_gfga-1, _cacb); _faba != nil {
			return _faba
		}
	}
	_cfbf := _cfga.BytesPerLine*8 - _cfga.Width
	if _cfbf == 0 {
		return nil
	}
	_babd -= 8
	_cacb = byte(_bcegb>>uint(_babd)&0xff) << uint(_cfbf)
	if _faba = _cfga.setByte(_bfea+_aecf, _cacb); _faba != nil {
		return _faba
	}
	return nil
}
func _egc(_gafg _d.RGBA) _d.NRGBA {
	switch _gafg.A {
	case 0xff:
		return _d.NRGBA{R: _gafg.R, G: _gafg.G, B: _gafg.B, A: 0xff}
	case 0x00:
		return _d.NRGBA{}
	default:
		_fed, _gca, _agae, _fgef := _gafg.RGBA()
		_fed = (_fed * 0xffff) / _fgef
		_gca = (_gca * 0xffff) / _fgef
		_agae = (_agae * 0xffff) / _fgef
		return _d.NRGBA{R: uint8(_fed >> 8), G: uint8(_gca >> 8), B: uint8(_agae >> 8), A: uint8(_fgef >> 8)}
	}
}

type ColorConverter interface {
	Convert(_gcg _ee.Image) (Image, error)
}

func AddDataPadding(width, height, bitsPerComponent, colorComponents int, data []byte) ([]byte, error) {
	_gdba := BytesPerLine(width, bitsPerComponent, colorComponents)
	if _gdba == width*colorComponents*bitsPerComponent/8 {
		return data, nil
	}
	_bcbe := width * colorComponents * bitsPerComponent
	_febb := _gdba * 8
	_dbgd := 8 - (_febb - _bcbe)
	_dedd := _eeg.NewReader(data)
	_efag := _gdba - 1
	_gdff := make([]byte, _efag)
	_ccbdge := make([]byte, height*_gdba)
	_fcbg := _eeg.NewWriterMSB(_ccbdge)
	var _gdaf uint64
	var _cfdc error
	for _fgbg := 0; _fgbg < height; _fgbg++ {
		_, _cfdc = _dedd.Read(_gdff)
		if _cfdc != nil {
			return nil, _cfdc
		}
		_, _cfdc = _fcbg.Write(_gdff)
		if _cfdc != nil {
			return nil, _cfdc
		}
		_gdaf, _cfdc = _dedd.ReadBits(byte(_dbgd))
		if _cfdc != nil {
			return nil, _cfdc
		}
		_, _cfdc = _fcbg.WriteBits(_gdaf, _dbgd)
		if _cfdc != nil {
			return nil, _cfdc
		}
		_fcbg.FinishByte()
	}
	return _ccbdge, nil
}
func _cadc(_gdad *Monochrome, _eae, _afbb, _eadg, _ceaf int, _cbdc RasterOperator, _adaf *Monochrome, _gfda, _bfac int) error {
	var (
		_dgbbf        bool
		_bebb         bool
		_cdbb         byte
		_eega         int
		_ffef         int
		_eba          int
		_edce         int
		_efdc         bool
		_dfce         int
		_fged         int
		_ggg          int
		_ccbc         bool
		_caebc        byte
		_febbg        int
		_ecdf         int
		_afcc         int
		_ggefd        byte
		_ccad         int
		_cdea         int
		_cbccg        uint
		_egcd         uint
		_eagfa        byte
		_gdfgc        shift
		_bggd         bool
		_ceeg         bool
		_bfeb, _gbfgg int
	)
	if _gfda&7 != 0 {
		_cdea = 8 - (_gfda & 7)
	}
	if _eae&7 != 0 {
		_ffef = 8 - (_eae & 7)
	}
	if _cdea == 0 && _ffef == 0 {
		_eagfa = _afge[0]
	} else {
		if _ffef > _cdea {
			_cbccg = uint(_ffef - _cdea)
		} else {
			_cbccg = uint(8 - (_cdea - _ffef))
		}
		_egcd = 8 - _cbccg
		_eagfa = _afge[_cbccg]
	}
	if (_eae & 7) != 0 {
		_dgbbf = true
		_eega = 8 - (_eae & 7)
		_cdbb = _afge[_eega]
		_eba = _gdad.BytesPerLine*_afbb + (_eae >> 3)
		_edce = _adaf.BytesPerLine*_bfac + (_gfda >> 3)
		_ccad = 8 - (_gfda & 7)
		if _eega > _ccad {
			_gdfgc = _egfe
			if _eadg >= _cdea {
				_bggd = true
			}
		} else {
			_gdfgc = _cdfb
		}
	}
	if _eadg < _eega {
		_bebb = true
		_cdbb &= _gbdbe[8-_eega+_eadg]
	}
	if !_bebb {
		_dfce = (_eadg - _eega) >> 3
		if _dfce != 0 {
			_efdc = true
			_fged = _gdad.BytesPerLine*_afbb + ((_eae + _ffef) >> 3)
			_ggg = _adaf.BytesPerLine*_bfac + ((_gfda + _ffef) >> 3)
		}
	}
	_febbg = (_eae + _eadg) & 7
	if !(_bebb || _febbg == 0) {
		_ccbc = true
		_caebc = _gbdbe[_febbg]
		_ecdf = _gdad.BytesPerLine*_afbb + ((_eae + _ffef) >> 3) + _dfce
		_afcc = _adaf.BytesPerLine*_bfac + ((_gfda + _ffef) >> 3) + _dfce
		if _febbg > int(_egcd) {
			_ceeg = true
		}
	}
	switch _cbdc {
	case PixSrc:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], _ggefd, _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] = _ggefd
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], _ggefd, _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixNotSrc:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], ^_ggefd, _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] = ^_ggefd
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], ^_ggefd, _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixSrcOrDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], _ggefd|_gdad.Data[_eba], _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] |= _ggefd
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], _ggefd|_gdad.Data[_ecdf], _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixSrcAndDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], _ggefd&_gdad.Data[_eba], _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] &= _ggefd
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], _ggefd&_gdad.Data[_ecdf], _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixSrcXorDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], _ggefd^_gdad.Data[_eba], _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] ^= _ggefd
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], _ggefd^_gdad.Data[_ecdf], _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], ^_ggefd|_gdad.Data[_eba], _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] |= ^_ggefd
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], ^_ggefd|_gdad.Data[_ecdf], _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], ^_ggefd&_gdad.Data[_eba], _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] &= ^_ggefd
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], ^_ggefd&_gdad.Data[_ecdf], _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], _ggefd|^_gdad.Data[_eba], _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] = _ggefd | ^_gdad.Data[_fged+_gbfgg]
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], _ggefd|^_gdad.Data[_ecdf], _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], _ggefd&^_gdad.Data[_eba], _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] = _ggefd &^ _gdad.Data[_fged+_gbfgg]
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], _ggefd&^_gdad.Data[_ecdf], _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], ^(_ggefd | _gdad.Data[_eba]), _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] = ^(_ggefd | _gdad.Data[_fged+_gbfgg])
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], ^(_ggefd | _gdad.Data[_ecdf]), _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], ^(_ggefd & _gdad.Data[_eba]), _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] = ^(_ggefd & _gdad.Data[_fged+_gbfgg])
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], ^(_ggefd & _gdad.Data[_ecdf]), _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		if _dgbbf {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				if _gdfgc == _egfe {
					_ggefd = _adaf.Data[_edce] << _cbccg
					if _bggd {
						_ggefd = _ggdf(_ggefd, _adaf.Data[_edce+1]>>_egcd, _eagfa)
					}
				} else {
					_ggefd = _adaf.Data[_edce] >> _egcd
				}
				_gdad.Data[_eba] = _ggdf(_gdad.Data[_eba], ^(_ggefd ^ _gdad.Data[_eba]), _cdbb)
				_eba += _gdad.BytesPerLine
				_edce += _adaf.BytesPerLine
			}
		}
		if _efdc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				for _gbfgg = 0; _gbfgg < _dfce; _gbfgg++ {
					_ggefd = _ggdf(_adaf.Data[_ggg+_gbfgg]<<_cbccg, _adaf.Data[_ggg+_gbfgg+1]>>_egcd, _eagfa)
					_gdad.Data[_fged+_gbfgg] = ^(_ggefd ^ _gdad.Data[_fged+_gbfgg])
				}
				_fged += _gdad.BytesPerLine
				_ggg += _adaf.BytesPerLine
			}
		}
		if _ccbc {
			for _bfeb = 0; _bfeb < _ceaf; _bfeb++ {
				_ggefd = _adaf.Data[_afcc] << _cbccg
				if _ceeg {
					_ggefd = _ggdf(_ggefd, _adaf.Data[_afcc+1]>>_egcd, _eagfa)
				}
				_gdad.Data[_ecdf] = _ggdf(_gdad.Data[_ecdf], ^(_ggefd ^ _gdad.Data[_ecdf]), _caebc)
				_ecdf += _gdad.BytesPerLine
				_afcc += _adaf.BytesPerLine
			}
		}
	default:
		_a.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _cbdc)
		return _e.New("\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}
func (_caeg *NRGBA32) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtNRGBA32(x, y, _caeg.Width, _caeg.Data, _caeg.Alpha, _caeg.Decode)
}
func (_gdbf *NRGBA64) Validate() error {
	if len(_gdbf.Data) != 3*2*_gdbf.Width*_gdbf.Height {
		return _e.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func (_ccfe *Monochrome) ReduceBinary(factor float64) (*Monochrome, error) {
	_gece := _bdgc(uint(factor))
	if !IsPowerOf2(uint(factor)) {
		_gece++
	}
	_dfd := make([]int, _gece)
	for _acfcb := range _dfd {
		_dfd[_acfcb] = 4
	}
	_cedb, _abef := _dgf(_ccfe, _dfd...)
	if _abef != nil {
		return nil, _abef
	}
	return _cedb, nil
}
func _bgfe(_ddbb nrgba64, _fggd RGBA, _eeec _ee.Rectangle) {
	for _cdcg := 0; _cdcg < _eeec.Max.X; _cdcg++ {
		for _deda := 0; _deda < _eeec.Max.Y; _deda++ {
			_cbaca := _ddbb.NRGBA64At(_cdcg, _deda)
			_fggd.SetRGBA(_cdcg, _deda, _ebb(_cbaca))
		}
	}
}
func IsPowerOf2(n uint) bool { return n > 0 && (n&(n-1)) == 0 }
func _cffb(_ebed CMYK, _gcff NRGBA, _ecfe _ee.Rectangle) {
	for _abeda := 0; _abeda < _ecfe.Max.X; _abeda++ {
		for _egdbf := 0; _egdbf < _ecfe.Max.Y; _egdbf++ {
			_efaec := _ebed.CMYKAt(_abeda, _egdbf)
			_gcff.SetNRGBA(_abeda, _egdbf, _dfeg(_efaec))
		}
	}
}

type Gray4 struct{ ImageBase }

var _ _ee.Image = &Monochrome{}

func (_cec *CMYK32) At(x, y int) _d.Color { _bge, _ := _cec.ColorAt(x, y); return _bge }
func (_fbgb *NRGBA64) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtNRGBA64(x, y, _fbgb.Width, _fbgb.Data, _fbgb.Alpha, _fbgb.Decode)
}
func _bcbc(_gfab RGBA, _fbbe NRGBA, _egbdf _ee.Rectangle) {
	for _eaed := 0; _eaed < _egbdf.Max.X; _eaed++ {
		for _adeg := 0; _adeg < _egbdf.Max.Y; _adeg++ {
			_edac := _gfab.RGBAAt(_eaed, _adeg)
			_fbbe.SetNRGBA(_eaed, _adeg, _egc(_edac))
		}
	}
}
func _bed(_agc, _cdc int) *Monochrome {
	return &Monochrome{ImageBase: NewImageBase(_agc, _cdc, 1, 1, nil, nil, nil), ModelThreshold: 0x0f}
}
func init() { _dbc() }
func _b(_fb *Monochrome, _dge int) (*Monochrome, error) {
	if _fb == nil {
		return nil, _e.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _dge == 1 {
		return _fb.copy(), nil
	}
	if !IsPowerOf2(uint(_dge)) {
		return nil, _db.Errorf("\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006ci\u0064 \u0065x\u0070a\u006e\u0064\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _dge)
	}
	_fbg := _bc(_dge)
	return _edg(_fb, _dge, _fbg)
}
func (_fcda *Gray2) SetGray(x, y int, gray _d.Gray) {
	_fgdg := _gea(gray)
	_dccfbe := y * _fcda.BytesPerLine
	_eead := _dccfbe + (x >> 2)
	if _eead >= len(_fcda.Data) {
		return
	}
	_ebbd := _fgdg.Y >> 6
	_fcda.Data[_eead] = (_fcda.Data[_eead] & (^(0xc0 >> uint(2*((x)&3))))) | (_ebbd << uint(6-2*(x&3)))
}
func _fbad(_bdae Gray, _ccgb NRGBA, _dcbg _ee.Rectangle) {
	for _acfe := 0; _acfe < _dcbg.Max.X; _acfe++ {
		for _defg := 0; _defg < _dcbg.Max.Y; _defg++ {
			_cadfa := _bdae.GrayAt(_acfe, _defg)
			_ccgb.SetNRGBA(_acfe, _defg, _gecf(_cadfa))
		}
	}
}

type monochromeThresholdConverter struct {
	Threshold uint8
}

func (_caba *NRGBA16) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _caba.Width, Y: _caba.Height}}
}
func (_dbfg *RGBA32) Base() *ImageBase { return &_dbfg.ImageBase }
func _aggf(_cca _d.NRGBA64) _d.Gray {
	var _fcdd _d.NRGBA64
	if _cca == _fcdd {
		return _d.Gray{Y: 0xff}
	}
	_deg, _bcec, _dbe, _ := _cca.RGBA()
	_ffbe := (19595*_deg + 38470*_bcec + 7471*_dbe + 1<<15) >> 24
	return _d.Gray{Y: uint8(_ffbe)}
}

var _ _ee.Image = &NRGBA16{}

func (_abefg *NRGBA16) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtNRGBA16(x, y, _abefg.Width, _abefg.BytesPerLine, _abefg.Data, _abefg.Alpha, _abefg.Decode)
}
func ColorAtGray8BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_d.Gray, error) {
	_dbff := y*bytesPerLine + x
	if _dbff >= len(data) {
		return _d.Gray{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_fddb := data[_dbff]
	if len(decode) == 2 {
		_fddb = uint8(uint32(LinearInterpolate(float64(_fddb), 0, 255, decode[0], decode[1])) & 0xff)
	}
	return _d.Gray{Y: _fddb}, nil
}
func _dba(_dcd _d.Gray) _d.RGBA { return _d.RGBA{R: _dcd.Y, G: _dcd.Y, B: _dcd.Y, A: 0xff} }
func _dbbb(_fgda RGBA, _fegg CMYK, _ecd _ee.Rectangle) {
	for _fbb := 0; _fbb < _ecd.Max.X; _fbb++ {
		for _bfa := 0; _bfa < _ecd.Max.Y; _bfa++ {
			_dccf := _fgda.RGBAAt(_fbb, _bfa)
			_fegg.SetCMYK(_fbb, _bfa, _fae(_dccf))
		}
	}
}
func (_eed *Gray8) SetGray(x, y int, g _d.Gray) {
	_edgf := y*_eed.BytesPerLine + x
	if _edgf > len(_eed.Data)-1 {
		return
	}
	_eed.Data[_edgf] = g.Y
}
func (_cgdc *Monochrome) copy() *Monochrome {
	_daeg := _bed(_cgdc.Width, _cgdc.Height)
	_daeg.ModelThreshold = _cgdc.ModelThreshold
	_daeg.Data = make([]byte, len(_cgdc.Data))
	copy(_daeg.Data, _cgdc.Data)
	if len(_cgdc.Decode) != 0 {
		_daeg.Decode = make([]float64, len(_cgdc.Decode))
		copy(_daeg.Decode, _cgdc.Decode)
	}
	if len(_cgdc.Alpha) != 0 {
		_daeg.Alpha = make([]byte, len(_cgdc.Alpha))
		copy(_daeg.Alpha, _cgdc.Alpha)
	}
	return _daeg
}
func _ec(_fcd, _cee *Monochrome) (_cad error) {
	_cdg := _cee.BytesPerLine
	_eg := _fcd.BytesPerLine
	var _bba, _fad, _de, _feg, _cg int
	for _de = 0; _de < _cee.Height; _de++ {
		_bba = _de * _cdg
		_fad = 8 * _de * _eg
		for _feg = 0; _feg < _cdg; _feg++ {
			if _cad = _fcd.setEightBytes(_fad+_feg*8, _bg[_cee.Data[_bba+_feg]]); _cad != nil {
				return _cad
			}
		}
		for _cg = 1; _cg < 8; _cg++ {
			for _feg = 0; _feg < _eg; _feg++ {
				if _cad = _fcd.setByte(_fad+_cg*_eg+_feg, _fcd.Data[_fad+_feg]); _cad != nil {
					return _cad
				}
			}
		}
	}
	return nil
}
func ImgToGray(i _ee.Image) *_ee.Gray {
	if _bcgf, _fcbc := i.(*_ee.Gray); _fcbc {
		return _bcgf
	}
	_adegd := i.Bounds()
	_fbab := _ee.NewGray(_adegd)
	for _bfcf := 0; _bfcf < _adegd.Max.X; _bfcf++ {
		for _gdbd := 0; _gdbd < _adegd.Max.Y; _gdbd++ {
			_afe := i.At(_bfcf, _gdbd)
			_fbab.Set(_bfcf, _gdbd, _afe)
		}
	}
	return _fbab
}

var _ Image = &NRGBA16{}

func (_cff *Monochrome) setBit(_cfgf, _gfeg int) {
	_cff.Data[_cfgf+(_gfeg>>3)] |= 0x80 >> uint(_gfeg&7)
}
func (_gbed *NRGBA16) SetNRGBA(x, y int, c _d.NRGBA) {
	_dbca := y*_gbed.BytesPerLine + x*3/2
	if _dbca+1 >= len(_gbed.Data) {
		return
	}
	c = _faef(c)
	_gbed.setNRGBA(x, y, _dbca, c)
}
func _bc(_ebe int) []uint {
	var _ggfd []uint
	_fcf := _ebe
	_da := _fcf / 8
	if _da != 0 {
		for _dae := 0; _dae < _da; _dae++ {
			_ggfd = append(_ggfd, 8)
		}
		_bbf := _fcf % 8
		_fcf = 0
		if _bbf != 0 {
			_fcf = _bbf
		}
	}
	_adc := _fcf / 4
	if _adc != 0 {
		for _dca := 0; _dca < _adc; _dca++ {
			_ggfd = append(_ggfd, 4)
		}
		_eag := _fcf % 4
		_fcf = 0
		if _eag != 0 {
			_fcf = _eag
		}
	}
	_acgf := _fcf / 2
	if _acgf != 0 {
		for _bac := 0; _bac < _acgf; _bac++ {
			_ggfd = append(_ggfd, 2)
		}
	}
	return _ggfd
}
func (_bfcb *Gray4) SetGray(x, y int, g _d.Gray) {
	if x >= _bfcb.Width || y >= _bfcb.Height {
		return
	}
	g = _gcca(g)
	_bfcb.setGray(x, y, g)
}
func ColorAt(x, y, width, bitsPerColor, colorComponents, bytesPerLine int, data, alpha []byte, decode []float64) (_d.Color, error) {
	switch colorComponents {
	case 1:
		return ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine, data, decode)
	case 3:
		return ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor, data, alpha, decode)
	case 4:
		return ColorAtCMYK(x, y, width, data, decode)
	default:
		return nil, _db.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063o\u006c\u006f\u0072\u0020\u0063\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", colorComponents)
	}
}
func _cac() (_eb [256]uint16) {
	for _gge := 0; _gge < 256; _gge++ {
		if _gge&0x01 != 0 {
			_eb[_gge] |= 0x3
		}
		if _gge&0x02 != 0 {
			_eb[_gge] |= 0xc
		}
		if _gge&0x04 != 0 {
			_eb[_gge] |= 0x30
		}
		if _gge&0x08 != 0 {
			_eb[_gge] |= 0xc0
		}
		if _gge&0x10 != 0 {
			_eb[_gge] |= 0x300
		}
		if _gge&0x20 != 0 {
			_eb[_gge] |= 0xc00
		}
		if _gge&0x40 != 0 {
			_eb[_gge] |= 0x3000
		}
		if _gge&0x80 != 0 {
			_eb[_gge] |= 0xc000
		}
	}
	return _eb
}
func _dbgb(_gabdg _ee.Image, _bgec Image, _cccb _ee.Rectangle) {
	if _acbc, _ggdb := _gabdg.(SMasker); _ggdb && _acbc.HasAlpha() {
		_bgec.(SMasker).MakeAlpha()
	}
	switch _afbe := _gabdg.(type) {
	case Gray:
		_agbfc(_afbe, _bgec.(RGBA), _cccb)
	case NRGBA:
		_ageb(_afbe, _bgec.(RGBA), _cccb)
	case *_ee.NYCbCrA:
		_cccf(_afbe, _bgec.(RGBA), _cccb)
	case CMYK:
		_gdea(_afbe, _bgec.(RGBA), _cccb)
	case RGBA:
		_afdg(_afbe, _bgec.(RGBA), _cccb)
	case nrgba64:
		_bgfe(_afbe, _bgec.(RGBA), _cccb)
	default:
		_fega(_gabdg, _bgec, _cccb)
	}
}
func (_eac *Gray4) Histogram() (_fac [256]int) {
	for _bead := 0; _bead < _eac.Width; _bead++ {
		for _bdga := 0; _bdga < _eac.Height; _bdga++ {
			_fac[_eac.GrayAt(_bead, _bdga).Y]++
		}
	}
	return _fac
}
func (_eeae *Monochrome) Histogram() (_afd [256]int) {
	for _, _aafe := range _eeae.Data {
		_afd[0xff] += int(_bcdf[_eeae.Data[_aafe]])
	}
	return _afd
}

type Gray2 struct{ ImageBase }

var _ Gray = &Gray16{}

func _efef(_fbed _d.Color) _d.Color {
	_eddc := _d.GrayModel.Convert(_fbed).(_d.Gray)
	return _gea(_eddc)
}

var _ Gray = &Gray2{}
var _ _ee.Image = &Gray16{}

func _bgca(_gfgf int, _cafb int) int {
	if _gfgf < _cafb {
		return _gfgf
	}
	return _cafb
}
func (_ebfa *NRGBA32) At(x, y int) _d.Color { _bggdd, _ := _ebfa.ColorAt(x, y); return _bggdd }
func _ageb(_egbe NRGBA, _fdga RGBA, _abgd _ee.Rectangle) {
	for _dcbc := 0; _dcbc < _abgd.Max.X; _dcbc++ {
		for _dgagg := 0; _dgagg < _abgd.Max.Y; _dgagg++ {
			_fbce := _egbe.NRGBAAt(_dcbc, _dgagg)
			_fdga.SetRGBA(_dcbc, _dgagg, _eef(_fbce))
		}
	}
}
func (_dcba *Gray16) ColorModel() _d.Model { return _d.Gray16Model }
func _gfd(_egb, _ffd int, _bbc []byte) *Monochrome {
	_gbe := _bed(_egb, _ffd)
	_gbe.Data = _bbc
	return _gbe
}
func (_ebgf *RGBA32) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtRGBA32(x, y, _ebgf.Width, _ebgf.Data, _ebgf.Alpha, _ebgf.Decode)
}
func _dfe(_aaff, _ced *Monochrome, _cead []byte, _ebdc int) (_dfb error) {
	var (
		_edea, _cdce, _dcf, _aed, _gdfg, _adda, _gdec, _ddb int
		_dfc, _bdf                                          uint32
		_egf, _cdec                                         byte
		_abe                                                uint16
	)
	_eagf := make([]byte, 4)
	_fba := make([]byte, 4)
	for _dcf = 0; _dcf < _aaff.Height-1; _dcf, _aed = _dcf+2, _aed+1 {
		_edea = _dcf * _aaff.BytesPerLine
		_cdce = _aed * _ced.BytesPerLine
		for _gdfg, _adda = 0, 0; _gdfg < _ebdc; _gdfg, _adda = _gdfg+4, _adda+1 {
			for _gdec = 0; _gdec < 4; _gdec++ {
				_ddb = _edea + _gdfg + _gdec
				if _ddb <= len(_aaff.Data)-1 && _ddb < _edea+_aaff.BytesPerLine {
					_eagf[_gdec] = _aaff.Data[_ddb]
				} else {
					_eagf[_gdec] = 0x00
				}
				_ddb = _edea + _aaff.BytesPerLine + _gdfg + _gdec
				if _ddb <= len(_aaff.Data)-1 && _ddb < _edea+(2*_aaff.BytesPerLine) {
					_fba[_gdec] = _aaff.Data[_ddb]
				} else {
					_fba[_gdec] = 0x00
				}
			}
			_dfc = _f.BigEndian.Uint32(_eagf)
			_bdf = _f.BigEndian.Uint32(_fba)
			_bdf &= _dfc
			_bdf &= _bdf << 1
			_bdf &= 0xaaaaaaaa
			_dfc = _bdf | (_bdf << 7)
			_egf = byte(_dfc >> 24)
			_cdec = byte((_dfc >> 8) & 0xff)
			_ddb = _cdce + _adda
			if _ddb+1 == len(_ced.Data)-1 || _ddb+1 >= _cdce+_ced.BytesPerLine {
				_ced.Data[_ddb] = _cead[_egf]
				if _dfb = _ced.setByte(_ddb, _cead[_egf]); _dfb != nil {
					return _db.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _ddb)
				}
			} else {
				_abe = (uint16(_cead[_egf]) << 8) | uint16(_cead[_cdec])
				if _dfb = _ced.setTwoBytes(_ddb, _abe); _dfb != nil {
					return _db.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _ddb)
				}
				_adda++
			}
		}
	}
	return nil
}

var _ Image = &Monochrome{}

type NRGBA interface {
	NRGBAAt(_aedb, _baa int) _d.NRGBA
	SetNRGBA(_cfc, _caad int, _gedf _d.NRGBA)
}

func ColorAtGray2BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_d.Gray, error) {
	_fegc := y*bytesPerLine + x>>2
	if _fegc >= len(data) {
		return _d.Gray{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_dbab := data[_fegc] >> uint(6-(x&3)*2) & 3
	if len(decode) == 2 {
		_dbab = uint8(uint32(LinearInterpolate(float64(_dbab), 0, 3.0, decode[0], decode[1])) & 3)
	}
	return _d.Gray{Y: _dbab * 85}, nil
}
func _ggaec(_dgce *_ee.Gray) bool {
	for _bfaee := 0; _bfaee < len(_dgce.Pix); _bfaee++ {
		if !_egag(_dgce.Pix[_bfaee]) {
			return false
		}
	}
	return true
}

type Histogramer interface{ Histogram() [256]int }

func (_cefce *NRGBA16) At(x, y int) _d.Color { _bcdb, _ := _cefce.ColorAt(x, y); return _bcdb }
func _gbcc(_bfge *Monochrome, _bfca, _dfag, _ceff, _dcfc int, _gbdad RasterOperator, _bfda *Monochrome, _dacae, _fcddd int) error {
	if _bfge == nil {
		return _e.New("\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _gbdad == PixDst {
		return nil
	}
	switch _gbdad {
	case PixClr, PixSet, PixNotDst:
		_bbef(_bfge, _bfca, _dfag, _ceff, _dcfc, _gbdad)
		return nil
	}
	if _bfda == nil {
		_a.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _e.New("\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _fece := _faee(_bfge, _bfca, _dfag, _ceff, _dcfc, _gbdad, _bfda, _dacae, _fcddd); _fece != nil {
		return _fece
	}
	return nil
}
func _gea(_gcf _d.Gray) _d.Gray {
	_daea := _gcf.Y >> 6
	_daea |= _daea << 2
	_gcf.Y = _daea | _daea<<4
	return _gcf
}
func (_daege *Gray4) GrayAt(x, y int) _d.Gray {
	_ceeec, _ := ColorAtGray4BPC(x, y, _daege.BytesPerLine, _daege.Data, _daege.Decode)
	return _ceeec
}
func NewImage(width, height, bitsPerComponent, colorComponents int, data, alpha []byte, decode []float64) (Image, error) {
	_fga := NewImageBase(width, height, bitsPerComponent, colorComponents, data, alpha, decode)
	var _efaa Image
	switch colorComponents {
	case 1:
		switch bitsPerComponent {
		case 1:
			_efaa = &Monochrome{ImageBase: _fga, ModelThreshold: 0x0f}
		case 2:
			_efaa = &Gray2{ImageBase: _fga}
		case 4:
			_efaa = &Gray4{ImageBase: _fga}
		case 8:
			_efaa = &Gray8{ImageBase: _fga}
		case 16:
			_efaa = &Gray16{ImageBase: _fga}
		}
	case 3:
		switch bitsPerComponent {
		case 4:
			_efaa = &NRGBA16{ImageBase: _fga}
		case 8:
			_efaa = &NRGBA32{ImageBase: _fga}
		case 16:
			_efaa = &NRGBA64{ImageBase: _fga}
		}
	case 4:
		_efaa = &CMYK32{ImageBase: _fga}
	}
	if _efaa == nil {
		return nil, ErrInvalidImage
	}
	return _efaa, nil
}
func (_ecgg *Gray4) Validate() error {
	if len(_ecgg.Data) != _ecgg.Height*_ecgg.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func _adee(_fcfe _d.NRGBA) _d.Gray {
	var _bff _d.NRGBA
	if _fcfe == _bff {
		return _d.Gray{Y: 0xff}
	}
	_bfgb, _edbe, _dgg, _ := _fcfe.RGBA()
	_aga := (19595*_bfgb + 38470*_edbe + 7471*_dgg + 1<<15) >> 24
	return _d.Gray{Y: uint8(_aga)}
}
func _eef(_dged _d.NRGBA) _d.RGBA {
	_ceef, _cfge, _gef, _ccc := _dged.RGBA()
	return _d.RGBA{R: uint8(_ceef >> 8), G: uint8(_cfge >> 8), B: uint8(_gef >> 8), A: uint8(_ccc >> 8)}
}
func (_cecb *Monochrome) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtGray1BPC(x, y, _cecb.BytesPerLine, _cecb.Data, _cecb.Decode)
}

type ImageBase struct {
	Width, Height                     int
	BitsPerComponent, ColorComponents int
	Data, Alpha                       []byte
	Decode                            []float64
	BytesPerLine                      int
}

var _ Image = &NRGBA64{}

func (_adgd *ImageBase) copy() ImageBase {
	_eafe := *_adgd
	_eafe.Data = make([]byte, len(_adgd.Data))
	copy(_eafe.Data, _adgd.Data)
	return _eafe
}
func _edgec(_bgd CMYK, _feabg Gray, _beag _ee.Rectangle) {
	for _ebgc := 0; _ebgc < _beag.Max.X; _ebgc++ {
		for _ggdc := 0; _ggdc < _beag.Max.Y; _ggdc++ {
			_aeab := _fafc(_bgd.CMYKAt(_ebgc, _ggdc))
			_feabg.SetGray(_ebgc, _ggdc, _aeab)
		}
	}
}
func _bbbd(_dcged _ee.Image, _ggaee int) (_ee.Rectangle, bool, []byte) {
	_fddd := _dcged.Bounds()
	var (
		_adcg bool
		_bca  []byte
	)
	switch _faaf := _dcged.(type) {
	case SMasker:
		_adcg = _faaf.HasAlpha()
	case NRGBA, RGBA, *_ee.RGBA64, nrgba64, *_ee.NYCbCrA:
		_bca = make([]byte, _fddd.Max.X*_fddd.Max.Y*_ggaee)
	case *_ee.Paletted:
		var _dage bool
		for _, _eecb := range _faaf.Palette {
			_faefbd, _afcb, _gebe, _gfef := _eecb.RGBA()
			if _faefbd == 0 && _afcb == 0 && _gebe == 0 && _gfef != 0 {
				_dage = true
				break
			}
		}
		if _dage {
			_bca = make([]byte, _fddd.Max.X*_fddd.Max.Y*_ggaee)
		}
	}
	return _fddd, _adcg, _bca
}

var (
	_gbdbe = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_afge  = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

const (
	_egfe shift = iota
	_cdfb
)

func (_abgf *Gray16) At(x, y int) _d.Color { _ffe, _ := _abgf.ColorAt(x, y); return _ffe }
func _bea() (_dc [256]uint32) {
	for _cdf := 0; _cdf < 256; _cdf++ {
		if _cdf&0x01 != 0 {
			_dc[_cdf] |= 0xf
		}
		if _cdf&0x02 != 0 {
			_dc[_cdf] |= 0xf0
		}
		if _cdf&0x04 != 0 {
			_dc[_cdf] |= 0xf00
		}
		if _cdf&0x08 != 0 {
			_dc[_cdf] |= 0xf000
		}
		if _cdf&0x10 != 0 {
			_dc[_cdf] |= 0xf0000
		}
		if _cdf&0x20 != 0 {
			_dc[_cdf] |= 0xf00000
		}
		if _cdf&0x40 != 0 {
			_dc[_cdf] |= 0xf000000
		}
		if _cdf&0x80 != 0 {
			_dc[_cdf] |= 0xf0000000
		}
	}
	return _dc
}
func (_fefa *Gray4) Base() *ImageBase { return &_fefa.ImageBase }
func (_acfc *CMYK32) Validate() error {
	if len(_acfc.Data) != 4*_acfc.Width*_acfc.Height {
		return _e.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func (_dffad *ImageBase) Pix() []byte { return _dffad.Data }
func (_bbe *Gray4) setGray(_egeef int, _abf int, _fgbc _d.Gray) {
	_ecf := _abf * _bbe.BytesPerLine
	_gfge := _ecf + (_egeef >> 1)
	if _gfge >= len(_bbe.Data) {
		return
	}
	_fcaa := _fgbc.Y >> 4
	_bbe.Data[_gfge] = (_bbe.Data[_gfge] & (^(0xf0 >> uint(4*(_egeef&1))))) | (_fcaa << uint(4-4*(_egeef&1)))
}
func (_cdfc *Monochrome) AddPadding() (_gga error) {
	if _bdb := ((_cdfc.Width * _cdfc.Height) + 7) >> 3; len(_cdfc.Data) < _bdb {
		return _db.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(_cdfc.Data), _bdb)
	}
	_bdbc := _cdfc.Width % 8
	if _bdbc == 0 {
		return nil
	}
	_gae := _cdfc.Width / 8
	_bdgg := _eeg.NewReader(_cdfc.Data)
	_gbda := make([]byte, _cdfc.Height*_cdfc.BytesPerLine)
	_egec := _eeg.NewWriterMSB(_gbda)
	_cfec := make([]byte, _gae)
	var (
		_gdfe int
		_geg  uint64
	)
	for _gdfe = 0; _gdfe < _cdfc.Height; _gdfe++ {
		if _, _gga = _bdgg.Read(_cfec); _gga != nil {
			return _gga
		}
		if _, _gga = _egec.Write(_cfec); _gga != nil {
			return _gga
		}
		if _geg, _gga = _bdgg.ReadBits(byte(_bdbc)); _gga != nil {
			return _gga
		}
		if _gga = _egec.WriteByte(byte(_geg) << uint(8-_bdbc)); _gga != nil {
			return _gga
		}
	}
	_cdfc.Data = _egec.Data()
	return nil
}
func (_fcbad *Gray16) GrayAt(x, y int) _d.Gray {
	_dcgf, _ := _fcbad.ColorAt(x, y)
	return _d.Gray{Y: uint8(_dcgf.(_d.Gray16).Y >> 8)}
}
func _edg(_be *Monochrome, _ca int, _ab []uint) (*Monochrome, error) {
	_bb := _ca * _be.Width
	_ag := _ca * _be.Height
	_ccbd := _bed(_bb, _ag)
	for _ef, _ac := range _ab {
		var _cd error
		switch _ac {
		case 2:
			_cd = _ccd(_ccbd, _be)
		case 4:
			_cd = _eea(_ccbd, _be)
		case 8:
			_cd = _ec(_ccbd, _be)
		}
		if _cd != nil {
			return nil, _cd
		}
		if _ef != len(_ab)-1 {
			_be = _ccbd.copy()
		}
	}
	return _ccbd, nil
}
func (_ggbe *ImageBase) GetAlpha() []byte { return _ggbe.Alpha }
func (_cbeg *ImageBase) newAlpha() {
	_bdbe := BytesPerLine(_cbeg.Width, _cbeg.BitsPerComponent, 1)
	_cbeg.Alpha = make([]byte, _cbeg.Height*_bdbe)
}
func _ceg(_faefb _ee.Image) (Image, error) {
	if _cdffc, _aeabb := _faefb.(*NRGBA32); _aeabb {
		return _cdffc.Copy(), nil
	}
	_gfgg, _dbeg, _eaaf := _bbbd(_faefb, 1)
	_daab, _cfdca := NewImage(_gfgg.Max.X, _gfgg.Max.Y, 8, 3, nil, _eaaf, nil)
	if _cfdca != nil {
		return nil, _cfdca
	}
	_bgbgd(_faefb, _daab, _gfgg)
	if len(_eaaf) != 0 && !_dbeg {
		if _dbaeb := _efaga(_eaaf, _daab); _dbaeb != nil {
			return nil, _dbaeb
		}
	}
	return _daab, nil
}

var _ _ee.Image = &Gray2{}
var _ NRGBA = &NRGBA32{}

func (_adec *NRGBA32) SetNRGBA(x, y int, c _d.NRGBA) {
	_edebd := y*_adec.Width + x
	_fgcd := 3 * _edebd
	if _fgcd+2 >= len(_adec.Data) {
		return
	}
	_adec.setRGBA(_edebd, c)
}
func (_ddac *Gray16) SetGray(x, y int, g _d.Gray) {
	_aaga := (y*_ddac.BytesPerLine/2 + x) * 2
	if _aaga+1 >= len(_ddac.Data) {
		return
	}
	_ddac.Data[_aaga] = g.Y
	_ddac.Data[_aaga+1] = g.Y
}

var _ _ee.Image = &Gray8{}

func (_gfa *Gray8) Validate() error {
	if len(_gfa.Data) != _gfa.Height*_gfa.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}

type Monochrome struct {
	ImageBase
	ModelThreshold uint8
}

var _ Gray = &Gray4{}

func ColorAtCMYK(x, y, width int, data []byte, decode []float64) (_d.CMYK, error) {
	_cfg := 4 * (y*width + x)
	if _cfg+3 >= len(data) {
		return _d.CMYK{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	C := data[_cfg] & 0xff
	M := data[_cfg+1] & 0xff
	Y := data[_cfg+2] & 0xff
	K := data[_cfg+3] & 0xff
	if len(decode) == 8 {
		C = uint8(uint32(LinearInterpolate(float64(C), 0, 255, decode[0], decode[1])) & 0xff)
		M = uint8(uint32(LinearInterpolate(float64(M), 0, 255, decode[2], decode[3])) & 0xff)
		Y = uint8(uint32(LinearInterpolate(float64(Y), 0, 255, decode[4], decode[5])) & 0xff)
		K = uint8(uint32(LinearInterpolate(float64(K), 0, 255, decode[6], decode[7])) & 0xff)
	}
	return _d.CMYK{C: C, M: M, Y: Y, K: K}, nil
}
func (_beb *CMYK32) Copy() Image { return &CMYK32{ImageBase: _beb.copy()} }
func _abbf(_bfaa _ee.Image) (Image, error) {
	if _daed, _dcca := _bfaa.(*NRGBA16); _dcca {
		return _daed.Copy(), nil
	}
	_ccdff := _bfaa.Bounds()
	_gecb, _cbed := NewImage(_ccdff.Max.X, _ccdff.Max.Y, 4, 3, nil, nil, nil)
	if _cbed != nil {
		return nil, _cbed
	}
	_bgbgd(_bfaa, _gecb, _ccdff)
	return _gecb, nil
}
func _bgdb(_eabd *Monochrome, _dabg, _egefe int, _abeg, _bgee int, _adae RasterOperator) {
	var (
		_cbge   bool
		_aafa   bool
		_fcdddf int
		_ffeb   int
		_baeg   int
		_ffad   int
		_bfacc  bool
		_babg   byte
	)
	_defa := 8 - (_dabg & 7)
	_afdc := _afge[_defa]
	_ababc := _eabd.BytesPerLine*_egefe + (_dabg >> 3)
	if _abeg < _defa {
		_cbge = true
		_afdc &= _gbdbe[8-_defa+_abeg]
	}
	if !_cbge {
		_fcdddf = (_abeg - _defa) >> 3
		if _fcdddf != 0 {
			_aafa = true
			_ffeb = _ababc + 1
		}
	}
	_baeg = (_dabg + _abeg) & 7
	if !(_cbge || _baeg == 0) {
		_bfacc = true
		_babg = _gbdbe[_baeg]
		_ffad = _ababc + 1 + _fcdddf
	}
	var _efgda, _dgde int
	switch _adae {
	case PixClr:
		for _efgda = 0; _efgda < _bgee; _efgda++ {
			_eabd.Data[_ababc] = _ggdf(_eabd.Data[_ababc], 0x0, _afdc)
			_ababc += _eabd.BytesPerLine
		}
		if _aafa {
			for _efgda = 0; _efgda < _bgee; _efgda++ {
				for _dgde = 0; _dgde < _fcdddf; _dgde++ {
					_eabd.Data[_ffeb+_dgde] = 0x0
				}
				_ffeb += _eabd.BytesPerLine
			}
		}
		if _bfacc {
			for _efgda = 0; _efgda < _bgee; _efgda++ {
				_eabd.Data[_ffad] = _ggdf(_eabd.Data[_ffad], 0x0, _babg)
				_ffad += _eabd.BytesPerLine
			}
		}
	case PixSet:
		for _efgda = 0; _efgda < _bgee; _efgda++ {
			_eabd.Data[_ababc] = _ggdf(_eabd.Data[_ababc], 0xff, _afdc)
			_ababc += _eabd.BytesPerLine
		}
		if _aafa {
			for _efgda = 0; _efgda < _bgee; _efgda++ {
				for _dgde = 0; _dgde < _fcdddf; _dgde++ {
					_eabd.Data[_ffeb+_dgde] = 0xff
				}
				_ffeb += _eabd.BytesPerLine
			}
		}
		if _bfacc {
			for _efgda = 0; _efgda < _bgee; _efgda++ {
				_eabd.Data[_ffad] = _ggdf(_eabd.Data[_ffad], 0xff, _babg)
				_ffad += _eabd.BytesPerLine
			}
		}
	case PixNotDst:
		for _efgda = 0; _efgda < _bgee; _efgda++ {
			_eabd.Data[_ababc] = _ggdf(_eabd.Data[_ababc], ^_eabd.Data[_ababc], _afdc)
			_ababc += _eabd.BytesPerLine
		}
		if _aafa {
			for _efgda = 0; _efgda < _bgee; _efgda++ {
				for _dgde = 0; _dgde < _fcdddf; _dgde++ {
					_eabd.Data[_ffeb+_dgde] = ^(_eabd.Data[_ffeb+_dgde])
				}
				_ffeb += _eabd.BytesPerLine
			}
		}
		if _bfacc {
			for _efgda = 0; _efgda < _bgee; _efgda++ {
				_eabd.Data[_ffad] = _ggdf(_eabd.Data[_ffad], ^_eabd.Data[_ffad], _babg)
				_ffad += _eabd.BytesPerLine
			}
		}
	}
}

type nrgba64 interface {
	NRGBA64At(_fegd, _becfb int) _d.NRGBA64
	SetNRGBA64(_afdae, _bgbd int, _ffdg _d.NRGBA64)
}

func (_dcfg *Gray2) Set(x, y int, c _d.Color) {
	if x >= _dcfg.Width || y >= _dcfg.Height {
		return
	}
	_dffe := Gray2Model.Convert(c).(_d.Gray)
	_eaf := y * _dcfg.BytesPerLine
	_abb := _eaf + (x >> 2)
	_bagg := _dffe.Y >> 6
	_dcfg.Data[_abb] = (_dcfg.Data[_abb] & (^(0xc0 >> uint(2*((x)&3))))) | (_bagg << uint(6-2*(x&3)))
}
func _becab(_fbcf nrgba64, _ffadd NRGBA, _cga _ee.Rectangle) {
	for _cbca := 0; _cbca < _cga.Max.X; _cbca++ {
		for _fefg := 0; _fefg < _cga.Max.Y; _fefg++ {
			_cbcb := _fbcf.NRGBA64At(_cbca, _fefg)
			_ffadd.SetNRGBA(_cbca, _fefg, _cdcc(_cbcb))
		}
	}
}
func ScaleAlphaToMonochrome(data []byte, width, height int) ([]byte, error) {
	_gb := BytesPerLine(width, 8, 1)
	if len(data) < _gb*height {
		return nil, nil
	}
	_cc := &Gray8{NewImageBase(width, height, 8, 1, data, nil, nil)}
	_ccb, _dg := MonochromeConverter.Convert(_cc)
	if _dg != nil {
		return nil, _dg
	}
	return _ccb.Base().Data, nil
}
func (_dac *Monochrome) Scale(scale float64) (*Monochrome, error) {
	var _cgg bool
	_cgedg := scale
	if scale < 1 {
		_cgedg = 1 / scale
		_cgg = true
	}
	_geda := NextPowerOf2(uint(_cgedg))
	if InDelta(float64(_geda), _cgedg, 0.001) {
		if _cgg {
			return _dac.ReduceBinary(_cgedg)
		}
		return _dac.ExpandBinary(int(_geda))
	}
	_eebd := int(_g.RoundToEven(float64(_dac.Width) * scale))
	_dag := int(_g.RoundToEven(float64(_dac.Height) * scale))
	return _dac.ScaleLow(_eebd, _dag)
}
func _faee(_bbdd *Monochrome, _befb, _aafea int, _egaf, _efdf int, _ccea RasterOperator, _cfbc *Monochrome, _bfad, _dde int) error {
	var _cagd, _edbd, _aggd, _cgebb int
	if _befb < 0 {
		_bfad -= _befb
		_egaf += _befb
		_befb = 0
	}
	if _bfad < 0 {
		_befb -= _bfad
		_egaf += _bfad
		_bfad = 0
	}
	_cagd = _befb + _egaf - _bbdd.Width
	if _cagd > 0 {
		_egaf -= _cagd
	}
	_edbd = _bfad + _egaf - _cfbc.Width
	if _edbd > 0 {
		_egaf -= _edbd
	}
	if _aafea < 0 {
		_dde -= _aafea
		_efdf += _aafea
		_aafea = 0
	}
	if _dde < 0 {
		_aafea -= _dde
		_efdf += _dde
		_dde = 0
	}
	_aggd = _aafea + _efdf - _bbdd.Height
	if _aggd > 0 {
		_efdf -= _aggd
	}
	_cgebb = _dde + _efdf - _cfbc.Height
	if _cgebb > 0 {
		_efdf -= _cgebb
	}
	if _egaf <= 0 || _efdf <= 0 {
		return nil
	}
	var _gcec error
	switch {
	case _befb&7 == 0 && _bfad&7 == 0:
		_gcec = _gafe(_bbdd, _befb, _aafea, _egaf, _efdf, _ccea, _cfbc, _bfad, _dde)
	case _befb&7 == _bfad&7:
		_gcec = _geff(_bbdd, _befb, _aafea, _egaf, _efdf, _ccea, _cfbc, _bfad, _dde)
	default:
		_gcec = _cadc(_bbdd, _befb, _aafea, _egaf, _efdf, _ccea, _cfbc, _bfad, _dde)
	}
	if _gcec != nil {
		return _gcec
	}
	return nil
}
func NewImageBase(width int, height int, bitsPerComponent int, colorComponents int, data []byte, alpha []byte, decode []float64) ImageBase {
	_edfa := ImageBase{Width: width, Height: height, BitsPerComponent: bitsPerComponent, ColorComponents: colorComponents, Data: data, Alpha: alpha, Decode: decode, BytesPerLine: BytesPerLine(width, bitsPerComponent, colorComponents)}
	if data == nil {
		_edfa.Data = make([]byte, height*_edfa.BytesPerLine)
	}
	return _edfa
}

type SMasker interface {
	HasAlpha() bool
	GetAlpha() []byte
	MakeAlpha()
}

func (_dgbb *Gray8) Copy() Image { return &Gray8{ImageBase: _dgbb.copy()} }
func ColorAtNRGBA(x, y, width, bytesPerLine, bitsPerColor int, data, alpha []byte, decode []float64) (_d.Color, error) {
	switch bitsPerColor {
	case 4:
		return ColorAtNRGBA16(x, y, width, bytesPerLine, data, alpha, decode)
	case 8:
		return ColorAtNRGBA32(x, y, width, data, alpha, decode)
	case 16:
		return ColorAtNRGBA64(x, y, width, data, alpha, decode)
	default:
		return nil, _db.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0067\u0062\u0020b\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0061\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
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
	return nil, _db.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0043o\u006e\u0076\u0065\u0072\u0074\u0065\u0072\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072\u0073\u002e\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003a\u0020\u0025\u0064\u002c\u0020\u0043\u006f\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u003a \u0025\u0064", bitsPerComponent, colorComponents)
}
func _beddc(_aabc _ee.Image, _cfgfd uint8) *_ee.Gray {
	_fdgd := _aabc.Bounds()
	_bgeca := _ee.NewGray(_fdgd)
	var (
		_faafc _d.Color
		_ceaa  _d.Gray
	)
	for _bacg := 0; _bacg < _fdgd.Max.X; _bacg++ {
		for _bbagf := 0; _bbagf < _fdgd.Max.Y; _bbagf++ {
			_faafc = _aabc.At(_bacg, _bbagf)
			_bgeca.Set(_bacg, _bbagf, _faafc)
			_ceaa = _bgeca.GrayAt(_bacg, _bbagf)
			_bgeca.SetGray(_bacg, _bbagf, _d.Gray{Y: _cdbe(_ceaa.Y, _cfgfd)})
		}
	}
	return _bgeca
}
func GrayHistogram(g Gray) (_gggbc [256]int) {
	switch _edfg := g.(type) {
	case Histogramer:
		return _edfg.Histogram()
	case _ee.Image:
		_ddbg := _edfg.Bounds()
		for _fegf := 0; _fegf < _ddbg.Max.X; _fegf++ {
			for _bgbda := 0; _bgbda < _ddbg.Max.Y; _bgbda++ {
				_gggbc[g.GrayAt(_fegf, _bgbda).Y]++
			}
		}
		return _gggbc
	default:
		return [256]int{}
	}
}
func (_ged colorConverter) Convert(src _ee.Image) (Image, error) { return _ged._dfeb(src) }
func (_ddd *ImageBase) getByte(_eaag int) (byte, error) {
	if _eaag > len(_ddd.Data)-1 || _eaag < 0 {
		return 0, _db.Errorf("\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", _eaag)
	}
	return _ddd.Data[_eaag], nil
}
func _gegf(_fgfd *_ee.Gray16, _bcaf uint8) *_ee.Gray {
	_ccdb := _fgfd.Bounds()
	_cbfd := _ee.NewGray(_ccdb)
	for _gecfc := 0; _gecfc < _ccdb.Dx(); _gecfc++ {
		for _gfeb := 0; _gfeb < _ccdb.Dy(); _gfeb++ {
			_ffbeg := _fgfd.Gray16At(_gecfc, _gfeb)
			_cbfd.SetGray(_gecfc, _gfeb, _d.Gray{Y: _cdbe(uint8(_ffbeg.Y/256), _bcaf)})
		}
	}
	return _cbfd
}
func ColorAtGrayscale(x, y, bitsPerColor, bytesPerLine int, data []byte, decode []float64) (_d.Color, error) {
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
		return nil, _db.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0067\u0072\u0061\u0079\u0020\u0073c\u0061\u006c\u0065\u0020\u0062\u0069\u0074s\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020a\u006d\u006f\u0075\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", bitsPerColor)
	}
}
func (_ccce *NRGBA16) ColorModel() _d.Model { return NRGBA16Model }

type CMYK32 struct{ ImageBase }

func (_fcga *RGBA32) At(x, y int) _d.Color { _aaba, _ := _fcga.ColorAt(x, y); return _aaba }

var _ _ee.Image = &RGBA32{}

func ColorAtGray1BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_d.Gray, error) {
	_aca := y*bytesPerLine + x>>3
	if _aca >= len(data) {
		return _d.Gray{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_feb := data[_aca] >> uint(7-(x&7)) & 1
	if len(decode) == 2 {
		_feb = uint8(LinearInterpolate(float64(_feb), 0.0, 1.0, decode[0], decode[1])) & 1
	}
	return _d.Gray{Y: _feb * 255}, nil
}
func (_gcba *NRGBA64) Copy() Image { return &NRGBA64{ImageBase: _gcba.copy()} }
func _ccd(_fg, _ae *Monochrome) (_fbf error) {
	_ba := _ae.BytesPerLine
	_edc := _fg.BytesPerLine
	var (
		_dd                      byte
		_gf                      uint16
		_ge, _ace, _ga, _ea, _aa int
	)
	for _ga = 0; _ga < _ae.Height; _ga++ {
		_ge = _ga * _ba
		_ace = 2 * _ga * _edc
		for _ea = 0; _ea < _ba; _ea++ {
			_dd = _ae.Data[_ge+_ea]
			_gf = _cfb[_dd]
			_aa = _ace + _ea*2
			if _fg.BytesPerLine != _ae.BytesPerLine*2 && (_ea+1)*2 > _fg.BytesPerLine {
				_fbf = _fg.setByte(_aa, byte(_gf>>8))
			} else {
				_fbf = _fg.setTwoBytes(_aa, _gf)
			}
			if _fbf != nil {
				return _fbf
			}
		}
		for _ea = 0; _ea < _edc; _ea++ {
			_aa = _ace + _edc + _ea
			_dd = _fg.Data[_ace+_ea]
			if _fbf = _fg.setByte(_aa, _dd); _fbf != nil {
				return _fbf
			}
		}
	}
	return nil
}
func _ggdf(_fbbb, _bagfg, _dgge byte) byte { return (_fbbb &^ (_dgge)) | (_bagfg & _dgge) }
func ColorAtNRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_d.NRGBA, error) {
	_ebce := y*width + x
	_bdefe := 3 * _ebce
	if _bdefe+2 >= len(data) {
		return _d.NRGBA{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_fadc := uint8(0xff)
	if alpha != nil && len(alpha) > _ebce {
		_fadc = alpha[_ebce]
	}
	_aafg, _eabf, _ggaad := data[_bdefe], data[_bdefe+1], data[_bdefe+2]
	if len(decode) == 6 {
		_aafg = uint8(uint32(LinearInterpolate(float64(_aafg), 0, 255, decode[0], decode[1])) & 0xff)
		_eabf = uint8(uint32(LinearInterpolate(float64(_eabf), 0, 255, decode[2], decode[3])) & 0xff)
		_ggaad = uint8(uint32(LinearInterpolate(float64(_ggaad), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _d.NRGBA{R: _aafg, G: _eabf, B: _ggaad, A: _fadc}, nil
}

var _ _ee.Image = &NRGBA32{}

func (_bbgcf *NRGBA64) Base() *ImageBase { return &_bbgcf.ImageBase }
func _gbgd(_cbae _ee.Image, _bfae Image, _cbe _ee.Rectangle) {
	switch _aad := _cbae.(type) {
	case Gray:
		_babb(_aad, _bfae.(Gray), _cbe)
	case NRGBA:
		_ddcc(_aad, _bfae.(Gray), _cbe)
	case CMYK:
		_edgec(_aad, _bfae.(Gray), _cbe)
	case RGBA:
		_gccd(_aad, _bfae.(Gray), _cbe)
	default:
		_fega(_cbae, _bfae, _cbe)
	}
}
func _fgbcb(_abfe *_ee.NYCbCrA, _beee NRGBA, _bdc _ee.Rectangle) {
	for _fagg := 0; _fagg < _bdc.Max.X; _fagg++ {
		for _fcef := 0; _fcef < _bdc.Max.Y; _fcef++ {
			_dgca := _abfe.NYCbCrAAt(_fagg, _fcef)
			_beee.SetNRGBA(_fagg, _fcef, _egee(_dgca))
		}
	}
}
func _fdf(_bee Gray, _ggef CMYK, _fef _ee.Rectangle) {
	for _beca := 0; _beca < _fef.Max.X; _beca++ {
		for _effe := 0; _effe < _fef.Max.Y; _effe++ {
			_edeaa := _bee.GrayAt(_beca, _effe)
			_ggef.SetCMYK(_beca, _effe, _aafb(_edeaa))
		}
	}
}

var _ Gray = &Monochrome{}

func (_aedf *RGBA32) setRGBA(_faggg int, _aabd _d.RGBA) {
	_gdaa := 3 * _faggg
	_aedf.Data[_gdaa] = _aabd.R
	_aedf.Data[_gdaa+1] = _aabd.G
	_aedf.Data[_gdaa+2] = _aabd.B
	if _faggg < len(_aedf.Alpha) {
		_aedf.Alpha[_faggg] = _aabd.A
	}
}
func (_gccb *Gray16) Set(x, y int, c _d.Color) {
	_gdbe := (y*_gccb.BytesPerLine/2 + x) * 2
	if _gdbe+1 >= len(_gccb.Data) {
		return
	}
	_egbd := _d.Gray16Model.Convert(c).(_d.Gray16)
	_gccb.Data[_gdbe], _gccb.Data[_gdbe+1] = uint8(_egbd.Y>>8), uint8(_egbd.Y&0xff)
}
func (_egea *Gray8) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtGray8BPC(x, y, _egea.BytesPerLine, _egea.Data, _egea.Decode)
}
func (_dcae *Monochrome) clearBit(_dcb, _fbgd int) { _dcae.Data[_dcb] &= ^(0x80 >> uint(_fbgd&7)) }
func (_cgdaa *NRGBA16) Base() *ImageBase           { return &_cgdaa.ImageBase }
func (_gcda *Gray2) Base() *ImageBase              { return &_gcda.ImageBase }
func ImgToBinary(i _ee.Image, threshold uint8) *_ee.Gray {
	switch _dccc := i.(type) {
	case *_ee.Gray:
		if _ggaec(_dccc) {
			return _dccc
		}
		return _ccaf(_dccc, threshold)
	case *_ee.Gray16:
		return _gegf(_dccc, threshold)
	default:
		return _beddc(_dccc, threshold)
	}
}
func _gafe(_agdf *Monochrome, _gfag, _eccf, _febf, _fcac int, _fee RasterOperator, _fgcc *Monochrome, _baeb, _agceg int) error {
	var (
		_aeaa        byte
		_bbaeb       int
		_gbfg        int
		_ecbg, _eeaa int
		_fbfd, _bcbg int
	)
	_ecfc := _febf >> 3
	_dgd := _febf & 7
	if _dgd > 0 {
		_aeaa = _gbdbe[_dgd]
	}
	_bbaeb = _fgcc.BytesPerLine*_agceg + (_baeb >> 3)
	_gbfg = _agdf.BytesPerLine*_eccf + (_gfag >> 3)
	switch _fee {
	case PixSrc:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] = _fgcc.Data[_ecbg]
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], _fgcc.Data[_ecbg], _aeaa)
			}
		}
	case PixNotSrc:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] = ^(_fgcc.Data[_ecbg])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], ^_fgcc.Data[_ecbg], _aeaa)
			}
		}
	case PixSrcOrDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] |= _fgcc.Data[_ecbg]
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], _fgcc.Data[_ecbg]|_agdf.Data[_eeaa], _aeaa)
			}
		}
	case PixSrcAndDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] &= _fgcc.Data[_ecbg]
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], _fgcc.Data[_ecbg]&_agdf.Data[_eeaa], _aeaa)
			}
		}
	case PixSrcXorDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] ^= _fgcc.Data[_ecbg]
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], _fgcc.Data[_ecbg]^_agdf.Data[_eeaa], _aeaa)
			}
		}
	case PixNotSrcOrDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] |= ^(_fgcc.Data[_ecbg])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], ^(_fgcc.Data[_ecbg])|_agdf.Data[_eeaa], _aeaa)
			}
		}
	case PixNotSrcAndDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] &= ^(_fgcc.Data[_ecbg])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], ^(_fgcc.Data[_ecbg])&_agdf.Data[_eeaa], _aeaa)
			}
		}
	case PixSrcOrNotDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] = _fgcc.Data[_ecbg] | ^(_agdf.Data[_eeaa])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], _fgcc.Data[_ecbg]|^(_agdf.Data[_eeaa]), _aeaa)
			}
		}
	case PixSrcAndNotDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] = _fgcc.Data[_ecbg] &^ (_agdf.Data[_eeaa])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], _fgcc.Data[_ecbg]&^(_agdf.Data[_eeaa]), _aeaa)
			}
		}
	case PixNotPixSrcOrDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] = ^(_fgcc.Data[_ecbg] | _agdf.Data[_eeaa])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], ^(_fgcc.Data[_ecbg] | _agdf.Data[_eeaa]), _aeaa)
			}
		}
	case PixNotPixSrcAndDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] = ^(_fgcc.Data[_ecbg] & _agdf.Data[_eeaa])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], ^(_fgcc.Data[_ecbg] & _agdf.Data[_eeaa]), _aeaa)
			}
		}
	case PixNotPixSrcXorDst:
		for _fbfd = 0; _fbfd < _fcac; _fbfd++ {
			_ecbg = _bbaeb + _fbfd*_fgcc.BytesPerLine
			_eeaa = _gbfg + _fbfd*_agdf.BytesPerLine
			for _bcbg = 0; _bcbg < _ecfc; _bcbg++ {
				_agdf.Data[_eeaa] = ^(_fgcc.Data[_ecbg] ^ _agdf.Data[_eeaa])
				_eeaa++
				_ecbg++
			}
			if _dgd > 0 {
				_agdf.Data[_eeaa] = _ggdf(_agdf.Data[_eeaa], ^(_fgcc.Data[_ecbg] ^ _agdf.Data[_eeaa]), _aeaa)
			}
		}
	default:
		_a.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _fee)
		return _e.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_bcee *Gray8) Set(x, y int, c _d.Color) {
	_ecbe := y*_bcee.BytesPerLine + x
	if _ecbe > len(_bcee.Data)-1 {
		return
	}
	_afda := _d.GrayModel.Convert(c)
	_bcee.Data[_ecbe] = _afda.(_d.Gray).Y
}
func InDelta(expected, current, delta float64) bool {
	_dggbe := expected - current
	if _dggbe <= -delta || _dggbe >= delta {
		return false
	}
	return true
}
func _ccaf(_gdcd *_ee.Gray, _bbefd uint8) *_ee.Gray {
	_gccg := _gdcd.Bounds()
	_ecce := _ee.NewGray(_gccg)
	for _eeca := 0; _eeca < _gccg.Dx(); _eeca++ {
		for _ecggb := 0; _ecggb < _gccg.Dy(); _ecggb++ {
			_adeae := _gdcd.GrayAt(_eeca, _ecggb)
			_ecce.SetGray(_eeca, _ecggb, _d.Gray{Y: _cdbe(_adeae.Y, _bbefd)})
		}
	}
	return _ecce
}
func (_bfe *Gray2) At(x, y int) _d.Color { _dfdf, _ := _bfe.ColorAt(x, y); return _dfdf }
func _aafb(_fcdde _d.Gray) _d.CMYK       { return _d.CMYK{K: 0xff - _fcdde.Y} }
func _gc(_acg *Monochrome, _fea, _eff int) (*Monochrome, error) {
	if _acg == nil {
		return nil, _e.New("\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _fea <= 0 || _eff <= 0 {
		return nil, _e.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _fea == _eff {
		if _fea == 1 {
			return _acg.copy(), nil
		}
		if _fea == 2 || _fea == 4 || _fea == 8 {
			_cdgc, _bf := _b(_acg, _fea)
			if _bf != nil {
				return nil, _bf
			}
			return _cdgc, nil
		}
	}
	_bbag := _fea * _acg.Width
	_edb := _eff * _acg.Height
	_bdg := _bed(_bbag, _edb)
	_af := _bdg.BytesPerLine
	var (
		_gdf, _ad, _dga, _add, _ddg int
		_ccdd                       byte
		_deb                        error
	)
	for _ad = 0; _ad < _acg.Height; _ad++ {
		_gdf = _eff * _ad * _af
		for _dga = 0; _dga < _acg.Width; _dga++ {
			if _eee := _acg.getBitAt(_dga, _ad); _eee {
				_ddg = _fea * _dga
				for _add = 0; _add < _fea; _add++ {
					_bdg.setIndexedBit(_gdf*8 + _ddg + _add)
				}
			}
		}
		for _add = 1; _add < _eff; _add++ {
			_ead := _gdf + _add*_af
			for _dbd := 0; _dbd < _af; _dbd++ {
				if _ccdd, _deb = _bdg.getByte(_gdf + _dbd); _deb != nil {
					return nil, _deb
				}
				if _deb = _bdg.setByte(_ead+_dbd, _ccdd); _deb != nil {
					return nil, _deb
				}
			}
		}
	}
	return _bdg, nil
}

var (
	_cfb = _cac()
	_acf = _bea()
	_bg  = _efe()
)

type NRGBA32 struct{ ImageBase }

func MonochromeModel(threshold uint8) _d.Model { return monochromeModel(threshold) }
func (_ffa *Gray16) Histogram() (_bbce [256]int) {
	for _dafcd := 0; _dafcd < _ffa.Width; _dafcd++ {
		for _geee := 0; _geee < _ffa.Height; _geee++ {
			_bbce[_ffa.GrayAt(_dafcd, _geee).Y]++
		}
	}
	return _bbce
}
func _aced(_fafcb _d.RGBA) _d.Gray {
	_adg := (19595*uint32(_fafcb.R) + 38470*uint32(_fafcb.G) + 7471*uint32(_fafcb.B) + 1<<7) >> 16
	return _d.Gray{Y: uint8(_adg)}
}
func (_cfba *Monochrome) SetGray(x, y int, g _d.Gray) {
	_fddg := y*_cfba.BytesPerLine + x>>3
	if _fddg > len(_cfba.Data)-1 {
		return
	}
	g = _beg(g, monochromeModel(_cfba.ModelThreshold))
	_cfba.setGray(x, g, _fddg)
}
func (_dcde *Monochrome) InverseData() error {
	return _dcde.RasterOperation(0, 0, _dcde.Width, _dcde.Height, PixNotDst, nil, 0, 0)
}
func ConverterFunc(converterFunc func(_ceae _ee.Image) (Image, error)) ColorConverter {
	return colorConverter{_dfeb: converterFunc}
}
func _cge(_gdb, _ffc *Monochrome, _cea []byte, _gad int) (_cacf error) {
	var (
		_cged, _abg, _cacg, _ede, _addc, _acfb, _agcc, _cde int
		_bde, _fgc                                          uint32
		_gbc, _bdeg                                         byte
		_fca                                                uint16
	)
	_cacfd := make([]byte, 4)
	_cagc := make([]byte, 4)
	for _cacg = 0; _cacg < _gdb.Height-1; _cacg, _ede = _cacg+2, _ede+1 {
		_cged = _cacg * _gdb.BytesPerLine
		_abg = _ede * _ffc.BytesPerLine
		for _addc, _acfb = 0, 0; _addc < _gad; _addc, _acfb = _addc+4, _acfb+1 {
			for _agcc = 0; _agcc < 4; _agcc++ {
				_cde = _cged + _addc + _agcc
				if _cde <= len(_gdb.Data)-1 && _cde < _cged+_gdb.BytesPerLine {
					_cacfd[_agcc] = _gdb.Data[_cde]
				} else {
					_cacfd[_agcc] = 0x00
				}
				_cde = _cged + _gdb.BytesPerLine + _addc + _agcc
				if _cde <= len(_gdb.Data)-1 && _cde < _cged+(2*_gdb.BytesPerLine) {
					_cagc[_agcc] = _gdb.Data[_cde]
				} else {
					_cagc[_agcc] = 0x00
				}
			}
			_bde = _f.BigEndian.Uint32(_cacfd)
			_fgc = _f.BigEndian.Uint32(_cagc)
			_fgc |= _bde
			_fgc |= _fgc << 1
			_fgc &= 0xaaaaaaaa
			_bde = _fgc | (_fgc << 7)
			_gbc = byte(_bde >> 24)
			_bdeg = byte((_bde >> 8) & 0xff)
			_cde = _abg + _acfb
			if _cde+1 == len(_ffc.Data)-1 || _cde+1 >= _abg+_ffc.BytesPerLine {
				_ffc.Data[_cde] = _cea[_gbc]
			} else {
				_fca = (uint16(_cea[_gbc]) << 8) | uint16(_cea[_bdeg])
				if _cacf = _ffc.setTwoBytes(_cde, _fca); _cacf != nil {
					return _db.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _cde)
				}
				_acfb++
			}
		}
	}
	return nil
}
func _beg(_dgfd _d.Gray, _fdeb monochromeModel) _d.Gray {
	if _dgfd.Y > uint8(_fdeb) {
		return _d.Gray{Y: _g.MaxUint8}
	}
	return _d.Gray{}
}
func (_edef *RGBA32) Copy() Image { return &RGBA32{ImageBase: _edef.copy()} }
func (_bceg *CMYK32) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _bceg.Width, Y: _bceg.Height}}
}
func (_gfdc *NRGBA64) ColorModel() _d.Model { return _d.NRGBA64Model }
func (_fbecd *NRGBA64) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _fbecd.Width, Y: _fbecd.Height}}
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

type RasterOperator int

func (_ffab *ImageBase) setTwoBytes(_facd int, _gbfc uint16) error {
	if _facd+1 > len(_ffab.Data)-1 {
		return _e.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_ffab.Data[_facd] = byte((_gbfc & 0xff00) >> 8)
	_ffab.Data[_facd+1] = byte(_gbfc & 0xff)
	return nil
}
func ColorAtNRGBA16(x, y, width, bytesPerLine int, data, alpha []byte, decode []float64) (_d.NRGBA, error) {
	_beed := y*bytesPerLine + x*3/2
	if _beed+1 >= len(data) {
		return _d.NRGBA{}, _eeac(x, y)
	}
	const (
		_ecae = 0xf
		_eegg = uint8(0xff)
	)
	_baag := _eegg
	if alpha != nil {
		_baebb := y * BytesPerLine(width, 4, 1)
		if _baebb < len(alpha) {
			if x%2 == 0 {
				_baag = (alpha[_baebb] >> uint(4)) & _ecae
			} else {
				_baag = alpha[_baebb] & _ecae
			}
			_baag |= _baag << 4
		}
	}
	var _ddcca, _ecdb, _abc uint8
	if x*3%2 == 0 {
		_ddcca = (data[_beed] >> uint(4)) & _ecae
		_ecdb = data[_beed] & _ecae
		_abc = (data[_beed+1] >> uint(4)) & _ecae
	} else {
		_ddcca = data[_beed] & _ecae
		_ecdb = (data[_beed+1] >> uint(4)) & _ecae
		_abc = data[_beed+1] & _ecae
	}
	if len(decode) == 6 {
		_ddcca = uint8(uint32(LinearInterpolate(float64(_ddcca), 0, 15, decode[0], decode[1])) & 0xf)
		_ecdb = uint8(uint32(LinearInterpolate(float64(_ecdb), 0, 15, decode[2], decode[3])) & 0xf)
		_abc = uint8(uint32(LinearInterpolate(float64(_abc), 0, 15, decode[4], decode[5])) & 0xf)
	}
	return _d.NRGBA{R: (_ddcca << 4) | (_ddcca & 0xf), G: (_ecdb << 4) | (_ecdb & 0xf), B: (_abc << 4) | (_abc & 0xf), A: _baag}, nil
}
func _geff(_cgca *Monochrome, _cgga, _gdce, _ceed, _eagba int, _cbcc RasterOperator, _bggf *Monochrome, _efae, _ebde int) error {
	var (
		_eefc         bool
		_aabb         bool
		_acddb        int
		_fgfcd        int
		_geca         int
		_fcgf         bool
		_egeg         byte
		_gbcd         int
		_bgcgf        int
		_gcfe         int
		_effa, _fcacf int
	)
	_ebbb := 8 - (_cgga & 7)
	_egef := _afge[_ebbb]
	_dbce := _cgca.BytesPerLine*_gdce + (_cgga >> 3)
	_faadg := _bggf.BytesPerLine*_ebde + (_efae >> 3)
	if _ceed < _ebbb {
		_eefc = true
		_egef &= _gbdbe[8-_ebbb+_ceed]
	}
	if !_eefc {
		_acddb = (_ceed - _ebbb) >> 3
		if _acddb > 0 {
			_aabb = true
			_fgfcd = _dbce + 1
			_geca = _faadg + 1
		}
	}
	_gbcd = (_cgga + _ceed) & 7
	if !(_eefc || _gbcd == 0) {
		_fcgf = true
		_egeg = _gbdbe[_gbcd]
		_bgcgf = _dbce + 1 + _acddb
		_gcfe = _faadg + 1 + _acddb
	}
	switch _cbcc {
	case PixSrc:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], _bggf.Data[_faadg], _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] = _bggf.Data[_geca+_fcacf]
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], _bggf.Data[_gcfe], _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixNotSrc:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], ^_bggf.Data[_faadg], _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] = ^_bggf.Data[_geca+_fcacf]
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], ^_bggf.Data[_gcfe], _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixSrcOrDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], _bggf.Data[_faadg]|_cgca.Data[_dbce], _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] |= _bggf.Data[_geca+_fcacf]
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], _bggf.Data[_gcfe]|_cgca.Data[_bgcgf], _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixSrcAndDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], _bggf.Data[_faadg]&_cgca.Data[_dbce], _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] &= _bggf.Data[_geca+_fcacf]
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], _bggf.Data[_gcfe]&_cgca.Data[_bgcgf], _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixSrcXorDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], _bggf.Data[_faadg]^_cgca.Data[_dbce], _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] ^= _bggf.Data[_geca+_fcacf]
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], _bggf.Data[_gcfe]^_cgca.Data[_bgcgf], _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixNotSrcOrDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], ^(_bggf.Data[_faadg])|_cgca.Data[_dbce], _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] |= ^(_bggf.Data[_geca+_fcacf])
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], ^(_bggf.Data[_gcfe])|_cgca.Data[_bgcgf], _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixNotSrcAndDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], ^(_bggf.Data[_faadg])&_cgca.Data[_dbce], _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] &= ^_bggf.Data[_geca+_fcacf]
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], ^(_bggf.Data[_gcfe])&_cgca.Data[_bgcgf], _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixSrcOrNotDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], _bggf.Data[_faadg]|^(_cgca.Data[_dbce]), _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] = _bggf.Data[_geca+_fcacf] | ^(_cgca.Data[_fgfcd+_fcacf])
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], _bggf.Data[_gcfe]|^(_cgca.Data[_bgcgf]), _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixSrcAndNotDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], _bggf.Data[_faadg]&^(_cgca.Data[_dbce]), _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] = _bggf.Data[_geca+_fcacf] &^ (_cgca.Data[_fgfcd+_fcacf])
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], _bggf.Data[_gcfe]&^(_cgca.Data[_bgcgf]), _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixNotPixSrcOrDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], ^(_bggf.Data[_faadg] | _cgca.Data[_dbce]), _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] = ^(_bggf.Data[_geca+_fcacf] | _cgca.Data[_fgfcd+_fcacf])
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], ^(_bggf.Data[_gcfe] | _cgca.Data[_bgcgf]), _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixNotPixSrcAndDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], ^(_bggf.Data[_faadg] & _cgca.Data[_dbce]), _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] = ^(_bggf.Data[_geca+_fcacf] & _cgca.Data[_fgfcd+_fcacf])
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], ^(_bggf.Data[_gcfe] & _cgca.Data[_bgcgf]), _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	case PixNotPixSrcXorDst:
		for _effa = 0; _effa < _eagba; _effa++ {
			_cgca.Data[_dbce] = _ggdf(_cgca.Data[_dbce], ^(_bggf.Data[_faadg] ^ _cgca.Data[_dbce]), _egef)
			_dbce += _cgca.BytesPerLine
			_faadg += _bggf.BytesPerLine
		}
		if _aabb {
			for _effa = 0; _effa < _eagba; _effa++ {
				for _fcacf = 0; _fcacf < _acddb; _fcacf++ {
					_cgca.Data[_fgfcd+_fcacf] = ^(_bggf.Data[_geca+_fcacf] ^ _cgca.Data[_fgfcd+_fcacf])
				}
				_fgfcd += _cgca.BytesPerLine
				_geca += _bggf.BytesPerLine
			}
		}
		if _fcgf {
			for _effa = 0; _effa < _eagba; _effa++ {
				_cgca.Data[_bgcgf] = _ggdf(_cgca.Data[_bgcgf], ^(_bggf.Data[_gcfe] ^ _cgca.Data[_bgcgf]), _egeg)
				_bgcgf += _cgca.BytesPerLine
				_gcfe += _bggf.BytesPerLine
			}
		}
	default:
		_a.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _cbcc)
		return _e.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func _gcca(_dfcf _d.Gray) _d.Gray { _dfcf.Y >>= 4; _dfcf.Y |= _dfcf.Y << 4; return _dfcf }
func (_faff *NRGBA16) Set(x, y int, c _d.Color) {
	_dagc := y*_faff.BytesPerLine + x*3/2
	if _dagc+1 >= len(_faff.Data) {
		return
	}
	_ceba := NRGBA16Model.Convert(c).(_d.NRGBA)
	_faff.setNRGBA(x, y, _dagc, _ceba)
}
func (_dfgdg *Monochrome) ExpandBinary(factor int) (*Monochrome, error) {
	if !IsPowerOf2(uint(factor)) {
		return nil, _db.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0065\u0078\u0070\u0061\u006e\u0064\u0020b\u0069n\u0061\u0072\u0079\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", factor)
	}
	return _b(_dfgdg, factor)
}

type NRGBA64 struct{ ImageBase }
type Gray interface {
	GrayAt(_fcba, _dbbg int) _d.Gray
	SetGray(_faad, _gbef int, _bae _d.Gray)
}

func _cecc(_daa _ee.Image) (Image, error) {
	if _gadf, _bda := _daa.(*Gray16); _bda {
		return _gadf.Copy(), nil
	}
	_gfdd := _daa.Bounds()
	_gfba, _bbg := NewImage(_gfdd.Max.X, _gfdd.Max.Y, 16, 1, nil, nil, nil)
	if _bbg != nil {
		return nil, _bbg
	}
	_gbgd(_daa, _gfba, _gfdd)
	return _gfba, nil
}
func (_gabb *RGBA32) ColorModel() _d.Model { return _d.NRGBAModel }
func (_cefd *Gray4) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _cefd.Width, Y: _cefd.Height}}
}
func (_efc *Monochrome) GrayAt(x, y int) _d.Gray {
	_ggae, _ := ColorAtGray1BPC(x, y, _efc.BytesPerLine, _efc.Data, _efc.Decode)
	return _ggae
}
func ColorAtGray4BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_d.Gray, error) {
	_bdba := y*bytesPerLine + x>>1
	if _bdba >= len(data) {
		return _d.Gray{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_cdb := data[_bdba] >> uint(4-(x&1)*4) & 0xf
	if len(decode) == 2 {
		_cdb = uint8(uint32(LinearInterpolate(float64(_cdb), 0, 15, decode[0], decode[1])) & 0xf)
	}
	return _d.Gray{Y: _cdb * 17 & 0xff}, nil
}
func (_acb *ImageBase) MakeAlpha() { _acb.newAlpha() }

var _ RGBA = &RGBA32{}

func _cefb(_bfb *Monochrome, _gbb, _fbec int, _bdegf, _degg int, _adea RasterOperator) {
	var (
		_edeb         int
		_geac         byte
		_aeee, _gafeg int
		_gceb         int
	)
	_adcf := _bdegf >> 3
	_dcdg := _bdegf & 7
	if _dcdg > 0 {
		_geac = _gbdbe[_dcdg]
	}
	_edeb = _bfb.BytesPerLine*_fbec + (_gbb >> 3)
	switch _adea {
	case PixClr:
		for _aeee = 0; _aeee < _degg; _aeee++ {
			_gceb = _edeb + _aeee*_bfb.BytesPerLine
			for _gafeg = 0; _gafeg < _adcf; _gafeg++ {
				_bfb.Data[_gceb] = 0x0
				_gceb++
			}
			if _dcdg > 0 {
				_bfb.Data[_gceb] = _ggdf(_bfb.Data[_gceb], 0x0, _geac)
			}
		}
	case PixSet:
		for _aeee = 0; _aeee < _degg; _aeee++ {
			_gceb = _edeb + _aeee*_bfb.BytesPerLine
			for _gafeg = 0; _gafeg < _adcf; _gafeg++ {
				_bfb.Data[_gceb] = 0xff
				_gceb++
			}
			if _dcdg > 0 {
				_bfb.Data[_gceb] = _ggdf(_bfb.Data[_gceb], 0xff, _geac)
			}
		}
	case PixNotDst:
		for _aeee = 0; _aeee < _degg; _aeee++ {
			_gceb = _edeb + _aeee*_bfb.BytesPerLine
			for _gafeg = 0; _gafeg < _adcf; _gafeg++ {
				_bfb.Data[_gceb] = ^_bfb.Data[_gceb]
				_gceb++
			}
			if _dcdg > 0 {
				_bfb.Data[_gceb] = _ggdf(_bfb.Data[_gceb], ^_bfb.Data[_gceb], _geac)
			}
		}
	}
}
func _bgf(_caae _d.NYCbCrA) _d.RGBA {
	_eefg, _ggee, _afbda, _fgb := _egee(_caae).RGBA()
	return _d.RGBA{R: uint8(_eefg >> 8), G: uint8(_ggee >> 8), B: uint8(_afbda >> 8), A: uint8(_fgb >> 8)}
}
func (_cagdc *NRGBA16) setNRGBA(_fbd, _dfad, _fdcf int, _eddb _d.NRGBA) {
	if _fbd*3%2 == 0 {
		_cagdc.Data[_fdcf] = (_eddb.R>>4)<<4 | (_eddb.G >> 4)
		_cagdc.Data[_fdcf+1] = (_eddb.B>>4)<<4 | (_cagdc.Data[_fdcf+1] & 0xf)
	} else {
		_cagdc.Data[_fdcf] = (_cagdc.Data[_fdcf] & 0xf0) | (_eddb.R >> 4)
		_cagdc.Data[_fdcf+1] = (_eddb.G>>4)<<4 | (_eddb.B >> 4)
	}
	if _cagdc.Alpha != nil {
		_geaa := _dfad * BytesPerLine(_cagdc.Width, 4, 1)
		if _geaa < len(_cagdc.Alpha) {
			if _fbd%2 == 0 {
				_cagdc.Alpha[_geaa] = (_eddb.A>>uint(4))<<uint(4) | (_cagdc.Alpha[_fdcf] & 0xf)
			} else {
				_cagdc.Alpha[_geaa] = (_cagdc.Alpha[_geaa] & 0xf0) | (_eddb.A >> uint(4))
			}
		}
	}
}
func (_ceee *Monochrome) setGrayBit(_dcfe, _bef int) { _ceee.Data[_dcfe] |= 0x80 >> uint(_bef&7) }

type shift int

func (_gdcea *NRGBA16) NRGBAAt(x, y int) _d.NRGBA {
	_bcea, _ := ColorAtNRGBA16(x, y, _gdcea.Width, _gdcea.BytesPerLine, _gdcea.Data, _gdcea.Alpha, _gdcea.Decode)
	return _bcea
}
func _eegd(_afc _d.NRGBA) _d.CMYK {
	_fec, _ecec, _afbd, _ := _afc.RGBA()
	_edd, _fdfg, _ddf, _gbd := _d.RGBToCMYK(uint8(_fec>>8), uint8(_ecec>>8), uint8(_afbd>>8))
	return _d.CMYK{C: _edd, M: _fdfg, Y: _ddf, K: _gbd}
}
func (_fff *Gray4) Copy() Image { return &Gray4{ImageBase: _fff.copy()} }

var _ Image = &RGBA32{}

func _adfe(_bfdf _d.CMYK) _d.RGBA {
	_effg, _caea, _cfa := _d.CMYKToRGB(_bfdf.C, _bfdf.M, _bfdf.Y, _bfdf.K)
	return _d.RGBA{R: _effg, G: _caea, B: _cfa, A: 0xff}
}
func (_ecb *CMYK32) CMYKAt(x, y int) _d.CMYK {
	_aceg, _ := ColorAtCMYK(x, y, _ecb.Width, _ecb.Data, _ecb.Decode)
	return _aceg
}
func (_egd *CMYK32) Base() *ImageBase { return &_egd.ImageBase }

var _ NRGBA = &NRGBA16{}

func (_egbf *Gray8) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _egbf.Width, Y: _egbf.Height}}
}
func (_eefgb *Gray8) At(x, y int) _d.Color { _cbff, _ := _eefgb.ColorAt(x, y); return _cbff }
func (_bgad *NRGBA64) NRGBA64At(x, y int) _d.NRGBA64 {
	_becgf, _ := ColorAtNRGBA64(x, y, _bgad.Width, _bgad.Data, _bgad.Alpha, _bgad.Decode)
	return _becgf
}

var ErrInvalidImage = _e.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")

func (_eab *Monochrome) getBitAt(_bfcg, _gfg int) bool {
	_ddab := _gfg*_eab.BytesPerLine + (_bfcg >> 3)
	_bfde := _bfcg & 0x07
	_ebc := uint(7 - _bfde)
	if _ddab > len(_eab.Data)-1 {
		return false
	}
	if (_eab.Data[_ddab]>>_ebc)&0x01 >= 1 {
		return true
	}
	return false
}
func (_bgbg *Gray16) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _bgbg.Width, Y: _bgbg.Height}}
}
func _gec(_cfe _ee.Image) (Image, error) {
	if _ece, _bdgb := _cfe.(*CMYK32); _bdgb {
		return _ece.Copy(), nil
	}
	_gda := _cfe.Bounds()
	_gfec, _bedc := NewImage(_gda.Max.X, _gda.Max.Y, 8, 4, nil, nil, nil)
	if _bedc != nil {
		return nil, _bedc
	}
	switch _acge := _cfe.(type) {
	case CMYK:
		_adfb(_acge, _gfec.(CMYK), _gda)
	case Gray:
		_fdf(_acge, _gfec.(CMYK), _gda)
	case NRGBA:
		_dfee(_acge, _gfec.(CMYK), _gda)
	case RGBA:
		_dbbb(_acge, _gfec.(CMYK), _gda)
	default:
		_fega(_cfe, _gfec, _gda)
	}
	return _gfec, nil
}
func (_faac *Monochrome) IsUnpadded() bool { return (_faac.Width * _faac.Height) == len(_faac.Data) }
func (_bcce *Monochrome) getBit(_gcag, _fag int) uint8 {
	return _bcce.Data[_gcag+(_fag>>3)] >> uint(7-(_fag&7)) & 1
}
func _df(_dff, _gfe *Monochrome, _bce []byte, _agb int) (_cfbe error) {
	var (
		_fda, _cab, _bdd, _bga, _bbb, _ebd, _ffcf, _fcab int
		_fgf, _dda, _ebf, _dad                           uint32
		_dab, _dcc                                       byte
		_fcb                                             uint16
	)
	_gde := make([]byte, 4)
	_debf := make([]byte, 4)
	for _bdd = 0; _bdd < _dff.Height-1; _bdd, _bga = _bdd+2, _bga+1 {
		_fda = _bdd * _dff.BytesPerLine
		_cab = _bga * _gfe.BytesPerLine
		for _bbb, _ebd = 0, 0; _bbb < _agb; _bbb, _ebd = _bbb+4, _ebd+1 {
			for _ffcf = 0; _ffcf < 4; _ffcf++ {
				_fcab = _fda + _bbb + _ffcf
				if _fcab <= len(_dff.Data)-1 && _fcab < _fda+_dff.BytesPerLine {
					_gde[_ffcf] = _dff.Data[_fcab]
				} else {
					_gde[_ffcf] = 0x00
				}
				_fcab = _fda + _dff.BytesPerLine + _bbb + _ffcf
				if _fcab <= len(_dff.Data)-1 && _fcab < _fda+(2*_dff.BytesPerLine) {
					_debf[_ffcf] = _dff.Data[_fcab]
				} else {
					_debf[_ffcf] = 0x00
				}
			}
			_fgf = _f.BigEndian.Uint32(_gde)
			_dda = _f.BigEndian.Uint32(_debf)
			_ebf = _fgf & _dda
			_ebf |= _ebf << 1
			_dad = _fgf | _dda
			_dad &= _dad << 1
			_dda = _ebf | _dad
			_dda &= 0xaaaaaaaa
			_fgf = _dda | (_dda << 7)
			_dab = byte(_fgf >> 24)
			_dcc = byte((_fgf >> 8) & 0xff)
			_fcab = _cab + _ebd
			if _fcab+1 == len(_gfe.Data)-1 || _fcab+1 >= _cab+_gfe.BytesPerLine {
				if _cfbe = _gfe.setByte(_fcab, _bce[_dab]); _cfbe != nil {
					return _db.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _fcab)
				}
			} else {
				_fcb = (uint16(_bce[_dab]) << 8) | uint16(_bce[_dcc])
				if _cfbe = _gfe.setTwoBytes(_fcab, _fcb); _cfbe != nil {
					return _db.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fcab)
				}
				_ebd++
			}
		}
	}
	return nil
}

type Gray16 struct{ ImageBase }

func (_afba *NRGBA32) Set(x, y int, c _d.Color) {
	_dfeed := y*_afba.Width + x
	_bggb := 3 * _dfeed
	if _bggb+2 >= len(_afba.Data) {
		return
	}
	_gfde := _d.NRGBAModel.Convert(c).(_d.NRGBA)
	_afba.setRGBA(_dfeed, _gfde)
}
func _egee(_aab _d.NYCbCrA) _d.NRGBA {
	_edff := int32(_aab.Y) * 0x10101
	_eagb := int32(_aab.Cb) - 128
	_eefd := int32(_aab.Cr) - 128
	_bage := _edff + 91881*_eefd
	if uint32(_bage)&0xff000000 == 0 {
		_bage >>= 8
	} else {
		_bage = ^(_bage >> 31) & 0xffff
	}
	_bfga := _edff - 22554*_eagb - 46802*_eefd
	if uint32(_bfga)&0xff000000 == 0 {
		_bfga >>= 8
	} else {
		_bfga = ^(_bfga >> 31) & 0xffff
	}
	_eeb := _edff + 116130*_eagb
	if uint32(_eeb)&0xff000000 == 0 {
		_eeb >>= 8
	} else {
		_eeb = ^(_eeb >> 31) & 0xffff
	}
	return _d.NRGBA{R: uint8(_bage >> 8), G: uint8(_bfga >> 8), B: uint8(_eeb >> 8), A: _aab.A}
}
func (_ddda *ImageBase) setEightFullBytes(_agbd int, _fcc uint64) error {
	if _agbd+7 > len(_ddda.Data)-1 {
		return _e.New("\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_ddda.Data[_agbd] = byte((_fcc & 0xff00000000000000) >> 56)
	_ddda.Data[_agbd+1] = byte((_fcc & 0xff000000000000) >> 48)
	_ddda.Data[_agbd+2] = byte((_fcc & 0xff0000000000) >> 40)
	_ddda.Data[_agbd+3] = byte((_fcc & 0xff00000000) >> 32)
	_ddda.Data[_agbd+4] = byte((_fcc & 0xff000000) >> 24)
	_ddda.Data[_agbd+5] = byte((_fcc & 0xff0000) >> 16)
	_ddda.Data[_agbd+6] = byte((_fcc & 0xff00) >> 8)
	_ddda.Data[_agbd+7] = byte(_fcc & 0xff)
	return nil
}
func _egag(_cffc uint8) bool {
	if _cffc == 0 || _cffc == 255 {
		return true
	}
	return false
}
func _afdg(_cegb, _afbbb RGBA, _cgec _ee.Rectangle) {
	for _dbfc := 0; _dbfc < _cgec.Max.X; _dbfc++ {
		for _bbac := 0; _bbac < _cgec.Max.Y; _bbac++ {
			_afbbb.SetRGBA(_dbfc, _bbac, _cegb.RGBAAt(_dbfc, _bbac))
		}
	}
}
func (_agd *CMYK32) SetCMYK(x, y int, c _d.CMYK) {
	_adef := 4 * (y*_agd.Width + x)
	if _adef+3 >= len(_agd.Data) {
		return
	}
	_agd.Data[_adef] = c.C
	_agd.Data[_adef+1] = c.M
	_agd.Data[_adef+2] = c.Y
	_agd.Data[_adef+3] = c.K
}
func _def() (_gbf []byte) {
	_gbf = make([]byte, 256)
	for _cgeb := 0; _cgeb < 256; _cgeb++ {
		_fde := byte(_cgeb)
		_gbf[_fde] = (_fde & 0x01) | ((_fde & 0x04) >> 1) | ((_fde & 0x10) >> 2) | ((_fde & 0x40) >> 3) | ((_fde & 0x02) << 3) | ((_fde & 0x08) << 2) | ((_fde & 0x20) << 1) | (_fde & 0x80)
	}
	return _gbf
}
func (_ddaf *Gray4) At(x, y int) _d.Color { _cafd, _ := _ddaf.ColorAt(x, y); return _cafd }
func _fae(_fbbc _d.RGBA) _d.CMYK {
	_fdd, _cbd, _bbba, _fggf := _d.RGBToCMYK(_fbbc.R, _fbbc.G, _fbbc.B)
	return _d.CMYK{C: _fdd, M: _cbd, Y: _bbba, K: _fggf}
}
func _cccf(_egab *_ee.NYCbCrA, _fdff RGBA, _fbcb _ee.Rectangle) {
	for _gfae := 0; _gfae < _fbcb.Max.X; _gfae++ {
		for _abbbb := 0; _abbbb < _fbcb.Max.Y; _abbbb++ {
			_faeee := _egab.NYCbCrAAt(_gfae, _abbbb)
			_fdff.SetRGBA(_gfae, _abbbb, _bgf(_faeee))
		}
	}
}
func _ggfga(_bedf _ee.Image) (Image, error) {
	if _dffa, _cbdb := _bedf.(*Gray4); _cbdb {
		return _dffa.Copy(), nil
	}
	_dcab := _bedf.Bounds()
	_dbae, _fdda := NewImage(_dcab.Max.X, _dcab.Max.Y, 4, 1, nil, nil, nil)
	if _fdda != nil {
		return nil, _fdda
	}
	_gbgd(_bedf, _dbae, _dcab)
	return _dbae, nil
}
func AutoThresholdTriangle(histogram [256]int) uint8 {
	var _afde, _gcbf, _acbd, _cgac int
	for _feag := 0; _feag < len(histogram); _feag++ {
		if histogram[_feag] > 0 {
			_afde = _feag
			break
		}
	}
	if _afde > 0 {
		_afde--
	}
	for _cccc := 255; _cccc > 0; _cccc-- {
		if histogram[_cccc] > 0 {
			_cgac = _cccc
			break
		}
	}
	if _cgac < 255 {
		_cgac++
	}
	for _caag := 0; _caag < 256; _caag++ {
		if histogram[_caag] > _gcbf {
			_acbd = _caag
			_gcbf = histogram[_caag]
		}
	}
	var _bade bool
	if (_acbd - _afde) < (_cgac - _acbd) {
		_bade = true
		var _gcagd int
		_begd := 255
		for _gcagd < _begd {
			_ggaaf := histogram[_gcagd]
			histogram[_gcagd] = histogram[_begd]
			histogram[_begd] = _ggaaf
			_gcagd++
			_begd--
		}
		_afde = 255 - _cgac
		_acbd = 255 - _acbd
	}
	if _afde == _acbd {
		return uint8(_afde)
	}
	_bagb := float64(histogram[_acbd])
	_gafc := float64(_afde - _acbd)
	_gcccd := _g.Sqrt(_bagb*_bagb + _gafc*_gafc)
	_bagb /= _gcccd
	_gafc /= _gcccd
	_gcccd = _bagb*float64(_afde) + _gafc*float64(histogram[_afde])
	_facdc := _afde
	var _dbef float64
	for _badeb := _afde + 1; _badeb <= _acbd; _badeb++ {
		_aefb := _bagb*float64(_badeb) + _gafc*float64(histogram[_badeb]) - _gcccd
		if _aefb > _dbef {
			_facdc = _badeb
			_dbef = _aefb
		}
	}
	_facdc--
	if _bade {
		var _fabb int
		_bfaf := 255
		for _fabb < _bfaf {
			_bced := histogram[_fabb]
			histogram[_fabb] = histogram[_bfaf]
			histogram[_bfaf] = _bced
			_fabb++
			_bfaf--
		}
		return uint8(255 - _facdc)
	}
	return uint8(_facdc)
}
func (_abbb *Gray8) Histogram() (_abba [256]int) {
	for _becg := 0; _becg < len(_abbb.Data); _becg++ {
		_abba[_abbb.Data[_becg]]++
	}
	return _abba
}
func _bbef(_ceab *Monochrome, _agfe, _edeaac, _cfgab, _cbdcd int, _dddd RasterOperator) {
	if _agfe < 0 {
		_cfgab += _agfe
		_agfe = 0
	}
	_ababf := _agfe + _cfgab - _ceab.Width
	if _ababf > 0 {
		_cfgab -= _ababf
	}
	if _edeaac < 0 {
		_cbdcd += _edeaac
		_edeaac = 0
	}
	_cafg := _edeaac + _cbdcd - _ceab.Height
	if _cafg > 0 {
		_cbdcd -= _cafg
	}
	if _cfgab <= 0 || _cbdcd <= 0 {
		return
	}
	if (_agfe & 7) == 0 {
		_cefb(_ceab, _agfe, _edeaac, _cfgab, _cbdcd, _dddd)
	} else {
		_bgdb(_ceab, _agfe, _edeaac, _cfgab, _cbdcd, _dddd)
	}
}

var _ Image = &Gray16{}
var (
	MonochromeConverter = ConverterFunc(_bfc)
	Gray2Converter      = ConverterFunc(_badd)
	Gray4Converter      = ConverterFunc(_ggfga)
	GrayConverter       = ConverterFunc(_gbdb)
	Gray16Converter     = ConverterFunc(_cecc)
	NRGBA16Converter    = ConverterFunc(_abbf)
	NRGBAConverter      = ConverterFunc(_ceg)
	NRGBA64Converter    = ConverterFunc(_egafg)
	RGBAConverter       = ConverterFunc(_ebbbe)
	CMYKConverter       = ConverterFunc(_gec)
)

func (_aeea *NRGBA32) Base() *ImageBase   { return &_aeea.ImageBase }
func (_becf *Gray4) ColorModel() _d.Model { return Gray4Model }
func (_edee *RGBA32) SetRGBA(x, y int, c _d.RGBA) {
	_gggb := y*_edee.Width + x
	_debe := 3 * _gggb
	if _debe+2 >= len(_edee.Data) {
		return
	}
	_edee.setRGBA(_gggb, c)
}
func _edag(_bfba _ee.Image, _ebca Image, _gcgd _ee.Rectangle) {
	if _bedd, _fgff := _bfba.(SMasker); _fgff && _bedd.HasAlpha() {
		_ebca.(SMasker).MakeAlpha()
	}
	_fega(_bfba, _ebca, _gcgd)
}
func LinearInterpolate(x, xmin, xmax, ymin, ymax float64) float64 {
	if _g.Abs(xmax-xmin) < 0.000001 {
		return ymin
	}
	_deff := ymin + (x-xmin)*(ymax-ymin)/(xmax-xmin)
	return _deff
}
func (_fcbd *NRGBA16) Copy() Image { return &NRGBA16{ImageBase: _fcbd.copy()} }
func (_aebe *Gray2) Copy() Image   { return &Gray2{ImageBase: _aebe.copy()} }
func _ddcc(_aebeb NRGBA, _bgac Gray, _ggaa _ee.Rectangle) {
	for _fdc := 0; _fdc < _ggaa.Max.X; _fdc++ {
		for _cfed := 0; _cfed < _ggaa.Max.Y; _cfed++ {
			_becac := _gcc(_aebeb.NRGBAAt(_fdc, _cfed))
			_bgac.SetGray(_fdc, _cfed, _becac)
		}
	}
}

var _ Image = &CMYK32{}

func _gbdb(_gcdd _ee.Image) (Image, error) {
	if _egcc, _adab := _gcdd.(*Gray8); _adab {
		return _egcc.Copy(), nil
	}
	_ggd := _gcdd.Bounds()
	_ffcd, _fade := NewImage(_ggd.Max.X, _ggd.Max.Y, 8, 1, nil, nil, nil)
	if _fade != nil {
		return nil, _fade
	}
	_gbgd(_gcdd, _ffcd, _ggd)
	return _ffcd, nil
}
func _fefe(_cagdb _d.Color) _d.Color {
	_aeag := _d.NRGBAModel.Convert(_cagdb).(_d.NRGBA)
	return _faef(_aeag)
}
func _afb(_edf, _bec *Monochrome, _ccf []byte, _bcc int) (_gab error) {
	var (
		_dgeg, _efg, _gdde, _gee, _ddc, _ecc, _ebdb, _ega int
		_fcg, _fdg, _cbg, _agg                            uint32
		_dceg, _gfc                                       byte
		_cba                                              uint16
	)
	_faa := make([]byte, 4)
	_adf := make([]byte, 4)
	for _gdde = 0; _gdde < _edf.Height-1; _gdde, _gee = _gdde+2, _gee+1 {
		_dgeg = _gdde * _edf.BytesPerLine
		_efg = _gee * _bec.BytesPerLine
		for _ddc, _ecc = 0, 0; _ddc < _bcc; _ddc, _ecc = _ddc+4, _ecc+1 {
			for _ebdb = 0; _ebdb < 4; _ebdb++ {
				_ega = _dgeg + _ddc + _ebdb
				if _ega <= len(_edf.Data)-1 && _ega < _dgeg+_edf.BytesPerLine {
					_faa[_ebdb] = _edf.Data[_ega]
				} else {
					_faa[_ebdb] = 0x00
				}
				_ega = _dgeg + _edf.BytesPerLine + _ddc + _ebdb
				if _ega <= len(_edf.Data)-1 && _ega < _dgeg+(2*_edf.BytesPerLine) {
					_adf[_ebdb] = _edf.Data[_ega]
				} else {
					_adf[_ebdb] = 0x00
				}
			}
			_fcg = _f.BigEndian.Uint32(_faa)
			_fdg = _f.BigEndian.Uint32(_adf)
			_cbg = _fcg & _fdg
			_cbg |= _cbg << 1
			_agg = _fcg | _fdg
			_agg &= _agg << 1
			_fdg = _cbg & _agg
			_fdg &= 0xaaaaaaaa
			_fcg = _fdg | (_fdg << 7)
			_dceg = byte(_fcg >> 24)
			_gfc = byte((_fcg >> 8) & 0xff)
			_ega = _efg + _ecc
			if _ega+1 == len(_bec.Data)-1 || _ega+1 >= _efg+_bec.BytesPerLine {
				if _gab = _bec.setByte(_ega, _ccf[_dceg]); _gab != nil {
					return _db.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _ega)
				}
			} else {
				_cba = (uint16(_ccf[_dceg]) << 8) | uint16(_ccf[_gfc])
				if _gab = _bec.setTwoBytes(_ega, _cba); _gab != nil {
					return _db.Errorf("s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _ega)
				}
				_ecc++
			}
		}
	}
	return nil
}
func _faef(_gdee _d.NRGBA) _d.NRGBA {
	_gdee.R = _gdee.R>>4 | (_gdee.R>>4)<<4
	_gdee.G = _gdee.G>>4 | (_gdee.G>>4)<<4
	_gdee.B = _gdee.B>>4 | (_gdee.B>>4)<<4
	return _gdee
}
func _cdbe(_gebea, _aabe uint8) uint8 {
	if _gebea < _aabe {
		return 255
	}
	return 0
}
func (_bcb *Gray16) Validate() error {
	if len(_bcb.Data) != _bcb.Height*_bcb.BytesPerLine {
		return ErrInvalidImage
	}
	return nil
}
func (_abd *Gray2) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtGray2BPC(x, y, _abd.BytesPerLine, _abd.Data, _abd.Decode)
}
func (_fcad *ImageBase) setFourBytes(_faeg int, _ceb uint32) error {
	if _faeg+3 > len(_fcad.Data)-1 {
		return _db.Errorf("\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _faeg)
	}
	_fcad.Data[_faeg] = byte((_ceb & 0xff000000) >> 24)
	_fcad.Data[_faeg+1] = byte((_ceb & 0xff0000) >> 16)
	_fcad.Data[_faeg+2] = byte((_ceb & 0xff00) >> 8)
	_fcad.Data[_faeg+3] = byte(_ceb & 0xff)
	return nil
}
func _dfee(_eca NRGBA, _fafa CMYK, _egde _ee.Rectangle) {
	for _aeg := 0; _aeg < _egde.Max.X; _aeg++ {
		for _ffg := 0; _ffg < _egde.Max.Y; _ffg++ {
			_ecg := _eca.NRGBAAt(_aeg, _ffg)
			_fafa.SetCMYK(_aeg, _ffg, _eegd(_ecg))
		}
	}
}
func _dgf(_fgd *Monochrome, _efa ...int) (_gdd *Monochrome, _cae error) {
	if _fgd == nil {
		return nil, _e.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_efa) == 0 {
		return nil, _e.New("\u0074h\u0065\u0072e\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020o\u006e\u0065\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006f\u0066 \u0072\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e")
	}
	_dce := _def()
	_gdd = _fgd
	for _, _ade := range _efa {
		if _ade <= 0 {
			break
		}
		_gdd, _cae = _ege(_gdd, _ade, _dce)
		if _cae != nil {
			return nil, _cae
		}
	}
	return _gdd, nil
}
func (_bab *Monochrome) ResolveDecode() error {
	if len(_bab.Decode) != 2 {
		return nil
	}
	if _bab.Decode[0] == 1 && _bab.Decode[1] == 0 {
		if _bbd := _bab.InverseData(); _bbd != nil {
			return _bbd
		}
		_bab.Decode = nil
	}
	return nil
}

type RGBA32 struct{ ImageBase }

func _fdfa(_bcfe _d.Color) _d.Color {
	_fdeef := _d.GrayModel.Convert(_bcfe).(_d.Gray)
	return _gcca(_fdeef)
}
func _adfb(_bfd, _egfc CMYK, _fgg _ee.Rectangle) {
	for _ffb := 0; _ffb < _fgg.Max.X; _ffb++ {
		for _dcg := 0; _dcg < _fgg.Max.Y; _dcg++ {
			_egfc.SetCMYK(_ffb, _dcg, _bfd.CMYKAt(_ffb, _dcg))
		}
	}
}

var _ Image = &Gray4{}

func _dbc() {
	for _eede := 0; _eede < 256; _eede++ {
		_bcdf[_eede] = uint8(_eede&0x1) + (uint8(_eede>>1) & 0x1) + (uint8(_eede>>2) & 0x1) + (uint8(_eede>>3) & 0x1) + (uint8(_eede>>4) & 0x1) + (uint8(_eede>>5) & 0x1) + (uint8(_eede>>6) & 0x1) + (uint8(_eede>>7) & 0x1)
	}
}
func (_abag *NRGBA32) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _abag.Width, Y: _abag.Height}}
}
func ColorAtGray16BPC(x, y, bytesPerLine int, data []byte, decode []float64) (_d.Gray16, error) {
	_ecee := (y*bytesPerLine/2 + x) * 2
	if _ecee+1 >= len(data) {
		return _d.Gray16{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_bfgbf := uint16(data[_ecee])<<8 | uint16(data[_ecee+1])
	if len(decode) == 2 {
		_bfgbf = uint16(uint64(LinearInterpolate(float64(_bfgbf), 0, 65535, decode[0], decode[1])))
	}
	return _d.Gray16{Y: _bfgbf}, nil
}
func (_bafe *NRGBA32) ColorModel() _d.Model { return _d.NRGBAModel }
func _bfc(_ddfd _ee.Image) (Image, error) {
	if _bgc, _becd := _ddfd.(*Monochrome); _becd {
		return _bgc, nil
	}
	_cdff := _ddfd.Bounds()
	var _bcd Gray
	switch _bdgf := _ddfd.(type) {
	case Gray:
		_bcd = _bdgf
	case NRGBA:
		_bcd = &Gray8{ImageBase: NewImageBase(_cdff.Max.X, _cdff.Max.Y, 8, 1, nil, nil, nil)}
		_bffc(_bcd, _bdgf, _cdff)
	case nrgba64:
		_bcd = &Gray8{ImageBase: NewImageBase(_cdff.Max.X, _cdff.Max.Y, 8, 1, nil, nil, nil)}
		_fbbg(_bcd, _bdgf, _cdff)
	default:
		_gabd, _bbae := GrayConverter.Convert(_ddfd)
		if _bbae != nil {
			return nil, _bbae
		}
		_bcd = _gabd.(Gray)
	}
	_eeeb, _efd := NewImage(_cdff.Max.X, _cdff.Max.Y, 1, 1, nil, nil, nil)
	if _efd != nil {
		return nil, _efd
	}
	_dbg := _eeeb.(*Monochrome)
	_fadd := AutoThresholdTriangle(GrayHistogram(_bcd))
	for _adde := 0; _adde < _cdff.Max.X; _adde++ {
		for _fdea := 0; _fdea < _cdff.Max.Y; _fdea++ {
			_efbg := _beg(_bcd.GrayAt(_adde, _fdea), monochromeModel(_fadd))
			_dbg.SetGray(_adde, _fdea, _efbg)
		}
	}
	return _eeeb, nil
}
func (_cgc *Monochrome) Copy() Image {
	return &Monochrome{ImageBase: _cgc.ImageBase.copy(), ModelThreshold: _cgc.ModelThreshold}
}

var _ Image = &NRGBA32{}

func (_agbe *RGBA32) RGBAAt(x, y int) _d.RGBA {
	_babc, _ := ColorAtRGBA32(x, y, _agbe.Width, _agbe.Data, _agbe.Alpha, _agbe.Decode)
	return _babc
}
func ColorAtRGBA32(x, y, width int, data, alpha []byte, decode []float64) (_d.RGBA, error) {
	_egga := y*width + x
	_bfbe := 3 * _egga
	if _bfbe+2 >= len(data) {
		return _d.RGBA{}, _db.Errorf("\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006ea\u0074\u0065\u0073\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064,\u0020\u0025\u0064\u0029", x, y)
	}
	_dfaa := uint8(0xff)
	if alpha != nil && len(alpha) > _egga {
		_dfaa = alpha[_egga]
	}
	_aaaa, _beaf, _cadfae := data[_bfbe], data[_bfbe+1], data[_bfbe+2]
	if len(decode) == 6 {
		_aaaa = uint8(uint32(LinearInterpolate(float64(_aaaa), 0, 255, decode[0], decode[1])) & 0xff)
		_beaf = uint8(uint32(LinearInterpolate(float64(_beaf), 0, 255, decode[2], decode[3])) & 0xff)
		_cadfae = uint8(uint32(LinearInterpolate(float64(_cadfae), 0, 255, decode[4], decode[5])) & 0xff)
	}
	return _d.RGBA{R: _aaaa, G: _beaf, B: _cadfae, A: _dfaa}, nil
}
func (_cbc *CMYK32) ColorModel() _d.Model { return _d.CMYKModel }
func (_efgb *RGBA32) Validate() error {
	if len(_efgb.Data) != 3*_efgb.Width*_efgb.Height {
		return _e.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func BytesPerLine(width, bitsPerComponent, colorComponents int) int {
	return ((width*bitsPerComponent)*colorComponents + 7) >> 3
}
func _egafg(_edgb _ee.Image) (Image, error) {
	if _gbfd, _bcgb := _edgb.(*NRGBA64); _bcgb {
		return _gbfd.Copy(), nil
	}
	_fbga, _caead, _ccff := _bbbd(_edgb, 2)
	_cadf, _afdcg := NewImage(_fbga.Max.X, _fbga.Max.Y, 16, 3, nil, _ccff, nil)
	if _afdcg != nil {
		return nil, _afdcg
	}
	_edag(_edgb, _cadf, _fbga)
	if len(_ccff) != 0 && !_caead {
		if _fceb := _efaga(_ccff, _cadf); _fceb != nil {
			return nil, _fceb
		}
	}
	return _cadf, nil
}
func (_dafc *Gray16) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtGray16BPC(x, y, _dafc.BytesPerLine, _dafc.Data, _dafc.Decode)
}
func (_bgcg *Gray2) ColorModel() _d.Model { return Gray2Model }
func (_bgcf *NRGBA32) Validate() error {
	if len(_bgcf.Data) != 3*_bgcf.Width*_bgcf.Height {
		return _e.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}
func _ebb(_ada _d.NRGBA64) _d.RGBA {
	_caeb, _cbac, _bgae, _bccc := _ada.RGBA()
	return _d.RGBA{R: uint8(_caeb >> 8), G: uint8(_cbac >> 8), B: uint8(_bgae >> 8), A: uint8(_bccc >> 8)}
}
func _ege(_caef *Monochrome, _aaf int, _aac []byte) (_caa *Monochrome, _aef error) {
	const _dbb = "\u0072\u0065d\u0075\u0063\u0065R\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079"
	if _caef == nil {
		return nil, _e.New("\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _aaf < 1 || _aaf > 4 {
		return nil, _e.New("\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _caef.Height <= 1 {
		return nil, _e.New("\u0073\u006f\u0075rc\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020m\u0075s\u0074 \u0062e\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027")
	}
	_caa = _bed(_caef.Width/2, _caef.Height/2)
	if _aac == nil {
		_aac = _def()
	}
	_faf := _bgca(_caef.BytesPerLine, 2*_caa.BytesPerLine)
	switch _aaf {
	case 1:
		_aef = _cge(_caef, _caa, _aac, _faf)
	case 2:
		_aef = _df(_caef, _caa, _aac, _faf)
	case 3:
		_aef = _afb(_caef, _caa, _aac, _faf)
	case 4:
		_aef = _dfe(_caef, _caa, _aac, _faf)
	}
	if _aef != nil {
		return nil, _aef
	}
	return _caa, nil
}
func (_dede *NRGBA32) NRGBAAt(x, y int) _d.NRGBA {
	_age, _ := ColorAtNRGBA32(x, y, _dede.Width, _dede.Data, _dede.Alpha, _dede.Decode)
	return _age
}
func _babb(_bcg, _egdb Gray, _fdgb _ee.Rectangle) {
	for _fecb := 0; _fecb < _fdgb.Max.X; _fecb++ {
		for _ggfe := 0; _ggfe < _fdgb.Max.Y; _ggfe++ {
			_egdb.SetGray(_fecb, _ggfe, _bcg.GrayAt(_fecb, _ggfe))
		}
	}
}
func (_fege *Monochrome) setGray(_aeb int, _acee _d.Gray, _fbac int) {
	if _acee.Y == 0 {
		_fege.clearBit(_fbac, _aeb)
	} else {
		_fege.setGrayBit(_fbac, _aeb)
	}
}
func _agbfc(_bcbed Gray, _fdcfb RGBA, _cbbd _ee.Rectangle) {
	for _fegcc := 0; _fegcc < _cbbd.Max.X; _fegcc++ {
		for _eadb := 0; _eadb < _cbbd.Max.Y; _eadb++ {
			_ggbg := _bcbed.GrayAt(_fegcc, _eadb)
			_fdcfb.SetRGBA(_fegcc, _eadb, _dba(_ggbg))
		}
	}
}

type colorConverter struct {
	_dfeb func(_ebee _ee.Image) (Image, error)
}

func (_cef *CMYK32) ColorAt(x, y int) (_d.Color, error) {
	return ColorAtCMYK(x, y, _cef.Width, _cef.Data, _cef.Decode)
}
func IsGrayImgBlackAndWhite(i *_ee.Gray) bool { return _ggaec(i) }
func _gcc(_dfa _d.NRGBA) _d.Gray {
	_cbf, _aba, _gcce, _ := _dfa.RGBA()
	_ebfb := (19595*_cbf + 38470*_aba + 7471*_gcce + 1<<15) >> 24
	return _d.Gray{Y: uint8(_ebfb)}
}

type Gray8 struct{ ImageBase }

func (_eceee *RGBA32) Set(x, y int, c _d.Color) {
	_eecd := y*_eceee.Width + x
	_ggge := 3 * _eecd
	if _ggge+2 >= len(_eceee.Data) {
		return
	}
	_abec := _d.RGBAModel.Convert(c).(_d.RGBA)
	_eceee.setRGBA(_eecd, _abec)
}
func _adcc(_edgd, _fgccb NRGBA, _gefb _ee.Rectangle) {
	for _cbcce := 0; _cbcce < _gefb.Max.X; _cbcce++ {
		for _cefbc := 0; _cefbc < _gefb.Max.Y; _cefbc++ {
			_fgccb.SetNRGBA(_cbcce, _cefbc, _edgd.NRGBAAt(_cbcce, _cefbc))
		}
	}
}
func NextPowerOf2(n uint) uint {
	if IsPowerOf2(n) {
		return n
	}
	return 1 << (_bdgc(n) + 1)
}
func FromGoImage(i _ee.Image) (Image, error) {
	switch _daca := i.(type) {
	case Image:
		return _daca.Copy(), nil
	case Gray:
		return GrayConverter.Convert(i)
	case *_ee.Gray16:
		return Gray16Converter.Convert(i)
	case CMYK:
		return CMYKConverter.Convert(i)
	case *_ee.NRGBA64:
		return NRGBA64Converter.Convert(i)
	default:
		return NRGBAConverter.Convert(i)
	}
}

var _ Gray = &Gray8{}

type RGBA interface {
	RGBAAt(_caacc, _cfeb int) _d.RGBA
	SetRGBA(_gggd, _gcbe int, _ffba _d.RGBA)
}
type NRGBA16 struct{ ImageBase }
type monochromeModel uint8

func _bffc(_gcb Gray, _gcad NRGBA, _bbad _ee.Rectangle) {
	for _fdee := 0; _fdee < _bbad.Max.X; _fdee++ {
		for _ggb := 0; _ggb < _bbad.Max.Y; _ggb++ {
			_gbca := _adee(_gcad.NRGBAAt(_fdee, _ggb))
			_gcb.SetGray(_fdee, _ggb, _gbca)
		}
	}
}
func _ebbbe(_daeda _ee.Image) (Image, error) {
	if _acgd, _efcb := _daeda.(*RGBA32); _efcb {
		return _acgd.Copy(), nil
	}
	_cbfb, _eabg, _ddacf := _bbbd(_daeda, 1)
	_bgaa := &RGBA32{ImageBase: NewImageBase(_cbfb.Max.X, _cbfb.Max.Y, 8, 3, nil, _ddacf, nil)}
	_dbgb(_daeda, _bgaa, _cbfb)
	if len(_ddacf) != 0 && !_eabg {
		if _abdd := _efaga(_ddacf, _bgaa); _abdd != nil {
			return nil, _abdd
		}
	}
	return _bgaa, nil
}
func (_dbdc *Gray8) GrayAt(x, y int) _d.Gray {
	_baf, _ := ColorAtGray8BPC(x, y, _dbdc.BytesPerLine, _dbdc.Data, _dbdc.Decode)
	return _baf
}
func (_ecaf *Monochrome) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Monochrome, sx, sy int) error {
	return _gbcc(_ecaf, dx, dy, dw, dh, op, src, sx, sy)
}
func (_fbcc *RGBA32) Bounds() _ee.Rectangle {
	return _ee.Rectangle{Max: _ee.Point{X: _fbcc.Width, Y: _fbcc.Height}}
}
func (_cbda *NRGBA64) At(x, y int) _d.Color { _ffdgf, _ := _cbda.ColorAt(x, y); return _ffdgf }
func _bgbgd(_gbbb _ee.Image, _gffg Image, _ebgcd _ee.Rectangle) {
	if _efgg, _cdead := _gbbb.(SMasker); _cdead && _efgg.HasAlpha() {
		_gffg.(SMasker).MakeAlpha()
	}
	switch _efcf := _gbbb.(type) {
	case Gray:
		_fbad(_efcf, _gffg.(NRGBA), _ebgcd)
	case NRGBA:
		_adcc(_efcf, _gffg.(NRGBA), _ebgcd)
	case *_ee.NYCbCrA:
		_fgbcb(_efcf, _gffg.(NRGBA), _ebgcd)
	case CMYK:
		_cffb(_efcf, _gffg.(NRGBA), _ebgcd)
	case RGBA:
		_bcbc(_efcf, _gffg.(NRGBA), _ebgcd)
	case nrgba64:
		_becab(_efcf, _gffg.(NRGBA), _ebgcd)
	default:
		_fega(_gbbb, _gffg, _ebgcd)
	}
}
func (_debd *Monochrome) Set(x, y int, c _d.Color) {
	_cfd := y*_debd.BytesPerLine + x>>3
	if _cfd > len(_debd.Data)-1 {
		return
	}
	_ffcg := _debd.ColorModel().Convert(c).(_d.Gray)
	_debd.setGray(x, _ffcg, _cfd)
}
func (_bgag *NRGBA16) Validate() error {
	if len(_bgag.Data) != 3*_bgag.Width*_bgag.Height/2 {
		return _e.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006da\u0067\u0065\u0020\u0064\u0061\u0074\u0061 s\u0069\u007a\u0065\u0020f\u006f\u0072\u0020\u0070\u0072\u006f\u0076\u0069\u0064ed\u0020\u0064i\u006d\u0065\u006e\u0073\u0069\u006f\u006e\u0073")
	}
	return nil
}

type CMYK interface {
	CMYKAt(_gfb, _ffdd int) _d.CMYK
	SetCMYK(_gded, _fge int, _aaa _d.CMYK)
}
