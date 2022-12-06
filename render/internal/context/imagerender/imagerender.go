package imagerender

import (
	_ag "errors"
	_ba "fmt"
	_e "image"
	_da "image/color"
	_f "image/draw"
	_d "math"
	_ae "sort"
	_a "strings"

	_ab "bitbucket.org/shenghui0779/gopdf/common"
	_c "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_af "bitbucket.org/shenghui0779/gopdf/render/internal/context"
	_bc "github.com/unidoc/freetype/raster"
	_bae "golang.org/x/image/draw"
	_bb "golang.org/x/image/font"
	_bg "golang.org/x/image/math/f64"
	_ce "golang.org/x/image/math/fixed"
)

func (_abg *Context) QuadraticTo(x1, y1, x2, y2 float64) {
	if !_abg._gff {
		_abg.MoveTo(x1, y1)
	}
	x1, y1 = _abg.Transform(x1, y1)
	x2, y2 = _abg.Transform(x2, y2)
	_gad := _c.NewPoint(x1, y1)
	_ffc := _c.NewPoint(x2, y2)
	_cgab := _adbe(_gad)
	_afe := _adbe(_ffc)
	_abg._gge.Add2(_cgab, _afe)
	_abg._feeb.Add2(_cgab, _afe)
	_abg._gf = _ffc
}
func _ed(_ef, _eda, _gc, _baf, _cae, _eec, _afd, _ec, _ga float64) (_cg, _cga float64) {
	_cbfa := 1 - _ga
	_bad := _cbfa * _cbfa * _cbfa
	_ebg := 3 * _cbfa * _cbfa * _ga
	_bcf := 3 * _cbfa * _ga * _ga
	_dgd := _ga * _ga * _ga
	_cg = _bad*_ef + _ebg*_gc + _bcf*_cae + _dgd*_afd
	_cga = _bad*_eda + _ebg*_baf + _bcf*_eec + _dgd*_ec
	return
}
func (_bdd stops) Len() int { return len(_bdd) }

type stop struct {
	_eggb float64
	_cdfg _da.Color
}

