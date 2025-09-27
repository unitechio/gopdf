package svg

import (
	_bg "encoding/xml"
	_b "fmt"
	_g "io"
	_f "math"
	_ac "os"
	_af "strconv"
	_a "strings"
	_d "unicode"

	_be "unitechio/gopdf/gopdf/common"
	_fc "unitechio/gopdf/gopdf/contentstream"
	_gf "unitechio/gopdf/gopdf/contentstream/draw"
	_gd "unitechio/gopdf/gopdf/internal/graphic2d"
	_ab "golang.org/x/net/html/charset"
)

func _gfeg(_cae []token) ([]*Command, error) {
	var (
		_eeeg []*Command
		_aaga []float64
	)
	for _efe := len(_cae) - 1; _efe >= 0; _efe-- {
		_aca := _cae[_efe]
		if _aca._gda {
			_dgf := _bgd._cad[_a.ToLower(_aca._dbb)]
			_fff := len(_aaga)
			if _dgf == 0 && _fff == 0 {
				_ggg := &Command{Symbol: _aca._dbb}
				_eeeg = append([]*Command{_ggg}, _eeeg...)
			} else if _dgf != 0 && _fff%_dgf == 0 {
				_faae := _fff / _dgf
				for _dcg := 0; _dcg < _faae; _dcg++ {
					_eeg := _aca._dbb
					if _eeg == "\u006d" && _dcg < _faae-1 {
						_eeg = "\u006c"
					}
					if _eeg == "\u004d" && _dcg < _faae-1 {
						_eeg = "\u004c"
					}
					_ccg := &Command{_eeg, _bcfd(_aaga[:_dgf])}
					_eeeg = append([]*Command{_ccg}, _eeeg...)
					_aaga = _aaga[_dgf:]
				}
			} else {
				_egag := pathParserError{"I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006fr\u0020" + _aca._dbb}
				return nil, _egag
			}
		} else {
			_adae, _gdd := _ged(_aca._dbb, 64)
			if _gdd != nil {
				return nil, _gdd
			}
			_aaga = append(_aaga, _adae)
		}
	}
	return _eeeg, nil
}

func (_ffgc *Path) compare(_ebg *Path) bool {
	if len(_ffgc.Subpaths) != len(_ebg.Subpaths) {
		return false
	}
	for _cgf, _bge := range _ffgc.Subpaths {
		if !_bge.compare(_ebg.Subpaths[_cgf]) {
			return false
		}
	}
	return true
}

type token struct {
	_dbb string
	_gda bool
}

func (_fgc *GraphicSVG) setDefaultScaling(_fef float64) {
	_fgc._cgaa = _fef
	if _fgc.Style != nil && _fgc.Style.StrokeWidth > 0 {
		_fgc.Style.StrokeWidth = _fgc.Style.StrokeWidth * _fgc._cgaa
	}
	for _, _eee := range _fgc.Children {
		_eee.setDefaultScaling(_fef)
	}
}

func _dde(_aac string) (_efdb []float64, _adfg error) {
	var _aafb float64
	_cefd := 0
	_cedg := true
	for _eega, _fbcd := range _aac {
		if _fbcd == '.' {
			if _cedg {
				_cedg = false
				continue
			}
			_aafb, _adfg = _ged(_aac[_cefd:_eega], 64)
			if _adfg != nil {
				return
			}
			_efdb = append(_efdb, _aafb)
			_cefd = _eega
		}
	}
	_aafb, _adfg = _ged(_aac[_cefd:], 64)
	if _adfg != nil {
		return
	}
	_efdb = append(_efdb, _aafb)
	return
}

func _bbc(_egg *GraphicSVG, _afab *_fc.ContentCreator) {
	_afab.Add_q()
	_egg.Style.toContentStream(_afab)
	_fgb, _ccf := _ged(_egg.Attributes["\u0063\u0078"], 64)
	if _ccf != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _ccf.Error())
	}
	_cgg, _ccf := _ged(_egg.Attributes["\u0063\u0079"], 64)
	if _ccf != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _ccf.Error())
	}
	_gae, _ccf := _ged(_egg.Attributes["\u0072\u0078"], 64)
	if _ccf != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _ccf.Error())
	}
	_ccfd, _ccf := _ged(_egg.Attributes["\u0072\u0079"], 64)
	if _ccf != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _ccf.Error())
	}
	_gef := _gae * _egg._cgaa
	_fbc := _ccfd * _egg._cgaa
	_ced := _fgb * _egg._cgaa
	_gfd := _cgg * _egg._cgaa
	_fedb := _gef * _df
	_dfd := _fbc * _df
	_bae := _gf.NewCubicBezierPath()
	_bae = _bae.AppendCurve(_gf.NewCubicBezierCurve(-_gef, 0, -_gef, _dfd, -_fedb, _fbc, 0, _fbc))
	_bae = _bae.AppendCurve(_gf.NewCubicBezierCurve(0, _fbc, _fedb, _fbc, _gef, _dfd, _gef, 0))
	_bae = _bae.AppendCurve(_gf.NewCubicBezierCurve(_gef, 0, _gef, -_dfd, _fedb, -_fbc, 0, -_fbc))
	_bae = _bae.AppendCurve(_gf.NewCubicBezierCurve(0, -_fbc, -_fedb, -_fbc, -_gef, -_dfd, -_gef, 0))
	_bae = _bae.Offset(_ced, _gfd)
	if _egg.Style.StrokeWidth > 0 {
		_bae = _bae.Offset(_egg.Style.StrokeWidth/2, _egg.Style.StrokeWidth/2)
	}
	_gf.DrawBezierPathWithCreator(_bae, _afab)
	if _egg.Style.FillColor != "" && _egg.Style.StrokeColor != "" {
		_afab.Add_B()
	} else if _egg.Style.FillColor != "" {
		_afab.Add_f()
	} else if _egg.Style.StrokeColor != "" {
		_afab.Add_S()
	}
	_afab.Add_h()
	_afab.Add_Q()
}

