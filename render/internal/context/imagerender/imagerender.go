package imagerender

import (
	_b "errors"
	_df "fmt"
	_c "image"
	_a "image/color"
	_ae "image/draw"
	_e "math"
	_ebe "sort"
	_eb "strings"

	_de "bitbucket.org/shenghui0779/gopdf/common"
	_f "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_g "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_dc "github.com/unidoc/freetype/raster"
	_fc "golang.org/x/image/draw"
	_def "golang.org/x/image/font"
	_cd "golang.org/x/image/math/f64"
	_be "golang.org/x/image/math/fixed"
)

func (_bed *Context) SetFillRule(fillRule _g.FillRule) { _bed._agb = fillRule }
func (_fbcb *Context) MeasureString(s string, face _def.Face) (_gec, _baa float64) {
	_dedc := &_def.Drawer{Face: face}
	_gfca := _dedc.MeasureString(s)
	return float64(_gfca >> 6), _fbcb._fgf.Tf.Size
}
func (_dbf *Context) ClosePath() {
	if _dbf._gdd {
		_dfc := _dfce(_dbf._fbfg)
		_dbf._ge.Add1(_dfc)
		_dbf._af.Add1(_dfc)
		_dbf._bbc = _dbf._fbfg
	}
}
func (_ccg *Context) LineWidth() float64               { return _ccg._fag }
func (_bfcb *Context) DrawImage(im _c.Image, x, y int) { _bfcb.DrawImageAnchored(im, x, y, 0, 0) }
func (_bbd *Context) ClearPath() {
	_bbd._ge.Clear()
	_bbd._af.Clear()
	_bbd._gdd = false
}
func (_dgd *solidPattern) ColorAt(x, y int) _a.Color { return _dgd._agde }
func (_fccc *Context) RotateAbout(angle, x, y float64) {
	_fccc.Translate(x, y)
	_fccc.Rotate(angle)
	_fccc.Translate(-x, -y)
}

var (
	_ecg = _eccf(_a.White)
	_gff = _eccf(_a.Black)
)

func (_cac *Context) FillPreserve() {
	var _dddf _dc.Painter
	if _cac._bba == nil {
		if _fgd, _fff := _cac._bfeg.(*solidPattern); _fff {
			_ebc := _dc.NewRGBAPainter(_cac._feg)
			_ebc.SetColor(_fgd._agde)
			_dddf = _ebc
		}
	}
	if _dddf == nil {
		_dddf = _efeb(_cac._feg, _cac._bba, _cac._bfeg)
	}
	_cac.fill(_dddf)
}
func (_cda *Context) stroke(_aeaa _dc.Painter) {
	_cca := _cda._ge
	if len(_cda._ea) > 0 {
		_cca = _abaa(_cca, _cda._ea, _cda._bc)
	} else {
		_cca = _bff(_ddbg(_cca))
	}
	_cee := _cda._ef
	_cee.UseNonZeroWinding = true
	_cee.Clear()
	_gba := (_cda._gdg.ScalingFactorX() + _cda._gdg.ScalingFactorY()) / 2
	_cee.AddStroke(_cca, _aecb(_cda._fag*_gba), _cda.capper(), _cda.joiner())
	_cee.Rasterize(_aeaa)
}
func (_ddac *Context) NewSubPath() {
	if _ddac._gdd {
		_ddac._af.Add1(_dfce(_ddac._fbfg))
	}
	_ddac._gdd = false
}
func (_bcf *Context) QuadraticTo(x1, y1, x2, y2 float64) {
	if !_bcf._gdd {
		_bcf.MoveTo(x1, y1)
	}
	x1, y1 = _bcf.Transform(x1, y1)
	x2, y2 = _bcf.Transform(x2, y2)
	_efb := _f.NewPoint(x1, y1)
	_bfda := _f.NewPoint(x2, y2)
	_faf := _dfce(_efb)
	_fba := _dfce(_bfda)
	_bcf._ge.Add2(_faf, _fba)
	_bcf._af.Add2(_faf, _fba)
	_bcf._bbc = _bfda
}

type surfacePattern struct {
	_gcd  _c.Image
	_bedb repeatOp
}

