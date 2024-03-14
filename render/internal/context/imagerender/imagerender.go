package imagerender

import (
	_ca "errors"
	_ab "fmt"
	_ad "image"
	_ce "image/color"
	_g "image/draw"
	_d "math"
	_af "sort"
	_c "strings"

	_ec "bitbucket.org/shenghui0779/gopdf/common"
	_e "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_adf "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_ed "github.com/unidoc/freetype/raster"
	_f "golang.org/x/image/draw"
	_ga "golang.org/x/image/font"
	_ff "golang.org/x/image/math/f64"
	_gf "golang.org/x/image/math/fixed"
)

func (_ede *Context) SetStrokeStyle(pattern _adf.Pattern) { _ede._gbf = pattern }
func _ge(_ba, _db, _cd, _gaf, _cee, _gc, _ac, _dde, _cec float64) (_gaa, _bd float64) {
	_fga := 1 - _cec
	_ege := _fga * _fga * _fga
	_be := 3 * _fga * _fga * _cec
	_egc := 3 * _fga * _cec * _cec
	_fgcf := _cec * _cec * _cec
	_gaa = _ege*_ba + _be*_cd + _egc*_cee + _fgcf*_ac
	_bd = _ege*_db + _be*_gaf + _egc*_gc + _fgcf*_dde
	return
}

func (_fba *Context) SetFillStyle(pattern _adf.Pattern) {
	if _dfab, _bbde := pattern.(*solidPattern); _bbde {
		_fba._acc = _dfab._fcg
	}
	_fba._adb = pattern
}
func (_adbc *Context) SetPixel(x, y int) { _adbc._gfd.Set(x, y, _adbc._acc) }

type patternPainter struct {
	_dbb  *_ad.RGBA
	_gfa  *_ad.Alpha
	_ebdg _adf.Pattern
}

func (_bbebd *radialGradient) ColorAt(x, y int) _ce.Color {
	if len(_bbebd._abaa) == 0 {
		return _ce.Transparent
	}
	_dfd, _bcb := float64(x)+0.5-_bbebd._ggbb._aeg, float64(y)+0.5-_bbebd._ggbb._cdc
	_gec := _fdba(_dfd, _bcb, _bbebd._ggbb._ecd, _bbebd._fce._aeg, _bbebd._fce._cdc, _bbebd._fce._ecd)
	_fafc := _fdba(_dfd, _bcb, -_bbebd._ggbb._ecd, _dfd, _bcb, _bbebd._ggbb._ecd)
	if _bbebd._gdc == 0 {
		if _gec == 0 {
			return _ce.Transparent
		}
		_ddd := 0.5 * _fafc / _gec
		if _ddd*_bbebd._fce._ecd >= _bbebd._fca {
			return _age(_ddd, _bbebd._abaa)
		}
		return _ce.Transparent
	}
	_dgb := _fdba(_gec, _bbebd._gdc, 0, _gec, -_fafc, 0)
	if _dgb >= 0 {
		_cde := _d.Sqrt(_dgb)
		_bcg := (_gec + _cde) * _bbebd._fdb
		_gbgb := (_gec - _cde) * _bbebd._fdb
		if _bcg*_bbebd._fce._ecd >= _bbebd._fca {
			return _age(_bcg, _bbebd._abaa)
		} else if _gbgb*_bbebd._fce._ecd >= _bbebd._fca {
			return _age(_gbgb, _bbebd._abaa)
		}
	}
	return _ce.Transparent
}

func _ebafd(_bdec, _gbfg uint32, _egg float64) uint8 {
	return uint8(int32(float64(_bdec)*(1.0-_egg)+float64(_gbfg)*_egg) >> 8)
}
func (_dec *Context) FillPattern() _adf.Pattern { return _dec._adb }
func (_ceae *Context) ShearAbout(sx, sy, x, y float64) {
	_ceae.Translate(x, y)
	_ceae.Shear(sx, sy)
	_ceae.Translate(-x, -y)
}

func (_ggd *Context) AsMask() *_ad.Alpha {
	_gafb := _ad.NewAlpha(_ggd._gfd.Bounds())
	_f.Draw(_gafb, _ggd._gfd.Bounds(), _ggd._gfd, _ad.Point{}, _f.Src)
	return _gafb
}

func (_daa *Context) StrokePreserve() {
	var _cbda _ed.Painter
	if _daa._fded == nil {
		if _bca, _eec := _daa._gbf.(*solidPattern); _eec {
			_adfe := _ed.NewRGBAPainter(_daa._gfd)
			_adfe.SetColor(_bca._fcg)
			_cbda = _adfe
		}
	}
	if _cbda == nil {
		_cbda = _fcc(_daa._gfd, _daa._fded, _daa._gbf)
	}
	_daa.stroke(_cbda)
}

func (_dcd *Context) InvertMask() {
	if _dcd._fded == nil {
		_dcd._fded = _ad.NewAlpha(_dcd._gfd.Bounds())
	} else {
		for _eece, _bad := range _dcd._fded.Pix {
			_dcd._fded.Pix[_eece] = 255 - _bad
		}
	}
}