func (_ged *Context) fill(_bgg _bc.Painter) {
	_bdc := _ged._feeb
	if _ged._gff {
		_bdc = make(_bc.Path, len(_ged._feeb))
		copy(_bdc, _ged._feeb)
		_bdc.Add1(_adbe(_ged._cef))
	}
	_dee := _ged._def
	_dee.UseNonZeroWinding = _ged._dgc == _af.FillRuleWinding
	_dee.Clear()
	_dee.AddPath(_bdc)
	_dee.Rasterize(_bgg)
}
func (_fde *Context) SetFillRGBA(r, g, b, a float64) {
	_gac := _da.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_fde._db = _gac
	_fde._gagc = _fdaf(_gac)
}
func (_eag *Context) LineWidth() float64 { return _eag._dae }
func (_cgad *Context) SetRGBA(r, g, b, a float64) {
	_cgad._db = _da.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_cgad.setFillAndStrokeColor(_cgad._db)
}
func (_dge *Context) DrawRectangle(x, y, w, h float64) {
	_dge.NewSubPath()
	_dge.MoveTo(x, y)
	_dge.LineTo(x+w, y)
	_dge.LineTo(x+w, y+h)
	_dge.LineTo(x, y+h)
	_dge.ClosePath()
}
func _gdddd(_efcd *_e.RGBA, _bce *_e.Alpha, _dff _af.Pattern) *patternPainter {
	return &patternPainter{_efcd, _bce, _dff}
}
func _eea(_feg _bc.Path) [][]_c.Point {
	var _aec [][]_c.Point
	var _aga []_c.Point
	var _cbfe, _cbff float64
	for _gdadd := 0; _gdadd < len(_feg); {
		switch _feg[_gdadd] {
		case 0:
			if len(_aga) > 0 {
				_aec = append(_aec, _aga)
				_aga = nil
			}
			_agd := _fgce(_feg[_gdadd+1])
			_fbcb := _fgce(_feg[_gdadd+2])
			_aga = append(_aga, _c.NewPoint(_agd, _fbcb))
			_cbfe, _cbff = _agd, _fbcb
			_gdadd += 4
		case 1:
			_cbgg := _fgce(_feg[_gdadd+1])
			_beg := _fgce(_feg[_gdadd+2])
			_aga = append(_aga, _c.NewPoint(_cbgg, _beg))
			_cbfe, _cbff = _cbgg, _beg
			_gdadd += 4
		case 2:
			_gga := _fgce(_feg[_gdadd+1])
			_bda := _fgce(_feg[_gdadd+2])
			_efa := _fgce(_feg[_gdadd+3])
			_aca := _fgce(_feg[_gdadd+4])
			_dbee := _dc(_cbfe, _cbff, _gga, _bda, _efa, _aca)
			_aga = append(_aga, _dbee...)
			_cbfe, _cbff = _efa, _aca
			_gdadd += 6
		case 3:
			_dage := _fgce(_feg[_gdadd+1])
			_ccfg := _fgce(_feg[_gdadd+2])
			_cdfa := _fgce(_feg[_gdadd+3])
			_dcg := _fgce(_feg[_gdadd+4])
			_cacc := _fgce(_feg[_gdadd+5])
			_ecag := _fgce(_feg[_gdadd+6])
			_abfbf := _bd(_cbfe, _cbff, _dage, _ccfg, _cdfa, _dcg, _cacc, _ecag)
			_aga = append(_aga, _abfbf...)
			_cbfe, _cbff = _cacc, _ecag
			_gdadd += 8
		default:
			_ab.Log.Debug("\u0057\u0041\u0052\u004e: \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0074\u0068\u003a\u0020%\u0076", _feg)
			return _aec
		}
	}
	if len(_aga) > 0 {
		_aec = append(_aec, _aga)
	}
	return _aec
}
func (_fccd *surfacePattern) ColorAt(x, y int) _da.Color {
	_ggba := _fccd._adga.Bounds()
	switch _fccd._begg {
	case _cadg:
		if y >= _ggba.Dy() {
			return _da.Transparent
		}
	case _cdcd:
		if x >= _ggba.Dx() {
			return _da.Transparent
		}
	case _fga:
		if x >= _ggba.Dx() || y >= _ggba.Dy() {
			return _da.Transparent
		}
	}
	x = x%_ggba.Dx() + _ggba.Min.X
	y = y%_ggba.Dy() + _ggba.Min.Y
	return _fccd._adga.At(x, y)
}
func NewContextForImage(im _e.Image) *Context { return NewContextForRGBA(_gaca(im)) }
func (_cead *Context) capper() _bc.Capper {
	switch _cead._afc {
	case _af.LineCapButt:
		return _bc.ButtCapper
	case _af.LineCapRound:
		return _bc.RoundCapper
	case _af.LineCapSquare:
		return _bc.SquareCapper
	}
	return nil
}
func _bec(_adf float64, _fad stops) _da.Color {
	if _adf <= 0.0 || len(_fad) == 1 {
		return _fad[0]._cdfg
	}
	_aaee := _fad[len(_fad)-1]
	if _adf >= _aaee._eggb {
		return _aaee._cdfg
	}
	for _gfb, _dfe := range _fad[1:] {
		if _adf < _dfe._eggb {
			_adf = (_adf - _fad[_gfb]._eggb) / (_dfe._eggb - _fad[_gfb]._eggb)
			return _fdfg(_fad[_gfb]._cdfg, _dfe._cdfg, _adf)
		}
	}
	return _aaee._cdfg
}
func _ebad(_dddd, _ffg uint32, _abac float64) uint8 {
	return uint8(int32(float64(_dddd)*(1.0-_abac)+float64(_ffg)*_abac) >> 8)
}
func (_cad *Context) SetRGB(r, g, b float64) { _cad.SetRGBA(r, g, b, 1) }
func _fefb(_bdbg [][]_c.Point) _bc.Path {
	var _gadg _bc.Path
	for _, _dcc := range _bdbg {
		var _dadgc _ce.Point26_6
		for _dcgf, _bged := range _dcc {
			_eeg := _adbe(_bged)
			if _dcgf == 0 {
				_gadg.Start(_eeg)
			} else {
				_efab := _eeg.X - _dadgc.X
				_geda := _eeg.Y - _dadgc.Y
				if _efab < 0 {
					_efab = -_efab
				}
				if _geda < 0 {
					_geda = -_geda
				}
				if _efab+_geda > 8 {
					_gadg.Add1(_eeg)
				}
			}
			_dadgc = _eeg
		}
	}
	return _gadg
}
func (_eca *Context) FillPreserve() {
	var _cbd _bc.Painter
	if _eca._fdd == nil {
		if _gdd, _cecg := _eca._gagc.(*solidPattern); _cecg {
			_afdf := _bc.NewRGBAPainter(_eca._bga)
			_afdf.SetColor(_gdd._bebf)
			_cbd = _afdf
		}
	}
	if _cbd == nil {
		_cbd = _gdddd(_eca._bga, _eca._fdd, _eca._gagc)
	}
	_eca.fill(_cbd)
}
func _edab(_cgbd, _ccb, _baad, _fef float64) _af.Gradient {
	_adg := &linearGradient{_ddd: _cgbd, _cac: _ccb, _fgc: _baad, _dafc: _fef}
	return _adg
}
func (_dab *Context) FillPattern() _af.Pattern { return _dab._gagc }
func (_cc *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	if !_cc._gff {
		_cc.MoveTo(x1, y1)
	}
	_eae, _gcg := _cc._gf.X, _cc._gf.Y
	x1, y1 = _cc.Transform(x1, y1)
	x2, y2 = _cc.Transform(x2, y2)
	x3, y3 = _cc.Transform(x3, y3)
	_eg := _bd(_eae, _gcg, x1, y1, x2, y2, x3, y3)
	_aaf := _adbe(_cc._gf)
	for _, _agf := range _eg[1:] {
		_edc := _adbe(_agf)
		if _edc == _aaf {
			continue
		}
		_aaf = _edc
		_cc._gge.Add1(_edc)
		_cc._feeb.Add1(_edc)
		_cc._gf = _agf
	}
}
func (_gcc *Context) InvertMask() {
	if _gcc._fdd == nil {
		_gcc._fdd = _e.NewAlpha(_gcc._bga.Bounds())
	} else {
		for _eaeb, _dea := range _gcc._fdd.Pix {
			_gcc._fdd.Pix[_eaeb] = 255 - _dea
		}
	}
}
func (_cea *Context) SetStrokeRGBA(r, g, b, a float64) {
	_bbbg := _da.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), uint8(a * 255)}
	_cea._cgb = _fdaf(_bbbg)
}
func _bdcf(_dabf, _eddc, _fbe, _gbdgb, _gbbc, _dgbc float64) _af.Gradient {
	_gdad := circle{_dabf, _eddc, _fbe}
	_abaa := circle{_gbdgb, _gbbc, _dgbc}
	_baag := circle{_gbdgb - _dabf, _gbbc - _eddc, _dgbc - _fbe}
	_fdba := _fca(_baag._ddc, _baag._fba, -_baag._gedc, _baag._ddc, _baag._fba, _baag._gedc)
	var _cbcfb float64
	if _fdba != 0 {
		_cbcfb = 1.0 / _fdba
	}
	_aded := -_gdad._gedc
	_bed := &radialGradient{_fgb: _gdad, _fff: _abaa, _dagg: _baag, _dadg: _fdba, _ebe: _cbcfb, _afcb: _aded}
	return _bed
}
func (_bdca *Context) drawRegularPolygon(_bbbb int, _fbg, _cge, _ggc, _gfa float64) {
	_abgc := 2 * _d.Pi / float64(_bbbb)
	_gfa -= _d.Pi / 2
	if _bbbb%2 == 0 {
		_gfa += _abgc / 2
	}
	_bdca.NewSubPath()
	for _ffe := 0; _ffe < _bbbb; _ffe++ {
		_eaf := _gfa + _abgc*float64(_ffe)
		_bdca.LineTo(_fbg+_ggc*_d.Cos(_eaf), _cge+_ggc*_d.Sin(_eaf))
	}
	_bdca.ClosePath()
}
func (_gbdf *Context) Translate(x, y float64) { _gbdf._baab = _gbdf._baab.Translate(x, y) }
func (_cggb *Context) SetRGBA255(r, g, b, a int) {
	_cggb._db = _da.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	_cggb.setFillAndStrokeColor(_cggb._db)
}

