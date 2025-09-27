package unichart

import (
	_g "bytes"
	_c "fmt"
	_dd "image/color"
	_fb "io"
	_d "math"

	_ca "unitechio/gopdf/gopdf/common"
	_dgd "unitechio/gopdf/gopdf/contentstream"
	_dg "unitechio/gopdf/gopdf/contentstream/draw"
	_ff "unitechio/gopdf/gopdf/core"
	_e "unitechio/gopdf/gopdf/model"
	_a "github.com/unidoc/unichart/render"
)

type Renderer struct {
	_b    int
	_df   int
	_dge  float64
	_fc   *_dgd.ContentCreator
	_ee   *_e.PdfPageResources
	_dgea _dd.Color
	_fe   _dd.Color
	_gd   float64
	_ed   *_e.PdfFont
	_de   float64
	_ad   _dd.Color
	_eg   float64
	_bg   map[*_e.PdfFont]_ff.PdfObjectName
}

func (_dga *Renderer) MeasureText(text string) _a.Box {
	_bb := _dga._de
	_be, _dag := _dga._ed.GetFontDescriptor()
	if _dag != nil {
		_ca.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	} else {
		_feg, _fbc := _be.GetCapHeight()
		if _fbc != nil {
			_ca.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _fbc)
		} else {
			_bb = _feg / 1000.0 * _dga._de
		}
	}
	var (
		_afg = 0.0
		_bc  = _dga.wrapText(text)
	)
	for _, _ba := range _bc {
		if _cee := _dga.getTextWidth(_ba); _cee > _afg {
			_afg = _cee
		}
	}
	_gcd := _a.NewBox(0, 0, int(_afg), int(_bb))
	if _ab := _dga._eg; _ab != 0 {
		_gcd = _gcd.Corners().Rotate(_ab).Box()
	}
	return _gcd
}

func (_ecg *Renderer) Save(w _fb.Writer) error {
	if w == nil {
		return nil
	}
	_, _aec := _fb.Copy(w, _g.NewBuffer(_ecg._fc.Bytes()))
	return _aec
}

func (_ea *Renderer) SetStrokeColor(color _dd.Color) {
	_ea._fe = color
	_cc, _ag, _ge, _ := _fdd(color)
	_ea._fc.Add_RG(_cc, _ag, _ge)
}
func (_cf *Renderer) MoveTo(x, y int) { _cf._fc.Add_m(float64(x), float64(y)) }
func (_fd *Renderer) SetFont(font _a.Font) {
	_cff, _cfff := font.(*_e.PdfFont)
	if !_cfff {
		_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
		return
	}
	_egf, _cfff := _fd._bg[_cff]
	if !_cfff {
		_egf = _bac("\u0046\u006f\u006e\u0074", 1, _fd._ee.HasFontByName)
		if _fbe := _fd._ee.SetFontByName(_egf, _cff.ToPdfObject()); _fbe != nil {
			_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0064d\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0076\u0020\u0074\u006f\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073", _cff)
		}
		_fd._bg[_cff] = _egf
	}
	_fd._fc.Add_Tf(_egf, _fd._de)
	_fd._ed = _cff
}

func (_bf *Renderer) QuadCurveTo(cx, cy, x, y int) {
	_bf._fc.Add_v(float64(x), float64(y), float64(cx), float64(cy))
}

func NewRenderer(cc *_dgd.ContentCreator, res *_e.PdfPageResources) func(int, int) (_a.Renderer, error) {
	return func(_da, _dac int) (_a.Renderer, error) {
		_ac := &Renderer{_b: _da, _df: _dac, _dge: 72, _fc: cc, _ee: res, _bg: map[*_e.PdfFont]_ff.PdfObjectName{}}
		_ac.ResetStyle()
		return _ac, nil
	}
}
func (_eb *Renderer) Close()          { _eb._fc.Add_h() }
func (_age *Renderer) Fill()          { _age._fc.Add_f() }
func (_gda *Renderer) FillStroke()    { _gda._fc.Add_B() }
func _cbd(_dgga float64) float64      { return _dgga * 180 / _d.Pi }
func (_gc *Renderer) GetDPI() float64 { return _gc._dge }
func (_ccf *Renderer) wrapText(_bef string) []string {
	var (
		_gge []string
		_cgf []rune
	)
	for _, _cd := range _bef {
		if _cd == '\n' {
			_gge = append(_gge, string(_cgf))
			_cgf = []rune{}
			continue
		}
		_cgf = append(_cgf, _cd)
	}
	if len(_cgf) > 0 {
		_gge = append(_gge, string(_cgf))
	}
	return _gge
}