func (_abdd *GraphicSVGStyle) toContentStream(_edc *_fc.ContentCreator) {
	if _abdd == nil {
		return
	}
	if _abdd.FillColor != "" {
		var _edbe, _afgb, _afd float64
		if _ebd, _bde := _gd.ColorMap[_abdd.FillColor]; _bde {
			_cbe, _aae, _bfbe, _ := _ebd.RGBA()
			_edbe, _afgb, _afd = float64(_cbe), float64(_aae), float64(_bfbe)
		} else {
			_edbe, _afgb, _afd = _aee(_abdd.FillColor)
		}
		_edc.Add_rg(_edbe, _afgb, _afd)
	}
	if _abdd.StrokeColor != "" {
		var _cca, _ffae, _dcd float64
		if _dfdf, _dcbg := _gd.ColorMap[_abdd.StrokeColor]; _dcbg {
			_ceea, _gcfc, _geb, _ := _dfdf.RGBA()
			_cca, _ffae, _dcd = float64(_ceea)/255.0, float64(_gcfc)/255.0, float64(_geb)/255.0
		} else {
			_cca, _ffae, _dcd = _aee(_abdd.StrokeColor)
		}
		_edc.Add_RG(_cca, _ffae, _dcd)
	}
	if _abdd.StrokeWidth > 0 {
		_edc.Add_w(_abdd.StrokeWidth)
	}
}

var _bgd commands

func _dacc() *GraphicSVGStyle {
	return &GraphicSVGStyle{FillColor: "\u00230\u0030\u0030\u0030\u0030\u0030", StrokeColor: "", StrokeWidth: 0}
}

func _cfb(_bbcf _bg.StartElement) *GraphicSVG {
	_ddc := &GraphicSVG{}
	_ece := make(map[string]string)
	for _, _bfdc := range _bbcf.Attr {
		_ece[_bfdc.Name.Local] = _bfdc.Value
	}
	_ddc.Name = _bbcf.Name.Local
	_ddc.Attributes = _ece
	_ddc._cgaa = 1
	if _ddc.Name == "\u0073\u0076\u0067" {
		_cd, _ecg := _edf(_ece["\u0076i\u0065\u0077\u0042\u006f\u0078"])
		if _ecg != nil {
			_be.Log.Debug("\u0055\u006ea\u0062\u006c\u0065\u0020t\u006f\u0020p\u0061\u0072\u0073\u0065\u0020\u0076\u0069\u0065w\u0042\u006f\u0078\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074e\u003a\u0020\u0025\u0076", _ecg)
			return nil
		}
		_ddc.ViewBox.X = _cd[0]
		_ddc.ViewBox.Y = _cd[1]
		_ddc.ViewBox.W = _cd[2]
		_ddc.ViewBox.H = _cd[3]
		_ddc.Width = _ddc.ViewBox.W
		_ddc.Height = _ddc.ViewBox.H
		if _edb, _bca := _ece["\u0077\u0069\u0064t\u0068"]; _bca {
			_eff, _dgd := _ged(_edb, 64)
			if _dgd != nil {
				_be.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0077\u0069\u0064\u0074\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _dgd)
				return nil
			}
			_ddc.Width = _eff
		}
		if _effa, _ggd := _ece["\u0068\u0065\u0069\u0067\u0068\u0074"]; _ggd {
			_acd, _ffa := _ged(_effa, 64)
			if _ffa != nil {
				_be.Log.Debug("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0070\u0061rse\u0020he\u0069\u0067\u0068\u0074\u0020\u0061\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _ffa)
				return nil
			}
			_ddc.Height = _acd
		}
		if _ddc.Width > 0 && _ddc.Height > 0 {
			_ddc._cgaa = _ddc.Width / _ddc.ViewBox.W
		}
	}
	return _ddc
}

