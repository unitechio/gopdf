package pdfa

import (
	_ea "errors"
	_b "fmt"
	_e "image/color"
	_bf "math"
	_a "sort"
	_g "strings"
	_c "time"

	_eg "unitechio/gopdf/gopdf/common"
	_df "unitechio/gopdf/gopdf/contentstream"
	_cb "unitechio/gopdf/gopdf/core"
	_cd "unitechio/gopdf/gopdf/internal/cmap"
	_dbg "unitechio/gopdf/gopdf/internal/imageutil"
	_egd "unitechio/gopdf/gopdf/internal/timeutils"
	_db "unitechio/gopdf/gopdf/model"
	_ebb "unitechio/gopdf/gopdf/model/internal/colorprofile"
	_f "unitechio/gopdf/gopdf/model/internal/docutil"
	_acd "unitechio/gopdf/gopdf/model/internal/fonts"
	_fcg "unitechio/gopdf/gopdf/model/xmputil"
	_eee "unitechio/gopdf/gopdf/model/xmputil/pdfaextension"
	_fc "unitechio/gopdf/gopdf/model/xmputil/pdfaid"
	_ee "github.com/adrg/sysfont"
	_da "github.com/trimmer-io/go-xmp/models/dc"
	_bfg "github.com/trimmer-io/go-xmp/models/pdf"
	_gf "github.com/trimmer-io/go-xmp/models/xmp_base"
	_ac "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_bc "github.com/trimmer-io/go-xmp/models/xmp_rights"
	_eb "github.com/trimmer-io/go-xmp/xmp"
)

func _dcab(_gbbg *_db.CompliancePdfReader) ViolatedRule            { return _ce }
func _eddd(_fdbed *_db.CompliancePdfReader) (_ebfd []ViolatedRule) { return _ebfd }
func _be() standardType                                            { return standardType{_ed: 2, _fd: "\u0055"} }

// NewProfile2U creates a new Profile2U with the given options.
func NewProfile2U(options *Profile2Options) *Profile2U {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_bdce(options)
	return &Profile2U{profile2{_fdbb: *options, _dgfgd: _be()}}
}

func _cg(_bbg []*_f.Image, _ead bool) error {
	_afd := _cb.PdfObjectName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
	if _ead {
		_afd = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	for _, _dab := range _bbg {
		if _dab.Colorspace == _afd {
			continue
		}
		_cda, _ggg := _db.NewXObjectImageFromStream(_dab.Stream)
		if _ggg != nil {
			return _ggg
		}
		_bbb, _ggg := _cda.ToImage()
		if _ggg != nil {
			return _ggg
		}
		_dd, _ggg := _bbb.ToGoImage()
		if _ggg != nil {
			return _ggg
		}
		var _gba _db.PdfColorspace
		if _ead {
			_gba = _db.NewPdfColorspaceDeviceCMYK()
			_dd, _ggg = _dbg.CMYKConverter.Convert(_dd)
		} else {
			_gba = _db.NewPdfColorspaceDeviceRGB()
			_dd, _ggg = _dbg.NRGBAConverter.Convert(_dd)
		}
		if _ggg != nil {
			return _ggg
		}
		_acb, _dg := _dd.(_dbg.Image)
		if !_dg {
			return _ea.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
		}
		_dbb := _acb.Base()
		_gcg := &_db.Image{Width: int64(_dbb.Width), Height: int64(_dbb.Height), BitsPerComponent: int64(_dbb.BitsPerComponent), ColorComponents: _dbb.ColorComponents, Data: _dbb.Data}
		_gcg.SetDecode(_dbb.Decode)
		_gcg.SetAlpha(_dbb.Alpha)
		if _ggg = _cda.SetImage(_gcg, _gba); _ggg != nil {
			return _ggg
		}
		_cda.ToPdfObject()
		_dab.ColorComponents = _dbb.ColorComponents
		_dab.Colorspace = _afd
	}
	return nil
}

// Part gets the PDF/A version level.
func (_ebc *profile1) Part() int { return _ebc._geg._ed }

// VerificationError is the PDF/A verification error structure, that contains all violated rules.
type VerificationError struct {
	// ViolatedRules are the rules that were violated during error verification.
	ViolatedRules []ViolatedRule

	// ConformanceLevel defines the standard on verification failed.
	ConformanceLevel int

	// ConformanceVariant is the standard variant used on verification.
	ConformanceVariant string
}

func _bgcc(_fgdb bool, _gcc standardType) (pageColorspaceOptimizeFunc, documentColorspaceOptimizeFunc) {
	var _cdda, _gfc, _bcgc bool
	_aaad := func(_aca *_f.Document, _cecc *_f.Page, _aced []*_f.Image) error {
		for _, _beg := range _aced {
			switch _beg.Colorspace {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				_gfc = true
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				_cdda = true
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				_bcgc = true
			}
		}
		_adc, _bdf := _cecc.GetContents()
		if !_bdf {
			return nil
		}
		for _, _ecb := range _adc {
			_edad, _deda := _ecb.GetData()
			if _deda != nil {
				continue
			}
			_edbc := _df.NewContentStreamParser(string(_edad))
			_fccf, _deda := _edbc.Parse()
			if _deda != nil {
				continue
			}
			for _, _bagf := range *_fccf {
				switch _bagf.Operand {
				case "\u0047", "\u0067":
					_gfc = true
				case "\u0052\u0047", "\u0072\u0067":
					_cdda = true
				case "\u004b", "\u006b":
					_bcgc = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_bagf.Params) == 0 {
						continue
					}
					_adce, _efbe := _cb.GetName(_bagf.Params[0])
					if !_efbe {
						continue
					}
					switch _adce.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_cdda = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_gfc = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_bcgc = true
					}
				}
			}
		}
		_eafb := _cecc.FindXObjectForms()
		for _, _ebbbe := range _eafb {
			_ceda := _df.NewContentStreamParser(string(_ebbbe.Stream))
			_ggbb, _def := _ceda.Parse()
			if _def != nil {
				continue
			}
			for _, _eeag := range *_ggbb {
				switch _eeag.Operand {
				case "\u0047", "\u0067":
					_gfc = true
				case "\u0052\u0047", "\u0072\u0067":
					_cdda = true
				case "\u004b", "\u006b":
					_bcgc = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_eeag.Params) == 0 {
						continue
					}
					_cgb, _gcbb := _cb.GetName(_eeag.Params[0])
					if !_gcbb {
						continue
					}
					switch _cgb.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_cdda = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_gfc = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_bcgc = true
					}
				}
			}
			_agba, _cddb := _cb.GetArray(_cecc.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
			if !_cddb {
				return nil
			}
			for _, _bebfd := range _agba.Elements() {
				_cge, _dabf := _cb.GetDict(_bebfd)
				if !_dabf {
					continue
				}
				_edba := _cge.Get("\u0043")
				if _edba == nil {
					continue
				}
				_dddf, _dabf := _cb.GetArray(_edba)
				if !_dabf {
					continue
				}
				switch _dddf.Len() {
				case 0:
				case 1:
					_gfc = true
				case 3:
					_cdda = true
				case 4:
					_bcgc = true
				}
			}
		}
		return nil
	}
	_feeg := func(_gfagf *_f.Document, _bgge []*_f.Image) error {
		_acbb, _egde := _gfagf.FindCatalog()
		if !_egde {
			return nil
		}
		_gfagb, _egde := _acbb.GetOutputIntents()
		if _egde && _gfagb.Len() > 0 {
			return nil
		}
		if !_egde {
			_gfagb = _acbb.NewOutputIntents()
		}
		if !(_cdda || _bcgc || _gfc) {
			return nil
		}
		defer _acbb.SetOutputIntents(_gfagb)
		if _cdda && !_bcgc && !_gfc {
			return _bgd(_gfagf, _gcc, _gfagb)
		}
		if _bcgc && !_cdda && !_gfc {
			return _ebed(_gcc, _gfagb)
		}
		if _gfc && !_cdda && !_bcgc {
			return _fge(_gcc, _gfagb)
		}
		if _cdda && _bcgc {
			if _aade := _cg(_bgge, _fgdb); _aade != nil {
				return _aade
			}
			if _fcdg := _dbeg(_gfagf, _fgdb); _fcdg != nil {
				return _fcdg
			}
			if _daac := _cdge(_gfagf, _fgdb); _daac != nil {
				return _daac
			}
			if _bdg := _ged(_gfagf, _fgdb); _bdg != nil {
				return _bdg
			}
			if _fgdb {
				return _ebed(_gcc, _gfagb)
			}
			return _bgd(_gfagf, _gcc, _gfagb)
		}
		return nil
	}
	return _aaad, _feeg
}
func _cede(_efcg *_db.CompliancePdfReader) ViolatedRule { return _ce }
func _befb(_bagbg *_cb.PdfObjectDictionary, _acead map[*_cb.PdfObjectStream][]byte, _dcef map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_ddcg = "\u0036.\u0033\u002e\u0038\u002d\u0031"
		_dcca = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_fbfdc, _gegd := _cb.GetStream(_bagbg.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _gegd {
		_, _afaf := _feac(_fbfdc, _acead, _dcef)
		if _afaf != nil {
			return _fdbe(_ddcg, _dcca)
		}
		return _ce
	}
	_gacfg, _gegd := _cb.GetName(_bagbg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_gegd {
		return _fdbe(_ddcg, _dcca)
	}
	switch _gacfg.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _ce
	}
	return _fdbe(_ddcg, _dcca)
}

func _dbf(_cec []_cb.PdfObject) (*documentImages, error) {
	_cdc := _cb.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_ebbb := make(map[*_cb.PdfObjectStream]struct{})
	_acg := make(map[_cb.PdfObject]struct{})
	var (
		_ffa, _fe, _cae bool
		_ga             []*imageInfo
		_eea            error
	)
	for _, _gg := range _cec {
		_adb, _fef := _cb.GetStream(_gg)
		if !_fef {
			continue
		}
		if _, _cee := _ebbb[_adb]; _cee {
			continue
		}
		_ebbb[_adb] = struct{}{}
		_ffg := _adb.PdfObjectDictionary.Get(_cdc)
		_de, _fef := _cb.GetName(_ffg)
		if !_fef || string(*_de) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if _faf := _adb.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b"); _faf != nil {
			_acg[_faf] = struct{}{}
		}
		_bb := imageInfo{BitsPerComponent: 8, Stream: _adb}
		_bb.ColorSpace, _eea = _db.DetermineColorspaceNameFromPdfObject(_adb.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _eea != nil {
			return nil, _eea
		}
		if _bd, _eeef := _cb.GetIntVal(_adb.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _eeef {
			_bb.BitsPerComponent = _bd
		}
		if _bda, _fgd := _cb.GetIntVal(_adb.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _fgd {
			_bb.Width = _bda
		}
		if _bag, _bcg := _cb.GetIntVal(_adb.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _bcg {
			_bb.Height = _bag
		}
		switch _bb.ColorSpace {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_cae = true
			_bb.ColorComponents = 1
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_ffa = true
			_bb.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_fe = true
			_bb.ColorComponents = 4
		default:
			_bb._fg = true
		}
		_ga = append(_ga, &_bb)
	}
	if len(_acg) > 0 {
		if len(_acg) == len(_ga) {
			_ga = nil
		} else {
			_bfe := make([]*imageInfo, len(_ga)-len(_acg))
			var _gc int
			for _, _bbe := range _ga {
				if _, _af := _acg[_bbe.Stream]; _af {
					continue
				}
				_bfe[_gc] = _bbe
				_gc++
			}
			_ga = _bfe
		}
	}
	return &documentImages{_acda: _ffa, _dbd: _fe, _fa: _cae, _ca: _acg, _bee: _ga}, nil
}

// StandardName gets the name of the standard.
func (_edaa *profile1) StandardName() string {
	return _b.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0031\u0025s", _edaa._geg._fd)
}

func _fde(_cgde *_db.PdfPageResources, _cgdg *_df.ContentStreamOperations, _cfbb bool) ([]byte, error) {
	var _dccg bool
	for _, _gfaec := range *_cgdg {
	_gca:
		switch _gfaec.Operand {
		case "\u0042\u0049":
			_ddce, _fead := _gfaec.Params[0].(*_df.ContentStreamInlineImage)
			if !_fead {
				break
			}
			_fab, _dga := _ddce.GetColorSpace(_cgde)
			if _dga != nil {
				return nil, _dga
			}
			switch _fab.(type) {
			case *_db.PdfColorspaceDeviceCMYK:
				if _cfbb {
					break _gca
				}
			case *_db.PdfColorspaceDeviceGray:
			case *_db.PdfColorspaceDeviceRGB:
				if !_cfbb {
					break _gca
				}
			default:
				break _gca
			}
			_dccg = true
			_ffge, _dga := _ddce.ToImage(_cgde)
			if _dga != nil {
				return nil, _dga
			}
			_bff, _dga := _ffge.ToGoImage()
			if _dga != nil {
				return nil, _dga
			}
			if _cfbb {
				_bff, _dga = _dbg.CMYKConverter.Convert(_bff)
			} else {
				_bff, _dga = _dbg.NRGBAConverter.Convert(_bff)
			}
			if _dga != nil {
				return nil, _dga
			}
			_bdd, _fead := _bff.(_dbg.Image)
			if !_fead {
				return nil, _ea.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
			}
			_aafg := _bdd.Base()
			_fced := _db.Image{Width: int64(_aafg.Width), Height: int64(_aafg.Height), BitsPerComponent: int64(_aafg.BitsPerComponent), ColorComponents: _aafg.ColorComponents, Data: _aafg.Data}
			_fced.SetDecode(_aafg.Decode)
			_fced.SetAlpha(_aafg.Alpha)
			_cbd, _dga := _ddce.GetEncoder()
			if _dga != nil {
				_cbd = _cb.NewFlateEncoder()
			}
			_dgea, _dga := _df.NewInlineImageFromImage(_fced, _cbd)
			if _dga != nil {
				return nil, _dga
			}
			_gfaec.Params[0] = _dgea
		case "\u0047", "\u0067":
			if len(_gfaec.Params) != 1 {
				break
			}
			_gadc, _cgba := _cb.GetNumberAsFloat(_gfaec.Params[0])
			if _cgba != nil {
				break
			}
			if _cfbb {
				_gfaec.Params = []_cb.PdfObject{_cb.MakeFloat(0), _cb.MakeFloat(0), _cb.MakeFloat(0), _cb.MakeFloat(1 - _gadc)}
				_efc := "\u004b"
				if _gfaec.Operand == "\u0067" {
					_efc = "\u006b"
				}
				_gfaec.Operand = _efc
			} else {
				_gfaec.Params = []_cb.PdfObject{_cb.MakeFloat(_gadc), _cb.MakeFloat(_gadc), _cb.MakeFloat(_gadc)}
				_dgfd := "\u0052\u0047"
				if _gfaec.Operand == "\u0067" {
					_dgfd = "\u0072\u0067"
				}
				_gfaec.Operand = _dgfd
			}
			_dccg = true
		case "\u0052\u0047", "\u0072\u0067":
			if !_cfbb {
				break
			}
			if len(_gfaec.Params) != 3 {
				break
			}
			_ffc, _cab := _cb.GetNumbersAsFloat(_gfaec.Params)
			if _cab != nil {
				break
			}
			_dccg = true
			_facf, _adfb, _bfgg := _ffc[0], _ffc[1], _ffc[2]
			_eaa, _acee, _fdag, _eff := _e.RGBToCMYK(uint8(_facf*255), uint8(_adfb*255), uint8(255*_bfgg))
			_gfaec.Params = []_cb.PdfObject{_cb.MakeFloat(float64(_eaa) / 255), _cb.MakeFloat(float64(_acee) / 255), _cb.MakeFloat(float64(_fdag) / 255), _cb.MakeFloat(float64(_eff) / 255)}
			_cgdeg := "\u004b"
			if _gfaec.Operand == "\u0072\u0067" {
				_cgdeg = "\u006b"
			}
			_gfaec.Operand = _cgdeg
		case "\u004b", "\u006b":
			if _cfbb {
				break
			}
			if len(_gfaec.Params) != 4 {
				break
			}
			_cfbe, _gbed := _cb.GetNumbersAsFloat(_gfaec.Params)
			if _gbed != nil {
				break
			}
			_agd, _cedb, _caaf, _aceed := _cfbe[0], _cfbe[1], _cfbe[2], _cfbe[3]
			_gac, _aadf, _dccf := _e.CMYKToRGB(uint8(255*_agd), uint8(255*_cedb), uint8(255*_caaf), uint8(255*_aceed))
			_gfaec.Params = []_cb.PdfObject{_cb.MakeFloat(float64(_gac) / 255), _cb.MakeFloat(float64(_aadf) / 255), _cb.MakeFloat(float64(_dccf) / 255)}
			_ega := "\u0052\u0047"
			if _gfaec.Operand == "\u006b" {
				_ega = "\u0072\u0067"
			}
			_gfaec.Operand = _ega
			_dccg = true
		}
	}
	if !_dccg {
		return nil, nil
	}
	_fabb := _df.NewContentCreator()
	for _, _gde := range *_cgdg {
		_fabb.AddOperand(*_gde)
	}
	_dfbg := _fabb.Bytes()
	return _dfbg, nil
}
func _cgaag(_gadf *_db.CompliancePdfReader) []ViolatedRule { return nil }

// Conformance gets the PDF/A conformance.
func (_baeb *profile2) Conformance() string { return _baeb._dgfgd._fd }

func _dcd(_cfd *_f.Document, _fgf int) {
	if _cfd.Version.Major == 0 {
		_cfd.Version.Major = 1
	}
	if _cfd.Version.Minor < _fgf {
		_cfd.Version.Minor = _fgf
	}
}

var _ Profile = (*Profile1A)(nil)

func _cfb(_cfbf *_f.Document, _eega func() _c.Time) error {
	_ccb, _ebdc := _db.NewPdfInfoFromObject(_cfbf.Info)
	if _ebdc != nil {
		return _ebdc
	}
	if _fcbdf := _cgeb(_ccb, _eega); _fcbdf != nil {
		return _fcbdf
	}
	_cfbf.Info = _ccb.ToPdfObject()
	return nil
}

// ApplyStandard tries to change the content of the writer to match the PDF/A-1 standard.
// Implements model.StandardApplier.
func (_ageg *profile1) ApplyStandard(document *_f.Document) (_bfad error) {
	_dcd(document, 4)
	if _bfad = _cfb(document, _ageg._fbg.Now); _bfad != nil {
		return _bfad
	}
	if _bfad = _ecad(document); _bfad != nil {
		return _bfad
	}
	_gaec, _dfa := _bgcc(_ageg._fbg.CMYKDefaultColorSpace, _ageg._geg)
	_bfad = _eef(document, []pageColorspaceOptimizeFunc{_bagb, _gaec}, []documentColorspaceOptimizeFunc{_dfa})
	if _bfad != nil {
		return _bfad
	}
	_cgaa(document)
	if _bfad = _daf(document, _ageg._geg._ed); _bfad != nil {
		return _bfad
	}
	if _bfad = _baea(document); _bfad != nil {
		return _bfad
	}
	if _bfad = _dce(document); _bfad != nil {
		return _bfad
	}
	if _bfad = _ggb(document); _bfad != nil {
		return _bfad
	}
	if _bfad = _cdf(document); _bfad != nil {
		return _bfad
	}
	if _ageg._geg._fd == "\u0041" {
		_cbce(document)
	}
	if _bfad = _fgg(document, _ageg._geg._ed); _bfad != nil {
		return _bfad
	}
	if _bfad = _gbc(document); _bfad != nil {
		return _bfad
	}
	if _efeg := _fed(document, _ageg._geg, _ageg._fbg.Xmp); _efeg != nil {
		return _efeg
	}
	if _ageg._geg == _fdc() {
		if _bfad = _gd(document); _bfad != nil {
			return _bfad
		}
	}
	if _bfad = _dag(document); _bfad != nil {
		return _bfad
	}
	return nil
}

func _ddcf(_gfbe *_db.CompliancePdfReader) (_deae []ViolatedRule) {
	var _bfafc, _ade, _dfdc, _dgff, _bfcc, _cdaa bool
	_adab := map[*_cb.PdfObjectStream]struct{}{}
	for _, _aceef := range _gfbe.GetObjectNums() {
		if _bfafc && _ade && _bfcc && _dfdc && _dgff && _cdaa {
			return _deae
		}
		_bgaa, _fegaa := _gfbe.GetIndirectObjectByNumber(_aceef)
		if _fegaa != nil {
			continue
		}
		_afcd, _gbdba := _cb.GetStream(_bgaa)
		if !_gbdba {
			continue
		}
		if _, _gbdba = _adab[_afcd]; _gbdba {
			continue
		}
		_adab[_afcd] = struct{}{}
		_gbbad, _gbdba := _cb.GetName(_afcd.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_gbdba {
			continue
		}
		if !_dgff {
			if _afcd.Get("\u0052\u0065\u0066") != nil {
				_deae = append(_deae, _fdbe("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_dgff = true
			}
		}
		if _gbbad.String() == "\u0050\u0053" {
			if !_cdaa {
				_deae = append(_deae, _fdbe("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_cdaa = true
				continue
			}
		}
		if _gbbad.String() == "\u0046\u006f\u0072\u006d" {
			if _ade && _dfdc && _dgff {
				continue
			}
			if !_ade && _afcd.Get("\u004f\u0050\u0049") != nil {
				_deae = append(_deae, _fdbe("\u0036.\u0032\u002e\u0034\u002d\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0028\u0049\u006d\u0061\u0067\u0065\u0020\u006f\u0072\u0020\u0046\u006f\u0072\u006d\u0029\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u004fP\u0049\u0020\u006b\u0065\u0079\u002e"))
				_ade = true
			}
			if !_dfdc {
				if _afcd.Get("\u0050\u0053") != nil {
					_dfdc = true
				}
				if _gdad := _afcd.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _gdad != nil && !_dfdc {
					if _aec, _gaedg := _cb.GetName(_gdad); _gaedg && *_aec == "\u0050\u0053" {
						_dfdc = true
					}
				}
				if _dfdc {
					_deae = append(_deae, _fdbe("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _gbbad.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_bfafc && _afcd.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_deae = append(_deae, _fdbe("\u0036.\u0032\u002e\u0034\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_bfafc = true
		}
		if !_bfcc && _afcd.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_efde, _acgb := _cb.GetBool(_afcd.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _acgb && bool(*_efde) {
				continue
			}
			_deae = append(_deae, _fdbe("\u0036.\u0032\u002e\u0034\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_bfcc = true
		}
	}
	return _deae
}
func (_fag *documentImages) hasOnlyDeviceCMYK() bool { return _fag._dbd && !_fag._acda && !_fag._fa }
func _cbce(_fgeac *_f.Document) {
	_ebef, _fbgfg := _fgeac.FindCatalog()
	if !_fbgfg {
		return
	}
	_bcee, _fbgfg := _ebef.GetMarkInfo()
	if !_fbgfg {
		_bcee = _cb.MakeDict()
	}
	_bacb, _fbgfg := _cb.GetBool(_bcee.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
	if !_fbgfg || !bool(*_bacb) {
		_bcee.Set("\u004d\u0061\u0072\u006b\u0065\u0064", _cb.MakeBool(true))
		_ebef.SetMarkInfo(_bcee)
	}
}

func _efff(_abaed *_db.CompliancePdfReader) (_accd ViolatedRule) {
	_bbbe, _becbg := _eagdc(_abaed)
	if !_becbg {
		return _ce
	}
	_ccfdb, _becbg := _cb.GetDict(_bbbe.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_becbg {
		return _ce
	}
	_edgec, _becbg := _cb.GetArray(_ccfdb.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_becbg {
		return _ce
	}
	for _gagf := 0; _gagf < _edgec.Len(); _gagf++ {
		_dafb, _edfe := _cb.GetDict(_edgec.Get(_gagf))
		if !_edfe {
			continue
		}
		if _dafb.Get("\u0041") != nil {
			return _fdbe("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
		if _dafb.Get("\u0041\u0041") != nil {
			return _fdbe("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
	}
	return _ce
}

func _bbfe(_cbfg *_db.CompliancePdfReader) (_gfbf []ViolatedRule) {
	var _aag, _bffe bool
	_dbacc := func() bool { return _aag && _bffe }
	for _, _beea := range _cbfg.GetObjectNums() {
		_efea, _dfdf := _cbfg.GetIndirectObjectByNumber(_beea)
		if _dfdf != nil {
			_eg.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _beea, _dfdf)
			continue
		}
		_acgca, _afdf := _cb.GetDict(_efea)
		if !_afdf {
			continue
		}
		_dbcfb, _afdf := _cb.GetName(_acgca.Get("\u0054\u0079\u0070\u0065"))
		if !_afdf {
			continue
		}
		if *_dbcfb != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_dadd, _afdf := _cb.GetName(_acgca.Get("\u0053"))
		if !_afdf {
			if !_aag {
				_gfbf = append(_gfbf, _fdbe("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_aag = true
				if _dbacc() {
					return _gfbf
				}
			}
			continue
		}
		switch _db.PdfActionType(*_dadd) {
		case _db.ActionTypeLaunch, _db.ActionTypeSound, _db.ActionTypeMovie, _db.ActionTypeResetForm, _db.ActionTypeImportData, _db.ActionTypeJavaScript:
			if !_aag {
				_gfbf = append(_gfbf, _fdbe("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_aag = true
				if _dbacc() {
					return _gfbf
				}
			}
			continue
		case _db.ActionTypeNamed:
			if !_bffe {
				_gdde, _dgbg := _cb.GetName(_acgca.Get("\u004e"))
				if !_dgbg {
					_gfbf = append(_gfbf, _fdbe("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bffe = true
					if _dbacc() {
						return _gfbf
					}
					continue
				}
				switch *_gdde {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_gfbf = append(_gfbf, _fdbe("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bffe = true
					if _dbacc() {
						return _gfbf
					}
					continue
				}
			}
		}
	}
	return _gfbf
}
func (_aa *documentImages) hasUncalibratedImages() bool { return _aa._acda || _aa._dbd || _aa._fa }
func _ccac(_aace *_db.CompliancePdfReader) (_dbdbf []ViolatedRule) {
	var _effa, _dgb, _dggb, _bcbc, _aafe, _cddd bool
	_fgee := func() bool { return _effa && _dgb && _dggb && _bcbc && _aafe && _cddd }
	_bbga := func(_dae *_cb.PdfObjectDictionary) bool {
		if !_effa && _dae.Get("\u0054\u0052") != nil {
			_effa = true
			_dbdbf = append(_dbdbf, _fdbe("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _ggccc := _dae.Get("\u0054\u0052\u0032"); !_dgb && _ggccc != nil {
			_bbdbc, _gdbd := _cb.GetName(_ggccc)
			if !_gdbd || (_gdbd && *_bbdbc != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_dgb = true
				_dbdbf = append(_dbdbf, _fdbe("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _fgee() {
					return true
				}
			}
		}
		if _bbge := _dae.Get("\u0053\u004d\u0061s\u006b"); !_dggb && _bbge != nil {
			_dcbe, _dade := _cb.GetName(_bbge)
			if !_dade || (_dade && *_dcbe != "\u004e\u006f\u006e\u0065") {
				_dggb = true
				_dbdbf = append(_dbdbf, _fdbe("\u0036\u002e\u0034-\u0031", "\u0049\u0066\u0020\u0061\u006e \u0053\u004d\u0061\u0073\u006b\u0020\u006be\u0079\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0069\u0074s\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u004e\u006f\u006ee\u002e"))
				if _fgee() {
					return true
				}
			}
		}
		if _cbbf := _dae.Get("\u0043\u0041"); !_aafe && _cbbf != nil {
			_cadab, _adcc := _cb.GetNumberAsFloat(_cbbf)
			if _adcc == nil && _cadab != 1.0 {
				_aafe = true
				_dbdbf = append(_dbdbf, _fdbe("\u0036\u002e\u0034-\u0035", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0043\u0041 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _fgee() {
					return true
				}
			}
		}
		if _edae := _dae.Get("\u0063\u0061"); !_cddd && _edae != nil {
			_gegb, _gagg := _cb.GetNumberAsFloat(_edae)
			if _gagg == nil && _gegb != 1.0 {
				_cddd = true
				_dbdbf = append(_dbdbf, _fdbe("\u0036\u002e\u0034-\u0036", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0063\u0061 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _fgee() {
					return true
				}
			}
		}
		if _dfff := _dae.Get("\u0042\u004d"); !_bcbc && _dfff != nil {
			_aegcb, _ebfec := _cb.GetName(_dfff)
			if _ebfec {
				switch _aegcb.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
				default:
					_bcbc = true
					_dbdbf = append(_dbdbf, _fdbe("\u0036\u002e\u0034-\u0034", "T\u0068\u0065\u0020\u0066\u006f\u006cl\u006f\u0077\u0069\u006e\u0067 \u006b\u0065y\u0073\u002c\u0020\u0069\u0066 \u0070res\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047S\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065 \u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0077n\u003a\u0020\u0042\u004d\u0020\u002d\u0020\u004e\u006f\u0072m\u0061\u006c\u0020\u006f\u0072\u0020\u0043\u006f\u006d\u0070\u0061t\u0069\u0062\u006c\u0065\u002e"))
					if _fgee() {
						return true
					}
				}
			}
		}
		return false
	}
	for _, _dbbg := range _aace.PageList {
		_bcfd := _dbbg.Resources
		if _bcfd == nil {
			continue
		}
		if _bcfd.ExtGState == nil {
			continue
		}
		_dgag, _ecbbg := _cb.GetDict(_bcfd.ExtGState)
		if !_ecbbg {
			continue
		}
		_fagdg := _dgag.Keys()
		for _, _abfe := range _fagdg {
			_cagb, _eebb := _cb.GetDict(_dgag.Get(_abfe))
			if !_eebb {
				continue
			}
			if _bbga(_cagb) {
				return _dbdbf
			}
		}
	}
	for _, _gffc := range _aace.PageList {
		_gbae := _gffc.Resources
		if _gbae == nil {
			continue
		}
		_fbagb, _beba := _cb.GetDict(_gbae.XObject)
		if !_beba {
			continue
		}
		for _, _agdgd := range _fbagb.Keys() {
			_cafgb, _cggc := _cb.GetStream(_fbagb.Get(_agdgd))
			if !_cggc {
				continue
			}
			_gaee, _cggc := _cb.GetDict(_cafgb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_cggc {
				continue
			}
			_ceffd, _cggc := _cb.GetDict(_gaee.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_cggc {
				continue
			}
			for _, _bfbac := range _ceffd.Keys() {
				_aegbg, _bdaf := _cb.GetDict(_ceffd.Get(_bfbac))
				if !_bdaf {
					continue
				}
				if _bbga(_aegbg) {
					return _dbdbf
				}
			}
		}
	}
	return _dbdbf
}

// Conformance gets the PDF/A conformance.
func (_cfdb *profile1) Conformance() string { return _cfdb._geg._fd }

func _fge(_beeff standardType, _bege *_f.OutputIntents) error {
	_dabg, _gbdb := _ebb.NewISOCoatedV2Gray1CBasOutputIntent(_beeff.outputIntentSubtype())
	if _gbdb != nil {
		return _gbdb
	}
	if _gbdb = _bege.Add(_dabg.ToPdfObject()); _gbdb != nil {
		return _gbdb
	}
	return nil
}

func _eagdc(_becb *_db.CompliancePdfReader) (*_cb.PdfObjectDictionary, bool) {
	_debeb, _ggccd := _becb.GetTrailer()
	if _ggccd != nil {
		_eg.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _ggccd)
		return nil, false
	}
	_bbfee, _afacf := _debeb.Get("\u0052\u006f\u006f\u0074").(*_cb.PdfObjectReference)
	if !_afacf {
		_eg.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return nil, false
	}
	_ddde, _afacf := _cb.GetDict(_cb.ResolveReference(_bbfee))
	if !_afacf {
		_eg.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return nil, false
	}
	return _ddde, true
}

func _ggce(_acge *_db.CompliancePdfReader) ViolatedRule {
	_cbed := map[*_cb.PdfObjectStream]struct{}{}
	for _, _daadf := range _acge.PageList {
		if _daadf.Resources == nil && _daadf.Contents == nil {
			continue
		}
		if _fbdg := _daadf.GetPageDict(); _fbdg != nil {
			_abgeb, _gbdad := _cb.GetDict(_fbdg.Get("\u0047\u0072\u006fu\u0070"))
			if _gbdad {
				if _gade := _abgeb.Get("\u0053"); _gade != nil {
					_gdbed, _fgaf := _cb.GetName(_gade)
					if _fgaf && _gdbed.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						return _fdbe("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
					}
				}
			}
		}
		if _daadf.Resources != nil {
			if _fabbe, _gedd := _cb.GetDict(_daadf.Resources.XObject); _gedd {
				for _, _dbfcc := range _fabbe.Keys() {
					_dgbd, _bcaf := _cb.GetStream(_fabbe.Get(_dbfcc))
					if !_bcaf {
						continue
					}
					if _, _gcbc := _cbed[_dgbd]; _gcbc {
						continue
					}
					_agda, _bcaf := _cb.GetDict(_dgbd.Get("\u0047\u0072\u006fu\u0070"))
					if !_bcaf {
						_cbed[_dgbd] = struct{}{}
						continue
					}
					_bebaa := _agda.Get("\u0053")
					if _bebaa != nil {
						_dgdge, _edccb := _cb.GetName(_bebaa)
						if _edccb && _dgdge.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _fdbe("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_cbed[_dgbd] = struct{}{}
					continue
				}
			}
		}
		if _daadf.Contents != nil {
			_abeg, _gedb := _daadf.GetContentStreams()
			if _gedb != nil {
				continue
			}
			for _, _bbddg := range _abeg {
				_gdfef, _beec := _df.NewContentStreamParser(_bbddg).Parse()
				if _beec != nil {
					continue
				}
				for _, _dcdf := range *_gdfef {
					if len(_dcdf.Params) == 0 {
						continue
					}
					_fafe, _edbe := _cb.GetName(_dcdf.Params[0])
					if !_edbe {
						continue
					}
					_cggf, _fcafg := _daadf.Resources.GetXObjectByName(*_fafe)
					if _fcafg != _db.XObjectTypeForm {
						continue
					}
					if _, _deab := _cbed[_cggf]; _deab {
						continue
					}
					_aedc, _edbe := _cb.GetDict(_cggf.Get("\u0047\u0072\u006fu\u0070"))
					if !_edbe {
						_cbed[_cggf] = struct{}{}
						continue
					}
					_begda := _aedc.Get("\u0053")
					if _begda != nil {
						_dfce, _gccgb := _cb.GetName(_begda)
						if _gccgb && _dfce.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _fdbe("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_cbed[_cggf] = struct{}{}
				}
			}
		}
	}
	return _ce
}

func _adge(_beeb *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_gffae = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0032"
		_fbcaa = "IS\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068a\u0074\u0020\u0061\u006c\u006c\u0020\u0065m\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0073\u0020\u0069n\u0020t\u0068e\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006fG\u0049\u0044M\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079 \u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020\u0061\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070p\u0069\u006e\u0067 f\u0072\u006f\u006d \u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079p\u0068 \u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006fr\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u002e"
	)
	var _deegf string
	if _eabfbc, _ecda := _cb.GetName(_beeb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ecda {
		_deegf = _eabfbc.String()
	}
	if _deegf != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _ce
	}
	if _beeb.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _fdbe(_gffae, _fbcaa)
	}
	return _ce
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-2 standard.
func (_gacb *profile2) ValidateStandard(r *_db.CompliancePdfReader) error {
	_fcaf := VerificationError{ConformanceLevel: _gacb._dgfgd._ed, ConformanceVariant: _gacb._dgfgd._fd}
	if _dbcb := _ccdg(r); _dbcb != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _dbcb)
	}
	if _dfc := _gbea(r); _dfc != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _dfc)
	}
	if _fdgc := _abdec(r); _fdgc != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _fdgc)
	}
	if _beegf := _cefgg(r); _beegf != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _beegf)
	}
	if _caaa := _adec(r); _caaa != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _caaa)
	}
	if _gegf := _cfga(r); len(_gegf) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _gegf...)
	}
	if _ebfc := _gaaf(r); len(_ebfc) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _ebfc...)
	}
	if _eeb := _dfea(r); len(_eeb) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _eeb...)
	}
	if _bgab := _bebg(r); _bgab != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _bgab)
	}
	if _ggefd := _bdca(r); len(_ggefd) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _ggefd...)
	}
	if _ege := _aee(r); len(_ege) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _ege...)
	}
	if _cfee := _eddf(r); _cfee != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _cfee)
	}
	if _cgaebe := _bagfb(r); len(_cgaebe) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _cgaebe...)
	}
	if _dbdb := _dfaf(r); len(_dbdb) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _dbdb...)
	}
	if _gadd := _agdc(r); _gadd != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _gadd)
	}
	if _afac := _gbef(r); len(_afac) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _afac...)
	}
	if _dbcd := _ecec(r); len(_dbcd) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _dbcd...)
	}
	if _dceg := _dddg(r); _dceg != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _dceg)
	}
	if _bdcf := _geged(r); len(_bdcf) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _bdcf...)
	}
	if _gdaf := _cecag(r, _gacb._dgfgd); len(_gdaf) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _gdaf...)
	}
	if _fbdd := _dfbag(r); len(_fbdd) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _fbdd...)
	}
	if _bfdc := _ecee(r); len(_bfdc) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _bfdc...)
	}
	if _acbe := _eddd(r); len(_acbe) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _acbe...)
	}
	if _cead := _efff(r); _cead != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _cead)
	}
	if _bebdf := _deabd(r); len(_bebdf) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _bebdf...)
	}
	if _bdbdc := _abcdb(r); _bdbdc != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _bdbdc)
	}
	if _bffc := _gfeg(r, _gacb._dgfgd, false); len(_bffc) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _bffc...)
	}
	if _gacb._dgfgd == _ad() {
		if _gfefd := _gcee(r); len(_gfefd) != 0 {
			_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _gfefd...)
		}
	}
	if _geb := _cgeeg(r); len(_geb) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _geb...)
	}
	if _dfcd := _egagd(r); len(_dfcd) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _dfcd...)
	}
	if _gbbb := _gega(r); len(_gbbb) != 0 {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _gbbb...)
	}
	if _abe := _afcc(r); _abe != _ce {
		_fcaf.ViolatedRules = append(_fcaf.ViolatedRules, _abe)
	}
	if len(_fcaf.ViolatedRules) > 0 {
		_a.Slice(_fcaf.ViolatedRules, func(_afc, _fgge int) bool {
			return _fcaf.ViolatedRules[_afc].RuleNo < _fcaf.ViolatedRules[_fgge].RuleNo
		})
		return _fcaf
	}
	return nil
}