type stops []stop

func _dfc(_gegf _bc.Path, _aabe []float64, _gdaddg float64) _bc.Path {
	return _fefb(_gggd(_eea(_gegf), _aabe, _gdaddg))
}
func (_dad *linearGradient) AddColorStop(offset float64, color _da.Color) {
	_dad._ccd = append(_dad._ccd, stop{_eggb: offset, _cdfg: color})
	_ae.Sort(_dad._ccd)
}
func (_bcba *Context) Transform(x, y float64) (_aeb, _gbc float64) {
	return _bcba._baab.Transform(x, y)
}
func (_dbb *Context) SetMask(mask *_e.Alpha) error {
	if mask.Bounds().Size() != _dbb._bga.Bounds().Size() {
		return _ag.New("\u006d\u0061\u0073\u006b\u0020\u0073i\u007a\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068 \u0063\u006f\u006e\u0074\u0065\u0078\u0074 \u0073\u0069\u007a\u0065")
	}
	_dbb._fdd = mask
	return nil
}
func (_aagf *Context) ScaleAbout(sx, sy, x, y float64) {
	_aagf.Translate(x, y)
	_aagf.Scale(sx, sy)
	_aagf.Translate(-x, -y)
}
func (_fcc *Context) Identity() { _fcc._baab = _c.IdentityMatrix() }

const (
	_gaab repeatOp = iota
	_cadg
	_cdcd
	_fga
)

type Context struct {
	_bcd  int
	_abb  int
	_def  *_bc.Rasterizer
	_bga  *_e.RGBA
	_fdd  *_e.Alpha
	_db   _da.Color
	_gagc _af.Pattern
	_cgb  _af.Pattern
	_gge  _bc.Path
	_feeb _bc.Path
	_cef  _c.Point
	_gf   _c.Point
	_gff  bool
	_fdb  []float64
	_bcg  float64
	_dae  float64
	_afc  _af.LineCap
	_dec  _af.LineJoin
	_dgc  _af.FillRule
	_baab _c.Matrix
	_gfg  _af.TextState
	_dcf  []*Context
}

func (_abdf *Context) joiner() _bc.Joiner {
	switch _abdf._dec {
	case _af.LineJoinBevel:
		return _bc.BevelJoiner
	case _af.LineJoinRound:
		return _bc.RoundJoiner
	}
	return nil
}
func (_efc *solidPattern) ColorAt(x, y int) _da.Color { return _efc._bebf }
func (_gbd *Context) MoveTo(x, y float64) {
	if _gbd._gff {
		_gbd._feeb.Add1(_adbe(_gbd._cef))
	}
	x, y = _gbd.Transform(x, y)
	_ecc := _c.NewPoint(x, y)
	_abbd := _adbe(_ecc)
	_gbd._gge.Start(_abbd)
	_gbd._feeb.Start(_abbd)
	_gbd._cef = _ecc
	_gbd._gf = _ecc
	_gbd._gff = true
}

var (
	_aff = _fdaf(_da.White)
	_cfa = _fdaf(_da.Black)
)