func _ffb(_ffdb *GraphicSVG, _dee *_fc.ContentCreator) {
	_dee.Add_q()
	_ffdb.Style.toContentStream(_dee)
	_beb, _gge := _edf(_ffdb.Attributes["\u0070\u006f\u0069\u006e\u0074\u0073"])
	if _gge != nil {
		_be.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076", _gge)
		return
	}
	if len(_beb)%2 > 0 {
		_be.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068")
		return
	}
	for _bbd := 0; _bbd < len(_beb); {
		if _bbd == 0 {
			_dee.Add_m(_beb[_bbd]*_ffdb._cgaa, _beb[_bbd+1]*_ffdb._cgaa)
		} else {
			_dee.Add_l(_beb[_bbd]*_ffdb._cgaa, _beb[_bbd+1]*_ffdb._cgaa)
		}
		_bbd += 2
	}
	_dee.Add_l(_beb[0]*_ffdb._cgaa, _beb[1]*_ffdb._cgaa)
	if _ffdb.Style.FillColor != "" && _ffdb.Style.StrokeColor != "" {
		_dee.Add_B()
	} else if _ffdb.Style.FillColor != "" {
		_dee.Add_f()
	} else if _ffdb.Style.StrokeColor != "" {
		_dee.Add_S()
	}
	_dee.Add_h()
	_dee.Add_Q()
}

type GraphicSVGStyle struct {
	FillColor   string
	StrokeColor string
	StrokeWidth float64
}

func _eggf(_ecd float64) int { return int(_ecd + _f.Copysign(0.5, _ecd)) }
func (_abd *GraphicSVG) SetScaling(xFactor, yFactor float64) {
	_bce := _abd.Width / _abd.ViewBox.W
	_fec := _abd.Height / _abd.ViewBox.H
	_abd.setDefaultScaling(_f.Max(_bce, _fec))
	for _, _cbfe := range _abd.Children {
		_cbfe.SetScaling(xFactor, yFactor)
	}
}

func _agb(_cbfea *_bg.Decoder) (*GraphicSVG, error) {
	for {
		_gefa, _bcbf := _cbfea.Token()
		if _gefa == nil && _bcbf == _g.EOF {
			break
		}
		if _bcbf != nil {
			return nil, _bcbf
		}
		switch _ede := _gefa.(type) {
		case _bg.StartElement:
			return _cfb(_ede), nil
		}
	}
	return &GraphicSVG{}, nil
}

func ParseFromString(svgStr string) (*GraphicSVG, error) {
	return ParseFromStream(_a.NewReader(svgStr))
}

