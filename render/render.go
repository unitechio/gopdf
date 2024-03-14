package render

import (
	_ag "errors"
	_cd "fmt"
	_ec "image"
	_eb "image/color"
	_e "image/draw"
	_f "image/jpeg"
	_ed "image/png"
	_db "math"
	_c "os"
	_g "path/filepath"
	_d "strings"

	_ef "bitbucket.org/shenghui0779/gopdf/annotator"
	_ceb "bitbucket.org/shenghui0779/gopdf/common"
	_gd "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ce "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_df "bitbucket.org/shenghui0779/gopdf/core"
	_cb "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_eg "bitbucket.org/shenghui0779/gopdf/model"
	_fe "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_b "bitbucket.org/shenghui0779/gopdf/render/internal/context/imagerender"
	_ff "github.com/adrg/sysfont"
	_fa "golang.org/x/image/draw"
)

func (_gaf renderer) processGradient(_ded _fe.Context, _gcdb *_gd.ContentStreamOperation, _fcge *_eg.PdfPageResources, _bdec *_df.PdfObjectName) (_fe.Gradient, error) {
	if _eccd, _ccg := _fcge.GetPatternByName(*_bdec); _ccg && _eccd.IsShading() {
		_dfa := _eccd.GetAsShadingPattern().Shading
		_bbcff, _, _baa := _gaf.processShading(_ded, _dfa)
		if _baa != nil {
			return nil, _baa
		}
		return _bbcff, nil
	}
	return nil, nil
}

func _aaa(_fcabe, _cfa, _gdfg float64) _ce.BoundingBox {
	return _ce.Path{Points: []_ce.Point{_ce.NewPoint(0, 0).Rotate(_gdfg), _ce.NewPoint(_fcabe, 0).Rotate(_gdfg), _ce.NewPoint(0, _cfa).Rotate(_gdfg), _ce.NewPoint(_fcabe, _cfa).Rotate(_gdfg)}}.GetBoundingBox()
}

// Render converts the specified PDF page into an image, flattens annotations by default and returns the result.
func (_gdg *ImageDevice) Render(page *_eg.PdfPage) (_ec.Image, error) {
	return _gdg.RenderWithOpts(page, false)
}

