package imagerender

import (
	_dd "errors"
	_de "fmt"
	_fb "image"
	_da "image/color"
	_a "image/draw"
	_b "math"
	_c "sort"
	_d "strings"

	_fbb "bitbucket.org/shenghui0779/gopdf/common"
	_ca "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_g "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_ed "github.com/unidoc/freetype/raster"
	_e "golang.org/x/image/draw"
	_ec "golang.org/x/image/font"
	_fd "golang.org/x/image/math/f64"
	_bg "golang.org/x/image/math/fixed"
)

func (_bbd *Context) NewSubPath() {
	if _bbd._df {
		_bbd._eea.Add1(_eda(_bbd._dea))
	}
	_bbd._df = false
}
func (_baa *Context) Height() int { return _baa._ggd }
func _abdg(_gegd _fb.Image, _eede repeatOp) _g.Pattern {
	return &surfacePattern{_fcaa: _gegd, _adee: _eede}
}
func (_effg *Context) SetLineWidth(lineWidth float64) { _effg._bgd = lineWidth }
func (_fadb *Context) capper() _ed.Capper {
	switch _fadb._eaag {
	case _g.LineCapButt:
		return _ed.ButtCapper
	case _g.LineCapRound:
		return _ed.RoundCapper
	case _g.LineCapSquare:
		return _ed.SquareCapper
	}
	return nil
}
func _cc(_bee, _ded, _af, _dag, _gc, _ad float64) []_ca.Point {
	_fad := (_b.Hypot(_af-_bee, _dag-_ded) + _b.Hypot(_gc-_af, _ad-_dag))
	_fbbg := int(_fad + 0.5)
	if _fbbg < 4 {
		_fbbg = 4
	}
	_dg := float64(_fbbg) - 1
	_bc := make([]_ca.Point, _fbbg)
	for _ef := 0; _ef < _fbbg; _ef++ {
		_aab := float64(_ef) / _dg
		_cd, _cf := _gd(_bee, _ded, _af, _dag, _gc, _ad, _aab)
		_bc[_ef] = _ca.NewPoint(_cd, _cf)
	}
	return _bc
}
func (_gbe *Context) drawString(_bca string, _aaa _ec.Face, _gcba, _edb float64) {
	_dgd := &_ec.Drawer{Src: _fb.NewUniform(_gbe._geg), Face: _aaa, Dot: _eda(_ca.NewPoint(_gcba, _edb))}
	_eeg := rune(-1)
	for _, _daggb := range _bca {
		if _eeg >= 0 {
			_dgd.Dot.X += _dgd.Face.Kern(_eeg, _daggb)
		}
		_gcag, _cedf, _bfe, _fgb, _fbgc := _dgd.Face.Glyph(_dgd.Dot, _daggb)
		if !_fbgc {
			continue
		}
		_bbf := _gcag.Sub(_gcag.Min)
		_ebc := _fb.NewRGBA(_bbf)
		_e.DrawMask(_ebc, _bbf, _dgd.Src, _fb.Point{}, _cedf, _bfe, _e.Over)
		var _acc *_e.Options
		if _gbe._aca != nil {
			_acc = &_e.Options{DstMask: _gbe._aca, DstMaskP: _fb.Point{}}
		}
		_ddf := _gbe._afg.Clone().Translate(float64(_gcag.Min.X), float64(_gcag.Min.Y))
		_ddb := _fd.Aff3{_ddf[0], _ddf[3], _ddf[6], _ddf[1], _ddf[4], _ddf[7]}
		_e.BiLinear.Transform(_gbe._agd, _ddb, _ebc, _bbf, _e.Over, _acc)
		_dgd.Dot.X += _fgb
		_eeg = _daggb
	}
}
func (_cfa *Context) AsMask() *_fb.Alpha {
	_dad := _fb.NewAlpha(_cfa._agd.Bounds())
	_e.Draw(_dad, _cfa._agd.Bounds(), _cfa._agd, _fb.Point{}, _e.Src)
	return _dad
}
func (_dc *Context) setFillAndStrokeColor(_ecf _da.Color) {
	_dc._geg = _ecf
	_dc._aabb = _cecd(_ecf)
	_dc._gdc = _cecd(_ecf)
}
func _gd(_fa, _ae, _aa, _ff, _eg, _dda, _ge float64) (_be, _dee float64) {
	_def := 1 - _ge
	_gf := _def * _def
	_deec := 2 * _def * _ge
	_bgc := _ge * _ge
	_be = _gf*_fa + _deec*_aa + _bgc*_eg
	_dee = _gf*_ae + _deec*_ff + _bgc*_dda
	return
}
func (_eef *Context) ResetClip() { _eef._aca = nil }
func _aegb(_ecae _bg.Int26_6) float64 {
	const _baf, _bdga = 6, 1<<6 - 1
	if _ecae >= 0 {
		return float64(_ecae>>_baf) + float64(_ecae&_bdga)/64
	}
	_ecae = -_ecae
	if _ecae >= 0 {
		return -(float64(_ecae>>_baf) + float64(_ecae&_bdga)/64)
	}
	return 0
}
func (_gcf *Context) StrokePreserve() {
	var _aff _ed.Painter
	if _gcf._aca == nil {
		if _cac, _cebb := _gcf._gdc.(*solidPattern); _cebb {
			_gfe := _ed.NewRGBAPainter(_gcf._agd)
			_gfe.SetColor(_cac._bbfc)
			_aff = _gfe
		}
	}
	if _aff == nil {
		_aff = _aagb(_gcf._agd, _gcf._aca, _gcf._gdc)
	}
	_gcf.stroke(_aff)
}
func NewContextForImage(im _fb.Image) *Context                     { return NewContextForRGBA(_cdc(im)) }
func (_gac *Context) Transform(x, y float64) (_dgec, _dga float64) { return _gac._afg.Transform(x, y) }
func (_bfc *Context) SetMask(mask *_fb.Alpha) error {
	if mask.Bounds().Size() != _bfc._agd.Bounds().Size() {
		return _dd.New("\u006d\u0061\u0073\u006b\u0020\u0073i\u007a\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068 \u0063\u006f\u006e\u0074\u0065\u0078\u0074 \u0073\u0069\u007a\u0065")
	}
	_bfc._aca = mask
	return nil
}
func (_gca *Context) drawRegularPolygon(_bdbbg int, _adce, _gefa, _dge, _gdb float64) {
	_dcc := 2 * _b.Pi / float64(_bdbbg)
	_gdb -= _b.Pi / 2
	if _bdbbg%2 == 0 {
		_gdb += _dcc / 2
	}
	_gca.NewSubPath()
	for _efc := 0; _efc < _bdbbg; _efc++ {
		_cea := _gdb + _dcc*float64(_efc)
		_gca.LineTo(_adce+_dge*_b.Cos(_cea), _gefa+_dge*_b.Sin(_cea))
	}
	_gca.ClosePath()
}
func (_fcef *Context) DrawPoint(x, y, r float64) {
	_fcef.Push()
	_ecdf, _gdd := _fcef.Transform(x, y)
	_fcef.Identity()
	_fcef.DrawCircle(_ecdf, _gdd, r)
	_fcef.Pop()
}