func _cc(_bea *GraphicSVG, _fe *_fc.ContentCreator) {
	_fe.Add_q()
	_bea.Style.toContentStream(_fe)
	_ge, _cee := _acb(_bea.Attributes["\u0064"])
	if _cee != nil {
		_be.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025s", _cee.Error())
	}
	var (
		_dc, _da  = 0.0, 0.0
		_bed, _eb = 0.0, 0.0
		_ceg      *Command
	)
	for _, _ef := range _ge.Subpaths {
		for _, _ee := range _ef.Commands {
			switch _a.ToLower(_ee.Symbol) {
			case "\u006d":
				_bed, _eb = _ee.Params[0]*_bea._cgaa, _ee.Params[1]*_bea._cgaa
				if !_ee.isAbsolute() {
					_bed, _eb = _dc+_bed-_bea.ViewBox.X, _da+_eb-_bea.ViewBox.Y
				}
				_fe.Add_m(_ddf(_bed, 3), _ddf(_eb, 3))
				_dc, _da = _bed, _eb
			case "\u0063":
				_ga, _bc, _eg, _ba, _gb, _bab := _ee.Params[0]*_bea._cgaa, _ee.Params[1]*_bea._cgaa, _ee.Params[2]*_bea._cgaa, _ee.Params[3]*_bea._cgaa, _ee.Params[4]*_bea._cgaa, _ee.Params[5]*_bea._cgaa
				if !_ee.isAbsolute() {
					_ga, _bc, _eg, _ba, _gb, _bab = _dc+_ga, _da+_bc, _dc+_eg, _da+_ba, _dc+_gb, _da+_bab
				}
				_fe.Add_c(_ddf(_ga, 3), _ddf(_bc, 3), _ddf(_eg, 3), _ddf(_ba, 3), _ddf(_gb, 3), _ddf(_bab, 3))
				_dc, _da = _gb, _bab
			case "\u0073":
				_cg, _ega, _bb, _cga := _ee.Params[0]*_bea._cgaa, _ee.Params[1]*_bea._cgaa, _ee.Params[2]*_bea._cgaa, _ee.Params[3]*_bea._cgaa
				if !_ee.isAbsolute() {
					_cg, _ega, _bb, _cga = _dc+_cg, _da+_ega, _dc+_bb, _da+_cga
				}
				_fe.Add_c(_ddf(_dc, 3), _ddf(_da, 3), _ddf(_cg, 3), _ddf(_ega, 3), _ddf(_bb, 3), _ddf(_cga, 3))
				_dc, _da = _bb, _cga
			case "\u006c":
				_fg, _gc := _ee.Params[0]*_bea._cgaa, _ee.Params[1]*_bea._cgaa
				if !_ee.isAbsolute() {
					_fg, _gc = _dc+_fg, _da+_gc
				}
				_fe.Add_l(_ddf(_fg, 3), _ddf(_gc, 3))
				_dc, _da = _fg, _gc
			case "\u0068":
				_de := _ee.Params[0] * _bea._cgaa
				if !_ee.isAbsolute() {
					_de = _dc + _de
				}
				_fe.Add_l(_ddf(_de, 3), _ddf(_da, 3))
				_dc = _de
			case "\u0076":
				_bf := _ee.Params[0] * _bea._cgaa
				if !_ee.isAbsolute() {
					_bf = _da + _bf
				}
				_fe.Add_l(_ddf(_dc, 3), _ddf(_bf, 3))
				_da = _bf
			case "\u0071":
				_ae, _ebc, _afg, _eed := _ee.Params[0]*_bea._cgaa, _ee.Params[1]*_bea._cgaa, _ee.Params[2]*_bea._cgaa, _ee.Params[3]*_bea._cgaa
				if !_ee.isAbsolute() {
					_ae, _ebc, _afg, _eed = _dc+_ae, _da+_ebc, _dc+_afg, _da+_eed
				}
				_ea, _bd := _gd.QuadraticToCubicBezier(_dc, _da, _ae, _ebc, _afg, _eed)
				_fe.Add_c(_ddf(_ea.X, 3), _ddf(_ea.Y, 3), _ddf(_bd.X, 3), _ddf(_bd.Y, 3), _ddf(_afg, 3), _ddf(_eed, 3))
				_dc, _da = _afg, _eed
			case "\u0074":
				var _ag, _eaf _gd.Point
				_ff, _bfb := _ee.Params[0]*_bea._cgaa, _ee.Params[1]*_bea._cgaa
				if !_ee.isAbsolute() {
					_ff, _bfb = _dc+_ff, _da+_bfb
				}
				if _ceg != nil && _a.ToLower(_ceg.Symbol) == "\u0071" {
					_dg := _gd.Point{X: _ceg.Params[0] * _bea._cgaa, Y: _ceg.Params[1] * _bea._cgaa}
					_bcb := _gd.Point{X: _ceg.Params[2] * _bea._cgaa, Y: _ceg.Params[3] * _bea._cgaa}
					_bac := _bcb.Mul(2.0).Sub(_dg)
					_ag, _eaf = _gd.QuadraticToCubicBezier(_dc, _da, _bac.X, _bac.Y, _ff, _bfb)
				}
				_fe.Add_c(_ddf(_ag.X, 3), _ddf(_ag.Y, 3), _ddf(_eaf.X, 3), _ddf(_eaf.Y, 3), _ddf(_ff, 3), _ddf(_bfb, 3))
				_dc, _da = _ff, _bfb
			case "\u0061":
				_ad, _eag := _ee.Params[0]*_bea._cgaa, _ee.Params[1]*_bea._cgaa
				_ca := _ee.Params[2]
				_dfb := _ee.Params[3] > 0
				_ggf := _ee.Params[4] > 0
				_dcb, _baa := _ee.Params[5]*_bea._cgaa, _ee.Params[6]*_bea._cgaa
				if !_ee.isAbsolute() {
					_dcb, _baa = _dc+_dcb, _da+_baa
				}
				_dgg := _gd.EllipseToCubicBeziers(_dc, _da, _ad, _eag, _ca, _dfb, _ggf, _dcb, _baa)
				for _, _eaa := range _dgg {
					_fe.Add_c(_ddf(_eaa[1].X, 3), _ddf((_eaa[1].Y), 3), _ddf((_eaa[2].X), 3), _ddf((_eaa[2].Y), 3), _ddf((_eaa[3].X), 3), _ddf((_eaa[3].Y), 3))
				}
				_dc, _da = _dcb, _baa
			case "\u007a":
				_fe.Add_h()
			}
			_ceg = _ee
		}
	}
	if _bea.Style.FillColor != "" && _bea.Style.StrokeColor != "" {
		_fe.Add_B()
	} else if _bea.Style.FillColor != "" {
		_fe.Add_f()
	} else if _bea.Style.StrokeColor != "" {
		_fe.Add_S()
	}
	_fe.Add_h()
	_fe.Add_Q()
}

func _gaf(_bga *GraphicSVG, _ege *_fc.ContentCreator) {
	_ege.Add_q()
	_bga.Style.toContentStream(_ege)
	_dab, _aeb := _ged(_bga.Attributes["\u0078\u0031"], 64)
	if _aeb != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _aeb.Error())
	}
	_begc, _aeb := _ged(_bga.Attributes["\u0079\u0031"], 64)
	if _aeb != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _aeb.Error())
	}
	_fbg, _aeb := _ged(_bga.Attributes["\u0078\u0032"], 64)
	if _aeb != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _aeb.Error())
	}
	_bggb, _aeb := _ged(_bga.Attributes["\u0079\u0032"], 64)
	if _aeb != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _aeb.Error())
	}
	_ege.Add_m(_dab*_bga._cgaa, _begc*_bga._cgaa)
	_ege.Add_l(_fbg*_bga._cgaa, _bggb*_bga._cgaa)
	if _bga.Style.FillColor != "" && _bga.Style.StrokeColor != "" {
		_ege.Add_B()
	} else if _bga.Style.FillColor != "" {
		_ege.Add_f()
	} else if _bga.Style.StrokeColor != "" {
		_ege.Add_S()
	}
	_ege.Add_h()
	_ege.Add_Q()
}