func _gcff(_bcbe [][]_e.Point) _ed.Path {
	var _ecdf _ed.Path
	for _, _cca := range _bcbe {
		var _baa _gf.Point26_6
		for _afdg, _ceab := range _cca {
			_ggdg := _edbe(_ceab)
			if _afdg == 0 {
				_ecdf.Start(_ggdg)
			} else {
				_decg := _ggdg.X - _baa.X
				_cbc := _ggdg.Y - _baa.Y
				if _decg < 0 {
					_decg = -_decg
				}
				if _cbc < 0 {
					_cbc = -_cbc
				}
				if _decg+_cbc > 8 {
					_ecdf.Add1(_ggdg)
				}
			}
			_baa = _ggdg
		}
	}
	return _ecdf
}
func (_dag *Context) SetDashOffset(offset float64) { _dag._aa = offset }
func _eed(_aae _ad.Image, _edd repeatOp) _adf.Pattern {
	return &surfacePattern{_ecfa: _aae, _efbf: _edd}
}
func (_cga *Context) Shear(x, y float64) { _cga._bdf.Shear(x, y) }
func (_abc *Context) Image() _ad.Image   { return _abc._gfd }

type solidPattern struct{ _fcg _ce.Color }

func (_babe *Context) Stroke()                        { _babe.StrokePreserve(); _babe.ClearPath() }
func _dddf(_cdbd float64) _gf.Int26_6                 { return _gf.Int26_6(_cdbd * 64) }
func (_ggf *Context) SetLineCap(lineCap _adf.LineCap) { _ggf._caa = lineCap }
func (_ddb *solidPattern) ColorAt(x, y int) _ce.Color { return _ddb._fcg }
func (_ced *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	if !_ced._eae {
		_ced.MoveTo(x1, y1)
	}
	_dea, _dcg := _ced._dcf.X, _ced._dcf.Y
	x1, y1 = _ced.Transform(x1, y1)
	x2, y2 = _ced.Transform(x2, y2)
	x3, y3 = _ced.Transform(x3, y3)
	_dagg := _bc(_dea, _dcg, x1, y1, x2, y2, x3, y3)
	_add := _edbe(_ced._dcf)
	for _, _dgce := range _dagg[1:] {
		_ceea := _edbe(_dgce)
		if _ceea == _add {
			continue
		}
		_add = _ceea
		_ced._gg.Add1(_ceea)
		_ced._df.Add1(_ceea)
		_ced._dcf = _dgce
	}
}
func _ecg(_cdbb _ce.Color) _adf.Pattern { return &solidPattern{_fcg: _cdbb} }
func (_cbe *Context) Height() int       { return _cbe._fdc }
func NewRadialGradient(x0, y0, r0, x1, y1, r1 float64) _adf.Gradient {
	_cfb := circle{x0, y0, r0}
	_acega := circle{x1, y1, r1}
	_dacc := circle{x1 - x0, y1 - y0, r1 - r0}
	_eeb := _fdba(_dacc._aeg, _dacc._cdc, -_dacc._ecd, _dacc._aeg, _dacc._cdc, _dacc._ecd)
	var _bdfc float64
	if _eeb != 0 {
		_bdfc = 1.0 / _eeb
	}
	_abbd := -_cfb._ecd
	_facf := &radialGradient{_ggbb: _cfb, _acb: _acega, _fce: _dacc, _gdc: _eeb, _fdb: _bdfc, _fca: _abbd}
	return _facf
}
func (_bf *Context) SetColor(c _ce.Color)        { _bf.setFillAndStrokeColor(c) }
func (_bgb *Context) TextState() *_adf.TextState { return &_bgb._acf }
func _edbe(_ebe _e.Point) _gf.Point26_6          { return _gf.Point26_6{X: _dddf(_ebe.X), Y: _dddf(_ebe.Y)} }
func (_dge *Context) DrawRectangle(x, y, w, h float64) {
	_dge.NewSubPath()
	_dge.MoveTo(x, y)
	_dge.LineTo(x+w, y)
	_dge.LineTo(x+w, y+h)
	_dge.LineTo(x, y+h)
	_dge.ClosePath()
}

func (_dbc *Context) SetFillRGBA(r, g, b, a float64) {
	_dbf := _ce.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_dbc._acc = _dbf
	_dbc._adb = _ecg(_dbf)
}
func (_da *Context) SetDash(dashes ...float64) { _da._cef = dashes }

type surfacePattern struct {
	_ecfa _ad.Image
	_efbf repeatOp
}

func NewLinearGradient(x0, y0, x1, y1 float64) _adf.Gradient {
	_fec := &linearGradient{_egf: x0, _bdeb: y0, _cedd: x1, _fda: y1}
	return _fec
}