type solidPattern struct{ _bbfc _da.Color }

func (_ecd *Context) Image() _fb.Image { return _ecd._agd }
func (_aea *Context) Stroke()          { _aea.StrokePreserve(); _aea.ClearPath() }
func (_eeag *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	if !_eeag._df {
		_eeag.MoveTo(x1, y1)
	}
	_eba, _gdg := _eeag._efa.X, _eeag._efa.Y
	x1, y1 = _eeag.Transform(x1, y1)
	x2, y2 = _eeag.Transform(x2, y2)
	x3, y3 = _eeag.Transform(x3, y3)
	_fff := _adb(_eba, _gdg, x1, y1, x2, y2, x3, y3)
	_faaf := _eda(_eeag._efa)
	for _, _dde := range _fff[1:] {
		_afcc := _eda(_dde)
		if _afcc == _faaf {
			continue
		}
		_faaf = _afcc
		_eeag._db.Add1(_afcc)
		_eeag._eea.Add1(_afcc)
		_eeag._efa = _dde
	}
}
func (_eedg *radialGradient) AddColorStop(offset float64, color _da.Color) {
	_eedg._gefb = append(_eedg._gefb, stop{_dede: offset, _gdbe: color})
	_c.Sort(_eedg._gefb)
}
func (_dba *Context) ClosePath() {
	if _dba._df {
		_aba := _eda(_dba._dea)
		_dba._db.Add1(_aba)
		_dba._eea.Add1(_aba)
		_dba._efa = _dba._dea
	}
}

type circle struct{ _ffb, _ebca, _gec float64 }

func (_cbfc *linearGradient) ColorAt(x, y int) _da.Color {
	if len(_cbfc._cgf) == 0 {
		return _da.Transparent
	}
	_eabb, _dfg := float64(x), float64(y)
	_cdf, _eeaga, _ebf, _aaae := _cbfc._bab, _cbfc._fef, _cbfc._dca, _cbfc._cded
	_ebbb, _ccaf := _ebf-_cdf, _aaae-_eeaga
	if _ccaf == 0 && _ebbb != 0 {
		return _bcf((_eabb-_cdf)/_ebbb, _cbfc._cgf)
	}
	if _ebbb == 0 && _ccaf != 0 {
		return _bcf((_dfg-_eeaga)/_ccaf, _cbfc._cgf)
	}
	_bcg := _ebbb*(_eabb-_cdf) + _ccaf*(_dfg-_eeaga)
	if _bcg < 0 {
		return _cbfc._cgf[0]._gdbe
	}
	_dffg := _b.Hypot(_ebbb, _ccaf)
	_bagg := ((_eabb-_cdf)*-_ccaf + (_dfg-_eeaga)*_ebbb) / (_dffg * _dffg)
	_faca, _egf := _cdf+_bagg*-_ccaf, _eeaga+_bagg*_ebbb
	_eac := _b.Hypot(_eabb-_faca, _dfg-_egf) / _dffg
	return _bcf(_eac, _cbfc._cgf)
}

type patternPainter struct {
	_bbbg *_fb.RGBA
	_dfdg *_fb.Alpha
	_fgbe _g.Pattern
}

