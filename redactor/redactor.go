package redactor

import (
	_f "errors"
	_c "fmt"
	_d "io"
	_e "regexp"
	_ad "sort"
	_b "strings"

	_g "bitbucket.org/shenghui0779/gopdf/common"
	_ba "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ea "bitbucket.org/shenghui0779/gopdf/core"
	_fd "bitbucket.org/shenghui0779/gopdf/creator"
	_ec "bitbucket.org/shenghui0779/gopdf/extractor"
	_ga "bitbucket.org/shenghui0779/gopdf/model"
)

func _fdd(_dbg *_ec.TextMarkArray) *_ga.PdfFont {
	_, _gda := _eb(_dbg)
	_dcgc := _dbg.Elements()[_gda]
	_abe := _dcgc.Font
	return _abe
}
func _ffc(_ebed, _agg string) []int {
	if len(_agg) == 0 {
		return nil
	}
	var _egbg []int
	for _beda := 0; _beda < len(_ebed); {
		_eaa := _b.Index(_ebed[_beda:], _agg)
		if _eaa < 0 {
			return _egbg
		}
		_egbg = append(_egbg, _beda+_eaa)
		_beda += _eaa + len(_agg)
	}
	return _egbg
}
func _ecaa(_ebc _ea.PdfObject, _egdg *_ga.PdfFont) (string, error) {
	_bcg, _dfd := _ea.GetStringBytes(_ebc)
	if !_dfd {
		return "", _ea.ErrTypeError
	}
	_afc := _egdg.BytesToCharcodes(_bcg)
	_bcdc, _ged, _gaaa := _egdg.CharcodesToStrings(_afc)
	if _gaaa > 0 {
		_g.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _ged, _gaaa)
	}
	_eabb := _b.Join(_bcdc, "")
	return _eabb, nil
}
func _edf(_efa []int, _gcd *_ec.TextMarkArray, _fdg string) (*_ec.TextMarkArray, matchedBBox, error) {
	_bbd := matchedBBox{}
	_fea := _efa[0]
	_bbcgf := _efa[1]
	_cbgc := len(_fdg) - len(_b.TrimLeft(_fdg, "\u0020"))
	_gcca := len(_fdg) - len(_b.TrimRight(_fdg, "\u0020\u000a"))
	_fea = _fea + _cbgc
	_bbcgf = _bbcgf - _gcca
	_fdg = _b.Trim(_fdg, "\u0020\u000a")
	_fce, _ccg := _gcd.RangeOffset(_fea, _bbcgf)
	if _ccg != nil {
		return nil, _bbd, _ccg
	}
	_bcae, _fcg := _fce.BBox()
	if !_fcg {
		return nil, _bbd, _c.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _fce)
	}
	_bbd = matchedBBox{_gdd: _fdg, _afcc: _bcae}
	return _fce, _bbd, nil
}
func (_cege *Redactor) redactPage(_fba *_ba.ContentStreamOperations, _gfaf *_ga.PdfPageResources) ([]matchedBBox, *_ba.ContentStreamOperations, error) {
	_egbc, _eff := _ec.NewFromContents(_fba.String(), _gfaf)
	if _eff != nil {
		return nil, nil, _eff
	}
	_ffcd, _, _, _eff := _egbc.ExtractPageText()
	_fba = _ffcd.GetContentStreamOps()
	if _eff != nil {
		return nil, nil, _eff
	}
	_cdcc := _ffcd.Marks()
	_dfc := _ffcd.Text()
	_edd := []matchedBBox{}
	_acad := make(map[_ea.PdfObject][]localSpanMarks)
	for _, _ggbc := range _cege._bdadb.Terms {
		_dca, _cgfd := _gff(_ggbc)
		if _cgfd != nil {
			return nil, nil, _cgfd
		}
		_degf, _cgfd := _dca.match(_dfc)
		if _cgfd != nil {
			return nil, nil, _cgfd
		}
		_fdee := _abc(_degf)
		for _cdg, _cbe := range _fdee {
			_afa := []matchedBBox{}
			for _, _ggc := range _cbe {
				_ddd, _afb, _dbad := _edf(_ggc, _cdcc, _cdg)
				if _dbad != nil {
					return nil, nil, _dbad
				}
				_ggf := _ddcd(_ddd)
				for _bcga, _efea := range _ggf {
					_cfab := localSpanMarks{_dbf: _efea, _gaaf: _bcga, _aca: _cdg}
					_gga, _ := _eb(_efea)
					if _gbfe, _cad := _acad[_gga]; _cad {
						_acad[_gga] = append(_gbfe, _cfab)
					} else {
						_acad[_gga] = []localSpanMarks{_cfab}
					}
				}
				_afa = append(_afa, _afb)
			}
			_edd = append(_edd, _afa...)
		}
	}
	_eff = _cf(_fba, _acad)
	if _eff != nil {
		return nil, nil, _eff
	}
	return _edd, _fba, nil
}
func _dcgf(_cda []localSpanMarks) (map[string][]localSpanMarks, []string) {
	_bge := make(map[string][]localSpanMarks)
	_gfed := []string{}
	for _, _gdb := range _cda {
		_gbg := _gdb._aca
		if _fae, _ace := _bge[_gbg]; _ace {
			_bge[_gbg] = append(_fae, _gdb)
		} else {
			_bge[_gbg] = []localSpanMarks{_gdb}
			_gfed = append(_gfed, _gbg)
		}
	}
	return _bge, _gfed
}

