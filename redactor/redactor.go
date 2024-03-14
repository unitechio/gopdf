package redactor

import (
	_fa "errors"
	_f "fmt"
	_fb "io"
	_a "regexp"
	_b "sort"
	_g "strings"

	_e "bitbucket.org/shenghui0779/gopdf/common"
	_dd "bitbucket.org/shenghui0779/gopdf/contentstream"
	_cb "bitbucket.org/shenghui0779/gopdf/core"
	_d "bitbucket.org/shenghui0779/gopdf/creator"
	_aa "bitbucket.org/shenghui0779/gopdf/extractor"
	_bb "bitbucket.org/shenghui0779/gopdf/model"
)

// Redact executes the redact operation on a pdf file and updates the content streams of all pages of the file.
func (_gba *Redactor) Redact() error {
	_bdbf, _gaddf := _gba._bag.GetNumPages()
	if _gaddf != nil {
		return _f.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0074\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020P\u0061\u0067\u0065\u0073")
	}
	_fggg := _gba._bgba.FillColor
	_cgfg := _gba._bgba.BorderWidth
	_aaafe := _gba._bgba.FillOpacity
	for _bgc := 1; _bgc <= _bdbf; _bgc++ {
		_edc, _ebgf := _gba._bag.GetPage(_bgc)
		if _ebgf != nil {
			return _ebgf
		}
		_gfa, _ebgf := _aa.New(_edc)
		if _ebgf != nil {
			return _ebgf
		}
		_cegf, _, _, _ebgf := _gfa.ExtractPageText()
		if _ebgf != nil {
			return _ebgf
		}
		_ddfg := _cegf.GetContentStreamOps()
		_feff, _gfdg, _ebgf := _gba.redactPage(_ddfg, _edc.Resources)
		if _gfdg == nil {
			_e.Log.Info("N\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u0020t\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065d \u0070\u0061\u0074t\u0061r\u006e\u002e")
			_gfdg = _ddfg
		}
		_baf := _dd.ContentStreamOperation{Operand: "\u006e"}
		*_gfdg = append(*_gfdg, &_baf)
		_edc.SetContentStreams([]string{_gfdg.String()}, _cb.NewFlateEncoder())
		if _ebgf != nil {
			return _ebgf
		}
		_ggg, _ebgf := _edc.GetMediaBox()
		if _ebgf != nil {
			return _ebgf
		}
		if _edc.MediaBox == nil {
			_edc.MediaBox = _ggg
		}
		if _bed := _gba._bcc.AddPage(_edc); _bed != nil {
			return _bed
		}
		_b.Slice(_feff, func(_eedb, _fce int) bool { return _feff[_eedb]._bgb < _feff[_fce]._bgb })
		_edef := _ggg.Ury
		for _, _baeg := range _feff {
			_ddff := _baeg._cfgg
			_cbc := _gba._bcc.NewRectangle(_ddff.Llx, _edef-_ddff.Lly, _ddff.Urx-_ddff.Llx, -(_ddff.Ury - _ddff.Lly))
			_cbc.SetFillColor(_fggg)
			_cbc.SetBorderWidth(_cgfg)
			_cbc.SetFillOpacity(_aaafe)
			if _ebf := _gba._bcc.Draw(_cbc); _ebf != nil {
				return nil
			}
		}
	}
	_gba._bcc.SetOutlineTree(_gba._bag.GetOutlineTree())
	return nil
}

func _ggd(_fee *_bb.PdfFont, _cegc _aa.TextMark) float64 {
	_ead := 0.001
	_de := _cegc.Th / 100
	if _fee.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_ead = 1
	}
	_aeg, _gfeb := _fee.GetRuneMetrics(' ')
	if !_gfeb {
		_aeg, _gfeb = _fee.GetCharMetrics(32)
	}
	if !_gfeb {
		_aeg, _ = _bb.DefaultFont().GetRuneMetrics(' ')
	}
	_ccf := _ead * ((_aeg.Wx*_cegc.FontSize + _cegc.Tc + _cegc.Tw) / _de)
	return _ccf
}