func _abe(_cdaa string) (_fffd, _ddd, _fda, _cgfc int) {
	_cdaa = _d.TrimPrefix(_cdaa, "\u0023")
	_cgfc = 255
	if len(_cdaa) == 3 {
		_bdfg := "\u00251\u0078\u0025\u0031\u0078\u0025\u0031x"
		_de.Sscanf(_cdaa, _bdfg, &_fffd, &_ddd, &_fda)
		_fffd |= _fffd << 4
		_ddd |= _ddd << 4
		_fda |= _fda << 4
	}
	if len(_cdaa) == 6 {
		_eddd := "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_de.Sscanf(_cdaa, _eddd, &_fffd, &_ddd, &_fda)
	}
	if len(_cdaa) == 8 {
		_bdag := "\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_de.Sscanf(_cdaa, _bdag, &_fffd, &_ddd, &_fda, &_cgfc)
	}
	return
}
func (_gfg stops) Len() int { return len(_gfg) }
func (_adc *Context) DrawRoundedRectangle(x, y, w, h, r float64) {
	_defg, _adg, _ddg, _eec := x, x+r, x+w-r, x+w
	_beff, _edg, _dbc, _gge := y, y+r, y+h-r, y+h
	_adc.NewSubPath()
	_adc.MoveTo(_adg, _beff)
	_adc.LineTo(_ddg, _beff)
	_adc.DrawArc(_ddg, _edg, r, _fbf(270), _fbf(360))
	_adc.LineTo(_eec, _dbc)
	_adc.DrawArc(_ddg, _dbc, r, _fbf(0), _fbf(90))
	_adc.LineTo(_adg, _gge)
	_adc.DrawArc(_adg, _dbc, r, _fbf(90), _fbf(180))
	_adc.LineTo(_defg, _edg)
	_adc.DrawArc(_adg, _edg, r, _fbf(180), _fbf(270))
	_adc.ClosePath()
}
func (_egeb *solidPattern) ColorAt(x, y int) _da.Color { return _egeb._bbfc }

type Context struct {
	_eab  int
	_ggd  int
	_bbe  *_ed.Rasterizer
	_agd  *_fb.RGBA
	_aca  *_fb.Alpha
	_geg  _da.Color
	_aabb _g.Pattern
	_gdc  _g.Pattern
	_db   _ed.Path
	_eea  _ed.Path
	_dea  _ca.Point
	_efa  _ca.Point
	_df   bool
	_eff  []float64
	_faa  float64
	_bgd  float64
	_eaag _g.LineCap
	_cae  _g.LineJoin
	_agg  _g.FillRule
	_afg  _ca.Matrix
	_beeb _g.TextState
	_ege  []*Context
}