func _fdd(_dff string, _agae _ec.Image, _febe int) error {
	_abaf, _acbg := _c.Create(_dff)
	if _acbg != nil {
		return _acbg
	}
	defer _abaf.Close()
	return _f.Encode(_abaf, _agae, &_f.Options{Quality: _febe})
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

func (_bcc renderer) renderContentStream(_edg _fe.Context, _efa string, _ea *_eg.PdfPageResources) error {
	_fbaa, _abe := _gd.NewContentStreamParser(_efa).Parse()
	if _abe != nil {
		return _abe
	}
	_eag := _edg.TextState()
	_eag.GlobalScale = _bcc._fdc
	_gg := map[string]*_fe.TextFont{}
	_ad := _ff.NewFinder(&_ff.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066", "\u002e\u0074\u0074\u0063"}})
	var _ceg *_gd.ContentStreamOperation
	_dg := _gd.NewContentStreamProcessor(*_fbaa)
	_dg.AddHandler(_gd.HandlerConditionEnumAllOperands, "", func(_cce *_gd.ContentStreamOperation, _cbf _gd.GraphicsState, _eff *_eg.PdfPageResources) error {
		_ceb.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0025\u0073", _cce.Operand)
		switch _cce.Operand {
		case "\u0071":
			_edg.Push()
		case "\u0051":
			_edg.Pop()
			_eag = _edg.TextState()
		case "\u0063\u006d":
			if len(_cce.Params) != 6 {
				return _ecd
			}
			_add, _fbc := _df.GetNumbersAsFloat(_cce.Params)
			if _fbc != nil {
				return _fbc
			}
			_ee := _cb.NewMatrix(_add[0], _add[1], _add[2], _add[3], _add[4], _add[5])
			_ceb.Log.Debug("\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074a\u0074\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078\u003a \u0025\u002b\u0076", _ee)
			_edg.SetMatrix(_edg.Matrix().Mult(_ee))
		case "\u0077":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_dgb, _bbf := _df.GetNumbersAsFloat(_cce.Params)
			if _bbf != nil {
				return _bbf
			}
			_edg.SetLineWidth(_dgb[0])
		case "\u004a":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_def, _eaa := _df.GetIntVal(_cce.Params[0])
			if !_eaa {
				return _eba
			}
			switch _def {
			case 0:
				_edg.SetLineCap(_fe.LineCapButt)
			case 1:
				_edg.SetLineCap(_fe.LineCapRound)
			case 2:
				_edg.SetLineCap(_fe.LineCapSquare)
			default:
				_ceb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u006ee\u0020\u0063\u0061\u0070\u0020\u0073\u0074\u0079\u006c\u0065:\u0020\u0025\u0064", _def)
				return _ecd
			}
		case "\u006a":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_afg, _da := _df.GetIntVal(_cce.Params[0])
			if !_da {
				return _eba
			}
			switch _afg {
			case 0:
				_edg.SetLineJoin(_fe.LineJoinBevel)
			case 1:
				_edg.SetLineJoin(_fe.LineJoinRound)
			case 2:
				_edg.SetLineJoin(_fe.LineJoinBevel)
			default:
				_ceb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u006e\u0065\u0020\u006a\u006f\u0069\u006e \u0073\u0074\u0079l\u0065:\u0020\u0025\u0064", _afg)
				return _ecd
			}
		case "\u004d":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_fbb, _daa := _df.GetNumbersAsFloat(_cce.Params)
			if _daa != nil {
				return _daa
			}
			_ = _fbb
			_ceb.Log.Debug("\u004di\u0074\u0065\u0072\u0020l\u0069\u006d\u0069\u0074\u0020n\u006ft\u0020s\u0075\u0070\u0070\u006f\u0072\u0074\u0065d")
		case "\u0064":
			if len(_cce.Params) != 2 {
				return _ecd
			}
			_aeb, _aa := _df.GetArray(_cce.Params[0])
			if !_aa {
				return _eba
			}
			_afe, _aa := _df.GetIntVal(_cce.Params[1])
			if !_aa {
				_, _fbd := _df.GetFloatVal(_cce.Params[1])
				if !_fbd {
					return _eba
				}
			}
			_adg, _gdd := _df.GetNumbersAsFloat(_aeb.Elements())
			if _gdd != nil {
				return _gdd
			}
			_edg.SetDash(_adg...)
			_ = _afe
			_ceb.Log.Debug("\u004c\u0069n\u0065\u0020\u0064\u0061\u0073\u0068\u0020\u0070\u0068\u0061\u0073\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006frt\u0065\u0064")
		case "\u0072\u0069":
			_ceb.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020i\u006e\u0074\u0065\u006e\u0074\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0069":
			_ceb.Log.Debug("\u0046\u006c\u0061\u0074\u006e\u0065\u0073\u0073\u0020\u0074\u006f\u006c\u0065\u0072\u0061n\u0063e\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0067\u0073":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_cef, _eae := _df.GetName(_cce.Params[0])
			if !_eae {
				return _eba
			}
			if _cef == nil {
				return _ecd
			}
			_aag, _eae := _eff.GetExtGState(*_cef)
			if !_eae {
				_ceb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074 \u0066i\u006ed\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u003a\u0020\u0025\u0073", *_cef)
				return _ag.New("\u0072e\u0073o\u0075\u0072\u0063\u0065\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064")
			}
			_gda, _eae := _df.GetDict(_aag)
			if !_eae {
				_ceb.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020c\u006f\u0075\u006c\u0064 ge\u0074 g\u0072\u0061\u0070\u0068\u0069\u0063\u0073 s\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074")
				return _eba
			}
			_ceb.Log.Debug("G\u0053\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _gda.String())
		case "\u006d":
			if len(_cce.Params) != 2 {
				_ceb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _ecd)
				return nil
			}
			_dfb, _eeb := _df.GetNumbersAsFloat(_cce.Params)
			if _eeb != nil {
				return _eeb
			}
			_ceb.Log.Debug("M\u006f\u0076\u0065\u0020\u0074\u006f\u003a\u0020\u0025\u0076", _dfb)
			_edg.NewSubPath()
			_edg.MoveTo(_dfb[0], _dfb[1])
		case "\u006c":
			if len(_cce.Params) != 2 {
				_ceb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _ecd)
				return nil
			}
			_geg, _afa := _df.GetNumbersAsFloat(_cce.Params)
			if _afa != nil {
				return _afa
			}
			_edg.LineTo(_geg[0], _geg[1])
		case "\u0063":
			if len(_cce.Params) != 6 {
				return _ecd
			}
			_eaac, _efd := _df.GetNumbersAsFloat(_cce.Params)
			if _efd != nil {
				return _efd
			}
			_ceb.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _eaac)
			_edg.CubicTo(_eaac[0], _eaac[1], _eaac[2], _eaac[3], _eaac[4], _eaac[5])
		case "\u0076", "\u0079":
			if len(_cce.Params) != 4 {
				return _ecd
			}
			_fdb, _cgg := _df.GetNumbersAsFloat(_cce.Params)
			if _cgg != nil {
				return _cgg
			}
			_ceb.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _fdb)
			_edg.QuadraticTo(_fdb[0], _fdb[1], _fdb[2], _fdb[3])
		case "\u0068":
			_edg.ClosePath()
			_edg.NewSubPath()
		case "\u0072\u0065":
			if len(_cce.Params) != 4 {
				return _ecd
			}
			_gddb, _ga := _df.GetNumbersAsFloat(_cce.Params)
			if _ga != nil {
				return _ga
			}
			_edg.DrawRectangle(_gddb[0], _gddb[1], _gddb[2], _gddb[3])
			_edg.NewSubPath()
		case "\u0053":
			_gfg, _egd := _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
			if _egd != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _egd)
				return _egd
			}
			_edf, _fdcc := _gfg.(*_eg.PdfColorDeviceRGB)
			if !_fdcc {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _egd
			}
			_edg.SetRGBA(_edf.R(), _edf.G(), _edf.B(), 1)
			_edg.Stroke()
		case "\u0073":
			_bf, _ca := _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
			if _ca != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ca)
				return _ca
			}
			_defd, _cbb := _bf.(*_eg.PdfColorDeviceRGB)
			if !_cbb {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _ca
			}
			_edg.ClosePath()
			_edg.NewSubPath()
			_edg.SetRGBA(_defd.R(), _defd.G(), _defd.B(), 1)
			_edg.Stroke()
		case "\u0066", "\u0046":
			_dbc, _gfa := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _gfa != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gfa)
				return _gfa
			}
			switch _bde := _dbc.(type) {
			case *_eg.PdfColorDeviceRGB:
				_edg.SetRGBA(_bde.R(), _bde.G(), _bde.B(), 1)
				_edg.SetFillRule(_fe.FillRuleWinding)
				_edg.Fill()
			case *_eg.PdfColorPattern:
				_edg.Fill()
			}
			_ceb.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
		case "\u0066\u002a":
			_dcg, _cae := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _cae != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cae)
				return _cae
			}
			_eec, _fae := _dcg.(*_eg.PdfColorDeviceRGB)
			if !_fae {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _cae
			}
			_edg.SetRGBA(_eec.R(), _eec.G(), _eec.B(), 1)
			_edg.SetFillRule(_fe.FillRuleEvenOdd)
			_edg.Fill()
		case "\u0042":
			_aef, _bfb := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _bfb != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bfb)
				return _bfb
			}
			switch _aec := _aef.(type) {
			case *_eg.PdfColorDeviceRGB:
				_edg.SetRGBA(_aec.R(), _aec.G(), _aec.B(), 1)
				_edg.SetFillRule(_fe.FillRuleWinding)
				_edg.FillPreserve()
				_aef, _bfb = _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
				if _bfb != nil {
					_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bfb)
					return _bfb
				}
				if _efg, _aad := _aef.(*_eg.PdfColorDeviceRGB); _aad {
					_edg.SetRGBA(_efg.R(), _efg.G(), _efg.B(), 1)
					_edg.Stroke()
				}
			case *_eg.PdfColorPattern:
				_edg.SetFillRule(_fe.FillRuleWinding)
				_edg.Fill()
				_edg.StrokePattern()
			}
		case "\u0042\u002a":
			_fag, _cfd := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _cfd != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cfd)
				return _cfd
			}
			switch _bg := _fag.(type) {
			case *_eg.PdfColorDeviceRGB:
				_edg.SetRGBA(_bg.R(), _bg.G(), _bg.B(), 1)
				_edg.SetFillRule(_fe.FillRuleEvenOdd)
				_edg.FillPreserve()
				_fag, _cfd = _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
				if _cfd != nil {
					_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cfd)
					return _cfd
				}
				if _edc, _ac := _fag.(*_eg.PdfColorDeviceRGB); _ac {
					_edg.SetRGBA(_edc.R(), _edc.G(), _edc.B(), 1)
					_edg.Stroke()
				}
			case *_eg.PdfColorPattern:
				_edg.SetFillRule(_fe.FillRuleEvenOdd)
				_edg.Fill()
				_edg.StrokePattern()
			}
		case "\u0062":
			_edg.ClosePath()
			_afga, _afd := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _afd != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _afd)
				return _afd
			}
			switch _eca := _afga.(type) {
			case *_eg.PdfColorDeviceRGB:
				_edg.SetRGBA(_eca.R(), _eca.G(), _eca.B(), 1)
				_edg.NewSubPath()
				_edg.SetFillRule(_fe.FillRuleWinding)
				_edg.FillPreserve()
				_afga, _afd = _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
				if _afd != nil {
					_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _afd)
					return _afd
				}
				if _dbb, _bea := _afga.(*_eg.PdfColorDeviceRGB); _bea {
					_edg.SetRGBA(_dbb.R(), _dbb.G(), _dbb.B(), 1)
					_edg.Stroke()
				}
			case *_eg.PdfColorPattern:
				_edg.NewSubPath()
				_edg.SetFillRule(_fe.FillRuleWinding)
				_edg.Fill()
				_edg.StrokePattern()
			}
		case "\u0062\u002a":
			_edg.ClosePath()
			_efgb, _gc := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _gc != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gc)
				return _gc
			}
			switch _gfaa := _efgb.(type) {
			case *_eg.PdfColorDeviceRGB:
				_edg.SetRGBA(_gfaa.R(), _gfaa.G(), _gfaa.B(), 1)
				_edg.NewSubPath()
				_edg.SetFillRule(_fe.FillRuleEvenOdd)
				_edg.FillPreserve()
				_efgb, _gc = _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
				if _gc != nil {
					_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gc)
					return _gc
				}
				if _agg, _edgb := _efgb.(*_eg.PdfColorDeviceRGB); _edgb {
					_edg.SetRGBA(_agg.R(), _agg.G(), _agg.B(), 1)
					_edg.Stroke()
				}
			case *_eg.PdfColorPattern:
				_edg.NewSubPath()
				_edg.SetFillRule(_fe.FillRuleEvenOdd)
				_edg.Fill()
				_edg.StrokePattern()
			}
		case "\u006e":
			_edg.ClearPath()
		case "\u0057":
			_edg.SetFillRule(_fe.FillRuleWinding)
			_edg.ClipPreserve()
		case "\u0057\u002a":
			_edg.SetFillRule(_fe.FillRuleEvenOdd)
			_edg.ClipPreserve()
		case "\u0072\u0067":
			_bca, _bdc := _cbf.ColorNonStroking.(*_eg.PdfColorDeviceRGB)
			if !_bdc {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_edg.SetFillRGBA(_bca.R(), _bca.G(), _bca.B(), 1)
		case "\u0052\u0047":
			_caee, _cgf := _cbf.ColorStroking.(*_eg.PdfColorDeviceRGB)
			if !_cgf {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_edg.SetStrokeRGBA(_caee.R(), _caee.G(), _caee.B(), 1)
		case "\u006b":
			_dcb, _ba := _cbf.ColorNonStroking.(*_eg.PdfColorDeviceCMYK)
			if !_ba {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_abb, _gbf := _cbf.ColorspaceNonStroking.ColorToRGB(_dcb)
			if _gbf != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_gde, _ba := _abb.(*_eg.PdfColorDeviceRGB)
			if !_ba {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _abb)
				return nil
			}
			_edg.SetFillRGBA(_gde.R(), _gde.G(), _gde.B(), 1)
		case "\u004b":
			_bag, _dgf := _cbf.ColorStroking.(*_eg.PdfColorDeviceCMYK)
			if !_dgf {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_gaa, _faeb := _cbf.ColorspaceStroking.ColorToRGB(_bag)
			if _faeb != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_gdf, _dgf := _gaa.(*_eg.PdfColorDeviceRGB)
			if !_dgf {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gaa)
				return nil
			}
			_edg.SetStrokeRGBA(_gdf.R(), _gdf.G(), _gdf.B(), 1)
		case "\u0067":
			_gcd, _fff := _cbf.ColorNonStroking.(*_eg.PdfColorDeviceGray)
			if !_fff {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_fca, _fef := _cbf.ColorspaceNonStroking.ColorToRGB(_gcd)
			if _fef != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_eeg, _fff := _fca.(*_eg.PdfColorDeviceRGB)
			if !_fff {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fca)
				return nil
			}
			_edg.SetFillRGBA(_eeg.R(), _eeg.G(), _eeg.B(), 1)
		case "\u0047":
			_ebd, _ecf := _cbf.ColorStroking.(*_eg.PdfColorDeviceGray)
			if !_ecf {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_fgg, _fcab := _cbf.ColorspaceStroking.ColorToRGB(_ebd)
			if _fcab != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_dfc, _ecf := _fgg.(*_eg.PdfColorDeviceRGB)
			if !_ecf {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fgg)
				return nil
			}
			_edg.SetStrokeRGBA(_dfc.R(), _dfc.G(), _dfc.B(), 1)
		case "\u0063\u0073":
			if len(_cce.Params) > 0 {
				if _eecd, _egdf := _df.GetName(_cce.Params[0]); _egdf && _eecd.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					break
				}
			}
			_cec, _fed := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _fed != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_cbg, _eeca := _cec.(*_eg.PdfColorDeviceRGB)
			if !_eeca {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cec)
				return nil
			}
			_edg.SetFillRGBA(_cbg.R(), _cbg.G(), _cbg.B(), 1)
		case "\u0073\u0063":
			_deg, _efaf := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _efaf != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_fcf, _abbf := _deg.(*_eg.PdfColorDeviceRGB)
			if !_abbf {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _deg)
				return nil
			}
			_edg.SetFillRGBA(_fcf.R(), _fcf.G(), _fcf.B(), 1)
		case "\u0073\u0063\u006e":
			if len(_cce.Params) > 0 && len(_ceg.Params) > 0 {
				if _bcaf, _gbc := _df.GetName(_ceg.Params[0]); _gbc && _bcaf.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					if _cad, _ffb := _df.GetName(_cce.Params[0]); _ffb {
						_becb, _ged := _bcc.processGradient(_edg, _cce, _eff, _cad)
						if _ged != nil {
							_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0064\u0061\u0074a\u003a\u0020\u0025\u0076", _ged)
							break
						}
						if _becb == nil {
							_ceb.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
							break
						}
						_edg.SetFillStyle(_becb)
						_edg.SetStrokeStyle(_becb)
						break
					}
				}
			}
			_degd, _afdc := _cbf.ColorspaceNonStroking.ColorToRGB(_cbf.ColorNonStroking)
			if _afdc != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorNonStroking)
				return nil
			}
			_bda, _ggd := _degd.(*_eg.PdfColorDeviceRGB)
			if !_ggd {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _degd)
				return nil
			}
			_edg.SetFillRGBA(_bda.R(), _bda.G(), _bda.B(), 1)
		case "\u0043\u0053":
			if len(_cce.Params) > 0 {
				if _abbb, _fce := _df.GetName(_cce.Params[0]); _fce && _abbb.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					break
				}
			}
			_cac, _dbe := _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
			if _dbe != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_dde, _dag := _cac.(*_eg.PdfColorDeviceRGB)
			if !_dag {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cac)
				return nil
			}
			_edg.SetStrokeRGBA(_dde.R(), _dde.G(), _dde.B(), 1)
		case "\u0053\u0043":
			_ggde, _ddf := _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
			if _ddf != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_fcc, _bgb := _ggde.(*_eg.PdfColorDeviceRGB)
			if !_bgb {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ggde)
				return nil
			}
			_edg.SetStrokeRGBA(_fcc.R(), _fcc.G(), _fcc.B(), 1)
		case "\u0053\u0043\u004e":
			if len(_cce.Params) > 0 && len(_ceg.Params) > 0 {
				if _gfd, _beag := _df.GetName(_ceg.Params[0]); _beag && _gfd.String() == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
					if _eab, _efb := _df.GetName(_cce.Params[0]); _efb {
						_acg, _gbcb := _bcc.processGradient(_edg, _cce, _eff, _eab)
						if _gbcb != nil {
							_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0064\u0061\u0074a\u003a\u0020\u0025\u0076", _gbcb)
							break
						}
						if _acg == nil {
							_ceb.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
							break
						}
						_edg.SetFillStyle(_acg)
						_edg.SetStrokeStyle(_acg)
						break
					}
				}
			}
			_fad, _ddd := _cbf.ColorspaceStroking.ColorToRGB(_cbf.ColorStroking)
			if _ddd != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _cbf.ColorStroking)
				return nil
			}
			_bdb, _gee := _fad.(*_eg.PdfColorDeviceRGB)
			if !_gee {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _fad)
				return nil
			}
			_edg.SetStrokeRGBA(_bdb.R(), _bdb.G(), _bdb.B(), 1)
		case "\u0073\u0068":
			if len(_cce.Params) != 1 {
				_ceb.Log.Debug("\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0020\u0070\u0061r\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
				break
			}
			_aab, _ccf := _df.GetName(_cce.Params[0])
			if !_ccf {
				_ceb.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020g\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006eg\u0020n\u0061\u006d\u0065")
				break
			}
			_acd, _ccf := _eff.GetShadingByName(*_aab)
			if !_ccf {
				_ceb.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020g\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006eg\u0020d\u0061\u0074\u0061")
				break
			}
			_gfdf, _gba, _bcd := _bcc.processShading(_edg, _acd)
			if _bcd != nil {
				_ceb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0065\u006e\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0073\u0068a\u0064\u0069\u006e\u0067\u0020d\u0061\u0074a\u003a\u0020\u0025\u0076", _bcd)
				break
			}
			if _gfdf == nil {
				_ceb.Log.Debug("\u0055\u006ek\u006e\u006f\u0077n\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074")
				break
			}
			_afdf, _bcd := _gba.ToFloat64Array()
			if _bcd != nil {
				_ceb.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073: \u0025\u0076", _bcd)
				break
			}
			_edg.DrawRectangle(_afdf[0], _afdf[1], _afdf[2], _afdf[3])
			_edg.NewSubPath()
			_edg.SetFillStyle(_gfdf)
			_edg.SetStrokeStyle(_gfdf)
			_edg.Fill()
		case "\u0044\u006f":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_ccd, _dfg := _df.GetName(_cce.Params[0])
			if !_dfg {
				return _eba
			}
			_, _dcee := _eff.GetXObjectByName(*_ccd)
			switch _dcee {
			case _eg.XObjectTypeImage:
				_ceb.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067e\u003a\u0020\u0025\u0073", _ccd.String())
				_aee, _cbbd := _eff.GetXObjectImageByName(*_ccd)
				if _cbbd != nil {
					return _cbbd
				}
				_gfb, _cbbd := _aee.ToImage()
				if _cbbd != nil {
					_ceb.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _cbbd)
					return nil
				}
				if _gbg := _aee.ColorSpace; _gbg != nil {
					var _bfg bool
					switch _gbg.(type) {
					case *_eg.PdfColorspaceSpecialIndexed:
						_bfg = true
					}
					if _bfg {
						if _fab, _abd := _gbg.ImageToRGB(*_gfb); _abd != nil {
							_ceb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006fnv\u0065r\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0052G\u0042\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020i\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
						} else {
							_gfb = &_fab
						}
					}
				}
				_ffe := _edg.FillPattern().ColorAt(0, 0)
				var _gef _ec.Image
				if _aee.Mask != nil {
					if _gef, _cbbd = _gff(_aee.Mask, _ffe); _cbbd != nil {
						_ceb.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				} else if _aee.SMask != nil {
					if _gef, _cbbd = _edac(_aee.SMask, _ffe); _cbbd != nil {
						_ceb.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				}
				var _eaef _ec.Image
				if _dcgb, _ := _df.GetBoolVal(_aee.ImageMask); _dcgb {
					_eaef = _ddb(_gfb, _ffe)
				} else {
					_eaef, _cbbd = _gfb.ToGoImage()
					if _cbbd != nil {
						_ceb.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _cbbd)
						return nil
					}
				}
				if _gef != nil {
					_eaef = _eeee(_eaef, _gef)
				}
				_edad := _eaef.Bounds()
				_edg.Push()
				_edg.Scale(1.0/float64(_edad.Dx()), -1.0/float64(_edad.Dy()))
				_edg.DrawImageAnchored(_eaef, 0, 0, 0, 1)
				_edg.Pop()
			case _eg.XObjectTypeForm:
				_ceb.Log.Debug("\u0058\u004fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u0073", _ccd.String())
				_dcea, _dcbb := _eff.GetXObjectFormByName(*_ccd)
				if _dcbb != nil {
					return _dcbb
				}
				_bbbc, _dcbb := _dcea.GetContentStream()
				if _dcbb != nil {
					return _dcbb
				}
				_ega := _dcea.Resources
				if _ega == nil {
					_ega = _eff
				}
				_edg.Push()
				if _dcea.Matrix != nil {
					_dgba, _ddfe := _df.GetArray(_dcea.Matrix)
					if !_ddfe {
						return _eba
					}
					_daab, _gaac := _df.GetNumbersAsFloat(_dgba.Elements())
					if _gaac != nil {
						return _gaac
					}
					if len(_daab) != 6 {
						return _ecd
					}
					_gca := _cb.NewMatrix(_daab[0], _daab[1], _daab[2], _daab[3], _daab[4], _daab[5])
					_edg.SetMatrix(_edg.Matrix().Mult(_gca))
				}
				if _dcea.BBox != nil {
					_effd, _ccb := _df.GetArray(_dcea.BBox)
					if !_ccb {
						return _eba
					}
					_agga, _ccbe := _df.GetNumbersAsFloat(_effd.Elements())
					if _ccbe != nil {
						return _ccbe
					}
					if len(_agga) != 4 {
						_ceb.Log.Debug("\u004c\u0065\u006e\u0020\u003d\u0020\u0025\u0064", len(_agga))
						return _ecd
					}
					_edg.DrawRectangle(_agga[0], _agga[1], _agga[2]-_agga[0], _agga[3]-_agga[1])
					_edg.SetRGBA(1, 0, 0, 1)
					_edg.Clip()
				} else {
					_ceb.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0052\u0065q\u0075\u0069\u0072e\u0064\u0020\u0042\u0042\u006f\u0078\u0020\u006d\u0069ss\u0069\u006e\u0067 \u006f\u006e \u0058\u004f\u0062\u006a\u0065\u0063t\u0020\u0046o\u0072\u006d")
				}
				_dcbb = _bcc.renderContentStream(_edg, string(_bbbc), _ega)
				if _dcbb != nil {
					return _dcbb
				}
				_edg.Pop()
			}
		case "\u0042\u0049":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_egb, _cga := _cce.Params[0].(*_gd.ContentStreamInlineImage)
			if !_cga {
				return nil
			}
			_bdd, _bee := _egb.ToImage(_eff)
			if _bee != nil {
				_ceb.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bee)
				return nil
			}
			_acga, _bee := _bdd.ToGoImage()
			if _bee != nil {
				_ceb.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u006day\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065.\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bee)
				return nil
			}
			_gcg := _acga.Bounds()
			_edg.Push()
			_edg.Scale(1.0/float64(_gcg.Dx()), -1.0/float64(_gcg.Dy()))
			_edg.DrawImageAnchored(_acga, 0, 0, 0, 1)
			_edg.Pop()
		case "\u0042\u0054":
			_eag.Reset()
		case "\u0045\u0054":
			_eag.Reset()
		case "\u0054\u0072":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_fcfg, _gdbd := _df.GetNumberAsFloat(_cce.Params[0])
			if _gdbd != nil {
				return _gdbd
			}
			_eag.Tr = _fe.TextRenderingMode(_fcfg)
		case "\u0054\u004c":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_faf, _eea := _df.GetNumberAsFloat(_cce.Params[0])
			if _eea != nil {
				return _eea
			}
			_eag.Tl = _faf
		case "\u0054\u0063":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_bff, _dfd := _df.GetNumberAsFloat(_cce.Params[0])
			if _dfd != nil {
				return _dfd
			}
			_ceb.Log.Debug("\u0054\u0063\u003a\u0020\u0025\u0076", _bff)
			_eag.Tc = _bff
		case "\u0054\u0077":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_fde, _ace := _df.GetNumberAsFloat(_cce.Params[0])
			if _ace != nil {
				return _ace
			}
			_ceb.Log.Debug("\u0054\u0077\u003a\u0020\u0025\u0076", _fde)
			_eag.Tw = _fde
		case "\u0054\u007a":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_cgaa, _dge := _df.GetNumberAsFloat(_cce.Params[0])
			if _dge != nil {
				return _dge
			}
			_eag.Th = _cgaa
		case "\u0054\u0073":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_eef, _ecgd := _df.GetNumberAsFloat(_cce.Params[0])
			if _ecgd != nil {
				return _ecgd
			}
			_eag.Ts = _eef
		case "\u0054\u0064":
			if len(_cce.Params) != 2 {
				return _ecd
			}
			_dgd, _adb := _df.GetNumbersAsFloat(_cce.Params)
			if _adb != nil {
				return _adb
			}
			_ceb.Log.Debug("\u0054\u0064\u003a\u0020\u0025\u0076", _dgd)
			_eag.ProcTd(_dgd[0], _dgd[1])
		case "\u0054\u0044":
			if len(_cce.Params) != 2 {
				return _ecd
			}
			_gea, _bgbf := _df.GetNumbersAsFloat(_cce.Params)
			if _bgbf != nil {
				return _bgbf
			}
			_ceb.Log.Debug("\u0054\u0044\u003a\u0020\u0025\u0076", _gea)
			_eag.ProcTD(_gea[0], _gea[1])
		case "\u0054\u002a":
			_eag.ProcTStar()
		case "\u0054\u006d":
			if len(_cce.Params) != 6 {
				return _ecd
			}
			_bgc, _dba := _df.GetNumbersAsFloat(_cce.Params)
			if _dba != nil {
				return _dba
			}
			_ceb.Log.Debug("\u0054\u0065x\u0074\u0020\u006da\u0074\u0072\u0069\u0078\u003a\u0020\u0025\u002b\u0076", _bgc)
			_eag.ProcTm(_bgc[0], _bgc[1], _bgc[2], _bgc[3], _bgc[4], _bgc[5])
		case "\u0027":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_abbg, _eee := _df.GetStringBytes(_cce.Params[0])
			if !_eee {
				return _eba
			}
			_ceb.Log.Debug("\u0027\u0020\u0073t\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_abbg))
			_eag.ProcQ(_abbg, _edg)
		case "\u0022":
			if len(_cce.Params) != 3 {
				return _ecd
			}
			_bffa, _fbg := _df.GetNumberAsFloat(_cce.Params[0])
			if _fbg != nil {
				return _fbg
			}
			_eeac, _fbg := _df.GetNumberAsFloat(_cce.Params[1])
			if _fbg != nil {
				return _fbg
			}
			_caf, _eegd := _df.GetStringBytes(_cce.Params[2])
			if !_eegd {
				return _eba
			}
			_eag.ProcDQ(_caf, _bffa, _eeac, _edg)
		case "\u0054\u006a":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_cde, _dgbe := _df.GetStringBytes(_cce.Params[0])
			if !_dgbe {
				return _eba
			}
			_ceb.Log.Debug("\u0054j\u0020s\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0060\u0025\u0073\u0060", string(_cde))
			_eag.ProcTj(_cde, _edg)
		case "\u0054\u004a":
			if len(_cce.Params) != 1 {
				return _ecd
			}
			_dca, _ggf := _df.GetArray(_cce.Params[0])
			if !_ggf {
				_ceb.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _dca)
				return _eba
			}
			_ceb.Log.Debug("\u0054\u004a\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u002b\u0076", _dca)
			for _, _adgd := range _dca.Elements() {
				switch _eebb := _adgd.(type) {
				case *_df.PdfObjectString:
					if _eebb != nil {
						_eag.ProcTj(_eebb.Bytes(), _edg)
					}
				case *_df.PdfObjectFloat, *_df.PdfObjectInteger:
					_fcg, _fagf := _df.GetNumberAsFloat(_eebb)
					if _fagf == nil {
						_eag.Translate(-_fcg*0.001*_eag.Tf.Size*_eag.Th/100.0, 0)
					}
				}
			}
		case "\u0054\u0066":
			if len(_cce.Params) != 2 {
				return _ecd
			}
			_ceb.Log.Debug("\u0025\u0023\u0076", _cce.Params)
			_beb, _eege := _df.GetName(_cce.Params[0])
			if !_eege || _beb == nil {
				_ceb.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006e\u0061m\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _cce.Params[0])
				return _eba
			}
			_ceb.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073", _beb.String())
			_bcg, _agb := _df.GetNumberAsFloat(_cce.Params[1])
			if _agb != nil {
				_ceb.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0069z\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _cce.Params[1])
				return _eba
			}
			_ceb.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0076", _bcg)
			_cfda, _ada := _eff.GetFontByName(*_beb)
			if !_ada {
				_ceb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0025s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _beb.String())
				return _ag.New("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
			}
			_ceb.Log.Debug("\u0046\u006f\u006e\u0074\u003a\u0020\u0025\u0054", _cfda)
			_dgg, _eege := _df.GetDict(_cfda)
			if !_eege {
				_ceb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0067e\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074")
				return _eba
			}
			_gfc, _agb := _eg.NewPdfFontFromPdfObject(_dgg)
			if _agb != nil {
				_ceb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
				return _agb
			}
			_efgd := _gfc.BaseFont()
			if _efgd == "" {
				_efgd = _beb.String()
			}
			_abg, _eege := _gg[_efgd]
			if !_eege {
				_abg, _agb = _fe.NewTextFont(_gfc, _bcg)
				if _agb != nil {
					_ceb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _agb)
				}
			}
			if _abg == nil {
				if len(_efgd) > 7 && _efgd[6] == '+' {
					_efgd = _efgd[7:]
				}
				_fdg := []string{_efgd, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
				for _, _fgc := range _fdg {
					_ceb.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _fgc)
					if _abg, _eege = _gg[_fgc]; _eege {
						break
					}
					_baf := _ad.Match(_fgc)
					if _baf == nil {
						_ceb.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _fgc)
						continue
					}
					_abg, _agb = _fe.NewTextFontFromPath(_baf.Filename, _bcg)
					if _agb != nil {
						_ceb.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _baf.Filename)
						continue
					}
					_ceb.Log.Debug("\u0053\u0075\u0062\u0073\u0074\u0069t\u0075\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073 \u0077\u0069\u0074\u0068\u0020\u0025\u0073 \u0028\u0025\u0073\u0029", _efgd, _baf.Name, _baf.Filename)
					_gg[_fgc] = _abg
					break
				}
			}
			if _abg == nil {
				_ceb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020n\u006f\u0074\u0020\u0066\u0069\u006ed\u0020\u0061\u006e\u0079\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065 \u0066\u006f\u006e\u0074")
				return _ag.New("\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0066\u0069\u006e\u0064\u0020a\u006ey\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u006e\u0074")
			}
			_eag.ProcTf(_abg.WithSize(_bcg, _gfc))
		case "\u0042\u004d\u0043", "\u0042\u0044\u0043":
		case "\u0045\u004d\u0043":
		default:
			_ceb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u006f\u0070\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u0073", _cce.Operand)
		}
		_ceg = _cce
		return nil
	})
	_abe = _dg.Process(_ea)
	if _abe != nil {
		return _abe
	}
	return nil
}