func _feca(_bcffb *_aa.TextMarkArray) []*_aa.TextMarkArray {
	_fcd := _bcffb.Elements()
	_acc := len(_fcd)
	var _bdbag _cb.PdfObject
	_ecg := []*_aa.TextMarkArray{}
	_eegg := &_aa.TextMarkArray{}
	_bga := -1
	for _ada, _bda := range _fcd {
		_ecbg := _bda.DirectObject
		_bga = _bda.Index
		if _ecbg == nil {
			_gdd := _fac(_bcffb, _ada, _bga)
			if _bdbag != nil {
				if _gdd == -1 || _gdd > _ada {
					_ecg = append(_ecg, _eegg)
					_eegg = &_aa.TextMarkArray{}
				}
			}
		} else if _ecbg != nil && _bdbag == nil {
			if _bga == 0 && _ada > 0 {
				_ecg = append(_ecg, _eegg)
				_eegg = &_aa.TextMarkArray{}
			}
		} else if _ecbg != nil && _bdbag != nil {
			if _ecbg != _bdbag {
				_ecg = append(_ecg, _eegg)
				_eegg = &_aa.TextMarkArray{}
			}
		}
		_bdbag = _ecbg
		_eegg.Append(_bda)
		if _ada == (_acc - 1) {
			_ecg = append(_ecg, _eegg)
		}
	}
	return _ecg
}

func _gd(_gaf []localSpanMarks) (map[string][]localSpanMarks, []string) {
	_faab := make(map[string][]localSpanMarks)
	_eae := []string{}
	for _, _ffab := range _gaf {
		_cde := _ffab._eec
		if _cec, _gace := _faab[_cde]; _gace {
			_faab[_cde] = append(_cec, _ffab)
		} else {
			_faab[_cde] = []localSpanMarks{_ffab}
			_eae = append(_eae, _cde)
		}
	}
	return _faab, _eae
}

func _fgc(_gebd localSpanMarks, _bdg *_aa.TextMarkArray, _gcg *_bb.PdfFont, _afd, _df string) ([]_cb.PdfObject, error) {
	_cbbd := _cgf(_bdg)
	Tj, _dac := _deg(_bdg)
	if _dac != nil {
		return nil, _dac
	}
	_fea := len(_afd)
	_ddg := len(_cbbd)
	_gadd := -1
	_cbd := _cb.MakeFloat(Tj)
	if _cbbd != _df {
		_fegb := _gebd._abd
		if _fegb == 0 {
			_gadd = _g.LastIndex(_afd, _cbbd)
		} else {
			_gadd = _g.Index(_afd, _cbbd)
		}
	} else {
		_gadd = _g.Index(_afd, _cbbd)
	}
	_cac := _gadd + _ddg
	_cee := []_cb.PdfObject{}
	if _gadd == 0 && _cac == _fea {
		_cee = append(_cee, _cbd)
	} else if _gadd == 0 && _cac < _fea {
		_dec := _acga(_afd[_cac:], _gcg)
		_dbgg := _cb.MakeStringFromBytes(_dec)
		_cee = append(_cee, _cbd, _dbgg)
	} else if _gadd > 0 && _cac >= _fea {
		_efb := _acga(_afd[:_gadd], _gcg)
		_ebg := _cb.MakeStringFromBytes(_efb)
		_cee = append(_cee, _ebg, _cbd)
	} else if _gadd > 0 && _cac < _fea {
		_egb := _acga(_afd[:_gadd], _gcg)
		_agcb := _acga(_afd[_cac:], _gcg)
		_aga := _cb.MakeStringFromBytes(_egb)
		_adb := _cb.MakeString(string(_agcb))
		_cee = append(_cee, _aga, _cbd, _adb)
	}
	return _cee, nil
}

func _eea(_bbac []*matchedIndex, _gbe [][]int) []*matchedIndex {
	_cccef := []*matchedIndex{}
	for _, _caed := range _bbac {
		_fbg, _bgdc := _cddc(_caed, _gbe)
		if _fbg {
			_gede := _gdg(_caed, _bgdc)
			_cccef = append(_cccef, _gede...)
		} else {
			_cccef = append(_cccef, _caed)
		}
	}
	return _cccef
}

func _cddc(_fffb *matchedIndex, _age [][]int) (bool, [][]int) {
	_ggb := [][]int{}
	for _, _cdec := range _age {
		if _fffb._ccff < _cdec[0] && _fffb._aaga > _cdec[1] {
			_ggb = append(_ggb, _cdec)
		}
	}
	return len(_ggb) > 0, _ggb
}