func (_cca *Context) DrawLine(x1, y1, x2, y2 float64) {
	_cca.MoveTo(x1, y1)
	_cca.LineTo(x2, y2)
}
func _cecd(_fcc _da.Color) _g.Pattern    { return &solidPattern{_bbfc: _fcc} }
func (_fdc *Context) Matrix() _ca.Matrix { return _fdc._afg }
func (_fgg *Context) Shear(x, y float64) { _fgg._afg.Shear(x, y) }
func (_afgg *Context) SetFillStyle(pattern _g.Pattern) {
	if _dab, _gbg := pattern.(*solidPattern); _gbg {
		_afgg._geg = _dab._bbfc
	}
	_afgg._aabb = pattern
}
func (_dec *Context) LineTo(x, y float64) {
	if !_dec._df {
		_dec.MoveTo(x, y)
	} else {
		x, y = _dec.Transform(x, y)
		_egb := _ca.NewPoint(x, y)
		_ebb := _eda(_egb)
		_dec._db.Add1(_ebb)
		_dec._eea.Add1(_ebb)
		_dec._efa = _egb
	}
}
func _bgdg(_ccab, _efad uint32, _ecag float64) uint8 {
	return uint8(int32(float64(_ccab)*(1.0-_ecag)+float64(_efad)*_ecag) >> 8)
}
func (_dgf *Context) SetHexColor(x string) {
	_bdfc, _cg, _ggf, _cge := _abe(x)
	_dgf.SetRGBA255(_bdfc, _cg, _ggf, _cge)
}
func (_dbe *Context) ClipPreserve() {
	_fec := _fb.NewAlpha(_fb.Rect(0, 0, _dbe._eab, _dbe._ggd))
	_dfdb := _ed.NewAlphaOverPainter(_fec)
	_dbe.fill(_dfdb)
	if _dbe._aca == nil {
		_dbe._aca = _fec
	} else {
		_fee := _fb.NewAlpha(_fb.Rect(0, 0, _dbe._eab, _dbe._ggd))
		_e.DrawMask(_fee, _fee.Bounds(), _fec, _fb.Point{}, _dbe._aca, _fb.Point{}, _e.Over)
		_dbe._aca = _fee
	}
}
func (_cdb *Context) Fill() { _cdb.FillPreserve(); _cdb.ClearPath() }
func _bea(_bdfa, _dgag _da.Color, _bac float64) _da.Color {
	_geb, _bce, _ecea, _daa := _bdfa.RGBA()
	_cfg, _abaa, _agcd, _gaa := _dgag.RGBA()
	return _da.RGBA{_bgdg(_geb, _cfg, _bac), _bgdg(_bce, _abaa, _bac), _bgdg(_ecea, _agcd, _bac), _bgdg(_daa, _gaa, _bac)}
}
func (_cgb *Context) SetRGBA(r, g, b, a float64) {
	_cgb._geg = _da.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_cgb.setFillAndStrokeColor(_cgb._geg)
}
func (_gb *Context) SetDashOffset(offset float64) { _gb._faa = offset }
func NewContext(width, height int) *Context {
	return NewContextForRGBA(_fb.NewRGBA(_fb.Rect(0, 0, width, height)))
}
func (_accc *Context) ShearAbout(sx, sy, x, y float64) {
	_accc.Translate(x, y)
	_accc.Shear(sx, sy)
	_accc.Translate(-x, -y)
}
func _cabd(_fgbc _ed.Path) [][]_ca.Point {
	var _aee [][]_ca.Point
	var _egcca []_ca.Point
	var _cgfd, _ggeb float64
	for _adcd := 0; _adcd < len(_fgbc); {
		switch _fgbc[_adcd] {
		case 0:
			if len(_egcca) > 0 {
				_aee = append(_aee, _egcca)
				_egcca = nil
			}
			_ccf := _aegb(_fgbc[_adcd+1])
			_cgg := _aegb(_fgbc[_adcd+2])
			_egcca = append(_egcca, _ca.NewPoint(_ccf, _cgg))
			_cgfd, _ggeb = _ccf, _cgg
			_adcd += 4
		case 1:
			_aga := _aegb(_fgbc[_adcd+1])
			_bgb := _aegb(_fgbc[_adcd+2])
			_egcca = append(_egcca, _ca.NewPoint(_aga, _bgb))
			_cgfd, _ggeb = _aga, _bgb
			_adcd += 4
		case 2:
			_aegd := _aegb(_fgbc[_adcd+1])
			_ffae := _aegb(_fgbc[_adcd+2])
			_efgc := _aegb(_fgbc[_adcd+3])
			_ggfg := _aegb(_fgbc[_adcd+4])
			_eged := _cc(_cgfd, _ggeb, _aegd, _ffae, _efgc, _ggfg)
			_egcca = append(_egcca, _eged...)
			_cgfd, _ggeb = _efgc, _ggfg
			_adcd += 6
		case 3:
			_cbg := _aegb(_fgbc[_adcd+1])
			_baaa := _aegb(_fgbc[_adcd+2])
			_agga := _aegb(_fgbc[_adcd+3])
			_faee := _aegb(_fgbc[_adcd+4])
			_gbegf := _aegb(_fgbc[_adcd+5])
			_fcd := _aegb(_fgbc[_adcd+6])
			_age := _adb(_cgfd, _ggeb, _cbg, _baaa, _agga, _faee, _gbegf, _fcd)
			_egcca = append(_egcca, _age...)
			_cgfd, _ggeb = _gbegf, _fcd
			_adcd += 8
		default:
			_fbb.Log.Debug("\u0057\u0041\u0052\u004e: \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0074\u0068\u003a\u0020%\u0076", _fgbc)
			return _aee
		}
	}
	if len(_egcca) > 0 {
		_aee = append(_aee, _egcca)
	}
	return _aee
}
func (_gde *Context) SetDash(dashes ...float64) { _gde._eff = dashes }
func (_eee *Context) DrawStringAnchored(s string, face _ec.Face, x, y, ax, ay float64) {
	_ccda, _cgeg := _eee.MeasureString(s, face)
	_eee.drawString(s, face, x-ax*_ccda, y+ay*_cgeg)
}
func (_fffg stops) Swap(i, j int) { _fffg[i], _fffg[j] = _fffg[j], _fffg[i] }
func _bd(_cde, _dgg, _gg, _bf, _ea, _daf, _egc, _bdf, _fe float64) (_ga, _bb float64) {
	_bdb := 1 - _fe
	_ac := _bdb * _bdb * _bdb
	_dae := 3 * _bdb * _bdb * _fe
	_ba := 3 * _bdb * _fe * _fe
	_ccd := _fe * _fe * _fe
	_ga = _ac*_cde + _dae*_gg + _ba*_ea + _ccd*_egc
	_bb = _ac*_dgg + _dae*_bf + _ba*_daf + _ccd*_bdf
	return
}
func _dagb(_caae float64) _bg.Int26_6 { return _bg.Int26_6(_caae * 64) }
func NewLinearGradient(x0, y0, x1, y1 float64) _g.Gradient {
	_fdfd := &linearGradient{_bab: x0, _fef: y0, _dca: x1, _cded: y1}
	return _fdfd
}
func (_cegg *patternPainter) Paint(ss []_ed.Span, done bool) {
	_cbfcc := _cegg._bbbg.Bounds()
	for _, _afb := range ss {
		if _afb.Y < _cbfcc.Min.Y {
			continue
		}
		if _afb.Y >= _cbfcc.Max.Y {
			return
		}
		if _afb.X0 < _cbfcc.Min.X {
			_afb.X0 = _cbfcc.Min.X
		}
		if _afb.X1 > _cbfcc.Max.X {
			_afb.X1 = _cbfcc.Max.X
		}
		if _afb.X0 >= _afb.X1 {
			continue
		}
		const _dcab = 1<<16 - 1
		_edcg := _afb.Y - _cegg._bbbg.Rect.Min.Y
		_dfge := _afb.X0 - _cegg._bbbg.Rect.Min.X
		_ffaa := (_afb.Y-_cegg._bbbg.Rect.Min.Y)*_cegg._bbbg.Stride + (_afb.X0-_cegg._bbbg.Rect.Min.X)*4
		_fdg := _ffaa + (_afb.X1-_afb.X0)*4
		for _cdd, _cadcb := _ffaa, _dfge; _cdd < _fdg; _cdd, _cadcb = _cdd+4, _cadcb+1 {
			_fgf := _afb.Alpha
			if _cegg._dfdg != nil {
				_fgf = _fgf * uint32(_cegg._dfdg.AlphaAt(_cadcb, _edcg).A) / 255
				if _fgf == 0 {
					continue
				}
			}
			_afa := _cegg._fgbe.ColorAt(_cadcb, _edcg)
			_dceef, _efd, _bffb, _fgbae := _afa.RGBA()
			_edbbd := uint32(_cegg._bbbg.Pix[_cdd+0])
			_cbe := uint32(_cegg._bbbg.Pix[_cdd+1])
			_aeag := uint32(_cegg._bbbg.Pix[_cdd+2])
			_cbad := uint32(_cegg._bbbg.Pix[_cdd+3])
			_aag := (_dcab - (_fgbae * _fgf / _dcab)) * 0x101
			_cegg._bbbg.Pix[_cdd+0] = uint8((_edbbd*_aag + _dceef*_fgf) / _dcab >> 8)
			_cegg._bbbg.Pix[_cdd+1] = uint8((_cbe*_aag + _efd*_fgf) / _dcab >> 8)
			_cegg._bbbg.Pix[_cdd+2] = uint8((_aeag*_aag + _bffb*_fgf) / _dcab >> 8)
			_cegg._bbbg.Pix[_cdd+3] = uint8((_cbad*_aag + _fgbae*_fgf) / _dcab >> 8)
		}
	}
}
func (_aeg *Context) Pop() {
	_cad := *_aeg
	_dff := _aeg._ege
	_bff := _dff[len(_dff)-1]
	*_aeg = *_bff
	_aeg._db = _cad._db
	_aeg._eea = _cad._eea
	_aeg._dea = _cad._dea
	_aeg._efa = _cad._efa
	_aeg._df = _cad._df
}
func (_gea *Context) SetFillRule(fillRule _g.FillRule) { _gea._agg = fillRule }
func (_cegb *Context) RotateAbout(angle, x, y float64) {
	_cegb.Translate(x, y)
	_cegb.Rotate(angle)
	_cegb.Translate(-x, -y)
}
func (_cag *Context) SetLineCap(lineCap _g.LineCap) { _cag._eaag = lineCap }
func NewContextForRGBA(im *_fb.RGBA) *Context {
	_gcg := im.Bounds().Size().X
	_cda := im.Bounds().Size().Y
	return &Context{_eab: _gcg, _ggd: _cda, _bbe: _ed.NewRasterizer(_gcg, _cda), _agd: im, _geg: _da.Transparent, _aabb: _bdbf, _gdc: _gcb, _bgd: 1, _agg: _g.FillRuleWinding, _afg: _ca.IdentityMatrix(), _beeb: _g.NewTextState()}
}
func NewRadialGradient(x0, y0, r0, x1, y1, r1 float64) _g.Gradient {
	_cfdb := circle{x0, y0, r0}
	_ebd := circle{x1, y1, r1}
	_ebe := circle{x1 - x0, y1 - y0, r1 - r0}
	_gfc := _acd(_ebe._ffb, _ebe._ebca, -_ebe._gec, _ebe._ffb, _ebe._ebca, _ebe._gec)
	var _fgcf float64
	if _gfc != 0 {
		_fgcf = 1.0 / _gfc
	}
	_eae := -_cfdb._gec
	_edbbb := &radialGradient{_ccc: _cfdb, _bde: _ebd, _cccf: _ebe, _eca: _gfc, _ffa: _fgcf, _gcfc: _eae}
	return _edbbb
}
func (_eaae *Context) DrawString(s string, face _ec.Face, x, y float64) {
	_eaae.DrawStringAnchored(s, face, x, y, 0, 0)
}
func (_bbc *Context) Rotate(angle float64) { _bbc._afg = _bbc._afg.Rotate(angle) }
func (_gba *Context) Scale(x, y float64)   { _gba._afg = _gba._afg.Scale(x, y) }
func (_aega *radialGradient) ColorAt(x, y int) _da.Color {
	if len(_aega._gefb) == 0 {
		return _da.Transparent
	}
	_abd, _fggd := float64(x)+0.5-_aega._ccc._ffb, float64(y)+0.5-_aega._ccc._ebca
	_cee := _acd(_abd, _fggd, _aega._ccc._gec, _aega._cccf._ffb, _aega._cccf._ebca, _aega._cccf._gec)
	_bfd := _acd(_abd, _fggd, -_aega._ccc._gec, _abd, _fggd, _aega._ccc._gec)
	if _aega._eca == 0 {
		if _cee == 0 {
			return _da.Transparent
		}
		_efg := 0.5 * _bfd / _cee
		if _efg*_aega._cccf._gec >= _aega._gcfc {
			return _bcf(_efg, _aega._gefb)
		}
		return _da.Transparent
	}
	_cba := _acd(_cee, _aega._eca, 0, _cee, -_bfd, 0)
	if _cba >= 0 {
		_bed := _b.Sqrt(_cba)
		_gcgf := (_cee + _bed) * _aega._ffa
		_ecac := (_cee - _bed) * _aega._ffa
		if _gcgf*_aega._cccf._gec >= _aega._gcfc {
			return _bcf(_gcgf, _aega._gefb)
		} else if _ecac*_aega._cccf._gec >= _aega._gcfc {
			return _bcf(_ecac, _aega._gefb)
		}
	}
	return _da.Transparent
}
func (_ceb *Context) ClearPath()                       { _ceb._db.Clear(); _ceb._eea.Clear(); _ceb._df = false }
func (_bfa *Context) SetLineJoin(lineJoin _g.LineJoin) { _bfa._cae = lineJoin }
func (_dfe *Context) Push()                            { _dafe := *_dfe; _dfe._ege = append(_dfe._ege, &_dafe) }
func (_bag *Context) DrawArc(x, y, r, angle1, angle2 float64) {
	_bag.DrawEllipticalArc(x, y, r, r, angle1, angle2)
}
func (_fadf *Context) stroke(_aac _ed.Painter) {
	_dbb := _fadf._db
	if len(_fadf._eff) > 0 {
		_dbb = _cce(_dbb, _fadf._eff, _fadf._faa)
	} else {
		_dbb = _edda(_cabd(_dbb))
	}
	_egec := _fadf._bbe
	_egec.UseNonZeroWinding = true
	_egec.Clear()
	_fc := (_fadf._afg.ScalingFactorX() + _fadf._afg.ScalingFactorY()) / 2
	_egec.AddStroke(_dbb, _dagb(_fadf._bgd*_fc), _fadf.capper(), _fadf.joiner())
	_egec.Rasterize(_aac)
}
func (_edd *Context) SetMatrix(m _ca.Matrix) { _edd._afg = m }
func (_ged *Context) DrawCircle(x, y, r float64) {
	_ged.NewSubPath()
	_ged.DrawEllipticalArc(x, y, r, r, 0, 2*_b.Pi)
	_ged.ClosePath()
}
func (_fca stops) Less(i, j int) bool { return _fca[i]._dede < _fca[j]._dede }
func (_caf *Context) DrawEllipse(x, y, rx, ry float64) {
	_caf.NewSubPath()
	_caf.DrawEllipticalArc(x, y, rx, ry, 0, 2*_b.Pi)
	_caf.ClosePath()
}
func (_fed *Context) SetPixel(x, y int) { _fed._agd.Set(x, y, _fed._geg) }
func (_ffff *Context) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64) {
	const _ced = 16
	for _bgfg := 0; _bgfg < _ced; _bgfg++ {
		_fac := float64(_bgfg+0) / _ced
		_abb := float64(_bgfg+1) / _ced
		_bdg := angle1 + (angle2-angle1)*_fac
		_gef := angle1 + (angle2-angle1)*_abb
		_gbb := x + rx*_b.Cos(_bdg)
		_acb := y + ry*_b.Sin(_bdg)
		_dac := x + rx*_b.Cos((_bdg+_gef)/2)
		_dce := y + ry*_b.Sin((_bdg+_gef)/2)
		_bdbb := x + rx*_b.Cos(_gef)
		_bbb := y + ry*_b.Sin(_gef)
		_dcdg := 2*_dac - _gbb/2 - _bdbb/2
		_gfea := 2*_dce - _acb/2 - _bbb/2
		if _bgfg == 0 {
			if _ffff._df {
				_ffff.LineTo(_gbb, _acb)
			} else {
				_ffff.MoveTo(_gbb, _acb)
			}
		}
		_ffff.QuadraticTo(_dcdg, _gfea, _bdbb, _bbb)
	}
}
func (_cgff *linearGradient) AddColorStop(offset float64, color _da.Color) {
	_cgff._cgf = append(_cgff._cgf, stop{_dede: offset, _gdbe: color})
	_c.Sort(_cgff._cgf)
}
func (_cfd *Context) SetRGB255(r, g, b int) { _cfd.SetRGBA255(r, g, b, 255) }
func _cdc(_dcaa _fb.Image) *_fb.RGBA {
	_ccdaf := _dcaa.Bounds()
	_cdbf := _fb.NewRGBA(_ccdaf)
	_a.Draw(_cdbf, _ccdaf, _dcaa, _ccdaf.Min, _a.Src)
	return _cdbf
}