func _adbe(_fgbg _c.Point) _ce.Point26_6 { return _ce.Point26_6{X: _agfe(_fgbg.X), Y: _agfe(_fgbg.Y)} }
func (_ade *radialGradient) AddColorStop(offset float64, color _da.Color) {
	_ade._bbaf = append(_ade._bbaf, stop{_eggb: offset, _cdfg: color})
	_ae.Sort(_ade._bbaf)
}
func _faeg(_eafd string) (_cdgd, _ccfgc, _becg, _abec int) {
	_eafd = _a.TrimPrefix(_eafd, "\u0023")
	_abec = 255
	if len(_eafd) == 3 {
		_bfb := "\u00251\u0078\u0025\u0031\u0078\u0025\u0031x"
		_ba.Sscanf(_eafd, _bfb, &_cdgd, &_ccfgc, &_becg)
		_cdgd |= _cdgd << 4
		_ccfgc |= _ccfgc << 4
		_becg |= _becg << 4
	}
	if len(_eafd) == 6 {
		_afee := "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_ba.Sscanf(_eafd, _afee, &_cdgd, &_ccfgc, &_becg)
	}
	if len(_eafd) == 8 {
		_eeca := "\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078"
		_ba.Sscanf(_eafd, _eeca, &_cdgd, &_ccfgc, &_becg, &_abec)
	}
	return
}
func (_dcfb stops) Swap(i, j int) { _dcfb[i], _dcfb[j] = _dcfb[j], _dcfb[i] }
func (_bgb *Context) ResetClip()  { _bgb._fdd = nil }
func (_bdb *Context) AsMask() *_e.Alpha {
	_dfa := _e.NewAlpha(_bdb._bga.Bounds())
	_bae.Draw(_dfa, _bdb._bga.Bounds(), _bdb._bga, _e.Point{}, _bae.Src)
	return _dfa
}
func (_afdb *Context) SetColor(c _da.Color) { _afdb.setFillAndStrokeColor(c) }
func _fdfg(_bggc, _eff _da.Color, _degc float64) _da.Color {
	_ebb, _ebbc, _bcga, _bcgaa := _bggc.RGBA()
	_eac, _agff, _ecacf, _fbed := _eff.RGBA()
	return _da.RGBA{_ebad(_ebb, _eac, _degc), _ebad(_ebbc, _agff, _degc), _ebad(_bcga, _ecacf, _degc), _ebad(_bcgaa, _fbed, _degc)}
}
func _agfe(_dgef float64) _ce.Int26_6 { return _ce.Int26_6(_dgef * 64) }
func (_ebfa *Context) DrawCircle(x, y, r float64) {
	_ebfa.NewSubPath()
	_ebfa.DrawEllipticalArc(x, y, r, r, 0, 2*_d.Pi)
	_ebfa.ClosePath()
}
func (_gae *Context) ClipPreserve() {
	_ccf := _e.NewAlpha(_e.Rect(0, 0, _gae._bcd, _gae._abb))
	_bfa := _bc.NewAlphaOverPainter(_ccf)
	_gae.fill(_bfa)
	if _gae._fdd == nil {
		_gae._fdd = _ccf
	} else {
		_gcb := _e.NewAlpha(_e.Rect(0, 0, _gae._bcd, _gae._abb))
		_bae.DrawMask(_gcb, _gcb.Bounds(), _ccf, _e.Point{}, _gae._fdd, _e.Point{}, _bae.Over)
		_gae._fdd = _gcb
	}
}
func (_gfga *Context) ClearPath() { _gfga._gge.Clear(); _gfga._feeb.Clear(); _gfga._gff = false }
func (_dfg *Context) StrokePreserve() {
	var _bcb _bc.Painter
	if _dfg._fdd == nil {
		if _ad, _egc := _dfg._cgb.(*solidPattern); _egc {
			_cfc := _bc.NewRGBAPainter(_dfg._bga)
			_cfc.SetColor(_ad._bebf)
			_bcb = _cfc
		}
	}
	if _bcb == nil {
		_bcb = _gdddd(_dfg._bga, _dfg._fdd, _dfg._cgb)
	}
	_dfg.stroke(_bcb)
}
func NewContextForRGBA(im *_e.RGBA) *Context {
	_dbc := im.Bounds().Size().X
	_gab := im.Bounds().Size().Y
	return &Context{_bcd: _dbc, _abb: _gab, _def: _bc.NewRasterizer(_dbc, _gab), _bga: im, _db: _da.Transparent, _gagc: _aff, _cgb: _cfa, _dae: 1, _dgc: _af.FillRuleWinding, _baab: _c.IdentityMatrix(), _gfg: _af.NewTextState()}
}
func (_aba *Context) Stroke() { _aba.StrokePreserve(); _aba.ClearPath() }
func (_gddd *Context) drawString(_cce string, _edb _bb.Face, _ffed, _ffa float64) {
	_bccb := &_bb.Drawer{Src: _e.NewUniform(_gddd._db), Face: _edb, Dot: _adbe(_c.NewPoint(_ffed, _ffa))}
	_dfag := rune(-1)
	for _, _dbe := range _cce {
		if _dfag >= 0 {
			_bccb.Dot.X += _bccb.Face.Kern(_dfag, _dbe)
		}
		_ccc, _dgb, _gead, _ecfc, _eba := _bccb.Face.Glyph(_bccb.Dot, _dbe)
		if !_eba {
			continue
		}
		_geed := _ccc.Sub(_ccc.Min)
		_bgd := _e.NewRGBA(_geed)
		_bae.DrawMask(_bgd, _geed, _bccb.Src, _e.Point{}, _dgb, _gead, _bae.Over)
		var _abe *_bae.Options
		if _gddd._fdd != nil {
			_abe = &_bae.Options{DstMask: _gddd._fdd, DstMaskP: _e.Point{}}
		}
		_ecg := _gddd._baab.Clone().Translate(float64(_ccc.Min.X), float64(_ccc.Min.Y))
		_dda := _bg.Aff3{_ecg[0], _ecg[3], _ecg[6], _ecg[1], _ecg[4], _ecg[7]}
		_bae.BiLinear.Transform(_gddd._bga, _dda, _bgd, _geed, _bae.Over, _abe)
		_bccb.Dot.X += _ecfc
		_dfag = _dbe
	}
}
func (_ebf *Context) SetLineCap(lineCap _af.LineCap)     { _ebf._afc = lineCap }
func (_gd *Context) Width() int                          { return _gd._bcd }
func (_ffbe *Context) SetLineJoin(lineJoin _af.LineJoin) { _ffbe._dec = lineJoin }
func (_gbf *Context) Push()                              { _beb := *_gbf; _gbf._dcf = append(_gbf._dcf, &_beb) }