func (_af *Renderer) ArcTo(cx, cy int, rx, ry, startAngle, deltaAngle float64) {
	startAngle = _cbd(2.0*_d.Pi - startAngle)
	deltaAngle = _cbd(-deltaAngle)
	_faf, _caa := deltaAngle, 1
	if _d.Abs(deltaAngle) > 90.0 {
		_caa = int(_d.Ceil(_d.Abs(deltaAngle) / 90.0))
		_faf = deltaAngle / float64(_caa)
	}
	var (
		_fg  = _ecga(_faf / 2)
		_daf = _d.Abs(4.0 / 3.0 * (1.0 - _d.Cos(_fg)) / _d.Sin(_fg))
		_acc = float64(cx)
		_aca = float64(cy)
	)
	for _gf := 0; _gf < _caa; _gf++ {
		_aad := _ecga(startAngle + float64(_gf)*_faf)
		_gdf := _ecga(startAngle + float64(_gf+1)*_faf)
		_dfc := _d.Cos(_aad)
		_ffa := _d.Cos(_gdf)
		_eeb := _d.Sin(_aad)
		_ced := _d.Sin(_gdf)
		var _edf []float64
		if _faf > 0 {
			_edf = []float64{_acc + rx*_dfc, _aca - ry*_eeb, _acc + rx*(_dfc-_daf*_eeb), _aca - ry*(_eeb+_daf*_dfc), _acc + rx*(_ffa+_daf*_ced), _aca - ry*(_ced-_daf*_ffa), _acc + rx*_ffa, _aca - ry*_ced}
		} else {
			_edf = []float64{_acc + rx*_dfc, _aca - ry*_eeb, _acc + rx*(_dfc+_daf*_eeb), _aca - ry*(_eeb-_daf*_dfc), _acc + rx*(_ffa-_daf*_ced), _aca - ry*(_ced+_daf*_ffa), _acc + rx*_ffa, _aca - ry*_ced}
		}
		if _gf == 0 {
			_af._fc.Add_l(_edf[0], _edf[1])
		}
		_af._fc.Add_c(_edf[2], _edf[3], _edf[4], _edf[5], _edf[6], _edf[7])
	}
}
func (_ffc *Renderer) ClearTextRotation() { _ffc._eg = 0 }
func _ecga(_fddf float64) float64         { return _fddf * _d.Pi / 180.0 }
func (_dae *Renderer) Circle(radius float64, x, y int) {
	_cg := radius
	if _add := _dae._gd; _add != 0 {
		_cg -= _add / 2
	}
	_daeg := _cg * 0.551784
	_ae := _dg.CubicBezierPath{Curves: []_dg.CubicBezierCurve{_dg.NewCubicBezierCurve(-_cg, 0, -_cg, _daeg, -_daeg, _cg, 0, _cg), _dg.NewCubicBezierCurve(0, _cg, _daeg, _cg, _cg, _daeg, _cg, 0), _dg.NewCubicBezierCurve(_cg, 0, _cg, -_daeg, _daeg, -_cg, 0, -_cg), _dg.NewCubicBezierCurve(0, -_cg, -_daeg, -_cg, -_cg, -_daeg, -_cg, 0)}}
	if _gac := _dae._gd; _gac != 0 {
		_ae = _ae.Offset(_gac/2, _gac/2)
	}
	_ae = _ae.Offset(float64(x), float64(y))
	_dg.DrawBezierPathWithCreator(_ae, _dae._fc)
}
func (_adb *Renderer) LineTo(x, y int) { _adb._fc.Add_l(float64(x), float64(y)) }
func (_ef *Renderer) SetStrokeDashArray(dashArray []float64) {
	_ccc := make([]int64, len(dashArray))
	for _fa, _cce := range dashArray {
		_ccc[_fa] = int64(_cce)
	}
	_ef._fc.Add_d(_ccc, 0)
}
func (_edb *Renderer) SetStrokeWidth(width float64) { _edb._gd = width; _edb._fc.Add_w(width) }
func (_cfg *Renderer) SetFontSize(size float64)     { _cfg._de = size }
func (_ec *Renderer) SetDPI(dpi float64)            { _ec._dge = dpi }
func (_ga *Renderer) Stroke()                       { _ga._fc.Add_S() }
func (_cb *Renderer) SetFillColor(color _dd.Color) {
	_cb._dgea = color
	_ce, _ede, _db, _ := _fdd(color)
	_cb._fc.Add_rg(_ce, _ede, _db)
}