const (
	_gecf repeatOp = iota
	_ddgc
	_gbag
	_bdea
)

func (_cfe *Context) MoveTo(x, y float64) {
	if _cfe._df {
		_cfe._eea.Add1(_eda(_cfe._dea))
	}
	x, y = _cfe.Transform(x, y)
	_bgg := _ca.NewPoint(x, y)
	_gfd := _eda(_bgg)
	_cfe._db.Start(_gfd)
	_cfe._eea.Start(_gfd)
	_cfe._dea = _bgg
	_cfe._efa = _bgg
	_cfe._df = true
}
func (_fdf *Context) Clip() { _fdf.ClipPreserve(); _fdf.ClearPath() }
func (_cab *Context) FillPreserve() {
	var _gbc _ed.Painter
	if _cab._aca == nil {
		if _fce, _fceb := _cab._aabb.(*solidPattern); _fceb {
			_fg := _ed.NewRGBAPainter(_cab._agd)
			_fg.SetColor(_fce._bbfc)
			_gbc = _fg
		}
	}
	if _gbc == nil {
		_gbc = _aagb(_cab._agd, _cab._aca, _cab._aabb)
	}
	_cab.fill(_gbc)
}
func (_dfd *Context) SetStrokeStyle(pattern _g.Pattern) { _dfd._gdc = pattern }
func (_bda *Context) SetRGBA255(r, g, b, a int) {
	_bda._geg = _da.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	_bda.setFillAndStrokeColor(_bda._geg)
}
func (_fdd *Context) DrawImage(im _fb.Image, x, y int) { _fdd.DrawImageAnchored(im, x, y, 0, 0) }
func (_caa *Context) SetColor(c _da.Color)             { _caa.setFillAndStrokeColor(c) }
func (_dcg *Context) TextState() *_g.TextState         { return &_dcg._beeb }
func (_aaf *Context) fill(_dagg _ed.Painter) {
	_aabe := _aaf._eea
	if _aaf._df {
		_aabe = make(_ed.Path, len(_aaf._eea))
		copy(_aabe, _aaf._eea)
		_aabe.Add1(_eda(_aaf._dea))
	}
	_abc := _aaf._bbe
	_abc.UseNonZeroWinding = _aaf._agg == _g.FillRuleWinding
	_abc.Clear()
	_abc.AddPath(_aabe)
	_abc.Rasterize(_dagg)
}
func _eda(_beee _ca.Point) _bg.Point26_6 { return _bg.Point26_6{X: _dagb(_beee.X), Y: _dagb(_beee.Y)} }

