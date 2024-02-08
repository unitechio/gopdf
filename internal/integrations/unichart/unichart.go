package unichart

import (
	_g "bytes"
	_b "fmt"
	_f "image/color"
	_a "io"
	_gc "math"

	_cd "bitbucket.org/shenghui0779/gopdf/common"
	_cdc "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ef "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_aeb "bitbucket.org/shenghui0779/gopdf/core"
	_ae "bitbucket.org/shenghui0779/gopdf/model"
	_c "github.com/unidoc/unichart/render"
)

func (_fd *Renderer) SetFillColor(color _f.Color) {
	_fd._gd = color
	_dc, _cde, _bc, _ := _gadc(color)
	_fd._aeg.Add_rg(_dc, _cde, _bc)
}
func (_acc *Renderer) FillStroke()                  { _acc._aeg.Add_B() }
func (_fgb *Renderer) SetStrokeWidth(width float64) { _fgb._gdc = width; _fgb._aeg.Add_w(width) }
func (_ec *Renderer) wrapText(_cfe string) []string {
	var (
		_db  []string
		_ede []rune
	)
	for _, _dce := range _cfe {
		if _dce == '\n' {
			_db = append(_db, string(_ede))
			_ede = []rune{}
			continue
		}
		_ede = append(_ede, _dce)
	}
	if len(_ede) > 0 {
		_db = append(_db, string(_ede))
	}
	return _db
}
func _aaa(_dfe string, _ace int, _aag func(_aeb.PdfObjectName) bool) _aeb.PdfObjectName {
	_efb := _aeb.PdfObjectName(_b.Sprintf("\u0025\u0073\u0025\u0064", _dfe, _ace))
	for _bgcde := _ace; _aag(_efb); {
		_bgcde++
		_efb = _aeb.PdfObjectName(_b.Sprintf("\u0025\u0073\u0025\u0064", _dfe, _bgcde))
	}
	return _efb
}
func _gadc(_eeg _f.Color) (float64, float64, float64, float64) {
	_fab, _fda, _edd, _ccdc := _ggf(_eeg)
	return float64(_fab) / 255, float64(_fda) / 255, float64(_edd) / 255, float64(_ccdc) / 255
}
func (_gea *Renderer) MeasureText(text string) _c.Box {
	_daa := _gea._ga
	_fcf, _fae := _gea._bg.GetFontDescriptor()
	if _fae != nil {
		_cd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	} else {
		_ccg, _gcbc := _fcf.GetCapHeight()
		if _gcbc != nil {
			_cd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _gcbc)
		} else {
			_daa = _ccg / 1000.0 * _gea._ga
		}
	}
	var (
		_bf  = 0.0
		_dac = _gea.wrapText(text)
	)
	for _, _egde := range _dac {
		if _gde := _gea.getTextWidth(_egde); _gde > _bf {
			_bf = _gde
		}
	}
	_cca := _c.NewBox(0, 0, int(_bf), int(_daa))
	if _gae := _gea._gcf; _gae != 0 {
		_cca = _cca.Corners().Rotate(_gae).Box()
	}
	return _cca
}
func _dfc(_bba float64) float64 { return _bba * 180 / _gc.Pi }
func (_dd *Renderer) ArcTo(cx, cy int, rx, ry, startAngle, deltaAngle float64) {
	startAngle = _dfc(2.0*_gc.Pi - startAngle)
	deltaAngle = _dfc(-deltaAngle)
	_gb, _ffb := deltaAngle, 1
	if _gc.Abs(deltaAngle) > 90.0 {
		_ffb = int(_gc.Ceil(_gc.Abs(deltaAngle) / 90.0))
		_gb = deltaAngle / float64(_ffb)
	}
	var (
		_cdd = _dgg(_gb / 2)
		_df  = _gc.Abs(4.0 / 3.0 * (1.0 - _gc.Cos(_cdd)) / _gc.Sin(_cdd))
		_ed  = float64(cx)
		_fcb = float64(cy)
	)
	for _fgd := 0; _fgd < _ffb; _fgd++ {
		_gdg := _dgg(startAngle + float64(_fgd)*_gb)
		_gdd := _dgg(startAngle + float64(_fgd+1)*_gb)
		_eb := _gc.Cos(_gdg)
		_ffc := _gc.Cos(_gdd)
		_gaca := _gc.Sin(_gdg)
		_af := _gc.Sin(_gdd)
		var _ddd []float64
		if _gb > 0 {
			_ddd = []float64{_ed + rx*_eb, _fcb - ry*_gaca, _ed + rx*(_eb-_df*_gaca), _fcb - ry*(_gaca+_df*_eb), _ed + rx*(_ffc+_df*_af), _fcb - ry*(_af-_df*_ffc), _ed + rx*_ffc, _fcb - ry*_af}
		} else {
			_ddd = []float64{_ed + rx*_eb, _fcb - ry*_gaca, _ed + rx*(_eb+_df*_gaca), _fcb - ry*(_gaca-_df*_eb), _ed + rx*(_ffc-_df*_af), _fcb - ry*(_af+_df*_ffc), _ed + rx*_ffc, _fcb - ry*_af}
		}
		if _fgd == 0 {
			_dd._aeg.Add_l(_ddd[0], _ddd[1])
		}
		_dd._aeg.Add_c(_ddd[2], _ddd[3], _ddd[4], _ddd[5], _ddd[6], _ddd[7])
	}
}
func (_fac *Renderer) SetClassName(name string) {}
func _dgg(_gbeb float64) float64                { return _gbeb * _gc.Pi / 180.0 }
func (_gcb *Renderer) LineTo(x, y int)          { _gcb._aeg.Add_l(float64(x), float64(y)) }
func (_gee *Renderer) SetFont(font _c.Font) {
	_gede, _ba := font.(*_ae.PdfFont)
	if !_ba {
		_cd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
		return
	}
	_def, _ba := _gee._de[_gede]
	if !_ba {
		_def = _aaa("\u0046\u006f\u006e\u0074", 1, _gee._fa.HasFontByName)
		if _gge := _gee._fa.SetFontByName(_def, _gede.ToPdfObject()); _gge != nil {
			_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0064d\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0076\u0020\u0074\u006f\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073", _gede)
		}
		_gee._de[_gede] = _def
	}
	_gee._aeg.Add_Tf(_def, _gee._ga)
	_gee._bg = _gede
}
func (_cg *Renderer) GetDPI() float64 { return _cg._d }
func (_acf *Renderer) Stroke()        { _acf._aeg.Add_S() }
func (_acd *Renderer) getTextWidth(_bda string) float64 {
	var _bad float64
	for _, _bea := range _bda {
		_gbe, _egf := _acd._bg.GetRuneMetrics(_bea)
		if !_egf {
			_cd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074", _bea)
		}
		_bad += _gbe.Wx
	}
	return _acd._ga * _bad / 1000.0
}

