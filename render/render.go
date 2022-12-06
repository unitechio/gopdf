package render

import (
	_a "errors"
	_d "fmt"
	_gb "image"
	_b "image/color"
	_cg "image/draw"
	_eb "image/jpeg"
	_ec "image/png"
	_g "math"
	_c "os"
	_cc "path/filepath"
	_ed "strings"

	_bd "bitbucket.org/shenghui0779/gopdf/common"
	_cgg "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ce "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_ee "bitbucket.org/shenghui0779/gopdf/core"
	_af "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_eg "bitbucket.org/shenghui0779/gopdf/model"
	_ae "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_bc "bitbucket.org/shenghui0779/gopdf/render/internal/context/imagerender"
	_ca "github.com/adrg/sysfont"
	_bg "golang.org/x/image/draw"
)

func _efd(_fff _ee.PdfObject, _dec _b.Color) (_gb.Image, error) {
	_gac, _bcg := _ee.GetStream(_fff)
	if !_bcg {
		return nil, nil
	}
	_fefb, _afd := _eg.NewXObjectImageFromStream(_gac)
	if _afd != nil {
		return nil, _afd
	}
	_afaf, _afd := _fefb.ToImage()
	if _afd != nil {
		return nil, _afd
	}
	return _eeec(_afaf, _dec), nil
}

type renderer struct{ _gba float64 }