func _febf(_ecfd _ed.Path) [][]_e.Point {
	var _efde [][]_e.Point
	var _cgf []_e.Point
	var _adg, _ecaa float64
	for _afc := 0; _afc < len(_ecfd); {
		switch _ecfd[_afc] {
		case 0:
			if len(_cgf) > 0 {
				_efde = append(_efde, _cgf)
				_cgf = nil
			}
			_cggf := _gbeb(_ecfd[_afc+1])
			_baccd := _gbeb(_ecfd[_afc+2])
			_cgf = append(_cgf, _e.NewPoint(_cggf, _baccd))
			_adg, _ecaa = _cggf, _baccd
			_afc += 4
		case 1:
			_bce := _gbeb(_ecfd[_afc+1])
			_bafa := _gbeb(_ecfd[_afc+2])
			_cgf = append(_cgf, _e.NewPoint(_bce, _bafa))
			_adg, _ecaa = _bce, _bafa
			_afc += 4
		case 2:
			_gfff := _gbeb(_ecfd[_afc+1])
			_baef := _gbeb(_ecfd[_afc+2])
			_fdaf := _gbeb(_ecfd[_afc+3])
			_eea := _gbeb(_ecfd[_afc+4])
			_dcb := _ffc(_adg, _ecaa, _gfff, _baef, _fdaf, _eea)
			_cgf = append(_cgf, _dcb...)
			_adg, _ecaa = _fdaf, _eea
			_afc += 6
		case 3:
			_dab := _gbeb(_ecfd[_afc+1])
			_cdcd := _gbeb(_ecfd[_afc+2])
			_egcb := _gbeb(_ecfd[_afc+3])
			_dfgc := _gbeb(_ecfd[_afc+4])
			_beaf := _gbeb(_ecfd[_afc+5])
			_aaa := _gbeb(_ecfd[_afc+6])
			_ead := _bc(_adg, _ecaa, _dab, _cdcd, _egcb, _dfgc, _beaf, _aaa)
			_cgf = append(_cgf, _ead...)
			_adg, _ecaa = _beaf, _aaa
			_afc += 8
		default:
			_ec.Log.Debug("\u0057\u0041\u0052\u004e: \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0074\u0068\u003a\u0020%\u0076", _ecfd)
			return _efde
		}
	}
	if len(_cgf) > 0 {
		_efde = append(_efde, _cgf)
	}
	return _efde
}
func (_ebae *Context) SetRGB255(r, g, b int)             { _ebae.SetRGBA255(r, g, b, 255) }
func (_gbg *Context) SetLineJoin(lineJoin _adf.LineJoin) { _gbg._bbd = lineJoin }
func (_ffca stops) Len() int                             { return len(_ffca) }
func (_ebd *Context) SetHexColor(x string) {
	_ffa, _fbf, _cad, _abb := _aaff(x)
	_ebd.SetRGBA255(_ffa, _fbf, _cad, _abb)
}

func (_aaf *Context) DrawString(s string, face _ga.Face, x, y float64) {
	_aaf.DrawStringAnchored(s, face, x, y, 0, 0)
}

func (_ffg *Context) ScaleAbout(sx, sy, x, y float64) {
	_ffg.Translate(x, y)
	_ffg.Scale(sx, sy)
	_ffg.Translate(-x, -y)
}

type stops []stop

func _aaff(_fgeea string) (_ged, _gggg, _beafa, _deaa int) {
	_fgeea = _c.TrimPrefix(_fgeea, "\u0023")
	_deaa = 255
	if len(_fgeea) == 3 {
		_daeg := "\u00251\u0078\u0025\u0031\u0078\u0025\u0031x"
		_ab.Sscanf(_fgeea, _daeg, &_ged, &_gggg, &_beafa)
		_ged |= _ged << 4
		_gggg |= _gggg << 4
		_beafa |= _beafa << 4
	}
	if len(_fgeea) == 6 {
		_gfeg := "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_ab.Sscanf(_fgeea, _gfeg, &_ged, &_gggg, &_beafa)
	}
	if len(_fgeea) == 8 {
		_eecb := "\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_ab.Sscanf(_fgeea, _eecb, &_ged, &_gggg, &_beafa, &_deaa)
	}
	return
}

func _age(_cag float64, _baea stops) _ce.Color {
	if _cag <= 0.0 || len(_baea) == 1 {
		return _baea[0]._dbac
	}
	_dfc := _baea[len(_baea)-1]
	if _cag >= _dfc._aba {
		return _dfc._dbac
	}
	for _dfg, _bacc := range _baea[1:] {
		if _cag < _bacc._aba {
			_cag = (_cag - _baea[_dfg]._aba) / (_bacc._aba - _baea[_dfg]._aba)
			return _cedc(_baea[_dfg]._dbac, _bacc._dbac, _cag)
		}
	}
	return _dfc._dbac
}

func (_abbe *Context) DrawArc(x, y, r, angle1, angle2 float64) {
	_abbe.DrawEllipticalArc(x, y, r, r, angle1, angle2)
}

func (_gcef *Context) DrawCircle(x, y, r float64) {
	_gcef.NewSubPath()
	_gcef.DrawEllipticalArc(x, y, r, r, 0, 2*_d.Pi)
	_gcef.ClosePath()
}

func (_dbcd *linearGradient) AddColorStop(offset float64, color _ce.Color) {
	_dbcd._ddg = append(_dbcd._ddg, stop{_aba: offset, _dbac: color})
	_af.Sort(_dbcd._ddg)
}
func (_bbeb *Context) Push() { _gba := *_bbeb; _bbeb._ddef = append(_bbeb._ddef, &_gba) }
func (_ffb *radialGradient) AddColorStop(offset float64, color _ce.Color) {
	_ffb._abaa = append(_ffb._abaa, stop{_aba: offset, _dbac: color})
	_af.Sort(_ffb._abaa)
}

func (_cba *Context) MeasureString(s string, face _ga.Face) (_agd, _eaf float64) {
	_edff := &_ga.Drawer{Face: face}
	_gfg := _edff.MeasureString(s)
	return float64(_gfg >> 6), _cba._acf.Tf.Size
}

func _ffc(_fgc, _cg, _cf, _cc, _dca, _gb float64) []_e.Point {
	_cea := (_d.Hypot(_cf-_fgc, _cc-_cg) + _d.Hypot(_dca-_cf, _gb-_cc))
	_ef := int(_cea + 0.5)
	if _ef < 4 {
		_ef = 4
	}
	_gag := float64(_ef) - 1
	_fb := make([]_e.Point, _ef)
	for _fdf := 0; _fdf < _ef; _fdf++ {
		_fac := float64(_fdf) / _gag
		_dg, _ea := _fg(_fgc, _cg, _cf, _cc, _dca, _gb, _fac)
		_fb[_fdf] = _e.NewPoint(_dg, _ea)
	}
	return _fb
}

