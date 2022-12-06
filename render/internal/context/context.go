package context

import (
	_b "errors"
	_f "image"
	_c "image/color"

	_ba "bitbucket.org/shenghui0779/gopdf/core"
	_bdg "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_bd "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_bdge "bitbucket.org/shenghui0779/gopdf/model"
	_ad "github.com/unidoc/freetype/truetype"
	_e "golang.org/x/image/font"
)

func NewTextFont(font *_bdge.PdfFont, size float64) (*TextFont, error) {
	_ada := font.FontDescriptor()
	if _ada == nil {
		return nil, _b.New("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072")
	}
	_ce, _aeb := _ba.GetStream(_ada.FontFile2)
	if !_aeb {
		return nil, _b.New("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_cb, _edb := _ba.DecodeStream(_ce)
	if _edb != nil {
		return nil, _edb
	}
	_fbg, _edb := _ad.Parse(_cb)
	if _edb != nil {
		return nil, _edb
	}
	return &TextFont{Font: font, Size: size, _bcf: _fbg}, nil
}
func (_fbf *TextState) ProcTf(font *TextFont) { _fbf.Tf = font }

type Gradient interface {
	Pattern
	AddColorStop(_adc float64, _cf _c.Color)
}
type TextRenderingMode int
type LineCap int

func (_ddd *TextState) ProcTm(a, b, c, d, e, f float64) {
	_ddd.Tm = _bd.NewMatrix(a, b, c, d, e, f)
	_ddd.Tlm = _ddd.Tm.Clone()
}
func (_agd *TextFont) BytesToCharcodes(data []byte) []_bdg.CharCode {
	if _agd._fbb != nil {
		return _agd._fbb.BytesToCharcodes(data)
	}
	return _agd.Font.BytesToCharcodes(data)
}
func (_dbca *TextState) ProcTj(data []byte, ctx Context) {
	_ebc := _dbca.Tf.Size
	_cfba := _dbca.Th / 100.0
	_gac := _dbca.GlobalScale
	_bbe := _bd.NewMatrix(_ebc*_cfba, 0, 0, _ebc, 0, _dbca.Ts)
	_df := ctx.Matrix()
	_cag := _df.Clone().Mult(_dbca.Tm.Clone().Mult(_bbe)).ScalingFactorY()
	_ceb := _dbca.Tf.NewFace(_cag)
	_gff := _dbca.Tf.BytesToCharcodes(data)
	for _, _cff := range _gff {
		_dbb, _cae := _dbca.Tf.CharcodeToRunes(_cff)
		_ccf := string(_cae)
		if _ccf == "\u0000" {
			continue
		}
		_cgg := _df.Clone().Mult(_dbca.Tm.Clone().Mult(_bbe))
		_aga := _cgg.ScalingFactorY()
		_cgg = _cgg.Scale(1/_aga, -1/_aga)
		if _dbca.Tr != TextRenderingModeInvisible {
			ctx.SetMatrix(_cgg)
			ctx.DrawString(_ccf, _ceb, 0, 0)
			ctx.SetMatrix(_df)
		}
		_dfa := 0.0
		if _ccf == "\u0020" {
			_dfa = _dbca.Tw
		}
		_eda, _, _ddg := _dbca.Tf.GetCharMetrics(_dbb)
		if _ddg {
			_eda = _eda * 0.001 * _ebc
		} else {
			_eda, _ = ctx.MeasureString(_ccf, _ceb)
			_eda = _eda / _gac
		}
		_ef := (_eda + _dbca.Tc + _dfa) * _cfba
		_dbca.Tm = _dbca.Tm.Mult(_bd.TranslationMatrix(_ef, 0))
	}
}

type Pattern interface{ ColorAt(_bc, _d int) _c.Color }

func NewTextFontFromPath(filePath string, size float64) (*TextFont, error) {
	_deeb, _bgbe := _bdge.NewPdfFontFromTTFFile(filePath)
	if _bgbe != nil {
		return nil, _bgbe
	}
	return NewTextFont(_deeb, size)
}
func (_cagg *TextState) ProcDQ(data []byte, aw, ac float64, ctx Context) {
	_cagg.Tw = aw
	_cagg.Tc = ac
	_cagg.ProcQ(data, ctx)
}

const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

func (_ced *TextState) ProcQ(data []byte, ctx Context) { _ced.ProcTStar(); _ced.ProcTj(data, ctx) }

const (
	LineCapRound LineCap = iota
	LineCapButt
	LineCapSquare
)

type TextState struct {
	Tc          float64
	Tw          float64
	Th          float64
	Tl          float64
	Tf          *TextFont
	Ts          float64
	Tm          _bd.Matrix
	Tlm         _bd.Matrix
	Tr          TextRenderingMode
	GlobalScale float64
}
type LineJoin int
type Context interface {
	Push()
	Pop()
	Matrix() _bd.Matrix
	SetMatrix(_g _bd.Matrix)
	Translate(_bcg, _ac float64)
	Scale(_cc, _ca float64)
	Rotate(_ga float64)
	MoveTo(_ge, _aa float64)
	LineTo(_ee, _de float64)
	CubicTo(_eg, _bg, _bca, _ab, _dg, _dgf float64)
	QuadraticTo(_gf, _ed, _bb, _fg float64)
	NewSubPath()
	ClosePath()
	ClearPath()
	Clip()
	ClipPreserve()
	ResetClip()
	LineWidth() float64
	SetLineWidth(_ega float64)
	SetLineCap(_ec LineCap)
	SetLineJoin(_af LineJoin)
	SetDash(_afd ...float64)
	SetDashOffset(_dee float64)
	Fill()
	FillPreserve()
	Stroke()
	StrokePreserve()
	SetRGBA(_gb, _aaa, _ff, _ag float64)
	SetFillRGBA(_ea, _gc, _ae, _ede float64)
	SetFillStyle(_be Pattern)
	SetFillRule(_age FillRule)
	SetStrokeRGBA(_bbf, _ade, _dgd, _bgg float64)
	SetStrokeStyle(_gg Pattern)
	FillPattern() Pattern
	StrokePattern() Pattern
	TextState() *TextState
	DrawString(_dc string, _fe _e.Face, _gfc, _afdd float64)
	MeasureString(_bgb string, _cfb _e.Face) (_bae, _gbc float64)
	DrawRectangle(_dd, _baef, _edf, _cd float64)
	DrawImage(_gga _f.Image, _gea, _cfg int)
	DrawImageAnchored(_fa _f.Image, _ecg, _bge int, _bbc, _fb float64)
	Height() int
	Width() int
}
type FillRule int

func (_egd *TextFont) WithSize(size float64, originalFont *_bdge.PdfFont) *TextFont {
	return &TextFont{Font: _egd.Font, Size: size, _bcf: _egd._bcf, _fbb: originalFont}
}

type TextFont struct {
	Font *_bdge.PdfFont
	Size float64
	_bcf *_ad.Font
	_fbb *_bdge.PdfFont
}

func (_ace *TextState) ProcTStar() { _ace.ProcTd(0, -_ace.Tl) }
func (_aac *TextState) Translate(tx, ty float64) {
	_aac.Tm = _aac.Tm.Mult(_bd.TranslationMatrix(tx, ty))
}
func (_eb *TextFont) GetCharMetrics(code _bdg.CharCode) (float64, float64, bool) {
	if _db, _dbc := _eb.Font.GetCharMetrics(code); _dbc && _db.Wx != 0 {
		return _db.Wx, _db.Wy, _dbc
	}
	if _eb._fbb == nil {
		return 0, 0, false
	}
	_fab, _adf := _eb._fbb.GetCharMetrics(code)
	return _fab.Wx, _fab.Wy, _adf && _fab.Wx != 0
}
func (_ffa *TextState) Reset() {
	_ffa.Tm = _bd.IdentityMatrix()
	_ffa.Tlm = _bd.IdentityMatrix()
}
func (_cdg *TextState) ProcTd(tx, ty float64) {
	_cdg.Tlm.Concat(_bd.TranslationMatrix(tx, ty))
	_cdg.Tm = _cdg.Tlm.Clone()
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
const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
)

func (_aca *TextFont) NewFace(size float64) _e.Face {
	return _ad.NewFace(_aca._bcf, &_ad.Options{Size: size})
}
func (_bbg *TextFont) CharcodeToRunes(charcode _bdg.CharCode) (_bdg.CharCode, []rune) {
	_ffg := []_bdg.CharCode{charcode}
	if _bbg._fbb == nil || _bbg._fbb == _bbg.Font {
		if _bbg.Font.IsSimple() && _bbg._bcf != nil {
			if _dgc := _bbg._bcf.Index(rune(charcode)); _dgc > 0 {
				return charcode, []rune{rune(charcode)}
			}
		}
		return charcode, _bbg.Font.CharcodesToUnicode(_ffg)
	}
	_cda := _bbg._fbb.CharcodesToUnicode(_ffg)
	_bcc, _ := _bbg.Font.RunesToCharcodeBytes(_cda)
	_aed := _bbg.Font.BytesToCharcodes(_bcc)
	_gae := charcode
	if len(_aed) > 0 && _aed[0] != 0 {
		_gae = _aed[0]
	}
	return _gae, _cda
}
func NewTextState() TextState {
	return TextState{Th: 100, Tm: _bd.IdentityMatrix(), Tlm: _bd.IdentityMatrix()}
}
func (_egb *TextState) ProcTD(tx, ty float64) {
	_egb.Tl = -ty
	_egb.ProcTd(tx, ty)
}
