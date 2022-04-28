package render

import (
	_g "errors"
	_de "fmt"
	_ab "image"
	_ea "image/color"
	_dc "image/draw"
	_fe "image/jpeg"
	_db "image/png"
	_e "math"
	_gd "os"
	_a "path/filepath"
	_d "strings"

	_b "bitbucket.org/shenghui0779/gopdf/common"
	_dd "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ad "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_fd "bitbucket.org/shenghui0779/gopdf/core"
	_aba "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_gb "bitbucket.org/shenghui0779/gopdf/model"
	_deb "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_gc "bitbucket.org/shenghui0779/gopdf/render/internal/context/imagerender"
	_df "github.com/adrg/sysfont"
	_ga "golang.org/x/image/draw"
)

func _cccb(_gfc string, _aee _ab.Image) error {
	_bead, _gad := _gd.Create(_gfc)
	if _gad != nil {
		return _gad
	}
	defer _bead.Close()
	return _db.Encode(_bead, _aee)
}

var (
	_bec = _g.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_dac = _g.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func _aea(_abe, _abef _ab.Image) _ab.Image {
	_ggbfc, _ddaf := _abef.Bounds().Size(), _abe.Bounds().Size()
	_dga, _bgg := _ggbfc.X, _ggbfc.Y
	if _ddaf.X > _dga {
		_dga = _ddaf.X
	}
	if _ddaf.Y > _bgg {
		_bgg = _ddaf.Y
	}
	_dgde := _ab.Rect(0, 0, _dga, _bgg)
	if _ggbfc.X != _dga || _ggbfc.Y != _bgg {
		_fba := _ab.NewRGBA(_dgde)
		_ga.BiLinear.Scale(_fba, _dgde, _abe, _abef.Bounds(), _ga.Over, nil)
		_abef = _fba
	}
	if _ddaf.X != _dga || _ddaf.Y != _bgg {
		_eeb := _ab.NewRGBA(_dgde)
		_ga.BiLinear.Scale(_eeb, _dgde, _abe, _abe.Bounds(), _ga.Over, nil)
		_abe = _eeb
	}
	_aab := _ab.NewRGBA(_dgde)
	_ga.DrawMask(_aab, _dgde, _abe, _ab.Point{}, _abef, _ab.Point{}, _ga.Over)
	return _aab
}
func _gbaaeg(_fcg string, _bcae _ab.Image, _aecf int) error {
	_eacf, _gdfc := _gd.Create(_fcg)
	if _gdfc != nil {
		return _gdfc
	}
	defer _eacf.Close()
	return _fe.Encode(_eacf, _bcae, &_fe.Options{Quality: _aecf})
}
func _dad(_abaa *_gb.Image, _ddd _ea.Color) _ab.Image {
	_eaga, _gaad := int(_abaa.Width), int(_abaa.Height)
	_acg := _ab.NewRGBA(_ab.Rect(0, 0, _eaga, _gaad))
	for _ggfd := 0; _ggfd < _gaad; _ggfd++ {
		for _dbf := 0; _dbf < _eaga; _dbf++ {
			_gfge, _aade := _abaa.ColorAt(_dbf, _ggfd)
			if _aade != nil {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _dbf, _ggfd)
				continue
			}
			_edb, _fbe, _gcgf, _ := _gfge.RGBA()
			var _bgdd _ea.Color
			if _edb+_fbe+_gcgf == 0 {
				_bgdd = _ea.Transparent
			} else {
				_bgdd = _ddd
			}
			_acg.Set(_dbf, _ggfd, _bgdd)
		}
	}
	return _acg
}
func _abec(_bef _fd.PdfObject, _ebec _ea.Color) (_ab.Image, error) {
	_cca, _fcb := _fd.GetStream(_bef)
	if !_fcb {
		return nil, nil
	}
	_bde, _fdg := _gb.NewXObjectImageFromStream(_cca)
	if _fdg != nil {
		return nil, _fdg
	}
	_fdee, _fdg := _bde.ToImage()
	if _fdg != nil {
		return nil, _fdg
	}
	return _dad(_fdee, _ebec), nil
}
func _aga(_dgb, _eced, _afg float64) _ad.BoundingBox {
	return _ad.Path{Points: []_ad.Point{_ad.NewPoint(0, 0).Rotate(_afg), _ad.NewPoint(_dgb, 0).Rotate(_afg), _ad.NewPoint(0, _eced).Rotate(_afg), _ad.NewPoint(_dgb, _eced).Rotate(_afg)}}.GetBoundingBox()
}

// NewImageDevice returns a new image device.
func NewImageDevice() *ImageDevice {
	return &ImageDevice{}
}
func (_gbg renderer) renderPage(_def _deb.Context, _ddf *_gb.PdfPage, _abf _aba.Matrix) error {
	_dfb, _af := _ddf.GetAllContentStreams()
	if _af != nil {
		return _af
	}
	if _eb := _abf; !_eb.Identity() {
		_dfb = _de.Sprintf("%\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066 \u0025\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066\u0020\u0063\u006d\u0020%\u0073", _eb[0], _eb[1], _eb[3], _eb[4], _eb[6], _eb[7], _dfb)
	}
	_def.Translate(0, float64(_def.Height()))
	_def.Scale(1, -1)
	_def.Push()
	_def.SetRGBA(1, 1, 1, 1)
	_def.DrawRectangle(0, 0, float64(_def.Width()), float64(_def.Height()))
	_def.Fill()
	_def.Pop()
	_def.SetLineWidth(1.0)
	_def.SetRGBA(0, 0, 0, 1)
	return _gbg.renderContentStream(_def, _dfb, _ddf.Resources)
}