type matchedIndex struct {
	_cbdc int
	_gbeg int
	_fbf  string
}

func _bd(_acd *_ga.PdfFont, _bfg _ec.TextMark) float64 {
	_baf := 0.001
	_eeg := _bfg.Th / 100
	if _acd.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_baf = 1
	}
	_feb, _gbf := _acd.GetRuneMetrics(' ')
	if !_gbf {
		_feb, _gbf = _acd.GetCharMetrics(32)
	}
	if !_gbf {
		_feb, _ = _ga.DefaultFont().GetRuneMetrics(' ')
	}
	_cg := _baf * ((_feb.Wx*_bfg.FontSize + _bfg.Tc + _bfg.Tw) / _eeg)
	return _cg
}

// New instantiates a Redactor object with given PdfReader and `regex` pattern.
func New(reader *_ga.PdfReader, opts *RedactionOptions, rectProps *RectangleProps) *Redactor {
	if rectProps == nil {
		rectProps = RedactRectanglePropsNew()
	}
	return &Redactor{_dag: reader, _bdadb: opts, _dgf: _fd.New(), _abd: rectProps}
}
func _bcge(_gecd *_ec.TextMarkArray, _fcc int, _edb int) int {
	_ecd := _gecd.Elements()
	_bbb := _fcc - 1
	_dcgfb := _fcc + 1
	_cec := -1
	if _bbb >= 0 {
		_ggae := _ecd[_bbb]
		_dgd := _ggae.ObjString
		_cca := len(_dgd)
		_aga := _ggae.Index
		if _aga+1 < _cca {
			_cec = _bbb
			return _cec
		}
	}
	if _dcgfb < len(_ecd) {
		_dbdg := _ecd[_dcgfb]
		_bgad := _dbdg.ObjString
		if _bgad[0] != _dbdg.Text {
			_cec = _dcgfb
			return _cec
		}
	}
	return _cec
}