func _bcfd(_ccba []float64) []float64 {
	for _ggdd, _ecb := 0, len(_ccba)-1; _ggdd < _ecb; _ggdd, _ecb = _ggdd+1, _ecb-1 {
		_ccba[_ggdd], _ccba[_ecb] = _ccba[_ecb], _ccba[_ggdd]
	}
	return _ccba
}

func _acb(_gde string) (*Path, error) {
	_bgd = _cdg()
	_dca, _egad := _gfeg(_gdf(_gde))
	if _egad != nil {
		return nil, _egad
	}
	return _dbaa(_dca), nil
}

func _ddf(_ffe float64, _bgde int) float64 {
	_deab := _f.Pow(10, float64(_bgde))
	return float64(_eggf(_ffe*_deab)) / _deab
}

func _edbb(_bfg string) (_fcaec, _aacf string) {
	if _bfg == "" || (_bfg[len(_bfg)-1] >= '0' && _bfg[len(_bfg)-1] <= '9') {
		return _bfg, ""
	}
	_fcaec = _bfg
	for _, _gcdf := range _abg {
		if _a.Contains(_fcaec, _gcdf) {
			_aacf = _gcdf
		}
		_fcaec = _a.TrimSuffix(_fcaec, _gcdf)
	}
	return
}

const (
	_c  = 0.72
	_fd = 28.3464
	_gg = _fd / 10
	_df = 0.551784
)

func (_cdgb *Subpath) compare(_fbf *Subpath) bool {
	if len(_cdgb.Commands) != len(_fbf.Commands) {
		return false
	}
	for _aec, _dafc := range _cdgb.Commands {
		if !_dafc.compare(_fbf.Commands[_aec]) {
			return false
		}
	}
	return true
}

func _gdf(_cea string) []token {
	var (
		_ecgf []token
		_fedg string
	)
	for _, _bacd := range _cea {
		_dba := string(_bacd)
		switch {
		case _bgd.isCommand(_dba):
			_ecgf, _fedg = _egdg(_ecgf, _fedg)
			_ecgf = append(_ecgf, token{_dba, true})
		case _dba == "\u002e":
			if _fedg == "" {
				_fedg = "\u0030"
			}
			if _a.Contains(_fedg, _dba) {
				_ecgf = append(_ecgf, token{_fedg, false})
				_fedg = "\u0030"
			}
			fallthrough
		case _dba >= "\u0030" && _dba <= "\u0039" || _dba == "\u0065":
			_fedg += _dba
		case _dba == "\u002d":
			if _a.HasSuffix(_fedg, "\u0065") {
				_fedg += _dba
			} else {
				_ecgf, _ = _egdg(_ecgf, _fedg)
				_fedg = _dba
			}
		default:
			_ecgf, _fedg = _egdg(_ecgf, _fedg)
		}
	}
	_ecgf, _ = _egdg(_ecgf, _fedg)
	return _ecgf
}

type Command struct {
	Symbol string
	Params []float64
}
type pathParserError struct{ _aaf string }

func (_edec pathParserError) Error() string { return _edec._aaf }
func _edf(_eggc string) ([]float64, error) {
	_dff := -1
	var _adge []float64
	_gfa := ' '
	for _fcaf, _fbcg := range _eggc {
		if !_d.IsNumber(_fbcg) && _fbcg != '.' && !(_fbcg == '-' && _gfa == 'e') && _fbcg != 'e' {
			if _dff != -1 {
				_cdb, _bgeg := _dde(_eggc[_dff:_fcaf])
				if _bgeg != nil {
					return _adge, _bgeg
				}
				_adge = append(_adge, _cdb...)
			}
			if _fbcg == '-' {
				_dff = _fcaf
			} else {
				_dff = -1
			}
		} else if _dff == -1 {
			_dff = _fcaf
		}
		_gfa = _fbcg
	}
	if _dff != -1 && _dff != len(_eggc) {
		_eef, _cef := _dde(_eggc[_dff:])
		if _cef != nil {
			return _adge, _cef
		}
		_adge = append(_adge, _eef...)
	}
	return _adge, nil
}

