package mdp

import (
	_a "errors"
	_ef "fmt"

	_ed "bitbucket.org/shenghui0779/gopdf/core"
)

func (_dfb *defaultDiffPolicy) compareRevisions(_ba *_ed.PdfParser, _de *_ed.PdfParser) (*DiffResults, error) {
	var _ffe error
	_dfb._g, _ffe = _de.GetUpdatedObjects(_ba)
	if _ffe != nil {
		return &DiffResults{}, _ffe
	}
	if len(_dfb._g) == 0 {
		return &DiffResults{}, nil
	}
	_bd := _de.GetRevisionNumber()
	_ge, _gg := _ed.GetIndirect(_ed.ResolveReference(_ba.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	_ffd, _db := _ed.GetIndirect(_ed.ResolveReference(_de.GetTrailer().Get("\u0052\u006f\u006f\u0074")))
	if !_gg || !_db {
		return &DiffResults{}, _a.New("\u0065\u0072\u0072o\u0072\u0020\u0077\u0068i\u006c\u0065\u0020\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0072\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_be, _gg := _ed.GetDict(_ed.ResolveReference(_ge.PdfObject))
	_ec, _db := _ed.GetDict(_ed.ResolveReference(_ffd.PdfObject))
	if !_gg || !_db {
		return &DiffResults{}, _a.New("\u0065\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020a\u0020\u0072\u006f\u006f\u0074\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
	}
	if _ggg, _gc := _ed.GetIndirect(_ec.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _gc {
		_agc, _ga := _ed.GetDict(_ggg)
		if !_ga {
			return &DiffResults{}, _a.New("\u0065\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_deg := make([]_ed.PdfObject, 0)
		if _aa, _bf := _ed.GetIndirect(_be.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")); _bf {
			if _dfe, _bb := _ed.GetDict(_aa); _bb {
				if _dc, _efc := _ed.GetArray(_dfe.Get("\u0046\u0069\u0065\u006c\u0064\u0073")); _efc {
					_deg = _dc.Elements()
				}
			}
		}
		_ggb, _ga := _ed.GetArray(_agc.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_ga {
			return &DiffResults{}, _a.New("\u0065\u0072r\u006f\u0072\u0020\u0077h\u0069\u006ce\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0066i\u0065\u006c\u0064\u0073")
		}
		if _dg := _dfb.compareFields(_bd, _deg, _ggb.Elements()); _dg != nil {
			return &DiffResults{}, _dg
		}
	}
	_fc, _degc := _ed.GetIndirect(_ec.Get("\u0050\u0061\u0067e\u0073"))
	if !_degc {
		return &DiffResults{}, _a.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_cf, _degc := _ed.GetIndirect(_be.Get("\u0050\u0061\u0067e\u0073"))
	if !_degc {
		return &DiffResults{}, _a.New("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	if _fg := _dfb.comparePages(_bd, _cf, _fc); _fg != nil {
		return &DiffResults{}, _fg
	}
	return _dfb._d, nil
}
func (_gba *DiffResults) addErrorWithDescription(_cba int, _egd string) {
	if _gba.Errors == nil {
		_gba.Errors = make([]*DiffResult, 0)
	}
	_gba.Errors = append(_gba.Errors, &DiffResult{Revision: _cba, Description: _egd})
}

// DocMDPPermission is values for set up access permissions for DocMDP.
// (Section 12.8.2.2, Table 254 - Entries in a signature dictionary p. 471 in PDF32000_2008).
type DocMDPPermission int64
type defaultDiffPolicy struct {
	_g map[int64]_ed.PdfObject
	_d *DiffResults
	_f DocMDPPermission
}

func (_gde *DiffResults) addWarningWithDescription(_ad int, _fda string) {
	if _gde.Warnings == nil {
		_gde.Warnings = make([]*DiffResult, 0)
	}
	_gde.Warnings = append(_gde.Warnings, &DiffResult{Revision: _ad, Description: _fda})
}
func (_dbc *defaultDiffPolicy) comparePages(_ffde int, _ea, _ggf *_ed.PdfIndirectObject) error {
	if _, _bdd := _dbc._g[_ggf.ObjectNumber]; _bdd {
		_dbc._d.addErrorWithDescription(_ffde, "\u0050a\u0067e\u0073\u0020\u0077\u0065\u0072e\u0020\u0063h\u0061\u006e\u0067\u0065\u0064")
	}
	_fad, _gd := _ed.GetDict(_ggf.PdfObject)
	_dfg, _eff := _ed.GetDict(_ea.PdfObject)
	if !_gd || !_eff {
		return _a.New("\u0075n\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061g\u0065\u0073\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_dfac, _gd := _ed.GetArray(_fad.Get("\u004b\u0069\u0064\u0073"))
	_ded, _eff := _ed.GetArray(_dfg.Get("\u004b\u0069\u0064\u0073"))
	if !_gd || !_eff {
		return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061\u0067\u0065s\u0027 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
	}
	_dge := _dfac.Len()
	if _dge > _ded.Len() {
		_dge = _ded.Len()
	}
	for _bcc := 0; _bcc < _dge; _bcc++ {
		_ccf, _bga := _ed.GetIndirect(_ed.ResolveReference(_ded.Get(_bcc)))
		_cff, _ce := _ed.GetIndirect(_ed.ResolveReference(_dfac.Get(_bcc)))
		if !_bga || !_ce {
			return _a.New("\u0075\u006e\u0065\u0078pe\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065c\u0074")
		}
		if _ccf.ObjectNumber != _cff.ObjectNumber {
			_dbc._d.addErrorWithDescription(_ffde, _ef.Sprintf("p\u0061\u0067\u0065\u0020#%\u0064 \u0077\u0061\u0073\u0020\u0072e\u0070\u006c\u0061\u0063\u0065\u0064", _bcc))
		}
		_eae, _bga := _ed.GetDict(_cff)
		_faa, _ce := _ed.GetDict(_ccf)
		if !_bga || !_ce {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065'\u0073 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079")
		}
		_bdb, _bfa := _gcb(_eae.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _bfa != nil {
			return _bfa
		}
		_gb, _bfa := _gcb(_faa.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _bfa != nil {
			return _bfa
		}
		if _egf := _dbc.compareAnnots(_ffde, _gb, _bdb); _egf != nil {
			return _egf
		}
	}
	for _ede := _dge + 1; _ede <= _dfac.Len(); _ede++ {
		_dbc._d.addErrorWithDescription(_ffde, _ef.Sprintf("\u0070a\u0067e\u0020\u0023\u0025\u0064\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _ede))
	}
	for _eb := _dge + 1; _eb <= _ded.Len(); _eb++ {
		_dbc._d.addErrorWithDescription(_ffde, _ef.Sprintf("p\u0061g\u0065\u0020\u0023\u0025\u0064\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _eb))
	}
	return nil
}

// String returns the state of the warning.
func (_gfa *DiffResult) String() string {
	return _ef.Sprintf("\u0025\u0073\u0020\u0069n \u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0023\u0025\u0064", _gfa.Description, _gfa.Revision)
}

// DiffPolicy interface for comparing two revisions of the Pdf document.
type DiffPolicy interface {

	// ReviewFile should check the revisions of the old and new parsers
	// and evaluate the differences between the revisions.
	// Each implementation of this interface must decide
	// how to handle cases where there are multiple revisions between the old and new revisions.
	ReviewFile(_feac *_ed.PdfParser, _cbf *_ed.PdfParser, _gfg *MDPParameters) (*DiffResults, error)
}

// DiffResults describes the results of the DiffPolicy.
type DiffResults struct {
	Warnings []*DiffResult
	Errors   []*DiffResult
}

func (_ggd *defaultDiffPolicy) compareAnnots(_gcg int, _fcd, _cg []_ed.PdfObject) error {
	_fe := make(map[int64]*_ed.PdfObjectDictionary)
	for _, _agcb := range _fcd {
		_baad, _fea := _ed.GetIndirect(_agcb)
		if !_fea {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_dgc, _fea := _ed.GetDict(_baad.PdfObject)
		if !_fea {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_fe[_baad.ObjectNumber] = _dgc
	}
	for _, _aaa := range _cg {
		_gbb, _dbe := _ed.GetIndirect(_aaa)
		if !_dbe {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_cd, _dbe := _ed.GetDict(_gbb.PdfObject)
		if !_dbe {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_bfg, _ := _ed.GetStringVal(_cd.Get("\u0054"))
		_bff, _ := _ed.GetNameVal(_cd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if _, _abg := _fe[_gbb.ObjectNumber]; !_abg {
			switch _ggd._f {
			case NoRestrictions, FillFormsAndAnnots:
				_ggd._d.addWarningWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bff, _bfg))
			default:
				_dbd, _aef := _ed.GetDict(_gbb.PdfObject)
				if !_aef {
					return _a.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
				}
				_aae, _aef := _ed.GetNameVal(_dbd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
				if !_aef {
					return _a.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
				}
				if _aae == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					switch _ggd._f {
					case NoRestrictions, FillFormsAndAnnots, FillForms:
						_ggd._d.addWarningWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bff, _bfg))
					default:
						_ggd._d.addErrorWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bff, _bfg))
					}
				} else {
					_ggd._d.addErrorWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064", _bff, _bfg))
				}
			}
		} else {
			delete(_fe, _gbb.ObjectNumber)
			if _gdf, _bgb := _ggd._g[_gbb.ObjectNumber]; _bgb {
				switch _ggd._f {
				case NoRestrictions, FillFormsAndAnnots:
					_ggd._d.addWarningWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bff, _bfg))
				default:
					_age, _bbb := _ed.GetIndirect(_gdf)
					if !_bbb {
						return _a.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_dga, _bbb := _ed.GetDict(_age.PdfObject)
					if !_bbb {
						return _a.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e")
					}
					_dgcg, _bbb := _ed.GetNameVal(_dga.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
					if !_bbb {
						return _a.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065")
					}
					if _dgcg == "\u0057\u0069\u0064\u0067\u0065\u0074" {
						switch _ggd._f {
						case NoRestrictions, FillFormsAndAnnots, FillForms:
							_ggd._d.addWarningWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bff, _bfg))
						default:
							_ggd._d.addErrorWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bff, _bfg))
						}
					} else {
						_ggd._d.addErrorWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064", _bff, _bfg))
					}
				}
			}
		}
	}
	for _, _eag := range _fe {
		_fbc, _ := _ed.GetStringVal(_eag.Get("\u0054"))
		_aed, _ := _ed.GetNameVal(_eag.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		switch _ggd._f {
		case NoRestrictions, FillFormsAndAnnots:
			_ggd._d.addWarningWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _aed, _fbc))
		default:
			_ggd._d.addErrorWithDescription(_gcg, _ef.Sprintf("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064", _aed, _fbc))
		}
	}
	return nil
}

// DiffResult describes the warning or the error for the DiffPolicy results.
type DiffResult struct {
	Revision    int
	Description string
}

// MDPParameters describes parameters for the MDP checks (now only DocMDP).
type MDPParameters struct{ DocMDPLevel DocMDPPermission }

// IsPermitted returns true if changes permitted.
func (_aad *DiffResults) IsPermitted() bool { return len(_aad.Errors) == 0 }

// ReviewFile implementation of DiffPolicy interface
// The default policy only checks the next types of objects:
// Page, Pages (container for page objects), Annot, Annots (container for annotation objects), Field.
// It checks adding, removing and modifying objects of these types.
func (_c *defaultDiffPolicy) ReviewFile(oldParser *_ed.PdfParser, newParser *_ed.PdfParser, params *MDPParameters) (*DiffResults, error) {
	if oldParser.GetRevisionNumber() > newParser.GetRevisionNumber() {
		return nil, _a.New("\u006f\u006c\u0064\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u006e\u0065\u0077\u0020r\u0065\u0076i\u0073\u0069\u006f\u006e")
	}
	if oldParser.GetRevisionNumber() == newParser.GetRevisionNumber() {
		if oldParser != newParser {
			return nil, _a.New("\u0073\u0061m\u0065\u0020\u0072\u0065v\u0069\u0073i\u006f\u006e\u0073\u002c\u0020\u0062\u0075\u0074 \u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0070\u0061r\u0073\u0065\u0072\u0073")
		}
		return &DiffResults{}, nil
	}
	if params == nil {
		_c._f = NoRestrictions
	} else {
		_c._f = params.DocMDPLevel
	}
	_ff := &DiffResults{}
	for _df := oldParser.GetRevisionNumber() + 1; _df <= newParser.GetRevisionNumber(); _df++ {
		_b, _ag := newParser.GetRevision(_df - 1)
		if _ag != nil {
			return nil, _ag
		}
		_ae, _ag := newParser.GetRevision(_df)
		if _ag != nil {
			return nil, _ag
		}
		_da, _ag := _c.compareRevisions(_b, _ae)
		if _ag != nil {
			return nil, _ag
		}
		_ff.Warnings = append(_ff.Warnings, _da.Warnings...)
		_ff.Errors = append(_ff.Errors, _da.Errors...)
	}
	return _ff, nil
}
func _gcb(_bbc _ed.PdfObject) ([]_ed.PdfObject, error) {
	_cfcc := make([]_ed.PdfObject, 0)
	if _bbc != nil {
		_aaed := _bbc
		if _fd, _dbg := _ed.GetIndirect(_bbc); _dbg {
			_aaed = _fd.PdfObject
		}
		if _dee, _cae := _ed.GetArray(_aaed); _cae {
			_cfcc = _dee.Elements()
		} else {
			return nil, _a.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0061n\u006eo\u0074s\u0027\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
	}
	return _cfcc, nil
}
func (_ffc *DiffResults) addWarning(_bfc *DiffResult) {
	if _ffc.Warnings == nil {
		_ffc.Warnings = make([]*DiffResult, 0)
	}
	_ffc.Warnings = append(_ffc.Warnings, _bfc)
}

const (
	NoRestrictions     DocMDPPermission = 0
	NoChanges          DocMDPPermission = 1
	FillForms          DocMDPPermission = 2
	FillFormsAndAnnots DocMDPPermission = 3
)

func (_fb *defaultDiffPolicy) compareFields(_beb int, _bc, _dfa []_ed.PdfObject) error {
	_gf := make(map[int64]*_ed.PdfObjectDictionary)
	for _, _ggbd := range _bc {
		_fgb, _bed := _ed.GetIndirect(_ggbd)
		if !_bed {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_dd, _bed := _ed.GetDict(_fgb.PdfObject)
		if !_bed {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_gf[_fgb.ObjectNumber] = _dd
	}
	for _, _bfb := range _dfa {
		_efe, _eg := _ed.GetIndirect(_bfb)
		if !_eg {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		_cfc, _eg := _ed.GetDict(_efe.PdfObject)
		if !_eg {
			return _a.New("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065")
		}
		T := _cfc.Get("\u0054")
		if _, _dfc := _fb._g[_efe.ObjectNumber]; _dfc {
			switch _fb._f {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_fb._d.addWarningWithDescription(_beb, _ef.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			default:
				_fb._d.addErrorWithDescription(_beb, _ef.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", T))
			}
		}
		if _, _edb := _gf[_efe.ObjectNumber]; !_edb {
			switch _fb._f {
			case NoRestrictions, FillForms, FillFormsAndAnnots:
				_fb._d.addWarningWithDescription(_beb, _ef.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _cfc.Get("\u0054")))
			default:
				_fb._d.addErrorWithDescription(_beb, _ef.Sprintf("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064", _cfc.Get("\u0054")))
			}
		} else {
			delete(_gf, _efe.ObjectNumber)
			if _, _bg := _fb._g[_efe.ObjectNumber]; _bg {
				switch _fb._f {
				case NoRestrictions, FillForms, FillFormsAndAnnots:
					_fb._d.addWarningWithDescription(_beb, _ef.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _cfc.Get("\u0054")))
				default:
					_fb._d.addErrorWithDescription(_beb, _ef.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064", _cfc.Get("\u0054")))
				}
			}
		}
		if FT, _cc := _ed.GetNameVal(_cfc.Get("\u0046\u0054")); _cc {
			if FT == "\u0053\u0069\u0067" {
				if _ab, _bae := _ed.GetIndirect(_cfc.Get("\u0056")); _bae {
					if _, _cb := _fb._g[_ab.ObjectNumber]; _cb {
						switch _fb._f {
						case NoRestrictions, FillForms, FillFormsAndAnnots:
							_fb._d.addWarningWithDescription(_beb, _ef.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						default:
							_fb._d.addErrorWithDescription(_beb, _ef.Sprintf("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064", T))
						}
					}
				}
			}
		}
	}
	for _, _fa := range _gf {
		switch _fb._f {
		case NoRestrictions:
			_fb._d.addWarningWithDescription(_beb, _ef.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _fa.Get("\u0054")))
		default:
			_fb._d.addErrorWithDescription(_beb, _ef.Sprintf("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064", _fa.Get("\u0054")))
		}
	}
	return nil
}
func (_gfd *DiffResults) addError(_edf *DiffResult) {
	if _gfd.Errors == nil {
		_gfd.Errors = make([]*DiffResult, 0)
	}
	_gfd.Errors = append(_gfd.Errors, _edf)
}
func NewDefaultDiffPolicy() DiffPolicy { return &defaultDiffPolicy{_g: nil, _d: &DiffResults{}, _f: 0} }