// RedactRectanglePropsNew return a new pointer to a default RectangleProps object.
func RedactRectanglePropsNew() *RectangleProps {
	return &RectangleProps{FillColor: _fd.ColorBlack, BorderWidth: 0.0, FillOpacity: 1.0}
}
func _cgb(_dfb []placeHolders) []replacement {
	_agd := []replacement{}
	for _, _ebag := range _dfb {
		_bgd := _ebag._eg
		_gbgb := _ebag._gf
		_cge := _ebag._ff
		for _, _bde := range _bgd {
			_dac := replacement{_dc: _gbgb, _bf: _cge, _dg: _bde}
			_agd = append(_agd, _dac)
		}
	}
	_ad.Slice(_agd, func(_bdaf, _adbf int) bool { return _agd[_bdaf]._dg < _agd[_adbf]._dg })
	return _agd
}
func _fag(_ac *_ba.ContentStreamOperation, _ee _ea.PdfObject, _gdg []localSpanMarks) error {
	_dee, _ef := _ea.GetArray(_ac.Params[0])
	_cfg := []_ea.PdfObject{}
	if !_ef {
		_g.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _ac)
		return _c.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _ac)
	}
	_fdc, _age := _dcgf(_gdg)
	if len(_age) == 1 {
		_cc := _age[0]
		_bgf := _fdc[_cc]
		if len(_bgf) == 1 {
			_fgd := _bgf[0]
			_df := _fgd._dbf
			_fff := _fdd(_df)
			_cebc, _bcc := _ecaa(_ee, _fff)
			if _bcc != nil {
				return _bcc
			}
			_adb, _bcc := _ced(_fgd, _df, _fff, _cebc, _cc)
			if _bcc != nil {
				return _bcc
			}
			for _, _aged := range _dee.Elements() {
				if _aged == _ee {
					_cfg = append(_cfg, _adb...)
				} else {
					_cfg = append(_cfg, _aged)
				}
			}
		} else {
			_da := _bgf[0]._dbf
			_cee := _fdd(_da)
			_ecg, _cdd := _ecaa(_ee, _cee)
			if _cdd != nil {
				return _cdd
			}
			_egd, _cdd := _bcb(_ecg, _bgf)
			if _cdd != nil {
				return _cdd
			}
			_ed := _cgb(_egd)
			_deg := _fbg(_ecg, _ed, _cee)
			for _, _gag := range _dee.Elements() {
				if _gag == _ee {
					_cfg = append(_cfg, _deg...)
				} else {
					_cfg = append(_cfg, _gag)
				}
			}
		}
		_ac.Params[0] = _ea.MakeArray(_cfg...)
	} else if len(_age) > 1 {
		_dd := _gdg[0]
		_fec := _dd._dbf
		_, _af := _eb(_fec)
		_gcc := _fec.Elements()[_af]
		_adc := _gcc.Font
		_bb, _eba := _ecaa(_ee, _adc)
		if _eba != nil {
			return _eba
		}
		_dab, _eba := _bcb(_bb, _gdg)
		if _eba != nil {
			return _eba
		}
		_gec := _cgb(_dab)
		_bfd := _fbg(_bb, _gec, _adc)
		for _, _fc := range _dee.Elements() {
			if _fc == _ee {
				_cfg = append(_cfg, _bfd...)
			} else {
				_cfg = append(_cfg, _fc)
			}
		}
		_ac.Params[0] = _ea.MakeArray(_cfg...)
	}
	return nil
}
func _abb(_cdb, _ccbb, _cbd float64) float64 {
	_cbd = _cbd / 100
	_eggg := (-1000 * _cdb) / (_ccbb * _cbd)
	return _eggg
}

// RedactionOptions is a collection of RedactionTerm objects.
type RedactionOptions struct{ Terms []RedactionTerm }