func (_bba *GraphicSVG) toContentStream(_cbcb *_fc.ContentCreator) {
	_adf, _deec := _aag(_bba.Attributes, _bba._cgaa)
	if _deec != nil {
		_be.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _deec)
	}
	_bba.Style = _adf
	switch _bba.Name {
	case "\u0070\u0061\u0074\u0068":
		_cc(_bba, _cbcb)
		for _, _adg := range _bba.Children {
			_adg.toContentStream(_cbcb)
		}
	case "\u0072\u0065\u0063\u0074":
		_afa(_bba, _cbcb)
		for _, _bgaf := range _bba.Children {
			_bgaf.toContentStream(_cbcb)
		}
	case "\u0063\u0069\u0072\u0063\u006c\u0065":
		_gbe(_bba, _cbcb)
		for _, _bacb := range _bba.Children {
			_bacb.toContentStream(_cbcb)
		}
	case "\u0065l\u006c\u0069\u0070\u0073\u0065":
		_bbc(_bba, _cbcb)
		for _, _bacc := range _bba.Children {
			_bacc.toContentStream(_cbcb)
		}
	case "\u0070\u006f\u006c\u0079\u006c\u0069\u006e\u0065":
		_ffd(_bba, _cbcb)
		for _, _fca := range _bba.Children {
			_fca.toContentStream(_cbcb)
		}
	case "\u0070o\u006c\u0079\u0067\u006f\u006e":
		_ffb(_bba, _cbcb)
		for _, _faa := range _bba.Children {
			_faa.toContentStream(_cbcb)
		}
	case "\u006c\u0069\u006e\u0065":
		_gaf(_bba, _cbcb)
		for _, _aa := range _bba.Children {
			_aa.toContentStream(_cbcb)
		}
	case "\u0067":
		_dea, _gcdb := _bba.Attributes["\u0066\u0069\u006c\u006c"]
		_efd, _fcae := _bba.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"]
		_dac, _dfge := _bba.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]
		for _, _gbec := range _bba.Children {
			if _, _begb := _gbec.Attributes["\u0066\u0069\u006c\u006c"]; !_begb && _gcdb {
				_gbec.Attributes["\u0066\u0069\u006c\u006c"] = _dea
			}
			if _, _cf := _gbec.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"]; !_cf && _fcae {
				_gbec.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"] = _efd
			}
			if _, _eedgc := _gbec.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]; !_eedgc && _dfge {
				_gbec.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"] = _dac
			}
			_gbec.toContentStream(_cbcb)
		}
	}
}

type GraphicSVG struct {
	ViewBox    struct{ X, Y, W, H float64 }
	Name       string
	Attributes map[string]string
	Children   []*GraphicSVG
	Content    string
	Style      *GraphicSVGStyle
	Width      float64
	Height     float64
	_cgaa      float64
}

