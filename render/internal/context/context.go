package context

import (
	_d "errors"
	_af "image"
	_e "image/color"

	_afg "bitbucket.org/shenghui0779/gopdf/core"
	_ef "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_da "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_dg "bitbucket.org/shenghui0779/gopdf/model"
	_dgb "github.com/unidoc/freetype/truetype"
	_f "golang.org/x/image/font"
)

type LineJoin int

func (_df *TextFont) NewFace(size float64) _f.Face {
	return _dgb.NewFace(_df._fe, &_dgb.Options{Size: size})
}

type TextRenderingMode int

const (
	TextRenderingModeFill TextRenderingMode = iota
	TextRenderingModeStroke
	TextRenderingModeFillStroke
	TextRenderingModeInvisible
	TextRenderingModeFillClip
	TextRenderingModeStrokeClip
	TextRenderingModeFillStrokeClip
	TextRenderingModeClip
)

func (_bfc *TextState) ProcTj(data []byte, ctx Context) {
	_gfa := _bfc.Tf.Size
	_fa := _bfc.Th / 100.0
	_ffd := _bfc.GlobalScale
	_ega := _da.NewMatrix(_gfa*_fa, 0, 0, _gfa, 0, _bfc.Ts)
	_ddc := ctx.Matrix()
	_gca := _ddc.Clone().Mult(_bfc.Tm.Clone().Mult(_ega)).ScalingFactorY()
	_aag := _bfc.Tf.NewFace(_gca)
	_bag := _bfc.Tf.BytesToCharcodes(data)
	for _, _aga := range _bag {
		_ffdc, _cad := _bfc.Tf.CharcodeToRunes(_aga)
		_cef := string(_cad)
		if _cef == "\u0000" {
			continue
		}
		_aba := _ddc.Clone().Mult(_bfc.Tm.Clone().Mult(_ega))
		_fg := _aba.ScalingFactorY()
		_aba = _aba.Scale(1/_fg, -1/_fg)
		if _bfc.Tr != TextRenderingModeInvisible {
			ctx.SetMatrix(_aba)
			ctx.DrawString(_cef, _aag, 0, 0)
			ctx.SetMatrix(_ddc)
		}
		_ggg := 0.0
		if _cef == "\u0020" {
			_ggg = _bfc.Tw
		}
		_fef, _, _bcd := _bfc.Tf.GetCharMetrics(_ffdc)
		if _bcd {
			_fef = _fef * 0.001 * _gfa
		} else {
			_fef, _ = ctx.MeasureString(_cef, _aag)
			_fef = _fef / _ffd
		}
		_beg := (_fef + _bfc.Tc + _ggg) * _fa
		_bfc.Tm = _bfc.Tm.Mult(_da.TranslationMatrix(_beg, 0))
	}
}
func (_ff *TextState) ProcTm(a, b, c, d, e, f float64) {
	_ff.Tm = _da.NewMatrix(a, b, c, d, e, f)
	_ff.Tlm = _ff.Tm.Clone()
}

type Context interface {
	Push()
	Pop()
	Matrix() _da.Matrix
	SetMatrix(_cb _da.Matrix)
	Translate(_fc, _eg float64)
	Scale(_cd, _fd float64)
	Rotate(_ac float64)
	MoveTo(_afgg, _fdd float64)
	LineTo(_b, _aca float64)
	CubicTo(_afd, _cfc, _fdg, _ad, _fcf, _dd float64)
	QuadraticTo(_ee, _efg, _cfe, _ddd float64)
	NewSubPath()
	ClosePath()
	ClearPath()
	Clip()
	ClipPreserve()
	ResetClip()
	LineWidth() float64
	SetLineWidth(_bf float64)
	SetLineCap(_cg LineCap)
	SetLineJoin(_g LineJoin)
	SetDash(_fcd ...float64)
	SetDashOffset(_ga float64)
	Fill()
	FillPreserve()
	Stroke()
	StrokePreserve()
	SetRGBA(_fdgf, _ba, _ec, _be float64)
	SetFillRGBA(_gf, _ca, _ea, _caa float64)
	SetFillStyle(_aa Pattern)
	SetFillRule(_ag FillRule)
	SetStrokeRGBA(_bfb, _efc, _gc, _gg float64)
	SetStrokeStyle(_dgbe Pattern)
	FillPattern() Pattern
	StrokePattern() Pattern
	TextState() *TextState
	DrawString(_bc string, _eb _f.Face, _gfd, _cce float64)
	MeasureString(_gd string, _ge _f.Face) (_ab, _dc float64)
	DrawRectangle(_ebc, _dcd, _agg, _ged float64)
	DrawImage(_bg _af.Image, _de, _ade int)
	DrawImageAnchored(_cgd _af.Image, _cbf, _dac int, _egg, _acg float64)
	Height() int
	Width() int
}
type LineCap int

