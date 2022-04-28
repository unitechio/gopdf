package pdfa

import (
	_f "errors"
	_c "fmt"
	_d "image/color"
	_b "math"
	_be "sort"
	_a "strings"
	_bee "time"

	_ge "bitbucket.org/shenghui0779/gopdf/common"
	_ee "bitbucket.org/shenghui0779/gopdf/contentstream"
	_cb "bitbucket.org/shenghui0779/gopdf/core"
	_ad "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_eg "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_ac "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_g "bitbucket.org/shenghui0779/gopdf/model"
	_gf "bitbucket.org/shenghui0779/gopdf/model/internal/colorprofile"
	_gc "bitbucket.org/shenghui0779/gopdf/model/internal/docutil"
	_ae "bitbucket.org/shenghui0779/gopdf/model/internal/fonts"
	_gga "bitbucket.org/shenghui0779/gopdf/model/xmputil"
	_cc "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_ea "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaid"
	_fg "github.com/adrg/sysfont"
	_gg "github.com/trimmer-io/go-xmp/models/dc"
	_df "github.com/trimmer-io/go-xmp/models/pdf"
	_fed "github.com/trimmer-io/go-xmp/models/xmp_base"
	_dea "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_de "github.com/trimmer-io/go-xmp/models/xmp_rights"
	_fe "github.com/trimmer-io/go-xmp/xmp"
)