func _fgbg(_bbf, _ddcc, _acgf, _bge string) (string, bool) {
	_eadg := _g.Index(_bbf, _ddcc)
	if _eadg == -1 {
		return "", false
	}
	_egb := _g.Index(_bbf, _acgf)
	if _egb == -1 {
		return "", false
	}
	if _egb < _eadg {
		return "", false
	}
	return _bbf[:_eadg] + _ddcc + _bge + _bbf[_egb:], true
}
func _ebg() standardType { return standardType{_ed: 1, _fd: "\u0042"} }
func _agdge(_agbc *_db.PdfFont, _fdfb *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_eaca = "\u0036.\u0033\u002e\u0037\u002d\u0031"
		_fbbf = "\u0041\u006cl \u006e\u006f\u006e\u002d\u0073\u0079\u006db\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0066o\u006e\u0074s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020e\u0069\u0074h\u0065\u0072\u0020\u004d\u0061\u0063\u0052\u006f\u006d\u0061\u006e\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041\u006e\u0073i\u0045n\u0063\u006f\u0064\u0069n\u0067\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072\u0020t\u0068\u0065 \u0045n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020t\u0068e\u0020\u0046o\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0072\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066\u006f\u0072 \u0074\u0068\u0065\u0020\u0042\u0061\u0073\u0065\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065 \u0064i\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0069s\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0074\u0068e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006be\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046\u006f\u006e\u0074 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u0049\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e, \u006eo\u0020n\u006f\u006e\u002d\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079p\u0065 \u0066\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0020\u0061\u0020\u0044\u0069\u0066\u0066e\u0072\u0065\u006e\u0063\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0075n\u006c\u0065s\u0073\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0074h\u0065\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u006e\u0061\u006d\u0065\u0073 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006ec\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0061\u0072\u0065\u0020\u006c\u0069\u0073\u0074\u0065\u0064 \u0069\u006e \u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065 G\u006c\u0079\u0070\u0068\u0020\u004c\u0069\u0073t\u0020\u0061\u006e\u0064\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066o\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072a\u006d\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073t\u0020\u0074\u0068\u0065\u0020\u004d\u0069\u0063\u0072o\u0073o\u0066\u0074\u0020\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0028\u0033\u002c\u0031 \u2013 P\u006c\u0061\u0074\u0066\u006f\u0072\u006d\u0020I\u0044\u003d\u0033\u002c\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067 I\u0044\u003d\u0031\u0029\u0020\u0065\u006e\u0063\u006f\u0064i\u006e\u0067 \u0069\u006e\u0020t\u0068\u0065\u0020'\u0063\u006d\u0061\u0070\u0027\u0020\u0074\u0061\u0062\u006c\u0065\u002e"
	)
	var _baad string
	if _afe, _abgae := _cb.GetName(_fdfb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _abgae {
		_baad = _afe.String()
	}
	if _baad != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _ce
	}
	_eaada := _agbc.FontDescriptor()
	_abbd, _fcac := _cb.GetIntVal(_eaada.Flags)
	if !_fcac {
		_eg.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _fdbe(_eaca, _fbbf)
	}
	_cedee := (uint32(_abbd) >> 3) != 0
	if _cedee {
		return _ce
	}
	_dcfd, _fcac := _cb.GetName(_fdfb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if !_fcac {
		return _fdbe(_eaca, _fbbf)
	}
	switch _dcfd.String() {
	case "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067":
		return _ce
	default:
		return _fdbe(_eaca, _fbbf)
	}
}

// Profile2A is the implementation of the PDF/A-2A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2A struct{ profile2 }