func (_cbe *radialGradient) AddColorStop(offset float64, color _a.Color) {
	_cbe._ffa = append(_cbe._ffa, stop{_efe: offset, _gebd: color})
	_ebe.Sort(_cbe._ffa)
}
func (_bea *Context) SetLineWidth(lineWidth float64) { _bea._fag = lineWidth }
func (_ecgf *Context) Matrix() _f.Matrix             { return _ecgf._gdg }
func (_aad *Context) Shear(x, y float64)             { _aad._gdg.Shear(x, y) }
func _daea(_daa, _eagd, _ddfb, _ggcc, _bab, _fbde float64) _g.Gradient {
	_ece := circle{_daa, _eagd, _ddfb}
	_fda := circle{_ggcc, _bab, _fbde}
	_ffffc := circle{_ggcc - _daa, _bab - _eagd, _fbde - _ddfb}
	_aeea := _egbb(_ffffc._aebf, _ffffc._fbdc, -_ffffc._acb, _ffffc._aebf, _ffffc._fbdc, _ffffc._acb)
	var _fccaf float64
	if _aeea != 0 {
		_fccaf = 1.0 / _aeea
	}
	_agc := -_ece._acb
	_beb := &radialGradient{_abd: _ece, _dca: _fda, _cade: _ffffc, _gada: _aeea, _ccae: _fccaf, _ecbg: _agc}
	return _beb
}
func (_abbb *Context) ResetClip() { _abbb._bba = nil }
func (_ddf *Context) DrawLine(x1, y1, x2, y2 float64) {
	_ddf.MoveTo(x1, y1)
	_ddf.LineTo(x2, y2)
}
func _bff(_gaf [][]_f.Point) _dc.Path {
	var _bcdc _dc.Path
	for _, _bdaf := range _gaf {
		var _dgc _be.Point26_6
		for _ebca, _fadc := range _bdaf {
			_babg := _dfce(_fadc)
			if _ebca == 0 {
				_bcdc.Start(_babg)
			} else {
				_dfba := _babg.X - _dgc.X
				_eebd := _babg.Y - _dgc.Y
				if _dfba < 0 {
					_dfba = -_dfba
				}
				if _eebd < 0 {
					_eebd = -_eebd
				}
				if _dfba+_eebd > 8 {
					_bcdc.Add1(_babg)
				}
			}
			_dgc = _babg
		}
	}
	return _bcdc
}
func (_fcc *Context) MoveTo(x, y float64) {
	if _fcc._gdd {
		_fcc._af.Add1(_dfce(_fcc._fbfg))
	}
	x, y = _fcc.Transform(x, y)
	_dee := _f.NewPoint(x, y)
	_fcdc := _dfce(_dee)
	_fcc._ge.Start(_fcdc)
	_fcc._af.Start(_fcdc)
	_fcc._fbfg = _dee
	_fcc._bbc = _dee
	_fcc._gdd = true
}
func (_gfef *Context) SetMask(mask *_c.Alpha) error {
	if mask.Bounds().Size() != _gfef._feg.Bounds().Size() {
		return _b.New("\u006d\u0061\u0073\u006b\u0020\u0073i\u007a\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068 \u0063\u006f\u006e\u0074\u0065\u0078\u0074 \u0073\u0069\u007a\u0065")
	}
	_gfef._bba = mask
	return nil
}
func (_cab *Context) SetRGBA255(r, g, b, a int) {
	_cab._egg = _a.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	_cab.setFillAndStrokeColor(_cab._egg)
}
func NewContextForImage(im _c.Image) *Context { return NewContextForRGBA(_babgf(im)) }
func (_aa *Context) SetStrokeRGBA(r, g, b, a float64) {
	_ddg := _a.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_aa._fg = _eccf(_ddg)
}

type linearGradient struct {
	_gffe, _fgde, _agfc, _cagb float64
	_feda                      stops
}

func (_cacc *Context) Identity() { _cacc._gdg = _f.IdentityMatrix() }
func (_add *Context) Clear() {
	_fdb := _c.NewUniform(_add._egg)
	_fc.Draw(_add._feg, _add._feg.Bounds(), _fdb, _c.Point{}, _fc.Src)
}
func (_fegd *Context) Clip()               { _fegd.ClipPreserve(); _fegd.ClearPath() }
func (_baee *Context) SetColor(c _a.Color) { _baee.setFillAndStrokeColor(c) }
func _dfce(_gabf _f.Point) _be.Point26_6   { return _be.Point26_6{X: _aecb(_gabf.X), Y: _aecb(_gabf.Y)} }
func (_fad *Context) Image() _c.Image      { return _fad._feg }
func _egbe(_agcc float64, _efae stops) _a.Color {
	if _agcc <= 0.0 || len(_efae) == 1 {
		return _efae[0]._gebd
	}
	_fadbb := _efae[len(_efae)-1]
	if _agcc >= _fadbb._efe {
		return _fadbb._gebd
	}
	for _bbca, _agg := range _efae[1:] {
		if _agcc < _agg._efe {
			_agcc = (_agcc - _efae[_bbca]._efe) / (_agg._efe - _efae[_bbca]._efe)
			return _gag(_efae[_bbca]._gebd, _agg._gebd, _agcc)
		}
	}
	return _fadbb._gebd
}
func (_ffff *Context) Rotate(angle float64) { _ffff._gdg = _ffff._gdg.Rotate(angle) }
func (_ffb *Context) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64) {
	const _dcd = 16
	for _fab := 0; _fab < _dcd; _fab++ {
		_fbd := float64(_fab+0) / _dcd
		_cde := float64(_fab+1) / _dcd
		_eeb := angle1 + (angle2-angle1)*_fbd
		_gfa := angle1 + (angle2-angle1)*_cde
		_eccg := x + rx*_e.Cos(_eeb)
		_agfb := y + ry*_e.Sin(_eeb)
		_gad := x + rx*_e.Cos((_eeb+_gfa)/2)
		_ffbc := y + ry*_e.Sin((_eeb+_gfa)/2)
		_ded := x + rx*_e.Cos(_gfa)
		_cf := y + ry*_e.Sin(_gfa)
		_beg := 2*_gad - _eccg/2 - _ded/2
		_cgg := 2*_ffbc - _agfb/2 - _cf/2
		if _fab == 0 {
			if _ffb._gdd {
				_ffb.LineTo(_eccg, _agfb)
			} else {
				_ffb.MoveTo(_eccg, _agfb)
			}
		}
		_ffb.QuadraticTo(_beg, _cgg, _ded, _cf)
	}
}
func (_aag *Context) DrawRectangle(x, y, w, h float64) {
	_aag.NewSubPath()
	_aag.MoveTo(x, y)
	_aag.LineTo(x+w, y)
	_aag.LineTo(x+w, y+h)
	_aag.LineTo(x, y+h)
	_aag.ClosePath()
}