func _bc(_gcf, _bef, _acd, _cb, _cbd, _abd, _cgc, _bee float64) []_e.Point {
	_fde := (_d.Hypot(_acd-_gcf, _cb-_bef) + _d.Hypot(_cbd-_acd, _abd-_cb) + _d.Hypot(_cgc-_cbd, _bee-_abd))
	_cac := int(_fde + 0.5)
	if _cac < 4 {
		_cac = 4
	}
	_dgc := float64(_cac) - 1
	_bab := make([]_e.Point, _cac)
	for _gdd := 0; _gdd < _cac; _gdd++ {
		_gbe := float64(_gdd) / _dgc
		_ddec, _fc := _ge(_gcf, _bef, _acd, _cb, _cbd, _abd, _cgc, _bee, _gbe)
		_bab[_gdd] = _e.NewPoint(_ddec, _fc)
	}
	return _bab
}

func (_efc *Context) stroke(_gae _ed.Painter) {
	_dda := _efc._gg
	if len(_efc._cef) > 0 {
		_dda = _dcc(_dda, _efc._cef, _efc._aa)
	} else {
		_dda = _gcff(_febf(_dda))
	}
	_afg := _efc._ee
	_afg.UseNonZeroWinding = true
	_afg.Clear()
	_acce := (_efc._bdf.ScalingFactorX() + _efc._bdf.ScalingFactorY()) / 2
	_afg.AddStroke(_dda, _dddf(_efc._afe*_acce), _efc.capper(), _efc.joiner())
	_afg.Rasterize(_gae)
}

func (_baf *Context) SetRGBA255(r, g, b, a int) {
	_baf._acc = _ce.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	_baf.setFillAndStrokeColor(_baf._acc)
}

func (_dae *Context) ClipPreserve() {
	_ceg := _ad.NewAlpha(_ad.Rect(0, 0, _dae._gee, _dae._fdc))
	_cbeg := _ed.NewAlphaOverPainter(_ceg)
	_dae.fill(_cbeg)
	if _dae._fded == nil {
		_dae._fded = _ceg
	} else {
		_fgd := _ad.NewAlpha(_ad.Rect(0, 0, _dae._gee, _dae._fdc))
		_f.DrawMask(_fgd, _fgd.Bounds(), _ceg, _ad.Point{}, _dae._fded, _ad.Point{}, _f.Over)
		_dae._fded = _fgd
	}
}

type radialGradient struct {
	_ggbb, _acb, _fce circle
	_gdc, _fdb        float64
	_fca              float64
	_abaa             stops
}

func (_dfa *Context) SetLineWidth(lineWidth float64)    { _dfa._afe = lineWidth }
func (_gabc *Context) DrawImage(im _ad.Image, x, y int) { _gabc.DrawImageAnchored(im, x, y, 0, 0) }
func _fg(_fe, _b, _dd, _fa, _gab, _fd, _ae float64) (_eb, _bb float64) {
	_gd := 1 - _ae
	_eg := _gd * _gd
	_ag := 2 * _gd * _ae
	_dc := _ae * _ae
	_eb = _eg*_fe + _ag*_dd + _dc*_gab
	_bb = _eg*_b + _ag*_fa + _dc*_fd
	return
}

func (_edfg *Context) DrawPoint(x, y, r float64) {
	_edfg.Push()
	_gagf, _dfb := _edfg.Transform(x, y)
	_edfg.Identity()
	_edfg.DrawCircle(_gagf, _dfb, r)
	_edfg.Pop()
}

func (_bec *Context) SetStrokeRGBA(r, g, b, a float64) {
	_afa := _ce.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_bec._gbf = _ecg(_afa)
}
func (_eaeb *Context) ClearPath()                      { _eaeb._gg.Clear(); _eaeb._df.Clear(); _eaeb._eae = false }
func (_edfd *Context) DrawLine(x1, y1, x2, y2 float64) { _edfd.MoveTo(x1, y1); _edfd.LineTo(x2, y2) }
func (_ebaf *Context) Rotate(angle float64)            { _ebaf._bdf = _ebaf._bdf.Rotate(angle) }
func (_feg *Context) DrawRoundedRectangle(x, y, w, h, r float64) {
	_bac, _dac, _dcge, _eff := x, x+r, x+w-r, x+w
	_gfb, _bde, _fdd, _cbde := y, y+r, y+h-r, y+h
	_feg.NewSubPath()
	_feg.MoveTo(_dac, _gfb)
	_feg.LineTo(_dcge, _gfb)
	_feg.DrawArc(_dcge, _bde, r, _baee(270), _baee(360))
	_feg.LineTo(_eff, _fdd)
	_feg.DrawArc(_dcge, _fdd, r, _baee(0), _baee(90))
	_feg.LineTo(_dac, _cbde)
	_feg.DrawArc(_dac, _fdd, r, _baee(90), _baee(180))
	_feg.LineTo(_bac, _bde)
	_feg.DrawArc(_dac, _bde, r, _baee(180), _baee(270))
	_feg.ClosePath()
}
func (_gcd *Context) Scale(x, y float64) { _gcd._bdf = _gcd._bdf.Scale(x, y) }
func _fdba(_fbb, _fafg, _bdba, _afaf, _egfc, _eafg float64) float64 {
	return _fbb*_afaf + _fafg*_egfc + _bdba*_eafg
}