// ImageDevice is used to render PDF pages to image targets.
type ImageDevice struct {
	renderer

	// OutputWidth represents the width of the rendered images in pixels.
	// The heights of the output images are calculated based on the selected
	// width and the original height of each rendered page.
	OutputWidth int
}

func _edac(_eag *_gb.Image, _gfb _ea.Color) _ab.Image {
	_fcgg, _acf := int(_eag.Width), int(_eag.Height)
	_dccg := _ab.NewRGBA(_ab.Rect(0, 0, _fcgg, _acf))
	for _gcba := 0; _gcba < _acf; _gcba++ {
		for _cfeg := 0; _cfeg < _fcgg; _cfeg++ {
			_geb, _eff := _eag.ColorAt(_cfeg, _gcba)
			if _eff != nil {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _cfeg, _gcba)
				continue
			}
			_dba, _dab, _dfd, _ := _geb.RGBA()
			var _ecbd _ea.Color
			if _dba+_dab+_dfd == 0 {
				_ecbd = _gfb
			} else {
				_ecbd = _ea.Transparent
			}
			_dccg.Set(_cfeg, _gcba, _ecbd)
		}
	}
	return _dccg
}

type renderer struct{ _gda float64 }

func _gfg(_cccd _fd.PdfObject, _ggg _ea.Color) (_ab.Image, error) {
	_cfd, _becg := _fd.GetStream(_cccd)
	if !_becg {
		return nil, nil
	}
	_bdbg, _ddc := _gb.NewXObjectImageFromStream(_cfd)
	if _ddc != nil {
		return nil, _ddc
	}
	_bbdb, _ddc := _bdbg.ToImage()
	if _ddc != nil {
		return nil, _ddc
	}
	return _edac(_bbdb, _ggg), nil
}
func (_fa renderer) renderContentStream(_dbe _deb.Context, _ecb string, _gg *_gb.PdfPageResources) error {
	_bed, _bg := _dd.NewContentStreamParser(_ecb).Parse()
	if _bg != nil {
		return _bg
	}
	_cf := _dbe.TextState()
	_cf.GlobalScale = _fa._gda
	_adc := map[string]*_deb.TextFont{}
	_afe := _df.NewFinder(&_df.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066", "\u002e\u0074\u0074\u0063"}})
	_fgg := _dd.NewContentStreamProcessor(*_bed)
	_fgg.AddHandler(_dd.HandlerConditionEnumAllOperands, "", func(_fec *_dd.ContentStreamOperation, _gdd _dd.GraphicsState, _fc *_gb.PdfPageResources) error {
		_b.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0025\u0073", _fec.Operand)
		switch _fec.Operand {
		case "\u0071":
			_dbe.Push()
		case "\u0051":
			_dbe.Pop()
			_cf = _dbe.TextState()
		case "\u0063\u006d":
			if len(_fec.Params) != 6 {
				return _dac
			}
			_ge, _dbc := _fd.GetNumbersAsFloat(_fec.Params)
			if _dbc != nil {
				return _dbc
			}
			_aac := _aba.NewMatrix(_ge[0], _ge[1], _ge[2], _ge[3], _ge[4], _ge[5])
			_b.Log.Debug("\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074a\u0074\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078\u003a \u0025\u002b\u0076", _aac)
			_dbe.SetMatrix(_dbe.Matrix().Mult(_aac))
		case "\u0077":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_fcd, _ced := _fd.GetNumbersAsFloat(_fec.Params)
			if _ced != nil {
				return _ced
			}
			_dbe.SetLineWidth(_fcd[0])
		case "\u004a":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_dda, _ecd := _fd.GetIntVal(_fec.Params[0])
			if !_ecd {
				return _bec
			}
			switch _dda {
			case 0:
				_dbe.SetLineCap(_deb.LineCapButt)
			case 1:
				_dbe.SetLineCap(_deb.LineCapRound)
			case 2:
				_dbe.SetLineCap(_deb.LineCapSquare)
			default:
				_b.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u006ee\u0020\u0063\u0061\u0070\u0020\u0073\u0074\u0079\u006c\u0065:\u0020\u0025\u0064", _dda)
				return _dac
			}
		case "\u006a":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_acb, _abd := _fd.GetIntVal(_fec.Params[0])
			if !_abd {
				return _bec
			}
			switch _acb {
			case 0:
				_dbe.SetLineJoin(_deb.LineJoinBevel)
			case 1:
				_dbe.SetLineJoin(_deb.LineJoinRound)
			case 2:
				_dbe.SetLineJoin(_deb.LineJoinBevel)
			default:
				_b.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u006e\u0065\u0020\u006a\u006f\u0069\u006e \u0073\u0074\u0079l\u0065:\u0020\u0025\u0064", _acb)
				return _dac
			}
		case "\u004d":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_adb, _gf := _fd.GetNumbersAsFloat(_fec.Params)
			if _gf != nil {
				return _gf
			}
			_ = _adb
			_b.Log.Debug("\u004di\u0074\u0065\u0072\u0020l\u0069\u006d\u0069\u0074\u0020n\u006ft\u0020s\u0075\u0070\u0070\u006f\u0072\u0074\u0065d")
		case "\u0064":
			if len(_fec.Params) != 2 {
				return _dac
			}
			_cb, _fde := _fd.GetArray(_fec.Params[0])
			if !_fde {
				return _bec
			}
			_agb, _fde := _fd.GetIntVal(_fec.Params[1])
			if !_fde {
				return _bec
			}
			_dcb, _cce := _fd.GetNumbersAsFloat(_cb.Elements())
			if _cce != nil {
				return _cce
			}
			_dbe.SetDash(_dcb...)
			_ = _agb
			_b.Log.Debug("\u004c\u0069n\u0065\u0020\u0064\u0061\u0073\u0068\u0020\u0070\u0068\u0061\u0073\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006frt\u0065\u0064")
		case "\u0072\u0069":
			_b.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020i\u006e\u0074\u0065\u006e\u0074\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0069":
			_b.Log.Debug("\u0046\u006c\u0061\u0074\u006e\u0065\u0073\u0073\u0020\u0074\u006f\u006c\u0065\u0072\u0061n\u0063e\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0067\u0073":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_ggb, _ee := _fd.GetName(_fec.Params[0])
			if !_ee {
				return _bec
			}
			if _ggb == nil {
				return _dac
			}
			_cfb, _ee := _fc.GetExtGState(*_ggb)
			if !_ee {
				_b.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074 \u0066i\u006ed\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u003a\u0020\u0025\u0073", *_ggb)
				return _g.New("\u0072e\u0073o\u0075\u0072\u0063\u0065\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064")
			}
			_eaa, _ee := _fd.GetDict(_cfb)
			if !_ee {
				_b.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020c\u006f\u0075\u006c\u0064 ge\u0074 g\u0072\u0061\u0070\u0068\u0069\u0063\u0073 s\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074")
				return _bec
			}
			_b.Log.Debug("G\u0053\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _eaa.String())
		case "\u006d":
			if len(_fec.Params) != 2 {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _dac)
				return nil
			}
			_eca, _egea := _fd.GetNumbersAsFloat(_fec.Params)
			if _egea != nil {
				return _egea
			}
			_b.Log.Debug("M\u006f\u0076\u0065\u0020\u0074\u006f\u003a\u0020\u0025\u0076", _eca)
			_dbe.NewSubPath()
			_dbe.MoveTo(_eca[0], _eca[1])
		case "\u006c":
			if len(_fec.Params) != 2 {
				_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _dac)
				return nil
			}
			_fcf, _gbe := _fd.GetNumbersAsFloat(_fec.Params)
			if _gbe != nil {
				return _gbe
			}
			_dbe.LineTo(_fcf[0], _fcf[1])
		case "\u0063":
			if len(_fec.Params) != 6 {
				return _dac
			}
			_agd, _eac := _fd.GetNumbersAsFloat(_fec.Params)
			if _eac != nil {
				return _eac
			}
			_b.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _agd)
			_dbe.CubicTo(_agd[0], _agd[1], _agd[2], _agd[3], _agd[4], _agd[5])
		case "\u0076", "\u0079":
			if len(_fec.Params) != 4 {
				return _dac
			}
			_fdb, _gaa := _fd.GetNumbersAsFloat(_fec.Params)
			if _gaa != nil {
				return _gaa
			}
			_b.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _fdb)
			_dbe.QuadraticTo(_fdb[0], _fdb[1], _fdb[2], _fdb[3])
		case "\u0068":
			_dbe.ClosePath()
			_dbe.NewSubPath()
		case "\u0072\u0065":
			if len(_fec.Params) != 4 {
				return _dac
			}
			_ed, _gdc := _fd.GetNumbersAsFloat(_fec.Params)
			if _gdc != nil {
				return _gdc
			}
			_dbe.DrawRectangle(_ed[0], _ed[1], _ed[2], _ed[3])
			_dbe.NewSubPath()
		case "\u0053":
			_fecb, _ece := _gdd.ColorspaceStroking.ColorToRGB(_gdd.ColorStroking)
			if _ece != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ece)
				return _ece
			}
			_cfbf, _fed := _fecb.(*_gb.PdfColorDeviceRGB)
			if !_fed {
				_b.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _ece
			}
			_dbe.SetRGBA(_cfbf.R(), _cfbf.G(), _cfbf.B(), 1)
			_dbe.Stroke()
		case "\u0073":
			_ff, _ffa := _gdd.ColorspaceStroking.ColorToRGB(_gdd.ColorStroking)
			if _ffa != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ffa)
				return _ffa
			}
			_fab, _bcf := _ff.(*_gb.PdfColorDeviceRGB)
			if !_bcf {
				_b.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _ffa
			}
			_dbe.ClosePath()
			_dbe.NewSubPath()
			_dbe.SetRGBA(_fab.R(), _fab.G(), _fab.B(), 1)
			_dbe.Stroke()
		case "\u0066", "\u0046":
			_eeg, _fggc := _gdd.ColorspaceNonStroking.ColorToRGB(_gdd.ColorNonStroking)
			if _fggc != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fggc)
				return _fggc
			}
			_bd, _fca := _eeg.(*_gb.PdfColorDeviceRGB)
			if !_fca {
				_b.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _fggc
			}
			_dbe.SetRGBA(_bd.R(), _bd.G(), _bd.B(), 1)
			_dbe.SetFillRule(_deb.FillRuleWinding)
			_dbe.Fill()
		case "\u0066\u002a":
			_cd, _aec := _gdd.ColorspaceNonStroking.ColorToRGB(_gdd.ColorNonStroking)
			if _aec != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _aec)
				return _aec
			}
			_aaa, _adfg := _cd.(*_gb.PdfColorDeviceRGB)
			if !_adfg {
				_b.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _aec
			}
			_dbe.SetRGBA(_aaa.R(), _aaa.G(), _aaa.B(), 1)
			_dbe.SetFillRule(_deb.FillRuleEvenOdd)
			_dbe.Fill()
		case "\u0042":
			_gac, _cea := _gdd.ColorspaceNonStroking.ColorToRGB(_gdd.ColorNonStroking)
			if _cea != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cea)
				return _cea
			}
			_egc := _gac.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_egc.R(), _egc.G(), _egc.B(), 1)
			_dbe.SetFillRule(_deb.FillRuleWinding)
			_dbe.FillPreserve()
			_gac, _cea = _gdd.ColorspaceStroking.ColorToRGB(_gdd.ColorStroking)
			if _cea != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cea)
				return _cea
			}
			_egc = _gac.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_egc.R(), _egc.G(), _egc.B(), 1)
			_dbe.Stroke()
		case "\u0042\u002a":
			_gbb, _ggd := _gdd.ColorspaceNonStroking.ColorToRGB(_gdd.ColorNonStroking)
			if _ggd != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ggd)
				return _ggd
			}
			_fb := _gbb.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_fb.R(), _fb.G(), _fb.B(), 1)
			_dbe.SetFillRule(_deb.FillRuleEvenOdd)
			_dbe.FillPreserve()
			_gbb, _ggd = _gdd.ColorspaceStroking.ColorToRGB(_gdd.ColorStroking)
			if _ggd != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ggd)
				return _ggd
			}
			_fb = _gbb.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_fb.R(), _fb.G(), _fb.B(), 1)
			_dbe.Stroke()
		case "\u0062":
			_eed, _cbd := _gdd.ColorspaceNonStroking.ColorToRGB(_gdd.ColorNonStroking)
			if _cbd != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbd)
				return _cbd
			}
			_beab := _eed.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_beab.R(), _beab.G(), _beab.B(), 1)
			_dbe.ClosePath()
			_dbe.NewSubPath()
			_dbe.SetFillRule(_deb.FillRuleWinding)
			_dbe.FillPreserve()
			_eed, _cbd = _gdd.ColorspaceStroking.ColorToRGB(_gdd.ColorStroking)
			if _cbd != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbd)
				return _cbd
			}
			_beab = _eed.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_beab.R(), _beab.G(), _beab.B(), 1)
			_dbe.Stroke()
		case "\u0062\u002a":
			_dbe.ClosePath()
			_ega, _dg := _gdd.ColorspaceNonStroking.ColorToRGB(_gdd.ColorNonStroking)
			if _dg != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dg)
				return _dg
			}
			_bfe := _ega.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_bfe.R(), _bfe.G(), _bfe.B(), 1)
			_dbe.NewSubPath()
			_dbe.SetFillRule(_deb.FillRuleEvenOdd)
			_dbe.FillPreserve()
			_ega, _dg = _gdd.ColorspaceStroking.ColorToRGB(_gdd.ColorStroking)
			if _dg != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dg)
				return _dg
			}
			_bfe = _ega.(*_gb.PdfColorDeviceRGB)
			_dbe.SetRGBA(_bfe.R(), _bfe.G(), _bfe.B(), 1)
			_dbe.Stroke()
		case "\u006e":
			_dbe.ClearPath()
		case "\u0057":
			_dbe.SetFillRule(_deb.FillRuleWinding)
			_dbe.ClipPreserve()
		case "\u0057\u002a":
			_dbe.SetFillRule(_deb.FillRuleEvenOdd)
			_dbe.ClipPreserve()
		case "\u0072\u0067":
			_fad, _gbaa := _gdd.ColorNonStroking.(*_gb.PdfColorDeviceRGB)
			if !_gbaa {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorNonStroking)
				return nil
			}
			_dbe.SetFillRGBA(_fad.R(), _fad.G(), _fad.B(), 1)
		case "\u0052\u0047":
			_dgd, _ggf := _gdd.ColorStroking.(*_gb.PdfColorDeviceRGB)
			if !_ggf {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorStroking)
				return nil
			}
			_dbe.SetStrokeRGBA(_dgd.R(), _dgd.G(), _dgd.B(), 1)
		case "\u006b":
			_ebab, _bdg := _gdd.ColorNonStroking.(*_gb.PdfColorDeviceCMYK)
			if !_bdg {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorNonStroking)
				return nil
			}
			_dcbc, _bb := _gdd.ColorspaceNonStroking.ColorToRGB(_ebab)
			if _bb != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorNonStroking)
				return nil
			}
			_cbe, _bdg := _dcbc.(*_gb.PdfColorDeviceRGB)
			if !_bdg {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dcbc)
				return nil
			}
			_dbe.SetFillRGBA(_cbe.R(), _cbe.G(), _cbe.B(), 1)
		case "\u004b":
			_ccg, _dcc := _gdd.ColorStroking.(*_gb.PdfColorDeviceCMYK)
			if !_dcc {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorStroking)
				return nil
			}
			_bgd, _adfc := _gdd.ColorspaceStroking.ColorToRGB(_ccg)
			if _adfc != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorStroking)
				return nil
			}
			_dfc, _dcc := _bgd.(*_gb.PdfColorDeviceRGB)
			if !_dcc {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bgd)
				return nil
			}
			_dbe.SetStrokeRGBA(_dfc.R(), _dfc.G(), _dfc.B(), 1)
		case "\u0067":
			_ccee, _aaf := _gdd.ColorNonStroking.(*_gb.PdfColorDeviceGray)
			if !_aaf {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorNonStroking)
				return nil
			}
			_bbe, _gdb := _gdd.ColorspaceNonStroking.ColorToRGB(_ccee)
			if _gdb != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorNonStroking)
				return nil
			}
			_dbd, _aaf := _bbe.(*_gb.PdfColorDeviceRGB)
			if !_aaf {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bbe)
				return nil
			}
			_dbe.SetFillRGBA(_dbd.R(), _dbd.G(), _dbd.B(), 1)
		case "\u0047":
			_cdd, _adbf := _gdd.ColorStroking.(*_gb.PdfColorDeviceGray)
			if !_adbf {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorStroking)
				return nil
			}
			_fea, _ccb := _gdd.ColorspaceStroking.ColorToRGB(_cdd)
			if _ccb != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorStroking)
				return nil
			}
			_cgd, _adbf := _fea.(*_gb.PdfColorDeviceRGB)
			if !_adbf {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fea)
				return nil
			}
			_dbe.SetStrokeRGBA(_cgd.R(), _cgd.G(), _cgd.B(), 1)
		case "\u0063\u0073", "\u0073\u0063", "\u0073\u0063\u006e":
			_cba, _fcfc := _gdd.ColorspaceNonStroking.ColorToRGB(_gdd.ColorNonStroking)
			if _fcfc != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorNonStroking)
				return nil
			}
			_faa, _dec := _cba.(*_gb.PdfColorDeviceRGB)
			if !_dec {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cba)
				return nil
			}
			_dbe.SetFillRGBA(_faa.R(), _faa.G(), _faa.B(), 1)
		case "\u0043\u0053", "\u0053\u0043", "\u0053\u0043\u004e":
			_dce, _bfec := _gdd.ColorspaceStroking.ColorToRGB(_gdd.ColorStroking)
			if _bfec != nil {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gdd.ColorStroking)
				return nil
			}
			_ged, _abda := _dce.(*_gb.PdfColorDeviceRGB)
			if !_abda {
				_b.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dce)
				return nil
			}
			_dbe.SetStrokeRGBA(_ged.R(), _ged.G(), _ged.B(), 1)
		case "\u0044\u006f":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_gbaae, _gdca := _fd.GetName(_fec.Params[0])
			if !_gdca {
				return _bec
			}
			_, _gbf := _fc.GetXObjectByName(*_gbaae)
			switch _gbf {
			case _gb.XObjectTypeImage:
				_b.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067e\u003a\u0020\u0025\u0073", _gbaae.String())
				_ba, _ceg := _fc.GetXObjectImageByName(*_gbaae)
				if _ceg != nil {
					return _ceg
				}
				_bge, _ceg := _ba.ToImage()
				if _ceg != nil {
					return _ceg
				}
				if _ecad := _ba.ColorSpace; _ecad != nil {
					var _bdgb bool
					switch _ecad.(type) {
					case *_gb.PdfColorspaceSpecialIndexed:
						_bdgb = true
					}
					if _bdgb {
						if _defg, _egg := _ecad.ImageToRGB(*_bge); _egg != nil {
							_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006fnv\u0065r\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0052G\u0042\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020i\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
						} else {
							_bge = &_defg
						}
					}
				}
				_ggbf := _dbe.FillPattern().ColorAt(0, 0)
				var _fgc _ab.Image
				if _ba.Mask != nil {
					if _fgc, _ceg = _gfg(_ba.Mask, _ggbf); _ceg != nil {
						_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				} else if _ba.SMask != nil {
					if _fgc, _ceg = _abec(_ba.SMask, _ggbf); _ceg != nil {
						_b.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				}
				var _fgcg _ab.Image
				if _ffaa, _ := _fd.GetBoolVal(_ba.ImageMask); _ffaa {
					_fgcg = _edac(_bge, _ggbf)
				} else {
					_fgcg, _ceg = _bge.ToGoImage()
					if _ceg != nil {
						return _ceg
					}
				}
				if _fgc != nil {
					_fgcg = _aea(_fgcg, _fgc)
				}
				_ebc := _fgcg.Bounds()
				_dbe.Push()
				_dbe.Scale(1.0/float64(_ebc.Dx()), -1.0/float64(_ebc.Dy()))
				_dbe.DrawImageAnchored(_fgcg, 0, 0, 0, 1)
				_dbe.Pop()
			case _gb.XObjectTypeForm:
				_b.Log.Debug("\u0058\u004fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u0073", _gbaae.String())
				_cfe, _bgc := _fc.GetXObjectFormByName(*_gbaae)
				if _bgc != nil {
					return _bgc
				}
				_cbf, _bgc := _cfe.GetContentStream()
				if _bgc != nil {
					return _bgc
				}
				_dfg := _cfe.Resources
				if _dfg == nil {
					_dfg = _fc
				}
				_dbe.Push()
				if _cfe.Matrix != nil {
					_bagf, _dca := _fd.GetArray(_cfe.Matrix)
					if !_dca {
						return _bec
					}
					_fgf, _bagg := _fd.GetNumbersAsFloat(_bagf.Elements())
					if _bagg != nil {
						return _bagg
					}
					if len(_fgf) != 6 {
						return _dac
					}
					_bbc := _aba.NewMatrix(_fgf[0], _fgf[1], _fgf[2], _fgf[3], _fgf[4], _fgf[5])
					_dbe.SetMatrix(_dbe.Matrix().Mult(_bbc))
				}
				if _cfe.BBox != nil {
					_ceeg, _gcd := _fd.GetArray(_cfe.BBox)
					if !_gcd {
						return _bec
					}
					_bba, _acbb := _fd.GetNumbersAsFloat(_ceeg.Elements())
					if _acbb != nil {
						return _acbb
					}
					if len(_bba) != 4 {
						_b.Log.Debug("\u004c\u0065\u006e\u0020\u003d\u0020\u0025\u0064", len(_bba))
						return _dac
					}
					_dbe.DrawRectangle(_bba[0], _bba[1], _bba[2]-_bba[0], _bba[3]-_bba[1])
					_dbe.SetRGBA(1, 0, 0, 1)
					_dbe.Clip()
				} else {
					_b.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0052\u0065q\u0075\u0069\u0072e\u0064\u0020\u0042\u0042\u006f\u0078\u0020\u006d\u0069ss\u0069\u006e\u0067 \u006f\u006e \u0058\u004f\u0062\u006a\u0065\u0063t\u0020\u0046o\u0072\u006d")
				}
				_bgc = _fa.renderContentStream(_dbe, string(_cbf), _dfg)
				if _bgc != nil {
					return _bgc
				}
				_dbe.Pop()
			}
		case "\u0042\u0049":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_cab, _ebe := _fec.Params[0].(*_dd.ContentStreamInlineImage)
			if !_ebe {
				return nil
			}
			_dfa, _fbd := _cab.ToImage(_fc)
			if _fbd != nil {
				return _fbd
			}
			_bfg, _fbd := _dfa.ToGoImage()
			if _fbd != nil {
				return _fbd
			}
			_fdd := _bfg.Bounds()
			_dbe.Push()
			_dbe.Scale(1.0/float64(_fdd.Dx()), -1.0/float64(_fdd.Dy()))
			_dbe.DrawImageAnchored(_bfg, 0, 0, 0, 1)
			_dbe.Pop()
		case "\u0042\u0054":
			_cf.Reset()
		case "\u0045\u0054":
			_cf.Reset()
		case "\u0054\u0072":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_geg, _edd := _fd.GetNumberAsFloat(_fec.Params[0])
			if _edd != nil {
				return _edd
			}
			_cf.Tr = _deb.TextRenderingMode(_geg)
		case "\u0054\u004c":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_gcb, _ceeb := _fd.GetNumberAsFloat(_fec.Params[0])
			if _ceeb != nil {
				return _ceeb
			}
			_cf.Tl = _gcb
		case "\u0054\u0063":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_defd, _bedb := _fd.GetNumberAsFloat(_fec.Params[0])
			if _bedb != nil {
				return _bedb
			}
			_b.Log.Debug("\u0054\u0063\u003a\u0020\u0025\u0076", _defd)
			_cf.Tc = _defd
		case "\u0054\u0077":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_ecac, _gfa := _fd.GetNumberAsFloat(_fec.Params[0])
			if _gfa != nil {
				return _gfa
			}
			_b.Log.Debug("\u0054\u0077\u003a\u0020\u0025\u0076", _ecac)
			_cf.Tw = _ecac
		case "\u0054\u007a":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_gcbf, _daa := _fd.GetNumberAsFloat(_fec.Params[0])
			if _daa != nil {
				return _daa
			}
			_cf.Th = _gcbf
		case "\u0054\u0073":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_aece, _ccga := _fd.GetNumberAsFloat(_fec.Params[0])
			if _ccga != nil {
				return _ccga
			}
			_cf.Ts = _aece
		case "\u0054\u0064":
			if len(_fec.Params) != 2 {
				return _dac
			}
			_cac, _fdc := _fd.GetNumbersAsFloat(_fec.Params)
			if _fdc != nil {
				return _fdc
			}
			_b.Log.Debug("\u0054\u0064\u003a\u0020\u0025\u0076", _cac)
			_cf.ProcTd(_cac[0], _cac[1])
		case "\u0054\u0044":
			if len(_fec.Params) != 2 {
				return _dac
			}
			_gef, _baf := _fd.GetNumbersAsFloat(_fec.Params)
			if _baf != nil {
				return _baf
			}
			_b.Log.Debug("\u0054\u0044\u003a\u0020\u0025\u0076", _gef)
			_cf.ProcTD(_gef[0], _gef[1])
		case "\u0054\u002a":
			_cf.ProcTStar()
		case "\u0054\u006d":
			if len(_fec.Params) != 6 {
				return _dac
			}
			_bbab, _bac := _fd.GetNumbersAsFloat(_fec.Params)
			if _bac != nil {
				return _bac
			}
			_b.Log.Debug("\u0054\u0065x\u0074\u0020\u006da\u0074\u0072\u0069\u0078\u003a\u0020\u0025\u002b\u0076", _bbab)
			_cf.ProcTm(_bbab[0], _bbab[1], _bbab[2], _bbab[3], _bbab[4], _bbab[5])
		case "\u0027":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_ffd, _aeg := _fd.GetStringBytes(_fec.Params[0])
			if !_aeg {
				return _bec
			}
			_b.Log.Debug("\u0027\u0020\u0073t\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_ffd))
			_cf.ProcQ(_ffd, _dbe)
		case "\u0022":
			if len(_fec.Params) != 3 {
				return _dac
			}
			_cfa, _bdb := _fd.GetNumberAsFloat(_fec.Params[0])
			if _bdb != nil {
				return _bdb
			}
			_bbeb, _bdb := _fd.GetNumberAsFloat(_fec.Params[1])
			if _bdb != nil {
				return _bdb
			}
			_gcfd, _bca := _fd.GetStringBytes(_fec.Params[2])
			if !_bca {
				return _bec
			}
			_cf.ProcDQ(_gcfd, _cfa, _bbeb, _dbe)
		case "\u0054\u006a":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_bbag, _cddd := _fd.GetStringBytes(_fec.Params[0])
			if !_cddd {
				return _bec
			}
			_b.Log.Debug("\u0054j\u0020s\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0060\u0025\u0073\u0060", string(_bbag))
			_cf.ProcTj(_bbag, _dbe)
		case "\u0054\u004a":
			if len(_fec.Params) != 1 {
				return _dac
			}
			_bcb, _eda := _fd.GetArray(_fec.Params[0])
			if !_eda {
				_b.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _bcb)
				return _bec
			}
			_b.Log.Debug("\u0054\u004a\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u002b\u0076", _bcb)
			for _, _aef := range _bcb.Elements() {
				switch _bfa := _aef.(type) {
				case *_fd.PdfObjectString:
					if _bfa != nil {
						_cf.ProcTj(_bfa.Bytes(), _dbe)
					}
				case *_fd.PdfObjectFloat, *_fd.PdfObjectInteger:
					_dcg, _bbad := _fd.GetNumberAsFloat(_bfa)
					if _bbad == nil {
						_cf.Translate(-_dcg*0.001*_cf.Tf.Size*_cf.Th/100.0, 0)
					}
				}
			}
		case "\u0054\u0066":
			if len(_fec.Params) != 2 {
				return _dac
			}
			_b.Log.Debug("\u0025\u0023\u0076", _fec.Params)
			_defe, _adbd := _fd.GetName(_fec.Params[0])
			if !_adbd || _defe == nil {
				_b.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006e\u0061m\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _fec.Params[0])
				return _bec
			}
			_b.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073", _defe.String())
			_egf, _dgc := _fd.GetNumberAsFloat(_fec.Params[1])
			if _dgc != nil {
				_b.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0069z\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _fec.Params[1])
				return _bec
			}
			_b.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0076", _egf)
			_bbd, _dcaf := _fc.GetFontByName(*_defe)
			if !_dcaf {
				_b.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0025s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _defe.String())
				return _g.New("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
			}
			_b.Log.Debug("\u0046\u006f\u006e\u0074\u003a\u0020\u0025\u0054", _bbd)
			_geda, _adbd := _fd.GetDict(_bbd)
			if !_adbd {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0067e\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074")
				return _bec
			}
			_ecc, _dgc := _gb.NewPdfFontFromPdfObject(_geda)
			if _dgc != nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
				return _dgc
			}
			_cfeb := _ecc.BaseFont()
			if _cfeb == "" {
				_cfeb = _defe.String()
			}
			_eggd, _adbd := _adc[_cfeb]
			if !_adbd {
				_eggd, _dgc = _deb.NewTextFont(_ecc, _egf)
				if _dgc != nil {
					_b.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgc)
				}
			}
			if _eggd == nil {
				if len(_cfeb) > 7 && _cfeb[6] == '+' {
					_cfeb = _cfeb[7:]
				}
				_dcfa := []string{_cfeb, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
				for _, _bcg := range _dcfa {
					_b.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _bcg)
					if _eggd, _adbd = _adc[_bcg]; _adbd {
						break
					}
					_ccgg := _afe.Match(_bcg)
					if _ccgg == nil {
						_b.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _bcg)
						continue
					}
					_eggd, _dgc = _deb.NewTextFontFromPath(_ccgg.Filename, _egf)
					if _dgc != nil {
						_b.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _ccgg.Filename)
						continue
					}
					_b.Log.Debug("\u0053\u0075\u0062\u0073\u0074\u0069t\u0075\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073 \u0077\u0069\u0074\u0068\u0020\u0025\u0073 \u0028\u0025\u0073\u0029", _cfeb, _ccgg.Name, _ccgg.Filename)
					_adc[_bcg] = _eggd
					break
				}
			}
			if _eggd == nil {
				_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020n\u006f\u0074\u0020\u0066\u0069\u006ed\u0020\u0061\u006e\u0079\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065 \u0066\u006f\u006e\u0074")
				return _g.New("\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0066\u0069\u006e\u0064\u0020a\u006ey\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u006e\u0074")
			}
			_cf.ProcTf(_eggd.WithSize(_egf, _ecc))
		case "\u0042\u004d\u0043", "\u0042\u0044\u0043":
		case "\u0045\u004d\u0043":
		default:
			_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u006f\u0070\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u0073", _fec.Operand)
		}
		return nil
	})
	_bg = _fgg.Process(_gg)
	if _bg != nil {
		return _bg
	}
	return nil
}

