package render

import (
	_b "errors"
	_cge "fmt"
	_db "image"
	_ea "image/color"
	_ce "image/draw"
	_ced "image/jpeg"
	_cg "image/png"
	_d "math"
	_cb "os"
	_f "path/filepath"
	_e "strings"

	_dee "bitbucket.org/shenghui0779/gopdf/annotator"
	_dd "bitbucket.org/shenghui0779/gopdf/common"
	_eg "bitbucket.org/shenghui0779/gopdf/contentstream"
	_be "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_de "bitbucket.org/shenghui0779/gopdf/core"
	_ba "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ed "bitbucket.org/shenghui0779/gopdf/model"
	_a "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_g "bitbucket.org/shenghui0779/gopdf/render/internal/context/imagerender"
	_cbf "github.com/adrg/sysfont"
	_eb "golang.org/x/image/draw"
)

// RenderWithOpts converts the specified PDF page into an image, optionally flattens annotations and returns the result.
func (_ff *ImageDevice) RenderWithOpts(page *_ed.PdfPage, skipFlattening bool) (_db.Image, error) {
	_af, _cf := page.GetMediaBox()
	if _cf != nil {
		return nil, _cf
	}
	_af.Normalize()
	_cc := page.CropBox
	var _cgg, _afc float64
	if _cc != nil {
		_cc.Normalize()
		_cgg, _afc = _cc.Width(), _cc.Height()
	}
	_fff := page.Rotate
	_ee, _cef, _da, _edc := _af.Llx, _af.Lly, _af.Width(), _af.Height()
	_edd := _ba.IdentityMatrix()
	if _fff != nil && *_fff%360 != 0 && *_fff%90 == 0 {
		_ffa := -float64(*_fff)
		_eaa := _fgef(_da, _edc, _ffa)
		_edd = _edd.Translate((_eaa.Width-_da)/2+_da/2, (_eaa.Height-_edc)/2+_edc/2).Rotate(_ffa*_d.Pi/180).Translate(-_da/2, -_edc/2)
		_da, _edc = _eaa.Width, _eaa.Height
		if _cc != nil {
			_eeb := _fgef(_cgg, _afc, _ffa)
			_cgg, _afc = _eeb.Width, _eeb.Height
		}
	}
	if _ee != 0 || _cef != 0 {
		_edd = _edd.Translate(-_ee, -_cef)
	}
	_ff._fee = 1.0
	if _ff.OutputWidth != 0 {
		_bac := _da
		if _cc != nil {
			_bac = _cgg
		}
		_ff._fee = float64(_ff.OutputWidth) / _bac
		_da, _edc, _cgg, _afc = _da*_ff._fee, _edc*_ff._fee, _cgg*_ff._fee, _afc*_ff._fee
		_edd = _ba.ScaleMatrix(_ff._fee, _ff._fee).Mult(_edd)
	}
	_gb := _g.NewContext(int(_da), int(_edc))
	if _fd := _ff.renderPage(_gb, page, _edd, skipFlattening); _fd != nil {
		return nil, _fd
	}
	_bed := _gb.Image()
	if _cc != nil {
		_ae, _fe := (_cc.Llx-_ee)*_ff._fee, (_cc.Lly-_cef)*_ff._fee
		_bef := _db.Rect(0, 0, int(_cgg), int(_afc))
		_fea := _db.Pt(int(_ae), int(_edc-_fe-_afc))
		_bf := _db.NewRGBA(_bef)
		_ce.Draw(_bf, _bef, _bed, _fea, _ce.Src)
		_bed = _bf
	}
	return _bed, nil
}

// NewImageDevice returns a new image device.
func NewImageDevice() *ImageDevice {
	return &ImageDevice{}
}

// ImageDevice is used to render PDF pages to image targets.
type ImageDevice struct {
	renderer

	// OutputWidth represents the width of the rendered images in pixels.
	// The heights of the output images are calculated based on the selected
	// width and the original height of each rendered page.
	OutputWidth int
}