type surfacePattern struct {
	_adga _e.Image
	_begg repeatOp
}

func (_egb *Context) Matrix() _c.Matrix     { return _egb._baab }
func (_dgg *Context) SetRGB255(r, g, b int) { _dgg.SetRGBA255(r, g, b, 255) }
func (_fdde *Context) RotateAbout(angle, x, y float64) {
	_fdde.Translate(x, y)
	_fdde.Rotate(angle)
	_fdde.Translate(-x, -y)
}
func _fca(_bcbg, _eafb, _beag, _bde, _gbg, _gaa float64) float64 {
	return _bcbg*_bde + _eafb*_gbg + _beag*_gaa
}
func _dc(_bbb, _ff, _fb, _de, _cag, _aac float64) []_c.Point {
	_fc := (_d.Hypot(_fb-_bbb, _de-_ff) + _d.Hypot(_cag-_fb, _aac-_de))
	_eb := int(_fc + 0.5)
	if _eb < 4 {
		_eb = 4
	}
	_fcg := float64(_eb) - 1
	_g := make([]_c.Point, _eb)
	for _gg := 0; _gg < _eb; _gg++ {
		_fee := float64(_gg) / _fcg
		_cbf, _ee := _cb(_bbb, _ff, _fb, _de, _cag, _aac, _fee)
		_g[_gg] = _c.NewPoint(_cbf, _ee)
	}
	return _g
}
func (_gaf *Context) Clear() {
	_badb := _e.NewUniform(_gaf._db)
	_bae.Draw(_gaf._bga, _gaf._bga.Bounds(), _badb, _e.Point{}, _bae.Src)
}
func (_afg *Context) SetFillStyle(pattern _af.Pattern) {
	if _bbd, _bgc := pattern.(*solidPattern); _bgc {
		_afg._db = _bbd._bebf
	}
	_afg._gagc = pattern
}
func _fdaf(_baba _da.Color) _af.Pattern { return &solidPattern{_bebf: _baba} }
func (_acc *Context) stroke(_gacg _bc.Painter) {
	_bfd := _acc._gge
	if len(_acc._fdb) > 0 {
		_bfd = _dfc(_bfd, _acc._fdb, _acc._bcg)
	} else {
		_bfd = _fefb(_eea(_bfd))
	}
	_bea := _acc._def
	_bea.UseNonZeroWinding = true
	_bea.Clear()
	_cefd := (_acc._baab.ScalingFactorX() + _acc._baab.ScalingFactorY()) / 2
	_bea.AddStroke(_bfd, _agfe(_acc._dae*_cefd), _acc.capper(), _acc.joiner())
	_bea.Rasterize(_gacg)
}
func (_cfg *Context) SetDash(dashes ...float64) { _cfg._fdb = dashes }
func (_ceb *Context) Shear(x, y float64)        { _ceb._baab.Shear(x, y) }

type patternPainter struct {
	_bgdc *_e.RGBA
	_cgf  *_e.Alpha
	_ega  _af.Pattern
}

