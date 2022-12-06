package unichart

import (
	_e "bytes"
	_ba "fmt"
	_ge "image/color"
	_g "io"
	_a "math"

	_eb "bitbucket.org/shenghui0779/gopdf/common"
	_bae "bitbucket.org/shenghui0779/gopdf/contentstream"
	_bad "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_f "bitbucket.org/shenghui0779/gopdf/core"
	_gg "bitbucket.org/shenghui0779/gopdf/model"
	_c "github.com/unidoc/unichart/render"
)

func (_bg *Renderer) SetStrokeDashArray(dashArray []float64) {
	_acf := make([]int64, len(dashArray))
	for _ffd, _bde := range dashArray {
		_acf[_ffd] = int64(_bde)
	}
	_bg._ac.Add_d(_acf, 0)
}
func (_aca *Renderer) SetFontSize(size float64) { _aca._gc = size }
func (_bea *Renderer) QuadCurveTo(cx, cy, x, y int) {
	_bea._ac.Add_v(float64(x), float64(y), float64(cx), float64(cy))
}
func (_adb *Renderer) MoveTo(x, y int)                 { _adb._ac.Add_m(float64(x), float64(y)) }
func (_gda *Renderer) SetTextRotation(radians float64) { _gda._dda = _faa(-radians) }
func (_eeb *Renderer) Save(w _g.Writer) error {
	if w == nil {
		return nil
	}
	_, _acg := _g.Copy(w, _e.NewBuffer(_eeb._ac.Bytes()))
	return _acg
}
func (_fc *Renderer) LineTo(x, y int) { _fc._ac.Add_l(float64(x), float64(y)) }
func (_bbd *Renderer) Stroke()        { _bbd._ac.Add_S() }
func _bge(_fgfd string, _cg int, _efd func(_f.PdfObjectName) bool) _f.PdfObjectName {
	_dfb := _f.PdfObjectName(_ba.Sprintf("\u0025\u0073\u0025\u0064", _fgfd, _cg))
	for _cbgc := _cg; _efd(_dfb); {
		_cbgc++
		_dfb = _f.PdfObjectName(_ba.Sprintf("\u0025\u0073\u0025\u0064", _fgfd, _cbgc))
	}
	return _dfb
}
func NewRenderer(cc *_bae.ContentCreator, res *_gg.PdfPageResources) func(int, int) (_c.Renderer, error) {
	return func(_da, _cb int) (_c.Renderer, error) {
		_gb := &Renderer{_bd: _da, _ff: _cb, _ad: 72, _ac: cc, _ggd: res, _ef: map[*_gg.PdfFont]_f.PdfObjectName{}}
		_gb.ResetStyle()
		return _gb, nil
	}
}
func (_ddc *Renderer) ResetStyle() {
	_ddc.SetFillColor(_ge.Black)
	_ddc.SetStrokeColor(_ge.Transparent)
	_ddc.SetStrokeWidth(0)
	_ddc.SetFont(_gg.DefaultFont())
	_ddc.SetFontColor(_ge.Black)
	_ddc.SetFontSize(12)
	_ddc.SetTextRotation(0)
}
func (_de *Renderer) Fill()                         { _de._ac.Add_f() }
func _dgg(_deg float64) float64                     { return _deg * _a.Pi / 180.0 }
func (_fgb *Renderer) SetStrokeWidth(width float64) { _fgb._ffa = width; _fgb._ac.Add_w(width) }
func (_gee *Renderer) SetClassName(name string)     {}
func (_fgc *Renderer) Text(text string, x, y int) {
	_fgc._ac.Add_q()
	_fgc.SetFont(_fgc._dd)
	_fec, _ae, _ce, _ := _beg(_fgc._bc)
	_fgc._ac.Add_rg(_fec, _ae, _ce)
	_fgc._ac.Translate(float64(x), float64(y)).Scale(1, -1)
	if _add := _fgc._dda; _add != 0 {
		_fgc._ac.RotateDeg(_add)
	}
	_fgc._ac.Add_BT().Add_TL(_fgc._gc)
	var (
		_ggf  = _fgc._dd.Encoder()
		_fedb = _fgc.wrapText(text)
		_fcga = len(_fedb)
	)
	for _cc, _addg := range _fedb {
		_fgc._ac.Add_TJ(_f.MakeStringFromBytes(_ggf.Encode(_addg)))
		if _cc != _fcga-1 {
			_fgc._ac.Add_Tstar()
		}
	}
	_fgc._ac.Add_ET()
	_fgc._ac.Add_Q()
}
func (_ccd *Renderer) wrapText(_afe string) []string {
	var (
		_bed []string
		_baa []rune
	)
	for _, _cef := range _afe {
		if _cef == '\n' {
			_bed = append(_bed, string(_baa))
			_baa = []rune{}
			continue
		}
		_baa = append(_baa, _cef)
	}
	if len(_baa) > 0 {
		_bed = append(_bed, string(_baa))
	}
	return _bed
}
func (_cd *Renderer) GetDPI() float64    { return _cd._ad }
func (_ea *Renderer) SetDPI(dpi float64) { _ea._ad = dpi }
func (_ffb *Renderer) FillStroke()       { _ffb._ac.Add_B() }
func (_be *Renderer) SetStrokeColor(color _ge.Color) {
	_be._fe = color
	_cda, _aa, _ca, _ := _beg(color)
	_be._ac.Add_RG(_cda, _aa, _ca)
}
func (_gd *Renderer) ArcTo(cx, cy int, rx, ry, startAngle, deltaAngle float64) {
	startAngle = _faa(2.0*_a.Pi - startAngle)
	deltaAngle = _faa(-deltaAngle)
	_ffaf, _gdb := deltaAngle, 1
	if _a.Abs(deltaAngle) > 90.0 {
		_gdb = int(_a.Ceil(_a.Abs(deltaAngle) / 90.0))
		_ffaf = deltaAngle / float64(_gdb)
	}
	var (
		_bf  = _dgg(_ffaf / 2)
		_bcg = _a.Abs(4.0 / 3.0 * (1.0 - _a.Cos(_bf)) / _a.Sin(_bf))
		_ggg = float64(cx)
		_cf  = float64(cy)
	)
	for _cdg := 0; _cdg < _gdb; _cdg++ {
		_cbb := _dgg(startAngle + float64(_cdg)*_ffaf)
		_ec := _dgg(startAngle + float64(_cdg+1)*_ffaf)
		_ffdg := _a.Cos(_cbb)
		_fed := _a.Cos(_ec)
		_ee := _a.Sin(_cbb)
		_bb := _a.Sin(_ec)
		var _af []float64
		if _ffaf > 0 {
			_af = []float64{_ggg + rx*_ffdg, _cf - ry*_ee, _ggg + rx*(_ffdg-_bcg*_ee), _cf - ry*(_ee+_bcg*_ffdg), _ggg + rx*(_fed+_bcg*_bb), _cf - ry*(_bb-_bcg*_fed), _ggg + rx*_fed, _cf - ry*_bb}
		} else {
			_af = []float64{_ggg + rx*_ffdg, _cf - ry*_ee, _ggg + rx*(_ffdg+_bcg*_ee), _cf - ry*(_ee-_bcg*_ffdg), _ggg + rx*(_fed-_bcg*_bb), _cf - ry*(_bb+_bcg*_fed), _ggg + rx*_fed, _cf - ry*_bb}
		}
		if _cdg == 0 {
			_gd._ac.Add_l(_af[0], _af[1])
		}
		_gd._ac.Add_c(_af[2], _af[3], _af[4], _af[5], _af[6], _af[7])
	}
}
func (_bga *Renderer) Close() { _bga._ac.Add_h() }
func (_daf *Renderer) SetFillColor(color _ge.Color) {
	_daf._d = color
	_aae, _cbg, _fg, _ := _beg(color)
	_daf._ac.Add_rg(_aae, _cbg, _fg)
}
func (_fgd *Renderer) Circle(radius float64, x, y int) {
	_ded := radius
	if _ggb := _fgd._ffa; _ggb != 0 {
		_ded -= _ggb / 2
	}
	_efe := _ded * 0.551784
	_fa := _bad.CubicBezierPath{Curves: []_bad.CubicBezierCurve{_bad.NewCubicBezierCurve(-_ded, 0, -_ded, _efe, -_efe, _ded, 0, _ded), _bad.NewCubicBezierCurve(0, _ded, _efe, _ded, _ded, _efe, _ded, 0), _bad.NewCubicBezierCurve(_ded, 0, _ded, -_efe, _efe, -_ded, 0, -_ded), _bad.NewCubicBezierCurve(0, -_ded, -_efe, -_ded, -_ded, -_efe, -_ded, 0)}}
	if _acc := _fgd._ffa; _acc != 0 {
		_fa = _fa.Offset(_acc/2, _acc/2)
	}
	_fa = _fa.Offset(float64(x), float64(y))
	_bad.DrawBezierPathWithCreator(_fa, _fgd._ac)
}
func (_dcf *Renderer) ClearTextRotation() { _dcf._dda = 0 }

