package redactor

import (
	_f "errors"
	_af "fmt"
	_e "io"
	_d "regexp"
	_c "sort"
	_g "strings"

	_gg "bitbucket.org/shenghui0779/gopdf/common"
	_eb "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ef "bitbucket.org/shenghui0779/gopdf/core"
	_ge "bitbucket.org/shenghui0779/gopdf/creator"
	_gf "bitbucket.org/shenghui0779/gopdf/extractor"
	_dd "bitbucket.org/shenghui0779/gopdf/model"
)

func _cga(_fddb *_gf.TextMarkArray) []*_gf.TextMarkArray {
	_afae := _fddb.Elements()
	_ddcf := len(_afae)
	var _fag _ef.PdfObject
	_ffed := []*_gf.TextMarkArray{}
	_fdeg := &_gf.TextMarkArray{}
	_fgbc := -1
	for _fbdda, _aed := range _afae {
		_edf := _aed.DirectObject
		_fgbc = _aed.Index
		if _edf == nil {
			_cagb := _cef(_fddb, _fbdda, _fgbc)
			if _fag != nil {
				if _cagb == -1 || _cagb > _fbdda {
					_ffed = append(_ffed, _fdeg)
					_fdeg = &_gf.TextMarkArray{}
				}
			}
		} else if _edf != nil && _fag == nil {
			if _fgbc == 0 && _fbdda > 0 {
				_ffed = append(_ffed, _fdeg)
				_fdeg = &_gf.TextMarkArray{}
			}
		} else if _edf != nil && _fag != nil {
			if _edf != _fag {
				_ffed = append(_ffed, _fdeg)
				_fdeg = &_gf.TextMarkArray{}
			}
		}
		_fag = _edf
		_fdeg.Append(_aed)
		if _fbdda == (_ddcf - 1) {
			_ffed = append(_ffed, _fdeg)
		}
	}
	return _ffed
}

type replacement struct {
	_ggd string
	_bc  float64
	_db  int
}

func _fb(_fg *_gf.TextMarkArray) int {
	_fa := 0
	_ee := _fg.Elements()
	if _ee[0].Text == "\u0020" {
		_fa++
	}
	if _ee[_fg.Len()-1].Text == "\u0020" {
		_fa++
	}
	return _fa
}

type regexMatcher struct{ _dfdg RedactionTerm }

// RedactionTerm holds the regexp pattern and the replacement string for the redaction process.
type RedactionTerm struct{ Pattern *_d.Regexp }