func _cf(_gfe *_ba.ContentStreamOperations, _gaa map[_ea.PdfObject][]localSpanMarks) error {
	for _ag, _cd := range _gaa {
		if _ag == nil {
			continue
		}
		_fg, _cfa, _ge := _ecgb(_gfe, _ag)
		if !_ge {
			_g.Log.Debug("Pd\u0066\u004fb\u006a\u0065\u0063\u0074\u0020\u0025\u0073\u006e\u006ft\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0069\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073\u0074r\u0065a\u006d\u0020\u006f\u0070\u0065\u0072\u0061\u0074i\u006fn\u0020\u0025s", _ag, _gfe)
			return nil
		}
		if _fg.Operand == "\u0054\u006a" {
			_ce := _bda(_fg, _ag, _cd)
			if _ce != nil {
				return _ce
			}
		} else if _fg.Operand == "\u0054\u004a" {
			_be := _fag(_fg, _ag, _cd)
			if _be != nil {
				return _be
			}
		} else if _fg.Operand == "\u0027" || _fg.Operand == "\u0022" {
			_gc := _cb(_gfe, _fg.Operand, _cfa)
			if _gc != nil {
				return _gc
			}
			_gc = _bda(_fg, _ag, _cd)
			if _gc != nil {
				return _gc
			}
		}
	}
	return nil
}
func _ca(_gdgg string, _dcg *_ga.PdfFont) []byte {
	_acc, _db := _dcg.StringToCharcodeBytes(_gdgg)
	if _db != 0 {
		_g.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0073\u006fm\u0065\u0020\u0072un\u0065\u0073\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065d\u002e\u000a\u0009\u0025\u0073\u0020\u002d\u003e \u0025\u0076", _gdgg, _acc)
	}
	return _acc
}
func _ddc(_cddb *_ec.TextMarkArray) string {
	_aaf := ""
	for _, _bcaf := range _cddb.Elements() {
		_aaf += _bcaf.Text
	}
	return _aaf
}
func _gff(_gccf RedactionTerm) (*regexMatcher, error) { return &regexMatcher{_abdd: _gccf}, nil }
func _ddcd(_adbc *_ec.TextMarkArray) []*_ec.TextMarkArray {
	_ggbd := _adbc.Elements()
	_abg := len(_ggbd)
	var _abbd _ea.PdfObject
	_fggc := []*_ec.TextMarkArray{}
	_fgdd := &_ec.TextMarkArray{}
	_cedf := -1
	for _eag, _beb := range _ggbd {
		_gefc := _beb.DirectObject
		_cedf = _beb.Index
		if _gefc == nil {
			_dfdb := _bcge(_adbc, _eag, _cedf)
			if _abbd != nil {
				if _dfdb == -1 || _dfdb > _eag {
					_fggc = append(_fggc, _fgdd)
					_fgdd = &_ec.TextMarkArray{}
				}
			}
		} else if _gefc != nil && _abbd == nil {
			if _cedf == 0 && _eag > 0 {
				_fggc = append(_fggc, _fgdd)
				_fgdd = &_ec.TextMarkArray{}
			}
		} else if _gefc != nil && _abbd != nil {
			if _gefc != _abbd {
				_fggc = append(_fggc, _fgdd)
				_fgdd = &_ec.TextMarkArray{}
			}
		}
		_abbd = _gefc
		_fgdd.Append(_beb)
		if _eag == (_abg - 1) {
			_fggc = append(_fggc, _fgdd)
		}
	}
	return _fggc
}
func _ced(_bfc localSpanMarks, _bcd *_ec.TextMarkArray, _gba *_ga.PdfFont, _ffe, _acbf string) ([]_ea.PdfObject, error) {
	_fagc := _ddc(_bcd)
	Tj, _fb := _adaaf(_bcd)
	if _fb != nil {
		return nil, _fb
	}
	_fbd := len(_ffe)
	_eca := len(_fagc)
	_bgea := -1
	_gcb := _ea.MakeFloat(Tj)
	if _fagc != _acbf {
		_efc := _bfc._gaaf
		if _efc == 0 {
			_bgea = _b.LastIndex(_ffe, _fagc)
		} else {
			_bgea = _b.Index(_ffe, _fagc)
		}
	} else {
		_bgea = _b.Index(_ffe, _fagc)
	}
	_cgf := _bgea + _eca
	_deb := []_ea.PdfObject{}
	if _bgea == 0 && _cgf == _fbd {
		_deb = append(_deb, _gcb)
	} else if _bgea == 0 && _cgf < _fbd {
		_bca := _ca(_ffe[_cgf:], _gba)
		_cdf := _ea.MakeStringFromBytes(_bca)
		_deb = append(_deb, _gcb, _cdf)
	} else if _bgea > 0 && _cgf >= _fbd {
		_faea := _ca(_ffe[:_bgea], _gba)
		_gbe := _ea.MakeStringFromBytes(_faea)
		_deb = append(_deb, _gbe, _gcb)
	} else if _bgea > 0 && _cgf < _fbd {
		_cefd := _ca(_ffe[:_bgea], _gba)
		_fbc := _ca(_ffe[_cgf:], _gba)
		_efe := _ea.MakeStringFromBytes(_cefd)
		_befd := _ea.MakeString(string(_fbc))
		_deb = append(_deb, _efe, _gcb, _befd)
	}
	return _deb, nil
}
func _ecgb(_bff *_ba.ContentStreamOperations, PdfObj _ea.PdfObject) (*_ba.ContentStreamOperation, int, bool) {
	for _fgf, _ebe := range *_bff {
		_dcb := _ebe.Operand
		if _dcb == "\u0054\u006a" {
			_daaf := _ea.TraceToDirectObject(_ebe.Params[0])
			if _daaf == PdfObj {
				return _ebe, _fgf, true
			}
		} else if _dcb == "\u0054\u004a" {
			_dec, _bbg := _ea.GetArray(_ebe.Params[0])
			if !_bbg {
				return nil, _fgf, _bbg
			}
			for _, _bag := range _dec.Elements() {
				if _bag == PdfObj {
					return _ebe, _fgf, true
				}
			}
		} else if _dcb == "\u0022" {
			_gfa := _ea.TraceToDirectObject(_ebe.Params[2])
			if _gfa == PdfObj {
				return _ebe, _fgf, true
			}
		} else if _dcb == "\u0027" {
			_bba := _ea.TraceToDirectObject(_ebe.Params[0])
			if _bba == PdfObj {
				return _ebe, _fgf, true
			}
		}
	}
	return nil, -1, false
}
func _adaaf(_dad *_ec.TextMarkArray) (float64, error) {
	_cfgd, _fab := _dad.BBox()
	if !_fab {
		return 0.0, _c.Errorf("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073", _dad)
	}
	_bafd := _fa(_dad)
	_dbd := 0.0
	_, _fge := _eb(_dad)
	_aea := _dad.Elements()[_fge]
	_egg := _aea.Font
	if _bafd > 0 {
		_dbd = _bd(_egg, _aea)
	}
	_fcb := (_cfgd.Urx - _cfgd.Llx)
	_fcb = _fcb + _dbd*float64(_bafd)
	Tj := _abb(_fcb, _aea.FontSize, _aea.Th)
	return Tj, nil
}

