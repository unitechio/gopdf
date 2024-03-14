package docutil

import (
	_f "errors"
	_d "fmt"

	_db "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/core"
)

func (_fac *Catalog) SetMarkInfo(mi _b.PdfObject) {
	_ba := _b.MakeIndirectObject(mi)
	_fac.Object.Set("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f", _ba)
	_fac._a.Objects = append(_fac._a.Objects, _ba)
}

type Document struct {
	ID             [2]string
	Version        _b.Version
	Objects        []_b.PdfObject
	Info           _b.PdfObject
	Crypt          *_b.PdfCrypt
	UseHashBasedID bool
}

func (_dbc *Catalog) GetMetadata() (*_b.PdfObjectStream, bool) {
	_ac, _cd := _b.GetStream(_dbc.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	return _ac, _cd
}

func (_ebc Content) GetData() ([]byte, error) {
	_gecb, _fcf := _b.NewEncoderFromStream(_ebc.Stream)
	if _fcf != nil {
		return nil, _fcf
	}
	_fcdc, _fcf := _gecb.DecodeStream(_ebc.Stream)
	if _fcf != nil {
		return nil, _fcf
	}
	return _fcdc, nil
}

func (_bdc *Content) SetData(data []byte) error {
	_efe, _afdc := _b.MakeStream(data, _b.NewFlateEncoder())
	if _afdc != nil {
		return _afdc
	}
	_gag, _dfd := _b.GetArray(_bdc._gec.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_dfd && _bdc._def == 0 {
		_bdc._gec.Object.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _efe)
	} else {
		if _afdc = _gag.Set(_bdc._def, _efe); _afdc != nil {
			return _afdc
		}
	}
	_bdc._gec._ecg.Objects = append(_bdc._gec._ecg.Objects, _efe)
	return nil
}

func (_cf *Catalog) GetOutputIntents() (*OutputIntents, bool) {
	_ef := _cf.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073")
	if _ef == nil {
		return nil, false
	}
	_ade, _adg := _b.GetIndirect(_ef)
	if !_adg {
		return nil, false
	}
	_ec, _dae := _b.GetArray(_ade.PdfObject)
	if !_dae {
		return nil, false
	}
	return &OutputIntents{_dg: _ade, _fc: _ec, _fag: _cf._a}, true
}

func (_be *Catalog) GetPages() ([]Page, bool) {
	_c, _de := _b.GetDict(_be.Object.Get("\u0050\u0061\u0067e\u0073"))
	if !_de {
		return nil, false
	}
	_e, _ea := _b.GetArray(_c.Get("\u004b\u0069\u0064\u0073"))
	if !_ea {
		return nil, false
	}
	_fe := make([]Page, _e.Len())
	for _da, _ad := range _e.Elements() {
		_ce, _bg := _b.GetDict(_ad)
		if !_bg {
			continue
		}
		_fe[_da] = Page{Object: _ce, _afd: _da + 1, _ecg: _be._a}
	}
	return _fe, true
}
func (_bed *OutputIntents) Len() int { return _bed._fc.Len() }

type Content struct {
	Stream *_b.PdfObjectStream
	_def   int
	_gec   Page
}

func (_gdc *Catalog) SetOutputIntents(outputIntents *OutputIntents) {
	if _cc := _gdc.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"); _cc != nil {
		for _bd, _gg := range _gdc._a.Objects {
			if _gg == _cc {
				if outputIntents._dg == _cc {
					return
				}
				_gdc._a.Objects = append(_gdc._a.Objects[:_bd], _gdc._a.Objects[_bd+1:]...)
				break
			}
		}
	}
	_bf := outputIntents._dg
	if _bf == nil {
		_bf = _b.MakeIndirectObject(outputIntents._fc)
	}
	_gdc.Object.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _bf)
	_gdc._a.Objects = append(_gdc._a.Objects, _bf)
}

func (_af *Catalog) SetVersion() {
	_af.Object.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _b.MakeName(_d.Sprintf("\u0025\u0064\u002e%\u0064", _af._a.Version.Major, _af._a.Version.Minor)))
}