func (_cab *Redactor) redactPage(_bfe *_dd.ContentStreamOperations, _cbab *_bb.PdfPageResources) ([]matchedBBox, *_dd.ContentStreamOperations, error) {
	_gcgf, _egab := _aa.NewFromContents(_bfe.String(), _cbab)
	if _egab != nil {
		return nil, nil, _egab
	}
	_ffb, _, _, _egab := _gcgf.ExtractPageText()
	if _egab != nil {
		return nil, nil, _egab
	}
	_bfe = _ffb.GetContentStreamOps()
	_fae := _ffb.Marks()
	_eee := _ffb.Text()
	_eee, _bfg := _cgg(_eee)
	_aba := []matchedBBox{}
	_bgfc := make(map[_cb.PdfObject][]localSpanMarks)
	_dga := []*targetMap{}
	for _, _afdc := range _cab._ffaf.Terms {
		_agfe, _eab := _bace(_afdc)
		if _eab != nil {
			return nil, nil, _eab
		}
		_abc, _eab := _agfe.match(_eee)
		if _eab != nil {
			return nil, nil, _eab
		}
		_abc = _eea(_abc, _bfg)
		_fbec := _ddbe(_abc)
		_dga = append(_dga, _fbec...)
	}
	_bdff(_dga)
	for _, _gcdc := range _dga {
		_dab := _gcdc._edgd
		_gbf := _gcdc._fbcf
		_cfa := []matchedBBox{}
		for _, _cabb := range _gbf {
			_dbae, _eccc, _ded := _edd(_cabb, _fae, _dab)
			if _ded != nil {
				return nil, nil, _ded
			}
			_cda := _feca(_dbae)
			for _cabf, _bfc := range _cda {
				_dcb := localSpanMarks{_gefe: _bfc, _abd: _cabf, _eec: _dab}
				_ceea, _ := _be(_bfc)
				if _gacb, _bgd := _bgfc[_ceea]; _bgd {
					_bgfc[_ceea] = append(_gacb, _dcb)
				} else {
					_bgfc[_ceea] = []localSpanMarks{_dcb}
				}
			}
			_cfa = append(_cfa, _eccc)
		}
		_aba = append(_aba, _cfa...)
	}
	_egab = _fbc(_bfe, _bgfc)
	if _egab != nil {
		return nil, nil, _egab
	}
	return _aba, _bfe, nil
}

func _gbbb(_dgbc int, _aeb []int) bool {
	for _, _bffg := range _aeb {
		if _bffg == _dgbc {
			return true
		}
	}
	return false
}

// RedactionTerm holds the regexp pattern and the replacement string for the redaction process.
type RedactionTerm struct{ Pattern *_a.Regexp }

func _fedd(_dbc []placeHolders) []replacement {
	_ecb := []replacement{}
	for _, _afe := range _dbc {
		_deba := _afe._ga
		_eccb := _afe._ed
		_fdd := _afe._ac
		for _, _aea := range _deba {
			_cag := replacement{_ff: _eccb, _gc: _fdd, _da: _aea}
			_ecb = append(_ecb, _cag)
		}
	}
	_b.Slice(_ecb, func(_bab, _ecf int) bool { return _ecb[_bab]._da < _ecb[_ecf]._da })
	return _ecb
}

func _cgf(_agf *_aa.TextMarkArray) string {
	_cfc := ""
	for _, _aegc := range _agf.Elements() {
		_cfc += _aegc.Text
	}
	return _cfc
}

type matchedBBox struct {
	_cfgg _bb.PdfRectangle
	_bgb  string
}

func _deg(_dagg *_aa.TextMarkArray) (float64, error) {
	_eag, _fgcb := _dagg.BBox()
	if !_fgcb {
		return 0.0, _f.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _dagg)
	}
	_bdf := _ef(_dagg)
	_fbb := 0.0
	_, _gbb := _be(_dagg)
	_eaeb := _dagg.Elements()[_gbb]
	_dcfc := _eaeb.Font
	if _bdf > 0 {
		_fbb = _ggd(_dcfc, _eaeb)
	}
	_cdd := (_eag.Urx - _eag.Llx)
	_cdd = _cdd + _fbb*float64(_bdf)
	Tj := _cfed(_cdd, _eaeb.FontSize, _eaeb.Th)
	return Tj, nil
}