type solidPattern struct{ _agde _a.Color }

func (_gdb *Context) setFillAndStrokeColor(_aea _a.Color) {
	_gdb._egg = _aea
	_gdb._bfeg = _eccf(_aea)
	_gdb._fg = _eccf(_aea)
}
func _gg(_ddd, _ccb, _aed, _bb, _gc, _dde, _gda, _caf, _abc float64) (_da, _fdc float64) {
	_gb := 1 - _abc
	_gdc := _gb * _gb * _gb
	_cce := 3 * _gb * _gb * _abc
	_dbe := 3 * _gb * _abc * _abc
	_ag := _abc * _abc * _abc
	_da = _gdc*_ddd + _cce*_aed + _dbe*_gc + _ag*_gda
	_fdc = _gdc*_ccb + _cce*_bb + _dbe*_dde + _ag*_caf
	return
}
func (_agd *Context) SetFillRGBA(r, g, b, a float64) {
	_gdad := _a.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_agd._egg = _gdad
	_agd._bfeg = _eccf(_gdad)
}
func (_ddc *Context) fill(_egd _dc.Painter) {
	_egde := _ddc._af
	if _ddc._gdd {
		_egde = make(_dc.Path, len(_ddc._af))
		copy(_egde, _ddc._af)
		_egde.Add1(_dfce(_ddc._fbfg))
	}
	_bdad := _ddc._ef
	_bdad.UseNonZeroWinding = _ddc._agb == _g.FillRuleWinding
	_bdad.Clear()
	_bdad.AddPath(_egde)
	_bdad.Rasterize(_egd)
}

type Context struct {
	_gcb  int
	_cdf  int
	_ef   *_dc.Rasterizer
	_feg  *_c.RGBA
	_bba  *_c.Alpha
	_egg  _a.Color
	_bfeg _g.Pattern
	_fg   _g.Pattern
	_ge   _dc.Path
	_af   _dc.Path
	_fbfg _f.Point
	_bbc  _f.Point
	_gdd  bool
	_ea   []float64
	_bc   float64
	_fag  float64
	_bae  _g.LineCap
	_bfd  _g.LineJoin
	_agb  _g.FillRule
	_gdg  _f.Matrix
	_fgf  _g.TextState
	_fbc  []*Context
}

func _gag(_eed, _daf _a.Color, _gddf float64) _a.Color {
	_gggd, _dfb, _agfbf, _bcgg := _eed.RGBA()
	_gaag, _gfg, _fagd, _dbc := _daf.RGBA()
	return _a.RGBA{_fcgg(_gggd, _gaag, _gddf), _fcgg(_dfb, _gfg, _gddf), _fcgg(_agfbf, _fagd, _gddf), _fcgg(_bcgg, _dbc, _gddf)}
}
func (_adc *Context) InvertMask() {
	if _adc._bba == nil {
		_adc._bba = _c.NewAlpha(_adc._feg.Bounds())
	} else {
		for _cag, _cdae := range _adc._bba.Pix {
			_adc._bba.Pix[_cag] = 255 - _cdae
		}
	}
}
func (_agfba *linearGradient) ColorAt(x, y int) _a.Color {
	if len(_agfba._feda) == 0 {
		return _a.Transparent
	}
	_abcb, _dfg := float64(x), float64(y)
	_cbf, _dedb, _fafd, _cfda := _agfba._gffe, _agfba._fgde, _agfba._agfc, _agfba._cagb
	_bada, _bdef := _fafd-_cbf, _cfda-_dedb
	if _bdef == 0 && _bada != 0 {
		return _egbe((_abcb-_cbf)/_bada, _agfba._feda)
	}
	if _bada == 0 && _bdef != 0 {
		return _egbe((_dfg-_dedb)/_bdef, _agfba._feda)
	}
	_deb := _bada*(_abcb-_cbf) + _bdef*(_dfg-_dedb)
	if _deb < 0 {
		return _agfba._feda[0]._gebd
	}
	_aeb := _e.Hypot(_bada, _bdef)
	_gaa := ((_abcb-_cbf)*-_bdef + (_dfg-_dedb)*_bada) / (_aeb * _aeb)
	_dcb, _fdcb := _cbf+_gaa*-_bdef, _dedb+_gaa*_bada
	_dcbb := _e.Hypot(_abcb-_dcb, _dfg-_fdcb) / _aeb
	return _egbe(_dcbb, _agfba._feda)
}
func (_bcea *Context) DrawEllipse(x, y, rx, ry float64) {
	_bcea.NewSubPath()
	_bcea.DrawEllipticalArc(x, y, rx, ry, 0, 2*_e.Pi)
	_bcea.ClosePath()
}
func (_dabc *Context) AsMask() *_c.Alpha {
	_cbb := _c.NewAlpha(_dabc._feg.Bounds())
	_fc.Draw(_cbb, _dabc._feg.Bounds(), _dabc._feg, _c.Point{}, _fc.Src)
	return _cbb
}

type repeatOp int

func (_gdf *Context) Push() { _edg := *_gdf; _gdf._fbc = append(_gdf._fbc, &_edg) }
func NewContextForRGBA(im *_c.RGBA) *Context {
	_bfc := im.Bounds().Size().X
	_fdd := im.Bounds().Size().Y
	return &Context{_gcb: _bfc, _cdf: _fdd, _ef: _dc.NewRasterizer(_bfc, _fdd), _feg: im, _egg: _a.Transparent, _bfeg: _ecg, _fg: _gff, _fag: 1, _agb: _g.FillRuleWinding, _gdg: _f.IdentityMatrix(), _fgf: _g.NewTextState()}
}

