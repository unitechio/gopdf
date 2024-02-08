package svg

import (
	_gb "encoding/xml"
	_e "fmt"
	_a "io"
	_eb "math"
	_ca "os"
	_gg "strconv"
	_g "strings"
	_c "unicode"

	_ggd "bitbucket.org/shenghui0779/gopdf/common"
	_be "bitbucket.org/shenghui0779/gopdf/contentstream"
	_f "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_af "bitbucket.org/shenghui0779/gopdf/internal/graphic2d"
	_d "golang.org/x/net/html/charset"
)

func (_ddb *GraphicSVG) SetScaling(xFactor, yFactor float64) {
	_defa := _ddb.Width / _ddb.ViewBox.W
	_edef := _ddb.Height / _ddb.ViewBox.H
	_ddb.setDefaultScaling(_eb.Max(_defa, _edef))
	for _, _afd := range _ddb.Children {
		_afd.SetScaling(xFactor, yFactor)
	}
}
func (_fac *GraphicSVG) Decode(decoder *_gb.Decoder) error {
	for {
		_gaeb, _aac := decoder.Token()
		if _gaeb == nil && _aac == _a.EOF {
			break
		}
		if _aac != nil {
			return _aac
		}
		switch _bbfg := _gaeb.(type) {
		case _gb.StartElement:
			_efa := _bc(_bbfg)
			_ffe := _efa.Decode(decoder)
			if _ffe != nil {
				return _ffe
			}
			_fac.Children = append(_fac.Children, _efa)
		case _gb.CharData:
			_edd := _g.TrimSpace(string(_bbfg))
			if _edd != "" {
				_fac.Content = string(_bbfg)
			}
		case _gb.EndElement:
			if _bbfg.Name.Local == _fac.Name {
				return nil
			}
		}
	}
	return nil
}
func _cbf() commands {
	var _cccb = map[string]int{"\u006d": 2, "\u007a": 0, "\u006c": 2, "\u0068": 1, "\u0076": 1, "\u0063": 6, "\u0073": 4, "\u0071": 4, "\u0074": 2, "\u0061": 7}
	var _gfd []string
	for _fadc := range _cccb {
		_gfd = append(_gfd, _fadc)
	}
	return commands{_gfd, _cccb, "\u006d", "\u007a"}
}
func _ebd(_aega map[string]string, _ecc float64) (*GraphicSVGStyle, error) {
	_ege := _efaa()
	_dgf, _dfa := _aega["\u0066\u0069\u006c\u006c"]
	if _dfa {
		_ege.FillColor = _dgf
		if _dgf == "\u006e\u006f\u006e\u0065" {
			_ege.FillColor = ""
		}
	}
	_dead, _aae := _aega["\u0073\u0074\u0072\u006f\u006b\u0065"]
	if _aae {
		_ege.StrokeColor = _dead
		if _dead == "\u006e\u006f\u006e\u0065" {
			_ege.StrokeColor = ""
		}
	}
	_dab, _cef := _aega["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]
	if _cef {
		_bcd, _gebd := _eccc(_dab, 64)
		if _gebd != nil {
			return nil, _gebd
		}
		_ege.StrokeWidth = _bcd * _ecc
	}
	return _ege, nil
}

type Path struct{ Subpaths []*Subpath }

func (_cgbe *Path) compare(_adef *Path) bool {
	if len(_cgbe.Subpaths) != len(_adef.Subpaths) {
		return false
	}
	for _ggc, _efg := range _cgbe.Subpaths {
		if !_efg.compare(_adef.Subpaths[_ggc]) {
			return false
		}
	}
	return true
}
func (_efaf *Subpath) compare(_fcab *Subpath) bool {
	if len(_efaf.Commands) != len(_fcab.Commands) {
		return false
	}
	for _aaa, _fec := range _efaf.Commands {
		if !_fec.compare(_fcab.Commands[_aaa]) {
			return false
		}
	}
	return true
}
func _cc(_ag *GraphicSVG, _fb *_be.ContentCreator) {
	_fb.Add_q()
	_ag.Style.toContentStream(_fb)
	_eg, _ac := _ffeagg(_ag.Attributes["\u0064"])
	if _ac != nil {
		_ggd.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025s", _ac.Error())
	}
	var (
		_cd, _agg = 0.0, 0.0
		_df, _gbb = 0.0, 0.0
		_bb       *Command
	)
	for _, _bba := range _eg.Subpaths {
		for _, _dd := range _bba.Commands {
			switch _g.ToLower(_dd.Symbol) {
			case "\u006d":
				_df, _gbb = _dd.Params[0]*_ag._edc, _dd.Params[1]*_ag._edc
				if !_dd.isAbsolute() {
					_df, _gbb = _cd+_df-_ag.ViewBox.X, _agg+_gbb-_ag.ViewBox.Y
				}
				_fb.Add_m(_cfaa(_df, 3), _cfaa(_gbb, 3))
				_cd, _agg = _df, _gbb
			case "\u0063":
				_ad, _ba, _afg, _da, _dae, _gba := _dd.Params[0]*_ag._edc, _dd.Params[1]*_ag._edc, _dd.Params[2]*_ag._edc, _dd.Params[3]*_ag._edc, _dd.Params[4]*_ag._edc, _dd.Params[5]*_ag._edc
				if !_dd.isAbsolute() {
					_ad, _ba, _afg, _da, _dae, _gba = _cd+_ad, _agg+_ba, _cd+_afg, _agg+_da, _cd+_dae, _agg+_gba
				}
				_fb.Add_c(_cfaa(_ad, 3), _cfaa(_ba, 3), _cfaa(_afg, 3), _cfaa(_da, 3), _cfaa(_dae, 3), _cfaa(_gba, 3))
				_cd, _agg = _dae, _gba
			case "\u0073":
				_gd, _ebg, _aa, _eda := _dd.Params[0]*_ag._edc, _dd.Params[1]*_ag._edc, _dd.Params[2]*_ag._edc, _dd.Params[3]*_ag._edc
				if !_dd.isAbsolute() {
					_gd, _ebg, _aa, _eda = _cd+_gd, _agg+_ebg, _cd+_aa, _agg+_eda
				}
				_fb.Add_c(_cfaa(_cd, 3), _cfaa(_agg, 3), _cfaa(_gd, 3), _cfaa(_ebg, 3), _cfaa(_aa, 3), _cfaa(_eda, 3))
				_cd, _agg = _aa, _eda
			case "\u006c":
				_cf, _cba := _dd.Params[0]*_ag._edc, _dd.Params[1]*_ag._edc
				if !_dd.isAbsolute() {
					_cf, _cba = _cd+_cf, _agg+_cba
				}
				_fb.Add_l(_cfaa(_cf, 3), _cfaa(_cba, 3))
				_cd, _agg = _cf, _cba
			case "\u0068":
				_fa := _dd.Params[0] * _ag._edc
				if !_dd.isAbsolute() {
					_fa = _cd + _fa
				}
				_fb.Add_l(_cfaa(_fa, 3), _cfaa(_agg, 3))
				_cd = _fa
			case "\u0076":
				_bgc := _dd.Params[0] * _ag._edc
				if !_dd.isAbsolute() {
					_bgc = _agg + _bgc
				}
				_fb.Add_l(_cfaa(_cd, 3), _cfaa(_bgc, 3))
				_agg = _bgc
			case "\u0071":
				_fg, _bga, _ae, _cfb := _dd.Params[0]*_ag._edc, _dd.Params[1]*_ag._edc, _dd.Params[2]*_ag._edc, _dd.Params[3]*_ag._edc
				if !_dd.isAbsolute() {
					_fg, _bga, _ae, _cfb = _cd+_fg, _agg+_bga, _cd+_ae, _agg+_cfb
				}
				_afb, _ga := _af.QuadraticToCubicBezier(_cd, _agg, _fg, _bga, _ae, _cfb)
				_fb.Add_c(_cfaa(_afb.X, 3), _cfaa(_afb.Y, 3), _cfaa(_ga.X, 3), _cfaa(_ga.Y, 3), _cfaa(_ae, 3), _cfaa(_cfb, 3))
				_cd, _agg = _ae, _cfb
			case "\u0074":
				var _baf, _ee _af.Point
				_ff, _eae := _dd.Params[0]*_ag._edc, _dd.Params[1]*_ag._edc
				if !_dd.isAbsolute() {
					_ff, _eae = _cd+_ff, _agg+_eae
				}
				if _bb != nil && _g.ToLower(_bb.Symbol) == "\u0071" {
					_ef := _af.Point{X: _bb.Params[0] * _ag._edc, Y: _bb.Params[1] * _ag._edc}
					_db := _af.Point{X: _bb.Params[2] * _ag._edc, Y: _bb.Params[3] * _ag._edc}
					_daf := _db.Mul(2.0).Sub(_ef)
					_baf, _ee = _af.QuadraticToCubicBezier(_cd, _agg, _daf.X, _daf.Y, _ff, _eae)
				}
				_fb.Add_c(_cfaa(_baf.X, 3), _cfaa(_baf.Y, 3), _cfaa(_ee.X, 3), _cfaa(_ee.Y, 3), _cfaa(_ff, 3), _cfaa(_eae, 3))
				_cd, _agg = _ff, _eae
			case "\u0061":
				_dbb, _agf := _dd.Params[0]*_ag._edc, _dd.Params[1]*_ag._edc
				_ede := _dd.Params[2]
				_ada := _dd.Params[3] > 0
				_afgg := _dd.Params[4] > 0
				_gea, _aee := _dd.Params[5]*_ag._edc, _dd.Params[6]*_ag._edc
				if !_dd.isAbsolute() {
					_gea, _aee = _cd+_gea, _agg+_aee
				}
				_cdb := _af.EllipseToCubicBeziers(_cd, _agg, _dbb, _agf, _ede, _ada, _afgg, _gea, _aee)
				for _, _fc := range _cdb {
					_fb.Add_c(_cfaa(_fc[1].X, 3), _cfaa((_fc[1].Y), 3), _cfaa((_fc[2].X), 3), _cfaa((_fc[2].Y), 3), _cfaa((_fc[3].X), 3), _cfaa((_fc[3].Y), 3))
				}
				_cd, _agg = _gea, _aee
			case "\u007a":
				_fb.Add_h()
			}
			_bb = _dd
		}
	}
	if _ag.Style.FillColor != "" && _ag.Style.StrokeColor != "" {
		_fb.Add_B()
	} else if _ag.Style.FillColor != "" {
		_fb.Add_f()
	} else if _ag.Style.StrokeColor != "" {
		_fb.Add_S()
	}
	_fb.Add_h()
	_fb.Add_Q()
}
func _aab(_dg *GraphicSVG, _fe *_be.ContentCreator) {
	_fe.Add_q()
	_dg.Style.toContentStream(_fe)
	_cgb, _bf := _eabd(_dg.Attributes["\u0070\u006f\u0069\u006e\u0074\u0073"])
	if _bf != nil {
		_ggd.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076", _bf)
		return
	}
	if len(_cgb)%2 > 0 {
		_ggd.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068")
		return
	}
	for _ffa := 0; _ffa < len(_cgb); {
		if _ffa == 0 {
			_fe.Add_m(_cgb[_ffa]*_dg._edc, _cgb[_ffa+1]*_dg._edc)
		} else {
			_fe.Add_l(_cgb[_ffa]*_dg._edc, _cgb[_ffa+1]*_dg._edc)
		}
		_ffa += 2
	}
	_fe.Add_l(_cgb[0]*_dg._edc, _cgb[1]*_dg._edc)
	if _dg.Style.FillColor != "" && _dg.Style.StrokeColor != "" {
		_fe.Add_B()
	} else if _dg.Style.FillColor != "" {
		_fe.Add_f()
	} else if _dg.Style.StrokeColor != "" {
		_fe.Add_S()
	}
	_fe.Add_h()
	_fe.Add_Q()
}
func _caab(_bbab []*Command) *Path {
	_agbb := &Path{}
	var _eefa []*Command
	for _ccbc, _ccgg := range _bbab {
		switch _g.ToLower(_ccgg.Symbol) {
		case _eca._bgcb:
			if len(_eefa) > 0 {
				_agbb.Subpaths = append(_agbb.Subpaths, &Subpath{_eefa})
			}
			_eefa = []*Command{_ccgg}
		case _eca._eeca:
			_eefa = append(_eefa, _ccgg)
			_agbb.Subpaths = append(_agbb.Subpaths, &Subpath{_eefa})
			_eefa = []*Command{}
		default:
			_eefa = append(_eefa, _ccgg)
			if len(_bbab) == _ccbc+1 {
				_agbb.Subpaths = append(_agbb.Subpaths, &Subpath{_eefa})
			}
		}
	}
	return _agbb
}
func _dcba(_faaa string) (_aeda, _dfge string) {
	if _faaa == "" || (_faaa[len(_faaa)-1] >= '0' && _faaa[len(_faaa)-1] <= '9') {
		return _faaa, ""
	}
	_aeda = _faaa
	for _, _bab := range _ea {
		if _g.Contains(_aeda, _bab) {
			_dfge = _bab
		}
		_aeda = _g.TrimSuffix(_aeda, _bab)
	}
	return
}

const (
	_ec  = 0.72
	_ebe = 28.3464
	_ge  = _ebe / 10
	_ed  = 0.551784
)

func _eccc(_gbg string, _eefg int) (float64, error) {
	_aedbf, _bfcg := _dcba(_gbg)
	_ccda, _cdbg := _gg.ParseFloat(_aedbf, _eefg)
	if _cdbg != nil {
		return 0, _cdbg
	}
	if _bdd, _dbbb := _de[_bfcg]; _dbbb {
		_ccda = _ccda * _bdd
	} else {
		_ccda = _ccda * _ec
	}
	return _ccda, nil
}
func (_defag *Command) isAbsolute() bool { return _defag.Symbol == _g.ToUpper(_defag.Symbol) }
func _eabd(_bfef string) ([]float64, error) {
	_eeaa := -1
	var _aagd []float64
	_abg := ' '
	for _beg, _aff := range _bfef {
		if !_c.IsNumber(_aff) && _aff != '.' && !(_aff == '-' && _abg == 'e') && _aff != 'e' {
			if _eeaa != -1 {
				_efe, _feg := _eecd(_bfef[_eeaa:_beg])
				if _feg != nil {
					return _aagd, _feg
				}
				_aagd = append(_aagd, _efe...)
			}
			if _aff == '-' {
				_eeaa = _beg
			} else {
				_eeaa = -1
			}
		} else if _eeaa == -1 {
			_eeaa = _beg
		}
		_abg = _aff
	}
	if _eeaa != -1 && _eeaa != len(_bfef) {
		_dabe, _ecb := _eecd(_bfef[_eeaa:])
		if _ecb != nil {
			return _aagd, _ecb
		}
		_aagd = append(_aagd, _dabe...)
	}
	return _aagd, nil
}

type GraphicSVGStyle struct {
	FillColor   string
	StrokeColor string
	StrokeWidth float64
}

func (_bdg *GraphicSVGStyle) toContentStream(_gdg *_be.ContentCreator) {
	if _bdg == nil {
		return
	}
	if _bdg.FillColor != "" {
		var _cec, _ffb, _fcfb float64
		if _ebf, _dcb := _af.ColorMap[_bdg.FillColor]; _dcb {
			_cgg, _aadd, _cbbb, _ := _ebf.RGBA()
			_cec, _ffb, _fcfb = float64(_cgg), float64(_aadd), float64(_cbbb)
		} else {
			_cec, _ffb, _fcfb = _gbbe(_bdg.FillColor)
		}
		_gdg.Add_rg(_cec, _ffb, _fcfb)
	}
	if _bdg.StrokeColor != "" {
		var _fcd, _ddg, _fcdb float64
		if _fgf, _acd := _af.ColorMap[_bdg.StrokeColor]; _acd {
			_gcdc, _gdf, _efab, _ := _fgf.RGBA()
			_fcd, _ddg, _fcdb = float64(_gcdc)/255.0, float64(_gdf)/255.0, float64(_efab)/255.0
		} else {
			_fcd, _ddg, _fcdb = _gbbe(_bdg.StrokeColor)
		}
		_gdg.Add_RG(_fcd, _ddg, _fcdb)
	}
	if _bdg.StrokeWidth > 0 {
		_gdg.Add_w(_bdg.StrokeWidth)
	}
}
func _efaa() *GraphicSVGStyle {
	return &GraphicSVGStyle{FillColor: "\u00230\u0030\u0030\u0030\u0030\u0030", StrokeColor: "", StrokeWidth: 0}
}
func (_cad *GraphicSVG) setDefaultScaling(_dee float64) {
	_cad._edc = _dee
	if _cad.Style != nil && _cad.Style.StrokeWidth > 0 {
		_cad.Style.StrokeWidth = _cad.Style.StrokeWidth * _cad._edc
	}
	for _, _eec := range _cad.Children {
		_eec.setDefaultScaling(_dee)
	}
}
func _eea(_ceg *_gb.Decoder) (*GraphicSVG, error) {
	for {
		_cdc, _ead := _ceg.Token()
		if _cdc == nil && _ead == _a.EOF {
			break
		}
		if _ead != nil {
			return nil, _ead
		}
		switch _ade := _cdc.(type) {
		case _gb.StartElement:
			return _bc(_ade), nil
		}
	}
	return &GraphicSVG{}, nil
}

type commands struct {
	_fff  []string
	_cfe  map[string]int
	_bgcb string
	_eeca string
}

func (_gfc *Command) compare(_afbe *Command) bool {
	if _gfc.Symbol != _afbe.Symbol {
		return false
	}
	for _dbd, _eadb := range _gfc.Params {
		if _eadb != _afbe.Params[_dbd] {
			return false
		}
	}
	return true
}
func (_fbd *commands) isCommand(_adc string) bool {
	for _, _acf := range _fbd._fff {
		if _g.ToLower(_adc) == _acf {
			return true
		}
	}
	return false
}

type Subpath struct{ Commands []*Command }

func _ffeagg(_dff string) (*Path, error) {
	_eca = _cbf()
	_dfgf, _agfa := _fcb(_fefa(_dff))
	if _agfa != nil {
		return nil, _agfa
	}
	return _caab(_dfgf), nil
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
	_edc       float64
}

func _gbbe(_aada string) (_aca, _fba, _efaff float64) {
	if (len(_aada) != 4 && len(_aada) != 7) || _aada[0] != '#' {
		_ggd.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _aada)
		return _aca, _fba, _efaff
	}
	var _eeb, _add, _ffeac int
	if len(_aada) == 4 {
		var _fbbb, _bbbg, _abe int
		_abd, _bcc := _e.Sscanf(_aada, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_fbbb, &_bbbg, &_abe)
		if _bcc != nil {
			_ggd.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _aada, _bcc)
			return _aca, _fba, _efaff
		}
		if _abd != 3 {
			_ggd.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _aada)
			return _aca, _fba, _efaff
		}
		_eeb = _fbbb*16 + _fbbb
		_add = _bbbg*16 + _bbbg
		_ffeac = _abe*16 + _abe
	} else {
		_gdfe, _afc := _e.Sscanf(_aada, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_eeb, &_add, &_ffeac)
		if _afc != nil {
			_ggd.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", _aada)
			return _aca, _fba, _efaff
		}
		if _gdfe != 3 {
			_ggd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", _aada, _gdfe)
			return _aca, _fba, _efaff
		}
	}
	_ecgb := float64(_eeb) / 255.0
	_acda := float64(_add) / 255.0
	_fegd := float64(_ffeac) / 255.0
	return _ecgb, _acda, _fegd
}
func _ccgd(_dfe []float64) []float64 {
	for _aaba, _egc := 0, len(_dfe)-1; _aaba < _egc; _aaba, _egc = _aaba+1, _egc-1 {
		_dfe[_aaba], _dfe[_egc] = _dfe[_egc], _dfe[_aaba]
	}
	return _dfe
}
func (_bfa pathParserError) Error() string { return _bfa._cea }