// Write writes the content of `re.creator` to writer of type io.Writer interface.
func (_gfdc *Redactor) Write(writer _fb.Writer) error { return _gfdc._bcc.Write(writer) }

func _bdff(_daga []*targetMap) {
	for _cbabf, _gbba := range _daga {
		for _bafe, _eeda := range _daga {
			if _cbabf != _bafe {
				_dae, _edee := _eca(*_gbba, *_eeda)
				if _dae {
					_bge(_eeda, _edee)
				}
			}
		}
	}
}

type placeHolders struct {
	_ga []int
	_ed string
	_ac float64
}
type matchedIndex struct {
	_ccff  int
	_aaga  int
	_adbbf string
}

// RectangleProps defines properties of the redaction rectangle to be drawn.
type RectangleProps struct {
	FillColor   _d.Color
	BorderWidth float64
	FillOpacity float64
}

func _eca(_ccc, _dcd targetMap) (bool, []int) {
	_fbf := _g.Contains(_ccc._edgd, _dcd._edgd)
	var _cadbb []int
	for _, _adcc := range _ccc._fbcf {
		for _dcda, _ege := range _dcd._fbcf {
			if _ege[0] >= _adcc[0] && _ege[1] <= _adcc[1] {
				_cadbb = append(_cadbb, _dcda)
			}
		}
	}
	return _fbf, _cadbb
}

// RedactRectanglePropsNew return a new pointer to a default RectangleProps object.
func RedactRectanglePropsNew() *RectangleProps {
	return &RectangleProps{FillColor: _d.ColorBlack, BorderWidth: 0.0, FillOpacity: 1.0}
}

func _feaf(_fgd, _eedf string) []int {
	if len(_eedf) == 0 {
		return nil
	}
	var _gbc []int
	for _bffc := 0; _bffc < len(_fgd); {
		_fdb := _g.Index(_fgd[_bffc:], _eedf)
		if _fdb < 0 {
			return _gbc
		}
		_gbc = append(_gbc, _bffc+_fdb)
		_bffc += _fdb + len(_eedf)
	}
	return _gbc
}

type localSpanMarks struct {
	_gefe *_aa.TextMarkArray
	_abd  int
	_eec  string
}

func (_aagd *regexMatcher) match(_fca string) ([]*matchedIndex, error) {
	_ccb := _aagd._afg.Pattern
	if _ccb == nil {
		return nil, _fa.New("\u006e\u006f\u0020\u0070at\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006d\u0070\u0069\u006c\u0065\u0064")
	}
	var (
		_fbfc = _ccb.FindAllStringIndex(_fca, -1)
		_bafg []*matchedIndex
	)
	for _, _edgg := range _fbfc {
		_bafg = append(_bafg, &matchedIndex{_ccff: _edgg[0], _aaga: _edgg[1], _adbbf: _fca[_edgg[0]:_edgg[1]]})
	}
	return _bafg, nil
}

func _acga(_cbe string, _gcd *_bb.PdfFont) []byte {
	_eege, _bfa := _gcd.StringToCharcodeBytes(_cbe)
	if _bfa != 0 {
		_e.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0073\u006fm\u0065\u0020\u0072un\u0065\u0073\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065d\u002e\u000a\u0009\u0025\u0073\u0020\u002d\u003e \u0025\u0076", _cbe, _eege)
	}
	return _eege
}

func _aad(_bbf _cb.PdfObject, _dbaa *_bb.PdfFont) (string, error) {
	_adde, _bbd := _cb.GetStringBytes(_bbf)
	if !_bbd {
		return "", _cb.ErrTypeError
	}
	_adbd := _dbaa.BytesToCharcodes(_adde)
	_cdbg, _dage, _aff := _dbaa.CharcodesToStrings(_adbd)
	if _aff > 0 {
		_e.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _dage, _aff)
	}
	_dgc := _g.Join(_cdbg, "")
	return _dgc, nil
}

func _cfed(_cfec, _ede, _bdd float64) float64 {
	_bdd = _bdd / 100
	_gga := (-1000 * _cfec) / (_ede * _bdd)
	return _gga
}

