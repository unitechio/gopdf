package unichart

import (
	_e "bytes"
	_ef "fmt"
	_fd "image/color"
	_a "io"
	_c "math"

	_b "bitbucket.org/shenghui0779/gopdf/common"
	_eb "bitbucket.org/shenghui0779/gopdf/contentstream"
	_d "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_ae "bitbucket.org/shenghui0779/gopdf/core"
	_g "bitbucket.org/shenghui0779/gopdf/model"
	_eg "github.com/unidoc/unichart/render"
)

func (_dae *Renderer) MeasureText(text string) _eg.Box {
	_bg := _dae._bc
	_fbd, _bag := _dae._da.GetFontDescriptor()
	if _bag != nil {
		_b.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	} else {
		_cdf, _ebb := _fbd.GetCapHeight()
		if _ebb != nil {
			_b.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _ebb)
		} else {
			_bg = _cdf / 1000.0 * _dae._bc
		}
	}
	var (
		_bee = 0.0
		_adb = _dae.wrapText(text)
	)
	for _, _gf := range _adb {
		if _baa := _dae.getTextWidth(_gf); _baa > _bee {
			_bee = _baa
		}
	}
	_baf := _eg.NewBox(0, 0, int(_bee), int(_bg))
	if _bbb := _dae._ad; _bbb != 0 {
		_baf = _baf.Corners().Rotate(_bbb).Box()
	}
	return _baf
}
func NewRenderer(cc *_eb.ContentCreator, res *_g.PdfPageResources) func(int, int) (_eg.Renderer, error) {
	return func(_efb, _fe int) (_eg.Renderer, error) {
		_ecc := &Renderer{_cf: _efb, _ac: _fe, _ba: 72, _acb: cc, _ec: res, _aa: map[*_g.PdfFont]_ae.PdfObjectName{}}
		_ecc.ResetStyle()
		return _ecc, nil
	}
}
func (_aeb *Renderer) MoveTo(x, y int) { _aeb._acb.Add_m(float64(x), float64(y)) }
func (_gg *Renderer) GetDPI() float64  { return _gg._ba }

type Renderer struct {
	_cf  int
	_ac  int
	_ba  float64
	_acb *_eb.ContentCreator
	_ec  *_g.PdfPageResources
	_af  _fd.Color
	_ebf _fd.Color
	_cfd float64
	_da  *_g.PdfFont
	_bc  float64
	_afc _fd.Color
	_ad  float64
	_aa  map[*_g.PdfFont]_ae.PdfObjectName
}