var (
	_ea = []string{"\u0063\u006d", "\u006d\u006d", "\u0070\u0078", "\u0070\u0074"}
	_de = map[string]float64{"\u0063\u006d": _ebe, "\u006d\u006d": _ge, "\u0070\u0078": _ec, "\u0070\u0074": 1}
)

func _ged(_bfac float64) int { return int(_bfac + _eb.Copysign(0.5, _bfac)) }
func _fefa(_egcd string) []token {
	var (
		_gcda []token
		_fge  string
	)
	for _, _fbe := range _egcd {
		_cabb := string(_fbe)
		switch {
		case _eca.isCommand(_cabb):
			_gcda, _fge = _fef(_gcda, _fge)
			_gcda = append(_gcda, token{_cabb, true})
		case _cabb == "\u002e":
			if _fge == "" {
				_fge = "\u0030"
			}
			if _g.Contains(_fge, _cabb) {
				_gcda = append(_gcda, token{_fge, false})
				_fge = "\u0030"
			}
			fallthrough
		case _cabb >= "\u0030" && _cabb <= "\u0039" || _cabb == "\u0065":
			_fge += _cabb
		case _cabb == "\u002d":
			if _g.HasSuffix(_fge, "\u0065") {
				_fge += _cabb
			} else {
				_gcda, _ = _fef(_gcda, _fge)
				_fge = _cabb
			}
		default:
			_gcda, _fge = _fef(_gcda, _fge)
		}
	}
	_gcda, _ = _fef(_gcda, _fge)
	return _gcda
}
func (_dgd *GraphicSVG) ToContentCreator(cc *_be.ContentCreator, scaleX, scaleY, translateX, translateY float64) *_be.ContentCreator {
	if _dgd.Name == "\u0073\u0076\u0067" {
		_dgd.SetScaling(scaleX, scaleY)
		cc.Add_cm(1, 0, 0, 1, translateX, translateY)
		_dgd.setDefaultScaling(_dgd._edc)
		cc.Add_q()
		_dafc := _eb.Max(scaleX, scaleY)
		cc.Add_re(_dgd.ViewBox.X*_dafc, _dgd.ViewBox.Y*_dafc, _dgd.ViewBox.W*_dafc, _dgd.ViewBox.H*_dafc)
		cc.Add_W()
		cc.Add_n()
		for _, _aeg := range _dgd.Children {
			_aeg.ViewBox = _dgd.ViewBox
			_aeg.toContentStream(cc)
		}
		cc.Add_Q()
		return cc
	}
	return nil
}
func _cfaa(_dbbf float64, _bbac int) float64 {
	_dad := _eb.Pow(10, float64(_bbac))
	return float64(_ged(_dbbf*_dad)) / _dad
}
func _geaf(_gdd *GraphicSVG, _gef *_be.ContentCreator) {
	_gef.Add_q()
	_gdd.Style.toContentStream(_gef)
	_cab, _fd := _eccc(_gdd.Attributes["\u0063\u0078"], 64)
	if _fd != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _fd.Error())
	}
	_ggb, _fd := _eccc(_gdd.Attributes["\u0063\u0079"], 64)
	if _fd != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _fd.Error())
	}
	_cbaa, _fd := _eccc(_gdd.Attributes["\u0072"], 64)
	if _fd != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0072\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _fd.Error())
	}
	_ddc := _cbaa * _gdd._edc
	_gc := _cbaa * _gdd._edc
	_fcf := _ddc * _ed
	_caa := _gc * _ed
	_ddd := _f.NewCubicBezierPath()
	_ddd = _ddd.AppendCurve(_f.NewCubicBezierCurve(-_ddc, 0, -_ddc, _caa, -_fcf, _gc, 0, _gc))
	_ddd = _ddd.AppendCurve(_f.NewCubicBezierCurve(0, _gc, _fcf, _gc, _ddc, _caa, _ddc, 0))
	_ddd = _ddd.AppendCurve(_f.NewCubicBezierCurve(_ddc, 0, _ddc, -_caa, _fcf, -_gc, 0, -_gc))
	_ddd = _ddd.AppendCurve(_f.NewCubicBezierCurve(0, -_gc, -_fcf, -_gc, -_ddc, -_caa, -_ddc, 0))
	_ddd = _ddd.Offset(_cab*_gdd._edc, _ggb*_gdd._edc)
	if _gdd.Style.StrokeWidth > 0 {
		_ddd = _ddd.Offset(_gdd.Style.StrokeWidth/2, _gdd.Style.StrokeWidth/2)
	}
	_f.DrawBezierPathWithCreator(_ddd, _gef)
	if _gdd.Style.FillColor != "" && _gdd.Style.StrokeColor != "" {
		_gef.Add_B()
	} else if _gdd.Style.FillColor != "" {
		_gef.Add_f()
	} else if _gdd.Style.StrokeColor != "" {
		_gef.Add_S()
	}
	_gef.Add_h()
	_gef.Add_Q()
}
func (_ccc *GraphicSVG) toContentStream(_agb *_be.ContentCreator) {
	_def, _aed := _ebd(_ccc.Attributes, _ccc._edc)
	if _aed != nil {
		_ggd.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _aed)
	}
	_ccc.Style = _def
	switch _ccc.Name {
	case "\u0070\u0061\u0074\u0068":
		_cc(_ccc, _agb)
		for _, _dea := range _ccc.Children {
			_dea.toContentStream(_agb)
		}
	case "\u0072\u0065\u0063\u0074":
		_ce(_ccc, _agb)
		for _, _abf := range _ccc.Children {
			_abf.toContentStream(_agb)
		}
	case "\u0063\u0069\u0072\u0063\u006c\u0065":
		_geaf(_ccc, _agb)
		for _, _gga := range _ccc.Children {
			_gga.toContentStream(_agb)
		}
	case "\u0065l\u006c\u0069\u0070\u0073\u0065":
		_bd(_ccc, _agb)
		for _, _eab := range _ccc.Children {
			_eab.toContentStream(_agb)
		}
	case "\u0070\u006f\u006c\u0079\u006c\u0069\u006e\u0065":
		_gab(_ccc, _agb)
		for _, _fgc := range _ccc.Children {
			_fgc.toContentStream(_agb)
		}
	case "\u0070o\u006c\u0079\u0067\u006f\u006e":
		_aab(_ccc, _agb)
		for _, _cff := range _ccc.Children {
			_cff.toContentStream(_agb)
		}
	case "\u006c\u0069\u006e\u0065":
		_gf(_ccc, _agb)
		for _, _dbeg := range _ccc.Children {
			_dbeg.toContentStream(_agb)
		}
	case "\u0067":
		_fdb, _bad := _ccc.Attributes["\u0066\u0069\u006c\u006c"]
		_bbb, _cge := _ccc.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"]
		_aeb, _fbg := _ccc.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]
		for _, _bag := range _ccc.Children {
			if _, _dde := _bag.Attributes["\u0066\u0069\u006c\u006c"]; !_dde && _bad {
				_bag.Attributes["\u0066\u0069\u006c\u006c"] = _fdb
			}
			if _, _cfc := _bag.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"]; !_cfc && _cge {
				_bag.Attributes["\u0073\u0074\u0072\u006f\u006b\u0065"] = _bbb
			}
			if _, _dbg := _bag.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]; !_dbg && _fbg {
				_bag.Attributes["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"] = _aeb
			}
			_bag.toContentStream(_agb)
		}
	}
}
func _eecd(_aebf string) (_acfc []float64, _fdc error) {
	var _gagag float64
	_bfee := 0
	_dgfg := true
	for _edec, _afdb := range _aebf {
		if _afdb == '.' {
			if _dgfg {
				_dgfg = false
				continue
			}
			_gagag, _fdc = _eccc(_aebf[_bfee:_edec], 64)
			if _fdc != nil {
				return
			}
			_acfc = append(_acfc, _gagag)
			_bfee = _edec
		}
	}
	_gagag, _fdc = _eccc(_aebf[_bfee:], 64)
	if _fdc != nil {
		return
	}
	_acfc = append(_acfc, _gagag)
	return
}
func _bd(_fcg *GraphicSVG, _dba *_be.ContentCreator) {
	_dba.Add_q()
	_fcg.Style.toContentStream(_dba)
	_cbb, _fdd := _eccc(_fcg.Attributes["\u0063\u0078"], 64)
	if _fdd != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _fdd.Error())
	}
	_dc, _fdd := _eccc(_fcg.Attributes["\u0063\u0079"], 64)
	if _fdd != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _fdd.Error())
	}
	_ccg, _fdd := _eccc(_fcg.Attributes["\u0072\u0078"], 64)
	if _fdd != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _fdd.Error())
	}
	_egb, _fdd := _eccc(_fcg.Attributes["\u0072\u0079"], 64)
	if _fdd != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _fdd.Error())
	}
	_egf := _ccg * _fcg._edc
	_acb := _egb * _fcg._edc
	_gag := _cbb * _fcg._edc
	_bac := _dc * _fcg._edc
	_bgg := _egf * _ed
	_gcd := _acb * _ed
	_gae := _f.NewCubicBezierPath()
	_gae = _gae.AppendCurve(_f.NewCubicBezierCurve(-_egf, 0, -_egf, _gcd, -_bgg, _acb, 0, _acb))
	_gae = _gae.AppendCurve(_f.NewCubicBezierCurve(0, _acb, _bgg, _acb, _egf, _gcd, _egf, 0))
	_gae = _gae.AppendCurve(_f.NewCubicBezierCurve(_egf, 0, _egf, -_gcd, _bgg, -_acb, 0, -_acb))
	_gae = _gae.AppendCurve(_f.NewCubicBezierCurve(0, -_acb, -_bgg, -_acb, -_egf, -_gcd, -_egf, 0))
	_gae = _gae.Offset(_gag, _bac)
	if _fcg.Style.StrokeWidth > 0 {
		_gae = _gae.Offset(_fcg.Style.StrokeWidth/2, _fcg.Style.StrokeWidth/2)
	}
	_f.DrawBezierPathWithCreator(_gae, _dba)
	if _fcg.Style.FillColor != "" && _fcg.Style.StrokeColor != "" {
		_dba.Add_B()
	} else if _fcg.Style.FillColor != "" {
		_dba.Add_f()
	} else if _fcg.Style.StrokeColor != "" {
		_dba.Add_S()
	}
	_dba.Add_h()
	_dba.Add_Q()
}
func ParseFromFile(path string) (*GraphicSVG, error) {
	_cac, _ccgf := _ca.Open(path)
	if _ccgf != nil {
		return nil, _ccgf
	}
	defer _cac.Close()
	return ParseFromStream(_cac)
}