type placeHolders struct {
	_eg []int
	_gf string
	_ff float64
}

// RedactionTerm holds the regexp pattern and the replacement string for the redaction process.
type RedactionTerm struct{ Pattern *_e.Regexp }

// RectangleProps defines properties of the redaction rectangle to be drawn.
type RectangleProps struct {
	FillColor   _fd.Color
	BorderWidth float64
	FillOpacity float64
}

func _bda(_egb *_ba.ContentStreamOperation, _daa _ea.PdfObject, _acb []localSpanMarks) error {
	var _dbe *_ea.PdfObjectArray
	_cbgf, _edc := _dcgf(_acb)
	if len(_edc) == 1 {
		_dfa := _edc[0]
		_aa := _cbgf[_dfa]
		if len(_aa) == 1 {
			_bga := _aa[0]
			_eae := _bga._dbf
			_bef := _fdd(_eae)
			_faga, _fgg := _ecaa(_daa, _bef)
			if _fgg != nil {
				return _fgg
			}
			_ebb, _fgg := _ced(_bga, _eae, _bef, _faga, _dfa)
			if _fgg != nil {
				return _fgg
			}
			_dbe = _ea.MakeArray(_ebb...)
		} else {
			_egc := _aa[0]._dbf
			_bdf := _fdd(_egc)
			_agb, _ccb := _ecaa(_daa, _bdf)
			if _ccb != nil {
				return _ccb
			}
			_daf, _ccb := _bcb(_agb, _aa)
			if _ccb != nil {
				return _ccb
			}
			_fda := _cgb(_daf)
			_gbc := _fbg(_agb, _fda, _bdf)
			_dbe = _ea.MakeArray(_gbc...)
		}
	} else if len(_edc) > 1 {
		_bcce := _acb[0]
		_dabe := _bcce._dbf
		_, _egdd := _eb(_dabe)
		_cbc := _dabe.Elements()[_egdd]
		_bdad := _cbc.Font
		_gfg, _gadb := _ecaa(_daa, _bdad)
		if _gadb != nil {
			return _gadb
		}
		_cff, _gadb := _bcb(_gfg, _acb)
		if _gadb != nil {
			return _gadb
		}
		_dcc := _cgb(_cff)
		_eec := _fbg(_gfg, _dcc, _bdad)
		_dbe = _ea.MakeArray(_eec...)
	}
	_egb.Params[0] = _dbe
	_egb.Operand = "\u0054\u004a"
	return nil
}