func _aag(_dae map[string]string, _ecec float64) (*GraphicSVGStyle, error) {
	_eaab := _dacc()
	_efb, _dda := _dae["\u0066\u0069\u006c\u006c"]
	if _dda {
		_eaab.FillColor = _efb
		if _efb == "\u006e\u006f\u006e\u0065" {
			_eaab.FillColor = ""
		}
	}
	_geac, _bdd := _dae["\u0073\u0074\u0072\u006f\u006b\u0065"]
	if _bdd {
		_eaab.StrokeColor = _geac
		if _geac == "\u006e\u006f\u006e\u0065" {
			_eaab.StrokeColor = ""
		}
	}
	_adc, _gfc := _dae["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]
	if _gfc {
		_gee, _bedg := _ged(_adc, 64)
		if _bedg != nil {
			return nil, _bedg
		}
		_eaab.StrokeWidth = _gee * _ecec
	}
	return _eaab, nil
}

func ParseFromFile(path string) (*GraphicSVG, error) {
	_ccb, _fga := _ac.Open(path)
	if _fga != nil {
		return nil, _fga
	}
	defer _ccb.Close()
	return ParseFromStream(_ccb)
}

func (_fa *GraphicSVG) Decode(decoder *_bg.Decoder) error {
	for {
		_fce, _fgg := decoder.Token()
		if _fce == nil && _fgg == _g.EOF {
			break
		}
		if _fgg != nil {
			return _fgg
		}
		switch _cbf := _fce.(type) {
		case _bg.StartElement:
			_fcb := _cfb(_cbf)
			_acg := _fcb.Decode(decoder)
			if _acg != nil {
				return _acg
			}
			_fa.Children = append(_fa.Children, _fcb)
		case _bg.CharData:
			_egd := _a.TrimSpace(string(_cbf))
			if _egd != "" {
				_fa.Content = string(_cbf)
			}
		case _bg.EndElement:
			if _cbf.Name.Local == _fa.Name {
				return nil
			}
		}
	}
	return nil
}
func (_abfc *Command) isAbsolute() bool { return _abfc.Symbol == _a.ToUpper(_abfc.Symbol) }

type Path struct{ Subpaths []*Subpath }

func (_bbf *commands) isCommand(_eecd string) bool {
	for _, _ada := range _bbf._daad {
		if _a.ToLower(_eecd) == _ada {
			return true
		}
	}
	return false
}

func (_feb *Command) compare(_daf *Command) bool {
	if _feb.Symbol != _daf.Symbol {
		return false
	}
	for _bbcc, _fgdg := range _feb.Params {
		if _fgdg != _daf.Params[_bbcc] {
			return false
		}
	}
	return true
}

type commands struct {
	_daad []string
	_cad  map[string]int
	_cce  string
	_gfcf string
}

func _egdg(_fbb []token, _ebf string) ([]token, string) {
	if _ebf != "" {
		_fbb = append(_fbb, token{_ebf, false})
		_ebf = ""
	}
	return _fbb, _ebf
}

func _cdg() commands {
	_cag := map[string]int{"\u006d": 2, "\u007a": 0, "\u006c": 2, "\u0068": 1, "\u0076": 1, "\u0063": 6, "\u0073": 4, "\u0071": 4, "\u0074": 2, "\u0061": 7}
	var _dbf []string
	for _bad := range _cag {
		_dbf = append(_dbf, _bad)
	}
	return commands{_dbf, _cag, "\u006d", "\u007a"}
}

func _dbaa(_ebdf []*Command) *Path {
	_afdc := &Path{}
	var _ggda []*Command
	for _bcad, _cbed := range _ebdf {
		switch _a.ToLower(_cbed.Symbol) {
		case _bgd._cce:
			if len(_ggda) > 0 {
				_afdc.Subpaths = append(_afdc.Subpaths, &Subpath{_ggda})
			}
			_ggda = []*Command{_cbed}
		case _bgd._gfcf:
			_ggda = append(_ggda, _cbed)
			_afdc.Subpaths = append(_afdc.Subpaths, &Subpath{_ggda})
			_ggda = []*Command{}
		default:
			_ggda = append(_ggda, _cbed)
			if len(_ebdf) == _bcad+1 {
				_afdc.Subpaths = append(_afdc.Subpaths, &Subpath{_ggda})
			}
		}
	}
	return _afdc
}

var (
	_abg = []string{"\u0063\u006d", "\u006d\u006d", "\u0070\u0078", "\u0070\u0074"}
	_ce  = map[string]float64{"\u0063\u006d": _fd, "\u006d\u006d": _gg, "\u0070\u0078": _c, "\u0070\u0074": 1}
)

type Subpath struct{ Commands []*Command }

func _afa(_fed *GraphicSVG, _bag *_fc.ContentCreator) {
	_bag.Add_q()
	_fed.Style.toContentStream(_bag)
	_gcf, _eec := _ged(_fed.Attributes["\u0078"], 64)
	if _eec != nil {
		_be.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _eec.Error())
	}
	_ffg, _eec := _ged(_fed.Attributes["\u0079"], 64)
	if _eec != nil {
		_be.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _eec.Error())
	}
	_eba, _eec := _ged(_fed.Attributes["\u0077\u0069\u0064t\u0068"], 64)
	if _eec != nil {
		_be.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _eec.Error())
	}
	_fcc, _eec := _ged(_fed.Attributes["\u0068\u0065\u0069\u0067\u0068\u0074"], 64)
	if _eec != nil {
		_be.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077h\u0069\u006c\u0065 \u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _eec.Error())
	}
	_bag.Add_re(_gcf*_fed._cgaa, _ffg*_fed._cgaa, _eba*_fed._cgaa, _fcc*_fed._cgaa)
	if _fed.Style.FillColor != "" && _fed.Style.StrokeColor != "" {
		_bag.Add_B()
	} else if _fed.Style.FillColor != "" {
		_bag.Add_f()
	} else if _fed.Style.StrokeColor != "" {
		_bag.Add_S()
	}
	_bag.Add_Q()
}

func _ffd(_efg *GraphicSVG, _cggb *_fc.ContentCreator) {
	_cggb.Add_q()
	_efg.Style.toContentStream(_cggb)
	_bgg, _beg := _edf(_efg.Attributes["\u0070\u006f\u0069\u006e\u0074\u0073"])
	if _beg != nil {
		_be.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076", _beg)
		return
	}
	if len(_bgg)%2 > 0 {
		_be.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068")
		return
	}
	for _bcf := 0; _bcf < len(_bgg); {
		if _bcf == 0 {
			_cggb.Add_m(_bgg[_bcf]*_efg._cgaa, _bgg[_bcf+1]*_efg._cgaa)
		} else {
			_cggb.Add_l(_bgg[_bcf]*_efg._cgaa, _bgg[_bcf+1]*_efg._cgaa)
		}
		_bcf += 2
	}
	if _efg.Style.FillColor != "" && _efg.Style.StrokeColor != "" {
		_cggb.Add_B()
	} else if _efg.Style.FillColor != "" {
		_cggb.Add_f()
	} else if _efg.Style.StrokeColor != "" {
		_cggb.Add_S()
	}
	_cggb.Add_h()
	_cggb.Add_Q()
}

func (_cbc *GraphicSVG) ToContentCreator(cc *_fc.ContentCreator, scaleX, scaleY, translateX, translateY float64) *_fc.ContentCreator {
	if _cbc.Name == "\u0073\u0076\u0067" {
		_cbc.SetScaling(scaleX, scaleY)
		cc.Add_cm(1, 0, 0, 1, translateX, translateY)
		_cbc.setDefaultScaling(_cbc._cgaa)
		cc.Add_q()
		_bgf := _f.Max(scaleX, scaleY)
		cc.Add_re(_cbc.ViewBox.X*_bgf, _cbc.ViewBox.Y*_bgf, _cbc.ViewBox.W*_bgf, _cbc.ViewBox.H*_bgf)
		cc.Add_W()
		cc.Add_n()
		for _, _ec := range _cbc.Children {
			_ec.ViewBox = _cbc.ViewBox
			_ec.toContentStream(cc)
		}
		cc.Add_Q()
		return cc
	}
	return nil
}