// Redact executes the redact operation on a pdf file and updates the content streams of all pages of the file.
func (_fee *Redactor) Redact() error {
	_fceg, _cdaa := _fee._gee.GetNumPages()
	if _cdaa != nil {
		return _af.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0074\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020P\u0061\u0067\u0065\u0073")
	}
	_eeb := _fee._gcg.FillColor
	_fda := _fee._gcg.BorderWidth
	_gbff := _fee._gcg.FillOpacity
	for _gdac := 1; _gdac <= _fceg; _gdac++ {
		_eege, _ega := _fee._gee.GetPage(_gdac)
		if _ega != nil {
			return _ega
		}
		_geg, _ega := _gf.New(_eege)
		if _ega != nil {
			return _ega
		}
		_ced, _, _, _ega := _geg.ExtractPageText()
		if _ega != nil {
			return _ega
		}
		_bff := _ced.GetContentStreamOps()
		_fbg, _afe, _ega := _fee.redactPage(_bff, _eege.Resources)
		if _afe == nil {
			_gg.Log.Info("N\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u0020t\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065d \u0070\u0061\u0074t\u0061r\u006e\u002e")
			_afe = _bff
		}
		_fgfe := _eb.ContentStreamOperation{Operand: "\u006e"}
		*_afe = append(*_afe, &_fgfe)
		_eege.SetContentStreams([]string{_afe.String()}, _ef.NewFlateEncoder())
		if _ega != nil {
			return _ega
		}
		_dgeb, _ega := _eege.GetMediaBox()
		if _ega != nil {
			return _ega
		}
		if _eege.MediaBox == nil {
			_eege.MediaBox = _dgeb
		}
		if _ccad := _fee._afbb.AddPage(_eege); _ccad != nil {
			return _ccad
		}
		_c.Slice(_fbg, func(_gce, _ebe int) bool { return _fbg[_gce]._gfc < _fbg[_ebe]._gfc })
		_abf := _dgeb.Ury
		for _, _dbf := range _fbg {
			_geb := _dbf._ddd
			_edgd := _fee._afbb.NewRectangle(_geb.Llx, _abf-_geb.Lly, _geb.Urx-_geb.Llx, -(_geb.Ury - _geb.Lly))
			_edgd.SetFillColor(_eeb)
			_edgd.SetBorderWidth(_fda)
			_edgd.SetFillOpacity(_gbff)
			if _ffa := _fee._afbb.Draw(_edgd); _ffa != nil {
				return nil
			}
		}
	}
	_fee._afbb.SetOutlineTree(_fee._gee.GetOutlineTree())
	return nil
}
func _ce(_eea *_dd.PdfFont, _fd _gf.TextMark) float64 {
	_ffe := 0.001
	_feg := _fd.Th / 100
	if _eea.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_ffe = 1
	}
	_ebdb, _fdg := _eea.GetRuneMetrics(' ')
	if !_fdg {
		_ebdb, _fdg = _eea.GetCharMetrics(32)
	}
	if !_fdg {
		_ebdb, _ = _dd.DefaultFont().GetRuneMetrics(' ')
	}
	_bcf := _ffe * ((_ebdb.Wx*_fd.FontSize + _fd.Tc + _fd.Tw) / _feg)
	return _bcf
}
func _deb(_bbgd localSpanMarks, _dac *_gf.TextMarkArray, _bad *_dd.PdfFont, _aae, _aeec string) ([]_ef.PdfObject, error) {
	_ade := _cgce(_dac)
	Tj, _cfa := _dgef(_dac)
	if _cfa != nil {
		return nil, _cfa
	}
	_egbd := len(_aae)
	_egfd := len(_ade)
	_cgeb := -1
	_fea := _ef.MakeFloat(Tj)
	if _ade != _aeec {
		_fdde := _bbgd._ede
		if _fdde == 0 {
			_cgeb = _g.LastIndex(_aae, _ade)
		} else {
			_cgeb = _g.Index(_aae, _ade)
		}
	} else {
		_cgeb = _g.Index(_aae, _ade)
	}
	_bae := _cgeb + _egfd
	_dgg := []_ef.PdfObject{}
	if _cgeb == 0 && _bae == _egbd {
		_dgg = append(_dgg, _fea)
	} else if _cgeb == 0 && _bae < _egbd {
		_adg := _beb(_aae[_bae:], _bad)
		_eef := _ef.MakeStringFromBytes(_adg)
		_dgg = append(_dgg, _fea, _eef)
	} else if _cgeb > 0 && _bae >= _egbd {
		_dce := _beb(_aae[:_cgeb], _bad)
		_ga := _ef.MakeStringFromBytes(_dce)
		_dgg = append(_dgg, _ga, _fea)
	} else if _cgeb > 0 && _bae < _egbd {
		_gff := _beb(_aae[:_cgeb], _bad)
		_gae := _beb(_aae[_bae:], _bad)
		_bdga := _ef.MakeStringFromBytes(_gff)
		_acb := _ef.MakeString(string(_gae))
		_dgg = append(_dgg, _bdga, _fea, _acb)
	}
	return _dgg, nil
}
func _cd(_abgc *_gf.TextMarkArray) (_ef.PdfObject, int) {
	var _df _ef.PdfObject
	_ged := -1
	for _dadd, _be := range _abgc.Elements() {
		_df = _be.DirectObject
		_ged = _dadd
		if _df != nil {
			break
		}
	}
	return _df, _ged
}
func _dg(_gc *_eb.ContentStreamOperation, _bef _ef.PdfObject, _de []localSpanMarks) error {
	_bb, _befd := _ef.GetArray(_gc.Params[0])
	_gbd := []_ef.PdfObject{}
	if !_befd {
		_gg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _gc)
		return _af.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _gc)
	}
	_ff, _dgf := _cf(_de)
	if len(_dgf) == 1 {
		_ada := _dgf[0]
		_bbf := _ff[_ada]
		if len(_bbf) == 1 {
			_gga := _bbf[0]
			_cba := _gga._edc
			_ggdb := _fca(_cba)
			_egf, _ebf := _accd(_bef, _ggdb)
			if _ebf != nil {
				return _ebf
			}
			_bbe, _ebf := _deb(_gga, _cba, _ggdb, _egf, _ada)
			if _ebf != nil {
				return _ebf
			}
			for _, _bbg := range _bb.Elements() {
				if _bbg == _bef {
					_gbd = append(_gbd, _bbe...)
				} else {
					_gbd = append(_gbd, _bbg)
				}
			}
		} else {
			_eae := _bbf[0]._edc
			_dbe := _fca(_eae)
			_ag, _eaeb := _accd(_bef, _dbe)
			if _eaeb != nil {
				return _eaeb
			}
			_fe, _eaeb := _ecdc(_ag, _bbf)
			if _eaeb != nil {
				return _eaeb
			}
			_ace := _egg(_fe)
			_bg := _cag(_ag, _ace, _dbe)
			for _, _gd := range _bb.Elements() {
				if _gd == _bef {
					_gbd = append(_gbd, _bg...)
				} else {
					_gbd = append(_gbd, _gd)
				}
			}
		}
		_gc.Params[0] = _ef.MakeArray(_gbd...)
	} else if len(_dgf) > 1 {
		_fge := _de[0]
		_ecf := _fge._edc
		_, _ebd := _cd(_ecf)
		_daf := _ecf.Elements()[_ebd]
		_ece := _daf.Font
		_eac, _ddg := _accd(_bef, _ece)
		if _ddg != nil {
			return _ddg
		}
		_dab, _ddg := _ecdc(_eac, _de)
		if _ddg != nil {
			return _ddg
		}
		_bd := _egg(_dab)
		_fc := _cag(_eac, _bd, _ece)
		for _, _bf := range _bb.Elements() {
			if _bf == _bef {
				_gbd = append(_gbd, _fc...)
			} else {
				_gbd = append(_gbd, _bf)
			}
		}
		_gc.Params[0] = _ef.MakeArray(_gbd...)
	}
	return nil
}
func _eff(_dbg, _cac, _decg float64) float64 {
	_decg = _decg / 100
	_cddc := (-1000 * _dbg) / (_cac * _decg)
	return _cddc
}