func (_cbc *TextState) Translate(tx, ty float64) {
	_cbc.Tm = _cbc.Tm.Mult(_da.TranslationMatrix(tx, ty))
}

type TextState struct {
	Tc          float64
	Tw          float64
	Th          float64
	Tl          float64
	Tf          *TextFont
	Ts          float64
	Tm          _da.Matrix
	Tlm         _da.Matrix
	Tr          TextRenderingMode
	GlobalScale float64
}

func (_fcc *TextFont) WithSize(size float64, originalFont *_dg.PdfFont) *TextFont {
	return &TextFont{Font: _fcc.Font, Size: size, _fe: _fcc._fe, _dbc: originalFont}
}
func NewTextState() TextState {
	return TextState{Th: 100, Tm: _da.IdentityMatrix(), Tlm: _da.IdentityMatrix()}
}
func (_bab *TextFont) CharcodeToRunes(charcode _ef.CharCode) (_ef.CharCode, []rune) {
	_ed := []_ef.CharCode{charcode}
	if _bab._dbc == nil || _bab._dbc == _bab.Font {
		if _bab.Font.IsSimple() && _bab._fe != nil {
			if _ce := _bab._fe.Index(rune(charcode)); _ce > 0 {
				return charcode, []rune{rune(charcode)}
			}
		}
		return charcode, _bab.Font.CharcodesToUnicode(_ed)
	}
	_eae := _bab._dbc.CharcodesToUnicode(_ed)
	_ced, _ := _bab.Font.RunesToCharcodeBytes(_eae)
	_eec := _bab.Font.BytesToCharcodes(_ced)
	_gdd := charcode
	if len(_eec) > 0 && _eec[0] != 0 {
		_gdd = _eec[0]
	}
	return _gdd, _eae
}

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
)

func (_abg *TextState) ProcDQ(data []byte, aw, ac float64, ctx Context) {
	_abg.Tw = aw
	_abg.Tc = ac
	_abg.ProcQ(data, ctx)
}
func (_ddb *TextState) ProcTStar()                     { _ddb.ProcTd(0, -_ddb.Tl) }
func (_aad *TextState) ProcQ(data []byte, ctx Context) { _aad.ProcTStar(); _aad.ProcTj(data, ctx) }

const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

func (_dfg *TextFont) GetCharMetrics(code _ef.CharCode) (float64, float64, bool) {
	if _ada, _edg := _dfg.Font.GetCharMetrics(code); _edg && _ada.Wx != 0 {
		return _ada.Wx, _ada.Wy, _edg
	}
	if _dfg._dbc == nil {
		return 0, 0, false
	}
	_fdc, _bcf := _dfg._dbc.GetCharMetrics(code)
	return _fdc.Wx, _fdc.Wy, _bcf && _fdc.Wx != 0
}
func (_fge *TextState) Reset() {
	_fge.Tm = _da.IdentityMatrix()
	_fge.Tlm = _da.IdentityMatrix()
}
func (_dbb *TextState) ProcTD(tx, ty float64) { _dbb.Tl = -ty; _dbb.ProcTd(tx, ty) }

const (
	LineCapRound LineCap = iota
	LineCapButt
	LineCapSquare
)

func NewTextFont(font *_dg.PdfFont, size float64) (*TextFont, error) {
	_gag := font.FontDescriptor()
	if _gag == nil {
		return nil, _d.New("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072")
	}
	_ead, _ggc := _afg.GetStream(_gag.FontFile2)
	if !_ggc {
		return nil, _d.New("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_bee, _ddg := _afg.DecodeStream(_ead)
	if _ddg != nil {
		return nil, _ddg
	}
	_bea, _ddg := _dgb.Parse(_bee)
	if _ddg != nil {
		return nil, _ddg
	}
	return &TextFont{Font: font, Size: size, _fe: _bea}, nil
}

type Pattern interface {
	ColorAt(_c, _db int) _e.Color
}

func (_eab *TextFont) BytesToCharcodes(data []byte) []_ef.CharCode {
	if _eab._dbc != nil {
		return _eab._dbc.BytesToCharcodes(data)
	}
	return _eab.Font.BytesToCharcodes(data)
}

type TextFont struct {
	Font *_dg.PdfFont
	Size float64
	_fe  *_dgb.Font
	_dbc *_dg.PdfFont
}
type Gradient interface {
	Pattern
	AddColorStop(_cf float64, _cc _e.Color)
}
type FillRule int

func (_ae *TextState) ProcTf(font *TextFont) { _ae.Tf = font }
func (_eaec *TextState) ProcTd(tx, ty float64) {
	_eaec.Tlm.Concat(_da.TranslationMatrix(tx, ty))
	_eaec.Tm = _eaec.Tlm.Clone()
}
func NewTextFontFromPath(filePath string, size float64) (*TextFont, error) {
	_gef, _adb := _dg.NewPdfFontFromTTFFile(filePath)
	if _adb != nil {
		return nil, _adb
	}
	return NewTextFont(_gef, size)
}