func (_caac *Context) drawString(_ecf string, _aceg _ga.Face, _eagc, _gbff float64) {
	_caag := &_ga.Drawer{Src: _ad.NewUniform(_caac._acc), Face: _aceg, Dot: _edbe(_e.NewPoint(_eagc, _gbff))}
	_fge := rune(-1)
	for _, _ccge := range _ecf {
		if _fge >= 0 {
			_caag.Dot.X += _caag.Face.Kern(_fge, _ccge)
		}
		_fea, _aec, _beae, _efd, _ebba := _caag.Face.Glyph(_caag.Dot, _ccge)
		if !_ebba {
			continue
		}
		_cfcc := _fea.Sub(_fea.Min)
		_ecef := _ad.NewRGBA(_cfcc)
		_f.DrawMask(_ecef, _cfcc, _caag.Src, _ad.Point{}, _aec, _beae, _f.Over)
		var _cbec *_f.Options
		if _caac._fded != nil {
			_cbec = &_f.Options{DstMask: _caac._fded, DstMaskP: _ad.Point{}}
		}
		_afbc := _caac._bdf.Clone().Translate(float64(_fea.Min.X), float64(_fea.Min.Y))
		_fee := _ff.Aff3{_afbc[0], _afbc[3], _afbc[6], _afbc[1], _afbc[4], _afbc[7]}
		_f.BiLinear.Transform(_caac._gfd, _fee, _ecef, _cfcc, _f.Over, _cbec)
		_caag.Dot.X += _efd
		_fge = _ccge
	}
}

type repeatOp int

func (_gbc *Context) ResetClip()                         { _gbc._fded = nil }
func (_bdaa *Context) SetMatrix(m _e.Matrix)             { _bdaa._bdf = m }
func (_afb *Context) StrokePattern() _adf.Pattern        { return _afb._gbf }
func (_faa *Context) SetFillRule(fillRule _adf.FillRule) { _faa._cdf = fillRule }
func (_bae stops) Less(i, j int) bool                    { return _bae[i]._aba < _bae[j]._aba }
func (_ddee stops) Swap(i, j int)                        { _ddee[i], _ddee[j] = _ddee[j], _ddee[i] }
func (_edf *Context) SetRGBA(r, g, b, a float64) {
	_edf._acc = _ce.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_edf.setFillAndStrokeColor(_edf._acc)
}

const (
	_fbd repeatOp = iota
	_ccd
	_ecde
	_afba
)

func (_cgd *Context) joiner() _ed.Joiner {
	switch _cgd._bbd {
	case _adf.LineJoinBevel:
		return _ed.BevelJoiner
	case _adf.LineJoinRound:
		return _ed.RoundJoiner
	}
	return nil
}

func (_ecfaa *patternPainter) Paint(ss []_ed.Span, done bool) {
	_dgf := _ecfaa._dbb.Bounds()
	for _, _ecdg := range ss {
		if _ecdg.Y < _dgf.Min.Y {
			continue
		}
		if _ecdg.Y >= _dgf.Max.Y {
			return
		}
		if _ecdg.X0 < _dgf.Min.X {
			_ecdg.X0 = _dgf.Min.X
		}
		if _ecdg.X1 > _dgf.Max.X {
			_ecdg.X1 = _dgf.Max.X
		}
		if _ecdg.X0 >= _ecdg.X1 {
			continue
		}
		const _bcbc = 1<<16 - 1
		_ceed := _ecdg.Y - _ecfaa._dbb.Rect.Min.Y
		_ebf := _ecdg.X0 - _ecfaa._dbb.Rect.Min.X
		_aeb := (_ecdg.Y-_ecfaa._dbb.Rect.Min.Y)*_ecfaa._dbb.Stride + (_ecdg.X0-_ecfaa._dbb.Rect.Min.X)*4
		_dcga := _aeb + (_ecdg.X1-_ecdg.X0)*4
		for _aafa, _fcgg := _aeb, _ebf; _aafa < _dcga; _aafa, _fcgg = _aafa+4, _fcgg+1 {
			_ffad := _ecdg.Alpha
			if _ecfaa._gfa != nil {
				_ffad = _ffad * uint32(_ecfaa._gfa.AlphaAt(_fcgg, _ceed).A) / 255
				if _ffad == 0 {
					continue
				}
			}
			_eaa := _ecfaa._ebdg.ColorAt(_fcgg, _ceed)
			_ddaf, _aea, _dfge, _gfae := _eaa.RGBA()
			_abbf := uint32(_ecfaa._dbb.Pix[_aafa+0])
			_aga := uint32(_ecfaa._dbb.Pix[_aafa+1])
			_gfe := uint32(_ecfaa._dbb.Pix[_aafa+2])
			_dcdg := uint32(_ecfaa._dbb.Pix[_aafa+3])
			_fgca := (_bcbc - (_gfae * _ffad / _bcbc)) * 0x101
			_ecfaa._dbb.Pix[_aafa+0] = uint8((_abbf*_fgca + _ddaf*_ffad) / _bcbc >> 8)
			_ecfaa._dbb.Pix[_aafa+1] = uint8((_aga*_fgca + _aea*_ffad) / _bcbc >> 8)
			_ecfaa._dbb.Pix[_aafa+2] = uint8((_gfe*_fgca + _dfge*_ffad) / _bcbc >> 8)
			_ecfaa._dbb.Pix[_aafa+3] = uint8((_dcdg*_fgca + _gfae*_ffad) / _bcbc >> 8)
		}
	}
}

