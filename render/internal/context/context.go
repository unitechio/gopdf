package context

import (
	_c "errors"
	_df "image"
	_dc "image/color"
	_b "strings"

	_fg "bitbucket.org/shenghui0779/gopdf/core"
	_cb "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_a "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_dfa "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_f "bitbucket.org/shenghui0779/gopdf/model"
	_cd "github.com/unidoc/freetype/truetype"
	_cg "golang.org/x/image/font"
)

func (_bgd *TextState) ProcTm(a, b, c, d, e, f float64) {
	_bgd.Tm = _dfa.NewMatrix(a, b, c, d, e, f)
	_bgd.Tlm = _bgd.Tm.Clone()
}

type Pattern interface {
	ColorAt(_ce, _da int) _dc.Color
}
type TextFont struct {
	Font  *_f.PdfFont
	Size  float64
	_cged *_cd.Font
	_gba  *_f.PdfFont
}

func (_gbd *TextState) ProcTD(tx, ty float64) { _gbd.Tl = -ty; _gbd.ProcTd(tx, ty) }

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
)

func (_fbd *TextFont) CharcodeToRunes(charcode _a.CharCode) (_a.CharCode, []rune) {
	_bgg := []_a.CharCode{charcode}
	if _fbd._gba == nil || _fbd._gba == _fbd.Font {
		return _fbd.charcodeToRunesSimple(charcode)
	}
	_fca := _fbd._gba.CharcodesToUnicode(_bgg)
	_cc, _ := _fbd.Font.RunesToCharcodeBytes(_fca)
	_cbaa := _fbd.Font.BytesToCharcodes(_cc)
	_de := charcode
	if len(_cbaa) > 0 && _cbaa[0] != 0 {
		_de = _cbaa[0]
	}
	if string(_fca) == string(_cb.MissingCodeRune) && _fbd._gba.BaseFont() == _fbd.Font.BaseFont() {
		return _fbd.charcodeToRunesSimple(charcode)
	}
	return _de, _fca
}
func (_gee *TextState) ProcTd(tx, ty float64) {
	_gee.Tlm.Concat(_dfa.TranslationMatrix(tx, ty))
	_gee.Tm = _gee.Tlm.Clone()
}
func (_bag *TextFont) GetCharMetrics(code _a.CharCode) (float64, float64, bool) {
	if _cbg, _gdb := _bag.Font.GetCharMetrics(code); _gdb && _cbg.Wx != 0 {
		return _cbg.Wx, _cbg.Wy, _gdb
	}
	if _bag._gba == nil {
		return 0, 0, false
	}
	_ccb, _dgac := _bag._gba.GetCharMetrics(code)
	return _ccb.Wx, _ccb.Wy, _dgac && _ccb.Wx != 0
}
func (_gcfe *TextState) ProcTf(font *TextFont) { _gcfe.Tf = font }
func (_dac *TextState) Translate(tx, ty float64) {
	_dac.Tm = _dac.Tm.Mult(_dfa.TranslationMatrix(tx, ty))
}
func (_ege *TextFont) NewFace(size float64) _cg.Face {
	return _cd.NewFace(_ege._cged, &_cd.Options{Size: size})
}

type Context interface {
	Push()
	Pop()
	Matrix() _dfa.Matrix
	SetMatrix(_ag _dfa.Matrix)
	Translate(_g, _e float64)
	Scale(_dd, _dcf float64)
	Rotate(_ea float64)
	MoveTo(_ba, _ff float64)
	LineTo(_fc, _db float64)
	CubicTo(_ffa, _gb, _eaa, _dfg, _ed, _bf float64)
	QuadraticTo(_bg, _dg, _bdf, _edc float64)
	NewSubPath()
	ClosePath()
	ClearPath()
	Clip()
	ClipPreserve()
	ResetClip()
	LineWidth() float64
	SetLineWidth(_fa float64)
	SetLineCap(_ddb LineCap)
	SetLineJoin(_ef LineJoin)
	SetDash(_dfb ...float64)
	SetDashOffset(_ge float64)
	Fill()
	FillPreserve()
	Stroke()
	StrokePreserve()
	SetRGBA(_bfb, _bad, _ae, _cge float64)
	SetFillRGBA(_cf, _af, _gc, _gcb float64)
	SetFillStyle(_dga Pattern)
	SetFillRule(_aeg FillRule)
	SetStrokeRGBA(_gf, _gd, _aa, _cbc float64)
	SetStrokeStyle(_daa Pattern)
	FillPattern() Pattern
	StrokePattern() Pattern
	TextState() *TextState
	DrawString(_ffe string, _cgb _cg.Face, _dfe, _eg float64)
	MeasureString(_egg string, _ee _cg.Face) (_ac, _afg float64)
	DrawRectangle(_aac, _gfe, _bfc, _bb float64)
	DrawImage(_fcd _df.Image, _ec, _cee int)
	DrawImageAnchored(_fae _df.Image, _bff, _ecc int, _ca, _efb float64)
	Height() int
	Width() int
}

func (_deg *TextState) ProcTStar() { _deg.ProcTd(0, -_deg.Tl) }
func (_ddbg *TextFont) BytesToCharcodes(data []byte) []_a.CharCode {
	if _ddbg._gba != nil {
		return _ddbg._gba.BytesToCharcodes(data)
	}
	return _ddbg.Font.BytesToCharcodes(data)
}
func (_bc *TextFont) WithSize(size float64, originalFont *_f.PdfFont) *TextFont {
	return &TextFont{Font: _bc.Font, Size: size, _cged: _bc._cged, _gba: originalFont}
}