type matchedIndex struct {
	_faa int
	_gaf int
	_fab string
}

func _abg(_dad *_eb.ContentStreamOperations, _cb string, _ba int) error {
	_egb := _eb.ContentStreamOperations{}
	var _gb _eb.ContentStreamOperation
	for _aee, _ebc := range *_dad {
		if _aee == _ba {
			if _cb == "\u0027" {
				_ecd := _eb.ContentStreamOperation{Operand: "\u0054\u002a"}
				_egb = append(_egb, &_ecd)
				_gb.Params = _ebc.Params
				_gb.Operand = "\u0054\u006a"
				_egb = append(_egb, &_gb)
			} else if _cb == "\u0022" {
				_gfg := _ebc.Params[:2]
				Tc, Tw := _gfg[0], _gfg[1]
				_ad := _eb.ContentStreamOperation{Params: []_ef.PdfObject{Tc}, Operand: "\u0054\u0063"}
				_egb = append(_egb, &_ad)
				_ad = _eb.ContentStreamOperation{Params: []_ef.PdfObject{Tw}, Operand: "\u0054\u0077"}
				_egb = append(_egb, &_ad)
				_gb.Params = []_ef.PdfObject{_ebc.Params[2]}
				_gb.Operand = "\u0054\u006a"
				_egb = append(_egb, &_gb)
			}
		}
		_egb = append(_egb, _ebc)
	}
	*_dad = _egb
	return nil
}
func _bab(_ddc *_eb.ContentStreamOperation, _abc _ef.PdfObject, _bdg []localSpanMarks) error {
	var _fdd *_ef.PdfObjectArray
	_ege, _gcd := _cf(_bdg)
	if len(_gcd) == 1 {
		_bdd := _gcd[0]
		_gbdf := _ege[_bdd]
		if len(_gbdf) == 1 {
			_bcc := _gbdf[0]
			_cge := _bcc._edc
			_gbf := _fca(_cge)
			_dcc, _ebg := _accd(_abc, _gbf)
			if _ebg != nil {
				return _ebg
			}
			_acg, _ebg := _deb(_bcc, _cge, _gbf, _dcc, _bdd)
			if _ebg != nil {
				return _ebg
			}
			_fdd = _ef.MakeArray(_acg...)
		} else {
			_aea := _gbdf[0]._edc
			_acd := _fca(_aea)
			_dcb, _fec := _accd(_abc, _acd)
			if _fec != nil {
				return _fec
			}
			_bdc, _fec := _ecdc(_dcb, _gbdf)
			if _fec != nil {
				return _fec
			}
			_gfb := _egg(_bdc)
			_fbe := _cag(_dcb, _gfb, _acd)
			_fdd = _ef.MakeArray(_fbe...)
		}
	} else if len(_gcd) > 1 {
		_cdd := _bdg[0]
		_aab := _cdd._edc
		_, _bfa := _cd(_aab)
		_dfb := _aab.Elements()[_bfa]
		_cged := _dfb.Font
		_gcf, _gef := _accd(_abc, _cged)
		if _gef != nil {
			return _gef
		}
		_cda, _gef := _ecdc(_gcf, _bdg)
		if _gef != nil {
			return _gef
		}
		_gec := _egg(_cda)
		_fgf := _cag(_gcf, _gec, _cged)
		_fdd = _ef.MakeArray(_fgf...)
	}
	_ddc.Params[0] = _fdd
	_ddc.Operand = "\u0054\u004a"
	return nil
}