type radialGradient struct {
	_ccc, _bde, _cccf circle
	_eca, _ffa        float64
	_gcfc             float64
	_gefb             stops
}

func (_dgc *Context) Identity() { _dgc._afg = _ca.IdentityMatrix() }
func _bcf(_gbeg float64, _ffbe stops) _da.Color {
	if _gbeg <= 0.0 || len(_ffbe) == 1 {
		return _ffbe[0]._gdbe
	}
	_ece := _ffbe[len(_ffbe)-1]
	if _gbeg >= _ece._dede {
		return _ece._gdbe
	}
	for _deg, _bcac := range _ffbe[1:] {
		if _gbeg < _bcac._dede {
			_gbeg = (_gbeg - _ffbe[_deg]._dede) / (_bcac._dede - _ffbe[_deg]._dede)
			return _bea(_ffbe[_deg]._gdbe, _bcac._gdbe, _gbeg)
		}
	}
	return _ece._gdbe
}
func (_edc *Context) QuadraticTo(x1, y1, x2, y2 float64) {
	if !_edc._df {
		_edc.MoveTo(x1, y1)
	}
	x1, y1 = _edc.Transform(x1, y1)
	x2, y2 = _edc.Transform(x2, y2)
	_bdc := _ca.NewPoint(x1, y1)
	_agf := _ca.NewPoint(x2, y2)
	_ab := _eda(_bdc)
	_cbf := _eda(_agf)
	_edc._db.Add2(_ab, _cbf)
	_edc._eea.Add2(_ab, _cbf)
	_edc._efa = _agf
}
func (_cec *Context) joiner() _ed.Joiner {
	switch _cec._cae {
	case _g.LineJoinBevel:
		return _ed.BevelJoiner
	case _g.LineJoinRound:
		return _ed.RoundJoiner
	}
	return nil
}
func _acd(_fgc, _edbb, _geca, _babd, _ccce, _baga float64) float64 {
	return _fgc*_babd + _edbb*_ccce + _geca*_baga
}
func (_gfb *surfacePattern) ColorAt(x, y int) _da.Color {
	_gab := _gfb._fcaa.Bounds()
	switch _gfb._adee {
	case _ddgc:
		if y >= _gab.Dy() {
			return _da.Transparent
		}
	case _gbag:
		if x >= _gab.Dx() {
			return _da.Transparent
		}
	case _bdea:
		if x >= _gab.Dx() || y >= _gab.Dy() {
			return _da.Transparent
		}
	}
	x = x%_gab.Dx() + _gab.Min.X
	y = y%_gab.Dy() + _gab.Min.Y
	return _gfb._fcaa.At(x, y)
}
func (_eed *Context) Translate(x, y float64) { _eed._afg = _eed._afg.Translate(x, y) }
func _cce(_gbd _ed.Path, _bbaca []float64, _gfga float64) _ed.Path {
	return _edda(_afcd(_cabd(_gbd), _bbaca, _gfga))
}
func (_ce *Context) LineWidth() float64 { return _ce._bgd }
func _aagb(_bcfg *_fb.RGBA, _faac *_fb.Alpha, _fbe _g.Pattern) *patternPainter {
	return &patternPainter{_bcfg, _faac, _fbe}
}
func (_bef *Context) Clear() {
	_ceg := _fb.NewUniform(_bef._geg)
	_e.Draw(_bef._agd, _bef._agd.Bounds(), _ceg, _fb.Point{}, _e.Src)
}
func (_dcd *Context) StrokePattern() _g.Pattern { return _dcd._gdc }
func (_efaf *Context) Width() int               { return _efaf._eab }
func (_ace *Context) DrawRectangle(x, y, w, h float64) {
	_ace.NewSubPath()
	_ace.MoveTo(x, y)
	_ace.LineTo(x+w, y)
	_ace.LineTo(x+w, y+h)
	_ace.LineTo(x, y+h)
	_ace.ClosePath()
}
func (_aeb *Context) SetRGB(r, g, b float64) { _aeb.SetRGBA(r, g, b, 1) }
func _afcd(_afge [][]_ca.Point, _egg []float64, _bdbg float64) [][]_ca.Point {
	var _bga [][]_ca.Point
	if len(_egg) == 0 {
		return _afge
	}
	if len(_egg) == 1 {
		_egg = append(_egg, _egg[0])
	}
	for _, _beab := range _afge {
		if len(_beab) < 2 {
			continue
		}
		_gcd := _beab[0]
		_bbac := 1
		_cabdg := 0
		_cfdg := 0.0
		if _bdbg != 0 {
			var _ceba float64
			for _, _fdcc := range _egg {
				_ceba += _fdcc
			}
			_bdbg = _b.Mod(_bdbg, _ceba)
			if _bdbg < 0 {
				_bdbg += _ceba
			}
			for _bgbd, _acba := range _egg {
				_bdbg -= _acba
				if _bdbg < 0 {
					_cabdg = _bgbd
					_cfdg = _acba + _bdbg
					break
				}
			}
		}
		var _dcf []_ca.Point
		_dcf = append(_dcf, _gcd)
		for _bbac < len(_beab) {
			_baac := _egg[_cabdg]
			_cebag := _beab[_bbac]
			_bfcg := _gcd.Distance(_cebag)
			_deb := _baac - _cfdg
			if _bfcg > _deb {
				_cfef := _deb / _bfcg
				_ebdc := _gcd.Interpolate(_cebag, _cfef)
				_dcf = append(_dcf, _ebdc)
				if _cabdg%2 == 0 && len(_dcf) > 1 {
					_bga = append(_bga, _dcf)
				}
				_dcf = nil
				_dcf = append(_dcf, _ebdc)
				_cfdg = 0
				_gcd = _ebdc
				_cabdg = (_cabdg + 1) % len(_egg)
			} else {
				_dcf = append(_dcf, _cebag)
				_gcd = _cebag
				_cfdg += _bfcg
				_bbac++
			}
		}
		if _cabdg%2 == 0 && len(_dcf) > 1 {
			_bga = append(_bga, _dcf)
		}
	}
	return _bga
}
func (_egef *Context) SetFillRGBA(r, g, b, a float64) {
	_cb := _da.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_egef._geg = _cb
	_egef._aabb = _cecd(_cb)
}
func (_bdcd *Context) MeasureString(s string, face _ec.Face) (_efcc, _bdd float64) {
	_dedg := &_ec.Drawer{Face: face}
	_dcee := _dedg.MeasureString(s)
	return float64(_dcee >> 6), _bdcd._beeb.Tf.Size
}
func (_ecfe *Context) FillPattern() _g.Pattern { return _ecfe._aabb }
func (_afc *Context) SetStrokeRGBA(r, g, b, a float64) {
	_afd := _da.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_afc._gdc = _cecd(_afd)
}
func _adb(_bae, _eb, _gfa, _cff, _egd, _ee, _egcc, _bgf float64) []_ca.Point {
	_ag := (_b.Hypot(_gfa-_bae, _cff-_eb) + _b.Hypot(_egd-_gfa, _ee-_cff) + _b.Hypot(_egcc-_egd, _bgf-_ee))
	_fae := int(_ag + 0.5)
	if _fae < 4 {
		_fae = 4
	}
	_daga := float64(_fae) - 1
	_fba := make([]_ca.Point, _fae)
	for _ade := 0; _ade < _fae; _ade++ {
		_agb := float64(_ade) / _daga
		_bba, _eaa := _bd(_bae, _eb, _gfa, _cff, _egd, _ee, _egcc, _bgf, _agb)
		_fba[_ade] = _ca.NewPoint(_bba, _eaa)
	}
	return _fba
}
func _fbf(_gff float64) float64 { return _gff * _b.Pi / 180 }

