package mdp

import (
	_ae "errors"
	_e "fmt"

	_g "unitechio/gopdf/gopdf/core"
)

func (_gbg *DiffResults) addWarningWithDescription(_cec int, _fdea string) {
	if _gbg.Warnings == nil {
		_gbg.Warnings = make([]*DiffResult, 0)
	}
	_gbg.Warnings = append(_gbg.Warnings, &DiffResult{Revision: _cec, Description: _fdea})
}

func (_ce *defaultDiffPolicy) compareRevisions(_ab *_g.PdfParser, _abe *_g.PdfParser) (*DiffResults, error) {
	var _bg error
	_ce._c, _bg = _abe.GetUpdatedObjects(_ab)
	if _bg != nil {
		return &DiffResults{}, _bg
	}
	if len(_ce._c) == 0 {
		return &DiffResults{}, nil
	}
	_fga := _abe.GetRevisionNumber()
	_gb, _de := _g.GetIndirect(_g.ResolveReference(_ab.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	_da, _ef := _g.GetIndirect(_g.ResolveReference(_abe.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	if !_de || !_ef {
		return &DiffResults{}, _ae.New("\u0065\u0072\u0072o\u0072\u0020\u0077\u0068i\u006c\u0065\u0020\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0072\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_bgc, _de := _g.GetDict(_g.ResolveReference(_gb.PdfObject))
	_efa, _ef := _g.GetDict(_g.ResolveReference(_da.PdfObject))
	if !_de || !_ef {
		return &DiffResults{}, _ae.New("\u0065\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020a\u0020\u0072\u006f\u006f\u0074\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
	}
	if _fef, _be := _g.GetIndirect(_efa.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _be {
		_ga, _bf := _g.GetDict(_fef)
		if !_bf {
			return &DiffResults{}, _ae.New("\u0065\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_bb := make([]_g.PdfObject, 0)
		if _daf, _bd := _g.GetIndirect(_bgc.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _bd {
			if _dc, _ad := _g.GetDict(_daf); _ad {
				if _feb, _cdc := _g.GetArray(_dc.Get("\u0046\u0069\u0065\u006c\u0064\u0073")); _cdc {
					_bb = _feb.Elements()
				}
			}
		}
		_ac, _bf := _g.GetArray(_ga.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_bf {
			return &DiffResults{}, _ae.New("\u0065\u0072r\u006f\u0072\u0020\u0077h\u0069\u006ce\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0066i\u0065\u006c\u0064\u0073")
		}
		if _db := _ce.compareFields(_fga, _bb, _ac.Elements()); _db != nil {
			return &DiffResults{}, _db
		}
	}
	_acd, _fa := _g.GetIndirect(_efa.Get("\u0050\u0061\u0067e\u0073"))
	if !_fa {
		return &DiffResults{}, _ae.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_ca, _fa := _g.GetIndirect(_bgc.Get("\u0050\u0061\u0067e\u0073"))
	if !_fa {
		return &DiffResults{}, _ae.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	if _ed := _ce.comparePages(_fga, _ca, _acd); _ed != nil {
		return &DiffResults{}, _ed
	}
	return _ce._f, nil
}

func (_egc *defaultDiffPolicy) comparePages(_afb int, _ega, _eda *_g.PdfIndirectObject) error {
	if _, _efc := _egc._c[_eda.ObjectNumber]; _efc {
		_egc._f.addErrorWithDescription(_afb, "\u0050a\u0067e\u0073\u0020\u0077\u0065\u0072e\u0020\u0063h\u0061\u006e\u0067\u0065\u0064")
	}
	_ggg, _baa := _g.GetDict(_eda.PdfObject)
	_gge, _fda := _g.GetDict(_ega.PdfObject)
	if !_baa || !_fda {
		return _ae.New("\u0075n\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061g\u0065\u0073\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_ffb, _baa := _g.GetArray(_ggg.Get("\u004b\u0069\u0064\u0073"))
	_ea, _fda := _g.GetArray(_gge.Get("\u004b\u0069\u0064\u0073"))
	if !_baa || !_fda {
		return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061\u0067\u0065s\u0027 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
	}
	_dbf := _ffb.Len()
	if _dbf > _ea.Len() {
		_dbf = _ea.Len()
	}
	for _babe := 0; _babe < _dbf; _babe++ {
		_cga, _dafa := _g.GetIndirect(_g.ResolveReference(_ea.Get(_babe)))
		_fgc, _gggg := _g.GetIndirect(_g.ResolveReference(_ffb.Get(_babe)))
		if !_dafa || !_gggg {
			return _ae.New("\u0075\u006e\u0065\u0078pe\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065c\u0074")
		}
		if _cga.ObjectNumber != _fgc.ObjectNumber {
			_egc._f.addErrorWithDescription(_afb, _e.Sprintf("p\u0061\u0067\u0065\u0020#%\u0064 \u0077\u0061\u0073\u0020\u0072e\u0070\u006c\u0061\u0063\u0065\u0064", _babe))
		}
		_ee, _dafa := _g.GetDict(_fgc)
		_fbg, _gggg := _g.GetDict(_cga)
		if !_dafa || !_gggg {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065'\u0073 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
		}
		_bdc, _aef := _ecd(_ee.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _aef != nil {
			return _aef
		}
		_eac, _aef := _ecd(_fbg.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _aef != nil {
			return _aef
		}
		if _ddg := _egc.compareAnnots(_afb, _eac, _bdc); _ddg != nil {
			return _ddg
		}
	}
	for _df := _dbf + 1; _df <= _ffb.Len(); _df++ {
		_egc._f.addErrorWithDescription(_afb, _e.Sprintf("\u0070a\u0067e\u0020\u0023\u0025\u0064\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _df))
	}
	for _gdb := _dbf + 1; _gdb <= _ea.Len(); _gdb++ {
		_egc._f.addErrorWithDescription(_afb, _e.Sprintf("p\u0061g\u0065\u0020\u0023\u0025\u0064\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _gdb))
	}
	return nil
}

// DocMDPPermission is values for set up access permissions for DocMDP.
// (Section 12.8.2.2, Table 254 - Entries in a signature dictionary p. 471 in PDF32000_2008).
type DocMDPPermission int64

func (_bda *DiffResults) addErrorWithDescription(_fba int, _gcf string) {
	if _bda.Errors == nil {
		_bda.Errors = make([]*DiffResult, 0)
	}
	_bda.Errors = append(_bda.Errors, &DiffResult{Revision: _fba, Description: _gcf})
}

// DiffResults describes the results of the DiffPolicy.
type DiffResults struct {
	Warnings []*DiffResult
	Errors   []*DiffResult
}

// String returns the state of the warning.
func (_fagf *DiffResult) String() string {
	return _e.Sprintf("\u0025\u0073\u0020\u0069n \u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0023\u0025\u0064", _fagf.Description, _fagf.Revision)
}

func _ecd(_gfa _g.PdfObject) ([]_g.PdfObject, error) {
	_aed := make([]_g.PdfObject, 0)
	if _gfa != nil {
		_bbc := _gfa
		if _fgb, _aa := _g.GetIndirect(_gfa); _aa {
			_bbc = _fgb.PdfObject
		}
		if _agb, _bcc := _g.GetArray(_bbc); _bcc {
			_aed = _agb.Elements()
		} else {
			return nil, _ae.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0061n\u006eo\u0074s\u0027\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
	}
	return _aed, nil
}

func (_dce *defaultDiffPolicy) compareFields(_cb int, _ag, _af []_g.PdfObject) error {
	_ge := make(map[int64]*_g.PdfObjectDictionary)
	for _, _gee := range _ag {
		_gc, _fc := _g.GetIndirect(_gee)
		if !_fc {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_agf, _fc := _g.GetDict(_gc.PdfObject)
		if !_fc {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_ge[_gc.ObjectNumber] = _agf
	}
	for _, _ada := range _af {
		_efg, _fag := _g.GetIndirect(_ada)
		if !_fag {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_cbg, _fag := _g.GetDict(_efg.PdfObject)
		if !_fag {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		T := _cbg.Get("\u0054")
		if _, _fb := _dce._c[_efg.ObjectNumber]; _fb {
			switch _dce._fe {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_dce._f.addWarningWithDescription(_cb, _e.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			default:
				_dce._f.addErrorWithDescription(_cb, _e.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			}
		}
		if _, _bag := _ge[_efg.ObjectNumber]; !_bag {
			switch _dce._fe {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_dce._f.addWarningWithDescription(_cb, _e.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _cbg.Get("\u0054")))
			default:
				_dce._f.addErrorWithDescription(_cb, _e.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _cbg.Get("\u0054")))
			}
		} else {
			delete(_ge, _efg.ObjectNumber)
			if _, _def := _dce._c[_efg.ObjectNumber]; _def {
				switch _dce._fe {
				case NoRestrictions, FillForms, FillFormsAndAnnots:
					_dce._f.addWarningWithDescription(_cb, _e.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _cbg.Get("\u0054")))
				default:
					_dce._f.addErrorWithDescription(_cb, _e.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _cbg.Get("\u0054")))
				}
			}
		}
		if FT, _efgd := _g.GetNameVal(_cbg.Get("\u0046\u0054")); _efgd {
			if FT == "\u0053\u0069\u0067" {
				if _caf, _cg := _g.GetIndirect(_cbg.Get("\u0056")); _cg {
					if _, _dd := _dce._c[_caf.ObjectNumber]; _dd {
						switch _dce._fe {
						case NoRestrictions, FillForms, FillFormsAndAnnots:
							_dce._f.addWarningWithDescription(_cb, _e.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						default:
							_dce._f.addErrorWithDescription(_cb, _e.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						}
					}
				}
			}
		}
	}
	for _, _abf := range _ge {
		switch _dce._fe {
		case NoRestrictions:
			_dce._f.addWarningWithDescription(_cb, _e.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _abf.Get("\u0054")))
		default:
			_dce._f.addErrorWithDescription(_cb, _e.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _abf.Get("\u0054")))
		}
	}
	return nil
}

// IsPermitted returns true if changes permitted.
func (_fgcf *DiffResults) IsPermitted() bool { return len(_fgcf.Errors) == 0 }

// MDPParameters describes parameters for the MDP checks (now only DocMDP).
type MDPParameters struct{ DocMDPLevel DocMDPPermission }

func NewDefaultDiffPolicy() DiffPolicy {
	return &defaultDiffPolicy{_c: nil, _f: &DiffResults{}, _fe: 0}
}

func (_efgg *DiffResults) addWarning(_bed *DiffResult) {
	if _efgg.Warnings == nil {
		_efgg.Warnings = make([]*DiffResult, 0)
	}
	_efgg.Warnings = append(_efgg.Warnings, _bed)
}

const (
	NoRestrictions     DocMDPPermission = 0
	NoChanges          DocMDPPermission = 1
	FillForms          DocMDPPermission = 2
	FillFormsAndAnnots DocMDPPermission = 3
)

// DiffResult describes the warning or the error for the DiffPolicy results.
type DiffResult struct {
	Revision    int
	Description string
}

// ReviewFile implementation of DiffPolicy interface
// The default policy only checks the next types of objects:
// Page, Pages (container for page objects), Annot, Annots (container for annotation objects), Field.
// It checks adding, removing and modifying objects of these types.
func (_eg *defaultDiffPolicy) ReviewFile(oldParser *_g.PdfParser, newParser *_g.PdfParser, params *MDPParameters) (*DiffResults, error) {
	if oldParser.GetRevisionNumber() > newParser.GetRevisionNumber() {
		return nil, _ae.New("\u006f\u006c\u0064\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u006e\u0065\u0077\u0020r\u0065\u0076i\u0073\u0069\u006f\u006e")
	}
	if oldParser.GetRevisionNumber() == newParser.GetRevisionNumber() {
		if oldParser != newParser {
			return nil, _ae.New("\u0073\u0061m\u0065\u0020\u0072\u0065v\u0069\u0073i\u006f\u006e\u0073\u002c\u0020\u0062\u0075\u0074 \u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0070\u0061r\u0073\u0065\u0072\u0073")
		}
		return &DiffResults{}, nil
	}
	if params == nil {
		_eg._fe = NoRestrictions
	} else {
		_eg._fe = params.DocMDPLevel
	}
	_b := &DiffResults{}
	for _gf := oldParser.GetRevisionNumber() + 1; _gf <= newParser.GetRevisionNumber(); _gf++ {
		_d, _fg := newParser.GetRevision(_gf - 1)
		if _fg != nil {
			return nil, _fg
		}
		_cd, _fg := newParser.GetRevision(_gf)
		if _fg != nil {
			return nil, _fg
		}
		_ba, _fg := _eg.compareRevisions(_d, _cd)
		if _fg != nil {
			return nil, _fg
		}
		_b.Warnings = append(_b.Warnings, _ba.Warnings...)
		_b.Errors = append(_b.Errors, _ba.Errors...)
	}
	return _b, nil
}

type defaultDiffPolicy struct {
	_c  map[int64]_g.PdfObject
	_f  *DiffResults
	_fe DocMDPPermission
}

// DiffPolicy interface for comparing two revisions of the Pdf document.
type DiffPolicy interface {
	// ReviewFile should check the revisions of the old and new parsers
	// and evaluate the differences between the revisions.
	// Each implementation of this interface must decide
	// how to handle cases where there are multiple revisions between the old and new revisions.
	ReviewFile(_bfg *_g.PdfParser, _bff *_g.PdfParser, _ffed *MDPParameters) (*DiffResults, error)
}

func (_cgaf *DiffResults) addError(_abg *DiffResult) {
	if _cgaf.Errors == nil {
		_cgaf.Errors = make([]*DiffResult, 0)
	}
	_cgaf.Errors = append(_cgaf.Errors, _abg)
}

func (_eef *defaultDiffPolicy) compareAnnots(_geg int, _eee, _ffbb []_g.PdfObject) error {
	_ece := make(map[int64]*_g.PdfObjectDictionary)
	for _, _ace := range _eee {
		_afe, _ggea := _g.GetIndirect(_ace)
		if !_ggea {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_geb, _ggea := _g.GetDict(_afe.PdfObject)
		if !_ggea {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_ece[_afe.ObjectNumber] = _geb
	}
	for _, _fac := range _ffbb {
		_fca, _bc := _g.GetIndirect(_fac)
		if !_bc {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_age, _bc := _g.GetDict(_fca.PdfObject)
		if !_bc {
			return _ae.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_bge, _ := _g.GetStringVal(_age.Get("\u0054"))
		_gac, _ := _g.GetNameVal(_age.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if _, _gege := _ece[_fca.ObjectNumber]; !_gege {
			switch _eef._fe {
			case NoRestrictions, FillFormsAndAnnots:
				_eef._f.addWarningWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gac, _bge))
			default:
				_faa, _fdg := _g.GetDict(_fca.PdfObject)
				if !_fdg {
					return _ae.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
				}
				_edaf, _fdg := _g.GetNameVal(_faa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
				if !_fdg {
					return _ae.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
				}
				if _edaf == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					switch _eef._fe {
					case NoRestrictions, FillFormsAndAnnots, FillForms:
						_eef._f.addWarningWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gac, _bge))
					default:
						_eef._f.addErrorWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gac, _bge))
					}
				} else {
					_eef._f.addErrorWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gac, _bge))
				}
			}
		} else {
			delete(_ece, _fca.ObjectNumber)
			if _ffe, _bae := _eef._c[_fca.ObjectNumber]; _bae {
				switch _eef._fe {
				case NoRestrictions, FillFormsAndAnnots:
					_eef._f.addWarningWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gac, _bge))
				default:
					_bgef, _adc := _g.GetIndirect(_ffe)
					if !_adc {
						return _ae.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_cda, _adc := _g.GetDict(_bgef.PdfObject)
					if !_adc {
						return _ae.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_fdb, _adc := _g.GetNameVal(_cda.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
					if !_adc {
						return _ae.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
					}
					if _fdb == "\u0057\u0069\u0064\u0067\u0065\u0074" {
						switch _eef._fe {
						case NoRestrictions, FillFormsAndAnnots, FillForms:
							_eef._f.addWarningWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gac, _bge))
						default:
							_eef._f.addErrorWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gac, _bge))
						}
					} else {
						_eef._f.addErrorWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gac, _bge))
					}
				}
			}
		}
	}
	for _, _fec := range _ece {
		_beg, _ := _g.GetStringVal(_fec.Get("\u0054"))
		_dcf, _ := _g.GetNameVal(_fec.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		switch _eef._fe {
		case NoRestrictions, FillFormsAndAnnots:
			_eef._f.addWarningWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _dcf, _beg))
		default:
			_eef._f.addErrorWithDescription(_geg, _e.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _dcf, _beg))
		}
	}
	return nil
}