// Redact executes the redact operation on a pdf file and updates the content streams of all pages of the file.
func (_beg *Redactor) Redact() error {
	_eegf, _ggb := _beg._dag.GetNumPages()
	if _ggb != nil {
		return _c.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0074\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020P\u0061\u0067\u0065\u0073")
	}
	_abbb := _beg._abd.FillColor
	_efcc := _beg._abd.BorderWidth
	_fafc := _beg._abd.FillOpacity
	for _aad := 1; _aad <= _eegf; _aad++ {
		_afe, _baa := _beg._dag.GetPage(_aad)
		if _baa != nil {
			return _baa
		}
		_ebg, _baa := _ec.New(_afe)
		if _baa != nil {
			return _baa
		}
		_fbe, _, _, _baa := _ebg.ExtractPageText()
		if _baa != nil {
			return _baa
		}
		_efcg := _fbe.GetContentStreamOps()
		_agdf, _gbeb, _baa := _beg.redactPage(_efcg, _afe.Resources)
		if _gbeb == nil {
			_g.Log.Info("N\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u0020t\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065d \u0070\u0061\u0074t\u0061r\u006e\u002e")
			_gbeb = _efcg
		}
		_afe.SetContentStreams([]string{_gbeb.String()}, _ea.NewFlateEncoder())
		if _baa != nil {
			return _baa
		}
		_aagg, _baa := _afe.GetMediaBox()
		if _baa != nil {
			return _baa
		}
		if _afe.MediaBox == nil {
			_afe.MediaBox = _aagg
		}
		if _bcdcg := _beg._dgf.AddPage(_afe); _bcdcg != nil {
			return _bcdcg
		}
		_ebaa := _aagg.Ury
		for _, _fga := range _agdf {
			_gac := _fga._afcc
			_dba := _beg._dgf.NewRectangle(_gac.Llx, _ebaa-_gac.Lly, _gac.Urx-_gac.Llx, -(_gac.Ury - _gac.Lly))
			_dba.SetFillColor(_abbb)
			_dba.SetBorderWidth(_efcc)
			_dba.SetFillOpacity(_fafc)
			if _cgee := _beg._dgf.Draw(_dba); _cgee != nil {
				return nil
			}
		}
	}
	_beg._dgf.SetOutlineTree(_beg._dag.GetOutlineTree())
	return nil
}

// WriteToFile writes the redacted document to file specified by `outputPath`.
func (_dcd *Redactor) WriteToFile(outputPath string) error {
	if _afd := _dcd._dgf.WriteToFile(outputPath); _afd != nil {
		return _c.Errorf("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0077\u0072\u0069\u0074\u0065\u0020t\u0068\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0066\u0069\u006c\u0065")
	}
	return nil
}

type matchedBBox struct {
	_afcc _ga.PdfRectangle
	_gdd  string
}

func _cb(_gb *_ba.ContentStreamOperations, _cfc string, _ab int) error {
	_bc := _ba.ContentStreamOperations{}
	var _cbg _ba.ContentStreamOperation
	for _cdc, _bg := range *_gb {
		if _cdc == _ab {
			if _cfc == "\u0027" {
				_de := _ba.ContentStreamOperation{Operand: "\u0054\u002a"}
				_bc = append(_bc, &_de)
				_cbg.Params = _bg.Params
				_cbg.Operand = "\u0054\u006a"
				_bc = append(_bc, &_cbg)
			} else if _cfc == "\u0022" {
				_cef := _bg.Params[:2]
				Tc, Tw := _cef[0], _cef[1]
				_bad := _ba.ContentStreamOperation{Params: []_ea.PdfObject{Tc}, Operand: "\u0054\u0063"}
				_bc = append(_bc, &_bad)
				_bad = _ba.ContentStreamOperation{Params: []_ea.PdfObject{Tw}, Operand: "\u0054\u0077"}
				_bc = append(_bc, &_bad)
				_cbg.Params = []_ea.PdfObject{_bg.Params[2]}
				_cbg.Operand = "\u0054\u006a"
				_bc = append(_bc, &_cbg)
			}
		}
		_bc = append(_bc, _bg)
	}
	*_gb = _bc
	return nil
}
func _fbg(_bee string, _aee []replacement, _acef *_ga.PdfFont) []_ea.PdfObject {
	_bbcd := []_ea.PdfObject{}
	_ada := 0
	_edcb := _bee
	for _fee, _ddf := range _aee {
		_ecga := _ddf._dg
		_adaa := _ddf._bf
		_efb := _ddf._dc
		_gee := _ea.MakeFloat(_adaa)
		_beec := _bee[_ada:_ecga]
		_dea := _ca(_beec, _acef)
		_bbe := _ea.MakeStringFromBytes(_dea)
		_bbcd = append(_bbcd, _bbe)
		_bbcd = append(_bbcd, _gee)
		_bgc := _ecga + len(_efb)
		_edcb = _bee[_bgc:]
		_ada = _bgc
		if _fee == len(_aee)-1 {
			_dea = _ca(_edcb, _acef)
			_bbe = _ea.MakeStringFromBytes(_dea)
			_bbcd = append(_bbcd, _bbe)
		}
	}
	return _bbcd
}
func (_cbeb *regexMatcher) match(_bcdd string) ([]*matchedIndex, error) {
	_cedbb := _cbeb._abdd.Pattern
	if _cedbb == nil {
		return nil, _f.New("\u006e\u006f\u0020\u0070at\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006d\u0070\u0069\u006c\u0065\u0064")
	}
	var (
		_cgba = _cedbb.FindAllStringIndex(_bcdd, -1)
		_dcae []*matchedIndex
	)
	for _, _cgdd := range _cgba {
		_dcae = append(_dcae, &matchedIndex{_cbdc: _cgdd[0], _gbeg: _cgdd[1], _fbf: _bcdd[_cgdd[0]:_cgdd[1]]})
	}
	return _dcae, nil
}