type linearGradient struct {
	_bab, _fef, _dca, _cded float64
	_cgf                    stops
}

func _edda(_gfgb [][]_ca.Point) _ed.Path {
	var _dggb _ed.Path
	for _, _cbd := range _gfgb {
		var _gbcg _bg.Point26_6
		for _gcdd, _gbga := range _cbd {
			_fgba := _eda(_gbga)
			if _gcdd == 0 {
				_dggb.Start(_fgba)
			} else {
				_bbcg := _fgba.X - _gbcg.X
				_faf := _fgba.Y - _gbcg.Y
				if _bbcg < 0 {
					_bbcg = -_bbcg
				}
				if _faf < 0 {
					_faf = -_faf
				}
				if _bbcg+_faf > 8 {
					_dggb.Add1(_fgba)
				}
			}
			_gbcg = _fgba
		}
	}
	return _dggb
}

type surfacePattern struct {
	_fcaa _fb.Image
	_adee repeatOp
}
type stop struct {
	_dede float64
	_gdbe _da.Color
}

func (_gfaf *Context) ScaleAbout(sx, sy, x, y float64) {
	_gfaf.Translate(x, y)
	_gfaf.Scale(sx, sy)
	_gfaf.Translate(-x, -y)
}
func (_cedc *Context) DrawImageAnchored(im _fb.Image, x, y int, ax, ay float64) {
	_beb := im.Bounds().Size()
	x -= int(ax * float64(_beb.X))
	y -= int(ay * float64(_beb.Y))
	_egcf := _e.BiLinear
	_dbcf := _cedc._afg.Clone().Translate(float64(x), float64(y))
	_ccag := _fd.Aff3{_dbcf[0], _dbcf[3], _dbcf[6], _dbcf[1], _dbcf[4], _dbcf[7]}
	if _cedc._aca == nil {
		_egcf.Transform(_cedc._agd, _ccag, im, im.Bounds(), _e.Over, nil)
	} else {
		_egcf.Transform(_cedc._agd, _ccag, im, im.Bounds(), _e.Over, &_e.Options{DstMask: _cedc._aca, DstMaskP: _fb.Point{}})
	}
}

var (
	_bdbf = _cecd(_da.White)
	_gcb  = _cecd(_da.Black)
)

func (_gae *Context) InvertMask() {
	if _gae._aca == nil {
		_gae._aca = _fb.NewAlpha(_gae._agd.Bounds())
	} else {
		for _cfff, _agc := range _gae._aca.Pix {
			_gae._aca.Pix[_cfff] = 255 - _agc
		}
	}
}

type repeatOp int
type stops []stop