func _deeg(_dfcf _a.Gradient, _agb *_ed.PdfFunctionType3, _dabd _ed.PdfColorspace, _aae []float64) (_a.Gradient, error) {
	var _bcec error
	for _gcbf := 0; _gcbf < len(_agb.Functions); _gcbf++ {
		if _dbca, _gdf := _agb.Functions[_gcbf].(*_ed.PdfFunctionType2); _gdf {
			_dfcf, _bcec = _eaag(_dfcf, _dbca, _dabd, _aae[_gcbf+1], _gcbf == 0)
			if _bcec != nil {
				return nil, _bcec
			}
		}
	}
	return _dfcf, nil
}
func (_gba renderer) processLinearShading(_fgbg _a.Context, _afee *_ed.PdfShading) (_a.Gradient, *_de.PdfObjectArray, error) {
	_fge := _afee.GetContext().(*_ed.PdfShadingType2)
	if len(_fge.Function) == 0 {
		return nil, nil, _b.New("\u006e\u006f\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074 \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0073\u006b\u0069\u0070\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	_ged, _cce := _fge.Coords.ToFloat64Array()
	if _cce != nil {
		return nil, nil, _b.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020\u0073h\u0061\u0064\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	_eac := _afee.ColorSpace
	_gge, _eaf := _fgbg.Matrix().Transform(_ged[0], _ged[1])
	_bada, _gbgd := _fgbg.Matrix().Transform(_ged[2], _ged[3])
	_cdd := _g.NewLinearGradient(_gge, _eaf, _bada, _gbgd)
	_bge := _de.MakeArrayFromFloats([]float64{0, 0, 1, 1})
	for _, _dce := range _ged {
		if _dce > 1 {
			_bge = _fge.Coords
			break
		}
	}
	if _fdf, _ebbf := _fge.Function[0].(*_ed.PdfFunctionType2); _ebbf {
		_cdd, _cce = _eaag(_cdd, _fdf, _eac, 1.0, true)
	} else if _ece, _cfg := _fge.Function[0].(*_ed.PdfFunctionType3); _cfg {
		_egce := append([]float64{0}, _ece.Bounds...)
		_egce = append(_egce, 1.0)
		_cdd, _cce = _deeg(_cdd, _ece, _eac, _egce)
	}
	return _cdd, _bge, _cce
}

// PdfShadingType defines PDF shading types.
// Source: PDF32000_2008.pdf. Chapter 8.7.4.5
type PdfShadingType int64

func (_gc renderer) renderContentStream(_gf _a.Context, _fb string, _ffaf *_ed.PdfPageResources) error {
	_dbb, _ccf := _eg.NewContentStreamParser(_fb).Parse()
	if _ccf != nil {
		return _ccf
	}
	_ge := _gf.TextState()
	_ge.GlobalScale = _gc._fee
	_gd := map[string]*_a.TextFont{}
	_bd := _cbf.NewFinder(&_cbf.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066", "\u002e\u0074\u0074\u0063"}})
	var _cfd *_eg.ContentStreamOperation
	_ccg := _eg.NewContentStreamProcessor(*_dbb)
	_ccg.AddHandler(_eg.HandlerConditionEnumAllOperands, "", func(_bc *_eg.ContentStreamOperation, _gcg _eg.GraphicsState, _egfe *_ed.PdfPageResources) error {
		_dd.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0025\u0073", _bc.Operand)
		switch _bc.Operand {
		case "\u0071":
			_gf.Push()
		case "\u0051":
			_gf.Pop()
			_ge = _gf.TextState()
		case "\u0063\u006d":
			if len(_bc.Params) != 6 {
				return _afg
			}
			_ef, _fec := _de.GetNumbersAsFloat(_bc.Params)
			if _fec != nil {
				return _fec
			}
			_dbc := _ba.NewMatrix(_ef[0], _ef[1], _ef[2], _ef[3], _ef[4], _ef[5])
			_dd.Log.Debug("\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074a\u0074\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078\u003a \u0025\u002b\u0076", _dbc)
			_gf.SetMatrix(_gf.Matrix().Mult(_dbc))
		case "\u0077":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_agg, _dc := _de.GetNumbersAsFloat(_bc.Params)
			if _dc != nil {
				return _dc
			}
			_gf.SetLineWidth(_agg[0])
		case "\u004a":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_ddb, _ccfg := _de.GetIntVal(_bc.Params[0])
			if !_ccfg {
				return _ag
			}
			switch _ddb {
			case 0:
				_gf.SetLineCap(_a.LineCapButt)
			case 1:
				_gf.SetLineCap(_a.LineCapRound)
			case 2:
				_gf.SetLineCap(_a.LineCapSquare)
			default:
				_dd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u006ee\u0020\u0063\u0061\u0070\u0020\u0073\u0074\u0079\u006c\u0065:\u0020\u0025\u0064", _ddb)
				return _afg
			}
		case "\u006a":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_cgf, _ab := _de.GetIntVal(_bc.Params[0])
			if !_ab {
				return _ag
			}
			switch _cgf {
			case 0:
				_gf.SetLineJoin(_a.LineJoinBevel)
			case 1:
				_gf.SetLineJoin(_a.LineJoinRound)
			case 2:
				_gf.SetLineJoin(_a.LineJoinBevel)
			default:
				_dd.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u006e\u0065\u0020\u006a\u006f\u0069\u006e \u0073\u0074\u0079l\u0065:\u0020\u0025\u0064", _cgf)
				return _afg
			}
		case "\u004d":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_fecc, _gag := _de.GetNumbersAsFloat(_bc.Params)
			if _gag != nil {
				return _gag
			}
			_ = _fecc
			_dd.Log.Debug("\u004di\u0074\u0065\u0072\u0020l\u0069\u006d\u0069\u0074\u0020n\u006ft\u0020s\u0075\u0070\u0070\u006f\u0072\u0074\u0065d")
		case "\u0064":
			if len(_bc.Params) != 2 {
				return _afg
			}
			_dff, _aa := _de.GetArray(_bc.Params[0])
			if !_aa {
				return _ag
			}
			_ddg, _aa := _de.GetIntVal(_bc.Params[1])
			if !_aa {
				_, _gea := _de.GetFloatVal(_bc.Params[1])
				if !_gea {
					return _ag
				}
			}
			_fdce, _bb := _de.GetNumbersAsFloat(_dff.Elements())
			if _bb != nil {
				return _bb
			}
			_gf.SetDash(_fdce...)
			_ = _ddg
			_dd.Log.Debug("\u004c\u0069n\u0065\u0020\u0064\u0061\u0073\u0068\u0020\u0070\u0068\u0061\u0073\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006frt\u0065\u0064")
		case "\u0072\u0069":
			_dd.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020i\u006e\u0074\u0065\u006e\u0074\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0069":
			_dd.Log.Debug("\u0046\u006c\u0061\u0074\u006e\u0065\u0073\u0073\u0020\u0074\u006f\u006c\u0065\u0072\u0061n\u0063e\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0067\u0073":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_ggd, _dba := _de.GetName(_bc.Params[0])
			if !_dba {
				return _ag
			}
			if _ggd == nil {
				return _afg
			}
			_gcge, _dba := _egfe.GetExtGState(*_ggd)
			if !_dba {
				_dd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074 \u0066i\u006ed\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u003a\u0020\u0025\u0073", *_ggd)
				return _b.New("\u0072e\u0073o\u0075\u0072\u0063\u0065\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064")
			}
			_aeb, _dba := _de.GetDict(_gcge)
			if !_dba {
				_dd.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020c\u006f\u0075\u006c\u0064 ge\u0074 g\u0072\u0061\u0070\u0068\u0069\u0063\u0073 s\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074")
				return _ag
			}
			_dd.Log.Debug("G\u0053\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _aeb.String())
		case "\u006d":
			if len(_bc.Params) != 2 {
				_dd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _afg)
				return nil
			}
			_gfd, _ac := _de.GetNumbersAsFloat(_bc.Params)
			if _ac != nil {
				return _ac
			}
			_dd.Log.Debug("M\u006f\u0076\u0065\u0020\u0074\u006f\u003a\u0020\u0025\u0076", _gfd)
			_gf.NewSubPath()
			_gf.MoveTo(_gfd[0], _gfd[1])
		case "\u006c":
			if len(_bc.Params) != 2 {
				_dd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _afg)
				return nil
			}
			_ffg, _ega := _de.GetNumbersAsFloat(_bc.Params)
			if _ega != nil {
				return _ega
			}
			_gf.LineTo(_ffg[0], _ffg[1])
		case "\u0063":
			if len(_bc.Params) != 6 {
				return _afg
			}
			_acf, _acd := _de.GetNumbersAsFloat(_bc.Params)
			if _acd != nil {
				return _acd
			}
			_dd.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _acf)
			_gf.CubicTo(_acf[0], _acf[1], _acf[2], _acf[3], _acf[4], _acf[5])
		case "\u0076", "\u0079":
			if len(_bc.Params) != 4 {
				return _afg
			}
			_aba, _ad := _de.GetNumbersAsFloat(_bc.Params)
			if _ad != nil {
				return _ad
			}
			_dd.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _aba)
			_gf.QuadraticTo(_aba[0], _aba[1], _aba[2], _aba[3])
		case "\u0068":
			_gf.ClosePath()
			_gf.NewSubPath()
		case "\u0072\u0065":
			if len(_bc.Params) != 4 {
				return _afg
			}
			_ffc, _ada := _de.GetNumbersAsFloat(_bc.Params)
			if _ada != nil {
				return _ada
			}
			_gf.DrawRectangle(_ffc[0], _ffc[1], _ffc[2], _ffc[3])
			_gf.NewSubPath()
		case "\u0053":
			_dde, _ddeb := _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
			if _ddeb != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ddeb)
				return _ddeb
			}
			_cga, _ebc := _dde.(*_ed.PdfColorDeviceRGB)
			if !_ebc {
				_dd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _ddeb
			}
			_gf.SetRGBA(_cga.R(), _cga.G(), _cga.B(), 1)
			_gf.Stroke()
		case "\u0073":
			_gff, _cefb := _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
			if _cefb != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cefb)
				return _cefb
			}
			_bfg, _afe := _gff.(*_ed.PdfColorDeviceRGB)
			if !_afe {
				_dd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _cefb
			}
			_gf.ClosePath()
			_gf.NewSubPath()
			_gf.SetRGBA(_bfg.R(), _bfg.G(), _bfg.B(), 1)
			_gf.Stroke()
		case "\u0066", "\u0046":
			_gbe, _fg := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _fg != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fg)
				return _fg
			}
			switch _ede := _gbe.(type) {
			case *_ed.PdfColorDeviceRGB:
				_gf.SetRGBA(_ede.R(), _ede.G(), _ede.B(), 1)
				_gf.SetFillRule(_a.FillRuleWinding)
				_gf.Fill()
			case *_ed.PdfColorPattern:
				_gf.Fill()
			}
			_dd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
		case "\u0066\u002a":
			_bbg, _bea := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _bea != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bea)
				return _bea
			}
			_gffe, _cbe := _bbg.(*_ed.PdfColorDeviceRGB)
			if !_cbe {
				_dd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _bea
			}
			_gf.SetRGBA(_gffe.R(), _gffe.G(), _gffe.B(), 1)
			_gf.SetFillRule(_a.FillRuleEvenOdd)
			_gf.Fill()
		case "\u0042":
			_bab, _ebg := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _ebg != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ebg)
				return _ebg
			}
			switch _afb := _bab.(type) {
			case *_ed.PdfColorDeviceRGB:
				_gf.SetRGBA(_afb.R(), _afb.G(), _afb.B(), 1)
				_gf.SetFillRule(_a.FillRuleWinding)
				_gf.FillPreserve()
				_bab, _ebg = _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
				if _ebg != nil {
					_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ebg)
					return _ebg
				}
				if _egfa, _ceca := _bab.(*_ed.PdfColorDeviceRGB); _ceca {
					_gf.SetRGBA(_egfa.R(), _egfa.G(), _egfa.B(), 1)
					_gf.Stroke()
				}
			case *_ed.PdfColorPattern:
				_gf.SetFillRule(_a.FillRuleWinding)
				_gf.Fill()
				_gf.StrokePattern()
			}
		case "\u0042\u002a":
			_fda, _aebf := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _aebf != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _aebf)
				return _aebf
			}
			switch _fef := _fda.(type) {
			case *_ed.PdfColorDeviceRGB:
				_gf.SetRGBA(_fef.R(), _fef.G(), _fef.B(), 1)
				_gf.SetFillRule(_a.FillRuleEvenOdd)
				_gf.FillPreserve()
				_fda, _aebf = _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
				if _aebf != nil {
					_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _aebf)
					return _aebf
				}
				if _dad, _eed := _fda.(*_ed.PdfColorDeviceRGB); _eed {
					_gf.SetRGBA(_dad.R(), _dad.G(), _dad.B(), 1)
					_gf.Stroke()
				}
			case *_ed.PdfColorPattern:
				_gf.SetFillRule(_a.FillRuleEvenOdd)
				_gf.Fill()
				_gf.StrokePattern()
			}
		case "\u0062":
			_gf.ClosePath()
			_gca, _ggc := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _ggc != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ggc)
				return _ggc
			}
			switch _gbd := _gca.(type) {
			case *_ed.PdfColorDeviceRGB:
				_gf.SetRGBA(_gbd.R(), _gbd.G(), _gbd.B(), 1)
				_gf.NewSubPath()
				_gf.SetFillRule(_a.FillRuleWinding)
				_gf.FillPreserve()
				_gca, _ggc = _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
				if _ggc != nil {
					_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ggc)
					return _ggc
				}
				if _gbea, _dcc := _gca.(*_ed.PdfColorDeviceRGB); _dcc {
					_gf.SetRGBA(_gbea.R(), _gbea.G(), _gbea.B(), 1)
					_gf.Stroke()
				}
			case *_ed.PdfColorPattern:
				_gf.NewSubPath()
				_gf.SetFillRule(_a.FillRuleWinding)
				_gf.Fill()
				_gf.StrokePattern()
			}
		case "\u0062\u002a":
			_gf.ClosePath()
			_eec, _agae := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _agae != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _agae)
				return _agae
			}
			switch _cfb := _eec.(type) {
			case *_ed.PdfColorDeviceRGB:
				_gf.SetRGBA(_cfb.R(), _cfb.G(), _cfb.B(), 1)
				_gf.NewSubPath()
				_gf.SetFillRule(_a.FillRuleEvenOdd)
				_gf.FillPreserve()
				_eec, _agae = _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
				if _agae != nil {
					_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _agae)
					return _agae
				}
				if _ggce, _baa := _eec.(*_ed.PdfColorDeviceRGB); _baa {
					_gf.SetRGBA(_ggce.R(), _ggce.G(), _ggce.B(), 1)
					_gf.Stroke()
				}
			case *_ed.PdfColorPattern:
				_gf.NewSubPath()
				_gf.SetFillRule(_a.FillRuleEvenOdd)
				_gf.Fill()
				_gf.StrokePattern()
			}
		case "\u006e":
			_gf.ClearPath()
		case "\u0057":
			_gf.SetFillRule(_a.FillRuleWinding)
			_gf.ClipPreserve()
		case "\u0057\u002a":
			_gf.SetFillRule(_a.FillRuleEvenOdd)
			_gf.ClipPreserve()
		case "\u0072\u0067":
			_gdg, _bfgf := _gcg.ColorNonStroking.(*_ed.PdfColorDeviceRGB)
			if !_bfgf {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_gf.SetFillRGBA(_gdg.R(), _gdg.G(), _gdg.B(), 1)
		case "\u0052\u0047":
			_ffab, _ebb := _gcg.ColorStroking.(*_ed.PdfColorDeviceRGB)
			if !_ebb {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_gf.SetStrokeRGBA(_ffab.R(), _ffab.G(), _ffab.B(), 1)
		case "\u006b":
			_fead, _abe := _gcg.ColorNonStroking.(*_ed.PdfColorDeviceCMYK)
			if !_abe {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_egfee, _cbfb := _gcg.ColorspaceNonStroking.ColorToRGB(_fead)
			if _cbfb != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_bfa, _abe := _egfee.(*_ed.PdfColorDeviceRGB)
			if !_abe {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _egfee)
				return nil
			}
			_gf.SetFillRGBA(_bfa.R(), _bfa.G(), _bfa.B(), 1)
		case "\u004b":
			_bfe, _acdf := _gcg.ColorStroking.(*_ed.PdfColorDeviceCMYK)
			if !_acdf {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_gfe, _fc := _gcg.ColorspaceStroking.ColorToRGB(_bfe)
			if _fc != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_fba, _acdf := _gfe.(*_ed.PdfColorDeviceRGB)
			if !_acdf {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gfe)
				return nil
			}
			_gf.SetStrokeRGBA(_fba.R(), _fba.G(), _fba.B(), 1)
		case "\u0067":
			_cbff, _dbbe := _gcg.ColorNonStroking.(*_ed.PdfColorDeviceGray)
			if !_dbbe {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_dga, _dgcb := _gcg.ColorspaceNonStroking.ColorToRGB(_cbff)
			if _dgcb != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_aebb, _dbbe := _dga.(*_ed.PdfColorDeviceRGB)
			if !_dbbe {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dga)
				return nil
			}
			_gf.SetFillRGBA(_aebb.R(), _aebb.G(), _aebb.B(), 1)
		case "\u0047":
			_aea, _cbaf := _gcg.ColorStroking.(*_ed.PdfColorDeviceGray)
			if !_cbaf {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_fcg, _ffgc := _gcg.ColorspaceStroking.ColorToRGB(_aea)
			if _ffgc != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_dfe, _cbaf := _fcg.(*_ed.PdfColorDeviceRGB)
			if !_cbaf {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fcg)
				return nil
			}
			_gf.SetStrokeRGBA(_dfe.R(), _dfe.G(), _dfe.B(), 1)
		case "\u0063\u0073":
			if len(_bc.Params) > 0 {
				if _ccc, _cfbc := _de.GetName(_bc.Params[0]); _cfbc && _ccc.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					break
				}
			}
			_fca, _dgab := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _dgab != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_bdd, _ccfgc := _fca.(*_ed.PdfColorDeviceRGB)
			if !_ccfgc {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fca)
				return nil
			}
			_gf.SetFillRGBA(_bdd.R(), _bdd.G(), _bdd.B(), 1)
		case "\u0073\u0063":
			_fdbe, _gbb := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _gbb != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_gad, _dab := _fdbe.(*_ed.PdfColorDeviceRGB)
			if !_dab {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fdbe)
				return nil
			}
			_gf.SetFillRGBA(_gad.R(), _gad.G(), _gad.B(), 1)
		case "\u0073\u0063\u006e":
			if len(_bc.Params) > 0 && len(_cfd.Params) > 0 {
				if _eedb, _cd := _de.GetName(_cfd.Params[0]); _cd && _eedb.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					if _cbce, _fde := _de.GetName(_bc.Params[0]); _fde {
						_def, _dgf := _gc.processGradient(_gf, _bc, _egfe, _cbce)
						if _dgf != nil {
							_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0064\u0061\u0074a\u003a\u0020\u0025\u0076", _dgf)
							break
						}
						if _def == nil {
							_dd.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
							break
						}
						_gf.SetFillStyle(_def)
						_gf.SetStrokeStyle(_def)
						break
					}
				}
			}
			_gce, _cbffd := _gcg.ColorspaceNonStroking.ColorToRGB(_gcg.ColorNonStroking)
			if _cbffd != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorNonStroking)
				return nil
			}
			_acg, _abaa := _gce.(*_ed.PdfColorDeviceRGB)
			if !_abaa {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gce)
				return nil
			}
			_gf.SetFillRGBA(_acg.R(), _acg.G(), _acg.B(), 1)
		case "\u0043\u0053":
			if len(_bc.Params) > 0 {
				if _dbd, _eee := _de.GetName(_bc.Params[0]); _eee && _dbd.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					break
				}
			}
			_gga, _ffe := _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
			if _ffe != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_abd, _aaf := _gga.(*_ed.PdfColorDeviceRGB)
			if !_aaf {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gga)
				return nil
			}
			_gf.SetStrokeRGBA(_abd.R(), _abd.G(), _abd.B(), 1)
		case "\u0053\u0043":
			_baf, _aca := _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
			if _aca != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_afa, _fffg := _baf.(*_ed.PdfColorDeviceRGB)
			if !_fffg {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _baf)
				return nil
			}
			_gf.SetStrokeRGBA(_afa.R(), _afa.G(), _afa.B(), 1)
		case "\u0053\u0043\u004e":
			if len(_bc.Params) > 0 && len(_cfd.Params) > 0 {
				if _fgd, _cgec := _de.GetName(_cfd.Params[0]); _cgec && _fgd.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					if _egeb, _egg := _de.GetName(_bc.Params[0]); _egg {
						_bbb, _bfb := _gc.processGradient(_gf, _bc, _egfe, _egeb)
						if _bfb != nil {
							_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0064\u0061\u0074a\u003a\u0020\u0025\u0076", _bfb)
							break
						}
						if _bbb == nil {
							_dd.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
							break
						}
						_gf.SetFillStyle(_bbb)
						_gf.SetStrokeStyle(_bbb)
						break
					}
				}
			}
			_ceb, _egd := _gcg.ColorspaceStroking.ColorToRGB(_gcg.ColorStroking)
			if _egd != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gcg.ColorStroking)
				return nil
			}
			_gcgd, _adg := _ceb.(*_ed.PdfColorDeviceRGB)
			if !_adg {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ceb)
				return nil
			}
			_gf.SetStrokeRGBA(_gcgd.R(), _gcgd.G(), _gcgd.B(), 1)
		case "\u0073\u0068":
			if len(_bc.Params) != 1 {
				_dd.Log.Debug("\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0020\u0070\u0061r\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
				break
			}
			_gcf, _cggd := _de.GetName(_bc.Params[0])
			if !_cggd {
				_dd.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020g\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006eg\u0020n\u0061\u006d\u0065")
				break
			}
			_cece, _cggd := _egfe.GetShadingByName(*_gcf)
			if !_cggd {
				_dd.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020g\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006eg\u0020d\u0061\u0074\u0061")
				break
			}
			_gda, _gaf, _bfgfc := _gc.processShading(_gf, _cece)
			if _bfgfc != nil {
				_dd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006e\u0067\u0020d\u0061\u0074a\u003a\u0020\u0025\u0076", _bfgfc)
				break
			}
			if _gda == nil {
				_dd.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
				break
			}
			_gcef, _bfgfc := _gaf.ToFloat64Array()
			if _bfgfc != nil {
				_dd.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073: \u0025\u0076", _bfgfc)
				break
			}
			_gf.DrawRectangle(_gcef[0], _gcef[1], _gcef[2], _gcef[3])
			_gf.NewSubPath()
			_gf.SetFillStyle(_gda)
			_gf.SetStrokeStyle(_gda)
			_gf.Fill()
		case "\u0044\u006f":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_eecd, _dbcb := _de.GetName(_bc.Params[0])
			if !_dbcb {
				return _ag
			}
			_, _cbab := _egfe.GetXObjectByName(*_eecd)
			switch _cbab {
			case _ed.XObjectTypeImage:
				_dd.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067e\u003a\u0020\u0025\u0073", _eecd.String())
				_gee, _daf := _egfe.GetXObjectImageByName(*_eecd)
				if _daf != nil {
					return _daf
				}
				_cfa, _daf := _gee.ToImage()
				if _daf != nil {
					_dd.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _daf)
					return nil
				}
				if _egae := _gee.ColorSpace; _egae != nil {
					var _ebca bool
					switch _egae.(type) {
					case *_ed.PdfColorspaceSpecialIndexed:
						_ebca = true
					}
					if _ebca {
						if _feef, _bbf := _egae.ImageToRGB(*_cfa); _bbf != nil {
							_dd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006fnv\u0065r\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0052G\u0042\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020i\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
						} else {
							_cfa = &_feef
						}
					}
				}
				_gceg := _gf.FillPattern().ColorAt(0, 0)
				var _befc _db.Image
				if _gee.Mask != nil {
					if _befc, _daf = _bba(_gee.Mask, _gceg); _daf != nil {
						_dd.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				} else if _gee.SMask != nil {
					if _befc, _daf = _abdg(_gee.SMask, _gceg); _daf != nil {
						_dd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				}
				var _bda _db.Image
				if _cbeg, _ := _de.GetBoolVal(_gee.ImageMask); _cbeg {
					_bda = _cea(_cfa, _gceg)
				} else {
					_bda, _daf = _cfa.ToGoImage()
					if _daf != nil {
						_dd.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _daf)
						return nil
					}
				}
				if _befc != nil {
					_bda = _cgba(_bda, _befc)
				}
				_ddgd := _bda.Bounds()
				_gf.Push()
				_gf.Scale(1.0/float64(_ddgd.Dx()), -1.0/float64(_ddgd.Dy()))
				_gf.DrawImageAnchored(_bda, 0, 0, 0, 1)
				_gf.Pop()
			case _ed.XObjectTypeForm:
				_dd.Log.Debug("\u0058\u004fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u0073", _eecd.String())
				_eeg, _cbcd := _egfe.GetXObjectFormByName(*_eecd)
				if _cbcd != nil {
					return _cbcd
				}
				_bcff, _cbcd := _eeg.GetContentStream()
				if _cbcd != nil {
					return _cbcd
				}
				_ddd := _eeg.Resources
				if _ddd == nil {
					_ddd = _egfe
				}
				_gf.Push()
				if _eeg.Matrix != nil {
					_gef, _fcf := _de.GetArray(_eeg.Matrix)
					if !_fcf {
						return _ag
					}
					_aggf, _bfeg := _de.GetNumbersAsFloat(_gef.Elements())
					if _bfeg != nil {
						return _bfeg
					}
					if len(_aggf) != 6 {
						return _afg
					}
					_bfeb := _ba.NewMatrix(_aggf[0], _aggf[1], _aggf[2], _aggf[3], _aggf[4], _aggf[5])
					_gf.SetMatrix(_gf.Matrix().Mult(_bfeb))
				}
				if _eeg.BBox != nil {
					_ca, _ace := _de.GetArray(_eeg.BBox)
					if !_ace {
						return _ag
					}
					_efa, _cgge := _de.GetNumbersAsFloat(_ca.Elements())
					if _cgge != nil {
						return _cgge
					}
					if len(_efa) != 4 {
						_dd.Log.Debug("\u004c\u0065\u006e\u0020\u003d\u0020\u0025\u0064", len(_efa))
						return _afg
					}
					_gf.DrawRectangle(_efa[0], _efa[1], _efa[2]-_efa[0], _efa[3]-_efa[1])
					_gf.SetRGBA(1, 0, 0, 1)
					_gf.Clip()
				} else {
					_dd.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0052\u0065q\u0075\u0069\u0072e\u0064\u0020\u0042\u0042\u006f\u0078\u0020\u006d\u0069ss\u0069\u006e\u0067 \u006f\u006e \u0058\u004f\u0062\u006a\u0065\u0063t\u0020\u0046o\u0072\u006d")
				}
				_cbcd = _gc.renderContentStream(_gf, string(_bcff), _ddd)
				if _cbcd != nil {
					return _cbcd
				}
				_gf.Pop()
			}
		case "\u0042\u0049":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_afgd, _gaa := _bc.Params[0].(*_eg.ContentStreamInlineImage)
			if !_gaa {
				return nil
			}
			_fcaf, _ceg := _afgd.ToImage(_egfe)
			if _ceg != nil {
				_dd.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _ceg)
				return nil
			}
			_beb, _ceg := _fcaf.ToGoImage()
			if _ceg != nil {
				_dd.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _ceg)
				return nil
			}
			_aab := _beb.Bounds()
			_gf.Push()
			_gf.Scale(1.0/float64(_aab.Dx()), -1.0/float64(_aab.Dy()))
			_gf.DrawImageAnchored(_beb, 0, 0, 0, 1)
			_gf.Pop()
		case "\u0042\u0054":
			_ge.Reset()
		case "\u0045\u0054":
			_ge.Reset()
		case "\u0054\u0072":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_bdc, _fbg := _de.GetNumberAsFloat(_bc.Params[0])
			if _fbg != nil {
				return _fbg
			}
			_ge.Tr = _a.TextRenderingMode(_bdc)
		case "\u0054\u004c":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_fdd, _bade := _de.GetNumberAsFloat(_bc.Params[0])
			if _bade != nil {
				return _bade
			}
			_ge.Tl = _fdd
		case "\u0054\u0063":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_bcd, _bg := _de.GetNumberAsFloat(_bc.Params[0])
			if _bg != nil {
				return _bg
			}
			_dd.Log.Debug("\u0054\u0063\u003a\u0020\u0025\u0076", _bcd)
			_ge.Tc = _bcd
		case "\u0054\u0077":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_dbde, _fcd := _de.GetNumberAsFloat(_bc.Params[0])
			if _fcd != nil {
				return _fcd
			}
			_dd.Log.Debug("\u0054\u0077\u003a\u0020\u0025\u0076", _dbde)
			_ge.Tw = _dbde
		case "\u0054\u007a":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_bebc, _gbgf := _de.GetNumberAsFloat(_bc.Params[0])
			if _gbgf != nil {
				return _gbgf
			}
			_ge.Th = _bebc
		case "\u0054\u0073":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_fdca, _gfg := _de.GetNumberAsFloat(_bc.Params[0])
			if _gfg != nil {
				return _gfg
			}
			_ge.Ts = _fdca
		case "\u0054\u0064":
			if len(_bc.Params) != 2 {
				return _afg
			}
			_afge, _bce := _de.GetNumbersAsFloat(_bc.Params)
			if _bce != nil {
				return _bce
			}
			_dd.Log.Debug("\u0054\u0064\u003a\u0020\u0025\u0076", _afge)
			_ge.ProcTd(_afge[0], _afge[1])
		case "\u0054\u0044":
			if len(_bc.Params) != 2 {
				return _afg
			}
			_cbb, _abb := _de.GetNumbersAsFloat(_bc.Params)
			if _abb != nil {
				return _abb
			}
			_dd.Log.Debug("\u0054\u0044\u003a\u0020\u0025\u0076", _cbb)
			_ge.ProcTD(_cbb[0], _cbb[1])
		case "\u0054\u002a":
			_ge.ProcTStar()
		case "\u0054\u006d":
			if len(_bc.Params) != 6 {
				return _afg
			}
			_adaf, _ddf := _de.GetNumbersAsFloat(_bc.Params)
			if _ddf != nil {
				return _ddf
			}
			_dd.Log.Debug("\u0054\u0065x\u0074\u0020\u006da\u0074\u0072\u0069\u0078\u003a\u0020\u0025\u002b\u0076", _adaf)
			_ge.ProcTm(_adaf[0], _adaf[1], _adaf[2], _adaf[3], _adaf[4], _adaf[5])
		case "\u0027":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_bbd, _fgg := _de.GetStringBytes(_bc.Params[0])
			if !_fgg {
				return _ag
			}
			_dd.Log.Debug("\u0027\u0020\u0073t\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_bbd))
			_ge.ProcQ(_bbd, _gf)
		case "\u0022":
			if len(_bc.Params) != 3 {
				return _afg
			}
			_gcd, _ec := _de.GetNumberAsFloat(_bc.Params[0])
			if _ec != nil {
				return _ec
			}
			_bace, _ec := _de.GetNumberAsFloat(_bc.Params[1])
			if _ec != nil {
				return _ec
			}
			_edg, _fbc := _de.GetStringBytes(_bc.Params[2])
			if !_fbc {
				return _ag
			}
			_ge.ProcDQ(_edg, _gcd, _bace, _gf)
		case "\u0054\u006a":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_aff, _ccgb := _de.GetStringBytes(_bc.Params[0])
			if !_ccgb {
				return _ag
			}
			_dd.Log.Debug("\u0054j\u0020s\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0060\u0025\u0073\u0060", string(_aff))
			_ge.ProcTj(_aff, _gf)
		case "\u0054\u004a":
			if len(_bc.Params) != 1 {
				return _afg
			}
			_fece, _dadb := _de.GetArray(_bc.Params[0])
			if !_dadb {
				_dd.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _fece)
				return _ag
			}
			_dd.Log.Debug("\u0054\u004a\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u002b\u0076", _fece)
			for _, _eag := range _fece.Elements() {
				switch _agge := _eag.(type) {
				case *_de.PdfObjectString:
					if _agge != nil {
						_ge.ProcTj(_agge.Bytes(), _gf)
					}
				case *_de.PdfObjectFloat, *_de.PdfObjectInteger:
					_cdg, _edb := _de.GetNumberAsFloat(_agge)
					if _edb == nil {
						_ge.Translate(-_cdg*0.001*_ge.Tf.Size*_ge.Th/100.0, 0)
					}
				}
			}
		case "\u0054\u0066":
			if len(_bc.Params) != 2 {
				return _afg
			}
			_dd.Log.Debug("\u0025\u0023\u0076", _bc.Params)
			_eedg, _ebbe := _de.GetName(_bc.Params[0])
			if !_ebbe || _eedg == nil {
				_dd.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006e\u0061m\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _bc.Params[0])
				return _ag
			}
			_dd.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073", _eedg.String())
			_acac, _cbae := _de.GetNumberAsFloat(_bc.Params[1])
			if _cbae != nil {
				_dd.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0069z\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _bc.Params[1])
				return _ag
			}
			_dd.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0076", _acac)
			_badg, _dbeb := _egfe.GetFontByName(*_eedg)
			if !_dbeb {
				_dd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0025s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _eedg.String())
				return _b.New("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
			}
			_dd.Log.Debug("\u0046\u006f\u006e\u0074\u003a\u0020\u0025\u0054", _badg)
			_aabf, _ebbe := _de.GetDict(_badg)
			if !_ebbe {
				_dd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0067e\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074")
				return _ag
			}
			_fa, _cbae := _ed.NewPdfFontFromPdfObject(_aabf)
			if _cbae != nil {
				_dd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
				return _cbae
			}
			_deed := _fa.BaseFont()
			if _deed == "" {
				_deed = _eedg.String()
			}
			_cdgb, _ebbe := _gd[_deed]
			if !_ebbe {
				_cdgb, _cbae = _a.NewTextFont(_fa, _acac)
				if _cbae != nil {
					_dd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cbae)
				}
			}
			if _cdgb == nil {
				if len(_deed) > 7 && _deed[6] == '+' {
					_deed = _deed[7:]
				}
				_bafg := []string{_deed, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
				for _, _bgf := range _bafg {
					_dd.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _bgf)
					if _cdgb, _ebbe = _gd[_bgf]; _ebbe {
						break
					}
					_gcec := _bd.Match(_bgf)
					if _gcec == nil {
						_dd.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _bgf)
						continue
					}
					_cdgb, _cbae = _a.NewTextFontFromPath(_gcec.Filename, _acac)
					if _cbae != nil {
						_dd.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _gcec.Filename)
						continue
					}
					_dd.Log.Debug("\u0053\u0075\u0062\u0073\u0074\u0069t\u0075\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073 \u0077\u0069\u0074\u0068\u0020\u0025\u0073 \u0028\u0025\u0073\u0029", _deed, _gcec.Name, _gcec.Filename)
					_gd[_bgf] = _cdgb
					break
				}
			}
			if _cdgb == nil {
				_dd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020n\u006f\u0074\u0020\u0066\u0069\u006ed\u0020\u0061\u006e\u0079\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065 \u0066\u006f\u006e\u0074")
				return _b.New("\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0066\u0069\u006e\u0064\u0020a\u006ey\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u006e\u0074")
			}
			_ge.ProcTf(_cdgb.WithSize(_acac, _fa))
		case "\u0042\u004d\u0043", "\u0042\u0044\u0043":
		case "\u0045\u004d\u0043":
		default:
			_dd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u006f\u0070\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u0073", _bc.Operand)
		}
		_cfd = _bc
		return nil
	})
	_ccf = _ccg.Process(_ffaf)
	if _ccf != nil {
		return _ccf
	}
	return nil
}