func _aee(_fac string) (_dbd, _dabb, _bdeg float64) {
	if (len(_fac) != 4 && len(_fac) != 7) || _fac[0] != '#' {
		_be.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _fac)
		return _dbd, _dabb, _bdeg
	}
	var _def, _gbd, _cfd int
	if len(_fac) == 4 {
		var _dcf, _fbcc, _ggfg int
		_fgdgb, _feff := _b.Sscanf(_fac, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_dcf, &_fbcc, &_ggfg)
		if _feff != nil {
			_be.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _fac, _feff)
			return _dbd, _dabb, _bdeg
		}
		if _fgdgb != 3 {
			_be.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _fac)
			return _dbd, _dabb, _bdeg
		}
		_def = _dcf*16 + _dcf
		_gbd = _fbcc*16 + _fbcc
		_cfd = _ggfg*16 + _ggfg
	} else {
		_agcg, _fdf := _b.Sscanf(_fac, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_def, &_gbd, &_cfd)
		if _fdf != nil {
			_be.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _fac)
			return _dbd, _dabb, _bdeg
		}
		if _agcg != 3 {
			_be.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", _fac, _agcg)
			return _dbd, _dabb, _bdeg
		}
	}
	_fcaa := float64(_def) / 255.0
	_ffbe := float64(_gbd) / 255.0
	_gaa := float64(_cfd) / 255.0
	return _fcaa, _ffbe, _gaa
}

func ParseFromStream(source _g.Reader) (*GraphicSVG, error) {
	_dcdf := _bg.NewDecoder(source)
	_dcdf.CharsetReader = _ab.NewReaderLabel
	_cedd, _egf := _agb(_dcdf)
	if _egf != nil {
		return nil, _egf
	}
	if _babg := _cedd.Decode(_dcdf); _babg != nil && _babg != _g.EOF {
		return nil, _babg
	}
	return _cedd, nil
}

func _ged(_gdg string, _gddg int) (float64, error) {
	_ebgc, _abb := _edbb(_gdg)
	_eeef, _bbb := _af.ParseFloat(_ebgc, _gddg)
	if _bbb != nil {
		return 0, _bbb
	}
	if _deg, _fcaef := _ce[_abb]; _fcaef {
		_eeef = _eeef * _deg
	} else {
		_eeef = _eeef * _c
	}
	return _eeef, nil
}

func _gbe(_abf *GraphicSVG, _eedg *_fc.ContentCreator) {
	_eedg.Add_q()
	_abf.Style.toContentStream(_eedg)
	_ed, _dfg := _ged(_abf.Attributes["\u0063\u0078"], 64)
	if _dfg != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _dfg.Error())
	}
	_fb, _dfg := _ged(_abf.Attributes["\u0063\u0079"], 64)
	if _dfg != nil {
		_be.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _dfg.Error())
	}
	_fge, _dfg := _ged(_abf.Attributes["\u0072"], 64)
	if _dfg != nil {
		_be.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0072\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _dfg.Error())
	}
	_dd := _fge * _abf._cgaa
	_gfe := _fge * _abf._cgaa
	_feg := _dd * _df
	_agc := _gfe * _df
	_dge := _gf.NewCubicBezierPath()
	_dge = _dge.AppendCurve(_gf.NewCubicBezierCurve(-_dd, 0, -_dd, _agc, -_feg, _gfe, 0, _gfe))
	_dge = _dge.AppendCurve(_gf.NewCubicBezierCurve(0, _gfe, _feg, _gfe, _dd, _agc, _dd, 0))
	_dge = _dge.AppendCurve(_gf.NewCubicBezierCurve(_dd, 0, _dd, -_agc, _feg, -_gfe, 0, -_gfe))
	_dge = _dge.AppendCurve(_gf.NewCubicBezierCurve(0, -_gfe, -_feg, -_gfe, -_dd, -_agc, -_dd, 0))
	_dge = _dge.Offset(_ed*_abf._cgaa, _fb*_abf._cgaa)
	if _abf.Style.StrokeWidth > 0 {
		_dge = _dge.Offset(_abf.Style.StrokeWidth/2, _abf.Style.StrokeWidth/2)
	}
	_gf.DrawBezierPathWithCreator(_dge, _eedg)
	if _abf.Style.FillColor != "" && _abf.Style.StrokeColor != "" {
		_eedg.Add_B()
	} else if _abf.Style.FillColor != "" {
		_eedg.Add_f()
	} else if _abf.Style.StrokeColor != "" {
		_eedg.Add_S()
	}
	_eedg.Add_h()
	_eedg.Add_Q()
}