// Redactor represents a Redactor object.
type Redactor struct {
	_gee  *_dd.PdfReader
	_gffa *RedactionOptions
	_afbb *_ge.Creator
	_gcg  *RectangleProps
}

func _acbg(_cddg []*matchedIndex, _egfc [][]int) []*matchedIndex {
	_gdgb := []*matchedIndex{}
	for _, _gbg := range _cddg {
		_cbaa, _age := _efb(_gbg, _egfc)
		if _cbaa {
			_aeefe := _cgef(_gbg, _age)
			_gdgb = append(_gdgb, _aeefe...)
		} else {
			_gdgb = append(_gdgb, _gbg)
		}
	}
	return _gdgb
}
func _accd(_bbc _ef.PdfObject, _dgd *_dd.PdfFont) (string, error) {
	_dgdf, _bgcd := _ef.GetStringBytes(_bbc)
	if !_bgcd {
		return "", _ef.ErrTypeError
	}
	_dcea := _dgd.BytesToCharcodes(_dgdf)
	_aeff, _fega, _aff := _dgd.CharcodesToStrings(_dcea)
	if _aff > 0 {
		_gg.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _fega, _aff)
	}
	_dage := _g.Join(_aeff, "")
	return _dage, nil
}
func _daa(_cgf int, _fga []int) bool {
	for _, _abca := range _fga {
		if _abca == _cgf {
			return true
		}
	}
	return false
}
func (_eccc *Redactor) redactPage(_eca *_eb.ContentStreamOperations, _dagf *_dd.PdfPageResources) ([]matchedBBox, *_eb.ContentStreamOperations, error) {
	_adb, _gegf := _gf.NewFromContents(_eca.String(), _dagf)
	if _gegf != nil {
		return nil, nil, _gegf
	}
	_bca, _, _, _gegf := _adb.ExtractPageText()
	if _gegf != nil {
		return nil, nil, _gegf
	}
	_eca = _bca.GetContentStreamOps()
	_debd := _bca.Marks()
	_ffc := _bca.Text()
	_ffc, _dba := _aafa(_ffc)
	_adgf := []matchedBBox{}
	_fgb := make(map[_ef.PdfObject][]localSpanMarks)
	_aaa := []*targetMap{}
	for _, _dabg := range _eccc._gffa.Terms {
		_bfg, _eda := _bfb(_dabg)
		if _eda != nil {
			return nil, nil, _eda
		}
		_cab, _eda := _bfg.match(_ffc)
		if _eda != nil {
			return nil, nil, _eda
		}
		_cab = _acbg(_cab, _dba)
		_aec := _gdd(_cab)
		_aaa = append(_aaa, _aec...)
	}
	_ffd(_aaa)
	for _, _egad := range _aaa {
		_cff := _egad._gea
		_daca := _egad._aeef
		_aaeb := []matchedBBox{}
		for _, _dcba := range _daca {
			_acgb, _dacb, _add := _gfba(_dcba, _debd, _cff)
			if _add != nil {
				return nil, nil, _add
			}
			_cccf := _cga(_acgb)
			for _dabc, _gceb := range _cccf {
				_abcf := localSpanMarks{_edc: _gceb, _ede: _dabc, _fffa: _cff}
				_cggd, _ := _cd(_gceb)
				if _bgce, _gdga := _fgb[_cggd]; _gdga {
					_fgb[_cggd] = append(_bgce, _abcf)
				} else {
					_fgb[_cggd] = []localSpanMarks{_abcf}
				}
			}
			_aaeb = append(_aaeb, _dacb)
		}
		_adgf = append(_adgf, _aaeb...)
	}
	_gegf = _cg(_eca, _fgb)
	if _gegf != nil {
		return nil, nil, _gegf
	}
	return _adgf, _eca, nil
}
func _ffd(_ccae []*targetMap) {
	for _aafd, _feeg := range _ccae {
		for _bcfb, _dcd := range _ccae {
			if _aafd != _bcfb {
				_afa, _aaef := _fgaa(*_feeg, *_dcd)
				if _afa {
					_dace(_dcd, _aaef)
				}
			}
		}
	}
}
func _efb(_eeba *matchedIndex, _caf [][]int) (bool, [][]int) {
	_aedg := [][]int{}
	for _, _caba := range _caf {
		if _eeba._faa < _caba[0] && _eeba._gaf > _caba[1] {
			_aedg = append(_aedg, _caba)
		}
	}
	return len(_aedg) > 0, _aedg
}