// RenderToPath converts the specified PDF page into an image and saves the
// result at the specified location.
func (_fdc *ImageDevice) RenderToPath(page *_ed.PdfPage, outputPath string) error {
	_dgc, _ga := _fdc.Render(page)
	if _ga != nil {
		return _ga
	}
	_fdb := _e.ToLower(_f.Ext(outputPath))
	if _fdb == "" {
		return _b.New("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065cog\u006eiz\u0065\u0020\u006f\u0075\u0074\u0070\u0075t \u0066\u0069\u006c\u0065\u0020\u0074\u0079p\u0065")
	}
	switch _fdb {
	case "\u002e\u0070\u006e\u0067":
		return _bcfb(outputPath, _dgc)
	case "\u002e\u006a\u0070\u0067", "\u002e\u006a\u0070e\u0067":
		return _ggdf(outputPath, _dgc, 100)
	}
	return _cge.Errorf("\u0075\u006e\u0072\u0065\u0063\u006fg\u006e\u0069\u007a\u0065\u0064\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020f\u0069\u006c\u0065\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _fdb)
}
func _bcfb(_gde string, _gbc _db.Image) error {
	_ecf, _ebbd := _cb.Create(_gde)
	if _ebbd != nil {
		return _ebbd
	}
	defer _ecf.Close()
	return _cg.Encode(_ecf, _gbc)
}
func _cgba(_fbf, _eedd _db.Image) _db.Image {
	_bec, _aabff := _eedd.Bounds().Size(), _fbf.Bounds().Size()
	_cag, _bcfe := _bec.X, _bec.Y
	if _aabff.X > _cag {
		_cag = _aabff.X
	}
	if _aabff.Y > _bcfe {
		_bcfe = _aabff.Y
	}
	_efaa := _db.Rect(0, 0, _cag, _bcfe)
	if _bec.X != _cag || _bec.Y != _bcfe {
		_cggb := _db.NewRGBA(_efaa)
		_eb.BiLinear.Scale(_cggb, _efaa, _fbf, _eedd.Bounds(), _eb.Over, nil)
		_eedd = _cggb
	}
	if _aabff.X != _cag || _aabff.Y != _bcfe {
		_ddgb := _db.NewRGBA(_efaa)
		_eb.BiLinear.Scale(_ddgb, _efaa, _fbf, _fbf.Bounds(), _eb.Over, nil)
		_fbf = _ddgb
	}
	_daa := _db.NewRGBA(_efaa)
	_eb.DrawMask(_daa, _efaa, _fbf, _db.Point{}, _eedd, _db.Point{}, _eb.Over)
	return _daa
}
func _ggdf(_ecee string, _gdgg _db.Image, _gae int) error {
	_dgb, _fdg := _cb.Create(_ecee)
	if _fdg != nil {
		return _fdg
	}
	defer _dgb.Close()
	return _ced.Encode(_dgb, _gdgg, &_ced.Options{Quality: _gae})
}
func _cea(_bbbc *_ed.Image, _cfec _ea.Color) _db.Image {
	_bgfaa, _gdd := int(_bbbc.Width), int(_bbbc.Height)
	_cbafb := _db.NewRGBA(_db.Rect(0, 0, _bgfaa, _gdd))
	for _efda := 0; _efda < _gdd; _efda++ {
		for _fdcg := 0; _fdcg < _bgfaa; _fdcg++ {
			_befb, _cgfg := _bbbc.ColorAt(_fdcg, _efda)
			if _cgfg != nil {
				_dd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _fdcg, _efda)
				continue
			}
			_dgfg, _defa, _gaef, _ := _befb.RGBA()
			var _daea _ea.Color
			if _dgfg+_defa+_gaef == 0 {
				_daea = _cfec
			} else {
				_daea = _ea.Transparent
			}
			_cbafb.Set(_fdcg, _efda, _daea)
		}
	}
	return _cbafb
}
func (_dffe renderer) processRadialShading(_bgec _a.Context, _cbg *_ed.PdfShading) (_a.Gradient, *_de.PdfObjectArray, error) {
	_gcc := _cbg.GetContext().(*_ed.PdfShadingType3)
	if len(_gcc.Function) == 0 {
		return nil, nil, _b.New("\u006e\u006f\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074 \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0073\u006b\u0069\u0070\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	_efg, _cac := _gcc.Coords.ToFloat64Array()
	if _cac != nil {
		return nil, nil, _b.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020\u0073h\u0061\u0064\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	_cae := _cbg.ColorSpace
	_deea := _de.MakeArrayFromFloats([]float64{0, 0, 1, 1})
	var _eddb, _cfbf, _cfbd, _eace, _gdb, _efc float64
	_eddb, _cfbf = _bgec.Matrix().Transform(_efg[0], _efg[1])
	_cfbd, _eace = _bgec.Matrix().Transform(_efg[3], _efg[4])
	_gdb, _ = _bgec.Matrix().Transform(_efg[2], 0)
	_efc, _ = _bgec.Matrix().Transform(_efg[5], 0)
	_cdge, _ := _bgec.Matrix().Translation()
	_gdb -= _cdge
	_efc -= _cdge
	for _cbcda, _ccd := range _efg {
		if _cbcda == 2 || _cbcda == 5 {
			continue
		}
		if _ccd > 1.0 {
			_eda := _d.Min(_eddb-_gdb, _cfbd-_efc)
			_fgba := _d.Min(_cfbf-_gdb, _eace-_efc)
			_abef := _d.Max(_eddb+_gdb, _cfbd+_efc)
			_gebd := _d.Max(_cfbf+_gdb, _eace+_efc)
			_dfc := _abef - _eda
			_fag := _fgba - _gebd
			_deea = _de.MakeArrayFromFloats([]float64{_eda, _fgba, _dfc, _fag})
			break
		}
	}
	_dfb := _g.NewRadialGradient(_eddb, _cfbf, _gdb, _cfbd, _eace, _efc)
	if _cee, _bcbb := _gcc.Function[0].(*_ed.PdfFunctionType2); _bcbb {
		_dfb, _cac = _eaag(_dfb, _cee, _cae, 1.0, true)
	} else if _ebd, _cacf := _gcc.Function[0].(*_ed.PdfFunctionType3); _cacf {
		_bgfa := append([]float64{0}, _ebd.Bounds...)
		_bgfa = append(_bgfa, 1.0)
		_dfb, _cac = _deeg(_dfb, _ebd, _cae, _bgfa)
	}
	if _cac != nil {
		return nil, nil, _cac
	}
	return _dfb, _deea, nil
}

var (
	_ag  = _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_afg = _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func _cecb(_cefbf *_ed.Image, _gddg _ea.Color) _db.Image {
	_eff, _gdeb := int(_cefbf.Width), int(_cefbf.Height)
	_gaeb := _db.NewRGBA(_db.Rect(0, 0, _eff, _gdeb))
	for _acdd := 0; _acdd < _gdeb; _acdd++ {
		for _egga := 0; _egga < _eff; _egga++ {
			_gdc, _acgd := _cefbf.ColorAt(_egga, _acdd)
			if _acgd != nil {
				_dd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _egga, _acdd)
				continue
			}
			_gged, _cdf, _dcce, _ := _gdc.RGBA()
			var _fgeg _ea.Color
			if _gged+_cdf+_dcce == 0 {
				_fgeg = _ea.Transparent
			} else {
				_fgeg = _gddg
			}
			_gaeb.Set(_egga, _acdd, _fgeg)
		}
	}
	return _gaeb
}

const (
	ShadingTypeFunctionBased PdfShadingType = 1
	ShadingTypeAxial         PdfShadingType = 2
	ShadingTypeRadial        PdfShadingType = 3
	ShadingTypeFreeForm      PdfShadingType = 4
	ShadingTypeLatticeForm   PdfShadingType = 5
	ShadingTypeCoons         PdfShadingType = 6
	ShadingTypeTensorProduct PdfShadingType = 7
)

func _bba(_adab _de.PdfObject, _gbaf _ea.Color) (_db.Image, error) {
	_fac, _cbgg := _de.GetStream(_adab)
	if !_cbgg {
		return nil, nil
	}
	_gab, _edcg := _ed.NewXObjectImageFromStream(_fac)
	if _edcg != nil {
		return nil, _edcg
	}
	_dbbc, _edcg := _gab.ToImage()
	if _edcg != nil {
		return nil, _edcg
	}
	return _cea(_dbbc, _gbaf), nil
}

type renderer struct{ _fee float64 }

func _eaag(_edf _a.Gradient, _gfde *_ed.PdfFunctionType2, _gcb _ed.PdfColorspace, _fdcf float64, _dfbe bool) (_a.Gradient, error) {
	switch _gcb.(type) {
	case *_ed.PdfColorspaceDeviceRGB:
		if len(_gfde.C0) != 3 || len(_gfde.C1) != 3 {
			return nil, _b.New("\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0052\u0047\u0042\u0020\u0063o\u006co\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
		}
		_cggc := _gfde.C0
		_cfe := _gfde.C1
		if _dfbe {
			_edf.AddColorStop(0.0, _ea.RGBA{R: uint8(_cggc[0] * 255), G: uint8(_cggc[1] * 255), B: uint8(_cggc[2] * 255), A: 255})
		}
		_edf.AddColorStop(_fdcf, _ea.RGBA{R: uint8(_cfe[0] * 255), G: uint8(_cfe[1] * 255), B: uint8(_cfe[2] * 255), A: 255})
	case *_ed.PdfColorspaceDeviceCMYK:
		if len(_gfde.C0) != 4 || len(_gfde.C1) != 4 {
			return nil, _b.New("\u0069\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072 \u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
		}
		_dec := _gfde.C0
		_afea := _gfde.C1
		if _dfbe {
			_edf.AddColorStop(0.0, _ea.CMYK{C: uint8(_dec[0] * 255), M: uint8(_dec[1] * 255), Y: uint8(_dec[2] * 255), K: uint8(_dec[3] * 255)})
		}
		_edf.AddColorStop(_fdcf, _ea.CMYK{C: uint8(_afea[0] * 255), M: uint8(_afea[1] * 255), Y: uint8(_afea[2] * 255), K: uint8(_afea[3] * 255)})
	default:
		return nil, _cge.Errorf("u\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072 \u0073\u0070\u0061c\u0065:\u0020\u0025\u0073", _gcb.String())
	}
	return _edf, nil
}
func _abdg(_ebbdg _de.PdfObject, _fbcb _ea.Color) (_db.Image, error) {
	_gdebf, _aebe := _de.GetStream(_ebbdg)
	if !_aebe {
		return nil, nil
	}
	_cdda, _cad := _ed.NewXObjectImageFromStream(_gdebf)
	if _cad != nil {
		return nil, _cad
	}
	_eded, _cad := _cdda.ToImage()
	if _cad != nil {
		return nil, _cad
	}
	return _cecb(_eded, _fbcb), nil
}
func (_aad renderer) processShading(_dfa _a.Context, _dgg *_ed.PdfShading) (_a.Gradient, *_de.PdfObjectArray, error) {
	_dfac := int64(*_dgg.ShadingType)
	if _dfac == int64(ShadingTypeAxial) {
		return _aad.processLinearShading(_dfa, _dgg)
	} else if _dfac == int64(ShadingTypeRadial) {
		return _aad.processRadialShading(_dfa, _dgg)
	} else {
		_dd.Log.Debug(_cge.Sprintf("\u0050r\u006f\u0063e\u0073\u0073\u0069n\u0067\u0020\u0067\u0072\u0061\u0064\u0069e\u006e\u0074\u0020\u0074\u0079\u0070e\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074 \u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064", _dfac))
	}
	return nil, nil, nil
}
func (_ecg renderer) processGradient(_fgb _a.Context, _fggc *_eg.ContentStreamOperation, _bcb *_ed.PdfPageResources, _dfg *_de.PdfObjectName) (_a.Gradient, error) {
	if _egc, _bcc := _bcb.GetPatternByName(*_dfg); _bcc && _egc.IsShading() {
		_dffc := _egc.GetAsShadingPattern().Shading
		_geb, _, _gffea := _ecg.processShading(_fgb, _dffc)
		if _gffea != nil {
			return nil, _gffea
		}
		return _geb, nil
	}
	return nil, nil
}
func (_ege renderer) renderPage(_aga _a.Context, _gbg *_ed.PdfPage, _cec _ba.Matrix, _cba bool) error {
	if !_cba {
		_gg := _ed.FieldFlattenOpts{AnnotFilterFunc: func(_bad *_ed.PdfAnnotation) bool {
			switch _bad.GetContext().(type) {
			case *_ed.PdfAnnotationLine:
				return true
			case *_ed.PdfAnnotationSquare:
				return true
			case *_ed.PdfAnnotationCircle:
				return true
			case *_ed.PdfAnnotationPolygon:
				return true
			case *_ed.PdfAnnotationPolyLine:
				return true
			}
			return false
		}}
		_dbe := _dee.FieldAppearance{}
		_dae := _gbg.FlattenFieldsWithOpts(_dbe, &_gg)
		if _dae != nil {
			_dd.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0064u\u0072\u0069n\u0067\u0020\u0061\u006e\u006e\u006f\u0074\u0061t\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0074\u0074\u0065\u006e\u0069n\u0067\u0020\u0025\u0076", _dae)
		}
	}
	_dgd, _egf := _gbg.GetAllContentStreams()
	if _egf != nil {
		return _egf
	}
	if _dda := _cec; !_dda.Identity() {
		_dgd = _cge.Sprintf("%\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066 \u0025\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066\u0020\u0063\u006d\u0020%\u0073", _dda[0], _dda[1], _dda[3], _dda[4], _dda[6], _dda[7], _dgd)
	}
	_aga.Translate(0, float64(_aga.Height()))
	_aga.Scale(1, -1)
	_aga.Push()
	_aga.SetRGBA(1, 1, 1, 1)
	_aga.DrawRectangle(0, 0, float64(_aga.Width()), float64(_aga.Height()))
	_aga.Fill()
	_aga.Pop()
	_aga.SetLineWidth(1.0)
	_aga.SetRGBA(0, 0, 0, 1)
	return _ege.renderContentStream(_aga, _dgd, _gbg.Resources)
}
func _fgef(_ccdc, _fecb, _adc float64) _be.BoundingBox {
	return _be.Path{Points: []_be.Point{_be.NewPoint(0, 0).Rotate(_adc), _be.NewPoint(_ccdc, 0).Rotate(_adc), _be.NewPoint(0, _fecb).Rotate(_adc), _be.NewPoint(_ccdc, _fecb).Rotate(_adc)}}.GetBoundingBox()
}

// Render converts the specified PDF page into an image, flattens annotations by default and returns the result.
func (_dg *ImageDevice) Render(page *_ed.PdfPage) (_db.Image, error) {
	return _dg.RenderWithOpts(page, false)
}