func (_gagc *Context) DrawImageAnchored(im _ad.Image, x, y int, ax, ay float64) {
	_dagf := im.Bounds().Size()
	x -= int(ax * float64(_dagf.X))
	y -= int(ay * float64(_dagf.Y))
	_ecaf := _f.BiLinear
	_cdg := _gagc._bdf.Clone().Translate(float64(x), float64(y))
	_gdg := _ff.Aff3{_cdg[0], _cdg[3], _cdg[6], _cdg[1], _cdg[4], _cdg[7]}
	if _gagc._fded == nil {
		_ecaf.Transform(_gagc._gfd, _gdg, im, im.Bounds(), _f.Over, nil)
	} else {
		_ecaf.Transform(_gagc._gfd, _gdg, im, im.Bounds(), _f.Over, &_f.Options{DstMask: _gagc._fded, DstMaskP: _ad.Point{}})
	}
}
func (_cbdc *Context) SetRGB(r, g, b float64) { _cbdc.SetRGBA(r, g, b, 1) }
func (_daad *linearGradient) ColorAt(x, y int) _ce.Color {
	if len(_daad._ddg) == 0 {
		return _ce.Transparent
	}
	_faad, _fdfc := float64(x), float64(y)
	_fab, _gafd, _gca, _ecec := _daad._egf, _daad._bdeb, _daad._cedd, _daad._fda
	_cedde, _eac := _gca-_fab, _ecec-_gafd
	if _eac == 0 && _cedde != 0 {
		return _age((_faad-_fab)/_cedde, _daad._ddg)
	}
	if _cedde == 0 && _eac != 0 {
		return _age((_fdfc-_gafd)/_eac, _daad._ddg)
	}
	_ada := _cedde*(_faad-_fab) + _eac*(_fdfc-_gafd)
	if _ada < 0 {
		return _daad._ddg[0]._dbac
	}
	_eaeff := _d.Hypot(_cedde, _eac)
	_gceb := ((_faad-_fab)*-_eac + (_fdfc-_gafd)*_cedde) / (_eaeff * _eaeff)
	_aafb, _faca := _fab+_gceb*-_eac, _gafd+_gceb*_cedde
	_gafbg := _d.Hypot(_faad-_aafb, _fdfc-_faca) / _eaeff
	return _age(_gafbg, _daad._ddg)
}

func (_cege *Context) drawRegularPolygon(_ece int, _gffa, _bdbd, _bba, _ecce float64) {
	_bbg := 2 * _d.Pi / float64(_ece)
	_ecce -= _d.Pi / 2
	if _ece%2 == 0 {
		_ecce += _bbg / 2
	}
	_cege.NewSubPath()
	for _aad := 0; _aad < _ece; _aad++ {
		_gffad := _ecce + _bbg*float64(_aad)
		_cege.LineTo(_gffa+_bba*_d.Cos(_gffad), _bdbd+_bba*_d.Sin(_gffad))
	}
	_cege.ClosePath()
}

func _ddfe(_bbac [][]_e.Point, _eafb []float64, _fdec float64) [][]_e.Point {
	var _eeae [][]_e.Point
	if len(_eafb) == 0 {
		return _bbac
	}
	if len(_eafb) == 1 {
		_eafb = append(_eafb, _eafb[0])
	}
	for _, _ggg := range _bbac {
		if len(_ggg) < 2 {
			continue
		}
		_bbaf := _ggg[0]
		_acee := 1
		_gdcg := 0
		_fgee := 0.0
		if _fdec != 0 {
			var _fgda float64
			for _, _dce := range _eafb {
				_fgda += _dce
			}
			_fdec = _d.Mod(_fdec, _fgda)
			if _fdec < 0 {
				_fdec += _fgda
			}
			for _ddda, _cbdeg := range _eafb {
				_fdec -= _cbdeg
				if _fdec < 0 {
					_gdcg = _ddda
					_fgee = _cbdeg + _fdec
					break
				}
			}
		}
		var _gaaf []_e.Point
		_gaaf = append(_gaaf, _bbaf)
		for _acee < len(_ggg) {
			_ccca := _eafb[_gdcg]
			_ecb := _ggg[_acee]
			_ecad := _bbaf.Distance(_ecb)
			_fgg := _ccca - _fgee
			if _ecad > _fgg {
				_gddf := _fgg / _ecad
				_ggc := _bbaf.Interpolate(_ecb, _gddf)
				_gaaf = append(_gaaf, _ggc)
				if _gdcg%2 == 0 && len(_gaaf) > 1 {
					_eeae = append(_eeae, _gaaf)
				}
				_gaaf = nil
				_gaaf = append(_gaaf, _ggc)
				_fgee = 0
				_bbaf = _ggc
				_gdcg = (_gdcg + 1) % len(_eafb)
			} else {
				_gaaf = append(_gaaf, _ecb)
				_bbaf = _ecb
				_fgee += _ecad
				_acee++
			}
		}
		if _gdcg%2 == 0 && len(_gaaf) > 1 {
			_eeae = append(_eeae, _gaaf)
		}
	}
	return _eeae
}

func (_afad *Context) capper() _ed.Capper {
	switch _afad._caa {
	case _adf.LineCapButt:
		return _ed.ButtCapper
	case _adf.LineCapRound:
		return _ed.RoundCapper
	case _adf.LineCapSquare:
		return _ed.SquareCapper
	}
	return nil
}