func (_decg *Context) DrawRoundedRectangle(x, y, w, h, r float64) {
	_gadf, _gda, _gcga, _fdf := x, x+r, x+w-r, x+w
	_bagd, _add, _acg, _gegc := y, y+r, y+h-r, y+h
	_decg.NewSubPath()
	_decg.MoveTo(_gda, _bagd)
	_decg.LineTo(_gcga, _bagd)
	_decg.DrawArc(_gcga, _add, r, _eadg(270), _eadg(360))
	_decg.LineTo(_fdf, _acg)
	_decg.DrawArc(_gcga, _acg, r, _eadg(0), _eadg(90))
	_decg.LineTo(_gda, _gegc)
	_decg.DrawArc(_gda, _acg, r, _eadg(90), _eadg(180))
	_decg.LineTo(_gadf, _add)
	_decg.DrawArc(_gda, _add, r, _eadg(180), _eadg(270))
	_decg.ClosePath()
}
func _gaca(_cccg _e.Image) *_e.RGBA {
	_ggfb := _cccg.Bounds()
	_fcbf := _e.NewRGBA(_ggfb)
	_f.Draw(_fcbf, _ggfb, _cccg, _ggfb.Min, _f.Src)
	return _fcbf
}
func (_gee *Context) TextState() *_af.TextState { return &_gee._gfg }
func (_bbdd stops) Less(i, j int) bool          { return _bbdd[i]._eggb < _bbdd[j]._eggb }
func (_adc *Context) ShearAbout(sx, sy, x, y float64) {
	_adc.Translate(x, y)
	_adc.Shear(sx, sy)
	_adc.Translate(-x, -y)
}
func NewContext(width, height int) *Context {
	return NewContextForRGBA(_e.NewRGBA(_e.Rect(0, 0, width, height)))
}
func _bd(_abd, _ge, _ecf, _cgg, _fd, _ac, _abf, _gb float64) []_c.Point {
	_cbc := (_d.Hypot(_ecf-_abd, _cgg-_ge) + _d.Hypot(_fd-_ecf, _ac-_cgg) + _d.Hypot(_abf-_fd, _gb-_ac))
	_ffb := int(_cbc + 0.5)
	if _ffb < 4 {
		_ffb = 4
	}
	_deg := float64(_ffb) - 1
	_baa := make([]_c.Point, _ffb)
	for _edg := 0; _edg < _ffb; _edg++ {
		_gag := float64(_edg) / _deg
		_cbg, _ea := _ed(_abd, _ge, _ecf, _cgg, _fd, _ac, _abf, _gb, _gag)
		_baa[_edg] = _c.NewPoint(_cbg, _ea)
	}
	return _baa
}
func _abad(_ccg _e.Image, _faea repeatOp) _af.Pattern {
	return &surfacePattern{_adga: _ccg, _begg: _faea}
}
func (_ecb *Context) SetHexColor(x string) {
	_gbe, _geg, _bbc, _ggb := _faeg(x)
	_ecb.SetRGBA255(_gbe, _geg, _bbc, _ggb)
}
func _gggd(_aegcc [][]_c.Point, _dgee []float64, _bdf float64) [][]_c.Point {
	var _abfbd [][]_c.Point
	if len(_dgee) == 0 {
		return _aegcc
	}
	if len(_dgee) == 1 {
		_dgee = append(_dgee, _dgee[0])
	}
	for _, _gega := range _aegcc {
		if len(_gega) < 2 {
			continue
		}
		_fbad := _gega[0]
		_egd := 1
		_cgag := 0
		_bbg := 0.0
		if _bdf != 0 {
			var _dadf float64
			for _, _bdaa := range _dgee {
				_dadf += _bdaa
			}
			_bdf = _d.Mod(_bdf, _dadf)
			if _bdf < 0 {
				_bdf += _dadf
			}
			for _aebc, _aee := range _dgee {
				_bdf -= _aee
				if _bdf < 0 {
					_cgag = _aebc
					_bbg = _aee + _bdf
					break
				}
			}
		}
		var _cage []_c.Point
		_cage = append(_cage, _fbad)
		for _egd < len(_gega) {
			_ace := _dgee[_cgag]
			_dbec := _gega[_egd]
			_daff := _fbad.Distance(_dbec)
			_gba := _ace - _bbg
			if _daff > _gba {
				_bee := _gba / _daff
				_cbcfe := _fbad.Interpolate(_dbec, _bee)
				_cage = append(_cage, _cbcfe)
				if _cgag%2 == 0 && len(_cage) > 1 {
					_abfbd = append(_abfbd, _cage)
				}
				_cage = nil
				_cage = append(_cage, _cbcfe)
				_bbg = 0
				_fbad = _cbcfe
				_cgag = (_cgag + 1) % len(_dgee)
			} else {
				_cage = append(_cage, _dbec)
				_fbad = _dbec
				_bbg += _daff
				_egd++
			}
		}
		if _cgag%2 == 0 && len(_cage) > 1 {
			_abfbd = append(_abfbd, _cage)
		}
	}
	return _abfbd
}

type repeatOp int