func _gaeg(_cff *_dd.ContentStreamOperations, PdfObj _cb.PdfObject) (*_dd.ContentStreamOperation, int, bool) {
	for _agda, _bdgb := range *_cff {
		_febb := _bdgb.Operand
		if _febb == "\u0054\u006a" {
			_cadb := _cb.TraceToDirectObject(_bdgb.Params[0])
			if _cadb == PdfObj {
				return _bdgb, _agda, true
			}
		} else if _febb == "\u0054\u004a" {
			_deac, _geg := _cb.GetArray(_bdgb.Params[0])
			if !_geg {
				return nil, _agda, _geg
			}
			for _, _cefeg := range _deac.Elements() {
				if _cefeg == PdfObj {
					return _bdgb, _agda, true
				}
			}
		} else if _febb == "\u0022" {
			_abb := _cb.TraceToDirectObject(_bdgb.Params[2])
			if _abb == PdfObj {
				return _bdgb, _agda, true
			}
		} else if _febb == "\u0027" {
			_bca := _cb.TraceToDirectObject(_bdgb.Params[0])
			if _bca == PdfObj {
				return _bdgb, _agda, true
			}
		}
	}
	return nil, -1, false
}

func _cgg(_ccce string) (string, [][]int) {
	_gcec := _a.MustCompile("\u005c\u006e")
	_bffd := _gcec.FindAllStringIndex(_ccce, -1)
	_cadea := _gcec.ReplaceAllString(_ccce, "\u0020")
	return _cadea, _bffd
}

func _gdg(_fgb *matchedIndex, _cbf [][]int) []*matchedIndex {
	_fbaf := []*matchedIndex{}
	_gebda := _fgb._ccff
	_ceef := _gebda
	_cea := _fgb._adbbf
	_dabg := 0
	for _, _gcc := range _cbf {
		_eage := _gcc[0] - _gebda
		if _dabg >= _eage {
			continue
		}
		_gegd := _cea[_dabg:_eage]
		_bef := &matchedIndex{_adbbf: _gegd, _ccff: _ceef, _aaga: _gcc[0]}
		if len(_g.TrimSpace(_gegd)) != 0 {
			_fbaf = append(_fbaf, _bef)
		}
		_dabg = _gcc[1] - _gebda
		_ceef = _gebda + _dabg
	}
	_bdce := _cea[_dabg:]
	_dff := &matchedIndex{_adbbf: _bdce, _ccff: _ceef, _aaga: _fgb._aaga}
	if len(_g.TrimSpace(_bdce)) != 0 {
		_fbaf = append(_fbaf, _dff)
	}
	return _fbaf
}

type targetMap struct {
	_edgd string
	_fbcf [][]int
}

// WriteToFile writes the redacted document to file specified by `outputPath`.
func (_ece *Redactor) WriteToFile(outputPath string) error {
	if _dcfcg := _ece._bcc.WriteToFile(outputPath); _dcfcg != nil {
		return _f.Errorf("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0077\u0072\u0069\u0074\u0065\u0020t\u0068\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0066\u0069\u006c\u0065")
	}
	return nil
}

func _ggdd(_fdf *_dd.ContentStreamOperation, _dcc _cb.PdfObject, _ec []localSpanMarks) error {
	var _gab *_cb.PdfObjectArray
	_gfc, _db := _gd(_ec)
	if len(_db) == 1 {
		_dcfg := _db[0]
		_aag := _gfc[_dcfg]
		if len(_aag) == 1 {
			_ag := _aag[0]
			_cadf := _ag._gefe
			_acf := _bg(_cadf)
			_fda, _ffd := _aad(_dcc, _acf)
			if _ffd != nil {
				return _ffd
			}
			_feg, _ffd := _fgc(_ag, _cadf, _acf, _fda, _dcfg)
			if _ffd != nil {
				return _ffd
			}
			_gab = _cb.MakeArray(_feg...)
		} else {
			_cbb := _aag[0]._gefe
			_agc := _bg(_cbb)
			_ecc, _cdg := _aad(_dcc, _agc)
			if _cdg != nil {
				return _cdg
			}
			_cfge, _cdg := _eed(_ecc, _aag)
			if _cdg != nil {
				return _cdg
			}
			_dee := _fedd(_cfge)
			_bdb := _dba(_ecc, _dee, _agc)
			_gab = _cb.MakeArray(_bdb...)
		}
	} else if len(_db) > 1 {
		_cef := _ec[0]
		_effc := _cef._gefe
		_, _deb := _be(_effc)
		_fdab := _effc.Elements()[_deb]
		_gec := _fdab.Font
		_cefe, _bae := _aad(_dcc, _gec)
		if _bae != nil {
			return _bae
		}
		_dbg, _bae := _eed(_cefe, _ec)
		if _bae != nil {
			return _bae
		}
		_gff := _fedd(_dbg)
		_add := _dba(_cefe, _gff, _gec)
		_gab = _cb.MakeArray(_add...)
	}
	_fdf.Params[0] = _gab
	_fdf.Operand = "\u0054\u004a"
	return nil
}