// RectangleProps defines properties of the redaction rectangle to be drawn.
type RectangleProps struct {
	FillColor   _ge.Color
	BorderWidth float64
	FillOpacity float64
}
type localSpanMarks struct {
	_edc  *_gf.TextMarkArray
	_ede  int
	_fffa string
}

// RedactionOptions is a collection of RedactionTerm objects.
type RedactionOptions struct{ Terms []RedactionTerm }

func _beb(_eaea string, _cgg *_dd.PdfFont) []byte {
	_fba, _gdg := _cgg.StringToCharcodeBytes(_eaea)
	if _gdg != 0 {
		_gg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0073\u006fm\u0065\u0020\u0072un\u0065\u0073\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065d\u002e\u000a\u0009\u0025\u0073\u0020\u002d\u003e \u0025\u0076", _eaea, _fba)
	}
	return _fba
}

// WriteToFile writes the redacted document to file specified by `outputPath`.
func (_fgbb *Redactor) WriteToFile(outputPath string) error {
	if _fabb := _fgbb._afbb.WriteToFile(outputPath); _fabb != nil {
		return _af.Errorf("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0077\u0072\u0069\u0074\u0065\u0020t\u0068\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0066\u0069\u006c\u0065")
	}
	return nil
}
func _dgef(_feab *_gf.TextMarkArray) (float64, error) {
	_cbf, _aba := _feab.BBox()
	if !_aba {
		return 0.0, _af.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _feab)
	}
	_cad := _fb(_feab)
	_dfg := 0.0
	_, _bea := _cd(_feab)
	_cca := _feab.Elements()[_bea]
	_fae := _cca.Font
	if _cad > 0 {
		_dfg = _ce(_fae, _cca)
	}
	_ecc := (_cbf.Urx - _cbf.Llx)
	_ecc = _ecc + _dfg*float64(_cad)
	Tj := _eff(_ecc, _cca.FontSize, _cca.Th)
	return Tj, nil
}
func _bga(_eaa, _aaf string) []int {
	if len(_aaf) == 0 {
		return nil
	}
	var _bcb []int
	for _bcbg := 0; _bcbg < len(_eaa); {
		_dbee := _g.Index(_eaa[_bcbg:], _aaf)
		if _dbee < 0 {
			return _bcb
		}
		_bcb = append(_bcb, _bcbg+_dbee)
		_bcbg += _dbee + len(_aaf)
	}
	return _bcb
}

type matchedBBox struct {
	_ddd _dd.PdfRectangle
	_gfc string
}