type regexMatcher struct{ _abdd RedactionTerm }

func _eb(_ceb *_ec.TextMarkArray) (_ea.PdfObject, int) {
	var _fe _ea.PdfObject
	_bed := -1
	for _gad, _ece := range _ceb.Elements() {
		_fe = _ece.DirectObject
		_bed = _gad
		if _fe != nil {
			break
		}
	}
	return _fe, _bed
}

type replacement struct {
	_dc string
	_bf float64
	_dg int
}
type localSpanMarks struct {
	_dbf  *_ec.TextMarkArray
	_gaaf int
	_aca  string
}

func _fa(_gd *_ec.TextMarkArray) int {
	_eab := 0
	_fde := _gd.Elements()
	if _fde[0].Text == "\u0020" {
		_eab++
	}
	if _fde[_gd.Len()-1].Text == "\u0020" {
		_eab++
	}
	return _eab
}
func _bcb(_fecg string, _aag []localSpanMarks) ([]placeHolders, error) {
	_gg := ""
	_ceg := []placeHolders{}
	for _efg, _ecb := range _aag {
		_cdcg := _ecb._dbf
		_egf := _ecb._aca
		_ae := _ddc(_cdcg)
		_ccd, _aba := _adaaf(_cdcg)
		if _aba != nil {
			return nil, _aba
		}
		if _ae != _gg {
			var _bdb []int
			if _efg == 0 && _egf != _ae {
				_dcf := _b.Index(_fecg, _ae)
				_bdb = []int{_dcf}
			} else if _efg == len(_aag)-1 {
				_gadg := _b.LastIndex(_fecg, _ae)
				_bdb = []int{_gadg}
			} else {
				_bdb = _ffc(_fecg, _ae)
			}
			_dafe := placeHolders{_eg: _bdb, _gf: _ae, _ff: _ccd}
			_ceg = append(_ceg, _dafe)
		}
		_gg = _ae
	}
	return _ceg, nil
}

// Redactor represtents a Redactor object.
type Redactor struct {
	_dag   *_ga.PdfReader
	_bdadb *RedactionOptions
	_dgf   *_fd.Creator
	_abd   *RectangleProps
}

func _abc(_cgd []*matchedIndex) map[string][][]int {
	_fcbb := make(map[string][][]int)
	for _, _dbb := range _cgd {
		_ccdb := _dbb._fbf
		_cedb := []int{_dbb._cbdc, _dbb._gbeg}
		if _ede, _bec := _fcbb[_ccdb]; _bec {
			_fcbb[_ccdb] = append(_ede, _cedb)
		} else {
			_fcbb[_ccdb] = [][]int{_cedb}
		}
	}
	return _fcbb
}

// Write writes the content of `re.creator` to writer of type io.Writer interface.
func (_gcf *Redactor) Write(writer _d.Writer) error { return _gcf._dgf.Write(writer) }