func _dba(_fdfg string, _bad []replacement, _gfgf *_bb.PdfFont) []_cb.PdfObject {
	_gbd := []_cb.PdfObject{}
	_gece := 0
	_dfg := _fdfg
	for _aacf, _fdg := range _bad {
		_caa := _fdg._da
		_aaaf := _fdg._gc
		_edg := _fdg._ff
		_ddf := _cb.MakeFloat(_aaaf)
		if _gece > _caa || _caa == -1 {
			continue
		}
		_bfab := _fdfg[_gece:_caa]
		_cbaf := _acga(_bfab, _gfgf)
		_gabg := _cb.MakeStringFromBytes(_cbaf)
		_gbd = append(_gbd, _gabg)
		_gbd = append(_gbd, _ddf)
		_dbe := _caa + len(_edg)
		_dfg = _fdfg[_dbe:]
		_gece = _dbe
		if _aacf == len(_bad)-1 {
			_cbaf = _acga(_dfg, _gfgf)
			_gabg = _cb.MakeStringFromBytes(_cbaf)
			_gbd = append(_gbd, _gabg)
		}
	}
	return _gbd
}

func _bg(_faf *_aa.TextMarkArray) *_bb.PdfFont {
	_, _aaa := _be(_faf)
	_bcd := _faf.Elements()[_aaa]
	_gae := _bcd.Font
	return _gae
}

func _ef(_fgg *_aa.TextMarkArray) int {
	_af := 0
	_acg := _fgg.Elements()
	if _acg[0].Text == "\u0020" {
		_af++
	}
	if _acg[_fgg.Len()-1].Text == "\u0020" {
		_af++
	}
	return _af
}

// New instantiates a Redactor object with given PdfReader and `regex` pattern.
func New(reader *_bb.PdfReader, opts *RedactionOptions, rectProps *RectangleProps) *Redactor {
	if rectProps == nil {
		rectProps = RedactRectanglePropsNew()
	}
	return &Redactor{_bag: reader, _ffaf: opts, _bcc: _d.New(), _bgba: rectProps}
}

func _be(_fe *_aa.TextMarkArray) (_cb.PdfObject, int) {
	var _dg _cb.PdfObject
	_ce := -1
	for _fad, _cf := range _fe.Elements() {
		_dg = _cf.DirectObject
		_ce = _fad
		if _dg != nil {
			break
		}
	}
	return _dg, _ce
}
func _bace(_gbbae RedactionTerm) (*regexMatcher, error) { return &regexMatcher{_afg: _gbbae}, nil }
func _ddbe(_dfc []*matchedIndex) []*targetMap {
	_cgc := make(map[string][][]int)
	_cccc := []*targetMap{}
	for _, _dda := range _dfc {
		_addeg := _dda._adbbf
		_bffa := []int{_dda._ccff, _dda._aaga}
		if _caf, _cbda := _cgc[_addeg]; _cbda {
			_cgc[_addeg] = append(_caf, _bffa)
		} else {
			_cgc[_addeg] = [][]int{_bffa}
		}
	}
	for _fba, _fag := range _cgc {
		_cddd := &targetMap{_edgd: _fba, _fbcf: _fag}
		_cccc = append(_cccc, _cddd)
	}
	return _cccc
}

// Redactor represents a Redactor object.
type Redactor struct {
	_bag  *_bb.PdfReader
	_ffaf *RedactionOptions
	_bcc  *_d.Creator
	_bgba *RectangleProps
}

