package mdp

import (
	_d "errors"
	_g "fmt"

	_dd "bitbucket.org/shenghui0779/gopdf/core"
)

func NewDefaultDiffPolicy() DiffPolicy {
	return &defaultDiffPolicy{_a: nil, _e: &DiffResults{}, _ac: 0}
}
func (_efcc *defaultDiffPolicy) comparePages(_cee int, _fdbc, _cdd *_dd.PdfIndirectObject) error {
	if _, _bfc := _efcc._a[_cdd.ObjectNumber]; _bfc {
		_efcc._e.addErrorWithDescription(_cee, "\u0050a\u0067e\u0073\u0020\u0077\u0065\u0072e\u0020\u0063h\u0061\u006e\u0067\u0065\u0064")
	}
	_gc, _bc := _dd.GetDict(_cdd.PdfObject)
	_aac, _dge := _dd.GetDict(_fdbc.PdfObject)
	if !_bc || !_dge {
		return _d.New("\u0075n\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061g\u0065\u0073\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_ea, _bc := _dd.GetArray(_gc.Get("\u004b\u0069\u0064\u0073"))
	_dga, _dge := _dd.GetArray(_aac.Get("\u004b\u0069\u0064\u0073"))
	if !_bc || !_dge {
		return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061\u0067\u0065s\u0027 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
	}
	_ad := _ea.Len()
	if _ad > _dga.Len() {
		_ad = _dga.Len()
	}
	for _cff := 0; _cff < _ad; _cff++ {
		_fga, _affa := _dd.GetIndirect(_dd.ResolveReference(_dga.Get(_cff)))
		_ggd, _aga := _dd.GetIndirect(_dd.ResolveReference(_ea.Get(_cff)))
		if !_affa || !_aga {
			return _d.New("\u0075\u006e\u0065\u0078pe\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065c\u0074")
		}
		if _fga.ObjectNumber != _ggd.ObjectNumber {
			_efcc._e.addErrorWithDescription(_cee, _g.Sprintf("p\u0061\u0067\u0065\u0020#%\u0064 \u0077\u0061\u0073\u0020\u0072e\u0070\u006c\u0061\u0063\u0065\u0064", _cff))
		}
		_bae, _affa := _dd.GetDict(_ggd)
		_bdf, _aga := _dd.GetDict(_fga)
		if !_affa || !_aga {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065'\u0073 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
		}
		_fb, _fcf := _caa(_bae.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _fcf != nil {
			return _fcf
		}
		_ec, _fcf := _caa(_bdf.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _fcf != nil {
			return _fcf
		}
		if _dfb := _efcc.compareAnnots(_cee, _ec, _fb); _dfb != nil {
			return _dfb
		}
	}
	for _agag := _ad + 1; _agag <= _ea.Len(); _agag++ {
		_efcc._e.addErrorWithDescription(_cee, _g.Sprintf("\u0070a\u0067e\u0020\u0023\u0025\u0064\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _agag))
	}
	for _cge := _ad + 1; _cge <= _dga.Len(); _cge++ {
		_efcc._e.addErrorWithDescription(_cee, _g.Sprintf("p\u0061g\u0065\u0020\u0023\u0025\u0064\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _cge))
	}
	return nil
}
func (_eeg *defaultDiffPolicy) compareFields(_ebe int, _dg, _dgf []_dd.PdfObject) error {
	_db := make(map[int64]*_dd.PdfObjectDictionary)
	for _, _aff := range _dg {
		_ce, _df := _dd.GetIndirect(_aff)
		if !_df {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_cca, _df := _dd.GetDict(_ce.PdfObject)
		if !_df {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_db[_ce.ObjectNumber] = _cca
	}
	for _, _dgd := range _dgf {
		_dbb, _fa := _dd.GetIndirect(_dgd)
		if !_fa {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_fef, _fa := _dd.GetDict(_dbb.PdfObject)
		if !_fa {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		T := _fef.Get("\u0054")
		if _, _ddc := _eeg._a[_dbb.ObjectNumber]; _ddc {
			switch _eeg._ac {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_eeg._e.addWarningWithDescription(_ebe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			default:
				_eeg._e.addErrorWithDescription(_ebe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			}
		}
		if _, _da := _db[_dbb.ObjectNumber]; !_da {
			switch _eeg._ac {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_eeg._e.addWarningWithDescription(_ebe, _g.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _fef.Get("\u0054")))
			default:
				_eeg._e.addErrorWithDescription(_ebe, _g.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _fef.Get("\u0054")))
			}
		} else {
			delete(_db, _dbb.ObjectNumber)
			if _, _cd := _eeg._a[_dbb.ObjectNumber]; _cd {
				switch _eeg._ac {
				case NoRestrictions, FillForms, FillFormsAndAnnots:
					_eeg._e.addWarningWithDescription(_ebe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _fef.Get("\u0054")))
				default:
					_eeg._e.addErrorWithDescription(_ebe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _fef.Get("\u0054")))
				}
			}
		}
		if FT, _aec := _dd.GetNameVal(_fef.Get("\u0046\u0054")); _aec {
			if FT == "\u0053\u0069\u0067" {
				if _bb, _cf := _dd.GetIndirect(_fef.Get("\u0056")); _cf {
					if _, _egd := _eeg._a[_bb.ObjectNumber]; _egd {
						switch _eeg._ac {
						case NoRestrictions, FillForms, FillFormsAndAnnots:
							_eeg._e.addWarningWithDescription(_ebe, _g.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						default:
							_eeg._e.addErrorWithDescription(_ebe, _g.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						}
					}
				}
			}
		}
	}
	for _, _ba := range _db {
		switch _eeg._ac {
		case NoRestrictions:
			_eeg._e.addWarningWithDescription(_ebe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _ba.Get("\u0054")))
		default:
			_eeg._e.addErrorWithDescription(_ebe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _ba.Get("\u0054")))
		}
	}
	return nil
}
func (_ggde *DiffResults) addError(_ece *DiffResult) {
	if _ggde.Errors == nil {
		_ggde.Errors = make([]*DiffResult, 0)
	}
	_ggde.Errors = append(_ggde.Errors, _ece)
}

// DocMDPPermission is values for set up access permissions for DocMDP.
// (Section 12.8.2.2, Table 254 - Entries in a signature dictionary p. 471 in PDF32000_2008).
type DocMDPPermission int64

// ReviewFile implementation of DiffPolicy interface
// The default policy only checks the next types of objects:
// Page, Pages (container for page objects), Annot, Annots (container for annotation objects), Field.
// It checks adding, removing and modifying objects of these types.
func (_f *defaultDiffPolicy) ReviewFile(oldParser *_dd.PdfParser, newParser *_dd.PdfParser, params *MDPParameters) (*DiffResults, error) {
	if oldParser.GetRevisionNumber() > newParser.GetRevisionNumber() {
		return nil, _d.New("\u006f\u006c\u0064\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u006e\u0065\u0077\u0020r\u0065\u0076i\u0073\u0069\u006f\u006e")
	}
	if oldParser.GetRevisionNumber() == newParser.GetRevisionNumber() {
		if oldParser != newParser {
			return nil, _d.New("\u0073\u0061m\u0065\u0020\u0072\u0065v\u0069\u0073i\u006f\u006e\u0073\u002c\u0020\u0062\u0075\u0074 \u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0070\u0061r\u0073\u0065\u0072\u0073")
		}
		return &DiffResults{}, nil
	}
	if params == nil {
		_f._ac = NoRestrictions
	} else {
		_f._ac = params.DocMDPLevel
	}
	_de := &DiffResults{}
	for _cc := oldParser.GetRevisionNumber() + 1; _cc <= newParser.GetRevisionNumber(); _cc++ {
		_cb, _dea := newParser.GetRevision(_cc - 1)
		if _dea != nil {
			return nil, _dea
		}
		_eb, _dea := newParser.GetRevision(_cc)
		if _dea != nil {
			return nil, _dea
		}
		_gg, _dea := _f.compareRevisions(_cb, _eb)
		if _dea != nil {
			return nil, _dea
		}
		_de.Warnings = append(_de.Warnings, _gg.Warnings...)
		_de.Errors = append(_de.Errors, _gg.Errors...)
	}
	return _de, nil
}

// IsPermitted returns true if changes permitted.
func (_ebga *DiffResults) IsPermitted() bool { return len(_ebga.Errors) == 0 }
func (_dgeb *DiffResults) addWarning(_be *DiffResult) {
	if _dgeb.Warnings == nil {
		_dgeb.Warnings = make([]*DiffResult, 0)
	}
	_dgeb.Warnings = append(_dgeb.Warnings, _be)
}

// DiffResult describes the warning or the error for the DiffPolicy results.
type DiffResult struct {
	Revision    int
	Description string
}

func (_bcf *DiffResults) addErrorWithDescription(_ebg int, _ggg string) {
	if _bcf.Errors == nil {
		_bcf.Errors = make([]*DiffResult, 0)
	}
	_bcf.Errors = append(_bcf.Errors, &DiffResult{Revision: _ebg, Description: _ggg})
}
func (_eg *defaultDiffPolicy) compareRevisions(_b *_dd.PdfParser, _cg *_dd.PdfParser) (*DiffResults, error) {
	var _ca error
	_eg._a, _ca = _cg.GetUpdatedObjects(_b)
	if _ca != nil {
		return &DiffResults{}, _ca
	}
	if len(_eg._a) == 0 {
		return &DiffResults{}, nil
	}
	_ef := _cg.GetRevisionNumber()
	_fg, _dec := _dd.GetIndirect(_dd.ResolveReference(_b.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	_ee, _bd := _dd.GetIndirect(_dd.ResolveReference(_cg.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	if !_dec || !_bd {
		return &DiffResults{}, _d.New("\u0065\u0072\u0072o\u0072\u0020\u0077\u0068i\u006c\u0065\u0020\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0072\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_ae, _dec := _dd.GetDict(_dd.ResolveReference(_fg.PdfObject))
	_bf, _bd := _dd.GetDict(_dd.ResolveReference(_ee.PdfObject))
	if !_dec || !_bd {
		return &DiffResults{}, _d.New("\u0065\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020a\u0020\u0072\u006f\u006f\u0074\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
	}
	if _ag, _ff := _dd.GetIndirect(_bf.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _ff {
		_ffc, _af := _dd.GetDict(_ag)
		if !_af {
			return &DiffResults{}, _d.New("\u0065\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_gb := make([]_dd.PdfObject, 0)
		if _eed, _fd := _dd.GetIndirect(_ae.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _fd {
			if _bdb, _fc := _dd.GetDict(_eed); _fc {
				if _ed, _efd := _dd.GetArray(_bdb.Get("\u0046\u0069\u0065\u006c\u0064\u0073")); _efd {
					_gb = _ed.Elements()
				}
			}
		}
		_ab, _af := _dd.GetArray(_ffc.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_af {
			return &DiffResults{}, _d.New("\u0065\u0072r\u006f\u0072\u0020\u0077h\u0069\u006ce\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0066i\u0065\u006c\u0064\u0073")
		}
		if _fge := _eg.compareFields(_ef, _gb, _ab.Elements()); _fge != nil {
			return &DiffResults{}, _fge
		}
	}
	_fdb, _fe := _dd.GetIndirect(_bf.Get("\u0050\u0061\u0067e\u0073"))
	if !_fe {
		return &DiffResults{}, _d.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_ffa, _fe := _dd.GetIndirect(_ae.Get("\u0050\u0061\u0067e\u0073"))
	if !_fe {
		return &DiffResults{}, _d.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	if _ded := _eg.comparePages(_ef, _ffa, _fdb); _ded != nil {
		return &DiffResults{}, _ded
	}
	return _eg._e, nil
}

// DiffResults describes the results of the DiffPolicy.
type DiffResults struct {
	Warnings []*DiffResult
	Errors   []*DiffResult
}

func (_acg *DiffResults) addWarningWithDescription(_affc int, _fcc string) {
	if _acg.Warnings == nil {
		_acg.Warnings = make([]*DiffResult, 0)
	}
	_acg.Warnings = append(_acg.Warnings, &DiffResult{Revision: _affc, Description: _fcc})
}

// DiffPolicy interface for comparing two revisions of the Pdf document.
type DiffPolicy interface {

	// ReviewFile should check the revisions of the old and new parsers
	// and evaluate the differences between the revisions.
	// Each implementation of this interface must decide
	// how to handle cases where there are multiple revisions between the old and new revisions.
	ReviewFile(_dgc *_dd.PdfParser, _afc *_dd.PdfParser, _bbd *MDPParameters) (*DiffResults, error)
}
type defaultDiffPolicy struct {
	_a  map[int64]_dd.PdfObject
	_e  *DiffResults
	_ac DocMDPPermission
}

// MDPParameters describes parameters for the MDP checks (now only DocMDP).
type MDPParameters struct{ DocMDPLevel DocMDPPermission }

const (
	NoRestrictions     DocMDPPermission = 0
	NoChanges          DocMDPPermission = 1
	FillForms          DocMDPPermission = 2
	FillFormsAndAnnots DocMDPPermission = 3
)

// String returns the state of the warning.
func (_ecd *DiffResult) String() string {
	return _g.Sprintf("\u0025\u0073\u0020\u0069n \u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0023\u0025\u0064", _ecd.Description, _ecd.Revision)
}
func _caa(_dc _dd.PdfObject) ([]_dd.PdfObject, error) {
	_fbg := make([]_dd.PdfObject, 0)
	if _dc != nil {
		_ged := _dc
		if _bfd, _fbd := _dd.GetIndirect(_dc); _fbd {
			_ged = _bfd.PdfObject
		}
		if _fegc, _fbf := _dd.GetArray(_ged); _fbf {
			_fbg = _fegc.Elements()
		} else {
			return nil, _d.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0061n\u006eo\u0074s\u0027\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
	}
	return _fbg, nil
}
func (_feg *defaultDiffPolicy) compareAnnots(_dfg int, _ada, _abb []_dd.PdfObject) error {
	_def := make(map[int64]*_dd.PdfObjectDictionary)
	for _, _aee := range _ada {
		_faa, _ebf := _dd.GetIndirect(_aee)
		if !_ebf {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_dee, _ebf := _dd.GetDict(_faa.PdfObject)
		if !_ebf {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_def[_faa.ObjectNumber] = _dee
	}
	for _, _dgfg := range _abb {
		_bbb, _dbc := _dd.GetIndirect(_dgfg)
		if !_dbc {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_deb, _dbc := _dd.GetDict(_bbb.PdfObject)
		if !_dbc {
			return _d.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_dfgf, _ := _dd.GetStringVal(_deb.Get("\u0054"))
		_egf, _ := _dd.GetNameVal(_deb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if _, _ceg := _def[_bbb.ObjectNumber]; !_ceg {
			switch _feg._ac {
			case NoRestrictions, FillFormsAndAnnots:
				_feg._e.addWarningWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _egf, _dfgf))
			default:
				_fefe, _ge := _dd.GetDict(_bbb.PdfObject)
				if !_ge {
					return _d.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
				}
				_afb, _ge := _dd.GetNameVal(_fefe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
				if !_ge {
					return _d.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
				}
				if _afb == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					switch _feg._ac {
					case NoRestrictions, FillFormsAndAnnots, FillForms:
						_feg._e.addWarningWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _egf, _dfgf))
					default:
						_feg._e.addErrorWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _egf, _dfgf))
					}
				} else {
					_feg._e.addErrorWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _egf, _dfgf))
				}
			}
		} else {
			delete(_def, _bbb.ObjectNumber)
			if _caf, _abbc := _feg._a[_bbb.ObjectNumber]; _abbc {
				switch _feg._ac {
				case NoRestrictions, FillFormsAndAnnots:
					_feg._e.addWarningWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _egf, _dfgf))
				default:
					_gef, _daa := _dd.GetIndirect(_caf)
					if !_daa {
						return _d.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_dgda, _daa := _dd.GetDict(_gef.PdfObject)
					if !_daa {
						return _d.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_ecg, _daa := _dd.GetNameVal(_dgda.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
					if !_daa {
						return _d.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
					}
					if _ecg == "\u0057\u0069\u0064\u0067\u0065\u0074" {
						switch _feg._ac {
						case NoRestrictions, FillFormsAndAnnots, FillForms:
							_feg._e.addWarningWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _egf, _dfgf))
						default:
							_feg._e.addErrorWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _egf, _dfgf))
						}
					} else {
						_feg._e.addErrorWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _egf, _dfgf))
					}
				}
			}
		}
	}
	for _, _bad := range _def {
		_cfa, _ := _dd.GetStringVal(_bad.Get("\u0054"))
		_dba, _ := _dd.GetNameVal(_bad.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		switch _feg._ac {
		case NoRestrictions, FillFormsAndAnnots:
			_feg._e.addWarningWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _dba, _cfa))
		default:
			_feg._e.addErrorWithDescription(_dfg, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _dba, _cfa))
		}
	}
	return nil
}