type Gradient interface {
	Pattern
	AddColorStop(_cba float64, _bd _dc.Color)
}

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

func (_cdb *TextState) ProcQ(data []byte, ctx Context) {
	_cdb.ProcTStar()
	_cdb.ProcTj(data, ctx)
}

type LineCap int

func (_eag *TextState) ProcDQ(data []byte, aw, ac float64, ctx Context) {
	_eag.Tw = aw
	_eag.Tc = ac
	_eag.ProcQ(data, ctx)
}

type TextState struct {
	Tc          float64
	Tw          float64
	Th          float64
	Tl          float64
	Tf          *TextFont
	Ts          float64
	Tm          _dfa.Matrix
	Tlm         _dfa.Matrix
	Tr          TextRenderingMode
	GlobalScale float64
}

func (_acc *TextFont) charcodeToRunesSimple(_cbd _a.CharCode) (_a.CharCode, []rune) {
	_bdc := []_a.CharCode{_cbd}
	if _acc.Font.IsSimple() && _acc._cged != nil {
		if _ddg := _acc._cged.Index(rune(_cbd)); _ddg > 0 {
			return _cbd, []rune{rune(_cbd)}
		}
	}
	if _acc._cged != nil && !_acc._cged.HasCmap() && _b.Contains(_acc.Font.Encoder().String(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-") {
		if _fge := _acc._cged.Index(rune(_cbd)); _fge > 0 {
			return _cbd, []rune{rune(_cbd)}
		}
	}
	return _cbd, _acc.Font.CharcodesToUnicode(_bdc)
}
func (_be *TextState) ProcTj(data []byte, ctx Context) {
	_fce := _be.Tf.Size
	_cgab := _be.Th / 100.0
	_ad := _be.GlobalScale
	_fgg := _dfa.NewMatrix(_fce*_cgab, 0, 0, _fce, 0, _be.Ts)
	_deb := ctx.Matrix()
	_dbc := _deb.Clone().Mult(_be.Tm.Clone().Mult(_fgg)).ScalingFactorY()
	_aag := _be.Tf.NewFace(_dbc)
	_ged := _be.Tf.BytesToCharcodes(data)
	for _, _gcf := range _ged {
		_bab, _fba := _be.Tf.CharcodeToRunes(_gcf)
		_eee := string(_fba)
		if _eee == "\u0000" {
			continue
		}
		_gcc := _deb.Clone().Mult(_be.Tm.Clone().Mult(_fgg))
		_fgb := _gcc.ScalingFactorY()
		_gcc = _gcc.Scale(1/_fgb, -1/_fgb)
		if _be.Tr != TextRenderingModeInvisible {
			ctx.SetMatrix(_gcc)
			ctx.DrawString(_eee, _aag, 0, 0)
			ctx.SetMatrix(_deb)
		}
		_eef := 0.0
		if _eee == "\u0020" {
			_eef = _be.Tw
		}
		_ece, _, _egf := _be.Tf.GetCharMetrics(_bab)
		if _egf {
			_ece = _ece * 0.001 * _fce
		} else {
			_ece, _ = ctx.MeasureString(_eee, _aag)
			_ece = _ece / _ad
		}
		_eb := (_ece + _be.Tc + _eef) * _cgab
		_be.Tm = _be.Tm.Mult(_dfa.TranslationMatrix(_eb, 0))
	}
}

type FillRule int

func NewTextState() TextState {
	return TextState{Th: 100, Tm: _dfa.IdentityMatrix(), Tlm: _dfa.IdentityMatrix()}
}
func NewTextFontFromPath(filePath string, size float64) (*TextFont, error) {
	_egb, _fb := _f.NewPdfFontFromTTFFile(filePath)
	if _fb != nil {
		return nil, _fb
	}
	return NewTextFont(_egb, size)
}
func (_gcfc *TextState) Reset() { _gcfc.Tm = _dfa.IdentityMatrix(); _gcfc.Tlm = _dfa.IdentityMatrix() }

const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

type LineJoin int

const (
	LineCapRound LineCap = iota
	LineCapButt
	LineCapSquare
)

func NewTextFont(font *_f.PdfFont, size float64) (*TextFont, error) {
	_cgf := font.FontDescriptor()
	if _cgf == nil {
		return nil, _c.New("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072")
	}
	_acg, _dfec := _fg.GetStream(_cgf.FontFile2)
	if !_dfec {
		return nil, _c.New("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_eac, _cfe := _fg.DecodeStream(_acg)
	if _cfe != nil {
		return nil, _cfe
	}
	_dfc, _cfe := _cd.Parse(_eac)
	if _cfe != nil {
		return nil, _cfe
	}
	_gcd := font.FontDescriptor().FontName.String()
	_aae := len(_gcd) > 7 && _gcd[6] == '+'
	if !_dfc.HasCmap() && (!_b.Contains(font.Encoder().String(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-") || !_aae) {
		return nil, _c.New("\u006e\u006f c\u006d\u0061\u0070 \u0061\u006e\u0064\u0020enc\u006fdi\u006e\u0067\u0020\u0069\u0073\u0020\u006eot\u0020\u0069\u0064\u0065\u006e\u0074\u0069t\u0079")
	}
	return &TextFont{Font: font, Size: size, _cged: _dfc}, nil
}

type TextRenderingMode int