func (_cedf *Context) fill(_eca _ed.Painter) {
	_cbb := _cedf._df
	if _cedf._eae {
		_cbb = make(_ed.Path, len(_cedf._df))
		copy(_cbb, _cedf._df)
		_cbb.Add1(_edbe(_cedf._de))
	}
	_ccg := _cedf._ee
	_ccg.UseNonZeroWinding = _cedf._cdf == _adf.FillRuleWinding
	_ccg.Clear()
	_ccg.AddPath(_cbb)
	_ccg.Rasterize(_eca)
}

func NewContextForRGBA(im *_ad.RGBA) *Context {
	_eba := im.Bounds().Size().X
	_bbe := im.Bounds().Size().Y
	return &Context{_gee: _eba, _fdc: _bbe, _ee: _ed.NewRasterizer(_eba, _bbe), _gfd: im, _acc: _ce.Transparent, _adb: _gafe, _gbf: _ebb, _afe: 1, _cdf: _adf.FillRuleWinding, _bdf: _e.IdentityMatrix(), _acf: _adf.NewTextState()}
}

func NewContext(width, height int) *Context {
	return NewContextForRGBA(_ad.NewRGBA(_ad.Rect(0, 0, width, height)))
}

type linearGradient struct {
	_egf, _bdeb, _cedd, _fda float64
	_ddg                     stops
}

func (_ace *Context) Width() int { return _ace._gee }
func (_cbdd *Context) Clip()     { _cbdd.ClipPreserve(); _cbdd.ClearPath() }
func (_bfb *Context) Clear() {
	_ggb := _ad.NewUniform(_bfb._acc)
	_f.Draw(_bfb._gfd, _bfb._gfd.Bounds(), _ggb, _ad.Point{}, _f.Src)
}

func (_bda *Context) NewSubPath() {
	if _bda._eae {
		_bda._df.Add1(_edbe(_bda._de))
	}
	_bda._eae = false
}

func _dcc(_acab _ed.Path, _gea []float64, _cbddc float64) _ed.Path {
	return _gcff(_ddfe(_febf(_acab), _gea, _cbddc))
}

type stop struct {
	_aba  float64
	_dbac _ce.Color
}

func (_bg *Context) setFillAndStrokeColor(_bgd _ce.Color) {
	_bg._acc = _bgd
	_bg._adb = _ecg(_bgd)
	_bg._gbf = _ecg(_bgd)
}

func _beb(_acad _ad.Image) *_ad.RGBA {
	_bga := _acad.Bounds()
	_ffadc := _ad.NewRGBA(_bga)
	_g.Draw(_ffadc, _bga, _acad, _bga.Min, _g.Src)
	return _ffadc
}
func (_bfe *Context) Transform(x, y float64) (_adff, _daf float64) { return _bfe._bdf.Transform(x, y) }

var (
	_gafe = _ecg(_ce.White)
	_ebb  = _ecg(_ce.Black)
)

func (_fbab *Context) LineTo(x, y float64) {
	if !_fbab._eae {
		_fbab.MoveTo(x, y)
	} else {
		x, y = _fbab.Transform(x, y)
		_dbd := _e.NewPoint(x, y)
		_bge := _edbe(_dbd)
		_fbab._gg.Add1(_bge)
		_fbab._df.Add1(_bge)
		_fbab._dcf = _dbd
	}
}

type circle struct{ _aeg, _cdc, _ecd float64 }

func (_ega *Context) LineWidth() float64 { return _ega._afe }
func (_agc *surfacePattern) ColorAt(x, y int) _ce.Color {
	_egcc := _agc._ecfa.Bounds()
	switch _agc._efbf {
	case _ccd:
		if y >= _egcc.Dy() {
			return _ce.Transparent
		}
	case _ecde:
		if x >= _egcc.Dx() {
			return _ce.Transparent
		}
	case _afba:
		if x >= _egcc.Dx() || y >= _egcc.Dy() {
			return _ce.Transparent
		}
	}
	x = x%_egcc.Dx() + _egcc.Min.X
	y = y%_egcc.Dy() + _egcc.Min.Y
	return _agc._ecfa.At(x, y)
}

func (_afd *Context) SetMask(mask *_ad.Alpha) error {
	if mask.Bounds().Size() != _afd._gfd.Bounds().Size() {
		return _ca.New("\u006d\u0061\u0073\u006b\u0020\u0073i\u007a\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068 \u0063\u006f\u006e\u0074\u0065\u0078\u0074 \u0073\u0069\u007a\u0065")
	}
	_afd._fded = mask
	return nil
}

func (_feb *Context) RotateAbout(angle, x, y float64) {
	_feb.Translate(x, y)
	_feb.Rotate(angle)
	_feb.Translate(-x, -y)
}

func (_cgg *Context) FillPreserve() {
	var _bdb _ed.Painter
	if _cgg._fded == nil {
		if _eee, _fbg := _cgg._adb.(*solidPattern); _fbg {
			_deaf := _ed.NewRGBAPainter(_cgg._gfd)
			_deaf.SetColor(_eee._fcg)
			_bdb = _deaf
		}
	}
	if _bdb == nil {
		_bdb = _fcc(_cgg._gfd, _cgg._fded, _cgg._adb)
	}
	_cgg.fill(_bdb)
}
func (_dba *Context) Fill() { _dba.FillPreserve(); _dba.ClearPath() }
func (_edfgb *Context) DrawEllipse(x, y, rx, ry float64) {
	_edfgb.NewSubPath()
	_edfgb.DrawEllipticalArc(x, y, rx, ry, 0, 2*_d.Pi)
	_edfgb.ClosePath()
}

