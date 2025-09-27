package context

import (
	_b "errors"
	_cd "image"
	_a "image/color"
	_f "strings"

	_d "unitechio/gopdf/gopdf/core"
	_af "unitechio/gopdf/gopdf/internal/cmap"
	_ef "unitechio/gopdf/gopdf/internal/textencoding"
	_e "unitechio/gopdf/gopdf/internal/transform"
	_cb "unitechio/gopdf/gopdf/model"
	_g "github.com/unidoc/freetype/truetype"
	_fg "golang.org/x/image/font"
)

func (_cgf *TextFont) CharcodeToRunes(charcode _ef.CharCode) (_ef.CharCode, []rune) {
	_ee := []_ef.CharCode{charcode}
	if _cgf._edf == nil || _cgf._edf == _cgf.Font {
		return _cgf.charcodeToRunesSimple(charcode)
	}
	_ffe := _cgf._edf.CharcodesToUnicode(_ee)
	_ac, _ := _cgf.Font.RunesToCharcodeBytes(_ffe)
	_ffd := _cgf.Font.BytesToCharcodes(_ac)
	_gab := charcode
	if len(_ffd) > 0 && _ffd[0] != 0 {
		_gab = _ffd[0]
	}
	if string(_ffe) == string(_af.MissingCodeRune) && _cgf._edf.BaseFont() == _cgf.Font.BaseFont() {
		return _cgf.charcodeToRunesSimple(charcode)
	}
	return _gab, _ffe
}
func (_abg *TextState) ProcTf(font *TextFont) { _abg.Tf = font }
func (_cda *TextFont) GetCharMetrics(code _ef.CharCode) (float64, float64, bool) {
	if _eb, _efdc := _cda.Font.GetCharMetrics(code); _efdc && _eb.Wx != 0 {
		return _eb.Wx, _eb.Wy, _efdc
	}
	if _cda._edf == nil {
		return 0, 0, false
	}
	_fcc, _baa := _cda._edf.GetCharMetrics(code)
	return _fcc.Wx, _fcc.Wy, _baa && _fcc.Wx != 0
}

type LineJoin int