func _eabf(_bbbg *_db.CompliancePdfReader) ViolatedRule {
	_cgee := _bbbg.ParserMetadata().HeaderCommentBytes()
	if _cgee[0] > 127 && _cgee[1] > 127 && _cgee[2] > 127 && _cgee[3] > 127 {
		return _ce
	}
	return _fdbe("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}

func _gaaf(_abfea *_db.CompliancePdfReader) (_bbff []ViolatedRule) {
	var _gcfe, _fded, _dccab bool
	if _abfea.ParserMetadata().HasNonConformantStream() {
		_bbff = []ViolatedRule{_fdbe("\u0036.\u0031\u002e\u0037\u002d\u0032", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _ceabc := range _abfea.GetObjectNums() {
		_cdddb, _ := _abfea.GetIndirectObjectByNumber(_ceabc)
		if _cdddb == nil {
			continue
		}
		_fegga, _fefa := _cb.GetStream(_cdddb)
		if !_fefa {
			continue
		}
		if !_gcfe {
			_bffeg := _fegga.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _bffeg == nil {
				_bbff = append(_bbff, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_gcfe = true
			} else {
				_abea, _ebde := _cb.GetIntVal(_bffeg)
				if !_ebde {
					_bbff = append(_bbff, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0031", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_gcfe = true
				} else {
					if len(_fegga.Stream) != _abea {
						_bbff = append(_bbff, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_gcfe = true
					}
				}
			}
		}
		if !_fded {
			if _fegga.Get("\u0046") != nil {
				_fded = true
				_bbff = append(_bbff, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _fegga.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_fded {
				_fded = true
				_bbff = append(_bbff, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _fegga.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_fded {
				_fded = true
				_bbff = append(_bbff, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_dccab {
			_fgfa, _egdd := _cb.GetName(_cb.TraceToDirectObject(_fegga.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_egdd {
				continue
			}
			if *_fgfa == _cb.StreamEncodingFilterNameLZW {
				_dccab = true
				_bbff = append(_bbff, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0034", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _bbff
}

func _bagb(_ffe *_f.Document, _ebae *_f.Page, _ccfb []*_f.Image) error {
	for _, _dagg := range _ccfb {
		if _dagg.SMask == nil {
			continue
		}
		_gffg, _cdbf := _db.NewXObjectImageFromStream(_dagg.Stream)
		if _cdbf != nil {
			return _cdbf
		}
		_dded, _cdbf := _gffg.ToImage()
		if _cdbf != nil {
			return _cdbf
		}
		_cad, _cdbf := _dded.ToGoImage()
		if _cdbf != nil {
			return _cdbf
		}
		_bgag, _cdbf := _dbg.RGBAConverter.Convert(_cad)
		if _cdbf != nil {
			return _cdbf
		}
		_gbec := _bgag.Base()
		_gcge := &_db.Image{Width: int64(_gbec.Width), Height: int64(_gbec.Height), BitsPerComponent: int64(_gbec.BitsPerComponent), ColorComponents: _gbec.ColorComponents, Data: _gbec.Data}
		_gcge.SetDecode(_gbec.Decode)
		_gcge.SetAlpha(_gbec.Alpha)
		if _cdbf = _gffg.SetImage(_gcge, nil); _cdbf != nil {
			return _cdbf
		}
		_gffg.SMask = _cb.MakeNull()
		var _bfcd _cb.PdfObject
		_acbc := -1
		for _acbc, _bfcd = range _ffe.Objects {
			if _bfcd == _dagg.SMask.Stream {
				break
			}
		}
		if _acbc != -1 {
			_ffe.Objects = append(_ffe.Objects[:_acbc], _ffe.Objects[_acbc+1:]...)
		}
		_dagg.SMask = nil
		_gffg.ToPdfObject()
	}
	return nil
}
func (_cf *documentImages) hasOnlyDeviceRGB() bool { return _cf._acda && !_cf._dbd && !_cf._fa }

var _ Profile = (*Profile2B)(nil)

func _cfga(_dbgg *_db.CompliancePdfReader) (_baae []ViolatedRule) {
	if _dbgg.ParserMetadata().HasOddLengthHexStrings() {
		_baae = append(_baae, _fdbe("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _dbgg.ParserMetadata().HasOddLengthHexStrings() {
		_baae = append(_baae, _fdbe("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _baae
}

func _ecad(_eeee *_f.Document) error {
	_dbda, _ffb := _eeee.FindCatalog()
	if !_ffb {
		return _ea.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_dbda.SetVersion()
	return nil
}

func _cdge(_eac *_f.Document, _dgf bool) error {
	_cgaeb, _fgc := _eac.GetPages()
	if !_fgc {
		return nil
	}
	for _, _cde := range _cgaeb {
		_gfad := _cde.FindXObjectForms()
		for _, _fac := range _gfad {
			_fea, _fggd := _db.NewXObjectFormFromStream(_fac)
			if _fggd != nil {
				return _fggd
			}
			_dacg, _fggd := _fea.GetContentStream()
			if _fggd != nil {
				return _fggd
			}
			_eeca := _df.NewContentStreamParser(string(_dacg))
			_faeg, _fggd := _eeca.Parse()
			if _fggd != nil {
				return _fggd
			}
			_fgda, _fggd := _fde(_fea.Resources, _faeg, _dgf)
			if _fggd != nil {
				return _fggd
			}
			if len(_fgda) == 0 {
				continue
			}
			if _fggd = _fea.SetContentStream(_fgda, _cb.NewFlateEncoder()); _fggd != nil {
				return _fggd
			}
			_fea.ToPdfObject()
		}
	}
	return nil
}
func _ecec(_acace *_db.CompliancePdfReader) (_cgccc []ViolatedRule) { return _cgccc }
func _ceag(_cgbf *_db.CompliancePdfReader) (_aegc ViolatedRule) {
	for _, _gfbg := range _cgbf.GetObjectNums() {
		_abc, _dbfc := _cgbf.GetIndirectObjectByNumber(_gfbg)
		if _dbfc != nil {
			continue
		}
		_ggcb, _gfdb := _cb.GetStream(_abc)
		if !_gfdb {
			continue
		}
		_fcdf, _gfdb := _cb.GetName(_ggcb.Get("\u0054\u0079\u0070\u0065"))
		if !_gfdb {
			continue
		}
		if *_fcdf != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_agcc, _gfdb := _cb.GetName(_ggcb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_gfdb {
			continue
		}
		if *_agcc == "\u0050\u0053" {
			return _fdbe("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073.")
		}
	}
	return _aegc
}

type profile1 struct {
	_geg standardType
	_fbg Profile1Options
}

func _gfba(_cagg *_f.Document) error {
	_bca := func(_aefb *_cb.PdfObjectDictionary) error {
		if _aefb.Get("\u0054\u0052") != nil {
			_eg.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_aefb.Remove("\u0054\u0052")
		}
		_cgad := _aefb.Get("\u0054\u0052\u0032")
		if _cgad != nil {
			_gaef := _cgad.String()
			if _gaef != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_eg.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_aefb.Set("\u0054\u0052\u0032", _cb.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		if _aefb.Get("\u0048\u0054\u0050") != nil {
			_eg.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0073\u0020\u0048\u0054P\u0020\u006b\u0065\u0079")
			_aefb.Remove("\u0048\u0054\u0050")
		}
		_ffac := _aefb.Get("\u0042\u004d")
		if _ffac != nil {
			_egac, _ddbe := _cb.GetName(_ffac)
			if !_ddbe {
				_eg.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_egac = _cb.MakeName("")
			}
			_ddf := _egac.String()
			switch _ddf {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
			default:
				_aefb.Set("\u0042\u004d", _cb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		return nil
	}
	_ggd, _begd := _cagg.GetPages()
	if !_begd {
		return nil
	}
	for _, _ecdg := range _ggd {
		_ecbb, _ggef := _ecdg.GetResources()
		if !_ggef {
			continue
		}
		_gcgg, _bfdb := _cb.GetDict(_ecbb.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_bfdb {
			return nil
		}
		_fbbg := _gcgg.Keys()
		for _, _efgf := range _fbbg {
			_aegb, _bec := _cb.GetDict(_gcgg.Get(_efgf))
			if !_bec {
				continue
			}
			_cdgf := _bca(_aegb)
			if _cdgf != nil {
				continue
			}
		}
	}
	for _, _afda := range _ggd {
		_fega, _gfaed := _afda.GetContents()
		if !_gfaed {
			return nil
		}
		for _, _daae := range _fega {
			_fdda, _agc := _daae.GetData()
			if _agc != nil {
				continue
			}
			_ddff := _df.NewContentStreamParser(string(_fdda))
			_deg, _agc := _ddff.Parse()
			if _agc != nil {
				continue
			}
			for _, _dbgbda := range *_deg {
				if len(_dbgbda.Params) == 0 {
					continue
				}
				_, _daacd := _cb.GetName(_dbgbda.Params[0])
				if !_daacd {
					continue
				}
				_caafa, _bgccb := _afda.GetResourcesXObject()
				if !_bgccb {
					continue
				}
				for _, _bdeg := range _caafa.Keys() {
					_ecgg, _caed := _cb.GetStream(_caafa.Get(_bdeg))
					if !_caed {
						continue
					}
					_gacf, _caed := _cb.GetDict(_ecgg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_caed {
						continue
					}
					_gea, _caed := _cb.GetDict(_gacf.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_caed {
						continue
					}
					for _, _ceff := range _gea.Keys() {
						_eefc, _afdb := _cb.GetDict(_gea.Get(_ceff))
						if !_afdb {
							continue
						}
						_ebbe := _bca(_eefc)
						if _ebbe != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}

var _ Profile = (*Profile1B)(nil)

func _cgeeg(_dfgfg *_db.CompliancePdfReader) (_eabgg []ViolatedRule) {
	_gdcea := _dfgfg.GetObjectNums()
	for _, _cfece := range _gdcea {
		_fabe, _fcfb := _dfgfg.GetIndirectObjectByNumber(_cfece)
		if _fcfb != nil {
			continue
		}
		_eedgb, _ccgea := _cb.GetDict(_fabe)
		if !_ccgea {
			continue
		}
		_dceab, _ccgea := _cb.GetName(_eedgb.Get("\u0054\u0079\u0070\u0065"))
		if !_ccgea {
			continue
		}
		if _dceab.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _eedgb.Get("\u0045\u0046") != nil {
			if _eedgb.Get("\u0046") == nil || _eedgb.Get("\u0045\u0046") == nil {
				_eabgg = append(_eabgg, _fdbe("\u0036\u002e\u0038-\u0032", "\u0054h\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063t\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020t\u0068\u0065\u0020\u0046\u0020a\u006e\u0064\u0020\u0055\u0046\u0020\u006b\u0065\u0079\u0073\u002e"))
			}
			if _eedgb.Get("\u0041\u0046\u0052\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070") == nil {
				_eabgg = append(_eabgg, _fdbe("\u0036\u002e\u0038-\u0033", "\u0049\u006e\u0020\u006f\u0072d\u0065\u0072\u0020\u0074\u006f\u0020\u0065\u006e\u0061\u0062\u006c\u0065\u0020i\u0064\u0065nt\u0069\u0066\u0069c\u0061\u0074\u0069o\u006e\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0074\u0068\u0065\u0020fi\u006ce\u0020\u0073\u0070\u0065\u0063\u0069f\u0069c\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020c\u006f\u006e\u0074e\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u0074\u002c\u0020\u0061\u0020\u006e\u0065\u0077\u0020(\u0072\u0065\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006be\u0079\u0020h\u0061\u0073\u0020\u0062e\u0065\u006e\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020a\u006e\u0064\u0020\u0069\u0074s \u0070\u0072e\u0073\u0065n\u0063\u0065\u0020\u0028\u0069\u006e\u0020\u0074\u0068e\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0029\u0020\u0069\u0073\u0020\u0072\u0065q\u0075\u0069\u0072e\u0064\u002e"))
			}
			break
		}
	}
	return _eabgg
}

func _fdagg(_afgd *_db.PdfFont, _ffcca *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_fga  = "\u0036.\u0033\u002e\u0035\u002d\u0032"
		_ggcc = "\u0046\u006f\u0072\u0020\u0061l\u006c\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074 \u0073\u0075bs\u0065\u0074\u0073 \u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074he\u0020f\u006f\u006e\u0074\u0020\u0064\u0065s\u0063r\u0069\u0070\u0074o\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006ec\u006c\u0075\u0064e\u0020\u0061\u0020\u0043\u0068\u0061\u0072\u0053\u0065\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u006c\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072 \u006e\u0061\u006d\u0065\u0073\u0020d\u0065\u0066i\u006e\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020f\u006f\u006e\u0074\u0020s\u0075\u0062\u0073\u0065\u0074, \u0061\u0073 \u0064\u0065s\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e \u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0061\u0062\u006ce\u0020\u0035\u002e1\u0038\u002e"
	)
	var _agad string
	if _dgcb, _acacb := _cb.GetName(_ffcca.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _acacb {
		_agad = _dgcb.String()
	}
	if _agad != "\u0054\u0079\u0070e\u0031" {
		return _ce
	}
	if _acd.IsStdFont(_acd.StdFontName(_afgd.BaseFont())) {
		return _ce
	}
	_bfgb := _afgd.FontDescriptor()
	if _bfgb.CharSet == nil {
		return _fdbe(_fga, _ggcc)
	}
	return _ce
}

func _eddc(_cadc *_f.Document) error {
	_bbgg, _ddag := _cadc.FindCatalog()
	if !_ddag {
		return _ea.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_gafg, _ddag := _cb.GetDict(_bbgg.Object.Get("\u0050\u0065\u0072m\u0073"))
	if _ddag {
		_fgde := _cb.MakeDict()
		_bcb := _gafg.Keys()
		for _, _gffa := range _bcb {
			if _gffa.String() == "\u0055\u0052\u0033" || _gffa.String() == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				_fgde.Set(_gffa, _gafg.Get(_gffa))
			}
		}
		_bbgg.Object.Set("\u0050\u0065\u0072m\u0073", _fgde)
	}
	return nil
}

func _abdec(_faggc *_db.CompliancePdfReader) ViolatedRule {
	_deca, _ddfeb := _faggc.PdfReader.GetTrailer()
	if _ddfeb != nil {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _deca.Get("\u0049\u0044") == nil {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _deca.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _ce
}

func _edgg(_ebfef *_cb.PdfObjectDictionary, _gfbaf map[*_cb.PdfObjectStream][]byte, _bgef map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_fcaa = "\u0036.\u0033\u002e\u0033\u002d\u0033"
		_abeb = "\u0041\u006cl \u0043\u004d\u0061\u0070\u0073\u0020\u0075\u0073e\u0064 \u0077i\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072m\u0069n\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020a\u006e\u0064\u0020\u0049\u0064\u0065\u006et\u0069\u0074\u0079-\u0056\u002c\u0020\u0073\u0068a\u006c\u006c \u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064es\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u00205\u002e\u0036\u002e\u0034\u002e"
	)
	var _cdga string
	if _egcab, _cfafa := _cb.GetName(_ebfef.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cfafa {
		_cdga = _egcab.String()
	}
	if _cdga != "\u0054\u0079\u0070e\u0030" {
		return _ce
	}
	_bbdd := _ebfef.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _dffe, _acacd := _cb.GetName(_bbdd); _acacd {
		switch _dffe.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _ce
		default:
			return _fdbe(_fcaa, _abeb)
		}
	}
	_ffcf, _abdd := _cb.GetStream(_bbdd)
	if !_abdd {
		return _fdbe(_fcaa, _abeb)
	}
	_, _bcec := _feac(_ffcf, _gfbaf, _bgef)
	if _bcec != nil {
		return _fdbe(_fcaa, _abeb)
	}
	return _ce
}

// Profile2Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile2Options struct {
	// CMYKDefaultColorSpace is an option that refers PDF/A
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _c.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}
type profile2 struct {
	_dgfgd standardType
	_fdbb  Profile2Options
}

func _dggg(_cdagg *_db.CompliancePdfReader) (_bgdd ViolatedRule) {
	for _, _abf := range _cdagg.GetObjectNums() {
		_cbfag, _fgfd := _cdagg.GetIndirectObjectByNumber(_abf)
		if _fgfd != nil {
			continue
		}
		_fgbd, _egag := _cb.GetStream(_cbfag)
		if !_egag {
			continue
		}
		_ffde, _egag := _cb.GetName(_fgbd.Get("\u0054\u0079\u0070\u0065"))
		if !_egag {
			continue
		}
		if *_ffde != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _fgbd.Get("\u0052\u0065\u0066") != nil {
			return _fdbe("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e")
		}
	}
	return _bgdd
}

func _bgdb(_ceef *_db.CompliancePdfReader, _eeba standardType) (_bfcb []ViolatedRule) {
	var _ggea, _degb, _aaeg, _dbbb, _gfdg, _bbcg, _dbcf, _baa, _cfcef, _agdga, _cacg bool
	_bdga := func() bool {
		return _ggea && _degb && _aaeg && _dbbb && _gfdg && _bbcg && _dbcf && _baa && _cfcef && _agdga && _cacg
	}
	_gbbgg := map[*_cb.PdfObjectStream]*_cd.CMap{}
	_bcgd := map[*_cb.PdfObjectStream][]byte{}
	_beed := map[_cb.PdfObject]*_db.PdfFont{}
	for _, _agggd := range _ceef.GetObjectNums() {
		_dcebg, _dagf := _ceef.GetIndirectObjectByNumber(_agggd)
		if _dagf != nil {
			continue
		}
		_dbgd, _facd := _cb.GetDict(_dcebg)
		if !_facd {
			continue
		}
		_dged, _facd := _cb.GetName(_dbgd.Get("\u0054\u0079\u0070\u0065"))
		if !_facd {
			continue
		}
		if *_dged != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_dff, _dagf := _db.NewPdfFontFromPdfObject(_dbgd)
		if _dagf != nil {
			_eg.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _dagf)
			continue
		}
		_beed[_dbgd] = _dff
	}
	for _, _dedg := range _ceef.PageList {
		_fdfd, _agbdd := _dedg.GetContentStreams()
		if _agbdd != nil {
			_eg.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _gccf := range _fdfd {
			_gcfg := _df.NewContentStreamParser(_gccf)
			_aaaf, _fggf := _gcfg.Parse()
			if _fggf != nil {
				_eg.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _fggf)
				continue
			}
			var _agfb bool
			for _, _cebdd := range *_aaaf {
				if _cebdd.Operand != "\u0054\u0072" {
					continue
				}
				if len(_cebdd.Params) != 1 {
					_eg.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_cebdd.Params))
					continue
				}
				_bgbd, _bbag := _cb.GetIntVal(_cebdd.Params[0])
				if !_bbag {
					_eg.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _bgbd == 3 {
					_agfb = true
					break
				}
			}
			for _, _efgff := range *_aaaf {
				if _efgff.Operand != "\u0054\u0066" {
					continue
				}
				if len(_efgff.Params) != 2 {
					_eg.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_efgff.Params))
					continue
				}
				_egec, _fdad := _cb.GetName(_efgff.Params[0])
				if !_fdad {
					_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _efgff)
					continue
				}
				_dagfd, _ccge := _dedg.Resources.GetFontByName(*_egec)
				if !_ccge {
					_eg.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_ffbe, _fdad := _cb.GetDict(_dagfd)
				if !_fdad {
					_eg.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_ffcc, _fdad := _beed[_ffbe]
				if !_fdad {
					var _badf error
					_ffcc, _badf = _db.NewPdfFontFromPdfObject(_ffbe)
					if _badf != nil {
						_eg.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _badf)
						continue
					}
					_beed[_ffbe] = _ffcc
				}
				if !_ggea {
					_acag := _fec(_ffbe, _bcgd, _gbbgg)
					if _acag != _ce {
						_bfcb = append(_bfcb, _acag)
						_ggea = true
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_degb {
					_gcceg := _ffgd(_ffbe)
					if _gcceg != _ce {
						_bfcb = append(_bfcb, _gcceg)
						_degb = true
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_aaeg {
					_acbcb := _edgg(_ffbe, _bcgd, _gbbgg)
					if _acbcb != _ce {
						_bfcb = append(_bfcb, _acbcb)
						_aaeg = true
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_dbbb {
					_dcfa := _aebf(_ffbe, _bcgd, _gbbgg)
					if _dcfa != _ce {
						_bfcb = append(_bfcb, _dcfa)
						_dbbb = true
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_gfdg {
					_cbgca := _aede(_ffcc, _ffbe, _agfb)
					if _cbgca != _ce {
						_gfdg = true
						_bfcb = append(_bfcb, _cbgca)
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_bbcg {
					_cgccg := _fdagg(_ffcc, _ffbe)
					if _cgccg != _ce {
						_bbcg = true
						_bfcb = append(_bfcb, _cgccg)
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_dbcf {
					_efaa := _edcbf(_ffcc, _ffbe)
					if _efaa != _ce {
						_dbcf = true
						_bfcb = append(_bfcb, _efaa)
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_baa {
					_abfd := _agdge(_ffcc, _ffbe)
					if _abfd != _ce {
						_baa = true
						_bfcb = append(_bfcb, _abfd)
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_cfcef {
					_afga := _dbad(_ffcc, _ffbe)
					if _afga != _ce {
						_cfcef = true
						_bfcb = append(_bfcb, _afga)
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_agdga {
					_dbbbg := _cafg(_ffcc, _ffbe)
					if _dbbbg != _ce {
						_agdga = true
						_bfcb = append(_bfcb, _dbbbg)
						if _bdga() {
							return _bfcb
						}
					}
				}
				if !_cacg && _eeba._fd == "\u0041" {
					_ebad := _befb(_ffbe, _bcgd, _gbbgg)
					if _ebad != _ce {
						_cacg = true
						_bfcb = append(_bfcb, _ebad)
						if _bdga() {
							return _bfcb
						}
					}
				}
			}
		}
	}
	return _bfcb
}

func _gd(_gfa *_f.Document) error {
	_caef, _cca := _gfa.FindCatalog()
	if !_cca {
		return nil
	}
	_, _cca = _cb.GetDict(_caef.Object.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
	if !_cca {
		_cbe := _cb.MakeDict()
		_cbe.Set("\u0054\u0079\u0070\u0065", _cb.MakeName("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
		_caef.Object.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074", _cbe)
	}
	return nil
}

func _cefgg(_dbcc *_db.CompliancePdfReader) ViolatedRule {
	if _dbcc.ParserMetadata().HasDataAfterEOF() {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _ce
}

func _baea(_gfff *_f.Document) error {
	_aac, _abd := _gfff.GetPages()
	if !_abd {
		return nil
	}
	for _, _dgc := range _aac {
		_gcf := _dgc.FindXObjectForms()
		for _, _afa := range _gcf {
			_bggf, _deb := _cb.GetDict(_afa.Get("\u0047\u0072\u006fu\u0070"))
			if _deb {
				if _dba := _bggf.Get("\u0053"); _dba != nil {
					_fcee, _bfc := _cb.GetName(_dba)
					if _bfc && _fcee.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_afa.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_gdfe, _cebd := _dgc.GetResourcesXObject()
		if _cebd {
			_gccg, _dgfg := _cb.GetDict(_gdfe.Get("\u0047\u0072\u006fu\u0070"))
			if _dgfg {
				_feaa := _gccg.Get("\u0053")
				if _feaa != nil {
					_cgbg, _gda := _cb.GetName(_feaa)
					if _gda && _cgbg.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_gdfe.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_gddd, _adcb := _cb.GetDict(_dgc.Object.Get("\u0047\u0072\u006fu\u0070"))
		if _adcb {
			_ecag := _gddd.Get("\u0053")
			if _ecag != nil {
				_gbgg, _ceg := _cb.GetName(_ecag)
				if _ceg && _gbgg.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
					_dgc.Object.Remove("\u0047\u0072\u006fu\u0070")
				}
			}
		}
	}
	return nil
}

func _fec(_ccaa *_cb.PdfObjectDictionary, _ddafg map[*_cb.PdfObjectStream][]byte, _ceddf map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_ffeg  = "\u0046\u006f\u0072 \u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029\u0020\u0066\u006f\u006et \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0064 \u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043I\u0044\u0053y\u0073\u0074\u0065\u006d\u0049nf\u006f\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006f\u0066\u0020i\u0074\u0073\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0061\u006e\u0064 \u0043\u004d\u0061\u0070 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0063\u006f\u006d\u0070\u0061\u0074i\u0062\u006c\u0065\u002e\u0020\u0049\u006e\u0020o\u0074\u0068\u0065\u0072\u0020\u0077\u006f\u0072\u0064\u0073\u002c\u0020\u0074\u0068\u0065\u0020R\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004fr\u0064\u0065\u0072\u0069\u006e\u0067 \u0073\u0074\u0072i\u006e\u0067\u0073\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0066\u006f\u0072 \u0074\u0068\u0061\u0074\u0020\u0066o\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u002c\u0020u\u006el\u0065ss \u0074\u0068\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073 \u0049\u0064\u0065\u006e\u0074\u0069t\u0079\u002d\u0048\u0020o\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074y\u002dV\u002e"
		_gcbad = "\u0036.\u0033\u002e\u0033\u002d\u0031"
	)
	var _befa string
	if _cdbcf, _dabfe := _cb.GetName(_ccaa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dabfe {
		_befa = _cdbcf.String()
	}
	if _befa != "\u0054\u0079\u0070e\u0030" {
		return _ce
	}
	_ddec := _ccaa.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _fgbf, _gabd := _cb.GetName(_ddec); _gabd {
		switch _fgbf.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _ce
		}
		_cefg, _agggf := _cd.LoadPredefinedCMap(_fgbf.String())
		if _agggf != nil {
			return _fdbe(_gcbad, _ffeg)
		}
		_acfa := _cefg.CIDSystemInfo()
		if _acfa.Ordering != _acfa.Registry {
			return _fdbe(_gcbad, _ffeg)
		}
		return _ce
	}
	_ccfg, _dbdc := _cb.GetStream(_ddec)
	if !_dbdc {
		return _fdbe(_gcbad, _ffeg)
	}
	_fecb, _dfbe := _feac(_ccfg, _ddafg, _ceddf)
	if _dfbe != nil {
		return _fdbe(_gcbad, _ffeg)
	}
	_dedge := _fecb.CIDSystemInfo()
	if _dedge.Ordering != _dedge.Registry {
		return _fdbe(_gcbad, _ffeg)
	}
	return _ce
}
func _fdbe(_ada string, _fcd string) ViolatedRule { return ViolatedRule{RuleNo: _ada, Detail: _fcd} }

// Profile2B is the implementation of the PDF/A-2B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2B struct{ profile2 }

func _fedc(_agdf *_db.CompliancePdfReader) ViolatedRule {
	_dacgb, _bacg := _agdf.PdfReader.GetTrailer()
	if _bacg != nil {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _dacgb.Get("\u0049\u0044") == nil {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _dacgb.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _ce
}

func _gfdd(_cgae *_f.Document) (*_cb.PdfObjectDictionary, bool) {
	_gagc, _bebf := _cgae.FindCatalog()
	if !_bebf {
		return nil, false
	}
	_acgc, _bebf := _cb.GetArray(_gagc.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_bebf {
		return nil, false
	}
	if _acgc.Len() == 0 {
		return nil, false
	}
	return _cb.GetDict(_acgc.Get(0))
}

// ViolatedRule is the structure that defines violated PDF/A rule.
type ViolatedRule struct {
	RuleNo string
	Detail string
}

func _deed(_dcea string, _bdfg string, _efgg string) (string, bool) {
	_ccag := _g.Index(_dcea, _bdfg)
	if _ccag == -1 {
		return "", false
	}
	_ccag += len(_bdfg)
	_edac := _g.Index(_dcea[_ccag:], _efgg)
	if _edac == -1 {
		return "", false
	}
	_edac = _ccag + _edac
	return _dcea[_ccag:_edac], true
}

func _ecee(_eebe *_db.CompliancePdfReader) (_gbaec []ViolatedRule) {
	for _, _bbfa := range _eebe.GetObjectNums() {
		_dgeg, _deee := _eebe.GetIndirectObjectByNumber(_bbfa)
		if _deee != nil {
			continue
		}
		_eefcg, _bffbc := _cb.GetDict(_dgeg)
		if !_bffbc {
			continue
		}
		_cccd, _bffbc := _cb.GetName(_eefcg.Get("\u0054\u0079\u0070\u0065"))
		if !_bffbc {
			continue
		}
		if _cccd.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_agcb, _bffbc := _cb.GetBool(_eefcg.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _bffbc && bool(*_agcb) {
			_gbaec = append(_gbaec, _fdbe("\u0036.\u0034\u002e\u0031\u002d\u0033", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
		if _eefcg.Get("\u0058\u0046\u0041") != nil {
			_gbaec = append(_gbaec, _fdbe("\u0036.\u0034\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020i\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020t\u0068\u0061\u0074\u0020f\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065 \u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d \u006b\u0065\u0079\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006f\u0066 \u0061 \u0050\u0044F\u002fA\u002d\u0032\u0020\u0066ile\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065n\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u0046\u0041\u0020\u006b\u0065y."))
		}
	}
	_dbebb, _bfgbg := _eagdc(_eebe)
	if _bfgbg && _dbebb.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_gbaec = append(_gbaec, _fdbe("\u0036.\u0034\u002e\u0032\u002d\u0032", "\u0041\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006cog\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u002e"))
	}
	return _gbaec
}

func _eacc(_daba *_cb.PdfObjectDictionary, _caecg map[*_cb.PdfObjectStream][]byte, _bedae map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_feba = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0033"
		_gead = "\u0041\u006c\u006c \u0043\u004d\u0061\u0070s\u0020\u0075\u0073ed\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074 th\u006f\u0073\u0065\u0020\u006ci\u0073\u0074\u0065\u0064\u0020i\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u00320\u00300\u002d1\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0035\u002e\u0032\u002c\u0020\u0054\u0061\u0062\u006c\u0065 \u0031\u00318,\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e \u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064e\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u00300\u0030-\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0035\u002e"
	)
	var _fdaa string
	if _cacf, _gaga := _cb.GetName(_daba.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _gaga {
		_fdaa = _cacf.String()
	}
	if _fdaa != "\u0054\u0079\u0070e\u0030" {
		return _ce
	}
	_cdca := _daba.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _agee, _cdce := _cb.GetName(_cdca); _cdce {
		switch _agee.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _ce
		default:
			return _fdbe(_feba, _gead)
		}
	}
	_bdceb, _gcgdb := _cb.GetStream(_cdca)
	if !_gcgdb {
		return _fdbe(_feba, _gead)
	}
	_, _aefbg := _feac(_bdceb, _caecg, _bedae)
	if _aefbg != nil {
		return _fdbe(_feba, _gead)
	}
	return _ce
}

// Profile1B is the implementation of the PDF/A-1B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1B struct{ profile1 }

func _begf(_dbdca *_db.CompliancePdfReader) (_dgfgb ViolatedRule) {
	_aacgg, _aegag := _eagdc(_dbdca)
	if !_aegag {
		return _ce
	}
	if _aacgg.Get("\u0041\u0041") != nil {
		return _fdbe("\u0036.\u0036\u002e\u0032\u002d\u0033", "\u0054\u0068e\u0020\u0064\u006f\u0063\u0075\u006d\u0065n\u0074 \u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065n\u0074r\u0079 \u0066\u006f\u0072 \u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063\u0074i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
	}
	return _ce
}
func _dfea(_cfec *_db.CompliancePdfReader) []ViolatedRule { return nil }
func _bagfb(_aabe *_db.CompliancePdfReader) (_fdgd []ViolatedRule) {
	var _dfbea, _acaga, _dbdee, _addg bool
	_fbddd := func() bool { return _dfbea && _acaga && _dbdee && _addg }
	_dacb, _abfa := _ffbc(_aabe)
	var _gbdbd _ebb.ProfileHeader
	if _abfa {
		_gbdbd, _ = _ebb.ParseHeader(_dacb.DestOutputProfile)
	}
	_ffdc := map[_cb.PdfObject]struct{}{}
	var _aeafa func(_ebeg _db.PdfColorspace) bool
	_aeafa = func(_face _db.PdfColorspace) bool {
		switch _ecgb := _face.(type) {
		case *_db.PdfColorspaceDeviceGray:
			if !_dfbea {
				if !_abfa {
					_fdgd = append(_fdgd, _fdbe("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d4", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064 \u0069\u0066\u0020\u0061\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065p\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0047\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0075r \u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u0077\u0068\u0065n \u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072a\u0079\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u002c o\u0072\u0020\u0069\u0066\u0020\u0061\u0020\u0050\u0044\u0046\u002fA\u0020\u004f\u0075tp\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002e"))
					_dfbea = true
					if _fbddd() {
						return true
					}
				}
			}
		case *_db.PdfColorspaceDeviceRGB:
			if !_acaga {
				if !_abfa || _gbdbd.ColorSpace != _ebb.ColorSpaceRGB {
					_fdgd = append(_fdgd, _fdbe("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d2", "\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042\u0020\u0073\u0068\u0061\u006cl\u0020\u006f\u006e\u006c\u0079\u0020\u0062e\u0020\u0075\u0073\u0065\u0064\u0020\u0069f\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063e\u0020\u0069n\u0064\u0065\u0070e\u006e\u0064\u0065\u006et \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0052\u0047\u0042\u0020\u0063\u006fl\u006f\u0075r\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020b\u0065\u0065\u006e\u0020s\u0065\u0074 \u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020c\u006flou\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020i\u0073\u0020\u0075\u0073\u0065\u0064\u002c\u0020\u006f\u0072\u0020if\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044F\u002f\u0041\u0020\u004fut\u0070\u0075\u0074\u0049\u006e\u0074\u0065n\u0074\u0020t\u0068\u0061t\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u006e\u0020\u0052\u0047\u0042\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061\u0074io\u006e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u002e"))
					_acaga = true
					if _fbddd() {
						return true
					}
				}
			}
		case *_db.PdfColorspaceDeviceCMYK:
			if !_dbdee {
				if !_abfa || _gbdbd.ColorSpace != _ebb.ColorSpaceCMYK {
					_fdgd = append(_fdgd, _fdbe("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d3", "\u0044e\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020\u0073hal\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u0066\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065\u0070\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044ef\u0061\u0075\u006c\u0074\u0043\u004d\u0059K\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061s\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u006fr \u0069\u0066\u0020\u0061\u0020\u0044e\u0076\u0069\u0063\u0065\u004e\u002d\u0062\u0061\u0073\u0065\u0064\u0020\u0044\u0065f\u0061\u0075\u006c\u0074\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006f\u0075r\u0020\u0073\u0070\u0061\u0063e\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e\u0020\u0073\u0065\u0074\u0020\u0077\u0068\u0065\u006e\u0020\u0074h\u0065\u0020\u0044\u0065\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006fu\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u0072\u0020t\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u0020\u0043\u004d\u0059\u004b\u0020d\u0065\u0073\u0074\u0069\u006e\u0061t\u0069\u006f\u006e\u0020\u0070r\u006f\u0066\u0069\u006c\u0065\u002e"))
					_dbdee = true
					if _fbddd() {
						return true
					}
				}
			}
		case *_db.PdfColorspaceICCBased:
			if !_addg {
				_cfdc, _adbba := _ebb.ParseHeader(_ecgb.Data)
				if _adbba != nil {
					_eg.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _adbba)
					_fdgd = append(_fdgd, func() ViolatedRule {
						return _fdbe("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e")
					}())
					_addg = true
					if _fbddd() {
						return true
					}
				}
				if !_addg {
					var _bcafd, _dfcb bool
					switch _cfdc.DeviceClass {
					case _ebb.DeviceClassPRTR, _ebb.DeviceClassMNTR, _ebb.DeviceClassSCNR, _ebb.DeviceClassSPAC:
					default:
						_bcafd = true
					}
					switch _cfdc.ColorSpace {
					case _ebb.ColorSpaceRGB, _ebb.ColorSpaceCMYK, _ebb.ColorSpaceGRAY, _ebb.ColorSpaceLAB:
					default:
						_dfcb = true
					}
					if _bcafd || _dfcb {
						_fdgd = append(_fdgd, _fdbe("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e"))
						_addg = true
						if _fbddd() {
							return true
						}
					}
				}
			}
			if _ecgb.Alternate != nil {
				return _aeafa(_ecgb.Alternate)
			}
		}
		return false
	}
	for _, _fddf := range _aabe.GetObjectNums() {
		_dgdd, _cdea := _aabe.GetIndirectObjectByNumber(_fddf)
		if _cdea != nil {
			continue
		}
		_gbged, _bfddf := _cb.GetStream(_dgdd)
		if !_bfddf {
			continue
		}
		_aabb, _bfddf := _cb.GetName(_gbged.Get("\u0054\u0079\u0070\u0065"))
		if !_bfddf || _aabb.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_abbdc, _bfddf := _cb.GetName(_gbged.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_bfddf {
			continue
		}
		_ffdc[_gbged] = struct{}{}
		switch _abbdc.String() {
		case "\u0049\u006d\u0061g\u0065":
			_bgddf, _ebcg := _db.NewXObjectImageFromStream(_gbged)
			if _ebcg != nil {
				continue
			}
			_ffdc[_gbged] = struct{}{}
			if _aeafa(_bgddf.ColorSpace) {
				return _fdgd
			}
		case "\u0046\u006f\u0072\u006d":
			_fccd, _dfbd := _cb.GetDict(_gbged.Get("\u0047\u0072\u006fu\u0070"))
			if !_dfbd {
				continue
			}
			_accb := _fccd.Get("\u0043\u0053")
			if _accb == nil {
				continue
			}
			_dgad, _gfac := _db.NewPdfColorspaceFromPdfObject(_accb)
			if _gfac != nil {
				continue
			}
			if _aeafa(_dgad) {
				return _fdgd
			}
		}
	}
	for _, _bade := range _aabe.PageList {
		_cagbf, _ddceg := _bade.GetContentStreams()
		if _ddceg != nil {
			continue
		}
		for _, _dbgc := range _cagbf {
			_ffad, _ecef := _df.NewContentStreamParser(_dbgc).Parse()
			if _ecef != nil {
				continue
			}
			for _, _ebefc := range *_ffad {
				if len(_ebefc.Params) > 1 {
					continue
				}
				switch _ebefc.Operand {
				case "\u0042\u0049":
					_afddf, _fgbebg := _ebefc.Params[0].(*_df.ContentStreamInlineImage)
					if !_fgbebg {
						continue
					}
					_daea, _ggee := _afddf.GetColorSpace(_bade.Resources)
					if _ggee != nil {
						continue
					}
					if _aeafa(_daea) {
						return _fdgd
					}
				case "\u0044\u006f":
					_edge, _dfbaee := _cb.GetName(_ebefc.Params[0])
					if !_dfbaee {
						continue
					}
					_bgfgb, _ecddf := _bade.Resources.GetXObjectByName(*_edge)
					if _, _decg := _ffdc[_bgfgb]; _decg {
						continue
					}
					switch _ecddf {
					case _db.XObjectTypeImage:
						_cdgaf, _eabfb := _db.NewXObjectImageFromStream(_bgfgb)
						if _eabfb != nil {
							continue
						}
						_ffdc[_bgfgb] = struct{}{}
						if _aeafa(_cdgaf.ColorSpace) {
							return _fdgd
						}
					case _db.XObjectTypeForm:
						_fcggc, _gcde := _cb.GetDict(_bgfgb.Get("\u0047\u0072\u006fu\u0070"))
						if !_gcde {
							continue
						}
						_acef, _gcde := _cb.GetName(_fcggc.Get("\u0043\u0053"))
						if !_gcde {
							continue
						}
						_fceda, _bffd := _db.NewPdfColorspaceFromPdfObject(_acef)
						if _bffd != nil {
							continue
						}
						_ffdc[_bgfgb] = struct{}{}
						if _aeafa(_fceda) {
							return _fdgd
						}
					}
				}
			}
		}
	}
	return _fdgd
}

func _afae(_abde *_f.Document) error {
	for _, _cdbb := range _abde.Objects {
		_ggf, _aga := _cb.GetDict(_cdbb)
		if !_aga {
			continue
		}
		_gbge := _ggf.Get("\u0054\u0079\u0070\u0065")
		if _gbge == nil {
			continue
		}
		if _gfbd, _bfef := _cb.GetName(_gbge); _bfef && _gfbd.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_gcgf, _gcca := _cb.GetBool(_ggf.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _gcca && bool(*_gcgf) {
			_ggf.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cb.MakeBool(false))
		}
		if _ggf.Get("\u0058\u0046\u0041") != nil {
			_ggf.Remove("\u0058\u0046\u0041")
		}
	}
	_bba, _dcde := _abde.FindCatalog()
	if !_dcde {
		return _ea.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _bba.Object.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_bba.Object.Remove("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067")
	}
	return nil
}

func _aebf(_dddb *_cb.PdfObjectDictionary, _ccbda map[*_cb.PdfObjectStream][]byte, _deeb map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_bcff  = "\u0036.\u0033\u002e\u0033\u002d\u0034"
		_agccg = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _febed string
	if _aggb, _eedg := _cb.GetName(_dddb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _eedg {
		_febed = _aggb.String()
	}
	if _febed != "\u0054\u0079\u0070e\u0030" {
		return _ce
	}
	_dgdg := _dddb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _ffca := _cb.GetName(_dgdg); _ffca {
		return _ce
	}
	_cabg, _bfca := _cb.GetStream(_dgdg)
	if !_bfca {
		return _fdbe(_bcff, _agccg)
	}
	_bgccbf, _fgfe := _feac(_cabg, _ccbda, _deeb)
	if _fgfe != nil {
		return _fdbe(_bcff, _agccg)
	}
	_dcae, _gbeg := _cb.GetIntVal(_cabg.Get("\u0057\u004d\u006fd\u0065"))
	_fgbeb, _fgeg := _bgccbf.WMode()
	if _gbeg && _fgeg {
		if _fgbeb != _dcae {
			return _fdbe(_bcff, _agccg)
		}
	}
	if (_gbeg && !_fgeg) || (!_gbeg && _fgeg) {
		return _fdbe(_bcff, _agccg)
	}
	return _ce
}

func _dfcee(_cegca *_db.PdfFont, _bfgdg *_cb.PdfObjectDictionary, _fcca bool) ViolatedRule {
	const (
		_fgac  = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0034\u002d\u0031"
		_abgdg = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006ds\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020f\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0072e\u006e\u0064\u0065\u0072\u0069\u006eg\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020w\u0069t\u0068\u0069\u006e\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u002c \u0061\u0073\u0020\u0064\u0065\u0066\u0069n\u0065\u0064 \u0069\u006e\u0020\u0049S\u004f\u0020\u0033\u0032\u00300\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0039\u002e"
	)
	if _fcca {
		return _ce
	}
	_geega := _cegca.FontDescriptor()
	var _bagc string
	if _dgec, _cabf := _cb.GetName(_bfgdg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cabf {
		_bagc = _dgec.String()
	}
	switch _bagc {
	case "\u0054\u0079\u0070e\u0031":
		if _geega.FontFile == nil {
			return _fdbe(_fgac, _abgdg)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _geega.FontFile2 == nil {
			return _fdbe(_fgac, _abgdg)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _geega.FontFile3 == nil {
			return _fdbe(_fgac, _abgdg)
		}
	}
	return _ce
}

func _geged(_edgb *_db.CompliancePdfReader) (_gddda []ViolatedRule) {
	var _geaa, _bbbbb, _bcaaf, _fadb, _bcac, _cffe bool
	_ecgd := func() bool { return _geaa && _bbbbb && _bcaaf && _fadb && _bcac && _cffe }
	for _, _ebaee := range _edgb.PageList {
		if _ebaee.Resources == nil {
			continue
		}
		_gfgf, _ffdce := _cb.GetDict(_ebaee.Resources.Font)
		if !_ffdce {
			continue
		}
		for _, _eddfc := range _gfgf.Keys() {
			_gbgab, _cegf := _cb.GetDict(_gfgf.Get(_eddfc))
			if !_cegf {
				if !_geaa {
					_gddda = append(_gddda, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006e\u0064\u0020\u0066on\u0074 \u0070\u0072\u006fg\u0072\u0061\u006ds\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072mi\u006e\u0067\u0020\u0066\u0069\u006ce\u002c\u0020\u0072\u0065\u0067\u0061\u0072\u0064\u006c\u0065s\u0073\u0020\u006f\u0066\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006eg m\u006f\u0064\u0065\u0020\u0075\u0073\u0061\u0067\u0065\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0074\u0068e\u0020\u0070\u0072o\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0069\u006e \u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031:\u0032\u0030\u0030\u0038\u002c \u0039\u002e\u0036\u0020a\u006e\u0064\u0020\u0039.\u0037\u002e"))
					_geaa = true
					if _ecgd() {
						return _gddda
					}
				}
				continue
			}
			if _egdgf, _ffagd := _cb.GetName(_gbgab.Get("\u0054\u0079\u0070\u0065")); !_geaa && (!_ffagd || _egdgf.String() != "\u0046\u006f\u006e\u0074") {
				_gddda = append(_gddda, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_geaa = true
				if _ecgd() {
					return _gddda
				}
			}
			_ebbfb, _bgeae := _db.NewPdfFontFromPdfObject(_gbgab)
			if _bgeae != nil {
				continue
			}
			var _gdddaf string
			if _gged, _dgfb := _cb.GetName(_gbgab.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dgfb {
				_gdddaf = _gged.String()
			}
			if !_bbbbb {
				switch _gdddaf {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_bbbbb = true
					_gddda = append(_gddda, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _ecgd() {
						return _gddda
					}
				}
			}
			if !_bcaaf {
				if _gdddaf != "\u0054\u0079\u0070e\u0033" {
					_fgeab, _efdf := _cb.GetName(_gbgab.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_efdf || _fgeab.String() == "" {
						_gddda = append(_gddda, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_bcaaf = true
						if _ecgd() {
							return _gddda
						}
					}
				}
			}
			if _gdddaf != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_adfc := _acd.IsStdFont(_acd.StdFontName(_ebbfb.BaseFont()))
			if _adfc {
				continue
			}
			_aabce, _aafa := _cb.GetIntVal(_gbgab.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_aafa && !_fadb {
				_gddda = append(_gddda, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_fadb = true
				if _ecgd() {
					return _gddda
				}
			}
			_gbaa, _fgbb := _cb.GetIntVal(_gbgab.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_fgbb && !_bcac {
				_gddda = append(_gddda, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_bcac = true
				if _ecgd() {
					return _gddda
				}
			}
			if !_cffe {
				_agdaa, _ebadb := _cb.GetArray(_gbgab.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_ebadb || !_aafa || !_fgbb || _agdaa.Len() != _gbaa-_aabce+1 {
					_gddda = append(_gddda, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_cffe = true
					if _ecgd() {
						return _gddda
					}
				}
			}
		}
	}
	return _gddda
}

func _edff(_cgcc *_db.CompliancePdfReader) (_gadaf ViolatedRule) {
	for _, _ebdb := range _cgcc.GetObjectNums() {
		_dgfdg, _gdfaa := _cgcc.GetIndirectObjectByNumber(_ebdb)
		if _gdfaa != nil {
			continue
		}
		_fgeb, _ffcdd := _cb.GetStream(_dgfdg)
		if !_ffcdd {
			continue
		}
		_abgf, _ffcdd := _cb.GetName(_fgeb.Get("\u0054\u0079\u0070\u0065"))
		if !_ffcdd {
			continue
		}
		if *_abgf != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_fbgd, _ffcdd := _cb.GetName(_fgeb.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_ffcdd {
			continue
		}
		if *_fbgd == "\u0050\u0053" {
			return _fdbe("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _fgeb.Get("\u0050\u0053") != nil {
			return _fdbe("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _gadaf
}

// Profile is the model.StandardImplementer enhanced by the information about the profile conformance level.
type Profile interface {
	_db.StandardImplementer
	Conformance() string
	Part() int
}

// Error implements error interface.
func (_bg VerificationError) Error() string {
	_ff := _g.Builder{}
	_ff.WriteString("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u003a\u0020")
	_ff.WriteString(_b.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _bg.ConformanceLevel, _bg.ConformanceVariant))
	_ff.WriteString("\u0020\u0056\u0069\u006f\u006c\u0061\u0074\u0065\u0064\u0020\u0072\u0075l\u0065\u0073\u003a\u0020")
	for _egg, _baf := range _bg.ViolatedRules {
		_ff.WriteString(_baf.String())
		if _egg != len(_bg.ViolatedRules)-1 {
			_ff.WriteRune('\n')
		}
	}
	return _ff.String()
}
func _agef(_dcb *_db.CompliancePdfReader) ViolatedRule { return _ce }
func _ad() standardType                                { return standardType{_ed: 2, _fd: "\u0041"} }
func (_fb standardType) outputIntentSubtype() _db.PdfOutputIntentType {
	switch _fb._ed {
	case 1:
		return _db.PdfOutputIntentTypeA1
	case 2:
		return _db.PdfOutputIntentTypeA2
	case 3:
		return _db.PdfOutputIntentTypeA3
	case 4:
		return _db.PdfOutputIntentTypeA4
	default:
		return 0
	}
}

func _eaff(_gefda *_db.CompliancePdfReader) (*_cb.PdfObjectDictionary, bool) {
	_afddb, _fbed := _eagdc(_gefda)
	if !_fbed {
		return nil, false
	}
	_bafdf, _fbed := _cb.GetArray(_afddb.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_fbed {
		return nil, false
	}
	if _bafdf.Len() == 0 {
		return nil, false
	}
	return _cb.GetDict(_bafdf.Get(0))
}

func _fgdac(_eacg *_f.Document) error {
	_bdge, _dcee := _eacg.GetPages()
	if !_dcee {
		return nil
	}
	for _, _cece := range _bdge {
		_acae, _beeg := _cb.GetArray(_cece.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_beeg {
			continue
		}
		for _, _aabc := range _acae.Elements() {
			_aabc = _cb.ResolveReference(_aabc)
			if _, _gabg := _aabc.(*_cb.PdfObjectNull); _gabg {
				continue
			}
			_gga, _eagd := _cb.GetDict(_aabc)
			if !_eagd {
				continue
			}
			_fagb, _ := _cb.GetIntVal(_gga.Get("\u0046"))
			_fagb &= ^(1 << 0)
			_fagb &= ^(1 << 1)
			_fagb &= ^(1 << 5)
			_fagb &= ^(1 << 8)
			_fagb |= 1 << 2
			_gga.Set("\u0046", _cb.MakeInteger(int64(_fagb)))
			_dggf := false
			if _acc := _gga.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _acc != nil {
				_aafc, _bfba := _cb.GetName(_acc)
				if _bfba && _aafc.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_dggf = true
					if _gga.Get("\u0041\u0041") != nil {
						_gga.Remove("\u0041\u0041")
					}
					if _gga.Get("\u0041") != nil {
						_gga.Remove("\u0041")
					}
				}
				if _bfba && _aafc.String() == "\u0054\u0065\u0078\u0074" {
					_ffbg, _ := _cb.GetIntVal(_gga.Get("\u0046"))
					_ffbg |= 1 << 3
					_ffbg |= 1 << 4
					_gga.Set("\u0046", _cb.MakeInteger(int64(_ffbg)))
				}
			}
			_dgfgc, _eagd := _cb.GetDict(_gga.Get("\u0041\u0050"))
			if _eagd {
				_aaca := _dgfgc.Get("\u004e")
				if _aaca == nil {
					continue
				}
				if len(_dgfgc.Keys()) > 1 {
					_dgfgc.Clear()
					_dgfgc.Set("\u004e", _aaca)
				}
				if _dggf {
					_daab, _dacf := _cb.GetName(_gga.Get("\u0046\u0054"))
					if _dacf && *_daab == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}

func _dfaf(_fffb *_db.CompliancePdfReader) (_bgde []ViolatedRule) {
	var _adeb, _gfgd, _ceagg, _bdcdd, _ggbfe, _cdfd, _aefbb bool
	_ccdbe := func() bool { return _adeb && _gfgd && _ceagg && _bdcdd && _ggbfe && _cdfd && _aefbb }
	_dbeb := func(_decgf *_cb.PdfObjectDictionary) bool {
		if !_adeb && _decgf.Get("\u0054\u0052") != nil {
			_adeb = true
			_bgde = append(_bgde, _fdbe("\u0036.\u0032\u002e\u0035\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _ecdgdb := _decgf.Get("\u0054\u0052\u0032"); !_gfgd && _ecdgdb != nil {
			_acdb, _gbfaf := _cb.GetName(_ecdgdb)
			if !_gbfaf || (_gbfaf && *_acdb != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_gfgd = true
				_bgde = append(_bgde, _fdbe("\u0036.\u0032\u002e\u0035\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _ccdbe() {
					return true
				}
			}
		}
		if !_ceagg && _decgf.Get("\u0048\u0054\u0050") != nil {
			_ceagg = true
			_bgde = append(_bgde, _fdbe("\u0036.\u0032\u002e\u0035\u002d\u0033", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020th\u0065\u0020\u0048\u0054\u0050\u0020\u006b\u0065\u0079\u002e"))
		}
		_dcaa, _geca := _cb.GetDict(_decgf.Get("\u0048\u0054"))
		if _geca {
			if _aaacg := _dcaa.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u0054\u0079\u0070\u0065"); !_bdcdd && _aaacg != nil {
				_dgcbc, _deef := _cb.GetInt(_aaacg)
				if !_deef || (_deef && !(*_dgcbc == 1 || *_dgcbc == 5)) {
					_bgde = append(_bgde, _fdbe("\u0020\u0036\u002e\u0032\u002e\u0035\u002d\u0034", "\u0041\u006c\u006c\u0020\u0068\u0061\u006c\u0066\u0074\u006f\u006e\u0065\u0073\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006ce\u0020\u0073h\u0061\u006c\u006c\u0020h\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0031\u0020\u006f\u0072\u0020\u0035 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0048\u0061l\u0066\u0074\u006fn\u0065\u0054\u0079\u0070\u0065\u0020\u006be\u0079\u002e"))
					if _ccdbe() {
						return true
					}
				}
			}
			if _gfade := _dcaa.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u004e\u0061\u006d\u0065"); !_ggbfe && _gfade != nil {
				_ggbfe = true
				_bgde = append(_bgde, _fdbe("\u0036.\u0032\u002e\u0035\u002d\u0035", "\u0048\u0061\u006c\u0066\u0074o\u006e\u0065\u0073\u0020\u0069\u006e\u0020a\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0048\u0061\u006c\u0066\u0074\u006f\u006e\u0065N\u0061\u006d\u0065\u0020\u006b\u0065y\u002e"))
				if _ccdbe() {
					return true
				}
			}
		}
		_, _gcbg := _ffbc(_fffb)
		var _dbdce bool
		_ccdba, _geca := _cb.GetDict(_decgf.Get("\u0047\u0072\u006fu\u0070"))
		if _geca {
			_, _bcddg := _cb.GetName(_ccdba.Get("\u0043\u0053"))
			if _bcddg {
				_dbdce = true
			}
		}
		if _gabgg := _decgf.Get("\u0042\u004d"); !_cdfd && !_aefbb && _gabgg != nil {
			_bdfe, _agddd := _cb.GetName(_gabgg)
			if _agddd {
				switch _bdfe.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
				default:
					_cdfd = true
					_bgde = append(_bgde, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u004f\u006el\u0079\u0020\u0062\u006c\u0065\u006e\u0064\u0020\u006d\u006f\u0064\u0065\u0073\u0020\u0074h\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0073\u0070\u0065c\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0049\u0053O\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a2\u0030\u0030\u0038\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075e\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0042M\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0065\u0078t\u0065\u006e\u0064\u0065\u0064\u0020\u0067\u0072\u0061\u0070\u0068\u0069\u0063\u0020\u0073\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					if _ccdbe() {
						return true
					}
				}
				if _bdfe.String() != "\u004e\u006f\u0072\u006d\u0061\u006c" && !_gcbg && !_dbdce {
					_aefbb = true
					_bgde = append(_bgde, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
					if _ccdbe() {
						return true
					}
				}
			}
		}
		if _, _geca = _cb.GetDict(_decgf.Get("\u0053\u004d\u0061s\u006b")); !_aefbb && _geca && !_gcbg && !_dbdce {
			_aefbb = true
			_bgde = append(_bgde, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
			if _ccdbe() {
				return true
			}
		}
		if _gege := _decgf.Get("\u0043\u0041"); !_aefbb && _gege != nil && !_gcbg && !_dbdce {
			_gdce, _fabba := _cb.GetNumberAsFloat(_gege)
			if _fabba == nil && _gdce < 1.0 {
				_aefbb = true
				_bgde = append(_bgde, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _ccdbe() {
					return true
				}
			}
		}
		if _cfba := _decgf.Get("\u0063\u0061"); !_aefbb && _cfba != nil && !_gcbg && !_dbdce {
			_fedd, _eefe := _cb.GetNumberAsFloat(_cfba)
			if _eefe == nil && _fedd < 1.0 {
				_aefbb = true
				_bgde = append(_bgde, _fdbe("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _ccdbe() {
					return true
				}
			}
		}
		return false
	}
	for _, _fcfg := range _fffb.PageList {
		_dccae := _fcfg.Resources
		if _dccae == nil {
			continue
		}
		if _dccae.ExtGState == nil {
			continue
		}
		_bfcdd, _ddga := _cb.GetDict(_dccae.ExtGState)
		if !_ddga {
			continue
		}
		_bdde := _bfcdd.Keys()
		for _, _fdbd := range _bdde {
			_aaacga, _eaab := _cb.GetDict(_bfcdd.Get(_fdbd))
			if !_eaab {
				continue
			}
			if _dbeb(_aaacga) {
				return _bgde
			}
		}
	}
	for _, _dega := range _fffb.PageList {
		_dagc := _dega.Resources
		if _dagc == nil {
			continue
		}
		_eaba, _eecad := _cb.GetDict(_dagc.XObject)
		if !_eecad {
			continue
		}
		for _, _afeg := range _eaba.Keys() {
			_bffg, _fgggg := _cb.GetStream(_eaba.Get(_afeg))
			if !_fgggg {
				continue
			}
			_ggfg, _fgggg := _cb.GetDict(_bffg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_fgggg {
				continue
			}
			_abfdg, _fgggg := _cb.GetDict(_ggfg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_fgggg {
				continue
			}
			for _, _gbeb := range _abfdg.Keys() {
				_daef, _adbf := _cb.GetDict(_abfdg.Get(_gbeb))
				if !_adbf {
					continue
				}
				if _dbeb(_daef) {
					return _bgde
				}
			}
		}
	}
	return _bgde
}

func _gcee(_ffff *_db.CompliancePdfReader) (_dadee []ViolatedRule) {
	_facb := true
	_beff, _aaaff := _ffff.GetCatalogMarkInfo()
	if !_aaaff {
		_facb = false
	} else {
		_abdac, _acagc := _cb.GetDict(_beff)
		if _acagc {
			_gbbe, _ffbed := _cb.GetBool(_abdac.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_gbbe) || !_ffbed {
				_facb = false
			}
		} else {
			_facb = false
		}
	}
	if !_facb {
		_dadee = append(_dadee, _fdbe("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_dbccd, _aaaff := _ffff.GetCatalogStructTreeRoot()
	if !_aaaff {
		_dadee = append(_dadee, _fdbe("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_gfgdg, _aaaff := _cb.GetDict(_dbccd)
	if _aaaff {
		_gfbc, _ebead := _cb.GetName(_gfgdg.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _ebead {
			_caba, _defc := _cb.GetDict(_gfbc)
			if _defc {
				for _, _defb := range _caba.Keys() {
					_bcba := _caba.Get(_defb)
					if _bcba == nil {
						_dadee = append(_dadee, _fdbe("\u0036.\u0037\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _dadee
}

func _adee(_gcaga *_cb.PdfObjectDictionary, _gdcgf map[*_cb.PdfObjectStream][]byte, _fgggc map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_aaab  = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0034"
		_agfgb = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _ecefb string
	if _ebdcf, _ecga := _cb.GetName(_gcaga.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ecga {
		_ecefb = _ebdcf.String()
	}
	if _ecefb != "\u0054\u0079\u0070e\u0030" {
		return _ce
	}
	_gdea := _gcaga.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _aacfb := _cb.GetName(_gdea); _aacfb {
		return _ce
	}
	_fbbb, _bbbd := _cb.GetStream(_gdea)
	if !_bbbd {
		return _fdbe(_aaab, _agfgb)
	}
	_eeagg, _fbba := _feac(_fbbb, _gdcgf, _fgggc)
	if _fbba != nil {
		return _fdbe(_aaab, _agfgb)
	}
	_aece, _daaa := _cb.GetIntVal(_fbbb.Get("\u0057\u004d\u006fd\u0065"))
	_bbgfe, _eccg := _eeagg.WMode()
	if _daaa && _eccg {
		if _bbgfe != _aece {
			return _fdbe(_aaab, _agfgb)
		}
	}
	if (_daaa && !_eccg) || (!_daaa && _eccg) {
		return _fdbe(_aaab, _agfgb)
	}
	return _ce
}

func _cgeb(_cced *_db.PdfInfo, _baef func() _c.Time) error {
	var _dea *_db.PdfDate
	if _cced.CreationDate == nil {
		_fcdb, _ebag := _db.NewPdfDateFromTime(_baef())
		if _ebag != nil {
			return _ebag
		}
		_dea = &_fcdb
		_cced.CreationDate = _dea
	}
	if _cced.ModifiedDate == nil {
		if _dea != nil {
			_bgce, _fbdc := _db.NewPdfDateFromTime(_baef())
			if _fbdc != nil {
				return _fbdc
			}
			_dea = &_bgce
		}
		_cced.ModifiedDate = _dea
	}
	return nil
}

func _bgd(_gae *_f.Document, _bfee standardType, _ecdc *_f.OutputIntents) error {
	var (
		_dbde *_db.PdfOutputIntent
		_feb  error
	)
	if _gae.Version.Minor <= 7 {
		_dbde, _feb = _ebb.NewSRGBv2OutputIntent(_bfee.outputIntentSubtype())
	} else {
		_dbde, _feb = _ebb.NewSRGBv4OutputIntent(_bfee.outputIntentSubtype())
	}
	if _feb != nil {
		return _feb
	}
	if _feb = _ecdc.Add(_dbde.ToPdfObject()); _feb != nil {
		return _feb
	}
	return nil
}

// String gets a string representation of the violated rule.
func (_fdb ViolatedRule) String() string {
	return _b.Sprintf("\u0025\u0073\u003a\u0020\u0025\u0073", _fdb.RuleNo, _fdb.Detail)
}

func _gdac(_eggg *_cb.PdfObjectDictionary, _bcaab map[*_cb.PdfObjectStream][]byte, _ggdb map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_ecde = "\u0046\u006f\u0072\u0020\u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070o\u0073\u0069\u0074e\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029 \u0066\u006fn\u0074\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f \u0065\u006e\u0074\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074\u0073 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0066\u006fl\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0072\u0065l\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070. \u0049\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0054\u0079\u0070\u0065\u0020\u0030 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0069\u0073\u0020I\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056\u002c\u0020\u0061\u006e\u0079\u0020v\u0061\u006c\u0075\u0065\u0073\u0020\u006f\u0066\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u002c\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u002c\u0020\u0061\u006e\u0064\u0020\u0053up\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069n\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0065\u006e\u0074r\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044F\u006f\u006e\u0074\u002e\u0020\u004f\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u0072\u0072\u0065\u0073\u0070\u006f\u006e\u0064\u0069\u006e\u0067\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0069\u006e\u0020\u0062\u006f\u0074h\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006cl\u0020\u0062\u0065\u0020i\u0064en\u0074\u0069\u0063\u0061\u006c\u002c \u0061n\u0064\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0070\u0070l\u0065\u006d\u0065\u006e\u0074 \u006b\u0065\u0079\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0043I\u0044S\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0067re\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f t\u0068\u0065\u0020\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043M\u0061p\u002e"
		_adea = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0031"
	)
	var _ggeab string
	if _cfad, _cedbg := _cb.GetName(_eggg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cedbg {
		_ggeab = _cfad.String()
	}
	if _ggeab != "\u0054\u0079\u0070e\u0030" {
		return _ce
	}
	_ggff := _eggg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _ebaeec, _aafeg := _cb.GetName(_ggff); _aafeg {
		switch _ebaeec.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _ce
		}
		_gece, _abae := _cd.LoadPredefinedCMap(_ebaeec.String())
		if _abae != nil {
			return _fdbe(_adea, _ecde)
		}
		_acec := _gece.CIDSystemInfo()
		if _acec.Ordering != _acec.Registry {
			return _fdbe(_adea, _ecde)
		}
		return _ce
	}
	_acfcf, _fgeba := _cb.GetStream(_ggff)
	if !_fgeba {
		return _fdbe(_adea, _ecde)
	}
	_gdfag, _ceded := _feac(_acfcf, _bcaab, _ggdb)
	if _ceded != nil {
		return _fdbe(_adea, _ecde)
	}
	_ffed := _gdfag.CIDSystemInfo()
	if _ffed.Ordering != _ffed.Registry {
		return _fdbe(_adea, _ecde)
	}
	return _ce
}

func _daf(_gaf *_f.Document, _dfe int) error {
	_gge := map[*_cb.PdfObjectStream]struct{}{}
	for _, _dgfde := range _gaf.Objects {
		_ecg, _ebf := _cb.GetStream(_dgfde)
		if !_ebf {
			continue
		}
		if _, _ebf = _gge[_ecg]; _ebf {
			continue
		}
		_gge[_ecg] = struct{}{}
		_afdc, _ebf := _cb.GetName(_ecg.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_ebf {
			continue
		}
		if _ecg.Get("\u0052\u0065\u0066") != nil {
			_ecg.Remove("\u0052\u0065\u0066")
		}
		if _afdc.String() == "\u0050\u0053" {
			_ecg.Remove("\u0050\u0053")
			continue
		}
		if _afdc.String() == "\u0046\u006f\u0072\u006d" {
			if _ecg.Get("\u004f\u0050\u0049") != nil {
				_ecg.Remove("\u004f\u0050\u0049")
			}
			if _ecg.Get("\u0050\u0053") != nil {
				_ecg.Remove("\u0050\u0053")
			}
			if _bead := _ecg.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _bead != nil {
				if _aeaf, _fcab := _cb.GetName(_bead); _fcab && *_aeaf == "\u0050\u0053" {
					_ecg.Remove("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032")
				}
			}
			continue
		}
		if _afdc.String() == "\u0049\u006d\u0061g\u0065" {
			_fegg, _cfeg := _cb.GetBool(_ecg.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _cfeg && bool(*_fegg) {
				_ecg.Set("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _cb.MakeBool(false))
			}
			if _dfe == 2 {
				if _ecg.Get("\u004f\u0050\u0049") != nil {
					_ecg.Remove("\u004f\u0050\u0049")
				}
			}
			if _ecg.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
				_ecg.Remove("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073")
			}
			continue
		}
	}
	return nil
}

type pageColorspaceOptimizeFunc func(_fce *_f.Document, _bfbc *_f.Page, _gbga []*_f.Image) error

func (_ba standardType) String() string {
	return _b.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _ba._ed, _ba._fd)
}

func _baba(_gbca *_db.CompliancePdfReader) (_dbdd []ViolatedRule) {
	var _bgac, _cgea, _bdedb bool
	if _gbca.ParserMetadata().HasNonConformantStream() {
		_dbdd = []ViolatedRule{_fdbe("\u0036.\u0031\u002e\u0037\u002d\u0031", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _fbag := range _gbca.GetObjectNums() {
		_defe, _ := _gbca.GetIndirectObjectByNumber(_fbag)
		if _defe == nil {
			continue
		}
		_cffb, _feec := _cb.GetStream(_defe)
		if !_feec {
			continue
		}
		if !_bgac {
			_cecf := _cffb.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _cecf == nil {
				_dbdd = append(_dbdd, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_bgac = true
			} else {
				_ecbbf, _ebfa := _cb.GetIntVal(_cecf)
				if !_ebfa {
					_dbdd = append(_dbdd, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0032", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_bgac = true
				} else {
					if len(_cffb.Stream) != _ecbbf {
						_dbdd = append(_dbdd, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_bgac = true
					}
				}
			}
		}
		if !_cgea {
			if _cffb.Get("\u0046") != nil {
				_cgea = true
				_dbdd = append(_dbdd, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _cffb.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_cgea {
				_cgea = true
				_dbdd = append(_dbdd, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _cffb.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_cgea {
				_cgea = true
				_dbdd = append(_dbdd, _fdbe("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_bdedb {
			_gfea, _fgdaf := _cb.GetName(_cb.TraceToDirectObject(_cffb.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_fgdaf {
				continue
			}
			if *_gfea == _cb.StreamEncodingFilterNameLZW {
				_bdedb = true
				_dbdd = append(_dbdd, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0031", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _dbdd
}

func _eef(_cbcd *_f.Document, _eecg []pageColorspaceOptimizeFunc, _gaa []documentColorspaceOptimizeFunc) error {
	_afg, _afge := _cbcd.GetPages()
	if !_afge {
		return nil
	}
	var _bde []*_f.Image
	for _bab, _cce := range _afg {
		_eggb, _babb := _cce.FindXObjectImages()
		if _babb != nil {
			return _babb
		}
		for _, _gfag := range _eecg {
			if _babb = _gfag(_cbcd, &_afg[_bab], _eggb); _babb != nil {
				return _babb
			}
		}
		_bde = append(_bde, _eggb...)
	}
	for _, _edb := range _gaa {
		if _fgb := _edb(_cbcd, _bde); _fgb != nil {
			return _fgb
		}
	}
	return nil
}

func _gbef(_dbcgc *_db.CompliancePdfReader) (_fdga []ViolatedRule) {
	var _ecfb, _gdgge, _edada, _cagbb, _adccb, _egfg, _ccff bool
	_fadf := map[*_cb.PdfObjectStream]struct{}{}
	for _, _fbgc := range _dbcgc.GetObjectNums() {
		if _ecfb && _gdgge && _adccb && _edada && _cagbb && _egfg && _ccff {
			return _fdga
		}
		_ddgd, _cgfd := _dbcgc.GetIndirectObjectByNumber(_fbgc)
		if _cgfd != nil {
			continue
		}
		_eggf, _fagfd := _cb.GetStream(_ddgd)
		if !_fagfd {
			continue
		}
		if _, _fagfd = _fadf[_eggf]; _fagfd {
			continue
		}
		_fadf[_eggf] = struct{}{}
		_deec, _fagfd := _cb.GetName(_eggf.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_fagfd {
			continue
		}
		if !_cagbb {
			if _eggf.Get("\u0052\u0065\u0066") != nil {
				_fdga = append(_fdga, _fdbe("\u0036.\u0032\u002e\u0039\u002d\u0032", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_cagbb = true
			}
		}
		if _deec.String() == "\u0050\u0053" {
			if !_egfg {
				_fdga = append(_fdga, _fdbe("\u0036.\u0032\u002e\u0039\u002d\u0033", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_egfg = true
				continue
			}
		}
		if _deec.String() == "\u0046\u006f\u0072\u006d" {
			if _gdgge && _edada && _cagbb {
				continue
			}
			if !_gdgge && _eggf.Get("\u004f\u0050\u0049") != nil {
				_fdga = append(_fdga, _fdbe("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d \u0058\u004f\u0062j\u0065\u0063\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
				_gdgge = true
			}
			if !_edada {
				if _eggf.Get("\u0050\u0053") != nil {
					_edada = true
				}
				if _agag := _eggf.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _agag != nil && !_edada {
					if _cbcaee, _ccfd := _cb.GetName(_agag); _ccfd && *_cbcaee == "\u0050\u0053" {
						_edada = true
					}
				}
				if _edada {
					_fdga = append(_fdga, _fdbe("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065y \u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006cu\u0065 o\u0066 \u0050\u0053\u0020\u0061\u006e\u0064\u0020t\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _deec.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_ecfb && _eggf.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_fdga = append(_fdga, _fdbe("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_ecfb = true
		}
		if !_ccff && _eggf.Get("\u004f\u0050\u0049") != nil {
			_fdga = append(_fdga, _fdbe("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
			_ccff = true
		}
		if !_adccb && _eggf.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_gafa, _cfbg := _cb.GetBool(_eggf.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _cfbg && bool(*_gafa) {
				continue
			}
			_fdga = append(_fdga, _fdbe("\u0036.\u0032\u002e\u0038\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_adccb = true
		}
	}
	return _fdga
}

// NewProfile2B creates a new Profile2B with the given options.
func NewProfile2B(options *Profile2Options) *Profile2B {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_bdce(options)
	return &Profile2B{profile2{_fdbb: *options, _dgfgd: _ae()}}
}

func _ddbea(_gbag *_db.CompliancePdfReader) ViolatedRule {
	for _, _ebea := range _gbag.GetObjectNums() {
		_eced, _cgbc := _gbag.GetIndirectObjectByNumber(_ebea)
		if _cgbc != nil {
			continue
		}
		_bfdcc, _adfd := _cb.GetStream(_eced)
		if !_adfd {
			continue
		}
		_bdbdb, _adfd := _cb.GetName(_bfdcc.Get("\u0054\u0079\u0070\u0065"))
		if !_adfd {
			continue
		}
		if *_bdbdb != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _bfdcc.Get("\u0053\u004d\u0061s\u006b") != nil {
			return _fdbe("\u0036\u002e\u0034-\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e \u0053\u004d\u0061\u0073\u006b\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _ce
}

// NewProfile2A creates a new Profile2A with given options.
func NewProfile2A(options *Profile2Options) *Profile2A {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_bdce(options)
	return &Profile2A{profile2{_fdbb: *options, _dgfgd: _ad()}}
}

// StandardName gets the name of the standard.
func (_faae *profile2) StandardName() string {
	return _b.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0032\u0025s", _faae._dgfgd._fd)
}

func _bcfb(_dgbb *_db.CompliancePdfReader, _gaecb standardType, _gefc bool) (_fcdba []ViolatedRule) {
	_aada, _aedce := _eagdc(_dgbb)
	if !_aedce {
		return []ViolatedRule{_fdbe("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_egdb := _aada.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _egdb == nil {
		return []ViolatedRule{_fdbe("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u006e\u006f\u0020\u0027\u004d\u0065\u0074\u0061d\u0061\u0074\u0061' \u006b\u0065\u0079\u0020\u0066\u006fu\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u002e"), _fdbe("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	_fcedf, _aedce := _cb.GetStream(_egdb)
	if !_aedce {
		return []ViolatedRule{_fdbe("\u0036.\u0037\u002e\u0032\u002d\u0032", "\u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0027\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"), _fdbe("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	if _fcedf.Get("\u0046\u0069\u006c\u0074\u0065\u0072") != nil {
		_fcdba = append(_fcdba, _fdbe("\u0036.\u0037\u002e\u0032\u002d\u0032", "M\u0065\u0074a\u0064\u0061\u0074\u0061\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u006b\u0065y\u002e"))
	}
	_begc, _dbed := _fcg.LoadDocument(_fcedf.Stream)
	if _dbed != nil {
		return []ViolatedRule{_fdbe("\u0036.\u0037\u002e\u0039\u002d\u0031", "The\u0020\u006d\u0065\u0074a\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0061\u006e\u0064\u0020\u0077\u0065\u006c\u006c\u0020\u0066\u006f\u0072\u006de\u0064\u0020\u0050\u0044\u0046\u0041\u0045\u0078\u0074e\u006e\u0073\u0069\u006f\u006e\u0020\u0053\u0063\u0068\u0065\u006da\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073\u002e")}
	}
	_bgbc := _begc.GetGoXmpDocument()
	var _cgff []*_eb.Namespace
	for _, _gbage := range _bgbc.Namespaces() {
		switch _gbage.Name {
		case _da.NsDc.Name, _bfg.NsPDF.Name, _gf.NsXmp.Name, _bc.NsXmpRights.Name, _fc.Namespace.Name, _eee.Namespace.Name, _ac.NsXmpMM.Name, _eee.FieldNS.Name, _eee.SchemaNS.Name, _eee.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_cgff = append(_cgff, _gbage)
	}
	_febeb := true
	_egcg, _dbed := _begc.GetPdfaExtensionSchemas()
	if _dbed == nil {
		for _, _acged := range _cgff {
			var _dced bool
			for _dbbe := range _egcg {
				if _acged.URI == _egcg[_dbbe].NamespaceURI {
					_dced = true
					break
				}
			}
			if !_dced {
				_febeb = false
				break
			}
		}
	} else {
		_febeb = false
	}
	if !_febeb {
		_fcdba = append(_fcdba, _fdbe("\u0036.\u0037\u002e\u0039\u002d\u0032", "\u0050\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073 \u0073\u0070\u0065\u0063\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0058M\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0073\u0068\u0061\u006cl\u0020\u0075\u0073\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073 \u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002c\u0020\u006f\u0072\u0020\u0065\u0078\u0074\u0065ns\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u0074\u0068\u0061\u0074 \u0063\u006f\u006d\u0070\u006c\u0079\u0020\u0077\u0069\u0074h\u0020\u0058\u004d\u0050\u0020\u0053\u0070e\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002e"))
	}
	_fagf, _dbed := _dgbb.GetPdfInfo()
	if _dbed == nil {
		if !_aefe(_fagf, _begc) {
			_fcdba = append(_fcdba, _fdbe("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
		}
	} else if _, _cbdcd := _begc.GetMediaManagement(); _cbdcd {
		_fcdba = append(_fcdba, _fdbe("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
	}
	_cddaa, _aedce := _begc.GetPdfAID()
	if !_aedce {
		_fcdba = append(_fcdba, _fdbe("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _cddaa.Part != _gaecb._ed {
			_fcdba = append(_fcdba, _fdbe("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _gaecb._fd == "\u0041" && _cddaa.Conformance != "\u0041" {
			_fcdba = append(_fcdba, _fdbe("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		} else if _gaecb._fd == "\u0042" && (_cddaa.Conformance != "\u0041" && _cddaa.Conformance != "\u0042") {
			_fcdba = append(_fcdba, _fdbe("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		}
	}
	return _fcdba
}

var _ Profile = (*Profile2U)(nil)

func _ccdg(_ecdb *_db.CompliancePdfReader) ViolatedRule {
	if _ecdb.ParserMetadata().HeaderPosition() != 0 {
		return _fdbe("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	if _ecdb.PdfVersion().Major != 1 {
		return _fdbe("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	if _ecdb.PdfVersion().Minor < 0 || _ecdb.PdfVersion().Minor > 7 {
		return _fdbe("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	return _ce
}

func _eceg(_fad *_db.CompliancePdfReader) ViolatedRule {
	_gbba, _cbgf := _fad.GetTrailer()
	if _cbgf != nil {
		_eg.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _cbgf)
		return _ce
	}
	_ecbe, _abda := _gbba.Get("\u0052\u006f\u006f\u0074").(*_cb.PdfObjectReference)
	if !_abda {
		_eg.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return _ce
	}
	_dfbga, _abda := _cb.GetDict(_cb.ResolveReference(_ecbe))
	if !_abda {
		_eg.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return _ce
	}
	if _dfbga.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073") != nil {
		return _fdbe("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063u\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u004f\u0043\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	}
	return _ce
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-1 standard.
func (_fbaf *profile1) ValidateStandard(r *_db.CompliancePdfReader) error {
	_eede := VerificationError{ConformanceLevel: _fbaf._geg._ed, ConformanceVariant: _fbaf._geg._fd}
	if _dbga := _cbea(r); _dbga != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _dbga)
	}
	if _cbg := _eabf(r); _cbg != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _cbg)
	}
	if _cea := _fedc(r); _cea != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _cea)
	}
	if _bffb := _befg(r); _bffb != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _bffb)
	}
	if _gbggb := _agef(r); _gbggb != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _gbggb)
	}
	if _gefe := _gebb(r); len(_gefe) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _gefe...)
	}
	if _fbcc := _dcab(r); _fbcc != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _fbcc)
	}
	if _eggd := _ccedd(r); len(_eggd) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _eggd...)
	}
	if _addd := _baba(r); len(_addd) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _addd...)
	}
	if _aafb := _cgaag(r); len(_aafb) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _aafb...)
	}
	if _fbac := _bdcd(r); _fbac != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _fbac)
	}
	if _fagg := _eebc(r); len(_fagg) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _fagg...)
	}
	if _aaec := _cbcec(r); len(_aaec) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _aaec...)
	}
	if _afb := _eceg(r); _afb != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _afb)
	}
	if _gaeb := _cadd(r, false); len(_gaeb) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _gaeb...)
	}
	if _ddfe := _ddcf(r); len(_ddfe) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _ddfe...)
	}
	if _ffd := _edff(r); _ffd != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _ffd)
	}
	if _cbgc := _dggg(r); _cbgc != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _cbgc)
	}
	if _fbgf := _ceag(r); _fbgf != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _fbgf)
	}
	if _adcgb := _cede(r); _adcgb != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _adcgb)
	}
	if _cebb := _fcbc(r); _cebb != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _cebb)
	}
	if _ddbg := _ddfb(r); len(_ddbg) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _ddbg...)
	}
	if _bdbd := _bgdb(r, _fbaf._geg); len(_bdbd) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _bdbd...)
	}
	if _ceeeg := _ccac(r); len(_ceeeg) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _ceeeg...)
	}
	if _bdgd := _ddbea(r); _bdgd != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _bdgd)
	}
	if _bafc := _ggce(r); _bafc != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _bafc)
	}
	if _gabe := _bbbb(r); len(_gabe) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _gabe...)
	}
	if _agbd := _bbfe(r); len(_agbd) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _agbd...)
	}
	if _agf := _dfed(r); _agf != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _agf)
	}
	if _feed := _begf(r); _feed != _ce {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _feed)
	}
	if _bced := _bcfb(r, _fbaf._geg, false); len(_bced) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _bced...)
	}
	if _fbaf._geg == _fdc() {
		if _gaed := _efgc(r); len(_gaed) != 0 {
			_eede.ViolatedRules = append(_eede.ViolatedRules, _gaed...)
		}
	}
	if _debd := _edee(r); len(_debd) != 0 {
		_eede.ViolatedRules = append(_eede.ViolatedRules, _debd...)
	}
	if len(_eede.ViolatedRules) > 0 {
		_a.Slice(_eede.ViolatedRules, func(_beee, _acf int) bool {
			return _eede.ViolatedRules[_beee].RuleNo < _eede.ViolatedRules[_acf].RuleNo
		})
		return _eede
	}
	return nil
}

func _abcdb(_caee *_db.CompliancePdfReader) (_dgeb ViolatedRule) {
	_efdfd, _fdbfb := _eagdc(_caee)
	if !_fdbfb {
		return _ce
	}
	if _efdfd.Get("\u0041\u0041") != nil {
		return _fdbe("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0054h\u0065\u0020\u0064\u006fc\u0075m\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020a\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079 \u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079\u002e")
	}
	return _ce
}

func _dfed(_bdgaa *_db.CompliancePdfReader) (_ebbf ViolatedRule) {
	_bebdfb, _edcab := _eagdc(_bdgaa)
	if !_edcab {
		return _ce
	}
	_ddgc, _edcab := _cb.GetDict(_bebdfb.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_edcab {
		return _ce
	}
	_dccc, _edcab := _cb.GetArray(_ddgc.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_edcab {
		return _ce
	}
	for _aded := 0; _aded < _dccc.Len(); _aded++ {
		_addf, _bcgaa := _cb.GetDict(_dccc.Get(_aded))
		if !_bcgaa {
			continue
		}
		if _addf.Get("\u0041\u0041") != nil {
			return _fdbe("\u0036.\u0036\u002e\u0032\u002d\u0032", "\u0041\u0020F\u0069\u0065\u006cd\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079 s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061n\u0020A\u0041\u0020\u0065\u006e\u0074\u0072y f\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069on\u0061l\u002d\u0061\u0063\u0074i\u006fn\u0073 \u0064\u0069c\u0074\u0069on\u0061\u0072\u0079\u002e")
		}
	}
	return _ce
}

// Profile1A is the implementation of the PDF/A-1A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1A struct{ profile1 }

func _cdf(_bgc *_f.Document) error {
	_fca := map[string]*_cb.PdfObjectDictionary{}
	_dfb := _ee.NewFinder(&_ee.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066"}})
	_gcba := map[_cb.PdfObject]struct{}{}
	_bfbg := map[_cb.PdfObject]struct{}{}
	for _, _agb := range _bgc.Objects {
		_fcf, _fdbf := _cb.GetDict(_agb)
		if !_fdbf {
			continue
		}
		_ab := _fcf.Get("\u0054\u0079\u0070\u0065")
		if _ab == nil {
			continue
		}
		if _deeg, _dfdg := _cb.GetName(_ab); _dfdg && _deeg.String() != "\u0046\u006f\u006e\u0074" {
			continue
		}
		if _, _gad := _gcba[_agb]; _gad {
			continue
		}
		_beae, _ede := _db.NewPdfFontFromPdfObject(_fcf)
		if _ede != nil {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
			return _ede
		}
		_eeed, _ede := _beae.GetFontDescriptor()
		if _ede != nil {
			return _ede
		}
		if _eeed != nil && (_eeed.FontFile != nil || _eeed.FontFile2 != nil || _eeed.FontFile3 != nil) {
			continue
		}
		_bad := _beae.BaseFont()
		if _bad == "" {
			return _b.Errorf("\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0073\u0020\u0073\u0079\u006e\u0074\u0061\u0078\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069d\u0020\u002d\u0020\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u0075\u006ed\u0065\u0066\u0069n\u0065\u0064\u003a\u0020\u0025\u0073", _fcf.String())
		}
		_dc, _ecf := _fca[_bad]
		if !_ecf {
			if len(_bad) > 7 && _bad[6] == '+' {
				_bad = _bad[7:]
			}
			_ccd := []string{_bad, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
			for _, _cdff := range _ccd {
				_eg.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _cdff)
				if _dc, _ecf = _fca[_cdff]; _ecf {
					break
				}
				_dbe := _dfb.Match(_cdff)
				if _dbe == nil {
					_eg.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _cdff)
					continue
				}
				_fcgg, _gbg := _db.NewPdfFontFromTTFFile(_dbe.Filename)
				if _gbg != nil {
					return _gbg
				}
				_fefd := _fcgg.FontDescriptor()
				if _fefd.FontFile != nil {
					if _, _ecf = _bfbg[_fefd.FontFile]; !_ecf {
						_bgc.Objects = append(_bgc.Objects, _fefd.FontFile)
						_bfbg[_fefd.FontFile] = struct{}{}
					}
				}
				if _fefd.FontFile2 != nil {
					if _, _ecf = _bfbg[_fefd.FontFile2]; !_ecf {
						_bgc.Objects = append(_bgc.Objects, _fefd.FontFile2)
						_bfbg[_fefd.FontFile2] = struct{}{}
					}
				}
				if _fefd.FontFile3 != nil {
					if _, _ecf = _bfbg[_fefd.FontFile3]; !_ecf {
						_bgc.Objects = append(_bgc.Objects, _fefd.FontFile3)
						_bfbg[_fefd.FontFile3] = struct{}{}
					}
				}
				_gag, _egc := _fcgg.ToPdfObject().(*_cb.PdfIndirectObject)
				if !_egc {
					_eg.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
					continue
				}
				_cdd, _egc := _gag.PdfObject.(*_cb.PdfObjectDictionary)
				if !_egc {
					_eg.Log.Debug("\u0046\u006fn\u0074\u0020\u0074\u0079p\u0065\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
					continue
				}
				_fca[_cdff] = _cdd
				_dc = _cdd
				break
			}
			if _dc == nil {
				_eg.Log.Debug("\u004e\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006eg\u0020\u0066\u006f\u006e\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u003a\u0020\u0025\u0073", _beae.BaseFont())
				return _ea.New("\u006e\u006f m\u0061\u0074\u0063h\u0069\u006e\u0067\u0020fon\u0074 f\u006f\u0075\u006e\u0064\u0020\u0069\u006e t\u0068\u0065\u0020\u0073\u0079\u0073\u0074e\u006d")
			}
		}
		for _, _bcgb := range _dc.Keys() {
			_fcf.Set(_bcgb, _dc.Get(_bcgb))
		}
		_fbf := _dc.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _fbf != nil {
			if _, _ecf = _bfbg[_fbf]; !_ecf {
				_bgc.Objects = append(_bgc.Objects, _fbf)
				_bfbg[_fbf] = struct{}{}
			}
		}
		_gcba[_agb] = struct{}{}
		_dca := _fcf.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
		if _dca != nil {
			_bgc.Objects = append(_bgc.Objects, _dca)
			_bfbg[_dca] = struct{}{}
		}
	}
	return nil
}

func _cafg(_acfc *_db.PdfFont, _cbdcf *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_dece = "\u0036.\u0033\u002e\u0037\u002d\u0033"
		_fffg = "\u0046\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0027\u0020\u0022\u0063\u006d\u0061\u0070\u0022\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0073\u0079\u006d\u0062o\u006c\u0069c\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0065\u0078\u0061\u0063\u0074\u006cy\u0020\u006f\u006ee\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u002e"
	)
	var _adga string
	if _dfbbe, _cbcae := _cb.GetName(_cbdcf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cbcae {
		_adga = _dfbbe.String()
	}
	if _adga != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _ce
	}
	_bcdd := _acfc.FontDescriptor()
	_bfdcd, _fbce := _cb.GetIntVal(_bcdd.Flags)
	if !_fbce {
		_eg.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _fdbe(_dece, _fffg)
	}
	_baga := (uint32(_bfdcd) >> 3) != 0
	if !_baga {
		return _ce
	}
	return _ce
}

func _dddg(_gefag *_db.CompliancePdfReader) (_dccef ViolatedRule) {
	for _, _gdbdb := range _gefag.GetObjectNums() {
		_ffbf, _bfcge := _gefag.GetIndirectObjectByNumber(_gdbdb)
		if _bfcge != nil {
			continue
		}
		_gbfe, _efaeb := _cb.GetStream(_ffbf)
		if !_efaeb {
			continue
		}
		_agefb, _efaeb := _cb.GetName(_gbfe.Get("\u0054\u0079\u0070\u0065"))
		if !_efaeb {
			continue
		}
		if *_agefb != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_, _efaeb = _cb.GetName(_gbfe.Get("\u004f\u0050\u0049"))
		if _efaeb {
			return _fdbe("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		_gagb, _efaeb := _cb.GetName(_gbfe.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_efaeb {
			continue
		}
		if *_gagb == "\u0050\u0053" {
			return _fdbe("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _gbfe.Get("\u0050\u0053") != nil {
			return _fdbe("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _dccef
}

func _cadd(_bebb *_db.CompliancePdfReader, _acac bool) (_ccdb []ViolatedRule) {
	var _cedg, _fgbe, _edca, _bdede, _dbac, _ebdd, _gcbd bool
	_dgeac := func() bool { return _cedg && _fgbe && _edca && _bdede && _dbac && _ebdd && _gcbd }
	_fbca, _aead := _ffbc(_bebb)
	var _edef _ebb.ProfileHeader
	if _aead {
		_edef, _ = _ebb.ParseHeader(_fbca.DestOutputProfile)
	}
	var _gcdc bool
	_caec := map[_cb.PdfObject]struct{}{}
	var _aba func(_fege _db.PdfColorspace) bool
	_aba = func(_gcggb _db.PdfColorspace) bool {
		switch _bfdd := _gcggb.(type) {
		case *_db.PdfColorspaceDeviceGray:
			if !_ebdd {
				if !_aead {
					_gcdc = true
					_ccdb = append(_ccdb, _fdbe("\u0036.\u0032\u002e\u0033\u002d\u0034", "\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0075s\u0065\u0064\u0020\u006f\u006el\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006ce\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020O\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u002e"))
					_ebdd = true
					if _dgeac() {
						return true
					}
				}
			}
		case *_db.PdfColorspaceDeviceRGB:
			if !_bdede {
				if !_aead || _edef.ColorSpace != _ebb.ColorSpaceRGB {
					_gcdc = true
					_ccdb = append(_ccdb, _fdbe("\u0036.\u0032\u002e\u0033\u002d\u0032", "\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065 \u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074In\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020u\u0073es\u0020a\u006e\u0020\u0052\u0047\u0042\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u002e"))
					_bdede = true
					if _dgeac() {
						return true
					}
				}
			}
		case *_db.PdfColorspaceDeviceCMYK:
			if !_dbac {
				if !_aead || _edef.ColorSpace != _ebb.ColorSpaceCMYK {
					_gcdc = true
					_ccdb = append(_ccdb, _fdbe("\u0036.\u0032\u002e\u0033\u002d\u0033", "\u0044\u0065\u0076\u0069\u0063e\u0043\u004d\u0059\u004b \u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074h\u0065\u0020\u0066\u0069\u006ce \u0068\u0061\u0073\u0020\u0061 \u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068a\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0061\u006e \u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072\u0020s\u0070\u0061\u0063e\u002e"))
					_dbac = true
					if _dgeac() {
						return true
					}
				}
			}
		case *_db.PdfColorspaceICCBased:
			if !_edca || !_gcbd {
				_cfegd, _eadcg := _ebb.ParseHeader(_bfdd.Data)
				if _eadcg != nil {
					_eg.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _eadcg)
					_ccdb = append(_ccdb, func() ViolatedRule {
						return _fdbe("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035")
					}())
					_edca = true
					if _dgeac() {
						return true
					}
				}
				if !_edca {
					var _gffb, _cefa bool
					switch _cfegd.DeviceClass {
					case _ebb.DeviceClassPRTR, _ebb.DeviceClassMNTR, _ebb.DeviceClassSCNR, _ebb.DeviceClassSPAC:
					default:
						_gffb = true
					}
					switch _cfegd.ColorSpace {
					case _ebb.ColorSpaceRGB, _ebb.ColorSpaceCMYK, _ebb.ColorSpaceGRAY, _ebb.ColorSpaceLAB:
					default:
						_cefa = true
					}
					if _gffb || _cefa {
						_ccdb = append(_ccdb, _fdbe("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035"))
						_edca = true
						if _dgeac() {
							return true
						}
					}
				}
				if !_gcbd {
					_aaff, _ := _cb.GetStream(_bfdd.GetContainingPdfObject())
					if _aaff.Get("\u004e") == nil || (_bfdd.N == 1 && _cfegd.ColorSpace != _ebb.ColorSpaceGRAY) || (_bfdd.N == 3 && !(_cfegd.ColorSpace == _ebb.ColorSpaceRGB || _cfegd.ColorSpace == _ebb.ColorSpaceLAB)) || (_bfdd.N == 4 && _cfegd.ColorSpace != _ebb.ColorSpaceCMYK) {
						_ccdb = append(_ccdb, _fdbe("\u0036.\u0032\u002e\u0033\u002d\u0035", "\u0049\u0066\u0020a\u006e\u0020u\u006e\u0063\u0061\u006c\u0069\u0062\u0072a\u0074\u0065\u0064\u0020\u0063\u006fl\u006f\u0072 \u0073\u0070\u0061c\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u0069\u006c\u0065 \u0074\u0068\u0065\u006e \u0074\u0068\u0061\u0074 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041-\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020\u0069\u006e\u0020\u0036\u002e\u0032\u002e\u0032\u002e"))
						_gcbd = true
						if _dgeac() {
							return true
						}
					}
				}
			}
			if _bfdd.Alternate != nil {
				return _aba(_bfdd.Alternate)
			}
		}
		return false
	}
	for _, _bgacb := range _bebb.GetObjectNums() {
		_cffd, _ggfb := _bebb.GetIndirectObjectByNumber(_bgacb)
		if _ggfb != nil {
			continue
		}
		_ccbd, _gfbab := _cb.GetStream(_cffd)
		if !_gfbab {
			continue
		}
		_cedfd, _gfbab := _cb.GetName(_ccbd.Get("\u0054\u0079\u0070\u0065"))
		if !_gfbab || _cedfd.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_debe, _gfbab := _cb.GetName(_ccbd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_gfbab {
			continue
		}
		_caec[_ccbd] = struct{}{}
		switch _debe.String() {
		case "\u0049\u006d\u0061g\u0065":
			_dfgf, _fgcg := _db.NewXObjectImageFromStream(_ccbd)
			if _fgcg != nil {
				continue
			}
			_caec[_ccbd] = struct{}{}
			if _aba(_dfgf.ColorSpace) {
				return _ccdb
			}
		case "\u0046\u006f\u0072\u006d":
			_efgd, _ggeb := _cb.GetDict(_ccbd.Get("\u0047\u0072\u006fu\u0070"))
			if !_ggeb {
				continue
			}
			_gafe := _efgd.Get("\u0043\u0053")
			if _gafe == nil {
				continue
			}
			_cfcb, _gfgb := _db.NewPdfColorspaceFromPdfObject(_gafe)
			if _gfgb != nil {
				continue
			}
			if _aba(_cfcb) {
				return _ccdb
			}
		}
	}
	for _, _eeff := range _bebb.PageList {
		_cgeae, _aegaf := _eeff.GetContentStreams()
		if _aegaf != nil {
			continue
		}
		for _, _bgbg := range _cgeae {
			_bgdab, _cbdc := _df.NewContentStreamParser(_bgbg).Parse()
			if _cbdc != nil {
				continue
			}
			for _, _abee := range *_bgdab {
				if len(_abee.Params) > 1 {
					continue
				}
				switch _abee.Operand {
				case "\u0042\u0049":
					_ffda, _dcfg := _abee.Params[0].(*_df.ContentStreamInlineImage)
					if !_dcfg {
						continue
					}
					_accf, _cdbcc := _ffda.GetColorSpace(_eeff.Resources)
					if _cdbcc != nil {
						continue
					}
					if _aba(_accf) {
						return _ccdb
					}
				case "\u0044\u006f":
					_ggefg, _acga := _cb.GetName(_abee.Params[0])
					if !_acga {
						continue
					}
					_gfed, _cfff := _eeff.Resources.GetXObjectByName(*_ggefg)
					if _, _efedd := _caec[_gfed]; _efedd {
						continue
					}
					switch _cfff {
					case _db.XObjectTypeImage:
						_dgdf, _faba := _db.NewXObjectImageFromStream(_gfed)
						if _faba != nil {
							continue
						}
						_caec[_gfed] = struct{}{}
						if _aba(_dgdf.ColorSpace) {
							return _ccdb
						}
					case _db.XObjectTypeForm:
						_efab, _fbcg := _cb.GetDict(_gfed.Get("\u0047\u0072\u006fu\u0070"))
						if !_fbcg {
							continue
						}
						_bfaf, _fbcg := _cb.GetName(_efab.Get("\u0043\u0053"))
						if !_fbcg {
							continue
						}
						_fgcgc, _cgdac := _db.NewPdfColorspaceFromPdfObject(_bfaf)
						if _cgdac != nil {
							continue
						}
						_caec[_gfed] = struct{}{}
						if _aba(_fgcgc) {
							return _ccdb
						}
					}
				}
			}
		}
	}
	if !_gcdc {
		return _ccdb
	}
	if (_edef.DeviceClass == _ebb.DeviceClassPRTR || _edef.DeviceClass == _ebb.DeviceClassMNTR) && (_edef.ColorSpace == _ebb.ColorSpaceRGB || _edef.ColorSpace == _ebb.ColorSpaceCMYK || _edef.ColorSpace == _ebb.ColorSpaceGRAY) {
		return _ccdb
	}
	if !_acac {
		return _ccdb
	}
	_adae, _bceb := _eagdc(_bebb)
	if !_bceb {
		return _ccdb
	}
	_afcb, _bceb := _cb.GetArray(_adae.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_bceb {
		_ccdb = append(_ccdb, _fdbe("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e"), _fdbe("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
		return _ccdb
	}
	if _afcb.Len() > 1 {
		_efbef := map[*_cb.PdfObjectDictionary]struct{}{}
		for _ffcd := 0; _ffcd < _afcb.Len(); _ffcd++ {
			_cceddc, _bdee := _cb.GetDict(_afcb.Get(_ffcd))
			if !_bdee {
				continue
			}
			if _ffcd == 0 {
				_efbef[_cceddc] = struct{}{}
				continue
			}
			if _, _dbea := _efbef[_cceddc]; !_dbea {
				_ccdb = append(_ccdb, _fdbe("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
				break
			}
		}
	}
	return _ccdb
}

func _bdca(_bacba *_db.CompliancePdfReader) (_geee []ViolatedRule) {
	_ebdf, _gbggd := _eagdc(_bacba)
	if !_gbggd {
		return _geee
	}
	_cfcbd, _gbggd := _cb.GetDict(_ebdf.Get("\u0050\u0065\u0072m\u0073"))
	if !_gbggd {
		return _geee
	}
	_cadb := _cfcbd.Keys()
	for _, _ebdba := range _cadb {
		if _ebdba.String() != "\u0055\u0052\u0033" && _ebdba.String() != "\u0044\u006f\u0063\u004d\u0044\u0050" {
			_geee = append(_geee, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "\u004e\u006f\u0020\u006b\u0065\u0079\u0073 \u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0055\u0052\u0033 \u0061n\u0064\u0020\u0044\u006f\u0063\u004dD\u0050\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u002e"))
		}
	}
	return _geee
}

func _dce(_cfeac *_f.Document) error {
	_agg, _ebbc := _cfeac.GetPages()
	if !_ebbc {
		return nil
	}
	for _, _dacc := range _agg {
		_gbfd, _dfbb := _cb.GetArray(_dacc.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_dfbb {
			continue
		}
		for _, _cbfe := range _gbfd.Elements() {
			_cbfe = _cb.ResolveReference(_cbfe)
			if _, _edd := _cbfe.(*_cb.PdfObjectNull); _edd {
				continue
			}
			_ebaa, _ecca := _cb.GetDict(_cbfe)
			if !_ecca {
				continue
			}
			_adbb, _ := _cb.GetIntVal(_ebaa.Get("\u0046"))
			_adbb &= ^(1 << 0)
			_adbb &= ^(1 << 1)
			_adbb &= ^(1 << 5)
			_adbb |= 1 << 2
			_ebaa.Set("\u0046", _cb.MakeInteger(int64(_adbb)))
			_bef := false
			if _bbd := _ebaa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _bbd != nil {
				_fcag, _fcfd := _cb.GetName(_bbd)
				if _fcfd && _fcag.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_bef = true
					if _ebaa.Get("\u0041\u0041") != nil {
						_ebaa.Remove("\u0041\u0041")
					}
				}
			}
			if _ebaa.Get("\u0043") != nil || _ebaa.Get("\u0049\u0043") != nil {
				_cgbd, _cfdd := _gfdd(_cfeac)
				if !_cfdd {
					_ebaa.Remove("\u0043")
					_ebaa.Remove("\u0049\u0043")
				} else {
					_fbfd, _fdca := _cb.GetIntVal(_cgbd.Get("\u004e"))
					if !_fdca || _fbfd != 3 {
						_ebaa.Remove("\u0043")
						_ebaa.Remove("\u0049\u0043")
					}
				}
			}
			_caf, _ecca := _cb.GetDict(_ebaa.Get("\u0041\u0050"))
			if _ecca {
				_agdg := _caf.Get("\u004e")
				if _agdg == nil {
					continue
				}
				if len(_caf.Keys()) > 1 {
					_caf.Clear()
					_caf.Set("\u004e", _agdg)
				}
				if _bef {
					_dbbc, _fagd := _cb.GetName(_ebaa.Get("\u0046\u0054"))
					if _fagd && *_dbbc == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}

type documentImages struct {
	_acda, _dbd, _fa bool
	_ca              map[_cb.PdfObject]struct{}
	_bee             []*imageInfo
}

func _dbad(_gefd *_db.PdfFont, _eae *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_egace = "\u0036.\u0033\u002e\u0037\u002d\u0032"
		_bdcb  = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _ecdd string
	if _fgag, _cfcc := _cb.GetName(_eae.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cfcc {
		_ecdd = _fgag.String()
	}
	if _ecdd != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _ce
	}
	_bbaa := _gefd.FontDescriptor()
	_defef, _egbb := _cb.GetIntVal(_bbaa.Flags)
	if !_egbb {
		_eg.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _fdbe(_egace, _bdcb)
	}
	_bdfd := (uint32(_defef) >> 3) & 1
	_gdcg := _bdfd != 0
	if !_gdcg {
		return _ce
	}
	if _eae.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _fdbe(_egace, _bdcb)
	}
	return _ce
}

func _ged(_aaac *_f.Document, _feg bool) error {
	_aeg, _fdd := _aaac.GetPages()
	if !_fdd {
		return nil
	}
	for _, _fbff := range _aeg {
		_bcf, _bgf := _cb.GetArray(_fbff.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_bgf {
			continue
		}
		for _, _gfb := range _bcf.Elements() {
			_cgda, _ffag := _cb.GetDict(_gfb)
			if !_ffag {
				continue
			}
			_egcc := _cgda.Get("\u0043")
			if _egcc == nil {
				continue
			}
			_adf, _ffag := _cb.GetArray(_egcc)
			if !_ffag {
				continue
			}
			_gcbf, _fae := _adf.GetAsFloat64Slice()
			if _fae != nil {
				return _fae
			}
			switch _adf.Len() {
			case 0, 1:
				if _feg {
					_cgda.Set("\u0043", _cb.MakeArrayFromIntegers([]int{1, 1, 1, 1}))
				} else {
					_cgda.Set("\u0043", _cb.MakeArrayFromIntegers([]int{1, 1, 1}))
				}
			case 3:
				if _feg {
					_aadc, _gada, _dcfc, _bded := _e.RGBToCMYK(uint8(_gcbf[0]*255), uint8(_gcbf[1]*255), uint8(_gcbf[2]*255))
					_cgda.Set("\u0043", _cb.MakeArrayFromFloats([]float64{float64(_aadc) / 255, float64(_gada) / 255, float64(_dcfc) / 255, float64(_bded) / 255}))
				}
			case 4:
				if !_feg {
					_bbcb, _ecd, _cfc := _e.CMYKToRGB(uint8(_gcbf[0]*255), uint8(_gcbf[1]*255), uint8(_gcbf[2]*255), uint8(_gcbf[3]*255))
					_cgda.Set("\u0043", _cb.MakeArrayFromFloats([]float64{float64(_bbcb) / 255, float64(_ecd) / 255, float64(_cfc) / 255}))
				}
			}
		}
	}
	return nil
}

func _cecag(_agaf *_db.CompliancePdfReader, _gbdg standardType) (_eefg []ViolatedRule) {
	var _ebaf, _dcdcb, _ebgcf, _deece, _dgggg, _gffbb, _bbed bool
	_abgaeb := func() bool { return _ebaf && _dcdcb && _ebgcf && _deece && _dgggg && _gffbb && _bbed }
	_fece := map[*_cb.PdfObjectStream]*_cd.CMap{}
	_dfdgf := map[*_cb.PdfObjectStream][]byte{}
	_febfe := map[_cb.PdfObject]*_db.PdfFont{}
	for _, _ddaff := range _agaf.GetObjectNums() {
		_fbfe, _bfeb := _agaf.GetIndirectObjectByNumber(_ddaff)
		if _bfeb != nil {
			continue
		}
		_afgf, _dedb := _cb.GetDict(_fbfe)
		if !_dedb {
			continue
		}
		_gabf, _dedb := _cb.GetName(_afgf.Get("\u0054\u0079\u0070\u0065"))
		if !_dedb {
			continue
		}
		if *_gabf != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_cdgbf, _bfeb := _db.NewPdfFontFromPdfObject(_afgf)
		if _bfeb != nil {
			_eg.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _bfeb)
			continue
		}
		_febfe[_afgf] = _cdgbf
	}
	for _, _efca := range _agaf.PageList {
		_dedae, _abef := _efca.GetContentStreams()
		if _abef != nil {
			_eg.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _eeefd := range _dedae {
			_eedggc := _df.NewContentStreamParser(_eeefd)
			_gbfg, _fegd := _eedggc.Parse()
			if _fegd != nil {
				_eg.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _fegd)
				continue
			}
			var _bebfb bool
			for _, _badb := range *_gbfg {
				if _badb.Operand != "\u0054\u0072" {
					continue
				}
				if len(_badb.Params) != 1 {
					_eg.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_badb.Params))
					continue
				}
				_acaf, _agddcc := _cb.GetIntVal(_badb.Params[0])
				if !_agddcc {
					_eg.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _acaf == 3 {
					_bebfb = true
					break
				}
			}
			for _, _dbbd := range *_gbfg {
				if _dbbd.Operand != "\u0054\u0066" {
					continue
				}
				if len(_dbbd.Params) != 2 {
					_eg.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_dbbd.Params))
					continue
				}
				_bfgca, _dbfd := _cb.GetName(_dbbd.Params[0])
				if !_dbfd {
					_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _dbbd)
					continue
				}
				_cegg, _gbceb := _efca.Resources.GetFontByName(*_bfgca)
				if !_gbceb {
					_eg.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_fdfdg, _dbfd := _cb.GetDict(_cegg)
				if !_dbfd {
					_eg.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_daggg, _dbfd := _febfe[_fdfdg]
				if !_dbfd {
					var _ddab error
					_daggg, _ddab = _db.NewPdfFontFromPdfObject(_fdfdg)
					if _ddab != nil {
						_eg.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _ddab)
						continue
					}
					_febfe[_fdfdg] = _daggg
				}
				if !_ebaf {
					_acff := _gdac(_fdfdg, _dfdgf, _fece)
					if _acff != _ce {
						_eefg = append(_eefg, _acff)
						_ebaf = true
						if _abgaeb() {
							return _eefg
						}
					}
				}
				if !_dcdcb {
					_fcgfe := _adge(_fdfdg)
					if _fcgfe != _ce {
						_eefg = append(_eefg, _fcgfe)
						_dcdcb = true
						if _abgaeb() {
							return _eefg
						}
					}
				}
				if !_ebgcf {
					_ccgd := _eacc(_fdfdg, _dfdgf, _fece)
					if _ccgd != _ce {
						_eefg = append(_eefg, _ccgd)
						_ebgcf = true
						if _abgaeb() {
							return _eefg
						}
					}
				}
				if !_deece {
					_acaec := _adee(_fdfdg, _dfdgf, _fece)
					if _acaec != _ce {
						_eefg = append(_eefg, _acaec)
						_deece = true
						if _abgaeb() {
							return _eefg
						}
					}
				}
				if !_dgggg {
					_dgeed := _dfcee(_daggg, _fdfdg, _bebfb)
					if _dgeed != _ce {
						_dgggg = true
						_eefg = append(_eefg, _dgeed)
						if _abgaeb() {
							return _eefg
						}
					}
				}
				if !_gffbb {
					_gfab := _ceeb(_daggg, _fdfdg)
					if _gfab != _ce {
						_gffbb = true
						_eefg = append(_eefg, _gfab)
						if _abgaeb() {
							return _eefg
						}
					}
				}
				if !_bbed && (_gbdg._fd == "\u0041" || _gbdg._fd == "\u0055") {
					_efgde := _bdgc(_fdfdg, _dfdgf, _fece)
					if _efgde != _ce {
						_bbed = true
						_eefg = append(_eefg, _efgde)
						if _abgaeb() {
							return _eefg
						}
					}
				}
			}
		}
	}
	return _eefg
}

func _bbbb(_abbg *_db.CompliancePdfReader) (_aeded []ViolatedRule) {
	var _eadge, _cdgeg, _ccfc, _dbgf, _fgce, _bbfc, _bgfa bool
	_febf := func() bool { return _eadge && _cdgeg && _ccfc && _dbgf && _fgce && _bbfc && _bgfa }
	for _, _dgdb := range _abbg.PageList {
		_bbdf, _cded := _dgdb.GetAnnotations()
		if _cded != nil {
			_eg.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _cded)
			continue
		}
		for _, _acfeb := range _bbdf {
			if !_eadge {
				switch _acfeb.GetContext().(type) {
				case *_db.PdfAnnotationFileAttachment, *_db.PdfAnnotationSound, *_db.PdfAnnotationMovie, nil:
					_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020p\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e\u0020\u0041d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020F\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063\u0068\u006de\u006e\u0074\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u0020\u0061\u006e\u0064\u0020\u004d\u006f\u0076\u0069e\u0020\u0074\u0079\u0070\u0065s \u0073ha\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_eadge = true
					if _febf() {
						return _aeded
					}
				}
			}
			_fbafc, _ccfca := _cb.GetDict(_acfeb.GetContainingPdfObject())
			if !_ccfca {
				continue
			}
			if !_cdgeg {
				_eabg, _bdbdd := _cb.GetFloatVal(_fbafc.Get("\u0043\u0041"))
				if _bdbdd && _eabg != 1.0 {
					_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0033\u002d\u0031", "\u0041\u006e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073h\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0043\u0041\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u002e\u0030\u002e"))
					_cdgeg = true
					if _febf() {
						return _aeded
					}
				}
			}
			if !_ccfc {
				_bgfg, _gfgg := _cb.GetIntVal(_fbafc.Get("\u0046"))
				if !(_gfgg && _bgfg&4 == 4 && _bgfg&1 == 0 && _bgfg&2 == 0 && _bgfg&32 == 0) {
					_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0033\u002d\u0032", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0074\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u002e\u0020\u0054\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065 s\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020I\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u0020\u0061\u006e\u0064\u0020\u004e\u006f\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020b\u0069\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073e\u0074\u0020t\u006f\u0020\u0030\u002e"))
					_ccfc = true
					if _febf() {
						return _aeded
					}
				}
			}
			if !_dbgf {
				_egce, _ebfcd := _cb.GetDict(_fbafc.Get("\u0041\u0050"))
				if _ebfcd {
					_begg := _egce.Get("\u004e")
					if _begg == nil || len(_egce.Keys()) > 1 {
						_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_dbgf = true
						if _febf() {
							return _aeded
						}
						continue
					}
					_, _gdgg := _acfeb.GetContext().(*_db.PdfAnnotationWidget)
					if _gdgg {
						_fcda, _eafg := _cb.GetName(_fbafc.Get("\u0046\u0054"))
						if _eafg && *_fcda == "\u0042\u0074\u006e" {
							if _, _fage := _cb.GetDict(_begg); !_fage {
								_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_dbgf = true
								if _febf() {
									return _aeded
								}
								continue
							}
						}
					}
					_, _agbb := _cb.GetStream(_begg)
					if !_agbb {
						_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_dbgf = true
						if _febf() {
							return _aeded
						}
						continue
					}
				}
			}
			if !_fgce {
				if _fbafc.Get("\u0043") != nil || _fbafc.Get("\u0049\u0043") != nil {
					_cade, _afdd := _eaff(_abbg)
					if !_afdd {
						_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
						_fgce = true
						if _febf() {
							return _aeded
						}
					} else {
						_bcgf, _cbdg := _cb.GetIntVal(_cade.Get("\u004e"))
						if !_cbdg || _bcgf != 3 {
							_aeded = append(_aeded, _fdbe("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
							_fgce = true
							if _febf() {
								return _aeded
							}
						}
					}
				}
			}
			_befd, _dfga := _acfeb.GetContext().(*_db.PdfAnnotationWidget)
			if !_dfga {
				continue
			}
			if !_bbfc {
				if _befd.A != nil {
					_aeded = append(_aeded, _fdbe("\u0036.\u0036\u002e\u0031\u002d\u0033", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_bbfc = true
					if _febf() {
						return _aeded
					}
				}
			}
			if !_bgfa {
				if _befd.AA != nil {
					_aeded = append(_aeded, _fdbe("\u0036.\u0036\u002e\u0032\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_bgfa = true
					if _febf() {
						return _aeded
					}
				}
			}
		}
	}
	return _aeded
}

type documentColorspaceOptimizeFunc func(_gbb *_f.Document, _dfg []*_f.Image) error

func _eebc(_fbda *_db.CompliancePdfReader) (_cfgd []ViolatedRule) {
	_bgb := _fbda.GetObjectNums()
	for _, _bfgd := range _bgb {
		_ddgg, _cegd := _fbda.GetIndirectObjectByNumber(_bfgd)
		if _cegd != nil {
			continue
		}
		_gfdc, _beeeg := _cb.GetDict(_ddgg)
		if !_beeeg {
			continue
		}
		_abgd, _beeeg := _cb.GetName(_gfdc.Get("\u0054\u0079\u0070\u0065"))
		if !_beeeg {
			continue
		}
		if _abgd.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _gfdc.Get("\u0045\u0046") != nil {
			_cfgd = append(_cfgd, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0031", "\u0041 \u0066\u0069\u006c\u0065 \u0073p\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046 \u0033\u002e\u0031\u0030\u002e\u0032\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0045\u0046 \u006be\u0079\u002e"))
			break
		}
	}
	_dege, _bgea := _eagdc(_fbda)
	if !_bgea {
		return _cfgd
	}
	_efed, _bgea := _cb.GetDict(_dege.Get("\u004e\u0061\u006de\u0073"))
	if !_bgea {
		return _cfgd
	}
	if _efed.Get("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065\u0073") != nil {
		_cfgd = append(_cfgd, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0032", "\u0041\u0020\u0066i\u006c\u0065\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020d\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020PD\u0046 \u0052\u0065\u0066er\u0065\u006e\u0063\u0065\u0020\u0033\u002e6\u002e\u0033\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0045m\u0062\u0065\u0064\u0064\u0065\u0064\u0046i\u006c\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
	}
	return _cfgd
}

func _gbea(_efgdc *_db.CompliancePdfReader) ViolatedRule {
	_bcedc := _efgdc.ParserMetadata().HeaderCommentBytes()
	if _bcedc[0] > 127 && _bcedc[1] > 127 && _bcedc[2] > 127 && _bcedc[3] > 127 {
		return _ce
	}
	return _fdbe("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}

var _ Profile = (*Profile2A)(nil)

func _bdcd(_badd *_db.CompliancePdfReader) ViolatedRule {
	for _, _gcag := range _badd.PageList {
		_bcc := _gcag.GetContentStreamObjs()
		for _, _fgfg := range _bcc {
			_fgfg = _cb.TraceToDirectObject(_fgfg)
			var _aacaa string
			switch _cfbfd := _fgfg.(type) {
			case *_cb.PdfObjectString:
				_aacaa = _cfbfd.Str()
			case *_cb.PdfObjectStream:
				_acbg, _gacc := _cb.GetName(_cb.TraceToDirectObject(_cfbfd.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
				if _gacc {
					if *_acbg == _cb.StreamEncodingFilterNameLZW {
						return _fdbe("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
					}
				}
				_abdef, _ceae := _cb.DecodeStream(_cfbfd)
				if _ceae != nil {
					_eg.Log.Debug("\u0045r\u0072\u003a\u0020\u0025\u0076", _ceae)
					continue
				}
				_aacaa = string(_abdef)
			default:
				_eg.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063t\u003a\u0020\u0025\u0054", _fgfg)
				continue
			}
			_eeeg := _df.NewContentStreamParser(_aacaa)
			_eaad, _daccf := _eeeg.Parse()
			if _daccf != nil {
				_eg.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d:\u0020\u0025\u0076", _daccf)
				continue
			}
			for _, _eebf := range *_eaad {
				if !(_eebf.Operand == "\u0042\u0049" && len(_eebf.Params) == 1) {
					continue
				}
				_bbcbf, _ecac := _eebf.Params[0].(*_df.ContentStreamInlineImage)
				if !_ecac {
					continue
				}
				_bgec, _eacgb := _bbcbf.GetEncoder()
				if _eacgb != nil {
					_eg.Log.Debug("\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006ee\u0020\u0069\u006d\u0061\u0067\u0065 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _eacgb)
					continue
				}
				if _bgec.GetFilterName() == _cb.StreamEncodingFilterNameLZW {
					return _fdbe("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
				}
			}
		}
	}
	return _ce
}

func _afcc(_fcaca *_db.CompliancePdfReader) (_bdgf ViolatedRule) {
	_cdage, _dacfg := _eagdc(_fcaca)
	if !_dacfg {
		return _ce
	}
	if _cdage.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		return _fdbe("\u0036\u002e\u0031\u0031\u002d\u0031", "Th\u0065\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063a\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020R\u0065q\u0075\u0069\u0072\u0065\u006d\u0065\u006e\u0074s\u0020k\u0065\u0079.")
	}
	return _ce
}

func _fgg(_cedf *_f.Document, _gdf int) error {
	for _, _aaa := range _cedf.Objects {
		_cedd, _ecc := _cb.GetDict(_aaa)
		if !_ecc {
			continue
		}
		_aea := _cedd.Get("\u0054\u0079\u0070\u0065")
		if _aea == nil {
			continue
		}
		if _dcf, _egfb := _cb.GetName(_aea); _egfb && _dcf.String() != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_dabb, _gcd := _cb.GetName(_cedd.Get("\u0053"))
		if !_gcd {
			continue
		}
		switch _db.PdfActionType(*_dabb) {
		case _db.ActionTypeLaunch, _db.ActionTypeSound, _db.ActionTypeMovie, _db.ActionTypeResetForm, _db.ActionTypeImportData, _db.ActionTypeJavaScript:
			_cedd.Remove("\u0053")
		case _db.ActionTypeHide, _db.ActionTypeSetOCGState, _db.ActionTypeRendition, _db.ActionTypeTrans, _db.ActionTypeGoTo3DView:
			if _gdf == 2 {
				_cedd.Remove("\u0053")
			}
		case _db.ActionTypeNamed:
			_fbad, _eaf := _cb.GetName(_cedd.Get("\u004e"))
			if !_eaf {
				continue
			}
			switch *_fbad {
			case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
			default:
				_cedd.Remove("\u004e")
			}
		}
	}
	return nil
}

func _adec(_gcea *_db.CompliancePdfReader) ViolatedRule {
	_edbca := _gcea.ParserMetadata()
	if _edbca.HasInvalidSeparationAfterXRef() {
		return _fdbe("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _ce
}

func _cgaa(_adbe *_f.Document) {
	if _adbe.ID[0] != "" && _adbe.ID[1] != "" {
		return
	}
	_adbe.UseHashBasedID = true
}

// NewProfile1A creates a new Profile1A with given options.
func NewProfile1A(options *Profile1Options) *Profile1A {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_cgc(options)
	return &Profile1A{profile1{_fbg: *options, _geg: _fdc()}}
}

// DefaultProfile1Options are the default options for the Profile1.
func DefaultProfile1Options() *Profile1Options {
	return &Profile1Options{Now: _c.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}

func _edee(_bacga *_db.CompliancePdfReader) (_acce []ViolatedRule) {
	for _, _fdeb := range _bacga.GetObjectNums() {
		_fffgf, _cdeg := _bacga.GetIndirectObjectByNumber(_fdeb)
		if _cdeg != nil {
			continue
		}
		_dcfb, _adaa := _cb.GetDict(_fffgf)
		if !_adaa {
			continue
		}
		_cegeb, _adaa := _cb.GetName(_dcfb.Get("\u0054\u0079\u0070\u0065"))
		if !_adaa {
			continue
		}
		if _cegeb.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_dfba, _adaa := _cb.GetBool(_dcfb.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if !_adaa {
			return _acce
		}
		if bool(*_dfba) {
			_acce = append(_acce, _fdbe("\u0036\u002e\u0039-\u0031", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
	}
	return _acce
}

// ApplyStandard tries to change the content of the writer to match the PDF/A-2 standard.
// Implements model.StandardApplier.
func (_aggg *profile2) ApplyStandard(document *_f.Document) (_acea error) {
	_dcd(document, 7)
	if _acea = _cfb(document, _aggg._fdbb.Now); _acea != nil {
		return _acea
	}
	if _acea = _ecad(document); _acea != nil {
		return _acea
	}
	_gdge, _gfffe := _bgcc(_aggg._fdbb.CMYKDefaultColorSpace, _aggg._dgfgd)
	_acea = _eef(document, []pageColorspaceOptimizeFunc{_gdge}, []documentColorspaceOptimizeFunc{_gfffe})
	if _acea != nil {
		return _acea
	}
	_cgaa(document)
	if _acea = _eddc(document); _acea != nil {
		return _acea
	}
	if _acea = _daf(document, _aggg._dgfgd._ed); _acea != nil {
		return _acea
	}
	if _acea = _fgdac(document); _acea != nil {
		return _acea
	}
	if _acea = _gfba(document); _acea != nil {
		return _acea
	}
	if _acea = _cdf(document); _acea != nil {
		return _acea
	}
	if _acea = _afae(document); _acea != nil {
		return _acea
	}
	if _aggg._dgfgd._fd == "\u0041" {
		_cbce(document)
	}
	if _acea = _fgg(document, _aggg._dgfgd._ed); _acea != nil {
		return _acea
	}
	if _acea = _gbc(document); _acea != nil {
		return _acea
	}
	if _aedb := _fed(document, _aggg._dgfgd, _aggg._fdbb.Xmp); _aedb != nil {
		return _aedb
	}
	if _aggg._dgfgd == _ad() {
		if _acea = _gd(document); _acea != nil {
			return _acea
		}
	}
	if _acea = _bgdg(document); _acea != nil {
		return _acea
	}
	if _acea = _aege(document); _acea != nil {
		return _acea
	}
	if _acea = _gafb(document); _acea != nil {
		return _acea
	}
	return nil
}

func _dfbag(_fdfa *_db.CompliancePdfReader) (_ffcb []ViolatedRule) {
	var _bedb, _abdaa, _becf, _fgbc, _gdggg, _ccea, _efcd bool
	_dddda := func() bool { return _bedb && _abdaa && _becf && _fgbc && _gdggg && _ccea && _efcd }
	for _, _cfdg := range _fdfa.PageList {
		_dadb, _bfcf := _cfdg.GetAnnotations()
		if _bfcf != nil {
			_eg.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _bfcf)
			continue
		}
		for _, _gecg := range _dadb {
			if !_bedb {
				switch _gecg.GetContext().(type) {
				case *_db.PdfAnnotationScreen, *_db.PdfAnnotation3D, *_db.PdfAnnotationSound, *_db.PdfAnnotationMovie, nil:
					_ffcb = append(_ffcb, _fdbe("\u0036.\u0033\u002e\u0031\u002d\u0031", "\u0041nn\u006f\u0074\u0061\u0074i\u006f\u006e t\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065r\u006d\u0069t\u0074\u0065\u0064\u002e\u0020\u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0033\u0044\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u0053\u0063\u0072\u0065\u0065\u006e\u0020\u0061n\u0064\u0020\u004d\u006f\u0076\u0069\u0065\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bedb = true
					if _dddda() {
						return _ffcb
					}
				}
			}
			_dgcc, _gafd := _cb.GetDict(_gecg.GetContainingPdfObject())
			if !_gafd {
				continue
			}
			_, _bdfb := _gecg.GetContext().(*_db.PdfAnnotationPopup)
			if !_bdfb && !_abdaa {
				_, _baeeb := _cb.GetIntVal(_dgcc.Get("\u0046"))
				if !_baeeb {
					_ffcb = append(_ffcb, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0045\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072i\u0065\u0073\u0020\u0077\u0068\u006fs\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0069\u0073\u0020\u0050\u006f\u0070u\u0070\u002c\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065y."))
					_abdaa = true
					if _dddda() {
						return _ffcb
					}
				}
			}
			if !_becf {
				_fegc, _decd := _cb.GetIntVal(_dgcc.Get("\u0046"))
				if _decd && !(_fegc&4 == 4 && _fegc&1 == 0 && _fegc&2 == 0 && _fegc&32 == 0 && _fegc&256 == 0) {
					_ffcb = append(_ffcb, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0032", "I\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073\u0068\u0061l\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020\u0049\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u002c\u0020\u0054\u006f\u0067\u0067\u006c\u0065\u004e\u006f\u0056\u0069\u0065\u0077\u002c\u0020\u0061\u006e\u0064 \u004eo\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020s\u0065\u0074\u0020t\u006f\u0020\u0030."))
					_becf = true
					if _dddda() {
						return _ffcb
					}
				}
			}
			_, _effda := _gecg.GetContext().(*_db.PdfAnnotationText)
			if _effda && !_fgbc {
				_cgbb, _ffcga := _cb.GetIntVal(_dgcc.Get("\u0046"))
				if _ffcga && !(_cgbb&8 == 8 && _cgbb&16 == 16) {
					_ffcb = append(_ffcb, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0033", "\u0054\u0065\u0078\u0074\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069o\u006e\u0020\u0068\u0061\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006ca\u0067\u0073\u0020\u004e\u006f\u005a\u006f\u006f\u006d\u0020\u006f\u0072\u0020\u004e\u006f\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u0073\u0065t\u0020\u0074\u006f\u0020\u0030\u002e"))
					_fgbc = true
					if _dddda() {
						return _ffcb
					}
				}
			}
			if !_gdggg {
				_edgfa, _afegd := _cb.GetDict(_dgcc.Get("\u0041\u0050"))
				if _afegd {
					_fede := _edgfa.Get("\u004e")
					if _fede == nil || len(_edgfa.Keys()) > 1 {
						_ffcb = append(_ffcb, _fdbe("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_gdggg = true
						if _dddda() {
							return _ffcb
						}
						continue
					}
					_, _adfde := _gecg.GetContext().(*_db.PdfAnnotationWidget)
					if _adfde {
						_dbba, _adda := _cb.GetName(_dgcc.Get("\u0046\u0054"))
						if _adda && *_dbba == "\u0042\u0074\u006e" {
							if _, _gggfe := _cb.GetDict(_fede); !_gggfe {
								_ffcb = append(_ffcb, _fdbe("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_gdggg = true
								if _dddda() {
									return _ffcb
								}
								continue
							}
						}
					}
					_, _bbaed := _cb.GetStream(_fede)
					if !_bbaed {
						_ffcb = append(_ffcb, _fdbe("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_gdggg = true
						if _dddda() {
							return _ffcb
						}
						continue
					}
				}
			}
			_fbdb, _bbdc := _gecg.GetContext().(*_db.PdfAnnotationWidget)
			if !_bbdc {
				continue
			}
			if !_ccea {
				if _fbdb.A != nil {
					_ffcb = append(_ffcb, _fdbe("\u0036.\u0034\u002e\u0031\u002d\u0031", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_ccea = true
					if _dddda() {
						return _ffcb
					}
				}
			}
			if !_efcd {
				if _fbdb.AA != nil {
					_ffcb = append(_ffcb, _fdbe("\u0036.\u0034\u002e\u0031\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_efcd = true
					if _dddda() {
						return _ffcb
					}
				}
			}
		}
	}
	return _ffcb
}

type imageInfo struct {
	ColorSpace       _cb.PdfObjectName
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_cb.PdfObjectStream
	_fg              bool
}

func _agdc(_eacf *_db.CompliancePdfReader) ViolatedRule { return _ce }
func _ggb(_fee *_f.Document) error {
	_bfb := func(_cac *_cb.PdfObjectDictionary) error {
		if _aad := _cac.Get("\u0053\u004d\u0061s\u006b"); _aad != nil {
			_cac.Set("\u0053\u004d\u0061s\u006b", _cb.MakeName("\u004e\u006f\u006e\u0065"))
		}
		_gfe := _cac.Get("\u0043\u0041")
		if _gfe != nil {
			_egf, _gbf := _cb.GetNumberAsFloat(_gfe)
			if _gbf != nil {
				_eg.Log.Debug("\u0045x\u0074\u0047S\u0074\u0061\u0074\u0065 \u006f\u0062\u006ae\u0063\u0074\u0020\u0043\u0041\u0020\u0076\u0061\u006cue\u0020\u0069\u0073 \u006e\u006ft\u0020\u0061\u0020\u0066\u006c\u006fa\u0074\u003a \u0025\u0076", _gbf)
				_egf = 0
			}
			if _egf != 1.0 {
				_cac.Set("\u0043\u0041", _cb.MakeFloat(1.0))
			}
		}
		_gfe = _cac.Get("\u0063\u0061")
		if _gfe != nil {
			_cgd, _gggd := _cb.GetNumberAsFloat(_gfe)
			if _gggd != nil {
				_eg.Log.Debug("\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0027\u0063\u0061\u0027\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _gggd)
				_cgd = 0
			}
			if _cgd != 1.0 {
				_cac.Set("\u0063\u0061", _cb.MakeFloat(1.0))
			}
		}
		_ceee := _cac.Get("\u0042\u004d")
		if _ceee != nil {
			_cbf, _gff := _cb.GetName(_ceee)
			if !_gff {
				_eg.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_cbf = _cb.MakeName("")
			}
			_dda := _cbf.String()
			switch _dda {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
			default:
				_cac.Set("\u0042\u004d", _cb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		_ddd := _cac.Get("\u0054\u0052")
		if _ddd != nil {
			_eg.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_cac.Remove("\u0054\u0052")
		}
		_dge := _cac.Get("\u0054\u0052\u0032")
		if _dge != nil {
			_fdf := _dge.String()
			if _fdf != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_eg.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_cac.Set("\u0054\u0052\u0032", _cb.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		return nil
	}
	_efb, _bace := _fee.GetPages()
	if !_bace {
		return nil
	}
	for _, _dfdb := range _efb {
		_aae, _bgg := _dfdb.GetResources()
		if !_bgg {
			continue
		}
		_cdb, _gbfa := _cb.GetDict(_aae.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_gbfa {
			return nil
		}
		_ag := _cdb.Keys()
		for _, _aadb := range _ag {
			_efg, _dbc := _cb.GetDict(_cdb.Get(_aadb))
			if !_dbc {
				continue
			}
			_cgdd := _bfb(_efg)
			if _cgdd != nil {
				continue
			}
		}
	}
	for _, _efe := range _efb {
		_eadc, _cdg := _efe.GetContents()
		if !_cdg {
			return nil
		}
		for _, _ge := range _eadc {
			_cfe, _bggg := _ge.GetData()
			if _bggg != nil {
				continue
			}
			_efa := _df.NewContentStreamParser(string(_cfe))
			_cc, _bggg := _efa.Parse()
			if _bggg != nil {
				continue
			}
			for _, _eab := range *_cc {
				if len(_eab.Params) == 0 {
					continue
				}
				_, _aeb := _cb.GetName(_eab.Params[0])
				if !_aeb {
					continue
				}
				_ced, _eba := _efe.GetResourcesXObject()
				if !_eba {
					continue
				}
				for _, _ddc := range _ced.Keys() {
					_ccf, _gcb := _cb.GetStream(_ced.Get(_ddc))
					if !_gcb {
						continue
					}
					_fbd, _gcb := _cb.GetDict(_ccf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_gcb {
						continue
					}
					_fbaa, _gcb := _cb.GetDict(_fbd.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_gcb {
						continue
					}
					for _, _ddb := range _fbaa.Keys() {
						_cga, _fda := _cb.GetDict(_fbaa.Get(_ddb))
						if !_fda {
							continue
						}
						_gce := _bfb(_cga)
						if _gce != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}

// NewProfile1B creates a new Profile1B with the given options.
func NewProfile1B(options *Profile1Options) *Profile1B {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_cgc(options)
	return &Profile1B{profile1{_fbg: *options, _geg: _ebg()}}
}

func _gdbe(_cegc *_db.CompliancePdfReader) (_febe []ViolatedRule) {
	_gecf, _ccg := _eagdc(_cegc)
	if !_ccg {
		return _febe
	}
	_ffdg := _fdbe("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e")
	_dcbc, _ccg := _cb.GetArray(_gecf.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_ccg {
		_febe = append(_febe, _ffdg)
		return _febe
	}
	_gdfa := _fdbe("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e")
	if _dcbc.Len() > 1 {
		_bfgec := map[*_cb.PdfObjectDictionary]struct{}{}
		for _dcad := 0; _dcad < _dcbc.Len(); _dcad++ {
			_beda, _gfcc := _cb.GetDict(_dcbc.Get(_dcad))
			if !_gfcc {
				_febe = append(_febe, _ffdg)
				return _febe
			}
			if _dcad == 0 {
				_bfgec[_beda] = struct{}{}
				continue
			}
			if _, _dcega := _bfgec[_beda]; !_dcega {
				_febe = append(_febe, _gdfa)
				break
			}
		}
	} else if _dcbc.Len() == 0 {
		_febe = append(_febe, _ffdg)
		return _febe
	}
	_bcgae, _ccg := _cb.GetDict(_dcbc.Get(0))
	if !_ccg {
		_febe = append(_febe, _ffdg)
		return _febe
	}
	if _aeag, _fcce := _cb.GetName(_bcgae.Get("\u0053")); !_fcce || (*_aeag) != "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411" {
		_febe = append(_febe, _ffdg)
		return _febe
	}
	_feaae, _cdffa := _db.NewPdfOutputIntentFromPdfObject(_bcgae)
	if _cdffa != nil {
		_eg.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020i\u006et\u0065\u006e\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _cdffa)
		return _febe
	}
	_bgggg, _cdffa := _ebb.ParseHeader(_feaae.DestOutputProfile)
	if _cdffa != nil {
		_eg.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066\u0061i\u006c\u0065d\u003a\u0020\u0025\u0076", _cdffa)
		return _febe
	}
	if (_bgggg.DeviceClass == _ebb.DeviceClassPRTR || _bgggg.DeviceClass == _ebb.DeviceClassMNTR) && (_bgggg.ColorSpace == _ebb.ColorSpaceRGB || _bgggg.ColorSpace == _ebb.ColorSpaceCMYK || _bgggg.ColorSpace == _ebb.ColorSpaceGRAY) {
		return _febe
	}
	_febe = append(_febe, _ffdg)
	return _febe
}
func _ae() standardType { return standardType{_ed: 2, _fd: "\u0042"} }
func _gafb(_fff *_f.Document) error {
	_fdbee, _fdg := _fff.FindCatalog()
	if !_fdg {
		return _ea.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _fdbee.Object.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		_fdbee.Object.Remove("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073")
	}
	return nil
}
func (_fbb *documentImages) hasOnlyDeviceGray() bool { return _fbb._fa && !_fbb._acda && !_fbb._dbd }
func _ffbc(_agdd *_db.CompliancePdfReader) (*_db.PdfOutputIntent, bool) {
	_bgfde, _ccda := _eaff(_agdd)
	if !_ccda {
		return nil, false
	}
	_baec, _effd := _db.NewPdfOutputIntentFromPdfObject(_bgfde)
	if _effd != nil {
		return nil, false
	}
	return _baec, true
}

func _aefe(_gffgc *_db.PdfInfo, _bdec *_fcg.Document) bool {
	_cffdc, _ecdgd := _bdec.GetPdfInfo()
	if !_ecdgd {
		return false
	}
	if _cffdc.InfoDict == nil {
		return false
	}
	_bcbca, _egfbb := _db.NewPdfInfoFromObject(_cffdc.InfoDict)
	if _egfbb != nil {
		return false
	}
	if _gffgc.Creator != nil {
		if _bcbca.Creator == nil || _bcbca.Creator.String() != _gffgc.Creator.String() {
			return false
		}
	}
	if _gffgc.CreationDate != nil {
		if _bcbca.CreationDate == nil || !_bcbca.CreationDate.ToGoTime().Equal(_gffgc.CreationDate.ToGoTime()) {
			return false
		}
	}
	if _gffgc.ModifiedDate != nil {
		if _bcbca.ModifiedDate == nil || !_bcbca.ModifiedDate.ToGoTime().Equal(_gffgc.ModifiedDate.ToGoTime()) {
			return false
		}
	}
	if _gffgc.Producer != nil {
		if _bcbca.Producer == nil || _bcbca.Producer.String() != _gffgc.Producer.String() {
			return false
		}
	}
	if _gffgc.Keywords != nil {
		if _bcbca.Keywords == nil || _bcbca.Keywords.String() != _gffgc.Keywords.String() {
			return false
		}
	}
	if _gffgc.Trapped != nil {
		if _bcbca.Trapped == nil {
			return false
		}
		switch _gffgc.Trapped.String() {
		case "\u0054\u0072\u0075\u0065":
			if _bcbca.Trapped.String() != "\u0054\u0072\u0075\u0065" {
				return false
			}
		case "\u0046\u0061\u006cs\u0065":
			if _bcbca.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		default:
			if _bcbca.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		}
	}
	if _gffgc.Title != nil {
		if _bcbca.Title == nil || _bcbca.Title.String() != _gffgc.Title.String() {
			return false
		}
	}
	if _gffgc.Subject != nil {
		if _bcbca.Subject == nil || _bcbca.Subject.String() != _gffgc.Subject.String() {
			return false
		}
	}
	return true
}

func _ffgd(_ceadf *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_bgecc = "\u0036.\u0033\u002e\u0033\u002d\u0032"
		_cgg   = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0054y\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020f\u006f\u0072 \u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067,\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0020\u0064\u0069c\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006f\u0047\u0049D\u004d\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070\u0070\u0069\u006e\u0067\u0020\u0066\u0072\u006f\u006d\u0020\u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061s d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020P\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0054a\u0062\u006c\u0065\u0020\u0035\u002e\u00313"
	)
	var _faacf string
	if _dcdc, _gace := _cb.GetName(_ceadf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _gace {
		_faacf = _dcdc.String()
	}
	if _faacf != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _ce
	}
	if _ceadf.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _fdbe(_bgecc, _cgg)
	}
	return _ce
}

func _eddf(_ebda *_db.CompliancePdfReader) ViolatedRule {
	for _, _bdea := range _ebda.PageList {
		_bcbe, _deced := _bdea.GetContentStreams()
		if _deced != nil {
			continue
		}
		for _, _dacdf := range _bcbe {
			_aefc := _df.NewContentStreamParser(_dacdf)
			_, _deced = _aefc.Parse()
			if _deced != nil {
				return _fdbe("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _ce
}

func _aee(_fafd *_db.CompliancePdfReader) (_cfgda []ViolatedRule) {
	var (
		_dbgbf, _bbfb, _gdafb, _eedgg, _abbe bool
		_befda                               func(_cb.PdfObject)
	)
	_befda = func(_cba _cb.PdfObject) {
		switch _gcgc := _cba.(type) {
		case *_cb.PdfObjectInteger:
			if !_dbgbf && (int64(*_gcgc) > _bf.MaxInt32 || int64(*_gcgc) < -_bf.MaxInt32) {
				_cfgda = append(_cfgda, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_dbgbf = true
			}
		case *_cb.PdfObjectFloat:
			if !_bbfb && (_bf.Abs(float64(*_gcgc)) > _bf.MaxFloat32) {
				_cfgda = append(_cfgda, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0032", "\u0041 \u0063\u006f\u006e\u0066orm\u0069\u006e\u0067\u0020f\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0061\u006c\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0075\u0074\u0073\u0069de\u0020\u0074\u0068e\u0020\u0072\u0061\u006e\u0067e\u0020o\u0066\u0020\u002b\u002f\u002d\u0033\u002e\u0034\u00303\u0020\u0078\u0020\u0031\u0030\u005e\u0033\u0038\u002e"))
			}
		case *_cb.PdfObjectString:
			if !_gdafb && len([]byte(_gcgc.Str())) > 32767 {
				_cfgda = append(_cfgda, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0033\u0032\u0037\u0036\u0037\u002e"))
				_gdafb = true
			}
		case *_cb.PdfObjectName:
			if !_eedgg && len([]byte(*_gcgc)) > 127 {
				_cfgda = append(_cfgda, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_eedgg = true
			}
		case *_cb.PdfObjectArray:
			for _, _dcbef := range _gcgc.Elements() {
				_befda(_dcbef)
			}
			if !_abbe && (_gcgc.Len() == 4 || _gcgc.Len() == 5) {
				_gdeb, _ebgc := _cb.GetName(_gcgc.Get(0))
				if !_ebgc {
					return
				}
				if *_gdeb != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_cccf := _gcgc.Get(1)
				_cccf = _cb.TraceToDirectObject(_cccf)
				_bcbg, _ebgc := _cb.GetArray(_cccf)
				if !_ebgc {
					return
				}
				if _bcbg.Len() > 32 {
					_cfgda = append(_cfgda, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0039", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006e\u0075\u006db\u0065\u0072\u0020\u006ff\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0069\u0073\u0020\u0033\u0032\u002e"))
					_abbe = true
				}
			}
		case *_cb.PdfObjectDictionary:
			_gbaeg := _gcgc.Keys()
			for _fddb, _ceca := range _gbaeg {
				_befda(&_gbaeg[_fddb])
				_befda(_gcgc.Get(_ceca))
			}
		case *_cb.PdfObjectStream:
			_befda(_gcgc.PdfObjectDictionary)
		case *_cb.PdfObjectStreams:
			for _, _edfb := range _gcgc.Elements() {
				_befda(_edfb)
			}
		case *_cb.PdfObjectReference:
			_befda(_gcgc.Resolve())
		}
	}
	_bfcg := _fafd.GetObjectNums()
	if len(_bfcg) > 8388607 {
		_cfgda = append(_cfgda, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _dcfgd := range _bfcg {
		_agddc, _cgfa := _fafd.GetIndirectObjectByNumber(_dcfgd)
		if _cgfa != nil {
			continue
		}
		_agdgb := _cb.TraceToDirectObject(_agddc)
		_befda(_agdgb)
	}
	return _cfgda
}

func _gebb(_cecb *_db.CompliancePdfReader) (_bgda []ViolatedRule) {
	_fbde := _cecb.ParserMetadata()
	if _fbde.HasInvalidSubsectionHeader() {
		_bgda = append(_bgda, _fdbe("\u0036.\u0031\u002e\u0034\u002d\u0031", "\u006e\u0020\u0061\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0073\u0065c\u0074\u0069\u006f\u006e\u0020h\u0065a\u0064\u0065\u0072\u0020t\u0068\u0065\u0020\u0073\u0074\u0061\u0072t\u0069\u006e\u0067\u0020\u006fb\u006a\u0065\u0063\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0072\u0061n\u0067e\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0053\u0050\u0041C\u0045\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u0028\u0032\u0030\u0068\u0029\u002e"))
	}
	if _fbde.HasInvalidSeparationAfterXRef() {
		_bgda = append(_bgda, _fdbe("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e"))
	}
	return _bgda
}

func _fcbc(_gcce *_db.CompliancePdfReader) ViolatedRule {
	for _, _bbec := range _gcce.PageList {
		_ebaec, _ggab := _bbec.GetContentStreams()
		if _ggab != nil {
			continue
		}
		for _, _ffdgd := range _ebaec {
			_abag := _df.NewContentStreamParser(_ffdgd)
			_, _ggab = _abag.Parse()
			if _ggab != nil {
				return _fdbe("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _ce
}

func _ebed(_bbeg standardType, _dddd *_f.OutputIntents) error {
	_bafd, _dbec := _ebb.NewCmykIsoCoatedV2OutputIntent(_bbeg.outputIntentSubtype())
	if _dbec != nil {
		return _dbec
	}
	if _dbec = _dddd.Add(_bafd.ToPdfObject()); _dbec != nil {
		return _dbec
	}
	return nil
}

var _ce = ViolatedRule{}

func _ccedd(_aacf *_db.CompliancePdfReader) (_aaecb []ViolatedRule) {
	if _aacf.ParserMetadata().HasOddLengthHexStrings() {
		_aaecb = append(_aaecb, _fdbe("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _aacf.ParserMetadata().HasOddLengthHexStrings() {
		_aaecb = append(_aaecb, _fdbe("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _aaecb
}

func _befg(_eacga *_db.CompliancePdfReader) ViolatedRule {
	if _eacga.ParserMetadata().HasDataAfterEOF() {
		return _fdbe("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _ce
}

func _cgc(_ffea *Profile1Options) {
	if _ffea.Now == nil {
		_ffea.Now = _c.Now
	}
}

func _efgc(_dbgbe *_db.CompliancePdfReader) (_gfaf []ViolatedRule) {
	_fgad := true
	_bceda, _egdeb := _dbgbe.GetCatalogMarkInfo()
	if !_egdeb {
		_fgad = false
	} else {
		_ebddd, _gdcc := _cb.GetDict(_bceda)
		if _gdcc {
			_fcgc, _ffcg := _cb.GetBool(_ebddd.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_fcgc) || !_ffcg {
				_fgad = false
			}
		} else {
			_fgad = false
		}
	}
	if !_fgad {
		_gfaf = append(_gfaf, _fdbe("\u0036.\u0038\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_efef, _egdeb := _dbgbe.GetCatalogStructTreeRoot()
	if !_egdeb {
		_gfaf = append(_gfaf, _fdbe("\u0036.\u0038\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_fbfa, _egdeb := _cb.GetDict(_efef)
	if _egdeb {
		_fgef, _fgfb := _cb.GetName(_fbfa.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _fgfb {
			_ebeb, _cfae := _cb.GetDict(_fgef)
			if _cfae {
				for _, _bada := range _ebeb.Keys() {
					_dcabc := _ebeb.Get(_bada)
					if _dcabc == nil {
						_gfaf = append(_gfaf, _fdbe("\u0036.\u0038\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _gfaf
}
func _fdc() standardType { return standardType{_ed: 1, _fd: "\u0041"} }
func _ddfb(_eage *_db.CompliancePdfReader) (_fefdd []ViolatedRule) {
	var _aecc, _eebcb, _dcg, _geff, _aacg, _gfcg, _ggbf bool
	_dgde := func() bool { return _aecc && _eebcb && _dcg && _geff && _aacg && _gfcg && _ggbf }
	for _, _fbge := range _eage.PageList {
		if _fbge.Resources == nil {
			continue
		}
		_dgeag, _bgdad := _cb.GetDict(_fbge.Resources.Font)
		if !_bgdad {
			continue
		}
		for _, _aeada := range _dgeag.Keys() {
			_cege, _abge := _cb.GetDict(_dgeag.Get(_aeada))
			if !_abge {
				if !_aecc {
					_fefdd = append(_fefdd, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006fn\u0074\u0073\u0020\u0075\u0073e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0066\u006f\u0072m\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0073\u0020d\u0065\u0066\u0069\u006e\u0065d \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035\u002e\u0035\u002e"))
					_aecc = true
					if _dgde() {
						return _fefdd
					}
				}
				continue
			}
			if _ddbb, _aegdg := _cb.GetName(_cege.Get("\u0054\u0079\u0070\u0065")); !_aecc && (!_aegdg || _ddbb.String() != "\u0046\u006f\u006e\u0074") {
				_fefdd = append(_fefdd, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_aecc = true
				if _dgde() {
					return _fefdd
				}
			}
			_fdbbg, _fafb := _db.NewPdfFontFromPdfObject(_cege)
			if _fafb != nil {
				continue
			}
			var _ffaa string
			if _decb, _eagg := _cb.GetName(_cege.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _eagg {
				_ffaa = _decb.String()
			}
			if !_eebcb {
				switch _ffaa {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_eebcb = true
					_fefdd = append(_fefdd, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _dgde() {
						return _fefdd
					}
				}
			}
			if !_dcg {
				if _ffaa != "\u0054\u0079\u0070e\u0033" {
					_edbb, _gfge := _cb.GetName(_cege.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_gfge || _edbb.String() == "" {
						_fefdd = append(_fefdd, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_dcg = true
						if _dgde() {
							return _fefdd
						}
					}
				}
			}
			if _ffaa != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_caac := _acd.IsStdFont(_acd.StdFontName(_fdbbg.BaseFont()))
			if _caac {
				continue
			}
			_cdgb, _eddg := _cb.GetIntVal(_cege.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_eddg && !_geff {
				_fefdd = append(_fefdd, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_geff = true
				if _dgde() {
					return _fefdd
				}
			}
			_ggcg, _cfcg := _cb.GetIntVal(_cege.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_cfcg && !_aacg {
				_fefdd = append(_fefdd, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_aacg = true
				if _dgde() {
					return _fefdd
				}
			}
			if !_gfcg {
				_fbbe, _cdgbe := _cb.GetArray(_cege.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_cdgbe || !_eddg || !_cfcg || _fbbe.Len() != _ggcg-_cdgb+1 {
					_fefdd = append(_fefdd, _fdbe("\u0036.\u0033\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_gfcg = true
					if _dgde() {
						return _fefdd
					}
				}
			}
		}
	}
	return _fefdd
}

func _deabd(_dfeb *_db.CompliancePdfReader) (_cdfe []ViolatedRule) {
	var _cegcf, _fdbfa bool
	_ceffg := func() bool { return _cegcf && _fdbfa }
	for _, _cbab := range _dfeb.GetObjectNums() {
		_gcad, _daddg := _dfeb.GetIndirectObjectByNumber(_cbab)
		if _daddg != nil {
			_eg.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _cbab, _daddg)
			continue
		}
		_abeff, _faee := _cb.GetDict(_gcad)
		if !_faee {
			continue
		}
		_fbga, _faee := _cb.GetName(_abeff.Get("\u0054\u0079\u0070\u0065"))
		if !_faee {
			continue
		}
		if *_fbga != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_babba, _faee := _cb.GetName(_abeff.Get("\u0053"))
		if !_faee {
			if !_cegcf {
				_cdfe = append(_cdfe, _fdbe("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_cegcf = true
				if _ceffg() {
					return _cdfe
				}
			}
			continue
		}
		switch _db.PdfActionType(*_babba) {
		case _db.ActionTypeLaunch, _db.ActionTypeSound, _db.ActionTypeMovie, _db.ActionTypeResetForm, _db.ActionTypeImportData, _db.ActionTypeJavaScript, _db.ActionTypeHide, _db.ActionTypeSetOCGState, _db.ActionTypeRendition, _db.ActionTypeTrans, _db.ActionTypeGoTo3DView:
			if !_cegcf {
				_cdfe = append(_cdfe, _fdbe("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_cegcf = true
				if _ceffg() {
					return _cdfe
				}
			}
			continue
		case _db.ActionTypeNamed:
			if !_fdbfa {
				_fbafa, _feecd := _cb.GetName(_abeff.Get("\u004e"))
				if !_feecd {
					_cdfe = append(_cdfe, _fdbe("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fdbfa = true
					if _ceffg() {
						return _cdfe
					}
					continue
				}
				switch *_fbafa {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_cdfe = append(_cdfe, _fdbe("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fdbfa = true
					if _ceffg() {
						return _cdfe
					}
					continue
				}
			}
		}
	}
	return _cdfe
}

// Part gets the PDF/A version level.
func (_abga *profile2) Part() int { return _abga._dgfgd._ed }

// DefaultProfile2Options are the default options for the Profile2.
func DefaultProfile2Options() *Profile2Options {
	return &Profile2Options{Now: _c.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}

func _bdce(_cfaf *Profile2Options) {
	if _cfaf.Now == nil {
		_cfaf.Now = _c.Now
	}
}

// Profile1Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile1Options struct {
	// CMYKDefaultColorSpace is an option that refers PDF/A-1
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _c.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func _aede(_ecdf *_db.PdfFont, _ggba *_cb.PdfObjectDictionary, _fegad bool) ViolatedRule {
	const (
		_dfbf = "\u0036.\u0033\u002e\u0034\u002d\u0031"
		_gdc  = "\u0054\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u0020s\u0068\u0061\u006cl\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006ce\u002c\u0020a\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0035\u002e\u0038\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0077h\u0065\u006e\u0020\u0074\u0068\u0065 \u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0065\u0078\u0063\u006cu\u0073i\u0076\u0065\u006c\u0079\u0020\u0077\u0069t\u0068\u0020\u0074\u0065\u0078\u0074\u0020\u0072e\u006ed\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020\u0033\u002e"
	)
	if _fegad {
		return _ce
	}
	_edcb := _ecdf.FontDescriptor()
	var _gfcga string
	if _cada, _abab := _cb.GetName(_ggba.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _abab {
		_gfcga = _cada.String()
	}
	switch _gfcga {
	case "\u0054\u0079\u0070e\u0031":
		if _edcb.FontFile == nil {
			return _fdbe(_dfbf, _gdc)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _edcb.FontFile2 == nil {
			return _fdbe(_dfbf, _gdc)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _edcb.FontFile3 == nil {
			return _fdbe(_dfbf, _gdc)
		}
	}
	return _ce
}

func _ceeb(_fgfeb *_db.PdfFont, _bfbb *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_fdbdg = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0036\u002d\u0033"
		_deff  = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _dead string
	if _bddf, _ccdf := _cb.GetName(_bfbb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ccdf {
		_dead = _bddf.String()
	}
	if _dead != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _ce
	}
	_gbbd := _fgfeb.FontDescriptor()
	_feag, _ceeeb := _cb.GetIntVal(_gbbd.Flags)
	if !_ceeeb {
		_eg.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _fdbe(_fdbdg, _deff)
	}
	_dcdd := (uint32(_feag) >> 3) & 1
	_fdfc := _dcdd != 0
	if !_fdfc {
		return _ce
	}
	if _bfbb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _fdbe(_fdbdg, _deff)
	}
	return _ce
}

func _dag(_fcc *_f.Document) error {
	for _, _ebd := range _fcc.Objects {
		_gbaf, _ddbc := _cb.GetDict(_ebd)
		if !_ddbc {
			continue
		}
		_dcc := _gbaf.Get("\u0054\u0079\u0070\u0065")
		if _dcc == nil {
			continue
		}
		if _gec, _aef := _cb.GetName(_dcc); _aef && _gec.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_cddg, _dabc := _cb.GetBool(_gbaf.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _dabc {
			if bool(*_cddg) {
				_gbaf.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cb.MakeBool(false))
			}
		}
		_fcbd := _gbaf.Get("\u0041")
		if _fcbd != nil {
			_gbaf.Remove("\u0041")
		}
		_dbbf, _dabc := _cb.GetArray(_gbaf.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if _dabc {
			for _geef := 0; _geef < _dbbf.Len(); _geef++ {
				_bce, _dbgb := _cb.GetDict(_dbbf.Get(_geef))
				if !_dbgb {
					continue
				}
				if _bce.Get("\u0041\u0041") != nil {
					_bce.Remove("\u0041\u0041")
				}
			}
		}
	}
	return nil
}

func _gega(_febg *_db.CompliancePdfReader) (_eccf []ViolatedRule) {
	_bgeg, _dfbaf := _eagdc(_febg)
	if !_dfbaf {
		return _eccf
	}
	_aeaeg, _dfbaf := _cb.GetDict(_bgeg.Get("\u004e\u0061\u006de\u0073"))
	if !_dfbaf {
		return _eccf
	}
	if _aeaeg.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_eccf = append(_eccf, _fdbe("\u0036\u002e\u0031\u0030\u002d\u0031", "T\u0068\u0065\u0072e\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u006e\u006f\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0050\u0072\u0065s\u0065\u006e\u0074a\u0074\u0069\u006f\u006e\u0073\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002e"))
	}
	return _eccf
}

type colorspaceModification struct {
	_daa _dbg.ColorConverter
	_ef  _db.PdfColorspace
}

func _ec(_bfge *_db.XObjectImage, _bac imageModifications) error {
	_gb, _eca := _bfge.ToImage()
	if _eca != nil {
		return _eca
	}
	if _bac._aed != nil {
		_bfge.Filter = _bac._aed
	}
	_bfgc := _cb.MakeDict()
	_bfgc.Set("\u0051u\u0061\u006c\u0069\u0074\u0079", _cb.MakeInteger(100))
	_bfgc.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", _cb.MakeInteger(1))
	_bfge.Decode = nil
	if _eca = _bfge.SetImage(_gb, nil); _eca != nil {
		return _eca
	}
	_bfge.ToPdfObject()
	return nil
}

func _gfeg(_affd *_db.CompliancePdfReader, _acgcaf standardType, _bbdcc bool) (_bdcbd []ViolatedRule) {
	_eegb, _aebb := _eagdc(_affd)
	if !_aebb {
		return []ViolatedRule{_fdbe("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_dcadf := _eegb.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _dcadf == nil {
		return []ViolatedRule{_fdbe("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_gdfd, _aebb := _cb.GetStream(_dcadf)
	if !_aebb {
		return []ViolatedRule{_fdbe("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_gdegd, _efce := _fcg.LoadDocument(_gdfd.Stream)
	if _efce != nil {
		return []ViolatedRule{_fdbe("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d4", "\u0041\u006c\u006c\u0020\u006de\u0074\u0061\u0064a\u0074\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020i\u006e \u0074\u0068\u0065\u0020\u0050\u0044\u0046 \u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065ci\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002e\u0020\u0041\u006c\u006c\u0020c\u006fn\u0074\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0061\u006c\u006c\u0020\u0058\u004d\u0050\u0020p\u0061\u0063\u006b\u0065\u0074\u0073 \u0073h\u0061\u006c\u006c \u0062\u0065\u0020\u0077\u0065\u006c\u006c\u002d\u0066o\u0072\u006de\u0064")}
	}
	_eadd := _gdegd.GetGoXmpDocument()
	var _abbgf []*_eb.Namespace
	for _, _fcfgc := range _eadd.Namespaces() {
		switch _fcfgc.Name {
		case _da.NsDc.Name, _bfg.NsPDF.Name, _gf.NsXmp.Name, _bc.NsXmpRights.Name, _fc.Namespace.Name, _eee.Namespace.Name, _ac.NsXmpMM.Name, _eee.FieldNS.Name, _eee.SchemaNS.Name, _eee.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_abbgf = append(_abbgf, _fcfgc)
	}
	_cggd := true
	_aeae, _efce := _gdegd.GetPdfaExtensionSchemas()
	if _efce == nil {
		for _, _ddbee := range _abbgf {
			var _geba bool
			for _bcgdc := range _aeae {
				if _ddbee.URI == _aeae[_bcgdc].NamespaceURI {
					_geba = true
					break
				}
			}
			if !_geba {
				_cggd = false
				break
			}
		}
	} else {
		_cggd = false
	}
	if !_cggd {
		_bdcbd = append(_bdcbd, _fdbe("\u0036.\u0036\u002e\u0032\u002e\u0033\u002d7", "\u0041\u006c\u006c\u0020\u0070\u0072\u006f\u0070e\u0072\u0074\u0069e\u0073\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072m\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0075s\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063he\u006da\u0073 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002c\u0020\u0049\u0053\u004f\u0020\u0031\u00390\u0030\u0035-\u0031\u0020\u006f\u0072\u0020\u0074h\u0069s\u0020\u0070\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u0020\u0031\u0039\u0030\u0030\u0035\u002c\u0020o\u0072\u0020\u0061\u006e\u0079\u0020e\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073c\u0068\u0065\u006das\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006fm\u0070\u006c\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0036\u002e\u0036\u002e\u0032.\u0033\u002e\u0032\u002e"))
	}
	_gaedd, _aebb := _gdegd.GetPdfAID()
	if !_aebb {
		_bdcbd = append(_bdcbd, _fdbe("\u0036.\u0036\u002e\u0034\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _gaedd.Part != _acgcaf._ed {
			_bdcbd = append(_bdcbd, _fdbe("\u0036.\u0036\u002e\u0034\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _acgcaf._fd == "\u0041" && _gaedd.Conformance != "\u0041" {
			_bdcbd = append(_bdcbd, _fdbe("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _acgcaf._fd == "\u0055" && (_gaedd.Conformance != "\u0041" && _gaedd.Conformance != "\u0055") {
			_bdcbd = append(_bdcbd, _fdbe("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _acgcaf._fd == "\u0042" && (_gaedd.Conformance != "\u0041" && _gaedd.Conformance != "\u0042" && _gaedd.Conformance != "\u0055") {
			_bdcbd = append(_bdcbd, _fdbe("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		}
	}
	return _bdcbd
}

func _gbc(_gfd *_f.Document) error {
	_ded, _cag := _gfd.FindCatalog()
	if !_cag {
		return _ea.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_, _cag = _cb.GetDict(_ded.Object.Get("\u0041\u0041"))
	if !_cag {
		return nil
	}
	_ded.Object.Remove("\u0041\u0041")
	return nil
}

func _dbeg(_gdb *_f.Document, _geeg bool) error {
	_cbfd, _aab := _gdb.GetPages()
	if !_aab {
		return nil
	}
	for _, _gfef := range _cbfd {
		_dedc, _faa := _gfef.GetContents()
		if !_faa {
			continue
		}
		var _cgaed *_db.PdfPageResources
		_fefg, _faa := _gfef.GetResources()
		if _faa {
			_cgaed, _ = _db.NewPdfPageResourcesFromDict(_fefg)
		}
		for _gef, _cff := range _dedc {
			_bed, _adcg := _cff.GetData()
			if _adcg != nil {
				continue
			}
			_gfae := _df.NewContentStreamParser(string(_bed))
			_cef, _adcg := _gfae.Parse()
			if _adcg != nil {
				continue
			}
			_baee, _adcg := _fde(_cgaed, _cef, _geeg)
			if _adcg != nil {
				return _adcg
			}
			if _baee == nil {
				continue
			}
			if _adcg = (&_dedc[_gef]).SetData(_baee); _adcg != nil {
				return _adcg
			}
		}
	}
	return nil
}

type standardType struct {
	_ed int
	_fd string
}

func _bgdg(_afdae *_f.Document) error {
	_ccc, _fggg := _afdae.FindCatalog()
	if !_fggg {
		return _ea.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_gaaa, _fggg := _cb.GetDict(_ccc.Object.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_fggg {
		return nil
	}
	_cfce, _fggg := _cb.GetDict(_gaaa.Get("\u0044"))
	if _fggg {
		if _cfce.Get("\u0041\u0053") != nil {
			_cfce.Remove("\u0041\u0053")
		}
	}
	_gefg, _fggg := _cb.GetArray(_gaaa.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _fggg {
		for _ebac := 0; _ebac < _gefg.Len(); _ebac++ {
			_dad, _dacfb := _cb.GetDict(_gefg.Get(_ebac))
			if !_dacfb {
				continue
			}
			if _dad.Get("\u0041\u0053") != nil {
				_dad.Remove("\u0041\u0053")
			}
		}
	}
	return nil
}
func _bebg(_ffbb *_db.CompliancePdfReader) ViolatedRule { return _ce }
func _fed(_beb *_f.Document, _bfd standardType, _feee XmpOptions) error {
	_efbf, _cfea := _beb.FindCatalog()
	if !_cfea {
		return nil
	}
	var _cfg *_fcg.Document
	_ace, _cfea := _efbf.GetMetadata()
	if !_cfea {
		_cfg = _fcg.NewDocument()
	} else {
		var _bgga error
		_cfg, _bgga = _fcg.LoadDocument(_ace.Stream)
		if _bgga != nil {
			return _bgga
		}
	}
	_ece := _fcg.PdfInfoOptions{InfoDict: _beb.Info, PdfVersion: _b.Sprintf("\u0025\u0064\u002e%\u0064", _beb.Version.Major, _beb.Version.Minor), Copyright: _feee.Copyright, Overwrite: true}
	_bbc, _cfea := _efbf.GetMarkInfo()
	if _cfea {
		_daad, _bae := _cb.GetBool(_bbc.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
		if _bae && bool(*_daad) {
			_ece.Marked = true
		}
	}
	if _ddg := _cfg.SetPdfInfo(&_ece); _ddg != nil {
		return _ddg
	}
	if _ebe := _cfg.SetPdfAID(_bfd._ed, _bfd._fd); _ebe != nil {
		return _ebe
	}
	_aaf := _fcg.MediaManagementOptions{OriginalDocumentID: _feee.OriginalDocumentID, DocumentID: _feee.DocumentID, InstanceID: _feee.InstanceID, NewDocumentID: !_feee.NewDocumentVersion, ModifyComment: "O\u0070\u0074\u0069\u006d\u0069\u007ae\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u006f\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064"}
	_dde, _cfea := _cb.GetDict(_beb.Info)
	if _cfea {
		if _gceb, _cbc := _cb.GetString(_dde.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _cbc && _gceb.String() != "" {
			_dgg, _bebd := _egd.ParsePdfTime(_gceb.String())
			if _bebd != nil {
				return _b.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _bebd)
			}
			_aaf.ModifyDate = _dgg
		}
	}
	if _fcba := _cfg.SetMediaManagement(&_aaf); _fcba != nil {
		return _fcba
	}
	if _aebg := _cfg.SetPdfAExtension(); _aebg != nil {
		return _aebg
	}
	_gbe, _bga := _cfg.MarshalIndent(_feee.MarshalPrefix, _feee.MarshalIndent)
	if _bga != nil {
		return _bga
	}
	if _gbd := _efbf.SetMetadata(_gbe); _gbd != nil {
		return _gbd
	}
	return nil
}

func _cbea(_acdcc *_db.CompliancePdfReader) ViolatedRule {
	if _acdcc.ParserMetadata().HeaderPosition() != 0 {
		return _fdbe("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	return _ce
}

func _egagd(_bdged *_db.CompliancePdfReader) (_babd []ViolatedRule) {
	_fdee := func(_dbfe *_cb.PdfObjectDictionary, _agdgdf *[]string, _caeg *[]ViolatedRule) error {
		_cgdgc := _dbfe.Get("\u004e\u0061\u006d\u0065")
		if _cgdgc == nil || len(_cgdgc.String()) == 0 {
			*_caeg = append(*_caeg, _fdbe("\u0036\u002e\u0039-\u0031", "\u0045\u0061\u0063\u0068\u0020o\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0063\u006fn\u0066\u0069\u0067\u0075r\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002e"))
		}
		for _, _bfeg := range *_agdgdf {
			if _bfeg == _cgdgc.String() {
				*_caeg = append(*_caeg, _fdbe("\u0036\u002e\u0039-\u0032", "\u0045\u0061\u0063\u0068\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0066\u0069\u0067\u0075\u0072a\u0074\u0069\u006fn\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020N\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002c w\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075ni\u0071\u0075\u0065 \u0061\u006d\u006f\u006e\u0067\u0073\u0074\u0020\u0061\u006c\u006c\u0020o\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006fn\u0074\u0065\u006e\u0074 \u0063\u006f\u006e\u0066\u0069\u0067u\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061\u0072\u0069\u0065\u0073\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0074\u0068e\u0020\u0050\u0044\u0046\u002fA\u002d\u0032\u0020\u0066\u0069l\u0065\u002e"))
			} else {
				*_agdgdf = append(*_agdgdf, _cgdgc.String())
			}
		}
		if _dbfe.Get("\u0041\u0053") != nil {
			*_caeg = append(*_caeg, _fdbe("\u0036\u002e\u0039-\u0034", "Th\u0065\u0020\u0041\u0053\u0020\u006b\u0065y \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0070\u0070\u0065\u0061r\u0020\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0063\u006fn\u0066\u0069\u0067\u0075\u0072\u0061\u0074\u0069\u006fn\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
		}
		return nil
	}
	_bcbce, _eaaf := _eagdc(_bdged)
	if !_eaaf {
		return _babd
	}
	_fagbb, _eaaf := _cb.GetDict(_bcbce.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_eaaf {
		return _babd
	}
	var _ddba []string
	_bccc, _eaaf := _cb.GetDict(_fagbb.Get("\u0044"))
	if _eaaf {
		_fdee(_bccc, &_ddba, &_babd)
	}
	_feaga, _eaaf := _cb.GetArray(_fagbb.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _eaaf {
		for _bcab := 0; _bcab < _feaga.Len(); _bcab++ {
			_afgag, _gcdd := _cb.GetDict(_feaga.Get(_bcab))
			if !_gcdd {
				continue
			}
			_fdee(_afgag, &_ddba, &_babd)
		}
	}
	return _babd
}

func _edcbf(_edcg *_db.PdfFont, _edcc *_cb.PdfObjectDictionary) ViolatedRule {
	const (
		_aggd = "\u0036.\u0033\u002e\u0035\u002d\u0033"
		_egfa = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0073 \u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077i\u0074\u0068\u0069n\u0020\u0061\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0064\u0065s\u0063\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020\u0043\u0049\u0044\u0053\u0065\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0079\u0069\u006eg\u0020\u0077\u0068i\u0063\u0068\u0020\u0043\u0049\u0044\u0073 \u0061\u0072e\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e \u0074\u0068\u0065\u0020\u0065\u006d\u0062\u0065\u0064d\u0065\u0064\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0066\u0069l\u0065,\u0020\u0061\u0073 \u0064\u0065\u0073\u0063\u0072\u0069b\u0065\u0064 \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020\u0054ab\u006c\u0065\u0020\u0035.\u00320\u002e"
	)
	var _dgee string
	if _agbe, _abcd := _cb.GetName(_edcc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _abcd {
		_dgee = _agbe.String()
	}
	switch _dgee {
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_ccec := _edcg.FontDescriptor()
		if _ccec.CIDSet == nil {
			return _fdbe(_aggd, _egfa)
		}
		return _ce
	default:
		return _ce
	}
}

func _aege(_aega *_f.Document) error {
	_gdeg, _fcea := _aega.FindCatalog()
	if !_fcea {
		return _ea.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_bggc, _fcea := _cb.GetDict(_gdeg.Object.Get("\u004e\u0061\u006de\u0073"))
	if !_fcea {
		return nil
	}
	if _bggc.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_bggc.Remove("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073")
	}
	return nil
}

func _bdgc(_gdfc *_cb.PdfObjectDictionary, _dggc map[*_cb.PdfObjectStream][]byte, _gafee map[*_cb.PdfObjectStream]*_cd.CMap) ViolatedRule {
	const (
		_fcedb = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0037\u002d\u0031"
		_ecge  = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_gcegf, _ggdg := _cb.GetStream(_gdfc.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _ggdg {
		_, _beaec := _feac(_gcegf, _dggc, _gafee)
		if _beaec != nil {
			return _fdbe(_fcedb, _ecge)
		}
		return _ce
	}
	_fgdg, _ggdg := _cb.GetName(_gdfc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_ggdg {
		return _fdbe(_fcedb, _ecge)
	}
	switch _fgdg.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _ce
	}
	return _fdbe(_fcedb, _ecge)
}

func _cbcec(_cdbfa *_db.CompliancePdfReader) (_ebfe []ViolatedRule) {
	var (
		_dcce, _agea, _ggda, _eagdb, _dbega, _bafe, _dfee bool
		_egca                                             func(_cb.PdfObject)
	)
	_egca = func(_cdba _cb.PdfObject) {
		switch _faef := _cdba.(type) {
		case *_cb.PdfObjectInteger:
			if !_dcce && (int64(*_faef) > _bf.MaxInt32 || int64(*_faef) < -_bf.MaxInt32) {
				_ebfe = append(_ebfe, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_dcce = true
			}
		case *_cb.PdfObjectFloat:
			if !_agea && (_bf.Abs(float64(*_faef)) > 32767.0) {
				_ebfe = append(_ebfe, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0032", "\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0020\u0072\u0065\u0061\u006c\u0020\u0076\u0061\u006c\u0075\u0065\u0020m\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u006c\u0065s\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075a\u006c\u0020\u0074\u006f\u0020\u00332\u0037\u0036\u0037.\u0030\u002e"))
			}
		case *_cb.PdfObjectString:
			if !_ggda && len([]byte(_faef.Str())) > 65535 {
				_ebfe = append(_ebfe, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0036\u0035\u0035\u0033\u0035\u002e"))
				_ggda = true
			}
		case *_cb.PdfObjectName:
			if !_eagdb && len([]byte(*_faef)) > 127 {
				_ebfe = append(_ebfe, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_eagdb = true
			}
		case *_cb.PdfObjectArray:
			if !_dbega && _faef.Len() > 8191 {
				_ebfe = append(_ebfe, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0035", "\u004d\u0061\u0078\u0069\u006d\u0075m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020(\u0069\u006e\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0029\u0020\u0069s\u00208\u0031\u0039\u0031\u002e"))
				_dbega = true
			}
			for _, _dgd := range _faef.Elements() {
				_egca(_dgd)
			}
			if !_dfee && (_faef.Len() == 4 || _faef.Len() == 5) {
				_fafg, _gcgd := _cb.GetName(_faef.Get(0))
				if !_gcgd {
					return
				}
				if *_fafg != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_dafa := _faef.Get(1)
				_dafa = _cb.TraceToDirectObject(_dafa)
				_ffdd, _gcgd := _cb.GetArray(_dafa)
				if !_gcgd {
					return
				}
				if _ffdd.Len() > 8 {
					_ebfe = append(_ebfe, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0039", "\u004d\u0061\u0078i\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0069\u0073\u0020\u0038\u002e"))
					_dfee = true
				}
			}
		case *_cb.PdfObjectDictionary:
			_bgfd := _faef.Keys()
			if !_bafe && len(_bgfd) > 4095 {
				_ebfe = append(_ebfe, _fdbe("\u0036.\u0031\u002e\u0031\u0032\u002d\u00311", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u0063\u0061\u0070\u0061\u0063\u0069\u0074y\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0028\u0069\u006e\u0020\u0065\u006e\u0074\u0072\u0069es\u0029\u0020\u0069\u0073\u0020\u0034\u0030\u0039\u0035\u002e"))
				_bafe = true
			}
			for _aadba, _bccb := range _bgfd {
				_egca(&_bgfd[_aadba])
				_egca(_faef.Get(_bccb))
			}
		case *_cb.PdfObjectStream:
			_egca(_faef.PdfObjectDictionary)
		case *_cb.PdfObjectStreams:
			for _, _agfg := range _faef.Elements() {
				_egca(_agfg)
			}
		case *_cb.PdfObjectReference:
			_egca(_faef.Resolve())
		}
	}
	_ddaf := _cdbfa.GetObjectNums()
	if len(_ddaf) > 8388607 {
		_ebfe = append(_ebfe, _fdbe("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _eecc := range _ddaf {
		_bcga, _dage := _cdbfa.GetIndirectObjectByNumber(_eecc)
		if _dage != nil {
			continue
		}
		_cbeg := _cb.TraceToDirectObject(_bcga)
		_egca(_cbeg)
	}
	return _ebfe
}

func _feac(_cacd *_cb.PdfObjectStream, _fbfac map[*_cb.PdfObjectStream][]byte, _abbgg map[*_cb.PdfObjectStream]*_cd.CMap) (*_cd.CMap, error) {
	_fbef, _bcgag := _abbgg[_cacd]
	if !_bcgag {
		var _egbd error
		_badg, _cafff := _fbfac[_cacd]
		if !_cafff {
			_badg, _egbd = _cb.DecodeStream(_cacd)
			if _egbd != nil {
				_eg.Log.Debug("\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _egbd)
				return nil, _egbd
			}
			_fbfac[_cacd] = _badg
		}
		_fbef, _egbd = _cd.LoadCmapFromData(_badg, false)
		if _egbd != nil {
			return nil, _egbd
		}
		_abbgg[_cacd] = _fbef
	}
	return _fbef, nil
}

// Profile2U is the implementation of the PDF/A-2U standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2U struct{ profile2 }

// Validate checks if provided input document reader matches given PDF/A profile.
func Validate(d *_db.CompliancePdfReader, profile Profile) error { return profile.ValidateStandard(d) }

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
type imageModifications struct {
	_eag *colorspaceModification
	_aed _cb.StreamEncoder
}