type Renderer struct {
	_bd  int
	_ff  int
	_ad  float64
	_ac  *_bae.ContentCreator
	_ggd *_gg.PdfPageResources
	_d   _ge.Color
	_fe  _ge.Color
	_ffa float64
	_dd  *_gg.PdfFont
	_gc  float64
	_bc  _ge.Color
	_dda float64
	_ef  map[*_gg.PdfFont]_f.PdfObjectName
}

func _fdf(_dce _ge.Color) (uint8, uint8, uint8, uint8) {
	_ccf, _afc, _efb, _ebd := _dce.RGBA()
	return uint8(_ccf >> 8), uint8(_afc >> 8), uint8(_efb >> 8), uint8(_ebd >> 8)
}
func (_ccb *Renderer) MeasureText(text string) _c.Box {
	_ag := _ccb._gc
	_dg, _dc := _ccb._dd.GetFontDescriptor()
	if _dc != nil {
		_eb.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	} else {
		_fd, _fde := _dg.GetCapHeight()
		if _fde != nil {
			_eb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _fde)
		} else {
			_ag = _fd / 1000.0 * _ccb._gc
		}
	}
	var (
		_gcb = 0.0
		_fce = _ccb.wrapText(text)
	)
	for _, _fcf := range _fce {
		if _addb := _ccb.getTextWidth(_fcf); _addb > _gcb {
			_gcb = _addb
		}
	}
	_df := _c.NewBox(0, 0, int(_gcb), int(_ag))
	if _def := _ccb._dda; _def != 0 {
		_df = _df.Corners().Rotate(_def).Box()
	}
	return _df
}
func (_cfg *Renderer) getTextWidth(_fgf string) float64 {
	var _dca float64
	for _, _dbg := range _fgf {
		_feb, _gdc := _cfg._dd.GetRuneMetrics(_dbg)
		if !_gdc {
			_eb.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074", _dbg)
		}
		_dca += _feb.Wx
	}
	return _cfg._gc * _dca / 1000.0
}
func (_ed *Renderer) SetFontColor(color _ge.Color) { _ed._bc = color }
func (_fb *Renderer) SetFont(font _c.Font) {
	_fcg, _db := font.(*_gg.PdfFont)
	if !_db {
		_eb.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
		return
	}
	_bff, _db := _fb._ef[_fcg]
	if !_db {
		_bff = _bge("\u0046\u006f\u006e\u0074", 1, _fb._ggd.HasFontByName)
		if _aac := _fb._ggd.SetFontByName(_bff, _fcg.ToPdfObject()); _aac != nil {
			_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0064d\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0076\u0020\u0074\u006f\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073", _fcg)
		}
		_fb._ef[_fcg] = _bff
	}
	_fb._ac.Add_Tf(_bff, _fb._gc)
	_fb._dd = _fcg
}
func _faa(_cgc float64) float64 { return _cgc * 180 / _a.Pi }
func _beg(_bag _ge.Color) (float64, float64, float64, float64) {
	_dgf, _bdf, _ccdc, _adc := _fdf(_bag)
	return float64(_dgf) / 255, float64(_bdf) / 255, float64(_ccdc) / 255, float64(_adc) / 255
}