const (
	_aff repeatOp = iota
	_aaa
	_gfbe
	_ggggg
)

type radialGradient struct {
	_abd, _dca, _cade circle
	_gada, _ccae      float64
	_ecbg             float64
	_ffa              stops
}

func NewContext(width, height int) *Context {
	return NewContextForRGBA(_c.NewRGBA(_c.Rect(0, 0, width, height)))
}
func (_fadb *Context) SetMatrix(m _f.Matrix) { _fadb._gdg = m }
func (_aef *Context) LineTo(x, y float64) {
	if !_aef._gdd {
		_aef.MoveTo(x, y)
	} else {
		x, y = _aef.Transform(x, y)
		_ggc := _f.NewPoint(x, y)
		_eag := _dfce(_ggc)
		_aef._ge.Add1(_eag)
		_aef._af.Add1(_eag)
		_aef._bbc = _ggc
	}
}

type circle struct{ _aebf, _fbdc, _acb float64 }

func (_gbdc *Context) TextState() *_g.TextState { return &_gbdc._fgf }
func (_ddfe *Context) DrawImageAnchored(im _c.Image, x, y int, ax, ay float64) {
	_ffdb := im.Bounds().Size()
	x -= int(ax * float64(_ffdb.X))
	y -= int(ay * float64(_ffdb.Y))
	_ecge := _fc.BiLinear
	_dbfc := _ddfe._gdg.Clone().Translate(float64(x), float64(y))
	_aabb := _cd.Aff3{_dbfc[0], _dbfc[3], _dbfc[6], _dbfc[1], _dbfc[4], _dbfc[7]}
	if _ddfe._bba == nil {
		_ecge.Transform(_ddfe._feg, _aabb, im, im.Bounds(), _fc.Over, nil)
	} else {
		_ecge.Transform(_ddfe._feg, _aabb, im, im.Bounds(), _fc.Over, &_fc.Options{DstMask: _ddfe._bba, DstMaskP: _c.Point{}})
	}
}
func (_egc *Context) Height() int { return _egc._cdf }
func (_bcg *Context) SetRGBA(r, g, b, a float64) {
	_bcg._egg = _a.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_bcg.setFillAndStrokeColor(_bcg._egg)
}
func (_gbb *Context) StrokePattern() _g.Pattern { return _gbb._fg }
func (_eac stops) Less(i, j int) bool           { return _eac[i]._efe < _eac[j]._efe }
func (_ddgd *surfacePattern) ColorAt(x, y int) _a.Color {
	_deec := _ddgd._gcd.Bounds()
	switch _ddgd._bedb {
	case _aaa:
		if y >= _deec.Dy() {
			return _a.Transparent
		}
	case _gfbe:
		if x >= _deec.Dx() {
			return _a.Transparent
		}
	case _ggggg:
		if x >= _deec.Dx() || y >= _deec.Dy() {
			return _a.Transparent
		}
	}
	x = x%_deec.Dx() + _deec.Min.X
	y = y%_deec.Dy() + _deec.Min.Y
	return _ddgd._gcd.At(x, y)
}
func _abaa(_cagf _dc.Path, _ccaef []float64, _cgf float64) _dc.Path {
	return _bff(_gadg(_ddbg(_cagf), _ccaef, _cgf))
}
func _fcgg(_fdcbd, _gbe uint32, _bec float64) uint8 {
	return uint8(int32(float64(_fdcbd)*(1.0-_bec)+float64(_gbe)*_bec) >> 8)
}
func _gcf(_gfb, _bdf, _bde, _age, _ecc, _fbf, _ffd, _cba float64) []_f.Point {
	_gbc := (_e.Hypot(_bde-_gfb, _age-_bdf) + _e.Hypot(_ecc-_bde, _fbf-_age) + _e.Hypot(_ffd-_ecc, _cba-_fbf))
	_abb := int(_gbc + 0.5)
	if _abb < 4 {
		_abb = 4
	}
	_dec := float64(_abb) - 1
	_ba := make([]_f.Point, _abb)
	for _dab := 0; _dab < _abb; _dab++ {
		_ad := float64(_dab) / _dec
		_ac, _aba := _gg(_gfb, _bdf, _bde, _age, _ecc, _fbf, _ffd, _cba, _ad)
		_ba[_dab] = _f.NewPoint(_ac, _aba)
	}
	return _ba
}
func _efeb(_aefe *_c.RGBA, _fccag *_c.Alpha, _gea _g.Pattern) *patternPainter {
	return &patternPainter{_aefe, _fccag, _gea}
}
func (_bcec stops) Swap(i, j int) { _bcec[i], _bcec[j] = _bcec[j], _bcec[i] }
func (_cgbb *Context) Fill()      { _cgbb.FillPreserve(); _cgbb.ClearPath() }
func (_gdcc *Context) DrawPoint(x, y, r float64) {
	_gdcc.Push()
	_bcd, _geb := _gdcc.Transform(x, y)
	_gdcc.Identity()
	_gdcc.DrawCircle(_bcd, _geb, r)
	_gdcc.Pop()
}
func _bbdd(_cfa, _aceg, _cced, _aga float64) _g.Gradient {
	_bbfb := &linearGradient{_gffe: _cfa, _fgde: _aceg, _agfc: _cced, _cagb: _aga}
	return _bbfb
}
func (_ecca *Context) ScaleAbout(sx, sy, x, y float64) {
	_ecca.Translate(x, y)
	_ecca.Scale(sx, sy)
	_ecca.Translate(-x, -y)
}
func (_dfec *radialGradient) ColorAt(x, y int) _a.Color {
	if len(_dfec._ffa) == 0 {
		return _a.Transparent
	}
	_aeaac, _ddee := float64(x)+0.5-_dfec._abd._aebf, float64(y)+0.5-_dfec._abd._fbdc
	_eee := _egbb(_aeaac, _ddee, _dfec._abd._acb, _dfec._cade._aebf, _dfec._cade._fbdc, _dfec._cade._acb)
	_deg := _egbb(_aeaac, _ddee, -_dfec._abd._acb, _aeaac, _ddee, _dfec._abd._acb)
	if _dfec._gada == 0 {
		if _eee == 0 {
			return _a.Transparent
		}
		_baed := 0.5 * _deg / _eee
		if _baed*_dfec._cade._acb >= _dfec._ecbg {
			return _egbe(_baed, _dfec._ffa)
		}
		return _a.Transparent
	}
	_dga := _egbb(_eee, _dfec._gada, 0, _eee, -_deg, 0)
	if _dga >= 0 {
		_ebgg := _e.Sqrt(_dga)
		_dfa := (_eee + _ebgg) * _dfec._ccae
		_ecfd := (_eee - _ebgg) * _dfec._ccae
		if _dfa*_dfec._cade._acb >= _dfec._ecbg {
			return _egbe(_dfa, _dfec._ffa)
		} else if _ecfd*_dfec._cade._acb >= _dfec._ecbg {
			return _egbe(_ecfd, _dfec._ffa)
		}
	}
	return _a.Transparent
}
func (_fcg *Context) SetLineJoin(lineJoin _g.LineJoin) { _fcg._bfd = lineJoin }
func (_agf *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	if !_agf._gdd {
		_agf.MoveTo(x1, y1)
	}
	_ggcf, _cef := _agf._bbc.X, _agf._bbc.Y
	x1, y1 = _agf.Transform(x1, y1)
	x2, y2 = _agf.Transform(x2, y2)
	x3, y3 = _agf.Transform(x3, y3)
	_fee := _gcf(_ggcf, _cef, x1, y1, x2, y2, x3, y3)
	_bef := _dfce(_agf._bbc)
	for _, _fadg := range _fee[1:] {
		_ecb := _dfce(_fadg)
		if _ecb == _bef {
			continue
		}
		_bef = _ecb
		_agf._ge.Add1(_ecb)
		_agf._af.Add1(_ecb)
		_agf._bbc = _fadg
	}
}
func (_aab *Context) Stroke() { _aab.StrokePreserve(); _aab.ClearPath() }
func (_cfd *Context) DrawStringAnchored(s string, face _def.Face, x, y, ax, ay float64) {
	_dadb, _eagea := _cfd.MeasureString(s, face)
	_cfd.drawString(s, face, x-ax*_dadb, y+ay*_eagea)
}
func (_bfa *Context) DrawString(s string, face _def.Face, x, y float64) {
	_bfa.DrawStringAnchored(s, face, x, y, 0, 0)
}
func (_gbcb *Context) joiner() _dc.Joiner {
	switch _gbcb._bfd {
	case _g.LineJoinBevel:
		return _dc.BevelJoiner
	case _g.LineJoinRound:
		return _dc.RoundJoiner
	}
	return nil
}
func (_eagb *Context) ShearAbout(sx, sy, x, y float64) {
	_eagb.Translate(x, y)
	_eagb.Shear(sx, sy)
	_eagb.Translate(-x, -y)
}
func (_bced *Context) Translate(x, y float64) { _bced._gdg = _bced._gdg.Translate(x, y) }
func _aecb(_eeec float64) _be.Int26_6         { return _be.Int26_6(_eeec * 64) }
func (_gbd *Context) DrawArc(x, y, r, angle1, angle2 float64) {
	_gbd.DrawEllipticalArc(x, y, r, r, angle1, angle2)
}
func (_gee *linearGradient) AddColorStop(offset float64, color _a.Color) {
	_gee._feda = append(_gee._feda, stop{_efe: offset, _gebd: color})
	_ebe.Sort(_gee._feda)
}
func (_fdcd *Context) DrawRoundedRectangle(x, y, w, h, r float64) {
	_cec, _cbbc, _abbc, _adb := x, x+r, x+w-r, x+w
	_ace, _bca, _efa, _ga := y, y+r, y+h-r, y+h
	_fdcd.NewSubPath()
	_fdcd.MoveTo(_cbbc, _ace)
	_fdcd.LineTo(_abbc, _ace)
	_fdcd.DrawArc(_abbc, _bca, r, _eef(270), _eef(360))
	_fdcd.LineTo(_adb, _efa)
	_fdcd.DrawArc(_abbc, _efa, r, _eef(0), _eef(90))
	_fdcd.LineTo(_cbbc, _ga)
	_fdcd.DrawArc(_cbbc, _efa, r, _eef(90), _eef(180))
	_fdcd.LineTo(_cec, _bca)
	_fdcd.DrawArc(_cbbc, _bca, r, _eef(180), _eef(270))
	_fdcd.ClosePath()
}