func _fbc(_ad *_dd.ContentStreamOperations, _ca map[_cb.PdfObject][]localSpanMarks) error {
	for _dc, _gg := range _ca {
		if _dc == nil {
			continue
		}
		_cc, _ge, _bba := _gaeg(_ad, _dc)
		if !_bba {
			_e.Log.Debug("Pd\u0066\u004fb\u006a\u0065\u0063\u0074\u0020\u0025\u0073\u006e\u006ft\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0069\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073\u0074r\u0065a\u006d\u0020\u006f\u0070\u0065\u0072\u0061\u0074i\u006fn\u0020\u0025s", _dc, _ad)
			return nil
		}
		if _cc.Operand == "\u0054\u006a" {
			_aab := _ggdd(_cc, _dc, _gg)
			if _aab != nil {
				return _aab
			}
		} else if _cc.Operand == "\u0054\u004a" {
			_fc := _ceg(_cc, _dc, _gg)
			if _fc != nil {
				return _fc
			}
		} else if _cc.Operand == "\u0027" || _cc.Operand == "\u0022" {
			_bc := _ee(_ad, _cc.Operand, _ge)
			if _bc != nil {
				return _bc
			}
			_bc = _ggdd(_cc, _dc, _gg)
			if _bc != nil {
				return _bc
			}
		}
	}
	return nil
}

// RedactionOptions is a collection of RedactionTerm objects.
type RedactionOptions struct{ Terms []RedactionTerm }

func _bge(_cbea *targetMap, _addd []int) {
	var _edec [][]int
	for _cbdf, _eece := range _cbea._fbcf {
		if _gbbb(_cbdf, _addd) {
			continue
		}
		_edec = append(_edec, _eece)
	}
	_cbea._fbcf = _edec
}

func _edd(_bfb []int, _aef *_aa.TextMarkArray, _dbb string) (*_aa.TextMarkArray, matchedBBox, error) {
	_feddb := matchedBBox{}
	_ebga := _bfb[0]
	_ggc := _bfb[1]
	_aaf := len(_dbb) - len(_g.TrimLeft(_dbb, "\u0020"))
	_bcff := len(_dbb) - len(_g.TrimRight(_dbb, "\u0020\u000a"))
	_ebga = _ebga + _aaf
	_ggc = _ggc - _bcff
	_dbb = _g.Trim(_dbb, "\u0020\u000a")
	_gfge, _cca := _aef.RangeOffset(_ebga, _ggc)
	if _cca != nil {
		return nil, _feddb, _cca
	}
	_geea, _bfef := _gfge.BBox()
	if !_bfef {
		return nil, _feddb, _f.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _gfge)
	}
	_feddb = matchedBBox{_bgb: _dbb, _cfgg: _geea}
	return _gfge, _feddb, nil
}

func _eed(_bgf string, _cba []localSpanMarks) ([]placeHolders, error) {
	_gdf := ""
	_ab := []placeHolders{}
	for _aac, _cg := range _cba {
		_fed := _cg._gefe
		_bcb := _cg._eec
		_fec := _cgf(_fed)
		_gee, _agd := _deg(_fed)
		if _agd != nil {
			return nil, _agd
		}
		if _fec != _gdf {
			var _aca []int
			if _aac == 0 && _bcb != _fec {
				_bff := _g.Index(_bgf, _fec)
				_aca = []int{_bff}
			} else if _aac == len(_cba)-1 {
				_fcf := _g.LastIndex(_bgf, _fec)
				_aca = []int{_fcf}
			} else {
				_aca = _feaf(_bgf, _fec)
			}
			_adc := placeHolders{_ga: _aca, _ed: _fec, _ac: _gee}
			_ab = append(_ab, _adc)
		}
		_gdf = _fec
	}
	return _ab, nil
}