type token struct {
	_ffea string
	_dcf  bool
}

func ParseFromStream(source _a.Reader) (*GraphicSVG, error) {
	_afda := _gb.NewDecoder(source)
	_afda.CharsetReader = _d.NewReaderLabel
	_ced, _eee := _eea(_afda)
	if _eee != nil {
		return nil, _eee
	}
	if _aaf := _ced.Decode(_afda); _aaf != nil && _aaf != _a.EOF {
		return nil, _aaf
	}
	return _ced, nil
}

type pathParserError struct{ _cea string }

var _eca commands

func _gab(_aag *GraphicSVG, _ccd *_be.ContentCreator) {
	_ccd.Add_q()
	_aag.Style.toContentStream(_ccd)
	_ded, _cdd := _eabd(_aag.Attributes["\u0070\u006f\u0069\u006e\u0074\u0073"])
	if _cdd != nil {
		_ggd.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076", _cdd)
		return
	}
	if len(_ded)%2 > 0 {
		_ggd.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068")
		return
	}
	for _fga := 0; _fga < len(_ded); {
		if _fga == 0 {
			_ccd.Add_m(_ded[_fga]*_aag._edc, _ded[_fga+1]*_aag._edc)
		} else {
			_ccd.Add_l(_ded[_fga]*_aag._edc, _ded[_fga+1]*_aag._edc)
		}
		_fga += 2
	}
	if _aag.Style.FillColor != "" && _aag.Style.StrokeColor != "" {
		_ccd.Add_B()
	} else if _aag.Style.FillColor != "" {
		_ccd.Add_f()
	} else if _aag.Style.StrokeColor != "" {
		_ccd.Add_S()
	}
	_ccd.Add_h()
	_ccd.Add_Q()
}
func _gf(_bbf *GraphicSVG, _egg *_be.ContentCreator) {
	_egg.Add_q()
	_bbf.Style.toContentStream(_egg)
	_eeg, _geb := _eccc(_bbf.Attributes["\u0078\u0031"], 64)
	if _geb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _geb.Error())
	}
	_ab, _geb := _eccc(_bbf.Attributes["\u0079\u0031"], 64)
	if _geb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _geb.Error())
	}
	_eag, _geb := _eccc(_bbf.Attributes["\u0078\u0032"], 64)
	if _geb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _geb.Error())
	}
	_aeee, _geb := _eccc(_bbf.Attributes["\u0079\u0032"], 64)
	if _geb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _geb.Error())
	}
	_egg.Add_m(_eeg*_bbf._edc, _ab*_bbf._edc)
	_egg.Add_l(_eag*_bbf._edc, _aeee*_bbf._edc)
	if _bbf.Style.FillColor != "" && _bbf.Style.StrokeColor != "" {
		_egg.Add_B()
	} else if _bbf.Style.FillColor != "" {
		_egg.Add_f()
	} else if _bbf.Style.StrokeColor != "" {
		_egg.Add_S()
	}
	_egg.Add_h()
	_egg.Add_Q()
}

