package mdp

import (
	_e "errors"
	_g "fmt"

	_c "bitbucket.org/shenghui0779/gopdf/core"
)

// IsPermitted returns true if changes permitted.
func (_gde *DiffResults) IsPermitted() bool { return len(_gde.Errors) == 0 }
func (_bb *DiffResults) addErrorWithDescription(_afg int, _gaa string) {
	if _bb.Errors == nil {
		_bb.Errors = make([]*DiffResult, 0)
	}
	_bb.Errors = append(_bb.Errors, &DiffResult{Revision: _afg, Description: _gaa})
}

// String returns the state of the warning.
func (_gea *DiffResult) String() string {
	return _g.Sprintf("\u0025\u0073\u0020\u0069n \u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0023\u0025\u0064", _gea.Description, _gea.Revision)
}

// DiffResult describes the warning or the error for the DiffPolicy results.
type DiffResult struct {
	Revision    int
	Description string
}

// ReviewFile implementation of DiffPolicy interface
// The default policy only checks the next types of objects:
// Page, Pages (container for page objects), Annot, Annots (container for annotation objects), Field.
// It checks adding, removing and modifying objects of these types.
func (_cb *defaultDiffPolicy) ReviewFile(oldParser *_c.PdfParser, newParser *_c.PdfParser, params *MDPParameters) (*DiffResults, error) {
	if oldParser.GetRevisionNumber() > newParser.GetRevisionNumber() {
		return nil, _e.New("\u006f\u006c\u0064\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u006e\u0065\u0077\u0020r\u0065\u0076i\u0073\u0069\u006f\u006e")
	}
	if oldParser.GetRevisionNumber() == newParser.GetRevisionNumber() {
		if oldParser != newParser {
			return nil, _e.New("\u0073\u0061m\u0065\u0020\u0072\u0065v\u0069\u0073i\u006f\u006e\u0073\u002c\u0020\u0062\u0075\u0074 \u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0070\u0061r\u0073\u0065\u0072\u0073")
		}
		return &DiffResults{}, nil
	}
	if params == nil {
		_cb._d = NoRestrictions
	} else {
		_cb._d = params.DocMDPLevel
	}
	_a := &DiffResults{}
	for _ag := oldParser.GetRevisionNumber() + 1; _ag <= newParser.GetRevisionNumber(); _ag++ {
		_cf, _fg := newParser.GetRevision(_ag - 1)
		if _fg != nil {
			return nil, _fg
		}
		_b, _fg := newParser.GetRevision(_ag)
		if _fg != nil {
			return nil, _fg
		}
		_ge, _fg := _cb.compareRevisions(_cf, _b)
		if _fg != nil {
			return nil, _fg
		}
		_a.Warnings = append(_a.Warnings, _ge.Warnings...)
		_a.Errors = append(_a.Errors, _ge.Errors...)
	}
	return _a, nil
}

// MDPParameters describes parameters for the MDP checks (now only DocMDP).
type MDPParameters struct{ DocMDPLevel DocMDPPermission }

func NewDefaultDiffPolicy() DiffPolicy {
	return &defaultDiffPolicy{_ef: nil, _eg: &DiffResults{}, _d: 0}
}

// DiffResults describes the results of the DiffPolicy.
type DiffResults struct {
	Warnings []*DiffResult
	Errors   []*DiffResult
}