var (
	_ef = _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_ff = _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

// RenderToPath converts the specified PDF page into an image and saves the
// result at the specified location.
func (_dd *ImageDevice) RenderToPath(page *_eg.PdfPage, outputPath string) error {
	_fae, _fgf := _dd.Render(page)
	if _fgf != nil {
		return _fgf
	}
	_bf := _ed.ToLower(_cc.Ext(outputPath))
	if _bf == "" {
		return _a.New("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065cog\u006eiz\u0065\u0020\u006f\u0075\u0074\u0070\u0075t \u0066\u0069\u006c\u0065\u0020\u0074\u0079p\u0065")
	}
	switch _bf {
	case "\u002e\u0070\u006e\u0067":
		return _bcd(outputPath, _fae)
	case "\u002e\u006a\u0070\u0067", "\u002e\u006a\u0070e\u0067":
		return _ffdg(outputPath, _fae, 100)
	}
	return _d.Errorf("\u0075\u006e\u0072\u0065\u0063\u006fg\u006e\u0069\u007a\u0065\u0064\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020f\u0069\u006c\u0065\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _bf)
}
func (_bdg renderer) renderContentStream(_afa _ae.Context, _be string, _bbb *_eg.PdfPageResources) error {
	_ccd, _bfa := _cgg.NewContentStreamParser(_be).Parse()
	if _bfa != nil {
		return _bfa
	}
	_ba := _afa.TextState()
	_ba.GlobalScale = _bdg._gba
	_aba := map[string]*_ae.TextFont{}
	_cdcg := _ca.NewFinder(&_ca.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066", "\u002e\u0074\u0074\u0063"}})
	_edg := _cgg.NewContentStreamProcessor(*_ccd)
	_edg.AddHandler(_cgg.HandlerConditionEnumAllOperands, "", func(_bec *_cgg.ContentStreamOperation, _ad _cgg.GraphicsState, _edb *_eg.PdfPageResources) error {
		_bd.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0025\u0073", _bec.Operand)
		switch _bec.Operand {
		case "\u0071":
			_afa.Push()
		case "\u0051":
			_afa.Pop()
			_ba = _afa.TextState()
		case "\u0063\u006d":
			if len(_bec.Params) != 6 {
				return _ff
			}
			_fb, _bfe := _ee.GetNumbersAsFloat(_bec.Params)
			if _bfe != nil {
				return _bfe
			}
			_dbd := _af.NewMatrix(_fb[0], _fb[1], _fb[2], _fb[3], _fb[4], _fb[5])
			_bd.Log.Debug("\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074a\u0074\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078\u003a \u0025\u002b\u0076", _dbd)
			_afa.SetMatrix(_afa.Matrix().Mult(_dbd))
		case "\u0077":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_bad, _cdbe := _ee.GetNumbersAsFloat(_bec.Params)
			if _cdbe != nil {
				return _cdbe
			}
			_afa.SetLineWidth(_bad[0])
		case "\u004a":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_cga, _aade := _ee.GetIntVal(_bec.Params[0])
			if !_aade {
				return _ef
			}
			switch _cga {
			case 0:
				_afa.SetLineCap(_ae.LineCapButt)
			case 1:
				_afa.SetLineCap(_ae.LineCapRound)
			case 2:
				_afa.SetLineCap(_ae.LineCapSquare)
			default:
				_bd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u006ee\u0020\u0063\u0061\u0070\u0020\u0073\u0074\u0079\u006c\u0065:\u0020\u0025\u0064", _cga)
				return _ff
			}
		case "\u006a":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_ffe, _dga := _ee.GetIntVal(_bec.Params[0])
			if !_dga {
				return _ef
			}
			switch _ffe {
			case 0:
				_afa.SetLineJoin(_ae.LineJoinBevel)
			case 1:
				_afa.SetLineJoin(_ae.LineJoinRound)
			case 2:
				_afa.SetLineJoin(_ae.LineJoinBevel)
			default:
				_bd.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u006e\u0065\u0020\u006a\u006f\u0069\u006e \u0073\u0074\u0079l\u0065:\u0020\u0025\u0064", _ffe)
				return _ff
			}
		case "\u004d":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_fgd, _cb := _ee.GetNumbersAsFloat(_bec.Params)
			if _cb != nil {
				return _cb
			}
			_ = _fgd
			_bd.Log.Debug("\u004di\u0074\u0065\u0072\u0020l\u0069\u006d\u0069\u0074\u0020n\u006ft\u0020s\u0075\u0070\u0070\u006f\u0072\u0074\u0065d")
		case "\u0064":
			if len(_bec.Params) != 2 {
				return _ff
			}
			_afb, _gfb := _ee.GetArray(_bec.Params[0])
			if !_gfb {
				return _ef
			}
			_de, _gfb := _ee.GetIntVal(_bec.Params[1])
			if !_gfb {
				return _ef
			}
			_cdcgd, _eea := _ee.GetNumbersAsFloat(_afb.Elements())
			if _eea != nil {
				return _eea
			}
			_afa.SetDash(_cdcgd...)
			_ = _de
			_bd.Log.Debug("\u004c\u0069n\u0065\u0020\u0064\u0061\u0073\u0068\u0020\u0070\u0068\u0061\u0073\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006frt\u0065\u0064")
		case "\u0072\u0069":
			_bd.Log.Debug("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020i\u006e\u0074\u0065\u006e\u0074\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0069":
			_bd.Log.Debug("\u0046\u006c\u0061\u0074\u006e\u0065\u0073\u0073\u0020\u0074\u006f\u006c\u0065\u0072\u0061n\u0063e\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		case "\u0067\u0073":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_cbd, _ga := _ee.GetName(_bec.Params[0])
			if !_ga {
				return _ef
			}
			if _cbd == nil {
				return _ff
			}
			_ccb, _ga := _edb.GetExtGState(*_cbd)
			if !_ga {
				_bd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074 \u0066i\u006ed\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u003a\u0020\u0025\u0073", *_cbd)
				return _a.New("\u0072e\u0073o\u0075\u0072\u0063\u0065\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064")
			}
			_df, _ga := _ee.GetDict(_ccb)
			if !_ga {
				_bd.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020c\u006f\u0075\u006c\u0064 ge\u0074 g\u0072\u0061\u0070\u0068\u0069\u0063\u0073 s\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074")
				return _ef
			}
			_bd.Log.Debug("G\u0053\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _df.String())
		case "\u006d":
			if len(_bec.Params) != 2 {
				_bd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _ff)
				return nil
			}
			_fgff, _eeg := _ee.GetNumbersAsFloat(_bec.Params)
			if _eeg != nil {
				return _eeg
			}
			_bd.Log.Debug("M\u006f\u0076\u0065\u0020\u0074\u006f\u003a\u0020\u0025\u0076", _fgff)
			_afa.NewSubPath()
			_afa.MoveTo(_fgff[0], _fgff[1])
		case "\u006c":
			if len(_bec.Params) != 2 {
				_bd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _ff)
				return nil
			}
			_dfe, _fd := _ee.GetNumbersAsFloat(_bec.Params)
			if _fd != nil {
				return _fd
			}
			_afa.LineTo(_dfe[0], _dfe[1])
		case "\u0063":
			if len(_bec.Params) != 6 {
				return _ff
			}
			_geg, _dda := _ee.GetNumbersAsFloat(_bec.Params)
			if _dda != nil {
				return _dda
			}
			_bd.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _geg)
			_afa.CubicTo(_geg[0], _geg[1], _geg[2], _geg[3], _geg[4], _geg[5])
		case "\u0076", "\u0079":
			if len(_bec.Params) != 4 {
				return _ff
			}
			_ddc, _fbe := _ee.GetNumbersAsFloat(_bec.Params)
			if _fbe != nil {
				return _fbe
			}
			_bd.Log.Debug("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076", _ddc)
			_afa.QuadraticTo(_ddc[0], _ddc[1], _ddc[2], _ddc[3])
		case "\u0068":
			_afa.ClosePath()
			_afa.NewSubPath()
		case "\u0072\u0065":
			if len(_bec.Params) != 4 {
				return _ff
			}
			_egc, _cgd := _ee.GetNumbersAsFloat(_bec.Params)
			if _cgd != nil {
				return _cgd
			}
			_afa.DrawRectangle(_egc[0], _egc[1], _egc[2], _egc[3])
			_afa.NewSubPath()
		case "\u0053":
			_fc, _faee := _ad.ColorspaceStroking.ColorToRGB(_ad.ColorStroking)
			if _faee != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _faee)
				return _faee
			}
			_edd, _beg := _fc.(*_eg.PdfColorDeviceRGB)
			if !_beg {
				_bd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _faee
			}
			_afa.SetRGBA(_edd.R(), _edd.G(), _edd.B(), 1)
			_afa.Stroke()
		case "\u0073":
			_ded, _dc := _ad.ColorspaceStroking.ColorToRGB(_ad.ColorStroking)
			if _dc != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dc)
				return _dc
			}
			_ea, _ffd := _ded.(*_eg.PdfColorDeviceRGB)
			if !_ffd {
				_bd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _dc
			}
			_afa.ClosePath()
			_afa.NewSubPath()
			_afa.SetRGBA(_ea.R(), _ea.G(), _ea.B(), 1)
			_afa.Stroke()
		case "\u0066", "\u0046":
			_baa, _bbf := _ad.ColorspaceNonStroking.ColorToRGB(_ad.ColorNonStroking)
			if _bbf != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bbf)
				return _bbf
			}
			_fe, _cad := _baa.(*_eg.PdfColorDeviceRGB)
			if !_cad {
				_bd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _bbf
			}
			_afa.SetRGBA(_fe.R(), _fe.G(), _fe.B(), 1)
			_afa.SetFillRule(_ae.FillRuleWinding)
			_afa.Fill()
		case "\u0066\u002a":
			_afc, _aae := _ad.ColorspaceNonStroking.ColorToRGB(_ad.ColorNonStroking)
			if _aae != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _aae)
				return _aae
			}
			_bgb, _dff := _afc.(*_eg.PdfColorDeviceRGB)
			if !_dff {
				_bd.Log.Debug("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072")
				return _aae
			}
			_afa.SetRGBA(_bgb.R(), _bgb.G(), _bgb.B(), 1)
			_afa.SetFillRule(_ae.FillRuleEvenOdd)
			_afa.Fill()
		case "\u0042":
			_eed, _bgd := _ad.ColorspaceNonStroking.ColorToRGB(_ad.ColorNonStroking)
			if _bgd != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bgd)
				return _bgd
			}
			_gc := _eed.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_gc.R(), _gc.G(), _gc.B(), 1)
			_afa.SetFillRule(_ae.FillRuleWinding)
			_afa.FillPreserve()
			_eed, _bgd = _ad.ColorspaceStroking.ColorToRGB(_ad.ColorStroking)
			if _bgd != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bgd)
				return _bgd
			}
			_gc = _eed.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_gc.R(), _gc.G(), _gc.B(), 1)
			_afa.Stroke()
		case "\u0042\u002a":
			_cba, _aeb := _ad.ColorspaceNonStroking.ColorToRGB(_ad.ColorNonStroking)
			if _aeb != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _aeb)
				return _aeb
			}
			_bbd := _cba.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_bbd.R(), _bbd.G(), _bbd.B(), 1)
			_afa.SetFillRule(_ae.FillRuleEvenOdd)
			_afa.FillPreserve()
			_cba, _aeb = _ad.ColorspaceStroking.ColorToRGB(_ad.ColorStroking)
			if _aeb != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _aeb)
				return _aeb
			}
			_bbd = _cba.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_bbd.R(), _bbd.G(), _bbd.B(), 1)
			_afa.Stroke()
		case "\u0062":
			_fef, _ddf := _ad.ColorspaceNonStroking.ColorToRGB(_ad.ColorNonStroking)
			if _ddf != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ddf)
				return _ddf
			}
			_gfg := _fef.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_gfg.R(), _gfg.G(), _gfg.B(), 1)
			_afa.ClosePath()
			_afa.NewSubPath()
			_afa.SetFillRule(_ae.FillRuleWinding)
			_afa.FillPreserve()
			_fef, _ddf = _ad.ColorspaceStroking.ColorToRGB(_ad.ColorStroking)
			if _ddf != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ddf)
				return _ddf
			}
			_gfg = _fef.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_gfg.R(), _gfg.G(), _gfg.B(), 1)
			_afa.Stroke()
		case "\u0062\u002a":
			_afa.ClosePath()
			_gg, _eae := _ad.ColorspaceNonStroking.ColorToRGB(_ad.ColorNonStroking)
			if _eae != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _eae)
				return _eae
			}
			_gcb := _gg.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_gcb.R(), _gcb.G(), _gcb.B(), 1)
			_afa.NewSubPath()
			_afa.SetFillRule(_ae.FillRuleEvenOdd)
			_afa.FillPreserve()
			_gg, _eae = _ad.ColorspaceStroking.ColorToRGB(_ad.ColorStroking)
			if _eae != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _eae)
				return _eae
			}
			_gcb = _gg.(*_eg.PdfColorDeviceRGB)
			_afa.SetRGBA(_gcb.R(), _gcb.G(), _gcb.B(), 1)
			_afa.Stroke()
		case "\u006e":
			_afa.ClearPath()
		case "\u0057":
			_afa.SetFillRule(_ae.FillRuleWinding)
			_afa.ClipPreserve()
		case "\u0057\u002a":
			_afa.SetFillRule(_ae.FillRuleEvenOdd)
			_afa.ClipPreserve()
		case "\u0072\u0067":
			_fde, _fdb := _ad.ColorNonStroking.(*_eg.PdfColorDeviceRGB)
			if !_fdb {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorNonStroking)
				return nil
			}
			_afa.SetFillRGBA(_fde.R(), _fde.G(), _fde.B(), 1)
		case "\u0052\u0047":
			_gfd, _aaa := _ad.ColorStroking.(*_eg.PdfColorDeviceRGB)
			if !_aaa {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorStroking)
				return nil
			}
			_afa.SetStrokeRGBA(_gfd.R(), _gfd.G(), _gfd.B(), 1)
		case "\u006b":
			_aeg, _dea := _ad.ColorNonStroking.(*_eg.PdfColorDeviceCMYK)
			if !_dea {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorNonStroking)
				return nil
			}
			_dde, _bed := _ad.ColorspaceNonStroking.ColorToRGB(_aeg)
			if _bed != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorNonStroking)
				return nil
			}
			_bbff, _dea := _dde.(*_eg.PdfColorDeviceRGB)
			if !_dea {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _dde)
				return nil
			}
			_afa.SetFillRGBA(_bbff.R(), _bbff.G(), _bbff.B(), 1)
		case "\u004b":
			_ebd, _dbf := _ad.ColorStroking.(*_eg.PdfColorDeviceCMYK)
			if !_dbf {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorStroking)
				return nil
			}
			_gda, _ege := _ad.ColorspaceStroking.ColorToRGB(_ebd)
			if _ege != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorStroking)
				return nil
			}
			_cf, _dbf := _gda.(*_eg.PdfColorDeviceRGB)
			if !_dbf {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _gda)
				return nil
			}
			_afa.SetStrokeRGBA(_cf.R(), _cf.G(), _cf.B(), 1)
		case "\u0067":
			_agf, _aef := _ad.ColorNonStroking.(*_eg.PdfColorDeviceGray)
			if !_aef {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorNonStroking)
				return nil
			}
			_afg, _ebda := _ad.ColorspaceNonStroking.ColorToRGB(_agf)
			if _ebda != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorNonStroking)
				return nil
			}
			_eda, _aef := _afg.(*_eg.PdfColorDeviceRGB)
			if !_aef {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _afg)
				return nil
			}
			_afa.SetFillRGBA(_eda.R(), _eda.G(), _eda.B(), 1)
		case "\u0047":
			_dfb, _aaeg := _ad.ColorStroking.(*_eg.PdfColorDeviceGray)
			if !_aaeg {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorStroking)
				return nil
			}
			_bede, _gdf := _ad.ColorspaceStroking.ColorToRGB(_dfb)
			if _gdf != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorStroking)
				return nil
			}
			_aebg, _aaeg := _bede.(*_eg.PdfColorDeviceRGB)
			if !_aaeg {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _bede)
				return nil
			}
			_afa.SetStrokeRGBA(_aebg.R(), _aebg.G(), _aebg.B(), 1)
		case "\u0063\u0073", "\u0073\u0063", "\u0073\u0063\u006e":
			_agg, _aaf := _ad.ColorspaceNonStroking.ColorToRGB(_ad.ColorNonStroking)
			if _aaf != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorNonStroking)
				return nil
			}
			_faf, _ddea := _agg.(*_eg.PdfColorDeviceRGB)
			if !_ddea {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _agg)
				return nil
			}
			_afa.SetFillRGBA(_faf.R(), _faf.G(), _faf.B(), 1)
		case "\u0043\u0053", "\u0053\u0043", "\u0053\u0043\u004e":
			_ddeg, _adf := _ad.ColorspaceStroking.ColorToRGB(_ad.ColorStroking)
			if _adf != nil {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ad.ColorStroking)
				return nil
			}
			_ddb, _cef := _ddeg.(*_eg.PdfColorDeviceRGB)
			if !_cef {
				_bd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076", _ddeg)
				return nil
			}
			_afa.SetStrokeRGBA(_ddb.R(), _ddb.G(), _ddb.B(), 1)
		case "\u0044\u006f":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_cfg, _edbf := _ee.GetName(_bec.Params[0])
			if !_edbf {
				return _ef
			}
			_, _aff := _edb.GetXObjectByName(*_cfg)
			switch _aff {
			case _eg.XObjectTypeImage:
				_bd.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067e\u003a\u0020\u0025\u0073", _cfg.String())
				_cgag, _dbac := _edb.GetXObjectImageByName(*_cfg)
				if _dbac != nil {
					return _dbac
				}
				_cfgg, _dbac := _cgag.ToImage()
				if _dbac != nil {
					return _dbac
				}
				if _gca := _cgag.ColorSpace; _gca != nil {
					var _fcb bool
					switch _gca.(type) {
					case *_eg.PdfColorspaceSpecialIndexed:
						_fcb = true
					}
					if _fcb {
						if _gfac, _ggb := _gca.ImageToRGB(*_cfgg); _ggb != nil {
							_bd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006fnv\u0065r\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0052G\u0042\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020i\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
						} else {
							_cfgg = &_gfac
						}
					}
				}
				_ccba := _afa.FillPattern().ColorAt(0, 0)
				var _dca _gb.Image
				if _cgag.Mask != nil {
					if _dca, _dbac = _dfd(_cgag.Mask, _ccba); _dbac != nil {
						_bd.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				} else if _cgag.SMask != nil {
					if _dca, _dbac = _efd(_cgag.SMask, _ccba); _dbac != nil {
						_bd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
					}
				}
				var _cec _gb.Image
				if _bda, _ := _ee.GetBoolVal(_cgag.ImageMask); _bda {
					_cec = _ecc(_cfgg, _ccba)
				} else {
					_cec, _dbac = _cfgg.ToGoImage()
					if _dbac != nil {
						return _dbac
					}
				}
				if _dca != nil {
					_cec = _gab(_cec, _dca)
				}
				_agfa := _cec.Bounds()
				_afa.Push()
				_afa.Scale(1.0/float64(_agfa.Dx()), -1.0/float64(_agfa.Dy()))
				_afa.DrawImageAnchored(_cec, 0, 0, 0, 1)
				_afa.Pop()
			case _eg.XObjectTypeForm:
				_bd.Log.Debug("\u0058\u004fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u0073", _cfg.String())
				_aec, _eeae := _edb.GetXObjectFormByName(*_cfg)
				if _eeae != nil {
					return _eeae
				}
				_eaf, _eeae := _aec.GetContentStream()
				if _eeae != nil {
					return _eeae
				}
				_gaf := _aec.Resources
				if _gaf == nil {
					_gaf = _edb
				}
				_afa.Push()
				if _aec.Matrix != nil {
					_gee, _abc := _ee.GetArray(_aec.Matrix)
					if !_abc {
						return _ef
					}
					_fgcf, _cgaa := _ee.GetNumbersAsFloat(_gee.Elements())
					if _cgaa != nil {
						return _cgaa
					}
					if len(_fgcf) != 6 {
						return _ff
					}
					_aac := _af.NewMatrix(_fgcf[0], _fgcf[1], _fgcf[2], _fgcf[3], _fgcf[4], _fgcf[5])
					_afa.SetMatrix(_afa.Matrix().Mult(_aac))
				}
				if _aec.BBox != nil {
					_bfed, _gga := _ee.GetArray(_aec.BBox)
					if !_gga {
						return _ef
					}
					_fgb, _fgdg := _ee.GetNumbersAsFloat(_bfed.Elements())
					if _fgdg != nil {
						return _fgdg
					}
					if len(_fgb) != 4 {
						_bd.Log.Debug("\u004c\u0065\u006e\u0020\u003d\u0020\u0025\u0064", len(_fgb))
						return _ff
					}
					_afa.DrawRectangle(_fgb[0], _fgb[1], _fgb[2]-_fgb[0], _fgb[3]-_fgb[1])
					_afa.SetRGBA(1, 0, 0, 1)
					_afa.Clip()
				} else {
					_bd.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0052\u0065q\u0075\u0069\u0072e\u0064\u0020\u0042\u0042\u006f\u0078\u0020\u006d\u0069ss\u0069\u006e\u0067 \u006f\u006e \u0058\u004f\u0062\u006a\u0065\u0063t\u0020\u0046o\u0072\u006d")
				}
				_eeae = _bdg.renderContentStream(_afa, string(_eaf), _gaf)
				if _eeae != nil {
					return _eeae
				}
				_afa.Pop()
			}
		case "\u0042\u0049":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_bbg, _bfaf := _bec.Params[0].(*_cgg.ContentStreamInlineImage)
			if !_bfaf {
				return nil
			}
			_fec, _abaf := _bbg.ToImage(_edb)
			if _abaf != nil {
				return _abaf
			}
			_ddaa, _abaf := _fec.ToGoImage()
			if _abaf != nil {
				return _abaf
			}
			_aeca := _ddaa.Bounds()
			_afa.Push()
			_afa.Scale(1.0/float64(_aeca.Dx()), -1.0/float64(_aeca.Dy()))
			_afa.DrawImageAnchored(_ddaa, 0, 0, 0, 1)
			_afa.Pop()
		case "\u0042\u0054":
			_ba.Reset()
		case "\u0045\u0054":
			_ba.Reset()
		case "\u0054\u0072":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_ggd, _ebe := _ee.GetNumberAsFloat(_bec.Params[0])
			if _ebe != nil {
				return _ebe
			}
			_ba.Tr = _ae.TextRenderingMode(_ggd)
		case "\u0054\u004c":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_cgc, _bcc := _ee.GetNumberAsFloat(_bec.Params[0])
			if _bcc != nil {
				return _bcc
			}
			_ba.Tl = _cgc
		case "\u0054\u0063":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_fgdb, _aadb := _ee.GetNumberAsFloat(_bec.Params[0])
			if _aadb != nil {
				return _aadb
			}
			_bd.Log.Debug("\u0054\u0063\u003a\u0020\u0025\u0076", _fgdb)
			_ba.Tc = _fgdb
		case "\u0054\u0077":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_ggc, _ccdc := _ee.GetNumberAsFloat(_bec.Params[0])
			if _ccdc != nil {
				return _ccdc
			}
			_bd.Log.Debug("\u0054\u0077\u003a\u0020\u0025\u0076", _ggc)
			_ba.Tw = _ggc
		case "\u0054\u007a":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_dbb, _gfc := _ee.GetNumberAsFloat(_bec.Params[0])
			if _gfc != nil {
				return _gfc
			}
			_ba.Th = _dbb
		case "\u0054\u0073":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_ddba, _dbbc := _ee.GetNumberAsFloat(_bec.Params[0])
			if _dbbc != nil {
				return _dbbc
			}
			_ba.Ts = _ddba
		case "\u0054\u0064":
			if len(_bec.Params) != 2 {
				return _ff
			}
			_bfac, _bdfb := _ee.GetNumbersAsFloat(_bec.Params)
			if _bdfb != nil {
				return _bdfb
			}
			_bd.Log.Debug("\u0054\u0064\u003a\u0020\u0025\u0076", _bfac)
			_ba.ProcTd(_bfac[0], _bfac[1])
		case "\u0054\u0044":
			if len(_bec.Params) != 2 {
				return _ff
			}
			_cab, _ecd := _ee.GetNumbersAsFloat(_bec.Params)
			if _ecd != nil {
				return _ecd
			}
			_bd.Log.Debug("\u0054\u0044\u003a\u0020\u0025\u0076", _cab)
			_ba.ProcTD(_cab[0], _cab[1])
		case "\u0054\u002a":
			_ba.ProcTStar()
		case "\u0054\u006d":
			if len(_bec.Params) != 6 {
				return _ff
			}
			_da, _abf := _ee.GetNumbersAsFloat(_bec.Params)
			if _abf != nil {
				return _abf
			}
			_bd.Log.Debug("\u0054\u0065x\u0074\u0020\u006da\u0074\u0072\u0069\u0078\u003a\u0020\u0025\u002b\u0076", _da)
			_ba.ProcTm(_da[0], _da[1], _da[2], _da[3], _da[4], _da[5])
		case "\u0027":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_ebaa, _eaec := _ee.GetStringBytes(_bec.Params[0])
			if !_eaec {
				return _ef
			}
			_bd.Log.Debug("\u0027\u0020\u0073t\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_ebaa))
			_ba.ProcQ(_ebaa, _afa)
		case "\u0022":
			if len(_bec.Params) != 3 {
				return _ff
			}
			_cfe, _cda := _ee.GetNumberAsFloat(_bec.Params[0])
			if _cda != nil {
				return _cda
			}
			_fbea, _cda := _ee.GetNumberAsFloat(_bec.Params[1])
			if _cda != nil {
				return _cda
			}
			_eef, _gef := _ee.GetStringBytes(_bec.Params[2])
			if !_gef {
				return _ef
			}
			_ba.ProcDQ(_eef, _cfe, _fbea, _afa)
		case "\u0054\u006a":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_bfaff, _aeff := _ee.GetStringBytes(_bec.Params[0])
			if !_aeff {
				return _ef
			}
			_bd.Log.Debug("\u0054j\u0020s\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0060\u0025\u0073\u0060", string(_bfaff))
			_ba.ProcTj(_bfaff, _afa)
		case "\u0054\u004a":
			if len(_bec.Params) != 1 {
				return _ff
			}
			_aab, _fcg := _ee.GetArray(_bec.Params[0])
			if !_fcg {
				_bd.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _aab)
				return _ef
			}
			_bd.Log.Debug("\u0054\u004a\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u002b\u0076", _aab)
			for _, _affg := range _aab.Elements() {
				switch _dgg := _affg.(type) {
				case *_ee.PdfObjectString:
					if _dgg != nil {
						_ba.ProcTj(_dgg.Bytes(), _afa)
					}
				case *_ee.PdfObjectFloat, *_ee.PdfObjectInteger:
					_bde, _bff := _ee.GetNumberAsFloat(_dgg)
					if _bff == nil {
						_ba.Translate(-_bde*0.001*_ba.Tf.Size*_ba.Th/100.0, 0)
					}
				}
			}
		case "\u0054\u0066":
			if len(_bec.Params) != 2 {
				return _ff
			}
			_bd.Log.Debug("\u0025\u0023\u0076", _bec.Params)
			_fca, _ffg := _ee.GetName(_bec.Params[0])
			if !_ffg || _fca == nil {
				_bd.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006e\u0061m\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _bec.Params[0])
				return _ef
			}
			_bd.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073", _fca.String())
			_egg, _bfc := _ee.GetNumberAsFloat(_bec.Params[1])
			if _bfc != nil {
				_bd.Log.Debug("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0069z\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _bec.Params[1])
				return _ef
			}
			_bd.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0076", _egg)
			_gfbf, _aee := _edb.GetFontByName(*_fca)
			if !_aee {
				_bd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0025s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _fca.String())
				return _a.New("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
			}
			_bd.Log.Debug("\u0046\u006f\u006e\u0074\u003a\u0020\u0025\u0054", _gfbf)
			_cbaf, _ffg := _ee.GetDict(_gfbf)
			if !_ffg {
				_bd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0067e\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074")
				return _ef
			}
			_ebb, _bfc := _eg.NewPdfFontFromPdfObject(_cbaf)
			if _bfc != nil {
				_bd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
				return _bfc
			}
			_cca := _ebb.BaseFont()
			if _cca == "" {
				_cca = _fca.String()
			}
			_cdbc, _ffg := _aba[_cca]
			if !_ffg {
				_cdbc, _bfc = _ae.NewTextFont(_ebb, _egg)
				if _bfc != nil {
					_bd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bfc)
				}
			}
			if _cdbc == nil {
				if len(_cca) > 7 && _cca[6] == '+' {
					_cca = _cca[7:]
				}
				_ebeb := []string{_cca, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
				for _, _aabc := range _ebeb {
					_bd.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _aabc)
					if _cdbc, _ffg = _aba[_aabc]; _ffg {
						break
					}
					_afcf := _cdcg.Match(_aabc)
					if _afcf == nil {
						_bd.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _aabc)
						continue
					}
					_cdbc, _bfc = _ae.NewTextFontFromPath(_afcf.Filename, _egg)
					if _bfc != nil {
						_bd.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _afcf.Filename)
						continue
					}
					_bd.Log.Debug("\u0053\u0075\u0062\u0073\u0074\u0069t\u0075\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073 \u0077\u0069\u0074\u0068\u0020\u0025\u0073 \u0028\u0025\u0073\u0029", _cca, _afcf.Name, _afcf.Filename)
					_aba[_aabc] = _cdbc
					break
				}
			}
			if _cdbc == nil {
				_bd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020n\u006f\u0074\u0020\u0066\u0069\u006ed\u0020\u0061\u006e\u0079\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065 \u0066\u006f\u006e\u0074")
				return _a.New("\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0066\u0069\u006e\u0064\u0020a\u006ey\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u006e\u0074")
			}
			_ba.ProcTf(_cdbc.WithSize(_egg, _ebb))
		case "\u0042\u004d\u0043", "\u0042\u0044\u0043":
		case "\u0045\u004d\u0043":
		default:
			_bd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u006f\u0070\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u0073", _bec.Operand)
		}
		return nil
	})
	_bfa = _edg.Process(_bbb)
	if _bfa != nil {
		return _bfa
	}
	return nil
}
func _ffdg(_ac string, _bbbc _gb.Image, _bdfe int) error {
	_cbae, _dbfa := _c.Create(_ac)
	if _dbfa != nil {
		return _dbfa
	}
	defer _cbae.Close()
	return _eb.Encode(_cbae, _bbbc, &_eb.Options{Quality: _bdfe})
}