func _dace(_fed *targetMap, _egadb []int) {
	var _bbea [][]int
	for _dfd, _eggf := range _fed._aeef {
		if _daa(_dfd, _egadb) {
			continue
		}
		_bbea = append(_bbea, _eggf)
	}
	_fed._aeef = _bbea
}
func _cgef(_fdc *matchedIndex, _abbe [][]int) []*matchedIndex {
	_aefff := []*matchedIndex{}
	_caed := _fdc._faa
	_bafe := _caed
	_dfe := _fdc._fab
	_ggg := 0
	for _, _dgfde := range _abbe {
		_abgb := _dgfde[0] - _caed
		if _ggg >= _abgb {
			continue
		}
		_daff := _dfe[_ggg:_abgb]
		_fcec := &matchedIndex{_fab: _daff, _faa: _bafe, _gaf: _dgfde[0]}
		if len(_g.TrimSpace(_daff)) != 0 {
			_aefff = append(_aefff, _fcec)
		}
		_ggg = _dgfde[1] - _caed
		_bafe = _caed + _ggg
	}
	_fdgf := _dfe[_ggg:]
	_bebc := &matchedIndex{_fab: _fdgf, _faa: _bafe, _gaf: _fdc._gaf}
	if len(_g.TrimSpace(_fdgf)) != 0 {
		_aefff = append(_aefff, _bebc)
	}
	return _aefff
}
func _cag(_fde string, _egc []replacement, _beg *_dd.PdfFont) []_ef.PdfObject {
	_eacd := []_ef.PdfObject{}
	_geff := 0
	_bee := _fde
	for _ggc, _eeg := range _egc {
		_cdc := _eeg._db
		_efa := _eeg._bc
		_fbac := _eeg._ggd
		_gbc := _ef.MakeFloat(_efa)
		if _geff > _cdc || _cdc == -1 {
			continue
		}
		_eggg := _fde[_geff:_cdc]
		_efag := _beb(_eggg, _beg)
		_deca := _ef.MakeStringFromBytes(_efag)
		_eacd = append(_eacd, _deca)
		_eacd = append(_eacd, _gbc)
		_agc := _cdc + len(_fbac)
		_bee = _fde[_agc:]
		_geff = _agc
		if _ggc == len(_egc)-1 {
			_efag = _beb(_bee, _beg)
			_deca = _ef.MakeStringFromBytes(_efag)
			_eacd = append(_eacd, _deca)
		}
	}
	return _eacd
}
func (_dbed *regexMatcher) match(_bbcc string) ([]*matchedIndex, error) {
	_bbda := _dbed._dfdg.Pattern
	if _bbda == nil {
		return nil, _f.New("\u006e\u006f\u0020\u0070at\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006d\u0070\u0069\u006c\u0065\u0064")
	}
	var (
		_agf  = _bbda.FindAllStringIndex(_bbcc, -1)
		_fddf []*matchedIndex
	)
	for _, _ggfb := range _agf {
		_fddf = append(_fddf, &matchedIndex{_faa: _ggfb[0], _gaf: _ggfb[1], _fab: _bbcc[_ggfb[0]:_ggfb[1]]})
	}
	return _fddf, nil
}