// Render converts the specified PDF page into an image and returns the result.
func (_gcg *ImageDevice) Render(page *_gb.PdfPage) (_ab.Image, error) {
	_bf, _c := page.GetMediaBox()
	if _c != nil {
		return nil, _c
	}
	_bf.Normalize()
	_gba := page.CropBox
	var _cc, _da float64
	if _gba != nil {
		_gba.Normalize()
		_cc, _da = _gba.Width(), _gba.Height()
	}
	_ce := page.Rotate
	_gcf, _cg, _cge, _aa := _bf.Llx, _bf.Lly, _bf.Width(), _bf.Height()
	_cga := _aba.IdentityMatrix()
	if _ce != nil && *_ce%360 != 0 && *_ce%90 == 0 {
		_ege := -float64(*_ce)
		_be := _aga(_cge, _aa, _ege)
		_cga = _cga.Translate((_be.Width-_cge)/2+_cge/2, (_be.Height-_aa)/2+_aa/2).Rotate(_ege*_e.Pi/180).Translate(-_cge/2, -_aa/2)
		_cge, _aa = _be.Width, _be.Height
		if _gba != nil {
			_ag := _aga(_cc, _da, _ege)
			_cc, _da = _ag.Width, _ag.Height
		}
	}
	if _gcf != 0 || _cg != 0 {
		_cga = _cga.Translate(-_gcf, -_cg)
	}
	_gcg._gda = 1.0
	if _gcg.OutputWidth != 0 {
		_aad := _cge
		if _gba != nil {
			_aad = _cc
		}
		_gcg._gda = float64(_gcg.OutputWidth) / _aad
		_cge, _aa, _cc, _da = _cge*_gcg._gda, _aa*_gcg._gda, _cc*_gcg._gda, _da*_gcg._gda
		_cga = _aba.ScaleMatrix(_gcg._gda, _gcg._gda).Mult(_cga)
	}
	_fg := _gc.NewContext(int(_cge), int(_aa))
	if _ac := _gcg.renderPage(_fg, page, _cga); _ac != nil {
		return nil, _ac
	}
	_ca := _fg.Image()
	if _gba != nil {
		_ec, _bc := (_gba.Llx-_gcf)*_gcg._gda, (_gba.Lly-_cg)*_gcg._gda
		_cee := _ab.Rect(0, 0, int(_cc), int(_da))
		_ef := _ab.Pt(int(_ec), int(_aa-_bc-_da))
		_gdf := _ab.NewRGBA(_cee)
		_dc.Draw(_gdf, _cee, _ca, _ef, _dc.Src)
		_ca = _gdf
	}
	return _ca, nil
}

// RenderToPath converts the specified PDF page into an image and saves the
// result at the specified location.
func (_ae *ImageDevice) RenderToPath(page *_gb.PdfPage, outputPath string) error {
	_fga, _dcf := _ae.Render(page)
	if _dcf != nil {
		return _dcf
	}
	_bea := _d.ToLower(_a.Ext(outputPath))
	if _bea == "" {
		return _g.New("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065cog\u006eiz\u0065\u0020\u006f\u0075\u0074\u0070\u0075t \u0066\u0069\u006c\u0065\u0020\u0074\u0079p\u0065")
	}
	switch _bea {
	case "\u002e\u0070\u006e\u0067":
		return _cccb(outputPath, _fga)
	case "\u002e\u006a\u0070\u0067", "\u002e\u006a\u0070e\u0067":
		return _gbaaeg(outputPath, _fga, 100)
	}
	return _de.Errorf("\u0075\u006e\u0072\u0065\u0063\u006fg\u006e\u0069\u007a\u0065\u0064\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020f\u0069\u006c\u0065\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _bea)
}