type stop struct {
	_efe  float64
	_gebd _a.Color
}

func (_cgd *Context) Scale(x, y float64) { _cgd._gdg = _cgd._gdg.Scale(x, y) }
func (_cdfc *Context) drawString(_bbe string, _aeaf _def.Face, _ebd, _gdeg float64) {
	_gcg := &_def.Drawer{Src: _c.NewUniform(_cdfc._egg), Face: _aeaf, Dot: _dfce(_f.NewPoint(_ebd, _gdeg))}
	_gbf := rune(-1)
	for _, _ada := range _bbe {
		if _gbf >= 0 {
			_gcg.Dot.X += _gcg.Face.Kern(_gbf, _ada)
		}
		_bbag, _fcca, _cggg, _gfeg, _fea := _gcg.Face.Glyph(_gcg.Dot, _ada)
		if !_fea {
			continue
		}
		_fadd := _bbag.Sub(_bbag.Min)
		_ecbc := _c.NewRGBA(_fadd)
		_fc.DrawMask(_ecbc, _fadd, _gcg.Src, _c.Point{}, _fcca, _cggg, _fc.Over)
		var _ebf *_fc.Options
		if _cdfc._bba != nil {
			_ebf = &_fc.Options{DstMask: _cdfc._bba, DstMaskP: _c.Point{}}
		}
		_fedd := _cdfc._gdg.Clone().Translate(float64(_bbag.Min.X), float64(_bbag.Min.Y))
		_cae := _cd.Aff3{_fedd[0], _fedd[3], _fedd[6], _fedd[1], _fedd[4], _fedd[7]}
		_fc.BiLinear.Transform(_cdfc._feg, _cae, _ecbc, _fadd, _fc.Over, _ebf)
		_gcg.Dot.X += _gfeg
		_gbf = _ada
	}
}
func (_gfff *Context) Transform(x, y float64) (_ced, _bbf float64) { return _gfff._gdg.Transform(x, y) }
func _ggb(_defg _c.Image, _fdag repeatOp) _g.Pattern {
	return &surfacePattern{_gcd: _defg, _bedb: _fdag}
}
func _ddbg(_gcbc _dc.Path) [][]_f.Point {
	var _gfbc [][]_f.Point
	var _dcda []_f.Point
	var _dfcb, _afa float64
	for _dada := 0; _dada < len(_gcbc); {
		switch _gcbc[_dada] {
		case 0:
			if len(_dcda) > 0 {
				_gfbc = append(_gfbc, _dcda)
				_dcda = nil
			}
			_gga := _edd(_gcbc[_dada+1])
			_fgb := _edd(_gcbc[_dada+2])
			_dcda = append(_dcda, _f.NewPoint(_gga, _fgb))
			_dfcb, _afa = _gga, _fgb
			_dada += 4
		case 1:
			_cga := _edd(_gcbc[_dada+1])
			_debb := _edd(_gcbc[_dada+2])
			_dcda = append(_dcda, _f.NewPoint(_cga, _debb))
			_dfcb, _afa = _cga, _debb
			_dada += 4
		case 2:
			_ffg := _edd(_gcbc[_dada+1])
			_aec := _edd(_gcbc[_dada+2])
			_aedb := _edd(_gcbc[_dada+3])
			_bbde := _edd(_gcbc[_dada+4])
			_fdcc := _dfe(_dfcb, _afa, _ffg, _aec, _aedb, _bbde)
			_dcda = append(_dcda, _fdcc...)
			_dfcb, _afa = _aedb, _bbde
			_dada += 6
		case 3:
			_adg := _edd(_gcbc[_dada+1])
			_bdfc := _edd(_gcbc[_dada+2])
			_fgdb := _edd(_gcbc[_dada+3])
			_fga := _edd(_gcbc[_dada+4])
			_ead := _edd(_gcbc[_dada+5])
			_cge := _edd(_gcbc[_dada+6])
			_bag := _gcf(_dfcb, _afa, _adg, _bdfc, _fgdb, _fga, _ead, _cge)
			_dcda = append(_dcda, _bag...)
			_dfcb, _afa = _ead, _cge
			_dada += 8
		default:
			_de.Log.Debug("\u0057\u0041\u0052\u004e: \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0074\u0068\u003a\u0020%\u0076", _gcbc)
			return _gfbc
		}
	}
	if len(_dcda) > 0 {
		_gfbc = append(_gfbc, _dcda)
	}
	return _gfbc
}
func _egbb(_agfg, _afd, _daeg, _ecbb, _eace, _daeba float64) float64 {
	return _agfg*_ecbb + _afd*_eace + _daeg*_daeba
}
func (_ed *Context) DrawCircle(x, y, r float64) {
	_ed.NewSubPath()
	_ed.DrawEllipticalArc(x, y, r, r, 0, 2*_e.Pi)
	_ed.ClosePath()
}
func _gadg(_gca [][]_f.Point, _ecbe []float64, _cbc float64) [][]_f.Point {
	var _eabg [][]_f.Point
	if len(_ecbe) == 0 {
		return _gca
	}
	if len(_ecbe) == 1 {
		_ecbe = append(_ecbe, _ecbe[0])
	}
	for _, _bcdd := range _gca {
		if len(_bcdd) < 2 {
			continue
		}
		_cgde := _bcdd[0]
		_dddb := 1
		_ega := 0
		_cbcb := 0.0
		if _cbc != 0 {
			var _eebg float64
			for _, _aece := range _ecbe {
				_eebg += _aece
			}
			_cbc = _e.Mod(_cbc, _eebg)
			if _cbc < 0 {
				_cbc += _eebg
			}
			for _eec, _dedd := range _ecbe {
				_cbc -= _dedd
				if _cbc < 0 {
					_ega = _eec
					_cbcb = _dedd + _cbc
					break
				}
			}
		}
		var _ebce []_f.Point
		_ebce = append(_ebce, _cgde)
		for _dddb < len(_bcdd) {
			_cabd := _ecbe[_ega]
			_fddb := _bcdd[_dddb]
			_addd := _cgde.Distance(_fddb)
			_bgc := _cabd - _cbcb
			if _addd > _bgc {
				_ccgf := _bgc / _addd
				_ebcd := _cgde.Interpolate(_fddb, _ccgf)
				_ebce = append(_ebce, _ebcd)
				if _ega%2 == 0 && len(_ebce) > 1 {
					_eabg = append(_eabg, _ebce)
				}
				_ebce = nil
				_ebce = append(_ebce, _ebcd)
				_cbcb = 0
				_cgde = _ebcd
				_ega = (_ega + 1) % len(_ecbe)
			} else {
				_ebce = append(_ebce, _fddb)
				_cgde = _fddb
				_cbcb += _addd
				_dddb++
			}
		}
		if _ega%2 == 0 && len(_ebce) > 1 {
			_eabg = append(_eabg, _ebce)
		}
	}
	return _eabg
}
func (_ebg *Context) Width() int { return _ebg._gcb }
func (_bee *Context) ClipPreserve() {
	_afc := _c.NewAlpha(_c.Rect(0, 0, _bee._gcb, _bee._cdf))
	_bad := _dc.NewAlphaOverPainter(_afc)
	_bee.fill(_bad)
	if _bee._bba == nil {
		_bee._bba = _afc
	} else {
		_fcde := _c.NewAlpha(_c.Rect(0, 0, _bee._gcb, _bee._cdf))
		_fc.DrawMask(_fcde, _fcde.Bounds(), _afc, _c.Point{}, _bee._bba, _c.Point{}, _fc.Over)
		_bee._bba = _fcde
	}
}
func (_fcce stops) Len() int { return len(_fcce) }
func _babgf(_dcf _c.Image) *_c.RGBA {
	_fdg := _dcf.Bounds()
	_dea := _c.NewRGBA(_fdg)
	_ae.Draw(_dea, _fdg, _dcf, _fdg.Min, _ae.Src)
	return _dea
}
func (_dae *Context) drawRegularPolygon(_ddab int, _fcb, _aefc, _gggg, _dfcd float64) {
	_bdb := 2 * _e.Pi / float64(_ddab)
	_dfcd -= _e.Pi / 2
	if _ddab%2 == 0 {
		_dfcd += _bdb / 2
	}
	_dae.NewSubPath()
	for _eage := 0; _eage < _ddab; _eage++ {
		_begd := _dfcd + _bdb*float64(_eage)
		_dae.LineTo(_fcb+_gggg*_e.Cos(_begd), _aefc+_gggg*_e.Sin(_begd))
	}
	_dae.ClosePath()
}
func (_dece *Context) SetLineCap(lineCap _g.LineCap) { _dece._bae = lineCap }
func _eef(_bgcf float64) float64                     { return _bgcf * _e.Pi / 180 }
func (_cbbg *Context) SetPixel(x, y int)             { _cbbg._feg.Set(x, y, _cbbg._egg) }
func (_bda *Context) SetHexColor(x string) {
	_dda, _bce, _gde, _bedc := _gbg(x)
	_bda.SetRGBA255(_dda, _bce, _gde, _bedc)
}
func (_dg *Context) StrokePreserve() {
	var _cad _dc.Painter
	if _dg._bba == nil {
		if _fbe, _ee := _dg._fg.(*solidPattern); _ee {
			_cadc := _dc.NewRGBAPainter(_dg._feg)
			_cadc.SetColor(_fbe._agde)
			_cad = _cadc
		}
	}
	if _cad == nil {
		_cad = _efeb(_dg._feg, _dg._bba, _dg._fg)
	}
	_dg.stroke(_cad)
}
func (_cgb *Context) SetFillStyle(pattern _g.Pattern) {
	if _eab, _cgc := pattern.(*solidPattern); _cgc {
		_cgb._egg = _eab._agde
	}
	_cgb._bfeg = pattern
}
func _edd(_agbd _be.Int26_6) float64 {
	const _cfe, _gebe = 6, 1<<6 - 1
	if _agbd >= 0 {
		return float64(_agbd>>_cfe) + float64(_agbd&_gebe)/64
	}
	_agbd = -_agbd
	if _agbd >= 0 {
		return -(float64(_agbd>>_cfe) + float64(_agbd&_gebe)/64)
	}
	return 0
}
func _dfe(_cg, _bg, _db, _eg, _ce, _ab float64) []_f.Point {
	_ddb := (_e.Hypot(_db-_cg, _eg-_bg) + _e.Hypot(_ce-_db, _ab-_eg))
	_bfe := int(_ddb + 0.5)
	if _bfe < 4 {
		_bfe = 4
	}
	_fbg := float64(_bfe) - 1
	_cb := make([]_f.Point, _bfe)
	for _bd := 0; _bd < _bfe; _bd++ {
		_ec := float64(_bd) / _fbg
		_fe, _gfe := _dd(_cg, _bg, _db, _eg, _ce, _ab, _ec)
		_cb[_bd] = _f.NewPoint(_fe, _gfe)
	}
	return _cb
}
func (_egb *Context) SetStrokeStyle(pattern _g.Pattern) { _egb._fg = pattern }
func (_aedg *Context) Pop() {
	_daeb := *_aedg
	_gggf := _aedg._fbc
	_agbe := _gggf[len(_gggf)-1]
	*_aedg = *_agbe
	_aedg._ge = _daeb._ge
	_aedg._af = _daeb._af
	_aedg._fbfg = _daeb._fbfg
	_aedg._bbc = _daeb._bbc
	_aedg._gdd = _daeb._gdd
}
func (_fed *Context) capper() _dc.Capper {
	switch _fed._bae {
	case _g.LineCapButt:
		return _dc.ButtCapper
	case _g.LineCapRound:
		return _dc.RoundCapper
	case _g.LineCapSquare:
		return _dc.SquareCapper
	}
	return nil
}
func (_abcd *Context) SetRGB(r, g, b float64) { _abcd.SetRGBA(r, g, b, 1) }
func (_dgf *patternPainter) Paint(ss []_dc.Span, done bool) {
	_fdab := _dgf._bdaff.Bounds()
	for _, _feb := range ss {
		if _feb.Y < _fdab.Min.Y {
			continue
		}
		if _feb.Y >= _fdab.Max.Y {
			return
		}
		if _feb.X0 < _fdab.Min.X {
			_feb.X0 = _fdab.Min.X
		}
		if _feb.X1 > _fdab.Max.X {
			_feb.X1 = _fdab.Max.X
		}
		if _feb.X0 >= _feb.X1 {
			continue
		}
		const _ceb = 1<<16 - 1
		_gfgf := _feb.Y - _dgf._bdaff.Rect.Min.Y
		_dge := _feb.X0 - _dgf._bdaff.Rect.Min.X
		_gfgfg := (_feb.Y-_dgf._bdaff.Rect.Min.Y)*_dgf._bdaff.Stride + (_feb.X0-_dgf._bdaff.Rect.Min.X)*4
		_fdae := _gfgfg + (_feb.X1-_feb.X0)*4
		for _bfb, _cdfe := _gfgfg, _dge; _bfb < _fdae; _bfb, _cdfe = _bfb+4, _cdfe+1 {
			_dbb := _feb.Alpha
			if _dgf._bcef != nil {
				_dbb = _dbb * uint32(_dgf._bcef.AlphaAt(_cdfe, _gfgf).A) / 255
				if _dbb == 0 {
					continue
				}
			}
			_gcgg := _dgf._bfdg.ColorAt(_cdfe, _gfgf)
			_gge, _edc, _bcge, _ebfd := _gcgg.RGBA()
			_gfega := uint32(_dgf._bdaff.Pix[_bfb+0])
			_bfdab := uint32(_dgf._bdaff.Pix[_bfb+1])
			_dgeb := uint32(_dgf._bdaff.Pix[_bfb+2])
			_bdc := uint32(_dgf._bdaff.Pix[_bfb+3])
			_fac := (_ceb - (_ebfd * _dbb / _ceb)) * 0x101
			_dgf._bdaff.Pix[_bfb+0] = uint8((_gfega*_fac + _gge*_dbb) / _ceb >> 8)
			_dgf._bdaff.Pix[_bfb+1] = uint8((_bfdab*_fac + _edc*_dbb) / _ceb >> 8)
			_dgf._bdaff.Pix[_bfb+2] = uint8((_dgeb*_fac + _bcge*_dbb) / _ceb >> 8)
			_dgf._bdaff.Pix[_bfb+3] = uint8((_bdc*_fac + _ebfd*_dbb) / _ceb >> 8)
		}
	}
}
func (_ggg *Context) SetDashOffset(offset float64) { _ggg._bc = offset }
func _gbg(_aca string) (_baf, _ddeg, _begf, _gab int) {
	_aca = _eb.TrimPrefix(_aca, "\u0023")
	_gab = 255
	if len(_aca) == 3 {
		_eced := "\u00251\u0078\u0025\u0031\u0078\u0025\u0031x"
		_df.Sscanf(_aca, _eced, &_baf, &_ddeg, &_begf)
		_baf |= _baf << 4
		_ddeg |= _ddeg << 4
		_begf |= _begf << 4
	}
	if len(_aca) == 6 {
		_ecd := "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_df.Sscanf(_aca, _ecd, &_baf, &_ddeg, &_begf)
	}
	if len(_aca) == 8 {
		_aebb := "\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_df.Sscanf(_aca, _aebb, &_baf, &_ddeg, &_begf, &_gab)
	}
	return
}
func _eccf(_bcfb _a.Color) _g.Pattern { return &solidPattern{_agde: _bcfb} }

type patternPainter struct {
	_bdaff *_c.RGBA
	_bcef  *_c.Alpha
	_bfdg  _g.Pattern
}

func _dd(_gf, _fa, _gd, _fb, _ca, _ff, _dfd float64) (_aee, _fd float64) {
	_fcd := 1 - _dfd
	_cc := _fcd * _fcd
	_bf := 2 * _fcd * _dfd
	_gfc := _dfd * _dfd
	_aee = _cc*_gf + _bf*_gd + _gfc*_ca
	_fd = _cc*_fa + _bf*_fb + _gfc*_ff
	return
}
func (_ecf *Context) SetRGB255(r, g, b int) { _ecf.SetRGBA255(r, g, b, 255) }

type stops []stop

func (_gfeb *Context) FillPattern() _g.Pattern  { return _gfeb._bfeg }
func (_dad *Context) SetDash(dashes ...float64) { _dad._ea = dashes }