// Render converts the specified PDF page into an image and returns the result.
func (_ge *ImageDevice) Render(page *_eg.PdfPage) (_gb.Image, error) {
	_ebf, _caa := page.GetMediaBox()
	if _caa != nil {
		return nil, _caa
	}
	_ebf.Normalize()
	_dg := page.CropBox
	var _db, _f float64
	if _dg != nil {
		_dg.Normalize()
		_db, _f = _dg.Width(), _dg.Height()
	}
	_fg := page.Rotate
	_cd, _aa, _gf, _dbg := _ebf.Llx, _ebf.Lly, _ebf.Width(), _ebf.Height()
	_dba := _af.IdentityMatrix()
	if _fg != nil && *_fg%360 != 0 && *_fg%90 == 0 {
		_fa := -float64(*_fg)
		_eee := _bcgf(_gf, _dbg, _fa)
		_dba = _dba.Translate((_eee.Width-_gf)/2+_gf/2, (_eee.Height-_dbg)/2+_dbg/2).Rotate(_fa*_g.Pi/180).Translate(-_gf/2, -_dbg/2)
		_gf, _dbg = _eee.Width, _eee.Height
		if _dg != nil {
			_aea := _bcgf(_db, _f, _fa)
			_db, _f = _aea.Width, _aea.Height
		}
	}
	if _cd != 0 || _aa != 0 {
		_dba = _dba.Translate(-_cd, -_aa)
	}
	_ge._gba = 1.0
	if _ge.OutputWidth != 0 {
		_cdb := _gf
		if _dg != nil {
			_cdb = _db
		}
		_ge._gba = float64(_ge.OutputWidth) / _cdb
		_gf, _dbg, _db, _f = _gf*_ge._gba, _dbg*_ge._gba, _db*_ge._gba, _f*_ge._gba
		_dba = _af.ScaleMatrix(_ge._gba, _ge._gba).Mult(_dba)
	}
	_gfa := _bc.NewContext(int(_gf), int(_dbg))
	if _cdc := _ge.renderPage(_gfa, page, _dba); _cdc != nil {
		return nil, _cdc
	}
	_cea := _gfa.Image()
	if _dg != nil {
		_gd, _fgc := (_dg.Llx-_cd)*_ge._gba, (_dg.Lly-_aa)*_ge._gba
		_ebfc := _gb.Rect(0, 0, int(_db), int(_f))
		_bdf := _gb.Pt(int(_gd), int(_dbg-_fgc-_f))
		_eba := _gb.NewRGBA(_ebfc)
		_cg.Draw(_eba, _ebfc, _cea, _bdf, _cg.Src)
		_cea = _eba
	}
	return _cea, nil
}

