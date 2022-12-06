package docutil

import (
	_b "errors"
	_f "fmt"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/core"
)

func (_fe *Catalog) GetMetadata() (*_d.PdfObjectStream, bool) {
	_ee, _eb := _d.GetStream(_fe.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	return _ee, _eb
}
func (_gda *Catalog) SetMetadata(data []byte) error {
	_gf, _ed := _d.MakeStream(data, nil)
	if _ed != nil {
		return _ed
	}
	_gf.Set("\u0054\u0079\u0070\u0065", _d.MakeName("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	_gf.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _d.MakeName("\u0058\u004d\u004c"))
	_gda.Object.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _gf)
	_gda._g.Objects = append(_gda._g.Objects, _gf)
	return nil
}
func (_gca *Document) FindCatalog() (*Catalog, bool) {
	var _abd *_d.PdfObjectDictionary
	for _, _eag := range _gca.Objects {
		_cf, _be := _d.GetDict(_eag)
		if !_be {
			continue
		}
		if _bd, _efc := _d.GetName(_cf.Get("\u0054\u0079\u0070\u0065")); _efc && *_bd == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
			_abd = _cf
			break
		}
	}
	if _abd == nil {
		return nil, false
	}
	return &Catalog{Object: _abd, _g: _gca}, true
}
func (_fcg Content) GetData() ([]byte, error) {
	_ff, _cad := _d.NewEncoderFromStream(_fcg.Stream)
	if _cad != nil {
		return nil, _cad
	}
	_cgfg, _cad := _ff.DecodeStream(_fcg.Stream)
	if _cad != nil {
		return nil, _cad
	}
	return _cgfg, nil
}
func (_gc *Catalog) GetMarkInfo() (*_d.PdfObjectDictionary, bool) {
	_da, _dd := _d.GetDict(_gc.Object.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"))
	return _da, _dd
}
func (_cfb Page) GetResources() (*_d.PdfObjectDictionary, bool) {
	return _d.GetDict(_cfb.Object.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
}

type ImageSMask struct {
	Image  *Image
	Stream *_d.PdfObjectStream
}

func (_e *Catalog) SetVersion() {
	_e.Object.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _d.MakeName(_f.Sprintf("\u0025\u0064\u002e%\u0064", _e._g.Version.Major, _e._g.Version.Minor)))
}
func (_cde *OutputIntents) Get(i int) (OutputIntent, bool) {
	if _cde._fb == nil {
		return OutputIntent{}, false
	}
	if i >= _cde._fb.Len() {
		return OutputIntent{}, false
	}
	_egf := _cde._fb.Get(i)
	_agf, _abf := _d.GetIndirect(_egf)
	if !_abf {
		_dg, _efb := _d.GetDict(_egf)
		return OutputIntent{Object: _dg}, _efb
	}
	_eff, _ccd := _d.GetDict(_agf.PdfObject)
	return OutputIntent{Object: _eff}, _ccd
}
func (_cd *Catalog) GetPages() ([]Page, bool) {
	_gd, _ec := _d.GetDict(_cd.Object.Get("\u0050\u0061\u0067e\u0073"))
	if !_ec {
		return nil, false
	}
	_cg, _ecb := _d.GetArray(_gd.Get("\u004b\u0069\u0064\u0073"))
	if !_ecb {
		return nil, false
	}
	_gdb := make([]Page, _cg.Len())
	for _cgb, _ae := range _cg.Elements() {
		_ca, _cc := _d.GetDict(_ae)
		if !_cc {
			continue
		}
		_gdb[_cgb] = Page{Object: _ca, _bag: _cgb + 1, _fgg: _cd._g}
	}
	return _gdb, true
}
func (_ab *Catalog) SetMarkInfo(mi _d.PdfObject) {
	_db := _d.MakeIndirectObject(mi)
	_ab.Object.Set("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f", _db)
	_ab._g.Objects = append(_ab._g.Objects, _db)
}
func (_aec *OutputIntents) Len() int { return _aec._fb.Len() }

type Document struct {
	ID             [2]string
	Version        _d.Version
	Objects        []_d.PdfObject
	Info           _d.PdfObject
	Crypt          *_d.PdfCrypt
	UseHashBasedID bool
}
type OutputIntents struct {
	_fb  *_d.PdfObjectArray
	_fee *Document
	_cba *_d.PdfIndirectObject
}

func (_fc Page) FindXObjectImages() ([]*Image, error) {
	_af, _aa := _fc.GetResourcesXObject()
	if !_aa {
		return nil, nil
	}
	var _bg []*Image
	var _fggf error
	_bda := map[*_d.PdfObjectStream]int{}
	_dab := map[*_d.PdfObjectStream]struct{}{}
	var _beg int
	for _, _fcc := range _af.Keys() {
		_bcf, _fcd := _d.GetStream(_af.Get(_fcc))
		if !_fcd {
			continue
		}
		if _, _df := _bda[_bcf]; _df {
			continue
		}
		_ebg, _edd := _d.GetName(_bcf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_edd || _ebg.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_dbd := Image{BitsPerComponent: 8, Stream: _bcf, Name: string(_fcc)}
		if _dbd.Colorspace, _fggf = _fdb(_bcf.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _fggf != nil {
			_a.Log.Error("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073", _fggf)
			continue
		}
		if _deg, _gcf := _d.GetIntVal(_bcf.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _gcf {
			_dbd.BitsPerComponent = _deg
		}
		if _dcd, _ace := _d.GetIntVal(_bcf.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _ace {
			_dbd.Width = _dcd
		}
		if _feaef, _cbed := _d.GetIntVal(_bcf.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _cbed {
			_dbd.Height = _feaef
		}
		if _cfd, _dbe := _d.GetStream(_bcf.Get("\u0053\u004d\u0061s\u006b")); _dbe {
			_dbd.SMask = &ImageSMask{Image: &_dbd, Stream: _cfd}
			_dab[_cfd] = struct{}{}
		}
		switch _dbd.Colorspace {
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_dbd.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_dbd.ColorComponents = 1
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_dbd.ColorComponents = 4
		default:
			_dbd.ColorComponents = -1
		}
		_bda[_bcf] = _beg
		_bg = append(_bg, &_dbd)
		_beg++
	}
	var _agg []int
	for _, _cbb := range _bg {
		if _cbb.SMask != nil {
			_gcb, _cdd := _bda[_cbb.SMask.Stream]
			if _cdd {
				_agg = append(_agg, _gcb)
			}
		}
	}
	_acg := make([]*Image, len(_bg)-len(_agg))
	_beg = 0
_efe:
	for _ega, _ccde := range _bg {
		for _, _cac := range _agg {
			if _ega == _cac {
				continue _efe
			}
		}
		_acg[_beg] = _ccde
		_beg++
	}
	return _bg, nil
}
func (_ccf *OutputIntents) Add(oi _d.PdfObject) error {
	_eg, _fbf := oi.(*_d.PdfObjectDictionary)
	if !_fbf {
		return _b.New("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _cdb, _aed := _d.GetStream(_eg.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _aed {
		_ccf._fee.Objects = append(_ccf._fee.Objects, _cdb)
	}
	_dc, _ef := oi.(*_d.PdfIndirectObject)
	if !_ef {
		_dc = _d.MakeIndirectObject(oi)
	}
	if _ccf._fb == nil {
		_ccf._fb = _d.MakeArray(_dc)
	} else {
		_ccf._fb.Append(_dc)
	}
	_ccf._fee.Objects = append(_ccf._fee.Objects, _dc)
	return nil
}
func (_gga Page) GetContents() ([]Content, bool) {
	_ce, _dda := _d.GetArray(_gga.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_dda {
		_dgb, _dec := _d.GetStream(_gga.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if !_dec {
			return nil, false
		}
		return []Content{{Stream: _dgb, _eeg: _gga, _dfb: 0}}, true
	}
	_ac := make([]Content, _ce.Len())
	for _bcd, _faf := range _ce.Elements() {
		_gad, _bea := _d.GetStream(_faf)
		if !_bea {
			continue
		}
		_ac[_bcd] = Content{Stream: _gad, _eeg: _gga, _dfb: _bcd}
	}
	return _ac, true
}
func (_cb *Catalog) SetOutputIntents(outputIntents *OutputIntents) {
	if _ecc := _cb.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"); _ecc != nil {
		for _ag, _fea := range _cb._g.Objects {
			if _fea == _ecc {
				if outputIntents._cba == _ecc {
					return
				}
				_cb._g.Objects = append(_cb._g.Objects[:_ag], _cb._g.Objects[_ag+1:]...)
				break
			}
		}
	}
	_fd := outputIntents._cba
	if _fd == nil {
		_fd = _d.MakeIndirectObject(outputIntents._fb)
	}
	_cb.Object.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _fd)
	_cb._g.Objects = append(_cb._g.Objects, _fd)
}
func (_bdb *Page) Number() int                         { return _bdb._bag }
func (_eca *Catalog) NewOutputIntents() *OutputIntents { return &OutputIntents{_fee: _eca._g} }
func (_gg *Document) GetPages() ([]Page, bool) {
	_bbc, _eaa := _gg.FindCatalog()
	if !_eaa {
		return nil, false
	}
	return _bbc.GetPages()
}

type Page struct {
	_bag   int
	Object *_d.PdfObjectDictionary
	_fgg   *Document
}
type OutputIntent struct {
	Object *_d.PdfObjectDictionary
}

func (_bc Page) GetResourcesXObject() (*_d.PdfObjectDictionary, bool) {
	_cgf, _edg := _bc.GetResources()
	if !_edg {
		return nil, false
	}
	return _d.GetDict(_cgf.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
}
func (_fbb Page) FindXObjectForms() []*_d.PdfObjectStream {
	_aedg, _bgf := _fbb.GetResourcesXObject()
	if !_bgf {
		return nil
	}
	_ege := map[*_d.PdfObjectStream]struct{}{}
	var _abe func(_bab *_d.PdfObjectDictionary, _egc map[*_d.PdfObjectStream]struct{})
	_abe = func(_cab *_d.PdfObjectDictionary, _cbdf map[*_d.PdfObjectStream]struct{}) {
		for _, _acd := range _cab.Keys() {
			_gdad, _cdbc := _d.GetStream(_cab.Get(_acd))
			if !_cdbc {
				continue
			}
			if _, _gec := _cbdf[_gdad]; _gec {
				continue
			}
			_dcb, _fda := _d.GetName(_gdad.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
			if !_fda || _dcb.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_cbdf[_gdad] = struct{}{}
			_bac, _fda := _d.GetDict(_gdad.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_fda {
				continue
			}
			_ddb, _cdbg := _d.GetDict(_bac.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
			if _cdbg {
				_abe(_ddb, _cbdf)
			}
		}
	}
	_abe(_aedg, _ege)
	var _edcc []*_d.PdfObjectStream
	for _cacc := range _ege {
		_edcc = append(_edcc, _cacc)
	}
	return _edcc
}
func (_ead *Document) AddIndirectObject(indirect *_d.PdfIndirectObject) {
	for _, _efg := range _ead.Objects {
		if _efg == indirect {
			return
		}
	}
	_ead.Objects = append(_ead.Objects, indirect)
}
func _fdb(_fg _d.PdfObject) (_d.PdfObjectName, error) {
	var _cbd *_d.PdfObjectName
	var _ede *_d.PdfObjectArray
	if _bbcb, _fa := _fg.(*_d.PdfIndirectObject); _fa {
		if _cag, _fab := _bbcb.PdfObject.(*_d.PdfObjectArray); _fab {
			_ede = _cag
		} else if _ge, _fae := _bbcb.PdfObject.(*_d.PdfObjectName); _fae {
			_cbd = _ge
		}
	} else if _gag, _dge := _fg.(*_d.PdfObjectArray); _dge {
		_ede = _gag
	} else if _feae, _de := _fg.(*_d.PdfObjectName); _de {
		_cbd = _feae
	}
	if _cbd != nil {
		switch *_cbd {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_cbd, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_cbd, nil
		}
	}
	if _ede != nil && _ede.Len() > 0 {
		if _fed, _edf := _ede.Get(0).(*_d.PdfObjectName); _edf {
			switch *_fed {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _ede.Len() == 1 {
					return *_fed, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_fed, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_fed, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_fed, nil
			}
		}
	}
	return "", nil
}
func (_eec *Content) SetData(data []byte) error {
	_afb, _dcba := _d.MakeStream(data, _d.NewFlateEncoder())
	if _dcba != nil {
		return _dcba
	}
	_bca, _ad := _d.GetArray(_eec._eeg.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_ad && _eec._dfb == 0 {
		_eec._eeg.Object.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _afb)
	} else {
		if _dcba = _bca.Set(_eec._dfb, _afb); _dcba != nil {
			return _dcba
		}
	}
	_eec._eeg._fgg.Objects = append(_eec._eeg._fgg.Objects, _afb)
	return nil
}

type Image struct {
	Name             string
	Width            int
	Height           int
	Colorspace       _d.PdfObjectName
	ColorComponents  int
	BitsPerComponent int
	SMask            *ImageSMask
	Stream           *_d.PdfObjectStream
}

func (_bb *Catalog) HasMetadata() bool {
	_ga := _bb.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _ga != nil
}
func (_gfg *Catalog) GetOutputIntents() (*OutputIntents, bool) {
	_cdg := _gfg.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073")
	if _cdg == nil {
		return nil, false
	}
	_edc, _gaa := _d.GetIndirect(_cdg)
	if !_gaa {
		return nil, false
	}
	_gab, _cbe := _d.GetArray(_edc.PdfObject)
	if !_cbe {
		return nil, false
	}
	return &OutputIntents{_cba: _edc, _fb: _gab, _fee: _gfg._g}, true
}

type Catalog struct {
	Object *_d.PdfObjectDictionary
	_g     *Document
}
type Content struct {
	Stream *_d.PdfObjectStream
	_dfb   int
	_eeg   Page
}

func (_ba *Document) AddStream(stream *_d.PdfObjectStream) {
	for _, _dag := range _ba.Objects {
		if _dag == stream {
			return
		}
	}
	_ba.Objects = append(_ba.Objects, stream)
}