func (_ggd *TextState) ProcTD(tx, ty float64) { _ggd.Tl = -ty; _ggd.ProcTd(tx, ty) }
func NewTextFont(font *_cb.PdfFont, size float64) (*TextFont, error) {
	_beb := font.FontDescriptor()
	if _beb == nil {
		return nil, _b.New("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072")
	}
	_edb, _gea := _d.GetStream(_beb.FontFile2)
	if !_gea {
		return nil, _b.New("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_efd, _aeb := _d.DecodeStream(_edb)
	if _aeb != nil {
		return nil, _aeb
	}
	_fd, _aeb := _g.Parse(_efd)
	if _aeb != nil {
		return nil, _aeb
	}
	_bdd := font.FontDescriptor().FontName.String()
	_ead := len(_bdd) > 7 && _bdd[6] == '+'
	if !_fd.HasCmap() && (!_f.Contains(font.Encoder().String(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-") || !_ead) {
		return nil, _b.New("\u006e\u006f c\u006d\u0061\u0070 \u0061\u006e\u0064\u0020enc\u006fdi\u006e\u0067\u0020\u0069\u0073\u0020\u006eot\u0020\u0069\u0064\u0065\u006e\u0074\u0069t\u0079")
	}
	return &TextFont{Font: font, Size: size, _ga: _fd}, nil
}

type Pattern interface{ ColorAt(_bg, _fb int) _a.Color }

func (_cbd *TextFont) WithSize(size float64, originalFont *_cb.PdfFont) *TextFont {
	return &TextFont{Font: _cbd.Font, Size: size, _ga: _cbd._ga, _edf: originalFont}
}
func (_dae *TextState) ProcTStar() { _dae.ProcTd(0, -_dae.Tl) }

type Context interface {
	Push()
	Pop()
	Matrix() _e.Matrix
	SetMatrix(_ec _e.Matrix)
	Translate(_ab, _aa float64)
	Scale(_ge, _gd float64)
	Rotate(_abc float64)
	MoveTo(_gc, _eg float64)
	LineTo(_fc, _da float64)
	CubicTo(_gdb, _cbf, _ce, _cg, _ecb, _df float64)
	QuadraticTo(_dc, _aae, _egg, _fa float64)
	NewSubPath()
	ClosePath()
	ClearPath()
	Clip()
	ClipPreserve()
	ResetClip()
	LineWidth() float64
	SetLineWidth(_bd float64)
	SetLineCap(_de LineCap)
	SetLineJoin(_bc LineJoin)
	SetDash(_gf ...float64)
	SetDashOffset(_afb float64)
	Fill()
	FillPreserve()
	Stroke()
	StrokePreserve()
	SetRGBA(_ecc, _dca, _be, _gb float64)
	SetFillRGBA(_db, _ca, _gec, _cbc float64)
	SetFillStyle(_dg Pattern)
	SetFillRule(_ea FillRule)
	SetStrokeRGBA(_dgd, _gg, _abe, _dfc float64)
	SetStrokeStyle(_gde Pattern)
	FillPattern() Pattern
	StrokePattern() Pattern
	TextState() *TextState
	DrawString(_eggc string, _ed _fg.Face, _eccg, _fca float64)
	MeasureString(_dce string, _ff _fg.Face) (_ege, _aaa float64)
	DrawRectangle(_aac, _cf, _bgf, _dfca float64)
	DrawImage(_ba _cd.Image, _bb, _bga int)
	DrawImageAnchored(_cae _cd.Image, _ece, _ae int, _aaf, _ccb float64)
	Height() int
	Width() int
}

func (_cdc *TextState) ProcQ(data []byte, ctx Context) {
	_cdc.ProcTStar()
	_cdc.ProcTj(data, ctx)
}

type TextRenderingMode int

func NewTextFontFromPath(filePath string, size float64) (*TextFont, error) {
	_cgb, _dee := _cb.NewPdfFontFromTTFFile(filePath)
	if _dee != nil {
		return nil, _dee
	}
	return NewTextFont(_cgb, size)
}

const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

func (_cef *TextFont) NewFace(size float64) _fg.Face {
	return _g.NewFace(_cef._ga, &_g.Options{Size: size})
}

func (_gca *TextFont) BytesToCharcodes(data []byte) []_ef.CharCode {
	if _gca._edf != nil {
		return _gca._edf.BytesToCharcodes(data)
	}
	return _gca.Font.BytesToCharcodes(data)
}

func (_gdbb *TextState) ProcTj(data []byte, ctx Context) {
	_ccbf := _gdbb.Tf.Size
	_ade := _gdbb.Th / 100.0
	_dd := _gdbb.GlobalScale
	_ag := _e.NewMatrix(_ccbf*_ade, 0, 0, _ccbf, 0, _gdbb.Ts)
	_eac := ctx.Matrix()
	_afc := _eac.Clone().Mult(_gdbb.Tm.Clone().Mult(_ag)).ScalingFactorY()
	_bee := _gdbb.Tf.NewFace(_afc)
	_dcb := _gdbb.Tf.BytesToCharcodes(data)
	for _, _cbce := range _dcb {
		_ded, _dgc := _gdbb.Tf.CharcodeToRunes(_cbce)
		_ebf := string(_dgc)
		if _ebf == "\u0000" {
			continue
		}
		_afa := _eac.Clone().Mult(_gdbb.Tm.Clone().Mult(_ag))
		_ceg := _afa.ScalingFactorY()
		_afa = _afa.Scale(1/_ceg, -1/_ceg)
		if _gdbb.Tr != TextRenderingModeInvisible {
			ctx.SetMatrix(_afa)
			ctx.DrawString(_ebf, _bee, 0, 0)
			ctx.SetMatrix(_eac)
		}
		_afd := 0.0
		if _ebf == "\u0020" {
			_afd = _gdbb.Tw
		}
		_dcc, _, _ceb := _gdbb.Tf.GetCharMetrics(_ded)
		if _ceb {
			_dcc = _dcc * 0.001 * _ccbf
		} else {
			_dcc, _ = ctx.MeasureString(_ebf, _bee)
			_dcc = _dcc / _dd
		}
		_dab := (_dcc + _gdbb.Tc + _afd) * _ade
		_gdbb.Tm = _gdbb.Tm.Mult(_e.TranslationMatrix(_dab, 0))
	}
}

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
)

func (_gcd *TextState) Reset() { _gcd.Tm = _e.IdentityMatrix(); _gcd.Tlm = _e.IdentityMatrix() }

type TextState struct {
	Tc          float64
	Tw          float64
	Th          float64
	Tl          float64
	Tf          *TextFont
	Ts          float64
	Tm          _e.Matrix
	Tlm         _e.Matrix
	Tr          TextRenderingMode
	GlobalScale float64
}
type Gradient interface {
	Pattern
	AddColorStop(_cc float64, _cde _a.Color)
}

func (_fge *TextState) ProcTm(a, b, c, d, e, f float64) {
	_fge.Tm = _e.NewMatrix(a, b, c, d, e, f)
	_fge.Tlm = _fge.Tm.Clone()
}

func (_bdf *TextState) ProcTd(tx, ty float64) {
	_bdf.Tlm.Concat(_e.TranslationMatrix(tx, ty))
	_bdf.Tm = _bdf.Tlm.Clone()
}

type TextFont struct {
	Font *_cb.PdfFont
	Size float64
	_ga  *_g.Font
	_edf *_cb.PdfFont
}

const (
	LineCapRound LineCap = iota
	LineCapButt
	LineCapSquare
)

func NewTextState() TextState {
	return TextState{Th: 100, Tm: _e.IdentityMatrix(), Tlm: _e.IdentityMatrix()}
}

func (_gag *TextState) Translate(tx, ty float64) {
	_gag.Tm = _gag.Tm.Mult(_e.TranslationMatrix(tx, ty))
}

func (_cdd *TextFont) charcodeToRunesSimple(_eff _ef.CharCode) (_ef.CharCode, []rune) {
	_edd := []_ef.CharCode{_eff}
	if _cdd.Font.IsSimple() && _cdd._ga != nil {
		if _ad := _cdd._ga.Index(rune(_eff)); _ad > 0 {
			return _eff, []rune{rune(_eff)}
		}
	}
	if _cdd._ga != nil && !_cdd._ga.HasCmap() && _f.Contains(_cdd.Font.Encoder().String(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-") {
		if _cdf := _cdd._ga.Index(rune(_eff)); _cdf > 0 {
			return _eff, []rune{rune(_eff)}
		}
	}
	return _eff, _cdd.Font.CharcodesToUnicode(_edd)
}

func (_bgb *TextState) ProcDQ(data []byte, aw, ac float64, ctx Context) {
	_bgb.Tw = aw
	_bgb.Tc = ac
	_bgb.ProcQ(data, ctx)
}

type LineCap int

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

type FillRule int