func _ggd() standardType { return standardType{_da: 2, _dag: "\u0055"} }
func _gfgd(_egfc *_gc.Document) error {
	_gecb := func(_aeb *_cb.PdfObjectDictionary) error {
		if _aeb.Get("\u0054\u0052") != nil {
			_ge.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_aeb.Remove("\u0054\u0052")
		}
		_dgg := _aeb.Get("\u0054\u0052\u0032")
		if _dgg != nil {
			_aced := _dgg.String()
			if _aced != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_ge.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_aeb.Set("\u0054\u0052\u0032", _cb.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		if _aeb.Get("\u0048\u0054\u0050") != nil {
			_ge.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0073\u0020\u0048\u0054P\u0020\u006b\u0065\u0079")
			_aeb.Remove("\u0048\u0054\u0050")
		}
		_eege := _aeb.Get("\u0042\u004d")
		if _eege != nil {
			_acba, _cbea := _cb.GetName(_eege)
			if !_cbea {
				_ge.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_acba = _cb.MakeName("")
			}
			_abd := _acba.String()
			switch _abd {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
			default:
				_aeb.Set("\u0042\u004d", _cb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		return nil
	}
	_aff, _dgab := _egfc.GetPages()
	if !_dgab {
		return nil
	}
	for _, _eccc := range _aff {
		_ebdc, _cgfad := _eccc.GetResources()
		if !_cgfad {
			continue
		}
		_gebf, _agbb := _cb.GetDict(_ebdc.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_agbb {
			return nil
		}
		_dafa := _gebf.Keys()
		for _, _fca := range _dafa {
			_geba, _bddc := _cb.GetDict(_gebf.Get(_fca))
			if !_bddc {
				continue
			}
			_dbbf := _gecb(_geba)
			if _dbbf != nil {
				continue
			}
		}
	}
	for _, _eggd := range _aff {
		_eceb, _ceefb := _eggd.GetContents()
		if !_ceefb {
			return nil
		}
		for _, _ggad := range _eceb {
			_gbgc, _eecge := _ggad.GetData()
			if _eecge != nil {
				continue
			}
			_gfccd := _ee.NewContentStreamParser(string(_gbgc))
			_efbe, _eecge := _gfccd.Parse()
			if _eecge != nil {
				continue
			}
			for _, _agfa := range *_efbe {
				if len(_agfa.Params) == 0 {
					continue
				}
				_, _defe := _cb.GetName(_agfa.Params[0])
				if !_defe {
					continue
				}
				_bac, _ffafg := _eggd.GetResourcesXObject()
				if !_ffafg {
					continue
				}
				for _, _ddeeg := range _bac.Keys() {
					_cegb, _dbde := _cb.GetStream(_bac.Get(_ddeeg))
					if !_dbde {
						continue
					}
					_dbdb, _dbde := _cb.GetDict(_cegb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_dbde {
						continue
					}
					_fcee, _dbde := _cb.GetDict(_dbdb.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_dbde {
						continue
					}
					for _, _dacf := range _fcee.Keys() {
						_gefd, _bbed := _cb.GetDict(_fcee.Get(_dacf))
						if !_bbed {
							continue
						}
						_ebde := _gecb(_gefd)
						if _ebde != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}

// Profile is the model.StandardImplementer enhanced by the information about the profile conformance level.
type Profile interface {
	_g.StandardImplementer
	Conformance() string
	Part() int
}

func _bgaa(_addd *_cb.PdfObjectDictionary, _acbb map[*_cb.PdfObjectStream][]byte, _edfd map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_bcfe = "\u0036.\u0033\u002e\u0033\u002d\u0034"
		_becc = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _eeegd string
	if _bbdad, _efgc := _cb.GetName(_addd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _efgc {
		_eeegd = _bbdad.String()
	}
	if _eeegd != "\u0054\u0079\u0070e\u0030" {
		return _dfa
	}
	_abdc := _addd.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _eede := _cb.GetName(_abdc); _eede {
		return _dfa
	}
	_cbeec, _gfbgf := _cb.GetStream(_abdc)
	if !_gfbgf {
		return _dd(_bcfe, _becc)
	}
	_gagb, _bbcga := _cdfbg(_cbeec, _acbb, _edfd)
	if _bbcga != nil {
		return _dd(_bcfe, _becc)
	}
	_agaed, _dbec := _cb.GetIntVal(_cbeec.Get("\u0057\u004d\u006fd\u0065"))
	_cdcce, _fbdd := _gagb.WMode()
	if _dbec && _fbdd {
		if _cdcce != _agaed {
			return _dd(_bcfe, _becc)
		}
	}
	if (_dbec && !_fbdd) || (!_dbec && _fbdd) {
		return _dd(_bcfe, _becc)
	}
	return _dfa
}

// NewProfile2U creates a new Profile2U with the given options.
func NewProfile2U(options *Profile2Options) *Profile2U {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_eecgb(options)
	return &Profile2U{profile2{_afef: *options, _gaeg: _ggd()}}
}
func _gcae(_gffe *_gc.Document) error {
	_dece, _dfda := _gffe.GetPages()
	if !_dfda {
		return nil
	}
	for _, _gbcf := range _dece {
		_daba, _fceg := _cb.GetArray(_gbcf.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_fceg {
			continue
		}
		for _, _cbda := range _daba.Elements() {
			_cbda = _cb.ResolveReference(_cbda)
			if _, _bea := _cbda.(*_cb.PdfObjectNull); _bea {
				continue
			}
			_dagb, _dgaa := _cb.GetDict(_cbda)
			if !_dgaa {
				continue
			}
			_dfga, _ := _cb.GetIntVal(_dagb.Get("\u0046"))
			_dfga &= ^(1 << 0)
			_dfga &= ^(1 << 1)
			_dfga &= ^(1 << 5)
			_dfga |= 1 << 2
			_dagb.Set("\u0046", _cb.MakeInteger(int64(_dfga)))
			_cbcc := false
			if _ffbe := _dagb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _ffbe != nil {
				_bfaf, _ddefb := _cb.GetName(_ffbe)
				if _ddefb && _bfaf.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_cbcc = true
					if _dagb.Get("\u0041\u0041") != nil {
						_dagb.Remove("\u0041\u0041")
					}
				}
			}
			if _dagb.Get("\u0043") != nil || _dagb.Get("\u0049\u0043") != nil {
				_gage, _fbcf := _bbgg(_gffe)
				if !_fbcf {
					_dagb.Remove("\u0043")
					_dagb.Remove("\u0049\u0043")
				} else {
					_bege, _begc := _cb.GetIntVal(_gage.Get("\u004e"))
					if !_begc || _bege != 3 {
						_dagb.Remove("\u0043")
						_dagb.Remove("\u0049\u0043")
					}
				}
			}
			_gcdf, _dgaa := _cb.GetDict(_dagb.Get("\u0041\u0050"))
			if _dgaa {
				_fggc := _gcdf.Get("\u004e")
				if _fggc == nil {
					continue
				}
				if len(_gcdf.Keys()) > 1 {
					_gcdf.Clear()
					_gcdf.Set("\u004e", _fggc)
				}
				if _cbcc {
					_agbc, _aaed := _cb.GetName(_dagb.Get("\u0046\u0054"))
					if _aaed && *_agbc == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}
func (_db *documentImages) hasOnlyDeviceRGB() bool { return _db._egc && !_db._eb && !_db._cf }
func _aaac(_bbecad *_g.CompliancePdfReader, _acbfb standardType, _dadca bool) (_fdcd []ViolatedRule) {
	_edfc, _bbac := _addf(_bbecad)
	if !_bbac {
		return []ViolatedRule{_dd("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_bgae := _edfc.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _bgae == nil {
		return []ViolatedRule{_dd("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u006e\u006f\u0020\u0027\u004d\u0065\u0074\u0061d\u0061\u0074\u0061' \u006b\u0065\u0079\u0020\u0066\u006fu\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u002e"), _dd("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	_bcedg, _bbac := _cb.GetStream(_bgae)
	if !_bbac {
		return []ViolatedRule{_dd("\u0036.\u0037\u002e\u0032\u002d\u0032", "\u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0027\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"), _dd("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	if _bcedg.Get("\u0046\u0069\u006c\u0074\u0065\u0072") != nil {
		_fdcd = append(_fdcd, _dd("\u0036.\u0037\u002e\u0032\u002d\u0032", "M\u0065\u0074a\u0064\u0061\u0074\u0061\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u006b\u0065y\u002e"))
	}
	_fcafc, _gecfe := _gga.LoadDocument(_bcedg.Stream)
	if _gecfe != nil {
		return []ViolatedRule{_dd("\u0036.\u0037\u002e\u0039\u002d\u0031", "The\u0020\u006d\u0065\u0074a\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0061\u006e\u0064\u0020\u0077\u0065\u006c\u006c\u0020\u0066\u006f\u0072\u006de\u0064\u0020\u0050\u0044\u0046\u0041\u0045\u0078\u0074e\u006e\u0073\u0069\u006f\u006e\u0020\u0053\u0063\u0068\u0065\u006da\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073\u002e")}
	}
	_dcfc := _fcafc.GetGoXmpDocument()
	var _gaeea []*_fe.Namespace
	for _, _cbdgb := range _dcfc.Namespaces() {
		switch _cbdgb.Name {
		case _gg.NsDc.Name, _df.NsPDF.Name, _fed.NsXmp.Name, _de.NsXmpRights.Name, _ea.Namespace.Name, _cc.Namespace.Name, _dea.NsXmpMM.Name, _cc.FieldNS.Name, _cc.SchemaNS.Name, _cc.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_gaeea = append(_gaeea, _cbdgb)
	}
	_efad := true
	_gabef, _gecfe := _fcafc.GetPdfaExtensionSchemas()
	if _gecfe == nil {
		for _, _gdee := range _gaeea {
			var _agdef bool
			for _cceg := range _gabef {
				if _gdee.URI == _gabef[_cceg].NamespaceURI {
					_agdef = true
					break
				}
			}
			if !_agdef {
				_efad = false
				break
			}
		}
	} else {
		_efad = false
	}
	if !_efad {
		_fdcd = append(_fdcd, _dd("\u0036.\u0037\u002e\u0039\u002d\u0032", "\u0050\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073 \u0073\u0070\u0065\u0063\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0058M\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0073\u0068\u0061\u006cl\u0020\u0075\u0073\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073 \u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002c\u0020\u006f\u0072\u0020\u0065\u0078\u0074\u0065ns\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u0074\u0068\u0061\u0074 \u0063\u006f\u006d\u0070\u006c\u0079\u0020\u0077\u0069\u0074h\u0020\u0058\u004d\u0050\u0020\u0053\u0070e\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002e"))
	}
	_dbda, _gecfe := _bbecad.GetPdfInfo()
	if _gecfe == nil {
		if !_afaf(_dbda, _fcafc) {
			_fdcd = append(_fdcd, _dd("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
		}
	} else if _, _ebgb := _fcafc.GetMediaManagement(); _ebgb {
		_fdcd = append(_fdcd, _dd("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
	}
	_caaea, _bbac := _fcafc.GetPdfAID()
	if !_bbac {
		_fdcd = append(_fdcd, _dd("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _caaea.Part != _acbfb._da {
			_fdcd = append(_fdcd, _dd("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _acbfb._dag == "\u0041" && _caaea.Conformance != "\u0041" {
			_fdcd = append(_fdcd, _dd("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		} else if _acbfb._dag == "\u0042" && (_caaea.Conformance != "\u0041" && _caaea.Conformance != "\u0042") {
			_fdcd = append(_fdcd, _dd("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		}
	}
	return _fdcd
}

// Part gets the PDF/A version level.
func (_cdc *profile1) Part() int { return _cdc._efc._da }
func _cabe(_bbaa *_g.CompliancePdfReader, _afgd standardType) (_fbbe []ViolatedRule) {
	var _deef, _dceb, _geg, _facb, _adcad, _cfcdg, _bggbc, _bbfd, _bcca, _ddga, _acbae bool
	_fcc := func() bool {
		return _deef && _dceb && _geg && _facb && _adcad && _cfcdg && _bggbc && _bbfd && _bcca && _ddga && _acbae
	}
	_aagc := map[*_cb.PdfObjectStream]*_ad.CMap{}
	_fagd := map[*_cb.PdfObjectStream][]byte{}
	_acda := map[_cb.PdfObject]*_g.PdfFont{}
	for _, _dcccg := range _bbaa.GetObjectNums() {
		_dded, _edfea := _bbaa.GetIndirectObjectByNumber(_dcccg)
		if _edfea != nil {
			continue
		}
		_eeae, _bced := _cb.GetDict(_dded)
		if !_bced {
			continue
		}
		_caba, _bced := _cb.GetName(_eeae.Get("\u0054\u0079\u0070\u0065"))
		if !_bced {
			continue
		}
		if *_caba != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_deea, _edfea := _g.NewPdfFontFromPdfObject(_eeae)
		if _edfea != nil {
			_ge.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _edfea)
			continue
		}
		_acda[_eeae] = _deea
	}
	for _, _afd := range _bbaa.PageList {
		_edaf, _edga := _afd.GetContentStreams()
		if _edga != nil {
			_ge.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _cdce := range _edaf {
			_fcbg := _ee.NewContentStreamParser(_cdce)
			_dfebb, _cgg := _fcbg.Parse()
			if _cgg != nil {
				_ge.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _cgg)
				continue
			}
			var _gccg bool
			for _, _edeaf := range *_dfebb {
				if _edeaf.Operand != "\u0054\u0072" {
					continue
				}
				if len(_edeaf.Params) != 1 {
					_ge.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_edeaf.Params))
					continue
				}
				_cbec, _efda := _cb.GetIntVal(_edeaf.Params[0])
				if !_efda {
					_ge.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _cbec == 3 {
					_gccg = true
					break
				}
			}
			for _, _acgb := range *_dfebb {
				if _acgb.Operand != "\u0054\u0066" {
					continue
				}
				if len(_acgb.Params) != 2 {
					_ge.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_acgb.Params))
					continue
				}
				_bcfg, _egfa := _cb.GetName(_acgb.Params[0])
				if !_egfa {
					_ge.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _acgb)
					continue
				}
				_bdgdb, _aaca := _afd.Resources.GetFontByName(*_bcfg)
				if !_aaca {
					_ge.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_cbbf, _egfa := _cb.GetDict(_bdgdb)
				if !_egfa {
					_ge.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_ffge, _egfa := _acda[_cbbf]
				if !_egfa {
					var _eeeg error
					_ffge, _eeeg = _g.NewPdfFontFromPdfObject(_cbbf)
					if _eeeg != nil {
						_ge.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _eeeg)
						continue
					}
					_acda[_cbbf] = _ffge
				}
				if !_deef {
					_ffbg := _bffa(_cbbf, _fagd, _aagc)
					if _ffbg != _dfa {
						_fbbe = append(_fbbe, _ffbg)
						_deef = true
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_dceb {
					_fcdg := _dbfb(_cbbf)
					if _fcdg != _dfa {
						_fbbe = append(_fbbe, _fcdg)
						_dceb = true
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_geg {
					_ceaeg := _debc(_cbbf, _fagd, _aagc)
					if _ceaeg != _dfa {
						_fbbe = append(_fbbe, _ceaeg)
						_geg = true
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_facb {
					_bedda := _bgaa(_cbbf, _fagd, _aagc)
					if _bedda != _dfa {
						_fbbe = append(_fbbe, _bedda)
						_facb = true
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_adcad {
					_aage := _bggfg(_ffge, _cbbf, _gccg)
					if _aage != _dfa {
						_adcad = true
						_fbbe = append(_fbbe, _aage)
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_cfcdg {
					_abgfef := _ffbbb(_ffge, _cbbf)
					if _abgfef != _dfa {
						_cfcdg = true
						_fbbe = append(_fbbe, _abgfef)
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_bggbc {
					_faba := _eaff(_ffge, _cbbf)
					if _faba != _dfa {
						_bggbc = true
						_fbbe = append(_fbbe, _faba)
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_bbfd {
					_bfbea := _dbaa(_ffge, _cbbf)
					if _bfbea != _dfa {
						_bbfd = true
						_fbbe = append(_fbbe, _bfbea)
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_bcca {
					_bcb := _babg(_ffge, _cbbf)
					if _bcb != _dfa {
						_bcca = true
						_fbbe = append(_fbbe, _bcb)
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_ddga {
					_agae := _ebcb(_ffge, _cbbf)
					if _agae != _dfa {
						_ddga = true
						_fbbe = append(_fbbe, _agae)
						if _fcc() {
							return _fbbe
						}
					}
				}
				if !_acbae && _afgd._dag == "\u0041" {
					_afdb := _efbd(_cbbf, _fagd, _aagc)
					if _afdb != _dfa {
						_acbae = true
						_fbbe = append(_fbbe, _afdb)
						if _fcc() {
							return _fbbe
						}
					}
				}
			}
		}
	}
	return _fbbe
}
func (_fb standardType) String() string {
	return _c.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _fb._da, _fb._dag)
}
func _egfcf(_cebb *_g.CompliancePdfReader) ViolatedRule {
	_adcdc := _cebb.ParserMetadata()
	if _adcdc.HasInvalidSeparationAfterXRef() {
		return _dd("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _dfa
}
func _dgda(_faae standardType, _dge *_gc.OutputIntents) error {
	_decd, _gcb := _gf.NewISOCoatedV2Gray1CBasOutputIntent(_faae.outputIntentSubtype())
	if _gcb != nil {
		return _gcb
	}
	if _gcb = _dge.Add(_decd.ToPdfObject()); _gcb != nil {
		return _gcb
	}
	return nil
}
func _gdde(_afbg *_g.CompliancePdfReader) (*_g.PdfOutputIntent, bool) {
	_fbddd, _dabab := _ebga(_afbg)
	if !_dabab {
		return nil, false
	}
	_egedc, _bagd := _g.NewPdfOutputIntentFromPdfObject(_fbddd)
	if _bagd != nil {
		return nil, false
	}
	return _egedc, true
}

// Profile2Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile2Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _bee.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func _deec(_dbfba *_g.CompliancePdfReader) (_afbef []ViolatedRule) {
	_edcg, _febegf := _addf(_dbfba)
	if !_febegf {
		return _afbef
	}
	_agfb, _febegf := _cb.GetDict(_edcg.Get("\u004e\u0061\u006de\u0073"))
	if !_febegf {
		return _afbef
	}
	if _agfb.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_afbef = append(_afbef, _dd("\u0036\u002e\u0031\u0030\u002d\u0031", "T\u0068\u0065\u0072e\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u006e\u006f\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0050\u0072\u0065s\u0065\u006e\u0074a\u0074\u0069\u006f\u006e\u0073\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002e"))
	}
	return _afbef
}
func _debc(_dfca *_cb.PdfObjectDictionary, _gea map[*_cb.PdfObjectStream][]byte, _adge map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_facbg = "\u0036.\u0033\u002e\u0033\u002d\u0033"
		_eafeb = "\u0041\u006cl \u0043\u004d\u0061\u0070\u0073\u0020\u0075\u0073e\u0064 \u0077i\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072m\u0069n\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020a\u006e\u0064\u0020\u0049\u0064\u0065\u006et\u0069\u0074\u0079-\u0056\u002c\u0020\u0073\u0068a\u006c\u006c \u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064es\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u00205\u002e\u0036\u002e\u0034\u002e"
	)
	var _aceg string
	if _bdc, _aaaf := _cb.GetName(_dfca.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _aaaf {
		_aceg = _bdc.String()
	}
	if _aceg != "\u0054\u0079\u0070e\u0030" {
		return _dfa
	}
	_gccd := _dfca.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _cfeg, _ffeaf := _cb.GetName(_gccd); _ffeaf {
		switch _cfeg.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _dfa
		default:
			return _dd(_facbg, _eafeb)
		}
	}
	_dfege, _bbgfb := _cb.GetStream(_gccd)
	if !_bbgfb {
		return _dd(_facbg, _eafeb)
	}
	_, _fbde := _cdfbg(_dfege, _gea, _adge)
	if _fbde != nil {
		return _dd(_facbg, _eafeb)
	}
	return _dfa
}

type documentColorspaceOptimizeFunc func(_eac *_gc.Document, _bcc []*_gc.Image) error

func _adab(_gged *_g.CompliancePdfReader) (_cgfc []ViolatedRule) {
	_dfcg, _fcaa := _addf(_gged)
	if !_fcaa {
		return _cgfc
	}
	_cdea := _dd("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e")
	_bagf, _fcaa := _cb.GetArray(_dfcg.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_fcaa {
		_cgfc = append(_cgfc, _cdea)
		return _cgfc
	}
	_eddc := _dd("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e")
	if _bagf.Len() > 1 {
		_egaba := map[*_cb.PdfObjectDictionary]struct{}{}
		for _bbda := 0; _bbda < _bagf.Len(); _bbda++ {
			_fbcb, _cdca := _cb.GetDict(_bagf.Get(_bbda))
			if !_cdca {
				_cgfc = append(_cgfc, _cdea)
				return _cgfc
			}
			if _bbda == 0 {
				_egaba[_fbcb] = struct{}{}
				continue
			}
			if _, _dcff := _egaba[_fbcb]; !_dcff {
				_cgfc = append(_cgfc, _eddc)
				break
			}
		}
	} else if _bagf.Len() == 0 {
		_cgfc = append(_cgfc, _cdea)
		return _cgfc
	}
	_adga, _fcaa := _cb.GetDict(_bagf.Get(0))
	if !_fcaa {
		_cgfc = append(_cgfc, _cdea)
		return _cgfc
	}
	if _dccc, _cdbc := _cb.GetName(_adga.Get("\u0053")); !_cdbc || (*_dccc) != "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411" {
		_cgfc = append(_cgfc, _cdea)
		return _cgfc
	}
	_fgbg, _egd := _g.NewPdfOutputIntentFromPdfObject(_adga)
	if _egd != nil {
		_ge.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020i\u006et\u0065\u006e\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _egd)
		return _cgfc
	}
	_gbga, _egd := _gf.ParseHeader(_fgbg.DestOutputProfile)
	if _egd != nil {
		_ge.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066\u0061i\u006c\u0065d\u003a\u0020\u0025\u0076", _egd)
		return _cgfc
	}
	if (_gbga.DeviceClass == _gf.DeviceClassPRTR || _gbga.DeviceClass == _gf.DeviceClassMNTR) && (_gbga.ColorSpace == _gf.ColorSpaceRGB || _gbga.ColorSpace == _gf.ColorSpaceCMYK || _gbga.ColorSpace == _gf.ColorSpaceGRAY) {
		return _cgfc
	}
	_cgfc = append(_cgfc, _cdea)
	return _cgfc
}
func _efbd(_babbc *_cb.PdfObjectDictionary, _ecgcf map[*_cb.PdfObjectStream][]byte, _fdfad map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_fedad = "\u0036.\u0033\u002e\u0038\u002d\u0031"
		_cccg  = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_ffgfg, _dbab := _cb.GetStream(_babbc.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _dbab {
		_, _fagg := _cdfbg(_ffgfg, _ecgcf, _fdfad)
		if _fagg != nil {
			return _dd(_fedad, _cccg)
		}
		return _dfa
	}
	_cbde, _dbab := _cb.GetName(_babbc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_dbab {
		return _dd(_fedad, _cccg)
	}
	switch _cbde.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _dfa
	}
	return _dd(_fedad, _cccg)
}
func _ddda(_dgb *_g.CompliancePdfReader) (_febaf []ViolatedRule) {
	var (
		_ddfed, _ceab, _dged, _cbga, _ffead, _egeda, _eefb bool
		_ecgg                                              func(_cb.PdfObject)
	)
	_ecgg = func(_egfe _cb.PdfObject) {
		switch _bedd := _egfe.(type) {
		case *_cb.PdfObjectInteger:
			if !_ddfed && (int64(*_bedd) > _b.MaxInt32 || int64(*_bedd) < -_b.MaxInt32) {
				_febaf = append(_febaf, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_ddfed = true
			}
		case *_cb.PdfObjectFloat:
			if !_ceab && (_b.Abs(float64(*_bedd)) > 32767.0) {
				_febaf = append(_febaf, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0032", "\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0020\u0072\u0065\u0061\u006c\u0020\u0076\u0061\u006c\u0075\u0065\u0020m\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u006c\u0065s\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075a\u006c\u0020\u0074\u006f\u0020\u00332\u0037\u0036\u0037.\u0030\u002e"))
			}
		case *_cb.PdfObjectString:
			if !_dged && len([]byte(_bedd.Str())) > 65535 {
				_febaf = append(_febaf, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0036\u0035\u0035\u0033\u0035\u002e"))
				_dged = true
			}
		case *_cb.PdfObjectName:
			if !_cbga && len([]byte(*_bedd)) > 127 {
				_febaf = append(_febaf, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_cbga = true
			}
		case *_cb.PdfObjectArray:
			if !_ffead && _bedd.Len() > 8191 {
				_febaf = append(_febaf, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0035", "\u004d\u0061\u0078\u0069\u006d\u0075m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020(\u0069\u006e\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0029\u0020\u0069s\u00208\u0031\u0039\u0031\u002e"))
				_ffead = true
			}
			for _, _bedf := range _bedd.Elements() {
				_ecgg(_bedf)
			}
			if !_eefb && (_bedd.Len() == 4 || _bedd.Len() == 5) {
				_gddd, _eeed := _cb.GetName(_bedd.Get(0))
				if !_eeed {
					return
				}
				if *_gddd != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_cbee := _bedd.Get(1)
				_cbee = _cb.TraceToDirectObject(_cbee)
				_bcde, _eeed := _cb.GetArray(_cbee)
				if !_eeed {
					return
				}
				if _bcde.Len() > 8 {
					_febaf = append(_febaf, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0039", "\u004d\u0061\u0078i\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0069\u0073\u0020\u0038\u002e"))
					_eefb = true
				}
			}
		case *_cb.PdfObjectDictionary:
			_bece := _bedd.Keys()
			if !_egeda && len(_bece) > 4095 {
				_febaf = append(_febaf, _dd("\u0036.\u0031\u002e\u0031\u0032\u002d\u00311", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u0063\u0061\u0070\u0061\u0063\u0069\u0074y\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0028\u0069\u006e\u0020\u0065\u006e\u0074\u0072\u0069es\u0029\u0020\u0069\u0073\u0020\u0034\u0030\u0039\u0035\u002e"))
				_egeda = true
			}
			for _aabcf, _ccgf := range _bece {
				_ecgg(&_bece[_aabcf])
				_ecgg(_bedd.Get(_ccgf))
			}
		case *_cb.PdfObjectStream:
			_ecgg(_bedd.PdfObjectDictionary)
		case *_cb.PdfObjectStreams:
			for _, _fecd := range _bedd.Elements() {
				_ecgg(_fecd)
			}
		case *_cb.PdfObjectReference:
			_ecgg(_bedd.Resolve())
		}
	}
	_dfagg := _dgb.GetObjectNums()
	if len(_dfagg) > 8388607 {
		_febaf = append(_febaf, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _dbae := range _dfagg {
		_cffc, _fagc := _dgb.GetIndirectObjectByNumber(_dbae)
		if _fagc != nil {
			continue
		}
		_agbg := _cb.TraceToDirectObject(_cffc)
		_ecgg(_agbg)
	}
	return _febaf
}
func _cgdf(_faaa *_gc.Document) error {
	_ffc, _bga := _faaa.GetPages()
	if !_bga {
		return nil
	}
	for _, _gba := range _ffc {
		_afg := _gba.FindXObjectForms()
		for _, _dffg := range _afg {
			_dffgg, _gcg := _cb.GetDict(_dffg.Get("\u0047\u0072\u006fu\u0070"))
			if _gcg {
				if _gbf := _dffgg.Get("\u0053"); _gbf != nil {
					_bfdf, _bdd := _cb.GetName(_gbf)
					if _bdd && _bfdf.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_dffg.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_ecgc, _cfae := _gba.GetResourcesXObject()
		if _cfae {
			_baefe, _acg := _cb.GetDict(_ecgc.Get("\u0047\u0072\u006fu\u0070"))
			if _acg {
				_gffdb := _baefe.Get("\u0053")
				if _gffdb != nil {
					_cgbd, _gbee := _cb.GetName(_gffdb)
					if _gbee && _cgbd.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_ecgc.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_ffafb, _dbcd := _cb.GetDict(_gba.Object.Get("\u0047\u0072\u006fu\u0070"))
		if _dbcd {
			_ccg := _ffafb.Get("\u0053")
			if _ccg != nil {
				_adca, _agbf := _cb.GetName(_ccg)
				if _agbf && _adca.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
					_gba.Object.Remove("\u0047\u0072\u006fu\u0070")
				}
			}
		}
	}
	return nil
}
func _faab(_ebd *_g.PdfInfo, _dfega func() _bee.Time) error {
	var _begg *_g.PdfDate
	if _ebd.CreationDate == nil {
		_dgfd, _daa := _g.NewPdfDateFromTime(_dfega())
		if _daa != nil {
			return _daa
		}
		_begg = &_dgfd
		_ebd.CreationDate = _begg
	}
	if _ebd.ModifiedDate == nil {
		if _begg != nil {
			_fdde, _cfbd := _g.NewPdfDateFromTime(_dfega())
			if _cfbd != nil {
				return _cfbd
			}
			_begg = &_fdde
		}
		_ebd.ModifiedDate = _begg
	}
	return nil
}
func _gcbg(_gfgg *_g.CompliancePdfReader) (_geee []ViolatedRule) {
	for _, _acabe := range _gfgg.GetObjectNums() {
		_accfa, _gcgae := _gfgg.GetIndirectObjectByNumber(_acabe)
		if _gcgae != nil {
			continue
		}
		_dabaec, _ggcd := _cb.GetDict(_accfa)
		if !_ggcd {
			continue
		}
		_faff, _ggcd := _cb.GetName(_dabaec.Get("\u0054\u0079\u0070\u0065"))
		if !_ggcd {
			continue
		}
		if _faff.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_fcfg, _ggcd := _cb.GetBool(_dabaec.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if !_ggcd {
			return _geee
		}
		if bool(*_fcfg) {
			_geee = append(_geee, _dd("\u0036\u002e\u0039-\u0031", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
	}
	return _geee
}
func _fdfd(_eeec standardType, _fad *_gc.OutputIntents) error {
	_ddef, _badbc := _gf.NewCmykIsoCoatedV2OutputIntent(_eeec.outputIntentSubtype())
	if _badbc != nil {
		return _badbc
	}
	if _badbc = _fad.Add(_ddef.ToPdfObject()); _badbc != nil {
		return _badbc
	}
	return nil
}
func _gda(_egab *_gc.Document, _feb standardType, _eaag XmpOptions) error {
	_gade, _defg := _egab.FindCatalog()
	if !_defg {
		return nil
	}
	var _cee *_gga.Document
	_cgbb, _defg := _gade.GetMetadata()
	if !_defg {
		_cee = _gga.NewDocument()
	} else {
		var _dbc error
		_cee, _dbc = _gga.LoadDocument(_cgbb.Stream)
		if _dbc != nil {
			return _dbc
		}
	}
	_fdd := _gga.PdfInfoOptions{InfoDict: _egab.Info, PdfVersion: _c.Sprintf("\u0025\u0064\u002e%\u0064", _egab.Version.Major, _egab.Version.Minor), Copyright: _eaag.Copyright, Overwrite: true}
	_cca, _defg := _gade.GetMarkInfo()
	if _defg {
		_deda, _dec := _cb.GetBool(_cca.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
		if _dec && bool(*_deda) {
			_fdd.Marked = true
		}
	}
	if _cgf := _cee.SetPdfInfo(&_fdd); _cgf != nil {
		return _cgf
	}
	if _ecc := _cee.SetPdfAID(_feb._da, _feb._dag); _ecc != nil {
		return _ecc
	}
	_eaae := _gga.MediaManagementOptions{OriginalDocumentID: _eaag.OriginalDocumentID, DocumentID: _eaag.DocumentID, InstanceID: _eaag.InstanceID, NewDocumentID: !_eaag.NewDocumentVersion, ModifyComment: "O\u0070\u0074\u0069\u006d\u0069\u007ae\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u006f\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064"}
	_fac, _defg := _cb.GetDict(_egab.Info)
	if _defg {
		if _dfbga, _bbe := _cb.GetString(_fac.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _bbe && _dfbga.String() != "" {
			_cgfd, _ecec := _ac.ParsePdfTime(_dfbga.String())
			if _ecec != nil {
				return _c.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _ecec)
			}
			_eaae.ModifyDate = _cgfd
		}
	}
	if _dfc := _cee.SetMediaManagement(&_eaae); _dfc != nil {
		return _dfc
	}
	if _agd := _cee.SetPdfAExtension(); _agd != nil {
		return _agd
	}
	_bce, _fab := _cee.MarshalIndent(_eaag.MarshalPrefix, _eaag.MarshalIndent)
	if _fab != nil {
		return _fab
	}
	if _gfbg := _gade.SetMetadata(_bce); _gfbg != nil {
		return _gfbg
	}
	return nil
}
func _abgdf(_fdbc *_g.CompliancePdfReader) []ViolatedRule { return nil }

var _dfa = ViolatedRule{}

// Validate checks if provided input document reader matches given PDF/A profile.
func Validate(d *_g.CompliancePdfReader, profile Profile) error { return profile.ValidateStandard(d) }

type imageModifications struct {
	_bad *colorspaceModification
	_af  _cb.StreamEncoder
}

func _bfbc(_fgaf *_g.CompliancePdfReader) (_ebacb []ViolatedRule) {
	if _fgaf.ParserMetadata().HasOddLengthHexStrings() {
		_ebacb = append(_ebacb, _dd("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _fgaf.ParserMetadata().HasOddLengthHexStrings() {
		_ebacb = append(_ebacb, _dd("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _ebacb
}
func _dbagec(_cceac *_g.CompliancePdfReader, _fcede standardType) (_dgdbd []ViolatedRule) {
	var _gdac, _deca, _ebeebb, _feabb, _gegd, _dega, _fafg bool
	_accfc := func() bool { return _gdac && _deca && _ebeebb && _feabb && _gegd && _dega && _fafg }
	_dccf := map[*_cb.PdfObjectStream]*_ad.CMap{}
	_gffge := map[*_cb.PdfObjectStream][]byte{}
	_gcbc := map[_cb.PdfObject]*_g.PdfFont{}
	for _, _gggeb := range _cceac.GetObjectNums() {
		_ccfg, _aaegg := _cceac.GetIndirectObjectByNumber(_gggeb)
		if _aaegg != nil {
			continue
		}
		_bcdce, _gfad := _cb.GetDict(_ccfg)
		if !_gfad {
			continue
		}
		_babbg, _gfad := _cb.GetName(_bcdce.Get("\u0054\u0079\u0070\u0065"))
		if !_gfad {
			continue
		}
		if *_babbg != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_gcbge, _aaegg := _g.NewPdfFontFromPdfObject(_bcdce)
		if _aaegg != nil {
			_ge.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _aaegg)
			continue
		}
		_gcbc[_bcdce] = _gcbge
	}
	for _, _dbebg := range _cceac.PageList {
		_bedge, _acdb := _dbebg.GetContentStreams()
		if _acdb != nil {
			_ge.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _beeb := range _bedge {
			_fbfec := _ee.NewContentStreamParser(_beeb)
			_egbc, _ccab := _fbfec.Parse()
			if _ccab != nil {
				_ge.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ccab)
				continue
			}
			var _gdea bool
			for _, _ffagg := range *_egbc {
				if _ffagg.Operand != "\u0054\u0072" {
					continue
				}
				if len(_ffagg.Params) != 1 {
					_ge.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_ffagg.Params))
					continue
				}
				_cadec, _fccgb := _cb.GetIntVal(_ffagg.Params[0])
				if !_fccgb {
					_ge.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _cadec == 3 {
					_gdea = true
					break
				}
			}
			for _, _bcdge := range *_egbc {
				if _bcdge.Operand != "\u0054\u0066" {
					continue
				}
				if len(_bcdge.Params) != 2 {
					_ge.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_bcdge.Params))
					continue
				}
				_aace, _gbdb := _cb.GetName(_bcdge.Params[0])
				if !_gbdb {
					_ge.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _bcdge)
					continue
				}
				_dcegf, _dbcgb := _dbebg.Resources.GetFontByName(*_aace)
				if !_dbcgb {
					_ge.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_fbefg, _gbdb := _cb.GetDict(_dcegf)
				if !_gbdb {
					_ge.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_gfcga, _gbdb := _gcbc[_fbefg]
				if !_gbdb {
					var _dbfbf error
					_gfcga, _dbfbf = _g.NewPdfFontFromPdfObject(_fbefg)
					if _dbfbf != nil {
						_ge.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _dbfbf)
						continue
					}
					_gcbc[_fbefg] = _gfcga
				}
				if !_gdac {
					_aedad := _ffce(_fbefg, _gffge, _dccf)
					if _aedad != _dfa {
						_dgdbd = append(_dgdbd, _aedad)
						_gdac = true
						if _accfc() {
							return _dgdbd
						}
					}
				}
				if !_deca {
					_ccagb := _aeffe(_fbefg)
					if _ccagb != _dfa {
						_dgdbd = append(_dgdbd, _ccagb)
						_deca = true
						if _accfc() {
							return _dgdbd
						}
					}
				}
				if !_ebeebb {
					_efff := _daebe(_fbefg, _gffge, _dccf)
					if _efff != _dfa {
						_dgdbd = append(_dgdbd, _efff)
						_ebeebb = true
						if _accfc() {
							return _dgdbd
						}
					}
				}
				if !_feabb {
					_febdb := _gdbe(_fbefg, _gffge, _dccf)
					if _febdb != _dfa {
						_dgdbd = append(_dgdbd, _febdb)
						_feabb = true
						if _accfc() {
							return _dgdbd
						}
					}
				}
				if !_gegd {
					_daecd := _abbe(_gfcga, _fbefg, _gdea)
					if _daecd != _dfa {
						_gegd = true
						_dgdbd = append(_dgdbd, _daecd)
						if _accfc() {
							return _dgdbd
						}
					}
				}
				if !_dega {
					_eebg := _fedaa(_gfcga, _fbefg)
					if _eebg != _dfa {
						_dega = true
						_dgdbd = append(_dgdbd, _eebg)
						if _accfc() {
							return _dgdbd
						}
					}
				}
				if !_fafg && (_fcede._dag == "\u0041" || _fcede._dag == "\u0055") {
					_ddec := _fada(_fbefg, _gffge, _dccf)
					if _ddec != _dfa {
						_fafg = true
						_dgdbd = append(_dgdbd, _ddec)
						if _accfc() {
							return _dgdbd
						}
					}
				}
			}
		}
	}
	return _dgdbd
}
func _fcccb(_agea *_g.CompliancePdfReader) (_eeca ViolatedRule) {
	_cfgf, _eeab := _addf(_agea)
	if !_eeab {
		return _dfa
	}
	if _cfgf.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		return _dd("\u0036\u002e\u0031\u0031\u002d\u0031", "Th\u0065\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063a\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020R\u0065q\u0075\u0069\u0072\u0065\u006d\u0065\u006e\u0074s\u0020k\u0065\u0079.")
	}
	return _dfa
}

var _ Profile = (*Profile2A)(nil)

func _ccgbe(_adce *_g.CompliancePdfReader) ViolatedRule {
	_bdfe := _adce.ParserMetadata().HeaderCommentBytes()
	if _bdfe[0] > 127 && _bdfe[1] > 127 && _bdfe[2] > 127 && _bdfe[3] > 127 {
		return _dfa
	}
	return _dd("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}

var _ Profile = (*Profile1A)(nil)

func (_bg standardType) outputIntentSubtype() _g.PdfOutputIntentType {
	switch _bg._da {
	case 1:
		return _g.PdfOutputIntentTypeA1
	case 2:
		return _g.PdfOutputIntentTypeA2
	case 3:
		return _g.PdfOutputIntentTypeA3
	case 4:
		return _g.PdfOutputIntentTypeA4
	default:
		return 0
	}
}
func _acade(_acbbeg *_g.CompliancePdfReader) (_deggc []ViolatedRule) {
	var (
		_aaeag, _ebfc, _abac, _bdga, _cagaf bool
		_bgfgb                              func(_cb.PdfObject)
	)
	_bgfgb = func(_dedg _cb.PdfObject) {
		switch _agfff := _dedg.(type) {
		case *_cb.PdfObjectInteger:
			if !_aaeag && (int64(*_agfff) > _b.MaxInt32 || int64(*_agfff) < -_b.MaxInt32) {
				_deggc = append(_deggc, _dd("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_aaeag = true
			}
		case *_cb.PdfObjectFloat:
			if !_ebfc && (_b.Abs(float64(*_agfff)) > _b.MaxFloat32) {
				_deggc = append(_deggc, _dd("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0032", "\u0041 \u0063\u006f\u006e\u0066orm\u0069\u006e\u0067\u0020f\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0061\u006c\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0075\u0074\u0073\u0069de\u0020\u0074\u0068e\u0020\u0072\u0061\u006e\u0067e\u0020o\u0066\u0020\u002b\u002f\u002d\u0033\u002e\u0034\u00303\u0020\u0078\u0020\u0031\u0030\u005e\u0033\u0038\u002e"))
			}
		case *_cb.PdfObjectString:
			if !_abac && len([]byte(_agfff.Str())) > 32767 {
				_deggc = append(_deggc, _dd("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0033\u0032\u0037\u0036\u0037\u002e"))
				_abac = true
			}
		case *_cb.PdfObjectName:
			if !_bdga && len([]byte(*_agfff)) > 127 {
				_deggc = append(_deggc, _dd("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_bdga = true
			}
		case *_cb.PdfObjectArray:
			for _, _egda := range _agfff.Elements() {
				_bgfgb(_egda)
			}
			if !_cagaf && (_agfff.Len() == 4 || _agfff.Len() == 5) {
				_gcgf, _cbgf := _cb.GetName(_agfff.Get(0))
				if !_cbgf {
					return
				}
				if *_gcgf != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_cbaef := _agfff.Get(1)
				_cbaef = _cb.TraceToDirectObject(_cbaef)
				_gabda, _cbgf := _cb.GetArray(_cbaef)
				if !_cbgf {
					return
				}
				if _gabda.Len() > 32 {
					_deggc = append(_deggc, _dd("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0039", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006e\u0075\u006db\u0065\u0072\u0020\u006ff\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0069\u0073\u0020\u0033\u0032\u002e"))
					_cagaf = true
				}
			}
		case *_cb.PdfObjectDictionary:
			_ggfb := _agfff.Keys()
			for _gacde, _adcee := range _ggfb {
				_bgfgb(&_ggfb[_gacde])
				_bgfgb(_agfff.Get(_adcee))
			}
		case *_cb.PdfObjectStream:
			_bgfgb(_agfff.PdfObjectDictionary)
		case *_cb.PdfObjectStreams:
			for _, _ffab := range _agfff.Elements() {
				_bgfgb(_ffab)
			}
		case *_cb.PdfObjectReference:
			_bgfgb(_agfff.Resolve())
		}
	}
	_dddde := _acbbeg.GetObjectNums()
	if len(_dddde) > 8388607 {
		_deggc = append(_deggc, _dd("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _fcbc := range _dddde {
		_ffca, _bdag := _acbbeg.GetIndirectObjectByNumber(_fcbc)
		if _bdag != nil {
			continue
		}
		_eefdb := _cb.TraceToDirectObject(_ffca)
		_bgfgb(_eefdb)
	}
	return _deggc
}
func _dbbeb(_fgcd *_g.CompliancePdfReader) ViolatedRule {
	if _fgcd.ParserMetadata().HasDataAfterEOF() {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _dfa
}
func _cgcc(_bcaa *_gc.Document) error {
	_ffea, _eafe := _bcaa.FindCatalog()
	if !_eafe {
		return _f.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_ccad, _eafe := _cb.GetDict(_ffea.Object.Get("\u004e\u0061\u006de\u0073"))
	if !_eafe {
		return nil
	}
	if _ccad.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_ccad.Remove("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073")
	}
	return nil
}

// NewProfile2A creates a new Profile2A with given options.
func NewProfile2A(options *Profile2Options) *Profile2A {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_eecgb(options)
	return &Profile2A{profile2{_afef: *options, _gaeg: _bf()}}
}

// Profile2U is the implementation of the PDF/A-2U standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2U struct{ profile2 }

func _dbbc(_bggf *_g.CompliancePdfReader, _age bool) (_ddeg []ViolatedRule) {
	var _ccebc, _efec, _eeac, _gdc, _bfdg, _bggb, _bgda bool
	_edfe := func() bool { return _ccebc && _efec && _eeac && _gdc && _bfdg && _bggb && _bgda }
	_acf, _cfbf := _gdde(_bggf)
	var _dgae _gf.ProfileHeader
	if _cfbf {
		_dgae, _ = _gf.ParseHeader(_acf.DestOutputProfile)
	}
	var _ade bool
	_gbcg := map[_cb.PdfObject]struct{}{}
	var _cfcd func(_aabe _g.PdfColorspace) bool
	_cfcd = func(_dabae _g.PdfColorspace) bool {
		switch _aefa := _dabae.(type) {
		case *_g.PdfColorspaceDeviceGray:
			if !_bggb {
				if !_cfbf {
					_ade = true
					_ddeg = append(_ddeg, _dd("\u0036.\u0032\u002e\u0033\u002d\u0034", "\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0075s\u0065\u0064\u0020\u006f\u006el\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006ce\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020O\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u002e"))
					_bggb = true
					if _edfe() {
						return true
					}
				}
			}
		case *_g.PdfColorspaceDeviceRGB:
			if !_gdc {
				if !_cfbf || _dgae.ColorSpace != _gf.ColorSpaceRGB {
					_ade = true
					_ddeg = append(_ddeg, _dd("\u0036.\u0032\u002e\u0033\u002d\u0032", "\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065 \u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074In\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020u\u0073es\u0020a\u006e\u0020\u0052\u0047\u0042\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u002e"))
					_gdc = true
					if _edfe() {
						return true
					}
				}
			}
		case *_g.PdfColorspaceDeviceCMYK:
			if !_bfdg {
				if !_cfbf || _dgae.ColorSpace != _gf.ColorSpaceCMYK {
					_ade = true
					_ddeg = append(_ddeg, _dd("\u0036.\u0032\u002e\u0033\u002d\u0033", "\u0044\u0065\u0076\u0069\u0063e\u0043\u004d\u0059\u004b \u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074h\u0065\u0020\u0066\u0069\u006ce \u0068\u0061\u0073\u0020\u0061 \u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068a\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0061\u006e \u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072\u0020s\u0070\u0061\u0063e\u002e"))
					_bfdg = true
					if _edfe() {
						return true
					}
				}
			}
		case *_g.PdfColorspaceICCBased:
			if !_eeac || !_bgda {
				_baab, _bgbag := _gf.ParseHeader(_aefa.Data)
				if _bgbag != nil {
					_ge.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _bgbag)
					_ddeg = append(_ddeg, func() ViolatedRule {
						return _dd("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035")
					}())
					_eeac = true
					if _edfe() {
						return true
					}
				}
				if !_eeac {
					var _degdc, _babb bool
					switch _baab.DeviceClass {
					case _gf.DeviceClassPRTR, _gf.DeviceClassMNTR, _gf.DeviceClassSCNR, _gf.DeviceClassSPAC:
					default:
						_degdc = true
					}
					switch _baab.ColorSpace {
					case _gf.ColorSpaceRGB, _gf.ColorSpaceCMYK, _gf.ColorSpaceGRAY, _gf.ColorSpaceLAB:
					default:
						_babb = true
					}
					if _degdc || _babb {
						_ddeg = append(_ddeg, _dd("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035"))
						_eeac = true
						if _edfe() {
							return true
						}
					}
				}
				if !_bgda {
					_fbga, _ := _cb.GetStream(_aefa.GetContainingPdfObject())
					if _fbga.Get("\u004e") == nil || (_aefa.N == 1 && _baab.ColorSpace != _gf.ColorSpaceGRAY) || (_aefa.N == 3 && !(_baab.ColorSpace == _gf.ColorSpaceRGB || _baab.ColorSpace == _gf.ColorSpaceLAB)) || (_aefa.N == 4 && _baab.ColorSpace != _gf.ColorSpaceCMYK) {
						_ddeg = append(_ddeg, _dd("\u0036.\u0032\u002e\u0033\u002d\u0035", "\u0049\u0066\u0020a\u006e\u0020u\u006e\u0063\u0061\u006c\u0069\u0062\u0072a\u0074\u0065\u0064\u0020\u0063\u006fl\u006f\u0072 \u0073\u0070\u0061c\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u0069\u006c\u0065 \u0074\u0068\u0065\u006e \u0074\u0068\u0061\u0074 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041-\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020\u0069\u006e\u0020\u0036\u002e\u0032\u002e\u0032\u002e"))
						_bgda = true
						if _edfe() {
							return true
						}
					}
				}
			}
			if _aefa.Alternate != nil {
				return _cfcd(_aefa.Alternate)
			}
		}
		return false
	}
	for _, _agce := range _bggf.GetObjectNums() {
		_efgf, _cdfb := _bggf.GetIndirectObjectByNumber(_agce)
		if _cdfb != nil {
			continue
		}
		_dbca, _aefd := _cb.GetStream(_efgf)
		if !_aefd {
			continue
		}
		_acec, _aefd := _cb.GetName(_dbca.Get("\u0054\u0079\u0070\u0065"))
		if !_aefd || _acec.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_ecdc, _aefd := _cb.GetName(_dbca.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_aefd {
			continue
		}
		_gbcg[_dbca] = struct{}{}
		switch _ecdc.String() {
		case "\u0049\u006d\u0061g\u0065":
			_cbdd, _ceag := _g.NewXObjectImageFromStream(_dbca)
			if _ceag != nil {
				continue
			}
			_gbcg[_dbca] = struct{}{}
			if _cfcd(_cbdd.ColorSpace) {
				return _ddeg
			}
		case "\u0046\u006f\u0072\u006d":
			_beaa, _ffgfa := _cb.GetDict(_dbca.Get("\u0047\u0072\u006fu\u0070"))
			if !_ffgfa {
				continue
			}
			_bcded := _beaa.Get("\u0043\u0053")
			if _bcded == nil {
				continue
			}
			_deced, _aabd := _g.NewPdfColorspaceFromPdfObject(_bcded)
			if _aabd != nil {
				continue
			}
			if _cfcd(_deced) {
				return _ddeg
			}
		}
	}
	for _, _gbbf := range _bggf.PageList {
		_gffb, _gdgf := _gbbf.GetContentStreams()
		if _gdgf != nil {
			continue
		}
		for _, _abgfe := range _gffb {
			_deadf, _adea := _ee.NewContentStreamParser(_abgfe).Parse()
			if _adea != nil {
				continue
			}
			for _, _aaab := range *_deadf {
				if len(_aaab.Params) > 1 {
					continue
				}
				switch _aaab.Operand {
				case "\u0042\u0049":
					_fdbe, _fecb := _aaab.Params[0].(*_ee.ContentStreamInlineImage)
					if !_fecb {
						continue
					}
					_ffgae, _aafa := _fdbe.GetColorSpace(_gbbf.Resources)
					if _aafa != nil {
						continue
					}
					if _cfcd(_ffgae) {
						return _ddeg
					}
				case "\u0044\u006f":
					_afc, _fecge := _cb.GetName(_aaab.Params[0])
					if !_fecge {
						continue
					}
					_eedb, _cbce := _gbbf.Resources.GetXObjectByName(*_afc)
					if _, _dada := _gbcg[_eedb]; _dada {
						continue
					}
					switch _cbce {
					case _g.XObjectTypeImage:
						_gfge, _gfab := _g.NewXObjectImageFromStream(_eedb)
						if _gfab != nil {
							continue
						}
						_gbcg[_eedb] = struct{}{}
						if _cfcd(_gfge.ColorSpace) {
							return _ddeg
						}
					case _g.XObjectTypeForm:
						_ccade, _gefdc := _cb.GetDict(_eedb.Get("\u0047\u0072\u006fu\u0070"))
						if !_gefdc {
							continue
						}
						_bcdb, _gefdc := _cb.GetName(_ccade.Get("\u0043\u0053"))
						if !_gefdc {
							continue
						}
						_gdfa, _aegd := _g.NewPdfColorspaceFromPdfObject(_bcdb)
						if _aegd != nil {
							continue
						}
						_gbcg[_eedb] = struct{}{}
						if _cfcd(_gdfa) {
							return _ddeg
						}
					}
				}
			}
		}
	}
	if !_ade {
		return _ddeg
	}
	if (_dgae.DeviceClass == _gf.DeviceClassPRTR || _dgae.DeviceClass == _gf.DeviceClassMNTR) && (_dgae.ColorSpace == _gf.ColorSpaceRGB || _dgae.ColorSpace == _gf.ColorSpaceCMYK || _dgae.ColorSpace == _gf.ColorSpaceGRAY) {
		return _ddeg
	}
	if !_age {
		return _ddeg
	}
	_cbddb, _agced := _addf(_bggf)
	if !_agced {
		return _ddeg
	}
	_fcbb, _agced := _cb.GetArray(_cbddb.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_agced {
		_ddeg = append(_ddeg, _dd("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e"), _dd("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
		return _ddeg
	}
	if _fcbb.Len() > 1 {
		_eefa := map[*_cb.PdfObjectDictionary]struct{}{}
		for _daec := 0; _daec < _fcbb.Len(); _daec++ {
			_fcbd, _ceea := _cb.GetDict(_fcbb.Get(_daec))
			if !_ceea {
				continue
			}
			if _daec == 0 {
				_eefa[_fcbd] = struct{}{}
				continue
			}
			if _, _bfgf := _eefa[_fcbd]; !_bfgf {
				_ddeg = append(_ddeg, _dd("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
				break
			}
		}
	}
	return _ddeg
}
func _edbg(_ddgg *_g.CompliancePdfReader) (_gffeg ViolatedRule) {
	_fcaea, _daae := _addf(_ddgg)
	if !_daae {
		return _dfa
	}
	if _fcaea.Get("\u0041\u0041") != nil {
		return _dd("\u0036.\u0036\u002e\u0032\u002d\u0033", "\u0054\u0068e\u0020\u0064\u006f\u0063\u0075\u006d\u0065n\u0074 \u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065n\u0074r\u0079 \u0066\u006f\u0072 \u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063\u0074i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
	}
	return _dfa
}

type pageColorspaceOptimizeFunc func(_feba *_gc.Document, _gcaa *_gc.Page, _dgc []*_gc.Image) error

func _abbe(_dgfa *_g.PdfFont, _cdfdc *_cb.PdfObjectDictionary, _fdfdc bool) ViolatedRule {
	const (
		_ebcd  = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0034\u002d\u0031"
		_gffea = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006ds\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020f\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0072e\u006e\u0064\u0065\u0072\u0069\u006eg\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020w\u0069t\u0068\u0069\u006e\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u002c \u0061\u0073\u0020\u0064\u0065\u0066\u0069n\u0065\u0064 \u0069\u006e\u0020\u0049S\u004f\u0020\u0033\u0032\u00300\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0039\u002e"
	)
	if _fdfdc {
		return _dfa
	}
	_eeafe := _dgfa.FontDescriptor()
	var _bddd string
	if _ffced, _adgg := _cb.GetName(_cdfdc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _adgg {
		_bddd = _ffced.String()
	}
	switch _bddd {
	case "\u0054\u0079\u0070e\u0031":
		if _eeafe.FontFile == nil {
			return _dd(_ebcd, _gffea)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _eeafe.FontFile2 == nil {
			return _dd(_ebcd, _gffea)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _eeafe.FontFile3 == nil {
			return _dd(_ebcd, _gffea)
		}
	}
	return _dfa
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-2 standard.
func (_aca *profile2) ValidateStandard(r *_g.CompliancePdfReader) error {
	_dbaf := VerificationError{ConformanceLevel: _aca._gaeg._da, ConformanceVariant: _aca._gaeg._dag}
	if _agc := _dgedf(r); _agc != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _agc)
	}
	if _eadg := _gcgb(r); _eadg != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _eadg)
	}
	if _deaeb := _abaa(r); _deaeb != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _deaeb)
	}
	if _abcf := _fcbf(r); _abcf != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _abcf)
	}
	if _ffd := _egfcf(r); _ffd != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _ffd)
	}
	if _dcf := _bfbc(r); len(_dcf) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _dcf...)
	}
	if _eggc := _gagc(r); len(_eggc) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _eggc...)
	}
	if _gabe := _abgdf(r); len(_gabe) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _gabe...)
	}
	if _gaaca := _edbe(r); _gaaca != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _gaaca)
	}
	if _bda := _aegc(r); len(_bda) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _bda...)
	}
	if _adg := _acade(r); len(_adg) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _adg...)
	}
	if _cbae := _baed(r); _cbae != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _cbae)
	}
	if _ffdf := _becgc(r); len(_ffdf) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _ffdf...)
	}
	if _ebdb := _ccfag(r); len(_ebdb) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _ebdb...)
	}
	if _ebef := _bacg(r); _ebef != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _ebef)
	}
	if _dgacf := _cdcca(r); len(_dgacf) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _dgacf...)
	}
	if _gggg := _ebcbc(r); len(_gggg) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _gggg...)
	}
	if _bega := _ebaff(r); _bega != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _bega)
	}
	if _bedb := _ecbg(r); len(_bedb) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _bedb...)
	}
	if _gbec := _dbagec(r, _aca._gaeg); len(_gbec) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _gbec...)
	}
	if _gdf := _fbeae(r); len(_gdf) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _gdf...)
	}
	if _caga := _ggacc(r); len(_caga) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _caga...)
	}
	if _dfad := _afagd(r); len(_dfad) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _dfad...)
	}
	if _decb := _fgdd(r); _decb != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _decb)
	}
	if _cadb := _aecf(r); len(_cadb) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _cadb...)
	}
	if _ceca := _ddbff(r); _ceca != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _ceca)
	}
	if _defcb := _aagg(r, _aca._gaeg, false); len(_defcb) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _defcb...)
	}
	if _aca._gaeg == _bf() {
		if _cabf := _fbbb(r); len(_cabf) != 0 {
			_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _cabf...)
		}
	}
	if _fdgc := _bagb(r); len(_fdgc) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _fdgc...)
	}
	if _baaf := _gefdb(r); len(_baaf) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _baaf...)
	}
	if _eaac := _deec(r); len(_eaac) != 0 {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _eaac...)
	}
	if _ebff := _fcccb(r); _ebff != _dfa {
		_dbaf.ViolatedRules = append(_dbaf.ViolatedRules, _ebff)
	}
	if len(_dbaf.ViolatedRules) > 0 {
		_be.Slice(_dbaf.ViolatedRules, func(_affc, _fbae int) bool {
			return _dbaf.ViolatedRules[_affc].RuleNo < _dbaf.ViolatedRules[_fbae].RuleNo
		})
		return _dbaf
	}
	return nil
}

// StandardName gets the name of the standard.
func (_aaee *profile2) StandardName() string {
	return _c.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0032\u0025s", _aaee._gaeg._dag)
}
func _cfdg(_abfe *Profile1Options) {
	if _abfe.Now == nil {
		_abfe.Now = _bee.Now
	}
}
func _cgcb(_bdgd *_gc.Document, _aae []pageColorspaceOptimizeFunc, _bbd []documentColorspaceOptimizeFunc) error {
	_afa, _bgbc := _bdgd.GetPages()
	if !_bgbc {
		return nil
	}
	var _abgb []*_gc.Image
	for _bbge, _edd := range _afa {
		_ceba, _gcaad := _edd.FindXObjectImages()
		if _gcaad != nil {
			return _gcaad
		}
		for _, _agg := range _aae {
			if _gcaad = _agg(_bdgd, &_afa[_bbge], _ceba); _gcaad != nil {
				return _gcaad
			}
		}
		_abgb = append(_abgb, _ceba...)
	}
	for _, _ccag := range _bbd {
		if _dbe := _ccag(_bdgd, _abgb); _dbe != nil {
			return _dbe
		}
	}
	return nil
}
func _dfgf(_dcfd *_g.CompliancePdfReader) (_eegd ViolatedRule) {
	for _, _egge := range _dcfd.GetObjectNums() {
		_fgefe, _abfd := _dcfd.GetIndirectObjectByNumber(_egge)
		if _abfd != nil {
			continue
		}
		_eaef, _egbf := _cb.GetStream(_fgefe)
		if !_egbf {
			continue
		}
		_gdfg, _egbf := _cb.GetName(_eaef.Get("\u0054\u0079\u0070\u0065"))
		if !_egbf {
			continue
		}
		if *_gdfg != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_abbdb, _egbf := _cb.GetName(_eaef.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_egbf {
			continue
		}
		if *_abbdb == "\u0050\u0053" {
			return _dd("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073.")
		}
	}
	return _eegd
}
func _ebga(_gffbc *_g.CompliancePdfReader) (*_cb.PdfObjectDictionary, bool) {
	_bdeb, _dggb := _addf(_gffbc)
	if !_dggb {
		return nil, false
	}
	_efdg, _dggb := _cb.GetArray(_bdeb.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_dggb {
		return nil, false
	}
	if _efdg.Len() == 0 {
		return nil, false
	}
	return _cb.GetDict(_efdg.Get(0))
}
func _gcd(_ddeb *_gc.Document) error {
	_acc, _eeff := _ddeb.FindCatalog()
	if !_eeff {
		return nil
	}
	_, _eeff = _cb.GetDict(_acc.Object.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
	if !_eeff {
		_ddfe := _cb.MakeDict()
		_ddfe.Set("\u0054\u0079\u0070\u0065", _cb.MakeName("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
		_acc.Object.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074", _ddfe)
	}
	return nil
}
func _bcea(_cgbc *_gc.Document) error {
	for _, _degd := range _cgbc.Objects {
		_cdedf, _cfeb := _cb.GetDict(_degd)
		if !_cfeb {
			continue
		}
		_ddfa := _cdedf.Get("\u0054\u0079\u0070\u0065")
		if _ddfa == nil {
			continue
		}
		if _acef, _dbeba := _cb.GetName(_ddfa); _dbeba && _acef.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_cgbba, _dabc := _cb.GetBool(_cdedf.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _dabc && bool(*_cgbba) {
			_cdedf.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cb.MakeBool(false))
		}
		if _cdedf.Get("\u0058\u0046\u0041") != nil {
			_cdedf.Remove("\u0058\u0046\u0041")
		}
	}
	_bcgc, _bfg := _cgbc.FindCatalog()
	if !_bfg {
		return _f.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _bcgc.Object.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_bcgc.Object.Remove("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067")
	}
	return nil
}

// DefaultProfile2Options are the default options for the Profile2.
func DefaultProfile2Options() *Profile2Options {
	return &Profile2Options{Now: _bee.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}
func _fbb(_ffg bool, _bdbg standardType) (pageColorspaceOptimizeFunc, documentColorspaceOptimizeFunc) {
	var _bae, _adba, _babc bool
	_bdf := func(_bbf *_gc.Document, _feg *_gc.Page, _bbec []*_gc.Image) error {
		for _, _ggb := range _bbec {
			switch _ggb.Colorspace {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				_adba = true
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				_bae = true
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				_babc = true
			}
		}
		_egcf, _dgd := _feg.GetContents()
		if !_dgd {
			return nil
		}
		for _, _gfc := range _egcf {
			_fgde, _baeb := _gfc.GetData()
			if _baeb != nil {
				continue
			}
			_fgf := _ee.NewContentStreamParser(string(_fgde))
			_ddcd, _baeb := _fgf.Parse()
			if _baeb != nil {
				continue
			}
			for _, _fcfd := range *_ddcd {
				switch _fcfd.Operand {
				case "\u0047", "\u0067":
					_adba = true
				case "\u0052\u0047", "\u0072\u0067":
					_bae = true
				case "\u004b", "\u006b":
					_babc = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_fcfd.Params) == 0 {
						continue
					}
					_bbee, _eaga := _cb.GetName(_fcfd.Params[0])
					if !_eaga {
						continue
					}
					switch _bbee.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_bae = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_adba = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_babc = true
					}
				}
			}
		}
		_edae := _feg.FindXObjectForms()
		for _, _caeg := range _edae {
			_bgba := _ee.NewContentStreamParser(string(_caeg.Stream))
			_eeaf, _eadf := _bgba.Parse()
			if _eadf != nil {
				continue
			}
			for _, _ggce := range *_eeaf {
				switch _ggce.Operand {
				case "\u0047", "\u0067":
					_adba = true
				case "\u0052\u0047", "\u0072\u0067":
					_bae = true
				case "\u004b", "\u006b":
					_babc = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_ggce.Params) == 0 {
						continue
					}
					_caad, _aaea := _cb.GetName(_ggce.Params[0])
					if !_aaea {
						continue
					}
					switch _caad.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_bae = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_adba = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_babc = true
					}
				}
			}
			_bfae, _ebc := _cb.GetArray(_feg.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
			if !_ebc {
				return nil
			}
			for _, _fbcc := range _bfae.Elements() {
				_gffd, _ebcg := _cb.GetDict(_fbcc)
				if !_ebcg {
					continue
				}
				_ceaf := _gffd.Get("\u0043")
				if _ceaf == nil {
					continue
				}
				_gbb, _ebcg := _cb.GetArray(_ceaf)
				if !_ebcg {
					continue
				}
				switch _gbb.Len() {
				case 0:
				case 1:
					_adba = true
				case 3:
					_bae = true
				case 4:
					_babc = true
				}
			}
		}
		return nil
	}
	_dad := func(_ede *_gc.Document, _cbdb []*_gc.Image) error {
		_cbef, _cfb := _ede.FindCatalog()
		if !_cfb {
			return nil
		}
		_gaab, _cfb := _cbef.GetOutputIntents()
		if _cfb && _gaab.Len() > 0 {
			return nil
		}
		if !_cfb {
			_gaab = _cbef.NewOutputIntents()
		}
		if !(_bae || _babc || _adba) {
			return nil
		}
		defer _cbef.SetOutputIntents(_gaab)
		if _bae && !_babc && !_adba {
			return _cbdg(_ede, _bdbg, _gaab)
		}
		if _babc && !_bae && !_adba {
			return _fdfd(_bdbg, _gaab)
		}
		if _adba && !_bae && !_babc {
			return _dgda(_bdbg, _gaab)
		}
		if _bae && _babc {
			if _bbdf := _ef(_cbdb, _ffg); _bbdf != nil {
				return _bbdf
			}
			if _ddb := _ged(_ede, _ffg); _ddb != nil {
				return _ddb
			}
			if _fba := _caae(_ede, _ffg); _fba != nil {
				return _fba
			}
			if _dbbd := _ggcg(_ede, _ffg); _dbbd != nil {
				return _dbbd
			}
			if _ffg {
				return _fdfd(_bdbg, _gaab)
			}
			return _cbdg(_ede, _bdbg, _gaab)
		}
		return nil
	}
	return _bdf, _dad
}
func _adaf(_feab *_g.CompliancePdfReader) (_fcfcc ViolatedRule) {
	for _, _eadb := range _feab.GetObjectNums() {
		_affcc, _cafe := _feab.GetIndirectObjectByNumber(_eadb)
		if _cafe != nil {
			continue
		}
		_ecfb, _cdg := _cb.GetStream(_affcc)
		if !_cdg {
			continue
		}
		_bceb, _cdg := _cb.GetName(_ecfb.Get("\u0054\u0079\u0070\u0065"))
		if !_cdg {
			continue
		}
		if *_bceb != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _ecfb.Get("\u0052\u0065\u0066") != nil {
			return _dd("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e")
		}
	}
	return _fcfcc
}
func _ffbbb(_efac *_g.PdfFont, _ddgdc *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_ggcge = "\u0036.\u0033\u002e\u0035\u002d\u0032"
		_addb  = "\u0046\u006f\u0072\u0020\u0061l\u006c\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074 \u0073\u0075bs\u0065\u0074\u0073 \u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074he\u0020f\u006f\u006e\u0074\u0020\u0064\u0065s\u0063r\u0069\u0070\u0074o\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006ec\u006c\u0075\u0064e\u0020\u0061\u0020\u0043\u0068\u0061\u0072\u0053\u0065\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u006c\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072 \u006e\u0061\u006d\u0065\u0073\u0020d\u0065\u0066i\u006e\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020f\u006f\u006e\u0074\u0020s\u0075\u0062\u0073\u0065\u0074, \u0061\u0073 \u0064\u0065s\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e \u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0061\u0062\u006ce\u0020\u0035\u002e1\u0038\u002e"
	)
	var _fgcde string
	if _gegc, _ebceg := _cb.GetName(_ddgdc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ebceg {
		_fgcde = _gegc.String()
	}
	if _fgcde != "\u0054\u0079\u0070e\u0031" {
		return _dfa
	}
	if _ae.IsStdFont(_ae.StdFontName(_efac.BaseFont())) {
		return _dfa
	}
	_abgba := _efac.FontDescriptor()
	if _abgba.CharSet == nil {
		return _dd(_ggcge, _addb)
	}
	return _dfa
}
func _addf(_acbfc *_g.CompliancePdfReader) (*_cb.PdfObjectDictionary, bool) {
	_eebd, _beba := _acbfc.GetTrailer()
	if _beba != nil {
		_ge.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _beba)
		return nil, false
	}
	_bcec, _dadd := _eebd.Get("\u0052\u006f\u006f\u0074").(*_cb.PdfObjectReference)
	if !_dadd {
		_ge.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return nil, false
	}
	_fdfc, _dadd := _cb.GetDict(_cb.ResolveReference(_bcec))
	if !_dadd {
		_ge.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return nil, false
	}
	return _fdfc, true
}
func _cbb() standardType { return standardType{_da: 1, _dag: "\u0041"} }
func _fbcg(_ecf, _gbfb, _ddbf, _agaa string) (string, bool) {
	_accf := _a.Index(_ecf, _gbfb)
	if _accf == -1 {
		return "", false
	}
	_dbg := _a.Index(_ecf, _ddbf)
	if _dbg == -1 {
		return "", false
	}
	if _dbg < _accf {
		return "", false
	}
	return _ecf[:_accf] + _gbfb + _agaa + _ecf[_dbg:], true
}
func _cbcbe(_cbad *_g.CompliancePdfReader) ViolatedRule {
	if _cbad.ParserMetadata().HeaderPosition() != 0 {
		return _dd("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	return _dfa
}
func _ceac(_gcgd *_g.CompliancePdfReader) (_dfec []ViolatedRule) {
	var _eacb, _gbac, _ccea, _ffbbd, _bfdd, _gdae, _cdge bool
	_bfbe := func() bool { return _eacb && _gbac && _ccea && _ffbbd && _bfdd && _gdae && _cdge }
	for _, _edab := range _gcgd.PageList {
		if _edab.Resources == nil {
			continue
		}
		_agdb, _aafd := _cb.GetDict(_edab.Resources.Font)
		if !_aafd {
			continue
		}
		for _, _bcag := range _agdb.Keys() {
			_acab, _cbabb := _cb.GetDict(_agdb.Get(_bcag))
			if !_cbabb {
				if !_eacb {
					_dfec = append(_dfec, _dd("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006fn\u0074\u0073\u0020\u0075\u0073e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0066\u006f\u0072m\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0073\u0020d\u0065\u0066\u0069\u006e\u0065d \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035\u002e\u0035\u002e"))
					_eacb = true
					if _bfbe() {
						return _dfec
					}
				}
				continue
			}
			if _gfae, _afag := _cb.GetName(_acab.Get("\u0054\u0079\u0070\u0065")); !_eacb && (!_afag || _gfae.String() != "\u0046\u006f\u006e\u0074") {
				_dfec = append(_dfec, _dd("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_eacb = true
				if _bfbe() {
					return _dfec
				}
			}
			_fffb, _egdf := _g.NewPdfFontFromPdfObject(_acab)
			if _egdf != nil {
				continue
			}
			var _cga string
			if _dbac, _fede := _cb.GetName(_acab.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fede {
				_cga = _dbac.String()
			}
			if !_gbac {
				switch _cga {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_gbac = true
					_dfec = append(_dfec, _dd("\u0036.\u0033\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _bfbe() {
						return _dfec
					}
				}
			}
			if !_ccea {
				if _cga != "\u0054\u0079\u0070e\u0033" {
					_dgde, _febe := _cb.GetName(_acab.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_febe || _dgde.String() == "" {
						_dfec = append(_dfec, _dd("\u0036.\u0033\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_ccea = true
						if _bfbe() {
							return _dfec
						}
					}
				}
			}
			if _cga != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_bcac := _ae.IsStdFont(_ae.StdFontName(_fffb.BaseFont()))
			if _bcac {
				continue
			}
			_gfbgg, _gbbee := _cb.GetIntVal(_acab.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_gbbee && !_ffbbd {
				_dfec = append(_dfec, _dd("\u0036.\u0033\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_ffbbd = true
				if _bfbe() {
					return _dfec
				}
			}
			_cgda, _deggd := _cb.GetIntVal(_acab.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_deggd && !_bfdd {
				_dfec = append(_dfec, _dd("\u0036.\u0033\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_bfdd = true
				if _bfbe() {
					return _dfec
				}
			}
			if !_gdae {
				_aeeed, _cdcd := _cb.GetArray(_acab.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_cdcd || !_gbbee || !_deggd || _aeeed.Len() != _cgda-_gfbgg+1 {
					_dfec = append(_dfec, _dd("\u0036.\u0033\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_gdae = true
					if _bfbe() {
						return _dfec
					}
				}
			}
		}
	}
	return _dfec
}
func _aegc(_bedg *_g.CompliancePdfReader) (_gdgg []ViolatedRule) {
	_cecd, _abeeg := _addf(_bedg)
	if !_abeeg {
		return _gdgg
	}
	_gfabf, _abeeg := _cb.GetDict(_cecd.Get("\u0050\u0065\u0072m\u0073"))
	if !_abeeg {
		return _gdgg
	}
	_ebeeb := _gfabf.Keys()
	for _, _bgfb := range _ebeeb {
		if _bgfb.String() != "\u0055\u0052\u0033" && _bgfb.String() != "\u0044\u006f\u0063\u004d\u0044\u0050" {
			_gdgg = append(_gdgg, _dd("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "\u004e\u006f\u0020\u006b\u0065\u0079\u0073 \u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0055\u0052\u0033 \u0061n\u0064\u0020\u0044\u006f\u0063\u004dD\u0050\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u002e"))
		}
	}
	return _gdgg
}
func _bf() standardType { return standardType{_da: 2, _dag: "\u0041"} }
func _cbd(_fffe *_gc.Document, _cde int) {
	if _fffe.Version.Major == 0 {
		_fffe.Version.Major = 1
	}
	if _fffe.Version.Minor < _cde {
		_fffe.Version.Minor = _cde
	}
}
func _abaa(_ebee *_g.CompliancePdfReader) ViolatedRule {
	_fbeea, _ddgf := _ebee.PdfReader.GetTrailer()
	if _ddgf != nil {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _fbeea.Get("\u0049\u0044") == nil {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _fbeea.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _dfa
}

type standardType struct {
	_da  int
	_dag string
}

// Conformance gets the PDF/A conformance.
func (_bfaec *profile2) Conformance() string { return _bfaec._gaeg._dag }
func _edfa(_fdge *_gc.Document) {
	_gcee, _bde := _fdge.FindCatalog()
	if !_bde {
		return
	}
	_gabb, _bde := _gcee.GetMarkInfo()
	if !_bde {
		_gabb = _cb.MakeDict()
	}
	_aafb, _bde := _cb.GetBool(_gabb.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
	if !_bde || !bool(*_aafb) {
		_gabb.Set("\u004d\u0061\u0072\u006b\u0065\u0064", _cb.MakeBool(true))
		_gcee.SetMarkInfo(_gabb)
	}
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-1 standard.
func (_aad *profile1) ValidateStandard(r *_g.CompliancePdfReader) error {
	_fea := VerificationError{ConformanceLevel: _aad._efc._da, ConformanceVariant: _aad._efc._dag}
	if _eeb := _cbcbe(r); _eeb != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _eeb)
	}
	if _ccb := _ccgbe(r); _ccb != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _ccb)
	}
	if _abeg := _bcdc(r); _abeg != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _abeg)
	}
	if _bdbd := _dbbeb(r); _bdbd != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _bdbd)
	}
	if _aaae := _egee(r); _aaae != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _aaae)
	}
	if _ccdb := _cage(r); len(_ccdb) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _ccdb...)
	}
	if _ddea := _gecg(r); _ddea != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _ddea)
	}
	if _beb := _efae(r); len(_beb) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _beb...)
	}
	if _egb := _affe(r); len(_egb) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _egb...)
	}
	if _edc := _eccb(r); len(_edc) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _edc...)
	}
	if _defc := _egcgg(r); _defc != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _defc)
	}
	if _gfe := _dbdbe(r); len(_gfe) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _gfe...)
	}
	if _efeb := _ddda(r); len(_efeb) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _efeb...)
	}
	if _acbf := _cefg(r); _acbf != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _acbf)
	}
	if _agac := _dbbc(r, false); len(_agac) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _agac...)
	}
	if _bffg := _gdga(r); len(_bffg) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _bffg...)
	}
	if _gaba := _gdcf(r); _gaba != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _gaba)
	}
	if _ddeag := _adaf(r); _ddeag != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _ddeag)
	}
	if _abde := _dfgf(r); _abde != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _abde)
	}
	if _abdd := _cfcda(r); _abdd != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _abdd)
	}
	if _aaeg := _cdccf(r); _aaeg != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _aaeg)
	}
	if _gbbe := _ceac(r); len(_gbbe) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _gbbe...)
	}
	if _gbcfe := _cabe(r, _aad._efc); len(_gbcfe) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _gbcfe...)
	}
	if _ebfgb := _cfac(r); len(_ebfgb) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _ebfgb...)
	}
	if _ecbd := _bdgde(r); _ecbd != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _ecbd)
	}
	if _dcaa := _fcdf(r); _dcaa != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _dcaa)
	}
	if _aaef := _caac(r); len(_aaef) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _aaef...)
	}
	if _dfee := _efdgd(r); len(_dfee) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _dfee...)
	}
	if _ecfa := _gfaa(r); _ecfa != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _ecfa)
	}
	if _afab := _edbg(r); _afab != _dfa {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _afab)
	}
	if _affd := _aaac(r, _aad._efc, false); len(_affd) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _affd...)
	}
	if _aad._efc == _cbb() {
		if _dbgd := _bbaae(r); len(_dbgd) != 0 {
			_fea.ViolatedRules = append(_fea.ViolatedRules, _dbgd...)
		}
	}
	if _fgfb := _gcbg(r); len(_fgfb) != 0 {
		_fea.ViolatedRules = append(_fea.ViolatedRules, _fgfb...)
	}
	if len(_fea.ViolatedRules) > 0 {
		_be.Slice(_fea.ViolatedRules, func(_ccgb, _bfdfg int) bool {
			return _fea.ViolatedRules[_ccgb].RuleNo < _fea.ViolatedRules[_bfdfg].RuleNo
		})
		return _fea
	}
	return nil
}

