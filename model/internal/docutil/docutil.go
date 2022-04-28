package docutil

import (
	_cg "errors"
	_ca "fmt"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/core"
)

type OutputIntents struct {
	_ce *_b.PdfObjectArray
	_gg *Document
	_ea *_b.PdfIndirectObject
}

func (_be *Catalog) HasMetadata() bool {
	_cf := _be.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _cf != nil
}
func (_caa Content) GetData() ([]byte, error) {
	_ge, _acac := _b.NewEncoderFromStream(_caa.Stream)
	if _acac != nil {
		return nil, _acac
	}
	_gga, _acac := _ge.DecodeStream(_caa.Stream)
	if _acac != nil {
		return nil, _acac
	}
	return _gga, nil
}
func (_efc *Document) GetPages() ([]Page, bool) {
	_acc, _ec := _efc.FindCatalog()
	if !_ec {
		return nil, false
	}
	return _acc.GetPages()
}
func (_ebc Page) FindXObjectForms() []*_b.PdfObjectStream {
	_bb, _ddb := _ebc.GetResourcesXObject()
	if !_ddb {
		return nil
	}
	_fda := map[*_b.PdfObjectStream]struct{}{}
	var _ag func(_bga *_b.PdfObjectDictionary, _ced map[*_b.PdfObjectStream]struct{})
	_ag = func(_feb *_b.PdfObjectDictionary, _aaf map[*_b.PdfObjectStream]struct{}) {
		for _, _cfd := range _feb.Keys() {
			_af, _dbe := _b.GetStream(_feb.Get(_cfd))
			if !_dbe {
				continue
			}
			if _, _gce := _aaf[_af]; _gce {
				continue
			}
			_cab, _bbb := _b.GetName(_af.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
			if !_bbb || _cab.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_aaf[_af] = struct{}{}
			_bcf, _bbb := _b.GetDict(_af.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_bbb {
				continue
			}
			_ecb, _ceeg := _b.GetDict(_bcf.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
			if _ceeg {
				_ag(_ecb, _aaf)
			}
		}
	}
	_ag(_bb, _fda)
	var _ggbc []*_b.PdfObjectStream
	for _ada := range _fda {
		_ggbc = append(_ggbc, _ada)
	}
	return _ggbc
}

type Catalog struct {
	Object *_b.PdfObjectDictionary
	_ba    *Document
}
type ImageSMask struct {
	Image  *Image
	Stream *_b.PdfObjectStream
}
type Page struct {
	_bdb   int
	Object *_b.PdfObjectDictionary
	_cgf   *Document
}

func (_bdd *Catalog) NewOutputIntents() *OutputIntents { return &OutputIntents{_gg: _bdd._ba} }
func (_ff *Catalog) SetOutputIntents(outputIntents *OutputIntents) {
	if _bde := _ff.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"); _bde != nil {
		for _g, _gd := range _ff._ba.Objects {
			if _gd == _bde {
				if outputIntents._ea == _bde {
					return
				}
				_ff._ba.Objects = append(_ff._ba.Objects[:_g], _ff._ba.Objects[_g+1:]...)
				break
			}
		}
	}
	_abb := outputIntents._ea
	if _abb == nil {
		_abb = _b.MakeIndirectObject(outputIntents._ce)
	}
	_ff.Object.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _abb)
	_ff._ba.Objects = append(_ff._ba.Objects, _abb)
}
func (_df *Catalog) SetMetadata(data []byte) error {
	_efd, _cd := _b.MakeStream(data, nil)
	if _cd != nil {
		return _cd
	}
	_efd.Set("\u0054\u0079\u0070\u0065", _b.MakeName("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	_efd.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _b.MakeName("\u0058\u004d\u004c"))
	_df.Object.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _efd)
	_df._ba.Objects = append(_df._ba.Objects, _efd)
	return nil
}
func (_bc *OutputIntents) Get(i int) (OutputIntent, bool) {
	if _bc._ce == nil {
		return OutputIntent{}, false
	}
	if i >= _bc._ce.Len() {
		return OutputIntent{}, false
	}
	_adg := _bc._ce.Get(i)
	_gdcb, _ee := _b.GetIndirect(_adg)
	if !_ee {
		_eed, _dc := _b.GetDict(_adg)
		return OutputIntent{Object: _eed}, _dc
	}
	_eff, _bfe := _b.GetDict(_gdcb.PdfObject)
	return OutputIntent{Object: _eff}, _bfe
}
func (_dda *Document) FindCatalog() (*Catalog, bool) {
	var _fdf *_b.PdfObjectDictionary
	for _, _fa := range _dda.Objects {
		_aa, _da := _b.GetDict(_fa)
		if !_da {
			continue
		}
		if _ggd, _cef := _b.GetName(_aa.Get("\u0054\u0079\u0070\u0065")); _cef && *_ggd == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
			_fdf = _aa
			break
		}
	}
	if _fdf == nil {
		return nil, false
	}
	return &Catalog{Object: _fdf, _ba: _dda}, true
}
func (_fef *Document) AddIndirectObject(indirect *_b.PdfIndirectObject) {
	for _, _dfe := range _fef.Objects {
		if _dfe == indirect {
			return
		}
	}
	_fef.Objects = append(_fef.Objects, indirect)
}
func (_bfb *Catalog) GetOutputIntents() (*OutputIntents, bool) {
	_bda := _bfb.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073")
	if _bda == nil {
		return nil, false
	}
	_fd, _ccb := _b.GetIndirect(_bda)
	if !_ccb {
		return nil, false
	}
	_bea, _dfdg := _b.GetArray(_fd.PdfObject)
	if !_dfdg {
		return nil, false
	}
	return &OutputIntents{_ea: _fd, _ce: _bea, _gg: _bfb._ba}, true
}
func (_gaa Page) FindXObjectImages() ([]*Image, error) {
	_cae, _caf := _gaa.GetResourcesXObject()
	if !_caf {
		return nil, nil
	}
	var _fefb []*Image
	var _bdab error
	_cfc := map[*_b.PdfObjectStream]int{}
	_abf := map[*_b.PdfObjectStream]struct{}{}
	var _bdeg int
	for _, _cgd := range _cae.Keys() {
		_cdd, _fefbf := _b.GetStream(_cae.Get(_cgd))
		if !_fefbf {
			continue
		}
		if _, _ebd := _cfc[_cdd]; _ebd {
			continue
		}
		_ccd, _gad := _b.GetName(_cdd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_gad || _ccd.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_de := Image{BitsPerComponent: 8, Stream: _cdd, Name: string(_cgd)}
		if _de.Colorspace, _bdab = _fb(_cdd.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _bdab != nil {
			_a.Log.Error("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073", _bdab)
			continue
		}
		if _db, _fff := _b.GetIntVal(_cdd.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _fff {
			_de.BitsPerComponent = _db
		}
		if _def, _dg := _b.GetIntVal(_cdd.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _dg {
			_de.Width = _def
		}
		if _cee, _ecd := _b.GetIntVal(_cdd.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _ecd {
			_de.Height = _cee
		}
		if _aab, _ecg := _b.GetStream(_cdd.Get("\u0053\u004d\u0061s\u006b")); _ecg {
			_de.SMask = &ImageSMask{Image: &_de, Stream: _aab}
			_abf[_aab] = struct{}{}
		}
		switch _de.Colorspace {
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_de.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_de.ColorComponents = 1
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_de.ColorComponents = 4
		default:
			_de.ColorComponents = -1
		}
		_cfc[_cdd] = _bdeg
		_fefb = append(_fefb, &_de)
		_bdeg++
	}
	var _gb []int
	for _, _cca := range _fefb {
		if _cca.SMask != nil {
			_ggb, _gf := _cfc[_cca.SMask.Stream]
			if _gf {
				_gb = append(_gb, _ggb)
			}
		}
	}
	_eedd := make([]*Image, len(_fefb)-len(_gb))
	_bdeg = 0
_ed:
	for _gag, _fac := range _fefb {
		for _, _abd := range _gb {
			if _gag == _abd {
				continue _ed
			}
		}
		_eedd[_bdeg] = _fac
		_bdeg++
	}
	return _fefb, nil
}
func (_dd *OutputIntents) Add(oi _b.PdfObject) error {
	_gdc, _eaf := oi.(*_b.PdfObjectDictionary)
	if !_eaf {
		return _cg.New("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _cce, _bdc := _b.GetStream(_gdc.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _bdc {
		_dd._gg.Objects = append(_dd._gg.Objects, _cce)
	}
	_ac, _bfbf := oi.(*_b.PdfIndirectObject)
	if !_bfbf {
		_ac = _b.MakeIndirectObject(oi)
	}
	if _dd._ce == nil {
		_dd._ce = _b.MakeArray(_ac)
	} else {
		_dd._ce.Append(_ac)
	}
	_dd._gg.Objects = append(_dd._gg.Objects, _ac)
	return nil
}

type OutputIntent struct{ Object *_b.PdfObjectDictionary }

func (_ggcb Page) GetContents() ([]Content, bool) {
	_ecc, _egb := _b.GetArray(_ggcb.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_egb {
		_bfg, _bge := _b.GetStream(_ggcb.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if !_bge {
			return nil, false
		}
		return []Content{{Stream: _bfg, _efcb: _ggcb, _ffa: 0}}, true
	}
	_ggdb := make([]Content, _ecc.Len())
	for _egf, _bfeb := range _ecc.Elements() {
		_fbf, _cgfb := _b.GetStream(_bfeb)
		if !_cgfb {
			continue
		}
		_ggdb[_egf] = Content{Stream: _fbf, _efcb: _ggcb, _ffa: _egf}
	}
	return _ggdb, true
}
func (_dcab Page) GetResourcesXObject() (*_b.PdfObjectDictionary, bool) {
	_eb, _eeg := _dcab.GetResources()
	if !_eeg {
		return nil, false
	}
	return _b.GetDict(_eb.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
}

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

func (_bafa *Document) AddStream(stream *_b.PdfObjectStream) {
	for _, _daa := range _bafa.Objects {
		if _daa == stream {
			return
		}
	}
	_bafa.Objects = append(_bafa.Objects, stream)
}
func (_beac Page) GetResources() (*_b.PdfObjectDictionary, bool) {
	return _b.GetDict(_beac.Object.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
}
func (_ab *Catalog) SetVersion() {
	_ab.Object.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _b.MakeName(_ca.Sprintf("\u0025\u0064\u002e%\u0064", _ab._ba.Version.Major, _ab._ba.Version.Minor)))
}
func (_d *Catalog) GetPages() ([]Page, bool) {
	_f, _cb := _b.GetDict(_d.Object.Get("\u0050\u0061\u0067e\u0073"))
	if !_cb {
		return nil, false
	}
	_cc, _bd := _b.GetArray(_f.Get("\u004b\u0069\u0064\u0073"))
	if !_bd {
		return nil, false
	}
	_ccg := make([]Page, _cc.Len())
	for _baf, _e := range _cc.Elements() {
		_ef, _eg := _b.GetDict(_e)
		if !_eg {
			continue
		}
		_ccg[_baf] = Page{Object: _ef, _bdb: _baf + 1, _cgf: _d._ba}
	}
	return _ccg, true
}

type Content struct {
	Stream *_b.PdfObjectStream
	_ffa   int
	_efcb  Page
}

func (_fgb *Content) SetData(data []byte) error {
	_bfd, _dbg := _b.MakeStream(data, _b.NewFlateEncoder())
	if _dbg != nil {
		return _dbg
	}
	_fbe, _cfe := _b.GetArray(_fgb._efcb.Object.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_cfe && _fgb._ffa == 0 {
		_fgb._efcb.Object.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _bfd)
	} else {
		if _dbg = _fbe.Set(_fgb._ffa, _bfd); _dbg != nil {
			return _dbg
		}
	}
	_fgb._efcb._cgf.Objects = append(_fgb._efcb._cgf.Objects, _bfd)
	return nil
}
func (_bf *Catalog) GetMarkInfo() (*_b.PdfObjectDictionary, bool) {
	_dfd, _ae := _b.GetDict(_bf.Object.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"))
	return _dfd, _ae
}
func (_ga *OutputIntents) Len() int { return _ga._ce.Len() }
func (_ad *Catalog) SetMarkInfo(mi _b.PdfObject) {
	_fg := _b.MakeIndirectObject(mi)
	_ad.Object.Set("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f", _fg)
	_ad._ba.Objects = append(_ad._ba.Objects, _fg)
}

type Document struct {
	ID             [2]string
	Version        _b.Version
	Objects        []_b.PdfObject
	Info           _b.PdfObject
	Crypt          *_b.PdfCrypt
	UseHashBasedID bool
}

func (_bag *Catalog) GetMetadata() (*_b.PdfObjectStream, bool) {
	_fe, _bg := _b.GetStream(_bag.Object.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"))
	return _fe, _bg
}
func (_cgdf *Page) Number() int { return _cgdf._bdb }
func _fb(_fgdf _b.PdfObject) (_b.PdfObjectName, error) {
	var _dfa *_b.PdfObjectName
	var _aca *_b.PdfObjectArray
	if _dca, _bed := _fgdf.(*_b.PdfIndirectObject); _bed {
		if _bfc, _fbc := _dca.PdfObject.(*_b.PdfObjectArray); _fbc {
			_aca = _bfc
		} else if _daaa, _fba := _dca.PdfObject.(*_b.PdfObjectName); _fba {
			_dfa = _daaa
		}
	} else if _dfdc, _fgdg := _fgdf.(*_b.PdfObjectArray); _fgdg {
		_aca = _dfdc
	} else if _fec, _ggc := _fgdf.(*_b.PdfObjectName); _ggc {
		_dfa = _fec
	}
	if _dfa != nil {
		switch *_dfa {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_dfa, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_dfa, nil
		}
	}
	if _aca != nil && _aca.Len() > 0 {
		if _bec, _dac := _aca.Get(0).(*_b.PdfObjectName); _dac {
			switch *_bec {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _aca.Len() == 1 {
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