func (_ff *Renderer) SetStrokeColor(color _fd.Color) {
	_ff._ebf = color
	_fg, _cd, _aag, _ := _cgb(color)
	_ff._acb.Add_RG(_fg, _cd, _aag)
}
func (_adg *Renderer) SetFillColor(color _fd.Color) {
	_adg._af = color
	_ebe, _ed, _aed, _ := _cgb(color)
	_adg._acb.Add_rg(_ebe, _ed, _aed)
}
func (_fae *Renderer) SetTextRotation(radians float64) { _fae._ad = _dgd(-radians) }
func (_cfcc *Renderer) Fill()                          { _cfcc._acb.Add_f() }
func _aebc(_aae _fd.Color) (uint8, uint8, uint8, uint8) {
	_fde, _edg, _egc, _baag := _aae.RGBA()
	return uint8(_fde >> 8), uint8(_edg >> 8), uint8(_egc >> 8), uint8(_baag >> 8)
}
func (_efg *Renderer) wrapText(_deg string) []string {
	var (
		_bbf []string
		_cfb []rune
	)
	for _, _bce := range _deg {
		if _bce == '\n' {
			_bbf = append(_bbf, string(_cfb))
			_cfb = []rune{}
			continue
		}
		_cfb = append(_cfb, _bce)
	}
	if len(_cfb) > 0 {
		_bbf = append(_bbf, string(_cfb))
	}
	return _bbf
}
func (_fc *Renderer) QuadCurveTo(cx, cy, x, y int) {
	_fc._acb.Add_v(float64(x), float64(y), float64(cx), float64(cy))
}
func (_gad *Renderer) SetFontColor(color _fd.Color) { _gad._afc = color }
func (_gd *Renderer) Text(text string, x, y int) {
	_gd._acb.Add_q()
	_gd.SetFont(_gd._da)
	_agc, _abfb, _ccf, _ := _cgb(_gd._afc)
	_gd._acb.Add_rg(_agc, _abfb, _ccf)
	_gd._acb.Translate(float64(x), float64(y)).Scale(1, -1)
	if _fda := _gd._ad; _fda != 0 {
		_gd._acb.RotateDeg(_fda)
	}
	_gd._acb.Add_BT().Add_TL(_gd._bc)
	var (
		_ccd = _gd._da.Encoder()
		_afe = _gd.wrapText(text)
		_cbd = len(_afe)
	)
	for _aff, _faf := range _afe {
		_gd._acb.Add_TJ(_ae.MakeStringFromBytes(_ccd.Encode(_faf)))
		if _aff != _cbd-1 {
			_gd._acb.Add_Tstar()
		}
	}
	_gd._acb.Add_ET()
	_gd._acb.Add_Q()
}
func (_dc *Renderer) ResetStyle() {
	_dc.SetFillColor(_fd.Black)
	_dc.SetStrokeColor(_fd.Transparent)
	_dc.SetStrokeWidth(0)
	_dc.SetFont(_g.DefaultFont())
	_dc.SetFontColor(_fd.Black)
	_dc.SetFontSize(12)
	_dc.SetTextRotation(0)
}
func _dgd(_gaa float64) float64                 { return _gaa * 180 / _c.Pi }
func (_cfc *Renderer) SetClassName(name string) {}
func _ceb(_abg float64) float64                 { return _abg * _c.Pi / 180.0 }
func (_ce *Renderer) Circle(radius float64, x, y int) {
	_ag := radius
	if _bca := _ce._cfd; _bca != 0 {
		_ag -= _bca / 2
	}
	_fcg := _ag * 0.551784
	_abeb := _d.CubicBezierPath{Curves: []_d.CubicBezierCurve{_d.NewCubicBezierCurve(-_ag, 0, -_ag, _fcg, -_fcg, _ag, 0, _ag), _d.NewCubicBezierCurve(0, _ag, _fcg, _ag, _ag, _fcg, _ag, 0), _d.NewCubicBezierCurve(_ag, 0, _ag, -_fcg, _fcg, -_ag, 0, -_ag), _d.NewCubicBezierCurve(0, -_ag, -_fcg, -_ag, -_ag, -_fcg, -_ag, 0)}}
	if _deb := _ce._cfd; _deb != 0 {
		_abeb = _abeb.Offset(_deb/2, _deb/2)
	}
	_abeb = _abeb.Offset(float64(x), float64(y))
	_d.DrawBezierPathWithCreator(_abeb, _ce._acb)
}
func _cgb(_df _fd.Color) (float64, float64, float64, float64) {
	_egg, _agg, _cbb, _bge := _aebc(_df)
	return float64(_egg) / 255, float64(_agg) / 255, float64(_cbb) / 255, float64(_bge) / 255
}
func (_ecb *Renderer) SetFont(font _eg.Font) {
	_dde, _eed := font.(*_g.PdfFont)
	if !_eed {
		_b.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
		return
	}
	_bb, _eed := _ecb._aa[_dde]
	if !_eed {
		_bb = _fdbg("\u0046\u006f\u006e\u0074", 1, _ecb._ec.HasFontByName)
		if _abf := _ecb._ec.SetFontByName(_bb, _dde.ToPdfObject()); _abf != nil {
			_b.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0064d\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0076\u0020\u0074\u006f\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073", _dde)
		}
		_ecb._aa[_dde] = _bb
	}
	_ecb._acb.Add_Tf(_bb, _ecb._bc)
	_ecb._da = _dde
}
func (_gc *Renderer) SetStrokeDashArray(dashArray []float64) {
	_bd := make([]int64, len(dashArray))
	for _de, _dd := range dashArray {
		_bd[_de] = int64(_dd)
	}
	_gc._acb.Add_d(_bd, 0)
}
func (_efd *Renderer) LineTo(x, y int)         { _efd._acb.Add_l(float64(x), float64(y)) }
func (_cge *Renderer) Close()                  { _cge._acb.Add_h() }
func (_be *Renderer) SetFontSize(size float64) { _be._bc = size }
func (_ga *Renderer) ArcTo(cx, cy int, rx, ry, startAngle, deltaAngle float64) {
	startAngle = _dgd(2.0*_c.Pi - startAngle)
	deltaAngle = _dgd(-deltaAngle)
	_eeb, _afd := deltaAngle, 1
	if _c.Abs(deltaAngle) > 90.0 {
		_afd = int(_c.Ceil(_c.Abs(deltaAngle) / 90.0))
		_eeb = deltaAngle / float64(_afd)
	}
	var (
		_fdb  = _ceb(_eeb / 2)
		_dg   = _c.Abs(4.0 / 3.0 * (1.0 - _c.Cos(_fdb)) / _c.Sin(_fdb))
		_afdb = float64(cx)
		_ea   = float64(cy)
	)
	for _fcc := 0; _fcc < _afd; _fcc++ {
		_cc := _ceb(startAngle + float64(_fcc)*_eeb)
		_ab := _ceb(startAngle + float64(_fcc+1)*_eeb)
		_fb := _c.Cos(_cc)
		_ge := _c.Cos(_ab)
		_gge := _c.Sin(_cc)
		_ege := _c.Sin(_ab)
		var _eef []float64
		if _eeb > 0 {
			_eef = []float64{_afdb + rx*_fb, _ea - ry*_gge, _afdb + rx*(_fb-_dg*_gge), _ea - ry*(_gge+_dg*_fb), _afdb + rx*(_ge+_dg*_ege), _ea - ry*(_ege-_dg*_ge), _afdb + rx*_ge, _ea - ry*_ege}
		} else {
			_eef = []float64{_afdb + rx*_fb, _ea - ry*_gge, _afdb + rx*(_fb+_dg*_gge), _ea - ry*(_gge-_dg*_fb), _afdb + rx*(_ge-_dg*_ege), _ea - ry*(_ege+_dg*_ge), _afdb + rx*_ge, _ea - ry*_ege}
		}
		if _fcc == 0 {
			_ga._acb.Add_l(_eef[0], _eef[1])
		}
		_ga._acb.Add_c(_eef[2], _eef[3], _eef[4], _eef[5], _eef[6], _eef[7])
	}
}
func _fdbg(_gb string, _bbc int, _bcb func(_ae.PdfObjectName) bool) _ae.PdfObjectName {
	_bdg := _ae.PdfObjectName(_ef.Sprintf("\u0025\u0073\u0025\u0064", _gb, _bbc))
	for _ddb := _bbc; _bcb(_bdg); {
		_ddb++
		_bdg = _ae.PdfObjectName(_ef.Sprintf("\u0025\u0073\u0025\u0064", _gb, _ddb))
	}
	return _bdg
}
func (_fa *Renderer) SetDPI(dpi float64) { _fa._ba = dpi }
func (_fgg *Renderer) getTextWidth(_fca string) float64 {
	var _gfc float64
	for _, _ccde := range _fca {
		_bed, _fec := _fgg._da.GetRuneMetrics(_ccde)
		if !_fec {
			_b.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074", _ccde)
		}
		_gfc += _bed.Wx
	}
	return _fgg._bc * _gfc / 1000.0
}
func (_cgf *Renderer) ClearTextRotation() { _cgf._ad = 0 }
func (_ebeg *Renderer) Save(w _a.Writer) error {
	if w == nil {
		return nil
	}
	_, _eeg := _a.Copy(w, _e.NewBuffer(_ebeg._acb.Bytes()))
	return _eeg
}
func (_abe *Renderer) Stroke()                     { _abe._acb.Add_S() }
func (_fdf *Renderer) FillStroke()                 { _fdf._acb.Add_B() }
func (_cb *Renderer) SetStrokeWidth(width float64) { _cb._cfd = width; _cb._acb.Add_w(width) }