func (_bda *Renderer) getTextWidth(_adg string) float64 {
	var _ega float64
	for _, _abe := range _adg {
		_bca, _ccg := _bda._ed.GetRuneMetrics(_abe)
		if !_ccg {
			_ca.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074", _abe)
		}
		_ega += _bca.Wx
	}
	return _bda._de * _ega / 1000.0
}

func (_dbd *Renderer) Text(text string, x, y int) {
	_dbd._fc.Add_q()
	_dbd.SetFont(_dbd._ed)
	_agg, _eee, _gg, _ := _fdd(_dbd._ad)
	_dbd._fc.Add_rg(_agg, _eee, _gg)
	_dbd._fc.Translate(float64(x), float64(y)).Scale(1, -1)
	if _cca := _dbd._eg; _cca != 0 {
		_dbd._fc.RotateDeg(_cca)
	}
	_dbd._fc.Add_BT().Add_TL(_dbd._de)
	var (
		_ddd = _dbd._ed.Encoder()
		_dgg = _dbd.wrapText(text)
		_edg = len(_dgg)
	)
	for _fac, _gfg := range _dgg {
		_dbd._fc.Add_TJ(_ff.MakeStringFromBytes(_ddd.Encode(_gfg)))
		if _fac != _edg-1 {
			_dbd._fc.Add_Tstar()
		}
	}
	_dbd._fc.Add_ET()
	_dbd._fc.Add_Q()
}

func (_aa *Renderer) ResetStyle() {
	_aa.SetFillColor(_dd.Black)
	_aa.SetStrokeColor(_dd.Transparent)
	_aa.SetStrokeWidth(0)
	_aa.SetFont(_e.DefaultFont())
	_aa.SetFontColor(_dd.Black)
	_aa.SetFontSize(12)
	_aa.SetTextRotation(0)
}

func _abg(_ffae _dd.Color) (uint8, uint8, uint8, uint8) {
	_aecf, _eca, _gcc, _dfd := _ffae.RGBA()
	return uint8(_aecf >> 8), uint8(_eca >> 8), uint8(_gcc >> 8), uint8(_dfd >> 8)
}

func _fdd(_abb _dd.Color) (float64, float64, float64, float64) {
	_dgb, _eac, _ggd, _gccd := _abg(_abb)
	return float64(_dgb) / 255, float64(_eac) / 255, float64(_ggd) / 255, float64(_gccd) / 255
}
func (_cfga *Renderer) SetTextRotation(radians float64) { _cfga._eg = _cbd(-radians) }
func _bac(_eea string, _cgfd int, _gea func(_ff.PdfObjectName) bool) _ff.PdfObjectName {
	_gab := _ff.PdfObjectName(_c.Sprintf("\u0025\u0073\u0025\u0064", _eea, _cgfd))
	for _eacb := _cgfd; _gea(_gab); {
		_eacb++
		_gab = _ff.PdfObjectName(_c.Sprintf("\u0025\u0073\u0025\u0064", _eea, _eacb))
	}
	return _gab
}
func (_fcb *Renderer) SetClassName(name string)      {}
func (_edba *Renderer) SetFontColor(color _dd.Color) { _edba._ad = color }