type Command struct {
	Symbol string
	Params []float64
}

func _fcb(_gcc []token) ([]*Command, error) {
	var (
		_eac []*Command
		_cbg []float64
	)
	for _bdc := len(_gcc) - 1; _bdc >= 0; _bdc-- {
		_cae := _gcc[_bdc]
		if _cae._dcf {
			_gaga := _eca._cfe[_g.ToLower(_cae._ffea)]
			_aggb := len(_cbg)
			if _gaga == 0 && _aggb == 0 {
				_dcc := &Command{Symbol: _cae._ffea}
				_eac = append([]*Command{_dcc}, _eac...)
			} else if _gaga != 0 && _aggb%_gaga == 0 {
				_ffeag := _aggb / _gaga
				for _bfg := 0; _bfg < _ffeag; _bfg++ {
					_dfb := _cae._ffea
					if _dfb == "\u006d" && _bfg < _ffeag-1 {
						_dfb = "\u006c"
					}
					if _dfb == "\u004d" && _bfg < _ffeag-1 {
						_dfb = "\u004c"
					}
					_defd := &Command{_dfb, _ccgd(_cbg[:_gaga])}
					_eac = append([]*Command{_defd}, _eac...)
					_cbg = _cbg[_gaga:]
				}
			} else {
				_cbfb := pathParserError{"I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006fr\u0020" + _cae._ffea}
				return nil, _cbfb
			}
		} else {
			_ceb, _cbe := _eccc(_cae._ffea, 64)
			if _cbe != nil {
				return nil, _cbe
			}
			_cbg = append(_cbg, _ceb)
		}
	}
	return _eac, nil
}
func _ce(_fca *GraphicSVG, _cca *_be.ContentCreator) {
	_cca.Add_q()
	_fca.Style.toContentStream(_cca)
	_dbe, _ccb := _eccc(_fca.Attributes["\u0078"], 64)
	if _ccb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _ccb.Error())
	}
	_bbd, _ccb := _eccc(_fca.Attributes["\u0079"], 64)
	if _ccb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _ccb.Error())
	}
	_cg, _ccb := _eccc(_fca.Attributes["\u0077\u0069\u0064t\u0068"], 64)
	if _ccb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _ccb.Error())
	}
	_cfg, _ccb := _eccc(_fca.Attributes["\u0068\u0065\u0069\u0067\u0068\u0074"], 64)
	if _ccb != nil {
		_ggd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0077h\u0069\u006c\u0065 \u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _ccb.Error())
	}
	_cca.Add_re(_dbe*_fca._edc, _bbd*_fca._edc, _cg*_fca._edc, _cfg*_fca._edc)
	if _fca.Style.FillColor != "" && _fca.Style.StrokeColor != "" {
		_cca.Add_B()
	} else if _fca.Style.FillColor != "" {
		_cca.Add_f()
	} else if _fca.Style.StrokeColor != "" {
		_cca.Add_S()
	}
	_cca.Add_Q()
}
func _fef(_ddf []token, _ccag string) ([]token, string) {
	if _ccag != "" {
		_ddf = append(_ddf, token{_ccag, false})
		_ccag = ""
	}
	return _ddf, _ccag
}
func ParseFromString(svgStr string) (*GraphicSVG, error) {
	return ParseFromStream(_g.NewReader(svgStr))
}
func _bc(_faa _gb.StartElement) *GraphicSVG {
	_bed := &GraphicSVG{}
	_edf := make(map[string]string)
	for _, _cee := range _faa.Attr {
		_edf[_cee.Name.Local] = _cee.Value
	}
	_bed.Name = _faa.Name.Local
	_bed.Attributes = _edf
	_bed._edc = 1
	if _bed.Name == "\u0073\u0076\u0067" {
		_bfe, _fad := _eabd(_edf["\u0076i\u0065\u0077\u0042\u006f\u0078"])
		if _fad != nil {
			_ggd.Log.Debug("\u0055\u006ea\u0062\u006c\u0065\u0020t\u006f\u0020p\u0061\u0072\u0073\u0065\u0020\u0076\u0069\u0065w\u0042\u006f\u0078\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074e\u003a\u0020\u0025\u0076", _fad)
			return nil
		}
		_bed.ViewBox.X = _bfe[0]
		_bed.ViewBox.Y = _bfe[1]
		_bed.ViewBox.W = _bfe[2]
		_bed.ViewBox.H = _bfe[3]
		_bed.Width = _bed.ViewBox.W
		_bed.Height = _bed.ViewBox.H
		if _dfg, _aad := _edf["\u0077\u0069\u0064t\u0068"]; _aad {
			_afge, _gca := _eccc(_dfg, 64)
			if _gca != nil {
				_ggd.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0077\u0069\u0064\u0074\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _gca)
				return nil
			}
			_bed.Width = _afge
		}
		if _aeeb, _fde := _edf["\u0068\u0065\u0069\u0067\u0068\u0074"]; _fde {
			_ecg, _dbf := _eccc(_aeeb, 64)
			if _dbf != nil {
				_ggd.Log.Debug("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0070\u0061rse\u0020he\u0069\u0067\u0068\u0074\u0020\u0061\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076", _dbf)
				return nil
			}
			_bed.Height = _ecg
		}
		if _bed.Width > 0 && _bed.Height > 0 {
			_bed._edc = _bed.Width / _bed.ViewBox.W
		}
	}
	return _bed
}
