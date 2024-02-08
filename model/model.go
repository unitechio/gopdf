package model

import (
	_ba "bufio"
	_bc "bytes"
	_f "crypto/md5"
	_ec "crypto/rand"
	_fb "crypto/sha1"
	_bb "crypto/x509"
	_bad "encoding/binary"
	_be "encoding/hex"
	_bf "errors"
	_b "fmt"
	_ed "hash"
	_gd "image"
	_edg "image/color"
	_ "image/gif"
	_ "image/png"
	_cf "io"
	_gf "io/ioutil"
	_cg "math"
	_cb "math/rand"
	_eb "os"
	_c "regexp"
	_gc "sort"
	_fbb "strconv"
	_ga "strings"
	_e "sync"
	_a "time"
	_af "unicode"
	_gbf "unicode/utf8"

	_ag "bitbucket.org/shenghui0779/gopdf/common"
	_dg "bitbucket.org/shenghui0779/gopdf/core"
	_gbd "bitbucket.org/shenghui0779/gopdf/core/security"
	_afb "bitbucket.org/shenghui0779/gopdf/core/security/crypt"
	_ff "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_fc "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_fcd "bitbucket.org/shenghui0779/gopdf/internal/sampling"
	_bd "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_bcc "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_fec "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_eba "bitbucket.org/shenghui0779/gopdf/model/internal/docutil"
	_bbg "bitbucket.org/shenghui0779/gopdf/model/internal/fonts"
	_ecb "bitbucket.org/shenghui0779/gopdf/model/mdp"
	_fe "bitbucket.org/shenghui0779/gopdf/model/sigutil"
	_gfd "bitbucket.org/shenghui0779/gopdf/ps"
	_eg "github.com/unidoc/pkcs7"
	_bfc "github.com/unidoc/unitype"
	_ge "golang.org/x/xerrors"
)

func (_dbgf *PdfReader) newPdfAnnotationRichMediaFromDict(_cffb *_dg.PdfObjectDictionary) (*PdfAnnotationRichMedia, error) {
	_cbdeg := &PdfAnnotationRichMedia{}
	_cbdeg.RichMediaSettings = _cffb.Get("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073")
	_cbdeg.RichMediaContent = _cffb.Get("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074")
	return _cbdeg, nil
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the page using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened for the page.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
func (_ecfc *PdfPage) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	_cbdgf := map[*PdfAnnotation]bool{}
	_ffccb, _egef := _ecfc.GetAnnotations()
	if _egef != nil {
		return _egef
	}
	_dcbb := false
	for _, _cfae := range _ffccb {
		if opts.AnnotFilterFunc != nil {
			_cbdgf[_cfae] = opts.AnnotFilterFunc(_cfae)
		} else {
			_cbdgf[_cfae] = true
		}
		if _cbdgf[_cfae] {
			_dcbb = true
		}
	}
	if !_dcbb {
		return nil
	}
	return _ecfc.flattenFieldsWithOpts(appgen, opts, _cbdgf)
}

// GetPdfName returns the PDF name used to indicate the border style.
// (Table 166 p. 395).
func (_fccb *BorderStyle) GetPdfName() string {
	switch *_fccb {
	case BorderStyleSolid:
		return "\u0053"
	case BorderStyleDashed:
		return "\u0044"
	case BorderStyleBeveled:
		return "\u0042"
	case BorderStyleInset:
		return "\u0049"
	case BorderStyleUnderline:
		return "\u0055"
	}
	return ""
}

// PdfActionTrans represents a trans action.
type PdfActionTrans struct {
	*PdfAction
	Trans _dg.PdfObject
}

// DefaultImageHandler is the default implementation of the ImageHandler using the standard go library.
type DefaultImageHandler struct{}

// String returns string value of output intent for given type
// ISO_19005-2 6.2.3: GTS_PDFA1 value should be used for PDF/A-1, A-2 and A-3 at least
func (_cdcc PdfOutputIntentType) String() string {
	switch _cdcc {
	case PdfOutputIntentTypeA1:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeA2:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeA3:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeA4:
		return "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411"
	case PdfOutputIntentTypeX:
		return "\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0058"
	default:
		return "\u0055N\u0044\u0045\u0046\u0049\u004e\u0045D"
	}
}

// ToPdfObject implements interface PdfModel.
func (_bee *PdfActionImportData) ToPdfObject() _dg.PdfObject {
	_bee.PdfAction.ToPdfObject()
	_bga := _bee._cbd
	_fce := _bga.PdfObject.(*_dg.PdfObjectDictionary)
	_fce.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeImportData)))
	if _bee.F != nil {
		_fce.Set("\u0046", _bee.F.ToPdfObject())
	}
	return _bga
}

// NewPdfField returns an initialized PdfField.
func NewPdfField() *PdfField { return &PdfField{_egce: _dg.MakeIndirectObject(_dg.MakeDict())} }

// Items returns all children outline items.
func (_dadce *OutlineItem) Items() []*OutlineItem { return _dadce.Entries }

// ToPdfObject returns the PDF representation of the colorspace.
func (_fcdf *PdfColorspaceDeviceGray) ToPdfObject() _dg.PdfObject {
	return _dg.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079")
}
func (_eafa SignatureValidationResult) String() string {
	var _egacb _bc.Buffer
	_egacb.WriteString(_b.Sprintf("\u004ea\u006d\u0065\u003a\u0020\u0025\u0073\n", _eafa.Name))
	if _eafa.Date._bgfdb > 0 {
		_egacb.WriteString(_b.Sprintf("\u0044a\u0074\u0065\u003a\u0020\u0025\u0073\n", _eafa.Date.ToGoTime().String()))
	} else {
		_egacb.WriteString("\u0044\u0061\u0074\u0065 n\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_eafa.Reason) > 0 {
		_egacb.WriteString(_b.Sprintf("R\u0065\u0061\u0073\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _eafa.Reason))
	} else {
		_egacb.WriteString("N\u006f \u0072\u0065\u0061\u0073\u006f\u006e\u0020\u0073p\u0065\u0063\u0069\u0066ie\u0064\u000a")
	}
	if len(_eafa.Location) > 0 {
		_egacb.WriteString(_b.Sprintf("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _eafa.Location))
	} else {
		_egacb.WriteString("\u004c\u006f\u0063at\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_eafa.ContactInfo) > 0 {
		_egacb.WriteString(_b.Sprintf("\u0043\u006f\u006e\u0074\u0061\u0063\u0074\u0020\u0049\u006e\u0066\u006f:\u0020\u0025\u0073\u000a", _eafa.ContactInfo))
	} else {
		_egacb.WriteString("C\u006f\u006e\u0074\u0061\u0063\u0074 \u0069\u006e\u0066\u006f\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063i\u0066i\u0065\u0064\u000a")
	}
	_egacb.WriteString(_b.Sprintf("F\u0069\u0065\u006c\u0064\u0073\u003a\u0020\u0025\u0064\u000a", len(_eafa.Fields)))
	if _eafa.IsSigned {
		_egacb.WriteString("S\u0069\u0067\u006e\u0065\u0064\u003a \u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073\u0020\u0073i\u0067n\u0065\u0064\u000a")
	} else {
		_egacb.WriteString("\u0053\u0069\u0067\u006eed\u003a\u0020\u004e\u006f\u0074\u0020\u0073\u0069\u0067\u006e\u0065\u0064\u000a")
	}
	if _eafa.IsVerified {
		_egacb.WriteString("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0076\u0061\u006c\u0069\u0064\u0061t\u0069\u006f\u006e\u003a\u0020\u0049\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u000a")
	} else {
		_egacb.WriteString("\u0053\u0069\u0067\u006e\u0061\u0074u\u0072\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e:\u0020\u0049\u0073\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u000a")
	}
	if _eafa.IsTrusted {
		_egacb.WriteString("\u0054\u0072\u0075\u0073\u0074\u0065\u0064\u003a\u0020\u0043\u0065\u0072\u0074\u0069\u0066i\u0063a\u0074\u0065\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0073\u0074\u0065\u0064\u000a")
	} else {
		_egacb.WriteString("\u0054\u0072\u0075s\u0074\u0065\u0064\u003a \u0055\u006e\u0074\u0072\u0075\u0073\u0074e\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u000a")
	}
	if !_eafa.GeneralizedTime.IsZero() {
		_egacb.WriteString(_b.Sprintf("G\u0065n\u0065\u0072\u0061\u006c\u0069\u007a\u0065\u0064T\u0069\u006d\u0065\u003a %\u0073\u000a", _eafa.GeneralizedTime.String()))
	}
	if _eafa.DiffResults != nil {
		_egacb.WriteString(_b.Sprintf("\u0064\u0069\u0066\u0066 i\u0073\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u003a\u0020\u0025v\u000a", _eafa.DiffResults.IsPermitted()))
		if len(_eafa.DiffResults.Warnings) > 0 {
			_egacb.WriteString("\u004d\u0044\u0050\u0020\u0077\u0061\u0072\u006e\u0069n\u0067\u0073\u003a\u000a")
			for _, _gafd := range _eafa.DiffResults.Warnings {
				_egacb.WriteString(_b.Sprintf("\u0009\u0025\u0073\u000a", _gafd))
			}
		}
		if len(_eafa.DiffResults.Errors) > 0 {
			_egacb.WriteString("\u004d\u0044\u0050 \u0065\u0072\u0072\u006f\u0072\u0073\u003a\u000a")
			for _, _fefc := range _eafa.DiffResults.Errors {
				_egacb.WriteString(_b.Sprintf("\u0009\u0025\u0073\u000a", _fefc))
			}
		}
	}
	if _eafa.IsCrlFound {
		_egacb.WriteString("R\u0065\u0076\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0061\u0074\u0061\u003a\u0020C\u0052\u004c\u0020f\u006fu\u006e\u0064\u000a")
	} else {
		_egacb.WriteString("\u0052\u0065\u0076o\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0061\u0074\u0061\u003a \u0043\u0052\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u000a")
	}
	if _eafa.IsOcspFound {
		_egacb.WriteString("\u0052\u0065\u0076\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0061\u0074\u0061:\u0020O\u0043\u0053\u0050\u0020\u0066\u006f\u0075n\u0064\u000a")
	} else {
		_egacb.WriteString("\u0052\u0065\u0076\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0061\u0074\u0061:\u0020O\u0043\u0053\u0050\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u000a")
	}
	return _egacb.String()
}

// NewPdfActionThread returns a new "thread" action.
func NewPdfActionThread() *PdfActionThread {
	_ebcc := NewPdfAction()
	_gfdc := &PdfActionThread{}
	_gfdc.PdfAction = _ebcc
	_ebcc.SetContext(_gfdc)
	return _gfdc
}

// GetContainingPdfObject returns the containing object for the PdfField, i.e. an indirect object
// containing the field dictionary.
func (_defd *PdfField) GetContainingPdfObject() _dg.PdfObject { return _defd._egce }

// ToPdfObject implements interface PdfModel.
func (_gdgf *PdfAnnotationLink) ToPdfObject() _dg.PdfObject {
	_gdgf.PdfAnnotation.ToPdfObject()
	_baca := _gdgf._cdf
	_bfdg := _baca.PdfObject.(*_dg.PdfObjectDictionary)
	_bfdg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u004c\u0069\u006e\u006b"))
	if _gdgf._ece != nil && _gdgf._ece._bg != nil {
		_bfdg.Set("\u0041", _gdgf._ece._bg.ToPdfObject())
	} else if _gdgf.A != nil {
		_bfdg.Set("\u0041", _gdgf.A)
	}
	_bfdg.SetIfNotNil("\u0044\u0065\u0073\u0074", _gdgf.Dest)
	_bfdg.SetIfNotNil("\u0048", _gdgf.H)
	_bfdg.SetIfNotNil("\u0050\u0041", _gdgf.PA)
	_bfdg.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _gdgf.QuadPoints)
	_bfdg.SetIfNotNil("\u0042\u0053", _gdgf.BS)
	return _baca
}

// FullName returns the full name of the field as in rootname.parentname.partialname.
func (_gebg *PdfField) FullName() (string, error) {
	var _aaccb _bc.Buffer
	_efebd := []string{}
	if _gebg.T != nil {
		_efebd = append(_efebd, _gebg.T.Decoded())
	}
	_dcff := map[*PdfField]bool{}
	_dcff[_gebg] = true
	_adga := _gebg.Parent
	for _adga != nil {
		if _, _ffde := _dcff[_adga]; _ffde {
			return _aaccb.String(), _bf.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		if _adga.T == nil {
			return _aaccb.String(), _bf.New("\u0066\u0069el\u0064\u0020\u0070a\u0072\u0074\u0069\u0061l n\u0061me\u0020\u0028\u0054\u0029\u0020\u006e\u006ft \u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064")
		}
		_efebd = append(_efebd, _adga.T.Decoded())
		_dcff[_adga] = true
		_adga = _adga.Parent
	}
	for _ecedf := len(_efebd) - 1; _ecedf >= 0; _ecedf-- {
		_aaccb.WriteString(_efebd[_ecedf])
		if _ecedf > 0 {
			_aaccb.WriteString("\u002e")
		}
	}
	return _aaccb.String(), nil
}
func (_fbfd *PdfAppender) updateObjectsDeep(_bebc _dg.PdfObject, _bagd map[_dg.PdfObject]struct{}) {
	if _bagd == nil {
		_bagd = map[_dg.PdfObject]struct{}{}
	}
	if _, _ggc := _bagd[_bebc]; _ggc || _bebc == nil {
		return
	}
	_bagd[_bebc] = struct{}{}
	_ebcg := _dg.ResolveReferencesDeep(_bebc, _fbfd._ebaa)
	if _ebcg != nil {
		_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ebcg)
	}
	switch _aea := _bebc.(type) {
	case *_dg.PdfIndirectObject:
		switch {
		case _aea.GetParser() == _fbfd._debg._baad:
			return
		case _aea.GetParser() == _fbfd.Reader._baad:
			_afgc, _ := _fbfd._debg.GetIndirectObjectByNumber(int(_aea.ObjectNumber))
			_aaea, _acba := _afgc.(*_dg.PdfIndirectObject)
			if _acba && _aaea != nil {
				if _aaea.PdfObject != _aea.PdfObject && _aaea.PdfObject.WriteString() != _aea.PdfObject.WriteString() {
					if _ga.Contains(_aea.PdfObject.WriteString(), "\u002f\u0053\u0069\u0067") && _ga.Contains(_aea.PdfObject.WriteString(), "\u002f\u0053\u0075\u0062\u0074\u0079\u0070\u0065") {
						return
					}
					_fbfd.addNewObject(_bebc)
					_fbfd._gebe[_bebc] = _aea.ObjectNumber
				}
			}
		default:
			_fbfd.addNewObject(_bebc)
		}
		_fbfd.updateObjectsDeep(_aea.PdfObject, _bagd)
	case *_dg.PdfObjectArray:
		for _, _ggcg := range _aea.Elements() {
			_fbfd.updateObjectsDeep(_ggcg, _bagd)
		}
	case *_dg.PdfObjectDictionary:
		for _, _bfe := range _aea.Keys() {
			_fbfd.updateObjectsDeep(_aea.Get(_bfe), _bagd)
		}
	case *_dg.PdfObjectStreams:
		if _aea.GetParser() != _fbfd._debg._baad {
			for _, _bcfeb := range _aea.Elements() {
				_fbfd.updateObjectsDeep(_bcfeb, _bagd)
			}
		}
	case *_dg.PdfObjectStream:
		switch {
		case _aea.GetParser() == _fbfd._debg._baad:
			return
		case _aea.GetParser() == _fbfd.Reader._baad:
			if _bgcc, _gdag := _fbfd._debg._baad.LookupByReference(_aea.PdfObjectReference); _gdag == nil {
				var _agdg bool
				if _aeac, _fccee := _dg.GetStream(_bgcc); _fccee && _bc.Equal(_aeac.Stream, _aea.Stream) {
					_agdg = true
				}
				if _cabg, _ebb := _dg.GetDict(_bgcc); _agdg && _ebb {
					_agdg = _cabg.WriteString() == _aea.PdfObjectDictionary.WriteString()
				}
				if _agdg {
					return
				}
			}
			if _aea.ObjectNumber != 0 {
				_fbfd._gebe[_bebc] = _aea.ObjectNumber
			}
		default:
			if _, _dcadf := _fbfd._efa[_bebc]; !_dcadf {
				_fbfd.addNewObject(_bebc)
			}
		}
		_fbfd.updateObjectsDeep(_aea.PdfObjectDictionary, _bagd)
	}
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the reader using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
// At the end of the process, the AcroForm contains all the fields which were
// not flattened. If all fields are flattened, the reader's AcroForm field
// is set to nil.
func (_ebfd *PdfReader) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	return _ebfd.flattenFieldsWithOpts(false, appgen, opts)
}

// ToPdfObject implements interface PdfModel.
func (_fdbf *PdfAnnotationPopup) ToPdfObject() _dg.PdfObject {
	_fdbf.PdfAnnotation.ToPdfObject()
	_cfda := _fdbf._cdf
	_fgdd := _cfda.PdfObject.(*_dg.PdfObjectDictionary)
	_fgdd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0050\u006f\u0070u\u0070"))
	_fgdd.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _fdbf.Parent)
	_fgdd.SetIfNotNil("\u004f\u0070\u0065\u006e", _fdbf.Open)
	return _cfda
}

// ImageToRGB converts ICCBased colorspace image to RGB and returns the result.
func (_ebdg *PdfColorspaceICCBased) ImageToRGB(img Image) (Image, error) {
	if _ebdg.Alternate == nil {
		_ag.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _ebdg.N == 1 {
			_ag.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_efff := NewPdfColorspaceDeviceGray()
			return _efff.ImageToRGB(img)
		} else if _ebdg.N == 3 {
			_ag.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return img, nil
		} else if _ebdg.N == 4 {
			_ag.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_bgade := NewPdfColorspaceDeviceCMYK()
			return _bgade.ImageToRGB(img)
		} else {
			return img, _bf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_ag.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _ebdg)
	_ggaf, _decd := _ebdg.Alternate.ImageToRGB(img)
	_ag.Log.Trace("I\u0043C\u0020\u0049\u006e\u0070\u0075\u0074\u0020\u0069m\u0061\u0067\u0065\u003a %\u002b\u0076", img)
	_ag.Log.Trace("I\u0043\u0043\u0020\u004fut\u0070u\u0074\u0020\u0069\u006d\u0061g\u0065\u003a\u0020\u0025\u002b\u0076", _ggaf)
	return _ggaf, _decd
}
func _ffgc(_aggca _dg.PdfObject) (*PdfBorderStyle, error) {
	_bgbg := &PdfBorderStyle{}
	_bgbg._cbcb = _aggca
	var _gcfee *_dg.PdfObjectDictionary
	_aggca = _dg.TraceToDirectObject(_aggca)
	_gcfee, _dfb := _aggca.(*_dg.PdfObjectDictionary)
	if !_dfb {
		return nil, _bf.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	if _egae := _gcfee.Get("\u0054\u0079\u0070\u0065"); _egae != nil {
		_dffc, _eged := _egae.(*_dg.PdfObjectName)
		if !_eged {
			_ag.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u0077\u0069th\u0020\u0054\u0079\u0070e\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061me\u0020\u006fb\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0054", _egae)
		} else {
			if *_dffc != "\u0042\u006f\u0072\u0064\u0065\u0072" {
				_ag.Log.Debug("W\u0061\u0072\u006e\u0069\u006e\u0067,\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020B\u006f\u0072\u0064e\u0072:\u0020\u0025\u0073", *_dffc)
			}
		}
	}
	if _abdg := _gcfee.Get("\u0057"); _abdg != nil {
		_gbfd, _cada := _dg.GetNumberAsFloat(_abdg)
		if _cada != nil {
			_ag.Log.Debug("\u0045\u0072\u0072\u006fr \u0072\u0065\u0074\u0072\u0069\u0065\u0076\u0069\u006e\u0067\u0020\u0057\u003a\u0020%\u0076", _cada)
			return nil, _cada
		}
		_bgbg.W = &_gbfd
	}
	if _fedf := _gcfee.Get("\u0053"); _fedf != nil {
		_fede, _egga := _fedf.(*_dg.PdfObjectName)
		if !_egga {
			return nil, _bf.New("\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u0053\u0020\u006e\u006ft\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		var _fbdd BorderStyle
		switch *_fede {
		case "\u0053":
			_fbdd = BorderStyleSolid
		case "\u0044":
			_fbdd = BorderStyleDashed
		case "\u0042":
			_fbdd = BorderStyleBeveled
		case "\u0049":
			_fbdd = BorderStyleInset
		case "\u0055":
			_fbdd = BorderStyleUnderline
		default:
			_ag.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020s\u0074\u0079\u006c\u0065\u0020n\u0061\u006d\u0065\u0020\u0025\u0073", *_fede)
			return nil, _bf.New("\u0073\u0074\u0079\u006ce \u0074\u0079\u0070\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065c\u006b")
		}
		_bgbg.S = &_fbdd
	}
	if _aggf := _gcfee.Get("\u0044"); _aggf != nil {
		_becg, _acab := _aggf.(*_dg.PdfObjectArray)
		if !_acab {
			_ag.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044\u0020\u0064a\u0073\u0068\u0020\u006e\u006f\u0074\u0020a\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0054", _aggf)
			return nil, _bf.New("\u0062o\u0072\u0064\u0065\u0072 \u0044\u0020\u0074\u0079\u0070e\u0020c\u0068e\u0063\u006b\u0020\u0065\u0072\u0072\u006fr")
		}
		_fegg, _fbgb := _becg.ToIntegerArray()
		if _fbgb != nil {
			_ag.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044 \u0050\u0072\u006fbl\u0065\u006d\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u006e\u0074\u0065\u0067e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u003a \u0025\u0076", _fbgb)
			return nil, _fbgb
		}
		_bgbg.D = &_fegg
	}
	return _bgbg, nil
}
func _ecef(_fggbf _dg.PdfObject) (*PdfColorspaceDeviceNAttributes, error) {
	_ggfca := &PdfColorspaceDeviceNAttributes{}
	var _cbfec *_dg.PdfObjectDictionary
	switch _bcfb := _fggbf.(type) {
	case *_dg.PdfIndirectObject:
		_ggfca._ffbg = _bcfb
		var _ffdg bool
		_cbfec, _ffdg = _bcfb.PdfObject.(*_dg.PdfObjectDictionary)
		if !_ffdg {
			_ag.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
			return nil, _bf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	case *_dg.PdfObjectDictionary:
		_cbfec = _bcfb
	case *_dg.PdfObjectReference:
		_agccb := _bcfb.Resolve()
		return _ecef(_agccb)
	default:
		_ag.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
		return nil, _bf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _bdgag := _cbfec.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _bdgag != nil {
		_acad, _eacdc := _dg.TraceToDirectObject(_bdgag).(*_dg.PdfObjectName)
		if !_eacdc {
			_ag.Log.Error("\u0044\u0065vi\u0063\u0065\u004e \u0061\u0074\u0074\u0072ibu\u0074e \u0053\u0075\u0062\u0074\u0079\u0070\u0065 t\u0079\u0070\u0065\u0020\u0065\u0072\u0072o\u0072")
			return nil, _bf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ggfca.Subtype = _acad
	}
	if _dggg := _cbfec.Get("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s"); _dggg != nil {
		_ggfca.Colorants = _dggg
	}
	if _deee := _cbfec.Get("\u0050r\u006f\u0063\u0065\u0073\u0073"); _deee != nil {
		_ggfca.Process = _deee
	}
	if _ecbaba := _cbfec.Get("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073"); _ecbaba != nil {
		_ggfca.MixingHints = _ecbaba
	}
	return _ggfca, nil
}

// SetDecode sets the decode image float slice.
func (_caaeca *Image) SetDecode(decode []float64) { _caaeca._gfbb = decode }

// AddExtension adds the specified extension to the Extensions dictionary.
// See section 7.1.2 "Extensions Dictionary" (pp. 108-109 PDF32000_2008).
func (_eccdb *PdfWriter) AddExtension(extName, baseVersion string, extLevel int) {
	_gaecd, _bgbdg := _dg.GetDict(_eccdb._ecdf.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073"))
	if !_bgbdg {
		_gaecd = _dg.MakeDict()
		_eccdb._ecdf.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073", _gaecd)
	}
	_badda, _bgbdg := _dg.GetDict(_gaecd.Get(_dg.PdfObjectName(extName)))
	if !_bgbdg {
		_badda = _dg.MakeDict()
		_gaecd.Set(_dg.PdfObjectName(extName), _badda)
	}
	if _gedfe, _ := _dg.GetNameVal(_badda.Get("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e")); _gedfe != baseVersion {
		_badda.Set("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e", _dg.MakeName(baseVersion))
	}
	if _bbfdd, _ := _dg.GetIntVal(_badda.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c")); _bbfdd != extLevel {
		_badda.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c", _dg.MakeInteger(int64(extLevel)))
	}
}
func _bccf(_beb _dg.PdfObject) (*PdfFilespec, error) {
	if _beb == nil {
		return nil, nil
	}
	return NewPdfFilespecFromObj(_beb)
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the A, B and C components of the color.
func (_eaegc *PdfColorspaceCalRGB) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_beba, _ggddc := _dg.GetNumbersAsFloat(objects)
	if _ggddc != nil {
		return nil, _ggddc
	}
	return _eaegc.ColorFromFloats(_beba)
}

// ColorToRGB converts a CalGray color to an RGB color.
func (_acea *PdfColorspaceCalGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_adcc, _eacd := color.(*PdfColorCalGray)
	if !_eacd {
		_ag.Log.Debug("\u0049n\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006eo\u0074\u0020\u0063\u0061\u006c\u0020\u0067\u0072\u0061\u0079")
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	ANorm := _adcc.Val()
	X := _acea.WhitePoint[0] * _cg.Pow(ANorm, _acea.Gamma)
	Y := _acea.WhitePoint[1] * _cg.Pow(ANorm, _acea.Gamma)
	Z := _acea.WhitePoint[2] * _cg.Pow(ANorm, _acea.Gamma)
	_gdda := 3.240479*X + -1.537150*Y + -0.498535*Z
	_cafg := -0.969256*X + 1.875992*Y + 0.041556*Z
	_fgfed := 0.055648*X + -0.204043*Y + 1.057311*Z
	_gdda = _cg.Min(_cg.Max(_gdda, 0), 1.0)
	_cafg = _cg.Min(_cg.Max(_cafg, 0), 1.0)
	_fgfed = _cg.Min(_cg.Max(_fgfed, 0), 1.0)
	return NewPdfColorDeviceRGB(_gdda, _cafg, _fgfed), nil
}

// PdfColorPattern represents a pattern color.
type PdfColorPattern struct {
	Color       PdfColor
	PatternName _dg.PdfObjectName
}

func (_cce *PdfReader) loadAction(_beadc _dg.PdfObject) (*PdfAction, error) {
	if _aagg, _aefd := _dg.GetIndirect(_beadc); _aefd {
		_gdgc, _beadg := _cce.newPdfActionFromIndirectObject(_aagg)
		if _beadg != nil {
			return nil, _beadg
		}
		return _gdgc, nil
	} else if !_dg.IsNullObject(_beadc) {
		return nil, _bf.New("\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074 \u0074\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, nil
}
func (_edea *PdfReader) newPdfActionTransFromDict(_cag *_dg.PdfObjectDictionary) (*PdfActionTrans, error) {
	return &PdfActionTrans{Trans: _cag.Get("\u0054\u0072\u0061n\u0073")}, nil
}

// NewPdfColorspaceDeviceRGB returns a new RGB colorspace object.
func NewPdfColorspaceDeviceRGB() *PdfColorspaceDeviceRGB { return &PdfColorspaceDeviceRGB{} }

var _decga = _c.MustCompile("\u005b\\\u006e\u005c\u0072\u005d\u002b")

func (_agdd *PdfReader) newPdfAnnotationPrinterMarkFromDict(_bcge *_dg.PdfObjectDictionary) (*PdfAnnotationPrinterMark, error) {
	_bcfgb := PdfAnnotationPrinterMark{}
	_bcfgb.MN = _bcge.Get("\u004d\u004e")
	return &_bcfgb, nil
}

// String returns a string describing the font descriptor.
func (_adddb *PdfFontDescriptor) String() string {
	var _cge []string
	if _adddb.FontName != nil {
		_cge = append(_cge, _adddb.FontName.String())
	}
	if _adddb.FontFamily != nil {
		_cge = append(_cge, _adddb.FontFamily.String())
	}
	if _adddb.fontFile != nil {
		_cge = append(_cge, _adddb.fontFile.String())
	}
	if _adddb._gbcg != nil {
		_cge = append(_cge, _adddb._gbcg.String())
	}
	_cge = append(_cge, _b.Sprintf("\u0046\u006f\u006et\u0046\u0069\u006c\u0065\u0033\u003d\u0025\u0074", _adddb.FontFile3 != nil))
	return _b.Sprintf("\u0046\u004f\u004e\u0054_D\u0045\u0053\u0043\u0052\u0049\u0050\u0054\u004f\u0052\u007b\u0025\u0073\u007d", _ga.Join(_cge, "\u002c\u0020"))
}
func (_dbee *PdfAppender) addNewObject(_dbca _dg.PdfObject) {
	if _, _dbcb := _dbee._efa[_dbca]; !_dbcb {
		_dbee._dgfd = append(_dbee._dgfd, _dbca)
		_dbee._efa[_dbca] = struct{}{}
	}
}

// ToPdfObject implements interface PdfModel.
func (_bgf *PdfActionLaunch) ToPdfObject() _dg.PdfObject {
	_bgf.PdfAction.ToPdfObject()
	_ecba := _bgf._cbd
	_adb := _ecba.PdfObject.(*_dg.PdfObjectDictionary)
	_adb.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeLaunch)))
	if _bgf.F != nil {
		_adb.Set("\u0046", _bgf.F.ToPdfObject())
	}
	_adb.SetIfNotNil("\u0057\u0069\u006e", _bgf.Win)
	_adb.SetIfNotNil("\u004d\u0061\u0063", _bgf.Mac)
	_adb.SetIfNotNil("\u0055\u006e\u0069\u0078", _bgf.Unix)
	_adb.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _bgf.NewWindow)
	return _ecba
}

// FieldAppearanceGenerator generates appearance stream for a given field.
type FieldAppearanceGenerator interface {
	ContentStreamWrapper
	GenerateAppearanceDict(_aegf *PdfAcroForm, _bafcea *PdfField, _faed *PdfAnnotationWidget) (*_dg.PdfObjectDictionary, error)
}

// ToPdfObject implements interface PdfModel.
func (_cbfbc *PdfAnnotationProjection) ToPdfObject() _dg.PdfObject {
	_cbfbc.PdfAnnotation.ToPdfObject()
	_gdfdb := _cbfbc._cdf
	_dcag := _gdfdb.PdfObject.(*_dg.PdfObjectDictionary)
	_cbfbc.PdfAnnotationMarkup.appendToPdfDictionary(_dcag)
	return _gdfdb
}
func (_dgaba PdfFont) actualFont() pdfFont {
	if _dgaba._cadf == nil {
		_ag.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0061\u0063\u0074\u0075\u0061\u006c\u0046\u006f\u006e\u0074\u002e\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c.\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _dgaba)
	}
	return _dgaba._cadf
}

// GetNumComponents returns the number of color components (3 for CalRGB).
func (_bafce *PdfColorCalRGB) GetNumComponents() int { return 3 }

// SetNameDictionary sets the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_gafca *PdfWriter) SetNameDictionary(names _dg.PdfObject) error {
	if names == nil {
		return nil
	}
	_ag.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u004e\u0061\u006d\u0065\u0073\u002e\u002e\u002e")
	_gafca._ecdf.Set("\u004e\u0061\u006de\u0073", names)
	return _gafca.addObjects(names)
}

// PdfBorderStyle represents a border style dictionary (12.5.4 Border Styles p. 394).
type PdfBorderStyle struct {
	W     *float64
	S     *BorderStyle
	D     *[]int
	_cbcb _dg.PdfObject
}

// A returns the value of the A component of the color.
func (_bfab *PdfColorCalRGB) A() float64 { return _bfab[0] }
func (_daba *PdfReader) newPdfAnnotationSquareFromDict(_gba *_dg.PdfObjectDictionary) (*PdfAnnotationSquare, error) {
	_dce := PdfAnnotationSquare{}
	_dfca, _bdcf := _daba.newPdfAnnotationMarkupFromDict(_gba)
	if _bdcf != nil {
		return nil, _bdcf
	}
	_dce.PdfAnnotationMarkup = _dfca
	_dce.BS = _gba.Get("\u0042\u0053")
	_dce.IC = _gba.Get("\u0049\u0043")
	_dce.BE = _gba.Get("\u0042\u0045")
	_dce.RD = _gba.Get("\u0052\u0044")
	return &_dce, nil
}

// PdfVersion returns version of the PDF file.
func (_bcgbf *PdfReader) PdfVersion() _dg.Version { return _bcgbf._baad.PdfVersion() }
func (_adbd *PdfReader) newPdfActionNamedFromDict(_gagd *_dg.PdfObjectDictionary) (*PdfActionNamed, error) {
	return &PdfActionNamed{N: _gagd.Get("\u004e")}, nil
}

// PdfAnnotationCaret represents Caret annotations.
// (Section 12.5.6.11).
type PdfAnnotationCaret struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	RD _dg.PdfObject
	Sy _dg.PdfObject
}

// PageProcessCallback callback function used in page loading
// that could be used to modify the page content.
//
// If an error is returned, the `ToWriter` process would fail.
//
// This callback, if defined, will take precedence over `PageCallback` callback.
type PageProcessCallback func(_fgefa int, _adee *PdfPage) error

// Encoder returns the font's text encoder.
func (_faga pdfFontType0) Encoder() _bd.TextEncoder { return _faga._ggec }
func (_abad *PdfAcroForm) fill(_geagcb FieldValueProvider, _dcaca FieldAppearanceGenerator) error {
	if _abad == nil {
		return nil
	}
	_bdce, _dgcg := _geagcb.FieldValues()
	if _dgcg != nil {
		return _dgcg
	}
	for _, _fdee := range _abad.AllFields() {
		_aeag := _fdee.PartialName()
		_bdegd, _badea := _bdce[_aeag]
		if !_badea {
			if _gbaac, _efbgb := _fdee.FullName(); _efbgb == nil {
				_bdegd, _badea = _bdce[_gbaac]
			}
		}
		if !_badea {
			_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _aeag)
			continue
		}
		if _egfe := _ebbce(_fdee, _bdegd); _egfe != nil {
			return _egfe
		}
		if _dcaca == nil {
			continue
		}
		for _, _aafg := range _fdee.Annotations {
			_gdaf, _debgf := _dcaca.GenerateAppearanceDict(_abad, _fdee, _aafg)
			if _debgf != nil {
				return _debgf
			}
			_aafg.AP = _gdaf
			_aafg.ToPdfObject()
		}
	}
	return nil
}

// ToPdfOutline returns a low level PdfOutline object, based on the current
// instance.
func (_gdfde *Outline) ToPdfOutline() *PdfOutline {
	_fecab := NewPdfOutline()
	var _fcege []*PdfOutlineItem
	var _bbbba int64
	var _fdbba *PdfOutlineItem
	for _, _ebfb := range _gdfde.Entries {
		_aeec, _bccgg := _ebfb.ToPdfOutlineItem()
		_aeec.Parent = &_fecab.PdfOutlineTreeNode
		if _fdbba != nil {
			_fdbba.Next = &_aeec.PdfOutlineTreeNode
			_aeec.Prev = &_fdbba.PdfOutlineTreeNode
		}
		_fcege = append(_fcege, _aeec)
		_bbbba += _bccgg
		_fdbba = _aeec
	}
	_eddbb := int64(len(_fcege))
	_bbbba += _eddbb
	if _eddbb > 0 {
		_fecab.First = &_fcege[0].PdfOutlineTreeNode
		_fecab.Last = &_fcege[_eddbb-1].PdfOutlineTreeNode
		_fecab.Count = &_bbbba
	}
	return _fecab
}

// ToGray returns a PdfColorDeviceGray color based on the current RGB color.
func (_gcdb *PdfColorDeviceRGB) ToGray() *PdfColorDeviceGray {
	_efbea := 0.3*_gcdb.R() + 0.59*_gcdb.G() + 0.11*_gcdb.B()
	_efbea = _cg.Min(_cg.Max(_efbea, 0.0), 1.0)
	return NewPdfColorDeviceGray(_efbea)
}
func _ccgad(_caaec _dg.PdfObject) (*PdfFunctionType3, error) {
	_ffeg := &PdfFunctionType3{}
	var _baeeg *_dg.PdfObjectDictionary
	if _cfgfc, _facf := _caaec.(*_dg.PdfIndirectObject); _facf {
		_gggf, _ababf := _cfgfc.PdfObject.(*_dg.PdfObjectDictionary)
		if !_ababf {
			return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ffeg._defdg = _cfgfc
		_baeeg = _gggf
	} else if _bafd, _gfadd := _caaec.(*_dg.PdfObjectDictionary); _gfadd {
		_baeeg = _bafd
	} else {
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_eedee, _bdaff := _dg.TraceToDirectObject(_baeeg.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_dg.PdfObjectArray)
	if !_bdaff {
		_ag.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _eedee.Len() != 2 {
		_ag.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _bf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_gabcf, _fddea := _eedee.ToFloat64Array()
	if _fddea != nil {
		return nil, _fddea
	}
	_ffeg.Domain = _gabcf
	_eedee, _bdaff = _dg.TraceToDirectObject(_baeeg.Get("\u0052\u0061\u006eg\u0065")).(*_dg.PdfObjectArray)
	if _bdaff {
		if _eedee.Len() < 0 || _eedee.Len()%2 != 0 {
			return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_edgdd, _gebaec := _eedee.ToFloat64Array()
		if _gebaec != nil {
			return nil, _gebaec
		}
		_ffeg.Range = _edgdd
	}
	_eedee, _bdaff = _dg.TraceToDirectObject(_baeeg.Get("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es")).(*_dg.PdfObjectArray)
	if !_bdaff {
		_ag.Log.Error("\u0046\u0075\u006ect\u0069\u006f\u006e\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_ffeg.Functions = []PdfFunction{}
	for _, _gaacb := range _eedee.Elements() {
		_fgaff, _ebgbb := _agec(_gaacb)
		if _ebgbb != nil {
			return nil, _ebgbb
		}
		_ffeg.Functions = append(_ffeg.Functions, _fgaff)
	}
	_eedee, _bdaff = _dg.TraceToDirectObject(_baeeg.Get("\u0042\u006f\u0075\u006e\u0064\u0073")).(*_dg.PdfObjectArray)
	if !_bdaff {
		_ag.Log.Error("B\u006fu\u006e\u0064\u0073\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_edbb, _fddea := _eedee.ToFloat64Array()
	if _fddea != nil {
		return nil, _fddea
	}
	_ffeg.Bounds = _edbb
	if len(_ffeg.Bounds) != len(_ffeg.Functions)-1 {
		_ag.Log.Error("B\u006f\u0075\u006e\u0064\u0073\u0020\u0028\u0025\u0064)\u0020\u0061\u006e\u0064\u0020\u006e\u0075m \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029 n\u006f\u0074 \u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067", len(_ffeg.Bounds), len(_ffeg.Functions))
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_eedee, _bdaff = _dg.TraceToDirectObject(_baeeg.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_dg.PdfObjectArray)
	if !_bdaff {
		_ag.Log.Error("E\u006ec\u006f\u0064\u0065\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_bbafg, _fddea := _eedee.ToFloat64Array()
	if _fddea != nil {
		return nil, _fddea
	}
	_ffeg.Encode = _bbafg
	if len(_ffeg.Encode) != 2*len(_ffeg.Functions) {
		_ag.Log.Error("\u004c\u0065\u006e\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0028\u0025\u0064\u0029 \u0061\u006e\u0064\u0020\u006e\u0075\u006d\u0020\u0066\u0075\u006e\u0063\u0074i\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029\u0020\u006e\u006f\u0074 m\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u0075\u0070", len(_ffeg.Encode), len(_ffeg.Functions))
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return _ffeg, nil
}

// ImageToRGB converts an image with samples in Separation CS to an image with samples specified in
// DeviceRGB CS.
func (_aedeg *PdfColorspaceSpecialSeparation) ImageToRGB(img Image) (Image, error) {
	_fbfa := _fcd.NewReader(img.getBase())
	_ggce := _fc.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), _aedeg.AlternateSpace.GetNumComponents(), nil, img._dgeb, nil)
	_dffcc := _fcd.NewWriter(_ggce)
	_ddcc := _cg.Pow(2, float64(img.BitsPerComponent)) - 1
	_ag.Log.Trace("\u0053\u0065\u0070a\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u002d\u003e\u0020\u0054\u006f\u0052\u0047\u0042\u0020\u0063o\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	_ag.Log.Trace("\u0054i\u006et\u0054\u0072\u0061\u006e\u0073f\u006f\u0072m\u003a\u0020\u0025\u002b\u0076", _aedeg.TintTransform)
	_bbfe := _aedeg.AlternateSpace.DecodeArray()
	var (
		_dded  uint32
		_bgadc error
	)
	for {
		_dded, _bgadc = _fbfa.ReadSample()
		if _bgadc == _cf.EOF {
			break
		}
		if _bgadc != nil {
			return img, _bgadc
		}
		_bafe := float64(_dded) / _ddcc
		_ffcb, _fage := _aedeg.TintTransform.Evaluate([]float64{_bafe})
		if _fage != nil {
			return img, _fage
		}
		for _bgdf, _bdcd := range _ffcb {
			_cfgc := _fc.LinearInterpolate(_bdcd, _bbfe[_bgdf*2], _bbfe[_bgdf*2+1], 0, 1)
			if _fage = _dffcc.WriteSample(uint32(_cfgc * _ddcc)); _fage != nil {
				return img, _fage
			}
		}
	}
	return _aedeg.AlternateSpace.ImageToRGB(_edcf(&_ggce))
}

// ToInteger convert to an integer format.
func (_dcae *PdfColorCalRGB) ToInteger(bits int) [3]uint32 {
	_cdgbb := _cg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_cdgbb * _dcae.A()), uint32(_cdgbb * _dcae.B()), uint32(_cdgbb * _dcae.C())}
}
func _cefdc(_eeac *_dg.PdfObjectDictionary, _bebae *fontCommon) (*pdfFontType0, error) {
	_addfe, _dfgbf := _dg.GetArray(_eeac.Get("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073"))
	if !_dfgbf {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006cid\u0020\u0044\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073\u0020\u002d\u0020\u006e\u006f\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079 \u0025\u0073", _bebae)
		return nil, _dg.ErrRangeError
	}
	if _addfe.Len() != 1 {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0041\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020(%\u0064\u0029", _addfe.Len())
		return nil, _dg.ErrRangeError
	}
	_eceff, _acda := _gegbd(_addfe.Get(0), false)
	if _acda != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046a\u0069\u006c\u0065d \u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u0020\u0066\u006f\u006et\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076 \u0025\u0073", _acda, _bebae)
		return nil, _acda
	}
	_gdcb := _cbbgd(_bebae)
	_gdcb.DescendantFont = _eceff
	_agfe, _dfgbf := _dg.GetNameVal(_eeac.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if _dfgbf {
		if _agfe == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048" || _agfe == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056" {
			_gdcb._ggec = _bd.NewIdentityTextEncoder(_agfe)
		} else if _ff.IsPredefinedCMap(_agfe) {
			_gdcb._bfae, _acda = _ff.LoadPredefinedCMap(_agfe)
			if _acda != nil {
				_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _agfe, _acda)
			}
		} else {
			_ag.Log.Debug("\u0055\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0063\u006da\u0070\u0020\u0025\u0071", _agfe)
		}
	}
	if _efaac := _eceff.baseFields()._ecfb; _efaac != nil {
		if _fdbbf := _efaac.Name(); _fdbbf == "\u0041d\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002d\u0055\u0043\u0053\u0032" || _fdbbf == "\u0041\u0064\u006f\u0062\u0065\u002d\u0047\u0042\u0031-\u0055\u0043\u0053\u0032" || _fdbbf == "\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031-\u0055\u0043\u0053\u0032" || _fdbbf == "\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006f\u0072\u0065\u0061\u0031-\u0055\u0043\u0053\u0032" {
			_gdcb._ggec = _bd.NewCMapEncoder(_agfe, _gdcb._bfae, _efaac)
		}
	}
	return _gdcb, nil
}

// SetLocation sets the `Location` field of the signature.
func (_daefd *PdfSignature) SetLocation(location string) { _daefd.Location = _dg.MakeString(location) }

// ToPdfObject implements model.PdfModel interface.
func (_ebcbd *PdfOutputIntent) ToPdfObject() _dg.PdfObject {
	if _ebcbd._cfaef == nil {
		_ebcbd._cfaef = _dg.MakeDict()
	}
	_ccfbe := _ebcbd._cfaef
	if _ebcbd.Type != "" {
		_ccfbe.Set("\u0054\u0079\u0070\u0065", _dg.MakeName(_ebcbd.Type))
	}
	_ccfbe.Set("\u0053", _dg.MakeName(_ebcbd.S.String()))
	if _ebcbd.OutputCondition != "" {
		_ccfbe.Set("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e", _dg.MakeString(_ebcbd.OutputCondition))
	}
	_ccfbe.Set("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r", _dg.MakeString(_ebcbd.OutputConditionIdentifier))
	_ccfbe.Set("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065", _dg.MakeString(_ebcbd.RegistryName))
	if _ebcbd.Info != "" {
		_ccfbe.Set("\u0049\u006e\u0066\u006f", _dg.MakeString(_ebcbd.Info))
	}
	if len(_ebcbd.DestOutputProfile) != 0 {
		_ggga, _fafbe := _dg.MakeStream(_ebcbd.DestOutputProfile, _dg.NewFlateEncoder())
		if _fafbe != nil {
			_ag.Log.Error("\u004d\u0061\u006b\u0065\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0044\u0065s\u0074\u004f\u0075\u0074\u0070\u0075t\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fafbe)
		}
		_ggga.PdfObjectDictionary.Set("\u004e", _dg.MakeInteger(int64(_ebcbd.ColorComponents)))
		_cgec := make([]float64, _ebcbd.ColorComponents*2)
		for _fdadf := 0; _fdadf < _ebcbd.ColorComponents*2; _fdadf++ {
			_edage := 0.0
			if _fdadf%2 != 0 {
				_edage = 1.0
			}
			_cgec[_fdadf] = _edage
		}
		_ggga.PdfObjectDictionary.Set("\u0052\u0061\u006eg\u0065", _dg.MakeArrayFromFloats(_cgec))
		_ccfbe.Set("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065", _ggga)
	}
	return _ccfbe
}
func (_gbb *PdfReader) newPdfActionSoundFromDict(_bae *_dg.PdfObjectDictionary) (*PdfActionSound, error) {
	return &PdfActionSound{Sound: _bae.Get("\u0053\u006f\u0075n\u0064"), Volume: _bae.Get("\u0056\u006f\u006c\u0075\u006d\u0065"), Synchronous: _bae.Get("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073"), Repeat: _bae.Get("\u0052\u0065\u0070\u0065\u0061\u0074"), Mix: _bae.Get("\u004d\u0069\u0078")}, nil
}

// WriteToFile writes the Appender output to file specified by path.
func (_gaef *PdfAppender) WriteToFile(outputPath string) error {
	_ffge, _fccec := _eb.Create(outputPath)
	if _fccec != nil {
		return _fccec
	}
	defer _ffge.Close()
	return _gaef.Write(_ffge)
}

// ToPdfObject converts colorspace to a PDF object. [/Indexed base hival lookup]
func (_gbgd *PdfColorspaceSpecialIndexed) ToPdfObject() _dg.PdfObject {
	_bdacd := _dg.MakeArray(_dg.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"))
	_bdacd.Append(_gbgd.Base.ToPdfObject())
	_bdacd.Append(_dg.MakeInteger(int64(_gbgd.HiVal)))
	_bdacd.Append(_gbgd.Lookup)
	if _gbgd._fffd != nil {
		_gbgd._fffd.PdfObject = _bdacd
		return _gbgd._fffd
	}
	return _bdacd
}

// PdfAnnotationSound represents Sound annotations.
// (Section 12.5.6.16).
type PdfAnnotationSound struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Sound _dg.PdfObject
	Name  _dg.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_aad *PdfActionGoToR) ToPdfObject() _dg.PdfObject {
	_aad.PdfAction.ToPdfObject()
	_gfb := _aad._cbd
	_ecg := _gfb.PdfObject.(*_dg.PdfObjectDictionary)
	_ecg.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeGoToR)))
	if _aad.F != nil {
		_ecg.Set("\u0046", _aad.F.ToPdfObject())
	}
	_ecg.SetIfNotNil("\u0044", _aad.D)
	_ecg.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _aad.NewWindow)
	return _gfb
}

// PdfColorPatternType3 represents a color shading pattern type 3 (Radial).
type PdfColorPatternType3 struct {
	Color       PdfColor
	PatternName _dg.PdfObjectName
}

// NewPdfAnnotationPolyLine returns a new polyline annotation.
func NewPdfAnnotationPolyLine() *PdfAnnotationPolyLine {
	_edef := NewPdfAnnotation()
	_cagg := &PdfAnnotationPolyLine{}
	_cagg.PdfAnnotation = _edef
	_cagg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_edef.SetContext(_cagg)
	return _cagg
}

// ToPdfObject implements interface PdfModel.
func (_dfgc *PdfFilespec) ToPdfObject() _dg.PdfObject {
	_fgag := _dfgc.getDict()
	_fgag.Clear()
	_fgag.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063"))
	_fgag.SetIfNotNil("\u0046\u0053", _dfgc.FS)
	_fgag.SetIfNotNil("\u0046", _dfgc.F)
	_fgag.SetIfNotNil("\u0055\u0046", _dfgc.UF)
	_fgag.SetIfNotNil("\u0044\u004f\u0053", _dfgc.DOS)
	_fgag.SetIfNotNil("\u004d\u0061\u0063", _dfgc.Mac)
	_fgag.SetIfNotNil("\u0055\u006e\u0069\u0078", _dfgc.Unix)
	_fgag.SetIfNotNil("\u0049\u0044", _dfgc.ID)
	_fgag.SetIfNotNil("\u0056", _dfgc.V)
	_fgag.SetIfNotNil("\u0045\u0046", _dfgc.EF)
	_fgag.SetIfNotNil("\u0052\u0046", _dfgc.RF)
	_fgag.SetIfNotNil("\u0044\u0065\u0073\u0063", _dfgc.Desc)
	_fgag.SetIfNotNil("\u0043\u0049", _dfgc.CI)
	return _dfgc._bcbe
}

// GetNumComponents returns the number of color components (1 for CalGray).
func (_ddfd *PdfColorCalGray) GetNumComponents() int { return 1 }
func (_debdc *PdfFunctionType0) processSamples() error {
	_gdfcg := _fcd.ResampleBytes(_debdc._bgdg, _debdc.BitsPerSample)
	_debdc._fecfc = _gdfcg
	return nil
}

// Outline represents a PDF outline dictionary (Table 152 - p. 376).
// Currently, the Outline object can only be used to construct PDF outlines.
type Outline struct {
	Entries []*OutlineItem `json:"entries,omitempty"`
}

func (_fbgbe *PdfAppender) replaceObject(_aagc, _gfee _dg.PdfObject) {
	switch _efgd := _aagc.(type) {
	case *_dg.PdfIndirectObject:
		_fbgbe._gebe[_gfee] = _efgd.ObjectNumber
	case *_dg.PdfObjectStream:
		_fbgbe._gebe[_gfee] = _efgd.ObjectNumber
	}
}

// NewPdfAnnotationProjection returns a new projection annotation.
func NewPdfAnnotationProjection() *PdfAnnotationProjection {
	_ecc := NewPdfAnnotation()
	_egf := &PdfAnnotationProjection{}
	_egf.PdfAnnotation = _ecc
	_egf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ecc.SetContext(_egf)
	return _egf
}
func (_efgb *PdfReader) newPdfAnnotationPolygonFromDict(_bgdbg *_dg.PdfObjectDictionary) (*PdfAnnotationPolygon, error) {
	_aed := PdfAnnotationPolygon{}
	_cfg, _bgec := _efgb.newPdfAnnotationMarkupFromDict(_bgdbg)
	if _bgec != nil {
		return nil, _bgec
	}
	_aed.PdfAnnotationMarkup = _cfg
	_aed.Vertices = _bgdbg.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_aed.LE = _bgdbg.Get("\u004c\u0045")
	_aed.BS = _bgdbg.Get("\u0042\u0053")
	_aed.IC = _bgdbg.Get("\u0049\u0043")
	_aed.BE = _bgdbg.Get("\u0042\u0045")
	_aed.IT = _bgdbg.Get("\u0049\u0054")
	_aed.Measure = _bgdbg.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_aed, nil
}
func (_eadeg *PdfWriter) checkCrossReferenceStream() bool {
	_debace := _eadeg._efacd.Major > 1 || (_eadeg._efacd.Major == 1 && _eadeg._efacd.Minor > 4)
	if _eadeg._eaacb != nil {
		_debace = *_eadeg._eaacb
	}
	return _debace
}

// Read reads an image and loads into a new Image object with an RGB
// colormap and 8 bits per component.
func (_dbeg DefaultImageHandler) Read(reader _cf.Reader) (*Image, error) {
	_gafac, _, _ebaaa := _gd.Decode(reader)
	if _ebaaa != nil {
		_ag.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _ebaaa)
		return nil, _ebaaa
	}
	return _dbeg.NewImageFromGoImage(_gafac)
}

// NewOutline returns a new outline instance.
func NewOutline() *Outline { return &Outline{} }
func _dcbd(_eegd _dg.PdfObject) (*PdfColorspaceLab, error) {
	_dbfg := NewPdfColorspaceLab()
	if _bfbg, _eeda := _eegd.(*_dg.PdfIndirectObject); _eeda {
		_dbfg._aadb = _bfbg
	}
	_eegd = _dg.TraceToDirectObject(_eegd)
	_ecbc, _cbbaf := _eegd.(*_dg.PdfObjectArray)
	if !_cbbaf {
		return nil, _b.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _ecbc.Len() != 2 {
		return nil, _b.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_eegd = _dg.TraceToDirectObject(_ecbc.Get(0))
	_cagfec, _cbbaf := _eegd.(*_dg.PdfObjectName)
	if !_cbbaf {
		return nil, _b.Errorf("\u006c\u0061\u0062\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if *_cagfec != "\u004c\u0061\u0062" {
		return nil, _b.Errorf("n\u006ft\u0020\u0061\u0020\u004c\u0061\u0062\u0020\u0063o\u006c\u006f\u0072\u0073pa\u0063\u0065")
	}
	_eegd = _dg.TraceToDirectObject(_ecbc.Get(1))
	_edgc, _cbbaf := _eegd.(*_dg.PdfObjectDictionary)
	if !_cbbaf {
		return nil, _b.Errorf("c\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006ev\u0061l\u0069\u0064")
	}
	_eegd = _edgc.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_eegd = _dg.TraceToDirectObject(_eegd)
	_eeae, _cbbaf := _eegd.(*_dg.PdfObjectArray)
	if !_cbbaf {
		return nil, _b.Errorf("\u004c\u0061\u0062\u0020In\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069n\u0074")
	}
	if _eeae.Len() != 3 {
		return nil, _b.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
	}
	_geec, _dgec := _eeae.GetAsFloat64Slice()
	if _dgec != nil {
		return nil, _dgec
	}
	_dbfg.WhitePoint = _geec
	_eegd = _edgc.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _eegd != nil {
		_eegd = _dg.TraceToDirectObject(_eegd)
		_edae, _dbcd := _eegd.(*_dg.PdfObjectArray)
		if !_dbcd {
			return nil, _b.Errorf("\u004c\u0061\u0062: \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
		}
		if _edae.Len() != 3 {
			return nil, _b.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
		}
		_bcfea, _bebac := _edae.GetAsFloat64Slice()
		if _bebac != nil {
			return nil, _bebac
		}
		_dbfg.BlackPoint = _bcfea
	}
	_eegd = _edgc.Get("\u0052\u0061\u006eg\u0065")
	if _eegd != nil {
		_eegd = _dg.TraceToDirectObject(_eegd)
		_dbccf, _bbaf := _eegd.(*_dg.PdfObjectArray)
		if !_bbaf {
			_ag.Log.Error("\u0052\u0061n\u0067\u0065\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
			return nil, _b.Errorf("\u004ca\u0062:\u0020\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if _dbccf.Len() != 4 {
			_ag.Log.Error("\u0052\u0061\u006e\u0067\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020e\u0072\u0072\u006f\u0072")
			return nil, _b.Errorf("\u004c\u0061b\u003a\u0020\u0052a\u006e\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_gbbc, _gecfa := _dbccf.GetAsFloat64Slice()
		if _gecfa != nil {
			return nil, _gecfa
		}
		_dbfg.Range = _gbbc
	}
	return _dbfg, nil
}

// EnableByName LTV enables the signature dictionary of the PDF AcroForm
// field identified the specified name. The signing certificate chain is
// extracted from the signature dictionary. Optionally, additional certificates
// can be specified through the `extraCerts` parameter. The LTV client attempts
// to build the certificate chain up to a trusted root by downloading any
// missing certificates.
func (_geacf *LTV) EnableByName(name string, extraCerts []*_bb.Certificate) error {
	_eggcf := _geacf._ebdgd._debg.AcroForm
	for _, _faeac := range _eggcf.AllFields() {
		_cagee, _ := _faeac.GetContext().(*PdfFieldSignature)
		if _cagee == nil {
			continue
		}
		if _cafce := _cagee.PartialName(); _cafce != name {
			continue
		}
		return _geacf.Enable(_cagee.V, extraCerts)
	}
	return nil
}

// RemovePage removes a page by number.
func (_ggac *PdfAppender) RemovePage(pageNum int) {
	_dgd := pageNum - 1
	_ggac._ggdd = append(_ggac._ggdd[0:_dgd], _ggac._ggdd[pageNum:]...)
}

// GetPage returns the PdfPage model for the specified page number.
func (_fadfb *PdfReader) GetPage(pageNumber int) (*PdfPage, error) {
	if _fadfb._baad.GetCrypter() != nil && !_fadfb._baad.IsAuthenticated() {
		return nil, _b.Errorf("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	if len(_fadfb._daddd) < pageNumber {
		return nil, _bf.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0028\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0074o\u006f\u0020\u0073\u0068\u006f\u0072\u0074\u0029")
	}
	_fcacb := pageNumber - 1
	if _fcacb < 0 {
		return nil, _b.Errorf("\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0069\u006e\u0067\u0020\u006d\u0075\u0073t\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0061\u0074\u0020\u0031")
	}
	_gcdaa := _fadfb.PageList[_fcacb]
	return _gcdaa, nil
}

// GetBorderWidth returns the border style's width.
func (_fdbec *PdfBorderStyle) GetBorderWidth() float64 {
	if _fdbec.W == nil {
		return 1
	}
	return *_fdbec.W
}

// PdfAnnotationUnderline represents Underline annotations.
// (Section 12.5.6.10).
type PdfAnnotationUnderline struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _dg.PdfObject
}

func (_ggbbf *PdfWriter) setCatalogVersion() {
	_ggbbf._ecdf.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _dg.MakeName(_b.Sprintf("\u0025\u0064\u002e%\u0064", _ggbbf._efacd.Major, _ggbbf._efacd.Minor)))
}

// SetContext sets the specific fielddata type, e.g. would be PdfFieldButton for a button field.
func (_gbddg *PdfField) SetContext(ctx PdfModel) { _gbddg._bdfg = ctx }
func (_fbae *PdfReader) newPdfAnnotationSquigglyFromDict(_gee *_dg.PdfObjectDictionary) (*PdfAnnotationSquiggly, error) {
	_gfdag := PdfAnnotationSquiggly{}
	_egab, _baec := _fbae.newPdfAnnotationMarkupFromDict(_gee)
	if _baec != nil {
		return nil, _baec
	}
	_gfdag.PdfAnnotationMarkup = _egab
	_gfdag.QuadPoints = _gee.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_gfdag, nil
}

// NewOutlineItem returns a new outline item instance.
func NewOutlineItem(title string, dest OutlineDest) *OutlineItem {
	return &OutlineItem{Title: title, Dest: dest}
}

// GetCatalogMarkInfo gets catalog MarkInfo object.
func (_adbbe *PdfReader) GetCatalogMarkInfo() (_dg.PdfObject, bool) {
	if _adbbe._gccfb == nil {
		return nil, false
	}
	_gdafag := _adbbe._gccfb.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
	return _gdafag, _gdafag != nil
}

// PdfActionMovie represents a movie action.
type PdfActionMovie struct {
	*PdfAction
	Annotation _dg.PdfObject
	T          _dg.PdfObject
	Operation  _dg.PdfObject
}

func _bgdgf(_bced _dg.PdfObject) (*PdfShading, error) {
	_ddbgg := &PdfShading{}
	var _fdfdb *_dg.PdfObjectDictionary
	if _fada, _bcfda := _dg.GetIndirect(_bced); _bcfda {
		_ddbgg._bcfbg = _fada
		_dgea, _bagcc := _fada.PdfObject.(*_dg.PdfObjectDictionary)
		if !_bagcc {
			_ag.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _dg.ErrTypeError
		}
		_fdfdb = _dgea
	} else if _aceb, _affa := _dg.GetStream(_bced); _affa {
		_ddbgg._bcfbg = _aceb
		_fdfdb = _aceb.PdfObjectDictionary
	} else if _becbef, _agaa := _dg.GetDict(_bced); _agaa {
		_ddbgg._bcfbg = _becbef
		_fdfdb = _becbef
	} else {
		_ag.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", _bced)
		return nil, _dg.ErrTypeError
	}
	if _fdfdb == nil {
		_ag.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _bf.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_bced = _fdfdb.Get("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _bced == nil {
		_ag.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073si\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bced = _dg.TraceToDirectObject(_bced)
	_cdba, _gfbdc := _bced.(*_dg.PdfObjectInteger)
	if !_gfbdc {
		_ag.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066o\u0072 \u0073h\u0061d\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _bced)
		return nil, _dg.ErrTypeError
	}
	if *_cdba < 1 || *_cdba > 7 {
		_ag.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u006e\u006ft\u0020\u0031\u002d\u0037\u0020(\u0067\u006ft\u0020\u0025\u0064\u0029", *_cdba)
		return nil, _dg.ErrTypeError
	}
	_ddbgg.ShadingType = _cdba
	_bced = _fdfdb.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")
	if _bced == nil {
		_ag.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0020\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065\u0020e\u006e\u0074\u0072\u0079\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_cbgfb, _cagef := NewPdfColorspaceFromPdfObject(_bced)
	if _cagef != nil {
		_ag.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065: \u0025\u0076", _cagef)
		return nil, _cagef
	}
	_ddbgg.ColorSpace = _cbgfb
	_bced = _fdfdb.Get("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064")
	if _bced != nil {
		_bced = _dg.TraceToDirectObject(_bced)
		_bgcfd, _ebdbb := _bced.(*_dg.PdfObjectArray)
		if !_ebdbb {
			_ag.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _bced)
			return nil, _dg.ErrTypeError
		}
		_ddbgg.Background = _bgcfd
	}
	_bced = _fdfdb.Get("\u0042\u0042\u006f\u0078")
	if _bced != nil {
		_bced = _dg.TraceToDirectObject(_bced)
		_cdfdg, _dfcba := _bced.(*_dg.PdfObjectArray)
		if !_dfcba {
			_ag.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _bced)
			return nil, _dg.ErrTypeError
		}
		_bcffb, _abfg := NewPdfRectangle(*_cdfdg)
		if _abfg != nil {
			_ag.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _abfg)
			return nil, _abfg
		}
		_ddbgg.BBox = _bcffb
	}
	_bced = _fdfdb.Get("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s")
	if _bced != nil {
		_bced = _dg.TraceToDirectObject(_bced)
		_cddcbg, _cfaad := _bced.(*_dg.PdfObjectBool)
		if !_cfaad {
			_ag.Log.Debug("A\u006e\u0074\u0069\u0041\u006c\u0069\u0061\u0073\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020ty\u0070\u0065\u002c\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020bo\u006f\u006c \u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bced)
			return nil, _dg.ErrTypeError
		}
		_ddbgg.AntiAlias = _cddcbg
	}
	switch *_cdba {
	case 1:
		_dcceg, _gdeg := _ggdcc(_fdfdb)
		if _gdeg != nil {
			return nil, _gdeg
		}
		_dcceg.PdfShading = _ddbgg
		_ddbgg._eeddb = _dcceg
		return _ddbgg, nil
	case 2:
		_gfdf, _ecbfb := _dgag(_fdfdb)
		if _ecbfb != nil {
			return nil, _ecbfb
		}
		_gfdf.PdfShading = _ddbgg
		_ddbgg._eeddb = _gfdf
		return _ddbgg, nil
	case 3:
		_dfafa, _gfege := _caff(_fdfdb)
		if _gfege != nil {
			return nil, _gfege
		}
		_dfafa.PdfShading = _ddbgg
		_ddbgg._eeddb = _dfafa
		return _ddbgg, nil
	case 4:
		_cfgbb, _fafa := _eaac(_fdfdb)
		if _fafa != nil {
			return nil, _fafa
		}
		_cfgbb.PdfShading = _ddbgg
		_ddbgg._eeddb = _cfgbb
		return _ddbgg, nil
	case 5:
		_fcgae, _dafdd := _edaac(_fdfdb)
		if _dafdd != nil {
			return nil, _dafdd
		}
		_fcgae.PdfShading = _ddbgg
		_ddbgg._eeddb = _fcgae
		return _ddbgg, nil
	case 6:
		_feeef, _ecdea := _eabg(_fdfdb)
		if _ecdea != nil {
			return nil, _ecdea
		}
		_feeef.PdfShading = _ddbgg
		_ddbgg._eeddb = _feeef
		return _ddbgg, nil
	case 7:
		_fbfdf, _cdced := _dfdb(_fdfdb)
		if _cdced != nil {
			return nil, _cdced
		}
		_fbfdf.PdfShading = _ddbgg
		_ddbgg._eeddb = _fbfdf
		return _ddbgg, nil
	}
	return nil, _bf.New("u\u006ek\u006e\u006f\u0077\u006e\u0020\u0073\u0068\u0061d\u0069\u006e\u0067\u0020ty\u0070\u0065")
}
func (_gcdfe *PdfWriter) optimize() error {
	if _gcdfe._egeac == nil {
		return nil
	}
	var _fgcbfa error
	_gcdfe._agaba, _fgcbfa = _gcdfe._egeac.Optimize(_gcdfe._agaba)
	if _fgcbfa != nil {
		return _fgcbfa
	}
	_edbgb := make(map[_dg.PdfObject]struct{}, len(_gcdfe._agaba))
	for _, _gefeb := range _gcdfe._agaba {
		_edbgb[_gefeb] = struct{}{}
	}
	_gcdfe._fdbfa = _edbgb
	return nil
}

// String returns a string representation of what flags are set.
func (_fgad FieldFlag) String() string {
	_gebae := ""
	if _fgad == FieldFlagClear {
		_gebae = "\u0043\u006c\u0065a\u0072"
		return _gebae
	}
	if _fgad&FieldFlagReadOnly > 0 {
		_gebae += "\u007cR\u0065\u0061\u0064\u004f\u006e\u006cy"
	}
	if _fgad&FieldFlagRequired > 0 {
		_gebae += "\u007cR\u0065\u0071\u0075\u0069\u0072\u0065d"
	}
	if _fgad&FieldFlagNoExport > 0 {
		_gebae += "\u007cN\u006f\u0045\u0078\u0070\u006f\u0072t"
	}
	if _fgad&FieldFlagNoToggleToOff > 0 {
		_gebae += "\u007c\u004e\u006f\u0054\u006f\u0067\u0067\u006c\u0065T\u006f\u004f\u0066\u0066"
	}
	if _fgad&FieldFlagRadio > 0 {
		_gebae += "\u007c\u0052\u0061\u0064\u0069\u006f"
	}
	if _fgad&FieldFlagPushbutton > 0 {
		_gebae += "|\u0050\u0075\u0073\u0068\u0062\u0075\u0074\u0074\u006f\u006e"
	}
	if _fgad&FieldFlagRadiosInUnision > 0 {
		_gebae += "\u007c\u0052a\u0064\u0069\u006fs\u0049\u006e\u0055\u006e\u0069\u0073\u0069\u006f\u006e"
	}
	if _fgad&FieldFlagMultiline > 0 {
		_gebae += "\u007c\u004d\u0075\u006c\u0074\u0069\u006c\u0069\u006e\u0065"
	}
	if _fgad&FieldFlagPassword > 0 {
		_gebae += "\u007cP\u0061\u0073\u0073\u0077\u006f\u0072d"
	}
	if _fgad&FieldFlagFileSelect > 0 {
		_gebae += "|\u0046\u0069\u006c\u0065\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _fgad&FieldFlagDoNotScroll > 0 {
		_gebae += "\u007c\u0044\u006fN\u006f\u0074\u0053\u0063\u0072\u006f\u006c\u006c"
	}
	if _fgad&FieldFlagComb > 0 {
		_gebae += "\u007c\u0043\u006fm\u0062"
	}
	if _fgad&FieldFlagRichText > 0 {
		_gebae += "\u007cR\u0069\u0063\u0068\u0054\u0065\u0078t"
	}
	if _fgad&FieldFlagDoNotSpellCheck > 0 {
		_gebae += "\u007c\u0044o\u004e\u006f\u0074S\u0070\u0065\u006c\u006c\u0043\u0068\u0065\u0063\u006b"
	}
	if _fgad&FieldFlagCombo > 0 {
		_gebae += "\u007c\u0043\u006f\u006d\u0062\u006f"
	}
	if _fgad&FieldFlagEdit > 0 {
		_gebae += "\u007c\u0045\u0064i\u0074"
	}
	if _fgad&FieldFlagSort > 0 {
		_gebae += "\u007c\u0053\u006fr\u0074"
	}
	if _fgad&FieldFlagMultiSelect > 0 {
		_gebae += "\u007c\u004d\u0075l\u0074\u0069\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _fgad&FieldFlagCommitOnSelChange > 0 {
		_gebae += "\u007cC\u006fm\u006d\u0069\u0074\u004f\u006eS\u0065\u006cC\u0068\u0061\u006e\u0067\u0065"
	}
	return _ga.Trim(_gebae, "\u007c")
}

// NewPdfAnnotationHighlight returns a new text highlight annotation.
func NewPdfAnnotationHighlight() *PdfAnnotationHighlight {
	_dbc := NewPdfAnnotation()
	_ceb := &PdfAnnotationHighlight{}
	_ceb.PdfAnnotation = _dbc
	_ceb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dbc.SetContext(_ceb)
	return _ceb
}
func (_gggea *PdfWriter) setWriter(_fbfcg _cf.Writer) {
	_gggea._fbbfc = _gggea._bgggdg
	_gggea._bddfa = _ba.NewWriter(_fbfcg)
}

// AllFields returns a flattened list of all fields in the form.
func (_aebgc *PdfAcroForm) AllFields() []*PdfField {
	if _aebgc == nil {
		return nil
	}
	var _gbef []*PdfField
	if _aebgc.Fields != nil {
		for _, _afga := range *_aebgc.Fields {
			_gbef = append(_gbef, _aacd(_afga)...)
		}
	}
	return _gbef
}
func _adddc(_adcb _dg.PdfObject) (*PdfColorspaceICCBased, error) {
	_cdag := &PdfColorspaceICCBased{}
	if _eedae, _cceg := _adcb.(*_dg.PdfIndirectObject); _cceg {
		_cdag._fbdff = _eedae
	}
	_adcb = _dg.TraceToDirectObject(_adcb)
	_gfga, _bccb := _adcb.(*_dg.PdfObjectArray)
	if !_bccb {
		return nil, _b.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _gfga.Len() != 2 {
		return nil, _b.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020c\u006f\u006c\u006fr\u0073p\u0061\u0063\u0065")
	}
	_adcb = _dg.TraceToDirectObject(_gfga.Get(0))
	_cbfc, _bccb := _adcb.(*_dg.PdfObjectName)
	if !_bccb {
		return nil, _b.Errorf("\u0049\u0043\u0043B\u0061\u0073\u0065\u0064 \u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_cbfc != "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064" {
		return nil, _b.Errorf("\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0049\u0043\u0043\u0042a\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_adcb = _gfga.Get(1)
	_gacgd, _bccb := _dg.GetStream(_adcb)
	if !_bccb {
		_ag.Log.Error("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006e\u006f\u0074\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067\u0020t\u006f\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020%\u0054", _adcb)
		return nil, _b.Errorf("\u0049\u0043\u0043Ba\u0073\u0065\u0064\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_ebac := _gacgd.PdfObjectDictionary
	_bcfca, _bccb := _ebac.Get("\u004e").(*_dg.PdfObjectInteger)
	if !_bccb {
		return nil, _b.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006d\u0069\u0073\u0073\u0069n\u0067\u0020\u004e\u0020\u0066\u0072\u006fm\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074")
	}
	if *_bcfca != 1 && *_bcfca != 3 && *_bcfca != 4 {
		return nil, _b.Errorf("\u0049\u0043\u0043\u0042\u0061s\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u006e\u006f\u0074\u0020\u0031\u002c\u0033\u002c\u0034\u0029")
	}
	_cdag.N = int(*_bcfca)
	if _abae := _ebac.Get("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e"); _abae != nil {
		_cabe, _cgcaf := NewPdfColorspaceFromPdfObject(_abae)
		if _cgcaf != nil {
			return nil, _cgcaf
		}
		_cdag.Alternate = _cabe
	}
	if _bbbcc := _ebac.Get("\u0052\u0061\u006eg\u0065"); _bbbcc != nil {
		_bbbcc = _dg.TraceToDirectObject(_bbbcc)
		_fegaa, _dddbe := _bbbcc.(*_dg.PdfObjectArray)
		if !_dddbe {
			return nil, _b.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0061\u006e \u0061r\u0072\u0061\u0079")
		}
		if _fegaa.Len() != 2*_cdag.N {
			return nil, _b.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0052\u0061\u006e\u0067e\u0020\u0077\u0072\u006f\u006e\u0067 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073")
		}
		_eccg, _bgfc := _fegaa.GetAsFloat64Slice()
		if _bgfc != nil {
			return nil, _bgfc
		}
		_cdag.Range = _eccg
	} else {
		_cdag.Range = make([]float64, 2*_cdag.N)
		for _bace := 0; _bace < _cdag.N; _bace++ {
			_cdag.Range[2*_bace] = 0.0
			_cdag.Range[2*_bace+1] = 1.0
		}
	}
	if _bagb := _ebac.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _bagb != nil {
		_dcaed, _dcagg := _bagb.(*_dg.PdfObjectStream)
		if !_dcagg {
			return nil, _b.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u004de\u0074\u0061\u0064\u0061\u0074\u0061\u0020n\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		_cdag.Metadata = _dcaed
	}
	_baaf, _abdcd := _dg.DecodeStream(_gacgd)
	if _abdcd != nil {
		return nil, _abdcd
	}
	_cdag.Data = _baaf
	_cdag._dafg = _gacgd
	return _cdag, nil
}

// HasColorspaceByName checks if the colorspace with the specified name exists in the page resources.
func (_fggg *PdfPageResources) HasColorspaceByName(keyName _dg.PdfObjectName) bool {
	_eaceb, _aefa := _fggg.GetColorspaces()
	if _aefa != nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _aefa)
		return false
	}
	if _eaceb == nil {
		return false
	}
	_, _aceed := _eaceb.Colorspaces[string(keyName)]
	return _aceed
}

// DecodeArray returns the component range values for the Indexed colorspace.
func (_fdff *PdfColorspaceSpecialIndexed) DecodeArray() []float64 {
	return []float64{0, float64(_fdff.HiVal)}
}

// GetContainingPdfObject implements model.PdfModel interface.
func (_eadbe *PdfOutputIntent) GetContainingPdfObject() _dg.PdfObject { return _eadbe._cfaef }

// PdfActionGoTo3DView represents a GoTo3DView action.
type PdfActionGoTo3DView struct {
	*PdfAction
	TA _dg.PdfObject
	V  _dg.PdfObject
}

// Width returns the width of `rect`.
func (_fbcdc *PdfRectangle) Width() float64 { return _cg.Abs(_fbcdc.Urx - _fbcdc.Llx) }

// PdfActionGoToE represents a GoToE action.
type PdfActionGoToE struct {
	*PdfAction
	F         *PdfFilespec
	D         _dg.PdfObject
	NewWindow _dg.PdfObject
	T         _dg.PdfObject
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_ccccb pdfCIDFontType0) GetRuneMetrics(r rune) (_bbg.CharMetrics, bool) {
	return _bbg.CharMetrics{Wx: _ccccb._gfdca}, true
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_cedf *PdfColorspaceDeviceGray) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bgcb, _febe := _dg.GetNumbersAsFloat(objects)
	if _febe != nil {
		return nil, _febe
	}
	return _cedf.ColorFromFloats(_bgcb)
}

// AddPages adds pages to be appended to the end of the source PDF.
func (_dgcb *PdfAppender) AddPages(pages ...*PdfPage) {
	for _, _geda := range pages {
		_geda = _geda.Duplicate()
		_dccea(_geda)
		_dgcb._ggdd = append(_dgcb._ggdd, _geda)
	}
}
func (_dace *pdfFontType0) subsetRegistered() error {
	_gegdg, _bcdgc := _dace.DescendantFont._cadf.(*pdfCIDFontType2)
	if !_bcdgc {
		_ag.Log.Debug("\u0046\u006fnt\u0020\u006e\u006ft\u0020\u0073\u0075\u0070por\u0074ed\u0020\u0066\u006f\u0072\u0020\u0073\u0075bs\u0065\u0074\u0074\u0069\u006e\u0067\u0020%\u0054", _dace.DescendantFont)
		return nil
	}
	if _gegdg == nil {
		return nil
	}
	if _gegdg._ccfb == nil {
		_ag.Log.Debug("\u004d\u0069\u0073si\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return nil
	}
	if _dace._ggec == nil {
		_ag.Log.Debug("\u004e\u006f\u0020e\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0073\u0075\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0067\u006e\u006f\u0072\u0065\u0064")
		return nil
	}
	_gcdg, _bcdgc := _dg.GetStream(_gegdg._ccfb.FontFile2)
	if !_bcdgc {
		_ag.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _bf.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_degc, _babc := _dg.DecodeStream(_gcdg)
	if _babc != nil {
		_ag.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _babc)
		return _babc
	}
	_bfabf, _babc := _bfc.Parse(_bc.NewReader(_degc))
	if _babc != nil {
		_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_gcdg.Stream))
		return _babc
	}
	var _bggb []rune
	var _ffcba *_bfc.Font
	switch _bfgfg := _dace._ggec.(type) {
	case *_bd.TrueTypeFontEncoder:
		_bggb = _bfgfg.RegisteredRunes()
		_ffcba, _babc = _bfabf.SubsetKeepRunes(_bggb)
		if _babc != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _babc)
			return _babc
		}
		_bfgfg.SubsetRegistered()
	case *_bd.IdentityEncoder:
		_bggb = _bfgfg.RegisteredRunes()
		_fcga := make([]_bfc.GlyphIndex, len(_bggb))
		for _dfgf, _aedb := range _bggb {
			_fcga[_dfgf] = _bfc.GlyphIndex(_aedb)
		}
		_ffcba, _babc = _bfabf.SubsetKeepIndices(_fcga)
		if _babc != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _babc)
			return _babc
		}
	case _bd.SimpleEncoder:
		_bbadb := _bfgfg.Charcodes()
		for _, _ebcdc := range _bbadb {
			_dbcbc, _gbaf := _bfgfg.CharcodeToRune(_ebcdc)
			if !_gbaf {
				_ag.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0020\u0074\u006f \u0072\u0075\u006e\u0065\u003a\u0020\u0025\u0064", _ebcdc)
				continue
			}
			_bggb = append(_bggb, _dbcbc)
		}
	default:
		return _b.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069\u006eg\u003a\u0020\u0025\u0054", _dace._ggec)
	}
	var _ebff _bc.Buffer
	_babc = _ffcba.Write(&_ebff)
	if _babc != nil {
		_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _babc)
		return _babc
	}
	if _dace._ecfb != nil {
		_aedd := make(map[_ff.CharCode]rune, len(_bggb))
		for _, _fdce := range _bggb {
			_daae, _bafeg := _dace._ggec.RuneToCharcode(_fdce)
			if !_bafeg {
				continue
			}
			_aedd[_ff.CharCode(_daae)] = _fdce
		}
		_dace._ecfb = _ff.NewToUnicodeCMap(_aedd)
	}
	_gcdg, _babc = _dg.MakeStream(_ebff.Bytes(), _dg.NewFlateEncoder())
	if _babc != nil {
		_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _babc)
		return _babc
	}
	_gcdg.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _dg.MakeInteger(int64(_ebff.Len())))
	if _eeee, _ffcdf := _dg.GetStream(_gegdg._ccfb.FontFile2); _ffcdf {
		*_eeee = *_gcdg
	} else {
		_gegdg._ccfb.FontFile2 = _gcdg
	}
	_abefc := _daffeb()
	if len(_dace._ecbf) > 0 {
		_dace._ecbf = _bbgcf(_dace._ecbf, _abefc)
	}
	if len(_gegdg._ecbf) > 0 {
		_gegdg._ecbf = _bbgcf(_gegdg._ecbf, _abefc)
	}
	if len(_dace._cefg) > 0 {
		_dace._cefg = _bbgcf(_dace._cefg, _abefc)
	}
	if _gegdg._ccfb != nil {
		_ccbdb, _cgdc := _dg.GetName(_gegdg._ccfb.FontName)
		if _cgdc && len(_ccbdb.String()) > 0 {
			_edfac := _bbgcf(_ccbdb.String(), _abefc)
			_gegdg._ccfb.FontName = _dg.MakeName(_edfac)
		}
	}
	return nil
}

// FieldFlattenOpts defines a set of options which can be used to configure
// the field flattening process.
type FieldFlattenOpts struct {

	// FilterFunc allows filtering the form fields used in the flattening
	// process. If the filter function returns true, the field is flattened,
	// otherwise it is skipped.
	// If a non-terminal field is discarded, all of its children (the fields
	// present in the Kids array) are discarded as well.
	// Non-terminal fields are kept in the AcroForm if one or more of their
	// child fields have not been selected for flattening.
	// If a filter function is not provided, all form fields are flattened.
	FilterFunc FieldFilterFunc

	// AnnotFilterFunc allows filtering the annotations in the flattening
	// process. If the filter function returns true, the annotation is flattened,
	// otherwise it is skipped.
	AnnotFilterFunc AnnotFilterFunc
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_dgab *PdfColorspaceCalGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_agacc := vals[0]
	if _agacc < 0.0 || _agacc > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _agacc)
		return nil, ErrColorOutOfRange
	}
	_aebd := NewPdfColorCalGray(_agacc)
	return _aebd, nil
}
func _gegbd(_dgegc _dg.PdfObject, _fdfd bool) (*PdfFont, error) {
	_fafca, _agagf, _beaef := _cbaa(_dgegc)
	if _fafca != nil {
		_acbbg(_fafca)
	}
	if _beaef != nil {
		if _beaef == ErrType1CFontNotSupported {
			_fbccc, _ccddg := _cfaa(_fafca, _agagf, nil)
			if _ccddg != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _agagf, _ccddg)
				return nil, _beaef
			}
			return &PdfFont{_cadf: _fbccc}, _beaef
		}
		return nil, _beaef
	}
	_aefb := &PdfFont{}
	switch _agagf._bcga {
	case "\u0054\u0079\u0070e\u0030":
		if !_fdfd {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u004c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u00650\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _agagf)
			return nil, _bf.New("\u0063\u0079\u0063\u006cic\u0061\u006c\u0020\u0074\u0079\u0070\u0065\u0030\u0020\u006c\u006f\u0061\u0064\u0069n\u0067")
		}
		_cefcc, _geff := _cefdc(_fafca, _agagf)
		if _geff != nil {
			_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0057\u0068\u0069l\u0065\u0020\u006c\u006f\u0061\u0064\u0069ng\u0020\u0054\u0079\u0070e\u0030\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066on\u0074\u003d%\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _agagf, _geff)
			return nil, _geff
		}
		_aefb._cadf = _cefcc
	case "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		var _dafc *pdfFontSimple
		_bdbb, _abgb := _bbg.NewStdFontByName(_bbg.StdFontName(_agagf._ecbf))
		if _abgb {
			_fbaaa := _gecgd(_bdbb)
			_aefb._cadf = &_fbaaa
			_aadg := _dg.TraceToDirectObject(_fbaaa.ToPdfObject())
			_fgcac, _gdae, _ccebc := _cbaa(_aadg)
			if _ccebc != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _agagf, _fbaaa)
				return nil, _ccebc
			}
			for _, _abga := range _fafca.Keys() {
				_fgcac.Set(_abga, _fafca.Get(_abga))
			}
			_dafc, _ccebc = _cfaa(_fgcac, _gdae, _fbaaa._dbdb)
			if _ccebc != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _agagf, _fbaaa)
				return nil, _ccebc
			}
			_dafc._cfagf = _fbaaa._cfagf
			_dafc._bfdee = _fbaaa._bfdee
			if _dafc._gagbe == nil {
				_dafc._gagbe = _fbaaa._gagbe
			}
		} else {
			_dafc, _beaef = _cfaa(_fafca, _agagf, nil)
			if _beaef != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _agagf, _beaef)
				return nil, _beaef
			}
		}
		_beaef = _dafc.addEncoding()
		if _beaef != nil {
			return nil, _beaef
		}
		if _abgb {
			_dafc.updateStandard14Font()
		}
		if _abgb && _dafc._bdeed == nil && _dafc._dbdb == nil {
			_ag.Log.Error("\u0073\u0069\u006d\u0070\u006c\u0065\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _dafc)
			_ag.Log.Error("\u0066n\u0074\u003d\u0025\u002b\u0076", _bdbb)
		}
		if len(_dafc._cfagf) == 0 {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u006f\u0020\u0077\u0069d\u0074h\u0073.\u0020\u0066\u006f\u006e\u0074\u003d\u0025s", _dafc)
		}
		_aefb._cadf = _dafc
	case "\u0054\u0079\u0070e\u0033":
		_cgaa, _feffa := _bagde(_fafca, _agagf)
		if _feffa != nil {
			_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020W\u0068\u0069\u006c\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0074y\u0070\u0065\u0033\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _feffa)
			return nil, _feffa
		}
		_aefb._cadf = _cgaa
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030":
		_fafg, _abfe := _fbff(_fafca, _agagf)
		if _abfe != nil {
			_ag.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0057\u0068i\u006c\u0065\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069d \u0066\u006f\u006et\u0020\u0074y\u0070\u0065\u0030\u0020\u0066\u006fn\u0074\u003a \u0025\u0076", _abfe)
			return nil, _abfe
		}
		_aefb._cadf = _fafg
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_ebafa, _eecf := _dadaa(_fafca, _agagf)
		if _eecf != nil {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0057\u0068\u0069l\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069\u0064\u0020f\u006f\u006e\u0074\u0020\u0074yp\u0065\u0032\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _agagf, _eecf)
			return nil, _eecf
		}
		_aefb._cadf = _ebafa
	default:
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020U\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020f\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073", _agagf)
		return nil, _b.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0066\u006f\u006e\u0074\u0020\u0074y\u0070\u0065\u003a\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _agagf)
	}
	return _aefb, nil
}
func _cfaa(_geef *_dg.PdfObjectDictionary, _dbbf *fontCommon, _gedbf _bd.TextEncoder) (*pdfFontSimple, error) {
	_dddef := _beabe(_dbbf)
	_dddef._dbdb = _gedbf
	if _gedbf == nil {
		_bfgab := _geef.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
		if _bfgab == nil {
			_bfgab = _dg.MakeInteger(0)
		}
		_dddef.FirstChar = _bfgab
		_cfgb, _eecg := _dg.GetIntVal(_bfgab)
		if !_eecg {
			_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _bfgab)
			return nil, _dg.ErrTypeError
		}
		_dbce := _bd.CharCode(_cfgb)
		_bfgab = _geef.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
		if _bfgab == nil {
			_bfgab = _dg.MakeInteger(255)
		}
		_dddef.LastChar = _bfgab
		_cfgb, _eecg = _dg.GetIntVal(_bfgab)
		if !_eecg {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _bfgab)
			return nil, _dg.ErrTypeError
		}
		_aacf := _bd.CharCode(_cfgb)
		_dddef._cfagf = make(map[_bd.CharCode]float64)
		_bfgab = _geef.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _bfgab != nil {
			_dddef.Widths = _bfgab
			_cage, _fafe := _dg.GetArray(_bfgab)
			if !_fafe {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _bfgab)
				return nil, _dg.ErrTypeError
			}
			_fbdg, _bbaeg := _cage.ToFloat64Array()
			if _bbaeg != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
				return nil, _bbaeg
			}
			if len(_fbdg) != int(_aacf-_dbce+1) {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _aacf-_dbce+1, len(_fbdg))
				return nil, _dg.ErrRangeError
			}
			for _cgcg, _bccdf := range _fbdg {
				_dddef._cfagf[_dbce+_bd.CharCode(_cgcg)] = _bccdf
			}
		}
	}
	_dddef.Encoding = _dg.TraceToDirectObject(_geef.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	return _dddef, nil
}

// ColorToRGB verifies that the input color is an RGB color. Method exists in
// order to satisfy the PdfColorspace interface.
func (_dcgf *PdfColorspaceDeviceRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_gafc, _gcfea := color.(*PdfColorDeviceRGB)
	if !_gcfea {
		_ag.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072 \u006e\u006f\u0074\u0020\u0064\u0065\u0076\u0069\u0063\u0065 \u0052\u0047\u0042")
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return _gafc, nil
}

// PdfActionThread represents a thread action.
type PdfActionThread struct {
	*PdfAction
	F *PdfFilespec
	D _dg.PdfObject
	B _dg.PdfObject
}

// PdfShadingType2 is an Axial shading.
type PdfShadingType2 struct {
	*PdfShading
	Coords   *_dg.PdfObjectArray
	Domain   *_dg.PdfObjectArray
	Function []PdfFunction
	Extend   *_dg.PdfObjectArray
}

// G returns the value of the green component of the color.
func (_egege *PdfColorDeviceRGB) G() float64 { return _egege[1] }
func (_dcbab *PdfWriter) writeDocumentVersion() {
	if _dcbab._bbac {
		_dcbab.writeString("\u000a")
	} else {
		_dcbab.writeString(_b.Sprintf("\u0025\u0025\u0050D\u0046\u002d\u0025\u0064\u002e\u0025\u0064\u000a", _dcbab._efacd.Major, _dcbab._efacd.Minor))
		_dcbab.writeString("\u0025\u00e2\u00e3\u00cf\u00d3\u000a")
	}
}

// NewPdfColorspaceLab returns a new Lab colorspace object.
func NewPdfColorspaceLab() *PdfColorspaceLab {
	_bfccb := &PdfColorspaceLab{}
	_bfccb.BlackPoint = []float64{0.0, 0.0, 0.0}
	_bfccb.Range = []float64{-100, 100, -100, 100}
	return _bfccb
}

// ToPdfObject implements interface PdfModel.
func (_fgd *PdfAnnotationFreeText) ToPdfObject() _dg.PdfObject {
	_fgd.PdfAnnotation.ToPdfObject()
	_dae := _fgd._cdf
	_aca := _dae.PdfObject.(*_dg.PdfObjectDictionary)
	_fgd.PdfAnnotationMarkup.appendToPdfDictionary(_aca)
	_aca.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074"))
	_aca.SetIfNotNil("\u0044\u0041", _fgd.DA)
	_aca.SetIfNotNil("\u0051", _fgd.Q)
	_aca.SetIfNotNil("\u0052\u0043", _fgd.RC)
	_aca.SetIfNotNil("\u0044\u0053", _fgd.DS)
	_aca.SetIfNotNil("\u0043\u004c", _fgd.CL)
	_aca.SetIfNotNil("\u0049\u0054", _fgd.IT)
	_aca.SetIfNotNil("\u0042\u0045", _fgd.BE)
	_aca.SetIfNotNil("\u0052\u0044", _fgd.RD)
	_aca.SetIfNotNil("\u0042\u0053", _fgd.BS)
	_aca.SetIfNotNil("\u004c\u0045", _fgd.LE)
	return _dae
}

// ColorFromPdfObjects loads the color from PDF objects.
// The first objects (if present) represent the color in underlying colorspace.  The last one represents
// the name of the pattern.
func (_dedcdb *PdfColorspaceSpecialPattern) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) < 1 {
		return nil, _bf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_egdb := &PdfColorPattern{}
	_adfdf, _ddba := objects[len(objects)-1].(*_dg.PdfObjectName)
	if !_ddba {
		_ag.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020a\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", objects[len(objects)-1])
		return nil, ErrTypeCheck
	}
	_egdb.PatternName = *_adfdf
	if len(objects) > 1 {
		_cbcc := objects[0 : len(objects)-1]
		if _dedcdb.UnderlyingCS == nil {
			_ag.Log.Debug("P\u0061\u0074t\u0065\u0072\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0077\u0069\u0074\u0068\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u0020\u0062\u0075\u0074\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0063\u0073\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
			return nil, _bf.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
		}
		_abde, _abaaf := _dedcdb.UnderlyingCS.ColorFromPdfObjects(_cbcc)
		if _abaaf != nil {
			_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0076\u0069\u0061\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0063\u0073\u003a\u0020\u0025\u0076", _abaaf)
			return nil, _abaaf
		}
		_egdb.Color = _abde
	}
	return _egdb, nil
}
func _fdabf(_ecbeb _dg.PdfObject) {
	_ag.Log.Debug("\u006f\u0062\u006a\u003a\u0020\u0025\u0054\u0020\u0025\u0073", _ecbeb, _ecbeb.String())
	if _dgcae, _abbec := _ecbeb.(*_dg.PdfObjectStream); _abbec {
		_bdfa, _gcfce := _dg.DecodeStream(_dgcae)
		if _gcfce != nil {
			_ag.Log.Debug("\u0045r\u0072\u006f\u0072\u003a\u0020\u0025v", _gcfce)
			return
		}
		_ag.Log.Debug("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _bdfa)
	} else if _eggagb, _cdfee := _ecbeb.(*_dg.PdfIndirectObject); _cdfee {
		_ag.Log.Debug("\u0025\u0054\u0020%\u0076", _eggagb.PdfObject, _eggagb.PdfObject)
		_ag.Log.Debug("\u0025\u0073", _eggagb.PdfObject.String())
	}
}

// PdfColorspace interface defines the common methods of a PDF colorspace.
// The colorspace defines the data storage format for each color and color representation.
//
// Device based colorspace, specified by name
// - /DeviceGray
// - /DeviceRGB
// - /DeviceCMYK
//
// CIE based colorspace specified by [name, dictionary]
// - [/CalGray dict]
// - [/CalRGB dict]
// - [/Lab dict]
// - [/ICCBased dict]
//
// Special colorspaces
// - /Pattern
// - /Indexed
// - /Separation
// - /DeviceN
//
// Work is in progress to support all colorspaces. At the moment ICCBased color spaces fall back to the alternate
// colorspace which works OK in most cases. For full color support, will need fully featured ICC support.
type PdfColorspace interface {

	// String returns the PdfColorspace's name.
	String() string

	// ImageToRGB converts an Image in a given PdfColorspace to an RGB image.
	ImageToRGB(Image) (Image, error)

	// ColorToRGB converts a single color in a given PdfColorspace to an RGB color.
	ColorToRGB(_dcadd PdfColor) (PdfColor, error)

	// GetNumComponents returns the number of components in the PdfColorspace.
	GetNumComponents() int

	// ToPdfObject returns a PdfObject representation of the PdfColorspace.
	ToPdfObject() _dg.PdfObject

	// ColorFromPdfObjects returns a PdfColor in the given PdfColorspace from an array of PdfObject where each
	// PdfObject represents a numeric value.
	ColorFromPdfObjects(_eggag []_dg.PdfObject) (PdfColor, error)

	// ColorFromFloats returns a new PdfColor based on input color components for a given PdfColorspace.
	ColorFromFloats(_gccc []float64) (PdfColor, error)

	// DecodeArray returns the Decode array for the PdfColorSpace, i.e. the range of each component.
	DecodeArray() []float64
}

// NewPdfColorPattern returns an empty color pattern.
func NewPdfColorPattern() *PdfColorPattern { _egcba := &PdfColorPattern{}; return _egcba }
func _gadec(_gdbb []byte) (_ggdfa, _fdfeb string, _bcgga error) {
	_ag.Log.Trace("g\u0065\u0074\u0041\u0053CI\u0049S\u0065\u0063\u0074\u0069\u006fn\u0073\u003a\u0020\u0025\u0064\u0020", len(_gdbb))
	_ggca := _caga.FindIndex(_gdbb)
	if _ggca == nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0067\u0065\u0074\u0041\u0053\u0043\u0049\u0049\u0053\u0065\u0063\u0074\u0069o\u006e\u0073\u002e\u0020\u004e\u006f\u0020d\u0069\u0063\u0074\u002e")
		return "", "", _dg.ErrTypeError
	}
	_gfcc := _ggca[1]
	_cbbf := _ga.Index(string(_gdbb[_gfcc:]), _degb)
	if _cbbf < 0 {
		_ggdfa = string(_gdbb[_gfcc:])
		return _ggdfa, "", nil
	}
	_eecfa := _gfcc + _cbbf
	_ggdfa = string(_gdbb[_gfcc:_eecfa])
	_cfgce := _eecfa
	_cbbf = _ga.Index(string(_gdbb[_cfgce:]), _cgcga)
	if _cbbf < 0 {
		_ag.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0067e\u0074\u0041\u0053\u0043\u0049\u0049\u0053e\u0063\u0074\u0069\u006f\u006e\u0073\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bcgga)
		return "", "", _dg.ErrTypeError
	}
	_beecg := _cfgce + _cbbf
	_fdfeb = string(_gdbb[_cfgce:_beecg])
	return _ggdfa, _fdfeb, nil
}

// ToPdfObject implements interface PdfModel.
func (_ecgb *PdfAnnotationWatermark) ToPdfObject() _dg.PdfObject {
	_ecgb.PdfAnnotation.ToPdfObject()
	_bfdc := _ecgb._cdf
	_cgb := _bfdc.PdfObject.(*_dg.PdfObjectDictionary)
	_cgb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k"))
	_cgb.SetIfNotNil("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074", _ecgb.FixedPrint)
	return _bfdc
}

// PdfActionSubmitForm represents a submitForm action.
type PdfActionSubmitForm struct {
	*PdfAction
	F      *PdfFilespec
	Fields _dg.PdfObject
	Flags  _dg.PdfObject
}

func (_aacb *PdfReader) newPdfAnnotationPolyLineFromDict(_bbab *_dg.PdfObjectDictionary) (*PdfAnnotationPolyLine, error) {
	_fge := PdfAnnotationPolyLine{}
	_dgf, _defg := _aacb.newPdfAnnotationMarkupFromDict(_bbab)
	if _defg != nil {
		return nil, _defg
	}
	_fge.PdfAnnotationMarkup = _dgf
	_fge.Vertices = _bbab.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_fge.LE = _bbab.Get("\u004c\u0045")
	_fge.BS = _bbab.Get("\u0042\u0053")
	_fge.IC = _bbab.Get("\u0049\u0043")
	_fge.BE = _bbab.Get("\u0042\u0045")
	_fge.IT = _bbab.Get("\u0049\u0054")
	_fge.Measure = _bbab.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_fge, nil
}

// ContentStreamWrapper wraps the Page's contentstream into q ... Q blocks.
type ContentStreamWrapper interface{ WrapContentStream(_afgf *PdfPage) error }

// NewPdfActionRendition returns a new "rendition" action.
func NewPdfActionRendition() *PdfActionRendition {
	_dc := NewPdfAction()
	_eade := &PdfActionRendition{}
	_eade.PdfAction = _dc
	_dc.SetContext(_eade)
	return _eade
}

var _addac = map[string]string{"\u0053\u0079\u006d\u0062\u006f\u006c": "\u0053\u0079\u006d\u0062\u006f\u006c\u0045\u006e\u0063o\u0064\u0069\u006e\u0067", "\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073": "Z\u0061p\u0066\u0044\u0069\u006e\u0067\u0062\u0061\u0074s\u0045\u006e\u0063\u006fdi\u006e\u0067"}

// ToPdfObject implements interface PdfModel.
func (_eec *PdfAnnotationSquare) ToPdfObject() _dg.PdfObject {
	_eec.PdfAnnotation.ToPdfObject()
	_gaf := _eec._cdf
	_fdba := _gaf.PdfObject.(*_dg.PdfObjectDictionary)
	if _eec.PdfAnnotationMarkup != nil {
		_eec.PdfAnnotationMarkup.appendToPdfDictionary(_fdba)
	}
	_fdba.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0053\u0071\u0075\u0061\u0072\u0065"))
	_fdba.SetIfNotNil("\u0042\u0053", _eec.BS)
	_fdba.SetIfNotNil("\u0049\u0043", _eec.IC)
	_fdba.SetIfNotNil("\u0042\u0045", _eec.BE)
	_fdba.SetIfNotNil("\u0052\u0044", _eec.RD)
	return _gaf
}

// SetName sets the `Name` field of the signature.
func (_eaag *PdfSignature) SetName(name string) { _eaag.Name = _dg.MakeString(name) }

// Clear clears flag fl from the flag and returns the resulting flag.
func (_fgbc FieldFlag) Clear(fl FieldFlag) FieldFlag { return FieldFlag(_fgbc.Mask() &^ fl.Mask()) }

// NewPdfAnnotationPopup returns a new popup annotation.
func NewPdfAnnotationPopup() *PdfAnnotationPopup {
	_egad := NewPdfAnnotation()
	_cebf := &PdfAnnotationPopup{}
	_cebf.PdfAnnotation = _egad
	_egad.SetContext(_cebf)
	return _cebf
}

// SetAction sets the PDF action for the annotation link.
func (_fgfa *PdfAnnotationLink) SetAction(action *PdfAction) {
	_fgfa._ece = action
	if action == nil {
		_fgfa.A = nil
	}
}

// DecodeArray returns the range of color component values in DeviceGray colorspace.
func (_cebfg *PdfColorspaceDeviceGray) DecodeArray() []float64 { return []float64{0, 1.0} }

// Fill populates `form` with values provided by `provider`.
func (_accef *PdfAcroForm) Fill(provider FieldValueProvider) error { return _accef.fill(provider, nil) }

// SetPdfProducer sets the Producer attribute of the output PDF.
func SetPdfProducer(producer string) { _fgefgf.Lock(); defer _fgefgf.Unlock(); _faff = producer }
func (_cdgfgc *PdfAcroForm) filteredFields(_egegd FieldFilterFunc, _agfab bool) []*PdfField {
	if _cdgfgc == nil {
		return nil
	}
	return _abcgf(_cdgfgc.Fields, _egegd, _agfab)
}

// NewPdfFilespecFromObj creates and returns a new PdfFilespec object.
func NewPdfFilespecFromObj(obj _dg.PdfObject) (*PdfFilespec, error) {
	_cbfbg := &PdfFilespec{}
	var _ffba *_dg.PdfObjectDictionary
	if _eaddd, _cbgd := _dg.GetIndirect(obj); _cbgd {
		_cbfbg._bcbe = _eaddd
		_bcee, _bddc := _dg.GetDict(_eaddd.PdfObject)
		if !_bddc {
			_ag.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _dg.ErrTypeError
		}
		_ffba = _bcee
	} else if _gafb, _bbgfa := _dg.GetDict(obj); _bbgfa {
		_cbfbg._bcbe = _gafb
		_ffba = _gafb
	} else {
		_ag.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", obj)
		return nil, _dg.ErrTypeError
	}
	if _ffba == nil {
		_ag.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _bf.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _ggba := _ffba.Get("\u0054\u0079\u0070\u0065"); _ggba != nil {
		_eaafb, _ecdef := _ggba.(*_dg.PdfObjectName)
		if !_ecdef {
			_ag.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _ggba)
		} else {
			if *_eaafb != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
				_ag.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070e\u0063\u0074\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063 \u0028\u0025\u0073\u0029", *_eaafb)
			}
		}
	}
	if _gdbd := _ffba.Get("\u0046\u0053"); _gdbd != nil {
		_cbfbg.FS = _gdbd
	}
	if _ccfd := _ffba.Get("\u0046"); _ccfd != nil {
		_cbfbg.F = _ccfd
	}
	if _ecff := _ffba.Get("\u0055\u0046"); _ecff != nil {
		_cbfbg.UF = _ecff
	}
	if _gcedg := _ffba.Get("\u0044\u004f\u0053"); _gcedg != nil {
		_cbfbg.DOS = _gcedg
	}
	if _cdebb := _ffba.Get("\u004d\u0061\u0063"); _cdebb != nil {
		_cbfbg.Mac = _cdebb
	}
	if _cgcfb := _ffba.Get("\u0055\u006e\u0069\u0078"); _cgcfb != nil {
		_cbfbg.Unix = _cgcfb
	}
	if _cgce := _ffba.Get("\u0049\u0044"); _cgce != nil {
		_cbfbg.ID = _cgce
	}
	if _abdgeb := _ffba.Get("\u0056"); _abdgeb != nil {
		_cbfbg.V = _abdgeb
	}
	if _aabe := _ffba.Get("\u0045\u0046"); _aabe != nil {
		_cbfbg.EF = _aabe
	}
	if _ebcff := _ffba.Get("\u0052\u0046"); _ebcff != nil {
		_cbfbg.RF = _ebcff
	}
	if _ecafg := _ffba.Get("\u0044\u0065\u0073\u0063"); _ecafg != nil {
		_cbfbg.Desc = _ecafg
	}
	if _debed := _ffba.Get("\u0043\u0049"); _debed != nil {
		_cbfbg.CI = _debed
	}
	return _cbfbg, nil
}

// ToPdfObject implements interface PdfModel.
func (_add *PdfActionGoTo3DView) ToPdfObject() _dg.PdfObject {
	_add.PdfAction.ToPdfObject()
	_fcc := _add._cbd
	_eac := _fcc.PdfObject.(*_dg.PdfObjectDictionary)
	_eac.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeGoTo3DView)))
	_eac.SetIfNotNil("\u0054\u0041", _add.TA)
	_eac.SetIfNotNil("\u0056", _add.V)
	return _fcc
}

// GetContainingPdfObject returns the XObject Form's containing object (indirect object).
func (_cbfcg *XObjectForm) GetContainingPdfObject() _dg.PdfObject { return _cbfcg._ebaeb }
func (_abecg *PdfReader) buildNameNodes(_afdc *_dg.PdfIndirectObject, _edgda map[_dg.PdfObject]struct{}) error {
	if _afdc == nil {
		return nil
	}
	if _, _ggfg := _edgda[_afdc]; _ggfg {
		_ag.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _afdc.ObjectNumber)
		return nil
	}
	_edgda[_afdc] = struct{}{}
	_ceebbb, _fgeff := _afdc.PdfObject.(*_dg.PdfObjectDictionary)
	if !_fgeff {
		return _bf.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _dcdg, _ddgaa := _dg.GetDict(_ceebbb.Get("\u0044\u0065\u0073t\u0073")); _ddgaa {
		_fceb, _cfgbd := _dg.GetArray(_dcdg.Get("\u004b\u0069\u0064\u0073"))
		if !_cfgbd {
			return _bf.New("\u0049n\u0076\u0061\u006c\u0069d\u0020\u004b\u0069\u0064\u0073 \u0061r\u0072a\u0079\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_ag.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _fceb)
		for _fcfe, _bcgcd := range _fceb.Elements() {
			_feafb, _gcgcce := _dg.GetIndirect(_bcgcd)
			if !_gcgcce {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u0068\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002d \u0028\u0025\u0073\u0029", _feafb)
				return _bf.New("\u0063h\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006ed\u0069r\u0065c\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
			}
			_fceb.Set(_fcfe, _feafb)
			_fggec := _abecg.buildNameNodes(_feafb, _edgda)
			if _fggec != nil {
				return _fggec
			}
		}
	}
	if _eccgd, _ddcdc := _dg.GetDict(_ceebbb); _ddcdc {
		if !_dg.IsNullObject(_eccgd.Get("\u004b\u0069\u0064\u0073")) {
			if _acadg, _cfac := _dg.GetArray(_eccgd.Get("\u004b\u0069\u0064\u0073")); _cfac {
				for _dfafg, _eaeec := range _acadg.Elements() {
					if _agdfe, _cged := _dg.GetIndirect(_eaeec); _cged {
						_acadg.Set(_dfafg, _agdfe)
						_gffeg := _abecg.buildNameNodes(_agdfe, _edgda)
						if _gffeg != nil {
							return _gffeg
						}
					}
				}
			}
		}
	}
	return nil
}

// NewPdfDateFromTime will create a PdfDate based on the given time
func NewPdfDateFromTime(timeObj _a.Time) (PdfDate, error) {
	_efefee := timeObj.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_ffgab, _ := _fbb.ParseInt(_efefee[1:3], 10, 32)
	_edeec, _ := _fbb.ParseInt(_efefee[4:6], 10, 32)
	return PdfDate{_bgfdb: int64(timeObj.Year()), _gbbge: int64(timeObj.Month()), _cbad: int64(timeObj.Day()), _cade: int64(timeObj.Hour()), _ccgbg: int64(timeObj.Minute()), _aacfe: int64(timeObj.Second()), _ggbdg: _efefee[0], _bdde: _ffgab, _gafe: _edeec}, nil
}

// ToPdfObject implements interface PdfModel.
func (_eda *PdfAnnotationSound) ToPdfObject() _dg.PdfObject {
	_eda.PdfAnnotation.ToPdfObject()
	_gcea := _eda._cdf
	_aegb := _gcea.PdfObject.(*_dg.PdfObjectDictionary)
	_eda.PdfAnnotationMarkup.appendToPdfDictionary(_aegb)
	_aegb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0053\u006f\u0075n\u0064"))
	_aegb.SetIfNotNil("\u0053\u006f\u0075n\u0064", _eda.Sound)
	_aegb.SetIfNotNil("\u004e\u0061\u006d\u0065", _eda.Name)
	return _gcea
}

// SetSubtype sets the Subtype S for given PdfOutputIntent.
func (_aadgf *PdfOutputIntent) SetSubtype(subtype PdfOutputIntentType) error {
	if !subtype.IsValid() {
		return _bf.New("\u0070\u0072o\u0076\u0069\u0064\u0065d\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u004f\u0075t\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0053\u0075b\u0054\u0079\u0070\u0065")
	}
	_aadgf.S = subtype
	return nil
}

// PdfActionGoToR represents a GoToR action.
type PdfActionGoToR struct {
	*PdfAction
	F         *PdfFilespec
	D         _dg.PdfObject
	NewWindow _dg.PdfObject
}

// AddPage adds a page to the PDF file. The new page should be an indirect object.
func (_fbgdc *PdfWriter) AddPage(page *PdfPage) error {
	const _geegdg = "\u006d\u006f\u0064el\u003a\u0050\u0064\u0066\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_dccea(page)
	_dfaba := page.ToPdfObject()
	_ag.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
	_ag.Log.Trace("\u0041p\u0070\u0065\u006e\u0064i\u006e\u0067\u0020\u0074\u006f \u0070a\u0067e\u0020\u006c\u0069\u0073\u0074\u0020\u0025T", _dfaba)
	_febcb, _abcgc := _dg.GetIndirect(_dfaba)
	if !_abcgc {
		return _bf.New("\u0070\u0061\u0067\u0065\u0020\u0073h\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_ag.Log.Trace("\u0025\u0073", _febcb)
	_ag.Log.Trace("\u0025\u0073", _febcb.PdfObject)
	_gfeca, _abcgc := _dg.GetDict(_febcb.PdfObject)
	if !_abcgc {
		return _bf.New("\u0070\u0061\u0067e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068o\u0075l\u0064 \u0062e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_dgbc, _abcgc := _dg.GetName(_gfeca.Get("\u0054\u0079\u0070\u0065"))
	if !_abcgc {
		return _b.Errorf("\u0070\u0061\u0067\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054y\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u006e\u0061m\u0065\u0020\u0028%\u0054\u0029", _gfeca.Get("\u0054\u0079\u0070\u0065"))
	}
	if _dgbc.String() != "\u0050\u0061\u0067\u0065" {
		return _bf.New("\u0066\u0069e\u006c\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u0050\u0061\u0067\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069re\u0064\u0029")
	}
	_babcg := []_dg.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	_edcgb, _bddfb := _dg.GetIndirect(_gfeca.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	_ag.Log.Trace("P\u0061g\u0065\u0020\u0050\u0061\u0072\u0065\u006e\u0074:\u0020\u0025\u0054\u0020(%\u0076\u0029", _gfeca.Get("\u0050\u0061\u0072\u0065\u006e\u0074"), _bddfb)
	for _bddfb {
		_ag.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _edcgb)
		_bdcac, _eadga := _dg.GetDict(_edcgb.PdfObject)
		if !_eadga {
			return _bf.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		for _, _dcdge := range _babcg {
			_ag.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _dcdge)
			if _gfeca.Get(_dcdge) != nil {
				_ag.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
				continue
			}
			if _dcgbe := _bdcac.Get(_dcdge); _dcgbe != nil {
				_ag.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _dcdge)
				_gfeca.Set(_dcdge, _dcgbe)
			}
		}
		_edcgb, _bddfb = _dg.GetIndirect(_bdcac.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
		_ag.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _bdcac.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	}
	_ag.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c \u0064\u006f\u006e\u0065")
	_gfeca.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _fbgdc._gbgb)
	_febcb.PdfObject = _gfeca
	_abccg, _abcgc := _dg.GetDict(_fbgdc._gbgb.PdfObject)
	if !_abcgc {
		return _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_deecb, _abcgc := _dg.GetArray(_abccg.Get("\u004b\u0069\u0064\u0073"))
	if !_abcgc {
		return _bf.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_deecb.Append(_febcb)
	_fbgdc._fegdf[_gfeca] = struct{}{}
	_edecd, _abcgc := _dg.GetInt(_abccg.Get("\u0043\u006f\u0075n\u0074"))
	if !_abcgc {
		return _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	*_edecd = *_edecd + 1
	_fbgdc.addObject(_febcb)
	_efagd := _fbgdc.addObjects(_gfeca)
	if _efagd != nil {
		return _efagd
	}
	return nil
}
func (_bfggf fontCommon) asPdfObjectDictionary(_dgacb string) *_dg.PdfObjectDictionary {
	if _dgacb != "" && _bfggf._bcga != "" && _dgacb != _bfggf._bcga {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061\u0073\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020O\u0076\u0065\u0072\u0072\u0069\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0062t\u0079\u0070\u0065\u0020\u0074\u006f \u0025\u0023\u0071 \u0025\u0073", _dgacb, _bfggf)
	} else if _dgacb == "" && _bfggf._bcga == "" {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061s\u0050\u0064\u0066Ob\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074y\u0070\u0065\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _bfggf)
	} else if _bfggf._bcga == "" {
		_bfggf._bcga = _dgacb
	}
	_babb := _dg.MakeDict()
	_babb.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0046\u006f\u006e\u0074"))
	_babb.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _dg.MakeName(_bfggf._ecbf))
	_babb.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName(_bfggf._bcga))
	if _bfggf._ccfb != nil {
		_babb.Set("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072", _bfggf._ccfb.ToPdfObject())
	}
	if _bfggf._ebbff != nil {
		_babb.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _bfggf._ebbff)
	} else if _bfggf._ecfb != nil {
		_gdcf, _bebf := _bfggf._ecfb.Stream()
		if _bebf != nil {
			_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0067\u0065\u0074\u0020C\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0065r\u0072\u003d\u0025\u0076", _bebf)
		} else {
			_babb.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _gdcf)
		}
	}
	return _babb
}

// C returns the value of the C component of the color.
func (_aebe *PdfColorCalRGB) C() float64 { return _aebe[2] }

// SetBorderWidth sets the style's border width.
func (_gaeea *PdfBorderStyle) SetBorderWidth(width float64) { _gaeea.W = &width }

// AcroFormRepairOptions contains options for rebuilding the AcroForm.
type AcroFormRepairOptions struct{}

// SetPdfTitle sets the Title attribute of the output PDF.
func SetPdfTitle(title string) { _fgefgf.Lock(); defer _fgefgf.Unlock(); _fgfda = title }

// NewStandardPdfOutputIntent creates a new standard PdfOutputIntent.
func NewStandardPdfOutputIntent(outputCondition, outputConditionIdentifier, registryName string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, RegistryName: registryName, DestOutputProfile: destOutputProfile, ColorComponents: colorComponents, _cfaef: _dg.MakeDict()}
}

// DecodeArray returns the range of color component values in the ICCBased colorspace.
func (_becb *PdfColorspaceICCBased) DecodeArray() []float64 { return _becb.Range }

var (
	_caga    = _c.MustCompile("\u005cd\u002b\u0020\u0064\u0069c\u0074\u005c\u0073\u002b\u0028d\u0075p\u005cs\u002b\u0029\u003f\u0062\u0065\u0067\u0069n")
	_caegd   = _c.MustCompile("\u005e\u005cs\u002a\u002f\u0028\u005c\u0053\u002b\u003f\u0029\u005c\u0073\u002b\u0028\u002e\u002b\u003f\u0029\u005c\u0073\u002b\u0064\u0065\u0066\\s\u002a\u0024")
	_bacefb  = _c.MustCompile("\u005e\u005c\u0073*\u0064\u0075\u0070\u005c\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002a\u002f\u0028\u005c\u0077\u002b\u003f\u0029\u0028\u003f\u003a\u005c\u002e\u005c\u0064\u002b)\u003f\u005c\u0073\u002b\u0070\u0075\u0074\u0024")
	_degb    = "\u002f\u0045\u006e\u0063od\u0069\u006e\u0067\u0020\u0032\u0035\u0036\u0020\u0061\u0072\u0072\u0061\u0079"
	_cgcga   = "\u0072\u0065\u0061d\u006f\u006e\u006c\u0079\u0020\u0064\u0065\u0066"
	_cccafaf = "\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u0066\u0069\u006c\u0065\u0020e\u0065\u0078\u0065\u0063"
)

// NewPdfAnnotationFileAttachment returns a new file attachment annotation.
func NewPdfAnnotationFileAttachment() *PdfAnnotationFileAttachment {
	_bcfc := NewPdfAnnotation()
	_dafd := &PdfAnnotationFileAttachment{}
	_dafd.PdfAnnotation = _bcfc
	_dafd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_bcfc.SetContext(_dafd)
	return _dafd
}

// SetAnnotations sets the annotations list.
func (_aafb *PdfPage) SetAnnotations(annotations []*PdfAnnotation) { _aafb._cadgg = annotations }

// PdfAnnotationTrapNet represents TrapNet annotations.
// (Section 12.5.6.21).
type PdfAnnotationTrapNet struct{ *PdfAnnotation }

func (_deffff *PdfReader) loadDSS() (*DSS, error) {
	if _deffff._baad.GetCrypter() != nil && !_deffff._baad.IsAuthenticated() {
		return nil, _b.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_bfff := _deffff._gccfb.Get("\u0044\u0053\u0053")
	if _bfff == nil {
		return nil, nil
	}
	_dfbd, _ := _dg.GetIndirect(_bfff)
	_bfff = _dg.TraceToDirectObject(_bfff)
	switch _cgeb := _bfff.(type) {
	case *_dg.PdfObjectNull:
		return nil, nil
	case *_dg.PdfObjectDictionary:
		return _ceca(_dfbd, _cgeb)
	}
	return nil, _b.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u0053\u0053 \u0065\u006e\u0074\u0072y \u0025\u0054", _bfff)
}

// PdfAnnotationStrikeOut represents StrikeOut annotations.
// (Section 12.5.6.10).
type PdfAnnotationStrikeOut struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _dg.PdfObject
}

// DefaultFont returns the default font, which is currently the built in Helvetica.
func DefaultFont() *PdfFont {
	_gdgag, _gaefc := _bbg.NewStdFontByName(HelveticaName)
	if !_gaefc {
		panic("\u0048\u0065lv\u0065\u0074\u0069c\u0061\u0020\u0073\u0068oul\u0064 a\u006c\u0077\u0061\u0079\u0073\u0020\u0062e \u0061\u0076\u0061\u0069\u006c\u0061\u0062l\u0065")
	}
	_fcdg := _gecgd(_gdgag)
	return &PdfFont{_cadf: &_fcdg}
}

// GetContext returns a reference to the subpattern entry: either PdfTilingPattern or PdfShadingPattern.
func (_egac *PdfPattern) GetContext() PdfModel { return _egac._cgdcc }

// PdfAnnotationRichMedia represents Rich Media annotations.
type PdfAnnotationRichMedia struct {
	*PdfAnnotation
	RichMediaSettings _dg.PdfObject
	RichMediaContent  _dg.PdfObject
}

// SetDocInfo sets the document /Info metadata.
// This will overwrite any globally declared document info.
func (_dfgd *PdfAppender) SetDocInfo(info *PdfInfo) { _dfgd._fega = info }
func _edda(_gbca *_dg.PdfObjectStream) (*PdfFunctionType4, error) {
	_cecgf := &PdfFunctionType4{}
	_cecgf._ggfdf = _gbca
	_dadd := _gbca.PdfObjectDictionary
	_bfcbb, _ffggf := _dg.TraceToDirectObject(_dadd.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_dg.PdfObjectArray)
	if !_ffggf {
		_ag.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _bfcbb.Len()%2 != 0 {
		_ag.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _bf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_cgcfe, _aead := _bfcbb.ToFloat64Array()
	if _aead != nil {
		return nil, _aead
	}
	_cecgf.Domain = _cgcfe
	_bfcbb, _ffggf = _dg.TraceToDirectObject(_dadd.Get("\u0052\u0061\u006eg\u0065")).(*_dg.PdfObjectArray)
	if _ffggf {
		if _bfcbb.Len() < 0 || _bfcbb.Len()%2 != 0 {
			return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_fcfb, _gcfeed := _bfcbb.ToFloat64Array()
		if _gcfeed != nil {
			return nil, _gcfeed
		}
		_cecgf.Range = _fcfb
	}
	_bfec, _aead := _dg.DecodeStream(_gbca)
	if _aead != nil {
		return nil, _aead
	}
	_cecgf._cbdbd = _bfec
	_fdaa := _gfd.NewPSParser(_bfec)
	_gbffb, _aead := _fdaa.Parse()
	if _aead != nil {
		return nil, _aead
	}
	_cecgf.Program = _gbffb
	return _cecgf, nil
}
func (_gbcb *PdfReader) newPdfAnnotation3DFromDict(_dddg *_dg.PdfObjectDictionary) (*PdfAnnotation3D, error) {
	_bdf := PdfAnnotation3D{}
	_bdf.T3DD = _dddg.Get("\u0033\u0044\u0044")
	_bdf.T3DV = _dddg.Get("\u0033\u0044\u0056")
	_bdf.T3DA = _dddg.Get("\u0033\u0044\u0041")
	_bdf.T3DI = _dddg.Get("\u0033\u0044\u0049")
	_bdf.T3DB = _dddg.Get("\u0033\u0044\u0042")
	return &_bdf, nil
}
func (_dad *PdfReader) newPdfActionThreadFromDict(_faea *_dg.PdfObjectDictionary) (*PdfActionThread, error) {
	_ggfc, _egc := _bccf(_faea.Get("\u0046"))
	if _egc != nil {
		return nil, _egc
	}
	return &PdfActionThread{D: _faea.Get("\u0044"), B: _faea.Get("\u0042"), F: _ggfc}, nil
}

// ToPdfObject returns a stream object.
func (_eaaca *XObjectForm) ToPdfObject() _dg.PdfObject {
	_dgada := _eaaca._ebaeb
	_dbafd := _dgada.PdfObjectDictionary
	if _eaaca.Filter != nil {
		_dbafd = _eaaca.Filter.MakeStreamDict()
		_dgada.PdfObjectDictionary = _dbafd
	}
	_dbafd.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_dbafd.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0046\u006f\u0072\u006d"))
	_dbafd.SetIfNotNil("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065", _eaaca.FormType)
	_dbafd.SetIfNotNil("\u0042\u0042\u006f\u0078", _eaaca.BBox)
	_dbafd.SetIfNotNil("\u004d\u0061\u0074\u0072\u0069\u0078", _eaaca.Matrix)
	if _eaaca.Resources != nil {
		_dbafd.SetIfNotNil("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _eaaca.Resources.ToPdfObject())
	}
	_dbafd.SetIfNotNil("\u0047\u0072\u006fu\u0070", _eaaca.Group)
	_dbafd.SetIfNotNil("\u0052\u0065\u0066", _eaaca.Ref)
	_dbafd.SetIfNotNil("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061", _eaaca.MetaData)
	_dbafd.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _eaaca.PieceInfo)
	_dbafd.SetIfNotNil("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _eaaca.LastModified)
	_dbafd.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _eaaca.StructParent)
	_dbafd.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _eaaca.StructParents)
	_dbafd.SetIfNotNil("\u004f\u0050\u0049", _eaaca.OPI)
	_dbafd.SetIfNotNil("\u004f\u0043", _eaaca.OC)
	_dbafd.SetIfNotNil("\u004e\u0061\u006d\u0065", _eaaca.Name)
	_dbafd.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dg.MakeInteger(int64(len(_eaaca.Stream))))
	_dgada.Stream = _eaaca.Stream
	return _dgada
}
func (_abbcb *pdfFontSimple) getFontEncoding() (_gcff string, _ecbce map[_bd.CharCode]_bd.GlyphName, _ebbc error) {
	_gcff = "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"
	if _bbada, _dgca := _addac[_abbcb._ecbf]; _dgca {
		_gcff = _bbada
	} else if _abbcb.fontFlags()&_bbfd != 0 {
		for _acgee, _gdab := range _addac {
			if _ga.Contains(_abbcb._ecbf, _acgee) {
				_gcff = _gdab
				break
			}
		}
	}
	if _abbcb.Encoding == nil {
		return _gcff, nil, nil
	}
	switch _cbbe := _abbcb.Encoding.(type) {
	case *_dg.PdfObjectName:
		return string(*_cbbe), nil, nil
	case *_dg.PdfObjectDictionary:
		_agggd, _ffbfc := _dg.GetName(_cbbe.Get("\u0042\u0061\u0073e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
		if _ffbfc {
			_gcff = _agggd.String()
		}
		if _agbag := _cbbe.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"); _agbag != nil {
			_cgbfea, _cbcef := _dg.GetArray(_agbag)
			if !_cbcef {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042a\u0064\u0020\u0066on\u0074\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u003d\u0025\u002b\u0076\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073=\u0025\u0054", _cbbe, _cbbe.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"))
				return "", nil, _dg.ErrTypeError
			}
			_ecbce, _ebbc = _bd.FromFontDifferences(_cgbfea)
		}
		return _gcff, _ecbce, _ebbc
	default:
		_ag.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0072\u0020\u0064\u0069\u0063t\u0020\u0028\u0025\u0054\u0029\u0020\u0025\u0073", _abbcb.Encoding, _abbcb.Encoding)
		return "", nil, _dg.ErrTypeError
	}
}

// NewPdfColorCalGray returns a new CalGray color.
func NewPdfColorCalGray(grayVal float64) *PdfColorCalGray {
	_ggbb := PdfColorCalGray(grayVal)
	return &_ggbb
}

var (
	ErrRequiredAttributeMissing = _bf.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074t\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
	ErrInvalidAttribute         = _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	ErrTypeCheck                = _bf.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	_dgaa                       = _bf.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrEncrypted                = _bf.New("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	ErrNoFont                   = _bf.New("\u0066\u006fn\u0074\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	ErrFontNotSupported         = _ge.Errorf("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u006fn\u0074\u0020\u0028\u0025\u0077\u0029", _dg.ErrNotSupported)
	ErrType1CFontNotSupported   = _ge.Errorf("\u0054y\u0070\u00651\u0043\u0020\u0066o\u006e\u0074\u0073\u0020\u0061\u0072\u0065 \u006e\u006f\u0074\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0028\u0025\u0077\u0029", _dg.ErrNotSupported)
	ErrType3FontNotSupported    = _ge.Errorf("\u0054y\u0070\u00653\u0020\u0066\u006f\u006et\u0073\u0020\u0061r\u0065\u0020\u006e\u006f\u0074\u0020\u0063\u0075\u0072re\u006e\u0074\u006cy\u0020\u0073u\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0028%\u0077\u0029", _dg.ErrNotSupported)
	ErrTTCmapNotSupported       = _ge.Errorf("\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065d\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070e\u0020\u0063\u006d\u0061\u0070\u0020\u0066\u006f\u0072\u006d\u0061\u0074 \u0028\u0025\u0077\u0029", _dg.ErrNotSupported)
	ErrSignNotEnoughSpace       = _ge.Errorf("\u0069\u006e\u0073\u0075\u0066\u0066\u0069c\u0069\u0065\u006et\u0020\u0073\u0070a\u0063\u0065 \u0061\u006c\u006c\u006f\u0063\u0061t\u0065d \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
	ErrSignNoCertificates       = _ge.Errorf("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065tri\u0065ve\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0063\u0068\u0061i\u006e")
)

// Items returns all children outline items.
func (_dfce *Outline) Items() []*OutlineItem { return _dfce.Entries }

// NewPdfAnnotationScreen returns a new screen annotation.
func NewPdfAnnotationScreen() *PdfAnnotationScreen {
	_bead := NewPdfAnnotation()
	_gae := &PdfAnnotationScreen{}
	_gae.PdfAnnotation = _bead
	_bead.SetContext(_gae)
	return _gae
}

// ToPdfObject implements interface PdfModel.
func (_dfc *PdfActionSubmitForm) ToPdfObject() _dg.PdfObject {
	_dfc.PdfAction.ToPdfObject()
	_gea := _dfc._cbd
	_bac := _gea.PdfObject.(*_dg.PdfObjectDictionary)
	_bac.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeSubmitForm)))
	if _dfc.F != nil {
		_bac.Set("\u0046", _dfc.F.ToPdfObject())
	}
	_bac.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _dfc.Fields)
	_bac.SetIfNotNil("\u0046\u006c\u0061g\u0073", _dfc.Flags)
	return _gea
}

// SetForms sets the Acroform for a PDF file.
func (_fbdbb *PdfWriter) SetForms(form *PdfAcroForm) error { _fbdbb._geabe = form; return nil }
func _dfgb(_dedb *XObjectForm) (*PdfRectangle, bool, error) {
	if _eeaed, _cdcd := _dedb.BBox.(*_dg.PdfObjectArray); _cdcd {
		_ddfe, _abcea := NewPdfRectangle(*_eeaed)
		if _abcea != nil {
			return nil, false, _abcea
		}
		if _dbde, _ddca := _dedb.Matrix.(*_dg.PdfObjectArray); _ddca {
			_cgcac, _ebedfg := _dbde.ToFloat64Array()
			if _ebedfg != nil {
				return nil, false, _ebedfg
			}
			_edfa := _fec.IdentityMatrix()
			if len(_cgcac) == 6 {
				_edfa = _fec.NewMatrix(_cgcac[0], _cgcac[1], _cgcac[2], _cgcac[3], _cgcac[4], _cgcac[5])
			}
			_ddfe.Transform(_edfa)
			return _ddfe, true, nil
		}
		return _ddfe, false, nil
	}
	return nil, false, _bf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063e\u0020\u0042\u0042\u006f\u0078\u0020\u0074y\u0070\u0065")
}

// SetCatalogMetadata sets the catalog metadata (XMP) stream object.
func (_ffbgc *PdfWriter) SetCatalogMetadata(meta _dg.PdfObject) error {
	if meta == nil {
		_ffbgc._ecdf.Remove("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
		return nil
	}
	_bbfc, _dgedc := _dg.GetStream(meta)
	if !_dgedc {
		return _bf.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006d\u0065\u0074\u0061\u0064a\u0074\u0061\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0073t\u0072\u0065\u0061\u006d")
	}
	_ffbgc.addObject(_bbfc)
	_ffbgc._ecdf.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _bbfc)
	return nil
}
func (_dcg *PdfReader) newPdfAnnotationMovieFromDict(_dcad *_dg.PdfObjectDictionary) (*PdfAnnotationMovie, error) {
	_afbb := PdfAnnotationMovie{}
	_afbb.T = _dcad.Get("\u0054")
	_afbb.Movie = _dcad.Get("\u004d\u006f\u0076i\u0065")
	_afbb.A = _dcad.Get("\u0041")
	return &_afbb, nil
}

// NewDSS returns a new DSS dictionary.
func NewDSS() *DSS {
	return &DSS{_agdcg: _dg.MakeIndirectObject(_dg.MakeDict()), VRI: map[string]*VRI{}}
}

// ToPdfObject implements interface PdfModel.
func (_deda *PdfAnnotationMovie) ToPdfObject() _dg.PdfObject {
	_deda.PdfAnnotation.ToPdfObject()
	_gaaa := _deda._cdf
	_bffb := _gaaa.PdfObject.(*_dg.PdfObjectDictionary)
	_bffb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u004d\u006f\u0076i\u0065"))
	_bffb.SetIfNotNil("\u0054", _deda.T)
	_bffb.SetIfNotNil("\u004d\u006f\u0076i\u0065", _deda.Movie)
	_bffb.SetIfNotNil("\u0041", _deda.A)
	return _gaaa
}

// AlphaMap performs mapping of alpha data for transformations. Allows custom filtering of alpha data etc.
func (_addgf *Image) AlphaMap(mapFunc AlphaMapFunc) {
	for _aeeaa, _ggdc := range _addgf._dgeb {
		_addgf._dgeb[_aeeaa] = mapFunc(_ggdc)
	}
}

// NewPdfColorCalRGB returns a new CalRBG color.
func NewPdfColorCalRGB(a, b, c float64) *PdfColorCalRGB {
	_gbbb := PdfColorCalRGB{a, b, c}
	return &_gbbb
}

// GetContainingPdfObject implements interface PdfModel.
func (_aa *PdfAction) GetContainingPdfObject() _dg.PdfObject { return _aa._cbd }

// SignatureHandler interface defines the common functionality for PDF signature handlers, which
// need to be capable of validating digital signatures and signing PDF documents.
type SignatureHandler interface {

	// IsApplicable checks if a given signature dictionary `sig` is applicable for the signature handler.
	// For example a signature of type `adbe.pkcs7.detached` might not fit for a rsa.sha1 handler.
	IsApplicable(_fdfda *PdfSignature) bool

	// Validate validates a PDF signature against a given digest (hash) such as that determined
	// for an input file. Returns validation results.
	Validate(_eeagbe *PdfSignature, _efabg Hasher) (SignatureValidationResult, error)

	// InitSignature prepares the signature dictionary for signing. This involves setting all
	// necessary fields, and also allocating sufficient space to the Contents so that the
	// finalized signature can be inserted once the hash is calculated.
	InitSignature(_cgeeg *PdfSignature) error

	// NewDigest creates a new digest/hasher based on the signature dictionary and handler.
	NewDigest(_caccf *PdfSignature) (Hasher, error)

	// Sign receives the hash `digest` (for example hash of an input file), and signs based
	// on the signature dictionary `sig` and applies the signature data to the signature
	// dictionary Contents field.
	Sign(_dgagg *PdfSignature, _bfebg Hasher) error
}

// NewPdfAnnotationStrikeOut returns a new text strikeout annotation.
func NewPdfAnnotationStrikeOut() *PdfAnnotationStrikeOut {
	_fdcd := NewPdfAnnotation()
	_egb := &PdfAnnotationStrikeOut{}
	_egb.PdfAnnotation = _fdcd
	_egb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fdcd.SetContext(_egb)
	return _egb
}

// PdfAnnotationMarkup represents additional fields for mark-up annotations.
// (Section 12.5.6.2 p. 399).
type PdfAnnotationMarkup struct {
	T            _dg.PdfObject
	Popup        *PdfAnnotationPopup
	CA           _dg.PdfObject
	RC           _dg.PdfObject
	CreationDate _dg.PdfObject
	IRT          _dg.PdfObject
	Subj         _dg.PdfObject
	RT           _dg.PdfObject
	IT           _dg.PdfObject
	ExData       _dg.PdfObject
}

// ColorToRGB only converts color used with uncolored patterns (defined in underlying colorspace).  Does not go into the
// pattern objects and convert those.  If that is desired, needs to be done separately.  See for example
// grayscale conversion example in unidoc-examples repo.
func (_gecc *PdfColorspaceSpecialPattern) ColorToRGB(color PdfColor) (PdfColor, error) {
	_eagbg, _dceb := color.(*PdfColorPattern)
	if !_dceb {
		_ag.Log.Debug("\u0043\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0070a\u0074\u0074\u0065\u0072\u006e\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", color)
		return nil, ErrTypeCheck
	}
	if _eagbg.Color == nil {
		return color, nil
	}
	if _gecc.UnderlyingCS == nil {
		return nil, _bf.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	return _gecc.UnderlyingCS.ColorToRGB(_eagbg.Color)
}
func (_abdgf *PdfWriter) writeObjectsInStreams(_bbfbg map[_dg.PdfObject]bool) error {
	for _, _bebcf := range _abdgf._agaba {
		if _bcag := _bbfbg[_bebcf]; _bcag {
			continue
		}
		_deeb := int64(0)
		switch _gbge := _bebcf.(type) {
		case *_dg.PdfIndirectObject:
			_deeb = _gbge.ObjectNumber
		case *_dg.PdfObjectStream:
			_deeb = _gbge.ObjectNumber
		case *_dg.PdfObjectStreams:
			_deeb = _gbge.ObjectNumber
		default:
			_ag.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0055n\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0074\u0079\u0070\u0065 \u0069\u006e\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0054", _bebcf)
			return ErrTypeCheck
		}
		if _abdgf._fadcg != nil && _bebcf != _abdgf._acdag {
			_aeda := _abdgf._fadcg.Encrypt(_bebcf, _deeb, 0)
			if _aeda != nil {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067\u0020(%\u0073\u0029", _aeda)
				return _aeda
			}
		}
		_abdgf.writeObject(int(_deeb), _bebcf)
	}
	return nil
}
func (_abbd *PdfReader) traverseObjectData(_dbff _dg.PdfObject) error {
	return _dg.ResolveReferencesDeep(_dbff, _abbd._addfg)
}

// PdfColorspaceDeviceNAttributes contains additional information about the components of colour space that
// conforming readers may use. Conforming readers need not use the alternateSpace and tintTransform parameters,
// and may instead use a custom blending algorithms, along with other information provided in the attributes
// dictionary if present.
type PdfColorspaceDeviceNAttributes struct {
	Subtype     *_dg.PdfObjectName
	Colorants   _dg.PdfObject
	Process     _dg.PdfObject
	MixingHints _dg.PdfObject
	_ffbg       *_dg.PdfIndirectObject
}

// ToPdfObject implements interface PdfModel.
func (_age *PdfAnnotationRichMedia) ToPdfObject() _dg.PdfObject {
	_age.PdfAnnotation.ToPdfObject()
	_gdfg := _age._cdf
	_cgda := _gdfg.PdfObject.(*_dg.PdfObjectDictionary)
	_cgda.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a"))
	_cgda.SetIfNotNil("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073", _age.RichMediaSettings)
	_cgda.SetIfNotNil("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074", _age.RichMediaContent)
	return _gdfg
}

// GetRevision returns the specific version of the PdfReader for the current Pdf document
func (_gbee *PdfReader) GetRevision(revisionNumber int) (*PdfReader, error) {
	_feeeg := _gbee._baad.GetRevisionNumber()
	if revisionNumber < 0 || revisionNumber > _feeeg {
		return nil, _bf.New("w\u0072\u006f\u006e\u0067 r\u0065v\u0069\u0073\u0069\u006f\u006e \u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if revisionNumber == _feeeg {
		return _gbee, nil
	}
	if _gbee._aggag[revisionNumber] != nil {
		return _gbee._aggag[revisionNumber], nil
	}
	_bbdcad := _gbee
	for _ebbcb := _feeeg - 1; _ebbcb >= revisionNumber; _ebbcb-- {
		_agbd, _abfd := _bbdcad.GetPreviousRevision()
		if _abfd != nil {
			return nil, _abfd
		}
		_gbee._aggag[_ebbcb] = _agbd
		_bbdcad = _agbd
	}
	return _bbdcad, nil
}

// ToPdfObject converts the PdfFont object to its PDF representation.
func (_dcfbb *PdfFont) ToPdfObject() _dg.PdfObject {
	if _dcfbb._cadf == nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0066\u006f\u006e\u0074 \u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return _dg.MakeNull()
	}
	return _dcfbb._cadf.ToPdfObject()
}
func (_ggeb *PdfReader) newPdfSignatureReferenceFromDict(_fadfe *_dg.PdfObjectDictionary) (*PdfSignatureReference, error) {
	if _abdcc, _dfgbb := _ggeb._cadfa.GetModelFromPrimitive(_fadfe).(*PdfSignatureReference); _dfgbb {
		return _abdcc, nil
	}
	_cgcbe := &PdfSignatureReference{_eccgb: _fadfe, Data: _fadfe.Get("\u0044\u0061\u0074\u0061")}
	var _dceedc bool
	_cgcbe.Type, _ = _dg.GetName(_fadfe.Get("\u0054\u0079\u0070\u0065"))
	_cgcbe.TransformMethod, _dceedc = _dg.GetName(_fadfe.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"))
	if !_dceedc {
		_ag.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0072\u0061\u006e\u0073\u0066o\u0072\u006dM\u0065\u0074h\u006f\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020in\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020m\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_cgcbe.TransformParams, _ = _dg.GetDict(_fadfe.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"))
	_cgcbe.DigestMethod, _ = _dg.GetName(_fadfe.Get("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064"))
	return _cgcbe, nil
}
func _dfdb(_gcegf *_dg.PdfObjectDictionary) (*PdfShadingType7, error) {
	_bgeeg := PdfShadingType7{}
	_afdce := _gcegf.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _afdce == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fgfba, _efcda := _afdce.(*_dg.PdfObjectInteger)
	if !_efcda {
		_ag.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _afdce)
		return nil, _dg.ErrTypeError
	}
	_bgeeg.BitsPerCoordinate = _fgfba
	_afdce = _gcegf.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _afdce == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_fgfba, _efcda = _afdce.(*_dg.PdfObjectInteger)
	if !_efcda {
		_ag.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _afdce)
		return nil, _dg.ErrTypeError
	}
	_bgeeg.BitsPerComponent = _fgfba
	_afdce = _gcegf.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _afdce == nil {
		_ag.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fgfba, _efcda = _afdce.(*_dg.PdfObjectInteger)
	if !_efcda {
		_ag.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _afdce)
		return nil, _dg.ErrTypeError
	}
	_bgeeg.BitsPerComponent = _fgfba
	_afdce = _gcegf.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _afdce == nil {
		_ag.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_aebb, _efcda := _afdce.(*_dg.PdfObjectArray)
	if !_efcda {
		_ag.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _afdce)
		return nil, _dg.ErrTypeError
	}
	_bgeeg.Decode = _aebb
	if _aecbf := _gcegf.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _aecbf != nil {
		_bgeeg.Function = []PdfFunction{}
		if _cfgbdb, _gggeea := _aecbf.(*_dg.PdfObjectArray); _gggeea {
			for _, _gaeac := range _cfgbdb.Elements() {
				_bbcd, _bfgfe := _agec(_gaeac)
				if _bfgfe != nil {
					_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bfgfe)
					return nil, _bfgfe
				}
				_bgeeg.Function = append(_bgeeg.Function, _bbcd)
			}
		} else {
			_aaega, _deab := _agec(_aecbf)
			if _deab != nil {
				_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _deab)
				return nil, _deab
			}
			_bgeeg.Function = append(_bgeeg.Function, _aaega)
		}
	}
	return &_bgeeg, nil
}

// NewPdfColorspaceSpecialSeparation returns a new separation color.
func NewPdfColorspaceSpecialSeparation() *PdfColorspaceSpecialSeparation {
	_ebbde := &PdfColorspaceSpecialSeparation{}
	return _ebbde
}

// SetShadingByName sets a shading resource specified by keyName.
func (_cbbcf *PdfPageResources) SetShadingByName(keyName _dg.PdfObjectName, shadingObj _dg.PdfObject) error {
	if _cbbcf.Shading == nil {
		_cbbcf.Shading = _dg.MakeDict()
	}
	_dgddc, _abddg := _dg.GetDict(_cbbcf.Shading)
	if !_abddg {
		return _dg.ErrTypeError
	}
	_dgddc.Set(keyName, shadingObj)
	return nil
}

// SubsetRegistered subsets the font to only the glyphs that have been registered by the encoder.
//
// NOTE: This only works on fonts that support subsetting. For unsupported fonts this is a no-op, although a debug
// message is emitted.  Currently supported fonts are embedded Truetype CID fonts (type 0).
//
// NOTE: Make sure to call this soon before writing (once all needed runes have been registered).
// If using package creator, use its EnableFontSubsetting method instead.
func (_cegc *PdfFont) SubsetRegistered() error {
	switch _eegab := _cegc._cadf.(type) {
	case *pdfFontType0:
		_gegg := _eegab.subsetRegistered()
		if _gegg != nil {
			_ag.Log.Debug("\u0053\u0075b\u0073\u0065\u0074 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _gegg)
			return _gegg
		}
		if _eegab._bedc != nil {
			if _eegab._ggec != nil {
				_eegab._ggec.ToPdfObject()
			}
			_eegab.ToPdfObject()
		}
	default:
		_ag.Log.Debug("F\u006f\u006e\u0074\u0020\u0025\u0054 \u0064\u006f\u0065\u0073\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006fr\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069n\u0067", _eegab)
	}
	return nil
}

// GetContentStreamObjs returns a slice of PDF objects containing the content
// streams of the page.
func (_dccab *PdfPage) GetContentStreamObjs() []_dg.PdfObject {
	if _dccab.Contents == nil {
		return nil
	}
	_cgaee := _dg.TraceToDirectObject(_dccab.Contents)
	if _dgfg, _eegaa := _cgaee.(*_dg.PdfObjectArray); _eegaa {
		return _dgfg.Elements()
	}
	return []_dg.PdfObject{_cgaee}
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_caeba *PdfShadingType4) ToPdfObject() _dg.PdfObject {
	_caeba.PdfShading.ToPdfObject()
	_cefea, _baed := _caeba.getShadingDict()
	if _baed != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _caeba.BitsPerCoordinate != nil {
		_cefea.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _caeba.BitsPerCoordinate)
	}
	if _caeba.BitsPerComponent != nil {
		_cefea.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _caeba.BitsPerComponent)
	}
	if _caeba.BitsPerFlag != nil {
		_cefea.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _caeba.BitsPerFlag)
	}
	if _caeba.Decode != nil {
		_cefea.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _caeba.Decode)
	}
	if _caeba.Function != nil {
		if len(_caeba.Function) == 1 {
			_cefea.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _caeba.Function[0].ToPdfObject())
		} else {
			_bbefg := _dg.MakeArray()
			for _, _bcgfad := range _caeba.Function {
				_bbefg.Append(_bcgfad.ToPdfObject())
			}
			_cefea.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _bbefg)
		}
	}
	return _caeba._bcfbg
}

// PdfInfo holds document information that will overwrite
// document information global variables defined above.
type PdfInfo struct {
	Title        *_dg.PdfObjectString
	Author       *_dg.PdfObjectString
	Subject      *_dg.PdfObjectString
	Keywords     *_dg.PdfObjectString
	Creator      *_dg.PdfObjectString
	Producer     *_dg.PdfObjectString
	CreationDate *PdfDate
	ModifiedDate *PdfDate
	Trapped      *_dg.PdfObjectName
	_dccdg       *_dg.PdfObjectDictionary
}

// NewPdfWriter initializes a new PdfWriter.
func NewPdfWriter() PdfWriter {
	_edcg := PdfWriter{}
	_edcg._fdbfa = map[_dg.PdfObject]struct{}{}
	_edcg._agaba = []_dg.PdfObject{}
	_edcg._ccgade = map[_dg.PdfObject][]*_dg.PdfObjectDictionary{}
	_edcg._cffaa = map[_dg.PdfObject]struct{}{}
	_edcg._efacd.Major = 1
	_edcg._efacd.Minor = 3
	_fdeee := _dg.MakeDict()
	_abag := []struct {
		_gdgd   _dg.PdfObjectName
		_gagfce string
	}{{"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _bbadae()}, {"\u0043r\u0065\u0061\u0074\u006f\u0072", _fgcbf()}, {"\u0041\u0075\u0074\u0068\u006f\u0072", _dgdbe()}, {"\u0053u\u0062\u006a\u0065\u0063\u0074", _aedcd()}, {"\u0054\u0069\u0074l\u0065", _daece()}, {"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _bbece()}}
	for _, _fedbd := range _abag {
		if _fedbd._gagfce != "" {
			_fdeee.Set(_fedbd._gdgd, _dg.MakeString(_fedbd._gagfce))
		}
	}
	if _dcadfb := _bgcbab(); !_dcadfb.IsZero() {
		if _becdg, _dadec := NewPdfDateFromTime(_dcadfb); _dadec == nil {
			_fdeee.Set("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _becdg.ToPdfObject())
		}
	}
	if _abdac := _ffeac(); !_abdac.IsZero() {
		if _agbdd, _ffaab := NewPdfDateFromTime(_abdac); _ffaab == nil {
			_fdeee.Set("\u004do\u0064\u0044\u0061\u0074\u0065", _agbdd.ToPdfObject())
		}
	}
	_fefba := _dg.PdfIndirectObject{}
	_fefba.PdfObject = _fdeee
	_edcg._efbfa = &_fefba
	_edcg.addObject(&_fefba)
	_afabg := _dg.PdfIndirectObject{}
	_gcgcg := _dg.MakeDict()
	_gcgcg.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0043a\u0074\u0061\u006c\u006f\u0067"))
	_afabg.PdfObject = _gcgcg
	_edcg._fadee = &_afabg
	_edcg.addObject(_edcg._fadee)
	_gebfe, _cefaae := _aeegd("\u0077")
	if _cefaae != nil {
		_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cefaae)
	}
	_edcg._gdfbg = _gebfe
	_addfa := _dg.PdfIndirectObject{}
	_egfed := _dg.MakeDict()
	_egfed.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0050\u0061\u0067e\u0073"))
	_bebgd := _dg.PdfObjectArray{}
	_egfed.Set("\u004b\u0069\u0064\u0073", &_bebgd)
	_egfed.Set("\u0043\u006f\u0075n\u0074", _dg.MakeInteger(0))
	_addfa.PdfObject = _egfed
	_edcg._gbgb = &_addfa
	_edcg._fegdf = map[_dg.PdfObject]struct{}{}
	_edcg.addObject(_edcg._gbgb)
	_gcgcg.Set("\u0050\u0061\u0067e\u0073", &_addfa)
	_edcg._ecdf = _gcgcg
	_ag.Log.Trace("\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0025\u0073", _afabg)
	return _edcg
}

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_faceb *PdfWriter) SetOptimizer(optimizer Optimizer) { _faceb._egeac = optimizer }

// PdfAnnotationSquare represents Square annotations.
// (Section 12.5.6.8).
type PdfAnnotationSquare struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _dg.PdfObject
	IC _dg.PdfObject
	BE _dg.PdfObject
	RD _dg.PdfObject
}

// VRI represents a Validation-Related Information dictionary.
// The VRI dictionary contains validation data in the form of
// certificates, OCSP and CRL information, for a single signature.
// See ETSI TS 102 778-4 V1.1.1 for more information.
type VRI struct {
	Cert []*_dg.PdfObjectStream
	OCSP []*_dg.PdfObjectStream
	CRL  []*_dg.PdfObjectStream
	TU   *_dg.PdfObjectString
	TS   *_dg.PdfObjectString
}

// SetType sets the field button's type.  Can be one of:
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
// This sets the field's flag appropriately.
func (_efef *PdfFieldButton) SetType(btype ButtonType) {
	_adbae := uint32(0)
	if _efef.Ff != nil {
		_adbae = uint32(*_efef.Ff)
	}
	switch btype {
	case ButtonTypePush:
		_adbae |= FieldFlagPushbutton.Mask()
	case ButtonTypeRadio:
		_adbae |= FieldFlagRadio.Mask()
	}
	_efef.Ff = _dg.MakeInteger(int64(_adbae))
}

// PdfFontDescriptor specifies metrics and other attributes of a font and can refer to a FontFile
// for embedded fonts.
// 9.8 Font Descriptors (page 281)
type PdfFontDescriptor struct {
	FontName     _dg.PdfObject
	FontFamily   _dg.PdfObject
	FontStretch  _dg.PdfObject
	FontWeight   _dg.PdfObject
	Flags        _dg.PdfObject
	FontBBox     _dg.PdfObject
	ItalicAngle  _dg.PdfObject
	Ascent       _dg.PdfObject
	Descent      _dg.PdfObject
	Leading      _dg.PdfObject
	CapHeight    _dg.PdfObject
	XHeight      _dg.PdfObject
	StemV        _dg.PdfObject
	StemH        _dg.PdfObject
	AvgWidth     _dg.PdfObject
	MaxWidth     _dg.PdfObject
	MissingWidth _dg.PdfObject
	FontFile     _dg.PdfObject
	FontFile2    _dg.PdfObject
	FontFile3    _dg.PdfObject
	CharSet      _dg.PdfObject
	_gcdf        int
	_adggg       float64
	*fontFile
	_gbcg *_bbg.TtfType

	// Additional entries for CIDFonts
	Style  _dg.PdfObject
	Lang   _dg.PdfObject
	FD     _dg.PdfObject
	CIDSet _dg.PdfObject
	_acag  *_dg.PdfIndirectObject
}

// PdfColorspaceCalGray represents CalGray color space.
type PdfColorspaceCalGray struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      float64
	_gcbf      *_dg.PdfIndirectObject
}

// ImageToRGB converts Lab colorspace image to RGB and returns the result.
func (_gab *PdfColorspaceLab) ImageToRGB(img Image) (Image, error) {
	_acfa := func(_acbb float64) float64 {
		if _acbb >= 6.0/29 {
			return _acbb * _acbb * _acbb
		}
		return 108.0 / 841 * (_acbb - 4.0/29.0)
	}
	_feeee := img._gfbb
	if len(_feeee) != 6 {
		_ag.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u002d\u0020\u004c\u0061\u0062\u0020\u0044e\u0063\u006f\u0064\u0065\u0020\u0072\u0061\u006e\u0067e\u0020\u0021\u003d\u0020\u0036\u002e\u002e\u002e\u0020\u0075\u0073\u0065\u0020\u005b0\u0020\u0031\u0030\u0030\u0020\u0061\u006d\u0069\u006e\u0020\u0061\u006d\u0061\u0078\u0020\u0062\u006d\u0069\u006e\u0020\u0062\u006d\u0061\u0078\u005d\u0020\u0064\u0065\u0066\u0061u\u006c\u0074\u0020\u0064\u0065\u0063\u006f\u0064\u0065 \u0061\u0072r\u0061\u0079")
		_feeee = _gab.DecodeArray()
	}
	_cef := _fcd.NewReader(img.getBase())
	_adaf := _fc.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, img._dgeb, img._gfbb)
	_gfca := _fcd.NewWriter(_adaf)
	_dcf := _cg.Pow(2, float64(img.BitsPerComponent)) - 1
	_dbfe := make([]uint32, 3)
	var (
		_dcabf                                             error
		Ls, As, Bs, L, M, N, X, Y, Z, _adgdd, _dcd, _afgef float64
	)
	for {
		_dcabf = _cef.ReadSamples(_dbfe)
		if _dcabf == _cf.EOF {
			break
		} else if _dcabf != nil {
			return img, _dcabf
		}
		Ls = float64(_dbfe[0]) / _dcf
		As = float64(_dbfe[1]) / _dcf
		Bs = float64(_dbfe[2]) / _dcf
		Ls = _fc.LinearInterpolate(Ls, 0.0, 1.0, _feeee[0], _feeee[1])
		As = _fc.LinearInterpolate(As, 0.0, 1.0, _feeee[2], _feeee[3])
		Bs = _fc.LinearInterpolate(Bs, 0.0, 1.0, _feeee[4], _feeee[5])
		L = (Ls+16)/116 + As/500
		M = (Ls + 16) / 116
		N = (Ls+16)/116 - Bs/200
		X = _gab.WhitePoint[0] * _acfa(L)
		Y = _gab.WhitePoint[1] * _acfa(M)
		Z = _gab.WhitePoint[2] * _acfa(N)
		_adgdd = 3.240479*X + -1.537150*Y + -0.498535*Z
		_dcd = -0.969256*X + 1.875992*Y + 0.041556*Z
		_afgef = 0.055648*X + -0.204043*Y + 1.057311*Z
		_adgdd = _cg.Min(_cg.Max(_adgdd, 0), 1.0)
		_dcd = _cg.Min(_cg.Max(_dcd, 0), 1.0)
		_afgef = _cg.Min(_cg.Max(_afgef, 0), 1.0)
		_dbfe[0] = uint32(_adgdd * _dcf)
		_dbfe[1] = uint32(_dcd * _dcf)
		_dbfe[2] = uint32(_afgef * _dcf)
		if _dcabf = _gfca.WriteSamples(_dbfe); _dcabf != nil {
			return img, _dcabf
		}
	}
	return _edcf(&_adaf), nil
}

const (
	BorderEffectNoEffect BorderEffect = iota
	BorderEffectCloudy   BorderEffect = iota
)

// String returns a string that describes `base`.
func (_aaaf fontCommon) String() string {
	return _b.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0073\u007d", _aaaf.coreString())
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a CalRGB device.
func (_adfdc *PdfColorspaceCalRGB) GetNumComponents() int { return 3 }

// ToPdfObject implements interface PdfModel.
// Note: Call the sub-annotation's ToPdfObject to set both the generic and non-generic information.
func (_fgfe *PdfAnnotation) ToPdfObject() _dg.PdfObject {
	_cdfd := _fgfe._cdf
	_abdcf := _cdfd.PdfObject.(*_dg.PdfObjectDictionary)
	_abdcf.Clear()
	_abdcf.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0041\u006e\u006eo\u0074"))
	_abdcf.SetIfNotNil("\u0052\u0065\u0063\u0074", _fgfe.Rect)
	_abdcf.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _fgfe.Contents)
	_abdcf.SetIfNotNil("\u0050", _fgfe.P)
	_abdcf.SetIfNotNil("\u004e\u004d", _fgfe.NM)
	_abdcf.SetIfNotNil("\u004d", _fgfe.M)
	_abdcf.SetIfNotNil("\u0046", _fgfe.F)
	_abdcf.SetIfNotNil("\u0041\u0050", _fgfe.AP)
	_abdcf.SetIfNotNil("\u0041\u0053", _fgfe.AS)
	_abdcf.SetIfNotNil("\u0042\u006f\u0072\u0064\u0065\u0072", _fgfe.Border)
	_abdcf.SetIfNotNil("\u0043", _fgfe.C)
	_abdcf.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _fgfe.StructParent)
	_abdcf.SetIfNotNil("\u004f\u0043", _fgfe.OC)
	return _cdfd
}

// GetPrimitiveFromModel returns the primitive object corresponding to the input `model`.
func (_daad *modelManager) GetPrimitiveFromModel(model PdfModel) _dg.PdfObject {
	_fbacd, _bgcbb := _daad._aaddgg[model]
	if !_bgcbb {
		return nil
	}
	return _fbacd
}

// ReplaceAcroForm replaces the acrobat form. It appends a new form to the Pdf which
// replaces the original AcroForm.
func (_cfba *PdfAppender) ReplaceAcroForm(acroForm *PdfAcroForm) {
	if acroForm != nil {
		_cfba.updateObjectsDeep(acroForm.ToPdfObject(), nil)
	}
	_cfba._afab = acroForm
}
func (_dee *PdfReader) newPdfAnnotationRedactFromDict(_cea *_dg.PdfObjectDictionary) (*PdfAnnotationRedact, error) {
	_fabb := PdfAnnotationRedact{}
	_afcd, _caggb := _dee.newPdfAnnotationMarkupFromDict(_cea)
	if _caggb != nil {
		return nil, _caggb
	}
	_fabb.PdfAnnotationMarkup = _afcd
	_fabb.QuadPoints = _cea.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_fabb.IC = _cea.Get("\u0049\u0043")
	_fabb.RO = _cea.Get("\u0052\u004f")
	_fabb.OverlayText = _cea.Get("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074")
	_fabb.Repeat = _cea.Get("\u0052\u0065\u0070\u0065\u0061\u0074")
	_fabb.DA = _cea.Get("\u0044\u0041")
	_fabb.Q = _cea.Get("\u0051")
	return &_fabb, nil
}
func (_dbec *PdfReader) loadOutlines() (*PdfOutlineTreeNode, error) {
	if _dbec._baad.GetCrypter() != nil && !_dbec._baad.IsAuthenticated() {
		return nil, _b.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_dcada := _dbec._gccfb
	_cfdbe := _dcada.Get("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	if _cfdbe == nil {
		return nil, nil
	}
	_ag.Log.Trace("\u002d\u0048\u0061\u0073\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	_accgf := _dg.ResolveReference(_cfdbe)
	_ag.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0020\u0072\u006f\u006f\u0074\u003a\u0020\u0025\u0076", _accgf)
	if _fdbaa := _dg.IsNullObject(_accgf); _fdbaa {
		_ag.Log.Trace("\u004f\u0075\u0074li\u006e\u0065\u0020\u0072\u006f\u006f\u0074\u0020\u0069s\u0020n\u0075l\u006c \u002d\u0020\u006e\u006f\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
		return nil, nil
	}
	_eddbf, _aedf := _accgf.(*_dg.PdfIndirectObject)
	if !_aedf {
		if _, _gacca := _dg.GetDict(_accgf); !_gacca {
			_ag.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			return nil, nil
		}
		_ag.Log.Debug("\u004f\u0075t\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u0069s\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		_eddbf = _dg.MakeIndirectObject(_accgf)
	}
	_eabba, _aedf := _eddbf.PdfObject.(*_dg.PdfObjectDictionary)
	if !_aedf {
		return nil, _bf.New("\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y")
	}
	_ag.Log.Trace("O\u0075\u0074\u006c\u0069ne\u0020r\u006f\u006f\u0074\u0020\u0064i\u0063\u0074\u003a\u0020\u0025\u0076", _eabba)
	_gcef, _, _bfafe := _dbec.buildOutlineTree(_eddbf, nil, nil, nil)
	if _bfafe != nil {
		return nil, _bfafe
	}
	_ag.Log.Trace("\u0052\u0065\u0073\u0075\u006c\u0074\u0069\u006e\u0067\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065:\u0020\u0025\u0076", _gcef)
	return _gcef, nil
}
func _fecf(_abac []byte) []byte {
	const _ecfdc = 52845
	const _afcbc = 22719
	_ccad := 55665
	for _, _dbbdb := range _abac[:4] {
		_ccad = (int(_dbbdb)+_ccad)*_ecfdc + _afcbc
	}
	_cagea := make([]byte, len(_abac)-4)
	for _gbdg, _eeafa := range _abac[4:] {
		_cagea[_gbdg] = byte(int(_eeafa) ^ _ccad>>8)
		_ccad = (int(_eeafa)+_ccad)*_ecfdc + _afcbc
	}
	return _cagea
}

// AlphaMapFunc represents a alpha mapping function: byte -> byte. Can be used for
// thresholding the alpha channel, i.e. setting all alpha values below threshold to transparent.
type AlphaMapFunc func(_eadca byte) byte

// GetNumComponents returns the number of color components of the underlying
// colorspace device.
func (_ccdc *PdfColorspaceSpecialPattern) GetNumComponents() int {
	return _ccdc.UnderlyingCS.GetNumComponents()
}
func (_gegc fontCommon) isCIDFont() bool {
	if _gegc._bcga == "" {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0069\u0073\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u002e\u0020\u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gegc)
	}
	_cfec := false
	switch _gegc._bcga {
	case "\u0054\u0079\u0070e\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_cfec = true
	}
	_ag.Log.Trace("i\u0073\u0043\u0049\u0044\u0046\u006fn\u0074\u003a\u0020\u0069\u0073\u0043\u0049\u0044\u003d%\u0074\u0020\u0066o\u006et\u003d\u0025\u0073", _cfec, _gegc)
	return _cfec
}

// PdfSignatureReference represents a PDF signature reference dictionary and is used for signing via form signature fields.
// (Section 12.8.1, Table 253 - Entries in a signature reference dictionary p. 469 in PDF32000_2008).
type PdfSignatureReference struct {
	_eccgb          *_dg.PdfObjectDictionary
	Type            *_dg.PdfObjectName
	TransformMethod *_dg.PdfObjectName
	TransformParams _dg.PdfObject
	Data            _dg.PdfObject
	DigestMethod    *_dg.PdfObjectName
}

// NewPdfSignature creates a new PdfSignature object.
func NewPdfSignature(handler SignatureHandler) *PdfSignature {
	_gbagb := &PdfSignature{Type: _dg.MakeName("\u0053\u0069\u0067"), Handler: handler}
	_bcddg := &pdfSignDictionary{PdfObjectDictionary: _dg.MakeDict(), _cgdea: &handler, _cdfca: _gbagb}
	_gbagb._bbda = _dg.MakeIndirectObject(_bcddg)
	return _gbagb
}

// AppendContentStream adds content stream by string.  Appends to the last
// contentstream instance if many.
func (_afaf *PdfPage) AppendContentStream(contentStr string) error {
	_debea, _dbda := _afaf.GetContentStreams()
	if _dbda != nil {
		return _dbda
	}
	if len(_debea) == 0 {
		_debea = []string{contentStr}
		return _afaf.SetContentStreams(_debea, _dg.NewFlateEncoder())
	}
	var _gefe _bc.Buffer
	_gefe.WriteString(_debea[len(_debea)-1])
	_gefe.WriteString("\u000a")
	_gefe.WriteString(contentStr)
	_debea[len(_debea)-1] = _gefe.String()
	return _afaf.SetContentStreams(_debea, _dg.NewFlateEncoder())
}
func (_fbaag *pdfFontSimple) addEncoding() error {
	var (
		_dfgbd string
		_ggfag map[_bd.CharCode]_bd.GlyphName
		_ecgbe _bd.SimpleEncoder
	)
	if _fbaag.Encoder() != nil {
		_bdcfd, _eeeee := _fbaag.Encoder().(_bd.SimpleEncoder)
		if _eeeee && _bdcfd != nil {
			_dfgbd = _bdcfd.BaseName()
		}
	}
	if _fbaag.Encoding != nil {
		_efdbc, _gbcdd, _fbed := _fbaag.getFontEncoding()
		if _fbed != nil {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0073\u0065F\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0053u\u0062t\u0079\u0070\u0065\u003d\u0025\u0071\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003d\u0025\u0073 \u0028\u0025\u0054\u0029\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fbaag._ecbf, _fbaag._bcga, _fbaag.Encoding, _fbaag.Encoding, _fbed)
			return _fbed
		}
		if _efdbc != "" {
			_dfgbd = _efdbc
		}
		_ggfag = _gbcdd
		_ecgbe, _fbed = _bd.NewSimpleTextEncoder(_dfgbd, _ggfag)
		if _fbed != nil {
			return _fbed
		}
	}
	if _ecgbe == nil {
		_fgdcd := _fbaag._ccfb
		if _fgdcd != nil {
			switch _fbaag._bcga {
			case "\u0054\u0079\u0070e\u0031":
				if _fgdcd.fontFile != nil && _fgdcd.fontFile._beaeb != nil {
					_ag.Log.Debug("\u0055\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006et\u0046\u0069\u006c\u0065")
					_ecgbe = _fgdcd.fontFile._beaeb
				}
			case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
				if _fgdcd._gbcg != nil {
					_ag.Log.Debug("\u0055s\u0069n\u0067\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032")
					_ddbg, _bggeg := _fgdcd._gbcg.MakeEncoder()
					if _bggeg == nil {
						_ecgbe = _ddbg
					}
					if _fbaag._ecfb == nil {
						_fbaag._ecfb = _fgdcd._gbcg.MakeToUnicode()
					}
				}
			}
		}
	}
	if _ecgbe != nil {
		if _ggfag != nil {
			_ag.Log.Trace("\u0064\u0069\u0066fe\u0072\u0065\u006e\u0063\u0065\u0073\u003d\u0025\u002b\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _ggfag, _fbaag.baseFields())
			_ecgbe = _bd.ApplyDifferences(_ecgbe, _ggfag)
		}
		_fbaag.SetEncoder(_ecgbe)
	}
	return nil
}

// CustomKeys returns all custom info keys as list.
func (_deaa *PdfInfo) CustomKeys() []string {
	if _deaa._dccdg == nil {
		return nil
	}
	_eccgc := make([]string, len(_deaa._dccdg.Keys()))
	for _, _dfaf := range _deaa._dccdg.Keys() {
		_eccgc = append(_eccgc, _dfaf.String())
	}
	return _eccgc
}
func (_bgbfd *PdfReader) newPdfFieldFromIndirectObject(_acfe *_dg.PdfIndirectObject, _baecf *PdfField) (*PdfField, error) {
	if _fgcc, _bdbde := _bgbfd._cadfa.GetModelFromPrimitive(_acfe).(*PdfField); _bdbde {
		return _fgcc, nil
	}
	_eabc, _gabf := _dg.GetDict(_acfe)
	if !_gabf {
		return nil, _b.Errorf("\u0050\u0064f\u0046\u0069\u0065\u006c\u0064 \u0069\u006e\u0064\u0069\u0072e\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_egbfb := NewPdfField()
	_egbfb._egce = _acfe
	_egbfb._egce.PdfObject = _eabc
	if _dagea, _cdff := _dg.GetName(_eabc.Get("\u0046\u0054")); _cdff {
		_egbfb.FT = _dagea
	}
	if _baecf != nil {
		_egbfb.Parent = _baecf
	}
	_egbfb.T, _ = _eabc.Get("\u0054").(*_dg.PdfObjectString)
	_egbfb.TU, _ = _eabc.Get("\u0054\u0055").(*_dg.PdfObjectString)
	_egbfb.TM, _ = _eabc.Get("\u0054\u004d").(*_dg.PdfObjectString)
	_egbfb.Ff, _ = _eabc.Get("\u0046\u0066").(*_dg.PdfObjectInteger)
	_egbfb.V = _eabc.Get("\u0056")
	_egbfb.DV = _eabc.Get("\u0044\u0056")
	_egbfb.AA = _eabc.Get("\u0041\u0041")
	if DA := _eabc.Get("\u0044\u0041"); DA != nil {
		DA, _ := _dg.GetString(DA)
		_egbfb.VariableText = &VariableText{DA: DA}
		Q, _ := _eabc.Get("\u0051").(*_dg.PdfObjectInteger)
		DS, _ := _eabc.Get("\u0044\u0053").(*_dg.PdfObjectString)
		RV := _eabc.Get("\u0052\u0056")
		_egbfb.VariableText.Q = Q
		_egbfb.VariableText.DS = DS
		_egbfb.VariableText.RV = RV
	}
	_fdae := _egbfb.FT
	if _fdae == nil && _baecf != nil {
		_fdae = _baecf.FT
	}
	if _fdae != nil {
		switch *_fdae {
		case "\u0054\u0078":
			_ggff, _eedgc := _gacf(_eabc)
			if _eedgc != nil {
				return nil, _eedgc
			}
			_ggff.PdfField = _egbfb
			_egbfb._bdfg = _ggff
		case "\u0043\u0068":
			_gacd, _ddded := _aaaa(_eabc)
			if _ddded != nil {
				return nil, _ddded
			}
			_gacd.PdfField = _egbfb
			_egbfb._bdfg = _gacd
		case "\u0042\u0074\u006e":
			_dcdb, _cgcf := _cdfe(_eabc)
			if _cgcf != nil {
				return nil, _cgcf
			}
			_dcdb.PdfField = _egbfb
			_egbfb._bdfg = _dcdb
		case "\u0053\u0069\u0067":
			_ffacg, _abcg := _bgbfd.newPdfFieldSignatureFromDict(_eabc)
			if _abcg != nil {
				return nil, _abcg
			}
			_ffacg.PdfField = _egbfb
			_egbfb._bdfg = _ffacg
		default:
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", *_egbfb.FT)
			return nil, _bf.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079p\u0065")
		}
	}
	if _adda, _fefd := _dg.GetName(_eabc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fefd {
		if *_adda == "\u0057\u0069\u0064\u0067\u0065\u0074" {
			_cceb, _gbbcb := _bgbfd.newPdfAnnotationFromIndirectObject(_acfe)
			if _gbbcb != nil {
				return nil, _gbbcb
			}
			_ababd, _gedb := _cceb.GetContext().(*PdfAnnotationWidget)
			if !_gedb {
				return nil, _bf.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0067e\u0074 \u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006fn")
			}
			_ababd._cgg = _egbfb
			_ababd.Parent = _egbfb._egce
			_egbfb.Annotations = append(_egbfb.Annotations, _ababd)
			return _egbfb, nil
		}
	}
	_edfec := true
	if _cgaf, _afdda := _dg.GetArray(_eabc.Get("\u004b\u0069\u0064\u0073")); _afdda {
		_ddcbc := make([]*_dg.PdfIndirectObject, 0, _cgaf.Len())
		for _, _aagb := range _cgaf.Elements() {
			_adfcc, _bega := _dg.GetIndirect(_aagb)
			if !_bega {
				_gggd, _dbccg := _dg.GetStream(_aagb)
				if _dbccg && _gggd.PdfObjectDictionary != nil {
					_gedc, _egdfe := _dg.GetNameVal(_gggd.Get("\u0054\u0079\u0070\u0065"))
					if _egdfe && _gedc == "\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061" {
						_ag.Log.Debug("E\u0052RO\u0052:\u0020f\u006f\u0072\u006d\u0020\u0066i\u0065\u006c\u0064 \u004b\u0069\u0064\u0073\u0020a\u0072\u0072\u0061y\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0069n\u0076\u0061\u006cid \u004d\u0065\u0074\u0061\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e")
						continue
					}
				}
				return nil, _bf.New("n\u006f\u0074\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0028\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065\u006cd\u0029")
			}
			_fdde, _eeecf := _dg.GetDict(_adfcc)
			if !_eeecf {
				return nil, ErrTypeCheck
			}
			if _edfec {
				_edfec = !_gef(_fdde)
			}
			_ddcbc = append(_ddcbc, _adfcc)
		}
		for _, _afgd := range _ddcbc {
			if _edfec {
				_ggbg, _ebacd := _bgbfd.newPdfAnnotationFromIndirectObject(_afgd)
				if _ebacd != nil {
					_ag.Log.Debug("\u0045r\u0072\u006fr\u0020\u006c\u006fa\u0064\u0069\u006e\u0067\u0020\u0077\u0069d\u0067\u0065\u0074\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0072 \u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _ebacd)
					return nil, _ebacd
				}
				_aaged, _gfgaf := _ggbg._egcg.(*PdfAnnotationWidget)
				if !_gfgaf {
					return nil, ErrTypeCheck
				}
				_aaged._cgg = _egbfb
				_egbfb.Annotations = append(_egbfb.Annotations, _aaged)
			} else {
				_beggg, _bcbc := _bgbfd.newPdfFieldFromIndirectObject(_afgd, _egbfb)
				if _bcbc != nil {
					_ag.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0068\u0069\u006c\u0064\u0020\u0066\u0069\u0065\u006c\u0064: \u0025\u0076", _bcbc)
					return nil, _bcbc
				}
				_egbfb.Kids = append(_egbfb.Kids, _beggg)
			}
		}
	}
	return _egbfb, nil
}

// ToPdfObject returns a *PdfIndirectObject containing a *PdfObjectArray representation of the DeviceN colorspace.
/*
	Format: [/DeviceN names alternateSpace tintTransform]
	    or: [/DeviceN names alternateSpace tintTransform attributes]
*/
func (_fdbd *PdfColorspaceDeviceN) ToPdfObject() _dg.PdfObject {
	_gggef := _dg.MakeArray(_dg.MakeName("\u0044e\u0076\u0069\u0063\u0065\u004e"))
	_gggef.Append(_fdbd.ColorantNames)
	_gggef.Append(_fdbd.AlternateSpace.ToPdfObject())
	_gggef.Append(_fdbd.TintTransform.ToPdfObject())
	if _fdbd.Attributes != nil {
		_gggef.Append(_fdbd.Attributes.ToPdfObject())
	}
	if _fdbd._effag != nil {
		_fdbd._effag.PdfObject = _gggef
		return _fdbd._effag
	}
	return _gggef
}

// GetContainingPdfObject returns the container of the PdfAcroForm (indirect object).
func (_dbcec *PdfAcroForm) GetContainingPdfObject() _dg.PdfObject { return _dbcec._bebfe }

type fontCommon struct {
	_ecbf  string
	_bcga  string
	_cefg  string
	_ebbff _dg.PdfObject
	_ecfb  *_ff.CMap
	_ccfb  *PdfFontDescriptor
	_bgggd int64
}

// ToPdfObject implements interface PdfModel.
func (_cgcb *PdfBorderStyle) ToPdfObject() _dg.PdfObject {
	_acgc := _dg.MakeDict()
	if _cgcb._cbcb != nil {
		if _gfad, _gdcca := _cgcb._cbcb.(*_dg.PdfIndirectObject); _gdcca {
			_gfad.PdfObject = _acgc
		}
	}
	_acgc.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0042\u006f\u0072\u0064\u0065\u0072"))
	if _cgcb.W != nil {
		_acgc.Set("\u0057", _dg.MakeFloat(*_cgcb.W))
	}
	if _cgcb.S != nil {
		_acgc.Set("\u0053", _dg.MakeName(_cgcb.S.GetPdfName()))
	}
	if _cgcb.D != nil {
		_acgc.Set("\u0044", _dg.MakeArrayFromIntegers(*_cgcb.D))
	}
	if _cgcb._cbcb != nil {
		return _cgcb._cbcb
	}
	return _acgc
}

// NewXObjectImageFromStream builds the image xobject from a stream object.
// An image dictionary is the dictionary portion of a stream object representing an image XObject.
func NewXObjectImageFromStream(stream *_dg.PdfObjectStream) (*XObjectImage, error) {
	_gacgfc := &XObjectImage{}
	_gacgfc._abfb = stream
	_ccbeg := *(stream.PdfObjectDictionary)
	_eecc, _ddabcc := _dg.NewEncoderFromStream(stream)
	if _ddabcc != nil {
		return nil, _ddabcc
	}
	_gacgfc.Filter = _eecc
	if _ffdcde := _dg.TraceToDirectObject(_ccbeg.Get("\u0057\u0069\u0064t\u0068")); _ffdcde != nil {
		_gdbc, _cead := _ffdcde.(*_dg.PdfObjectInteger)
		if !_cead {
			return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_ceabce := int64(*_gdbc)
		_gacgfc.Width = &_ceabce
	} else {
		return nil, _bf.New("\u0077\u0069\u0064\u0074\u0068\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _cccac := _dg.TraceToDirectObject(_ccbeg.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _cccac != nil {
		_egaad, _caadd := _cccac.(*_dg.PdfObjectInteger)
		if !_caadd {
			return nil, _bf.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_gdac := int64(*_egaad)
		_gacgfc.Height = &_gdac
	} else {
		return nil, _bf.New("\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	if _aedg := _dg.TraceToDirectObject(_ccbeg.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _aedg != nil {
		_cacf, _cfbef := NewPdfColorspaceFromPdfObject(_aedg)
		if _cfbef != nil {
			return nil, _cfbef
		}
		_gacgfc.ColorSpace = _cacf
	} else {
		_ag.Log.Debug("\u0058O\u0062\u006a\u0065c\u0074\u0020\u0049m\u0061ge\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067 1\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065n\u0074\u0020\u002d\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047r\u0061\u0079")
		_gacgfc.ColorSpace = NewPdfColorspaceDeviceGray()
	}
	if _ecggf := _dg.TraceToDirectObject(_ccbeg.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _ecggf != nil {
		_fgaeg, _cdcdfb := _ecggf.(*_dg.PdfObjectInteger)
		if !_cdcdfb {
			return nil, _bf.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_cbafa := int64(*_fgaeg)
		_gacgfc.BitsPerComponent = &_cbafa
	}
	_gacgfc.Intent = _ccbeg.Get("\u0049\u006e\u0074\u0065\u006e\u0074")
	_gacgfc.ImageMask = _ccbeg.Get("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k")
	_gacgfc.Mask = _ccbeg.Get("\u004d\u0061\u0073\u006b")
	_gacgfc.Decode = _ccbeg.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	_gacgfc.Interpolate = _ccbeg.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065")
	_gacgfc.Alternatives = _ccbeg.Get("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073")
	_gacgfc.SMask = _ccbeg.Get("\u0053\u004d\u0061s\u006b")
	_gacgfc.SMaskInData = _ccbeg.Get("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061")
	_gacgfc.Matte = _ccbeg.Get("\u004d\u0061\u0074t\u0065")
	_gacgfc.Name = _ccbeg.Get("\u004e\u0061\u006d\u0065")
	_gacgfc.StructParent = _ccbeg.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_gacgfc.ID = _ccbeg.Get("\u0049\u0044")
	_gacgfc.OPI = _ccbeg.Get("\u004f\u0050\u0049")
	_gacgfc.Metadata = _ccbeg.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	_gacgfc.OC = _ccbeg.Get("\u004f\u0043")
	_gacgfc.Stream = stream.Stream
	return _gacgfc, nil
}
func _baac(_fbbb _dg.PdfObject) (*PdfColorspaceSpecialIndexed, error) {
	_ceda := NewPdfColorspaceSpecialIndexed()
	if _gffb, _bbbe := _fbbb.(*_dg.PdfIndirectObject); _bbbe {
		_ceda._fffd = _gffb
	}
	_fbbb = _dg.TraceToDirectObject(_fbbb)
	_edfd, _bdbf := _fbbb.(*_dg.PdfObjectArray)
	if !_bdbf {
		return nil, _b.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _edfd.Len() != 4 {
		return nil, _b.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_fbbb = _edfd.Get(0)
	_acac, _bdbf := _fbbb.(*_dg.PdfObjectName)
	if !_bdbf {
		return nil, _b.Errorf("\u0069n\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
	}
	if *_acac != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _b.Errorf("\u0069\u006e\u0064\u0065xe\u0064\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u006e\u0061m\u0065")
	}
	_fbbb = _edfd.Get(1)
	_egbf, _acgcf := DetermineColorspaceNameFromPdfObject(_fbbb)
	if _acgcf != nil {
		return nil, _acgcf
	}
	if _egbf == "\u0049n\u0064\u0065\u0078\u0065\u0064" || _egbf == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
		_ag.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0049\u006e\u0064e\u0078\u0065\u0064\u002f\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043S\u0020\u0061\u0073\u0020\u0062\u0061\u0073\u0065\u0020\u0028\u0025v\u0029", _egbf)
		return nil, _dgaa
	}
	_caabf, _acgcf := NewPdfColorspaceFromPdfObject(_fbbb)
	if _acgcf != nil {
		return nil, _acgcf
	}
	_ceda.Base = _caabf
	_fbbb = _edfd.Get(2)
	_dgba, _acgcf := _dg.GetNumberAsInt64(_fbbb)
	if _acgcf != nil {
		return nil, _acgcf
	}
	if _dgba > 255 {
		return nil, _b.Errorf("\u0069n\u0064\u0065\u0078\u0065d\u0020\u0043\u0053\u003a\u0020I\u006ev\u0061l\u0069\u0064\u0020\u0068\u0069\u0076\u0061l")
	}
	_ceda.HiVal = int(_dgba)
	_fbbb = _edfd.Get(3)
	_ceda.Lookup = _fbbb
	_fbbb = _dg.TraceToDirectObject(_fbbb)
	var _edaa []byte
	if _dffcb, _gadbd := _fbbb.(*_dg.PdfObjectString); _gadbd {
		_edaa = _dffcb.Bytes()
		_ag.Log.Trace("\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0020\u0064", _edaa)
	} else if _cbge, _accg := _fbbb.(*_dg.PdfObjectStream); _accg {
		_ag.Log.Trace("\u0049n\u0064e\u0078\u0065\u0064\u0020\u0073t\u0072\u0065a\u006d\u003a\u0020\u0025\u0073", _fbbb.String())
		_ag.Log.Trace("\u0045\u006e\u0063\u006fde\u0064\u0020\u0028\u0025\u0064\u0029\u0020\u003a\u0020\u0025\u0023\u0020\u0078", len(_cbge.Stream), _cbge.Stream)
		_cbfg, _gdagg := _dg.DecodeStream(_cbge)
		if _gdagg != nil {
			return nil, _gdagg
		}
		_ag.Log.Trace("\u0044e\u0063o\u0064\u0065\u0064\u0020\u0028%\u0064\u0029 \u003a\u0020\u0025\u0020\u0058", len(_cbfg), _cbfg)
		_edaa = _cbfg
	} else {
		_ag.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _fbbb)
		return nil, _b.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076a\u006ci\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if len(_edaa) < _ceda.Base.GetNumComponents()*(_ceda.HiVal+1) {
		_ag.Log.Debug("\u0050\u0044\u0046\u0020\u0049\u006e\u0063o\u006d\u0070\u0061t\u0069\u0062\u0069\u006ci\u0074\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074\u006f\u006f\u0020\u0073\u0068\u006f\u0072\u0074")
		_ag.Log.Debug("\u0046\u0061i\u006c\u002c\u0020\u006c\u0065\u006e\u0028\u0064\u0061\u0074\u0061\u0029\u003a\u0020\u0025\u0064\u002c\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u002c\u0020\u0068\u0069\u0056\u0061\u006c\u003a\u0020\u0025\u0064", len(_edaa), _ceda.Base.GetNumComponents(), _ceda.HiVal)
	} else {
		_edaa = _edaa[:_ceda.Base.GetNumComponents()*(_ceda.HiVal+1)]
	}
	_ceda._bcea = _edaa
	return _ceda, nil
}

// SetDate sets the `M` field of the signature.
func (_efbd *PdfSignature) SetDate(date _a.Time, format string) {
	if format == "" {
		format = "\u0044\u003a\u003200\u0036\u0030\u0031\u0030\u0032\u0031\u0035\u0030\u0034\u0030\u0035\u002d\u0030\u0037\u0027\u0030\u0030\u0027"
	}
	_efbd.M = _dg.MakeString(date.Format(format))
}

// PdfShadingType7 is a Tensor-product patch mesh.
type PdfShadingType7 struct {
	*PdfShading
	BitsPerCoordinate *_dg.PdfObjectInteger
	BitsPerComponent  *_dg.PdfObjectInteger
	BitsPerFlag       *_dg.PdfObjectInteger
	Decode            *_dg.PdfObjectArray
	Function          []PdfFunction
}

func (_cddb *PdfAppender) mergeResources(_ccda, _cebd _dg.PdfObject, _ceed map[_dg.PdfObjectName]_dg.PdfObjectName) _dg.PdfObject {
	if _cebd == nil && _ccda == nil {
		return nil
	}
	if _cebd == nil {
		return _ccda
	}
	_fbgbb, _dede := _dg.GetDict(_cebd)
	if !_dede {
		return _ccda
	}
	if _ccda == nil {
		_gcc := _dg.MakeDict()
		_gcc.Merge(_fbgbb)
		return _cebd
	}
	_edce, _dede := _dg.GetDict(_ccda)
	if !_dede {
		_ag.Log.Error("\u0045\u0072\u0072or\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		_edce = _dg.MakeDict()
	}
	for _, _egcb := range _fbgbb.Keys() {
		if _dfgaf, _gagf := _ceed[_egcb]; _gagf {
			_edce.Set(_dfgaf, _fbgbb.Get(_egcb))
		} else {
			_edce.Set(_egcb, _fbgbb.Get(_egcb))
		}
	}
	return _edce
}

// GetFontByName gets the font specified by keyName. Returns the PdfObject which
// the entry refers to. Returns a bool value indicating whether or not the entry was found.
func (_aeadg *PdfPageResources) GetFontByName(keyName _dg.PdfObjectName) (_dg.PdfObject, bool) {
	if _aeadg.Font == nil {
		return nil, false
	}
	_fbbgb, _gfagc := _dg.TraceToDirectObject(_aeadg.Font).(*_dg.PdfObjectDictionary)
	if !_gfagc {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _dg.TraceToDirectObject(_aeadg.Font))
		return nil, false
	}
	if _fgbgf := _fbbgb.Get(keyName); _fgbgf != nil {
		return _fgbgf, true
	}
	return nil, false
}
func _deg(_cgae _dg.PdfObject) (*PdfFontDescriptor, error) {
	_gbbf := &PdfFontDescriptor{}
	_cgae = _dg.ResolveReference(_cgae)
	if _bcbg, _cbdb := _cgae.(*_dg.PdfIndirectObject); _cbdb {
		_gbbf._acag = _bcbg
		_cgae = _bcbg.PdfObject
	}
	_baeb, _bcggg := _dg.GetDict(_cgae)
	if !_bcggg {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0044\u0065\u0073c\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _cgae)
		return nil, _dg.ErrTypeError
	}
	if _gagfc := _baeb.Get("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"); _gagfc != nil {
		_gbbf.FontName = _gagfc
	} else {
		_ag.Log.Debug("\u0049n\u0063\u006fm\u0070\u0061\u0074\u0069b\u0069\u006c\u0069t\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u004e\u0061me\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006di\u0073\u0073i\u006e\u0067")
	}
	_beafb, _ := _dg.GetName(_gbbf.FontName)
	if _fcac := _baeb.Get("\u0054\u0079\u0070\u0065"); _fcac != nil {
		_gdefe, _beadd := _fcac.(*_dg.PdfObjectName)
		if !_beadd || string(*_gdefe) != "\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072" {
			_ag.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072i\u0070t\u006f\u0072\u0020\u0054y\u0070\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0054\u0029\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _fcac, _beafb, _gbbf.FontName)
		}
	} else {
		_ag.Log.Trace("\u0049\u006ec\u006f\u006d\u0070\u0061\u0074i\u0062\u0069\u006c\u0069\u0074y\u003a\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _beafb, _gbbf.FontName)
	}
	_gbbf.FontFamily = _baeb.Get("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079")
	_gbbf.FontStretch = _baeb.Get("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068")
	_gbbf.FontWeight = _baeb.Get("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074")
	_gbbf.Flags = _baeb.Get("\u0046\u006c\u0061g\u0073")
	_gbbf.FontBBox = _baeb.Get("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078")
	_gbbf.ItalicAngle = _baeb.Get("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065")
	_gbbf.Ascent = _baeb.Get("\u0041\u0073\u0063\u0065\u006e\u0074")
	_gbbf.Descent = _baeb.Get("\u0044e\u0073\u0063\u0065\u006e\u0074")
	_gbbf.Leading = _baeb.Get("\u004ce\u0061\u0064\u0069\u006e\u0067")
	_gbbf.CapHeight = _baeb.Get("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t")
	_gbbf.XHeight = _baeb.Get("\u0058H\u0065\u0069\u0067\u0068\u0074")
	_gbbf.StemV = _baeb.Get("\u0053\u0074\u0065m\u0056")
	_gbbf.StemH = _baeb.Get("\u0053\u0074\u0065m\u0048")
	_gbbf.AvgWidth = _baeb.Get("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068")
	_gbbf.MaxWidth = _baeb.Get("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068")
	_gbbf.MissingWidth = _baeb.Get("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068")
	_gbbf.FontFile = _baeb.Get("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065")
	_gbbf.FontFile2 = _baeb.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652")
	_gbbf.FontFile3 = _baeb.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653")
	_gbbf.CharSet = _baeb.Get("\u0043h\u0061\u0072\u0053\u0065\u0074")
	_gbbf.Style = _baeb.Get("\u0053\u0074\u0079l\u0065")
	_gbbf.Lang = _baeb.Get("\u004c\u0061\u006e\u0067")
	_gbbf.FD = _baeb.Get("\u0046\u0044")
	_gbbf.CIDSet = _baeb.Get("\u0043\u0049\u0044\u0053\u0065\u0074")
	if _gbbf.Flags != nil {
		if _bcgb, _fedg := _dg.GetIntVal(_gbbf.Flags); _fedg {
			_gbbf._gcdf = _bcgb
		}
	}
	if _gbbf.MissingWidth != nil {
		if _fbca, _adfac := _dg.GetNumberAsFloat(_gbbf.MissingWidth); _adfac == nil {
			_gbbf._adggg = _fbca
		}
	}
	if _gbbf.FontFile != nil {
		_bffec, _faee := _bedbb(_gbbf.FontFile)
		if _faee != nil {
			return _gbbf, _faee
		}
		_ag.Log.Trace("f\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u003d\u0025\u0073", _bffec)
		_gbbf.fontFile = _bffec
	}
	if _gbbf.FontFile2 != nil {
		_fedd, _bgdfa := _bbg.NewFontFile2FromPdfObject(_gbbf.FontFile2)
		if _bgdfa != nil {
			return _gbbf, _bgdfa
		}
		_ag.Log.Trace("\u0066\u006f\u006et\u0046\u0069\u006c\u0065\u0032\u003d\u0025\u0073", _fedd.String())
		_gbbf._gbcg = &_fedd
	}
	return _gbbf, nil
}

// NewPdfAnnotationWatermark returns a new watermark annotation.
func NewPdfAnnotationWatermark() *PdfAnnotationWatermark {
	_eaeg := NewPdfAnnotation()
	_gec := &PdfAnnotationWatermark{}
	_gec.PdfAnnotation = _eaeg
	_eaeg.SetContext(_gec)
	return _gec
}

// IsCID returns true if the underlying font is CID.
func (_gadd *PdfFont) IsCID() bool { return _gadd.baseFields().isCIDFont() }

// ColorToRGB converts a CMYK32 color to an RGB color.
func (_afbf *PdfColorspaceDeviceCMYK) ColorToRGB(color PdfColor) (PdfColor, error) {
	_dcadde, _bfef := color.(*PdfColorDeviceCMYK)
	if !_bfef {
		_ag.Log.Debug("I\u006e\u0070\u0075\u0074\u0020\u0063o\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0064e\u0076\u0069\u0063e\u0020c\u006d\u0079\u006b")
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_aede := _dcadde.C()
	_bgcdb := _dcadde.M()
	_dfde := _dcadde.Y()
	_afabe := _dcadde.K()
	_aede = _aede*(1-_afabe) + _afabe
	_bgcdb = _bgcdb*(1-_afabe) + _afabe
	_dfde = _dfde*(1-_afabe) + _afabe
	_fgddc := 1 - _aede
	_cgdaf := 1 - _bgcdb
	_ggb := 1 - _dfde
	return NewPdfColorDeviceRGB(_fgddc, _cgdaf, _ggb), nil
}
func (_ebbf *PdfColorspaceSpecialIndexed) String() string {
	return "\u0049n\u0064\u0065\u0078\u0065\u0064"
}

// AddCerts adds certificates to DSS.
func (_cfbb *DSS) AddCerts(certs [][]byte) ([]*_dg.PdfObjectStream, error) {
	return _cfbb.add(&_cfbb.Certs, _cfbb._agbg, certs)
}

// GetFontDescriptor returns the font descriptor for `font`.
func (_ageba PdfFont) GetFontDescriptor() (*PdfFontDescriptor, error) {
	return _ageba._cadf.getFontDescriptor(), nil
}

// ToPdfObject implements interface PdfModel.
func (_cfc *PdfActionGoTo) ToPdfObject() _dg.PdfObject {
	_cfc.PdfAction.ToPdfObject()
	_bdd := _cfc._cbd
	_bbb := _bdd.PdfObject.(*_dg.PdfObjectDictionary)
	_bbb.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeGoTo)))
	_bbb.SetIfNotNil("\u0044", _cfc.D)
	return _bdd
}

// ImageToRGB returns the passed in image. Method exists in order to satisfy
// the PdfColorspace interface.
func (_dgb *PdfColorspaceDeviceRGB) ImageToRGB(img Image) (Image, error) { return img, nil }

// DecodeArray returns the range of color component values in CalGray colorspace.
func (_gdfge *PdfColorspaceCalGray) DecodeArray() []float64 { return []float64{0.0, 1.0} }

// NewStandard14FontWithEncoding returns the standard 14 font named `basefont` as a *PdfFont and
// a TextEncoder that encodes all the runes in `alphabet`, or an error if this is not possible.
// An error can occur if `basefont` is not one the standard 14 font names.
func NewStandard14FontWithEncoding(basefont StdFontName, alphabet map[rune]int) (*PdfFont, _bd.SimpleEncoder, error) {
	_dbbc, _gbeb := _eece(basefont)
	if _gbeb != nil {
		return nil, nil, _gbeb
	}
	_eaaef, _ebefc := _dbbc.Encoder().(_bd.SimpleEncoder)
	if !_ebefc {
		return nil, nil, _b.Errorf("\u006f\u006e\u006c\u0079\u0020s\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u002c\u0020\u0067\u006f\u0074\u0020\u0025\u0054", _dbbc.Encoder())
	}
	_dbccb := make(map[rune]_bd.GlyphName)
	for _gcgdc := range alphabet {
		if _, _bfaa := _eaaef.RuneToCharcode(_gcgdc); !_bfaa {
			_, _ffbc := _dbbc._bfdee.Read(_gcgdc)
			if !_ffbc {
				_ag.Log.Trace("r\u0075\u006e\u0065\u0020\u0025\u0023x\u003d\u0025\u0071\u0020\u006e\u006f\u0074\u0020\u0069n\u0020\u0074\u0068e\u0020f\u006f\u006e\u0074", _gcgdc, _gcgdc)
				continue
			}
			_dfdc, _ffbc := _bd.RuneToGlyph(_gcgdc)
			if !_ffbc {
				_ag.Log.Debug("\u006eo\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u0020\u0025\u0023\u0078\u003d\u0025\u0071", _gcgdc, _gcgdc)
				continue
			}
			if len(_dbccb) >= 255 {
				return nil, nil, _bf.New("\u0074\u006f\u006f\u0020\u006d\u0061\u006e\u0079\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020f\u006f\u0072\u0020\u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067")
			}
			_dbccb[_gcgdc] = _dfdc
		}
	}
	var (
		_fcbbb []_bd.CharCode
		_gade  []_bd.CharCode
	)
	for _fgca := _bd.CharCode(1); _fgca <= 0xff; _fgca++ {
		_aagbf, _bgea := _eaaef.CharcodeToRune(_fgca)
		if !_bgea {
			_fcbbb = append(_fcbbb, _fgca)
			continue
		}
		if _, _bgea = alphabet[_aagbf]; !_bgea {
			_gade = append(_gade, _fgca)
		}
	}
	_edab := append(_fcbbb, _gade...)
	if len(_edab) < len(_dbccb) {
		return nil, nil, _b.Errorf("n\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u006ec\u006f\u0064\u0065\u0020\u0025\u0064\u0020ru\u006e\u0065\u0073\u002c \u0062\u0075\u0074\u0020\u0068\u0061\u0076\u0065\u0020on\u006c\u0079 \u0025\u0064\u0020\u0073\u006c\u006f\u0074\u0073", len(_dbccb), len(_edab))
	}
	_gabff := make([]rune, 0, len(_dbccb))
	for _ecfd := range _dbccb {
		_gabff = append(_gabff, _ecfd)
	}
	_gc.Slice(_gabff, func(_bdace, _fafdc int) bool { return _gabff[_bdace] < _gabff[_fafdc] })
	_cgfe := make(map[_bd.CharCode]_bd.GlyphName, len(_gabff))
	for _, _cgdf := range _gabff {
		_cegd := _edab[0]
		_edab = _edab[1:]
		_cgfe[_cegd] = _dbccb[_cgdf]
	}
	_eaaef = _bd.ApplyDifferences(_eaaef, _cgfe)
	_dbbc.SetEncoder(_eaaef)
	return &PdfFont{_cadf: &_dbbc}, _eaaef, nil
}

// SetPatternByName sets a pattern resource specified by keyName.
func (_agbf *PdfPageResources) SetPatternByName(keyName _dg.PdfObjectName, pattern _dg.PdfObject) error {
	if _agbf.Pattern == nil {
		_agbf.Pattern = _dg.MakeDict()
	}
	_gcgde, _febaf := _dg.GetDict(_agbf.Pattern)
	if !_febaf {
		return _dg.ErrTypeError
	}
	_gcgde.Set(keyName, pattern)
	return nil
}

// AddAnnotation appends `annot` to the list of page annotations.
func (_cddcb *PdfPage) AddAnnotation(annot *PdfAnnotation) {
	if _cddcb._cadgg == nil {
		_cddcb.GetAnnotations()
	}
	_cddcb._cadgg = append(_cddcb._cadgg, annot)
}

// NewPdfColorspaceDeviceGray returns a new grayscale colorspace.
func NewPdfColorspaceDeviceGray() *PdfColorspaceDeviceGray { return &PdfColorspaceDeviceGray{} }

// CharcodesToUnicode converts the character codes `charcodes` to a slice of runes.
// How it works:
//  1. Use the ToUnicode CMap if there is one.
//  2. Use the underlying font's encoding.
func (_gfbg *PdfFont) CharcodesToUnicode(charcodes []_bd.CharCode) []rune {
	_cdeg, _, _ := _gfbg.CharcodesToUnicodeWithStats(charcodes)
	return _cdeg
}

// GetNamedDestinations returns the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_gbed *PdfReader) GetNamedDestinations() (_dg.PdfObject, error) {
	_gceed := _dg.ResolveReference(_gbed._gccfb.Get("\u0044\u0065\u0073t\u0073"))
	if _gceed == nil {
		return nil, nil
	}
	if !_gbed._dadcef {
		_abefg := _gbed.traverseObjectData(_gceed)
		if _abefg != nil {
			return nil, _abefg
		}
	}
	return _gceed, nil
}

// ToPdfObject implements interface PdfModel.
func (_daa *PdfAnnotationFileAttachment) ToPdfObject() _dg.PdfObject {
	_daa.PdfAnnotation.ToPdfObject()
	_fgeg := _daa._cdf
	_ceg := _fgeg.PdfObject.(*_dg.PdfObjectDictionary)
	_daa.PdfAnnotationMarkup.appendToPdfDictionary(_ceg)
	_ceg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074"))
	_ceg.SetIfNotNil("\u0046\u0053", _daa.FS)
	_ceg.SetIfNotNil("\u004e\u0061\u006d\u0065", _daa.Name)
	return _fgeg
}

// PdfAction represents an action in PDF (section 12.6 p. 412).
type PdfAction struct {
	_bg  PdfModel
	Type _dg.PdfObject
	S    _dg.PdfObject
	Next _dg.PdfObject
	_cbd *_dg.PdfIndirectObject
}

// GetNumPages returns the number of pages in the document.
func (_dfgca *PdfReader) GetNumPages() (int, error) {
	if _dfgca._baad.GetCrypter() != nil && !_dfgca._baad.IsAuthenticated() {
		return 0, _b.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	return len(_dfgca._daddd), nil
}

// NewOutlineDest returns a new outline destination which can be used
// with outline items.
func NewOutlineDest(page int64, x, y float64) OutlineDest {
	return OutlineDest{Page: page, Mode: "\u0058\u0059\u005a", X: x, Y: y}
}
func _agec(_fdfebc _dg.PdfObject) (PdfFunction, error) {
	_fdfebc = _dg.ResolveReference(_fdfebc)
	if _abeac, _beabd := _fdfebc.(*_dg.PdfObjectStream); _beabd {
		_fgef := _abeac.PdfObjectDictionary
		_dcgge, _gfdd := _fgef.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_dg.PdfObjectInteger)
		if !_gfdd {
			_ag.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _bf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_dcgge == 0 {
			return _adcgg(_abeac)
		} else if *_dcgge == 4 {
			return _edda(_abeac)
		} else {
			return nil, _bf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _aaccg, _aaae := _fdfebc.(*_dg.PdfIndirectObject); _aaae {
		_eedf, _afdbd := _aaccg.PdfObject.(*_dg.PdfObjectDictionary)
		if !_afdbd {
			_ag.Log.Error("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006eg\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _bf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		_gead, _afdbd := _eedf.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_dg.PdfObjectInteger)
		if !_afdbd {
			_ag.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _bf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_gead == 2 {
			return _fdgf(_aaccg)
		} else if *_gead == 3 {
			return _ccgad(_aaccg)
		} else {
			return nil, _bf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _gaaad, _gcfaf := _fdfebc.(*_dg.PdfObjectDictionary); _gcfaf {
		_edcaf, _ccff := _gaaad.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_dg.PdfObjectInteger)
		if !_ccff {
			_ag.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _bf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_edcaf == 2 {
			return _fdgf(_gaaad)
		} else if *_edcaf == 3 {
			return _ccgad(_gaaad)
		} else {
			return nil, _bf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else {
		_ag.Log.Debug("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0054\u0079\u0070\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0023\u0076", _fdfebc)
		return nil, _bf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_beade *PdfColorspaceICCBased) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _beade.Alternate == nil {
		if _beade.N == 1 {
			_dcee := NewPdfColorspaceDeviceGray()
			return _dcee.ColorFromFloats(vals)
		} else if _beade.N == 3 {
			_fdeg := NewPdfColorspaceDeviceRGB()
			return _fdeg.ColorFromFloats(vals)
		} else if _beade.N == 4 {
			_baeee := NewPdfColorspaceDeviceCMYK()
			return _baeee.ColorFromFloats(vals)
		} else {
			return nil, _bf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _beade.Alternate.ColorFromFloats(vals)
}
func (_gecfacb *PdfWriter) setHashIDs(_eccbcf _ed.Hash) error {
	_acaaa := _eccbcf.Sum(nil)
	if _gecfacb._cedaf == "" {
		_gecfacb._cedaf = _be.EncodeToString(_acaaa[:8])
	}
	_gecfacb.setDocumentIDs(_gecfacb._cedaf, _be.EncodeToString(_acaaa[8:]))
	return nil
}
func _adcgg(_ddcaf *_dg.PdfObjectStream) (*PdfFunctionType0, error) {
	_edagc := &PdfFunctionType0{}
	_edagc._bedaf = _ddcaf
	_dege := _ddcaf.PdfObjectDictionary
	_bbfgf, _gdbdf := _dg.TraceToDirectObject(_dege.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_dg.PdfObjectArray)
	if !_gdbdf {
		_ag.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _bbfgf.Len() < 0 || _bbfgf.Len()%2 != 0 {
		_ag.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _bf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_edagc.NumInputs = _bbfgf.Len() / 2
	_fgcff, _fdgec := _bbfgf.ToFloat64Array()
	if _fdgec != nil {
		return nil, _fdgec
	}
	_edagc.Domain = _fgcff
	_bbfgf, _gdbdf = _dg.TraceToDirectObject(_dege.Get("\u0052\u0061\u006eg\u0065")).(*_dg.PdfObjectArray)
	if !_gdbdf {
		_ag.Log.Error("\u0052\u0061\u006e\u0067e \u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _bbfgf.Len() < 0 || _bbfgf.Len()%2 != 0 {
		return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_edagc.NumOutputs = _bbfgf.Len() / 2
	_bdgac, _fdgec := _bbfgf.ToFloat64Array()
	if _fdgec != nil {
		return nil, _fdgec
	}
	_edagc.Range = _bdgac
	_bbfgf, _gdbdf = _dg.TraceToDirectObject(_dege.Get("\u0053\u0069\u007a\u0065")).(*_dg.PdfObjectArray)
	if !_gdbdf {
		_ag.Log.Error("\u0053i\u007ae\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_fbdgf, _fdgec := _bbfgf.ToIntegerArray()
	if _fdgec != nil {
		return nil, _fdgec
	}
	if len(_fbdgf) != _edagc.NumInputs {
		_ag.Log.Error("T\u0061\u0062\u006c\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063h\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072 o\u0066\u0020\u0069n\u0070u\u0074\u0073")
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_edagc.Size = _fbdgf
	_fgegg, _gdbdf := _dg.TraceToDirectObject(_dege.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065")).(*_dg.PdfObjectInteger)
	if !_gdbdf {
		_ag.Log.Error("B\u0069\u0074\u0073\u0050\u0065\u0072S\u0061\u006d\u0070\u006c\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if *_fgegg != 1 && *_fgegg != 2 && *_fgegg != 4 && *_fgegg != 8 && *_fgegg != 12 && *_fgegg != 16 && *_fgegg != 24 && *_fgegg != 32 {
		_ag.Log.Error("\u0042\u0069\u0074s \u0070\u0065\u0072\u0020\u0073\u0061\u006d\u0070\u006ce\u0020o\u0075t\u0073i\u0064\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064\u0029", *_fgegg)
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_edagc.BitsPerSample = int(*_fgegg)
	_edagc.Order = 1
	_afbdgg, _gdbdf := _dg.TraceToDirectObject(_dege.Get("\u004f\u0072\u0064e\u0072")).(*_dg.PdfObjectInteger)
	if _gdbdf {
		if *_afbdgg != 1 && *_afbdgg != 3 {
			_ag.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u006fr\u0064\u0065r\u0020\u0028\u0025\u0064\u0029", *_afbdgg)
			return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		_edagc.Order = int(*_afbdgg)
	}
	_bbfgf, _gdbdf = _dg.TraceToDirectObject(_dege.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_dg.PdfObjectArray)
	if _gdbdf {
		_ebffc, _geaf := _bbfgf.ToFloat64Array()
		if _geaf != nil {
			return nil, _geaf
		}
		_edagc.Encode = _ebffc
	}
	_bbfgf, _gdbdf = _dg.TraceToDirectObject(_dege.Get("\u0044\u0065\u0063\u006f\u0064\u0065")).(*_dg.PdfObjectArray)
	if _gdbdf {
		_edacd, _abec := _bbfgf.ToFloat64Array()
		if _abec != nil {
			return nil, _abec
		}
		_edagc.Decode = _edacd
	}
	_cfeae, _fdgec := _dg.DecodeStream(_ddcaf)
	if _fdgec != nil {
		return nil, _fdgec
	}
	_edagc._bgdg = _cfeae
	return _edagc, nil
}
func _geee(_acabc _dg.PdfObject) (*PdfColorspaceCalGray, error) {
	_ddgc := NewPdfColorspaceCalGray()
	if _gagc, _baga := _acabc.(*_dg.PdfIndirectObject); _baga {
		_ddgc._gcbf = _gagc
	}
	_acabc = _dg.TraceToDirectObject(_acabc)
	_cbdg, _eceg := _acabc.(*_dg.PdfObjectArray)
	if !_eceg {
		return nil, _b.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _cbdg.Len() != 2 {
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0061\u006cG\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_acabc = _dg.TraceToDirectObject(_cbdg.Get(0))
	_abgg, _eceg := _acabc.(*_dg.PdfObjectName)
	if !_eceg {
		return nil, _b.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u0020\u006e\u0061m\u0065\u0020\u006e\u006f\u0074\u0020\u0061 \u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_abgg != "\u0043a\u006c\u0047\u0072\u0061\u0079" {
		return nil, _b.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0043\u0061\u006c\u0047\u0072\u0061\u0079 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_acabc = _dg.TraceToDirectObject(_cbdg.Get(1))
	_fbfdc, _eceg := _acabc.(*_dg.PdfObjectDictionary)
	if !_eceg {
		return nil, _b.Errorf("\u0043\u0061lG\u0072\u0061\u0079 \u0064\u0069\u0063\u0074 no\u0074 a\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u006f\u0062\u006a\u0065c\u0074")
	}
	_acabc = _fbfdc.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_acabc = _dg.TraceToDirectObject(_acabc)
	_dccaa, _eceg := _acabc.(*_dg.PdfObjectArray)
	if !_eceg {
		return nil, _b.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020W\u0068\u0069\u0074e\u0050o\u0069\u006e\u0074")
	}
	if _dccaa.Len() != 3 {
		return nil, _b.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
	}
	_ccba, _bbf := _dccaa.GetAsFloat64Slice()
	if _bbf != nil {
		return nil, _bbf
	}
	_ddgc.WhitePoint = _ccba
	_acabc = _fbfdc.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _acabc != nil {
		_acabc = _dg.TraceToDirectObject(_acabc)
		_ddde, _abce := _acabc.(*_dg.PdfObjectArray)
		if !_abce {
			return nil, _b.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020B\u006c\u0061\u0063k\u0050o\u0069\u006e\u0074")
		}
		if _ddde.Len() != 3 {
			return nil, _b.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061c\u006b\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
		}
		_eddb, _bedb := _ddde.GetAsFloat64Slice()
		if _bedb != nil {
			return nil, _bedb
		}
		_ddgc.BlackPoint = _eddb
	}
	_acabc = _fbfdc.Get("\u0047\u0061\u006dm\u0061")
	if _acabc != nil {
		_acabc = _dg.TraceToDirectObject(_acabc)
		_gcfb, _edgb := _dg.GetNumberAsFloat(_acabc)
		if _edgb != nil {
			return nil, _b.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0067\u0061\u006d\u006d\u0061\u0020\u006e\u006ft\u0020\u0061\u0020n\u0075m\u0062\u0065\u0072")
		}
		_ddgc.Gamma = _gcfb
	}
	return _ddgc, nil
}

// ToWriter creates a new writer from the current reader, based on the specified options.
// If no options are provided, all reader properties are copied to the writer.
func (_eedaeb *PdfReader) ToWriter(opts *ReaderToWriterOpts) (*PdfWriter, error) {
	_efab := NewPdfWriter()
	if opts == nil {
		opts = &ReaderToWriterOpts{}
	}
	_ebefa, _adef := _eedaeb.GetNumPages()
	if _adef != nil {
		_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _adef)
		return nil, _adef
	}
	for _cddgc := 1; _cddgc <= _ebefa; _cddgc++ {
		_abdf, _fgcce := _eedaeb.GetPage(_cddgc)
		if _fgcce != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fgcce)
			return nil, _fgcce
		}
		if opts.PageProcessCallback != nil {
			_fgcce = opts.PageProcessCallback(_cddgc, _abdf)
			if _fgcce != nil {
				_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fgcce)
				return nil, _fgcce
			}
		} else if opts.PageCallback != nil {
			opts.PageCallback(_cddgc, _abdf)
		}
		_fgcce = _efab.AddPage(_abdf)
		if _fgcce != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fgcce)
			return nil, _fgcce
		}
	}
	_efab._efacd = _eedaeb.PdfVersion()
	if !opts.SkipInfo {
		_ecccf, _egag := _eedaeb.GetPdfInfo()
		if _egag != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _egag)
		} else {
			_efab._efbfa.PdfObject = _ecccf.ToPdfObject()
		}
	}
	if !opts.SkipMetadata {
		if _dcegf := _eedaeb._gccfb.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _dcegf != nil {
			if _caca := _efab.SetCatalogMetadata(_dcegf); _caca != nil {
				return nil, _caca
			}
		}
	}
	if !opts.SkipAcroForm {
		_dgcba := _efab.SetForms(_eedaeb.AcroForm)
		if _dgcba != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgcba)
			return nil, _dgcba
		}
	}
	if !opts.SkipOutlines {
		_efab.AddOutlineTree(_eedaeb.GetOutlineTree())
	}
	if !opts.SkipOCProperties {
		_daddaf, _gddef := _eedaeb.GetOCProperties()
		if _gddef != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gddef)
		} else {
			_gddef = _efab.SetOCProperties(_daddaf)
			if _gddef != nil {
				_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gddef)
			}
		}
	}
	if !opts.SkipPageLabels {
		_cddf, _dgcf := _eedaeb.GetPageLabels()
		if _dgcf != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgcf)
		} else {
			_dgcf = _efab.SetPageLabels(_cddf)
			if _dgcf != nil {
				_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgcf)
			}
		}
	}
	if !opts.SkipNamedDests {
		_aadca, _ebgga := _eedaeb.GetNamedDestinations()
		if _ebgga != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ebgga)
		} else {
			_ebgga = _efab.SetNamedDestinations(_aadca)
			if _ebgga != nil {
				_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ebgga)
			}
		}
	}
	if !opts.SkipNameDictionary {
		_afcgcg, _gfbe := _eedaeb.GetNameDictionary()
		if _gfbe != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gfbe)
		} else {
			_gfbe = _efab.SetNameDictionary(_afcgcg)
			if _gfbe != nil {
				_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gfbe)
			}
		}
	}
	if !opts.SkipRotation && _eedaeb.Rotate != nil {
		if _gccfc := _efab.SetRotation(*_eedaeb.Rotate); _gccfc != nil {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gccfc)
		}
	}
	return &_efab, nil
}

// SetPageLabels sets the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_cgcag *PdfWriter) SetPageLabels(pageLabels _dg.PdfObject) error {
	if pageLabels == nil {
		return nil
	}
	_ag.Log.Trace("\u0053\u0065t\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006cs.\u002e\u002e")
	_cgcag._ecdf.Set("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073", pageLabels)
	return _cgcag.addObjects(pageLabels)
}

// PdfShadingType6 is a Coons patch mesh.
type PdfShadingType6 struct {
	*PdfShading
	BitsPerCoordinate *_dg.PdfObjectInteger
	BitsPerComponent  *_dg.PdfObjectInteger
	BitsPerFlag       *_dg.PdfObjectInteger
	Decode            *_dg.PdfObjectArray
	Function          []PdfFunction
}

// SetFillImage attach a model.Image to push button.
func (_gegd *PdfFieldButton) SetFillImage(image *Image) {
	if _gegd.IsPush() {
		_gegd._gddc = image
	}
}

var _fecec = map[string]struct{}{"\u0046\u0054": {}, "\u004b\u0069\u0064\u0073": {}, "\u0054": {}, "\u0054\u0055": {}, "\u0054\u004d": {}, "\u0046\u0066": {}, "\u0056": {}, "\u0044\u0056": {}, "\u0041\u0041": {}, "\u0044\u0041": {}, "\u0051": {}, "\u0044\u0053": {}, "\u0052\u0056": {}}

// GetContainingPdfObject returns the container of the outline item (indirect object).
func (_ggeafa *PdfOutlineItem) GetContainingPdfObject() _dg.PdfObject { return _ggeafa._fbgf }
func _agea(_dcdeb _dg.PdfObject) (*PdfPattern, error) {
	_bgadec := &PdfPattern{}
	var _gacddg *_dg.PdfObjectDictionary
	if _edbg, _egde := _dg.GetIndirect(_dcdeb); _egde {
		_bgadec._eacce = _edbg
		_abceg, _degea := _edbg.PdfObject.(*_dg.PdfObjectDictionary)
		if !_degea {
			_ag.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0028g\u006f\u0074\u0020%\u0054\u0029", _edbg.PdfObject)
			return nil, _dg.ErrTypeError
		}
		_gacddg = _abceg
	} else if _aafde, _fgab := _dg.GetStream(_dcdeb); _fgab {
		_bgadec._eacce = _aafde
		_gacddg = _aafde.PdfObjectDictionary
	} else {
		_ag.Log.Debug("\u0050a\u0074\u0074e\u0072\u006e\u0020\u006eo\u0074\u0020\u0061n\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 o\u0062\u006a\u0065c\u0074\u0020o\u0072\u0020\u0073\u0074\u0072\u0065a\u006d\u002e \u0025\u0054", _dcdeb)
		return nil, _dg.ErrTypeError
	}
	_cfbae := _gacddg.Get("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065")
	if _cfbae == nil {
		_ag.Log.Debug("\u0050\u0064\u0066\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069n\u0067\u0020\u0050\u0061\u0074t\u0065\u0072n\u0054\u0079\u0070\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_cdad, _edbbg := _cfbae.(*_dg.PdfObjectInteger)
	if !_edbbg {
		_ag.Log.Debug("\u0050\u0061tt\u0065\u0072\u006e \u0074\u0079\u0070\u0065 no\u0074 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067er\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _cfbae)
		return nil, _dg.ErrTypeError
	}
	if *_cdad != 1 && *_cdad != 2 {
		_ag.Log.Debug("\u0050\u0061\u0074\u0074e\u0072\u006e\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d \u0031/\u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", *_cdad)
		return nil, _dg.ErrRangeError
	}
	_bgadec.PatternType = int64(*_cdad)
	switch *_cdad {
	case 1:
		_dfead, _ecedd := _afbc(_gacddg)
		if _ecedd != nil {
			return nil, _ecedd
		}
		_dfead.PdfPattern = _bgadec
		_bgadec._cgdcc = _dfead
		return _bgadec, nil
	case 2:
		_bffcd, _beadeb := _adabe(_gacddg)
		if _beadeb != nil {
			return nil, _beadeb
		}
		_bffcd.PdfPattern = _bgadec
		_bgadec._cgdcc = _bffcd
		return _bgadec, nil
	}
	return nil, _bf.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
}
func (_caa *PdfReader) newPdfActionResetFormFromDict(_bgb *_dg.PdfObjectDictionary) (*PdfActionResetForm, error) {
	return &PdfActionResetForm{Fields: _bgb.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _bgb.Get("\u0046\u006c\u0061g\u0073")}, nil
}

// NewPdfAnnotation3D returns a new 3d annotation.
func NewPdfAnnotation3D() *PdfAnnotation3D {
	_cggg := NewPdfAnnotation()
	_fff := &PdfAnnotation3D{}
	_fff.PdfAnnotation = _cggg
	_cggg.SetContext(_fff)
	return _fff
}

// DecodeArray returns the range of color component values in DeviceCMYK colorspace.
func (_dedec *PdfColorspaceDeviceCMYK) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// ToPdfObject implements interface PdfModel.
func (_cagc *PdfAnnotationWidget) ToPdfObject() _dg.PdfObject {
	_cagc.PdfAnnotation.ToPdfObject()
	_dbbb := _cagc._cdf
	_cdgdg := _dbbb.PdfObject.(*_dg.PdfObjectDictionary)
	if _cagc._dgac {
		return _dbbb
	}
	_cagc._dgac = true
	_cdgdg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0057\u0069\u0064\u0067\u0065\u0074"))
	_cdgdg.SetIfNotNil("\u0048", _cagc.H)
	_cdgdg.SetIfNotNil("\u004d\u004b", _cagc.MK)
	_cdgdg.SetIfNotNil("\u0041", _cagc.A)
	_cdgdg.SetIfNotNil("\u0041\u0041", _cagc.AA)
	_cdgdg.SetIfNotNil("\u0042\u0053", _cagc.BS)
	_abaa := _cagc.Parent
	if _cagc._cgg != nil {
		if _cagc._cgg._egce == _cagc._cdf {
			_cagc._cgg.ToPdfObject()
		}
		_abaa = _cagc._cgg.GetContainingPdfObject()
	}
	if _abaa != _dbbb {
		_cdgdg.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _abaa)
	}
	_cagc._dgac = false
	return _dbbb
}

// AcroFormNeedsRepair returns true if the document contains widget annotations
// linked to fields which are not referenced in the AcroForm. The AcroForm can
// be repaired using the RepairAcroForm method of the reader.
func (_dddge *PdfReader) AcroFormNeedsRepair() (bool, error) {
	var _gefga []*PdfField
	if _dddge.AcroForm != nil {
		_gefga = _dddge.AcroForm.AllFields()
	}
	_beeab := make(map[*PdfField]struct{}, len(_gefga))
	for _, _adcff := range _gefga {
		_beeab[_adcff] = struct{}{}
	}
	for _, _cgbfd := range _dddge.PageList {
		_cbgcce, _dfebb := _cgbfd.GetAnnotations()
		if _dfebb != nil {
			return false, _dfebb
		}
		for _, _ffffg := range _cbgcce {
			_efec, _cdgbe := _ffffg.GetContext().(*PdfAnnotationWidget)
			if !_cdgbe {
				continue
			}
			_befaa := _efec.Field()
			if _befaa == nil {
				return true, nil
			}
			if _, _gbba := _beeab[_befaa]; !_gbba {
				return true, nil
			}
		}
	}
	return false, nil
}

// FontDescriptor returns font's PdfFontDescriptor. This may be a builtin descriptor for standard 14
// fonts but must be an explicit descriptor for other fonts.
func (_gbgfb *PdfFont) FontDescriptor() *PdfFontDescriptor {
	if _gbgfb.baseFields()._ccfb != nil {
		return _gbgfb.baseFields()._ccfb
	}
	if _bffeb := _gbgfb._cadf.getFontDescriptor(); _bffeb != nil {
		return _bffeb
	}
	_ag.Log.Error("\u0041\u006cl \u0066\u006f\u006et\u0073\u0020\u0068\u0061ve \u0061 D\u0065\u0073\u0063\u0072\u0069\u0070\u0074or\u002e\u0020\u0066\u006f\u006e\u0074\u003d%\u0073", _gbgfb)
	return nil
}
func _cbaa(_ebdb _dg.PdfObject) (*_dg.PdfObjectDictionary, *fontCommon, error) {
	_cebg := &fontCommon{}
	if _efbg, _eabag := _ebdb.(*_dg.PdfIndirectObject); _eabag {
		_cebg._bgggd = _efbg.ObjectNumber
	}
	_gafg, _bdee := _dg.GetDict(_ebdb)
	if !_bdee {
		_ag.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _ebdb)
		return nil, nil, ErrFontNotSupported
	}
	_aafa, _bdee := _dg.GetNameVal(_gafg.Get("\u0054\u0079\u0070\u0065"))
	if !_bdee {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	if _aafa != "\u0046\u006f\u006e\u0074" {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u0020\u0049\u006e\u0063\u006f\u006d\u0070\u0061t\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u003d\u0025\u0071\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0025\u0071.", _aafa, "\u0046\u006f\u006e\u0074")
		return nil, nil, _dg.ErrTypeError
	}
	_bbfgd, _bdee := _dg.GetNameVal(_gafg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_bdee {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020F\u006f\u006e\u0074 \u0049\u006e\u0063o\u006d\u0070a\u0074\u0069\u0062\u0069\u006c\u0069t\u0079. \u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	_cebg._bcga = _bbfgd
	_dgbae, _bdee := _dg.GetNameVal(_gafg.Get("\u004e\u0061\u006d\u0065"))
	if _bdee {
		_cebg._cefg = _dgbae
	}
	_aabbd := _gafg.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _aabbd != nil {
		_cebg._ebbff = _dg.TraceToDirectObject(_aabbd)
		_dfdeg, _eeagb := _bfbgg(_cebg._ebbff, _cebg)
		if _eeagb != nil {
			return _gafg, _cebg, _eeagb
		}
		_cebg._ecfb = _dfdeg
	} else if _bbfgd == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" || _bbfgd == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_beaf, _fgdcg := _ff.NewCIDSystemInfo(_gafg.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
		if _fgdcg != nil {
			return _gafg, _cebg, _fgdcg
		}
		_ddda := _b.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0055\u0043\u0053\u0032", _beaf.Registry, _beaf.Ordering)
		if _ff.IsPredefinedCMap(_ddda) {
			_cebg._ecfb, _fgdcg = _ff.LoadPredefinedCMap(_ddda)
			if _fgdcg != nil {
				_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _ddda, _fgdcg)
			}
		}
	}
	_fced := _gafg.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
	if _fced != nil {
		_ffdd, _fgga := _deg(_fced)
		if _fgga != nil {
			_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0042\u0061\u0064\u0020\u0066\u006f\u006et\u0020d\u0065s\u0063r\u0069\u0070\u0074\u006f\u0072\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fgga)
			return _gafg, _cebg, _fgga
		}
		_cebg._ccfb = _ffdd
	}
	if _bbfgd != "\u0054\u0079\u0070e\u0033" {
		_bccd, _eacdd := _dg.GetNameVal(_gafg.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
		if !_eacdd {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u006f\u006et\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069t\u0079\u002e\u0020\u0042\u0061se\u0046\u006f\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
			return _gafg, _cebg, ErrRequiredAttributeMissing
		}
		_cebg._ecbf = _bccd
	}
	return _gafg, _cebg, nil
}

// Flags returns the field flags for the field accounting for any inherited flags.
func (_eeb *PdfField) Flags() FieldFlag {
	var _bcbb FieldFlag
	_fdabc, _fffdg := _eeb.inherit(func(_cgcbd *PdfField) bool {
		if _cgcbd.Ff != nil {
			_bcbb = FieldFlag(*_cgcbd.Ff)
			return true
		}
		return false
	})
	if _fffdg != nil {
		_ag.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0069\u006e\u0067\u0020\u0066\u006c\u0061\u0067\u0073\u0020\u0076\u0069\u0061\u0020\u0069\u006e\u0068\u0065\u0072\u0069t\u0061\u006e\u0063\u0065\u003a\u0020\u0025\u0076", _fffdg)
	}
	if !_fdabc {
		_ag.Log.Trace("N\u006f\u0020\u0066\u0069\u0065\u006cd\u0020\u0066\u006c\u0061\u0067\u0073 \u0066\u006f\u0075\u006e\u0064\u0020\u002d \u0061\u0073\u0073\u0075\u006d\u0065\u0020\u0063\u006c\u0065a\u0072")
	}
	return _bcbb
}

// SetContext sets the sub action (context).
func (_gfe *PdfAction) SetContext(ctx PdfModel) { _gfe._bg = ctx }

// SetDSS sets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_fddd *PdfAppender) SetDSS(dss *DSS) {
	if dss != nil {
		_fddd.updateObjectsDeep(dss.ToPdfObject(), nil)
	}
	_fddd._cfdd = dss
}

// NewPdfAnnotationRedact returns a new redact annotation.
func NewPdfAnnotationRedact() *PdfAnnotationRedact {
	_eadg := NewPdfAnnotation()
	_gdg := &PdfAnnotationRedact{}
	_gdg.PdfAnnotation = _eadg
	_gdg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_eadg.SetContext(_gdg)
	return _gdg
}

// B returns the value of the B component of the color.
func (_gbac *PdfColorLab) B() float64 { return _gbac[2] }
func _cdfe(_cdec *_dg.PdfObjectDictionary) (*PdfFieldButton, error) {
	_dadc := &PdfFieldButton{}
	_dadc.PdfField = NewPdfField()
	_dadc.PdfField.SetContext(_dadc)
	_dadc.Opt, _ = _dg.GetArray(_cdec.Get("\u004f\u0070\u0074"))
	_afcgc := NewPdfAnnotationWidget()
	_afcgc.A, _ = _dg.GetDict(_cdec.Get("\u0041"))
	_afcgc.AP, _ = _dg.GetDict(_cdec.Get("\u0041\u0050"))
	_afcgc.SetContext(_dadc)
	_dadc.PdfField.Annotations = append(_dadc.PdfField.Annotations, _afcgc)
	return _dadc, nil
}

// Encoder returns the font's text encoder.
func (_bfge *PdfFont) Encoder() _bd.TextEncoder {
	_ffccg := _bfge.actualFont()
	if _ffccg == nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0045n\u0063\u006f\u0064er\u0020\u006e\u006f\u0074\u0020\u0069m\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066o\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d%\u0023\u0054", _bfge._cadf)
		return nil
	}
	return _ffccg.Encoder()
}

// PdfPageResources is a Page resources model.
// Implements PdfModel.
type PdfPageResources struct {
	ExtGState  _dg.PdfObject
	ColorSpace _dg.PdfObject
	Pattern    _dg.PdfObject
	Shading    _dg.PdfObject
	XObject    _dg.PdfObject
	Font       _dg.PdfObject
	ProcSet    _dg.PdfObject
	Properties _dg.PdfObject
	_ddcfb     *_dg.PdfObjectDictionary
	_dadae     *PdfPageResourcesColorspaces
}

// GetContainingPdfObject returns the container of the resources object (indirect object).
func (_ffeef *PdfPageResources) GetContainingPdfObject() _dg.PdfObject { return _ffeef._ddcfb }

// VariableText contains the common attributes of a variable text.
// The VariableText is typically not used directly, but is can encapsulate by PdfField
// See section 12.7.3.3 "Variable Text" and Table 222 (pp. 434-436 PDF32000_2008).
type VariableText struct {
	DA *_dg.PdfObjectString
	Q  *_dg.PdfObjectInteger
	DS *_dg.PdfObjectString
	RV _dg.PdfObject
}

// SignatureValidationResult defines the response from the signature validation handler.
type SignatureValidationResult struct {

	// List of errors when validating the signature.
	Errors      []string
	IsSigned    bool
	IsVerified  bool
	IsTrusted   bool
	Fields      []*PdfField
	Name        string
	Date        PdfDate
	Reason      string
	Location    string
	ContactInfo string
	DiffResults *_ecb.DiffResults
	IsCrlFound  bool
	IsOcspFound bool

	// GeneralizedTime is the time at which the time-stamp token has been created by the TSA (RFC 3161).
	GeneralizedTime _a.Time
}

// Encoder returns the font's text encoder.
func (_dfda pdfCIDFontType2) Encoder() _bd.TextEncoder { return _dfda._dcde }

type pdfFontSimple struct {
	fontCommon
	_gccg  *_dg.PdfIndirectObject
	_cfagf map[_bd.CharCode]float64
	_bdeed _bd.TextEncoder
	_dbdb  _bd.TextEncoder
	_gagbe *PdfFontDescriptor

	// Encoding is subject to limitations that are described in 9.6.6, "Character Encoding".
	// BaseFont is derived differently.
	FirstChar _dg.PdfObject
	LastChar  _dg.PdfObject
	Widths    _dg.PdfObject
	Encoding  _dg.PdfObject
	_bfdee    *_bbg.RuneCharSafeMap
}

// Size returns the width and the height of the page. The method reports
// the page dimensions as displayed by a PDF viewer (i.e. page rotation is
// taken into account).
func (_bfbe *PdfPage) Size() (float64, float64, error) {
	_ecbe, _bfaea := _bfbe.GetMediaBox()
	if _bfaea != nil {
		return 0, 0, _bfaea
	}
	_gagba, _deff := _ecbe.Width(), _ecbe.Height()
	_dbgdf, _bfaea := _bfbe.GetRotate()
	if _bfaea != nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _bfaea.Error())
	}
	if _acaae := _dbgdf; _acaae%360 != 0 && _acaae%90 == 0 {
		if _dddd := (360 + _acaae%360) % 360; _dddd == 90 || _dddd == 270 {
			_gagba, _deff = _deff, _gagba
		}
	}
	return _gagba, _deff, nil
}

var (
	CourierName              = _bbg.CourierName
	CourierBoldName          = _bbg.CourierBoldName
	CourierObliqueName       = _bbg.CourierObliqueName
	CourierBoldObliqueName   = _bbg.CourierBoldObliqueName
	HelveticaName            = _bbg.HelveticaName
	HelveticaBoldName        = _bbg.HelveticaBoldName
	HelveticaObliqueName     = _bbg.HelveticaObliqueName
	HelveticaBoldObliqueName = _bbg.HelveticaBoldObliqueName
	SymbolName               = _bbg.SymbolName
	ZapfDingbatsName         = _bbg.ZapfDingbatsName
	TimesRomanName           = _bbg.TimesRomanName
	TimesBoldName            = _bbg.TimesBoldName
	TimesItalicName          = _bbg.TimesItalicName
	TimesBoldItalicName      = _bbg.TimesBoldItalicName
)

// PdfColorCalGray represents a CalGray colorspace.
type PdfColorCalGray float64

func (_efdb *PdfReader) newPdfAnnotationFreeTextFromDict(_bgde *_dg.PdfObjectDictionary) (*PdfAnnotationFreeText, error) {
	_decg := PdfAnnotationFreeText{}
	_cbgc, _dffd := _efdb.newPdfAnnotationMarkupFromDict(_bgde)
	if _dffd != nil {
		return nil, _dffd
	}
	_decg.PdfAnnotationMarkup = _cbgc
	_decg.DA = _bgde.Get("\u0044\u0041")
	_decg.Q = _bgde.Get("\u0051")
	_decg.RC = _bgde.Get("\u0052\u0043")
	_decg.DS = _bgde.Get("\u0044\u0053")
	_decg.CL = _bgde.Get("\u0043\u004c")
	_decg.IT = _bgde.Get("\u0049\u0054")
	_decg.BE = _bgde.Get("\u0042\u0045")
	_decg.RD = _bgde.Get("\u0052\u0044")
	_decg.BS = _bgde.Get("\u0042\u0053")
	_decg.LE = _bgde.Get("\u004c\u0045")
	return &_decg, nil
}
func (_ffegb *PdfWriter) writeObjects() {
	_ag.Log.Trace("\u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0025d\u0020\u006f\u0062\u006a", len(_ffegb._agaba))
	_ffegb._fffge = make(map[int]crossReference)
	_ffegb._fffge[0] = crossReference{Type: 0, ObjectNumber: 0, Generation: 0xFFFF}
	if _ffegb._cfcdcb.ObjectMap != nil {
		for _fccg, _eafdd := range _ffegb._cfcdcb.ObjectMap {
			if _fccg == 0 {
				continue
			}
			if _eafdd.XType == _dg.XrefTypeObjectStream {
				_fagea := crossReference{Type: 2, ObjectNumber: _eafdd.OsObjNumber, Index: _eafdd.OsObjIndex}
				_ffegb._fffge[_fccg] = _fagea
			}
			if _eafdd.XType == _dg.XrefTypeTableEntry {
				_dgbce := crossReference{Type: 1, ObjectNumber: _eafdd.ObjectNumber, Offset: _eafdd.Offset}
				_ffegb._fffge[_fccg] = _dgbce
			}
		}
	}
}

// PdfFunctionType2 defines an exponential interpolation of one input value and n
// output values:
//
//	f(x) = y_0, ..., y_(n-1)
//
// y_j = C0_j + x^N * (C1_j - C0_j); for 0 <= j < n
// When N=1 ; linear interpolation between C0 and C1.
type PdfFunctionType2 struct {
	Domain []float64
	Range  []float64
	C0     []float64
	C1     []float64
	N      float64
	_acfde *_dg.PdfIndirectObject
}

func _abeda(_gaddg *_dg.PdfIndirectObject) (*PdfOutline, error) {
	_dcccac, _gacbg := _gaddg.PdfObject.(*_dg.PdfObjectDictionary)
	if !_gacbg {
		return nil, _b.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_dfbcg := NewPdfOutline()
	if _ccgac := _dcccac.Get("\u0054\u0079\u0070\u0065"); _ccgac != nil {
		_ecbga, _cgfgg := _ccgac.(*_dg.PdfObjectName)
		if _cgfgg {
			if *_ecbga != "\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073" {
				_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u004f\u0075\u0074l\u0069\u006e\u0065s\u0020(\u0025\u0073\u0029", *_ecbga)
			}
		}
	}
	if _gbdge := _dcccac.Get("\u0043\u006f\u0075n\u0074"); _gbdge != nil {
		_cdfdbf, _fbcgb := _dg.GetNumberAsInt64(_gbdge)
		if _fbcgb != nil {
			return nil, _fbcgb
		}
		_dfbcg.Count = &_cdfdbf
	}
	return _dfbcg, nil
}
func (_dcab *PdfColorspaceDeviceGray) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
}

// NewPdfAnnotationUnderline returns a new text underline annotation.
func NewPdfAnnotationUnderline() *PdfAnnotationUnderline {
	_gcgb := NewPdfAnnotation()
	_bgd := &PdfAnnotationUnderline{}
	_bgd.PdfAnnotation = _gcgb
	_bgd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gcgb.SetContext(_bgd)
	return _bgd
}

// GetPdfVersion gets the version of the PDF used within this document.
func (_fcgbc *PdfWriter) GetPdfVersion() string { return _fcgbc.getPdfVersion() }
func _cbbgd(_cfbe *fontCommon) *pdfFontType0    { return &pdfFontType0{fontCommon: *_cfbe} }

// NewPdfReaderFromFile creates a new PdfReader from the speficied PDF file.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderFromFile(pdfFile string, opts *ReaderOpts) (*PdfReader, *_eb.File, error) {
	const _dgccd = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0046\u0072\u006f\u006dF\u0069\u006c\u0065"
	_bagbb, _fcaga := _eb.Open(pdfFile)
	if _fcaga != nil {
		return nil, nil, _fcaga
	}
	_faedb, _fcaga := _dcdd(_bagbb, opts, true, _dgccd)
	if _fcaga != nil {
		_bagbb.Close()
		return nil, nil, _fcaga
	}
	return _faedb, _bagbb, nil
}

// GetStandardApplier gets currently used StandardApplier..
func (_abbcg *PdfWriter) GetStandardApplier() StandardApplier { return _abbcg._afafd }
func _gacbe(_affd _dg.PdfObject) (map[_bd.CharCode]float64, error) {
	if _affd == nil {
		return nil, nil
	}
	_gadfb, _ebab := _dg.GetArray(_affd)
	if !_ebab {
		return nil, nil
	}
	_aabc := map[_bd.CharCode]float64{}
	_affg := _gadfb.Len()
	for _dggb := 0; _dggb < _affg-1; _dggb++ {
		_dggad := _dg.TraceToDirectObject(_gadfb.Get(_dggb))
		_gccfe, _decf := _dg.GetIntVal(_dggad)
		if !_decf {
			return nil, _b.Errorf("\u0042a\u0064\u0020\u0066\u006fn\u0074\u0020\u0057\u0020\u006fb\u006a0\u003a \u0069\u003d\u0025\u0064\u0020\u0025\u0023v", _dggb, _dggad)
		}
		_dggb++
		if _dggb > _affg-1 {
			return nil, _b.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _gadfb)
		}
		_dffad := _dg.TraceToDirectObject(_gadfb.Get(_dggb))
		switch _dffad.(type) {
		case *_dg.PdfObjectArray:
			_acbda, _ := _dg.GetArray(_dffad)
			if _efca, _fgge := _acbda.ToFloat64Array(); _fgge == nil {
				for _agccg := 0; _agccg < len(_efca); _agccg++ {
					_aabc[_bd.CharCode(_gccfe+_agccg)] = _efca[_agccg]
				}
			} else {
				return nil, _b.Errorf("\u0042\u0061\u0064 \u0066\u006f\u006e\u0074 \u0057\u0020\u0061\u0072\u0072\u0061\u0079 \u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _dggb, _dffad)
			}
		case *_dg.PdfObjectInteger:
			_fgbg, _fagb := _dg.GetIntVal(_dffad)
			if !_fagb {
				return nil, _b.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _dggb, _dffad)
			}
			_dggb++
			if _dggb > _affg-1 {
				return nil, _b.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _gadfb)
			}
			_gagce := _gadfb.Get(_dggb)
			_bcdge, _bccfd := _dg.GetNumberAsFloat(_gagce)
			if _bccfd != nil {
				return nil, _b.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0032\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _dggb, _gagce)
			}
			for _cdee := _gccfe; _cdee <= _fgbg; _cdee++ {
				_aabc[_bd.CharCode(_cdee)] = _bcdge
			}
		default:
			return nil, _b.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057 \u006f\u0062\u006a\u0031\u0020\u0074\u0079p\u0065\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _dggb, _dffad)
		}
	}
	return _aabc, nil
}

// ToInteger convert to an integer format.
func (_fagf *PdfColorCalGray) ToInteger(bits int) uint32 {
	_geag := _cg.Pow(2, float64(bits)) - 1
	return uint32(_geag * _fagf.Val())
}
func (_eefgd *PdfSignature) extractChainFromPKCS7() ([]*_bb.Certificate, error) {
	_dgccc, _cfceg := _eg.Parse(_eefgd.Contents.Bytes())
	if _cfceg != nil {
		return nil, _cfceg
	}
	return _dgccc.Certificates, nil
}

// ToPdfObject converts the font to a PDF representation.
func (_gddaa *pdfFontType3) ToPdfObject() _dg.PdfObject {
	if _gddaa._aedbb == nil {
		_gddaa._aedbb = &_dg.PdfIndirectObject{}
	}
	_gcffe := _gddaa.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0033")
	_gddaa._aedbb.PdfObject = _gcffe
	if _gddaa.FirstChar != nil {
		_gcffe.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _gddaa.FirstChar)
	}
	if _gddaa.LastChar != nil {
		_gcffe.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _gddaa.LastChar)
	}
	if _gddaa.Widths != nil {
		_gcffe.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _gddaa.Widths)
	}
	if _gddaa.Encoding != nil {
		_gcffe.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _gddaa.Encoding)
	} else if _gddaa._geffb != nil {
		_beebff := _gddaa._geffb.ToPdfObject()
		if _beebff != nil {
			_gcffe.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _beebff)
		}
	}
	if _gddaa.FontBBox != nil {
		_gcffe.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _gddaa.FontBBox)
	}
	if _gddaa.FontMatrix != nil {
		_gcffe.Set("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0069\u0072\u0078", _gddaa.FontMatrix)
	}
	if _gddaa.CharProcs != nil {
		_gcffe.Set("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s", _gddaa.CharProcs)
	}
	if _gddaa.Resources != nil {
		_gcffe.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _gddaa.Resources)
	}
	return _gddaa._aedbb
}
func (_efgdd *PdfFilespec) getDict() *_dg.PdfObjectDictionary {
	if _gcega, _abgd := _efgdd._bcbe.(*_dg.PdfIndirectObject); _abgd {
		_ccbc, _fbabe := _gcega.PdfObject.(*_dg.PdfObjectDictionary)
		if !_fbabe {
			return nil
		}
		return _ccbc
	} else if _ddccd, _ccdb := _efgdd._bcbe.(*_dg.PdfObjectDictionary); _ccdb {
		return _ddccd
	} else {
		_ag.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074 \u0074\u0079p\u0065\u0020(\u0025T\u0029", _efgdd._bcbe)
		return nil
	}
}

// NewPdfColorspaceCalRGB returns a new CalRGB colorspace object.
func NewPdfColorspaceCalRGB() *PdfColorspaceCalRGB {
	_daga := &PdfColorspaceCalRGB{}
	_daga.BlackPoint = []float64{0.0, 0.0, 0.0}
	_daga.Gamma = []float64{1.0, 1.0, 1.0}
	_daga.Matrix = []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	return _daga
}

// PdfColorspaceLab is a L*, a*, b* 3 component colorspace.
type PdfColorspaceLab struct {
	WhitePoint []float64
	BlackPoint []float64
	Range      []float64
	_aadb      *_dg.PdfIndirectObject
}

func (_bgbb Image) getBase() _fc.ImageBase {
	return _fc.NewImageBase(int(_bgbb.Width), int(_bgbb.Height), int(_bgbb.BitsPerComponent), _bgbb.ColorComponents, _bgbb.Data, _bgbb._dgeb, _bgbb._gfbb)
}

// Compress is yet to be implemented.
// Should be able to compress in terms of JPEG quality parameter,
// and DPI threshold (need to know bounding area dimensions).
func (_ecafd DefaultImageHandler) Compress(input *Image, quality int64) (*Image, error) {
	return input, nil
}

// ToPdfObject convert PdfInfo to pdf object.
func (_geeb *PdfInfo) ToPdfObject() _dg.PdfObject {
	_fgdc := _dg.MakeDict()
	_fgdc.SetIfNotNil("\u0054\u0069\u0074l\u0065", _geeb.Title)
	_fgdc.SetIfNotNil("\u0041\u0075\u0074\u0068\u006f\u0072", _geeb.Author)
	_fgdc.SetIfNotNil("\u0053u\u0062\u006a\u0065\u0063\u0074", _geeb.Subject)
	_fgdc.SetIfNotNil("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _geeb.Keywords)
	_fgdc.SetIfNotNil("\u0043r\u0065\u0061\u0074\u006f\u0072", _geeb.Creator)
	_fgdc.SetIfNotNil("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _geeb.Producer)
	_fgdc.SetIfNotNil("\u0054r\u0061\u0070\u0070\u0065\u0064", _geeb.Trapped)
	if _geeb.CreationDate != nil {
		_fgdc.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _geeb.CreationDate.ToPdfObject())
	}
	if _geeb.ModifiedDate != nil {
		_fgdc.SetIfNotNil("\u004do\u0064\u0044\u0061\u0074\u0065", _geeb.ModifiedDate.ToPdfObject())
	}
	for _, _begf := range _geeb._dccdg.Keys() {
		_fgdc.SetIfNotNil(_begf, _geeb._dccdg.Get(_begf))
	}
	return _fgdc
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the L, A and B components of the color.
func (_bfdf *PdfColorspaceLab) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gfge, _faab := _dg.GetNumbersAsFloat(objects)
	if _faab != nil {
		return nil, _faab
	}
	return _bfdf.ColorFromFloats(_gfge)
}

// PdfAnnotationFileAttachment represents FileAttachment annotations.
// (Section 12.5.6.15).
type PdfAnnotationFileAttachment struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	FS   _dg.PdfObject
	Name _dg.PdfObject
}

// NewPdfReader returns a new PdfReader for an input io.ReadSeeker interface. Can be used to read PDF from
// memory or file. Immediately loads and traverses the PDF structure including pages and page contents (if
// not encrypted). Loads entire document structure into memory.
// Alternatively a lazy-loading reader can be created with NewPdfReaderLazy which loads only references,
// and references are loaded from disk into memory on an as-needed basis.
func NewPdfReader(rs _cf.ReadSeeker) (*PdfReader, error) {
	const _gcfeea = "\u006do\u0064e\u006c\u003a\u004e\u0065\u0077P\u0064\u0066R\u0065\u0061\u0064\u0065\u0072"
	return _dcdd(rs, &ReaderOpts{}, false, _gcfeea)
}

// ToPdfObject implements interface PdfModel.
func (_geg *PdfActionJavaScript) ToPdfObject() _dg.PdfObject {
	_geg.PdfAction.ToPdfObject()
	_bacb := _geg._cbd
	_eed := _bacb.PdfObject.(*_dg.PdfObjectDictionary)
	_eed.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeJavaScript)))
	_eed.SetIfNotNil("\u004a\u0053", _geg.JS)
	return _bacb
}
func (_fggc *PdfAnnotationMarkup) appendToPdfDictionary(_ebgge *_dg.PdfObjectDictionary) {
	_ebgge.SetIfNotNil("\u0054", _fggc.T)
	if _fggc.Popup != nil {
		_ebgge.Set("\u0050\u006f\u0070u\u0070", _fggc.Popup.ToPdfObject())
	}
	_ebgge.SetIfNotNil("\u0043\u0041", _fggc.CA)
	_ebgge.SetIfNotNil("\u0052\u0043", _fggc.RC)
	_ebgge.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _fggc.CreationDate)
	_ebgge.SetIfNotNil("\u0049\u0052\u0054", _fggc.IRT)
	_ebgge.SetIfNotNil("\u0053\u0075\u0062\u006a", _fggc.Subj)
	_ebgge.SetIfNotNil("\u0052\u0054", _fggc.RT)
	_ebgge.SetIfNotNil("\u0049\u0054", _fggc.IT)
	_ebgge.SetIfNotNil("\u0045\u0078\u0044\u0061\u0074\u0061", _fggc.ExData)
}

// ToPdfObject returns the choice field dictionary within an indirect object (container).
func (_dcegd *PdfFieldChoice) ToPdfObject() _dg.PdfObject {
	_dcegd.PdfField.ToPdfObject()
	_cdef := _dcegd._egce
	_bcfcc := _cdef.PdfObject.(*_dg.PdfObjectDictionary)
	_bcfcc.Set("\u0046\u0054", _dg.MakeName("\u0043\u0068"))
	if _dcegd.Opt != nil {
		_bcfcc.Set("\u004f\u0070\u0074", _dcegd.Opt)
	}
	if _dcegd.TI != nil {
		_bcfcc.Set("\u0054\u0049", _dcegd.TI)
	}
	if _dcegd.I != nil {
		_bcfcc.Set("\u0049", _dcegd.I)
	}
	return _cdef
}
func (_fcde *PdfReader) newPdfAnnotationStrikeOut(_gdeb *_dg.PdfObjectDictionary) (*PdfAnnotationStrikeOut, error) {
	_ddd := PdfAnnotationStrikeOut{}
	_fcae, _cgdg := _fcde.newPdfAnnotationMarkupFromDict(_gdeb)
	if _cgdg != nil {
		return nil, _cgdg
	}
	_ddd.PdfAnnotationMarkup = _fcae
	_ddd.QuadPoints = _gdeb.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_ddd, nil
}

// PdfShadingType5 is a Lattice-form Gouraud-shaded triangle mesh.
type PdfShadingType5 struct {
	*PdfShading
	BitsPerCoordinate *_dg.PdfObjectInteger
	BitsPerComponent  *_dg.PdfObjectInteger
	VerticesPerRow    *_dg.PdfObjectInteger
	Decode            *_dg.PdfObjectArray
	Function          []PdfFunction
}

func _edaac(_gegdb *_dg.PdfObjectDictionary) (*PdfShadingType5, error) {
	_cdcbg := PdfShadingType5{}
	_gddba := _gegdb.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _gddba == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_feffe, _eddd := _gddba.(*_dg.PdfObjectInteger)
	if !_eddd {
		_ag.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gddba)
		return nil, _dg.ErrTypeError
	}
	_cdcbg.BitsPerCoordinate = _feffe
	_gddba = _gegdb.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _gddba == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_feffe, _eddd = _gddba.(*_dg.PdfObjectInteger)
	if !_eddd {
		_ag.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _gddba)
		return nil, _dg.ErrTypeError
	}
	_cdcbg.BitsPerComponent = _feffe
	_gddba = _gegdb.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077")
	if _gddba == nil {
		_ag.Log.Debug("\u0052\u0065\u0071u\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0056\u0065\u0072\u0074\u0069c\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077")
		return nil, ErrRequiredAttributeMissing
	}
	_feffe, _eddd = _gddba.(*_dg.PdfObjectInteger)
	if !_eddd {
		_ag.Log.Debug("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006et\u0065\u0067\u0065\u0072\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _gddba)
		return nil, _dg.ErrTypeError
	}
	_cdcbg.VerticesPerRow = _feffe
	_gddba = _gegdb.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _gddba == nil {
		_ag.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fbce, _eddd := _gddba.(*_dg.PdfObjectArray)
	if !_eddd {
		_ag.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _gddba)
		return nil, _dg.ErrTypeError
	}
	_cdcbg.Decode = _fbce
	if _bdge := _gegdb.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _bdge != nil {
		_cdcbg.Function = []PdfFunction{}
		if _gdbbf, _beede := _bdge.(*_dg.PdfObjectArray); _beede {
			for _, _ebfab := range _gdbbf.Elements() {
				_ecabe, _acgeg := _agec(_ebfab)
				if _acgeg != nil {
					_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _acgeg)
					return nil, _acgeg
				}
				_cdcbg.Function = append(_cdcbg.Function, _ecabe)
			}
		} else {
			_eaaee, _deaec := _agec(_bdge)
			if _deaec != nil {
				_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _deaec)
				return nil, _deaec
			}
			_cdcbg.Function = append(_cdcbg.Function, _eaaee)
		}
	}
	return &_cdcbg, nil
}
func _dbdg(_acef string) map[string]string {
	_agded := _decga.Split(_acef, -1)
	_abgc := map[string]string{}
	for _, _fbfca := range _agded {
		_feagd := _caegd.FindStringSubmatch(_fbfca)
		if _feagd == nil {
			continue
		}
		_gbga, _efda := _feagd[1], _feagd[2]
		_abgc[_gbga] = _efda
	}
	return _abgc
}
func (_eggdb *PdfWriter) writeBytes(_aadef []byte) {
	if _eggdb._ffefc != nil {
		return
	}
	_dfgafg, _gggdb := _eggdb._bddfa.Write(_aadef)
	_eggdb._fbbfc += int64(_dfgafg)
	_eggdb._ffefc = _gggdb
}

// NewPdfActionGoToR returns a new "go to remote" action.
func NewPdfActionGoToR() *PdfActionGoToR {
	_ee := NewPdfAction()
	_ea := &PdfActionGoToR{}
	_ea.PdfAction = _ee
	_ee.SetContext(_ea)
	return _ea
}

const (
	ButtonTypeCheckbox ButtonType = iota
	ButtonTypePush     ButtonType = iota
	ButtonTypeRadio    ButtonType = iota
)

func _gcga(_egadf string) (map[_bd.CharCode]_bd.GlyphName, error) {
	_daag := _ga.Split(_egadf, "\u000a")
	_cdgfg := make(map[_bd.CharCode]_bd.GlyphName)
	for _, _bfaba := range _daag {
		_ffbcg := _bacefb.FindStringSubmatch(_bfaba)
		if _ffbcg == nil {
			continue
		}
		_cfagd, _ccdba := _ffbcg[1], _ffbcg[2]
		_bgdbe, _dbddf := _fbb.Atoi(_cfagd)
		if _dbddf != nil {
			_ag.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0042\u0061\u0064\u0020\u0065\u006e\u0063\u006fd\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u002e \u0025\u0071", _bfaba)
			return nil, _dg.ErrTypeError
		}
		_cdgfg[_bd.CharCode(_bgdbe)] = _bd.GlyphName(_ccdba)
	}
	_ag.Log.Trace("g\u0065\u0074\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0073\u003a\u0020\u006b\u0065\u0079V\u0061\u006c\u0075e\u0073=\u0025\u0023\u0076", _cdgfg)
	return _cdgfg, nil
}

// CharcodesToUnicodeWithStats is identical to CharcodesToUnicode except it returns more statistical
// information about hits and misses from the reverse mapping process.
// NOTE: The number of runes returned may be greater than the number of charcodes.
// TODO(peterwilliams97): Deprecate in v4 and use only CharcodesToStrings()
func (_bbeg *PdfFont) CharcodesToUnicodeWithStats(charcodes []_bd.CharCode) (_ebgb []rune, _ddga, _gdeaf int) {
	_aaf, _ddga, _gdeaf := _bbeg.CharcodesToStrings(charcodes)
	return []rune(_ga.Join(_aaf, "")), _ddga, _gdeaf
}

// PdfActionJavaScript represents a javaScript action.
type PdfActionJavaScript struct {
	*PdfAction
	JS _dg.PdfObject
}

// NewPdfOutlineTree returns an initialized PdfOutline tree.
func NewPdfOutlineTree() *PdfOutline {
	_fdeeg := NewPdfOutline()
	_fdeeg._baddf = &_fdeeg
	return _fdeeg
}

// BytesToCharcodes converts the bytes in a PDF string to character codes.
func (_bgge *PdfFont) BytesToCharcodes(data []byte) []_bd.CharCode {
	_ag.Log.Trace("\u0042\u0079\u0074es\u0054\u006f\u0043\u0068\u0061\u0072\u0063\u006f\u0064e\u0073:\u0020d\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", data, data)
	if _agfd, _ccdfg := _bgge._cadf.(*pdfFontType0); _ccdfg && _agfd._bfae != nil {
		if _bcef, _cfad := _agfd.bytesToCharcodes(data); _cfad {
			return _bcef
		}
	}
	var (
		_fffc  = make([]_bd.CharCode, 0, len(data)+len(data)%2)
		_ffdcd = _bgge.baseFields()
	)
	if _ffdcd._ecfb != nil {
		if _ebggb, _eggd := _ffdcd._ecfb.BytesToCharcodes(data); _eggd {
			for _, _cdcg := range _ebggb {
				_fffc = append(_fffc, _bd.CharCode(_cdcg))
			}
			return _fffc
		}
	}
	if _ffdcd.isCIDFont() {
		if len(data) == 1 {
			data = []byte{0, data[0]}
		}
		if len(data)%2 != 0 {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0064\u0061\u0074\u0061\u003d\u0025\u002b\u0076\u0020t\u006f\u0020\u0065\u0076\u0065n\u0020\u006ce\u006e\u0067\u0074\u0068", data)
			data = append(data, 0)
		}
		for _fffgf := 0; _fffgf < len(data); _fffgf += 2 {
			_cefe := uint16(data[_fffgf])<<8 | uint16(data[_fffgf+1])
			_fffc = append(_fffc, _bd.CharCode(_cefe))
		}
	} else {
		for _, _cafge := range data {
			_fffc = append(_fffc, _bd.CharCode(_cafge))
		}
	}
	return _fffc
}

// GenerateHashMaps generates DSS hashmaps for Certificates, OCSPs and CRLs to make sure they are unique.
func (_bdcdg *DSS) GenerateHashMaps() error {
	_cfdb, _febed := _bdcdg.generateHashMap(_bdcdg.Certs)
	if _febed != nil {
		return _febed
	}
	_bggd, _febed := _bdcdg.generateHashMap(_bdcdg.OCSPs)
	if _febed != nil {
		return _febed
	}
	_cgde, _febed := _bdcdg.generateHashMap(_bdcdg.CRLs)
	if _febed != nil {
		return _febed
	}
	_bdcdg._agbg = _cfdb
	_bdcdg._bggg = _bggd
	_bdcdg._dcdf = _cgde
	return nil
}

// GetPerms returns the Permissions dictionary
func (_gfec *PdfReader) GetPerms() *Permissions { return _gfec._fbcaa }

// GetExtGState gets the ExtGState specified by keyName. Returns a bool
// indicating whether it was found or not.
func (_gdcff *PdfPageResources) GetExtGState(keyName _dg.PdfObjectName) (_dg.PdfObject, bool) {
	if _gdcff.ExtGState == nil {
		return nil, false
	}
	_eddfdg, _bccbf := _dg.TraceToDirectObject(_gdcff.ExtGState).(*_dg.PdfObjectDictionary)
	if !_bccbf {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gdcff.ExtGState)
		return nil, false
	}
	if _aebee := _eddfdg.Get(keyName); _aebee != nil {
		return _aebee, true
	}
	return nil, false
}
func (_cdde *PdfColorspaceCalGray) String() string { return "\u0043a\u006c\u0047\u0072\u0061\u0079" }

// FillWithAppearance populates `form` with values provided by `provider`.
// If not nil, `appGen` is used to generate appearance dictionaries for the
// field annotations, based on the specified settings. Otherwise, appearance
// generation is skipped.
// e.g.: appGen := annotator.FieldAppearance{OnlyIfMissing: true, RegenerateTextFields: true}
// NOTE: In next major version this functionality will be part of Fill. (v4)
func (_dcfg *PdfAcroForm) FillWithAppearance(provider FieldValueProvider, appGen FieldAppearanceGenerator) error {
	_degde := _dcfg.fill(provider, appGen)
	if _degde != nil {
		return _degde
	}
	if _, _cddea := provider.(FieldImageProvider); _cddea {
		_degde = _dcfg.fillImageWithAppearance(provider.(FieldImageProvider), appGen)
	}
	return _degde
}

// PdfShadingPatternType3 is shading patterns that will use a Type 3 shading pattern (Radial).
type PdfShadingPatternType3 struct {
	*PdfPattern
	Shading   *PdfShadingType3
	Matrix    *_dg.PdfObjectArray
	ExtGState _dg.PdfObject
}

// PdfFieldSignature signature field represents digital signatures and optional data for authenticating
// the name of the signer and verifying document contents.
type PdfFieldSignature struct {
	*PdfField
	*PdfAnnotationWidget
	V    *PdfSignature
	Lock *_dg.PdfIndirectObject
	SV   *_dg.PdfIndirectObject
}

// PdfColorspaceSpecialIndexed is an indexed color space is a lookup table, where the input element
// is an index to the lookup table and the output is a color defined in the lookup table in the Base
// colorspace.
// [/Indexed base hival lookup]
type PdfColorspaceSpecialIndexed struct {
	Base   PdfColorspace
	HiVal  int
	Lookup _dg.PdfObject
	_bcea  []byte
	_fffd  *_dg.PdfIndirectObject
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for an RGB device.
func (_gadbb *PdfColorspaceDeviceRGB) GetNumComponents() int { return 3 }

// ToPdfObject returns the PDF representation of the colorspace.
func (_addf *PdfColorspaceSpecialPattern) ToPdfObject() _dg.PdfObject {
	if _addf.UnderlyingCS == nil {
		return _dg.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e")
	}
	_dgg := _dg.MakeArray(_dg.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_dgg.Append(_addf.UnderlyingCS.ToPdfObject())
	if _addf._cbff != nil {
		_addf._cbff.PdfObject = _dgg
		return _addf._cbff
	}
	return _dgg
}
func (_aggc *PdfReader) newPdfAnnotationInkFromDict(_ffed *_dg.PdfObjectDictionary) (*PdfAnnotationInk, error) {
	_ccde := PdfAnnotationInk{}
	_egg, _gbcf := _aggc.newPdfAnnotationMarkupFromDict(_ffed)
	if _gbcf != nil {
		return nil, _gbcf
	}
	_ccde.PdfAnnotationMarkup = _egg
	_ccde.InkList = _ffed.Get("\u0049n\u006b\u004c\u0069\u0073\u0074")
	_ccde.BS = _ffed.Get("\u0042\u0053")
	return &_ccde, nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// A, B and C components of the color. The values of the elements should be
// between 0 and 1.
func (_eccb *PdfColorspaceCalRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gfdb := vals[0]
	if _gfdb < 0.0 || _gfdb > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _gfdb)
		return nil, ErrColorOutOfRange
	}
	_ada := vals[1]
	if _ada < 0.0 || _ada > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _ada)
		return nil, ErrColorOutOfRange
	}
	_egea := vals[2]
	if _egea < 0.0 || _egea > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _egea)
		return nil, ErrColorOutOfRange
	}
	_gegea := NewPdfColorCalRGB(_gfdb, _ada, _egea)
	return _gegea, nil
}
func (_gdada *PdfAcroForm) fillImageWithAppearance(_gece FieldImageProvider, _beea FieldAppearanceGenerator) error {
	if _gdada == nil {
		return nil
	}
	_agga, _gcfef := _gece.FieldImageValues()
	if _gcfef != nil {
		return _gcfef
	}
	for _, _ggbab := range _gdada.AllFields() {
		_ddee := _ggbab.PartialName()
		_cggd, _agbe := _agga[_ddee]
		if !_agbe {
			if _dgcaf, _eegga := _ggbab.FullName(); _eegga == nil {
				_cggd, _agbe = _agga[_dgcaf]
			}
		}
		if !_agbe {
			_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _ddee)
			continue
		}
		switch _eaedg := _ggbab.GetContext().(type) {
		case *PdfFieldButton:
			if _eaedg.IsPush() {
				_eaedg.SetFillImage(_cggd)
			}
		}
		if _beea == nil {
			continue
		}
		for _, _bbce := range _ggbab.Annotations {
			_abbce, _ddgff := _beea.GenerateAppearanceDict(_gdada, _ggbab, _bbce)
			if _ddgff != nil {
				return _ddgff
			}
			_bbce.AP = _abbce
			_bbce.ToPdfObject()
		}
	}
	return nil
}

// GetPdfInfo returns the PDF info dictionary.
func (_fgfgb *PdfReader) GetPdfInfo() (*PdfInfo, error) {
	_dcec, _gbdfa := _fgfgb.GetTrailer()
	if _gbdfa != nil {
		return nil, _gbdfa
	}
	var _gbace *_dg.PdfObjectDictionary
	_cafa := _dcec.Get("\u0049\u006e\u0066\u006f")
	switch _edcc := _cafa.(type) {
	case *_dg.PdfObjectReference:
		_bbade := _edcc
		_cafa, _gbdfa = _fgfgb.GetIndirectObjectByNumber(int(_bbade.ObjectNumber))
		_cafa = _dg.TraceToDirectObject(_cafa)
		if _gbdfa != nil {
			return nil, _gbdfa
		}
		_gbace, _ = _cafa.(*_dg.PdfObjectDictionary)
	case *_dg.PdfObjectDictionary:
		_gbace = _edcc
	}
	if _gbace == nil {
		return nil, _bf.New("I\u006e\u0066\u006f\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0070r\u0065s\u0065\u006e\u0074")
	}
	_agfb, _gbdfa := NewPdfInfoFromObject(_gbace)
	if _gbdfa != nil {
		return nil, _gbdfa
	}
	return _agfb, nil
}

// NewPdfAnnotationText returns a new text annotation.
func NewPdfAnnotationText() *PdfAnnotationText {
	_afg := NewPdfAnnotation()
	_ggg := &PdfAnnotationText{}
	_ggg.PdfAnnotation = _afg
	_ggg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_afg.SetContext(_ggg)
	return _ggg
}
func (_eggcfb *PdfWriter) getPdfVersion() string {
	return _b.Sprintf("\u0025\u0064\u002e%\u0064", _eggcfb._efacd.Major, _eggcfb._efacd.Minor)
}

// Enable LTV enables the specified signature. The signing certificate
// chain is extracted from the signature dictionary. Optionally, additional
// certificates can be specified through the `extraCerts` parameter.
// The LTV client attempts to build the certificate chain up to a trusted root
// by downloading any missing certificates.
func (_ffgga *LTV) Enable(sig *PdfSignature, extraCerts []*_bb.Certificate) error {
	if _geeea := _ffgga.validateSig(sig); _geeea != nil {
		return _geeea
	}
	_ebedd, _deae := _ffgga.generateVRIKey(sig)
	if _deae != nil {
		return _deae
	}
	if _, _bgbcb := _ffgga._abca.VRI[_ebedd]; _bgbcb && _ffgga.SkipExisting {
		return nil
	}
	_bagf, _deae := sig.GetCerts()
	if _deae != nil {
		return _deae
	}
	return _ffgga.enable(_bagf, extraCerts, _ebedd)
}

// GetXHeight returns the XHeight of the font `descriptor`.
func (_becga *PdfFontDescriptor) GetXHeight() (float64, error) {
	return _dg.GetNumberAsFloat(_becga.XHeight)
}

// PdfFieldText represents a text field where user can enter text.
type PdfFieldText struct {
	*PdfField
	DA     *_dg.PdfObjectString
	Q      *_dg.PdfObjectInteger
	DS     *_dg.PdfObjectString
	RV     _dg.PdfObject
	MaxLen *_dg.PdfObjectInteger
}

// SetPdfCreator sets the Creator attribute of the output PDF.
func SetPdfCreator(creator string) { _fgefgf.Lock(); defer _fgefgf.Unlock(); _cebb = creator }
func (_acf *PdfReader) newPdfAnnotationTextFromDict(_bag *_dg.PdfObjectDictionary) (*PdfAnnotationText, error) {
	_bdgg := PdfAnnotationText{}
	_bdc, _gagda := _acf.newPdfAnnotationMarkupFromDict(_bag)
	if _gagda != nil {
		return nil, _gagda
	}
	_bdgg.PdfAnnotationMarkup = _bdc
	_bdgg.Open = _bag.Get("\u004f\u0070\u0065\u006e")
	_bdgg.Name = _bag.Get("\u004e\u0061\u006d\u0065")
	_bdgg.State = _bag.Get("\u0053\u0074\u0061t\u0065")
	_bdgg.StateModel = _bag.Get("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c")
	return &_bdgg, nil
}

// ToPdfObject returns the button field dictionary within an indirect object.
func (_decgb *PdfFieldButton) ToPdfObject() _dg.PdfObject {
	_decgb.PdfField.ToPdfObject()
	_gdad := _decgb._egce
	_daeg := _gdad.PdfObject.(*_dg.PdfObjectDictionary)
	_daeg.Set("\u0046\u0054", _dg.MakeName("\u0042\u0074\u006e"))
	if _decgb.Opt != nil {
		_daeg.Set("\u004f\u0070\u0074", _decgb.Opt)
	}
	return _gdad
}

// PdfAnnotationProjection represents Projection annotations.
type PdfAnnotationProjection struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
}

// FieldFilterFunc represents a PDF field filtering function. If the function
// returns true, the PDF field is kept, otherwise it is discarded.
type FieldFilterFunc func(*PdfField) bool

// SetEncoder sets the encoding for the underlying font.
// TODO(peterwilliams97): Change function signature to SetEncoder(encoder *textencoding.simpleEncoder).
// TODO(gunnsth): Makes sense if SetEncoder is removed from the interface fonts.Font as proposed in PR #260.
func (_gdaa *pdfFontSimple) SetEncoder(encoder _bd.TextEncoder) { _gdaa._bdeed = encoder }
func (_bea *PdfReader) newPdfActionRenditionFromDict(_gad *_dg.PdfObjectDictionary) (*PdfActionRendition, error) {
	return &PdfActionRendition{R: _gad.Get("\u0052"), AN: _gad.Get("\u0041\u004e"), OP: _gad.Get("\u004f\u0050"), JS: _gad.Get("\u004a\u0053")}, nil
}

// Mask returns the uin32 bitmask for the specific flag.
func (_cefd FieldFlag) Mask() uint32 { return uint32(_cefd) }

// AppendContentBytes creates a PDF stream from `cs` and appends it to the
// array of streams specified by the pages's Contents entry.
// If `wrapContents` is true, the content stream of the page is wrapped using
// a `q/Q` operator pair, so that its state does not affect the appended
// content stream.
func (_gggee *PdfPage) AppendContentBytes(cs []byte, wrapContents bool) error {
	_gbea := _gggee.GetContentStreamObjs()
	wrapContents = wrapContents && len(_gbea) > 0
	_dbbda := _dg.NewFlateEncoder()
	_bagdg := _dg.MakeArray()
	if wrapContents {
		_eace, _efedfb := _dg.MakeStream([]byte("\u0071\u000a"), _dbbda)
		if _efedfb != nil {
			return _efedfb
		}
		_bagdg.Append(_eace)
	}
	_bagdg.Append(_gbea...)
	if wrapContents {
		_effc, _dcfbf := _dg.MakeStream([]byte("\u000a\u0051\u000a"), _dbbda)
		if _dcfbf != nil {
			return _dcfbf
		}
		_bagdg.Append(_effc)
	}
	_fgbea, _bebee := _dg.MakeStream(cs, _dbbda)
	if _bebee != nil {
		return _bebee
	}
	_bagdg.Append(_fgbea)
	_gggee.Contents = _bagdg
	return nil
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_bcff *PdfColorspaceLab) ToPdfObject() _dg.PdfObject {
	_ccg := _dg.MakeArray()
	_ccg.Append(_dg.MakeName("\u004c\u0061\u0062"))
	_gacb := _dg.MakeDict()
	if _bcff.WhitePoint != nil {
		_bdggd := _dg.MakeArray(_dg.MakeFloat(_bcff.WhitePoint[0]), _dg.MakeFloat(_bcff.WhitePoint[1]), _dg.MakeFloat(_bcff.WhitePoint[2]))
		_gacb.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _bdggd)
	} else {
		_ag.Log.Error("\u004c\u0061\u0062: \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057h\u0069t\u0065P\u006fi\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	if _bcff.BlackPoint != nil {
		_geeec := _dg.MakeArray(_dg.MakeFloat(_bcff.BlackPoint[0]), _dg.MakeFloat(_bcff.BlackPoint[1]), _dg.MakeFloat(_bcff.BlackPoint[2]))
		_gacb.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _geeec)
	}
	if _bcff.Range != nil {
		_afge := _dg.MakeArray(_dg.MakeFloat(_bcff.Range[0]), _dg.MakeFloat(_bcff.Range[1]), _dg.MakeFloat(_bcff.Range[2]), _dg.MakeFloat(_bcff.Range[3]))
		_gacb.Set("\u0052\u0061\u006eg\u0065", _afge)
	}
	_ccg.Append(_gacb)
	if _bcff._aadb != nil {
		_bcff._aadb.PdfObject = _ccg
		return _bcff._aadb
	}
	return _ccg
}

// ToPdfObject converts rectangle to a PDF object.
func (_cggga *PdfRectangle) ToPdfObject() _dg.PdfObject {
	return _dg.MakeArray(_dg.MakeFloat(_cggga.Llx), _dg.MakeFloat(_cggga.Lly), _dg.MakeFloat(_cggga.Urx), _dg.MakeFloat(_cggga.Ury))
}

// GetXObjectFormByName returns the XObjectForm with the specified name from the
// page resources, if it exists.
func (_bggcd *PdfPageResources) GetXObjectFormByName(keyName _dg.PdfObjectName) (*XObjectForm, error) {
	_fgcbda, _ggcf := _bggcd.GetXObjectByName(keyName)
	if _fgcbda == nil {
		return nil, nil
	}
	if _ggcf != XObjectTypeForm {
		return nil, _bf.New("\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006f\u0072\u006d")
	}
	_dggc, _gecfc := NewXObjectFormFromStream(_fgcbda)
	if _gecfc != nil {
		return nil, _gecfc
	}
	return _dggc, nil
}

// GetOutlinesFlattened returns a flattened list of tree nodes and titles.
// NOTE: for most use cases, it is recommended to use the high-level GetOutlines
// method instead, which also provides information regarding the destination
// of the outline items.
func (_dffda *PdfReader) GetOutlinesFlattened() ([]*PdfOutlineTreeNode, []string, error) {
	var _gebfa []*PdfOutlineTreeNode
	var _cfaf []string
	var _cdagc func(*PdfOutlineTreeNode, *[]*PdfOutlineTreeNode, *[]string, int)
	_cdagc = func(_cbgcc *PdfOutlineTreeNode, _fgfae *[]*PdfOutlineTreeNode, _abebg *[]string, _bdefg int) {
		if _cbgcc == nil {
			return
		}
		if _cbgcc._baddf == nil {
			_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020M\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006e\u006fd\u0065\u002e\u0063o\u006et\u0065\u0078\u0074")
			return
		}
		_dcegc, _fgfad := _cbgcc._baddf.(*PdfOutlineItem)
		if _fgfad {
			*_fgfae = append(*_fgfae, &_dcegc.PdfOutlineTreeNode)
			_gbafg := _ga.Repeat("\u0020", _bdefg*2) + _dcegc.Title.Decoded()
			*_abebg = append(*_abebg, _gbafg)
		}
		if _cbgcc.First != nil {
			_cdefg := _ga.Repeat("\u0020", _bdefg*2) + "\u002b"
			*_abebg = append(*_abebg, _cdefg)
			_cdagc(_cbgcc.First, _fgfae, _abebg, _bdefg+1)
		}
		if _fgfad && _dcegc.Next != nil {
			_cdagc(_dcegc.Next, _fgfae, _abebg, _bdefg)
		}
	}
	_cdagc(_dffda._fcfc, &_gebfa, &_cfaf, 0)
	return _gebfa, _cfaf, nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_gfeeb *PdfColorspaceDeviceCMYK) ToPdfObject() _dg.PdfObject {
	return _dg.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b")
}
func (_gedf *PdfReader) newPdfAnnotationStampFromDict(_cdgb *_dg.PdfObjectDictionary) (*PdfAnnotationStamp, error) {
	_afbd := PdfAnnotationStamp{}
	_beee, _bcg := _gedf.newPdfAnnotationMarkupFromDict(_cdgb)
	if _bcg != nil {
		return nil, _bcg
	}
	_afbd.PdfAnnotationMarkup = _beee
	_afbd.Name = _cdgb.Get("\u004e\u0061\u006d\u0065")
	return &_afbd, nil
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_cgbfc *PdfFunctionType2) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_ag.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cecae := []float64{0.0}
	if _cgbfc.C0 != nil {
		_cecae = _cgbfc.C0
	}
	_gagbb := []float64{1.0}
	if _cgbfc.C1 != nil {
		_gagbb = _cgbfc.C1
	}
	var _ccacf []float64
	for _fbfe := 0; _fbfe < len(_cecae); _fbfe++ {
		_cfee := _cecae[_fbfe] + _cg.Pow(x[0], _cgbfc.N)*(_gagbb[_fbfe]-_cecae[_fbfe])
		_ccacf = append(_ccacf, _cfee)
	}
	return _ccacf, nil
}

// ToPdfObject returns the PDF representation of the pattern.
func (_gfed *PdfPattern) ToPdfObject() _dg.PdfObject {
	_gcfbc := _gfed.getDict()
	_gcfbc.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_gcfbc.Set("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065", _dg.MakeInteger(_gfed.PatternType))
	return _gfed._eacce
}

// PdfFieldChoice represents a choice field which includes scrollable list boxes and combo boxes.
type PdfFieldChoice struct {
	*PdfField
	Opt *_dg.PdfObjectArray
	TI  *_dg.PdfObjectInteger
	I   *_dg.PdfObjectArray
}

// NewCompliancePdfReader creates a PdfReader or an input io.ReadSeeker that during reading will scan the files for the
// metadata details. It could be used for the PDF standard implementations like PDF/A or PDF/X.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
func NewCompliancePdfReader(rs _cf.ReadSeeker) (*CompliancePdfReader, error) {
	const _dabd = "\u006d\u006f\u0064\u0065l\u003a\u004e\u0065\u0077\u0043\u006f\u006d\u0070\u006c\u0069a\u006ec\u0065\u0050\u0064\u0066\u0052\u0065\u0061d\u0065\u0072"
	_fabcg, _adebg := _dcdd(rs, &ReaderOpts{ComplianceMode: true}, false, _dabd)
	if _adebg != nil {
		return nil, _adebg
	}
	return &CompliancePdfReader{PdfReader: _fabcg}, nil
}

// NewPdfActionURI returns a new "Uri" action.
func NewPdfActionURI() *PdfActionURI {
	_dga := NewPdfAction()
	_afd := &PdfActionURI{}
	_afd.PdfAction = _dga
	_dga.SetContext(_afd)
	return _afd
}

// Permissions specify a permissions dictionary (PDF 1.5).
// (Section 12.8.4, Table 258 - Entries in a permissions dictionary p. 477 in PDF32000_2008).
type Permissions struct {
	DocMDP *PdfSignature
	_fbdaa *_dg.PdfObjectDictionary
}

// ParserMetadata gets the parser  metadata.
func (_adba *CompliancePdfReader) ParserMetadata() _dg.ParserMetadata {
	if _adba._acgf == (_dg.ParserMetadata{}) {
		_adba._acgf, _ = _adba._baad.ParserMetadata()
	}
	return _adba._acgf
}
func _gfebd(_dcffb _dg.PdfObject) (string, error) {
	_dcffb = _dg.TraceToDirectObject(_dcffb)
	switch _gffbd := _dcffb.(type) {
	case *_dg.PdfObjectString:
		return _gffbd.Str(), nil
	case *_dg.PdfObjectStream:
		_cfdda, _ebdf := _dg.DecodeStream(_gffbd)
		if _ebdf != nil {
			return "", _ebdf
		}
		return string(_cfdda), nil
	}
	return "", _b.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _dcffb)
}

// PdfAnnotationWatermark represents Watermark annotations.
// (Section 12.5.6.22).
type PdfAnnotationWatermark struct {
	*PdfAnnotation
	FixedPrint _dg.PdfObject
}

// ToPdfObject returns the PDF representation of the DSS dictionary.
func (_gfbd *DSS) ToPdfObject() _dg.PdfObject {
	_bgeg := _gfbd._agdcg.PdfObject.(*_dg.PdfObjectDictionary)
	_bgeg.Clear()
	_gaeb := _dg.MakeDict()
	for _gcgd, _dcfe := range _gfbd.VRI {
		_gaeb.Set(*_dg.MakeName(_gcgd), _dcfe.ToPdfObject())
	}
	_bgeg.SetIfNotNil("\u0043\u0065\u0072t\u0073", _fecee(_gfbd.Certs))
	_bgeg.SetIfNotNil("\u004f\u0043\u0053P\u0073", _fecee(_gfbd.OCSPs))
	_bgeg.SetIfNotNil("\u0043\u0052\u004c\u0073", _fecee(_gfbd.CRLs))
	_bgeg.Set("\u0056\u0052\u0049", _gaeb)
	return _gfbd._agdcg
}
func _ggbff(_fcgd *fontCommon) *pdfCIDFontType0 { return &pdfCIDFontType0{fontCommon: *_fcgd} }

// NewPdfAnnotationLine returns a new line annotation.
func NewPdfAnnotationLine() *PdfAnnotationLine {
	_adbda := NewPdfAnnotation()
	_dba := &PdfAnnotationLine{}
	_dba.PdfAnnotation = _adbda
	_dba.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_adbda.SetContext(_dba)
	return _dba
}

const (
	_gafad = 0x00001
	_agcf  = 0x00002
	_bbfd  = 0x00004
	_caaca = 0x00008
	_gbaba = 0x00020
	_bfeb  = 0x00040
	_ggee  = 0x10000
	_gabe  = 0x20000
	_edgfe = 0x40000
)

// RunesToCharcodeBytes maps the provided runes to charcode bytes and it
// returns the resulting slice of bytes, along with the number of runes which
// could not be converted. If the number of misses is 0, all runes were
// successfully converted.
func (_efgde *PdfFont) RunesToCharcodeBytes(data []rune) ([]byte, int) {
	var _adcd []_bd.TextEncoder
	var _ecac _bd.CMapEncoder
	if _fdfaf := _efgde.baseFields()._ecfb; _fdfaf != nil {
		_ecac = _bd.NewCMapEncoder("", nil, _fdfaf)
	}
	_dafca := _efgde.Encoder()
	if _dafca != nil {
		switch _befa := _dafca.(type) {
		case _bd.SimpleEncoder:
			_abbcc := _befa.BaseName()
			if _, _aedc := _eabb[_abbcc]; _aedc {
				_adcd = append(_adcd, _dafca)
			}
		}
	}
	if len(_adcd) == 0 {
		if _efgde.baseFields()._ecfb != nil {
			_adcd = append(_adcd, _ecac)
		}
		if _dafca != nil {
			_adcd = append(_adcd, _dafca)
		}
	}
	var _bgac _bc.Buffer
	var _egaeb int
	for _, _cgfd := range data {
		var _ffcda bool
		for _, _cabbf := range _adcd {
			if _gbbga := _cabbf.Encode(string(_cgfd)); len(_gbbga) > 0 {
				_bgac.Write(_gbbga)
				_ffcda = true
				break
			}
		}
		if !_ffcda {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020f\u0061\u0069\u006ce\u0064\u0020\u0074\u006f \u006d\u0061\u0070\u0020\u0072\u0075\u006e\u0065\u0020\u0060\u0025\u002b\u0071\u0060\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065", _cgfd)
			_egaeb++
		}
	}
	if _egaeb != 0 {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0061\u006cl\u0020\u0072\u0075\u006e\u0065\u0073\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072c\u006fd\u0065\u0073\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0052\u0075\u006e\u0065\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073 \u0065\u006e\u0063\u006f\u0064e\u0072\u0073\u003d\u0025\u002b\u0076", len(data), _egaeb, _efgde, _adcd)
	}
	return _bgac.Bytes(), _egaeb
}

// NewPdfAnnotationWidget returns an initialized annotation widget.
func NewPdfAnnotationWidget() *PdfAnnotationWidget {
	_ceec := NewPdfAnnotation()
	_bbbb := &PdfAnnotationWidget{}
	_bbbb.PdfAnnotation = _ceec
	_ceec.SetContext(_bbbb)
	return _bbbb
}
func (_fdgg *pdfCIDFontType0) baseFields() *fontCommon { return &_fdgg.fontCommon }

// IsEncrypted returns true if the PDF file is encrypted.
func (_cdgdc *PdfReader) IsEncrypted() (bool, error) { return _cdgdc._baad.IsEncrypted() }

// SetOCProperties sets the optional content properties.
func (_eaaae *PdfWriter) SetOCProperties(ocProperties _dg.PdfObject) error {
	_edcee := _eaaae._ecdf
	if ocProperties != nil {
		_ag.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u004f\u0043\u0020\u0050\u0072o\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002e\u002e\u002e")
		_edcee.Set("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", ocProperties)
		return _eaaae.addObjects(ocProperties)
	}
	return nil
}

// AddOutlineTree adds outlines to a PDF file.
func (_ebgd *PdfWriter) AddOutlineTree(outlineTree *PdfOutlineTreeNode) { _ebgd._cedcg = outlineTree }

// String returns a string representation of PdfTransformParamsDocMDP.
func (_fcgdd *PdfTransformParamsDocMDP) String() string {
	return _b.Sprintf("\u0025\u0073\u0020\u0050\u003a\u0020\u0025\u0073\u0020V\u003a\u0020\u0025\u0073", _fcgdd.Type, _fcgdd.P, _fcgdd.V)
}

// ToGoImage converts the unidoc Image to a golang Image structure.
func (_bdefc *Image) ToGoImage() (_gd.Image, error) {
	_ag.Log.Trace("\u0043\u006f\u006e\u0076er\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u006f\u0020\u0069\u006d\u0061g\u0065")
	_feecb, _adfde := _fc.NewImage(int(_bdefc.Width), int(_bdefc.Height), int(_bdefc.BitsPerComponent), _bdefc.ColorComponents, _bdefc.Data, _bdefc._dgeb, _bdefc._gfbb)
	if _adfde != nil {
		return nil, _adfde
	}
	return _feecb, nil
}

// NewPdfAnnotation returns an initialized generic PDF annotation model.
func NewPdfAnnotation() *PdfAnnotation {
	_dbg := &PdfAnnotation{}
	_dbg._cdf = _dg.MakeIndirectObject(_dg.MakeDict())
	return _dbg
}

// ParsePdfObject parses input pdf object into given output intent.
func (_ggfac *PdfOutputIntent) ParsePdfObject(object _dg.PdfObject) error {
	_adddg, _fadfg := _dg.GetDict(object)
	if !_fadfg {
		_ag.Log.Error("\u0055\u006e\u006bno\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020%\u0054 \u0066o\u0072 \u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074", object)
		return _bf.New("\u0075\u006e\u006b\u006e\u006fw\u006e\u0020\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074")
	}
	_ggfac._cfaef = _adddg
	_ggfac.Type, _ = _adddg.GetString("\u0054\u0079\u0070\u0065")
	_eggab, _fadfg := _adddg.GetString("\u0053")
	if _fadfg {
		switch _eggab {
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411":
			_ggfac.S = PdfOutputIntentTypeA1
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00412":
			_ggfac.S = PdfOutputIntentTypeA2
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00413":
			_ggfac.S = PdfOutputIntentTypeA3
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00414":
			_ggfac.S = PdfOutputIntentTypeA4
		case "\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0058":
			_ggfac.S = PdfOutputIntentTypeX
		}
	}
	_ggfac.OutputCondition, _ = _adddg.GetString("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e")
	_ggfac.OutputConditionIdentifier, _ = _adddg.GetString("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r")
	_ggfac.RegistryName, _ = _adddg.GetString("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065")
	_ggfac.Info, _ = _adddg.GetString("\u0049\u006e\u0066\u006f")
	if _fdfbe, _daeb := _dg.GetStream(_adddg.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _daeb {
		_ggfac.ColorComponents, _ = _dg.GetIntVal(_fdfbe.Get("\u004e"))
		_feagg, _agaeb := _dg.DecodeStream(_fdfbe)
		if _agaeb != nil {
			return _agaeb
		}
		_ggfac.DestOutputProfile = _feagg
	}
	return nil
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_edaaf pdfCIDFontType0) GetCharMetrics(code _bd.CharCode) (_bbg.CharMetrics, bool) {
	_bcaed := _edaaf._gfdca
	if _aaddg, _dcbeb := _edaaf._gfcba[code]; _dcbeb {
		_bcaed = _aaddg
	}
	return _bbg.CharMetrics{Wx: _bcaed}, true
}

// ToPdfObject return the CalGray colorspace as a PDF object (name dictionary).
func (_egbb *PdfColorspaceCalGray) ToPdfObject() _dg.PdfObject {
	_fcf := &_dg.PdfObjectArray{}
	_fcf.Append(_dg.MakeName("\u0043a\u006c\u0047\u0072\u0061\u0079"))
	_abeg := _dg.MakeDict()
	if _egbb.WhitePoint != nil {
		_abeg.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _dg.MakeArray(_dg.MakeFloat(_egbb.WhitePoint[0]), _dg.MakeFloat(_egbb.WhitePoint[1]), _dg.MakeFloat(_egbb.WhitePoint[2])))
	} else {
		_ag.Log.Error("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006fi\u006e\u0074\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _egbb.BlackPoint != nil {
		_abeg.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _dg.MakeArray(_dg.MakeFloat(_egbb.BlackPoint[0]), _dg.MakeFloat(_egbb.BlackPoint[1]), _dg.MakeFloat(_egbb.BlackPoint[2])))
	}
	_abeg.Set("\u0047\u0061\u006dm\u0061", _dg.MakeFloat(_egbb.Gamma))
	_fcf.Append(_abeg)
	if _egbb._gcbf != nil {
		_egbb._gcbf.PdfObject = _fcf
		return _egbb._gcbf
	}
	return _fcf
}

type modelManager struct {
	_aaddgg map[PdfModel]_dg.PdfObject
	_eccbc  map[_dg.PdfObject]PdfModel
}

// GetCharMetrics returns the character metrics for the specified character code.  A bool flag is
// returned to indicate whether or not the entry was found in the glyph to charcode mapping.
// How it works:
//  1. Return a value the /Widths array (charWidths) if there is one.
//  2. If the font has the same name as a standard 14 font then return width=250.
//  3. Otherwise return no match and let the caller substitute a default.
func (_ddacc pdfFontSimple) GetCharMetrics(code _bd.CharCode) (_bbg.CharMetrics, bool) {
	if _abcbf, _cfecg := _ddacc._cfagf[code]; _cfecg {
		return _bbg.CharMetrics{Wx: _abcbf}, true
	}
	if _bbg.IsStdFont(_bbg.StdFontName(_ddacc._ecbf)) {
		return _bbg.CharMetrics{Wx: 250}, true
	}
	return _bbg.CharMetrics{}, false
}
func _ceca(_edb *_dg.PdfIndirectObject, _eggf *_dg.PdfObjectDictionary) (*DSS, error) {
	if _edb == nil {
		_edb = _dg.MakeIndirectObject(nil)
	}
	_edb.PdfObject = _dg.MakeDict()
	_eaga := map[string]*VRI{}
	if _aabf, _aagcg := _dg.GetDict(_eggf.Get("\u0056\u0052\u0049")); _aagcg {
		for _, _fffdf := range _aabf.Keys() {
			if _cfcd, _ddgf := _dg.GetDict(_aabf.Get(_fffdf)); _ddgf {
				_eaga[_ga.ToUpper(_fffdf.String())] = _eceba(_cfcd)
			}
		}
	}
	return &DSS{Certs: _eeafe(_eggf.Get("\u0043\u0065\u0072t\u0073")), OCSPs: _eeafe(_eggf.Get("\u004f\u0043\u0053P\u0073")), CRLs: _eeafe(_eggf.Get("\u0043\u0052\u004c\u0073")), VRI: _eaga, _agdcg: _edb}, nil
}

// NewPdfActionHide returns a new "hide" action.
func NewPdfActionHide() *PdfActionHide {
	_aee := NewPdfAction()
	_ecd := &PdfActionHide{}
	_ecd.PdfAction = _aee
	_aee.SetContext(_ecd)
	return _ecd
}

// ToPdfObject returns an indirect object containing the signature field dictionary.
func (_efed *PdfFieldSignature) ToPdfObject() _dg.PdfObject {
	if _efed.PdfAnnotationWidget != nil {
		_efed.PdfAnnotationWidget.ToPdfObject()
	}
	_efed.PdfField.ToPdfObject()
	_dagfc := _efed._egce
	_caaeb := _dagfc.PdfObject.(*_dg.PdfObjectDictionary)
	_caaeb.SetIfNotNil("\u0046\u0054", _dg.MakeName("\u0053\u0069\u0067"))
	_caaeb.SetIfNotNil("\u004c\u006f\u0063\u006b", _efed.Lock)
	_caaeb.SetIfNotNil("\u0053\u0056", _efed.SV)
	if _efed.V != nil {
		_caaeb.SetIfNotNil("\u0056", _efed.V.ToPdfObject())
	}
	return _dagfc
}
func _ababg(_ffccd _dg.PdfObject) (*PdfPageResourcesColorspaces, error) {
	_aege := &PdfPageResourcesColorspaces{}
	if _adag, _ecbed := _ffccd.(*_dg.PdfIndirectObject); _ecbed {
		_aege._dcebf = _adag
		_ffccd = _adag.PdfObject
	}
	_affe, _ddeda := _dg.GetDict(_ffccd)
	if !_ddeda {
		return nil, _bf.New("\u0043\u0053\u0020at\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_aege.Names = []string{}
	_aege.Colorspaces = map[string]PdfColorspace{}
	for _, _dcdc := range _affe.Keys() {
		_ddbf := _affe.Get(_dcdc)
		_aege.Names = append(_aege.Names, string(_dcdc))
		_aggbc, _fddb := NewPdfColorspaceFromPdfObject(_ddbf)
		if _fddb != nil {
			return nil, _fddb
		}
		_aege.Colorspaces[string(_dcdc)] = _aggbc
	}
	return _aege, nil
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_dabce pdfCIDFontType2) GetCharMetrics(code _bd.CharCode) (_bbg.CharMetrics, bool) {
	if _dfbe, _afed := _dabce._ddeb[code]; _afed {
		return _bbg.CharMetrics{Wx: _dfbe}, true
	}
	_fcag := rune(code)
	_bgfcb, _adfbb := _dabce._ebgc[_fcag]
	if !_adfbb {
		_bgfcb = int(_dabce._bfbc)
	}
	return _bbg.CharMetrics{Wx: float64(_bgfcb)}, true
}

// NewPermissions returns a new permissions object.
func NewPermissions(docMdp *PdfSignature) *Permissions {
	_fdef := Permissions{}
	_fdef.DocMDP = docMdp
	_ggafd := _dg.MakeDict()
	_ggafd.Set("\u0044\u006f\u0063\u004d\u0044\u0050", docMdp.ToPdfObject())
	_fdef._fbdaa = _ggafd
	return &_fdef
}

// ImageHandler interface implements common image loading and processing tasks.
// Implementing as an interface allows for the possibility to use non-standard libraries for faster
// loading and processing of images.
type ImageHandler interface {

	// Read any image type and load into a new Image object.
	Read(_aaeeg _cf.Reader) (*Image, error)

	// NewImageFromGoImage loads a NRGBA32 unidoc Image from a standard Go image structure.
	NewImageFromGoImage(_caaa _gd.Image) (*Image, error)

	// NewGrayImageFromGoImage loads a grayscale unidoc Image from a standard Go image structure.
	NewGrayImageFromGoImage(_efgge _gd.Image) (*Image, error)

	// Compress an image.
	Compress(_cfcde *Image, _bgbga int64) (*Image, error)
}

func _egcfg(_cgdda *XObjectImage) error {
	if _cgdda.SMask == nil {
		return nil
	}
	_dddag, _dbae := _cgdda.SMask.(*_dg.PdfObjectStream)
	if !_dbae {
		_ag.Log.Debug("\u0053\u004da\u0073\u006b\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0053\u0074re\u0061\u006d")
		return _dg.ErrTypeError
	}
	_ggeee := _dddag.PdfObjectDictionary
	_edgcd := _ggeee.Get("\u004d\u0061\u0074t\u0065")
	if _edgcd == nil {
		return nil
	}
	_cddfd, _gcfagd := _ceafc(_edgcd.(*_dg.PdfObjectArray))
	if _gcfagd != nil {
		return _gcfagd
	}
	_bbfca := _dg.MakeArrayFromFloats([]float64{_cddfd})
	_ggeee.SetIfNotNil("\u004d\u0061\u0074t\u0065", _bbfca)
	return nil
}

// Write writes out the PDF.
func (_ccge *PdfWriter) Write(writer _cf.Writer) error {
	_ag.Log.Trace("\u0057r\u0069\u0074\u0065\u0028\u0029")
	if _cdfbg := _ccge.writeOutlines(); _cdfbg != nil {
		return _cdfbg
	}
	if _cdfbg := _ccge.writeAcroFormFields(); _cdfbg != nil {
		return _cdfbg
	}
	_ccge.checkPendingObjects()
	if _cdfbg := _ccge.writeOutputIntents(); _cdfbg != nil {
		return _cdfbg
	}
	_ccge.setCatalogVersion()
	_ccge.copyObjects()
	if _cdfbg := _ccge.optimize(); _cdfbg != nil {
		return _cdfbg
	}
	if _cdfbg := _ccge.optimizeDocument(); _cdfbg != nil {
		return _cdfbg
	}
	var _ccgab _ed.Hash
	if _ccge._affea {
		_ccgab = _f.New()
		writer = _cf.MultiWriter(_ccgab, writer)
	}
	_ccge.setWriter(writer)
	_dedeed := _ccge.checkCrossReferenceStream()
	_ccffa, _dedeed := _ccge.mapObjectStreams(_dedeed)
	_ccge.adjustXRefAffectedVersion(_dedeed)
	_ccge.writeDocumentVersion()
	_ccge.updateObjectNumbers()
	_ccge.writeObjects()
	if _cdfbg := _ccge.writeObjectsInStreams(_ccffa); _cdfbg != nil {
		return _cdfbg
	}
	_dbabe := _ccge._fbbfc
	var _dfdf int
	for _bbbbb := range _ccge._fffge {
		if _bbbbb > _dfdf {
			_dfdf = _bbbbb
		}
	}
	if _ccge._affea {
		if _cdfbg := _ccge.setHashIDs(_ccgab); _cdfbg != nil {
			return _cdfbg
		}
	}
	if _dedeed {
		if _cdfbg := _ccge.writeXRefStreams(_dfdf, _dbabe); _cdfbg != nil {
			return _cdfbg
		}
	} else {
		_ccge.writeTrailer(_dfdf)
	}
	_ccge.makeOffSetReference(_dbabe)
	if _cdfbg := _ccge.flushWriter(); _cdfbg != nil {
		return _cdfbg
	}
	return nil
}

// PdfDate represents a date, which is a PDF string of the form:
// (D:YYYYMMDDHHmmSSOHH'mm)
type PdfDate struct {
	_bgfdb int64
	_gbbge int64
	_cbad  int64
	_cade  int64
	_ccgbg int64
	_aacfe int64
	_ggbdg byte
	_bdde  int64
	_gafe  int64
}

func (_fefde *LTV) buildCertChain(_afgdg, _fgdae []*_bb.Certificate) ([]*_bb.Certificate, map[string]*_bb.Certificate, error) {
	_dcgae := map[string]*_bb.Certificate{}
	for _, _ccdg := range _afgdg {
		_dcgae[_ccdg.Subject.CommonName] = _ccdg
	}
	_fegf := _afgdg
	for _, _cggb := range _fgdae {
		_cbedf := _cggb.Subject.CommonName
		if _, _gdead := _dcgae[_cbedf]; _gdead {
			continue
		}
		_dcgae[_cbedf] = _cggb
		_fegf = append(_fegf, _cggb)
	}
	if len(_fegf) == 0 {
		return nil, nil, ErrSignNoCertificates
	}
	var _cdbeg error
	for _faaf := _fegf[0]; _faaf != nil && !_fefde.CertClient.IsCA(_faaf); {
		_bfgda, _beadeg := _dcgae[_faaf.Issuer.CommonName]
		if !_beadeg {
			if _bfgda, _cdbeg = _fefde.CertClient.GetIssuer(_faaf); _cdbeg != nil {
				_ag.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _cdbeg)
				break
			}
			_dcgae[_faaf.Issuer.CommonName] = _bfgda
			_fegf = append(_fegf, _bfgda)
		}
		_faaf = _bfgda
	}
	return _fegf, _dcgae, nil
}

// GetContainingPdfObject implements interface PdfModel.
func (_abeb *PdfFilespec) GetContainingPdfObject() _dg.PdfObject { return _abeb._bcbe }
func _bfbgg(_gegcc _dg.PdfObject, _gbfgf *fontCommon) (*_ff.CMap, error) {
	_dcce, _aabd := _dg.GetStream(_gegcc)
	if !_aabd {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0074\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0054\u006f\u0043m\u0061\u0070\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0054\u0029", _gegcc)
		return nil, _dg.ErrTypeError
	}
	_debge, _efaab := _dg.DecodeStream(_dcce)
	if _efaab != nil {
		return nil, _efaab
	}
	_dceaff, _efaab := _ff.LoadCmapFromData(_debge, !_gbfgf.isCIDFont())
	if _efaab != nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u004e\u0075\u006d\u0062\u0065\u0072\u003d\u0025\u0064\u0020\u0065\u0072r=\u0025\u0076", _dcce.ObjectNumber, _efaab)
	}
	return _dceaff, _efaab
}

// PdfTransformParamsDocMDP represents a transform parameters dictionary for the DocMDP method and is used to detect
// modifications relative to a signature field that is signed by the author of a document.
// (Section 12.8.2.2, Table 254 - Entries in the DocMDP transform parameters dictionary p. 471 in PDF32000_2008).
type PdfTransformParamsDocMDP struct {
	Type *_dg.PdfObjectName
	P    *_dg.PdfObjectInteger
	V    *_dg.PdfObjectName
}

// PdfAnnotationLine represents Line annotations.
// (Section 12.5.6.7).
type PdfAnnotationLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	L       _dg.PdfObject
	BS      _dg.PdfObject
	LE      _dg.PdfObject
	IC      _dg.PdfObject
	LL      _dg.PdfObject
	LLE     _dg.PdfObject
	Cap     _dg.PdfObject
	IT      _dg.PdfObject
	LLO     _dg.PdfObject
	CP      _dg.PdfObject
	Measure _dg.PdfObject
	CO      _dg.PdfObject
}

// ImageToRGB converts an Image in a given PdfColorspace to an RGB image.
func (_deec *PdfColorspaceDeviceN) ImageToRGB(img Image) (Image, error) {
	_dfbc := _fcd.NewReader(img.getBase())
	_ecf := _fc.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, nil, img._dgeb, img._gfbb)
	_ceaf := _fcd.NewWriter(_ecf)
	_eeag := _cg.Pow(2, float64(img.BitsPerComponent)) - 1
	_fdacf := _deec.GetNumComponents()
	_fdbad := make([]uint32, _fdacf)
	_dcafd := make([]float64, _fdacf)
	for {
		_cfgca := _dfbc.ReadSamples(_fdbad)
		if _cfgca == _cf.EOF {
			break
		} else if _cfgca != nil {
			return img, _cfgca
		}
		for _acbe := 0; _acbe < _fdacf; _acbe++ {
			_fgaf := float64(_fdbad[_acbe]) / _eeag
			_dcafd[_acbe] = _fgaf
		}
		_cfbaa, _cfgca := _deec.TintTransform.Evaluate(_dcafd)
		if _cfgca != nil {
			return img, _cfgca
		}
		for _, _fffg := range _cfbaa {
			_fffg = _cg.Min(_cg.Max(0, _fffg), 1.0)
			if _cfgca = _ceaf.WriteSample(uint32(_fffg * _eeag)); _cfgca != nil {
				return img, _cfgca
			}
		}
	}
	return _deec.AlternateSpace.ImageToRGB(_edcf(&_ecf))
}

// NewLTV returns a new LTV client.
func NewLTV(appender *PdfAppender) (*LTV, error) {
	_cedae := appender.Reader.DSS
	if _cedae == nil {
		_cedae = NewDSS()
	}
	if _abgab := _cedae.GenerateHashMaps(); _abgab != nil {
		return nil, _abgab
	}
	return &LTV{CertClient: _fe.NewCertClient(), OCSPClient: _fe.NewOCSPClient(), CRLClient: _fe.NewCRLClient(), SkipExisting: true, _ebdgd: appender, _abca: _cedae}, nil
}

// NewPdfColorLab returns a new Lab color.
func NewPdfColorLab(l, a, b float64) *PdfColorLab { _gbbg := PdfColorLab{l, a, b}; return &_gbbg }

// NewXObjectForm creates a brand new XObject Form. Creates a new underlying PDF object stream primitive.
func NewXObjectForm() *XObjectForm {
	_cggaa := &XObjectForm{}
	_fgae := &_dg.PdfObjectStream{}
	_fgae.PdfObjectDictionary = _dg.MakeDict()
	_cggaa._ebaeb = _fgae
	return _cggaa
}

// GetAnnotations returns the list of page annotations for `page`. If not loaded attempts to load the
// annotations, otherwise returns the loaded list.
func (_baafe *PdfPage) GetAnnotations() ([]*PdfAnnotation, error) {
	if _baafe._cadgg != nil {
		return _baafe._cadgg, nil
	}
	if _baafe.Annots == nil {
		_baafe._cadgg = []*PdfAnnotation{}
		return nil, nil
	}
	if _baafe._cbbcc == nil {
		_baafe._cadgg = []*PdfAnnotation{}
		return nil, nil
	}
	_gdefc, _dcfc := _baafe._cbbcc.loadAnnotations(_baafe.Annots)
	if _dcfc != nil {
		return nil, _dcfc
	}
	if _gdefc == nil {
		_baafe._cadgg = []*PdfAnnotation{}
	}
	_baafe._cadgg = _gdefc
	return _baafe._cadgg, nil
}
func _daffeb() string {
	_geed := "\u0051\u0057\u0045\u0052\u0054\u0059\u0055\u0049\u004f\u0050\u0041S\u0044\u0046\u0047\u0048\u004a\u004b\u004c\u005a\u0058\u0043V\u0042\u004e\u004d"
	var _afeed _bc.Buffer
	for _ebcd := 0; _ebcd < 6; _ebcd++ {
		_afeed.WriteRune(rune(_geed[_cb.Intn(len(_geed))]))
	}
	return _afeed.String()
}
func _adabe(_acgcg *_dg.PdfObjectDictionary) (*PdfShadingPattern, error) {
	_aaggf := &PdfShadingPattern{}
	_gcfc := _acgcg.Get("\u0053h\u0061\u0064\u0069\u006e\u0067")
	if _gcfc == nil {
		_ag.Log.Debug("\u0053h\u0061d\u0069\u006e\u0067\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_dffe, _aeaf := _bgdgf(_gcfc)
	if _aeaf != nil {
		_ag.Log.Debug("\u0045r\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006eg\u0020s\u0068a\u0064\u0069\u006e\u0067\u003a\u0020\u0025v", _aeaf)
		return nil, _aeaf
	}
	_aaggf.Shading = _dffe
	if _ccbce := _acgcg.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _ccbce != nil {
		_cgbff, _efgba := _ccbce.(*_dg.PdfObjectArray)
		if !_efgba {
			_ag.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ccbce)
			return nil, _dg.ErrTypeError
		}
		_aaggf.Matrix = _cgbff
	}
	if _cbfbge := _acgcg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _cbfbge != nil {
		_aaggf.ExtGState = _cbfbge
	}
	return _aaggf, nil
}
func (_cagd *PdfReader) newPdfAnnotationTrapNetFromDict(_adbb *_dg.PdfObjectDictionary) (*PdfAnnotationTrapNet, error) {
	_adc := PdfAnnotationTrapNet{}
	return &_adc, nil
}

// ToPdfObject implements interface PdfModel.
func (_bcda *PdfAnnotationScreen) ToPdfObject() _dg.PdfObject {
	_bcda.PdfAnnotation.ToPdfObject()
	_abef := _bcda._cdf
	_eeab := _abef.PdfObject.(*_dg.PdfObjectDictionary)
	_eeab.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0053\u0063\u0072\u0065\u0065\u006e"))
	_eeab.SetIfNotNil("\u0054", _bcda.T)
	_eeab.SetIfNotNil("\u004d\u004b", _bcda.MK)
	_eeab.SetIfNotNil("\u0041", _bcda.A)
	_eeab.SetIfNotNil("\u0041\u0041", _bcda.AA)
	return _abef
}

// NewGrayImageFromGoImage creates a new grayscale unidoc Image from a golang Image.
func (_ecad DefaultImageHandler) NewGrayImageFromGoImage(goimg _gd.Image) (*Image, error) {
	_cbfeg := goimg.Bounds()
	_fcgeb := &Image{Width: int64(_cbfeg.Dx()), Height: int64(_cbfeg.Dy()), ColorComponents: 1, BitsPerComponent: 8}
	switch _eedd := goimg.(type) {
	case *_gd.Gray:
		if len(_eedd.Pix) != _cbfeg.Dx()*_cbfeg.Dy() {
			_ecab, _eeabg := _fc.GrayConverter.Convert(goimg)
			if _eeabg != nil {
				return nil, _eeabg
			}
			_fcgeb.Data = _ecab.Pix()
		} else {
			_fcgeb.Data = _eedd.Pix
		}
	case *_gd.Gray16:
		_fcgeb.BitsPerComponent = 16
		if len(_eedd.Pix) != _cbfeg.Dx()*_cbfeg.Dy()*2 {
			_dadda, _fbfbb := _fc.Gray16Converter.Convert(goimg)
			if _fbfbb != nil {
				return nil, _fbfbb
			}
			_fcgeb.Data = _dadda.Pix()
		} else {
			_fcgeb.Data = _eedd.Pix
		}
	case _fc.Image:
		_gefg := _eedd.Base()
		if _gefg.ColorComponents == 1 {
			_fcgeb.BitsPerComponent = int64(_gefg.BitsPerComponent)
			_fcgeb.Data = _gefg.Data
			return _fcgeb, nil
		}
		_fdec, _eefac := _fc.GrayConverter.Convert(goimg)
		if _eefac != nil {
			return nil, _eefac
		}
		_fcgeb.Data = _fdec.Pix()
	default:
		_dagg, _cccfa := _fc.GrayConverter.Convert(goimg)
		if _cccfa != nil {
			return nil, _cccfa
		}
		_fcgeb.Data = _dagg.Pix()
	}
	return _fcgeb, nil
}

// SetImageHandler sets the image handler used by the package.
func SetImageHandler(imgHandling ImageHandler) { ImageHandling = imgHandling }

// HasXObjectByName checks if an XObject with a specified keyName is defined.
func (_dbafe *PdfPageResources) HasXObjectByName(keyName _dg.PdfObjectName) bool {
	_dadggb, _ := _dbafe.GetXObjectByName(keyName)
	return _dadggb != nil
}
func (_eggc *PdfReader) newPdfAnnotationFileAttachmentFromDict(_adeg *_dg.PdfObjectDictionary) (*PdfAnnotationFileAttachment, error) {
	_deca := PdfAnnotationFileAttachment{}
	_bcgf, _gffg := _eggc.newPdfAnnotationMarkupFromDict(_adeg)
	if _gffg != nil {
		return nil, _gffg
	}
	_deca.PdfAnnotationMarkup = _bcgf
	_deca.FS = _adeg.Get("\u0046\u0053")
	_deca.Name = _adeg.Get("\u004e\u0061\u006d\u0065")
	return &_deca, nil
}

// SetNamedDestinations sets the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_cdcee *PdfWriter) SetNamedDestinations(dests _dg.PdfObject) error {
	if dests == nil {
		return nil
	}
	_ag.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0044\u0065\u0073\u0074\u0073\u002e\u002e\u002e")
	_cdcee._ecdf.Set("\u0044\u0065\u0073t\u0073", dests)
	return _cdcee.addObjects(dests)
}

// GetContainingPdfObject returns the container of the pattern object (indirect object).
func (_ggdb *PdfPattern) GetContainingPdfObject() _dg.PdfObject { return _ggdb._eacce }

// ToPdfObject returns the PDF representation of the colorspace.
func (_cfga *PdfColorspaceDeviceRGB) ToPdfObject() _dg.PdfObject {
	return _dg.MakeName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
}

// GetCapHeight returns the CapHeight of the font `descriptor`.
func (_ceae *PdfFontDescriptor) GetCapHeight() (float64, error) {
	return _dg.GetNumberAsFloat(_ceae.CapHeight)
}

// Encoder returns the font's text encoder.
func (_dcfd pdfCIDFontType0) Encoder() _bd.TextEncoder { return _dcfd._afbff }
func _dffca(_gdgb _dg.PdfObject) (*PdfColorspaceSpecialPattern, error) {
	_ag.Log.Trace("\u004e\u0065\u0077\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u0043\u0053\u0020\u0066\u0072\u006fm\u0020\u006f\u0062\u006a\u003a\u0020\u0025\u0073\u0020\u0025\u0054", _gdgb.String(), _gdgb)
	_gbddf := NewPdfColorspaceSpecialPattern()
	if _gcccc, _cgff := _gdgb.(*_dg.PdfIndirectObject); _cgff {
		_gbddf._cbff = _gcccc
	}
	_gdgb = _dg.TraceToDirectObject(_gdgb)
	if _ffcd, _gdca := _gdgb.(*_dg.PdfObjectName); _gdca {
		if *_ffcd != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			return nil, _b.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
		return _gbddf, nil
	}
	_dccdc, _fcegc := _gdgb.(*_dg.PdfObjectArray)
	if !_fcegc {
		_ag.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061t\u0074\u0065\u0072\u006e\u0020\u0043\u0053 \u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0023\u0076", _gdgb)
		return nil, _b.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0050\u0061\u0074\u0074e\u0072n\u0020C\u0053\u0020\u006f\u0062\u006a\u0065\u0063t")
	}
	if _dccdc.Len() != 1 && _dccdc.Len() != 2 {
		_ag.Log.Error("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079\u003a %\u0023\u0076", _dccdc)
		return nil, _b.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065r\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_gdgb = _dccdc.Get(0)
	if _cddc, _gffc := _gdgb.(*_dg.PdfObjectName); _gffc {
		if *_cddc != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			_ag.Log.Error("\u0049\u006e\u0076al\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020C\u0053 \u0061r\u0072\u0061\u0079\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0023\u0076", _cddc)
			return nil, _b.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
	}
	if _dccdc.Len() > 1 {
		_gdgb = _dccdc.Get(1)
		_gdgb = _dg.TraceToDirectObject(_gdgb)
		_fadf, _gaab := NewPdfColorspaceFromPdfObject(_gdgb)
		if _gaab != nil {
			return nil, _gaab
		}
		_gbddf.UnderlyingCS = _fadf
	}
	_ag.Log.Trace("R\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0077i\u0074\u0068\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079in\u0067\u0020\u0063s\u003a \u0025\u0054", _gbddf.UnderlyingCS)
	return _gbddf, nil
}

// PdfFieldButton represents a button field which includes push buttons, checkboxes, and radio buttons.
type PdfFieldButton struct {
	*PdfField
	Opt   *_dg.PdfObjectArray
	_gddc *Image
}

// NewPdfFontFromTTFFile loads a TTF font file and returns a PdfFont type
// that can be used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTFFile.
func NewPdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_ebdd, _bdbg := _eb.Open(filePath)
	if _bdbg != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020T\u0054F\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _bdbg)
		return nil, _bdbg
	}
	defer _ebdd.Close()
	return NewPdfFontFromTTF(_ebdd)
}

// UpdatePage updates the `page` in the new revision if it has changed.
func (_ffea *PdfAppender) UpdatePage(page *PdfPage) { _ffea.updateObjectsDeep(page.ToPdfObject(), nil) }

// GetPageAsIndirectObject returns the page as a dictionary within an PdfIndirectObject.
func (_bbcfg *PdfPage) GetPageAsIndirectObject() *_dg.PdfIndirectObject { return _bbcfg._cggbe }

// PdfColorspaceSpecialSeparation is a Separation colorspace.
// At the moment the colour space is set to a Separation space, the conforming reader shall determine whether the
// device has an available colorant (e.g. dye) corresponding to the name of the requested space. If so, the conforming
// reader shall ignore the alternateSpace and tintTransform parameters; subsequent painting operations within the
// space shall apply the designated colorant directly, according to the tint values supplied.
//
// Format: [/Separation name alternateSpace tintTransform]
type PdfColorspaceSpecialSeparation struct {
	ColorantName   *_dg.PdfObjectName
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	_fecg          *_dg.PdfIndirectObject
}

const (
	TrappedUnknown PdfInfoTrapped = "\u0055n\u006b\u006e\u006f\u0077\u006e"
	TrappedTrue    PdfInfoTrapped = "\u0054\u0072\u0075\u0065"
	TrappedFalse   PdfInfoTrapped = "\u0046\u0061\u006cs\u0065"
)

// SetFontByName sets the font specified by keyName to the given object.
func (_ffffd *PdfPageResources) SetFontByName(keyName _dg.PdfObjectName, obj _dg.PdfObject) error {
	if _ffffd.Font == nil {
		_ffffd.Font = _dg.MakeDict()
	}
	_cgcgb, _ggfe := _dg.TraceToDirectObject(_ffffd.Font).(*_dg.PdfObjectDictionary)
	if !_ggfe {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _dg.TraceToDirectObject(_ffffd.Font))
		return _dg.ErrTypeError
	}
	_cgcgb.Set(keyName, obj)
	return nil
}

// SetImage updates XObject Image with new image data.
func (_dbfdf *XObjectImage) SetImage(img *Image, cs PdfColorspace) error {
	_dbfdf.Filter.UpdateParams(img.GetParamsDict())
	_fecbb, _dfgdg := _dbfdf.Filter.EncodeBytes(img.Data)
	if _dfgdg != nil {
		return _dfgdg
	}
	_dbfdf.Stream = _fecbb
	_gdba := img.Width
	_dbfdf.Width = &_gdba
	_ceadd := img.Height
	_dbfdf.Height = &_ceadd
	_cabbd := img.BitsPerComponent
	_dbfdf.BitsPerComponent = &_cabbd
	if cs == nil {
		if img.ColorComponents == 1 {
			_dbfdf.ColorSpace = NewPdfColorspaceDeviceGray()
		} else if img.ColorComponents == 3 {
			_dbfdf.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			_dbfdf.ColorSpace = NewPdfColorspaceDeviceCMYK()
		} else {
			return _bf.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_dbfdf.ColorSpace = cs
	}
	return nil
}

// SetContentStream sets the pattern cell's content stream.
func (_geab *PdfTilingPattern) SetContentStream(content []byte, encoder _dg.StreamEncoder) error {
	_bbcec, _feaef := _geab._eacce.(*_dg.PdfObjectStream)
	if !_feaef {
		_ag.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _geab._eacce)
		return _dg.ErrTypeError
	}
	if encoder == nil {
		encoder = _dg.NewRawEncoder()
	}
	_aaeg := _bbcec.PdfObjectDictionary
	_gdcab := encoder.MakeStreamDict()
	_aaeg.Merge(_gdcab)
	_gbcdf, _fggcc := encoder.EncodeBytes(content)
	if _fggcc != nil {
		return _fggcc
	}
	_aaeg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dg.MakeInteger(int64(len(_gbcdf))))
	_bbcec.Stream = _gbcdf
	return nil
}

// ColorFromPdfObjects returns a new PdfColor based on input color components. The input PdfObjects should
// be numeric.
func (_cabd *PdfColorspaceDeviceN) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != _cabd.GetNumComponents() {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cbbb, _ddggg := _dg.GetNumbersAsFloat(objects)
	if _ddggg != nil {
		return nil, _ddggg
	}
	return _cabd.ColorFromFloats(_cbbb)
}

// PdfTilingPattern is a Tiling pattern that consists of repetitions of a pattern cell with defined intervals.
// It is a type 1 pattern. (PatternType = 1).
// A tiling pattern is represented by a stream object, where the stream content is
// a content stream that describes the pattern cell.
type PdfTilingPattern struct {
	*PdfPattern
	PaintType  *_dg.PdfObjectInteger
	TilingType *_dg.PdfObjectInteger
	BBox       *PdfRectangle
	XStep      *_dg.PdfObjectFloat
	YStep      *_dg.PdfObjectFloat
	Resources  *PdfPageResources
	Matrix     *_dg.PdfObjectArray
}

// ToPdfObject implements interface PdfModel.
func (_bcae *PdfAnnotationText) ToPdfObject() _dg.PdfObject {
	_bcae.PdfAnnotation.ToPdfObject()
	_edgf := _bcae._cdf
	_cfd := _edgf.PdfObject.(*_dg.PdfObjectDictionary)
	if _bcae.PdfAnnotationMarkup != nil {
		_bcae.PdfAnnotationMarkup.appendToPdfDictionary(_cfd)
	}
	_cfd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0054\u0065\u0078\u0074"))
	_cfd.SetIfNotNil("\u004f\u0070\u0065\u006e", _bcae.Open)
	_cfd.SetIfNotNil("\u004e\u0061\u006d\u0065", _bcae.Name)
	_cfd.SetIfNotNil("\u0053\u0074\u0061t\u0065", _bcae.State)
	_cfd.SetIfNotNil("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c", _bcae.StateModel)
	return _edgf
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_ffgg *PdfColorspaceSpecialSeparation) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ffbf := vals[0]
	_abedf := []float64{_ffbf}
	_gggga, _gbaa := _ffgg.TintTransform.Evaluate(_abedf)
	if _gbaa != nil {
		_ag.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0065: \u0025\u0076", _gbaa)
		_ag.Log.Trace("\u0054\u0069\u006e\u0074 t\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _ffgg.TintTransform)
		return nil, _gbaa
	}
	_ag.Log.Trace("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0043\u006f\u006c\u006fr\u0046\u0072\u006f\u006d\u0046\u006c\u006f\u0061\u0074\u0073\u0028\u0025\u002bv\u0029\u0020\u006f\u006e\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061te\u0053\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0023\u0076", _gggga, _ffgg.AlternateSpace)
	_febef, _gbaa := _ffgg.AlternateSpace.ColorFromFloats(_gggga)
	if _gbaa != nil {
		_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u002c\u0020\u0066a\u0069\u006c\u0065d \u0074\u006f\u0020\u0065\u0076\u0061l\u0075\u0061\u0074\u0065\u0020\u0069\u006e\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u003a \u0025\u0076", _gbaa)
		return nil, _gbaa
	}
	return _febef, nil
}

// PdfColorDeviceGray represents a grayscale color value that shall be represented by a single number in the
// range 0.0 to 1.0 where 0.0 corresponds to black and 1.0 to white.
type PdfColorDeviceGray float64

func (_cbce *PdfReader) newPdfActionImportDataFromDict(_cfea *_dg.PdfObjectDictionary) (*PdfActionImportData, error) {
	_gde, _fbc := _bccf(_cfea.Get("\u0046"))
	if _fbc != nil {
		return nil, _fbc
	}
	return &PdfActionImportData{F: _gde}, nil
}
func (_cgbea *PdfPage) setContainer(_eggfa *_dg.PdfIndirectObject) {
	_eggfa.PdfObject = _cgbea._bfdge
	_cgbea._cggbe = _eggfa
}

// FieldFlag represents form field flags. Some of the flags can apply to all types of fields whereas other
// flags are specific.
type FieldFlag uint32

// BorderEffect represents a border effect (Table 167 p. 395).
type BorderEffect int

// NewPdfColorspaceDeviceCMYK returns a new CMYK32 colorspace object.
func NewPdfColorspaceDeviceCMYK() *PdfColorspaceDeviceCMYK { return &PdfColorspaceDeviceCMYK{} }
func (_aabdgg *LTV) generateVRIKey(_dcga *PdfSignature) (string, error) {
	_afea, _afdg := _cbeaf(_dcga.Contents.Bytes())
	if _afdg != nil {
		return "", _afdg
	}
	return _ga.ToUpper(_be.EncodeToString(_afea)), nil
}

// GetOutlineTree returns the outline tree.
func (_cgfa *PdfReader) GetOutlineTree() *PdfOutlineTreeNode { return _cgfa._fcfc }

// ConvertToBinary converts current image into binary (bi-level) format.
// Binary images are composed of single bits per pixel (only black or white).
// If provided image has more color components, then it would be converted into binary image using
// histogram auto threshold function.
func (_ebge *Image) ConvertToBinary() error {
	if _ebge.ColorComponents == 1 && _ebge.BitsPerComponent == 1 {
		return nil
	}
	_ebdae, _bccdg := _ebge.ToGoImage()
	if _bccdg != nil {
		return _bccdg
	}
	_fdbdc, _bccdg := _fc.MonochromeConverter.Convert(_ebdae)
	if _bccdg != nil {
		return _bccdg
	}
	_ebge.Data = _fdbdc.Base().Data
	_ebge._dgeb, _bccdg = _fc.ScaleAlphaToMonochrome(_ebge._dgeb, int(_ebge.Width), int(_ebge.Height))
	if _bccdg != nil {
		return _bccdg
	}
	_ebge.BitsPerComponent = 1
	_ebge.ColorComponents = 1
	_ebge._gfbb = nil
	return nil
}

// HasPatternByName checks whether a pattern object is defined by the specified keyName.
func (_cffbc *PdfPageResources) HasPatternByName(keyName _dg.PdfObjectName) bool {
	_, _fbcfc := _cffbc.GetPatternByName(keyName)
	return _fbcfc
}

// Has checks if flag fl is set in flag and returns true if so, false otherwise.
func (_cgad FieldFlag) Has(fl FieldFlag) bool { return (_cgad.Mask() & fl.Mask()) > 0 }

// PartialName returns the partial name of the field.
func (_bgbf *PdfField) PartialName() string {
	_ceecb := ""
	if _bgbf.T != nil {
		_ceecb = _bgbf.T.Decoded()
	} else {
		_ag.Log.Debug("\u0046\u0069el\u0064\u0020\u006di\u0073\u0073\u0069\u006eg T\u0020fi\u0065\u006c\u0064\u0020\u0028\u0069\u006eco\u006d\u0070\u0061\u0074\u0069\u0062\u006ce\u0029")
	}
	return _ceecb
}

// AddExtGState add External Graphics State (GState). The gsDict can be specified
// either directly as a dictionary or an indirect object containing a dictionary.
func (_acfc *PdfPageResources) AddExtGState(gsName _dg.PdfObjectName, gsDict _dg.PdfObject) error {
	if _acfc.ExtGState == nil {
		_acfc.ExtGState = _dg.MakeDict()
	}
	_gcded := _acfc.ExtGState
	_afffc, _bcece := _dg.TraceToDirectObject(_gcded).(*_dg.PdfObjectDictionary)
	if !_bcece {
		_ag.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020e\u0072r\u006f\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u002f\u0025\u0054\u0029", _gcded, _dg.TraceToDirectObject(_gcded))
		return _dg.ErrTypeError
	}
	_afffc.Set(gsName, gsDict)
	return nil
}

// NewPdfOutline returns an initialized PdfOutline.
func NewPdfOutline() *PdfOutline {
	_debdd := &PdfOutline{_bbef: _dg.MakeIndirectObject(_dg.MakeDict())}
	_debdd._baddf = _debdd
	return _debdd
}

// StdFontName represents name of a standard font.
type StdFontName = _bbg.StdFontName

// ReaderOpts defines options for creating PdfReader instances.
type ReaderOpts struct {

	// Password password of the PDF file encryption.
	// Default: empty ("").
	Password string

	// LazyLoad set if the PDF file would be loaded using lazy-loading mode.
	// Default: true.
	LazyLoad bool

	// ComplianceMode set if parsed PDF file should contain meta information for the verifiers of the compliance standards like PDF/A.
	ComplianceMode bool
}

func (_cccg *PdfReader) newPdfAnnotationFromIndirectObject(_bafb *_dg.PdfIndirectObject) (*PdfAnnotation, error) {
	_bbe, _bba := _bafb.PdfObject.(*_dg.PdfObjectDictionary)
	if !_bba {
		return nil, _b.Errorf("\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _cccg._cadfa.GetModelFromPrimitive(_bbe); model != nil {
		_ebae, _dca := model.(*PdfAnnotation)
		if !_dca {
			return nil, _b.Errorf("\u0063\u0061\u0063\u0068\u0065\u0064 \u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050D\u0046\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e")
		}
		return _ebae, nil
	}
	_dcccc := &PdfAnnotation{}
	_dcccc._cdf = _bafb
	_cccg._cadfa.Register(_bbe, _dcccc)
	if _gggg := _bbe.Get("\u0054\u0079\u0070\u0065"); _gggg != nil {
		_fcea, _cadg := _gggg.(*_dg.PdfObjectName)
		if !_cadg {
			_ag.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _gggg)
		} else {
			if *_fcea != "\u0041\u006e\u006eo\u0074" {
				_ag.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070\u0065\u0063\u0074\u0065d\u0020\u0054\u0079\u0070\u0065\u0020\u0021=\u0020\u0041\u006e\u006e\u006f\u0074\u0020\u0028\u0025\u0073\u0029", *_fcea)
			}
		}
	}
	if _fcec := _bbe.Get("\u0052\u0065\u0063\u0074"); _fcec != nil {
		_dcccc.Rect = _fcec
	}
	if _aega := _bbe.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _aega != nil {
		_dcccc.Contents = _aega
	}
	if _faeg := _bbe.Get("\u0050"); _faeg != nil {
		_dcccc.P = _faeg
	}
	if _ffd := _bbe.Get("\u004e\u004d"); _ffd != nil {
		_dcccc.NM = _ffd
	}
	if _bfcd := _bbe.Get("\u004d"); _bfcd != nil {
		_dcccc.M = _bfcd
	}
	if _ffac := _bbe.Get("\u0046"); _ffac != nil {
		_dcccc.F = _ffac
	}
	if _ddc := _bbe.Get("\u0041\u0050"); _ddc != nil {
		_dcccc.AP = _ddc
	}
	if _bec := _bbe.Get("\u0041\u0053"); _bec != nil {
		_dcccc.AS = _bec
	}
	if _cbfb := _bbe.Get("\u0042\u006f\u0072\u0064\u0065\u0072"); _cbfb != nil {
		_dcccc.Border = _cbfb
	}
	if _bdg := _bbe.Get("\u0043"); _bdg != nil {
		_dcccc.C = _bdg
	}
	if _gff := _bbe.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074"); _gff != nil {
		_dcccc.StructParent = _gff
	}
	if _fac := _bbe.Get("\u004f\u0043"); _fac != nil {
		_dcccc.OC = _fac
	}
	_agg := _bbe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")
	if _agg == nil {
		_ag.Log.Debug("\u0057\u0041\u0052\u004e\u0049\u004e\u0047:\u0020\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079 \u0069s\u0073\u0075\u0065\u0020\u002d\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073u\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074\u0079p\u0065")
		_dcccc._egcg = nil
		return _dcccc, nil
	}
	_deb, _feea := _agg.(*_dg.PdfObjectName)
	if !_feea {
		_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065 !\u003d\u0020n\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _agg)
		return nil, _b.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d n\u0061\u006d\u0065 \u0028%\u0054\u0029", _agg)
	}
	switch *_deb {
	case "\u0054\u0065\u0078\u0074":
		_ace, _dbe := _cccg.newPdfAnnotationTextFromDict(_bbe)
		if _dbe != nil {
			return nil, _dbe
		}
		_ace.PdfAnnotation = _dcccc
		_dcccc._egcg = _ace
		return _dcccc, nil
	case "\u004c\u0069\u006e\u006b":
		_cgcd, _adfa := _cccg.newPdfAnnotationLinkFromDict(_bbe)
		if _adfa != nil {
			return nil, _adfa
		}
		_cgcd.PdfAnnotation = _dcccc
		_dcccc._egcg = _cgcd
		return _dcccc, nil
	case "\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074":
		_aag, _gcbe := _cccg.newPdfAnnotationFreeTextFromDict(_bbe)
		if _gcbe != nil {
			return nil, _gcbe
		}
		_aag.PdfAnnotation = _dcccc
		_dcccc._egcg = _aag
		return _dcccc, nil
	case "\u004c\u0069\u006e\u0065":
		_dccg, _ffb := _cccg.newPdfAnnotationLineFromDict(_bbe)
		if _ffb != nil {
			return nil, _ffb
		}
		_dccg.PdfAnnotation = _dcccc
		_dcccc._egcg = _dccg
		_ag.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041N\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0061\u006e\u006eo\u0074\u0020\u0028\u0025\u0054\u0029\u003a \u0025\u002b\u0076\u000a", _dcccc, _dcccc)
		_ag.Log.Trace("\u004c\u0049\u004eE\u0020\u0041\u004e\u004eO\u0054\u0041\u0054\u0049\u004f\u004e\u003a \u0063\u0074\u0078\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u002b\u0076\u000a", _dccg, _dccg)
		_ag.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041\u004e\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u0020\u004d\u0061\u0072\u006b\u0075\u0070\u003a\u0020c\u0074\u0078\u0020\u0028\u0025T\u0029\u003a \u0025\u002b\u0076\u000a", _dccg.PdfAnnotationMarkup, _dccg.PdfAnnotationMarkup)
		return _dcccc, nil
	case "\u0053\u0071\u0075\u0061\u0072\u0065":
		_acee, _agag := _cccg.newPdfAnnotationSquareFromDict(_bbe)
		if _agag != nil {
			return nil, _agag
		}
		_acee.PdfAnnotation = _dcccc
		_dcccc._egcg = _acee
		return _dcccc, nil
	case "\u0043\u0069\u0072\u0063\u006c\u0065":
		_fdbe, _cbe := _cccg.newPdfAnnotationCircleFromDict(_bbe)
		if _cbe != nil {
			return nil, _cbe
		}
		_fdbe.PdfAnnotation = _dcccc
		_dcccc._egcg = _fdbe
		return _dcccc, nil
	case "\u0050o\u006c\u0079\u0067\u006f\u006e":
		_dff, _fdf := _cccg.newPdfAnnotationPolygonFromDict(_bbe)
		if _fdf != nil {
			return nil, _fdf
		}
		_dff.PdfAnnotation = _dcccc
		_dcccc._egcg = _dff
		return _dcccc, nil
	case "\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065":
		_ddab, _fda := _cccg.newPdfAnnotationPolyLineFromDict(_bbe)
		if _fda != nil {
			return nil, _fda
		}
		_ddab.PdfAnnotation = _dcccc
		_dcccc._egcg = _ddab
		return _dcccc, nil
	case "\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t":
		_cbde, _ffec := _cccg.newPdfAnnotationHighlightFromDict(_bbe)
		if _ffec != nil {
			return nil, _ffec
		}
		_cbde.PdfAnnotation = _dcccc
		_dcccc._egcg = _cbde
		return _dcccc, nil
	case "\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee":
		_bgda, _ebcb := _cccg.newPdfAnnotationUnderlineFromDict(_bbe)
		if _ebcb != nil {
			return nil, _ebcb
		}
		_bgda.PdfAnnotation = _dcccc
		_dcccc._egcg = _bgda
		return _dcccc, nil
	case "\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079":
		_accd, _fbg := _cccg.newPdfAnnotationSquigglyFromDict(_bbe)
		if _fbg != nil {
			return nil, _fbg
		}
		_accd.PdfAnnotation = _dcccc
		_dcccc._egcg = _accd
		return _dcccc, nil
	case "\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t":
		_bffc, _eag := _cccg.newPdfAnnotationStrikeOut(_bbe)
		if _eag != nil {
			return nil, _eag
		}
		_bffc.PdfAnnotation = _dcccc
		_dcccc._egcg = _bffc
		return _dcccc, nil
	case "\u0043\u0061\u0072e\u0074":
		_aaa, _bcdf := _cccg.newPdfAnnotationCaretFromDict(_bbe)
		if _bcdf != nil {
			return nil, _bcdf
		}
		_aaa.PdfAnnotation = _dcccc
		_dcccc._egcg = _aaa
		return _dcccc, nil
	case "\u0053\u0074\u0061m\u0070":
		_caac, _cbcf := _cccg.newPdfAnnotationStampFromDict(_bbe)
		if _cbcf != nil {
			return nil, _cbcf
		}
		_caac.PdfAnnotation = _dcccc
		_dcccc._egcg = _caac
		return _dcccc, nil
	case "\u0049\u006e\u006b":
		_fdac, _dcca := _cccg.newPdfAnnotationInkFromDict(_bbe)
		if _dcca != nil {
			return nil, _dcca
		}
		_fdac.PdfAnnotation = _dcccc
		_dcccc._egcg = _fdac
		return _dcccc, nil
	case "\u0050\u006f\u0070u\u0070":
		_fbe, _dadg := _cccg.newPdfAnnotationPopupFromDict(_bbe)
		if _dadg != nil {
			return nil, _dadg
		}
		_fbe.PdfAnnotation = _dcccc
		_dcccc._egcg = _fbe
		return _dcccc, nil
	case "\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074":
		_eded, _bcfg := _cccg.newPdfAnnotationFileAttachmentFromDict(_bbe)
		if _bcfg != nil {
			return nil, _bcfg
		}
		_eded.PdfAnnotation = _dcccc
		_dcccc._egcg = _eded
		return _dcccc, nil
	case "\u0053\u006f\u0075n\u0064":
		_fca, _ebe := _cccg.newPdfAnnotationSoundFromDict(_bbe)
		if _ebe != nil {
			return nil, _ebe
		}
		_fca.PdfAnnotation = _dcccc
		_dcccc._egcg = _fca
		return _dcccc, nil
	case "\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a":
		_dccb, _eaef := _cccg.newPdfAnnotationRichMediaFromDict(_bbe)
		if _eaef != nil {
			return nil, _eaef
		}
		_dccb.PdfAnnotation = _dcccc
		_dcccc._egcg = _dccb
		return _dcccc, nil
	case "\u004d\u006f\u0076i\u0065":
		_gbc, _fece := _cccg.newPdfAnnotationMovieFromDict(_bbe)
		if _fece != nil {
			return nil, _fece
		}
		_gbc.PdfAnnotation = _dcccc
		_dcccc._egcg = _gbc
		return _dcccc, nil
	case "\u0053\u0063\u0072\u0065\u0065\u006e":
		_gbcd, _cadge := _cccg.newPdfAnnotationScreenFromDict(_bbe)
		if _cadge != nil {
			return nil, _cadge
		}
		_gbcd.PdfAnnotation = _dcccc
		_dcccc._egcg = _gbcd
		return _dcccc, nil
	case "\u0057\u0069\u0064\u0067\u0065\u0074":
		_cadd, _ccdde := _cccg.newPdfAnnotationWidgetFromDict(_bbe)
		if _ccdde != nil {
			return nil, _ccdde
		}
		_cadd.PdfAnnotation = _dcccc
		_dcccc._egcg = _cadd
		return _dcccc, nil
	case "P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b":
		_beec, _dfcg := _cccg.newPdfAnnotationPrinterMarkFromDict(_bbe)
		if _dfcg != nil {
			return nil, _dfcg
		}
		_beec.PdfAnnotation = _dcccc
		_dcccc._egcg = _beec
		return _dcccc, nil
	case "\u0054r\u0061\u0070\u004e\u0065\u0074":
		_bbad, _bfa := _cccg.newPdfAnnotationTrapNetFromDict(_bbe)
		if _bfa != nil {
			return nil, _bfa
		}
		_bbad.PdfAnnotation = _dcccc
		_dcccc._egcg = _bbad
		return _dcccc, nil
	case "\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k":
		_dbgc, _ffc := _cccg.newPdfAnnotationWatermarkFromDict(_bbe)
		if _ffc != nil {
			return nil, _ffc
		}
		_dbgc.PdfAnnotation = _dcccc
		_dcccc._egcg = _dbgc
		return _dcccc, nil
	case "\u0033\u0044":
		_eef, _bade := _cccg.newPdfAnnotation3DFromDict(_bbe)
		if _bade != nil {
			return nil, _bade
		}
		_eef.PdfAnnotation = _dcccc
		_dcccc._egcg = _eef
		return _dcccc, nil
	case "\u0050\u0072\u006f\u006a\u0065\u0063\u0074\u0069\u006f\u006e":
		_efga, _afcf := _cccg.newPdfAnnotationProjectionFromDict(_bbe)
		if _afcf != nil {
			return nil, _afcf
		}
		_efga.PdfAnnotation = _dcccc
		_dcccc._egcg = _efga
		return _dcccc, nil
	case "\u0052\u0065\u0064\u0061\u0063\u0074":
		_ddbb, _fcce := _cccg.newPdfAnnotationRedactFromDict(_bbe)
		if _fcce != nil {
			return nil, _fcce
		}
		_ddbb.PdfAnnotation = _dcccc
		_dcccc._egcg = _ddbb
		return _dcccc, nil
	}
	_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069\u006fn\u003a\u0020\u0025\u0073", *_deb)
	return nil, nil
}

// ToPdfObject returns the PDF representation of the outline tree node.
func (_beac *PdfOutlineTreeNode) ToPdfObject() _dg.PdfObject { return _beac.GetContext().ToPdfObject() }

// ToPdfObject implements interface PdfModel.
func (_ddceg *PdfAnnotationSquiggly) ToPdfObject() _dg.PdfObject {
	_ddceg.PdfAnnotation.ToPdfObject()
	_dbcc := _ddceg._cdf
	_ceba := _dbcc.PdfObject.(*_dg.PdfObjectDictionary)
	_ddceg.PdfAnnotationMarkup.appendToPdfDictionary(_ceba)
	_ceba.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079"))
	_ceba.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _ddceg.QuadPoints)
	return _dbcc
}
func (_gefd fontCommon) fontFlags() int {
	if _gefd._ccfb == nil {
		return 0
	}
	return _gefd._ccfb._gcdf
}

// PdfAppender appends new PDF content to an existing PDF document via incremental updates.
type PdfAppender struct {
	_cga   _cf.ReadSeeker
	_gfab  *_dg.PdfParser
	_debg  *PdfReader
	Reader *PdfReader
	_ggdd  []*PdfPage
	_afab  *PdfAcroForm
	_cfdd  *DSS
	_accbd *Permissions
	_ceef  _dg.XrefTable
	_bcdg  int64
	_ceecc int
	_dgfd  []_dg.PdfObject
	_efa   map[_dg.PdfObject]struct{}
	_gebe  map[_dg.PdfObject]int64
	_eede  map[_dg.PdfObject]struct{}
	_ebaa  map[_dg.PdfObject]struct{}
	_dcaf  int64
	_fdbb  bool
	_acb   string
	_ccfc  *EncryptOptions
	_fega  *PdfInfo
}

// GetType returns the button field type which returns one of the following
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
func (_dcebc *PdfFieldButton) GetType() ButtonType {
	_febf := ButtonTypeCheckbox
	if _dcebc.Ff != nil {
		if (uint32(*_dcebc.Ff) & FieldFlagPushbutton.Mask()) > 0 {
			_febf = ButtonTypePush
		} else if (uint32(*_dcebc.Ff) & FieldFlagRadio.Mask()) > 0 {
			_febf = ButtonTypeRadio
		}
	}
	return _febf
}
func _fbff(_afegg *_dg.PdfObjectDictionary, _aeed *fontCommon) (*pdfCIDFontType0, error) {
	if _aeed._bcga != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" {
		_ag.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0030\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _aeed)
		return nil, _dg.ErrRangeError
	}
	_bfcbf := _ggbff(_aeed)
	_bcaec, _ffga := _dg.GetDict(_afegg.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_ffga {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _aeed)
		return nil, ErrRequiredAttributeMissing
	}
	_bfcbf.CIDSystemInfo = _bcaec
	_bfcbf.DW = _afegg.Get("\u0044\u0057")
	_bfcbf.W = _afegg.Get("\u0057")
	_bfcbf.DW2 = _afegg.Get("\u0044\u0057\u0032")
	_bfcbf.W2 = _afegg.Get("\u0057\u0032")
	_bfcbf._gfdca = 1000.0
	if _cbbba, _bfaf := _dg.GetNumberAsFloat(_bfcbf.DW); _bfaf == nil {
		_bfcbf._gfdca = _cbbba
	}
	_bfga, _bbbcea := _gacbe(_bfcbf.W)
	if _bbbcea != nil {
		return nil, _bbbcea
	}
	if _bfga == nil {
		_bfga = map[_bd.CharCode]float64{}
	}
	_bfcbf._gfcba = _bfga
	return _bfcbf, nil
}
func _fgcbf() string {
	_fgefgf.Lock()
	defer _fgefgf.Unlock()
	if len(_cebb) > 0 {
		return _cebb
	}
	return "\u0055n\u0069\u0044\u006f\u0063 \u002d\u0020\u0068\u0074\u0074p\u003a/\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069o"
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_feae *PdfPageResourcesColorspaces) ToPdfObject() _dg.PdfObject {
	_ffggfa := _dg.MakeDict()
	for _, _fgbed := range _feae.Names {
		_ffggfa.Set(_dg.PdfObjectName(_fgbed), _feae.Colorspaces[_fgbed].ToPdfObject())
	}
	if _feae._dcebf != nil {
		_feae._dcebf.PdfObject = _ffggfa
		return _feae._dcebf
	}
	return _ffggfa
}
func (_gbcfcg *PdfWriter) addObject(_dceag _dg.PdfObject) bool {
	_cdagb := _gbcfcg.hasObject(_dceag)
	if !_cdagb {
		_cfafe := _dg.ResolveReferencesDeep(_dceag, _gbcfcg._cffaa)
		if _cfafe != nil {
			_ag.Log.Debug("E\u0052R\u004f\u0052\u003a\u0020\u0025\u0076\u0020\u002d \u0073\u006b\u0069\u0070pi\u006e\u0067", _cfafe)
		}
		_gbcfcg._agaba = append(_gbcfcg._agaba, _dceag)
		_gbcfcg._fdbfa[_dceag] = struct{}{}
		return true
	}
	return false
}

// AddFont adds a font dictionary to the Font resources.
func (_ecbfc *PdfPage) AddFont(name _dg.PdfObjectName, font _dg.PdfObject) error {
	if _ecbfc.Resources == nil {
		_ecbfc.Resources = NewPdfPageResources()
	}
	if _ecbfc.Resources.Font == nil {
		_ecbfc.Resources.Font = _dg.MakeDict()
	}
	_gegaa, _bcdd := _dg.TraceToDirectObject(_ecbfc.Resources.Font).(*_dg.PdfObjectDictionary)
	if !_bcdd {
		_ag.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0066\u006f\u006et \u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a \u0025\u0076", _dg.TraceToDirectObject(_ecbfc.Resources.Font))
		return _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gegaa.Set(name, font)
	return nil
}
func (_caacf *PdfWriter) makeOffSetReference(_dgdbf int64) {
	_cdeaa := _b.Sprintf("\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u000a\u0025\u0064\u000a", _dgdbf)
	_caacf.writeString(_cdeaa)
	_caacf.writeString("\u0025\u0025\u0045\u004f\u0046\u000a")
}
func (_gfged *PdfWriter) copyObject(_cddfe _dg.PdfObject, _adce map[_dg.PdfObject]_dg.PdfObject, _ceegc map[_dg.PdfObject]struct{}, _dcdfb bool) _dg.PdfObject {
	_dbcf := !_gfged._bbac && _ceegc != nil
	if _acbfc, _bfbggc := _adce[_cddfe]; _bfbggc {
		if _dbcf && !_dcdfb {
			delete(_ceegc, _cddfe)
		}
		return _acbfc
	}
	if _cddfe == nil {
		_efgbc := _dg.MakeNull()
		return _efgbc
	}
	_dffcf := _cddfe
	switch _bcfec := _cddfe.(type) {
	case *_dg.PdfObjectArray:
		_dagff := _dg.MakeArray()
		_dffcf = _dagff
		_adce[_cddfe] = _dffcf
		for _, _cfbdd := range _bcfec.Elements() {
			_dagff.Append(_gfged.copyObject(_cfbdd, _adce, _ceegc, _dcdfb))
		}
	case *_dg.PdfObjectStreams:
		_fgbd := &_dg.PdfObjectStreams{PdfObjectReference: _bcfec.PdfObjectReference}
		_dffcf = _fgbd
		_adce[_cddfe] = _dffcf
		for _, _acabe := range _bcfec.Elements() {
			_fgbd.Append(_gfged.copyObject(_acabe, _adce, _ceegc, _dcdfb))
		}
	case *_dg.PdfObjectStream:
		_baebc := &_dg.PdfObjectStream{Stream: _bcfec.Stream, PdfObjectReference: _bcfec.PdfObjectReference}
		_dffcf = _baebc
		_adce[_cddfe] = _dffcf
		_baebc.PdfObjectDictionary = _gfged.copyObject(_bcfec.PdfObjectDictionary, _adce, _ceegc, _dcdfb).(*_dg.PdfObjectDictionary)
	case *_dg.PdfObjectDictionary:
		var _gfcaa bool
		if _dbcf && !_dcdfb {
			if _cdggc, _ := _dg.GetNameVal(_bcfec.Get("\u0054\u0079\u0070\u0065")); _cdggc == "\u0050\u0061\u0067\u0065" {
				_, _bbadc := _gfged._fegdf[_bcfec]
				_dcdfb = !_bbadc
				_gfcaa = _dcdfb
			}
		}
		_bfbea := _dg.MakeDict()
		_dffcf = _bfbea
		_adce[_cddfe] = _dffcf
		for _, _gfafd := range _bcfec.Keys() {
			_bfbea.Set(_gfafd, _gfged.copyObject(_bcfec.Get(_gfafd), _adce, _ceegc, _dcdfb))
		}
		if _gfcaa {
			_dffcf = _dg.MakeNull()
			_dcdfb = false
		}
	case *_dg.PdfIndirectObject:
		_cdfdf := &_dg.PdfIndirectObject{PdfObjectReference: _bcfec.PdfObjectReference}
		_dffcf = _cdfdf
		_adce[_cddfe] = _dffcf
		_cdfdf.PdfObject = _gfged.copyObject(_bcfec.PdfObject, _adce, _ceegc, _dcdfb)
	case *_dg.PdfObjectString:
		_dcfdb := *_bcfec
		_dffcf = &_dcfdb
		_adce[_cddfe] = _dffcf
	case *_dg.PdfObjectName:
		_gegccb := *_bcfec
		_dffcf = &_gegccb
		_adce[_cddfe] = _dffcf
	case *_dg.PdfObjectNull:
		_dffcf = _dg.MakeNull()
		_adce[_cddfe] = _dffcf
	case *_dg.PdfObjectInteger:
		_dgacg := *_bcfec
		_dffcf = &_dgacg
		_adce[_cddfe] = _dffcf
	case *_dg.PdfObjectReference:
		_ggbc := *_bcfec
		_dffcf = &_ggbc
		_adce[_cddfe] = _dffcf
	case *_dg.PdfObjectFloat:
		_dfbgg := *_bcfec
		_dffcf = &_dfbgg
		_adce[_cddfe] = _dffcf
	case *_dg.PdfObjectBool:
		_dggbf := *_bcfec
		_dffcf = &_dggbf
		_adce[_cddfe] = _dffcf
	case *pdfSignDictionary:
		_gcefe := &pdfSignDictionary{PdfObjectDictionary: _dg.MakeDict(), _cgdea: _bcfec._cgdea, _cdfca: _bcfec._cdfca}
		_dffcf = _gcefe
		_adce[_cddfe] = _dffcf
		for _, _ggcec := range _bcfec.Keys() {
			_gcefe.Set(_ggcec, _gfged.copyObject(_bcfec.Get(_ggcec), _adce, _ceegc, _dcdfb))
		}
	default:
		_ag.Log.Info("\u0054\u004f\u0044\u004f\u0028\u0061\u0035\u0069\u0029\u003a\u0020\u0069\u006dp\u006c\u0065\u006d\u0065\u006e\u0074 \u0063\u006f\u0070\u0079\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0066\u006fr\u0020\u0025\u002b\u0076", _cddfe)
	}
	if _dbcf && _dcdfb {
		_ceegc[_cddfe] = struct{}{}
	}
	return _dffcf
}

// NewPdfActionLaunch returns a new "launch" action.
func NewPdfActionLaunch() *PdfActionLaunch {
	_fdg := NewPdfAction()
	_ad := &PdfActionLaunch{}
	_ad.PdfAction = _fdg
	_fdg.SetContext(_ad)
	return _ad
}

// String returns a human readable description of `fontfile`.
func (_abcbg *fontFile) String() string {
	_eeca := "\u005b\u004e\u006f\u006e\u0065\u005d"
	if _abcbg._beaeb != nil {
		_eeca = _abcbg._beaeb.String()
	}
	return _b.Sprintf("\u0046O\u004e\u0054\u0046\u0049\u004c\u0045\u007b\u0025\u0023\u0071\u0020e\u006e\u0063\u006f\u0064\u0065\u0072\u003d\u0025\u0073\u007d", _abcbg._ecgd, _eeca)
}

// GetContentStreamWithEncoder returns the pattern cell's content stream and its encoder
func (_ddebe *PdfTilingPattern) GetContentStreamWithEncoder() ([]byte, _dg.StreamEncoder, error) {
	_bgcba, _eeebf := _ddebe._eacce.(*_dg.PdfObjectStream)
	if !_eeebf {
		_ag.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ddebe._eacce)
		return nil, nil, _dg.ErrTypeError
	}
	_aegef, _ecge := _dg.DecodeStream(_bgcba)
	if _ecge != nil {
		_ag.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020s\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0065\u0072\u0072\u003a\u0020\u0025\u0076", _ecge)
		return nil, nil, _ecge
	}
	_begee, _ecge := _dg.NewEncoderFromStream(_bgcba)
	if _ecge != nil {
		_ag.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020f\u0069\u006e\u0064\u0069\u006e\u0067 \u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _ecge)
		return nil, nil, _ecge
	}
	return _aegef, _begee, nil
}

// ToPdfObject implements interface PdfModel.
func (_fbcab *PdfSignatureReference) ToPdfObject() _dg.PdfObject {
	_eaafc := _dg.MakeDict()
	_eaafc.SetIfNotNil("\u0054\u0079\u0070\u0065", _fbcab.Type)
	_eaafc.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064", _fbcab.TransformMethod)
	_eaafc.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073", _fbcab.TransformParams)
	_eaafc.SetIfNotNil("\u0044\u0061\u0074\u0061", _fbcab.Data)
	_eaafc.SetIfNotNil("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064", _fbcab.DigestMethod)
	return _eaafc
}
func (_fffad *pdfFontSimple) updateStandard14Font() {
	_afca, _eagg := _fffad.Encoder().(_bd.SimpleEncoder)
	if !_eagg {
		_ag.Log.Error("\u0057\u0072\u006f\u006e\u0067\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074y\u0070e\u003a\u0020\u0025\u0054\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u002e", _fffad.Encoder(), _fffad)
		return
	}
	_cgcba := _afca.Charcodes()
	_fffad._cfagf = make(map[_bd.CharCode]float64, len(_cgcba))
	for _, _gcda := range _cgcba {
		_fgec, _ := _afca.CharcodeToRune(_gcda)
		_bffa, _ := _fffad._bfdee.Read(_fgec)
		_fffad._cfagf[_gcda] = _bffa.Wx
	}
}

// GetParamsDict returns *core.PdfObjectDictionary with a set of basic image parameters.
func (_gfdg *Image) GetParamsDict() *_dg.PdfObjectDictionary {
	_dadgb := _dg.MakeDict()
	_dadgb.Set("\u0057\u0069\u0064t\u0068", _dg.MakeInteger(_gfdg.Width))
	_dadgb.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _dg.MakeInteger(_gfdg.Height))
	_dadgb.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _dg.MakeInteger(int64(_gfdg.ColorComponents)))
	_dadgb.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _dg.MakeInteger(_gfdg.BitsPerComponent))
	return _dadgb
}
func (_fcab *PdfColorspaceSpecialPattern) String() string {
	return "\u0050a\u0074\u0074\u0065\u0072\u006e"
}

// NewPdfAction returns an initialized generic PDF action model.
func NewPdfAction() *PdfAction {
	_cc := &PdfAction{}
	_cc._cbd = _dg.MakeIndirectObject(_dg.MakeDict())
	return _cc
}

// NewPdfOutputIntentFromPdfObject creates a new PdfOutputIntent from the input core.PdfObject.
func NewPdfOutputIntentFromPdfObject(object _dg.PdfObject) (*PdfOutputIntent, error) {
	_ffadff := &PdfOutputIntent{}
	if _fddaa := _ffadff.ParsePdfObject(object); _fddaa != nil {
		return nil, _fddaa
	}
	return _ffadff, nil
}

// PdfColorspaceSpecialPattern is a Pattern colorspace.
// Can be defined either as /Pattern or with an underlying colorspace [/Pattern cs].
type PdfColorspaceSpecialPattern struct {
	UnderlyingCS PdfColorspace
	_cbff        *_dg.PdfIndirectObject
}

// IsTerminal returns true for terminal fields, false otherwise.
// Terminal fields are fields whose descendants are only widget annotations.
func (_cecf *PdfField) IsTerminal() bool { return len(_cecf.Kids) == 0 }

// PdfReader represents a PDF file reader. It is a frontend to the lower level parsing mechanism and provides
// a higher level access to work with PDF structure and information, such as the page structure etc.
type PdfReader struct {
	_baad    *_dg.PdfParser
	_dfdg    _dg.PdfObject
	_eeeef   *_dg.PdfIndirectObject
	_eefc    *_dg.PdfObjectDictionary
	_daddd   []*_dg.PdfIndirectObject
	PageList []*PdfPage
	_eeeg    int
	_gccfb   *_dg.PdfObjectDictionary
	_fcfc    *PdfOutlineTreeNode
	AcroForm *PdfAcroForm
	DSS      *DSS
	Rotate   *int64
	_fbcaa   *Permissions
	_ggafe   map[*PdfReader]*PdfReader
	_aggag   []*PdfReader
	_cadfa   *modelManager
	_dadcef  bool
	_addfg   map[_dg.PdfObject]struct{}
	_efcfa   _cf.ReadSeeker
	_aafgg   string
	_cfbbb   bool
	_gcecfe  *ReaderOpts
	_eadef   bool
}

// GetRuneMetrics returns the char metrics for a rune.
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
// well give them 0 width. There is no need for the bool return.
func (_gadf *PdfFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_eefg := _gadf.actualFont()
	if _eefg == nil {
		_ag.Log.Debug("ER\u0052\u004fR\u003a\u0020\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u004e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020f\u006fr\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079p\u0065=\u0025\u0023T", _gadf._cadf)
		return _bbg.CharMetrics{}, false
	}
	if _cdfdb, _afbdg := _eefg.GetRuneMetrics(r); _afbdg {
		return _cdfdb, true
	}
	if _cfdf, _aaaaee := _gadf.GetFontDescriptor(); _aaaaee == nil && _cfdf != nil {
		return _bbg.CharMetrics{Wx: _cfdf._adggg}, true
	}
	_ag.Log.Debug("\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061r\u004d\u0065\u0074\u0072i\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gadf)
	return _bbg.CharMetrics{}, false
}

// AnnotFilterFunc represents a PDF annotation filtering function. If the function
// returns true, the annotation is kept, otherwise it is discarded.
type AnnotFilterFunc func(*PdfAnnotation) bool

// PdfAnnotationRedact represents Redact annotations.
// (Section 12.5.6.23).
type PdfAnnotationRedact struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints  _dg.PdfObject
	IC          _dg.PdfObject
	RO          _dg.PdfObject
	OverlayText _dg.PdfObject
	Repeat      _dg.PdfObject
	DA          _dg.PdfObject
	Q           _dg.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_agffg *PdfTransformParamsDocMDP) ToPdfObject() _dg.PdfObject {
	_feccg := _dg.MakeDict()
	_feccg.SetIfNotNil("\u0054\u0079\u0070\u0065", _agffg.Type)
	_feccg.SetIfNotNil("\u0056", _agffg.V)
	_feccg.SetIfNotNil("\u0050", _agffg.P)
	return _feccg
}

// NewPdfActionGoTo returns a new "go to" action.
func NewPdfActionGoTo() *PdfActionGoTo {
	_ebc := NewPdfAction()
	_ce := &PdfActionGoTo{}
	_ce.PdfAction = _ebc
	_ebc.SetContext(_ce)
	return _ce
}

// GetNumComponents returns the number of color components (1 for Separation).
func (_egaa *PdfColorspaceSpecialSeparation) GetNumComponents() int { return 1 }

var _eabb = map[string]struct{}{"\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074\u0045\u006e\u0063o\u0064\u0069\u006e\u0067": {}, "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}}

// ToJBIG2Image converts current image to the core.JBIG2Image.
func (_fdad *Image) ToJBIG2Image() (*_dg.JBIG2Image, error) {
	_gfbbg, _dbdgc := _fdad.ToGoImage()
	if _dbdgc != nil {
		return nil, _dbdgc
	}
	return _dg.GoImageToJBIG2(_gfbbg, _dg.JB2ImageAutoThreshold)
}

// GetAlphabet returns a map of the runes in `text` and their frequencies.
func GetAlphabet(text string) map[rune]int {
	_fgbe := map[rune]int{}
	for _, _efcg := range text {
		_fgbe[_efcg]++
	}
	return _fgbe
}

// ToPdfOutlineItem returns a low level PdfOutlineItem object,
// based on the current instance.
func (_bbadd *OutlineItem) ToPdfOutlineItem() (*PdfOutlineItem, int64) {
	_abgfg := NewPdfOutlineItem()
	_abgfg.Title = _dg.MakeEncodedString(_bbadd.Title, true)
	_abgfg.Dest = _bbadd.Dest.ToPdfObject()
	var _ffgefc []*PdfOutlineItem
	var _cfcga int64
	var _dgbaf *PdfOutlineItem
	for _, _ffafe := range _bbadd.Entries {
		_egaef, _bbgb := _ffafe.ToPdfOutlineItem()
		_egaef.Parent = &_abgfg.PdfOutlineTreeNode
		if _dgbaf != nil {
			_dgbaf.Next = &_egaef.PdfOutlineTreeNode
			_egaef.Prev = &_dgbaf.PdfOutlineTreeNode
		}
		_ffgefc = append(_ffgefc, _egaef)
		_cfcga += _bbgb
		_dgbaf = _egaef
	}
	_ebdda := len(_ffgefc)
	_cfcga += int64(_ebdda)
	if _ebdda > 0 {
		_abgfg.First = &_ffgefc[0].PdfOutlineTreeNode
		_abgfg.Last = &_ffgefc[_ebdda-1].PdfOutlineTreeNode
		_abgfg.Count = &_cfcga
	}
	return _abgfg, _cfcga
}

// NewPdfShadingType3 creates an empty shading type 3 dictionary.
func NewPdfShadingType3() *PdfShadingType3 {
	_egcd := &PdfShadingType3{}
	_egcd.PdfShading = &PdfShading{}
	_egcd.PdfShading._bcfbg = _dg.MakeIndirectObject(_dg.MakeDict())
	_egcd.PdfShading._eeddb = _egcd
	return _egcd
}

// GetPageDict converts the Page to a PDF object dictionary.
func (_ceaa *PdfPage) GetPageDict() *_dg.PdfObjectDictionary {
	_dgde := _ceaa._bfdge
	_dgde.Clear()
	_dgde.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0050\u0061\u0067\u0065"))
	_dgde.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _ceaa.Parent)
	if _ceaa.LastModified != nil {
		_dgde.Set("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _ceaa.LastModified.ToPdfObject())
	}
	if _ceaa.Resources != nil {
		_dgde.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _ceaa.Resources.ToPdfObject())
	}
	if _ceaa.CropBox != nil {
		_dgde.Set("\u0043r\u006f\u0070\u0042\u006f\u0078", _ceaa.CropBox.ToPdfObject())
	}
	if _ceaa.MediaBox != nil {
		_dgde.Set("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", _ceaa.MediaBox.ToPdfObject())
	}
	if _ceaa.BleedBox != nil {
		_dgde.Set("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078", _ceaa.BleedBox.ToPdfObject())
	}
	if _ceaa.TrimBox != nil {
		_dgde.Set("\u0054r\u0069\u006d\u0042\u006f\u0078", _ceaa.TrimBox.ToPdfObject())
	}
	if _ceaa.ArtBox != nil {
		_dgde.Set("\u0041\u0072\u0074\u0042\u006f\u0078", _ceaa.ArtBox.ToPdfObject())
	}
	_dgde.SetIfNotNil("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f", _ceaa.BoxColorInfo)
	_dgde.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _ceaa.Contents)
	if _ceaa.Rotate != nil {
		_dgde.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _dg.MakeInteger(*_ceaa.Rotate))
	}
	_dgde.SetIfNotNil("\u0047\u0072\u006fu\u0070", _ceaa.Group)
	_dgde.SetIfNotNil("\u0054\u0068\u0075m\u0062", _ceaa.Thumb)
	_dgde.SetIfNotNil("\u0042", _ceaa.B)
	_dgde.SetIfNotNil("\u0044\u0075\u0072", _ceaa.Dur)
	_dgde.SetIfNotNil("\u0054\u0072\u0061n\u0073", _ceaa.Trans)
	_dgde.SetIfNotNil("\u0041\u0041", _ceaa.AA)
	_dgde.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _ceaa.Metadata)
	_dgde.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _ceaa.PieceInfo)
	_dgde.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _ceaa.StructParents)
	_dgde.SetIfNotNil("\u0049\u0044", _ceaa.ID)
	_dgde.SetIfNotNil("\u0050\u005a", _ceaa.PZ)
	_dgde.SetIfNotNil("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f", _ceaa.SeparationInfo)
	_dgde.SetIfNotNil("\u0054\u0061\u0062\u0073", _ceaa.Tabs)
	_dgde.SetIfNotNil("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064", _ceaa.TemplateInstantiated)
	_dgde.SetIfNotNil("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s", _ceaa.PresSteps)
	_dgde.SetIfNotNil("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074", _ceaa.UserUnit)
	_dgde.SetIfNotNil("\u0056\u0050", _ceaa.VP)
	if _ceaa._cadgg != nil {
		_eeacb := _dg.MakeArray()
		for _, _gbdagg := range _ceaa._cadgg {
			if _eaadb := _gbdagg.GetContext(); _eaadb != nil {
				_eeacb.Append(_eaadb.ToPdfObject())
			} else {
				_eeacb.Append(_gbdagg.ToPdfObject())
			}
		}
		if _eeacb.Len() > 0 {
			_dgde.Set("\u0041\u006e\u006e\u006f\u0074\u0073", _eeacb)
		}
	} else if _ceaa.Annots != nil {
		_dgde.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0073", _ceaa.Annots)
	}
	return _dgde
}

// GetContainingPdfObject returns the container of the outline tree node (indirect object).
func (_ecea *PdfOutlineTreeNode) GetContainingPdfObject() _dg.PdfObject {
	return _ecea.GetContext().GetContainingPdfObject()
}

// NewPdfColorPatternType2 returns an empty color shading pattern type 2 (Axial).
func NewPdfColorPatternType2() *PdfColorPatternType2 { _cdaf := &PdfColorPatternType2{}; return _cdaf }

// IsColored specifies if the pattern is colored.
func (_adaa *PdfTilingPattern) IsColored() bool {
	if _adaa.PaintType != nil && *_adaa.PaintType == 1 {
		return true
	}
	return false
}

// SignatureHandlerDocMDP extends SignatureHandler with the ValidateWithOpts method for checking the DocMDP policy.
type SignatureHandlerDocMDP interface {
	SignatureHandler

	// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser
	// ValidateWithOpts shall contain Validate call
	ValidateWithOpts(_aadcb *PdfSignature, _cccbg Hasher, _gbcebg SignatureHandlerDocMDPParams) (SignatureValidationResult, error)
}
type pdfCIDFontType2 struct {
	fontCommon
	_ddccf *_dg.PdfIndirectObject
	_dcde  _bd.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// Dictionary that defines the character collection of the CIDFont (required).
	// See Table 116.
	CIDSystemInfo *_dg.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW  _dg.PdfObject
	W   _dg.PdfObject
	DW2 _dg.PdfObject
	W2  _dg.PdfObject

	// CIDs to glyph indices mapping (optional).
	CIDToGIDMap _dg.PdfObject
	_ddeb       map[_bd.CharCode]float64
	_bfbc       float64
	_ebgc       map[rune]int
}

func _aeegd(_baefe string) (string, error) {
	var _fggea _bc.Buffer
	_fggea.WriteString(_baefe)
	_bddad := make([]byte, 8+16)
	_gecag := _a.Now().UTC().UnixNano()
	_bad.BigEndian.PutUint64(_bddad, uint64(_gecag))
	_, _aaadd := _ec.Read(_bddad[8:])
	if _aaadd != nil {
		return "", _aaadd
	}
	_fggea.WriteString(_be.EncodeToString(_bddad))
	return _fggea.String(), nil
}

// ImageToGray returns a new grayscale image based on the passed in RGB image.
func (_efgga *PdfColorspaceDeviceRGB) ImageToGray(img Image) (Image, error) {
	if img.ColorComponents != 3 {
		return img, _bf.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u0020\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042")
	}
	_aadd, _abcc := _fc.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._dgeb, img._gfbb)
	if _abcc != nil {
		return img, _abcc
	}
	_bbea, _abcc := _fc.GrayConverter.Convert(_aadd)
	if _abcc != nil {
		return img, _abcc
	}
	return _edcf(_bbea.Base()), nil
}

// XObjectType represents the type of an XObject.
type XObjectType int

// GetAllContentStreams gets all the content streams for a page as one string.
func (_efdac *PdfPage) GetAllContentStreams() (string, error) {
	_gdddg, _dacba := _efdac.GetContentStreams()
	if _dacba != nil {
		return "", _dacba
	}
	return _ga.Join(_gdddg, "\u0020"), nil
}

// SetPdfSubject sets the Subject attribute of the output PDF.
func SetPdfSubject(subject string) { _fgefgf.Lock(); defer _fgefgf.Unlock(); _adgf = subject }

// PdfColorPatternType2 represents a color shading pattern type 2 (Axial).
type PdfColorPatternType2 struct {
	Color       PdfColor
	PatternName _dg.PdfObjectName
}

func (_deacf *PdfWriter) writeXRefStreams(_abeed int, _baecfb int64) error {
	_bafg := _abeed + 1
	_deacf._fffge[_bafg] = crossReference{Type: 1, ObjectNumber: _bafg, Offset: _baecfb}
	_bdacdc := _bc.NewBuffer(nil)
	_fcaf := _dg.MakeArray()
	for _febea := 0; _febea <= _abeed; {
		for ; _febea <= _abeed; _febea++ {
			_cdcdd, _bgdbggd := _deacf._fffge[_febea]
			if _bgdbggd && (!_deacf._bbac || _deacf._bbac && (_cdcdd.Type == 1 && _cdcdd.Offset >= _deacf._fafab || _cdcdd.Type == 0)) {
				break
			}
		}
		var _dbgg int
		for _dbgg = _febea + 1; _dbgg <= _abeed; _dbgg++ {
			_gbcc, _dcebb := _deacf._fffge[_dbgg]
			if _dcebb && (!_deacf._bbac || _deacf._bbac && (_gbcc.Type == 1 && _gbcc.Offset > _deacf._fafab)) {
				continue
			}
			break
		}
		_fcaf.Append(_dg.MakeInteger(int64(_febea)), _dg.MakeInteger(int64(_dbgg-_febea)))
		for _gcca := _febea; _gcca < _dbgg; _gcca++ {
			_edfdb := _deacf._fffge[_gcca]
			switch _edfdb.Type {
			case 0:
				_bad.Write(_bdacdc, _bad.BigEndian, byte(0))
				_bad.Write(_bdacdc, _bad.BigEndian, uint32(0))
				_bad.Write(_bdacdc, _bad.BigEndian, uint16(0xFFFF))
			case 1:
				_bad.Write(_bdacdc, _bad.BigEndian, byte(1))
				_bad.Write(_bdacdc, _bad.BigEndian, uint32(_edfdb.Offset))
				_bad.Write(_bdacdc, _bad.BigEndian, uint16(_edfdb.Generation))
			case 2:
				_bad.Write(_bdacdc, _bad.BigEndian, byte(2))
				_bad.Write(_bdacdc, _bad.BigEndian, uint32(_edfdb.ObjectNumber))
				_bad.Write(_bdacdc, _bad.BigEndian, uint16(_edfdb.Index))
			}
		}
		_febea = _dbgg + 1
	}
	_ececc, _bdgba := _dg.MakeStream(_bdacdc.Bytes(), _dg.NewFlateEncoder())
	if _bdgba != nil {
		return _bdgba
	}
	_ececc.ObjectNumber = int64(_bafg)
	_ececc.PdfObjectDictionary.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0058\u0052\u0065\u0066"))
	_ececc.PdfObjectDictionary.Set("\u0057", _dg.MakeArray(_dg.MakeInteger(1), _dg.MakeInteger(4), _dg.MakeInteger(2)))
	_ececc.PdfObjectDictionary.Set("\u0049\u006e\u0064e\u0078", _fcaf)
	_ececc.PdfObjectDictionary.Set("\u0053\u0069\u007a\u0065", _dg.MakeInteger(int64(_bafg)))
	_ececc.PdfObjectDictionary.Set("\u0049\u006e\u0066\u006f", _deacf._efbfa)
	_ececc.PdfObjectDictionary.Set("\u0052\u006f\u006f\u0074", _deacf._fadee)
	if _deacf._bbac && _deacf._eege > 0 {
		_ececc.PdfObjectDictionary.Set("\u0050\u0072\u0065\u0076", _dg.MakeInteger(_deacf._eege))
	}
	if _deacf._fadcg != nil {
		_ececc.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _deacf._acdag)
	}
	if _deacf._agbeg == nil && _deacf._cedaf != "" && _deacf._ceaab != "" {
		_deacf._agbeg = _dg.MakeArray(_dg.MakeHexString(_deacf._cedaf), _dg.MakeHexString(_deacf._ceaab))
	}
	if _deacf._agbeg != nil {
		_ag.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _deacf._agbeg)
		_ececc.Set("\u0049\u0044", _deacf._agbeg)
	}
	_deacf.writeObject(int(_ececc.ObjectNumber), _ececc)
	return nil
}

// PdfModel is a higher level PDF construct which can be collapsed into a PdfObject.
// Each PdfModel has an underlying PdfObject and vice versa (one-to-one).
// Under normal circumstances there should only be one copy of each.
// Copies can be made, but care must be taken to do it properly.
type PdfModel interface {
	ToPdfObject() _dg.PdfObject
	GetContainingPdfObject() _dg.PdfObject
}

// IsCheckbox returns true if the button field represents a checkbox, false otherwise.
func (_eeec *PdfFieldButton) IsCheckbox() bool { return _eeec.GetType() == ButtonTypeCheckbox }

// Hasher is the interface that wraps the basic Write method.
type Hasher interface {
	Write(_gaedg []byte) (_bdebe int, _bffba error)
}

// SetPdfCreationDate sets the CreationDate attribute of the output PDF.
func SetPdfCreationDate(creationDate _a.Time) {
	_fgefgf.Lock()
	defer _fgefgf.Unlock()
	_ccfea = creationDate
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_cdbe *PdfColorspaceSpecialSeparation) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gbab, _fgac := _dg.GetNumbersAsFloat(objects)
	if _fgac != nil {
		return nil, _fgac
	}
	return _cdbe.ColorFromFloats(_gbab)
}

// GetNumComponents returns the number of color components (3 for Lab).
func (_beebf *PdfColorLab) GetNumComponents() int { return 3 }
func (_afdf *PdfReader) newPdfActionLaunchFromDict(_gge *_dg.PdfObjectDictionary) (*PdfActionLaunch, error) {
	_aebc, _aba := _bccf(_gge.Get("\u0046"))
	if _aba != nil {
		return nil, _aba
	}
	return &PdfActionLaunch{Win: _gge.Get("\u0057\u0069\u006e"), Mac: _gge.Get("\u004d\u0061\u0063"), Unix: _gge.Get("\u0055\u006e\u0069\u0078"), NewWindow: _gge.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _aebc}, nil
}

// PdfColorspaceDeviceGray represents a grayscale colorspace.
type PdfColorspaceDeviceGray struct{}

// GetObjectNums returns the object numbers of the PDF objects in the file
// Numbered objects are either indirect objects or stream objects.
// e.g. objNums := pdfReader.GetObjectNums()
// The underlying objects can then be accessed with
// pdfReader.GetIndirectObjectByNumber(objNums[0]) for the first available object.
func (_ffdgc *PdfReader) GetObjectNums() []int { return _ffdgc._baad.GetObjectNums() }

// PdfShadingType3 is a Radial shading.
type PdfShadingType3 struct {
	*PdfShading
	Coords   *_dg.PdfObjectArray
	Domain   *_dg.PdfObjectArray
	Function []PdfFunction
	Extend   *_dg.PdfObjectArray
}

// ToPdfObject implements interface PdfModel.
func (_begg *PdfActionGoToE) ToPdfObject() _dg.PdfObject {
	_begg.PdfAction.ToPdfObject()
	_fee := _begg._cbd
	_afc := _fee.PdfObject.(*_dg.PdfObjectDictionary)
	_afc.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeGoToE)))
	if _begg.F != nil {
		_afc.Set("\u0046", _begg.F.ToPdfObject())
	}
	_afc.SetIfNotNil("\u0044", _begg.D)
	_afc.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _begg.NewWindow)
	_afc.SetIfNotNil("\u0054", _begg.T)
	return _fee
}

// NewXObjectFormFromStream builds the Form XObject from a stream object.
// TODO: Should this be exposed? Consider different access points.
func NewXObjectFormFromStream(stream *_dg.PdfObjectStream) (*XObjectForm, error) {
	_gegda := &XObjectForm{}
	_gegda._ebaeb = stream
	_edddf := *(stream.PdfObjectDictionary)
	_accgb, _fbbeg := _dg.NewEncoderFromStream(stream)
	if _fbbeg != nil {
		return nil, _fbbeg
	}
	_gegda.Filter = _accgb
	if _dgfda := _edddf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _dgfda != nil {
		_caaad, _effe := _dgfda.(*_dg.PdfObjectName)
		if !_effe {
			return nil, _bf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if *_caaad != "\u0046\u006f\u0072\u006d" {
			_ag.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
			return nil, _bf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
		}
	}
	if _ecbda := _edddf.Get("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065"); _ecbda != nil {
		_gegda.FormType = _ecbda
	}
	if _cafef := _edddf.Get("\u0042\u0042\u006f\u0078"); _cafef != nil {
		_gegda.BBox = _cafef
	}
	if _bbbeca := _edddf.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _bbbeca != nil {
		_gegda.Matrix = _bbbeca
	}
	if _cggc := _edddf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _cggc != nil {
		_cggc = _dg.TraceToDirectObject(_cggc)
		_dfdab, _gaebg := _cggc.(*_dg.PdfObjectDictionary)
		if !_gaebg {
			_ag.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0058\u004f\u0062j\u0065c\u0074\u0020\u0046\u006f\u0072\u006d\u0020\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u002c\u0020\u0070\u006f\u0069\u006e\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _dg.ErrTypeError
		}
		_bede, _egfda := NewPdfPageResourcesFromDict(_dfdab)
		if _egfda != nil {
			_ag.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0020\u0072\u0065\u0073\u006f\u0075rc\u0065\u0073")
			return nil, _egfda
		}
		_gegda.Resources = _bede
		_ag.Log.Trace("\u0046\u006f\u0072\u006d r\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u003a\u0020\u0025\u0023\u0076", _gegda.Resources)
	}
	_gegda.Group = _edddf.Get("\u0047\u0072\u006fu\u0070")
	_gegda.Ref = _edddf.Get("\u0052\u0065\u0066")
	_gegda.MetaData = _edddf.Get("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061")
	_gegda.PieceInfo = _edddf.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o")
	_gegda.LastModified = _edddf.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064")
	_gegda.StructParent = _edddf.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_gegda.StructParents = _edddf.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073")
	_gegda.OPI = _edddf.Get("\u004f\u0050\u0049")
	_gegda.OC = _edddf.Get("\u004f\u0043")
	_gegda.Name = _edddf.Get("\u004e\u0061\u006d\u0065")
	_gegda.Stream = stream.Stream
	return _gegda, nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_bbbbd *PdfShadingType6) ToPdfObject() _dg.PdfObject {
	_bbbbd.PdfShading.ToPdfObject()
	_feafd, _bdff := _bbbbd.getShadingDict()
	if _bdff != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _bbbbd.BitsPerCoordinate != nil {
		_feafd.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _bbbbd.BitsPerCoordinate)
	}
	if _bbbbd.BitsPerComponent != nil {
		_feafd.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _bbbbd.BitsPerComponent)
	}
	if _bbbbd.BitsPerFlag != nil {
		_feafd.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _bbbbd.BitsPerFlag)
	}
	if _bbbbd.Decode != nil {
		_feafd.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _bbbbd.Decode)
	}
	if _bbbbd.Function != nil {
		if len(_bbbbd.Function) == 1 {
			_feafd.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _bbbbd.Function[0].ToPdfObject())
		} else {
			_gefaf := _dg.MakeArray()
			for _, _ddag := range _bbbbd.Function {
				_gefaf.Append(_ddag.ToPdfObject())
			}
			_feafd.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _gefaf)
		}
	}
	return _bbbbd._bcfbg
}
func (_ecgdg *PdfWriter) writeAcroFormFields() error {
	if _ecgdg._geabe == nil {
		return nil
	}
	_ag.Log.Trace("\u0057r\u0069t\u0069\u006e\u0067\u0020\u0061c\u0072\u006f \u0066\u006f\u0072\u006d\u0073")
	_geefg := _ecgdg._geabe.ToPdfObject()
	_ag.Log.Trace("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _geefg)
	_ecgdg._ecdf.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _geefg)
	_dcfad := _ecgdg.addObjects(_geefg)
	if _dcfad != nil {
		return _dcfad
	}
	return nil
}

// ColorFromPdfObjects gets the color from a series of pdf objects (3 for rgb).
func (_ddabd *PdfColorspaceDeviceRGB) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fegb, _fgcf := _dg.GetNumbersAsFloat(objects)
	if _fgcf != nil {
		return nil, _fgcf
	}
	return _ddabd.ColorFromFloats(_fegb)
}

type pdfSignDictionary struct {
	*_dg.PdfObjectDictionary
	_cgdea  *SignatureHandler
	_cdfca  *PdfSignature
	_cagb   int64
	_aaeae  int
	_gdbfe  int
	_egda   int
	_ebaaaa int
}

// UpdateXObjectImageFromImage creates a new XObject Image from an
// Image object `img` and default masks from xobjIn.
// The default masks are overridden if img.hasAlpha
// If `encoder` is nil, uses raw encoding (none).
func UpdateXObjectImageFromImage(xobjIn *XObjectImage, img *Image, cs PdfColorspace, encoder _dg.StreamEncoder) (*XObjectImage, error) {
	if encoder == nil {
		var _fbea error
		encoder, _fbea = img.getSuitableEncoder()
		if _fbea != nil {
			_ag.Log.Debug("F\u0061\u0069l\u0075\u0072\u0065\u0020\u006f\u006e\u0020\u0066\u0069\u006e\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0069\u0074\u0061b\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072,\u0020\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b\u0020\u0074\u006f\u0020R\u0061\u0077\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _fbea)
			encoder = _dg.NewRawEncoder()
		}
	}
	encoder.UpdateParams(img.GetParamsDict())
	_bagda, _caged := encoder.EncodeBytes(img.Data)
	if _caged != nil {
		_ag.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _caged)
		return nil, _caged
	}
	_efcfe := NewXObjectImage()
	_bdbbf := img.Width
	_ebaac := img.Height
	_efcfe.Width = &_bdbbf
	_efcfe.Height = &_ebaac
	_gbfdf := img.BitsPerComponent
	_efcfe.BitsPerComponent = &_gbfdf
	_efcfe.Filter = encoder
	_efcfe.Stream = _bagda
	if cs == nil {
		if img.ColorComponents == 1 {
			_efcfe.ColorSpace = NewPdfColorspaceDeviceGray()
			if img.BitsPerComponent == 16 {
				switch encoder.(type) {
				case *_dg.DCTEncoder:
					_efcfe.ColorSpace = NewPdfColorspaceDeviceRGB()
					_gbfdf = 8
					_efcfe.BitsPerComponent = &_gbfdf
				}
			}
		} else if img.ColorComponents == 3 {
			_efcfe.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			switch encoder.(type) {
			case *_dg.DCTEncoder:
				_efcfe.ColorSpace = NewPdfColorspaceDeviceRGB()
			default:
				_efcfe.ColorSpace = NewPdfColorspaceDeviceCMYK()
			}
		} else {
			return nil, _bf.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_efcfe.ColorSpace = cs
	}
	if len(img._dgeb) != 0 {
		_fgdb := NewXObjectImage()
		_fgdb.Filter = encoder
		_bacac, _bbfcc := encoder.EncodeBytes(img._dgeb)
		if _bbfcc != nil {
			_ag.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _bbfcc)
			return nil, _bbfcc
		}
		_fgdb.Stream = _bacac
		_fgdb.BitsPerComponent = _efcfe.BitsPerComponent
		_fgdb.Width = &img.Width
		_fgdb.Height = &img.Height
		_fgdb.ColorSpace = NewPdfColorspaceDeviceGray()
		_efcfe.SMask = _fgdb.ToPdfObject()
	} else {
		_efcfe.SMask = xobjIn.SMask
		_efcfe.ImageMask = xobjIn.ImageMask
		if _efcfe.ColorSpace.GetNumComponents() == 1 {
			_egcfg(_efcfe)
		}
	}
	return _efcfe, nil
}

// AddCRLs adds CRLs to DSS.
func (_cdbb *DSS) AddCRLs(crls [][]byte) ([]*_dg.PdfObjectStream, error) {
	return _cdbb.add(&_cdbb.CRLs, _cdbb._dcdf, crls)
}
func (_ffa *PdfAnnotation) String() string {
	_efd := ""
	_gcfa, _dab := _ffa.ToPdfObject().(*_dg.PdfIndirectObject)
	if _dab {
		_efd = _b.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _ffa._egcg, _gcfa.PdfObject.String())
	}
	return _efd
}
func (_gegbf *PdfReader) newPdfFieldSignatureFromDict(_dfeg *_dg.PdfObjectDictionary) (*PdfFieldSignature, error) {
	_fggbd := &PdfFieldSignature{}
	_begc, _ggbd := _dg.GetIndirect(_dfeg.Get("\u0056"))
	if _ggbd {
		var _ececg error
		_fggbd.V, _ececg = _gegbf.newPdfSignatureFromIndirect(_begc)
		if _ececg != nil {
			return nil, _ececg
		}
	}
	_fggbd.Lock, _ = _dg.GetIndirect(_dfeg.Get("\u004c\u006f\u0063\u006b"))
	_fggbd.SV, _ = _dg.GetIndirect(_dfeg.Get("\u0053\u0056"))
	return _fggbd, nil
}

// HasExtGState checks if ExtGState name is available.
func (_defcdb *PdfPage) HasExtGState(name _dg.PdfObjectName) bool {
	if _defcdb.Resources == nil {
		return false
	}
	if _defcdb.Resources.ExtGState == nil {
		return false
	}
	_cbbec, _fcegg := _dg.TraceToDirectObject(_defcdb.Resources.ExtGState).(*_dg.PdfObjectDictionary)
	if !_fcegg {
		_ag.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _dg.TraceToDirectObject(_defcdb.Resources.ExtGState))
		return false
	}
	_ecgf := _cbbec.Get(name)
	_eaade := _ecgf != nil
	return _eaade
}
func (_adcf *Image) samplesAddPadding(_fgaab []uint32) []uint32 {
	_cegde := _fc.BytesPerLine(int(_adcf.Width), int(_adcf.BitsPerComponent), _adcf.ColorComponents) * (8 / int(_adcf.BitsPerComponent))
	_bfaca := _cegde * int(_adcf.Height)
	if len(_fgaab) == _bfaca {
		return _fgaab
	}
	_bgaa := make([]uint32, _bfaca)
	_eegac := int(_adcf.Width) * _adcf.ColorComponents
	for _defb := 0; _defb < int(_adcf.Height); _defb++ {
		_cgba := _defb * int(_adcf.Width)
		_cfdfdd := _defb * _cegde
		for _egaba := 0; _egaba < _eegac; _egaba++ {
			_bgaa[_cfdfdd+_egaba] = _fgaab[_cgba+_egaba]
		}
	}
	return _bgaa
}
func _abcgf(_aegabe *[]*PdfField, _adfe FieldFilterFunc, _gceab bool) []*PdfField {
	if _aegabe == nil {
		return nil
	}
	_egaf := *_aegabe
	if len(*_aegabe) == 0 {
		return nil
	}
	_faedg := _egaf[:0]
	if _adfe == nil {
		_adfe = func(*PdfField) bool { return true }
	}
	var _bbafa []*PdfField
	for _, _fecad := range _egaf {
		_ccbdg := _adfe(_fecad)
		if _ccbdg {
			_bbafa = append(_bbafa, _fecad)
			if len(_fecad.Kids) > 0 {
				_bbafa = append(_bbafa, _abcgf(&_fecad.Kids, _adfe, _gceab)...)
			}
		}
		if !_gceab || !_ccbdg || len(_fecad.Kids) > 0 {
			_faedg = append(_faedg, _fecad)
		}
	}
	*_aegabe = _faedg
	return _bbafa
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_bdbgc *PdfShadingType1) ToPdfObject() _dg.PdfObject {
	_bdbgc.PdfShading.ToPdfObject()
	_gfcfc, _fdfebg := _bdbgc.getShadingDict()
	if _fdfebg != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _bdbgc.Domain != nil {
		_gfcfc.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _bdbgc.Domain)
	}
	if _bdbgc.Matrix != nil {
		_gfcfc.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _bdbgc.Matrix)
	}
	if _bdbgc.Function != nil {
		if len(_bdbgc.Function) == 1 {
			_gfcfc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _bdbgc.Function[0].ToPdfObject())
		} else {
			_aagce := _dg.MakeArray()
			for _, _aecc := range _bdbgc.Function {
				_aagce.Append(_aecc.ToPdfObject())
			}
			_gfcfc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _aagce)
		}
	}
	return _bdbgc._bcfbg
}

// PdfActionHide represents a hide action.
type PdfActionHide struct {
	*PdfAction
	T _dg.PdfObject
	H _dg.PdfObject
}

// ToPdfObject returns a PDF object representation of the outline.
func (_efdda *Outline) ToPdfObject() _dg.PdfObject { return _efdda.ToPdfOutline().ToPdfObject() }

// PdfFunctionType3 defines stitching of the subdomains of several 1-input functions to produce
// a single new 1-input function.
type PdfFunctionType3 struct {
	Domain    []float64
	Range     []float64
	Functions []PdfFunction
	Bounds    []float64
	Encode    []float64
	_defdg    *_dg.PdfIndirectObject
}

// Evaluate runs the function. Input is [x1 x2 x3].
func (_beccd *PdfFunctionType4) Evaluate(xVec []float64) ([]float64, error) {
	if _beccd._agbgd == nil {
		_beccd._agbgd = _gfd.NewPSExecutor(_beccd.Program)
	}
	var _fdag []_gfd.PSObject
	for _, _gcffa := range xVec {
		_fdag = append(_fdag, _gfd.MakeReal(_gcffa))
	}
	_eaaa, _cdfbc := _beccd._agbgd.Execute(_fdag)
	if _cdfbc != nil {
		return nil, _cdfbc
	}
	_gggfe, _cdfbc := _gfd.PSObjectArrayToFloat64Array(_eaaa)
	if _cdfbc != nil {
		return nil, _cdfbc
	}
	return _gggfe, nil
}

// GetContext returns the context of the outline tree node, which is either a
// *PdfOutline or a *PdfOutlineItem. The method returns nil for uninitialized
// tree nodes.
func (_ccab *PdfOutlineTreeNode) GetContext() PdfModel {
	if _baddd, _cgef := _ccab._baddf.(*PdfOutline); _cgef {
		return _baddd
	}
	if _acfb, _bgegb := _ccab._baddf.(*PdfOutlineItem); _bgegb {
		return _acfb
	}
	_ag.Log.Debug("\u0045\u0052RO\u0052\u0020\u0049n\u0076\u0061\u006c\u0069d o\u0075tl\u0069\u006e\u0065\u0020\u0074\u0072\u0065e \u006e\u006f\u0064\u0065\u0020\u0069\u0074e\u006d")
	return nil
}
func _ececd(_bgff []byte) bool {
	if len(_bgff) < 4 {
		return true
	}
	for _cgada := range _bgff[:4] {
		_dfbf := rune(_cgada)
		if !_af.Is(_af.ASCII_Hex_Digit, _dfbf) && !_af.IsSpace(_dfbf) {
			return true
		}
	}
	return false
}

// PdfColorspaceCalRGB stores A, B, C components
type PdfColorspaceCalRGB struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      []float64
	Matrix     []float64
	_dfab      *_dg.PdfObjectDictionary
	_ccae      *_dg.PdfIndirectObject
}

func _eaac(_feed *_dg.PdfObjectDictionary) (*PdfShadingType4, error) {
	_fdeac := PdfShadingType4{}
	_aeede := _feed.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _aeede == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_cedbe, _cegbd := _aeede.(*_dg.PdfObjectInteger)
	if !_cegbd {
		_ag.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _aeede)
		return nil, _dg.ErrTypeError
	}
	_fdeac.BitsPerCoordinate = _cedbe
	_aeede = _feed.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _aeede == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_cedbe, _cegbd = _aeede.(*_dg.PdfObjectInteger)
	if !_cegbd {
		_ag.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _aeede)
		return nil, _dg.ErrTypeError
	}
	_fdeac.BitsPerComponent = _cedbe
	_aeede = _feed.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _aeede == nil {
		_ag.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_cedbe, _cegbd = _aeede.(*_dg.PdfObjectInteger)
	if !_cegbd {
		_ag.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _aeede)
		return nil, _dg.ErrTypeError
	}
	_fdeac.BitsPerComponent = _cedbe
	_aeede = _feed.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _aeede == nil {
		_ag.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_cdca, _cegbd := _aeede.(*_dg.PdfObjectArray)
	if !_cegbd {
		_ag.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _aeede)
		return nil, _dg.ErrTypeError
	}
	_fdeac.Decode = _cdca
	_aeede = _feed.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _aeede == nil {
		_ag.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_fdeac.Function = []PdfFunction{}
	if _bdgcc, _dedee := _aeede.(*_dg.PdfObjectArray); _dedee {
		for _, _bedac := range _bdgcc.Elements() {
			_ecfe, _dgdde := _agec(_bedac)
			if _dgdde != nil {
				_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dgdde)
				return nil, _dgdde
			}
			_fdeac.Function = append(_fdeac.Function, _ecfe)
		}
	} else {
		_efefe, _abaad := _agec(_aeede)
		if _abaad != nil {
			_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _abaad)
			return nil, _abaad
		}
		_fdeac.Function = append(_fdeac.Function, _efefe)
	}
	return &_fdeac, nil
}

// NewPdfAnnotationInk returns a new ink annotation.
func NewPdfAnnotationInk() *PdfAnnotationInk {
	_ebf := NewPdfAnnotation()
	_fgg := &PdfAnnotationInk{}
	_fgg.PdfAnnotation = _ebf
	_fgg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ebf.SetContext(_fgg)
	return _fgg
}
func (_bfddd *pdfFontType3) baseFields() *fontCommon { return &_bfddd.fontCommon }
func (_aeg *PdfReader) newPdfActionSubmitFormFromDict(_acc *_dg.PdfObjectDictionary) (*PdfActionSubmitForm, error) {
	_dec, _agcc := _bccf(_acc.Get("\u0046"))
	if _agcc != nil {
		return nil, _agcc
	}
	return &PdfActionSubmitForm{F: _dec, Fields: _acc.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _acc.Get("\u0046\u006c\u0061g\u0073")}, nil
}
func (_bdgb *PdfColorspaceDeviceRGB) String() string {
	return "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
}
func (_dafaf *XObjectImage) getParamsDict() *_dg.PdfObjectDictionary {
	_agabg := _dg.MakeDict()
	_agabg.Set("\u0057\u0069\u0064t\u0068", _dg.MakeInteger(*_dafaf.Width))
	_agabg.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _dg.MakeInteger(*_dafaf.Height))
	_agabg.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _dg.MakeInteger(int64(_dafaf.ColorSpace.GetNumComponents())))
	_agabg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _dg.MakeInteger(*_dafaf.BitsPerComponent))
	return _agabg
}

// DetermineColorspaceNameFromPdfObject determines PDF colorspace from a PdfObject.  Returns the colorspace name and
// an error on failure. If the colorspace was not found, will return an empty string.
func DetermineColorspaceNameFromPdfObject(obj _dg.PdfObject) (_dg.PdfObjectName, error) {
	var _bfcg *_dg.PdfObjectName
	var _abeaf *_dg.PdfObjectArray
	if _gbad, _cebe := obj.(*_dg.PdfIndirectObject); _cebe {
		if _cacc, _dafa := _gbad.PdfObject.(*_dg.PdfObjectArray); _dafa {
			_abeaf = _cacc
		} else if _fabc, _aegba := _gbad.PdfObject.(*_dg.PdfObjectName); _aegba {
			_bfcg = _fabc
		}
	} else if _eacb, _eaed := obj.(*_dg.PdfObjectArray); _eaed {
		_abeaf = _eacb
	} else if _afcfg, _aegd := obj.(*_dg.PdfObjectName); _aegd {
		_bfcg = _afcfg
	}
	if _bfcg != nil {
		switch *_bfcg {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_bfcg, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_bfcg, nil
		}
	}
	if _abeaf != nil && _abeaf.Len() > 0 {
		if _aded, _cbfe := _abeaf.Get(0).(*_dg.PdfObjectName); _cbfe {
			switch *_aded {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _abeaf.Len() == 1 {
					return *_aded, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_aded, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_aded, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_aded, nil
			}
		}
	}
	return "", nil
}

// PdfShading represents a shading dictionary. There are 7 types of shading,
// indicatedby the shading type variable:
// 1: Function-based shading.
// 2: Axial shading.
// 3: Radial shading.
// 4: Free-form Gouraud-shaded triangle mesh.
// 5: Lattice-form Gouraud-shaded triangle mesh.
// 6: Coons patch mesh.
// 7: Tensor-product patch mesh.
// types 4-7 are contained in a stream object, where the dictionary is given by the stream dictionary.
type PdfShading struct {
	ShadingType *_dg.PdfObjectInteger
	ColorSpace  PdfColorspace
	Background  *_dg.PdfObjectArray
	BBox        *PdfRectangle
	AntiAlias   *_dg.PdfObjectBool
	_eeddb      PdfModel
	_bcfbg      _dg.PdfObject
}

// AddExtGState adds a graphics state to the XObject resources.
func (_fcgeg *PdfPage) AddExtGState(name _dg.PdfObjectName, egs *_dg.PdfObjectDictionary) error {
	if _fcgeg.Resources == nil {
		_fcgeg.Resources = NewPdfPageResources()
	}
	if _fcgeg.Resources.ExtGState == nil {
		_fcgeg.Resources.ExtGState = _dg.MakeDict()
	}
	_bcdec, _cgcbc := _dg.TraceToDirectObject(_fcgeg.Resources.ExtGState).(*_dg.PdfObjectDictionary)
	if !_cgcbc {
		_ag.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _dg.TraceToDirectObject(_fcgeg.Resources.ExtGState))
		return _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_bcdec.Set(name, egs)
	return nil
}

type pdfFontType0 struct {
	fontCommon
	_bedc          *_dg.PdfIndirectObject
	_ggec          _bd.TextEncoder
	Encoding       _dg.PdfObject
	DescendantFont *PdfFont
	_bfae          *_ff.CMap
}

// GetContainingPdfObject implements interface PdfModel.
func (_aebbc *PdfSignatureReference) GetContainingPdfObject() _dg.PdfObject { return _aebbc._eccgb }

// PdfAnnotationPolyLine represents PolyLine annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolyLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _dg.PdfObject
	LE       _dg.PdfObject
	BS       _dg.PdfObject
	IC       _dg.PdfObject
	BE       _dg.PdfObject
	IT       _dg.PdfObject
	Measure  _dg.PdfObject
}

func (_cee *PdfReader) newPdfActionGotoRFromDict(_ceeb *_dg.PdfObjectDictionary) (*PdfActionGoToR, error) {
	_eeg, _beeb := _bccf(_ceeb.Get("\u0046"))
	if _beeb != nil {
		return nil, _beeb
	}
	return &PdfActionGoToR{D: _ceeb.Get("\u0044"), NewWindow: _ceeb.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _eeg}, nil
}

// PdfColorspaceICCBased format [/ICCBased stream]
//
// The stream shall contain the ICC profile.
// A conforming reader shall support ICC.1:2004:10 as required by PDF 1.7, which will enable it
// to properly render all embedded ICC profiles regardless of the PDF version
//
// In the current implementation, we rely on the alternative colormap provided.
type PdfColorspaceICCBased struct {
	N         int
	Alternate PdfColorspace

	// If omitted ICC not supported: then use DeviceGray,
	// DeviceRGB or DeviceCMYK for N=1,3,4 respectively.
	Range    []float64
	Metadata *_dg.PdfObjectStream
	Data     []byte
	_fbdff   *_dg.PdfIndirectObject
	_dafg    *_dg.PdfObjectStream
}

func (_cca *PdfReader) newPdfAnnotationWatermarkFromDict(_gac *_dg.PdfObjectDictionary) (*PdfAnnotationWatermark, error) {
	_abbe := PdfAnnotationWatermark{}
	_abbe.FixedPrint = _gac.Get("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074")
	return &_abbe, nil
}

// UpdateObject marks `obj` as updated and to be included in the following revision.
func (_ccag *PdfAppender) UpdateObject(obj _dg.PdfObject) {
	_ccag.replaceObject(obj, obj)
	if _, _edde := _ccag._efa[obj]; !_edde {
		_ccag._dgfd = append(_ccag._dgfd, obj)
		_ccag._efa[obj] = struct{}{}
	}
}

// CheckAccessRights checks access rights and permissions for a specified password.  If either user/owner
// password is specified,  full rights are granted, otherwise the access rights are specified by the
// Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_fegbg *PdfReader) CheckAccessRights(password []byte) (bool, _gbd.Permissions, error) {
	return _fegbg._baad.CheckAccessRights(password)
}

// ToPdfObject returns a stream object.
func (_bbaae *XObjectImage) ToPdfObject() _dg.PdfObject {
	_cddfec := _bbaae._abfb
	_ceea := _cddfec.PdfObjectDictionary
	if _bbaae.Filter != nil {
		_ceea = _bbaae.Filter.MakeStreamDict()
		_cddfec.PdfObjectDictionary = _ceea
	}
	_ceea.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_ceea.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0049\u006d\u0061g\u0065"))
	_ceea.Set("\u0057\u0069\u0064t\u0068", _dg.MakeInteger(*(_bbaae.Width)))
	_ceea.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _dg.MakeInteger(*(_bbaae.Height)))
	if _bbaae.BitsPerComponent != nil {
		_ceea.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _dg.MakeInteger(*(_bbaae.BitsPerComponent)))
	}
	if _bbaae.ColorSpace != nil {
		_ceea.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _bbaae.ColorSpace.ToPdfObject())
	}
	_ceea.SetIfNotNil("\u0049\u006e\u0074\u0065\u006e\u0074", _bbaae.Intent)
	_ceea.SetIfNotNil("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k", _bbaae.ImageMask)
	_ceea.SetIfNotNil("\u004d\u0061\u0073\u006b", _bbaae.Mask)
	_cbbef := _ceea.Get("\u0044\u0065\u0063\u006f\u0064\u0065") != nil
	if _bbaae.Decode == nil && _cbbef {
		_ceea.Remove("\u0044\u0065\u0063\u006f\u0064\u0065")
	} else if _bbaae.Decode != nil {
		_ceea.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _bbaae.Decode)
	}
	_ceea.SetIfNotNil("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _bbaae.Interpolate)
	_ceea.SetIfNotNil("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073", _bbaae.Alternatives)
	_ceea.SetIfNotNil("\u0053\u004d\u0061s\u006b", _bbaae.SMask)
	_ceea.SetIfNotNil("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061", _bbaae.SMaskInData)
	_ceea.SetIfNotNil("\u004d\u0061\u0074t\u0065", _bbaae.Matte)
	_ceea.SetIfNotNil("\u004e\u0061\u006d\u0065", _bbaae.Name)
	_ceea.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _bbaae.StructParent)
	_ceea.SetIfNotNil("\u0049\u0044", _bbaae.ID)
	_ceea.SetIfNotNil("\u004f\u0050\u0049", _bbaae.OPI)
	_ceea.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _bbaae.Metadata)
	_ceea.SetIfNotNil("\u004f\u0043", _bbaae.OC)
	_ceea.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dg.MakeInteger(int64(len(_bbaae.Stream))))
	_cddfec.Stream = _bbaae.Stream
	return _cddfec
}

// ToInteger convert to an integer format.
func (_dcbe *PdfColorDeviceRGB) ToInteger(bits int) [3]uint32 {
	_dag := _cg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_dag * _dcbe.R()), uint32(_dag * _dcbe.G()), uint32(_dag * _dcbe.B())}
}

// NewPdfFontFromTTF loads a TTF font and returns a PdfFont type that can be
// used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTF.
func NewPdfFontFromTTF(r _cf.ReadSeeker) (*PdfFont, error) {
	const _gcba = _bd.CharCode(32)
	const _begca = _bd.CharCode(255)
	_aebec, _ggad := _gf.ReadAll(r)
	if _ggad != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _ggad)
		return nil, _ggad
	}
	_cffae, _ggad := _bbg.TtfParse(_bc.NewReader(_aebec))
	if _ggad != nil {
		_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0054\u0054F\u0020\u0066\u006fn\u0074:\u0020\u0025\u0076", _ggad)
		return nil, _ggad
	}
	_gefbg := &pdfFontSimple{_cfagf: make(map[_bd.CharCode]float64), fontCommon: fontCommon{_bcga: "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065"}}
	_gefbg._bdeed = _bd.NewWinAnsiEncoder()
	_gefbg._ecbf = _cffae.PostScriptName
	_gefbg.FirstChar = _dg.MakeInteger(int64(_gcba))
	_gefbg.LastChar = _dg.MakeInteger(int64(_begca))
	_cbae := 1000.0 / float64(_cffae.UnitsPerEm)
	if len(_cffae.Widths) <= 0 {
		return nil, _bf.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_baaca := _cbae * float64(_cffae.Widths[0])
	_adgdf := make([]float64, 0, _begca-_gcba+1)
	for _acbag := _gcba; _acbag <= _begca; _acbag++ {
		_aabce, _dgad := _gefbg.Encoder().CharcodeToRune(_acbag)
		if !_dgad {
			_ag.Log.Debug("\u0052u\u006e\u0065\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u0028c\u006f\u0064\u0065\u003a\u0020\u0025\u0064)", _acbag)
			_adgdf = append(_adgdf, _baaca)
			continue
		}
		_afec, _egaaa := _cffae.Chars[_aabce]
		if !_egaaa {
			_ag.Log.Debug("R\u0075\u006e\u0065\u0020no\u0074 \u0069\u006e\u0020\u0054\u0054F\u0020\u0043\u0068\u0061\u0072\u0073")
			_adgdf = append(_adgdf, _baaca)
			continue
		}
		_fdddf := _cbae * float64(_cffae.Widths[_afec])
		_adgdf = append(_adgdf, _fdddf)
	}
	_gefbg.Widths = _dg.MakeIndirectObject(_dg.MakeArrayFromFloats(_adgdf))
	if len(_adgdf) < int(_begca-_gcba+1) {
		_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u006f\u0066\u0020\u0077\u0069\u0064\u0074\u0068s,\u0020\u0025\u0064 \u003c \u0025\u0064", len(_adgdf), 255-32+1)
		return nil, _dg.ErrRangeError
	}
	for _egfff := _gcba; _egfff <= _begca; _egfff++ {
		_gefbg._cfagf[_egfff] = _adgdf[_egfff-_gcba]
	}
	_gefbg.Encoding = _dg.MakeName("\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	_cdfac := &PdfFontDescriptor{}
	_cdfac.FontName = _dg.MakeName(_cffae.PostScriptName)
	_cdfac.Ascent = _dg.MakeFloat(_cbae * float64(_cffae.TypoAscender))
	_cdfac.Descent = _dg.MakeFloat(_cbae * float64(_cffae.TypoDescender))
	_cdfac.CapHeight = _dg.MakeFloat(_cbae * float64(_cffae.CapHeight))
	_cdfac.FontBBox = _dg.MakeArrayFromFloats([]float64{_cbae * float64(_cffae.Xmin), _cbae * float64(_cffae.Ymin), _cbae * float64(_cffae.Xmax), _cbae * float64(_cffae.Ymax)})
	_cdfac.ItalicAngle = _dg.MakeFloat(_cffae.ItalicAngle)
	_cdfac.MissingWidth = _dg.MakeFloat(_cbae * float64(_cffae.Widths[0]))
	_eegdb, _ggad := _dg.MakeStream(_aebec, _dg.NewFlateEncoder())
	if _ggad != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _ggad)
		return nil, _ggad
	}
	_eegdb.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _dg.MakeInteger(int64(len(_aebec))))
	_cdfac.FontFile2 = _eegdb
	if _cffae.Bold {
		_cdfac.StemV = _dg.MakeInteger(120)
	} else {
		_cdfac.StemV = _dg.MakeInteger(70)
	}
	_defge := _gbaba
	if _cffae.IsFixedPitch {
		_defge |= _gafad
	}
	if _cffae.ItalicAngle != 0 {
		_defge |= _bfeb
	}
	_cdfac.Flags = _dg.MakeInteger(int64(_defge))
	_gefbg._ccfb = _cdfac
	_cfgf := &PdfFont{_cadf: _gefbg}
	return _cfgf, nil
}

// PdfActionResetForm represents a resetForm action.
type PdfActionResetForm struct {
	*PdfAction
	Fields _dg.PdfObject
	Flags  _dg.PdfObject
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_debbb *PdfShadingPattern) ToPdfObject() _dg.PdfObject {
	_debbb.PdfPattern.ToPdfObject()
	_deecg := _debbb.getDict()
	if _debbb.Shading != nil {
		_deecg.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _debbb.Shading.ToPdfObject())
	}
	if _debbb.Matrix != nil {
		_deecg.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _debbb.Matrix)
	}
	if _debbb.ExtGState != nil {
		_deecg.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _debbb.ExtGState)
	}
	return _debbb._eacce
}

// ToPdfObject converts the font to a PDF representation.
func (_dgfeg *pdfFontType0) ToPdfObject() _dg.PdfObject {
	if _dgfeg._bedc == nil {
		_dgfeg._bedc = &_dg.PdfIndirectObject{}
	}
	_fggcb := _dgfeg.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0030")
	_dgfeg._bedc.PdfObject = _fggcb
	if _dgfeg.Encoding != nil {
		_fggcb.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _dgfeg.Encoding)
	} else if _dgfeg._ggec != nil {
		_fggcb.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _dgfeg._ggec.ToPdfObject())
	}
	if _dgfeg.DescendantFont != nil {
		_fggcb.Set("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073", _dg.MakeArray(_dgfeg.DescendantFont.ToPdfObject()))
	}
	return _dgfeg._bedc
}

// StandardImplementer is an interface that defines specified PDF standards like PDF/A-1A (pdfa.Profile1A)
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardImplementer interface {
	StandardValidator
	StandardApplier

	// StandardName gets the human-readable name of the standard.
	StandardName() string
}

// PdfActionLaunch represents a launch action.
type PdfActionLaunch struct {
	*PdfAction
	F         *PdfFilespec
	Win       _dg.PdfObject
	Mac       _dg.PdfObject
	Unix      _dg.PdfObject
	NewWindow _dg.PdfObject
}

// Sign signs a specific page with a digital signature.
// The signature field parameter must have a valid signature dictionary
// specified by its V field.
func (_ddgg *PdfAppender) Sign(pageNum int, field *PdfFieldSignature) error {
	if field == nil {
		return _bf.New("\u0073\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 n\u0069\u006c")
	}
	_ffad := field.V
	if _ffad == nil {
		return _bf.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061r\u0079 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_cadgf := pageNum - 1
	if _cadgf < 0 || _cadgf > len(_ddgg._ggdd)-1 {
		return _b.Errorf("\u0070\u0061\u0067\u0065\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064", pageNum)
	}
	_cdbg := _ddgg.Reader.PageList[_cadgf]
	field.P = _cdbg.ToPdfObject()
	if field.T == nil || field.T.String() == "" {
		field.T = _dg.MakeString(_b.Sprintf("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072\u0065\u0020\u0025\u0064", pageNum))
	}
	_cdbg.AddAnnotation(field.PdfAnnotationWidget.PdfAnnotation)
	if _ddgg._afab == _ddgg._debg.AcroForm {
		_ddgg._afab = _ddgg.Reader.AcroForm
	}
	_dedcd := _ddgg._afab
	if _dedcd == nil {
		_dedcd = NewPdfAcroForm()
	}
	_dedcd.SigFlags = _dg.MakeInteger(3)
	if _dedcd.NeedAppearances != nil {
		_dedcd.NeedAppearances = nil
	}
	_acff := append(_dedcd.AllFields(), field.PdfField)
	_dedcd.Fields = &_acff
	_ddgg.ReplaceAcroForm(_dedcd)
	_ddgg.UpdatePage(_cdbg)
	_ddgg._ggdd[_cadgf] = _cdbg
	if _, _bdag := field.V.GetDocMDPPermission(); _bdag {
		_ddgg._accbd = NewPermissions(field.V)
	}
	return nil
}

type pdfFontType3 struct {
	fontCommon
	_aedbb *_dg.PdfIndirectObject

	// These fields are specific to Type 3 fonts.
	CharProcs  _dg.PdfObject
	Encoding   _dg.PdfObject
	FontBBox   _dg.PdfObject
	FontMatrix _dg.PdfObject
	FirstChar  _dg.PdfObject
	LastChar   _dg.PdfObject
	Widths     _dg.PdfObject
	Resources  _dg.PdfObject
	_bcde      map[_bd.CharCode]float64
	_geffb     _bd.TextEncoder
}

func (_edcde *PdfPattern) getDict() *_dg.PdfObjectDictionary {
	if _bcfgf, _cdce := _edcde._eacce.(*_dg.PdfIndirectObject); _cdce {
		_gecd, _fcbgf := _bcfgf.PdfObject.(*_dg.PdfObjectDictionary)
		if !_fcbgf {
			return nil
		}
		return _gecd
	} else if _fegef, _bfgc := _edcde._eacce.(*_dg.PdfObjectStream); _bfgc {
		return _fegef.PdfObjectDictionary
	} else {
		_ag.Log.Debug("\u0054r\u0079\u0069\u006e\u0067\u0020\u0074\u006f a\u0063\u0063\u0065\u0073\u0073\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063t \u0074\u0079\u0070e\u0020\u0028\u0025\u0054\u0029", _edcde._eacce)
		return nil
	}
}
func _edcf(_bcfcb *_fc.ImageBase) (_decbg Image) {
	_decbg.Width = int64(_bcfcb.Width)
	_decbg.Height = int64(_bcfcb.Height)
	_decbg.BitsPerComponent = int64(_bcfcb.BitsPerComponent)
	_decbg.ColorComponents = _bcfcb.ColorComponents
	_decbg.Data = _bcfcb.Data
	_decbg._gfbb = _bcfcb.Decode
	_decbg._dgeb = _bcfcb.Alpha
	return _decbg
}
func (_gfeg *PdfShading) getShadingDict() (*_dg.PdfObjectDictionary, error) {
	_gfebe := _gfeg._bcfbg
	if _gdaaf, _gdbbg := _gfebe.(*_dg.PdfIndirectObject); _gdbbg {
		_ffdb, _ecafgd := _gdaaf.PdfObject.(*_dg.PdfObjectDictionary)
		if !_ecafgd {
			return nil, _dg.ErrTypeError
		}
		return _ffdb, nil
	} else if _gfeea, _bccfdd := _gfebe.(*_dg.PdfObjectStream); _bccfdd {
		return _gfeea.PdfObjectDictionary, nil
	} else if _gfdda, _cagcc := _gfebe.(*_dg.PdfObjectDictionary); _cagcc {
		return _gfdda, nil
	} else {
		_ag.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0063\u0063\u0065s\u0073\u0020\u0073\u0068\u0061\u0064\u0069n\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079")
		return nil, _dg.ErrTypeError
	}
}

// PdfActionGoTo represents a GoTo action.
type PdfActionGoTo struct {
	*PdfAction
	D _dg.PdfObject
}

// SetContentStream updates the content stream with specified encoding.
// If encoding is null, will use the xform.Filter object or Raw encoding if not set.
func (_egfec *XObjectForm) SetContentStream(content []byte, encoder _dg.StreamEncoder) error {
	_gbgg := content
	if encoder == nil {
		if _egfec.Filter != nil {
			encoder = _egfec.Filter
		} else {
			encoder = _dg.NewRawEncoder()
		}
	}
	_afde, _aefgb := encoder.EncodeBytes(_gbgg)
	if _aefgb != nil {
		return _aefgb
	}
	_gbgg = _afde
	_egfec.Stream = _gbgg
	_egfec.Filter = encoder
	return nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_dfcaf *PdfColorspaceSpecialSeparation) ToPdfObject() _dg.PdfObject {
	_begd := _dg.MakeArray(_dg.MakeName("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"))
	_begd.Append(_dfcaf.ColorantName)
	_begd.Append(_dfcaf.AlternateSpace.ToPdfObject())
	_begd.Append(_dfcaf.TintTransform.ToPdfObject())
	if _dfcaf._fecg != nil {
		_dfcaf._fecg.PdfObject = _begd
		return _dfcaf._fecg
	}
	return _begd
}

// ToPdfObject implements interface PdfModel.
func (_eab *PdfAnnotationStamp) ToPdfObject() _dg.PdfObject {
	_eab.PdfAnnotation.ToPdfObject()
	_daee := _eab._cdf
	_afdd := _daee.PdfObject.(*_dg.PdfObjectDictionary)
	_eab.PdfAnnotationMarkup.appendToPdfDictionary(_afdd)
	_afdd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0053\u0074\u0061m\u0070"))
	_afdd.SetIfNotNil("\u004e\u0061\u006d\u0065", _eab.Name)
	return _daee
}

// ToPdfObject converts PdfAcroForm to a PdfObject, i.e. an indirect object containing the
// AcroForm dictionary.
func (_bceab *PdfAcroForm) ToPdfObject() _dg.PdfObject {
	_egcbf := _bceab._bebfe
	_cgdfg := _egcbf.PdfObject.(*_dg.PdfObjectDictionary)
	if _bceab.Fields != nil {
		_cfgbc := _dg.PdfObjectArray{}
		for _, _gcbdc := range *_bceab.Fields {
			_afda := _gcbdc.GetContext()
			if _afda != nil {
				_cfgbc.Append(_afda.ToPdfObject())
			} else {
				_cfgbc.Append(_gcbdc.ToPdfObject())
			}
		}
		_cgdfg.Set("\u0046\u0069\u0065\u006c\u0064\u0073", &_cfgbc)
	}
	if _bceab.NeedAppearances != nil {
		_cgdfg.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _bceab.NeedAppearances)
	} else {
		if _cdgdf := _cgdfg.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _cdgdf != nil {
			_cgdfg.Remove("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073")
		}
	}
	if _bceab.SigFlags != nil {
		_cgdfg.Set("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073", _bceab.SigFlags)
	}
	if _bceab.CO != nil {
		_cgdfg.Set("\u0043\u004f", _bceab.CO)
	}
	if _bceab.DR != nil {
		_cgdfg.Set("\u0044\u0052", _bceab.DR.ToPdfObject())
	}
	if _bceab.DA != nil {
		_cgdfg.Set("\u0044\u0041", _bceab.DA)
	}
	if _bceab.Q != nil {
		_cgdfg.Set("\u0051", _bceab.Q)
	}
	if _bceab.XFA != nil {
		_cgdfg.Set("\u0058\u0046\u0041", _bceab.XFA)
	}
	if _bceab.ADBEEchoSign != nil {
		_cgdfg.Set("\u0041\u0044\u0042\u0045\u005f\u0045\u0063\u0068\u006f\u0053\u0069\u0067\u006e", _bceab.ADBEEchoSign)
	}
	return _egcbf
}
func (_abaec *pdfCIDFontType2) baseFields() *fontCommon { return &_abaec.fontCommon }

// GetContentStream returns the XObject Form's content stream.
func (_eccbcc *XObjectForm) GetContentStream() ([]byte, error) {
	_bebge, _fdfaa := _dg.DecodeStream(_eccbcc._ebaeb)
	if _fdfaa != nil {
		return nil, _fdfaa
	}
	return _bebge, nil
}

const (
	FieldFlagClear             FieldFlag = 0
	FieldFlagReadOnly          FieldFlag = 1
	FieldFlagRequired          FieldFlag = (1 << 1)
	FieldFlagNoExport          FieldFlag = (2 << 1)
	FieldFlagNoToggleToOff     FieldFlag = (1 << 14)
	FieldFlagRadio             FieldFlag = (1 << 15)
	FieldFlagPushbutton        FieldFlag = (1 << 16)
	FieldFlagRadiosInUnision   FieldFlag = (1 << 25)
	FieldFlagMultiline         FieldFlag = (1 << 12)
	FieldFlagPassword          FieldFlag = (1 << 13)
	FieldFlagFileSelect        FieldFlag = (1 << 20)
	FieldFlagDoNotScroll       FieldFlag = (1 << 23)
	FieldFlagComb              FieldFlag = (1 << 24)
	FieldFlagRichText          FieldFlag = (1 << 26)
	FieldFlagDoNotSpellCheck   FieldFlag = (1 << 22)
	FieldFlagCombo             FieldFlag = (1 << 17)
	FieldFlagEdit              FieldFlag = (1 << 18)
	FieldFlagSort              FieldFlag = (1 << 19)
	FieldFlagMultiSelect       FieldFlag = (1 << 21)
	FieldFlagCommitOnSelChange FieldFlag = (1 << 27)
)

func (_afa *PdfReader) newPdfActionMovieFromDict(_fgf *_dg.PdfObjectDictionary) (*PdfActionMovie, error) {
	return &PdfActionMovie{Annotation: _fgf.Get("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e"), T: _fgf.Get("\u0054"), Operation: _fgf.Get("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn")}, nil
}

// GetColorspaceByName returns the colorspace with the specified name from the page resources.
func (_eaadbe *PdfPageResources) GetColorspaceByName(keyName _dg.PdfObjectName) (PdfColorspace, bool) {
	_gebed, _bdcbe := _eaadbe.GetColorspaces()
	if _bdcbe != nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _bdcbe)
		return nil, false
	}
	if _gebed == nil {
		return nil, false
	}
	_aagba, _fagac := _gebed.Colorspaces[string(keyName)]
	if !_fagac {
		return nil, false
	}
	return _aagba, true
}
func (_afege *PdfReader) lookupPageByObject(_fcef _dg.PdfObject) (*PdfPage, error) {
	return nil, _bf.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}
func _bedbb(_gaca _dg.PdfObject) (*fontFile, error) {
	_ag.Log.Trace("\u006e\u0065\u0077\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0046\u0072\u006f\u006dP\u0064f\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u006f\u0062\u006a\u003d\u0025\u0073", _gaca)
	_fcgea := &fontFile{}
	_gaca = _dg.TraceToDirectObject(_gaca)
	_ceaec, _dcaga := _gaca.(*_dg.PdfObjectStream)
	if !_dcaga {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020F\u006f\u006et\u0046\u0069\u006c\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0028\u0025\u0054\u0029", _gaca)
		return nil, _dg.ErrTypeError
	}
	_cced := _ceaec.PdfObjectDictionary
	_cefbb, _ebdc := _dg.DecodeStream(_ceaec)
	if _ebdc != nil {
		return nil, _ebdc
	}
	_dcaede, _dcaga := _dg.GetNameVal(_cced.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_dcaga {
		_fcgea._geffe = _dcaede
		if _dcaede == "\u0054\u0079\u0070\u0065\u0031\u0043" {
			_ag.Log.Debug("T\u0079\u0070\u0065\u0031\u0043\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u006e\u006f\u0074 s\u0075\u0070\u0070o\u0072t\u0065\u0064")
			return nil, ErrType1CFontNotSupported
		}
	}
	_facg, _ := _dg.GetIntVal(_cced.Get("\u004ce\u006e\u0067\u0074\u0068\u0031"))
	_gcecf, _ := _dg.GetIntVal(_cced.Get("\u004ce\u006e\u0067\u0074\u0068\u0032"))
	if _facg > len(_cefbb) {
		_facg = len(_cefbb)
	}
	if _facg+_gcecf > len(_cefbb) {
		_gcecf = len(_cefbb) - _facg
	}
	_ebgcc := _cefbb[:_facg]
	var _gceec []byte
	if _gcecf > 0 {
		_gceec = _cefbb[_facg : _facg+_gcecf]
	}
	if _facg > 0 && _gcecf > 0 {
		_ccegb := _fcgea.loadFromSegments(_ebgcc, _gceec)
		if _ccegb != nil {
			return nil, _ccegb
		}
	}
	return _fcgea, nil
}

// GetEncryptionMethod returns a descriptive information string about the encryption method used.
func (_ggcdb *PdfReader) GetEncryptionMethod() string {
	_cfdgf := _ggcdb._baad.GetCrypter()
	return _cfdgf.String()
}
func (_gga *PdfReader) newPdfActionSetOCGStateFromDict(_aeea *_dg.PdfObjectDictionary) (*PdfActionSetOCGState, error) {
	return &PdfActionSetOCGState{State: _aeea.Get("\u0053\u0074\u0061t\u0065"), PreserveRB: _aeea.Get("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042")}, nil
}

// ToPdfObject implements interface PdfModel.
func (_bcd *PdfActionMovie) ToPdfObject() _dg.PdfObject {
	_bcd.PdfAction.ToPdfObject()
	_aadf := _bcd._cbd
	_bge := _aadf.PdfObject.(*_dg.PdfObjectDictionary)
	_bge.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeMovie)))
	_bge.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e", _bcd.Annotation)
	_bge.SetIfNotNil("\u0054", _bcd.T)
	_bge.SetIfNotNil("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn", _bcd.Operation)
	return _aadf
}
func _gecgd(_cfbce _bbg.StdFont) pdfFontSimple {
	_gbdae := _cfbce.Descriptor()
	return pdfFontSimple{fontCommon: fontCommon{_bcga: "\u0054\u0079\u0070e\u0031", _ecbf: _cfbce.Name()}, _bfdee: _cfbce.GetMetricsTable(), _gagbe: &PdfFontDescriptor{FontName: _dg.MakeName(string(_gbdae.Name)), FontFamily: _dg.MakeName(_gbdae.Family), FontWeight: _dg.MakeFloat(float64(_gbdae.Weight)), Flags: _dg.MakeInteger(int64(_gbdae.Flags)), FontBBox: _dg.MakeArrayFromFloats(_gbdae.BBox[:]), ItalicAngle: _dg.MakeFloat(_gbdae.ItalicAngle), Ascent: _dg.MakeFloat(_gbdae.Ascent), Descent: _dg.MakeFloat(_gbdae.Descent), CapHeight: _dg.MakeFloat(_gbdae.CapHeight), XHeight: _dg.MakeFloat(_gbdae.XHeight), StemV: _dg.MakeFloat(_gbdae.StemV), StemH: _dg.MakeFloat(_gbdae.StemH)}, _dbdb: _cfbce.Encoder()}
}

const (
	_ PdfOutputIntentType = iota
	PdfOutputIntentTypeA1
	PdfOutputIntentTypeA2
	PdfOutputIntentTypeA3
	PdfOutputIntentTypeA4
	PdfOutputIntentTypeX
)

// SetContext sets the sub annotation (context).
func (_cabb *PdfAnnotation) SetContext(ctx PdfModel) { _cabb._egcg = ctx }
func _dgag(_fcbec *_dg.PdfObjectDictionary) (*PdfShadingType2, error) {
	_geegd := PdfShadingType2{}
	_abgad := _fcbec.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _abgad == nil {
		_ag.Log.Debug("R\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u003a\u0020\u0020\u0043\u006f\u006f\u0072d\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_ffcac, _dcggf := _abgad.(*_dg.PdfObjectArray)
	if !_dcggf {
		_ag.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _abgad)
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _ffcac.Len() != 4 {
		_ag.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0034\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _ffcac.Len())
		return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	}
	_geegd.Coords = _ffcac
	if _addca := _fcbec.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _addca != nil {
		_addca = _dg.TraceToDirectObject(_addca)
		_bacec, _bfddc := _addca.(*_dg.PdfObjectArray)
		if !_bfddc {
			_ag.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _addca)
			return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_geegd.Domain = _bacec
	}
	_abgad = _fcbec.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _abgad == nil {
		_ag.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_geegd.Function = []PdfFunction{}
	if _dddbae, _dafgb := _abgad.(*_dg.PdfObjectArray); _dafgb {
		for _, _ddabc := range _dddbae.Elements() {
			_eeeec, _gagee := _agec(_ddabc)
			if _gagee != nil {
				_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _gagee)
				return nil, _gagee
			}
			_geegd.Function = append(_geegd.Function, _eeeec)
		}
	} else {
		_bgdgfa, _gffgf := _agec(_abgad)
		if _gffgf != nil {
			_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _gffgf)
			return nil, _gffgf
		}
		_geegd.Function = append(_geegd.Function, _bgdgfa)
	}
	if _cgee := _fcbec.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _cgee != nil {
		_cgee = _dg.TraceToDirectObject(_cgee)
		_dagac, _ecafc := _cgee.(*_dg.PdfObjectArray)
		if !_ecafc {
			_ag.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cgee)
			return nil, _dg.ErrTypeError
		}
		if _dagac.Len() != 2 {
			_ag.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _dagac.Len())
			return nil, ErrInvalidAttribute
		}
		_geegd.Extend = _dagac
	}
	return &_geegd, nil
}

// PdfFont represents an underlying font structure which can be of type:
// - Type0
// - Type1
// - TrueType
// etc.
type PdfFont struct{ _cadf pdfFont }

// ToPdfObject implements interface PdfModel.
func (_cgcdf *PdfAnnotationCaret) ToPdfObject() _dg.PdfObject {
	_cgcdf.PdfAnnotation.ToPdfObject()
	_eadd := _cgcdf._cdf
	_gffe := _eadd.PdfObject.(*_dg.PdfObjectDictionary)
	_cgcdf.PdfAnnotationMarkup.appendToPdfDictionary(_gffe)
	_gffe.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0043\u0061\u0072e\u0074"))
	_gffe.SetIfNotNil("\u0052\u0044", _cgcdf.RD)
	_gffe.SetIfNotNil("\u0053\u0079", _cgcdf.Sy)
	return _eadd
}

// GetContext returns the PdfField context which is the more specific field data type, e.g. PdfFieldButton
// for a button field.
func (_adafg *PdfField) GetContext() PdfModel { return _adafg._bdfg }

// IsPush returns true if the button field represents a push button, false otherwise.
func (_ggab *PdfFieldButton) IsPush() bool { return _ggab.GetType() == ButtonTypePush }

// NewPdfAnnotationPrinterMark returns a new printermark annotation.
func NewPdfAnnotationPrinterMark() *PdfAnnotationPrinterMark {
	_cgca := NewPdfAnnotation()
	_ggge := &PdfAnnotationPrinterMark{}
	_ggge.PdfAnnotation = _cgca
	_cgca.SetContext(_ggge)
	return _ggge
}

// GetContainingPdfObject implements interface PdfModel.
func (_fbfcf *PdfSignature) GetContainingPdfObject() _dg.PdfObject { return _fbfcf._bbda }
func (_bagg *PdfWriter) setDocumentIDs(_aacece, _cefff string) {
	_bagg._agbeg = _dg.MakeArray(_dg.MakeHexString(_aacece), _dg.MakeHexString(_cefff))
}

// NewPdfColorspaceSpecialPattern returns a new pattern color.
func NewPdfColorspaceSpecialPattern() *PdfColorspaceSpecialPattern {
	return &PdfColorspaceSpecialPattern{}
}

// GetNumComponents returns the number of color components (1 for grayscale).
func (_acdc *PdfColorDeviceGray) GetNumComponents() int { return 1 }

// NewPdfAnnotationTrapNet returns a new trapnet annotation.
func NewPdfAnnotationTrapNet() *PdfAnnotationTrapNet {
	_gbeg := NewPdfAnnotation()
	_bab := &PdfAnnotationTrapNet{}
	_bab.PdfAnnotation = _gbeg
	_gbeg.SetContext(_bab)
	return _bab
}
func (_fgdfc *PdfColorspaceSpecialSeparation) String() string {
	return "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"
}

// SetSamples convert samples to byte-data and sets for the image.
// NOTE: The method resamples the data and this could lead to high memory usage,
// especially on large images. It should be used only when it is not possible
// to work with the image byte data directly.
func (_fcgaf *Image) SetSamples(samples []uint32) {
	if _fcgaf.BitsPerComponent < 8 {
		samples = _fcgaf.samplesAddPadding(samples)
	}
	_gaeg := _fcd.ResampleUint32(samples, int(_fcgaf.BitsPerComponent), 8)
	_efde := make([]byte, len(_gaeg))
	for _fgccd, _abaeg := range _gaeg {
		_efde[_fgccd] = byte(_abaeg)
	}
	_fcgaf.Data = _efde
}
func _deag() *modelManager {
	_afac := modelManager{}
	_afac._aaddgg = map[PdfModel]_dg.PdfObject{}
	_afac._eccbc = map[_dg.PdfObject]PdfModel{}
	return &_afac
}
func (_adea *LTV) enable(_fbfec, _gafgd []*_bb.Certificate, _fcbgd string) error {
	_cdab, _cgaab, _facb := _adea.buildCertChain(_fbfec, _gafgd)
	if _facb != nil {
		return _facb
	}
	_afcge, _facb := _adea.getCerts(_cdab)
	if _facb != nil {
		return _facb
	}
	var _dcfda, _afgce [][]byte
	if _adea.OCSPClient != nil {
		_dcfda, _facb = _adea.getOCSPs(_cdab, _cgaab)
		if _facb != nil {
			return _facb
		}
	}
	if _adea.CRLClient != nil {
		_afgce, _facb = _adea.getCRLs(_cdab)
		if _facb != nil {
			return _facb
		}
	}
	_ddbde := _adea._abca
	_fdccg, _facb := _ddbde.AddCerts(_afcge)
	if _facb != nil {
		return _facb
	}
	_dagc, _facb := _ddbde.AddOCSPs(_dcfda)
	if _facb != nil {
		return _facb
	}
	_feeg, _facb := _ddbde.AddCRLs(_afgce)
	if _facb != nil {
		return _facb
	}
	if _fcbgd != "" {
		_ddbde.VRI[_fcbgd] = &VRI{Cert: _fdccg, OCSP: _dagc, CRL: _feeg}
	}
	_adea._ebdgd.SetDSS(_ddbde)
	return nil
}

// ImageToRGB returns an error since an image cannot be defined in a pattern colorspace.
func (_bebg *PdfColorspaceSpecialPattern) ImageToRGB(img Image) (Image, error) {
	_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0061\u0074\u0074\u0065\u0072n \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	return img, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0066\u006f\u0072\u0020\u0069m\u0061\u0067\u0065\u0020\u0028p\u0061\u0074t\u0065\u0072\u006e\u0029")
}

// GetAction returns the PDF action for the annotation link.
func (_cec *PdfAnnotationLink) GetAction() (*PdfAction, error) {
	if _cec._ece != nil {
		return _cec._ece, nil
	}
	if _cec.A == nil {
		return nil, nil
	}
	if _cec._abea == nil {
		return nil, nil
	}
	_cfab, _ced := _cec._abea.loadAction(_cec.A)
	if _ced != nil {
		return nil, _ced
	}
	_cec._ece = _cfab
	return _cec._ece, nil
}

// Duplicate creates a duplicate page based on the current one and returns it.
func (_ffcf *PdfPage) Duplicate() *PdfPage {
	_gbcfc := *_ffcf
	_gbcfc._bfdge = _dg.MakeDict()
	_gbcfc._cggbe = _dg.MakeIndirectObject(_gbcfc._bfdge)
	_gbcfc._gaed = *_gbcfc._bfdge
	return &_gbcfc
}

// PdfBorderEffect represents a PDF border effect.
type PdfBorderEffect struct {
	S *BorderEffect
	I *float64
}

// SetContext set the sub annotation (context).
func (_dcfecg *PdfShading) SetContext(ctx PdfModel) { _dcfecg._eeddb = ctx }

// GetContext returns the annotation context which contains the specific type-dependent context.
// The context represents the subannotation.
func (_def *PdfAnnotation) GetContext() PdfModel {
	if _def == nil {
		return nil
	}
	return _def._egcg
}
func (_edeag *PdfReader) newPdfAcroFormFromDict(_gaabd *_dg.PdfIndirectObject, _effd *_dg.PdfObjectDictionary) (*PdfAcroForm, error) {
	_dacb := NewPdfAcroForm()
	if _gaabd != nil {
		_dacb._bebfe = _gaabd
		_gaabd.PdfObject = _dg.MakeDict()
	}
	if _dggbd := _effd.Get("\u0046\u0069\u0065\u006c\u0064\u0073"); _dggbd != nil && !_dg.IsNullObject(_dggbd) {
		_aacfb, _gcad := _dg.GetArray(_dggbd)
		if !_gcad {
			return nil, _b.Errorf("\u0066i\u0065\u006c\u0064\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _dggbd)
		}
		var _eagfa []*PdfField
		for _, _baegg := range _aacfb.Elements() {
			_eddfd, _gcdfd := _dg.GetIndirect(_baegg)
			if !_gcdfd {
				if _, _gdde := _baegg.(*_dg.PdfObjectNull); _gdde {
					_ag.Log.Trace("\u0053k\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072 \u006e\u0075\u006c\u006c\u0020\u0066\u0069\u0065\u006c\u0064")
					continue
				}
				_ag.Log.Debug("\u0046\u0069\u0065\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0064 \u0069\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _baegg)
				return nil, _b.Errorf("\u0066\u0069\u0065l\u0064\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_ggeec, _ccbe := _edeag.newPdfFieldFromIndirectObject(_eddfd, nil)
			if _ccbe != nil {
				return nil, _ccbe
			}
			_ag.Log.Trace("\u0041\u0063\u0072\u006fFo\u0072\u006d\u0020\u0046\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u002b\u0076", *_ggeec)
			_eagfa = append(_eagfa, _ggeec)
		}
		_dacb.Fields = &_eagfa
	}
	if _cefaa := _effd.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _cefaa != nil {
		_bdeg, _bfce := _dg.GetBool(_cefaa)
		if _bfce {
			_dacb.NeedAppearances = _bdeg
		} else {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u0065\u0065\u0064\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063e\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _cefaa)
		}
	}
	if _cefge := _effd.Get("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073"); _cefge != nil {
		_bffce, _fafgb := _dg.GetInt(_cefge)
		if _fafgb {
			_dacb.SigFlags = _bffce
		} else {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073 \u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cefge)
		}
	}
	if _eabcg := _effd.Get("\u0043\u004f"); _eabcg != nil {
		_adab, _dcfbc := _dg.GetArray(_eabcg)
		if _dcfbc {
			_dacb.CO = _adab
		} else {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u004f\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _eabcg)
		}
	}
	if _cecgg := _effd.Get("\u0044\u0052"); _cecgg != nil {
		if _bbff, _gdgcg := _dg.GetDict(_cecgg); _gdgcg {
			_aafe, _dcac := NewPdfPageResourcesFromDict(_bbff)
			if _dcac != nil {
				_ag.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0044R\u003a\u0020\u0025\u0076", _dcac)
				return nil, _dcac
			}
			_dacb.DR = _aafe
		} else {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0052\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _cecgg)
		}
	}
	if _degd := _effd.Get("\u0044\u0041"); _degd != nil {
		_becd, _gacgb := _dg.GetString(_degd)
		if _gacgb {
			_dacb.DA = _becd
		} else {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0041\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _degd)
		}
	}
	if _bccff := _effd.Get("\u0051"); _bccff != nil {
		_adcda, _afcfb := _dg.GetInt(_bccff)
		if _afcfb {
			_dacb.Q = _adcda
		} else {
			_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0051\u0020\u0069\u006e\u0076a\u006ci\u0064 \u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _bccff)
		}
	}
	if _afbfa := _effd.Get("\u0058\u0046\u0041"); _afbfa != nil {
		_dacb.XFA = _afbfa
	}
	if _dgfa := _effd.Get("\u0041\u0044\u0042\u0045\u005f\u0045\u0063\u0068\u006f\u0053\u0069\u0067\u006e"); _dgfa != nil {
		_dacb.ADBEEchoSign = _dgfa
	}
	_dacb.ToPdfObject()
	return _dacb, nil
}
func (_bdedc *PdfWriter) hasObject(_begfd _dg.PdfObject) bool {
	_, _cabca := _bdedc._fdbfa[_begfd]
	return _cabca
}

// PdfActionSetOCGState represents a SetOCGState action.
type PdfActionSetOCGState struct {
	*PdfAction
	State      _dg.PdfObject
	PreserveRB _dg.PdfObject
}

func (_aef *PdfReader) newPdfActionURIFromDict(_eaa *_dg.PdfObjectDictionary) (*PdfActionURI, error) {
	return &PdfActionURI{URI: _eaa.Get("\u0055\u0052\u0049"), IsMap: _eaa.Get("\u0049\u0073\u004da\u0070")}, nil
}
func _beabe(_debca *fontCommon) *pdfFontSimple { return &pdfFontSimple{fontCommon: *_debca} }
func (_adgd *PdfReader) newPdfActionFromIndirectObject(_afcg *_dg.PdfIndirectObject) (*PdfAction, error) {
	_bdb, _ddg := _afcg.PdfObject.(*_dg.PdfObjectDictionary)
	if !_ddg {
		return nil, _b.Errorf("\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u006e\u006f\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _adgd._cadfa.GetModelFromPrimitive(_bdb); model != nil {
		_daf, _ccc := model.(*PdfAction)
		if !_ccc {
			return nil, _b.Errorf("\u0063\u0061c\u0068\u0065\u0064\u0020\u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0061\u0063ti\u006f\u006e")
		}
		return _daf, nil
	}
	_bce := &PdfAction{}
	_bce._cbd = _afcg
	_adgd._cadfa.Register(_bdb, _bce)
	if _abe := _bdb.Get("\u0054\u0079\u0070\u0065"); _abe != nil {
		_acd, _fae := _abe.(*_dg.PdfObjectName)
		if !_fae {
			_ag.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _abe)
		} else {
			if *_acd != "\u0041\u0063\u0074\u0069\u006f\u006e" {
				_ag.Log.Trace("\u0055\u006e\u0073u\u0073\u0070\u0065\u0063t\u0065\u0064\u0020\u0054\u0079\u0070\u0065 \u0021\u003d\u0020\u0041\u0063\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", *_acd)
			}
			_bce.Type = _acd
		}
	}
	if _abg := _bdb.Get("\u004e\u0065\u0078\u0074"); _abg != nil {
		_bce.Next = _abg
	}
	if _edc := _bdb.Get("\u0053"); _edc != nil {
		_bce.S = _edc
	}
	_dac, _bed := _bce.S.(*_dg.PdfObjectName)
	if !_bed {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _bce.S)
		return nil, _b.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0053\u0020\u006f\u0062\u006a\u0065c\u0074 \u0074y\u0070e\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _bce.S)
	}
	_badc := PdfActionType(_dac.String())
	switch _badc {
	case ActionTypeGoTo:
		_ebg, _fde := _adgd.newPdfActionGotoFromDict(_bdb)
		if _fde != nil {
			return nil, _fde
		}
		_ebg.PdfAction = _bce
		_bce._bg = _ebg
		return _bce, nil
	case ActionTypeGoToR:
		_fcb, _bcdb := _adgd.newPdfActionGotoRFromDict(_bdb)
		if _bcdb != nil {
			return nil, _bcdb
		}
		_fcb.PdfAction = _bce
		_bce._bg = _fcb
		return _bce, nil
	case ActionTypeGoToE:
		_abed, _efe := _adgd.newPdfActionGotoEFromDict(_bdb)
		if _efe != nil {
			return nil, _efe
		}
		_abed.PdfAction = _bce
		_bce._bg = _abed
		return _bce, nil
	case ActionTypeLaunch:
		_de, _eff := _adgd.newPdfActionLaunchFromDict(_bdb)
		if _eff != nil {
			return nil, _eff
		}
		_de.PdfAction = _bce
		_bce._bg = _de
		return _bce, nil
	case ActionTypeThread:
		_cdd, _bfd := _adgd.newPdfActionThreadFromDict(_bdb)
		if _bfd != nil {
			return nil, _bfd
		}
		_cdd.PdfAction = _bce
		_bce._bg = _cdd
		return _bce, nil
	case ActionTypeURI:
		_cdg, _dcc := _adgd.newPdfActionURIFromDict(_bdb)
		if _dcc != nil {
			return nil, _dcc
		}
		_cdg.PdfAction = _bce
		_bce._bg = _cdg
		return _bce, nil
	case ActionTypeSound:
		_cbf, _aga := _adgd.newPdfActionSoundFromDict(_bdb)
		if _aga != nil {
			return nil, _aga
		}
		_cbf.PdfAction = _bce
		_bce._bg = _cbf
		return _bce, nil
	case ActionTypeMovie:
		_gce, _cgc := _adgd.newPdfActionMovieFromDict(_bdb)
		if _cgc != nil {
			return nil, _cgc
		}
		_gce.PdfAction = _bce
		_bce._bg = _gce
		return _bce, nil
	case ActionTypeHide:
		_adf, _dccc := _adgd.newPdfActionHideFromDict(_bdb)
		if _dccc != nil {
			return nil, _dccc
		}
		_adf.PdfAction = _bce
		_bce._bg = _adf
		return _bce, nil
	case ActionTypeNamed:
		_efg, _gdfb := _adgd.newPdfActionNamedFromDict(_bdb)
		if _gdfb != nil {
			return nil, _gdfb
		}
		_efg.PdfAction = _bce
		_bce._bg = _efg
		return _bce, nil
	case ActionTypeSubmitForm:
		_gag, _gfda := _adgd.newPdfActionSubmitFormFromDict(_bdb)
		if _gfda != nil {
			return nil, _gfda
		}
		_gag.PdfAction = _bce
		_bce._bg = _gag
		return _bce, nil
	case ActionTypeResetForm:
		_fg, _ebd := _adgd.newPdfActionResetFormFromDict(_bdb)
		if _ebd != nil {
			return nil, _ebd
		}
		_fg.PdfAction = _bce
		_bce._bg = _fg
		return _bce, nil
	case ActionTypeImportData:
		_gbe, _dgc := _adgd.newPdfActionImportDataFromDict(_bdb)
		if _dgc != nil {
			return nil, _dgc
		}
		_gbe.PdfAction = _bce
		_bce._bg = _gbe
		return _bce, nil
	case ActionTypeSetOCGState:
		_cgf, _bddb := _adgd.newPdfActionSetOCGStateFromDict(_bdb)
		if _bddb != nil {
			return nil, _bddb
		}
		_cgf.PdfAction = _bce
		_bce._bg = _cgf
		return _bce, nil
	case ActionTypeRendition:
		_db, _agd := _adgd.newPdfActionRenditionFromDict(_bdb)
		if _agd != nil {
			return nil, _agd
		}
		_db.PdfAction = _bce
		_bce._bg = _db
		return _bce, nil
	case ActionTypeTrans:
		_cfe, _cde := _adgd.newPdfActionTransFromDict(_bdb)
		if _cde != nil {
			return nil, _cde
		}
		_cfe.PdfAction = _bce
		_bce._bg = _cfe
		return _bce, nil
	case ActionTypeGoTo3DView:
		_gceg, _dbb := _adgd.newPdfActionGoTo3DViewFromDict(_bdb)
		if _dbb != nil {
			return nil, _dbb
		}
		_gceg.PdfAction = _bce
		_bce._bg = _gceg
		return _bce, nil
	case ActionTypeJavaScript:
		_feff, _adfc := _adgd.newPdfActionJavaScriptFromDict(_bdb)
		if _adfc != nil {
			return nil, _adfc
		}
		_feff.PdfAction = _bce
		_bce._bg = _feff
		return _bce, nil
	}
	_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006eg\u0020u\u006ek\u006eo\u0077\u006e\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073", _badc)
	return nil, nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_bdbdg *PdfColorspaceDeviceGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gfcfa := vals[0]
	if _gfcfa < 0.0 || _gfcfa > 1.0 {
		_ag.Log.Debug("\u0049\u006eco\u006d\u0070\u0061t\u0069\u0062\u0069\u006city\u003a R\u0061\u006e\u0067\u0065\u0020\u006f\u0075ts\u0069\u0064\u0065\u0020\u005b\u0030\u002c1\u005d")
	}
	if _gfcfa < 0.0 {
		_gfcfa = 0.0
	} else if _gfcfa > 1.0 {
		_gfcfa = 1.0
	}
	return NewPdfColorDeviceGray(_gfcfa), nil
}

// NewReaderOpts generates a default `ReaderOpts` instance.
func NewReaderOpts() *ReaderOpts { return &ReaderOpts{Password: "", LazyLoad: true} }

// Val returns the color value.
func (_gege *PdfColorDeviceGray) Val() float64    { return float64(*_gege) }
func (_gcdc *PdfColorspaceCalRGB) String() string { return "\u0043\u0061\u006c\u0052\u0047\u0042" }

// ToPdfObject recursively builds the Outline tree PDF object.
func (_ffadf *PdfOutline) ToPdfObject() _dg.PdfObject {
	_dfee := _ffadf._bbef
	_geccd := _dfee.PdfObject.(*_dg.PdfObjectDictionary)
	_geccd.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073"))
	if _ffadf.First != nil {
		_geccd.Set("\u0046\u0069\u0072s\u0074", _ffadf.First.ToPdfObject())
	}
	if _ffadf.Last != nil {
		_geccd.Set("\u004c\u0061\u0073\u0074", _ffadf.Last.GetContext().GetContainingPdfObject())
	}
	if _ffadf.Parent != nil {
		_geccd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _ffadf.Parent.GetContext().GetContainingPdfObject())
	}
	if _ffadf.Count != nil {
		_geccd.Set("\u0043\u006f\u0075n\u0074", _dg.MakeInteger(*_ffadf.Count))
	}
	return _dfee
}

// NewPdfPage returns a new PDF page.
func NewPdfPage() *PdfPage {
	_fabed := PdfPage{}
	_fabed._bfdge = _dg.MakeDict()
	_fabed.Resources = NewPdfPageResources()
	_gdfe := _dg.PdfIndirectObject{}
	_gdfe.PdfObject = _fabed._bfdge
	_fabed._cggbe = &_gdfe
	_fabed._gaed = *_fabed._bfdge
	return &_fabed
}

// ToPdfObject implements interface PdfModel.
func (_acg *PdfActionHide) ToPdfObject() _dg.PdfObject {
	_acg.PdfAction.ToPdfObject()
	_dfg := _acg._cbd
	_dd := _dfg.PdfObject.(*_dg.PdfObjectDictionary)
	_dd.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeHide)))
	_dd.SetIfNotNil("\u0054", _acg.T)
	_dd.SetIfNotNil("\u0048", _acg.H)
	return _dfg
}

// ToPdfObject implements interface PdfModel.
func (_cfa *PdfActionNamed) ToPdfObject() _dg.PdfObject {
	_cfa.PdfAction.ToPdfObject()
	_bgg := _cfa._cbd
	_feg := _bgg.PdfObject.(*_dg.PdfObjectDictionary)
	_feg.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeNamed)))
	_feg.SetIfNotNil("\u004e", _cfa.N)
	return _bgg
}

// ReaderToWriterOpts options used to generate a PdfWriter.
type ReaderToWriterOpts struct {
	SkipAcroForm        bool
	SkipInfo            bool
	SkipNameDictionary  bool
	SkipNamedDests      bool
	SkipOCProperties    bool
	SkipOutlines        bool
	SkipPageLabels      bool
	SkipRotation        bool
	SkipMetadata        bool
	PageProcessCallback PageProcessCallback

	// Deprecated: will be removed in v4. Use PageProcessCallback instead.
	PageCallback PageCallback
}
type fontFile struct {
	_ecgd  string
	_geffe string
	_beaeb _bd.SimpleEncoder
}

// ToPdfObject implements interface PdfModel.
func (_fdga *PdfAnnotationUnderline) ToPdfObject() _dg.PdfObject {
	_fdga.PdfAnnotation.ToPdfObject()
	_bdaa := _fdga._cdf
	_bafc := _bdaa.PdfObject.(*_dg.PdfObjectDictionary)
	_fdga.PdfAnnotationMarkup.appendToPdfDictionary(_bafc)
	_bafc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee"))
	_bafc.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _fdga.QuadPoints)
	return _bdaa
}

// NewReaderForText makes a new PdfReader for an input PDF content string. For use in testing.
func NewReaderForText(txt string) *PdfReader {
	return &PdfReader{_addfg: map[_dg.PdfObject]struct{}{}, _cadfa: _deag(), _baad: _dg.NewParserFromString(txt)}
}

// ImageToRGB convert 1-component grayscale data to 3-component RGB.
func (_egbac *PdfColorspaceDeviceGray) ImageToRGB(img Image) (Image, error) {
	if img.ColorComponents != 1 {
		return img, _bf.New("\u0074\u0068e \u0070\u0072\u006fv\u0069\u0064\u0065\u0064 im\u0061ge\u0020\u0069\u0073\u0020\u006e\u006f\u0074 g\u0072\u0061\u0079\u0020\u0073\u0063\u0061l\u0065")
	}
	_eadge, _egd := _fc.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._dgeb, img._gfbb)
	if _egd != nil {
		return img, _egd
	}
	_bcaeb, _egd := _fc.NRGBAConverter.Convert(_eadge)
	if _egd != nil {
		return img, _egd
	}
	_bcgg := _edcf(_bcaeb.Base())
	_ag.Log.Trace("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079\u0020\u002d>\u0020\u0052\u0047\u0042")
	_ag.Log.Trace("s\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", img.Data)
	_ag.Log.Trace("\u0052G\u0042 \u0073\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", _bcgg.Data)
	_ag.Log.Trace("\u0025\u0076\u0020\u002d\u003e\u0020\u0025\u0076", img, _bcgg)
	return _bcgg, nil
}

// NewPdfSignatureReferenceDocMDP returns PdfSignatureReference for the transformParams.
func NewPdfSignatureReferenceDocMDP(transformParams *PdfTransformParamsDocMDP) *PdfSignatureReference {
	return &PdfSignatureReference{Type: _dg.MakeName("\u0053\u0069\u0067\u0052\u0065\u0066"), TransformMethod: _dg.MakeName("\u0044\u006f\u0063\u004d\u0044\u0050"), TransformParams: transformParams.ToPdfObject()}
}
func (_facec *PdfAcroForm) signatureFields() []*PdfFieldSignature {
	var _gceeg []*PdfFieldSignature
	for _, _fbbd := range _facec.AllFields() {
		switch _cbfd := _fbbd.GetContext().(type) {
		case *PdfFieldSignature:
			_dddgf := _cbfd
			_gceeg = append(_gceeg, _dddgf)
		}
	}
	return _gceeg
}

// SetAlpha sets the alpha layer for the image.
func (_abcbb *Image) SetAlpha(alpha []byte) { _abcbb._dgeb = alpha }

// NewPdfActionSetOCGState returns a new "named" action.
func NewPdfActionSetOCGState() *PdfActionSetOCGState {
	_ccd := NewPdfAction()
	_gfc := &PdfActionSetOCGState{}
	_gfc.PdfAction = _ccd
	_ccd.SetContext(_gfc)
	return _gfc
}

// BaseFont returns the font's "BaseFont" field.
func (_ffef *PdfFont) BaseFont() string { return _ffef.baseFields()._ecbf }

// PdfOutline represents a PDF outline dictionary (Table 152 - p. 376).
type PdfOutline struct {
	PdfOutlineTreeNode
	Parent *PdfOutlineTreeNode
	Count  *int64
	_bbef  *_dg.PdfIndirectObject
}

// PdfShadingType1 is a Function-based shading.
type PdfShadingType1 struct {
	*PdfShading
	Domain   *_dg.PdfObjectArray
	Matrix   *_dg.PdfObjectArray
	Function []PdfFunction
}

// ImageToRGB convert an indexed image to RGB.
func (_dfcd *PdfColorspaceSpecialIndexed) ImageToRGB(img Image) (Image, error) {
	N := _dfcd.Base.GetNumComponents()
	if N < 1 {
		return Image{}, _b.Errorf("\u0062\u0061d \u0062\u0061\u0073e\u0020\u0063\u006f\u006cors\u0070ac\u0065\u0020\u004e\u0075\u006d\u0043\u006fmp\u006f\u006e\u0065\u006e\u0074\u0073\u003d%\u0064", N)
	}
	_edec := _fc.NewImageBase(int(img.Width), int(img.Height), 8, N, nil, img._dgeb, img._gfbb)
	_edadd := _fcd.NewReader(img.getBase())
	_dcaa := _fcd.NewWriter(_edec)
	var (
		_bbcg  uint32
		_egdbe int
		_fefa  error
	)
	for {
		_bbcg, _fefa = _edadd.ReadSample()
		if _fefa == _cf.EOF {
			break
		} else if _fefa != nil {
			return img, _fefa
		}
		_egdbe = int(_bbcg)
		_ag.Log.Trace("\u0049\u006ed\u0065\u0078\u0065\u0064\u003a\u0020\u0069\u006e\u0064\u0065\u0078\u003d\u0025\u0064\u0020\u004e\u003d\u0025\u0064\u0020\u006c\u0075t=\u0025\u0064", _egdbe, N, len(_dfcd._bcea))
		if (_egdbe+1)*N > len(_dfcd._bcea) {
			_egdbe = len(_dfcd._bcea)/N - 1
			_ag.Log.Trace("C\u006c\u0069\u0070\u0070in\u0067 \u0074\u006f\u0020\u0069\u006ed\u0065\u0078\u003a\u0020\u0025\u0064", _egdbe)
			if _egdbe < 0 {
				_ag.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069p\u0020\u0069\u006e\u0064\u0065\u0078.\u0020\u0049\u0073\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006ce\u0020\u0064\u0061\u006d\u0061\u0067\u0065\u0064\u003f")
				break
			}
		}
		for _dgdf := _egdbe * N; _dgdf < (_egdbe+1)*N; _dgdf++ {
			if _fefa = _dcaa.WriteSample(uint32(_dfcd._bcea[_dgdf])); _fefa != nil {
				return img, _fefa
			}
		}
	}
	return _dfcd.Base.ImageToRGB(_edcf(&_edec))
}

// SetRotation sets the rotation of all pages added to writer. The rotation is
// specified in degrees and must be a multiple of 90.
// The Rotate field of individual pages has priority over the global rotation.
func (_gfdde *PdfWriter) SetRotation(rotate int64) error {
	_cbdegg, _cfdff := _dg.GetDict(_gfdde._gbgb)
	if !_cfdff {
		return ErrTypeCheck
	}
	_cbdegg.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _dg.MakeInteger(rotate))
	return nil
}

// GetPreviousRevision returns the previous revision of PdfReader for the Pdf document
func (_cbca *PdfReader) GetPreviousRevision() (*PdfReader, error) {
	if _cbca._baad.GetRevisionNumber() == 0 {
		return nil, _bf.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065xi\u0073\u0074")
	}
	if _dabdc, _eebd := _cbca._ggafe[_cbca]; _eebd {
		return _dabdc, nil
	}
	_afeb, _gcfdc := _cbca._baad.GetPreviousRevisionReadSeeker()
	if _gcfdc != nil {
		return nil, _gcfdc
	}
	_ccgacb, _gcfdc := _dcdd(_afeb, _cbca._gcecfe, _cbca._eadef, "\u006do\u0064\u0065\u006c\u003aG\u0065\u0074\u0050\u0072\u0065v\u0069o\u0075s\u0052\u0065\u0076\u0069\u0073\u0069\u006fn")
	if _gcfdc != nil {
		return nil, _gcfdc
	}
	_cbca._aggag[_cbca._baad.GetRevisionNumber()-1] = _ccgacb
	_cbca._ggafe[_cbca] = _ccgacb
	_ccgacb._ggafe = _cbca._ggafe
	return _ccgacb, nil
}
func (_bbdca *PdfPage) getParentResources() (*PdfPageResources, error) {
	_bgfff := _bbdca.Parent
	for _bgfff != nil {
		_gbcbd, _cgddc := _dg.GetDict(_bgfff)
		if !_cgddc {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020n\u006f\u0064\u0065")
			return nil, _bf.New("i\u006e\u0076\u0061\u006cid\u0020p\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		if _fddc := _gbcbd.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _fddc != nil {
			_egfg, _feegd := _dg.GetDict(_fddc)
			if !_feegd {
				return nil, _bf.New("i\u006e\u0076\u0061\u006cid\u0020r\u0065\u0073\u006f\u0075\u0072c\u0065\u0020\u0064\u0069\u0063\u0074")
			}
			_eacac, _gcde := NewPdfPageResourcesFromDict(_egfg)
			if _gcde != nil {
				return nil, _gcde
			}
			return _eacac, nil
		}
		_bgfff = _gbcbd.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, nil
}

// ToPdfObject converts date to a PDF string object.
func (_efadb *PdfDate) ToPdfObject() _dg.PdfObject {
	_fdcee := _b.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _efadb._bgfdb, _efadb._gbbge, _efadb._cbad, _efadb._cade, _efadb._ccgbg, _efadb._aacfe, _efadb._ggbdg, _efadb._bdde, _efadb._gafe)
	return _dg.MakeString(_fdcee)
}

// ReplacePage replaces the original page to a new page.
func (_dbgcb *PdfAppender) ReplacePage(pageNum int, page *PdfPage) {
	_bagc := pageNum - 1
	for _acae := range _dbgcb._ggdd {
		if _acae == _bagc {
			_fgdf := page.Duplicate()
			_dccea(_fgdf)
			_dbgcb._ggdd[_acae] = _fgdf
		}
	}
}
func (_fcbgb *pdfFontSimple) baseFields() *fontCommon { return &_fcbgb.fontCommon }

var ErrColorOutOfRange = _bf.New("\u0063o\u006co\u0072\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")

func (_gbda *PdfReader) newPdfAnnotationWidgetFromDict(_efgad *_dg.PdfObjectDictionary) (*PdfAnnotationWidget, error) {
	_cbbg := PdfAnnotationWidget{}
	_cbbg.H = _efgad.Get("\u0048")
	_cbbg.MK = _efgad.Get("\u004d\u004b")
	_cbbg.A = _efgad.Get("\u0041")
	_cbbg.AA = _efgad.Get("\u0041\u0041")
	_cbbg.BS = _efgad.Get("\u0042\u0053")
	_cbbg.Parent = _efgad.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	return &_cbbg, nil
}

// SetPdfKeywords sets the Keywords attribute of the output PDF.
func SetPdfKeywords(keywords string) { _fgefgf.Lock(); defer _fgefgf.Unlock(); _abbgf = keywords }
func _fdfee(_becc _dg.PdfObject) (*PdfColorspaceDeviceN, error) {
	_dddba := NewPdfColorspaceDeviceN()
	if _ceag, _bcac := _becc.(*_dg.PdfIndirectObject); _bcac {
		_dddba._effag = _ceag
	}
	_becc = _dg.TraceToDirectObject(_becc)
	_ecbab, _gdfgg := _becc.(*_dg.PdfObjectArray)
	if !_gdfgg {
		return nil, _b.Errorf("\u0064\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if _ecbab.Len() != 4 && _ecbab.Len() != 5 {
		return nil, _b.Errorf("\u0064\u0065\u0076ic\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0049\u006ec\u006fr\u0072e\u0063t\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_becc = _ecbab.Get(0)
	_bbfa, _gdfgg := _becc.(*_dg.PdfObjectName)
	if !_gdfgg {
		return nil, _b.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_bbfa != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
		return nil, _b.Errorf("\u0064\u0065v\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020na\u006d\u0065")
	}
	_becc = _ecbab.Get(1)
	_becc = _dg.TraceToDirectObject(_becc)
	_edcec, _gdfgg := _becc.(*_dg.PdfObjectArray)
	if !_gdfgg {
		return nil, _b.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0049\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_dddba.ColorantNames = _edcec
	_becc = _ecbab.Get(2)
	_caeg, _afdfa := NewPdfColorspaceFromPdfObject(_becc)
	if _afdfa != nil {
		return nil, _afdfa
	}
	_dddba.AlternateSpace = _caeg
	_fffa, _afdfa := _agec(_ecbab.Get(3))
	if _afdfa != nil {
		return nil, _afdfa
	}
	_dddba.TintTransform = _fffa
	if _ecbab.Len() == 5 {
		_cgge, _eefae := _ecef(_ecbab.Get(4))
		if _eefae != nil {
			return nil, _eefae
		}
		_dddba.Attributes = _cgge
	}
	return _dddba, nil
}

// GetSamples converts the raw byte slice into samples which are stored in a uint32 bit array.
// Each sample is represented by BitsPerComponent consecutive bits in the raw data.
// NOTE: The method resamples the image byte data before returning the result and
// this could lead to high memory usage, especially on large images. It should
// be avoided, when possible. It is recommended to access the Data field of the
// image directly or use the ColorAt method to extract individual pixels.
func (_ffca *Image) GetSamples() []uint32 {
	_cbccd := _fcd.ResampleBytes(_ffca.Data, int(_ffca.BitsPerComponent))
	if _ffca.BitsPerComponent < 8 {
		_cbccd = _ffca.samplesTrimPadding(_cbccd)
	}
	_ffgcc := int(_ffca.Width) * int(_ffca.Height) * _ffca.ColorComponents
	if len(_cbccd) < _ffgcc {
		_ag.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u0066\u0065\u0077\u0020\u0073a\u006d\u0070\u006c\u0065\u0073\u0020(\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064\u0029", len(_cbccd), _ffgcc)
		return _cbccd
	} else if len(_cbccd) > _ffgcc {
		_ag.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u006d\u0061\u006e\u0079\u0020s\u0061\u006d\u0070\u006c\u0065\u0073 \u0028\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078p\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064", len(_cbccd), _ffgcc)
		_cbccd = _cbccd[:_ffgcc]
	}
	return _cbccd
}

// DSS represents a Document Security Store dictionary.
// The DSS dictionary contains both global and signature specific validation
// information. The certificates and revocation data in the `Certs`, `OCSPs`,
// and `CRLs` fields can be used to validate any signature in the document.
// Additionally, the VRI entry contains validation data per signature.
// The keys in the VRI entry are calculated as upper(hex(sha1(sig.Contents))).
// The values are VRI dictionaries containing certificates and revocation
// information used for validating a single signature.
// See ETSI TS 102 778-4 V1.1.1 for more information.
type DSS struct {
	_agdcg *_dg.PdfIndirectObject
	Certs  []*_dg.PdfObjectStream
	OCSPs  []*_dg.PdfObjectStream
	CRLs   []*_dg.PdfObjectStream
	VRI    map[string]*VRI
	_agbg  map[string]*_dg.PdfObjectStream
	_bggg  map[string]*_dg.PdfObjectStream
	_dcdf  map[string]*_dg.PdfObjectStream
}

func (_fag *PdfReader) newPdfActionGotoEFromDict(_fbf *_dg.PdfObjectDictionary) (*PdfActionGoToE, error) {
	_cdb, _ege := _bccf(_fbf.Get("\u0046"))
	if _ege != nil {
		return nil, _ege
	}
	return &PdfActionGoToE{D: _fbf.Get("\u0044"), NewWindow: _fbf.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), T: _fbf.Get("\u0054"), F: _cdb}, nil
}

// ToPdfObject returns colorspace in a PDF object format [name stream]
func (_fcegf *PdfColorspaceICCBased) ToPdfObject() _dg.PdfObject {
	_dgef := &_dg.PdfObjectArray{}
	_dgef.Append(_dg.MakeName("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"))
	var _gebb *_dg.PdfObjectStream
	if _fcegf._dafg != nil {
		_gebb = _fcegf._dafg
	} else {
		_gebb = &_dg.PdfObjectStream{}
	}
	_fceab := _dg.MakeDict()
	_fceab.Set("\u004e", _dg.MakeInteger(int64(_fcegf.N)))
	if _fcegf.Alternate != nil {
		_fceab.Set("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e", _fcegf.Alternate.ToPdfObject())
	}
	if _fcegf.Metadata != nil {
		_fceab.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _fcegf.Metadata)
	}
	if _fcegf.Range != nil {
		var _badd []_dg.PdfObject
		for _, _efaa := range _fcegf.Range {
			_badd = append(_badd, _dg.MakeFloat(_efaa))
		}
		_fceab.Set("\u0052\u0061\u006eg\u0065", _dg.MakeArray(_badd...))
	}
	_fceab.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dg.MakeInteger(int64(len(_fcegf.Data))))
	_gebb.Stream = _fcegf.Data
	_gebb.PdfObjectDictionary = _fceab
	_dgef.Append(_gebb)
	if _fcegf._fbdff != nil {
		_fcegf._fbdff.PdfObject = _dgef
		return _fcegf._fbdff
	}
	return _dgef
}

// ValidateSignatures validates digital signatures in the document.
func (_bdfff *PdfReader) ValidateSignatures(handlers []SignatureHandler) ([]SignatureValidationResult, error) {
	if _bdfff.AcroForm == nil {
		return nil, nil
	}
	if _bdfff.AcroForm.Fields == nil {
		return nil, nil
	}
	type sigFieldPair struct {
		_fbaba *PdfSignature
		_ecfcf *PdfField
		_fgccg SignatureHandler
	}
	var _efcae []*sigFieldPair
	for _, _fcdff := range _bdfff.AcroForm.AllFields() {
		if _fcdff.V == nil {
			continue
		}
		if _eabe, _bfecf := _dg.GetDict(_fcdff.V); _bfecf {
			if _bcgcfb, _bgcdg := _dg.GetNameVal(_eabe.Get("\u0054\u0079\u0070\u0065")); _bgcdg && (_bcgcfb == "\u0053\u0069\u0067" || _bcgcfb == "\u0044\u006f\u0063T\u0069\u006d\u0065\u0053\u0074\u0061\u006d\u0070") {
				_aegaa, _gbdfad := _dg.GetIndirect(_fcdff.V)
				if !_gbdfad {
					_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0020\u0069s\u0020\u006e\u0069\u006c")
					return nil, ErrTypeCheck
				}
				_dbbae, _bdaba := _bdfff.newPdfSignatureFromIndirect(_aegaa)
				if _bdaba != nil {
					return nil, _bdaba
				}
				var _bffdc SignatureHandler
				for _, _ecae := range handlers {
					if _ecae.IsApplicable(_dbbae) {
						_bffdc = _ecae
						break
					}
				}
				_efcae = append(_efcae, &sigFieldPair{_fbaba: _dbbae, _ecfcf: _fcdff, _fgccg: _bffdc})
			}
		}
	}
	var _eefgg []SignatureValidationResult
	for _, _dfae := range _efcae {
		_fdggb := SignatureValidationResult{IsSigned: true, Fields: []*PdfField{_dfae._ecfcf}}
		if _dfae._fgccg == nil {
			_fdggb.Errors = append(_fdggb.Errors, "\u0068a\u006ed\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
			_eefgg = append(_eefgg, _fdggb)
			continue
		}
		_afggg, _bgaad := _dfae._fgccg.NewDigest(_dfae._fbaba)
		if _bgaad != nil {
			_fdggb.Errors = append(_fdggb.Errors, "\u0064\u0069\u0067e\u0073\u0074\u0020\u0065\u0072\u0072\u006f\u0072", _bgaad.Error())
			_eefgg = append(_eefgg, _fdggb)
			continue
		}
		_eeff := _dfae._fbaba.ByteRange
		if _eeff == nil {
			_fdggb.Errors = append(_fdggb.Errors, "\u0042\u0079\u0074\u0065\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
			_eefgg = append(_eefgg, _fdggb)
			continue
		}
		for _fgbce := 0; _fgbce < _eeff.Len(); _fgbce = _fgbce + 2 {
			_dadgc, _ := _dg.GetNumberAsInt64(_eeff.Get(_fgbce))
			_bface, _ := _dg.GetIntVal(_eeff.Get(_fgbce + 1))
			if _, _bgadcf := _bdfff._efcfa.Seek(_dadgc, _cf.SeekStart); _bgadcf != nil {
				return nil, _bgadcf
			}
			_dgdef := make([]byte, _bface)
			if _, _cfcef := _bdfff._efcfa.Read(_dgdef); _cfcef != nil {
				return nil, _cfcef
			}
			_afggg.Write(_dgdef)
		}
		var _bcgfd SignatureValidationResult
		if _gbcbe, _gbade := _dfae._fgccg.(SignatureHandlerDocMDP); _gbade {
			_bcgfd, _bgaad = _gbcbe.ValidateWithOpts(_dfae._fbaba, _afggg, SignatureHandlerDocMDPParams{Parser: _bdfff._baad})
		} else {
			_bcgfd, _bgaad = _dfae._fgccg.Validate(_dfae._fbaba, _afggg)
		}
		if _bgaad != nil {
			_ag.Log.Debug("E\u0052\u0052\u004f\u0052: \u0025v\u0020\u0028\u0025\u0054\u0029 \u002d\u0020\u0073\u006b\u0069\u0070", _bgaad, _dfae._fgccg)
			_bcgfd.Errors = append(_bcgfd.Errors, _bgaad.Error())
		}
		_bcgfd.Name = _dfae._fbaba.Name.Decoded()
		_bcgfd.Reason = _dfae._fbaba.Reason.Decoded()
		if _dfae._fbaba.M != nil {
			_ddbgf, _fadd := NewPdfDate(_dfae._fbaba.M.String())
			if _fadd != nil {
				_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fadd)
				_bcgfd.Errors = append(_bcgfd.Errors, _fadd.Error())
				continue
			}
			_bcgfd.Date = _ddbgf
		}
		_bcgfd.ContactInfo = _dfae._fbaba.ContactInfo.Decoded()
		_bcgfd.Location = _dfae._fbaba.Location.Decoded()
		_bcgfd.Fields = _fdggb.Fields
		_eefgg = append(_eefgg, _bcgfd)
	}
	return _eefgg, nil
}

// ToPdfObject implements interface PdfModel.
func (_cbc *PdfActionThread) ToPdfObject() _dg.PdfObject {
	_cbc.PdfAction.ToPdfObject()
	_adg := _cbc._cbd
	_efb := _adg.PdfObject.(*_dg.PdfObjectDictionary)
	_efb.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeThread)))
	if _cbc.F != nil {
		_efb.Set("\u0046", _cbc.F.ToPdfObject())
	}
	_efb.SetIfNotNil("\u0044", _cbc.D)
	_efb.SetIfNotNil("\u0042", _cbc.B)
	return _adg
}

// WriteToFile writes the output PDF to file.
func (_ceebg *PdfWriter) WriteToFile(outputFilePath string) error {
	_cafb, _bffcc := _eb.Create(outputFilePath)
	if _bffcc != nil {
		return _bffcc
	}
	defer _cafb.Close()
	return _ceebg.Write(_cafb)
}
func _ceafc(_fcca *_dg.PdfObjectArray) (float64, error) {
	_fgagfe, _gedcdd := _fcca.ToFloat64Array()
	if _gedcdd != nil {
		_ag.Log.Debug("\u0042\u0061\u0064\u0020\u004d\u0061\u0074\u0074\u0065\u0020\u0061\u0072\u0072\u0061\u0079:\u0020m\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fcca, _gedcdd)
	}
	switch len(_fgagfe) {
	case 1:
		return _fgagfe[0], nil
	case 3:
		_fcecc := PdfColorspaceDeviceRGB{}
		_agebg, _fafdd := _fcecc.ColorFromFloats(_fgagfe)
		if _fafdd != nil {
			return 0.0, _fafdd
		}
		return _agebg.(*PdfColorDeviceRGB).ToGray().Val(), nil
	case 4:
		_efaag := PdfColorspaceDeviceCMYK{}
		_adebge, _cdbege := _efaag.ColorFromFloats(_fgagfe)
		if _cdbege != nil {
			return 0.0, _cdbege
		}
		_gcdd, _cdbege := _efaag.ColorToRGB(_adebge.(*PdfColorDeviceCMYK))
		if _cdbege != nil {
			return 0.0, _cdbege
		}
		return _gcdd.(*PdfColorDeviceRGB).ToGray().Val(), nil
	}
	_gedcdd = _bf.New("\u0062a\u0064 \u004d\u0061\u0074\u0074\u0065\u0020\u0063\u006f\u006c\u006f\u0072")
	_ag.Log.Error("\u0074\u006f\u0047ra\u0079\u003a\u0020\u006d\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fcca, _gedcdd)
	return 0.0, _gedcdd
}

// GetNumComponents returns the number of color components.
func (_eca *PdfColorspaceICCBased) GetNumComponents() int { return _eca.N }

// ToPdfObject returns the PDF representation of the tiling pattern.
func (_cbec *PdfTilingPattern) ToPdfObject() _dg.PdfObject {
	_cbec.PdfPattern.ToPdfObject()
	_aaga := _cbec.getDict()
	if _cbec.PaintType != nil {
		_aaga.Set("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e", _cbec.PaintType)
	}
	if _cbec.TilingType != nil {
		_aaga.Set("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _cbec.TilingType)
	}
	if _cbec.BBox != nil {
		_aaga.Set("\u0042\u0042\u006f\u0078", _cbec.BBox.ToPdfObject())
	}
	if _cbec.XStep != nil {
		_aaga.Set("\u0058\u0053\u0074e\u0070", _cbec.XStep)
	}
	if _cbec.YStep != nil {
		_aaga.Set("\u0059\u0053\u0074e\u0070", _cbec.YStep)
	}
	if _cbec.Resources != nil {
		_aaga.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _cbec.Resources.ToPdfObject())
	}
	if _cbec.Matrix != nil {
		_aaga.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _cbec.Matrix)
	}
	return _cbec._eacce
}
func _eabg(_gdee *_dg.PdfObjectDictionary) (*PdfShadingType6, error) {
	_adfbbe := PdfShadingType6{}
	_ebacf := _gdee.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _ebacf == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_baeeec, _cdcdfg := _ebacf.(*_dg.PdfObjectInteger)
	if !_cdcdfg {
		_ag.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ebacf)
		return nil, _dg.ErrTypeError
	}
	_adfbbe.BitsPerCoordinate = _baeeec
	_ebacf = _gdee.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _ebacf == nil {
		_ag.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_baeeec, _cdcdfg = _ebacf.(*_dg.PdfObjectInteger)
	if !_cdcdfg {
		_ag.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _ebacf)
		return nil, _dg.ErrTypeError
	}
	_adfbbe.BitsPerComponent = _baeeec
	_ebacf = _gdee.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _ebacf == nil {
		_ag.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_baeeec, _cdcdfg = _ebacf.(*_dg.PdfObjectInteger)
	if !_cdcdfg {
		_ag.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _ebacf)
		return nil, _dg.ErrTypeError
	}
	_adfbbe.BitsPerComponent = _baeeec
	_ebacf = _gdee.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _ebacf == nil {
		_ag.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_cefcf, _cdcdfg := _ebacf.(*_dg.PdfObjectArray)
	if !_cdcdfg {
		_ag.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ebacf)
		return nil, _dg.ErrTypeError
	}
	_adfbbe.Decode = _cefcf
	if _gbafgc := _gdee.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _gbafgc != nil {
		_adfbbe.Function = []PdfFunction{}
		if _eedec, _ddada := _gbafgc.(*_dg.PdfObjectArray); _ddada {
			for _, _aabg := range _eedec.Elements() {
				_cfdgfg, _cbgdd := _agec(_aabg)
				if _cbgdd != nil {
					_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cbgdd)
					return nil, _cbgdd
				}
				_adfbbe.Function = append(_adfbbe.Function, _cfdgfg)
			}
		} else {
			_fcddf, _dgdfd := _agec(_gbafgc)
			if _dgdfd != nil {
				_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dgdfd)
				return nil, _dgdfd
			}
			_adfbbe.Function = append(_adfbbe.Function, _fcddf)
		}
	}
	return &_adfbbe, nil
}
func _dcdd(_ageca _cf.ReadSeeker, _ebafd *ReaderOpts, _aaab bool, _aaaga string) (*PdfReader, error) {
	if _ebafd == nil {
		_ebafd = NewReaderOpts()
	}
	_ccgb := *_ebafd
	_bcfga := &PdfReader{_efcfa: _ageca, _addfg: map[_dg.PdfObject]struct{}{}, _cadfa: _deag(), _dadcef: _ebafd.LazyLoad, _cfbbb: _ebafd.ComplianceMode, _eadef: _aaab, _gcecfe: &_ccgb}
	_cbdf, _faeea := _aeegd("\u0072")
	if _faeea != nil {
		return nil, _faeea
	}
	_bcfga._aafgg = _cbdf
	var _fcagff *_dg.PdfParser
	if !_bcfga._cfbbb {
		_fcagff, _faeea = _dg.NewParser(_ageca)
	} else {
		_fcagff, _faeea = _dg.NewCompliancePdfParser(_ageca)
	}
	if _faeea != nil {
		return nil, _faeea
	}
	_bcfga._baad = _fcagff
	_afccb, _faeea := _bcfga.IsEncrypted()
	if _faeea != nil {
		return nil, _faeea
	}
	if !_afccb {
		_faeea = _bcfga.loadStructure()
		if _faeea != nil {
			return nil, _faeea
		}
	} else if _aaab {
		_afbde, _agfag := _bcfga.Decrypt([]byte(_ebafd.Password))
		if _agfag != nil {
			return nil, _agfag
		}
		if !_afbde {
			return nil, _bf.New("\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0064\u0065c\u0072\u0079\u0070\u0074\u0020\u0070\u0061\u0073\u0073w\u006f\u0072\u0064\u0020p\u0072\u006f\u0074\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u002d\u0020\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0073\u0070\u0065\u0063\u0069\u0066y\u0020\u0070\u0061s\u0073\u0020\u0074\u006f\u0020\u0044\u0065\u0063\u0072\u0079\u0070\u0074")
		}
	}
	_bcfga._ggafe = make(map[*PdfReader]*PdfReader)
	_bcfga._aggag = make([]*PdfReader, _fcagff.GetRevisionNumber())
	return _bcfga, nil
}

// PdfAnnotationWidget represents Widget annotations.
// Note: Widget annotations are used to display form fields.
// (Section 12.5.6.19).
type PdfAnnotationWidget struct {
	*PdfAnnotation
	H      _dg.PdfObject
	MK     _dg.PdfObject
	A      _dg.PdfObject
	AA     _dg.PdfObject
	BS     _dg.PdfObject
	Parent _dg.PdfObject
	_cgg   *PdfField
	_dgac  bool
}

func _caff(_debac *_dg.PdfObjectDictionary) (*PdfShadingType3, error) {
	_eegabg := PdfShadingType3{}
	_ddccb := _debac.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _ddccb == nil {
		_ag.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0043\u006f\u006f\u0072\u0064\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_cecfb, _fdgff := _ddccb.(*_dg.PdfObjectArray)
	if !_fdgff {
		_ag.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ddccb)
		return nil, _dg.ErrTypeError
	}
	if _cecfb.Len() != 6 {
		_ag.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0036\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _cecfb.Len())
		return nil, ErrInvalidAttribute
	}
	_eegabg.Coords = _cecfb
	if _egaeg := _debac.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _egaeg != nil {
		_egaeg = _dg.TraceToDirectObject(_egaeg)
		_cbded, _addbcf := _egaeg.(*_dg.PdfObjectArray)
		if !_addbcf {
			_ag.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _egaeg)
			return nil, _dg.ErrTypeError
		}
		_eegabg.Domain = _cbded
	}
	_ddccb = _debac.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _ddccb == nil {
		_ag.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_eegabg.Function = []PdfFunction{}
	if _ggdda, _bfcce := _ddccb.(*_dg.PdfObjectArray); _bfcce {
		for _, _babca := range _ggdda.Elements() {
			_bfcbg, _fdgad := _agec(_babca)
			if _fdgad != nil {
				_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _fdgad)
				return nil, _fdgad
			}
			_eegabg.Function = append(_eegabg.Function, _bfcbg)
		}
	} else {
		_gfbbb, _cdcdf := _agec(_ddccb)
		if _cdcdf != nil {
			_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cdcdf)
			return nil, _cdcdf
		}
		_eegabg.Function = append(_eegabg.Function, _gfbbb)
	}
	if _cagfb := _debac.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _cagfb != nil {
		_cagfb = _dg.TraceToDirectObject(_cagfb)
		_aeaa, _geacb := _cagfb.(*_dg.PdfObjectArray)
		if !_geacb {
			_ag.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cagfb)
			return nil, _dg.ErrTypeError
		}
		if _aeaa.Len() != 2 {
			_ag.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _aeaa.Len())
			return nil, ErrInvalidAttribute
		}
		_eegabg.Extend = _aeaa
	}
	return &_eegabg, nil
}

// ToPdfObject implements interface PdfModel.
func (_dfga *PdfActionRendition) ToPdfObject() _dg.PdfObject {
	_dfga.PdfAction.ToPdfObject()
	_cbbd := _dfga._cbd
	_fed := _cbbd.PdfObject.(*_dg.PdfObjectDictionary)
	_fed.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeRendition)))
	_fed.SetIfNotNil("\u0052", _dfga.R)
	_fed.SetIfNotNil("\u0041\u004e", _dfga.AN)
	_fed.SetIfNotNil("\u004f\u0050", _dfga.OP)
	_fed.SetIfNotNil("\u004a\u0053", _dfga.JS)
	return _cbbd
}

// NewBorderStyle returns an initialized PdfBorderStyle.
func NewBorderStyle() *PdfBorderStyle { _begge := &PdfBorderStyle{}; return _begge }
func (_cddg *PdfReader) newPdfAnnotationProjectionFromDict(_fdfa *_dg.PdfObjectDictionary) (*PdfAnnotationProjection, error) {
	_aegab := &PdfAnnotationProjection{}
	_dedc, _ddce := _cddg.newPdfAnnotationMarkupFromDict(_fdfa)
	if _ddce != nil {
		return nil, _ddce
	}
	_aegab.PdfAnnotationMarkup = _dedc
	return _aegab, nil
}
func _bbgcf(_dgbe, _becbf string) string {
	if _ga.Contains(_dgbe, "\u002b") {
		_faeb := _ga.Split(_dgbe, "\u002b")
		if len(_faeb) == 2 {
			_dgbe = _faeb[1]
		}
	}
	return _becbf + "\u002b" + _dgbe
}

// GetXObjectByName gets XObject by name.
func (_abefd *PdfPage) GetXObjectByName(name _dg.PdfObjectName) (_dg.PdfObject, bool) {
	_ecgg, _dcgee := _abefd.Resources.XObject.(*_dg.PdfObjectDictionary)
	if !_dcgee {
		return nil, false
	}
	if _ccbcg := _ecgg.Get(name); _ccbcg != nil {
		return _ccbcg, true
	}
	return nil, false
}

// DecodeArray returns the range of color component values in the Lab colorspace.
func (_ffbbe *PdfColorspaceLab) DecodeArray() []float64 {
	_dfcac := []float64{0, 100}
	if _ffbbe.Range != nil && len(_ffbbe.Range) == 4 {
		_dfcac = append(_dfcac, _ffbbe.Range...)
	} else {
		_dfcac = append(_dfcac, -100, 100, -100, 100)
	}
	return _dfcac
}

// SetVersion sets the PDF version of the output file.
func (_gcadg *PdfWriter) SetVersion(majorVersion, minorVersion int) {
	_gcadg._efacd.Major = majorVersion
	_gcadg._efacd.Minor = minorVersion
}
func _bbece() string { _fgefgf.Lock(); defer _fgefgf.Unlock(); return _abbgf }

// NewPdfActionMovie returns a new "movie" action.
func NewPdfActionMovie() *PdfActionMovie {
	_ede := NewPdfAction()
	_fa := &PdfActionMovie{}
	_fa.PdfAction = _ede
	_ede.SetContext(_fa)
	return _fa
}

// ToPdfObject implements interface PdfModel.
func (_bcf *PdfActionSetOCGState) ToPdfObject() _dg.PdfObject {
	_bcf.PdfAction.ToPdfObject()
	_ccdd := _bcf._cbd
	_dda := _ccdd.PdfObject.(*_dg.PdfObjectDictionary)
	_dda.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeSetOCGState)))
	_dda.SetIfNotNil("\u0053\u0074\u0061t\u0065", _bcf.State)
	_dda.SetIfNotNil("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042", _bcf.PreserveRB)
	return _ccdd
}

// MergePageWith appends page content to source Pdf file page content.
func (_bfgd *PdfAppender) MergePageWith(pageNum int, page *PdfPage) error {
	_bdea := pageNum - 1
	var _ddea *PdfPage
	for _edgd, _bcgfa := range _bfgd._ggdd {
		if _edgd == _bdea {
			_ddea = _bcgfa
		}
	}
	if _ddea == nil {
		return _b.Errorf("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073o\u0075\u0072\u0063\u0065\u0020\u0064o\u0063\u0075\u006de\u006e\u0074", pageNum)
	}
	if _ddea._cggbe != nil && _ddea._cggbe.GetParser() == _bfgd._debg._baad {
		_ddea = _ddea.Duplicate()
		_bfgd._ggdd[_bdea] = _ddea
	}
	page = page.Duplicate()
	_dccea(page)
	_bgcce := _fgea(_ddea)
	_caf := _fgea(page)
	_bfcc := make(map[_dg.PdfObjectName]_dg.PdfObjectName)
	for _ebad := range _caf {
		if _, _aeae := _bgcce[_ebad]; _aeae {
			for _eee := 1; true; _eee++ {
				_decc := _dg.PdfObjectName(string(_ebad) + _fbb.Itoa(_eee))
				if _, _bffe := _bgcce[_decc]; !_bffe {
					_bfcc[_ebad] = _decc
					break
				}
			}
		}
	}
	_cffa, _fafd := page.GetContentStreams()
	if _fafd != nil {
		return _fafd
	}
	_baef, _fafd := _ddea.GetContentStreams()
	if _fafd != nil {
		return _fafd
	}
	for _efeb, _eaaf := range _cffa {
		for _cgfb, _gfce := range _bfcc {
			_eaaf = _ga.Replace(_eaaf, "\u002f"+string(_cgfb), "\u002f"+string(_gfce), -1)
		}
		_cffa[_efeb] = _eaaf
	}
	_baef = append(_baef, _cffa...)
	if _gagdg := _ddea.SetContentStreams(_baef, _dg.NewFlateEncoder()); _gagdg != nil {
		return _gagdg
	}
	_ddea._cadgg = append(_ddea._cadgg, page._cadgg...)
	if _ddea.Resources == nil {
		_ddea.Resources = NewPdfPageResources()
	}
	if page.Resources != nil {
		_ddea.Resources.Font = _bfgd.mergeResources(_ddea.Resources.Font, page.Resources.Font, _bfcc)
		_ddea.Resources.XObject = _bfgd.mergeResources(_ddea.Resources.XObject, page.Resources.XObject, _bfcc)
		_ddea.Resources.Properties = _bfgd.mergeResources(_ddea.Resources.Properties, page.Resources.Properties, _bfcc)
		if _ddea.Resources.ProcSet == nil {
			_ddea.Resources.ProcSet = page.Resources.ProcSet
		}
		_ddea.Resources.Shading = _bfgd.mergeResources(_ddea.Resources.Shading, page.Resources.Shading, _bfcc)
		_ddea.Resources.ExtGState = _bfgd.mergeResources(_ddea.Resources.ExtGState, page.Resources.ExtGState, _bfcc)
	}
	_ggcd, _fafd := _ddea.GetMediaBox()
	if _fafd != nil {
		return _fafd
	}
	_fdfe, _fafd := page.GetMediaBox()
	if _fafd != nil {
		return _fafd
	}
	var _gdb bool
	if _ggcd.Llx > _fdfe.Llx {
		_ggcd.Llx = _fdfe.Llx
		_gdb = true
	}
	if _ggcd.Lly > _fdfe.Lly {
		_ggcd.Lly = _fdfe.Lly
		_gdb = true
	}
	if _ggcd.Urx < _fdfe.Urx {
		_ggcd.Urx = _fdfe.Urx
		_gdb = true
	}
	if _ggcd.Ury < _fdfe.Ury {
		_ggcd.Ury = _fdfe.Ury
		_gdb = true
	}
	if _gdb {
		_ddea.MediaBox = _ggcd
	}
	return nil
}

// ColorToRGB converts a Lab color to an RGB color.
func (_gegb *PdfColorspaceLab) ColorToRGB(color PdfColor) (PdfColor, error) {
	_bage := func(_bgad float64) float64 {
		if _bgad >= 6.0/29 {
			return _bgad * _bgad * _bgad
		}
		return 108.0 / 841 * (_bgad - 4.0/29.0)
	}
	_ddbdf, _gecgb := color.(*PdfColorLab)
	if !_gecgb {
		_ag.Log.Debug("\u0069\u006e\u0070\u0075t \u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u006c\u0061\u0062")
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	LStar := _ddbdf.L()
	AStar := _ddbdf.A()
	BStar := _ddbdf.B()
	L := (LStar+16)/116 + AStar/500
	M := (LStar + 16) / 116
	N := (LStar+16)/116 - BStar/200
	X := _gegb.WhitePoint[0] * _bage(L)
	Y := _gegb.WhitePoint[1] * _bage(M)
	Z := _gegb.WhitePoint[2] * _bage(N)
	_eadda := 3.240479*X + -1.537150*Y + -0.498535*Z
	_fdgb := -0.969256*X + 1.875992*Y + 0.041556*Z
	_gceda := 0.055648*X + -0.204043*Y + 1.057311*Z
	_eadda = _cg.Min(_cg.Max(_eadda, 0), 1.0)
	_fdgb = _cg.Min(_cg.Max(_fdgb, 0), 1.0)
	_gceda = _cg.Min(_cg.Max(_gceda, 0), 1.0)
	return NewPdfColorDeviceRGB(_eadda, _fdgb, _gceda), nil
}

// NewPdfActionGoTo3DView returns a new "goTo3DView" action.
func NewPdfActionGoTo3DView() *PdfActionGoTo3DView {
	_beg := NewPdfAction()
	_bgc := &PdfActionGoTo3DView{}
	_bgc.PdfAction = _beg
	_beg.SetContext(_bgc)
	return _bgc
}
func _ebbce(_cfgcc *PdfField, _gdfcc _dg.PdfObject) error {
	switch _cfgcc.GetContext().(type) {
	case *PdfFieldText:
		switch _dggf := _gdfcc.(type) {
		case *_dg.PdfObjectName:
			_cagdc := _dggf
			_ag.Log.Debug("\u0055\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u003a\u0020\u0047\u006f\u0074 \u0056\u0020\u0061\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u003e\u0020c\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f s\u0074\u0072\u0069\u006e\u0067\u0020\u0027\u0025\u0073\u0027", _cagdc.String())
			_cfgcc.V = _dg.MakeEncodedString(_dggf.String(), true)
		case *_dg.PdfObjectString:
			_cfgcc.V = _dg.MakeEncodedString(_dggf.String(), true)
		default:
			_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0056\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u0028\u0025\u0023\u0076\u0029", _dggf, _dggf)
		}
	case *PdfFieldButton:
		switch _gdfcc.(type) {
		case *_dg.PdfObjectName:
			if len(_gdfcc.String()) > 0 {
				_cfgcc.V = _gdfcc
				_gcab(_cfgcc, _gdfcc)
			}
		case *_dg.PdfObjectString:
			if len(_gdfcc.String()) > 0 {
				_cfgcc.V = _dg.MakeName(_gdfcc.String())
				_gcab(_cfgcc, _cfgcc.V)
			}
		default:
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _cfgcc.PartialName(), _gdfcc)
			_cfgcc.V = _gdfcc
		}
	case *PdfFieldChoice:
		switch _gdfcc.(type) {
		case *_dg.PdfObjectName:
			if len(_gdfcc.String()) > 0 {
				_cfgcc.V = _dg.MakeString(_gdfcc.String())
				_gcab(_cfgcc, _gdfcc)
			}
		case *_dg.PdfObjectString:
			if len(_gdfcc.String()) > 0 {
				_cfgcc.V = _gdfcc
				_gcab(_cfgcc, _dg.MakeName(_gdfcc.String()))
			}
		default:
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _cfgcc.PartialName(), _gdfcc)
			_cfgcc.V = _gdfcc
		}
	case *PdfFieldSignature:
		_ag.Log.Debug("\u0054\u004f\u0044\u004f\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0061\u0070\u0070e\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0079\u0065\u0074\u003a\u0020\u0025\u0073\u002f\u0025v", _cfgcc.PartialName(), _gdfcc)
	}
	return nil
}

// SetOpenAction sets the OpenAction in the PDF catalog.
// The value shall be either an array defining a destination (12.3.2 "Destinations" PDF32000_2008),
// or an action dictionary representing an action (12.6 "Actions" PDF32000_2008).
func (_fbfac *PdfWriter) SetOpenAction(dest _dg.PdfObject) error {
	if dest == nil || _dg.IsNullObject(dest) {
		return nil
	}
	_fbfac._ecdf.Set("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e", dest)
	return _fbfac.addObjects(dest)
}

// Decrypt decrypts the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful,
// false otherwise.
func (_ddaa *PdfReader) Decrypt(password []byte) (bool, error) {
	_ccbec, _ebege := _ddaa._baad.Decrypt(password)
	if _ebege != nil {
		return false, _ebege
	}
	if !_ccbec {
		return false, nil
	}
	_ebege = _ddaa.loadStructure()
	if _ebege != nil {
		_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f \u006co\u0061d\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", _ebege)
		return false, _ebege
	}
	return true, nil
}

// HasXObjectByName checks if has XObject resource by name.
func (_caag *PdfPage) HasXObjectByName(name _dg.PdfObjectName) bool {
	_adfg, _caegc := _caag.Resources.XObject.(*_dg.PdfObjectDictionary)
	if !_caegc {
		return false
	}
	if _aecb := _adfg.Get(name); _aecb != nil {
		return true
	}
	return false
}

// GetContainingPdfObject returns the container of the DSS (indirect object).
func (_cacb *DSS) GetContainingPdfObject() _dg.PdfObject { return _cacb._agdcg }
func (_gcf *PdfReader) newPdfActionGotoFromDict(_baf *_dg.PdfObjectDictionary) (*PdfActionGoTo, error) {
	return &PdfActionGoTo{D: _baf.Get("\u0044")}, nil
}

// Optimizer is the interface that performs optimization of PDF object structure for output writing.
//
// Optimize receives a slice of input `objects`, performs optimization, including removing, replacing objects and
// output the optimized slice of objects.
type Optimizer interface {
	Optimize(_egffgc []_dg.PdfObject) ([]_dg.PdfObject, error)
}

// ToPdfObject returns the text field dictionary within an indirect object (container).
func (_gede *PdfFieldText) ToPdfObject() _dg.PdfObject {
	_gede.PdfField.ToPdfObject()
	_dgdb := _gede._egce
	_deba := _dgdb.PdfObject.(*_dg.PdfObjectDictionary)
	_deba.Set("\u0046\u0054", _dg.MakeName("\u0054\u0078"))
	if _gede.DA != nil {
		_deba.Set("\u0044\u0041", _gede.DA)
	}
	if _gede.Q != nil {
		_deba.Set("\u0051", _gede.Q)
	}
	if _gede.DS != nil {
		_deba.Set("\u0044\u0053", _gede.DS)
	}
	if _gede.RV != nil {
		_deba.Set("\u0052\u0056", _gede.RV)
	}
	if _gede.MaxLen != nil {
		_deba.Set("\u004d\u0061\u0078\u004c\u0065\u006e", _gede.MaxLen)
	}
	return _dgdb
}
func (_agaga *pdfFontType0) bytesToCharcodes(_daec []byte) ([]_bd.CharCode, bool) {
	if _agaga._bfae == nil {
		return nil, false
	}
	_ddgad, _acbd := _agaga._bfae.BytesToCharcodes(_daec)
	if !_acbd {
		return nil, false
	}
	_bggc := make([]_bd.CharCode, len(_ddgad))
	for _gccb, _eegg := range _ddgad {
		_bggc[_gccb] = _bd.CharCode(_eegg)
	}
	return _bggc, true
}

// GetCerts returns the signature certificate chain.
func (_bgba *PdfSignature) GetCerts() ([]*_bb.Certificate, error) {
	var _cdcab []func() ([]*_bb.Certificate, error)
	switch _fbedg, _ := _dg.GetNameVal(_bgba.SubFilter); _fbedg {
	case "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064", "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064":
		_cdcab = append(_cdcab, _bgba.extractChainFromPKCS7, _bgba.extractChainFromCert)
	case "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031":
		_cdcab = append(_cdcab, _bgba.extractChainFromCert)
	case "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031":
		_cdcab = append(_cdcab, _bgba.extractChainFromPKCS7)
	default:
		return nil, _b.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020S\u0075b\u0046i\u006c\u0074\u0065\u0072\u003a\u0020\u0025s", _fbedg)
	}
	for _, _efad := range _cdcab {
		_agece, _ggfda := _efad()
		if _ggfda != nil {
			return nil, _ggfda
		}
		if len(_agece) > 0 {
			return _agece, nil
		}
	}
	return nil, ErrSignNoCertificates
}

// Field returns the parent form field of the widget annotation, if one exists.
// NOTE: the method returns nil if the parent form field has not been parsed.
func (_cfca *PdfAnnotationWidget) Field() *PdfField { return _cfca._cgg }

// PdfAnnotation3D represents 3D annotations.
// (Section 13.6.2).
type PdfAnnotation3D struct {
	*PdfAnnotation
	T3DD _dg.PdfObject
	T3DV _dg.PdfObject
	T3DA _dg.PdfObject
	T3DI _dg.PdfObject
	T3DB _dg.PdfObject
}

func (_ebfgd *PdfWriter) writeString(_eadfg string) {
	if _ebfgd._ffefc != nil {
		return
	}
	_accaf, _gacab := _ebfgd._bddfa.WriteString(_eadfg)
	_ebfgd._fbbfc += int64(_accaf)
	_ebfgd._ffefc = _gacab
}

// NewPdfFieldSignature returns an initialized signature field.
func NewPdfFieldSignature(signature *PdfSignature) *PdfFieldSignature {
	_ffgef := &PdfFieldSignature{}
	_ffgef.PdfField = NewPdfField()
	_ffgef.PdfField.SetContext(_ffgef)
	_ffgef.PdfAnnotationWidget = NewPdfAnnotationWidget()
	_ffgef.PdfAnnotationWidget.SetContext(_ffgef)
	_ffgef.PdfAnnotationWidget._cdf = _ffgef.PdfField._egce
	_ffgef.T = _dg.MakeString("")
	_ffgef.F = _dg.MakeInteger(132)
	_ffgef.V = signature
	return _ffgef
}

// GetCatalogStructTreeRoot gets the catalog StructTreeRoot object.
func (_ffgf *PdfReader) GetCatalogStructTreeRoot() (_dg.PdfObject, bool) {
	if _ffgf._gccfb == nil {
		return nil, false
	}
	_gfbbe := _ffgf._gccfb.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074")
	return _gfbbe, _gfbbe != nil
}

// GetOutlines returns a high-level Outline object, based on the outline tree
// of the reader.
func (_febfg *PdfReader) GetOutlines() (*Outline, error) {
	if _febfg == nil {
		return nil, _bf.New("\u0063\u0061n\u006e\u006f\u0074\u0020c\u0072\u0065a\u0074\u0065\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u006e\u0069\u006c\u0020\u0072e\u0061\u0064\u0065\u0072")
	}
	_accge := _febfg.GetOutlineTree()
	if _accge == nil {
		return nil, _bf.New("\u0074\u0068\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u0064\u006f\u0065\u0073\u0020n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065")
	}
	var _fcee func(_abeff *PdfOutlineTreeNode, _eccca *[]*OutlineItem)
	_fcee = func(_fcedc *PdfOutlineTreeNode, _cfggb *[]*OutlineItem) {
		if _fcedc == nil {
			return
		}
		if _fcedc._baddf == nil {
			_ag.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020m\u0069\u0073\u0073\u0069ng \u006fut\u006c\u0069\u006e\u0065\u0020\u0065\u006etr\u0079\u0020\u0063\u006f\u006e\u0074\u0065x\u0074")
			return
		}
		var _aegcd *OutlineItem
		if _fgdfab, _edgba := _fcedc._baddf.(*PdfOutlineItem); _edgba {
			_gdgcgc := _fgdfab.Dest
			if (_gdgcgc == nil || _dg.IsNullObject(_gdgcgc)) && _fgdfab.A != nil {
				if _facd, _dcacc := _dg.GetDict(_fgdfab.A); _dcacc {
					if _dfdegg, _adfae := _dg.GetArray(_facd.Get("\u0044")); _adfae {
						_gdgcgc = _dfdegg
					} else {
						_gafab, _baebd := _dg.GetString(_facd.Get("\u0044"))
						if !_baebd {
							return
						}
						_daef, _baebd := _febfg._gccfb.Get("\u004e\u0061\u006de\u0073").(*_dg.PdfObjectReference)
						if !_baebd {
							return
						}
						_gfgeg, _cbab := _febfg._baad.LookupByReference(*_daef)
						if _cbab != nil {
							_ag.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0028\u0025\u0073\u0029", _cbab.Error())
							return
						}
						_daege, _baebd := _gfgeg.(*_dg.PdfIndirectObject)
						if !_baebd {
							return
						}
						_gabd := map[_dg.PdfObject]struct{}{}
						_cbab = _febfg.buildNameNodes(_daege, _gabd)
						if _cbab != nil {
							_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006ea\u006d\u0065\u0020\u006e\u006fd\u0065\u0073 \u0028\u0025\u0073\u0029", _cbab.Error())
							return
						}
						for _bfed := range _gabd {
							_daegeg, _adae := _dg.GetDict(_bfed)
							if !_adae {
								continue
							}
							_bgcdd, _adae := _dg.GetArray(_daegeg.Get("\u004e\u0061\u006de\u0073"))
							if !_adae {
								continue
							}
							for _ebbe, _aeca := range _bgcdd.Elements() {
								switch _aeca.(type) {
								case *_dg.PdfObjectString:
									if _aeca.String() == _gafab.String() {
										if _efcd := _bgcdd.Get(_ebbe + 1); _efcd != nil {
											if _baff, _efaeg := _dg.GetDict(_efcd); _efaeg {
												_gdgcgc = _baff.Get("\u0044")
												break
											}
										}
									}
								}
							}
						}
					}
				}
			}
			var _bddd OutlineDest
			if _gdgcgc != nil && !_dg.IsNullObject(_gdgcgc) {
				if _badbg, _dfcdg := _eadf(_gdgcgc, _febfg); _dfcdg == nil {
					_bddd = *_badbg
				} else {
					_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020p\u0061\u0072\u0073\u0065\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065\u0073\u0074\u0020\u0028\u0025\u0076\u0029\u003a\u0020\u0025\u0076", _gdgcgc, _dfcdg)
				}
			}
			_aegcd = NewOutlineItem(_fgdfab.Title.Decoded(), _bddd)
			*_cfggb = append(*_cfggb, _aegcd)
			if _fgdfab.Next != nil {
				_fcee(_fgdfab.Next, _cfggb)
			}
		}
		if _fcedc.First != nil {
			if _aegcd != nil {
				_cfggb = &_aegcd.Entries
			}
			_fcee(_fcedc.First, _cfggb)
		}
	}
	_gddeg := NewOutline()
	_fcee(_accge, &_gddeg.Entries)
	return _gddeg, nil
}
func (_dabc *PdfReader) newPdfAnnotationPopupFromDict(_gdcc *_dg.PdfObjectDictionary) (*PdfAnnotationPopup, error) {
	_geb := PdfAnnotationPopup{}
	_geb.Parent = _gdcc.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	_geb.Open = _gdcc.Get("\u004f\u0070\u0065\u006e")
	return &_geb, nil
}

// GetMediaBox gets the inheritable media box value, either from the page
// or a higher up page/pages struct.
func (_ebfga *PdfPage) GetMediaBox() (*PdfRectangle, error) {
	if _ebfga.MediaBox != nil {
		return _ebfga.MediaBox, nil
	}
	_dccad := _ebfga.Parent
	for _dccad != nil {
		_ecebe, _cdccd := _dg.GetDict(_dccad)
		if !_cdccd {
			return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _cbfbb := _ecebe.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _cbfbb != nil {
			_acdba, _bfaac := _dg.GetArray(_cbfbb)
			if !_bfaac {
				return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0064\u0069a\u0020\u0062\u006f\u0078")
			}
			_aefbf, _fccc := NewPdfRectangle(*_acdba)
			if _fccc != nil {
				return nil, _fccc
			}
			return _aefbf, nil
		}
		_dccad = _ecebe.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, _bf.New("m\u0065\u0064\u0069\u0061 b\u006fx\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
}

// B returns the value of the blue component of the color.
func (_edfe *PdfColorDeviceRGB) B() float64 { return _edfe[2] }
func _dccea(_dcagbd *PdfPage) {
	_aecbb := _dg.PdfObjectName("\u0055\u0046\u0031")
	if !_dcagbd.Resources.HasFontByName(_aecbb) {
		_dcagbd.Resources.SetFontByName(_aecbb, DefaultFont().ToPdfObject())
	}
	var _agadf []string
	_agadf = append(_agadf, "\u0071")
	_agadf = append(_agadf, "\u0042\u0054")
	_agadf = append(_agadf, _b.Sprintf("\u002f%\u0073\u0020\u0031\u0034\u0020\u0054f", _aecbb.String()))
	_agadf = append(_agadf, "\u0031\u0020\u0030\u0020\u0030\u0020\u0072\u0067")
	_agadf = append(_agadf, "\u0031\u0030\u0020\u0031\u0030\u0020\u0054\u0064")
	_eegfcg := "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064\u0020\u0055\u006e\u0069\u0044o\u0063\u0020\u002d\u0020\u0047\u0065\u0074\u0020\u0061\u0020\u006c\u0069\u0063e\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073:/\u002f\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f"
	_agadf = append(_agadf, _b.Sprintf("\u0028%\u0073\u0029\u0020\u0054\u006a", _eegfcg))
	_agadf = append(_agadf, "\u0045\u0054")
	_agadf = append(_agadf, "\u0051")
	_fcaeb := _ga.Join(_agadf, "\u000a")
	_dcagbd.AddContentStreamByString(_fcaeb)
	_dcagbd.ToPdfObject()
}

// NewPdfActionSound returns a new "sound" action.
func NewPdfActionSound() *PdfActionSound {
	_gcb := NewPdfAction()
	_ac := &PdfActionSound{}
	_ac.PdfAction = _gcb
	_gcb.SetContext(_ac)
	return _ac
}

// SetContentStreams sets the content streams based on a string array. Will make
// 1 object stream for each string and reference from the page Contents.
// Each stream will be encoded using the encoding specified by the StreamEncoder,
// if empty, will use identity encoding (raw data).
func (_bgabg *PdfPage) SetContentStreams(cStreams []string, encoder _dg.StreamEncoder) error {
	if len(cStreams) == 0 {
		_bgabg.Contents = nil
		return nil
	}
	if encoder == nil {
		encoder = _dg.NewRawEncoder()
	}
	var _bggge []*_dg.PdfObjectStream
	for _, _fgeea := range cStreams {
		_abbba := &_dg.PdfObjectStream{}
		_fbdec := encoder.MakeStreamDict()
		_acbbc, _aeecb := encoder.EncodeBytes([]byte(_fgeea))
		if _aeecb != nil {
			return _aeecb
		}
		_fbdec.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dg.MakeInteger(int64(len(_acbbc))))
		_abbba.PdfObjectDictionary = _fbdec
		_abbba.Stream = _acbbc
		_bggge = append(_bggge, _abbba)
	}
	if len(_bggge) == 1 {
		_bgabg.Contents = _bggge[0]
	} else {
		_aacfg := _dg.MakeArray()
		for _, _aaad := range _bggge {
			_aacfg.Append(_aaad)
		}
		_bgabg.Contents = _aacfg
	}
	return nil
}

// GetColorspaces loads PdfPageResourcesColorspaces from `r.ColorSpace` and returns an error if there
// is a problem loading. Once loaded, the same object is returned on multiple calls.
func (_cfcdb *PdfPageResources) GetColorspaces() (*PdfPageResourcesColorspaces, error) {
	if _cfcdb._dadae != nil {
		return _cfcdb._dadae, nil
	}
	if _cfcdb.ColorSpace == nil {
		return nil, nil
	}
	_bcec, _fcaea := _ababg(_cfcdb.ColorSpace)
	if _fcaea != nil {
		return nil, _fcaea
	}
	_cfcdb._dadae = _bcec
	return _cfcdb._dadae, nil
}

// SetColorspaceByName adds the provided colorspace to the page resources.
func (_fdacg *PdfPageResources) SetColorspaceByName(keyName _dg.PdfObjectName, cs PdfColorspace) error {
	_gebgb, _dbfdc := _fdacg.GetColorspaces()
	if _dbfdc != nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _dbfdc)
		return _dbfdc
	}
	if _gebgb == nil {
		_gebgb = NewPdfPageResourcesColorspaces()
		_fdacg.SetColorSpace(_gebgb)
	}
	_gebgb.Set(keyName, cs)
	return nil
}
func (_cfag *PdfReader) newPdfActionGoTo3DViewFromDict(_beed *_dg.PdfObjectDictionary) (*PdfActionGoTo3DView, error) {
	return &PdfActionGoTo3DView{TA: _beed.Get("\u0054\u0041"), V: _beed.Get("\u0056")}, nil
}

// BorderStyle defines border type, typically used for annotations.
type BorderStyle int

func (_adfcb *LTV) getOCSPs(_bcca []*_bb.Certificate, _gdafa map[string]*_bb.Certificate) ([][]byte, error) {
	_debaa := make([][]byte, 0, len(_bcca))
	for _, _fgde := range _bcca {
		for _, _febbd := range _fgde.OCSPServer {
			if _adfcb.CertClient.IsCA(_fgde) {
				continue
			}
			_abafb, _bcgfc := _gdafa[_fgde.Issuer.CommonName]
			if !_bcgfc {
				_ag.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _fgcae, _ededa := _adfcb.OCSPClient.MakeRequest(_febbd, _fgde, _abafb)
			if _ededa != nil {
				_ag.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _ededa)
				continue
			}
			_debaa = append(_debaa, _fgcae)
		}
	}
	return _debaa, nil
}

// ToPdfObject returns the PDF representation of the page resources.
func (_baba *PdfPageResources) ToPdfObject() _dg.PdfObject {
	_bgee := _baba._ddcfb
	_bgee.SetIfNotNil("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _baba.ExtGState)
	if _baba._dadae != nil {
		_baba.ColorSpace = _baba._dadae.ToPdfObject()
	}
	_bgee.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _baba.ColorSpace)
	_bgee.SetIfNotNil("\u0050a\u0074\u0074\u0065\u0072\u006e", _baba.Pattern)
	_bgee.SetIfNotNil("\u0053h\u0061\u0064\u0069\u006e\u0067", _baba.Shading)
	_bgee.SetIfNotNil("\u0058O\u0062\u006a\u0065\u0063\u0074", _baba.XObject)
	_bgee.SetIfNotNil("\u0046\u006f\u006e\u0074", _baba.Font)
	_bgee.SetIfNotNil("\u0050r\u006f\u0063\u0053\u0065\u0074", _baba.ProcSet)
	_bgee.SetIfNotNil("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", _baba.Properties)
	return _bgee
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_bbdf *PdfColorspaceSpecialIndexed) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	N := _bbdf.Base.GetNumComponents()
	_dacdf := int(vals[0]) * N
	if _dacdf < 0 || (_dacdf+N-1) >= len(_bbdf._bcea) {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _dacdf)
		return nil, ErrColorOutOfRange
	}
	_gbfg := _bbdf._bcea[_dacdf : _dacdf+N]
	var _geae []float64
	for _, _ffcc := range _gbfg {
		_geae = append(_geae, float64(_ffcc)/255.0)
	}
	_agfac, _bebe := _bbdf.Base.ColorFromFloats(_geae)
	if _bebe != nil {
		return nil, _bebe
	}
	return _agfac, nil
}

// PdfShadingPattern is a Shading patterns that provide a smooth transition between colors across an area to be painted,
// i.e. color(x,y) = f(x,y) at each point.
// It is a type 2 pattern (PatternType = 2).
type PdfShadingPattern struct {
	*PdfPattern
	Shading   *PdfShading
	Matrix    *_dg.PdfObjectArray
	ExtGState _dg.PdfObject
}

// A returns the value of the A component of the color.
func (_dfad *PdfColorLab) A() float64 { return _dfad[1] }

// GetFillImage get attached model.Image in push button.
func (_gdbf *PdfFieldButton) GetFillImage() *Image {
	if _gdbf.IsPush() {
		return _gdbf._gddc
	}
	return nil
}

// IsSimple returns true if `font` is a simple font.
func (_agba *PdfFont) IsSimple() bool { _, _edcdda := _agba._cadf.(*pdfFontSimple); return _edcdda }

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// component PDF objects.
func (_aeege *PdfColorspaceICCBased) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if _aeege.Alternate == nil {
		if _aeege.N == 1 {
			_feec := NewPdfColorspaceDeviceGray()
			return _feec.ColorFromPdfObjects(objects)
		} else if _aeege.N == 3 {
			_dea := NewPdfColorspaceDeviceRGB()
			return _dea.ColorFromPdfObjects(objects)
		} else if _aeege.N == 4 {
			_dage := NewPdfColorspaceDeviceCMYK()
			return _dage.ColorFromPdfObjects(objects)
		} else {
			return nil, _bf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _aeege.Alternate.ColorFromPdfObjects(objects)
}
func _afbc(_ecdefc *_dg.PdfObjectDictionary) (*PdfTilingPattern, error) {
	_ecbgf := &PdfTilingPattern{}
	_fead := _ecdefc.Get("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e")
	if _fead == nil {
		_ag.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054\u0079\u0070\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bdgaab, _gdaeg := _fead.(*_dg.PdfObjectInteger)
	if !_gdaeg {
		_ag.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054y\u0070\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _fead)
		return nil, _dg.ErrTypeError
	}
	_ecbgf.PaintType = _bdgaab
	_fead = _ecdefc.Get("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _fead == nil {
		_ag.Log.Debug("\u0054i\u006ci\u006e\u0067\u0054\u0079\u0070e\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fdfc, _gdaeg := _fead.(*_dg.PdfObjectInteger)
	if !_gdaeg {
		_ag.Log.Debug("\u0054\u0069\u006cin\u0067\u0054\u0079\u0070\u0065\u0020\u006e\u006f\u0074 \u0061n\u0020i\u006et\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _fead)
		return nil, _dg.ErrTypeError
	}
	_ecbgf.TilingType = _fdfc
	_fead = _ecdefc.Get("\u0042\u0042\u006f\u0078")
	if _fead == nil {
		_ag.Log.Debug("\u0042\u0042\u006fx\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fead = _dg.TraceToDirectObject(_fead)
	_bbbec, _gdaeg := _fead.(*_dg.PdfObjectArray)
	if !_gdaeg {
		_ag.Log.Debug("\u0042B\u006f\u0078 \u0073\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061y\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _fead)
		return nil, _dg.ErrTypeError
	}
	_ccbcf, _bdagca := NewPdfRectangle(*_bbbec)
	if _bdagca != nil {
		_ag.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _bdagca)
		return nil, _bdagca
	}
	_ecbgf.BBox = _ccbcf
	_fead = _ecdefc.Get("\u0058\u0053\u0074e\u0070")
	if _fead == nil {
		_ag.Log.Debug("\u0058\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fddaae, _bdagca := _dg.GetNumberAsFloat(_fead)
	if _bdagca != nil {
		_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0058S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _fddaae)
		return nil, _bdagca
	}
	_ecbgf.XStep = _dg.MakeFloat(_fddaae)
	_fead = _ecdefc.Get("\u0059\u0053\u0074e\u0070")
	if _fead == nil {
		_ag.Log.Debug("\u0059\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bdebd, _bdagca := _dg.GetNumberAsFloat(_fead)
	if _bdagca != nil {
		_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0059S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _bdebd)
		return nil, _bdagca
	}
	_ecbgf.YStep = _dg.MakeFloat(_bdebd)
	_fead = _ecdefc.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _fead == nil {
		_ag.Log.Debug("\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_ecdefc, _gdaeg = _dg.TraceToDirectObject(_fead).(*_dg.PdfObjectDictionary)
	if !_gdaeg {
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _fead)
	}
	_bedf, _bdagca := NewPdfPageResourcesFromDict(_ecdefc)
	if _bdagca != nil {
		return nil, _bdagca
	}
	_ecbgf.Resources = _bedf
	if _ddadb := _ecdefc.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _ddadb != nil {
		_dfed, _bdeedg := _ddadb.(*_dg.PdfObjectArray)
		if !_bdeedg {
			_ag.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ddadb)
			return nil, _dg.ErrTypeError
		}
		_ecbgf.Matrix = _dfed
	}
	return _ecbgf, nil
}

// Image interface is a basic representation of an image used in PDF.
// The colorspace is not specified, but must be known when handling the image.
type Image struct {
	Width            int64
	Height           int64
	BitsPerComponent int64
	ColorComponents  int
	Data             []byte
	_dgeb            []byte
	_gfbb            []float64
}

func (_cede fontCommon) coreString() string {
	_abede := ""
	if _cede._ccfb != nil {
		_abede = _cede._ccfb.String()
	}
	return _b.Sprintf("\u0025#\u0071\u0020%\u0023\u0071\u0020%\u0071\u0020\u006f\u0062\u006a\u003d\u0025d\u0020\u0054\u006f\u0055\u006e\u0069c\u006f\u0064\u0065\u003d\u0025\u0074\u0020\u0066\u006c\u0061\u0067s\u003d\u0030\u0078\u0025\u0030\u0078\u0020\u0025\u0073", _cede._bcga, _cede._ecbf, _cede._cefg, _cede._bgggd, _cede._ebbff != nil, _cede.fontFlags(), _abede)
}

// ToPdfObject implements interface PdfModel.
func (_cba *PdfActionTrans) ToPdfObject() _dg.PdfObject {
	_cba.PdfAction.ToPdfObject()
	_cbg := _cba._cbd
	_cd := _cbg.PdfObject.(*_dg.PdfObjectDictionary)
	_cd.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeTrans)))
	_cd.SetIfNotNil("\u0054\u0072\u0061n\u0073", _cba.Trans)
	return _cbg
}
func (_aae *PdfReader) newPdfActionHideFromDict(_fcbd *_dg.PdfObjectDictionary) (*PdfActionHide, error) {
	return &PdfActionHide{T: _fcbd.Get("\u0054"), H: _fcbd.Get("\u0048")}, nil
}
func (_abfa *pdfCIDFontType0) getFontDescriptor() *PdfFontDescriptor { return _abfa._ccfb }

// NewPdfPageResourcesColorspaces returns a new PdfPageResourcesColorspaces object.
func NewPdfPageResourcesColorspaces() *PdfPageResourcesColorspaces {
	_aadeg := &PdfPageResourcesColorspaces{}
	_aadeg.Names = []string{}
	_aadeg.Colorspaces = map[string]PdfColorspace{}
	_aadeg._dcebf = &_dg.PdfIndirectObject{}
	return _aadeg
}

// GetShadingByName gets the shading specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_dged *PdfPageResources) GetShadingByName(keyName _dg.PdfObjectName) (*PdfShading, bool) {
	if _dged.Shading == nil {
		return nil, false
	}
	_cegad, _gafag := _dg.TraceToDirectObject(_dged.Shading).(*_dg.PdfObjectDictionary)
	if !_gafag {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0068\u0061d\u0069\u006e\u0067\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _dged.Shading)
		return nil, false
	}
	if _bfdac := _cegad.Get(keyName); _bfdac != nil {
		_dece, _fbgbef := _bgdgf(_bfdac)
		if _fbgbef != nil {
			_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0073\u0068\u0061\u0064\u0069\u006eg\u003a\u0020\u0025\u0076", _fbgbef)
			return nil, false
		}
		return _dece, true
	}
	return nil, false
}

// PdfAnnotationPolygon represents Polygon annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolygon struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _dg.PdfObject
	LE       _dg.PdfObject
	BS       _dg.PdfObject
	IC       _dg.PdfObject
	BE       _dg.PdfObject
	IT       _dg.PdfObject
	Measure  _dg.PdfObject
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_cbbc *PdfOutlineItem) ToPdfObject() _dg.PdfObject {
	_bcgd := _cbbc._fbgf
	_bgga := _bcgd.PdfObject.(*_dg.PdfObjectDictionary)
	_bgga.Set("\u0054\u0069\u0074l\u0065", _cbbc.Title)
	if _cbbc.A != nil {
		_bgga.Set("\u0041", _cbbc.A)
	}
	if _efaf := _bgga.Get("\u0053\u0045"); _efaf != nil {
		_bgga.Remove("\u0053\u0045")
	}
	if _cbbc.C != nil {
		_bgga.Set("\u0043", _cbbc.C)
	}
	if _cbbc.Dest != nil {
		_bgga.Set("\u0044\u0065\u0073\u0074", _cbbc.Dest)
	}
	if _cbbc.F != nil {
		_bgga.Set("\u0046", _cbbc.F)
	}
	if _cbbc.Count != nil {
		_bgga.Set("\u0043\u006f\u0075n\u0074", _dg.MakeInteger(*_cbbc.Count))
	}
	if _cbbc.Next != nil {
		_bgga.Set("\u004e\u0065\u0078\u0074", _cbbc.Next.ToPdfObject())
	}
	if _cbbc.First != nil {
		_bgga.Set("\u0046\u0069\u0072s\u0074", _cbbc.First.ToPdfObject())
	}
	if _cbbc.Prev != nil {
		_bgga.Set("\u0050\u0072\u0065\u0076", _cbbc.Prev.GetContext().GetContainingPdfObject())
	}
	if _cbbc.Last != nil {
		_bgga.Set("\u004c\u0061\u0073\u0074", _cbbc.Last.GetContext().GetContainingPdfObject())
	}
	if _cbbc.Parent != nil {
		_bgga.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _cbbc.Parent.GetContext().GetContainingPdfObject())
	}
	return _bcgd
}

// PdfFunctionType0 uses a sequence of sample values (contained in a stream) to provide an approximation
// for functions whose domains and ranges are bounded. The samples are organized as an m-dimensional
// table in which each entry has n components
type PdfFunctionType0 struct {
	Domain        []float64
	Range         []float64
	NumInputs     int
	NumOutputs    int
	Size          []int
	BitsPerSample int
	Order         int
	Encode        []float64
	Decode        []float64
	_bgdg         []byte
	_fecfc        []uint32
	_bedaf        *_dg.PdfObjectStream
}

// SetXObjectImageByName adds the provided XObjectImage to the page resources.
// The added XObjectImage is identified by the specified name.
func (_bacaa *PdfPageResources) SetXObjectImageByName(keyName _dg.PdfObjectName, ximg *XObjectImage) error {
	_dbcgg := ximg.ToPdfObject().(*_dg.PdfObjectStream)
	_ebbeg := _bacaa.SetXObjectByName(keyName, _dbcgg)
	return _ebbeg
}

var _ pdfFont = (*pdfCIDFontType0)(nil)

// ToPdfObject returns the PDF representation of the function.
func (_dfbb *PdfFunctionType0) ToPdfObject() _dg.PdfObject {
	if _dfbb._bedaf == nil {
		_dfbb._bedaf = &_dg.PdfObjectStream{}
	}
	_eacc := _dg.MakeDict()
	_eacc.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _dg.MakeInteger(0))
	_gbfa := &_dg.PdfObjectArray{}
	for _, _gcac := range _dfbb.Domain {
		_gbfa.Append(_dg.MakeFloat(_gcac))
	}
	_eacc.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _gbfa)
	_addbc := &_dg.PdfObjectArray{}
	for _, _dggdb := range _dfbb.Range {
		_addbc.Append(_dg.MakeFloat(_dggdb))
	}
	_eacc.Set("\u0052\u0061\u006eg\u0065", _addbc)
	_gbfgd := &_dg.PdfObjectArray{}
	for _, _agdcf := range _dfbb.Size {
		_gbfgd.Append(_dg.MakeInteger(int64(_agdcf)))
	}
	_eacc.Set("\u0053\u0069\u007a\u0065", _gbfgd)
	_eacc.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065", _dg.MakeInteger(int64(_dfbb.BitsPerSample)))
	if _dfbb.Order != 1 {
		_eacc.Set("\u004f\u0072\u0064e\u0072", _dg.MakeInteger(int64(_dfbb.Order)))
	}
	_eacc.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dg.MakeInteger(int64(len(_dfbb._bgdg))))
	_dfbb._bedaf.Stream = _dfbb._bgdg
	_dfbb._bedaf.PdfObjectDictionary = _eacc
	return _dfbb._bedaf
}

// ColorToRGB converts a color in Separation colorspace to RGB colorspace.
func (_aaed *PdfColorspaceSpecialSeparation) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _aaed.AlternateSpace == nil {
		return nil, _bf.New("\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020c\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	return _aaed.AlternateSpace.ColorToRGB(color)
}
func (_cebeg *PdfReader) loadPerms() (*Permissions, error) {
	if _adbed := _cebeg._gccfb.Get("\u0050\u0065\u0072m\u0073"); _adbed != nil {
		if _cdgdff, _bbga := _dg.GetDict(_adbed); _bbga {
			_ecaff := _cdgdff.Get("\u0044\u006f\u0063\u004d\u0044\u0050")
			if _ecaff == nil {
				return nil, nil
			}
			if _afgfg, _fbeca := _dg.GetIndirect(_ecaff); _fbeca {
				_aadea, _cfbd := _cebeg.newPdfSignatureFromIndirect(_afgfg)
				if _cfbd != nil {
					return nil, _cfbd
				}
				return NewPermissions(_aadea), nil
			}
			return nil, _b.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u006f\u0063M\u0044\u0050\u0020\u0065nt\u0072\u0079")
		}
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0050\u0065\u0072\u006d\u0073\u0020\u0065\u006e\u0074\u0072\u0079")
	}
	return nil, nil
}
func (_beafe *PdfReader) loadForms() (*PdfAcroForm, error) {
	if _beafe._baad.GetCrypter() != nil && !_beafe._baad.IsAuthenticated() {
		return nil, _b.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_afacg := _beafe._gccfb
	_cdge := _afacg.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")
	if _cdge == nil {
		return nil, nil
	}
	_feaf, _eecd := _dg.GetIndirect(_cdge)
	_cdge = _dg.TraceToDirectObject(_cdge)
	if _dg.IsNullObject(_cdge) {
		_ag.Log.Trace("\u0041\u0063\u0072of\u006f\u0072\u006d\u0020\u0069\u0073\u0020\u0061\u0020n\u0075l\u006c \u006fb\u006a\u0065\u0063\u0074\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029\u000a")
		return nil, nil
	}
	_bgadf, _agge := _dg.GetDict(_cdge)
	if !_agge {
		_ag.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0041\u0063\u0072\u006fF\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _cdge)
		_ag.Log.Debug("\u0044\u006f\u0065\u0073 n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0073")
		return nil, _b.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0061\u0063\u0072\u006ff\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _cdge)
	}
	_ag.Log.Trace("\u0048\u0061\u0073\u0020\u0041\u0063\u0072\u006f\u0020f\u006f\u0072\u006d\u0073")
	_ag.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0041\u0063r\u006ff\u006f\u0072\u006d\u0073\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065")
	if !_beafe._dadcef {
		_ecaa := _beafe.traverseObjectData(_bgadf)
		if _ecaa != nil {
			_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0074\u0072a\u0076\u0065\u0072\u0073\u0065\u0020\u0041\u0063\u0072\u006fFo\u0072\u006d\u0073 \u0028%\u0073\u0029", _ecaa)
			return nil, _ecaa
		}
	}
	_fbgfb, _cedca := _beafe.newPdfAcroFormFromDict(_feaf, _bgadf)
	if _cedca != nil {
		return nil, _cedca
	}
	_fbgfb._eeecb = !_eecd
	return _fbgfb, nil
}

// ColorToRGB converts a ICCBased color to an RGB color.
func (_fdfg *PdfColorspaceICCBased) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _fdfg.Alternate == nil {
		_ag.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _fdfg.N == 1 {
			_ag.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_cgbb := NewPdfColorspaceDeviceGray()
			return _cgbb.ColorToRGB(color)
		} else if _fdfg.N == 3 {
			_ag.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return color, nil
		} else if _fdfg.N == 4 {
			_ag.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_gdff := NewPdfColorspaceDeviceCMYK()
			return _gdff.ColorToRGB(color)
		} else {
			return nil, _bf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_ag.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _fdfg)
	return _fdfg.Alternate.ColorToRGB(color)
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_gabg pdfFontType0) GetRuneMetrics(r rune) (_bbg.CharMetrics, bool) {
	if _gabg.DescendantFont == nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _gabg)
		return _bbg.CharMetrics{}, false
	}
	return _gabg.DescendantFont.GetRuneMetrics(r)
}
func (_ebdgg *PdfWriter) seekByName(_cdgbbf _dg.PdfObject, _dfaec []string, _bcaab string) ([]_dg.PdfObject, error) {
	_ag.Log.Trace("\u0053\u0065\u0065\u006b\u0020\u0062\u0079\u0020\u006e\u0061\u006d\u0065.\u002e\u0020\u0025\u0054", _cdgbbf)
	var _eada []_dg.PdfObject
	if _faeed, _daecea := _cdgbbf.(*_dg.PdfIndirectObject); _daecea {
		return _ebdgg.seekByName(_faeed.PdfObject, _dfaec, _bcaab)
	}
	if _gfgba, _bgbbd := _cdgbbf.(*_dg.PdfObjectStream); _bgbbd {
		return _ebdgg.seekByName(_gfgba.PdfObjectDictionary, _dfaec, _bcaab)
	}
	if _gaaae, _degg := _cdgbbf.(*_dg.PdfObjectDictionary); _degg {
		_ag.Log.Trace("\u0044\u0069\u0063\u0074")
		for _, _deaag := range _gaaae.Keys() {
			_cfaff := _gaaae.Get(_deaag)
			if string(_deaag) == _bcaab {
				_eada = append(_eada, _cfaff)
			}
			for _, _bddcd := range _dfaec {
				if string(_deaag) == _bddcd {
					_ag.Log.Trace("\u0046\u006f\u006c\u006c\u006f\u0077\u0020\u006b\u0065\u0079\u0020\u0025\u0073", _bddcd)
					_ddcdf, _agaee := _ebdgg.seekByName(_cfaff, _dfaec, _bcaab)
					if _agaee != nil {
						return _eada, _agaee
					}
					_eada = append(_eada, _ddcdf...)
					break
				}
			}
		}
		return _eada, nil
	}
	return _eada, nil
}

// Encrypt encrypts the output file with a specified user/owner password.
func (_agadd *PdfWriter) Encrypt(userPass, ownerPass []byte, options *EncryptOptions) error {
	_dcdeg := RC4_128bit
	if options != nil {
		_dcdeg = options.Algorithm
	}
	_ceagc := _gbd.PermOwner
	if options != nil {
		_ceagc = options.Permissions
	}
	var _gcaba _afb.Filter
	switch _dcdeg {
	case RC4_128bit:
		_gcaba = _afb.NewFilterV2(16)
	case AES_128bit:
		_gcaba = _afb.NewFilterAESV2()
	case AES_256bit:
		_gcaba = _afb.NewFilterAESV3()
	default:
		return _b.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020a\u006cg\u006fr\u0069\u0074\u0068\u006d\u003a\u0020\u0025v", options.Algorithm)
	}
	_deccab, _gdeed, _fdca := _dg.PdfCryptNewEncrypt(_gcaba, userPass, ownerPass, _ceagc)
	if _fdca != nil {
		return _fdca
	}
	_agadd._fadcg = _deccab
	if _gdeed.Major != 0 {
		_agadd.SetVersion(_gdeed.Major, _gdeed.Minor)
	}
	_agadd._fbfbf = _gdeed.Encrypt
	_agadd._cedaf, _agadd._ceaab = _gdeed.ID0, _gdeed.ID1
	_gdbda := _dg.MakeIndirectObject(_gdeed.Encrypt)
	_agadd._acdag = _gdbda
	_agadd.addObject(_gdbda)
	return nil
}

// ColorFromPdfObjects gets the color from a series of pdf objects (4 for cmyk).
func (_baee *PdfColorspaceDeviceCMYK) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 4 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bbbce, _cbcgc := _dg.GetNumbersAsFloat(objects)
	if _cbcgc != nil {
		return nil, _cbcgc
	}
	return _baee.ColorFromFloats(_bbbce)
}

// PdfOutputIntent provides a means for matching the color characteristics of a PDF document with
// those of a target output device.
// Multiple PdfOutputIntents allows the production process to be customized to the expected workflow and the specific
// tools available.
type PdfOutputIntent struct {

	// Type is an optional PDF object that this dictionary describes.
	// If present, must be OutputIntent for an output intent dictionary.
	Type string

	// S defines the OutputIntent subtype which should match the standard used in given document i.e:
	// for PDF/X use PdfOutputIntentTypeX.
	S PdfOutputIntentType

	// OutputCondition is an optional field that is identifying the intended output device or production condition in
	// human-readable form. This is preferred method of defining such a string for presentation to the user.
	OutputCondition string

	// OutputConditionIdentifier is a required field identifying the intended output device or production condition in
	// human or machine-readable form. If human-readable, this string may be used
	// in lieu of an OutputCondition for presentation to the user.
	// A typical value for this entry would be the name of a production condition  maintained
	// in  an  industry-standard registry such  as the ICC Characterization Data Registry
	// If the intended production condition is not a recognized standard, the value Custom is recommended for this entry.
	// the DestOutputProfile entry defines the ICC profile, and the Info entry is used for further
	// human-readable identification.
	OutputConditionIdentifier string

	// RegistryName is an optional string field (conventionally URI) identifying the registry in which the condition
	// designated by OutputConditionIdentifier is defined.
	RegistryName string

	// Info is a required field if OutputConditionIdentifier does not specify a standard production condition.
	// A human-readable text string containing additional information  or comments about intended
	// target device or production condition.
	Info string

	// DestOutputProfile is required if OutputConditionIdentifier does not specify a standard production condition.
	// It is an ICC profile stream defining the transformation from the PDF document's source colors to output device colorants.
	DestOutputProfile []byte

	// ColorComponents is the number of color components supported by given output profile.
	ColorComponents int
	_cfaef          *_dg.PdfObjectDictionary
}

// PdfOutputIntentType is the subtype of the given PdfOutputIntent.
type PdfOutputIntentType int

// DecodeArray returns the component range values for the Separation colorspace.
func (_gcec *PdfColorspaceSpecialSeparation) DecodeArray() []float64 { return []float64{0, 1.0} }

// ToPdfObject converts the pdfCIDFontType2 to a PDF representation.
func (_bcbbd *pdfCIDFontType2) ToPdfObject() _dg.PdfObject {
	if _bcbbd._ddccf == nil {
		_bcbbd._ddccf = &_dg.PdfIndirectObject{}
	}
	_dgae := _bcbbd.baseFields().asPdfObjectDictionary("\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032")
	_bcbbd._ddccf.PdfObject = _dgae
	if _bcbbd.CIDSystemInfo != nil {
		_dgae.Set("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f", _bcbbd.CIDSystemInfo)
	}
	if _bcbbd.DW != nil {
		_dgae.Set("\u0044\u0057", _bcbbd.DW)
	}
	if _bcbbd.DW2 != nil {
		_dgae.Set("\u0044\u0057\u0032", _bcbbd.DW2)
	}
	if _bcbbd.W != nil {
		_dgae.Set("\u0057", _bcbbd.W)
	}
	if _bcbbd.W2 != nil {
		_dgae.Set("\u0057\u0032", _bcbbd.W2)
	}
	if _bcbbd.CIDToGIDMap != nil {
		_dgae.Set("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070", _bcbbd.CIDToGIDMap)
	}
	return _bcbbd._ddccf
}

// GetNumComponents returns the number of color components (4 for CMYK32).
func (_bfcb *PdfColorDeviceCMYK) GetNumComponents() int { return 4 }

// PdfActionSound represents a sound action.
type PdfActionSound struct {
	*PdfAction
	Sound       _dg.PdfObject
	Volume      _dg.PdfObject
	Synchronous _dg.PdfObject
	Repeat      _dg.PdfObject
	Mix         _dg.PdfObject
}

// PdfActionRendition represents a Rendition action.
type PdfActionRendition struct {
	*PdfAction
	R  _dg.PdfObject
	AN _dg.PdfObject
	OP _dg.PdfObject
	JS _dg.PdfObject
}

// GetSubFilter returns SubFilter value or empty string.
func (_faac *pdfSignDictionary) GetSubFilter() string {
	_bdcbbd := _faac.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r")
	if _bdcbbd == nil {
		return ""
	}
	if _bdcdga, _fgbad := _dg.GetNameVal(_bdcbbd); _fgbad {
		return _bdcdga
	}
	return ""
}
func (_ecec *PdfReader) newPdfAnnotationSoundFromDict(_ebcf *_dg.PdfObjectDictionary) (*PdfAnnotationSound, error) {
	_cbcg := PdfAnnotationSound{}
	_badb, _dde := _ecec.newPdfAnnotationMarkupFromDict(_ebcf)
	if _dde != nil {
		return nil, _dde
	}
	_cbcg.PdfAnnotationMarkup = _badb
	_cbcg.Name = _ebcf.Get("\u004e\u0061\u006d\u0065")
	_cbcg.Sound = _ebcf.Get("\u0053\u006f\u0075n\u0064")
	return &_cbcg, nil
}

// GetContentStreams returns the content stream as an array of strings.
func (_cadgb *PdfPage) GetContentStreams() ([]string, error) {
	_afbbe := _cadgb.GetContentStreamObjs()
	var _bgbgd []string
	for _, _baea := range _afbbe {
		_afcce, _edfdg := _gfebd(_baea)
		if _edfdg != nil {
			return nil, _edfdg
		}
		_bgbgd = append(_bgbgd, _afcce)
	}
	return _bgbgd, nil
}

// DecodeArray returns the component range values for the DeviceN colorspace.
// [0 1.0 0 1.0 ...] for each color component.
func (_bbgf *PdfColorspaceDeviceN) DecodeArray() []float64 {
	var _dafgd []float64
	for _cgbg := 0; _cgbg < _bbgf.GetNumComponents(); _cgbg++ {
		_dafgd = append(_dafgd, 0.0, 1.0)
	}
	return _dafgd
}
func _eeafe(_cdcfe _dg.PdfObject) []*_dg.PdfObjectStream {
	if _cdcfe == nil {
		return nil
	}
	_geddc, _cddfg := _dg.GetArray(_cdcfe)
	if !_cddfg || _geddc.Len() == 0 {
		return nil
	}
	_ddbfa := make([]*_dg.PdfObjectStream, 0, _geddc.Len())
	for _, _agacb := range _geddc.Elements() {
		if _abegb, _fggag := _dg.GetStream(_agacb); _fggag {
			_ddbfa = append(_ddbfa, _abegb)
		}
	}
	return _ddbfa
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 4 for a CMYK32 device.
func (_fdbbb *PdfColorspaceDeviceCMYK) GetNumComponents() int { return 4 }

// HasExtGState checks whether a font is defined by the specified keyName.
func (_cafe *PdfPageResources) HasExtGState(keyName _dg.PdfObjectName) bool {
	_, _cdfaa := _cafe.GetFontByName(keyName)
	return _cdfaa
}

// GetContext returns the action context which contains the specific type-dependent context.
// The context represents the subaction.
func (_fef *PdfAction) GetContext() PdfModel {
	if _fef == nil {
		return nil
	}
	return _fef._bg
}

// Normalize swaps (Llx,Urx) if Urx < Llx, and (Lly,Ury) if Ury < Lly.
func (_cfcee *PdfRectangle) Normalize() {
	if _cfcee.Llx > _cfcee.Urx {
		_cfcee.Llx, _cfcee.Urx = _cfcee.Urx, _cfcee.Llx
	}
	if _cfcee.Lly > _cfcee.Ury {
		_cfcee.Lly, _cfcee.Ury = _cfcee.Ury, _cfcee.Lly
	}
}

// LTV represents an LTV (Long-Term Validation) client. It is used to LTV
// enable signatures by adding validation and revocation data (certificate,
// OCSP and CRL information) to the DSS dictionary of a PDF document.
//
// LTV is added through the DSS by:
//   - Adding certificates, OCSP and CRL information in the global scope of the
//     DSS. The global data is used for validating any of the signatures present
//     in the document.
//   - Adding certificates, OCSP and CRL information for a single signature,
//     through an entry in the VRI dictionary of the DSS. The added data is used
//     for validating that particular signature only. This is the recommended
//     method for adding validation data for a signature. However, this is not
//     is not possible in the same revision the signature is applied. Validation
//     data for a signature is added based on the Contents entry of the signature,
//     which is known only after the revision is written. Even if the Contents
//     are known (e.g. when signing externally), updating the DSS at that point
//     would invalidate the calculated signature. As a result, if adding LTV
//     in the same revision is a requirement, use the first method.
//     See LTV.EnableChain.
//
// The client applies both methods, when possible.
//
// If `LTV.SkipExisting` is set to true (the default), validations are
// not added for signatures which are already present in the VRI entry of the
// document's DSS dictionary.
type LTV struct {

	// CertClient is the client used to retrieve certificates.
	CertClient *_fe.CertClient

	// OCSPClient is the client used to retrieve OCSP validation information.
	OCSPClient *_fe.OCSPClient

	// CRLClient is the client used to retrieve CRL validation information.
	CRLClient *_fe.CRLClient

	// SkipExisting specifies whether existing signature validations
	// should be skipped.
	SkipExisting bool
	_ebdgd       *PdfAppender
	_abca        *DSS
}

func (_bbadf *PdfWriter) checkPendingObjects() {
	for _ggccg, _dbbab := range _bbadf._ccgade {
		if !_bbadf.hasObject(_ggccg) {
			_ag.Log.Debug("\u0057\u0041\u0052\u004e\u0020\u0050\u0065n\u0064\u0069\u006eg\u0020\u006f\u0062j\u0065\u0063t\u0020\u0025\u002b\u0076\u0020\u0025T\u0020(%\u0070\u0029\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0077\u0072\u0069\u0074\u0069\u006e\u0067", _ggccg, _ggccg, _ggccg)
			for _, _gbffd := range _dbbab {
				for _, _cgaae := range _gbffd.Keys() {
					_fffef := _gbffd.Get(_cgaae)
					if _fffef == _ggccg {
						_ag.Log.Debug("\u0050e\u006e\u0064i\u006e\u0067\u0020\u006fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0061nd\u0020\u0072\u0065p\u006c\u0061c\u0065\u0064\u0020\u0077\u0069\u0074h\u0020\u006eu\u006c\u006c")
						_gbffd.Set(_cgaae, _dg.MakeNull())
						break
					}
				}
			}
		}
	}
}

// NewPdfActionSubmitForm returns a new "submit form" action.
func NewPdfActionSubmitForm() *PdfActionSubmitForm {
	_da := NewPdfAction()
	_ead := &PdfActionSubmitForm{}
	_ead.PdfAction = _da
	_da.SetContext(_ead)
	return _ead
}

// ToPdfObject implements interface PdfModel.
func (_abab *PdfAnnotationPolyLine) ToPdfObject() _dg.PdfObject {
	_abab.PdfAnnotation.ToPdfObject()
	_fafc := _abab._cdf
	_ccf := _fafc.PdfObject.(*_dg.PdfObjectDictionary)
	_abab.PdfAnnotationMarkup.appendToPdfDictionary(_ccf)
	_ccf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065"))
	_ccf.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _abab.Vertices)
	_ccf.SetIfNotNil("\u004c\u0045", _abab.LE)
	_ccf.SetIfNotNil("\u0042\u0053", _abab.BS)
	_ccf.SetIfNotNil("\u0049\u0043", _abab.IC)
	_ccf.SetIfNotNil("\u0042\u0045", _abab.BE)
	_ccf.SetIfNotNil("\u0049\u0054", _abab.IT)
	_ccf.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _abab.Measure)
	return _fafc
}

const (
	BorderStyleSolid     BorderStyle = iota
	BorderStyleDashed    BorderStyle = iota
	BorderStyleBeveled   BorderStyle = iota
	BorderStyleInset     BorderStyle = iota
	BorderStyleUnderline BorderStyle = iota
)

func _bbadae() string {
	_fgefgf.Lock()
	defer _fgefgf.Unlock()
	return _faff
}

// String returns the name of the colorspace (DeviceN).
func (_caae *PdfColorspaceDeviceN) String() string { return "\u0044e\u0076\u0069\u0063\u0065\u004e" }
func (_daabag *PdfWriter) addObjects(_dcbba _dg.PdfObject) error {
	_ag.Log.Trace("\u0041d\u0064i\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073\u0021")
	if _bgcbc, _ebcbf := _dcbba.(*_dg.PdfIndirectObject); _ebcbf {
		_ag.Log.Trace("\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074")
		_ag.Log.Trace("\u002d \u0025\u0073\u0020\u0028\u0025\u0070)", _dcbba, _bgcbc)
		_ag.Log.Trace("\u002d\u0020\u0025\u0073", _bgcbc.PdfObject)
		if _daabag.addObject(_bgcbc) {
			_cdada := _daabag.addObjects(_bgcbc.PdfObject)
			if _cdada != nil {
				return _cdada
			}
		}
		return nil
	}
	if _ddcff, _afbcb := _dcbba.(*_dg.PdfObjectStream); _afbcb {
		_ag.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d")
		_ag.Log.Trace("\u002d \u0025\u0073\u0020\u0025\u0070", _dcbba, _dcbba)
		if _daabag.addObject(_ddcff) {
			_fgcbdg := _daabag.addObjects(_ddcff.PdfObjectDictionary)
			if _fgcbdg != nil {
				return _fgcbdg
			}
		}
		return nil
	}
	if _acec, _ceabc := _dcbba.(*_dg.PdfObjectDictionary); _ceabc {
		_ag.Log.Trace("\u0044\u0069\u0063\u0074")
		_ag.Log.Trace("\u002d\u0020\u0025\u0073", _dcbba)
		for _, _ecbd := range _acec.Keys() {
			_fcdaf := _acec.Get(_ecbd)
			if _caadg, _caagg := _fcdaf.(*_dg.PdfObjectReference); _caagg {
				_fcdaf = _caadg.Resolve()
				_acec.Set(_ecbd, _fcdaf)
			}
			if _ecbd != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _ddcegf := _daabag.addObjects(_fcdaf); _ddcegf != nil {
					return _ddcegf
				}
			} else {
				if _, _faef := _fcdaf.(*_dg.PdfObjectNull); _faef {
					continue
				}
				if _caceg := _daabag.hasObject(_fcdaf); !_caceg {
					_ag.Log.Debug("P\u0061\u0072\u0065\u006e\u0074\u0020o\u0062\u006a\u0020\u006e\u006f\u0074 \u0061\u0064\u0064\u0065\u0064\u0020\u0079e\u0074\u0021\u0021\u0020\u0025\u0054\u0020\u0025\u0070\u0020%\u0076", _fcdaf, _fcdaf, _fcdaf)
					_daabag._ccgade[_fcdaf] = append(_daabag._ccgade[_fcdaf], _acec)
				}
			}
		}
		return nil
	}
	if _gecbb, _cfbec := _dcbba.(*_dg.PdfObjectArray); _cfbec {
		_ag.Log.Trace("\u0041\u0072\u0072a\u0079")
		_ag.Log.Trace("\u002d\u0020\u0025\u0073", _dcbba)
		if _gecbb == nil {
			return _bf.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _dgbafe, _gbbd := range _gecbb.Elements() {
			if _gcfg, _edfc := _gbbd.(*_dg.PdfObjectReference); _edfc {
				_gbbd = _gcfg.Resolve()
				_gecbb.Set(_dgbafe, _gbbd)
			}
			if _eecdf := _daabag.addObjects(_gbbd); _eecdf != nil {
				return _eecdf
			}
		}
		return nil
	}
	if _, _dbecf := _dcbba.(*_dg.PdfObjectReference); _dbecf {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0062\u0065\u0020\u0061\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u002d\u0020\u0067\u006f\u0074 \u0025\u0023\u0076\u0021", _dcbba)
		return _bf.New("r\u0065\u0066\u0065\u0072en\u0063e\u0020\u006e\u006f\u0074\u0020a\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	return nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_gbgf *PdfColorspaceSpecialPattern) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _gbgf.UnderlyingCS == nil {
		return nil, _bf.New("u\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0043\u0053\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
	}
	return _gbgf.UnderlyingCS.ColorFromFloats(vals)
}

// NewPdfAcroForm returns a new PdfAcroForm with an initialized container (indirect object).
func NewPdfAcroForm() *PdfAcroForm {
	return &PdfAcroForm{Fields: &[]*PdfField{}, _bebfe: _dg.MakeIndirectObject(_dg.MakeDict())}
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_cbfbcg *PdfColorspaceSpecialIndexed) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fbac, _gdffe := _dg.GetNumbersAsFloat(objects)
	if _gdffe != nil {
		return nil, _gdffe
	}
	return _cbfbcg.ColorFromFloats(_fbac)
}

// ToPdfObject returns the PDF representation of the function.
func (_aedda *PdfFunctionType3) ToPdfObject() _dg.PdfObject {
	_adafa := _dg.MakeDict()
	_adafa.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _dg.MakeInteger(3))
	_cefec := &_dg.PdfObjectArray{}
	for _, _bgaf := range _aedda.Domain {
		_cefec.Append(_dg.MakeFloat(_bgaf))
	}
	_adafa.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _cefec)
	if _aedda.Range != nil {
		_fdgac := &_dg.PdfObjectArray{}
		for _, _caaebd := range _aedda.Range {
			_fdgac.Append(_dg.MakeFloat(_caaebd))
		}
		_adafa.Set("\u0052\u0061\u006eg\u0065", _fdgac)
	}
	if _aedda.Functions != nil {
		_bbdc := &_dg.PdfObjectArray{}
		for _, _bdgca := range _aedda.Functions {
			_bbdc.Append(_bdgca.ToPdfObject())
		}
		_adafa.Set("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es", _bbdc)
	}
	if _aedda.Bounds != nil {
		_abaf := &_dg.PdfObjectArray{}
		for _, _accda := range _aedda.Bounds {
			_abaf.Append(_dg.MakeFloat(_accda))
		}
		_adafa.Set("\u0042\u006f\u0075\u006e\u0064\u0073", _abaf)
	}
	if _aedda.Encode != nil {
		_cdfff := &_dg.PdfObjectArray{}
		for _, _ccagfg := range _aedda.Encode {
			_cdfff.Append(_dg.MakeFloat(_ccagfg))
		}
		_adafa.Set("\u0045\u006e\u0063\u006f\u0064\u0065", _cdfff)
	}
	if _aedda._defdg != nil {
		_aedda._defdg.PdfObject = _adafa
		return _aedda._defdg
	}
	return _adafa
}

// Write writes the Appender output to io.Writer.
// It can only be called once and further invocations will result in an error.
func (_dfcb *PdfAppender) Write(w _cf.Writer) error {
	if _dfcb._fdbb {
		return _bf.New("\u0061\u0070\u0070\u0065\u006e\u0064\u0065\u0072\u0020\u0077\u0072\u0069\u0074e\u0020\u0063\u0061\u006e\u0020\u006fn\u006c\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0076\u006f\u006b\u0065\u0064 \u006f\u006e\u0063\u0065")
	}
	_adbf := NewPdfWriter()
	_bfdgc, _gged := _dg.GetDict(_adbf._gbgb)
	if !_gged {
		return _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_cadaf, _gged := _bfdgc.Get("\u004b\u0069\u0064\u0073").(*_dg.PdfObjectArray)
	if !_gged {
		return _bf.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_efbf, _gged := _bfdgc.Get("\u0043\u006f\u0075n\u0074").(*_dg.PdfObjectInteger)
	if !_gged {
		return _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	_gfba := _dfcb._debg._baad
	_fceg := _gfba.GetTrailer()
	if _fceg == nil {
		return _bf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_edgfc, _gged := _dg.GetIndirect(_fceg.Get("\u0052\u006f\u006f\u0074"))
	if !_gged {
		return _bf.New("c\u0061\u0074\u0061\u006c\u006f\u0067 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064")
	}
	_ddf, _gged := _dg.GetDict(_edgfc)
	if !_gged {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _edgfc, *_fceg)
		return _bf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	for _, _agfa := range _ddf.Keys() {
		if _adbf._ecdf.Get(_agfa) == nil {
			_efcf := _ddf.Get(_agfa)
			_adbf._ecdf.Set(_agfa, _efcf)
		}
	}
	if _dfcb._afab != nil {
		if _dfcb._afab._eeecb {
			if _gbcff := _dg.TraceToDirectObject(_dfcb._afab.ToPdfObject()); !_dg.IsNullObject(_gbcff) {
				_adbf._ecdf.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _gbcff)
				_dfcb.updateObjectsDeep(_gbcff, nil)
			} else {
				_ag.Log.Debug("\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020t\u0072\u0061\u0063e\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0020o\u0062\u006a\u0065\u0063\u0074, \u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u002e")
			}
		} else {
			_adbf._ecdf.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _dfcb._afab.ToPdfObject())
			_dfcb.updateObjectsDeep(_dfcb._afab.ToPdfObject(), nil)
		}
	}
	if _dfcb._cfdd != nil {
		_dfcb.updateObjectsDeep(_dfcb._cfdd.ToPdfObject(), nil)
		_adbf._ecdf.Set("\u0044\u0053\u0053", _dfcb._cfdd.GetContainingPdfObject())
	}
	if _dfcb._accbd != nil {
		_adbf._ecdf.Set("\u0050\u0065\u0072m\u0073", _dfcb._accbd.ToPdfObject())
		_dfcb.updateObjectsDeep(_dfcb._accbd.ToPdfObject(), nil)
	}
	if _adbf._efacd.Major < 2 {
		_adbf.AddExtension("\u0045\u0053\u0049\u0043", "\u0031\u002e\u0037", 5)
		_adbf.AddExtension("\u0041\u0044\u0042\u0045", "\u0031\u002e\u0037", 8)
	}
	if _ebba, _defe := _dg.GetDict(_fceg.Get("\u0049\u006e\u0066\u006f")); _defe {
		if _eagf, _cdgc := _dg.GetDict(_adbf._efbfa); _cdgc {
			for _, _ebccb := range _ebba.Keys() {
				if _eagf.Get(_ebccb) == nil {
					_eagf.Set(_ebccb, _ebba.Get(_ebccb))
				}
			}
		}
	}
	if _dfcb._fega != nil {
		_adbf._efbfa = _dg.MakeIndirectObject(_dfcb._fega.ToPdfObject())
	}
	_dfcb.addNewObject(_adbf._efbfa)
	_dfcb.addNewObject(_adbf._fadee)
	_bgcg := false
	if len(_dfcb._debg.PageList) != len(_dfcb._ggdd) {
		_bgcg = true
	} else {
		for _cdc := range _dfcb._debg.PageList {
			switch {
			case _dfcb._ggdd[_cdc] == _dfcb._debg.PageList[_cdc]:
			case _dfcb._ggdd[_cdc] == _dfcb.Reader.PageList[_cdc]:
			default:
				_bgcg = true
			}
			if _bgcg {
				break
			}
		}
	}
	if _bgcg {
		_dfcb.updateObjectsDeep(_adbf._gbgb, nil)
	} else {
		_dfcb._eede[_adbf._gbgb] = struct{}{}
	}
	_adbf._gbgb.ObjectNumber = _dfcb.Reader._eeeef.ObjectNumber
	_dfcb._gebe[_adbf._gbgb] = _dfcb.Reader._eeeef.ObjectNumber
	_gbdag := []_dg.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	for _, _fbcd := range _dfcb._ggdd {
		_bfgf := _fbcd.ToPdfObject()
		*_efbf = *_efbf + 1
		if _acca, _dge := _bfgf.(*_dg.PdfIndirectObject); _dge && _acca.GetParser() == _dfcb._debg._baad {
			_cadaf.Append(&_acca.PdfObjectReference)
			continue
		}
		if _cgdd, _gcgbf := _dg.GetDict(_bfgf); _gcgbf {
			_ggfd, _edag := _cgdd.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_dg.PdfIndirectObject)
			for _edag {
				_ag.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _ggfd)
				_fged, _feggc := _ggfd.PdfObject.(*_dg.PdfObjectDictionary)
				if !_feggc {
					return _bf.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
				}
				for _, _dfe := range _gbdag {
					_ag.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _dfe)
					if _gebc := _cgdd.Get(_dfe); _gebc != nil {
						_ag.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
						if len(_fbcd._gaed.Keys()) > 0 && !_bgcg {
							_cbba := _fbcd._gaed
							if _befbb := _cbba.Get(_dfe); _befbb != nil {
								if _gebc != _befbb {
									_ag.Log.Trace("\u0049\u006e\u0068\u0065\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u006f\u0072\u0069\u0067i\u006ea\u006c\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073\u002c\u0020\u0025\u0054", _dfe, _befbb)
									_cgdd.Set(_dfe, _befbb)
								}
							}
						}
						continue
					}
					if _debd := _fged.Get(_dfe); _debd != nil {
						_ag.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _dfe)
						_cgdd.Set(_dfe, _debd)
					}
				}
				_ggfd, _edag = _fged.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_dg.PdfIndirectObject)
				_ag.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _fged.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
			}
			if _bgcg {
				_cgdd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _adbf._gbgb)
			}
		}
		_dfcb.updateObjectsDeep(_bfgf, nil)
		_cadaf.Append(_bfgf)
	}
	if _, _ebed := _dfcb._cga.Seek(0, _cf.SeekStart); _ebed != nil {
		return _ebed
	}
	_defgc := make(map[SignatureHandler]_cf.Writer)
	_addg := _dg.MakeArray()
	for _, _ffdc := range _dfcb._dgfd {
		if _gcd, _eedg := _dg.GetIndirect(_ffdc); _eedg {
			if _ccca, _fbbe := _gcd.PdfObject.(*pdfSignDictionary); _fbbe {
				_abf := *_ccca._cgdea
				var _fdbeb error
				_defgc[_abf], _fdbeb = _abf.NewDigest(_ccca._cdfca)
				if _fdbeb != nil {
					return _fdbeb
				}
				_addg.Append(_dg.MakeInteger(0xfffff), _dg.MakeInteger(0xfffff))
			}
		}
	}
	if _addg.Len() > 0 {
		_addg.Append(_dg.MakeInteger(0xfffff), _dg.MakeInteger(0xfffff))
	}
	for _, _cdeb := range _dfcb._dgfd {
		if _ceebb, _eddf := _dg.GetIndirect(_cdeb); _eddf {
			if _ccb, _ebda := _ceebb.PdfObject.(*pdfSignDictionary); _ebda {
				_ccb.Set("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _addg)
			}
		}
	}
	_bdbd := len(_defgc) > 0
	var _gaaf _cf.Reader = _dfcb._cga
	if _bdbd {
		_cgac := make([]_cf.Writer, 0, len(_defgc))
		for _, _aeeb := range _defgc {
			_cgac = append(_cgac, _aeeb)
		}
		_gaaf = _cf.TeeReader(_dfcb._cga, _cf.MultiWriter(_cgac...))
	}
	_ddeag, _aab := _cf.Copy(w, _gaaf)
	if _aab != nil {
		return _aab
	}
	if len(_dfcb._dgfd) == 0 {
		return nil
	}
	_adbf._bgggdg = _ddeag
	_adbf.ObjNumOffset = _dfcb._ceecc
	_adbf._bbac = true
	_adbf._cfcdcb = _dfcb._ceef
	_adbf._eege = _dfcb._bcdg
	_adbf._fafab = _dfcb._dcaf
	_adbf._efacd = _dfcb._debg.PdfVersion()
	_adbf._ecabb = _dfcb._gebe
	_adbf._fadcg = _dfcb._gfab.GetCrypter()
	_adbf._acdag = _dfcb._gfab.GetEncryptObj()
	_fbdf := _dfcb._gfab.GetXrefType()
	if _fbdf != nil {
		_gfg := *_fbdf == _dg.XrefTypeObjectStream
		_adbf._eaacb = &_gfg
	}
	_adbf._fdbfa = map[_dg.PdfObject]struct{}{}
	_adbf._agaba = []_dg.PdfObject{}
	for _, _ccbg := range _dfcb._dgfd {
		if _, _abaag := _dfcb._eede[_ccbg]; _abaag {
			continue
		}
		_adbf.addObject(_ccbg)
	}
	_bgab := w
	if _bdbd {
		_bgab = _bc.NewBuffer(nil)
	}
	if _dfcb._acb != "" && _adbf._fadcg == nil {
		_adbf.Encrypt([]byte(_dfcb._acb), []byte(_dfcb._acb), _dfcb._ccfc)
	}
	if _fbdb := _fceg.Get("\u0049\u0044"); _fbdb != nil {
		if _fcg, _ffff := _dg.GetArray(_fbdb); _ffff {
			_adbf._agbeg = _fcg
		}
	}
	if _caace := _adbf.Write(_bgab); _caace != nil {
		return _caace
	}
	if _bdbd {
		_fegd := _bgab.(*_bc.Buffer).Bytes()
		_ebef := _dg.MakeArray()
		var _badcc []*pdfSignDictionary
		var _edad int64
		for _, _ccbd := range _adbf._agaba {
			if _abda, _cgbe := _dg.GetIndirect(_ccbd); _cgbe {
				if _feee, _faa := _abda.PdfObject.(*pdfSignDictionary); _faa {
					_badcc = append(_badcc, _feee)
					_fcge := _feee._cagb + int64(_feee._aaeae)
					_ebef.Append(_dg.MakeInteger(_edad), _dg.MakeInteger(_fcge-_edad))
					_edad = _feee._cagb + int64(_feee._gdbfe)
				}
			}
		}
		_ebef.Append(_dg.MakeInteger(_edad), _dg.MakeInteger(_ddeag+int64(len(_fegd))-_edad))
		_gced := []byte(_ebef.WriteString())
		for _, _cccaf := range _badcc {
			_bdgc := int(_cccaf._cagb - _ddeag)
			for _ccefc := _cccaf._egda; _ccefc < _cccaf._ebaaaa; _ccefc++ {
				_fegd[_bdgc+_ccefc] = ' '
			}
			_agae := _fegd[_bdgc+_cccaf._egda : _bdgc+_cccaf._ebaaaa]
			copy(_agae, _gced)
		}
		var _dceg int
		for _, _gfaf := range _badcc {
			_ffbb := int(_gfaf._cagb - _ddeag)
			_feggca := _fegd[_dceg : _ffbb+_gfaf._aaeae]
			_ddaf := *_gfaf._cgdea
			_defgc[_ddaf].Write(_feggca)
			_dceg = _ffbb + _gfaf._gdbfe
		}
		for _, _egbg := range _badcc {
			_gcge := _fegd[_dceg:]
			_fbeg := *_egbg._cgdea
			_defgc[_fbeg].Write(_gcge)
		}
		for _, _dccd := range _badcc {
			_gdef := int(_dccd._cagb - _ddeag)
			_efcb := *_dccd._cgdea
			_ddae := _defgc[_efcb]
			if _adgg := _efcb.Sign(_dccd._cdfca, _ddae); _adgg != nil {
				return _adgg
			}
			_dccd._cdfca.ByteRange = _ebef
			_egba := []byte(_dccd._cdfca.Contents.WriteString())
			for _fcbb := _dccd._egda; _fcbb < _dccd._ebaaaa; _fcbb++ {
				_fegd[_gdef+_fcbb] = ' '
			}
			for _decca := _dccd._aaeae; _decca < _dccd._gdbfe; _decca++ {
				_fegd[_gdef+_decca] = ' '
			}
			_fade := _fegd[_gdef+_dccd._egda : _gdef+_dccd._ebaaaa]
			copy(_fade, _gced)
			_fade = _fegd[_gdef+_dccd._aaeae : _gdef+_dccd._gdbfe]
			copy(_fade, _egba)
		}
		_bgcd := _bc.NewBuffer(_fegd)
		_, _aab = _cf.Copy(w, _bgcd)
		if _aab != nil {
			return _aab
		}
	}
	_dfcb._fdbb = true
	return nil
}

// SetFilter sets compression filter. Decodes with current filter sets and
// encodes the data with the new filter.
func (_ceffb *XObjectImage) SetFilter(encoder _dg.StreamEncoder) error {
	_bdfgg := _ceffb.Stream
	_gfcde, _fdeec := _ceffb.Filter.DecodeBytes(_bdfgg)
	if _fdeec != nil {
		return _fdeec
	}
	_ceffb.Filter = encoder
	encoder.UpdateParams(_ceffb.getParamsDict())
	_bdfgg, _fdeec = encoder.EncodeBytes(_gfcde)
	if _fdeec != nil {
		return _fdeec
	}
	_ceffb.Stream = _bdfgg
	return nil
}

// NewPdfAppender creates a new Pdf appender from a Pdf reader.
func NewPdfAppender(reader *PdfReader) (*PdfAppender, error) {
	_eega := &PdfAppender{_cga: reader._efcfa, Reader: reader, _gfab: reader._baad, _ebaa: reader._addfg}
	_aage, _gaac := _eega._cga.Seek(0, _cf.SeekEnd)
	if _gaac != nil {
		return nil, _gaac
	}
	_eega._dcaf = _aage
	if _, _gaac = _eega._cga.Seek(0, _cf.SeekStart); _gaac != nil {
		return nil, _gaac
	}
	_eega._debg, _gaac = NewPdfReader(_eega._cga)
	if _gaac != nil {
		return nil, _gaac
	}
	for _, _abbca := range _eega.Reader.GetObjectNums() {
		if _eega._ceecc < _abbca {
			_eega._ceecc = _abbca
		}
	}
	_eega._ceef = _eega._gfab.GetXrefTable()
	_eega._bcdg = _eega._gfab.GetXrefOffset()
	_eega._ggdd = append(_eega._ggdd, _eega._debg.PageList...)
	_eega._efa = make(map[_dg.PdfObject]struct{})
	_eega._gebe = make(map[_dg.PdfObject]int64)
	_eega._eede = make(map[_dg.PdfObject]struct{})
	_eega._afab = _eega._debg.AcroForm
	_eega._cfdd = _eega._debg.DSS
	return _eega, nil
}

type pdfCIDFontType0 struct {
	fontCommon
	_adcdd *_dg.PdfIndirectObject
	_afbff _bd.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// (Required) Dictionary that defines the character collection of the CIDFont.
	// See Table 116.
	CIDSystemInfo *_dg.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW     _dg.PdfObject
	W      _dg.PdfObject
	DW2    _dg.PdfObject
	W2     _dg.PdfObject
	_gfcba map[_bd.CharCode]float64
	_gfdca float64
}

// GetXObjectByName returns the XObject with the specified keyName and the object type.
func (_bcdad *PdfPageResources) GetXObjectByName(keyName _dg.PdfObjectName) (*_dg.PdfObjectStream, XObjectType) {
	if _bcdad.XObject == nil {
		return nil, XObjectTypeUndefined
	}
	_fddcc, _cfce := _dg.TraceToDirectObject(_bcdad.XObject).(*_dg.PdfObjectDictionary)
	if !_cfce {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _dg.TraceToDirectObject(_bcdad.XObject))
		return nil, XObjectTypeUndefined
	}
	if _dagad := _fddcc.Get(keyName); _dagad != nil {
		_fbfg, _cedgd := _dg.GetStream(_dagad)
		if !_cedgd {
			_ag.Log.Debug("X\u004f\u0062\u006a\u0065\u0063\u0074 \u006e\u006f\u0074\u0020\u0070\u006fi\u006e\u0074\u0069\u006e\u0067\u0020\u0074o\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020%\u0054", _dagad)
			return nil, XObjectTypeUndefined
		}
		_bffecg := _fbfg.PdfObjectDictionary
		_bdbbd, _cedgd := _dg.TraceToDirectObject(_bffecg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")).(*_dg.PdfObjectName)
		if !_cedgd {
			_ag.Log.Debug("\u0058\u004fbj\u0065\u0063\u0074 \u0053\u0075\u0062\u0074ype\u0020no\u0074\u0020\u0061\u0020\u004e\u0061\u006de,\u0020\u0064\u0069\u0063\u0074\u003a\u0020%\u0073", _bffecg.String())
			return nil, XObjectTypeUndefined
		}
		if *_bdbbd == "\u0049\u006d\u0061g\u0065" {
			return _fbfg, XObjectTypeImage
		} else if *_bdbbd == "\u0046\u006f\u0072\u006d" {
			return _fbfg, XObjectTypeForm
		} else if *_bdbbd == "\u0050\u0053" {
			return _fbfg, XObjectTypePS
		} else {
			_ag.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0053\u0075b\u0074\u0079\u0070\u0065\u0020\u006e\u006ft\u0020\u006b\u006e\u006f\u0077\u006e\u0020\u0028\u0025\u0073\u0029", *_bdbbd)
			return nil, XObjectTypeUndefined
		}
	} else {
		return nil, XObjectTypeUndefined
	}
}

// IsValid checks if the given pdf output intent type is valid.
func (_dgfab PdfOutputIntentType) IsValid() bool {
	return _dgfab >= PdfOutputIntentTypeA1 && _dgfab <= PdfOutputIntentTypeX
}
func (_bffg *PdfReader) loadStructure() error {
	if _bffg._baad.GetCrypter() != nil && !_bffg._baad.IsAuthenticated() {
		return _b.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_dfff := _bffg._baad.GetTrailer()
	if _dfff == nil {
		return _b.Errorf("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_cecfe, _fgeae := _dfff.Get("\u0052\u006f\u006f\u0074").(*_dg.PdfObjectReference)
	if !_fgeae {
		return _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052\u006f\u006ft\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a \u0025\u0073\u0029", _dfff)
	}
	_gffbc, _feef := _bffg._baad.LookupByReference(*_cecfe)
	if _feef != nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0072\u006f\u006f\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _feef)
		return _feef
	}
	_fffgfd, _fgeae := _gffbc.(*_dg.PdfIndirectObject)
	if !_fgeae {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _gffbc, *_dfff)
		return _bf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_cbbdf, _fgeae := (*_fffgfd).PdfObject.(*_dg.PdfObjectDictionary)
	if !_fgeae {
		_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0061t\u0061\u006c\u006fg\u0020(\u0025\u0073\u0029", _fffgfd.PdfObject)
		return _bf.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_ag.Log.Trace("C\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _cbbdf)
	_aafef, _fgeae := _cbbdf.Get("\u0050\u0061\u0067e\u0073").(*_dg.PdfObjectReference)
	if !_fgeae {
		return _bf.New("\u0070\u0061\u0067\u0065\u0073\u0020\u0069\u006e\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0061\u0020\u0072e\u0066\u0065r\u0065\u006e\u0063\u0065")
	}
	_ffbaa, _feef := _bffg._baad.LookupByReference(*_aafef)
	if _feef != nil {
		_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020F\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020r\u0065\u0061\u0064 \u0070a\u0067\u0065\u0073")
		return _feef
	}
	_eggdc, _fgeae := _ffbaa.(*_dg.PdfIndirectObject)
	if !_fgeae {
		_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020P\u0061\u0067\u0065\u0073\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0069n\u0076a\u006c\u0069\u0064")
		_ag.Log.Debug("\u006f\u0070\u003a\u0020\u0025\u0070", _eggdc)
		return _bf.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_ccabd, _fgeae := _eggdc.PdfObject.(*_dg.PdfObjectDictionary)
	if !_fgeae {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065c\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0073\u0029", _eggdc)
		return _bf.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_bafdb, _fgeae := _dg.GetInt(_ccabd.Get("\u0043\u006f\u0075n\u0074"))
	if !_fgeae {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0050\u0061\u0067\u0065\u0073\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return _bf.New("\u0070\u0061\u0067\u0065s \u0063\u006f\u0075\u006e\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _, _fgeae = _dg.GetName(_ccabd.Get("\u0054\u0079\u0070\u0065")); !_fgeae {
		_ag.Log.Debug("\u0050\u0061\u0067\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020T\u0079\u0070\u0065\u0020\u0066\u0069\u0065\u006cd\u0020n\u006f\u0074\u0020\u0073\u0065\u0074\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0054\u0079p\u0065\u0020\u0074\u006f\u0020\u0050\u0061\u0067\u0065\u0073\u002e")
		_ccabd.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0050\u0061\u0067e\u0073"))
	}
	if _gadecg, _bafa := _dg.GetInt(_ccabd.Get("\u0052\u006f\u0074\u0061\u0074\u0065")); _bafa {
		_dfada := int64(*_gadecg)
		_bffg.Rotate = &_dfada
	}
	_bffg._dfdg = _cecfe
	_bffg._gccfb = _cbbdf
	_bffg._eefc = _ccabd
	_bffg._eeeef = _eggdc
	_bffg._eeeg = int(*_bafdb)
	_bffg._daddd = []*_dg.PdfIndirectObject{}
	_dbfeg := map[_dg.PdfObject]struct{}{}
	_feef = _bffg.buildPageList(_eggdc, nil, _dbfeg)
	if _feef != nil {
		return _feef
	}
	_ag.Log.Trace("\u002d\u002d\u002d")
	_ag.Log.Trace("\u0054\u004f\u0043")
	_ag.Log.Trace("\u0050\u0061\u0067e\u0073")
	_ag.Log.Trace("\u0025\u0064\u003a\u0020\u0025\u0073", len(_bffg._daddd), _bffg._daddd)
	_bffg._fcfc, _feef = _bffg.loadOutlines()
	if _feef != nil {
		_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075i\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 t\u0072\u0065\u0065 \u0028%\u0073\u0029", _feef)
		return _feef
	}
	_bffg.AcroForm, _feef = _bffg.loadForms()
	if _feef != nil {
		return _feef
	}
	_bffg.DSS, _feef = _bffg.loadDSS()
	if _feef != nil {
		return _feef
	}
	_bffg._fbcaa, _feef = _bffg.loadPerms()
	if _feef != nil {
		return _feef
	}
	return nil
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_gage pdfCIDFontType2) GetRuneMetrics(r rune) (_bbg.CharMetrics, bool) {
	_beaeff, _gaeba := _gage._ebgc[r]
	if !_gaeba {
		_bdeee, _fbfbe := _dg.GetInt(_gage.DW)
		if !_fbfbe {
			return _bbg.CharMetrics{}, false
		}
		_beaeff = int(*_bdeee)
	}
	return _bbg.CharMetrics{Wx: float64(_beaeff)}, true
}

// EnableAll LTV enables all signatures in the PDF document.
// The signing certificate chain is extracted from each signature dictionary.
// Optionally, additional certificates can be specified through the
// `extraCerts` parameter. The LTV client attempts to build the certificate
// chain up to a trusted root by downloading any missing certificates.
func (_eabfc *LTV) EnableAll(extraCerts []*_bb.Certificate) error {
	_ddccfa := _eabfc._ebdgd._debg.AcroForm
	for _, _faegd := range _ddccfa.AllFields() {
		_cabdd, _ := _faegd.GetContext().(*PdfFieldSignature)
		if _cabdd == nil {
			continue
		}
		_bdeeec := _cabdd.V
		if _agbc := _eabfc.validateSig(_bdeeec); _agbc != nil {
			_ag.Log.Debug("\u0057\u0041\u0052N\u003a\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _agbc)
		}
		if _acdb := _eabfc.Enable(_bdeeec, extraCerts); _acdb != nil {
			return _acdb
		}
	}
	return nil
}

// NewPdfActionNamed returns a new "named" action.
func NewPdfActionNamed() *PdfActionNamed {
	_gda := NewPdfAction()
	_aeb := &PdfActionNamed{}
	_aeb.PdfAction = _gda
	_gda.SetContext(_aeb)
	return _aeb
}
func (_bccg *PdfReader) newPdfAnnotationLinkFromDict(_dffa *_dg.PdfObjectDictionary) (*PdfAnnotationLink, error) {
	_gbdd := PdfAnnotationLink{}
	_gbdd.A = _dffa.Get("\u0041")
	_gbdd.Dest = _dffa.Get("\u0044\u0065\u0073\u0074")
	_gbdd.H = _dffa.Get("\u0048")
	_gbdd.PA = _dffa.Get("\u0050\u0041")
	_gbdd.QuadPoints = _dffa.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_gbdd.BS = _dffa.Get("\u0042\u0053")
	return &_gbdd, nil
}

// FieldValueProvider provides field values from a data source such as FDF, JSON or any other.
type FieldValueProvider interface {
	FieldValues() (map[string]_dg.PdfObject, error)
}

// PdfPageResourcesColorspaces contains the colorspace in the PdfPageResources.
// Needs to have matching name and colorspace map entry. The Names define the order.
type PdfPageResourcesColorspaces struct {
	Names       []string
	Colorspaces map[string]PdfColorspace
	_dcebf      *_dg.PdfIndirectObject
}

// PdfColor interface represents a generic color in PDF.
type PdfColor interface{}

// NewPdfAnnotationLink returns a new link annotation.
func NewPdfAnnotationLink() *PdfAnnotationLink {
	_ded := NewPdfAnnotation()
	_fdcc := &PdfAnnotationLink{}
	_fdcc.PdfAnnotation = _ded
	_ded.SetContext(_fdcc)
	return _fdcc
}

// String returns a string that describes `font`.
func (_cdgaf *PdfFont) String() string {
	_gabc := ""
	if _cdgaf._cadf.Encoder() != nil {
		_gabc = _cdgaf._cadf.Encoder().String()
	}
	return _b.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0054\u0020\u0025s\u0020\u0025\u0073\u007d", _cdgaf._cadf, _cdgaf.baseFields().coreString(), _gabc)
}

// NewPdfColorDeviceCMYK returns a new CMYK32 color.
func NewPdfColorDeviceCMYK(c, m, y, k float64) *PdfColorDeviceCMYK {
	_ccac := PdfColorDeviceCMYK{c, m, y, k}
	return &_ccac
}

// A PdfPattern can represent a Pattern, either a tiling pattern or a shading pattern.
// Note that all patterns shall be treated as colours; a Pattern colour space shall be established with the CS or cs
// operator just like other colour spaces, and a particular pattern shall be installed as the current colour with the
// SCN or scn operator.
type PdfPattern struct {

	// Type: Pattern
	PatternType int64
	_cgdcc      PdfModel
	_eacce      _dg.PdfObject
}

// GetContainingPdfObject returns the container of the shading object (indirect object).
func (_eeafd *PdfShading) GetContainingPdfObject() _dg.PdfObject { return _eeafd._bcfbg }

// Register registers (caches) a model to primitive object relationship.
func (_ffae *modelManager) Register(primitive _dg.PdfObject, model PdfModel) {
	_ffae._aaddgg[model] = primitive
	_ffae._eccbc[primitive] = model
}

var _ddcf = map[string]struct{}{"\u0054\u0069\u0074l\u0065": {}, "\u0041\u0075\u0074\u0068\u006f\u0072": {}, "\u0053u\u0062\u006a\u0065\u0063\u0074": {}, "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073": {}, "\u0043r\u0065\u0061\u0074\u006f\u0072": {}, "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072": {}, "\u0054r\u0061\u0070\u0070\u0065\u0064": {}, "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065": {}, "\u004do\u0064\u0044\u0061\u0074\u0065": {}}

// NewCustomPdfOutputIntent creates a new custom PdfOutputIntent.
func NewCustomPdfOutputIntent(outputCondition, outputConditionIdentifier, info string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, Info: info, DestOutputProfile: destOutputProfile, _cfaef: _dg.MakeDict(), ColorComponents: colorComponents}
}

// ToPdfObject implements interface PdfModel.
func (_abb *PdfActionSound) ToPdfObject() _dg.PdfObject {
	_abb.PdfAction.ToPdfObject()
	_gdf := _abb._cbd
	_ged := _gdf.PdfObject.(*_dg.PdfObjectDictionary)
	_ged.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeSound)))
	_ged.SetIfNotNil("\u0053\u006f\u0075n\u0064", _abb.Sound)
	_ged.SetIfNotNil("\u0056\u006f\u006c\u0075\u006d\u0065", _abb.Volume)
	_ged.SetIfNotNil("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073", _abb.Synchronous)
	_ged.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _abb.Repeat)
	_ged.SetIfNotNil("\u004d\u0069\u0078", _abb.Mix)
	return _gdf
}

// NewPdfActionGoToE returns a new "go to embedded" action.
func NewPdfActionGoToE() *PdfActionGoToE {
	_ae := NewPdfAction()
	_ecbg := &PdfActionGoToE{}
	_ecbg.PdfAction = _ae
	_ae.SetContext(_ecbg)
	return _ecbg
}

// ButtonType represents the subtype of a button field, can be one of:
// - Checkbox (ButtonTypeCheckbox)
// - PushButton (ButtonTypePushButton)
// - RadioButton (ButtonTypeRadioButton)
type ButtonType int

// ToPdfObject returns a PdfObject representation of PdfColorspaceDeviceNAttributes as a PdfObjectDictionary directly
// or indirectly within an indirect object container.
func (_cefc *PdfColorspaceDeviceNAttributes) ToPdfObject() _dg.PdfObject {
	_cebdb := _dg.MakeDict()
	if _cefc.Subtype != nil {
		_cebdb.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _cefc.Subtype)
	}
	_cebdb.SetIfNotNil("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s", _cefc.Colorants)
	_cebdb.SetIfNotNil("\u0050r\u006f\u0063\u0065\u0073\u0073", _cefc.Process)
	_cebdb.SetIfNotNil("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073", _cefc.MixingHints)
	if _cefc._ffbg != nil {
		_cefc._ffbg.PdfObject = _cebdb
		return _cefc._ffbg
	}
	return _cebdb
}
func _bagde(_fcdd *_dg.PdfObjectDictionary, _abgde *fontCommon) (*pdfFontType3, error) {
	_ecebae := _aacff(_abgde)
	_cagec := _fcdd.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
	if _cagec == nil {
		_cagec = _dg.MakeInteger(0)
	}
	_ecebae.FirstChar = _cagec
	_dbbed, _abbeb := _dg.GetIntVal(_cagec)
	if !_abbeb {
		_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _cagec)
		return nil, _dg.ErrTypeError
	}
	_gbcda := _bd.CharCode(_dbbed)
	_cagec = _fcdd.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
	if _cagec == nil {
		_cagec = _dg.MakeInteger(255)
	}
	_ecebae.LastChar = _cagec
	_dbbed, _abbeb = _dg.GetIntVal(_cagec)
	if !_abbeb {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _cagec)
		return nil, _dg.ErrTypeError
	}
	_gccfd := _bd.CharCode(_dbbed)
	_cagec = _fcdd.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _cagec != nil {
		_ecebae.Resources = _cagec
	}
	_cagec = _fcdd.Get("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s")
	if _cagec == nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0068\u0061\u0072\u0050\u0072\u006f\u0063\u0073\u0020(%\u0076\u0029", _cagec)
		return nil, _dg.ErrNotSupported
	}
	_ecebae.CharProcs = _cagec
	_cagec = _fcdd.Get("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078")
	if _cagec == nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0046\u006f\u006et\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0028\u0025\u0076\u0029", _cagec)
		return nil, _dg.ErrNotSupported
	}
	_ecebae.FontMatrix = _cagec
	_ecebae._bcde = make(map[_bd.CharCode]float64)
	_cagec = _fcdd.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
	if _cagec != nil {
		_ecebae.Widths = _cagec
		_fegae, _dbfgc := _dg.GetArray(_cagec)
		if !_dbfgc {
			_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _cagec)
			return nil, _dg.ErrTypeError
		}
		_cbgg, _gdcg := _fegae.ToFloat64Array()
		if _gdcg != nil {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _gdcg
		}
		if len(_cbgg) != int(_gccfd-_gbcda+1) {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _gccfd-_gbcda+1, len(_cbgg))
			return nil, _dg.ErrRangeError
		}
		_bcbdc, _dbfgc := _dg.GetArray(_ecebae.FontMatrix)
		if !_dbfgc {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _bcbdc)
			return nil, _gdcg
		}
		_fbegf, _gdcg := _bcbdc.ToFloat64Array()
		if _gdcg != nil {
			_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020c\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0046o\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _gdcg
		}
		_bebd := _fec.NewMatrix(_fbegf[0], _fbegf[1], _fbegf[2], _fbegf[3], _fbegf[4], _fbegf[5])
		for _gebcc, _eeagf := range _cbgg {
			_fdda, _ := _bebd.Transform(_eeagf, _eeagf)
			_ecebae._bcde[_gbcda+_bd.CharCode(_gebcc)] = _fdda
		}
	}
	_ecebae.Encoding = _dg.TraceToDirectObject(_fcdd.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	_cdfb := _fcdd.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _cdfb != nil {
		_ecebae._ebbff = _dg.TraceToDirectObject(_cdfb)
		_fecd, _dfba := _bfbgg(_ecebae._ebbff, &_ecebae.fontCommon)
		if _dfba != nil {
			return nil, _dfba
		}
		_ecebae._ecfb = _fecd
	}
	if _bbbf := _ecebae._ecfb; _bbbf != nil {
		_ecebae._geffb = _bd.NewCMapEncoder("", nil, _bbbf)
	} else {
		_ecebae._geffb = _bd.NewPdfDocEncoder()
	}
	return _ecebae, nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_faad *PdfShadingType5) ToPdfObject() _dg.PdfObject {
	_faad.PdfShading.ToPdfObject()
	_acgg, _ebagc := _faad.getShadingDict()
	if _ebagc != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _faad.BitsPerCoordinate != nil {
		_acgg.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _faad.BitsPerCoordinate)
	}
	if _faad.BitsPerComponent != nil {
		_acgg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _faad.BitsPerComponent)
	}
	if _faad.VerticesPerRow != nil {
		_acgg.Set("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077", _faad.VerticesPerRow)
	}
	if _faad.Decode != nil {
		_acgg.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _faad.Decode)
	}
	if _faad.Function != nil {
		if len(_faad.Function) == 1 {
			_acgg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _faad.Function[0].ToPdfObject())
		} else {
			_cbdcd := _dg.MakeArray()
			for _, _fdfde := range _faad.Function {
				_cbdcd.Append(_fdfde.ToPdfObject())
			}
			_acgg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _cbdcd)
		}
	}
	return _faad._bcfbg
}

// GetContainingPdfObject returns the container of the outline (indirect object).
func (_cdgdfb *PdfOutline) GetContainingPdfObject() _dg.PdfObject { return _cdgdfb._bbef }

// EncryptOptions represents encryption options for an output PDF.
type EncryptOptions struct {
	Permissions _gbd.Permissions
	Algorithm   EncryptionAlgorithm
}

func _aacff(_caeaa *fontCommon) *pdfFontType3 { return &pdfFontType3{fontCommon: *_caeaa} }

// NewPdfAnnotationCircle returns a new circle annotation.
func NewPdfAnnotationCircle() *PdfAnnotationCircle {
	_gcfe := NewPdfAnnotation()
	_gdfd := &PdfAnnotationCircle{}
	_gdfd.PdfAnnotation = _gcfe
	_gdfd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gcfe.SetContext(_gdfd)
	return _gdfd
}
func (_agde *PdfReader) flattenFieldsWithOpts(_efbc bool, _gdce FieldAppearanceGenerator, _gcgg *FieldFlattenOpts) error {
	if _gcgg == nil {
		_gcgg = &FieldFlattenOpts{}
	}
	var _fgfde bool
	_bbgd := map[*PdfAnnotation]bool{}
	{
		var _ggbba []*PdfField
		_gacdd := _agde.AcroForm
		if _gacdd != nil {
			if _gcgg.FilterFunc != nil {
				_ggbba = _gacdd.filteredFields(_gcgg.FilterFunc, true)
				_fgfde = _gacdd.Fields != nil && len(*_gacdd.Fields) > 0
			} else {
				_ggbba = _gacdd.AllFields()
			}
		}
		for _, _dddca := range _ggbba {
			if len(_dddca.Annotations) < 1 {
				_ag.Log.Debug("\u004e\u006f\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006ed\u0020\u0066\u006f\u0072\u003a\u0020\u0025v\u002c\u0020\u006c\u006f\u006f\u006b\u0020\u0069\u006e\u0074\u006f \u004b\u0069\u0064\u0073\u0020\u004f\u0062\u006a\u0065\u0063\u0074", _dddca.PartialName())
				for _gbcfa, _gddd := range _dddca.Kids {
					for _, _gfdaa := range _gddd.Annotations {
						_bbgd[_gfdaa.PdfAnnotation] = _dddca.V != nil
						if _gddd.V == nil {
							_gddd.V = _dddca.V
						}
						if _gddd.T == nil {
							_gddd.T = _dg.MakeString(_b.Sprintf("\u0025\u0073\u0023%\u0064", _dddca.PartialName(), _gbcfa))
						}
						if _gdce != nil {
							_afaa, _ddad := _gdce.GenerateAppearanceDict(_gacdd, _gddd, _gfdaa)
							if _ddad != nil {
								return _ddad
							}
							_gfdaa.AP = _afaa
						}
					}
				}
			}
			for _, _ccec := range _dddca.Annotations {
				_bbgd[_ccec.PdfAnnotation] = _dddca.V != nil
				if _gdce != nil {
					_badee, _cdefd := _gdce.GenerateAppearanceDict(_gacdd, _dddca, _ccec)
					if _cdefd != nil {
						return _cdefd
					}
					_ccec.AP = _badee
				}
			}
		}
	}
	if _efbc {
		for _, _fcgb := range _agde.PageList {
			_fecb, _dccgac := _fcgb.GetAnnotations()
			if _dccgac != nil {
				return _dccgac
			}
			for _, _gefb := range _fecb {
				_bbgd[_gefb] = true
			}
		}
	}
	for _, _ebaag := range _agde.PageList {
		_dada := _ebaag.flattenFieldsWithOpts(_gdce, _gcgg, _bbgd)
		if _dada != nil {
			return _dada
		}
	}
	if !_fgfde {
		_agde.AcroForm = nil
	}
	return nil
}

// StandardValidator is the interface that is used for the PDF StandardImplementer validation for the PDF document.
// It is using a CompliancePdfReader which is expected to give more Metadata during reading process.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardValidator interface {

	// ValidateStandard checks if the input reader
	ValidateStandard(_edge *CompliancePdfReader) error
}

// ColorAt returns the color of the image pixel specified by the x and y coordinates.
func (_abfaa *Image) ColorAt(x, y int) (_edg.Color, error) {
	_bcfa := _fc.BytesPerLine(int(_abfaa.Width), int(_abfaa.BitsPerComponent), _abfaa.ColorComponents)
	switch _abfaa.ColorComponents {
	case 1:
		return _fc.ColorAtGrayscale(x, y, int(_abfaa.BitsPerComponent), _bcfa, _abfaa.Data, _abfaa._gfbb)
	case 3:
		return _fc.ColorAtNRGBA(x, y, int(_abfaa.Width), _bcfa, int(_abfaa.BitsPerComponent), _abfaa.Data, _abfaa._dgeb, _abfaa._gfbb)
	case 4:
		return _fc.ColorAtCMYK(x, y, int(_abfaa.Width), _abfaa.Data, _abfaa._gfbb)
	}
	_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 i\u006da\u0067\u0065\u002e\u0020\u0025\u0064\u0020\u0063\u006f\u006d\u0070\u006fn\u0065\u006e\u0074\u0073\u002c\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072 \u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _abfaa.ColorComponents, _abfaa.BitsPerComponent)
	return nil, _bf.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061g\u0065 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
}

// GetOptimizer returns current PDF optimizer.
func (_dgbfe *PdfWriter) GetOptimizer() Optimizer { return _dgbfe._egeac }
func _aacd(_efeba *PdfField) []*PdfField {
	_fegab := []*PdfField{_efeba}
	for _, _geac := range _efeba.Kids {
		_fegab = append(_fegab, _aacd(_geac)...)
	}
	return _fegab
}

// GetContainingPdfObject returns the container of the image object (indirect object).
func (_dacbb *XObjectImage) GetContainingPdfObject() _dg.PdfObject { return _dacbb._abfb }

// R returns the value of the red component of the color.
func (_egeg *PdfColorDeviceRGB) R() float64 { return _egeg[0] }

// CharcodeBytesToUnicode converts PDF character codes `data` to a Go unicode string.
//
// 9.10 Extraction of Text Content (page 292)
// The process of finding glyph descriptions in OpenType fonts by a conforming reader shall be the following:
//   - For Type 1 fonts using “CFF” tables, the process shall be as described in 9.6.6.2, "Encodings
//     for Type 1 Fonts".
//   - For TrueType fonts using “glyf” tables, the process shall be as described in 9.6.6.4,
//     "Encodings for TrueType Fonts". Since this process sometimes produces ambiguous results,
//     conforming writers, instead of using a simple font, shall use a Type 0 font with an Identity-H
//     encoding and use the glyph indices as character codes, as described following Table 118.
func (_bbee *PdfFont) CharcodeBytesToUnicode(data []byte) (string, int, int) {
	_cega, _, _aaedg := _bbee.CharcodesToUnicodeWithStats(_bbee.BytesToCharcodes(data))
	_fbaca := _bd.ExpandLigatures(_cega)
	return _fbaca, _gbf.RuneCountInString(_fbaca), _aaedg
}

// NewPdfActionResetForm returns a new "reset form" action.
func NewPdfActionResetForm() *PdfActionResetForm {
	_ab := NewPdfAction()
	_ggf := &PdfActionResetForm{}
	_ggf.PdfAction = _ab
	_ab.SetContext(_ggf)
	return _ggf
}

// NewPdfAnnotationRichMedia returns a new rich media annotation.
func NewPdfAnnotationRichMedia() *PdfAnnotationRichMedia {
	_cad := NewPdfAnnotation()
	_efgg := &PdfAnnotationRichMedia{}
	_efgg.PdfAnnotation = _cad
	_cad.SetContext(_efgg)
	return _efgg
}

var _ pdfFont = (*pdfFontType3)(nil)

// EnableChain adds the specified certificate chain and validation data (OCSP
// and CRL information) for it to the global scope of the document DSS. The
// added data is used for validating any of the signatures present in the
// document. The LTV client attempts to build the certificate chain up to a
// trusted root by downloading any missing certificates.
func (_fbbea *LTV) EnableChain(chain []*_bb.Certificate) error { return _fbbea.enable(nil, chain, "") }

// EncryptionAlgorithm is used in EncryptOptions to change the default algorithm used to encrypt the document.
type EncryptionAlgorithm int

func (_aacec *PdfPage) flattenFieldsWithOpts(_ccgg FieldAppearanceGenerator, _cfcg *FieldFlattenOpts, _gccf map[*PdfAnnotation]bool) error {
	var _eaded []*PdfAnnotation
	if _ccgg != nil {
		if _edaec := _ccgg.WrapContentStream(_aacec); _edaec != nil {
			return _edaec
		}
	}
	_ceeg, _daaa := _aacec.GetAnnotations()
	if _daaa != nil {
		return _daaa
	}
	for _, _bdef := range _ceeg {
		_bgeca, _eacdb := _gccf[_bdef]
		if !_eacdb && _cfcg.AnnotFilterFunc != nil {
			if _, _aaag := _bdef.GetContext().(*PdfAnnotationWidget); !_aaag {
				_eacdb = _cfcg.AnnotFilterFunc(_bdef)
			}
		}
		if !_eacdb {
			_eaded = append(_eaded, _bdef)
			continue
		}
		switch _bdef.GetContext().(type) {
		case *PdfAnnotationPopup:
			continue
		case *PdfAnnotationLink:
			continue
		case *PdfAnnotationProjection:
			continue
		}
		_efbag, _bbgg, _gfgag := _dbag(_bdef)
		if _gfgag != nil {
			if !_bgeca {
				_ag.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0056\u0020\u002d\u003e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075t\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065am\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072")
				continue
			}
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d,\u0020\u0065\u0072\u0072\u0020\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006bi\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072", _gfgag)
			continue
		}
		if _efbag == nil {
			continue
		}
		_eadb := _aacec.Resources.GenerateXObjectName()
		_aacec.Resources.SetXObjectFormByName(_eadb, _efbag)
		_cgcdb, _gaccg, _gfgag := _dfgb(_efbag)
		if _gfgag != nil {
			_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0061\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u004d\u0061\u0074\u0072\u0069\u0078\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0078\u0066\u006f\u0072\u006d\u0020\u0062\u0062\u006f\u0078\u0020\u0061\u0064\u006a\u0075\u0073t\u006d\u0065\u006e\u0074\u003a \u0025\u0076", _gfgag)
		} else {
			_gegf := _fec.IdentityMatrix()
			_gegf = _gegf.Translate(-_cgcdb.Llx, -_cgcdb.Lly)
			if _gaccg {
				_aeaed := 0.0
				if _cgcdb.Width() > 0 {
					_aeaed = _bbgg.Width() / _cgcdb.Width()
				}
				_aaaae := 0.0
				if _cgcdb.Height() > 0 {
					_aaaae = _bbgg.Height() / _cgcdb.Height()
				}
				_gegf = _gegf.Scale(_aeaed, _aaaae)
			}
			_bbgg.Transform(_gegf)
		}
		_cbag := _cg.Min(_bbgg.Llx, _bbgg.Urx)
		_afee := _cg.Min(_bbgg.Lly, _bbgg.Ury)
		var _adege []string
		_adege = append(_adege, "\u0071")
		_adege = append(_adege, _b.Sprintf("\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020%\u002e6\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0063\u006d", 1.0, 0.0, 0.0, 1.0, _cbag, _afee))
		_adege = append(_adege, _b.Sprintf("\u002f\u0025\u0073\u0020\u0044\u006f", _eadb.String()))
		_adege = append(_adege, "\u0051")
		_fdegg := _ga.Join(_adege, "\u000a")
		_gfgag = _aacec.AppendContentStream(_fdegg)
		if _gfgag != nil {
			return _gfgag
		}
		if _efbag.Resources != nil {
			_daffe, _ffbd := _dg.GetDict(_efbag.Resources.Font)
			if _ffbd {
				for _, _caba := range _daffe.Keys() {
					if !_aacec.Resources.HasFontByName(_caba) {
						_aacec.Resources.SetFontByName(_caba, _daffe.Get(_caba))
					}
				}
			}
		}
	}
	if len(_eaded) > 0 {
		_aacec._cadgg = _eaded
	} else {
		_aacec._cadgg = []*PdfAnnotation{}
	}
	return nil
}

// PdfAnnotationHighlight represents Highlight annotations.
// (Section 12.5.6.10).
type PdfAnnotationHighlight struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _dg.PdfObject
}

// NewPdfShadingPatternType2 creates an empty shading pattern type 2 object.
func NewPdfShadingPatternType2() *PdfShadingPatternType2 {
	_bcdgcg := &PdfShadingPatternType2{}
	_bcdgcg.Matrix = _dg.MakeArrayFromIntegers([]int{1, 0, 0, 1, 0, 0})
	_bcdgcg.PdfPattern = &PdfPattern{}
	_bcdgcg.PdfPattern.PatternType = int64(*_dg.MakeInteger(2))
	_bcdgcg.PdfPattern._cgdcc = _bcdgcg
	_bcdgcg.PdfPattern._eacce = _dg.MakeIndirectObject(_dg.MakeDict())
	return _bcdgcg
}

// AddOCSPs adds OCSPs to DSS.
func (_gbfc *DSS) AddOCSPs(ocsps [][]byte) ([]*_dg.PdfObjectStream, error) {
	return _gbfc.add(&_gbfc.OCSPs, _gbfc._bggg, ocsps)
}

// NewPdfAnnotationPolygon returns a new polygon annotation.
func NewPdfAnnotationPolygon() *PdfAnnotationPolygon {
	_cggf := NewPdfAnnotation()
	_ceebf := &PdfAnnotationPolygon{}
	_ceebf.PdfAnnotation = _cggf
	_ceebf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cggf.SetContext(_ceebf)
	return _ceebf
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_gfff *PdfShadingPatternType2) ToPdfObject() _dg.PdfObject {
	_gfff.PdfPattern.ToPdfObject()
	_fffeb := _gfff.getDict()
	if _gfff.Shading != nil {
		_fffeb.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _gfff.Shading.ToPdfObject())
	}
	if _gfff.Matrix != nil {
		_fffeb.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _gfff.Matrix)
	}
	if _gfff.ExtGState != nil {
		_fffeb.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _gfff.ExtGState)
	}
	return _gfff._eacce
}

// DecodeArray returns an empty slice as there are no components associated with pattern colorspace.
func (_debb *PdfColorspaceSpecialPattern) DecodeArray() []float64 { return []float64{} }

// NewPdfReaderLazy creates a new PdfReader for `rs` in lazy-loading mode. The difference
// from NewPdfReader is that in lazy-loading mode, objects are only loaded into memory when needed
// rather than entire structure being loaded into memory on reader creation.
// Note that it may make sense to use the lazy-load reader when processing only parts of files,
// rather than loading entire file into memory. Example: splitting a few pages from a large PDF file.
func NewPdfReaderLazy(rs _cf.ReadSeeker) (*PdfReader, error) {
	const _fbaf = "\u006d\u006f\u0064\u0065l:\u004e\u0065\u0077\u0050\u0064\u0066\u0052\u0065\u0061\u0064\u0065\u0072\u004c\u0061z\u0079"
	return _dcdd(rs, &ReaderOpts{LazyLoad: true}, false, _fbaf)
}

// Resample resamples the image data converting from current BitsPerComponent to a target BitsPerComponent
// value.  Sets the image's BitsPerComponent to the target value following resampling.
//
// For example, converting an 8-bit RGB image to 1-bit grayscale (common for scanned images):
//
//	// Convert RGB image to grayscale.
//	rgbColorSpace := pdf.NewPdfColorspaceDeviceRGB()
//	grayImage, err := rgbColorSpace.ImageToGray(rgbImage)
//	if err != nil {
//	  return err
//	}
//	// Resample as 1 bit.
//	grayImage.Resample(1)
func (_bddcb *Image) Resample(targetBitsPerComponent int64) {
	if _bddcb.BitsPerComponent == targetBitsPerComponent {
		return
	}
	_ebggd := _bddcb.GetSamples()
	if targetBitsPerComponent < _bddcb.BitsPerComponent {
		_ffcbb := _bddcb.BitsPerComponent - targetBitsPerComponent
		for _ceege := range _ebggd {
			_ebggd[_ceege] >>= uint(_ffcbb)
		}
	} else if targetBitsPerComponent > _bddcb.BitsPerComponent {
		_dbed := targetBitsPerComponent - _bddcb.BitsPerComponent
		for _bfefg := range _ebggd {
			_ebggd[_bfefg] <<= uint(_dbed)
		}
	}
	_bddcb.BitsPerComponent = targetBitsPerComponent
	if _bddcb.BitsPerComponent < 8 {
		_bddcb.resampleLowBits(_ebggd)
		return
	}
	_cabbg := _fc.BytesPerLine(int(_bddcb.Width), int(_bddcb.BitsPerComponent), _bddcb.ColorComponents)
	_gded := make([]byte, _cabbg*int(_bddcb.Height))
	var (
		_bgfcc, _acdgc, _eabf, _dacab int
		_fcdfb                        uint32
	)
	for _eabf = 0; _eabf < int(_bddcb.Height); _eabf++ {
		_bgfcc = _eabf * _cabbg
		_acdgc = (_eabf+1)*_cabbg - 1
		_bedce := _fcd.ResampleUint32(_ebggd[_bgfcc:_acdgc], int(targetBitsPerComponent), 8)
		for _dacab, _fcdfb = range _bedce {
			_gded[_dacab+_bgfcc] = byte(_fcdfb)
		}
	}
	_bddcb.Data = _gded
}
func (_ddbbg *fontFile) loadFromSegments(_aagcc, _acaef []byte) error {
	_ag.Log.Trace("\u006c\u006f\u0061dF\u0072\u006f\u006d\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u0020\u0025\u0064", len(_aagcc), len(_acaef))
	_fadgd := _ddbbg.parseASCIIPart(_aagcc)
	if _fadgd != nil {
		return _fadgd
	}
	_ag.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _ddbbg)
	if len(_acaef) == 0 {
		return nil
	}
	_ag.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _ddbbg)
	return nil
}

// PdfColorspaceDeviceCMYK represents a CMYK32 colorspace.
type PdfColorspaceDeviceCMYK struct{}

// ToPdfObject returns the PdfFontDescriptor as a PDF dictionary inside an indirect object.
func (_aebg *PdfFontDescriptor) ToPdfObject() _dg.PdfObject {
	_afgdd := _dg.MakeDict()
	if _aebg._acag == nil {
		_aebg._acag = &_dg.PdfIndirectObject{}
	}
	_aebg._acag.PdfObject = _afgdd
	_afgdd.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072"))
	if _aebg.FontName != nil {
		_afgdd.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _aebg.FontName)
	}
	if _aebg.FontFamily != nil {
		_afgdd.Set("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079", _aebg.FontFamily)
	}
	if _aebg.FontStretch != nil {
		_afgdd.Set("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068", _aebg.FontStretch)
	}
	if _aebg.FontWeight != nil {
		_afgdd.Set("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074", _aebg.FontWeight)
	}
	if _aebg.Flags != nil {
		_afgdd.Set("\u0046\u006c\u0061g\u0073", _aebg.Flags)
	}
	if _aebg.FontBBox != nil {
		_afgdd.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _aebg.FontBBox)
	}
	if _aebg.ItalicAngle != nil {
		_afgdd.Set("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065", _aebg.ItalicAngle)
	}
	if _aebg.Ascent != nil {
		_afgdd.Set("\u0041\u0073\u0063\u0065\u006e\u0074", _aebg.Ascent)
	}
	if _aebg.Descent != nil {
		_afgdd.Set("\u0044e\u0073\u0063\u0065\u006e\u0074", _aebg.Descent)
	}
	if _aebg.Leading != nil {
		_afgdd.Set("\u004ce\u0061\u0064\u0069\u006e\u0067", _aebg.Leading)
	}
	if _aebg.CapHeight != nil {
		_afgdd.Set("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t", _aebg.CapHeight)
	}
	if _aebg.XHeight != nil {
		_afgdd.Set("\u0058H\u0065\u0069\u0067\u0068\u0074", _aebg.XHeight)
	}
	if _aebg.StemV != nil {
		_afgdd.Set("\u0053\u0074\u0065m\u0056", _aebg.StemV)
	}
	if _aebg.StemH != nil {
		_afgdd.Set("\u0053\u0074\u0065m\u0048", _aebg.StemH)
	}
	if _aebg.AvgWidth != nil {
		_afgdd.Set("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068", _aebg.AvgWidth)
	}
	if _aebg.MaxWidth != nil {
		_afgdd.Set("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068", _aebg.MaxWidth)
	}
	if _aebg.MissingWidth != nil {
		_afgdd.Set("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068", _aebg.MissingWidth)
	}
	if _aebg.FontFile != nil {
		_afgdd.Set("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065", _aebg.FontFile)
	}
	if _aebg.FontFile2 != nil {
		_afgdd.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652", _aebg.FontFile2)
	}
	if _aebg.FontFile3 != nil {
		_afgdd.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653", _aebg.FontFile3)
	}
	if _aebg.CharSet != nil {
		_afgdd.Set("\u0043h\u0061\u0072\u0053\u0065\u0074", _aebg.CharSet)
	}
	if _aebg.Style != nil {
		_afgdd.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _aebg.FontName)
	}
	if _aebg.Lang != nil {
		_afgdd.Set("\u004c\u0061\u006e\u0067", _aebg.Lang)
	}
	if _aebg.FD != nil {
		_afgdd.Set("\u0046\u0044", _aebg.FD)
	}
	if _aebg.CIDSet != nil {
		_afgdd.Set("\u0043\u0049\u0044\u0053\u0065\u0074", _aebg.CIDSet)
	}
	return _aebg._acag
}

// PdfColorCalRGB represents a color in the Colorimetric CIE RGB colorspace.
// A, B, C components
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorCalRGB [3]float64

// NewPdfPageResourcesFromDict creates and returns a new PdfPageResources object
// from the input dictionary.
func NewPdfPageResourcesFromDict(dict *_dg.PdfObjectDictionary) (*PdfPageResources, error) {
	_bdeeb := NewPdfPageResources()
	if _egedd := dict.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _egedd != nil {
		_bdeeb.ExtGState = _egedd
	}
	if _fgcg := dict.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"); _fgcg != nil && !_dg.IsNullObject(_fgcg) {
		_bdeeb.ColorSpace = _fgcg
	}
	if _eeebfg := dict.Get("\u0050a\u0074\u0074\u0065\u0072\u006e"); _eeebfg != nil {
		_bdeeb.Pattern = _eeebfg
	}
	if _feac := dict.Get("\u0053h\u0061\u0064\u0069\u006e\u0067"); _feac != nil {
		_bdeeb.Shading = _feac
	}
	if _fdcca := dict.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"); _fdcca != nil {
		_bdeeb.XObject = _fdcca
	}
	if _bddg := _dg.ResolveReference(dict.Get("\u0046\u006f\u006e\u0074")); _bddg != nil {
		_bdeeb.Font = _bddg
	}
	if _bceae := dict.Get("\u0050r\u006f\u0063\u0053\u0065\u0074"); _bceae != nil {
		_bdeeb.ProcSet = _bceae
	}
	if _deea := dict.Get("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"); _deea != nil {
		_bdeeb.Properties = _deea
	}
	return _bdeeb, nil
}

// GetContainingPdfObject implements interface PdfModel.
func (_bcad *PdfAnnotation) GetContainingPdfObject() _dg.PdfObject { return _bcad._cdf }
func _dbag(_cdffd *PdfAnnotation) (*XObjectForm, *PdfRectangle, error) {
	_cdgbc, _adcg := _dg.GetDict(_cdffd.AP)
	if !_adcg {
		return nil, nil, _bf.New("f\u0069\u0065\u006c\u0064\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0041\u0050\u0020d\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079")
	}
	if _cdgbc == nil {
		return nil, nil, nil
	}
	_deecf, _adcg := _dg.GetArray(_cdffd.Rect)
	if !_adcg || _deecf.Len() != 4 {
		return nil, nil, _bf.New("\u0072\u0065\u0063t\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_dbab, _cbagg := NewPdfRectangle(*_deecf)
	if _cbagg != nil {
		return nil, nil, _cbagg
	}
	_adff := _dg.TraceToDirectObject(_cdgbc.Get("\u004e"))
	switch _agff := _adff.(type) {
	case *_dg.PdfObjectStream:
		_dgbd := _agff
		_cccafa, _ddgde := NewXObjectFormFromStream(_dgbd)
		return _cccafa, _dbab, _ddgde
	case *_dg.PdfObjectDictionary:
		_cgggf := _agff
		_effbb, _deac := _dg.GetName(_cdffd.AS)
		if !_deac {
			return nil, nil, nil
		}
		if _cgggf.Get(*_effbb) == nil {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0041\u0053\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
			return nil, nil, nil
		}
		_cabc, _deac := _dg.GetStream(_cgggf.Get(*_effbb))
		if !_deac {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006ce \u0074\u006f\u0020\u0061\u0063\u0063e\u0073\u0073\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0066\u006f\u0072 \u0025\u0076", _effbb)
			return nil, nil, _bf.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		}
		_caad, _aggb := NewXObjectFormFromStream(_cabc)
		return _caad, _dbab, _aggb
	}
	_ag.Log.Debug("\u0049\u006e\u0076\u0061li\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u004e\u003a\u0020%\u0054", _adff)
	return nil, nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
}
func _daece() string { _fgefgf.Lock(); defer _fgefgf.Unlock(); return _fgfda }

// PdfOutlineTreeNode contains common fields used by the outline and outline
// item objects.
type PdfOutlineTreeNode struct {
	_baddf interface{}
	First  *PdfOutlineTreeNode
	Last   *PdfOutlineTreeNode
}

func (_fccd *PdfWriter) mapObjectStreams(_feggg bool) (map[_dg.PdfObject]bool, bool) {
	_afefb := make(map[_dg.PdfObject]bool)
	for _, _fcfcc := range _fccd._agaba {
		if _cbada, _ccfeg := _fcfcc.(*_dg.PdfObjectStreams); _ccfeg {
			_feggg = true
			for _, _egegbg := range _cbada.Elements() {
				_afefb[_egegbg] = true
				if _gdcgb, _gcaa := _egegbg.(*_dg.PdfIndirectObject); _gcaa {
					_afefb[_gdcgb.PdfObject] = true
				}
			}
		}
	}
	return _afefb, _feggg
}

// GetContentStream returns the pattern cell's content stream
func (_fdfef *PdfTilingPattern) GetContentStream() ([]byte, error) {
	_ddgcc, _, _edfea := _fdfef.GetContentStreamWithEncoder()
	return _ddgcc, _edfea
}

// NewPdfDate returns a new PdfDate object from a PDF date string (see 7.9.4 Dates).
// format: "D: YYYYMMDDHHmmSSOHH'mm"
func NewPdfDate(dateStr string) (PdfDate, error) {
	_dgee, _bcba := _bcc.ParsePdfTime(dateStr)
	if _bcba != nil {
		return PdfDate{}, _bcba
	}
	return NewPdfDateFromTime(_dgee)
}
func (_edee *PdfReader) newPdfAnnotationMarkupFromDict(_dbf *_dg.PdfObjectDictionary) (*PdfAnnotationMarkup, error) {
	_dccga := &PdfAnnotationMarkup{}
	if _effa := _dbf.Get("\u0054"); _effa != nil {
		_dccga.T = _effa
	}
	if _cagf := _dbf.Get("\u0050\u006f\u0070u\u0070"); _cagf != nil {
		_fdd, _gdc := _cagf.(*_dg.PdfIndirectObject)
		if !_gdc {
			if _, _faf := _cagf.(*_dg.PdfObjectNull); !_faf {
				return nil, _bf.New("p\u006f\u0070\u0075\u0070\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074\u0020t\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			}
		} else {
			_cagfe, _bbeb := _edee.newPdfAnnotationFromIndirectObject(_fdd)
			if _bbeb != nil {
				return nil, _bbeb
			}
			if _cagfe != nil {
				_fgc, _gdd := _cagfe._egcg.(*PdfAnnotationPopup)
				if !_gdd {
					return nil, _bf.New("\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0020\u0070\u006f\u0070\u0075\u0070\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e")
				}
				_dccga.Popup = _fgc
			}
		}
	}
	if _aac := _dbf.Get("\u0043\u0041"); _aac != nil {
		_dccga.CA = _aac
	}
	if _bcfe := _dbf.Get("\u0052\u0043"); _bcfe != nil {
		_dccga.RC = _bcfe
	}
	if _bfgg := _dbf.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"); _bfgg != nil {
		_dccga.CreationDate = _bfgg
	}
	if _bgdb := _dbf.Get("\u0049\u0052\u0054"); _bgdb != nil {
		_dccga.IRT = _bgdb
	}
	if _agab := _dbf.Get("\u0053\u0075\u0062\u006a"); _agab != nil {
		_dccga.Subj = _agab
	}
	if _efbe := _dbf.Get("\u0052\u0054"); _efbe != nil {
		_dccga.RT = _efbe
	}
	if _fbfc := _dbf.Get("\u0049\u0054"); _fbfc != nil {
		_dccga.IT = _fbfc
	}
	if _cfeaf := _dbf.Get("\u0045\u0078\u0044\u0061\u0074\u0061"); _cfeaf != nil {
		_dccga.ExData = _cfeaf
	}
	return _dccga, nil
}

// GetNameDictionary returns the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_afgad *PdfReader) GetNameDictionary() (_dg.PdfObject, error) {
	_dgbdc := _dg.ResolveReference(_afgad._gccfb.Get("\u004e\u0061\u006de\u0073"))
	if _dgbdc == nil {
		return nil, nil
	}
	if !_afgad._dadcef {
		_aeecg := _afgad.traverseObjectData(_dgbdc)
		if _aeecg != nil {
			return nil, _aeecg
		}
	}
	return _dgbdc, nil
}
func _gacf(_aeee *_dg.PdfObjectDictionary) (*PdfFieldText, error) {
	_fbbfe := &PdfFieldText{}
	_fbbfe.DA, _ = _dg.GetString(_aeee.Get("\u0044\u0041"))
	_fbbfe.Q, _ = _dg.GetInt(_aeee.Get("\u0051"))
	_fbbfe.DS, _ = _dg.GetString(_aeee.Get("\u0044\u0053"))
	_fbbfe.RV = _aeee.Get("\u0052\u0056")
	_fbbfe.MaxLen, _ = _dg.GetInt(_aeee.Get("\u004d\u0061\u0078\u004c\u0065\u006e"))
	return _fbbfe, nil
}

// GetRotate gets the inheritable rotate value, either from the page
// or a higher up page/pages struct.
func (_dfdd *PdfPage) GetRotate() (int64, error) {
	if _dfdd.Rotate != nil {
		return *_dfdd.Rotate, nil
	}
	_gbecb := _dfdd.Parent
	for _gbecb != nil {
		_gbgfff, _bbdde := _dg.GetDict(_gbecb)
		if !_bbdde {
			return 0, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _gbcebc := _gbgfff.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _gbcebc != nil {
			_baegf, _fbgfg := _dg.GetInt(_gbcebc)
			if !_fbgfg {
				return 0, _bf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0074a\u0074\u0065\u0020\u0076al\u0075\u0065")
			}
			if _baegf != nil {
				return int64(*_baegf), nil
			}
			return 0, _bf.New("\u0072\u006f\u0074\u0061te\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		_gbecb = _gbgfff.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return 0, _bf.New("\u0072o\u0074a\u0074\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
}
func _fecee(_ffdee []*_dg.PdfObjectStream) *_dg.PdfObjectArray {
	if len(_ffdee) == 0 {
		return nil
	}
	_faebf := make([]_dg.PdfObject, 0, len(_ffdee))
	for _, _gfedd := range _ffdee {
		_faebf = append(_faebf, _gfedd)
	}
	return _dg.MakeArray(_faebf...)
}
func _dgagb() string { return _ag.Version }

// GetCharMetrics returns the char metrics for character code `code`.
// How it works:
//  1. It calls the GetCharMetrics function for the underlying font, either a simple font or
//     a Type0 font. The underlying font GetCharMetrics() functions do direct charcode ➞  metrics
//     mappings.
//  2. If the underlying font's GetCharMetrics() doesn't have a CharMetrics for `code` then a
//     a CharMetrics with the FontDescriptor's /MissingWidth is returned.
//  3. If there is no /MissingWidth then a failure is returned.
//
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
// well give them 0 width. There is no need for the bool return.
//
// TODO(gunnsth): Reconsider whether needed or if can map via GlyphName.
func (_daca *PdfFont) GetCharMetrics(code _bd.CharCode) (CharMetrics, bool) {
	var _bcgea _bbg.CharMetrics
	switch _dcgcg := _daca._cadf.(type) {
	case *pdfFontSimple:
		if _cffag, _egfdg := _dcgcg.GetCharMetrics(code); _egfdg {
			return _cffag, _egfdg
		}
	case *pdfFontType0:
		if _bgce, _fgfgf := _dcgcg.GetCharMetrics(code); _fgfgf {
			return _bgce, _fgfgf
		}
	case *pdfCIDFontType0:
		if _ebfa, _fdeb := _dcgcg.GetCharMetrics(code); _fdeb {
			return _ebfa, _fdeb
		}
	case *pdfCIDFontType2:
		if _bgca, _bgbc := _dcgcg.GetCharMetrics(code); _bgbc {
			return _bgca, _bgbc
		}
	case *pdfFontType3:
		if _bbcfc, _fgeb := _dcgcg.GetCharMetrics(code); _fgeb {
			return _bbcfc, _fgeb
		}
	default:
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020G\u0065\u0074\u0043h\u0061\u0072\u004de\u0074\u0072i\u0063\u0073\u0020\u006e\u006f\u0074 \u0069mp\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d\u0025\u0054\u002e", _daca._cadf)
		return _bcgea, false
	}
	if _bgcfb, _cbgeg := _daca.GetFontDescriptor(); _cbgeg == nil && _bgcfb != nil {
		return _bbg.CharMetrics{Wx: _bgcfb._adggg}, true
	}
	_ag.Log.Debug("\u0047\u0065\u0074\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069c\u0073\u0020\u0066\u006f\u0072 \u0066\u006fn\u0074\u003d\u0025\u0073", _daca)
	return _bcgea, false
}

// PdfOutlineItem represents an outline item dictionary (Table 153 - pp. 376 - 377).
type PdfOutlineItem struct {
	PdfOutlineTreeNode
	Title  *_dg.PdfObjectString
	Parent *PdfOutlineTreeNode
	Prev   *PdfOutlineTreeNode
	Next   *PdfOutlineTreeNode
	Count  *int64
	Dest   _dg.PdfObject
	A      _dg.PdfObject
	SE     _dg.PdfObject
	C      _dg.PdfObject
	F      _dg.PdfObject
	_fbgf  *_dg.PdfIndirectObject
}

// NewPdfColorspaceFromPdfObject loads a PdfColorspace from a PdfObject.  Returns an error if there is
// a failure in loading.
func NewPdfColorspaceFromPdfObject(obj _dg.PdfObject) (PdfColorspace, error) {
	if obj == nil {
		return nil, nil
	}
	var _cdfc *_dg.PdfIndirectObject
	var _acdg *_dg.PdfObjectName
	var _ecdb *_dg.PdfObjectArray
	if _gafa, _cgag := obj.(*_dg.PdfIndirectObject); _cgag {
		_cdfc = _gafa
	}
	obj = _dg.TraceToDirectObject(obj)
	switch _egbge := obj.(type) {
	case *_dg.PdfObjectArray:
		_ecdb = _egbge
	case *_dg.PdfObjectName:
		_acdg = _egbge
	}
	if _acdg != nil {
		switch *_acdg {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			return NewPdfColorspaceDeviceGray(), nil
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			return NewPdfColorspaceDeviceRGB(), nil
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return NewPdfColorspaceDeviceCMYK(), nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return NewPdfColorspaceSpecialPattern(), nil
		default:
			_ag.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0025\u0073", *_acdg)
			return nil, _dgaa
		}
	}
	if _ecdb != nil && _ecdb.Len() > 0 {
		var _eced _dg.PdfObject = _cdfc
		if _cdfc == nil {
			_eced = _ecdb
		}
		if _ageb, _cgbc := _dg.GetName(_ecdb.Get(0)); _cgbc {
			switch _ageb.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				if _ecdb.Len() == 1 {
					return NewPdfColorspaceDeviceGray(), nil
				}
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				if _ecdb.Len() == 1 {
					return NewPdfColorspaceDeviceRGB(), nil
				}
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _ecdb.Len() == 1 {
					return NewPdfColorspaceDeviceCMYK(), nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079":
				return _geee(_eced)
			case "\u0043\u0061\u006c\u0052\u0047\u0042":
				return _gdea(_eced)
			case "\u004c\u0061\u0062":
				return _dcbd(_eced)
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064":
				return _adddc(_eced)
			case "\u0050a\u0074\u0074\u0065\u0072\u006e":
				return _dffca(_eced)
			case "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return _baac(_eced)
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e":
				return _fbbg(_eced)
			case "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return _fdfee(_eced)
			default:
				_ag.Log.Debug("A\u0072\u0072\u0061\u0079\u0020\u0077i\u0074\u0068\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0061m\u0065:\u0020\u0025\u0073", *_ageb)
			}
		}
	}
	_ag.Log.Debug("\u0050\u0044\u0046\u0020\u0046i\u006c\u0065\u0020\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", obj.String())
	return nil, ErrTypeCheck
}
func (_gaa *PdfReader) newPdfActionJavaScriptFromDict(_bff *_dg.PdfObjectDictionary) (*PdfActionJavaScript, error) {
	return &PdfActionJavaScript{JS: _bff.Get("\u004a\u0053")}, nil
}

// ColorToRGB converts gray -> rgb for a single color component.
func (_eagb *PdfColorspaceDeviceGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_bcbd, _ccagf := color.(*PdfColorDeviceGray)
	if !_ccagf {
		_ag.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006fr\u0020\u006e\u006f\u0074\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0067\u0072\u0061\u0079\u0020\u0025\u0054", color)
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return NewPdfColorDeviceRGB(float64(*_bcbd), float64(*_bcbd), float64(*_bcbd)), nil
}

// ToInteger convert to an integer format.
func (_dadf *PdfColorDeviceCMYK) ToInteger(bits int) [4]uint32 {
	_bfb := _cg.Pow(2, float64(bits)) - 1
	return [4]uint32{uint32(_bfb * _dadf.C()), uint32(_bfb * _dadf.M()), uint32(_bfb * _dadf.Y()), uint32(_bfb * _dadf.K())}
}
func (_feag *PdfField) inherit(_adgb func(*PdfField) bool) (bool, error) {
	_decaf := map[*PdfField]bool{}
	_fcbe := false
	_bdcb := _feag
	for _bdcb != nil {
		if _, _cdeaf := _decaf[_bdcb]; _cdeaf {
			return false, _bf.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		_cgcc := _adgb(_bdcb)
		if _cgcc {
			_fcbe = true
			break
		}
		_decaf[_bdcb] = true
		_bdcb = _bdcb.Parent
	}
	return _fcbe, nil
}

var _ pdfFont = (*pdfFontSimple)(nil)

// SetPdfModifiedDate sets the ModDate attribute of the output PDF.
func SetPdfModifiedDate(modifiedDate _a.Time) {
	_fgefgf.Lock()
	defer _fgefgf.Unlock()
	_efafac = modifiedDate
}

// NewPdfAnnotationMovie returns a new movie annotation.
func NewPdfAnnotationMovie() *PdfAnnotationMovie {
	_fcda := NewPdfAnnotation()
	_dabe := &PdfAnnotationMovie{}
	_dabe.PdfAnnotation = _fcda
	_fcda.SetContext(_dabe)
	return _dabe
}

// NewPdfColorspaceICCBased returns a new ICCBased colorspace object.
func NewPdfColorspaceICCBased(N int) (*PdfColorspaceICCBased, error) {
	_eaad := &PdfColorspaceICCBased{}
	if N != 1 && N != 3 && N != 4 {
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u0031/\u0033\u002f\u0034\u0029")
	}
	_eaad.N = N
	return _eaad, nil
}
func (_bcfd *PdfReader) newPdfPageFromDict(_ggef *_dg.PdfObjectDictionary) (*PdfPage, error) {
	_efedf := NewPdfPage()
	_efedf._bfdge = _ggef
	_efedf._gaed = *_ggef
	_daagd := *_ggef
	_fcgag, _bacae := _daagd.Get("\u0054\u0079\u0070\u0065").(*_dg.PdfObjectName)
	if !_bacae {
		return nil, _bf.New("\u006d\u0069ss\u0069\u006e\u0067/\u0069\u006e\u0076\u0061lid\u0020Pa\u0067\u0065\u0020\u0064\u0069\u0063\u0074io\u006e\u0061\u0072\u0079\u0020\u0054\u0079p\u0065")
	}
	if *_fcgag != "\u0050\u0061\u0067\u0065" {
		return nil, _bf.New("\u0070\u0061\u0067\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u0050a\u0067\u0065")
	}
	if _abega := _daagd.Get("\u0050\u0061\u0072\u0065\u006e\u0074"); _abega != nil {
		_efedf.Parent = _abega
	}
	if _dfgdd := _daagd.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"); _dfgdd != nil {
		_fffdgd, _adebe := _dg.GetString(_dfgdd)
		if !_adebe {
			return nil, _bf.New("\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u004c\u0061\u0073\u0074\u004d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0021=\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		_egcbe, _fdceb := NewPdfDate(_fffdgd.Str())
		if _fdceb != nil {
			return nil, _fdceb
		}
		_efedf.LastModified = &_egcbe
	}
	if _gbag := _daagd.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _gbag != nil && !_dg.IsNullObject(_gbag) {
		_afae, _aeff := _dg.GetDict(_gbag)
		if !_aeff {
			return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _gbag)
		}
		var _dfege error
		_efedf.Resources, _dfege = NewPdfPageResourcesFromDict(_afae)
		if _dfege != nil {
			return nil, _dfege
		}
	} else {
		_gdeag, _afbg := _efedf.getParentResources()
		if _afbg != nil {
			return nil, _afbg
		}
		if _gdeag == nil {
			_gdeag = NewPdfPageResources()
		}
		_efedf.Resources = _gdeag
	}
	if _eaecg := _daagd.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _eaecg != nil {
		_gaec, _fdbed := _dg.GetArray(_eaecg)
		if !_fdbed {
			return nil, _bf.New("\u0070\u0061\u0067\u0065\u0020\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _ggacc error
		_efedf.MediaBox, _ggacc = NewPdfRectangle(*_gaec)
		if _ggacc != nil {
			return nil, _ggacc
		}
	}
	if _eaff := _daagd.Get("\u0043r\u006f\u0070\u0042\u006f\u0078"); _eaff != nil {
		_dadde, _aefgf := _dg.GetArray(_eaff)
		if !_aefgf {
			return nil, _bf.New("\u0070a\u0067\u0065\u0020\u0043r\u006f\u0070\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _beag error
		_efedf.CropBox, _beag = NewPdfRectangle(*_dadde)
		if _beag != nil {
			return nil, _beag
		}
	}
	if _cbea := _daagd.Get("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078"); _cbea != nil {
		_efea, _febbb := _dg.GetArray(_cbea)
		if !_febbb {
			return nil, _bf.New("\u0070\u0061\u0067\u0065\u0020\u0042\u006c\u0065\u0065\u0064\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _bgcdf error
		_efedf.BleedBox, _bgcdf = NewPdfRectangle(*_efea)
		if _bgcdf != nil {
			return nil, _bgcdf
		}
	}
	if _eccdd := _daagd.Get("\u0054r\u0069\u006d\u0042\u006f\u0078"); _eccdd != nil {
		_abee, _bbgdg := _dg.GetArray(_eccdd)
		if !_bbgdg {
			return nil, _bf.New("\u0070a\u0067\u0065\u0020\u0054r\u0069\u006d\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _aecf error
		_efedf.TrimBox, _aecf = NewPdfRectangle(*_abee)
		if _aecf != nil {
			return nil, _aecf
		}
	}
	if _ffee := _daagd.Get("\u0041\u0072\u0074\u0042\u006f\u0078"); _ffee != nil {
		_afff, _eafg := _dg.GetArray(_ffee)
		if !_eafg {
			return nil, _bf.New("\u0070a\u0067\u0065\u0020\u0041\u0072\u0074\u0042\u006f\u0078\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		var _ggfcfa error
		_efedf.ArtBox, _ggfcfa = NewPdfRectangle(*_afff)
		if _ggfcfa != nil {
			return nil, _ggfcfa
		}
	}
	if _ggaff := _daagd.Get("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f"); _ggaff != nil {
		_efedf.BoxColorInfo = _ggaff
	}
	if _gbgffe := _daagd.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _gbgffe != nil {
		_efedf.Contents = _gbgffe
	}
	if _bcgaa := _daagd.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _bcgaa != nil {
		_ddaff, _bbcce := _dg.GetNumberAsInt64(_bcgaa)
		if _bbcce != nil {
			return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0067e\u0020\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_efedf.Rotate = &_ddaff
	}
	if _cdbd := _daagd.Get("\u0047\u0072\u006fu\u0070"); _cdbd != nil {
		_efedf.Group = _cdbd
	}
	if _ceage := _daagd.Get("\u0054\u0068\u0075m\u0062"); _ceage != nil {
		_efedf.Thumb = _ceage
	}
	if _gddcb := _daagd.Get("\u0042"); _gddcb != nil {
		_efedf.B = _gddcb
	}
	if _cdead := _daagd.Get("\u0044\u0075\u0072"); _cdead != nil {
		_efedf.Dur = _cdead
	}
	if _cfcae := _daagd.Get("\u0054\u0072\u0061n\u0073"); _cfcae != nil {
		_efedf.Trans = _cfcae
	}
	if _dacaf := _daagd.Get("\u0041\u0041"); _dacaf != nil {
		_efedf.AA = _dacaf
	}
	if _cdda := _daagd.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _cdda != nil {
		_efedf.Metadata = _cdda
	}
	if _bbbef := _daagd.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o"); _bbbef != nil {
		_efedf.PieceInfo = _bbbef
	}
	if _fegdc := _daagd.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073"); _fegdc != nil {
		_efedf.StructParents = _fegdc
	}
	if _dddfa := _daagd.Get("\u0049\u0044"); _dddfa != nil {
		_efedf.ID = _dddfa
	}
	if _gccbd := _daagd.Get("\u0050\u005a"); _gccbd != nil {
		_efedf.PZ = _gccbd
	}
	if _gafcf := _daagd.Get("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f"); _gafcf != nil {
		_efedf.SeparationInfo = _gafcf
	}
	if _bbefc := _daagd.Get("\u0054\u0061\u0062\u0073"); _bbefc != nil {
		_efedf.Tabs = _bbefc
	}
	if _eedb := _daagd.Get("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064"); _eedb != nil {
		_efedf.TemplateInstantiated = _eedb
	}
	if _dcffe := _daagd.Get("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s"); _dcffe != nil {
		_efedf.PresSteps = _dcffe
	}
	if _bffecc := _daagd.Get("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074"); _bffecc != nil {
		_efedf.UserUnit = _bffecc
	}
	if _eafe := _daagd.Get("\u0056\u0050"); _eafe != nil {
		_efedf.VP = _eafe
	}
	if _cbfcb := _daagd.Get("\u0041\u006e\u006e\u006f\u0074\u0073"); _cbfcb != nil {
		_efedf.Annots = _cbfcb
	}
	_efedf._cbbcc = _bcfd
	return _efedf, nil
}

// NewCompositePdfFontFromTTF loads a composite TTF font. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTF.
func NewCompositePdfFontFromTTF(r _cf.ReadSeeker) (*PdfFont, error) {
	_dcgg, _egedc := _gf.ReadAll(r)
	if _egedc != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _egedc)
		return nil, _egedc
	}
	_cdcb, _egedc := _bbg.TtfParse(_bc.NewReader(_dcgg))
	if _egedc != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067 \u0074\u0074\u0066\u0020\u0066\u006f\u006et\u003a\u0020\u0025\u0076", _egedc)
		return nil, _egedc
	}
	_fcaa := &pdfCIDFontType2{fontCommon: fontCommon{_bcga: "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032"}, CIDToGIDMap: _dg.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079")}
	if len(_cdcb.Widths) <= 0 {
		return nil, _bf.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_abbg := 1000.0 / float64(_cdcb.UnitsPerEm)
	_ceff := _abbg * float64(_cdcb.Widths[0])
	_ecbaf := make(map[rune]int)
	_cadb := make(map[_bbg.GID]int)
	_fbece := _bbg.GID(len(_cdcb.Widths))
	for _aefe, _bfcba := range _cdcb.Chars {
		if _bfcba > _fbece-1 {
			continue
		}
		_ccbgg := int(_abbg * float64(_cdcb.Widths[_bfcba]))
		_ecbaf[_aefe] = _ccbgg
		_cadb[_bfcba] = _ccbgg
	}
	_fcaa._ebgc = _ecbaf
	_fcaa.DW = _dg.MakeInteger(int64(_ceff))
	_gbff := _ggcb(_cadb, uint16(_fbece))
	_fcaa.W = _dg.MakeIndirectObject(_gbff)
	_fbege := _dg.MakeDict()
	_fbege.Set("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067", _dg.MakeString("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"))
	_fbege.Set("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079", _dg.MakeString("\u0041\u0064\u006fb\u0065"))
	_fbege.Set("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074", _dg.MakeInteger(0))
	_fcaa.CIDSystemInfo = _fbege
	_aabdg := &PdfFontDescriptor{FontName: _dg.MakeName(_cdcb.PostScriptName), Ascent: _dg.MakeFloat(_abbg * float64(_cdcb.TypoAscender)), Descent: _dg.MakeFloat(_abbg * float64(_cdcb.TypoDescender)), CapHeight: _dg.MakeFloat(_abbg * float64(_cdcb.CapHeight)), FontBBox: _dg.MakeArrayFromFloats([]float64{_abbg * float64(_cdcb.Xmin), _abbg * float64(_cdcb.Ymin), _abbg * float64(_cdcb.Xmax), _abbg * float64(_cdcb.Ymax)}), ItalicAngle: _dg.MakeFloat(_cdcb.ItalicAngle), MissingWidth: _dg.MakeFloat(_ceff)}
	_gefa, _egedc := _dg.MakeStream(_dcgg, _dg.NewFlateEncoder())
	if _egedc != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _egedc)
		return nil, _egedc
	}
	_gefa.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _dg.MakeInteger(int64(len(_dcgg))))
	_aabdg.FontFile2 = _gefa
	if _cdcb.Bold {
		_aabdg.StemV = _dg.MakeInteger(120)
	} else {
		_aabdg.StemV = _dg.MakeInteger(70)
	}
	_ggbbc := _bbfd
	if _cdcb.IsFixedPitch {
		_ggbbc |= _gafad
	}
	if _cdcb.ItalicAngle != 0 {
		_ggbbc |= _bfeb
	}
	_aabdg.Flags = _dg.MakeInteger(int64(_ggbbc))
	_fcaa._ecbf = _cdcb.PostScriptName
	_fcaa._ccfb = _aabdg
	_faca := pdfFontType0{fontCommon: fontCommon{_bcga: "\u0054\u0079\u0070e\u0030", _ecbf: _cdcb.PostScriptName}, DescendantFont: &PdfFont{_cadf: _fcaa}, Encoding: _dg.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048"), _ggec: _cdcb.NewEncoder()}
	if len(_cdcb.Chars) > 0 {
		_gebfg := make(map[_ff.CharCode]rune, len(_cdcb.Chars))
		for _cafc, _cggeb := range _cdcb.Chars {
			_gegfd := _ff.CharCode(_cggeb)
			if _caaef, _ggccd := _gebfg[_gegfd]; !_ggccd || (_ggccd && _caaef > _cafc) {
				_gebfg[_gegfd] = _cafc
			}
		}
		_faca._ecfb = _ff.NewToUnicodeCMap(_gebfg)
	}
	_cffbf := PdfFont{_cadf: &_faca}
	return &_cffbf, nil
}
func _aaaa(_gbgff *_dg.PdfObjectDictionary) (*PdfFieldChoice, error) {
	_fbab := &PdfFieldChoice{}
	_fbab.Opt, _ = _dg.GetArray(_gbgff.Get("\u004f\u0070\u0074"))
	_fbab.TI, _ = _dg.GetInt(_gbgff.Get("\u0054\u0049"))
	_fbab.I, _ = _dg.GetArray(_gbgff.Get("\u0049"))
	return _fbab, nil
}
func (_edfg *DSS) add(_bgbd *[]*_dg.PdfObjectStream, _cdcf map[string]*_dg.PdfObjectStream, _dgeg [][]byte) ([]*_dg.PdfObjectStream, error) {
	_abcb := make([]*_dg.PdfObjectStream, 0, len(_dgeg))
	for _, _fbbf := range _dgeg {
		_dagf, _dggd := _cbeaf(_fbbf)
		if _dggd != nil {
			return nil, _dggd
		}
		_cefa, _fddg := _cdcf[string(_dagf)]
		if !_fddg {
			_cefa, _dggd = _dg.MakeStream(_fbbf, _dg.NewRawEncoder())
			if _dggd != nil {
				return nil, _dggd
			}
			_cdcf[string(_dagf)] = _cefa
			*_bgbd = append(*_bgbd, _cefa)
		}
		_abcb = append(_abcb, _cefa)
	}
	return _abcb, nil
}

// PdfActionImportData represents a importData action.
type PdfActionImportData struct {
	*PdfAction
	F *PdfFilespec
}

func _acbbg(_acdgd *_dg.PdfObjectDictionary) {
	_cefb, _gedg := _dg.GetArray(_acdgd.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
	_dbcca, _edga := _dg.GetIntVal(_acdgd.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
	_dgfe, _edfgd := _dg.GetIntVal(_acdgd.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
	if _gedg && _edga && _edfgd {
		_fceaa := _cefb.Len()
		if _fceaa != _dgfe-_dbcca+1 {
			_ag.Log.Debug("\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076\u002c\u0020\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u003a\u0020\u0025\u0076", _fceaa, _dgfe)
			_fffe := _dg.PdfObjectInteger(_dbcca + _fceaa - 1)
			_acdgd.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", &_fffe)
		}
	}
}

// PdfColorspaceDeviceRGB represents an RGB colorspace.
type PdfColorspaceDeviceRGB struct{}

// RepairAcroForm attempts to rebuild the AcroForm fields using the widget
// annotations present in the document pages. Pass nil for the opts parameter
// in order to use the default options.
// NOTE: Currently, the opts parameter is declared in order to enable adding
// future options, but passing nil will always result in the default options
// being used.
func (_aecd *PdfReader) RepairAcroForm(opts *AcroFormRepairOptions) error {
	var _aceg []*PdfField
	_ddeeb := map[*_dg.PdfIndirectObject]struct{}{}
	for _, _fbbbd := range _aecd.PageList {
		_eabbe, _abebe := _fbbbd.GetAnnotations()
		if _abebe != nil {
			return _abebe
		}
		for _, _cdcfg := range _eabbe {
			var _ebcbb *PdfField
			switch _fdced := _cdcfg.GetContext().(type) {
			case *PdfAnnotationWidget:
				if _fdced._cgg != nil {
					_ebcbb = _fdced._cgg
					break
				}
				if _gaege, _eedgg := _dg.GetIndirect(_fdced.Parent); _eedgg {
					_ebcbb, _abebe = _aecd.newPdfFieldFromIndirectObject(_gaege, nil)
					if _abebe == nil {
						break
					}
					_ag.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _gaege, _abebe)
				}
				if _fdced._cdf != nil {
					_ebcbb, _abebe = _aecd.newPdfFieldFromIndirectObject(_fdced._cdf, nil)
					if _abebe == nil {
						break
					}
					_ag.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _fdced._cdf, _abebe)
				}
			}
			if _ebcbb == nil {
				continue
			}
			if _, _cagce := _ddeeb[_ebcbb._egce]; _cagce {
				continue
			}
			_ddeeb[_ebcbb._egce] = struct{}{}
			_aceg = append(_aceg, _ebcbb)
		}
	}
	if len(_aceg) == 0 {
		return nil
	}
	if _aecd.AcroForm == nil {
		_aecd.AcroForm = NewPdfAcroForm()
	}
	_aecd.AcroForm.Fields = &_aceg
	return nil
}
func (_debeg *PdfReader) buildPageList(_gdfea *_dg.PdfIndirectObject, _eegf *_dg.PdfIndirectObject, _bcaa map[_dg.PdfObject]struct{}) error {
	if _gdfea == nil {
		return nil
	}
	if _, _cefgb := _bcaa[_gdfea]; _cefgb {
		_ag.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _gdfea.ObjectNumber)
		return nil
	}
	_bcaa[_gdfea] = struct{}{}
	_fcefd, _dgbdf := _gdfea.PdfObject.(*_dg.PdfObjectDictionary)
	if !_dgbdf {
		return _bf.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_eabbae, _dgbdf := (*_fcefd).Get("\u0054\u0079\u0070\u0065").(*_dg.PdfObjectName)
	if !_dgbdf {
		if _fcefd.Get("\u004b\u0069\u0064\u0073") == nil {
			return _bf.New("\u006e\u006f\u0064\u0065 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0054\u0079p\u0065 \u0028\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0029")
		}
		_ag.Log.Debug("ER\u0052\u004fR\u003a\u0020\u006e\u006f\u0064\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067\u0020\u0054\u0079\u0070\u0065\u002c\u0020\u0062\u0075\u0074\u0020\u0068\u0061\u0073\u0020\u004b\u0069\u0064\u0073\u002e\u0020\u0041\u0073\u0073u\u006di\u006e\u0067\u0020\u0050\u0061\u0067\u0065\u0073 \u006eo\u0064\u0065.")
		_eabbae = _dg.MakeName("\u0050\u0061\u0067e\u0073")
		_fcefd.Set("\u0054\u0079\u0070\u0065", _eabbae)
	}
	_ag.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0050a\u0067\u0065\u004c\u0069\u0073\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0074y\u0070\u0065\u003a\u0020\u0025\u0073\u0020(\u0025\u002b\u0076\u0029", *_eabbae, _gdfea)
	if *_eabbae == "\u0050\u0061\u0067\u0065" {
		_caaaa, _edbgf := _debeg.newPdfPageFromDict(_fcefd)
		if _edbgf != nil {
			return _edbgf
		}
		_caaaa.setContainer(_gdfea)
		if _eegf != nil {
			_fcefd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _eegf)
		}
		_debeg._daddd = append(_debeg._daddd, _gdfea)
		_debeg.PageList = append(_debeg.PageList, _caaaa)
		return nil
	}
	if *_eabbae != "\u0050\u0061\u0067e\u0073" {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0063\u006fnt\u0065n\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u0021\u0020\u0028\u0025\u0073\u0029", _eabbae)
		return _bf.New("\u0074\u0061\u0062\u006c\u0065\u0020o\u0066\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073 \u006fb\u006a\u0065\u0063\u0074")
	}
	if _eegf != nil {
		_fcefd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _eegf)
	}
	if !_debeg._dadcef {
		_cbfde := _debeg.traverseObjectData(_gdfea)
		if _cbfde != nil {
			return _cbfde
		}
	}
	_egfb, _bcbbb := _debeg._baad.Resolve(_fcefd.Get("\u004b\u0069\u0064\u0073"))
	if _bcbbb != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u004b\u0069\u0064\u0073\u0020\u006fb\u006a\u0065\u0063\u0074")
		return _bcbbb
	}
	var _gfeda *_dg.PdfObjectArray
	_gfeda, _dgbdf = _egfb.(*_dg.PdfObjectArray)
	if !_dgbdf {
		_abgfa, _deaeb := _egfb.(*_dg.PdfIndirectObject)
		if !_deaeb {
			return _bf.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_gfeda, _dgbdf = _abgfa.PdfObject.(*_dg.PdfObjectArray)
		if !_dgbdf {
			return _bf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u0069\u006ed\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
	}
	_ag.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _gfeda)
	for _aacg, _bcbea := range _gfeda.Elements() {
		_egddb, _cddd := _dg.GetIndirect(_bcbea)
		if !_cddd {
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 \u006f\u0062\u006a\u0065\u0063t\u0020\u002d \u0028\u0025\u0073\u0029", _egddb)
			return _bf.New("\u0070a\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_gfeda.Set(_aacg, _egddb)
		_bcbbb = _debeg.buildPageList(_egddb, _gdfea, _bcaa)
		if _bcbbb != nil {
			return _bcbbb
		}
	}
	return nil
}

// GetPatternByName gets the pattern specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_acdbb *PdfPageResources) GetPatternByName(keyName _dg.PdfObjectName) (*PdfPattern, bool) {
	if _acdbb.Pattern == nil {
		return nil, false
	}
	_afbe, _efedb := _dg.TraceToDirectObject(_acdbb.Pattern).(*_dg.PdfObjectDictionary)
	if !_efedb {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074t\u0065\u0072\u006e\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _acdbb.Pattern)
		return nil, false
	}
	if _debbc := _afbe.Get(keyName); _debbc != nil {
		_cfgfe, _ccedc := _agea(_debbc)
		if _ccedc != nil {
			_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0070\u0061\u0074\u0074\u0065\u0072n\u003a\u0020\u0025\u0076", _ccedc)
			return nil, false
		}
		return _cfgfe, true
	}
	return nil, false
}

// GetOCProperties returns the optional content properties PdfObject.
func (_dcaaf *PdfReader) GetOCProperties() (_dg.PdfObject, error) {
	_cbee := _dcaaf._gccfb
	_efee := _cbee.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	_efee = _dg.ResolveReference(_efee)
	if !_dcaaf._dadcef {
		_gdebg := _dcaaf.traverseObjectData(_efee)
		if _gdebg != nil {
			return nil, _gdebg
		}
	}
	return _efee, nil
}

// PdfFilespec represents a file specification which can either refer to an external or embedded file.
type PdfFilespec struct {
	Type  _dg.PdfObject
	FS    _dg.PdfObject
	F     _dg.PdfObject
	UF    _dg.PdfObject
	DOS   _dg.PdfObject
	Mac   _dg.PdfObject
	Unix  _dg.PdfObject
	ID    _dg.PdfObject
	V     _dg.PdfObject
	EF    _dg.PdfObject
	RF    _dg.PdfObject
	Desc  _dg.PdfObject
	CI    _dg.PdfObject
	_bcbe _dg.PdfObject
}

// PdfField contains the common attributes of a form field. The context object contains the specific field data
// which can represent a button, text, choice or signature.
// The PdfField is typically not used directly, but is encapsulated by the more specific field types such as
// PdfFieldButton etc (i.e. the context attribute).
type PdfField struct {
	_bdfg        PdfModel
	_egce        *_dg.PdfIndirectObject
	Parent       *PdfField
	Annotations  []*PdfAnnotationWidget
	Kids         []*PdfField
	FT           *_dg.PdfObjectName
	T            *_dg.PdfObjectString
	TU           *_dg.PdfObjectString
	TM           *_dg.PdfObjectString
	Ff           *_dg.PdfObjectInteger
	V            _dg.PdfObject
	DV           _dg.PdfObject
	AA           _dg.PdfObject
	VariableText *VariableText
}

// NewPdfFilespec returns an initialized generic PDF filespec model.
func NewPdfFilespec() *PdfFilespec {
	_ggcef := &PdfFilespec{}
	_ggcef._bcbe = _dg.MakeIndirectObject(_dg.MakeDict())
	return _ggcef
}

// GetTrailer returns the PDF's trailer dictionary.
func (_agdf *PdfReader) GetTrailer() (*_dg.PdfObjectDictionary, error) {
	_aagf := _agdf._baad.GetTrailer()
	if _aagf == nil {
		return nil, _bf.New("\u0074r\u0061i\u006c\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	return _aagf, nil
}

// PdfWriter handles outputing PDF content.
type PdfWriter struct {
	_fadee         *_dg.PdfIndirectObject
	_gbgb          *_dg.PdfIndirectObject
	_fegdf         map[_dg.PdfObject]struct{}
	_agaba         []_dg.PdfObject
	_fdbfa         map[_dg.PdfObject]struct{}
	_fcbf          []*_dg.PdfIndirectObject
	_cedcg         *PdfOutlineTreeNode
	_ecdf          *_dg.PdfObjectDictionary
	_cacg          []_dg.PdfObject
	_efbfa         *_dg.PdfIndirectObject
	_bddfa         *_ba.Writer
	_fbbfc         int64
	_ffefc         error
	_fadcg         *_dg.PdfCrypt
	_fbfbf         *_dg.PdfObjectDictionary
	_acdag         *_dg.PdfIndirectObject
	_agbeg         *_dg.PdfObjectArray
	_efacd         _dg.Version
	_eaacb         *bool
	_ccgade        map[_dg.PdfObject][]*_dg.PdfObjectDictionary
	_geabe         *PdfAcroForm
	_egeac         Optimizer
	_afafd         StandardApplier
	_fffge         map[int]crossReference
	_bgggdg        int64
	ObjNumOffset   int
	_bbac          bool
	_cfcdcb        _dg.XrefTable
	_eege          int64
	_fafab         int64
	_ecabb         map[_dg.PdfObject]int64
	_cffaa         map[_dg.PdfObject]struct{}
	_gdfbg         string
	_ceab          []*PdfOutputIntent
	_affea         bool
	_cedaf, _ceaab string
}

// PdfAnnotationInk represents Ink annotations.
// (Section 12.5.6.13).
type PdfAnnotationInk struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	InkList _dg.PdfObject
	BS      _dg.PdfObject
}

// ColorToRGB converts a DeviceN color to an RGB color.
func (_bfabc *PdfColorspaceDeviceN) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _bfabc.AlternateSpace == nil {
		return nil, _bf.New("\u0044\u0065\u0076\u0069\u0063\u0065N\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _bfabc.AlternateSpace.ColorToRGB(color)
}

// Insert adds a top level outline item in the outline,
// at the specified index.
func (_acaa *Outline) Insert(index uint, item *OutlineItem) {
	_cbgf := uint(len(_acaa.Entries))
	if index > _cbgf {
		index = _cbgf
	}
	_acaa.Entries = append(_acaa.Entries[:index], append([]*OutlineItem{item}, _acaa.Entries[index:]...)...)
}

// NewXObjectImage returns a new XObjectImage.
func NewXObjectImage() *XObjectImage {
	_cebga := &XObjectImage{}
	_bbega := &_dg.PdfObjectStream{}
	_bbega.PdfObjectDictionary = _dg.MakeDict()
	_cebga._abfb = _bbega
	return _cebga
}
func _ggdcc(_fefe *_dg.PdfObjectDictionary) (*PdfShadingType1, error) {
	_fcedb := PdfShadingType1{}
	if _ageg := _fefe.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _ageg != nil {
		_ageg = _dg.TraceToDirectObject(_ageg)
		_cedge, _fbbge := _ageg.(*_dg.PdfObjectArray)
		if !_fbbge {
			_ag.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ageg)
			return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fcedb.Domain = _cedge
	}
	if _bfdeg := _fefe.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _bfdeg != nil {
		_bfdeg = _dg.TraceToDirectObject(_bfdeg)
		_ffce, _cceec := _bfdeg.(*_dg.PdfObjectArray)
		if !_cceec {
			_ag.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _bfdeg)
			return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fcedb.Matrix = _ffce
	}
	_gecde := _fefe.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _gecde == nil {
		_ag.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_fcedb.Function = []PdfFunction{}
	if _dfag, _cgfgb := _gecde.(*_dg.PdfObjectArray); _cgfgb {
		for _, _ebebbc := range _dfag.Elements() {
			_cbfbe, _fbgc := _agec(_ebebbc)
			if _fbgc != nil {
				_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _fbgc)
				return nil, _fbgc
			}
			_fcedb.Function = append(_fcedb.Function, _cbfbe)
		}
	} else {
		_bcdc, _dgfaa := _agec(_gecde)
		if _dgfaa != nil {
			_ag.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dgfaa)
			return nil, _dgfaa
		}
		_fcedb.Function = append(_fcedb.Function, _bcdc)
	}
	return &_fcedb, nil
}

// PdfAnnotationPrinterMark represents PrinterMark annotations.
// (Section 12.5.6.20).
type PdfAnnotationPrinterMark struct {
	*PdfAnnotation
	MN _dg.PdfObject
}

// PdfColorDeviceRGB represents a color in DeviceRGB colorspace with R, G, B components, where component is
// defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceRGB [3]float64

// PdfAnnotationStamp represents Stamp annotations.
// (Section 12.5.6.12).
type PdfAnnotationStamp struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Name _dg.PdfObject
}

func (_gadcg *PdfReader) resolveReference(_efdab *_dg.PdfObjectReference) (_dg.PdfObject, bool, error) {
	_gfcbd, _efdabd := _gadcg._baad.ObjCache[int(_efdab.ObjectNumber)]
	if !_efdabd {
		_ag.Log.Trace("R\u0065\u0061\u0064\u0065r \u004co\u006f\u006b\u0075\u0070\u0020r\u0065\u0066\u003a\u0020\u0025\u0073", _efdab)
		_fbaee, _gfgbc := _gadcg._baad.LookupByReference(*_efdab)
		if _gfgbc != nil {
			return nil, false, _gfgbc
		}
		_gadcg._baad.ObjCache[int(_efdab.ObjectNumber)] = _fbaee
		return _fbaee, false, nil
	}
	return _gfcbd, true, nil
}
func _eece(_abbed StdFontName) (pdfFontSimple, error) {
	_beda, _fbef := _bbg.NewStdFontByName(_abbed)
	if !_fbef {
		return pdfFontSimple{}, ErrFontNotSupported
	}
	_dggda := _gecgd(_beda)
	return _dggda, nil
}

// NewStandard14FontMustCompile returns the standard 14 font named `basefont` as a *PdfFont.
// If `basefont` is one of the 14 Standard14Font values defined above then NewStandard14FontMustCompile
// is guaranteed to succeed.
func NewStandard14FontMustCompile(basefont StdFontName) *PdfFont {
	_fggd, _adegg := NewStandard14Font(basefont)
	if _adegg != nil {
		panic(_b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0031\u0034\u0046\u006f\u006e\u0074 \u0025\u0023\u0071", basefont))
	}
	return _fggd
}
func _aedcd() string { _fgefgf.Lock(); defer _fgefgf.Unlock(); return _adgf }

// PdfAnnotationText represents Text annotations.
// (Section 12.5.6.4 p. 402).
type PdfAnnotationText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Open       _dg.PdfObject
	Name       _dg.PdfObject
	State      _dg.PdfObject
	StateModel _dg.PdfObject
}

func (_dceab *PdfWriter) updateObjectNumbers() {
	_fedee := _dceab.ObjNumOffset
	_dgbb := 0
	for _, _fdggd := range _dceab._agaba {
		_aeged := int64(_dgbb + 1 + _fedee)
		_bbaaf := true
		if _dceab._bbac {
			if _ecfg, _cgccc := _dceab._ecabb[_fdggd]; _cgccc {
				_aeged = _ecfg
				_bbaaf = false
			}
		}
		switch _gdfbf := _fdggd.(type) {
		case *_dg.PdfIndirectObject:
			_gdfbf.ObjectNumber = _aeged
			_gdfbf.GenerationNumber = 0
		case *_dg.PdfObjectStream:
			_gdfbf.ObjectNumber = _aeged
			_gdfbf.GenerationNumber = 0
		case *_dg.PdfObjectStreams:
			_gdfbf.ObjectNumber = _aeged
			_gdfbf.GenerationNumber = 0
		default:
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u0020%\u0054\u0020\u002d\u0020\u0073\u006b\u0069p\u0070\u0069\u006e\u0067", _gdfbf)
			continue
		}
		if _bbaaf {
			_dgbb++
		}
	}
	_cbcbd := func(_fbbdg _dg.PdfObject) int64 {
		switch _aecbfa := _fbbdg.(type) {
		case *_dg.PdfIndirectObject:
			return _aecbfa.ObjectNumber
		case *_dg.PdfObjectStream:
			return _aecbfa.ObjectNumber
		case *_dg.PdfObjectStreams:
			return _aecbfa.ObjectNumber
		}
		return 0
	}
	_gc.SliceStable(_dceab._agaba, func(_gcecc, _defbc int) bool { return _cbcbd(_dceab._agaba[_gcecc]) < _cbcbd(_dceab._agaba[_defbc]) })
}

// IsRadio returns true if the button field represents a radio button, false otherwise.
func (_ccefe *PdfFieldButton) IsRadio() bool { return _ccefe.GetType() == ButtonTypeRadio }

// NewOutlineBookmark returns an initialized PdfOutlineItem for a given bookmark title and page.
func NewOutlineBookmark(title string, page *_dg.PdfIndirectObject) *PdfOutlineItem {
	_abade := PdfOutlineItem{}
	_abade._baddf = &_abade
	_abade.Title = _dg.MakeString(title)
	_cedg := _dg.MakeArray()
	_cedg.Append(page)
	_cedg.Append(_dg.MakeName("\u0046\u0069\u0074"))
	_abade.Dest = _cedg
	return &_abade
}
func (_ecda *PdfReader) loadAnnotations(_dabfe _dg.PdfObject) ([]*PdfAnnotation, error) {
	_bgccf, _gfabe := _dg.GetArray(_dabfe)
	if !_gfabe {
		return nil, _b.Errorf("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	var _eaee []*PdfAnnotation
	for _, _beagg := range _bgccf.Elements() {
		_beagg = _dg.ResolveReference(_beagg)
		if _, _addc := _beagg.(*_dg.PdfObjectNull); _addc {
			continue
		}
		_fabf, _bdebb := _beagg.(*_dg.PdfObjectDictionary)
		_bbbcd, _ffagg := _beagg.(*_dg.PdfIndirectObject)
		if _bdebb {
			_bbbcd = &_dg.PdfIndirectObject{}
			_bbbcd.PdfObject = _fabf
		} else {
			if !_ffagg {
				return nil, _b.Errorf("\u0061\u006eno\u0074\u0061\u0074i\u006f\u006e\u0020\u006eot \u0069n \u0061\u006e\u0020\u0069\u006e\u0064\u0069re\u0063\u0074\u0020\u006f\u0062\u006a\u0065c\u0074")
			}
		}
		_agcfb, _cegb := _ecda.newPdfAnnotationFromIndirectObject(_bbbcd)
		if _cegb != nil {
			return nil, _cegb
		}
		switch _accf := _agcfb.GetContext().(type) {
		case *PdfAnnotationWidget:
			for _, _gfdcaa := range _ecda.AcroForm.AllFields() {
				if _gfdcaa._egce == _accf.Parent {
					_accf._cgg = _gfdcaa
					break
				}
			}
		}
		if _agcfb != nil {
			_eaee = append(_eaee, _agcfb)
		}
	}
	return _eaee, nil
}

// GetContainingPdfObject returns the page as a dictionary within an PdfIndirectObject.
func (_befda *PdfPage) GetContainingPdfObject() _dg.PdfObject { return _befda._cggbe }

// GetDescent returns the Descent of the font `descriptor`.
func (_ffaca *PdfFontDescriptor) GetDescent() (float64, error) {
	return _dg.GetNumberAsFloat(_ffaca.Descent)
}

// ImageToRGB converts CalRGB colorspace image to RGB and returns the result.
func (_cdgg *PdfColorspaceCalRGB) ImageToRGB(img Image) (Image, error) {
	_aeeg := _fcd.NewReader(img.getBase())
	_acge := _fc.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_dcgc := _fcd.NewWriter(_acge)
	_beae := _cg.Pow(2, float64(img.BitsPerComponent)) - 1
	_gacg := make([]uint32, 3)
	var (
		_gdefa                                    error
		_beff, _ggbe, _fgba, _ccdeg, _eaec, _dddb float64
	)
	for {
		_gdefa = _aeeg.ReadSamples(_gacg)
		if _gdefa == _cf.EOF {
			break
		} else if _gdefa != nil {
			return img, _gdefa
		}
		_beff = float64(_gacg[0]) / _beae
		_ggbe = float64(_gacg[1]) / _beae
		_fgba = float64(_gacg[2]) / _beae
		_ccdeg = _cdgg.Matrix[0]*_cg.Pow(_beff, _cdgg.Gamma[0]) + _cdgg.Matrix[3]*_cg.Pow(_ggbe, _cdgg.Gamma[1]) + _cdgg.Matrix[6]*_cg.Pow(_fgba, _cdgg.Gamma[2])
		_eaec = _cdgg.Matrix[1]*_cg.Pow(_beff, _cdgg.Gamma[0]) + _cdgg.Matrix[4]*_cg.Pow(_ggbe, _cdgg.Gamma[1]) + _cdgg.Matrix[7]*_cg.Pow(_fgba, _cdgg.Gamma[2])
		_dddb = _cdgg.Matrix[2]*_cg.Pow(_beff, _cdgg.Gamma[0]) + _cdgg.Matrix[5]*_cg.Pow(_ggbe, _cdgg.Gamma[1]) + _cdgg.Matrix[8]*_cg.Pow(_fgba, _cdgg.Gamma[2])
		_beff = 3.240479*_ccdeg + -1.537150*_eaec + -0.498535*_dddb
		_ggbe = -0.969256*_ccdeg + 1.875992*_eaec + 0.041556*_dddb
		_fgba = 0.055648*_ccdeg + -0.204043*_eaec + 1.057311*_dddb
		_beff = _cg.Min(_cg.Max(_beff, 0), 1.0)
		_ggbe = _cg.Min(_cg.Max(_ggbe, 0), 1.0)
		_fgba = _cg.Min(_cg.Max(_fgba, 0), 1.0)
		_gacg[0] = uint32(_beff * _beae)
		_gacg[1] = uint32(_ggbe * _beae)
		_gacg[2] = uint32(_fgba * _beae)
		if _gdefa = _dcgc.WriteSamples(_gacg); _gdefa != nil {
			return img, _gdefa
		}
	}
	return _edcf(&_acge), nil
}

var _ pdfFont = (*pdfCIDFontType2)(nil)

func (_fgfgg *PdfWriter) writeOutputIntents() error {
	if len(_fgfgg._ceab) == 0 {
		return nil
	}
	_cbdgg := make([]_dg.PdfObject, len(_fgfgg._ceab))
	for _dbage, _fccece := range _fgfgg._ceab {
		_afecd := _fccece.ToPdfObject()
		_cbdgg[_dbage] = _dg.MakeIndirectObject(_afecd)
	}
	_dabfec := _dg.MakeIndirectObject(_dg.MakeArray(_cbdgg...))
	_fgfgg._ecdf.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _dabfec)
	if _bbccg := _fgfgg.addObjects(_dabfec); _bbccg != nil {
		return _bbccg
	}
	return nil
}

// Add appends a top level outline item to the outline.
func (_caecd *Outline) Add(item *OutlineItem) { _caecd.Entries = append(_caecd.Entries, item) }

// ToPdfObject implements interface PdfModel.
func (_cab *PdfActionURI) ToPdfObject() _dg.PdfObject {
	_cab.PdfAction.ToPdfObject()
	_fad := _cab._cbd
	_bda := _fad.PdfObject.(*_dg.PdfObjectDictionary)
	_bda.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeURI)))
	_bda.SetIfNotNil("\u0055\u0052\u0049", _cab.URI)
	_bda.SetIfNotNil("\u0049\u0073\u004da\u0070", _cab.IsMap)
	return _fad
}

// C returns the value of the cyan component of the color.
func (_fgcfa *PdfColorDeviceCMYK) C() float64 { return _fgcfa[0] }
func _fdgf(_bgcgd _dg.PdfObject) (*PdfFunctionType2, error) {
	_gfcee := &PdfFunctionType2{}
	var _agggcd *_dg.PdfObjectDictionary
	if _cfcgb, _bege := _bgcgd.(*_dg.PdfIndirectObject); _bege {
		_geggf, _fege := _cfcgb.PdfObject.(*_dg.PdfObjectDictionary)
		if !_fege {
			return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_gfcee._acfde = _cfcgb
		_agggcd = _geggf
	} else if _adbec, _fcdbd := _bgcgd.(*_dg.PdfObjectDictionary); _fcdbd {
		_agggcd = _adbec
	} else {
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_ag.Log.Trace("\u0046U\u004e\u0043\u0032\u003a\u0020\u0025s", _agggcd.String())
	_fgff, _cace := _dg.TraceToDirectObject(_agggcd.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_dg.PdfObjectArray)
	if !_cace {
		_ag.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _bf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _fgff.Len() < 0 || _fgff.Len()%2 != 0 {
		_ag.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u0072\u0061\u006e\u0067e\u0020\u0069\u006e\u0076al\u0069\u0064")
		return nil, _bf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_gddb, _gaea := _fgff.ToFloat64Array()
	if _gaea != nil {
		return nil, _gaea
	}
	_gfcee.Domain = _gddb
	_fgff, _cace = _dg.TraceToDirectObject(_agggcd.Get("\u0052\u0061\u006eg\u0065")).(*_dg.PdfObjectArray)
	if _cace {
		if _fgff.Len() < 0 || _fgff.Len()%2 != 0 {
			return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_aeeeb, _ebcge := _fgff.ToFloat64Array()
		if _ebcge != nil {
			return nil, _ebcge
		}
		_gfcee.Range = _aeeeb
	}
	_fgff, _cace = _dg.TraceToDirectObject(_agggcd.Get("\u0043\u0030")).(*_dg.PdfObjectArray)
	if _cace {
		_dcfec, _efgbd := _fgff.ToFloat64Array()
		if _efgbd != nil {
			return nil, _efgbd
		}
		_gfcee.C0 = _dcfec
	}
	_fgff, _cace = _dg.TraceToDirectObject(_agggcd.Get("\u0043\u0031")).(*_dg.PdfObjectArray)
	if _cace {
		_gdfcb, _cgcbdg := _fgff.ToFloat64Array()
		if _cgcbdg != nil {
			return nil, _cgcbdg
		}
		_gfcee.C1 = _gdfcb
	}
	if len(_gfcee.C0) != len(_gfcee.C1) {
		_ag.Log.Error("\u0043\u0030\u0020\u0061nd\u0020\u0043\u0031\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069n\u0067")
		return nil, _dg.ErrRangeError
	}
	N, _gaea := _dg.GetNumberAsFloat(_dg.TraceToDirectObject(_agggcd.Get("\u004e")))
	if _gaea != nil {
		_ag.Log.Error("\u004e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020o\u0072\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u002c\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _agggcd.String())
		return nil, _gaea
	}
	_gfcee.N = N
	return _gfcee, nil
}

// ApplyStandard is used to apply changes required on the document to match the rules required by the input standard.
// The writer's content would be changed after all the document parts are already established during the Write method.
// A good example of the StandardApplier could be a PDF/A Profile (i.e.: pdfa.Profile1A). In such a case PdfWriter would
// set up all rules required by that Profile.
func (_bbefgf *PdfWriter) ApplyStandard(optimizer StandardApplier) { _bbefgf._afafd = optimizer }

// GetRuneMetrics returns the character metrics for the rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_bdagc pdfFontSimple) GetRuneMetrics(r rune) (_bbg.CharMetrics, bool) {
	if _bdagc._bfdee != nil {
		_eabd, _gaag := _bdagc._bfdee.Read(r)
		if _gaag {
			return _eabd, true
		}
	}
	_bece := _bdagc.Encoder()
	if _bece == nil {
		_ag.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _bdagc)
		return _bbg.CharMetrics{}, false
	}
	_ebeg, _cfcdc := _bece.RuneToCharcode(r)
	if !_cfcdc {
		if r != ' ' {
			_ag.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _bdagc)
		}
		return _bbg.CharMetrics{}, false
	}
	_adafb, _fadg := _bdagc.GetCharMetrics(_ebeg)
	return _adafb, _fadg
}

// B returns the value of the B component of the color.
func (_ecega *PdfColorCalRGB) B() float64 { return _ecega[1] }

// PdfAnnotationPopup represents Popup annotations.
// (Section 12.5.6.14).
type PdfAnnotationPopup struct {
	*PdfAnnotation
	Parent _dg.PdfObject
	Open   _dg.PdfObject
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_dfebg *PdfShadingPatternType3) ToPdfObject() _dg.PdfObject {
	_dfebg.PdfPattern.ToPdfObject()
	_feceg := _dfebg.getDict()
	if _dfebg.Shading != nil {
		_feceg.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _dfebg.Shading.ToPdfObject())
	}
	if _dfebg.Matrix != nil {
		_feceg.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _dfebg.Matrix)
	}
	if _dfebg.ExtGState != nil {
		_feceg.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _dfebg.ExtGState)
	}
	return _dfebg._eacce
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_egfdde *PdfShadingType3) ToPdfObject() _dg.PdfObject {
	_egfdde.PdfShading.ToPdfObject()
	_becce, _ffeea := _egfdde.getShadingDict()
	if _ffeea != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _egfdde.Coords != nil {
		_becce.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _egfdde.Coords)
	}
	if _egfdde.Domain != nil {
		_becce.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _egfdde.Domain)
	}
	if _egfdde.Function != nil {
		if len(_egfdde.Function) == 1 {
			_becce.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _egfdde.Function[0].ToPdfObject())
		} else {
			_gddea := _dg.MakeArray()
			for _, _cfcge := range _egfdde.Function {
				_gddea.Append(_cfcge.ToPdfObject())
			}
			_becce.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _gddea)
		}
	}
	if _egfdde.Extend != nil {
		_becce.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _egfdde.Extend)
	}
	return _egfdde._bcfbg
}
func (_abdc *PdfReader) newPdfAnnotationCircleFromDict(_gfag *_dg.PdfObjectDictionary) (*PdfAnnotationCircle, error) {
	_bcb := PdfAnnotationCircle{}
	_fdcdc, _gfcf := _abdc.newPdfAnnotationMarkupFromDict(_gfag)
	if _gfcf != nil {
		return nil, _gfcf
	}
	_bcb.PdfAnnotationMarkup = _fdcdc
	_bcb.BS = _gfag.Get("\u0042\u0053")
	_bcb.IC = _gfag.Get("\u0049\u0043")
	_bcb.BE = _gfag.Get("\u0042\u0045")
	_bcb.RD = _gfag.Get("\u0052\u0044")
	return &_bcb, nil
}

// GenerateXObjectName generates an unused XObject name that can be used for
// adding new XObjects. Uses format XObj1, XObj2, ...
func (_gbgabg *PdfPageResources) GenerateXObjectName() _dg.PdfObjectName {
	_effg := 1
	for {
		_cegf := _dg.MakeName(_b.Sprintf("\u0058\u004f\u0062\u006a\u0025\u0064", _effg))
		if !_gbgabg.HasXObjectByName(*_cegf) {
			return *_cegf
		}
		_effg++
	}
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_efeae *PdfShadingType2) ToPdfObject() _dg.PdfObject {
	_efeae.PdfShading.ToPdfObject()
	_bgcbaa, _cdabb := _efeae.getShadingDict()
	if _cdabb != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _bgcbaa == nil {
		_ag.Log.Error("\u0053\u0068\u0061\u0064in\u0067\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		return nil
	}
	if _efeae.Coords != nil {
		_bgcbaa.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _efeae.Coords)
	}
	if _efeae.Domain != nil {
		_bgcbaa.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _efeae.Domain)
	}
	if _efeae.Function != nil {
		if len(_efeae.Function) == 1 {
			_bgcbaa.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _efeae.Function[0].ToPdfObject())
		} else {
			_cdadb := _dg.MakeArray()
			for _, _ddfdf := range _efeae.Function {
				_cdadb.Append(_ddfdf.ToPdfObject())
			}
			_bgcbaa.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _cdadb)
		}
	}
	if _efeae.Extend != nil {
		_bgcbaa.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _efeae.Extend)
	}
	return _efeae._bcfbg
}

// OutlineDest represents the destination of an outline item.
// It holds the page and the position on the page an outline item points to.
type OutlineDest struct {
	PageObj *_dg.PdfIndirectObject `json:"-"`
	Page    int64                  `json:"page"`
	Mode    string                 `json:"mode"`
	X       float64                `json:"x"`
	Y       float64                `json:"y"`
	Zoom    float64                `json:"zoom"`
}

// GetContext returns a reference to the subshading entry as represented by PdfShadingType1-7.
func (_ffded *PdfShading) GetContext() PdfModel { return _ffded._eeddb }

// GetDocMDPPermission returns the DocMDP level of the restrictions
func (_eadgb *PdfSignature) GetDocMDPPermission() (_ecb.DocMDPPermission, bool) {
	for _, _fgbcd := range _eadgb.Reference.Elements() {
		if _abfc, _dcgb := _dg.GetDict(_fgbcd); _dcgb {
			if _cgefa, _ddbfe := _dg.GetNameVal(_abfc.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _ddbfe && _cgefa == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				if _aefba, _dgfeb := _dg.GetDict(_abfc.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _dgfeb {
					if P, _dade := _dg.GetIntVal(_aefba.Get("\u0050")); _dade {
						return _ecb.DocMDPPermission(P), true
					}
				}
			}
		}
	}
	return 0, false
}

// PdfActionURI represents an URI action.
type PdfActionURI struct {
	*PdfAction
	URI   _dg.PdfObject
	IsMap _dg.PdfObject
}

// GetAscent returns the Ascent of the font `descriptor`.
func (_ebeb *PdfFontDescriptor) GetAscent() (float64, error) {
	return _dg.GetNumberAsFloat(_ebeb.Ascent)
}
func _cbeaf(_ebaea []byte) ([]byte, error) {
	_dbeaa := _fb.New()
	if _, _fdeggb := _cf.Copy(_dbeaa, _bc.NewReader(_ebaea)); _fdeggb != nil {
		return nil, _fdeggb
	}
	return _dbeaa.Sum(nil), nil
}

// Inspect inspects the object types, subtypes and content in the PDF file returning a map of
// object type to number of instances of each.
func (_ddddf *PdfReader) Inspect() (map[string]int, error) { return _ddddf._baad.Inspect() }
func (_bca *PdfReader) newPdfAnnotationHighlightFromDict(_gcbea *_dg.PdfObjectDictionary) (*PdfAnnotationHighlight, error) {
	_cgd := PdfAnnotationHighlight{}
	_cdgd, _cac := _bca.newPdfAnnotationMarkupFromDict(_gcbea)
	if _cac != nil {
		return nil, _cac
	}
	_cgd.PdfAnnotationMarkup = _cdgd
	_cgd.QuadPoints = _gcbea.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_cgd, nil
}

// ToPdfObject implements interface PdfModel.
func (_dfge *PdfAnnotation3D) ToPdfObject() _dg.PdfObject {
	_dfge.PdfAnnotation.ToPdfObject()
	_ebaed := _dfge._cdf
	_aagd := _ebaed.PdfObject.(*_dg.PdfObjectDictionary)
	_aagd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0033\u0044"))
	_aagd.SetIfNotNil("\u0033\u0044\u0044", _dfge.T3DD)
	_aagd.SetIfNotNil("\u0033\u0044\u0056", _dfge.T3DV)
	_aagd.SetIfNotNil("\u0033\u0044\u0041", _dfge.T3DA)
	_aagd.SetIfNotNil("\u0033\u0044\u0049", _dfge.T3DI)
	_aagd.SetIfNotNil("\u0033\u0044\u0042", _dfge.T3DB)
	return _ebaed
}

// NewPdfActionJavaScript returns a new "javaScript" action.
func NewPdfActionJavaScript() *PdfActionJavaScript {
	_dfa := NewPdfAction()
	_fdc := &PdfActionJavaScript{}
	_fdc.PdfAction = _dfa
	_dfa.SetContext(_fdc)
	return _fdc
}

var _cfeab = false

func (_cegadb *PdfSignature) extractChainFromCert() ([]*_bb.Certificate, error) {
	var _ffeb *_dg.PdfObjectArray
	switch _bcaefb := _cegadb.Cert.(type) {
	case *_dg.PdfObjectString:
		_ffeb = _dg.MakeArray(_bcaefb)
	case *_dg.PdfObjectArray:
		_ffeb = _bcaefb
	default:
		return nil, _b.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _bcaefb)
	}
	var _daafc _bc.Buffer
	for _, _dfeeb := range _ffeb.Elements() {
		_ffgae, _dcece := _dg.GetString(_dfeeb)
		if !_dcece {
			return nil, _b.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054", _dfeeb)
		}
		if _, _adbdae := _daafc.Write(_ffgae.Bytes()); _adbdae != nil {
			return nil, _adbdae
		}
	}
	return _bb.ParseCertificates(_daafc.Bytes())
}

// PdfActionNamed represents a named action.
type PdfActionNamed struct {
	*PdfAction
	N _dg.PdfObject
}

func (_ffg *PdfReader) newPdfAnnotationLineFromDict(_edd *_dg.PdfObjectDictionary) (*PdfAnnotationLine, error) {
	_abgf := PdfAnnotationLine{}
	_abd, _fba := _ffg.newPdfAnnotationMarkupFromDict(_edd)
	if _fba != nil {
		return nil, _fba
	}
	_abgf.PdfAnnotationMarkup = _abd
	_abgf.L = _edd.Get("\u004c")
	_abgf.BS = _edd.Get("\u0042\u0053")
	_abgf.LE = _edd.Get("\u004c\u0045")
	_abgf.IC = _edd.Get("\u0049\u0043")
	_abgf.LL = _edd.Get("\u004c\u004c")
	_abgf.LLE = _edd.Get("\u004c\u004c\u0045")
	_abgf.Cap = _edd.Get("\u0043\u0061\u0070")
	_abgf.IT = _edd.Get("\u0049\u0054")
	_abgf.LLO = _edd.Get("\u004c\u004c\u004f")
	_abgf.CP = _edd.Get("\u0043\u0050")
	_abgf.Measure = _edd.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	_abgf.CO = _edd.Get("\u0043\u004f")
	return &_abgf, nil
}

// GetAsTilingPattern returns a tiling pattern. Check with IsTiling() prior to using this.
func (_bdbe *PdfPattern) GetAsTilingPattern() *PdfTilingPattern {
	return _bdbe._cgdcc.(*PdfTilingPattern)
}

// ToPdfObject implements interface PdfModel.
func (_fd *PdfAction) ToPdfObject() _dg.PdfObject {
	_gcg := _fd._cbd
	_ef := _gcg.PdfObject.(*_dg.PdfObjectDictionary)
	_ef.Clear()
	_ef.Set("\u0054\u0079\u0070\u0065", _dg.MakeName("\u0041\u0063\u0074\u0069\u006f\u006e"))
	_ef.SetIfNotNil("\u0053", _fd.S)
	_ef.SetIfNotNil("\u004e\u0065\u0078\u0074", _fd.Next)
	return _gcg
}

// FieldImageProvider provides fields images for specified fields.
type FieldImageProvider interface {
	FieldImageValues() (map[string]*Image, error)
}

func (_fgfdb *PdfReader) newPdfOutlineItemFromIndirectObject(_fbecee *_dg.PdfIndirectObject) (*PdfOutlineItem, error) {
	_ggdfc, _ebdbc := _fbecee.PdfObject.(*_dg.PdfObjectDictionary)
	if !_ebdbc {
		return nil, _b.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_ecfbb := NewPdfOutlineItem()
	_fcagf := _ggdfc.Get("\u0054\u0069\u0074l\u0065")
	if _fcagf == nil {
		return nil, _b.Errorf("\u006d\u0069\u0073s\u0069\u006e\u0067\u0020\u0054\u0069\u0074\u006c\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0049\u0074\u0065\u006d\u0020\u0028r\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	_acgd, _cedc := _dg.GetString(_fcagf)
	if !_cedc {
		return nil, _b.Errorf("\u0074\u0069\u0074le\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0054\u0029", _fcagf)
	}
	_ecfbb.Title = _acgd
	if _edcfg := _ggdfc.Get("\u0043\u006f\u0075n\u0074"); _edcfg != nil {
		_feege, _feda := _edcfg.(*_dg.PdfObjectInteger)
		if !_feda {
			return nil, _b.Errorf("\u0063o\u0075\u006e\u0074\u0020n\u006f\u0074\u0020\u0061\u006e \u0069n\u0074e\u0067\u0065\u0072\u0020\u0028\u0025\u0054)", _edcfg)
		}
		_geafd := int64(*_feege)
		_ecfbb.Count = &_geafd
	}
	if _edfge := _ggdfc.Get("\u0044\u0065\u0073\u0074"); _edfge != nil {
		_ecfbb.Dest = _dg.ResolveReference(_edfge)
		if !_fgfdb._dadcef {
			_caggf := _fgfdb.traverseObjectData(_ecfbb.Dest)
			if _caggf != nil {
				return nil, _caggf
			}
		}
	}
	if _bcfcf := _ggdfc.Get("\u0041"); _bcfcf != nil {
		_ecfbb.A = _dg.ResolveReference(_bcfcf)
		if !_fgfdb._dadcef {
			_gfcge := _fgfdb.traverseObjectData(_ecfbb.A)
			if _gfcge != nil {
				return nil, _gfcge
			}
		}
	}
	if _acdd := _ggdfc.Get("\u0053\u0045"); _acdd != nil {
		_ecfbb.SE = nil
	}
	if _egcc := _ggdfc.Get("\u0043"); _egcc != nil {
		_ecfbb.C = _dg.ResolveReference(_egcc)
	}
	if _cfge := _ggdfc.Get("\u0046"); _cfge != nil {
		_ecfbb.F = _dg.ResolveReference(_cfge)
	}
	return _ecfbb, nil
}

// PdfAnnotationSquiggly represents Squiggly annotations.
// (Section 12.5.6.10).
type PdfAnnotationSquiggly struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _dg.PdfObject
}

// SetXObjectByName adds the XObject from the passed in stream to the page resources.
// The added XObject is identified by the specified name.
func (_ddadbf *PdfPageResources) SetXObjectByName(keyName _dg.PdfObjectName, stream *_dg.PdfObjectStream) error {
	if _ddadbf.XObject == nil {
		_ddadbf.XObject = _dg.MakeDict()
	}
	_acbf := _dg.TraceToDirectObject(_ddadbf.XObject)
	_dgddd, _feggd := _acbf.(*_dg.PdfObjectDictionary)
	if !_feggd {
		_ag.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0058\u004f\u0062j\u0065\u0063\u0074\u002c\u0020\u0067\u006f\u0074\u0020\u0025T\u002f\u0025\u0054", _ddadbf.XObject, _acbf)
		return _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_dgddd.Set(keyName, stream)
	return nil
}

// NewStandard14Font returns the standard 14 font named `basefont` as a *PdfFont, or an error if it
// `basefont` is not one of the standard 14 font names.
func NewStandard14Font(basefont StdFontName) (*PdfFont, error) {
	_caeb, _gbec := _eece(basefont)
	if _gbec != nil {
		return nil, _gbec
	}
	if basefont != SymbolName && basefont != ZapfDingbatsName {
		_caeb._bdeed = _bd.NewWinAnsiEncoder()
	}
	return &PdfFont{_cadf: &_caeb}, nil
}

// Val returns the value of the color.
func (_ggea *PdfColorCalGray) Val() float64 { return float64(*_ggea) }

// ColorFromFloats returns a new PdfColor based on input color components.
func (_bggf *PdfColorspaceDeviceN) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != _bggf.GetNumComponents() {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_acgb, _abdcg := _bggf.TintTransform.Evaluate(vals)
	if _abdcg != nil {
		return nil, _abdcg
	}
	_bfda, _abdcg := _bggf.AlternateSpace.ColorFromFloats(_acgb)
	if _abdcg != nil {
		return nil, _abdcg
	}
	return _bfda, nil
}

// GetCatalogMetadata gets the catalog defined XMP Metadata.
func (_dacf *PdfReader) GetCatalogMetadata() (_dg.PdfObject, bool) {
	if _dacf._gccfb == nil {
		return nil, false
	}
	_dbedf := _dacf._gccfb.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _dbedf, _dbedf != nil
}

// GetRevisionNumber returns the version of the current Pdf document
func (_gbde *PdfReader) GetRevisionNumber() int { return _gbde._baad.GetRevisionNumber() }

// PdfPage represents a page in a PDF document. (7.7.3.3 - Table 30).
type PdfPage struct {
	Parent               _dg.PdfObject
	LastModified         *PdfDate
	Resources            *PdfPageResources
	CropBox              *PdfRectangle
	MediaBox             *PdfRectangle
	BleedBox             *PdfRectangle
	TrimBox              *PdfRectangle
	ArtBox               *PdfRectangle
	BoxColorInfo         _dg.PdfObject
	Contents             _dg.PdfObject
	Rotate               *int64
	Group                _dg.PdfObject
	Thumb                _dg.PdfObject
	B                    _dg.PdfObject
	Dur                  _dg.PdfObject
	Trans                _dg.PdfObject
	AA                   _dg.PdfObject
	Metadata             _dg.PdfObject
	PieceInfo            _dg.PdfObject
	StructParents        _dg.PdfObject
	ID                   _dg.PdfObject
	PZ                   _dg.PdfObject
	SeparationInfo       _dg.PdfObject
	Tabs                 _dg.PdfObject
	TemplateInstantiated _dg.PdfObject
	PresSteps            _dg.PdfObject
	UserUnit             _dg.PdfObject
	VP                   _dg.PdfObject
	Annots               _dg.PdfObject
	_cadgg               []*PdfAnnotation
	_bfdge               *_dg.PdfObjectDictionary
	_cggbe               *_dg.PdfIndirectObject
	_gaed                _dg.PdfObjectDictionary
	_cbbcc               *PdfReader
}

// ToPdfObject implements interface PdfModel.
func (_gcgc *PdfAnnotationInk) ToPdfObject() _dg.PdfObject {
	_gcgc.PdfAnnotation.ToPdfObject()
	_gcgcf := _gcgc._cdf
	_abdd := _gcgcf.PdfObject.(*_dg.PdfObjectDictionary)
	_gcgc.PdfAnnotationMarkup.appendToPdfDictionary(_abdd)
	_abdd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0049\u006e\u006b"))
	_abdd.SetIfNotNil("\u0049n\u006b\u004c\u0069\u0073\u0074", _gcgc.InkList)
	_abdd.SetIfNotNil("\u0042\u0053", _gcgc.BS)
	return _gcgcf
}

// GetModelFromPrimitive returns the model corresponding to the `primitive` PdfObject.
func (_ffgd *modelManager) GetModelFromPrimitive(primitive _dg.PdfObject) PdfModel {
	model, _cccb := _ffgd._eccbc[primitive]
	if !_cccb {
		return nil
	}
	return model
}

// PageFromIndirectObject returns the PdfPage and page number for a given indirect object.
func (_cbcba *PdfReader) PageFromIndirectObject(ind *_dg.PdfIndirectObject) (*PdfPage, int, error) {
	if len(_cbcba.PageList) != len(_cbcba._daddd) {
		return nil, 0, _bf.New("\u0070\u0061\u0067\u0065\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	for _bdec, _bebcb := range _cbcba._daddd {
		if _bebcb == ind {
			return _cbcba.PageList[_bdec], _bdec + 1, nil
		}
	}
	return nil, 0, _bf.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// Set sets the colorspace corresponding to key. Add to Names if not set.
func (_dgdbd *PdfPageResourcesColorspaces) Set(key _dg.PdfObjectName, val PdfColorspace) {
	if _, _fbfbbg := _dgdbd.Colorspaces[string(key)]; !_fbfbbg {
		_dgdbd.Names = append(_dgdbd.Names, string(key))
	}
	_dgdbd.Colorspaces[string(key)] = val
}

// ColorToRGB converts an Indexed color to an RGB color.
func (_gbcde *PdfColorspaceSpecialIndexed) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _gbcde.Base == nil {
		return nil, _bf.New("\u0069\u006e\u0064\u0065\u0078\u0065d\u0020\u0062\u0061\u0073\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _gbcde.Base.ColorToRGB(color)
}

// NewCompositePdfFontFromTTFFile loads a composite font from a TTF font file. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTFFile.
func NewCompositePdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_ecgbf, _cegdd := _eb.Open(filePath)
	if _cegdd != nil {
		_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006f\u0070\u0065\u006e\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _cegdd)
		return nil, _cegdd
	}
	defer _ecgbf.Close()
	return NewCompositePdfFontFromTTF(_ecgbf)
}
func (_efdaa *PdfWriter) copyObjects() {
	_bgcae := make(map[_dg.PdfObject]_dg.PdfObject)
	_bdefe := make([]_dg.PdfObject, 0, len(_efdaa._agaba))
	_efcdb := make(map[_dg.PdfObject]struct{}, len(_efdaa._agaba))
	_ebedc := make(map[_dg.PdfObject]struct{})
	for _, _dbdc := range _efdaa._agaba {
		_febae := _efdaa.copyObject(_dbdc, _bgcae, _ebedc, false)
		if _, _gacad := _ebedc[_dbdc]; _gacad {
			continue
		}
		_bdefe = append(_bdefe, _febae)
		_efcdb[_febae] = struct{}{}
	}
	_efdaa._agaba = _bdefe
	_efdaa._fdbfa = _efcdb
	_efdaa._efbfa = _efdaa.copyObject(_efdaa._efbfa, _bgcae, nil, false).(*_dg.PdfIndirectObject)
	_efdaa._fadee = _efdaa.copyObject(_efdaa._fadee, _bgcae, nil, false).(*_dg.PdfIndirectObject)
	if _efdaa._acdag != nil {
		_efdaa._acdag = _efdaa.copyObject(_efdaa._acdag, _bgcae, nil, false).(*_dg.PdfIndirectObject)
	}
	if _efdaa._bbac {
		_abdca := make(map[_dg.PdfObject]int64)
		for _bded, _gfaac := range _efdaa._ecabb {
			if _ccbf, _ddgb := _bgcae[_bded]; _ddgb {
				_abdca[_ccbf] = _gfaac
			} else {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020a\u0070\u0070\u0065n\u0064\u0020\u006d\u006fd\u0065\u0020\u002d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u0070\u0079\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070")
			}
		}
		_efdaa._ecabb = _abdca
	}
}

const (
	ActionTypeGoTo        PdfActionType = "\u0047\u006f\u0054\u006f"
	ActionTypeGoTo3DView  PdfActionType = "\u0047\u006f\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077"
	ActionTypeGoToE       PdfActionType = "\u0047\u006f\u0054o\u0045"
	ActionTypeGoToR       PdfActionType = "\u0047\u006f\u0054o\u0052"
	ActionTypeHide        PdfActionType = "\u0048\u0069\u0064\u0065"
	ActionTypeImportData  PdfActionType = "\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061"
	ActionTypeJavaScript  PdfActionType = "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"
	ActionTypeLaunch      PdfActionType = "\u004c\u0061\u0075\u006e\u0063\u0068"
	ActionTypeMovie       PdfActionType = "\u004d\u006f\u0076i\u0065"
	ActionTypeNamed       PdfActionType = "\u004e\u0061\u006de\u0064"
	ActionTypeRendition   PdfActionType = "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn"
	ActionTypeResetForm   PdfActionType = "\u0052e\u0073\u0065\u0074\u0046\u006f\u0072m"
	ActionTypeSetOCGState PdfActionType = "S\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065"
	ActionTypeSound       PdfActionType = "\u0053\u006f\u0075n\u0064"
	ActionTypeSubmitForm  PdfActionType = "\u0053\u0075\u0062\u006d\u0069\u0074\u0046\u006f\u0072\u006d"
	ActionTypeThread      PdfActionType = "\u0054\u0068\u0072\u0065\u0061\u0064"
	ActionTypeTrans       PdfActionType = "\u0054\u0072\u0061n\u0073"
	ActionTypeURI         PdfActionType = "\u0055\u0052\u0049"
)

// NewPdfColorDeviceRGB returns a new PdfColorDeviceRGB based on the r,g,b component values.
func NewPdfColorDeviceRGB(r, g, b float64) *PdfColorDeviceRGB {
	_eabaf := PdfColorDeviceRGB{r, g, b}
	return &_eabaf
}
func _ggcb(_eaefc map[_bbg.GID]int, _cdegd uint16) *_dg.PdfObjectArray {
	_bbeac := &_dg.PdfObjectArray{}
	_egbfc := _bbg.GID(_cdegd)
	for _cgbfe := _bbg.GID(0); _cgbfe < _egbfc; {
		_fgagf, _edabf := _eaefc[_cgbfe]
		if !_edabf {
			_cgbfe++
			continue
		}
		_dfgcf := _cgbfe
		for _dfeb := _dfgcf + 1; _dfeb < _egbfc; _dfeb++ {
			if _cfdfd, _ggag := _eaefc[_dfeb]; !_ggag || _fgagf != _cfdfd {
				break
			}
			_dfgcf = _dfeb
		}
		_bbeac.Append(_dg.MakeInteger(int64(_cgbfe)))
		_bbeac.Append(_dg.MakeInteger(int64(_dfgcf)))
		_bbeac.Append(_dg.MakeInteger(int64(_fgagf)))
		_cgbfe = _dfgcf + 1
	}
	return _bbeac
}

// NewPdfActionImportData returns a new "import data" action.
func NewPdfActionImportData() *PdfActionImportData {
	_ca := NewPdfAction()
	_cbb := &PdfActionImportData{}
	_cbb.PdfAction = _ca
	_ca.SetContext(_cbb)
	return _cbb
}
func (_bbdd *PdfColorspaceDeviceCMYK) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
}

// SetReason sets the `Reason` field of the signature.
func (_dgbeg *PdfSignature) SetReason(reason string) {
	_dgbeg.Reason = _dg.MakeEncodedString(reason, true)
}

// ToPdfObject implements interface PdfModel.
func (_fcgg *PdfSignature) ToPdfObject() _dg.PdfObject {
	_eggbf := _fcgg._bbda
	var _bcefg *_dg.PdfObjectDictionary
	if _beadde, _gaaca := _eggbf.PdfObject.(*pdfSignDictionary); _gaaca {
		_bcefg = _beadde.PdfObjectDictionary
	} else {
		_bcefg = _eggbf.PdfObject.(*_dg.PdfObjectDictionary)
	}
	_bcefg.SetIfNotNil("\u0054\u0079\u0070\u0065", _fcgg.Type)
	_bcefg.SetIfNotNil("\u0046\u0069\u006c\u0074\u0065\u0072", _fcgg.Filter)
	_bcefg.SetIfNotNil("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r", _fcgg.SubFilter)
	_bcefg.SetIfNotNil("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _fcgg.ByteRange)
	_bcefg.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _fcgg.Contents)
	_bcefg.SetIfNotNil("\u0043\u0065\u0072\u0074", _fcgg.Cert)
	_bcefg.SetIfNotNil("\u004e\u0061\u006d\u0065", _fcgg.Name)
	_bcefg.SetIfNotNil("\u0052\u0065\u0061\u0073\u006f\u006e", _fcgg.Reason)
	_bcefg.SetIfNotNil("\u004d", _fcgg.M)
	_bcefg.SetIfNotNil("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e", _fcgg.Reference)
	_bcefg.SetIfNotNil("\u0043h\u0061\u006e\u0067\u0065\u0073", _fcgg.Changes)
	_bcefg.SetIfNotNil("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f", _fcgg.ContactInfo)
	return _eggbf
}

// IsTiling specifies if the pattern is a tiling pattern.
func (_ebce *PdfPattern) IsTiling() bool { return _ebce.PatternType == 1 }

// ToPdfObject returns a PDF object representation of the outline destination.
func (_fcacd OutlineDest) ToPdfObject() _dg.PdfObject {
	if (_fcacd.PageObj == nil && _fcacd.Page < 0) || _fcacd.Mode == "" {
		return _dg.MakeNull()
	}
	_cdbgg := _dg.MakeArray()
	if _fcacd.PageObj != nil {
		_cdbgg.Append(_fcacd.PageObj)
	} else {
		_cdbgg.Append(_dg.MakeInteger(_fcacd.Page))
	}
	_cdbgg.Append(_dg.MakeName(_fcacd.Mode))
	switch _fcacd.Mode {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		_cdbgg.Append(_dg.MakeFloat(_fcacd.Y))
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		_cdbgg.Append(_dg.MakeFloat(_fcacd.X))
	case "\u0058\u0059\u005a":
		_cdbgg.Append(_dg.MakeFloat(_fcacd.X))
		_cdbgg.Append(_dg.MakeFloat(_fcacd.Y))
		_cdbgg.Append(_dg.MakeFloat(_fcacd.Zoom))
	default:
		_cdbgg.Set(1, _dg.MakeName("\u0046\u0069\u0074"))
	}
	return _cdbgg
}
func (_fbfecf *PdfWriter) writeObject(_dffg int, _fgefd _dg.PdfObject) {
	_ag.Log.Trace("\u0057\u0072\u0069\u0074\u0065\u0020\u006f\u0062\u006a \u0023\u0025\u0064\u000a", _dffg)
	if _addab, _ebbdc := _fgefd.(*_dg.PdfIndirectObject); _ebbdc {
		_fbfecf._fffge[_dffg] = crossReference{Type: 1, Offset: _fbfecf._fbbfc, Generation: _addab.GenerationNumber}
		_adfcd := _b.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _dffg)
		if _deaea, _bddfc := _addab.PdfObject.(*pdfSignDictionary); _bddfc {
			_deaea._cagb = _fbfecf._fbbfc + int64(len(_adfcd))
		}
		if _addab.PdfObject == nil {
			_ag.Log.Debug("E\u0072\u0072\u006fr\u003a\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0027\u0073\u0020\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020b\u0065\u0020\u006e\u0069l\u0020\u002d\u0020\u0073e\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063t\u004e\u0075\u006c\u006c")
			_addab.PdfObject = _dg.MakeNull()
		}
		_adfcd += _addab.PdfObject.WriteString()
		_adfcd += "\u000a\u0065\u006e\u0064\u006f\u0062\u006a\u000a"
		_fbfecf.writeString(_adfcd)
		return
	}
	if _caccb, _fdfbec := _fgefd.(*_dg.PdfObjectStream); _fdfbec {
		_fbfecf._fffge[_dffg] = crossReference{Type: 1, Offset: _fbfecf._fbbfc, Generation: _caccb.GenerationNumber}
		_afcac := _b.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _dffg)
		_afcac += _caccb.PdfObjectDictionary.WriteString()
		_afcac += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_fbfecf.writeString(_afcac)
		_fbfecf.writeBytes(_caccb.Stream)
		_fbfecf.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	if _dfec, _accga := _fgefd.(*_dg.PdfObjectStreams); _accga {
		_fbfecf._fffge[_dffg] = crossReference{Type: 1, Offset: _fbfecf._fbbfc, Generation: _dfec.GenerationNumber}
		_eebf := _b.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _dffg)
		var _bdage []string
		var _aadbf string
		var _aadcd int64
		for _defbg, _ebfgg := range _dfec.Elements() {
			_caddb, _geedg := _ebfgg.(*_dg.PdfIndirectObject)
			if !_geedg {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065am\u0073 \u004e\u0020\u0025\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006es\u0020\u006e\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u0070\u0064\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0076", _dffg, _ebfgg)
				continue
			}
			_fgecd := _caddb.PdfObject.WriteString() + "\u0020"
			_aadbf = _aadbf + _fgecd
			_bdage = append(_bdage, _b.Sprintf("\u0025\u0064\u0020%\u0064", _caddb.ObjectNumber, _aadcd))
			_fbfecf._fffge[int(_caddb.ObjectNumber)] = crossReference{Type: 2, ObjectNumber: _dffg, Index: _defbg}
			_aadcd = _aadcd + int64(len([]byte(_fgecd)))
		}
		_cdcbf := _ga.Join(_bdage, "\u0020") + "\u0020"
		_gdgcf := _dg.NewFlateEncoder()
		_cfbbe := _gdgcf.MakeStreamDict()
		_cfbbe.Set(_dg.PdfObjectName("\u0054\u0079\u0070\u0065"), _dg.MakeName("\u004f\u0062\u006a\u0053\u0074\u006d"))
		_cfadc := int64(_dfec.Len())
		_cfbbe.Set(_dg.PdfObjectName("\u004e"), _dg.MakeInteger(_cfadc))
		_acgegd := int64(len(_cdcbf))
		_cfbbe.Set(_dg.PdfObjectName("\u0046\u0069\u0072s\u0074"), _dg.MakeInteger(_acgegd))
		_aaec, _ := _gdgcf.EncodeBytes([]byte(_cdcbf + _aadbf))
		_cbgcg := int64(len(_aaec))
		_cfbbe.Set(_dg.PdfObjectName("\u004c\u0065\u006e\u0067\u0074\u0068"), _dg.MakeInteger(_cbgcg))
		_eebf += _cfbbe.WriteString()
		_eebf += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_fbfecf.writeString(_eebf)
		_fbfecf.writeBytes(_aaec)
		_fbfecf.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	_fbfecf.writeString(_fgefd.WriteString())
}

// SetXObjectFormByName adds the provided XObjectForm to the page resources.
// The added XObjectForm is identified by the specified name.
func (_afbec *PdfPageResources) SetXObjectFormByName(keyName _dg.PdfObjectName, xform *XObjectForm) error {
	_fbcgg := xform.ToPdfObject().(*_dg.PdfObjectStream)
	_ffdac := _afbec.SetXObjectByName(keyName, _fbcgg)
	return _ffdac
}

// NewPdfRectangle creates a PDF rectangle object based on an input array of 4 integers.
// Defining the lower left (LL) and upper right (UR) corners with
// floating point numbers.
func NewPdfRectangle(arr _dg.PdfObjectArray) (*PdfRectangle, error) {
	_aedfc := PdfRectangle{}
	if arr.Len() != 4 {
		return nil, _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065\u0020\u0061\u0072r\u0061\u0079\u002c\u0020\u006c\u0065\u006e \u0021\u003d\u0020\u0034")
	}
	var _beege error
	_aedfc.Llx, _beege = _dg.GetNumberAsFloat(arr.Get(0))
	if _beege != nil {
		return nil, _beege
	}
	_aedfc.Lly, _beege = _dg.GetNumberAsFloat(arr.Get(1))
	if _beege != nil {
		return nil, _beege
	}
	_aedfc.Urx, _beege = _dg.GetNumberAsFloat(arr.Get(2))
	if _beege != nil {
		return nil, _beege
	}
	_aedfc.Ury, _beege = _dg.GetNumberAsFloat(arr.Get(3))
	if _beege != nil {
		return nil, _beege
	}
	return &_aedfc, nil
}

// NewPdfAnnotationSound returns a new sound annotation.
func NewPdfAnnotationSound() *PdfAnnotationSound {
	_ffe := NewPdfAnnotation()
	_afe := &PdfAnnotationSound{}
	_afe.PdfAnnotation = _ffe
	_afe.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ffe.SetContext(_afe)
	return _afe
}

// WriteString outputs the object as it is to be written to file.
func (_agca *PdfTransformParamsDocMDP) WriteString() string { return _agca.ToPdfObject().WriteString() }

// PdfSignature represents a PDF signature dictionary and is used for signing via form signature fields.
// (Section 12.8, Table 252 - Entries in a signature dictionary p. 475 in PDF32000_2008).
type PdfSignature struct {
	Handler SignatureHandler
	_bbda   *_dg.PdfIndirectObject

	// Type: Sig/DocTimeStamp
	Type         *_dg.PdfObjectName
	Filter       *_dg.PdfObjectName
	SubFilter    *_dg.PdfObjectName
	Contents     *_dg.PdfObjectString
	Cert         _dg.PdfObject
	ByteRange    *_dg.PdfObjectArray
	Reference    *_dg.PdfObjectArray
	Changes      *_dg.PdfObjectArray
	Name         *_dg.PdfObjectString
	M            *_dg.PdfObjectString
	Location     *_dg.PdfObjectString
	Reason       *_dg.PdfObjectString
	ContactInfo  *_dg.PdfObjectString
	R            *_dg.PdfObjectInteger
	V            *_dg.PdfObjectInteger
	PropBuild    *_dg.PdfObjectDictionary
	PropAuthTime *_dg.PdfObjectInteger
	PropAuthType *_dg.PdfObjectName
}

func _dadaa(_aefg *_dg.PdfObjectDictionary, _acfd *fontCommon) (*pdfCIDFontType2, error) {
	if _acfd._bcga != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_ag.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0032\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _acfd)
		return nil, _dg.ErrRangeError
	}
	_ecefd := _ebccc(_acfd)
	_gfgec, _fgcb := _dg.GetDict(_aefg.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_fgcb {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _acfd)
		return nil, ErrRequiredAttributeMissing
	}
	_ecefd.CIDSystemInfo = _gfgec
	_ecefd.DW = _aefg.Get("\u0044\u0057")
	_ecefd.W = _aefg.Get("\u0057")
	_ecefd.DW2 = _aefg.Get("\u0044\u0057\u0032")
	_ecefd.W2 = _aefg.Get("\u0057\u0032")
	_ecefd.CIDToGIDMap = _aefg.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070")
	_ecefd._bfbc = 1000.0
	if _edcb, _fgddcd := _dg.GetNumberAsFloat(_ecefd.DW); _fgddcd == nil {
		_ecefd._bfbc = _edcb
	}
	_acfee, _aegg := _gacbe(_ecefd.W)
	if _aegg != nil {
		return nil, _aegg
	}
	if _acfee == nil {
		_acfee = map[_bd.CharCode]float64{}
	}
	_ecefd._ddeb = _acfee
	return _ecefd, nil
}

// NewPdfAnnotationCaret returns a new caret annotation.
func NewPdfAnnotationCaret() *PdfAnnotationCaret {
	_gdfa := NewPdfAnnotation()
	_cedb := &PdfAnnotationCaret{}
	_cedb.PdfAnnotation = _gdfa
	_cedb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gdfa.SetContext(_cedb)
	return _cedb
}

var ImageHandling ImageHandler = DefaultImageHandler{}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a grayscale device.
func (_fadc *PdfColorspaceDeviceGray) GetNumComponents() int { return 1 }

// Evaluate runs the function on the passed in slice and returns the results.
func (_dbcg *PdfFunctionType0) Evaluate(x []float64) ([]float64, error) {
	if len(x) != _dbcg.NumInputs {
		_ag.Log.Error("\u004eu\u006d\u0062e\u0072\u0020\u006f\u0066 \u0069\u006e\u0070u\u0074\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061tc\u0068\u0069\u006eg\u0020\u0077h\u0061\u0074\u0020\u0069\u0073\u0020n\u0065\u0065d\u0065\u0064")
		return nil, _bf.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _dbcg._fecfc == nil {
		_ecbba := _dbcg.processSamples()
		if _ecbba != nil {
			return nil, _ecbba
		}
	}
	_dabag := _dbcg.Encode
	if _dabag == nil {
		_dabag = []float64{}
		for _eacbc := 0; _eacbc < len(_dbcg.Size); _eacbc++ {
			_dabag = append(_dabag, 0)
			_dabag = append(_dabag, float64(_dbcg.Size[_eacbc]-1))
		}
	}
	_acaf := _dbcg.Decode
	if _acaf == nil {
		_acaf = _dbcg.Range
	}
	_adfcg := make([]int, len(x))
	for _abgfe := 0; _abgfe < len(x); _abgfe++ {
		_adgca := x[_abgfe]
		_bbcc := _cg.Min(_cg.Max(_adgca, _dbcg.Domain[2*_abgfe]), _dbcg.Domain[2*_abgfe+1])
		_cgdcg := _fc.LinearInterpolate(_bbcc, _dbcg.Domain[2*_abgfe], _dbcg.Domain[2*_abgfe+1], _dabag[2*_abgfe], _dabag[2*_abgfe+1])
		_babe := _cg.Min(_cg.Max(_cgdcg, 0), float64(_dbcg.Size[_abgfe]-1))
		_befdg := int(_cg.Floor(_babe + 0.5))
		if _befdg < 0 {
			_befdg = 0
		} else if _befdg > _dbcg.Size[_abgfe] {
			_befdg = _dbcg.Size[_abgfe] - 1
		}
		_adfcg[_abgfe] = _befdg
	}
	_acgfa := _adfcg[0]
	for _bdgaa := 1; _bdgaa < _dbcg.NumInputs; _bdgaa++ {
		_fbabc := _adfcg[_bdgaa]
		for _agggc := 0; _agggc < _bdgaa; _agggc++ {
			_fbabc *= _dbcg.Size[_agggc]
		}
		_acgfa += _fbabc
	}
	_acgfa *= _dbcg.NumOutputs
	var _beef []float64
	for _ggeg := 0; _ggeg < _dbcg.NumOutputs; _ggeg++ {
		_ddgaf := _acgfa + _ggeg
		if _ddgaf >= len(_dbcg._fecfc) {
			_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u006e\u006ft\u0020\u0065\u006eo\u0075\u0067\u0068\u0020\u0069\u006ep\u0075\u0074\u0020sa\u006dp\u006c\u0065\u0073\u0020\u0074\u006f\u0020d\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0076\u0061lu\u0065\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			continue
		}
		_fecc := _dbcg._fecfc[_ddgaf]
		_gecce := _fc.LinearInterpolate(float64(_fecc), 0, _cg.Pow(2, float64(_dbcg.BitsPerSample)), _acaf[2*_ggeg], _acaf[2*_ggeg+1])
		_eacf := _cg.Min(_cg.Max(_gecce, _dbcg.Range[2*_ggeg]), _dbcg.Range[2*_ggeg+1])
		_beef = append(_beef, _eacf)
	}
	return _beef, nil
}

// OutlineItem represents a PDF outline item dictionary (Table 153 - pp. 376 - 377).
type OutlineItem struct {
	Title   string         `json:"title"`
	Dest    OutlineDest    `json:"dest"`
	Entries []*OutlineItem `json:"entries,omitempty"`
}

// HasShadingByName checks whether a shading is defined by the specified keyName.
func (_cadbd *PdfPageResources) HasShadingByName(keyName _dg.PdfObjectName) bool {
	_, _bbffc := _cadbd.GetShadingByName(keyName)
	return _bbffc
}

// GetContainingPdfObject gets the primitive used to parse the color space.
func (_aaeeb *PdfColorspaceICCBased) GetContainingPdfObject() _dg.PdfObject { return _aaeeb._dafg }

const (
	RC4_128bit = EncryptionAlgorithm(iota)
	AES_128bit
	AES_256bit
)

// NewPdfTransformParamsDocMDP create a PdfTransformParamsDocMDP with the specific permissions.
func NewPdfTransformParamsDocMDP(permission _ecb.DocMDPPermission) *PdfTransformParamsDocMDP {
	return &PdfTransformParamsDocMDP{Type: _dg.MakeName("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"), P: _dg.MakeInteger(int64(permission)), V: _dg.MakeName("\u0031\u002e\u0032")}
}

// ToPdfObject implements interface PdfModel.
func (_bdcff *PdfAnnotationCircle) ToPdfObject() _dg.PdfObject {
	_bdcff.PdfAnnotation.ToPdfObject()
	_eaae := _bdcff._cdf
	_gbg := _eaae.PdfObject.(*_dg.PdfObjectDictionary)
	_bdcff.PdfAnnotationMarkup.appendToPdfDictionary(_gbg)
	_gbg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0043\u0069\u0072\u0063\u006c\u0065"))
	_gbg.SetIfNotNil("\u0042\u0053", _bdcff.BS)
	_gbg.SetIfNotNil("\u0049\u0043", _bdcff.IC)
	_gbg.SetIfNotNil("\u0042\u0045", _bdcff.BE)
	_gbg.SetIfNotNil("\u0052\u0044", _bdcff.RD)
	return _eaae
}

// ToPdfObject converts the pdfFontSimple to its PDF representation for outputting.
func (_gfcg *pdfFontSimple) ToPdfObject() _dg.PdfObject {
	if _gfcg._gccg == nil {
		_gfcg._gccg = &_dg.PdfIndirectObject{}
	}
	_gddf := _gfcg.baseFields().asPdfObjectDictionary("")
	_gfcg._gccg.PdfObject = _gddf
	if _gfcg.FirstChar != nil {
		_gddf.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _gfcg.FirstChar)
	}
	if _gfcg.LastChar != nil {
		_gddf.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _gfcg.LastChar)
	}
	if _gfcg.Widths != nil {
		_gddf.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _gfcg.Widths)
	}
	if _gfcg.Encoding != nil {
		_gddf.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _gfcg.Encoding)
	} else if _gfcg._bdeed != nil {
		_ddcfg := _gfcg._bdeed.ToPdfObject()
		if _ddcfg != nil {
			_gddf.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _ddcfg)
		}
	}
	return _gfcg._gccg
}

// ColorToRGB converts a CalRGB color to an RGB color.
func (_eceb *PdfColorspaceCalRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_edac, _ffda := color.(*PdfColorCalRGB)
	if !_ffda {
		_ag.Log.Debug("\u0049\u006e\u0070ut\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0063\u0061\u006c\u0020\u0072\u0067\u0062")
		return nil, _bf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_fbcc := _edac.A()
	_fggb := _edac.B()
	_efba := _edac.C()
	X := _eceb.Matrix[0]*_cg.Pow(_fbcc, _eceb.Gamma[0]) + _eceb.Matrix[3]*_cg.Pow(_fggb, _eceb.Gamma[1]) + _eceb.Matrix[6]*_cg.Pow(_efba, _eceb.Gamma[2])
	Y := _eceb.Matrix[1]*_cg.Pow(_fbcc, _eceb.Gamma[0]) + _eceb.Matrix[4]*_cg.Pow(_fggb, _eceb.Gamma[1]) + _eceb.Matrix[7]*_cg.Pow(_efba, _eceb.Gamma[2])
	Z := _eceb.Matrix[2]*_cg.Pow(_fbcc, _eceb.Gamma[0]) + _eceb.Matrix[5]*_cg.Pow(_fggb, _eceb.Gamma[1]) + _eceb.Matrix[8]*_cg.Pow(_efba, _eceb.Gamma[2])
	_beecbd := 3.240479*X + -1.537150*Y + -0.498535*Z
	_bbgc := -0.969256*X + 1.875992*Y + 0.041556*Z
	_fdab := 0.055648*X + -0.204043*Y + 1.057311*Z
	_beecbd = _cg.Min(_cg.Max(_beecbd, 0), 1.0)
	_bbgc = _cg.Min(_cg.Max(_bbgc, 0), 1.0)
	_fdab = _cg.Min(_cg.Max(_fdab, 0), 1.0)
	return NewPdfColorDeviceRGB(_beecbd, _bbgc, _fdab), nil
}

// ImageToRGB converts image in CalGray color space to RGB (A, B, C -> X, Y, Z).
func (_bbc *PdfColorspaceCalGray) ImageToRGB(img Image) (Image, error) {
	_edcdd := _fcd.NewReader(img.getBase())
	_egee := _fc.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_bdgbg := _fcd.NewWriter(_egee)
	_cgcde := _cg.Pow(2, float64(img.BitsPerComponent)) - 1
	_dccca := make([]uint32, 3)
	var (
		_adfd                               uint32
		ANorm, X, Y, Z, _febb, _beeg, _addb float64
		_cabgd                              error
	)
	for {
		_adfd, _cabgd = _edcdd.ReadSample()
		if _cabgd == _cf.EOF {
			break
		} else if _cabgd != nil {
			return img, _cabgd
		}
		ANorm = float64(_adfd) / _cgcde
		X = _bbc.WhitePoint[0] * _cg.Pow(ANorm, _bbc.Gamma)
		Y = _bbc.WhitePoint[1] * _cg.Pow(ANorm, _bbc.Gamma)
		Z = _bbc.WhitePoint[2] * _cg.Pow(ANorm, _bbc.Gamma)
		_febb = 3.240479*X + -1.537150*Y + -0.498535*Z
		_beeg = -0.969256*X + 1.875992*Y + 0.041556*Z
		_addb = 0.055648*X + -0.204043*Y + 1.057311*Z
		_febb = _cg.Min(_cg.Max(_febb, 0), 1.0)
		_beeg = _cg.Min(_cg.Max(_beeg, 0), 1.0)
		_addb = _cg.Min(_cg.Max(_addb, 0), 1.0)
		_dccca[0] = uint32(_febb * _cgcde)
		_dccca[1] = uint32(_beeg * _cgcde)
		_dccca[2] = uint32(_addb * _cgcde)
		if _cabgd = _bdgbg.WriteSamples(_dccca); _cabgd != nil {
			return img, _cabgd
		}
	}
	return _edcf(&_egee), nil
}
func (_dcbef *PdfReader) buildOutlineTree(_eeceb _dg.PdfObject, _ccebd *PdfOutlineTreeNode, _gacae *PdfOutlineTreeNode, _gebbf map[_dg.PdfObject]struct{}) (*PdfOutlineTreeNode, *PdfOutlineTreeNode, error) {
	if _gebbf == nil {
		_gebbf = map[_dg.PdfObject]struct{}{}
	}
	_gebbf[_eeceb] = struct{}{}
	_dcba, _cgbgc := _eeceb.(*_dg.PdfIndirectObject)
	if !_cgbgc {
		return nil, nil, _b.Errorf("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0063\u006f\u006et\u0061\u0069\u006e\u0065r\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _eeceb)
	}
	_gcdbf, _dagdc := _dcba.PdfObject.(*_dg.PdfObjectDictionary)
	if !_dagdc {
		return nil, nil, _bf.New("\u006e\u006f\u0074 a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_ag.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0074\u0072\u0065\u0065\u003a\u0020d\u0069\u0063\u0074\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0076\u0029\u0020p\u003a\u0020\u0025\u0070", _gcdbf, _dcba, _dcba)
	if _dgefd := _gcdbf.Get("\u0054\u0069\u0074l\u0065"); _dgefd != nil {
		_fbddc, _eadcc := _dcbef.newPdfOutlineItemFromIndirectObject(_dcba)
		if _eadcc != nil {
			return nil, nil, _eadcc
		}
		_fbddc.Parent = _ccebd
		_fbddc.Prev = _gacae
		_bfbf := _dg.ResolveReference(_gcdbf.Get("\u0046\u0069\u0072s\u0074"))
		if _, _cbef := _gebbf[_bfbf]; _bfbf != nil && _bfbf != _dcba && !_cbef {
			if !_dg.IsNullObject(_bfbf) {
				_edgcb, _cdccc, _fgedb := _dcbef.buildOutlineTree(_bfbf, &_fbddc.PdfOutlineTreeNode, nil, _gebbf)
				if _fgedb != nil {
					_ag.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0069\u0074\u0065\u006d\u0020\u0074\u0072\u0065\u0065\u003a \u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020n\u006f\u0064\u0065\u0020\u0063\u0068\u0069\u006c\u0064\u0072\u0065n\u002e", _fgedb)
				} else {
					_fbddc.First = _edgcb
					_fbddc.Last = _cdccc
				}
			}
		}
		_afdba := _dg.ResolveReference(_gcdbf.Get("\u004e\u0065\u0078\u0074"))
		if _, _dbded := _gebbf[_afdba]; _afdba != nil && _afdba != _dcba && !_dbded {
			if !_dg.IsNullObject(_afdba) {
				_gfbc, _bgfg, _feaggf := _dcbef.buildOutlineTree(_afdba, _ccebd, &_fbddc.PdfOutlineTreeNode, _gebbf)
				if _feaggf != nil {
					_ag.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065\u0020\u0066\u006f\u0072\u0020\u004ee\u0078\u0074\u0020\u006e\u006f\u0064\u0065\u003a\u0020\u0025\u0076\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006e\u006f\u0064e\u002e", _feaggf)
				} else {
					_fbddc.Next = _gfbc
					return &_fbddc.PdfOutlineTreeNode, _bgfg, nil
				}
			}
		}
		return &_fbddc.PdfOutlineTreeNode, &_fbddc.PdfOutlineTreeNode, nil
	}
	_egfdgd, _gfgae := _abeda(_dcba)
	if _gfgae != nil {
		return nil, nil, _gfgae
	}
	_egfdgd.Parent = _ccebd
	if _badddf := _gcdbf.Get("\u0046\u0069\u0072s\u0074"); _badddf != nil {
		_badddf = _dg.ResolveReference(_badddf)
		if _, _aggcb := _gebbf[_badddf]; _badddf != nil && _badddf != _dcba && !_aggcb {
			_fegda := _dg.TraceToDirectObject(_badddf)
			if _, _eddgd := _fegda.(*_dg.PdfObjectNull); !_eddgd && _fegda != nil {
				_cbbae, _egffa, _bcbgf := _dcbef.buildOutlineTree(_badddf, &_egfdgd.PdfOutlineTreeNode, nil, _gebbf)
				if _bcbgf != nil {
					_ag.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020b\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0074\u0072\u0065\u0065\u003a\u0020\u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006e\u006f\u0064\u0065 \u0063\u0068i\u006c\u0064r\u0065n\u002e", _bcbgf)
				} else {
					_egfdgd.First = _cbbae
					_egfdgd.Last = _egffa
				}
			}
		}
	}
	return &_egfdgd.PdfOutlineTreeNode, &_egfdgd.PdfOutlineTreeNode, nil
}

// Subtype returns the font's "Subtype" field.
func (_afgfd *PdfFont) Subtype() string {
	_cgbf := _afgfd.baseFields()._bcga
	if _ebcad, _daab := _afgfd._cadf.(*pdfFontType0); _daab {
		_cgbf = _cgbf + "\u003a" + _ebcad.DescendantFont.Subtype()
	}
	return _cgbf
}

// ToPdfObject implements interface PdfModel.
func (_ddac *PdfAnnotationPolygon) ToPdfObject() _dg.PdfObject {
	_ddac.PdfAnnotation.ToPdfObject()
	_edf := _ddac._cdf
	_egff := _edf.PdfObject.(*_dg.PdfObjectDictionary)
	_ddac.PdfAnnotationMarkup.appendToPdfDictionary(_egff)
	_egff.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0050o\u006c\u0079\u0067\u006f\u006e"))
	_egff.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _ddac.Vertices)
	_egff.SetIfNotNil("\u004c\u0045", _ddac.LE)
	_egff.SetIfNotNil("\u0042\u0053", _ddac.BS)
	_egff.SetIfNotNil("\u0049\u0043", _ddac.IC)
	_egff.SetIfNotNil("\u0042\u0045", _ddac.BE)
	_egff.SetIfNotNil("\u0049\u0054", _ddac.IT)
	_egff.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _ddac.Measure)
	return _edf
}

// ToPdfObject implements interface PdfModel.
func (_eea *PdfAnnotationStrikeOut) ToPdfObject() _dg.PdfObject {
	_eea.PdfAnnotation.ToPdfObject()
	_bddf := _eea._cdf
	_bdga := _bddf.PdfObject.(*_dg.PdfObjectDictionary)
	_eea.PdfAnnotationMarkup.appendToPdfDictionary(_bdga)
	_bdga.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t"))
	_bdga.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _eea.QuadPoints)
	return _bddf
}

var (
	_fgefgf _e.Mutex
	_bdeaf  = ""
	_ccfea  _a.Time
	_cebb   = ""
	_abbgf  = ""
	_efafac _a.Time
	_faff   = ""
	_adgf   = ""
	_fgfda  = ""
)

// Height returns the height of `rect`.
func (_fdgge *PdfRectangle) Height() float64 { return _cg.Abs(_fdgge.Ury - _fdgge.Lly) }

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_dggdbg *PdfShading) ToPdfObject() _dg.PdfObject {
	_bfcaf := _dggdbg._bcfbg
	_gdfbe, _effba := _dggdbg.getShadingDict()
	if _effba != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _dggdbg.ShadingType != nil {
		_gdfbe.Set("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _dggdbg.ShadingType)
	}
	if _dggdbg.ColorSpace != nil {
		_gdfbe.Set("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _dggdbg.ColorSpace.ToPdfObject())
	}
	if _dggdbg.Background != nil {
		_gdfbe.Set("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064", _dggdbg.Background)
	}
	if _dggdbg.BBox != nil {
		_gdfbe.Set("\u0042\u0042\u006f\u0078", _dggdbg.BBox.ToPdfObject())
	}
	if _dggdbg.AntiAlias != nil {
		_gdfbe.Set("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s", _dggdbg.AntiAlias)
	}
	return _bfcaf
}

// GetXObjectImageByName returns the XObjectImage with the specified name from the
// page resources, if it exists.
func (_acbc *PdfPageResources) GetXObjectImageByName(keyName _dg.PdfObjectName) (*XObjectImage, error) {
	_gfcgc, _fgaga := _acbc.GetXObjectByName(keyName)
	if _gfcgc == nil {
		return nil, nil
	}
	if _fgaga != XObjectTypeImage {
		return nil, _bf.New("\u006e\u006f\u0074 \u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065")
	}
	_fddcb, _adgdc := NewXObjectImageFromStream(_gfcgc)
	if _adgdc != nil {
		return nil, _adgdc
	}
	return _fddcb, nil
}

// PdfRectangle is a definition of a rectangle.
type PdfRectangle struct {
	Llx float64
	Lly float64
	Urx float64
	Ury float64
}

func (_daac *PdfWriter) setDocInfo(_aaecd _dg.PdfObject) {
	if _daac.hasObject(_daac._efbfa) {
		delete(_daac._fdbfa, _daac._efbfa)
		delete(_daac._cffaa, _daac._efbfa)
		for _bafgc, _ddeebg := range _daac._agaba {
			if _ddeebg == _daac._efbfa {
				copy(_daac._agaba[_bafgc:], _daac._agaba[_bafgc+1:])
				_daac._agaba[len(_daac._agaba)-1] = nil
				_daac._agaba = _daac._agaba[:len(_daac._agaba)-1]
				break
			}
		}
	}
	_fgggf := _dg.PdfIndirectObject{}
	_fgggf.PdfObject = _aaecd
	_daac._efbfa = &_fgggf
	_daac.addObject(&_fgggf)
}

// NewPdfReaderWithOpts creates a new PdfReader for an input io.ReadSeeker interface
// with a ReaderOpts.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderWithOpts(rs _cf.ReadSeeker, opts *ReaderOpts) (*PdfReader, error) {
	const _dfdcb = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0057\u0069\u0074\u0068O\u0070\u0074\u0073"
	return _dcdd(rs, opts, true, _dfdcb)
}

// DecodeArray returns the range of color component values in DeviceRGB colorspace.
func (_fgda *PdfColorspaceDeviceRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}
func _fgea(_dbd *PdfPage) map[_dg.PdfObjectName]_dg.PdfObject {
	_efac := make(map[_dg.PdfObjectName]_dg.PdfObject)
	if _dbd.Resources == nil {
		return _efac
	}
	if _dbd.Resources.Font != nil {
		if _cdfa, _afbda := _dg.GetDict(_dbd.Resources.Font); _afbda {
			for _, _befb := range _cdfa.Keys() {
				_efac[_befb] = _cdfa.Get(_befb)
			}
		}
	}
	if _dbd.Resources.ExtGState != nil {
		if _fga, _ebag := _dg.GetDict(_dbd.Resources.ExtGState); _ebag {
			for _, _gdec := range _fga.Keys() {
				_efac[_gdec] = _fga.Get(_gdec)
			}
		}
	}
	if _dbd.Resources.XObject != nil {
		if _cfbc, _bfddb := _dg.GetDict(_dbd.Resources.XObject); _bfddb {
			for _, _gfeb := range _cfbc.Keys() {
				_efac[_gfeb] = _cfbc.Get(_gfeb)
			}
		}
	}
	if _dbd.Resources.Pattern != nil {
		if _fbcg, _dcge := _dg.GetDict(_dbd.Resources.Pattern); _dcge {
			for _, _adbe := range _fbcg.Keys() {
				_efac[_adbe] = _fbcg.Get(_adbe)
			}
		}
	}
	if _dbd.Resources.Shading != nil {
		if _gagg, _afad := _dg.GetDict(_dbd.Resources.Shading); _afad {
			for _, _gcfag := range _gagg.Keys() {
				_efac[_gcfag] = _gagg.Get(_gcfag)
			}
		}
	}
	if _dbd.Resources.ProcSet != nil {
		if _ddbd, _ccef := _dg.GetDict(_dbd.Resources.ProcSet); _ccef {
			for _, _aggg := range _ddbd.Keys() {
				_efac[_aggg] = _ddbd.Get(_aggg)
			}
		}
	}
	if _dbd.Resources.Properties != nil {
		if _agac, _abbc := _dg.GetDict(_dbd.Resources.Properties); _abbc {
			for _, _cccce := range _agac.Keys() {
				_efac[_cccce] = _agac.Get(_cccce)
			}
		}
	}
	return _efac
}
func (_dfabae *PdfWriter) adjustXRefAffectedVersion(_cgggd bool) {
	if _cgggd && _dfabae._efacd.Major == 1 && _dfabae._efacd.Minor < 5 {
		_dfabae._efacd.Minor = 5
	}
}
func (_gega *PdfFont) baseFields() *fontCommon {
	if _gega._cadf == nil {
		_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0062\u0061\u0073\u0065\u0046\u0069\u0065l\u0064s\u002e \u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e")
		return nil
	}
	return _gega._cadf.baseFields()
}
func (_eeagd *pdfFontType0) baseFields() *fontCommon { return &_eeagd.fontCommon }

// Set applies flag fl to the flag's bitmask and returns the combined flag.
func (_gaga FieldFlag) Set(fl FieldFlag) FieldFlag { return FieldFlag(_gaga.Mask() | fl.Mask()) }

// Encoder returns the font's text encoder.
func (_cgace *pdfFontSimple) Encoder() _bd.TextEncoder {
	if _cgace._bdeed != nil {
		return _cgace._bdeed
	}
	if _cgace._dbdb != nil {
		return _cgace._dbdb
	}
	_aefdg, _ := _bd.NewSimpleTextEncoder("\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", nil)
	return _aefdg
}

// CharMetrics represents width and height metrics of a glyph.
type CharMetrics = _bbg.CharMetrics

// HasFontByName checks if has font resource by name.
func (_gfgd *PdfPage) HasFontByName(name _dg.PdfObjectName) bool {
	_faccg, _bdefd := _gfgd.Resources.Font.(*_dg.PdfObjectDictionary)
	if !_bdefd {
		return false
	}
	if _adffg := _faccg.Get(name); _adffg != nil {
		return true
	}
	return false
}

// PdfColorspaceDeviceN represents a DeviceN color space. DeviceN color spaces are similar to Separation color
// spaces, except they can contain an arbitrary number of color components.
/*
	Format: [/DeviceN names alternateSpace tintTransform]
        or: [/DeviceN names alternateSpace tintTransform attributes]
*/
type PdfColorspaceDeviceN struct {
	ColorantNames  *_dg.PdfObjectArray
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	Attributes     *PdfColorspaceDeviceNAttributes
	_effag         *_dg.PdfIndirectObject
}

func (_fegbc *pdfFontSimple) getFontDescriptor() *PdfFontDescriptor {
	if _cccf := _fegbc._ccfb; _cccf != nil {
		return _cccf
	}
	return _fegbc._gagbe
}

// NewPdfAnnotationFreeText returns a new free text annotation.
func NewPdfAnnotationFreeText() *PdfAnnotationFreeText {
	_fbd := NewPdfAnnotation()
	_fdb := &PdfAnnotationFreeText{}
	_fdb.PdfAnnotation = _fbd
	_fdb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fbd.SetContext(_fdb)
	return _fdb
}

// M returns the value of the magenta component of the color.
func (_cfef *PdfColorDeviceCMYK) M() float64 { return _cfef[1] }

// FlattenFields flattens the form fields and annotations for the PDF loaded in `pdf` and makes
// non-editable.
// Looks up all widget annotations corresponding to form fields and flattens them by drawing the content
// through the content stream rather than annotations.
// References to flattened annotations will be removed from Page Annots array. For fields the AcroForm entry
// will be emptied.
// When `allannots` is true, all annotations will be flattened. Keep false if want to keep non-form related
// annotations intact.
// When `appgen` is not nil, it will be used to generate appearance streams for the field annotations.
func (_dagd *PdfReader) FlattenFields(allannots bool, appgen FieldAppearanceGenerator) error {
	return _dagd.flattenFieldsWithOpts(allannots, appgen, nil)
}

// Encoder returns the font's text encoder.
func (_eage pdfFontType3) Encoder() _bd.TextEncoder { return _eage._geffb }

// NewPdfShadingPatternType3 creates an empty shading pattern type 3 object.
func NewPdfShadingPatternType3() *PdfShadingPatternType3 {
	_faeaf := &PdfShadingPatternType3{}
	_faeaf.Matrix = _dg.MakeArrayFromIntegers([]int{1, 0, 0, 1, 0, 0})
	_faeaf.PdfPattern = &PdfPattern{}
	_faeaf.PdfPattern.PatternType = int64(*_dg.MakeInteger(2))
	_faeaf.PdfPattern._cgdcc = _faeaf
	_faeaf.PdfPattern._eacce = _dg.MakeIndirectObject(_dg.MakeDict())
	return _faeaf
}

// ToPdfObject converts the pdfCIDFontType0 to a PDF representation.
func (_fefb *pdfCIDFontType0) ToPdfObject() _dg.PdfObject { return _dg.MakeNull() }

// ToPdfObject implements interface PdfModel.
func (_ade *PdfActionResetForm) ToPdfObject() _dg.PdfObject {
	_ade.PdfAction.ToPdfObject()
	_eae := _ade._cbd
	_bfg := _eae.PdfObject.(*_dg.PdfObjectDictionary)
	_bfg.SetIfNotNil("\u0053", _dg.MakeName(string(ActionTypeResetForm)))
	_bfg.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _ade.Fields)
	_bfg.SetIfNotNil("\u0046\u006c\u0061g\u0073", _ade.Flags)
	return _eae
}

// NewPdfAnnotationStamp returns a new stamp annotation.
func NewPdfAnnotationStamp() *PdfAnnotationStamp {
	_cff := NewPdfAnnotation()
	_bfdd := &PdfAnnotationStamp{}
	_bfdd.PdfAnnotation = _cff
	_bfdd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cff.SetContext(_bfdd)
	return _bfdd
}

// PdfShadingPatternType2 is shading patterns that will use a Type 2 shading pattern (Axial).
type PdfShadingPatternType2 struct {
	*PdfPattern
	Shading   *PdfShadingType2
	Matrix    *_dg.PdfObjectArray
	ExtGState _dg.PdfObject
}

// ToGoTime returns the date in time.Time format.
func (_ebadb PdfDate) ToGoTime() _a.Time {
	_gcdee := int(_ebadb._bdde*60*60 + _ebadb._gafe*60)
	switch _ebadb._ggbdg {
	case '-':
		_gcdee = -_gcdee
	case 'Z':
		_gcdee = 0
	}
	_adeac := _b.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _ebadb._ggbdg, _ebadb._bdde, _ebadb._gafe)
	_eaecc := _a.FixedZone(_adeac, _gcdee)
	return _a.Date(int(_ebadb._bgfdb), _a.Month(_ebadb._gbbge), int(_ebadb._cbad), int(_ebadb._cade), int(_ebadb._ccgbg), int(_ebadb._aacfe), 0, _eaecc)
}
func _eceba(_bcgeb *_dg.PdfObjectDictionary) *VRI {
	_ecde, _ := _dg.GetString(_bcgeb.Get("\u0054\u0055"))
	_adgcc, _ := _dg.GetString(_bcgeb.Get("\u0054\u0053"))
	return &VRI{Cert: _eeafe(_bcgeb.Get("\u0043\u0065\u0072\u0074")), OCSP: _eeafe(_bcgeb.Get("\u004f\u0043\u0053\u0050")), CRL: _eeafe(_bcgeb.Get("\u0043\u0052\u004c")), TU: _ecde, TS: _adgcc}
}

// GetCustomInfo returns a custom info value for the specified name.
func (_aacc *PdfInfo) GetCustomInfo(name string) *_dg.PdfObjectString {
	var _egfdd *_dg.PdfObjectString
	if _aacc._dccdg == nil {
		return _egfdd
	}
	if _ggbf, _fcdbb := _aacc._dccdg.Get(*_dg.MakeName(name)).(*_dg.PdfObjectString); _fcdbb {
		_egfdd = _ggbf
	}
	return _egfdd
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a Lab device.
func (_fcdfd *PdfColorspaceLab) GetNumComponents() int { return 3 }

// ToPdfObject implements interface PdfModel.
func (_ecbgd *PdfAnnotationHighlight) ToPdfObject() _dg.PdfObject {
	_ecbgd.PdfAnnotation.ToPdfObject()
	_dacd := _ecbgd._cdf
	_eggcg := _dacd.PdfObject.(*_dg.PdfObjectDictionary)
	_ecbgd.PdfAnnotationMarkup.appendToPdfDictionary(_eggcg)
	_eggcg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t"))
	_eggcg.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _ecbgd.QuadPoints)
	return _dacd
}

// PdfShadingType4 is a Free-form Gouraud-shaded triangle mesh.
type PdfShadingType4 struct {
	*PdfShading
	BitsPerCoordinate *_dg.PdfObjectInteger
	BitsPerComponent  *_dg.PdfObjectInteger
	BitsPerFlag       *_dg.PdfObjectInteger
	Decode            *_dg.PdfObjectArray
	Function          []PdfFunction
}

// ToPdfObject sets the common field elements.
// Note: Call the more field context's ToPdfObject to set both the generic and
// non-generic information.
func (_ddef *PdfField) ToPdfObject() _dg.PdfObject {
	_afcb := _ddef._egce
	_ebee := _afcb.PdfObject.(*_dg.PdfObjectDictionary)
	_gbaec := _dg.MakeArray()
	for _, _eafd := range _ddef.Kids {
		_gbaec.Append(_eafd.ToPdfObject())
	}
	for _, _efdd := range _ddef.Annotations {
		if _efdd._cdf != _ddef._egce {
			_gbaec.Append(_efdd.GetContext().ToPdfObject())
		}
	}
	if _ddef.Parent != nil {
		_ebee.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _ddef.Parent.GetContainingPdfObject())
	}
	if _gbaec.Len() > 0 {
		_ebee.Set("\u004b\u0069\u0064\u0073", _gbaec)
	}
	_ebee.SetIfNotNil("\u0046\u0054", _ddef.FT)
	_ebee.SetIfNotNil("\u0054", _ddef.T)
	_ebee.SetIfNotNil("\u0054\u0055", _ddef.TU)
	_ebee.SetIfNotNil("\u0054\u004d", _ddef.TM)
	_ebee.SetIfNotNil("\u0046\u0066", _ddef.Ff)
	_ebee.SetIfNotNil("\u0056", _ddef.V)
	_ebee.SetIfNotNil("\u0044\u0056", _ddef.DV)
	_ebee.SetIfNotNil("\u0041\u0041", _ddef.AA)
	if _ddef.VariableText != nil {
		_ebee.SetIfNotNil("\u0044\u0041", _ddef.VariableText.DA)
		_ebee.SetIfNotNil("\u0051", _ddef.VariableText.Q)
		_ebee.SetIfNotNil("\u0044\u0053", _ddef.VariableText.DS)
		_ebee.SetIfNotNil("\u0052\u0056", _ddef.VariableText.RV)
	}
	return _afcb
}

// NewPdfInfoFromObject creates a new PdfInfo from the input core.PdfObject.
func NewPdfInfoFromObject(obj _dg.PdfObject) (*PdfInfo, error) {
	var _aggcd PdfInfo
	_bafcd, _fcbc := obj.(*_dg.PdfObjectDictionary)
	if !_fcbc {
		return nil, _b.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", obj)
	}
	for _, _ffag := range _bafcd.Keys() {
		switch _ffag {
		case "\u0054\u0069\u0074l\u0065":
			_aggcd.Title, _ = _dg.GetString(_bafcd.Get("\u0054\u0069\u0074l\u0065"))
		case "\u0041\u0075\u0074\u0068\u006f\u0072":
			_aggcd.Author, _ = _dg.GetString(_bafcd.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
		case "\u0053u\u0062\u006a\u0065\u0063\u0074":
			_aggcd.Subject, _ = _dg.GetString(_bafcd.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
		case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
			_aggcd.Keywords, _ = _dg.GetString(_bafcd.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
		case "\u0043r\u0065\u0061\u0074\u006f\u0072":
			_aggcd.Creator, _ = _dg.GetString(_bafcd.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
		case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
			_aggcd.Producer, _ = _dg.GetString(_bafcd.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
		case "\u0054r\u0061\u0070\u0070\u0065\u0064":
			_aggcd.Trapped, _ = _dg.GetName(_bafcd.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
		case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
			if _dedeg, _ebca := _dg.GetString(_bafcd.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _ebca && _dedeg.String() != "" {
				_fdfb, _cadc := NewPdfDate(_dedeg.String())
				if _cadc != nil {
					return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _cadc)
				}
				_aggcd.CreationDate = &_fdfb
			}
		case "\u004do\u0064\u0044\u0061\u0074\u0065":
			if _cbed, _gedfa := _dg.GetString(_bafcd.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _gedfa && _cbed.String() != "" {
				_cfgg, _aegc := NewPdfDate(_cbed.String())
				if _aegc != nil {
					return nil, _b.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _aegc)
				}
				_aggcd.ModifiedDate = &_cfgg
			}
		default:
			_bacd, _ := _dg.GetString(_bafcd.Get(_ffag))
			if _aggcd._dccdg == nil {
				_aggcd._dccdg = _dg.MakeDict()
			}
			_aggcd._dccdg.Set(_ffag, _bacd)
		}
	}
	return &_aggcd, nil
}
func (_agaf *PdfColorspaceICCBased) String() string {
	return "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"
}
func (_badeef *pdfCIDFontType2) getFontDescriptor() *PdfFontDescriptor { return _badeef._ccfb }
func _dgdbe() string                                                   { _fgefgf.Lock(); defer _fgefgf.Unlock(); return _bdeaf }

// ToPdfObject returns the PDF representation of the function.
func (_acdcg *PdfFunctionType2) ToPdfObject() _dg.PdfObject {
	_dgaf := _dg.MakeDict()
	_dgaf.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _dg.MakeInteger(2))
	_eggb := &_dg.PdfObjectArray{}
	for _, _gggc := range _acdcg.Domain {
		_eggb.Append(_dg.MakeFloat(_gggc))
	}
	_dgaf.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _eggb)
	if _acdcg.Range != nil {
		_faged := &_dg.PdfObjectArray{}
		for _, _egega := range _acdcg.Range {
			_faged.Append(_dg.MakeFloat(_egega))
		}
		_dgaf.Set("\u0052\u0061\u006eg\u0065", _faged)
	}
	if _acdcg.C0 != nil {
		_defcd := &_dg.PdfObjectArray{}
		for _, _bcbba := range _acdcg.C0 {
			_defcd.Append(_dg.MakeFloat(_bcbba))
		}
		_dgaf.Set("\u0043\u0030", _defcd)
	}
	if _acdcg.C1 != nil {
		_gaafg := &_dg.PdfObjectArray{}
		for _, _fgedc := range _acdcg.C1 {
			_gaafg.Append(_dg.MakeFloat(_fgedc))
		}
		_dgaf.Set("\u0043\u0031", _gaafg)
	}
	_dgaf.Set("\u004e", _dg.MakeFloat(_acdcg.N))
	if _acdcg._acfde != nil {
		_acdcg._acfde.PdfObject = _dgaf
		return _acdcg._acfde
	}
	return _dgaf
}
func (_ebgg *PdfReader) newPdfAnnotationScreenFromDict(_gadb *_dg.PdfObjectDictionary) (*PdfAnnotationScreen, error) {
	_bdac := PdfAnnotationScreen{}
	_bdac.T = _gadb.Get("\u0054")
	_bdac.MK = _gadb.Get("\u004d\u004b")
	_bdac.A = _gadb.Get("\u0041")
	_bdac.AA = _gadb.Get("\u0041\u0041")
	return &_bdac, nil
}

// NewPdfColorPatternType3 returns an empty color shading pattern type 3 (Radial).
func NewPdfColorPatternType3() *PdfColorPatternType3 { _egbc := &PdfColorPatternType3{}; return _egbc }
func (_eddgb *PdfReader) newPdfSignatureFromIndirect(_cece *_dg.PdfIndirectObject) (*PdfSignature, error) {
	_dcgad, _eefca := _cece.PdfObject.(*_dg.PdfObjectDictionary)
	if !_eefca {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		return nil, ErrTypeCheck
	}
	if _aceea, _defeb := _eddgb._cadfa.GetModelFromPrimitive(_cece).(*PdfSignature); _defeb {
		return _aceea, nil
	}
	_gagac := &PdfSignature{}
	_gagac._bbda = _cece
	_gagac.Type, _ = _dg.GetName(_dcgad.Get("\u0054\u0079\u0070\u0065"))
	_gagac.Filter, _eefca = _dg.GetName(_dcgad.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if !_eefca {
		_ag.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053i\u0067\u006e\u0061\u0074\u0075r\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_gagac.SubFilter, _ = _dg.GetName(_dcgad.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r"))
	_gagac.Contents, _eefca = _dg.GetString(_dcgad.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_eefca {
		_ag.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	if _ebdde, _gabed := _dg.GetArray(_dcgad.Get("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e")); _gabed {
		_gagac.Reference = _dg.MakeArray()
		for _, _gebcf := range _ebdde.Elements() {
			_eeegg, _aaade := _dg.GetDict(_gebcf)
			if !_aaade {
				_ag.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020R\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0061\u0074\u0065\u0064")
				return nil, ErrInvalidAttribute
			}
			_edcfgg, _cggbb := _eddgb.newPdfSignatureReferenceFromDict(_eeegg)
			if _cggbb != nil {
				return nil, _cggbb
			}
			_gagac.Reference.Append(_edcfgg.ToPdfObject())
		}
	}
	_gagac.Cert = _dcgad.Get("\u0043\u0065\u0072\u0074")
	_gagac.ByteRange, _ = _dg.GetArray(_dcgad.Get("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e"))
	_gagac.Changes, _ = _dg.GetArray(_dcgad.Get("\u0043h\u0061\u006e\u0067\u0065\u0073"))
	_gagac.Name, _ = _dg.GetString(_dcgad.Get("\u004e\u0061\u006d\u0065"))
	_gagac.M, _ = _dg.GetString(_dcgad.Get("\u004d"))
	_gagac.Location, _ = _dg.GetString(_dcgad.Get("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"))
	_gagac.Reason, _ = _dg.GetString(_dcgad.Get("\u0052\u0065\u0061\u0073\u006f\u006e"))
	_gagac.ContactInfo, _ = _dg.GetString(_dcgad.Get("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f"))
	_gagac.R, _ = _dg.GetInt(_dcgad.Get("\u0052"))
	_gagac.V, _ = _dg.GetInt(_dcgad.Get("\u0056"))
	_gagac.PropBuild, _ = _dg.GetDict(_dcgad.Get("\u0050\u0072\u006f\u0070\u005f\u0042\u0075\u0069\u006c\u0064"))
	_gagac.PropAuthTime, _ = _dg.GetInt(_dcgad.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0069\u006d\u0065"))
	_gagac.PropAuthType, _ = _dg.GetName(_dcgad.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0079\u0070\u0065"))
	_eddgb._cadfa.Register(_cece, _gagac)
	return _gagac, nil
}
func _ebccc(_adfb *fontCommon) *pdfCIDFontType2 { return &pdfCIDFontType2{fontCommon: *_adfb} }
func _gcab(_caec *PdfField, _caaebb _dg.PdfObject) {
	for _, _dfea := range _caec.Annotations {
		_dfea.AS = _caaebb
		_dfea.ToPdfObject()
	}
}

// Insert adds an outline item as a child of the current outline item,
// at the specified index.
func (_agcce *OutlineItem) Insert(index uint, item *OutlineItem) {
	_dffb := uint(len(_agcce.Entries))
	if index > _dffb {
		index = _dffb
	}
	_agcce.Entries = append(_agcce.Entries[:index], append([]*OutlineItem{item}, _agcce.Entries[index:]...)...)
}

// SetContext sets the sub pattern (context).  Either PdfTilingPattern or PdfShadingPattern.
func (_ecbgc *PdfPattern) SetContext(ctx PdfModel) { _ecbgc._cgdcc = ctx }

// PdfAnnotationCircle represents Circle annotations.
// (Section 12.5.6.8).
type PdfAnnotationCircle struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _dg.PdfObject
	IC _dg.PdfObject
	BE _dg.PdfObject
	RD _dg.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_affge *Permissions) ToPdfObject() _dg.PdfObject { return _affge._fbdaa }

// NewPdfColorspaceDeviceN returns an initialized PdfColorspaceDeviceN.
func NewPdfColorspaceDeviceN() *PdfColorspaceDeviceN {
	_dceed := &PdfColorspaceDeviceN{}
	return _dceed
}
func (_acdf *PdfWriter) optimizeDocument() error {
	if _acdf._afafd == nil {
		return nil
	}
	_agdfd, _faefc := _dg.GetDict(_acdf._efbfa)
	if !_faefc {
		return _bf.New("\u0061\u006e\u0020in\u0066\u006f\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_eecdb := _eba.Document{ID: [2]string{_acdf._cedaf, _acdf._ceaab}, Version: _acdf._efacd, Objects: _acdf._agaba, Info: _agdfd, Crypt: _acdf._fadcg, UseHashBasedID: _acdf._affea}
	if _bdfdd := _acdf._afafd.ApplyStandard(&_eecdb); _bdfdd != nil {
		return _bdfdd
	}
	_acdf._cedaf, _acdf._ceaab = _eecdb.ID[0], _eecdb.ID[1]
	_acdf._efacd = _eecdb.Version
	_acdf._agaba = _eecdb.Objects
	_acdf._efbfa.PdfObject = _eecdb.Info
	_acdf._affea = _eecdb.UseHashBasedID
	_acdf._fadcg = _eecdb.Crypt
	_dfgfb := make(map[_dg.PdfObject]struct{}, len(_acdf._agaba))
	for _, _bdddf := range _acdf._agaba {
		_dfgfb[_bdddf] = struct{}{}
	}
	_acdf._fdbfa = _dfgfb
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_fgdde *PdfAnnotationPrinterMark) ToPdfObject() _dg.PdfObject {
	_fgdde.PdfAnnotation.ToPdfObject()
	_fcba := _fgdde._cdf
	_ddgd := _fcba.PdfObject.(*_dg.PdfObjectDictionary)
	_ddgd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b"))
	_ddgd.SetIfNotNil("\u004d\u004e", _fgdde.MN)
	return _fcba
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_gcbd *PdfColorspaceCalGray) ColorFromPdfObjects(objects []_dg.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gdagf, _gecf := _dg.GetNumbersAsFloat(objects)
	if _gecf != nil {
		return nil, _gecf
	}
	return _gcbd.ColorFromFloats(_gdagf)
}
func _gdea(_ccdf _dg.PdfObject) (*PdfColorspaceCalRGB, error) {
	_ggcc := NewPdfColorspaceCalRGB()
	if _gbce, _ffffa := _ccdf.(*_dg.PdfIndirectObject); _ffffa {
		_ggcc._ccae = _gbce
	}
	_ccdf = _dg.TraceToDirectObject(_ccdf)
	_bbfg, _gcfd := _ccdf.(*_dg.PdfObjectArray)
	if !_gcfd {
		return nil, _b.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _bbfg.Len() != 2 {
		return nil, _b.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_ccdf = _dg.TraceToDirectObject(_bbfg.Get(0))
	_addd, _gcfd := _ccdf.(*_dg.PdfObjectName)
	if !_gcfd {
		return nil, _b.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	if *_addd != "\u0043\u0061\u006c\u0052\u0047\u0042" {
		return nil, _b.Errorf("\u006e\u006f\u0074 a\u0020\u0043\u0061\u006c\u0052\u0047\u0042\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_ccdf = _dg.TraceToDirectObject(_bbfg.Get(1))
	_afef, _gcfd := _ccdf.(*_dg.PdfObjectDictionary)
	if !_gcfd {
		return nil, _b.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	_ccdf = _afef.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_ccdf = _dg.TraceToDirectObject(_ccdf)
	_aece, _gcfd := _ccdf.(*_dg.PdfObjectArray)
	if !_gcfd {
		return nil, _b.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050o\u0069\u006e\u0074")
	}
	if _aece.Len() != 3 {
		return nil, _b.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0057h\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_fabcd, _bgfd := _aece.GetAsFloat64Slice()
	if _bgfd != nil {
		return nil, _bgfd
	}
	_ggcc.WhitePoint = _fabcd
	_ccdf = _afef.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _ccdf != nil {
		_ccdf = _dg.TraceToDirectObject(_ccdf)
		_dddea, _fgee := _ccdf.(*_dg.PdfObjectArray)
		if !_fgee {
			return nil, _b.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050o\u0069\u006e\u0074")
		}
		if _dddea.Len() != 3 {
			return nil, _b.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0042l\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		_cagdd, _gedfd := _dddea.GetAsFloat64Slice()
		if _gedfd != nil {
			return nil, _gedfd
		}
		_ggcc.BlackPoint = _cagdd
	}
	_ccdf = _afef.Get("\u0047\u0061\u006dm\u0061")
	if _ccdf != nil {
		_ccdf = _dg.TraceToDirectObject(_ccdf)
		_cfcb, _eggcd := _ccdf.(*_dg.PdfObjectArray)
		if !_eggcd {
			return nil, _b.Errorf("C\u0061\u006c\u0052\u0047B:\u0020I\u006e\u0076\u0061\u006c\u0069d\u0020\u0047\u0061\u006d\u006d\u0061")
		}
		if _cfcb.Len() != 3 {
			return nil, _b.Errorf("C\u0061\u006c\u0052\u0047\u0042\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0047a\u006d\u006d\u0061 \u0061r\u0072\u0061\u0079")
		}
		_cae, _cgdb := _cfcb.GetAsFloat64Slice()
		if _cgdb != nil {
			return nil, _cgdb
		}
		_ggcc.Gamma = _cae
	}
	_ccdf = _afef.Get("\u004d\u0061\u0074\u0072\u0069\u0078")
	if _ccdf != nil {
		_ccdf = _dg.TraceToDirectObject(_ccdf)
		_eaf, _cfeac := _ccdf.(*_dg.PdfObjectArray)
		if !_cfeac {
			return nil, _b.Errorf("\u0043\u0061\u006c\u0052GB\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004d\u0061\u0074\u0072i\u0078")
		}
		if _eaf.Len() != 9 {
			_ag.Log.Error("\u004d\u0061t\u0072\u0069\u0078 \u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _eaf.String())
			return nil, _b.Errorf("\u0043\u0061\u006c\u0052G\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u004da\u0074\u0072\u0069\u0078\u0020\u0061\u0072r\u0061\u0079")
		}
		_eefa, _ebcfb := _eaf.GetAsFloat64Slice()
		if _ebcfb != nil {
			return nil, _ebcfb
		}
		_ggcc.Matrix = _eefa
	}
	return _ggcc, nil
}

// GetPageLabels returns the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_cecad *PdfReader) GetPageLabels() (_dg.PdfObject, error) {
	_ecbfd := _dg.ResolveReference(_cecad._gccfb.Get("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073"))
	if _ecbfd == nil {
		return nil, nil
	}
	if !_cecad._dadcef {
		_cgga := _cecad.traverseObjectData(_ecbfd)
		if _cgga != nil {
			return nil, _cgga
		}
	}
	return _ecbfd, nil
}

// StandardApplier is the interface that performs optimization of the whole PDF document.
// As a result an input document is being changed by the optimizer.
// The writer than takes back all it's parts and overwrites it.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardApplier interface {
	ApplyStandard(_cgeg *_eba.Document) error
}

// AddImageResource adds an image to the XObject resources.
func (_gdaee *PdfPage) AddImageResource(name _dg.PdfObjectName, ximg *XObjectImage) error {
	var _bceg *_dg.PdfObjectDictionary
	if _gdaee.Resources.XObject == nil {
		_bceg = _dg.MakeDict()
		_gdaee.Resources.XObject = _bceg
	} else {
		var _ggged bool
		_bceg, _ggged = (_gdaee.Resources.XObject).(*_dg.PdfObjectDictionary)
		if !_ggged {
			return _bf.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0078\u0072\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020\u0074\u0079p\u0065")
		}
	}
	_bceg.Set(name, ximg.ToPdfObject())
	return nil
}

// CharcodesToStrings returns the unicode strings corresponding to `charcodes`.
// The int returns are the number of strings and the number of unconvereted codes.
// NOTE: The number of strings returned is equal to the number of charcodes
func (_eebc *PdfFont) CharcodesToStrings(charcodes []_bd.CharCode) ([]string, int, int) {
	_eeaf := _eebc.baseFields()
	_bffd := make([]string, 0, len(charcodes))
	_ggeaf := 0
	_abdaa := _eebc.Encoder()
	_gfcd := _eeaf._ecfb != nil && _eebc.IsSimple() && _eebc.Subtype() == "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" && !_ga.Contains(_eeaf._ecfb.Name(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")
	if !_gfcd && _abdaa != nil {
		switch _gbdf := _abdaa.(type) {
		case _bd.SimpleEncoder:
			_faege := _gbdf.BaseName()
			if _, _ecgbc := _eabb[_faege]; _ecgbc {
				for _, _gacdda := range charcodes {
					if _fdfab, _ccbge := _abdaa.CharcodeToRune(_gacdda); _ccbge {
						_bffd = append(_bffd, string(_fdfab))
					} else {
						_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _gacdda, charcodes, _eeaf.isCIDFont(), _eebc, _abdaa)
						_ggeaf++
						_bffd = append(_bffd, _ff.MissingCodeString)
					}
				}
				return _bffd, len(_bffd), _ggeaf
			}
		}
	}
	for _, _eafdb := range charcodes {
		if _eeaf._ecfb != nil {
			if _bgdaf, _ccaa := _eeaf._ecfb.CharcodeToUnicode(_ff.CharCode(_eafdb)); _ccaa {
				_bffd = append(_bffd, _bgdaf)
				continue
			}
		}
		if _abdaa != nil {
			if _gbfgg, _edgfg := _abdaa.CharcodeToRune(_eafdb); _edgfg {
				_bffd = append(_bffd, string(_gbfgg))
				continue
			}
		}
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _eafdb, charcodes, _eeaf.isCIDFont(), _eebc, _abdaa)
		_ggeaf++
		_bffd = append(_bffd, _ff.MissingCodeString)
	}
	if _ggeaf != 0 {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0043\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0074\u006f\u0020u\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020i\u006ep\u0075\u0074\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0043\u0068\u0061\u0072\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073", len(charcodes), _ggeaf, _eebc)
	}
	return _bffd, len(_bffd), _ggeaf
}
func (_gebeg *PdfWriter) writeOutlines() error {
	if _gebeg._cedcg == nil {
		return nil
	}
	_ag.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0054\u0072\u0065\u0065\u003a\u0020\u0025\u002b\u0076", _gebeg._cedcg)
	_cfgfg := _gebeg._cedcg.ToPdfObject()
	_ag.Log.Trace("\u004fu\u0074\u006c\u0069\u006e\u0065\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0025\u0054\u002c\u0020\u0070\u003a\u0025\u0070\u0029", _cfgfg, _cfgfg, _cfgfg)
	_gebeg._ecdf.Set("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073", _cfgfg)
	_dfbdf := _gebeg.addObjects(_cfgfg)
	if _dfbdf != nil {
		return _dfbdf
	}
	return nil
}

// SetDocInfo set document info.
// This will overwrite any globally declared document info.
func (_bbca *PdfWriter) SetDocInfo(info *PdfInfo) { _bbca.setDocInfo(info.ToPdfObject()) }

// ToPdfObject implements interface PdfModel.
func (_aec *PdfAnnotationTrapNet) ToPdfObject() _dg.PdfObject {
	_aec.PdfAnnotation.ToPdfObject()
	_bef := _aec._cdf
	_egfd := _bef.PdfObject.(*_dg.PdfObjectDictionary)
	_egfd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0054r\u0061\u0070\u004e\u0065\u0074"))
	return _bef
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// red, green and blue components of the color. The values of the elements
// should be between 0 and 1.
func (_baeg *PdfColorspaceDeviceRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bdab := vals[0]
	if _bdab < 0.0 || _bdab > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bdab)
		return nil, ErrColorOutOfRange
	}
	_aaee := vals[1]
	if _aaee < 0.0 || _aaee > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bdab)
		return nil, ErrColorOutOfRange
	}
	_accc := vals[2]
	if _accc < 0.0 || _accc > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bdab)
		return nil, ErrColorOutOfRange
	}
	_ddfb := NewPdfColorDeviceRGB(_bdab, _aaee, _accc)
	return _ddfb, nil
}
func (_cdgf *pdfFontType0) getFontDescriptor() *PdfFontDescriptor {
	if _cdgf._ccfb == nil && _cdgf.DescendantFont != nil {
		return _cdgf.DescendantFont.FontDescriptor()
	}
	return _cdgf._ccfb
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_abbbg *PdfShadingType7) ToPdfObject() _dg.PdfObject {
	_abbbg.PdfShading.ToPdfObject()
	_ebbg, _geagg := _abbbg.getShadingDict()
	if _geagg != nil {
		_ag.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _abbbg.BitsPerCoordinate != nil {
		_ebbg.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _abbbg.BitsPerCoordinate)
	}
	if _abbbg.BitsPerComponent != nil {
		_ebbg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _abbbg.BitsPerComponent)
	}
	if _abbbg.BitsPerFlag != nil {
		_ebbg.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _abbbg.BitsPerFlag)
	}
	if _abbbg.Decode != nil {
		_ebbg.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _abbbg.Decode)
	}
	if _abbbg.Function != nil {
		if len(_abbbg.Function) == 1 {
			_ebbg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _abbbg.Function[0].ToPdfObject())
		} else {
			_agabc := _dg.MakeArray()
			for _, _eegfc := range _abbbg.Function {
				_agabc.Append(_eegfc.ToPdfObject())
			}
			_ebbg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _agabc)
		}
	}
	return _abbbg._bcfbg
}

// Transform rectangle with the supplied matrix.
func (_dffaa *PdfRectangle) Transform(transformMatrix _fec.Matrix) {
	_dffaa.Llx, _dffaa.Lly = transformMatrix.Transform(_dffaa.Llx, _dffaa.Lly)
	_dffaa.Urx, _dffaa.Ury = transformMatrix.Transform(_dffaa.Urx, _dffaa.Ury)
	_dffaa.Normalize()
}

// GetDSS gets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_dbaf *PdfAppender) GetDSS() (_agf *DSS) { return _dbaf._cfdd }

// SignatureHandlerDocMDPParams describe the specific parameters for the SignatureHandlerEx
// These parameters describe how to check the difference between revisions.
// Revisions of the document get from the PdfParser.
type SignatureHandlerDocMDPParams struct {
	Parser     *_dg.PdfParser
	DiffPolicy _ecb.DiffPolicy
}

const (
	XObjectTypeUndefined XObjectType = iota
	XObjectTypeImage
	XObjectTypeForm
	XObjectTypePS
	XObjectTypeUnknown
)

// GetCharMetrics returns the char metrics for character code `code`.
func (_dfcbf pdfFontType0) GetCharMetrics(code _bd.CharCode) (_bbg.CharMetrics, bool) {
	if _dfcbf.DescendantFont == nil {
		_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _dfcbf)
		return _bbg.CharMetrics{}, false
	}
	return _dfcbf.DescendantFont.GetCharMetrics(code)
}
func (_bdfd *PdfColorspaceLab) String() string { return "\u004c\u0061\u0062" }
func (_babg *PdfReader) newPdfAnnotationCaretFromDict(_edca *_dg.PdfObjectDictionary) (*PdfAnnotationCaret, error) {
	_daff := PdfAnnotationCaret{}
	_fefg, _ecce := _babg.newPdfAnnotationMarkupFromDict(_edca)
	if _ecce != nil {
		return nil, _ecce
	}
	_daff.PdfAnnotationMarkup = _fefg
	_daff.RD = _edca.Get("\u0052\u0044")
	_daff.Sy = _edca.Get("\u0053\u0079")
	return &_daff, nil
}
func (_cgafg *pdfFontType3) getFontDescriptor() *PdfFontDescriptor { return _cgafg._ccfb }
func _ffeac() _a.Time                                              { _fgefgf.Lock(); defer _fgefgf.Unlock(); return _efafac }

// NewPdfAppenderWithOpts creates a new Pdf appender from a Pdf reader with options.
func NewPdfAppenderWithOpts(reader *PdfReader, opts *ReaderOpts, encryptOptions *EncryptOptions) (*PdfAppender, error) {
	_baa := &PdfAppender{_cga: reader._efcfa, Reader: reader, _gfab: reader._baad, _ebaa: reader._addfg}
	_beecb, _bcadd := _baa._cga.Seek(0, _cf.SeekEnd)
	if _bcadd != nil {
		return nil, _bcadd
	}
	_baa._dcaf = _beecb
	if _, _bcadd = _baa._cga.Seek(0, _cf.SeekStart); _bcadd != nil {
		return nil, _bcadd
	}
	_baa._debg, _bcadd = NewPdfReaderWithOpts(_baa._cga, opts)
	if _bcadd != nil {
		return nil, _bcadd
	}
	for _, _gaff := range _baa.Reader.GetObjectNums() {
		if _baa._ceecc < _gaff {
			_baa._ceecc = _gaff
		}
	}
	_baa._ceef = _baa._gfab.GetXrefTable()
	_baa._bcdg = _baa._gfab.GetXrefOffset()
	_baa._ggdd = append(_baa._ggdd, _baa._debg.PageList...)
	_baa._efa = make(map[_dg.PdfObject]struct{})
	_baa._gebe = make(map[_dg.PdfObject]int64)
	_baa._eede = make(map[_dg.PdfObject]struct{})
	_baa._afab = _baa._debg.AcroForm
	_baa._cfdd = _baa._debg.DSS
	if opts != nil {
		_baa._acb = opts.Password
	}
	if encryptOptions != nil {
		_baa._ccfc = encryptOptions
	}
	return _baa, nil
}
func _gef(_egec *_dg.PdfObjectDictionary) bool {
	for _, _bbdb := range _egec.Keys() {
		if _, _efae := _fecec[_bbdb.String()]; _efae {
			return true
		}
	}
	return false
}

// PdfAnnotationLink represents Link annotations.
// (Section 12.5.6.5 p. 403).
type PdfAnnotationLink struct {
	*PdfAnnotation
	A          _dg.PdfObject
	Dest       _dg.PdfObject
	H          _dg.PdfObject
	PA         _dg.PdfObject
	QuadPoints _dg.PdfObject
	BS         _dg.PdfObject
	_ece       *PdfAction
	_abea      *PdfReader
}

// NewPdfColorspaceCalGray returns a new CalGray colorspace object.
func NewPdfColorspaceCalGray() *PdfColorspaceCalGray {
	_aaead := &PdfColorspaceCalGray{}
	_aaead.BlackPoint = []float64{0.0, 0.0, 0.0}
	_aaead.Gamma = 1
	return _aaead
}

// SetColorSpace sets `r` colorspace object to `colorspace`.
func (_fgdeb *PdfPageResources) SetColorSpace(colorspace *PdfPageResourcesColorspaces) {
	_fgdeb._dadae = colorspace
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_ggedg pdfFontType3) GetCharMetrics(code _bd.CharCode) (_bbg.CharMetrics, bool) {
	if _ddfc, _fbdc := _ggedg._bcde[code]; _fbdc {
		return _bbg.CharMetrics{Wx: _ddfc}, true
	}
	if _bbg.IsStdFont(_bbg.StdFontName(_ggedg._ecbf)) {
		return _bbg.CharMetrics{Wx: 250}, true
	}
	return _bbg.CharMetrics{}, false
}

var _ pdfFont = (*pdfFontType0)(nil)

// PdfActionType represents an action type in PDF (section 12.6.4 p. 417).
type PdfActionType string

// ColorFromFloats returns a new PdfColorDevice based on the input slice of
// color components. The slice should contain four elements representing the
// cyan, magenta, yellow and key components of the color. The values of the
// elements should be between 0 and 1.
func (_agb *PdfColorspaceDeviceCMYK) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 4 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ebbd := vals[0]
	if _ebbd < 0.0 || _ebbd > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _ebbd)
		return nil, ErrColorOutOfRange
	}
	_dbgcd := vals[1]
	if _dbgcd < 0.0 || _dbgcd > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _dbgcd)
		return nil, ErrColorOutOfRange
	}
	_debe := vals[2]
	if _debe < 0.0 || _debe > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _debe)
		return nil, ErrColorOutOfRange
	}
	_cda := vals[3]
	if _cda < 0.0 || _cda > 1.0 {
		_ag.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _cda)
		return nil, ErrColorOutOfRange
	}
	_ddcb := NewPdfColorDeviceCMYK(_ebbd, _dbgcd, _debe, _cda)
	return _ddcb, nil
}

// Y returns the value of the yellow component of the color.
func (_bcgc *PdfColorDeviceCMYK) Y() float64 { return _bcgc[2] }

// ToPdfObject returns the PDF representation of the VRI dictionary.
func (_adbg *VRI) ToPdfObject() *_dg.PdfObjectDictionary {
	_cdea := _dg.MakeDict()
	_cdea.SetIfNotNil(_dg.PdfObjectName("\u0043\u0065\u0072\u0074"), _fecee(_adbg.Cert))
	_cdea.SetIfNotNil(_dg.PdfObjectName("\u004f\u0043\u0053\u0050"), _fecee(_adbg.OCSP))
	_cdea.SetIfNotNil(_dg.PdfObjectName("\u0043\u0052\u004c"), _fecee(_adbg.CRL))
	_cdea.SetIfNotNil("\u0054\u0055", _adbg.TU)
	_cdea.SetIfNotNil("\u0054\u0053", _adbg.TS)
	return _cdea
}

// StringToCharcodeBytes maps the provided string runes to charcode bytes and
// it returns the resulting slice of bytes, along with the number of runes
// which could not be converted. If the number of misses is 0, all string runes
// were successfully converted.
func (_egffg *PdfFont) StringToCharcodeBytes(str string) ([]byte, int) {
	return _egffg.RunesToCharcodeBytes([]rune(str))
}

// WriteString outputs the object as it is to be written to file.
func (_gefc *pdfSignDictionary) WriteString() string {
	_gefc._aaeae = 0
	_gefc._gdbfe = 0
	_gefc._egda = 0
	_gefc._ebaaaa = 0
	_fdaga := _bc.NewBuffer(nil)
	_fdaga.WriteString("\u003c\u003c")
	for _, _egcff := range _gefc.Keys() {
		_dacbc := _gefc.Get(_egcff)
		switch _egcff {
		case "\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e":
			_fdaga.WriteString(_egcff.WriteString())
			_fdaga.WriteString("\u0020")
			_gefc._egda = _fdaga.Len()
			_fdaga.WriteString(_dacbc.WriteString())
			_fdaga.WriteString("\u0020")
			_gefc._ebaaaa = _fdaga.Len() - 1
		case "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073":
			_fdaga.WriteString(_egcff.WriteString())
			_fdaga.WriteString("\u0020")
			_gefc._aaeae = _fdaga.Len()
			_fdaga.WriteString(_dacbc.WriteString())
			_fdaga.WriteString("\u0020")
			_gefc._gdbfe = _fdaga.Len() - 1
		default:
			_fdaga.WriteString(_egcff.WriteString())
			_fdaga.WriteString("\u0020")
			_fdaga.WriteString(_dacbc.WriteString())
		}
	}
	_fdaga.WriteString("\u003e\u003e")
	return _fdaga.String()
}

// String implements interface PdfObject.
func (_gg *PdfAction) String() string {
	_df, _agc := _gg.ToPdfObject().(*_dg.PdfIndirectObject)
	if _agc {
		return _b.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _gg._bg, _df.PdfObject.String())
	}
	return ""
}

// DecodeArray returns the range of color component values in CalRGB colorspace.
func (_fea *PdfColorspaceCalRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}
func _fbbg(_geagc _dg.PdfObject) (*PdfColorspaceSpecialSeparation, error) {
	_debgd := NewPdfColorspaceSpecialSeparation()
	if _dgga, _fcbg := _geagc.(*_dg.PdfIndirectObject); _fcbg {
		_debgd._fecg = _dgga
	}
	_geagc = _dg.TraceToDirectObject(_geagc)
	_becbe, _beab := _geagc.(*_dg.PdfObjectArray)
	if !_beab {
		return nil, _b.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062je\u0063\u0074")
	}
	if _becbe.Len() != 4 {
		return nil, _b.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074i\u006f\u006e \u0043\u0053\u003a\u0020\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
	}
	_geagc = _becbe.Get(0)
	_bgcf, _beab := _geagc.(*_dg.PdfObjectName)
	if !_beab {
		return nil, _b.Errorf("\u0073\u0065\u0070ar\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a \u0069n\u0076a\u006ci\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_bgcf != "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e" {
		return nil, _b.Errorf("\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020w\u0072o\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	_geagc = _becbe.Get(1)
	_bgcf, _beab = _geagc.(*_dg.PdfObjectName)
	if !_beab {
		return nil, _b.Errorf("\u0073\u0065pa\u0072\u0061\u0074i\u006f\u006e\u0020\u0043S: \u0049nv\u0061\u006c\u0069\u0064\u0020\u0063\u006flo\u0072\u0061\u006e\u0074\u0020\u006e\u0061m\u0065")
	}
	_debgd.ColorantName = _bgcf
	_geagc = _becbe.Get(2)
	_afdb, _cebc := NewPdfColorspaceFromPdfObject(_geagc)
	if _cebc != nil {
		return nil, _cebc
	}
	_debgd.AlternateSpace = _afdb
	_ggdf, _cebc := _agec(_becbe.Get(3))
	if _cebc != nil {
		return nil, _cebc
	}
	_debgd.TintTransform = _ggdf
	return _debgd, nil
}

// NewPdfShadingType2 creates an empty shading type 2 dictionary.
func NewPdfShadingType2() *PdfShadingType2 {
	_acgeec := &PdfShadingType2{}
	_acgeec.PdfShading = &PdfShading{}
	_acgeec.PdfShading._bcfbg = _dg.MakeIndirectObject(_dg.MakeDict())
	_acgeec.PdfShading._eeddb = _acgeec
	return _acgeec
}
func (_bfggfb *LTV) getCerts(_gfgb []*_bb.Certificate) ([][]byte, error) {
	_fgefg := make([][]byte, 0, len(_gfgb))
	for _, _abdcb := range _gfgb {
		_fgefg = append(_fgefg, _abdcb.Raw)
	}
	return _fgefg, nil
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_begcae pdfFontType3) GetRuneMetrics(r rune) (_bbg.CharMetrics, bool) {
	_bgfcd := _begcae.Encoder()
	if _bgfcd == nil {
		_ag.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _begcae)
		return _bbg.CharMetrics{}, false
	}
	_befd, _bdca := _bgfcd.RuneToCharcode(r)
	if !_bdca {
		if r != ' ' {
			_ag.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _begcae)
		}
		return _bbg.CharMetrics{}, false
	}
	_geca, _aeeba := _begcae.GetCharMetrics(_befd)
	return _geca, _aeeba
}

// GetContainingPdfObject implements interface PdfModel.
func (_faabf *Permissions) GetContainingPdfObject() _dg.PdfObject { return _faabf._fbdaa }

// ToImage converts an object to an Image which can be transformed or saved out.
// The image data is decoded and the Image returned.
func (_ccfcd *XObjectImage) ToImage() (*Image, error) {
	_ebcfba := &Image{}
	if _ccfcd.Height == nil {
		return nil, _bf.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_ebcfba.Height = *_ccfcd.Height
	if _ccfcd.Width == nil {
		return nil, _bf.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_ebcfba.Width = *_ccfcd.Width
	if _ccfcd.BitsPerComponent == nil {
		switch _ccfcd.Filter.(type) {
		case *_dg.CCITTFaxEncoder, *_dg.JBIG2Encoder:
			_ebcfba.BitsPerComponent = 1
		case *_dg.LZWEncoder, *_dg.RunLengthEncoder:
			_ebcfba.BitsPerComponent = 8
		default:
			return nil, _bf.New("\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006fm\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
	} else {
		_ebcfba.BitsPerComponent = *_ccfcd.BitsPerComponent
	}
	_ebcfba.ColorComponents = _ccfcd.ColorSpace.GetNumComponents()
	_ccfcd._abfb.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _dg.MakeInteger(int64(_ebcfba.ColorComponents)))
	_cafgb, _bdad := _dg.DecodeStream(_ccfcd._abfb)
	if _bdad != nil {
		return nil, _bdad
	}
	_ebcfba.Data = _cafgb
	if _ccfcd.Decode != nil {
		_dbbbe, _ebcccf := _ccfcd.Decode.(*_dg.PdfObjectArray)
		if !_ebcccf {
			_ag.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020D\u0065\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074")
			return nil, _bf.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_edebd, _faedd := _dbbbe.ToFloat64Array()
		if _faedd != nil {
			return nil, _faedd
		}
		switch _ccfcd.ColorSpace.(type) {
		case *PdfColorspaceDeviceCMYK:
			_dcggb := _ccfcd.ColorSpace.DecodeArray()
			if _dcggb[0] == _edebd[0] && _dcggb[1] == _edebd[1] && _dcggb[2] == _edebd[2] && _dcggb[3] == _edebd[3] {
				_ebcfba._gfbb = _edebd
			}
		default:
			_ebcfba._gfbb = _edebd
		}
	}
	return _ebcfba, nil
}

// SetFlag sets the flag for the field.
func (_ddgcd *PdfField) SetFlag(flag FieldFlag) { _ddgcd.Ff = _dg.MakeInteger(int64(flag)) }
func (_daegege *PdfWriter) flushWriter() error {
	if _daegege._ffefc == nil {
		_daegege._ffefc = _daegege._bddfa.Flush()
	}
	return _daegege._ffefc
}

// NewPdfColorspaceSpecialIndexed returns a new Indexed color.
func NewPdfColorspaceSpecialIndexed() *PdfColorspaceSpecialIndexed {
	return &PdfColorspaceSpecialIndexed{HiVal: 255}
}

// ToOutlineTree returns a low level PdfOutlineTreeNode object, based on
// the current instance.
func (_gffca *Outline) ToOutlineTree() *PdfOutlineTreeNode {
	return &_gffca.ToPdfOutline().PdfOutlineTreeNode
}

// GetNumComponents returns the number of input color components, i.e. that are input to the tint transform.
func (_fbec *PdfColorspaceDeviceN) GetNumComponents() int { return _fbec.ColorantNames.Len() }

// PdfAnnotationMovie represents Movie annotations.
// (Section 12.5.6.17).
type PdfAnnotationMovie struct {
	*PdfAnnotation
	T     _dg.PdfObject
	Movie _dg.PdfObject
	A     _dg.PdfObject
}

// GetVersion gets the document version.
func (_cbdff *PdfWriter) GetVersion() _dg.Version { return _cbdff._efacd }

// ImageToRGB converts an image in CMYK32 colorspace to an RGB image.
func (_ecga *PdfColorspaceDeviceCMYK) ImageToRGB(img Image) (Image, error) {
	_ag.Log.Trace("\u0043\u004d\u0059\u004b\u0033\u0032\u0020\u002d\u003e\u0020\u0052\u0047\u0042")
	_ag.Log.Trace("I\u006d\u0061\u0067\u0065\u0020\u0042P\u0043\u003a\u0020\u0025\u0064\u002c \u0043\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020%\u0064", img.BitsPerComponent, img.ColorComponents)
	_ag.Log.Trace("\u004c\u0065\u006e \u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0064", len(img.Data))
	_ag.Log.Trace("H\u0065\u0069\u0067\u0068t:\u0020%\u0064\u002c\u0020\u0057\u0069d\u0074\u0068\u003a\u0020\u0025\u0064", img.Height, img.Width)
	_dddc, _dbgd := _fc.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._dgeb, img._gfbb)
	if _dbgd != nil {
		return Image{}, _dbgd
	}
	_gceag, _dbgd := _fc.NRGBAConverter.Convert(_dddc)
	if _dbgd != nil {
		return Image{}, _dbgd
	}
	return _edcf(_gceag.Base()), nil
}

// ToPdfObject implements interface PdfModel.
func (_ebgf *PdfAnnotationRedact) ToPdfObject() _dg.PdfObject {
	_ebgf.PdfAnnotation.ToPdfObject()
	_gdgfe := _ebgf._cdf
	_cbac := _gdgfe.PdfObject.(*_dg.PdfObjectDictionary)
	_ebgf.PdfAnnotationMarkup.appendToPdfDictionary(_cbac)
	_cbac.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u0052\u0065\u0064\u0061\u0063\u0074"))
	_cbac.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _ebgf.QuadPoints)
	_cbac.SetIfNotNil("\u0049\u0043", _ebgf.IC)
	_cbac.SetIfNotNil("\u0052\u004f", _ebgf.RO)
	_cbac.SetIfNotNil("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074", _ebgf.OverlayText)
	_cbac.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _ebgf.Repeat)
	_cbac.SetIfNotNil("\u0044\u0041", _ebgf.DA)
	_cbac.SetIfNotNil("\u0051", _ebgf.Q)
	return _gdgfe
}

// GetNumComponents returns the number of color components (1 for Indexed).
func (_afgb *PdfColorspaceSpecialIndexed) GetNumComponents() int { return 1 }
func _eadf(_eabage _dg.PdfObject, _facecc *PdfReader) (*OutlineDest, error) {
	_eacff, _cabdda := _dg.GetArray(_eabage)
	if !_cabdda {
		return nil, _bf.New("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0064\u0065\u0073t\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_gcgf := _eacff.Len()
	if _gcgf < 2 {
		return nil, _b.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u006fu\u0074\u006c\u0069n\u0065\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061ti\u006f\u006e\u0020a\u0072\u0072a\u0079\u0020\u006c\u0065\u006e\u0067t\u0068\u003a \u0025\u0064", _gcgf)
	}
	_bdda := &OutlineDest{Mode: "\u0046\u0069\u0074"}
	_bccda := _eacff.Get(0)
	if _dbcaf, _eccd := _dg.GetIndirect(_bccda); _eccd {
		if _, _dabfd, _bcaef := _facecc.PageFromIndirectObject(_dbcaf); _bcaef == nil {
			_bdda.Page = int64(_dabfd - 1)
		} else {
			_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020g\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u002b\u0076", _dbcaf)
		}
		_bdda.PageObj = _dbcaf
	} else if _eeeb, _dbabg := _dg.GetIntVal(_bccda); _dbabg {
		if _eeeb >= 0 && _eeeb < len(_facecc.PageList) {
			_bdda.PageObj = _facecc.PageList[_eeeb].GetPageAsIndirectObject()
		} else {
			_ag.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _eeeb)
		}
		_bdda.Page = int64(_eeeb)
	} else {
		return nil, _b.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006f\u0075\u0074\u006cine\u0020de\u0073\u0074\u0069\u006e\u0061\u0074\u0069on\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0054", _bccda)
	}
	_eaca, _cabdda := _dg.GetNameVal(_eacff.Get(1))
	if !_cabdda {
		_ag.Log.Debug("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006d\u0061\u0067\u006e\u0069\u0066\u0069\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _eacff.Get(1))
		return _bdda, nil
	}
	switch _eaca {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		if _gcgf > 2 {
			_bdda.Y, _ = _dg.GetNumberAsFloat(_dg.TraceToDirectObject(_eacff.Get(2)))
		}
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		if _gcgf > 2 {
			_bdda.X, _ = _dg.GetNumberAsFloat(_dg.TraceToDirectObject(_eacff.Get(2)))
		}
	case "\u0058\u0059\u005a":
		if _gcgf > 4 {
			_bdda.X, _ = _dg.GetNumberAsFloat(_dg.TraceToDirectObject(_eacff.Get(2)))
			_bdda.Y, _ = _dg.GetNumberAsFloat(_dg.TraceToDirectObject(_eacff.Get(3)))
			_bdda.Zoom, _ = _dg.GetNumberAsFloat(_dg.TraceToDirectObject(_eacff.Get(4)))
		}
	default:
		_eaca = "\u0046\u0069\u0074"
	}
	_bdda.Mode = _eaca
	return _bdda, nil
}

// PdfColorLab represents a color in the L*, a*, b* 3 component colorspace.
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorLab [3]float64

// ToInteger convert to an integer format.
func (_gagb *PdfColorDeviceGray) ToInteger(bits int) uint32 {
	_fgfb := _cg.Pow(2, float64(bits)) - 1
	return uint32(_fgfb * _gagb.Val())
}
func (_acedc *fontFile) parseASCIIPart(_dbaa []byte) error {
	if len(_dbaa) < 2 || string(_dbaa[:2]) != "\u0025\u0021" {
		return _bf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0074a\u0072\u0074\u0020\u006f\u0066\u0020\u0041S\u0043\u0049\u0049\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_daed, _adgcca, _cfdag := _gadec(_dbaa)
	if _cfdag != nil {
		return _cfdag
	}
	_aafad := _dbdg(_daed)
	_acedc._ecgd = _aafad["\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"]
	if _acedc._ecgd == "" {
		_ag.Log.Debug("\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0020\u0068a\u0073\u0020\u006e\u006f\u0020\u002f\u0046\u006f\u006e\u0074N\u0061\u006d\u0065")
	}
	if _adgcca != "" {
		_egcf, _ggdef := _gcga(_adgcca)
		if _ggdef != nil {
			return _ggdef
		}
		_bgecc, _ggdef := _bd.NewCustomSimpleTextEncoder(_egcf, nil)
		if _ggdef != nil {
			_ag.Log.Debug("\u0045\u0052\u0052\u004fR\u0020\u003a\u0055\u004e\u004b\u004e\u004f\u0057\u004e\u0020G\u004cY\u0050\u0048\u003a\u0020\u0065\u0072\u0072=\u0025\u0076", _ggdef)
			return nil
		}
		_acedc._beaeb = _bgecc
	}
	return nil
}
func (_accb *PdfReader) newPdfAnnotationUnderlineFromDict(_gbae *_dg.PdfObjectDictionary) (*PdfAnnotationUnderline, error) {
	_cccc := PdfAnnotationUnderline{}
	_gggec, _gaee := _accb.newPdfAnnotationMarkupFromDict(_gbae)
	if _gaee != nil {
		return nil, _gaee
	}
	_cccc.PdfAnnotationMarkup = _gggec
	_cccc.QuadPoints = _gbae.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_cccc, nil
}

// NewPdfActionTrans returns a new "trans" action.
func NewPdfActionTrans() *PdfActionTrans {
	_ecbb := NewPdfAction()
	_abc := &PdfActionTrans{}
	_abc.PdfAction = _ecbb
	_ecbb.SetContext(_abc)
	return _abc
}
func (_befg *Image) resampleLowBits(_dbfd []uint32) {
	_ddge := _fc.BytesPerLine(int(_befg.Width), int(_befg.BitsPerComponent), _befg.ColorComponents)
	_gaacbf := make([]byte, _befg.ColorComponents*_ddge*int(_befg.Height))
	_fcaeg := int(_befg.BitsPerComponent) * _befg.ColorComponents * int(_befg.Width)
	_gabceb := uint8(8)
	var (
		_gbacf, _fgbee int
		_eagbgf        uint32
	)
	for _ccfbf := 0; _ccfbf < int(_befg.Height); _ccfbf++ {
		_fgbee = _ccfbf * _ddge
		for _dfgg := 0; _dfgg < _fcaeg; _dfgg++ {
			_eagbgf = _dbfd[_gbacf]
			_gabceb -= uint8(_befg.BitsPerComponent)
			_gaacbf[_fgbee] |= byte(_eagbgf) << _gabceb
			if _gabceb == 0 {
				_gabceb = 8
				_fgbee++
			}
			_gbacf++
		}
	}
	_befg.Data = _gaacbf
}
func (_bdcbb *LTV) getCRLs(_ccee []*_bb.Certificate) ([][]byte, error) {
	_bacdc := make([][]byte, 0, len(_ccee))
	for _, _bbfb := range _ccee {
		for _, _adfeb := range _bbfb.CRLDistributionPoints {
			if _bdcbb.CertClient.IsCA(_bbfb) {
				continue
			}
			_fedc, _aaafg := _bdcbb.CRLClient.MakeRequest(_adfeb, _bbfb)
			if _aaafg != nil {
				_ag.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _aaafg)
				continue
			}
			_bacdc = append(_bacdc, _fedc)
		}
	}
	return _bacdc, nil
}

// String returns a string representation of the field.
func (_afgg *PdfField) String() string {
	if _befc, _fadeg := _afgg.ToPdfObject().(*_dg.PdfIndirectObject); _fadeg {
		return _b.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _afgg._bdfg, _befc.PdfObject.String())
	}
	return ""
}
func (_efaaf *LTV) validateSig(_facbb *PdfSignature) error {
	if _facbb == nil || _facbb.Contents == nil || len(_facbb.Contents.Bytes()) == 0 {
		return _b.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0066\u0069\u0065l\u0064:\u0020\u0025\u0076", _facbb)
	}
	return nil
}

// NewImageFromGoImage creates a new NRGBA32 unidoc Image from a golang Image.
// If `goimg` is grayscale (*goimage.Gray8) then calls NewGrayImageFromGoImage instead.
func (_bgfeg DefaultImageHandler) NewImageFromGoImage(goimg _gd.Image) (*Image, error) {
	_babd, _dgcc := _fc.FromGoImage(goimg)
	if _dgcc != nil {
		return nil, _dgcc
	}
	_aafd := _edcf(_babd.Base())
	return &_aafd, nil
}
func (_gbdgb *PdfWriter) writeTrailer(_efeffc int) {
	_gbdgb.writeString("\u0078\u0072\u0065\u0066\u000d\u000a")
	for _bbbbbf := 0; _bbbbbf <= _efeffc; {
		for ; _bbbbbf <= _efeffc; _bbbbbf++ {
			_ccgf, _fbfbbe := _gbdgb._fffge[_bbbbbf]
			if _fbfbbe && (!_gbdgb._bbac || _gbdgb._bbac && (_ccgf.Type == 1 && _ccgf.Offset >= _gbdgb._fafab || _ccgf.Type == 0)) {
				break
			}
		}
		var _gfadf int
		for _gfadf = _bbbbbf + 1; _gfadf <= _efeffc; _gfadf++ {
			_dcdda, _edff := _gbdgb._fffge[_gfadf]
			if _edff && (!_gbdgb._bbac || _gbdgb._bbac && (_dcdda.Type == 1 && _dcdda.Offset > _gbdgb._fafab)) {
				continue
			}
			break
		}
		_gcae := _b.Sprintf("\u0025d\u0020\u0025\u0064\u000d\u000a", _bbbbbf, _gfadf-_bbbbbf)
		_gbdgb.writeString(_gcae)
		for _dccae := _bbbbbf; _dccae < _gfadf; _dccae++ {
			_fcceeb := _gbdgb._fffge[_dccae]
			switch _fcceeb.Type {
			case 0:
				_gcae = _b.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u0066\u000d\u000a", 0, 65535)
				_gbdgb.writeString(_gcae)
			case 1:
				_gcae = _b.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u006e\u000d\u000a", _fcceeb.Offset, 0)
				_gbdgb.writeString(_gcae)
			}
		}
		_bbbbbf = _gfadf + 1
	}
	_fdebc := _dg.MakeDict()
	_fdebc.Set("\u0049\u006e\u0066\u006f", _gbdgb._efbfa)
	_fdebc.Set("\u0052\u006f\u006f\u0074", _gbdgb._fadee)
	_fdebc.Set("\u0053\u0069\u007a\u0065", _dg.MakeInteger(int64(_efeffc+1)))
	if _gbdgb._bbac && _gbdgb._eege > 0 {
		_fdebc.Set("\u0050\u0072\u0065\u0076", _dg.MakeInteger(_gbdgb._eege))
	}
	if _gbdgb._fadcg != nil {
		_fdebc.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _gbdgb._acdag)
	}
	if _gbdgb._agbeg == nil && _gbdgb._cedaf != "" && _gbdgb._ceaab != "" {
		_gbdgb._agbeg = _dg.MakeArray(_dg.MakeHexString(_gbdgb._cedaf), _dg.MakeHexString(_gbdgb._ceaab))
	}
	if _gbdgb._agbeg != nil {
		_fdebc.Set("\u0049\u0044", _gbdgb._agbeg)
		_ag.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _gbdgb._agbeg)
	}
	_gbdgb.writeString("\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u000a")
	_gbdgb.writeString(_fdebc.WriteString())
	_gbdgb.writeString("\u000a")
}

// NewPdfFontFromPdfObject loads a PdfFont from the dictionary `fontObj`.  If there is a problem an
// error is returned.
func NewPdfFontFromPdfObject(fontObj _dg.PdfObject) (*PdfFont, error) { return _gegbd(fontObj, true) }

// NewPdfColorDeviceGray returns a new grayscale color based on an input grayscale float value in range [0-1].
func NewPdfColorDeviceGray(grayVal float64) *PdfColorDeviceGray {
	_cebaa := PdfColorDeviceGray(grayVal)
	return &_cebaa
}

// K returns the value of the key component of the color.
func (_eadc *PdfColorDeviceCMYK) K() float64 { return _eadc[3] }

// ToPdfObject returns the PDF representation of the function.
func (_gbabe *PdfFunctionType4) ToPdfObject() _dg.PdfObject {
	_daaf := _gbabe._ggfdf
	if _daaf == nil {
		_gbabe._ggfdf = &_dg.PdfObjectStream{}
		_daaf = _gbabe._ggfdf
	}
	_aabbb := _dg.MakeDict()
	_aabbb.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _dg.MakeInteger(4))
	_dbad := &_dg.PdfObjectArray{}
	for _, _dbgcbg := range _gbabe.Domain {
		_dbad.Append(_dg.MakeFloat(_dbgcbg))
	}
	_aabbb.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _dbad)
	_ffaa := &_dg.PdfObjectArray{}
	for _, _dcagb := range _gbabe.Range {
		_ffaa.Append(_dg.MakeFloat(_dcagb))
	}
	_aabbb.Set("\u0052\u0061\u006eg\u0065", _ffaa)
	if _gbabe._cbdbd == nil && _gbabe.Program != nil {
		_gbabe._cbdbd = []byte(_gbabe.Program.String())
	}
	_aabbb.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _dg.MakeInteger(int64(len(_gbabe._cbdbd))))
	_daaf.Stream = _gbabe._cbdbd
	_daaf.PdfObjectDictionary = _aabbb
	return _daaf
}

// XObjectImage (Table 89 in 8.9.5.1).
// Implements PdfModel interface.
type XObjectImage struct {

	// ColorSpace       PdfObject
	Width            *int64
	Height           *int64
	ColorSpace       PdfColorspace
	BitsPerComponent *int64
	Filter           _dg.StreamEncoder
	Intent           _dg.PdfObject
	ImageMask        _dg.PdfObject
	Mask             _dg.PdfObject
	Matte            _dg.PdfObject
	Decode           _dg.PdfObject
	Interpolate      _dg.PdfObject
	Alternatives     _dg.PdfObject
	SMask            _dg.PdfObject
	SMaskInData      _dg.PdfObject
	Name             _dg.PdfObject
	StructParent     _dg.PdfObject
	ID               _dg.PdfObject
	OPI              _dg.PdfObject
	Metadata         _dg.PdfObject
	OC               _dg.PdfObject
	Stream           []byte
	_abfb            *_dg.PdfObjectStream
}

// ToUnicode returns the name of the font's "ToUnicode" field if there is one, or "" if there isn't.
func (_bfac *PdfFont) ToUnicode() string {
	if _bfac.baseFields()._ecfb == nil {
		return ""
	}
	return _bfac.baseFields()._ecfb.Name()
}

// ToPdfObject implements interface PdfModel.
func (_ggd *PdfAnnotationLine) ToPdfObject() _dg.PdfObject {
	_ggd.PdfAnnotation.ToPdfObject()
	_aadc := _ggd._cdf
	_fafb := _aadc.PdfObject.(*_dg.PdfObjectDictionary)
	_ggd.PdfAnnotationMarkup.appendToPdfDictionary(_fafb)
	_fafb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _dg.MakeName("\u004c\u0069\u006e\u0065"))
	_fafb.SetIfNotNil("\u004c", _ggd.L)
	_fafb.SetIfNotNil("\u0042\u0053", _ggd.BS)
	_fafb.SetIfNotNil("\u004c\u0045", _ggd.LE)
	_fafb.SetIfNotNil("\u0049\u0043", _ggd.IC)
	_fafb.SetIfNotNil("\u004c\u004c", _ggd.LL)
	_fafb.SetIfNotNil("\u004c\u004c\u0045", _ggd.LLE)
	_fafb.SetIfNotNil("\u0043\u0061\u0070", _ggd.Cap)
	_fafb.SetIfNotNil("\u0049\u0054", _ggd.IT)
	_fafb.SetIfNotNil("\u004c\u004c\u004f", _ggd.LLO)
	_fafb.SetIfNotNil("\u0043\u0050", _ggd.CP)
	_fafb.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _ggd.Measure)
	_fafb.SetIfNotNil("\u0043\u004f", _ggd.CO)
	return _aadc
}
func _bgcbab() _a.Time { _fgefgf.Lock(); defer _fgefgf.Unlock(); return _ccfea }

// GetNumComponents returns the number of color components (3 for RGB).
func (_fbge *PdfColorDeviceRGB) GetNumComponents() int { return 3 }

// AddWatermarkImage adds a watermark to the page.
func (_agad *PdfPage) AddWatermarkImage(ximg *XObjectImage, opt WatermarkImageOptions) error {
	_edaca, _aaef := _agad.GetMediaBox()
	if _aaef != nil {
		return _aaef
	}
	_ddfba := _edaca.Urx - _edaca.Llx
	_fbad := _edaca.Ury - _edaca.Lly
	_bfabcb := float64(*ximg.Width)
	_egbae := (_ddfba - _bfabcb) / 2
	if opt.FitToWidth {
		_bfabcb = _ddfba
		_egbae = 0
	}
	_cbdd := _fbad
	_ecacc := float64(0)
	if opt.PreserveAspectRatio {
		_cbdd = _bfabcb * float64(*ximg.Height) / float64(*ximg.Width)
		_ecacc = (_fbad - _cbdd) / 2
	}
	if _agad.Resources == nil {
		_agad.Resources = NewPdfPageResources()
	}
	_bcdga := 0
	_cbbgg := _dg.PdfObjectName(_b.Sprintf("\u0049\u006d\u0077%\u0064", _bcdga))
	for _agad.Resources.HasXObjectByName(_cbbgg) {
		_bcdga++
		_cbbgg = _dg.PdfObjectName(_b.Sprintf("\u0049\u006d\u0077%\u0064", _bcdga))
	}
	_aaef = _agad.AddImageResource(_cbbgg, ximg)
	if _aaef != nil {
		return _aaef
	}
	_bcdga = 0
	_aefgff := _dg.PdfObjectName(_b.Sprintf("\u0047\u0053\u0025\u0064", _bcdga))
	for _agad.HasExtGState(_aefgff) {
		_bcdga++
		_aefgff = _dg.PdfObjectName(_b.Sprintf("\u0047\u0053\u0025\u0064", _bcdga))
	}
	_dadaaf := _dg.MakeDict()
	_dadaaf.Set("\u0042\u004d", _dg.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	_dadaaf.Set("\u0043\u0041", _dg.MakeFloat(opt.Alpha))
	_dadaaf.Set("\u0063\u0061", _dg.MakeFloat(opt.Alpha))
	_aaef = _agad.AddExtGState(_aefgff, _dadaaf)
	if _aaef != nil {
		return _aaef
	}
	_bbcgd := _b.Sprintf("\u0071\u000a"+"\u002f%\u0073\u0020\u0067\u0073\u000a"+"%\u002e\u0030\u0066\u0020\u0030\u00200\u0020\u0025\u002e\u0030\u0066\u0020\u0025\u002e\u0034f\u0020\u0025\u002e4\u0066 \u0063\u006d\u000a"+"\u002f%\u0073\u0020\u0044\u006f\u000a"+"\u0051", _aefgff, _bfabcb, _cbdd, _egbae, _ecacc, _cbbgg)
	_agad.AddContentStreamByString(_bbcgd)
	return nil
}

// NewPdfAnnotationSquare returns a new square annotation.
func NewPdfAnnotationSquare() *PdfAnnotationSquare {
	_fgb := NewPdfAnnotation()
	_ddb := &PdfAnnotationSquare{}
	_ddb.PdfAnnotation = _fgb
	_ddb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fgb.SetContext(_ddb)
	return _ddb
}

// AddCustomInfo adds a custom info into document info dictionary.
func (_ggfa *PdfInfo) AddCustomInfo(name string, value string) error {
	if _ggfa._dccdg == nil {
		_ggfa._dccdg = _dg.MakeDict()
	}
	if _, _bdbc := _ddcf[name]; _bdbc {
		return _b.Errorf("\u0063\u0061\u006e\u006e\u006ft\u0020\u0075\u0073\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072\u0064 \u0069\u006e\u0066\u006f\u0020\u006b\u0065\u0079\u0020\u0025\u0073\u0020\u0061\u0073\u0020\u0063\u0075\u0073\u0074\u006f\u006d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u006b\u0065y", name)
	}
	_ggfa._dccdg.SetIfNotNil(*_dg.MakeName(name), _dg.MakeString(value))
	return nil
}

// NewPdfOutlineItem returns an initialized PdfOutlineItem.
func NewPdfOutlineItem() *PdfOutlineItem {
	_geebb := &PdfOutlineItem{_fbgf: _dg.MakeIndirectObject(_dg.MakeDict())}
	_geebb._baddf = _geebb
	return _geebb
}

// GetAsShadingPattern returns a shading pattern. Check with IsShading() prior to using this.
func (_bfaff *PdfPattern) GetAsShadingPattern() *PdfShadingPattern {
	return _bfaff._cgdcc.(*PdfShadingPattern)
}

// PdfAnnotation represents an annotation in PDF (section 12.5 p. 389).
type PdfAnnotation struct {
	_egcg        PdfModel
	Rect         _dg.PdfObject
	Contents     _dg.PdfObject
	P            _dg.PdfObject
	NM           _dg.PdfObject
	M            _dg.PdfObject
	F            _dg.PdfObject
	AP           _dg.PdfObject
	AS           _dg.PdfObject
	Border       _dg.PdfObject
	C            _dg.PdfObject
	StructParent _dg.PdfObject
	OC           _dg.PdfObject
	_cdf         *_dg.PdfIndirectObject
}

// AddContentStreamByString adds content stream by string. Puts the content
// string into a stream object and points the content stream towards it.
func (_abcf *PdfPage) AddContentStreamByString(contentStr string) error {
	_egafg, _cfgge := _dg.MakeStream([]byte(contentStr), _dg.NewFlateEncoder())
	if _cfgge != nil {
		return _cfgge
	}
	if _abcf.Contents == nil {
		_abcf.Contents = _egafg
	} else {
		_cabf := _dg.TraceToDirectObject(_abcf.Contents)
		_bgceg, _fffdgf := _cabf.(*_dg.PdfObjectArray)
		if !_fffdgf {
			_bgceg = _dg.MakeArray(_cabf)
		}
		_bgceg.Append(_egafg)
		_abcf.Contents = _bgceg
	}
	return nil
}

// NewXObjectImageFromImage creates a new XObject Image from an image object
// with default options. If encoder is nil, uses raw encoding (none).
func NewXObjectImageFromImage(img *Image, cs PdfColorspace, encoder _dg.StreamEncoder) (*XObjectImage, error) {
	_gaegd := NewXObjectImage()
	return UpdateXObjectImageFromImage(_gaegd, img, cs, encoder)
}

// XObjectForm (Table 95 in 8.10.2).
type XObjectForm struct {
	Filter        _dg.StreamEncoder
	FormType      _dg.PdfObject
	BBox          _dg.PdfObject
	Matrix        _dg.PdfObject
	Resources     *PdfPageResources
	Group         _dg.PdfObject
	Ref           _dg.PdfObject
	MetaData      _dg.PdfObject
	PieceInfo     _dg.PdfObject
	LastModified  _dg.PdfObject
	StructParent  _dg.PdfObject
	StructParents _dg.PdfObject
	OPI           _dg.PdfObject
	OC            _dg.PdfObject
	Name          _dg.PdfObject

	// Stream data.
	Stream []byte
	_ebaeb *_dg.PdfObjectStream
}

// CompliancePdfReader is a wrapper over PdfReader that is used for verifying if the input Pdf document matches the
// compliance rules of standards like PDF/A.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type CompliancePdfReader struct {
	*PdfReader
	_acgf _dg.ParserMetadata
}

// GetStructRoot gets the StructTreeRoot object
func (_geggg *PdfPage) GetStructTreeRoot() (*_dg.PdfObject, bool) {
	_ebfg, _bgbcba := _geggg._cbbcc.GetCatalogStructTreeRoot()
	return &_ebfg, _bgbcba
}
func (_gecaf *Image) getSuitableEncoder() (_dg.StreamEncoder, error) {
	var (
		_caed, _dgdd = int(_gecaf.Width), int(_gecaf.Height)
		_gabce       = make(map[string]bool)
		_dfcbc       = true
		_bdeaa       = false
		_gbcbg       = func() *_dg.DCTEncoder { return _dg.NewDCTEncoder() }
		_fgcbd       = func() *_dg.DCTEncoder { _ddebf := _dg.NewDCTEncoder(); _ddebf.BitsPerComponent = 16; return _ddebf }
	)
	for _fbfad := 0; _fbfad < _dgdd; _fbfad++ {
		for _eaea := 0; _eaea < _caed; _eaea++ {
			_bfgac, _agfdc := _gecaf.ColorAt(_eaea, _fbfad)
			if _agfdc != nil {
				return nil, _agfdc
			}
			_dbbdc, _egabc, _eaead, _fcbdb := _bfgac.RGBA()
			if _dfcbc && (_dbbdc != _egabc || _dbbdc != _eaead || _egabc != _eaead) {
				_dfcbc = false
			}
			if !_bdeaa {
				switch _bfgac.(type) {
				case _edg.NRGBA:
					_bdeaa = _fcbdb > 0
				}
			}
			_gabce[_b.Sprintf("\u0025\u0064\u002c\u0025\u0064\u002c\u0025\u0064", _dbbdc, _egabc, _eaead)] = true
			if len(_gabce) > 2 && _bdeaa {
				return _fgcbd(), nil
			}
		}
	}
	if _bdeaa || len(_gecaf._dgeb) > 0 {
		return _dg.NewFlateEncoder(), nil
	}
	if len(_gabce) <= 2 {
		_cedeg := _gecaf.ConvertToBinary()
		if _cedeg != nil {
			return nil, _cedeg
		}
		return _dg.NewJBIG2Encoder(), nil
	}
	if _dfcbc {
		return _gbcbg(), nil
	}
	if _gecaf.ColorComponents == 1 {
		if _gecaf.BitsPerComponent == 1 {
			return _dg.NewJBIG2Encoder(), nil
		} else if _gecaf.BitsPerComponent == 8 {
			_aacca := _dg.NewDCTEncoder()
			_aacca.ColorComponents = 1
			return _aacca, nil
		}
	} else if _gecaf.ColorComponents == 3 {
		if _gecaf.BitsPerComponent == 8 {
			return _gbcbg(), nil
		} else if _gecaf.BitsPerComponent == 16 {
			return _fgcbd(), nil
		}
	} else if _gecaf.ColorComponents == 4 {
		_degdec := _fgcbd()
		_degdec.ColorComponents = 4
		return _degdec, nil
	}
	return _fgcbd(), nil
}
func (_bgbcg *Image) samplesTrimPadding(_dgfdd []uint32) []uint32 {
	_aade := _bgbcg.ColorComponents * int(_bgbcg.Width) * int(_bgbcg.Height)
	if len(_dgfdd) == _aade {
		return _dgfdd
	}
	_adcde := make([]uint32, _aade)
	_cfdc := int(_bgbcg.Width) * _bgbcg.ColorComponents
	var _fegeg, _cgdef, _gcfeg, _dgbf int
	_affde := _fc.BytesPerLine(int(_bgbcg.Width), int(_bgbcg.BitsPerComponent), _bgbcg.ColorComponents)
	for _fegeg = 0; _fegeg < int(_bgbcg.Height); _fegeg++ {
		_cgdef = _fegeg * int(_bgbcg.Width)
		_gcfeg = _fegeg * _affde
		for _dgbf = 0; _dgbf < _cfdc; _dgbf++ {
			_adcde[_cgdef+_dgbf] = _dgfdd[_gcfeg+_dgbf]
		}
	}
	return _adcde
}

// Initialize initializes the PdfSignature.
func (_deccd *PdfSignature) Initialize() error {
	if _deccd.Handler == nil {
		return _bf.New("\u0073\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0068\u0061\u006e\u0064\u006c\u0065r\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	return _deccd.Handler.InitSignature(_deccd)
}

// PdfFunction interface represents the common methods of a function in PDF.
type PdfFunction interface {
	Evaluate([]float64) ([]float64, error)
	ToPdfObject() _dg.PdfObject
}

// PageCallback callback function used in page loading
// that could be used to modify the page content.
//
// Deprecated: will be removed in v4. Use PageProcessCallback instead.
type PageCallback func(_gebab int, _ceffc *PdfPage)

// HasFontByName checks whether a font is defined by the specified keyName.
func (_cdefgc *PdfPageResources) HasFontByName(keyName _dg.PdfObjectName) bool {
	_, _bggce := _cdefgc.GetFontByName(keyName)
	return _bggce
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_dagb *PdfColorspaceCalRGB) ToPdfObject() _dg.PdfObject {
	_ggcgb := &_dg.PdfObjectArray{}
	_ggcgb.Append(_dg.MakeName("\u0043\u0061\u006c\u0052\u0047\u0042"))
	_gbceb := _dg.MakeDict()
	if _dagb.WhitePoint != nil {
		_bgfe := _dg.MakeArray(_dg.MakeFloat(_dagb.WhitePoint[0]), _dg.MakeFloat(_dagb.WhitePoint[1]), _dg.MakeFloat(_dagb.WhitePoint[2]))
		_gbceb.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _bgfe)
	} else {
		_ag.Log.Error("\u0043\u0061l\u0052\u0047\u0042\u003a \u004d\u0069s\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _dagb.BlackPoint != nil {
		_ddcba := _dg.MakeArray(_dg.MakeFloat(_dagb.BlackPoint[0]), _dg.MakeFloat(_dagb.BlackPoint[1]), _dg.MakeFloat(_dagb.BlackPoint[2]))
		_gbceb.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _ddcba)
	}
	if _dagb.Gamma != nil {
		_dgfdf := _dg.MakeArray(_dg.MakeFloat(_dagb.Gamma[0]), _dg.MakeFloat(_dagb.Gamma[1]), _dg.MakeFloat(_dagb.Gamma[2]))
		_gbceb.Set("\u0047\u0061\u006dm\u0061", _dgfdf)
	}
	if _dagb.Matrix != nil {
		_caab := _dg.MakeArray(_dg.MakeFloat(_dagb.Matrix[0]), _dg.MakeFloat(_dagb.Matrix[1]), _dg.MakeFloat(_dagb.Matrix[2]), _dg.MakeFloat(_dagb.Matrix[3]), _dg.MakeFloat(_dagb.Matrix[4]), _dg.MakeFloat(_dagb.Matrix[5]), _dg.MakeFloat(_dagb.Matrix[6]), _dg.MakeFloat(_dagb.Matrix[7]), _dg.MakeFloat(_dagb.Matrix[8]))
		_gbceb.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _caab)
	}
	_ggcgb.Append(_gbceb)
	if _dagb._ccae != nil {
		_dagb._ccae.PdfObject = _ggcgb
		return _dagb._ccae
	}
	return _ggcgb
}

// ToPdfObject returns a PDF object representation of the outline item.
func (_gbcbc *OutlineItem) ToPdfObject() _dg.PdfObject {
	_gdcbb, _ := _gbcbc.ToPdfOutlineItem()
	return _gdcbb.ToPdfObject()
}

// PdfAnnotationScreen represents Screen annotations.
// (Section 12.5.6.18).
type PdfAnnotationScreen struct {
	*PdfAnnotation
	T  _dg.PdfObject
	MK _dg.PdfObject
	A  _dg.PdfObject
	AA _dg.PdfObject
}

// NewPdfAnnotationSquiggly returns a new text squiggly annotation.
func NewPdfAnnotationSquiggly() *PdfAnnotationSquiggly {
	_gfa := NewPdfAnnotation()
	_fab := &PdfAnnotationSquiggly{}
	_fab.PdfAnnotation = _gfa
	_fab.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gfa.SetContext(_fab)
	return _fab
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a CalGray device.
func (_eaeb *PdfColorspaceCalGray) GetNumComponents() int { return 1 }

// ToPdfObject converts the PdfPage to a dictionary within an indirect object container.
func (_gcged *PdfPage) ToPdfObject() _dg.PdfObject {
	_gggde := _gcged._cggbe
	_gcged.GetPageDict()
	return _gggde
}

// GetIndirectObjectByNumber retrieves and returns a specific PdfObject by object number.
func (_ffffb *PdfReader) GetIndirectObjectByNumber(number int) (_dg.PdfObject, error) {
	_agce, _dcgfc := _ffffb._baad.LookupByNumber(number)
	return _agce, _dcgfc
}

type crossReference struct {
	Type int

	// Type 1
	Offset     int64
	Generation int64

	// Type 2
	ObjectNumber int
	Index        int
}

// SetPdfAuthor sets the Author attribute of the output PDF.
func SetPdfAuthor(author string) { _fgefgf.Lock(); defer _fgefgf.Unlock(); _bdeaf = author }

// IsShading specifies if the pattern is a shading pattern.
func (_aceae *PdfPattern) IsShading() bool { return _aceae.PatternType == 2 }

// PdfInfoTrapped specifies pdf trapped information.
type PdfInfoTrapped string

// Add appends an outline item as a child of the current outline item.
func (_ebcac *OutlineItem) Add(item *OutlineItem) { _ebcac.Entries = append(_ebcac.Entries, item) }

// ToInteger convert to an integer format.
func (_dabf *PdfColorLab) ToInteger(bits int) [3]uint32 {
	_adgc := _cg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_adgc * _dabf.L()), uint32(_adgc * _dabf.A()), uint32(_adgc * _dabf.B())}
}

type pdfFont interface {
	_bbg.Font

	// ToPdfObject returns a PDF representation of the font and implements interface Model.
	ToPdfObject() _dg.PdfObject
	getFontDescriptor() *PdfFontDescriptor
	baseFields() *fontCommon
}

// PdfAcroForm represents the AcroForm dictionary used for representation of form data in PDF.
type PdfAcroForm struct {
	Fields          *[]*PdfField
	NeedAppearances *_dg.PdfObjectBool
	SigFlags        *_dg.PdfObjectInteger
	CO              *_dg.PdfObjectArray
	DR              *PdfPageResources
	DA              *_dg.PdfObjectString
	Q               *_dg.PdfObjectInteger
	XFA             _dg.PdfObject

	// ADBEEchoSign extra objects from Adobe Acrobat, causing signature invalid if not exists.
	ADBEEchoSign _dg.PdfObject
	_bebfe       *_dg.PdfIndirectObject
	_eeecb       bool
}

// PdfColorDeviceCMYK is a CMYK32 color, where each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceCMYK [4]float64

// L returns the value of the L component of the color.
func (_cffc *PdfColorLab) L() float64 { return _cffc[0] }

// NewPdfPageResources returns a new PdfPageResources object.
func NewPdfPageResources() *PdfPageResources {
	_dead := &PdfPageResources{}
	_dead._ddcfb = _dg.MakeDict()
	return _dead
}

// PdfAnnotationFreeText represents FreeText annotations.
// (Section 12.5.6.6).
type PdfAnnotationFreeText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	DA _dg.PdfObject
	Q  _dg.PdfObject
	RC _dg.PdfObject
	DS _dg.PdfObject
	CL _dg.PdfObject
	IT _dg.PdfObject
	BE _dg.PdfObject
	RD _dg.PdfObject
	BS _dg.PdfObject
	LE _dg.PdfObject
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// L (range 0-100), A (range -100-100) and B (range -100-100) components of
// the color.
func (_ggfcf *PdfColorspaceLab) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gbegf := vals[0]
	if _gbegf < 0.0 || _gbegf > 100.0 {
		_ag.Log.Debug("\u004c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030-\u0031\u0030\u0030\u0029", _gbegf)
		return nil, ErrColorOutOfRange
	}
	_egdf := vals[1]
	_bdeb := float64(-100)
	_bffeg := float64(100)
	if len(_ggfcf.Range) > 1 {
		_bdeb = _ggfcf.Range[0]
		_bffeg = _ggfcf.Range[1]
	}
	if _egdf < _bdeb || _egdf > _bffeg {
		_ag.Log.Debug("\u0041\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _egdf, _bdeb, _bffeg)
		return nil, ErrColorOutOfRange
	}
	_fgaa := vals[2]
	_edgg := float64(-100)
	_cgfg := float64(100)
	if len(_ggfcf.Range) > 3 {
		_edgg = _ggfcf.Range[2]
		_cgfg = _ggfcf.Range[3]
	}
	if _fgaa < _edgg || _fgaa > _cgfg {
		_ag.Log.Debug("\u0062\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _fgaa, _edgg, _cgfg)
		return nil, ErrColorOutOfRange
	}
	_bfde := NewPdfColorLab(_gbegf, _egdf, _fgaa)
	return _bfde, nil
}

// PdfFunctionType4 is a Postscript calculator functions.
type PdfFunctionType4 struct {
	Domain  []float64
	Range   []float64
	Program *_gfd.PSProgram
	_agbgd  *_gfd.PSExecutor
	_cbdbd  []byte
	_ggfdf  *_dg.PdfObjectStream
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_afecg *PdfFunctionType3) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_ag.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _bf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return nil, _bf.New("\u006e\u006f\u0074\u0020im\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
}
func (_ggde *DSS) generateHashMap(_dceaf []*_dg.PdfObjectStream) (map[string]*_dg.PdfObjectStream, error) {
	_feca := map[string]*_dg.PdfObjectStream{}
	for _, _geecd := range _dceaf {
		_bcgcf, _face := _dg.DecodeStream(_geecd)
		if _face != nil {
			return nil, _face
		}
		_cbdc, _face := _cbeaf(_bcgcf)
		if _face != nil {
			return nil, _face
		}
		_feca[string(_cbdc)] = _geecd
	}
	return _feca, nil
}

// WatermarkImageOptions contains options for configuring the watermark process.
type WatermarkImageOptions struct {
	Alpha               float64
	FitToWidth          bool
	PreserveAspectRatio bool
}
