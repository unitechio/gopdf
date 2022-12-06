package model

import (
	_b "bufio"
	_ede "bytes"
	_bf "crypto/md5"
	_d "crypto/rand"
	_a "crypto/sha1"
	_bg "crypto/x509"
	_ba "encoding/binary"
	_ed "encoding/hex"
	_ceg "errors"
	_ee "fmt"
	_g "hash"
	_gf "image"
	_be "image/color"
	_ "image/gif"
	_ "image/png"
	_f "io"
	_bb "io/ioutil"
	_ced "math"
	_edc "math/rand"
	_db "os"
	_cc "regexp"
	_dd "sort"
	_gb "strconv"
	_dac "strings"
	_c "sync"
	_ce "time"
	_da "unicode"
	_ca "unicode/utf8"

	_ad "bitbucket.org/shenghui0779/gopdf/common"
	_cde "bitbucket.org/shenghui0779/gopdf/core"
	_ccg "bitbucket.org/shenghui0779/gopdf/core/security"
	_cd "bitbucket.org/shenghui0779/gopdf/core/security/crypt"
	_fb "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_ff "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_cae "bitbucket.org/shenghui0779/gopdf/internal/sampling"
	_gc "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_cce "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_adb "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ab "bitbucket.org/shenghui0779/gopdf/model/internal/docutil"
	_fe "bitbucket.org/shenghui0779/gopdf/model/internal/fonts"
	_cdc "bitbucket.org/shenghui0779/gopdf/model/mdp"
	_bbe "bitbucket.org/shenghui0779/gopdf/model/sigutil"
	_dg "bitbucket.org/shenghui0779/gopdf/ps"
	_cb "github.com/unidoc/pkcs7"
	_ef "github.com/unidoc/unitype"
	_beb "golang.org/x/xerrors"
)

func (_egac *PdfReader) newPdfActionSubmitFormFromDict(_cbf *_cde.PdfObjectDictionary) (*PdfActionSubmitForm, error) {
	_gagc, _gee := _beed(_cbf.Get("\u0046"))
	if _gee != nil {
		return nil, _gee
	}
	return &PdfActionSubmitForm{F: _gagc, Fields: _cbf.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _cbf.Get("\u0046\u006c\u0061g\u0073")}, nil
}

// PdfActionSubmitForm represents a submitForm action.
type PdfActionSubmitForm struct {
	*PdfAction
	F      *PdfFilespec
	Fields _cde.PdfObject
	Flags  _cde.PdfObject
}

// ToPdfObject converts the pdfFontSimple to its PDF representation for outputting.
func (_gbcef *pdfFontSimple) ToPdfObject() _cde.PdfObject {
	if _gbcef._acebaa == nil {
		_gbcef._acebaa = &_cde.PdfIndirectObject{}
	}
	_aafc := _gbcef.baseFields().asPdfObjectDictionary("")
	_gbcef._acebaa.PdfObject = _aafc
	if _gbcef.FirstChar != nil {
		_aafc.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _gbcef.FirstChar)
	}
	if _gbcef.LastChar != nil {
		_aafc.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _gbcef.LastChar)
	}
	if _gbcef.Widths != nil {
		_aafc.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _gbcef.Widths)
	}
	if _gbcef.Encoding != nil {
		_aafc.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _gbcef.Encoding)
	} else if _gbcef._efeaf != nil {
		_eagga := _gbcef._efeaf.ToPdfObject()
		if _eagga != nil {
			_aafc.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _eagga)
		}
	}
	return _gbcef._acebaa
}

// PdfAnnotation3D represents 3D annotations.
// (Section 13.6.2).
type PdfAnnotation3D struct {
	*PdfAnnotation
	T3DD _cde.PdfObject
	T3DV _cde.PdfObject
	T3DA _cde.PdfObject
	T3DI _cde.PdfObject
	T3DB _cde.PdfObject
}

func _egadg(_acbea *PdfAnnotation) (*XObjectForm, *PdfRectangle, error) {
	_dagc, _ddec := _cde.GetDict(_acbea.AP)
	if !_ddec {
		return nil, nil, _ceg.New("f\u0069\u0065\u006c\u0064\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0041\u0050\u0020d\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079")
	}
	if _dagc == nil {
		return nil, nil, nil
	}
	_bgaa, _ddec := _cde.GetArray(_acbea.Rect)
	if !_ddec || _bgaa.Len() != 4 {
		return nil, nil, _ceg.New("\u0072\u0065\u0063t\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_efdb, _cbag := NewPdfRectangle(*_bgaa)
	if _cbag != nil {
		return nil, nil, _cbag
	}
	_abadg := _cde.TraceToDirectObject(_dagc.Get("\u004e"))
	switch _ffggf := _abadg.(type) {
	case *_cde.PdfObjectStream:
		_dgfbb := _ffggf
		_dbed, _cfcg := NewXObjectFormFromStream(_dgfbb)
		return _dbed, _efdb, _cfcg
	case *_cde.PdfObjectDictionary:
		_cacdg := _ffggf
		_gdfa, _defae := _cde.GetName(_acbea.AS)
		if !_defae {
			return nil, nil, nil
		}
		if _cacdg.Get(*_gdfa) == nil {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0041\u0053\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
			return nil, nil, nil
		}
		_bdceg, _defae := _cde.GetStream(_cacdg.Get(*_gdfa))
		if !_defae {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006ce \u0074\u006f\u0020\u0061\u0063\u0063e\u0073\u0073\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0066\u006f\u0072 \u0025\u0076", _gdfa)
			return nil, nil, _ceg.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		}
		_aedd, _gfcc := NewXObjectFormFromStream(_bdceg)
		return _aedd, _efdb, _gfcc
	}
	_ad.Log.Debug("\u0049\u006e\u0076\u0061li\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u004e\u003a\u0020%\u0054", _abadg)
	return nil, nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
}
func (_ffdeg *PdfWriter) updateObjectNumbers() {
	_geabf := _ffdeg.ObjNumOffset
	_ffda := 0
	for _, _begabb := range _ffdeg._egbccc {
		_accgf := int64(_ffda + 1 + _geabf)
		_dedba := true
		if _ffdeg._aabfe {
			if _adgca, _gcgfe := _ffdeg._dgad[_begabb]; _gcgfe {
				_accgf = _adgca
				_dedba = false
			}
		}
		switch _gbggf := _begabb.(type) {
		case *_cde.PdfIndirectObject:
			_gbggf.ObjectNumber = _accgf
			_gbggf.GenerationNumber = 0
		case *_cde.PdfObjectStream:
			_gbggf.ObjectNumber = _accgf
			_gbggf.GenerationNumber = 0
		case *_cde.PdfObjectStreams:
			_gbggf.ObjectNumber = _accgf
			_gbggf.GenerationNumber = 0
		default:
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u0020%\u0054\u0020\u002d\u0020\u0073\u006b\u0069p\u0070\u0069\u006e\u0067", _gbggf)
			continue
		}
		if _dedba {
			_ffda++
		}
	}
	_bdeea := func(_bgcgb _cde.PdfObject) int64 {
		switch _cbdee := _bgcgb.(type) {
		case *_cde.PdfIndirectObject:
			return _cbdee.ObjectNumber
		case *_cde.PdfObjectStream:
			return _cbdee.ObjectNumber
		case *_cde.PdfObjectStreams:
			return _cbdee.ObjectNumber
		}
		return 0
	}
	_dd.SliceStable(_ffdeg._egbccc, func(_dgddb, _abbcc int) bool { return _bdeea(_ffdeg._egbccc[_dgddb]) < _bdeea(_ffdeg._egbccc[_abbcc]) })
}

// NewPdfColorspaceSpecialPattern returns a new pattern color.
func NewPdfColorspaceSpecialPattern() *PdfColorspaceSpecialPattern {
	return &PdfColorspaceSpecialPattern{}
}
func (_fafdeb *PdfFilespec) getDict() *_cde.PdfObjectDictionary {
	if _aeea, _dfbfef := _fafdeb._bgac.(*_cde.PdfIndirectObject); _dfbfef {
		_egae, _ebfee := _aeea.PdfObject.(*_cde.PdfObjectDictionary)
		if !_ebfee {
			return nil
		}
		return _egae
	} else if _bggfd, _fbgbe := _fafdeb._bgac.(*_cde.PdfObjectDictionary); _fbgbe {
		return _bggfd
	} else {
		_ad.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074 \u0074\u0079p\u0065\u0020(\u0025T\u0029", _fafdeb._bgac)
		return nil
	}
}
func (_cbda *DSS) addCRLs(_fgecg [][]byte) ([]*_cde.PdfObjectStream, error) {
	return _cbda.add(&_cbda.CRLs, _cbda._fbbd, _fgecg)
}

// NewPdfAnnotationProjection returns a new projection annotation.
func NewPdfAnnotationProjection() *PdfAnnotationProjection {
	_ceb := NewPdfAnnotation()
	_fgf := &PdfAnnotationProjection{}
	_fgf.PdfAnnotation = _ceb
	_fgf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ceb.SetContext(_fgf)
	return _fgf
}
func (_bgbe *PdfReader) newPdfAnnotationInkFromDict(_eaca *_cde.PdfObjectDictionary) (*PdfAnnotationInk, error) {
	_bdb := PdfAnnotationInk{}
	_gda, _baf := _bgbe.newPdfAnnotationMarkupFromDict(_eaca)
	if _baf != nil {
		return nil, _baf
	}
	_bdb.PdfAnnotationMarkup = _gda
	_bdb.InkList = _eaca.Get("\u0049n\u006b\u004c\u0069\u0073\u0074")
	_bdb.BS = _eaca.Get("\u0042\u0053")
	return &_bdb, nil
}

// FullName returns the full name of the field as in rootname.parentname.partialname.
func (_geda *PdfField) FullName() (string, error) {
	var _fbbca _ede.Buffer
	_ffae := []string{}
	if _geda.T != nil {
		_ffae = append(_ffae, _geda.T.Decoded())
	}
	_bafbc := map[*PdfField]bool{}
	_bafbc[_geda] = true
	_ggebe := _geda.Parent
	for _ggebe != nil {
		if _, _begf := _bafbc[_ggebe]; _begf {
			return _fbbca.String(), _ceg.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		if _ggebe.T == nil {
			return _fbbca.String(), _ceg.New("\u0066\u0069el\u0064\u0020\u0070a\u0072\u0074\u0069\u0061l n\u0061me\u0020\u0028\u0054\u0029\u0020\u006e\u006ft \u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064")
		}
		_ffae = append(_ffae, _ggebe.T.Decoded())
		_bafbc[_ggebe] = true
		_ggebe = _ggebe.Parent
	}
	for _accg := len(_ffae) - 1; _accg >= 0; _accg-- {
		_fbbca.WriteString(_ffae[_accg])
		if _accg > 0 {
			_fbbca.WriteString("\u002e")
		}
	}
	return _fbbca.String(), nil
}
func (_egf *PdfReader) newPdfAnnotationPrinterMarkFromDict(_abgg *_cde.PdfObjectDictionary) (*PdfAnnotationPrinterMark, error) {
	_agdc := PdfAnnotationPrinterMark{}
	_agdc.MN = _abgg.Get("\u004d\u004e")
	return &_agdc, nil
}

// PdfActionTrans represents a trans action.
type PdfActionTrans struct {
	*PdfAction
	Trans _cde.PdfObject
}

// NewCompositePdfFontFromTTFFile loads a composite font from a TTF font file. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTFFile.
func NewCompositePdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_bceb, _cgbg := _db.Open(filePath)
	if _cgbg != nil {
		_ad.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006f\u0070\u0065\u006e\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _cgbg)
		return nil, _cgbg
	}
	defer _bceb.Close()
	return NewCompositePdfFontFromTTF(_bceb)
}

// DecodeArray returns the range of color component values in the ICCBased colorspace.
func (_bgf *PdfColorspaceICCBased) DecodeArray() []float64 { return _bgf.Range }

// PdfAnnotationMovie represents Movie annotations.
// (Section 12.5.6.17).
type PdfAnnotationMovie struct {
	*PdfAnnotation
	T     _cde.PdfObject
	Movie _cde.PdfObject
	A     _cde.PdfObject
}

// EncryptOptions represents encryption options for an output PDF.
type EncryptOptions struct {
	Permissions _ccg.Permissions
	Algorithm   EncryptionAlgorithm
}

// GetContainingPdfObject returns the container of the resources object (indirect object).
func (_fdffb *PdfPageResources) GetContainingPdfObject() _cde.PdfObject { return _fdffb._eaegg }

// GetDocMDPPermission returns the DocMDP level of the restrictions
func (_deggd *PdfSignature) GetDocMDPPermission() (_cdc.DocMDPPermission, bool) {
	for _, _dfgde := range _deggd.Reference.Elements() {
		if _gdbdg, _face := _cde.GetDict(_dfgde); _face {
			if _ffgce, _gebc := _cde.GetNameVal(_gdbdg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _gebc && _ffgce == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				if _bfeegd, _dgaeb := _cde.GetDict(_gdbdg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _dgaeb {
					if P, _eefadd := _cde.GetIntVal(_bfeegd.Get("\u0050")); _eefadd {
						return _cdc.DocMDPPermission(P), true
					}
				}
			}
		}
	}
	return 0, false
}

// GetAction returns the PDF action for the annotation link.
func (_bfba *PdfAnnotationLink) GetAction() (*PdfAction, error) {
	if _bfba._beg != nil {
		return _bfba._beg, nil
	}
	if _bfba.A == nil {
		return nil, nil
	}
	if _bfba._gegc == nil {
		return nil, nil
	}
	_eda, _cgg := _bfba._gegc.loadAction(_bfba.A)
	if _cgg != nil {
		return nil, _cgg
	}
	_bfba._beg = _eda
	return _bfba._beg, nil
}

// ToGoTime returns the date in time.Time format.
func (_dfbef PdfDate) ToGoTime() _ce.Time {
	_cccb := int(_dfbef._gafaf*60*60 + _dfbef._bbcgcf*60)
	switch _dfbef._efdde {
	case '-':
		_cccb = -_cccb
	case 'Z':
		_cccb = 0
	}
	_fabegg := _ee.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _dfbef._efdde, _dfbef._gafaf, _dfbef._bbcgcf)
	_gfbgc := _ce.FixedZone(_fabegg, _cccb)
	return _ce.Date(int(_dfbef._aedbfg), _ce.Month(_dfbef._babce), int(_dfbef._edbef), int(_dfbef._aeac), int(_dfbef._accgb), int(_dfbef._ebfeg), 0, _gfbgc)
}

// NewPdfAnnotationMovie returns a new movie annotation.
func NewPdfAnnotationMovie() *PdfAnnotationMovie {
	_abdg := NewPdfAnnotation()
	_ecb := &PdfAnnotationMovie{}
	_ecb.PdfAnnotation = _abdg
	_abdg.SetContext(_ecb)
	return _ecb
}

// GetContext returns a reference to the subpattern entry: either PdfTilingPattern or PdfShadingPattern.
func (_gffbc *PdfPattern) GetContext() PdfModel { return _gffbc._abddb }

// ToPdfObject implements interface PdfModel.
func (_aae *PdfAnnotationSquare) ToPdfObject() _cde.PdfObject {
	_aae.PdfAnnotation.ToPdfObject()
	_efbf := _aae._bddg
	_gcbg := _efbf.PdfObject.(*_cde.PdfObjectDictionary)
	if _aae.PdfAnnotationMarkup != nil {
		_aae.PdfAnnotationMarkup.appendToPdfDictionary(_gcbg)
	}
	_gcbg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0053\u0071\u0075\u0061\u0072\u0065"))
	_gcbg.SetIfNotNil("\u0042\u0053", _aae.BS)
	_gcbg.SetIfNotNil("\u0049\u0043", _aae.IC)
	_gcbg.SetIfNotNil("\u0042\u0045", _aae.BE)
	_gcbg.SetIfNotNil("\u0052\u0044", _aae.RD)
	return _efbf
}

// ImageToGray returns a new grayscale image based on the passed in RGB image.
func (_faea *PdfColorspaceDeviceRGB) ImageToGray(img Image) (Image, error) {
	if img.ColorComponents != 3 {
		return img, _ceg.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u0020\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042")
	}
	_bfgdg, _eccf := _ff.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._deegf, img._aaafb)
	if _eccf != nil {
		return img, _eccf
	}
	_gbdea, _eccf := _ff.GrayConverter.Convert(_bfgdg)
	if _eccf != nil {
		return img, _eccf
	}
	return _bddb(_gbdea.Base()), nil
}

// GetPdfVersion gets the version of the PDF used within this document.
func (_fddag *PdfWriter) GetPdfVersion() string { return _fddag.getPdfVersion() }

// ToPdfObject returns the PDF representation of the page resources.
func (_dcfac *PdfPageResources) ToPdfObject() _cde.PdfObject {
	_bdgab := _dcfac._eaegg
	_bdgab.SetIfNotNil("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _dcfac.ExtGState)
	if _dcfac._bfff != nil {
		_dcfac.ColorSpace = _dcfac._bfff.ToPdfObject()
	}
	_bdgab.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _dcfac.ColorSpace)
	_bdgab.SetIfNotNil("\u0050a\u0074\u0074\u0065\u0072\u006e", _dcfac.Pattern)
	_bdgab.SetIfNotNil("\u0053h\u0061\u0064\u0069\u006e\u0067", _dcfac.Shading)
	_bdgab.SetIfNotNil("\u0058O\u0062\u006a\u0065\u0063\u0074", _dcfac.XObject)
	_bdgab.SetIfNotNil("\u0046\u006f\u006e\u0074", _dcfac.Font)
	_bdgab.SetIfNotNil("\u0050r\u006f\u0063\u0053\u0065\u0074", _dcfac.ProcSet)
	_bdgab.SetIfNotNil("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", _dcfac.Properties)
	return _bdgab
}

// DecodeArray returns the component range values for the DeviceN colorspace.
// [0 1.0 0 1.0 ...] for each color component.
func (_bddgf *PdfColorspaceDeviceN) DecodeArray() []float64 {
	var _aaccb []float64
	for _adbd := 0; _adbd < _bddgf.GetNumComponents(); _adbd++ {
		_aaccb = append(_aaccb, 0.0, 1.0)
	}
	return _aaccb
}

// FieldFilterFunc represents a PDF field filtering function. If the function
// returns true, the PDF field is kept, otherwise it is discarded.
type FieldFilterFunc func(*PdfField) bool

func (_eabc *PdfColorspaceDeviceCMYK) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
}

// NewPdfAnnotationRichMedia returns a new rich media annotation.
func NewPdfAnnotationRichMedia() *PdfAnnotationRichMedia {
	_fecb := NewPdfAnnotation()
	_agdg := &PdfAnnotationRichMedia{}
	_agdg.PdfAnnotation = _fecb
	_fecb.SetContext(_agdg)
	return _agdg
}
func _beed(_adea _cde.PdfObject) (*PdfFilespec, error) {
	if _adea == nil {
		return nil, nil
	}
	return NewPdfFilespecFromObj(_adea)
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for an RGB device.
func (_debf *PdfColorspaceDeviceRGB) GetNumComponents() int { return 3 }

// ToPdfObject returns a *PdfIndirectObject containing a *PdfObjectArray representation of the DeviceN colorspace.
// Format: [/DeviceN names alternateSpace tintTransform]
//     or: [/DeviceN names alternateSpace tintTransform attributes]
func (_ggfe *PdfColorspaceDeviceN) ToPdfObject() _cde.PdfObject {
	_aeaab := _cde.MakeArray(_cde.MakeName("\u0044e\u0076\u0069\u0063\u0065\u004e"))
	_aeaab.Append(_ggfe.ColorantNames)
	_aeaab.Append(_ggfe.AlternateSpace.ToPdfObject())
	_aeaab.Append(_ggfe.TintTransform.ToPdfObject())
	if _ggfe.Attributes != nil {
		_aeaab.Append(_ggfe.Attributes.ToPdfObject())
	}
	if _ggfe._cdac != nil {
		_ggfe._cdac.PdfObject = _aeaab
		return _ggfe._cdac
	}
	return _aeaab
}
func _aadfe(_abaad *XObjectImage) error {
	if _abaad.SMask == nil {
		return nil
	}
	_dggcf, _ddgcc := _abaad.SMask.(*_cde.PdfObjectStream)
	if !_ddgcc {
		_ad.Log.Debug("\u0053\u004da\u0073\u006b\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0053\u0074re\u0061\u006d")
		return _cde.ErrTypeError
	}
	_fdgbce := _dggcf.PdfObjectDictionary
	_baaee := _fdgbce.Get("\u004d\u0061\u0074t\u0065")
	if _baaee == nil {
		return nil
	}
	_gddba, _ggca := _bagfg(_baaee.(*_cde.PdfObjectArray))
	if _ggca != nil {
		return _ggca
	}
	_decfcc := _cde.MakeArrayFromFloats([]float64{_gddba})
	_fdgbce.SetIfNotNil("\u004d\u0061\u0074t\u0065", _decfcc)
	return nil
}
func (_dbffa *PdfWriter) getPdfVersion() string {
	return _ee.Sprintf("\u0025\u0064\u002e%\u0064", _dbffa._cgdcc.Major, _dbffa._cgdcc.Minor)
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// L (range 0-100), A (range -100-100) and B (range -100-100) components of
// the color.
func (_fgc *PdfColorspaceLab) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_aeggf := vals[0]
	if _aeggf < 0.0 || _aeggf > 100.0 {
		_ad.Log.Debug("\u004c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030-\u0031\u0030\u0030\u0029", _aeggf)
		return nil, ErrColorOutOfRange
	}
	_gdgaaf := vals[1]
	_eafbc := float64(-100)
	_ebdg := float64(100)
	if len(_fgc.Range) > 1 {
		_eafbc = _fgc.Range[0]
		_ebdg = _fgc.Range[1]
	}
	if _gdgaaf < _eafbc || _gdgaaf > _ebdg {
		_ad.Log.Debug("\u0041\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _gdgaaf, _eafbc, _ebdg)
		return nil, ErrColorOutOfRange
	}
	_gcbcd := vals[2]
	_ggc := float64(-100)
	_cef := float64(100)
	if len(_fgc.Range) > 3 {
		_ggc = _fgc.Range[2]
		_cef = _fgc.Range[3]
	}
	if _gcbcd < _ggc || _gcbcd > _cef {
		_ad.Log.Debug("\u0062\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _gcbcd, _ggc, _cef)
		return nil, ErrColorOutOfRange
	}
	_gabd := NewPdfColorLab(_aeggf, _gdgaaf, _gcbcd)
	return _gabd, nil
}

// ToPdfObject implements interface PdfModel.
func (_cbdb *PdfActionJavaScript) ToPdfObject() _cde.PdfObject {
	_cbdb.PdfAction.ToPdfObject()
	_cdef := _cbdb._bc
	_cgd := _cdef.PdfObject.(*_cde.PdfObjectDictionary)
	_cgd.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeJavaScript)))
	_cgd.SetIfNotNil("\u004a\u0053", _cbdb.JS)
	return _cdef
}

// GetContentStream returns the XObject Form's content stream.
func (_gafdf *XObjectForm) GetContentStream() ([]byte, error) {
	_bcedc, _bgdeg := _cde.DecodeStream(_gafdf._ecfbd)
	if _bgdeg != nil {
		return nil, _bgdeg
	}
	return _bcedc, nil
}

// SignatureHandlerDocMDPParams describe the specific parameters for the SignatureHandlerEx
// These parameters describe how to check the difference between revisions.
// Revisions of the document get from the PdfParser.
type SignatureHandlerDocMDPParams struct {
	Parser     *_cde.PdfParser
	DiffPolicy _cdc.DiffPolicy
}

// ToPdfObject implements interface PdfModel.
func (_fcbf *PdfAnnotationProjection) ToPdfObject() _cde.PdfObject {
	_fcbf.PdfAnnotation.ToPdfObject()
	_ecaf := _fcbf._bddg
	_dbaf := _ecaf.PdfObject.(*_cde.PdfObjectDictionary)
	_fcbf.PdfAnnotationMarkup.appendToPdfDictionary(_dbaf)
	return _ecaf
}

// ToPdfObject implements interface PdfModel.
func (_feac *PdfAnnotationWidget) ToPdfObject() _cde.PdfObject {
	_feac.PdfAnnotation.ToPdfObject()
	_abfb := _feac._bddg
	_dcff := _abfb.PdfObject.(*_cde.PdfObjectDictionary)
	if _feac._bga {
		return _abfb
	}
	_feac._bga = true
	_dcff.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0057\u0069\u0064\u0067\u0065\u0074"))
	_dcff.SetIfNotNil("\u0048", _feac.H)
	_dcff.SetIfNotNil("\u004d\u004b", _feac.MK)
	_dcff.SetIfNotNil("\u0041", _feac.A)
	_dcff.SetIfNotNil("\u0041\u0041", _feac.AA)
	_dcff.SetIfNotNil("\u0042\u0053", _feac.BS)
	_edca := _feac.Parent
	if _feac._dbf != nil {
		if _feac._dbf._afgc == _feac._bddg {
			_feac._dbf.ToPdfObject()
		}
		_edca = _feac._dbf.GetContainingPdfObject()
	}
	if _edca != _abfb {
		_dcff.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _edca)
	}
	_feac._bga = false
	return _abfb
}

// ToPdfObject implements interface PdfModel.
func (_bad *PdfAnnotationText) ToPdfObject() _cde.PdfObject {
	_bad.PdfAnnotation.ToPdfObject()
	_ggfc := _bad._bddg
	_aaaf := _ggfc.PdfObject.(*_cde.PdfObjectDictionary)
	if _bad.PdfAnnotationMarkup != nil {
		_bad.PdfAnnotationMarkup.appendToPdfDictionary(_aaaf)
	}
	_aaaf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0054\u0065\u0078\u0074"))
	_aaaf.SetIfNotNil("\u004f\u0070\u0065\u006e", _bad.Open)
	_aaaf.SetIfNotNil("\u004e\u0061\u006d\u0065", _bad.Name)
	_aaaf.SetIfNotNil("\u0053\u0074\u0061t\u0065", _bad.State)
	_aaaf.SetIfNotNil("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c", _bad.StateModel)
	return _ggfc
}

// NewPdfReader returns a new PdfReader for an input io.ReadSeeker interface. Can be used to read PDF from
// memory or file. Immediately loads and traverses the PDF structure including pages and page contents (if
// not encrypted). Loads entire document structure into memory.
// Alternatively a lazy-loading reader can be created with NewPdfReaderLazy which loads only references,
// and references are loaded from disk into memory on an as-needed basis.
func NewPdfReader(rs _f.ReadSeeker) (*PdfReader, error) {
	const _gdfbe = "\u006do\u0064e\u006c\u003a\u004e\u0065\u0077P\u0064\u0066R\u0065\u0061\u0064\u0065\u0072"
	return _gfceb(rs, &ReaderOpts{}, false, _gdfbe)
}

// SetContext sets the sub action (context).
func (_cac *PdfAction) SetContext(ctx PdfModel) { _cac._bgd = ctx }

// ToPdfObject implements interface PdfModel.
func (_bdd *PdfActionURI) ToPdfObject() _cde.PdfObject {
	_bdd.PdfAction.ToPdfObject()
	_ddb := _bdd._bc
	_ddbg := _ddb.PdfObject.(*_cde.PdfObjectDictionary)
	_ddbg.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeURI)))
	_ddbg.SetIfNotNil("\u0055\u0052\u0049", _bdd.URI)
	_ddbg.SetIfNotNil("\u0049\u0073\u004da\u0070", _bdd.IsMap)
	return _ddb
}

// GetFontDescriptor returns the font descriptor for `font`.
func (_bgfb PdfFont) GetFontDescriptor() (*PdfFontDescriptor, error) {
	return _bgfb._gbcff.getFontDescriptor(), nil
}

// NewCompositePdfFontFromTTF loads a composite TTF font. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTF.
func NewCompositePdfFontFromTTF(r _f.ReadSeeker) (*PdfFont, error) {
	_ffafc, _gagca := _bb.ReadAll(r)
	if _gagca != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _gagca)
		return nil, _gagca
	}
	_dcfa, _gagca := _fe.TtfParse(_ede.NewReader(_ffafc))
	if _gagca != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067 \u0074\u0074\u0066\u0020\u0066\u006f\u006et\u003a\u0020\u0025\u0076", _gagca)
		return nil, _gagca
	}
	_ebggd := &pdfCIDFontType2{fontCommon: fontCommon{_dcbc: "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032"}, CIDToGIDMap: _cde.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079")}
	if len(_dcfa.Widths) <= 0 {
		return nil, _ceg.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_gbffbf := 1000.0 / float64(_dcfa.UnitsPerEm)
	_gfab := _gbffbf * float64(_dcfa.Widths[0])
	_dfgfc := make(map[rune]int)
	_cfagb := make(map[_fe.GID]int)
	_bbdb := _fe.GID(len(_dcfa.Widths))
	for _eccc, _fbdfd := range _dcfa.Chars {
		if _fbdfd > _bbdb-1 {
			continue
		}
		_fbcg := int(_gbffbf * float64(_dcfa.Widths[_fbdfd]))
		_dfgfc[_eccc] = _fbcg
		_cfagb[_fbdfd] = _fbcg
	}
	_ebggd._dfefe = _dfgfc
	_ebggd.DW = _cde.MakeInteger(int64(_gfab))
	_geee := _eefgd(_cfagb, uint16(_bbdb))
	_ebggd.W = _cde.MakeIndirectObject(_geee)
	_beea := _cde.MakeDict()
	_beea.Set("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067", _cde.MakeString("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"))
	_beea.Set("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079", _cde.MakeString("\u0041\u0064\u006fb\u0065"))
	_beea.Set("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074", _cde.MakeInteger(0))
	_ebggd.CIDSystemInfo = _beea
	_ggbf := &PdfFontDescriptor{FontName: _cde.MakeName(_dcfa.PostScriptName), Ascent: _cde.MakeFloat(_gbffbf * float64(_dcfa.TypoAscender)), Descent: _cde.MakeFloat(_gbffbf * float64(_dcfa.TypoDescender)), CapHeight: _cde.MakeFloat(_gbffbf * float64(_dcfa.CapHeight)), FontBBox: _cde.MakeArrayFromFloats([]float64{_gbffbf * float64(_dcfa.Xmin), _gbffbf * float64(_dcfa.Ymin), _gbffbf * float64(_dcfa.Xmax), _gbffbf * float64(_dcfa.Ymax)}), ItalicAngle: _cde.MakeFloat(_dcfa.ItalicAngle), MissingWidth: _cde.MakeFloat(_gfab)}
	_dabdc, _gagca := _cde.MakeStream(_ffafc, _cde.NewFlateEncoder())
	if _gagca != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _gagca)
		return nil, _gagca
	}
	_dabdc.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _cde.MakeInteger(int64(len(_ffafc))))
	_ggbf.FontFile2 = _dabdc
	if _dcfa.Bold {
		_ggbf.StemV = _cde.MakeInteger(120)
	} else {
		_ggbf.StemV = _cde.MakeInteger(70)
	}
	_dbfcg := _edbca
	if _dcfa.IsFixedPitch {
		_dbfcg |= _gbcd
	}
	if _dcfa.ItalicAngle != 0 {
		_dbfcg |= _ecebb
	}
	_ggbf.Flags = _cde.MakeInteger(int64(_dbfcg))
	_ebggd._eeab = _dcfa.PostScriptName
	_ebggd._fagf = _ggbf
	_dbegb := pdfFontType0{fontCommon: fontCommon{_dcbc: "\u0054\u0079\u0070e\u0030", _eeab: _dcfa.PostScriptName}, DescendantFont: &PdfFont{_gbcff: _ebggd}, Encoding: _cde.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048"), _cffef: _dcfa.NewEncoder()}
	if len(_dcfa.Chars) > 0 {
		_ebada := make(map[_fb.CharCode]rune, len(_dcfa.Chars))
		for _fbga, _addgc := range _dcfa.Chars {
			_facde := _fb.CharCode(_addgc)
			if _daadg, _fbcbf := _ebada[_facde]; !_fbcbf || (_fbcbf && _daadg > _fbga) {
				_ebada[_facde] = _fbga
			}
		}
		_dbegb._ggebg = _fb.NewToUnicodeCMap(_ebada)
	}
	_eaaff := PdfFont{_gbcff: &_dbegb}
	return &_eaaff, nil
}
func (_gbbe *PdfAcroForm) fill(_acae FieldValueProvider, _cfeb FieldAppearanceGenerator) error {
	if _gbbe == nil {
		return nil
	}
	_ebcc, _bafa := _acae.FieldValues()
	if _bafa != nil {
		return _bafa
	}
	for _, _abdbe := range _gbbe.AllFields() {
		_egcb := _abdbe.PartialName()
		_bgad, _ddbf := _ebcc[_egcb]
		if !_ddbf {
			if _fecf, _cgda := _abdbe.FullName(); _cgda == nil {
				_bgad, _ddbf = _ebcc[_fecf]
			}
		}
		if !_ddbf {
			_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _egcb)
			continue
		}
		if _baab := _deaee(_abdbe, _bgad); _baab != nil {
			return _baab
		}
		if _cfeb == nil {
			continue
		}
		for _, _eabcb := range _abdbe.Annotations {
			_bbcgc, _fabc := _cfeb.GenerateAppearanceDict(_gbbe, _abdbe, _eabcb)
			if _fabc != nil {
				return _fabc
			}
			_eabcb.AP = _bbcgc
			_eabcb.ToPdfObject()
		}
	}
	return nil
}

// Val returns the color value.
func (_fgfab *PdfColorDeviceGray) Val() float64 { return float64(*_fgfab) }

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
	Metadata *_cde.PdfObjectStream
	Data     []byte
	_cffb    *_cde.PdfIndirectObject
	_fdea    *_cde.PdfObjectStream
}

// NewPdfAction returns an initialized generic PDF action model.
func NewPdfAction() *PdfAction {
	_dge := &PdfAction{}
	_dge._bc = _cde.MakeIndirectObject(_cde.MakeDict())
	return _dge
}

// NewPdfAnnotation3D returns a new 3d annotation.
func NewPdfAnnotation3D() *PdfAnnotation3D {
	_cgdb := NewPdfAnnotation()
	_gga := &PdfAnnotation3D{}
	_gga.PdfAnnotation = _cgdb
	_cgdb.SetContext(_gga)
	return _gga
}

// PdfAnnotationProjection represents Projection annotations.
type PdfAnnotationProjection struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
}

// ToPdfObject returns the PDF representation of the function.
func (_cdfc *PdfFunctionType0) ToPdfObject() _cde.PdfObject {
	if _cdfc._fefbf == nil {
		_cdfc._fefbf = &_cde.PdfObjectStream{}
	}
	_daagag := _cde.MakeDict()
	_daagag.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _cde.MakeInteger(0))
	_ecbgf := &_cde.PdfObjectArray{}
	for _, _fgbcc := range _cdfc.Domain {
		_ecbgf.Append(_cde.MakeFloat(_fgbcc))
	}
	_daagag.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _ecbgf)
	_agbac := &_cde.PdfObjectArray{}
	for _, _ggaff := range _cdfc.Range {
		_agbac.Append(_cde.MakeFloat(_ggaff))
	}
	_daagag.Set("\u0052\u0061\u006eg\u0065", _agbac)
	_acbcd := &_cde.PdfObjectArray{}
	for _, _gfgbg := range _cdfc.Size {
		_acbcd.Append(_cde.MakeInteger(int64(_gfgbg)))
	}
	_daagag.Set("\u0053\u0069\u007a\u0065", _acbcd)
	_daagag.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065", _cde.MakeInteger(int64(_cdfc.BitsPerSample)))
	if _cdfc.Order != 1 {
		_daagag.Set("\u004f\u0072\u0064e\u0072", _cde.MakeInteger(int64(_cdfc.Order)))
	}
	_daagag.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cde.MakeInteger(int64(len(_cdfc._gfbb))))
	_cdfc._fefbf.Stream = _cdfc._gfbb
	_cdfc._fefbf.PdfObjectDictionary = _daagag
	return _cdfc._fefbf
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a CalRGB device.
func (_fadc *PdfColorspaceCalRGB) GetNumComponents() int { return 3 }

// NewPdfAnnotationLine returns a new line annotation.
func NewPdfAnnotationLine() *PdfAnnotationLine {
	_aed := NewPdfAnnotation()
	_egd := &PdfAnnotationLine{}
	_egd.PdfAnnotation = _aed
	_egd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_aed.SetContext(_egd)
	return _egd
}
func (_eefg *PdfReader) newPdfAnnotationRichMediaFromDict(_eadb *_cde.PdfObjectDictionary) (*PdfAnnotationRichMedia, error) {
	_ccgf := &PdfAnnotationRichMedia{}
	_ccgf.RichMediaSettings = _eadb.Get("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073")
	_ccgf.RichMediaContent = _eadb.Get("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074")
	return _ccgf, nil
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

// FieldImageProvider provides fields images for specified fields.
type FieldImageProvider interface {
	FieldImageValues() (map[string]*Image, error)
}

// NewPdfInfoFromObject creates a new PdfInfo from the input core.PdfObject.
func NewPdfInfoFromObject(obj _cde.PdfObject) (*PdfInfo, error) {
	var _abad PdfInfo
	_cacdb, _cbcg := obj.(*_cde.PdfObjectDictionary)
	if !_cbcg {
		return nil, _ee.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", obj)
	}
	for _, _bagd := range _cacdb.Keys() {
		switch _bagd {
		case "\u0054\u0069\u0074l\u0065":
			_abad.Title, _ = _cde.GetString(_cacdb.Get("\u0054\u0069\u0074l\u0065"))
		case "\u0041\u0075\u0074\u0068\u006f\u0072":
			_abad.Author, _ = _cde.GetString(_cacdb.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
		case "\u0053u\u0062\u006a\u0065\u0063\u0074":
			_abad.Subject, _ = _cde.GetString(_cacdb.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
		case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
			_abad.Keywords, _ = _cde.GetString(_cacdb.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
		case "\u0043r\u0065\u0061\u0074\u006f\u0072":
			_abad.Creator, _ = _cde.GetString(_cacdb.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
		case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
			_abad.Producer, _ = _cde.GetString(_cacdb.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
		case "\u0054r\u0061\u0070\u0070\u0065\u0064":
			_abad.Trapped, _ = _cde.GetName(_cacdb.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
		case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
			if _efedd, _fedg := _cde.GetString(_cacdb.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _fedg && _efedd.String() != "" {
				_afceaa, _gbaa := NewPdfDate(_efedd.String())
				if _gbaa != nil {
					return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _gbaa)
				}
				_abad.CreationDate = &_afceaa
			}
		case "\u004do\u0064\u0044\u0061\u0074\u0065":
			if _fbbg, _deaae := _cde.GetString(_cacdb.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _deaae && _fbbg.String() != "" {
				_adacd, _eecb := NewPdfDate(_fbbg.String())
				if _eecb != nil {
					return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _eecb)
				}
				_abad.ModifiedDate = &_adacd
			}
		default:
			_cbcgg, _ := _cde.GetString(_cacdb.Get(_bagd))
			if _abad._ccef == nil {
				_abad._ccef = _cde.MakeDict()
			}
			_abad._ccef.Set(_bagd, _cbcgg)
		}
	}
	return &_abad, nil
}

// GetPage returns the PdfPage model for the specified page number.
func (_edbe *PdfReader) GetPage(pageNumber int) (*PdfPage, error) {
	if _edbe._aggcgb.GetCrypter() != nil && !_edbe._aggcgb.IsAuthenticated() {
		return nil, _ee.Errorf("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	if len(_edbe._gbfgg) < pageNumber {
		return nil, _ceg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0028\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0074o\u006f\u0020\u0073\u0068\u006f\u0072\u0074\u0029")
	}
	_acbfe := pageNumber - 1
	if _acbfe < 0 {
		return nil, _ee.Errorf("\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0069\u006e\u0067\u0020\u006d\u0075\u0073t\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0061\u0074\u0020\u0031")
	}
	_fcbfa := _edbe.PageList[_acbfe]
	return _fcbfa, nil
}
func (_bgcd Image) getBase() _ff.ImageBase {
	return _ff.NewImageBase(int(_bgcd.Width), int(_bgcd.Height), int(_bgcd.BitsPerComponent), _bgcd.ColorComponents, _bgcd.Data, _bgcd._deegf, _bgcd._aaafb)
}

// SetBorderWidth sets the style's border width.
func (_ebge *PdfBorderStyle) SetBorderWidth(width float64) { _ebge.W = &width }

// AddWatermarkImage adds a watermark to the page.
func (_eabf *PdfPage) AddWatermarkImage(ximg *XObjectImage, opt WatermarkImageOptions) error {
	_ebdff, _acea := _eabf.GetMediaBox()
	if _acea != nil {
		return _acea
	}
	_bdcae := _ebdff.Urx - _ebdff.Llx
	_fddeba := _ebdff.Ury - _ebdff.Lly
	_gaaa := float64(*ximg.Width)
	_cacb := (_bdcae - _gaaa) / 2
	if opt.FitToWidth {
		_gaaa = _bdcae
		_cacb = 0
	}
	_cgdcd := _fddeba
	_dcec := float64(0)
	if opt.PreserveAspectRatio {
		_cgdcd = _gaaa * float64(*ximg.Height) / float64(*ximg.Width)
		_dcec = (_fddeba - _cgdcd) / 2
	}
	if _eabf.Resources == nil {
		_eabf.Resources = NewPdfPageResources()
	}
	_bbcdbc := 0
	_agbc := _cde.PdfObjectName(_ee.Sprintf("\u0049\u006d\u0077%\u0064", _bbcdbc))
	for _eabf.Resources.HasXObjectByName(_agbc) {
		_bbcdbc++
		_agbc = _cde.PdfObjectName(_ee.Sprintf("\u0049\u006d\u0077%\u0064", _bbcdbc))
	}
	_acea = _eabf.AddImageResource(_agbc, ximg)
	if _acea != nil {
		return _acea
	}
	_bbcdbc = 0
	_dfafe := _cde.PdfObjectName(_ee.Sprintf("\u0047\u0053\u0025\u0064", _bbcdbc))
	for _eabf.HasExtGState(_dfafe) {
		_bbcdbc++
		_dfafe = _cde.PdfObjectName(_ee.Sprintf("\u0047\u0053\u0025\u0064", _bbcdbc))
	}
	_fabba := _cde.MakeDict()
	_fabba.Set("\u0042\u004d", _cde.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	_fabba.Set("\u0043\u0041", _cde.MakeFloat(opt.Alpha))
	_fabba.Set("\u0063\u0061", _cde.MakeFloat(opt.Alpha))
	_acea = _eabf.AddExtGState(_dfafe, _fabba)
	if _acea != nil {
		return _acea
	}
	_faebc := _ee.Sprintf("\u0071\u000a"+"\u002f%\u0073\u0020\u0067\u0073\u000a"+"%\u002e\u0030\u0066\u0020\u0030\u00200\u0020\u0025\u002e\u0030\u0066\u0020\u0025\u002e\u0034f\u0020\u0025\u002e4\u0066 \u0063\u006d\u000a"+"\u002f%\u0073\u0020\u0044\u006f\u000a"+"\u0051", _dfafe, _gaaa, _cgdcd, _cacb, _dcec, _agbc)
	_eabf.AddContentStreamByString(_faebc)
	return nil
}

const (
	_gbcd  = 0x00001
	_bafg  = 0x00002
	_edbca = 0x00004
	_ffaf  = 0x00008
	_gbdf  = 0x00020
	_ecebb = 0x00040
	_affa  = 0x10000
	_fabfb = 0x20000
	_ddbcb = 0x40000
)

func (_dcdfe *PdfWriter) writeAcroFormFields() error {
	if _dcdfe._fcacfe == nil {
		return nil
	}
	_ad.Log.Trace("\u0057r\u0069t\u0069\u006e\u0067\u0020\u0061c\u0072\u006f \u0066\u006f\u0072\u006d\u0073")
	_ccgfg := _dcdfe._fcacfe.ToPdfObject()
	_ad.Log.Trace("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _ccgfg)
	_dcdfe._fedbb.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _ccgfg)
	_ebebe := _dcdfe.addObjects(_ccgfg)
	if _ebebe != nil {
		return _ebebe
	}
	return nil
}
func (_bdce *PdfAnnotation) String() string {
	_bgc := ""
	_edd, _acb := _bdce.ToPdfObject().(*_cde.PdfIndirectObject)
	if _acb {
		_bgc = _ee.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _bdce._bea, _edd.PdfObject.String())
	}
	return _bgc
}

// GetRuneMetrics returns the char metrics for a rune.
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
//                       well give them 0 width. There is no need for the bool return.
func (_agee *PdfFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_efcd := _agee.actualFont()
	if _efcd == nil {
		_ad.Log.Debug("ER\u0052\u004fR\u003a\u0020\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u004e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020f\u006fr\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079p\u0065=\u0025\u0023T", _agee._gbcff)
		return _fe.CharMetrics{}, false
	}
	if _geadd, _geagd := _efcd.GetRuneMetrics(r); _geagd {
		return _geadd, true
	}
	if _bfbac, _bfac := _agee.GetFontDescriptor(); _bfac == nil && _bfbac != nil {
		return _fe.CharMetrics{Wx: _bfbac._efgg}, true
	}
	_ad.Log.Debug("\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061r\u004d\u0065\u0074\u0072i\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _agee)
	return _fe.CharMetrics{}, false
}

const (
	BorderStyleSolid     BorderStyle = iota
	BorderStyleDashed    BorderStyle = iota
	BorderStyleBeveled   BorderStyle = iota
	BorderStyleInset     BorderStyle = iota
	BorderStyleUnderline BorderStyle = iota
)

// PdfAnnotationPrinterMark represents PrinterMark annotations.
// (Section 12.5.6.20).
type PdfAnnotationPrinterMark struct {
	*PdfAnnotation
	MN _cde.PdfObject
}

// AppendContentBytes creates a PDF stream from `cs` and appends it to the
// array of streams specified by the pages's Contents entry.
// If `wrapContents` is true, the content stream of the page is wrapped using
// a `q/Q` operator pair, so that its state does not affect the appended
// content stream.
func (_gcbbg *PdfPage) AppendContentBytes(cs []byte, wrapContents bool) error {
	_gacbc := _gcbbg.GetContentStreamObjs()
	wrapContents = wrapContents && len(_gacbc) > 0
	_cfbc := _cde.NewFlateEncoder()
	_aafaa := _cde.MakeArray()
	if wrapContents {
		_cdfca, _bdcf := _cde.MakeStream([]byte("\u0071\u000a"), _cfbc)
		if _bdcf != nil {
			return _bdcf
		}
		_aafaa.Append(_cdfca)
	}
	_aafaa.Append(_gacbc...)
	if wrapContents {
		_fcffd, _cggcd := _cde.MakeStream([]byte("\u000a\u0051\u000a"), _cfbc)
		if _cggcd != nil {
			return _cggcd
		}
		_aafaa.Append(_fcffd)
	}
	_gffbg, _bgfc := _cde.MakeStream(cs, _cfbc)
	if _bgfc != nil {
		return _bgfc
	}
	_aafaa.Append(_gffbg)
	_gcbbg.Contents = _aafaa
	return nil
}
func _dcfg(_cbgf _cde.PdfObject) (*PdfColorspaceSpecialSeparation, error) {
	_beedd := NewPdfColorspaceSpecialSeparation()
	if _fgae, _geea := _cbgf.(*_cde.PdfIndirectObject); _geea {
		_beedd._bcdc = _fgae
	}
	_cbgf = _cde.TraceToDirectObject(_cbgf)
	_adaa, _gegge := _cbgf.(*_cde.PdfObjectArray)
	if !_gegge {
		return nil, _ee.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062je\u0063\u0074")
	}
	if _adaa.Len() != 4 {
		return nil, _ee.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074i\u006f\u006e \u0043\u0053\u003a\u0020\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
	}
	_cbgf = _adaa.Get(0)
	_geeb, _gegge := _cbgf.(*_cde.PdfObjectName)
	if !_gegge {
		return nil, _ee.Errorf("\u0073\u0065\u0070ar\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a \u0069n\u0076a\u006ci\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_geeb != "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e" {
		return nil, _ee.Errorf("\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020w\u0072o\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	_cbgf = _adaa.Get(1)
	_geeb, _gegge = _cbgf.(*_cde.PdfObjectName)
	if !_gegge {
		return nil, _ee.Errorf("\u0073\u0065pa\u0072\u0061\u0074i\u006f\u006e\u0020\u0043S: \u0049nv\u0061\u006c\u0069\u0064\u0020\u0063\u006flo\u0072\u0061\u006e\u0074\u0020\u006e\u0061m\u0065")
	}
	_beedd.ColorantName = _geeb
	_cbgf = _adaa.Get(2)
	_bdcbc, _fbae := NewPdfColorspaceFromPdfObject(_cbgf)
	if _fbae != nil {
		return nil, _fbae
	}
	_beedd.AlternateSpace = _bdcbc
	_eafa, _fbae := _cfdbb(_adaa.Get(3))
	if _fbae != nil {
		return nil, _fbae
	}
	_beedd.TintTransform = _eafa
	return _beedd, nil
}
func _gfceb(_fdaa _f.ReadSeeker, _geefc *ReaderOpts, _cagded bool, _ddage string) (*PdfReader, error) {
	if _geefc == nil {
		_geefc = NewReaderOpts()
	}
	_cdfacb := *_geefc
	_cafcaf := &PdfReader{_caecc: _fdaa, _efbdd: map[_cde.PdfObject]struct{}{}, _bedfa: _acaga(), _cdgee: _geefc.LazyLoad, _gcegag: _geefc.ComplianceMode, _cgebe: _cagded, _cbbc: &_cdfacb}
	_fbbad, _fffee := _bbeaa("\u0072")
	if _fffee != nil {
		return nil, _fffee
	}
	_cafcaf._gfbcg = _fbbad
	var _fgaa *_cde.PdfParser
	if !_cafcaf._gcegag {
		_fgaa, _fffee = _cde.NewParser(_fdaa)
	} else {
		_fgaa, _fffee = _cde.NewCompliancePdfParser(_fdaa)
	}
	if _fffee != nil {
		return nil, _fffee
	}
	_cafcaf._aggcgb = _fgaa
	_gbdee, _fffee := _cafcaf.IsEncrypted()
	if _fffee != nil {
		return nil, _fffee
	}
	if !_gbdee {
		_fffee = _cafcaf.loadStructure()
		if _fffee != nil {
			return nil, _fffee
		}
	} else if _cagded {
		_dcgaa, _fgbb := _cafcaf.Decrypt([]byte(_geefc.Password))
		if _fgbb != nil {
			return nil, _fgbb
		}
		if !_dcgaa {
			return nil, _ceg.New("\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0064\u0065c\u0072\u0079\u0070\u0074\u0020\u0070\u0061\u0073\u0073w\u006f\u0072\u0064\u0020p\u0072\u006f\u0074\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u002d\u0020\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0073\u0070\u0065\u0063\u0069\u0066y\u0020\u0070\u0061s\u0073\u0020\u0074\u006f\u0020\u0044\u0065\u0063\u0072\u0079\u0070\u0074")
		}
	}
	_cafcaf._egaba = make(map[*PdfReader]*PdfReader)
	_cafcaf._fgfbe = make([]*PdfReader, _fgaa.GetRevisionNumber())
	return _cafcaf, nil
}
func _ecg(_fbfce *_cde.PdfObjectDictionary) *VRI {
	_cgea, _ := _cde.GetString(_fbfce.Get("\u0054\u0055"))
	_ddfba, _ := _cde.GetString(_fbfce.Get("\u0054\u0053"))
	return &VRI{Cert: _ebbec(_fbfce.Get("\u0043\u0065\u0072\u0074")), OCSP: _ebbec(_fbfce.Get("\u004f\u0043\u0053\u0050")), CRL: _ebbec(_fbfce.Get("\u0043\u0052\u004c")), TU: _cgea, TS: _ddfba}
}

// NewPdfAnnotationLink returns a new link annotation.
func NewPdfAnnotationLink() *PdfAnnotationLink {
	_ccc := NewPdfAnnotation()
	_dfg := &PdfAnnotationLink{}
	_dfg.PdfAnnotation = _ccc
	_ccc.SetContext(_dfg)
	return _dfg
}
func (_abbcd *PdfWriter) adjustXRefAffectedVersion(_edeaf bool) {
	if _edeaf && _abbcd._cgdcc.Major == 1 && _abbcd._cgdcc.Minor < 5 {
		_abbcd._cgdcc.Minor = 5
	}
}

// FieldAppearanceGenerator generates appearance stream for a given field.
type FieldAppearanceGenerator interface {
	ContentStreamWrapper
	GenerateAppearanceDict(_acbge *PdfAcroForm, _badd *PdfField, _eccfc *PdfAnnotationWidget) (*_cde.PdfObjectDictionary, error)
}

// NewPdfAnnotationStamp returns a new stamp annotation.
func NewPdfAnnotationStamp() *PdfAnnotationStamp {
	_gbg := NewPdfAnnotation()
	_cdg := &PdfAnnotationStamp{}
	_cdg.PdfAnnotation = _gbg
	_cdg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gbg.SetContext(_cdg)
	return _cdg
}
func (_dfd *PdfReader) newPdfActionMovieFromDict(_agg *_cde.PdfObjectDictionary) (*PdfActionMovie, error) {
	return &PdfActionMovie{Annotation: _agg.Get("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e"), T: _agg.Get("\u0054"), Operation: _agg.Get("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn")}, nil
}

// ToPdfObject implements interface PdfModel.
func (_dcaec *PdfSignatureReference) ToPdfObject() _cde.PdfObject {
	_edga := _cde.MakeDict()
	_edga.SetIfNotNil("\u0054\u0079\u0070\u0065", _dcaec.Type)
	_edga.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064", _dcaec.TransformMethod)
	_edga.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073", _dcaec.TransformParams)
	_edga.SetIfNotNil("\u0044\u0061\u0074\u0061", _dcaec.Data)
	_edga.SetIfNotNil("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064", _dcaec.DigestMethod)
	return _edga
}

// NewPdfFilespecFromObj creates and returns a new PdfFilespec object.
func NewPdfFilespecFromObj(obj _cde.PdfObject) (*PdfFilespec, error) {
	_egbef := &PdfFilespec{}
	var _fgdc *_cde.PdfObjectDictionary
	if _geag, _baef := _cde.GetIndirect(obj); _baef {
		_egbef._bgac = _geag
		_eaad, _cebc := _cde.GetDict(_geag.PdfObject)
		if !_cebc {
			_ad.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _cde.ErrTypeError
		}
		_fgdc = _eaad
	} else if _decfe, _fefc := _cde.GetDict(obj); _fefc {
		_egbef._bgac = _decfe
		_fgdc = _decfe
	} else {
		_ad.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", obj)
		return nil, _cde.ErrTypeError
	}
	if _fgdc == nil {
		_ad.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _ceg.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _dcgc := _fgdc.Get("\u0054\u0079\u0070\u0065"); _dcgc != nil {
		_cgfe, _abec := _dcgc.(*_cde.PdfObjectName)
		if !_abec {
			_ad.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _dcgc)
		} else {
			if *_cgfe != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
				_ad.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070e\u0063\u0074\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063 \u0028\u0025\u0073\u0029", *_cgfe)
			}
		}
	}
	if _fbdg := _fgdc.Get("\u0046\u0053"); _fbdg != nil {
		_egbef.FS = _fbdg
	}
	if _ddba := _fgdc.Get("\u0046"); _ddba != nil {
		_egbef.F = _ddba
	}
	if _dgebb := _fgdc.Get("\u0055\u0046"); _dgebb != nil {
		_egbef.UF = _dgebb
	}
	if _cdae := _fgdc.Get("\u0044\u004f\u0053"); _cdae != nil {
		_egbef.DOS = _cdae
	}
	if _abac := _fgdc.Get("\u004d\u0061\u0063"); _abac != nil {
		_egbef.Mac = _abac
	}
	if _gcbdg := _fgdc.Get("\u0055\u006e\u0069\u0078"); _gcbdg != nil {
		_egbef.Unix = _gcbdg
	}
	if _ffaa := _fgdc.Get("\u0049\u0044"); _ffaa != nil {
		_egbef.ID = _ffaa
	}
	if _efaf := _fgdc.Get("\u0056"); _efaf != nil {
		_egbef.V = _efaf
	}
	if _daef := _fgdc.Get("\u0045\u0046"); _daef != nil {
		_egbef.EF = _daef
	}
	if _eadfd := _fgdc.Get("\u0052\u0046"); _eadfd != nil {
		_egbef.RF = _eadfd
	}
	if _adbfg := _fgdc.Get("\u0044\u0065\u0073\u0063"); _adbfg != nil {
		_egbef.Desc = _adbfg
	}
	if _daccg := _fgdc.Get("\u0043\u0049"); _daccg != nil {
		_egbef.CI = _daccg
	}
	return _egbef, nil
}
func (_fad *PdfReader) newPdfAnnotationPolygonFromDict(_faef *_cde.PdfObjectDictionary) (*PdfAnnotationPolygon, error) {
	_gae := PdfAnnotationPolygon{}
	_edad, _bebb := _fad.newPdfAnnotationMarkupFromDict(_faef)
	if _bebb != nil {
		return nil, _bebb
	}
	_gae.PdfAnnotationMarkup = _edad
	_gae.Vertices = _faef.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_gae.LE = _faef.Get("\u004c\u0045")
	_gae.BS = _faef.Get("\u0042\u0053")
	_gae.IC = _faef.Get("\u0049\u0043")
	_gae.BE = _faef.Get("\u0042\u0045")
	_gae.IT = _faef.Get("\u0049\u0054")
	_gae.Measure = _faef.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_gae, nil
}

// GetPdfInfo returns the PDF info dictionary.
func (_cccdc *PdfReader) GetPdfInfo() (*PdfInfo, error) {
	_ddcde, _bfbab := _cccdc.GetTrailer()
	if _bfbab != nil {
		return nil, _bfbab
	}
	var _egcc *_cde.PdfObjectDictionary
	_cecgb := _ddcde.Get("\u0049\u006e\u0066\u006f")
	switch _cbcgc := _cecgb.(type) {
	case *_cde.PdfObjectReference:
		_dcegcd := _cbcgc
		_cecgb, _bfbab = _cccdc.GetIndirectObjectByNumber(int(_dcegcd.ObjectNumber))
		_cecgb = _cde.TraceToDirectObject(_cecgb)
		if _bfbab != nil {
			return nil, _bfbab
		}
		_egcc, _ = _cecgb.(*_cde.PdfObjectDictionary)
	case *_cde.PdfObjectDictionary:
		_egcc = _cbcgc
	}
	if _egcc == nil {
		return nil, _ceg.New("I\u006e\u0066\u006f\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0070r\u0065s\u0065\u006e\u0074")
	}
	_dfeea, _bfbab := NewPdfInfoFromObject(_egcc)
	if _bfbab != nil {
		return nil, _bfbab
	}
	return _dfeea, nil
}

// NewPdfAnnotationText returns a new text annotation.
func NewPdfAnnotationText() *PdfAnnotationText {
	_fdb := NewPdfAnnotation()
	_gfg := &PdfAnnotationText{}
	_gfg.PdfAnnotation = _fdb
	_gfg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fdb.SetContext(_gfg)
	return _gfg
}
func (_acgbg *PdfReader) newPdfFieldFromIndirectObject(_acaa *_cde.PdfIndirectObject, _fffe *PdfField) (*PdfField, error) {
	if _eada, _afdb := _acgbg._bedfa.GetModelFromPrimitive(_acaa).(*PdfField); _afdb {
		return _eada, nil
	}
	_bgcg, _faa := _cde.GetDict(_acaa)
	if !_faa {
		return nil, _ee.Errorf("\u0050\u0064f\u0046\u0069\u0065\u006c\u0064 \u0069\u006e\u0064\u0069\u0072e\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_geffe := NewPdfField()
	_geffe._afgc = _acaa
	_geffe._afgc.PdfObject = _bgcg
	if _gedb, _bcee := _cde.GetName(_bgcg.Get("\u0046\u0054")); _bcee {
		_geffe.FT = _gedb
	}
	if _fffe != nil {
		_geffe.Parent = _fffe
	}
	_geffe.T, _ = _bgcg.Get("\u0054").(*_cde.PdfObjectString)
	_geffe.TU, _ = _bgcg.Get("\u0054\u0055").(*_cde.PdfObjectString)
	_geffe.TM, _ = _bgcg.Get("\u0054\u004d").(*_cde.PdfObjectString)
	_geffe.Ff, _ = _bgcg.Get("\u0046\u0066").(*_cde.PdfObjectInteger)
	_geffe.V = _bgcg.Get("\u0056")
	_geffe.DV = _bgcg.Get("\u0044\u0056")
	_geffe.AA = _bgcg.Get("\u0041\u0041")
	if DA := _bgcg.Get("\u0044\u0041"); DA != nil {
		DA, _ := _cde.GetString(DA)
		_geffe.VariableText = &VariableText{DA: DA}
		Q, _ := _bgcg.Get("\u0051").(*_cde.PdfObjectInteger)
		DS, _ := _bgcg.Get("\u0044\u0053").(*_cde.PdfObjectString)
		RV := _bgcg.Get("\u0052\u0056")
		_geffe.VariableText.Q = Q
		_geffe.VariableText.DS = DS
		_geffe.VariableText.RV = RV
	}
	_bdbe := _geffe.FT
	if _bdbe == nil && _fffe != nil {
		_bdbe = _fffe.FT
	}
	if _bdbe != nil {
		switch *_bdbe {
		case "\u0054\u0078":
			_adgd, _deae := _gbad(_bgcg)
			if _deae != nil {
				return nil, _deae
			}
			_adgd.PdfField = _geffe
			_geffe._ecfg = _adgd
		case "\u0043\u0068":
			_gcee, _ecfgf := _gccb(_bgcg)
			if _ecfgf != nil {
				return nil, _ecfgf
			}
			_gcee.PdfField = _geffe
			_geffe._ecfg = _gcee
		case "\u0042\u0074\u006e":
			_gceb, _gdef := _aebaa(_bgcg)
			if _gdef != nil {
				return nil, _gdef
			}
			_gceb.PdfField = _geffe
			_geffe._ecfg = _gceb
		case "\u0053\u0069\u0067":
			_dedd, _bbfa := _acgbg.newPdfFieldSignatureFromDict(_bgcg)
			if _bbfa != nil {
				return nil, _bbfa
			}
			_dedd.PdfField = _geffe
			_geffe._ecfg = _dedd
		default:
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", *_geffe.FT)
			return nil, _ceg.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079p\u0065")
		}
	}
	if _debc, _fdeeb := _cde.GetName(_bgcg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fdeeb {
		if *_debc == "\u0057\u0069\u0064\u0067\u0065\u0074" {
			_gcfab, _gebg := _acgbg.newPdfAnnotationFromIndirectObject(_acaa)
			if _gebg != nil {
				return nil, _gebg
			}
			_cccg, _gcccc := _gcfab.GetContext().(*PdfAnnotationWidget)
			if !_gcccc {
				return nil, _ceg.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0067e\u0074 \u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006fn")
			}
			_cccg._dbf = _geffe
			_cccg.Parent = _geffe._afgc
			_geffe.Annotations = append(_geffe.Annotations, _cccg)
			return _geffe, nil
		}
	}
	_cbdc := true
	if _bfde, _aebe := _cde.GetArray(_bgcg.Get("\u004b\u0069\u0064\u0073")); _aebe {
		_bada := make([]*_cde.PdfIndirectObject, 0, _bfde.Len())
		for _, _gbfb := range _bfde.Elements() {
			_gcda, _eadf := _cde.GetIndirect(_gbfb)
			if !_eadf {
				_edea, _fbee := _cde.GetStream(_gbfb)
				if _fbee && _edea.PdfObjectDictionary != nil {
					_bdfg, _acfb := _cde.GetNameVal(_edea.Get("\u0054\u0079\u0070\u0065"))
					if _acfb && _bdfg == "\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061" {
						_ad.Log.Debug("E\u0052RO\u0052:\u0020f\u006f\u0072\u006d\u0020\u0066i\u0065\u006c\u0064 \u004b\u0069\u0064\u0073\u0020a\u0072\u0072\u0061y\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0069n\u0076\u0061\u006cid \u004d\u0065\u0074\u0061\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e")
						continue
					}
				}
				return nil, _ceg.New("n\u006f\u0074\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0028\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065\u006cd\u0029")
			}
			_eeda, _caab := _cde.GetDict(_gcda)
			if !_caab {
				return nil, ErrTypeCheck
			}
			if _cbdc {
				_cbdc = !_aeeb(_eeda)
			}
			_bada = append(_bada, _gcda)
		}
		for _, _ggdg := range _bada {
			if _cbdc {
				_dffc, _gfgfb := _acgbg.newPdfAnnotationFromIndirectObject(_ggdg)
				if _gfgfb != nil {
					_ad.Log.Debug("\u0045r\u0072\u006fr\u0020\u006c\u006fa\u0064\u0069\u006e\u0067\u0020\u0077\u0069d\u0067\u0065\u0074\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0072 \u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _gfgfb)
					return nil, _gfgfb
				}
				_bbcdd, _abef := _dffc._bea.(*PdfAnnotationWidget)
				if !_abef {
					return nil, ErrTypeCheck
				}
				_bbcdd._dbf = _geffe
				_geffe.Annotations = append(_geffe.Annotations, _bbcdd)
			} else {
				_aedbgc, _abfdbe := _acgbg.newPdfFieldFromIndirectObject(_ggdg, _geffe)
				if _abfdbe != nil {
					_ad.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0068\u0069\u006c\u0064\u0020\u0066\u0069\u0065\u006c\u0064: \u0025\u0076", _abfdbe)
					return nil, _abfdbe
				}
				_geffe.Kids = append(_geffe.Kids, _aedbgc)
			}
		}
	}
	return _geffe, nil
}

// PdfColorDeviceCMYK is a CMYK32 color, where each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceCMYK [4]float64

func (_afad *PdfReader) newPdfAnnotationUnderlineFromDict(_ffb *_cde.PdfObjectDictionary) (*PdfAnnotationUnderline, error) {
	_aafd := PdfAnnotationUnderline{}
	_fggf, _gdda := _afad.newPdfAnnotationMarkupFromDict(_ffb)
	if _gdda != nil {
		return nil, _gdda
	}
	_aafd.PdfAnnotationMarkup = _fggf
	_aafd.QuadPoints = _ffb.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_aafd, nil
}

// NewPdfActionGoToR returns a new "go to remote" action.
func NewPdfActionGoToR() *PdfActionGoToR {
	_cg := NewPdfAction()
	_efb := &PdfActionGoToR{}
	_efb.PdfAction = _cg
	_cg.SetContext(_efb)
	return _efb
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// A, B and C components of the color. The values of the elements should be
// between 0 and 1.
func (_acbe *PdfColorspaceCalRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_aede := vals[0]
	if _aede < 0.0 || _aede > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _aede)
		return nil, ErrColorOutOfRange
	}
	_fge := vals[1]
	if _fge < 0.0 || _fge > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fge)
		return nil, ErrColorOutOfRange
	}
	_abfba := vals[2]
	if _abfba < 0.0 || _abfba > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _abfba)
		return nil, ErrColorOutOfRange
	}
	_fgacb := NewPdfColorCalRGB(_aede, _fge, _abfba)
	return _fgacb, nil
}

// NewPdfColorCalRGB returns a new CalRBG color.
func NewPdfColorCalRGB(a, b, c float64) *PdfColorCalRGB {
	_aeee := PdfColorCalRGB{a, b, c}
	return &_aeee
}

// AlphaMapFunc represents a alpha mapping function: byte -> byte. Can be used for
// thresholding the alpha channel, i.e. setting all alpha values below threshold to transparent.
type AlphaMapFunc func(_agfb byte) byte

// ToPdfObject sets the common field elements.
// Note: Call the more field context's ToPdfObject to set both the generic and
// non-generic information.
func (_adaf *PdfField) ToPdfObject() _cde.PdfObject {
	_gbcb := _adaf._afgc
	_gddfd := _gbcb.PdfObject.(*_cde.PdfObjectDictionary)
	_fcgf := _cde.MakeArray()
	for _, _gace := range _adaf.Kids {
		_fcgf.Append(_gace.ToPdfObject())
	}
	for _, _dabd := range _adaf.Annotations {
		if _dabd._bddg != _adaf._afgc {
			_fcgf.Append(_dabd.GetContext().ToPdfObject())
		}
	}
	if _adaf.Parent != nil {
		_gddfd.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _adaf.Parent.GetContainingPdfObject())
	}
	if _fcgf.Len() > 0 {
		_gddfd.Set("\u004b\u0069\u0064\u0073", _fcgf)
	}
	_gddfd.SetIfNotNil("\u0046\u0054", _adaf.FT)
	_gddfd.SetIfNotNil("\u0054", _adaf.T)
	_gddfd.SetIfNotNil("\u0054\u0055", _adaf.TU)
	_gddfd.SetIfNotNil("\u0054\u004d", _adaf.TM)
	_gddfd.SetIfNotNil("\u0046\u0066", _adaf.Ff)
	_gddfd.SetIfNotNil("\u0056", _adaf.V)
	_gddfd.SetIfNotNil("\u0044\u0056", _adaf.DV)
	_gddfd.SetIfNotNil("\u0041\u0041", _adaf.AA)
	if _adaf.VariableText != nil {
		_gddfd.SetIfNotNil("\u0044\u0041", _adaf.VariableText.DA)
		_gddfd.SetIfNotNil("\u0051", _adaf.VariableText.Q)
		_gddfd.SetIfNotNil("\u0044\u0053", _adaf.VariableText.DS)
		_gddfd.SetIfNotNil("\u0052\u0056", _adaf.VariableText.RV)
	}
	return _gbcb
}

// ToPdfObject implements interface PdfModel.
func (_gfaa *PdfActionGoToR) ToPdfObject() _cde.PdfObject {
	_gfaa.PdfAction.ToPdfObject()
	_dcf := _gfaa._bc
	_cba := _dcf.PdfObject.(*_cde.PdfObjectDictionary)
	_cba.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeGoToR)))
	if _gfaa.F != nil {
		_cba.Set("\u0046", _gfaa.F.ToPdfObject())
	}
	_cba.SetIfNotNil("\u0044", _gfaa.D)
	_cba.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _gfaa.NewWindow)
	return _dcf
}
func (_bddd *PdfColorspaceICCBased) String() string {
	return "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"
}

// SetContext sets the specific fielddata type, e.g. would be PdfFieldButton for a button field.
func (_addd *PdfField) SetContext(ctx PdfModel) { _addd._ecfg = ctx }

// PdfColorspaceDeviceGray represents a grayscale colorspace.
type PdfColorspaceDeviceGray struct{}

// PdfInfo holds document information that will overwrite
// document information global variables defined above.
type PdfInfo struct {
	Title        *_cde.PdfObjectString
	Author       *_cde.PdfObjectString
	Subject      *_cde.PdfObjectString
	Keywords     *_cde.PdfObjectString
	Creator      *_cde.PdfObjectString
	Producer     *_cde.PdfObjectString
	CreationDate *PdfDate
	ModifiedDate *PdfDate
	Trapped      *_cde.PdfObjectName
	_ccef        *_cde.PdfObjectDictionary
}

// ValidateSignatures validates digital signatures in the document.
func (_cdbe *PdfReader) ValidateSignatures(handlers []SignatureHandler) ([]SignatureValidationResult, error) {
	if _cdbe.AcroForm == nil {
		return nil, nil
	}
	if _cdbe.AcroForm.Fields == nil {
		return nil, nil
	}
	type sigFieldPair struct {
		_faddg  *PdfSignature
		_cgadc  *PdfField
		_egcbdc SignatureHandler
	}
	var _gcddg []*sigFieldPair
	for _, _ggafd := range _cdbe.AcroForm.AllFields() {
		if _ggafd.V == nil {
			continue
		}
		if _gbeac, _aeddd := _cde.GetDict(_ggafd.V); _aeddd {
			if _fcdfb, _dagde := _cde.GetNameVal(_gbeac.Get("\u0054\u0079\u0070\u0065")); _dagde && _fcdfb == "\u0053\u0069\u0067" {
				_fcgcc, _edbcf := _cde.GetIndirect(_ggafd.V)
				if !_edbcf {
					_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0020\u0069s\u0020\u006e\u0069\u006c")
					return nil, ErrTypeCheck
				}
				_abfad, _gbfda := _cdbe.newPdfSignatureFromIndirect(_fcgcc)
				if _gbfda != nil {
					return nil, _gbfda
				}
				var _bdfeg SignatureHandler
				for _, _eecee := range handlers {
					if _eecee.IsApplicable(_abfad) {
						_bdfeg = _eecee
						break
					}
				}
				_gcddg = append(_gcddg, &sigFieldPair{_faddg: _abfad, _cgadc: _ggafd, _egcbdc: _bdfeg})
			}
		}
	}
	var _cbgge []SignatureValidationResult
	for _, _eeaga := range _gcddg {
		_aegba := SignatureValidationResult{IsSigned: true, Fields: []*PdfField{_eeaga._cgadc}}
		if _eeaga._egcbdc == nil {
			_aegba.Errors = append(_aegba.Errors, "\u0068a\u006ed\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
			_cbgge = append(_cbgge, _aegba)
			continue
		}
		_ffbdg, _fadde := _eeaga._egcbdc.NewDigest(_eeaga._faddg)
		if _fadde != nil {
			_aegba.Errors = append(_aegba.Errors, "\u0064\u0069\u0067e\u0073\u0074\u0020\u0065\u0072\u0072\u006f\u0072", _fadde.Error())
			_cbgge = append(_cbgge, _aegba)
			continue
		}
		_gggee := _eeaga._faddg.ByteRange
		if _gggee == nil {
			_aegba.Errors = append(_aegba.Errors, "\u0042\u0079\u0074\u0065\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
			_cbgge = append(_cbgge, _aegba)
			continue
		}
		for _fcfe := 0; _fcfe < _gggee.Len(); _fcfe = _fcfe + 2 {
			_gffbggd, _ := _cde.GetNumberAsInt64(_gggee.Get(_fcfe))
			_abfe, _ := _cde.GetIntVal(_gggee.Get(_fcfe + 1))
			if _, _adedg := _cdbe._caecc.Seek(_gffbggd, _f.SeekStart); _adedg != nil {
				return nil, _adedg
			}
			_eaeff := make([]byte, _abfe)
			if _, _gcgac := _cdbe._caecc.Read(_eaeff); _gcgac != nil {
				return nil, _gcgac
			}
			_ffbdg.Write(_eaeff)
		}
		var _ageab SignatureValidationResult
		if _gdbae, _cffaf := _eeaga._egcbdc.(SignatureHandlerDocMDP); _cffaf {
			_ageab, _fadde = _gdbae.ValidateWithOpts(_eeaga._faddg, _ffbdg, SignatureHandlerDocMDPParams{Parser: _cdbe._aggcgb})
		} else {
			_ageab, _fadde = _eeaga._egcbdc.Validate(_eeaga._faddg, _ffbdg)
		}
		if _fadde != nil {
			_ad.Log.Debug("E\u0052\u0052\u004f\u0052: \u0025v\u0020\u0028\u0025\u0054\u0029 \u002d\u0020\u0073\u006b\u0069\u0070", _fadde, _eeaga._egcbdc)
			_ageab.Errors = append(_ageab.Errors, _fadde.Error())
		}
		_ageab.Name = _eeaga._faddg.Name.Decoded()
		_ageab.Reason = _eeaga._faddg.Reason.Decoded()
		if _eeaga._faddg.M != nil {
			_feaab, _badfa := NewPdfDate(_eeaga._faddg.M.String())
			if _badfa != nil {
				_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _badfa)
				_ageab.Errors = append(_ageab.Errors, _badfa.Error())
				continue
			}
			_ageab.Date = _feaab
		}
		_ageab.ContactInfo = _eeaga._faddg.ContactInfo.Decoded()
		_ageab.Location = _eeaga._faddg.Location.Decoded()
		_ageab.Fields = _aegba.Fields
		_cbgge = append(_cbgge, _ageab)
	}
	return _cbgge, nil
}
func (_cecc fontCommon) fontFlags() int {
	if _cecc._fagf == nil {
		return 0
	}
	return _cecc._fagf._gbffb
}

// AddCustomInfo adds a custom info into document info dictionary.
func (_dbae *PdfInfo) AddCustomInfo(name string, value string) error {
	if _dbae._ccef == nil {
		_dbae._ccef = _cde.MakeDict()
	}
	if _, _fgbde := _fdga[name]; _fgbde {
		return _ee.Errorf("\u0063\u0061\u006e\u006e\u006ft\u0020\u0075\u0073\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072\u0064 \u0069\u006e\u0066\u006f\u0020\u006b\u0065\u0079\u0020\u0025\u0073\u0020\u0061\u0073\u0020\u0063\u0075\u0073\u0074\u006f\u006d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u006b\u0065y", name)
	}
	_dbae._ccef.SetIfNotNil(*_cde.MakeName(name), _cde.MakeString(value))
	return nil
}
func _eefcg(_dcgcg _cde.PdfObject) (*PdfFunctionType3, error) {
	_egcec := &PdfFunctionType3{}
	var _agbdd *_cde.PdfObjectDictionary
	if _fgcfg, _aege := _dcgcg.(*_cde.PdfIndirectObject); _aege {
		_egcg, _egdeg := _fgcfg.PdfObject.(*_cde.PdfObjectDictionary)
		if !_egdeg {
			return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_egcec._fbbgd = _fgcfg
		_agbdd = _egcg
	} else if _ceceeb, _fafg := _dcgcg.(*_cde.PdfObjectDictionary); _fafg {
		_agbdd = _ceceeb
	} else {
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_aegec, _cccge := _cde.TraceToDirectObject(_agbdd.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_cde.PdfObjectArray)
	if !_cccge {
		_ad.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _aegec.Len() != 2 {
		_ad.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _ceg.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_dfba, _abdea := _aegec.ToFloat64Array()
	if _abdea != nil {
		return nil, _abdea
	}
	_egcec.Domain = _dfba
	_aegec, _cccge = _cde.TraceToDirectObject(_agbdd.Get("\u0052\u0061\u006eg\u0065")).(*_cde.PdfObjectArray)
	if _cccge {
		if _aegec.Len() < 0 || _aegec.Len()%2 != 0 {
			return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_ecbcf, _cgbbe := _aegec.ToFloat64Array()
		if _cgbbe != nil {
			return nil, _cgbbe
		}
		_egcec.Range = _ecbcf
	}
	_aegec, _cccge = _cde.TraceToDirectObject(_agbdd.Get("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es")).(*_cde.PdfObjectArray)
	if !_cccge {
		_ad.Log.Error("\u0046\u0075\u006ect\u0069\u006f\u006e\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_egcec.Functions = []PdfFunction{}
	for _, _dfeaf := range _aegec.Elements() {
		_ffec, _gbbag := _cfdbb(_dfeaf)
		if _gbbag != nil {
			return nil, _gbbag
		}
		_egcec.Functions = append(_egcec.Functions, _ffec)
	}
	_aegec, _cccge = _cde.TraceToDirectObject(_agbdd.Get("\u0042\u006f\u0075\u006e\u0064\u0073")).(*_cde.PdfObjectArray)
	if !_cccge {
		_ad.Log.Error("B\u006fu\u006e\u0064\u0073\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_ccceb, _abdea := _aegec.ToFloat64Array()
	if _abdea != nil {
		return nil, _abdea
	}
	_egcec.Bounds = _ccceb
	if len(_egcec.Bounds) != len(_egcec.Functions)-1 {
		_ad.Log.Error("B\u006f\u0075\u006e\u0064\u0073\u0020\u0028\u0025\u0064)\u0020\u0061\u006e\u0064\u0020\u006e\u0075m \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029 n\u006f\u0074 \u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067", len(_egcec.Bounds), len(_egcec.Functions))
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_aegec, _cccge = _cde.TraceToDirectObject(_agbdd.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_cde.PdfObjectArray)
	if !_cccge {
		_ad.Log.Error("E\u006ec\u006f\u0064\u0065\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_dcca, _abdea := _aegec.ToFloat64Array()
	if _abdea != nil {
		return nil, _abdea
	}
	_egcec.Encode = _dcca
	if len(_egcec.Encode) != 2*len(_egcec.Functions) {
		_ad.Log.Error("\u004c\u0065\u006e\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0028\u0025\u0064\u0029 \u0061\u006e\u0064\u0020\u006e\u0075\u006d\u0020\u0066\u0075\u006e\u0063\u0074i\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029\u0020\u006e\u006f\u0074 m\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u0075\u0070", len(_egcec.Encode), len(_egcec.Functions))
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return _egcec, nil
}

// ToPdfObject implements interface PdfModel.
func (_aeg *PdfActionGoTo) ToPdfObject() _cde.PdfObject {
	_aeg.PdfAction.ToPdfObject()
	_cacd := _aeg._bc
	_cfb := _cacd.PdfObject.(*_cde.PdfObjectDictionary)
	_cfb.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeGoTo)))
	_cfb.SetIfNotNil("\u0044", _aeg.D)
	return _cacd
}

// PdfShadingType7 is a Tensor-product patch mesh.
type PdfShadingType7 struct {
	*PdfShading
	BitsPerCoordinate *_cde.PdfObjectInteger
	BitsPerComponent  *_cde.PdfObjectInteger
	BitsPerFlag       *_cde.PdfObjectInteger
	Decode            *_cde.PdfObjectArray
	Function          []PdfFunction
}

func _dagcd(_fdabc *_cde.PdfObjectDictionary) (*PdfTilingPattern, error) {
	_dffa := &PdfTilingPattern{}
	_adcf := _fdabc.Get("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e")
	if _adcf == nil {
		_ad.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054\u0079\u0070\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_aefd, _gadab := _adcf.(*_cde.PdfObjectInteger)
	if !_gadab {
		_ad.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054y\u0070\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _adcf)
		return nil, _cde.ErrTypeError
	}
	_dffa.PaintType = _aefd
	_adcf = _fdabc.Get("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _adcf == nil {
		_ad.Log.Debug("\u0054i\u006ci\u006e\u0067\u0054\u0079\u0070e\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_eebb, _gadab := _adcf.(*_cde.PdfObjectInteger)
	if !_gadab {
		_ad.Log.Debug("\u0054\u0069\u006cin\u0067\u0054\u0079\u0070\u0065\u0020\u006e\u006f\u0074 \u0061n\u0020i\u006et\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _adcf)
		return nil, _cde.ErrTypeError
	}
	_dffa.TilingType = _eebb
	_adcf = _fdabc.Get("\u0042\u0042\u006f\u0078")
	if _adcf == nil {
		_ad.Log.Debug("\u0042\u0042\u006fx\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_adcf = _cde.TraceToDirectObject(_adcf)
	_dfcgg, _gadab := _adcf.(*_cde.PdfObjectArray)
	if !_gadab {
		_ad.Log.Debug("\u0042B\u006f\u0078 \u0073\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061y\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _adcf)
		return nil, _cde.ErrTypeError
	}
	_ggfg, _gggd := NewPdfRectangle(*_dfcgg)
	if _gggd != nil {
		_ad.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _gggd)
		return nil, _gggd
	}
	_dffa.BBox = _ggfg
	_adcf = _fdabc.Get("\u0058\u0053\u0074e\u0070")
	if _adcf == nil {
		_ad.Log.Debug("\u0058\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_gggbe, _gggd := _cde.GetNumberAsFloat(_adcf)
	if _gggd != nil {
		_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0058S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _gggbe)
		return nil, _gggd
	}
	_dffa.XStep = _cde.MakeFloat(_gggbe)
	_adcf = _fdabc.Get("\u0059\u0053\u0074e\u0070")
	if _adcf == nil {
		_ad.Log.Debug("\u0059\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_faaf, _gggd := _cde.GetNumberAsFloat(_adcf)
	if _gggd != nil {
		_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0059S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _faaf)
		return nil, _gggd
	}
	_dffa.YStep = _cde.MakeFloat(_faaf)
	_adcf = _fdabc.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _adcf == nil {
		_ad.Log.Debug("\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fdabc, _gadab = _cde.TraceToDirectObject(_adcf).(*_cde.PdfObjectDictionary)
	if !_gadab {
		return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _adcf)
	}
	_fabca, _gggd := NewPdfPageResourcesFromDict(_fdabc)
	if _gggd != nil {
		return nil, _gggd
	}
	_dffa.Resources = _fabca
	if _eccgg := _fdabc.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _eccgg != nil {
		_cgbbea, _dgaef := _eccgg.(*_cde.PdfObjectArray)
		if !_dgaef {
			_ad.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _eccgg)
			return nil, _cde.ErrTypeError
		}
		_dffa.Matrix = _cgbbea
	}
	return _dffa, nil
}

// AddExtension adds the specified extension to the Extensions dictionary.
// See section 7.1.2 "Extensions Dictionary" (pp. 108-109 PDF32000_2008).
func (_baebe *PdfWriter) AddExtension(extName, baseVersion string, extLevel int) {
	_ecefc, _cdfeg := _cde.GetDict(_baebe._fedbb.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073"))
	if !_cdfeg {
		_ecefc = _cde.MakeDict()
		_baebe._fedbb.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073", _ecefc)
	}
	_dfab, _cdfeg := _cde.GetDict(_ecefc.Get(_cde.PdfObjectName(extName)))
	if !_cdfeg {
		_dfab = _cde.MakeDict()
		_ecefc.Set(_cde.PdfObjectName(extName), _dfab)
	}
	if _eege, _ := _cde.GetNameVal(_dfab.Get("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e")); _eege != baseVersion {
		_dfab.Set("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e", _cde.MakeName(baseVersion))
	}
	if _eccfdb, _ := _cde.GetIntVal(_dfab.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c")); _eccfdb != extLevel {
		_dfab.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c", _cde.MakeInteger(int64(extLevel)))
	}
}

// Encoder returns the font's text encoder.
func (_bedfg pdfCIDFontType2) Encoder() _gc.TextEncoder { return _bedfg._egga }
func _bddb(_cdgg *_ff.ImageBase) (_dafb Image) {
	_dafb.Width = int64(_cdgg.Width)
	_dafb.Height = int64(_cdgg.Height)
	_dafb.BitsPerComponent = int64(_cdgg.BitsPerComponent)
	_dafb.ColorComponents = _cdgg.ColorComponents
	_dafb.Data = _cdgg.Data
	_dafb._aaafb = _cdgg.Decode
	_dafb._deegf = _cdgg.Alpha
	return _dafb
}
func (_cedaff *PdfWriter) mapObjectStreams(_fefbd bool) (map[_cde.PdfObject]bool, bool) {
	_ggfa := make(map[_cde.PdfObject]bool)
	for _, _fcecg := range _cedaff._egbccc {
		if _gabeb, _cfdfb := _fcecg.(*_cde.PdfObjectStreams); _cfdfb {
			_fefbd = true
			for _, _acca := range _gabeb.Elements() {
				_ggfa[_acca] = true
				if _gdeffb, _ebdbe := _acca.(*_cde.PdfIndirectObject); _ebdbe {
					_ggfa[_gdeffb.PdfObject] = true
				}
			}
		}
	}
	return _ggfa, _fefbd
}

// GetContext returns the context of the outline tree node, which is either a
// *PdfOutline or a *PdfOutlineItem. The method returns nil for uninitialized
// tree nodes.
func (_dfbcd *PdfOutlineTreeNode) GetContext() PdfModel {
	if _bbgf, _afeg := _dfbcd._fbeea.(*PdfOutline); _afeg {
		return _bbgf
	}
	if _gedaa, _cfbg := _dfbcd._fbeea.(*PdfOutlineItem); _cfbg {
		return _gedaa
	}
	_ad.Log.Debug("\u0045\u0052RO\u0052\u0020\u0049n\u0076\u0061\u006c\u0069d o\u0075tl\u0069\u006e\u0065\u0020\u0074\u0072\u0065e \u006e\u006f\u0064\u0065\u0020\u0069\u0074e\u006d")
	return nil
}

// M returns the value of the magenta component of the color.
func (_bcgc *PdfColorDeviceCMYK) M() float64 { return _bcgc[1] }

// ColorFromPdfObjects gets the color from a series of pdf objects (4 for cmyk).
func (_gfcg *PdfColorspaceDeviceCMYK) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 4 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cafc, _egcd := _cde.GetNumbersAsFloat(objects)
	if _egcd != nil {
		return nil, _egcd
	}
	return _gfcg.ColorFromFloats(_cafc)
}

// GetType returns the button field type which returns one of the following
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
func (_aagd *PdfFieldButton) GetType() ButtonType {
	_faecd := ButtonTypeCheckbox
	if _aagd.Ff != nil {
		if (uint32(*_aagd.Ff) & FieldFlagPushbutton.Mask()) > 0 {
			_faecd = ButtonTypePush
		} else if (uint32(*_aagd.Ff) & FieldFlagRadio.Mask()) > 0 {
			_faecd = ButtonTypeRadio
		}
	}
	return _faecd
}

// Encoder returns the font's text encoder.
func (_bgeec pdfFontType3) Encoder() _gc.TextEncoder { return _bgeec._bdcege }

// SetOCProperties sets the optional content properties.
func (_acfe *PdfWriter) SetOCProperties(ocProperties _cde.PdfObject) error {
	_fccae := _acfe._fedbb
	if ocProperties != nil {
		_ad.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u004f\u0043\u0020\u0050\u0072o\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002e\u002e\u002e")
		_fccae.Set("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", ocProperties)
		return _acfe.addObjects(ocProperties)
	}
	return nil
}
func (_caf *PdfReader) newPdfAnnotationFreeTextFromDict(_efac *_cde.PdfObjectDictionary) (*PdfAnnotationFreeText, error) {
	_cbcb := PdfAnnotationFreeText{}
	_defab, _gba := _caf.newPdfAnnotationMarkupFromDict(_efac)
	if _gba != nil {
		return nil, _gba
	}
	_cbcb.PdfAnnotationMarkup = _defab
	_cbcb.DA = _efac.Get("\u0044\u0041")
	_cbcb.Q = _efac.Get("\u0051")
	_cbcb.RC = _efac.Get("\u0052\u0043")
	_cbcb.DS = _efac.Get("\u0044\u0053")
	_cbcb.CL = _efac.Get("\u0043\u004c")
	_cbcb.IT = _efac.Get("\u0049\u0054")
	_cbcb.BE = _efac.Get("\u0042\u0045")
	_cbcb.RD = _efac.Get("\u0052\u0044")
	_cbcb.BS = _efac.Get("\u0042\u0053")
	_cbcb.LE = _efac.Get("\u004c\u0045")
	return &_cbcb, nil
}

// SetAlpha sets the alpha layer for the image.
func (_eccfb *Image) SetAlpha(alpha []byte) { _eccfb._deegf = alpha }
func (_gdgb *PdfReader) newPdfAnnotationTextFromDict(_defa *_cde.PdfObjectDictionary) (*PdfAnnotationText, error) {
	_age := PdfAnnotationText{}
	_bbg, _bdff := _gdgb.newPdfAnnotationMarkupFromDict(_defa)
	if _bdff != nil {
		return nil, _bdff
	}
	_age.PdfAnnotationMarkup = _bbg
	_age.Open = _defa.Get("\u004f\u0070\u0065\u006e")
	_age.Name = _defa.Get("\u004e\u0061\u006d\u0065")
	_age.State = _defa.Get("\u0053\u0074\u0061t\u0065")
	_age.StateModel = _defa.Get("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c")
	return &_age, nil
}
func (_fbabga *PdfWriter) copyObject(_fbfced _cde.PdfObject, _eggaa map[_cde.PdfObject]_cde.PdfObject, _fgcbb map[_cde.PdfObject]struct{}, _aedff bool) _cde.PdfObject {
	_bgbgf := !_fbabga._aabfe && _fgcbb != nil
	if _dfda, _abfga := _eggaa[_fbfced]; _abfga {
		if _bgbgf && !_aedff {
			delete(_fgcbb, _fbfced)
		}
		return _dfda
	}
	if _fbfced == nil {
		_bdfefb := _cde.MakeNull()
		return _bdfefb
	}
	_eeaag := _fbfced
	switch _eeeeg := _fbfced.(type) {
	case *_cde.PdfObjectArray:
		_efbbb := _cde.MakeArray()
		_eeaag = _efbbb
		_eggaa[_fbfced] = _eeaag
		for _, _gfgd := range _eeeeg.Elements() {
			_efbbb.Append(_fbabga.copyObject(_gfgd, _eggaa, _fgcbb, _aedff))
		}
	case *_cde.PdfObjectStreams:
		_efagdg := &_cde.PdfObjectStreams{PdfObjectReference: _eeeeg.PdfObjectReference}
		_eeaag = _efagdg
		_eggaa[_fbfced] = _eeaag
		for _, _dbaea := range _eeeeg.Elements() {
			_efagdg.Append(_fbabga.copyObject(_dbaea, _eggaa, _fgcbb, _aedff))
		}
	case *_cde.PdfObjectStream:
		_abagb := &_cde.PdfObjectStream{Stream: _eeeeg.Stream, PdfObjectReference: _eeeeg.PdfObjectReference}
		_eeaag = _abagb
		_eggaa[_fbfced] = _eeaag
		_abagb.PdfObjectDictionary = _fbabga.copyObject(_eeeeg.PdfObjectDictionary, _eggaa, _fgcbb, _aedff).(*_cde.PdfObjectDictionary)
	case *_cde.PdfObjectDictionary:
		var _bceba bool
		if _bgbgf && !_aedff {
			if _dcbdb, _ := _cde.GetNameVal(_eeeeg.Get("\u0054\u0079\u0070\u0065")); _dcbdb == "\u0050\u0061\u0067\u0065" {
				_, _gdfcb := _fbabga._fbeeb[_eeeeg]
				_aedff = !_gdfcb
				_bceba = _aedff
			}
		}
		_afacf := _cde.MakeDict()
		_eeaag = _afacf
		_eggaa[_fbfced] = _eeaag
		for _, _eegfe := range _eeeeg.Keys() {
			_afacf.Set(_eegfe, _fbabga.copyObject(_eeeeg.Get(_eegfe), _eggaa, _fgcbb, _aedff))
		}
		if _bceba {
			_eeaag = _cde.MakeNull()
			_aedff = false
		}
	case *_cde.PdfIndirectObject:
		_gdbad := &_cde.PdfIndirectObject{PdfObjectReference: _eeeeg.PdfObjectReference}
		_eeaag = _gdbad
		_eggaa[_fbfced] = _eeaag
		_gdbad.PdfObject = _fbabga.copyObject(_eeeeg.PdfObject, _eggaa, _fgcbb, _aedff)
	case *_cde.PdfObjectString:
		_gfbaa := *_eeeeg
		_eeaag = &_gfbaa
		_eggaa[_fbfced] = _eeaag
	case *_cde.PdfObjectName:
		_ggaad := *_eeeeg
		_eeaag = &_ggaad
		_eggaa[_fbfced] = _eeaag
	case *_cde.PdfObjectNull:
		_eeaag = _cde.MakeNull()
		_eggaa[_fbfced] = _eeaag
	case *_cde.PdfObjectInteger:
		_afebb := *_eeeeg
		_eeaag = &_afebb
		_eggaa[_fbfced] = _eeaag
	case *_cde.PdfObjectReference:
		_ecbde := *_eeeeg
		_eeaag = &_ecbde
		_eggaa[_fbfced] = _eeaag
	case *_cde.PdfObjectFloat:
		_dfebf := *_eeeeg
		_eeaag = &_dfebf
		_eggaa[_fbfced] = _eeaag
	case *_cde.PdfObjectBool:
		_bgfdb := *_eeeeg
		_eeaag = &_bgfdb
		_eggaa[_fbfced] = _eeaag
	case *pdfSignDictionary:
		_acabd := &pdfSignDictionary{PdfObjectDictionary: _cde.MakeDict(), _addcc: _eeeeg._addcc, _afafa: _eeeeg._afafa}
		_eeaag = _acabd
		_eggaa[_fbfced] = _eeaag
		for _, _afaec := range _eeeeg.Keys() {
			_acabd.Set(_afaec, _fbabga.copyObject(_eeeeg.Get(_afaec), _eggaa, _fgcbb, _aedff))
		}
	default:
		_ad.Log.Info("\u0054\u004f\u0044\u004f\u0028\u0061\u0035\u0069\u0029\u003a\u0020\u0069\u006dp\u006c\u0065\u006d\u0065\u006e\u0074 \u0063\u006f\u0070\u0079\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0066\u006fr\u0020\u0025\u002b\u0076", _fbfced)
	}
	if _bgbgf && _aedff {
		_fgcbb[_fbfced] = struct{}{}
	}
	return _eeaag
}
func (_gddg *PdfAppender) addNewObject(_eegf _cde.PdfObject) {
	if _, _ggbd := _gddg._aadd[_eegf]; !_ggbd {
		_gddg._bfec = append(_gddg._bfec, _eegf)
		_gddg._aadd[_eegf] = struct{}{}
	}
}

// Permissions specify a permissions dictionary (PDF 1.5).
// (Section 12.8.4, Table 258 - Entries in a permissions dictionary p. 477 in PDF32000_2008).
type Permissions struct {
	DocMDP *PdfSignature
	_dafaf *_cde.PdfObjectDictionary
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

// GetXObjectByName returns the XObject with the specified keyName and the object type.
func (_fcddb *PdfPageResources) GetXObjectByName(keyName _cde.PdfObjectName) (*_cde.PdfObjectStream, XObjectType) {
	if _fcddb.XObject == nil {
		return nil, XObjectTypeUndefined
	}
	_gbccb, _bcfd := _cde.TraceToDirectObject(_fcddb.XObject).(*_cde.PdfObjectDictionary)
	if !_bcfd {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cde.TraceToDirectObject(_fcddb.XObject))
		return nil, XObjectTypeUndefined
	}
	if _eece := _gbccb.Get(keyName); _eece != nil {
		_aggac, _gbafb := _cde.GetStream(_eece)
		if !_gbafb {
			_ad.Log.Debug("X\u004f\u0062\u006a\u0065\u0063\u0074 \u006e\u006f\u0074\u0020\u0070\u006fi\u006e\u0074\u0069\u006e\u0067\u0020\u0074o\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020%\u0054", _eece)
			return nil, XObjectTypeUndefined
		}
		_fdda := _aggac.PdfObjectDictionary
		_fgfbd, _gbafb := _cde.TraceToDirectObject(_fdda.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")).(*_cde.PdfObjectName)
		if !_gbafb {
			_ad.Log.Debug("\u0058\u004fbj\u0065\u0063\u0074 \u0053\u0075\u0062\u0074ype\u0020no\u0074\u0020\u0061\u0020\u004e\u0061\u006de,\u0020\u0064\u0069\u0063\u0074\u003a\u0020%\u0073", _fdda.String())
			return nil, XObjectTypeUndefined
		}
		if *_fgfbd == "\u0049\u006d\u0061g\u0065" {
			return _aggac, XObjectTypeImage
		} else if *_fgfbd == "\u0046\u006f\u0072\u006d" {
			return _aggac, XObjectTypeForm
		} else if *_fgfbd == "\u0050\u0053" {
			return _aggac, XObjectTypePS
		} else {
			_ad.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0053\u0075b\u0074\u0079\u0070\u0065\u0020\u006e\u006ft\u0020\u006b\u006e\u006f\u0077\u006e\u0020\u0028\u0025\u0073\u0029", *_fgfbd)
			return nil, XObjectTypeUndefined
		}
	} else {
		return nil, XObjectTypeUndefined
	}
}

// NewPdfColorspaceSpecialIndexed returns a new Indexed color.
func NewPdfColorspaceSpecialIndexed() *PdfColorspaceSpecialIndexed {
	return &PdfColorspaceSpecialIndexed{HiVal: 255}
}

// NewPdfAnnotationInk returns a new ink annotation.
func NewPdfAnnotationInk() *PdfAnnotationInk {
	_ecf := NewPdfAnnotation()
	_bfe := &PdfAnnotationInk{}
	_bfe.PdfAnnotation = _ecf
	_bfe.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ecf.SetContext(_bfe)
	return _bfe
}

// ToPdfObject returns the PDF representation of the function.
func (_eedb *PdfFunctionType4) ToPdfObject() _cde.PdfObject {
	_gggff := _eedb._eaded
	if _gggff == nil {
		_eedb._eaded = &_cde.PdfObjectStream{}
		_gggff = _eedb._eaded
	}
	_agdf := _cde.MakeDict()
	_agdf.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _cde.MakeInteger(4))
	_aeeeg := &_cde.PdfObjectArray{}
	for _, _aedgc := range _eedb.Domain {
		_aeeeg.Append(_cde.MakeFloat(_aedgc))
	}
	_agdf.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _aeeeg)
	_debca := &_cde.PdfObjectArray{}
	for _, _gedbf := range _eedb.Range {
		_debca.Append(_cde.MakeFloat(_gedbf))
	}
	_agdf.Set("\u0052\u0061\u006eg\u0065", _debca)
	if _eedb._bfbd == nil && _eedb.Program != nil {
		_eedb._bfbd = []byte(_eedb.Program.String())
	}
	_agdf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cde.MakeInteger(int64(len(_eedb._bfbd))))
	_gggff.Stream = _eedb._bfbd
	_gggff.PdfObjectDictionary = _agdf
	return _gggff
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
	_gbgda *_cde.PdfIndirectObject
	Certs  []*_cde.PdfObjectStream
	OCSPs  []*_cde.PdfObjectStream
	CRLs   []*_cde.PdfObjectStream
	VRI    map[string]*VRI
	_cddc  map[string]*_cde.PdfObjectStream
	_ecgd  map[string]*_cde.PdfObjectStream
	_fbbd  map[string]*_cde.PdfObjectStream
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_baee *PdfShadingType3) ToPdfObject() _cde.PdfObject {
	_baee.PdfShading.ToPdfObject()
	_cabaaf, _dbcba := _baee.getShadingDict()
	if _dbcba != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _baee.Coords != nil {
		_cabaaf.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _baee.Coords)
	}
	if _baee.Domain != nil {
		_cabaaf.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _baee.Domain)
	}
	if _baee.Function != nil {
		if len(_baee.Function) == 1 {
			_cabaaf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _baee.Function[0].ToPdfObject())
		} else {
			_gaebg := _cde.MakeArray()
			for _, _ddfd := range _baee.Function {
				_gaebg.Append(_ddfd.ToPdfObject())
			}
			_cabaaf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _gaebg)
		}
	}
	if _baee.Extend != nil {
		_cabaaf.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _baee.Extend)
	}
	return _baee._dffg
}

// PdfAnnotationInk represents Ink annotations.
// (Section 12.5.6.13).
type PdfAnnotationInk struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	InkList _cde.PdfObject
	BS      _cde.PdfObject
}

func (_cadaa *PdfWriter) writeTrailer(_dabdca int) {
	_cadaa.writeString("\u0078\u0072\u0065\u0066\u000d\u000a")
	for _feceb := 0; _feceb <= _dabdca; {
		for ; _feceb <= _dabdca; _feceb++ {
			_fcec, _ddcag := _cadaa._gfdac[_feceb]
			if _ddcag && (!_cadaa._aabfe || _cadaa._aabfe && (_fcec.Type == 1 && _fcec.Offset >= _cadaa._fgged || _fcec.Type == 0)) {
				break
			}
		}
		var _aadee int
		for _aadee = _feceb + 1; _aadee <= _dabdca; _aadee++ {
			_bedag, _feeb := _cadaa._gfdac[_aadee]
			if _feeb && (!_cadaa._aabfe || _cadaa._aabfe && (_bedag.Type == 1 && _bedag.Offset > _cadaa._fgged)) {
				continue
			}
			break
		}
		_ddfeac := _ee.Sprintf("\u0025d\u0020\u0025\u0064\u000d\u000a", _feceb, _aadee-_feceb)
		_cadaa.writeString(_ddfeac)
		for _ffcfa := _feceb; _ffcfa < _aadee; _ffcfa++ {
			_acbfa := _cadaa._gfdac[_ffcfa]
			switch _acbfa.Type {
			case 0:
				_ddfeac = _ee.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u0066\u000d\u000a", 0, 65535)
				_cadaa.writeString(_ddfeac)
			case 1:
				_ddfeac = _ee.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u006e\u000d\u000a", _acbfa.Offset, 0)
				_cadaa.writeString(_ddfeac)
			}
		}
		_feceb = _aadee + 1
	}
	_gccdc := _cde.MakeDict()
	_gccdc.Set("\u0049\u006e\u0066\u006f", _cadaa._fdgbc)
	_gccdc.Set("\u0052\u006f\u006f\u0074", _cadaa._eacge)
	_gccdc.Set("\u0053\u0069\u007a\u0065", _cde.MakeInteger(int64(_dabdca+1)))
	if _cadaa._aabfe && _cadaa._gbbda > 0 {
		_gccdc.Set("\u0050\u0072\u0065\u0076", _cde.MakeInteger(_cadaa._gbbda))
	}
	if _cadaa._ccgbe != nil {
		_gccdc.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _cadaa._bdbdb)
	}
	if _cadaa._bgacb == nil && _cadaa._gfegf != "" && _cadaa._dcace != "" {
		_cadaa._bgacb = _cde.MakeArray(_cde.MakeHexString(_cadaa._gfegf), _cde.MakeHexString(_cadaa._dcace))
	}
	if _cadaa._bgacb != nil {
		_gccdc.Set("\u0049\u0044", _cadaa._bgacb)
		_ad.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _cadaa._bgacb)
	}
	_cadaa.writeString("\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u000a")
	_cadaa.writeString(_gccdc.WriteString())
	_cadaa.writeString("\u000a")
}

// DecodeArray returns the component range values for the Indexed colorspace.
func (_gggb *PdfColorspaceSpecialIndexed) DecodeArray() []float64 {
	return []float64{0, float64(_gggb.HiVal)}
}

// ToPdfObject implements interface PdfModel.
func (_bbeebd *PdfTransformParamsDocMDP) ToPdfObject() _cde.PdfObject {
	_cbadb := _cde.MakeDict()
	_cbadb.SetIfNotNil("\u0054\u0079\u0070\u0065", _bbeebd.Type)
	_cbadb.SetIfNotNil("\u0056", _bbeebd.V)
	_cbadb.SetIfNotNil("\u0050", _bbeebd.P)
	return _cbadb
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_dcefd *PdfShadingType2) ToPdfObject() _cde.PdfObject {
	_dcefd.PdfShading.ToPdfObject()
	_cdgfc, _afbgc := _dcefd.getShadingDict()
	if _afbgc != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _cdgfc == nil {
		_ad.Log.Error("\u0053\u0068\u0061\u0064in\u0067\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		return nil
	}
	if _dcefd.Coords != nil {
		_cdgfc.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _dcefd.Coords)
	}
	if _dcefd.Domain != nil {
		_cdgfc.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _dcefd.Domain)
	}
	if _dcefd.Function != nil {
		if len(_dcefd.Function) == 1 {
			_cdgfc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dcefd.Function[0].ToPdfObject())
		} else {
			_acggcf := _cde.MakeArray()
			for _, _dgeaf := range _dcefd.Function {
				_acggcf.Append(_dgeaf.ToPdfObject())
			}
			_cdgfc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _acggcf)
		}
	}
	if _dcefd.Extend != nil {
		_cdgfc.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _dcefd.Extend)
	}
	return _dcefd._dffg
}

// ToPdfObject converts rectangle to a PDF object.
func (_dfcc *PdfRectangle) ToPdfObject() _cde.PdfObject {
	return _cde.MakeArray(_cde.MakeFloat(_dfcc.Llx), _cde.MakeFloat(_dfcc.Lly), _cde.MakeFloat(_dfcc.Urx), _cde.MakeFloat(_dfcc.Ury))
}

// PdfAnnotationTrapNet represents TrapNet annotations.
// (Section 12.5.6.21).
type PdfAnnotationTrapNet struct{ *PdfAnnotation }

func (_ddfa *PdfReader) buildPageList(_fabfbe *_cde.PdfIndirectObject, _eefcgc *_cde.PdfIndirectObject, _dbbge map[_cde.PdfObject]struct{}) error {
	if _fabfbe == nil {
		return nil
	}
	if _, _gcedf := _dbbge[_fabfbe]; _gcedf {
		_ad.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _fabfbe.ObjectNumber)
		return nil
	}
	_dbbge[_fabfbe] = struct{}{}
	_bdbbc, _dbddf := _fabfbe.PdfObject.(*_cde.PdfObjectDictionary)
	if !_dbddf {
		return _ceg.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_beff, _dbddf := (*_bdbbc).Get("\u0054\u0079\u0070\u0065").(*_cde.PdfObjectName)
	if !_dbddf {
		if _bdbbc.Get("\u004b\u0069\u0064\u0073") == nil {
			return _ceg.New("\u006e\u006f\u0064\u0065 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0054\u0079p\u0065 \u0028\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0029")
		}
		_ad.Log.Debug("ER\u0052\u004fR\u003a\u0020\u006e\u006f\u0064\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067\u0020\u0054\u0079\u0070\u0065\u002c\u0020\u0062\u0075\u0074\u0020\u0068\u0061\u0073\u0020\u004b\u0069\u0064\u0073\u002e\u0020\u0041\u0073\u0073u\u006di\u006e\u0067\u0020\u0050\u0061\u0067\u0065\u0073 \u006eo\u0064\u0065.")
		_beff = _cde.MakeName("\u0050\u0061\u0067e\u0073")
		_bdbbc.Set("\u0054\u0079\u0070\u0065", _beff)
	}
	_ad.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0050a\u0067\u0065\u004c\u0069\u0073\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0074y\u0070\u0065\u003a\u0020\u0025\u0073\u0020(\u0025\u002b\u0076\u0029", *_beff, _fabfbe)
	if *_beff == "\u0050\u0061\u0067\u0065" {
		_beaec, _cabge := _ddfa.newPdfPageFromDict(_bdbbc)
		if _cabge != nil {
			return _cabge
		}
		_beaec.setContainer(_fabfbe)
		if _eefcgc != nil {
			_bdbbc.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _eefcgc)
		}
		_ddfa._gbfgg = append(_ddfa._gbfgg, _fabfbe)
		_ddfa.PageList = append(_ddfa.PageList, _beaec)
		return nil
	}
	if *_beff != "\u0050\u0061\u0067e\u0073" {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0063\u006fnt\u0065n\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u0021\u0020\u0028\u0025\u0073\u0029", _beff)
		return _ceg.New("\u0074\u0061\u0062\u006c\u0065\u0020o\u0066\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073 \u006fb\u006a\u0065\u0063\u0074")
	}
	if _eefcgc != nil {
		_bdbbc.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _eefcgc)
	}
	if !_ddfa._cdgee {
		_defbf := _ddfa.traverseObjectData(_fabfbe)
		if _defbf != nil {
			return _defbf
		}
	}
	_gcbag, _edeg := _ddfa._aggcgb.Resolve(_bdbbc.Get("\u004b\u0069\u0064\u0073"))
	if _edeg != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u004b\u0069\u0064\u0073\u0020\u006fb\u006a\u0065\u0063\u0074")
		return _edeg
	}
	var _cded *_cde.PdfObjectArray
	_cded, _dbddf = _gcbag.(*_cde.PdfObjectArray)
	if !_dbddf {
		_dbfcgg, _bedfcf := _gcbag.(*_cde.PdfIndirectObject)
		if !_bedfcf {
			return _ceg.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_cded, _dbddf = _dbfcgg.PdfObject.(*_cde.PdfObjectArray)
		if !_dbddf {
			return _ceg.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u0069\u006ed\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
	}
	_ad.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _cded)
	for _dcedg, _eddgf := range _cded.Elements() {
		_adcge, _gaccc := _cde.GetIndirect(_eddgf)
		if !_gaccc {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 \u006f\u0062\u006a\u0065\u0063t\u0020\u002d \u0028\u0025\u0073\u0029", _adcge)
			return _ceg.New("\u0070a\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_cded.Set(_dcedg, _adcge)
		_edeg = _ddfa.buildPageList(_adcge, _fabfbe, _dbbge)
		if _edeg != nil {
			return _edeg
		}
	}
	return nil
}

// ToPdfObject convert PdfInfo to pdf object.
func (_efcb *PdfInfo) ToPdfObject() _cde.PdfObject {
	_fafd := _cde.MakeDict()
	_fafd.SetIfNotNil("\u0054\u0069\u0074l\u0065", _efcb.Title)
	_fafd.SetIfNotNil("\u0041\u0075\u0074\u0068\u006f\u0072", _efcb.Author)
	_fafd.SetIfNotNil("\u0053u\u0062\u006a\u0065\u0063\u0074", _efcb.Subject)
	_fafd.SetIfNotNil("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _efcb.Keywords)
	_fafd.SetIfNotNil("\u0043r\u0065\u0061\u0074\u006f\u0072", _efcb.Creator)
	_fafd.SetIfNotNil("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _efcb.Producer)
	_fafd.SetIfNotNil("\u0054r\u0061\u0070\u0070\u0065\u0064", _efcb.Trapped)
	if _efcb.CreationDate != nil {
		_fafd.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _efcb.CreationDate.ToPdfObject())
	}
	if _efcb.ModifiedDate != nil {
		_fafd.SetIfNotNil("\u004do\u0064\u0044\u0061\u0074\u0065", _efcb.ModifiedDate.ToPdfObject())
	}
	for _, _gabbb := range _efcb._ccef.Keys() {
		_fafd.SetIfNotNil(_gabbb, _efcb._ccef.Get(_gabbb))
	}
	return _fafd
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_cafca *PdfColorspaceLab) ToPdfObject() _cde.PdfObject {
	_ebggcf := _cde.MakeArray()
	_ebggcf.Append(_cde.MakeName("\u004c\u0061\u0062"))
	_aaab := _cde.MakeDict()
	if _cafca.WhitePoint != nil {
		_dgce := _cde.MakeArray(_cde.MakeFloat(_cafca.WhitePoint[0]), _cde.MakeFloat(_cafca.WhitePoint[1]), _cde.MakeFloat(_cafca.WhitePoint[2]))
		_aaab.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _dgce)
	} else {
		_ad.Log.Error("\u004c\u0061\u0062: \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057h\u0069t\u0065P\u006fi\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	if _cafca.BlackPoint != nil {
		_aaba := _cde.MakeArray(_cde.MakeFloat(_cafca.BlackPoint[0]), _cde.MakeFloat(_cafca.BlackPoint[1]), _cde.MakeFloat(_cafca.BlackPoint[2]))
		_aaab.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _aaba)
	}
	if _cafca.Range != nil {
		_agf := _cde.MakeArray(_cde.MakeFloat(_cafca.Range[0]), _cde.MakeFloat(_cafca.Range[1]), _cde.MakeFloat(_cafca.Range[2]), _cde.MakeFloat(_cafca.Range[3]))
		_aaab.Set("\u0052\u0061\u006eg\u0065", _agf)
	}
	_ebggcf.Append(_aaab)
	if _cafca._gbbd != nil {
		_cafca._gbbd.PdfObject = _ebggcf
		return _cafca._gbbd
	}
	return _ebggcf
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a grayscale device.
func (_gdbe *PdfColorspaceDeviceGray) GetNumComponents() int { return 1 }

// GetColorspaces loads PdfPageResourcesColorspaces from `r.ColorSpace` and returns an error if there
// is a problem loading. Once loaded, the same object is returned on multiple calls.
func (_efdabe *PdfPageResources) GetColorspaces() (*PdfPageResourcesColorspaces, error) {
	if _efdabe._bfff != nil {
		return _efdabe._bfff, nil
	}
	if _efdabe.ColorSpace == nil {
		return nil, nil
	}
	_cbcd, _agabc := _dddcf(_efdabe.ColorSpace)
	if _agabc != nil {
		return nil, _agabc
	}
	_efdabe._bfff = _cbcd
	return _efdabe._bfff, nil
}

// ToPdfObject converts the font to a PDF representation.
func (_fbdd *pdfFontType0) ToPdfObject() _cde.PdfObject {
	if _fbdd._dcda == nil {
		_fbdd._dcda = &_cde.PdfIndirectObject{}
	}
	_edcbe := _fbdd.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0030")
	_fbdd._dcda.PdfObject = _edcbe
	if _fbdd.Encoding != nil {
		_edcbe.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _fbdd.Encoding)
	} else if _fbdd._cffef != nil {
		_edcbe.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _fbdd._cffef.ToPdfObject())
	}
	if _fbdd.DescendantFont != nil {
		_edcbe.Set("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073", _cde.MakeArray(_fbdd.DescendantFont.ToPdfObject()))
	}
	return _fbdd._dcda
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_cdaeb *PdfOutline) ToPdfObject() _cde.PdfObject {
	_feged := _cdaeb._dgcdc
	_afabc := _feged.PdfObject.(*_cde.PdfObjectDictionary)
	_afabc.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073"))
	if _cdaeb.First != nil {
		_afabc.Set("\u0046\u0069\u0072s\u0074", _cdaeb.First.ToPdfObject())
	}
	if _cdaeb.Last != nil {
		_afabc.Set("\u004c\u0061\u0073\u0074", _cdaeb.Last.GetContext().GetContainingPdfObject())
	}
	if _cdaeb.Parent != nil {
		_afabc.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _cdaeb.Parent.GetContext().GetContainingPdfObject())
	}
	if _cdaeb.Count != nil {
		_afabc.Set("\u0043\u006f\u0075n\u0074", _cde.MakeInteger(*_cdaeb.Count))
	}
	return _feged
}

// NewPdfFontFromTTFFile loads a TTF font file and returns a PdfFont type
// that can be used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTFFile.
func NewPdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_feec, _fdgga := _db.Open(filePath)
	if _fdgga != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020T\u0054F\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _fdgga)
		return nil, _fdgga
	}
	defer _feec.Close()
	return NewPdfFontFromTTF(_feec)
}

// Insert adds an outline item as a child of the current outline item,
// at the specified index.
func (_cdbad *OutlineItem) Insert(index uint, item *OutlineItem) {
	_aece := uint(len(_cdbad.Entries))
	if index > _aece {
		index = _aece
	}
	_cdbad.Entries = append(_cdbad.Entries[:index], append([]*OutlineItem{item}, _cdbad.Entries[index:]...)...)
}

// ToPdfOutlineItem returns a low level PdfOutlineItem object,
// based on the current instance.
func (_ebgd *OutlineItem) ToPdfOutlineItem() (*PdfOutlineItem, int64) {
	_efdf := NewPdfOutlineItem()
	_efdf.Title = _cde.MakeEncodedString(_ebgd.Title, true)
	_efdf.Dest = _ebgd.Dest.ToPdfObject()
	var _cfaef []*PdfOutlineItem
	var _fcffea int64
	var _fdcb *PdfOutlineItem
	for _, _gfegd := range _ebgd.Entries {
		_fcgeag, _ggfda := _gfegd.ToPdfOutlineItem()
		_fcgeag.Parent = &_efdf.PdfOutlineTreeNode
		if _fdcb != nil {
			_fdcb.Next = &_fcgeag.PdfOutlineTreeNode
			_fcgeag.Prev = &_fdcb.PdfOutlineTreeNode
		}
		_cfaef = append(_cfaef, _fcgeag)
		_fcffea += _ggfda
		_fdcb = _fcgeag
	}
	_bgce := len(_cfaef)
	_fcffea += int64(_bgce)
	if _bgce > 0 {
		_efdf.First = &_cfaef[0].PdfOutlineTreeNode
		_efdf.Last = &_cfaef[_bgce-1].PdfOutlineTreeNode
		_efdf.Count = &_fcffea
	}
	return _efdf, _fcffea
}

// ToPdfObject implements interface PdfModel.
func (_beag *PdfBorderStyle) ToPdfObject() _cde.PdfObject {
	_fddd := _cde.MakeDict()
	if _beag._dggc != nil {
		if _gdgc, _fdcg := _beag._dggc.(*_cde.PdfIndirectObject); _fdcg {
			_gdgc.PdfObject = _fddd
		}
	}
	_fddd.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0042\u006f\u0072\u0064\u0065\u0072"))
	if _beag.W != nil {
		_fddd.Set("\u0057", _cde.MakeFloat(*_beag.W))
	}
	if _beag.S != nil {
		_fddd.Set("\u0053", _cde.MakeName(_beag.S.GetPdfName()))
	}
	if _beag.D != nil {
		_fddd.Set("\u0044", _cde.MakeArrayFromIntegers(*_beag.D))
	}
	if _beag._dggc != nil {
		return _beag._dggc
	}
	return _fddd
}

// B returns the value of the blue component of the color.
func (_deeg *PdfColorDeviceRGB) B() float64 { return _deeg[2] }

// ToPdfObject implements interface PdfModel.
func (_eedaa *PdfFilespec) ToPdfObject() _cde.PdfObject {
	_eecdf := _eedaa.getDict()
	_eecdf.Clear()
	_eecdf.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063"))
	_eecdf.SetIfNotNil("\u0046\u0053", _eedaa.FS)
	_eecdf.SetIfNotNil("\u0046", _eedaa.F)
	_eecdf.SetIfNotNil("\u0055\u0046", _eedaa.UF)
	_eecdf.SetIfNotNil("\u0044\u004f\u0053", _eedaa.DOS)
	_eecdf.SetIfNotNil("\u004d\u0061\u0063", _eedaa.Mac)
	_eecdf.SetIfNotNil("\u0055\u006e\u0069\u0078", _eedaa.Unix)
	_eecdf.SetIfNotNil("\u0049\u0044", _eedaa.ID)
	_eecdf.SetIfNotNil("\u0056", _eedaa.V)
	_eecdf.SetIfNotNil("\u0045\u0046", _eedaa.EF)
	_eecdf.SetIfNotNil("\u0052\u0046", _eedaa.RF)
	_eecdf.SetIfNotNil("\u0044\u0065\u0073\u0063", _eedaa.Desc)
	_eecdf.SetIfNotNil("\u0043\u0049", _eedaa.CI)
	return _eedaa._bgac
}
func _eaddb(_fddc *_cde.PdfObjectDictionary, _begbc *fontCommon, _gcdg _gc.TextEncoder) (*pdfFontSimple, error) {
	_gbebe := _efge(_begbc)
	_gbebe._facga = _gcdg
	if _gcdg == nil {
		_cfca := _fddc.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
		if _cfca == nil {
			_cfca = _cde.MakeInteger(0)
		}
		_gbebe.FirstChar = _cfca
		_debce, _egce := _cde.GetIntVal(_cfca)
		if !_egce {
			_ad.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _cfca)
			return nil, _cde.ErrTypeError
		}
		_afadg := _gc.CharCode(_debce)
		_cfca = _fddc.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
		if _cfca == nil {
			_cfca = _cde.MakeInteger(255)
		}
		_gbebe.LastChar = _cfca
		_debce, _egce = _cde.GetIntVal(_cfca)
		if !_egce {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _cfca)
			return nil, _cde.ErrTypeError
		}
		_gdgec := _gc.CharCode(_debce)
		_gbebe._fecg = make(map[_gc.CharCode]float64)
		_cfca = _fddc.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _cfca != nil {
			_gbebe.Widths = _cfca
			_edfeg, _cfcce := _cde.GetArray(_cfca)
			if !_cfcce {
				_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _cfca)
				return nil, _cde.ErrTypeError
			}
			_dgebd, _dbfg := _edfeg.ToFloat64Array()
			if _dbfg != nil {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
				return nil, _dbfg
			}
			if len(_dgebd) != int(_gdgec-_afadg+1) {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _gdgec-_afadg+1, len(_dgebd))
				return nil, _cde.ErrRangeError
			}
			for _dfgd, _deccg := range _dgebd {
				_gbebe._fecg[_afadg+_gc.CharCode(_dfgd)] = _deccg
			}
		}
	}
	_gbebe.Encoding = _cde.TraceToDirectObject(_fddc.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	return _gbebe, nil
}

// ColorFromFloats returns a new PdfColorDevice based on the input slice of
// color components. The slice should contain four elements representing the
// cyan, magenta, yellow and key components of the color. The values of the
// elements should be between 0 and 1.
func (_bafb *PdfColorspaceDeviceCMYK) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 4 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_geedf := vals[0]
	if _geedf < 0.0 || _geedf > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _geedf)
		return nil, ErrColorOutOfRange
	}
	_beeg := vals[1]
	if _beeg < 0.0 || _beeg > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _beeg)
		return nil, ErrColorOutOfRange
	}
	_aggeb := vals[2]
	if _aggeb < 0.0 || _aggeb > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _aggeb)
		return nil, ErrColorOutOfRange
	}
	_ceef := vals[3]
	if _ceef < 0.0 || _ceef > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _ceef)
		return nil, ErrColorOutOfRange
	}
	_cecee := NewPdfColorDeviceCMYK(_geedf, _beeg, _aggeb, _ceef)
	return _cecee, nil
}
func (_dgbc *pdfFontSimple) getFontEncoding() (_gcaec string, _bafed map[_gc.CharCode]_gc.GlyphName, _gcbf error) {
	_gcaec = "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"
	if _adcbe, _gfabe := _dfbc[_dgbc._eeab]; _gfabe {
		_gcaec = _adcbe
	} else if _dgbc.fontFlags()&_edbca != 0 {
		for _becac, _feeg := range _dfbc {
			if _dac.Contains(_dgbc._eeab, _becac) {
				_gcaec = _feeg
				break
			}
		}
	}
	if _dgbc.Encoding == nil {
		return _gcaec, nil, nil
	}
	switch _ddfea := _dgbc.Encoding.(type) {
	case *_cde.PdfObjectName:
		return string(*_ddfea), nil, nil
	case *_cde.PdfObjectDictionary:
		_bcfc, _gaeb := _cde.GetName(_ddfea.Get("\u0042\u0061\u0073e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
		if _gaeb {
			_gcaec = _bcfc.String()
		}
		if _gbea := _ddfea.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"); _gbea != nil {
			_fedba, _beec := _cde.GetArray(_gbea)
			if !_beec {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042a\u0064\u0020\u0066on\u0074\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u003d\u0025\u002b\u0076\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073=\u0025\u0054", _ddfea, _ddfea.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"))
				return "", nil, _cde.ErrTypeError
			}
			_bafed, _gcbf = _gc.FromFontDifferences(_fedba)
		}
		return _gcaec, _bafed, _gcbf
	default:
		_ad.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0072\u0020\u0064\u0069\u0063t\u0020\u0028\u0025\u0054\u0029\u0020\u0025\u0073", _dgbc.Encoding, _dgbc.Encoding)
		return "", nil, _cde.ErrTypeError
	}
}

// PdfColorspaceCalRGB stores A, B, C components
type PdfColorspaceCalRGB struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      []float64
	Matrix     []float64
	_bfbf      *_cde.PdfObjectDictionary
	_efab      *_cde.PdfIndirectObject
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_afd *PdfColorspaceSpecialIndexed) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_edfc, _gaef := _cde.GetNumbersAsFloat(objects)
	if _gaef != nil {
		return nil, _gaef
	}
	return _afd.ColorFromFloats(_edfc)
}
func _bfbdc(_gagcf _cde.PdfObject) (*PdfShading, error) {
	_cabcb := &PdfShading{}
	var _fggef *_cde.PdfObjectDictionary
	if _caabc, _abbbc := _cde.GetIndirect(_gagcf); _abbbc {
		_cabcb._dffg = _caabc
		_ccgfa, _dbde := _caabc.PdfObject.(*_cde.PdfObjectDictionary)
		if !_dbde {
			_ad.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _cde.ErrTypeError
		}
		_fggef = _ccgfa
	} else if _cafeeg, _gbeg := _cde.GetStream(_gagcf); _gbeg {
		_cabcb._dffg = _cafeeg
		_fggef = _cafeeg.PdfObjectDictionary
	} else if _faag, _dbfgc := _cde.GetDict(_gagcf); _dbfgc {
		_cabcb._dffg = _faag
		_fggef = _faag
	} else {
		_ad.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", _gagcf)
		return nil, _cde.ErrTypeError
	}
	if _fggef == nil {
		_ad.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _ceg.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_gagcf = _fggef.Get("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _gagcf == nil {
		_ad.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073si\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_gagcf = _cde.TraceToDirectObject(_gagcf)
	_cebgd, _gegdc := _gagcf.(*_cde.PdfObjectInteger)
	if !_gegdc {
		_ad.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066o\u0072 \u0073h\u0061d\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _gagcf)
		return nil, _cde.ErrTypeError
	}
	if *_cebgd < 1 || *_cebgd > 7 {
		_ad.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u006e\u006ft\u0020\u0031\u002d\u0037\u0020(\u0067\u006ft\u0020\u0025\u0064\u0029", *_cebgd)
		return nil, _cde.ErrTypeError
	}
	_cabcb.ShadingType = _cebgd
	_gagcf = _fggef.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")
	if _gagcf == nil {
		_ad.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0020\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065\u0020e\u006e\u0074\u0072\u0079\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_geedc, _eeac := NewPdfColorspaceFromPdfObject(_gagcf)
	if _eeac != nil {
		_ad.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065: \u0025\u0076", _eeac)
		return nil, _eeac
	}
	_cabcb.ColorSpace = _geedc
	_gagcf = _fggef.Get("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064")
	if _gagcf != nil {
		_gagcf = _cde.TraceToDirectObject(_gagcf)
		_cfbafe, _fggcb := _gagcf.(*_cde.PdfObjectArray)
		if !_fggcb {
			_ad.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _gagcf)
			return nil, _cde.ErrTypeError
		}
		_cabcb.Background = _cfbafe
	}
	_gagcf = _fggef.Get("\u0042\u0042\u006f\u0078")
	if _gagcf != nil {
		_gagcf = _cde.TraceToDirectObject(_gagcf)
		_aaac, _aaadb := _gagcf.(*_cde.PdfObjectArray)
		if !_aaadb {
			_ad.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _gagcf)
			return nil, _cde.ErrTypeError
		}
		_gddfg, _cface := NewPdfRectangle(*_aaac)
		if _cface != nil {
			_ad.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _cface)
			return nil, _cface
		}
		_cabcb.BBox = _gddfg
	}
	_gagcf = _fggef.Get("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s")
	if _gagcf != nil {
		_gagcf = _cde.TraceToDirectObject(_gagcf)
		_ebgb, _fcad := _gagcf.(*_cde.PdfObjectBool)
		if !_fcad {
			_ad.Log.Debug("A\u006e\u0074\u0069\u0041\u006c\u0069\u0061\u0073\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020ty\u0070\u0065\u002c\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020bo\u006f\u006c \u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gagcf)
			return nil, _cde.ErrTypeError
		}
		_cabcb.AntiAlias = _ebgb
	}
	switch *_cebgd {
	case 1:
		_dcgcga, _cdacd := _gccbg(_fggef)
		if _cdacd != nil {
			return nil, _cdacd
		}
		_dcgcga.PdfShading = _cabcb
		_cabcb._dgfac = _dcgcga
		return _cabcb, nil
	case 2:
		_bcfbg, _bebeg := _eebfb(_fggef)
		if _bebeg != nil {
			return nil, _bebeg
		}
		_bcfbg.PdfShading = _cabcb
		_cabcb._dgfac = _bcfbg
		return _cabcb, nil
	case 3:
		_dcfcd, _ccgd := _eecae(_fggef)
		if _ccgd != nil {
			return nil, _ccgd
		}
		_dcfcd.PdfShading = _cabcb
		_cabcb._dgfac = _dcfcd
		return _cabcb, nil
	case 4:
		_gbbgf, _fbac := _cgcge(_fggef)
		if _fbac != nil {
			return nil, _fbac
		}
		_gbbgf.PdfShading = _cabcb
		_cabcb._dgfac = _gbbgf
		return _cabcb, nil
	case 5:
		_egcfc, _ffbee := _fggeg(_fggef)
		if _ffbee != nil {
			return nil, _ffbee
		}
		_egcfc.PdfShading = _cabcb
		_cabcb._dgfac = _egcfc
		return _cabcb, nil
	case 6:
		_fgbccd, _cbffb := _bcfa(_fggef)
		if _cbffb != nil {
			return nil, _cbffb
		}
		_fgbccd.PdfShading = _cabcb
		_cabcb._dgfac = _fgbccd
		return _cabcb, nil
	case 7:
		_dbfee, _bdae := _gedgbf(_fggef)
		if _bdae != nil {
			return nil, _bdae
		}
		_dbfee.PdfShading = _cabcb
		_cabcb._dgfac = _dbfee
		return _cabcb, nil
	}
	return nil, _ceg.New("u\u006ek\u006e\u006f\u0077\u006e\u0020\u0073\u0068\u0061d\u0069\u006e\u0067\u0020ty\u0070\u0065")
}
func _ceba(_gegeb _cde.PdfObject) (*PdfColorspaceDeviceNAttributes, error) {
	_dggbg := &PdfColorspaceDeviceNAttributes{}
	var _dggf *_cde.PdfObjectDictionary
	switch _ecbae := _gegeb.(type) {
	case *_cde.PdfIndirectObject:
		_dggbg._deaf = _ecbae
		var _aeded bool
		_dggf, _aeded = _ecbae.PdfObject.(*_cde.PdfObjectDictionary)
		if !_aeded {
			_ad.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
			return nil, _ceg.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	case *_cde.PdfObjectDictionary:
		_dggf = _ecbae
	case *_cde.PdfObjectReference:
		_egff := _ecbae.Resolve()
		return _ceba(_egff)
	default:
		_ad.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
		return nil, _ceg.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _agdgg := _dggf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _agdgg != nil {
		_aeab, _eefe := _cde.TraceToDirectObject(_agdgg).(*_cde.PdfObjectName)
		if !_eefe {
			_ad.Log.Error("\u0044\u0065vi\u0063\u0065\u004e \u0061\u0074\u0074\u0072ibu\u0074e \u0053\u0075\u0062\u0074\u0079\u0070\u0065 t\u0079\u0070\u0065\u0020\u0065\u0072\u0072o\u0072")
			return nil, _ceg.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_dggbg.Subtype = _aeab
	}
	if _bccg := _dggf.Get("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s"); _bccg != nil {
		_dggbg.Colorants = _bccg
	}
	if _fdgg := _dggf.Get("\u0050r\u006f\u0063\u0065\u0073\u0073"); _fdgg != nil {
		_dggbg.Process = _fdgg
	}
	if _gdgbe := _dggf.Get("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073"); _gdgbe != nil {
		_dggbg.MixingHints = _gdgbe
	}
	return _dggbg, nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// red, green and blue components of the color. The values of the elements
// should be between 0 and 1.
func (_gaaf *PdfColorspaceDeviceRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bdbfe := vals[0]
	if _bdbfe < 0.0 || _bdbfe > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bdbfe)
		return nil, ErrColorOutOfRange
	}
	_eafb := vals[1]
	if _eafb < 0.0 || _eafb > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bdbfe)
		return nil, ErrColorOutOfRange
	}
	_dcdg := vals[2]
	if _dcdg < 0.0 || _dcdg > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bdbfe)
		return nil, ErrColorOutOfRange
	}
	_dddbd := NewPdfColorDeviceRGB(_bdbfe, _eafb, _dcdg)
	return _dddbd, nil
}

// GetSamples converts the raw byte slice into samples which are stored in a uint32 bit array.
// Each sample is represented by BitsPerComponent consecutive bits in the raw data.
// NOTE: The method resamples the image byte data before returning the result and
// this could lead to high memory usage, especially on large images. It should
// be avoided, when possible. It is recommended to access the Data field of the
// image directly or use the ColorAt method to extract individual pixels.
func (_aafbc *Image) GetSamples() []uint32 {
	_gafdg := _cae.ResampleBytes(_aafbc.Data, int(_aafbc.BitsPerComponent))
	if _aafbc.BitsPerComponent < 8 {
		_gafdg = _aafbc.samplesTrimPadding(_gafdg)
	}
	_bcfcf := int(_aafbc.Width) * int(_aafbc.Height) * _aafbc.ColorComponents
	if len(_gafdg) < _bcfcf {
		_ad.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u0066\u0065\u0077\u0020\u0073a\u006d\u0070\u006c\u0065\u0073\u0020(\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064\u0029", len(_gafdg), _bcfcf)
		return _gafdg
	} else if len(_gafdg) > _bcfcf {
		_ad.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u006d\u0061\u006e\u0079\u0020s\u0061\u006d\u0070\u006c\u0065\u0073 \u0028\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078p\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064", len(_gafdg), _bcfcf)
		_gafdg = _gafdg[:_bcfcf]
	}
	return _gafdg
}
func _fgecb(_eecbc _cde.PdfObject) (map[_gc.CharCode]float64, error) {
	if _eecbc == nil {
		return nil, nil
	}
	_geae, _aegda := _cde.GetArray(_eecbc)
	if !_aegda {
		return nil, nil
	}
	_fbfee := map[_gc.CharCode]float64{}
	_bdece := _geae.Len()
	for _bgca := 0; _bgca < _bdece-1; _bgca++ {
		_gbfgb := _cde.TraceToDirectObject(_geae.Get(_bgca))
		_bgbag, _efea := _cde.GetIntVal(_gbfgb)
		if !_efea {
			return nil, _ee.Errorf("\u0042a\u0064\u0020\u0066\u006fn\u0074\u0020\u0057\u0020\u006fb\u006a0\u003a \u0069\u003d\u0025\u0064\u0020\u0025\u0023v", _bgca, _gbfgb)
		}
		_bgca++
		if _bgca > _bdece-1 {
			return nil, _ee.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _geae)
		}
		_ddgaa := _cde.TraceToDirectObject(_geae.Get(_bgca))
		switch _ddgaa.(type) {
		case *_cde.PdfObjectArray:
			_dgecd, _ := _cde.GetArray(_ddgaa)
			if _cadbf, _beega := _dgecd.ToFloat64Array(); _beega == nil {
				for _eaaa := 0; _eaaa < len(_cadbf); _eaaa++ {
					_fbfee[_gc.CharCode(_bgbag+_eaaa)] = _cadbf[_eaaa]
				}
			} else {
				return nil, _ee.Errorf("\u0042\u0061\u0064 \u0066\u006f\u006e\u0074 \u0057\u0020\u0061\u0072\u0072\u0061\u0079 \u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _bgca, _ddgaa)
			}
		case *_cde.PdfObjectInteger:
			_ceaec, _gfca := _cde.GetIntVal(_ddgaa)
			if !_gfca {
				return nil, _ee.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _bgca, _ddgaa)
			}
			_bgca++
			if _bgca > _bdece-1 {
				return nil, _ee.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _geae)
			}
			_eeaa := _geae.Get(_bgca)
			_gbeba, _cgcc := _cde.GetNumberAsFloat(_eeaa)
			if _cgcc != nil {
				return nil, _ee.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0032\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _bgca, _eeaa)
			}
			for _effcc := _bgbag; _effcc <= _ceaec; _effcc++ {
				_fbfee[_gc.CharCode(_effcc)] = _gbeba
			}
		default:
			return nil, _ee.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057 \u006f\u0062\u006a\u0031\u0020\u0074\u0079p\u0065\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _bgca, _ddgaa)
		}
	}
	return _fbfee, nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_fafb *PdfColorspaceDeviceGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bbgd := vals[0]
	if _bbgd < 0.0 || _bbgd > 1.0 {
		_ad.Log.Debug("\u0049\u006eco\u006d\u0070\u0061t\u0069\u0062\u0069\u006city\u003a R\u0061\u006e\u0067\u0065\u0020\u006f\u0075ts\u0069\u0064\u0065\u0020\u005b\u0030\u002c1\u005d")
	}
	if _bbgd < 0.0 {
		_bbgd = 0.0
	} else if _bbgd > 1.0 {
		_bbgd = 1.0
	}
	return NewPdfColorDeviceGray(_bbgd), nil
}

// NewPdfAnnotationFreeText returns a new free text annotation.
func NewPdfAnnotationFreeText() *PdfAnnotationFreeText {
	_gdb := NewPdfAnnotation()
	_dda := &PdfAnnotationFreeText{}
	_dda.PdfAnnotation = _gdb
	_dda.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gdb.SetContext(_dda)
	return _dda
}
func _cgebb(_ebbd _cde.PdfObject) (*_cde.PdfObjectDictionary, *fontCommon, error) {
	_bccb := &fontCommon{}
	if _becca, _adfc := _ebbd.(*_cde.PdfIndirectObject); _adfc {
		_bccb._dbadd = _becca.ObjectNumber
	}
	_aabdg, _bgeg := _cde.GetDict(_ebbd)
	if !_bgeg {
		_ad.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _ebbd)
		return nil, nil, ErrFontNotSupported
	}
	_befd, _bgeg := _cde.GetNameVal(_aabdg.Get("\u0054\u0079\u0070\u0065"))
	if !_bgeg {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	if _befd != "\u0046\u006f\u006e\u0074" {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u0020\u0049\u006e\u0063\u006f\u006d\u0070\u0061t\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u003d\u0025\u0071\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0025\u0071.", _befd, "\u0046\u006f\u006e\u0074")
		return nil, nil, _cde.ErrTypeError
	}
	_acgaf, _bgeg := _cde.GetNameVal(_aabdg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_bgeg {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020F\u006f\u006e\u0074 \u0049\u006e\u0063o\u006d\u0070a\u0074\u0069\u0062\u0069\u006c\u0069t\u0079. \u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	_bccb._dcbc = _acgaf
	_aeebg, _bgeg := _cde.GetNameVal(_aabdg.Get("\u004e\u0061\u006d\u0065"))
	if _bgeg {
		_bccb._fddeb = _aeebg
	}
	_fgab := _aabdg.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _fgab != nil {
		_bccb._dfae = _cde.TraceToDirectObject(_fgab)
		_ddgc, _gfae := _bebbe(_bccb._dfae, _bccb)
		if _gfae != nil {
			return _aabdg, _bccb, _gfae
		}
		_bccb._ggebg = _ddgc
	} else if _acgaf == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" || _acgaf == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_ggffd, _bfca := _fb.NewCIDSystemInfo(_aabdg.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
		if _bfca != nil {
			return _aabdg, _bccb, _bfca
		}
		_gefdg := _ee.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0055\u0043\u0053\u0032", _ggffd.Registry, _ggffd.Ordering)
		if _fb.IsPredefinedCMap(_gefdg) {
			_bccb._ggebg, _bfca = _fb.LoadPredefinedCMap(_gefdg)
			if _bfca != nil {
				_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _gefdg, _bfca)
			}
		}
	}
	_afca := _aabdg.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
	if _afca != nil {
		_bgea, _afffc := _ggad(_afca)
		if _afffc != nil {
			_ad.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0042\u0061\u0064\u0020\u0066\u006f\u006et\u0020d\u0065s\u0063r\u0069\u0070\u0074\u006f\u0072\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _afffc)
			return _aabdg, _bccb, _afffc
		}
		_bccb._fagf = _bgea
	}
	if _acgaf != "\u0054\u0079\u0070e\u0033" {
		_bdbde, _acabf := _cde.GetNameVal(_aabdg.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
		if !_acabf {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u006f\u006et\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069t\u0079\u002e\u0020\u0042\u0061se\u0046\u006f\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
			return _aabdg, _bccb, ErrRequiredAttributeMissing
		}
		_bccb._eeab = _bdbde
	}
	return _aabdg, _bccb, nil
}
func (_bgcc *PdfReader) lookupPageByObject(_gfba _cde.PdfObject) (*PdfPage, error) {
	return nil, _ceg.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// PdfColorPattern represents a pattern color.
type PdfColorPattern struct {
	Color       PdfColor
	PatternName _cde.PdfObjectName
}

// PdfColorDeviceRGB represents a color in DeviceRGB colorspace with R, G, B components, where component is
// defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceRGB [3]float64

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_adge *PdfColorspaceCalGray) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ddfb, _bdga := _cde.GetNumbersAsFloat(objects)
	if _bdga != nil {
		return nil, _bdga
	}
	return _adge.ColorFromFloats(_ddfb)
}

// ToPdfObject implements interface PdfModel.
func (_bcb *PdfActionMovie) ToPdfObject() _cde.PdfObject {
	_bcb.PdfAction.ToPdfObject()
	_cbe := _bcb._bc
	_ffc := _cbe.PdfObject.(*_cde.PdfObjectDictionary)
	_ffc.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeMovie)))
	_ffc.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e", _bcb.Annotation)
	_ffc.SetIfNotNil("\u0054", _bcb.T)
	_ffc.SetIfNotNil("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn", _bcb.Operation)
	return _cbe
}
func _bcfg(_aggcg _cde.PdfObject) (*PdfColorspaceLab, error) {
	_gbae := NewPdfColorspaceLab()
	if _cgggd, _dagd := _aggcg.(*_cde.PdfIndirectObject); _dagd {
		_gbae._gbbd = _cgggd
	}
	_aggcg = _cde.TraceToDirectObject(_aggcg)
	_adfac, _eagd := _aggcg.(*_cde.PdfObjectArray)
	if !_eagd {
		return nil, _ee.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _adfac.Len() != 2 {
		return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_aggcg = _cde.TraceToDirectObject(_adfac.Get(0))
	_eaaf, _eagd := _aggcg.(*_cde.PdfObjectName)
	if !_eagd {
		return nil, _ee.Errorf("\u006c\u0061\u0062\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if *_eaaf != "\u004c\u0061\u0062" {
		return nil, _ee.Errorf("n\u006ft\u0020\u0061\u0020\u004c\u0061\u0062\u0020\u0063o\u006c\u006f\u0072\u0073pa\u0063\u0065")
	}
	_aggcg = _cde.TraceToDirectObject(_adfac.Get(1))
	_gacd, _eagd := _aggcg.(*_cde.PdfObjectDictionary)
	if !_eagd {
		return nil, _ee.Errorf("c\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006ev\u0061l\u0069\u0064")
	}
	_aggcg = _gacd.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_aggcg = _cde.TraceToDirectObject(_aggcg)
	_ababfb, _eagd := _aggcg.(*_cde.PdfObjectArray)
	if !_eagd {
		return nil, _ee.Errorf("\u004c\u0061\u0062\u0020In\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069n\u0074")
	}
	if _ababfb.Len() != 3 {
		return nil, _ee.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
	}
	_gdge, _gegda := _ababfb.GetAsFloat64Slice()
	if _gegda != nil {
		return nil, _gegda
	}
	_gbae.WhitePoint = _gdge
	_aggcg = _gacd.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _aggcg != nil {
		_aggcg = _cde.TraceToDirectObject(_aggcg)
		_fbeg, _gbec := _aggcg.(*_cde.PdfObjectArray)
		if !_gbec {
			return nil, _ee.Errorf("\u004c\u0061\u0062: \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
		}
		if _fbeg.Len() != 3 {
			return nil, _ee.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
		}
		_bdcdd, _daad := _fbeg.GetAsFloat64Slice()
		if _daad != nil {
			return nil, _daad
		}
		_gbae.BlackPoint = _bdcdd
	}
	_aggcg = _gacd.Get("\u0052\u0061\u006eg\u0065")
	if _aggcg != nil {
		_aggcg = _cde.TraceToDirectObject(_aggcg)
		_dffbf, _aedg := _aggcg.(*_cde.PdfObjectArray)
		if !_aedg {
			_ad.Log.Error("\u0052\u0061n\u0067\u0065\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
			return nil, _ee.Errorf("\u004ca\u0062:\u0020\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if _dffbf.Len() != 4 {
			_ad.Log.Error("\u0052\u0061\u006e\u0067\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020e\u0072\u0072\u006f\u0072")
			return nil, _ee.Errorf("\u004c\u0061b\u003a\u0020\u0052a\u006e\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ecae, _aabbe := _dffbf.GetAsFloat64Slice()
		if _aabbe != nil {
			return nil, _aabbe
		}
		_gbae.Range = _ecae
	}
	return _gbae, nil
}

// SetEncoder sets the encoding for the underlying font.
// TODO(peterwilliams97): Change function signature to SetEncoder(encoder *textencoding.simpleEncoder).
// TODO(gunnsth): Makes sense if SetEncoder is removed from the interface fonts.Font as proposed in PR #260.
func (_aebg *pdfFontSimple) SetEncoder(encoder _gc.TextEncoder) { _aebg._efeaf = encoder }

var ImageHandling ImageHandler = DefaultImageHandler{}
var _ pdfFont = (*pdfCIDFontType2)(nil)

// DecodeArray returns the component range values for the Separation colorspace.
func (_cbfd *PdfColorspaceSpecialSeparation) DecodeArray() []float64 { return []float64{0, 1.0} }

// ToPdfObject implements interface PdfModel.
func (_gbfa *PdfAnnotationPopup) ToPdfObject() _cde.PdfObject {
	_gbfa.PdfAnnotation.ToPdfObject()
	_eaf := _gbfa._bddg
	_bbegc := _eaf.PdfObject.(*_cde.PdfObjectDictionary)
	_bbegc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0050\u006f\u0070u\u0070"))
	_bbegc.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _gbfa.Parent)
	_bbegc.SetIfNotNil("\u004f\u0070\u0065\u006e", _gbfa.Open)
	return _eaf
}

// PdfActionType represents an action type in PDF (section 12.6.4 p. 417).
type PdfActionType string

// Write writes the Appender output to io.Writer.
// It can only be called once and further invocations will result in an error.
func (_faee *PdfAppender) Write(w _f.Writer) error {
	if _faee._gcbc {
		return _ceg.New("\u0061\u0070\u0070\u0065\u006e\u0064\u0065\u0072\u0020\u0077\u0072\u0069\u0074e\u0020\u0063\u0061\u006e\u0020\u006fn\u006c\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0076\u006f\u006b\u0065\u0064 \u006f\u006e\u0063\u0065")
	}
	_aeba := NewPdfWriter()
	_daab, _bafe := _cde.GetDict(_aeba._acbgeg)
	if !_bafe {
		return _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_gege, _bafe := _daab.Get("\u004b\u0069\u0064\u0073").(*_cde.PdfObjectArray)
	if !_bafe {
		return _ceg.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_eggf, _bafe := _daab.Get("\u0043\u006f\u0075n\u0074").(*_cde.PdfObjectInteger)
	if !_bafe {
		return _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	_cfedb := _faee._cfad._aggcgb
	_edf := _cfedb.GetTrailer()
	if _edf == nil {
		return _ceg.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_fbaf, _bafe := _cde.GetIndirect(_edf.Get("\u0052\u006f\u006f\u0074"))
	if !_bafe {
		return _ceg.New("c\u0061\u0074\u0061\u006c\u006f\u0067 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064")
	}
	_gcaf, _bafe := _cde.GetDict(_fbaf)
	if !_bafe {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _fbaf, *_edf)
		return _ceg.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	for _, _fdcga := range _gcaf.Keys() {
		if _aeba._fedbb.Get(_fdcga) == nil {
			_eecd := _gcaf.Get(_fdcga)
			_aeba._fedbb.Set(_fdcga, _eecd)
		}
	}
	if _faee._cbeaa != nil {
		_aeba._fedbb.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _faee._cbeaa.ToPdfObject())
		_faee.updateObjectsDeep(_faee._cbeaa.ToPdfObject(), nil)
	}
	if _faee._ffac != nil {
		_faee.updateObjectsDeep(_faee._ffac.ToPdfObject(), nil)
		_aeba._fedbb.Set("\u0044\u0053\u0053", _faee._ffac.GetContainingPdfObject())
	}
	if _faee._bfbe != nil {
		_aeba._fedbb.Set("\u0050\u0065\u0072m\u0073", _faee._bfbe.ToPdfObject())
		_faee.updateObjectsDeep(_faee._bfbe.ToPdfObject(), nil)
	}
	if _aeba._cgdcc.Major < 2 {
		_aeba.AddExtension("\u0045\u0053\u0049\u0043", "\u0031\u002e\u0037", 5)
		_aeba.AddExtension("\u0041\u0044\u0042\u0045", "\u0031\u002e\u0037", 8)
	}
	if _dfea, _adbb := _cde.GetDict(_edf.Get("\u0049\u006e\u0066\u006f")); _adbb {
		if _gbgf, _cgae := _cde.GetDict(_aeba._fdgbc); _cgae {
			for _, _bde := range _dfea.Keys() {
				if _gbgf.Get(_bde) == nil {
					_gbgf.Set(_bde, _dfea.Get(_bde))
				}
			}
		}
	}
	if _faee._gdaa != nil {
		_aeba._fdgbc = _cde.MakeIndirectObject(_faee._gdaa.ToPdfObject())
	}
	_faee.addNewObject(_aeba._fdgbc)
	_faee.addNewObject(_aeba._eacge)
	_fgac := false
	if len(_faee._cfad.PageList) != len(_faee._gfb) {
		_fgac = true
	} else {
		for _dcc := range _faee._cfad.PageList {
			switch {
			case _faee._gfb[_dcc] == _faee._cfad.PageList[_dcc]:
			case _faee._gfb[_dcc] == _faee.Reader.PageList[_dcc]:
			default:
				_fgac = true
			}
			if _fgac {
				break
			}
		}
	}
	if _fgac {
		_faee.updateObjectsDeep(_aeba._acbgeg, nil)
	} else {
		_faee._deb[_aeba._acbgeg] = struct{}{}
	}
	_aeba._acbgeg.ObjectNumber = _faee.Reader._geadg.ObjectNumber
	_faee._adfa[_aeba._acbgeg] = _faee.Reader._geadg.ObjectNumber
	_beca := []_cde.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	for _, _adgf := range _faee._gfb {
		_ccga := _adgf.ToPdfObject()
		*_eggf = *_eggf + 1
		if _dcaf, _addb := _ccga.(*_cde.PdfIndirectObject); _addb && _dcaf.GetParser() == _faee._cfad._aggcgb {
			_gege.Append(&_dcaf.PdfObjectReference)
			continue
		}
		if _cged, _adab := _cde.GetDict(_ccga); _adab {
			_ccdca, _gfcf := _cged.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_cde.PdfIndirectObject)
			for _gfcf {
				_ad.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _ccdca)
				_gcag, _acgf := _ccdca.PdfObject.(*_cde.PdfObjectDictionary)
				if !_acgf {
					return _ceg.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
				}
				for _, _egc := range _beca {
					_ad.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _egc)
					if _cged.Get(_egc) != nil {
						_ad.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
						continue
					}
					if _gged := _gcag.Get(_egc); _gged != nil {
						_ad.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _egc)
						_cged.Set(_egc, _gged)
					}
				}
				_ccdca, _gfcf = _gcag.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_cde.PdfIndirectObject)
				_ad.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _gcag.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
			}
			_cged.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _aeba._acbgeg)
		}
		_faee.updateObjectsDeep(_ccga, nil)
		_gege.Append(_ccga)
	}
	if _, _bac := _faee._aac.Seek(0, _f.SeekStart); _bac != nil {
		return _bac
	}
	_bcbg := make(map[SignatureHandler]_f.Writer)
	_gaabbe := _cde.MakeArray()
	for _, _edfb := range _faee._bfec {
		if _dcabb, _fcge := _cde.GetIndirect(_edfb); _fcge {
			if _cfg, _feb := _dcabb.PdfObject.(*pdfSignDictionary); _feb {
				_egacb := *_cfg._addcc
				var _dggcd error
				_bcbg[_egacb], _dggcd = _egacb.NewDigest(_cfg._afafa)
				if _dggcd != nil {
					return _dggcd
				}
				_gaabbe.Append(_cde.MakeInteger(0xfffff), _cde.MakeInteger(0xfffff))
			}
		}
	}
	if _gaabbe.Len() > 0 {
		_gaabbe.Append(_cde.MakeInteger(0xfffff), _cde.MakeInteger(0xfffff))
	}
	for _, _acgd := range _faee._bfec {
		if _fgad, _fgdf := _cde.GetIndirect(_acgd); _fgdf {
			if _gaea, _cedc := _fgad.PdfObject.(*pdfSignDictionary); _cedc {
				_gaea.Set("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _gaabbe)
			}
		}
	}
	_ggff := len(_bcbg) > 0
	var _afada _f.Reader = _faee._aac
	if _ggff {
		_dfeg := make([]_f.Writer, 0, len(_bcbg))
		for _, _fega := range _bcbg {
			_dfeg = append(_dfeg, _fega)
		}
		_afada = _f.TeeReader(_faee._aac, _f.MultiWriter(_dfeg...))
	}
	_fcfd, _decd := _f.Copy(w, _afada)
	if _decd != nil {
		return _decd
	}
	if len(_faee._bfec) == 0 {
		return nil
	}
	_aeba._fgca = _fcfd
	_aeba.ObjNumOffset = _faee._afcea
	_aeba._aabfe = true
	_aeba._bdgaa = _faee._eccee
	_aeba._gbbda = _faee._aabb
	_aeba._fgged = _faee._efag
	_aeba._cgdcc = _faee._cfad.PdfVersion()
	_aeba._dgad = _faee._adfa
	_aeba._ccgbe = _faee._gdgf.GetCrypter()
	_aeba._bdbdb = _faee._gdgf.GetEncryptObj()
	_gcae := _faee._gdgf.GetXrefType()
	if _gcae != nil {
		_edbg := *_gcae == _cde.XrefTypeObjectStream
		_aeba._cagac = &_edbg
	}
	_aeba._bccde = map[_cde.PdfObject]struct{}{}
	_aeba._egbccc = []_cde.PdfObject{}
	for _, _gbaf := range _faee._bfec {
		if _, _bfgd := _faee._deb[_gbaf]; _bfgd {
			continue
		}
		_aeba.addObject(_gbaf)
	}
	_ecea := w
	if _ggff {
		_ecea = _ede.NewBuffer(nil)
	}
	if _faee._dgdff != "" && _aeba._ccgbe == nil {
		_aeba.Encrypt([]byte(_faee._dgdff), []byte(_faee._dgdff), _faee._fcga)
	}
	if _dfcf := _edf.Get("\u0049\u0044"); _dfcf != nil {
		if _cggd, _cgab := _cde.GetArray(_dfcf); _cgab {
			_aeba._bgacb = _cggd
		}
	}
	if _ccaf := _aeba.Write(_ecea); _ccaf != nil {
		return _ccaf
	}
	if _ggff {
		_cdfa := _ecea.(*_ede.Buffer).Bytes()
		_cfag := _cde.MakeArray()
		var _cacdc []*pdfSignDictionary
		var _bdbba int64
		for _, _gceg := range _aeba._egbccc {
			if _adcb, _ffd := _cde.GetIndirect(_gceg); _ffd {
				if _eae, _ecba := _adcb.PdfObject.(*pdfSignDictionary); _ecba {
					_cacdc = append(_cacdc, _eae)
					_dced := _eae._aafab + int64(_eae._gggca)
					_cfag.Append(_cde.MakeInteger(_bdbba), _cde.MakeInteger(_dced-_bdbba))
					_bdbba = _eae._aafab + int64(_eae._cgedg)
				}
			}
		}
		_cfag.Append(_cde.MakeInteger(_bdbba), _cde.MakeInteger(_fcfd+int64(len(_cdfa))-_bdbba))
		_geed := []byte(_cfag.WriteString())
		for _, _gdce := range _cacdc {
			_cacde := int(_gdce._aafab - _fcfd)
			for _fbca := _gdce._faba; _fbca < _gdce._edgb; _fbca++ {
				_cdfa[_cacde+_fbca] = ' '
			}
			_abaac := _cdfa[_cacde+_gdce._faba : _cacde+_gdce._edgb]
			copy(_abaac, _geed)
		}
		var _bcef int
		for _, _acc := range _cacdc {
			_bedf := int(_acc._aafab - _fcfd)
			_dddg := _cdfa[_bcef : _bedf+_acc._gggca]
			_gcgbf := *_acc._addcc
			_bcbg[_gcgbf].Write(_dddg)
			_bcef = _bedf + _acc._cgedg
		}
		for _, _bdec := range _cacdc {
			_bddgd := _cdfa[_bcef:]
			_cgdc := *_bdec._addcc
			_bcbg[_cgdc].Write(_bddgd)
		}
		for _, _ddgg := range _cacdc {
			_gbag := int(_ddgg._aafab - _fcfd)
			_abgca := *_ddgg._addcc
			_gcba := _bcbg[_abgca]
			if _deacd := _abgca.Sign(_ddgg._afafa, _gcba); _deacd != nil {
				return _deacd
			}
			_ddgg._afafa.ByteRange = _cfag
			_fcgd := []byte(_ddgg._afafa.Contents.WriteString())
			for _abgd := _ddgg._faba; _abgd < _ddgg._edgb; _abgd++ {
				_cdfa[_gbag+_abgd] = ' '
			}
			for _dbbd := _ddgg._gggca; _dbbd < _ddgg._cgedg; _dbbd++ {
				_cdfa[_gbag+_dbbd] = ' '
			}
			_cdeb := _cdfa[_gbag+_ddgg._faba : _gbag+_ddgg._edgb]
			copy(_cdeb, _geed)
			_cdeb = _cdfa[_gbag+_ddgg._gggca : _gbag+_ddgg._cgedg]
			copy(_cdeb, _fcgd)
		}
		_cebdb := _ede.NewBuffer(_cdfa)
		_, _decd = _f.Copy(w, _cebdb)
		if _decd != nil {
			return _decd
		}
	}
	_faee._gcbc = true
	return nil
}

// ToInteger convert to an integer format.
func (_gffca *PdfColorDeviceRGB) ToInteger(bits int) [3]uint32 {
	_fced := _ced.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_fced * _gffca.R()), uint32(_fced * _gffca.G()), uint32(_fced * _gffca.B())}
}

// AppendContentStream adds content stream by string.  Appends to the last
// contentstream instance if many.
func (_ddee *PdfPage) AppendContentStream(contentStr string) error {
	_becae, _bebbb := _ddee.GetContentStreams()
	if _bebbb != nil {
		return _bebbb
	}
	if len(_becae) == 0 {
		_becae = []string{contentStr}
		return _ddee.SetContentStreams(_becae, _cde.NewFlateEncoder())
	}
	var _aaebd _ede.Buffer
	_aaebd.WriteString(_becae[len(_becae)-1])
	_aaebd.WriteString("\u000a")
	_aaebd.WriteString(contentStr)
	_becae[len(_becae)-1] = _aaebd.String()
	return _ddee.SetContentStreams(_becae, _cde.NewFlateEncoder())
}

// B returns the value of the B component of the color.
func (_ffgd *PdfColorLab) B() float64 { return _ffgd[2] }

// IsColored specifies if the pattern is colored.
func (_fcdd *PdfTilingPattern) IsColored() bool {
	if _fcdd.PaintType != nil && *_fcdd.PaintType == 1 {
		return true
	}
	return false
}

// NewPdfWriter initializes a new PdfWriter.
func NewPdfWriter() PdfWriter {
	_adfeec := PdfWriter{}
	_adfeec._bccde = map[_cde.PdfObject]struct{}{}
	_adfeec._egbccc = []_cde.PdfObject{}
	_adfeec._egffc = map[_cde.PdfObject][]*_cde.PdfObjectDictionary{}
	_adfeec._gbdfb = map[_cde.PdfObject]struct{}{}
	_adfeec._cgdcc.Major = 1
	_adfeec._cgdcc.Minor = 3
	_gfcgb := _cde.MakeDict()
	_bgdffd := []struct {
		_gfeef _cde.PdfObjectName
		_dgbba string
	}{{"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _gfdbd()}, {"\u0043r\u0065\u0061\u0074\u006f\u0072", _ffbc()}, {"\u0041\u0075\u0074\u0068\u006f\u0072", _befce()}, {"\u0053u\u0062\u006a\u0065\u0063\u0074", _gaabf()}, {"\u0054\u0069\u0074l\u0065", _cgcfa()}, {"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _ffabc()}}
	for _, _fbfed := range _bgdffd {
		if _fbfed._dgbba != "" {
			_gfcgb.Set(_fbfed._gfeef, _cde.MakeString(_fbfed._dgbba))
		}
	}
	if _fdfed := _dgdg(); !_fdfed.IsZero() {
		if _dgeac, _ecgbd := NewPdfDateFromTime(_fdfed); _ecgbd == nil {
			_gfcgb.Set("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _dgeac.ToPdfObject())
		}
	}
	if _defbe := _abdae(); !_defbe.IsZero() {
		if _afba, _fbbfc := NewPdfDateFromTime(_defbe); _fbbfc == nil {
			_gfcgb.Set("\u004do\u0064\u0044\u0061\u0074\u0065", _afba.ToPdfObject())
		}
	}
	_bedfab := _cde.PdfIndirectObject{}
	_bedfab.PdfObject = _gfcgb
	_adfeec._fdgbc = &_bedfab
	_adfeec.addObject(&_bedfab)
	_bfdaf := _cde.PdfIndirectObject{}
	_ddebdd := _cde.MakeDict()
	_ddebdd.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0043a\u0074\u0061\u006c\u006f\u0067"))
	_bfdaf.PdfObject = _ddebdd
	_adfeec._eacge = &_bfdaf
	_adfeec.addObject(_adfeec._eacge)
	_gaaee, _dfdbd := _bbeaa("\u0077")
	if _dfdbd != nil {
		_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dfdbd)
	}
	_adfeec._bedaa = _gaaee
	_dgdfb := _cde.PdfIndirectObject{}
	_caedg := _cde.MakeDict()
	_caedg.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0050\u0061\u0067e\u0073"))
	_bdbbd := _cde.PdfObjectArray{}
	_caedg.Set("\u004b\u0069\u0064\u0073", &_bdbbd)
	_caedg.Set("\u0043\u006f\u0075n\u0074", _cde.MakeInteger(0))
	_dgdfb.PdfObject = _caedg
	_adfeec._acbgeg = &_dgdfb
	_adfeec._fbeeb = map[_cde.PdfObject]struct{}{}
	_adfeec.addObject(_adfeec._acbgeg)
	_ddebdd.Set("\u0050\u0061\u0067e\u0073", &_dgdfb)
	_adfeec._fedbb = _ddebdd
	_ad.Log.Trace("\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0025\u0073", _bfdaf)
	return _adfeec
}

// GetCharMetrics returns the character metrics for the specified character code.  A bool flag is
// returned to indicate whether or not the entry was found in the glyph to charcode mapping.
// How it works:
//  1) Return a value the /Widths array (charWidths) if there is one.
//  2) If the font has the same name as a standard 14 font then return width=250.
//  3) Otherwise return no match and let the caller substitute a default.
func (_fgfaa pdfFontSimple) GetCharMetrics(code _gc.CharCode) (_fe.CharMetrics, bool) {
	if _dbeb, _cfdb := _fgfaa._fecg[code]; _cfdb {
		return _fe.CharMetrics{Wx: _dbeb}, true
	}
	if _fe.IsStdFont(_fe.StdFontName(_fgfaa._eeab)) {
		return _fe.CharMetrics{Wx: 250}, true
	}
	return _fe.CharMetrics{}, false
}

// CharcodesToStrings returns the unicode strings corresponding to `charcodes`.
// The int returns are the number of strings and the number of unconvereted codes.
// NOTE: The number of strings returned is equal to the number of charcodes
func (_bggaa *PdfFont) CharcodesToStrings(charcodes []_gc.CharCode) ([]string, int, int) {
	_gaddd := _bggaa.baseFields()
	_bced := make([]string, 0, len(charcodes))
	_fbgbb := 0
	_gafed := _bggaa.Encoder()
	_dfbff := _gaddd._ggebg != nil && _bggaa.IsSimple() && _bggaa.Subtype() == "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" && !_dac.Contains(_gaddd._ggebg.Name(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")
	if !_dfbff && _gafed != nil {
		switch _ddede := _gafed.(type) {
		case _gc.SimpleEncoder:
			_aeffb := _ddede.BaseName()
			if _, _gcfdg := _egee[_aeffb]; _gcfdg {
				for _, _ecdb := range charcodes {
					if _gffgb, _ecef := _gafed.CharcodeToRune(_ecdb); _ecef {
						_bced = append(_bced, string(_gffgb))
					} else {
						_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _ecdb, charcodes, _gaddd.isCIDFont(), _bggaa, _gafed)
						_fbgbb++
						_bced = append(_bced, _fb.MissingCodeString)
					}
				}
				return _bced, len(_bced), _fbgbb
			}
		}
	}
	for _, _dgcb := range charcodes {
		if _gaddd._ggebg != nil {
			if _cega, _adgeb := _gaddd._ggebg.CharcodeToUnicode(_fb.CharCode(_dgcb)); _adgeb {
				_bced = append(_bced, _cega)
				continue
			}
		}
		if _gafed != nil {
			if _egfeg, _gdaeb := _gafed.CharcodeToRune(_dgcb); _gdaeb {
				_bced = append(_bced, string(_egfeg))
				continue
			}
		}
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _dgcb, charcodes, _gaddd.isCIDFont(), _bggaa, _gafed)
		_fbgbb++
		_bced = append(_bced, _fb.MissingCodeString)
	}
	if _fbgbb != 0 {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0043\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0074\u006f\u0020u\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020i\u006ep\u0075\u0074\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0043\u0068\u0061\u0072\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073", len(charcodes), _fbgbb, _bggaa)
	}
	return _bced, len(_bced), _fbgbb
}

// EnableChain adds the specified certificate chain and validation data (OCSP
// and CRL information) for it to the global scope of the document DSS. The
// added data is used for validating any of the signatures present in the
// document. The LTV client attempts to build the certificate chain up to a
// trusted root by downloading any missing certificates.
func (_fecbf *LTV) EnableChain(chain []*_bg.Certificate) error { return _fecbf.enable(nil, chain, "") }

// NewPdfAnnotationWatermark returns a new watermark annotation.
func NewPdfAnnotationWatermark() *PdfAnnotationWatermark {
	_aegg := NewPdfAnnotation()
	_bgbd := &PdfAnnotationWatermark{}
	_bgbd.PdfAnnotation = _aegg
	_aegg.SetContext(_bgbd)
	return _bgbd
}

// PdfFontDescriptor specifies metrics and other attributes of a font and can refer to a FontFile
// for embedded fonts.
// 9.8 Font Descriptors (page 281)
type PdfFontDescriptor struct {
	FontName     _cde.PdfObject
	FontFamily   _cde.PdfObject
	FontStretch  _cde.PdfObject
	FontWeight   _cde.PdfObject
	Flags        _cde.PdfObject
	FontBBox     _cde.PdfObject
	ItalicAngle  _cde.PdfObject
	Ascent       _cde.PdfObject
	Descent      _cde.PdfObject
	Leading      _cde.PdfObject
	CapHeight    _cde.PdfObject
	XHeight      _cde.PdfObject
	StemV        _cde.PdfObject
	StemH        _cde.PdfObject
	AvgWidth     _cde.PdfObject
	MaxWidth     _cde.PdfObject
	MissingWidth _cde.PdfObject
	FontFile     _cde.PdfObject
	FontFile2    _cde.PdfObject
	FontFile3    _cde.PdfObject
	CharSet      _cde.PdfObject
	_gbffb       int
	_efgg        float64
	*fontFile
	_ggdga *_fe.TtfType

	// Additional entries for CIDFonts
	Style  _cde.PdfObject
	Lang   _cde.PdfObject
	FD     _cde.PdfObject
	CIDSet _cde.PdfObject
	_fbff  *_cde.PdfIndirectObject
}

// NewXObjectImageFromStream builds the image xobject from a stream object.
// An image dictionary is the dictionary portion of a stream object representing an image XObject.
func NewXObjectImageFromStream(stream *_cde.PdfObjectStream) (*XObjectImage, error) {
	_cgdff := &XObjectImage{}
	_cgdff._bbaed = stream
	_cgeed := *(stream.PdfObjectDictionary)
	_cgbga, _cbgee := _cde.NewEncoderFromStream(stream)
	if _cbgee != nil {
		return nil, _cbgee
	}
	_cgdff.Filter = _cgbga
	if _edcd := _cde.TraceToDirectObject(_cgeed.Get("\u0057\u0069\u0064t\u0068")); _edcd != nil {
		_aedbe, _afcfg := _edcd.(*_cde.PdfObjectInteger)
		if !_afcfg {
			return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_eeba := int64(*_aedbe)
		_cgdff.Width = &_eeba
	} else {
		return nil, _ceg.New("\u0077\u0069\u0064\u0074\u0068\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _acgcg := _cde.TraceToDirectObject(_cgeed.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _acgcg != nil {
		_gacbf, _edfdf := _acgcg.(*_cde.PdfObjectInteger)
		if !_edfdf {
			return nil, _ceg.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_aacfd := int64(*_gacbf)
		_cgdff.Height = &_aacfd
	} else {
		return nil, _ceg.New("\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	if _gaccf := _cde.TraceToDirectObject(_cgeed.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _gaccf != nil {
		_bbebf, _egdcg := NewPdfColorspaceFromPdfObject(_gaccf)
		if _egdcg != nil {
			return nil, _egdcg
		}
		_cgdff.ColorSpace = _bbebf
	} else {
		_ad.Log.Debug("\u0058O\u0062\u006a\u0065c\u0074\u0020\u0049m\u0061ge\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067 1\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065n\u0074\u0020\u002d\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047r\u0061\u0079")
		_cgdff.ColorSpace = NewPdfColorspaceDeviceGray()
	}
	if _cegga := _cde.TraceToDirectObject(_cgeed.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _cegga != nil {
		_bfdg, _aagbg := _cegga.(*_cde.PdfObjectInteger)
		if !_aagbg {
			return nil, _ceg.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_bafaba := int64(*_bfdg)
		_cgdff.BitsPerComponent = &_bafaba
	}
	_cgdff.Intent = _cgeed.Get("\u0049\u006e\u0074\u0065\u006e\u0074")
	_cgdff.ImageMask = _cgeed.Get("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k")
	_cgdff.Mask = _cgeed.Get("\u004d\u0061\u0073\u006b")
	_cgdff.Decode = _cgeed.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	_cgdff.Interpolate = _cgeed.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065")
	_cgdff.Alternatives = _cgeed.Get("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073")
	_cgdff.SMask = _cgeed.Get("\u0053\u004d\u0061s\u006b")
	_cgdff.SMaskInData = _cgeed.Get("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061")
	_cgdff.Matte = _cgeed.Get("\u004d\u0061\u0074t\u0065")
	_cgdff.Name = _cgeed.Get("\u004e\u0061\u006d\u0065")
	_cgdff.StructParent = _cgeed.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_cgdff.ID = _cgeed.Get("\u0049\u0044")
	_cgdff.OPI = _cgeed.Get("\u004f\u0050\u0049")
	_cgdff.Metadata = _cgeed.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	_cgdff.OC = _cgeed.Get("\u004f\u0043")
	_cgdff.Stream = stream.Stream
	return _cgdff, nil
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_ggdb pdfCIDFontType2) GetCharMetrics(code _gc.CharCode) (_fe.CharMetrics, bool) {
	if _aggcd, _cbfdc := _ggdb._fdgc[code]; _cbfdc {
		return _fe.CharMetrics{Wx: _aggcd}, true
	}
	_cagde := rune(code)
	_fccag, _gfege := _ggdb._dfefe[_cagde]
	if !_gfege {
		_fccag = int(_ggdb._eedac)
	}
	return _fe.CharMetrics{Wx: float64(_fccag)}, true
}
func (_dbfa *PdfReader) newPdfAnnotationCircleFromDict(_ffg *_cde.PdfObjectDictionary) (*PdfAnnotationCircle, error) {
	_abaa := PdfAnnotationCircle{}
	_ebcg, _gedg := _dbfa.newPdfAnnotationMarkupFromDict(_ffg)
	if _gedg != nil {
		return nil, _gedg
	}
	_abaa.PdfAnnotationMarkup = _ebcg
	_abaa.BS = _ffg.Get("\u0042\u0053")
	_abaa.IC = _ffg.Get("\u0049\u0043")
	_abaa.BE = _ffg.Get("\u0042\u0045")
	_abaa.RD = _ffg.Get("\u0052\u0044")
	return &_abaa, nil
}

// PdfAnnotationWatermark represents Watermark annotations.
// (Section 12.5.6.22).
type PdfAnnotationWatermark struct {
	*PdfAnnotation
	FixedPrint _cde.PdfObject
}

// NewPdfAnnotationPopup returns a new popup annotation.
func NewPdfAnnotationPopup() *PdfAnnotationPopup {
	_dcd := NewPdfAnnotation()
	_bgdf := &PdfAnnotationPopup{}
	_bgdf.PdfAnnotation = _dcd
	_dcd.SetContext(_bgdf)
	return _bgdf
}

// HasXObjectByName checks if an XObject with a specified keyName is defined.
func (_dcedfa *PdfPageResources) HasXObjectByName(keyName _cde.PdfObjectName) bool {
	_fbcegf, _ := _dcedfa.GetXObjectByName(keyName)
	return _fbcegf != nil
}

// PdfAnnotationPolyLine represents PolyLine annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolyLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _cde.PdfObject
	LE       _cde.PdfObject
	BS       _cde.PdfObject
	IC       _cde.PdfObject
	BE       _cde.PdfObject
	IT       _cde.PdfObject
	Measure  _cde.PdfObject
}

// CompliancePdfReader is a wrapper over PdfReader that is used for verifying if the input Pdf document matches the
// compliance rules of standards like PDF/A.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
type CompliancePdfReader struct {
	*PdfReader
	_abbc _cde.ParserMetadata
}

// NewPdfColorspaceFromPdfObject loads a PdfColorspace from a PdfObject.  Returns an error if there is
// a failure in loading.
func NewPdfColorspaceFromPdfObject(obj _cde.PdfObject) (PdfColorspace, error) {
	if obj == nil {
		return nil, nil
	}
	var _bcbb *_cde.PdfIndirectObject
	var _gcega *_cde.PdfObjectName
	var _ebd *_cde.PdfObjectArray
	if _gbace, _aged := obj.(*_cde.PdfIndirectObject); _aged {
		_bcbb = _gbace
	}
	obj = _cde.TraceToDirectObject(obj)
	switch _bagb := obj.(type) {
	case *_cde.PdfObjectArray:
		_ebd = _bagb
	case *_cde.PdfObjectName:
		_gcega = _bagb
	}
	if _gcega != nil {
		switch *_gcega {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			return NewPdfColorspaceDeviceGray(), nil
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			return NewPdfColorspaceDeviceRGB(), nil
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return NewPdfColorspaceDeviceCMYK(), nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return NewPdfColorspaceSpecialPattern(), nil
		default:
			_ad.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0025\u0073", *_gcega)
			return nil, _cdab
		}
	}
	if _ebd != nil && _ebd.Len() > 0 {
		var _ffca _cde.PdfObject = _bcbb
		if _bcbb == nil {
			_ffca = _ebd
		}
		if _eggcf, _dgcf := _cde.GetName(_ebd.Get(0)); _dgcf {
			switch _eggcf.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				if _ebd.Len() == 1 {
					return NewPdfColorspaceDeviceGray(), nil
				}
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				if _ebd.Len() == 1 {
					return NewPdfColorspaceDeviceRGB(), nil
				}
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _ebd.Len() == 1 {
					return NewPdfColorspaceDeviceCMYK(), nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079":
				return _deec(_ffca)
			case "\u0043\u0061\u006c\u0052\u0047\u0042":
				return _eddbb(_ffca)
			case "\u004c\u0061\u0062":
				return _bcfg(_ffca)
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064":
				return _gfce(_ffca)
			case "\u0050a\u0074\u0074\u0065\u0072\u006e":
				return _agca(_ffca)
			case "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return _eafbf(_ffca)
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e":
				return _dcfg(_ffca)
			case "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return _fecd(_ffca)
			default:
				_ad.Log.Debug("A\u0072\u0072\u0061\u0079\u0020\u0077i\u0074\u0068\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0061m\u0065:\u0020\u0025\u0073", *_eggcf)
			}
		}
	}
	_ad.Log.Debug("\u0050\u0044\u0046\u0020\u0046i\u006c\u0065\u0020\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", obj.String())
	return nil, ErrTypeCheck
}

// UpdatePage updates the `page` in the new revision if it has changed.
func (_fgfcf *PdfAppender) UpdatePage(page *PdfPage) {
	_fgfcf.updateObjectsDeep(page.ToPdfObject(), nil)
}

// NewPdfFontFromTTF loads a TTF font and returns a PdfFont type that can be
// used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTF.
func NewPdfFontFromTTF(r _f.ReadSeeker) (*PdfFont, error) {
	const _fadb = _gc.CharCode(32)
	const _eeaab = _gc.CharCode(255)
	_bbfg, _aadfga := _bb.ReadAll(r)
	if _aadfga != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _aadfga)
		return nil, _aadfga
	}
	_ccca, _aadfga := _fe.TtfParse(_ede.NewReader(_bbfg))
	if _aadfga != nil {
		_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0054\u0054F\u0020\u0066\u006fn\u0074:\u0020\u0025\u0076", _aadfga)
		return nil, _aadfga
	}
	_dddbc := &pdfFontSimple{_fecg: make(map[_gc.CharCode]float64), fontCommon: fontCommon{_dcbc: "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065"}}
	_dddbc._efeaf = _gc.NewWinAnsiEncoder()
	_dddbc._eeab = _ccca.PostScriptName
	_dddbc.FirstChar = _cde.MakeInteger(int64(_fadb))
	_dddbc.LastChar = _cde.MakeInteger(int64(_eeaab))
	_agbb := 1000.0 / float64(_ccca.UnitsPerEm)
	if len(_ccca.Widths) <= 0 {
		return nil, _ceg.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_gced := _agbb * float64(_ccca.Widths[0])
	_feff := make([]float64, 0, _eeaab-_fadb+1)
	for _abeg := _fadb; _abeg <= _eeaab; _abeg++ {
		_aeedce, _bcfgf := _dddbc.Encoder().CharcodeToRune(_abeg)
		if !_bcfgf {
			_ad.Log.Debug("\u0052u\u006e\u0065\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u0028c\u006f\u0064\u0065\u003a\u0020\u0025\u0064)", _abeg)
			_feff = append(_feff, _gced)
			continue
		}
		_cfgfc, _gfagc := _ccca.Chars[_aeedce]
		if !_gfagc {
			_ad.Log.Debug("R\u0075\u006e\u0065\u0020no\u0074 \u0069\u006e\u0020\u0054\u0054F\u0020\u0043\u0068\u0061\u0072\u0073")
			_feff = append(_feff, _gced)
			continue
		}
		_afbbg := _agbb * float64(_ccca.Widths[_cfgfc])
		_feff = append(_feff, _afbbg)
	}
	_dddbc.Widths = _cde.MakeIndirectObject(_cde.MakeArrayFromFloats(_feff))
	if len(_feff) < int(_eeaab-_fadb+1) {
		_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u006f\u0066\u0020\u0077\u0069\u0064\u0074\u0068s,\u0020\u0025\u0064 \u003c \u0025\u0064", len(_feff), 255-32+1)
		return nil, _cde.ErrRangeError
	}
	for _bfbb := _fadb; _bfbb <= _eeaab; _bfbb++ {
		_dddbc._fecg[_bfbb] = _feff[_bfbb-_fadb]
	}
	_dddbc.Encoding = _cde.MakeName("\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	_dabec := &PdfFontDescriptor{}
	_dabec.FontName = _cde.MakeName(_ccca.PostScriptName)
	_dabec.Ascent = _cde.MakeFloat(_agbb * float64(_ccca.TypoAscender))
	_dabec.Descent = _cde.MakeFloat(_agbb * float64(_ccca.TypoDescender))
	_dabec.CapHeight = _cde.MakeFloat(_agbb * float64(_ccca.CapHeight))
	_dabec.FontBBox = _cde.MakeArrayFromFloats([]float64{_agbb * float64(_ccca.Xmin), _agbb * float64(_ccca.Ymin), _agbb * float64(_ccca.Xmax), _agbb * float64(_ccca.Ymax)})
	_dabec.ItalicAngle = _cde.MakeFloat(_ccca.ItalicAngle)
	_dabec.MissingWidth = _cde.MakeFloat(_agbb * float64(_ccca.Widths[0]))
	_gegebb, _aadfga := _cde.MakeStream(_bbfg, _cde.NewFlateEncoder())
	if _aadfga != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _aadfga)
		return nil, _aadfga
	}
	_gegebb.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _cde.MakeInteger(int64(len(_bbfg))))
	_dabec.FontFile2 = _gegebb
	if _ccca.Bold {
		_dabec.StemV = _cde.MakeInteger(120)
	} else {
		_dabec.StemV = _cde.MakeInteger(70)
	}
	_fbdfc := _gbdf
	if _ccca.IsFixedPitch {
		_fbdfc |= _gbcd
	}
	if _ccca.ItalicAngle != 0 {
		_fbdfc |= _ecebb
	}
	_dabec.Flags = _cde.MakeInteger(int64(_fbdfc))
	_dddbc._fagf = _dabec
	_gdbgg := &PdfFont{_gbcff: _dddbc}
	return _gdbgg, nil
}

// GetStandardApplier gets currently used StandardApplier..
func (_edcafc *PdfWriter) GetStandardApplier() StandardApplier { return _edcafc._cfaf }
func _fdcef(_bebcg *PdfPage) {
	_gggcd := _cde.PdfObjectName("\u0055\u0046\u0031")
	if !_bebcg.Resources.HasFontByName(_gggcd) {
		_bebcg.Resources.SetFontByName(_gggcd, DefaultFont().ToPdfObject())
	}
	var _acdac []string
	_acdac = append(_acdac, "\u0071")
	_acdac = append(_acdac, "\u0042\u0054")
	_acdac = append(_acdac, _ee.Sprintf("\u002f%\u0073\u0020\u0031\u0034\u0020\u0054f", _gggcd.String()))
	_acdac = append(_acdac, "\u0031\u0020\u0030\u0020\u0030\u0020\u0072\u0067")
	_acdac = append(_acdac, "\u0031\u0030\u0020\u0031\u0030\u0020\u0054\u0064")
	_acdac = append(_acdac, "\u0045\u0054")
	_acdac = append(_acdac, "\u0051")
	_dbddb := _dac.Join(_acdac, "\u000a")
	_bebcg.AddContentStreamByString(_dbddb)
	_bebcg.ToPdfObject()
}

// ColorToRGB converts a CalGray color to an RGB color.
func (_fcab *PdfColorspaceCalGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ccacg, _fbfg := color.(*PdfColorCalGray)
	if !_fbfg {
		_ad.Log.Debug("\u0049n\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006eo\u0074\u0020\u0063\u0061\u006c\u0020\u0067\u0072\u0061\u0079")
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	ANorm := _ccacg.Val()
	X := _fcab.WhitePoint[0] * _ced.Pow(ANorm, _fcab.Gamma)
	Y := _fcab.WhitePoint[1] * _ced.Pow(ANorm, _fcab.Gamma)
	Z := _fcab.WhitePoint[2] * _ced.Pow(ANorm, _fcab.Gamma)
	_gfed := 3.240479*X + -1.537150*Y + -0.498535*Z
	_fcdc := -0.969256*X + 1.875992*Y + 0.041556*Z
	_eea := 0.055648*X + -0.204043*Y + 1.057311*Z
	_gfed = _ced.Min(_ced.Max(_gfed, 0), 1.0)
	_fcdc = _ced.Min(_ced.Max(_fcdc, 0), 1.0)
	_eea = _ced.Min(_ced.Max(_eea, 0), 1.0)
	return NewPdfColorDeviceRGB(_gfed, _fcdc, _eea), nil
}

var ErrColorOutOfRange = _ceg.New("\u0063o\u006co\u0072\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")

// PdfColorDeviceGray represents a grayscale color value that shall be represented by a single number in the
// range 0.0 to 1.0 where 0.0 corresponds to black and 1.0 to white.
type PdfColorDeviceGray float64

// NewStandardPdfOutputIntent creates a new standard PdfOutputIntent.
func NewStandardPdfOutputIntent(outputCondition, outputConditionIdentifier, registryName string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, RegistryName: registryName, DestOutputProfile: destOutputProfile, ColorComponents: colorComponents, _acbcb: _cde.MakeDict()}
}

// OutlineItem represents a PDF outline item dictionary (Table 153 - pp. 376 - 377).
type OutlineItem struct {
	Title   string         `json:"title"`
	Dest    OutlineDest    `json:"dest"`
	Entries []*OutlineItem `json:"entries,omitempty"`
}

func _efge(_edag *fontCommon) *pdfFontSimple { return &pdfFontSimple{fontCommon: *_edag} }

// NewStandard14FontMustCompile returns the standard 14 font named `basefont` as a *PdfFont.
// If `basefont` is one of the 14 Standard14Font values defined above then NewStandard14FontMustCompile
// is guaranteed to succeed.
func NewStandard14FontMustCompile(basefont StdFontName) *PdfFont {
	_effce, _defc := NewStandard14Font(basefont)
	if _defc != nil {
		panic(_ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0031\u0034\u0046\u006f\u006e\u0074 \u0025\u0023\u0071", basefont))
	}
	return _effce
}

// NewPdfActionGoTo3DView returns a new "goTo3DView" action.
func NewPdfActionGoTo3DView() *PdfActionGoTo3DView {
	_gdg := NewPdfAction()
	_gg := &PdfActionGoTo3DView{}
	_gg.PdfAction = _gdg
	_gdg.SetContext(_gg)
	return _gg
}

// PdfAnnotationFreeText represents FreeText annotations.
// (Section 12.5.6.6).
type PdfAnnotationFreeText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	DA _cde.PdfObject
	Q  _cde.PdfObject
	RC _cde.PdfObject
	DS _cde.PdfObject
	CL _cde.PdfObject
	IT _cde.PdfObject
	BE _cde.PdfObject
	RD _cde.PdfObject
	BS _cde.PdfObject
	LE _cde.PdfObject
}

var _ pdfFont = (*pdfCIDFontType0)(nil)

// PdfFont represents an underlying font structure which can be of type:
// - Type0
// - Type1
// - TrueType
// etc.
type PdfFont struct{ _gbcff pdfFont }

// PdfActionGoToE represents a GoToE action.
type PdfActionGoToE struct {
	*PdfAction
	F         *PdfFilespec
	D         _cde.PdfObject
	NewWindow _cde.PdfObject
	T         _cde.PdfObject
}

// PdfAnnotationHighlight represents Highlight annotations.
// (Section 12.5.6.10).
type PdfAnnotationHighlight struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _cde.PdfObject
}

// String returns a string representation of PdfTransformParamsDocMDP.
func (_ddfbc *PdfTransformParamsDocMDP) String() string {
	return _ee.Sprintf("\u0025\u0073\u0020\u0050\u003a\u0020\u0025\u0073\u0020V\u003a\u0020\u0025\u0073", _ddfbc.Type, _ddfbc.P, _ddfbc.V)
}

// GetIndirectObjectByNumber retrieves and returns a specific PdfObject by object number.
func (_aaegg *PdfReader) GetIndirectObjectByNumber(number int) (_cde.PdfObject, error) {
	_ggeed, _fdcd := _aaegg._aggcgb.LookupByNumber(number)
	return _ggeed, _fdcd
}

// RepairAcroForm attempts to rebuild the AcroForm fields using the widget
// annotations present in the document pages. Pass nil for the opts parameter
// in order to use the default options.
// NOTE: Currently, the opts parameter is declared in order to enable adding
// future options, but passing nil will always result in the default options
// being used.
func (_bdgdc *PdfReader) RepairAcroForm(opts *AcroFormRepairOptions) error {
	var _ebdfb []*PdfField
	_eagac := map[*_cde.PdfIndirectObject]struct{}{}
	for _, _eedf := range _bdgdc.PageList {
		_efceg, _gdefa := _eedf.GetAnnotations()
		if _gdefa != nil {
			return _gdefa
		}
		for _, _bbfgb := range _efceg {
			var _aedeg *PdfField
			switch _fcgcb := _bbfgb.GetContext().(type) {
			case *PdfAnnotationWidget:
				if _fcgcb._dbf != nil {
					_aedeg = _fcgcb._dbf
					break
				}
				if _acgeb, _ecadf := _cde.GetIndirect(_fcgcb.Parent); _ecadf {
					_aedeg, _gdefa = _bdgdc.newPdfFieldFromIndirectObject(_acgeb, nil)
					if _gdefa == nil {
						break
					}
					_ad.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _acgeb, _gdefa)
				}
				if _fcgcb._bddg != nil {
					_aedeg, _gdefa = _bdgdc.newPdfFieldFromIndirectObject(_fcgcb._bddg, nil)
					if _gdefa == nil {
						break
					}
					_ad.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _fcgcb._bddg, _gdefa)
				}
			}
			if _aedeg == nil {
				continue
			}
			if _, _abbdf := _eagac[_aedeg._afgc]; _abbdf {
				continue
			}
			_eagac[_aedeg._afgc] = struct{}{}
			_ebdfb = append(_ebdfb, _aedeg)
		}
	}
	if len(_ebdfb) == 0 {
		return nil
	}
	if _bdgdc.AcroForm == nil {
		_bdgdc.AcroForm = NewPdfAcroForm()
	}
	_bdgdc.AcroForm.Fields = &_ebdfb
	return nil
}
func (_dcbcfc *PdfWriter) addObjects(_baefb _cde.PdfObject) error {
	_ad.Log.Trace("\u0041d\u0064i\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073\u0021")
	if _gbfdg, _fdbeg := _baefb.(*_cde.PdfIndirectObject); _fdbeg {
		_ad.Log.Trace("\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074")
		_ad.Log.Trace("\u002d \u0025\u0073\u0020\u0028\u0025\u0070)", _baefb, _gbfdg)
		_ad.Log.Trace("\u002d\u0020\u0025\u0073", _gbfdg.PdfObject)
		if _dcbcfc.addObject(_gbfdg) {
			_bfeff := _dcbcfc.addObjects(_gbfdg.PdfObject)
			if _bfeff != nil {
				return _bfeff
			}
		}
		return nil
	}
	if _fdccd, _dgdcg := _baefb.(*_cde.PdfObjectStream); _dgdcg {
		_ad.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d")
		_ad.Log.Trace("\u002d \u0025\u0073\u0020\u0025\u0070", _baefb, _baefb)
		if _dcbcfc.addObject(_fdccd) {
			_fcaa := _dcbcfc.addObjects(_fdccd.PdfObjectDictionary)
			if _fcaa != nil {
				return _fcaa
			}
		}
		return nil
	}
	if _bgcga, _fgea := _baefb.(*_cde.PdfObjectDictionary); _fgea {
		_ad.Log.Trace("\u0044\u0069\u0063\u0074")
		_ad.Log.Trace("\u002d\u0020\u0025\u0073", _baefb)
		for _, _eaagd := range _bgcga.Keys() {
			_facgag := _bgcga.Get(_eaagd)
			if _ceca, _edebb := _facgag.(*_cde.PdfObjectReference); _edebb {
				_facgag = _ceca.Resolve()
				_bgcga.Set(_eaagd, _facgag)
			}
			if _eaagd != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _fdbed := _dcbcfc.addObjects(_facgag); _fdbed != nil {
					return _fdbed
				}
			} else {
				if _, _gdbc := _facgag.(*_cde.PdfObjectNull); _gdbc {
					continue
				}
				if _cfeaa := _dcbcfc.hasObject(_facgag); !_cfeaa {
					_ad.Log.Debug("P\u0061\u0072\u0065\u006e\u0074\u0020o\u0062\u006a\u0020\u006e\u006f\u0074 \u0061\u0064\u0064\u0065\u0064\u0020\u0079e\u0074\u0021\u0021\u0020\u0025\u0054\u0020\u0025\u0070\u0020%\u0076", _facgag, _facgag, _facgag)
					_dcbcfc._egffc[_facgag] = append(_dcbcfc._egffc[_facgag], _bgcga)
				}
			}
		}
		return nil
	}
	if _ccgfc, _bcggd := _baefb.(*_cde.PdfObjectArray); _bcggd {
		_ad.Log.Trace("\u0041\u0072\u0072a\u0079")
		_ad.Log.Trace("\u002d\u0020\u0025\u0073", _baefb)
		if _ccgfc == nil {
			return _ceg.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _fggcf, _acagb := range _ccgfc.Elements() {
			if _aafg, _caegb := _acagb.(*_cde.PdfObjectReference); _caegb {
				_acagb = _aafg.Resolve()
				_ccgfc.Set(_fggcf, _acagb)
			}
			if _daegf := _dcbcfc.addObjects(_acagb); _daegf != nil {
				return _daegf
			}
		}
		return nil
	}
	if _, _bfaca := _baefb.(*_cde.PdfObjectReference); _bfaca {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0062\u0065\u0020\u0061\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u002d\u0020\u0067\u006f\u0074 \u0025\u0023\u0076\u0021", _baefb)
		return _ceg.New("r\u0065\u0066\u0065\u0072en\u0063e\u0020\u006e\u006f\u0074\u0020a\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	return nil
}
func (_gdbed *PdfReader) loadOutlines() (*PdfOutlineTreeNode, error) {
	if _gdbed._aggcgb.GetCrypter() != nil && !_gdbed._aggcgb.IsAuthenticated() {
		return nil, _ee.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_fabce := _gdbed._efabe
	_degcd := _fabce.Get("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	if _degcd == nil {
		return nil, nil
	}
	_ad.Log.Trace("\u002d\u0048\u0061\u0073\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	_gaec := _cde.ResolveReference(_degcd)
	_ad.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0020\u0072\u006f\u006f\u0074\u003a\u0020\u0025\u0076", _gaec)
	if _fgbcb := _cde.IsNullObject(_gaec); _fgbcb {
		_ad.Log.Trace("\u004f\u0075\u0074li\u006e\u0065\u0020\u0072\u006f\u006f\u0074\u0020\u0069s\u0020n\u0075l\u006c \u002d\u0020\u006e\u006f\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
		return nil, nil
	}
	_ddcfd, _bbeeg := _gaec.(*_cde.PdfIndirectObject)
	if !_bbeeg {
		if _, _egdc := _cde.GetDict(_gaec); !_egdc {
			_ad.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			return nil, nil
		}
		_ad.Log.Debug("\u004f\u0075t\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u0069s\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		_ddcfd = _cde.MakeIndirectObject(_gaec)
	}
	_aead, _bbeeg := _ddcfd.PdfObject.(*_cde.PdfObjectDictionary)
	if !_bbeeg {
		return nil, _ceg.New("\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y")
	}
	_ad.Log.Trace("O\u0075\u0074\u006c\u0069ne\u0020r\u006f\u006f\u0074\u0020\u0064i\u0063\u0074\u003a\u0020\u0025\u0076", _aead)
	_ggcd, _, _baafgb := _gdbed.buildOutlineTree(_ddcfd, nil, nil, nil)
	if _baafgb != nil {
		return nil, _baafgb
	}
	_ad.Log.Trace("\u0052\u0065\u0073\u0075\u006c\u0074\u0069\u006e\u0067\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065:\u0020\u0025\u0076", _ggcd)
	return _ggcd, nil
}

// ToPdfObject converts the pdfCIDFontType2 to a PDF representation.
func (_gadeg *pdfCIDFontType2) ToPdfObject() _cde.PdfObject {
	if _gadeg._abca == nil {
		_gadeg._abca = &_cde.PdfIndirectObject{}
	}
	_fccg := _gadeg.baseFields().asPdfObjectDictionary("\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032")
	_gadeg._abca.PdfObject = _fccg
	if _gadeg.CIDSystemInfo != nil {
		_fccg.Set("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f", _gadeg.CIDSystemInfo)
	}
	if _gadeg.DW != nil {
		_fccg.Set("\u0044\u0057", _gadeg.DW)
	}
	if _gadeg.DW2 != nil {
		_fccg.Set("\u0044\u0057\u0032", _gadeg.DW2)
	}
	if _gadeg.W != nil {
		_fccg.Set("\u0057", _gadeg.W)
	}
	if _gadeg.W2 != nil {
		_fccg.Set("\u0057\u0032", _gadeg.W2)
	}
	if _gadeg.CIDToGIDMap != nil {
		_fccg.Set("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070", _gadeg.CIDToGIDMap)
	}
	return _gadeg._abca
}

// ToPdfObject implements interface PdfModel.
func (_bba *PdfActionSound) ToPdfObject() _cde.PdfObject {
	_bba.PdfAction.ToPdfObject()
	_cda := _bba._bc
	_cbd := _cda.PdfObject.(*_cde.PdfObjectDictionary)
	_cbd.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeSound)))
	_cbd.SetIfNotNil("\u0053\u006f\u0075n\u0064", _bba.Sound)
	_cbd.SetIfNotNil("\u0056\u006f\u006c\u0075\u006d\u0065", _bba.Volume)
	_cbd.SetIfNotNil("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073", _bba.Synchronous)
	_cbd.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _bba.Repeat)
	_cbd.SetIfNotNil("\u004d\u0069\u0078", _bba.Mix)
	return _cda
}

// GetAnnotations returns the list of page annotations for `page`. If not loaded attempts to load the
// annotations, otherwise returns the loaded list.
func (_cdfec *PdfPage) GetAnnotations() ([]*PdfAnnotation, error) {
	if _cdfec._cefe != nil {
		return _cdfec._cefe, nil
	}
	if _cdfec.Annots == nil {
		_cdfec._cefe = []*PdfAnnotation{}
		return nil, nil
	}
	if _cdfec._gfcbe == nil {
		_cdfec._cefe = []*PdfAnnotation{}
		return nil, nil
	}
	_fbgg, _eefgf := _cdfec._gfcbe.loadAnnotations(_cdfec.Annots)
	if _eefgf != nil {
		return nil, _eefgf
	}
	if _fbgg == nil {
		_cdfec._cefe = []*PdfAnnotation{}
	}
	_cdfec._cefe = _fbgg
	return _cdfec._cefe, nil
}

// NewPdfColorLab returns a new Lab color.
func NewPdfColorLab(l, a, b float64) *PdfColorLab { _bcbe := PdfColorLab{l, a, b}; return &_bcbe }

// NewPdfDate returns a new PdfDate object from a PDF date string (see 7.9.4 Dates).
// format: "D: YYYYMMDDHHmmSSOHH'mm"
func NewPdfDate(dateStr string) (PdfDate, error) {
	_bcefa, _aefce := _cce.ParsePdfTime(dateStr)
	if _aefce != nil {
		return PdfDate{}, _aefce
	}
	return NewPdfDateFromTime(_bcefa)
}

// StandardValidator is the interface that is used for the PDF StandardImplementer validation for the PDF document.
// It is using a CompliancePdfReader which is expected to give more Metadata during reading process.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
type StandardValidator interface {

	// ValidateStandard checks if the input reader
	ValidateStandard(_ddebd *CompliancePdfReader) error
}

func (_aceg *PdfAppender) updateObjectsDeep(_bdbg _cde.PdfObject, _agag map[_cde.PdfObject]struct{}) {
	if _agag == nil {
		_agag = map[_cde.PdfObject]struct{}{}
	}
	if _, _ggeb := _agag[_bdbg]; _ggeb || _bdbg == nil {
		return
	}
	_agag[_bdbg] = struct{}{}
	_bdcg := _cde.ResolveReferencesDeep(_bdbg, _aceg._egbc)
	if _bdcg != nil {
		_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bdcg)
	}
	switch _ggaa := _bdbg.(type) {
	case *_cde.PdfIndirectObject:
		switch {
		case _ggaa.GetParser() == _aceg._cfad._aggcgb:
			return
		case _ggaa.GetParser() == _aceg.Reader._aggcgb:
			_dgfa, _ := _aceg._cfad.GetIndirectObjectByNumber(int(_ggaa.ObjectNumber))
			_bfag, _bbgb := _dgfa.(*_cde.PdfIndirectObject)
			if _bbgb && _bfag != nil {
				if _bfag.PdfObject != _ggaa.PdfObject && _bfag.PdfObject.WriteString() != _ggaa.PdfObject.WriteString() {
					_aceg.addNewObject(_bdbg)
					_aceg._adfa[_bdbg] = _ggaa.ObjectNumber
				}
			}
		default:
			_aceg.addNewObject(_bdbg)
		}
		_aceg.updateObjectsDeep(_ggaa.PdfObject, _agag)
	case *_cde.PdfObjectArray:
		for _, _ggbc := range _ggaa.Elements() {
			_aceg.updateObjectsDeep(_ggbc, _agag)
		}
	case *_cde.PdfObjectDictionary:
		for _, _bbcg := range _ggaa.Keys() {
			_aceg.updateObjectsDeep(_ggaa.Get(_bbcg), _agag)
		}
	case *_cde.PdfObjectStreams:
		if _ggaa.GetParser() != _aceg._cfad._aggcgb {
			for _, _ddaf := range _ggaa.Elements() {
				_aceg.updateObjectsDeep(_ddaf, _agag)
			}
		}
	case *_cde.PdfObjectStream:
		switch {
		case _ggaa.GetParser() == _aceg._cfad._aggcgb:
			return
		case _ggaa.GetParser() == _aceg.Reader._aggcgb:
			if _aeb, _babf := _aceg._cfad._aggcgb.LookupByReference(_ggaa.PdfObjectReference); _babf == nil {
				var _cbfc bool
				if _bafd, _fegfb := _cde.GetStream(_aeb); _fegfb && _ede.Equal(_bafd.Stream, _ggaa.Stream) {
					_cbfc = true
				}
				if _cbffe, _eceg := _cde.GetDict(_aeb); _cbfc && _eceg {
					_cbfc = _cbffe.WriteString() == _ggaa.PdfObjectDictionary.WriteString()
				}
				if _cbfc {
					return
				}
			}
			if _ggaa.ObjectNumber != 0 {
				_aceg._adfa[_bdbg] = _ggaa.ObjectNumber
			}
		default:
			if _, _ebad := _aceg._aadd[_bdbg]; !_ebad {
				_aceg.addNewObject(_bdbg)
			}
		}
		_aceg.updateObjectsDeep(_ggaa.PdfObjectDictionary, _agag)
	}
}

// Hasher is the interface that wraps the basic Write method.
type Hasher interface {
	Write(_efdccd []byte) (_fgdag int, _fgbab error)
}

var (
	_dccfe  _c.Mutex
	_bfcdae = ""
	_gffce  _ce.Time
	_dbbgaa = ""
	_bdcea  = ""
	_cebaa  _ce.Time
	_adcff  = ""
	_gbccbb = ""
	_dgcdb  = ""
)

func (_cggc *PdfReader) newPdfAnnotationMarkupFromDict(_cbef *_cde.PdfObjectDictionary) (*PdfAnnotationMarkup, error) {
	_cfe := &PdfAnnotationMarkup{}
	if _eba := _cbef.Get("\u0054"); _eba != nil {
		_cfe.T = _eba
	}
	if _fgbd := _cbef.Get("\u0050\u006f\u0070u\u0070"); _fgbd != nil {
		_cecb, _fege := _fgbd.(*_cde.PdfIndirectObject)
		if !_fege {
			if _, _egdd := _fgbd.(*_cde.PdfObjectNull); !_egdd {
				return nil, _ceg.New("p\u006f\u0070\u0075\u0070\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074\u0020t\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			}
		} else {
			_ecfa, _afc := _cggc.newPdfAnnotationFromIndirectObject(_cecb)
			if _afc != nil {
				return nil, _afc
			}
			if _ecfa != nil {
				_cbgd, _bfbad := _ecfa._bea.(*PdfAnnotationPopup)
				if !_bfbad {
					return nil, _ceg.New("\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0020\u0070\u006f\u0070\u0075\u0070\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e")
				}
				_cfe.Popup = _cbgd
			}
		}
	}
	if _bcbc := _cbef.Get("\u0043\u0041"); _bcbc != nil {
		_cfe.CA = _bcbc
	}
	if _bfdc := _cbef.Get("\u0052\u0043"); _bfdc != nil {
		_cfe.RC = _bfdc
	}
	if _geca := _cbef.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"); _geca != nil {
		_cfe.CreationDate = _geca
	}
	if _adec := _cbef.Get("\u0049\u0052\u0054"); _adec != nil {
		_cfe.IRT = _adec
	}
	if _abfdb := _cbef.Get("\u0053\u0075\u0062\u006a"); _abfdb != nil {
		_cfe.Subj = _abfdb
	}
	if _cfc := _cbef.Get("\u0052\u0054"); _cfc != nil {
		_cfe.RT = _cfc
	}
	if _caed := _cbef.Get("\u0049\u0054"); _caed != nil {
		_cfe.IT = _caed
	}
	if _cbde := _cbef.Get("\u0045\u0078\u0044\u0061\u0074\u0061"); _cbde != nil {
		_cfe.ExData = _cbde
	}
	return _cfe, nil
}

// NewPdfColorspaceCalGray returns a new CalGray colorspace object.
func NewPdfColorspaceCalGray() *PdfColorspaceCalGray {
	_gcbd := &PdfColorspaceCalGray{}
	_gcbd.BlackPoint = []float64{0.0, 0.0, 0.0}
	_gcbd.Gamma = 1
	return _gcbd
}
func _gdfg(_dbede *fontCommon) *pdfCIDFontType2 { return &pdfCIDFontType2{fontCommon: *_dbede} }

// PdfColorspaceDeviceN represents a DeviceN color space. DeviceN color spaces are similar to Separation color
// spaces, except they can contain an arbitrary number of color components.
//
// Format: [/DeviceN names alternateSpace tintTransform]
//     or: [/DeviceN names alternateSpace tintTransform attributes]
type PdfColorspaceDeviceN struct {
	ColorantNames  *_cde.PdfObjectArray
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	Attributes     *PdfColorspaceDeviceNAttributes
	_cdac          *_cde.PdfIndirectObject
}

// NewPdfActionGoTo returns a new "go to" action.
func NewPdfActionGoTo() *PdfActionGoTo {
	_dc := NewPdfAction()
	_ddce := &PdfActionGoTo{}
	_ddce.PdfAction = _dc
	_dc.SetContext(_ddce)
	return _ddce
}

// PdfAnnotationRedact represents Redact annotations.
// (Section 12.5.6.23).
type PdfAnnotationRedact struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints  _cde.PdfObject
	IC          _cde.PdfObject
	RO          _cde.PdfObject
	OverlayText _cde.PdfObject
	Repeat      _cde.PdfObject
	DA          _cde.PdfObject
	Q           _cde.PdfObject
}

func _caag(_begfb *_cde.PdfObjectDictionary) {
	_ffbf, _ecgf := _cde.GetArray(_begfb.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
	_dfde, _ddbce := _cde.GetIntVal(_begfb.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
	_abcb, _bcbee := _cde.GetIntVal(_begfb.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
	if _ecgf && _ddbce && _bcbee {
		_dafad := _ffbf.Len()
		if _dafad != _abcb-_dfde+1 {
			_ad.Log.Debug("\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076\u002c\u0020\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u003a\u0020\u0025\u0076", _dafad, _abcb)
			_feaef := _cde.PdfObjectInteger(_dfde + _dafad - 1)
			_begfb.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", &_feaef)
		}
	}
}

// GetColorspaceByName returns the colorspace with the specified name from the page resources.
func (_dfebeg *PdfPageResources) GetColorspaceByName(keyName _cde.PdfObjectName) (PdfColorspace, bool) {
	_gdbd, _adgad := _dfebeg.GetColorspaces()
	if _adgad != nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _adgad)
		return nil, false
	}
	if _gdbd == nil {
		return nil, false
	}
	_edfef, _agcg := _gdbd.Colorspaces[string(keyName)]
	if !_agcg {
		return nil, false
	}
	return _edfef, true
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_afdc *PdfShadingType5) ToPdfObject() _cde.PdfObject {
	_afdc.PdfShading.ToPdfObject()
	_dbbb, _fdeae := _afdc.getShadingDict()
	if _fdeae != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _afdc.BitsPerCoordinate != nil {
		_dbbb.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _afdc.BitsPerCoordinate)
	}
	if _afdc.BitsPerComponent != nil {
		_dbbb.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _afdc.BitsPerComponent)
	}
	if _afdc.VerticesPerRow != nil {
		_dbbb.Set("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077", _afdc.VerticesPerRow)
	}
	if _afdc.Decode != nil {
		_dbbb.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _afdc.Decode)
	}
	if _afdc.Function != nil {
		if len(_afdc.Function) == 1 {
			_dbbb.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _afdc.Function[0].ToPdfObject())
		} else {
			_dadd := _cde.MakeArray()
			for _, _gddbf := range _afdc.Function {
				_dadd.Append(_gddbf.ToPdfObject())
			}
			_dbbb.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dadd)
		}
	}
	return _afdc._dffg
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the L, A and B components of the color.
func (_fgcf *PdfColorspaceLab) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bgee, _fdgf := _cde.GetNumbersAsFloat(objects)
	if _fdgf != nil {
		return nil, _fdgf
	}
	return _fgcf.ColorFromFloats(_bgee)
}
func _cbbf(_aeedc *_cde.PdfObjectDictionary, _egdbe *fontCommon) (*pdfFontType0, error) {
	_bdcbe, _gbce := _cde.GetArray(_aeedc.Get("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073"))
	if !_gbce {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006cid\u0020\u0044\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073\u0020\u002d\u0020\u006e\u006f\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079 \u0025\u0073", _egdbe)
		return nil, _cde.ErrRangeError
	}
	if _bdcbe.Len() != 1 {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0041\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020(%\u0064\u0029", _bdcbe.Len())
		return nil, _cde.ErrRangeError
	}
	_ccgb, _acbgf := _dgefe(_bdcbe.Get(0), false)
	if _acbgf != nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046a\u0069\u006c\u0065d \u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u0020\u0066\u006f\u006et\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076 \u0025\u0073", _acbgf, _egdbe)
		return nil, _acbgf
	}
	_adabc := _cfbaf(_egdbe)
	_adabc.DescendantFont = _ccgb
	_gafgf, _gbce := _cde.GetNameVal(_aeedc.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if _gbce {
		if _gafgf == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048" || _gafgf == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056" {
			_adabc._cffef = _gc.NewIdentityTextEncoder(_gafgf)
		} else if _fb.IsPredefinedCMap(_gafgf) {
			_adabc._dcdga, _acbgf = _fb.LoadPredefinedCMap(_gafgf)
			if _acbgf != nil {
				_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _gafgf, _acbgf)
			}
		} else {
			_ad.Log.Debug("\u0055\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0063\u006da\u0070\u0020\u0025\u0071", _gafgf)
		}
	}
	if _eedee := _ccgb.baseFields()._ggebg; _eedee != nil {
		if _gafee := _eedee.Name(); _gafee == "\u0041d\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002d\u0055\u0043\u0053\u0032" || _gafee == "\u0041\u0064\u006f\u0062\u0065\u002d\u0047\u0042\u0031-\u0055\u0043\u0053\u0032" || _gafee == "\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031-\u0055\u0043\u0053\u0032" || _gafee == "\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006f\u0072\u0065\u0061\u0031-\u0055\u0043\u0053\u0032" {
			_adabc._cffef = _gc.NewCMapEncoder(_gafgf, _adabc._dcdga, _eedee)
		}
	}
	return _adabc, nil
}
func _fegg(_gecgc *_cde.PdfObjectDictionary, _aeegg *fontCommon) (*pdfCIDFontType0, error) {
	if _aeegg._dcbc != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" {
		_ad.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0030\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _aeegg)
		return nil, _cde.ErrRangeError
	}
	_egagd := _agad(_aeegg)
	_gada, _agba := _cde.GetDict(_gecgc.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_agba {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _aeegg)
		return nil, ErrRequiredAttributeMissing
	}
	_egagd.CIDSystemInfo = _gada
	_egagd.DW = _gecgc.Get("\u0044\u0057")
	_egagd.W = _gecgc.Get("\u0057")
	_egagd.DW2 = _gecgc.Get("\u0044\u0057\u0032")
	_egagd.W2 = _gecgc.Get("\u0057\u0032")
	_egagd._bdega = 1000.0
	if _abga, _acbbb := _cde.GetNumberAsFloat(_egagd.DW); _acbbb == nil {
		_egagd._bdega = _abga
	}
	_dbca, _gcde := _fgecb(_egagd.W)
	if _gcde != nil {
		return nil, _gcde
	}
	if _dbca == nil {
		_dbca = map[_gc.CharCode]float64{}
	}
	_egagd._egfeb = _dbca
	return _egagd, nil
}

// ToGoImage converts the unidoc Image to a golang Image structure.
func (_edabd *Image) ToGoImage() (_gf.Image, error) {
	_ad.Log.Trace("\u0043\u006f\u006e\u0076er\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u006f\u0020\u0069\u006d\u0061g\u0065")
	_ffge, _dgecg := _ff.NewImage(int(_edabd.Width), int(_edabd.Height), int(_edabd.BitsPerComponent), _edabd.ColorComponents, _edabd.Data, _edabd._deegf, _edabd._aaafb)
	if _dgecg != nil {
		return nil, _dgecg
	}
	return _ffge, nil
}
func _eefgd(_gbdd map[_fe.GID]int, _dgag uint16) *_cde.PdfObjectArray {
	_fggd := &_cde.PdfObjectArray{}
	_eade := _fe.GID(_dgag)
	for _fafa := _fe.GID(0); _fafa < _eade; {
		_feadc, _gcad := _gbdd[_fafa]
		if !_gcad {
			_fafa++
			continue
		}
		_aggd := _fafa
		for _gcdf := _aggd + 1; _gcdf < _eade; _gcdf++ {
			if _aefe, _bfecd := _gbdd[_gcdf]; !_bfecd || _feadc != _aefe {
				break
			}
			_aggd = _gcdf
		}
		_fggd.Append(_cde.MakeInteger(int64(_fafa)))
		_fggd.Append(_cde.MakeInteger(int64(_aggd)))
		_fggd.Append(_cde.MakeInteger(int64(_feadc)))
		_fafa = _aggd + 1
	}
	return _fggd
}
func (_eafd *pdfCIDFontType2) getFontDescriptor() *PdfFontDescriptor { return _eafd._fagf }
func (_cbfgc *pdfFontType3) baseFields() *fontCommon                 { return &_cbfgc.fontCommon }

// Set applies flag fl to the flag's bitmask and returns the combined flag.
func (_bfcd FieldFlag) Set(fl FieldFlag) FieldFlag { return FieldFlag(_bfcd.Mask() | fl.Mask()) }

// NewPdfFilespec returns an initialized generic PDF filespec model.
func NewPdfFilespec() *PdfFilespec {
	_ccdcc := &PdfFilespec{}
	_ccdcc._bgac = _cde.MakeIndirectObject(_cde.MakeDict())
	return _ccdcc
}

// PdfColorLab represents a color in the L*, a*, b* 3 component colorspace.
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorLab [3]float64

func (_adfd *DSS) generateHashMap(_ddfbe []*_cde.PdfObjectStream) (map[string]*_cde.PdfObjectStream, error) {
	_bcbfd := map[string]*_cde.PdfObjectStream{}
	for _, _dcag := range _ddfbe {
		_fafde, _cbcf := _cde.DecodeStream(_dcag)
		if _cbcf != nil {
			return nil, _cbcf
		}
		_abbff, _cbcf := _dgaff(_fafde)
		if _cbcf != nil {
			return nil, _cbcf
		}
		_bcbfd[string(_abbff)] = _dcag
	}
	return _bcbfd, nil
}
func (_feaed *PdfReader) loadDSS() (*DSS, error) {
	if _feaed._aggcgb.GetCrypter() != nil && !_feaed._aggcgb.IsAuthenticated() {
		return nil, _ee.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_ceggd := _feaed._efabe.Get("\u0044\u0053\u0053")
	if _ceggd == nil {
		return nil, nil
	}
	_beaff, _ := _cde.GetIndirect(_ceggd)
	_ceggd = _cde.TraceToDirectObject(_ceggd)
	switch _gfbfe := _ceggd.(type) {
	case *_cde.PdfObjectNull:
		return nil, nil
	case *_cde.PdfObjectDictionary:
		return _gcfa(_beaff, _gfbfe)
	}
	return nil, _ee.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u0053\u0053 \u0065\u006e\u0074\u0072y \u0025\u0054", _ceggd)
}
func (_egfee *PdfWriter) checkLicense() error {
	return nil
}

// PdfAnnotationLine represents Line annotations.
// (Section 12.5.6.7).
type PdfAnnotationLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	L       _cde.PdfObject
	BS      _cde.PdfObject
	LE      _cde.PdfObject
	IC      _cde.PdfObject
	LL      _cde.PdfObject
	LLE     _cde.PdfObject
	Cap     _cde.PdfObject
	IT      _cde.PdfObject
	LLO     _cde.PdfObject
	CP      _cde.PdfObject
	Measure _cde.PdfObject
	CO      _cde.PdfObject
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_ceae pdfCIDFontType0) GetRuneMetrics(r rune) (_fe.CharMetrics, bool) {
	return _fe.CharMetrics{Wx: _ceae._bdega}, true
}

// PdfReader represents a PDF file reader. It is a frontend to the lower level parsing mechanism and provides
// a higher level access to work with PDF structure and information, such as the page structure etc.
type PdfReader struct {
	_aggcgb  *_cde.PdfParser
	_bgbage  _cde.PdfObject
	_geadg   *_cde.PdfIndirectObject
	_bdagc   *_cde.PdfObjectDictionary
	_gbfgg   []*_cde.PdfIndirectObject
	PageList []*PdfPage
	_bdaff   int
	_efabe   *_cde.PdfObjectDictionary
	_aegbb   *PdfOutlineTreeNode
	AcroForm *PdfAcroForm
	DSS      *DSS
	Rotate   *int64
	_ebbdb   *Permissions
	_egaba   map[*PdfReader]*PdfReader
	_fgfbe   []*PdfReader
	_bedfa   *modelManager
	_cdgee   bool
	_efbdd   map[_cde.PdfObject]struct{}
	_caecc   _f.ReadSeeker
	_gfbcg   string
	_gcegag  bool
	_cbbc    *ReaderOpts
	_cgebe   bool
}

func (_dddbe *PdfWriter) seekByName(_becdb _cde.PdfObject, _ddafe []string, _ffegd string) ([]_cde.PdfObject, error) {
	_ad.Log.Trace("\u0053\u0065\u0065\u006b\u0020\u0062\u0079\u0020\u006e\u0061\u006d\u0065.\u002e\u0020\u0025\u0054", _becdb)
	var _gdedg []_cde.PdfObject
	if _bgebc, _eefbd := _becdb.(*_cde.PdfIndirectObject); _eefbd {
		return _dddbe.seekByName(_bgebc.PdfObject, _ddafe, _ffegd)
	}
	if _dfaag, _caafg := _becdb.(*_cde.PdfObjectStream); _caafg {
		return _dddbe.seekByName(_dfaag.PdfObjectDictionary, _ddafe, _ffegd)
	}
	if _faaac, _geafd := _becdb.(*_cde.PdfObjectDictionary); _geafd {
		_ad.Log.Trace("\u0044\u0069\u0063\u0074")
		for _, _dfabc := range _faaac.Keys() {
			_bgceg := _faaac.Get(_dfabc)
			if string(_dfabc) == _ffegd {
				_gdedg = append(_gdedg, _bgceg)
			}
			for _, _dbdab := range _ddafe {
				if string(_dfabc) == _dbdab {
					_ad.Log.Trace("\u0046\u006f\u006c\u006c\u006f\u0077\u0020\u006b\u0065\u0079\u0020\u0025\u0073", _dbdab)
					_ffcda, _dedeb := _dddbe.seekByName(_bgceg, _ddafe, _ffegd)
					if _dedeb != nil {
						return _gdedg, _dedeb
					}
					_gdedg = append(_gdedg, _ffcda...)
					break
				}
			}
		}
		return _gdedg, nil
	}
	return _gdedg, nil
}

// ToPdfObject returns a PDF object representation of the outline destination.
func (_fdafd OutlineDest) ToPdfObject() _cde.PdfObject {
	if (_fdafd.PageObj == nil && _fdafd.Page < 0) || _fdafd.Mode == "" {
		return _cde.MakeNull()
	}
	_cegab := _cde.MakeArray()
	if _fdafd.PageObj != nil {
		_cegab.Append(_fdafd.PageObj)
	} else {
		_cegab.Append(_cde.MakeInteger(_fdafd.Page))
	}
	_cegab.Append(_cde.MakeName(_fdafd.Mode))
	switch _fdafd.Mode {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		_cegab.Append(_cde.MakeFloat(_fdafd.Y))
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		_cegab.Append(_cde.MakeFloat(_fdafd.X))
	case "\u0058\u0059\u005a":
		_cegab.Append(_cde.MakeFloat(_fdafd.X))
		_cegab.Append(_cde.MakeFloat(_fdafd.Y))
		_cegab.Append(_cde.MakeFloat(_fdafd.Zoom))
	default:
		_cegab.Set(1, _cde.MakeName("\u0046\u0069\u0074"))
	}
	return _cegab
}

// ColorFromFloats returns a new PdfColor based on input color components.
func (_fdaf *PdfColorspaceDeviceN) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != _fdaf.GetNumComponents() {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gcacc, _cfgcb := _fdaf.TintTransform.Evaluate(vals)
	if _cfgcb != nil {
		return nil, _cfgcb
	}
	_eggg, _cfgcb := _fdaf.AlternateSpace.ColorFromFloats(_gcacc)
	if _cfgcb != nil {
		return nil, _cfgcb
	}
	return _eggg, nil
}

// SetImageHandler sets the image handler used by the package.
func SetImageHandler(imgHandling ImageHandler) { ImageHandling = imgHandling }

// ToPdfObject returns the PDF representation of the tiling pattern.
func (_fccfg *PdfTilingPattern) ToPdfObject() _cde.PdfObject {
	_fccfg.PdfPattern.ToPdfObject()
	_dfdga := _fccfg.getDict()
	if _fccfg.PaintType != nil {
		_dfdga.Set("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e", _fccfg.PaintType)
	}
	if _fccfg.TilingType != nil {
		_dfdga.Set("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _fccfg.TilingType)
	}
	if _fccfg.BBox != nil {
		_dfdga.Set("\u0042\u0042\u006f\u0078", _fccfg.BBox.ToPdfObject())
	}
	if _fccfg.XStep != nil {
		_dfdga.Set("\u0058\u0053\u0074e\u0070", _fccfg.XStep)
	}
	if _fccfg.YStep != nil {
		_dfdga.Set("\u0059\u0053\u0074e\u0070", _fccfg.YStep)
	}
	if _fccfg.Resources != nil {
		_dfdga.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _fccfg.Resources.ToPdfObject())
	}
	if _fccfg.Matrix != nil {
		_dfdga.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _fccfg.Matrix)
	}
	return _fccfg._eecac
}

// NewPdfFieldSignature returns an initialized signature field.
func NewPdfFieldSignature(signature *PdfSignature) *PdfFieldSignature {
	_gfee := &PdfFieldSignature{}
	_gfee.PdfField = NewPdfField()
	_gfee.PdfField.SetContext(_gfee)
	_gfee.PdfAnnotationWidget = NewPdfAnnotationWidget()
	_gfee.PdfAnnotationWidget.SetContext(_gfee)
	_gfee.PdfAnnotationWidget._bddg = _gfee.PdfField._afgc
	_gfee.T = _cde.MakeString("")
	_gfee.F = _cde.MakeInteger(132)
	_gfee.V = signature
	return _gfee
}

// PdfColorspaceCalGray represents CalGray color space.
type PdfColorspaceCalGray struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      float64
	_ggffb     *_cde.PdfIndirectObject
}

func _deafd(_edbgf *fontCommon) *pdfFontType3 { return &pdfFontType3{fontCommon: *_edbgf} }
func (_feaa *pdfFontType0) subsetRegistered() error {
	_cggcc, _aaddb := _feaa.DescendantFont._gbcff.(*pdfCIDFontType2)
	if !_aaddb {
		_ad.Log.Debug("\u0046\u006fnt\u0020\u006e\u006ft\u0020\u0073\u0075\u0070por\u0074ed\u0020\u0066\u006f\u0072\u0020\u0073\u0075bs\u0065\u0074\u0074\u0069\u006e\u0067\u0020%\u0054", _feaa.DescendantFont)
		return nil
	}
	if _cggcc == nil {
		return nil
	}
	if _cggcc._fagf == nil {
		_ad.Log.Debug("\u004d\u0069\u0073si\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return nil
	}
	if _feaa._cffef == nil {
		_ad.Log.Debug("\u004e\u006f\u0020e\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0073\u0075\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0067\u006e\u006f\u0072\u0065\u0064")
		return nil
	}
	_gedgd, _aaddb := _cde.GetStream(_cggcc._fagf.FontFile2)
	if !_aaddb {
		_ad.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _ceg.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_efeb, _bfeca := _cde.DecodeStream(_gedgd)
	if _bfeca != nil {
		_ad.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bfeca)
		return _bfeca
	}
	_ceefd, _bfeca := _ef.Parse(_ede.NewReader(_efeb))
	if _bfeca != nil {
		_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_gedgd.Stream))
		return _bfeca
	}
	var _aaccc []rune
	var _cebbc *_ef.Font
	switch _eeage := _feaa._cffef.(type) {
	case *_gc.TrueTypeFontEncoder:
		_aaccc = _eeage.RegisteredRunes()
		_cebbc, _bfeca = _ceefd.SubsetKeepRunes(_aaccc)
		if _bfeca != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bfeca)
			return _bfeca
		}
		_eeage.SubsetRegistered()
	case *_gc.IdentityEncoder:
		_aaccc = _eeage.RegisteredRunes()
		_dfdc := make([]_ef.GlyphIndex, len(_aaccc))
		for _gafaa, _cdacg := range _aaccc {
			_dfdc[_gafaa] = _ef.GlyphIndex(_cdacg)
		}
		_cebbc, _bfeca = _ceefd.SubsetKeepIndices(_dfdc)
		if _bfeca != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bfeca)
			return _bfeca
		}
	case _gc.SimpleEncoder:
		_fbffc := _eeage.Charcodes()
		for _, _aabaf := range _fbffc {
			_gcece, _fffb := _eeage.CharcodeToRune(_aabaf)
			if !_fffb {
				_ad.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0020\u0074\u006f \u0072\u0075\u006e\u0065\u003a\u0020\u0025\u0064", _aabaf)
				continue
			}
			_aaccc = append(_aaccc, _gcece)
		}
	default:
		return _ee.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069\u006eg\u003a\u0020\u0025\u0054", _feaa._cffef)
	}
	var _dffcd _ede.Buffer
	_bfeca = _cebbc.Write(&_dffcd)
	if _bfeca != nil {
		_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bfeca)
		return _bfeca
	}
	if _feaa._ggebg != nil {
		_fcgea := make(map[_fb.CharCode]rune, len(_aaccc))
		for _, _cbcga := range _aaccc {
			_afcg, _bdfga := _feaa._cffef.RuneToCharcode(_cbcga)
			if !_bdfga {
				continue
			}
			_fcgea[_fb.CharCode(_afcg)] = _cbcga
		}
		_feaa._ggebg = _fb.NewToUnicodeCMap(_fcgea)
	}
	_gedgd, _bfeca = _cde.MakeStream(_dffcd.Bytes(), _cde.NewFlateEncoder())
	if _bfeca != nil {
		_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bfeca)
		return _bfeca
	}
	_gedgd.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _cde.MakeInteger(int64(_dffcd.Len())))
	if _cegaf, _bdgf := _cde.GetStream(_cggcc._fagf.FontFile2); _bdgf {
		*_cegaf = *_gedgd
	} else {
		_cggcc._fagf.FontFile2 = _gedgd
	}
	_dbggg := _abaf()
	if len(_feaa._eeab) > 0 {
		_feaa._eeab = _dede(_feaa._eeab, _dbggg)
	}
	if len(_cggcc._eeab) > 0 {
		_cggcc._eeab = _dede(_cggcc._eeab, _dbggg)
	}
	if len(_feaa._fddeb) > 0 {
		_feaa._fddeb = _dede(_feaa._fddeb, _dbggg)
	}
	if _cggcc._fagf != nil {
		_egacd, _ggbg := _cde.GetName(_cggcc._fagf.FontName)
		if _ggbg && len(_egacd.String()) > 0 {
			_aagf := _dede(_egacd.String(), _dbggg)
			_cggcc._fagf.FontName = _cde.MakeName(_aagf)
		}
	}
	return nil
}
func (_dgac *PdfReader) newPdfAnnotationPopupFromDict(_gbc *_cde.PdfObjectDictionary) (*PdfAnnotationPopup, error) {
	_ffcgf := PdfAnnotationPopup{}
	_ffcgf.Parent = _gbc.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	_ffcgf.Open = _gbc.Get("\u004f\u0070\u0065\u006e")
	return &_ffcgf, nil
}
func (_bcadg *PdfWriter) writeObject(_bffcf int, _feaf _cde.PdfObject) {
	_ad.Log.Trace("\u0057\u0072\u0069\u0074\u0065\u0020\u006f\u0062\u006a \u0023\u0025\u0064\u000a", _bffcf)
	if _agedg, _gbecd := _feaf.(*_cde.PdfIndirectObject); _gbecd {
		_bcadg._gfdac[_bffcf] = crossReference{Type: 1, Offset: _bcadg._eddbc, Generation: _agedg.GenerationNumber}
		_facge := _ee.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _bffcf)
		if _eaedc, _bedaf := _agedg.PdfObject.(*pdfSignDictionary); _bedaf {
			_eaedc._aafab = _bcadg._eddbc + int64(len(_facge))
		}
		if _agedg.PdfObject == nil {
			_ad.Log.Debug("E\u0072\u0072\u006fr\u003a\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0027\u0073\u0020\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020b\u0065\u0020\u006e\u0069l\u0020\u002d\u0020\u0073e\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063t\u004e\u0075\u006c\u006c")
			_agedg.PdfObject = _cde.MakeNull()
		}
		_facge += _agedg.PdfObject.WriteString()
		_facge += "\u000a\u0065\u006e\u0064\u006f\u0062\u006a\u000a"
		_bcadg.writeString(_facge)
		return
	}
	if _dgagab, _ccgde := _feaf.(*_cde.PdfObjectStream); _ccgde {
		_bcadg._gfdac[_bffcf] = crossReference{Type: 1, Offset: _bcadg._eddbc, Generation: _dgagab.GenerationNumber}
		_ebegf := _ee.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _bffcf)
		_ebegf += _dgagab.PdfObjectDictionary.WriteString()
		_ebegf += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_bcadg.writeString(_ebegf)
		_bcadg.writeBytes(_dgagab.Stream)
		_bcadg.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	if _ddfce, _daecd := _feaf.(*_cde.PdfObjectStreams); _daecd {
		_bcadg._gfdac[_bffcf] = crossReference{Type: 1, Offset: _bcadg._eddbc, Generation: _ddfce.GenerationNumber}
		_dbfbd := _ee.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _bffcf)
		var _fgggd []string
		var _bdcbg string
		var _dffbd int64
		for _bafdg, _dagbc := range _ddfce.Elements() {
			_ebfaf, _adbee := _dagbc.(*_cde.PdfIndirectObject)
			if !_adbee {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065am\u0073 \u004e\u0020\u0025\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006es\u0020\u006e\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u0070\u0064\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0076", _bffcf, _dagbc)
				continue
			}
			_fgfef := _ebfaf.PdfObject.WriteString() + "\u0020"
			_bdcbg = _bdcbg + _fgfef
			_fgggd = append(_fgggd, _ee.Sprintf("\u0025\u0064\u0020%\u0064", _ebfaf.ObjectNumber, _dffbd))
			_bcadg._gfdac[int(_ebfaf.ObjectNumber)] = crossReference{Type: 2, ObjectNumber: _bffcf, Index: _bafdg}
			_dffbd = _dffbd + int64(len([]byte(_fgfef)))
		}
		_ebcec := _dac.Join(_fgggd, "\u0020") + "\u0020"
		_afbbe := _cde.NewFlateEncoder()
		_gddgc := _afbbe.MakeStreamDict()
		_gddgc.Set(_cde.PdfObjectName("\u0054\u0079\u0070\u0065"), _cde.MakeName("\u004f\u0062\u006a\u0053\u0074\u006d"))
		_fbabd := int64(_ddfce.Len())
		_gddgc.Set(_cde.PdfObjectName("\u004e"), _cde.MakeInteger(_fbabd))
		_faaff := int64(len(_ebcec))
		_gddgc.Set(_cde.PdfObjectName("\u0046\u0069\u0072s\u0074"), _cde.MakeInteger(_faaff))
		_ggbac, _ := _afbbe.EncodeBytes([]byte(_ebcec + _bdcbg))
		_bdcbag := int64(len(_ggbac))
		_gddgc.Set(_cde.PdfObjectName("\u004c\u0065\u006e\u0067\u0074\u0068"), _cde.MakeInteger(_bdcbag))
		_dbfbd += _gddgc.WriteString()
		_dbfbd += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_bcadg.writeString(_dbfbd)
		_bcadg.writeBytes(_ggbac)
		_bcadg.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	_bcadg.writeString(_feaf.WriteString())
}

// NewPdfColorspaceDeviceCMYK returns a new CMYK32 colorspace object.
func NewPdfColorspaceDeviceCMYK() *PdfColorspaceDeviceCMYK { return &PdfColorspaceDeviceCMYK{} }
func (_fcaf *PdfReader) newPdfAnnotationFileAttachmentFromDict(_gdcd *_cde.PdfObjectDictionary) (*PdfAnnotationFileAttachment, error) {
	_efad := PdfAnnotationFileAttachment{}
	_bcc, _ebcga := _fcaf.newPdfAnnotationMarkupFromDict(_gdcd)
	if _ebcga != nil {
		return nil, _ebcga
	}
	_efad.PdfAnnotationMarkup = _bcc
	_efad.FS = _gdcd.Get("\u0046\u0053")
	_efad.Name = _gdcd.Get("\u004e\u0061\u006d\u0065")
	return &_efad, nil
}

// NewPdfColorspaceICCBased returns a new ICCBased colorspace object.
func NewPdfColorspaceICCBased(N int) (*PdfColorspaceICCBased, error) {
	_gfecg := &PdfColorspaceICCBased{}
	if N != 1 && N != 3 && N != 4 {
		return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u0031/\u0033\u002f\u0034\u0029")
	}
	_gfecg.N = N
	return _gfecg, nil
}
func (_bdgae *PdfWriter) addObject(_faed _cde.PdfObject) bool {
	_adad := _bdgae.hasObject(_faed)
	if !_adad {
		_bgabf := _cde.ResolveReferencesDeep(_faed, _bdgae._gbdfb)
		if _bgabf != nil {
			_ad.Log.Debug("E\u0052R\u004f\u0052\u003a\u0020\u0025\u0076\u0020\u002d \u0073\u006b\u0069\u0070pi\u006e\u0067", _bgabf)
		}
		_bdgae._egbccc = append(_bdgae._egbccc, _faed)
		_bdgae._bccde[_faed] = struct{}{}
		return true
	}
	return false
}
func _deec(_aadf _cde.PdfObject) (*PdfColorspaceCalGray, error) {
	_cdb := NewPdfColorspaceCalGray()
	if _ddggd, _fbcb := _aadf.(*_cde.PdfIndirectObject); _fbcb {
		_cdb._ggffb = _ddggd
	}
	_aadf = _cde.TraceToDirectObject(_aadf)
	_dfegc, _gaede := _aadf.(*_cde.PdfObjectArray)
	if !_gaede {
		return nil, _ee.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _dfegc.Len() != 2 {
		return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0061\u006cG\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_aadf = _cde.TraceToDirectObject(_dfegc.Get(0))
	_cbdf, _gaede := _aadf.(*_cde.PdfObjectName)
	if !_gaede {
		return nil, _ee.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u0020\u006e\u0061m\u0065\u0020\u006e\u006f\u0074\u0020\u0061 \u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_cbdf != "\u0043a\u006c\u0047\u0072\u0061\u0079" {
		return nil, _ee.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0043\u0061\u006c\u0047\u0072\u0061\u0079 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_aadf = _cde.TraceToDirectObject(_dfegc.Get(1))
	_aeaa, _gaede := _aadf.(*_cde.PdfObjectDictionary)
	if !_gaede {
		return nil, _ee.Errorf("\u0043\u0061lG\u0072\u0061\u0079 \u0064\u0069\u0063\u0074 no\u0074 a\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u006f\u0062\u006a\u0065c\u0074")
	}
	_aadf = _aeaa.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_aadf = _cde.TraceToDirectObject(_aadf)
	_cfgg, _gaede := _aadf.(*_cde.PdfObjectArray)
	if !_gaede {
		return nil, _ee.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020W\u0068\u0069\u0074e\u0050o\u0069\u006e\u0074")
	}
	if _cfgg.Len() != 3 {
		return nil, _ee.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
	}
	_acde, _ddbc := _cfgg.GetAsFloat64Slice()
	if _ddbc != nil {
		return nil, _ddbc
	}
	_cdb.WhitePoint = _acde
	_aadf = _aeaa.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _aadf != nil {
		_aadf = _cde.TraceToDirectObject(_aadf)
		_cdfac, _aadfg := _aadf.(*_cde.PdfObjectArray)
		if !_aadfg {
			return nil, _ee.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020B\u006c\u0061\u0063k\u0050o\u0069\u006e\u0074")
		}
		if _cdfac.Len() != 3 {
			return nil, _ee.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061c\u006b\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
		}
		_fdbf, _cbfa := _cdfac.GetAsFloat64Slice()
		if _cbfa != nil {
			return nil, _cbfa
		}
		_cdb.BlackPoint = _fdbf
	}
	_aadf = _aeaa.Get("\u0047\u0061\u006dm\u0061")
	if _aadf != nil {
		_aadf = _cde.TraceToDirectObject(_aadf)
		_febb, _baaa := _cde.GetNumberAsFloat(_aadf)
		if _baaa != nil {
			return nil, _ee.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0067\u0061\u006d\u006d\u0061\u0020\u006e\u006ft\u0020\u0061\u0020n\u0075m\u0062\u0065\u0072")
		}
		_cdb.Gamma = _febb
	}
	return _cdb, nil
}
func (_gccgg *PdfShading) getShadingDict() (*_cde.PdfObjectDictionary, error) {
	_cacdcf := _gccgg._dffg
	if _adfg, _cccf := _cacdcf.(*_cde.PdfIndirectObject); _cccf {
		_gbagaf, _aegf := _adfg.PdfObject.(*_cde.PdfObjectDictionary)
		if !_aegf {
			return nil, _cde.ErrTypeError
		}
		return _gbagaf, nil
	} else if _ddedc, _cffedg := _cacdcf.(*_cde.PdfObjectStream); _cffedg {
		return _ddedc.PdfObjectDictionary, nil
	} else if _gdebc, _beccb := _cacdcf.(*_cde.PdfObjectDictionary); _beccb {
		return _gdebc, nil
	} else {
		_ad.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0063\u0063\u0065s\u0073\u0020\u0073\u0068\u0061\u0064\u0069n\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079")
		return nil, _cde.ErrTypeError
	}
}
func (_eecc *PdfWriter) setHashIDs(_fdad _g.Hash) error {
	_agdff := _fdad.Sum(nil)
	if _eecc._gfegf == "" {
		_eecc._gfegf = _ed.EncodeToString(_agdff[:8])
	}
	_eecc.setDocumentIDs(_eecc._gfegf, _ed.EncodeToString(_agdff[8:]))
	return nil
}

// ImageToRGB converts CalRGB colorspace image to RGB and returns the result.
func (_fdgb *PdfColorspaceCalRGB) ImageToRGB(img Image) (Image, error) {
	_cage := _cae.NewReader(img.getBase())
	_facg := _ff.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_fgbc := _cae.NewWriter(_facg)
	_dcafc := _ced.Pow(2, float64(img.BitsPerComponent)) - 1
	_ccec := make([]uint32, 3)
	var (
		_cedb                                     error
		_bdcd, _eeag, _bdgb, _fcde, _cfgc, _dbead float64
	)
	for {
		_cedb = _cage.ReadSamples(_ccec)
		if _cedb == _f.EOF {
			break
		} else if _cedb != nil {
			return img, _cedb
		}
		_bdcd = float64(_ccec[0]) / _dcafc
		_eeag = float64(_ccec[1]) / _dcafc
		_bdgb = float64(_ccec[2]) / _dcafc
		_fcde = _fdgb.Matrix[0]*_ced.Pow(_bdcd, _fdgb.Gamma[0]) + _fdgb.Matrix[3]*_ced.Pow(_eeag, _fdgb.Gamma[1]) + _fdgb.Matrix[6]*_ced.Pow(_bdgb, _fdgb.Gamma[2])
		_cfgc = _fdgb.Matrix[1]*_ced.Pow(_bdcd, _fdgb.Gamma[0]) + _fdgb.Matrix[4]*_ced.Pow(_eeag, _fdgb.Gamma[1]) + _fdgb.Matrix[7]*_ced.Pow(_bdgb, _fdgb.Gamma[2])
		_dbead = _fdgb.Matrix[2]*_ced.Pow(_bdcd, _fdgb.Gamma[0]) + _fdgb.Matrix[5]*_ced.Pow(_eeag, _fdgb.Gamma[1]) + _fdgb.Matrix[8]*_ced.Pow(_bdgb, _fdgb.Gamma[2])
		_bdcd = 3.240479*_fcde + -1.537150*_cfgc + -0.498535*_dbead
		_eeag = -0.969256*_fcde + 1.875992*_cfgc + 0.041556*_dbead
		_bdgb = 0.055648*_fcde + -0.204043*_cfgc + 1.057311*_dbead
		_bdcd = _ced.Min(_ced.Max(_bdcd, 0), 1.0)
		_eeag = _ced.Min(_ced.Max(_eeag, 0), 1.0)
		_bdgb = _ced.Min(_ced.Max(_bdgb, 0), 1.0)
		_ccec[0] = uint32(_bdcd * _dcafc)
		_ccec[1] = uint32(_eeag * _dcafc)
		_ccec[2] = uint32(_bdgb * _dcafc)
		if _cedb = _fgbc.WriteSamples(_ccec); _cedb != nil {
			return img, _cedb
		}
	}
	return _bddb(&_facg), nil
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

// NewOutline returns a new outline instance.
func NewOutline() *Outline { return &Outline{} }

// PdfFieldButton represents a button field which includes push buttons, checkboxes, and radio buttons.
type PdfFieldButton struct {
	*PdfField
	Opt   *_cde.PdfObjectArray
	_dcbd *Image
}

// GetPdfName returns the PDF name used to indicate the border style.
// (Table 166 p. 395).
func (_cgcg *BorderStyle) GetPdfName() string {
	switch *_cgcg {
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

// GetNamedDestinations returns the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_cbcge *PdfReader) GetNamedDestinations() (_cde.PdfObject, error) {
	_dfdgf := _cde.ResolveReference(_cbcge._efabe.Get("\u0044\u0065\u0073t\u0073"))
	if _dfdgf == nil {
		return nil, nil
	}
	if !_cbcge._cdgee {
		_dcedf := _cbcge.traverseObjectData(_dfdgf)
		if _dcedf != nil {
			return nil, _dcedf
		}
	}
	return _dfdgf, nil
}

// NewPdfColorspaceDeviceN returns an initialized PdfColorspaceDeviceN.
func NewPdfColorspaceDeviceN() *PdfColorspaceDeviceN {
	_bcagb := &PdfColorspaceDeviceN{}
	return _bcagb
}

// GetShadingByName gets the shading specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_adda *PdfPageResources) GetShadingByName(keyName _cde.PdfObjectName) (*PdfShading, bool) {
	if _adda.Shading == nil {
		return nil, false
	}
	_bgffc, _fbadg := _cde.TraceToDirectObject(_adda.Shading).(*_cde.PdfObjectDictionary)
	if !_fbadg {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0068\u0061d\u0069\u006e\u0067\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _adda.Shading)
		return nil, false
	}
	if _cedcf := _bgffc.Get(keyName); _cedcf != nil {
		_ecdaf, _ecebbb := _bfbdc(_cedcf)
		if _ecebbb != nil {
			_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0073\u0068\u0061\u0064\u0069\u006eg\u003a\u0020\u0025\u0076", _ecebbb)
			return nil, false
		}
		return _ecdaf, true
	}
	return nil, false
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
	DiffResults *_cdc.DiffResults

	// GeneralizedTime is the time at which the time-stamp token has been created by the TSA (RFC 3161).
	GeneralizedTime _ce.Time
}

// ToPdfObject returns a stream object.
func (_fcbde *XObjectForm) ToPdfObject() _cde.PdfObject {
	_bfegc := _fcbde._ecfbd
	_fadbd := _bfegc.PdfObjectDictionary
	if _fcbde.Filter != nil {
		_fadbd = _fcbde.Filter.MakeStreamDict()
		_bfegc.PdfObjectDictionary = _fadbd
	}
	_fadbd.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_fadbd.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0046\u006f\u0072\u006d"))
	_fadbd.SetIfNotNil("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065", _fcbde.FormType)
	_fadbd.SetIfNotNil("\u0042\u0042\u006f\u0078", _fcbde.BBox)
	_fadbd.SetIfNotNil("\u004d\u0061\u0074\u0072\u0069\u0078", _fcbde.Matrix)
	if _fcbde.Resources != nil {
		_fadbd.SetIfNotNil("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _fcbde.Resources.ToPdfObject())
	}
	_fadbd.SetIfNotNil("\u0047\u0072\u006fu\u0070", _fcbde.Group)
	_fadbd.SetIfNotNil("\u0052\u0065\u0066", _fcbde.Ref)
	_fadbd.SetIfNotNil("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061", _fcbde.MetaData)
	_fadbd.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _fcbde.PieceInfo)
	_fadbd.SetIfNotNil("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _fcbde.LastModified)
	_fadbd.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _fcbde.StructParent)
	_fadbd.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _fcbde.StructParents)
	_fadbd.SetIfNotNil("\u004f\u0050\u0049", _fcbde.OPI)
	_fadbd.SetIfNotNil("\u004f\u0043", _fcbde.OC)
	_fadbd.SetIfNotNil("\u004e\u0061\u006d\u0065", _fcbde.Name)
	_fadbd.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cde.MakeInteger(int64(len(_fcbde.Stream))))
	_bfegc.Stream = _fcbde.Stream
	return _bfegc
}

// Add appends a top level outline item to the outline.
func (_gbbee *Outline) Add(item *OutlineItem) { _gbbee.Entries = append(_gbbee.Entries, item) }

// ToPdfObject returns the PDF representation of the DSS dictionary.
func (_bfbc *DSS) ToPdfObject() _cde.PdfObject {
	_daccd := _bfbc._gbgda.PdfObject.(*_cde.PdfObjectDictionary)
	_daccd.Clear()
	_dace := _cde.MakeDict()
	for _fdff, _afcfa := range _bfbc.VRI {
		_dace.Set(*_cde.MakeName(_fdff), _afcfa.ToPdfObject())
	}
	_daccd.SetIfNotNil("\u0043\u0065\u0072t\u0073", _dbgae(_bfbc.Certs))
	_daccd.SetIfNotNil("\u004f\u0043\u0053P\u0073", _dbgae(_bfbc.OCSPs))
	_daccd.SetIfNotNil("\u0043\u0052\u004c\u0073", _dbgae(_bfbc.CRLs))
	_daccd.Set("\u0056\u0052\u0049", _dace)
	return _bfbc._gbgda
}
func (_gbdfe *PdfWriter) hasObject(_eddcg _cde.PdfObject) bool {
	_, _ecbgg := _gbdfe._bccde[_eddcg]
	return _ecbgg
}

// PdfActionSound represents a sound action.
type PdfActionSound struct {
	*PdfAction
	Sound       _cde.PdfObject
	Volume      _cde.PdfObject
	Synchronous _cde.PdfObject
	Repeat      _cde.PdfObject
	Mix         _cde.PdfObject
}

// NewPdfColorspaceCalRGB returns a new CalRGB colorspace object.
func NewPdfColorspaceCalRGB() *PdfColorspaceCalRGB {
	_daae := &PdfColorspaceCalRGB{}
	_daae.BlackPoint = []float64{0.0, 0.0, 0.0}
	_daae.Gamma = []float64{1.0, 1.0, 1.0}
	_daae.Matrix = []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	return _daae
}
func _aabeb(_adaga _cde.PdfObject) {
	_ad.Log.Debug("\u006f\u0062\u006a\u003a\u0020\u0025\u0054\u0020\u0025\u0073", _adaga, _adaga.String())
	if _bagda, _bgbed := _adaga.(*_cde.PdfObjectStream); _bgbed {
		_aagae, _bafab := _cde.DecodeStream(_bagda)
		if _bafab != nil {
			_ad.Log.Debug("\u0045r\u0072\u006f\u0072\u003a\u0020\u0025v", _bafab)
			return
		}
		_ad.Log.Debug("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _aagae)
	} else if _gfdbgg, _eadgcc := _adaga.(*_cde.PdfIndirectObject); _eadgcc {
		_ad.Log.Debug("\u0025\u0054\u0020%\u0076", _gfdbgg.PdfObject, _gfdbgg.PdfObject)
		_ad.Log.Debug("\u0025\u0073", _gfdbgg.PdfObject.String())
	}
}

// ToPdfObject implements interface PdfModel.
func (_dfa *PdfActionLaunch) ToPdfObject() _cde.PdfObject {
	_dfa.PdfAction.ToPdfObject()
	_cab := _dfa._bc
	_gdc := _cab.PdfObject.(*_cde.PdfObjectDictionary)
	_gdc.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeLaunch)))
	if _dfa.F != nil {
		_gdc.Set("\u0046", _dfa.F.ToPdfObject())
	}
	_gdc.SetIfNotNil("\u0057\u0069\u006e", _dfa.Win)
	_gdc.SetIfNotNil("\u004d\u0061\u0063", _dfa.Mac)
	_gdc.SetIfNotNil("\u0055\u006e\u0069\u0078", _dfa.Unix)
	_gdc.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _dfa.NewWindow)
	return _cab
}

// SetName sets the `Name` field of the signature.
func (_ggfbd *PdfSignature) SetName(name string) { _ggfbd.Name = _cde.MakeString(name) }
func (_eegb *PdfReader) newPdfActionGotoRFromDict(_abbb *_cde.PdfObjectDictionary) (*PdfActionGoToR, error) {
	_bbeg, _adbe := _beed(_abbb.Get("\u0046"))
	if _adbe != nil {
		return nil, _adbe
	}
	return &PdfActionGoToR{D: _abbb.Get("\u0044"), NewWindow: _abbb.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _bbeg}, nil
}

// String returns a string that describes `base`.
func (_fefa fontCommon) String() string {
	return _ee.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0073\u007d", _fefa.coreString())
}
func (_gfaggb *LTV) getOCSPs(_eegfg []*_bg.Certificate, _egaff map[string]*_bg.Certificate) ([][]byte, error) {
	_egage := make([][]byte, 0, len(_eegfg))
	for _, _fadd := range _eegfg {
		for _, _feecb := range _fadd.OCSPServer {
			if _gfaggb.CertClient.IsCA(_fadd) {
				continue
			}
			_bacbe, _eeafg := _egaff[_fadd.Issuer.CommonName]
			if !_eeafg {
				_ad.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _cdaa, _fdac := _gfaggb.OCSPClient.MakeRequest(_feecb, _fadd, _bacbe)
			if _fdac != nil {
				_ad.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _fdac)
				continue
			}
			_egage = append(_egage, _cdaa)
		}
	}
	return _egage, nil
}

// PageCallback callback function used in page loading
// that could be used to modify the page content.
//
// Deprecated: will be removed in v4. Use PageProcessCallback instead.
type PageCallback func(_dfged int, _bade *PdfPage)

// ToPdfObject returns a stream object.
func (_egfbd *XObjectImage) ToPdfObject() _cde.PdfObject {
	_egcff := _egfbd._bbaed
	_ddbda := _egcff.PdfObjectDictionary
	if _egfbd.Filter != nil {
		_ddbda = _egfbd.Filter.MakeStreamDict()
		_egcff.PdfObjectDictionary = _ddbda
	}
	_ddbda.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_ddbda.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0049\u006d\u0061g\u0065"))
	_ddbda.Set("\u0057\u0069\u0064t\u0068", _cde.MakeInteger(*(_egfbd.Width)))
	_ddbda.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _cde.MakeInteger(*(_egfbd.Height)))
	if _egfbd.BitsPerComponent != nil {
		_ddbda.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _cde.MakeInteger(*(_egfbd.BitsPerComponent)))
	}
	if _egfbd.ColorSpace != nil {
		_ddbda.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _egfbd.ColorSpace.ToPdfObject())
	}
	_ddbda.SetIfNotNil("\u0049\u006e\u0074\u0065\u006e\u0074", _egfbd.Intent)
	_ddbda.SetIfNotNil("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k", _egfbd.ImageMask)
	_ddbda.SetIfNotNil("\u004d\u0061\u0073\u006b", _egfbd.Mask)
	_dgfag := _ddbda.Get("\u0044\u0065\u0063\u006f\u0064\u0065") != nil
	if _egfbd.Decode == nil && _dgfag {
		_ddbda.Remove("\u0044\u0065\u0063\u006f\u0064\u0065")
	} else if _egfbd.Decode != nil {
		_ddbda.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _egfbd.Decode)
	}
	_ddbda.SetIfNotNil("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _egfbd.Interpolate)
	_ddbda.SetIfNotNil("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073", _egfbd.Alternatives)
	_ddbda.SetIfNotNil("\u0053\u004d\u0061s\u006b", _egfbd.SMask)
	_ddbda.SetIfNotNil("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061", _egfbd.SMaskInData)
	_ddbda.SetIfNotNil("\u004d\u0061\u0074t\u0065", _egfbd.Matte)
	_ddbda.SetIfNotNil("\u004e\u0061\u006d\u0065", _egfbd.Name)
	_ddbda.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _egfbd.StructParent)
	_ddbda.SetIfNotNil("\u0049\u0044", _egfbd.ID)
	_ddbda.SetIfNotNil("\u004f\u0050\u0049", _egfbd.OPI)
	_ddbda.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _egfbd.Metadata)
	_ddbda.SetIfNotNil("\u004f\u0043", _egfbd.OC)
	_ddbda.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cde.MakeInteger(int64(len(_egfbd.Stream))))
	_egcff.Stream = _egfbd.Stream
	return _egcff
}
func (_fff *PdfReader) newPdfActionNamedFromDict(_dgf *_cde.PdfObjectDictionary) (*PdfActionNamed, error) {
	return &PdfActionNamed{N: _dgf.Get("\u004e")}, nil
}

// PdfAnnotationWidget represents Widget annotations.
// Note: Widget annotations are used to display form fields.
// (Section 12.5.6.19).
type PdfAnnotationWidget struct {
	*PdfAnnotation
	H      _cde.PdfObject
	MK     _cde.PdfObject
	A      _cde.PdfObject
	AA     _cde.PdfObject
	BS     _cde.PdfObject
	Parent _cde.PdfObject
	_dbf   *PdfField
	_bga   bool
}

// PdfFieldSignature signature field represents digital signatures and optional data for authenticating
// the name of the signer and verifying document contents.
type PdfFieldSignature struct {
	*PdfField
	*PdfAnnotationWidget
	V    *PdfSignature
	Lock *_cde.PdfIndirectObject
	SV   *_cde.PdfIndirectObject
}

// String returns a string representation of what flags are set.
func (_fdfa FieldFlag) String() string {
	_cfgb := ""
	if _fdfa == FieldFlagClear {
		_cfgb = "\u0043\u006c\u0065a\u0072"
		return _cfgb
	}
	if _fdfa&FieldFlagReadOnly > 0 {
		_cfgb += "\u007cR\u0065\u0061\u0064\u004f\u006e\u006cy"
	}
	if _fdfa&FieldFlagRequired > 0 {
		_cfgb += "\u007cR\u0065\u0071\u0075\u0069\u0072\u0065d"
	}
	if _fdfa&FieldFlagNoExport > 0 {
		_cfgb += "\u007cN\u006f\u0045\u0078\u0070\u006f\u0072t"
	}
	if _fdfa&FieldFlagNoToggleToOff > 0 {
		_cfgb += "\u007c\u004e\u006f\u0054\u006f\u0067\u0067\u006c\u0065T\u006f\u004f\u0066\u0066"
	}
	if _fdfa&FieldFlagRadio > 0 {
		_cfgb += "\u007c\u0052\u0061\u0064\u0069\u006f"
	}
	if _fdfa&FieldFlagPushbutton > 0 {
		_cfgb += "|\u0050\u0075\u0073\u0068\u0062\u0075\u0074\u0074\u006f\u006e"
	}
	if _fdfa&FieldFlagRadiosInUnision > 0 {
		_cfgb += "\u007c\u0052a\u0064\u0069\u006fs\u0049\u006e\u0055\u006e\u0069\u0073\u0069\u006f\u006e"
	}
	if _fdfa&FieldFlagMultiline > 0 {
		_cfgb += "\u007c\u004d\u0075\u006c\u0074\u0069\u006c\u0069\u006e\u0065"
	}
	if _fdfa&FieldFlagPassword > 0 {
		_cfgb += "\u007cP\u0061\u0073\u0073\u0077\u006f\u0072d"
	}
	if _fdfa&FieldFlagFileSelect > 0 {
		_cfgb += "|\u0046\u0069\u006c\u0065\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _fdfa&FieldFlagDoNotScroll > 0 {
		_cfgb += "\u007c\u0044\u006fN\u006f\u0074\u0053\u0063\u0072\u006f\u006c\u006c"
	}
	if _fdfa&FieldFlagComb > 0 {
		_cfgb += "\u007c\u0043\u006fm\u0062"
	}
	if _fdfa&FieldFlagRichText > 0 {
		_cfgb += "\u007cR\u0069\u0063\u0068\u0054\u0065\u0078t"
	}
	if _fdfa&FieldFlagDoNotSpellCheck > 0 {
		_cfgb += "\u007c\u0044o\u004e\u006f\u0074S\u0070\u0065\u006c\u006c\u0043\u0068\u0065\u0063\u006b"
	}
	if _fdfa&FieldFlagCombo > 0 {
		_cfgb += "\u007c\u0043\u006f\u006d\u0062\u006f"
	}
	if _fdfa&FieldFlagEdit > 0 {
		_cfgb += "\u007c\u0045\u0064i\u0074"
	}
	if _fdfa&FieldFlagSort > 0 {
		_cfgb += "\u007c\u0053\u006fr\u0074"
	}
	if _fdfa&FieldFlagMultiSelect > 0 {
		_cfgb += "\u007c\u004d\u0075l\u0074\u0069\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _fdfa&FieldFlagCommitOnSelChange > 0 {
		_cfgb += "\u007cC\u006fm\u006d\u0069\u0074\u004f\u006eS\u0065\u006cC\u0068\u0061\u006e\u0067\u0065"
	}
	return _dac.Trim(_cfgb, "\u007c")
}

// ImageToRGB converts an Image in a given PdfColorspace to an RGB image.
func (_cfcee *PdfColorspaceDeviceN) ImageToRGB(img Image) (Image, error) {
	_bcbd := _cae.NewReader(img.getBase())
	_ccag := _ff.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, nil, img._deegf, img._aaafb)
	_gcbb := _cae.NewWriter(_ccag)
	_ggge := _ced.Pow(2, float64(img.BitsPerComponent)) - 1
	_aefg := _cfcee.GetNumComponents()
	_afaaf := make([]uint32, _aefg)
	_gagb := make([]float64, _aefg)
	for {
		_adbbg := _bcbd.ReadSamples(_afaaf)
		if _adbbg == _f.EOF {
			break
		} else if _adbbg != nil {
			return img, _adbbg
		}
		for _bbbe := 0; _bbbe < _aefg; _bbbe++ {
			_ebfe := float64(_afaaf[_bbbe]) / _ggge
			_gagb[_bbbe] = _ebfe
		}
		_gbgd, _adbbg := _cfcee.TintTransform.Evaluate(_gagb)
		if _adbbg != nil {
			return img, _adbbg
		}
		for _, _ffdc := range _gbgd {
			_ffdc = _ced.Min(_ced.Max(0, _ffdc), 1.0)
			if _adbbg = _gcbb.WriteSample(uint32(_ffdc * _ggge)); _adbbg != nil {
				return img, _adbbg
			}
		}
	}
	return _cfcee.AlternateSpace.ImageToRGB(_bddb(&_ccag))
}

// GetNumComponents returns the number of input color components, i.e. that are input to the tint transform.
func (_cdbf *PdfColorspaceDeviceN) GetNumComponents() int { return _cdbf.ColorantNames.Len() }

// NewPdfRectangle creates a PDF rectangle object based on an input array of 4 integers.
// Defining the lower left (LL) and upper right (UR) corners with
// floating point numbers.
func NewPdfRectangle(arr _cde.PdfObjectArray) (*PdfRectangle, error) {
	_ddcdb := PdfRectangle{}
	if arr.Len() != 4 {
		return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065\u0020\u0061\u0072r\u0061\u0079\u002c\u0020\u006c\u0065\u006e \u0021\u003d\u0020\u0034")
	}
	var _abefg error
	_ddcdb.Llx, _abefg = _cde.GetNumberAsFloat(arr.Get(0))
	if _abefg != nil {
		return nil, _abefg
	}
	_ddcdb.Lly, _abefg = _cde.GetNumberAsFloat(arr.Get(1))
	if _abefg != nil {
		return nil, _abefg
	}
	_ddcdb.Urx, _abefg = _cde.GetNumberAsFloat(arr.Get(2))
	if _abefg != nil {
		return nil, _abefg
	}
	_ddcdb.Ury, _abefg = _cde.GetNumberAsFloat(arr.Get(3))
	if _abefg != nil {
		return nil, _abefg
	}
	return &_ddcdb, nil
}
func (_eadab *pdfFontSimple) getFontDescriptor() *PdfFontDescriptor {
	if _acdgf := _eadab._fagf; _acdgf != nil {
		return _acdgf
	}
	return _eadab._ccff
}

// RunesToCharcodeBytes maps the provided runes to charcode bytes and it
// returns the resulting slice of bytes, along with the number of runes which
// could not be converted. If the number of misses is 0, all runes were
// successfully converted.
func (_efeee *PdfFont) RunesToCharcodeBytes(data []rune) ([]byte, int) {
	var _gccbf []_gc.TextEncoder
	var _cbfae _gc.CMapEncoder
	if _acaaa := _efeee.baseFields()._ggebg; _acaaa != nil {
		_cbfae = _gc.NewCMapEncoder("", nil, _acaaa)
	}
	_fccb := _efeee.Encoder()
	if _fccb != nil {
		switch _dceg := _fccb.(type) {
		case _gc.SimpleEncoder:
			_cbfb := _dceg.BaseName()
			if _, _dfgf := _egee[_cbfb]; _dfgf {
				_gccbf = append(_gccbf, _fccb)
			}
		}
	}
	if len(_gccbf) == 0 {
		if _efeee.baseFields()._ggebg != nil {
			_gccbf = append(_gccbf, _cbfae)
		}
		if _fccb != nil {
			_gccbf = append(_gccbf, _fccb)
		}
	}
	var _bggg _ede.Buffer
	var _egdde int
	for _, _fbfe := range data {
		var _ccad bool
		for _, _gdcg := range _gccbf {
			if _ffgdc := _gdcg.Encode(string(_fbfe)); len(_ffgdc) > 0 {
				_bggg.Write(_ffgdc)
				_ccad = true
				break
			}
		}
		if !_ccad {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020f\u0061\u0069\u006ce\u0064\u0020\u0074\u006f \u006d\u0061\u0070\u0020\u0072\u0075\u006e\u0065\u0020\u0060\u0025\u002b\u0071\u0060\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065", _fbfe)
			_egdde++
		}
	}
	if _egdde != 0 {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0061\u006cl\u0020\u0072\u0075\u006e\u0065\u0073\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072c\u006fd\u0065\u0073\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0052\u0075\u006e\u0065\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073 \u0065\u006e\u0063\u006f\u0064e\u0072\u0073\u003d\u0025\u002b\u0076", len(data), _egdde, _efeee, _gccbf)
	}
	return _bggg.Bytes(), _egdde
}

// GetOutlineTree returns the outline tree.
func (_ggdbb *PdfReader) GetOutlineTree() *PdfOutlineTreeNode { return _ggdbb._aegbb }

var (
	CourierName              = _fe.CourierName
	CourierBoldName          = _fe.CourierBoldName
	CourierObliqueName       = _fe.CourierObliqueName
	CourierBoldObliqueName   = _fe.CourierBoldObliqueName
	HelveticaName            = _fe.HelveticaName
	HelveticaBoldName        = _fe.HelveticaBoldName
	HelveticaObliqueName     = _fe.HelveticaObliqueName
	HelveticaBoldObliqueName = _fe.HelveticaBoldObliqueName
	SymbolName               = _fe.SymbolName
	ZapfDingbatsName         = _fe.ZapfDingbatsName
	TimesRomanName           = _fe.TimesRomanName
	TimesBoldName            = _fe.TimesBoldName
	TimesItalicName          = _fe.TimesItalicName
	TimesBoldItalicName      = _fe.TimesBoldItalicName
)

// PdfAcroForm represents the AcroForm dictionary used for representation of form data in PDF.
type PdfAcroForm struct {
	Fields          *[]*PdfField
	NeedAppearances *_cde.PdfObjectBool
	SigFlags        *_cde.PdfObjectInteger
	CO              *_cde.PdfObjectArray
	DR              *PdfPageResources
	DA              *_cde.PdfObjectString
	Q               *_cde.PdfObjectInteger
	XFA             _cde.PdfObject
	_ecca           *_cde.PdfIndirectObject
}

// ImageHandler interface implements common image loading and processing tasks.
// Implementing as an interface allows for the possibility to use non-standard libraries for faster
// loading and processing of images.
type ImageHandler interface {

	// Read any image type and load into a new Image object.
	Read(_dcgg _f.Reader) (*Image, error)

	// NewImageFromGoImage loads a NRGBA32 unidoc Image from a standard Go image structure.
	NewImageFromGoImage(_gfcge _gf.Image) (*Image, error)

	// NewGrayImageFromGoImage loads a grayscale unidoc Image from a standard Go image structure.
	NewGrayImageFromGoImage(_beeff _gf.Image) (*Image, error)

	// Compress an image.
	Compress(_dbef *Image, _fdfb int64) (*Image, error)
}

// ColorToRGB converts an Indexed color to an RGB color.
func (_ddfc *PdfColorspaceSpecialIndexed) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _ddfc.Base == nil {
		return nil, _ceg.New("\u0069\u006e\u0064\u0065\u0078\u0065d\u0020\u0062\u0061\u0073\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _ddfc.Base.ColorToRGB(color)
}

// PdfAnnotationStrikeOut represents StrikeOut annotations.
// (Section 12.5.6.10).
type PdfAnnotationStrikeOut struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _cde.PdfObject
}

func (_bfd *PdfReader) newPdfActionGotoFromDict(_eec *_cde.PdfObjectDictionary) (*PdfActionGoTo, error) {
	return &PdfActionGoTo{D: _eec.Get("\u0044")}, nil
}
func _fbbce(_fcgda *_cde.PdfObjectDictionary) (*PdfShadingPattern, error) {
	_fcgc := &PdfShadingPattern{}
	_edbgg := _fcgda.Get("\u0053h\u0061\u0064\u0069\u006e\u0067")
	if _edbgg == nil {
		_ad.Log.Debug("\u0053h\u0061d\u0069\u006e\u0067\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_deee, _fcfce := _bfbdc(_edbgg)
	if _fcfce != nil {
		_ad.Log.Debug("\u0045r\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006eg\u0020s\u0068a\u0064\u0069\u006e\u0067\u003a\u0020\u0025v", _fcfce)
		return nil, _fcfce
	}
	_fcgc.Shading = _deee
	if _aadbc := _fcgda.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _aadbc != nil {
		_fdbad, _babcf := _aadbc.(*_cde.PdfObjectArray)
		if !_babcf {
			_ad.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _aadbc)
			return nil, _cde.ErrTypeError
		}
		_fcgc.Matrix = _fdbad
	}
	if _gagf := _fcgda.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _gagf != nil {
		_fcgc.ExtGState = _gagf
	}
	return _fcgc, nil
}

type pdfCIDFontType2 struct {
	fontCommon
	_abca *_cde.PdfIndirectObject
	_egga _gc.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// Dictionary that defines the character collection of the CIDFont (required).
	// See Table 116.
	CIDSystemInfo *_cde.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW  _cde.PdfObject
	W   _cde.PdfObject
	DW2 _cde.PdfObject
	W2  _cde.PdfObject

	// CIDs to glyph indices mapping (optional).
	CIDToGIDMap _cde.PdfObject
	_fdgc       map[_gc.CharCode]float64
	_eedac      float64
	_dfefe      map[rune]int
}

func _aebaa(_cacf *_cde.PdfObjectDictionary) (*PdfFieldButton, error) {
	_acec := &PdfFieldButton{}
	_acec.PdfField = NewPdfField()
	_acec.PdfField.SetContext(_acec)
	_acec.Opt, _ = _cde.GetArray(_cacf.Get("\u004f\u0070\u0074"))
	_efbdg := NewPdfAnnotationWidget()
	_efbdg.A, _ = _cde.GetDict(_cacf.Get("\u0041"))
	_efbdg.AP, _ = _cde.GetDict(_cacf.Get("\u0041\u0050"))
	_efbdg.SetContext(_acec)
	_acec.PdfField.Annotations = append(_acec.PdfField.Annotations, _efbdg)
	return _acec, nil
}

// PdfOutline represents a PDF outline dictionary (Table 152 - p. 376).
type PdfOutline struct {
	PdfOutlineTreeNode
	Parent *PdfOutlineTreeNode
	Count  *int64
	_dgcdc *_cde.PdfIndirectObject
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_dfaa *PdfColorspaceSpecialSeparation) ToPdfObject() _cde.PdfObject {
	_cfagc := _cde.MakeArray(_cde.MakeName("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"))
	_cfagc.Append(_dfaa.ColorantName)
	_cfagc.Append(_dfaa.AlternateSpace.ToPdfObject())
	_cfagc.Append(_dfaa.TintTransform.ToPdfObject())
	if _dfaa._bcdc != nil {
		_dfaa._bcdc.PdfObject = _cfagc
		return _dfaa._bcdc
	}
	return _cfagc
}

// ToInteger convert to an integer format.
func (_cced *PdfColorDeviceCMYK) ToInteger(bits int) [4]uint32 {
	_eafe := _ced.Pow(2, float64(bits)) - 1
	return [4]uint32{uint32(_eafe * _cced.C()), uint32(_eafe * _cced.M()), uint32(_eafe * _cced.Y()), uint32(_eafe * _cced.K())}
}

// DecodeArray returns the range of color component values in DeviceGray colorspace.
func (_bcgg *PdfColorspaceDeviceGray) DecodeArray() []float64 { return []float64{0, 1.0} }

// NewCustomPdfOutputIntent creates a new custom PdfOutputIntent.
func NewCustomPdfOutputIntent(outputCondition, outputConditionIdentifier, info string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, Info: info, DestOutputProfile: destOutputProfile, _acbcb: _cde.MakeDict(), ColorComponents: colorComponents}
}

const (
	RC4_128bit = EncryptionAlgorithm(iota)
	AES_128bit
	AES_256bit
)

// AnnotFilterFunc represents a PDF annotation filtering function. If the function
// returns true, the annotation is kept, otherwise it is discarded.
type AnnotFilterFunc func(*PdfAnnotation) bool

var _fdga = map[string]struct{}{"\u0054\u0069\u0074l\u0065": {}, "\u0041\u0075\u0074\u0068\u006f\u0072": {}, "\u0053u\u0062\u006a\u0065\u0063\u0074": {}, "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073": {}, "\u0043r\u0065\u0061\u0074\u006f\u0072": {}, "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072": {}, "\u0054r\u0061\u0070\u0070\u0065\u0064": {}, "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065": {}, "\u004do\u0064\u0044\u0061\u0074\u0065": {}}

// GetCustomInfo returns a custom info value for the specified name.
func (_bbbeg *PdfInfo) GetCustomInfo(name string) *_cde.PdfObjectString {
	var _cccc *_cde.PdfObjectString
	if _bbbeg._ccef == nil {
		return _cccc
	}
	if _bggab, _bgeea := _bbbeg._ccef.Get(*_cde.MakeName(name)).(*_cde.PdfObjectString); _bgeea {
		_cccc = _bggab
	}
	return _cccc
}

// GetPreviousRevision returns the previous revision of PdfReader for the Pdf document
func (_egfcg *PdfReader) GetPreviousRevision() (*PdfReader, error) {
	if _egfcg._aggcgb.GetRevisionNumber() == 0 {
		return nil, _ceg.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065xi\u0073\u0074")
	}
	if _ggab, _abdf := _egfcg._egaba[_egfcg]; _abdf {
		return _ggab, nil
	}
	_bbeebc, _baag := _egfcg._aggcgb.GetPreviousRevisionReadSeeker()
	if _baag != nil {
		return nil, _baag
	}
	_bbec, _baag := _gfceb(_bbeebc, _egfcg._cbbc, _egfcg._cgebe, "\u006do\u0064\u0065\u006c\u003aG\u0065\u0074\u0050\u0072\u0065v\u0069o\u0075s\u0052\u0065\u0076\u0069\u0073\u0069\u006fn")
	if _baag != nil {
		return nil, _baag
	}
	_egfcg._fgfbe[_egfcg._aggcgb.GetRevisionNumber()-1] = _bbec
	_egfcg._egaba[_egfcg] = _bbec
	_bbec._egaba = _egfcg._egaba
	return _bbec, nil
}

// SetPdfModifiedDate sets the ModDate attribute of the output PDF.
func SetPdfModifiedDate(modifiedDate _ce.Time) {
	_dccfe.Lock()
	defer _dccfe.Unlock()
	_cebaa = modifiedDate
}

// ToPdfObject implements interface PdfModel.
func (_abgeg *Permissions) ToPdfObject() _cde.PdfObject { return _abgeg._dafaf }

// SetXObjectByName adds the XObject from the passed in stream to the page resources.
// The added XObject is identified by the specified name.
func (_bdcce *PdfPageResources) SetXObjectByName(keyName _cde.PdfObjectName, stream *_cde.PdfObjectStream) error {
	if _bdcce.XObject == nil {
		_bdcce.XObject = _cde.MakeDict()
	}
	_aaggd := _cde.TraceToDirectObject(_bdcce.XObject)
	_ebcbe, _bffag := _aaggd.(*_cde.PdfObjectDictionary)
	if !_bffag {
		_ad.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0058\u004f\u0062j\u0065\u0063\u0074\u002c\u0020\u0067\u006f\u0074\u0020\u0025T\u002f\u0025\u0054", _bdcce.XObject, _aaggd)
		return _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_ebcbe.Set(keyName, stream)
	return nil
}

// StandardApplier is the interface that performs optimization of the whole PDF document.
// As a result an input document is being changed by the optimizer.
// The writer than takes back all it's parts and overwrites it.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
type StandardApplier interface {
	ApplyStandard(_abgb *_ab.Document) error
}

func (_dff *PdfReader) newPdfActionThreadFromDict(_cge *_cde.PdfObjectDictionary) (*PdfActionThread, error) {
	_abf, _dba := _beed(_cge.Get("\u0046"))
	if _dba != nil {
		return nil, _dba
	}
	return &PdfActionThread{D: _cge.Get("\u0044"), B: _cge.Get("\u0042"), F: _abf}, nil
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_geffeb *PdfFunctionType2) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_ad.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bcgb := []float64{0.0}
	if _geffeb.C0 != nil {
		_bcgb = _geffeb.C0
	}
	_cebf := []float64{1.0}
	if _geffeb.C1 != nil {
		_cebf = _geffeb.C1
	}
	var _fdbb []float64
	for _ddbfg := 0; _ddbfg < len(_bcgb); _ddbfg++ {
		_dbeac := _bcgb[_ddbfg] + _ced.Pow(x[0], _geffeb.N)*(_cebf[_ddbfg]-_bcgb[_ddbfg])
		_fdbb = append(_fdbb, _dbeac)
	}
	return _fdbb, nil
}
func (_febf *PdfWriter) setWriter(_bbdea _f.Writer) {
	_febf._eddbc = _febf._fgca
	_febf._cefdd = _b.NewWriter(_bbdea)
}

// NewPdfOutlineTree returns an initialized PdfOutline tree.
func NewPdfOutlineTree() *PdfOutline {
	_afbbb := NewPdfOutline()
	_afbbb._fbeea = &_afbbb
	return _afbbb
}

// PdfAnnotationCaret represents Caret annotations.
// (Section 12.5.6.11).
type PdfAnnotationCaret struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	RD _cde.PdfObject
	Sy _cde.PdfObject
}

// ToPdfObject returns colorspace in a PDF object format [name stream]
func (_gcfd *PdfColorspaceICCBased) ToPdfObject() _cde.PdfObject {
	_fdde := &_cde.PdfObjectArray{}
	_fdde.Append(_cde.MakeName("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"))
	var _ddfe *_cde.PdfObjectStream
	if _gcfd._fdea != nil {
		_ddfe = _gcfd._fdea
	} else {
		_ddfe = &_cde.PdfObjectStream{}
	}
	_cbac := _cde.MakeDict()
	_cbac.Set("\u004e", _cde.MakeInteger(int64(_gcfd.N)))
	if _gcfd.Alternate != nil {
		_cbac.Set("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e", _gcfd.Alternate.ToPdfObject())
	}
	if _gcfd.Metadata != nil {
		_cbac.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _gcfd.Metadata)
	}
	if _gcfd.Range != nil {
		var _daag []_cde.PdfObject
		for _, _efae := range _gcfd.Range {
			_daag = append(_daag, _cde.MakeFloat(_efae))
		}
		_cbac.Set("\u0052\u0061\u006eg\u0065", _cde.MakeArray(_daag...))
	}
	_cbac.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cde.MakeInteger(int64(len(_gcfd.Data))))
	_ddfe.Stream = _gcfd.Data
	_ddfe.PdfObjectDictionary = _cbac
	_fdde.Append(_ddfe)
	if _gcfd._cffb != nil {
		_gcfd._cffb.PdfObject = _fdde
		return _gcfd._cffb
	}
	return _fdde
}

// C returns the value of the C component of the color.
func (_cagg *PdfColorCalRGB) C() float64 { return _cagg[2] }

// PdfActionGoTo represents a GoTo action.
type PdfActionGoTo struct {
	*PdfAction
	D _cde.PdfObject
}

func _agca(_fgfcb _cde.PdfObject) (*PdfColorspaceSpecialPattern, error) {
	_ad.Log.Trace("\u004e\u0065\u0077\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u0043\u0053\u0020\u0066\u0072\u006fm\u0020\u006f\u0062\u006a\u003a\u0020\u0025\u0073\u0020\u0025\u0054", _fgfcb.String(), _fgfcb)
	_fcacd := NewPdfColorspaceSpecialPattern()
	if _bgfd, _dfgb := _fgfcb.(*_cde.PdfIndirectObject); _dfgb {
		_fcacd._ffcb = _bgfd
	}
	_fgfcb = _cde.TraceToDirectObject(_fgfcb)
	if _dece, _dgdffe := _fgfcb.(*_cde.PdfObjectName); _dgdffe {
		if *_dece != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			return nil, _ee.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
		return _fcacd, nil
	}
	_cfba, _acbf := _fgfcb.(*_cde.PdfObjectArray)
	if !_acbf {
		_ad.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061t\u0074\u0065\u0072\u006e\u0020\u0043\u0053 \u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0023\u0076", _fgfcb)
		return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0050\u0061\u0074\u0074e\u0072n\u0020C\u0053\u0020\u006f\u0062\u006a\u0065\u0063t")
	}
	if _cfba.Len() != 1 && _cfba.Len() != 2 {
		_ad.Log.Error("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079\u003a %\u0023\u0076", _cfba)
		return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065r\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_fgfcb = _cfba.Get(0)
	if _aaeg, _ebcgb := _fgfcb.(*_cde.PdfObjectName); _ebcgb {
		if *_aaeg != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			_ad.Log.Error("\u0049\u006e\u0076al\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020C\u0053 \u0061r\u0072\u0061\u0079\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0023\u0076", _aaeg)
			return nil, _ee.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
	}
	if _cfba.Len() > 1 {
		_fgfcb = _cfba.Get(1)
		_fgfcb = _cde.TraceToDirectObject(_fgfcb)
		_bgff, _dfdg := NewPdfColorspaceFromPdfObject(_fgfcb)
		if _dfdg != nil {
			return nil, _dfdg
		}
		_fcacd.UnderlyingCS = _bgff
	}
	_ad.Log.Trace("R\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0077i\u0074\u0068\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079in\u0067\u0020\u0063s\u003a \u0025\u0054", _fcacd.UnderlyingCS)
	return _fcacd, nil
}

// NewPdfTransformParamsDocMDP create a PdfTransformParamsDocMDP with the specific permissions.
func NewPdfTransformParamsDocMDP(permission _cdc.DocMDPPermission) *PdfTransformParamsDocMDP {
	return &PdfTransformParamsDocMDP{Type: _cde.MakeName("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"), P: _cde.MakeInteger(int64(permission)), V: _cde.MakeName("\u0031\u002e\u0032")}
}
func _aeeb(_fdge *_cde.PdfObjectDictionary) bool {
	for _, _bbff := range _fdge.Keys() {
		if _, _cceae := _gfac[_bbff.String()]; _cceae {
			return true
		}
	}
	return false
}

// GetContainingPdfObject returns the container of the outline (indirect object).
func (_gdefce *PdfOutline) GetContainingPdfObject() _cde.PdfObject { return _gdefce._dgcdc }

// NewPermissions returns a new permissions object.
func NewPermissions(docMdp *PdfSignature) *Permissions {
	_aaecb := Permissions{}
	_aaecb.DocMDP = docMdp
	_ddfbee := _cde.MakeDict()
	_ddfbee.Set("\u0044\u006f\u0063\u004d\u0044\u0050", docMdp.ToPdfObject())
	_aaecb._dafaf = _ddfbee
	return &_aaecb
}
func (_cbefd *PdfReader) buildNameNodes(_begab *_cde.PdfIndirectObject, _aagfe map[_cde.PdfObject]struct{}) error {
	if _begab == nil {
		return nil
	}
	if _, _gccbc := _aagfe[_begab]; _gccbc {
		_ad.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _begab.ObjectNumber)
		return nil
	}
	_aagfe[_begab] = struct{}{}
	_aafcea, _bafebg := _begab.PdfObject.(*_cde.PdfObjectDictionary)
	if !_bafebg {
		return _ceg.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _ecbaa, _cbbff := _cde.GetDict(_aafcea.Get("\u0044\u0065\u0073t\u0073")); _cbbff {
		_dcgb, _gbdag := _cde.GetArray(_ecbaa.Get("\u004b\u0069\u0064\u0073"))
		if !_gbdag {
			return _ceg.New("\u0049n\u0076\u0061\u006c\u0069d\u0020\u004b\u0069\u0064\u0073 \u0061r\u0072a\u0079\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_ad.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _dcgb)
		for _geeca, _cbgdc := range _dcgb.Elements() {
			_cgcgc, _egfcd := _cde.GetIndirect(_cbgdc)
			if !_egfcd {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u0068\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002d \u0028\u0025\u0073\u0029", _cgcgc)
				return _ceg.New("\u0063h\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006ed\u0069r\u0065c\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
			}
			_dcgb.Set(_geeca, _cgcgc)
			_ffgge := _cbefd.buildNameNodes(_cgcgc, _aagfe)
			if _ffgge != nil {
				return _ffgge
			}
		}
	}
	if _bgdged, _cdagc := _cde.GetDict(_aafcea); _cdagc {
		if !_cde.IsNullObject(_bgdged.Get("\u004b\u0069\u0064\u0073")) {
			if _bdeff, _fegda := _cde.GetArray(_bgdged.Get("\u004b\u0069\u0064\u0073")); _fegda {
				for _fdeaad, _aaegf := range _bdeff.Elements() {
					if _cdca, _gdccb := _cde.GetIndirect(_aaegf); _gdccb {
						_bdeff.Set(_fdeaad, _cdca)
						_bgaaf := _cbefd.buildNameNodes(_cdca, _aagfe)
						if _bgaaf != nil {
							return _bgaaf
						}
					}
				}
			}
		}
	}
	return nil
}

// ColorToRGB converts a ICCBased color to an RGB color.
func (_caaf *PdfColorspaceICCBased) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _caaf.Alternate == nil {
		_ad.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _caaf.N == 1 {
			_ad.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_fedb := NewPdfColorspaceDeviceGray()
			return _fedb.ColorToRGB(color)
		} else if _caaf.N == 3 {
			_ad.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return color, nil
		} else if _caaf.N == 4 {
			_ad.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_bfbec := NewPdfColorspaceDeviceCMYK()
			return _bfbec.ColorToRGB(color)
		} else {
			return nil, _ceg.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_ad.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _caaf)
	return _caaf.Alternate.ColorToRGB(color)
}

// Val returns the value of the color.
func (_fade *PdfColorCalGray) Val() float64 { return float64(*_fade) }

// SetContentStreams sets the content streams based on a string array. Will make
// 1 object stream for each string and reference from the page Contents.
// Each stream will be encoded using the encoding specified by the StreamEncoder,
// if empty, will use identity encoding (raw data).
func (_edbbb *PdfPage) SetContentStreams(cStreams []string, encoder _cde.StreamEncoder) error {
	if len(cStreams) == 0 {
		_edbbb.Contents = nil
		return nil
	}
	if encoder == nil {
		encoder = _cde.NewRawEncoder()
	}
	var _gdeed []*_cde.PdfObjectStream
	for _, _cabaa := range cStreams {
		_cabfd := &_cde.PdfObjectStream{}
		_aefef := encoder.MakeStreamDict()
		_ccbe, _cfged := encoder.EncodeBytes([]byte(_cabaa))
		if _cfged != nil {
			return _cfged
		}
		_aefef.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cde.MakeInteger(int64(len(_ccbe))))
		_cabfd.PdfObjectDictionary = _aefef
		_cabfd.Stream = _ccbe
		_gdeed = append(_gdeed, _cabfd)
	}
	if len(_gdeed) == 1 {
		_edbbb.Contents = _gdeed[0]
	} else {
		_fcfgc := _cde.MakeArray()
		for _, _gefe := range _gdeed {
			_fcfgc.Append(_gefe)
		}
		_edbbb.Contents = _fcfgc
	}
	return nil
}

// NewPdfActionHide returns a new "hide" action.
func NewPdfActionHide() *PdfActionHide {
	_gea := NewPdfAction()
	_gfa := &PdfActionHide{}
	_gfa.PdfAction = _gea
	_gea.SetContext(_gfa)
	return _gfa
}

// ColorToRGB converts gray -> rgb for a single color component.
func (_gbbf *PdfColorspaceDeviceGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_aggc, _ffbe := color.(*PdfColorDeviceGray)
	if !_ffbe {
		_ad.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006fr\u0020\u006e\u006f\u0074\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0067\u0072\u0061\u0079\u0020\u0025\u0054", color)
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return NewPdfColorDeviceRGB(float64(*_aggc), float64(*_aggc), float64(*_aggc)), nil
}

// PageProcessCallback callback function used in page loading
// that could be used to modify the page content.
//
// If an error is returned, the `ToWriter` process would fail.
//
// This callback, if defined, will take precedence over `PageCallback` callback.
type PageProcessCallback func(_fece int, _bagfb *PdfPage) error

func _cgcge(_fcba *_cde.PdfObjectDictionary) (*PdfShadingType4, error) {
	_dbdc := PdfShadingType4{}
	_agfcg := _fcba.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _agfcg == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_facfa, _ccbba := _agfcg.(*_cde.PdfObjectInteger)
	if !_ccbba {
		_ad.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _agfcg)
		return nil, _cde.ErrTypeError
	}
	_dbdc.BitsPerCoordinate = _facfa
	_agfcg = _fcba.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _agfcg == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_facfa, _ccbba = _agfcg.(*_cde.PdfObjectInteger)
	if !_ccbba {
		_ad.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _agfcg)
		return nil, _cde.ErrTypeError
	}
	_dbdc.BitsPerComponent = _facfa
	_agfcg = _fcba.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _agfcg == nil {
		_ad.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_facfa, _ccbba = _agfcg.(*_cde.PdfObjectInteger)
	if !_ccbba {
		_ad.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _agfcg)
		return nil, _cde.ErrTypeError
	}
	_dbdc.BitsPerComponent = _facfa
	_agfcg = _fcba.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _agfcg == nil {
		_ad.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_aaebb, _ccbba := _agfcg.(*_cde.PdfObjectArray)
	if !_ccbba {
		_ad.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _agfcg)
		return nil, _cde.ErrTypeError
	}
	_dbdc.Decode = _aaebb
	_agfcg = _fcba.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _agfcg == nil {
		_ad.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_dbdc.Function = []PdfFunction{}
	if _beaba, _afbf := _agfcg.(*_cde.PdfObjectArray); _afbf {
		for _, _egec := range _beaba.Elements() {
			_bfbee, _eefd := _cfdbb(_egec)
			if _eefd != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _eefd)
				return nil, _eefd
			}
			_dbdc.Function = append(_dbdc.Function, _bfbee)
		}
	} else {
		_ddaac, _debcac := _cfdbb(_agfcg)
		if _debcac != nil {
			_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _debcac)
			return nil, _debcac
		}
		_dbdc.Function = append(_dbdc.Function, _ddaac)
	}
	return &_dbdc, nil
}

// PdfSignatureReference represents a PDF signature reference dictionary and is used for signing via form signature fields.
// (Section 12.8.1, Table 253 - Entries in a signature reference dictionary p. 469 in PDF32000_2008).
type PdfSignatureReference struct {
	_fggee          *_cde.PdfObjectDictionary
	Type            *_cde.PdfObjectName
	TransformMethod *_cde.PdfObjectName
	TransformParams _cde.PdfObject
	Data            _cde.PdfObject
	DigestMethod    *_cde.PdfObjectName
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_gadfa pdfCIDFontType0) GetCharMetrics(code _gc.CharCode) (_fe.CharMetrics, bool) {
	_bdffa := _gadfa._bdega
	if _facaf, _dfded := _gadfa._egfeb[code]; _dfded {
		_bdffa = _facaf
	}
	return _fe.CharMetrics{Wx: _bdffa}, true
}

// SetSamples convert samples to byte-data and sets for the image.
// NOTE: The method resamples the data and this could lead to high memory usage,
// especially on large images. It should be used only when it is not possible
// to work with the image byte data directly.
func (_caebe *Image) SetSamples(samples []uint32) {
	if _caebe.BitsPerComponent < 8 {
		samples = _caebe.samplesAddPadding(samples)
	}
	_cfaac := _cae.ResampleUint32(samples, int(_caebe.BitsPerComponent), 8)
	_bedef := make([]byte, len(_cfaac))
	for _cafee, _acdaa := range _cfaac {
		_bedef[_cafee] = byte(_acdaa)
	}
	_caebe.Data = _bedef
}

// GetContext returns the PdfField context which is the more specific field data type, e.g. PdfFieldButton
// for a button field.
func (_dbad *PdfField) GetContext() PdfModel { return _dbad._ecfg }
func _ebag(_daaf string) map[string]string {
	_gbdcf := _caggc.Split(_daaf, -1)
	_gagda := map[string]string{}
	for _, _aeafc := range _gbdcf {
		_debe := _fafdd.FindStringSubmatch(_aeafc)
		if _debe == nil {
			continue
		}
		_aadac, _efbb := _debe[1], _debe[2]
		_gagda[_aadac] = _efbb
	}
	return _gagda
}

// SetContext sets the sub annotation (context).
func (_fgg *PdfAnnotation) SetContext(ctx PdfModel) { _fgg._bea = ctx }

// DecodeArray returns an empty slice as there are no components associated with pattern colorspace.
func (_dfgag *PdfColorspaceSpecialPattern) DecodeArray() []float64 { return []float64{} }

// ToPdfObject returns the PDF representation of the VRI dictionary.
func (_edebc *VRI) ToPdfObject() *_cde.PdfObjectDictionary {
	_dcce := _cde.MakeDict()
	_dcce.SetIfNotNil(_cde.PdfObjectName("\u0043\u0065\u0072\u0074"), _dbgae(_edebc.Cert))
	_dcce.SetIfNotNil(_cde.PdfObjectName("\u004f\u0043\u0053\u0050"), _dbgae(_edebc.OCSP))
	_dcce.SetIfNotNil(_cde.PdfObjectName("\u0043\u0052\u004c"), _dbgae(_edebc.CRL))
	_dcce.SetIfNotNil("\u0054\u0055", _edebc.TU)
	_dcce.SetIfNotNil("\u0054\u0053", _edebc.TS)
	return _dcce
}

// ColorFromPdfObjects gets the color from a series of pdf objects (3 for rgb).
func (_eeggg *PdfColorspaceDeviceRGB) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cfce, _fbbcc := _cde.GetNumbersAsFloat(objects)
	if _fbbcc != nil {
		return nil, _fbbcc
	}
	return _eeggg.ColorFromFloats(_cfce)
}

// ToPdfObject implements interface PdfModel.
func (_ade *PdfActionThread) ToPdfObject() _cde.PdfObject {
	_ade.PdfAction.ToPdfObject()
	_dca := _ade._bc
	_eg := _dca.PdfObject.(*_cde.PdfObjectDictionary)
	_eg.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeThread)))
	if _ade.F != nil {
		_eg.Set("\u0046", _ade.F.ToPdfObject())
	}
	_eg.SetIfNotNil("\u0044", _ade.D)
	_eg.SetIfNotNil("\u0042", _ade.B)
	return _dca
}
func _dede(_acabe, _ggfce string) string {
	if _dac.Contains(_acabe, "\u002b") {
		_abcca := _dac.Split(_acabe, "\u002b")
		if len(_abcca) == 2 {
			_acabe = _abcca[1]
		}
	}
	return _ggfce + "\u002b" + _acabe
}

// NewOutlineItem returns a new outline item instance.
func NewOutlineItem(title string, dest OutlineDest) *OutlineItem {
	return &OutlineItem{Title: title, Dest: dest}
}

// ToPdfObject return the CalGray colorspace as a PDF object (name dictionary).
func (_abgcc *PdfColorspaceCalGray) ToPdfObject() _cde.PdfObject {
	_febc := &_cde.PdfObjectArray{}
	_febc.Append(_cde.MakeName("\u0043a\u006c\u0047\u0072\u0061\u0079"))
	_eacae := _cde.MakeDict()
	if _abgcc.WhitePoint != nil {
		_eacae.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _cde.MakeArray(_cde.MakeFloat(_abgcc.WhitePoint[0]), _cde.MakeFloat(_abgcc.WhitePoint[1]), _cde.MakeFloat(_abgcc.WhitePoint[2])))
	} else {
		_ad.Log.Error("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006fi\u006e\u0074\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _abgcc.BlackPoint != nil {
		_eacae.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _cde.MakeArray(_cde.MakeFloat(_abgcc.BlackPoint[0]), _cde.MakeFloat(_abgcc.BlackPoint[1]), _cde.MakeFloat(_abgcc.BlackPoint[2])))
	}
	_eacae.Set("\u0047\u0061\u006dm\u0061", _cde.MakeFloat(_abgcc.Gamma))
	_febc.Append(_eacae)
	if _abgcc._ggffb != nil {
		_abgcc._ggffb.PdfObject = _febc
		return _abgcc._ggffb
	}
	return _febc
}

// GetContext returns the annotation context which contains the specific type-dependent context.
// The context represents the subannotation.
func (_bff *PdfAnnotation) GetContext() PdfModel {
	if _bff == nil {
		return nil
	}
	return _bff._bea
}

// NewPdfActionThread returns a new "thread" action.
func NewPdfActionThread() *PdfActionThread {
	_df := NewPdfAction()
	_fgd := &PdfActionThread{}
	_fgd.PdfAction = _df
	_df.SetContext(_fgd)
	return _fgd
}

// NewPdfActionSetOCGState returns a new "named" action.
func NewPdfActionSetOCGState() *PdfActionSetOCGState {
	_cag := NewPdfAction()
	_edg := &PdfActionSetOCGState{}
	_edg.PdfAction = _cag
	_cag.SetContext(_edg)
	return _edg
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_ffgc *PdfColorspaceDeviceRGB) ToPdfObject() _cde.PdfObject {
	return _cde.MakeName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
}

// ToInteger convert to an integer format.
func (_fbgb *PdfColorDeviceGray) ToInteger(bits int) uint32 {
	_dbfaa := _ced.Pow(2, float64(bits)) - 1
	return uint32(_dbfaa * _fbgb.Val())
}

// IsCID returns true if the underlying font is CID.
func (_acfc *PdfFont) IsCID() bool { return _acfc.baseFields().isCIDFont() }

// NewPdfActionImportData returns a new "import data" action.
func NewPdfActionImportData() *PdfActionImportData {
	_de := NewPdfAction()
	_ccf := &PdfActionImportData{}
	_ccf.PdfAction = _de
	_de.SetContext(_ccf)
	return _ccf
}

// HasXObjectByName checks if has XObject resource by name.
func (_egbcc *PdfPage) HasXObjectByName(name _cde.PdfObjectName) bool {
	_baae, _eadad := _egbcc.Resources.XObject.(*_cde.PdfObjectDictionary)
	if !_eadad {
		return false
	}
	if _cfedf := _baae.Get(name); _cfedf != nil {
		return true
	}
	return false
}

// PdfSignature represents a PDF signature dictionary and is used for signing via form signature fields.
// (Section 12.8, Table 252 - Entries in a signature dictionary p. 475 in PDF32000_2008).
type PdfSignature struct {
	Handler SignatureHandler
	_cabd   *_cde.PdfIndirectObject

	// Type: Sig/DocTimeStamp
	Type         *_cde.PdfObjectName
	Filter       *_cde.PdfObjectName
	SubFilter    *_cde.PdfObjectName
	Contents     *_cde.PdfObjectString
	Cert         _cde.PdfObject
	ByteRange    *_cde.PdfObjectArray
	Reference    *_cde.PdfObjectArray
	Changes      *_cde.PdfObjectArray
	Name         *_cde.PdfObjectString
	M            *_cde.PdfObjectString
	Location     *_cde.PdfObjectString
	Reason       *_cde.PdfObjectString
	ContactInfo  *_cde.PdfObjectString
	R            *_cde.PdfObjectInteger
	V            *_cde.PdfObjectInteger
	PropBuild    *_cde.PdfObjectDictionary
	PropAuthTime *_cde.PdfObjectInteger
	PropAuthType *_cde.PdfObjectName
}

func (_deac *PdfReader) newPdfAnnotationTrapNetFromDict(_gfaf *_cde.PdfObjectDictionary) (*PdfAnnotationTrapNet, error) {
	_fcac := PdfAnnotationTrapNet{}
	return &_fcac, nil
}

// NewPdfAnnotationHighlight returns a new text highlight annotation.
func NewPdfAnnotationHighlight() *PdfAnnotationHighlight {
	_dbbf := NewPdfAnnotation()
	_ffe := &PdfAnnotationHighlight{}
	_ffe.PdfAnnotation = _dbbf
	_ffe.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dbbf.SetContext(_ffe)
	return _ffe
}

// ToPdfObject implements interface PdfModel.
func (_agc *PdfActionResetForm) ToPdfObject() _cde.PdfObject {
	_agc.PdfAction.ToPdfObject()
	_bag := _agc._bc
	_aeff := _bag.PdfObject.(*_cde.PdfObjectDictionary)
	_aeff.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeResetForm)))
	_aeff.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _agc.Fields)
	_aeff.SetIfNotNil("\u0046\u006c\u0061g\u0073", _agc.Flags)
	return _bag
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_cabab pdfFontType3) GetRuneMetrics(r rune) (_fe.CharMetrics, bool) {
	_bbbg := _cabab.Encoder()
	if _bbbg == nil {
		_ad.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _cabab)
		return _fe.CharMetrics{}, false
	}
	_gagbe, _gcafg := _bbbg.RuneToCharcode(r)
	if !_gcafg {
		if r != ' ' {
			_ad.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _cabab)
		}
		return _fe.CharMetrics{}, false
	}
	_cgggg, _aabg := _cabab.GetCharMetrics(_gagbe)
	return _cgggg, _aabg
}

// DefaultFont returns the default font, which is currently the built in Helvetica.
func DefaultFont() *PdfFont {
	_ecgdd, _cbage := _fe.NewStdFontByName(HelveticaName)
	if !_cbage {
		panic("\u0048\u0065lv\u0065\u0074\u0069c\u0061\u0020\u0073\u0068oul\u0064 a\u006c\u0077\u0061\u0079\u0073\u0020\u0062e \u0061\u0076\u0061\u0069\u006c\u0061\u0062l\u0065")
	}
	_gecf := _gdaeg(_ecgdd)
	return &PdfFont{_gbcff: &_gecf}
}

// ToPdfObject implements interface PdfModel.
func (_ag *PdfActionGoToE) ToPdfObject() _cde.PdfObject {
	_ag.PdfAction.ToPdfObject()
	_cfbd := _ag._bc
	_eb := _cfbd.PdfObject.(*_cde.PdfObjectDictionary)
	_eb.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeGoToE)))
	if _ag.F != nil {
		_eb.Set("\u0046", _ag.F.ToPdfObject())
	}
	_eb.SetIfNotNil("\u0044", _ag.D)
	_eb.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _ag.NewWindow)
	_eb.SetIfNotNil("\u0054", _ag.T)
	return _cfbd
}

// GetContext returns the action context which contains the specific type-dependent context.
// The context represents the subaction.
func (_fg *PdfAction) GetContext() PdfModel {
	if _fg == nil {
		return nil
	}
	return _fg._bgd
}
func (_abaec *PdfWriter) checkCrossReferenceStream() bool {
	_dgddd := _abaec._cgdcc.Major > 1 || (_abaec._cgdcc.Major == 1 && _abaec._cgdcc.Minor > 4)
	if _abaec._cagac != nil {
		_dgddd = *_abaec._cagac
	}
	return _dgddd
}

// ToPdfObject implements interface PdfModel.
func (_ffa *PdfAnnotationFreeText) ToPdfObject() _cde.PdfObject {
	_ffa.PdfAnnotation.ToPdfObject()
	_aceb := _ffa._bddg
	_dcfc := _aceb.PdfObject.(*_cde.PdfObjectDictionary)
	_ffa.PdfAnnotationMarkup.appendToPdfDictionary(_dcfc)
	_dcfc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074"))
	_dcfc.SetIfNotNil("\u0044\u0041", _ffa.DA)
	_dcfc.SetIfNotNil("\u0051", _ffa.Q)
	_dcfc.SetIfNotNil("\u0052\u0043", _ffa.RC)
	_dcfc.SetIfNotNil("\u0044\u0053", _ffa.DS)
	_dcfc.SetIfNotNil("\u0043\u004c", _ffa.CL)
	_dcfc.SetIfNotNil("\u0049\u0054", _ffa.IT)
	_dcfc.SetIfNotNil("\u0042\u0045", _ffa.BE)
	_dcfc.SetIfNotNil("\u0052\u0044", _ffa.RD)
	_dcfc.SetIfNotNil("\u0042\u0053", _ffa.BS)
	_dcfc.SetIfNotNil("\u004c\u0045", _ffa.LE)
	return _aceb
}
func (_gbfbe fontCommon) isCIDFont() bool {
	if _gbfbe._dcbc == "" {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0069\u0073\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u002e\u0020\u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gbfbe)
	}
	_egfb := false
	switch _gbfbe._dcbc {
	case "\u0054\u0079\u0070e\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_egfb = true
	}
	_ad.Log.Trace("i\u0073\u0043\u0049\u0044\u0046\u006fn\u0074\u003a\u0020\u0069\u0073\u0043\u0049\u0044\u003d%\u0074\u0020\u0066o\u006et\u003d\u0025\u0073", _egfb, _gbfbe)
	return _egfb
}

// A returns the value of the A component of the color.
func (_bbde *PdfColorCalRGB) A() float64 { return _bbde[0] }

// NewPdfSignature creates a new PdfSignature object.
func NewPdfSignature(handler SignatureHandler) *PdfSignature {
	_bfdff := &PdfSignature{Type: _cde.MakeName("\u0053\u0069\u0067"), Handler: handler}
	_dfdce := &pdfSignDictionary{PdfObjectDictionary: _cde.MakeDict(), _addcc: &handler, _afafa: _bfdff}
	_bfdff._cabd = _cde.MakeIndirectObject(_dfdce)
	return _bfdff
}

// PdfColorspaceSpecialSeparation is a Separation colorspace.
// At the moment the colour space is set to a Separation space, the conforming reader shall determine whether the
// device has an available colorant (e.g. dye) corresponding to the name of the requested space. If so, the conforming
// reader shall ignore the alternateSpace and tintTransform parameters; subsequent painting operations within the
// space shall apply the designated colorant directly, according to the tint values supplied.
//
// Format: [/Separation name alternateSpace tintTransform]
type PdfColorspaceSpecialSeparation struct {
	ColorantName   *_cde.PdfObjectName
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	_bcdc          *_cde.PdfIndirectObject
}

// GetNumComponents returns the number of color components.
func (_cdba *PdfColorspaceICCBased) GetNumComponents() int { return _cdba.N }

// ContentStreamWrapper wraps the Page's contentstream into q ... Q blocks.
type ContentStreamWrapper interface{ WrapContentStream(_gaeg *PdfPage) error }

func (_abdd *PdfReader) newPdfAnnotationFromIndirectObject(_gfec *_cde.PdfIndirectObject) (*PdfAnnotation, error) {
	_aff, _gaa := _gfec.PdfObject.(*_cde.PdfObjectDictionary)
	if !_gaa {
		return nil, _ee.Errorf("\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _abdd._bedfa.GetModelFromPrimitive(_aff); model != nil {
		_baed, _fdef := model.(*PdfAnnotation)
		if !_fdef {
			return nil, _ee.Errorf("\u0063\u0061\u0063\u0068\u0065\u0064 \u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050D\u0046\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e")
		}
		return _baed, nil
	}
	_ddbe := &PdfAnnotation{}
	_ddbe._bddg = _gfec
	_abdd._bedfa.Register(_aff, _ddbe)
	if _dae := _aff.Get("\u0054\u0079\u0070\u0065"); _dae != nil {
		_gddf, _fdg := _dae.(*_cde.PdfObjectName)
		if !_fdg {
			_ad.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _dae)
		} else {
			if *_gddf != "\u0041\u006e\u006eo\u0074" {
				_ad.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070\u0065\u0063\u0074\u0065d\u0020\u0054\u0079\u0070\u0065\u0020\u0021=\u0020\u0041\u006e\u006e\u006f\u0074\u0020\u0028\u0025\u0073\u0029", *_gddf)
			}
		}
	}
	if _feab := _aff.Get("\u0052\u0065\u0063\u0074"); _feab != nil {
		_ddbe.Rect = _feab
	}
	if _fbb := _aff.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _fbb != nil {
		_ddbe.Contents = _fbb
	}
	if _bgg := _aff.Get("\u0050"); _bgg != nil {
		_ddbe.P = _bgg
	}
	if _bddgc := _aff.Get("\u004e\u004d"); _bddgc != nil {
		_ddbe.NM = _bddgc
	}
	if _gefd := _aff.Get("\u004d"); _gefd != nil {
		_ddbe.M = _gefd
	}
	if _dacf := _aff.Get("\u0046"); _dacf != nil {
		_ddbe.F = _dacf
	}
	if _aedb := _aff.Get("\u0041\u0050"); _aedb != nil {
		_ddbe.AP = _aedb
	}
	if _edeb := _aff.Get("\u0041\u0053"); _edeb != nil {
		_ddbe.AS = _edeb
	}
	if _aaf := _aff.Get("\u0042\u006f\u0072\u0064\u0065\u0072"); _aaf != nil {
		_ddbe.Border = _aaf
	}
	if _begd := _aff.Get("\u0043"); _begd != nil {
		_ddbe.C = _begd
	}
	if _gaga := _aff.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074"); _gaga != nil {
		_ddbe.StructParent = _gaga
	}
	if _ggdc := _aff.Get("\u004f\u0043"); _ggdc != nil {
		_ddbe.OC = _ggdc
	}
	_fae := _aff.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")
	if _fae == nil {
		_ad.Log.Debug("\u0057\u0041\u0052\u004e\u0049\u004e\u0047:\u0020\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079 \u0069s\u0073\u0075\u0065\u0020\u002d\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073u\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074\u0079p\u0065")
		_ddbe._bea = nil
		return _ddbe, nil
	}
	_fee, _aad := _fae.(*_cde.PdfObjectName)
	if !_aad {
		_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065 !\u003d\u0020n\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _fae)
		return nil, _ee.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d n\u0061\u006d\u0065 \u0028%\u0054\u0029", _fae)
	}
	switch *_fee {
	case "\u0054\u0065\u0078\u0074":
		_agbd, _bce := _abdd.newPdfAnnotationTextFromDict(_aff)
		if _bce != nil {
			return nil, _bce
		}
		_agbd.PdfAnnotation = _ddbe
		_ddbe._bea = _agbd
		return _ddbe, nil
	case "\u004c\u0069\u006e\u006b":
		_aga, _fdbe := _abdd.newPdfAnnotationLinkFromDict(_aff)
		if _fdbe != nil {
			return nil, _fdbe
		}
		_aga.PdfAnnotation = _ddbe
		_ddbe._bea = _aga
		return _ddbe, nil
	case "\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074":
		_cbb, _gdf := _abdd.newPdfAnnotationFreeTextFromDict(_aff)
		if _gdf != nil {
			return nil, _gdf
		}
		_cbb.PdfAnnotation = _ddbe
		_ddbe._bea = _cbb
		return _ddbe, nil
	case "\u004c\u0069\u006e\u0065":
		_cggg, _feee := _abdd.newPdfAnnotationLineFromDict(_aff)
		if _feee != nil {
			return nil, _feee
		}
		_cggg.PdfAnnotation = _ddbe
		_ddbe._bea = _cggg
		_ad.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041N\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0061\u006e\u006eo\u0074\u0020\u0028\u0025\u0054\u0029\u003a \u0025\u002b\u0076\u000a", _ddbe, _ddbe)
		_ad.Log.Trace("\u004c\u0049\u004eE\u0020\u0041\u004e\u004eO\u0054\u0041\u0054\u0049\u004f\u004e\u003a \u0063\u0074\u0078\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u002b\u0076\u000a", _cggg, _cggg)
		_ad.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041\u004e\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u0020\u004d\u0061\u0072\u006b\u0075\u0070\u003a\u0020c\u0074\u0078\u0020\u0028\u0025T\u0029\u003a \u0025\u002b\u0076\u000a", _cggg.PdfAnnotationMarkup, _cggg.PdfAnnotationMarkup)
		return _ddbe, nil
	case "\u0053\u0071\u0075\u0061\u0072\u0065":
		_bggf, _adbg := _abdd.newPdfAnnotationSquareFromDict(_aff)
		if _adbg != nil {
			return nil, _adbg
		}
		_bggf.PdfAnnotation = _ddbe
		_ddbe._bea = _bggf
		return _ddbe, nil
	case "\u0043\u0069\u0072\u0063\u006c\u0065":
		_cbab, _eacb := _abdd.newPdfAnnotationCircleFromDict(_aff)
		if _eacb != nil {
			return nil, _eacb
		}
		_cbab.PdfAnnotation = _ddbe
		_ddbe._bea = _cbab
		return _ddbe, nil
	case "\u0050o\u006c\u0079\u0067\u006f\u006e":
		_dcgfb, _bgdg := _abdd.newPdfAnnotationPolygonFromDict(_aff)
		if _bgdg != nil {
			return nil, _bgdg
		}
		_dcgfb.PdfAnnotation = _ddbe
		_ddbe._bea = _dcgfb
		return _ddbe, nil
	case "\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065":
		_fdf, _bffc := _abdd.newPdfAnnotationPolyLineFromDict(_aff)
		if _bffc != nil {
			return nil, _bffc
		}
		_fdf.PdfAnnotation = _ddbe
		_ddbe._bea = _fdf
		return _ddbe, nil
	case "\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t":
		_gfag, _gcgd := _abdd.newPdfAnnotationHighlightFromDict(_aff)
		if _gcgd != nil {
			return nil, _gcgd
		}
		_gfag.PdfAnnotation = _ddbe
		_ddbe._bea = _gfag
		return _ddbe, nil
	case "\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee":
		_bgab, _aab := _abdd.newPdfAnnotationUnderlineFromDict(_aff)
		if _aab != nil {
			return nil, _aab
		}
		_bgab.PdfAnnotation = _ddbe
		_ddbe._bea = _bgab
		return _ddbe, nil
	case "\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079":
		_gdbb, _gegd := _abdd.newPdfAnnotationSquigglyFromDict(_aff)
		if _gegd != nil {
			return nil, _gegd
		}
		_gdbb.PdfAnnotation = _ddbe
		_ddbe._bea = _gdbb
		return _ddbe, nil
	case "\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t":
		_cbff, _fefb := _abdd.newPdfAnnotationStrikeOut(_aff)
		if _fefb != nil {
			return nil, _fefb
		}
		_cbff.PdfAnnotation = _ddbe
		_ddbe._bea = _cbff
		return _ddbe, nil
	case "\u0043\u0061\u0072e\u0074":
		_feed, _fegf := _abdd.newPdfAnnotationCaretFromDict(_aff)
		if _fegf != nil {
			return nil, _fegf
		}
		_feed.PdfAnnotation = _ddbe
		_ddbe._bea = _feed
		return _ddbe, nil
	case "\u0053\u0074\u0061m\u0070":
		_fcc, _fbd := _abdd.newPdfAnnotationStampFromDict(_aff)
		if _fbd != nil {
			return nil, _fbd
		}
		_fcc.PdfAnnotation = _ddbe
		_ddbe._bea = _fcc
		return _ddbe, nil
	case "\u0049\u006e\u006b":
		_ggf, _fba := _abdd.newPdfAnnotationInkFromDict(_aff)
		if _fba != nil {
			return nil, _fba
		}
		_ggf.PdfAnnotation = _ddbe
		_ddbe._bea = _ggf
		return _ddbe, nil
	case "\u0050\u006f\u0070u\u0070":
		_adf, _fca := _abdd.newPdfAnnotationPopupFromDict(_aff)
		if _fca != nil {
			return nil, _fca
		}
		_adf.PdfAnnotation = _ddbe
		_ddbe._bea = _adf
		return _ddbe, nil
	case "\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074":
		_aaad, _fecbg := _abdd.newPdfAnnotationFileAttachmentFromDict(_aff)
		if _fecbg != nil {
			return nil, _fecbg
		}
		_aaad.PdfAnnotation = _ddbe
		_ddbe._bea = _aaad
		return _ddbe, nil
	case "\u0053\u006f\u0075n\u0064":
		_abfd, _efd := _abdd.newPdfAnnotationSoundFromDict(_aff)
		if _efd != nil {
			return nil, _efd
		}
		_abfd.PdfAnnotation = _ddbe
		_ddbe._bea = _abfd
		return _ddbe, nil
	case "\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a":
		_deaa, _baa := _abdd.newPdfAnnotationRichMediaFromDict(_aff)
		if _baa != nil {
			return nil, _baa
		}
		_deaa.PdfAnnotation = _ddbe
		_ddbe._bea = _deaa
		return _ddbe, nil
	case "\u004d\u006f\u0076i\u0065":
		_ffcg, _adgg := _abdd.newPdfAnnotationMovieFromDict(_aff)
		if _adgg != nil {
			return nil, _adgg
		}
		_ffcg.PdfAnnotation = _ddbe
		_ddbe._bea = _ffcg
		return _ddbe, nil
	case "\u0053\u0063\u0072\u0065\u0065\u006e":
		_ebg, _bcg := _abdd.newPdfAnnotationScreenFromDict(_aff)
		if _bcg != nil {
			return nil, _bcg
		}
		_ebg.PdfAnnotation = _ddbe
		_ddbe._bea = _ebg
		return _ddbe, nil
	case "\u0057\u0069\u0064\u0067\u0065\u0074":
		_aaadg, _gefdc := _abdd.newPdfAnnotationWidgetFromDict(_aff)
		if _gefdc != nil {
			return nil, _gefdc
		}
		_aaadg.PdfAnnotation = _ddbe
		_ddbe._bea = _aaadg
		return _ddbe, nil
	case "P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b":
		_cbffc, _gffgg := _abdd.newPdfAnnotationPrinterMarkFromDict(_aff)
		if _gffgg != nil {
			return nil, _gffgg
		}
		_cbffc.PdfAnnotation = _ddbe
		_ddbe._bea = _cbffc
		return _ddbe, nil
	case "\u0054r\u0061\u0070\u004e\u0065\u0074":
		_faeg, _cfa := _abdd.newPdfAnnotationTrapNetFromDict(_aff)
		if _cfa != nil {
			return nil, _cfa
		}
		_faeg.PdfAnnotation = _ddbe
		_ddbe._bea = _faeg
		return _ddbe, nil
	case "\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k":
		_ccb, _bcda := _abdd.newPdfAnnotationWatermarkFromDict(_aff)
		if _bcda != nil {
			return nil, _bcda
		}
		_ccb.PdfAnnotation = _ddbe
		_ddbe._bea = _ccb
		return _ddbe, nil
	case "\u0033\u0044":
		_edee, _adbf := _abdd.newPdfAnnotation3DFromDict(_aff)
		if _adbf != nil {
			return nil, _adbf
		}
		_edee.PdfAnnotation = _ddbe
		_ddbe._bea = _edee
		return _ddbe, nil
	case "\u0050\u0072\u006f\u006a\u0065\u0063\u0074\u0069\u006f\u006e":
		_abfc, _dbfc := _abdd.newPdfAnnotationProjectionFromDict(_aff)
		if _dbfc != nil {
			return nil, _dbfc
		}
		_abfc.PdfAnnotation = _ddbe
		_ddbe._bea = _abfc
		return _ddbe, nil
	case "\u0052\u0065\u0064\u0061\u0063\u0074":
		_ead, _afgd := _abdd.newPdfAnnotationRedactFromDict(_aff)
		if _afgd != nil {
			return nil, _afgd
		}
		_ead.PdfAnnotation = _ddbe
		_ddbe._bea = _ead
		return _ddbe, nil
	}
	_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069\u006fn\u003a\u0020\u0025\u0073", *_fee)
	return nil, nil
}

// GetVersion gets the document version.
func (_fadgd *PdfWriter) GetVersion() _cde.Version { return _fadgd._cgdcc }
func (_ebbe *LTV) enable(_efga, _bbfe []*_bg.Certificate, _cfdbd string) error {
	_efadb, _efdab, _degf := _ebbe.buildCertChain(_efga, _bbfe)
	if _degf != nil {
		return _degf
	}
	_adce, _degf := _ebbe.getCerts(_efadb)
	if _degf != nil {
		return _degf
	}
	var _gbgb, _bbcee [][]byte
	if _ebbe.OCSPClient != nil {
		_gbgb, _degf = _ebbe.getOCSPs(_efadb, _efdab)
		if _degf != nil {
			return _degf
		}
	}
	if _ebbe.CRLClient != nil {
		_bbcee, _degf = _ebbe.getCRLs(_efadb)
		if _degf != nil {
			return _degf
		}
	}
	_adfdd := _ebbe._geeeg
	_bacad, _degf := _adfdd.addCerts(_adce)
	if _degf != nil {
		return _degf
	}
	_cgdd, _degf := _adfdd.addOCSPs(_gbgb)
	if _degf != nil {
		return _degf
	}
	_cadg, _degf := _adfdd.addCRLs(_bbcee)
	if _degf != nil {
		return _degf
	}
	if _cfdbd != "" {
		_adfdd.VRI[_cfdbd] = &VRI{Cert: _bacad, OCSP: _cgdd, CRL: _cadg}
	}
	_ebbe._ffab.SetDSS(_adfdd)
	return nil
}

// AddExtGState adds a graphics state to the XObject resources.
func (_cbggc *PdfPage) AddExtGState(name _cde.PdfObjectName, egs *_cde.PdfObjectDictionary) error {
	if _cbggc.Resources == nil {
		_cbggc.Resources = NewPdfPageResources()
	}
	if _cbggc.Resources.ExtGState == nil {
		_cbggc.Resources.ExtGState = _cde.MakeDict()
	}
	_gggeg, _bccc := _cde.TraceToDirectObject(_cbggc.Resources.ExtGState).(*_cde.PdfObjectDictionary)
	if !_bccc {
		_ad.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _cde.TraceToDirectObject(_cbggc.Resources.ExtGState))
		return _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gggeg.Set(name, egs)
	return nil
}

// GetNameDictionary returns the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_fadae *PdfReader) GetNameDictionary() (_cde.PdfObject, error) {
	_aabfcg := _cde.ResolveReference(_fadae._efabe.Get("\u004e\u0061\u006de\u0073"))
	if _aabfcg == nil {
		return nil, nil
	}
	if !_fadae._cdgee {
		_egbea := _fadae.traverseObjectData(_aabfcg)
		if _egbea != nil {
			return nil, _egbea
		}
	}
	return _aabfcg, nil
}
func (_ceed fontCommon) asPdfObjectDictionary(_faae string) *_cde.PdfObjectDictionary {
	if _faae != "" && _ceed._dcbc != "" && _faae != _ceed._dcbc {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061\u0073\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020O\u0076\u0065\u0072\u0072\u0069\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0062t\u0079\u0070\u0065\u0020\u0074\u006f \u0025\u0023\u0071 \u0025\u0073", _faae, _ceed)
	} else if _faae == "" && _ceed._dcbc == "" {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061s\u0050\u0064\u0066Ob\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074y\u0070\u0065\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _ceed)
	} else if _ceed._dcbc == "" {
		_ceed._dcbc = _faae
	}
	_ceec := _cde.MakeDict()
	_ceec.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0046\u006f\u006e\u0074"))
	_ceec.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _cde.MakeName(_ceed._eeab))
	_ceec.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName(_ceed._dcbc))
	if _ceed._fagf != nil {
		_ceec.Set("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072", _ceed._fagf.ToPdfObject())
	}
	if _ceed._dfae != nil {
		_ceec.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _ceed._dfae)
	} else if _ceed._ggebg != nil {
		_efcee, _bgdgg := _ceed._ggebg.Stream()
		if _bgdgg != nil {
			_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0067\u0065\u0074\u0020C\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0065r\u0072\u003d\u0025\u0076", _bgdgg)
		} else {
			_ceec.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _efcee)
		}
	}
	return _ceec
}

// PdfActionGoToR represents a GoToR action.
type PdfActionGoToR struct {
	*PdfAction
	F         *PdfFilespec
	D         _cde.PdfObject
	NewWindow _cde.PdfObject
}

func (_gce *PdfReader) newPdfActionSetOCGStateFromDict(_bbaa *_cde.PdfObjectDictionary) (*PdfActionSetOCGState, error) {
	return &PdfActionSetOCGState{State: _bbaa.Get("\u0053\u0074\u0061t\u0065"), PreserveRB: _bbaa.Get("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042")}, nil
}

// AddImageResource adds an image to the XObject resources.
func (_bfcgf *PdfPage) AddImageResource(name _cde.PdfObjectName, ximg *XObjectImage) error {
	var _bdde *_cde.PdfObjectDictionary
	if _bfcgf.Resources.XObject == nil {
		_bdde = _cde.MakeDict()
		_bfcgf.Resources.XObject = _bdde
	} else {
		var _fabbg bool
		_bdde, _fabbg = (_bfcgf.Resources.XObject).(*_cde.PdfObjectDictionary)
		if !_fabbg {
			return _ceg.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0078\u0072\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020\u0074\u0079p\u0065")
		}
	}
	_bdde.Set(name, ximg.ToPdfObject())
	return nil
}

// ToPdfObject returns the text field dictionary within an indirect object (container).
func (_eggb *PdfFieldText) ToPdfObject() _cde.PdfObject {
	_eggb.PdfField.ToPdfObject()
	_gaaec := _eggb._afgc
	_bacb := _gaaec.PdfObject.(*_cde.PdfObjectDictionary)
	_bacb.Set("\u0046\u0054", _cde.MakeName("\u0054\u0078"))
	if _eggb.DA != nil {
		_bacb.Set("\u0044\u0041", _eggb.DA)
	}
	if _eggb.Q != nil {
		_bacb.Set("\u0051", _eggb.Q)
	}
	if _eggb.DS != nil {
		_bacb.Set("\u0044\u0053", _eggb.DS)
	}
	if _eggb.RV != nil {
		_bacb.Set("\u0052\u0056", _eggb.RV)
	}
	if _eggb.MaxLen != nil {
		_bacb.Set("\u004d\u0061\u0078\u004c\u0065\u006e", _eggb.MaxLen)
	}
	return _gaaec
}

// HasColorspaceByName checks if the colorspace with the specified name exists in the page resources.
func (_cbefb *PdfPageResources) HasColorspaceByName(keyName _cde.PdfObjectName) bool {
	_cefc, _beggg := _cbefb.GetColorspaces()
	if _beggg != nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _beggg)
		return false
	}
	if _cefc == nil {
		return false
	}
	_, _bfege := _cefc.Colorspaces[string(keyName)]
	return _bfege
}

// Encoder returns the font's text encoder.
func (_abgdc *PdfFont) Encoder() _gc.TextEncoder {
	_eecde := _abgdc.actualFont()
	if _eecde == nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0045n\u0063\u006f\u0064er\u0020\u006e\u006f\u0074\u0020\u0069m\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066o\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d%\u0023\u0054", _abgdc._gbcff)
		return nil
	}
	return _eecde.Encoder()
}
func (_dfgc *LTV) getCRLs(_cgbc []*_bg.Certificate) ([][]byte, error) {
	_gdfd := make([][]byte, 0, len(_cgbc))
	for _, _agagg := range _cgbc {
		for _, _ceag := range _agagg.CRLDistributionPoints {
			if _dfgc.CertClient.IsCA(_agagg) {
				continue
			}
			_ccfc, _dfce := _dfgc.CRLClient.MakeRequest(_ceag, _agagg)
			if _dfce != nil {
				_ad.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _dfce)
				continue
			}
			_gdfd = append(_gdfd, _ccfc)
		}
	}
	return _gdfd, nil
}
func (_ddg *PdfReader) newPdfActionRenditionFromDict(_afed *_cde.PdfObjectDictionary) (*PdfActionRendition, error) {
	return &PdfActionRendition{R: _afed.Get("\u0052"), AN: _afed.Get("\u0041\u004e"), OP: _afed.Get("\u004f\u0050"), JS: _afed.Get("\u004a\u0053")}, nil
}

// NewPdfActionRendition returns a new "rendition" action.
func NewPdfActionRendition() *PdfActionRendition {
	_gec := NewPdfAction()
	_ec := &PdfActionRendition{}
	_ec.PdfAction = _gec
	_gec.SetContext(_ec)
	return _ec
}

// PdfActionThread represents a thread action.
type PdfActionThread struct {
	*PdfAction
	F *PdfFilespec
	D _cde.PdfObject
	B _cde.PdfObject
}

const (
	XObjectTypeUndefined XObjectType = iota
	XObjectTypeImage
	XObjectTypeForm
	XObjectTypePS
	XObjectTypeUnknown
)

// Subtype returns the font's "Subtype" field.
func (_gfcd *PdfFont) Subtype() string {
	_cgcb := _gfcd.baseFields()._dcbc
	if _ecbc, _bdca := _gfcd._gbcff.(*pdfFontType0); _bdca {
		_cgcb = _cgcb + "\u003a" + _ecbc.DescendantFont.Subtype()
	}
	return _cgcb
}
func (_ggba *fontFile) loadFromSegments(_aacbg, _aadge []byte) error {
	_ad.Log.Trace("\u006c\u006f\u0061dF\u0072\u006f\u006d\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u0020\u0025\u0064", len(_aacbg), len(_aadge))
	_ecded := _ggba.parseASCIIPart(_aacbg)
	if _ecded != nil {
		return _ecded
	}
	_ad.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _ggba)
	if len(_aadge) == 0 {
		return nil
	}
	_ad.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _ggba)
	return nil
}

// GetContentStreamObjs returns a slice of PDF objects containing the content
// streams of the page.
func (_acgc *PdfPage) GetContentStreamObjs() []_cde.PdfObject {
	if _acgc.Contents == nil {
		return nil
	}
	_cadgb := _cde.TraceToDirectObject(_acgc.Contents)
	if _cbdgf, _dcef := _cadgb.(*_cde.PdfObjectArray); _dcef {
		return _cbdgf.Elements()
	}
	return []_cde.PdfObject{_cadgb}
}

// FieldValueProvider provides field values from a data source such as FDF, JSON or any other.
type FieldValueProvider interface {
	FieldValues() (map[string]_cde.PdfObject, error)
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
	_acbcb          *_cde.PdfObjectDictionary
}

// SignatureHandler interface defines the common functionality for PDF signature handlers, which
// need to be capable of validating digital signatures and signing PDF documents.
type SignatureHandler interface {

	// IsApplicable checks if a given signature dictionary `sig` is applicable for the signature handler.
	// For example a signature of type `adbe.pkcs7.detached` might not fit for a rsa.sha1 handler.
	IsApplicable(_dccbb *PdfSignature) bool

	// Validate validates a PDF signature against a given digest (hash) such as that determined
	// for an input file. Returns validation results.
	Validate(_fabbf *PdfSignature, _cgfb Hasher) (SignatureValidationResult, error)

	// InitSignature prepares the signature dictionary for signing. This involves setting all
	// necessary fields, and also allocating sufficient space to the Contents so that the
	// finalized signature can be inserted once the hash is calculated.
	InitSignature(_afbca *PdfSignature) error

	// NewDigest creates a new digest/hasher based on the signature dictionary and handler.
	NewDigest(_bgcgg *PdfSignature) (Hasher, error)

	// Sign receives the hash `digest` (for example hash of an input file), and signs based
	// on the signature dictionary `sig` and applies the signature data to the signature
	// dictionary Contents field.
	Sign(_caca *PdfSignature, _fbaecg Hasher) error
}

func _fggeg(_cdgde *_cde.PdfObjectDictionary) (*PdfShadingType5, error) {
	_babca := PdfShadingType5{}
	_begagc := _cdgde.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _begagc == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_aaefe, _ebgda := _begagc.(*_cde.PdfObjectInteger)
	if !_ebgda {
		_ad.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _begagc)
		return nil, _cde.ErrTypeError
	}
	_babca.BitsPerCoordinate = _aaefe
	_begagc = _cdgde.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _begagc == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_aaefe, _ebgda = _begagc.(*_cde.PdfObjectInteger)
	if !_ebgda {
		_ad.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _begagc)
		return nil, _cde.ErrTypeError
	}
	_babca.BitsPerComponent = _aaefe
	_begagc = _cdgde.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077")
	if _begagc == nil {
		_ad.Log.Debug("\u0052\u0065\u0071u\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0056\u0065\u0072\u0074\u0069c\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077")
		return nil, ErrRequiredAttributeMissing
	}
	_aaefe, _ebgda = _begagc.(*_cde.PdfObjectInteger)
	if !_ebgda {
		_ad.Log.Debug("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006et\u0065\u0067\u0065\u0072\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _begagc)
		return nil, _cde.ErrTypeError
	}
	_babca.VerticesPerRow = _aaefe
	_begagc = _cdgde.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _begagc == nil {
		_ad.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_afbd, _ebgda := _begagc.(*_cde.PdfObjectArray)
	if !_ebgda {
		_ad.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _begagc)
		return nil, _cde.ErrTypeError
	}
	_babca.Decode = _afbd
	if _cccef := _cdgde.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _cccef != nil {
		_babca.Function = []PdfFunction{}
		if _cefdb, _cbbdb := _cccef.(*_cde.PdfObjectArray); _cbbdb {
			for _, _dabga := range _cefdb.Elements() {
				_eaff, _ddbff := _cfdbb(_dabga)
				if _ddbff != nil {
					_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _ddbff)
					return nil, _ddbff
				}
				_babca.Function = append(_babca.Function, _eaff)
			}
		} else {
			_efgad, _accc := _cfdbb(_cccef)
			if _accc != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _accc)
				return nil, _accc
			}
			_babca.Function = append(_babca.Function, _efgad)
		}
	}
	return &_babca, nil
}

// NewPdfActionURI returns a new "Uri" action.
func NewPdfActionURI() *PdfActionURI {
	_caec := NewPdfAction()
	_bbc := &PdfActionURI{}
	_bbc.PdfAction = _caec
	_caec.SetContext(_bbc)
	return _bbc
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_bcdb *PdfColorspaceDeviceCMYK) ToPdfObject() _cde.PdfObject {
	return _cde.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b")
}
func (_aded *PdfReader) newPdfAnnotationRedactFromDict(_edaa *_cde.PdfObjectDictionary) (*PdfAnnotationRedact, error) {
	_abbf := PdfAnnotationRedact{}
	_gggg, _dfga := _aded.newPdfAnnotationMarkupFromDict(_edaa)
	if _dfga != nil {
		return nil, _dfga
	}
	_abbf.PdfAnnotationMarkup = _gggg
	_abbf.QuadPoints = _edaa.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_abbf.IC = _edaa.Get("\u0049\u0043")
	_abbf.RO = _edaa.Get("\u0052\u004f")
	_abbf.OverlayText = _edaa.Get("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074")
	_abbf.Repeat = _edaa.Get("\u0052\u0065\u0070\u0065\u0061\u0074")
	_abbf.DA = _edaa.Get("\u0044\u0041")
	_abbf.Q = _edaa.Get("\u0051")
	return &_abbf, nil
}

// CharcodesToUnicode converts the character codes `charcodes` to a slice of runes.
// How it works:
//  1) Use the ToUnicode CMap if there is one.
//  2) Use the underlying font's encoding.
func (_dgfd *PdfFont) CharcodesToUnicode(charcodes []_gc.CharCode) []rune {
	_dbdd, _, _ := _dgfd.CharcodesToUnicodeWithStats(charcodes)
	return _dbdd
}

// Inspect inspects the object types, subtypes and content in the PDF file returning a map of
// object type to number of instances of each.
func (_eaecg *PdfReader) Inspect() (map[string]int, error) { return _eaecg._aggcgb.Inspect() }
func (_gbe *PdfReader) newPdfAnnotationWatermarkFromDict(_bbb *_cde.PdfObjectDictionary) (*PdfAnnotationWatermark, error) {
	_dacc := PdfAnnotationWatermark{}
	_dacc.FixedPrint = _bbb.Get("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074")
	return &_dacc, nil
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
	ColorToRGB(_dddff PdfColor) (PdfColor, error)

	// GetNumComponents returns the number of components in the PdfColorspace.
	GetNumComponents() int

	// ToPdfObject returns a PdfObject representation of the PdfColorspace.
	ToPdfObject() _cde.PdfObject

	// ColorFromPdfObjects returns a PdfColor in the given PdfColorspace from an array of PdfObject where each
	// PdfObject represents a numeric value.
	ColorFromPdfObjects(_dcge []_cde.PdfObject) (PdfColor, error)

	// ColorFromFloats returns a new PdfColor based on input color components for a given PdfColorspace.
	ColorFromFloats(_cgge []float64) (PdfColor, error)

	// DecodeArray returns the Decode array for the PdfColorSpace, i.e. the range of each component.
	DecodeArray() []float64
}

func (_feda *fontFile) parseASCIIPart(_eaba []byte) error {
	if len(_eaba) < 2 || string(_eaba[:2]) != "\u0025\u0021" {
		return _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0074a\u0072\u0074\u0020\u006f\u0066\u0020\u0041S\u0043\u0049\u0049\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_cdag, _gcgbfb, _fgfag := _dggba(_eaba)
	if _fgfag != nil {
		return _fgfag
	}
	_gdgfa := _ebag(_cdag)
	_feda._fbffa = _gdgfa["\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"]
	if _feda._fbffa == "" {
		_ad.Log.Debug("\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0020\u0068a\u0073\u0020\u006e\u006f\u0020\u002f\u0046\u006f\u006e\u0074N\u0061\u006d\u0065")
	}
	if _gcgbfb != "" {
		_gdgd, _adgge := _bbefa(_gcgbfb)
		if _adgge != nil {
			return _adgge
		}
		_egeed, _adgge := _gc.NewCustomSimpleTextEncoder(_gdgd, nil)
		if _adgge != nil {
			_ad.Log.Debug("\u0045\u0052\u0052\u004fR\u0020\u003a\u0055\u004e\u004b\u004e\u004f\u0057\u004e\u0020G\u004cY\u0050\u0048\u003a\u0020\u0065\u0072\u0072=\u0025\u0076", _adgge)
			return nil
		}
		_feda._cafea = _egeed
	}
	return nil
}

// GetCharMetrics returns the char metrics for character code `code`.
// How it works:
//  1) It calls the GetCharMetrics function for the underlying font, either a simple font or
//     a Type0 font. The underlying font GetCharMetrics() functions do direct charcode ➞  metrics
//     mappings.
//  2) If the underlying font's GetCharMetrics() doesn't have a CharMetrics for `code` then a
//     a CharMetrics with the FontDescriptor's /MissingWidth is returned.
//  3) If there is no /MissingWidth then a failure is returned.
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
//                       well give them 0 width. There is no need for the bool return.
// TODO(gunnsth): Reconsider whether needed or if can map via GlyphName.
func (_ecgb *PdfFont) GetCharMetrics(code _gc.CharCode) (CharMetrics, bool) {
	var _ffbb _fe.CharMetrics
	switch _eaec := _ecgb._gbcff.(type) {
	case *pdfFontSimple:
		if _efdbf, _gcefe := _eaec.GetCharMetrics(code); _gcefe {
			return _efdbf, _gcefe
		}
	case *pdfFontType0:
		if _bcgcf, _eddg := _eaec.GetCharMetrics(code); _eddg {
			return _bcgcf, _eddg
		}
	case *pdfCIDFontType0:
		if _bdeg, _cccec := _eaec.GetCharMetrics(code); _cccec {
			return _bdeg, _cccec
		}
	case *pdfCIDFontType2:
		if _ebcag, _dacff := _eaec.GetCharMetrics(code); _dacff {
			return _ebcag, _dacff
		}
	case *pdfFontType3:
		if _ecafg, _dcaef := _eaec.GetCharMetrics(code); _dcaef {
			return _ecafg, _dcaef
		}
	default:
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020G\u0065\u0074\u0043h\u0061\u0072\u004de\u0074\u0072i\u0063\u0073\u0020\u006e\u006f\u0074 \u0069mp\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d\u0025\u0054\u002e", _ecgb._gbcff)
		return _ffbb, false
	}
	if _badg, _efff := _ecgb.GetFontDescriptor(); _efff == nil && _badg != nil {
		return _fe.CharMetrics{Wx: _badg._efgg}, true
	}
	_ad.Log.Debug("\u0047\u0065\u0074\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069c\u0073\u0020\u0066\u006f\u0072 \u0066\u006fn\u0074\u003d\u0025\u0073", _ecgb)
	return _ffbb, false
}
func (_ggbb *PdfFunctionType0) processSamples() error {
	_aade := _cae.ResampleBytes(_ggbb._gfbb, _ggbb.BitsPerSample)
	_ggbb._bgacg = _aade
	return nil
}
func (_fgbge *PdfWriter) writeXRefStreams(_edcef int, _dddfb int64) error {
	_ccge := _edcef + 1
	_fgbge._gfdac[_ccge] = crossReference{Type: 1, ObjectNumber: _ccge, Offset: _dddfb}
	_fdegg := _ede.NewBuffer(nil)
	_cbgb := _cde.MakeArray()
	for _ggafdd := 0; _ggafdd <= _edcef; {
		for ; _ggafdd <= _edcef; _ggafdd++ {
			_eefdb, _fbegc := _fgbge._gfdac[_ggafdd]
			if _fbegc && (!_fgbge._aabfe || _fgbge._aabfe && (_eefdb.Type == 1 && _eefdb.Offset >= _fgbge._fgged || _eefdb.Type == 0)) {
				break
			}
		}
		var _cagce int
		for _cagce = _ggafdd + 1; _cagce <= _edcef; _cagce++ {
			_edgbd, _efbbc := _fgbge._gfdac[_cagce]
			if _efbbc && (!_fgbge._aabfe || _fgbge._aabfe && (_edgbd.Type == 1 && _edgbd.Offset > _fgbge._fgged)) {
				continue
			}
			break
		}
		_cbgb.Append(_cde.MakeInteger(int64(_ggafdd)), _cde.MakeInteger(int64(_cagce-_ggafdd)))
		for _gaedfcb := _ggafdd; _gaedfcb < _cagce; _gaedfcb++ {
			_ccee := _fgbge._gfdac[_gaedfcb]
			switch _ccee.Type {
			case 0:
				_ba.Write(_fdegg, _ba.BigEndian, byte(0))
				_ba.Write(_fdegg, _ba.BigEndian, uint32(0))
				_ba.Write(_fdegg, _ba.BigEndian, uint16(0xFFFF))
			case 1:
				_ba.Write(_fdegg, _ba.BigEndian, byte(1))
				_ba.Write(_fdegg, _ba.BigEndian, uint32(_ccee.Offset))
				_ba.Write(_fdegg, _ba.BigEndian, uint16(_ccee.Generation))
			case 2:
				_ba.Write(_fdegg, _ba.BigEndian, byte(2))
				_ba.Write(_fdegg, _ba.BigEndian, uint32(_ccee.ObjectNumber))
				_ba.Write(_fdegg, _ba.BigEndian, uint16(_ccee.Index))
			}
		}
		_ggafdd = _cagce + 1
	}
	_eafbd, _fedd := _cde.MakeStream(_fdegg.Bytes(), _cde.NewFlateEncoder())
	if _fedd != nil {
		return _fedd
	}
	_eafbd.ObjectNumber = int64(_ccge)
	_eafbd.PdfObjectDictionary.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0058\u0052\u0065\u0066"))
	_eafbd.PdfObjectDictionary.Set("\u0057", _cde.MakeArray(_cde.MakeInteger(1), _cde.MakeInteger(4), _cde.MakeInteger(2)))
	_eafbd.PdfObjectDictionary.Set("\u0049\u006e\u0064e\u0078", _cbgb)
	_eafbd.PdfObjectDictionary.Set("\u0053\u0069\u007a\u0065", _cde.MakeInteger(int64(_ccge+1)))
	_eafbd.PdfObjectDictionary.Set("\u0049\u006e\u0066\u006f", _fgbge._fdgbc)
	_eafbd.PdfObjectDictionary.Set("\u0052\u006f\u006f\u0074", _fgbge._eacge)
	if _fgbge._aabfe && _fgbge._gbbda > 0 {
		_eafbd.PdfObjectDictionary.Set("\u0050\u0072\u0065\u0076", _cde.MakeInteger(_fgbge._gbbda))
	}
	if _fgbge._ccgbe != nil {
		_eafbd.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _fgbge._bdbdb)
	}
	if _fgbge._bgacb == nil && _fgbge._gfegf != "" && _fgbge._dcace != "" {
		_fgbge._bgacb = _cde.MakeArray(_cde.MakeHexString(_fgbge._gfegf), _cde.MakeHexString(_fgbge._dcace))
	}
	if _fgbge._bgacb != nil {
		_ad.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _fgbge._bgacb)
		_eafbd.Set("\u0049\u0044", _fgbge._bgacb)
	}
	_fgbge.writeObject(int(_eafbd.ObjectNumber), _eafbd)
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_ebeae *PdfAnnotationCircle) ToPdfObject() _cde.PdfObject {
	_ebeae.PdfAnnotation.ToPdfObject()
	_ebff := _ebeae._bddg
	_afce := _ebff.PdfObject.(*_cde.PdfObjectDictionary)
	_ebeae.PdfAnnotationMarkup.appendToPdfDictionary(_afce)
	_afce.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0043\u0069\u0072\u0063\u006c\u0065"))
	_afce.SetIfNotNil("\u0042\u0053", _ebeae.BS)
	_afce.SetIfNotNil("\u0049\u0043", _ebeae.IC)
	_afce.SetIfNotNil("\u0042\u0045", _ebeae.BE)
	_afce.SetIfNotNil("\u0052\u0044", _ebeae.RD)
	return _ebff
}
func (_adfec fontCommon) coreString() string {
	_acbaf := ""
	if _adfec._fagf != nil {
		_acbaf = _adfec._fagf.String()
	}
	return _ee.Sprintf("\u0025#\u0071\u0020%\u0023\u0071\u0020%\u0071\u0020\u006f\u0062\u006a\u003d\u0025d\u0020\u0054\u006f\u0055\u006e\u0069c\u006f\u0064\u0065\u003d\u0025\u0074\u0020\u0066\u006c\u0061\u0067s\u003d\u0030\u0078\u0025\u0030\u0078\u0020\u0025\u0073", _adfec._dcbc, _adfec._eeab, _adfec._fddeb, _adfec._dbadd, _adfec._dfae != nil, _adfec.fontFlags(), _acbaf)
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_ebadg pdfFontType0) GetCharMetrics(code _gc.CharCode) (_fe.CharMetrics, bool) {
	if _ebadg.DescendantFont == nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _ebadg)
		return _fe.CharMetrics{}, false
	}
	return _ebadg.DescendantFont.GetCharMetrics(code)
}

// PdfFunctionType3 defines stitching of the subdomains of several 1-input functions to produce
// a single new 1-input function.
type PdfFunctionType3 struct {
	Domain    []float64
	Range     []float64
	Functions []PdfFunction
	Bounds    []float64
	Encode    []float64
	_fbbgd    *_cde.PdfIndirectObject
}

// SetSubtype sets the Subtype S for given PdfOutputIntent.
func (_ddceb *PdfOutputIntent) SetSubtype(subtype PdfOutputIntentType) error {
	if !subtype.IsValid() {
		return _ceg.New("\u0070\u0072o\u0076\u0069\u0064\u0065d\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u004f\u0075t\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0053\u0075b\u0054\u0079\u0070\u0065")
	}
	_ddceb.S = subtype
	return nil
}

// PdfActionMovie represents a movie action.
type PdfActionMovie struct {
	*PdfAction
	Annotation _cde.PdfObject
	T          _cde.PdfObject
	Operation  _cde.PdfObject
}

// SetFilter sets compression filter. Decodes with current filter sets and
// encodes the data with the new filter.
func (_efcdg *XObjectImage) SetFilter(encoder _cde.StreamEncoder) error {
	_gfdga := _efcdg.Stream
	_baba, _fbgag := _efcdg.Filter.DecodeBytes(_gfdga)
	if _fbgag != nil {
		return _fbgag
	}
	_efcdg.Filter = encoder
	encoder.UpdateParams(_efcdg.getParamsDict())
	_gfdga, _fbgag = encoder.EncodeBytes(_baba)
	if _fbgag != nil {
		return _fbgag
	}
	_efcdg.Stream = _gfdga
	return nil
}

// NewPdfActionResetForm returns a new "reset form" action.
func NewPdfActionResetForm() *PdfActionResetForm {
	_dad := NewPdfAction()
	_bdc := &PdfActionResetForm{}
	_bdc.PdfAction = _dad
	_dad.SetContext(_bdc)
	return _bdc
}

// NewPdfAnnotationWidget returns an initialized annotation widget.
func NewPdfAnnotationWidget() *PdfAnnotationWidget {
	_acf := NewPdfAnnotation()
	_dag := &PdfAnnotationWidget{}
	_dag.PdfAnnotation = _acf
	_acf.SetContext(_dag)
	return _dag
}

// ToPdfObject implements interface PdfModel.
func (_ccab *PdfAnnotationSquiggly) ToPdfObject() _cde.PdfObject {
	_ccab.PdfAnnotation.ToPdfObject()
	_ebgg := _ccab._bddg
	_dfcdb := _ebgg.PdfObject.(*_cde.PdfObjectDictionary)
	_ccab.PdfAnnotationMarkup.appendToPdfDictionary(_dfcdb)
	_dfcdb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079"))
	_dfcdb.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _ccab.QuadPoints)
	return _ebgg
}

// GetRevision returns the specific version of the PdfReader for the current Pdf document
func (_ccfca *PdfReader) GetRevision(revisionNumber int) (*PdfReader, error) {
	_cfcbc := _ccfca._aggcgb.GetRevisionNumber()
	if revisionNumber < 0 || revisionNumber > _cfcbc {
		return nil, _ceg.New("w\u0072\u006f\u006e\u0067 r\u0065v\u0069\u0073\u0069\u006f\u006e \u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if revisionNumber == _cfcbc {
		return _ccfca, nil
	}
	if _ccfca._fgfbe[revisionNumber] != nil {
		return _ccfca._fgfbe[revisionNumber], nil
	}
	_gbgff := _ccfca
	for _ageeb := _cfcbc - 1; _ageeb >= revisionNumber; _ageeb-- {
		_gbage, _gafaaf := _gbgff.GetPreviousRevision()
		if _gafaaf != nil {
			return nil, _gafaaf
		}
		_ccfca._fgfbe[_ageeb] = _gbage
		_gbgff = _gbage
	}
	return _gbgff, nil
}
func _eabg(_eeae StdFontName) (pdfFontSimple, error) {
	_dabf, _dfbd := _fe.NewStdFontByName(_eeae)
	if !_dfbd {
		return pdfFontSimple{}, ErrFontNotSupported
	}
	_bggb := _gdaeg(_dabf)
	return _bggb, nil
}
func (_fbdda *pdfFontSimple) addEncoding() error {
	var (
		_dfag  string
		_agaad map[_gc.CharCode]_gc.GlyphName
		_defcb _gc.SimpleEncoder
	)
	if _fbdda.Encoder() != nil {
		_bfgg, _dfff := _fbdda.Encoder().(_gc.SimpleEncoder)
		if _dfff && _bfgg != nil {
			_dfag = _bfgg.BaseName()
		}
	}
	if _fbdda.Encoding != nil {
		_ffffa, _eadfc, _aabcg := _fbdda.getFontEncoding()
		if _aabcg != nil {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0073\u0065F\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0053u\u0062t\u0079\u0070\u0065\u003d\u0025\u0071\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003d\u0025\u0073 \u0028\u0025\u0054\u0029\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fbdda._eeab, _fbdda._dcbc, _fbdda.Encoding, _fbdda.Encoding, _aabcg)
			return _aabcg
		}
		if _ffffa != "" {
			_dfag = _ffffa
		}
		_agaad = _eadfc
		_defcb, _aabcg = _gc.NewSimpleTextEncoder(_dfag, _agaad)
		if _aabcg != nil {
			return _aabcg
		}
	}
	if _defcb == nil {
		_gcbda := _fbdda._fagf
		if _gcbda != nil {
			switch _fbdda._dcbc {
			case "\u0054\u0079\u0070e\u0031":
				if _gcbda.fontFile != nil && _gcbda.fontFile._cafea != nil {
					_ad.Log.Debug("\u0055\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006et\u0046\u0069\u006c\u0065")
					_defcb = _gcbda.fontFile._cafea
				}
			case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
				if _gcbda._ggdga != nil {
					_ad.Log.Debug("\u0055s\u0069n\u0067\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032")
					_gaedf, _dfad := _gcbda._ggdga.MakeEncoder()
					if _dfad == nil {
						_defcb = _gaedf
					}
				}
			}
		}
	}
	if _defcb != nil {
		if _agaad != nil {
			_ad.Log.Trace("\u0064\u0069\u0066fe\u0072\u0065\u006e\u0063\u0065\u0073\u003d\u0025\u002b\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _agaad, _fbdda.baseFields())
			_defcb = _gc.ApplyDifferences(_defcb, _agaad)
		}
		_fbdda.SetEncoder(_defcb)
	}
	return nil
}

// SetColorspaceByName adds the provided colorspace to the page resources.
func (_gccbce *PdfPageResources) SetColorspaceByName(keyName _cde.PdfObjectName, cs PdfColorspace) error {
	_cdfbe, _ddcbce := _gccbce.GetColorspaces()
	if _ddcbce != nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _ddcbce)
		return _ddcbce
	}
	if _cdfbe == nil {
		_cdfbe = NewPdfPageResourcesColorspaces()
		_gccbce.SetColorSpace(_cdfbe)
	}
	_cdfbe.Set(keyName, cs)
	return nil
}

// String returns the name of the colorspace (DeviceN).
func (_cabcc *PdfColorspaceDeviceN) String() string { return "\u0044e\u0076\u0069\u0063\u0065\u004e" }

// NewPdfAnnotationFileAttachment returns a new file attachment annotation.
func NewPdfAnnotationFileAttachment() *PdfAnnotationFileAttachment {
	_ace := NewPdfAnnotation()
	_gffg := &PdfAnnotationFileAttachment{}
	_gffg.PdfAnnotation = _ace
	_gffg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ace.SetContext(_gffg)
	return _gffg
}

// HasExtGState checks whether a font is defined by the specified keyName.
func (_fadee *PdfPageResources) HasExtGState(keyName _cde.PdfObjectName) bool {
	_, _bgedd := _fadee.GetFontByName(keyName)
	return _bgedd
}

// BorderEffect represents a border effect (Table 167 p. 395).
type BorderEffect int

// PdfColorspaceDeviceCMYK represents a CMYK32 colorspace.
type PdfColorspaceDeviceCMYK struct{}

// SetFlag sets the flag for the field.
func (_cbbea *PdfField) SetFlag(flag FieldFlag) { _cbbea.Ff = _cde.MakeInteger(int64(flag)) }

// FlattenFields flattens the form fields and annotations for the PDF loaded in `pdf` and makes
// non-editable.
// Looks up all widget annotations corresponding to form fields and flattens them by drawing the content
// through the content stream rather than annotations.
// References to flattened annotations will be removed from Page Annots array. For fields the AcroForm entry
// will be emptied.
// When `allannots` is true, all annotations will be flattened. Keep false if want to keep non-form related
// annotations intact.
// When `appgen` is not nil, it will be used to generate appearance streams for the field annotations.
func (_fcbbe *PdfReader) FlattenFields(allannots bool, appgen FieldAppearanceGenerator) error {
	return _fcbbe.flattenFieldsWithOpts(allannots, appgen, nil)
}

// NewDSS returns a new DSS dictionary.
func NewDSS() *DSS {
	return &DSS{_gbgda: _cde.MakeIndirectObject(_cde.MakeDict()), VRI: map[string]*VRI{}}
}
func (_gdfge *pdfFontSimple) updateStandard14Font() {
	_addc, _cafa := _gdfge.Encoder().(_gc.SimpleEncoder)
	if !_cafa {
		_ad.Log.Error("\u0057\u0072\u006f\u006e\u0067\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074y\u0070e\u003a\u0020\u0025\u0054\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u002e", _gdfge.Encoder(), _gdfge)
		return
	}
	_efbea := _addc.Charcodes()
	_gdfge._fecg = make(map[_gc.CharCode]float64, len(_efbea))
	for _, _eddc := range _efbea {
		_bgeb, _ := _addc.CharcodeToRune(_eddc)
		_fcabg, _ := _gdfge._gggc.Read(_bgeb)
		_gdfge._fecg[_eddc] = _fcabg.Wx
	}
}

// PdfAnnotationFileAttachment represents FileAttachment annotations.
// (Section 12.5.6.15).
type PdfAnnotationFileAttachment struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	FS   _cde.PdfObject
	Name _cde.PdfObject
}

func _bgcac(_dfge *PdfField) []*PdfField {
	_ebfg := []*PdfField{_dfge}
	for _, _aabe := range _dfge.Kids {
		_ebfg = append(_ebfg, _bgcac(_aabe)...)
	}
	return _ebfg
}
func (_gffggc *PdfWriter) writeOutlines() error {
	if _gffggc._bdfce == nil {
		return nil
	}
	_ad.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0054\u0072\u0065\u0065\u003a\u0020\u0025\u002b\u0076", _gffggc._bdfce)
	_ecgff := _gffggc._bdfce.ToPdfObject()
	_ad.Log.Trace("\u004fu\u0074\u006c\u0069\u006e\u0065\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0025\u0054\u002c\u0020\u0070\u003a\u0025\u0070\u0029", _ecgff, _ecgff, _ecgff)
	_gffggc._fedbb.Set("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073", _ecgff)
	_ddfdc := _gffggc.addObjects(_ecgff)
	if _ddfdc != nil {
		return _ddfdc
	}
	return nil
}

// ToPdfObject converts the PdfPage to a dictionary within an indirect object container.
func (_dcgfbd *PdfPage) ToPdfObject() _cde.PdfObject {
	_gefbb := _dcgfbd._dcaeff
	_dcgfbd.GetPageDict()
	return _gefbb
}

// AddExtGState add External Graphics State (GState). The gsDict can be specified
// either directly as a dictionary or an indirect object containing a dictionary.
func (_eebc *PdfPageResources) AddExtGState(gsName _cde.PdfObjectName, gsDict _cde.PdfObject) error {
	if _eebc.ExtGState == nil {
		_eebc.ExtGState = _cde.MakeDict()
	}
	_afged := _eebc.ExtGState
	_gcccd, _fgfd := _cde.TraceToDirectObject(_afged).(*_cde.PdfObjectDictionary)
	if !_fgfd {
		_ad.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020e\u0072r\u006f\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u002f\u0025\u0054\u0029", _afged, _cde.TraceToDirectObject(_afged))
		return _cde.ErrTypeError
	}
	_gcccd.Set(gsName, gsDict)
	return nil
}
func (_fdgcg *PdfWriter) optimizeDocument() error {
	if _fdgcg._cfaf == nil {
		return nil
	}
	_gefcff, _cedbb := _cde.GetDict(_fdgcg._fdgbc)
	if !_cedbb {
		return _ceg.New("\u0061\u006e\u0020in\u0066\u006f\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_aeadc := _ab.Document{ID: [2]string{_fdgcg._gfegf, _fdgcg._dcace}, Version: _fdgcg._cgdcc, Objects: _fdgcg._egbccc, Info: _gefcff, Crypt: _fdgcg._ccgbe, UseHashBasedID: _fdgcg._fcdff}
	if _faeda := _fdgcg._cfaf.ApplyStandard(&_aeadc); _faeda != nil {
		return _faeda
	}
	_fdgcg._gfegf, _fdgcg._dcace = _aeadc.ID[0], _aeadc.ID[1]
	_fdgcg._cgdcc = _aeadc.Version
	_fdgcg._egbccc = _aeadc.Objects
	_fdgcg._fdgbc.PdfObject = _aeadc.Info
	_fdgcg._fcdff = _aeadc.UseHashBasedID
	_fdgcg._ccgbe = _aeadc.Crypt
	_edcfa := make(map[_cde.PdfObject]struct{}, len(_fdgcg._egbccc))
	for _, _faegf := range _fdgcg._egbccc {
		_edcfa[_faegf] = struct{}{}
	}
	_fdgcg._bccde = _edcfa
	return nil
}

// PdfActionResetForm represents a resetForm action.
type PdfActionResetForm struct {
	*PdfAction
	Fields _cde.PdfObject
	Flags  _cde.PdfObject
}

// PdfShadingType5 is a Lattice-form Gouraud-shaded triangle mesh.
type PdfShadingType5 struct {
	*PdfShading
	BitsPerCoordinate *_cde.PdfObjectInteger
	BitsPerComponent  *_cde.PdfObjectInteger
	VerticesPerRow    *_cde.PdfObjectInteger
	Decode            *_cde.PdfObjectArray
	Function          []PdfFunction
}

func _cbbd(_acge _cde.PdfObject) (*PdfFunctionType2, error) {
	_dcaab := &PdfFunctionType2{}
	var _deab *_cde.PdfObjectDictionary
	if _aegb, _edddc := _acge.(*_cde.PdfIndirectObject); _edddc {
		_ffcf, _gfdae := _aegb.PdfObject.(*_cde.PdfObjectDictionary)
		if !_gfdae {
			return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_dcaab._bddc = _aegb
		_deab = _ffcf
	} else if _gbacc, _fbbcg := _acge.(*_cde.PdfObjectDictionary); _fbbcg {
		_deab = _gbacc
	} else {
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_ad.Log.Trace("\u0046U\u004e\u0043\u0032\u003a\u0020\u0025s", _deab.String())
	_gbccd, _becfe := _cde.TraceToDirectObject(_deab.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_cde.PdfObjectArray)
	if !_becfe {
		_ad.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _gbccd.Len() < 0 || _gbccd.Len()%2 != 0 {
		_ad.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u0072\u0061\u006e\u0067e\u0020\u0069\u006e\u0076al\u0069\u0064")
		return nil, _ceg.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_bcbdc, _babga := _gbccd.ToFloat64Array()
	if _babga != nil {
		return nil, _babga
	}
	_dcaab.Domain = _bcbdc
	_gbccd, _becfe = _cde.TraceToDirectObject(_deab.Get("\u0052\u0061\u006eg\u0065")).(*_cde.PdfObjectArray)
	if _becfe {
		if _gbccd.Len() < 0 || _gbccd.Len()%2 != 0 {
			return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_fdfac, _cdcdf := _gbccd.ToFloat64Array()
		if _cdcdf != nil {
			return nil, _cdcdf
		}
		_dcaab.Range = _fdfac
	}
	_gbccd, _becfe = _cde.TraceToDirectObject(_deab.Get("\u0043\u0030")).(*_cde.PdfObjectArray)
	if _becfe {
		_fdffd, _cccd := _gbccd.ToFloat64Array()
		if _cccd != nil {
			return nil, _cccd
		}
		_dcaab.C0 = _fdffd
	}
	_gbccd, _becfe = _cde.TraceToDirectObject(_deab.Get("\u0043\u0031")).(*_cde.PdfObjectArray)
	if _becfe {
		_adde, _ebbg := _gbccd.ToFloat64Array()
		if _ebbg != nil {
			return nil, _ebbg
		}
		_dcaab.C1 = _adde
	}
	if len(_dcaab.C0) != len(_dcaab.C1) {
		_ad.Log.Error("\u0043\u0030\u0020\u0061nd\u0020\u0043\u0031\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069n\u0067")
		return nil, _cde.ErrRangeError
	}
	N, _babga := _cde.GetNumberAsFloat(_cde.TraceToDirectObject(_deab.Get("\u004e")))
	if _babga != nil {
		_ad.Log.Error("\u004e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020o\u0072\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u002c\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _deab.String())
		return nil, _babga
	}
	_dcaab.N = N
	return _dcaab, nil
}

// PdfColorspaceDeviceRGB represents an RGB colorspace.
type PdfColorspaceDeviceRGB struct{}

// GetContainingPdfObject implements interface PdfModel.
func (_bcdfg *PdfSignature) GetContainingPdfObject() _cde.PdfObject { return _bcdfg._cabd }

// SetDocInfo sets the document /Info metadata.
// This will overwrite any globally declared document info.
func (_dabg *PdfAppender) SetDocInfo(info *PdfInfo) { _dabg._gdaa = info }

// NewPdfAppenderWithOpts creates a new Pdf appender from a Pdf reader with options.
func NewPdfAppenderWithOpts(reader *PdfReader, opts *ReaderOpts, encryptOptions *EncryptOptions) (*PdfAppender, error) {
	_aadg := &PdfAppender{_aac: reader._caecc, Reader: reader, _gdgf: reader._aggcgb, _egbc: reader._efbdd}
	_egbeb, _fabe := _aadg._aac.Seek(0, _f.SeekEnd)
	if _fabe != nil {
		return nil, _fabe
	}
	_aadg._efag = _egbeb
	if _, _fabe = _aadg._aac.Seek(0, _f.SeekStart); _fabe != nil {
		return nil, _fabe
	}
	_aadg._cfad, _fabe = NewPdfReaderWithOpts(_aadg._aac, opts)
	if _fabe != nil {
		return nil, _fabe
	}
	for _, _dgc := range _aadg.Reader.GetObjectNums() {
		if _aadg._afcea < _dgc {
			_aadg._afcea = _dgc
		}
	}
	_aadg._eccee = _aadg._gdgf.GetXrefTable()
	_aadg._aabb = _aadg._gdgf.GetXrefOffset()
	_aadg._gfb = append(_aadg._gfb, _aadg._cfad.PageList...)
	_aadg._aadd = make(map[_cde.PdfObject]struct{})
	_aadg._adfa = make(map[_cde.PdfObject]int64)
	_aadg._deb = make(map[_cde.PdfObject]struct{})
	_aadg._cbeaa = _aadg._cfad.AcroForm
	_aadg._ffac = _aadg._cfad.DSS
	if opts != nil {
		_aadg._dgdff = opts.Password
	}
	if encryptOptions != nil {
		_aadg._fcga = encryptOptions
	}
	return _aadg, nil
}

// Mask returns the uin32 bitmask for the specific flag.
func (_gedgg FieldFlag) Mask() uint32 { return uint32(_gedgg) }

// GetRotate gets the inheritable rotate value, either from the page
// or a higher up page/pages struct.
func (_gccebe *PdfPage) GetRotate() (int64, error) {
	if _gccebe.Rotate != nil {
		return *_gccebe.Rotate, nil
	}
	_degfd := _gccebe.Parent
	for _degfd != nil {
		_adafa, _dfcfg := _cde.GetDict(_degfd)
		if !_dfcfg {
			return 0, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _bcgd := _adafa.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _bcgd != nil {
			_ffcbd, _ebbfd := _cde.GetInt(_bcgd)
			if !_ebbfd {
				return 0, _ceg.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0074a\u0074\u0065\u0020\u0076al\u0075\u0065")
			}
			if _ffcbd != nil {
				return int64(*_ffcbd), nil
			}
			return 0, _ceg.New("\u0072\u006f\u0074\u0061te\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		_degfd = _adafa.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return 0, _ceg.New("\u0072o\u0074a\u0074\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
}

type pdfFontType3 struct {
	fontCommon
	_fddbc *_cde.PdfIndirectObject

	// These fields are specific to Type 3 fonts.
	CharProcs  _cde.PdfObject
	Encoding   _cde.PdfObject
	FontBBox   _cde.PdfObject
	FontMatrix _cde.PdfObject
	FirstChar  _cde.PdfObject
	LastChar   _cde.PdfObject
	Widths     _cde.PdfObject
	Resources  _cde.PdfObject
	_dedg      map[_gc.CharCode]float64
	_bdcege    _gc.TextEncoder
}

func _bbeaa(_ecgfg string) (string, error) {
	var _baafgbf _ede.Buffer
	_baafgbf.WriteString(_ecgfg)
	_ddace := make([]byte, 8+16)
	_fgga := _ce.Now().UTC().UnixNano()
	_ba.BigEndian.PutUint64(_ddace, uint64(_fgga))
	_, _eeeef := _d.Read(_ddace[8:])
	if _eeeef != nil {
		return "", _eeeef
	}
	_baafgbf.WriteString(_ed.EncodeToString(_ddace))
	return _baafgbf.String(), nil
}
func (_gag *PdfReader) newPdfActionSoundFromDict(_bfc *_cde.PdfObjectDictionary) (*PdfActionSound, error) {
	return &PdfActionSound{Sound: _bfc.Get("\u0053\u006f\u0075n\u0064"), Volume: _bfc.Get("\u0056\u006f\u006c\u0075\u006d\u0065"), Synchronous: _bfc.Get("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073"), Repeat: _bfc.Get("\u0052\u0065\u0070\u0065\u0061\u0074"), Mix: _bfc.Get("\u004d\u0069\u0078")}, nil
}

// C returns the value of the cyan component of the color.
func (_bbce *PdfColorDeviceCMYK) C() float64 { return _bbce[0] }

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_aebb *PdfColorspaceSpecialIndexed) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	N := _aebb.Base.GetNumComponents()
	_gbaga := int(vals[0]) * N
	if _gbaga < 0 || (_gbaga+N-1) >= len(_aebb._dafd) {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _gbaga)
		return nil, ErrColorOutOfRange
	}
	_afgdc := _aebb._dafd[_gbaga : _gbaga+N]
	var _bcfb []float64
	for _, _eage := range _afgdc {
		_bcfb = append(_bcfb, float64(_eage)/255.0)
	}
	_afcf, _dfbbc := _aebb.Base.ColorFromFloats(_bcfb)
	if _dfbbc != nil {
		return nil, _dfbbc
	}
	return _afcf, nil
}

// GetAsTilingPattern returns a tiling pattern. Check with IsTiling() prior to using this.
func (_cfdge *PdfPattern) GetAsTilingPattern() *PdfTilingPattern {
	return _cfdge._abddb.(*PdfTilingPattern)
}
func _dgefe(_ebde _cde.PdfObject, _eaeb bool) (*PdfFont, error) {
	_ggee, _gadd, _fbecf := _cgebb(_ebde)
	if _ggee != nil {
		_caag(_ggee)
	}
	if _fbecf != nil {
		if _fbecf == ErrType1CFontNotSupported {
			_dfebe, _eebgg := _eaddb(_ggee, _gadd, nil)
			if _eebgg != nil {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _gadd, _eebgg)
				return nil, _fbecf
			}
			return &PdfFont{_gbcff: _dfebe}, _fbecf
		}
		return nil, _fbecf
	}
	_gcec := &PdfFont{}
	switch _gadd._dcbc {
	case "\u0054\u0079\u0070e\u0030":
		if !_eaeb {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u004c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u00650\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gadd)
			return nil, _ceg.New("\u0063\u0079\u0063\u006cic\u0061\u006c\u0020\u0074\u0079\u0070\u0065\u0030\u0020\u006c\u006f\u0061\u0064\u0069n\u0067")
		}
		_gfff, _fbfb := _cbbf(_ggee, _gadd)
		if _fbfb != nil {
			_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0057\u0068\u0069l\u0065\u0020\u006c\u006f\u0061\u0064\u0069ng\u0020\u0054\u0079\u0070e\u0030\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066on\u0074\u003d%\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gadd, _fbfb)
			return nil, _fbfb
		}
		_gcec._gbcff = _gfff
	case "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		var _ecfed *pdfFontSimple
		_faaa, _acag := _fe.NewStdFontByName(_fe.StdFontName(_gadd._eeab))
		if _acag {
			_aaef := _gdaeg(_faaa)
			_gcec._gbcff = &_aaef
			_eddd := _cde.TraceToDirectObject(_aaef.ToPdfObject())
			_efdac, _fbdb, _gdafg := _cgebb(_eddd)
			if _gdafg != nil {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _gadd, _aaef)
				return nil, _gdafg
			}
			for _, _bafbe := range _ggee.Keys() {
				_efdac.Set(_bafbe, _ggee.Get(_bafbe))
			}
			_ecfed, _gdafg = _eaddb(_efdac, _fbdb, _aaef._facga)
			if _gdafg != nil {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _gadd, _aaef)
				return nil, _gdafg
			}
			_ecfed._fecg = _aaef._fecg
			_ecfed._gggc = _aaef._gggc
			if _ecfed._ccff == nil {
				_ecfed._ccff = _aaef._ccff
			}
		} else {
			_ecfed, _fbecf = _eaddb(_ggee, _gadd, nil)
			if _fbecf != nil {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _gadd, _fbecf)
				return nil, _fbecf
			}
		}
		_fbecf = _ecfed.addEncoding()
		if _fbecf != nil {
			return nil, _fbecf
		}
		if _acag {
			_ecfed.updateStandard14Font()
		}
		if _acag && _ecfed._efeaf == nil && _ecfed._facga == nil {
			_ad.Log.Error("\u0073\u0069\u006d\u0070\u006c\u0065\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _ecfed)
			_ad.Log.Error("\u0066n\u0074\u003d\u0025\u002b\u0076", _faaa)
		}
		if len(_ecfed._fecg) == 0 {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u006f\u0020\u0077\u0069d\u0074h\u0073.\u0020\u0066\u006f\u006e\u0074\u003d\u0025s", _ecfed)
		}
		_gcec._gbcff = _ecfed
	case "\u0054\u0079\u0070e\u0033":
		_degg, _babg := _decfc(_ggee, _gadd)
		if _babg != nil {
			_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020W\u0068\u0069\u006c\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0074y\u0070\u0065\u0033\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _babg)
			return nil, _babg
		}
		_gcec._gbcff = _degg
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030":
		_gbafe, _dfdd := _fegg(_ggee, _gadd)
		if _dfdd != nil {
			_ad.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0057\u0068i\u006c\u0065\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069d \u0066\u006f\u006et\u0020\u0074y\u0070\u0065\u0030\u0020\u0066\u006fn\u0074\u003a \u0025\u0076", _dfdd)
			return nil, _dfdd
		}
		_gcec._gbcff = _gbafe
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_fggg, _gcfad := _gdff(_ggee, _gadd)
		if _gcfad != nil {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0057\u0068\u0069l\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069\u0064\u0020f\u006f\u006e\u0074\u0020\u0074yp\u0065\u0032\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gadd, _gcfad)
			return nil, _gcfad
		}
		_gcec._gbcff = _fggg
	default:
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020U\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020f\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073", _gadd)
		return nil, _ee.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0066\u006f\u006e\u0074\u0020\u0074y\u0070\u0065\u003a\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gadd)
	}
	return _gcec, nil
}

// ToPdfObject converts PdfAcroForm to a PdfObject, i.e. an indirect object containing the
// AcroForm dictionary.
func (_cgaf *PdfAcroForm) ToPdfObject() _cde.PdfObject {
	_bgfde := _cgaf._ecca
	_efcad := _bgfde.PdfObject.(*_cde.PdfObjectDictionary)
	if _cgaf.Fields != nil {
		_ccfe := _cde.PdfObjectArray{}
		for _, _adecd := range *_cgaf.Fields {
			_dggg := _adecd.GetContext()
			if _dggg != nil {
				_ccfe.Append(_dggg.ToPdfObject())
			} else {
				_ccfe.Append(_adecd.ToPdfObject())
			}
		}
		_efcad.Set("\u0046\u0069\u0065\u006c\u0064\u0073", &_ccfe)
	}
	if _cgaf.NeedAppearances != nil {
		_efcad.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cgaf.NeedAppearances)
	}
	if _cgaf.SigFlags != nil {
		_efcad.Set("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073", _cgaf.SigFlags)
	}
	if _cgaf.CO != nil {
		_efcad.Set("\u0043\u004f", _cgaf.CO)
	}
	if _cgaf.DR != nil {
		_efcad.Set("\u0044\u0052", _cgaf.DR.ToPdfObject())
	}
	if _cgaf.DA != nil {
		_efcad.Set("\u0044\u0041", _cgaf.DA)
	}
	if _cgaf.Q != nil {
		_efcad.Set("\u0051", _cgaf.Q)
	}
	if _cgaf.XFA != nil {
		_efcad.Set("\u0058\u0046\u0041", _cgaf.XFA)
	}
	return _bgfde
}

// GetNumComponents returns the number of color components (3 for Lab).
func (_fbce *PdfColorLab) GetNumComponents() int { return 3 }
func (_cgad *PdfColorspaceDeviceRGB) String() string {
	return "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
}

var _caggc = _cc.MustCompile("\u005b\\\u006e\u005c\u0072\u005d\u002b")

// WriteToFile writes the Appender output to file specified by path.
func (_dfee *PdfAppender) WriteToFile(outputPath string) error {
	_edde, _egca := _db.Create(outputPath)
	if _egca != nil {
		return _egca
	}
	defer _edde.Close()
	return _dfee.Write(_edde)
}

type pdfFontType0 struct {
	fontCommon
	_dcda          *_cde.PdfIndirectObject
	_cffef         _gc.TextEncoder
	Encoding       _cde.PdfObject
	DescendantFont *PdfFont
	_dcdga         *_fb.CMap
}

func (_adcg *DSS) add(_defb *[]*_cde.PdfObjectStream, _dggfc map[string]*_cde.PdfObjectStream, _cgbf [][]byte) ([]*_cde.PdfObjectStream, error) {
	_gbcfc := make([]*_cde.PdfObjectStream, 0, len(_cgbf))
	for _, _fdee := range _cgbf {
		_gacc, _cfec := _dgaff(_fdee)
		if _cfec != nil {
			return nil, _cfec
		}
		_cggb, _dagf := _dggfc[string(_gacc)]
		if !_dagf {
			_cggb, _cfec = _cde.MakeStream(_fdee, _cde.NewRawEncoder())
			if _cfec != nil {
				return nil, _cfec
			}
			_dggfc[string(_gacc)] = _cggb
			*_defb = append(*_defb, _cggb)
		}
		_gbcfc = append(_gbcfc, _cggb)
	}
	return _gbcfc, nil
}

// PdfFunctionType2 defines an exponential interpolation of one input value and n
// output values:
//      f(x) = y_0, ..., y_(n-1)
// y_j = C0_j + x^N * (C1_j - C0_j); for 0 <= j < n
// When N=1 ; linear interpolation between C0 and C1.
type PdfFunctionType2 struct {
	Domain []float64
	Range  []float64
	C0     []float64
	C1     []float64
	N      float64
	_bddc  *_cde.PdfIndirectObject
}

func (_edefa *PdfPage) setContainer(_fgacf *_cde.PdfIndirectObject) {
	_fgacf.PdfObject = _edefa._gbbc
	_edefa._dcaeff = _fgacf
}

// NewPdfActionTrans returns a new "trans" action.
func NewPdfActionTrans() *PdfActionTrans {
	_fbe := NewPdfAction()
	_ddd := &PdfActionTrans{}
	_ddd.PdfAction = _fbe
	_fbe.SetContext(_ddd)
	return _ddd
}

// SetContext sets the sub pattern (context).  Either PdfTilingPattern or PdfShadingPattern.
func (_fbecfd *PdfPattern) SetContext(ctx PdfModel) { _fbecfd._abddb = ctx }
func (_ffbec *PdfReader) loadForms() (*PdfAcroForm, error) {
	if _ffbec._aggcgb.GetCrypter() != nil && !_ffbec._aggcgb.IsAuthenticated() {
		return nil, _ee.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_gcfdc := _ffbec._efabe
	_feggd := _gcfdc.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")
	if _feggd == nil {
		return nil, nil
	}
	_eggfb, _ := _cde.GetIndirect(_feggd)
	_feggd = _cde.TraceToDirectObject(_feggd)
	if _cde.IsNullObject(_feggd) {
		_ad.Log.Trace("\u0041\u0063\u0072of\u006f\u0072\u006d\u0020\u0069\u0073\u0020\u0061\u0020n\u0075l\u006c \u006fb\u006a\u0065\u0063\u0074\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029\u000a")
		return nil, nil
	}
	_gabbba, _adddb := _cde.GetDict(_feggd)
	if !_adddb {
		_ad.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0041\u0063\u0072\u006fF\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _feggd)
		_ad.Log.Debug("\u0044\u006f\u0065\u0073 n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0073")
		return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0061\u0063\u0072\u006ff\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _feggd)
	}
	_ad.Log.Trace("\u0048\u0061\u0073\u0020\u0041\u0063\u0072\u006f\u0020f\u006f\u0072\u006d\u0073")
	_ad.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0041\u0063r\u006ff\u006f\u0072\u006d\u0073\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065")
	if !_ffbec._cdgee {
		_dggbdb := _ffbec.traverseObjectData(_gabbba)
		if _dggbdb != nil {
			_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0074\u0072a\u0076\u0065\u0072\u0073\u0065\u0020\u0041\u0063\u0072\u006fFo\u0072\u006d\u0073 \u0028%\u0073\u0029", _dggbdb)
			return nil, _dggbdb
		}
	}
	_aeca, _cdfee := _ffbec.newPdfAcroFormFromDict(_eggfb, _gabbba)
	if _cdfee != nil {
		return nil, _cdfee
	}
	return _aeca, nil
}

// PdfColorspaceSpecialPattern is a Pattern colorspace.
// Can be defined either as /Pattern or with an underlying colorspace [/Pattern cs].
type PdfColorspaceSpecialPattern struct {
	UnderlyingCS PdfColorspace
	_ffcb        *_cde.PdfIndirectObject
}

func _bddge(_aeeea []byte) []byte {
	const _ccebf = 52845
	const _eafef = 22719
	_eafc := 55665
	for _, _dadfd := range _aeeea[:4] {
		_eafc = (int(_dadfd)+_eafc)*_ccebf + _eafef
	}
	_gfegef := make([]byte, len(_aeeea)-4)
	for _fabb, _ebcgd := range _aeeea[4:] {
		_gfegef[_fabb] = byte(int(_ebcgd) ^ _eafc>>8)
		_eafc = (int(_ebcgd)+_eafc)*_ccebf + _eafef
	}
	return _gfegef
}
func (_bbafb *pdfFontSimple) baseFields() *fontCommon { return &_bbafb.fontCommon }
func (_cdff *PdfWriter) setDocInfo(_cccecg _cde.PdfObject) {
	if _cdff.hasObject(_cdff._fdgbc) {
		delete(_cdff._bccde, _cdff._fdgbc)
		delete(_cdff._gbdfb, _cdff._fdgbc)
		for _bbddc, _deeeaa := range _cdff._egbccc {
			if _deeeaa == _cdff._fdgbc {
				copy(_cdff._egbccc[_bbddc:], _cdff._egbccc[_bbddc+1:])
				_cdff._egbccc[len(_cdff._egbccc)-1] = nil
				_cdff._egbccc = _cdff._egbccc[:len(_cdff._egbccc)-1]
				break
			}
		}
	}
	_degb := _cde.PdfIndirectObject{}
	_degb.PdfObject = _cccecg
	_cdff._fdgbc = &_degb
	_cdff.addObject(&_degb)
}
func (_fbde *PdfReader) newPdfFieldSignatureFromDict(_ffgae *_cde.PdfObjectDictionary) (*PdfFieldSignature, error) {
	_agaf := &PdfFieldSignature{}
	_fdca, _efbc := _cde.GetIndirect(_ffgae.Get("\u0056"))
	if _efbc {
		var _bebac error
		_agaf.V, _bebac = _fbde.newPdfSignatureFromIndirect(_fdca)
		if _bebac != nil {
			return nil, _bebac
		}
	}
	_agaf.Lock, _ = _cde.GetIndirect(_ffgae.Get("\u004c\u006f\u0063\u006b"))
	_agaf.SV, _ = _cde.GetIndirect(_ffgae.Get("\u0053\u0056"))
	return _agaf, nil
}

// PdfAnnotationStamp represents Stamp annotations.
// (Section 12.5.6.12).
type PdfAnnotationStamp struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Name _cde.PdfObject
}

// PdfAnnotationLink represents Link annotations.
// (Section 12.5.6.5 p. 403).
type PdfAnnotationLink struct {
	*PdfAnnotation
	A          _cde.PdfObject
	Dest       _cde.PdfObject
	H          _cde.PdfObject
	PA         _cde.PdfObject
	QuadPoints _cde.PdfObject
	BS         _cde.PdfObject
	_beg       *PdfAction
	_gegc      *PdfReader
}

// NewPdfAppender creates a new Pdf appender from a Pdf reader.
func NewPdfAppender(reader *PdfReader) (*PdfAppender, error) {
	_gbacg := &PdfAppender{_aac: reader._caecc, Reader: reader, _gdgf: reader._aggcgb, _egbc: reader._efbdd}
	_cdf, _agbf := _gbacg._aac.Seek(0, _f.SeekEnd)
	if _agbf != nil {
		return nil, _agbf
	}
	_gbacg._efag = _cdf
	if _, _agbf = _gbacg._aac.Seek(0, _f.SeekStart); _agbf != nil {
		return nil, _agbf
	}
	_gbacg._cfad, _agbf = NewPdfReader(_gbacg._aac)
	if _agbf != nil {
		return nil, _agbf
	}
	for _, _fdeg := range _gbacg.Reader.GetObjectNums() {
		if _gbacg._afcea < _fdeg {
			_gbacg._afcea = _fdeg
		}
	}
	_gbacg._eccee = _gbacg._gdgf.GetXrefTable()
	_gbacg._aabb = _gbacg._gdgf.GetXrefOffset()
	_gbacg._gfb = append(_gbacg._gfb, _gbacg._cfad.PageList...)
	_gbacg._aadd = make(map[_cde.PdfObject]struct{})
	_gbacg._adfa = make(map[_cde.PdfObject]int64)
	_gbacg._deb = make(map[_cde.PdfObject]struct{})
	_gbacg._cbeaa = _gbacg._cfad.AcroForm
	_gbacg._ffac = _gbacg._cfad.DSS
	return _gbacg, nil
}

// GetContainingPdfObject returns the container of the outline item (indirect object).
func (_abag *PdfOutlineItem) GetContainingPdfObject() _cde.PdfObject { return _abag._ccbf }

// UpdateObject marks `obj` as updated and to be included in the following revision.
func (_gdeb *PdfAppender) UpdateObject(obj _cde.PdfObject) {
	_gdeb.replaceObject(obj, obj)
	if _, _bbab := _gdeb._aadd[obj]; !_bbab {
		_gdeb._bfec = append(_gdeb._bfec, obj)
		_gdeb._aadd[obj] = struct{}{}
	}
}

// GetRuneMetrics returns the character metrics for the rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_fcfff pdfFontSimple) GetRuneMetrics(r rune) (_fe.CharMetrics, bool) {
	if _fcfff._gggc != nil {
		_cafd, _ddfbb := _fcfff._gggc.Read(r)
		if _ddfbb {
			return _cafd, true
		}
	}
	_fafef := _fcfff.Encoder()
	if _fafef == nil {
		_ad.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _fcfff)
		return _fe.CharMetrics{}, false
	}
	_afdee, _dfcdc := _fafef.RuneToCharcode(r)
	if !_dfcdc {
		if r != ' ' {
			_ad.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _fcfff)
		}
		return _fe.CharMetrics{}, false
	}
	_gfddf, _beae := _fcfff.GetCharMetrics(_afdee)
	return _gfddf, _beae
}
func (_dedgc *PdfPattern) getDict() *_cde.PdfObjectDictionary {
	if _dbdf, _gfcaf := _dedgc._eecac.(*_cde.PdfIndirectObject); _gfcaf {
		_cadff, _gbecf := _dbdf.PdfObject.(*_cde.PdfObjectDictionary)
		if !_gbecf {
			return nil
		}
		return _cadff
	} else if _aagc, _ccba := _dedgc._eecac.(*_cde.PdfObjectStream); _ccba {
		return _aagc.PdfObjectDictionary
	} else {
		_ad.Log.Debug("\u0054r\u0079\u0069\u006e\u0067\u0020\u0074\u006f a\u0063\u0063\u0065\u0073\u0073\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063t \u0074\u0079\u0070e\u0020\u0028\u0025\u0054\u0029", _dedgc._eecac)
		return nil
	}
}
func (_defde *PdfColorspaceSpecialIndexed) String() string {
	return "\u0049n\u0064\u0065\u0078\u0065\u0064"
}

// ToPdfObject implements interface PdfModel.
func (_abee *PdfAnnotationWatermark) ToPdfObject() _cde.PdfObject {
	_abee.PdfAnnotation.ToPdfObject()
	_ffce := _abee._bddg
	_beefb := _ffce.PdfObject.(*_cde.PdfObjectDictionary)
	_beefb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k"))
	_beefb.SetIfNotNil("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074", _abee.FixedPrint)
	return _ffce
}

// Image interface is a basic representation of an image used in PDF.
// The colorspace is not specified, but must be known when handling the image.
type Image struct {
	Width            int64
	Height           int64
	BitsPerComponent int64
	ColorComponents  int
	Data             []byte
	_deegf           []byte
	_aaafb           []float64
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_ebfdc *PdfShadingType6) ToPdfObject() _cde.PdfObject {
	_ebfdc.PdfShading.ToPdfObject()
	_ddff, _aedbf := _ebfdc.getShadingDict()
	if _aedbf != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _ebfdc.BitsPerCoordinate != nil {
		_ddff.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _ebfdc.BitsPerCoordinate)
	}
	if _ebfdc.BitsPerComponent != nil {
		_ddff.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _ebfdc.BitsPerComponent)
	}
	if _ebfdc.BitsPerFlag != nil {
		_ddff.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _ebfdc.BitsPerFlag)
	}
	if _ebfdc.Decode != nil {
		_ddff.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _ebfdc.Decode)
	}
	if _ebfdc.Function != nil {
		if len(_ebfdc.Function) == 1 {
			_ddff.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _ebfdc.Function[0].ToPdfObject())
		} else {
			_ccfade := _cde.MakeArray()
			for _, _dcece := range _ebfdc.Function {
				_ccfade.Append(_dcece.ToPdfObject())
			}
			_ddff.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _ccfade)
		}
	}
	return _ebfdc._dffg
}

// GetContainingPdfObject implements interface PdfModel.
func (_eggfcc *Permissions) GetContainingPdfObject() _cde.PdfObject { return _eggfcc._dafaf }

// GetSubFilter returns SubFilter value or empty string.
func (_aeeca *pdfSignDictionary) GetSubFilter() string {
	_bcdaeb := _aeeca.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r")
	if _bcdaeb == nil {
		return ""
	}
	if _aaec, _eebff := _cde.GetNameVal(_bcdaeb); _eebff {
		return _aaec
	}
	return ""
}

// ToPdfObject returns a PDF object representation of the outline.
func (_faegda *Outline) ToPdfObject() _cde.PdfObject { return _faegda.ToPdfOutline().ToPdfObject() }

// AcroFormNeedsRepair returns true if the document contains widget annotations
// linked to fields which are not referenced in the AcroForm. The AcroForm can
// be repaired using the RepairAcroForm method of the reader.
func (_bbdbd *PdfReader) AcroFormNeedsRepair() (bool, error) {
	var _dfeafc []*PdfField
	if _bbdbd.AcroForm != nil {
		_dfeafc = _bbdbd.AcroForm.AllFields()
	}
	_aaedc := make(map[*PdfField]struct{}, len(_dfeafc))
	for _, _efdge := range _dfeafc {
		_aaedc[_efdge] = struct{}{}
	}
	for _, _egade := range _bbdbd.PageList {
		_ddbcc, _ebega := _egade.GetAnnotations()
		if _ebega != nil {
			return false, _ebega
		}
		for _, _abfcd := range _ddbcc {
			_dfbda, _bcgce := _abfcd.GetContext().(*PdfAnnotationWidget)
			if !_bcgce {
				continue
			}
			_eaadc := _dfbda.Field()
			if _eaadc == nil {
				return true, nil
			}
			if _, _adbdg := _aaedc[_eaadc]; !_adbdg {
				return true, nil
			}
		}
	}
	return false, nil
}
func (_edabb *pdfFontType0) baseFields() *fontCommon { return &_edabb.fontCommon }

// ColorToRGB converts a CMYK32 color to an RGB color.
func (_bgge *PdfColorspaceDeviceCMYK) ColorToRGB(color PdfColor) (PdfColor, error) {
	_becc, _eadd := color.(*PdfColorDeviceCMYK)
	if !_eadd {
		_ad.Log.Debug("I\u006e\u0070\u0075\u0074\u0020\u0063o\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0064e\u0076\u0069\u0063e\u0020c\u006d\u0079\u006b")
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_ddcce := _becc.C()
	_aacb := _becc.M()
	_dgec := _becc.Y()
	_agea := _becc.K()
	_ddcce = _ddcce*(1-_agea) + _agea
	_aacb = _aacb*(1-_agea) + _agea
	_dgec = _dgec*(1-_agea) + _agea
	_gbcf := 1 - _ddcce
	_ggea := 1 - _aacb
	_gagaa := 1 - _dgec
	return NewPdfColorDeviceRGB(_gbcf, _ggea, _gagaa), nil
}
func _gdff(_fcbbc *_cde.PdfObjectDictionary, _bfgae *fontCommon) (*pdfCIDFontType2, error) {
	if _bfgae._dcbc != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_ad.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0032\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _bfgae)
		return nil, _cde.ErrRangeError
	}
	_cbadd := _gdfg(_bfgae)
	_aafdd, _dbbc := _cde.GetDict(_fcbbc.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_dbbc {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _bfgae)
		return nil, ErrRequiredAttributeMissing
	}
	_cbadd.CIDSystemInfo = _aafdd
	_cbadd.DW = _fcbbc.Get("\u0044\u0057")
	_cbadd.W = _fcbbc.Get("\u0057")
	_cbadd.DW2 = _fcbbc.Get("\u0044\u0057\u0032")
	_cbadd.W2 = _fcbbc.Get("\u0057\u0032")
	_cbadd.CIDToGIDMap = _fcbbc.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070")
	_cbadd._eedac = 1000.0
	if _gcab, _gefgb := _cde.GetNumberAsFloat(_cbadd.DW); _gefgb == nil {
		_cbadd._eedac = _gcab
	}
	_acaba, _bdcdc := _fgecb(_cbadd.W)
	if _bdcdc != nil {
		return nil, _bdcdc
	}
	if _acaba == nil {
		_acaba = map[_gc.CharCode]float64{}
	}
	_cbadd._fdgc = _acaba
	return _cbadd, nil
}
func (_eeggf *PdfWriter) copyObjects() {
	_agfe := make(map[_cde.PdfObject]_cde.PdfObject)
	_eacfd := make([]_cde.PdfObject, 0, len(_eeggf._egbccc))
	_bbffg := make(map[_cde.PdfObject]struct{}, len(_eeggf._egbccc))
	_aaacc := make(map[_cde.PdfObject]struct{})
	for _, _acccd := range _eeggf._egbccc {
		_gbgg := _eeggf.copyObject(_acccd, _agfe, _aaacc, false)
		if _, _eadead := _aaacc[_acccd]; _eadead {
			continue
		}
		_eacfd = append(_eacfd, _gbgg)
		_bbffg[_gbgg] = struct{}{}
	}
	_eeggf._egbccc = _eacfd
	_eeggf._bccde = _bbffg
	_eeggf._fdgbc = _eeggf.copyObject(_eeggf._fdgbc, _agfe, nil, false).(*_cde.PdfIndirectObject)
	_eeggf._eacge = _eeggf.copyObject(_eeggf._eacge, _agfe, nil, false).(*_cde.PdfIndirectObject)
	if _eeggf._bdbdb != nil {
		_eeggf._bdbdb = _eeggf.copyObject(_eeggf._bdbdb, _agfe, nil, false).(*_cde.PdfIndirectObject)
	}
	if _eeggf._aabfe {
		_bceef := make(map[_cde.PdfObject]int64)
		for _bgacgg, _aeebga := range _eeggf._dgad {
			if _ddbfa, _abacg := _agfe[_bgacgg]; _abacg {
				_bceef[_ddbfa] = _aeebga
			} else {
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020a\u0070\u0070\u0065n\u0064\u0020\u006d\u006fd\u0065\u0020\u002d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u0070\u0079\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070")
			}
		}
		_eeggf._dgad = _bceef
	}
}

// Transform rectangle with the supplied matrix.
func (_bcagdd *PdfRectangle) Transform(transformMatrix _adb.Matrix) {
	_bcagdd.Llx, _bcagdd.Lly = transformMatrix.Transform(_bcagdd.Llx, _bcagdd.Lly)
	_bcagdd.Urx, _bcagdd.Ury = transformMatrix.Transform(_bcagdd.Urx, _bcagdd.Ury)
	_bcagdd.Normalize()
}

// GetBorderWidth returns the border style's width.
func (_bed *PdfBorderStyle) GetBorderWidth() float64 {
	if _bed.W == nil {
		return 1
	}
	return *_bed.W
}

// PdfPage represents a page in a PDF document. (7.7.3.3 - Table 30).
type PdfPage struct {
	Parent               _cde.PdfObject
	LastModified         *PdfDate
	Resources            *PdfPageResources
	CropBox              *PdfRectangle
	MediaBox             *PdfRectangle
	BleedBox             *PdfRectangle
	TrimBox              *PdfRectangle
	ArtBox               *PdfRectangle
	BoxColorInfo         _cde.PdfObject
	Contents             _cde.PdfObject
	Rotate               *int64
	Group                _cde.PdfObject
	Thumb                _cde.PdfObject
	B                    _cde.PdfObject
	Dur                  _cde.PdfObject
	Trans                _cde.PdfObject
	AA                   _cde.PdfObject
	Metadata             _cde.PdfObject
	PieceInfo            _cde.PdfObject
	StructParents        _cde.PdfObject
	ID                   _cde.PdfObject
	PZ                   _cde.PdfObject
	SeparationInfo       _cde.PdfObject
	Tabs                 _cde.PdfObject
	TemplateInstantiated _cde.PdfObject
	PresSteps            _cde.PdfObject
	UserUnit             _cde.PdfObject
	VP                   _cde.PdfObject
	Annots               _cde.PdfObject
	_cefe                []*PdfAnnotation
	_gbbc                *_cde.PdfObjectDictionary
	_dcaeff              *_cde.PdfIndirectObject
	_gfcbe               *PdfReader
}

func _bagfg(_gdccc *_cde.PdfObjectArray) (float64, error) {
	_abefe, _ddccb := _gdccc.ToFloat64Array()
	if _ddccb != nil {
		_ad.Log.Debug("\u0042\u0061\u0064\u0020\u004d\u0061\u0074\u0074\u0065\u0020\u0061\u0072\u0072\u0061\u0079:\u0020m\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gdccc, _ddccb)
	}
	switch len(_abefe) {
	case 1:
		return _abefe[0], nil
	case 3:
		_aecfe := PdfColorspaceDeviceRGB{}
		_dabda, _bedg := _aecfe.ColorFromFloats(_abefe)
		if _bedg != nil {
			return 0.0, _bedg
		}
		return _dabda.(*PdfColorDeviceRGB).ToGray().Val(), nil
	case 4:
		_cbbae := PdfColorspaceDeviceCMYK{}
		_cgged, _eeeegb := _cbbae.ColorFromFloats(_abefe)
		if _eeeegb != nil {
			return 0.0, _eeeegb
		}
		_fdbegb, _eeeegb := _cbbae.ColorToRGB(_cgged.(*PdfColorDeviceCMYK))
		if _eeeegb != nil {
			return 0.0, _eeeegb
		}
		return _fdbegb.(*PdfColorDeviceRGB).ToGray().Val(), nil
	}
	_ddccb = _ceg.New("\u0062a\u0064 \u004d\u0061\u0074\u0074\u0065\u0020\u0063\u006f\u006c\u006f\u0072")
	_ad.Log.Error("\u0074\u006f\u0047ra\u0079\u003a\u0020\u006d\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gdccc, _ddccb)
	return 0.0, _ddccb
}

// ColorFromPdfObjects loads the color from PDF objects.
// The first objects (if present) represent the color in underlying colorspace.  The last one represents
// the name of the pattern.
func (_ddcbga *PdfColorspaceSpecialPattern) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) < 1 {
		return nil, _ceg.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_ccce := &PdfColorPattern{}
	_fabf, _bbaaa := objects[len(objects)-1].(*_cde.PdfObjectName)
	if !_bbaaa {
		_ad.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020a\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", objects[len(objects)-1])
		return nil, ErrTypeCheck
	}
	_ccce.PatternName = *_fabf
	if len(objects) > 1 {
		_edce := objects[0 : len(objects)-1]
		if _ddcbga.UnderlyingCS == nil {
			_ad.Log.Debug("P\u0061\u0074t\u0065\u0072\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0077\u0069\u0074\u0068\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u0020\u0062\u0075\u0074\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0063\u0073\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
			return nil, _ceg.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
		}
		_edab, _eadde := _ddcbga.UnderlyingCS.ColorFromPdfObjects(_edce)
		if _eadde != nil {
			_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0076\u0069\u0061\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0063\u0073\u003a\u0020\u0025\u0076", _eadde)
			return nil, _eadde
		}
		_ccce.Color = _edab
	}
	return _ccce, nil
}
func _gbad(_geadc *_cde.PdfObjectDictionary) (*PdfFieldText, error) {
	_eged := &PdfFieldText{}
	_eged.DA, _ = _cde.GetString(_geadc.Get("\u0044\u0041"))
	_eged.Q, _ = _cde.GetInt(_geadc.Get("\u0051"))
	_eged.DS, _ = _cde.GetString(_geadc.Get("\u0044\u0053"))
	_eged.RV = _geadc.Get("\u0052\u0056")
	_eged.MaxLen, _ = _cde.GetInt(_geadc.Get("\u004d\u0061\u0078\u004c\u0065\u006e"))
	return _eged, nil
}

// ToPdfObject implements interface PdfModel.
func (_aaeb *PdfAnnotationFileAttachment) ToPdfObject() _cde.PdfObject {
	_aaeb.PdfAnnotation.ToPdfObject()
	_caedd := _aaeb._bddg
	_dddf := _caedd.PdfObject.(*_cde.PdfObjectDictionary)
	_aaeb.PdfAnnotationMarkup.appendToPdfDictionary(_dddf)
	_dddf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074"))
	_dddf.SetIfNotNil("\u0046\u0053", _aaeb.FS)
	_dddf.SetIfNotNil("\u004e\u0061\u006d\u0065", _aaeb.Name)
	return _caedd
}
func _dggba(_cafccc []byte) (_ecbg, _agde string, _egfbg error) {
	_ad.Log.Trace("g\u0065\u0074\u0041\u0053CI\u0049S\u0065\u0063\u0074\u0069\u006fn\u0073\u003a\u0020\u0025\u0064\u0020", len(_cafccc))
	_dfbbg := _dfcde.FindIndex(_cafccc)
	if _dfbbg == nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0067\u0065\u0074\u0041\u0053\u0043\u0049\u0049\u0053\u0065\u0063\u0074\u0069o\u006e\u0073\u002e\u0020\u004e\u006f\u0020d\u0069\u0063\u0074\u002e")
		return "", "", _cde.ErrTypeError
	}
	_bgdge := _dfbbg[1]
	_cdabb := _dac.Index(string(_cafccc[_bgdge:]), _dcgd)
	if _cdabb < 0 {
		_ecbg = string(_cafccc[_bgdge:])
		return _ecbg, "", nil
	}
	_bbead := _bgdge + _cdabb
	_ecbg = string(_cafccc[_bgdge:_bbead])
	_bgedc := _bbead
	_cdabb = _dac.Index(string(_cafccc[_bgedc:]), _fafec)
	if _cdabb < 0 {
		_ad.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0067e\u0074\u0041\u0053\u0043\u0049\u0049\u0053e\u0063\u0074\u0069\u006f\u006e\u0073\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _egfbg)
		return "", "", _cde.ErrTypeError
	}
	_cfdg := _bgedc + _cdabb
	_agde = string(_cafccc[_bgedc:_cfdg])
	return _ecbg, _agde, nil
}
func _ggad(_daege _cde.PdfObject) (*PdfFontDescriptor, error) {
	_bafbee := &PdfFontDescriptor{}
	_daege = _cde.ResolveReference(_daege)
	if _cgcdb, _caggd := _daege.(*_cde.PdfIndirectObject); _caggd {
		_bafbee._fbff = _cgcdb
		_daege = _cgcdb.PdfObject
	}
	_gacb, _egbg := _cde.GetDict(_daege)
	if !_egbg {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0044\u0065\u0073c\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _daege)
		return nil, _cde.ErrTypeError
	}
	if _bcecd := _gacb.Get("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"); _bcecd != nil {
		_bafbee.FontName = _bcecd
	} else {
		_ad.Log.Debug("\u0049n\u0063\u006fm\u0070\u0061\u0074\u0069b\u0069\u006c\u0069t\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u004e\u0061me\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006di\u0073\u0073i\u006e\u0067")
	}
	_eegggc, _ := _cde.GetName(_bafbee.FontName)
	if _eead := _gacb.Get("\u0054\u0079\u0070\u0065"); _eead != nil {
		_ecbag, _gegde := _eead.(*_cde.PdfObjectName)
		if !_gegde || string(*_ecbag) != "\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072" {
			_ad.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072i\u0070t\u006f\u0072\u0020\u0054y\u0070\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0054\u0029\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _eead, _eegggc, _bafbee.FontName)
		}
	} else {
		_ad.Log.Trace("\u0049\u006ec\u006f\u006d\u0070\u0061\u0074i\u0062\u0069\u006c\u0069\u0074y\u003a\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _eegggc, _bafbee.FontName)
	}
	_bafbee.FontFamily = _gacb.Get("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079")
	_bafbee.FontStretch = _gacb.Get("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068")
	_bafbee.FontWeight = _gacb.Get("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074")
	_bafbee.Flags = _gacb.Get("\u0046\u006c\u0061g\u0073")
	_bafbee.FontBBox = _gacb.Get("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078")
	_bafbee.ItalicAngle = _gacb.Get("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065")
	_bafbee.Ascent = _gacb.Get("\u0041\u0073\u0063\u0065\u006e\u0074")
	_bafbee.Descent = _gacb.Get("\u0044e\u0073\u0063\u0065\u006e\u0074")
	_bafbee.Leading = _gacb.Get("\u004ce\u0061\u0064\u0069\u006e\u0067")
	_bafbee.CapHeight = _gacb.Get("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t")
	_bafbee.XHeight = _gacb.Get("\u0058H\u0065\u0069\u0067\u0068\u0074")
	_bafbee.StemV = _gacb.Get("\u0053\u0074\u0065m\u0056")
	_bafbee.StemH = _gacb.Get("\u0053\u0074\u0065m\u0048")
	_bafbee.AvgWidth = _gacb.Get("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068")
	_bafbee.MaxWidth = _gacb.Get("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068")
	_bafbee.MissingWidth = _gacb.Get("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068")
	_bafbee.FontFile = _gacb.Get("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065")
	_bafbee.FontFile2 = _gacb.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652")
	_bafbee.FontFile3 = _gacb.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653")
	_bafbee.CharSet = _gacb.Get("\u0043h\u0061\u0072\u0053\u0065\u0074")
	_bafbee.Style = _gacb.Get("\u0053\u0074\u0079l\u0065")
	_bafbee.Lang = _gacb.Get("\u004c\u0061\u006e\u0067")
	_bafbee.FD = _gacb.Get("\u0046\u0044")
	_bafbee.CIDSet = _gacb.Get("\u0043\u0049\u0044\u0053\u0065\u0074")
	if _bafbee.Flags != nil {
		if _edcgb, _bedff := _cde.GetIntVal(_bafbee.Flags); _bedff {
			_bafbee._gbffb = _edcgb
		}
	}
	if _bafbee.MissingWidth != nil {
		if _ceaf, _debfb := _cde.GetNumberAsFloat(_bafbee.MissingWidth); _debfb == nil {
			_bafbee._efgg = _ceaf
		}
	}
	if _bafbee.FontFile != nil {
		_cfcd, _fgcc := _cabad(_bafbee.FontFile)
		if _fgcc != nil {
			return _bafbee, _fgcc
		}
		_ad.Log.Trace("f\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u003d\u0025\u0073", _cfcd)
		_bafbee.fontFile = _cfcd
	}
	if _bafbee.FontFile2 != nil {
		_abeec, _ddca := _fe.NewFontFile2FromPdfObject(_bafbee.FontFile2)
		if _ddca != nil {
			return _bafbee, _ddca
		}
		_ad.Log.Trace("\u0066\u006f\u006et\u0046\u0069\u006c\u0065\u0032\u003d\u0025\u0073", _abeec.String())
		_bafbee._ggdga = &_abeec
	}
	return _bafbee, nil
}

// ToPdfObject implements interface PdfModel.
func (_agbg *PdfAnnotationStrikeOut) ToPdfObject() _cde.PdfObject {
	_agbg.PdfAnnotation.ToPdfObject()
	_eega := _agbg._bddg
	_cbad := _eega.PdfObject.(*_cde.PdfObjectDictionary)
	_agbg.PdfAnnotationMarkup.appendToPdfDictionary(_cbad)
	_cbad.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t"))
	_cbad.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _agbg.QuadPoints)
	return _eega
}

// SetDecode sets the decode image float slice.
func (_gebb *Image) SetDecode(decode []float64) { _gebb._aaafb = decode }
func _acaga() *modelManager {
	_fbcef := modelManager{}
	_fbcef._bfaea = map[PdfModel]_cde.PdfObject{}
	_fbcef._dfeee = map[_cde.PdfObject]PdfModel{}
	return &_fbcef
}

var _dfbc = map[string]string{"\u0053\u0079\u006d\u0062\u006f\u006c": "\u0053\u0079\u006d\u0062\u006f\u006c\u0045\u006e\u0063o\u0064\u0069\u006e\u0067", "\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073": "Z\u0061p\u0066\u0044\u0069\u006e\u0067\u0062\u0061\u0074s\u0045\u006e\u0063\u006fdi\u006e\u0067"}

func (_bec *PdfReader) newPdfActionHideFromDict(_gab *_cde.PdfObjectDictionary) (*PdfActionHide, error) {
	return &PdfActionHide{T: _gab.Get("\u0054"), H: _gab.Get("\u0048")}, nil
}

// NewPdfPageResources returns a new PdfPageResources object.
func NewPdfPageResources() *PdfPageResources {
	_gaecb := &PdfPageResources{}
	_gaecb._eaegg = _cde.MakeDict()
	return _gaecb
}

// GetOutlines returns a high-level Outline object, based on the outline tree
// of the reader.
func (_abge *PdfReader) GetOutlines() (*Outline, error) {
	if _abge == nil {
		return nil, _ceg.New("\u0063\u0061n\u006e\u006f\u0074\u0020c\u0072\u0065a\u0074\u0065\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u006e\u0069\u006c\u0020\u0072e\u0061\u0064\u0065\u0072")
	}
	_cacbd := _abge.GetOutlineTree()
	if _cacbd == nil {
		return nil, _ceg.New("\u0074\u0068\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u0064\u006f\u0065\u0073\u0020n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065")
	}
	var _fabbc func(_geafg *PdfOutlineTreeNode, _cafg *[]*OutlineItem)
	_fabbc = func(_bebab *PdfOutlineTreeNode, _edede *[]*OutlineItem) {
		if _bebab == nil {
			return
		}
		if _bebab._fbeea == nil {
			_ad.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020m\u0069\u0073\u0073\u0069ng \u006fut\u006c\u0069\u006e\u0065\u0020\u0065\u006etr\u0079\u0020\u0063\u006f\u006e\u0074\u0065x\u0074")
			return
		}
		var _ddcbc *OutlineItem
		if _edceb, _aagg := _bebab._fbeea.(*PdfOutlineItem); _aagg {
			_ffeaag := _edceb.Dest
			if (_ffeaag == nil || _cde.IsNullObject(_ffeaag)) && _edceb.A != nil {
				if _gbfed, _ccced := _cde.GetDict(_edceb.A); _ccced {
					if _gecab, _cgde := _cde.GetArray(_gbfed.Get("\u0044")); _cgde {
						_ffeaag = _gecab
					} else {
						_eddbg, _cdgda := _cde.GetString(_gbfed.Get("\u0044"))
						if !_cdgda {
							return
						}
						_bbdab, _cdgda := _abge._efabe.Get("\u004e\u0061\u006de\u0073").(*_cde.PdfObjectReference)
						if !_cdgda {
							return
						}
						_cddf, _deeea := _abge._aggcgb.LookupByReference(*_bbdab)
						if _deeea != nil {
							_ad.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0028\u0025\u0073\u0029", _deeea.Error())
							return
						}
						_fedce, _cdgda := _cddf.(*_cde.PdfIndirectObject)
						if !_cdgda {
							return
						}
						_decb := map[_cde.PdfObject]struct{}{}
						_deeea = _abge.buildNameNodes(_fedce, _decb)
						if _deeea != nil {
							_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006ea\u006d\u0065\u0020\u006e\u006fd\u0065\u0073 \u0028\u0025\u0073\u0029", _deeea.Error())
							return
						}
						for _fcdcb := range _decb {
							_eccfd, _dabff := _cde.GetDict(_fcdcb)
							if !_dabff {
								continue
							}
							_gagdg, _dabff := _cde.GetArray(_eccfd.Get("\u004e\u0061\u006de\u0073"))
							if !_dabff {
								continue
							}
							for _fgaag, _eagb := range _gagdg.Elements() {
								switch _eagb.(type) {
								case *_cde.PdfObjectString:
									if _eagb.String() == _eddbg.String() {
										if _dedgd := _gagdg.Get(_fgaag + 1); _dedgd != nil {
											if _ecda, _feaea := _cde.GetDict(_dedgd); _feaea {
												_ffeaag = _ecda.Get("\u0044")
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
			var _eabd OutlineDest
			if _ffeaag != nil && !_cde.IsNullObject(_ffeaag) {
				if _ccdcac, _affe := _fdce(_ffeaag, _abge); _affe == nil {
					_eabd = *_ccdcac
				} else {
					_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020p\u0061\u0072\u0073\u0065\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065\u0073\u0074\u0020\u0028\u0025\u0076\u0029\u003a\u0020\u0025\u0076", _ffeaag, _affe)
				}
			}
			_ddcbc = NewOutlineItem(_edceb.Title.Decoded(), _eabd)
			*_edede = append(*_edede, _ddcbc)
			if _edceb.Next != nil {
				_fabbc(_edceb.Next, _edede)
			}
		}
		if _bebab.First != nil {
			if _ddcbc != nil {
				_edede = &_ddcbc.Entries
			}
			_fabbc(_bebab.First, _edede)
		}
	}
	_dbeaf := NewOutline()
	_fabbc(_cacbd, &_dbeaf.Entries)
	return _dbeaf, nil
}
func (_ebea *PdfReader) newPdfAnnotationStrikeOut(_afbge *_cde.PdfObjectDictionary) (*PdfAnnotationStrikeOut, error) {
	_egdb := PdfAnnotationStrikeOut{}
	_efee, _aeeg := _ebea.newPdfAnnotationMarkupFromDict(_afbge)
	if _aeeg != nil {
		return nil, _aeeg
	}
	_egdb.PdfAnnotationMarkup = _efee
	_egdb.QuadPoints = _afbge.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_egdb, nil
}

// PdfBorderEffect represents a PDF border effect.
type PdfBorderEffect struct {
	S *BorderEffect
	I *float64
}

func (_eacd *PdfFont) baseFields() *fontCommon {
	if _eacd._gbcff == nil {
		_ad.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0062\u0061\u0073\u0065\u0046\u0069\u0065l\u0064s\u002e \u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e")
		return nil
	}
	return _eacd._gbcff.baseFields()
}

// NewPdfAnnotationSound returns a new sound annotation.
func NewPdfAnnotationSound() *PdfAnnotationSound {
	_gaf := NewPdfAnnotation()
	_bdf := &PdfAnnotationSound{}
	_bdf.PdfAnnotation = _gaf
	_bdf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gaf.SetContext(_bdf)
	return _bdf
}

// PdfRectangle is a definition of a rectangle.
type PdfRectangle struct {
	Llx float64
	Lly float64
	Urx float64
	Ury float64
}

// AllFields returns a flattened list of all fields in the form.
func (_ebcf *PdfAcroForm) AllFields() []*PdfField {
	if _ebcf == nil {
		return nil
	}
	var _ceee []*PdfField
	if _ebcf.Fields != nil {
		for _, _dcee := range *_ebcf.Fields {
			_ceee = append(_ceee, _bgcac(_dcee)...)
		}
	}
	return _ceee
}
func (_fcg *PdfReader) newPdfActionGotoEFromDict(_edb *_cde.PdfObjectDictionary) (*PdfActionGoToE, error) {
	_fgde, _fag := _beed(_edb.Get("\u0046"))
	if _fag != nil {
		return nil, _fag
	}
	return &PdfActionGoToE{D: _edb.Get("\u0044"), NewWindow: _edb.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), T: _edb.Get("\u0054"), F: _fgde}, nil
}
func (_afgdd *PdfReader) newPdfAnnotationLinkFromDict(_fcbc *_cde.PdfObjectDictionary) (*PdfAnnotationLink, error) {
	_acbd := PdfAnnotationLink{}
	_acbd.A = _fcbc.Get("\u0041")
	_acbd.Dest = _fcbc.Get("\u0044\u0065\u0073\u0074")
	_acbd.H = _fcbc.Get("\u0048")
	_acbd.PA = _fcbc.Get("\u0050\u0041")
	_acbd.QuadPoints = _fcbc.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_acbd.BS = _fcbc.Get("\u0042\u0053")
	return &_acbd, nil
}

// NewPdfColorspaceDeviceGray returns a new grayscale colorspace.
func NewPdfColorspaceDeviceGray() *PdfColorspaceDeviceGray { return &PdfColorspaceDeviceGray{} }

// ToPdfObject implements interface PdfModel.
func (_cgac *PdfAnnotationTrapNet) ToPdfObject() _cde.PdfObject {
	_cgac.PdfAnnotation.ToPdfObject()
	_ddcef := _cgac._bddg
	_gfgf := _ddcef.PdfObject.(*_cde.PdfObjectDictionary)
	_gfgf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0054r\u0061\u0070\u004e\u0065\u0074"))
	return _ddcef
}

// PdfTilingPattern is a Tiling pattern that consists of repetitions of a pattern cell with defined intervals.
// It is a type 1 pattern. (PatternType = 1).
// A tiling pattern is represented by a stream object, where the stream content is
// a content stream that describes the pattern cell.
type PdfTilingPattern struct {
	*PdfPattern
	PaintType  *_cde.PdfObjectInteger
	TilingType *_cde.PdfObjectInteger
	BBox       *PdfRectangle
	XStep      *_cde.PdfObjectFloat
	YStep      *_cde.PdfObjectFloat
	Resources  *PdfPageResources
	Matrix     *_cde.PdfObjectArray
}

// GetEncryptionMethod returns a descriptive information string about the encryption method used.
func (_ecdea *PdfReader) GetEncryptionMethod() string {
	_bgag := _ecdea._aggcgb.GetCrypter()
	return _bgag.String()
}
func _cfdbb(_acggc _cde.PdfObject) (PdfFunction, error) {
	_acggc = _cde.ResolveReference(_acggc)
	if _cfcab, _eeec := _acggc.(*_cde.PdfObjectStream); _eeec {
		_aaga := _cfcab.PdfObjectDictionary
		_fbgd, _fdggf := _aaga.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_cde.PdfObjectInteger)
		if !_fdggf {
			_ad.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _ceg.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_fbgd == 0 {
			return _gbfac(_cfcab)
		} else if *_fbgd == 4 {
			return _cefb(_cfcab)
		} else {
			return nil, _ceg.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _gfgfa, _gfbce := _acggc.(*_cde.PdfIndirectObject); _gfbce {
		_bdba, _bggabg := _gfgfa.PdfObject.(*_cde.PdfObjectDictionary)
		if !_bggabg {
			_ad.Log.Error("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006eg\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _ceg.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		_acda, _bggabg := _bdba.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_cde.PdfObjectInteger)
		if !_bggabg {
			_ad.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _ceg.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_acda == 2 {
			return _cbbd(_gfgfa)
		} else if *_acda == 3 {
			return _eefcg(_gfgfa)
		} else {
			return nil, _ceg.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _fbdga, _dafag := _acggc.(*_cde.PdfObjectDictionary); _dafag {
		_aage, _ceafd := _fbdga.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_cde.PdfObjectInteger)
		if !_ceafd {
			_ad.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _ceg.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_aage == 2 {
			return _cbbd(_fbdga)
		} else if *_aage == 3 {
			return _eefcg(_fbdga)
		} else {
			return nil, _ceg.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else {
		_ad.Log.Debug("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0054\u0079\u0070\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0023\u0076", _acggc)
		return nil, _ceg.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
}
func (_gbbcbe *PdfWriter) setCatalogVersion() {
	_gbbcbe._fedbb.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _cde.MakeName(_ee.Sprintf("\u0025\u0064\u002e%\u0064", _gbbcbe._cgdcc.Major, _gbbcbe._cgdcc.Minor)))
}

// GetContainingPdfObject returns the container of the shading object (indirect object).
func (_cbead *PdfShading) GetContainingPdfObject() _cde.PdfObject { return _cbead._dffg }

// NewXObjectForm creates a brand new XObject Form. Creates a new underlying PDF object stream primitive.
func NewXObjectForm() *XObjectForm {
	_edcgd := &XObjectForm{}
	_adfde := &_cde.PdfObjectStream{}
	_adfde.PdfObjectDictionary = _cde.MakeDict()
	_edcgd._ecfbd = _adfde
	return _edcgd
}
func (_cecbg *PdfColorspaceLab) String() string { return "\u004c\u0061\u0062" }
func _bbefa(_ffed string) (map[_gc.CharCode]_gc.GlyphName, error) {
	_faab := _dac.Split(_ffed, "\u000a")
	_fbaed := make(map[_gc.CharCode]_gc.GlyphName)
	for _, _aeeed := range _faab {
		_ddcd := _eadff.FindStringSubmatch(_aeeed)
		if _ddcd == nil {
			continue
		}
		_fdba, _ddggc := _ddcd[1], _ddcd[2]
		_fcdfd, _cggba := _gb.Atoi(_fdba)
		if _cggba != nil {
			_ad.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0042\u0061\u0064\u0020\u0065\u006e\u0063\u006fd\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u002e \u0025\u0071", _aeeed)
			return nil, _cde.ErrTypeError
		}
		_fbaed[_gc.CharCode(_fcdfd)] = _gc.GlyphName(_ddggc)
	}
	_ad.Log.Trace("g\u0065\u0074\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0073\u003a\u0020\u006b\u0065\u0079V\u0061\u006c\u0075e\u0073=\u0025\u0023\u0076", _fbaed)
	return _fbaed, nil
}
func (_dbgaf *PdfWriter) flushWriter() error {
	if _dbgaf._fggeef == nil {
		_dbgaf._fggeef = _dbgaf._cefdd.Flush()
	}
	return _dbgaf._fggeef
}

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_fcgce *PdfWriter) SetOptimizer(optimizer Optimizer) { _fcgce._cgee = optimizer }

// R returns the value of the red component of the color.
func (_gfafc *PdfColorDeviceRGB) R() float64 { return _gfafc[0] }

// ToPdfObject implements interface PdfModel.
func (_faf *PdfActionTrans) ToPdfObject() _cde.PdfObject {
	_faf.PdfAction.ToPdfObject()
	_agcc := _faf._bc
	_gcg := _agcc.PdfObject.(*_cde.PdfObjectDictionary)
	_gcg.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeTrans)))
	_gcg.SetIfNotNil("\u0054\u0072\u0061n\u0073", _faf.Trans)
	return _agcc
}

// PdfAnnotationText represents Text annotations.
// (Section 12.5.6.4 p. 402).
type PdfAnnotationText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Open       _cde.PdfObject
	Name       _cde.PdfObject
	State      _cde.PdfObject
	StateModel _cde.PdfObject
}

// NewPdfActionSound returns a new "sound" action.
func NewPdfActionSound() *PdfActionSound {
	_aa := NewPdfAction()
	_cf := &PdfActionSound{}
	_cf.PdfAction = _aa
	_aa.SetContext(_cf)
	return _cf
}

// ToPdfObject returns the PDF representation of the pattern.
func (_cdccae *PdfPattern) ToPdfObject() _cde.PdfObject {
	_beagb := _cdccae.getDict()
	_beagb.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_beagb.Set("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065", _cde.MakeInteger(_cdccae.PatternType))
	return _cdccae._eecac
}

// GetContainingPdfObject implements interface PdfModel.
func (_eaag *PdfFilespec) GetContainingPdfObject() _cde.PdfObject { return _eaag._bgac }
func (_eebf *PdfReader) newPdfAnnotationCaretFromDict(_geab *_cde.PdfObjectDictionary) (*PdfAnnotationCaret, error) {
	_caa := PdfAnnotationCaret{}
	_dfdb, _cgb := _eebf.newPdfAnnotationMarkupFromDict(_geab)
	if _cgb != nil {
		return nil, _cgb
	}
	_caa.PdfAnnotationMarkup = _dfdb
	_caa.RD = _geab.Get("\u0052\u0044")
	_caa.Sy = _geab.Get("\u0053\u0079")
	return &_caa, nil
}

// ToPdfObject returns the PdfFontDescriptor as a PDF dictionary inside an indirect object.
func (_bcegd *PdfFontDescriptor) ToPdfObject() _cde.PdfObject {
	_efdg := _cde.MakeDict()
	if _bcegd._fbff == nil {
		_bcegd._fbff = &_cde.PdfIndirectObject{}
	}
	_bcegd._fbff.PdfObject = _efdg
	_efdg.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072"))
	if _bcegd.FontName != nil {
		_efdg.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _bcegd.FontName)
	}
	if _bcegd.FontFamily != nil {
		_efdg.Set("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079", _bcegd.FontFamily)
	}
	if _bcegd.FontStretch != nil {
		_efdg.Set("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068", _bcegd.FontStretch)
	}
	if _bcegd.FontWeight != nil {
		_efdg.Set("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074", _bcegd.FontWeight)
	}
	if _bcegd.Flags != nil {
		_efdg.Set("\u0046\u006c\u0061g\u0073", _bcegd.Flags)
	}
	if _bcegd.FontBBox != nil {
		_efdg.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _bcegd.FontBBox)
	}
	if _bcegd.ItalicAngle != nil {
		_efdg.Set("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065", _bcegd.ItalicAngle)
	}
	if _bcegd.Ascent != nil {
		_efdg.Set("\u0041\u0073\u0063\u0065\u006e\u0074", _bcegd.Ascent)
	}
	if _bcegd.Descent != nil {
		_efdg.Set("\u0044e\u0073\u0063\u0065\u006e\u0074", _bcegd.Descent)
	}
	if _bcegd.Leading != nil {
		_efdg.Set("\u004ce\u0061\u0064\u0069\u006e\u0067", _bcegd.Leading)
	}
	if _bcegd.CapHeight != nil {
		_efdg.Set("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t", _bcegd.CapHeight)
	}
	if _bcegd.XHeight != nil {
		_efdg.Set("\u0058H\u0065\u0069\u0067\u0068\u0074", _bcegd.XHeight)
	}
	if _bcegd.StemV != nil {
		_efdg.Set("\u0053\u0074\u0065m\u0056", _bcegd.StemV)
	}
	if _bcegd.StemH != nil {
		_efdg.Set("\u0053\u0074\u0065m\u0048", _bcegd.StemH)
	}
	if _bcegd.AvgWidth != nil {
		_efdg.Set("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068", _bcegd.AvgWidth)
	}
	if _bcegd.MaxWidth != nil {
		_efdg.Set("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068", _bcegd.MaxWidth)
	}
	if _bcegd.MissingWidth != nil {
		_efdg.Set("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068", _bcegd.MissingWidth)
	}
	if _bcegd.FontFile != nil {
		_efdg.Set("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065", _bcegd.FontFile)
	}
	if _bcegd.FontFile2 != nil {
		_efdg.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652", _bcegd.FontFile2)
	}
	if _bcegd.FontFile3 != nil {
		_efdg.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653", _bcegd.FontFile3)
	}
	if _bcegd.CharSet != nil {
		_efdg.Set("\u0043h\u0061\u0072\u0053\u0065\u0074", _bcegd.CharSet)
	}
	if _bcegd.Style != nil {
		_efdg.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _bcegd.FontName)
	}
	if _bcegd.Lang != nil {
		_efdg.Set("\u004c\u0061\u006e\u0067", _bcegd.Lang)
	}
	if _bcegd.FD != nil {
		_efdg.Set("\u0046\u0044", _bcegd.FD)
	}
	if _bcegd.CIDSet != nil {
		_efdg.Set("\u0043\u0049\u0044\u0053\u0065\u0074", _bcegd.CIDSet)
	}
	return _bcegd._fbff
}
func (_fafbb SignatureValidationResult) String() string {
	var _fdeef _ede.Buffer
	_fdeef.WriteString(_ee.Sprintf("\u004ea\u006d\u0065\u003a\u0020\u0025\u0073\n", _fafbb.Name))
	if _fafbb.Date._aedbfg > 0 {
		_fdeef.WriteString(_ee.Sprintf("\u0044a\u0074\u0065\u003a\u0020\u0025\u0073\n", _fafbb.Date.ToGoTime().String()))
	} else {
		_fdeef.WriteString("\u0044\u0061\u0074\u0065 n\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_fafbb.Reason) > 0 {
		_fdeef.WriteString(_ee.Sprintf("R\u0065\u0061\u0073\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _fafbb.Reason))
	} else {
		_fdeef.WriteString("N\u006f \u0072\u0065\u0061\u0073\u006f\u006e\u0020\u0073p\u0065\u0063\u0069\u0066ie\u0064\u000a")
	}
	if len(_fafbb.Location) > 0 {
		_fdeef.WriteString(_ee.Sprintf("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _fafbb.Location))
	} else {
		_fdeef.WriteString("\u004c\u006f\u0063at\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_fafbb.ContactInfo) > 0 {
		_fdeef.WriteString(_ee.Sprintf("\u0043\u006f\u006e\u0074\u0061\u0063\u0074\u0020\u0049\u006e\u0066\u006f:\u0020\u0025\u0073\u000a", _fafbb.ContactInfo))
	} else {
		_fdeef.WriteString("C\u006f\u006e\u0074\u0061\u0063\u0074 \u0069\u006e\u0066\u006f\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063i\u0066i\u0065\u0064\u000a")
	}
	_fdeef.WriteString(_ee.Sprintf("F\u0069\u0065\u006c\u0064\u0073\u003a\u0020\u0025\u0064\u000a", len(_fafbb.Fields)))
	if _fafbb.IsSigned {
		_fdeef.WriteString("S\u0069\u0067\u006e\u0065\u0064\u003a \u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073\u0020\u0073i\u0067n\u0065\u0064\u000a")
	} else {
		_fdeef.WriteString("\u0053\u0069\u0067\u006eed\u003a\u0020\u004e\u006f\u0074\u0020\u0073\u0069\u0067\u006e\u0065\u0064\u000a")
	}
	if _fafbb.IsVerified {
		_fdeef.WriteString("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0076\u0061\u006c\u0069\u0064\u0061t\u0069\u006f\u006e\u003a\u0020\u0049\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u000a")
	} else {
		_fdeef.WriteString("\u0053\u0069\u0067\u006e\u0061\u0074u\u0072\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e:\u0020\u0049\u0073\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u000a")
	}
	if _fafbb.IsTrusted {
		_fdeef.WriteString("\u0054\u0072\u0075\u0073\u0074\u0065\u0064\u003a\u0020\u0043\u0065\u0072\u0074\u0069\u0066i\u0063a\u0074\u0065\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0073\u0074\u0065\u0064\u000a")
	} else {
		_fdeef.WriteString("\u0054\u0072\u0075s\u0074\u0065\u0064\u003a \u0055\u006e\u0074\u0072\u0075\u0073\u0074e\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u000a")
	}
	if !_fafbb.GeneralizedTime.IsZero() {
		_fdeef.WriteString(_ee.Sprintf("G\u0065n\u0065\u0072\u0061\u006c\u0069\u007a\u0065\u0064T\u0069\u006d\u0065\u003a %\u0073\u000a", _fafbb.GeneralizedTime.String()))
	}
	if _fafbb.DiffResults != nil {
		_fdeef.WriteString(_ee.Sprintf("\u0064\u0069\u0066\u0066 i\u0073\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u003a\u0020\u0025v\u000a", _fafbb.DiffResults.IsPermitted()))
		if len(_fafbb.DiffResults.Warnings) > 0 {
			_fdeef.WriteString("\u004d\u0044\u0050\u0020\u0077\u0061\u0072\u006e\u0069n\u0067\u0073\u003a\u000a")
			for _, _bagbe := range _fafbb.DiffResults.Warnings {
				_fdeef.WriteString(_ee.Sprintf("\u0009\u0025\u0073\u000a", _bagbe))
			}
		}
		if len(_fafbb.DiffResults.Errors) > 0 {
			_fdeef.WriteString("\u004d\u0044\u0050 \u0065\u0072\u0072\u006f\u0072\u0073\u003a\u000a")
			for _, _edefe := range _fafbb.DiffResults.Errors {
				_fdeef.WriteString(_ee.Sprintf("\u0009\u0025\u0073\u000a", _edefe))
			}
		}
	}
	return _fdeef.String()
}

// ToOutlineTree returns a low level PdfOutlineTreeNode object, based on
// the current instance.
func (_agafc *Outline) ToOutlineTree() *PdfOutlineTreeNode {
	return &_agafc.ToPdfOutline().PdfOutlineTreeNode
}

// Read reads an image and loads into a new Image object with an RGB
// colormap and 8 bits per component.
func (_bcgge DefaultImageHandler) Read(reader _f.Reader) (*Image, error) {
	_cbae, _, _aagba := _gf.Decode(reader)
	if _aagba != nil {
		_ad.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _aagba)
		return nil, _aagba
	}
	return _bcgge.NewImageFromGoImage(_cbae)
}

// IsValid checks if the given pdf output intent type is valid.
func (_defgd PdfOutputIntentType) IsValid() bool {
	return _defgd >= PdfOutputIntentTypeA1 && _defgd <= PdfOutputIntentTypeX
}

// GetContentStreamWithEncoder returns the pattern cell's content stream and its encoder
func (_accgc *PdfTilingPattern) GetContentStreamWithEncoder() ([]byte, _cde.StreamEncoder, error) {
	_dabb, _dcgea := _accgc._eecac.(*_cde.PdfObjectStream)
	if !_dcgea {
		_ad.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _accgc._eecac)
		return nil, nil, _cde.ErrTypeError
	}
	_fccdg, _gbbdf := _cde.DecodeStream(_dabb)
	if _gbbdf != nil {
		_ad.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020s\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0065\u0072\u0072\u003a\u0020\u0025\u0076", _gbbdf)
		return nil, nil, _gbbdf
	}
	_cadf, _gbbdf := _cde.NewEncoderFromStream(_dabb)
	if _gbbdf != nil {
		_ad.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020f\u0069\u006e\u0064\u0069\u006e\u0067 \u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _gbbdf)
		return nil, nil, _gbbdf
	}
	return _fccdg, _cadf, nil
}
func (_gbed *PdfReader) newPdfSignatureReferenceFromDict(_gcbfa *_cde.PdfObjectDictionary) (*PdfSignatureReference, error) {
	if _bdbca, _dcecb := _gbed._bedfa.GetModelFromPrimitive(_gcbfa).(*PdfSignatureReference); _dcecb {
		return _bdbca, nil
	}
	_afcc := &PdfSignatureReference{_fggee: _gcbfa, Data: _gcbfa.Get("\u0044\u0061\u0074\u0061")}
	var _gcaeb bool
	_afcc.Type, _ = _cde.GetName(_gcbfa.Get("\u0054\u0079\u0070\u0065"))
	_afcc.TransformMethod, _gcaeb = _cde.GetName(_gcbfa.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"))
	if !_gcaeb {
		_ad.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0072\u0061\u006e\u0073\u0066o\u0072\u006dM\u0065\u0074h\u006f\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020in\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020m\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_afcc.TransformParams, _ = _cde.GetDict(_gcbfa.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"))
	_afcc.DigestMethod, _ = _cde.GetName(_gcbfa.Get("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064"))
	return _afcc, nil
}
func (_egeb *PdfAcroForm) filteredFields(_eaeg FieldFilterFunc, _aedad bool) []*PdfField {
	if _egeb == nil {
		return nil
	}
	return _dfbeg(_egeb.Fields, _eaeg, _aedad)
}

// Resample resamples the image data converting from current BitsPerComponent to a target BitsPerComponent
// value.  Sets the image's BitsPerComponent to the target value following resampling.
//
// For example, converting an 8-bit RGB image to 1-bit grayscale (common for scanned images):
//   // Convert RGB image to grayscale.
//   rgbColorSpace := pdf.NewPdfColorspaceDeviceRGB()
//   grayImage, err := rgbColorSpace.ImageToGray(rgbImage)
//   if err != nil {
//     return err
//   }
//   // Resample as 1 bit.
//   grayImage.Resample(1)
func (_aaafe *Image) Resample(targetBitsPerComponent int64) {
	if _aaafe.BitsPerComponent == targetBitsPerComponent {
		return
	}
	_feca := _aaafe.GetSamples()
	if targetBitsPerComponent < _aaafe.BitsPerComponent {
		_efaef := _aaafe.BitsPerComponent - targetBitsPerComponent
		for _ffeaa := range _feca {
			_feca[_ffeaa] >>= uint(_efaef)
		}
	} else if targetBitsPerComponent > _aaafe.BitsPerComponent {
		_abadb := targetBitsPerComponent - _aaafe.BitsPerComponent
		for _gdefc := range _feca {
			_feca[_gdefc] <<= uint(_abadb)
		}
	}
	_aaafe.BitsPerComponent = targetBitsPerComponent
	if _aaafe.BitsPerComponent < 8 {
		_aaafe.resampleLowBits(_feca)
		return
	}
	_aecga := _ff.BytesPerLine(int(_aaafe.Width), int(_aaafe.BitsPerComponent), _aaafe.ColorComponents)
	_dgfdc := make([]byte, _aecga*int(_aaafe.Height))
	var (
		_egbfb, _aefgd, _abeee, _gadfb int
		_ebeg                          uint32
	)
	for _abeee = 0; _abeee < int(_aaafe.Height); _abeee++ {
		_egbfb = _abeee * _aecga
		_aefgd = (_abeee+1)*_aecga - 1
		_fcee := _cae.ResampleUint32(_feca[_egbfb:_aefgd], int(targetBitsPerComponent), 8)
		for _gadfb, _ebeg = range _fcee {
			_dgfdc[_gadfb+_egbfb] = byte(_ebeg)
		}
	}
	_aaafe.Data = _dgfdc
}

// PdfAnnotationPopup represents Popup annotations.
// (Section 12.5.6.14).
type PdfAnnotationPopup struct {
	*PdfAnnotation
	Parent _cde.PdfObject
	Open   _cde.PdfObject
}

func _bfad(_gfdgf []byte) bool {
	if len(_gfdgf) < 4 {
		return true
	}
	for _gdfff := range _gfdgf[:4] {
		_efbg := rune(_gdfff)
		if !_da.Is(_da.ASCII_Hex_Digit, _efbg) && !_da.IsSpace(_efbg) {
			return true
		}
	}
	return false
}

// Flags returns the field flags for the field accounting for any inherited flags.
func (_fggc *PdfField) Flags() FieldFlag {
	var _ceeg FieldFlag
	_gfbc, _fbcag := _fggc.inherit(func(_ebgf *PdfField) bool {
		if _ebgf.Ff != nil {
			_ceeg = FieldFlag(*_ebgf.Ff)
			return true
		}
		return false
	})
	if _fbcag != nil {
		_ad.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0069\u006e\u0067\u0020\u0066\u006c\u0061\u0067\u0073\u0020\u0076\u0069\u0061\u0020\u0069\u006e\u0068\u0065\u0072\u0069t\u0061\u006e\u0063\u0065\u003a\u0020\u0025\u0076", _fbcag)
	}
	if !_gfbc {
		_ad.Log.Trace("N\u006f\u0020\u0066\u0069\u0065\u006cd\u0020\u0066\u006c\u0061\u0067\u0073 \u0066\u006f\u0075\u006e\u0064\u0020\u002d \u0061\u0073\u0073\u0075\u006d\u0065\u0020\u0063\u006c\u0065a\u0072")
	}
	return _ceeg
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

// GetContainingPdfObject implements model.PdfModel interface.
func (_cggce *PdfOutputIntent) GetContainingPdfObject() _cde.PdfObject { return _cggce._acbcb }

// K returns the value of the key component of the color.
func (_dbgg *PdfColorDeviceCMYK) K() float64 { return _dbgg[3] }

// DecodeArray returns the range of color component values in CalGray colorspace.
func (_cbdg *PdfColorspaceCalGray) DecodeArray() []float64 { return []float64{0.0, 1.0} }

// NewPdfColorspaceDeviceRGB returns a new RGB colorspace object.
func NewPdfColorspaceDeviceRGB() *PdfColorspaceDeviceRGB { return &PdfColorspaceDeviceRGB{} }

// NewStandard14FontWithEncoding returns the standard 14 font named `basefont` as a *PdfFont and
// a TextEncoder that encodes all the runes in `alphabet`, or an error if this is not possible.
// An error can occur if `basefont` is not one the standard 14 font names.
func NewStandard14FontWithEncoding(basefont StdFontName, alphabet map[rune]int) (*PdfFont, _gc.SimpleEncoder, error) {
	_ddgac, _cfgfd := _eabg(basefont)
	if _cfgfd != nil {
		return nil, nil, _cfgfd
	}
	_cfdf, _aadag := _ddgac.Encoder().(_gc.SimpleEncoder)
	if !_aadag {
		return nil, nil, _ee.Errorf("\u006f\u006e\u006c\u0079\u0020s\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u002c\u0020\u0067\u006f\u0074\u0020\u0025\u0054", _ddgac.Encoder())
	}
	_eaedd := make(map[rune]_gc.GlyphName)
	for _gedgb := range alphabet {
		if _, _fedf := _cfdf.RuneToCharcode(_gedgb); !_fedf {
			_, _gcca := _ddgac._gggc.Read(_gedgb)
			if !_gcca {
				_ad.Log.Trace("r\u0075\u006e\u0065\u0020\u0025\u0023x\u003d\u0025\u0071\u0020\u006e\u006f\u0074\u0020\u0069n\u0020\u0074\u0068e\u0020f\u006f\u006e\u0074", _gedgb, _gedgb)
				continue
			}
			_cagec, _gcca := _gc.RuneToGlyph(_gedgb)
			if !_gcca {
				_ad.Log.Debug("\u006eo\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u0020\u0025\u0023\u0078\u003d\u0025\u0071", _gedgb, _gedgb)
				continue
			}
			if len(_eaedd) >= 255 {
				return nil, nil, _ceg.New("\u0074\u006f\u006f\u0020\u006d\u0061\u006e\u0079\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020f\u006f\u0072\u0020\u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067")
			}
			_eaedd[_gedgb] = _cagec
		}
	}
	var (
		_eecaa []_gc.CharCode
		_cgcd  []_gc.CharCode
	)
	for _gdaf := _gc.CharCode(1); _gdaf <= 0xff; _gdaf++ {
		_dafa, _dedaf := _cfdf.CharcodeToRune(_gdaf)
		if !_dedaf {
			_eecaa = append(_eecaa, _gdaf)
			continue
		}
		if _, _dedaf = alphabet[_dafa]; !_dedaf {
			_cgcd = append(_cgcd, _gdaf)
		}
	}
	_ebec := append(_eecaa, _cgcd...)
	if len(_ebec) < len(_eaedd) {
		return nil, nil, _ee.Errorf("n\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u006ec\u006f\u0064\u0065\u0020\u0025\u0064\u0020ru\u006e\u0065\u0073\u002c \u0062\u0075\u0074\u0020\u0068\u0061\u0076\u0065\u0020on\u006c\u0079 \u0025\u0064\u0020\u0073\u006c\u006f\u0074\u0073", len(_eaedd), len(_ebec))
	}
	_dgecf := make([]rune, 0, len(_eaedd))
	for _agdca := range _eaedd {
		_dgecf = append(_dgecf, _agdca)
	}
	_dd.Slice(_dgecf, func(_bfbg, _abcge int) bool { return _dgecf[_bfbg] < _dgecf[_abcge] })
	_fdab := make(map[_gc.CharCode]_gc.GlyphName, len(_dgecf))
	for _, _bgec := range _dgecf {
		_fdgfb := _ebec[0]
		_ebec = _ebec[1:]
		_fdab[_fdgfb] = _eaedd[_bgec]
	}
	_cfdf = _gc.ApplyDifferences(_cfdf, _fdab)
	_ddgac.SetEncoder(_cfdf)
	return &PdfFont{_gbcff: &_ddgac}, _cfdf, nil
}

// NewPdfOutputIntentFromPdfObject creates a new PdfOutputIntent from the input core.PdfObject.
func NewPdfOutputIntentFromPdfObject(object _cde.PdfObject) (*PdfOutputIntent, error) {
	_aebga := &PdfOutputIntent{}
	if _ebeef := _aebga.ParsePdfObject(object); _ebeef != nil {
		return nil, _ebeef
	}
	return _aebga, nil
}
func _gfdbd() string {
	_dccfe.Lock()
	defer _dccfe.Unlock()
	return _adcff
}

// NewXObjectImage returns a new XObjectImage.
func NewXObjectImage() *XObjectImage {
	_fafdc := &XObjectImage{}
	_cddcg := &_cde.PdfObjectStream{}
	_cddcg.PdfObjectDictionary = _cde.MakeDict()
	_fafdc._bbaed = _cddcg
	return _fafdc
}

// L returns the value of the L component of the color.
func (_afaf *PdfColorLab) L() float64 { return _afaf[0] }

// NewGrayImageFromGoImage creates a new grayscale unidoc Image from a golang Image.
func (_bgecba DefaultImageHandler) NewGrayImageFromGoImage(goimg _gf.Image) (*Image, error) {
	_ddced := goimg.Bounds()
	_fcbfe := &Image{Width: int64(_ddced.Dx()), Height: int64(_ddced.Dy()), ColorComponents: 1, BitsPerComponent: 8}
	switch _bbae := goimg.(type) {
	case *_gf.Gray:
		if len(_bbae.Pix) != _ddced.Dx()*_ddced.Dy() {
			_fcae, _cegfa := _ff.GrayConverter.Convert(goimg)
			if _cegfa != nil {
				return nil, _cegfa
			}
			_fcbfe.Data = _fcae.Pix()
		} else {
			_fcbfe.Data = _bbae.Pix
		}
	case *_gf.Gray16:
		_fcbfe.BitsPerComponent = 16
		if len(_bbae.Pix) != _ddced.Dx()*_ddced.Dy()*2 {
			_bbeeb, _cbgg := _ff.Gray16Converter.Convert(goimg)
			if _cbgg != nil {
				return nil, _cbgg
			}
			_fcbfe.Data = _bbeeb.Pix()
		} else {
			_fcbfe.Data = _bbae.Pix
		}
	case _ff.Image:
		_fgfb := _bbae.Base()
		if _fgfb.ColorComponents == 1 {
			_fcbfe.BitsPerComponent = int64(_fgfb.BitsPerComponent)
			_fcbfe.Data = _fgfb.Data
			return _fcbfe, nil
		}
		_agebb, _fcceg := _ff.GrayConverter.Convert(goimg)
		if _fcceg != nil {
			return nil, _fcceg
		}
		_fcbfe.Data = _agebb.Pix()
	default:
		_fbcgb, _fefag := _ff.GrayConverter.Convert(goimg)
		if _fefag != nil {
			return nil, _fefag
		}
		_fcbfe.Data = _fbcgb.Pix()
	}
	return _fcbfe, nil
}

// ToPdfObject implements interface PdfModel.
func (_ebfa *PdfAnnotationLine) ToPdfObject() _cde.PdfObject {
	_ebfa.PdfAnnotation.ToPdfObject()
	_egb := _ebfa._bddg
	_ecad := _egb.PdfObject.(*_cde.PdfObjectDictionary)
	_ebfa.PdfAnnotationMarkup.appendToPdfDictionary(_ecad)
	_ecad.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u004c\u0069\u006e\u0065"))
	_ecad.SetIfNotNil("\u004c", _ebfa.L)
	_ecad.SetIfNotNil("\u0042\u0053", _ebfa.BS)
	_ecad.SetIfNotNil("\u004c\u0045", _ebfa.LE)
	_ecad.SetIfNotNil("\u0049\u0043", _ebfa.IC)
	_ecad.SetIfNotNil("\u004c\u004c", _ebfa.LL)
	_ecad.SetIfNotNil("\u004c\u004c\u0045", _ebfa.LLE)
	_ecad.SetIfNotNil("\u0043\u0061\u0070", _ebfa.Cap)
	_ecad.SetIfNotNil("\u0049\u0054", _ebfa.IT)
	_ecad.SetIfNotNil("\u004c\u004c\u004f", _ebfa.LLO)
	_ecad.SetIfNotNil("\u0043\u0050", _ebfa.CP)
	_ecad.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _ebfa.Measure)
	_ecad.SetIfNotNil("\u0043\u004f", _ebfa.CO)
	return _egb
}
func (_bggc *DSS) addOCSPs(_ebeb [][]byte) ([]*_cde.PdfObjectStream, error) {
	return _bggc.add(&_bggc.OCSPs, _bggc._ecgd, _ebeb)
}

// GetContext returns a reference to the subshading entry as represented by PdfShadingType1-7.
func (_ggcdd *PdfShading) GetContext() PdfModel { return _ggcdd._dgfac }

// SetXObjectFormByName adds the provided XObjectForm to the page resources.
// The added XObjectForm is identified by the specified name.
func (_afaba *PdfPageResources) SetXObjectFormByName(keyName _cde.PdfObjectName, xform *XObjectForm) error {
	_gdbfa := xform.ToPdfObject().(*_cde.PdfObjectStream)
	_gffa := _afaba.SetXObjectByName(keyName, _gdbfa)
	return _gffa
}
func _befce() string {
	_dccfe.Lock()
	defer _dccfe.Unlock()
	return _bfcdae
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_dceb *PdfColorspaceDeviceGray) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bcag, _acgg := _cde.GetNumbersAsFloat(objects)
	if _acgg != nil {
		return nil, _acgg
	}
	return _dceb.ColorFromFloats(_bcag)
}

// PdfActionImportData represents a importData action.
type PdfActionImportData struct {
	*PdfAction
	F *PdfFilespec
}

func (_dbfde *PdfReader) newPdfAcroFormFromDict(_gaeda *_cde.PdfIndirectObject, _cggec *_cde.PdfObjectDictionary) (*PdfAcroForm, error) {
	_efca := NewPdfAcroForm()
	if _gaeda != nil {
		_efca._ecca = _gaeda
		_gaeda.PdfObject = _cde.MakeDict()
	}
	if _cdgd := _cggec.Get("\u0046\u0069\u0065\u006c\u0064\u0073"); _cdgd != nil && !_cde.IsNullObject(_cdgd) {
		_caeeg, _aggea := _cde.GetArray(_cdgd)
		if !_aggea {
			return nil, _ee.Errorf("\u0066i\u0065\u006c\u0064\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _cdgd)
		}
		var _facgc []*PdfField
		for _, _dgafe := range _caeeg.Elements() {
			_feacd, _gdba := _cde.GetIndirect(_dgafe)
			if !_gdba {
				if _, _bggag := _dgafe.(*_cde.PdfObjectNull); _bggag {
					_ad.Log.Trace("\u0053k\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072 \u006e\u0075\u006c\u006c\u0020\u0066\u0069\u0065\u006c\u0064")
					continue
				}
				_ad.Log.Debug("\u0046\u0069\u0065\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0064 \u0069\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _dgafe)
				return nil, _ee.Errorf("\u0066\u0069\u0065l\u0064\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_ccdcb, _fdgcd := _dbfde.newPdfFieldFromIndirectObject(_feacd, nil)
			if _fdgcd != nil {
				return nil, _fdgcd
			}
			_ad.Log.Trace("\u0041\u0063\u0072\u006fFo\u0072\u006d\u0020\u0046\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u002b\u0076", *_ccdcb)
			_facgc = append(_facgc, _ccdcb)
		}
		_efca.Fields = &_facgc
	}
	if _ecdc := _cggec.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _ecdc != nil {
		_bdcdg, _facf := _cde.GetBool(_ecdc)
		if _facf {
			_efca.NeedAppearances = _bdcdg
		} else {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u0065\u0065\u0064\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063e\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _ecdc)
		}
	}
	if _dcbf := _cggec.Get("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073"); _dcbf != nil {
		_feaee, _ffbbg := _cde.GetInt(_dcbf)
		if _ffbbg {
			_efca.SigFlags = _feaee
		} else {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073 \u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _dcbf)
		}
	}
	if _defe := _cggec.Get("\u0043\u004f"); _defe != nil {
		_bagg, _dcddf := _cde.GetArray(_defe)
		if _dcddf {
			_efca.CO = _bagg
		} else {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u004f\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _defe)
		}
	}
	if _edefd := _cggec.Get("\u0044\u0052"); _edefd != nil {
		if _cfga, _ffegg := _cde.GetDict(_edefd); _ffegg {
			_agbge, _ecbf := NewPdfPageResourcesFromDict(_cfga)
			if _ecbf != nil {
				_ad.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0044R\u003a\u0020\u0025\u0076", _ecbf)
				return nil, _ecbf
			}
			_efca.DR = _agbge
		} else {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0052\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _edefd)
		}
	}
	if _ecgbf := _cggec.Get("\u0044\u0041"); _ecgbf != nil {
		_dcdf, _ddad := _cde.GetString(_ecgbf)
		if _ddad {
			_efca.DA = _dcdf
		} else {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0041\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _ecgbf)
		}
	}
	if _cbebb := _cggec.Get("\u0051"); _cbebb != nil {
		_ebdga, _ffcd := _cde.GetInt(_cbebb)
		if _ffcd {
			_efca.Q = _ebdga
		} else {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0051\u0020\u0069\u006e\u0076a\u006ci\u0064 \u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _cbebb)
		}
	}
	if _bffg := _cggec.Get("\u0058\u0046\u0041"); _bffg != nil {
		_efca.XFA = _bffg
	}
	_efca.ToPdfObject()
	return _efca, nil
}

// StringToCharcodeBytes maps the provided string runes to charcode bytes and
// it returns the resulting slice of bytes, along with the number of runes
// which could not be converted. If the number of misses is 0, all string runes
// were successfully converted.
func (_adcaf *PdfFont) StringToCharcodeBytes(str string) ([]byte, int) {
	return _adcaf.RunesToCharcodeBytes([]rune(str))
}

// NewPdfPageResourcesFromDict creates and returns a new PdfPageResources object
// from the input dictionary.
func NewPdfPageResourcesFromDict(dict *_cde.PdfObjectDictionary) (*PdfPageResources, error) {
	_gdag := NewPdfPageResources()
	if _efcc := dict.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _efcc != nil {
		_gdag.ExtGState = _efcc
	}
	if _aagfd := dict.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"); _aagfd != nil && !_cde.IsNullObject(_aagfd) {
		_gdag.ColorSpace = _aagfd
	}
	if _becd := dict.Get("\u0050a\u0074\u0074\u0065\u0072\u006e"); _becd != nil {
		_gdag.Pattern = _becd
	}
	if _debg := dict.Get("\u0053h\u0061\u0064\u0069\u006e\u0067"); _debg != nil {
		_gdag.Shading = _debg
	}
	if _adgc := dict.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"); _adgc != nil {
		_gdag.XObject = _adgc
	}
	if _cgbdd := _cde.ResolveReference(dict.Get("\u0046\u006f\u006e\u0074")); _cgbdd != nil {
		_gdag.Font = _cgbdd
	}
	if _dbcd := dict.Get("\u0050r\u006f\u0063\u0053\u0065\u0074"); _dbcd != nil {
		_gdag.ProcSet = _dbcd
	}
	if _ceebc := dict.Get("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"); _ceebc != nil {
		_gdag.Properties = _ceebc
	}
	return _gdag, nil
}

// NewPdfAnnotation returns an initialized generic PDF annotation model.
func NewPdfAnnotation() *PdfAnnotation {
	_bbf := &PdfAnnotation{}
	_bbf._bddg = _cde.MakeIndirectObject(_cde.MakeDict())
	return _bbf
}
func _gdaeg(_eegbc _fe.StdFont) pdfFontSimple {
	_ccbbg := _eegbc.Descriptor()
	return pdfFontSimple{fontCommon: fontCommon{_dcbc: "\u0054\u0079\u0070e\u0031", _eeab: _eegbc.Name()}, _gggc: _eegbc.GetMetricsTable(), _ccff: &PdfFontDescriptor{FontName: _cde.MakeName(string(_ccbbg.Name)), FontFamily: _cde.MakeName(_ccbbg.Family), FontWeight: _cde.MakeFloat(float64(_ccbbg.Weight)), Flags: _cde.MakeInteger(int64(_ccbbg.Flags)), FontBBox: _cde.MakeArrayFromFloats(_ccbbg.BBox[:]), ItalicAngle: _cde.MakeFloat(_ccbbg.ItalicAngle), Ascent: _cde.MakeFloat(_ccbbg.Ascent), Descent: _cde.MakeFloat(_ccbbg.Descent), CapHeight: _cde.MakeFloat(_ccbbg.CapHeight), XHeight: _cde.MakeFloat(_ccbbg.XHeight), StemV: _cde.MakeFloat(_ccbbg.StemV), StemH: _cde.MakeFloat(_ccbbg.StemH)}, _facga: _eegbc.Encoder()}
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the A, B and C components of the color.
func (_cfge *PdfColorspaceCalRGB) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_agga, _cbeff := _cde.GetNumbersAsFloat(objects)
	if _cbeff != nil {
		return nil, _cbeff
	}
	return _cfge.ColorFromFloats(_agga)
}

// PdfAnnotationSquiggly represents Squiggly annotations.
// (Section 12.5.6.10).
type PdfAnnotationSquiggly struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _cde.PdfObject
}

func _cefb(_cbebbe *_cde.PdfObjectStream) (*PdfFunctionType4, error) {
	_dgfcgb := &PdfFunctionType4{}
	_dgfcgb._eaded = _cbebbe
	_eacaf := _cbebbe.PdfObjectDictionary
	_aggb, _dedga := _cde.TraceToDirectObject(_eacaf.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_cde.PdfObjectArray)
	if !_dedga {
		_ad.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _aggb.Len()%2 != 0 {
		_ad.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _ceg.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_baca, _ggbbe := _aggb.ToFloat64Array()
	if _ggbbe != nil {
		return nil, _ggbbe
	}
	_dgfcgb.Domain = _baca
	_aggb, _dedga = _cde.TraceToDirectObject(_eacaf.Get("\u0052\u0061\u006eg\u0065")).(*_cde.PdfObjectArray)
	if _dedga {
		if _aggb.Len() < 0 || _aggb.Len()%2 != 0 {
			return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_aggae, _eaebc := _aggb.ToFloat64Array()
		if _eaebc != nil {
			return nil, _eaebc
		}
		_dgfcgb.Range = _aggae
	}
	_bdbgg, _ggbbe := _cde.DecodeStream(_cbebbe)
	if _ggbbe != nil {
		return nil, _ggbbe
	}
	_dgfcgb._bfbd = _bdbgg
	_ggfee := _dg.NewPSParser(_bdbgg)
	_ebdf, _ggbbe := _ggfee.Parse()
	if _ggbbe != nil {
		return nil, _ggbbe
	}
	_dgfcgb.Program = _ebdf
	return _dgfcgb, nil
}

// XObjectType represents the type of an XObject.
type XObjectType int

// NewPdfAnnotationTrapNet returns a new trapnet annotation.
func NewPdfAnnotationTrapNet() *PdfAnnotationTrapNet {
	_dgae := NewPdfAnnotation()
	_gaba := &PdfAnnotationTrapNet{}
	_gaba.PdfAnnotation = _dgae
	_dgae.SetContext(_gaba)
	return _gaba
}

// GetXObjectFormByName returns the XObjectForm with the specified name from the
// page resources, if it exists.
func (_cdfag *PdfPageResources) GetXObjectFormByName(keyName _cde.PdfObjectName) (*XObjectForm, error) {
	_gfcca, _afeb := _cdfag.GetXObjectByName(keyName)
	if _gfcca == nil {
		return nil, nil
	}
	if _afeb != XObjectTypeForm {
		return nil, _ceg.New("\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006f\u0072\u006d")
	}
	_eaea, _afebg := NewXObjectFormFromStream(_gfcca)
	if _afebg != nil {
		return nil, _afebg
	}
	return _eaea, nil
}

// ToUnicode returns the name of the font's "ToUnicode" field if there is one, or "" if there isn't.
func (_dcfge *PdfFont) ToUnicode() string {
	if _dcfge.baseFields()._ggebg == nil {
		return ""
	}
	return _dcfge.baseFields()._ggebg.Name()
}

// SetReason sets the `Reason` field of the signature.
func (_agced *PdfSignature) SetReason(reason string) { _agced.Reason = _cde.MakeString(reason) }

// GetRevisionNumber returns the version of the current Pdf document
func (_dabce *PdfReader) GetRevisionNumber() int { return _dabce._aggcgb.GetRevisionNumber() }

// ToPdfObject returns the choice field dictionary within an indirect object (container).
func (_eefa *PdfFieldChoice) ToPdfObject() _cde.PdfObject {
	_eefa.PdfField.ToPdfObject()
	_egfg := _eefa._afgc
	_cdea := _egfg.PdfObject.(*_cde.PdfObjectDictionary)
	_cdea.Set("\u0046\u0054", _cde.MakeName("\u0043\u0068"))
	if _eefa.Opt != nil {
		_cdea.Set("\u004f\u0070\u0074", _eefa.Opt)
	}
	if _eefa.TI != nil {
		_cdea.Set("\u0054\u0049", _eefa.TI)
	}
	if _eefa.I != nil {
		_cdea.Set("\u0049", _eefa.I)
	}
	return _egfg
}

// WriteString outputs the object as it is to be written to file.
func (_agedf *pdfSignDictionary) WriteString() string {
	_agedf._gggca = 0
	_agedf._cgedg = 0
	_agedf._faba = 0
	_agedf._edgb = 0
	_gcbfg := _ede.NewBuffer(nil)
	_gcbfg.WriteString("\u003c\u003c")
	for _, _dagfg := range _agedf.Keys() {
		_gaded := _agedf.Get(_dagfg)
		switch _dagfg {
		case "\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e":
			_gcbfg.WriteString(_dagfg.WriteString())
			_gcbfg.WriteString("\u0020")
			_agedf._faba = _gcbfg.Len()
			_gcbfg.WriteString(_gaded.WriteString())
			_gcbfg.WriteString("\u0020")
			_agedf._edgb = _gcbfg.Len() - 1
		case "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073":
			_gcbfg.WriteString(_dagfg.WriteString())
			_gcbfg.WriteString("\u0020")
			_agedf._gggca = _gcbfg.Len()
			_gcbfg.WriteString(_gaded.WriteString())
			_gcbfg.WriteString("\u0020")
			_agedf._cgedg = _gcbfg.Len() - 1
		default:
			_gcbfg.WriteString(_dagfg.WriteString())
			_gcbfg.WriteString("\u0020")
			_gcbfg.WriteString(_gaded.WriteString())
		}
	}
	_gcbfg.WriteString("\u003e\u003e")
	return _gcbfg.String()
}

// ColorToRGB only converts color used with uncolored patterns (defined in underlying colorspace).  Does not go into the
// pattern objects and convert those.  If that is desired, needs to be done separately.  See for example
// grayscale conversion example in unidoc-examples repo.
func (_gfbg *PdfColorspaceSpecialPattern) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ebed, _bcggf := color.(*PdfColorPattern)
	if !_bcggf {
		_ad.Log.Debug("\u0043\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0070a\u0074\u0074\u0065\u0072\u006e\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", color)
		return nil, ErrTypeCheck
	}
	if _ebed.Color == nil {
		return color, nil
	}
	if _gfbg.UnderlyingCS == nil {
		return nil, _ceg.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	return _gfbg.UnderlyingCS.ColorToRGB(_ebed.Color)
}

// ToPdfObject implements interface PdfModel.
func (_fcdb *PdfAnnotationRichMedia) ToPdfObject() _cde.PdfObject {
	_fcdb.PdfAnnotation.ToPdfObject()
	_cea := _fcdb._bddg
	_daeg := _cea.PdfObject.(*_cde.PdfObjectDictionary)
	_daeg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a"))
	_daeg.SetIfNotNil("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073", _fcdb.RichMediaSettings)
	_daeg.SetIfNotNil("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074", _fcdb.RichMediaContent)
	return _cea
}

// PdfActionJavaScript represents a javaScript action.
type PdfActionJavaScript struct {
	*PdfAction
	JS _cde.PdfObject
}

// Field returns the parent form field of the widget annotation, if one exists.
// NOTE: the method returns nil if the parent form field has not been parsed.
func (_bcd *PdfAnnotationWidget) Field() *PdfField { return _bcd._dbf }
func _cfbaf(_bega *fontCommon) *pdfFontType0       { return &pdfFontType0{fontCommon: *_bega} }

// NewBorderStyle returns an initialized PdfBorderStyle.
func NewBorderStyle() *PdfBorderStyle { _egfa := &PdfBorderStyle{}; return _egfa }

// ParsePdfObject parses input pdf object into given output intent.
func (_befa *PdfOutputIntent) ParsePdfObject(object _cde.PdfObject) error {
	_egfebf, _gcbff := _cde.GetDict(object)
	if !_gcbff {
		_ad.Log.Error("\u0055\u006e\u006bno\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020%\u0054 \u0066o\u0072 \u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074", object)
		return _ceg.New("\u0075\u006e\u006b\u006e\u006fw\u006e\u0020\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074")
	}
	_befa._acbcb = _egfebf
	_befa.Type, _ = _egfebf.GetString("\u0054\u0079\u0070\u0065")
	_efcda, _gcbff := _egfebf.GetString("\u0053")
	if _gcbff {
		switch _efcda {
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411":
			_befa.S = PdfOutputIntentTypeA1
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00412":
			_befa.S = PdfOutputIntentTypeA2
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00413":
			_befa.S = PdfOutputIntentTypeA3
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00414":
			_befa.S = PdfOutputIntentTypeA4
		case "\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0058":
			_befa.S = PdfOutputIntentTypeX
		}
	}
	_befa.OutputCondition, _ = _egfebf.GetString("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e")
	_befa.OutputConditionIdentifier, _ = _egfebf.GetString("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r")
	_befa.RegistryName, _ = _egfebf.GetString("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065")
	_befa.Info, _ = _egfebf.GetString("\u0049\u006e\u0066\u006f")
	if _begbe, _cggf := _cde.GetStream(_egfebf.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _cggf {
		_befa.ColorComponents, _ = _cde.GetIntVal(_begbe.Get("\u004e"))
		_gbee, _gecge := _cde.DecodeStream(_begbe)
		if _gecge != nil {
			return _gecge
		}
		_befa.DestOutputProfile = _gbee
	}
	return nil
}
func _bcfa(_afbda *_cde.PdfObjectDictionary) (*PdfShadingType6, error) {
	_dgbce := PdfShadingType6{}
	_cbdbf := _afbda.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _cbdbf == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fcebe, _gdfffc := _cbdbf.(*_cde.PdfObjectInteger)
	if !_gdfffc {
		_ad.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cbdbf)
		return nil, _cde.ErrTypeError
	}
	_dgbce.BitsPerCoordinate = _fcebe
	_cbdbf = _afbda.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _cbdbf == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_fcebe, _gdfffc = _cbdbf.(*_cde.PdfObjectInteger)
	if !_gdfffc {
		_ad.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _cbdbf)
		return nil, _cde.ErrTypeError
	}
	_dgbce.BitsPerComponent = _fcebe
	_cbdbf = _afbda.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _cbdbf == nil {
		_ad.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fcebe, _gdfffc = _cbdbf.(*_cde.PdfObjectInteger)
	if !_gdfffc {
		_ad.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _cbdbf)
		return nil, _cde.ErrTypeError
	}
	_dgbce.BitsPerComponent = _fcebe
	_cbdbf = _afbda.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _cbdbf == nil {
		_ad.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_gdccf, _gdfffc := _cbdbf.(*_cde.PdfObjectArray)
	if !_gdfffc {
		_ad.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cbdbf)
		return nil, _cde.ErrTypeError
	}
	_dgbce.Decode = _gdccf
	if _dccg := _afbda.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _dccg != nil {
		_dgbce.Function = []PdfFunction{}
		if _gccgd, _ccdfg := _dccg.(*_cde.PdfObjectArray); _ccdfg {
			for _, _cecf := range _gccgd.Elements() {
				_fggfb, _bgbfc := _cfdbb(_cecf)
				if _bgbfc != nil {
					_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bgbfc)
					return nil, _bgbfc
				}
				_dgbce.Function = append(_dgbce.Function, _fggfb)
			}
		} else {
			_bdcee, _eabgad := _cfdbb(_dccg)
			if _eabgad != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _eabgad)
				return nil, _eabgad
			}
			_dgbce.Function = append(_dgbce.Function, _bdcee)
		}
	}
	return &_dgbce, nil
}

// PdfFilespec represents a file specification which can either refer to an external or embedded file.
type PdfFilespec struct {
	Type  _cde.PdfObject
	FS    _cde.PdfObject
	F     _cde.PdfObject
	UF    _cde.PdfObject
	DOS   _cde.PdfObject
	Mac   _cde.PdfObject
	Unix  _cde.PdfObject
	ID    _cde.PdfObject
	V     _cde.PdfObject
	EF    _cde.PdfObject
	RF    _cde.PdfObject
	Desc  _cde.PdfObject
	CI    _cde.PdfObject
	_bgac _cde.PdfObject
}

func _dbgae(_gfecge []*_cde.PdfObjectStream) *_cde.PdfObjectArray {
	if len(_gfecge) == 0 {
		return nil
	}
	_gcaccf := make([]_cde.PdfObject, 0, len(_gfecge))
	for _, _aagge := range _gfecge {
		_gcaccf = append(_gcaccf, _aagge)
	}
	return _cde.MakeArray(_gcaccf...)
}
func (_ecc *PdfReader) newPdfActionURIFromDict(_aaa *_cde.PdfObjectDictionary) (*PdfActionURI, error) {
	return &PdfActionURI{URI: _aaa.Get("\u0055\u0052\u0049"), IsMap: _aaa.Get("\u0049\u0073\u004da\u0070")}, nil
}

// EnableAll LTV enables all signatures in the PDF document.
// The signing certificate chain is extracted from each signature dictionary.
// Optionally, additional certificates can be specified through the
// `extraCerts` parameter. The LTV client attempts to build the certificate
// chain up to a trusted root by downloading any missing certificates.
func (_gded *LTV) EnableAll(extraCerts []*_bg.Certificate) error {
	_bbba := _gded._ffab._cfad.AcroForm
	for _, _ddafd := range _bbba.AllFields() {
		_cbcgga, _ := _ddafd.GetContext().(*PdfFieldSignature)
		if _cbcgga == nil {
			continue
		}
		_fcede := _cbcgga.V
		if _gedf := _gded.validateSig(_fcede); _gedf != nil {
			_ad.Log.Debug("\u0057\u0041\u0052N\u003a\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _gedf)
		}
		if _gfagcf := _gded.Enable(_fcede, extraCerts); _gfagcf != nil {
			return _gfagcf
		}
	}
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_agb *PdfActionSubmitForm) ToPdfObject() _cde.PdfObject {
	_agb.PdfAction.ToPdfObject()
	_fd := _agb._bc
	_gbb := _fd.PdfObject.(*_cde.PdfObjectDictionary)
	_gbb.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeSubmitForm)))
	if _agb.F != nil {
		_gbb.Set("\u0046", _agb.F.ToPdfObject())
	}
	_gbb.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _agb.Fields)
	_gbb.SetIfNotNil("\u0046\u006c\u0061g\u0073", _agb.Flags)
	return _fd
}
func (_dbb *PdfReader) newPdfActionResetFormFromDict(_dfcd *_cde.PdfObjectDictionary) (*PdfActionResetForm, error) {
	return &PdfActionResetForm{Fields: _dfcd.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _dfcd.Get("\u0046\u006c\u0061g\u0073")}, nil
}

// GetObjectNums returns the object numbers of the PDF objects in the file
// Numbered objects are either indirect objects or stream objects.
// e.g. objNums := pdfReader.GetObjectNums()
// The underlying objects can then be accessed with
// pdfReader.GetIndirectObjectByNumber(objNums[0]) for the first available object.
func (_fedcd *PdfReader) GetObjectNums() []int { return _fedcd._aggcgb.GetObjectNums() }
func (_dgga *PdfReader) newPdfActionImportDataFromDict(_gfe *_cde.PdfObjectDictionary) (*PdfActionImportData, error) {
	_agbe, _aee := _beed(_gfe.Get("\u0046"))
	if _aee != nil {
		return nil, _aee
	}
	return &PdfActionImportData{F: _agbe}, nil
}

const (
	_ PdfOutputIntentType = iota
	PdfOutputIntentTypeA1
	PdfOutputIntentTypeA2
	PdfOutputIntentTypeA3
	PdfOutputIntentTypeA4
	PdfOutputIntentTypeX
)

// ImageToRGB converts Lab colorspace image to RGB and returns the result.
func (_bged *PdfColorspaceLab) ImageToRGB(img Image) (Image, error) {
	_fded := func(_bagf float64) float64 {
		if _bagf >= 6.0/29 {
			return _bagf * _bagf * _bagf
		}
		return 108.0 / 841 * (_bagf - 4/29)
	}
	_fcbb := img._aaafb
	if len(_fcbb) != 6 {
		_ad.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u002d\u0020\u004c\u0061\u0062\u0020\u0044e\u0063\u006f\u0064\u0065\u0020\u0072\u0061\u006e\u0067e\u0020\u0021\u003d\u0020\u0036\u002e\u002e\u002e\u0020\u0075\u0073\u0065\u0020\u005b0\u0020\u0031\u0030\u0030\u0020\u0061\u006d\u0069\u006e\u0020\u0061\u006d\u0061\u0078\u0020\u0062\u006d\u0069\u006e\u0020\u0062\u006d\u0061\u0078\u005d\u0020\u0064\u0065\u0066\u0061u\u006c\u0074\u0020\u0064\u0065\u0063\u006f\u0064\u0065 \u0061\u0072r\u0061\u0079")
		_fcbb = _bged.DecodeArray()
	}
	_egbb := _cae.NewReader(img.getBase())
	_gfgg := _ff.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, img._deegf, img._aaafb)
	_dddc := _cae.NewWriter(_gfgg)
	_cefg := _ced.Pow(2, float64(img.BitsPerComponent)) - 1
	_efbe := make([]uint32, 3)
	var (
		_fgcb                                              error
		Ls, As, Bs, L, M, N, X, Y, Z, _efgd, _ebca, _ggbcb float64
	)
	for {
		_fgcb = _egbb.ReadSamples(_efbe)
		if _fgcb == _f.EOF {
			break
		} else if _fgcb != nil {
			return img, _fgcb
		}
		Ls = float64(_efbe[0]) / _cefg
		As = float64(_efbe[1]) / _cefg
		Bs = float64(_efbe[2]) / _cefg
		Ls = _ff.LinearInterpolate(Ls, 0.0, 1.0, _fcbb[0], _fcbb[1])
		As = _ff.LinearInterpolate(As, 0.0, 1.0, _fcbb[2], _fcbb[3])
		Bs = _ff.LinearInterpolate(Bs, 0.0, 1.0, _fcbb[4], _fcbb[5])
		L = (Ls+16)/116 + As/500
		M = (Ls + 16) / 116
		N = (Ls+16)/116 - Bs/200
		X = _bged.WhitePoint[0] * _fded(L)
		Y = _bged.WhitePoint[1] * _fded(M)
		Z = _bged.WhitePoint[2] * _fded(N)
		_efgd = 3.240479*X + -1.537150*Y + -0.498535*Z
		_ebca = -0.969256*X + 1.875992*Y + 0.041556*Z
		_ggbcb = 0.055648*X + -0.204043*Y + 1.057311*Z
		_efgd = _ced.Min(_ced.Max(_efgd, 0), 1.0)
		_ebca = _ced.Min(_ced.Max(_ebca, 0), 1.0)
		_ggbcb = _ced.Min(_ced.Max(_ggbcb, 0), 1.0)
		_efbe[0] = uint32(_efgd * _cefg)
		_efbe[1] = uint32(_ebca * _cefg)
		_efbe[2] = uint32(_ggbcb * _cefg)
		if _fgcb = _dddc.WriteSamples(_efbe); _fgcb != nil {
			return img, _fgcb
		}
	}
	return _bddb(&_gfgg), nil
}

// NewReaderForText makes a new PdfReader for an input PDF content string. For use in testing.
func NewReaderForText(txt string) *PdfReader {
	return &PdfReader{_efbdd: map[_cde.PdfObject]struct{}{}, _bedfa: _acaga(), _aggcgb: _cde.NewParserFromString(txt)}
}

// GetNumComponents returns the number of color components (1 for Indexed).
func (_cbbb *PdfColorspaceSpecialIndexed) GetNumComponents() int { return 1 }

// ToPdfObject implements interface PdfModel.
func (_dfb *PdfAnnotationLink) ToPdfObject() _cde.PdfObject {
	_dfb.PdfAnnotation.ToPdfObject()
	_bab := _dfb._bddg
	_dcdd := _bab.PdfObject.(*_cde.PdfObjectDictionary)
	_dcdd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u004c\u0069\u006e\u006b"))
	if _dfb._beg != nil && _dfb._beg._bgd != nil {
		_dcdd.Set("\u0041", _dfb._beg._bgd.ToPdfObject())
	} else if _dfb.A != nil {
		_dcdd.Set("\u0041", _dfb.A)
	}
	_dcdd.SetIfNotNil("\u0044\u0065\u0073\u0074", _dfb.Dest)
	_dcdd.SetIfNotNil("\u0048", _dfb.H)
	_dcdd.SetIfNotNil("\u0050\u0041", _dfb.PA)
	_dcdd.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _dfb.QuadPoints)
	_dcdd.SetIfNotNil("\u0042\u0053", _dfb.BS)
	return _bab
}

// FillWithAppearance populates `form` with values provided by `provider`.
// If not nil, `appGen` is used to generate appearance dictionaries for the
// field annotations, based on the specified settings. Otherwise, appearance
// generation is skipped.
// e.g.: appGen := annotator.FieldAppearance{OnlyIfMissing: true, RegenerateTextFields: true}
// NOTE: In next major version this functionality will be part of Fill. (v4)
func (_aafce *PdfAcroForm) FillWithAppearance(provider FieldValueProvider, appGen FieldAppearanceGenerator) error {
	_ecdd := _aafce.fill(provider, appGen)
	if _ecdd != nil {
		return _ecdd
	}
	if _, _bfcab := provider.(FieldImageProvider); _bfcab {
		_ecdd = _aafce.fillImageWithAppearance(provider.(FieldImageProvider), appGen)
	}
	return _ecdd
}

// GetOCProperties returns the optional content properties PdfObject.
func (_gfbd *PdfReader) GetOCProperties() (_cde.PdfObject, error) {
	_egbce := _gfbd._efabe
	_efacf := _egbce.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	_efacf = _cde.ResolveReference(_efacf)
	if !_gfbd._cdgee {
		_agadg := _gfbd.traverseObjectData(_efacf)
		if _agadg != nil {
			return nil, _agadg
		}
	}
	return _efacf, nil
}

// GetDescent returns the Descent of the font `descriptor`.
func (_dafdd *PdfFontDescriptor) GetDescent() (float64, error) {
	return _cde.GetNumberAsFloat(_dafdd.Descent)
}

// PdfColorspaceSpecialIndexed is an indexed color space is a lookup table, where the input element
// is an index to the lookup table and the output is a color defined in the lookup table in the Base
// colorspace.
// [/Indexed base hival lookup]
type PdfColorspaceSpecialIndexed struct {
	Base   PdfColorspace
	HiVal  int
	Lookup _cde.PdfObject
	_dafd  []byte
	_gcbcb *_cde.PdfIndirectObject
}

// CheckAccessRights checks access rights and permissions for a specified password.  If either user/owner
// password is specified,  full rights are granted, otherwise the access rights are specified by the
// Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_fgfcc *PdfReader) CheckAccessRights(password []byte) (bool, _ccg.Permissions, error) {
	return _fgfcc._aggcgb.CheckAccessRights(password)
}

// PdfFunction interface represents the common methods of a function in PDF.
type PdfFunction interface {
	Evaluate([]float64) ([]float64, error)
	ToPdfObject() _cde.PdfObject
}

// SetNamedDestinations sets the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_ccfb *PdfWriter) SetNamedDestinations(dests _cde.PdfObject) error {
	if dests == nil {
		return nil
	}
	_ad.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0044\u0065\u0073\u0074\u0073\u002e\u002e\u002e")
	_ccfb._fedbb.Set("\u0044\u0065\u0073t\u0073", dests)
	return _ccfb.addObjects(dests)
}

// GetPerms returns the Permissions dictionary
func (_ccbed *PdfReader) GetPerms() *Permissions { return _ccbed._ebbdb }
func (_bebe *PdfReader) loadStructure() error {
	if _bebe._aggcgb.GetCrypter() != nil && !_bebe._aggcgb.IsAuthenticated() {
		return _ee.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_dddfa := _bebe._aggcgb.GetTrailer()
	if _dddfa == nil {
		return _ee.Errorf("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_aefdb, _dgda := _dddfa.Get("\u0052\u006f\u006f\u0074").(*_cde.PdfObjectReference)
	if !_dgda {
		return _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052\u006f\u006ft\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a \u0025\u0073\u0029", _dddfa)
	}
	_cdga, _bafc := _bebe._aggcgb.LookupByReference(*_aefdb)
	if _bafc != nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0072\u006f\u006f\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _bafc)
		return _bafc
	}
	_fdfe, _dgda := _cdga.(*_cde.PdfIndirectObject)
	if !_dgda {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _cdga, *_dddfa)
		return _ceg.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_cbbcb, _dgda := (*_fdfe).PdfObject.(*_cde.PdfObjectDictionary)
	if !_dgda {
		_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0061t\u0061\u006c\u006fg\u0020(\u0025\u0073\u0029", _fdfe.PdfObject)
		return _ceg.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_ad.Log.Trace("C\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _cbbcb)
	_ggbea, _dgda := _cbbcb.Get("\u0050\u0061\u0067e\u0073").(*_cde.PdfObjectReference)
	if !_dgda {
		return _ceg.New("\u0070\u0061\u0067\u0065\u0073\u0020\u0069\u006e\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0061\u0020\u0072e\u0066\u0065r\u0065\u006e\u0063\u0065")
	}
	_bbgcf, _bafc := _bebe._aggcgb.LookupByReference(*_ggbea)
	if _bafc != nil {
		_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020F\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020r\u0065\u0061\u0064 \u0070a\u0067\u0065\u0073")
		return _bafc
	}
	_bbbcb, _dgda := _bbgcf.(*_cde.PdfIndirectObject)
	if !_dgda {
		_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020P\u0061\u0067\u0065\u0073\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0069n\u0076a\u006c\u0069\u0064")
		_ad.Log.Debug("\u006f\u0070\u003a\u0020\u0025\u0070", _bbbcb)
		return _ceg.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_ccfg, _dgda := _bbbcb.PdfObject.(*_cde.PdfObjectDictionary)
	if !_dgda {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065c\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0073\u0029", _bbbcb)
		return _ceg.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_aefcf, _dgda := _cde.GetInt(_ccfg.Get("\u0043\u006f\u0075n\u0074"))
	if !_dgda {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0050\u0061\u0067\u0065\u0073\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return _ceg.New("\u0070\u0061\u0067\u0065s \u0063\u006f\u0075\u006e\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _, _dgda = _cde.GetName(_ccfg.Get("\u0054\u0079\u0070\u0065")); !_dgda {
		_ad.Log.Debug("\u0050\u0061\u0067\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020T\u0079\u0070\u0065\u0020\u0066\u0069\u0065\u006cd\u0020n\u006f\u0074\u0020\u0073\u0065\u0074\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0054\u0079p\u0065\u0020\u0074\u006f\u0020\u0050\u0061\u0067\u0065\u0073\u002e")
		_ccfg.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0050\u0061\u0067e\u0073"))
	}
	if _dbag, _ebef := _cde.GetInt(_ccfg.Get("\u0052\u006f\u0074\u0061\u0074\u0065")); _ebef {
		_efaad := int64(*_dbag)
		_bebe.Rotate = &_efaad
	}
	_bebe._bgbage = _aefdb
	_bebe._efabe = _cbbcb
	_bebe._bdagc = _ccfg
	_bebe._geadg = _bbbcb
	_bebe._bdaff = int(*_aefcf)
	_bebe._gbfgg = []*_cde.PdfIndirectObject{}
	_abgad := map[_cde.PdfObject]struct{}{}
	_bafc = _bebe.buildPageList(_bbbcb, nil, _abgad)
	if _bafc != nil {
		return _bafc
	}
	_ad.Log.Trace("\u002d\u002d\u002d")
	_ad.Log.Trace("\u0054\u004f\u0043")
	_ad.Log.Trace("\u0050\u0061\u0067e\u0073")
	_ad.Log.Trace("\u0025\u0064\u003a\u0020\u0025\u0073", len(_bebe._gbfgg), _bebe._gbfgg)
	_bebe._aegbb, _bafc = _bebe.loadOutlines()
	if _bafc != nil {
		_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075i\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 t\u0072\u0065\u0065 \u0028%\u0073\u0029", _bafc)
		return _bafc
	}
	_bebe.AcroForm, _bafc = _bebe.loadForms()
	if _bafc != nil {
		return _bafc
	}
	_bebe.DSS, _bafc = _bebe.loadDSS()
	if _bafc != nil {
		return _bafc
	}
	_bebe._ebbdb, _bafc = _bebe.loadPerms()
	if _bafc != nil {
		return _bafc
	}
	return nil
}

// GetPageLabels returns the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_ceeec *PdfReader) GetPageLabels() (_cde.PdfObject, error) {
	_defec := _cde.ResolveReference(_ceeec._efabe.Get("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073"))
	if _defec == nil {
		return nil, nil
	}
	if !_ceeec._cdgee {
		_cacca := _ceeec.traverseObjectData(_defec)
		if _cacca != nil {
			return nil, _cacca
		}
	}
	return _defec, nil
}

// ButtonType represents the subtype of a button field, can be one of:
// - Checkbox (ButtonTypeCheckbox)
// - PushButton (ButtonTypePushButton)
// - RadioButton (ButtonTypeRadioButton)
type ButtonType int

func _cgcfa() string { _dccfe.Lock(); defer _dccfe.Unlock(); return _dgcdb }

// PdfModel is a higher level PDF construct which can be collapsed into a PdfObject.
// Each PdfModel has an underlying PdfObject and vice versa (one-to-one).
// Under normal circumstances there should only be one copy of each.
// Copies can be made, but care must be taken to do it properly.
type PdfModel interface {
	ToPdfObject() _cde.PdfObject
	GetContainingPdfObject() _cde.PdfObject
}

func (_eecf *PdfReader) loadAnnotations(_eefb _cde.PdfObject) ([]*PdfAnnotation, error) {
	_fbbb, _gcecd := _cde.GetArray(_eefb)
	if !_gcecd {
		return nil, _ee.Errorf("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	var _caac []*PdfAnnotation
	for _, _bbfff := range _fbbb.Elements() {
		_bbfff = _cde.ResolveReference(_bbfff)
		if _, _fefeef := _bbfff.(*_cde.PdfObjectNull); _fefeef {
			continue
		}
		_cfcdc, _egab := _bbfff.(*_cde.PdfObjectDictionary)
		_dadbf, _fadbc := _bbfff.(*_cde.PdfIndirectObject)
		if _egab {
			_dadbf = &_cde.PdfIndirectObject{}
			_dadbf.PdfObject = _cfcdc
		} else {
			if !_fadbc {
				return nil, _ee.Errorf("\u0061\u006eno\u0074\u0061\u0074i\u006f\u006e\u0020\u006eot \u0069n \u0061\u006e\u0020\u0069\u006e\u0064\u0069re\u0063\u0074\u0020\u006f\u0062\u006a\u0065c\u0074")
			}
		}
		_bgfg, _cfdbe := _eecf.newPdfAnnotationFromIndirectObject(_dadbf)
		if _cfdbe != nil {
			return nil, _cfdbe
		}
		switch _gaafa := _bgfg.GetContext().(type) {
		case *PdfAnnotationWidget:
			for _, _gbabc := range _eecf.AcroForm.AllFields() {
				if _gbabc._afgc == _gaafa.Parent {
					_gaafa._dbf = _gbabc
					break
				}
			}
		}
		if _bgfg != nil {
			_caac = append(_caac, _bgfg)
		}
	}
	return _caac, nil
}

// ToPdfObject implements interface PdfModel.
// Note: Call the sub-annotation's ToPdfObject to set both the generic and non-generic information.
func (_dec *PdfAnnotation) ToPdfObject() _cde.PdfObject {
	_ece := _dec._bddg
	_gdbf := _ece.PdfObject.(*_cde.PdfObjectDictionary)
	_gdbf.Clear()
	_gdbf.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0041\u006e\u006eo\u0074"))
	_gdbf.SetIfNotNil("\u0052\u0065\u0063\u0074", _dec.Rect)
	_gdbf.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _dec.Contents)
	_gdbf.SetIfNotNil("\u0050", _dec.P)
	_gdbf.SetIfNotNil("\u004e\u004d", _dec.NM)
	_gdbf.SetIfNotNil("\u004d", _dec.M)
	_gdbf.SetIfNotNil("\u0046", _dec.F)
	_gdbf.SetIfNotNil("\u0041\u0050", _dec.AP)
	_gdbf.SetIfNotNil("\u0041\u0053", _dec.AS)
	_gdbf.SetIfNotNil("\u0042\u006f\u0072\u0064\u0065\u0072", _dec.Border)
	_gdbf.SetIfNotNil("\u0043", _dec.C)
	_gdbf.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _dec.StructParent)
	_gdbf.SetIfNotNil("\u004f\u0043", _dec.OC)
	return _ece
}
func (_dbbde *PdfReader) traverseObjectData(_gfefe _cde.PdfObject) error {
	return _cde.ResolveReferencesDeep(_gfefe, _dbbde._efbdd)
}

var _egee = map[string]struct{}{"\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074\u0045\u006e\u0063o\u0064\u0069\u006e\u0067": {}, "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}}

// EnableByName LTV enables the signature dictionary of the PDF AcroForm
// field identified the specified name. The signing certificate chain is
// extracted from the signature dictionary. Optionally, additional certificates
// can be specified through the `extraCerts` parameter. The LTV client attempts
// to build the certificate chain up to a trusted root by downloading any
// missing certificates.
func (_bdeec *LTV) EnableByName(name string, extraCerts []*_bg.Certificate) error {
	_afgf := _bdeec._ffab._cfad.AcroForm
	for _, _dgdc := range _afgf.AllFields() {
		_agbaa, _ := _dgdc.GetContext().(*PdfFieldSignature)
		if _agbaa == nil {
			continue
		}
		if _gebbd := _agbaa.PartialName(); _gebbd != name {
			continue
		}
		return _bdeec.Enable(_agbaa.V, extraCerts)
	}
	return nil
}

// GetXObjectByName gets XObject by name.
func (_bcfgd *PdfPage) GetXObjectByName(name _cde.PdfObjectName) (_cde.PdfObject, bool) {
	_afgcg, _ebcca := _bcfgd.Resources.XObject.(*_cde.PdfObjectDictionary)
	if !_ebcca {
		return nil, false
	}
	if _dbgf := _afgcg.Get(name); _dbgf != nil {
		return _dbgf, true
	}
	return nil, false
}

// PdfOutlineTreeNode contains common fields used by the outline and outline
// item objects.
type PdfOutlineTreeNode struct {
	_fbeea interface{}
	First  *PdfOutlineTreeNode
	Last   *PdfOutlineTreeNode
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_ffggc pdfFontType0) GetRuneMetrics(r rune) (_fe.CharMetrics, bool) {
	if _ffggc.DescendantFont == nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _ffggc)
		return _fe.CharMetrics{}, false
	}
	return _ffggc.DescendantFont.GetRuneMetrics(r)
}
func (_fdfd PdfFont) actualFont() pdfFont {
	if _fdfd._gbcff == nil {
		_ad.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0061\u0063\u0074\u0075\u0061\u006c\u0046\u006f\u006e\u0074\u002e\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c.\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _fdfd)
	}
	return _fdfd._gbcff
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// component PDF objects.
func (_cegg *PdfColorspaceICCBased) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if _cegg.Alternate == nil {
		if _cegg.N == 1 {
			_eace := NewPdfColorspaceDeviceGray()
			return _eace.ColorFromPdfObjects(objects)
		} else if _cegg.N == 3 {
			_daed := NewPdfColorspaceDeviceRGB()
			return _daed.ColorFromPdfObjects(objects)
		} else if _cegg.N == 4 {
			_gcac := NewPdfColorspaceDeviceCMYK()
			return _gcac.ColorFromPdfObjects(objects)
		} else {
			return nil, _ceg.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _cegg.Alternate.ColorFromPdfObjects(objects)
}

// SetRotation sets the rotation of all pages added to writer. The rotation is
// specified in degrees and must be a multiple of 90.
// The Rotate field of individual pages has priority over the global rotation.
func (_eggfe *PdfWriter) SetRotation(rotate int64) error {
	_ccagg, _efdcf := _cde.GetDict(_eggfe._acbgeg)
	if !_efdcf {
		return ErrTypeCheck
	}
	_ccagg.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _cde.MakeInteger(rotate))
	return nil
}

// PdfOutlineItem represents an outline item dictionary (Table 153 - pp. 376 - 377).
type PdfOutlineItem struct {
	PdfOutlineTreeNode
	Title  *_cde.PdfObjectString
	Parent *PdfOutlineTreeNode
	Prev   *PdfOutlineTreeNode
	Next   *PdfOutlineTreeNode
	Count  *int64
	Dest   _cde.PdfObject
	A      _cde.PdfObject
	SE     _cde.PdfObject
	C      _cde.PdfObject
	F      _cde.PdfObject
	_ccbf  *_cde.PdfIndirectObject
}

func _gbeb(_gde _cde.PdfObject) (*PdfBorderStyle, error) {
	_egbe := &PdfBorderStyle{}
	_egbe._dggc = _gde
	var _edada *_cde.PdfObjectDictionary
	_gde = _cde.TraceToDirectObject(_gde)
	_edada, _bcdg := _gde.(*_cde.PdfObjectDictionary)
	if !_bcdg {
		return nil, _ceg.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	if _ccea := _edada.Get("\u0054\u0079\u0070\u0065"); _ccea != nil {
		_ebfc, _fdc := _ccea.(*_cde.PdfObjectName)
		if !_fdc {
			_ad.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u0077\u0069th\u0020\u0054\u0079\u0070e\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061me\u0020\u006fb\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0054", _ccea)
		} else {
			if *_ebfc != "\u0042\u006f\u0072\u0064\u0065\u0072" {
				_ad.Log.Debug("W\u0061\u0072\u006e\u0069\u006e\u0067,\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020B\u006f\u0072\u0064e\u0072:\u0020\u0025\u0073", *_ebfc)
			}
		}
	}
	if _cdad := _edada.Get("\u0057"); _cdad != nil {
		_bbcd, _ccae := _cde.GetNumberAsFloat(_cdad)
		if _ccae != nil {
			_ad.Log.Debug("\u0045\u0072\u0072\u006fr \u0072\u0065\u0074\u0072\u0069\u0065\u0076\u0069\u006e\u0067\u0020\u0057\u003a\u0020%\u0076", _ccae)
			return nil, _ccae
		}
		_egbe.W = &_bbcd
	}
	if _egddd := _edada.Get("\u0053"); _egddd != nil {
		_gbac, _gcgb := _egddd.(*_cde.PdfObjectName)
		if !_gcgb {
			return nil, _ceg.New("\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u0053\u0020\u006e\u006ft\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		var _acab BorderStyle
		switch *_gbac {
		case "\u0053":
			_acab = BorderStyleSolid
		case "\u0044":
			_acab = BorderStyleDashed
		case "\u0042":
			_acab = BorderStyleBeveled
		case "\u0049":
			_acab = BorderStyleInset
		case "\u0055":
			_acab = BorderStyleUnderline
		default:
			_ad.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020s\u0074\u0079\u006c\u0065\u0020n\u0061\u006d\u0065\u0020\u0025\u0073", *_gbac)
			return nil, _ceg.New("\u0073\u0074\u0079\u006ce \u0074\u0079\u0070\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065c\u006b")
		}
		_egbe.S = &_acab
	}
	if _efbfa := _edada.Get("\u0044"); _efbfa != nil {
		_baea, _ecce := _efbfa.(*_cde.PdfObjectArray)
		if !_ecce {
			_ad.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044\u0020\u0064a\u0073\u0068\u0020\u006e\u006f\u0074\u0020a\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0054", _efbfa)
			return nil, _ceg.New("\u0062o\u0072\u0064\u0065\u0072 \u0044\u0020\u0074\u0079\u0070e\u0020c\u0068e\u0063\u006b\u0020\u0065\u0072\u0072\u006fr")
		}
		_gdgbc, _dfbb := _baea.ToIntegerArray()
		if _dfbb != nil {
			_ad.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044 \u0050\u0072\u006fbl\u0065\u006d\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u006e\u0074\u0065\u0067e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u003a \u0025\u0076", _dfbb)
			return nil, _dfbb
		}
		_egbe.D = &_gdgbc
	}
	return _egbe, nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_beccbe *PdfShadingType1) ToPdfObject() _cde.PdfObject {
	_beccbe.PdfShading.ToPdfObject()
	_dcfcf, _deabg := _beccbe.getShadingDict()
	if _deabg != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _beccbe.Domain != nil {
		_dcfcf.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _beccbe.Domain)
	}
	if _beccbe.Matrix != nil {
		_dcfcf.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _beccbe.Matrix)
	}
	if _beccbe.Function != nil {
		if len(_beccbe.Function) == 1 {
			_dcfcf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _beccbe.Function[0].ToPdfObject())
		} else {
			_bebca := _cde.MakeArray()
			for _, _cedaf := range _beccbe.Function {
				_bebca.Append(_cedaf.ToPdfObject())
			}
			_dcfcf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _bebca)
		}
	}
	return _beccbe._dffg
}

// CharMetrics represents width and height metrics of a glyph.
type CharMetrics = _fe.CharMetrics

func _eegc(_bca *PdfPage) map[_cde.PdfObjectName]_cde.PdfObject {
	_ebfd := make(map[_cde.PdfObjectName]_cde.PdfObject)
	if _bca.Resources == nil {
		return _ebfd
	}
	if _bca.Resources.Font != nil {
		if _gccc, _ddcb := _cde.GetDict(_bca.Resources.Font); _ddcb {
			for _, _aafa := range _gccc.Keys() {
				_ebfd[_aafa] = _gccc.Get(_aafa)
			}
		}
	}
	if _bca.Resources.ExtGState != nil {
		if _gefg, _beee := _cde.GetDict(_bca.Resources.ExtGState); _beee {
			for _, _aceba := range _gefg.Keys() {
				_ebfd[_aceba] = _gefg.Get(_aceba)
			}
		}
	}
	if _bca.Resources.XObject != nil {
		if _faeb, _abfg := _cde.GetDict(_bca.Resources.XObject); _abfg {
			for _, _egg := range _faeb.Keys() {
				_ebfd[_egg] = _faeb.Get(_egg)
			}
		}
	}
	if _bca.Resources.Pattern != nil {
		if _eedc, _bgdb := _cde.GetDict(_bca.Resources.Pattern); _bgdb {
			for _, _dcba := range _eedc.Keys() {
				_ebfd[_dcba] = _eedc.Get(_dcba)
			}
		}
	}
	if _bca.Resources.Shading != nil {
		if _acdg, _ddag := _cde.GetDict(_bca.Resources.Shading); _ddag {
			for _, _fgge := range _acdg.Keys() {
				_ebfd[_fgge] = _acdg.Get(_fgge)
			}
		}
	}
	if _bca.Resources.ProcSet != nil {
		if _ffgf, _aadaf := _cde.GetDict(_bca.Resources.ProcSet); _aadaf {
			for _, _dcga := range _ffgf.Keys() {
				_ebfd[_dcga] = _ffgf.Get(_dcga)
			}
		}
	}
	if _bca.Resources.Properties != nil {
		if _ccac, _gcf := _cde.GetDict(_bca.Resources.Properties); _gcf {
			for _, _gabe := range _ccac.Keys() {
				_ebfd[_gabe] = _ccac.Get(_gabe)
			}
		}
	}
	return _ebfd
}

// ToPdfObject implements interface PdfModel.
func (_bdbf *PdfAnnotationUnderline) ToPdfObject() _cde.PdfObject {
	_bdbf.PdfAnnotation.ToPdfObject()
	_adc := _bdbf._bddg
	_bcf := _adc.PdfObject.(*_cde.PdfObjectDictionary)
	_bdbf.PdfAnnotationMarkup.appendToPdfDictionary(_bcf)
	_bcf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee"))
	_bcf.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _bdbf.QuadPoints)
	return _adc
}

// GetContainingPdfObject returns the containing object for the PdfField, i.e. an indirect object
// containing the field dictionary.
func (_eefca *PdfField) GetContainingPdfObject() _cde.PdfObject { return _eefca._afgc }

// GetMediaBox gets the inheritable media box value, either from the page
// or a higher up page/pages struct.
func (_dega *PdfPage) GetMediaBox() (*PdfRectangle, error) {
	if _dega.MediaBox != nil {
		return _dega.MediaBox, nil
	}
	_egfcc := _dega.Parent
	for _egfcc != nil {
		_fbbcac, _dcddd := _cde.GetDict(_egfcc)
		if !_dcddd {
			return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _ddgee := _fbbcac.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _ddgee != nil {
			_cgafc, _dcbcf := _cde.GetArray(_ddgee)
			if !_dcbcf {
				return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0064\u0069a\u0020\u0062\u006f\u0078")
			}
			_accbf, _bebdd := NewPdfRectangle(*_cgafc)
			if _bebdd != nil {
				return nil, _bebdd
			}
			return _accbf, nil
		}
		_egfcc = _fbbcac.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, _ceg.New("m\u0065\u0064\u0069\u0061 b\u006fx\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
}

// ToPdfObject implements interface PdfModel.
func (_fa *PdfActionImportData) ToPdfObject() _cde.PdfObject {
	_fa.PdfAction.ToPdfObject()
	_cbaa := _fa._bc
	_eff := _cbaa.PdfObject.(*_cde.PdfObjectDictionary)
	_eff.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeImportData)))
	if _fa.F != nil {
		_eff.Set("\u0046", _fa.F.ToPdfObject())
	}
	return _cbaa
}

// NewPdfAnnotationUnderline returns a new text underline annotation.
func NewPdfAnnotationUnderline() *PdfAnnotationUnderline {
	_gge := NewPdfAnnotation()
	_bgb := &PdfAnnotationUnderline{}
	_bgb.PdfAnnotation = _gge
	_bgb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gge.SetContext(_bgb)
	return _bgb
}

// AddAnnotation appends `annot` to the list of page annotations.
func (_eagdf *PdfPage) AddAnnotation(annot *PdfAnnotation) {
	if _eagdf._cefe == nil {
		_eagdf.GetAnnotations()
	}
	_eagdf._cefe = append(_eagdf._cefe, annot)
}

// ApplyStandard is used to apply changes required on the document to match the rules required by the input standard.
// The writer's content would be changed after all the document parts are already established during the Write method.
// A good example of the StandardApplier could be a PDF/A Profile (i.e.: pdfa.Profile1A). In such a case PdfWriter would
// set up all rules required by that Profile.
func (_ebfca *PdfWriter) ApplyStandard(optimizer StandardApplier) { _ebfca._cfaf = optimizer }

const (
	BorderEffectNoEffect BorderEffect = iota
	BorderEffectCloudy   BorderEffect = iota
)

// GetNumComponents returns the number of color components (3 for RGB).
func (_dbfd *PdfColorDeviceRGB) GetNumComponents() int { return 3 }

// PdfField contains the common attributes of a form field. The context object contains the specific field data
// which can represent a button, text, choice or signature.
// The PdfField is typically not used directly, but is encapsulated by the more specific field types such as
// PdfFieldButton etc (i.e. the context attribute).
type PdfField struct {
	_ecfg        PdfModel
	_afgc        *_cde.PdfIndirectObject
	Parent       *PdfField
	Annotations  []*PdfAnnotationWidget
	Kids         []*PdfField
	FT           *_cde.PdfObjectName
	T            *_cde.PdfObjectString
	TU           *_cde.PdfObjectString
	TM           *_cde.PdfObjectString
	Ff           *_cde.PdfObjectInteger
	V            _cde.PdfObject
	DV           _cde.PdfObject
	AA           _cde.PdfObject
	VariableText *VariableText
}

func (_gafc *PdfAnnotationMarkup) appendToPdfDictionary(_dab *_cde.PdfObjectDictionary) {
	_dab.SetIfNotNil("\u0054", _gafc.T)
	if _gafc.Popup != nil {
		_dab.Set("\u0050\u006f\u0070u\u0070", _gafc.Popup.ToPdfObject())
	}
	_dab.SetIfNotNil("\u0043\u0041", _gafc.CA)
	_dab.SetIfNotNil("\u0052\u0043", _gafc.RC)
	_dab.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _gafc.CreationDate)
	_dab.SetIfNotNil("\u0049\u0052\u0054", _gafc.IRT)
	_dab.SetIfNotNil("\u0053\u0075\u0062\u006a", _gafc.Subj)
	_dab.SetIfNotNil("\u0052\u0054", _gafc.RT)
	_dab.SetIfNotNil("\u0049\u0054", _gafc.IT)
	_dab.SetIfNotNil("\u0045\u0078\u0044\u0061\u0074\u0061", _gafc.ExData)
}

// SetOpenAction sets the OpenAction in the PDF catalog.
// The value shall be either an array defining a destination (12.3.2 "Destinations" PDF32000_2008),
// or an action dictionary representing an action (12.6 "Actions" PDF32000_2008).
func (_babb *PdfWriter) SetOpenAction(dest _cde.PdfObject) error {
	if dest == nil || _cde.IsNullObject(dest) {
		return nil
	}
	_babb._fedbb.Set("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e", dest)
	return _babb.addObjects(dest)
}

// CharcodesToUnicodeWithStats is identical to CharcodesToUnicode except it returns more statistical
// information about hits and misses from the reverse mapping process.
// NOTE: The number of runes returned may be greater than the number of charcodes.
// TODO(peterwilliams97): Deprecate in v4 and use only CharcodesToStrings()
func (_ccagc *PdfFont) CharcodesToUnicodeWithStats(charcodes []_gc.CharCode) (_agcaa []rune, _ddcf, _dabcg int) {
	_bbgbd, _ddcf, _dabcg := _ccagc.CharcodesToStrings(charcodes)
	return []rune(_dac.Join(_bbgbd, "")), _ddcf, _dabcg
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_aabca *PdfColorspaceICCBased) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _aabca.Alternate == nil {
		if _aabca.N == 1 {
			_egcf := NewPdfColorspaceDeviceGray()
			return _egcf.ColorFromFloats(vals)
		} else if _aabca.N == 3 {
			_ecd := NewPdfColorspaceDeviceRGB()
			return _ecd.ColorFromFloats(vals)
		} else if _aabca.N == 4 {
			_fgec := NewPdfColorspaceDeviceCMYK()
			return _fgec.ColorFromFloats(vals)
		} else {
			return nil, _ceg.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _aabca.Alternate.ColorFromFloats(vals)
}

// ReplacePage replaces the original page to a new page.
func (_gcd *PdfAppender) ReplacePage(pageNum int, page *PdfPage) {
	_gede := pageNum - 1
	for _bffd := range _gcd._gfb {
		if _bffd == _gede {
			_eede := page.Duplicate()
			_fdcef(_eede)
			_gcd._gfb[_bffd] = _eede
		}
	}
}
func (_dffde *DSS) generateHashMaps() error {
	_cdcd, _egfe := _dffde.generateHashMap(_dffde.Certs)
	if _egfe != nil {
		return _egfe
	}
	_aadgg, _egfe := _dffde.generateHashMap(_dffde.OCSPs)
	if _egfe != nil {
		return _egfe
	}
	_gfgb, _egfe := _dffde.generateHashMap(_dffde.CRLs)
	if _egfe != nil {
		return _egfe
	}
	_dffde._cddc = _cdcd
	_dffde._ecgd = _aadgg
	_dffde._fbbd = _gfgb
	return nil
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_dffe *PdfColorspaceCalRGB) ToPdfObject() _cde.PdfObject {
	_gbbgc := &_cde.PdfObjectArray{}
	_gbbgc.Append(_cde.MakeName("\u0043\u0061\u006c\u0052\u0047\u0042"))
	_baad := _cde.MakeDict()
	if _dffe.WhitePoint != nil {
		_fccd := _cde.MakeArray(_cde.MakeFloat(_dffe.WhitePoint[0]), _cde.MakeFloat(_dffe.WhitePoint[1]), _cde.MakeFloat(_dffe.WhitePoint[2]))
		_baad.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _fccd)
	} else {
		_ad.Log.Error("\u0043\u0061l\u0052\u0047\u0042\u003a \u004d\u0069s\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _dffe.BlackPoint != nil {
		_feea := _cde.MakeArray(_cde.MakeFloat(_dffe.BlackPoint[0]), _cde.MakeFloat(_dffe.BlackPoint[1]), _cde.MakeFloat(_dffe.BlackPoint[2]))
		_baad.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _feea)
	}
	if _dffe.Gamma != nil {
		_ffeg := _cde.MakeArray(_cde.MakeFloat(_dffe.Gamma[0]), _cde.MakeFloat(_dffe.Gamma[1]), _cde.MakeFloat(_dffe.Gamma[2]))
		_baad.Set("\u0047\u0061\u006dm\u0061", _ffeg)
	}
	if _dffe.Matrix != nil {
		_eag := _cde.MakeArray(_cde.MakeFloat(_dffe.Matrix[0]), _cde.MakeFloat(_dffe.Matrix[1]), _cde.MakeFloat(_dffe.Matrix[2]), _cde.MakeFloat(_dffe.Matrix[3]), _cde.MakeFloat(_dffe.Matrix[4]), _cde.MakeFloat(_dffe.Matrix[5]), _cde.MakeFloat(_dffe.Matrix[6]), _cde.MakeFloat(_dffe.Matrix[7]), _cde.MakeFloat(_dffe.Matrix[8]))
		_baad.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _eag)
	}
	_gbbgc.Append(_baad)
	if _dffe._efab != nil {
		_dffe._efab.PdfObject = _gbbgc
		return _dffe._efab
	}
	return _gbbgc
}

const (
	ButtonTypeCheckbox ButtonType = iota
	ButtonTypePush     ButtonType = iota
	ButtonTypeRadio    ButtonType = iota
)

func _bfeg(_ecec *_cde.PdfIndirectObject) (*PdfOutline, error) {
	_cbgcd, _gedc := _ecec.PdfObject.(*_cde.PdfObjectDictionary)
	if !_gedc {
		return nil, _ee.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_fcfg := NewPdfOutline()
	if _gffe := _cbgcd.Get("\u0054\u0079\u0070\u0065"); _gffe != nil {
		_bdfb, _dbcce := _gffe.(*_cde.PdfObjectName)
		if _dbcce {
			if *_bdfb != "\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073" {
				_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u004f\u0075\u0074l\u0069\u006e\u0065s\u0020(\u0025\u0073\u0029", *_bdfb)
			}
		}
	}
	if _fcfcb := _cbgcd.Get("\u0043\u006f\u0075n\u0074"); _fcfcb != nil {
		_cffefa, _fceb := _cde.GetNumberAsInt64(_fcfcb)
		if _fceb != nil {
			return nil, _fceb
		}
		_fcfg.Count = &_cffefa
	}
	return _fcfg, nil
}
func _gccbg(_gecba *_cde.PdfObjectDictionary) (*PdfShadingType1, error) {
	_ababg := PdfShadingType1{}
	if _efcg := _gecba.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _efcg != nil {
		_efcg = _cde.TraceToDirectObject(_efcg)
		_cfeg, _gccfc := _efcg.(*_cde.PdfObjectArray)
		if !_gccfc {
			_ad.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _efcg)
			return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ababg.Domain = _cfeg
	}
	if _eacbe := _gecba.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _eacbe != nil {
		_eacbe = _cde.TraceToDirectObject(_eacbe)
		_ebgged, _afee := _eacbe.(*_cde.PdfObjectArray)
		if !_afee {
			_ad.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _eacbe)
			return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ababg.Matrix = _ebgged
	}
	_gcgc := _gecba.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _gcgc == nil {
		_ad.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_ababg.Function = []PdfFunction{}
	if _ddae, _fdeac := _gcgc.(*_cde.PdfObjectArray); _fdeac {
		for _, _dadfe := range _ddae.Elements() {
			_gbdcc, _ccbea := _cfdbb(_dadfe)
			if _ccbea != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _ccbea)
				return nil, _ccbea
			}
			_ababg.Function = append(_ababg.Function, _gbdcc)
		}
	} else {
		_acfd, _eddgg := _cfdbb(_gcgc)
		if _eddgg != nil {
			_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _eddgg)
			return nil, _eddgg
		}
		_ababg.Function = append(_ababg.Function, _acfd)
	}
	return &_ababg, nil
}

// IsCheckbox returns true if the button field represents a checkbox, false otherwise.
func (_fgdd *PdfFieldButton) IsCheckbox() bool { return _fgdd.GetType() == ButtonTypeCheckbox }

// SetPdfProducer sets the Producer attribute of the output PDF.
func SetPdfProducer(producer string) { _dccfe.Lock(); defer _dccfe.Unlock(); _adcff = producer }

// ImageToRGB converts an image with samples in Separation CS to an image with samples specified in
// DeviceRGB CS.
func (_fafc *PdfColorspaceSpecialSeparation) ImageToRGB(img Image) (Image, error) {
	_fbabg := _cae.NewReader(img.getBase())
	_cgbb := _ff.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), _fafc.AlternateSpace.GetNumComponents(), nil, img._deegf, nil)
	_eebd := _cae.NewWriter(_cgbb)
	_fbea := _ced.Pow(2, float64(img.BitsPerComponent)) - 1
	_ad.Log.Trace("\u0053\u0065\u0070a\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u002d\u003e\u0020\u0054\u006f\u0052\u0047\u0042\u0020\u0063o\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	_ad.Log.Trace("\u0054i\u006et\u0054\u0072\u0061\u006e\u0073f\u006f\u0072m\u003a\u0020\u0025\u002b\u0076", _fafc.TintTransform)
	_fegef := _fafc.AlternateSpace.DecodeArray()
	var (
		_bcec uint32
		_bdfc error
	)
	for {
		_bcec, _bdfc = _fbabg.ReadSample()
		if _bdfc == _f.EOF {
			break
		}
		if _bdfc != nil {
			return img, _bdfc
		}
		_ceggb := float64(_bcec) / _fbea
		_acgfc, _dffd := _fafc.TintTransform.Evaluate([]float64{_ceggb})
		if _dffd != nil {
			return img, _dffd
		}
		for _aefa, _eeca := range _acgfc {
			_acbfg := _ff.LinearInterpolate(_eeca, _fegef[_aefa*2], _fegef[_aefa*2+1], 0, 1)
			if _dffd = _eebd.WriteSample(uint32(_acbfg * _fbea)); _dffd != nil {
				return img, _dffd
			}
		}
	}
	return _fafc.AlternateSpace.ImageToRGB(_bddb(&_cgbb))
}

// NewPdfField returns an initialized PdfField.
func NewPdfField() *PdfField { return &PdfField{_afgc: _cde.MakeIndirectObject(_cde.MakeDict())} }

// String returns string value of output intent for given type
// ISO_19005-2 6.2.3: GTS_PDFA1 value should be used for PDF/A-1, A-2 and A-3 at least
func (_dcgad PdfOutputIntentType) String() string {
	switch _dcgad {
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
	_gfbb         []byte
	_bgacg        []uint32
	_fefbf        *_cde.PdfObjectStream
}

func _gedgbf(_ebadb *_cde.PdfObjectDictionary) (*PdfShadingType7, error) {
	_daca := PdfShadingType7{}
	_bbbga := _ebadb.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _bbbga == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_gegfb, _agaab := _bbbga.(*_cde.PdfObjectInteger)
	if !_agaab {
		_ad.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bbbga)
		return nil, _cde.ErrTypeError
	}
	_daca.BitsPerCoordinate = _gegfb
	_bbbga = _ebadb.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _bbbga == nil {
		_ad.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_gegfb, _agaab = _bbbga.(*_cde.PdfObjectInteger)
	if !_agaab {
		_ad.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _bbbga)
		return nil, _cde.ErrTypeError
	}
	_daca.BitsPerComponent = _gegfb
	_bbbga = _ebadb.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _bbbga == nil {
		_ad.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_gegfb, _agaab = _bbbga.(*_cde.PdfObjectInteger)
	if !_agaab {
		_ad.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _bbbga)
		return nil, _cde.ErrTypeError
	}
	_daca.BitsPerComponent = _gegfb
	_bbbga = _ebadb.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _bbbga == nil {
		_ad.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_ecceg, _agaab := _bbbga.(*_cde.PdfObjectArray)
	if !_agaab {
		_ad.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _bbbga)
		return nil, _cde.ErrTypeError
	}
	_daca.Decode = _ecceg
	if _dcacb := _ebadb.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _dcacb != nil {
		_daca.Function = []PdfFunction{}
		if _acagf, _dceeg := _dcacb.(*_cde.PdfObjectArray); _dceeg {
			for _, _eebdb := range _acagf.Elements() {
				_cgbcf, _cecdg := _cfdbb(_eebdb)
				if _cecdg != nil {
					_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cecdg)
					return nil, _cecdg
				}
				_daca.Function = append(_daca.Function, _cgbcf)
			}
		} else {
			_bccfe, _cada := _cfdbb(_dcacb)
			if _cada != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cada)
				return nil, _cada
			}
			_daca.Function = append(_daca.Function, _bccfe)
		}
	}
	return &_daca, nil
}

// ToPdfObject returns the PDF representation of the function.
func (_gfedf *PdfFunctionType2) ToPdfObject() _cde.PdfObject {
	_cefgb := _cde.MakeDict()
	_cefgb.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _cde.MakeInteger(2))
	_bgfbc := &_cde.PdfObjectArray{}
	for _, _fadec := range _gfedf.Domain {
		_bgfbc.Append(_cde.MakeFloat(_fadec))
	}
	_cefgb.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _bgfbc)
	if _gfedf.Range != nil {
		_gbfec := &_cde.PdfObjectArray{}
		for _, _deeb := range _gfedf.Range {
			_gbfec.Append(_cde.MakeFloat(_deeb))
		}
		_cefgb.Set("\u0052\u0061\u006eg\u0065", _gbfec)
	}
	if _gfedf.C0 != nil {
		_fegb := &_cde.PdfObjectArray{}
		for _, _dgcdd := range _gfedf.C0 {
			_fegb.Append(_cde.MakeFloat(_dgcdd))
		}
		_cefgb.Set("\u0043\u0030", _fegb)
	}
	if _gfedf.C1 != nil {
		_badf := &_cde.PdfObjectArray{}
		for _, _bbbb := range _gfedf.C1 {
			_badf.Append(_cde.MakeFloat(_bbbb))
		}
		_cefgb.Set("\u0043\u0031", _badf)
	}
	_cefgb.Set("\u004e", _cde.MakeFloat(_gfedf.N))
	if _gfedf._bddc != nil {
		_gfedf._bddc.PdfObject = _cefgb
		return _gfedf._bddc
	}
	return _cefgb
}

// IsEncrypted returns true if the PDF file is encrypted.
func (_gbbfb *PdfReader) IsEncrypted() (bool, error)    { return _gbbfb._aggcgb.IsEncrypted() }
func (_effcg *pdfCIDFontType0) baseFields() *fontCommon { return &_effcg.fontCommon }

type fontCommon struct {
	_eeab  string
	_dcbc  string
	_fddeb string
	_dfae  _cde.PdfObject
	_ggebg *_fb.CMap
	_fagf  *PdfFontDescriptor
	_dbadd int64
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_bef *PdfColorspaceSpecialSeparation) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_dcgag, _facd := _cde.GetNumbersAsFloat(objects)
	if _facd != nil {
		return nil, _facd
	}
	return _bef.ColorFromFloats(_dcgag)
}

// String implements interface PdfObject.
func (_adba *PdfAction) String() string {
	_ge, _dga := _adba.ToPdfObject().(*_cde.PdfIndirectObject)
	if _dga {
		return _ee.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _adba._bgd, _ge.PdfObject.String())
	}
	return ""
}
func (_dbfeb *pdfFontType0) getFontDescriptor() *PdfFontDescriptor {
	if _dbfeb._fagf == nil && _dbfeb.DescendantFont != nil {
		return _dbfeb.DescendantFont.FontDescriptor()
	}
	return _dbfeb._fagf
}

// ToPdfObject implements interface PdfModel.
func (_dbc *PdfAnnotation3D) ToPdfObject() _cde.PdfObject {
	_dbc.PdfAnnotation.ToPdfObject()
	_gafd := _dbc._bddg
	_ada := _gafd.PdfObject.(*_cde.PdfObjectDictionary)
	_ada.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0033\u0044"))
	_ada.SetIfNotNil("\u0033\u0044\u0044", _dbc.T3DD)
	_ada.SetIfNotNil("\u0033\u0044\u0056", _dbc.T3DV)
	_ada.SetIfNotNil("\u0033\u0044\u0041", _dbc.T3DA)
	_ada.SetIfNotNil("\u0033\u0044\u0049", _dbc.T3DI)
	_ada.SetIfNotNil("\u0033\u0044\u0042", _dbc.T3DB)
	return _gafd
}

// SetPdfTitle sets the Title attribute of the output PDF.
func SetPdfTitle(title string) { _dccfe.Lock(); defer _dccfe.Unlock(); _dgcdb = title }

// Items returns all children outline items.
func (_dgbb *Outline) Items() []*OutlineItem { return _dgbb.Entries }

// String returns a string representation of the field.
func (_bcbeb *PdfField) String() string {
	if _gbaac, _cfff := _bcbeb.ToPdfObject().(*_cde.PdfIndirectObject); _cfff {
		return _ee.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _bcbeb._ecfg, _gbaac.PdfObject.String())
	}
	return ""
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_edfbf *PdfFunctionType0) Evaluate(x []float64) ([]float64, error) {
	if len(x) != _edfbf.NumInputs {
		_ad.Log.Error("\u004eu\u006d\u0062e\u0072\u0020\u006f\u0066 \u0069\u006e\u0070u\u0074\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061tc\u0068\u0069\u006eg\u0020\u0077h\u0061\u0074\u0020\u0069\u0073\u0020n\u0065\u0065d\u0065\u0064")
		return nil, _ceg.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _edfbf._bgacg == nil {
		_caadd := _edfbf.processSamples()
		if _caadd != nil {
			return nil, _caadd
		}
	}
	_fecba := _edfbf.Encode
	if _fecba == nil {
		_fecba = []float64{}
		for _cafdd := 0; _cafdd < len(_edfbf.Size); _cafdd++ {
			_fecba = append(_fecba, 0)
			_fecba = append(_fecba, float64(_edfbf.Size[_cafdd]-1))
		}
	}
	_acgge := _edfbf.Decode
	if _acgge == nil {
		_acgge = _edfbf.Range
	}
	_bffcc := make([]int, len(x))
	for _gecc := 0; _gecc < len(x); _gecc++ {
		_bgeda := x[_gecc]
		_bgedad := _ced.Min(_ced.Max(_bgeda, _edfbf.Domain[2*_gecc]), _edfbf.Domain[2*_gecc+1])
		_fcgbc := _ff.LinearInterpolate(_bgedad, _edfbf.Domain[2*_gecc], _edfbf.Domain[2*_gecc+1], _fecba[2*_gecc], _fecba[2*_gecc+1])
		_efaga := _ced.Min(_ced.Max(_fcgbc, 0), float64(_edfbf.Size[_gecc]-1))
		_adff := int(_ced.Floor(_efaga + 0.5))
		if _adff < 0 {
			_adff = 0
		} else if _adff > _edfbf.Size[_gecc] {
			_adff = _edfbf.Size[_gecc] - 1
		}
		_bffcc[_gecc] = _adff
	}
	_fafaf := _bffcc[0]
	for _gegcg := 1; _gegcg < _edfbf.NumInputs; _gegcg++ {
		_dccbg := _bffcc[_gegcg]
		for _aaadd := 0; _aaadd < _gegcg; _aaadd++ {
			_dccbg *= _edfbf.Size[_aaadd]
		}
		_fafaf += _dccbg
	}
	_fafaf *= _edfbf.NumOutputs
	var _cecg []float64
	for _gcdea := 0; _gcdea < _edfbf.NumOutputs; _gcdea++ {
		_caebg := _fafaf + _gcdea
		if _caebg >= len(_edfbf._bgacg) {
			_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u006e\u006ft\u0020\u0065\u006eo\u0075\u0067\u0068\u0020\u0069\u006ep\u0075\u0074\u0020sa\u006dp\u006c\u0065\u0073\u0020\u0074\u006f\u0020d\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0076\u0061lu\u0065\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			continue
		}
		_gcegf := _edfbf._bgacg[_caebg]
		_daff := _ff.LinearInterpolate(float64(_gcegf), 0, _ced.Pow(2, float64(_edfbf.BitsPerSample)), _acgge[2*_gcdea], _acgge[2*_gcdea+1])
		_ddbfb := _ced.Min(_ced.Max(_daff, _edfbf.Range[2*_gcdea]), _edfbf.Range[2*_gcdea+1])
		_cecg = append(_cecg, _ddbfb)
	}
	return _cecg, nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_afdeeg *PdfShading) ToPdfObject() _cde.PdfObject {
	_ecbce := _afdeeg._dffg
	_egfdc, _dafdc := _afdeeg.getShadingDict()
	if _dafdc != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _afdeeg.ShadingType != nil {
		_egfdc.Set("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _afdeeg.ShadingType)
	}
	if _afdeeg.ColorSpace != nil {
		_egfdc.Set("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _afdeeg.ColorSpace.ToPdfObject())
	}
	if _afdeeg.Background != nil {
		_egfdc.Set("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064", _afdeeg.Background)
	}
	if _afdeeg.BBox != nil {
		_egfdc.Set("\u0042\u0042\u006f\u0078", _afdeeg.BBox.ToPdfObject())
	}
	if _afdeeg.AntiAlias != nil {
		_egfdc.Set("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s", _afdeeg.AntiAlias)
	}
	return _ecbce
}
func (_bagab *XObjectImage) getParamsDict() *_cde.PdfObjectDictionary {
	_fecga := _cde.MakeDict()
	_fecga.Set("\u0057\u0069\u0064t\u0068", _cde.MakeInteger(*_bagab.Width))
	_fecga.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _cde.MakeInteger(*_bagab.Height))
	_fecga.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _cde.MakeInteger(int64(_bagab.ColorSpace.GetNumComponents())))
	_fecga.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _cde.MakeInteger(*_bagab.BitsPerComponent))
	return _fecga
}

// ColorToRGB converts a CalRGB color to an RGB color.
func (_cagb *PdfColorspaceCalRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_dccb, _fbab := color.(*PdfColorCalRGB)
	if !_fbab {
		_ad.Log.Debug("\u0049\u006e\u0070ut\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0063\u0061\u006c\u0020\u0072\u0067\u0062")
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_baaf := _dccb.A()
	_bgba := _dccb.B()
	_efaa := _dccb.C()
	X := _cagb.Matrix[0]*_ced.Pow(_baaf, _cagb.Gamma[0]) + _cagb.Matrix[3]*_ced.Pow(_bgba, _cagb.Gamma[1]) + _cagb.Matrix[6]*_ced.Pow(_efaa, _cagb.Gamma[2])
	Y := _cagb.Matrix[1]*_ced.Pow(_baaf, _cagb.Gamma[0]) + _cagb.Matrix[4]*_ced.Pow(_bgba, _cagb.Gamma[1]) + _cagb.Matrix[7]*_ced.Pow(_efaa, _cagb.Gamma[2])
	Z := _cagb.Matrix[2]*_ced.Pow(_baaf, _cagb.Gamma[0]) + _cagb.Matrix[5]*_ced.Pow(_bgba, _cagb.Gamma[1]) + _cagb.Matrix[8]*_ced.Pow(_efaa, _cagb.Gamma[2])
	_bbbc := 3.240479*X + -1.537150*Y + -0.498535*Z
	_fegd := -0.969256*X + 1.875992*Y + 0.041556*Z
	_efed := 0.055648*X + -0.204043*Y + 1.057311*Z
	_bbbc = _ced.Min(_ced.Max(_bbbc, 0), 1.0)
	_fegd = _ced.Min(_ced.Max(_fegd, 0), 1.0)
	_efed = _ced.Min(_ced.Max(_efed, 0), 1.0)
	return NewPdfColorDeviceRGB(_bbbc, _fegd, _efed), nil
}

// Items returns all children outline items.
func (_ebbde *OutlineItem) Items() []*OutlineItem { return _ebbde.Entries }
func (_adeb *PdfReader) newPdfAnnotationScreenFromDict(_edebe *_cde.PdfObjectDictionary) (*PdfAnnotationScreen, error) {
	_beaf := PdfAnnotationScreen{}
	_beaf.T = _edebe.Get("\u0054")
	_beaf.MK = _edebe.Get("\u004d\u004b")
	_beaf.A = _edebe.Get("\u0041")
	_beaf.AA = _edebe.Get("\u0041\u0041")
	return &_beaf, nil
}
func (_gadcg *PdfWriter) writeObjects() {
	_ad.Log.Trace("\u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0025d\u0020\u006f\u0062\u006a", len(_gadcg._egbccc))
	_gadcg._gfdac = make(map[int]crossReference)
	_gadcg._gfdac[0] = crossReference{Type: 0, ObjectNumber: 0, Generation: 0xFFFF}
	if _gadcg._bdgaa.ObjectMap != nil {
		for _ffaec, _ggdbc := range _gadcg._bdgaa.ObjectMap {
			if _ffaec == 0 {
				continue
			}
			if _ggdbc.XType == _cde.XrefTypeObjectStream {
				_afbcg := crossReference{Type: 2, ObjectNumber: _ggdbc.OsObjNumber, Index: _ggdbc.OsObjIndex}
				_gadcg._gfdac[_ffaec] = _afbcg
			}
			if _ggdbc.XType == _cde.XrefTypeTableEntry {
				_eccfa := crossReference{Type: 1, ObjectNumber: _ggdbc.ObjectNumber, Offset: _ggdbc.Offset}
				_gadcg._gfdac[_ffaec] = _eccfa
			}
		}
	}
}

// GetContainingPdfObject returns the container of the outline tree node (indirect object).
func (_ebecb *PdfOutlineTreeNode) GetContainingPdfObject() _cde.PdfObject {
	return _ebecb.GetContext().GetContainingPdfObject()
}

// SetLocation sets the `Location` field of the signature.
func (_afcge *PdfSignature) SetLocation(location string) { _afcge.Location = _cde.MakeString(location) }

// NewPdfAnnotationStrikeOut returns a new text strikeout annotation.
func NewPdfAnnotationStrikeOut() *PdfAnnotationStrikeOut {
	_fec := NewPdfAnnotation()
	_dee := &PdfAnnotationStrikeOut{}
	_dee.PdfAnnotation = _fec
	_dee.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fec.SetContext(_dee)
	return _dee
}

// NewPdfAnnotationPolyLine returns a new polyline annotation.
func NewPdfAnnotationPolyLine() *PdfAnnotationPolyLine {
	_dged := NewPdfAnnotation()
	_bffa := &PdfAnnotationPolyLine{}
	_bffa.PdfAnnotation = _dged
	_bffa.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dged.SetContext(_bffa)
	return _bffa
}

// ToPdfObject implements interface PdfModel.
func (_bd *PdfAction) ToPdfObject() _cde.PdfObject {
	_gd := _bd._bc
	_dbe := _gd.PdfObject.(*_cde.PdfObjectDictionary)
	_dbe.Clear()
	_dbe.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0041\u0063\u0074\u0069\u006f\u006e"))
	_dbe.SetIfNotNil("\u0053", _bd.S)
	_dbe.SetIfNotNil("\u004e\u0065\u0078\u0074", _bd.Next)
	return _gd
}

// BaseFont returns the font's "BaseFont" field.
func (_aaaa *PdfFont) BaseFont() string { return _aaaa.baseFields()._eeab }

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_ddcbg *PdfColorspaceSpecialPattern) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _ddcbg.UnderlyingCS == nil {
		return nil, _ceg.New("u\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0043\u0053\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
	}
	return _ddcbg.UnderlyingCS.ColorFromFloats(vals)
}
func (_abeb *PdfReader) loadAction(_efa _cde.PdfObject) (*PdfAction, error) {
	if _fead, _gfd := _cde.GetIndirect(_efa); _gfd {
		_cbgc, _fcf := _abeb.newPdfActionFromIndirectObject(_fead)
		if _fcf != nil {
			return nil, _fcf
		}
		return _cbgc, nil
	} else if !_cde.IsNullObject(_efa) {
		return nil, _ceg.New("\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074 \u0074\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, nil
}

type pdfFontSimple struct {
	fontCommon
	_acebaa *_cde.PdfIndirectObject
	_fecg   map[_gc.CharCode]float64
	_efeaf  _gc.TextEncoder
	_facga  _gc.TextEncoder
	_ccff   *PdfFontDescriptor

	// Encoding is subject to limitations that are described in 9.6.6, "Character Encoding".
	// BaseFont is derived differently.
	FirstChar _cde.PdfObject
	LastChar  _cde.PdfObject
	Widths    _cde.PdfObject
	Encoding  _cde.PdfObject
	_gggc     *_fe.RuneCharSafeMap
}

func (_bdecg *pdfFontType0) bytesToCharcodes(_bceed []byte) ([]_gc.CharCode, bool) {
	if _bdecg._dcdga == nil {
		return nil, false
	}
	_faeeg, _gabbc := _bdecg._dcdga.BytesToCharcodes(_bceed)
	if !_gabbc {
		return nil, false
	}
	_cdfea := make([]_gc.CharCode, len(_faeeg))
	for _bcdae, _bgecb := range _faeeg {
		_cdfea[_bcdae] = _gc.CharCode(_bgecb)
	}
	return _cdfea, true
}

// AddFont adds a font dictionary to the Font resources.
func (_eafdd *PdfPage) AddFont(name _cde.PdfObjectName, font _cde.PdfObject) error {
	if _eafdd.Resources == nil {
		_eafdd.Resources = NewPdfPageResources()
	}
	if _eafdd.Resources.Font == nil {
		_eafdd.Resources.Font = _cde.MakeDict()
	}
	_egdg, _ddedb := _cde.TraceToDirectObject(_eafdd.Resources.Font).(*_cde.PdfObjectDictionary)
	if !_ddedb {
		_ad.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0066\u006f\u006et \u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a \u0025\u0076", _cde.TraceToDirectObject(_eafdd.Resources.Font))
		return _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_egdg.Set(name, font)
	return nil
}
func (_fged *PdfAcroForm) fillImageWithAppearance(_cdbc FieldImageProvider, _ecgdf FieldAppearanceGenerator) error {
	if _fged == nil {
		return nil
	}
	_fbegd, _ccdb := _cdbc.FieldImageValues()
	if _ccdb != nil {
		return _ccdb
	}
	for _, _eeagc := range _fged.AllFields() {
		_fgda := _eeagc.PartialName()
		_aabee, _abba := _fbegd[_fgda]
		if !_abba {
			if _eccfe, _efabd := _eeagc.FullName(); _efabd == nil {
				_aabee, _abba = _fbegd[_eccfe]
			}
		}
		if !_abba {
			_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _fgda)
			continue
		}
		switch _gdgfc := _eeagc.GetContext().(type) {
		case *PdfFieldButton:
			if _gdgfc.IsPush() {
				_gdgfc.SetFillImage(_aabee)
			}
		}
		if _ecgdf == nil {
			continue
		}
		for _, _gefb := range _eeagc.Annotations {
			_fgfgb, _ecaec := _ecgdf.GenerateAppearanceDict(_fged, _eeagc, _gefb)
			if _ecaec != nil {
				return _ecaec
			}
			_gefb.AP = _fgfgb
			_gefb.ToPdfObject()
		}
	}
	return nil
}

// PdfShadingType2 is an Axial shading.
type PdfShadingType2 struct {
	*PdfShading
	Coords   *_cde.PdfObjectArray
	Domain   *_cde.PdfObjectArray
	Function []PdfFunction
	Extend   *_cde.PdfObjectArray
}

// NewPdfAnnotationRedact returns a new redact annotation.
func NewPdfAnnotationRedact() *PdfAnnotationRedact {
	_adg := NewPdfAnnotation()
	_feg := &PdfAnnotationRedact{}
	_feg.PdfAnnotation = _adg
	_feg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_adg.SetContext(_feg)
	return _feg
}

// GetNumPages returns the number of pages in the document.
func (_fdfab *PdfReader) GetNumPages() (int, error) {
	if _fdfab._aggcgb.GetCrypter() != nil && !_fdfab._aggcgb.IsAuthenticated() {
		return 0, _ee.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	return len(_fdfab._gbfgg), nil
}

type pdfSignDictionary struct {
	*_cde.PdfObjectDictionary
	_addcc *SignatureHandler
	_afafa *PdfSignature
	_aafab int64
	_gggca int
	_cgedg int
	_faba  int
	_edgb  int
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_ddac *PdfShadingType4) ToPdfObject() _cde.PdfObject {
	_ddac.PdfShading.ToPdfObject()
	_bacf, _cdgce := _ddac.getShadingDict()
	if _cdgce != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _ddac.BitsPerCoordinate != nil {
		_bacf.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _ddac.BitsPerCoordinate)
	}
	if _ddac.BitsPerComponent != nil {
		_bacf.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _ddac.BitsPerComponent)
	}
	if _ddac.BitsPerFlag != nil {
		_bacf.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _ddac.BitsPerFlag)
	}
	if _ddac.Decode != nil {
		_bacf.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _ddac.Decode)
	}
	if _ddac.Function != nil {
		if len(_ddac.Function) == 1 {
			_bacf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _ddac.Function[0].ToPdfObject())
		} else {
			_egbgc := _cde.MakeArray()
			for _, _abcgd := range _ddac.Function {
				_egbgc.Append(_abcgd.ToPdfObject())
			}
			_bacf.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _egbgc)
		}
	}
	return _ddac._dffg
}

// VRI represents a Validation-Related Information dictionary.
// The VRI dictionary contains validation data in the form of
// certificates, OCSP and CRL information, for a single signature.
// See ETSI TS 102 778-4 V1.1.1 for more information.
type VRI struct {
	Cert []*_cde.PdfObjectStream
	OCSP []*_cde.PdfObjectStream
	CRL  []*_cde.PdfObjectStream
	TU   *_cde.PdfObjectString
	TS   *_cde.PdfObjectString
}

func (_gbgc *PdfReader) newPdfAnnotationLineFromDict(_ebe *_cde.PdfObjectDictionary) (*PdfAnnotationLine, error) {
	_ccbg := PdfAnnotationLine{}
	_daa, _bfee := _gbgc.newPdfAnnotationMarkupFromDict(_ebe)
	if _bfee != nil {
		return nil, _bfee
	}
	_ccbg.PdfAnnotationMarkup = _daa
	_ccbg.L = _ebe.Get("\u004c")
	_ccbg.BS = _ebe.Get("\u0042\u0053")
	_ccbg.LE = _ebe.Get("\u004c\u0045")
	_ccbg.IC = _ebe.Get("\u0049\u0043")
	_ccbg.LL = _ebe.Get("\u004c\u004c")
	_ccbg.LLE = _ebe.Get("\u004c\u004c\u0045")
	_ccbg.Cap = _ebe.Get("\u0043\u0061\u0070")
	_ccbg.IT = _ebe.Get("\u0049\u0054")
	_ccbg.LLO = _ebe.Get("\u004c\u004c\u004f")
	_ccbg.CP = _ebe.Get("\u0043\u0050")
	_ccbg.Measure = _ebe.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	_ccbg.CO = _ebe.Get("\u0043\u004f")
	return &_ccbg, nil
}
func (_cfcgg *PdfReader) newPdfPageFromDict(_egggd *_cde.PdfObjectDictionary) (*PdfPage, error) {
	_bfaeag := NewPdfPage()
	_bfaeag._gbbc = _egggd
	_gfggf := *_egggd
	_cgeg, _fbeeg := _gfggf.Get("\u0054\u0079\u0070\u0065").(*_cde.PdfObjectName)
	if !_fbeeg {
		return nil, _ceg.New("\u006d\u0069ss\u0069\u006e\u0067/\u0069\u006e\u0076\u0061lid\u0020Pa\u0067\u0065\u0020\u0064\u0069\u0063\u0074io\u006e\u0061\u0072\u0079\u0020\u0054\u0079p\u0065")
	}
	if *_cgeg != "\u0050\u0061\u0067\u0065" {
		return nil, _ceg.New("\u0070\u0061\u0067\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u0050a\u0067\u0065")
	}
	if _cgbca := _gfggf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"); _cgbca != nil {
		_bfaeag.Parent = _cgbca
	}
	if _abbea := _gfggf.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"); _abbea != nil {
		_gage, _fddba := _cde.GetString(_abbea)
		if !_fddba {
			return nil, _ceg.New("\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u004c\u0061\u0073\u0074\u004d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0021=\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		_aadb, _gccff := NewPdfDate(_gage.Str())
		if _gccff != nil {
			return nil, _gccff
		}
		_bfaeag.LastModified = &_aadb
	}
	if _ggfed := _gfggf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _ggfed != nil && !_cde.IsNullObject(_ggfed) {
		_cafcd, _daaa := _cde.GetDict(_ggfed)
		if !_daaa {
			return nil, _ee.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _ggfed)
		}
		var _ebedc error
		_bfaeag.Resources, _ebedc = NewPdfPageResourcesFromDict(_cafcd)
		if _ebedc != nil {
			return nil, _ebedc
		}
	} else {
		_efdfe, _acfgd := _bfaeag.getParentResources()
		if _acfgd != nil {
			return nil, _acfgd
		}
		if _efdfe == nil {
			_efdfe = NewPdfPageResources()
		}
		_bfaeag.Resources = _efdfe
	}
	if _gaebf := _gfggf.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _gaebf != nil {
		_fbfa, _facac := _cde.GetArray(_gaebf)
		if !_facac {
			return nil, _ceg.New("\u0070\u0061\u0067\u0065\u0020\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _dfgdf error
		_bfaeag.MediaBox, _dfgdf = NewPdfRectangle(*_fbfa)
		if _dfgdf != nil {
			return nil, _dfgdf
		}
	}
	if _gbceb := _gfggf.Get("\u0043r\u006f\u0070\u0042\u006f\u0078"); _gbceb != nil {
		_ecgddf, _bafbg := _cde.GetArray(_gbceb)
		if !_bafbg {
			return nil, _ceg.New("\u0070a\u0067\u0065\u0020\u0043r\u006f\u0070\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _fffac error
		_bfaeag.CropBox, _fffac = NewPdfRectangle(*_ecgddf)
		if _fffac != nil {
			return nil, _fffac
		}
	}
	if _gcfae := _gfggf.Get("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078"); _gcfae != nil {
		_afac, _bcdcd := _cde.GetArray(_gcfae)
		if !_bcdcd {
			return nil, _ceg.New("\u0070\u0061\u0067\u0065\u0020\u0042\u006c\u0065\u0065\u0064\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _defdea error
		_bfaeag.BleedBox, _defdea = NewPdfRectangle(*_afac)
		if _defdea != nil {
			return nil, _defdea
		}
	}
	if _cbcc := _gfggf.Get("\u0054r\u0069\u006d\u0042\u006f\u0078"); _cbcc != nil {
		_ecbff, _ddcec := _cde.GetArray(_cbcc)
		if !_ddcec {
			return nil, _ceg.New("\u0070a\u0067\u0065\u0020\u0054r\u0069\u006d\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _gdgfe error
		_bfaeag.TrimBox, _gdgfe = NewPdfRectangle(*_ecbff)
		if _gdgfe != nil {
			return nil, _gdgfe
		}
	}
	if _bcebb := _gfggf.Get("\u0041\u0072\u0074\u0042\u006f\u0078"); _bcebb != nil {
		_cabfa, _ffggg := _cde.GetArray(_bcebb)
		if !_ffggg {
			return nil, _ceg.New("\u0070a\u0067\u0065\u0020\u0041\u0072\u0074\u0042\u006f\u0078\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		var _fbddc error
		_bfaeag.ArtBox, _fbddc = NewPdfRectangle(*_cabfa)
		if _fbddc != nil {
			return nil, _fbddc
		}
	}
	if _ccbgg := _gfggf.Get("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f"); _ccbgg != nil {
		_bfaeag.BoxColorInfo = _ccbgg
	}
	if _gdbbg := _gfggf.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _gdbbg != nil {
		_bfaeag.Contents = _gdbbg
	}
	if _edbff := _gfggf.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _edbff != nil {
		_dedde, _dedb := _cde.GetNumberAsInt64(_edbff)
		if _dedb != nil {
			return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0067e\u0020\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_bfaeag.Rotate = &_dedde
	}
	if _cgfg := _gfggf.Get("\u0047\u0072\u006fu\u0070"); _cgfg != nil {
		_bfaeag.Group = _cgfg
	}
	if _ecfgc := _gfggf.Get("\u0054\u0068\u0075m\u0062"); _ecfgc != nil {
		_bfaeag.Thumb = _ecfgc
	}
	if _agafcc := _gfggf.Get("\u0042"); _agafcc != nil {
		_bfaeag.B = _agafcc
	}
	if _dgaga := _gfggf.Get("\u0044\u0075\u0072"); _dgaga != nil {
		_bfaeag.Dur = _dgaga
	}
	if _dbec := _gfggf.Get("\u0054\u0072\u0061n\u0073"); _dbec != nil {
		_bfaeag.Trans = _dbec
	}
	if _eddcb := _gfggf.Get("\u0041\u0041"); _eddcb != nil {
		_bfaeag.AA = _eddcb
	}
	if _daee := _gfggf.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _daee != nil {
		_bfaeag.Metadata = _daee
	}
	if _fbeee := _gfggf.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o"); _fbeee != nil {
		_bfaeag.PieceInfo = _fbeee
	}
	if _cgbd := _gfggf.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073"); _cgbd != nil {
		_bfaeag.StructParents = _cgbd
	}
	if _caef := _gfggf.Get("\u0049\u0044"); _caef != nil {
		_bfaeag.ID = _caef
	}
	if _abbce := _gfggf.Get("\u0050\u005a"); _abbce != nil {
		_bfaeag.PZ = _abbce
	}
	if _baeag := _gfggf.Get("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f"); _baeag != nil {
		_bfaeag.SeparationInfo = _baeag
	}
	if _bcedb := _gfggf.Get("\u0054\u0061\u0062\u0073"); _bcedb != nil {
		_bfaeag.Tabs = _bcedb
	}
	if _ceggbe := _gfggf.Get("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064"); _ceggbe != nil {
		_bfaeag.TemplateInstantiated = _ceggbe
	}
	if _bbeee := _gfggf.Get("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s"); _bbeee != nil {
		_bfaeag.PresSteps = _bbeee
	}
	if _gegf := _gfggf.Get("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074"); _gegf != nil {
		_bfaeag.UserUnit = _gegf
	}
	if _degc := _gfggf.Get("\u0056\u0050"); _degc != nil {
		_bfaeag.VP = _degc
	}
	if _fbdfg := _gfggf.Get("\u0041\u006e\u006e\u006f\u0074\u0073"); _fbdfg != nil {
		_bfaeag.Annots = _fbdfg
	}
	_bfaeag._gfcbe = _cfcgg
	return _bfaeag, nil
}
func _gaabf() string {
	_dccfe.Lock()
	defer _dccfe.Unlock()
	return _gbccbb
}

const (
	TrappedUnknown PdfInfoTrapped = "\u0055n\u006b\u006e\u006f\u0077\u006e"
	TrappedTrue    PdfInfoTrapped = "\u0054\u0072\u0075\u0065"
	TrappedFalse   PdfInfoTrapped = "\u0046\u0061\u006cs\u0065"
)

// ToPdfObject returns a PDF object representation of the outline item.
func (_eadgc *OutlineItem) ToPdfObject() _cde.PdfObject {
	_fcacf, _ := _eadgc.ToPdfOutlineItem()
	return _fcacf.ToPdfObject()
}

// PdfAnnotationCircle represents Circle annotations.
// (Section 12.5.6.8).
type PdfAnnotationCircle struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _cde.PdfObject
	IC _cde.PdfObject
	BE _cde.PdfObject
	RD _cde.PdfObject
}

var _gfac = map[string]struct{}{"\u0046\u0054": {}, "\u004b\u0069\u0064\u0073": {}, "\u0054": {}, "\u0054\u0055": {}, "\u0054\u004d": {}, "\u0046\u0066": {}, "\u0056": {}, "\u0044\u0056": {}, "\u0041\u0041": {}, "\u0044\u0041": {}, "\u0051": {}, "\u0044\u0053": {}, "\u0052\u0056": {}}

func (_acba *PdfReader) newPdfAnnotationPolyLineFromDict(_ddcc *_cde.PdfObjectDictionary) (*PdfAnnotationPolyLine, error) {
	_abde := PdfAnnotationPolyLine{}
	_gfaad, _aada := _acba.newPdfAnnotationMarkupFromDict(_ddcc)
	if _aada != nil {
		return nil, _aada
	}
	_abde.PdfAnnotationMarkup = _gfaad
	_abde.Vertices = _ddcc.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_abde.LE = _ddcc.Get("\u004c\u0045")
	_abde.BS = _ddcc.Get("\u0042\u0053")
	_abde.IC = _ddcc.Get("\u0049\u0043")
	_abde.BE = _ddcc.Get("\u0042\u0045")
	_abde.IT = _ddcc.Get("\u0049\u0054")
	_abde.Measure = _ddcc.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_abde, nil
}

// PdfAnnotationScreen represents Screen annotations.
// (Section 12.5.6.18).
type PdfAnnotationScreen struct {
	*PdfAnnotation
	T  _cde.PdfObject
	MK _cde.PdfObject
	A  _cde.PdfObject
	AA _cde.PdfObject
}

// Initialize initializes the PdfSignature.
func (_dfebb *PdfSignature) Initialize() error {
	if _dfebb.Handler == nil {
		return _ceg.New("\u0073\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0068\u0061\u006e\u0064\u006c\u0065r\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	return _dfebb.Handler.InitSignature(_dfebb)
}

// PdfColorspaceDeviceNAttributes contains additional information about the components of colour space that
// conforming readers may use. Conforming readers need not use the alternateSpace and tintTransform parameters,
// and may instead use a custom blending algorithms, along with other information provided in the attributes
// dictionary if present.
type PdfColorspaceDeviceNAttributes struct {
	Subtype     *_cde.PdfObjectName
	Colorants   _cde.PdfObject
	Process     _cde.PdfObject
	MixingHints _cde.PdfObject
	_deaf       *_cde.PdfIndirectObject
}

// Decrypt decrypts the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful,
// false otherwise.
func (_eeeec *PdfReader) Decrypt(password []byte) (bool, error) {
	_egcbg, _cedebe := _eeeec._aggcgb.Decrypt(password)
	if _cedebe != nil {
		return false, _cedebe
	}
	if !_egcbg {
		return false, nil
	}
	_cedebe = _eeeec.loadStructure()
	if _cedebe != nil {
		_ad.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f \u006co\u0061d\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", _cedebe)
		return false, _cedebe
	}
	return true, nil
}

// Encoder returns the font's text encoder.
func (_bfbfg *pdfFontSimple) Encoder() _gc.TextEncoder {
	if _bfbfg._efeaf != nil {
		return _bfbfg._efeaf
	}
	if _bfbfg._facga != nil {
		return _bfbfg._facga
	}
	_dccc, _ := _gc.NewSimpleTextEncoder("\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", nil)
	return _dccc
}

// ColorAt returns the color of the image pixel specified by the x and y coordinates.
func (_agaac *Image) ColorAt(x, y int) (_be.Color, error) {
	_gdgbb := _ff.BytesPerLine(int(_agaac.Width), int(_agaac.BitsPerComponent), _agaac.ColorComponents)
	switch _agaac.ColorComponents {
	case 1:
		return _ff.ColorAtGrayscale(x, y, int(_agaac.BitsPerComponent), _gdgbb, _agaac.Data, _agaac._aaafb)
	case 3:
		return _ff.ColorAtNRGBA(x, y, int(_agaac.Width), _gdgbb, int(_agaac.BitsPerComponent), _agaac.Data, _agaac._deegf, _agaac._aaafb)
	case 4:
		return _ff.ColorAtCMYK(x, y, int(_agaac.Width), _agaac.Data, _agaac._aaafb)
	}
	_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 i\u006da\u0067\u0065\u002e\u0020\u0025\u0064\u0020\u0063\u006f\u006d\u0070\u006fn\u0065\u006e\u0074\u0073\u002c\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072 \u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _agaac.ColorComponents, _agaac.BitsPerComponent)
	return nil, _ceg.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061g\u0065 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
}

// GetContainingPdfObject returns the container of the DSS (indirect object).
func (_aacd *DSS) GetContainingPdfObject() _cde.PdfObject { return _aacd._gbgda }

// RemovePage removes a page by number.
func (_geff *PdfAppender) RemovePage(pageNum int) {
	_bcaa := pageNum - 1
	_geff._gfb = append(_geff._gfb[0:_bcaa], _geff._gfb[pageNum:]...)
}

// ToPdfObject implements interface PdfModel.
func (_bddf *PdfActionRendition) ToPdfObject() _cde.PdfObject {
	_bddf.PdfAction.ToPdfObject()
	_aba := _bddf._bc
	_afe := _aba.PdfObject.(*_cde.PdfObjectDictionary)
	_afe.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeRendition)))
	_afe.SetIfNotNil("\u0052", _bddf.R)
	_afe.SetIfNotNil("\u0041\u004e", _bddf.AN)
	_afe.SetIfNotNil("\u004f\u0050", _bddf.OP)
	_afe.SetIfNotNil("\u004a\u0053", _bddf.JS)
	return _aba
}

// ColorToRGB verifies that the input color is an RGB color. Method exists in
// order to satisfy the PdfColorspace interface.
func (_cabe *PdfColorspaceDeviceRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ccgc, _fdec := color.(*PdfColorDeviceRGB)
	if !_fdec {
		_ad.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072 \u006e\u006f\u0074\u0020\u0064\u0065\u0076\u0069\u0063\u0065 \u0052\u0047\u0042")
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return _ccgc, nil
}

// AddOutlineTree adds outlines to a PDF file.
func (_gbfba *PdfWriter) AddOutlineTree(outlineTree *PdfOutlineTreeNode) { _gbfba._bdfce = outlineTree }

// NewPdfColorspaceSpecialSeparation returns a new separation color.
func NewPdfColorspaceSpecialSeparation() *PdfColorspaceSpecialSeparation {
	_effc := &PdfColorspaceSpecialSeparation{}
	return _effc
}
func (_fefe *PdfReader) newPdfAnnotationHighlightFromDict(_fbgf *_cde.PdfObjectDictionary) (*PdfAnnotationHighlight, error) {
	_fgfc := PdfAnnotationHighlight{}
	_aggf, _dcb := _fefe.newPdfAnnotationMarkupFromDict(_fbgf)
	if _dcb != nil {
		return nil, _dcb
	}
	_fgfc.PdfAnnotationMarkup = _aggf
	_fgfc.QuadPoints = _fbgf.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_fgfc, nil
}
func (_gfda *PdfReader) flattenFieldsWithOpts(_gfge bool, _bdfe FieldAppearanceGenerator, _beeeg *FieldFlattenOpts) error {
	if _beeeg == nil {
		_beeeg = &FieldFlattenOpts{}
	}
	var _bdaa bool
	_gcgf := map[*PdfAnnotation]bool{}
	{
		var _acfg []*PdfField
		_gagbd := _gfda.AcroForm
		if _gagbd != nil {
			if _beeeg.FilterFunc != nil {
				_acfg = _gagbd.filteredFields(_beeeg.FilterFunc, true)
				_bdaa = _gagbd.Fields != nil && len(*_gagbd.Fields) > 0
			} else {
				_acfg = _gagbd.AllFields()
			}
		}
		for _, _bdea := range _acfg {
			for _, _dccd := range _bdea.Annotations {
				_gcgf[_dccd.PdfAnnotation] = _bdea.V != nil
				if _bdfe != nil {
					_adfee, _fgbga := _bdfe.GenerateAppearanceDict(_gagbd, _bdea, _dccd)
					if _fgbga != nil {
						return _fgbga
					}
					_dccd.AP = _adfee
				}
			}
		}
	}
	if _gfge {
		for _, _bbeef := range _gfda.PageList {
			_cebg, _beab := _bbeef.GetAnnotations()
			if _beab != nil {
				return _beab
			}
			for _, _dgdfc := range _cebg {
				_gcgf[_dgdfc] = true
			}
		}
	}
	for _, _eacf := range _gfda.PageList {
		var _ebdc []*PdfAnnotation
		if _bdfe != nil {
			if _aaea := _bdfe.WrapContentStream(_eacf); _aaea != nil {
				return _aaea
			}
		}
		_dfeab, _fbabf := _eacf.GetAnnotations()
		if _fbabf != nil {
			return _fbabf
		}
		for _, _dabe := range _dfeab {
			_cege, _gdfe := _gcgf[_dabe]
			if !_gdfe && _beeeg.AnnotFilterFunc != nil {
				if _, _eaac := _dabe.GetContext().(*PdfAnnotationWidget); !_eaac {
					_gdfe = _beeeg.AnnotFilterFunc(_dabe)
				}
			}
			if !_gdfe {
				_ebdc = append(_ebdc, _dabe)
				continue
			}
			switch _dabe.GetContext().(type) {
			case *PdfAnnotationPopup:
				continue
			case *PdfAnnotationLink:
				continue
			case *PdfAnnotationProjection:
				continue
			}
			_faegc, _edfa, _beccf := _egadg(_dabe)
			if _beccf != nil {
				if !_cege {
					_ad.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0056\u0020\u002d\u003e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075t\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065am\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072")
					continue
				}
				_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d,\u0020\u0065\u0072\u0072\u0020\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006bi\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072", _beccf)
				continue
			}
			if _faegc == nil {
				continue
			}
			_efba := _eacf.Resources.GenerateXObjectName()
			_eacf.Resources.SetXObjectFormByName(_efba, _faegc)
			_abbe, _beccf := _cdcca(_faegc)
			if _beccf != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0061\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u004d\u0061\u0074\u0072\u0069\u0078\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0078\u0066\u006f\u0072\u006d\u0020\u0062\u0062\u006f\u0078\u0020\u0061\u0064\u006a\u0075\u0073t\u006d\u0065\u006e\u0074\u003a \u0025\u0076", _beccf)
			} else {
				_cbec := _adb.IdentityMatrix()
				_cbec = _cbec.Translate(-_abbe.Llx, -_abbe.Lly)
				_cbec = _cbec.Scale(_edfa.Width()/_abbe.Width(), _edfa.Height()/_abbe.Height())
				_edfa.Transform(_cbec)
			}
			_aedf := _ced.Min(_edfa.Llx, _edfa.Urx)
			_cacgg := _ced.Min(_edfa.Lly, _edfa.Ury)
			var _gdbfd []string
			_gdbfd = append(_gdbfd, "\u0071")
			_gdbfd = append(_gdbfd, _ee.Sprintf("\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020%\u002e6\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0063\u006d", 1.0, 0.0, 0.0, 1.0, _aedf, _cacgg))
			_gdbfd = append(_gdbfd, _ee.Sprintf("\u002f\u0025\u0073\u0020\u0044\u006f", _efba.String()))
			_gdbfd = append(_gdbfd, "\u0051")
			_edef := _dac.Join(_gdbfd, "\u000a")
			_beccf = _eacf.AppendContentStream(_edef)
			if _beccf != nil {
				return _beccf
			}
			if _faegc.Resources != nil {
				_ffag, _gade := _cde.GetDict(_faegc.Resources.Font)
				if _gade {
					for _, _edfe := range _ffag.Keys() {
						if !_eacf.Resources.HasFontByName(_edfe) {
							_eacf.Resources.SetFontByName(_edfe, _ffag.Get(_edfe))
						}
					}
				}
			}
		}
		if len(_ebdc) > 0 {
			_eacf._cefe = _ebdc
		} else {
			_eacf._cefe = []*PdfAnnotation{}
		}
	}
	if !_bdaa {
		_gfda.AcroForm = nil
	}
	return nil
}

// NewPdfFontFromPdfObject loads a PdfFont from the dictionary `fontObj`.  If there is a problem an
// error is returned.
func NewPdfFontFromPdfObject(fontObj _cde.PdfObject) (*PdfFont, error) { return _dgefe(fontObj, true) }

// PdfAppender appends new PDF content to an existing PDF document via incremental updates.
type PdfAppender struct {
	_aac   _f.ReadSeeker
	_gdgf  *_cde.PdfParser
	_cfad  *PdfReader
	Reader *PdfReader
	_gfb   []*PdfPage
	_cbeaa *PdfAcroForm
	_ffac  *DSS
	_bfbe  *Permissions
	_eccee _cde.XrefTable
	_aabb  int64
	_afcea int
	_bfec  []_cde.PdfObject
	_aadd  map[_cde.PdfObject]struct{}
	_adfa  map[_cde.PdfObject]int64
	_deb   map[_cde.PdfObject]struct{}
	_egbc  map[_cde.PdfObject]struct{}
	_efag  int64
	_gcbc  bool
	_dgdff string
	_fcga  *EncryptOptions
	_gdaa  *PdfInfo
}

func _ccfa(_dcac *PdfField, _cdda _cde.PdfObject) {
	for _, _gaee := range _dcac.Annotations {
		_gaee.AS = _cdda
		_gaee.ToPdfObject()
	}
}

// PdfColor interface represents a generic color in PDF.
type PdfColor interface{}

// SetContentStream updates the content stream with specified encoding.
// If encoding is null, will use the xform.Filter object or Raw encoding if not set.
func (_gfdeg *XObjectForm) SetContentStream(content []byte, encoder _cde.StreamEncoder) error {
	_fefcc := content
	if encoder == nil {
		if _gfdeg.Filter != nil {
			encoder = _gfdeg.Filter
		} else {
			encoder = _cde.NewRawEncoder()
		}
	}
	_efdbg, _eeffg := encoder.EncodeBytes(_fefcc)
	if _eeffg != nil {
		return _eeffg
	}
	_fefcc = _efdbg
	_gfdeg.Stream = _fefcc
	_gfdeg.Filter = encoder
	return nil
}

// GetCatalogStructTreeRoot gets the catalog StructTreeRoot object.
func (_bcdad *PdfReader) GetCatalogStructTreeRoot() (_cde.PdfObject, bool) {
	if _bcdad._efabe == nil {
		return nil, false
	}
	_fdcce := _bcdad._efabe.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074")
	return _fdcce, _fdcce != nil
}
func _gcfa(_bbac *_cde.PdfIndirectObject, _dcbe *_cde.PdfObjectDictionary) (*DSS, error) {
	if _bbac == nil {
		_bbac = _cde.MakeIndirectObject(nil)
	}
	_bbac.PdfObject = _cde.MakeDict()
	_agbef := map[string]*VRI{}
	if _bcba, _gead := _cde.GetDict(_dcbe.Get("\u0056\u0052\u0049")); _gead {
		for _, _deca := range _bcba.Keys() {
			if _cede, _dbbg := _cde.GetDict(_bcba.Get(_deca)); _dbbg {
				_agbef[_dac.ToUpper(_deca.String())] = _ecg(_cede)
			}
		}
	}
	return &DSS{Certs: _ebbec(_dcbe.Get("\u0043\u0065\u0072t\u0073")), OCSPs: _ebbec(_dcbe.Get("\u004f\u0043\u0053P\u0073")), CRLs: _ebbec(_dcbe.Get("\u0043\u0052\u004c\u0073")), VRI: _agbef, _gbgda: _bbac}, nil
}

// String returns a human readable description of `fontfile`.
func (_caccd *fontFile) String() string {
	_gdaad := "\u005b\u004e\u006f\u006e\u0065\u005d"
	if _caccd._cafea != nil {
		_gdaad = _caccd._cafea.String()
	}
	return _ee.Sprintf("\u0046O\u004e\u0054\u0046\u0049\u004c\u0045\u007b\u0025\u0023\u0071\u0020e\u006e\u0063\u006f\u0064\u0065\u0072\u003d\u0025\u0073\u007d", _caccd._fbffa, _gdaad)
}

// PdfBorderStyle represents a border style dictionary (12.5.4 Border Styles p. 394).
type PdfBorderStyle struct {
	W     *float64
	S     *BorderStyle
	D     *[]int
	_dggc _cde.PdfObject
}

func _fdce(_bgege _cde.PdfObject, _gdfbf *PdfReader) (*OutlineDest, error) {
	_gadfc, _gega := _cde.GetArray(_bgege)
	if !_gega {
		return nil, _ceg.New("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0064\u0065\u0073t\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_bfacg := _gadfc.Len()
	if _bfacg < 2 {
		return nil, _ee.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u006fu\u0074\u006c\u0069n\u0065\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061ti\u006f\u006e\u0020a\u0072\u0072a\u0079\u0020\u006c\u0065\u006e\u0067t\u0068\u003a \u0025\u0064", _bfacg)
	}
	_bgeead := &OutlineDest{Mode: "\u0046\u0069\u0074"}
	_debfc := _gadfc.Get(0)
	if _dfgba, _cagge := _cde.GetIndirect(_debfc); _cagge {
		if _, _abgdb, _gefbe := _gdfbf.PageFromIndirectObject(_dfgba); _gefbe == nil {
			_bgeead.Page = int64(_abgdb - 1)
		} else {
			_ad.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020g\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u002b\u0076", _dfgba)
		}
		_bgeead.PageObj = _dfgba
	} else if _dedec, _dgfdg := _cde.GetIntVal(_debfc); _dgfdg {
		if _dedec >= 0 && _dedec < len(_gdfbf.PageList) {
			_bgeead.PageObj = _gdfbf.PageList[_dedec].GetPageAsIndirectObject()
		} else {
			_ad.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _dedec)
		}
		_bgeead.Page = int64(_dedec)
	} else {
		return nil, _ee.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006f\u0075\u0074\u006cine\u0020de\u0073\u0074\u0069\u006e\u0061\u0074\u0069on\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0054", _debfc)
	}
	_eeafd, _gega := _cde.GetNameVal(_gadfc.Get(1))
	if !_gega {
		_ad.Log.Debug("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006d\u0061\u0067\u006e\u0069\u0066\u0069\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _gadfc.Get(1))
		return _bgeead, nil
	}
	switch _eeafd {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		if _bfacg > 2 {
			_bgeead.Y, _ = _cde.GetNumberAsFloat(_cde.TraceToDirectObject(_gadfc.Get(2)))
		}
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		if _bfacg > 2 {
			_bgeead.X, _ = _cde.GetNumberAsFloat(_cde.TraceToDirectObject(_gadfc.Get(2)))
		}
	case "\u0058\u0059\u005a":
		if _bfacg > 4 {
			_bgeead.X, _ = _cde.GetNumberAsFloat(_cde.TraceToDirectObject(_gadfc.Get(2)))
			_bgeead.Y, _ = _cde.GetNumberAsFloat(_cde.TraceToDirectObject(_gadfc.Get(3)))
			_bgeead.Zoom, _ = _cde.GetNumberAsFloat(_cde.TraceToDirectObject(_gadfc.Get(4)))
		}
	default:
		_eeafd = "\u0046\u0069\u0074"
	}
	_bgeead.Mode = _eeafd
	return _bgeead, nil
}
func _decfc(_addbb *_cde.PdfObjectDictionary, _cfagg *fontCommon) (*pdfFontType3, error) {
	_aedgg := _deafd(_cfagg)
	_aaff := _addbb.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
	if _aaff == nil {
		_aaff = _cde.MakeInteger(0)
	}
	_aedgg.FirstChar = _aaff
	_fagfg, _faecdf := _cde.GetIntVal(_aaff)
	if !_faecdf {
		_ad.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _aaff)
		return nil, _cde.ErrTypeError
	}
	_gaefg := _gc.CharCode(_fagfg)
	_aaff = _addbb.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
	if _aaff == nil {
		_aaff = _cde.MakeInteger(255)
	}
	_aedgg.LastChar = _aaff
	_fagfg, _faecdf = _cde.GetIntVal(_aaff)
	if !_faecdf {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _aaff)
		return nil, _cde.ErrTypeError
	}
	_cbfe := _gc.CharCode(_fagfg)
	_aaff = _addbb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _aaff != nil {
		_aedgg.Resources = _aaff
	}
	_aaff = _addbb.Get("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s")
	if _aaff == nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0068\u0061\u0072\u0050\u0072\u006f\u0063\u0073\u0020(%\u0076\u0029", _aaff)
		return nil, _cde.ErrNotSupported
	}
	_aedgg.CharProcs = _aaff
	_aaff = _addbb.Get("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078")
	if _aaff == nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0046\u006f\u006et\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0028\u0025\u0076\u0029", _aaff)
		return nil, _cde.ErrNotSupported
	}
	_aedgg.FontMatrix = _aaff
	_aedgg._dedg = make(map[_gc.CharCode]float64)
	_aaff = _addbb.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
	if _aaff != nil {
		_aedgg.Widths = _aaff
		_bfcda, _gfagg := _cde.GetArray(_aaff)
		if !_gfagg {
			_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _aaff)
			return nil, _cde.ErrTypeError
		}
		_dgbca, _ffeac := _bfcda.ToFloat64Array()
		if _ffeac != nil {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _ffeac
		}
		if len(_dgbca) != int(_cbfe-_gaefg+1) {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _cbfe-_gaefg+1, len(_dgbca))
			return nil, _cde.ErrRangeError
		}
		_bbcdb, _gfagg := _cde.GetArray(_aedgg.FontMatrix)
		if !_gfagg {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _bbcdb)
			return nil, _ffeac
		}
		_caga, _ffeac := _bbcdb.ToFloat64Array()
		if _ffeac != nil {
			_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020c\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0046o\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _ffeac
		}
		_bfdba := _adb.NewMatrix(_caga[0], _caga[1], _caga[2], _caga[3], _caga[4], _caga[5])
		for _ecde, _efde := range _dgbca {
			_abcd, _ := _bfdba.Transform(_efde, _efde)
			_aedgg._dedg[_gaefg+_gc.CharCode(_ecde)] = _abcd
		}
	}
	_aedgg.Encoding = _cde.TraceToDirectObject(_addbb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	_fbceg := _addbb.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _fbceg != nil {
		_aedgg._dfae = _cde.TraceToDirectObject(_fbceg)
		_eaef, _bebg := _bebbe(_aedgg._dfae, &_aedgg.fontCommon)
		if _bebg != nil {
			return nil, _bebg
		}
		_aedgg._ggebg = _eaef
	}
	if _geabb := _aedgg._ggebg; _geabb != nil {
		_aedgg._bdcege = _gc.NewCMapEncoder("", nil, _geabb)
	} else {
		_aedgg._bdcege = _gc.NewPdfDocEncoder()
	}
	return _aedgg, nil
}

var (
	_dfcde = _cc.MustCompile("\u005cd\u002b\u0020\u0064\u0069c\u0074\u005c\u0073\u002b\u0028d\u0075p\u005cs\u002b\u0029\u003f\u0062\u0065\u0067\u0069n")
	_fafdd = _cc.MustCompile("\u005e\u005cs\u002a\u002f\u0028\u005c\u0053\u002b\u003f\u0029\u005c\u0073\u002b\u0028\u002e\u002b\u003f\u0029\u005c\u0073\u002b\u0064\u0065\u0066\\s\u002a\u0024")
	_eadff = _cc.MustCompile("\u005e\u005c\u0073*\u0064\u0075\u0070\u005c\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002a\u002f\u0028\u005c\u0077\u002b\u003f\u0029\u0028\u003f\u003a\u005c\u002e\u005c\u0064\u002b)\u003f\u005c\u0073\u002b\u0070\u0075\u0074\u0024")
	_dcgd  = "\u002f\u0045\u006e\u0063od\u0069\u006e\u0067\u0020\u0032\u0035\u0036\u0020\u0061\u0072\u0072\u0061\u0079"
	_fafec = "\u0072\u0065\u0061d\u006f\u006e\u006c\u0079\u0020\u0064\u0065\u0066"
	_dadf  = "\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u0066\u0069\u006c\u0065\u0020e\u0065\u0078\u0065\u0063"
)

// SetImage updates XObject Image with new image data.
func (_gbced *XObjectImage) SetImage(img *Image, cs PdfColorspace) error {
	_gbced.Filter.UpdateParams(img.GetParamsDict())
	_aeead, _efbfg := _gbced.Filter.EncodeBytes(img.Data)
	if _efbfg != nil {
		return _efbfg
	}
	_gbced.Stream = _aeead
	_cdbde := img.Width
	_gbced.Width = &_cdbde
	_fbbe := img.Height
	_gbced.Height = &_fbbe
	_fgebg := img.BitsPerComponent
	_gbced.BitsPerComponent = &_fgebg
	if cs == nil {
		if img.ColorComponents == 1 {
			_gbced.ColorSpace = NewPdfColorspaceDeviceGray()
		} else if img.ColorComponents == 3 {
			_gbced.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			_gbced.ColorSpace = NewPdfColorspaceDeviceCMYK()
		} else {
			return _ceg.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_gbced.ColorSpace = cs
	}
	return nil
}
func _gfce(_aeaf _cde.PdfObject) (*PdfColorspaceICCBased, error) {
	_addg := &PdfColorspaceICCBased{}
	if _fed, _bfbed := _aeaf.(*_cde.PdfIndirectObject); _bfbed {
		_addg._cffb = _fed
	}
	_aeaf = _cde.TraceToDirectObject(_aeaf)
	_ffff, _eagg := _aeaf.(*_cde.PdfObjectArray)
	if !_eagg {
		return nil, _ee.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _ffff.Len() != 2 {
		return nil, _ee.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020c\u006f\u006c\u006fr\u0073p\u0061\u0063\u0065")
	}
	_aeaf = _cde.TraceToDirectObject(_ffff.Get(0))
	_fbbcf, _eagg := _aeaf.(*_cde.PdfObjectName)
	if !_eagg {
		return nil, _ee.Errorf("\u0049\u0043\u0043B\u0061\u0073\u0065\u0064 \u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_fbbcf != "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064" {
		return nil, _ee.Errorf("\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0049\u0043\u0043\u0042a\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_aeaf = _ffff.Get(1)
	_fgbgd, _eagg := _cde.GetStream(_aeaf)
	if !_eagg {
		_ad.Log.Error("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006e\u006f\u0074\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067\u0020t\u006f\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020%\u0054", _aeaf)
		return nil, _ee.Errorf("\u0049\u0043\u0043Ba\u0073\u0065\u0064\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_ffeb := _fgbgd.PdfObjectDictionary
	_cfcc, _eagg := _ffeb.Get("\u004e").(*_cde.PdfObjectInteger)
	if !_eagg {
		return nil, _ee.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006d\u0069\u0073\u0073\u0069n\u0067\u0020\u004e\u0020\u0066\u0072\u006fm\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074")
	}
	if *_cfcc != 1 && *_cfcc != 3 && *_cfcc != 4 {
		return nil, _ee.Errorf("\u0049\u0043\u0043\u0042\u0061s\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u006e\u006f\u0074\u0020\u0031\u002c\u0033\u002c\u0034\u0029")
	}
	_addg.N = int(*_cfcc)
	if _fcef := _ffeb.Get("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e"); _fcef != nil {
		_egde, _gddaf := NewPdfColorspaceFromPdfObject(_fcef)
		if _gddaf != nil {
			return nil, _gddaf
		}
		_addg.Alternate = _egde
	}
	if _faec := _ffeb.Get("\u0052\u0061\u006eg\u0065"); _faec != nil {
		_faec = _cde.TraceToDirectObject(_faec)
		_baac, _bacc := _faec.(*_cde.PdfObjectArray)
		if !_bacc {
			return nil, _ee.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0061\u006e \u0061r\u0072\u0061\u0079")
		}
		if _baac.Len() != 2*_addg.N {
			return nil, _ee.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0052\u0061\u006e\u0067e\u0020\u0077\u0072\u006f\u006e\u0067 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073")
		}
		_ddda, _cgfd := _baac.GetAsFloat64Slice()
		if _cgfd != nil {
			return nil, _cgfd
		}
		_addg.Range = _ddda
	} else {
		_addg.Range = make([]float64, 2*_addg.N)
		for _gbgcf := 0; _gbgcf < _addg.N; _gbgcf++ {
			_addg.Range[2*_gbgcf] = 0.0
			_addg.Range[2*_gbgcf+1] = 1.0
		}
	}
	if _baec := _ffeb.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _baec != nil {
		_afecd, _cbbed := _baec.(*_cde.PdfObjectStream)
		if !_cbbed {
			return nil, _ee.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u004de\u0074\u0061\u0064\u0061\u0074\u0061\u0020n\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		_addg.Metadata = _afecd
	}
	_bgbae, _egbf := _cde.DecodeStream(_fgbgd)
	if _egbf != nil {
		return nil, _egbf
	}
	_addg.Data = _bgbae
	_addg._fdea = _fgbgd
	return _addg, nil
}

// NewPdfOutlineItem returns an initialized PdfOutlineItem.
func NewPdfOutlineItem() *PdfOutlineItem {
	_eaggf := &PdfOutlineItem{_ccbf: _cde.MakeIndirectObject(_cde.MakeDict())}
	_eaggf._fbeea = _eaggf
	return _eaggf
}

// SetType sets the field button's type.  Can be one of:
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
// This sets the field's flag appropriately.
func (_dabde *PdfFieldButton) SetType(btype ButtonType) {
	_faga := uint32(0)
	if _dabde.Ff != nil {
		_faga = uint32(*_dabde.Ff)
	}
	switch btype {
	case ButtonTypePush:
		_faga |= FieldFlagPushbutton.Mask()
	case ButtonTypeRadio:
		_faga |= FieldFlagRadio.Mask()
	}
	_dabde.Ff = _cde.MakeInteger(int64(_faga))
}

// ToPdfObject implements interface PdfModel.
func (_efbd *PdfActionGoTo3DView) ToPdfObject() _cde.PdfObject {
	_efbd.PdfAction.ToPdfObject()
	_dggb := _efbd._bc
	_bfb := _dggb.PdfObject.(*_cde.PdfObjectDictionary)
	_bfb.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeGoTo3DView)))
	_bfb.SetIfNotNil("\u0054\u0041", _efbd.TA)
	_bfb.SetIfNotNil("\u0056", _efbd.V)
	return _dggb
}

// NewReaderOpts generates a default `ReaderOpts` instance.
func NewReaderOpts() *ReaderOpts { return &ReaderOpts{Password: "", LazyLoad: true} }

// OutlineDest represents the destination of an outline item.
// It holds the page and the position on the page an outline item points to.
type OutlineDest struct {
	PageObj *_cde.PdfIndirectObject `json:"-"`
	Page    int64                   `json:"page"`
	Mode    string                  `json:"mode"`
	X       float64                 `json:"x"`
	Y       float64                 `json:"y"`
	Zoom    float64                 `json:"zoom"`
}

// PdfActionLaunch represents a launch action.
type PdfActionLaunch struct {
	*PdfAction
	F         *PdfFilespec
	Win       _cde.PdfObject
	Mac       _cde.PdfObject
	Unix      _cde.PdfObject
	NewWindow _cde.PdfObject
}

// GetCatalogMetadata gets the catalog defined XMP Metadata.
func (_deaff *PdfReader) GetCatalogMetadata() (_cde.PdfObject, bool) {
	if _deaff._efabe == nil {
		return nil, false
	}
	_cgegg := _deaff._efabe.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _cgegg, _cgegg != nil
}

// PdfFunctionType4 is a Postscript calculator functions.
type PdfFunctionType4 struct {
	Domain  []float64
	Range   []float64
	Program *_dg.PSProgram
	_bdfd   *_dg.PSExecutor
	_bfbd   []byte
	_eaded  *_cde.PdfObjectStream
}

func (_dfe *PdfReader) newPdfAnnotationProjectionFromDict(_ddbd *_cde.PdfObjectDictionary) (*PdfAnnotationProjection, error) {
	_cca := &PdfAnnotationProjection{}
	_agge, _dgbd := _dfe.newPdfAnnotationMarkupFromDict(_ddbd)
	if _dgbd != nil {
		return nil, _dgbd
	}
	_cca.PdfAnnotationMarkup = _agge
	return _cca, nil
}

// GetAsShadingPattern returns a shading pattern. Check with IsShading() prior to using this.
func (_bbfb *PdfPattern) GetAsShadingPattern() *PdfShadingPattern {
	return _bbfb._abddb.(*PdfShadingPattern)
}

// SetAnnotations sets the annotations list.
func (_ecadb *PdfPage) SetAnnotations(annotations []*PdfAnnotation) { _ecadb._cefe = annotations }

// LTV represents an LTV (Long-Term Validation) client. It is used to LTV
// enable signatures by adding validation and revocation data (certificate,
// OCSP and CRL information) to the DSS dictionary of a PDF document.
//
// LTV is added through the DSS by:
// - Adding certificates, OCSP and CRL information in the global scope of the
//   DSS. The global data is used for validating any of the signatures present
//   in the document.
// - Adding certificates, OCSP and CRL information for a single signature,
//   through an entry in the VRI dictionary of the DSS. The added data is used
//   for validating that particular signature only. This is the recommended
//   method for adding validation data for a signature. However, this is not
//   is not possible in the same revision the signature is applied. Validation
//   data for a signature is added based on the Contents entry of the signature,
//   which is known only after the revision is written. Even if the Contents
//   are known (e.g. when signing externally), updating the DSS at that point
//   would invalidate the calculated signature. As a result, if adding LTV
//   in the same revision is a requirement, use the first method.
//   See LTV.EnableChain.
// The client applies both methods, when possible.
//
// If `LTV.SkipExisting` is set to true (the default), validations are
// not added for signatures which are already present in the VRI entry of the
// document's DSS dictionary.
type LTV struct {

	// CertClient is the client used to retrieve certificates.
	CertClient *_bbe.CertClient

	// OCSPClient is the client used to retrieve OCSP validation information.
	OCSPClient *_bbe.OCSPClient

	// CRLClient is the client used to retrieve CRL validation information.
	CRLClient *_bbe.CRLClient

	// SkipExisting specifies whether existing signature validations
	// should be skipped.
	SkipExisting bool
	_ffab        *PdfAppender
	_geeeg       *DSS
}

// GetTrailer returns the PDF's trailer dictionary.
func (_gbbfda *PdfReader) GetTrailer() (*_cde.PdfObjectDictionary, error) {
	_begdg := _gbbfda._aggcgb.GetTrailer()
	if _begdg == nil {
		return nil, _ceg.New("\u0074r\u0061i\u006c\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	return _begdg, nil
}

// WatermarkImageOptions contains options for configuring the watermark process.
type WatermarkImageOptions struct {
	Alpha               float64
	FitToWidth          bool
	PreserveAspectRatio bool
}

// ToPdfObject implements interface PdfModel.
func (_cedg *PdfActionNamed) ToPdfObject() _cde.PdfObject {
	_cedg.PdfAction.ToPdfObject()
	_cedf := _cedg._bc
	_afa := _cedf.PdfObject.(*_cde.PdfObjectDictionary)
	_afa.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeNamed)))
	_afa.SetIfNotNil("\u004e", _cedg.N)
	return _cedf
}
func (_eecab *DSS) addCerts(_bbef [][]byte) ([]*_cde.PdfObjectStream, error) {
	return _eecab.add(&_eecab.Certs, _eecab._cddc, _bbef)
}

// NewOutlineBookmark returns an initialized PdfOutlineItem for a given bookmark title and page.
func NewOutlineBookmark(title string, page *_cde.PdfIndirectObject) *PdfOutlineItem {
	_deff := PdfOutlineItem{}
	_deff._fbeea = &_deff
	_deff.Title = _cde.MakeString(title)
	_cbged := _cde.MakeArray()
	_cbged.Append(page)
	_cbged.Append(_cde.MakeName("\u0046\u0069\u0074"))
	_deff.Dest = _cbged
	return &_deff
}

// HasFontByName checks if has font resource by name.
func (_bdgba *PdfPage) HasFontByName(name _cde.PdfObjectName) bool {
	_adeg, _bdaf := _bdgba.Resources.Font.(*_cde.PdfObjectDictionary)
	if !_bdaf {
		return false
	}
	if _ebbded := _adeg.Get(name); _ebbded != nil {
		return true
	}
	return false
}

// ToPdfObject implements model.PdfModel interface.
func (_ebbf *PdfOutputIntent) ToPdfObject() _cde.PdfObject {
	if _ebbf._acbcb == nil {
		_ebbf._acbcb = _cde.MakeDict()
	}
	_bdadf := _ebbf._acbcb
	if _ebbf.Type != "" {
		_bdadf.Set("\u0054\u0079\u0070\u0065", _cde.MakeName(_ebbf.Type))
	}
	_bdadf.Set("\u0053", _cde.MakeName(_ebbf.S.String()))
	if _ebbf.OutputCondition != "" {
		_bdadf.Set("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e", _cde.MakeString(_ebbf.OutputCondition))
	}
	_bdadf.Set("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r", _cde.MakeString(_ebbf.OutputConditionIdentifier))
	_bdadf.Set("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065", _cde.MakeString(_ebbf.RegistryName))
	if _ebbf.Info != "" {
		_bdadf.Set("\u0049\u006e\u0066\u006f", _cde.MakeString(_ebbf.Info))
	}
	if len(_ebbf.DestOutputProfile) != 0 {
		_adcca, _gdfgd := _cde.MakeStream(_ebbf.DestOutputProfile, _cde.NewFlateEncoder())
		if _gdfgd != nil {
			_ad.Log.Error("\u004d\u0061\u006b\u0065\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0044\u0065s\u0074\u004f\u0075\u0074\u0070\u0075t\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _gdfgd)
		}
		_adcca.PdfObjectDictionary.Set("\u004e", _cde.MakeInteger(int64(_ebbf.ColorComponents)))
		_gaaeb := make([]float64, _ebbf.ColorComponents*2)
		for _acdfc := 0; _acdfc < _ebbf.ColorComponents*2; _acdfc++ {
			_gcceb := 0.0
			if _acdfc%2 != 0 {
				_gcceb = 1.0
			}
			_gaaeb[_acdfc] = _gcceb
		}
		_adcca.PdfObjectDictionary.Set("\u0052\u0061\u006eg\u0065", _cde.MakeArrayFromFloats(_gaaeb))
		_bdadf.Set("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065", _adcca)
	}
	return _bdadf
}
func (_fbcee *PdfWriter) setDocumentIDs(_bfgga, _agdgd string) {
	_fbcee._bgacb = _cde.MakeArray(_cde.MakeHexString(_bfgga), _cde.MakeHexString(_agdgd))
}

// PdfTransformParamsDocMDP represents a transform parameters dictionary for the DocMDP method and is used to detect
// modifications relative to a signature field that is signed by the author of a document.
// (Section 12.8.2.2, Table 254 - Entries in the DocMDP transform parameters dictionary p. 471 in PDF32000_2008).
type PdfTransformParamsDocMDP struct {
	Type *_cde.PdfObjectName
	P    *_cde.PdfObjectInteger
	V    *_cde.PdfObjectName
}

// GetExtGState gets the ExtGState specified by keyName. Returns a bool
// indicating whether it was found or not.
func (_cfceeg *PdfPageResources) GetExtGState(keyName _cde.PdfObjectName) (_cde.PdfObject, bool) {
	if _cfceeg.ExtGState == nil {
		return nil, false
	}
	_baeaf, _ceeca := _cde.TraceToDirectObject(_cfceeg.ExtGState).(*_cde.PdfObjectDictionary)
	if !_ceeca {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cfceeg.ExtGState)
		return nil, false
	}
	if _edddf := _baeaf.Get(keyName); _edddf != nil {
		return _edddf, true
	}
	return nil, false
}

// ToPdfObject implements interface PdfModel.
func (_fgdg *PdfAnnotationSound) ToPdfObject() _cde.PdfObject {
	_fgdg.PdfAnnotation.ToPdfObject()
	_affb := _fgdg._bddg
	_aag := _affb.PdfObject.(*_cde.PdfObjectDictionary)
	_fgdg.PdfAnnotationMarkup.appendToPdfDictionary(_aag)
	_aag.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0053\u006f\u0075n\u0064"))
	_aag.SetIfNotNil("\u0053\u006f\u0075n\u0064", _fgdg.Sound)
	_aag.SetIfNotNil("\u004e\u0061\u006d\u0065", _fgdg.Name)
	return _affb
}

// ToPdfObject implements interface PdfModel.
func (_cacc *PdfAnnotationHighlight) ToPdfObject() _cde.PdfObject {
	_cacc.PdfAnnotation.ToPdfObject()
	_beeda := _cacc._bddg
	_acbc := _beeda.PdfObject.(*_cde.PdfObjectDictionary)
	_cacc.PdfAnnotationMarkup.appendToPdfDictionary(_acbc)
	_acbc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t"))
	_acbc.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _cacc.QuadPoints)
	return _beeda
}

// ToPdfObject implements interface PdfModel.
func (_ac *PdfActionHide) ToPdfObject() _cde.PdfObject {
	_ac.PdfAction.ToPdfObject()
	_dea := _ac._bc
	_gfc := _dea.PdfObject.(*_cde.PdfObjectDictionary)
	_gfc.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeHide)))
	_gfc.SetIfNotNil("\u0054", _ac.T)
	_gfc.SetIfNotNil("\u0048", _ac.H)
	return _dea
}

// NewPdfAnnotationScreen returns a new screen annotation.
func NewPdfAnnotationScreen() *PdfAnnotationScreen {
	_dcg := NewPdfAnnotation()
	_cdgc := &PdfAnnotationScreen{}
	_cdgc.PdfAnnotation = _dcg
	_dcg.SetContext(_cdgc)
	return _cdgc
}

// PdfAnnotation represents an annotation in PDF (section 12.5 p. 389).
type PdfAnnotation struct {
	_bea         PdfModel
	Rect         _cde.PdfObject
	Contents     _cde.PdfObject
	P            _cde.PdfObject
	NM           _cde.PdfObject
	M            _cde.PdfObject
	F            _cde.PdfObject
	AP           _cde.PdfObject
	AS           _cde.PdfObject
	Border       _cde.PdfObject
	C            _cde.PdfObject
	StructParent _cde.PdfObject
	OC           _cde.PdfObject
	_bddg        *_cde.PdfIndirectObject
}

// MergePageWith appends page content to source Pdf file page content.
func (_gac *PdfAppender) MergePageWith(pageNum int, page *PdfPage) error {
	_debb := pageNum - 1
	var _fga *PdfPage
	for _ddaa, _efdc := range _gac._gfb {
		if _ddaa == _debb {
			_fga = _efdc
		}
	}
	if _fga == nil {
		return _ee.Errorf("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073o\u0075\u0072\u0063\u0065\u0020\u0064o\u0063\u0075\u006de\u006e\u0074", pageNum)
	}
	if _fga._dcaeff != nil && _fga._dcaeff.GetParser() == _gac._cfad._aggcgb {
		_fga = _fga.Duplicate()
		_gac._gfb[_debb] = _fga
	}
	page = page.Duplicate()
	_fdcef(page)
	_cfdd := _eegc(_fga)
	_ebb := _eegc(page)
	_fdgd := make(map[_cde.PdfObjectName]_cde.PdfObjectName)
	for _gcef := range _ebb {
		if _, _gaabb := _cfdd[_gcef]; _gaabb {
			for _bbeb := 1; true; _bbeb++ {
				_dffb := _cde.PdfObjectName(string(_gcef) + _gb.Itoa(_bbeb))
				if _, _gfgfe := _cfdd[_dffb]; !_gfgfe {
					_fdgd[_gcef] = _dffb
					break
				}
			}
		}
	}
	_daf, _cgba := page.GetContentStreams()
	if _cgba != nil {
		return _cgba
	}
	_gegg, _cgba := _fga.GetContentStreams()
	if _cgba != nil {
		return _cgba
	}
	for _fbf, _dded := range _daf {
		for _adee, _deg := range _fdgd {
			_dded = _dac.Replace(_dded, "\u002f"+string(_adee), "\u002f"+string(_deg), -1)
		}
		_daf[_fbf] = _dded
	}
	_gegg = append(_gegg, _daf...)
	if _dfaf := _fga.SetContentStreams(_gegg, _cde.NewFlateEncoder()); _dfaf != nil {
		return _dfaf
	}
	_fga._cefe = append(_fga._cefe, page._cefe...)
	if _fga.Resources == nil {
		_fga.Resources = NewPdfPageResources()
	}
	if page.Resources != nil {
		_fga.Resources.Font = _gac.mergeResources(_fga.Resources.Font, page.Resources.Font, _fdgd)
		_fga.Resources.XObject = _gac.mergeResources(_fga.Resources.XObject, page.Resources.XObject, _fdgd)
		_fga.Resources.Properties = _gac.mergeResources(_fga.Resources.Properties, page.Resources.Properties, _fdgd)
		if _fga.Resources.ProcSet == nil {
			_fga.Resources.ProcSet = page.Resources.ProcSet
		}
		_fga.Resources.Shading = _gac.mergeResources(_fga.Resources.Shading, page.Resources.Shading, _fdgd)
		_fga.Resources.ExtGState = _gac.mergeResources(_fga.Resources.ExtGState, page.Resources.ExtGState, _fdgd)
	}
	_dbce, _cgba := _fga.GetMediaBox()
	if _cgba != nil {
		return _cgba
	}
	_gbab, _cgba := page.GetMediaBox()
	if _cgba != nil {
		return _cgba
	}
	var _eggc bool
	if _dbce.Llx > _gbab.Llx {
		_dbce.Llx = _gbab.Llx
		_eggc = true
	}
	if _dbce.Lly > _gbab.Lly {
		_dbce.Lly = _gbab.Lly
		_eggc = true
	}
	if _dbce.Urx < _gbab.Urx {
		_dbce.Urx = _gbab.Urx
		_eggc = true
	}
	if _dbce.Ury < _gbab.Ury {
		_dbce.Ury = _gbab.Ury
		_eggc = true
	}
	if _eggc {
		_fga.MediaBox = _dbce
	}
	return nil
}

// HasFontByName checks whether a font is defined by the specified keyName.
func (_aefdg *PdfPageResources) HasFontByName(keyName _cde.PdfObjectName) bool {
	_, _agbaaa := _aefdg.GetFontByName(keyName)
	return _agbaaa
}

// Width returns the width of `rect`.
func (_fdfc *PdfRectangle) Width() float64 { return _ced.Abs(_fdfc.Urx - _fdfc.Llx) }

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_cgeda *PdfColorspaceSpecialSeparation) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_edaac := vals[0]
	_bgbca := []float64{_edaac}
	_fggbg, _dbced := _cgeda.TintTransform.Evaluate(_bgbca)
	if _dbced != nil {
		_ad.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0065: \u0025\u0076", _dbced)
		_ad.Log.Trace("\u0054\u0069\u006e\u0074 t\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _cgeda.TintTransform)
		return nil, _dbced
	}
	_ad.Log.Trace("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0043\u006f\u006c\u006fr\u0046\u0072\u006f\u006d\u0046\u006c\u006f\u0061\u0074\u0073\u0028\u0025\u002bv\u0029\u0020\u006f\u006e\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061te\u0053\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0023\u0076", _fggbg, _cgeda.AlternateSpace)
	_cbge, _dbced := _cgeda.AlternateSpace.ColorFromFloats(_fggbg)
	if _dbced != nil {
		_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u002c\u0020\u0066a\u0069\u006c\u0065d \u0074\u006f\u0020\u0065\u0076\u0061l\u0075\u0061\u0074\u0065\u0020\u0069\u006e\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u003a \u0025\u0076", _dbced)
		return nil, _dbced
	}
	return _cbge, nil
}

// NewPdfColorDeviceGray returns a new grayscale color based on an input grayscale float value in range [0-1].
func NewPdfColorDeviceGray(grayVal float64) *PdfColorDeviceGray {
	_gdcf := PdfColorDeviceGray(grayVal)
	return &_gdcf
}

// A returns the value of the A component of the color.
func (_efgc *PdfColorLab) A() float64 { return _efgc[1] }

// Register registers (caches) a model to primitive object relationship.
func (_fabeg *modelManager) Register(primitive _cde.PdfObject, model PdfModel) {
	_fabeg._bfaea[model] = primitive
	_fabeg._dfeee[primitive] = model
}

// NewPdfAnnotationCaret returns a new caret annotation.
func NewPdfAnnotationCaret() *PdfAnnotationCaret {
	_fcdf := NewPdfAnnotation()
	_dcab := &PdfAnnotationCaret{}
	_dcab.PdfAnnotation = _fcdf
	_dcab.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fcdf.SetContext(_dcab)
	return _dcab
}
func (_dbdg *PdfReader) newPdfAnnotationSquigglyFromDict(_eeee *_cde.PdfObjectDictionary) (*PdfAnnotationSquiggly, error) {
	_bge := PdfAnnotationSquiggly{}
	_cgca, _aea := _dbdg.newPdfAnnotationMarkupFromDict(_eeee)
	if _aea != nil {
		return nil, _aea
	}
	_bge.PdfAnnotationMarkup = _cgca
	_bge.QuadPoints = _eeee.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_bge, nil
}

// ImageToRGB converts ICCBased colorspace image to RGB and returns the result.
func (_bdcc *PdfColorspaceICCBased) ImageToRGB(img Image) (Image, error) {
	if _bdcc.Alternate == nil {
		_ad.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _bdcc.N == 1 {
			_ad.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_gccf := NewPdfColorspaceDeviceGray()
			return _gccf.ImageToRGB(img)
		} else if _bdcc.N == 3 {
			_ad.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return img, nil
		} else if _bdcc.N == 4 {
			_ad.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_fgdgb := NewPdfColorspaceDeviceCMYK()
			return _fgdgb.ImageToRGB(img)
		} else {
			return img, _ceg.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_ad.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _bdcc)
	_eddf, _aeda := _bdcc.Alternate.ImageToRGB(img)
	_ad.Log.Trace("I\u0043C\u0020\u0049\u006e\u0070\u0075\u0074\u0020\u0069m\u0061\u0067\u0065\u003a %\u002b\u0076", img)
	_ad.Log.Trace("I\u0043\u0043\u0020\u004fut\u0070u\u0074\u0020\u0069\u006d\u0061g\u0065\u003a\u0020\u0025\u002b\u0076", _eddf)
	return _eddf, _aeda
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_cgdbg pdfCIDFontType2) GetRuneMetrics(r rune) (_fe.CharMetrics, bool) {
	_fdgdd, _cfea := _cgdbg._dfefe[r]
	if !_cfea {
		_fbabb, _eggd := _cde.GetInt(_cgdbg.DW)
		if !_eggd {
			return _fe.CharMetrics{}, false
		}
		_fdgdd = int(*_fbabb)
	}
	return _fe.CharMetrics{Wx: float64(_fdgdd)}, true
}

// PdfAnnotationSound represents Sound annotations.
// (Section 12.5.6.16).
type PdfAnnotationSound struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Sound _cde.PdfObject
	Name  _cde.PdfObject
}

// BorderStyle defines border type, typically used for annotations.
type BorderStyle int

func (_bfbcg *pdfCIDFontType2) baseFields() *fontCommon { return &_bfbcg.fontCommon }

// AcroFormRepairOptions contains options for rebuilding the AcroForm.
type AcroFormRepairOptions struct{}

// ColorToRGB converts a DeviceN color to an RGB color.
func (_gdebd *PdfColorspaceDeviceN) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _gdebd.AlternateSpace == nil {
		return nil, _ceg.New("\u0044\u0065\u0076\u0069\u0063\u0065N\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _gdebd.AlternateSpace.ColorToRGB(color)
}
func (_dbegba *PdfReader) newPdfSignatureFromIndirect(_cadffa *_cde.PdfIndirectObject) (*PdfSignature, error) {
	_ddbad, _efagd := _cadffa.PdfObject.(*_cde.PdfObjectDictionary)
	if !_efagd {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		return nil, ErrTypeCheck
	}
	if _caeba, _bcbbd := _dbegba._bedfa.GetModelFromPrimitive(_cadffa).(*PdfSignature); _bcbbd {
		return _caeba, nil
	}
	_afbgb := &PdfSignature{}
	_afbgb._cabd = _cadffa
	_afbgb.Type, _ = _cde.GetName(_ddbad.Get("\u0054\u0079\u0070\u0065"))
	_afbgb.Filter, _efagd = _cde.GetName(_ddbad.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if !_efagd {
		_ad.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053i\u0067\u006e\u0061\u0074\u0075r\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_afbgb.SubFilter, _ = _cde.GetName(_ddbad.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r"))
	_afbgb.Contents, _efagd = _cde.GetString(_ddbad.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_efagd {
		_ad.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	if _ecff, _adbde := _cde.GetArray(_ddbad.Get("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e")); _adbde {
		_afbgb.Reference = _cde.MakeArray()
		for _, _eebcb := range _ecff.Elements() {
			_agedd, _cbdfd := _cde.GetDict(_eebcb)
			if !_cbdfd {
				_ad.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020R\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0061\u0074\u0065\u0064")
				return nil, ErrInvalidAttribute
			}
			_afcgec, _gbgbd := _dbegba.newPdfSignatureReferenceFromDict(_agedd)
			if _gbgbd != nil {
				return nil, _gbgbd
			}
			_afbgb.Reference.Append(_afcgec.ToPdfObject())
		}
	}
	_afbgb.Cert = _ddbad.Get("\u0043\u0065\u0072\u0074")
	_afbgb.ByteRange, _ = _cde.GetArray(_ddbad.Get("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e"))
	_afbgb.Changes, _ = _cde.GetArray(_ddbad.Get("\u0043h\u0061\u006e\u0067\u0065\u0073"))
	_afbgb.Name, _ = _cde.GetString(_ddbad.Get("\u004e\u0061\u006d\u0065"))
	_afbgb.M, _ = _cde.GetString(_ddbad.Get("\u004d"))
	_afbgb.Location, _ = _cde.GetString(_ddbad.Get("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"))
	_afbgb.Reason, _ = _cde.GetString(_ddbad.Get("\u0052\u0065\u0061\u0073\u006f\u006e"))
	_afbgb.ContactInfo, _ = _cde.GetString(_ddbad.Get("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f"))
	_afbgb.R, _ = _cde.GetInt(_ddbad.Get("\u0052"))
	_afbgb.V, _ = _cde.GetInt(_ddbad.Get("\u0056"))
	_afbgb.PropBuild, _ = _cde.GetDict(_ddbad.Get("\u0050\u0072\u006f\u0070\u005f\u0042\u0075\u0069\u006c\u0064"))
	_afbgb.PropAuthTime, _ = _cde.GetInt(_ddbad.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0069\u006d\u0065"))
	_afbgb.PropAuthType, _ = _cde.GetName(_ddbad.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0079\u0070\u0065"))
	_dbegba._bedfa.Register(_cadffa, _afbgb)
	return _afbgb, nil
}

// NewPdfOutline returns an initialized PdfOutline.
func NewPdfOutline() *PdfOutline {
	_ebffd := &PdfOutline{_dgcdc: _cde.MakeIndirectObject(_cde.MakeDict())}
	_ebffd._fbeea = _ebffd
	return _ebffd
}
func (_fegfc *PdfColorspaceSpecialSeparation) String() string {
	return "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"
}

// ToPdfObject converts the font to a PDF representation.
func (_dabgc *pdfFontType3) ToPdfObject() _cde.PdfObject {
	if _dabgc._fddbc == nil {
		_dabgc._fddbc = &_cde.PdfIndirectObject{}
	}
	_bcedg := _dabgc.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0033")
	_dabgc._fddbc.PdfObject = _bcedg
	if _dabgc.FirstChar != nil {
		_bcedg.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _dabgc.FirstChar)
	}
	if _dabgc.LastChar != nil {
		_bcedg.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _dabgc.LastChar)
	}
	if _dabgc.Widths != nil {
		_bcedg.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _dabgc.Widths)
	}
	if _dabgc.Encoding != nil {
		_bcedg.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _dabgc.Encoding)
	} else if _dabgc._bdcege != nil {
		_gffb := _dabgc._bdcege.ToPdfObject()
		if _gffb != nil {
			_bcedg.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _gffb)
		}
	}
	if _dabgc.FontBBox != nil {
		_bcedg.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _dabgc.FontBBox)
	}
	if _dabgc.FontMatrix != nil {
		_bcedg.Set("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0069\u0072\u0078", _dabgc.FontMatrix)
	}
	if _dabgc.CharProcs != nil {
		_bcedg.Set("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s", _dabgc.CharProcs)
	}
	if _dabgc.Resources != nil {
		_bcedg.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _dabgc.Resources)
	}
	return _dabgc._fddbc
}

// NewPdfAnnotationSquiggly returns a new text squiggly annotation.
func NewPdfAnnotationSquiggly() *PdfAnnotationSquiggly {
	_ddga := NewPdfAnnotation()
	_fcd := &PdfAnnotationSquiggly{}
	_fcd.PdfAnnotation = _ddga
	_fcd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ddga.SetContext(_fcd)
	return _fcd
}

// ToWriter creates a new writer from the current reader, based on the specified options.
// If no options are provided, all reader properties are copied to the writer.
func (_deggf *PdfReader) ToWriter(opts *ReaderToWriterOpts) (*PdfWriter, error) {
	_cadgg := NewPdfWriter()
	if opts == nil {
		opts = &ReaderToWriterOpts{}
	}
	_eagf, _bdbc := _deggf.GetNumPages()
	if _bdbc != nil {
		_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bdbc)
		return nil, _bdbc
	}
	for _gebd := 1; _gebd <= _eagf; _gebd++ {
		_dbgd, _bcdfe := _deggf.GetPage(_gebd)
		if _bcdfe != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bcdfe)
			return nil, _bcdfe
		}
		if opts.PageProcessCallback != nil {
			_bcdfe = opts.PageProcessCallback(_gebd, _dbgd)
			if _bcdfe != nil {
				_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bcdfe)
				return nil, _bcdfe
			}
		} else if opts.PageCallback != nil {
			opts.PageCallback(_gebd, _dbgd)
		}
		_bcdfe = _cadgg.AddPage(_dbgd)
		if _bcdfe != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bcdfe)
			return nil, _bcdfe
		}
	}
	_cadgg._cgdcc = _deggf.PdfVersion()
	if !opts.SkipInfo {
		_bcebbb, _cccgc := _deggf.GetPdfInfo()
		if _cccgc != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cccgc)
		} else {
			_cadgg._fdgbc.PdfObject = _bcebbb.ToPdfObject()
		}
	}
	if !opts.SkipMetadata {
		if _cbagd := _deggf._efabe.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _cbagd != nil {
			if _facgae := _cadgg.SetCatalogMetadata(_cbagd); _facgae != nil {
				return nil, _facgae
			}
		}
	}
	if !opts.SkipAcroForm {
		_aeggeb := _cadgg.SetForms(_deggf.AcroForm)
		if _aeggeb != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aeggeb)
			return nil, _aeggeb
		}
	}
	if !opts.SkipOutlines {
		_cadgg.AddOutlineTree(_deggf.GetOutlineTree())
	}
	if !opts.SkipOCProperties {
		_bbdegd, _bcgga := _deggf.GetOCProperties()
		if _bcgga != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bcgga)
		} else {
			_bcgga = _cadgg.SetOCProperties(_bbdegd)
			if _bcgga != nil {
				_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bcgga)
			}
		}
	}
	if !opts.SkipPageLabels {
		_gccaa, _cgedae := _deggf.GetPageLabels()
		if _cgedae != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cgedae)
		} else {
			_cgedae = _cadgg.SetPageLabels(_gccaa)
			if _cgedae != nil {
				_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cgedae)
			}
		}
	}
	if !opts.SkipNamedDests {
		_dfcgd, _dgea := _deggf.GetNamedDestinations()
		if _dgea != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgea)
		} else {
			_dgea = _cadgg.SetNamedDestinations(_dfcgd)
			if _dgea != nil {
				_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgea)
			}
		}
	}
	if !opts.SkipNameDictionary {
		_cdagd, _gabfg := _deggf.GetNameDictionary()
		if _gabfg != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gabfg)
		} else {
			_gabfg = _cadgg.SetNameDictionary(_cdagd)
			if _gabfg != nil {
				_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gabfg)
			}
		}
	}
	if !opts.SkipRotation && _deggf.Rotate != nil {
		if _fbffae := _cadgg.SetRotation(*_deggf.Rotate); _fbffae != nil {
			_ad.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fbffae)
		}
	}
	return &_cadgg, nil
}

// GetNumComponents returns the number of color components (3 for CalRGB).
func (_ecaff *PdfColorCalRGB) GetNumComponents() int { return 3 }

// ToPdfObject implements interface PdfModel.
func (_fbbc *PdfAnnotationRedact) ToPdfObject() _cde.PdfObject {
	_fbbc.PdfAnnotation.ToPdfObject()
	_cafe := _fbbc._bddg
	_cegc := _cafe.PdfObject.(*_cde.PdfObjectDictionary)
	_fbbc.PdfAnnotationMarkup.appendToPdfDictionary(_cegc)
	_cegc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0052\u0065\u0064\u0061\u0063\u0074"))
	_cegc.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _fbbc.QuadPoints)
	_cegc.SetIfNotNil("\u0049\u0043", _fbbc.IC)
	_cegc.SetIfNotNil("\u0052\u004f", _fbbc.RO)
	_cegc.SetIfNotNil("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074", _fbbc.OverlayText)
	_cegc.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _fbbc.Repeat)
	_cegc.SetIfNotNil("\u0044\u0041", _fbbc.DA)
	_cegc.SetIfNotNil("\u0051", _fbbc.Q)
	return _cafe
}

// Compress is yet to be implemented.
// Should be able to compress in terms of JPEG quality parameter,
// and DPI threshold (need to know bounding area dimensions).
func (_dacd DefaultImageHandler) Compress(input *Image, quality int64) (*Image, error) {
	return input, nil
}

// GenerateXObjectName generates an unused XObject name that can be used for
// adding new XObjects. Uses format XObj1, XObj2, ...
func (_abda *PdfPageResources) GenerateXObjectName() _cde.PdfObjectName {
	_aedde := 1
	for {
		_gece := _cde.MakeName(_ee.Sprintf("\u0058\u004f\u0062\u006a\u0025\u0064", _aedde))
		if !_abda.HasXObjectByName(*_gece) {
			return *_gece
		}
		_aedde++
	}
}

// PdfActionGoTo3DView represents a GoTo3DView action.
type PdfActionGoTo3DView struct {
	*PdfAction
	TA _cde.PdfObject
	V  _cde.PdfObject
}

func (_gaed *PdfReader) newPdfAnnotationStampFromDict(_cebd *_cde.PdfObjectDictionary) (*PdfAnnotationStamp, error) {
	_add := PdfAnnotationStamp{}
	_acg, _adfe := _gaed.newPdfAnnotationMarkupFromDict(_cebd)
	if _adfe != nil {
		return nil, _adfe
	}
	_add.PdfAnnotationMarkup = _acg
	_add.Name = _cebd.Get("\u004e\u0061\u006d\u0065")
	return &_add, nil
}

// WriteToFile writes the output PDF to file.
func (_gbcgb *PdfWriter) WriteToFile(outputFilePath string) error {
	_fgaab, _cgafba := _db.Create(outputFilePath)
	if _cgafba != nil {
		return _cgafba
	}
	defer _fgaab.Close()
	return _gbcgb.Write(_fgaab)
}

// PdfDate represents a date, which is a PDF string of the form:
// (D:YYYYMMDDHHmmSSOHH'mm)
type PdfDate struct {
	_aedbfg int64
	_babce  int64
	_edbef  int64
	_aeac   int64
	_accgb  int64
	_ebfeg  int64
	_efdde  byte
	_gafaf  int64
	_bbcgcf int64
}

// NewPdfColorDeviceCMYK returns a new CMYK32 color.
func NewPdfColorDeviceCMYK(c, m, y, k float64) *PdfColorDeviceCMYK {
	_caeb := PdfColorDeviceCMYK{c, m, y, k}
	return &_caeb
}
func (_egdbcf *PdfWriter) writeBytes(_cegge []byte) {
	if _egdbcf._fggeef != nil {
		return
	}
	_fabcc, _gdefe := _egdbcf._cefdd.Write(_cegge)
	_egdbcf._eddbc += int64(_fabcc)
	_egdbcf._fggeef = _gdefe
}

// PdfAnnotationPolygon represents Polygon annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolygon struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _cde.PdfObject
	LE       _cde.PdfObject
	BS       _cde.PdfObject
	IC       _cde.PdfObject
	BE       _cde.PdfObject
	IT       _cde.PdfObject
	Measure  _cde.PdfObject
}

// GetOptimizer returns current PDF optimizer.
func (_aabbd *PdfWriter) GetOptimizer() Optimizer { return _aabbd._cgee }
func _agad(_efdcc *fontCommon) *pdfCIDFontType0   { return &pdfCIDFontType0{fontCommon: *_efdcc} }

// FieldFlag represents form field flags. Some of the flags can apply to all types of fields whereas other
// flags are specific.
type FieldFlag uint32

func (_bedd *PdfReader) loadPerms() (*Permissions, error) {
	if _efbcd := _bedd._efabe.Get("\u0050\u0065\u0072m\u0073"); _efbcd != nil {
		if _bfcdc, _adbc := _cde.GetDict(_efbcd); _adbc {
			_aabfc := _bfcdc.Get("\u0044\u006f\u0063\u004d\u0044\u0050")
			if _aabfc == nil {
				return nil, nil
			}
			if _aaefb, _facgaef := _cde.GetIndirect(_aabfc); _facgaef {
				_ceagc, _edbcac := _bedd.newPdfSignatureFromIndirect(_aaefb)
				if _edbcac != nil {
					return nil, _edbcac
				}
				return NewPermissions(_ceagc), nil
			}
			return nil, _ee.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u006f\u0063M\u0044\u0050\u0020\u0065nt\u0072\u0079")
		}
		return nil, _ee.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0050\u0065\u0072\u006d\u0073\u0020\u0065\u006e\u0074\u0072\u0079")
	}
	return nil, nil
}

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

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 4 for a CMYK32 device.
func (_dabc *PdfColorspaceDeviceCMYK) GetNumComponents() int { return 4 }

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a Lab device.
func (_fddb *PdfColorspaceLab) GetNumComponents() int { return 3 }
func _eafbf(_eaga _cde.PdfObject) (*PdfColorspaceSpecialIndexed, error) {
	_gcgdf := NewPdfColorspaceSpecialIndexed()
	if _dcaa, _accb := _eaga.(*_cde.PdfIndirectObject); _accb {
		_gcgdf._gcbcb = _dcaa
	}
	_eaga = _cde.TraceToDirectObject(_eaga)
	_gbdc, _egeg := _eaga.(*_cde.PdfObjectArray)
	if !_egeg {
		return nil, _ee.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _gbdc.Len() != 4 {
		return nil, _ee.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_eaga = _gbdc.Get(0)
	_febbc, _egeg := _eaga.(*_cde.PdfObjectName)
	if !_egeg {
		return nil, _ee.Errorf("\u0069n\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
	}
	if *_febbc != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _ee.Errorf("\u0069\u006e\u0064\u0065xe\u0064\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u006e\u0061m\u0065")
	}
	_eaga = _gbdc.Get(1)
	_afgbc, _cabb := DetermineColorspaceNameFromPdfObject(_eaga)
	if _cabb != nil {
		return nil, _cabb
	}
	if _afgbc == "\u0049n\u0064\u0065\u0078\u0065\u0064" || _afgbc == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
		_ad.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0049\u006e\u0064e\u0078\u0065\u0064\u002f\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043S\u0020\u0061\u0073\u0020\u0062\u0061\u0073\u0065\u0020\u0028\u0025v\u0029", _afgbc)
		return nil, _cdab
	}
	_decc, _cabb := NewPdfColorspaceFromPdfObject(_eaga)
	if _cabb != nil {
		return nil, _cabb
	}
	_gcgdf.Base = _decc
	_eaga = _gbdc.Get(2)
	_gafe, _cabb := _cde.GetNumberAsInt64(_eaga)
	if _cabb != nil {
		return nil, _cabb
	}
	if _gafe > 255 {
		return nil, _ee.Errorf("\u0069n\u0064\u0065\u0078\u0065d\u0020\u0043\u0053\u003a\u0020I\u006ev\u0061l\u0069\u0064\u0020\u0068\u0069\u0076\u0061l")
	}
	_gcgdf.HiVal = int(_gafe)
	_eaga = _gbdc.Get(3)
	_gcgdf.Lookup = _eaga
	_eaga = _cde.TraceToDirectObject(_eaga)
	var _dccba []byte
	if _efef, _eadg := _eaga.(*_cde.PdfObjectString); _eadg {
		_dccba = _efef.Bytes()
		_ad.Log.Trace("\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0020\u0064", _dccba)
	} else if _dbda, _fdcc := _eaga.(*_cde.PdfObjectStream); _fdcc {
		_ad.Log.Trace("\u0049n\u0064e\u0078\u0065\u0064\u0020\u0073t\u0072\u0065a\u006d\u003a\u0020\u0025\u0073", _eaga.String())
		_ad.Log.Trace("\u0045\u006e\u0063\u006fde\u0064\u0020\u0028\u0025\u0064\u0029\u0020\u003a\u0020\u0025\u0023\u0020\u0078", len(_dbda.Stream), _dbda.Stream)
		_bfaf, _aeed := _cde.DecodeStream(_dbda)
		if _aeed != nil {
			return nil, _aeed
		}
		_ad.Log.Trace("\u0044e\u0063o\u0064\u0065\u0064\u0020\u0028%\u0064\u0029 \u003a\u0020\u0025\u0020\u0058", len(_bfaf), _bfaf)
		_dccba = _bfaf
	} else {
		_ad.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _eaga)
		return nil, _ee.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076a\u006ci\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if len(_dccba) < _gcgdf.Base.GetNumComponents()*(_gcgdf.HiVal+1) {
		_ad.Log.Debug("\u0050\u0044\u0046\u0020\u0049\u006e\u0063o\u006d\u0070\u0061t\u0069\u0062\u0069\u006ci\u0074\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074\u006f\u006f\u0020\u0073\u0068\u006f\u0072\u0074")
		_ad.Log.Debug("\u0046\u0061i\u006c\u002c\u0020\u006c\u0065\u006e\u0028\u0064\u0061\u0074\u0061\u0029\u003a\u0020\u0025\u0064\u002c\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u002c\u0020\u0068\u0069\u0056\u0061\u006c\u003a\u0020\u0025\u0064", len(_dccba), _gcgdf.Base.GetNumComponents(), _gcgdf.HiVal)
	} else {
		_dccba = _dccba[:_gcgdf.Base.GetNumComponents()*(_gcgdf.HiVal+1)]
	}
	_gcgdf._dafd = _dccba
	return _gcgdf, nil
}

// DetermineColorspaceNameFromPdfObject determines PDF colorspace from a PdfObject.  Returns the colorspace name and
// an error on failure. If the colorspace was not found, will return an empty string.
func DetermineColorspaceNameFromPdfObject(obj _cde.PdfObject) (_cde.PdfObjectName, error) {
	var _bcbf *_cde.PdfObjectName
	var _gbff *_cde.PdfObjectArray
	if _gdgaa, _ddea := obj.(*_cde.PdfIndirectObject); _ddea {
		if _gaca, _ecfe := _gdgaa.PdfObject.(*_cde.PdfObjectArray); _ecfe {
			_gbff = _gaca
		} else if _bcaf, _edcb := _gdgaa.PdfObject.(*_cde.PdfObjectName); _edcb {
			_bcbf = _bcaf
		}
	} else if _cgaa, _ecfb := obj.(*_cde.PdfObjectArray); _ecfb {
		_gbff = _cgaa
	} else if _dddb, _gecg := obj.(*_cde.PdfObjectName); _gecg {
		_bcbf = _dddb
	}
	if _bcbf != nil {
		switch *_bcbf {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_bcbf, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_bcbf, nil
		}
	}
	if _gbff != nil && _gbff.Len() > 0 {
		if _fccf, _bfae := _gbff.Get(0).(*_cde.PdfObjectName); _bfae {
			switch *_fccf {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _gbff.Len() == 1 {
					return *_fccf, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_fccf, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_fccf, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_fccf, nil
			}
		}
	}
	return "", nil
}

// ToPdfObject converts date to a PDF string object.
func (_adegf *PdfDate) ToPdfObject() _cde.PdfObject {
	_gaffa := _ee.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _adegf._aedbfg, _adegf._babce, _adegf._edbef, _adegf._aeac, _adegf._accgb, _adegf._ebfeg, _adegf._efdde, _adegf._gafaf, _adegf._bbcgcf)
	return _cde.MakeString(_gaffa)
}

// Add appends an outline item as a child of the current outline item.
func (_caeg *OutlineItem) Add(item *OutlineItem) { _caeg.Entries = append(_caeg.Entries, item) }

// GetContentStreams returns the content stream as an array of strings.
func (_fdfbf *PdfPage) GetContentStreams() ([]string, error) {
	_ebecfe := _fdfbf.GetContentStreamObjs()
	var _abbd []string
	for _, _ceda := range _ebecfe {
		_affg, _daaeg := _fbggd(_ceda)
		if _daaeg != nil {
			return nil, _daaeg
		}
		_abbd = append(_abbd, _affg)
	}
	return _abbd, nil
}

// NewCompliancePdfReader creates a PdfReader or an input io.ReadSeeker that during reading will scan the files for the
// metadata details. It could be used for the PDF standard implementations like PDF/A or PDF/X.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
func NewCompliancePdfReader(rs _f.ReadSeeker) (*CompliancePdfReader, error) {
	const _eadda = "\u006d\u006f\u0064\u0065l\u003a\u004e\u0065\u0077\u0043\u006f\u006d\u0070\u006c\u0069a\u006ec\u0065\u0050\u0064\u0066\u0052\u0065\u0061d\u0065\u0072"
	_bafeg, _ffdcc := _gfceb(rs, &ReaderOpts{ComplianceMode: true}, false, _eadda)
	if _ffdcc != nil {
		return nil, _ffdcc
	}
	return &CompliancePdfReader{PdfReader: _bafeg}, nil
}
func _fecd(_gbfg _cde.PdfObject) (*PdfColorspaceDeviceN, error) {
	_cgeb := NewPdfColorspaceDeviceN()
	if _gdbg, _abcc := _gbfg.(*_cde.PdfIndirectObject); _abcc {
		_cgeb._cdac = _gdbg
	}
	_gbfg = _cde.TraceToDirectObject(_gbfg)
	_ccdg, _ffdg := _gbfg.(*_cde.PdfObjectArray)
	if !_ffdg {
		return nil, _ee.Errorf("\u0064\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if _ccdg.Len() != 4 && _ccdg.Len() != 5 {
		return nil, _ee.Errorf("\u0064\u0065\u0076ic\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0049\u006ec\u006fr\u0072e\u0063t\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_gbfg = _ccdg.Get(0)
	_dccf, _ffdg := _gbfg.(*_cde.PdfObjectName)
	if !_ffdg {
		return nil, _ee.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_dccf != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
		return nil, _ee.Errorf("\u0064\u0065v\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020na\u006d\u0065")
	}
	_gbfg = _ccdg.Get(1)
	_gbfg = _cde.TraceToDirectObject(_gbfg)
	_cgbbc, _ffdg := _gbfg.(*_cde.PdfObjectArray)
	if !_ffdg {
		return nil, _ee.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0049\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cgeb.ColorantNames = _cgbbc
	_gbfg = _ccdg.Get(2)
	_aedef, _aegge := NewPdfColorspaceFromPdfObject(_gbfg)
	if _aegge != nil {
		return nil, _aegge
	}
	_cgeb.AlternateSpace = _aedef
	_cgdf, _aegge := _cfdbb(_ccdg.Get(3))
	if _aegge != nil {
		return nil, _aegge
	}
	_cgeb.TintTransform = _cgdf
	if _ccdg.Len() == 5 {
		_bgga, _gecac := _ceba(_ccdg.Get(4))
		if _gecac != nil {
			return nil, _gecac
		}
		_cgeb.Attributes = _bgga
	}
	return _cgeb, nil
}

// IsRadio returns true if the button field represents a radio button, false otherwise.
func (_afff *PdfFieldButton) IsRadio() bool { return _afff.GetType() == ButtonTypeRadio }

// Encoder returns the font's text encoder.
func (_ccacd pdfFontType0) Encoder() _gc.TextEncoder { return _ccacd._cffef }
func _cgafb(_cffdb _cde.PdfObject) (*PdfPattern, error) {
	_dcegc := &PdfPattern{}
	var _gdfcd *_cde.PdfObjectDictionary
	if _fcgfc, _efbda := _cde.GetIndirect(_cffdb); _efbda {
		_dcegc._eecac = _fcgfc
		_cfcf, _fggfc := _fcgfc.PdfObject.(*_cde.PdfObjectDictionary)
		if !_fggfc {
			_ad.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0028g\u006f\u0074\u0020%\u0054\u0029", _fcgfc.PdfObject)
			return nil, _cde.ErrTypeError
		}
		_gdfcd = _cfcf
	} else if _dafff, _cgace := _cde.GetStream(_cffdb); _cgace {
		_dcegc._eecac = _dafff
		_gdfcd = _dafff.PdfObjectDictionary
	} else {
		_ad.Log.Debug("\u0050a\u0074\u0074e\u0072\u006e\u0020\u006eo\u0074\u0020\u0061n\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 o\u0062\u006a\u0065c\u0074\u0020o\u0072\u0020\u0073\u0074\u0072\u0065a\u006d\u002e \u0025\u0054", _cffdb)
		return nil, _cde.ErrTypeError
	}
	_fedc := _gdfcd.Get("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065")
	if _fedc == nil {
		_ad.Log.Debug("\u0050\u0064\u0066\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069n\u0067\u0020\u0050\u0061\u0074t\u0065\u0072n\u0054\u0079\u0070\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fddf, _eaab := _fedc.(*_cde.PdfObjectInteger)
	if !_eaab {
		_ad.Log.Debug("\u0050\u0061tt\u0065\u0072\u006e \u0074\u0079\u0070\u0065 no\u0074 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067er\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _fedc)
		return nil, _cde.ErrTypeError
	}
	if *_fddf != 1 && *_fddf != 2 {
		_ad.Log.Debug("\u0050\u0061\u0074\u0074e\u0072\u006e\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d \u0031/\u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", *_fddf)
		return nil, _cde.ErrRangeError
	}
	_dcegc.PatternType = int64(*_fddf)
	switch *_fddf {
	case 1:
		_eegaa, _fada := _dagcd(_gdfcd)
		if _fada != nil {
			return nil, _fada
		}
		_eegaa.PdfPattern = _dcegc
		_dcegc._abddb = _eegaa
		return _dcegc, nil
	case 2:
		_bcbdcb, _ebce := _fbbce(_gdfcd)
		if _ebce != nil {
			return nil, _ebce
		}
		_bcbdcb.PdfPattern = _dcegc
		_dcegc._abddb = _bcbdcb
		return _dcegc, nil
	}
	return nil, _ceg.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
}

// XObjectForm (Table 95 in 8.10.2).
type XObjectForm struct {
	Filter        _cde.StreamEncoder
	FormType      _cde.PdfObject
	BBox          _cde.PdfObject
	Matrix        _cde.PdfObject
	Resources     *PdfPageResources
	Group         _cde.PdfObject
	Ref           _cde.PdfObject
	MetaData      _cde.PdfObject
	PieceInfo     _cde.PdfObject
	LastModified  _cde.PdfObject
	StructParent  _cde.PdfObject
	StructParents _cde.PdfObject
	OPI           _cde.PdfObject
	OC            _cde.PdfObject
	Name          _cde.PdfObject

	// Stream data.
	Stream []byte
	_ecfbd *_cde.PdfObjectStream
}

func (_dbbdf *PdfReader) resolveReference(_gcdd *_cde.PdfObjectReference) (_cde.PdfObject, bool, error) {
	_gbacgb, _fcgeage := _dbbdf._aggcgb.ObjCache[int(_gcdd.ObjectNumber)]
	if !_fcgeage {
		_ad.Log.Trace("R\u0065\u0061\u0064\u0065r \u004co\u006f\u006b\u0075\u0070\u0020r\u0065\u0066\u003a\u0020\u0025\u0073", _gcdd)
		_addbbb, _ccdgd := _dbbdf._aggcgb.LookupByReference(*_gcdd)
		if _ccdgd != nil {
			return nil, false, _ccdgd
		}
		_dbbdf._aggcgb.ObjCache[int(_gcdd.ObjectNumber)] = _addbbb
		return _addbbb, false, nil
	}
	return _gbacgb, true, nil
}
func (_gdec *PdfAppender) replaceObject(_gdca, _fbc _cde.PdfObject) {
	switch _fce := _gdca.(type) {
	case *_cde.PdfIndirectObject:
		_gdec._adfa[_fbc] = _fce.ObjectNumber
	case *_cde.PdfObjectStream:
		_gdec._adfa[_fbc] = _fce.ObjectNumber
	}
}
func (_gbgcd *PdfReader) buildOutlineTree(_acbafd _cde.PdfObject, _fafda *PdfOutlineTreeNode, _bdfdd *PdfOutlineTreeNode, _dfffe map[_cde.PdfObject]struct{}) (*PdfOutlineTreeNode, *PdfOutlineTreeNode, error) {
	if _dfffe == nil {
		_dfffe = map[_cde.PdfObject]struct{}{}
	}
	_dfffe[_acbafd] = struct{}{}
	_cead, _acaf := _acbafd.(*_cde.PdfIndirectObject)
	if !_acaf {
		return nil, nil, _ee.Errorf("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0063\u006f\u006et\u0061\u0069\u006e\u0065r\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _acbafd)
	}
	_gfage, _edcf := _cead.PdfObject.(*_cde.PdfObjectDictionary)
	if !_edcf {
		return nil, nil, _ceg.New("\u006e\u006f\u0074 a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_ad.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0074\u0072\u0065\u0065\u003a\u0020d\u0069\u0063\u0074\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0076\u0029\u0020p\u003a\u0020\u0025\u0070", _gfage, _cead, _cead)
	if _aeecc := _gfage.Get("\u0054\u0069\u0074l\u0065"); _aeecc != nil {
		_fgbf, _eeeb := _gbgcd.newPdfOutlineItemFromIndirectObject(_cead)
		if _eeeb != nil {
			return nil, nil, _eeeb
		}
		_fgbf.Parent = _fafda
		_fgbf.Prev = _bdfdd
		_ebcgdb := _cde.ResolveReference(_gfage.Get("\u0046\u0069\u0072s\u0074"))
		if _, _aabcf := _dfffe[_ebcgdb]; _ebcgdb != nil && _ebcgdb != _cead && !_aabcf {
			if !_cde.IsNullObject(_ebcgdb) {
				_adefe, _dafg, _edagg := _gbgcd.buildOutlineTree(_ebcgdb, &_fgbf.PdfOutlineTreeNode, nil, _dfffe)
				if _edagg != nil {
					_ad.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0069\u0074\u0065\u006d\u0020\u0074\u0072\u0065\u0065\u003a \u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020n\u006f\u0064\u0065\u0020\u0063\u0068\u0069\u006c\u0064\u0072\u0065n\u002e", _edagg)
				} else {
					_fgbf.First = _adefe
					_fgbf.Last = _dafg
				}
			}
		}
		_aaaee := _cde.ResolveReference(_gfage.Get("\u004e\u0065\u0078\u0074"))
		if _, _debd := _dfffe[_aaaee]; _aaaee != nil && _aaaee != _cead && !_debd {
			if !_cde.IsNullObject(_aaaee) {
				_cdfb, _aefec, _edcbg := _gbgcd.buildOutlineTree(_aaaee, _fafda, &_fgbf.PdfOutlineTreeNode, _dfffe)
				if _edcbg != nil {
					_ad.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065\u0020\u0066\u006f\u0072\u0020\u004ee\u0078\u0074\u0020\u006e\u006f\u0064\u0065\u003a\u0020\u0025\u0076\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006e\u006f\u0064e\u002e", _edcbg)
				} else {
					_fgbf.Next = _cdfb
					return &_fgbf.PdfOutlineTreeNode, _aefec, nil
				}
			}
		}
		return &_fgbf.PdfOutlineTreeNode, &_fgbf.PdfOutlineTreeNode, nil
	}
	_fdae, _bfagg := _bfeg(_cead)
	if _bfagg != nil {
		return nil, nil, _bfagg
	}
	_fdae.Parent = _fafda
	if _bdffb := _gfage.Get("\u0046\u0069\u0072s\u0074"); _bdffb != nil {
		_bdffb = _cde.ResolveReference(_bdffb)
		if _, _gacf := _dfffe[_bdffb]; _bdffb != nil && _bdffb != _cead && !_gacf {
			_dccbaf := _cde.TraceToDirectObject(_bdffb)
			if _, _ecbd := _dccbaf.(*_cde.PdfObjectNull); !_ecbd && _dccbaf != nil {
				_aebbd, _dafagb, _cfbee := _gbgcd.buildOutlineTree(_bdffb, &_fdae.PdfOutlineTreeNode, nil, _dfffe)
				if _cfbee != nil {
					_ad.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020b\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0074\u0072\u0065\u0065\u003a\u0020\u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006e\u006f\u0064\u0065 \u0063\u0068i\u006c\u0064r\u0065n\u002e", _cfbee)
				} else {
					_fdae.First = _aebbd
					_fdae.Last = _dafagb
				}
			}
		}
	}
	return &_fdae.PdfOutlineTreeNode, &_fdae.PdfOutlineTreeNode, nil
}
func _dgdg() _ce.Time { _dccfe.Lock(); defer _dccfe.Unlock(); return _gffce }

// ToInteger convert to an integer format.
func (_aacc *PdfColorCalGray) ToInteger(bits int) uint32 {
	_fbge := _ced.Pow(2, float64(bits)) - 1
	return uint32(_fbge * _aacc.Val())
}

// SetDate sets the `M` field of the signature.
func (_gbcg *PdfSignature) SetDate(date _ce.Time, format string) {
	if format == "" {
		format = "\u0044\u003a\u003200\u0036\u0030\u0031\u0030\u0032\u0031\u0035\u0030\u0034\u0030\u0035\u002d\u0030\u0037\u0027\u0030\u0030\u0027"
	}
	_gbcg.M = _cde.MakeString(date.Format(format))
}

// SetNameDictionary sets the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_efgca *PdfWriter) SetNameDictionary(names _cde.PdfObject) error {
	if names == nil {
		return nil
	}
	_ad.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u004e\u0061\u006d\u0065\u0073\u002e\u002e\u002e")
	_efgca._fedbb.Set("\u004e\u0061\u006de\u0073", names)
	return _efgca.addObjects(names)
}

// PdfAnnotationUnderline represents Underline annotations.
// (Section 12.5.6.10).
type PdfAnnotationUnderline struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _cde.PdfObject
}

func (_cdgdc *PdfPage) getParentResources() (*PdfPageResources, error) {
	_bcafb := _cdgdc.Parent
	for _bcafb != nil {
		_agbbc, _gfcec := _cde.GetDict(_bcafb)
		if !_gfcec {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020n\u006f\u0064\u0065")
			return nil, _ceg.New("i\u006e\u0076\u0061\u006cid\u0020p\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		if _gfegda := _agbbc.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _gfegda != nil {
			_dcddc, _dfbee := _cde.GetDict(_gfegda)
			if !_dfbee {
				return nil, _ceg.New("i\u006e\u0076\u0061\u006cid\u0020r\u0065\u0073\u006f\u0075\u0072c\u0065\u0020\u0064\u0069\u0063\u0074")
			}
			_ebdca, _cadgd := NewPdfPageResourcesFromDict(_dcddc)
			if _cadgd != nil {
				return nil, _cadgd
			}
			return _ebdca, nil
		}
		_bcafb = _agbbc.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, nil
}

// GetContainingPdfObject returns the container of the PdfAcroForm (indirect object).
func (_gccfb *PdfAcroForm) GetContainingPdfObject() _cde.PdfObject { return _gccfb._ecca }

// ImageToRGB convert an indexed image to RGB.
func (_bfdb *PdfColorspaceSpecialIndexed) ImageToRGB(img Image) (Image, error) {
	N := _bfdb.Base.GetNumComponents()
	if N < 1 {
		return Image{}, _ee.Errorf("\u0062\u0061d \u0062\u0061\u0073e\u0020\u0063\u006f\u006cors\u0070ac\u0065\u0020\u004e\u0075\u006d\u0043\u006fmp\u006f\u006e\u0065\u006e\u0074\u0073\u003d%\u0064", N)
	}
	_ceea := _ff.NewImageBase(int(img.Width), int(img.Height), 8, N, nil, img._deegf, img._aaafb)
	_abcg := _cae.NewReader(img.getBase())
	_bdda := _cae.NewWriter(_ceea)
	var (
		_afab uint32
		_gabb int
		_badb error
	)
	for {
		_afab, _badb = _abcg.ReadSample()
		if _badb == _f.EOF {
			break
		} else if _badb != nil {
			return img, _badb
		}
		_gabb = int(_afab)
		_ad.Log.Trace("\u0049\u006ed\u0065\u0078\u0065\u0064\u003a\u0020\u0069\u006e\u0064\u0065\u0078\u003d\u0025\u0064\u0020\u004e\u003d\u0025\u0064\u0020\u006c\u0075t=\u0025\u0064", _gabb, N, len(_bfdb._dafd))
		if (_gabb+1)*N > len(_bfdb._dafd) {
			_gabb = len(_bfdb._dafd)/N - 1
			_ad.Log.Trace("C\u006c\u0069\u0070\u0070in\u0067 \u0074\u006f\u0020\u0069\u006ed\u0065\u0078\u003a\u0020\u0025\u0064", _gabb)
			if _gabb < 0 {
				_ad.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069p\u0020\u0069\u006e\u0064\u0065\u0078.\u0020\u0049\u0073\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006ce\u0020\u0064\u0061\u006d\u0061\u0067\u0065\u0064\u003f")
				break
			}
		}
		for _becf := _gabb * N; _becf < (_gabb+1)*N; _becf++ {
			if _badb = _bdda.WriteSample(uint32(_bfdb._dafd[_becf])); _badb != nil {
				return img, _badb
			}
		}
	}
	return _bfdb.Base.ImageToRGB(_bddb(&_ceea))
}

// Enable LTV enables the specified signature. The signing certificate
// chain is extracted from the signature dictionary. Optionally, additional
// certificates can be specified through the `extraCerts` parameter.
// The LTV client attempts to build the certificate chain up to a trusted root
// by downloading any missing certificates.
func (_cgdac *LTV) Enable(sig *PdfSignature, extraCerts []*_bg.Certificate) error {
	if _cefgg := _cgdac.validateSig(sig); _cefgg != nil {
		return _cefgg
	}
	_dcedc, _gcdab := _cgdac.generateVRIKey(sig)
	if _gcdab != nil {
		return _gcdab
	}
	if _, _ceeaa := _cgdac._geeeg.VRI[_dcedc]; _ceeaa && _cgdac.SkipExisting {
		return nil
	}
	_debfe, _gcdab := sig.GetCerts()
	if _gcdab != nil {
		return _gcdab
	}
	return _cgdac.enable(_debfe, extraCerts, _dcedc)
}
func _dddcf(_dgbf _cde.PdfObject) (*PdfPageResourcesColorspaces, error) {
	_bgged := &PdfPageResourcesColorspaces{}
	if _baeg, _bggd := _dgbf.(*_cde.PdfIndirectObject); _bggd {
		_bgged._ccdd = _baeg
		_dgbf = _baeg.PdfObject
	}
	_gfdbg, _bfcdb := _cde.GetDict(_dgbf)
	if !_bfcdb {
		return nil, _ceg.New("\u0043\u0053\u0020at\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_bgged.Names = []string{}
	_bgged.Colorspaces = map[string]PdfColorspace{}
	for _, _ffcge := range _gfdbg.Keys() {
		_dcgada := _gfdbg.Get(_ffcge)
		_bgged.Names = append(_bgged.Names, string(_ffcge))
		_bgef, _egdfc := NewPdfColorspaceFromPdfObject(_dcgada)
		if _egdfc != nil {
			return nil, _egdfc
		}
		_bgged.Colorspaces[string(_ffcge)] = _bgef
	}
	return _bgged, nil
}

// NewPdfDateFromTime will create a PdfDate based on the given time
func NewPdfDateFromTime(timeObj _ce.Time) (PdfDate, error) {
	_aefdd := timeObj.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_ccbgf, _ := _gb.ParseInt(_aefdd[1:3], 10, 32)
	_afgbcd, _ := _gb.ParseInt(_aefdd[4:6], 10, 32)
	return PdfDate{_aedbfg: int64(timeObj.Year()), _babce: int64(timeObj.Month()), _edbef: int64(timeObj.Day()), _aeac: int64(timeObj.Hour()), _accgb: int64(timeObj.Minute()), _ebfeg: int64(timeObj.Second()), _efdde: _aefdd[0], _gafaf: _ccbgf, _bbcgcf: _afgbcd}, nil
}

// SetColorSpace sets `r` colorspace object to `colorspace`.
func (_dbff *PdfPageResources) SetColorSpace(colorspace *PdfPageResourcesColorspaces) {
	_dbff._bfff = colorspace
}

// ToPdfObject returns an indirect object containing the signature field dictionary.
func (_gefge *PdfFieldSignature) ToPdfObject() _cde.PdfObject {
	if _gefge.PdfAnnotationWidget != nil {
		_gefge.PdfAnnotationWidget.ToPdfObject()
	}
	_gefge.PdfField.ToPdfObject()
	_fbba := _gefge._afgc
	_bdee := _fbba.PdfObject.(*_cde.PdfObjectDictionary)
	_bdee.SetIfNotNil("\u0046\u0054", _cde.MakeName("\u0053\u0069\u0067"))
	_bdee.SetIfNotNil("\u004c\u006f\u0063\u006b", _gefge.Lock)
	_bdee.SetIfNotNil("\u0053\u0056", _gefge.SV)
	if _gefge.V != nil {
		_bdee.SetIfNotNil("\u0056", _gefge.V.ToPdfObject())
	}
	return _fbba
}

// PartialName returns the partial name of the field.
func (_dggaf *PdfField) PartialName() string {
	_bdad := ""
	if _dggaf.T != nil {
		_bdad = _dggaf.T.Decoded()
	} else {
		_ad.Log.Debug("\u0046\u0069el\u0064\u0020\u006di\u0073\u0073\u0069\u006eg T\u0020fi\u0065\u006c\u0064\u0020\u0028\u0069\u006eco\u006d\u0070\u0061\u0074\u0069\u0062\u006ce\u0029")
	}
	return _bdad
}

// NewPdfReaderFromFile creates a new PdfReader from the speficied PDF file.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderFromFile(pdfFile string, opts *ReaderOpts) (*PdfReader, *_db.File, error) {
	const _gfbf = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0046\u0072\u006f\u006dF\u0069\u006c\u0065"
	_ggbe, _ebabac := _db.Open(pdfFile)
	if _ebabac != nil {
		return nil, nil, _ebabac
	}
	_gabc, _ebabac := _gfceb(_ggbe, opts, true, _gfbf)
	if _ebabac != nil {
		_ggbe.Close()
		return nil, nil, _ebabac
	}
	return _gabc, _ggbe, nil
}

// GetCerts returns the signature certificate chain.
func (_bdfdf *PdfSignature) GetCerts() ([]*_bg.Certificate, error) {
	var _bcad []func() ([]*_bg.Certificate, error)
	switch _dcabe, _ := _cde.GetNameVal(_bdfdf.SubFilter); _dcabe {
	case "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064", "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064":
		_bcad = append(_bcad, _bdfdf.extractChainFromPKCS7, _bdfdf.extractChainFromCert)
	case "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031":
		_bcad = append(_bcad, _bdfdf.extractChainFromCert)
	case "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031":
		_bcad = append(_bcad, _bdfdf.extractChainFromPKCS7)
	default:
		return nil, _ee.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020S\u0075b\u0046i\u006c\u0074\u0065\u0072\u003a\u0020\u0025s", _dcabe)
	}
	for _, _gaedfc := range _bcad {
		_egceg, _bfeeg := _gaedfc()
		if _bfeeg != nil {
			return nil, _bfeeg
		}
		if len(_egceg) > 0 {
			return _egceg, nil
		}
	}
	return nil, ErrSignNoCertificates
}

// Evaluate runs the function. Input is [x1 x2 x3].
func (_cfggb *PdfFunctionType4) Evaluate(xVec []float64) ([]float64, error) {
	if _cfggb._bdfd == nil {
		_cfggb._bdfd = _dg.NewPSExecutor(_cfggb.Program)
	}
	var _dcde []_dg.PSObject
	for _, _eadea := range xVec {
		_dcde = append(_dcde, _dg.MakeReal(_eadea))
	}
	_ccbge, _eefaf := _cfggb._bdfd.Execute(_dcde)
	if _eefaf != nil {
		return nil, _eefaf
	}
	_gbge, _eefaf := _dg.PSObjectArrayToFloat64Array(_ccbge)
	if _eefaf != nil {
		return nil, _eefaf
	}
	return _gbge, nil
}
func (_cgec *LTV) generateVRIKey(_baeb *PdfSignature) (string, error) {
	_adaab, _bbabf := _dgaff(_baeb.Contents.Bytes())
	if _bbabf != nil {
		return "", _bbabf
	}
	return _dac.ToUpper(_ed.EncodeToString(_adaab)), nil
}

// ToPdfObject implements interface PdfModel.
func (_gffbgg *PdfSignature) ToPdfObject() _cde.PdfObject {
	_ebfgf := _gffbgg._cabd
	var _gbaba *_cde.PdfObjectDictionary
	if _bdgef, _fceeb := _ebfgf.PdfObject.(*pdfSignDictionary); _fceeb {
		_gbaba = _bdgef.PdfObjectDictionary
	} else {
		_gbaba = _ebfgf.PdfObject.(*_cde.PdfObjectDictionary)
	}
	_gbaba.SetIfNotNil("\u0054\u0079\u0070\u0065", _gffbgg.Type)
	_gbaba.SetIfNotNil("\u0046\u0069\u006c\u0074\u0065\u0072", _gffbgg.Filter)
	_gbaba.SetIfNotNil("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r", _gffbgg.SubFilter)
	_gbaba.SetIfNotNil("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _gffbgg.ByteRange)
	_gbaba.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _gffbgg.Contents)
	_gbaba.SetIfNotNil("\u0043\u0065\u0072\u0074", _gffbgg.Cert)
	_gbaba.SetIfNotNil("\u004e\u0061\u006d\u0065", _gffbgg.Name)
	_gbaba.SetIfNotNil("\u0052\u0065\u0061\u0073\u006f\u006e", _gffbgg.Reason)
	_gbaba.SetIfNotNil("\u004d", _gffbgg.M)
	_gbaba.SetIfNotNil("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e", _gffbgg.Reference)
	_gbaba.SetIfNotNil("\u0043h\u0061\u006e\u0067\u0065\u0073", _gffbgg.Changes)
	_gbaba.SetIfNotNil("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f", _gffbgg.ContactInfo)
	return _ebfgf
}
func (_gbbg *PdfReader) newPdfAnnotationMovieFromDict(_efda *_cde.PdfObjectDictionary) (*PdfAnnotationMovie, error) {
	_defd := PdfAnnotationMovie{}
	_defd.T = _efda.Get("\u0054")
	_defd.Movie = _efda.Get("\u004d\u006f\u0076i\u0065")
	_defd.A = _efda.Get("\u0041")
	return &_defd, nil
}
func _ffbc() string {
	_dccfe.Lock()
	defer _dccfe.Unlock()
	return _dbbgaa
}

// PdfPageResources is a Page resources model.
// Implements PdfModel.
type PdfPageResources struct {
	ExtGState  _cde.PdfObject
	ColorSpace _cde.PdfObject
	Pattern    _cde.PdfObject
	Shading    _cde.PdfObject
	XObject    _cde.PdfObject
	Font       _cde.PdfObject
	ProcSet    _cde.PdfObject
	Properties _cde.PdfObject
	_eaegg     *_cde.PdfObjectDictionary
	_bfff      *PdfPageResourcesColorspaces
}

// PdfVersion returns version of the PDF file.
func (_eecfg *PdfReader) PdfVersion() _cde.Version { return _eecfg._aggcgb.PdfVersion() }

// ToPdfObject converts the PdfFont object to its PDF representation.
func (_egfd *PdfFont) ToPdfObject() _cde.PdfObject {
	if _egfd._gbcff == nil {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0066\u006f\u006e\u0074 \u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return _cde.MakeNull()
	}
	return _egfd._gbcff.ToPdfObject()
}

// NewPdfPage returns a new PDF page.
func NewPdfPage() *PdfPage {
	_caff := PdfPage{}
	_caff._gbbc = _cde.MakeDict()
	_caff.Resources = NewPdfPageResources()
	_fgeb := _cde.PdfIndirectObject{}
	_fgeb.PdfObject = _caff._gbbc
	_caff._dcaeff = &_fgeb
	return &_caff
}

// ToPdfObject implements interface PdfModel.
func (_gdbbd *PdfAnnotationCaret) ToPdfObject() _cde.PdfObject {
	_gdbbd.PdfAnnotation.ToPdfObject()
	_aabc := _gdbbd._bddg
	_acgb := _aabc.PdfObject.(*_cde.PdfObjectDictionary)
	_gdbbd.PdfAnnotationMarkup.appendToPdfDictionary(_acgb)
	_acgb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0043\u0061\u0072e\u0074"))
	_acgb.SetIfNotNil("\u0052\u0044", _gdbbd.RD)
	_acgb.SetIfNotNil("\u0053\u0079", _gdbbd.Sy)
	return _aabc
}

// SetDocInfo set document info.
// This will overwrite any globally declared document info.
func (_gaad *PdfWriter) SetDocInfo(info *PdfInfo) { _gaad.setDocInfo(info.ToPdfObject()) }

// A PdfPattern can represent a Pattern, either a tiling pattern or a shading pattern.
// Note that all patterns shall be treated as colours; a Pattern colour space shall be established with the CS or cs
// operator just like other colour spaces, and a particular pattern shall be installed as the current colour with the
// SCN or scn operator.
type PdfPattern struct {

	// Type: Pattern
	PatternType int64
	_abddb      PdfModel
	_eecac      _cde.PdfObject
}

var _ pdfFont = (*pdfFontSimple)(nil)

func _eecae(_gbebb *_cde.PdfObjectDictionary) (*PdfShadingType3, error) {
	_daba := PdfShadingType3{}
	_ccbef := _gbebb.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _ccbef == nil {
		_ad.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0043\u006f\u006f\u0072\u0064\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_fabec, _bcaae := _ccbef.(*_cde.PdfObjectArray)
	if !_bcaae {
		_ad.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ccbef)
		return nil, _cde.ErrTypeError
	}
	if _fabec.Len() != 6 {
		_ad.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0036\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _fabec.Len())
		return nil, ErrInvalidAttribute
	}
	_daba.Coords = _fabec
	if _ffaeg := _gbebb.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _ffaeg != nil {
		_ffaeg = _cde.TraceToDirectObject(_ffaeg)
		_bbfdc, _begag := _ffaeg.(*_cde.PdfObjectArray)
		if !_begag {
			_ad.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ffaeg)
			return nil, _cde.ErrTypeError
		}
		_daba.Domain = _bbfdc
	}
	_ccbef = _gbebb.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _ccbef == nil {
		_ad.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_daba.Function = []PdfFunction{}
	if _bfef, _dggfd := _ccbef.(*_cde.PdfObjectArray); _dggfd {
		for _, _fbaca := range _bfef.Elements() {
			_dgabd, _geaa := _cfdbb(_fbaca)
			if _geaa != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _geaa)
				return nil, _geaa
			}
			_daba.Function = append(_daba.Function, _dgabd)
		}
	} else {
		_gbbcf, _fadbf := _cfdbb(_ccbef)
		if _fadbf != nil {
			_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _fadbf)
			return nil, _fadbf
		}
		_daba.Function = append(_daba.Function, _gbbcf)
	}
	if _cefcf := _gbebb.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _cefcf != nil {
		_cefcf = _cde.TraceToDirectObject(_cefcf)
		_daccf, _fefcd := _cefcf.(*_cde.PdfObjectArray)
		if !_fefcd {
			_ad.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cefcf)
			return nil, _cde.ErrTypeError
		}
		if _daccf.Len() != 2 {
			_ad.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _daccf.Len())
			return nil, ErrInvalidAttribute
		}
		_daba.Extend = _daccf
	}
	return &_daba, nil
}

// ToPdfObject returns the button field dictionary within an indirect object.
func (_eacef *PdfFieldButton) ToPdfObject() _cde.PdfObject {
	_eacef.PdfField.ToPdfObject()
	_bebd := _eacef._afgc
	_gbba := _bebd.PdfObject.(*_cde.PdfObjectDictionary)
	_gbba.Set("\u0046\u0054", _cde.MakeName("\u0042\u0074\u006e"))
	if _eacef.Opt != nil {
		_gbba.Set("\u004f\u0070\u0074", _eacef.Opt)
	}
	return _bebd
}

// CustomKeys returns all custom info keys as list.
func (_dbaeg *PdfInfo) CustomKeys() []string {
	if _dbaeg._ccef == nil {
		return nil
	}
	_efedb := make([]string, len(_dbaeg._ccef.Keys()))
	for _, _bbfd := range _dbaeg._ccef.Keys() {
		_efedb = append(_efedb, _bbfd.String())
	}
	return _efedb
}
func (_afgb *PdfReader) newPdfAnnotationSoundFromDict(_beef *_cde.PdfObjectDictionary) (*PdfAnnotationSound, error) {
	_fdd := PdfAnnotationSound{}
	_eefc, _cggga := _afgb.newPdfAnnotationMarkupFromDict(_beef)
	if _cggga != nil {
		return nil, _cggga
	}
	_fdd.PdfAnnotationMarkup = _eefc
	_fdd.Name = _beef.Get("\u004e\u0061\u006d\u0065")
	_fdd.Sound = _beef.Get("\u0053\u006f\u0075n\u0064")
	return &_fdd, nil
}

// AddPages adds pages to be appended to the end of the source PDF.
func (_cdadg *PdfAppender) AddPages(pages ...*PdfPage) {
	for _, _cffe := range pages {
		_cffe = _cffe.Duplicate()
		_fdcef(_cffe)
		_cdadg._gfb = append(_cdadg._gfb, _cffe)
	}
}

// PdfAnnotationSquare represents Square annotations.
// (Section 12.5.6.8).
type PdfAnnotationSquare struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _cde.PdfObject
	IC _cde.PdfObject
	BE _cde.PdfObject
	RD _cde.PdfObject
}

// GetContainingPdfObject returns the container of the image object (indirect object).
func (_ccacf *XObjectImage) GetContainingPdfObject() _cde.PdfObject { return _ccacf._bbaed }
func (_adcce *PdfReader) newPdfOutlineItemFromIndirectObject(_ccdga *_cde.PdfIndirectObject) (*PdfOutlineItem, error) {
	_eeecc, _cecec := _ccdga.PdfObject.(*_cde.PdfObjectDictionary)
	if !_cecec {
		return nil, _ee.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_bfdae := NewPdfOutlineItem()
	_eeaeg := _eeecc.Get("\u0054\u0069\u0074l\u0065")
	if _eeaeg == nil {
		return nil, _ee.Errorf("\u006d\u0069\u0073s\u0069\u006e\u0067\u0020\u0054\u0069\u0074\u006c\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0049\u0074\u0065\u006d\u0020\u0028r\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	_eeafb, _fgfcd := _cde.GetString(_eeaeg)
	if !_fgfcd {
		return nil, _ee.Errorf("\u0074\u0069\u0074le\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0054\u0029", _eeaeg)
	}
	_bfdae.Title = _eeafb
	if _fggdg := _eeecc.Get("\u0043\u006f\u0075n\u0074"); _fggdg != nil {
		_bfcb, _ecdf := _fggdg.(*_cde.PdfObjectInteger)
		if !_ecdf {
			return nil, _ee.Errorf("\u0063o\u0075\u006e\u0074\u0020n\u006f\u0074\u0020\u0061\u006e \u0069n\u0074e\u0067\u0065\u0072\u0020\u0028\u0025\u0054)", _fggdg)
		}
		_dcgfd := int64(*_bfcb)
		_bfdae.Count = &_dcgfd
	}
	if _gdcbe := _eeecc.Get("\u0044\u0065\u0073\u0074"); _gdcbe != nil {
		_bfdae.Dest = _cde.ResolveReference(_gdcbe)
		if !_adcce._cdgee {
			_dcggb := _adcce.traverseObjectData(_bfdae.Dest)
			if _dcggb != nil {
				return nil, _dcggb
			}
		}
	}
	if _ceggc := _eeecc.Get("\u0041"); _ceggc != nil {
		_bfdae.A = _cde.ResolveReference(_ceggc)
		if !_adcce._cdgee {
			_gecacd := _adcce.traverseObjectData(_bfdae.A)
			if _gecacd != nil {
				return nil, _gecacd
			}
		}
	}
	if _bdbeb := _eeecc.Get("\u0053\u0045"); _bdbeb != nil {
		_bfdae.SE = nil
	}
	if _bgbg := _eeecc.Get("\u0043"); _bgbg != nil {
		_bfdae.C = _cde.ResolveReference(_bgbg)
	}
	if _bagea := _eeecc.Get("\u0046"); _bagea != nil {
		_bfdae.F = _cde.ResolveReference(_bagea)
	}
	return _bfdae, nil
}

// Optimizer is the interface that performs optimization of PDF object structure for output writing.
//
// Optimize receives a slice of input `objects`, performs optimization, including removing, replacing objects and
// output the optimized slice of objects.
type Optimizer interface {
	Optimize(_adaae []_cde.PdfObject) ([]_cde.PdfObject, error)
}

// ImageToRGB returns an error since an image cannot be defined in a pattern colorspace.
func (_ffgb *PdfColorspaceSpecialPattern) ImageToRGB(img Image) (Image, error) {
	_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0061\u0074\u0074\u0065\u0072n \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	return img, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0066\u006f\u0072\u0020\u0069m\u0061\u0067\u0065\u0020\u0028p\u0061\u0074t\u0065\u0072\u006e\u0029")
}

// NewPdfAnnotationPolygon returns a new polygon annotation.
func NewPdfAnnotationPolygon() *PdfAnnotationPolygon {
	_ebc := NewPdfAnnotation()
	_fcgb := &PdfAnnotationPolygon{}
	_fcgb.PdfAnnotation = _ebc
	_fcgb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ebc.SetContext(_fcgb)
	return _fcgb
}

// ToPdfObject implements interface PdfModel.
func (_aaed *PdfAnnotationInk) ToPdfObject() _cde.PdfObject {
	_aaed.PdfAnnotation.ToPdfObject()
	_adga := _aaed._bddg
	_aca := _adga.PdfObject.(*_cde.PdfObjectDictionary)
	_aaed.PdfAnnotationMarkup.appendToPdfDictionary(_aca)
	_aca.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0049\u006e\u006b"))
	_aca.SetIfNotNil("\u0049n\u006b\u004c\u0069\u0073\u0074", _aaed.InkList)
	_aca.SetIfNotNil("\u0042\u0053", _aaed.BS)
	return _adga
}

// GetPrimitiveFromModel returns the primitive object corresponding to the input `model`.
func (_gafdgg *modelManager) GetPrimitiveFromModel(model PdfModel) _cde.PdfObject {
	_eagaf, _cdfcc := _gafdgg._bfaea[model]
	if !_cdfcc {
		return nil
	}
	return _eagaf
}

// ToPdfObject returns the PDF representation of the outline tree node.
func (_acecg *PdfOutlineTreeNode) ToPdfObject() _cde.PdfObject {
	return _acecg.GetContext().ToPdfObject()
}

// ToInteger convert to an integer format.
func (_caad *PdfColorCalRGB) ToInteger(bits int) [3]uint32 {
	_afbc := _ced.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_afbc * _caad.A()), uint32(_afbc * _caad.B()), uint32(_afbc * _caad.C())}
}

// PdfInfoTrapped specifies pdf trapped information.
type PdfInfoTrapped string

// NewPdfAnnotationCircle returns a new circle annotation.
func NewPdfAnnotationCircle() *PdfAnnotationCircle {
	_cabf := NewPdfAnnotation()
	_bbfc := &PdfAnnotationCircle{}
	_bbfc.PdfAnnotation = _cabf
	_bbfc.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cabf.SetContext(_bbfc)
	return _bbfc
}
func (_ddfeg *PdfWriter) writeDocumentVersion() {
	if _ddfeg._aabfe {
		_ddfeg.writeString("\u000a")
	} else {
		_ddfeg.writeString(_ee.Sprintf("\u0025\u0025\u0050D\u0046\u002d\u0025\u0064\u002e\u0025\u0064\u000a", _ddfeg._cgdcc.Major, _ddfeg._cgdcc.Minor))
		_ddfeg.writeString("\u0025\u00e2\u00e3\u00cf\u00d3\u000a")
	}
}

// ToInteger convert to an integer format.
func (_faegd *PdfColorLab) ToInteger(bits int) [3]uint32 {
	_efce := _ced.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_efce * _faegd.L()), uint32(_efce * _faegd.A()), uint32(_efce * _faegd.B())}
}

// ToPdfObject implements interface PdfModel.
func (_aec *PdfAnnotationPrinterMark) ToPdfObject() _cde.PdfObject {
	_aec.PdfAnnotation.ToPdfObject()
	_dgee := _aec._bddg
	_cfd := _dgee.PdfObject.(*_cde.PdfObjectDictionary)
	_cfd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b"))
	_cfd.SetIfNotNil("\u004d\u004e", _aec.MN)
	return _dgee
}

// ToPdfObject returns the PDF representation of the function.
func (_bdag *PdfFunctionType3) ToPdfObject() _cde.PdfObject {
	_fdccb := _cde.MakeDict()
	_fdccb.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _cde.MakeInteger(3))
	_fcabd := &_cde.PdfObjectArray{}
	for _, _gedbg := range _bdag.Domain {
		_fcabd.Append(_cde.MakeFloat(_gedbg))
	}
	_fdccb.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _fcabd)
	if _bdag.Range != nil {
		_cgcf := &_cde.PdfObjectArray{}
		for _, _cbgeb := range _bdag.Range {
			_cgcf.Append(_cde.MakeFloat(_cbgeb))
		}
		_fdccb.Set("\u0052\u0061\u006eg\u0065", _cgcf)
	}
	if _bdag.Functions != nil {
		_cgaca := &_cde.PdfObjectArray{}
		for _, _efcbc := range _bdag.Functions {
			_cgaca.Append(_efcbc.ToPdfObject())
		}
		_fdccb.Set("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es", _cgaca)
	}
	if _bdag.Bounds != nil {
		_gefcf := &_cde.PdfObjectArray{}
		for _, _dcfb := range _bdag.Bounds {
			_gefcf.Append(_cde.MakeFloat(_dcfb))
		}
		_fdccb.Set("\u0042\u006f\u0075\u006e\u0064\u0073", _gefcf)
	}
	if _bdag.Encode != nil {
		_eggcb := &_cde.PdfObjectArray{}
		for _, _fbfgb := range _bdag.Encode {
			_eggcb.Append(_cde.MakeFloat(_fbfgb))
		}
		_fdccb.Set("\u0045\u006e\u0063\u006f\u0064\u0065", _eggcb)
	}
	if _bdag._fbbgd != nil {
		_bdag._fbbgd.PdfObject = _fdccb
		return _bdag._fbbgd
	}
	return _fdccb
}

// SetXObjectImageByName adds the provided XObjectImage to the page resources.
// The added XObjectImage is identified by the specified name.
func (_daecb *PdfPageResources) SetXObjectImageByName(keyName _cde.PdfObjectName, ximg *XObjectImage) error {
	_aced := ximg.ToPdfObject().(*_cde.PdfObjectStream)
	_ceeag := _daecb.SetXObjectByName(keyName, _aced)
	return _ceeag
}
func _ebbec(_feagfg _cde.PdfObject) []*_cde.PdfObjectStream {
	if _feagfg == nil {
		return nil
	}
	_ebagd, _cccffg := _cde.GetArray(_feagfg)
	if !_cccffg || _ebagd.Len() == 0 {
		return nil
	}
	_gfdag := make([]*_cde.PdfObjectStream, 0, _ebagd.Len())
	for _, _gdbfg := range _ebagd.Elements() {
		if _aebbf, _ebba := _cde.GetStream(_gdbfg); _ebba {
			_gfdag = append(_gfdag, _aebbf)
		}
	}
	return _gfdag
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_ebeaa *PdfPageResourcesColorspaces) ToPdfObject() _cde.PdfObject {
	_fbgbg := _cde.MakeDict()
	for _, _bdgdb := range _ebeaa.Names {
		_fbgbg.Set(_cde.PdfObjectName(_bdgdb), _ebeaa.Colorspaces[_bdgdb].ToPdfObject())
	}
	if _ebeaa._ccdd != nil {
		_ebeaa._ccdd.PdfObject = _fbgbg
		return _ebeaa._ccdd
	}
	return _fbgbg
}

// Encrypt encrypts the output file with a specified user/owner password.
func (_agbbf *PdfWriter) Encrypt(userPass, ownerPass []byte, options *EncryptOptions) error {
	_cbaf := RC4_128bit
	if options != nil {
		_cbaf = options.Algorithm
	}
	_dbcec := _ccg.PermOwner
	if options != nil {
		_dbcec = options.Permissions
	}
	var _bebgf _cd.Filter
	switch _cbaf {
	case RC4_128bit:
		_bebgf = _cd.NewFilterV2(16)
	case AES_128bit:
		_bebgf = _cd.NewFilterAESV2()
	case AES_256bit:
		_bebgf = _cd.NewFilterAESV3()
	default:
		return _ee.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020a\u006cg\u006fr\u0069\u0074\u0068\u006d\u003a\u0020\u0025v", options.Algorithm)
	}
	_deeed, _gafad, _fcbed := _cde.PdfCryptNewEncrypt(_bebgf, userPass, ownerPass, _dbcec)
	if _fcbed != nil {
		return _fcbed
	}
	_agbbf._ccgbe = _deeed
	if _gafad.Major != 0 {
		_agbbf.SetVersion(_gafad.Major, _gafad.Minor)
	}
	_agbbf._dcbae = _gafad.Encrypt
	_agbbf._gfegf, _agbbf._dcace = _gafad.ID0, _gafad.ID1
	_bcdbc := _cde.MakeIndirectObject(_gafad.Encrypt)
	_agbbf._bdbdb = _bcdbc
	_agbbf.addObject(_bcdbc)
	return nil
}
func _ffabc() string { _dccfe.Lock(); defer _dccfe.Unlock(); return _bdcea }

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_dgeb *PdfColorspaceCalGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bedfc := vals[0]
	if _bedfc < 0.0 || _bedfc > 1.0 {
		_ad.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bedfc)
		return nil, ErrColorOutOfRange
	}
	_ffde := NewPdfColorCalGray(_bedfc)
	return _ffde, nil
}
func (_fadcf *pdfCIDFontType0) getFontDescriptor() *PdfFontDescriptor { return _fadcf._fagf }

// ToPdfObject converts the pdfCIDFontType0 to a PDF representation.
func (_adgfb *pdfCIDFontType0) ToPdfObject() _cde.PdfObject { return _cde.MakeNull() }

// NewPdfActionJavaScript returns a new "javaScript" action.
func NewPdfActionJavaScript() *PdfActionJavaScript {
	_efe := NewPdfAction()
	_abg := &PdfActionJavaScript{}
	_abg.PdfAction = _efe
	_efe.SetContext(_abg)
	return _abg
}
func (_dcea *PdfSignature) extractChainFromPKCS7() ([]*_bg.Certificate, error) {
	_aabdc, _cabaab := _cb.Parse(_dcea.Contents.Bytes())
	if _cabaab != nil {
		return nil, _cabaab
	}
	return _aabdc.Certificates, nil
}

// PdfColorCalRGB represents a color in the Colorimetric CIE RGB colorspace.
// A, B, C components
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorCalRGB [3]float64

// NewPdfAnnotationSquare returns a new square annotation.
func NewPdfAnnotationSquare() *PdfAnnotationSquare {
	_fea := NewPdfAnnotation()
	_ddf := &PdfAnnotationSquare{}
	_ddf.PdfAnnotation = _fea
	_ddf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fea.SetContext(_ddf)
	return _ddf
}

// ToPdfObject implements interface PdfModel.
func (_gfeg *PdfAnnotationScreen) ToPdfObject() _cde.PdfObject {
	_gfeg.PdfAnnotation.ToPdfObject()
	_bcff := _gfeg._bddg
	_eddb := _bcff.PdfObject.(*_cde.PdfObjectDictionary)
	_eddb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0053\u0063\u0072\u0065\u0065\u006e"))
	_eddb.SetIfNotNil("\u0054", _gfeg.T)
	_eddb.SetIfNotNil("\u004d\u004b", _gfeg.MK)
	_eddb.SetIfNotNil("\u0041", _gfeg.A)
	_eddb.SetIfNotNil("\u0041\u0041", _gfeg.AA)
	return _bcff
}
func (_beba *PdfReader) newPdfActionTransFromDict(_dgaf *_cde.PdfObjectDictionary) (*PdfActionTrans, error) {
	return &PdfActionTrans{Trans: _dgaf.Get("\u0054\u0072\u0061n\u0073")}, nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_fbfc *PdfColorspaceSpecialPattern) ToPdfObject() _cde.PdfObject {
	if _fbfc.UnderlyingCS == nil {
		return _cde.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e")
	}
	_bfga := _cde.MakeArray(_cde.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_bfga.Append(_fbfc.UnderlyingCS.ToPdfObject())
	if _fbfc._ffcb != nil {
		_fbfc._ffcb.PdfObject = _bfga
		return _fbfc._ffcb
	}
	return _bfga
}

// Duplicate creates a duplicate page based on the current one and returns it.
func (_begfg *PdfPage) Duplicate() *PdfPage {
	_cbfbc := *_begfg
	_cbfbc._gbbc = _cde.MakeDict()
	_cbfbc._dcaeff = _cde.MakeIndirectObject(_cbfbc._gbbc)
	return &_cbfbc
}
func _deaee(_feagf *PdfField, _afabd _cde.PdfObject) error {
	switch _feagf.GetContext().(type) {
	case *PdfFieldText:
		switch _gfcb := _afabd.(type) {
		case *_cde.PdfObjectName:
			_defg := _gfcb
			_ad.Log.Debug("\u0055\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u003a\u0020\u0047\u006f\u0074 \u0056\u0020\u0061\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u003e\u0020c\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f s\u0074\u0072\u0069\u006e\u0067\u0020\u0027\u0025\u0073\u0027", _defg.String())
			_feagf.V = _cde.MakeEncodedString(_gfcb.String(), true)
		case *_cde.PdfObjectString:
			_feagf.V = _cde.MakeEncodedString(_gfcb.String(), true)
		default:
			_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0056\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u0028\u0025\u0023\u0076\u0029", _gfcb, _gfcb)
		}
	case *PdfFieldButton:
		switch _afabd.(type) {
		case *_cde.PdfObjectName:
			if len(_afabd.String()) > 0 {
				_feagf.V = _afabd
				_ccfa(_feagf, _afabd)
			}
		case *_cde.PdfObjectString:
			if len(_afabd.String()) > 0 {
				_feagf.V = _cde.MakeName(_afabd.String())
				_ccfa(_feagf, _feagf.V)
			}
		default:
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _feagf.PartialName(), _afabd)
			_feagf.V = _afabd
		}
	case *PdfFieldChoice:
		switch _afabd.(type) {
		case *_cde.PdfObjectName:
			if len(_afabd.String()) > 0 {
				_feagf.V = _cde.MakeString(_afabd.String())
				_ccfa(_feagf, _afabd)
			}
		case *_cde.PdfObjectString:
			if len(_afabd.String()) > 0 {
				_feagf.V = _afabd
				_ccfa(_feagf, _cde.MakeName(_afabd.String()))
			}
		default:
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _feagf.PartialName(), _afabd)
			_feagf.V = _afabd
		}
	case *PdfFieldSignature:
		_ad.Log.Debug("\u0054\u004f\u0044\u004f\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0061\u0070\u0070e\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0079\u0065\u0074\u003a\u0020\u0025\u0073\u002f\u0025v", _feagf.PartialName(), _afabd)
	}
	return nil
}

// SetForms sets the Acroform for a PDF file.
func (_dfec *PdfWriter) SetForms(form *PdfAcroForm) error { _dfec._fcacfe = form; return nil }

// FontDescriptor returns font's PdfFontDescriptor. This may be a builtin descriptor for standard 14
// fonts but must be an explicit descriptor for other fonts.
func (_dfcda *PdfFont) FontDescriptor() *PdfFontDescriptor {
	if _dfcda.baseFields()._fagf != nil {
		return _dfcda.baseFields()._fagf
	}
	if _ebdb := _dfcda._gbcff.getFontDescriptor(); _ebdb != nil {
		return _ebdb
	}
	_ad.Log.Error("\u0041\u006cl \u0066\u006f\u006et\u0073\u0020\u0068\u0061ve \u0061 D\u0065\u0073\u0063\u0072\u0069\u0070\u0074or\u002e\u0020\u0066\u006f\u006e\u0074\u003d%\u0073", _dfcda)
	return nil
}

// NewLTV returns a new LTV client.
func NewLTV(appender *PdfAppender) (*LTV, error) {
	_fcbfc := appender.Reader.DSS
	if _fcbfc == nil {
		_fcbfc = NewDSS()
	}
	if _gedec := _fcbfc.generateHashMaps(); _gedec != nil {
		return nil, _gedec
	}
	return &LTV{CertClient: _bbe.NewCertClient(), OCSPClient: _bbe.NewOCSPClient(), CRLClient: _bbe.NewCRLClient(), SkipExisting: true, _ffab: appender, _geeeg: _fcbfc}, nil
}

// GetPatternByName gets the pattern specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_agdcc *PdfPageResources) GetPatternByName(keyName _cde.PdfObjectName) (*PdfPattern, bool) {
	if _agdcc.Pattern == nil {
		return nil, false
	}
	_ebga, _gfbcf := _cde.TraceToDirectObject(_agdcc.Pattern).(*_cde.PdfObjectDictionary)
	if !_gfbcf {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074t\u0065\u0072\u006e\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _agdcc.Pattern)
		return nil, false
	}
	if _eebba := _ebga.Get(keyName); _eebba != nil {
		_agce, _defee := _cgafb(_eebba)
		if _defee != nil {
			_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0070\u0061\u0074\u0074\u0065\u0072n\u003a\u0020\u0025\u0076", _defee)
			return nil, false
		}
		return _agce, true
	}
	return nil, false
}

// WriteString outputs the object as it is to be written to file.
func (_babfd *PdfTransformParamsDocMDP) WriteString() string {
	return _babfd.ToPdfObject().WriteString()
}

// Clear clears flag fl from the flag and returns the resulting flag.
func (_baafg FieldFlag) Clear(fl FieldFlag) FieldFlag { return FieldFlag(_baafg.Mask() &^ fl.Mask()) }

// ImageToRGB convert 1-component grayscale data to 3-component RGB.
func (_cbcbe *PdfColorspaceDeviceGray) ImageToRGB(img Image) (Image, error) {
	if img.ColorComponents != 1 {
		return img, _ceg.New("\u0074\u0068e \u0070\u0072\u006fv\u0069\u0064\u0065\u0064 im\u0061ge\u0020\u0069\u0073\u0020\u006e\u006f\u0074 g\u0072\u0061\u0079\u0020\u0073\u0063\u0061l\u0065")
	}
	_bffe, _cdfg := _ff.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._deegf, img._aaafb)
	if _cdfg != nil {
		return img, _cdfg
	}
	_gadb, _cdfg := _ff.NRGBAConverter.Convert(_bffe)
	if _cdfg != nil {
		return img, _cdfg
	}
	_cfgf := _bddb(_gadb.Base())
	_ad.Log.Trace("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079\u0020\u002d>\u0020\u0052\u0047\u0042")
	_ad.Log.Trace("s\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", img.Data)
	_ad.Log.Trace("\u0052G\u0042 \u0073\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", _cfgf.Data)
	_ad.Log.Trace("\u0025\u0076\u0020\u002d\u003e\u0020\u0025\u0076", img, _cfgf)
	return _cfgf, nil
}

// GetContentStream returns the pattern cell's content stream
func (_bgbgd *PdfTilingPattern) GetContentStream() ([]byte, error) {
	_edgg, _, _fddbcb := _bgbgd.GetContentStreamWithEncoder()
	return _edgg, _fddbcb
}

// Y returns the value of the yellow component of the color.
func (_adac *PdfColorDeviceCMYK) Y() float64 { return _adac[2] }

// PdfFieldChoice represents a choice field which includes scrollable list boxes and combo boxes.
type PdfFieldChoice struct {
	*PdfField
	Opt *_cde.PdfObjectArray
	TI  *_cde.PdfObjectInteger
	I   *_cde.PdfObjectArray
}

// NewPdfColorDeviceRGB returns a new PdfColorDeviceRGB based on the r,g,b component values.
func NewPdfColorDeviceRGB(r, g, b float64) *PdfColorDeviceRGB {
	_gddb := PdfColorDeviceRGB{r, g, b}
	return &_gddb
}

// CharcodeBytesToUnicode converts PDF character codes `data` to a Go unicode string.
//
// 9.10 Extraction of Text Content (page 292)
// The process of finding glyph descriptions in OpenType fonts by a conforming reader shall be the following:
// • For Type 1 fonts using “CFF” tables, the process shall be as described in 9.6.6.2, "Encodings
//   for Type 1 Fonts".
// • For TrueType fonts using “glyf” tables, the process shall be as described in 9.6.6.4,
//   "Encodings for TrueType Fonts". Since this process sometimes produces ambiguous results,
//   conforming writers, instead of using a simple font, shall use a Type 0 font with an Identity-H
//   encoding and use the glyph indices as character codes, as described following Table 118.
func (_ebab *PdfFont) CharcodeBytesToUnicode(data []byte) (string, int, int) {
	_agab, _, _fgba := _ebab.CharcodesToUnicodeWithStats(_ebab.BytesToCharcodes(data))
	_egag := _gc.ExpandLigatures(_agab)
	return _egag, _ca.RuneCountInString(_egag), _fgba
}

// GetContainingPdfObject gets the primitive used to parse the color space.
func (_cfbdf *PdfColorspaceICCBased) GetContainingPdfObject() _cde.PdfObject { return _cfbdf._fdea }

// SetContentStream sets the pattern cell's content stream.
func (_ecbaec *PdfTilingPattern) SetContentStream(content []byte, encoder _cde.StreamEncoder) error {
	_bgacc, _bffae := _ecbaec._eecac.(*_cde.PdfObjectStream)
	if !_bffae {
		_ad.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ecbaec._eecac)
		return _cde.ErrTypeError
	}
	if encoder == nil {
		encoder = _cde.NewRawEncoder()
	}
	_gcafa := _bgacc.PdfObjectDictionary
	_ebcaa := encoder.MakeStreamDict()
	_gcafa.Merge(_ebcaa)
	_fcafb, _aabef := encoder.EncodeBytes(content)
	if _aabef != nil {
		return _aabef
	}
	_gcafa.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cde.MakeInteger(int64(len(_fcafb))))
	_bgacc.Stream = _fcafb
	return nil
}
func (_efg *PdfReader) newPdfAnnotationWidgetFromDict(_gdfc *_cde.PdfObjectDictionary) (*PdfAnnotationWidget, error) {
	_gaae := PdfAnnotationWidget{}
	_gaae.H = _gdfc.Get("\u0048")
	_gaae.MK = _gdfc.Get("\u004d\u004b")
	_gaae.A = _gdfc.Get("\u0041")
	_gaae.AA = _gdfc.Get("\u0041\u0041")
	_gaae.BS = _gdfc.Get("\u0042\u0053")
	_gaae.Parent = _gdfc.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	return &_gaae, nil
}

// G returns the value of the green component of the color.
func (_bgdff *PdfColorDeviceRGB) G() float64 { return _bgdff[1] }

// NewPdfColorCalGray returns a new CalGray color.
func NewPdfColorCalGray(grayVal float64) *PdfColorCalGray {
	_gafg := PdfColorCalGray(grayVal)
	return &_gafg
}

// NewOutlineDest returns a new outline destination which can be used
// with outline items.
func NewOutlineDest(page int64, x, y float64) OutlineDest {
	return OutlineDest{Page: page, Mode: "\u0058\u0059\u005a", X: x, Y: y}
}
func _fbggd(_ebaba _cde.PdfObject) (string, error) {
	_ebaba = _cde.TraceToDirectObject(_ebaba)
	switch _dgfgf := _ebaba.(type) {
	case *_cde.PdfObjectString:
		return _dgfgf.Str(), nil
	case *_cde.PdfObjectStream:
		_edeed, _ddcee := _cde.DecodeStream(_dgfgf)
		if _ddcee != nil {
			return "", _ddcee
		}
		return string(_edeed), nil
	}
	return "", _ee.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _ebaba)
}

// GetNumComponents returns the number of color components (1 for CalGray).
func (_fcfc *PdfColorCalGray) GetNumComponents() int { return 1 }

// PdfShadingType4 is a Free-form Gouraud-shaded triangle mesh.
type PdfShadingType4 struct {
	*PdfShading
	BitsPerCoordinate *_cde.PdfObjectInteger
	BitsPerComponent  *_cde.PdfObjectInteger
	BitsPerFlag       *_cde.PdfObjectInteger
	Decode            *_cde.PdfObjectArray
	Function          []PdfFunction
}

// PdfColorCalGray represents a CalGray colorspace.
type PdfColorCalGray float64
type pdfCIDFontType0 struct {
	fontCommon
	_bfbece *_cde.PdfIndirectObject
	_efdaa  _gc.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// (Required) Dictionary that defines the character collection of the CIDFont.
	// See Table 116.
	CIDSystemInfo *_cde.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW     _cde.PdfObject
	W      _cde.PdfObject
	DW2    _cde.PdfObject
	W2     _cde.PdfObject
	_egfeb map[_gc.CharCode]float64
	_bdega float64
}

// ImageToRGB converts an image in CMYK32 colorspace to an RGB image.
func (_feeea *PdfColorspaceDeviceCMYK) ImageToRGB(img Image) (Image, error) {
	_ad.Log.Trace("\u0043\u004d\u0059\u004b\u0033\u0032\u0020\u002d\u003e\u0020\u0052\u0047\u0042")
	_ad.Log.Trace("I\u006d\u0061\u0067\u0065\u0020\u0042P\u0043\u003a\u0020\u0025\u0064\u002c \u0043\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020%\u0064", img.BitsPerComponent, img.ColorComponents)
	_ad.Log.Trace("\u004c\u0065\u006e \u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0064", len(img.Data))
	_ad.Log.Trace("H\u0065\u0069\u0067\u0068t:\u0020%\u0064\u002c\u0020\u0057\u0069d\u0074\u0068\u003a\u0020\u0025\u0064", img.Height, img.Width)
	_decf, _cdd := _ff.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._deegf, img._aaafb)
	if _cdd != nil {
		return Image{}, _cdd
	}
	_gafa, _cdd := _ff.NRGBAConverter.Convert(_decf)
	if _cdd != nil {
		return Image{}, _cdd
	}
	return _bddb(_gafa.Base()), nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_cefa *PdfShadingType7) ToPdfObject() _cde.PdfObject {
	_cefa.PdfShading.ToPdfObject()
	_dcfd, _eaffc := _cefa.getShadingDict()
	if _eaffc != nil {
		_ad.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _cefa.BitsPerCoordinate != nil {
		_dcfd.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _cefa.BitsPerCoordinate)
	}
	if _cefa.BitsPerComponent != nil {
		_dcfd.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _cefa.BitsPerComponent)
	}
	if _cefa.BitsPerFlag != nil {
		_dcfd.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _cefa.BitsPerFlag)
	}
	if _cefa.Decode != nil {
		_dcfd.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _cefa.Decode)
	}
	if _cefa.Function != nil {
		if len(_cefa.Function) == 1 {
			_dcfd.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _cefa.Function[0].ToPdfObject())
		} else {
			_dafc := _cde.MakeArray()
			for _, _eefad := range _cefa.Function {
				_dafc.Append(_eefad.ToPdfObject())
			}
			_dcfd.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dafc)
		}
	}
	return _cefa._dffg
}

// SetPdfKeywords sets the Keywords attribute of the output PDF.
func SetPdfKeywords(keywords string) { _dccfe.Lock(); defer _dccfe.Unlock(); _bdcea = keywords }

// ToImage converts an object to an Image which can be transformed or saved out.
// The image data is decoded and the Image returned.
func (_ceefc *XObjectImage) ToImage() (*Image, error) {
	_gaffe := &Image{}
	if _ceefc.Height == nil {
		return nil, _ceg.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_gaffe.Height = *_ceefc.Height
	if _ceefc.Width == nil {
		return nil, _ceg.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_gaffe.Width = *_ceefc.Width
	if _ceefc.BitsPerComponent == nil {
		switch _ceefc.Filter.(type) {
		case *_cde.CCITTFaxEncoder, *_cde.JBIG2Encoder:
			_gaffe.BitsPerComponent = 1
		case *_cde.LZWEncoder, *_cde.RunLengthEncoder:
			_gaffe.BitsPerComponent = 8
		default:
			return nil, _ceg.New("\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006fm\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
	} else {
		_gaffe.BitsPerComponent = *_ceefc.BitsPerComponent
	}
	_gaffe.ColorComponents = _ceefc.ColorSpace.GetNumComponents()
	_ceefc._bbaed.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _cde.MakeInteger(int64(_gaffe.ColorComponents)))
	_cffff, _eaacc := _cde.DecodeStream(_ceefc._bbaed)
	if _eaacc != nil {
		return nil, _eaacc
	}
	_gaffe.Data = _cffff
	if _ceefc.Decode != nil {
		_ebdged, _bebeb := _ceefc.Decode.(*_cde.PdfObjectArray)
		if !_bebeb {
			_ad.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020D\u0065\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074")
			return nil, _ceg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_gefff, _cbaca := _ebdged.ToFloat64Array()
		if _cbaca != nil {
			return nil, _cbaca
		}
		_gaffe._aaafb = _gefff
	}
	return _gaffe, nil
}

// String returns a string that describes `font`.
func (_gdee *PdfFont) String() string {
	_dfef := ""
	if _gdee._gbcff.Encoder() != nil {
		_dfef = _gdee._gbcff.Encoder().String()
	}
	return _ee.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0054\u0020\u0025s\u0020\u0025\u0073\u007d", _gdee._gbcff, _gdee.baseFields().coreString(), _dfef)
}

// IsShading specifies if the pattern is a shading pattern.
func (_fgbcg *PdfPattern) IsShading() bool { return _fgbcg.PatternType == 2 }

// SetPdfCreator sets the Creator attribute of the output PDF.
func SetPdfCreator(creator string) { _dccfe.Lock(); defer _dccfe.Unlock(); _dbbgaa = creator }
func _bebbe(_cdge _cde.PdfObject, _bcce *fontCommon) (*_fb.CMap, error) {
	_bfge, _cdfe := _cde.GetStream(_cdge)
	if !_cdfe {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0074\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0054\u006f\u0043m\u0061\u0070\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0054\u0029", _cdge)
		return nil, _cde.ErrTypeError
	}
	_gadc, _aaafg := _cde.DecodeStream(_bfge)
	if _aaafg != nil {
		return nil, _aaafg
	}
	_dgacec, _aaafg := _fb.LoadCmapFromData(_gadc, !_bcce.isCIDFont())
	if _aaafg != nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u004e\u0075\u006d\u0062\u0065\u0072\u003d\u0025\u0064\u0020\u0065\u0072r=\u0025\u0076", _bfge.ObjectNumber, _aaafg)
	}
	return _dgacec, _aaafg
}

// ParserMetadata gets the parser  metadata.
func (_dfbfe *CompliancePdfReader) ParserMetadata() _cde.ParserMetadata {
	if _dfbfe._abbc == (_cde.ParserMetadata{}) {
		_dfbfe._abbc, _ = _dfbfe._aggcgb.ParserMetadata()
	}
	return _dfbfe._abbc
}

// PdfPageResourcesColorspaces contains the colorspace in the PdfPageResources.
// Needs to have matching name and colorspace map entry. The Names define the order.
type PdfPageResourcesColorspaces struct {
	Names       []string
	Colorspaces map[string]PdfColorspace
	_ccdd       *_cde.PdfIndirectObject
}

func (_dbea *PdfReader) newPdfAnnotation3DFromDict(_gcc *_cde.PdfObjectDictionary) (*PdfAnnotation3D, error) {
	_dgdf := PdfAnnotation3D{}
	_dgdf.T3DD = _gcc.Get("\u0033\u0044\u0044")
	_dgdf.T3DV = _gcc.Get("\u0033\u0044\u0056")
	_dgdf.T3DA = _gcc.Get("\u0033\u0044\u0041")
	_dgdf.T3DI = _gcc.Get("\u0033\u0044\u0049")
	_dgdf.T3DB = _gcc.Get("\u0033\u0044\u0042")
	return &_dgdf, nil
}

// PdfShadingType6 is a Coons patch mesh.
type PdfShadingType6 struct {
	*PdfShading
	BitsPerCoordinate *_cde.PdfObjectInteger
	BitsPerComponent  *_cde.PdfObjectInteger
	BitsPerFlag       *_cde.PdfObjectInteger
	Decode            *_cde.PdfObjectArray
	Function          []PdfFunction
}

// NewImageFromGoImage creates a new NRGBA32 unidoc Image from a golang Image.
// If `goimg` is grayscale (*goimage.Gray8) then calls NewGrayImageFromGoImage instead.
func (_dcafb DefaultImageHandler) NewImageFromGoImage(goimg _gf.Image) (*Image, error) {
	_bcebc, _ebeea := _ff.FromGoImage(goimg)
	if _ebeea != nil {
		return nil, _ebeea
	}
	_cdcg := _bddb(_bcebc.Base())
	return &_cdcg, nil
}

// NewPdfReaderLazy creates a new PdfReader for `rs` in lazy-loading mode. The difference
// from NewPdfReader is that in lazy-loading mode, objects are only loaded into memory when needed
// rather than entire structure being loaded into memory on reader creation.
// Note that it may make sense to use the lazy-load reader when processing only parts of files,
// rather than loading entire file into memory. Example: splitting a few pages from a large PDF file.
func NewPdfReaderLazy(rs _f.ReadSeeker) (*PdfReader, error) {
	const _cedgb = "\u006d\u006f\u0064\u0065l:\u004e\u0065\u0077\u0050\u0064\u0066\u0052\u0065\u0061\u0064\u0065\u0072\u004c\u0061z\u0079"
	return _gfceb(rs, &ReaderOpts{LazyLoad: true}, false, _cedgb)
}

// GetXObjectImageByName returns the XObjectImage with the specified name from the
// page resources, if it exists.
func (_cbedb *PdfPageResources) GetXObjectImageByName(keyName _cde.PdfObjectName) (*XObjectImage, error) {
	_gefed, _cgeca := _cbedb.GetXObjectByName(keyName)
	if _gefed == nil {
		return nil, nil
	}
	if _cgeca != XObjectTypeImage {
		return nil, _ceg.New("\u006e\u006f\u0074 \u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065")
	}
	_efaed, _dbdac := NewXObjectImageFromStream(_gefed)
	if _dbdac != nil {
		return nil, _dbdac
	}
	return _efaed, nil
}

// SetFillImage attach a model.Image to push button.
func (_dddd *PdfFieldButton) SetFillImage(image *Image) {
	if _dddd.IsPush() {
		_dddd._dcbd = image
	}
}

// GetFillImage get attached model.Image in push button.
func (_abea *PdfFieldButton) GetFillImage() *Image {
	if _abea.IsPush() {
		return _abea._dcbd
	}
	return nil
}

// PdfShadingType1 is a Function-based shading.
type PdfShadingType1 struct {
	*PdfShading
	Domain   *_cde.PdfObjectArray
	Matrix   *_cde.PdfObjectArray
	Function []PdfFunction
}

func (_bdcb *PdfReader) newPdfActionJavaScriptFromDict(_ggd *_cde.PdfObjectDictionary) (*PdfActionJavaScript, error) {
	return &PdfActionJavaScript{JS: _ggd.Get("\u004a\u0053")}, nil
}

// SetContext set the sub annotation (context).
func (_fcbe *PdfShading) SetContext(ctx PdfModel) { _fcbe._dgfac = ctx }

// Outline represents a PDF outline dictionary (Table 152 - p. 376).
// Currently, the Outline object can only be used to construct PDF outlines.
type Outline struct {
	Entries []*OutlineItem `json:"entries,omitempty"`
}

// SetPdfAuthor sets the Author attribute of the output PDF.
func SetPdfAuthor(author string) { _dccfe.Lock(); defer _dccfe.Unlock(); _bfcdae = author }
func _abdae() _ce.Time           { _dccfe.Lock(); defer _dccfe.Unlock(); return _cebaa }

// VariableText contains the common attributes of a variable text.
// The VariableText is typically not used directly, but is can encapsulate by PdfField
// See section 12.7.3.3 "Variable Text" and Table 222 (pp. 434-436 PDF32000_2008).
type VariableText struct {
	DA *_cde.PdfObjectString
	Q  *_cde.PdfObjectInteger
	DS *_cde.PdfObjectString
	RV _cde.PdfObject
}

func _eddbb(_egbcg _cde.PdfObject) (*PdfColorspaceCalRGB, error) {
	_gbbfd := NewPdfColorspaceCalRGB()
	if _gfef, _bbgc := _egbcg.(*_cde.PdfIndirectObject); _bbgc {
		_gbbfd._efab = _gfef
	}
	_egbcg = _cde.TraceToDirectObject(_egbcg)
	_ceaa, _aaee := _egbcg.(*_cde.PdfObjectArray)
	if !_aaee {
		return nil, _ee.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _ceaa.Len() != 2 {
		return nil, _ee.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_egbcg = _cde.TraceToDirectObject(_ceaa.Get(0))
	_bdffd, _aaee := _egbcg.(*_cde.PdfObjectName)
	if !_aaee {
		return nil, _ee.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	if *_bdffd != "\u0043\u0061\u006c\u0052\u0047\u0042" {
		return nil, _ee.Errorf("\u006e\u006f\u0074 a\u0020\u0043\u0061\u006c\u0052\u0047\u0042\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_egbcg = _cde.TraceToDirectObject(_ceaa.Get(1))
	_cabc, _aaee := _egbcg.(*_cde.PdfObjectDictionary)
	if !_aaee {
		return nil, _ee.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	_egbcg = _cabc.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_egbcg = _cde.TraceToDirectObject(_egbcg)
	_afecb, _aaee := _egbcg.(*_cde.PdfObjectArray)
	if !_aaee {
		return nil, _ee.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050o\u0069\u006e\u0074")
	}
	if _afecb.Len() != 3 {
		return nil, _ee.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0057h\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_ffga, _eaa := _afecb.GetAsFloat64Slice()
	if _eaa != nil {
		return nil, _eaa
	}
	_gbbfd.WhitePoint = _ffga
	_egbcg = _cabc.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _egbcg != nil {
		_egbcg = _cde.TraceToDirectObject(_egbcg)
		_acga, _dcbb := _egbcg.(*_cde.PdfObjectArray)
		if !_dcbb {
			return nil, _ee.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050o\u0069\u006e\u0074")
		}
		if _acga.Len() != 3 {
			return nil, _ee.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0042l\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		_acbb, _agae := _acga.GetAsFloat64Slice()
		if _agae != nil {
			return nil, _agae
		}
		_gbbfd.BlackPoint = _acbb
	}
	_egbcg = _cabc.Get("\u0047\u0061\u006dm\u0061")
	if _egbcg != nil {
		_egbcg = _cde.TraceToDirectObject(_egbcg)
		_bgbc, _ccebg := _egbcg.(*_cde.PdfObjectArray)
		if !_ccebg {
			return nil, _ee.Errorf("C\u0061\u006c\u0052\u0047B:\u0020I\u006e\u0076\u0061\u006c\u0069d\u0020\u0047\u0061\u006d\u006d\u0061")
		}
		if _bgbc.Len() != 3 {
			return nil, _ee.Errorf("C\u0061\u006c\u0052\u0047\u0042\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0047a\u006d\u006d\u0061 \u0061r\u0072\u0061\u0079")
		}
		_bfda, _dffbc := _bgbc.GetAsFloat64Slice()
		if _dffbc != nil {
			return nil, _dffbc
		}
		_gbbfd.Gamma = _bfda
	}
	_egbcg = _cabc.Get("\u004d\u0061\u0074\u0072\u0069\u0078")
	if _egbcg != nil {
		_egbcg = _cde.TraceToDirectObject(_egbcg)
		_eeaf, _gfafb := _egbcg.(*_cde.PdfObjectArray)
		if !_gfafb {
			return nil, _ee.Errorf("\u0043\u0061\u006c\u0052GB\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004d\u0061\u0074\u0072i\u0078")
		}
		if _eeaf.Len() != 9 {
			_ad.Log.Error("\u004d\u0061t\u0072\u0069\u0078 \u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _eeaf.String())
			return nil, _ee.Errorf("\u0043\u0061\u006c\u0052G\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u004da\u0074\u0072\u0069\u0078\u0020\u0061\u0072r\u0061\u0079")
		}
		_geec, _gdac := _eeaf.GetAsFloat64Slice()
		if _gdac != nil {
			return nil, _gdac
		}
		_gbbfd.Matrix = _geec
	}
	return _gbbfd, nil
}

// HasExtGState checks if ExtGState name is available.
func (_dgafdg *PdfPage) HasExtGState(name _cde.PdfObjectName) bool {
	if _dgafdg.Resources == nil {
		return false
	}
	if _dgafdg.Resources.ExtGState == nil {
		return false
	}
	_aaccg, _eccd := _cde.TraceToDirectObject(_dgafdg.Resources.ExtGState).(*_cde.PdfObjectDictionary)
	if !_eccd {
		_ad.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _cde.TraceToDirectObject(_dgafdg.Resources.ExtGState))
		return false
	}
	_gadg := _aaccg.Get(name)
	_eccca := _gadg != nil
	return _eccca
}

// PdfActionRendition represents a Rendition action.
type PdfActionRendition struct {
	*PdfAction
	R  _cde.PdfObject
	AN _cde.PdfObject
	OP _cde.PdfObject
	JS _cde.PdfObject
}

// SetCatalogMetadata sets the catalog metadata (XMP) stream object.
func (_dddfd *PdfWriter) SetCatalogMetadata(meta _cde.PdfObject) error {
	if meta == nil {
		_dddfd._fedbb.Remove("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
		return nil
	}
	_fdcfd, _cdbdd := _cde.GetStream(meta)
	if !_cdbdd {
		return _ceg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006d\u0065\u0074\u0061\u0064a\u0074\u0061\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0073t\u0072\u0065\u0061\u006d")
	}
	_dddfd.addObject(_fdcfd)
	_dddfd._fedbb.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _fdcfd)
	return nil
}
func (_gdeff *pdfFontType3) getFontDescriptor() *PdfFontDescriptor { return _gdeff._fagf }

// NewPdfSignatureReferenceDocMDP returns PdfSignatureReference for the transformParams.
func NewPdfSignatureReferenceDocMDP(transformParams *PdfTransformParamsDocMDP) *PdfSignatureReference {
	return &PdfSignatureReference{Type: _cde.MakeName("\u0053\u0069\u0067\u0052\u0065\u0066"), TransformMethod: _cde.MakeName("\u0044\u006f\u0063\u004d\u0044\u0050"), TransformParams: transformParams.ToPdfObject()}
}
func (_babbd *PdfWriter) optimize() error {
	if _babbd._cgee == nil {
		return nil
	}
	var _cfcbcg error
	_babbd._egbccc, _cfcbcg = _babbd._cgee.Optimize(_babbd._egbccc)
	if _cfcbcg != nil {
		return _cfcbcg
	}
	_eedeb := make(map[_cde.PdfObject]struct{}, len(_babbd._egbccc))
	for _, _afcgc := range _babbd._egbccc {
		_eedeb[_afcgc] = struct{}{}
	}
	_babbd._bccde = _eedeb
	return nil
}

// DecodeArray returns the range of color component values in CalRGB colorspace.
func (_aaadc *PdfColorspaceCalRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}
func (_eecbe *LTV) validateSig(_efdcg *PdfSignature) error {
	if _efdcg == nil || _efdcg.Contents == nil || len(_efdcg.Contents.Bytes()) == 0 {
		return _ee.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0066\u0069\u0065l\u0064:\u0020\u0025\u0076", _efdcg)
	}
	return nil
}

// PdfWriter handles outputing PDF content.
type PdfWriter struct {
	_eacge         *_cde.PdfIndirectObject
	_acbgeg        *_cde.PdfIndirectObject
	_fbeeb         map[_cde.PdfObject]struct{}
	_egbccc        []_cde.PdfObject
	_bccde         map[_cde.PdfObject]struct{}
	_ggbfg         []*_cde.PdfIndirectObject
	_bdfce         *PdfOutlineTreeNode
	_fedbb         *_cde.PdfObjectDictionary
	_gfdf          []_cde.PdfObject
	_fdgbc         *_cde.PdfIndirectObject
	_cefdd         *_b.Writer
	_eddbc         int64
	_fggeef        error
	_ccgbe         *_cde.PdfCrypt
	_dcbae         *_cde.PdfObjectDictionary
	_bdbdb         *_cde.PdfIndirectObject
	_bgacb         *_cde.PdfObjectArray
	_cgdcc         _cde.Version
	_cagac         *bool
	_egffc         map[_cde.PdfObject][]*_cde.PdfObjectDictionary
	_fcacfe        *PdfAcroForm
	_cgee          Optimizer
	_cfaf          StandardApplier
	_gfdac         map[int]crossReference
	_fgca          int64
	ObjNumOffset   int
	_aabfe         bool
	_bdgaa         _cde.XrefTable
	_gbbda         int64
	_fgged         int64
	_dgad          map[_cde.PdfObject]int64
	_gbdfb         map[_cde.PdfObject]struct{}
	_bedaa         string
	_cgdbf         []*PdfOutputIntent
	_fcdff         bool
	_gfegf, _dcace string
}

// NewPdfAcroForm returns a new PdfAcroForm with an intialized container (indirect object).
func NewPdfAcroForm() *PdfAcroForm {
	return &PdfAcroForm{Fields: &[]*PdfField{}, _ecca: _cde.MakeIndirectObject(_cde.MakeDict())}
}

// SetVersion sets the PDF version of the output file.
func (_cdfgd *PdfWriter) SetVersion(majorVersion, minorVersion int) {
	_cdfgd._cgdcc.Major = majorVersion
	_cdfgd._cgdcc.Minor = minorVersion
}

// ImageToRGB converts image in CalGray color space to RGB (A, B, C -> X, Y, Z).
func (_aagb *PdfColorspaceCalGray) ImageToRGB(img Image) (Image, error) {
	_ffea := _cae.NewReader(img.getBase())
	_aeef := _ff.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_cacg := _cae.NewWriter(_aeef)
	_dgcff := _ced.Pow(2, float64(img.BitsPerComponent)) - 1
	_acggf := make([]uint32, 3)
	var (
		_aedbg                              uint32
		ANorm, X, Y, Z, _fbag, _ddge, _deda float64
		_gbbfa                              error
	)
	for {
		_aedbg, _gbbfa = _ffea.ReadSample()
		if _gbbfa == _f.EOF {
			break
		} else if _gbbfa != nil {
			return img, _gbbfa
		}
		ANorm = float64(_aedbg) / _dgcff
		X = _aagb.WhitePoint[0] * _ced.Pow(ANorm, _aagb.Gamma)
		Y = _aagb.WhitePoint[1] * _ced.Pow(ANorm, _aagb.Gamma)
		Z = _aagb.WhitePoint[2] * _ced.Pow(ANorm, _aagb.Gamma)
		_fbag = 3.240479*X + -1.537150*Y + -0.498535*Z
		_ddge = -0.969256*X + 1.875992*Y + 0.041556*Z
		_deda = 0.055648*X + -0.204043*Y + 1.057311*Z
		_fbag = _ced.Min(_ced.Max(_fbag, 0), 1.0)
		_ddge = _ced.Min(_ced.Max(_ddge, 0), 1.0)
		_deda = _ced.Min(_ced.Max(_deda, 0), 1.0)
		_acggf[0] = uint32(_fbag * _dgcff)
		_acggf[1] = uint32(_ddge * _dgcff)
		_acggf[2] = uint32(_deda * _dgcff)
		if _gbbfa = _cacg.WriteSamples(_acggf); _gbbfa != nil {
			return img, _gbbfa
		}
	}
	return _bddb(&_aeef), nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_gffc *PdfColorspaceDeviceGray) ToPdfObject() _cde.PdfObject {
	return _cde.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079")
}

// DefaultImageHandler is the default implementation of the ImageHandler using the standard go library.
type DefaultImageHandler struct{}

// ToPdfObject implements interface PdfModel.
func (_bfea *PdfAnnotationStamp) ToPdfObject() _cde.PdfObject {
	_bfea.PdfAnnotation.ToPdfObject()
	_caba := _bfea._bddg
	_bceg := _caba.PdfObject.(*_cde.PdfObjectDictionary)
	_bfea.PdfAnnotationMarkup.appendToPdfDictionary(_bceg)
	_bceg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0053\u0074\u0061m\u0070"))
	_bceg.SetIfNotNil("\u004e\u0061\u006d\u0065", _bfea.Name)
	return _caba
}

// IsTerminal returns true for terminal fields, false otherwise.
// Terminal fields are fields whose descendants are only widget annotations.
func (_gbdg *PdfField) IsTerminal() bool { return len(_gbdg.Kids) == 0 }

// NewPdfReaderWithOpts creates a new PdfReader for an input io.ReadSeeker interface
// with a ReaderOpts.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderWithOpts(rs _f.ReadSeeker, opts *ReaderOpts) (*PdfReader, error) {
	const _ecdbe = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0057\u0069\u0074\u0068O\u0070\u0074\u0073"
	return _gfceb(rs, opts, true, _ecdbe)
}

// ToPdfObject implements interface PdfModel.
func (_cbbe *PdfAnnotationMovie) ToPdfObject() _cde.PdfObject {
	_cbbe.PdfAnnotation.ToPdfObject()
	_fgfa := _cbbe._bddg
	_aadab := _fgfa.PdfObject.(*_cde.PdfObjectDictionary)
	_aadab.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u004d\u006f\u0076i\u0065"))
	_aadab.SetIfNotNil("\u0054", _cbbe.T)
	_aadab.SetIfNotNil("\u004d\u006f\u0076i\u0065", _cbbe.Movie)
	_aadab.SetIfNotNil("\u0041", _cbbe.A)
	return _fgfa
}

// GetFontByName gets the font specified by keyName. Returns the PdfObject which
// the entry refers to. Returns a bool value indicating whether or not the entry was found.
func (_aagga *PdfPageResources) GetFontByName(keyName _cde.PdfObjectName) (_cde.PdfObject, bool) {
	if _aagga.Font == nil {
		return nil, false
	}
	_fbfba, _fafbe := _cde.TraceToDirectObject(_aagga.Font).(*_cde.PdfObjectDictionary)
	if !_fafbe {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _cde.TraceToDirectObject(_aagga.Font))
		return nil, false
	}
	if _accbfa := _fbfba.Get(keyName); _accbfa != nil {
		return _accbfa, true
	}
	return nil, false
}

// SignatureHandlerDocMDP extends SignatureHandler with the ValidateWithOpts method for checking the DocMDP policy.
type SignatureHandlerDocMDP interface {
	SignatureHandler

	// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser
	// ValidateWithOpts shall contain Validate call
	ValidateWithOpts(_ffcbc *PdfSignature, _aacce Hasher, _dccdd SignatureHandlerDocMDPParams) (SignatureValidationResult, error)
}

func _bfagge() string { return _ad.Version }

// NewXObjectFormFromStream builds the Form XObject from a stream object.
// TODO: Should this be exposed? Consider different access points.
func NewXObjectFormFromStream(stream *_cde.PdfObjectStream) (*XObjectForm, error) {
	_ddbaf := &XObjectForm{}
	_ddbaf._ecfbd = stream
	_ddcaf := *(stream.PdfObjectDictionary)
	_ffedf, _adae := _cde.NewEncoderFromStream(stream)
	if _adae != nil {
		return nil, _adae
	}
	_ddbaf.Filter = _ffedf
	if _gaafd := _ddcaf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _gaafd != nil {
		_cageeb, _facad := _gaafd.(*_cde.PdfObjectName)
		if !_facad {
			return nil, _ceg.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if *_cageeb != "\u0046\u006f\u0072\u006d" {
			_ad.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
			return nil, _ceg.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
		}
	}
	if _bdgdg := _ddcaf.Get("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065"); _bdgdg != nil {
		_ddbaf.FormType = _bdgdg
	}
	if _dgeag := _ddcaf.Get("\u0042\u0042\u006f\u0078"); _dgeag != nil {
		_ddbaf.BBox = _dgeag
	}
	if _gagef := _ddcaf.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _gagef != nil {
		_ddbaf.Matrix = _gagef
	}
	if _efaab := _ddcaf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _efaab != nil {
		_efaab = _cde.TraceToDirectObject(_efaab)
		_eggfg, _bbga := _efaab.(*_cde.PdfObjectDictionary)
		if !_bbga {
			_ad.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0058\u004f\u0062j\u0065c\u0074\u0020\u0046\u006f\u0072\u006d\u0020\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u002c\u0020\u0070\u006f\u0069\u006e\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _cde.ErrTypeError
		}
		_fggae, _cfbgg := NewPdfPageResourcesFromDict(_eggfg)
		if _cfbgg != nil {
			_ad.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0020\u0072\u0065\u0073\u006f\u0075rc\u0065\u0073")
			return nil, _cfbgg
		}
		_ddbaf.Resources = _fggae
		_ad.Log.Trace("\u0046\u006f\u0072\u006d r\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u003a\u0020\u0025\u0023\u0076", _ddbaf.Resources)
	}
	_ddbaf.Group = _ddcaf.Get("\u0047\u0072\u006fu\u0070")
	_ddbaf.Ref = _ddcaf.Get("\u0052\u0065\u0066")
	_ddbaf.MetaData = _ddcaf.Get("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061")
	_ddbaf.PieceInfo = _ddcaf.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o")
	_ddbaf.LastModified = _ddcaf.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064")
	_ddbaf.StructParent = _ddcaf.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_ddbaf.StructParents = _ddcaf.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073")
	_ddbaf.OPI = _ddcaf.Get("\u004f\u0050\u0049")
	_ddbaf.OC = _ddcaf.Get("\u004f\u0043")
	_ddbaf.Name = _ddcaf.Get("\u004e\u0061\u006d\u0065")
	_ddbaf.Stream = stream.Stream
	return _ddbaf, nil
}

// SetFontByName sets the font specified by keyName to the given object.
func (_ccfad *PdfPageResources) SetFontByName(keyName _cde.PdfObjectName, obj _cde.PdfObject) error {
	if _ccfad.Font == nil {
		_ccfad.Font = _cde.MakeDict()
	}
	_dafge, _adebe := _cde.TraceToDirectObject(_ccfad.Font).(*_cde.PdfObjectDictionary)
	if !_adebe {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _cde.TraceToDirectObject(_ccfad.Font))
		return _cde.ErrTypeError
	}
	_dafge.Set(keyName, obj)
	return nil
}

// NewPdfActionNamed returns a new "named" action.
func NewPdfActionNamed() *PdfActionNamed {
	_cceb := NewPdfAction()
	_abb := &PdfActionNamed{}
	_abb.PdfAction = _cceb
	_cceb.SetContext(_abb)
	return _abb
}

// ColorToRGB converts a color in Separation colorspace to RGB colorspace.
func (_daadb *PdfColorspaceSpecialSeparation) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _daadb.AlternateSpace == nil {
		return nil, _ceg.New("\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020c\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	return _daadb.AlternateSpace.ColorToRGB(color)
}

// NewPdfAnnotationPrinterMark returns a new printermark annotation.
func NewPdfAnnotationPrinterMark() *PdfAnnotationPrinterMark {
	_dcgf := NewPdfAnnotation()
	_dgd := &PdfAnnotationPrinterMark{}
	_dgd.PdfAnnotation = _dcgf
	_dcgf.SetContext(_dgd)
	return _dgd
}

// GetContainingPdfObject implements interface PdfModel.
func (_begb *PdfAnnotation) GetContainingPdfObject() _cde.PdfObject { return _begb._bddg }

// Set sets the colorspace corresponding to key. Add to Names if not set.
func (_fgfbf *PdfPageResourcesColorspaces) Set(key _cde.PdfObjectName, val PdfColorspace) {
	if _, _abdcc := _fgfbf.Colorspaces[string(key)]; !_abdcc {
		_fgfbf.Names = append(_fgfbf.Names, string(key))
	}
	_fgfbf.Colorspaces[string(key)] = val
}

var _ pdfFont = (*pdfFontType3)(nil)

// Insert adds a top level outline item in the outline,
// at the specified index.
func (_defge *Outline) Insert(index uint, item *OutlineItem) {
	_adef := uint(len(_defge.Entries))
	if index > _adef {
		index = _adef
	}
	_defge.Entries = append(_defge.Entries[:index], append([]*OutlineItem{item}, _defge.Entries[index:]...)...)
}
func _dgaff(_fgee []byte) ([]byte, error) {
	_agggb := _a.New()
	if _, _dfcb := _f.Copy(_agggb, _ede.NewReader(_fgee)); _dfcb != nil {
		return nil, _dfcb
	}
	return _agggb.Sum(nil), nil
}

// ToGray returns a PdfColorDeviceGray color based on the current RGB color.
func (_dbbdg *PdfColorDeviceRGB) ToGray() *PdfColorDeviceGray {
	_bdbd := 0.3*_dbbdg.R() + 0.59*_dbbdg.G() + 0.11*_dbbdg.B()
	_bdbd = _ced.Min(_ced.Max(_bdbd, 0.0), 1.0)
	return NewPdfColorDeviceGray(_bdbd)
}
func _cdcca(_acdf *XObjectForm) (*PdfRectangle, error) {
	if _bcbgd, _ceff := _acdf.BBox.(*_cde.PdfObjectArray); _ceff {
		_aggg, _deag := NewPdfRectangle(*_bcbgd)
		if _deag != nil {
			return nil, _deag
		}
		if _ddagb, _bfaa := _acdf.Matrix.(*_cde.PdfObjectArray); _bfaa {
			_cecd, _cfee := _ddagb.ToFloat64Array()
			if _cfee != nil {
				return nil, _cfee
			}
			_gfgc := _adb.IdentityMatrix()
			if len(_cecd) == 6 {
				_gfgc = _adb.NewMatrix(_cecd[0], _cecd[1], _cecd[2], _cecd[3], _cecd[4], _cecd[5])
			}
			_aggg.Transform(_gfgc)
			return _aggg, nil
		}
	}
	return nil, _ceg.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063e\u0020\u0042\u0042\u006f\u0078\u0020\u0074y\u0070\u0065")
}
func (_dgdd *PdfColorspaceCalGray) String() string { return "\u0043a\u006c\u0047\u0072\u0061\u0079" }

// GetAllContentStreams gets all the content streams for a page as one string.
func (_fdeaa *PdfPage) GetAllContentStreams() (string, error) {
	_fgcfe, _aaage := _fdeaa.GetContentStreams()
	if _aaage != nil {
		return "", _aaage
	}
	return _dac.Join(_fgcfe, "\u0020"), nil
}

// GetParamsDict returns *core.PdfObjectDictionary with a set of basic image parameters.
func (_efeeea *Image) GetParamsDict() *_cde.PdfObjectDictionary {
	_cfgfb := _cde.MakeDict()
	_cfgfb.Set("\u0057\u0069\u0064t\u0068", _cde.MakeInteger(_efeeea.Width))
	_cfgfb.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _cde.MakeInteger(_efeeea.Height))
	_cfgfb.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _cde.MakeInteger(int64(_efeeea.ColorComponents)))
	_cfgfb.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _cde.MakeInteger(_efeeea.BitsPerComponent))
	return _cfgfb
}

// GetAlphabet returns a map of the runes in `text` and their frequencies.
func GetAlphabet(text string) map[rune]int {
	_bdbbaa := map[rune]int{}
	for _, _afdf := range text {
		_bdbbaa[_afdf]++
	}
	return _bdbbaa
}

// GetContainingPdfObject returns the container of the pattern object (indirect object).
func (_cafb *PdfPattern) GetContainingPdfObject() _cde.PdfObject { return _cafb._eecac }
func (_egbbe *LTV) buildCertChain(_dbbgad, _bfdec []*_bg.Certificate) ([]*_bg.Certificate, map[string]*_bg.Certificate, error) {
	_efadd := map[string]*_bg.Certificate{}
	for _, _cfae := range _dbbgad {
		_efadd[_cfae.Subject.CommonName] = _cfae
	}
	_cefd := _dbbgad
	for _, _fefee := range _bfdec {
		_bfbbb := _fefee.Subject.CommonName
		if _, _gcacg := _efadd[_bfbbb]; _gcacg {
			continue
		}
		_efadd[_bfbbb] = _fefee
		_cefd = append(_cefd, _fefee)
	}
	if len(_cefd) == 0 {
		return nil, nil, ErrSignNoCertificates
	}
	var _agadd error
	for _bedee := _cefd[0]; _bedee != nil && !_egbbe.CertClient.IsCA(_bedee); {
		_agcd, _ageea := _efadd[_bedee.Issuer.CommonName]
		if !_ageea {
			if _agcd, _agadd = _egbbe.CertClient.GetIssuer(_bedee); _agadd != nil {
				_ad.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _agadd)
				break
			}
			_efadd[_bedee.Issuer.CommonName] = _agcd
			_cefd = append(_cefd, _agcd)
		}
		_bedee = _agcd
	}
	return _cefd, _efadd, nil
}

// NewPdfActionMovie returns a new "movie" action.
func NewPdfActionMovie() *PdfActionMovie {
	_af := NewPdfAction()
	_dgg := &PdfActionMovie{}
	_dgg.PdfAction = _af
	_af.SetContext(_dgg)
	return _dgg
}

// PdfActionHide represents a hide action.
type PdfActionHide struct {
	*PdfAction
	T _cde.PdfObject
	H _cde.PdfObject
}

// AddContentStreamByString adds content stream by string. Puts the content
// string into a stream object and points the content stream towards it.
func (_dbgb *PdfPage) AddContentStreamByString(contentStr string) error {
	_ccgg, _accf := _cde.MakeStream([]byte(contentStr), _cde.NewFlateEncoder())
	if _accf != nil {
		return _accf
	}
	if _dbgb.Contents == nil {
		_dbgb.Contents = _ccgg
	} else {
		_fdfg := _cde.TraceToDirectObject(_dbgb.Contents)
		_cbgfd, _bcbcc := _fdfg.(*_cde.PdfObjectArray)
		if !_bcbcc {
			_cbgfd = _cde.MakeArray(_fdfg)
		}
		_cbgfd.Append(_ccgg)
		_dbgb.Contents = _cbgfd
	}
	return nil
}

// String returns a string describing the font descriptor.
func (_ebecf *PdfFontDescriptor) String() string {
	var _eccg []string
	if _ebecf.FontName != nil {
		_eccg = append(_eccg, _ebecf.FontName.String())
	}
	if _ebecf.FontFamily != nil {
		_eccg = append(_eccg, _ebecf.FontFamily.String())
	}
	if _ebecf.fontFile != nil {
		_eccg = append(_eccg, _ebecf.fontFile.String())
	}
	if _ebecf._ggdga != nil {
		_eccg = append(_eccg, _ebecf._ggdga.String())
	}
	_eccg = append(_eccg, _ee.Sprintf("\u0046\u006f\u006et\u0046\u0069\u006c\u0065\u0033\u003d\u0025\u0074", _ebecf.FontFile3 != nil))
	return _ee.Sprintf("\u0046\u004f\u004e\u0054_D\u0045\u0053\u0043\u0052\u0049\u0050\u0054\u004f\u0052\u007b\u0025\u0073\u007d", _dac.Join(_eccg, "\u002c\u0020"))
}

// DecodeArray returns the range of color component values in DeviceRGB colorspace.
func (_eaed *PdfColorspaceDeviceRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// AlphaMap performs mapping of alpha data for transformations. Allows custom filtering of alpha data etc.
func (_dfbffc *Image) AlphaMap(mapFunc AlphaMapFunc) {
	for _cfagbd, _abfbe := range _dfbffc._deegf {
		_dfbffc._deegf[_cfagbd] = mapFunc(_abfbe)
	}
}

// SetShadingByName sets a shading resource specified by keyName.
func (_bcbfdf *PdfPageResources) SetShadingByName(keyName _cde.PdfObjectName, shadingObj _cde.PdfObject) error {
	if _bcbfdf.Shading == nil {
		_bcbfdf.Shading = _cde.MakeDict()
	}
	_dcfbc, _afae := _bcbfdf.Shading.(*_cde.PdfObjectDictionary)
	if !_afae {
		return _cde.ErrTypeError
	}
	_dcfbc.Set(keyName, shadingObj)
	return nil
}
func (_gbde *PdfColorspaceDeviceGray) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
}

// NewXObjectImageFromImage creates a new XObject Image from an image object
// with default options. If encoder is nil, uses raw encoding (none).
func NewXObjectImageFromImage(img *Image, cs PdfColorspace, encoder _cde.StreamEncoder) (*XObjectImage, error) {
	_aaabd := NewXObjectImage()
	return UpdateXObjectImageFromImage(_aaabd, img, cs, encoder)
}

// ToPdfObject implements interface PdfModel.
func (_bcca *PdfAnnotationPolyLine) ToPdfObject() _cde.PdfObject {
	_bcca.PdfAnnotation.ToPdfObject()
	_cedd := _bcca._bddg
	_egdbg := _cedd.PdfObject.(*_cde.PdfObjectDictionary)
	_bcca.PdfAnnotationMarkup.appendToPdfDictionary(_egdbg)
	_egdbg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065"))
	_egdbg.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _bcca.Vertices)
	_egdbg.SetIfNotNil("\u004c\u0045", _bcca.LE)
	_egdbg.SetIfNotNil("\u0042\u0053", _bcca.BS)
	_egdbg.SetIfNotNil("\u0049\u0043", _bcca.IC)
	_egdbg.SetIfNotNil("\u0042\u0045", _bcca.BE)
	_egdbg.SetIfNotNil("\u0049\u0054", _bcca.IT)
	_egdbg.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _bcca.Measure)
	return _cedd
}

// UpdateXObjectImageFromImage creates a new XObject Image from an
// Image object `img` and default masks from xobjIn.
// The default masks are overridden if img.hasAlpha
// If `encoder` is nil, uses raw encoding (none).
func UpdateXObjectImageFromImage(xobjIn *XObjectImage, img *Image, cs PdfColorspace, encoder _cde.StreamEncoder) (*XObjectImage, error) {
	if encoder == nil {
		encoder = _cde.NewRawEncoder()
	}
	encoder.UpdateParams(img.GetParamsDict())
	_cfacg, _caefc := encoder.EncodeBytes(img.Data)
	if _caefc != nil {
		_ad.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _caefc)
		return nil, _caefc
	}
	_ffcbg := NewXObjectImage()
	_egbae := img.Width
	_dbgfa := img.Height
	_ffcbg.Width = &_egbae
	_ffcbg.Height = &_dbgfa
	_eggdb := img.BitsPerComponent
	_ffcbg.BitsPerComponent = &_eggdb
	_ffcbg.Filter = encoder
	_ffcbg.Stream = _cfacg
	if cs == nil {
		if img.ColorComponents == 1 {
			_ffcbg.ColorSpace = NewPdfColorspaceDeviceGray()
		} else if img.ColorComponents == 3 {
			_ffcbg.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			_ffcbg.ColorSpace = NewPdfColorspaceDeviceCMYK()
		} else {
			return nil, _ceg.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_ffcbg.ColorSpace = cs
	}
	if len(img._deegf) != 0 {
		_dacfff := NewXObjectImage()
		_dacfff.Filter = encoder
		_aecf, _degag := encoder.EncodeBytes(img._deegf)
		if _degag != nil {
			_ad.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _degag)
			return nil, _degag
		}
		_dacfff.Stream = _aecf
		_dacfff.BitsPerComponent = _ffcbg.BitsPerComponent
		_dacfff.Width = &img.Width
		_dacfff.Height = &img.Height
		_dacfff.ColorSpace = NewPdfColorspaceDeviceGray()
		_ffcbg.SMask = _dacfff.ToPdfObject()
	} else {
		_ffcbg.SMask = xobjIn.SMask
		_ffcbg.ImageMask = xobjIn.ImageMask
		if _ffcbg.ColorSpace.GetNumComponents() == 1 {
			_aadfe(_ffcbg)
		}
	}
	return _ffcbg, nil
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

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a CalGray device.
func (_gabf *PdfColorspaceCalGray) GetNumComponents() int { return 1 }

// GetNumComponents returns the number of color components (4 for CMYK32).
func (_bgdbf *PdfColorDeviceCMYK) GetNumComponents() int { return 4 }

// NewPdfPageResourcesColorspaces returns a new PdfPageResourcesColorspaces object.
func NewPdfPageResourcesColorspaces() *PdfPageResourcesColorspaces {
	_ebdge := &PdfPageResourcesColorspaces{}
	_ebdge.Names = []string{}
	_ebdge.Colorspaces = map[string]PdfColorspace{}
	_ebdge._ccdd = &_cde.PdfIndirectObject{}
	return _ebdge
}
func (_dbga *PdfReader) newPdfActionGoTo3DViewFromDict(_bbee *_cde.PdfObjectDictionary) (*PdfActionGoTo3DView, error) {
	return &PdfActionGoTo3DView{TA: _bbee.Get("\u0054\u0041"), V: _bbee.Get("\u0056")}, nil
}

// PdfActionSetOCGState represents a SetOCGState action.
type PdfActionSetOCGState struct {
	*PdfAction
	State      _cde.PdfObject
	PreserveRB _cde.PdfObject
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_cfgff pdfFontType3) GetCharMetrics(code _gc.CharCode) (_fe.CharMetrics, bool) {
	if _gagd, _dbafg := _cfgff._dedg[code]; _dbafg {
		return _fe.CharMetrics{Wx: _gagd}, true
	}
	if _fe.IsStdFont(_fe.StdFontName(_cfgff._eeab)) {
		return _fe.CharMetrics{Wx: 250}, true
	}
	return _fe.CharMetrics{}, false
}

// SetAction sets the PDF action for the annotation link.
func (_cec *PdfAnnotationLink) SetAction(action *PdfAction) {
	_cec._beg = action
	if action == nil {
		_cec.A = nil
	}
}

// Encoder returns the font's text encoder.
func (_cadb pdfCIDFontType0) Encoder() _gc.TextEncoder { return _cadb._efdaa }
func (_gggf *PdfColorspaceSpecialPattern) String() string {
	return "\u0050a\u0074\u0074\u0065\u0072\u006e"
}
func (_bfacaa *PdfWriter) writeObjectsInStreams(_bbbgaf map[_cde.PdfObject]bool) error {
	for _, _bdbdc := range _bfacaa._egbccc {
		if _gaag := _bbbgaf[_bdbdc]; _gaag {
			continue
		}
		_adfaa := int64(0)
		switch _gcgba := _bdbdc.(type) {
		case *_cde.PdfIndirectObject:
			_adfaa = _gcgba.ObjectNumber
		case *_cde.PdfObjectStream:
			_adfaa = _gcgba.ObjectNumber
		case *_cde.PdfObjectStreams:
			_adfaa = _gcgba.ObjectNumber
		default:
			_ad.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0055n\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0074\u0079\u0070\u0065 \u0069\u006e\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0054", _bdbdc)
			return ErrTypeCheck
		}
		if _bfacaa._ccgbe != nil && _bdbdc != _bfacaa._bdbdb {
			_efcgg := _bfacaa._ccgbe.Encrypt(_bdbdc, _adfaa, 0)
			if _efcgg != nil {
				_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067\u0020(%\u0073\u0029", _efcgg)
				return _efcgg
			}
		}
		_bfacaa.writeObject(int(_adfaa), _bdbdc)
	}
	return nil
}
func (_dgfb *PdfAppender) mergeResources(_ebggc, _dgef _cde.PdfObject, _cdee map[_cde.PdfObjectName]_cde.PdfObjectName) _cde.PdfObject {
	if _dgef == nil && _ebggc == nil {
		return nil
	}
	if _dgef == nil {
		return _ebggc
	}
	_beda, _cad := _cde.GetDict(_dgef)
	if !_cad {
		return _ebggc
	}
	if _ebggc == nil {
		_dfbfc := _cde.MakeDict()
		_dfbfc.Merge(_beda)
		return _dgef
	}
	_dde, _cad := _cde.GetDict(_ebggc)
	if !_cad {
		_ad.Log.Error("\u0045\u0072\u0072or\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		_dde = _cde.MakeDict()
	}
	for _, _fgbg := range _beda.Keys() {
		if _cfed, _fadg := _cdee[_fgbg]; _fadg {
			_dde.Set(_cfed, _beda.Get(_fgbg))
		} else {
			_dde.Set(_fgbg, _beda.Get(_fgbg))
		}
	}
	return _dde
}

// AddPage adds a page to the PDF file. The new page should be an indirect object.
func (_afbad *PdfWriter) AddPage(page *PdfPage) error {
	const _daabf = "\u006d\u006f\u0064el\u003a\u0050\u0064\u0066\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_fdcef(page)
	_daffc := page.ToPdfObject()
	_ad.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
	_ad.Log.Trace("\u0041p\u0070\u0065\u006e\u0064i\u006e\u0067\u0020\u0074\u006f \u0070a\u0067e\u0020\u006c\u0069\u0073\u0074\u0020\u0025T", _daffc)
	_cfgea, _fgbfd := _cde.GetIndirect(_daffc)
	if !_fgbfd {
		return _ceg.New("\u0070\u0061\u0067\u0065\u0020\u0073h\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_ad.Log.Trace("\u0025\u0073", _cfgea)
	_ad.Log.Trace("\u0025\u0073", _cfgea.PdfObject)
	_fdcf, _fgbfd := _cde.GetDict(_cfgea.PdfObject)
	if !_fgbfd {
		return _ceg.New("\u0070\u0061\u0067e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068o\u0075l\u0064 \u0062e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_fffea, _fgbfd := _cde.GetName(_fdcf.Get("\u0054\u0079\u0070\u0065"))
	if !_fgbfd {
		return _ee.Errorf("\u0070\u0061\u0067\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054y\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u006e\u0061m\u0065\u0020\u0028%\u0054\u0029", _fdcf.Get("\u0054\u0079\u0070\u0065"))
	}
	if _fffea.String() != "\u0050\u0061\u0067\u0065" {
		return _ceg.New("\u0066\u0069e\u006c\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u0050\u0061\u0067\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069re\u0064\u0029")
	}
	_ecbe := []_cde.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	_ggga, _dabdcc := _cde.GetIndirect(_fdcf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	_ad.Log.Trace("P\u0061g\u0065\u0020\u0050\u0061\u0072\u0065\u006e\u0074:\u0020\u0025\u0054\u0020(%\u0076\u0029", _fdcf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"), _dabdcc)
	for _dabdcc {
		_ad.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _ggga)
		_abebf, _egeda := _cde.GetDict(_ggga.PdfObject)
		if !_egeda {
			return _ceg.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		for _, _beccbea := range _ecbe {
			_ad.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _beccbea)
			if _fdcf.Get(_beccbea) != nil {
				_ad.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
				continue
			}
			if _ecga := _abebf.Get(_beccbea); _ecga != nil {
				_ad.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _beccbea)
				_fdcf.Set(_beccbea, _ecga)
			}
		}
		_ggga, _dabdcc = _cde.GetIndirect(_abebf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
		_ad.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _abebf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	}
	_ad.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c \u0064\u006f\u006e\u0065")
	_fdcf.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _afbad._acbgeg)
	_cfgea.PdfObject = _fdcf
	_cfggf, _fgbfd := _cde.GetDict(_afbad._acbgeg.PdfObject)
	if !_fgbfd {
		return _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_cegd, _fgbfd := _cde.GetArray(_cfggf.Get("\u004b\u0069\u0064\u0073"))
	if !_fgbfd {
		return _ceg.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_cegd.Append(_cfgea)
	_afbad._fbeeb[_fdcf] = struct{}{}
	_edfec, _fgbfd := _cde.GetInt(_cfggf.Get("\u0043\u006f\u0075n\u0074"))
	if !_fgbfd {
		return _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	*_edfec = *_edfec + 1
	_afbad.addObject(_cfgea)
	_aabdf := _afbad.addObjects(_fdcf)
	if _aabdf != nil {
		return _aabdf
	}
	return nil
}
func (_fedabc *Image) samplesTrimPadding(_bdcba []uint32) []uint32 {
	_bgbdg := _fedabc.ColorComponents * int(_fedabc.Width) * int(_fedabc.Height)
	if len(_bdcba) == _bgbdg {
		return _bdcba
	}
	_dade := make([]uint32, _bgbdg)
	_aeaae := int(_fedabc.Width) * _fedabc.ColorComponents
	var _dbbga, _cebde, _dfaef, _dead int
	_cgdg := _ff.BytesPerLine(int(_fedabc.Width), int(_fedabc.BitsPerComponent), _fedabc.ColorComponents)
	for _dbbga = 0; _dbbga < int(_fedabc.Height); _dbbga++ {
		_cebde = _dbbga * int(_fedabc.Width)
		_dfaef = _dbbga * _cgdg
		for _dead = 0; _dead < _aeaae; _dead++ {
			_dade[_cebde+_dead] = _bdcba[_dfaef+_dead]
		}
	}
	return _dade
}

// PdfAnnotationRichMedia represents Rich Media annotations.
type PdfAnnotationRichMedia struct {
	*PdfAnnotation
	RichMediaSettings _cde.PdfObject
	RichMediaContent  _cde.PdfObject
}

// NewPdfActionGoToE returns a new "go to embedded" action.
func NewPdfActionGoToE() *PdfActionGoToE {
	_gff := NewPdfAction()
	_aef := &PdfActionGoToE{}
	_aef.PdfAction = _gff
	_gff.SetContext(_aef)
	return _aef
}

// BytesToCharcodes converts the bytes in a PDF string to character codes.
func (_feagg *PdfFont) BytesToCharcodes(data []byte) []_gc.CharCode {
	_ad.Log.Trace("\u0042\u0079\u0074es\u0054\u006f\u0043\u0068\u0061\u0072\u0063\u006f\u0064e\u0073:\u0020d\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", data, data)
	if _ccbb, _gggfa := _feagg._gbcff.(*pdfFontType0); _gggfa && _ccbb._dcdga != nil {
		if _cafcc, _adcbg := _ccbb.bytesToCharcodes(data); _adcbg {
			return _cafcc
		}
	}
	var (
		_cfcec = make([]_gc.CharCode, 0, len(data)+len(data)%2)
		_aeeab = _feagg.baseFields()
	)
	if _aeeab._ggebg != nil {
		if _bcbde, _bdge := _aeeab._ggebg.BytesToCharcodes(data); _bdge {
			for _, _gadf := range _bcbde {
				_cfcec = append(_cfcec, _gc.CharCode(_gadf))
			}
			return _cfcec
		}
	}
	if _aeeab.isCIDFont() {
		if len(data) == 1 {
			data = []byte{0, data[0]}
		}
		if len(data)%2 != 0 {
			_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0064\u0061\u0074\u0061\u003d\u0025\u002b\u0076\u0020t\u006f\u0020\u0065\u0076\u0065n\u0020\u006ce\u006e\u0067\u0074\u0068", data)
			data = append(data, 0)
		}
		for _geaf := 0; _geaf < len(data); _geaf += 2 {
			_ggaf := uint16(data[_geaf])<<8 | uint16(data[_geaf+1])
			_cfcec = append(_cfcec, _gc.CharCode(_ggaf))
		}
	} else {
		for _, _fgcbd := range data {
			_cfcec = append(_cfcec, _gc.CharCode(_fgcbd))
		}
	}
	return _cfcec
}
func (_gdecc *Image) resampleLowBits(_edbgfb []uint32) {
	_ebbc := _ff.BytesPerLine(int(_gdecc.Width), int(_gdecc.BitsPerComponent), _gdecc.ColorComponents)
	_dgfg := make([]byte, _gdecc.ColorComponents*_ebbc*int(_gdecc.Height))
	_ceab := int(_gdecc.BitsPerComponent) * _gdecc.ColorComponents * int(_gdecc.Width)
	_fcfa := uint8(8)
	var (
		_dgece, _ebae int
		_gdcb         uint32
	)
	for _gedag := 0; _gedag < int(_gdecc.Height); _gedag++ {
		_ebae = _gedag * _ebbc
		for _gegcb := 0; _gegcb < _ceab; _gegcb++ {
			_gdcb = _edbgfb[_dgece]
			_fcfa -= uint8(_gdecc.BitsPerComponent)
			_dgfg[_ebae] |= byte(_gdcb) << _fcfa
			if _fcfa == 0 {
				_fcfa = 8
				_ebae++
			}
			_dgece++
		}
	}
	_gdecc.Data = _dgfg
}
func (_becb *LTV) getCerts(_fbbga []*_bg.Certificate) ([][]byte, error) {
	_edbce := make([][]byte, 0, len(_fbbga))
	for _, _ggfd := range _fbbga {
		_edbce = append(_edbce, _ggfd.Raw)
	}
	return _edbce, nil
}
func (_dfc *PdfReader) newPdfActionLaunchFromDict(_gedd *_cde.PdfObjectDictionary) (*PdfActionLaunch, error) {
	_afbg, _fcb := _beed(_gedd.Get("\u0046"))
	if _fcb != nil {
		return nil, _fcb
	}
	return &PdfActionLaunch{Win: _gedd.Get("\u0057\u0069\u006e"), Mac: _gedd.Get("\u004d\u0061\u0063"), Unix: _gedd.Get("\u0055\u006e\u0069\u0078"), NewWindow: _gedd.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _afbg}, nil
}
func (_dbba *PdfWriter) writeOutputIntents() error {
	if len(_dbba._cgdbf) == 0 {
		return nil
	}
	_dagbe := make([]_cde.PdfObject, len(_dbba._cgdbf))
	for _fdgae, _daage := range _dbba._cgdbf {
		_bcde := _daage.ToPdfObject()
		_dagbe[_fdgae] = _cde.MakeIndirectObject(_bcde)
	}
	_aedgca := _cde.MakeIndirectObject(_cde.MakeArray(_dagbe...))
	_dbba._fedbb.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _aedgca)
	if _bgbb := _dbba.addObjects(_aedgca); _bgbb != nil {
		return _bgbb
	}
	return nil
}

// PdfActionURI represents an URI action.
type PdfActionURI struct {
	*PdfAction
	URI   _cde.PdfObject
	IsMap _cde.PdfObject
}

// GetNumComponents returns the number of color components (1 for grayscale).
func (_abeea *PdfColorDeviceGray) GetNumComponents() int { return 1 }

// DecodeArray returns the range of color component values in the Lab colorspace.
func (_gcea *PdfColorspaceLab) DecodeArray() []float64 {
	_faca := []float64{0, 100}
	if _gcea.Range != nil && len(_gcea.Range) == 4 {
		_faca = append(_faca, _gcea.Range...)
	} else {
		_faca = append(_faca, -100, 100, -100, 100)
	}
	return _faca
}

// PdfShadingType3 is a Radial shading.
type PdfShadingType3 struct {
	*PdfShading
	Coords   *_cde.PdfObjectArray
	Domain   *_cde.PdfObjectArray
	Function []PdfFunction
	Extend   *_cde.PdfObjectArray
}

func _gccb(_agcae *_cde.PdfObjectDictionary) (*PdfFieldChoice, error) {
	_ffgg := &PdfFieldChoice{}
	_ffgg.Opt, _ = _cde.GetArray(_agcae.Get("\u004f\u0070\u0074"))
	_ffgg.TI, _ = _cde.GetInt(_agcae.Get("\u0054\u0049"))
	_ffgg.I, _ = _cde.GetArray(_agcae.Get("\u0049"))
	return _ffgg, nil
}

// ToPdfObject implements interface PdfModel.
func (_eeg *PdfActionSetOCGState) ToPdfObject() _cde.PdfObject {
	_eeg.PdfAction.ToPdfObject()
	_eac := _eeg._bc
	_dbg := _eac.PdfObject.(*_cde.PdfObjectDictionary)
	_dbg.SetIfNotNil("\u0053", _cde.MakeName(string(ActionTypeSetOCGState)))
	_dbg.SetIfNotNil("\u0053\u0074\u0061t\u0065", _eeg.State)
	_dbg.SetIfNotNil("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042", _eeg.PreserveRB)
	return _eac
}

// PdfOutputIntentType is the subtype of the given PdfOutputIntent.
type PdfOutputIntentType int

// SetDSS sets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_gbfe *PdfAppender) SetDSS(dss *DSS) {
	if dss != nil {
		_gbfe.updateObjectsDeep(dss.ToPdfObject(), nil)
	}
	_gbfe._ffac = dss
}

// PdfShadingPattern is a Shading patterns that provide a smooth transition between colors across an area to be painted,
// i.e. color(x,y) = f(x,y) at each point.
// It is a type 2 pattern (PatternType = 2).
type PdfShadingPattern struct {
	*PdfPattern
	Shading   *PdfShading
	Matrix    *_cde.PdfObjectArray
	ExtGState _cde.PdfObject
}

func (_addce *PdfWriter) makeOffSetReference(_egcbda int64) {
	_cgbcd := _ee.Sprintf("\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u000a\u0025\u0064\u000a", _egcbda)
	_addce.writeString(_cgbcd)
	_addce.writeString("\u0025\u0025\u0045\u004f\u0046\u000a")
}

// PdfFieldText represents a text field where user can enter text.
type PdfFieldText struct {
	*PdfField
	DA     *_cde.PdfObjectString
	Q      *_cde.PdfObjectInteger
	DS     *_cde.PdfObjectString
	RV     _cde.PdfObject
	MaxLen *_cde.PdfObjectInteger
}

// NewPdfActionLaunch returns a new "launch" action.
func NewPdfActionLaunch() *PdfActionLaunch {
	_gcb := NewPdfAction()
	_bbea := &PdfActionLaunch{}
	_bbea.PdfAction = _gcb
	_gcb.SetContext(_bbea)
	return _bbea
}

// GetContainingPdfObject implements interface PdfModel.
func (_gcge *PdfSignatureReference) GetContainingPdfObject() _cde.PdfObject { return _gcge._fggee }

// GetContainingPdfObject returns the XObject Form's containing object (indirect object).
func (_fegc *XObjectForm) GetContainingPdfObject() _cde.PdfObject { return _fegc._ecfbd }
func _eebfb(_ggged *_cde.PdfObjectDictionary) (*PdfShadingType2, error) {
	_acef := PdfShadingType2{}
	_agbbcd := _ggged.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _agbbcd == nil {
		_ad.Log.Debug("R\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u003a\u0020\u0020\u0043\u006f\u006f\u0072d\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_eceec, _edbffg := _agbbcd.(*_cde.PdfObjectArray)
	if !_edbffg {
		_ad.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _agbbcd)
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _eceec.Len() != 4 {
		_ad.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0034\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _eceec.Len())
		return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	}
	_acef.Coords = _eceec
	if _befe := _ggged.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _befe != nil {
		_befe = _cde.TraceToDirectObject(_befe)
		_abcf, _edade := _befe.(*_cde.PdfObjectArray)
		if !_edade {
			_ad.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _befe)
			return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_acef.Domain = _abcf
	}
	_agbbcd = _ggged.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _agbbcd == nil {
		_ad.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_acef.Function = []PdfFunction{}
	if _acdad, _ffbd := _agbbcd.(*_cde.PdfObjectArray); _ffbd {
		for _, _adcgc := range _acdad.Elements() {
			_fdacf, _dabgb := _cfdbb(_adcgc)
			if _dabgb != nil {
				_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dabgb)
				return nil, _dabgb
			}
			_acef.Function = append(_acef.Function, _fdacf)
		}
	} else {
		_ddeb, _bcebg := _cfdbb(_agbbcd)
		if _bcebg != nil {
			_ad.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bcebg)
			return nil, _bcebg
		}
		_acef.Function = append(_acef.Function, _ddeb)
	}
	if _acdeg := _ggged.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _acdeg != nil {
		_acdeg = _cde.TraceToDirectObject(_acdeg)
		_faad, _ccebe := _acdeg.(*_cde.PdfObjectArray)
		if !_ccebe {
			_ad.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _acdeg)
			return nil, _cde.ErrTypeError
		}
		if _faad.Len() != 2 {
			_ad.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _faad.Len())
			return nil, ErrInvalidAttribute
		}
		_acef.Extend = _faad
	}
	return &_acef, nil
}

// Fill populates `form` with values provided by `provider`.
func (_bfeee *PdfAcroForm) Fill(provider FieldValueProvider) error { return _bfeee.fill(provider, nil) }

// IsPush returns true if the button field represents a push button, false otherwise.
func (_cbgcf *PdfFieldButton) IsPush() bool { return _cbgcf.GetType() == ButtonTypePush }
func _abaf() string {
	_cdce := "\u0051\u0057\u0045\u0052\u0054\u0059\u0055\u0049\u004f\u0050\u0041S\u0044\u0046\u0047\u0048\u004a\u004b\u004c\u005a\u0058\u0043V\u0042\u004e\u004d"
	var _fadf _ede.Buffer
	for _bccgd := 0; _bccgd < 6; _bccgd++ {
		_fadf.WriteRune(rune(_cdce[_edc.Intn(len(_cdce))]))
	}
	return _fadf.String()
}
func (_eef *PdfReader) newPdfAnnotationSquareFromDict(_baga *_cde.PdfObjectDictionary) (*PdfAnnotationSquare, error) {
	_cbea := PdfAnnotationSquare{}
	_eegd, _eed := _eef.newPdfAnnotationMarkupFromDict(_baga)
	if _eed != nil {
		return nil, _eed
	}
	_cbea.PdfAnnotationMarkup = _eegd
	_cbea.BS = _baga.Get("\u0042\u0053")
	_cbea.IC = _baga.Get("\u0049\u0043")
	_cbea.BE = _baga.Get("\u0042\u0045")
	_cbea.RD = _baga.Get("\u0052\u0044")
	return &_cbea, nil
}
func (_gdfb *PdfAcroForm) signatureFields() []*PdfFieldSignature {
	var _egbebg []*PdfFieldSignature
	for _, _cbfgb := range _gdfb.AllFields() {
		switch _fbafe := _cbfgb.GetContext().(type) {
		case *PdfFieldSignature:
			_eeagg := _fbafe
			_egbebg = append(_egbebg, _eeagg)
		}
	}
	return _egbebg
}

// ConvertToBinary converts current image into binary (bi-level) format.
// Binary images are composed of single bits per pixel (only black or white).
// If provided image has more color components, then it would be converted into binary image using
// histogram auto threshold function.
func (_dafddd *Image) ConvertToBinary() error {
	if _dafddd.ColorComponents == 1 && _dafddd.BitsPerComponent == 1 {
		return nil
	}
	_bdgd, _ageg := _dafddd.ToGoImage()
	if _ageg != nil {
		return _ageg
	}
	_deabe, _ageg := _ff.MonochromeConverter.Convert(_bdgd)
	if _ageg != nil {
		return _ageg
	}
	_dafddd.Data = _deabe.Base().Data
	_dafddd._deegf, _ageg = _ff.ScaleAlphaToMonochrome(_dafddd._deegf, int(_dafddd.Width), int(_dafddd.Height))
	if _ageg != nil {
		return _ageg
	}
	_dafddd.BitsPerComponent = 1
	_dafddd.ColorComponents = 1
	_dafddd._aaafb = nil
	return nil
}

// ToJBIG2Image converts current image to the core.JBIG2Image.
func (_ceeb *Image) ToJBIG2Image() (*_cde.JBIG2Image, error) {
	_cebaf, _aeaba := _ceeb.ToGoImage()
	if _aeaba != nil {
		return nil, _aeaba
	}
	return _cde.GoImageToJBIG2(_cebaf, _cde.JB2ImageAutoThreshold)
}

// GetOutlinesFlattened returns a flattened list of tree nodes and titles.
// NOTE: for most use cases, it is recommended to use the high-level GetOutlines
// method instead, which also provides information regarding the destination
// of the outline items.
func (_edgca *PdfReader) GetOutlinesFlattened() ([]*PdfOutlineTreeNode, []string, error) {
	var _ebdfc []*PdfOutlineTreeNode
	var _cdbfd []string
	var _aeegga func(*PdfOutlineTreeNode, *[]*PdfOutlineTreeNode, *[]string, int)
	_aeegga = func(_ggef *PdfOutlineTreeNode, _ecegf *[]*PdfOutlineTreeNode, _dbbcc *[]string, _ffdgd int) {
		if _ggef == nil {
			return
		}
		if _ggef._fbeea == nil {
			_ad.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020M\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006e\u006fd\u0065\u002e\u0063o\u006et\u0065\u0078\u0074")
			return
		}
		_cdbcc, _efdee := _ggef._fbeea.(*PdfOutlineItem)
		if _efdee {
			*_ecegf = append(*_ecegf, &_cdbcc.PdfOutlineTreeNode)
			_geacf := _dac.Repeat("\u0020", _ffdgd*2) + _cdbcc.Title.Decoded()
			*_dbbcc = append(*_dbbcc, _geacf)
		}
		if _ggef.First != nil {
			_gccg := _dac.Repeat("\u0020", _ffdgd*2) + "\u002b"
			*_dbbcc = append(*_dbbcc, _gccg)
			_aeegga(_ggef.First, _ecegf, _dbbcc, _ffdgd+1)
		}
		if _efdee && _cdbcc.Next != nil {
			_aeegga(_cdbcc.Next, _ecegf, _dbbcc, _ffdgd)
		}
	}
	_aeegga(_edgca._aegbb, &_ebdfc, &_cdbfd, 0)
	return _ebdfc, _cdbfd, nil
}

// ToPdfObject converts colorspace to a PDF object. [/Indexed base hival lookup]
func (_effe *PdfColorspaceSpecialIndexed) ToPdfObject() _cde.PdfObject {
	_dbfce := _cde.MakeArray(_cde.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"))
	_dbfce.Append(_effe.Base.ToPdfObject())
	_dbfce.Append(_cde.MakeInteger(int64(_effe.HiVal)))
	_dbfce.Append(_effe.Lookup)
	if _effe._gcbcb != nil {
		_effe._gcbcb.PdfObject = _dbfce
		return _effe._gcbcb
	}
	return _dbfce
}
func (_ega *PdfReader) newPdfActionFromIndirectObject(_abd *_cde.PdfIndirectObject) (*PdfAction, error) {
	_ged, _dbd := _abd.PdfObject.(*_cde.PdfObjectDictionary)
	if !_dbd {
		return nil, _ee.Errorf("\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u006e\u006f\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _ega._bedfa.GetModelFromPrimitive(_ged); model != nil {
		_ga, _ebf := model.(*PdfAction)
		if !_ebf {
			return nil, _ee.Errorf("\u0063\u0061c\u0068\u0065\u0064\u0020\u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0061\u0063ti\u006f\u006e")
		}
		return _ga, nil
	}
	_eee := &PdfAction{}
	_eee._bc = _abd
	_ega._bedfa.Register(_ged, _eee)
	if _afec := _ged.Get("\u0054\u0079\u0070\u0065"); _afec != nil {
		_abc, _afg := _afec.(*_cde.PdfObjectName)
		if !_afg {
			_ad.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _afec)
		} else {
			if *_abc != "\u0041\u0063\u0074\u0069\u006f\u006e" {
				_ad.Log.Trace("\u0055\u006e\u0073u\u0073\u0070\u0065\u0063t\u0065\u0064\u0020\u0054\u0079\u0070\u0065 \u0021\u003d\u0020\u0041\u0063\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", *_abc)
			}
			_eee.Type = _abc
		}
	}
	if _gdd := _ged.Get("\u004e\u0065\u0078\u0074"); _gdd != nil {
		_eee.Next = _gdd
	}
	if _cagd := _ged.Get("\u0053"); _cagd != nil {
		_eee.S = _cagd
	}
	_cga, _abe := _eee.S.(*_cde.PdfObjectName)
	if !_abe {
		_ad.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _eee.S)
		return nil, _ee.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0053\u0020\u006f\u0062\u006a\u0065c\u0074 \u0074y\u0070e\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _eee.S)
	}
	_afea := PdfActionType(_cga.String())
	switch _afea {
	case ActionTypeGoTo:
		_cgf, _afaa := _ega.newPdfActionGotoFromDict(_ged)
		if _afaa != nil {
			return nil, _afaa
		}
		_cgf.PdfAction = _eee
		_eee._bgd = _cgf
		return _eee, nil
	case ActionTypeGoToR:
		_def, _geg := _ega.newPdfActionGotoRFromDict(_ged)
		if _geg != nil {
			return nil, _geg
		}
		_def.PdfAction = _eee
		_eee._bgd = _def
		return _eee, nil
	case ActionTypeGoToE:
		_dgab, _eca := _ega.newPdfActionGotoEFromDict(_ged)
		if _eca != nil {
			return nil, _eca
		}
		_dgab.PdfAction = _eee
		_eee._bgd = _dgab
		return _eee, nil
	case ActionTypeLaunch:
		_fac, _abab := _ega.newPdfActionLaunchFromDict(_ged)
		if _abab != nil {
			return nil, _abab
		}
		_fac.PdfAction = _eee
		_eee._bgd = _fac
		return _eee, nil
	case ActionTypeThread:
		_ababc, _cbc := _ega.newPdfActionThreadFromDict(_ged)
		if _cbc != nil {
			return nil, _cbc
		}
		_ababc.PdfAction = _eee
		_eee._bgd = _ababc
		return _eee, nil
	case ActionTypeURI:
		_aefc, _fab := _ega.newPdfActionURIFromDict(_ged)
		if _fab != nil {
			return nil, _fab
		}
		_aefc.PdfAction = _eee
		_eee._bgd = _aefc
		return _eee, nil
	case ActionTypeSound:
		_bfa, _cgc := _ega.newPdfActionSoundFromDict(_ged)
		if _cgc != nil {
			return nil, _cgc
		}
		_bfa.PdfAction = _eee
		_eee._bgd = _bfa
		return _eee, nil
	case ActionTypeMovie:
		_acd, _ggb := _ega.newPdfActionMovieFromDict(_ged)
		if _ggb != nil {
			return nil, _ggb
		}
		_acd.PdfAction = _eee
		_eee._bgd = _acd
		return _eee, nil
	case ActionTypeHide:
		_afb, _cee := _ega.newPdfActionHideFromDict(_ged)
		if _cee != nil {
			return nil, _cee
		}
		_afb.PdfAction = _eee
		_eee._bgd = _afb
		return _eee, nil
	case ActionTypeNamed:
		_dcae, _eeb := _ega.newPdfActionNamedFromDict(_ged)
		if _eeb != nil {
			return nil, _eeb
		}
		_dcae.PdfAction = _eee
		_eee._bgd = _dcae
		return _eee, nil
	case ActionTypeSubmitForm:
		_cbg, _fda := _ega.newPdfActionSubmitFormFromDict(_ged)
		if _fda != nil {
			return nil, _fda
		}
		_cbg.PdfAction = _eee
		_eee._bgd = _cbg
		return _eee, nil
	case ActionTypeResetForm:
		_eegg, _eebg := _ega.newPdfActionResetFormFromDict(_ged)
		if _eebg != nil {
			return nil, _eebg
		}
		_eegg.PdfAction = _eee
		_eee._bgd = _eegg
		return _eee, nil
	case ActionTypeImportData:
		_bae, _fc := _ega.newPdfActionImportDataFromDict(_ged)
		if _fc != nil {
			return nil, _fc
		}
		_bae.PdfAction = _eee
		_eee._bgd = _bae
		return _eee, nil
	case ActionTypeSetOCGState:
		_dadb, _gbf := _ega.newPdfActionSetOCGStateFromDict(_ged)
		if _gbf != nil {
			return nil, _gbf
		}
		_dadb.PdfAction = _eee
		_eee._bgd = _dadb
		return _eee, nil
	case ActionTypeRendition:
		_fde, _agd := _ega.newPdfActionRenditionFromDict(_ged)
		if _agd != nil {
			return nil, _agd
		}
		_fde.PdfAction = _eee
		_eee._bgd = _fde
		return _eee, nil
	case ActionTypeTrans:
		_dgb, _cff := _ega.newPdfActionTransFromDict(_ged)
		if _cff != nil {
			return nil, _cff
		}
		_dgb.PdfAction = _eee
		_eee._bgd = _dgb
		return _eee, nil
	case ActionTypeGoTo3DView:
		_efc, _ggg := _ega.newPdfActionGoTo3DViewFromDict(_ged)
		if _ggg != nil {
			return nil, _ggg
		}
		_efc.PdfAction = _eee
		_eee._bgd = _efc
		return _eee, nil
	case ActionTypeJavaScript:
		_ege, _fgb := _ega.newPdfActionJavaScriptFromDict(_ged)
		if _fgb != nil {
			return nil, _fgb
		}
		_ege.PdfAction = _eee
		_eee._bgd = _ege
		return _eee, nil
	}
	_ad.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006eg\u0020u\u006ek\u006eo\u0077\u006e\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073", _afea)
	return nil, nil
}

// Sign signs a specific page with a digital signature.
// The signature field parameter must have a valid signature dictionary
// specified by its V field.
func (_dce *PdfAppender) Sign(pageNum int, field *PdfFieldSignature) error {
	if field == nil {
		return _ceg.New("\u0073\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 n\u0069\u006c")
	}
	_ababf := field.V
	if _ababf == nil {
		return _ceg.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061r\u0079 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_ebee := pageNum - 1
	if _ebee < 0 || _ebee > len(_dce._gfb)-1 {
		return _ee.Errorf("\u0070\u0061\u0067\u0065\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064", pageNum)
	}
	_fgaf := _dce.Reader.PageList[_ebee]
	field.P = _fgaf.ToPdfObject()
	if field.T == nil || field.T.String() == "" {
		field.T = _cde.MakeString(_ee.Sprintf("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072\u0065\u0020\u0025\u0064", pageNum))
	}
	_fgaf.AddAnnotation(field.PdfAnnotationWidget.PdfAnnotation)
	if _dce._cbeaa == _dce._cfad.AcroForm {
		_dce._cbeaa = _dce.Reader.AcroForm
	}
	_ccdc := _dce._cbeaa
	if _ccdc == nil {
		_ccdc = NewPdfAcroForm()
	}
	_ccdc.SigFlags = _cde.MakeInteger(3)
	_gad := append(_ccdc.AllFields(), field.PdfField)
	_ccdc.Fields = &_gad
	_dce.ReplaceAcroForm(_ccdc)
	_dce.UpdatePage(_fgaf)
	_dce._gfb[_ebee] = _fgaf
	if _, _bda := field.V.GetDocMDPPermission(); _bda {
		_dce._bfbe = NewPermissions(field.V)
	}
	return nil
}

// ColorToRGB converts a Lab color to an RGB color.
func (_agaa *PdfColorspaceLab) ColorToRGB(color PdfColor) (PdfColor, error) {
	_dfgaf := func(_fbbf float64) float64 {
		if _fbbf >= 6.0/29 {
			return _fbbf * _fbbf * _fbbf
		}
		return 108.0 / 841 * (_fbbf - 4/29)
	}
	_gecb, _dfcg := color.(*PdfColorLab)
	if !_dfcg {
		_ad.Log.Debug("\u0069\u006e\u0070\u0075t \u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u006c\u0061\u0062")
		return nil, _ceg.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	LStar := _gecb.L()
	AStar := _gecb.A()
	BStar := _gecb.B()
	L := (LStar+16)/116 + AStar/500
	M := (LStar + 16) / 116
	N := (LStar+16)/116 - BStar/200
	X := _agaa.WhitePoint[0] * _dfgaf(L)
	Y := _agaa.WhitePoint[1] * _dfgaf(M)
	Z := _agaa.WhitePoint[2] * _dfgaf(N)
	_aafb := 3.240479*X + -1.537150*Y + -0.498535*Z
	_egdf := -0.969256*X + 1.875992*Y + 0.041556*Z
	_gdcc := 0.055648*X + -0.204043*Y + 1.057311*Z
	_aafb = _ced.Min(_ced.Max(_aafb, 0), 1.0)
	_egdf = _ced.Min(_ced.Max(_egdf, 0), 1.0)
	_gdcc = _ced.Min(_ced.Max(_gdcc, 0), 1.0)
	return NewPdfColorDeviceRGB(_aafb, _egdf, _gdcc), nil
}

// B returns the value of the B component of the color.
func (_fggb *PdfColorCalRGB) B() float64 { return _fggb[1] }

// SubsetRegistered subsets the font to only the glyphs that have been registered by the encoder.
// NOTE: This only works on fonts that support subsetting. For unsupported fonts this is a no-op, although a debug
//   message is emitted.  Currently supported fonts are embedded Truetype CID fonts (type 0).
// NOTE: Make sure to call this soon before writing (once all needed runes have been registered).
// If using package creator, use its EnableFontSubsetting method instead.
func (_dcdc *PdfFont) SubsetRegistered() error {
	switch _dgace := _dcdc._gbcff.(type) {
	case *pdfFontType0:
		_bacce := _dgace.subsetRegistered()
		if _bacce != nil {
			_ad.Log.Debug("\u0053\u0075b\u0073\u0065\u0074 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bacce)
			return _bacce
		}
		if _dgace._dcda != nil {
			if _dgace._cffef != nil {
				_dgace._cffef.ToPdfObject()
			}
			_dgace.ToPdfObject()
		}
	default:
		_ad.Log.Debug("F\u006f\u006e\u0074\u0020\u0025\u0054 \u0064\u006f\u0065\u0073\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006fr\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069n\u0067", _dgace)
	}
	return nil
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_edbf *PdfFunctionType3) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_ad.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return nil, _ceg.New("\u006e\u006f\u0074\u0020im\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
}

// GetModelFromPrimitive returns the model corresponding to the `primitive` PdfObject.
func (_bbdaf *modelManager) GetModelFromPrimitive(primitive _cde.PdfObject) PdfModel {
	model, _gddfb := _bbdaf._dfeee[primitive]
	if !_gddfb {
		return nil
	}
	return model
}

// NewPdfColorspaceLab returns a new Lab colorspace object.
func NewPdfColorspaceLab() *PdfColorspaceLab {
	_edgc := &PdfColorspaceLab{}
	_edgc.BlackPoint = []float64{0.0, 0.0, 0.0}
	_edgc.Range = []float64{-100, 100, -100, 100}
	return _edgc
}

// Height returns the height of `rect`.
func (_dfdge *PdfRectangle) Height() float64 { return _ced.Abs(_dfdge.Ury - _dfdge.Lly) }

// GetNumComponents returns the number of color components (1 for Separation).
func (_eceb *PdfColorspaceSpecialSeparation) GetNumComponents() int { return 1 }

// PageFromIndirectObject returns the PdfPage and page number for a given indirect object.
func (_fbcfc *PdfReader) PageFromIndirectObject(ind *_cde.PdfIndirectObject) (*PdfPage, int, error) {
	if len(_fbcfc.PageList) != len(_fbcfc._gbfgg) {
		return nil, 0, _ceg.New("\u0070\u0061\u0067\u0065\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	for _eeff, _dgbg := range _fbcfc._gbfgg {
		if _dgbg == ind {
			return _fbcfc.PageList[_eeff], _eeff + 1, nil
		}
	}
	return nil, 0, _ceg.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// GetAscent returns the Ascent of the font `descriptor`.
func (_bfed *PdfFontDescriptor) GetAscent() (float64, error) {
	return _cde.GetNumberAsFloat(_bfed.Ascent)
}

// NewPdfActionSubmitForm returns a new "submit form" action.
func NewPdfActionSubmitForm() *PdfActionSubmitForm {
	_fef := NewPdfAction()
	_fbg := &PdfActionSubmitForm{}
	_fbg.PdfAction = _fef
	_fef.SetContext(_fbg)
	return _fbg
}
func (_cadd *Image) samplesAddPadding(_edac []uint32) []uint32 {
	_fcbd := _ff.BytesPerLine(int(_cadd.Width), int(_cadd.BitsPerComponent), _cadd.ColorComponents) * (8 / int(_cadd.BitsPerComponent))
	_bgdfe := _fcbd * int(_cadd.Height)
	if len(_edac) == _bgdfe {
		return _edac
	}
	_ffba := make([]uint32, _bgdfe)
	_fdagf := int(_cadd.Width) * _cadd.ColorComponents
	for _gdcee := 0; _gdcee < int(_cadd.Height); _gdcee++ {
		_bbcc := _gdcee * int(_cadd.Width)
		_deagc := _gdcee * _fcbd
		for _gdafc := 0; _gdafc < _fdagf; _gdafc++ {
			_ffba[_deagc+_gdafc] = _edac[_bbcc+_gdafc]
		}
	}
	return _ffba
}

// ReplaceAcroForm replaces the acrobat form. It appends a new form to the Pdf which
// replaces the original AcroForm.
func (_dgfc *PdfAppender) ReplaceAcroForm(acroForm *PdfAcroForm) {
	if acroForm != nil {
		_dgfc.updateObjectsDeep(acroForm.ToPdfObject(), nil)
	}
	_dgfc._cbeaa = acroForm
}

// IsTiling specifies if the pattern is a tiling pattern.
func (_fgcd *PdfPattern) IsTiling() bool { return _fgcd.PatternType == 1 }

// GetPageDict converts the Page to a PDF object dictionary.
func (_ecee *PdfPage) GetPageDict() *_cde.PdfObjectDictionary {
	_fgbdb := _ecee._gbbc
	_fgbdb.Clear()
	_fgbdb.Set("\u0054\u0079\u0070\u0065", _cde.MakeName("\u0050\u0061\u0067\u0065"))
	_fgbdb.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _ecee.Parent)
	if _ecee.LastModified != nil {
		_fgbdb.Set("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _ecee.LastModified.ToPdfObject())
	}
	if _ecee.Resources != nil {
		_fgbdb.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _ecee.Resources.ToPdfObject())
	}
	if _ecee.CropBox != nil {
		_fgbdb.Set("\u0043r\u006f\u0070\u0042\u006f\u0078", _ecee.CropBox.ToPdfObject())
	}
	if _ecee.MediaBox != nil {
		_fgbdb.Set("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", _ecee.MediaBox.ToPdfObject())
	}
	if _ecee.BleedBox != nil {
		_fgbdb.Set("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078", _ecee.BleedBox.ToPdfObject())
	}
	if _ecee.TrimBox != nil {
		_fgbdb.Set("\u0054r\u0069\u006d\u0042\u006f\u0078", _ecee.TrimBox.ToPdfObject())
	}
	if _ecee.ArtBox != nil {
		_fgbdb.Set("\u0041\u0072\u0074\u0042\u006f\u0078", _ecee.ArtBox.ToPdfObject())
	}
	_fgbdb.SetIfNotNil("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f", _ecee.BoxColorInfo)
	_fgbdb.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _ecee.Contents)
	if _ecee.Rotate != nil {
		_fgbdb.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _cde.MakeInteger(*_ecee.Rotate))
	}
	_fgbdb.SetIfNotNil("\u0047\u0072\u006fu\u0070", _ecee.Group)
	_fgbdb.SetIfNotNil("\u0054\u0068\u0075m\u0062", _ecee.Thumb)
	_fgbdb.SetIfNotNil("\u0042", _ecee.B)
	_fgbdb.SetIfNotNil("\u0044\u0075\u0072", _ecee.Dur)
	_fgbdb.SetIfNotNil("\u0054\u0072\u0061n\u0073", _ecee.Trans)
	_fgbdb.SetIfNotNil("\u0041\u0041", _ecee.AA)
	_fgbdb.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _ecee.Metadata)
	_fgbdb.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _ecee.PieceInfo)
	_fgbdb.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _ecee.StructParents)
	_fgbdb.SetIfNotNil("\u0049\u0044", _ecee.ID)
	_fgbdb.SetIfNotNil("\u0050\u005a", _ecee.PZ)
	_fgbdb.SetIfNotNil("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f", _ecee.SeparationInfo)
	_fgbdb.SetIfNotNil("\u0054\u0061\u0062\u0073", _ecee.Tabs)
	_fgbdb.SetIfNotNil("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064", _ecee.TemplateInstantiated)
	_fgbdb.SetIfNotNil("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s", _ecee.PresSteps)
	_fgbdb.SetIfNotNil("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074", _ecee.UserUnit)
	_fgbdb.SetIfNotNil("\u0056\u0050", _ecee.VP)
	if _ecee._cefe != nil {
		_fgeg := _cde.MakeArray()
		for _, _gagec := range _ecee._cefe {
			if _bbabc := _gagec.GetContext(); _bbabc != nil {
				_fgeg.Append(_bbabc.ToPdfObject())
			} else {
				_fgeg.Append(_gagec.ToPdfObject())
			}
		}
		if _fgeg.Len() > 0 {
			_fgbdb.Set("\u0041\u006e\u006e\u006f\u0074\u0073", _fgeg)
		}
	} else if _ecee.Annots != nil {
		_fgbdb.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0073", _ecee.Annots)
	}
	return _fgbdb
}

// StandardImplementer is an interface that defines specified PDF standards like PDF/A-1A (pdfa.Profile1A)
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
type StandardImplementer interface {
	StandardValidator
	StandardApplier

	// StandardName gets the human-readable name of the standard.
	StandardName() string
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_gcga *PdfOutlineItem) ToPdfObject() _cde.PdfObject {
	_bfdf := _gcga._ccbf
	_cffa := _bfdf.PdfObject.(*_cde.PdfObjectDictionary)
	_cffa.Set("\u0054\u0069\u0074l\u0065", _gcga.Title)
	if _gcga.A != nil {
		_cffa.Set("\u0041", _gcga.A)
	}
	if _aabgg := _cffa.Get("\u0053\u0045"); _aabgg != nil {
		_cffa.Remove("\u0053\u0045")
	}
	if _gcga.C != nil {
		_cffa.Set("\u0043", _gcga.C)
	}
	if _gcga.Dest != nil {
		_cffa.Set("\u0044\u0065\u0073\u0074", _gcga.Dest)
	}
	if _gcga.F != nil {
		_cffa.Set("\u0046", _gcga.F)
	}
	if _gcga.Count != nil {
		_cffa.Set("\u0043\u006f\u0075n\u0074", _cde.MakeInteger(*_gcga.Count))
	}
	if _gcga.Next != nil {
		_cffa.Set("\u004e\u0065\u0078\u0074", _gcga.Next.ToPdfObject())
	}
	if _gcga.First != nil {
		_cffa.Set("\u0046\u0069\u0072s\u0074", _gcga.First.ToPdfObject())
	}
	if _gcga.Prev != nil {
		_cffa.Set("\u0050\u0072\u0065\u0076", _gcga.Prev.GetContext().GetContainingPdfObject())
	}
	if _gcga.Last != nil {
		_cffa.Set("\u004c\u0061\u0073\u0074", _gcga.Last.GetContext().GetContainingPdfObject())
	}
	if _gcga.Parent != nil {
		_cffa.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _gcga.Parent.GetContext().GetContainingPdfObject())
	}
	return _bfdf
}
func (_bbdd *PdfWriter) checkPendingObjects() {
	for _eadc, _agcf := range _bbdd._egffc {
		if !_bbdd.hasObject(_eadc) {
			_ad.Log.Debug("\u0057\u0041\u0052\u004e\u0020\u0050\u0065n\u0064\u0069\u006eg\u0020\u006f\u0062j\u0065\u0063t\u0020\u0025\u002b\u0076\u0020\u0025T\u0020(%\u0070\u0029\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0077\u0072\u0069\u0074\u0069\u006e\u0067", _eadc, _eadc, _eadc)
			for _, _edbec := range _agcf {
				for _, _efgb := range _edbec.Keys() {
					_ceaeb := _edbec.Get(_efgb)
					if _ceaeb == _eadc {
						_ad.Log.Debug("\u0050e\u006e\u0064i\u006e\u0067\u0020\u006fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0061nd\u0020\u0072\u0065p\u006c\u0061c\u0065\u0064\u0020\u0077\u0069\u0074h\u0020\u006eu\u006c\u006c")
						_edbec.Set(_efgb, _cde.MakeNull())
						break
					}
				}
			}
		}
	}
}

// ToPdfObject returns a PdfObject representation of PdfColorspaceDeviceNAttributes as a PdfObjectDictionary directly
// or indirectly within an indirect object container.
func (_fbad *PdfColorspaceDeviceNAttributes) ToPdfObject() _cde.PdfObject {
	_adag := _cde.MakeDict()
	if _fbad.Subtype != nil {
		_adag.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _fbad.Subtype)
	}
	_adag.SetIfNotNil("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s", _fbad.Colorants)
	_adag.SetIfNotNil("\u0050r\u006f\u0063\u0065\u0073\u0073", _fbad.Process)
	_adag.SetIfNotNil("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073", _fbad.MixingHints)
	if _fbad._deaf != nil {
		_fbad._deaf.PdfObject = _adag
		return _fbad._deaf
	}
	return _adag
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the reader using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
// At the end of the process, the AcroForm contains all the fields which were
// not flattened. If all fields are flattened, the reader's AcroForm field
// is set to nil.
func (_fefbb *PdfReader) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	return _fefbb.flattenFieldsWithOpts(false, appgen, opts)
}

// GetCatalogMarkInfo gets catalog MarkInfo object.
func (_agdfb *PdfReader) GetCatalogMarkInfo() (_cde.PdfObject, bool) {
	if _agdfb._efabe == nil {
		return nil, false
	}
	_gefa := _agdfb._efabe.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
	return _gefa, _gefa != nil
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
	ShadingType *_cde.PdfObjectInteger
	ColorSpace  PdfColorspace
	Background  *_cde.PdfObjectArray
	BBox        *PdfRectangle
	AntiAlias   *_cde.PdfObjectBool
	_dgfac      PdfModel
	_dffg       _cde.PdfObject
}

// PdfAction represents an action in PDF (section 12.6 p. 412).
type PdfAction struct {
	_bgd PdfModel
	Type _cde.PdfObject
	S    _cde.PdfObject
	Next _cde.PdfObject
	_bc  *_cde.PdfIndirectObject
}

// SetPageLabels sets the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_acbcgd *PdfWriter) SetPageLabels(pageLabels _cde.PdfObject) error {
	if pageLabels == nil {
		return nil
	}
	_ad.Log.Trace("\u0053\u0065t\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006cs.\u002e\u002e")
	_acbcgd._fedbb.Set("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073", pageLabels)
	return _acbcgd.addObjects(pageLabels)
}

// ImageToRGB returns the passed in image. Method exists in order to satisfy
// the PdfColorspace interface.
func (_fbdf *PdfColorspaceDeviceRGB) ImageToRGB(img Image) (Image, error) { return img, nil }

// Write writes out the PDF.
func (_badaa *PdfWriter) Write(writer _f.Writer) error {
	_ad.Log.Trace("\u0057r\u0069\u0074\u0065\u0028\u0029")
	_bbgfe := _badaa.checkLicense()
	if _bbgfe != nil {
		return _bbgfe
	}
	if _bbgfe = _badaa.writeOutlines(); _bbgfe != nil {
		return _bbgfe
	}
	if _bbgfe = _badaa.writeAcroFormFields(); _bbgfe != nil {
		return _bbgfe
	}
	_badaa.checkPendingObjects()
	if _bbgfe = _badaa.writeOutputIntents(); _bbgfe != nil {
		return _bbgfe
	}
	_badaa.setCatalogVersion()
	_badaa.copyObjects()
	if _bbgfe = _badaa.optimize(); _bbgfe != nil {
		return _bbgfe
	}
	if _bbgfe = _badaa.optimizeDocument(); _bbgfe != nil {
		return _bbgfe
	}
	var _edfg _g.Hash
	if _badaa._fcdff {
		_edfg = _bf.New()
		writer = _f.MultiWriter(_edfg, writer)
	}
	_badaa.setWriter(writer)
	_abdde := _badaa.checkCrossReferenceStream()
	_cdaf, _abdde := _badaa.mapObjectStreams(_abdde)
	_badaa.adjustXRefAffectedVersion(_abdde)
	_badaa.writeDocumentVersion()
	_badaa.updateObjectNumbers()
	_badaa.writeObjects()
	if _bbgfe = _badaa.writeObjectsInStreams(_cdaf); _bbgfe != nil {
		return _bbgfe
	}
	_egabaa := _badaa._eddbc
	var _gefede int
	for _gafef := range _badaa._gfdac {
		if _gafef > _gefede {
			_gefede = _gafef
		}
	}
	if _badaa._fcdff {
		if _bbgfe = _badaa.setHashIDs(_edfg); _bbgfe != nil {
			return _bbgfe
		}
	}
	if _abdde {
		if _bbgfe = _badaa.writeXRefStreams(_gefede, _egabaa); _bbgfe != nil {
			return _bbgfe
		}
	} else {
		_badaa.writeTrailer(_gefede)
	}
	_badaa.makeOffSetReference(_egabaa)
	if _bbgfe = _badaa.flushWriter(); _bbgfe != nil {
		return _bbgfe
	}
	return nil
}

var _edba = false

// Normalize swaps (Llx,Urx) if Urx < Llx, and (Lly,Ury) if Ury < Lly.
func (_dagbb *PdfRectangle) Normalize() {
	if _dagbb.Llx > _dagbb.Urx {
		_dagbb.Llx, _dagbb.Urx = _dagbb.Urx, _dagbb.Llx
	}
	if _dagbb.Lly > _dagbb.Ury {
		_dagbb.Lly, _dagbb.Ury = _dagbb.Ury, _dagbb.Lly
	}
}
func (_fadea *PdfField) inherit(_bccf func(*PdfField) bool) (bool, error) {
	_cfac := map[*PdfField]bool{}
	_acbg := false
	_eagda := _fadea
	for _eagda != nil {
		if _, _fffa := _cfac[_eagda]; _fffa {
			return false, _ceg.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		_caeda := _bccf(_eagda)
		if _caeda {
			_acbg = true
			break
		}
		_cfac[_eagda] = true
		_eagda = _eagda.Parent
	}
	return _acbg, nil
}
func (_bgbf *PdfColorspaceCalRGB) String() string { return "\u0043\u0061\u006c\u0052\u0047\u0042" }

// Size returns the width and the height of the page. The method reports
// the page dimensions as displayed by a PDF viewer (i.e. page rotation is
// taken into account).
func (_baff *PdfPage) Size() (float64, float64, error) {
	_bgbfg, _dfadg := _baff.GetMediaBox()
	if _dfadg != nil {
		return 0, 0, _dfadg
	}
	_egbcb, _cabg := _bgbfg.Width(), _bgbfg.Height()
	_ccffb, _dfadg := _baff.GetRotate()
	if _dfadg != nil {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _dfadg.Error())
	}
	if _ccffe := _ccffb; _ccffe%360 != 0 && _ccffe%90 == 0 {
		if _eded := (360 + _ccffe%360) % 360; _eded == 90 || _eded == 270 {
			_egbcb, _cabg = _cabg, _egbcb
		}
	}
	return _egbcb, _cabg, nil
}

// EncryptionAlgorithm is used in EncryptOptions to change the default algorithm used to encrypt the document.
type EncryptionAlgorithm int

var (
	ErrRequiredAttributeMissing = _ceg.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074t\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
	ErrInvalidAttribute         = _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	ErrTypeCheck                = _ceg.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	_cdab                       = _ceg.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrEncrypted                = _ceg.New("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	ErrNoFont                   = _ceg.New("\u0066\u006fn\u0074\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	ErrFontNotSupported         = _beb.Errorf("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u006fn\u0074\u0020\u0028\u0025\u0077\u0029", _cde.ErrNotSupported)
	ErrType1CFontNotSupported   = _beb.Errorf("\u0054y\u0070\u00651\u0043\u0020\u0066o\u006e\u0074\u0073\u0020\u0061\u0072\u0065 \u006e\u006f\u0074\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0028\u0025\u0077\u0029", _cde.ErrNotSupported)
	ErrType3FontNotSupported    = _beb.Errorf("\u0054y\u0070\u00653\u0020\u0066\u006f\u006et\u0073\u0020\u0061r\u0065\u0020\u006e\u006f\u0074\u0020\u0063\u0075\u0072re\u006e\u0074\u006cy\u0020\u0073u\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0028%\u0077\u0029", _cde.ErrNotSupported)
	ErrTTCmapNotSupported       = _beb.Errorf("\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065d\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070e\u0020\u0063\u006d\u0061\u0070\u0020\u0066\u006f\u0072\u006d\u0061\u0074 \u0028\u0025\u0077\u0029", _cde.ErrNotSupported)
	ErrSignNotEnoughSpace       = _beb.Errorf("\u0069\u006e\u0073\u0075\u0066\u0066\u0069c\u0069\u0065\u006et\u0020\u0073\u0070a\u0063\u0065 \u0061\u006c\u006c\u006f\u0063\u0061t\u0065d \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
	ErrSignNoCertificates       = _beb.Errorf("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065tri\u0065ve\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0063\u0068\u0061i\u006e")
)

// StdFontName represents name of a standard font.
type StdFontName = _fe.StdFontName

// DecodeArray returns the range of color component values in DeviceCMYK colorspace.
func (_cfaa *PdfColorspaceDeviceCMYK) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// GetNumComponents returns the number of color components of the underlying
// colorspace device.
func (_cdebg *PdfColorspaceSpecialPattern) GetNumComponents() int {
	return _cdebg.UnderlyingCS.GetNumComponents()
}
func (_gcddgb *PdfWriter) writeString(_bcdgae string) {
	if _gcddgb._fggeef != nil {
		return
	}
	_bgde, _efadg := _gcddgb._cefdd.WriteString(_bcdgae)
	_gcddgb._eddbc += int64(_bgde)
	_gcddgb._fggeef = _efadg
}

// GetPageAsIndirectObject returns the page as a dictionary within an PdfIndirectObject.
func (_gbdbb *PdfPage) GetPageAsIndirectObject() *_cde.PdfIndirectObject { return _gbdbb._dcaeff }

// GetContainingPdfObject returns the page as a dictionary within an PdfIndirectObject.
func (_eabb *PdfPage) GetContainingPdfObject() _cde.PdfObject { return _eabb._dcaeff }

// GetContainingPdfObject implements interface PdfModel.
func (_bee *PdfAction) GetContainingPdfObject() _cde.PdfObject { return _bee._bc }

type modelManager struct {
	_bfaea map[PdfModel]_cde.PdfObject
	_dfeee map[_cde.PdfObject]PdfModel
}

func _dfbeg(_gefc *[]*PdfField, _ddgcg FieldFilterFunc, _dgcd bool) []*PdfField {
	if _gefc == nil {
		return nil
	}
	_dcffg := *_gefc
	if len(*_gefc) == 0 {
		return nil
	}
	_bccd := _dcffg[:0]
	if _ddgcg == nil {
		_ddgcg = func(*PdfField) bool { return true }
	}
	var _aacf []*PdfField
	for _, _eacbbf := range _dcffg {
		_fbcf := _ddgcg(_eacbbf)
		if _fbcf {
			_aacf = append(_aacf, _eacbbf)
			if len(_eacbbf.Kids) > 0 {
				_aacf = append(_aacf, _dfbeg(&_eacbbf.Kids, _ddgcg, _dgcd)...)
			}
		}
		if !_dgcd || !_fbcf || len(_eacbbf.Kids) > 0 {
			_bccd = append(_bccd, _eacbbf)
		}
	}
	*_gefc = _bccd
	return _aacf
}

type fontFile struct {
	_fbffa string
	_bfbba string
	_cafea _gc.SimpleEncoder
}

// SetPatternByName sets a pattern resource specified by keyName.
func (_gegbg *PdfPageResources) SetPatternByName(keyName _cde.PdfObjectName, pattern _cde.PdfObject) error {
	if _gegbg.Pattern == nil {
		_gegbg.Pattern = _cde.MakeDict()
	}
	_ggcf, _eacfb := _gegbg.Pattern.(*_cde.PdfObjectDictionary)
	if !_eacfb {
		return _cde.ErrTypeError
	}
	_ggcf.Set(keyName, pattern)
	return nil
}

type pdfFont interface {
	_fe.Font

	// ToPdfObject returns a PDF representation of the font and implements interface Model.
	ToPdfObject() _cde.PdfObject
	getFontDescriptor() *PdfFontDescriptor
	baseFields() *fontCommon
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_aecc *PdfShadingPattern) ToPdfObject() _cde.PdfObject {
	_aecc.PdfPattern.ToPdfObject()
	_bdef := _aecc.getDict()
	if _aecc.Shading != nil {
		_bdef.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _aecc.Shading.ToPdfObject())
	}
	if _aecc.Matrix != nil {
		_bdef.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _aecc.Matrix)
	}
	if _aecc.ExtGState != nil {
		_bdef.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _aecc.ExtGState)
	}
	return _aecc._eecac
}
func _gbfac(_dgafd *_cde.PdfObjectStream) (*PdfFunctionType0, error) {
	_gabbd := &PdfFunctionType0{}
	_gabbd._fefbf = _dgafd
	_eedae := _dgafd.PdfObjectDictionary
	_gbda, _acbcg := _cde.TraceToDirectObject(_eedae.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_cde.PdfObjectArray)
	if !_acbcg {
		_ad.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _gbda.Len() < 0 || _gbda.Len()%2 != 0 {
		_ad.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _ceg.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_gabbd.NumInputs = _gbda.Len() / 2
	_adage, _gdgg := _gbda.ToFloat64Array()
	if _gdgg != nil {
		return nil, _gdgg
	}
	_gabbd.Domain = _adage
	_gbda, _acbcg = _cde.TraceToDirectObject(_eedae.Get("\u0052\u0061\u006eg\u0065")).(*_cde.PdfObjectArray)
	if !_acbcg {
		_ad.Log.Error("\u0052\u0061\u006e\u0067e \u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _gbda.Len() < 0 || _gbda.Len()%2 != 0 {
		return nil, _ceg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_gabbd.NumOutputs = _gbda.Len() / 2
	_eabga, _gdgg := _gbda.ToFloat64Array()
	if _gdgg != nil {
		return nil, _gdgg
	}
	_gabbd.Range = _eabga
	_gbda, _acbcg = _cde.TraceToDirectObject(_eedae.Get("\u0053\u0069\u007a\u0065")).(*_cde.PdfObjectArray)
	if !_acbcg {
		_ad.Log.Error("\u0053i\u007ae\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_gbacf, _gdgg := _gbda.ToIntegerArray()
	if _gdgg != nil {
		return nil, _gdgg
	}
	if len(_gbacf) != _gabbd.NumInputs {
		_ad.Log.Error("T\u0061\u0062\u006c\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063h\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072 o\u0066\u0020\u0069n\u0070u\u0074\u0073")
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gabbd.Size = _gbacf
	_fdgcf, _acbcg := _cde.TraceToDirectObject(_eedae.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065")).(*_cde.PdfObjectInteger)
	if !_acbcg {
		_ad.Log.Error("B\u0069\u0074\u0073\u0050\u0065\u0072S\u0061\u006d\u0070\u006c\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
		return nil, _ceg.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if *_fdgcf != 1 && *_fdgcf != 2 && *_fdgcf != 4 && *_fdgcf != 8 && *_fdgcf != 12 && *_fdgcf != 16 && *_fdgcf != 24 && *_fdgcf != 32 {
		_ad.Log.Error("\u0042\u0069\u0074s \u0070\u0065\u0072\u0020\u0073\u0061\u006d\u0070\u006ce\u0020o\u0075t\u0073i\u0064\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064\u0029", *_fdgcf)
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gabbd.BitsPerSample = int(*_fdgcf)
	_gabbd.Order = 1
	_cfbe, _acbcg := _cde.TraceToDirectObject(_eedae.Get("\u004f\u0072\u0064e\u0072")).(*_cde.PdfObjectInteger)
	if _acbcg {
		if *_cfbe != 1 && *_cfbe != 3 {
			_ad.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u006fr\u0064\u0065r\u0020\u0028\u0025\u0064\u0029", *_cfbe)
			return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		_gabbd.Order = int(*_cfbe)
	}
	_gbda, _acbcg = _cde.TraceToDirectObject(_eedae.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_cde.PdfObjectArray)
	if _acbcg {
		_fedab, _fefg := _gbda.ToFloat64Array()
		if _fefg != nil {
			return nil, _fefg
		}
		_gabbd.Encode = _fedab
	}
	_gbda, _acbcg = _cde.TraceToDirectObject(_eedae.Get("\u0044\u0065\u0063\u006f\u0064\u0065")).(*_cde.PdfObjectArray)
	if _acbcg {
		_gcfg, _eefaa := _gbda.ToFloat64Array()
		if _eefaa != nil {
			return nil, _eefaa
		}
		_gabbd.Decode = _gcfg
	}
	_ccdf, _gdgg := _cde.DecodeStream(_dgafd)
	if _gdgg != nil {
		return nil, _gdgg
	}
	_gabbd._gfbb = _ccdf
	return _gabbd, nil
}

// NewStandard14Font returns the standard 14 font named `basefont` as a *PdfFont, or an error if it
// `basefont` is not one of the standard 14 font names.
func NewStandard14Font(basefont StdFontName) (*PdfFont, error) {
	_ggag, _afadf := _eabg(basefont)
	if _afadf != nil {
		return nil, _afadf
	}
	if basefont != SymbolName && basefont != ZapfDingbatsName {
		_ggag._efeaf = _gc.NewWinAnsiEncoder()
	}
	return &PdfFont{_gbcff: &_ggag}, nil
}

// ColorFromPdfObjects returns a new PdfColor based on input color components. The input PdfObjects should
// be numeric.
func (_fcgg *PdfColorspaceDeviceN) ColorFromPdfObjects(objects []_cde.PdfObject) (PdfColor, error) {
	if len(objects) != _fcgg.GetNumComponents() {
		return nil, _ceg.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_egdbc, _bead := _cde.GetNumbersAsFloat(objects)
	if _bead != nil {
		return nil, _bead
	}
	return _fcgg.ColorFromFloats(_egdbc)
}

var _ pdfFont = (*pdfFontType0)(nil)

// XObjectImage (Table 89 in 8.9.5.1).
// Implements PdfModel interface.
type XObjectImage struct {

	//ColorSpace       PdfObject
	Width            *int64
	Height           *int64
	ColorSpace       PdfColorspace
	BitsPerComponent *int64
	Filter           _cde.StreamEncoder
	Intent           _cde.PdfObject
	ImageMask        _cde.PdfObject
	Mask             _cde.PdfObject
	Matte            _cde.PdfObject
	Decode           _cde.PdfObject
	Interpolate      _cde.PdfObject
	Alternatives     _cde.PdfObject
	SMask            _cde.PdfObject
	SMaskInData      _cde.PdfObject
	Name             _cde.PdfObject
	StructParent     _cde.PdfObject
	ID               _cde.PdfObject
	OPI              _cde.PdfObject
	Metadata         _cde.PdfObject
	OC               _cde.PdfObject
	Stream           []byte
	_bbaed           *_cde.PdfObjectStream
}

// IsSimple returns true if `font` is a simple font.
func (_bgbcf *PdfFont) IsSimple() bool { _, _bgdgd := _bgbcf._gbcff.(*pdfFontSimple); return _bgdgd }

// PdfAnnotationMarkup represents additional fields for mark-up annotations.
// (Section 12.5.6.2 p. 399).
type PdfAnnotationMarkup struct {
	T            _cde.PdfObject
	Popup        *PdfAnnotationPopup
	CA           _cde.PdfObject
	RC           _cde.PdfObject
	CreationDate _cde.PdfObject
	IRT          _cde.PdfObject
	Subj         _cde.PdfObject
	RT           _cde.PdfObject
	IT           _cde.PdfObject
	ExData       _cde.PdfObject
}

// GetCapHeight returns the CapHeight of the font `descriptor`.
func (_ageb *PdfFontDescriptor) GetCapHeight() (float64, error) {
	return _cde.GetNumberAsFloat(_ageb.CapHeight)
}

// PdfActionNamed represents a named action.
type PdfActionNamed struct {
	*PdfAction
	N _cde.PdfObject
}

func (_addca *PdfSignature) extractChainFromCert() ([]*_bg.Certificate, error) {
	var _fddcf *_cde.PdfObjectArray
	switch _cefef := _addca.Cert.(type) {
	case *_cde.PdfObjectString:
		_fddcf = _cde.MakeArray(_cefef)
	case *_cde.PdfObjectArray:
		_fddcf = _cefef
	default:
		return nil, _ee.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _cefef)
	}
	var _cecce _ede.Buffer
	for _, _fcccf := range _fddcf.Elements() {
		_dccad, _efcdb := _cde.GetString(_fcccf)
		if !_efcdb {
			return nil, _ee.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054", _fcccf)
		}
		if _, _fdafdf := _cecce.Write(_dccad.Bytes()); _fdafdf != nil {
			return nil, _fdafdf
		}
	}
	return _bg.ParseCertificates(_cecce.Bytes())
}

// SetPdfSubject sets the Subject attribute of the output PDF.
func SetPdfSubject(subject string) { _dccfe.Lock(); defer _dccfe.Unlock(); _gbccbb = subject }

// Has checks if flag fl is set in flag and returns true if so, false otherwise.
func (_aabd FieldFlag) Has(fl FieldFlag) bool { return (_aabd.Mask() & fl.Mask()) > 0 }

// ToPdfOutline returns a low level PdfOutline object, based on the current
// instance.
func (_fdgef *Outline) ToPdfOutline() *PdfOutline {
	_adagc := NewPdfOutline()
	var _bbdee []*PdfOutlineItem
	var _fbgc int64
	var _bage *PdfOutlineItem
	for _, _gagae := range _fdgef.Entries {
		_fbgec, _cgcdfe := _gagae.ToPdfOutlineItem()
		_fbgec.Parent = &_adagc.PdfOutlineTreeNode
		if _bage != nil {
			_bage.Next = &_fbgec.PdfOutlineTreeNode
			_fbgec.Prev = &_bage.PdfOutlineTreeNode
		}
		_bbdee = append(_bbdee, _fbgec)
		_fbgc += _cgcdfe
		_bage = _fbgec
	}
	_agfc := int64(len(_bbdee))
	_fbgc += _agfc
	if _agfc > 0 {
		_adagc.First = &_bbdee[0].PdfOutlineTreeNode
		_adagc.Last = &_bbdee[_agfc-1].PdfOutlineTreeNode
		_adagc.Count = &_fbgc
	}
	return _adagc
}
func _cabad(_gcafd _cde.PdfObject) (*fontFile, error) {
	_ad.Log.Trace("\u006e\u0065\u0077\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0046\u0072\u006f\u006dP\u0064f\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u006f\u0062\u006a\u003d\u0025\u0073", _gcafd)
	_agebg := &fontFile{}
	_gcafd = _cde.TraceToDirectObject(_gcafd)
	_bbceb, _fgaba := _gcafd.(*_cde.PdfObjectStream)
	if !_fgaba {
		_ad.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020F\u006f\u006et\u0046\u0069\u006c\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0028\u0025\u0054\u0029", _gcafd)
		return nil, _cde.ErrTypeError
	}
	_bcdcb := _bbceb.PdfObjectDictionary
	_begg, _cddec := _cde.DecodeStream(_bbceb)
	if _cddec != nil {
		return nil, _cddec
	}
	_aabf, _fgaba := _cde.GetNameVal(_bcdcb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_fgaba {
		_agebg._bfbba = _aabf
		if _aabf == "\u0054\u0079\u0070\u0065\u0031\u0043" {
			_ad.Log.Debug("T\u0079\u0070\u0065\u0031\u0043\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u006e\u006f\u0074 s\u0075\u0070\u0070o\u0072t\u0065\u0064")
			return nil, ErrType1CFontNotSupported
		}
	}
	_egaf, _ := _cde.GetIntVal(_bcdcb.Get("\u004ce\u006e\u0067\u0074\u0068\u0031"))
	_ddgd, _ := _cde.GetIntVal(_bcdcb.Get("\u004ce\u006e\u0067\u0074\u0068\u0032"))
	if _egaf > len(_begg) {
		_egaf = len(_begg)
	}
	if _egaf+_ddgd > len(_begg) {
		_ddgd = len(_begg) - _egaf
	}
	_dfdda := _begg[:_egaf]
	var _dedeg []byte
	if _ddgd > 0 {
		_dedeg = _begg[_egaf : _egaf+_ddgd]
	}
	if _egaf > 0 && _ddgd > 0 {
		_ffdf := _agebg.loadFromSegments(_dfdda, _dedeg)
		if _ffdf != nil {
			return nil, _ffdf
		}
	}
	return _agebg, nil
}

// SetPdfCreationDate sets the CreationDate attribute of the output PDF.
func SetPdfCreationDate(creationDate _ce.Time) {
	_dccfe.Lock()
	defer _dccfe.Unlock()
	_gffce = creationDate
}

// ToPdfObject implements interface PdfModel.
func (_cbeb *PdfAnnotationPolygon) ToPdfObject() _cde.PdfObject {
	_cbeb.PdfAnnotation.ToPdfObject()
	_cbfg := _cbeb._bddg
	_bdg := _cbfg.PdfObject.(*_cde.PdfObjectDictionary)
	_cbeb.PdfAnnotationMarkup.appendToPdfDictionary(_bdg)
	_bdg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _cde.MakeName("\u0050o\u006c\u0079\u0067\u006f\u006e"))
	_bdg.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _cbeb.Vertices)
	_bdg.SetIfNotNil("\u004c\u0045", _cbeb.LE)
	_bdg.SetIfNotNil("\u0042\u0053", _cbeb.BS)
	_bdg.SetIfNotNil("\u0049\u0043", _cbeb.IC)
	_bdg.SetIfNotNil("\u0042\u0045", _cbeb.BE)
	_bdg.SetIfNotNil("\u0049\u0054", _cbeb.IT)
	_bdg.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _cbeb.Measure)
	return _cbfg
}

// PdfColorspaceLab is a L*, a*, b* 3 component colorspace.
type PdfColorspaceLab struct {
	WhitePoint []float64
	BlackPoint []float64
	Range      []float64
	_gbbd      *_cde.PdfIndirectObject
}