func (_baaf *Context) Pop() {
	_ecbd := *_baaf
	_dce := _baaf._dcf
	_cdg := _dce[len(_dce)-1]
	*_baaf = *_cdg
	_baaf._gge = _ecbd._gge
	_baaf._feeb = _ecbd._feeb
	_baaf._cef = _ecbd._cef
	_baaf._gf = _ecbd._gf
	_baaf._gff = _ecbd._gff
}
func (_abfb *Context) SetLineWidth(lineWidth float64)  { _abfb._dae = lineWidth }
func _eadg(_afeg float64) float64                      { return _afeg * _d.Pi / 180 }
func (_gbbd *Context) DrawImage(im _e.Image, x, y int) { _gbbd.DrawImageAnchored(im, x, y, 0, 0) }
func (_bab *Context) Rotate(angle float64)             { _bab._baab = _bab._baab.Rotate(angle) }
func (_baca *Context) LineTo(x, y float64) {
	if !_baca._gff {
		_baca.MoveTo(x, y)
	} else {
		x, y = _baca.Transform(x, y)
		_ecfd := _c.NewPoint(x, y)
		_dcd := _adbe(_ecfd)
		_baca._gge.Add1(_dcd)
		_baca._feeb.Add1(_dcd)
		_baca._gf = _ecfd
	}
}
func (_egg *Context) DrawLine(x1, y1, x2, y2 float64) { _egg.MoveTo(x1, y1); _egg.LineTo(x2, y2) }
func (_aagd *Context) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64) {
	const _eef = 16
	for _cd := 0; _cd < _eef; _cd++ {
		_aad := float64(_cd+0) / _eef
		_efb := float64(_cd+1) / _eef
		_gce := angle1 + (angle2-angle1)*_aad
		_caf := angle1 + (angle2-angle1)*_efb
		_cbca := x + rx*_d.Cos(_gce)
		_cdc := y + ry*_d.Sin(_gce)
		_fbc := x + rx*_d.Cos((_gce+_caf)/2)
		_eadb := y + ry*_d.Sin((_gce+_caf)/2)
		_ecac := x + rx*_d.Cos(_caf)
		_dba := y + ry*_d.Sin(_caf)
		_ede := 2*_fbc - _cbca/2 - _ecac/2
		_eccf := 2*_eadb - _cdc/2 - _dba/2
		if _cd == 0 {
			if _aagd._gff {
				_aagd.LineTo(_cbca, _cdc)
			} else {
				_aagd.MoveTo(_cbca, _cdc)
			}
		}
		_aagd.QuadraticTo(_ede, _eccf, _ecac, _dba)
	}
}
func _fgce(_cgc _ce.Int26_6) float64 {
	const _gcee, _fgbf = 6, 1<<6 - 1
	if _cgc >= 0 {
		return float64(_cgc>>_gcee) + float64(_cgc&_fgbf)/64
	}
	_cgc = -_cgc
	if _cgc >= 0 {
		return -(float64(_cgc>>_gcee) + float64(_cgc&_fgbf)/64)
	}
	return 0
}
func (_aae *Context) DrawImageAnchored(im _e.Image, x, y int, ax, ay float64) {
	_gdec := im.Bounds().Size()
	x -= int(ax * float64(_gdec.X))
	y -= int(ay * float64(_gdec.Y))
	_aab := _bae.BiLinear
	_gadfe := _aae._baab.Clone().Translate(float64(x), float64(y))
	_cdd := _bg.Aff3{_gadfe[0], _gadfe[3], _gadfe[6], _gadfe[1], _gadfe[4], _gadfe[7]}
	if _aae._fdd == nil {
		_aab.Transform(_aae._bga, _cdd, im, im.Bounds(), _bae.Over, nil)
	} else {
		_aab.Transform(_aae._bga, _cdd, im, im.Bounds(), _bae.Over, &_bae.Options{DstMask: _aae._fdd, DstMaskP: _e.Point{}})
	}
}
func (_gbb *Context) SetDashOffset(offset float64) { _gbb._bcg = offset }
func (_ecfe *Context) SetPixel(x, y int)           { _ecfe._bga.Set(x, y, _ecfe._db) }
func (_eaa *Context) Height() int                  { return _eaa._abb }
func (_gddg *Context) MeasureString(s string, face _bb.Face) (_cdf, _bef float64) {
	_cgea := &_bb.Drawer{Face: face}
	_cca := _cgea.MeasureString(s)
	return float64(_cca >> 6), _gddg._gfg.Tf.Size
}
func (_adb *Context) DrawStringAnchored(s string, face _bb.Face, x, y, ax, ay float64) {
	_gffg, _fcb := _adb.MeasureString(s, face)
	_adb.drawString(s, face, x-ax*_gffg, y+ay*_fcb)
}
func (_cbe *Context) DrawPoint(x, y, r float64) {
	_cbe.Push()
	_bba, _bgac := _cbe.Transform(x, y)
	_cbe.Identity()
	_cbe.DrawCircle(_bba, _bgac, r)
	_cbe.Pop()
}
func (_gbeg *radialGradient) ColorAt(x, y int) _da.Color {
	if len(_gbeg._bbaf) == 0 {
		return _da.Transparent
	}
	_dggd, _gbgf := float64(x)+0.5-_gbeg._fgb._ddc, float64(y)+0.5-_gbeg._fgb._fba
	_edcf := _fca(_dggd, _gbgf, _gbeg._fgb._gedc, _gbeg._dagg._ddc, _gbeg._dagg._fba, _gbeg._dagg._gedc)
	_dbca := _fca(_dggd, _gbgf, -_gbeg._fgb._gedc, _dggd, _gbgf, _gbeg._fgb._gedc)
	if _gbeg._dadg == 0 {
		if _edcf == 0 {
			return _da.Transparent
		}
		_defa := 0.5 * _dbca / _edcf
		if _defa*_gbeg._dagg._gedc >= _gbeg._afcb {
			return _bec(_defa, _gbeg._bbaf)
		}
		return _da.Transparent
	}
	_efbf := _fca(_edcf, _gbeg._dadg, 0, _edcf, -_dbca, 0)
	if _efbf >= 0 {
		_dddc := _d.Sqrt(_efbf)
		_ggf := (_edcf + _dddc) * _gbeg._ebe
		_aegc := (_edcf - _dddc) * _gbeg._ebe
		if _ggf*_gbeg._dagg._gedc >= _gbeg._afcb {
			return _bec(_ggf, _gbeg._bbaf)
		} else if _aegc*_gbeg._dagg._gedc >= _gbeg._afcb {
			return _bec(_aegc, _gbeg._bbaf)
		}
	}
	return _da.Transparent
}
func (_cab *linearGradient) ColorAt(x, y int) _da.Color {
	if len(_cab._ccd) == 0 {
		return _da.Transparent
	}
	_fddd, _edd := float64(x), float64(y)
	_gec, _fae, _caef, _baae := _cab._ddd, _cab._cac, _cab._fgc, _cab._dafc
	_egf, _gbdg := _caef-_gec, _baae-_fae
	if _gbdg == 0 && _egf != 0 {
		return _bec((_fddd-_gec)/_egf, _cab._ccd)
	}
	if _egf == 0 && _gbdg != 0 {
		return _bec((_edd-_fae)/_gbdg, _cab._ccd)
	}
	_bge := _egf*(_fddd-_gec) + _gbdg*(_edd-_fae)
	if _bge < 0 {
		return _cab._ccd[0]._cdfg
	}
	_abeg := _d.Hypot(_egf, _gbdg)
	_edeb := ((_fddd-_gec)*-_gbdg + (_edd-_fae)*_egf) / (_abeg * _abeg)
	_gbce, _dag := _gec+_edeb*-_gbdg, _fae+_edeb*_egf
	_cbcf := _d.Hypot(_fddd-_gbce, _edd-_dag) / _abeg
	return _bec(_cbcf, _cab._ccd)
}

type circle struct{ _ddc, _fba, _gedc float64 }