// Write writes the content of `re.creator` to writer of type io.Writer interface.
func (_cae *Redactor) Write(writer _e.Writer) error { return _cae._afbb.Write(writer) }
func _gfba(_gegc []int, _gedd *_gf.TextMarkArray, _cfag string) (*_gf.TextMarkArray, matchedBBox, error) {
	_ffea := matchedBBox{}
	_feec := _gegc[0]
	_gdf := _gegc[1]
	_cbc := len(_cfag) - len(_g.TrimLeft(_cfag, "\u0020"))
	_bda := len(_cfag) - len(_g.TrimRight(_cfag, "\u0020\u000a"))
	_feec = _feec + _cbc
	_gdf = _gdf - _bda
	_cfag = _g.Trim(_cfag, "\u0020\u000a")
	_afc, _defg := _gedd.RangeOffset(_feec, _gdf)
	if _defg != nil {
		return nil, _ffea, _defg
	}
	_faf, _ecag := _afc.BBox()
	if !_ecag {
		return nil, _ffea, _af.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _afc)
	}
	_ffea = matchedBBox{_gfc: _cfag, _ddd: _faf}
	return _afc, _ffea, nil
}
func _fgaa(_aecf, _cdcg targetMap) (bool, []int) {
	_cfb := _g.Contains(_aecf._gea, _cdcg._gea)
	var _ecfa []int
	for _, _edb := range _aecf._aeef {
		for _bbca, _fabc := range _cdcg._aeef {
			if _fabc[0] >= _edb[0] && _fabc[1] <= _edb[1] {
				_ecfa = append(_ecfa, _bbca)
			}
		}
	}
	return _cfb, _ecfa
}
func _cef(_bddg *_gf.TextMarkArray, _ecde int, _affd int) int {
	_fad := _bddg.Elements()
	_ecdf := _ecde - 1
	_abb := _ecde + 1
	_bgcde := -1
	if _ecdf >= 0 {
		_aeg := _fad[_ecdf]
		_ffcg := _aeg.ObjString
		_becg := len(_ffcg)
		_beab := _aeg.Index
		if _beab+1 < _becg {
			_bgcde = _ecdf
			return _bgcde
		}
	}
	if _abb < len(_fad) {
		_acf := _fad[_abb]
		_befb := _acf.ObjString
		if _befb[0] != _acf.Text {
			_bgcde = _abb
			return _bgcde
		}
	}
	return _bgcde
}
func _cg(_dag *_eb.ContentStreamOperations, _ec map[_ef.PdfObject][]localSpanMarks) error {
	for _eg, _ae := range _ec {
		if _eg == nil {
			continue
		}
		_ca, _cgc, _ab := _fbea(_dag, _eg)
		if !_ab {
			_gg.Log.Debug("Pd\u0066\u004fb\u006a\u0065\u0063\u0074\u0020\u0025\u0073\u006e\u006ft\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0069\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073\u0074r\u0065a\u006d\u0020\u006f\u0070\u0065\u0072\u0061\u0074i\u006fn\u0020\u0025s", _eg, _dag)
			return nil
		}
		if _ca.Operand == "\u0054\u006a" {
			_dc := _bab(_ca, _eg, _ae)
			if _dc != nil {
				return _dc
			}
		} else if _ca.Operand == "\u0054\u004a" {
			_ea := _dg(_ca, _eg, _ae)
			if _ea != nil {
				return _ea
			}
		} else if _ca.Operand == "\u0027" || _ca.Operand == "\u0022" {
			_ac := _abg(_dag, _ca.Operand, _cgc)
			if _ac != nil {
				return _ac
			}
			_ac = _bab(_ca, _eg, _ae)
			if _ac != nil {
				return _ac
			}
		}
	}
	return nil
}

type targetMap struct {
	_gea  string
	_aeef [][]int
}

func _aafa(_fecd string) (string, [][]int) {
	_acfb := _d.MustCompile("\u005c\u006e")
	_cfd := _acfb.FindAllStringIndex(_fecd, -1)
	_daba := _acfb.ReplaceAllString(_fecd, "\u0020")
	return _daba, _cfd
}

// RedactRectanglePropsNew return a new pointer to a default RectangleProps object.
func RedactRectanglePropsNew() *RectangleProps {
	return &RectangleProps{FillColor: _ge.ColorBlack, BorderWidth: 0.0, FillOpacity: 1.0}
}
func _bfb(_ebdgd RedactionTerm) (*regexMatcher, error) { return &regexMatcher{_dfdg: _ebdgd}, nil }
func _fca(_dec *_gf.TextMarkArray) *_dd.PdfFont {
	_, _gde := _cd(_dec)
	_aga := _dec.Elements()[_gde]
	_edg := _aga.Font
	return _edg
}
func _cgce(_aeb *_gf.TextMarkArray) string {
	_befg := ""
	for _, _dbc := range _aeb.Elements() {
		_befg += _dbc.Text
	}
	return _befg
}
func _fbea(_gfbg *_eb.ContentStreamOperations, PdfObj _ef.PdfObject) (*_eb.ContentStreamOperation, int, bool) {
	for _cfae, _ded := range *_gfbg {
		_beaa := _ded.Operand
		if _beaa == "\u0054\u006a" {
			_aeaa := _ef.TraceToDirectObject(_ded.Params[0])
			if _aeaa == PdfObj {
				return _ded, _cfae, true
			}
		} else if _beaa == "\u0054\u004a" {
			_fce, _afg := _ef.GetArray(_ded.Params[0])
			if !_afg {
				return nil, _cfae, _afg
			}
			for _, _acc := range _fce.Elements() {
				if _acc == PdfObj {
					return _ded, _cfae, true
				}
			}
		} else if _beaa == "\u0022" {
			_fbb := _ef.TraceToDirectObject(_ded.Params[2])
			if _fbb == PdfObj {
				return _ded, _cfae, true
			}
		} else if _beaa == "\u0027" {
			_dfbe := _ef.TraceToDirectObject(_ded.Params[0])
			if _dfbe == PdfObj {
				return _ded, _cfae, true
			}
		}
	}
	return nil, -1, false
}