func (_gd *Catalog) GetMarkInfo() (*_b.PdfObjectDictionary, bool) {
	_cdd, _fg := _b.GetDict(_gd.Object.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"))
	return _cdd, _fg
}

type ImageSMask struct {
	Image  *Image
	Stream *_b.PdfObjectStream
}

func (_ceg *OutputIntents) Get(i int) (OutputIntent, bool) {
	if _ceg._fc == nil {
		return OutputIntent{}, false
	}
	if i >= _ceg._fc.Len() {
		return OutputIntent{}, false
	}
	_cff := _ceg._fc.Get(i)
	_ecd, _bab := _b.GetIndirect(_cff)
	if !_bab {
		_dd, _dgc := _b.GetDict(_cff)
		return OutputIntent{Object: _dd}, _dgc
	}
	_gcb, _aa := _b.GetDict(_ecd.PdfObject)
	return OutputIntent{Object: _gcb}, _aa
}
func (_afg *Catalog) NewOutputIntents() *OutputIntents { return &OutputIntents{_fag: _afg._a} }

type Page struct {
	_afd   int
	Object *_b.PdfObjectDictionary
	_ecg   *Document
}

func (_bfc *OutputIntents) Add(oi _b.PdfObject) error {
	_dee, _eca := oi.(*_b.PdfObjectDictionary)
	if !_eca {
		return _f.New("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _gb, _gc := _b.GetStream(_dee.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _gc {
		_bfc._fag.Objects = append(_bfc._fag.Objects, _gb)
	}
	_cfd, _gf := oi.(*_b.PdfIndirectObject)
	if !_gf {
		_cfd = _b.MakeIndirectObject(oi)
	}
	if _bfc._fc == nil {
		_bfc._fc = _b.MakeArray(_cfd)
	} else {
		_bfc._fc.Append(_cfd)
	}
	_bfc._fag.Objects = append(_bfc._fag.Objects, _cfd)
	return nil
}

func _gcc(_cffb _b.PdfObject) (_b.PdfObjectName, error) {
	var _bgd *_b.PdfObjectName
	var _ge *_b.PdfObjectArray
	if _ee, _beda := _cffb.(*_b.PdfIndirectObject); _beda {
		if _ffb, _fd := _ee.PdfObject.(*_b.PdfObjectArray); _fd {
			_ge = _ffb
		} else if _aea, _dda := _ee.PdfObject.(*_b.PdfObjectName); _dda {
			_bgd = _aea
		}
	} else if _dgf, _agc := _cffb.(*_b.PdfObjectArray); _agc {
		_ge = _dgf
	} else if _ga, _cffg := _cffb.(*_b.PdfObjectName); _cffg {
		_bgd = _ga
	}
	if _bgd != nil {
		switch *_bgd {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_bgd, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_bgd, nil
		}
	}
	if _ge != nil && _ge.Len() > 0 {
		if _dfg, _ed := _ge.Get(0).(*_b.PdfObjectName); _ed {
			switch *_dfg {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _ge.Len() == 1 {
					return *_dfg, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_dfg, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_dfg, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_dfg, nil
			}
		}
	}
	return "", nil
}

func (_fgf Page) GetResources() (*_b.PdfObjectDictionary, bool) {
	return _b.GetDict(_fgf.Object.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
}

func (_fgfd Page) GetResourcesXObject() (*_b.PdfObjectDictionary, bool) {
	_geg, _fcd := _fgfd.GetResources()
	if !_fcd {
		return nil, false
	}
	return _b.GetDict(_geg.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
}

func (_gcbf Page) FindXObjectImages() ([]*Image, error) {
	_bc, _eacc := _gcbf.GetResourcesXObject()
	if !_eacc {
		return nil, nil
	}
	var _bag []*Image
	var _cga error
	_gee := map[*_b.PdfObjectStream]int{}
	_bca := map[*_b.PdfObjectStream]struct{}{}
	var _dfa int
	for _, _ged := range _bc.Keys() {
		_ead, _babd := _b.GetStream(_bc.Get(_ged))
		if !_babd {
			continue
		}
		if _, _fdc := _gee[_ead]; _fdc {
			continue
		}
		_bac, _ggb := _b.GetName(_ead.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_ggb || _bac.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_ffa := Image{BitsPerComponent: 8, Stream: _ead, Name: string(_ged)}
		if _ffa.Colorspace, _cga = _gcc(_ead.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _cga != nil {
			_db.Log.Error("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073", _cga)
			continue
		}
		if _dga, _gdb := _b.GetIntVal(_ead.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _gdb {
			_ffa.BitsPerComponent = _dga
		}
		if _bbf, _bad := _b.GetIntVal(_ead.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _bad {
			_ffa.Width = _bbf
		}
		if _ace, _ebe := _b.GetIntVal(_ead.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _ebe {
			_ffa.Height = _ace
		}
		if _cde, _edg := _b.GetStream(_ead.Get("\u0053\u004d\u0061s\u006b")); _edg {
			_ffa.SMask = &ImageSMask{Image: &_ffa, Stream: _cde}
			_bca[_cde] = struct{}{}
		}
		switch _ffa.Colorspace {
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_ffa.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_ffa.ColorComponents = 1
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_ffa.ColorComponents = 4
		default:
			_ffa.ColorComponents = -1
		}
		_gee[_ead] = _dfa
		_bag = append(_bag, &_ffa)
		_dfa++
	}
	var _cca []int
	for _, _acd := range _bag {
		if _acd.SMask != nil {
			_eg, _feg := _gee[_acd.SMask.Stream]
			if _feg {
				_cca = append(_cca, _eg)
			}
		}
	}
	_ccc := make([]*Image, len(_bag)-len(_cca))
	_dfa = 0
_dgg:
	for _bcf, _gfg := range _bag {
		for _, _ced := range _cca {
			if _bcf == _ced {
				continue _dgg
			}
		}
		_ccc[_dfa] = _gfg
		_dfa++
	}
	return _bag, nil
}

func (_bgf *Catalog) SetMetadata(data []byte) error {
	_eac, _fa := _b.MakeStream(data, nil)
	if _fa != nil {
		return _fa
	}
	_eac.Set("\u0054\u0079\u0070\u0065", _b.MakeName("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	_eac.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _b.MakeName("\u0058\u004d\u004c"))
	_bgf.Object.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _eac)
	_bgf._a.Objects = append(_bgf._a.Objects, _eac)
	return nil
}

func (_fca *Document) AddIndirectObject(indirect *_b.PdfIndirectObject) {
	for _, _ebf := range _fca.Objects {
		if _ebf == indirect {
			return
		}
	}
	_fca.Objects = append(_fca.Objects, indirect)
}

func (_ecb Page) FindXObjectForms() []*_b.PdfObjectStream {
	_bff, _gac := _ecb.GetResourcesXObject()
	if !_gac {
		return nil
	}
	_ffg := map[*_b.PdfObjectStream]struct{}{}
	var _dfe func(_ceda *_b.PdfObjectDictionary, _ege map[*_b.PdfObjectStream]struct{})
	_dfe = func(_cbc *_b.PdfObjectDictionary, _bfa map[*_b.PdfObjectStream]struct{}) {
		for _, _cef := range _cbc.Keys() {
			_ffbb, _eaccc := _b.GetStream(_cbc.Get(_cef))
			if !_eaccc {
				continue
			}
			if _, _gfe := _bfa[_ffbb]; _gfe {
				continue
			}
			_cea, _ddd := _b.GetName(_ffbb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
			if !_ddd || _cea.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_bfa[_ffbb] = struct{}{}
			_bga, _ddd := _b.GetDict(_ffbb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_ddd {
				continue
			}
			_gff, _edf := _b.GetDict(_bga.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
			if _edf {
				_dfe(_gff, _bfa)
			}
		}
	}
	_dfe(_bff, _ffg)
	var _cba []*_b.PdfObjectStream
	for _fed := range _ffg {
		_cba = append(_cba, _fed)
	}
	return _cba
}

func (_efg *Document) FindCatalog() (*Catalog, bool) {
	var _eb *_b.PdfObjectDictionary
	for _, _ab := range _efg.Objects {
		_ag, _bda := _b.GetDict(_ab)
		if !_bda {
			continue
		}
		if _deg, _bdf := _b.GetName(_ag.Get("\u0054\u0079\u0070\u0065")); _bdf && *_deg == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
			_eb = _ag
			break
		}
	}
	if _eb == nil {
		return nil, false
	}
	return &Catalog{Object: _eb, _a: _efg}, true
}

type OutputIntents struct {
	_fc  *_b.PdfObjectArray
	_fag *Document
	_dg  *_b.PdfIndirectObject
}

func (_babc Page) GetContents() ([]Content, bool) {
	_bee, _caf := _b.GetArray(_babc.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_caf {
		_dea, _cg := _b.GetStream(_babc.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if !_cg {
			return nil, false
		}
		return []Content{{Stream: _dea, _gec: _babc, _def: 0}}, true
	}
	_bdg := make([]Content, _bee.Len())
	for _cee, _bb := range _bee.Elements() {
		_dbf, _edd := _b.GetStream(_bb)
		if !_edd {
			continue
		}
		_bdg[_cee] = Content{Stream: _dbf, _gec: _babc, _def: _cee}
	}
	return _bdg, true
}

type OutputIntent struct{ Object *_b.PdfObjectDictionary }

func (_fb *Catalog) HasMetadata() bool {
	_ae := _fb.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _ae != nil
}
func (_dfc *Page) Number() int { return _dfc._afd }

type Image struct {
	Name             string
	Width            int
	Height           int
	Colorspace       _b.PdfObjectName
	ColorComponents  int
	BitsPerComponent int
	SMask            *ImageSMask
	Stream           *_b.PdfObjectStream
}
type Catalog struct {
	Object *_b.PdfObjectDictionary
	_a     *Document
}

func (_ca *Document) AddStream(stream *_b.PdfObjectStream) {
	for _, _df := range _ca.Objects {
		if _df == stream {
			return
		}
	}
	_ca.Objects = append(_ca.Objects, stream)
}

func (_aca *Document) GetPages() ([]Page, bool) {
	_cega, _cb := _aca.FindCatalog()
	if !_cb {
		return nil, false
	}
	return _cega.GetPages()
}