func _ceg(_gef *_dd.ContentStreamOperation, _dag _cb.PdfObject, _fbe []localSpanMarks) error {
	_cd, _gacf := _cb.GetArray(_gef.Params[0])
	_cad := []_cb.PdfObject{}
	if !_gacf {
		_e.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _gef)
		return _f.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _gef)
	}
	_bf, _gaa := _gd(_fbe)
	if len(_gaa) == 1 {
		_ae := _gaa[0]
		_ffa := _bf[_ae]
		if len(_ffa) == 1 {
			_gad := _ffa[0]
			_cfe := _gad._gefe
			_ea := _bg(_cfe)
			_ba, _cade := _aad(_dag, _ea)
			if _cade != nil {
				return _cade
			}
			_dgb, _cade := _fgc(_gad, _cfe, _ea, _ba, _ae)
			if _cade != nil {
				return _cade
			}
			for _, _ged := range _cd.Elements() {
				if _ged == _dag {
					_cad = append(_cad, _dgb...)
				} else {
					_cad = append(_cad, _ged)
				}
			}
		} else {
			_aabe := _ffa[0]._gefe
			_faa := _bg(_aabe)
			_eff, _bcf := _aad(_dag, _faa)
			if _bcf != nil {
				return _bcf
			}
			_bd, _bcf := _eed(_eff, _ffa)
			if _bcf != nil {
				return _bcf
			}
			_feb := _fedd(_bd)
			_gfb := _dba(_eff, _feb, _faa)
			for _, _gce := range _cd.Elements() {
				if _gce == _dag {
					_cad = append(_cad, _gfb...)
				} else {
					_cad = append(_cad, _gce)
				}
			}
		}
		_gef.Params[0] = _cb.MakeArray(_cad...)
	} else if len(_gaa) > 1 {
		_egf := _fbe[0]
		_afa := _egf._gefe
		_, _fd := _be(_afa)
		_cfg := _afa.Elements()[_fd]
		_gb := _cfg.Font
		_gacg, _fab := _aad(_dag, _gb)
		if _fab != nil {
			return _fab
		}
		_ega, _fab := _eed(_gacg, _fbe)
		if _fab != nil {
			return _fab
		}
		_fef := _fedd(_ega)
		_acb := _dba(_gacg, _fef, _gb)
		for _, _cdb := range _cd.Elements() {
			if _cdb == _dag {
				_cad = append(_cad, _acb...)
			} else {
				_cad = append(_cad, _cdb)
			}
		}
		_gef.Params[0] = _cb.MakeArray(_cad...)
	}
	return nil
}

type regexMatcher struct{ _afg RedactionTerm }

func _ee(_geb *_dd.ContentStreamOperations, _ddc string, _gf int) error {
	_eb := _dd.ContentStreamOperations{}
	var _fg _dd.ContentStreamOperation
	for _eg, _eeg := range *_geb {
		if _eg == _gf {
			if _ddc == "\u0027" {
				_gac := _dd.ContentStreamOperation{Operand: "\u0054\u002a"}
				_eb = append(_eb, &_gac)
				_fg.Params = _eeg.Params
				_fg.Operand = "\u0054\u006a"
				_eb = append(_eb, &_fg)
			} else if _ddc == "\u0022" {
				_dcf := _eeg.Params[:2]
				Tc, Tw := _dcf[0], _dcf[1]
				_gcb := _dd.ContentStreamOperation{Params: []_cb.PdfObject{Tc}, Operand: "\u0054\u0063"}
				_eb = append(_eb, &_gcb)
				_gcb = _dd.ContentStreamOperation{Params: []_cb.PdfObject{Tw}, Operand: "\u0054\u0077"}
				_eb = append(_eb, &_gcb)
				_fg.Params = []_cb.PdfObject{_eeg.Params[2]}
				_fg.Operand = "\u0054\u006a"
				_eb = append(_eb, &_fg)
			}
		}
		_eb = append(_eb, _eeg)
	}
	*_geb = _eb
	return nil
}

type replacement struct {
	_ff string
	_gc float64
	_da int
}

func _fac(_bdba *_aa.TextMarkArray, _dabc int, _cgb int) int {
	_gaea := _bdba.Elements()
	_cgfc := _dabc - 1
	_gebf := _dabc + 1
	_gcea := -1
	if _cgfc >= 0 {
		_gbcd := _gaea[_cgfc]
		_ggab := _gbcd.ObjString
		_fcc := len(_ggab)
		_fcff := _gbcd.Index
		if _fcff+1 < _fcc {
			_gcea = _cgfc
			return _gcea
		}
	}
	if _gebf < len(_gaea) {
		_efc := _gaea[_gebf]
		_afc := _efc.ObjString
		if _afc[0] != _efc.Text {
			_gcea = _gebf
			return _gcea
		}
	}
	return _gcea
}