type Renderer struct {
	_ge  int
	_ea  int
	_d   float64
	_aeg *_cdc.ContentCreator
	_fa  *_ae.PdfPageResources
	_gd  _f.Color
	_be  _f.Color
	_gdc float64
	_bg  *_ae.PdfFont
	_ga  float64
	_bgc _f.Color
	_gcf float64
	_de  map[*_ae.PdfFont]_aeb.PdfObjectName
}

func _ggf(_cfb _f.Color) (uint8, uint8, uint8, uint8) {
	_ege, _ffd, _ccdba, _bgf := _cfb.RGBA()
	return uint8(_ege >> 8), uint8(_ffd >> 8), uint8(_ccdba >> 8), uint8(_bgf >> 8)
}
func (_bbb *Renderer) ClearTextRotation() { _bbb._gcf = 0 }
func (_ebd *Renderer) Save(w _a.Writer) error {
	if w == nil {
		return nil
	}
	_, _gag := _a.Copy(w, _g.NewBuffer(_ebd._aeg.Bytes()))
	return _gag
}
func (_fdga *Renderer) SetFontColor(color _f.Color) { _fdga._bgc = color }
func (_gdb *Renderer) Fill()                        { _gdb._aeg.Add_f() }
func NewRenderer(cc *_cdc.ContentCreator, res *_ae.PdfPageResources) func(int, int) (_c.Renderer, error) {
	return func(_dg, _cc int) (_c.Renderer, error) {
		_gac := &Renderer{_ge: _dg, _ea: _cc, _d: 72, _aeg: cc, _fa: res, _de: map[*_ae.PdfFont]_aeb.PdfObjectName{}}
		_gac.ResetStyle()
		return _gac, nil
	}
}
func (_ccdb *Renderer) MoveTo(x, y int)                { _ccdb._aeg.Add_m(float64(x), float64(y)) }
func (_dcg *Renderer) SetTextRotation(radians float64) { _dcg._gcf = _dfc(-radians) }
func (_fef *Renderer) Text(text string, x, y int) {
	_fef._aeg.Add_q()
	_fef.SetFont(_fef._bg)
	_bd, _ebc, _bdc, _ := _gadc(_fef._bgc)
	_fef._aeg.Add_rg(_bd, _ebc, _bdc)
	_fef._aeg.Translate(float64(x), float64(y)).Scale(1, -1)
	if _egd := _fef._gcf; _egd != 0 {
		_fef._aeg.RotateDeg(_egd)
	}
	_fef._aeg.Add_BT().Add_TL(_fef._ga)
	var (
		_cff  = _fef._bg.Encoder()
		_bgcd = _fef.wrapText(text)
		_faf  = len(_bgcd)
	)
	for _cb, _eggc := range _bgcd {
		_fef._aeg.Add_TJ(_aeb.MakeStringFromBytes(_cff.Encode(_eggc)))
		if _cb != _faf-1 {
			_fef._aeg.Add_Tstar()
		}
	}
	_fef._aeg.Add_ET()
	_fef._aeg.Add_Q()
}
func (_gg *Renderer) SetStrokeColor(color _f.Color) {
	_gg._be = color
	_fg, _ged, _fc, _ := _gadc(color)
	_gg._aeg.Add_RG(_fg, _ged, _fc)
}
func (_aa *Renderer) Circle(radius float64, x, y int) {
	_fdg := radius
	if _gef := _aa._gdc; _gef != 0 {
		_fdg -= _gef / 2
	}
	_bb := _fdg * 0.551784
	_ca := _ef.CubicBezierPath{Curves: []_ef.CubicBezierCurve{_ef.NewCubicBezierCurve(-_fdg, 0, -_fdg, _bb, -_bb, _fdg, 0, _fdg), _ef.NewCubicBezierCurve(0, _fdg, _bb, _fdg, _fdg, _bb, _fdg, 0), _ef.NewCubicBezierCurve(_fdg, 0, _fdg, -_bb, _bb, -_fdg, 0, -_fdg), _ef.NewCubicBezierCurve(0, -_fdg, -_bb, -_fdg, -_fdg, -_bb, -_fdg, 0)}}
	if _cgf := _aa._gdc; _cgf != 0 {
		_ca = _ca.Offset(_cgf/2, _cgf/2)
	}
	_ca = _ca.Offset(float64(x), float64(y))
	_ef.DrawBezierPathWithCreator(_ca, _aa._aeg)
}
func (_ac *Renderer) QuadCurveTo(cx, cy, x, y int) {
	_ac._aeg.Add_v(float64(x), float64(y), float64(cx), float64(cy))
}
func (_da *Renderer) Close()                    { _da._aeg.Add_h() }
func (_egg *Renderer) SetFontSize(size float64) { _egg._ga = size }
func (_ee *Renderer) SetStrokeDashArray(dashArray []float64) {
	_efa := make([]int64, len(dashArray))
	for _cf, _eg := range dashArray {
		_efa[_cf] = int64(_eg)
	}
	_ee._aeg.Add_d(_efa, 0)
}
func (_ccd *Renderer) SetDPI(dpi float64) { _ccd._d = dpi }
func (_ab *Renderer) ResetStyle() {
	_ab.SetFillColor(_f.Black)
	_ab.SetStrokeColor(_f.Transparent)
	_ab.SetStrokeWidth(0)
	_ab.SetFont(_ae.DefaultFont())
	_ab.SetFontColor(_f.Black)
	_ab.SetFontSize(12)
	_ab.SetTextRotation(0)
}
