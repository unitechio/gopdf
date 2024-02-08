package docutil

import (
	_c "errors"
	_d "fmt"

	_e "bitbucket.org/shenghui0779/gopdf/common"
	_da "bitbucket.org/shenghui0779/gopdf/core"
)

func (_ddee Page) GetResources() (*_da.PdfObjectDictionary, bool) {
	return _da.GetDict(_ddee.Object.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
}

type OutputIntent struct{ Object *_da.PdfObjectDictionary }

func (_fde *Catalog) SetMarkInfo(mi _da.PdfObject) {
	_eg := _da.MakeIndirectObject(mi)
	_fde.Object.Set("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f", _eg)
	_fde._f.Objects = append(_fde._f.Objects, _eg)
}
func (_aed *OutputIntents) Get(i int) (OutputIntent, bool) {
	if _aed._aea == nil {
		return OutputIntent{}, false
	}
	if i >= _aed._aea.Len() {
		return OutputIntent{}, false
	}
	_db := _aed._aea.Get(i)
	_bae, _gec := _da.GetIndirect(_db)
	if !_gec {
		_dab, _gee := _da.GetDict(_db)
		return OutputIntent{Object: _dab}, _gee
	}
	_dde, _ege := _da.GetDict(_bae.PdfObject)
	return OutputIntent{Object: _dde}, _ege
}
func (_de *Page) Number() int                         { return _de._aeb }
func (_ae *Catalog) NewOutputIntents() *OutputIntents { return &OutputIntents{_egb: _ae._f} }
func (_cb *Catalog) GetMetadata() (*_da.PdfObjectStream, bool) {
	_af, _aff := _da.GetStream(_cb.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	return _af, _aff
}
func _ab(_bd _da.PdfObject) (_da.PdfObjectName, error) {
	var _ee *_da.PdfObjectName
	var _bdd *_da.PdfObjectArray
	if _cba, _dbg := _bd.(*_da.PdfIndirectObject); _dbg {
		if _cbe, _gf := _cba.PdfObject.(*_da.PdfObjectArray); _gf {
			_bdd = _cbe
		} else if _ffd, _ebe := _cba.PdfObject.(*_da.PdfObjectName); _ebe {
			_ee = _ffd
		}
	} else if _eaa, _gd := _bd.(*_da.PdfObjectArray); _gd {
		_bdd = _eaa
	} else if _dbge, _cag := _bd.(*_da.PdfObjectName); _cag {
		_ee = _dbge
	}
	if _ee != nil {
		switch *_ee {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_ee, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_ee, nil
		}
	}
	if _bdd != nil && _bdd.Len() > 0 {
		if _bec, _cgf := _bdd.Get(0).(*_da.PdfObjectName); _cgf {
			switch *_bec {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _bdd.Len() == 1 {
					return *_bec, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_bec, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_bec, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_bec, nil
			}
		}
	}
	return "", nil
}
func (_ad *Catalog) HasMetadata() bool {
	_ceg := _ad.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _ceg != nil
}

type Document struct {
	ID             [2]string
	Version        _da.Version
	Objects        []_da.PdfObject
	Info           _da.PdfObject
	Crypt          *_da.PdfCrypt
	UseHashBasedID bool
}
type Catalog struct {
	Object *_da.PdfObjectDictionary
	_f     *Document
}

func (_ec *Catalog) SetMetadata(data []byte) error {
	_fd, _ag := _da.MakeStream(data, nil)
	if _ag != nil {
		return _ag
	}
	_fd.Set("\u0054\u0079\u0070\u0065", _da.MakeName("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	_fd.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _da.MakeName("\u0058\u004d\u004c"))
	_ec.Object.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _fd)
	_ec._f.Objects = append(_ec._f.Objects, _fd)
	return nil
}

type ImageSMask struct {
	Image  *Image
	Stream *_da.PdfObjectStream
}

func (_g *Catalog) SetVersion() {
	_g.Object.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _da.MakeName(_d.Sprintf("\u0025\u0064\u002e%\u0064", _g._f.Version.Major, _g._f.Version.Minor)))
}
func (_df Page) GetContents() ([]Content, bool) {
	_geeb, _ebee := _da.GetArray(_df.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_ebee {
		_cge, _baeg := _da.GetStream(_df.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if !_baeg {
			return nil, false
		}
		return []Content{{Stream: _cge, _bf: _df, _ddf: 0}}, true
	}
	_dbe := make([]Content, _geeb.Len())
	for _eba, _ccc := range _geeb.Elements() {
		_dgg, _dfc := _da.GetStream(_ccc)
		if !_dfc {
			continue
		}
		_dbe[_eba] = Content{Stream: _dgg, _bf: _df, _ddf: _eba}
	}
	return _dbe, true
}
func (_b *Catalog) GetPages() ([]Page, bool) {
	_ac, _ef := _da.GetDict(_b.Object.Get("\u0050\u0061\u0067e\u0073"))
	if !_ef {
		return nil, false
	}
	_ce, _ba := _da.GetArray(_ac.Get("\u004b\u0069\u0064\u0073"))
	if !_ba {
		return nil, false
	}
	_fc := make([]Page, _ce.Len())
	for _gg, _gc := range _ce.Elements() {
		_eb, _cg := _da.GetDict(_gc)
		if !_cg {
			continue
		}
		_fc[_gg] = Page{Object: _eb, _aeb: _gg + 1, _gfb: _b._f}
	}
	return _fc, true
}
func (_bac *Document) GetPages() ([]Page, bool) {
	_dcf, _beb := _bac.FindCatalog()
	if !_beb {
		return nil, false
	}
	return _dcf.GetPages()
}
func (_cgd Content) GetData() ([]byte, error) {
	_afeb, _cad := _da.NewEncoderFromStream(_cgd.Stream)
	if _cad != nil {
		return nil, _cad
	}
	_bee, _cad := _afeb.DecodeStream(_cgd.Stream)
	if _cad != nil {
		return nil, _cad
	}
	return _bee, nil
}

type Content struct {
	Stream *_da.PdfObjectStream
	_ddf   int
	_bf    Page
}

func (_fdf Page) GetResourcesXObject() (*_da.PdfObjectDictionary, bool) {
	_bba, _cdcc := _fdf.GetResources()
	if !_cdcc {
		return nil, false
	}
	return _da.GetDict(_bba.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
}
func (_ca *Catalog) GetOutputIntents() (*OutputIntents, bool) {
	_dc := _ca.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073")
	if _dc == nil {
		return nil, false
	}
	_ge, _cea := _da.GetIndirect(_dc)
	if !_cea {
		return nil, false
	}
	_dac, _fff := _da.GetArray(_ge.PdfObject)
	if !_fff {
		return nil, false
	}
	return &OutputIntents{_fcb: _ge, _aea: _dac, _egb: _ca._f}, true
}

type OutputIntents struct {
	_aea *_da.PdfObjectArray
	_egb *Document
	_fcb *_da.PdfIndirectObject
}

func (_agd *Document) FindCatalog() (*Catalog, bool) {
	var _ebg *_da.PdfObjectDictionary
	for _, _gce := range _agd.Objects {
		_acg, _bag := _da.GetDict(_gce)
		if !_bag {
			continue
		}
		if _fdg, _aca := _da.GetName(_acg.Get("\u0054\u0079\u0070\u0065")); _aca && *_fdg == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
			_ebg = _acg
			break
		}
	}
	if _ebg == nil {
		return nil, false
	}
	return &Catalog{Object: _ebg, _f: _agd}, true
}
func (_dcd *OutputIntents) Len() int { return _dcd._aea.Len() }
func (_afc *Content) SetData(data []byte) error {
	_bef, _efc := _da.MakeStream(data, _da.NewFlateEncoder())
	if _efc != nil {
		return _efc
	}
	_ccd, _bgd := _da.GetArray(_afc._bf.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_bgd && _afc._ddf == 0 {
		_afc._bf.Object.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _bef)
	} else {
		if _efc = _ccd.Set(_afc._ddf, _bef); _efc != nil {
			return _efc
		}
	}
	_afc._bf._gfb.Objects = append(_afc._bf._gfb.Objects, _bef)
	return nil
}

type Image struct {
	Name             string
	Width            int
	Height           int
	Colorspace       _da.PdfObjectName
	ColorComponents  int
	BitsPerComponent int
	SMask            *ImageSMask
	Stream           *_da.PdfObjectStream
}

func (_agf Page) FindXObjectImages() ([]*Image, error) {
	_aaab, _dbf := _agf.GetResourcesXObject()
	if !_dbf {
		return nil, nil
	}
	var _eaaa []*Image
	var _bc error
	_dgb := map[*_da.PdfObjectStream]int{}
	_eae := map[*_da.PdfObjectStream]struct{}{}
	var _cac int
	for _, _ggg := range _aaab.Keys() {
		_dad, _agb := _da.GetStream(_aaab.Get(_ggg))
		if !_agb {
			continue
		}
		if _, _dadf := _dgb[_dad]; _dadf {
			continue
		}
		_acd, _cf := _da.GetName(_dad.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_cf || _acd.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_dba := Image{BitsPerComponent: 8, Stream: _dad, Name: string(_ggg)}
		if _dba.Colorspace, _bc = _ab(_dad.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _bc != nil {
			_e.Log.Error("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073", _bc)
			continue
		}
		if _cfa, _bg := _da.GetIntVal(_dad.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _bg {
			_dba.BitsPerComponent = _cfa
		}
		if _cef, _efd := _da.GetIntVal(_dad.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _efd {
			_dba.Width = _cef
		}
		if _dbfb, _afe := _da.GetIntVal(_dad.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _afe {
			_dba.Height = _dbfb
		}
		if _dga, _dfg := _da.GetStream(_dad.Get("\u0053\u004d\u0061s\u006b")); _dfg {
			_dba.SMask = &ImageSMask{Image: &_dba, Stream: _dga}
			_eae[_dga] = struct{}{}
		}
		switch _dba.Colorspace {
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_dba.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_dba.ColorComponents = 1
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_dba.ColorComponents = 4
		default:
			_dba.ColorComponents = -1
		}
		_dgb[_dad] = _cac
		_eaaa = append(_eaaa, &_dba)
		_cac++
	}
	var _cca []int
	for _, _gdd := range _eaaa {
		if _gdd.SMask != nil {
			_cfb, _affg := _dgb[_gdd.SMask.Stream]
			if _affg {
				_cca = append(_cca, _cfb)
			}
		}
	}
	_ceb := make([]*Image, len(_eaaa)-len(_cca))
	_cac = 0
_fda:
	for _ed, _cab := range _eaaa {
		for _, _ecc := range _cca {
			if _ed == _ecc {
				continue _fda
			}
		}
		_ceb[_cac] = _cab
		_cac++
	}
	return _eaaa, nil
}

type Page struct {
	_aeb   int
	Object *_da.PdfObjectDictionary
	_gfb   *Document
}

func (_fffb *OutputIntents) Add(oi _da.PdfObject) error {
	_ea, _be := oi.(*_da.PdfObjectDictionary)
	if !_be {
		return _c.New("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _ead, _fg := _da.GetStream(_ea.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _fg {
		_fffb._egb.Objects = append(_fffb._egb.Objects, _ead)
	}
	_aaa, _cc := oi.(*_da.PdfIndirectObject)
	if !_cc {
		_aaa = _da.MakeIndirectObject(oi)
	}
	if _fffb._aea == nil {
		_fffb._aea = _da.MakeArray(_aaa)
	} else {
		_fffb._aea.Append(_aaa)
	}
	_fffb._egb.Objects = append(_fffb._egb.Objects, _aaa)
	return nil
}
func (_ff *Catalog) GetMarkInfo() (*_da.PdfObjectDictionary, bool) {
	_cd, _cgb := _da.GetDict(_ff.Object.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"))
	return _cd, _cgb
}
func (_dd *Catalog) SetOutputIntents(outputIntents *OutputIntents) {
	if _cdc := _dd.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"); _cdc != nil {
		for _bb, _aga := range _dd._f.Objects {
			if _aga == _cdc {
				if outputIntents._fcb == _cdc {
					return
				}
				_dd._f.Objects = append(_dd._f.Objects[:_bb], _dd._f.Objects[_bb+1:]...)
				break
			}
		}
	}
	_aa := outputIntents._fcb
	if _aa == nil {
		_aa = _da.MakeIndirectObject(outputIntents._aea)
	}
	_dd.Object.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _aa)
	_dd._f.Objects = append(_dd._f.Objects, _aa)
}
func (_cdb Page) FindXObjectForms() []*_da.PdfObjectStream {
	_fad, _abd := _cdb.GetResourcesXObject()
	if !_abd {
		return nil
	}
	_gaa := map[*_da.PdfObjectStream]struct{}{}
	var _fbd func(_eeg *_da.PdfObjectDictionary, _bga map[*_da.PdfObjectStream]struct{})
	_fbd = func(_cfe *_da.PdfObjectDictionary, _cfc map[*_da.PdfObjectStream]struct{}) {
		for _, _gad := range _cfe.Keys() {
			_gag, _gab := _da.GetStream(_cfe.Get(_gad))
			if !_gab {
				continue
			}
			if _, _abc := _cfc[_gag]; _abc {
				continue
			}
			_ffa, _fbb := _da.GetName(_gag.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
			if !_fbb || _ffa.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_cfc[_gag] = struct{}{}
			_cee, _fbb := _da.GetDict(_gag.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_fbb {
				continue
			}
			_abe, _adf := _da.GetDict(_cee.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
			if _adf {
				_fbd(_abe, _cfc)
			}
		}
	}
	_fbd(_fad, _gaa)
	var _cfee []*_da.PdfObjectStream
	for _dcb := range _gaa {
		_cfee = append(_cfee, _dcb)
	}
	return _cfee
}
func (_ga *Document) AddIndirectObject(indirect *_da.PdfIndirectObject) {
	for _, _fab := range _ga.Objects {
		if _fab == indirect {
			return
		}
	}
	_ga.Objects = append(_ga.Objects, indirect)
}
func (_dg *Document) AddStream(stream *_da.PdfObjectStream) {
	for _, _ega := range _dg.Objects {
		if _ega == stream {
			return
		}
	}
	_dg.Objects = append(_dg.Objects, stream)
}