func _cedc(_abae, _dbe _ce.Color, _dbab float64) _ce.Color {
	_ggbe, _aff, _eceg, _fecc := _abae.RGBA()
	_cfe, _fead, _fbc, _gcc := _dbe.RGBA()
	return _ce.RGBA{_ebafd(_ggbe, _cfe, _dbab), _ebafd(_aff, _fead, _dbab), _ebafd(_eceg, _fbc, _dbab), _ebafd(_fecc, _gcc, _dbab)}
}
func (_eaef *Context) Identity() { _eaef._bdf = _e.IdentityMatrix() }
func (_dafg *Context) Pop() {
	_efb := *_dafg
	_bfg := _dafg._ddef
	_ddf := _bfg[len(_bfg)-1]
	*_dafg = *_ddf
	_dafg._gg = _efb._gg
	_dafg._df = _efb._df
	_dafg._de = _efb._de
	_dafg._dcf = _efb._dcf
	_dafg._eae = _efb._eae
}
func NewContextForImage(im _ad.Image) *Context { return NewContextForRGBA(_beb(im)) }
func _baee(_bbf float64) float64               { return _bbf * _d.Pi / 180 }
func (_cedg *Context) DrawStringAnchored(s string, face _ga.Face, x, y, ax, ay float64) {
	_acae, _bgg := _cedg.MeasureString(s, face)
	_cedg.drawString(s, face, x-ax*_acae, y+ay*_bgg)
}
func (_gfbg *Context) Translate(x, y float64) { _gfbg._bdf = _gfbg._bdf.Translate(x, y) }
func _fcc(_fbac *_ad.RGBA, _ded *_ad.Alpha, _dfcg _adf.Pattern) *patternPainter {
	return &patternPainter{_fbac, _ded, _dfcg}
}

func _gbeb(_fgcc _gf.Int26_6) float64 {
	const _deg, _bag = 6, 1<<6 - 1
	if _fgcc >= 0 {
		return float64(_fgcc>>_deg) + float64(_fgcc&_bag)/64
	}
	_fgcc = -_fgcc
	if _fgcc >= 0 {
		return -(float64(_fgcc>>_deg) + float64(_fgcc&_bag)/64)
	}
	return 0
}

func (_ccc *Context) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64) {
	const _gcfe = 16
	for _gce := 0; _gce < _gcfe; _gce++ {
		_bea := float64(_gce+0) / _gcfe
		_cdfe := float64(_gce+1) / _gcfe
		_egd := angle1 + (angle2-angle1)*_bea
		_faf := angle1 + (angle2-angle1)*_cdfe
		_cfc := x + rx*_d.Cos(_egd)
		_aca := y + ry*_d.Sin(_egd)
		_gff := x + rx*_d.Cos((_egd+_faf)/2)
		_ddac := y + ry*_d.Sin((_egd+_faf)/2)
		_ebc := x + rx*_d.Cos(_faf)
		_deb := y + ry*_d.Sin(_faf)
		_edc := 2*_gff - _cfc/2 - _ebc/2
		_edb := 2*_ddac - _aca/2 - _deb/2
		if _gce == 0 {
			if _ccc._eae {
				_ccc.LineTo(_cfc, _aca)
			} else {
				_ccc.MoveTo(_cfc, _aca)
			}
		}
		_ccc.QuadraticTo(_edc, _edb, _ebc, _deb)
	}
}

func (_ecc *Context) QuadraticTo(x1, y1, x2, y2 float64) {
	if !_ecc._eae {
		_ecc.MoveTo(x1, y1)
	}
	x1, y1 = _ecc.Transform(x1, y1)
	x2, y2 = _ecc.Transform(x2, y2)
	_gfc := _e.NewPoint(x1, y1)
	_aed := _e.NewPoint(x2, y2)
	_cdb := _edbe(_gfc)
	_bdfd := _edbe(_aed)
	_ecc._gg.Add2(_cdb, _bdfd)
	_ecc._df.Add2(_cdb, _bdfd)
	_ecc._dcf = _aed
}

func (_eag *Context) MoveTo(x, y float64) {
	if _eag._eae {
		_eag._df.Add1(_edbe(_eag._de))
	}
	x, y = _eag.Transform(x, y)
	_geb := _e.NewPoint(x, y)
	_efa := _edbe(_geb)
	_eag._gg.Start(_efa)
	_eag._df.Start(_efa)
	_eag._de = _geb
	_eag._dcf = _geb
	_eag._eae = true
}

func (_aedb *Context) ClosePath() {
	if _aedb._eae {
		_gge := _edbe(_aedb._de)
		_aedb._gg.Add1(_gge)
		_aedb._df.Add1(_gge)
		_aedb._dcf = _aedb._de
	}
}

type Context struct {
	_gee  int
	_fdc  int
	_ee   *_ed.Rasterizer
	_gfd  *_ad.RGBA
	_fded *_ad.Alpha
	_acc  _ce.Color
	_adb  _adf.Pattern
	_gbf  _adf.Pattern
	_gg   _ed.Path
	_df   _ed.Path
	_de   _e.Point
	_dcf  _e.Point
	_eae  bool
	_cef  []float64
	_aa   float64
	_afe  float64
	_caa  _adf.LineCap
	_bbd  _adf.LineJoin
	_cdf  _adf.FillRule
	_bdf  _e.Matrix
	_acf  _adf.TextState
	_ddef []*Context
}

func (_acced *Context) Matrix() _e.Matrix { return _acced._bdf }