// NewImageDevice returns a new image device.
func NewImageDevice() *ImageDevice {
	return &ImageDevice{}
}
func (_dbge renderer) renderPage(_aeae _ae.Context, _ab *_eg.PdfPage, _aad _af.Matrix) error {
	_efe, _gea := _ab.GetAllContentStreams()
	if _gea != nil {
		return _gea
	}
	if _ece := _aad; !_ece.Identity() {
		_efe = _d.Sprintf("%\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066 \u0025\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066\u0020\u0063\u006d\u0020%\u0073", _ece[0], _ece[1], _ece[3], _ece[4], _ece[6], _ece[7], _efe)
	}
	_aeae.Translate(0, float64(_aeae.Height()))
	_aeae.Scale(1, -1)
	_aeae.Push()
	_aeae.SetRGBA(1, 1, 1, 1)
	_aeae.DrawRectangle(0, 0, float64(_aeae.Width()), float64(_aeae.Height()))
	_aeae.Fill()
	_aeae.Pop()
	_aeae.SetLineWidth(1.0)
	_aeae.SetRGBA(0, 0, 0, 1)
	return _dbge.renderContentStream(_aeae, _efe, _ab.Resources)
}
func _dfd(_add _ee.PdfObject, _dgb _b.Color) (_gb.Image, error) {
	_dbe, _cdd := _ee.GetStream(_add)
	if !_cdd {
		return nil, nil
	}
	_bbbb, _fed := _eg.NewXObjectImageFromStream(_dbe)
	if _fed != nil {
		return nil, _fed
	}
	_dfbg, _fed := _bbbb.ToImage()
	if _fed != nil {
		return nil, _fed
	}
	return _ecc(_dfbg, _dgb), nil
}
func _gab(_fcge, _dcg _gb.Image) _gb.Image {
	_ced, _dfa := _dcg.Bounds().Size(), _fcge.Bounds().Size()
	_dgad, _fgffc := _ced.X, _ced.Y
	if _dfa.X > _dgad {
		_dgad = _dfa.X
	}
	if _dfa.Y > _fgffc {
		_fgffc = _dfa.Y
	}
	_cefb := _gb.Rect(0, 0, _dgad, _fgffc)
	if _ced.X != _dgad || _ced.Y != _fgffc {
		_ace := _gb.NewRGBA(_cefb)
		_bg.BiLinear.Scale(_ace, _cefb, _fcge, _dcg.Bounds(), _bg.Over, nil)
		_dcg = _ace
	}
	if _dfa.X != _dgad || _dfa.Y != _fgffc {
		_edgg := _gb.NewRGBA(_cefb)
		_bg.BiLinear.Scale(_edgg, _cefb, _fcge, _fcge.Bounds(), _bg.Over, nil)
		_fcge = _edgg
	}
	_eac := _gb.NewRGBA(_cefb)
	_bg.DrawMask(_eac, _cefb, _fcge, _gb.Point{}, _dcg, _gb.Point{}, _bg.Over)
	return _eac
}