var (
	_eba = _ag.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_ecd = _ag.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

// PdfShadingType defines PDF shading types.
// Source: PDF32000_2008.pdf. Chapter 8.7.4.5
type PdfShadingType int64

func (_fbgg renderer) processLinearShading(_gaad _fe.Context, _edcc *_eg.PdfShading) (_fe.Gradient, *_df.PdfObjectArray, error) {
	_cge := _edcc.GetContext().(*_eg.PdfShadingType2)
	if len(_cge.Function) == 0 {
		return nil, nil, _ag.New("\u006e\u006f\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074 \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0073\u006b\u0069\u0070\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	_fac, _bccg := _cge.Coords.ToFloat64Array()
	if _bccg != nil {
		return nil, nil, _ag.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020\u0073h\u0061\u0064\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	_daaa := _edcc.ColorSpace
	_afgc, _dgc := _gaad.Matrix().Transform(_fac[0], _fac[1])
	_cada, _ffbb := _gaad.Matrix().Transform(_fac[2], _fac[3])
	_gfga := _b.NewLinearGradient(_afgc, _dgc, _cada, _ffbb)
	_aba := _df.MakeArrayFromFloats([]float64{0, 0, 1, 1})
	for _, _aebc := range _fac {
		if _aebc > 1 {
			_aba = _cge.Coords
			break
		}
	}
	if _cefa, _acb := _cge.Function[0].(*_eg.PdfFunctionType2); _acb {
		_gfga, _bccg = _cgb(_gfga, _cefa, _daaa, 1.0, true)
	} else if _fbf, _dcf := _cge.Function[0].(*_eg.PdfFunctionType3); _dcf {
		_acf := append([]float64{0}, _fbf.Bounds...)
		_acf = append(_acf, 1.0)
		_gfga, _bccg = _efba(_gfga, _fbf, _daaa, _acf)
	}
	return _gfga, _aba, _bccg
}

func _cgb(_bgd _fe.Gradient, _abgc *_eg.PdfFunctionType2, _eeae _eg.PdfColorspace, _aeba float64, _bdde bool) (_fe.Gradient, error) {
	switch _eeae.(type) {
	case *_eg.PdfColorspaceDeviceRGB:
		if len(_abgc.C0) != 3 || len(_abgc.C1) != 3 {
			return nil, _ag.New("\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0052\u0047\u0042\u0020\u0063o\u006co\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
		}
		_edaf := _abgc.C0
		_cegd := _abgc.C1
		if _bdde {
			_bgd.AddColorStop(0.0, _eb.RGBA{R: uint8(_edaf[0] * 255), G: uint8(_edaf[1] * 255), B: uint8(_edaf[2] * 255), A: 255})
		}
		_bgd.AddColorStop(_aeba, _eb.RGBA{R: uint8(_cegd[0] * 255), G: uint8(_cegd[1] * 255), B: uint8(_cegd[2] * 255), A: 255})
	case *_eg.PdfColorspaceDeviceCMYK:
		if len(_abgc.C0) != 4 || len(_abgc.C1) != 4 {
			return nil, _ag.New("\u0069\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072 \u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
		}
		_eecg := _abgc.C0
		_egf := _abgc.C1
		if _bdde {
			_bgd.AddColorStop(0.0, _eb.CMYK{C: uint8(_eecg[0] * 255), M: uint8(_eecg[1] * 255), Y: uint8(_eecg[2] * 255), K: uint8(_eecg[3] * 255)})
		}
		_bgd.AddColorStop(_aeba, _eb.CMYK{C: uint8(_egf[0] * 255), M: uint8(_egf[1] * 255), Y: uint8(_egf[2] * 255), K: uint8(_egf[3] * 255)})
	default:
		return nil, _cd.Errorf("u\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072 \u0073\u0070\u0061c\u0065:\u0020\u0025\u0073", _eeae.String())
	}
	return _bgd, nil
}

func _eeee(_gedb, _gabb _ec.Image) _ec.Image {
	_ebe, _ead := _gabb.Bounds().Size(), _gedb.Bounds().Size()
	_ffcg, _cegc := _ebe.X, _ebe.Y
	if _ead.X > _ffcg {
		_ffcg = _ead.X
	}
	if _ead.Y > _cegc {
		_cegc = _ead.Y
	}
	_ebeb := _ec.Rect(0, 0, _ffcg, _cegc)
	if _ebe.X != _ffcg || _ebe.Y != _cegc {
		_cfb := _ec.NewRGBA(_ebeb)
		_fa.BiLinear.Scale(_cfb, _ebeb, _gedb, _gabb.Bounds(), _fa.Over, nil)
		_gabb = _cfb
	}
	if _ead.X != _ffcg || _ead.Y != _cegc {
		_bfa := _ec.NewRGBA(_ebeb)
		_fa.BiLinear.Scale(_bfa, _ebeb, _gedb, _gedb.Bounds(), _fa.Over, nil)
		_gedb = _bfa
	}
	_eeec := _ec.NewRGBA(_ebeb)
	_fa.DrawMask(_eeec, _ebeb, _gedb, _ec.Point{}, _gabb, _ec.Point{}, _fa.Over)
	return _eeec
}

// ImageDevice is used to render PDF pages to image targets.
type ImageDevice struct {
	renderer

	// OutputWidth represents the width of the rendered images in pixels.
	// The heights of the output images are calculated based on the selected
	// width and the original height of each rendered page.
	OutputWidth int
}

func _cafb(_dcd *_eg.Image, _cgge _eb.Color) _ec.Image {
	_abf, _agc := int(_dcd.Width), int(_dcd.Height)
	_efbc := _ec.NewRGBA(_ec.Rect(0, 0, _abf, _agc))
	for _gbbbd := 0; _gbbbd < _agc; _gbbbd++ {
		for _aac := 0; _aac < _abf; _aac++ {
			_bbdg, _ecdg := _dcd.ColorAt(_aac, _gbbbd)
			if _ecdg != nil {
				_ceb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _aac, _gbbbd)
				continue
			}
			_dga, _ede, _gcdg, _ := _bbdg.RGBA()
			var _gddd _eb.Color
			if _dga+_ede+_gcdg == 0 {
				_gddd = _eb.Transparent
			} else {
				_gddd = _cgge
			}
			_efbc.Set(_aac, _gbbbd, _gddd)
		}
	}
	return _efbc
}

func _gff(_bcb _df.PdfObject, _gded _eb.Color) (_ec.Image, error) {
	_gbe, _edda := _df.GetStream(_bcb)
	if !_edda {
		return nil, nil
	}
	_efbd, _fgd := _eg.NewXObjectImageFromStream(_gbe)
	if _fgd != nil {
		return nil, _fgd
	}
	_dbfe, _fgd := _efbd.ToImage()
	if _fgd != nil {
		return nil, _fgd
	}
	return _ddb(_dbfe, _gded), nil
}

func (_gcab renderer) processShading(_cggb _fe.Context, _cbbf *_eg.PdfShading) (_fe.Gradient, *_df.PdfObjectArray, error) {
	_dfda := int64(*_cbbf.ShadingType)
	if _dfda == int64(ShadingTypeAxial) {
		return _gcab.processLinearShading(_cggb, _cbbf)
	} else if _dfda == int64(ShadingTypeRadial) {
		return _gcab.processRadialShading(_cggb, _cbbf)
	} else {
		_ceb.Log.Debug(_cd.Sprintf("\u0050r\u006f\u0063e\u0073\u0073\u0069n\u0067\u0020\u0067\u0072\u0061\u0064\u0069e\u006e\u0074\u0020\u0074\u0079\u0070e\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074 \u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064", _dfda))
	}
	return nil, nil, nil
}

func _ddb(_gefd *_eg.Image, _ebc _eb.Color) _ec.Image {
	_fdda, _acbb := int(_gefd.Width), int(_gefd.Height)
	_beda := _ec.NewRGBA(_ec.Rect(0, 0, _fdda, _acbb))
	for _ebba := 0; _ebba < _acbb; _ebba++ {
		for _aedf := 0; _aedf < _fdda; _aedf++ {
			_fbdc, _gbae := _gefd.ColorAt(_aedf, _ebba)
			if _gbae != nil {
				_ceb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _aedf, _ebba)
				continue
			}
			_ddc, _gga, _ebcc, _ := _fbdc.RGBA()
			var _cgbd _eb.Color
			if _ddc+_gga+_ebcc == 0 {
				_cgbd = _ebc
			} else {
				_cgbd = _eb.Transparent
			}
			_beda.Set(_aedf, _ebba, _cgbd)
		}
	}
	return _beda
}

func (_dcc renderer) renderPage(_aff _fe.Context, _cf *_eg.PdfPage, _bbd _cb.Matrix, _cgd bool) error {
	if !_cgd {
		_dce := _eg.FieldFlattenOpts{AnnotFilterFunc: func(_bbc *_eg.PdfAnnotation) bool {
			switch _bbc.GetContext().(type) {
			case *_eg.PdfAnnotationLine:
				return true
			case *_eg.PdfAnnotationSquare:
				return true
			case *_eg.PdfAnnotationCircle:
				return true
			case *_eg.PdfAnnotationPolygon:
				return true
			case *_eg.PdfAnnotationPolyLine:
				return true
			}
			return false
		}}
		_cfe := _ef.FieldAppearance{}
		_bbb := _cf.FlattenFieldsWithOpts(_cfe, &_dce)
		if _bbb != nil {
			_ceb.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0064u\u0072\u0069n\u0067\u0020\u0061\u006e\u006e\u006f\u0074\u0061t\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0074\u0074\u0065\u006e\u0069n\u0067\u0020\u0025\u0076", _bbb)
		}
	}
	_cff, _cc := _cf.GetAllContentStreams()
	if _cc != nil {
		return _cc
	}
	if _gbb := _bbd; !_gbb.Identity() {
		_cff = _cd.Sprintf("%\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066 \u0025\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066\u0020\u0063\u006d\u0020%\u0073", _gbb[0], _gbb[1], _gbb[3], _gbb[4], _gbb[6], _gbb[7], _cff)
	}
	_aff.Translate(0, float64(_aff.Height()))
	_aff.Scale(1, -1)
	_aff.Push()
	_aff.SetRGBA(1, 1, 1, 1)
	_aff.DrawRectangle(0, 0, float64(_aff.Width()), float64(_aff.Height()))
	_aff.Fill()
	_aff.Pop()
	_aff.SetLineWidth(1.0)
	_aff.SetRGBA(0, 0, 0, 1)
	return _dcc.renderContentStream(_aff, _cff, _cf.Resources)
}

// RenderToPath converts the specified PDF page into an image and saves the
// result at the specified location.
func (_dc *ImageDevice) RenderToPath(page *_eg.PdfPage, outputPath string) error {
	_gb, _eda := _dc.Render(page)
	if _eda != nil {
		return _eda
	}
	_bdg := _d.ToLower(_g.Ext(outputPath))
	if _bdg == "" {
		return _ag.New("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065cog\u006eiz\u0065\u0020\u006f\u0075\u0074\u0070\u0075t \u0066\u0069\u006c\u0065\u0020\u0074\u0079p\u0065")
	}
	switch _bdg {
	case "\u002e\u0070\u006e\u0067":
		return _aed(outputPath, _gb)
	case "\u002e\u006a\u0070\u0067", "\u002e\u006a\u0070e\u0067":
		return _fdd(outputPath, _gb, 100)
	}
	return _cd.Errorf("\u0075\u006e\u0072\u0065\u0063\u006fg\u006e\u0069\u007a\u0065\u0064\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020f\u0069\u006c\u0065\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _bdg)
}

func _efba(_gcgb _fe.Gradient, _eabb *_eg.PdfFunctionType3, _ffc _eg.PdfColorspace, _afgag []float64) (_fe.Gradient, error) {
	var _aea error
	for _feb := 0; _feb < len(_eabb.Functions); _feb++ {
		if _afbg, _aaf := _eabb.Functions[_feb].(*_eg.PdfFunctionType2); _aaf {
			_gcgb, _aea = _cgb(_gcgb, _afbg, _ffc, _afgag[_feb+1], _feb == 0)
			if _aea != nil {
				return nil, _aea
			}
		}
	}
	return _gcgb, nil
}

func _edac(_dccb _df.PdfObject, _aeca _eb.Color) (_ec.Image, error) {
	_ggfa, _fffe := _df.GetStream(_dccb)
	if !_fffe {
		return nil, nil
	}
	_dgef, _dcab := _eg.NewXObjectImageFromStream(_ggfa)
	if _dcab != nil {
		return nil, _dcab
	}
	_dad, _dcab := _dgef.ToImage()
	if _dcab != nil {
		return nil, _dcab
	}
	return _cafb(_dad, _aeca), nil
}

// RenderWithOpts converts the specified PDF page into an image, optionally flattens annotations and returns the result.
func (_fd *ImageDevice) RenderWithOpts(page *_eg.PdfPage, skipFlattening bool) (_ec.Image, error) {
	_bd, _ab := page.GetMediaBox()
	if _ab != nil {
		return nil, _ab
	}
	_bd.Normalize()
	_fb := page.CropBox
	var _fba, _edd float64
	if _fb != nil {
		_fb.Normalize()
		_fba, _edd = _fb.Width(), _fb.Height()
	}
	_fg := page.Rotate
	_be, _dd, _af, _bb := _bd.Llx, _bd.Lly, _bd.Width(), _bd.Height()
	_ge := _cb.IdentityMatrix()
	if _fg != nil && *_fg%360 != 0 && *_fg%90 == 0 {
		_fea := -float64(*_fg)
		_bec := _aaa(_af, _bb, _fea)
		_ge = _ge.Translate((_bec.Width-_af)/2+_af/2, (_bec.Height-_bb)/2+_bb/2).Rotate(_fea*_db.Pi/180).Translate(-_af/2, -_bb/2)
		_af, _bb = _bec.Width, _bec.Height
		if _fb != nil {
			_cee := _aaa(_fba, _edd, _fea)
			_fba, _edd = _cee.Width, _cee.Height
		}
	}
	if _be != 0 || _dd != 0 {
		_ge = _ge.Translate(-_be, -_dd)
	}
	_fd._fdc = 1.0
	if _fd.OutputWidth != 0 {
		_gdb := _af
		if _fb != nil {
			_gdb = _fba
		}
		_fd._fdc = float64(_fd.OutputWidth) / _gdb
		_af, _bb, _fba, _edd = _af*_fd._fdc, _bb*_fd._fdc, _fba*_fd._fdc, _edd*_fd._fdc
		_ge = _cb.ScaleMatrix(_fd._fdc, _fd._fdc).Mult(_ge)
	}
	_cg := _b.NewContext(int(_af), int(_bb))
	if _ecg := _fd.renderPage(_cg, page, _ge, skipFlattening); _ecg != nil {
		return nil, _ecg
	}
	_bc := _cg.Image()
	if _fb != nil {
		_cebg, _fgf := (_fb.Llx-_be)*_fd._fdc, (_fb.Lly-_dd)*_fd._fdc
		_fbe := _ec.Rect(0, 0, int(_fba), int(_edd))
		_de := _ec.Pt(int(_cebg), int(_bb-_fgf-_edd))
		_gf := _ec.NewRGBA(_fbe)
		_e.Draw(_gf, _fbe, _bc, _de, _e.Src)
		_bc = _gf
	}
	return _bc, nil
}

func (_fabc renderer) processRadialShading(_gdbe _fe.Context, _abdg *_eg.PdfShading) (_fe.Gradient, *_df.PdfObjectArray, error) {
	_bed := _abdg.GetContext().(*_eg.PdfShadingType3)
	if len(_bed.Function) == 0 {
		return nil, nil, _ag.New("\u006e\u006f\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074 \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0073\u006b\u0069\u0070\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	_gac, _aeg := _bed.Coords.ToFloat64Array()
	if _aeg != nil {
		return nil, nil, _ag.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020\u0073h\u0061\u0064\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	_bfbg := _abdg.ColorSpace
	_afb := _df.MakeArrayFromFloats([]float64{0, 0, 1, 1})
	var _ffa, _gab, _ebg, _beee, _ebf, _dbec float64
	_ffa, _gab = _gdbe.Matrix().Transform(_gac[0], _gac[1])
	_ebg, _beee = _gdbe.Matrix().Transform(_gac[3], _gac[4])
	_ebf, _ = _gdbe.Matrix().Transform(_gac[2], 0)
	_dbec, _ = _gdbe.Matrix().Transform(_gac[5], 0)
	_effc, _ := _gdbe.Matrix().Translation()
	_ebf -= _effc
	_dbec -= _effc
	for _bba, _cdg := range _gac {
		if _bba == 2 || _bba == 5 {
			continue
		}
		if _cdg > 1.0 {
			_gfcd := _db.Min(_ffa-_ebf, _ebg-_dbec)
			_bccgg := _db.Min(_gab-_ebf, _beee-_dbec)
			_geag := _db.Max(_ffa+_ebf, _ebg+_dbec)
			_ggfb := _db.Max(_gab+_ebf, _beee+_dbec)
			_cege := _geag - _gfcd
			_fcgc := _bccgg - _ggfb
			_afb = _df.MakeArrayFromFloats([]float64{_gfcd, _bccgg, _cege, _fcgc})
			break
		}
	}
	_gfda := _b.NewRadialGradient(_ffa, _gab, _ebf, _ebg, _beee, _dbec)
	if _abgd, _ccc := _bed.Function[0].(*_eg.PdfFunctionType2); _ccc {
		_gfda, _aeg = _cgb(_gfda, _abgd, _bfbg, 1.0, true)
	} else if _ddfeg, _fcgf := _bed.Function[0].(*_eg.PdfFunctionType3); _fcgf {
		_dbf := append([]float64{0}, _ddfeg.Bounds...)
		_dbf = append(_dbf, 1.0)
		_gfda, _aeg = _efba(_gfda, _ddfeg, _bfbg, _dbf)
	}
	if _aeg != nil {
		return nil, nil, _aeg
	}
	return _gfda, _afb, nil
}

func _aed(_cgde string, _fadc _ec.Image) error {
	_eac, _ccbeg := _c.Create(_cgde)
	if _ccbeg != nil {
		return _ccbeg
	}
	defer _eac.Close()
	return _ed.Encode(_eac, _fadc)
}

// NewImageDevice returns a new image device.
func NewImageDevice() *ImageDevice {
	return &ImageDevice{}
}

type renderer struct{ _fdc float64 }