type placeHolders struct {
	_da []int
	_b  string
	_ed float64
}

func _egg(_ccc []placeHolders) []replacement {
	_cbd := []replacement{}
	for _, _cdb := range _ccc {
		_eeae := _cdb._da
		_fcd := _cdb._b
		_ebb := _cdb._ed
		for _, _bgc := range _eeae {
			_ebdg := replacement{_ggd: _fcd, _bc: _ebb, _db: _bgc}
			_cbd = append(_cbd, _ebdg)
		}
	}
	_c.Slice(_cbd, func(_dge, _fgc int) bool { return _cbd[_dge]._db < _cbd[_fgc]._db })
	return _cbd
}
func _ecdc(_cdf string, _aac []localSpanMarks) ([]placeHolders, error) {
	_cdfc := ""
	_fbdb := []placeHolders{}
	for _fbdd, _ccg := range _aac {
		_cdfg := _ccg._edc
		_decf := _ccg._fffa
		_ggf := _cgce(_cdfg)
		_dece, _afb := _dgef(_cdfg)
		if _afb != nil {
			return nil, _afb
		}
		if _ggf != _cdfc {
			var _gdc []int
			if _fbdd == 0 && _decf != _ggf {
				_acef := _g.Index(_cdf, _ggf)
				_gdc = []int{_acef}
			} else if _fbdd == len(_aac)-1 {
				_bbd := _g.LastIndex(_cdf, _ggf)
				_gdc = []int{_bbd}
			} else {
				_gdc = _bga(_cdf, _ggf)
			}
			_ebge := placeHolders{_da: _gdc, _b: _ggf, _ed: _dece}
			_fbdb = append(_fbdb, _ebge)
		}
		_cdfc = _ggf
	}
	return _fbdb, nil
}

// New instantiates a Redactor object with given PdfReader and `regex` pattern.
func New(reader *_dd.PdfReader, opts *RedactionOptions, rectProps *RectangleProps) *Redactor {
	if rectProps == nil {
		rectProps = RedactRectanglePropsNew()
	}
	return &Redactor{_gee: reader, _gffa: opts, _afbb: _ge.New(), _gcg: rectProps}
}
func _cf(_agg []localSpanMarks) (map[string][]localSpanMarks, []string) {
	_baf := make(map[string][]localSpanMarks)
	_defe := []string{}
	for _, _dda := range _agg {
		_cbg := _dda._fffa
		if _bgg, _fcf := _baf[_cbg]; _fcf {
			_baf[_cbg] = append(_bgg, _dda)
		} else {
			_baf[_cbg] = []localSpanMarks{_dda}
			_defe = append(_defe, _cbg)
		}
	}
	return _baf, _defe
}
func _gdd(_ccab []*matchedIndex) []*targetMap {
	_gebf := make(map[string][][]int)
	_egfe := []*targetMap{}
	for _, _adc := range _ccab {
		_fffe := _adc._fab
		_addb := []int{_adc._faa, _adc._gaf}
		if _fdgc, _efc := _gebf[_fffe]; _efc {
			_gebf[_fffe] = append(_fdgc, _addb)
		} else {
			_gebf[_fffe] = [][]int{_addb}
		}
	}
	for _eaaa, _bcba := range _gebf {
		_cde := &targetMap{_gea: _eaaa, _aeef: _bcba}
		_egfe = append(_egfe, _cde)
	}
	return _egfe
}