// ImageDevice is used to render PDF pages to image targets.
type ImageDevice struct {
	renderer

	// OutputWidth represents the width of the rendered images in pixels.
	// The heights of the output images are calculated based on the selected
	// width and the original height of each rendered page.
	OutputWidth int
}

func _ecc(_begc *_eg.Image, _fgg _b.Color) _gb.Image {
	_bag, _ebaad := int(_begc.Width), int(_begc.Height)
	_ade := _gb.NewRGBA(_gb.Rect(0, 0, _bag, _ebaad))
	for _egb := 0; _egb < _ebaad; _egb++ {
		for _cabb := 0; _cabb < _bag; _cabb++ {
			_caba, _ggf := _begc.ColorAt(_cabb, _egb)
			if _ggf != nil {
				_bd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _cabb, _egb)
				continue
			}
			_geef, _ede, _fee, _ := _caba.RGBA()
			var _gbd _b.Color
			if _geef+_ede+_fee == 0 {
				_gbd = _fgg
			} else {
				_gbd = _b.Transparent
			}
			_ade.Set(_cabb, _egb, _gbd)
		}
	}
	return _ade
}
func _eeec(_feg *_eg.Image, _cfd _b.Color) _gb.Image {
	_cce, _cbdg := int(_feg.Width), int(_feg.Height)
	_gdac := _gb.NewRGBA(_gb.Rect(0, 0, _cce, _cbdg))
	for _bbbg := 0; _bbbg < _cbdg; _bbbg++ {
		for _abb := 0; _abb < _cce; _abb++ {
			_adac, _gag := _feg.ColorAt(_abb, _bbbg)
			if _gag != nil {
				_bd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _abb, _bbbg)
				continue
			}
			_bgc, _egbc, _bgg, _ := _adac.RGBA()
			var _bbda _b.Color
			if _bgc+_egbc+_bgg == 0 {
				_bbda = _b.Transparent
			} else {
				_bbda = _cfd
			}
			_gdac.Set(_abb, _bbbg, _bbda)
		}
	}
	return _gdac
}
func _bcgf(_aeab, _cfb, _bce float64) _ce.BoundingBox {
	return _ce.Path{Points: []_ce.Point{_ce.NewPoint(0, 0).Rotate(_bce), _ce.NewPoint(_aeab, 0).Rotate(_bce), _ce.NewPoint(0, _cfb).Rotate(_bce), _ce.NewPoint(_aeab, _cfb).Rotate(_bce)}}.GetBoundingBox()
}
func _bcd(_deb string, _gdg _gb.Image) error {
	_fda, _ada := _c.Create(_deb)
	if _ada != nil {
		return _ada
	}
	defer _fda.Close()
	return _ec.Encode(_fda, _gdg)
}