func (_gb *defaultDiffPolicy) comparePages(_ggf int, _fgd, _dgg *_c.PdfIndirectObject) error {
	if _, _ggfd := _gb._ef[_dgg.ObjectNumber]; _ggfd {
		_gb._eg.addErrorWithDescription(_ggf, "\u0050a\u0067e\u0073\u0020\u0077\u0065\u0072e\u0020\u0063h\u0061\u006e\u0067\u0065\u0064")
	}
	_aef, _ecd := _c.GetDict(_dgg.PdfObject)
	_fgg, _bde := _c.GetDict(_fgd.PdfObject)
	if !_ecd || !_bde {
		return _e.New("\u0075n\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061g\u0065\u0073\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_bgd, _ecd := _c.GetArray(_aef.Get("\u004b\u0069\u0064\u0073"))
	_aa, _bde := _c.GetArray(_fgg.Get("\u004b\u0069\u0064\u0073"))
	if !_ecd || !_bde {
		return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061\u0067\u0065s\u0027 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
	}
	_ce := _bgd.Len()
	if _ce > _aa.Len() {
		_ce = _aa.Len()
	}
	for _gc := 0; _gc < _ce; _gc++ {
		_bab, _dbg := _c.GetIndirect(_c.ResolveReference(_aa.Get(_gc)))
		_bfda, _aeb := _c.GetIndirect(_c.ResolveReference(_bgd.Get(_gc)))
		if !_dbg || !_aeb {
			return _e.New("\u0075\u006e\u0065\u0078pe\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065c\u0074")
		}
		if _bab.ObjectNumber != _bfda.ObjectNumber {
			_gb._eg.addErrorWithDescription(_ggf, _g.Sprintf("p\u0061\u0067\u0065\u0020#%\u0064 \u0077\u0061\u0073\u0020\u0072e\u0070\u006c\u0061\u0063\u0065\u0064", _gc))
		}
		_ceg, _dbg := _c.GetDict(_bfda)
		_bdb, _aeb := _c.GetDict(_bab)
		if !_dbg || !_aeb {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065'\u0073 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
		}
		_ade, _cegf := _agb(_ceg.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _cegf != nil {
			return _cegf
		}
		_fgbc, _cegf := _agb(_bdb.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _cegf != nil {
			return _cegf
		}
		if _ggd := _gb.compareAnnots(_ggf, _fgbc, _ade); _ggd != nil {
			return _ggd
		}
	}
	for _ebf := _ce + 1; _ebf <= _bgd.Len(); _ebf++ {
		_gb._eg.addErrorWithDescription(_ggf, _g.Sprintf("\u0070a\u0067e\u0020\u0023\u0025\u0064\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _ebf))
	}
	for _fa := _ce + 1; _fa <= _aa.Len(); _fa++ {
		_gb._eg.addErrorWithDescription(_ggf, _g.Sprintf("p\u0061g\u0065\u0020\u0023\u0025\u0064\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _fa))
	}
	return nil
}
func (_gce *DiffResults) addWarning(_gfge *DiffResult) {
	if _gce.Warnings == nil {
		_gce.Warnings = make([]*DiffResult, 0)
	}
	_gce.Warnings = append(_gce.Warnings, _gfge)
}
func (_fbd *DiffResults) addError(_gfa *DiffResult) {
	if _fbd.Errors == nil {
		_fbd.Errors = make([]*DiffResult, 0)
	}
	_fbd.Errors = append(_fbd.Errors, _gfa)
}

// DiffPolicy interface for comparing two revisions of the Pdf document.
type DiffPolicy interface {

	// ReviewFile should check the revisions of the old and new parsers
	// and evaluate the differences between the revisions.
	// Each implementation of this interface must decide
	// how to handle cases where there are multiple revisions between the old and new revisions.
	ReviewFile(_fff *_c.PdfParser, _geaf *_c.PdfParser, _ffc *MDPParameters) (*DiffResults, error)
}

// DocMDPPermission is values for set up access permissions for DocMDP.
// (Section 12.8.2.2, Table 254 - Entries in a signature dictionary p. 471 in PDF32000_2008).
type DocMDPPermission int64
type defaultDiffPolicy struct {
	_ef map[int64]_c.PdfObject
	_eg *DiffResults
	_d  DocMDPPermission
}

func (_cg *defaultDiffPolicy) compareRevisions(_bg *_c.PdfParser, _cff *_c.PdfParser) (*DiffResults, error) {
	var _fc error
	_cg._ef, _fc = _cff.GetUpdatedObjects(_bg)
	if _fc != nil {
		return &DiffResults{}, _fc
	}
	if len(_cg._ef) == 0 {
		return &DiffResults{}, nil
	}
	_geg := _cff.GetRevisionNumber()
	_fd, _ac := _c.GetIndirect(_c.ResolveReference(_bg.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	_gee, _eb := _c.GetIndirect(_c.ResolveReference(_cff.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	if !_ac || !_eb {
		return &DiffResults{}, _e.New("\u0065\u0072\u0072o\u0072\u0020\u0077\u0068i\u006c\u0065\u0020\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0072\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_bf, _ac := _c.GetDict(_c.ResolveReference(_fd.PdfObject))
	_fcg, _eb := _c.GetDict(_c.ResolveReference(_gee.PdfObject))
	if !_ac || !_eb {
		return &DiffResults{}, _e.New("\u0065\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020a\u0020\u0072\u006f\u006f\u0074\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
	}
	if _dg, _fde := _c.GetIndirect(_fcg.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _fde {
		_fdef, _bff := _c.GetDict(_dg)
		if !_bff {
			return &DiffResults{}, _e.New("\u0065\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_gf := make([]_c.PdfObject, 0)
		if _ee, _bfd := _c.GetIndirect(_bf.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _bfd {
			if _bfc, _fdg := _c.GetDict(_ee); _fdg {
				if _be, _de := _c.GetArray(_bfc.Get("\u0046\u0069\u0065\u006c\u0064\u0073")); _de {
					_gf = _be.Elements()
				}
			}
		}
		_fca, _bff := _c.GetArray(_fdef.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_bff {
			return &DiffResults{}, _e.New("\u0065\u0072r\u006f\u0072\u0020\u0077h\u0069\u006ce\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0066i\u0065\u006c\u0064\u0073")
		}
		if _ba := _cg.compareFields(_geg, _gf, _fca.Elements()); _ba != nil {
			return &DiffResults{}, _ba
		}
	}
	_ff, _efa := _c.GetIndirect(_fcg.Get("\u0050\u0061\u0067e\u0073"))
	if !_efa {
		return &DiffResults{}, _e.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_dd, _efa := _c.GetIndirect(_bf.Get("\u0050\u0061\u0067e\u0073"))
	if !_efa {
		return &DiffResults{}, _e.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	if _cgd := _cg.comparePages(_geg, _dd, _ff); _cgd != nil {
		return &DiffResults{}, _cgd
	}
	return _cg._eg, nil
}

const (
	NoRestrictions     DocMDPPermission = 0
	NoChanges          DocMDPPermission = 1
	FillForms          DocMDPPermission = 2
	FillFormsAndAnnots DocMDPPermission = 3
)

func (_gd *DiffResults) addWarningWithDescription(_ffe int, _eff string) {
	if _gd.Warnings == nil {
		_gd.Warnings = make([]*DiffResult, 0)
	}
	_gd.Warnings = append(_gd.Warnings, &DiffResult{Revision: _ffe, Description: _eff})
}
func _agb(_fbb _c.PdfObject) ([]_c.PdfObject, error) {
	_dfe := make([]_c.PdfObject, 0)
	if _fbb != nil {
		_beg := _fbb
		if _egaa, _gab := _c.GetIndirect(_fbb); _gab {
			_beg = _egaa.PdfObject
		}
		if _gbb, _eee := _c.GetArray(_beg); _eee {
			_dfe = _gbb.Elements()
		} else {
			return nil, _e.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0061n\u006eo\u0074s\u0027\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
	}
	return _dfe, nil
}
func (_ecg *defaultDiffPolicy) compareAnnots(_dgd int, _cfa, _deg []_c.PdfObject) error {
	_dcc := make(map[int64]*_c.PdfObjectDictionary)
	for _, _gcf := range _cfa {
		_feb, _aae := _c.GetIndirect(_gcf)
		if !_aae {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_ccf, _aae := _c.GetDict(_feb.PdfObject)
		if !_aae {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_dcc[_feb.ObjectNumber] = _ccf
	}
	for _, _dff := range _deg {
		_cee, _feg := _c.GetIndirect(_dff)
		if !_feg {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_gge, _feg := _c.GetDict(_cee.PdfObject)
		if !_feg {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_cd, _ := _c.GetStringVal(_gge.Get("\u0054"))
		_gfba, _ := _c.GetNameVal(_gge.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if _, _bfe := _dcc[_cee.ObjectNumber]; !_bfe {
			switch _ecg._d {
			case NoRestrictions, FillFormsAndAnnots:
				_ecg._eg.addWarningWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gfba, _cd))
			default:
				_fcaf, _aaf := _c.GetDict(_cee.PdfObject)
				if !_aaf {
					return _e.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
				}
				_eca, _aaf := _c.GetNameVal(_fcaf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
				if !_aaf {
					return _e.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
				}
				if _eca == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					switch _ecg._d {
					case NoRestrictions, FillFormsAndAnnots, FillForms:
						_ecg._eg.addWarningWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gfba, _cd))
					default:
						_ecg._eg.addErrorWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gfba, _cd))
					}
				} else {
					_ecg._eg.addErrorWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _gfba, _cd))
				}
			}
		} else {
			delete(_dcc, _cee.ObjectNumber)
			if _fgbg, _gfg := _ecg._ef[_cee.ObjectNumber]; _gfg {
				switch _ecg._d {
				case NoRestrictions, FillFormsAndAnnots:
					_ecg._eg.addWarningWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gfba, _cd))
				default:
					_adg, _acb := _c.GetIndirect(_fgbg)
					if !_acb {
						return _e.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_gbe, _acb := _c.GetDict(_adg.PdfObject)
					if !_acb {
						return _e.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_ga, _acb := _c.GetNameVal(_gbe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
					if !_acb {
						return _e.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
					}
					if _ga == "\u0057\u0069\u0064\u0067\u0065\u0074" {
						switch _ecg._d {
						case NoRestrictions, FillFormsAndAnnots, FillForms:
							_ecg._eg.addWarningWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gfba, _cd))
						default:
							_ecg._eg.addErrorWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gfba, _cd))
						}
					} else {
						_ecg._eg.addErrorWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _gfba, _cd))
					}
				}
			}
		}
	}
	for _, _ddb := range _dcc {
		_bc, _ := _c.GetStringVal(_ddb.Get("\u0054"))
		_cfd, _ := _c.GetNameVal(_ddb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		switch _ecg._d {
		case NoRestrictions, FillFormsAndAnnots:
			_ecg._eg.addWarningWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _cfd, _bc))
		default:
			_ecg._eg.addErrorWithDescription(_dgd, _g.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _cfd, _bc))
		}
	}
	return nil
}
func (_df *defaultDiffPolicy) compareFields(_fe int, _ad, _fb []_c.PdfObject) error {
	_fgb := make(map[int64]*_c.PdfObjectDictionary)
	for _, _gg := range _ad {
		_bga, _af := _c.GetIndirect(_gg)
		if !_af {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_efaf, _af := _c.GetDict(_bga.PdfObject)
		if !_af {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_fgb[_bga.ObjectNumber] = _efaf
	}
	for _, _fbc := range _fb {
		_fbg, _cbd := _c.GetIndirect(_fbc)
		if !_cbd {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_dc, _cbd := _c.GetDict(_fbg.PdfObject)
		if !_cbd {
			return _e.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		T := _dc.Get("\u0054")
		if _, _abg := _df._ef[_fbg.ObjectNumber]; _abg {
			switch _df._d {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_df._eg.addWarningWithDescription(_fe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			default:
				_df._eg.addErrorWithDescription(_fe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			}
		}
		if _, _cfg := _fgb[_fbg.ObjectNumber]; !_cfg {
			switch _df._d {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_df._eg.addWarningWithDescription(_fe, _g.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _dc.Get("\u0054")))
			default:
				_df._eg.addErrorWithDescription(_fe, _g.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _dc.Get("\u0054")))
			}
		} else {
			delete(_fgb, _fbg.ObjectNumber)
			if _, _ae := _df._ef[_fbg.ObjectNumber]; _ae {
				switch _df._d {
				case NoRestrictions, FillForms, FillFormsAndAnnots:
					_df._eg.addWarningWithDescription(_fe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _dc.Get("\u0054")))
				default:
					_df._eg.addErrorWithDescription(_fe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _dc.Get("\u0054")))
				}
			}
		}
		if FT, _cc := _c.GetNameVal(_dc.Get("\u0046\u0054")); _cc {
			if FT == "\u0053\u0069\u0067" {
				if _da, _ec := _c.GetIndirect(_dc.Get("\u0056")); _ec {
					if _, _efe := _df._ef[_da.ObjectNumber]; _efe {
						switch _df._d {
						case NoRestrictions, FillForms, FillFormsAndAnnots:
							_df._eg.addWarningWithDescription(_fe, _g.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						default:
							_df._eg.addErrorWithDescription(_fe, _g.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						}
					}
				}
			}
		}
	}
	for _, _ecb := range _fgb {
		switch _df._d {
		case NoRestrictions:
			_df._eg.addWarningWithDescription(_fe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _ecb.Get("\u0054")))
		default:
			_df._eg.addErrorWithDescription(_fe, _g.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _ecb.Get("\u0054")))
		}
	}
	return nil
}