func (_daf *Context) Clip()                             { _daf.ClipPreserve(); _daf.ClearPath() }
func (_cec *Context) SetFillRule(fillRule _af.FillRule) { _cec._dgc = fillRule }

type linearGradient struct {
	_ddd, _cac, _fgc, _dafc float64
	_ccd                    stops
}

func (_gcbg *patternPainter) Paint(ss []_bc.Span, done bool) {
	_cagc := _gcbg._bgdc.Bounds()
	for _, _gceg := range ss {
		if _gceg.Y < _cagc.Min.Y {
			continue
		}
		if _gceg.Y >= _cagc.Max.Y {
			return
		}
		if _gceg.X0 < _cagc.Min.X {
			_gceg.X0 = _cagc.Min.X
		}
		if _gceg.X1 > _cagc.Max.X {
			_gceg.X1 = _cagc.Max.X
		}
		if _gceg.X0 >= _gceg.X1 {
			continue
		}
		const _gfgb = 1<<16 - 1
		_ebd := _gceg.Y - _gcbg._bgdc.Rect.Min.Y
		_bbbba := _gceg.X0 - _gcbg._bgdc.Rect.Min.X
		_fefg := (_gceg.Y-_gcbg._bgdc.Rect.Min.Y)*_gcbg._bgdc.Stride + (_gceg.X0-_gcbg._bgdc.Rect.Min.X)*4
		_fbb := _fefg + (_gceg.X1-_gceg.X0)*4
		for _fdc, _bdfb := _fefg, _bbbba; _fdc < _fbb; _fdc, _bdfb = _fdc+4, _bdfb+1 {
			_gbfg := _gceg.Alpha
			if _gcbg._cgf != nil {
				_gbfg = _gbfg * uint32(_gcbg._cgf.AlphaAt(_bdfb, _ebd).A) / 255
				if _gbfg == 0 {
					continue
				}
			}
			_efg := _gcbg._ega.ColorAt(_bdfb, _ebd)
			_baeb, _ebdc, _dgcg, _fcgd := _efg.RGBA()
			_cff := uint32(_gcbg._bgdc.Pix[_fdc+0])
			_bead := uint32(_gcbg._bgdc.Pix[_fdc+1])
			_deag := uint32(_gcbg._bgdc.Pix[_fdc+2])
			_ddf := uint32(_gcbg._bgdc.Pix[_fdc+3])
			_beee := (_gfgb - (_fcgd * _gbfg / _gfgb)) * 0x101
			_gcbg._bgdc.Pix[_fdc+0] = uint8((_cff*_beee + _baeb*_gbfg) / _gfgb >> 8)
			_gcbg._bgdc.Pix[_fdc+1] = uint8((_bead*_beee + _ebdc*_gbfg) / _gfgb >> 8)
			_gcbg._bgdc.Pix[_fdc+2] = uint8((_deag*_beee + _dgcg*_gbfg) / _gfgb >> 8)
			_gcbg._bgdc.Pix[_fdc+3] = uint8((_ddf*_beee + _fcgd*_gbfg) / _gfgb >> 8)
		}
	}
}
func (_ceg *Context) setFillAndStrokeColor(_fg _da.Color) {
	_ceg._db = _fg
	_ceg._gagc = _fdaf(_fg)
	_ceg._cgb = _fdaf(_fg)
}
func (_ead *Context) StrokePattern() _af.Pattern { return _ead._cgb }
func (_cgee *Context) DrawString(s string, face _bb.Face, x, y float64) {
	_cgee.DrawStringAnchored(s, face, x, y, 0, 0)
}

type radialGradient struct {
	_fgb, _fff, _dagg circle
	_dadg, _ebe       float64
	_afcb             float64
	_bbaf             stops
}

func _cb(_df, _aa, _cf, _ca, _fa, _fe, _bcc float64) (_bf, _dd float64) {
	_aag := 1 - _bcc
	_aaa := _aag * _aag
	_dg := 2 * _aag * _bcc
	_aeg := _bcc * _bcc
	_bf = _aaa*_df + _dg*_cf + _aeg*_fa
	_dd = _aaa*_aa + _dg*_ca + _aeg*_fe
	return
}
func (_gea *Context) DrawEllipse(x, y, rx, ry float64) {
	_gea.NewSubPath()
	_gea.DrawEllipticalArc(x, y, rx, ry, 0, 2*_d.Pi)
	_gea.ClosePath()
}
func (_be *Context) Image() _e.Image { return _be._bga }
func (_ggg *Context) DrawArc(x, y, r, angle1, angle2 float64) {
	_ggg.DrawEllipticalArc(x, y, r, r, angle1, angle2)
}
func (_fda *Context) NewSubPath() {
	if _fda._gff {
		_fda._feeb.Add1(_adbe(_fda._cef))
	}
	_fda._gff = false
}
func (_gbea *Context) Fill() {
	_gbea.FillPreserve()
	_gbea.ClearPath()
}
func (_eab *Context) ClosePath() {
	if _eab._gff {
		_gde := _adbe(_eab._cef)
		_eab._gge.Add1(_gde)
		_eab._feeb.Add1(_gde)
		_eab._gf = _eab._cef
	}
}

type solidPattern struct{ _bebf _da.Color }

func (_bac *Context) SetStrokeStyle(pattern _af.Pattern) { _bac._cgb = pattern }
func (_aaaa *Context) Scale(x, y float64)                { _aaaa._baab = _aaaa._baab.Scale(x, y) }
func (_dfaf *Context) SetMatrix(m _c.Matrix)             { _dfaf._baab = m }