// Conformance gets the PDF/A conformance.
func (_abge *profile1) Conformance() string { return _abge._efc._dag }
func _fbeae(_eafgae *_g.CompliancePdfReader) (_egcc []ViolatedRule) {
	var _dfecd, _bcgf, _fbgg, _facbd, _befdg, _cgfb, _gbae bool
	_fafca := func() bool { return _dfecd && _bcgf && _fbgg && _facbd && _befdg && _cgfb && _gbae }
	for _, _ffbf := range _eafgae.PageList {
		_cffea, _gfeac := _ffbf.GetAnnotations()
		if _gfeac != nil {
			_ge.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _gfeac)
			continue
		}
		for _, _edgf := range _cffea {
			if !_dfecd {
				switch _edgf.GetContext().(type) {
				case *_g.PdfAnnotationScreen, *_g.PdfAnnotation3D, *_g.PdfAnnotationSound, *_g.PdfAnnotationMovie, nil:
					_egcc = append(_egcc, _dd("\u0036.\u0033\u002e\u0031\u002d\u0031", "\u0041nn\u006f\u0074\u0061\u0074i\u006f\u006e t\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065r\u006d\u0069t\u0074\u0065\u0064\u002e\u0020\u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0033\u0044\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u0053\u0063\u0072\u0065\u0065\u006e\u0020\u0061n\u0064\u0020\u004d\u006f\u0076\u0069\u0065\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_dfecd = true
					if _fafca() {
						return _egcc
					}
				}
			}
			_fbgd, _fafab := _cb.GetDict(_edgf.GetContainingPdfObject())
			if !_fafab {
				continue
			}
			_, _dgebe := _edgf.GetContext().(*_g.PdfAnnotationPopup)
			if !_dgebe && !_bcgf {
				_, _eabc := _cb.GetIntVal(_fbgd.Get("\u0046"))
				if !_eabc {
					_egcc = append(_egcc, _dd("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0045\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072i\u0065\u0073\u0020\u0077\u0068\u006fs\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0069\u0073\u0020\u0050\u006f\u0070u\u0070\u002c\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065y."))
					_bcgf = true
					if _fafca() {
						return _egcc
					}
				}
			}
			if !_fbgg {
				_fccc, _cadg := _cb.GetIntVal(_fbgd.Get("\u0046"))
				if _cadg && !(_fccc&4 == 4 && _fccc&1 == 0 && _fccc&2 == 0 && _fccc&32 == 0 && _fccc&256 == 0) {
					_egcc = append(_egcc, _dd("\u0036.\u0033\u002e\u0032\u002d\u0032", "I\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073\u0068\u0061l\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020\u0049\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u002c\u0020\u0054\u006f\u0067\u0067\u006c\u0065\u004e\u006f\u0056\u0069\u0065\u0077\u002c\u0020\u0061\u006e\u0064 \u004eo\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020s\u0065\u0074\u0020t\u006f\u0020\u0030."))
					_fbgg = true
					if _fafca() {
						return _egcc
					}
				}
			}
			_, _eegge := _edgf.GetContext().(*_g.PdfAnnotationText)
			if _eegge && !_facbd {
				_efffg, _dbfbd := _cb.GetIntVal(_fbgd.Get("\u0046"))
				if _dbfbd && !(_efffg&8 == 8 && _efffg&16 == 16) {
					_egcc = append(_egcc, _dd("\u0036.\u0033\u002e\u0032\u002d\u0033", "\u0054\u0065\u0078\u0074\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069o\u006e\u0020\u0068\u0061\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006ca\u0067\u0073\u0020\u004e\u006f\u005a\u006f\u006f\u006d\u0020\u006f\u0072\u0020\u004e\u006f\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u0073\u0065t\u0020\u0074\u006f\u0020\u0030\u002e"))
					_facbd = true
					if _fafca() {
						return _egcc
					}
				}
			}
			if !_befdg {
				_abag, _bccf := _cb.GetDict(_fbgd.Get("\u0041\u0050"))
				if _bccf {
					_fgbe := _abag.Get("\u004e")
					if _fgbe == nil || len(_abag.Keys()) > 1 {
						_egcc = append(_egcc, _dd("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_befdg = true
						if _fafca() {
							return _egcc
						}
						continue
					}
					_, _fggbe := _edgf.GetContext().(*_g.PdfAnnotationWidget)
					if _fggbe {
						_dffc, _eggg := _cb.GetName(_fbgd.Get("\u0046\u0054"))
						if _eggg && *_dffc == "\u0042\u0074\u006e" {
							if _, _bbgc := _cb.GetDict(_fgbe); !_bbgc {
								_egcc = append(_egcc, _dd("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_befdg = true
								if _fafca() {
									return _egcc
								}
								continue
							}
						}
					}
					_, _ebeg := _cb.GetStream(_fgbe)
					if !_ebeg {
						_egcc = append(_egcc, _dd("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_befdg = true
						if _fafca() {
							return _egcc
						}
						continue
					}
				}
			}
			_dfbb, _ffgd := _edgf.GetContext().(*_g.PdfAnnotationWidget)
			if !_ffgd {
				continue
			}
			if !_cgfb {
				if _dfbb.A != nil {
					_egcc = append(_egcc, _dd("\u0036.\u0034\u002e\u0031\u002d\u0031", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_cgfb = true
					if _fafca() {
						return _egcc
					}
				}
			}
			if !_gbae {
				if _dfbb.AA != nil {
					_egcc = append(_egcc, _dd("\u0036.\u0034\u002e\u0031\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_gbae = true
					if _fafca() {
						return _egcc
					}
				}
			}
		}
	}
	return _egcc
}
func _gcgb(_ccde *_g.CompliancePdfReader) ViolatedRule {
	_fgbd := _ccde.ParserMetadata().HeaderCommentBytes()
	if _fgbd[0] > 127 && _fgbd[1] > 127 && _fgbd[2] > 127 && _fgbd[3] > 127 {
		return _dfa
	}
	return _dd("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}

// NewProfile2B creates a new Profile2B with the given options.
func NewProfile2B(options *Profile2Options) *Profile2B {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_eecgb(options)
	return &Profile2B{profile2{_afef: *options, _gaeg: _gce()}}
}

// Profile2A is the implementation of the PDF/A-2A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2A struct{ profile2 }

func _cfd(_gec []_cb.PdfObject) (*documentImages, error) {
	_egcg := _cb.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_aac := make(map[*_cb.PdfObjectStream]struct{})
	_ce := make(map[_cb.PdfObject]struct{})
	var (
		_gcf, _bgb, _fae bool
		_ebf             []*imageInfo
		_bb              error
	)
	for _, _cd := range _gec {
		_ba, _ca := _cb.GetStream(_cd)
		if !_ca {
			continue
		}
		if _, _ddd := _aac[_ba]; _ddd {
			continue
		}
		_aac[_ba] = struct{}{}
		_bgg := _ba.PdfObjectDictionary.Get(_egcg)
		_ggg, _ca := _cb.GetName(_bgg)
		if !_ca || string(*_ggg) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if _gff := _ba.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b"); _gff != nil {
			_ce[_gff] = struct{}{}
		}
		_gb := imageInfo{BitsPerComponent: 8, Stream: _ba}
		_gb.ColorSpace, _bb = _g.DetermineColorspaceNameFromPdfObject(_ba.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _bb != nil {
			return nil, _bb
		}
		if _dcg, _cbc := _cb.GetIntVal(_ba.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _cbc {
			_gb.BitsPerComponent = _dcg
		}
		if _ddf, _adf := _cb.GetIntVal(_ba.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _adf {
			_gb.Width = _ddf
		}
		if _eba, _adfe := _cb.GetIntVal(_ba.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _adfe {
			_gb.Height = _eba
		}
		switch _gb.ColorSpace {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_fae = true
			_gb.ColorComponents = 1
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_gcf = true
			_gb.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_bgb = true
			_gb.ColorComponents = 4
		default:
			_gb._dc = true
		}
		_ebf = append(_ebf, &_gb)
	}
	if len(_ce) > 0 {
		if len(_ce) == len(_ebf) {
			_ebf = nil
		} else {
			_gbc := make([]*imageInfo, len(_ebf)-len(_ce))
			var _bba int
			for _, _eec := range _ebf {
				if _, _eab := _ce[_eec.Stream]; _eab {
					continue
				}
				_gbc[_bba] = _eec
				_bba++
			}
			_ebf = _gbc
		}
	}
	return &documentImages{_egc: _gcf, _eb: _bgb, _cf: _fae, _fag: _ce, _geb: _ebf}, nil
}
func _cdfbg(_acffe *_cb.PdfObjectStream, _aaege map[*_cb.PdfObjectStream][]byte, _agcae map[*_cb.PdfObjectStream]*_ad.CMap) (*_ad.CMap, error) {
	_dcce, _fabd := _agcae[_acffe]
	if !_fabd {
		var _fbbd error
		_ccdac, _eefab := _aaege[_acffe]
		if !_eefab {
			_ccdac, _fbbd = _cb.DecodeStream(_acffe)
			if _fbbd != nil {
				_ge.Log.Debug("\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _fbbd)
				return nil, _fbbd
			}
			_aaege[_acffe] = _ccdac
		}
		_dcce, _fbbd = _ad.LoadCmapFromData(_ccdac, false)
		if _fbbd != nil {
			return nil, _fbbd
		}
		_agcae[_acffe] = _dcce
	}
	return _dcce, nil
}
func _affe(_dbebag *_g.CompliancePdfReader) (_ecd []ViolatedRule) {
	var _daaf, _fcac, _ddgd bool
	if _dbebag.ParserMetadata().HasNonConformantStream() {
		_ecd = []ViolatedRule{_dd("\u0036.\u0031\u002e\u0037\u002d\u0031", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _fcff := range _dbebag.GetObjectNums() {
		_gcef, _ := _dbebag.GetIndirectObjectByNumber(_fcff)
		if _gcef == nil {
			continue
		}
		_ebbe, _gcac := _cb.GetStream(_gcef)
		if !_gcac {
			continue
		}
		if !_daaf {
			_cbab := _ebbe.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _cbab == nil {
				_ecd = append(_ecd, _dd("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_daaf = true
			} else {
				_bbeb, _febae := _cb.GetIntVal(_cbab)
				if !_febae {
					_ecd = append(_ecd, _dd("\u0036.\u0031\u002e\u0037\u002d\u0032", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_daaf = true
				} else {
					if len(_ebbe.Stream) != _bbeb {
						_ecd = append(_ecd, _dd("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_daaf = true
					}
				}
			}
		}
		if !_fcac {
			if _ebbe.Get("\u0046") != nil {
				_fcac = true
				_ecd = append(_ecd, _dd("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _ebbe.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_fcac {
				_fcac = true
				_ecd = append(_ecd, _dd("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _ebbe.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_fcac {
				_fcac = true
				_ecd = append(_ecd, _dd("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_ddgd {
			_dee, _gafe := _cb.GetName(_cb.TraceToDirectObject(_ebbe.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_gafe {
				continue
			}
			if *_dee == _cb.StreamEncodingFilterNameLZW {
				_ddgd = true
				_ecd = append(_ecd, _dd("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0031", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _ecd
}
func _fada(_bfcg *_cb.PdfObjectDictionary, _cebca map[*_cb.PdfObjectStream][]byte, _cccd map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_eebeb = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0037\u002d\u0031"
		_cbff  = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_geab, _cdae := _cb.GetStream(_bfcg.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _cdae {
		_, _efdc := _cdfbg(_geab, _cebca, _cccd)
		if _efdc != nil {
			return _dd(_eebeb, _cbff)
		}
		return _dfa
	}
	_dbgc, _cdae := _cb.GetName(_bfcg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_cdae {
		return _dd(_eebeb, _cbff)
	}
	switch _dbgc.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _dfa
	}
	return _dd(_eebeb, _cbff)
}
func _dbaa(_agaec *_g.PdfFont, _adff *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_ddge = "\u0036.\u0033\u002e\u0037\u002d\u0031"
		_ecdg = "\u0041\u006cl \u006e\u006f\u006e\u002d\u0073\u0079\u006db\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0066o\u006e\u0074s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020e\u0069\u0074h\u0065\u0072\u0020\u004d\u0061\u0063\u0052\u006f\u006d\u0061\u006e\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041\u006e\u0073i\u0045n\u0063\u006f\u0064\u0069n\u0067\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072\u0020t\u0068\u0065 \u0045n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020t\u0068e\u0020\u0046o\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0072\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066\u006f\u0072 \u0074\u0068\u0065\u0020\u0042\u0061\u0073\u0065\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065 \u0064i\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0069s\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0074\u0068e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006be\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046\u006f\u006e\u0074 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u0049\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e, \u006eo\u0020n\u006f\u006e\u002d\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079p\u0065 \u0066\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0020\u0061\u0020\u0044\u0069\u0066\u0066e\u0072\u0065\u006e\u0063\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0075n\u006c\u0065s\u0073\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0074h\u0065\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u006e\u0061\u006d\u0065\u0073 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006ec\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0061\u0072\u0065\u0020\u006c\u0069\u0073\u0074\u0065\u0064 \u0069\u006e \u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065 G\u006c\u0079\u0070\u0068\u0020\u004c\u0069\u0073t\u0020\u0061\u006e\u0064\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066o\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072a\u006d\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073t\u0020\u0074\u0068\u0065\u0020\u004d\u0069\u0063\u0072o\u0073o\u0066\u0074\u0020\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0028\u0033\u002c\u0031 \u2013 P\u006c\u0061\u0074\u0066\u006f\u0072\u006d\u0020I\u0044\u003d\u0033\u002c\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067 I\u0044\u003d\u0031\u0029\u0020\u0065\u006e\u0063\u006f\u0064i\u006e\u0067 \u0069\u006e\u0020t\u0068\u0065\u0020'\u0063\u006d\u0061\u0070\u0027\u0020\u0074\u0061\u0062\u006c\u0065\u002e"
	)
	var _cege string
	if _dage, _acad := _cb.GetName(_adff.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _acad {
		_cege = _dage.String()
	}
	if _cege != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _dfa
	}
	_ffag := _agaec.FontDescriptor()
	_gffdf, _eggeb := _cb.GetIntVal(_ffag.Flags)
	if !_eggeb {
		_ge.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _dd(_ddge, _ecdg)
	}
	_bdbf := (uint32(_gffdf) >> 3) != 0
	if _bdbf {
		return _dfa
	}
	_ffdfb, _eggeb := _cb.GetName(_adff.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if !_eggeb {
		return _dd(_ddge, _ecdg)
	}
	switch _ffdfb.String() {
	case "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067":
		return _dfa
	default:
		return _dd(_ddge, _ecdg)
	}
}
func _aagg(_geaa *_g.CompliancePdfReader, _bccg standardType, _ccege bool) (_abbg []ViolatedRule) {
	_fegg, _bfdb := _addf(_geaa)
	if !_bfdb {
		return []ViolatedRule{_dd("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_abec := _fegg.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _abec == nil {
		return []ViolatedRule{_dd("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_bfdbe, _bfdb := _cb.GetStream(_abec)
	if !_bfdb {
		return []ViolatedRule{_dd("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_ggedf, _gbbc := _gga.LoadDocument(_bfdbe.Stream)
	if _gbbc != nil {
		return []ViolatedRule{_dd("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d4", "\u0041\u006c\u006c\u0020\u006de\u0074\u0061\u0064a\u0074\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020i\u006e \u0074\u0068\u0065\u0020\u0050\u0044\u0046 \u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065ci\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002e\u0020\u0041\u006c\u006c\u0020c\u006fn\u0074\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0061\u006c\u006c\u0020\u0058\u004d\u0050\u0020p\u0061\u0063\u006b\u0065\u0074\u0073 \u0073h\u0061\u006c\u006c \u0062\u0065\u0020\u0077\u0065\u006c\u006c\u002d\u0066o\u0072\u006de\u0064")}
	}
	_eaadf := _ggedf.GetGoXmpDocument()
	var _ebcdd []*_fe.Namespace
	for _, _cdef := range _eaadf.Namespaces() {
		switch _cdef.Name {
		case _gg.NsDc.Name, _df.NsPDF.Name, _fed.NsXmp.Name, _de.NsXmpRights.Name, _ea.Namespace.Name, _cc.Namespace.Name, _dea.NsXmpMM.Name, _cc.FieldNS.Name, _cc.SchemaNS.Name, _cc.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_ebcdd = append(_ebcdd, _cdef)
	}
	_gccfb := true
	_dbdfd, _gbbc := _ggedf.GetPdfaExtensionSchemas()
	if _gbbc == nil {
		for _, _agddb := range _ebcdd {
			var _fdeg bool
			for _fcfgg := range _dbdfd {
				if _agddb.URI == _dbdfd[_fcfgg].NamespaceURI {
					_fdeg = true
					break
				}
			}
			if !_fdeg {
				_gccfb = false
				break
			}
		}
	} else {
		_gccfb = false
	}
	if !_gccfb {
		_abbg = append(_abbg, _dd("\u0036.\u0036\u002e\u0032\u002e\u0033\u002d7", "\u0041\u006c\u006c\u0020\u0070\u0072\u006f\u0070e\u0072\u0074\u0069e\u0073\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072m\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0075s\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063he\u006da\u0073 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002c\u0020\u0049\u0053\u004f\u0020\u0031\u00390\u0030\u0035-\u0031\u0020\u006f\u0072\u0020\u0074h\u0069s\u0020\u0070\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u0020\u0031\u0039\u0030\u0030\u0035\u002c\u0020o\u0072\u0020\u0061\u006e\u0079\u0020e\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073c\u0068\u0065\u006das\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006fm\u0070\u006c\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0036\u002e\u0036\u002e\u0032.\u0033\u002e\u0032\u002e"))
	}
	_gbad, _bfdb := _ggedf.GetPdfAID()
	if !_bfdb {
		_abbg = append(_abbg, _dd("\u0036.\u0036\u002e\u0034\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _gbad.Part != _bccg._da {
			_abbg = append(_abbg, _dd("\u0036.\u0036\u002e\u0034\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _bccg._dag == "\u0041" && _gbad.Conformance != "\u0041" {
			_abbg = append(_abbg, _dd("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _bccg._dag == "\u0055" && (_gbad.Conformance != "\u0041" && _gbad.Conformance != "\u0055") {
			_abbg = append(_abbg, _dd("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _bccg._dag == "\u0042" && (_gbad.Conformance != "\u0041" && _gbad.Conformance != "\u0042" && _gbad.Conformance != "\u0055") {
			_abbg = append(_abbg, _dd("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		}
	}
	return _abbg
}

// String gets a string representation of the violated rule.
func (_adc ViolatedRule) String() string {
	return _c.Sprintf("\u0025\u0073\u003a\u0020\u0025\u0073", _adc.RuleNo, _adc.Detail)
}
func _ged(_gae *_gc.Document, _fdfb bool) error {
	_dfd, _cceb := _gae.GetPages()
	if !_cceb {
		return nil
	}
	for _, _adcb := range _dfd {
		_faga, _ffb := _adcb.GetContents()
		if !_ffb {
			continue
		}
		var _baaa *_g.PdfPageResources
		_gef, _ffb := _adcb.GetResources()
		if _ffb {
			_baaa, _ = _g.NewPdfPageResourcesFromDict(_gef)
		}
		for _dace, _baef := range _faga {
			_dbf, _agde := _baef.GetData()
			if _agde != nil {
				continue
			}
			_bca := _ee.NewContentStreamParser(string(_dbf))
			_gedc, _agde := _bca.Parse()
			if _agde != nil {
				continue
			}
			_ebg, _agde := _edb(_baaa, _gedc, _fdfb)
			if _agde != nil {
				return _agde
			}
			if _ebg == nil {
				continue
			}
			if _agde = (&_faga[_dace]).SetData(_ebg); _agde != nil {
				return _agde
			}
		}
	}
	return nil
}
func _bagb(_cebcb *_g.CompliancePdfReader) (_gbge []ViolatedRule) {
	_gcbcc := _cebcb.GetObjectNums()
	for _, _bafa := range _gcbcc {
		_eagg, _aabeg := _cebcb.GetIndirectObjectByNumber(_bafa)
		if _aabeg != nil {
			continue
		}
		_gcgdg, _adde := _cb.GetDict(_eagg)
		if !_adde {
			continue
		}
		_bfda, _adde := _cb.GetName(_gcgdg.Get("\u0054\u0079\u0070\u0065"))
		if !_adde {
			continue
		}
		if _bfda.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _gcgdg.Get("\u0045\u0046") != nil {
			if _gcgdg.Get("\u0046") == nil || _gcgdg.Get("\u0045\u0046") == nil {
				_gbge = append(_gbge, _dd("\u0036\u002e\u0038-\u0032", "\u0054h\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063t\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020t\u0068\u0065\u0020\u0046\u0020a\u006e\u0064\u0020\u0055\u0046\u0020\u006b\u0065\u0079\u0073\u002e"))
			}
			if _gcgdg.Get("\u0041\u0046\u0052\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070") == nil {
				_gbge = append(_gbge, _dd("\u0036\u002e\u0038-\u0033", "\u0049\u006e\u0020\u006f\u0072d\u0065\u0072\u0020\u0074\u006f\u0020\u0065\u006e\u0061\u0062\u006c\u0065\u0020i\u0064\u0065nt\u0069\u0066\u0069c\u0061\u0074\u0069o\u006e\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0074\u0068\u0065\u0020fi\u006ce\u0020\u0073\u0070\u0065\u0063\u0069f\u0069c\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020c\u006f\u006e\u0074e\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u0074\u002c\u0020\u0061\u0020\u006e\u0065\u0077\u0020(\u0072\u0065\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006be\u0079\u0020h\u0061\u0073\u0020\u0062e\u0065\u006e\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020a\u006e\u0064\u0020\u0069\u0074s \u0070\u0072e\u0073\u0065n\u0063\u0065\u0020\u0028\u0069\u006e\u0020\u0074\u0068e\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0029\u0020\u0069\u0073\u0020\u0072\u0065q\u0075\u0069\u0072e\u0064\u002e"))
			}
			break
		}
	}
	return _gbge
}

var _ Profile = (*Profile1B)(nil)

func _gefdb(_dcgd *_g.CompliancePdfReader) (_eggda []ViolatedRule) {
	_gfda := func(_dfebd *_cb.PdfObjectDictionary, _gcgce *[]string, _abegc *[]ViolatedRule) error {
		_dabfe := _dfebd.Get("\u004e\u0061\u006d\u0065")
		if _dabfe == nil || len(_dabfe.String()) == 0 {
			*_abegc = append(*_abegc, _dd("\u0036\u002e\u0039-\u0031", "\u0045\u0061\u0063\u0068\u0020o\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0063\u006fn\u0066\u0069\u0067\u0075r\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002e"))
		}
		for _, _ecbb := range *_gcgce {
			if _ecbb == _dabfe.String() {
				*_abegc = append(*_abegc, _dd("\u0036\u002e\u0039-\u0032", "\u0045\u0061\u0063\u0068\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0066\u0069\u0067\u0075\u0072a\u0074\u0069\u006fn\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020N\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002c w\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075ni\u0071\u0075\u0065 \u0061\u006d\u006f\u006e\u0067\u0073\u0074\u0020\u0061\u006c\u006c\u0020o\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006fn\u0074\u0065\u006e\u0074 \u0063\u006f\u006e\u0066\u0069\u0067u\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061\u0072\u0069\u0065\u0073\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0074\u0068e\u0020\u0050\u0044\u0046\u002fA\u002d\u0032\u0020\u0066\u0069l\u0065\u002e"))
			} else {
				*_gcgce = append(*_gcgce, _dabfe.String())
			}
		}
		if _dfebd.Get("\u0041\u0053") != nil {
			*_abegc = append(*_abegc, _dd("\u0036\u002e\u0039-\u0034", "Th\u0065\u0020\u0041\u0053\u0020\u006b\u0065y \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0070\u0070\u0065\u0061r\u0020\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0063\u006fn\u0066\u0069\u0067\u0075\u0072\u0061\u0074\u0069\u006fn\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
		}
		return nil
	}
	_bgef, _dacfd := _addf(_dcgd)
	if !_dacfd {
		return _eggda
	}
	_aggce, _dacfd := _cb.GetDict(_bgef.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_dacfd {
		return _eggda
	}
	var _agab []string
	_cbgag, _dacfd := _cb.GetDict(_aggce.Get("\u0044"))
	if _dacfd {
		_gfda(_cbgag, &_agab, &_eggda)
	}
	_cebcc, _dacfd := _cb.GetArray(_aggce.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _dacfd {
		for _cfgc := 0; _cfgc < _cebcc.Len(); _cfgc++ {
			_fgbda, _dbafa := _cb.GetDict(_cebcc.Get(_cfgc))
			if !_dbafa {
				continue
			}
			_gfda(_fgbda, &_agab, &_eggda)
		}
	}
	return _eggda
}

// Profile1Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile1Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A-1
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _bee.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func (_ccd *documentImages) hasOnlyDeviceCMYK() bool { return _ccd._eb && !_ccd._egc && !_ccd._cf }
func _ebe(_fga *_gc.Document) {
	if _fga.ID[0] != "" && _fga.ID[1] != "" {
		return
	}
	_fga.UseHashBasedID = true
}

// VerificationError is the PDF/A verification error structure, that contains all violated rules.
type VerificationError struct {

	// ViolatedRules are the rules that were violated during error verification.
	ViolatedRules []ViolatedRule

	// ConformanceLevel defines the standard on verification failed.
	ConformanceLevel int

	// ConformanceVariant is the standard variant used on verification.
	ConformanceVariant string
}

func _egcgg(_acge *_g.CompliancePdfReader) ViolatedRule {
	for _, _cbg := range _acge.PageList {
		_ggdg := _cbg.GetContentStreamObjs()
		for _, _bedcd := range _ggdg {
			_bedcd = _cb.TraceToDirectObject(_bedcd)
			var _aaec string
			switch _ddgb := _bedcd.(type) {
			case *_cb.PdfObjectString:
				_aaec = _ddgb.Str()
			case *_cb.PdfObjectStream:
				_eafc, _fbdg := _cb.GetName(_cb.TraceToDirectObject(_ddgb.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
				if _fbdg {
					if *_eafc == _cb.StreamEncodingFilterNameLZW {
						return _dd("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
					}
				}
				_faf, _ccbd := _cb.DecodeStream(_ddgb)
				if _ccbd != nil {
					_ge.Log.Debug("\u0045r\u0072\u003a\u0020\u0025\u0076", _ccbd)
					continue
				}
				_aaec = string(_faf)
			default:
				_ge.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063t\u003a\u0020\u0025\u0054", _bedcd)
				continue
			}
			_effdd := _ee.NewContentStreamParser(_aaec)
			_gfa, _adad := _effdd.Parse()
			if _adad != nil {
				_ge.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d:\u0020\u0025\u0076", _adad)
				continue
			}
			for _, _dacb := range *_gfa {
				if !(_dacb.Operand == "\u0042\u0049" && len(_dacb.Params) == 1) {
					continue
				}
				_facg, _bdbdd := _dacb.Params[0].(*_ee.ContentStreamInlineImage)
				if !_bdbdd {
					continue
				}
				_fgeda, _dabg := _facg.GetEncoder()
				if _dabg != nil {
					_ge.Log.Debug("\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006ee\u0020\u0069\u006d\u0061\u0067\u0065 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _dabg)
					continue
				}
				if _fgeda.GetFilterName() == _cb.StreamEncodingFilterNameLZW {
					return _dd("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
				}
			}
		}
	}
	return _dfa
}
func _fcdf(_gdec *_g.CompliancePdfReader) ViolatedRule {
	_begbb := map[*_cb.PdfObjectStream]struct{}{}
	for _, _gbfbc := range _gdec.PageList {
		if _gbfbc.Resources == nil && _gbfbc.Contents == nil {
			continue
		}
		if _cead := _gbfbc.GetPageDict(); _cead != nil {
			_eadbd, _dddf := _cb.GetDict(_cead.Get("\u0047\u0072\u006fu\u0070"))
			if _dddf {
				if _fcg := _eadbd.Get("\u0053"); _fcg != nil {
					_cade, _dggd := _cb.GetName(_fcg)
					if _dggd && _cade.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						return _dd("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
					}
				}
			}
		}
		if _gbfbc.Resources != nil {
			if _ffbbf, _egbe := _cb.GetDict(_gbfbc.Resources.XObject); _egbe {
				for _, _babd := range _ffbbf.Keys() {
					_eeaga, _bagg := _cb.GetStream(_ffbbf.Get(_babd))
					if !_bagg {
						continue
					}
					if _, _gbebd := _begbb[_eeaga]; _gbebd {
						continue
					}
					_adcd, _bagg := _cb.GetDict(_eeaga.Get("\u0047\u0072\u006fu\u0070"))
					if !_bagg {
						_begbb[_eeaga] = struct{}{}
						continue
					}
					_cgfdc := _adcd.Get("\u0053")
					if _cgfdc != nil {
						_egfb, _egfbb := _cb.GetName(_cgfdc)
						if _egfbb && _egfb.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _dd("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_begbb[_eeaga] = struct{}{}
					continue
				}
			}
		}
		if _gbfbc.Contents != nil {
			_acbbf, _cbbca := _gbfbc.GetContentStreams()
			if _cbbca != nil {
				continue
			}
			for _, _aeff := range _acbbf {
				_gfea, _daca := _ee.NewContentStreamParser(_aeff).Parse()
				if _daca != nil {
					continue
				}
				for _, _bgdf := range *_gfea {
					if len(_bgdf.Params) == 0 {
						continue
					}
					_gacd, _afcb := _cb.GetName(_bgdf.Params[0])
					if !_afcb {
						continue
					}
					_fadf, _afde := _gbfbc.Resources.GetXObjectByName(*_gacd)
					if _afde != _g.XObjectTypeForm {
						continue
					}
					if _, _daedc := _begbb[_fadf]; _daedc {
						continue
					}
					_fade, _afcb := _cb.GetDict(_fadf.Get("\u0047\u0072\u006fu\u0070"))
					if !_afcb {
						_begbb[_fadf] = struct{}{}
						continue
					}
					_daede := _fade.Get("\u0053")
					if _daede != nil {
						_adbe, _dgdb := _cb.GetName(_daede)
						if _dgdb && _adbe.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _dd("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_begbb[_fadf] = struct{}{}
				}
			}
		}
	}
	return _dfa
}
func _fgdd(_cagd *_g.CompliancePdfReader) (_cegf ViolatedRule) {
	_bgbf, _edef := _addf(_cagd)
	if !_edef {
		return _dfa
	}
	_fcgf, _edef := _cb.GetDict(_bgbf.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_edef {
		return _dfa
	}
	_dbfgf, _edef := _cb.GetArray(_fcgf.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_edef {
		return _dfa
	}
	for _acfd := 0; _acfd < _dbfgf.Len(); _acfd++ {
		_fegb, _afbf := _cb.GetDict(_dbfgf.Get(_acfd))
		if !_afbf {
			continue
		}
		if _fegb.Get("\u0041") != nil {
			return _dd("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
		if _fegb.Get("\u0041\u0041") != nil {
			return _dd("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
	}
	return _dfa
}
func _gfaa(_edcf *_g.CompliancePdfReader) (_fdba ViolatedRule) {
	_agbbe, _gbcag := _addf(_edcf)
	if !_gbcag {
		return _dfa
	}
	_bbafc, _gbcag := _cb.GetDict(_agbbe.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_gbcag {
		return _dfa
	}
	_egbef, _gbcag := _cb.GetArray(_bbafc.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_gbcag {
		return _dfa
	}
	for _ceaa := 0; _ceaa < _egbef.Len(); _ceaa++ {
		_cbfc, _aeaf := _cb.GetDict(_egbef.Get(_ceaa))
		if !_aeaf {
			continue
		}
		if _cbfc.Get("\u0041\u0041") != nil {
			return _dd("\u0036.\u0036\u002e\u0032\u002d\u0032", "\u0041\u0020F\u0069\u0065\u006cd\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079 s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061n\u0020A\u0041\u0020\u0065\u006e\u0074\u0072y f\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069on\u0061l\u002d\u0061\u0063\u0074i\u006fn\u0073 \u0064\u0069c\u0074\u0069on\u0061\u0072\u0079\u002e")
		}
	}
	return _dfa
}
func _cea(_befd *_gc.Document) error {
	_bag, _ecg := _befd.FindCatalog()
	if !_ecg {
		return _f.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_, _ecg = _cb.GetDict(_bag.Object.Get("\u0041\u0041"))
	if !_ecg {
		return nil
	}
	_bag.Object.Remove("\u0041\u0041")
	return nil
}

type colorspaceModification struct {
	_bc  _eg.ColorConverter
	_baa _g.PdfColorspace
}

func _gcaf(_dddb *_gc.Document) error {
	_fcb, _afb := _dddb.FindCatalog()
	if !_afb {
		return _f.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_cddf, _afb := _cb.GetDict(_fcb.Object.Get("\u0050\u0065\u0072m\u0073"))
	if _afb {
		_fgef := _cb.MakeDict()
		_beag := _cddf.Keys()
		for _, _dagc := range _beag {
			if _dagc.String() == "\u0055\u0052\u0033" || _dagc.String() == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				_fgef.Set(_dagc, _cddf.Get(_dagc))
			}
		}
		_fcb.Object.Set("\u0050\u0065\u0072m\u0073", _fgef)
	}
	return nil
}
func _ef(_fec []*_gc.Image, _gbg bool) error {
	_cg := _cb.PdfObjectName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
	if _gbg {
		_cg = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	for _, _caa := range _fec {
		if _caa.Colorspace == _cg {
			continue
		}
		_bed, _bec := _g.NewXObjectImageFromStream(_caa.Stream)
		if _bec != nil {
			return _bec
		}
		_egg, _bec := _bed.ToImage()
		if _bec != nil {
			return _bec
		}
		_dbb, _bec := _egg.ToGoImage()
		if _bec != nil {
			return _bec
		}
		var _dfb _g.PdfColorspace
		if _gbg {
			_dfb = _g.NewPdfColorspaceDeviceCMYK()
			_dbb, _bec = _eg.CMYKConverter.Convert(_dbb)
		} else {
			_dfb = _g.NewPdfColorspaceDeviceRGB()
			_dbb, _bec = _eg.NRGBAConverter.Convert(_dbb)
		}
		if _bec != nil {
			return _bec
		}
		_ga, _cgb := _dbb.(_eg.Image)
		if !_cgb {
			return _f.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
		}
		_fc := _ga.Base()
		_eea := &_g.Image{Width: int64(_fc.Width), Height: int64(_fc.Height), BitsPerComponent: int64(_fc.BitsPerComponent), ColorComponents: _fc.ColorComponents, Data: _fc.Data}
		_eea.SetDecode(_fc.Decode)
		_eea.SetAlpha(_fc.Alpha)
		if _bec = _bed.SetImage(_eea, _dfb); _bec != nil {
			return _bec
		}
		_bed.ToPdfObject()
		_caa.ColorComponents = _fc.ColorComponents
		_caa.Colorspace = _cg
	}
	return nil
}
func _beg(_efb *_gc.Document) error {
	_eaaa := func(_fcf *_cb.PdfObjectDictionary) error {
		if _cfg := _fcf.Get("\u0053\u004d\u0061s\u006b"); _cfg != nil {
			_fcf.Set("\u0053\u004d\u0061s\u006b", _cb.MakeName("\u004e\u006f\u006e\u0065"))
		}
		_fce := _fcf.Get("\u0043\u0041")
		if _fce != nil {
			_aaf, _caf := _cb.GetNumberAsFloat(_fce)
			if _caf != nil {
				_ge.Log.Debug("\u0045x\u0074\u0047S\u0074\u0061\u0074\u0065 \u006f\u0062\u006ae\u0063\u0074\u0020\u0043\u0041\u0020\u0076\u0061\u006cue\u0020\u0069\u0073 \u006e\u006ft\u0020\u0061\u0020\u0066\u006c\u006fa\u0074\u003a \u0025\u0076", _caf)
				_aaf = 0
			}
			if _aaf != 1.0 {
				_fcf.Set("\u0043\u0041", _cb.MakeFloat(1.0))
			}
		}
		_fce = _fcf.Get("\u0063\u0061")
		if _fce != nil {
			_adb, _fbf := _cb.GetNumberAsFloat(_fce)
			if _fbf != nil {
				_ge.Log.Debug("\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0027\u0063\u0061\u0027\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _fbf)
				_adb = 0
			}
			if _adb != 1.0 {
				_fcf.Set("\u0063\u0061", _cb.MakeFloat(1.0))
			}
		}
		_fgb := _fcf.Get("\u0042\u004d")
		if _fgb != nil {
			_ebb, _afe := _cb.GetName(_fgb)
			if !_afe {
				_ge.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_ebb = _cb.MakeName("")
			}
			_bggd := _ebb.String()
			switch _bggd {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
			default:
				_fcf.Set("\u0042\u004d", _cb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		_eag := _fcf.Get("\u0054\u0052")
		if _eag != nil {
			_ge.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_fcf.Remove("\u0054\u0052")
		}
		_cfc := _fcf.Get("\u0054\u0052\u0032")
		if _cfc != nil {
			_dde := _cfc.String()
			if _dde != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_ge.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_fcf.Set("\u0054\u0052\u0032", _cb.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		return nil
	}
	_ccc, _ggf := _efb.GetPages()
	if !_ggf {
		return nil
	}
	for _, _cgc := range _ccc {
		_bfa, _adbf := _cgc.GetResources()
		if !_adbf {
			continue
		}
		_gbca, _cfff := _cb.GetDict(_bfa.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_cfff {
			return nil
		}
		_dae := _gbca.Keys()
		for _, _eef := range _dae {
			_egf, _faef := _cb.GetDict(_gbca.Get(_eef))
			if !_faef {
				continue
			}
			_dfbe := _eaaa(_egf)
			if _dfbe != nil {
				continue
			}
		}
	}
	for _, _gaa := range _ccc {
		_fdb, _ebfg := _gaa.GetContents()
		if !_ebfg {
			return nil
		}
		for _, _bedc := range _fdb {
			_efe, _cba := _bedc.GetData()
			if _cba != nil {
				continue
			}
			_ab := _ee.NewContentStreamParser(string(_efe))
			_ead, _cba := _ab.Parse()
			if _cba != nil {
				continue
			}
			for _, _bbag := range *_ead {
				if len(_bbag.Params) == 0 {
					continue
				}
				_, _gcec := _cb.GetName(_bbag.Params[0])
				if !_gcec {
					continue
				}
				_add, _bfc := _gaa.GetResourcesXObject()
				if !_bfc {
					continue
				}
				for _, _gge := range _add.Keys() {
					_cbe, _efa := _cb.GetStream(_add.Get(_gge))
					if !_efa {
						continue
					}
					_gbgg, _efa := _cb.GetDict(_cbe.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_efa {
						continue
					}
					_eeg, _efa := _cb.GetDict(_gbgg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_efa {
						continue
					}
					for _, _ed := range _eeg.Keys() {
						_cge, _eecg := _cb.GetDict(_eeg.Get(_ed))
						if !_eecg {
							continue
						}
						_dfbeg := _eaaa(_cge)
						if _dfbeg != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}
func _cdcca(_bfbbg *_g.CompliancePdfReader) (_ggac []ViolatedRule) {
	var _ddcde, _gdeec, _ageb, _acbbeb, _ggdd, _cgdda, _bebd bool
	_bgge := map[*_cb.PdfObjectStream]struct{}{}
	for _, _egfce := range _bfbbg.GetObjectNums() {
		if _ddcde && _gdeec && _ggdd && _ageb && _acbbeb && _cgdda && _bebd {
			return _ggac
		}
		_gdgab, _bgfd := _bfbbg.GetIndirectObjectByNumber(_egfce)
		if _bgfd != nil {
			continue
		}
		_baabd, _cedc := _cb.GetStream(_gdgab)
		if !_cedc {
			continue
		}
		if _, _cedc = _bgge[_baabd]; _cedc {
			continue
		}
		_bgge[_baabd] = struct{}{}
		_dfegf, _cedc := _cb.GetName(_baabd.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_cedc {
			continue
		}
		if !_acbbeb {
			if _baabd.Get("\u0052\u0065\u0066") != nil {
				_ggac = append(_ggac, _dd("\u0036.\u0032\u002e\u0039\u002d\u0032", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_acbbeb = true
			}
		}
		if _dfegf.String() == "\u0050\u0053" {
			if !_cgdda {
				_ggac = append(_ggac, _dd("\u0036.\u0032\u002e\u0039\u002d\u0033", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_cgdda = true
				continue
			}
		}
		if _dfegf.String() == "\u0046\u006f\u0072\u006d" {
			if _gdeec && _ageb && _acbbeb {
				continue
			}
			if !_gdeec && _baabd.Get("\u004f\u0050\u0049") != nil {
				_ggac = append(_ggac, _dd("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d \u0058\u004f\u0062j\u0065\u0063\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
				_gdeec = true
			}
			if !_ageb {
				if _baabd.Get("\u0050\u0053") != nil {
					_ageb = true
				}
				if _ccefe := _baabd.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _ccefe != nil && !_ageb {
					if _ebca, _eaeff := _cb.GetName(_ccefe); _eaeff && *_ebca == "\u0050\u0053" {
						_ageb = true
					}
				}
				if _ageb {
					_ggac = append(_ggac, _dd("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065y \u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006cu\u0065 o\u0066 \u0050\u0053\u0020\u0061\u006e\u0064\u0020t\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _dfegf.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_ddcde && _baabd.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_ggac = append(_ggac, _dd("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_ddcde = true
		}
		if !_bebd && _baabd.Get("\u004f\u0050\u0049") != nil {
			_ggac = append(_ggac, _dd("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
			_bebd = true
		}
		if !_ggdd && _baabd.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_bcdd, _faca := _cb.GetBool(_baabd.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _faca && bool(*_bcdd) {
				continue
			}
			_ggac = append(_ggac, _dd("\u0036.\u0032\u002e\u0038\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_ggdd = true
		}
	}
	return _ggac
}
func _gecg(_dbgdf *_g.CompliancePdfReader) ViolatedRule { return _dfa }
func _ggcg(_ddee *_gc.Document, _feda bool) error {
	_defb, _egae := _ddee.GetPages()
	if !_egae {
		return nil
	}
	for _, _aeda := range _defb {
		_fbg, _ceed := _cb.GetArray(_aeda.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_ceed {
			continue
		}
		for _, _fbbg := range _fbg.Elements() {
			_fgc, _bfb := _cb.GetDict(_fbbg)
			if !_bfb {
				continue
			}
			_ebad := _fgc.Get("\u0043")
			if _ebad == nil {
				continue
			}
			_dba, _bfb := _cb.GetArray(_ebad)
			if !_bfb {
				continue
			}
			_cdd, _abfb := _dba.GetAsFloat64Slice()
			if _abfb != nil {
				return _abfb
			}
			switch _dba.Len() {
			case 0, 1:
				if _feda {
					_fgc.Set("\u0043", _cb.MakeArrayFromIntegers([]int{1, 1, 1, 1}))
				} else {
					_fgc.Set("\u0043", _cb.MakeArrayFromIntegers([]int{1, 1, 1}))
				}
			case 3:
				if _feda {
					_gaac, _acdg, _aggc, _cbdc := _d.RGBToCMYK(uint8(_cdd[0]*255), uint8(_cdd[1]*255), uint8(_cdd[2]*255))
					_fgc.Set("\u0043", _cb.MakeArrayFromFloats([]float64{float64(_gaac) / 255, float64(_acdg) / 255, float64(_aggc) / 255, float64(_cbdc) / 255}))
				}
			case 4:
				if !_feda {
					_bgc, _ecb, _ebba := _d.CMYKToRGB(uint8(_cdd[0]*255), uint8(_cdd[1]*255), uint8(_cdd[2]*255), uint8(_cdd[3]*255))
					_fgc.Set("\u0043", _cb.MakeArrayFromFloats([]float64{float64(_bgc) / 255, float64(_ecb) / 255, float64(_ebba) / 255}))
				}
			}
		}
	}
	return nil
}
func _ddbff(_ecgga *_g.CompliancePdfReader) (_dcag ViolatedRule) {
	_ggdc, _ceeb := _addf(_ecgga)
	if !_ceeb {
		return _dfa
	}
	if _ggdc.Get("\u0041\u0041") != nil {
		return _dd("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0054h\u0065\u0020\u0064\u006fc\u0075m\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020a\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079 \u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079\u002e")
	}
	return _dfa
}
func _bggfg(_fdcc *_g.PdfFont, _agfe *_cb.PdfObjectDictionary, _cggb bool) ViolatedRule {
	const (
		_fggb  = "\u0036.\u0033\u002e\u0034\u002d\u0031"
		_eeaag = "\u0054\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u0020s\u0068\u0061\u006cl\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006ce\u002c\u0020a\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0035\u002e\u0038\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0077h\u0065\u006e\u0020\u0074\u0068\u0065 \u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0065\u0078\u0063\u006cu\u0073i\u0076\u0065\u006c\u0079\u0020\u0077\u0069t\u0068\u0020\u0074\u0065\u0078\u0074\u0020\u0072e\u006ed\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020\u0033\u002e"
	)
	if _cggb {
		return _dfa
	}
	_gggf := _fdcc.FontDescriptor()
	var _gfef string
	if _daad, _cbefb := _cb.GetName(_agfe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cbefb {
		_gfef = _daad.String()
	}
	switch _gfef {
	case "\u0054\u0079\u0070e\u0031":
		if _gggf.FontFile == nil {
			return _dd(_fggb, _eeaag)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _gggf.FontFile2 == nil {
			return _dd(_fggb, _eeaag)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _gggf.FontFile3 == nil {
			return _dd(_fggb, _eeaag)
		}
	}
	return _dfa
}
func _cefg(_eded *_g.CompliancePdfReader) ViolatedRule {
	_cgdd, _agfd := _eded.GetTrailer()
	if _agfd != nil {
		_ge.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _agfd)
		return _dfa
	}
	_bbcg, _dgaf := _cgdd.Get("\u0052\u006f\u006f\u0074").(*_cb.PdfObjectReference)
	if !_dgaf {
		_ge.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return _dfa
	}
	_abbd, _dgaf := _cb.GetDict(_cb.ResolveReference(_bbcg))
	if !_dgaf {
		_ge.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return _dfa
	}
	if _abbd.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073") != nil {
		return _dd("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063u\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u004f\u0043\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	}
	return _dfa
}
func _daebe(_gcgfe *_cb.PdfObjectDictionary, _gecge map[*_cb.PdfObjectStream][]byte, _dcda map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_dffb = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0033"
		_dbfg = "\u0041\u006c\u006c \u0043\u004d\u0061\u0070s\u0020\u0075\u0073ed\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074 th\u006f\u0073\u0065\u0020\u006ci\u0073\u0074\u0065\u0064\u0020i\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u00320\u00300\u002d1\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0035\u002e\u0032\u002c\u0020\u0054\u0061\u0062\u006c\u0065 \u0031\u00318,\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e \u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064e\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u00300\u0030-\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0035\u002e"
	)
	var _dacab string
	if _gdfb, _gbcce := _cb.GetName(_gcgfe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _gbcce {
		_dacab = _gdfb.String()
	}
	if _dacab != "\u0054\u0079\u0070e\u0030" {
		return _dfa
	}
	_bggc := _gcgfe.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _ffgc, _abfbg := _cb.GetName(_bggc); _abfbg {
		switch _ffgc.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _dfa
		default:
			return _dd(_dffb, _dbfg)
		}
	}
	_dfdee, _cfea := _cb.GetStream(_bggc)
	if !_cfea {
		return _dd(_dffb, _dbfg)
	}
	_, _dgee := _cdfbg(_dfdee, _gecge, _dcda)
	if _dgee != nil {
		return _dd(_dffb, _dbfg)
	}
	return _dfa
}
func _cbdg(_bbff *_gc.Document, _gee standardType, _afac *_gc.OutputIntents) error {
	var (
		_efef *_g.PdfOutputIntent
		_dafc error
	)
	if _bbff.Version.Minor <= 7 {
		_efef, _dafc = _gf.NewSRGBv2OutputIntent(_gee.outputIntentSubtype())
	} else {
		_efef, _dafc = _gf.NewSRGBv4OutputIntent(_gee.outputIntentSubtype())
	}
	if _dafc != nil {
		return _dafc
	}
	if _dafc = _afac.Add(_efef.ToPdfObject()); _dafc != nil {
		return _dafc
	}
	return nil
}
func _bffa(_gcge *_cb.PdfObjectDictionary, _aea map[*_cb.PdfObjectStream][]byte, _gbag map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_ccfa = "\u0046\u006f\u0072 \u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029\u0020\u0066\u006f\u006et \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0064 \u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043I\u0044\u0053y\u0073\u0074\u0065\u006d\u0049nf\u006f\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006f\u0066\u0020i\u0074\u0073\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0061\u006e\u0064 \u0043\u004d\u0061\u0070 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0063\u006f\u006d\u0070\u0061\u0074i\u0062\u006c\u0065\u002e\u0020\u0049\u006e\u0020o\u0074\u0068\u0065\u0072\u0020\u0077\u006f\u0072\u0064\u0073\u002c\u0020\u0074\u0068\u0065\u0020R\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004fr\u0064\u0065\u0072\u0069\u006e\u0067 \u0073\u0074\u0072i\u006e\u0067\u0073\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0066\u006f\u0072 \u0074\u0068\u0061\u0074\u0020\u0066o\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u002c\u0020u\u006el\u0065ss \u0074\u0068\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073 \u0049\u0064\u0065\u006e\u0074\u0069t\u0079\u002d\u0048\u0020o\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074y\u002dV\u002e"
		_abgd = "\u0036.\u0033\u002e\u0033\u002d\u0031"
	)
	var _bcdg string
	if _dcb, _fbad := _cb.GetName(_gcge.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fbad {
		_bcdg = _dcb.String()
	}
	if _bcdg != "\u0054\u0079\u0070e\u0030" {
		return _dfa
	}
	_bcgg := _gcge.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _agca, _fee := _cb.GetName(_bcgg); _fee {
		switch _agca.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _dfa
		}
		_bfad, _ffbgf := _ad.LoadPredefinedCMap(_agca.String())
		if _ffbgf != nil {
			return _dd(_abgd, _ccfa)
		}
		_fedef := _bfad.CIDSystemInfo()
		if _fedef.Ordering != _fedef.Registry {
			return _dd(_abgd, _ccfa)
		}
		return _dfa
	}
	_eead, _gbaa := _cb.GetStream(_bcgg)
	if !_gbaa {
		return _dd(_abgd, _ccfa)
	}
	_bcdcb, _dgbf := _cdfbg(_eead, _aea, _gbag)
	if _dgbf != nil {
		return _dd(_abgd, _ccfa)
	}
	_gbacc := _bcdcb.CIDSystemInfo()
	if _gbacc.Ordering != _gbacc.Registry {
		return _dd(_abgd, _ccfa)
	}
	return _dfa
}
func _bacg(_eabec *_g.CompliancePdfReader) ViolatedRule { return _dfa }
func _cfac(_adeaf *_g.CompliancePdfReader) (_ddede []ViolatedRule) {
	var _bbaf, _deaa, _fgge, _gac, _babgd, _ccec bool
	_babf := func() bool { return _bbaf && _deaa && _fgge && _gac && _babgd && _ccec }
	_cegg := func(_gfbb *_cb.PdfObjectDictionary) bool {
		if !_bbaf && _gfbb.Get("\u0054\u0052") != nil {
			_bbaf = true
			_ddede = append(_ddede, _dd("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _bfbf := _gfbb.Get("\u0054\u0052\u0032"); !_deaa && _bfbf != nil {
			_cfga, _becec := _cb.GetName(_bfbf)
			if !_becec || (_becec && *_cfga != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_deaa = true
				_ddede = append(_ddede, _dd("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _babf() {
					return true
				}
			}
		}
		if _eage := _gfbb.Get("\u0053\u004d\u0061s\u006b"); !_fgge && _eage != nil {
			_aedg, _eace := _cb.GetName(_eage)
			if !_eace || (_eace && *_aedg != "\u004e\u006f\u006e\u0065") {
				_fgge = true
				_ddede = append(_ddede, _dd("\u0036\u002e\u0034-\u0031", "\u0049\u0066\u0020\u0061\u006e \u0053\u004d\u0061\u0073\u006b\u0020\u006be\u0079\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0069\u0074s\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u004e\u006f\u006ee\u002e"))
				if _babf() {
					return true
				}
			}
		}
		if _bfca := _gfbb.Get("\u0043\u0041"); !_babgd && _bfca != nil {
			_dgcg, _efdf := _cb.GetNumberAsFloat(_bfca)
			if _efdf == nil && _dgcg != 1.0 {
				_babgd = true
				_ddede = append(_ddede, _dd("\u0036\u002e\u0034-\u0035", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0043\u0041 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _babf() {
					return true
				}
			}
		}
		if _bfaa := _gfbb.Get("\u0063\u0061"); !_ccec && _bfaa != nil {
			_ccbc, _eaca := _cb.GetNumberAsFloat(_bfaa)
			if _eaca == nil && _ccbc != 1.0 {
				_ccec = true
				_ddede = append(_ddede, _dd("\u0036\u002e\u0034-\u0036", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0063\u0061 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _babf() {
					return true
				}
			}
		}
		if _beac := _gfbb.Get("\u0042\u004d"); !_gac && _beac != nil {
			_febd, _badba := _cb.GetName(_beac)
			if _badba {
				switch _febd.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
				default:
					_gac = true
					_ddede = append(_ddede, _dd("\u0036\u002e\u0034-\u0034", "T\u0068\u0065\u0020\u0066\u006f\u006cl\u006f\u0077\u0069\u006e\u0067 \u006b\u0065y\u0073\u002c\u0020\u0069\u0066 \u0070res\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047S\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065 \u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0077n\u003a\u0020\u0042\u004d\u0020\u002d\u0020\u004e\u006f\u0072m\u0061\u006c\u0020\u006f\u0072\u0020\u0043\u006f\u006d\u0070\u0061t\u0069\u0062\u006c\u0065\u002e"))
					if _babf() {
						return true
					}
				}
			}
		}
		return false
	}
	for _, _gebac := range _adeaf.PageList {
		_bacc := _gebac.Resources
		if _bacc == nil {
			continue
		}
		if _bacc.ExtGState == nil {
			continue
		}
		_efcf, _abeef := _cb.GetDict(_bacc.ExtGState)
		if !_abeef {
			continue
		}
		_gcce := _efcf.Keys()
		for _, _badee := range _gcce {
			_aefc, _ggge := _cb.GetDict(_efcf.Get(_badee))
			if !_ggge {
				continue
			}
			if _cegg(_aefc) {
				return _ddede
			}
		}
	}
	for _, _fcae := range _adeaf.PageList {
		_aaffb := _fcae.Resources
		if _aaffb == nil {
			continue
		}
		_defd, _gcbf := _cb.GetDict(_aaffb.XObject)
		if !_gcbf {
			continue
		}
		for _, _dcgbc := range _defd.Keys() {
			_cddc, _fabb := _cb.GetStream(_defd.Get(_dcgbc))
			if !_fabb {
				continue
			}
			_ceaec, _fabb := _cb.GetDict(_cddc.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_fabb {
				continue
			}
			_bbdag, _fabb := _cb.GetDict(_ceaec.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_fabb {
				continue
			}
			for _, _ccgfb := range _bbdag.Keys() {
				_agcf, _egfab := _cb.GetDict(_bbdag.Get(_ccgfb))
				if !_egfab {
					continue
				}
				if _cegg(_agcf) {
					return _ddede
				}
			}
		}
	}
	return _ddede
}
func _afaf(_baff *_g.PdfInfo, _eagd *_gga.Document) bool {
	_fgcc, _acbc := _eagd.GetPdfInfo()
	if !_acbc {
		return false
	}
	if _fgcc.InfoDict == nil {
		return false
	}
	_edda, _dcea := _g.NewPdfInfoFromObject(_fgcc.InfoDict)
	if _dcea != nil {
		return false
	}
	if _baff.Creator != nil {
		if _edda.Creator == nil || _edda.Creator.String() != _baff.Creator.String() {
			return false
		}
	}
	if _baff.CreationDate != nil {
		if _edda.CreationDate == nil || !_edda.CreationDate.ToGoTime().Equal(_baff.CreationDate.ToGoTime()) {
			return false
		}
	}
	if _baff.ModifiedDate != nil {
		if _edda.ModifiedDate == nil || !_edda.ModifiedDate.ToGoTime().Equal(_baff.ModifiedDate.ToGoTime()) {
			return false
		}
	}
	if _baff.Producer != nil {
		if _edda.Producer == nil || _edda.Producer.String() != _baff.Producer.String() {
			return false
		}
	}
	if _baff.Keywords != nil {
		if _edda.Keywords == nil || _edda.Keywords.String() != _baff.Keywords.String() {
			return false
		}
	}
	if _baff.Trapped != nil {
		if _edda.Trapped == nil {
			return false
		}
		switch _baff.Trapped.String() {
		case "\u0054\u0072\u0075\u0065":
			if _edda.Trapped.String() != "\u0054\u0072\u0075\u0065" {
				return false
			}
		case "\u0046\u0061\u006cs\u0065":
			if _edda.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		default:
			if _edda.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		}
	}
	if _baff.Title != nil {
		if _edda.Title == nil || _edda.Title.String() != _baff.Title.String() {
			return false
		}
	}
	if _baff.Subject != nil {
		if _edda.Subject == nil || _edda.Subject.String() != _baff.Subject.String() {
			return false
		}
	}
	return true
}
func _fcbf(_bbebc *_g.CompliancePdfReader) ViolatedRule {
	if _bbebc.ParserMetadata().HasDataAfterEOF() {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _dfa
}
func _badb(_aabc *_gc.Document, _bbg int) error {
	for _, _ada := range _aabc.Objects {
		_ece, _ffa := _cb.GetDict(_ada)
		if !_ffa {
			continue
		}
		_ceb := _ece.Get("\u0054\u0079\u0070\u0065")
		if _ceb == nil {
			continue
		}
		if _cab, _abg := _cb.GetName(_ceb); _abg && _cab.String() != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_ded, _abc := _cb.GetName(_ece.Get("\u0053"))
		if !_abc {
			continue
		}
		switch _g.PdfActionType(*_ded) {
		case _g.ActionTypeLaunch, _g.ActionTypeSound, _g.ActionTypeMovie, _g.ActionTypeResetForm, _g.ActionTypeImportData, _g.ActionTypeJavaScript:
			_ece.Remove("\u0053")
		case _g.ActionTypeHide, _g.ActionTypeSetOCGState, _g.ActionTypeRendition, _g.ActionTypeTrans, _g.ActionTypeGoTo3DView:
			if _bbg == 2 {
				_ece.Remove("\u0053")
			}
		case _g.ActionTypeNamed:
			_dcgb, _cfdb := _cb.GetName(_ece.Get("\u004e"))
			if !_cfdb {
				continue
			}
			switch *_dcgb {
			case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
			default:
				_ece.Remove("\u004e")
			}
		}
	}
	return nil
}
func _ffce(_dcae *_cb.PdfObjectDictionary, _defbb map[*_cb.PdfObjectStream][]byte, _cbcad map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_dgcgc = "\u0046\u006f\u0072\u0020\u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070o\u0073\u0069\u0074e\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029 \u0066\u006fn\u0074\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f \u0065\u006e\u0074\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074\u0073 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0066\u006fl\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0072\u0065l\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070. \u0049\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0054\u0079\u0070\u0065\u0020\u0030 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0069\u0073\u0020I\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056\u002c\u0020\u0061\u006e\u0079\u0020v\u0061\u006c\u0075\u0065\u0073\u0020\u006f\u0066\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u002c\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u002c\u0020\u0061\u006e\u0064\u0020\u0053up\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069n\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0065\u006e\u0074r\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044F\u006f\u006e\u0074\u002e\u0020\u004f\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u0072\u0072\u0065\u0073\u0070\u006f\u006e\u0064\u0069\u006e\u0067\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0069\u006e\u0020\u0062\u006f\u0074h\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006cl\u0020\u0062\u0065\u0020i\u0064en\u0074\u0069\u0063\u0061\u006c\u002c \u0061n\u0064\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0070\u0070l\u0065\u006d\u0065\u006e\u0074 \u006b\u0065\u0079\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0043I\u0044S\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0067re\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f t\u0068\u0065\u0020\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043M\u0061p\u002e"
		_efdfc = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0031"
	)
	var _cabac string
	if _edca, _becd := _cb.GetName(_dcae.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _becd {
		_cabac = _edca.String()
	}
	if _cabac != "\u0054\u0079\u0070e\u0030" {
		return _dfa
	}
	_adeab := _dcae.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _efagb, _bggeg := _cb.GetName(_adeab); _bggeg {
		switch _efagb.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _dfa
		}
		_ebfcg, _eadga := _ad.LoadPredefinedCMap(_efagb.String())
		if _eadga != nil {
			return _dd(_efdfc, _dgcgc)
		}
		_bdgdbf := _ebfcg.CIDSystemInfo()
		if _bdgdbf.Ordering != _bdgdbf.Registry {
			return _dd(_efdfc, _dgcgc)
		}
		return _dfa
	}
	_deada, _eacbc := _cb.GetStream(_adeab)
	if !_eacbc {
		return _dd(_efdfc, _dgcgc)
	}
	_egddb, _babcd := _cdfbg(_deada, _defbb, _cbcad)
	if _babcd != nil {
		return _dd(_efdfc, _dgcgc)
	}
	_cgdgb := _egddb.CIDSystemInfo()
	if _cgdgb.Ordering != _cgdgb.Registry {
		return _dd(_efdfc, _dgcgc)
	}
	return _dfa
}

// Part gets the PDF/A version level.
func (_cbcb *profile2) Part() int { return _cbcb._gaeg._da }
func _edee(_faed *_gc.Document) error {
	_feff, _fcea := _faed.FindCatalog()
	if !_fcea {
		return _f.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_cdeda, _fcea := _cb.GetDict(_feff.Object.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_fcea {
		return nil
	}
	_dadc, _fcea := _cb.GetDict(_cdeda.Get("\u0044"))
	if _fcea {
		if _dadc.Get("\u0041\u0053") != nil {
			_dadc.Remove("\u0041\u0053")
		}
	}
	_edgg, _fcea := _cb.GetArray(_cdeda.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _fcea {
		for _cegc := 0; _cegc < _edgg.Len(); _cegc++ {
			_aebd, _gecf := _cb.GetDict(_edgg.Get(_cegc))
			if !_gecf {
				continue
			}
			if _aebd.Get("\u0041\u0053") != nil {
				_aebd.Remove("\u0041\u0053")
			}
		}
	}
	return nil
}
func _eecgb(_ffec *Profile2Options) {
	if _ffec.Now == nil {
		_ffec.Now = _bee.Now
	}
}
func _becgc(_dfab *_g.CompliancePdfReader) (_fffaf []ViolatedRule) {
	var _accgd, _dfac, _ebgc, _cfage bool
	_gcfe := func() bool { return _accgd && _dfac && _ebgc && _cfage }
	_eadd, _gafb := _gdde(_dfab)
	var _affg _gf.ProfileHeader
	if _gafb {
		_affg, _ = _gf.ParseHeader(_eadd.DestOutputProfile)
	}
	_cdfd := map[_cb.PdfObject]struct{}{}
	var _efgb func(_gdce _g.PdfColorspace) bool
	_efgb = func(_dfde _g.PdfColorspace) bool {
		switch _fgcfe := _dfde.(type) {
		case *_g.PdfColorspaceDeviceGray:
			if !_accgd {
				if !_gafb {
					_fffaf = append(_fffaf, _dd("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d4", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064 \u0069\u0066\u0020\u0061\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065p\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0047\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0075r \u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u0077\u0068\u0065n \u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072a\u0079\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u002c o\u0072\u0020\u0069\u0066\u0020\u0061\u0020\u0050\u0044\u0046\u002fA\u0020\u004f\u0075tp\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002e"))
					_accgd = true
					if _gcfe() {
						return true
					}
				}
			}
		case *_g.PdfColorspaceDeviceRGB:
			if !_dfac {
				if !_gafb || _affg.ColorSpace != _gf.ColorSpaceRGB {
					_fffaf = append(_fffaf, _dd("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d2", "\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042\u0020\u0073\u0068\u0061\u006cl\u0020\u006f\u006e\u006c\u0079\u0020\u0062e\u0020\u0075\u0073\u0065\u0064\u0020\u0069f\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063e\u0020\u0069n\u0064\u0065\u0070e\u006e\u0064\u0065\u006et \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0052\u0047\u0042\u0020\u0063\u006fl\u006f\u0075r\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020b\u0065\u0065\u006e\u0020s\u0065\u0074 \u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020c\u006flou\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020i\u0073\u0020\u0075\u0073\u0065\u0064\u002c\u0020\u006f\u0072\u0020if\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044F\u002f\u0041\u0020\u004fut\u0070\u0075\u0074\u0049\u006e\u0074\u0065n\u0074\u0020t\u0068\u0061t\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u006e\u0020\u0052\u0047\u0042\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061\u0074io\u006e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u002e"))
					_dfac = true
					if _gcfe() {
						return true
					}
				}
			}
		case *_g.PdfColorspaceDeviceCMYK:
			if !_ebgc {
				if !_gafb || _affg.ColorSpace != _gf.ColorSpaceCMYK {
					_fffaf = append(_fffaf, _dd("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d3", "\u0044e\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020\u0073hal\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u0066\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065\u0070\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044ef\u0061\u0075\u006c\u0074\u0043\u004d\u0059K\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061s\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u006fr \u0069\u0066\u0020\u0061\u0020\u0044e\u0076\u0069\u0063\u0065\u004e\u002d\u0062\u0061\u0073\u0065\u0064\u0020\u0044\u0065f\u0061\u0075\u006c\u0074\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006f\u0075r\u0020\u0073\u0070\u0061\u0063e\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e\u0020\u0073\u0065\u0074\u0020\u0077\u0068\u0065\u006e\u0020\u0074h\u0065\u0020\u0044\u0065\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006fu\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u0072\u0020t\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u0020\u0043\u004d\u0059\u004b\u0020d\u0065\u0073\u0074\u0069\u006e\u0061t\u0069\u006f\u006e\u0020\u0070r\u006f\u0066\u0069\u006c\u0065\u002e"))
					_ebgc = true
					if _gcfe() {
						return true
					}
				}
			}
		case *_g.PdfColorspaceICCBased:
			if !_cfage {
				_aefag, _ggeg := _gf.ParseHeader(_fgcfe.Data)
				if _ggeg != nil {
					_ge.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ggeg)
					_fffaf = append(_fffaf, func() ViolatedRule {
						return _dd("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e")
					}())
					_cfage = true
					if _gcfe() {
						return true
					}
				}
				if !_cfage {
					var _egfaf, _bedce bool
					switch _aefag.DeviceClass {
					case _gf.DeviceClassPRTR, _gf.DeviceClassMNTR, _gf.DeviceClassSCNR, _gf.DeviceClassSPAC:
					default:
						_egfaf = true
					}
					switch _aefag.ColorSpace {
					case _gf.ColorSpaceRGB, _gf.ColorSpaceCMYK, _gf.ColorSpaceGRAY, _gf.ColorSpaceLAB:
					default:
						_bedce = true
					}
					if _egfaf || _bedce {
						_fffaf = append(_fffaf, _dd("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e"))
						_cfage = true
						if _gcfe() {
							return true
						}
					}
				}
			}
			if _fgcfe.Alternate != nil {
				return _efgb(_fgcfe.Alternate)
			}
		}
		return false
	}
	for _, _ggcc := range _dfab.GetObjectNums() {
		_bafd, _cccb := _dfab.GetIndirectObjectByNumber(_ggcc)
		if _cccb != nil {
			continue
		}
		_gdgcc, _gbdd := _cb.GetStream(_bafd)
		if !_gbdd {
			continue
		}
		_fbfd, _gbdd := _cb.GetName(_gdgcc.Get("\u0054\u0079\u0070\u0065"))
		if !_gbdd || _fbfd.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_gbgb, _gbdd := _cb.GetName(_gdgcc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_gbdd {
			continue
		}
		_cdfd[_gdgcc] = struct{}{}
		switch _gbgb.String() {
		case "\u0049\u006d\u0061g\u0065":
			_ffed, _defgg := _g.NewXObjectImageFromStream(_gdgcc)
			if _defgg != nil {
				continue
			}
			_cdfd[_gdgcc] = struct{}{}
			if _efgb(_ffed.ColorSpace) {
				return _fffaf
			}
		case "\u0046\u006f\u0072\u006d":
			_bddcf, _bafbf := _cb.GetDict(_gdgcc.Get("\u0047\u0072\u006fu\u0070"))
			if !_bafbf {
				continue
			}
			_geef := _bddcf.Get("\u0043\u0053")
			if _geef == nil {
				continue
			}
			_dcba, _dabf := _g.NewPdfColorspaceFromPdfObject(_geef)
			if _dabf != nil {
				continue
			}
			if _efgb(_dcba) {
				return _fffaf
			}
		}
	}
	for _, _afbe := range _dfab.PageList {
		_bdfd, _caea := _afbe.GetContentStreams()
		if _caea != nil {
			continue
		}
		for _, _fcfdf := range _bdfd {
			_dabd, _gfec := _ee.NewContentStreamParser(_fcfdf).Parse()
			if _gfec != nil {
				continue
			}
			for _, _feeda := range *_dabd {
				if len(_feeda.Params) > 1 {
					continue
				}
				switch _feeda.Operand {
				case "\u0042\u0049":
					_aaeb, _afefd := _feeda.Params[0].(*_ee.ContentStreamInlineImage)
					if !_afefd {
						continue
					}
					_ffaec, _dbdf := _aaeb.GetColorSpace(_afbe.Resources)
					if _dbdf != nil {
						continue
					}
					if _efgb(_ffaec) {
						return _fffaf
					}
				case "\u0044\u006f":
					_adgea, _dcfda := _cb.GetName(_feeda.Params[0])
					if !_dcfda {
						continue
					}
					_bfddc, _bgeg := _afbe.Resources.GetXObjectByName(*_adgea)
					if _, _dbabd := _cdfd[_bfddc]; _dbabd {
						continue
					}
					switch _bgeg {
					case _g.XObjectTypeImage:
						_gfcg, _cfec := _g.NewXObjectImageFromStream(_bfddc)
						if _cfec != nil {
							continue
						}
						_cdfd[_bfddc] = struct{}{}
						if _efgb(_gfcg.ColorSpace) {
							return _fffaf
						}
					case _g.XObjectTypeForm:
						_fcfge, _efegb := _cb.GetDict(_bfddc.Get("\u0047\u0072\u006fu\u0070"))
						if !_efegb {
							continue
						}
						_eegec, _efegb := _cb.GetName(_fcfge.Get("\u0043\u0053"))
						if !_efegb {
							continue
						}
						_aabf, _fdgf := _g.NewPdfColorspaceFromPdfObject(_eegec)
						if _fdgf != nil {
							continue
						}
						_cdfd[_bfddc] = struct{}{}
						if _efgb(_aabf) {
							return _fffaf
						}
					}
				}
			}
		}
	}
	return _fffaf
}

// Profile2B is the implementation of the PDF/A-2B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2B struct{ profile2 }

// ApplyStandard tries to change the content of the writer to match the PDF/A-2 standard.
// Implements model.StandardApplier.
func (_ffbb *profile2) ApplyStandard(document *_gc.Document) (_ffba error) {
	_cbd(document, 7)
	if _ffba = _abf(document, _ffbb._afef.Now); _ffba != nil {
		return _ffba
	}
	if _ffba = _ced(document); _ffba != nil {
		return _ffba
	}
	_bbffa, _cbacd := _fbb(_ffbb._afef.CMYKDefaultColorSpace, _ffbb._gaeg)
	_ffba = _cgcb(document, []pageColorspaceOptimizeFunc{_bbffa}, []documentColorspaceOptimizeFunc{_cbacd})
	if _ffba != nil {
		return _ffba
	}
	_ebe(document)
	if _ffba = _gcaf(document); _ffba != nil {
		return _ffba
	}
	if _ffba = _cgd(document, _ffbb._gaeg._da); _ffba != nil {
		return _ffba
	}
	if _ffba = _cad(document); _ffba != nil {
		return _ffba
	}
	if _ffba = _gfgd(document); _ffba != nil {
		return _ffba
	}
	if _ffba = _abb(document); _ffba != nil {
		return _ffba
	}
	if _ffba = _bcea(document); _ffba != nil {
		return _ffba
	}
	if _ffbb._gaeg._dag == "\u0041" {
		_edfa(document)
	}
	if _ffba = _badb(document, _ffbb._gaeg._da); _ffba != nil {
		return _ffba
	}
	if _ffba = _cea(document); _ffba != nil {
		return _ffba
	}
	if _gbea := _gda(document, _ffbb._gaeg, _ffbb._afef.Xmp); _gbea != nil {
		return _gbea
	}
	if _ffbb._gaeg == _bf() {
		if _ffba = _gcd(document); _ffba != nil {
			return _ffba
		}
	}
	if _ffba = _edee(document); _ffba != nil {
		return _ffba
	}
	if _ffba = _cgcc(document); _ffba != nil {
		return _ffba
	}
	if _ffba = _gcfa(document); _ffba != nil {
		return _ffba
	}
	return nil
}
func _dgedf(_effa *_g.CompliancePdfReader) ViolatedRule {
	if _effa.ParserMetadata().HeaderPosition() != 0 {
		return _dd("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	if _effa.PdfVersion().Major != 1 {
		return _dd("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	if _effa.PdfVersion().Minor < 0 || _effa.PdfVersion().Minor > 7 {
		return _dd("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	return _dfa
}
func _dbcb(_acddc string, _acbbe string, _dbdge string) (string, bool) {
	_cfcf := _a.Index(_acddc, _acbbe)
	if _cfcf == -1 {
		return "", false
	}
	_cfcf += len(_acbbe)
	_ffcd := _a.Index(_acddc[_cfcf:], _dbdge)
	if _ffcd == -1 {
		return "", false
	}
	_ffcd = _cfcf + _ffcd
	return _acddc[_cfcf:_ffcd], true
}
func _edg(_dffe *_gc.Document) error {
	for _, _fff := range _dffe.Objects {
		_cce, _dbd := _cb.GetDict(_fff)
		if !_dbd {
			continue
		}
		_cgeg := _cce.Get("\u0054\u0079\u0070\u0065")
		if _cgeg == nil {
			continue
		}
		if _eae, _eafg := _cb.GetName(_cgeg); _eafg && _eae.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_cec, _bab := _cb.GetBool(_cce.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _bab {
			if bool(*_cec) {
				_cce.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cb.MakeBool(false))
			}
		}
		_ggc := _cce.Get("\u0041")
		if _ggc != nil {
			_cce.Remove("\u0041")
		}
		_gfg, _bab := _cb.GetArray(_cce.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if _bab {
			for _fgg := 0; _fgg < _gfg.Len(); _fgg++ {
				_agf, _fgd := _cb.GetDict(_gfg.Get(_fgg))
				if !_fgd {
					continue
				}
				if _agf.Get("\u0041\u0041") != nil {
					_agf.Remove("\u0041\u0041")
				}
			}
		}
	}
	return nil
}

type documentImages struct {
	_egc, _eb, _cf bool
	_fag           map[_cb.PdfObject]struct{}
	_geb           []*imageInfo
}

// NewProfile1B creates a new Profile1B with the given options.
func NewProfile1B(options *Profile1Options) *Profile1B {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_cfdg(options)
	return &Profile1B{profile1{_gfcf: *options, _efc: _eaa()}}
}
func _gdcf(_fgdb *_g.CompliancePdfReader) (_fbee ViolatedRule) {
	for _, _cgfge := range _fgdb.GetObjectNums() {
		_gbcc, _agff := _fgdb.GetIndirectObjectByNumber(_cgfge)
		if _agff != nil {
			continue
		}
		_gabd, _ebce := _cb.GetStream(_gbcc)
		if !_ebce {
			continue
		}
		_ebgg, _ebce := _cb.GetName(_gabd.Get("\u0054\u0079\u0070\u0065"))
		if !_ebce {
			continue
		}
		if *_ebgg != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_gfbf, _ebce := _cb.GetName(_gabd.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_ebce {
			continue
		}
		if *_gfbf == "\u0050\u0053" {
			return _dd("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _gabd.Get("\u0050\u0053") != nil {
			return _dd("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _fbee
}

// NewProfile1A creates a new Profile1A with given options.
func NewProfile1A(options *Profile1Options) *Profile1A {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_cfdg(options)
	return &Profile1A{profile1{_gfcf: *options, _efc: _cbb()}}
}
func _abb(_bdg *_gc.Document) error {
	_eaf := map[string]*_cb.PdfObjectDictionary{}
	_gbe := _fg.NewFinder(&_fg.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066"}})
	_daf := map[_cb.PdfObject]struct{}{}
	_gfb := map[_cb.PdfObject]struct{}{}
	for _, _cag := range _bdg.Objects {
		_dfbg, _aed := _cb.GetDict(_cag)
		if !_aed {
			continue
		}
		_dff := _dfbg.Get("\u0054\u0079\u0070\u0065")
		if _dff == nil {
			continue
		}
		if _ff, _fbc := _cb.GetName(_dff); _fbc && _ff.String() != "\u0046\u006f\u006e\u0074" {
			continue
		}
		if _, _abee := _daf[_cag]; _abee {
			continue
		}
		_dfg, _ceg := _g.NewPdfFontFromPdfObject(_dfbg)
		if _ceg != nil {
			_ge.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
			return _ceg
		}
		_bgbe, _ceg := _dfg.GetFontDescriptor()
		if _ceg != nil {
			return _ceg
		}
		if _bgbe != nil && (_bgbe.FontFile != nil || _bgbe.FontFile2 != nil || _bgbe.FontFile3 != nil) {
			continue
		}
		_eged := _dfg.BaseFont()
		if _eged == "" {
			return _c.Errorf("\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0073\u0020\u0073\u0079\u006e\u0074\u0061\u0078\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069d\u0020\u002d\u0020\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u0075\u006ed\u0065\u0066\u0069n\u0065\u0064\u003a\u0020\u0025\u0073", _dfbg.String())
		}
		_befb, _bade := _eaf[_eged]
		if !_bade {
			if len(_eged) > 7 && _eged[6] == '+' {
				_eged = _eged[7:]
			}
			_dfag := []string{_eged, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
			for _, _acd := range _dfag {
				_ge.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _acd)
				if _befb, _bade = _eaf[_acd]; _bade {
					break
				}
				_aab := _gbe.Match(_acd)
				if _aab == nil {
					_ge.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _acd)
					continue
				}
				_eabb, _dce := _g.NewPdfFontFromTTFFile(_aab.Filename)
				if _dce != nil {
					return _dce
				}
				_fedg := _eabb.FontDescriptor()
				if _fedg.FontFile != nil {
					if _, _bade = _gfb[_fedg.FontFile]; !_bade {
						_bdg.Objects = append(_bdg.Objects, _fedg.FontFile)
						_gfb[_fedg.FontFile] = struct{}{}
					}
				}
				if _fedg.FontFile2 != nil {
					if _, _bade = _gfb[_fedg.FontFile2]; !_bade {
						_bdg.Objects = append(_bdg.Objects, _fedg.FontFile2)
						_gfb[_fedg.FontFile2] = struct{}{}
					}
				}
				if _fedg.FontFile3 != nil {
					if _, _bade = _gfb[_fedg.FontFile3]; !_bade {
						_bdg.Objects = append(_bdg.Objects, _fedg.FontFile3)
						_gfb[_fedg.FontFile3] = struct{}{}
					}
				}
				_fbcd, _bdb := _eabb.ToPdfObject().(*_cb.PdfIndirectObject)
				if !_bdb {
					_ge.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
					continue
				}
				_fef, _bdb := _fbcd.PdfObject.(*_cb.PdfObjectDictionary)
				if !_bdb {
					_ge.Log.Debug("\u0046\u006fn\u0074\u0020\u0074\u0079p\u0065\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
					continue
				}
				_eaf[_acd] = _fef
				_befb = _fef
				break
			}
			if _befb == nil {
				_ge.Log.Debug("\u004e\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006eg\u0020\u0066\u006f\u006e\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u003a\u0020\u0025\u0073", _dfg.BaseFont())
				return _f.New("\u006e\u006f m\u0061\u0074\u0063h\u0069\u006e\u0067\u0020fon\u0074 f\u006f\u0075\u006e\u0064\u0020\u0069\u006e t\u0068\u0065\u0020\u0073\u0079\u0073\u0074e\u006d")
			}
		}
		for _, _eff := range _befb.Keys() {
			_dfbg.Set(_eff, _befb.Get(_eff))
		}
		_cafd := _befb.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _cafd != nil {
			if _, _bade = _gfb[_cafd]; !_bade {
				_bdg.Objects = append(_bdg.Objects, _cafd)
				_gfb[_cafd] = struct{}{}
			}
		}
		_daf[_cag] = struct{}{}
		_gcad := _dfbg.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
		if _gcad != nil {
			_bdg.Objects = append(_bdg.Objects, _gcad)
			_gfb[_gcad] = struct{}{}
		}
	}
	return nil
}
func _babg(_badc *_g.PdfFont, _fgeg *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_begac = "\u0036.\u0033\u002e\u0037\u002d\u0032"
		_bgac  = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _fcffg string
	if _fddg, _baac := _cb.GetName(_fgeg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _baac {
		_fcffg = _fddg.String()
	}
	if _fcffg != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _dfa
	}
	_cecgc := _badc.FontDescriptor()
	_gfgec, _cbbc := _cb.GetIntVal(_cecgc.Flags)
	if !_cbbc {
		_ge.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _dd(_begac, _bgac)
	}
	_fggd := (uint32(_gfgec) >> 3) & 1
	_bgea := _fggd != 0
	if !_bgea {
		return _dfa
	}
	if _fgeg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _dd(_begac, _bgac)
	}
	return _dfa
}
func _fbbb(_eabf *_g.CompliancePdfReader) (_cbeeb []ViolatedRule) {
	_cffcb := true
	_fgfbb, _dfce := _eabf.GetCatalogMarkInfo()
	if !_dfce {
		_cffcb = false
	} else {
		_dadb, _beggf := _cb.GetDict(_fgfbb)
		if _beggf {
			_acgbf, _befa := _cb.GetBool(_dadb.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_acgbf) || !_befa {
				_cffcb = false
			}
		} else {
			_cffcb = false
		}
	}
	if !_cffcb {
		_cbeeb = append(_cbeeb, _dd("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_fbab, _dfce := _eabf.GetCatalogStructTreeRoot()
	if !_dfce {
		_cbeeb = append(_cbeeb, _dd("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_dbgf, _dfce := _cb.GetDict(_fbab)
	if _dfce {
		_eaefe, _ffeg := _cb.GetName(_dbgf.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _ffeg {
			_dfbee, _fggef := _cb.GetDict(_eaefe)
			if _fggef {
				for _, _bcbd := range _dfbee.Keys() {
					_ddegg := _dfbee.Get(_bcbd)
					if _ddegg == nil {
						_cbeeb = append(_cbeeb, _dd("\u0036.\u0037\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _cbeeb
}

var _ Profile = (*Profile2B)(nil)

func _gdbe(_egbb *_cb.PdfObjectDictionary, _ddfb map[*_cb.PdfObjectStream][]byte, _eegece map[*_cb.PdfObjectStream]*_ad.CMap) ViolatedRule {
	const (
		_efdgg = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0034"
		_bedbc = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _dcbc string
	if _afccd, _dfgaf := _cb.GetName(_egbb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dfgaf {
		_dcbc = _afccd.String()
	}
	if _dcbc != "\u0054\u0079\u0070e\u0030" {
		return _dfa
	}
	_ffaff := _egbb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _gagbf := _cb.GetName(_ffaff); _gagbf {
		return _dfa
	}
	_eebe, _gdacd := _cb.GetStream(_ffaff)
	if !_gdacd {
		return _dd(_efdgg, _bedbc)
	}
	_acdba, _ccfd := _cdfbg(_eebe, _ddfb, _eegece)
	if _ccfd != nil {
		return _dd(_efdgg, _bedbc)
	}
	_gaec, _cgeb := _cb.GetIntVal(_eebe.Get("\u0057\u004d\u006fd\u0065"))
	_agda, _cbcg := _acdba.WMode()
	if _cgeb && _cbcg {
		if _agda != _gaec {
			return _dd(_efdgg, _bedbc)
		}
	}
	if (_cgeb && !_cbcg) || (!_cgeb && _cbcg) {
		return _dd(_efdgg, _bedbc)
	}
	return _dfa
}
func _gagc(_cbba *_g.CompliancePdfReader) (_fbbcf []ViolatedRule) {
	var _egdd, _gfca, _dcfcf bool
	if _cbba.ParserMetadata().HasNonConformantStream() {
		_fbbcf = []ViolatedRule{_dd("\u0036.\u0031\u002e\u0037\u002d\u0032", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _adag := range _cbba.GetObjectNums() {
		_cdcee, _ := _cbba.GetIndirectObjectByNumber(_adag)
		if _cdcee == nil {
			continue
		}
		_bfbbb, _fdeba := _cb.GetStream(_cdcee)
		if !_fdeba {
			continue
		}
		if !_egdd {
			_dbba := _bfbbb.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _dbba == nil {
				_fbbcf = append(_fbbcf, _dd("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_egdd = true
			} else {
				_bfff, _dbcdb := _cb.GetIntVal(_dbba)
				if !_dbcdb {
					_fbbcf = append(_fbbcf, _dd("\u0036.\u0031\u002e\u0037\u002d\u0031", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_egdd = true
				} else {
					if len(_bfbbb.Stream) != _bfff {
						_fbbcf = append(_fbbcf, _dd("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_egdd = true
					}
				}
			}
		}
		if !_gfca {
			if _bfbbb.Get("\u0046") != nil {
				_gfca = true
				_fbbcf = append(_fbbcf, _dd("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _bfbbb.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_gfca {
				_gfca = true
				_fbbcf = append(_fbbcf, _dd("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _bfbbb.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_gfca {
				_gfca = true
				_fbbcf = append(_fbbcf, _dd("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_dcfcf {
			_bfag, _fceb := _cb.GetName(_cb.TraceToDirectObject(_bfbbb.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_fceb {
				continue
			}
			if *_bfag == _cb.StreamEncodingFilterNameLZW {
				_dcfcf = true
				_fbbcf = append(_fbbcf, _dd("\u0036.\u0031\u002e\u0037\u002d\u0034", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _fbbcf
}
func _aecf(_fedba *_g.CompliancePdfReader) (_babe []ViolatedRule) {
	var _fagde, _baadg bool
	_fege := func() bool { return _fagde && _baadg }
	for _, _dgad := range _fedba.GetObjectNums() {
		_cagea, _babdd := _fedba.GetIndirectObjectByNumber(_dgad)
		if _babdd != nil {
			_ge.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _dgad, _babdd)
			continue
		}
		_ebgbb, _fdebbd := _cb.GetDict(_cagea)
		if !_fdebbd {
			continue
		}
		_cefb, _fdebbd := _cb.GetName(_ebgbb.Get("\u0054\u0079\u0070\u0065"))
		if !_fdebbd {
			continue
		}
		if *_cefb != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_bfcge, _fdebbd := _cb.GetName(_ebgbb.Get("\u0053"))
		if !_fdebbd {
			if !_fagde {
				_babe = append(_babe, _dd("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_fagde = true
				if _fege() {
					return _babe
				}
			}
			continue
		}
		switch _g.PdfActionType(*_bfcge) {
		case _g.ActionTypeLaunch, _g.ActionTypeSound, _g.ActionTypeMovie, _g.ActionTypeResetForm, _g.ActionTypeImportData, _g.ActionTypeJavaScript, _g.ActionTypeHide, _g.ActionTypeSetOCGState, _g.ActionTypeRendition, _g.ActionTypeTrans, _g.ActionTypeGoTo3DView:
			if !_fagde {
				_babe = append(_babe, _dd("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_fagde = true
				if _fege() {
					return _babe
				}
			}
			continue
		case _g.ActionTypeNamed:
			if !_baadg {
				_bcdgf, _bcgdfe := _cb.GetName(_ebgbb.Get("\u004e"))
				if !_bcgdfe {
					_babe = append(_babe, _dd("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_baadg = true
					if _fege() {
						return _babe
					}
					continue
				}
				switch *_bcdgf {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_babe = append(_babe, _dd("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_baadg = true
					if _fege() {
						return _babe
					}
					continue
				}
			}
		}
	}
	return _babe
}
func _cad(_agdd *_gc.Document) error {
	_gdg, _gcde := _agdd.GetPages()
	if !_gcde {
		return nil
	}
	for _, _gbd := range _gdg {
		_afeb, _ecgd := _cb.GetArray(_gbd.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_ecgd {
			continue
		}
		for _, _aeg := range _afeb.Elements() {
			_aeg = _cb.ResolveReference(_aeg)
			if _, _ebda := _aeg.(*_cb.PdfObjectNull); _ebda {
				continue
			}
			_cgfg, _aagf := _cb.GetDict(_aeg)
			if !_aagf {
				continue
			}
			_gde, _ := _cb.GetIntVal(_cgfg.Get("\u0046"))
			_gde &= ^(1 << 0)
			_gde &= ^(1 << 1)
			_gde &= ^(1 << 5)
			_gde &= ^(1 << 8)
			_gde |= 1 << 2
			_cgfg.Set("\u0046", _cb.MakeInteger(int64(_gde)))
			_effg := false
			if _bfac := _cgfg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _bfac != nil {
				_fdg, _dgga := _cb.GetName(_bfac)
				if _dgga && _fdg.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_effg = true
					if _cgfg.Get("\u0041\u0041") != nil {
						_cgfg.Remove("\u0041\u0041")
					}
					if _cgfg.Get("\u0041") != nil {
						_cgfg.Remove("\u0041")
					}
				}
				if _dgga && _fdg.String() == "\u0054\u0065\u0078\u0074" {
					_deae, _ := _cb.GetIntVal(_cgfg.Get("\u0046"))
					_deae |= 1 << 3
					_deae |= 1 << 4
					_cgfg.Set("\u0046", _cb.MakeInteger(int64(_deae)))
				}
			}
			_fgca, _aagf := _cb.GetDict(_cgfg.Get("\u0041\u0050"))
			if _aagf {
				_cfce := _fgca.Get("\u004e")
				if _cfce == nil {
					continue
				}
				if len(_fgca.Keys()) > 1 {
					_fgca.Clear()
					_fgca.Set("\u004e", _cfce)
				}
				if _effg {
					_gbab, _begd := _cb.GetName(_cgfg.Get("\u0046\u0054"))
					if _begd && *_gbab == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}

// Error implements error interface.
func (_fd VerificationError) Error() string {
	_fa := _a.Builder{}
	_fa.WriteString("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u003a\u0020")
	_fa.WriteString(_c.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _fd.ConformanceLevel, _fd.ConformanceVariant))
	_fa.WriteString("\u0020\u0056\u0069\u006f\u006c\u0061\u0074\u0065\u0064\u0020\u0072\u0075l\u0065\u0073\u003a\u0020")
	for _fdf, _fde := range _fd.ViolatedRules {
		_fa.WriteString(_fde.String())
		if _fdf != len(_fd.ViolatedRules)-1 {
			_fa.WriteRune('\n')
		}
	}
	return _fa.String()
}
func _edbe(_ffbbdg *_g.CompliancePdfReader) ViolatedRule { return _dfa }
func _bcdc(_fced *_g.CompliancePdfReader) ViolatedRule {
	_bgce, _gdd := _fced.PdfReader.GetTrailer()
	if _gdd != nil {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _bgce.Get("\u0049\u0044") == nil {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _bgce.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _dd("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _dfa
}
func _caae(_efd *_gc.Document, _bgf bool) error {
	_gdb, _aec := _efd.GetPages()
	if !_aec {
		return nil
	}
	for _, _daed := range _gdb {
		_bffc := _daed.FindXObjectForms()
		for _, _bfbb := range _bffc {
			_efeg, _effb := _g.NewXObjectFormFromStream(_bfbb)
			if _effb != nil {
				return _effb
			}
			_gfcc, _effb := _efeg.GetContentStream()
			if _effb != nil {
				return _effb
			}
			_ggba := _ee.NewContentStreamParser(string(_gfcc))
			_bcef, _effb := _ggba.Parse()
			if _effb != nil {
				return _effb
			}
			_cdf, _effb := _edb(_efeg.Resources, _bcef, _bgf)
			if _effb != nil {
				return _effb
			}
			if len(_cdf) == 0 {
				continue
			}
			if _effb = _efeg.SetContentStream(_cdf, _cb.NewFlateEncoder()); _effb != nil {
				return _effb
			}
			_efeg.ToPdfObject()
		}
	}
	return nil
}
func _cfcda(_beec *_g.CompliancePdfReader) ViolatedRule { return _dfa }
func _egee(_daeb *_g.CompliancePdfReader) ViolatedRule  { return _dfa }
func _ebcb(_cegeg *_g.PdfFont, _fddd *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_feed = "\u0036.\u0033\u002e\u0037\u002d\u0033"
		_fbef = "\u0046\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0027\u0020\u0022\u0063\u006d\u0061\u0070\u0022\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0073\u0079\u006d\u0062o\u006c\u0069c\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0065\u0078\u0061\u0063\u0074\u006cy\u0020\u006f\u006ee\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u002e"
	)
	var _cgac string
	if _becg, _ecfc := _cb.GetName(_fddd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ecfc {
		_cgac = _becg.String()
	}
	if _cgac != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _dfa
	}
	_eabe := _cegeg.FontDescriptor()
	_bgeb, _cacg := _cb.GetIntVal(_eabe.Flags)
	if !_cacg {
		_ge.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _dd(_feed, _fbef)
	}
	_ebbc := (uint32(_bgeb) >> 3) != 0
	if !_ebbc {
		return _dfa
	}
	return _dfa
}
func _gcfa(_bgd *_gc.Document) error {
	_eeag, _cgee := _bgd.FindCatalog()
	if !_cgee {
		return _f.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _eeag.Object.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		_eeag.Object.Remove("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073")
	}
	return nil
}
func _gce() standardType { return standardType{_da: 2, _dag: "\u0042"} }
func _ggacc(_gbfce *_g.CompliancePdfReader) (_ccebce []ViolatedRule) {
	for _, _fdbd := range _gbfce.GetObjectNums() {
		_bfagb, _bccd := _gbfce.GetIndirectObjectByNumber(_fdbd)
		if _bccd != nil {
			continue
		}
		_acbcf, _gdddf := _cb.GetDict(_bfagb)
		if !_gdddf {
			continue
		}
		_dacce, _gdddf := _cb.GetName(_acbcf.Get("\u0054\u0079\u0070\u0065"))
		if !_gdddf {
			continue
		}
		if _dacce.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_cfead, _gdddf := _cb.GetBool(_acbcf.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _gdddf && bool(*_cfead) {
			_ccebce = append(_ccebce, _dd("\u0036.\u0034\u002e\u0031\u002d\u0033", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
		if _acbcf.Get("\u0058\u0046\u0041") != nil {
			_ccebce = append(_ccebce, _dd("\u0036.\u0034\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020i\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020t\u0068\u0061\u0074\u0020f\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065 \u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d \u006b\u0065\u0079\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006f\u0066 \u0061 \u0050\u0044F\u002fA\u002d\u0032\u0020\u0066ile\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065n\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u0046\u0041\u0020\u006b\u0065y."))
		}
	}
	_dgbfe, _bcgdf := _addf(_gbfce)
	if _bcgdf && _dgbfe.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_ccebce = append(_ccebce, _dd("\u0036.\u0034\u002e\u0032\u002d\u0032", "\u0041\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006cog\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u002e"))
	}
	return _ccebce
}
func _ebcbc(_dgcb *_g.CompliancePdfReader) (_efada []ViolatedRule) { return _efada }
func _bbgg(_fbd *_gc.Document) (*_cb.PdfObjectDictionary, bool) {
	_dga, _bff := _fbd.FindCatalog()
	if !_bff {
		return nil, false
	}
	_ddfcg, _bff := _cb.GetArray(_dga.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_bff {
		return nil, false
	}
	if _ddfcg.Len() == 0 {
		return nil, false
	}
	return _cb.GetDict(_ddfcg.Get(0))
}
func _eccb(_ceae *_g.CompliancePdfReader) []ViolatedRule { return nil }
func _dac(_eaba *_g.XObjectImage, _cae imageModifications) error {
	_dg, _cff := _eaba.ToImage()
	if _cff != nil {
		return _cff
	}
	if _cae._af != nil {
		_eaba.Filter = _cae._af
	}
	_baf := _cb.MakeDict()
	_baf.Set("\u0051u\u0061\u006c\u0069\u0074\u0079", _cb.MakeInteger(100))
	_baf.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", _cb.MakeInteger(1))
	_eaba.Decode = nil
	if _cff = _eaba.SetImage(_dg, nil); _cff != nil {
		return _cff
	}
	_eaba.ToPdfObject()
	return nil
}

var _ Profile = (*Profile2U)(nil)

func _cgd(_fbfe *_gc.Document, _cfe int) error {
	_fge := map[*_cb.PdfObjectStream]struct{}{}
	for _, _cffe := range _fbfe.Objects {
		_edf, _bbggb := _cb.GetStream(_cffe)
		if !_bbggb {
			continue
		}
		if _, _bbggb = _fge[_edf]; _bbggb {
			continue
		}
		_fge[_edf] = struct{}{}
		_accg, _bbggb := _cb.GetName(_edf.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_bbggb {
			continue
		}
		if _edf.Get("\u0052\u0065\u0066") != nil {
			_edf.Remove("\u0052\u0065\u0066")
		}
		if _accg.String() == "\u0050\u0053" {
			_edf.Remove("\u0050\u0053")
			continue
		}
		if _accg.String() == "\u0046\u006f\u0072\u006d" {
			if _edf.Get("\u004f\u0050\u0049") != nil {
				_edf.Remove("\u004f\u0050\u0049")
			}
			if _edf.Get("\u0050\u0053") != nil {
				_edf.Remove("\u0050\u0053")
			}
			if _faaf := _edf.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _faaf != nil {
				if _efba, _bge := _cb.GetName(_faaf); _bge && *_efba == "\u0050\u0053" {
					_edf.Remove("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032")
				}
			}
			continue
		}
		if _accg.String() == "\u0049\u006d\u0061g\u0065" {
			_ccfe, _dbeb := _cb.GetBool(_edf.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _dbeb && bool(*_ccfe) {
				_edf.Set("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _cb.MakeBool(false))
			}
			if _cfe == 2 {
				if _edf.Get("\u004f\u0050\u0049") != nil {
					_edf.Remove("\u004f\u0050\u0049")
				}
			}
			if _edf.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
				_edf.Remove("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073")
			}
			continue
		}
	}
	return nil
}

// Profile1A is the implementation of the PDF/A-1A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1A struct{ profile1 }

func _aeffe(_affdd *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_dbdbeb = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0032"
		_cbabe  = "IS\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068a\u0074\u0020\u0061\u006c\u006c\u0020\u0065m\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0073\u0020\u0069n\u0020t\u0068e\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006fG\u0049\u0044M\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079 \u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020\u0061\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070p\u0069\u006e\u0067 f\u0072\u006f\u006d \u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079p\u0068 \u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006fr\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u002e"
	)
	var _fdebb string
	if _fdab, _cecdf := _cb.GetName(_affdd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cecdf {
		_fdebb = _fdab.String()
	}
	if _fdebb != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _dfa
	}
	if _affdd.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _dd(_dbdbeb, _cbabe)
	}
	return _dfa
}
func _bdgde(_bfdgf *_g.CompliancePdfReader) ViolatedRule {
	for _, _beca := range _bfdgf.GetObjectNums() {
		_gbcfb, _gdcfa := _bfdgf.GetIndirectObjectByNumber(_beca)
		if _gdcfa != nil {
			continue
		}
		_eefg, _bbecf := _cb.GetStream(_gbcfb)
		if !_bbecf {
			continue
		}
		_deff, _bbecf := _cb.GetName(_eefg.Get("\u0054\u0079\u0070\u0065"))
		if !_bbecf {
			continue
		}
		if *_deff != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _eefg.Get("\u0053\u004d\u0061s\u006b") != nil {
			return _dd("\u0036\u002e\u0034-\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e \u0053\u004d\u0061\u0073\u006b\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _dfa
}
func (_ega *documentImages) hasOnlyDeviceGray() bool { return _ega._cf && !_ega._egc && !_ega._eb }

// XmpOptions are the options used by the optimization of the XMP metadata.
type XmpOptions struct {

	// Copyright information.
	Copyright string

	// OriginalDocumentID is the original document identifier.
	// By default, if this field is empty the value is extracted from the XMP Metadata or generated UUID.
	OriginalDocumentID string

	// DocumentID is the original document identifier.
	// By default, if this field is empty the value is extracted from the XMP Metadata or generated UUID.
	DocumentID string

	// InstanceID is the original document identifier.
	// By default, if this field is empty the value is set to generated UUID.
	InstanceID string

	// NewDocumentVersion is a flag that defines if a document was overwritten.
	// If the new document was created this should be true. On changing given document file, and overwriting it it should be true.
	NewDocumentVersion bool

	// MarshalIndent defines marshaling indent of the XMP metadata.
	MarshalIndent string

	// MarshalPrefix defines marshaling prefix of the XMP metadata.
	MarshalPrefix string
}

func _eaa() standardType { return standardType{_da: 1, _dag: "\u0042"} }
func _fedaa(_cgfcd *_g.PdfFont, _fdgg *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_gbfc = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0036\u002d\u0033"
		_bddg = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _fafc string
	if _ebceb, _ccdbc := _cb.GetName(_fdgg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ccdbc {
		_fafc = _ebceb.String()
	}
	if _fafc != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _dfa
	}
	_edcb := _cgfcd.FontDescriptor()
	_fcacf, _bfgc := _cb.GetIntVal(_edcb.Flags)
	if !_bfgc {
		_ge.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _dd(_gbfc, _bddg)
	}
	_eedea := (uint32(_fcacf) >> 3) & 1
	_fafd := _eedea != 0
	if !_fafd {
		return _dfa
	}
	if _fdgg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _dd(_gbfc, _bddg)
	}
	return _dfa
}

type imageInfo struct {
	ColorSpace       _cb.PdfObjectName
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_cb.PdfObjectStream
	_dc              bool
}

func _cdccf(_eefd *_g.CompliancePdfReader) ViolatedRule {
	for _, _aebf := range _eefd.PageList {
		_baad, _eabd := _aebf.GetContentStreams()
		if _eabd != nil {
			continue
		}
		for _, _cgbde := range _baad {
			_afgf := _ee.NewContentStreamParser(_cgbde)
			_, _eabd = _afgf.Parse()
			if _eabd != nil {
				return _dd("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _dfa
}

// ApplyStandard tries to change the content of the writer to match the PDF/A-1 standard.
// Implements model.StandardApplier.
func (_fcaf *profile1) ApplyStandard(document *_gc.Document) (_cffb error) {
	_cbd(document, 4)
	if _cffb = _abf(document, _fcaf._gfcf.Now); _cffb != nil {
		return _cffb
	}
	if _cffb = _ced(document); _cffb != nil {
		return _cffb
	}
	_fdfge, _acedc := _fbb(_fcaf._gfcf.CMYKDefaultColorSpace, _fcaf._efc)
	_cffb = _cgcb(document, []pageColorspaceOptimizeFunc{_gbeb, _fdfge}, []documentColorspaceOptimizeFunc{_acedc})
	if _cffb != nil {
		return _cffb
	}
	_ebe(document)
	if _cffb = _cgd(document, _fcaf._efc._da); _cffb != nil {
		return _cffb
	}
	if _cffb = _cgdf(document); _cffb != nil {
		return _cffb
	}
	if _cffb = _gcae(document); _cffb != nil {
		return _cffb
	}
	if _cffb = _beg(document); _cffb != nil {
		return _cffb
	}
	if _cffb = _abb(document); _cffb != nil {
		return _cffb
	}
	if _fcaf._efc._dag == "\u0041" {
		_edfa(document)
	}
	if _cffb = _badb(document, _fcaf._efc._da); _cffb != nil {
		return _cffb
	}
	if _cffb = _cea(document); _cffb != nil {
		return _cffb
	}
	if _agbd := _gda(document, _fcaf._efc, _fcaf._gfcf.Xmp); _agbd != nil {
		return _agbd
	}
	if _fcaf._efc == _cbb() {
		if _cffb = _gcd(document); _cffb != nil {
			return _cffb
		}
	}
	if _cffb = _edg(document); _cffb != nil {
		return _cffb
	}
	return nil
}
func _abf(_ace *_gc.Document, _ccda func() _bee.Time) error {
	_ddfc, _def := _g.NewPdfInfoFromObject(_ace.Info)
	if _def != nil {
		return _def
	}
	if _bfe := _faab(_ddfc, _ccda); _bfe != nil {
		return _bfe
	}
	_ace.Info = _ddfc.ToPdfObject()
	return nil
}
func _gbeb(_dfeg *_gc.Document, _gada *_gc.Page, _dgdf []*_gc.Image) error {
	for _, _bcefa := range _dgdf {
		if _bcefa.SMask == nil {
			continue
		}
		_efag, _ebae := _g.NewXObjectImageFromStream(_bcefa.Stream)
		if _ebae != nil {
			return _ebae
		}
		_cbcce, _ebae := _efag.ToImage()
		if _ebae != nil {
			return _ebae
		}
		_bcf, _ebae := _cbcce.ToGoImage()
		if _ebae != nil {
			return _ebae
		}
		_bdbb, _ebae := _eg.RGBAConverter.Convert(_bcf)
		if _ebae != nil {
			return _ebae
		}
		_bbca := _bdbb.Base()
		_cbbbg := &_g.Image{Width: int64(_bbca.Width), Height: int64(_bbca.Height), BitsPerComponent: int64(_bbca.BitsPerComponent), ColorComponents: _bbca.ColorComponents, Data: _bbca.Data}
		_cbbbg.SetDecode(_bbca.Decode)
		_cbbbg.SetAlpha(_bbca.Alpha)
		if _ebae = _efag.SetImage(_cbbbg, nil); _ebae != nil {
			return _ebae
		}
		_efag.SMask = _cb.MakeNull()
		var _dbfe _cb.PdfObject
		_ebac := -1
		for _ebac, _dbfe = range _dfeg.Objects {
			if _dbfe == _bcefa.SMask.Stream {
				break
			}
		}
		if _ebac != -1 {
			_dfeg.Objects = append(_dfeg.Objects[:_ebac], _dfeg.Objects[_ebac+1:]...)
		}
		_bcefa.SMask = nil
		_efag.ToPdfObject()
	}
	return nil
}
func _dbfb(_efbbg *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_eddd = "\u0036.\u0033\u002e\u0033\u002d\u0032"
		_fcfb = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0054y\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020f\u006f\u0072 \u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067,\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0020\u0064\u0069c\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006f\u0047\u0049D\u004d\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070\u0070\u0069\u006e\u0067\u0020\u0066\u0072\u006f\u006d\u0020\u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061s d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020P\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0054a\u0062\u006c\u0065\u0020\u0035\u002e\u00313"
	)
	var _bdgc string
	if _dfgb, _ggggc := _cb.GetName(_efbbg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ggggc {
		_bdgc = _dfgb.String()
	}
	if _bdgc != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _dfa
	}
	if _efbbg.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _dd(_eddd, _fcfb)
	}
	return _dfa
}
func _dd(_ag string, _ege string) ViolatedRule { return ViolatedRule{RuleNo: _ag, Detail: _ege} }

// ViolatedRule is the structure that defines violated PDF/A rule.
type ViolatedRule struct {
	RuleNo string
	Detail string
}

// DefaultProfile1Options are the default options for the Profile1.
func DefaultProfile1Options() *Profile1Options {
	return &Profile1Options{Now: _bee.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}
func _baed(_agba *_g.CompliancePdfReader) ViolatedRule {
	for _, _gedbc := range _agba.PageList {
		_fgcf, _dafg := _gedbc.GetContentStreams()
		if _dafg != nil {
			continue
		}
		for _, _aded := range _fgcf {
			_gfbc := _ee.NewContentStreamParser(_aded)
			_, _dafg = _gfbc.Parse()
			if _dafg != nil {
				return _dd("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _dfa
}
func _edb(_aef *_g.PdfPageResources, _fedb *_ee.ContentStreamOperations, _gag bool) ([]byte, error) {
	var _cded bool
	for _, _dcc := range *_fedb {
	_fgbf:
		switch _dcc.Operand {
		case "\u0042\u0049":
			_bbb, _eeaa := _dcc.Params[0].(*_ee.ContentStreamInlineImage)
			if !_eeaa {
				break
			}
			_afae, _ggcf := _bbb.GetColorSpace(_aef)
			if _ggcf != nil {
				return nil, _ggcf
			}
			switch _afae.(type) {
			case *_g.PdfColorspaceDeviceCMYK:
				if _gag {
					break _fgbf
				}
			case *_g.PdfColorspaceDeviceGray:
			case *_g.PdfColorspaceDeviceRGB:
				if !_gag {
					break _fgbf
				}
			default:
				break _fgbf
			}
			_cded = true
			_aee, _ggcf := _bbb.ToImage(_aef)
			if _ggcf != nil {
				return nil, _ggcf
			}
			_ddff, _ggcf := _aee.ToGoImage()
			if _ggcf != nil {
				return nil, _ggcf
			}
			if _gag {
				_ddff, _ggcf = _eg.CMYKConverter.Convert(_ddff)
			} else {
				_ddff, _ggcf = _eg.NRGBAConverter.Convert(_ddff)
			}
			if _ggcf != nil {
				return nil, _ggcf
			}
			_acb, _eeaa := _ddff.(_eg.Image)
			if !_eeaa {
				return nil, _f.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
			}
			_aeee := _acb.Base()
			_bcce := _g.Image{Width: int64(_aeee.Width), Height: int64(_aeee.Height), BitsPerComponent: int64(_aeee.BitsPerComponent), ColorComponents: _aeee.ColorComponents, Data: _aeee.Data}
			_bcce.SetDecode(_aeee.Decode)
			_bcce.SetAlpha(_aeee.Alpha)
			_eafga, _ggcf := _bbb.GetEncoder()
			if _ggcf != nil {
				_eafga = _cb.NewFlateEncoder()
			}
			_ffaf, _ggcf := _ee.NewInlineImageFromImage(_bcce, _eafga)
			if _ggcf != nil {
				return nil, _ggcf
			}
			_dcc.Params[0] = _ffaf
		case "\u0047", "\u0067":
			if len(_dcc.Params) != 1 {
				break
			}
			_cfgd, _aaff := _cb.GetNumberAsFloat(_dcc.Params[0])
			if _aaff != nil {
				break
			}
			if _gag {
				_dcc.Params = []_cb.PdfObject{_cb.MakeFloat(0), _cb.MakeFloat(0), _cb.MakeFloat(0), _cb.MakeFloat(1 - _cfgd)}
				_fgae := "\u004b"
				if _dcc.Operand == "\u0067" {
					_fgae = "\u006b"
				}
				_dcc.Operand = _fgae
			} else {
				_dcc.Params = []_cb.PdfObject{_cb.MakeFloat(_cfgd), _cb.MakeFloat(_cfgd), _cb.MakeFloat(_cfgd)}
				_bagc := "\u0052\u0047"
				if _dcc.Operand == "\u0067" {
					_bagc = "\u0072\u0067"
				}
				_dcc.Operand = _bagc
			}
			_cded = true
		case "\u0052\u0047", "\u0072\u0067":
			if !_gag {
				break
			}
			if len(_dcc.Params) != 3 {
				break
			}
			_dffd, _dceg := _cb.GetNumbersAsFloat(_dcc.Params)
			if _dceg != nil {
				break
			}
			_cded = true
			_edea, _bfd, _bbc := _dffd[0], _dffd[1], _dffd[2]
			_faa, _ffgf, _dgdg, _fdc := _d.RGBToCMYK(uint8(_edea*255), uint8(_bfd*255), uint8(255*_bbc))
			_dcc.Params = []_cb.PdfObject{_cb.MakeFloat(float64(_faa) / 255), _cb.MakeFloat(float64(_ffgf) / 255), _cb.MakeFloat(float64(_dgdg) / 255), _cb.MakeFloat(float64(_fdc) / 255)}
			_bcg := "\u004b"
			if _dcc.Operand == "\u0072\u0067" {
				_bcg = "\u006b"
			}
			_dcc.Operand = _bcg
		case "\u004b", "\u006b":
			if _gag {
				break
			}
			if len(_dcc.Params) != 4 {
				break
			}
			_ffe, _caeb := _cb.GetNumbersAsFloat(_dcc.Params)
			if _caeb != nil {
				break
			}
			_edgb, _gbbg, _bbeca, _fdfg := _ffe[0], _ffe[1], _ffe[2], _ffe[3]
			_gab, _bgga, _gaf := _d.CMYKToRGB(uint8(255*_edgb), uint8(255*_gbbg), uint8(255*_bbeca), uint8(255*_fdfg))
			_dcc.Params = []_cb.PdfObject{_cb.MakeFloat(float64(_gab) / 255), _cb.MakeFloat(float64(_bgga) / 255), _cb.MakeFloat(float64(_gaf) / 255)}
			_dab := "\u0052\u0047"
			if _dcc.Operand == "\u006b" {
				_dab = "\u0072\u0067"
			}
			_dcc.Operand = _dab
			_cded = true
		}
	}
	if !_cded {
		return nil, nil
	}
	_fffg := _ee.NewContentCreator()
	for _, _dgca := range *_fedb {
		_fffg.AddOperand(*_dgca)
	}
	_aga := _fffg.Bytes()
	return _aga, nil
}
func _ced(_beda *_gc.Document) error {
	_gcc, _fda := _beda.FindCatalog()
	if !_fda {
		return _f.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_gcc.SetVersion()
	return nil
}
func _cage(_affca *_g.CompliancePdfReader) (_cdfe []ViolatedRule) {
	_ddfd := _affca.ParserMetadata()
	if _ddfd.HasInvalidSubsectionHeader() {
		_cdfe = append(_cdfe, _dd("\u0036.\u0031\u002e\u0034\u002d\u0031", "\u006e\u0020\u0061\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0073\u0065c\u0074\u0069\u006f\u006e\u0020h\u0065a\u0064\u0065\u0072\u0020t\u0068\u0065\u0020\u0073\u0074\u0061\u0072t\u0069\u006e\u0067\u0020\u006fb\u006a\u0065\u0063\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0072\u0061n\u0067e\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0053\u0050\u0041C\u0045\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u0028\u0032\u0030\u0068\u0029\u002e"))
	}
	if _ddfd.HasInvalidSeparationAfterXRef() {
		_cdfe = append(_cdfe, _dd("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e"))
	}
	return _cdfe
}

// StandardName gets the name of the standard.
func (_dcac *profile1) StandardName() string {
	return _c.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0031\u0025s", _dcac._efc._dag)
}
func _ccfag(_gdad *_g.CompliancePdfReader) (_eecc []ViolatedRule) {
	var _fcgb, _bgbd, _eaffd, _cegge, _gcdc, _dcgea, _fdgb bool
	_efdd := func() bool { return _fcgb && _bgbd && _eaffd && _cegge && _gcdc && _dcgea && _fdgb }
	_bbdb := func(_gbacd *_cb.PdfObjectDictionary) bool {
		if !_fcgb && _gbacd.Get("\u0054\u0052") != nil {
			_fcgb = true
			_eecc = append(_eecc, _dd("\u0036.\u0032\u002e\u0035\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _ceafa := _gbacd.Get("\u0054\u0052\u0032"); !_bgbd && _ceafa != nil {
			_ccdda, _efde := _cb.GetName(_ceafa)
			if !_efde || (_efde && *_ccdda != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_bgbd = true
				_eecc = append(_eecc, _dd("\u0036.\u0032\u002e\u0035\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _efdd() {
					return true
				}
			}
		}
		if !_eaffd && _gbacd.Get("\u0048\u0054\u0050") != nil {
			_eaffd = true
			_eecc = append(_eecc, _dd("\u0036.\u0032\u002e\u0035\u002d\u0033", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020th\u0065\u0020\u0048\u0054\u0050\u0020\u006b\u0065\u0079\u002e"))
		}
		_cedbd, _egbfb := _cb.GetDict(_gbacd.Get("\u0048\u0054"))
		if _egbfb {
			if _afcg := _cedbd.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u0054\u0079\u0070\u0065"); !_cegge && _afcg != nil {
				_ddaf, _fafeg := _cb.GetInt(_afcg)
				if !_fafeg || (_fafeg && !(*_ddaf == 1 || *_ddaf == 5)) {
					_eecc = append(_eecc, _dd("\u0020\u0036\u002e\u0032\u002e\u0035\u002d\u0034", "\u0041\u006c\u006c\u0020\u0068\u0061\u006c\u0066\u0074\u006f\u006e\u0065\u0073\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006ce\u0020\u0073h\u0061\u006c\u006c\u0020h\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0031\u0020\u006f\u0072\u0020\u0035 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0048\u0061l\u0066\u0074\u006fn\u0065\u0054\u0079\u0070\u0065\u0020\u006be\u0079\u002e"))
					if _efdd() {
						return true
					}
				}
			}
			if _acdf := _cedbd.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u004e\u0061\u006d\u0065"); !_gcdc && _acdf != nil {
				_gcdc = true
				_eecc = append(_eecc, _dd("\u0036.\u0032\u002e\u0035\u002d\u0035", "\u0048\u0061\u006c\u0066\u0074o\u006e\u0065\u0073\u0020\u0069\u006e\u0020a\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0048\u0061\u006c\u0066\u0074\u006f\u006e\u0065N\u0061\u006d\u0065\u0020\u006b\u0065y\u002e"))
				if _efdd() {
					return true
				}
			}
		}
		_, _dfbea := _gdde(_gdad)
		var _cafa bool
		_ebgf, _egbfb := _cb.GetDict(_gbacd.Get("\u0047\u0072\u006fu\u0070"))
		if _egbfb {
			_, _gfaab := _cb.GetName(_ebgf.Get("\u0043\u0053"))
			if _gfaab {
				_cafa = true
			}
		}
		if _cgff := _gbacd.Get("\u0042\u004d"); !_dcgea && !_fdgb && _cgff != nil {
			_bfgfg, _ddca := _cb.GetName(_cgff)
			if _ddca {
				switch _bfgfg.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
				default:
					_dcgea = true
					_eecc = append(_eecc, _dd("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u004f\u006el\u0079\u0020\u0062\u006c\u0065\u006e\u0064\u0020\u006d\u006f\u0064\u0065\u0073\u0020\u0074h\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0073\u0070\u0065c\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0049\u0053O\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a2\u0030\u0030\u0038\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075e\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0042M\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0065\u0078t\u0065\u006e\u0064\u0065\u0064\u0020\u0067\u0072\u0061\u0070\u0068\u0069\u0063\u0020\u0073\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					if _efdd() {
						return true
					}
				}
				if _bfgfg.String() != "\u004e\u006f\u0072\u006d\u0061\u006c" && !_dfbea && !_cafa {
					_fdgb = true
					_eecc = append(_eecc, _dd("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
					if _efdd() {
						return true
					}
				}
			}
		}
		if _, _egbfb = _cb.GetDict(_gbacd.Get("\u0053\u004d\u0061s\u006b")); !_fdgb && _egbfb && !_dfbea && !_cafa {
			_fdgb = true
			_eecc = append(_eecc, _dd("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
			if _efdd() {
				return true
			}
		}
		if _cadd := _gbacd.Get("\u0043\u0041"); !_fdgb && _cadd != nil && !_dfbea && !_cafa {
			_dfdg, _edfce := _cb.GetNumberAsFloat(_cadd)
			if _edfce == nil && _dfdg < 1.0 {
				_fdgb = true
				_eecc = append(_eecc, _dd("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _efdd() {
					return true
				}
			}
		}
		if _fcfde := _gbacd.Get("\u0063\u0061"); !_fdgb && _fcfde != nil && !_dfbea && !_cafa {
			_efgd, _dbbaf := _cb.GetNumberAsFloat(_fcfde)
			if _dbbaf == nil && _efgd < 1.0 {
				_fdgb = true
				_eecc = append(_eecc, _dd("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _efdd() {
					return true
				}
			}
		}
		return false
	}
	for _, _fcdc := range _gdad.PageList {
		_fbaef := _fcdc.Resources
		if _fbaef == nil {
			continue
		}
		if _fbaef.ExtGState == nil {
			continue
		}
		_afcc, _dcbag := _cb.GetDict(_fbaef.ExtGState)
		if !_dcbag {
			continue
		}
		_eece := _afcc.Keys()
		for _, _bcgd := range _eece {
			_ecdca, _gfgf := _cb.GetDict(_afcc.Get(_bcgd))
			if !_gfgf {
				continue
			}
			if _bbdb(_ecdca) {
				return _eecc
			}
		}
	}
	for _, _dfecc := range _gdad.PageList {
		_adcg := _dfecc.Resources
		if _adcg == nil {
			continue
		}
		_cabd, _fbddg := _cb.GetDict(_adcg.XObject)
		if !_fbddg {
			continue
		}
		for _, _bcbb := range _cabd.Keys() {
			_geefg, _edabg := _cb.GetStream(_cabd.Get(_bcbb))
			if !_edabg {
				continue
			}
			_bafga, _edabg := _cb.GetDict(_geefg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_edabg {
				continue
			}
			_fggdg, _edabg := _cb.GetDict(_bafga.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_edabg {
				continue
			}
			for _, _fefbf := range _fggdg.Keys() {
				_fafa, _gccf := _cb.GetDict(_fggdg.Get(_fefbf))
				if !_gccf {
					continue
				}
				if _bbdb(_fafa) {
					return _eecc
				}
			}
		}
	}
	return _eecc
}

type profile2 struct {
	_gaeg standardType
	_afef Profile2Options
}

func _afagd(_febecb *_g.CompliancePdfReader) (_acgf []ViolatedRule) { return _acgf }
func _eaff(_eabbe *_g.PdfFont, _gcgc *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_dacc   = "\u0036.\u0033\u002e\u0035\u002d\u0033"
		_fgcdea = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0073 \u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077i\u0074\u0068\u0069n\u0020\u0061\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0064\u0065s\u0063\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020\u0043\u0049\u0044\u0053\u0065\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0079\u0069\u006eg\u0020\u0077\u0068i\u0063\u0068\u0020\u0043\u0049\u0044\u0073 \u0061\u0072e\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e \u0074\u0068\u0065\u0020\u0065\u006d\u0062\u0065\u0064d\u0065\u0064\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0066\u0069l\u0065,\u0020\u0061\u0073 \u0064\u0065\u0073\u0063\u0072\u0069b\u0065\u0064 \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020\u0054ab\u006c\u0065\u0020\u0035.\u00320\u002e"
	)
	var _cffeg string
	if _egeeg, _bbab := _cb.GetName(_gcgc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bbab {
		_cffeg = _egeeg.String()
	}
	switch _cffeg {
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_agbbc := _eabbe.FontDescriptor()
		if _agbbc.CIDSet == nil {
			return _dd(_dacc, _fgcdea)
		}
		return _dfa
	default:
		return _dfa
	}
}
func (_aa *documentImages) hasUncalibratedImages() bool { return _aa._egc || _aa._eb || _aa._cf }
func _efdgd(_fafe *_g.CompliancePdfReader) (_egbg []ViolatedRule) {
	var _efed, _ccgg bool
	_bffgg := func() bool { return _efed && _ccgg }
	for _, _dbdd := range _fafe.GetObjectNums() {
		_dfdf, _gacc := _fafe.GetIndirectObjectByNumber(_dbdd)
		if _gacc != nil {
			_ge.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _dbdd, _gacc)
			continue
		}
		_cbeg, _ddfdc := _cb.GetDict(_dfdf)
		if !_ddfdc {
			continue
		}
		_fdgcb, _ddfdc := _cb.GetName(_cbeg.Get("\u0054\u0079\u0070\u0065"))
		if !_ddfdc {
			continue
		}
		if *_fdgcb != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_egfbbb, _ddfdc := _cb.GetName(_cbeg.Get("\u0053"))
		if !_ddfdc {
			if !_efed {
				_egbg = append(_egbg, _dd("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_efed = true
				if _bffgg() {
					return _egbg
				}
			}
			continue
		}
		switch _g.PdfActionType(*_egfbbb) {
		case _g.ActionTypeLaunch, _g.ActionTypeSound, _g.ActionTypeMovie, _g.ActionTypeResetForm, _g.ActionTypeImportData, _g.ActionTypeJavaScript:
			if !_efed {
				_egbg = append(_egbg, _dd("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_efed = true
				if _bffgg() {
					return _egbg
				}
			}
			continue
		case _g.ActionTypeNamed:
			if !_ccgg {
				_eaagg, _gdgc := _cb.GetName(_cbeg.Get("\u004e"))
				if !_gdgc {
					_egbg = append(_egbg, _dd("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_ccgg = true
					if _bffgg() {
						return _egbg
					}
					continue
				}
				switch *_eaagg {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_egbg = append(_egbg, _dd("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_ccgg = true
					if _bffgg() {
						return _egbg
					}
					continue
				}
			}
		}
	}
	return _egbg
}
func _gdga(_ddgba *_g.CompliancePdfReader) (_abfca []ViolatedRule) {
	var _daee, _abdf, _cebc, _ccaf, _fcd, _gddc bool
	_ddebb := map[*_cb.PdfObjectStream]struct{}{}
	for _, _afee := range _ddgba.GetObjectNums() {
		if _daee && _abdf && _fcd && _cebc && _ccaf && _gddc {
			return _abfca
		}
		_acbff, _degg := _ddgba.GetIndirectObjectByNumber(_afee)
		if _degg != nil {
			continue
		}
		_efbb, _dced := _cb.GetStream(_acbff)
		if !_dced {
			continue
		}
		if _, _dced = _ddebb[_efbb]; _dced {
			continue
		}
		_ddebb[_efbb] = struct{}{}
		_dfeb, _dced := _cb.GetName(_efbb.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_dced {
			continue
		}
		if !_ccaf {
			if _efbb.Get("\u0052\u0065\u0066") != nil {
				_abfca = append(_abfca, _dd("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_ccaf = true
			}
		}
		if _dfeb.String() == "\u0050\u0053" {
			if !_gddc {
				_abfca = append(_abfca, _dd("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_gddc = true
				continue
			}
		}
		if _dfeb.String() == "\u0046\u006f\u0072\u006d" {
			if _abdf && _cebc && _ccaf {
				continue
			}
			if !_abdf && _efbb.Get("\u004f\u0050\u0049") != nil {
				_abfca = append(_abfca, _dd("\u0036.\u0032\u002e\u0034\u002d\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0028\u0049\u006d\u0061\u0067\u0065\u0020\u006f\u0072\u0020\u0046\u006f\u0072\u006d\u0029\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u004fP\u0049\u0020\u006b\u0065\u0079\u002e"))
				_abdf = true
			}
			if !_cebc {
				if _efbb.Get("\u0050\u0053") != nil {
					_cebc = true
				}
				if _eaad := _efbb.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _eaad != nil && !_cebc {
					if _fbaa, _ceee := _cb.GetName(_eaad); _ceee && *_fbaa == "\u0050\u0053" {
						_cebc = true
					}
				}
				if _cebc {
					_abfca = append(_abfca, _dd("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _dfeb.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_daee && _efbb.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_abfca = append(_abfca, _dd("\u0036.\u0032\u002e\u0034\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_daee = true
		}
		if !_fcd && _efbb.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_bgfg, _ebaf := _cb.GetBool(_efbb.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _ebaf && bool(*_bgfg) {
				continue
			}
			_abfca = append(_abfca, _dd("\u0036.\u0032\u002e\u0034\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_fcd = true
		}
	}
	return _abfca
}
func _caac(_gfac *_g.CompliancePdfReader) (_acdd []ViolatedRule) {
	var _aafdb, _aebe, _eacc, _bafb, _aacg, _cgga, _deceb bool
	_dcaf := func() bool { return _aafdb && _aebe && _eacc && _bafb && _aacg && _cgga && _deceb }
	for _, _febf := range _gfac.PageList {
		_fffga, _febc := _febf.GetAnnotations()
		if _febc != nil {
			_ge.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _febc)
			continue
		}
		for _, _egec := range _fffga {
			if !_aafdb {
				switch _egec.GetContext().(type) {
				case *_g.PdfAnnotationFileAttachment, *_g.PdfAnnotationSound, *_g.PdfAnnotationMovie, nil:
					_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020p\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e\u0020\u0041d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020F\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063\u0068\u006de\u006e\u0074\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u0020\u0061\u006e\u0064\u0020\u004d\u006f\u0076\u0069e\u0020\u0074\u0079\u0070\u0065s \u0073ha\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_aafdb = true
					if _dcaf() {
						return _acdd
					}
				}
			}
			_bfcab, _badf := _cb.GetDict(_egec.GetContainingPdfObject())
			if !_badf {
				continue
			}
			if !_aebe {
				_dggc, _cbf := _cb.GetFloatVal(_bfcab.Get("\u0043\u0041"))
				if _cbf && _dggc != 1.0 {
					_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0033\u002d\u0031", "\u0041\u006e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073h\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0043\u0041\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u002e\u0030\u002e"))
					_aebe = true
					if _dcaf() {
						return _acdd
					}
				}
			}
			if !_eacc {
				_dgdef, _adffa := _cb.GetIntVal(_bfcab.Get("\u0046"))
				if !(_adffa && _dgdef&4 == 4 && _dgdef&1 == 0 && _dgdef&2 == 0 && _dgdef&32 == 0) {
					_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0033\u002d\u0032", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0074\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u002e\u0020\u0054\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065 s\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020I\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u0020\u0061\u006e\u0064\u0020\u004e\u006f\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020b\u0069\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073e\u0074\u0020t\u006f\u0020\u0030\u002e"))
					_eacc = true
					if _dcaf() {
						return _acdd
					}
				}
			}
			if !_bafb {
				_gceee, _gace := _cb.GetDict(_bfcab.Get("\u0041\u0050"))
				if _gace {
					_dddd := _gceee.Get("\u004e")
					if _dddd == nil || len(_gceee.Keys()) > 1 {
						_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_bafb = true
						if _dcaf() {
							return _acdd
						}
						continue
					}
					_, _ceafe := _egec.GetContext().(*_g.PdfAnnotationWidget)
					if _ceafe {
						_gggee, _dgeb := _cb.GetName(_bfcab.Get("\u0046\u0054"))
						if _dgeb && *_gggee == "\u0042\u0074\u006e" {
							if _, _ccef := _cb.GetDict(_dddd); !_ccef {
								_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_bafb = true
								if _dcaf() {
									return _acdd
								}
								continue
							}
						}
					}
					_, _febeg := _cb.GetStream(_dddd)
					if !_febeg {
						_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_bafb = true
						if _dcaf() {
							return _acdd
						}
						continue
					}
				}
			}
			if !_aacg {
				if _bfcab.Get("\u0043") != nil || _bfcab.Get("\u0049\u0043") != nil {
					_gaee, _deffd := _ebga(_gfac)
					if !_deffd {
						_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
						_aacg = true
						if _dcaf() {
							return _acdd
						}
					} else {
						_cdgd, _debe := _cb.GetIntVal(_gaee.Get("\u004e"))
						if !_debe || _cdgd != 3 {
							_acdd = append(_acdd, _dd("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
							_aacg = true
							if _dcaf() {
								return _acdd
							}
						}
					}
				}
			}
			_ddgdg, _bcgb := _egec.GetContext().(*_g.PdfAnnotationWidget)
			if !_bcgb {
				continue
			}
			if !_cgga {
				if _ddgdg.A != nil {
					_acdd = append(_acdd, _dd("\u0036.\u0036\u002e\u0031\u002d\u0033", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_cgga = true
					if _dcaf() {
						return _acdd
					}
				}
			}
			if !_deceb {
				if _ddgdg.AA != nil {
					_acdd = append(_acdd, _dd("\u0036.\u0036\u002e\u0032\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_deceb = true
					if _dcaf() {
						return _acdd
					}
				}
			}
		}
	}
	return _acdd
}
func _bbaae(_ccfb *_g.CompliancePdfReader) (_eagab []ViolatedRule) {
	_ffae := true
	_cacc, _degda := _ccfb.GetCatalogMarkInfo()
	if !_degda {
		_ffae = false
	} else {
		_fgfd, _cda := _cb.GetDict(_cacc)
		if _cda {
			_dbcg, _bdef := _cb.GetBool(_fgfd.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_dbcg) || !_bdef {
				_ffae = false
			}
		} else {
			_ffae = false
		}
	}
	if !_ffae {
		_eagab = append(_eagab, _dd("\u0036.\u0038\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_cbca, _degda := _ccfb.GetCatalogStructTreeRoot()
	if !_degda {
		_eagab = append(_eagab, _dd("\u0036.\u0038\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_cgbg, _degda := _cb.GetDict(_cbca)
	if _degda {
		_bgcb, _fdfac := _cb.GetName(_cgbg.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _fdfac {
			_ffgaee, _acff := _cb.GetDict(_bgcb)
			if _acff {
				for _, _bcbg := range _ffgaee.Keys() {
					_fdeb := _ffgaee.Get(_bcbg)
					if _fdeb == nil {
						_eagab = append(_eagab, _dd("\u0036.\u0038\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _eagab
}
func _ecbg(_gdgb *_g.CompliancePdfReader) (_gdgac []ViolatedRule) {
	var _fbcff, _fbeb, _daece, _abege, _cagg, _gdeed bool
	_bcbe := func() bool { return _fbcff && _fbeb && _daece && _abege && _cagg && _gdeed }
	for _, _fbea := range _gdgb.PageList {
		if _fbea.Resources == nil {
			continue
		}
		_ebab, _eeeed := _cb.GetDict(_fbea.Resources.Font)
		if !_eeeed {
			continue
		}
		for _, _afebc := range _ebab.Keys() {
			_fdcg, _deac := _cb.GetDict(_ebab.Get(_afebc))
			if !_deac {
				if !_fbcff {
					_gdgac = append(_gdgac, _dd("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006e\u0064\u0020\u0066on\u0074 \u0070\u0072\u006fg\u0072\u0061\u006ds\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072mi\u006e\u0067\u0020\u0066\u0069\u006ce\u002c\u0020\u0072\u0065\u0067\u0061\u0072\u0064\u006c\u0065s\u0073\u0020\u006f\u0066\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006eg m\u006f\u0064\u0065\u0020\u0075\u0073\u0061\u0067\u0065\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0074\u0068e\u0020\u0070\u0072o\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0069\u006e \u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031:\u0032\u0030\u0030\u0038\u002c \u0039\u002e\u0036\u0020a\u006e\u0064\u0020\u0039.\u0037\u002e"))
					_fbcff = true
					if _bcbe() {
						return _gdgac
					}
				}
				continue
			}
			if _aefb, _fdbg := _cb.GetName(_fdcg.Get("\u0054\u0079\u0070\u0065")); !_fbcff && (!_fdbg || _aefb.String() != "\u0046\u006f\u006e\u0074") {
				_gdgac = append(_gdgac, _dd("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_fbcff = true
				if _bcbe() {
					return _gdgac
				}
			}
			_beeca, _acffb := _g.NewPdfFontFromPdfObject(_fdcg)
			if _acffb != nil {
				continue
			}
			var _gbgf string
			if _cbefc, _febec := _cb.GetName(_fdcg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _febec {
				_gbgf = _cbefc.String()
			}
			if !_fbeb {
				switch _gbgf {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_fbeb = true
					_gdgac = append(_gdgac, _dd("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _bcbe() {
						return _gdgac
					}
				}
			}
			if !_daece {
				if _gbgf != "\u0054\u0079\u0070e\u0033" {
					_cfge, _adfc := _cb.GetName(_fdcg.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_adfc || _cfge.String() == "" {
						_gdgac = append(_gdgac, _dd("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_daece = true
						if _bcbe() {
							return _gdgac
						}
					}
				}
			}
			if _gbgf != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_effc := _ae.IsStdFont(_ae.StdFontName(_beeca.BaseFont()))
			if _effc {
				continue
			}
			_ggeb, _beab := _cb.GetIntVal(_fdcg.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_beab && !_abege {
				_gdgac = append(_gdgac, _dd("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_abege = true
				if _bcbe() {
					return _gdgac
				}
			}
			_dgdd, _ecbfa := _cb.GetIntVal(_fdcg.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_ecbfa && !_cagg {
				_gdgac = append(_gdgac, _dd("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_cagg = true
				if _bcbe() {
					return _gdgac
				}
			}
			if !_gdeed {
				_fadg, _gefb := _cb.GetArray(_fdcg.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_gefb || !_beab || !_ecbfa || _fadg.Len() != _dgdd-_ggeb+1 {
					_gdgac = append(_gdgac, _dd("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_gdeed = true
					if _bcbe() {
						return _gdgac
					}
				}
			}
		}
	}
	return _gdgac
}
func _ebaff(_fabg *_g.CompliancePdfReader) (_bbcag ViolatedRule) {
	for _, _gegb := range _fabg.GetObjectNums() {
		_ccdbd, _bdad := _fabg.GetIndirectObjectByNumber(_gegb)
		if _bdad != nil {
			continue
		}
		_dccg, _fedag := _cb.GetStream(_ccdbd)
		if !_fedag {
			continue
		}
		_cgdg, _fedag := _cb.GetName(_dccg.Get("\u0054\u0079\u0070\u0065"))
		if !_fedag {
			continue
		}
		if *_cgdg != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_, _fedag = _cb.GetName(_dccg.Get("\u004f\u0050\u0049"))
		if _fedag {
			return _dd("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		_afff, _fedag := _cb.GetName(_dccg.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_fedag {
			continue
		}
		if *_afff == "\u0050\u0053" {
			return _dd("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _dccg.Get("\u0050\u0053") != nil {
			return _dd("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _bbcag
}
func _efae(_fcfc *_g.CompliancePdfReader) (_cddd []ViolatedRule) {
	if _fcfc.ParserMetadata().HasOddLengthHexStrings() {
		_cddd = append(_cddd, _dd("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _fcfc.ParserMetadata().HasOddLengthHexStrings() {
		_cddd = append(_cddd, _dd("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _cddd
}

type profile1 struct {
	_efc  standardType
	_gfcf Profile1Options
}

func _dbdbe(_cecg *_g.CompliancePdfReader) (_ccac []ViolatedRule) {
	_aba := _cecg.GetObjectNums()
	for _, _dda := range _aba {
		_cac, _begef := _cecg.GetIndirectObjectByNumber(_dda)
		if _begef != nil {
			continue
		}
		_dcgbb, _ecaa := _cb.GetDict(_cac)
		if !_ecaa {
			continue
		}
		_efg, _ecaa := _cb.GetName(_dcgbb.Get("\u0054\u0079\u0070\u0065"))
		if !_ecaa {
			continue
		}
		if _efg.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _dcgbb.Get("\u0045\u0046") != nil {
			_ccac = append(_ccac, _dd("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0031", "\u0041 \u0066\u0069\u006c\u0065 \u0073p\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046 \u0033\u002e\u0031\u0030\u002e\u0032\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0045\u0046 \u006be\u0079\u002e"))
			break
		}
	}
	_cdcc, _decda := _addf(_cecg)
	if !_decda {
		return _ccac
	}
	_gedf, _decda := _cb.GetDict(_cdcc.Get("\u004e\u0061\u006de\u0073"))
	if !_decda {
		return _ccac
	}
	if _gedf.Get("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065\u0073") != nil {
		_ccac = append(_ccac, _dd("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0032", "\u0041\u0020\u0066i\u006c\u0065\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020d\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020PD\u0046 \u0052\u0065\u0066er\u0065\u006e\u0063\u0065\u0020\u0033\u002e6\u002e\u0033\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0045m\u0062\u0065\u0064\u0064\u0065\u0064\u0046i\u006c\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
	}
	return _ccac
}

// Profile1B is the implementation of the PDF/A-1B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1B struct{ profile1 }
