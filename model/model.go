package model

import (
	_ba "bufio"
	_ca "bytes"
	_dd "crypto/md5"
	_gd "crypto/rand"
	_bd "crypto/sha1"
	_g "crypto/x509"
	_cb "encoding/binary"
	_d "encoding/hex"
	_gf "errors"
	_bg "fmt"
	_c "hash"
	_gdc "image"
	_e "image/color"
	_ "image/gif"
	_ "image/png"
	_ab "io"
	_ef "io/ioutil"
	_cbg "math"
	_ff "math/rand"
	_ed "os"
	_a "regexp"
	_ae "sort"
	_aa "strconv"
	_ee "strings"
	_bf "sync"
	_f "time"
	_cc "unicode"
	_de "unicode/utf8"

	_eg "bitbucket.org/shenghui0779/gopdf/common"
	_ebb "bitbucket.org/shenghui0779/gopdf/core"
	_fe "bitbucket.org/shenghui0779/gopdf/core/security"
	_fa "bitbucket.org/shenghui0779/gopdf/core/security/crypt"
	_ebe "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_dg "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_abg "bitbucket.org/shenghui0779/gopdf/internal/sampling"
	_da "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_fd "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_fef "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_bda "bitbucket.org/shenghui0779/gopdf/model/internal/docutil"
	_bad "bitbucket.org/shenghui0779/gopdf/model/internal/fonts"
	_ac "bitbucket.org/shenghui0779/gopdf/model/mdp"
	_cg "bitbucket.org/shenghui0779/gopdf/model/sigutil"
	_bc "bitbucket.org/shenghui0779/gopdf/ps"
	_gb "github.com/unidoc/pkcs7"
	_gbg "github.com/unidoc/unitype"
	_bfc "golang.org/x/xerrors"
)

// DefaultImageHandler is the default implementation of the ImageHandler using the standard go library.
type DefaultImageHandler struct{}

func (_fdfbe *PdfWriter) writeString(_ggfbf string) {
	if _fdfbe._bgef != nil {
		return
	}
	_fbdbg, _cbaba := _fdfbe._cbabb.WriteString(_ggfbf)
	_fdfbe._afedd += int64(_fbdbg)
	_fdfbe._bgef = _cbaba
}

// NewPdfAnnotationCircle returns a new circle annotation.
func NewPdfAnnotationCircle() *PdfAnnotationCircle {
	_cfg := NewPdfAnnotation()
	_geb := &PdfAnnotationCircle{}
	_geb.PdfAnnotation = _cfg
	_geb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cfg.SetContext(_geb)
	return _geb
}
func (_febag *Image) samplesTrimPadding(_efgc []uint32) []uint32 {
	_fdfge := _febag.ColorComponents * int(_febag.Width) * int(_febag.Height)
	if len(_efgc) == _fdfge {
		return _efgc
	}
	_bdbbd := make([]uint32, _fdfge)
	_daeae := int(_febag.Width) * _febag.ColorComponents
	var _aebae, _bege, _agabg, _ggcb int
	_bcefb := _dg.BytesPerLine(int(_febag.Width), int(_febag.BitsPerComponent), _febag.ColorComponents)
	for _aebae = 0; _aebae < int(_febag.Height); _aebae++ {
		_bege = _aebae * int(_febag.Width)
		_agabg = _aebae * _bcefb
		for _ggcb = 0; _ggcb < _daeae; _ggcb++ {
			_bdbbd[_bege+_ggcb] = _efgc[_agabg+_ggcb]
		}
	}
	return _bdbbd
}
func (_agaf *pdfCIDFontType0) getFontDescriptor() *PdfFontDescriptor { return _agaf._fbbd }

// NewPdfAnnotationMovie returns a new movie annotation.
func NewPdfAnnotationMovie() *PdfAnnotationMovie {
	_cdcd := NewPdfAnnotation()
	_ffee := &PdfAnnotationMovie{}
	_ffee.PdfAnnotation = _cdcd
	_cdcd.SetContext(_ffee)
	return _ffee
}
func (_bbgg *PdfReader) newPdfAnnotationProjectionFromDict(_ebbe *_ebb.PdfObjectDictionary) (*PdfAnnotationProjection, error) {
	_cbgb := &PdfAnnotationProjection{}
	_gcfe, _aafc := _bbgg.newPdfAnnotationMarkupFromDict(_ebbe)
	if _aafc != nil {
		return nil, _aafc
	}
	_cbgb.PdfAnnotationMarkup = _gcfe
	return _cbgb, nil
}

// ToPdfObject implements interface PdfModel.
func (_bae *PdfActionGoToE) ToPdfObject() _ebb.PdfObject {
	_bae.PdfAction.ToPdfObject()
	_gg := _bae._abe
	_ea := _gg.PdfObject.(*_ebb.PdfObjectDictionary)
	_ea.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeGoToE)))
	if _bae.F != nil {
		_ea.Set("\u0046", _bae.F.ToPdfObject())
	}
	_ea.SetIfNotNil("\u0044", _bae.D)
	_ea.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _bae.NewWindow)
	_ea.SetIfNotNil("\u0054", _bae.T)
	return _gg
}
func _abba(_ddfbd _ebb.PdfObject) (*PdfColorspaceSpecialPattern, error) {
	_eg.Log.Trace("\u004e\u0065\u0077\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u0043\u0053\u0020\u0066\u0072\u006fm\u0020\u006f\u0062\u006a\u003a\u0020\u0025\u0073\u0020\u0025\u0054", _ddfbd.String(), _ddfbd)
	_cdag := NewPdfColorspaceSpecialPattern()
	if _gggfc, _eceb := _ddfbd.(*_ebb.PdfIndirectObject); _eceb {
		_cdag._adegb = _gggfc
	}
	_ddfbd = _ebb.TraceToDirectObject(_ddfbd)
	if _bfgf, _gfaf := _ddfbd.(*_ebb.PdfObjectName); _gfaf {
		if *_bfgf != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			return nil, _bg.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
		return _cdag, nil
	}
	_dced, _cddc := _ddfbd.(*_ebb.PdfObjectArray)
	if !_cddc {
		_eg.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061t\u0074\u0065\u0072\u006e\u0020\u0043\u0053 \u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0023\u0076", _ddfbd)
		return nil, _bg.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0050\u0061\u0074\u0074e\u0072n\u0020C\u0053\u0020\u006f\u0062\u006a\u0065\u0063t")
	}
	if _dced.Len() != 1 && _dced.Len() != 2 {
		_eg.Log.Error("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079\u003a %\u0023\u0076", _dced)
		return nil, _bg.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065r\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_ddfbd = _dced.Get(0)
	if _egfcf, _cfdd := _ddfbd.(*_ebb.PdfObjectName); _cfdd {
		if *_egfcf != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			_eg.Log.Error("\u0049\u006e\u0076al\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020C\u0053 \u0061r\u0072\u0061\u0079\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0023\u0076", _egfcf)
			return nil, _bg.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
	}
	if _dced.Len() > 1 {
		_ddfbd = _dced.Get(1)
		_ddfbd = _ebb.TraceToDirectObject(_ddfbd)
		_bece, _cdec := NewPdfColorspaceFromPdfObject(_ddfbd)
		if _cdec != nil {
			return nil, _cdec
		}
		_cdag.UnderlyingCS = _bece
	}
	_eg.Log.Trace("R\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0077i\u0074\u0068\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079in\u0067\u0020\u0063s\u003a \u0025\u0054", _cdag.UnderlyingCS)
	return _cdag, nil
}
func _eacbd() _f.Time { _daddc.Lock(); defer _daddc.Unlock(); return _ccdff }

// NewPdfFilespecFromObj creates and returns a new PdfFilespec object.
func NewPdfFilespecFromObj(obj _ebb.PdfObject) (*PdfFilespec, error) {
	_edge := &PdfFilespec{}
	var _fada *_ebb.PdfObjectDictionary
	if _bfdgg, _gfdd := _ebb.GetIndirect(obj); _gfdd {
		_edge._gcge = _bfdgg
		_cdcc, _ffff := _ebb.GetDict(_bfdgg.PdfObject)
		if !_ffff {
			_eg.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _ebb.ErrTypeError
		}
		_fada = _cdcc
	} else if _ggefg, _dbdcd := _ebb.GetDict(obj); _dbdcd {
		_edge._gcge = _ggefg
		_fada = _ggefg
	} else {
		_eg.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", obj)
		return nil, _ebb.ErrTypeError
	}
	if _fada == nil {
		_eg.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _gf.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _cfcf := _fada.Get("\u0054\u0079\u0070\u0065"); _cfcf != nil {
		_fgadf, _dcccf := _cfcf.(*_ebb.PdfObjectName)
		if !_dcccf {
			_eg.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _cfcf)
		} else {
			if *_fgadf != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
				_eg.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070e\u0063\u0074\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063 \u0028\u0025\u0073\u0029", *_fgadf)
			}
		}
	}
	if _eedda := _fada.Get("\u0046\u0053"); _eedda != nil {
		_edge.FS = _eedda
	}
	if _gebaa := _fada.Get("\u0046"); _gebaa != nil {
		_edge.F = _gebaa
	}
	if _dbga := _fada.Get("\u0055\u0046"); _dbga != nil {
		_edge.UF = _dbga
	}
	if _bgdag := _fada.Get("\u0044\u004f\u0053"); _bgdag != nil {
		_edge.DOS = _bgdag
	}
	if _agcea := _fada.Get("\u004d\u0061\u0063"); _agcea != nil {
		_edge.Mac = _agcea
	}
	if _deed := _fada.Get("\u0055\u006e\u0069\u0078"); _deed != nil {
		_edge.Unix = _deed
	}
	if _bbba := _fada.Get("\u0049\u0044"); _bbba != nil {
		_edge.ID = _bbba
	}
	if _gdcf := _fada.Get("\u0056"); _gdcf != nil {
		_edge.V = _gdcf
	}
	if _egfe := _fada.Get("\u0045\u0046"); _egfe != nil {
		_edge.EF = _egfe
	}
	if _eadc := _fada.Get("\u0052\u0046"); _eadc != nil {
		_edge.RF = _eadc
	}
	if _gcgbb := _fada.Get("\u0044\u0065\u0073\u0063"); _gcgbb != nil {
		_edge.Desc = _gcgbb
	}
	if _bbae := _fada.Get("\u0043\u0049"); _bbae != nil {
		_edge.CI = _bbae
	}
	return _edge, nil
}

// NewPdfAnnotationSquiggly returns a new text squiggly annotation.
func NewPdfAnnotationSquiggly() *PdfAnnotationSquiggly {
	_bef := NewPdfAnnotation()
	_bcg := &PdfAnnotationSquiggly{}
	_bcg.PdfAnnotation = _bef
	_bcg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_bef.SetContext(_bcg)
	return _bcg
}

// ToPdfObject implements interface PdfModel.
func (_fgd *PdfActionGoToR) ToPdfObject() _ebb.PdfObject {
	_fgd.PdfAction.ToPdfObject()
	_gbeb := _fgd._abe
	_ebdd := _gbeb.PdfObject.(*_ebb.PdfObjectDictionary)
	_ebdd.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeGoToR)))
	if _fgd.F != nil {
		_ebdd.Set("\u0046", _fgd.F.ToPdfObject())
	}
	_ebdd.SetIfNotNil("\u0044", _fgd.D)
	_ebdd.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _fgd.NewWindow)
	return _gbeb
}

// FieldImageProvider provides fields images for specified fields.
type FieldImageProvider interface {
	FieldImageValues() (map[string]*Image, error)
}

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_dcfeg *PdfWriter) SetOptimizer(optimizer Optimizer) { _dcfeg._aadcd = optimizer }

// PdfActionThread represents a thread action.
type PdfActionThread struct {
	*PdfAction
	F *PdfFilespec
	D _ebb.PdfObject
	B _ebb.PdfObject
}

func _feaae() string {
	_daddc.Lock()
	defer _daddc.Unlock()
	if len(_afedf) > 0 {
		return _afedf
	}
	return "Go PDF"
}
func (_dadec *PdfReader) newPdfOutlineItemFromIndirectObject(_ccgee *_ebb.PdfIndirectObject) (*PdfOutlineItem, error) {
	_agffa, _ebed := _ccgee.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_ebed {
		return nil, _bg.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_gcdc := NewPdfOutlineItem()
	_fbdc := _agffa.Get("\u0054\u0069\u0074l\u0065")
	if _fbdc == nil {
		return nil, _bg.Errorf("\u006d\u0069\u0073s\u0069\u006e\u0067\u0020\u0054\u0069\u0074\u006c\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0049\u0074\u0065\u006d\u0020\u0028r\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	_ffebg, _ddbb := _ebb.GetString(_fbdc)
	if !_ddbb {
		return nil, _bg.Errorf("\u0074\u0069\u0074le\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0054\u0029", _fbdc)
	}
	_gcdc.Title = _ffebg
	if _cfeec := _agffa.Get("\u0043\u006f\u0075n\u0074"); _cfeec != nil {
		_cbcfb, _febb := _cfeec.(*_ebb.PdfObjectInteger)
		if !_febb {
			return nil, _bg.Errorf("\u0063o\u0075\u006e\u0074\u0020n\u006f\u0074\u0020\u0061\u006e \u0069n\u0074e\u0067\u0065\u0072\u0020\u0028\u0025\u0054)", _cfeec)
		}
		_ccacf := int64(*_cbcfb)
		_gcdc.Count = &_ccacf
	}
	if _dcebd := _agffa.Get("\u0044\u0065\u0073\u0074"); _dcebd != nil {
		_gcdc.Dest = _ebb.ResolveReference(_dcebd)
		if !_dadec._ceefa {
			_bcegf := _dadec.traverseObjectData(_gcdc.Dest)
			if _bcegf != nil {
				return nil, _bcegf
			}
		}
	}
	if _dfegb := _agffa.Get("\u0041"); _dfegb != nil {
		_gcdc.A = _ebb.ResolveReference(_dfegb)
		if !_dadec._ceefa {
			_gdgeg := _dadec.traverseObjectData(_gcdc.A)
			if _gdgeg != nil {
				return nil, _gdgeg
			}
		}
	}
	if _ddbcc := _agffa.Get("\u0053\u0045"); _ddbcc != nil {
		_gcdc.SE = nil
	}
	if _dafga := _agffa.Get("\u0043"); _dafga != nil {
		_gcdc.C = _ebb.ResolveReference(_dafga)
	}
	if _bggc := _agffa.Get("\u0046"); _bggc != nil {
		_gcdc.F = _ebb.ResolveReference(_bggc)
	}
	return _gcdc, nil
}

// PdfFunctionType3 defines stitching of the subdomains of several 1-input functions to produce
// a single new 1-input function.
type PdfFunctionType3 struct {
	Domain    []float64
	Range     []float64
	Functions []PdfFunction
	Bounds    []float64
	Encode    []float64
	_acac     *_ebb.PdfIndirectObject
}

func (_eacaf *LTV) generateVRIKey(_eceff *PdfSignature) (string, error) {
	_ffaed, _fefgg := _eaef(_eceff.Contents.Bytes())
	if _fefgg != nil {
		return "", _fefgg
	}
	return _ee.ToUpper(_d.EncodeToString(_ffaed)), nil
}

// PdfField contains the common attributes of a form field. The context object contains the specific field data
// which can represent a button, text, choice or signature.
// The PdfField is typically not used directly, but is encapsulated by the more specific field types such as
// PdfFieldButton etc (i.e. the context attribute).
type PdfField struct {
	_cada        PdfModel
	_cdfd        *_ebb.PdfIndirectObject
	Parent       *PdfField
	Annotations  []*PdfAnnotationWidget
	Kids         []*PdfField
	FT           *_ebb.PdfObjectName
	T            *_ebb.PdfObjectString
	TU           *_ebb.PdfObjectString
	TM           *_ebb.PdfObjectString
	Ff           *_ebb.PdfObjectInteger
	V            _ebb.PdfObject
	DV           _ebb.PdfObject
	AA           _ebb.PdfObject
	VariableText *VariableText
}

func (_baae *PdfReader) newPdfActionJavaScriptFromDict(_edefc *_ebb.PdfObjectDictionary) (*PdfActionJavaScript, error) {
	return &PdfActionJavaScript{JS: _edefc.Get("\u004a\u0053")}, nil
}
func _afdc(_ggafa *fontCommon) *pdfCIDFontType0 { return &pdfCIDFontType0{fontCommon: *_ggafa} }

// IsEncrypted returns true if the PDF file is encrypted.
func (_badbc *PdfReader) IsEncrypted() (bool, error) { return _badbc._cafdf.IsEncrypted() }

// GetContentStreams returns the content stream as an array of strings.
func (_eaac *PdfPage) GetContentStreams() ([]string, error) {
	_cadgf := _eaac.GetContentStreamObjs()
	var _aacf []string
	for _, _debg := range _cadgf {
		_eeafa, _ggcbg := _bfebg(_debg)
		if _ggcbg != nil {
			return nil, _ggcbg
		}
		_aacf = append(_aacf, _eeafa)
	}
	return _aacf, nil
}

// GetContainingPdfObject implements interface PdfModel.
func (_fea *PdfAction) GetContainingPdfObject() _ebb.PdfObject { return _fea._abe }
func (_dcga *PdfWriter) writeXRefStreams(_bfabe int, _bcfe int64) error {
	_bdcea := _bfabe + 1
	_dcga._bedfc[_bdcea] = crossReference{Type: 1, ObjectNumber: _bdcea, Offset: _bcfe}
	_gccdc := _ca.NewBuffer(nil)
	_daege := _ebb.MakeArray()
	for _aggfc := 0; _aggfc <= _bfabe; {
		for ; _aggfc <= _bfabe; _aggfc++ {
			_efcea, _egda := _dcga._bedfc[_aggfc]
			if _egda && (!_dcga._abffb || _dcga._abffb && (_efcea.Type == 1 && _efcea.Offset >= _dcga._ggbfg || _efcea.Type == 0)) {
				break
			}
		}
		var _aggbde int
		for _aggbde = _aggfc + 1; _aggbde <= _bfabe; _aggbde++ {
			_gaaed, _eccag := _dcga._bedfc[_aggbde]
			if _eccag && (!_dcga._abffb || _dcga._abffb && (_gaaed.Type == 1 && _gaaed.Offset > _dcga._ggbfg)) {
				continue
			}
			break
		}
		_daege.Append(_ebb.MakeInteger(int64(_aggfc)), _ebb.MakeInteger(int64(_aggbde-_aggfc)))
		for _geeaf := _aggfc; _geeaf < _aggbde; _geeaf++ {
			_cbcea := _dcga._bedfc[_geeaf]
			switch _cbcea.Type {
			case 0:
				_cb.Write(_gccdc, _cb.BigEndian, byte(0))
				_cb.Write(_gccdc, _cb.BigEndian, uint32(0))
				_cb.Write(_gccdc, _cb.BigEndian, uint16(0xFFFF))
			case 1:
				_cb.Write(_gccdc, _cb.BigEndian, byte(1))
				_cb.Write(_gccdc, _cb.BigEndian, uint32(_cbcea.Offset))
				_cb.Write(_gccdc, _cb.BigEndian, uint16(_cbcea.Generation))
			case 2:
				_cb.Write(_gccdc, _cb.BigEndian, byte(2))
				_cb.Write(_gccdc, _cb.BigEndian, uint32(_cbcea.ObjectNumber))
				_cb.Write(_gccdc, _cb.BigEndian, uint16(_cbcea.Index))
			}
		}
		_aggfc = _aggbde + 1
	}
	_bdgee, _abgec := _ebb.MakeStream(_gccdc.Bytes(), _ebb.NewFlateEncoder())
	if _abgec != nil {
		return _abgec
	}
	_bdgee.ObjectNumber = int64(_bdcea)
	_bdgee.PdfObjectDictionary.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0058\u0052\u0065\u0066"))
	_bdgee.PdfObjectDictionary.Set("\u0057", _ebb.MakeArray(_ebb.MakeInteger(1), _ebb.MakeInteger(4), _ebb.MakeInteger(2)))
	_bdgee.PdfObjectDictionary.Set("\u0049\u006e\u0064e\u0078", _daege)
	_bdgee.PdfObjectDictionary.Set("\u0053\u0069\u007a\u0065", _ebb.MakeInteger(int64(_bdcea+1)))
	_bdgee.PdfObjectDictionary.Set("\u0049\u006e\u0066\u006f", _dcga._eadfd)
	_bdgee.PdfObjectDictionary.Set("\u0052\u006f\u006f\u0074", _dcga._gegba)
	if _dcga._abffb && _dcga._bcage > 0 {
		_bdgee.PdfObjectDictionary.Set("\u0050\u0072\u0065\u0076", _ebb.MakeInteger(_dcga._bcage))
	}
	if _dcga._cgfde != nil {
		_bdgee.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _dcga._cbcaa)
	}
	if _dcga._eecfe == nil && _dcga._gfdea != "" && _dcga._gffb != "" {
		_dcga._eecfe = _ebb.MakeArray(_ebb.MakeHexString(_dcga._gfdea), _ebb.MakeHexString(_dcga._gffb))
	}
	if _dcga._eecfe != nil {
		_eg.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _dcga._eecfe)
		_bdgee.Set("\u0049\u0044", _dcga._eecfe)
	}
	_dcga.writeObject(int(_bdgee.ObjectNumber), _bdgee)
	return nil
}

// NewPdfActionGoTo3DView returns a new "goTo3DView" action.
func NewPdfActionGoTo3DView() *PdfActionGoTo3DView {
	_fbb := NewPdfAction()
	_ebd := &PdfActionGoTo3DView{}
	_ebd.PdfAction = _fbb
	_fbb.SetContext(_ebd)
	return _ebd
}

// SetDocInfo sets the document /Info metadata.
// This will overwrite any globally declared document info.
func (_ebdf *PdfAppender) SetDocInfo(info *PdfInfo) { _ebdf._eeee = info }

// SetDate sets the `M` field of the signature.
func (_fbagd *PdfSignature) SetDate(date _f.Time, format string) {
	if format == "" {
		format = "\u0044\u003a\u003200\u0036\u0030\u0031\u0030\u0032\u0031\u0035\u0030\u0034\u0030\u0035\u002d\u0030\u0037\u0027\u0030\u0030\u0027"
	}
	_fbagd.M = _ebb.MakeString(date.Format(format))
}

// DetermineColorspaceNameFromPdfObject determines PDF colorspace from a PdfObject.  Returns the colorspace name and
// an error on failure. If the colorspace was not found, will return an empty string.
func DetermineColorspaceNameFromPdfObject(obj _ebb.PdfObject) (_ebb.PdfObjectName, error) {
	var _cdbg *_ebb.PdfObjectName
	var _ggeb *_ebb.PdfObjectArray
	if _cfccf, _bcbbf := obj.(*_ebb.PdfIndirectObject); _bcbbf {
		if _ddfc, _dabe := _cfccf.PdfObject.(*_ebb.PdfObjectArray); _dabe {
			_ggeb = _ddfc
		} else if _bfbgf, _faba := _cfccf.PdfObject.(*_ebb.PdfObjectName); _faba {
			_cdbg = _bfbgf
		}
	} else if _gdcdg, _eeca := obj.(*_ebb.PdfObjectArray); _eeca {
		_ggeb = _gdcdg
	} else if _fabe, _fdde := obj.(*_ebb.PdfObjectName); _fdde {
		_cdbg = _fabe
	}
	if _cdbg != nil {
		switch *_cdbg {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_cdbg, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_cdbg, nil
		}
	}
	if _ggeb != nil && _ggeb.Len() > 0 {
		if _gbacc, _edcec := _ggeb.Get(0).(*_ebb.PdfObjectName); _edcec {
			switch *_gbacc {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _ggeb.Len() == 1 {
					return *_gbacc, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_gbacc, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_gbacc, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_gbacc, nil
			}
		}
	}
	return "", nil
}

// ToOutlineTree returns a low level PdfOutlineTreeNode object, based on
// the current instance.
func (_cgaea *Outline) ToOutlineTree() *PdfOutlineTreeNode {
	return &_cgaea.ToPdfOutline().PdfOutlineTreeNode
}
func _afag(_aebfb *_ebb.PdfIndirectObject, _ggdddb *_ebb.PdfObjectDictionary) (*DSS, error) {
	if _aebfb == nil {
		_aebfb = _ebb.MakeIndirectObject(nil)
	}
	_aebfb.PdfObject = _ebb.MakeDict()
	_bbdfe := map[string]*VRI{}
	if _ebfgg, _dace := _ebb.GetDict(_ggdddb.Get("\u0056\u0052\u0049")); _dace {
		for _, _dacg := range _ebfgg.Keys() {
			if _adege, _ageec := _ebb.GetDict(_ebfgg.Get(_dacg)); _ageec {
				_bbdfe[_ee.ToUpper(_dacg.String())] = _bcbe(_adege)
			}
		}
	}
	return &DSS{Certs: _ffaa(_ggdddb.Get("\u0043\u0065\u0072t\u0073")), OCSPs: _ffaa(_ggdddb.Get("\u004f\u0043\u0053P\u0073")), CRLs: _ffaa(_ggdddb.Get("\u0043\u0052\u004c\u0073")), VRI: _bbdfe, _fcgb: _aebfb}, nil
}

// Set applies flag fl to the flag's bitmask and returns the combined flag.
func (_gbcf FieldFlag) Set(fl FieldFlag) FieldFlag { return FieldFlag(_gbcf.Mask() | fl.Mask()) }

// NewPdfColorspaceDeviceGray returns a new grayscale colorspace.
func NewPdfColorspaceDeviceGray() *PdfColorspaceDeviceGray { return &PdfColorspaceDeviceGray{} }
func (_dbfee *PdfWriter) optimize() error {
	if _dbfee._aadcd == nil {
		return nil
	}
	var _fgdf error
	_dbfee._ebdgg, _fgdf = _dbfee._aadcd.Optimize(_dbfee._ebdgg)
	if _fgdf != nil {
		return _fgdf
	}
	_accgc := make(map[_ebb.PdfObject]struct{}, len(_dbfee._ebdgg))
	for _, _gaeff := range _dbfee._ebdgg {
		_accgc[_gaeff] = struct{}{}
	}
	_dbfee._ffffd = _accgc
	return nil
}

// PdfActionSetOCGState represents a SetOCGState action.
type PdfActionSetOCGState struct {
	*PdfAction
	State      _ebb.PdfObject
	PreserveRB _ebb.PdfObject
}

// NewPdfAnnotationSound returns a new sound annotation.
func NewPdfAnnotationSound() *PdfAnnotationSound {
	_eac := NewPdfAnnotation()
	_feae := &PdfAnnotationSound{}
	_feae.PdfAnnotation = _eac
	_feae.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_eac.SetContext(_feae)
	return _feae
}

// PdfActionMovie represents a movie action.
type PdfActionMovie struct {
	*PdfAction
	Annotation _ebb.PdfObject
	T          _ebb.PdfObject
	Operation  _ebb.PdfObject
}

// GetShadingByName gets the shading specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_dcgb *PdfPageResources) GetShadingByName(keyName _ebb.PdfObjectName) (*PdfShading, bool) {
	if _dcgb.Shading == nil {
		return nil, false
	}
	_egfcg, _cbbb := _ebb.TraceToDirectObject(_dcgb.Shading).(*_ebb.PdfObjectDictionary)
	if !_cbbb {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0068\u0061d\u0069\u006e\u0067\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _dcgb.Shading)
		return nil, false
	}
	if _gccd := _egfcg.Get(keyName); _gccd != nil {
		_fgcgf, _fdceb := _ggdfc(_gccd)
		if _fdceb != nil {
			_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0073\u0068\u0061\u0064\u0069\u006eg\u003a\u0020\u0025\u0076", _fdceb)
			return nil, false
		}
		return _fgcgf, true
	}
	return nil, false
}

// NewPdfColorspaceICCBased returns a new ICCBased colorspace object.
func NewPdfColorspaceICCBased(N int) (*PdfColorspaceICCBased, error) {
	_bebbf := &PdfColorspaceICCBased{}
	if N != 1 && N != 3 && N != 4 {
		return nil, _bg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u0031/\u0033\u002f\u0034\u0029")
	}
	_bebbf.N = N
	return _bebbf, nil
}

// GetContainingPdfObject returns the XObject Form's containing object (indirect object).
func (_cgadf *XObjectForm) GetContainingPdfObject() _ebb.PdfObject { return _cgadf._gebcd }
func _ccef(_agdbd _ebb.PdfObject) (*PdfPattern, error) {
	_gafde := &PdfPattern{}
	var _ceaa *_ebb.PdfObjectDictionary
	if _eaaec, _decb := _ebb.GetIndirect(_agdbd); _decb {
		_gafde._dcddc = _eaaec
		_afgge, _fbag := _eaaec.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_fbag {
			_eg.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0028g\u006f\u0074\u0020%\u0054\u0029", _eaaec.PdfObject)
			return nil, _ebb.ErrTypeError
		}
		_ceaa = _afgge
	} else if _cegcb, _gfcae := _ebb.GetStream(_agdbd); _gfcae {
		_gafde._dcddc = _cegcb
		_ceaa = _cegcb.PdfObjectDictionary
	} else {
		_eg.Log.Debug("\u0050a\u0074\u0074e\u0072\u006e\u0020\u006eo\u0074\u0020\u0061n\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 o\u0062\u006a\u0065c\u0074\u0020o\u0072\u0020\u0073\u0074\u0072\u0065a\u006d\u002e \u0025\u0054", _agdbd)
		return nil, _ebb.ErrTypeError
	}
	_adcaa := _ceaa.Get("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065")
	if _adcaa == nil {
		_eg.Log.Debug("\u0050\u0064\u0066\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069n\u0067\u0020\u0050\u0061\u0074t\u0065\u0072n\u0054\u0079\u0070\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_eefed, _ecaac := _adcaa.(*_ebb.PdfObjectInteger)
	if !_ecaac {
		_eg.Log.Debug("\u0050\u0061tt\u0065\u0072\u006e \u0074\u0079\u0070\u0065 no\u0074 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067er\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _adcaa)
		return nil, _ebb.ErrTypeError
	}
	if *_eefed != 1 && *_eefed != 2 {
		_eg.Log.Debug("\u0050\u0061\u0074\u0074e\u0072\u006e\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d \u0031/\u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", *_eefed)
		return nil, _ebb.ErrRangeError
	}
	_gafde.PatternType = int64(*_eefed)
	switch *_eefed {
	case 1:
		_gfcgf, _gdeec := _ggbd(_ceaa)
		if _gdeec != nil {
			return nil, _gdeec
		}
		_gfcgf.PdfPattern = _gafde
		_gafde._ffagg = _gfcgf
		return _gafde, nil
	case 2:
		_abeg, _bdcf := _cdbaf(_ceaa)
		if _bdcf != nil {
			return nil, _bdcf
		}
		_abeg.PdfPattern = _gafde
		_gafde._ffagg = _abeg
		return _gafde, nil
	}
	return nil, _gf.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a Lab device.
func (_bebe *PdfColorspaceLab) GetNumComponents() int { return 3 }

// GetDocMDPPermission returns the DocMDP level of the restrictions
func (_cdbcc *PdfSignature) GetDocMDPPermission() (_ac.DocMDPPermission, bool) {
	for _, _egaba := range _cdbcc.Reference.Elements() {
		if _baebf, _gdcag := _ebb.GetDict(_egaba); _gdcag {
			if _edcdd, _dcgc := _ebb.GetNameVal(_baebf.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _dcgc && _edcdd == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				if _bfegg, _dfbgbd := _ebb.GetDict(_baebf.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _dfbgbd {
					if P, _dbdgf := _ebb.GetIntVal(_bfegg.Get("\u0050")); _dbdgf {
						return _ac.DocMDPPermission(P), true
					}
				}
			}
		}
	}
	return 0, false
}
func _ccgda(_adaea *_ebb.PdfObjectDictionary) (*PdfShadingType4, error) {
	_eggcb := PdfShadingType4{}
	_bffcd := _adaea.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _bffcd == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_aegff, _ceed := _bffcd.(*_ebb.PdfObjectInteger)
	if !_ceed {
		_eg.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bffcd)
		return nil, _ebb.ErrTypeError
	}
	_eggcb.BitsPerCoordinate = _aegff
	_bffcd = _adaea.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _bffcd == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_aegff, _ceed = _bffcd.(*_ebb.PdfObjectInteger)
	if !_ceed {
		_eg.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _bffcd)
		return nil, _ebb.ErrTypeError
	}
	_eggcb.BitsPerComponent = _aegff
	_bffcd = _adaea.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _bffcd == nil {
		_eg.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_aegff, _ceed = _bffcd.(*_ebb.PdfObjectInteger)
	if !_ceed {
		_eg.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _bffcd)
		return nil, _ebb.ErrTypeError
	}
	_eggcb.BitsPerComponent = _aegff
	_bffcd = _adaea.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _bffcd == nil {
		_eg.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_ddefb, _ceed := _bffcd.(*_ebb.PdfObjectArray)
	if !_ceed {
		_eg.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _bffcd)
		return nil, _ebb.ErrTypeError
	}
	_eggcb.Decode = _ddefb
	_bffcd = _adaea.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _bffcd == nil {
		_eg.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_eggcb.Function = []PdfFunction{}
	if _aeaee, _ebbga := _bffcd.(*_ebb.PdfObjectArray); _ebbga {
		for _, _gceg := range _aeaee.Elements() {
			_adef, _cdbgb := _aagg(_gceg)
			if _cdbgb != nil {
				_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cdbgb)
				return nil, _cdbgb
			}
			_eggcb.Function = append(_eggcb.Function, _adef)
		}
	} else {
		_gdda, _bfeed := _aagg(_bffcd)
		if _bfeed != nil {
			_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bfeed)
			return nil, _bfeed
		}
		_eggcb.Function = append(_eggcb.Function, _gdda)
	}
	return &_eggcb, nil
}

// VRI represents a Validation-Related Information dictionary.
// The VRI dictionary contains validation data in the form of
// certificates, OCSP and CRL information, for a single signature.
// See ETSI TS 102 778-4 V1.1.1 for more information.
type VRI struct {
	Cert []*_ebb.PdfObjectStream
	OCSP []*_ebb.PdfObjectStream
	CRL  []*_ebb.PdfObjectStream
	TU   *_ebb.PdfObjectString
	TS   *_ebb.PdfObjectString
}

var _cbcfg = map[string]string{"\u0053\u0079\u006d\u0062\u006f\u006c": "\u0053\u0079\u006d\u0062\u006f\u006c\u0045\u006e\u0063o\u0064\u0069\u006e\u0067", "\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073": "Z\u0061p\u0066\u0044\u0069\u006e\u0067\u0062\u0061\u0074s\u0045\u006e\u0063\u006fdi\u006e\u0067"}

func _gfed(_cfdc, _ceccd string) string {
	if _ee.Contains(_cfdc, "\u002b") {
		_ccfag := _ee.Split(_cfdc, "\u002b")
		if len(_ccfag) == 2 {
			_cfdc = _ccfag[1]
		}
	}
	return _ceccd + "\u002b" + _cfdc
}

// NewPdfActionResetForm returns a new "reset form" action.
func NewPdfActionResetForm() *PdfActionResetForm {
	_dac := NewPdfAction()
	_abb := &PdfActionResetForm{}
	_abb.PdfAction = _dac
	_dac.SetContext(_abb)
	return _abb
}
func (_egeg *DSS) addOCSPs(_bdbaa [][]byte) ([]*_ebb.PdfObjectStream, error) {
	return _egeg.add(&_egeg.OCSPs, _egeg._cadd, _bdbaa)
}

// SetXObjectFormByName adds the provided XObjectForm to the page resources.
// The added XObjectForm is identified by the specified name.
func (_dffa *PdfPageResources) SetXObjectFormByName(keyName _ebb.PdfObjectName, xform *XObjectForm) error {
	_fegb := xform.ToPdfObject().(*_ebb.PdfObjectStream)
	_fbdb := _dffa.SetXObjectByName(keyName, _fegb)
	return _fbdb
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
	FieldFlagRichText          FieldFlag = (1 << 25)
	FieldFlagDoNotSpellCheck   FieldFlag = (1 << 22)
	FieldFlagCombo             FieldFlag = (1 << 17)
	FieldFlagEdit              FieldFlag = (1 << 18)
	FieldFlagSort              FieldFlag = (1 << 19)
	FieldFlagMultiSelect       FieldFlag = (1 << 21)
	FieldFlagCommitOnSelChange FieldFlag = (1 << 26)
)

// ToPdfObject implements interface PdfModel.
func (_agee *PdfAnnotationRichMedia) ToPdfObject() _ebb.PdfObject {
	_agee.PdfAnnotation.ToPdfObject()
	_dbadg := _agee._bdcd
	_cag := _dbadg.PdfObject.(*_ebb.PdfObjectDictionary)
	_cag.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a"))
	_cag.SetIfNotNil("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073", _agee.RichMediaSettings)
	_cag.SetIfNotNil("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074", _agee.RichMediaContent)
	return _dbadg
}
func (_dafe *DSS) generateHashMap(_fbed []*_ebb.PdfObjectStream) (map[string]*_ebb.PdfObjectStream, error) {
	_cbdd := map[string]*_ebb.PdfObjectStream{}
	for _, _bdab := range _fbed {
		_acaf, _gfga := _ebb.DecodeStream(_bdab)
		if _gfga != nil {
			return nil, _gfga
		}
		_geddd, _gfga := _eaef(_acaf)
		if _gfga != nil {
			return nil, _gfga
		}
		_cbdd[string(_geddd)] = _bdab
	}
	return _cbdd, nil
}
func (_faed *PdfReader) newPdfAnnotation3DFromDict(_efcc *_ebb.PdfObjectDictionary) (*PdfAnnotation3D, error) {
	_fdfc := PdfAnnotation3D{}
	_fdfc.T3DD = _efcc.Get("\u0033\u0044\u0044")
	_fdfc.T3DV = _efcc.Get("\u0033\u0044\u0056")
	_fdfc.T3DA = _efcc.Get("\u0033\u0044\u0041")
	_fdfc.T3DI = _efcc.Get("\u0033\u0044\u0049")
	_fdfc.T3DB = _efcc.Get("\u0033\u0044\u0042")
	return &_fdfc, nil
}

// PdfAppender appends new PDF content to an existing PDF document via incremental updates.
type PdfAppender struct {
	_ecce  _ab.ReadSeeker
	_gege  *_ebb.PdfParser
	_acfe  *PdfReader
	Reader *PdfReader
	_dfbg  []*PdfPage
	_bfef  *PdfAcroForm
	_eged  *DSS
	_eaaa  *Permissions
	_acfd  _ebb.XrefTable
	_cfag  int64
	_gbddb int
	_bfeg  []_ebb.PdfObject
	_ddfg  map[_ebb.PdfObject]struct{}
	_gbfa  map[_ebb.PdfObject]int64
	_eebc  map[_ebb.PdfObject]struct{}
	_agb   map[_ebb.PdfObject]struct{}
	_bee   int64
	_adca  bool
	_accg  string
	_gfba  *EncryptOptions
	_eeee  *PdfInfo
}

// PdfAnnotationUnderline represents Underline annotations.
// (Section 12.5.6.10).
type PdfAnnotationUnderline struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _ebb.PdfObject
}

// GetStandardApplier gets currently used StandardApplier..
func (_aaaac *PdfWriter) GetStandardApplier() StandardApplier { return _aaaac._cafac }

// Duplicate creates a duplicate page based on the current one and returns it.
func (_cced *PdfPage) Duplicate() *PdfPage {
	_fcggf := *_cced
	_fcggf._cdbfde = _ebb.MakeDict()
	_fcggf._defbb = _ebb.MakeIndirectObject(_fcggf._cdbfde)
	return &_fcggf
}
func (_fgagc *PdfWriter) writeTrailer(_dcfae int) {
	_fgagc.writeString("\u0078\u0072\u0065\u0066\u000d\u000a")
	for _cbcaaa := 0; _cbcaaa <= _dcfae; {
		for ; _cbcaaa <= _dcfae; _cbcaaa++ {
			_abadab, _baedg := _fgagc._bedfc[_cbcaaa]
			if _baedg && (!_fgagc._abffb || _fgagc._abffb && (_abadab.Type == 1 && _abadab.Offset >= _fgagc._ggbfg || _abadab.Type == 0)) {
				break
			}
		}
		var _edefdg int
		for _edefdg = _cbcaaa + 1; _edefdg <= _dcfae; _edefdg++ {
			_geabd, _gdfa := _fgagc._bedfc[_edefdg]
			if _gdfa && (!_fgagc._abffb || _fgagc._abffb && (_geabd.Type == 1 && _geabd.Offset > _fgagc._ggbfg)) {
				continue
			}
			break
		}
		_bbda := _bg.Sprintf("\u0025d\u0020\u0025\u0064\u000d\u000a", _cbcaaa, _edefdg-_cbcaaa)
		_fgagc.writeString(_bbda)
		for _egege := _cbcaaa; _egege < _edefdg; _egege++ {
			_fcgec := _fgagc._bedfc[_egege]
			switch _fcgec.Type {
			case 0:
				_bbda = _bg.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u0066\u000d\u000a", 0, 65535)
				_fgagc.writeString(_bbda)
			case 1:
				_bbda = _bg.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u006e\u000d\u000a", _fcgec.Offset, 0)
				_fgagc.writeString(_bbda)
			}
		}
		_cbcaaa = _edefdg + 1
	}
	_bgcce := _ebb.MakeDict()
	_bgcce.Set("\u0049\u006e\u0066\u006f", _fgagc._eadfd)
	_bgcce.Set("\u0052\u006f\u006f\u0074", _fgagc._gegba)
	_bgcce.Set("\u0053\u0069\u007a\u0065", _ebb.MakeInteger(int64(_dcfae+1)))
	if _fgagc._abffb && _fgagc._bcage > 0 {
		_bgcce.Set("\u0050\u0072\u0065\u0076", _ebb.MakeInteger(_fgagc._bcage))
	}
	if _fgagc._cgfde != nil {
		_bgcce.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _fgagc._cbcaa)
	}
	if _fgagc._eecfe == nil && _fgagc._gfdea != "" && _fgagc._gffb != "" {
		_fgagc._eecfe = _ebb.MakeArray(_ebb.MakeHexString(_fgagc._gfdea), _ebb.MakeHexString(_fgagc._gffb))
	}
	if _fgagc._eecfe != nil {
		_bgcce.Set("\u0049\u0044", _fgagc._eecfe)
		_eg.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _fgagc._eecfe)
	}
	_fgagc.writeString("\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u000a")
	_fgagc.writeString(_bgcce.WriteString())
	_fgagc.writeString("\u000a")
}

// ToPdfObject implements interface PdfModel.
func (_ddgcc *PdfAnnotationTrapNet) ToPdfObject() _ebb.PdfObject {
	_ddgcc.PdfAnnotation.ToPdfObject()
	_abee := _ddgcc._bdcd
	_cfgc := _abee.PdfObject.(*_ebb.PdfObjectDictionary)
	_cfgc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0054r\u0061\u0070\u004e\u0065\u0074"))
	return _abee
}

// BytesToCharcodes converts the bytes in a PDF string to character codes.
func (_deaaf *PdfFont) BytesToCharcodes(data []byte) []_da.CharCode {
	_eg.Log.Trace("\u0042\u0079\u0074es\u0054\u006f\u0043\u0068\u0061\u0072\u0063\u006f\u0064e\u0073:\u0020d\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", data, data)
	if _fcffa, _dfe := _deaaf._ebcad.(*pdfFontType0); _dfe && _fcffa._efeb != nil {
		if _bacg, _egcfd := _fcffa.bytesToCharcodes(data); _egcfd {
			return _bacg
		}
	}
	var (
		_cfaad = make([]_da.CharCode, 0, len(data)+len(data)%2)
		_cadg  = _deaaf.baseFields()
	)
	if _cadg._dcdd != nil {
		if _fgcce, _bfgc := _cadg._dcdd.BytesToCharcodes(data); _bfgc {
			for _, _fafda := range _fgcce {
				_cfaad = append(_cfaad, _da.CharCode(_fafda))
			}
			return _cfaad
		}
	}
	if _cadg.isCIDFont() {
		if len(data) == 1 {
			data = []byte{0, data[0]}
		}
		if len(data)%2 != 0 {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0064\u0061\u0074\u0061\u003d\u0025\u002b\u0076\u0020t\u006f\u0020\u0065\u0076\u0065n\u0020\u006ce\u006e\u0067\u0074\u0068", data)
			data = append(data, 0)
		}
		for _cefgd := 0; _cefgd < len(data); _cefgd += 2 {
			_ffde := uint16(data[_cefgd])<<8 | uint16(data[_cefgd+1])
			_cfaad = append(_cfaad, _da.CharCode(_ffde))
		}
	} else {
		for _, _dgaag := range data {
			_cfaad = append(_cfaad, _da.CharCode(_dgaag))
		}
	}
	return _cfaad
}
func (_badda *PdfReader) newPdfAnnotationLineFromDict(_gdf *_ebb.PdfObjectDictionary) (*PdfAnnotationLine, error) {
	_dfc := PdfAnnotationLine{}
	_age, _bbaa := _badda.newPdfAnnotationMarkupFromDict(_gdf)
	if _bbaa != nil {
		return nil, _bbaa
	}
	_dfc.PdfAnnotationMarkup = _age
	_dfc.L = _gdf.Get("\u004c")
	_dfc.BS = _gdf.Get("\u0042\u0053")
	_dfc.LE = _gdf.Get("\u004c\u0045")
	_dfc.IC = _gdf.Get("\u0049\u0043")
	_dfc.LL = _gdf.Get("\u004c\u004c")
	_dfc.LLE = _gdf.Get("\u004c\u004c\u0045")
	_dfc.Cap = _gdf.Get("\u0043\u0061\u0070")
	_dfc.IT = _gdf.Get("\u0049\u0054")
	_dfc.LLO = _gdf.Get("\u004c\u004c\u004f")
	_dfc.CP = _gdf.Get("\u0043\u0050")
	_dfc.Measure = _gdf.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	_dfc.CO = _gdf.Get("\u0043\u004f")
	return &_dfc, nil
}

// K returns the value of the key component of the color.
func (_abdgd *PdfColorDeviceCMYK) K() float64 { return _abdgd[3] }

// ToPdfObject implements interface PdfModel.
func (_ebge *PdfAnnotationInk) ToPdfObject() _ebb.PdfObject {
	_ebge.PdfAnnotation.ToPdfObject()
	_aebf := _ebge._bdcd
	_bbbg := _aebf.PdfObject.(*_ebb.PdfObjectDictionary)
	_ebge.PdfAnnotationMarkup.appendToPdfDictionary(_bbbg)
	_bbbg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0049\u006e\u006b"))
	_bbbg.SetIfNotNil("\u0049n\u006b\u004c\u0069\u0073\u0074", _ebge.InkList)
	_bbbg.SetIfNotNil("\u0042\u0053", _ebge.BS)
	return _aebf
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
	ColorToRGB(_ggc PdfColor) (PdfColor, error)

	// GetNumComponents returns the number of components in the PdfColorspace.
	GetNumComponents() int

	// ToPdfObject returns a PdfObject representation of the PdfColorspace.
	ToPdfObject() _ebb.PdfObject

	// ColorFromPdfObjects returns a PdfColor in the given PdfColorspace from an array of PdfObject where each
	// PdfObject represents a numeric value.
	ColorFromPdfObjects(_deeb []_ebb.PdfObject) (PdfColor, error)

	// ColorFromFloats returns a new PdfColor based on input color components for a given PdfColorspace.
	ColorFromFloats(_ecee []float64) (PdfColor, error)

	// DecodeArray returns the Decode array for the PdfColorSpace, i.e. the range of each component.
	DecodeArray() []float64
}

// NewPdfColorspaceDeviceN returns an initialized PdfColorspaceDeviceN.
func NewPdfColorspaceDeviceN() *PdfColorspaceDeviceN { _bagc := &PdfColorspaceDeviceN{}; return _bagc }
func _agcbd(_eefg _ebb.PdfObject) (*PdfColorspaceSpecialSeparation, error) {
	_bcda := NewPdfColorspaceSpecialSeparation()
	if _cggg, _gedd := _eefg.(*_ebb.PdfIndirectObject); _gedd {
		_bcda._cded = _cggg
	}
	_eefg = _ebb.TraceToDirectObject(_eefg)
	_cfccb, _eecef := _eefg.(*_ebb.PdfObjectArray)
	if !_eecef {
		return nil, _bg.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062je\u0063\u0074")
	}
	if _cfccb.Len() != 4 {
		return nil, _bg.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074i\u006f\u006e \u0043\u0053\u003a\u0020\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
	}
	_eefg = _cfccb.Get(0)
	_agfc, _eecef := _eefg.(*_ebb.PdfObjectName)
	if !_eecef {
		return nil, _bg.Errorf("\u0073\u0065\u0070ar\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a \u0069n\u0076a\u006ci\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_agfc != "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e" {
		return nil, _bg.Errorf("\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020w\u0072o\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	_eefg = _cfccb.Get(1)
	_agfc, _eecef = _eefg.(*_ebb.PdfObjectName)
	if !_eecef {
		return nil, _bg.Errorf("\u0073\u0065pa\u0072\u0061\u0074i\u006f\u006e\u0020\u0043S: \u0049nv\u0061\u006c\u0069\u0064\u0020\u0063\u006flo\u0072\u0061\u006e\u0074\u0020\u006e\u0061m\u0065")
	}
	_bcda.ColorantName = _agfc
	_eefg = _cfccb.Get(2)
	_gfbg, _aedga := NewPdfColorspaceFromPdfObject(_eefg)
	if _aedga != nil {
		return nil, _aedga
	}
	_bcda.AlternateSpace = _gfbg
	_gcfc, _aedga := _aagg(_cfccb.Get(3))
	if _aedga != nil {
		return nil, _aedga
	}
	_bcda.TintTransform = _gcfc
	return _bcda, nil
}
func _abeb(_fegc *PdfField, _ggff _ebb.PdfObject) error {
	switch _fegc.GetContext().(type) {
	case *PdfFieldText:
		switch _cbbf := _ggff.(type) {
		case *_ebb.PdfObjectName:
			_afcc := _cbbf
			_eg.Log.Debug("\u0055\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u003a\u0020\u0047\u006f\u0074 \u0056\u0020\u0061\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u003e\u0020c\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f s\u0074\u0072\u0069\u006e\u0067\u0020\u0027\u0025\u0073\u0027", _afcc.String())
			_fegc.V = _ebb.MakeEncodedString(_cbbf.String(), true)
		case *_ebb.PdfObjectString:
			_fegc.V = _ebb.MakeEncodedString(_cbbf.String(), true)
		default:
			_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0056\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u0028\u0025\u0023\u0076\u0029", _cbbf, _cbbf)
		}
	case *PdfFieldButton:
		switch _ggff.(type) {
		case *_ebb.PdfObjectName:
			if len(_ggff.String()) > 0 {
				_fegc.V = _ggff
				_bgfag(_fegc, _ggff)
			}
		case *_ebb.PdfObjectString:
			if len(_ggff.String()) > 0 {
				_fegc.V = _ebb.MakeName(_ggff.String())
				_bgfag(_fegc, _fegc.V)
			}
		default:
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _fegc.PartialName(), _ggff)
			_fegc.V = _ggff
		}
	case *PdfFieldChoice:
		switch _ggff.(type) {
		case *_ebb.PdfObjectName:
			if len(_ggff.String()) > 0 {
				_fegc.V = _ebb.MakeString(_ggff.String())
				_bgfag(_fegc, _ggff)
			}
		case *_ebb.PdfObjectString:
			if len(_ggff.String()) > 0 {
				_fegc.V = _ggff
				_bgfag(_fegc, _ebb.MakeName(_ggff.String()))
			}
		default:
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _fegc.PartialName(), _ggff)
			_fegc.V = _ggff
		}
	case *PdfFieldSignature:
		_eg.Log.Debug("\u0054\u004f\u0044\u004f\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0061\u0070\u0070e\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0079\u0065\u0074\u003a\u0020\u0025\u0073\u002f\u0025v", _fegc.PartialName(), _ggff)
	}
	return nil
}

// NewCustomPdfOutputIntent creates a new custom PdfOutputIntent.
func NewCustomPdfOutputIntent(outputCondition, outputConditionIdentifier, info string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, Info: info, DestOutputProfile: destOutputProfile, _faeb: _ebb.MakeDict(), ColorComponents: colorComponents}
}

// GetIndirectObjectByNumber retrieves and returns a specific PdfObject by object number.
func (_bfeeb *PdfReader) GetIndirectObjectByNumber(number int) (_ebb.PdfObject, error) {
	_bbcfe, _aefbg := _bfeeb._cafdf.LookupByNumber(number)
	return _bbcfe, _aefbg
}

type pdfCIDFontType2 struct {
	fontCommon
	_ddea *_ebb.PdfIndirectObject
	_aacb _da.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// Dictionary that defines the character collection of the CIDFont (required).
	// See Table 116.
	CIDSystemInfo *_ebb.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW  _ebb.PdfObject
	W   _ebb.PdfObject
	DW2 _ebb.PdfObject
	W2  _ebb.PdfObject

	// CIDs to glyph indices mapping (optional).
	CIDToGIDMap _ebb.PdfObject
	_dgbc       map[_da.CharCode]float64
	_bagcb      float64
	_dceb       map[rune]int
}

// Normalize swaps (Llx,Urx) if Urx < Llx, and (Lly,Ury) if Ury < Lly.
func (_fabab *PdfRectangle) Normalize() {
	if _fabab.Llx > _fabab.Urx {
		_fabab.Llx, _fabab.Urx = _fabab.Urx, _fabab.Llx
	}
	if _fabab.Lly > _fabab.Ury {
		_fabab.Lly, _fabab.Ury = _fabab.Ury, _fabab.Lly
	}
}

// NewPdfActionNamed returns a new "named" action.
func NewPdfActionNamed() *PdfActionNamed {
	_edb := NewPdfAction()
	_eea := &PdfActionNamed{}
	_eea.PdfAction = _edb
	_edb.SetContext(_eea)
	return _eea
}

// GetPdfInfo returns the PDF info dictionary.
func (_bgaec *PdfReader) GetPdfInfo() (*PdfInfo, error) {
	_fbaee, _bcad := _bgaec.GetTrailer()
	if _bcad != nil {
		return nil, _bcad
	}
	var _eecgb *_ebb.PdfObjectDictionary
	_fdeab := _fbaee.Get("\u0049\u006e\u0066\u006f")
	switch _adaf := _fdeab.(type) {
	case *_ebb.PdfObjectReference:
		_ffgac := _adaf
		_fdeab, _bcad = _bgaec.GetIndirectObjectByNumber(int(_ffgac.ObjectNumber))
		_fdeab = _ebb.TraceToDirectObject(_fdeab)
		if _bcad != nil {
			return nil, _bcad
		}
		_eecgb, _ = _fdeab.(*_ebb.PdfObjectDictionary)
	case *_ebb.PdfObjectDictionary:
		_eecgb = _adaf
	}
	if _eecgb == nil {
		return nil, _gf.New("I\u006e\u0066\u006f\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0070r\u0065s\u0065\u006e\u0074")
	}
	_afgec, _bcad := NewPdfInfoFromObject(_eecgb)
	if _bcad != nil {
		return nil, _bcad
	}
	return _afgec, nil
}

type pdfFontSimple struct {
	fontCommon
	_fafaf *_ebb.PdfIndirectObject
	_cdff  map[_da.CharCode]float64
	_ebcb  _da.TextEncoder
	_dacee _da.TextEncoder
	_adbd  *PdfFontDescriptor

	// Encoding is subject to limitations that are described in 9.6.6, "Character Encoding".
	// BaseFont is derived differently.
	FirstChar _ebb.PdfObject
	LastChar  _ebb.PdfObject
	Widths    _ebb.PdfObject
	Encoding  _ebb.PdfObject
	_ddgd     *_bad.RuneCharSafeMap
}

var (
	_aabec = _a.MustCompile("\u005cd\u002b\u0020\u0064\u0069c\u0074\u005c\u0073\u002b\u0028d\u0075p\u005cs\u002b\u0029\u003f\u0062\u0065\u0067\u0069n")
	_afgea = _a.MustCompile("\u005e\u005cs\u002a\u002f\u0028\u005c\u0053\u002b\u003f\u0029\u005c\u0073\u002b\u0028\u002e\u002b\u003f\u0029\u005c\u0073\u002b\u0064\u0065\u0066\\s\u002a\u0024")
	_bdddd = _a.MustCompile("\u005e\u005c\u0073*\u0064\u0075\u0070\u005c\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002a\u002f\u0028\u005c\u0077\u002b\u003f\u0029\u0028\u003f\u003a\u005c\u002e\u005c\u0064\u002b)\u003f\u005c\u0073\u002b\u0070\u0075\u0074\u0024")
	_aeecf = "\u002f\u0045\u006e\u0063od\u0069\u006e\u0067\u0020\u0032\u0035\u0036\u0020\u0061\u0072\u0072\u0061\u0079"
	_ccdgc = "\u0072\u0065\u0061d\u006f\u006e\u006c\u0079\u0020\u0064\u0065\u0066"
	_cfgg  = "\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u0066\u0069\u006c\u0065\u0020e\u0065\u0078\u0065\u0063"
)

// GetBorderWidth returns the border style's width.
func (_fbbg *PdfBorderStyle) GetBorderWidth() float64 {
	if _fbbg.W == nil {
		return 1
	}
	return *_fbbg.W
}
func (_eca *PdfReader) newPdfAnnotationScreenFromDict(_aace *_ebb.PdfObjectDictionary) (*PdfAnnotationScreen, error) {
	_eded := PdfAnnotationScreen{}
	_eded.T = _aace.Get("\u0054")
	_eded.MK = _aace.Get("\u004d\u004b")
	_eded.A = _aace.Get("\u0041")
	_eded.AA = _aace.Get("\u0041\u0041")
	return &_eded, nil
}
func (_cdd *PdfReader) newPdfAnnotationStampFromDict(_bca *_ebb.PdfObjectDictionary) (*PdfAnnotationStamp, error) {
	_eacc := PdfAnnotationStamp{}
	_ccd, _geg := _cdd.newPdfAnnotationMarkupFromDict(_bca)
	if _geg != nil {
		return nil, _geg
	}
	_eacc.PdfAnnotationMarkup = _ccd
	_eacc.Name = _bca.Get("\u004e\u0061\u006d\u0065")
	return &_eacc, nil
}
func (_afec *PdfReader) loadPerms() (*Permissions, error) {
	if _bcggb := _afec._fdgda.Get("\u0050\u0065\u0072m\u0073"); _bcggb != nil {
		if _ddcbf, _eggfd := _ebb.GetDict(_bcggb); _eggfd {
			_cbgaa := _ddcbf.Get("\u0044\u006f\u0063\u004d\u0044\u0050")
			if _cbgaa == nil {
				return nil, nil
			}
			if _edfdd, _geecf := _ebb.GetIndirect(_cbgaa); _geecf {
				_gcbeb, _cecee := _afec.newPdfSignatureFromIndirect(_edfdd)
				if _cecee != nil {
					return nil, _cecee
				}
				return NewPermissions(_gcbeb), nil
			}
			return nil, _bg.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u006f\u0063M\u0044\u0050\u0020\u0065nt\u0072\u0079")
		}
		return nil, _bg.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0050\u0065\u0072\u006d\u0073\u0020\u0065\u006e\u0074\u0072\u0079")
	}
	return nil, nil
}

// GetCatalogMarkInfo gets catalog MarkInfo object.
func (_cdcgb *PdfReader) GetCatalogMarkInfo() (_ebb.PdfObject, bool) {
	if _cdcgb._fdgda == nil {
		return nil, false
	}
	_fdgb := _cdcgb._fdgda.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
	return _fdgb, _fdgb != nil
}

// Compress is yet to be implemented.
// Should be able to compress in terms of JPEG quality parameter,
// and DPI threshold (need to know bounding area dimensions).
func (_dcbgd DefaultImageHandler) Compress(input *Image, quality int64) (*Image, error) {
	return input, nil
}
func _ggddd(_bbea _ebb.PdfObject) (*PdfBorderStyle, error) {
	_bfaa := &PdfBorderStyle{}
	_bfaa._dgaa = _bbea
	var _egbb *_ebb.PdfObjectDictionary
	_bbea = _ebb.TraceToDirectObject(_bbea)
	_egbb, _ecdf := _bbea.(*_ebb.PdfObjectDictionary)
	if !_ecdf {
		return nil, _gf.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	if _efccb := _egbb.Get("\u0054\u0079\u0070\u0065"); _efccb != nil {
		_accc, _bdgc := _efccb.(*_ebb.PdfObjectName)
		if !_bdgc {
			_eg.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u0077\u0069th\u0020\u0054\u0079\u0070e\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061me\u0020\u006fb\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0054", _efccb)
		} else {
			if *_accc != "\u0042\u006f\u0072\u0064\u0065\u0072" {
				_eg.Log.Debug("W\u0061\u0072\u006e\u0069\u006e\u0067,\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020B\u006f\u0072\u0064e\u0072:\u0020\u0025\u0073", *_accc)
			}
		}
	}
	if _fee := _egbb.Get("\u0057"); _fee != nil {
		_gfg, _dfbe := _ebb.GetNumberAsFloat(_fee)
		if _dfbe != nil {
			_eg.Log.Debug("\u0045\u0072\u0072\u006fr \u0072\u0065\u0074\u0072\u0069\u0065\u0076\u0069\u006e\u0067\u0020\u0057\u003a\u0020%\u0076", _dfbe)
			return nil, _dfbe
		}
		_bfaa.W = &_gfg
	}
	if _fab := _egbb.Get("\u0053"); _fab != nil {
		_egcc, _dbae := _fab.(*_ebb.PdfObjectName)
		if !_dbae {
			return nil, _gf.New("\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u0053\u0020\u006e\u006ft\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		var _gcgb BorderStyle
		switch *_egcc {
		case "\u0053":
			_gcgb = BorderStyleSolid
		case "\u0044":
			_gcgb = BorderStyleDashed
		case "\u0042":
			_gcgb = BorderStyleBeveled
		case "\u0049":
			_gcgb = BorderStyleInset
		case "\u0055":
			_gcgb = BorderStyleUnderline
		default:
			_eg.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020s\u0074\u0079\u006c\u0065\u0020n\u0061\u006d\u0065\u0020\u0025\u0073", *_egcc)
			return nil, _gf.New("\u0073\u0074\u0079\u006ce \u0074\u0079\u0070\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065c\u006b")
		}
		_bfaa.S = &_gcgb
	}
	if _ccfd := _egbb.Get("\u0044"); _ccfd != nil {
		_agcd, _eaeb := _ccfd.(*_ebb.PdfObjectArray)
		if !_eaeb {
			_eg.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044\u0020\u0064a\u0073\u0068\u0020\u006e\u006f\u0074\u0020a\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0054", _ccfd)
			return nil, _gf.New("\u0062o\u0072\u0064\u0065\u0072 \u0044\u0020\u0074\u0079\u0070e\u0020c\u0068e\u0063\u006b\u0020\u0065\u0072\u0072\u006fr")
		}
		_dcbe, _deec := _agcd.ToIntegerArray()
		if _deec != nil {
			_eg.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044 \u0050\u0072\u006fbl\u0065\u006d\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u006e\u0074\u0065\u0067e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u003a \u0025\u0076", _deec)
			return nil, _deec
		}
		_bfaa.D = &_dcbe
	}
	return _bfaa, nil
}
func (_ccafe *PdfWriter) setDocumentIDs(_geaeb, _dbgdd string) {
	_ccafe._eecfe = _ebb.MakeArray(_ebb.MakeHexString(_geaeb), _ebb.MakeHexString(_dbgdd))
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_efdf pdfCIDFontType0) GetCharMetrics(code _da.CharCode) (_bad.CharMetrics, bool) {
	_defe := _efdf._gbdb
	if _gbfe, _bbce := _efdf._afcac[code]; _bbce {
		_defe = _gbfe
	}
	return _bad.CharMetrics{Wx: _defe}, true
}

// NewPdfActionGoToE returns a new "go to embedded" action.
func NewPdfActionGoToE() *PdfActionGoToE {
	_cf := NewPdfAction()
	_dga := &PdfActionGoToE{}
	_dga.PdfAction = _cf
	_cf.SetContext(_dga)
	return _dga
}

// ImageToRGB converts an Image in a given PdfColorspace to an RGB image.
func (_gabcf *PdfColorspaceDeviceN) ImageToRGB(img Image) (Image, error) {
	_fagf := _abg.NewReader(img.getBase())
	_cfebf := _dg.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, nil, img._dagcb, img._dgcea)
	_eeefc := _abg.NewWriter(_cfebf)
	_gebdg := _cbg.Pow(2, float64(img.BitsPerComponent)) - 1
	_feg := _gabcf.GetNumComponents()
	_ceaca := make([]uint32, _feg)
	_aege := make([]float64, _feg)
	for {
		_cgcea := _fagf.ReadSamples(_ceaca)
		if _cgcea == _ab.EOF {
			break
		} else if _cgcea != nil {
			return img, _cgcea
		}
		for _bgdf := 0; _bgdf < _feg; _bgdf++ {
			_agdc := float64(_ceaca[_bgdf]) / _gebdg
			_aege[_bgdf] = _agdc
		}
		_ebgee, _cgcea := _gabcf.TintTransform.Evaluate(_aege)
		if _cgcea != nil {
			return img, _cgcea
		}
		for _, _fbfda := range _ebgee {
			_fbfda = _cbg.Min(_cbg.Max(0, _fbfda), 1.0)
			if _cgcea = _eeefc.WriteSample(uint32(_fbfda * _gebdg)); _cgcea != nil {
				return img, _cgcea
			}
		}
	}
	return _gabcf.AlternateSpace.ImageToRGB(_afacb(&_cfebf))
}

// NewReaderForText makes a new PdfReader for an input PDF content string. For use in testing.
func NewReaderForText(txt string) *PdfReader {
	return &PdfReader{_dfadc: map[_ebb.PdfObject]struct{}{}, _abbaca: _fadcd(), _cafdf: _ebb.NewParserFromString(txt)}
}

// GetContainingPdfObject implements interface PdfModel.
func (_gdgag *PdfAnnotation) GetContainingPdfObject() _ebb.PdfObject { return _gdgag._bdcd }
func _bafec(_gegdg string) (string, error) {
	var _gacbg _ca.Buffer
	_gacbg.WriteString(_gegdg)
	_cabeb := make([]byte, 8+16)
	_cfdced := _f.Now().UTC().UnixNano()
	_cb.BigEndian.PutUint64(_cabeb, uint64(_cfdced))
	_, _efbgg := _gd.Read(_cabeb[8:])
	if _efbgg != nil {
		return "", _efbgg
	}
	_gacbg.WriteString(_d.EncodeToString(_cabeb))
	return _gacbg.String(), nil
}

// Val returns the color value.
func (_ecef *PdfColorDeviceGray) Val() float64 { return float64(*_ecef) }

// BaseFont returns the font's "BaseFont" field.
func (_ecbad *PdfFont) BaseFont() string { return _ecbad.baseFields()._fdacg }

// ToPdfObject converts PdfAcroForm to a PdfObject, i.e. an indirect object containing the
// AcroForm dictionary.
func (_bbbed *PdfAcroForm) ToPdfObject() _ebb.PdfObject {
	_agdgc := _bbbed._adcg
	_cacfb := _agdgc.PdfObject.(*_ebb.PdfObjectDictionary)
	if _bbbed.Fields != nil {
		_egge := _ebb.PdfObjectArray{}
		for _, _bcacbc := range *_bbbed.Fields {
			_fefg := _bcacbc.GetContext()
			if _fefg != nil {
				_egge.Append(_fefg.ToPdfObject())
			} else {
				_egge.Append(_bcacbc.ToPdfObject())
			}
		}
		_cacfb.Set("\u0046\u0069\u0065\u006c\u0064\u0073", &_egge)
	}
	if _bbbed.NeedAppearances != nil {
		_cacfb.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _bbbed.NeedAppearances)
	}
	if _bbbed.SigFlags != nil {
		_cacfb.Set("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073", _bbbed.SigFlags)
	}
	if _bbbed.CO != nil {
		_cacfb.Set("\u0043\u004f", _bbbed.CO)
	}
	if _bbbed.DR != nil {
		_cacfb.Set("\u0044\u0052", _bbbed.DR.ToPdfObject())
	}
	if _bbbed.DA != nil {
		_cacfb.Set("\u0044\u0041", _bbbed.DA)
	}
	if _bbbed.Q != nil {
		_cacfb.Set("\u0051", _bbbed.Q)
	}
	if _bbbed.XFA != nil {
		_cacfb.Set("\u0058\u0046\u0041", _bbbed.XFA)
	}
	return _agdgc
}

// NewPdfActionImportData returns a new "import data" action.
func NewPdfActionImportData() *PdfActionImportData {
	_adc := NewPdfAction()
	_eeaa := &PdfActionImportData{}
	_eeaa.PdfAction = _adc
	_adc.SetContext(_eeaa)
	return _eeaa
}

// GetTrailer returns the PDF's trailer dictionary.
func (_dcgg *PdfReader) GetTrailer() (*_ebb.PdfObjectDictionary, error) {
	_dagbe := _dcgg._cafdf.GetTrailer()
	if _dagbe == nil {
		return nil, _gf.New("\u0074r\u0061i\u006c\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	return _dagbe, nil
}

// PdfSignature represents a PDF signature dictionary and is used for signing via form signature fields.
// (Section 12.8, Table 252 - Entries in a signature dictionary p. 475 in PDF32000_2008).
type PdfSignature struct {
	Handler SignatureHandler
	_ffbgc  *_ebb.PdfIndirectObject

	// Type: Sig/DocTimeStamp
	Type         *_ebb.PdfObjectName
	Filter       *_ebb.PdfObjectName
	SubFilter    *_ebb.PdfObjectName
	Contents     *_ebb.PdfObjectString
	Cert         _ebb.PdfObject
	ByteRange    *_ebb.PdfObjectArray
	Reference    *_ebb.PdfObjectArray
	Changes      *_ebb.PdfObjectArray
	Name         *_ebb.PdfObjectString
	M            *_ebb.PdfObjectString
	Location     *_ebb.PdfObjectString
	Reason       *_ebb.PdfObjectString
	ContactInfo  *_ebb.PdfObjectString
	R            *_ebb.PdfObjectInteger
	V            *_ebb.PdfObjectInteger
	PropBuild    *_ebb.PdfObjectDictionary
	PropAuthTime *_ebb.PdfObjectInteger
	PropAuthType *_ebb.PdfObjectName
}

// SetForms sets the Acroform for a PDF file.
func (_efbd *PdfWriter) SetForms(form *PdfAcroForm) error { _efbd._dbea = form; return nil }

// A returns the value of the A component of the color.
func (_gad *PdfColorLab) A() float64 { return _gad[1] }

// ImageToRGB converts ICCBased colorspace image to RGB and returns the result.
func (_fbdd *PdfColorspaceICCBased) ImageToRGB(img Image) (Image, error) {
	if _fbdd.Alternate == nil {
		_eg.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _fbdd.N == 1 {
			_eg.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_dbffc := NewPdfColorspaceDeviceGray()
			return _dbffc.ImageToRGB(img)
		} else if _fbdd.N == 3 {
			_eg.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return img, nil
		} else if _fbdd.N == 4 {
			_eg.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_dfdf := NewPdfColorspaceDeviceCMYK()
			return _dfdf.ImageToRGB(img)
		} else {
			return img, _gf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_eg.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _fbdd)
	_eggd, _bdfb := _fbdd.Alternate.ImageToRGB(img)
	_eg.Log.Trace("I\u0043C\u0020\u0049\u006e\u0070\u0075\u0074\u0020\u0069m\u0061\u0067\u0065\u003a %\u002b\u0076", img)
	_eg.Log.Trace("I\u0043\u0043\u0020\u004fut\u0070u\u0074\u0020\u0069\u006d\u0061g\u0065\u003a\u0020\u0025\u002b\u0076", _eggd)
	return _eggd, _bdfb
}

// GetAlphabet returns a map of the runes in `text` and their frequencies.
func GetAlphabet(text string) map[rune]int {
	_aegef := map[rune]int{}
	for _, _aafa := range text {
		_aegef[_aafa]++
	}
	return _aegef
}

// PdfColorspaceLab is a L*, a*, b* 3 component colorspace.
type PdfColorspaceLab struct {
	WhitePoint []float64
	BlackPoint []float64
	Range      []float64
	_bdga      *_ebb.PdfIndirectObject
}

// PdfInfo holds document information that will overwrite
// document information global variables defined above.
type PdfInfo struct {
	Title        *_ebb.PdfObjectString
	Author       *_ebb.PdfObjectString
	Subject      *_ebb.PdfObjectString
	Keywords     *_ebb.PdfObjectString
	Creator      *_ebb.PdfObjectString
	Producer     *_ebb.PdfObjectString
	CreationDate *PdfDate
	ModifiedDate *PdfDate
	Trapped      *_ebb.PdfObjectName
	_gcgf        *_ebb.PdfObjectDictionary
}

// PdfColorspaceDeviceCMYK represents a CMYK32 colorspace.
type PdfColorspaceDeviceCMYK struct{}

// PdfAnnotationTrapNet represents TrapNet annotations.
// (Section 12.5.6.21).
type PdfAnnotationTrapNet struct{ *PdfAnnotation }

// NewXObjectImage returns a new XObjectImage.
func NewXObjectImage() *XObjectImage {
	_daag := &XObjectImage{}
	_cdagc := &_ebb.PdfObjectStream{}
	_cdagc.PdfObjectDictionary = _ebb.MakeDict()
	_daag._fbeec = _cdagc
	return _daag
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_aaea *PdfShadingType3) ToPdfObject() _ebb.PdfObject {
	_aaea.PdfShading.ToPdfObject()
	_caebg, _fbbfe := _aaea.getShadingDict()
	if _fbbfe != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _aaea.Coords != nil {
		_caebg.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _aaea.Coords)
	}
	if _aaea.Domain != nil {
		_caebg.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _aaea.Domain)
	}
	if _aaea.Function != nil {
		if len(_aaea.Function) == 1 {
			_caebg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _aaea.Function[0].ToPdfObject())
		} else {
			_egfcb := _ebb.MakeArray()
			for _, _eaaag := range _aaea.Function {
				_egfcb.Append(_eaaag.ToPdfObject())
			}
			_caebg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _egfcb)
		}
	}
	if _aaea.Extend != nil {
		_caebg.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _aaea.Extend)
	}
	return _aaea._fbfae
}

// GetPageAsIndirectObject returns the page as a dictionary within an PdfIndirectObject.
func (_aeffa *PdfPage) GetPageAsIndirectObject() *_ebb.PdfIndirectObject { return _aeffa._defbb }

// AddExtension adds the specified extension to the Extensions dictionary.
// See section 7.1.2 "Extensions Dictionary" (pp. 108-109 PDF32000_2008).
func (_ececd *PdfWriter) AddExtension(extName, baseVersion string, extLevel int) {
	_babed, _fbeff := _ebb.GetDict(_ececd._dffegd.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073"))
	if !_fbeff {
		_babed = _ebb.MakeDict()
		_ececd._dffegd.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073", _babed)
	}
	_dgdbb, _fbeff := _ebb.GetDict(_babed.Get(_ebb.PdfObjectName(extName)))
	if !_fbeff {
		_dgdbb = _ebb.MakeDict()
		_babed.Set(_ebb.PdfObjectName(extName), _dgdbb)
	}
	if _gebaf, _ := _ebb.GetNameVal(_dgdbb.Get("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e")); _gebaf != baseVersion {
		_dgdbb.Set("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e", _ebb.MakeName(baseVersion))
	}
	if _dgage, _ := _ebb.GetIntVal(_dgdbb.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c")); _dgage != extLevel {
		_dgdbb.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c", _ebb.MakeInteger(int64(extLevel)))
	}
}

// ToPdfObject implements interface PdfModel.
func (_cgacb *PdfAnnotationSquare) ToPdfObject() _ebb.PdfObject {
	_cgacb.PdfAnnotation.ToPdfObject()
	_ebaa := _cgacb._bdcd
	_bdca := _ebaa.PdfObject.(*_ebb.PdfObjectDictionary)
	if _cgacb.PdfAnnotationMarkup != nil {
		_cgacb.PdfAnnotationMarkup.appendToPdfDictionary(_bdca)
	}
	_bdca.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0053\u0071\u0075\u0061\u0072\u0065"))
	_bdca.SetIfNotNil("\u0042\u0053", _cgacb.BS)
	_bdca.SetIfNotNil("\u0049\u0043", _cgacb.IC)
	_bdca.SetIfNotNil("\u0042\u0045", _cgacb.BE)
	_bdca.SetIfNotNil("\u0052\u0044", _cgacb.RD)
	return _ebaa
}

// PdfAnnotationPopup represents Popup annotations.
// (Section 12.5.6.14).
type PdfAnnotationPopup struct {
	*PdfAnnotation
	Parent _ebb.PdfObject
	Open   _ebb.PdfObject
}

// PdfColorDeviceRGB represents a color in DeviceRGB colorspace with R, G, B components, where component is
// defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceRGB [3]float64

func (_bcbfg *Image) samplesAddPadding(_gcga []uint32) []uint32 {
	_bcead := _dg.BytesPerLine(int(_bcbfg.Width), int(_bcbfg.BitsPerComponent), _bcbfg.ColorComponents) * (8 / int(_bcbfg.BitsPerComponent))
	_geda := _bcead * int(_bcbfg.Height)
	if len(_gcga) == _geda {
		return _gcga
	}
	_bcgae := make([]uint32, _geda)
	_adabg := int(_bcbfg.Width) * _bcbfg.ColorComponents
	for _cfegc := 0; _cfegc < int(_bcbfg.Height); _cfegc++ {
		_efbga := _cfegc * int(_bcbfg.Width)
		_edff := _cfegc * _bcead
		for _affb := 0; _affb < _adabg; _affb++ {
			_bcgae[_edff+_affb] = _gcga[_efbga+_affb]
		}
	}
	return _bcgae
}
func (_addg *PdfReader) newPdfAnnotationStrikeOut(_feb *_ebb.PdfObjectDictionary) (*PdfAnnotationStrikeOut, error) {
	_eda := PdfAnnotationStrikeOut{}
	_bcea, _daga := _addg.newPdfAnnotationMarkupFromDict(_feb)
	if _daga != nil {
		return nil, _daga
	}
	_eda.PdfAnnotationMarkup = _bcea
	_eda.QuadPoints = _feb.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_eda, nil
}
func _egaaa(_gggfe _ebb.PdfObject) (*PdfColorspaceDeviceN, error) {
	_aae := NewPdfColorspaceDeviceN()
	if _acdb, _cacc := _gggfe.(*_ebb.PdfIndirectObject); _cacc {
		_aae._gebb = _acdb
	}
	_gggfe = _ebb.TraceToDirectObject(_gggfe)
	_eeagg, _caeeg := _gggfe.(*_ebb.PdfObjectArray)
	if !_caeeg {
		return nil, _bg.Errorf("\u0064\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if _eeagg.Len() != 4 && _eeagg.Len() != 5 {
		return nil, _bg.Errorf("\u0064\u0065\u0076ic\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0049\u006ec\u006fr\u0072e\u0063t\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_gggfe = _eeagg.Get(0)
	_efce, _caeeg := _gggfe.(*_ebb.PdfObjectName)
	if !_caeeg {
		return nil, _bg.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_efce != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
		return nil, _bg.Errorf("\u0064\u0065v\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020na\u006d\u0065")
	}
	_gggfe = _eeagg.Get(1)
	_gggfe = _ebb.TraceToDirectObject(_gggfe)
	_faeg, _caeeg := _gggfe.(*_ebb.PdfObjectArray)
	if !_caeeg {
		return nil, _bg.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0049\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_aae.ColorantNames = _faeg
	_gggfe = _eeagg.Get(2)
	_bcged, _becc := NewPdfColorspaceFromPdfObject(_gggfe)
	if _becc != nil {
		return nil, _becc
	}
	_aae.AlternateSpace = _bcged
	_cgdaa, _becc := _aagg(_eeagg.Get(3))
	if _becc != nil {
		return nil, _becc
	}
	_aae.TintTransform = _cgdaa
	if _eeagg.Len() == 5 {
		_aeabg, _fdfcg := _gbgb(_eeagg.Get(4))
		if _fdfcg != nil {
			return nil, _fdfcg
		}
		_aae.Attributes = _aeabg
	}
	return _aae, nil
}
func _fcgc(_gaea _ebb.PdfObject) (*PdfColorspaceCalRGB, error) {
	_ccae := NewPdfColorspaceCalRGB()
	if _fabc, _gbb := _gaea.(*_ebb.PdfIndirectObject); _gbb {
		_ccae._aeac = _fabc
	}
	_gaea = _ebb.TraceToDirectObject(_gaea)
	_cedg, _gbacgg := _gaea.(*_ebb.PdfObjectArray)
	if !_gbacgg {
		return nil, _bg.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _cedg.Len() != 2 {
		return nil, _bg.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_gaea = _ebb.TraceToDirectObject(_cedg.Get(0))
	_eaced, _gbacgg := _gaea.(*_ebb.PdfObjectName)
	if !_gbacgg {
		return nil, _bg.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	if *_eaced != "\u0043\u0061\u006c\u0052\u0047\u0042" {
		return nil, _bg.Errorf("\u006e\u006f\u0074 a\u0020\u0043\u0061\u006c\u0052\u0047\u0042\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_gaea = _ebb.TraceToDirectObject(_cedg.Get(1))
	_egbe, _gbacgg := _gaea.(*_ebb.PdfObjectDictionary)
	if !_gbacgg {
		return nil, _bg.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	_gaea = _egbe.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_gaea = _ebb.TraceToDirectObject(_gaea)
	_fgbe, _gbacgg := _gaea.(*_ebb.PdfObjectArray)
	if !_gbacgg {
		return nil, _bg.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050o\u0069\u006e\u0074")
	}
	if _fgbe.Len() != 3 {
		return nil, _bg.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0057h\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_ebbf, _cgba := _fgbe.GetAsFloat64Slice()
	if _cgba != nil {
		return nil, _cgba
	}
	_ccae.WhitePoint = _ebbf
	_gaea = _egbe.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _gaea != nil {
		_gaea = _ebb.TraceToDirectObject(_gaea)
		_ccaee, _gaad := _gaea.(*_ebb.PdfObjectArray)
		if !_gaad {
			return nil, _bg.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050o\u0069\u006e\u0074")
		}
		if _ccaee.Len() != 3 {
			return nil, _bg.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0042l\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		_fccgf, _gafg := _ccaee.GetAsFloat64Slice()
		if _gafg != nil {
			return nil, _gafg
		}
		_ccae.BlackPoint = _fccgf
	}
	_gaea = _egbe.Get("\u0047\u0061\u006dm\u0061")
	if _gaea != nil {
		_gaea = _ebb.TraceToDirectObject(_gaea)
		_gaag, _fcfc := _gaea.(*_ebb.PdfObjectArray)
		if !_fcfc {
			return nil, _bg.Errorf("C\u0061\u006c\u0052\u0047B:\u0020I\u006e\u0076\u0061\u006c\u0069d\u0020\u0047\u0061\u006d\u006d\u0061")
		}
		if _gaag.Len() != 3 {
			return nil, _bg.Errorf("C\u0061\u006c\u0052\u0047\u0042\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0047a\u006d\u006d\u0061 \u0061r\u0072\u0061\u0079")
		}
		_bgfd, _acgd := _gaag.GetAsFloat64Slice()
		if _acgd != nil {
			return nil, _acgd
		}
		_ccae.Gamma = _bgfd
	}
	_gaea = _egbe.Get("\u004d\u0061\u0074\u0072\u0069\u0078")
	if _gaea != nil {
		_gaea = _ebb.TraceToDirectObject(_gaea)
		_gag, _dbdf := _gaea.(*_ebb.PdfObjectArray)
		if !_dbdf {
			return nil, _bg.Errorf("\u0043\u0061\u006c\u0052GB\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004d\u0061\u0074\u0072i\u0078")
		}
		if _gag.Len() != 9 {
			_eg.Log.Error("\u004d\u0061t\u0072\u0069\u0078 \u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _gag.String())
			return nil, _bg.Errorf("\u0043\u0061\u006c\u0052G\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u004da\u0074\u0072\u0069\u0078\u0020\u0061\u0072r\u0061\u0079")
		}
		_ebba, _fcfbg := _gag.GetAsFloat64Slice()
		if _fcfbg != nil {
			return nil, _fcfbg
		}
		_ccae.Matrix = _ebba
	}
	return _ccae, nil
}

// GetPdfName returns the PDF name used to indicate the border style.
// (Table 166 p. 395).
func (_acfcc *BorderStyle) GetPdfName() string {
	switch *_acfcc {
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

// GetFontDescriptor returns the font descriptor for `font`.
func (_ffdff PdfFont) GetFontDescriptor() (*PdfFontDescriptor, error) {
	return _ffdff._ebcad.getFontDescriptor(), nil
}

// FlattenFields flattens the form fields and annotations for the PDF loaded in `pdf` and makes
// non-editable.
// Looks up all widget annotations corresponding to form fields and flattens them by drawing the content
// through the content stream rather than annotations.
// References to flattened annotations will be removed from Page Annots array. For fields the AcroForm entry
// will be emptied.
// When `allannots` is true, all annotations will be flattened. Keep false if want to keep non-form related
// annotations intact.
// When `appgen` is not nil, it will be used to generate appearance streams for the field annotations.
func (_aeeg *PdfReader) FlattenFields(allannots bool, appgen FieldAppearanceGenerator) error {
	return _aeeg.flattenFieldsWithOpts(allannots, appgen, nil)
}
func (_dbac *PdfReader) newPdfAnnotationRedactFromDict(_dfb *_ebb.PdfObjectDictionary) (*PdfAnnotationRedact, error) {
	_gecd := PdfAnnotationRedact{}
	_bdag, _fagg := _dbac.newPdfAnnotationMarkupFromDict(_dfb)
	if _fagg != nil {
		return nil, _fagg
	}
	_gecd.PdfAnnotationMarkup = _bdag
	_gecd.QuadPoints = _dfb.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_gecd.IC = _dfb.Get("\u0049\u0043")
	_gecd.RO = _dfb.Get("\u0052\u004f")
	_gecd.OverlayText = _dfb.Get("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074")
	_gecd.Repeat = _dfb.Get("\u0052\u0065\u0070\u0065\u0061\u0074")
	_gecd.DA = _dfb.Get("\u0044\u0041")
	_gecd.Q = _dfb.Get("\u0051")
	return &_gecd, nil
}

// IsShading specifies if the pattern is a shading pattern.
func (_fbcgc *PdfPattern) IsShading() bool { return _fbcgc.PatternType == 2 }

// HasXObjectByName checks if an XObject with a specified keyName is defined.
func (_deaaa *PdfPageResources) HasXObjectByName(keyName _ebb.PdfObjectName) bool {
	_bbgge, _ := _deaaa.GetXObjectByName(keyName)
	return _bbgge != nil
}

// PdfActionJavaScript represents a javaScript action.
type PdfActionJavaScript struct {
	*PdfAction
	JS _ebb.PdfObject
}

func _ggbd(_cgggf *_ebb.PdfObjectDictionary) (*PdfTilingPattern, error) {
	_aefdc := &PdfTilingPattern{}
	_afdag := _cgggf.Get("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e")
	if _afdag == nil {
		_eg.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054\u0079\u0070\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bfed, _ffgcf := _afdag.(*_ebb.PdfObjectInteger)
	if !_ffgcf {
		_eg.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054y\u0070\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _afdag)
		return nil, _ebb.ErrTypeError
	}
	_aefdc.PaintType = _bfed
	_afdag = _cgggf.Get("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _afdag == nil {
		_eg.Log.Debug("\u0054i\u006ci\u006e\u0067\u0054\u0079\u0070e\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_eadcc, _ffgcf := _afdag.(*_ebb.PdfObjectInteger)
	if !_ffgcf {
		_eg.Log.Debug("\u0054\u0069\u006cin\u0067\u0054\u0079\u0070\u0065\u0020\u006e\u006f\u0074 \u0061n\u0020i\u006et\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _afdag)
		return nil, _ebb.ErrTypeError
	}
	_aefdc.TilingType = _eadcc
	_afdag = _cgggf.Get("\u0042\u0042\u006f\u0078")
	if _afdag == nil {
		_eg.Log.Debug("\u0042\u0042\u006fx\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_afdag = _ebb.TraceToDirectObject(_afdag)
	_cgcd, _ffgcf := _afdag.(*_ebb.PdfObjectArray)
	if !_ffgcf {
		_eg.Log.Debug("\u0042B\u006f\u0078 \u0073\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061y\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _afdag)
		return nil, _ebb.ErrTypeError
	}
	_fggbf, _dfcae := NewPdfRectangle(*_cgcd)
	if _dfcae != nil {
		_eg.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _dfcae)
		return nil, _dfcae
	}
	_aefdc.BBox = _fggbf
	_afdag = _cgggf.Get("\u0058\u0053\u0074e\u0070")
	if _afdag == nil {
		_eg.Log.Debug("\u0058\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fcce, _dfcae := _ebb.GetNumberAsFloat(_afdag)
	if _dfcae != nil {
		_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0058S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _fcce)
		return nil, _dfcae
	}
	_aefdc.XStep = _ebb.MakeFloat(_fcce)
	_afdag = _cgggf.Get("\u0059\u0053\u0074e\u0070")
	if _afdag == nil {
		_eg.Log.Debug("\u0059\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_aeagd, _dfcae := _ebb.GetNumberAsFloat(_afdag)
	if _dfcae != nil {
		_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0059S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _aeagd)
		return nil, _dfcae
	}
	_aefdc.YStep = _ebb.MakeFloat(_aeagd)
	_afdag = _cgggf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _afdag == nil {
		_eg.Log.Debug("\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_cgggf, _ffgcf = _ebb.TraceToDirectObject(_afdag).(*_ebb.PdfObjectDictionary)
	if !_ffgcf {
		return nil, _bg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _afdag)
	}
	_abegb, _dfcae := NewPdfPageResourcesFromDict(_cgggf)
	if _dfcae != nil {
		return nil, _dfcae
	}
	_aefdc.Resources = _abegb
	if _cgdef := _cgggf.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _cgdef != nil {
		_gdefg, _bdacc := _cgdef.(*_ebb.PdfObjectArray)
		if !_bdacc {
			_eg.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cgdef)
			return nil, _ebb.ErrTypeError
		}
		_aefdc.Matrix = _gdefg
	}
	return _aefdc, nil
}

// PdfColorDeviceCMYK is a CMYK32 color, where each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceCMYK [4]float64

const (
	TrappedUnknown PdfInfoTrapped = "\u0055n\u006b\u006e\u006f\u0077\u006e"
	TrappedTrue    PdfInfoTrapped = "\u0054\u0072\u0075\u0065"
	TrappedFalse   PdfInfoTrapped = "\u0046\u0061\u006cs\u0065"
)

// ColorFromFloats returns a new PdfColorDevice based on the input slice of
// color components. The slice should contain four elements representing the
// cyan, magenta, yellow and key components of the color. The values of the
// elements should be between 0 and 1.
func (_fcga *PdfColorspaceDeviceCMYK) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 4 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_agcf := vals[0]
	if _agcf < 0.0 || _agcf > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _agcf)
		return nil, ErrColorOutOfRange
	}
	_cgaf := vals[1]
	if _cgaf < 0.0 || _cgaf > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _cgaf)
		return nil, ErrColorOutOfRange
	}
	_fbdab := vals[2]
	if _fbdab < 0.0 || _fbdab > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fbdab)
		return nil, ErrColorOutOfRange
	}
	_cbda := vals[3]
	if _cbda < 0.0 || _cbda > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _cbda)
		return nil, ErrColorOutOfRange
	}
	_ddec := NewPdfColorDeviceCMYK(_agcf, _cgaf, _fbdab, _cbda)
	return _ddec, nil
}

// ToPdfObject returns the PDF representation of the pattern.
func (_fdbg *PdfPattern) ToPdfObject() _ebb.PdfObject {
	_aeef := _fdbg.getDict()
	_aeef.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_aeef.Set("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065", _ebb.MakeInteger(_fdbg.PatternType))
	return _fdbg._dcddc
}

const (
	XObjectTypeUndefined XObjectType = iota
	XObjectTypeImage
	XObjectTypeForm
	XObjectTypePS
	XObjectTypeUnknown
)

func (_dade *PdfReader) newPdfAnnotationCircleFromDict(_aef *_ebb.PdfObjectDictionary) (*PdfAnnotationCircle, error) {
	_bfde := PdfAnnotationCircle{}
	_bbac, _bed := _dade.newPdfAnnotationMarkupFromDict(_aef)
	if _bed != nil {
		return nil, _bed
	}
	_bfde.PdfAnnotationMarkup = _bbac
	_bfde.BS = _aef.Get("\u0042\u0053")
	_bfde.IC = _aef.Get("\u0049\u0043")
	_bfde.BE = _aef.Get("\u0042\u0045")
	_bfde.RD = _aef.Get("\u0052\u0044")
	return &_bfde, nil
}

// FieldFlag represents form field flags. Some of the flags can apply to all types of fields whereas other
// flags are specific.
type FieldFlag uint32

func (_ggcgg *PdfWriter) writeObjects() {
	_eg.Log.Trace("\u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0025d\u0020\u006f\u0062\u006a", len(_ggcgg._ebdgg))
	_ggcgg._bedfc = make(map[int]crossReference)
	_ggcgg._bedfc[0] = crossReference{Type: 0, ObjectNumber: 0, Generation: 0xFFFF}
	if _ggcgg._efdega.ObjectMap != nil {
		for _acdgg, _bcee := range _ggcgg._efdega.ObjectMap {
			if _acdgg == 0 {
				continue
			}
			if _bcee.XType == _ebb.XrefTypeObjectStream {
				_cddb := crossReference{Type: 2, ObjectNumber: _bcee.OsObjNumber, Index: _bcee.OsObjIndex}
				_ggcgg._bedfc[_acdgg] = _cddb
			}
			if _bcee.XType == _ebb.XrefTypeTableEntry {
				_bcfef := crossReference{Type: 1, ObjectNumber: _bcee.ObjectNumber, Offset: _bcee.Offset}
				_ggcgg._bedfc[_acdgg] = _bcfef
			}
		}
	}
}

// NewPdfAnnotationProjection returns a new projection annotation.
func NewPdfAnnotationProjection() *PdfAnnotationProjection {
	_acgg := NewPdfAnnotation()
	_gcc := &PdfAnnotationProjection{}
	_gcc.PdfAnnotation = _acgg
	_gcc.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_acgg.SetContext(_gcc)
	return _gcc
}

// GetContainingPdfObject returns the container of the outline (indirect object).
func (_eaadg *PdfOutline) GetContainingPdfObject() _ebb.PdfObject { return _eaadg._egee }

// A returns the value of the A component of the color.
func (_dccf *PdfColorCalRGB) A() float64 { return _dccf[0] }

// PdfFunction interface represents the common methods of a function in PDF.
type PdfFunction interface {
	Evaluate([]float64) ([]float64, error)
	ToPdfObject() _ebb.PdfObject
}

func (_fgdb *pdfFontSimple) getFontEncoding() (_agece string, _afbd map[_da.CharCode]_da.GlyphName, _dcfa error) {
	_agece = "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"
	if _gffdc, _gbec := _cbcfg[_fgdb._fdacg]; _gbec {
		_agece = _gffdc
	} else if _fgdb.fontFlags()&_dffee != 0 {
		for _cegf, _fefeb := range _cbcfg {
			if _ee.Contains(_fgdb._fdacg, _cegf) {
				_agece = _fefeb
				break
			}
		}
	}
	if _fgdb.Encoding == nil {
		return _agece, nil, nil
	}
	switch _egedae := _fgdb.Encoding.(type) {
	case *_ebb.PdfObjectName:
		return string(*_egedae), nil, nil
	case *_ebb.PdfObjectDictionary:
		_cdaee, _acfeg := _ebb.GetName(_egedae.Get("\u0042\u0061\u0073e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
		if _acfeg {
			_agece = _cdaee.String()
		}
		if _gebea := _egedae.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"); _gebea != nil {
			_eeefg, _bebg := _ebb.GetArray(_gebea)
			if !_bebg {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042a\u0064\u0020\u0066on\u0074\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u003d\u0025\u002b\u0076\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073=\u0025\u0054", _egedae, _egedae.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"))
				return "", nil, _ebb.ErrTypeError
			}
			_afbd, _dcfa = _da.FromFontDifferences(_eeefg)
		}
		return _agece, _afbd, _dcfa
	default:
		_eg.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0072\u0020\u0064\u0069\u0063t\u0020\u0028\u0025\u0054\u0029\u0020\u0025\u0073", _fgdb.Encoding, _fgdb.Encoding)
		return "", nil, _ebb.ErrTypeError
	}
}

// ToPdfObject implements interface PdfModel.
func (_bgcb *PdfFilespec) ToPdfObject() _ebb.PdfObject {
	_edfd := _bgcb.getDict()
	_edfd.Clear()
	_edfd.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063"))
	_edfd.SetIfNotNil("\u0046\u0053", _bgcb.FS)
	_edfd.SetIfNotNil("\u0046", _bgcb.F)
	_edfd.SetIfNotNil("\u0055\u0046", _bgcb.UF)
	_edfd.SetIfNotNil("\u0044\u004f\u0053", _bgcb.DOS)
	_edfd.SetIfNotNil("\u004d\u0061\u0063", _bgcb.Mac)
	_edfd.SetIfNotNil("\u0055\u006e\u0069\u0078", _bgcb.Unix)
	_edfd.SetIfNotNil("\u0049\u0044", _bgcb.ID)
	_edfd.SetIfNotNil("\u0056", _bgcb.V)
	_edfd.SetIfNotNil("\u0045\u0046", _bgcb.EF)
	_edfd.SetIfNotNil("\u0052\u0046", _bgcb.RF)
	_edfd.SetIfNotNil("\u0044\u0065\u0073\u0063", _bgcb.Desc)
	_edfd.SetIfNotNil("\u0043\u0049", _bgcb.CI)
	return _bgcb._gcge
}

// ColorToRGB converts a ICCBased color to an RGB color.
func (_agf *PdfColorspaceICCBased) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _agf.Alternate == nil {
		_eg.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _agf.N == 1 {
			_eg.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_aca := NewPdfColorspaceDeviceGray()
			return _aca.ColorToRGB(color)
		} else if _agf.N == 3 {
			_eg.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return color, nil
		} else if _agf.N == 4 {
			_eg.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_bedd := NewPdfColorspaceDeviceCMYK()
			return _bedd.ColorToRGB(color)
		} else {
			return nil, _gf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_eg.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _agf)
	return _agf.Alternate.ColorToRGB(color)
}

const (
	ButtonTypeCheckbox ButtonType = iota
	ButtonTypePush     ButtonType = iota
	ButtonTypeRadio    ButtonType = iota
)

func (_cgec *PdfColorspaceCalRGB) String() string { return "\u0043\u0061\u006c\u0052\u0047\u0042" }

// GetRotate gets the inheritable rotate value, either from the page
// or a higher up page/pages struct.
func (_eagc *PdfPage) GetRotate() (int64, error) {
	if _eagc.Rotate != nil {
		return *_eagc.Rotate, nil
	}
	_cgbg := _eagc.Parent
	for _cgbg != nil {
		_dfded, _aefg := _ebb.GetDict(_cgbg)
		if !_aefg {
			return 0, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _abagc := _dfded.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _abagc != nil {
			_cdbed, _dgec := _ebb.GetInt(_abagc)
			if !_dgec {
				return 0, _gf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0074a\u0074\u0065\u0020\u0076al\u0075\u0065")
			}
			if _cdbed != nil {
				return int64(*_cdbed), nil
			}
			return 0, _gf.New("\u0072\u006f\u0074\u0061te\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		_cgbg = _dfded.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return 0, _gf.New("\u0072o\u0074a\u0074\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_bfac *PdfColorspaceSpecialPattern) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _bfac.UnderlyingCS == nil {
		return nil, _gf.New("u\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0043\u0053\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
	}
	return _bfac.UnderlyingCS.ColorFromFloats(vals)
}

// Add appends a top level outline item to the outline.
func (_aaafc *Outline) Add(item *OutlineItem) { _aaafc.Entries = append(_aaafc.Entries, item) }

// DecodeArray returns an empty slice as there are no components associated with pattern colorspace.
func (_gbefd *PdfColorspaceSpecialPattern) DecodeArray() []float64 { return []float64{} }

// PdfFontDescriptor specifies metrics and other attributes of a font and can refer to a FontFile
// for embedded fonts.
// 9.8 Font Descriptors (page 281)
type PdfFontDescriptor struct {
	FontName     _ebb.PdfObject
	FontFamily   _ebb.PdfObject
	FontStretch  _ebb.PdfObject
	FontWeight   _ebb.PdfObject
	Flags        _ebb.PdfObject
	FontBBox     _ebb.PdfObject
	ItalicAngle  _ebb.PdfObject
	Ascent       _ebb.PdfObject
	Descent      _ebb.PdfObject
	Leading      _ebb.PdfObject
	CapHeight    _ebb.PdfObject
	XHeight      _ebb.PdfObject
	StemV        _ebb.PdfObject
	StemH        _ebb.PdfObject
	AvgWidth     _ebb.PdfObject
	MaxWidth     _ebb.PdfObject
	MissingWidth _ebb.PdfObject
	FontFile     _ebb.PdfObject
	FontFile2    _ebb.PdfObject
	FontFile3    _ebb.PdfObject
	CharSet      _ebb.PdfObject
	_gfbge       int
	_gbfgb       float64
	*fontFile
	_aeeb *_bad.TtfType

	// Additional entries for CIDFonts
	Style  _ebb.PdfObject
	Lang   _ebb.PdfObject
	FD     _ebb.PdfObject
	CIDSet _ebb.PdfObject
	_ccfcg *_ebb.PdfIndirectObject
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_ddcb *PdfColorspaceICCBased) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _ddcb.Alternate == nil {
		if _ddcb.N == 1 {
			_fgeb := NewPdfColorspaceDeviceGray()
			return _fgeb.ColorFromFloats(vals)
		} else if _ddcb.N == 3 {
			_gged := NewPdfColorspaceDeviceRGB()
			return _gged.ColorFromFloats(vals)
		} else if _ddcb.N == 4 {
			_cbed := NewPdfColorspaceDeviceCMYK()
			return _cbed.ColorFromFloats(vals)
		} else {
			return nil, _gf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _ddcb.Alternate.ColorFromFloats(vals)
}

// CustomKeys returns all custom info keys as list.
func (_badba *PdfInfo) CustomKeys() []string {
	if _badba._gcgf == nil {
		return nil
	}
	_agfcc := make([]string, len(_badba._gcgf.Keys()))
	for _, _dfaf := range _badba._gcgf.Keys() {
		_agfcc = append(_agfcc, _dfaf.String())
	}
	return _agfcc
}

// PdfAnnotationStrikeOut represents StrikeOut annotations.
// (Section 12.5.6.10).
type PdfAnnotationStrikeOut struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _ebb.PdfObject
}

// B returns the value of the B component of the color.
func (_bbd *PdfColorCalRGB) B() float64 { return _bbd[1] }

// ToPdfObject returns the PDF representation of the colorspace.
func (_cdfge *PdfPageResourcesColorspaces) ToPdfObject() _ebb.PdfObject {
	_dcebg := _ebb.MakeDict()
	for _, _eeac := range _cdfge.Names {
		_dcebg.Set(_ebb.PdfObjectName(_eeac), _cdfge.Colorspaces[_eeac].ToPdfObject())
	}
	if _cdfge._ddffd != nil {
		_cdfge._ddffd.PdfObject = _dcebg
		return _cdfge._ddffd
	}
	return _dcebg
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
	_faeb           *_ebb.PdfObjectDictionary
}

// GetNumComponents returns the number of color components of the underlying
// colorspace device.
func (_daac *PdfColorspaceSpecialPattern) GetNumComponents() int {
	return _daac.UnderlyingCS.GetNumComponents()
}

// PdfActionRendition represents a Rendition action.
type PdfActionRendition struct {
	*PdfAction
	R  _ebb.PdfObject
	AN _ebb.PdfObject
	OP _ebb.PdfObject
	JS _ebb.PdfObject
}

// G returns the value of the green component of the color.
func (_eabd *PdfColorDeviceRGB) G() float64 { return _eabd[1] }

// ToPdfObject implements interface PdfModel.
func (_egd *PdfActionRendition) ToPdfObject() _ebb.PdfObject {
	_egd.PdfAction.ToPdfObject()
	_cgdf := _egd._abe
	_ddef := _cgdf.PdfObject.(*_ebb.PdfObjectDictionary)
	_ddef.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeRendition)))
	_ddef.SetIfNotNil("\u0052", _egd.R)
	_ddef.SetIfNotNil("\u0041\u004e", _egd.AN)
	_ddef.SetIfNotNil("\u004f\u0050", _egd.OP)
	_ddef.SetIfNotNil("\u004a\u0053", _egd.JS)
	return _cgdf
}

// NewPdfAnnotationPolygon returns a new polygon annotation.
func NewPdfAnnotationPolygon() *PdfAnnotationPolygon {
	_cegg := NewPdfAnnotation()
	_dfg := &PdfAnnotationPolygon{}
	_dfg.PdfAnnotation = _cegg
	_dfg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cegg.SetContext(_dfg)
	return _dfg
}

const (
	_gfega = 0x00001
	_adfbd = 0x00002
	_dffee = 0x00004
	_dgfgd = 0x00008
	_cbfd  = 0x00020
	_ccgbe = 0x00040
	_fece  = 0x10000
	_ddge  = 0x20000
	_decd  = 0x40000
)

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// component PDF objects.
func (_dgae *PdfColorspaceICCBased) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if _dgae.Alternate == nil {
		if _dgae.N == 1 {
			_dfag := NewPdfColorspaceDeviceGray()
			return _dfag.ColorFromPdfObjects(objects)
		} else if _dgae.N == 3 {
			_bbbe := NewPdfColorspaceDeviceRGB()
			return _bbbe.ColorFromPdfObjects(objects)
		} else if _dgae.N == 4 {
			_fbae := NewPdfColorspaceDeviceCMYK()
			return _fbae.ColorFromPdfObjects(objects)
		} else {
			return nil, _gf.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _dgae.Alternate.ColorFromPdfObjects(objects)
}
func (_gccfa Image) getBase() _dg.ImageBase {
	return _dg.NewImageBase(int(_gccfa.Width), int(_gccfa.Height), int(_gccfa.BitsPerComponent), _gccfa.ColorComponents, _gccfa.Data, _gccfa._dagcb, _gccfa._dgcea)
}

// GetCustomInfo returns a custom info value for the specified name.
func (_deff *PdfInfo) GetCustomInfo(name string) *_ebb.PdfObjectString {
	var _cfaac *_ebb.PdfObjectString
	if _deff._gcgf == nil {
		return _cfaac
	}
	if _dfffg, _gafc := _deff._gcgf.Get(*_ebb.MakeName(name)).(*_ebb.PdfObjectString); _gafc {
		_cfaac = _dfffg
	}
	return _cfaac
}

// PdfActionType represents an action type in PDF (section 12.6.4 p. 417).
type PdfActionType string

func _fadcd() *modelManager {
	_gbed := modelManager{}
	_gbed._bgfdb = map[PdfModel]_ebb.PdfObject{}
	_gbed._aabcbdd = map[_ebb.PdfObject]PdfModel{}
	return &_gbed
}

// ToPdfObject implements interface PdfModel.
func (_eec *PdfActionLaunch) ToPdfObject() _ebb.PdfObject {
	_eec.PdfAction.ToPdfObject()
	_eeag := _eec._abe
	_fbd := _eeag.PdfObject.(*_ebb.PdfObjectDictionary)
	_fbd.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeLaunch)))
	if _eec.F != nil {
		_fbd.Set("\u0046", _eec.F.ToPdfObject())
	}
	_fbd.SetIfNotNil("\u0057\u0069\u006e", _eec.Win)
	_fbd.SetIfNotNil("\u004d\u0061\u0063", _eec.Mac)
	_fbd.SetIfNotNil("\u0055\u006e\u0069\u0078", _eec.Unix)
	_fbd.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _eec.NewWindow)
	return _eeag
}
func (_eefd *PdfWriter) getPdfVersion() string {
	return _bg.Sprintf("\u0025\u0064\u002e%\u0064", _eefd._efcge.Major, _eefd._efcge.Minor)
}
func _adcga(_bbegd *_ebb.PdfObjectDictionary) (*PdfShadingType2, error) {
	_fffe := PdfShadingType2{}
	_feeca := _bbegd.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _feeca == nil {
		_eg.Log.Debug("R\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u003a\u0020\u0020\u0043\u006f\u006f\u0072d\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_daaca, _ecbcg := _feeca.(*_ebb.PdfObjectArray)
	if !_ecbcg {
		_eg.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _feeca)
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _daaca.Len() != 4 {
		_eg.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0034\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _daaca.Len())
		return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	}
	_fffe.Coords = _daaca
	if _dbaefe := _bbegd.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _dbaefe != nil {
		_dbaefe = _ebb.TraceToDirectObject(_dbaefe)
		_ecdge, _bffed := _dbaefe.(*_ebb.PdfObjectArray)
		if !_bffed {
			_eg.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _dbaefe)
			return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fffe.Domain = _ecdge
	}
	_feeca = _bbegd.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _feeca == nil {
		_eg.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_fffe.Function = []PdfFunction{}
	if _afagg, _eebcbf := _feeca.(*_ebb.PdfObjectArray); _eebcbf {
		for _, _bgcbg := range _afagg.Elements() {
			_bddce, _ceaee := _aagg(_bgcbg)
			if _ceaee != nil {
				_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _ceaee)
				return nil, _ceaee
			}
			_fffe.Function = append(_fffe.Function, _bddce)
		}
	} else {
		_ebdfb, _ffbd := _aagg(_feeca)
		if _ffbd != nil {
			_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _ffbd)
			return nil, _ffbd
		}
		_fffe.Function = append(_fffe.Function, _ebdfb)
	}
	if _gceef := _bbegd.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _gceef != nil {
		_gceef = _ebb.TraceToDirectObject(_gceef)
		_adafb, _egddg := _gceef.(*_ebb.PdfObjectArray)
		if !_egddg {
			_eg.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _gceef)
			return nil, _ebb.ErrTypeError
		}
		if _adafb.Len() != 2 {
			_eg.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _adafb.Len())
			return nil, ErrInvalidAttribute
		}
		_fffe.Extend = _adafb
	}
	return &_fffe, nil
}

// ToPdfObject implements model.PdfModel interface.
func (_bfcfe *PdfOutputIntent) ToPdfObject() _ebb.PdfObject {
	if _bfcfe._faeb == nil {
		_bfcfe._faeb = _ebb.MakeDict()
	}
	_fgfb := _bfcfe._faeb
	if _bfcfe.Type != "" {
		_fgfb.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName(_bfcfe.Type))
	}
	_fgfb.Set("\u0053", _ebb.MakeName(_bfcfe.S.String()))
	if _bfcfe.OutputCondition != "" {
		_fgfb.Set("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e", _ebb.MakeString(_bfcfe.OutputCondition))
	}
	_fgfb.Set("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r", _ebb.MakeString(_bfcfe.OutputConditionIdentifier))
	_fgfb.Set("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065", _ebb.MakeString(_bfcfe.RegistryName))
	if _bfcfe.Info != "" {
		_fgfb.Set("\u0049\u006e\u0066\u006f", _ebb.MakeString(_bfcfe.Info))
	}
	if len(_bfcfe.DestOutputProfile) != 0 {
		_bacag, _fdcdc := _ebb.MakeStream(_bfcfe.DestOutputProfile, _ebb.NewFlateEncoder())
		if _fdcdc != nil {
			_eg.Log.Error("\u004d\u0061\u006b\u0065\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0044\u0065s\u0074\u004f\u0075\u0074\u0070\u0075t\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fdcdc)
		}
		_bacag.PdfObjectDictionary.Set("\u004e", _ebb.MakeInteger(int64(_bfcfe.ColorComponents)))
		_eabdg := make([]float64, _bfcfe.ColorComponents*2)
		for _bgdaa := 0; _bgdaa < _bfcfe.ColorComponents*2; _bgdaa++ {
			_ggfcb := 0.0
			if _bgdaa%2 != 0 {
				_ggfcb = 1.0
			}
			_eabdg[_bgdaa] = _ggfcb
		}
		_bacag.PdfObjectDictionary.Set("\u0052\u0061\u006eg\u0065", _ebb.MakeArrayFromFloats(_eabdg))
		_fgfb.Set("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065", _bacag)
	}
	return _fgfb
}

// HasFontByName checks if has font resource by name.
func (_cdgf *PdfPage) HasFontByName(name _ebb.PdfObjectName) bool {
	_dgge, _daecb := _cdgf.Resources.Font.(*_ebb.PdfObjectDictionary)
	if !_daecb {
		return false
	}
	if _eeeec := _dgge.Get(name); _eeeec != nil {
		return true
	}
	return false
}

// PdfAnnotation represents an annotation in PDF (section 12.5 p. 389).
type PdfAnnotation struct {
	_efd         PdfModel
	Rect         _ebb.PdfObject
	Contents     _ebb.PdfObject
	P            _ebb.PdfObject
	NM           _ebb.PdfObject
	M            _ebb.PdfObject
	F            _ebb.PdfObject
	AP           _ebb.PdfObject
	AS           _ebb.PdfObject
	Border       _ebb.PdfObject
	C            _ebb.PdfObject
	StructParent _ebb.PdfObject
	OC           _ebb.PdfObject
	_bdcd        *_ebb.PdfIndirectObject
}

func (_cggf *PdfAppender) updateObjectsDeep(_ged _ebb.PdfObject, _badbe map[_ebb.PdfObject]struct{}) {
	if _badbe == nil {
		_badbe = map[_ebb.PdfObject]struct{}{}
	}
	if _, _gfdb := _badbe[_ged]; _gfdb || _ged == nil {
		return
	}
	_badbe[_ged] = struct{}{}
	_adcag := _ebb.ResolveReferencesDeep(_ged, _cggf._agb)
	if _adcag != nil {
		_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _adcag)
	}
	switch _dfcf := _ged.(type) {
	case *_ebb.PdfIndirectObject:
		switch {
		case _dfcf.GetParser() == _cggf._acfe._cafdf:
			return
		case _dfcf.GetParser() == _cggf.Reader._cafdf:
			_gbgg, _ := _cggf._acfe.GetIndirectObjectByNumber(int(_dfcf.ObjectNumber))
			_adfc, _feec := _gbgg.(*_ebb.PdfIndirectObject)
			if _feec && _adfc != nil {
				if _adfc.PdfObject != _dfcf.PdfObject && _adfc.PdfObject.WriteString() != _dfcf.PdfObject.WriteString() {
					_cggf.addNewObject(_ged)
					_cggf._gbfa[_ged] = _dfcf.ObjectNumber
				}
			}
		default:
			_cggf.addNewObject(_ged)
		}
		_cggf.updateObjectsDeep(_dfcf.PdfObject, _badbe)
	case *_ebb.PdfObjectArray:
		for _, _eedb := range _dfcf.Elements() {
			_cggf.updateObjectsDeep(_eedb, _badbe)
		}
	case *_ebb.PdfObjectDictionary:
		for _, _gffe := range _dfcf.Keys() {
			_cggf.updateObjectsDeep(_dfcf.Get(_gffe), _badbe)
		}
	case *_ebb.PdfObjectStreams:
		if _dfcf.GetParser() != _cggf._acfe._cafdf {
			for _, _cdcf := range _dfcf.Elements() {
				_cggf.updateObjectsDeep(_cdcf, _badbe)
			}
		}
	case *_ebb.PdfObjectStream:
		switch {
		case _dfcf.GetParser() == _cggf._acfe._cafdf:
			return
		case _dfcf.GetParser() == _cggf.Reader._cafdf:
			if _cfagc, _cbd := _cggf._acfe._cafdf.LookupByReference(_dfcf.PdfObjectReference); _cbd == nil {
				var _gef bool
				if _geged, _gac := _ebb.GetStream(_cfagc); _gac && _ca.Equal(_geged.Stream, _dfcf.Stream) {
					_gef = true
				}
				if _bbbgg, _feag := _ebb.GetDict(_cfagc); _gef && _feag {
					_gef = _bbbgg.WriteString() == _dfcf.PdfObjectDictionary.WriteString()
				}
				if _gef {
					return
				}
			}
			if _dfcf.ObjectNumber != 0 {
				_cggf._gbfa[_ged] = _dfcf.ObjectNumber
			}
		default:
			if _, _fbgg := _cggf._ddfg[_ged]; !_fbgg {
				_cggf.addNewObject(_ged)
			}
		}
		_cggf.updateObjectsDeep(_dfcf.PdfObjectDictionary, _badbe)
	}
}

// GetOptimizer returns current PDF optimizer.
func (_dccaa *PdfWriter) GetOptimizer() Optimizer { return _dccaa._aadcd }

// SetVersion sets the PDF version of the output file.
func (_febac *PdfWriter) SetVersion(majorVersion, minorVersion int) {
	_febac._efcge.Major = majorVersion
	_febac._efcge.Minor = minorVersion
}

// FontDescriptor returns font's PdfFontDescriptor. This may be a builtin descriptor for standard 14
// fonts but must be an explicit descriptor for other fonts.
func (_fbfc *PdfFont) FontDescriptor() *PdfFontDescriptor {
	if _fbfc.baseFields()._fbbd != nil {
		return _fbfc.baseFields()._fbbd
	}
	if _dfggb := _fbfc._ebcad.getFontDescriptor(); _dfggb != nil {
		return _dfggb
	}
	_eg.Log.Error("\u0041\u006cl \u0066\u006f\u006et\u0073\u0020\u0068\u0061ve \u0061 D\u0065\u0073\u0063\u0072\u0069\u0070\u0074or\u002e\u0020\u0066\u006f\u006e\u0074\u003d%\u0073", _fbfc)
	return nil
}
func (_baef *PdfColorspaceDeviceRGB) String() string {
	return "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
}

// GetContext returns the context of the outline tree node, which is either a
// *PdfOutline or a *PdfOutlineItem. The method returns nil for uninitialized
// tree nodes.
func (_bbgeb *PdfOutlineTreeNode) GetContext() PdfModel {
	if _adebf, _ebgba := _bbgeb._geeee.(*PdfOutline); _ebgba {
		return _adebf
	}
	if _dfgcb, _gdbe := _bbgeb._geeee.(*PdfOutlineItem); _gdbe {
		return _dfgcb
	}
	_eg.Log.Debug("\u0045\u0052RO\u0052\u0020\u0049n\u0076\u0061\u006c\u0069d o\u0075tl\u0069\u006e\u0065\u0020\u0074\u0072\u0065e \u006e\u006f\u0064\u0065\u0020\u0069\u0074e\u006d")
	return nil
}
func _afacb(_eaad *_dg.ImageBase) (_edacc Image) {
	_edacc.Width = int64(_eaad.Width)
	_edacc.Height = int64(_eaad.Height)
	_edacc.BitsPerComponent = int64(_eaad.BitsPerComponent)
	_edacc.ColorComponents = _eaad.ColorComponents
	_edacc.Data = _eaad.Data
	_edacc._dgcea = _eaad.Decode
	_edacc._dagcb = _eaad.Alpha
	return _edacc
}

// DecodeArray returns the component range values for the DeviceN colorspace.
// [0 1.0 0 1.0 ...] for each color component.
func (_affe *PdfColorspaceDeviceN) DecodeArray() []float64 {
	var _bgfad []float64
	for _gdbc := 0; _gdbc < _affe.GetNumComponents(); _gdbc++ {
		_bgfad = append(_bgfad, 0.0, 1.0)
	}
	return _bgfad
}

// ToPdfObject implements interface PdfModel.
func (_edab *PdfAnnotationCaret) ToPdfObject() _ebb.PdfObject {
	_edab.PdfAnnotation.ToPdfObject()
	_befg := _edab._bdcd
	_gbee := _befg.PdfObject.(*_ebb.PdfObjectDictionary)
	_edab.PdfAnnotationMarkup.appendToPdfDictionary(_gbee)
	_gbee.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0043\u0061\u0072e\u0074"))
	_gbee.SetIfNotNil("\u0052\u0044", _edab.RD)
	_gbee.SetIfNotNil("\u0053\u0079", _edab.Sy)
	return _befg
}

// PageFromIndirectObject returns the PdfPage and page number for a given indirect object.
func (_eacbb *PdfReader) PageFromIndirectObject(ind *_ebb.PdfIndirectObject) (*PdfPage, int, error) {
	if len(_eacbb.PageList) != len(_eacbb._faebb) {
		return nil, 0, _gf.New("\u0070\u0061\u0067\u0065\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	for _bged, _fbaef := range _eacbb._faebb {
		if _fbaef == ind {
			return _eacbb.PageList[_bged], _bged + 1, nil
		}
	}
	return nil, 0, _gf.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// ToPdfObject implements interface PdfModel.
func (_fff *PdfAnnotationCircle) ToPdfObject() _ebb.PdfObject {
	_fff.PdfAnnotation.ToPdfObject()
	_bdcb := _fff._bdcd
	_bcec := _bdcb.PdfObject.(*_ebb.PdfObjectDictionary)
	_fff.PdfAnnotationMarkup.appendToPdfDictionary(_bcec)
	_bcec.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0043\u0069\u0072\u0063\u006c\u0065"))
	_bcec.SetIfNotNil("\u0042\u0053", _fff.BS)
	_bcec.SetIfNotNil("\u0049\u0043", _fff.IC)
	_bcec.SetIfNotNil("\u0042\u0045", _fff.BE)
	_bcec.SetIfNotNil("\u0052\u0044", _fff.RD)
	return _bdcb
}

// AddExtGState add External Graphics State (GState). The gsDict can be specified
// either directly as a dictionary or an indirect object containing a dictionary.
func (_beca *PdfPageResources) AddExtGState(gsName _ebb.PdfObjectName, gsDict _ebb.PdfObject) error {
	if _beca.ExtGState == nil {
		_beca.ExtGState = _ebb.MakeDict()
	}
	_egfad := _beca.ExtGState
	_abbce, _baabb := _ebb.TraceToDirectObject(_egfad).(*_ebb.PdfObjectDictionary)
	if !_baabb {
		_eg.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020e\u0072r\u006f\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u002f\u0025\u0054\u0029", _egfad, _ebb.TraceToDirectObject(_egfad))
		return _ebb.ErrTypeError
	}
	_abbce.Set(gsName, gsDict)
	return nil
}

// ColorToRGB converts a color in Separation colorspace to RGB colorspace.
func (_eegce *PdfColorspaceSpecialSeparation) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _eegce.AlternateSpace == nil {
		return nil, _gf.New("\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020c\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	return _eegce.AlternateSpace.ColorToRGB(color)
}
func _cgcga(_ggcdf *[]*PdfField, _ggebg FieldFilterFunc, _cacca bool) []*PdfField {
	if _ggcdf == nil {
		return nil
	}
	_dbdff := *_ggcdf
	if len(*_ggcdf) == 0 {
		return nil
	}
	_ffdfb := _dbdff[:0]
	if _ggebg == nil {
		_ggebg = func(*PdfField) bool { return true }
	}
	var _eedfc []*PdfField
	for _, _bagg := range _dbdff {
		_ggedc := _ggebg(_bagg)
		if _ggedc {
			_eedfc = append(_eedfc, _bagg)
			if len(_bagg.Kids) > 0 {
				_eedfc = append(_eedfc, _cgcga(&_bagg.Kids, _ggebg, _cacca)...)
			}
		}
		if !_cacca || !_ggedc || len(_bagg.Kids) > 0 {
			_ffdfb = append(_ffdfb, _bagg)
		}
	}
	*_ggcdf = _ffdfb
	return _eedfc
}

// XObjectType represents the type of an XObject.
type XObjectType int

// NewPdfFontFromPdfObject loads a PdfFont from the dictionary `fontObj`.  If there is a problem an
// error is returned.
func NewPdfFontFromPdfObject(fontObj _ebb.PdfObject) (*PdfFont, error) { return _ddacd(fontObj, true) }

// GetRuneMetrics returns the char metrics for a rune.
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
//                       well give them 0 width. There is no need for the bool return.
func (_bdafb *PdfFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_acdab := _bdafb.actualFont()
	if _acdab == nil {
		_eg.Log.Debug("ER\u0052\u004fR\u003a\u0020\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u004e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020f\u006fr\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079p\u0065=\u0025\u0023T", _bdafb._ebcad)
		return _bad.CharMetrics{}, false
	}
	if _ccfcc, _eaca := _acdab.GetRuneMetrics(r); _eaca {
		return _ccfcc, true
	}
	if _dbacdg, _bccc := _bdafb.GetFontDescriptor(); _bccc == nil && _dbacdg != nil {
		return _bad.CharMetrics{Wx: _dbacdg._gbfgb}, true
	}
	_eg.Log.Debug("\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061r\u004d\u0065\u0074\u0072i\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _bdafb)
	return _bad.CharMetrics{}, false
}

// PdfOutputIntentType is the subtype of the given PdfOutputIntent.
type PdfOutputIntentType int

func (_dfdg *PdfReader) newPdfAnnotationPrinterMarkFromDict(_dbcd *_ebb.PdfObjectDictionary) (*PdfAnnotationPrinterMark, error) {
	_bbag := PdfAnnotationPrinterMark{}
	_bbag.MN = _dbcd.Get("\u004d\u004e")
	return &_bbag, nil
}

// FullName returns the full name of the field as in rootname.parentname.partialname.
func (_bfdfg *PdfField) FullName() (string, error) {
	var _dceg _ca.Buffer
	_cafga := []string{}
	if _bfdfg.T != nil {
		_cafga = append(_cafga, _bfdfg.T.Decoded())
	}
	_dbacd := map[*PdfField]bool{}
	_dbacd[_bfdfg] = true
	_gefg := _bfdfg.Parent
	for _gefg != nil {
		if _, _ccfbg := _dbacd[_gefg]; _ccfbg {
			return _dceg.String(), _gf.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		if _gefg.T == nil {
			return _dceg.String(), _gf.New("\u0066\u0069el\u0064\u0020\u0070a\u0072\u0074\u0069\u0061l n\u0061me\u0020\u0028\u0054\u0029\u0020\u006e\u006ft \u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064")
		}
		_cafga = append(_cafga, _gefg.T.Decoded())
		_dbacd[_gefg] = true
		_gefg = _gefg.Parent
	}
	for _abdb := len(_cafga) - 1; _abdb >= 0; _abdb-- {
		_dceg.WriteString(_cafga[_abdb])
		if _abdb > 0 {
			_dceg.WriteString("\u002e")
		}
	}
	return _dceg.String(), nil
}

// Register registers (caches) a model to primitive object relationship.
func (_ecbd *modelManager) Register(primitive _ebb.PdfObject, model PdfModel) {
	_ecbd._bgfdb[model] = primitive
	_ecbd._aabcbdd[primitive] = model
}

// GetVersion gets the document version.
func (_efcaaa *PdfWriter) GetVersion() _ebb.Version { return _efcaaa._efcge }
func (_afbcd *PdfReader) lookupPageByObject(_ecac _ebb.PdfObject) (*PdfPage, error) {
	return nil, _gf.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}
func _gbab(_dgfeg *_ebb.PdfObjectStream) (*PdfFunctionType4, error) {
	_bacf := &PdfFunctionType4{}
	_bacf._gbgd = _dgfeg
	_cdef := _dgfeg.PdfObjectDictionary
	_bcaf, _adaa := _ebb.TraceToDirectObject(_cdef.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_ebb.PdfObjectArray)
	if !_adaa {
		_eg.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _bcaf.Len()%2 != 0 {
		_eg.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _gf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_gefc, _cbce := _bcaf.ToFloat64Array()
	if _cbce != nil {
		return nil, _cbce
	}
	_bacf.Domain = _gefc
	_bcaf, _adaa = _ebb.TraceToDirectObject(_cdef.Get("\u0052\u0061\u006eg\u0065")).(*_ebb.PdfObjectArray)
	if _adaa {
		if _bcaf.Len() < 0 || _bcaf.Len()%2 != 0 {
			return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_dabd, _ggaeb := _bcaf.ToFloat64Array()
		if _ggaeb != nil {
			return nil, _ggaeb
		}
		_bacf.Range = _dabd
	}
	_ddgde, _cbce := _ebb.DecodeStream(_dgfeg)
	if _cbce != nil {
		return nil, _cbce
	}
	_bacf._abbbb = _ddgde
	_dbgaa := _bc.NewPSParser(_ddgde)
	_adcef, _cbce := _dbgaa.Parse()
	if _cbce != nil {
		return nil, _cbce
	}
	_bacf.Program = _adcef
	return _bacf, nil
}

// RepairAcroForm attempts to rebuild the AcroForm fields using the widget
// annotations present in the document pages. Pass nil for the opts parameter
// in order to use the default options.
// NOTE: Currently, the opts parameter is declared in order to enable adding
// future options, but passing nil will always result in the default options
// being used.
func (_bfgfb *PdfReader) RepairAcroForm(opts *AcroFormRepairOptions) error {
	var _cefgf []*PdfField
	_gadd := map[*_ebb.PdfIndirectObject]struct{}{}
	for _, _cbcfc := range _bfgfb.PageList {
		_efgfd, _ceebd := _cbcfc.GetAnnotations()
		if _ceebd != nil {
			return _ceebd
		}
		for _, _gdbbf := range _efgfd {
			var _gacfbf *PdfField
			switch _gaccd := _gdbbf.GetContext().(type) {
			case *PdfAnnotationWidget:
				if _gaccd._gce != nil {
					_gacfbf = _gaccd._gce
					break
				}
				if _cedf, _cbebe := _ebb.GetIndirect(_gaccd.Parent); _cbebe {
					_gacfbf, _ceebd = _bfgfb.newPdfFieldFromIndirectObject(_cedf, nil)
					if _ceebd == nil {
						break
					}
					_eg.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _cedf, _ceebd)
				}
				if _gaccd._bdcd != nil {
					_gacfbf, _ceebd = _bfgfb.newPdfFieldFromIndirectObject(_gaccd._bdcd, nil)
					if _ceebd == nil {
						break
					}
					_eg.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _gaccd._bdcd, _ceebd)
				}
			}
			if _gacfbf == nil {
				continue
			}
			if _, _efdeg := _gadd[_gacfbf._cdfd]; _efdeg {
				continue
			}
			_gadd[_gacfbf._cdfd] = struct{}{}
			_cefgf = append(_cefgf, _gacfbf)
		}
	}
	if len(_cefgf) == 0 {
		return nil
	}
	if _bfgfb.AcroForm == nil {
		_bfgfb.AcroForm = NewPdfAcroForm()
	}
	_bfgfb.AcroForm.Fields = &_cefgf
	return nil
}

// SetPdfCreator sets the Creator attribute of the output PDF.
func SetPdfCreator(creator string) { _daddc.Lock(); defer _daddc.Unlock(); _afedf = creator }

// Permissions specify a permissions dictionary (PDF 1.5).
// (Section 12.8.4, Table 258 - Entries in a permissions dictionary p. 477 in PDF32000_2008).
type Permissions struct {
	DocMDP  *PdfSignature
	_bdgdgf *_ebb.PdfObjectDictionary
}

// PdfActionGoTo3DView represents a GoTo3DView action.
type PdfActionGoTo3DView struct {
	*PdfAction
	TA _ebb.PdfObject
	V  _ebb.PdfObject
}

func (_dcaac *LTV) enable(_acfb, _ecfcf []*_g.Certificate, _beegg string) error {
	_edfg, _adbag, _facbfe := _dcaac.buildCertChain(_acfb, _ecfcf)
	if _facbfe != nil {
		return _facbfe
	}
	_efgd, _facbfe := _dcaac.getCerts(_edfg)
	if _facbfe != nil {
		return _facbfe
	}
	_fefeg, _facbfe := _dcaac.getOCSPs(_edfg, _adbag)
	if _facbfe != nil {
		return _facbfe
	}
	_cdbc, _facbfe := _dcaac.getCRLs(_edfg)
	if _facbfe != nil {
		return _facbfe
	}
	_dagdg := _dcaac._dfdgf
	_acggb, _facbfe := _dagdg.addCerts(_efgd)
	if _facbfe != nil {
		return _facbfe
	}
	_fafac, _facbfe := _dagdg.addOCSPs(_fefeg)
	if _facbfe != nil {
		return _facbfe
	}
	_eaddg, _facbfe := _dagdg.addCRLs(_cdbc)
	if _facbfe != nil {
		return _facbfe
	}
	if _beegg != "" {
		_dagdg.VRI[_beegg] = &VRI{Cert: _acggb, OCSP: _fafac, CRL: _eaddg}
	}
	_dcaac._ggdbg.SetDSS(_dagdg)
	return nil
}

// NewPdfOutlineTree returns an initialized PdfOutline tree.
func NewPdfOutlineTree() *PdfOutline { _ddgg := NewPdfOutline(); _ddgg._geeee = &_ddgg; return _ddgg }
func _fggcd() string                 { _daddc.Lock(); defer _daddc.Unlock(); return _fceef }

// ToPdfObject implements interface PdfModel.
func (_ccgc *PdfBorderStyle) ToPdfObject() _ebb.PdfObject {
	_aaad := _ebb.MakeDict()
	if _ccgc._dgaa != nil {
		if _fbcg, _afed := _ccgc._dgaa.(*_ebb.PdfIndirectObject); _afed {
			_fbcg.PdfObject = _aaad
		}
	}
	_aaad.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0042\u006f\u0072\u0064\u0065\u0072"))
	if _ccgc.W != nil {
		_aaad.Set("\u0057", _ebb.MakeFloat(*_ccgc.W))
	}
	if _ccgc.S != nil {
		_aaad.Set("\u0053", _ebb.MakeName(_ccgc.S.GetPdfName()))
	}
	if _ccgc.D != nil {
		_aaad.Set("\u0044", _ebb.MakeArrayFromIntegers(*_ccgc.D))
	}
	if _ccgc._dgaa != nil {
		return _ccgc._dgaa
	}
	return _aaad
}

// ToPdfObject implements interface PdfModel.
func (_aabc *PdfActionSubmitForm) ToPdfObject() _ebb.PdfObject {
	_aabc.PdfAction.ToPdfObject()
	_deg := _aabc._abe
	_dbe := _deg.PdfObject.(*_ebb.PdfObjectDictionary)
	_dbe.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeSubmitForm)))
	if _aabc.F != nil {
		_dbe.Set("\u0046", _aabc.F.ToPdfObject())
	}
	_dbe.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _aabc.Fields)
	_dbe.SetIfNotNil("\u0046\u006c\u0061g\u0073", _aabc.Flags)
	return _deg
}

// PdfPageResources is a Page resources model.
// Implements PdfModel.
type PdfPageResources struct {
	ExtGState  _ebb.PdfObject
	ColorSpace _ebb.PdfObject
	Pattern    _ebb.PdfObject
	Shading    _ebb.PdfObject
	XObject    _ebb.PdfObject
	Font       _ebb.PdfObject
	ProcSet    _ebb.PdfObject
	Properties _ebb.PdfObject
	_efbed     *_ebb.PdfObjectDictionary
	_aaee      *PdfPageResourcesColorspaces
}

// SignatureHandler interface defines the common functionality for PDF signature handlers, which
// need to be capable of validating digital signatures and signing PDF documents.
type SignatureHandler interface {

	// IsApplicable checks if a given signature dictionary `sig` is applicable for the signature handler.
	// For example a signature of type `adbe.pkcs7.detached` might not fit for a rsa.sha1 handler.
	IsApplicable(_cfgae *PdfSignature) bool

	// Validate validates a PDF signature against a given digest (hash) such as that determined
	// for an input file. Returns validation results.
	Validate(_eeed *PdfSignature, _gdaba Hasher) (SignatureValidationResult, error)

	// InitSignature prepares the signature dictionary for signing. This involves setting all
	// necessary fields, and also allocating sufficient space to the Contents so that the
	// finalized signature can be inserted once the hash is calculated.
	InitSignature(_adfcf *PdfSignature) error

	// NewDigest creates a new digest/hasher based on the signature dictionary and handler.
	NewDigest(_abdae *PdfSignature) (Hasher, error)

	// Sign receives the hash `digest` (for example hash of an input file), and signs based
	// on the signature dictionary `sig` and applies the signature data to the signature
	// dictionary Contents field.
	Sign(_caafc *PdfSignature, _cabe Hasher) error
}

// SetPdfSubject sets the Subject attribute of the output PDF.
func SetPdfSubject(subject string) { _daddc.Lock(); defer _daddc.Unlock(); _feaca = subject }

// SetFlag sets the flag for the field.
func (_bggb *PdfField) SetFlag(flag FieldFlag) { _bggb.Ff = _ebb.MakeInteger(int64(flag)) }

// NewPdfField returns an initialized PdfField.
func NewPdfField() *PdfField { return &PdfField{_cdfd: _ebb.MakeIndirectObject(_ebb.MakeDict())} }

// GetNumComponents returns the number of color components (1 for grayscale).
func (_ddbg *PdfColorDeviceGray) GetNumComponents() int { return 1 }

// SetFilter sets compression filter. Decodes with current filter sets and
// encodes the data with the new filter.
func (_gddag *XObjectImage) SetFilter(encoder _ebb.StreamEncoder) error {
	_dcfdb := _gddag.Stream
	_cefb, _geebd := _gddag.Filter.DecodeBytes(_dcfdb)
	if _geebd != nil {
		return _geebd
	}
	_gddag.Filter = encoder
	encoder.UpdateParams(_gddag.getParamsDict())
	_dcfdb, _geebd = encoder.EncodeBytes(_cefb)
	if _geebd != nil {
		return _geebd
	}
	_gddag.Stream = _dcfdb
	return nil
}

// DecodeArray returns the range of color component values in CalRGB colorspace.
func (_feeg *PdfColorspaceCalRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// ToPdfObject returns a PDF object representation of the outline destination.
func (_fbef OutlineDest) ToPdfObject() _ebb.PdfObject {
	if (_fbef.PageObj == nil && _fbef.Page < 0) || _fbef.Mode == "" {
		return _ebb.MakeNull()
	}
	_caedb := _ebb.MakeArray()
	if _fbef.PageObj != nil {
		_caedb.Append(_fbef.PageObj)
	} else {
		_caedb.Append(_ebb.MakeInteger(_fbef.Page))
	}
	_caedb.Append(_ebb.MakeName(_fbef.Mode))
	switch _fbef.Mode {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		_caedb.Append(_ebb.MakeFloat(_fbef.Y))
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		_caedb.Append(_ebb.MakeFloat(_fbef.X))
	case "\u0058\u0059\u005a":
		_caedb.Append(_ebb.MakeFloat(_fbef.X))
		_caedb.Append(_ebb.MakeFloat(_fbef.Y))
		_caedb.Append(_ebb.MakeFloat(_fbef.Zoom))
	default:
		_caedb.Set(1, _ebb.MakeName("\u0046\u0069\u0074"))
	}
	return _caedb
}

// MergePageWith appends page content to source Pdf file page content.
func (_dcaf *PdfAppender) MergePageWith(pageNum int, page *PdfPage) error {
	_dcea := pageNum - 1
	var _ggge *PdfPage
	for _bfbe, _gcea := range _dcaf._dfbg {
		if _bfbe == _dcea {
			_ggge = _gcea
		}
	}
	if _ggge == nil {
		return _bg.Errorf("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073o\u0075\u0072\u0063\u0065\u0020\u0064o\u0063\u0075\u006de\u006e\u0074", pageNum)
	}
	if _ggge._defbb != nil && _ggge._defbb.GetParser() == _dcaf._acfe._cafdf {
		_ggge = _ggge.Duplicate()
		_dcaf._dfbg[_dcea] = _ggge
	}
	page = page.Duplicate()
	_bfegc(page)
	_caec := _adg(_ggge)
	_abccac := _adg(page)
	_aegd := make(map[_ebb.PdfObjectName]_ebb.PdfObjectName)
	for _dcc := range _abccac {
		if _, _cfbfb := _caec[_dcc]; _cfbfb {
			for _ebfc := 1; true; _ebfc++ {
				_gedc := _ebb.PdfObjectName(string(_dcc) + _aa.Itoa(_ebfc))
				if _, _edc := _caec[_gedc]; !_edc {
					_aegd[_dcc] = _gedc
					break
				}
			}
		}
	}
	_fdge, _edgfb := page.GetContentStreams()
	if _edgfb != nil {
		return _edgfb
	}
	_fbfe, _edgfb := _ggge.GetContentStreams()
	if _edgfb != nil {
		return _edgfb
	}
	for _bac, _aeba := range _fdge {
		for _cfadc, _agdfd := range _aegd {
			_aeba = _ee.Replace(_aeba, "\u002f"+string(_cfadc), "\u002f"+string(_agdfd), -1)
		}
		_fdge[_bac] = _aeba
	}
	_fbfe = append(_fbfe, _fdge...)
	if _fcfe := _ggge.SetContentStreams(_fbfe, _ebb.NewFlateEncoder()); _fcfe != nil {
		return _fcfe
	}
	_ggge._bbfed = append(_ggge._bbfed, page._bbfed...)
	if _ggge.Resources == nil {
		_ggge.Resources = NewPdfPageResources()
	}
	if page.Resources != nil {
		_ggge.Resources.Font = _dcaf.mergeResources(_ggge.Resources.Font, page.Resources.Font, _aegd)
		_ggge.Resources.XObject = _dcaf.mergeResources(_ggge.Resources.XObject, page.Resources.XObject, _aegd)
		_ggge.Resources.Properties = _dcaf.mergeResources(_ggge.Resources.Properties, page.Resources.Properties, _aegd)
		if _ggge.Resources.ProcSet == nil {
			_ggge.Resources.ProcSet = page.Resources.ProcSet
		}
		_ggge.Resources.Shading = _dcaf.mergeResources(_ggge.Resources.Shading, page.Resources.Shading, _aegd)
		_ggge.Resources.ExtGState = _dcaf.mergeResources(_ggge.Resources.ExtGState, page.Resources.ExtGState, _aegd)
	}
	_cgdfe, _edgfb := _ggge.GetMediaBox()
	if _edgfb != nil {
		return _edgfb
	}
	_fggg, _edgfb := page.GetMediaBox()
	if _edgfb != nil {
		return _edgfb
	}
	var _egfc bool
	if _cgdfe.Llx > _fggg.Llx {
		_cgdfe.Llx = _fggg.Llx
		_egfc = true
	}
	if _cgdfe.Lly > _fggg.Lly {
		_cgdfe.Lly = _fggg.Lly
		_egfc = true
	}
	if _cgdfe.Urx < _fggg.Urx {
		_cgdfe.Urx = _fggg.Urx
		_egfc = true
	}
	if _cgdfe.Ury < _fggg.Ury {
		_cgdfe.Ury = _fggg.Ury
		_egfc = true
	}
	if _egfc {
		_ggge.MediaBox = _cgdfe
	}
	return nil
}
func _febaee(_aaaa *_ebb.PdfObjectDictionary, _gacdg *fontCommon) (*pdfFontType3, error) {
	_fcba := _eega(_gacdg)
	_ffdd := _aaaa.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
	if _ffdd == nil {
		_ffdd = _ebb.MakeInteger(0)
	}
	_fcba.FirstChar = _ffdd
	_fdcfb, _cbef := _ebb.GetIntVal(_ffdd)
	if !_cbef {
		_eg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _ffdd)
		return nil, _ebb.ErrTypeError
	}
	_aegf := _da.CharCode(_fdcfb)
	_ffdd = _aaaa.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
	if _ffdd == nil {
		_ffdd = _ebb.MakeInteger(255)
	}
	_fcba.LastChar = _ffdd
	_fdcfb, _cbef = _ebb.GetIntVal(_ffdd)
	if !_cbef {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _ffdd)
		return nil, _ebb.ErrTypeError
	}
	_dacbd := _da.CharCode(_fdcfb)
	_ffdd = _aaaa.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _ffdd != nil {
		_fcba.Resources = _ffdd
	}
	_ffdd = _aaaa.Get("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s")
	if _ffdd == nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0068\u0061\u0072\u0050\u0072\u006f\u0063\u0073\u0020(%\u0076\u0029", _ffdd)
		return nil, _ebb.ErrNotSupported
	}
	_fcba.CharProcs = _ffdd
	_ffdd = _aaaa.Get("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078")
	if _ffdd == nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0046\u006f\u006et\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0028\u0025\u0076\u0029", _ffdd)
		return nil, _ebb.ErrNotSupported
	}
	_fcba.FontMatrix = _ffdd
	_fcba._aggbb = make(map[_da.CharCode]float64)
	_ffdd = _aaaa.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
	if _ffdd != nil {
		_fcba.Widths = _ffdd
		_agad, _cgfgc := _ebb.GetArray(_ffdd)
		if !_cgfgc {
			_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _ffdd)
			return nil, _ebb.ErrTypeError
		}
		_badae, _cgfce := _agad.ToFloat64Array()
		if _cgfce != nil {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _cgfce
		}
		if len(_badae) != int(_dacbd-_aegf+1) {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _dacbd-_aegf+1, len(_badae))
			return nil, _ebb.ErrRangeError
		}
		_ccac, _cgfgc := _ebb.GetArray(_fcba.FontMatrix)
		if !_cgfgc {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _ccac)
			return nil, _cgfce
		}
		_feegb, _cgfce := _ccac.ToFloat64Array()
		if _cgfce != nil {
			_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020c\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0046o\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _cgfce
		}
		_ebdbgg := _fef.NewMatrix(_feegb[0], _feegb[1], _feegb[2], _feegb[3], _feegb[4], _feegb[5])
		for _efdg, _gfcec := range _badae {
			_egccd, _ := _ebdbgg.Transform(_gfcec, _gfcec)
			_fcba._aggbb[_aegf+_da.CharCode(_efdg)] = _egccd
		}
	}
	_fcba.Encoding = _ebb.TraceToDirectObject(_aaaa.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	_cbcdg := _aaaa.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _cbcdg != nil {
		_fcba._baag = _ebb.TraceToDirectObject(_cbcdg)
		_ccdd, _gbfeb := _feef(_fcba._baag, &_fcba.fontCommon)
		if _gbfeb != nil {
			return nil, _gbfeb
		}
		_fcba._dcdd = _ccdd
	}
	if _gacg := _fcba._dcdd; _gacg != nil {
		_fcba._abcba = _da.NewCMapEncoder("", nil, _gacg)
	} else {
		_fcba._abcba = _da.NewPdfDocEncoder()
	}
	return _fcba, nil
}

// ToPdfObject returns the PDF representation of the outline tree node.
func (_ddfea *PdfOutlineTreeNode) ToPdfObject() _ebb.PdfObject {
	return _ddfea.GetContext().ToPdfObject()
}
func _dddgd(_ffbb string) (map[_da.CharCode]_da.GlyphName, error) {
	_ccacb := _ee.Split(_ffbb, "\u000a")
	_agfa := make(map[_da.CharCode]_da.GlyphName)
	for _, _ddfcd := range _ccacb {
		_bedgg := _bdddd.FindStringSubmatch(_ddfcd)
		if _bedgg == nil {
			continue
		}
		_abace, _eaegc := _bedgg[1], _bedgg[2]
		_cgffce, _ggede := _aa.Atoi(_abace)
		if _ggede != nil {
			_eg.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0042\u0061\u0064\u0020\u0065\u006e\u0063\u006fd\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u002e \u0025\u0071", _ddfcd)
			return nil, _ebb.ErrTypeError
		}
		_agfa[_da.CharCode(_cgffce)] = _da.GlyphName(_eaegc)
	}
	_eg.Log.Trace("g\u0065\u0074\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0073\u003a\u0020\u006b\u0065\u0079V\u0061\u006c\u0075e\u0073=\u0025\u0023\u0076", _agfa)
	return _agfa, nil
}

// GetPreviousRevision returns the previous revision of PdfReader for the Pdf document
func (_dgga *PdfReader) GetPreviousRevision() (*PdfReader, error) {
	if _dgga._cafdf.GetRevisionNumber() == 0 {
		return nil, _gf.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065xi\u0073\u0074")
	}
	if _befbd, _gdgfd := _dgga._cacfc[_dgga]; _gdgfd {
		return _befbd, nil
	}
	_eggdb, _bdaga := _dgga._cafdf.GetPreviousRevisionReadSeeker()
	if _bdaga != nil {
		return nil, _bdaga
	}
	_bgfg, _bdaga := _dcbd(_eggdb, _dgga._edcbc, _dgga._abadec, "\u006do\u0064\u0065\u006c\u003aG\u0065\u0074\u0050\u0072\u0065v\u0069o\u0075s\u0052\u0065\u0076\u0069\u0073\u0069\u006fn")
	if _bdaga != nil {
		return nil, _bdaga
	}
	_dgga._face[_dgga._cafdf.GetRevisionNumber()-1] = _bgfg
	_dgga._cacfc[_dgga] = _bgfg
	_bgfg._cacfc = _dgga._cacfc
	return _bgfg, nil
}
func _bdec() string {
	_ggfd := "\u0051\u0057\u0045\u0052\u0054\u0059\u0055\u0049\u004f\u0050\u0041S\u0044\u0046\u0047\u0048\u004a\u004b\u004c\u005a\u0058\u0043V\u0042\u004e\u004d"
	var _abdaa _ca.Buffer
	for _dgdg := 0; _dgdg < 6; _dgdg++ {
		_abdaa.WriteRune(rune(_ggfd[_ff.Intn(len(_ggfd))]))
	}
	return _abdaa.String()
}

// String returns a string that describes `base`.
func (_afcga fontCommon) String() string {
	return _bg.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0073\u007d", _afcga.coreString())
}

// ToPdfObject converts the pdfCIDFontType0 to a PDF representation.
func (_gacfbd *pdfCIDFontType0) ToPdfObject() _ebb.PdfObject { return _ebb.MakeNull() }
func (_dgad *PdfReader) newPdfAnnotationMarkupFromDict(_edgf *_ebb.PdfObjectDictionary) (*PdfAnnotationMarkup, error) {
	_cdae := &PdfAnnotationMarkup{}
	if _ecg := _edgf.Get("\u0054"); _ecg != nil {
		_cdae.T = _ecg
	}
	if _ggd := _edgf.Get("\u0050\u006f\u0070u\u0070"); _ggd != nil {
		_ccaa, _bddd := _ggd.(*_ebb.PdfIndirectObject)
		if !_bddd {
			if _, _ebag := _ggd.(*_ebb.PdfObjectNull); !_ebag {
				return nil, _gf.New("p\u006f\u0070\u0075\u0070\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074\u0020t\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			}
		} else {
			_bde, _ddga := _dgad.newPdfAnnotationFromIndirectObject(_ccaa)
			if _ddga != nil {
				return nil, _ddga
			}
			if _bde != nil {
				_dcb, _gfca := _bde._efd.(*PdfAnnotationPopup)
				if !_gfca {
					return nil, _gf.New("\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0020\u0070\u006f\u0070\u0075\u0070\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e")
				}
				_cdae.Popup = _dcb
			}
		}
	}
	if _eede := _edgf.Get("\u0043\u0041"); _eede != nil {
		_cdae.CA = _eede
	}
	if _cga := _edgf.Get("\u0052\u0043"); _cga != nil {
		_cdae.RC = _cga
	}
	if _dbc := _edgf.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"); _dbc != nil {
		_cdae.CreationDate = _dbc
	}
	if _fec := _edgf.Get("\u0049\u0052\u0054"); _fec != nil {
		_cdae.IRT = _fec
	}
	if _dbagd := _edgf.Get("\u0053\u0075\u0062\u006a"); _dbagd != nil {
		_cdae.Subj = _dbagd
	}
	if _fccbf := _edgf.Get("\u0052\u0054"); _fccbf != nil {
		_cdae.RT = _fccbf
	}
	if _cffa := _edgf.Get("\u0049\u0054"); _cffa != nil {
		_cdae.IT = _cffa
	}
	if _dcbg := _edgf.Get("\u0045\u0078\u0044\u0061\u0074\u0061"); _dcbg != nil {
		_cdae.ExData = _dcbg
	}
	return _cdae, nil
}
func (_ffdge SignatureValidationResult) String() string {
	var _cecdb _ca.Buffer
	_cecdb.WriteString(_bg.Sprintf("\u004ea\u006d\u0065\u003a\u0020\u0025\u0073\n", _ffdge.Name))
	if _ffdge.Date._dacdd > 0 {
		_cecdb.WriteString(_bg.Sprintf("\u0044a\u0074\u0065\u003a\u0020\u0025\u0073\n", _ffdge.Date.ToGoTime().String()))
	} else {
		_cecdb.WriteString("\u0044\u0061\u0074\u0065 n\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_ffdge.Reason) > 0 {
		_cecdb.WriteString(_bg.Sprintf("R\u0065\u0061\u0073\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _ffdge.Reason))
	} else {
		_cecdb.WriteString("N\u006f \u0072\u0065\u0061\u0073\u006f\u006e\u0020\u0073p\u0065\u0063\u0069\u0066ie\u0064\u000a")
	}
	if len(_ffdge.Location) > 0 {
		_cecdb.WriteString(_bg.Sprintf("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _ffdge.Location))
	} else {
		_cecdb.WriteString("\u004c\u006f\u0063at\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_ffdge.ContactInfo) > 0 {
		_cecdb.WriteString(_bg.Sprintf("\u0043\u006f\u006e\u0074\u0061\u0063\u0074\u0020\u0049\u006e\u0066\u006f:\u0020\u0025\u0073\u000a", _ffdge.ContactInfo))
	} else {
		_cecdb.WriteString("C\u006f\u006e\u0074\u0061\u0063\u0074 \u0069\u006e\u0066\u006f\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063i\u0066i\u0065\u0064\u000a")
	}
	_cecdb.WriteString(_bg.Sprintf("F\u0069\u0065\u006c\u0064\u0073\u003a\u0020\u0025\u0064\u000a", len(_ffdge.Fields)))
	if _ffdge.IsSigned {
		_cecdb.WriteString("S\u0069\u0067\u006e\u0065\u0064\u003a \u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073\u0020\u0073i\u0067n\u0065\u0064\u000a")
	} else {
		_cecdb.WriteString("\u0053\u0069\u0067\u006eed\u003a\u0020\u004e\u006f\u0074\u0020\u0073\u0069\u0067\u006e\u0065\u0064\u000a")
	}
	if _ffdge.IsVerified {
		_cecdb.WriteString("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0076\u0061\u006c\u0069\u0064\u0061t\u0069\u006f\u006e\u003a\u0020\u0049\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u000a")
	} else {
		_cecdb.WriteString("\u0053\u0069\u0067\u006e\u0061\u0074u\u0072\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e:\u0020\u0049\u0073\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u000a")
	}
	if _ffdge.IsTrusted {
		_cecdb.WriteString("\u0054\u0072\u0075\u0073\u0074\u0065\u0064\u003a\u0020\u0043\u0065\u0072\u0074\u0069\u0066i\u0063a\u0074\u0065\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0073\u0074\u0065\u0064\u000a")
	} else {
		_cecdb.WriteString("\u0054\u0072\u0075s\u0074\u0065\u0064\u003a \u0055\u006e\u0074\u0072\u0075\u0073\u0074e\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u000a")
	}
	if !_ffdge.GeneralizedTime.IsZero() {
		_cecdb.WriteString(_bg.Sprintf("G\u0065n\u0065\u0072\u0061\u006c\u0069\u007a\u0065\u0064T\u0069\u006d\u0065\u003a %\u0073\u000a", _ffdge.GeneralizedTime.String()))
	}
	if _ffdge.DiffResults != nil {
		_cecdb.WriteString(_bg.Sprintf("\u0064\u0069\u0066\u0066 i\u0073\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u003a\u0020\u0025v\u000a", _ffdge.DiffResults.IsPermitted()))
		if len(_ffdge.DiffResults.Warnings) > 0 {
			_cecdb.WriteString("\u004d\u0044\u0050\u0020\u0077\u0061\u0072\u006e\u0069n\u0067\u0073\u003a\u000a")
			for _, _dgbfeb := range _ffdge.DiffResults.Warnings {
				_cecdb.WriteString(_bg.Sprintf("\u0009\u0025\u0073\u000a", _dgbfeb))
			}
		}
		if len(_ffdge.DiffResults.Errors) > 0 {
			_cecdb.WriteString("\u004d\u0044\u0050 \u0065\u0072\u0072\u006f\u0072\u0073\u003a\u000a")
			for _, _eeaac := range _ffdge.DiffResults.Errors {
				_cecdb.WriteString(_bg.Sprintf("\u0009\u0025\u0073\u000a", _eeaac))
			}
		}
	}
	return _cecdb.String()
}

// ToPdfObject returns a stream object.
func (_bdbcd *XObjectForm) ToPdfObject() _ebb.PdfObject {
	_bgecb := _bdbcd._gebcd
	_bcbcf := _bgecb.PdfObjectDictionary
	if _bdbcd.Filter != nil {
		_bcbcf = _bdbcd.Filter.MakeStreamDict()
		_bgecb.PdfObjectDictionary = _bcbcf
	}
	_bcbcf.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_bcbcf.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0046\u006f\u0072\u006d"))
	_bcbcf.SetIfNotNil("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065", _bdbcd.FormType)
	_bcbcf.SetIfNotNil("\u0042\u0042\u006f\u0078", _bdbcd.BBox)
	_bcbcf.SetIfNotNil("\u004d\u0061\u0074\u0072\u0069\u0078", _bdbcd.Matrix)
	if _bdbcd.Resources != nil {
		_bcbcf.SetIfNotNil("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _bdbcd.Resources.ToPdfObject())
	}
	_bcbcf.SetIfNotNil("\u0047\u0072\u006fu\u0070", _bdbcd.Group)
	_bcbcf.SetIfNotNil("\u0052\u0065\u0066", _bdbcd.Ref)
	_bcbcf.SetIfNotNil("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061", _bdbcd.MetaData)
	_bcbcf.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _bdbcd.PieceInfo)
	_bcbcf.SetIfNotNil("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _bdbcd.LastModified)
	_bcbcf.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _bdbcd.StructParent)
	_bcbcf.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _bdbcd.StructParents)
	_bcbcf.SetIfNotNil("\u004f\u0050\u0049", _bdbcd.OPI)
	_bcbcf.SetIfNotNil("\u004f\u0043", _bdbcd.OC)
	_bcbcf.SetIfNotNil("\u004e\u0061\u006d\u0065", _bdbcd.Name)
	_bcbcf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ebb.MakeInteger(int64(len(_bdbcd.Stream))))
	_bgecb.Stream = _bdbcd.Stream
	return _bgecb
}
func (_bdddf *LTV) getCerts(_eecfbg []*_g.Certificate) ([][]byte, error) {
	_debf := make([][]byte, 0, len(_eecfbg))
	for _, _bcbg := range _eecfbg {
		_debf = append(_debf, _bcbg.Raw)
	}
	return _debf, nil
}

// PdfColor interface represents a generic color in PDF.
type PdfColor interface{}

// ToPdfObject implements interface PdfModel.
func (_cadc *PdfActionNamed) ToPdfObject() _ebb.PdfObject {
	_cadc.PdfAction.ToPdfObject()
	_dbdc := _cadc._abe
	_fcg := _dbdc.PdfObject.(*_ebb.PdfObjectDictionary)
	_fcg.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeNamed)))
	_fcg.SetIfNotNil("\u004e", _cadc.N)
	return _dbdc
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_affd *PdfColorspaceDeviceRGB) ToPdfObject() _ebb.PdfObject {
	return _ebb.MakeName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
}

// AlphaMapFunc represents a alpha mapping function: byte -> byte. Can be used for
// thresholding the alpha channel, i.e. setting all alpha values below threshold to transparent.
type AlphaMapFunc func(_gaegd byte) byte

// ToPdfObject returns the PDF representation of the function.
func (_afbcg *PdfFunctionType0) ToPdfObject() _ebb.PdfObject {
	if _afbcg._afgcc == nil {
		_afbcg._afgcc = &_ebb.PdfObjectStream{}
	}
	_fcae := _ebb.MakeDict()
	_fcae.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _ebb.MakeInteger(0))
	_ebafa := &_ebb.PdfObjectArray{}
	for _, _eedab := range _afbcg.Domain {
		_ebafa.Append(_ebb.MakeFloat(_eedab))
	}
	_fcae.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _ebafa)
	_fdfcb := &_ebb.PdfObjectArray{}
	for _, _fdfab := range _afbcg.Range {
		_fdfcb.Append(_ebb.MakeFloat(_fdfab))
	}
	_fcae.Set("\u0052\u0061\u006eg\u0065", _fdfcb)
	_agde := &_ebb.PdfObjectArray{}
	for _, _gegde := range _afbcg.Size {
		_agde.Append(_ebb.MakeInteger(int64(_gegde)))
	}
	_fcae.Set("\u0053\u0069\u007a\u0065", _agde)
	_fcae.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065", _ebb.MakeInteger(int64(_afbcg.BitsPerSample)))
	if _afbcg.Order != 1 {
		_fcae.Set("\u004f\u0072\u0064e\u0072", _ebb.MakeInteger(int64(_afbcg.Order)))
	}
	_fcae.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ebb.MakeInteger(int64(len(_afbcg._aeecg))))
	_afbcg._afgcc.Stream = _afbcg._aeecg
	_afbcg._afgcc.PdfObjectDictionary = _fcae
	return _afbcg._afgcc
}

// SetImageHandler sets the image handler used by the package.
func SetImageHandler(imgHandling ImageHandler) { ImageHandling = imgHandling }

// B returns the value of the blue component of the color.
func (_eecge *PdfColorDeviceRGB) B() float64 { return _eecge[2] }

// ImageToGray returns a new grayscale image based on the passed in RGB image.
func (_ddeb *PdfColorspaceDeviceRGB) ImageToGray(img Image) (Image, error) {
	if img.ColorComponents != 3 {
		return img, _gf.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u0020\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042")
	}
	_beeg, _cdad := _dg.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._dagcb, img._dgcea)
	if _cdad != nil {
		return img, _cdad
	}
	_bdaf, _cdad := _dg.GrayConverter.Convert(_beeg)
	if _cdad != nil {
		return img, _cdad
	}
	return _afacb(_bdaf.Base()), nil
}

// SetLocation sets the `Location` field of the signature.
func (_ebbgadb *PdfSignature) SetLocation(location string) {
	_ebbgadb.Location = _ebb.MakeString(location)
}

// GetNumComponents returns the number of color components (3 for CalRGB).
func (_cabd *PdfColorCalRGB) GetNumComponents() int { return 3 }

// PdfSignatureReference represents a PDF signature reference dictionary and is used for signing via form signature fields.
// (Section 12.8.1, Table 253 - Entries in a signature reference dictionary p. 469 in PDF32000_2008).
type PdfSignatureReference struct {
	_afacea         *_ebb.PdfObjectDictionary
	Type            *_ebb.PdfObjectName
	TransformMethod *_ebb.PdfObjectName
	TransformParams _ebb.PdfObject
	Data            _ebb.PdfObject
	DigestMethod    *_ebb.PdfObjectName
}

const (
	RC4_128bit = EncryptionAlgorithm(iota)
	AES_128bit
	AES_256bit
)

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

// SetContext sets the sub annotation (context).
func (_cfcc *PdfAnnotation) SetContext(ctx PdfModel) { _cfcc._efd = ctx }

// ToPdfObject implements interface PdfModel.
func (_abbd *PdfAnnotation3D) ToPdfObject() _ebb.PdfObject {
	_abbd.PdfAnnotation.ToPdfObject()
	_gfeb := _abbd._bdcd
	_afaa := _gfeb.PdfObject.(*_ebb.PdfObjectDictionary)
	_afaa.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0033\u0044"))
	_afaa.SetIfNotNil("\u0033\u0044\u0044", _abbd.T3DD)
	_afaa.SetIfNotNil("\u0033\u0044\u0056", _abbd.T3DV)
	_afaa.SetIfNotNil("\u0033\u0044\u0041", _abbd.T3DA)
	_afaa.SetIfNotNil("\u0033\u0044\u0049", _abbd.T3DI)
	_afaa.SetIfNotNil("\u0033\u0044\u0042", _abbd.T3DB)
	return _gfeb
}

// SetContext sets the sub action (context).
func (_acc *PdfAction) SetContext(ctx PdfModel) { _acc._ad = ctx }

// SetPdfCreationDate sets the CreationDate attribute of the output PDF.
func SetPdfCreationDate(creationDate _f.Time) {
	_daddc.Lock()
	defer _daddc.Unlock()
	_gecg = creationDate
}

// AddFont adds a font dictionary to the Font resources.
func (_acbb *PdfPage) AddFont(name _ebb.PdfObjectName, font _ebb.PdfObject) error {
	if _acbb.Resources == nil {
		_acbb.Resources = NewPdfPageResources()
	}
	if _acbb.Resources.Font == nil {
		_acbb.Resources.Font = _ebb.MakeDict()
	}
	_beef, _deab := _ebb.TraceToDirectObject(_acbb.Resources.Font).(*_ebb.PdfObjectDictionary)
	if !_deab {
		_eg.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0066\u006f\u006et \u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a \u0025\u0076", _ebb.TraceToDirectObject(_acbb.Resources.Font))
		return _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_beef.Set(name, font)
	return nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_aefd *PdfColorspaceSpecialIndexed) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	N := _aefd.Base.GetNumComponents()
	_ffbg := int(vals[0]) * N
	if _ffbg < 0 || (_ffbg+N-1) >= len(_aefd._eabg) {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _ffbg)
		return nil, ErrColorOutOfRange
	}
	_caeff := _aefd._eabg[_ffbg : _ffbg+N]
	var _aeab []float64
	for _, _fede := range _caeff {
		_aeab = append(_aeab, float64(_fede)/255.0)
	}
	_gbff, _dfgf := _aefd.Base.ColorFromFloats(_aeab)
	if _dfgf != nil {
		return nil, _dfgf
	}
	return _gbff, nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// L (range 0-100), A (range -100-100) and B (range -100-100) components of
// the color.
func (_cgffg *PdfColorspaceLab) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_afae := vals[0]
	if _afae < 0.0 || _afae > 100.0 {
		_eg.Log.Debug("\u004c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030-\u0031\u0030\u0030\u0029", _afae)
		return nil, ErrColorOutOfRange
	}
	_cagea := vals[1]
	_cfeb := float64(-100)
	_eafa := float64(100)
	if len(_cgffg.Range) > 1 {
		_cfeb = _cgffg.Range[0]
		_eafa = _cgffg.Range[1]
	}
	if _cagea < _cfeb || _cagea > _eafa {
		_eg.Log.Debug("\u0041\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _cagea, _cfeb, _eafa)
		return nil, ErrColorOutOfRange
	}
	_edgd := vals[2]
	_egbed := float64(-100)
	_ebgag := float64(100)
	if len(_cgffg.Range) > 3 {
		_egbed = _cgffg.Range[2]
		_ebgag = _cgffg.Range[3]
	}
	if _edgd < _egbed || _edgd > _ebgag {
		_eg.Log.Debug("\u0062\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _edgd, _egbed, _ebgag)
		return nil, ErrColorOutOfRange
	}
	_cgfe := NewPdfColorLab(_afae, _cagea, _edgd)
	return _cgfe, nil
}

// ColorToRGB converts gray -> rgb for a single color component.
func (_fgcd *PdfColorspaceDeviceGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_bbga, _acbe := color.(*PdfColorDeviceGray)
	if !_acbe {
		_eg.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006fr\u0020\u006e\u006f\u0074\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0067\u0072\u0061\u0079\u0020\u0025\u0054", color)
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return NewPdfColorDeviceRGB(float64(*_bbga), float64(*_bbga), float64(*_bbga)), nil
}

// ColorAt returns the color of the image pixel specified by the x and y coordinates.
func (_aaaff *Image) ColorAt(x, y int) (_e.Color, error) {
	_dbdce := _dg.BytesPerLine(int(_aaaff.Width), int(_aaaff.BitsPerComponent), _aaaff.ColorComponents)
	switch _aaaff.ColorComponents {
	case 1:
		return _dg.ColorAtGrayscale(x, y, int(_aaaff.BitsPerComponent), _dbdce, _aaaff.Data, _aaaff._dgcea)
	case 3:
		return _dg.ColorAtNRGBA(x, y, int(_aaaff.Width), _dbdce, int(_aaaff.BitsPerComponent), _aaaff.Data, _aaaff._dagcb, _aaaff._dgcea)
	case 4:
		return _dg.ColorAtCMYK(x, y, int(_aaaff.Width), _aaaff.Data, _aaaff._dgcea)
	}
	_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 i\u006da\u0067\u0065\u002e\u0020\u0025\u0064\u0020\u0063\u006f\u006d\u0070\u006fn\u0065\u006e\u0074\u0073\u002c\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072 \u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _aaaff.ColorComponents, _aaaff.BitsPerComponent)
	return nil, _gf.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061g\u0065 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
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
	ShadingType *_ebb.PdfObjectInteger
	ColorSpace  PdfColorspace
	Background  *_ebb.PdfObjectArray
	BBox        *PdfRectangle
	AntiAlias   *_ebb.PdfObjectBool
	_edgag      PdfModel
	_fbfae      _ebb.PdfObject
}

// GetSamples converts the raw byte slice into samples which are stored in a uint32 bit array.
// Each sample is represented by BitsPerComponent consecutive bits in the raw data.
// NOTE: The method resamples the image byte data before returning the result and
// this could lead to high memory usage, especially on large images. It should
// be avoided, when possible. It is recommended to access the Data field of the
// image directly or use the ColorAt method to extract individual pixels.
func (_bdcc *Image) GetSamples() []uint32 {
	_bgcc := _abg.ResampleBytes(_bdcc.Data, int(_bdcc.BitsPerComponent))
	if _bdcc.BitsPerComponent < 8 {
		_bgcc = _bdcc.samplesTrimPadding(_bgcc)
	}
	_gdgaga := int(_bdcc.Width) * int(_bdcc.Height) * _bdcc.ColorComponents
	if len(_bgcc) < _gdgaga {
		_eg.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u0066\u0065\u0077\u0020\u0073a\u006d\u0070\u006c\u0065\u0073\u0020(\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064\u0029", len(_bgcc), _gdgaga)
		return _bgcc
	} else if len(_bgcc) > _gdgaga {
		_eg.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u006d\u0061\u006e\u0079\u0020s\u0061\u006d\u0070\u006c\u0065\u0073 \u0028\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078p\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064", len(_bgcc), _gdgaga)
		_bgcc = _bgcc[:_gdgaga]
	}
	return _bgcc
}

// ToPdfObject implements interface PdfModel.
func (_aed *PdfActionGoTo) ToPdfObject() _ebb.PdfObject {
	_aed.PdfAction.ToPdfObject()
	_be := _aed._abe
	_faa := _be.PdfObject.(*_ebb.PdfObjectDictionary)
	_faa.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeGoTo)))
	_faa.SetIfNotNil("\u0044", _aed.D)
	return _be
}

// SetAlpha sets the alpha layer for the image.
func (_ddfde *Image) SetAlpha(alpha []byte) { _ddfde._dagcb = alpha }

// B returns the value of the B component of the color.
func (_ffead *PdfColorLab) B() float64 { return _ffead[2] }

// NewPdfColorspaceDeviceCMYK returns a new CMYK32 colorspace object.
func NewPdfColorspaceDeviceCMYK() *PdfColorspaceDeviceCMYK { return &PdfColorspaceDeviceCMYK{} }
func (_fdcd *PdfColorspaceSpecialPattern) String() string {
	return "\u0050a\u0074\u0074\u0065\u0072\u006e"
}

// GetCapHeight returns the CapHeight of the font `descriptor`.
func (_fgggd *PdfFontDescriptor) GetCapHeight() (float64, error) {
	return _ebb.GetNumberAsFloat(_fgggd.CapHeight)
}

// GetNumComponents returns the number of color components.
func (_cecf *PdfColorspaceICCBased) GetNumComponents() int { return _cecf.N }
func (_ge *PdfReader) newPdfActionGotoFromDict(_cef *_ebb.PdfObjectDictionary) (*PdfActionGoTo, error) {
	return &PdfActionGoTo{D: _cef.Get("\u0044")}, nil
}

// PdfWriter handles outputing PDF content.
type PdfWriter struct {
	_gegba        *_ebb.PdfIndirectObject
	_dggbf        *_ebb.PdfIndirectObject
	_afbdd        map[_ebb.PdfObject]struct{}
	_ebdgg        []_ebb.PdfObject
	_ffffd        map[_ebb.PdfObject]struct{}
	_addge        []*_ebb.PdfIndirectObject
	_bcbee        *PdfOutlineTreeNode
	_dffegd       *_ebb.PdfObjectDictionary
	_cgced        []_ebb.PdfObject
	_eadfd        *_ebb.PdfIndirectObject
	_cbabb        *_ba.Writer
	_afedd        int64
	_bgef         error
	_cgfde        *_ebb.PdfCrypt
	_gaccf        *_ebb.PdfObjectDictionary
	_cbcaa        *_ebb.PdfIndirectObject
	_eecfe        *_ebb.PdfObjectArray
	_efcge        _ebb.Version
	_adeff        *bool
	_eefeb        map[_ebb.PdfObject][]*_ebb.PdfObjectDictionary
	_dbea         *PdfAcroForm
	_aadcd        Optimizer
	_cafac        StandardApplier
	_bedfc        map[int]crossReference
	_fgdce        int64
	ObjNumOffset  int
	_abffb        bool
	_efdega       _ebb.XrefTable
	_bcage        int64
	_ggbfg        int64
	_cdgd         map[_ebb.PdfObject]int64
	_dcfg         map[_ebb.PdfObject]struct{}
	_cfecg        string
	_geced        []*PdfOutputIntent
	_abgcg        bool
	_gfdea, _gffb string
}

var _ pdfFont = (*pdfFontType3)(nil)

func (_gegbc *PdfWriter) adjustXRefAffectedVersion(_dfcad bool) {
	if _dfcad && _gegbc._efcge.Major == 1 && _gegbc._efcge.Minor < 5 {
		_gegbc._efcge.Minor = 5
	}
}
func (_gaf *PdfAnnotationMarkup) appendToPdfDictionary(_bbed *_ebb.PdfObjectDictionary) {
	_bbed.SetIfNotNil("\u0054", _gaf.T)
	if _gaf.Popup != nil {
		_bbed.Set("\u0050\u006f\u0070u\u0070", _gaf.Popup.ToPdfObject())
	}
	_bbed.SetIfNotNil("\u0043\u0041", _gaf.CA)
	_bbed.SetIfNotNil("\u0052\u0043", _gaf.RC)
	_bbed.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _gaf.CreationDate)
	_bbed.SetIfNotNil("\u0049\u0052\u0054", _gaf.IRT)
	_bbed.SetIfNotNil("\u0053\u0075\u0062\u006a", _gaf.Subj)
	_bbed.SetIfNotNil("\u0052\u0054", _gaf.RT)
	_bbed.SetIfNotNil("\u0049\u0054", _gaf.IT)
	_bbed.SetIfNotNil("\u0045\u0078\u0044\u0061\u0074\u0061", _gaf.ExData)
}

// Image interface is a basic representation of an image used in PDF.
// The colorspace is not specified, but must be known when handling the image.
type Image struct {
	Width            int64
	Height           int64
	BitsPerComponent int64
	ColorComponents  int
	Data             []byte
	_dagcb           []byte
	_dgcea           []float64
}

// NewPdfColorspaceSpecialSeparation returns a new separation color.
func NewPdfColorspaceSpecialSeparation() *PdfColorspaceSpecialSeparation {
	_bcfa := &PdfColorspaceSpecialSeparation{}
	return _bcfa
}

// PdfAnnotationSquare represents Square annotations.
// (Section 12.5.6.8).
type PdfAnnotationSquare struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _ebb.PdfObject
	IC _ebb.PdfObject
	BE _ebb.PdfObject
	RD _ebb.PdfObject
}

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
	CertClient *_cg.CertClient

	// OCSPClient is the client used to retrieve OCSP validation information.
	OCSPClient *_cg.OCSPClient

	// CRLClient is the client used to retrieve CRL validation information.
	CRLClient *_cg.CRLClient

	// SkipExisting specifies whether existing signature validations
	// should be skipped.
	SkipExisting bool
	_ggdbg       *PdfAppender
	_dfdgf       *DSS
}

const (
	_ PdfOutputIntentType = iota
	PdfOutputIntentTypeA1
	PdfOutputIntentTypeA2
	PdfOutputIntentTypeA3
	PdfOutputIntentTypeA4
	PdfOutputIntentTypeX
)

// IsRadio returns true if the button field represents a radio button, false otherwise.
func (_eedg *PdfFieldButton) IsRadio() bool { return _eedg.GetType() == ButtonTypeRadio }

// SignatureHandlerDocMDP extends SignatureHandler with the ValidateWithOpts method for checking the DocMDP policy.
type SignatureHandlerDocMDP interface {
	SignatureHandler

	// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser
	// ValidateWithOpts shall contain Validate call
	ValidateWithOpts(_dabaf *PdfSignature, _gaaef Hasher, _fgcfd SignatureHandlerDocMDPParams) (SignatureValidationResult, error)
}

// Add appends an outline item as a child of the current outline item.
func (_cacg *OutlineItem) Add(item *OutlineItem) { _cacg.Entries = append(_cacg.Entries, item) }

// L returns the value of the L component of the color.
func (_fgggb *PdfColorLab) L() float64 { return _fgggb[0] }
func (_bcgc *PdfReader) newPdfPageFromDict(_fbdca *_ebb.PdfObjectDictionary) (*PdfPage, error) {
	_ggbac := NewPdfPage()
	_ggbac._cdbfde = _fbdca
	_dbbf := *_fbdca
	_gdeg, _beaba := _dbbf.Get("\u0054\u0079\u0070\u0065").(*_ebb.PdfObjectName)
	if !_beaba {
		return nil, _gf.New("\u006d\u0069ss\u0069\u006e\u0067/\u0069\u006e\u0076\u0061lid\u0020Pa\u0067\u0065\u0020\u0064\u0069\u0063\u0074io\u006e\u0061\u0072\u0079\u0020\u0054\u0079p\u0065")
	}
	if *_gdeg != "\u0050\u0061\u0067\u0065" {
		return nil, _gf.New("\u0070\u0061\u0067\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u0050a\u0067\u0065")
	}
	if _cfea := _dbbf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"); _cfea != nil {
		_ggbac.Parent = _cfea
	}
	if _ceae := _dbbf.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"); _ceae != nil {
		_cbeg, _gadff := _ebb.GetString(_ceae)
		if !_gadff {
			return nil, _gf.New("\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u004c\u0061\u0073\u0074\u004d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0021=\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		_fgbee, _fgcg := NewPdfDate(_cbeg.Str())
		if _fgcg != nil {
			return nil, _fgcg
		}
		_ggbac.LastModified = &_fgbee
	}
	if _gdabd := _dbbf.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _gdabd != nil && !_ebb.IsNullObject(_gdabd) {
		_bgbec, _fdcgg := _ebb.GetDict(_gdabd)
		if !_fdcgg {
			return nil, _bg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _gdabd)
		}
		var _bfcab error
		_ggbac.Resources, _bfcab = NewPdfPageResourcesFromDict(_bgbec)
		if _bfcab != nil {
			return nil, _bfcab
		}
	} else {
		_agcbb, _aaafe := _ggbac.getParentResources()
		if _aaafe != nil {
			return nil, _aaafe
		}
		if _agcbb == nil {
			_agcbb = NewPdfPageResources()
		}
		_ggbac.Resources = _agcbb
	}
	if _ebbc := _dbbf.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _ebbc != nil {
		_gdbbe, _bgccb := _ebb.GetArray(_ebbc)
		if !_bgccb {
			return nil, _gf.New("\u0070\u0061\u0067\u0065\u0020\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _edfge error
		_ggbac.MediaBox, _edfge = NewPdfRectangle(*_gdbbe)
		if _edfge != nil {
			return nil, _edfge
		}
	}
	if _fcfeac := _dbbf.Get("\u0043r\u006f\u0070\u0042\u006f\u0078"); _fcfeac != nil {
		_fbadb, _gced := _ebb.GetArray(_fcfeac)
		if !_gced {
			return nil, _gf.New("\u0070a\u0067\u0065\u0020\u0043r\u006f\u0070\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _gcgaf error
		_ggbac.CropBox, _gcgaf = NewPdfRectangle(*_fbadb)
		if _gcgaf != nil {
			return nil, _gcgaf
		}
	}
	if _dddbb := _dbbf.Get("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078"); _dddbb != nil {
		_fedefa, _dadee := _ebb.GetArray(_dddbb)
		if !_dadee {
			return nil, _gf.New("\u0070\u0061\u0067\u0065\u0020\u0042\u006c\u0065\u0065\u0064\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _ecadc error
		_ggbac.BleedBox, _ecadc = NewPdfRectangle(*_fedefa)
		if _ecadc != nil {
			return nil, _ecadc
		}
	}
	if _cebb := _dbbf.Get("\u0054r\u0069\u006d\u0042\u006f\u0078"); _cebb != nil {
		_dgdea, _dgbb := _ebb.GetArray(_cebb)
		if !_dgbb {
			return nil, _gf.New("\u0070a\u0067\u0065\u0020\u0054r\u0069\u006d\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _fabb error
		_ggbac.TrimBox, _fabb = NewPdfRectangle(*_dgdea)
		if _fabb != nil {
			return nil, _fabb
		}
	}
	if _aaagg := _dbbf.Get("\u0041\u0072\u0074\u0042\u006f\u0078"); _aaagg != nil {
		_afdfe, _cgabb := _ebb.GetArray(_aaagg)
		if !_cgabb {
			return nil, _gf.New("\u0070a\u0067\u0065\u0020\u0041\u0072\u0074\u0042\u006f\u0078\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		var _fbgef error
		_ggbac.ArtBox, _fbgef = NewPdfRectangle(*_afdfe)
		if _fbgef != nil {
			return nil, _fbgef
		}
	}
	if _dggcd := _dbbf.Get("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f"); _dggcd != nil {
		_ggbac.BoxColorInfo = _dggcd
	}
	if _edaf := _dbbf.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _edaf != nil {
		_ggbac.Contents = _edaf
	}
	if _dbffag := _dbbf.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _dbffag != nil {
		_cgace, _bggfd := _ebb.GetNumberAsInt64(_dbffag)
		if _bggfd != nil {
			return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0067e\u0020\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_ggbac.Rotate = &_cgace
	}
	if _gfdfb := _dbbf.Get("\u0047\u0072\u006fu\u0070"); _gfdfb != nil {
		_ggbac.Group = _gfdfb
	}
	if _baafd := _dbbf.Get("\u0054\u0068\u0075m\u0062"); _baafd != nil {
		_ggbac.Thumb = _baafd
	}
	if _afbf := _dbbf.Get("\u0042"); _afbf != nil {
		_ggbac.B = _afbf
	}
	if _befeg := _dbbf.Get("\u0044\u0075\u0072"); _befeg != nil {
		_ggbac.Dur = _befeg
	}
	if _dfgcbb := _dbbf.Get("\u0054\u0072\u0061n\u0073"); _dfgcbb != nil {
		_ggbac.Trans = _dfgcbb
	}
	if _cceee := _dbbf.Get("\u0041\u0041"); _cceee != nil {
		_ggbac.AA = _cceee
	}
	if _cgabd := _dbbf.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _cgabd != nil {
		_ggbac.Metadata = _cgabd
	}
	if _ddgaa := _dbbf.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o"); _ddgaa != nil {
		_ggbac.PieceInfo = _ddgaa
	}
	if _fgcgg := _dbbf.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073"); _fgcgg != nil {
		_ggbac.StructParents = _fgcgg
	}
	if _bfebc := _dbbf.Get("\u0049\u0044"); _bfebc != nil {
		_ggbac.ID = _bfebc
	}
	if _ddaba := _dbbf.Get("\u0050\u005a"); _ddaba != nil {
		_ggbac.PZ = _ddaba
	}
	if _eecag := _dbbf.Get("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f"); _eecag != nil {
		_ggbac.SeparationInfo = _eecag
	}
	if _acab := _dbbf.Get("\u0054\u0061\u0062\u0073"); _acab != nil {
		_ggbac.Tabs = _acab
	}
	if _bcbbfa := _dbbf.Get("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064"); _bcbbfa != nil {
		_ggbac.TemplateInstantiated = _bcbbfa
	}
	if _aeagf := _dbbf.Get("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s"); _aeagf != nil {
		_ggbac.PresSteps = _aeagf
	}
	if _fbdf := _dbbf.Get("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074"); _fbdf != nil {
		_ggbac.UserUnit = _fbdf
	}
	if _gbbac := _dbbf.Get("\u0056\u0050"); _gbbac != nil {
		_ggbac.VP = _gbbac
	}
	if _bfcaa := _dbbf.Get("\u0041\u006e\u006e\u006f\u0074\u0073"); _bfcaa != nil {
		_ggbac.Annots = _bfcaa
	}
	_ggbac._ddab = _bcgc
	return _ggbac, nil
}

// String returns a human readable description of `fontfile`.
func (_bfgbe *fontFile) String() string {
	_dafgc := "\u005b\u004e\u006f\u006e\u0065\u005d"
	if _bfgbe._gega != nil {
		_dafgc = _bfgbe._gega.String()
	}
	return _bg.Sprintf("\u0046O\u004e\u0054\u0046\u0049\u004c\u0045\u007b\u0025\u0023\u0071\u0020e\u006e\u0063\u006f\u0064\u0065\u0072\u003d\u0025\u0073\u007d", _bfgbe._bgbcg, _dafgc)
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_bgce pdfCIDFontType0) GetRuneMetrics(r rune) (_bad.CharMetrics, bool) {
	return _bad.CharMetrics{Wx: _bgce._gbdb}, true
}
func _gebgc(_dacfb string) map[string]string {
	_ebbeg := _dbagf.Split(_dacfb, -1)
	_ffge := map[string]string{}
	for _, _gadf := range _ebbeg {
		_bfbee := _afgea.FindStringSubmatch(_gadf)
		if _bfbee == nil {
			continue
		}
		_eaagc, _geggg := _bfbee[1], _bfbee[2]
		_ffge[_eaagc] = _geggg
	}
	return _ffge
}

// ImageToRGB convert an indexed image to RGB.
func (_acddf *PdfColorspaceSpecialIndexed) ImageToRGB(img Image) (Image, error) {
	N := _acddf.Base.GetNumComponents()
	if N < 1 {
		return Image{}, _bg.Errorf("\u0062\u0061d \u0062\u0061\u0073e\u0020\u0063\u006f\u006cors\u0070ac\u0065\u0020\u004e\u0075\u006d\u0043\u006fmp\u006f\u006e\u0065\u006e\u0074\u0073\u003d%\u0064", N)
	}
	_caab := _dg.NewImageBase(int(img.Width), int(img.Height), 8, N, nil, img._dagcb, img._dgcea)
	_dggb := _abg.NewReader(img.getBase())
	_efee := _abg.NewWriter(_caab)
	var (
		_ceac  uint32
		_caedd int
		_fabae error
	)
	for {
		_ceac, _fabae = _dggb.ReadSample()
		if _fabae == _ab.EOF {
			break
		} else if _fabae != nil {
			return img, _fabae
		}
		_caedd = int(_ceac)
		_eg.Log.Trace("\u0049\u006ed\u0065\u0078\u0065\u0064\u003a\u0020\u0069\u006e\u0064\u0065\u0078\u003d\u0025\u0064\u0020\u004e\u003d\u0025\u0064\u0020\u006c\u0075t=\u0025\u0064", _caedd, N, len(_acddf._eabg))
		if (_caedd+1)*N > len(_acddf._eabg) {
			_caedd = len(_acddf._eabg)/N - 1
			_eg.Log.Trace("C\u006c\u0069\u0070\u0070in\u0067 \u0074\u006f\u0020\u0069\u006ed\u0065\u0078\u003a\u0020\u0025\u0064", _caedd)
			if _caedd < 0 {
				_eg.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069p\u0020\u0069\u006e\u0064\u0065\u0078.\u0020\u0049\u0073\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006ce\u0020\u0064\u0061\u006d\u0061\u0067\u0065\u0064\u003f")
				break
			}
		}
		for _ffdb := _caedd * N; _ffdb < (_caedd+1)*N; _ffdb++ {
			if _fabae = _efee.WriteSample(uint32(_acddf._eabg[_ffdb])); _fabae != nil {
				return img, _fabae
			}
		}
	}
	return _acddf.Base.ImageToRGB(_afacb(&_caab))
}

// ToPdfObject implements interface PdfModel.
func (_fdfd *PdfAnnotationPopup) ToPdfObject() _ebb.PdfObject {
	_fdfd.PdfAnnotation.ToPdfObject()
	_adagb := _fdfd._bdcd
	_aedf := _adagb.PdfObject.(*_ebb.PdfObjectDictionary)
	_aedf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0050\u006f\u0070u\u0070"))
	_aedf.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _fdfd.Parent)
	_aedf.SetIfNotNil("\u004f\u0070\u0065\u006e", _fdfd.Open)
	return _adagb
}

// GetContentStream returns the pattern cell's content stream
func (_dcccb *PdfTilingPattern) GetContentStream() ([]byte, error) {
	_bbfcg, _, _fbfgd := _dcccb.GetContentStreamWithEncoder()
	return _bbfcg, _fbfgd
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the reader using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
// At the end of the process, the AcroForm contains all the fields which were
// not flattened. If all fields are flattened, the reader's AcroForm field
// is set to nil.
func (_afff *PdfReader) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	return _afff.flattenFieldsWithOpts(false, appgen, opts)
}

// StdFontName represents name of a standard font.
type StdFontName = _bad.StdFontName

// PdfAnnotationMovie represents Movie annotations.
// (Section 12.5.6.17).
type PdfAnnotationMovie struct {
	*PdfAnnotation
	T     _ebb.PdfObject
	Movie _ebb.PdfObject
	A     _ebb.PdfObject
}

func (_fgea *PdfColorspaceDeviceGray) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
}

// NewGrayImageFromGoImage creates a new grayscale unidoc Image from a golang Image.
func (_fdbaa DefaultImageHandler) NewGrayImageFromGoImage(goimg _gdc.Image) (*Image, error) {
	_agfcb := goimg.Bounds()
	_eacag := &Image{Width: int64(_agfcb.Dx()), Height: int64(_agfcb.Dy()), ColorComponents: 1, BitsPerComponent: 8}
	switch _ffddc := goimg.(type) {
	case *_gdc.Gray:
		if len(_ffddc.Pix) != _agfcb.Dx()*_agfcb.Dy() {
			_cgdcb, _ecec := _dg.GrayConverter.Convert(goimg)
			if _ecec != nil {
				return nil, _ecec
			}
			_eacag.Data = _cgdcb.Pix()
		} else {
			_eacag.Data = _ffddc.Pix
		}
	case *_gdc.Gray16:
		_eacag.BitsPerComponent = 16
		if len(_ffddc.Pix) != _agfcb.Dx()*_agfcb.Dy()*2 {
			_faea, _eadg := _dg.Gray16Converter.Convert(goimg)
			if _eadg != nil {
				return nil, _eadg
			}
			_eacag.Data = _faea.Pix()
		} else {
			_eacag.Data = _ffddc.Pix
		}
	case _dg.Image:
		_dggbb := _ffddc.Base()
		if _dggbb.ColorComponents == 1 {
			_eacag.BitsPerComponent = int64(_dggbb.BitsPerComponent)
			_eacag.Data = _dggbb.Data
			return _eacag, nil
		}
		_fdaba, _facg := _dg.GrayConverter.Convert(goimg)
		if _facg != nil {
			return nil, _facg
		}
		_eacag.Data = _fdaba.Pix()
	default:
		_adbg, _cdbeg := _dg.GrayConverter.Convert(goimg)
		if _cdbeg != nil {
			return nil, _cdbeg
		}
		_eacag.Data = _adbg.Pix()
	}
	return _eacag, nil
}

// ToPdfObject implements interface PdfModel.
func (_bdad *PdfActionHide) ToPdfObject() _ebb.PdfObject {
	_bdad.PdfAction.ToPdfObject()
	_dde := _bdad._abe
	_fge := _dde.PdfObject.(*_ebb.PdfObjectDictionary)
	_fge.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeHide)))
	_fge.SetIfNotNil("\u0054", _bdad.T)
	_fge.SetIfNotNil("\u0048", _bdad.H)
	return _dde
}

// GetRevision returns the specific version of the PdfReader for the current Pdf document
func (_fdgeg *PdfReader) GetRevision(revisionNumber int) (*PdfReader, error) {
	_gfdcb := _fdgeg._cafdf.GetRevisionNumber()
	if revisionNumber < 0 || revisionNumber > _gfdcb {
		return nil, _gf.New("w\u0072\u006f\u006e\u0067 r\u0065v\u0069\u0073\u0069\u006f\u006e \u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if revisionNumber == _gfdcb {
		return _fdgeg, nil
	}
	if _fdgeg._face[revisionNumber] != nil {
		return _fdgeg._face[revisionNumber], nil
	}
	_bfeee := _fdgeg
	for _beccd := _gfdcb - 1; _beccd >= revisionNumber; _beccd-- {
		_eddae, _ceba := _bfeee.GetPreviousRevision()
		if _ceba != nil {
			return nil, _ceba
		}
		_fdgeg._face[_beccd] = _eddae
		_bfeee = _eddae
	}
	return _bfeee, nil
}

// DecodeArray returns the range of color component values in CalGray colorspace.
func (_gdd *PdfColorspaceCalGray) DecodeArray() []float64 { return []float64{0.0, 1.0} }

// ToPdfObject implements interface PdfModel.
func (_fbbe *PdfAnnotationFileAttachment) ToPdfObject() _ebb.PdfObject {
	_fbbe.PdfAnnotation.ToPdfObject()
	_bbfg := _fbbe._bdcd
	_cdda := _bbfg.PdfObject.(*_ebb.PdfObjectDictionary)
	_fbbe.PdfAnnotationMarkup.appendToPdfDictionary(_cdda)
	_cdda.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074"))
	_cdda.SetIfNotNil("\u0046\u0053", _fbbe.FS)
	_cdda.SetIfNotNil("\u004e\u0061\u006d\u0065", _fbbe.Name)
	return _bbfg
}

// SetColorSpace sets `r` colorspace object to `colorspace`.
func (_dbgf *PdfPageResources) SetColorSpace(colorspace *PdfPageResourcesColorspaces) {
	_dbgf._aaee = colorspace
}
func (_gabc *PdfReader) newPdfAnnotationMovieFromDict(_ffgb *_ebb.PdfObjectDictionary) (*PdfAnnotationMovie, error) {
	_fdab := PdfAnnotationMovie{}
	_fdab.T = _ffgb.Get("\u0054")
	_fdab.Movie = _ffgb.Get("\u004d\u006f\u0076i\u0065")
	_fdab.A = _ffgb.Get("\u0041")
	return &_fdab, nil
}

// NewPdfReaderFromFile creates a new PdfReader from the speficied PDF file.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderFromFile(pdfFile string, opts *ReaderOpts) (*PdfReader, *_ed.File, error) {
	const _acffb = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0046\u0072\u006f\u006dF\u0069\u006c\u0065"
	_gggba, _cdead := _ed.Open(pdfFile)
	if _cdead != nil {
		return nil, nil, _cdead
	}
	_gffff, _cdead := _dcbd(_gggba, opts, true, _acffb)
	if _cdead != nil {
		_gggba.Close()
		return nil, nil, _cdead
	}
	return _gffff, _gggba, nil
}
func (_dfeg *fontFile) parseASCIIPart(_febd []byte) error {
	if len(_febd) < 2 || string(_febd[:2]) != "\u0025\u0021" {
		return _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0074a\u0072\u0074\u0020\u006f\u0066\u0020\u0041S\u0043\u0049\u0049\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_ffcef, _daeeb, _addbbd := _gefgb(_febd)
	if _addbbd != nil {
		return _addbbd
	}
	_abag := _gebgc(_ffcef)
	_dfeg._bgbcg = _abag["\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"]
	if _dfeg._bgbcg == "" {
		_eg.Log.Debug("\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0020\u0068a\u0073\u0020\u006e\u006f\u0020\u002f\u0046\u006f\u006e\u0074N\u0061\u006d\u0065")
	}
	if _daeeb != "" {
		_edaef, _fddbe := _dddgd(_daeeb)
		if _fddbe != nil {
			return _fddbe
		}
		_egfef, _fddbe := _da.NewCustomSimpleTextEncoder(_edaef, nil)
		if _fddbe != nil {
			_eg.Log.Debug("\u0045\u0052\u0052\u004fR\u0020\u003a\u0055\u004e\u004b\u004e\u004f\u0057\u004e\u0020G\u004cY\u0050\u0048\u003a\u0020\u0065\u0072\u0072=\u0025\u0076", _fddbe)
			return nil
		}
		_dfeg._gega = _egfef
	}
	return nil
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
	_aeecg        []byte
	_egdg         []uint32
	_afgcc        *_ebb.PdfObjectStream
}

// SetPdfKeywords sets the Keywords attribute of the output PDF.
func SetPdfKeywords(keywords string) { _daddc.Lock(); defer _daddc.Unlock(); _fceef = keywords }

type pdfFontType0 struct {
	fontCommon
	_dfffc         *_ebb.PdfIndirectObject
	_bfdgc         _da.TextEncoder
	Encoding       _ebb.PdfObject
	DescendantFont *PdfFont
	_efeb          *_ebe.CMap
}

func (_bcfdg *PdfAppender) addNewObject(_eaba _ebb.PdfObject) {
	if _, _ebgeb := _bcfdg._ddfg[_eaba]; !_ebgeb {
		_bcfdg._bfeg = append(_bcfdg._bfeg, _eaba)
		_bcfdg._ddfg[_eaba] = struct{}{}
	}
}

// GetContainingPdfObject returns the container of the DSS (indirect object).
func (_egfd *DSS) GetContainingPdfObject() _ebb.PdfObject { return _egfd._fcgb }

// ColorToRGB converts an Indexed color to an RGB color.
func (_dgfe *PdfColorspaceSpecialIndexed) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _dgfe.Base == nil {
		return nil, _gf.New("\u0069\u006e\u0064\u0065\u0078\u0065d\u0020\u0062\u0061\u0073\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _dgfe.Base.ColorToRGB(color)
}
func (_addcb *PdfReader) loadStructure() error {
	if _addcb._cafdf.GetCrypter() != nil && !_addcb._cafdf.IsAuthenticated() {
		return _bg.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_dcafa := _addcb._cafdf.GetTrailer()
	if _dcafa == nil {
		return _bg.Errorf("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_ggacb, _aede := _dcafa.Get("\u0052\u006f\u006f\u0074").(*_ebb.PdfObjectReference)
	if !_aede {
		return _bg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052\u006f\u006ft\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a \u0025\u0073\u0029", _dcafa)
	}
	_bebeb, _fdbgb := _addcb._cafdf.LookupByReference(*_ggacb)
	if _fdbgb != nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0072\u006f\u006f\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _fdbgb)
		return _fdbgb
	}
	_eddde, _aede := _bebeb.(*_ebb.PdfIndirectObject)
	if !_aede {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _bebeb, *_dcafa)
		return _gf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_edcbb, _aede := (*_eddde).PdfObject.(*_ebb.PdfObjectDictionary)
	if !_aede {
		_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0061t\u0061\u006c\u006fg\u0020(\u0025\u0073\u0029", _eddde.PdfObject)
		return _gf.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_eg.Log.Trace("C\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _edcbb)
	_dgged, _aede := _edcbb.Get("\u0050\u0061\u0067e\u0073").(*_ebb.PdfObjectReference)
	if !_aede {
		return _gf.New("\u0070\u0061\u0067\u0065\u0073\u0020\u0069\u006e\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0061\u0020\u0072e\u0066\u0065r\u0065\u006e\u0063\u0065")
	}
	_gcbce, _fdbgb := _addcb._cafdf.LookupByReference(*_dgged)
	if _fdbgb != nil {
		_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020F\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020r\u0065\u0061\u0064 \u0070a\u0067\u0065\u0073")
		return _fdbgb
	}
	_cadaf, _aede := _gcbce.(*_ebb.PdfIndirectObject)
	if !_aede {
		_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020P\u0061\u0067\u0065\u0073\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0069n\u0076a\u006c\u0069\u0064")
		_eg.Log.Debug("\u006f\u0070\u003a\u0020\u0025\u0070", _cadaf)
		return _gf.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_bgcfb, _aede := _cadaf.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_aede {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065c\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0073\u0029", _cadaf)
		return _gf.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_agcef, _aede := _ebb.GetInt(_bgcfb.Get("\u0043\u006f\u0075n\u0074"))
	if !_aede {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0050\u0061\u0067\u0065\u0073\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return _gf.New("\u0070\u0061\u0067\u0065s \u0063\u006f\u0075\u006e\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _, _aede = _ebb.GetName(_bgcfb.Get("\u0054\u0079\u0070\u0065")); !_aede {
		_eg.Log.Debug("\u0050\u0061\u0067\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020T\u0079\u0070\u0065\u0020\u0066\u0069\u0065\u006cd\u0020n\u006f\u0074\u0020\u0073\u0065\u0074\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0054\u0079p\u0065\u0020\u0074\u006f\u0020\u0050\u0061\u0067\u0065\u0073\u002e")
		_bgcfb.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0050\u0061\u0067e\u0073"))
	}
	if _afgac, _ebaega := _ebb.GetInt(_bgcfb.Get("\u0052\u006f\u0074\u0061\u0074\u0065")); _ebaega {
		_fcdage := int64(*_afgac)
		_addcb.Rotate = &_fcdage
	}
	_addcb._agbbe = _ggacb
	_addcb._fdgda = _edcbb
	_addcb._egea = _bgcfb
	_addcb._eedbb = _cadaf
	_addcb._aadcb = int(*_agcef)
	_addcb._faebb = []*_ebb.PdfIndirectObject{}
	_ccabe := map[_ebb.PdfObject]struct{}{}
	_fdbgb = _addcb.buildPageList(_cadaf, nil, _ccabe)
	if _fdbgb != nil {
		return _fdbgb
	}
	_eg.Log.Trace("\u002d\u002d\u002d")
	_eg.Log.Trace("\u0054\u004f\u0043")
	_eg.Log.Trace("\u0050\u0061\u0067e\u0073")
	_eg.Log.Trace("\u0025\u0064\u003a\u0020\u0025\u0073", len(_addcb._faebb), _addcb._faebb)
	_addcb._fgbcg, _fdbgb = _addcb.loadOutlines()
	if _fdbgb != nil {
		_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075i\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 t\u0072\u0065\u0065 \u0028%\u0073\u0029", _fdbgb)
		return _fdbgb
	}
	_addcb.AcroForm, _fdbgb = _addcb.loadForms()
	if _fdbgb != nil {
		return _fdbgb
	}
	_addcb.DSS, _fdbgb = _addcb.loadDSS()
	if _fdbgb != nil {
		return _fdbgb
	}
	_addcb._fdca, _fdbgb = _addcb.loadPerms()
	if _fdbgb != nil {
		return _fdbgb
	}
	return nil
}
func (_ffcb *PdfWriter) updateObjectNumbers() {
	_eeegdb := _ffcb.ObjNumOffset
	_gddfd := 0
	for _, _abdde := range _ffcb._ebdgg {
		_gcgbf := int64(_gddfd + 1 + _eeegdb)
		_dedff := true
		if _ffcb._abffb {
			if _bbcdb, _cafdc := _ffcb._cdgd[_abdde]; _cafdc {
				_gcgbf = _bbcdb
				_dedff = false
			}
		}
		switch _geeda := _abdde.(type) {
		case *_ebb.PdfIndirectObject:
			_geeda.ObjectNumber = _gcgbf
			_geeda.GenerationNumber = 0
		case *_ebb.PdfObjectStream:
			_geeda.ObjectNumber = _gcgbf
			_geeda.GenerationNumber = 0
		case *_ebb.PdfObjectStreams:
			_geeda.ObjectNumber = _gcgbf
			_geeda.GenerationNumber = 0
		default:
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u0020%\u0054\u0020\u002d\u0020\u0073\u006b\u0069p\u0070\u0069\u006e\u0067", _geeda)
			continue
		}
		if _dedff {
			_gddfd++
		}
	}
	_bbffb := func(_efccdb _ebb.PdfObject) int64 {
		switch _gdaae := _efccdb.(type) {
		case *_ebb.PdfIndirectObject:
			return _gdaae.ObjectNumber
		case *_ebb.PdfObjectStream:
			return _gdaae.ObjectNumber
		case *_ebb.PdfObjectStreams:
			return _gdaae.ObjectNumber
		}
		return 0
	}
	_ae.SliceStable(_ffcb._ebdgg, func(_cebbg, _ddfab int) bool { return _bbffb(_ffcb._ebdgg[_cebbg]) < _bbffb(_ffcb._ebdgg[_ddfab]) })
}

// NewPdfOutline returns an initialized PdfOutline.
func NewPdfOutline() *PdfOutline {
	_dbefg := &PdfOutline{_egee: _ebb.MakeIndirectObject(_ebb.MakeDict())}
	_dbefg._geeee = _dbefg
	return _dbefg
}

// SetPdfProducer sets the Producer attribute of the output PDF.
func SetPdfProducer(producer string) { _daddc.Lock(); defer _daddc.Unlock(); _aaaae = producer }

// ParserMetadata gets the parser  metadata.
func (_aadg *CompliancePdfReader) ParserMetadata() _ebb.ParserMetadata {
	if _aadg._ecbbg == (_ebb.ParserMetadata{}) {
		_aadg._ecbbg, _ = _aadg._cafdf.ParserMetadata()
	}
	return _aadg._ecbbg
}
func (_bgfdf *pdfFontType0) bytesToCharcodes(_bgfde []byte) ([]_da.CharCode, bool) {
	if _bgfdf._efeb == nil {
		return nil, false
	}
	_cgcca, _fcdf := _bgfdf._efeb.BytesToCharcodes(_bgfde)
	if !_fcdf {
		return nil, false
	}
	_cbfb := make([]_da.CharCode, len(_cgcca))
	for _cadaa, _daea := range _cgcca {
		_cbfb[_cadaa] = _da.CharCode(_daea)
	}
	return _cbfb, true
}

// GetNumComponents returns the number of color components (4 for CMYK32).
func (_bdgg *PdfColorDeviceCMYK) GetNumComponents() int { return 4 }

// PdfShadingType2 is an Axial shading.
type PdfShadingType2 struct {
	*PdfShading
	Coords   *_ebb.PdfObjectArray
	Domain   *_ebb.PdfObjectArray
	Function []PdfFunction
	Extend   *_ebb.PdfObjectArray
}

// ColorToRGB converts a CalGray color to an RGB color.
func (_ddff *PdfColorspaceCalGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_cdea, _afab := color.(*PdfColorCalGray)
	if !_afab {
		_eg.Log.Debug("\u0049n\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006eo\u0074\u0020\u0063\u0061\u006c\u0020\u0067\u0072\u0061\u0079")
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	ANorm := _cdea.Val()
	X := _ddff.WhitePoint[0] * _cbg.Pow(ANorm, _ddff.Gamma)
	Y := _ddff.WhitePoint[1] * _cbg.Pow(ANorm, _ddff.Gamma)
	Z := _ddff.WhitePoint[2] * _cbg.Pow(ANorm, _ddff.Gamma)
	_fdce := 3.240479*X + -1.537150*Y + -0.498535*Z
	_bfff := -0.969256*X + 1.875992*Y + 0.041556*Z
	_abggf := 0.055648*X + -0.204043*Y + 1.057311*Z
	_fdce = _cbg.Min(_cbg.Max(_fdce, 0), 1.0)
	_bfff = _cbg.Min(_cbg.Max(_bfff, 0), 1.0)
	_abggf = _cbg.Min(_cbg.Max(_abggf, 0), 1.0)
	return NewPdfColorDeviceRGB(_fdce, _bfff, _abggf), nil
}
func (_dfbgb *PdfReader) buildNameNodes(_aegbb *_ebb.PdfIndirectObject, _cgfcd map[_ebb.PdfObject]struct{}) error {
	if _aegbb == nil {
		return nil
	}
	if _, _ddfbb := _cgfcd[_aegbb]; _ddfbb {
		_eg.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _aegbb.ObjectNumber)
		return nil
	}
	_cgfcd[_aegbb] = struct{}{}
	_dbbe, _dafae := _aegbb.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_dafae {
		return _gf.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _cecg, _gbfb := _ebb.GetDict(_dbbe.Get("\u0044\u0065\u0073t\u0073")); _gbfb {
		_beff, _deadg := _ebb.GetArray(_cecg.Get("\u004b\u0069\u0064\u0073"))
		if !_deadg {
			return _gf.New("\u0049n\u0076\u0061\u006c\u0069d\u0020\u004b\u0069\u0064\u0073 \u0061r\u0072a\u0079\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_eg.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _beff)
		for _bedaf, _bacff := range _beff.Elements() {
			_adadg, _bfgda := _ebb.GetIndirect(_bacff)
			if !_bfgda {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u0068\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002d \u0028\u0025\u0073\u0029", _adadg)
				return _gf.New("\u0063h\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006ed\u0069r\u0065c\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
			}
			_beff.Set(_bedaf, _adadg)
			_cfed := _dfbgb.buildNameNodes(_adadg, _cgfcd)
			if _cfed != nil {
				return _cfed
			}
		}
	}
	if _fcab, _eacae := _ebb.GetDict(_dbbe); _eacae {
		if !_ebb.IsNullObject(_fcab.Get("\u004b\u0069\u0064\u0073")) {
			if _ebedb, _cdgfd := _ebb.GetArray(_fcab.Get("\u004b\u0069\u0064\u0073")); _cdgfd {
				for _fbgb, _bgbf := range _ebedb.Elements() {
					if _dfbdc, _abca := _ebb.GetIndirect(_bgbf); _abca {
						_ebedb.Set(_fbgb, _dfbdc)
						_gfedgg := _dfbgb.buildNameNodes(_dfbdc, _cgfcd)
						if _gfedgg != nil {
							return _gfedgg
						}
					}
				}
			}
		}
	}
	return nil
}

// SetOCProperties sets the optional content properties.
func (_afbae *PdfWriter) SetOCProperties(ocProperties _ebb.PdfObject) error {
	_gbag := _afbae._dffegd
	if ocProperties != nil {
		_eg.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u004f\u0043\u0020\u0050\u0072o\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002e\u002e\u002e")
		_gbag.Set("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", ocProperties)
		return _afbae.addObjects(ocProperties)
	}
	return nil
}
func _dffb(_agcda _ebb.PdfObject) (*PdfColorspaceSpecialIndexed, error) {
	_gaecb := NewPdfColorspaceSpecialIndexed()
	if _gade, _ecge := _agcda.(*_ebb.PdfIndirectObject); _ecge {
		_gaecb._gdaef = _gade
	}
	_agcda = _ebb.TraceToDirectObject(_agcda)
	_facbf, _eeeb := _agcda.(*_ebb.PdfObjectArray)
	if !_eeeb {
		return nil, _bg.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _facbf.Len() != 4 {
		return nil, _bg.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_agcda = _facbf.Get(0)
	_geaa, _eeeb := _agcda.(*_ebb.PdfObjectName)
	if !_eeeb {
		return nil, _bg.Errorf("\u0069n\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
	}
	if *_geaa != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _bg.Errorf("\u0069\u006e\u0064\u0065xe\u0064\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u006e\u0061m\u0065")
	}
	_agcda = _facbf.Get(1)
	_fggge, _bbeg := DetermineColorspaceNameFromPdfObject(_agcda)
	if _bbeg != nil {
		return nil, _bbeg
	}
	if _fggge == "\u0049n\u0064\u0065\u0078\u0065\u0064" || _fggge == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
		_eg.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0049\u006e\u0064e\u0078\u0065\u0064\u002f\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043S\u0020\u0061\u0073\u0020\u0062\u0061\u0073\u0065\u0020\u0028\u0025v\u0029", _fggge)
		return nil, _fddb
	}
	_cdfc, _bbeg := NewPdfColorspaceFromPdfObject(_agcda)
	if _bbeg != nil {
		return nil, _bbeg
	}
	_gaecb.Base = _cdfc
	_agcda = _facbf.Get(2)
	_ggcc, _bbeg := _ebb.GetNumberAsInt64(_agcda)
	if _bbeg != nil {
		return nil, _bbeg
	}
	if _ggcc > 255 {
		return nil, _bg.Errorf("\u0069n\u0064\u0065\u0078\u0065d\u0020\u0043\u0053\u003a\u0020I\u006ev\u0061l\u0069\u0064\u0020\u0068\u0069\u0076\u0061l")
	}
	_gaecb.HiVal = int(_ggcc)
	_agcda = _facbf.Get(3)
	_gaecb.Lookup = _agcda
	_agcda = _ebb.TraceToDirectObject(_agcda)
	var _ebca []byte
	if _cafa, _cggb := _agcda.(*_ebb.PdfObjectString); _cggb {
		_ebca = _cafa.Bytes()
		_eg.Log.Trace("\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0020\u0064", _ebca)
	} else if _eaae, _gfbda := _agcda.(*_ebb.PdfObjectStream); _gfbda {
		_eg.Log.Trace("\u0049n\u0064e\u0078\u0065\u0064\u0020\u0073t\u0072\u0065a\u006d\u003a\u0020\u0025\u0073", _agcda.String())
		_eg.Log.Trace("\u0045\u006e\u0063\u006fde\u0064\u0020\u0028\u0025\u0064\u0029\u0020\u003a\u0020\u0025\u0023\u0020\u0078", len(_eaae.Stream), _eaae.Stream)
		_dcfb, _fggd := _ebb.DecodeStream(_eaae)
		if _fggd != nil {
			return nil, _fggd
		}
		_eg.Log.Trace("\u0044e\u0063o\u0064\u0065\u0064\u0020\u0028%\u0064\u0029 \u003a\u0020\u0025\u0020\u0058", len(_dcfb), _dcfb)
		_ebca = _dcfb
	} else {
		_eg.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _agcda)
		return nil, _bg.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076a\u006ci\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if len(_ebca) < _gaecb.Base.GetNumComponents()*(_gaecb.HiVal+1) {
		_eg.Log.Debug("\u0050\u0044\u0046\u0020\u0049\u006e\u0063o\u006d\u0070\u0061t\u0069\u0062\u0069\u006ci\u0074\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074\u006f\u006f\u0020\u0073\u0068\u006f\u0072\u0074")
		_eg.Log.Debug("\u0046\u0061i\u006c\u002c\u0020\u006c\u0065\u006e\u0028\u0064\u0061\u0074\u0061\u0029\u003a\u0020\u0025\u0064\u002c\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u002c\u0020\u0068\u0069\u0056\u0061\u006c\u003a\u0020\u0025\u0064", len(_ebca), _gaecb.Base.GetNumComponents(), _gaecb.HiVal)
	} else {
		_ebca = _ebca[:_gaecb.Base.GetNumComponents()*(_gaecb.HiVal+1)]
	}
	_gaecb._eabg = _ebca
	return _gaecb, nil
}

// NewPdfColorspaceLab returns a new Lab colorspace object.
func NewPdfColorspaceLab() *PdfColorspaceLab {
	_agcb := &PdfColorspaceLab{}
	_agcb.BlackPoint = []float64{0.0, 0.0, 0.0}
	_agcb.Range = []float64{-100, 100, -100, 100}
	return _agcb
}
func (_fdcee *PdfReader) loadDSS() (*DSS, error) {
	if _fdcee._cafdf.GetCrypter() != nil && !_fdcee._cafdf.IsAuthenticated() {
		return nil, _bg.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_egcgb := _fdcee._fdgda.Get("\u0044\u0053\u0053")
	if _egcgb == nil {
		return nil, nil
	}
	_cbecc, _ := _ebb.GetIndirect(_egcgb)
	_egcgb = _ebb.TraceToDirectObject(_egcgb)
	switch _ebffg := _egcgb.(type) {
	case *_ebb.PdfObjectNull:
		return nil, nil
	case *_ebb.PdfObjectDictionary:
		return _afag(_cbecc, _ebffg)
	}
	return nil, _bg.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u0053\u0053 \u0065\u006e\u0074\u0072y \u0025\u0054", _egcgb)
}

// DecodeArray returns the range of color component values in the ICCBased colorspace.
func (_cfcce *PdfColorspaceICCBased) DecodeArray() []float64 { return _cfcce.Range }

// ImageToRGB converts an image in CMYK32 colorspace to an RGB image.
func (_edea *PdfColorspaceDeviceCMYK) ImageToRGB(img Image) (Image, error) {
	_eg.Log.Trace("\u0043\u004d\u0059\u004b\u0033\u0032\u0020\u002d\u003e\u0020\u0052\u0047\u0042")
	_eg.Log.Trace("I\u006d\u0061\u0067\u0065\u0020\u0042P\u0043\u003a\u0020\u0025\u0064\u002c \u0043\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020%\u0064", img.BitsPerComponent, img.ColorComponents)
	_eg.Log.Trace("\u004c\u0065\u006e \u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0064", len(img.Data))
	_eg.Log.Trace("H\u0065\u0069\u0067\u0068t:\u0020%\u0064\u002c\u0020\u0057\u0069d\u0074\u0068\u003a\u0020\u0025\u0064", img.Height, img.Width)
	_feee, _abbgg := _dg.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._dagcb, img._dgcea)
	if _abbgg != nil {
		return Image{}, _abbgg
	}
	_bbagg, _abbgg := _dg.NRGBAConverter.Convert(_feee)
	if _abbgg != nil {
		return Image{}, _abbgg
	}
	return _afacb(_bbagg.Base()), nil
}

// ToPdfObject implements interface PdfModel.
func (_eee *PdfAnnotationSquiggly) ToPdfObject() _ebb.PdfObject {
	_eee.PdfAnnotation.ToPdfObject()
	_bbcd := _eee._bdcd
	_ecgc := _bbcd.PdfObject.(*_ebb.PdfObjectDictionary)
	_eee.PdfAnnotationMarkup.appendToPdfDictionary(_ecgc)
	_ecgc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079"))
	_ecgc.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _eee.QuadPoints)
	return _bbcd
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_gddbg *PdfShadingType6) ToPdfObject() _ebb.PdfObject {
	_gddbg.PdfShading.ToPdfObject()
	_egged, _dfdaba := _gddbg.getShadingDict()
	if _dfdaba != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _gddbg.BitsPerCoordinate != nil {
		_egged.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _gddbg.BitsPerCoordinate)
	}
	if _gddbg.BitsPerComponent != nil {
		_egged.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _gddbg.BitsPerComponent)
	}
	if _gddbg.BitsPerFlag != nil {
		_egged.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _gddbg.BitsPerFlag)
	}
	if _gddbg.Decode != nil {
		_egged.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _gddbg.Decode)
	}
	if _gddbg.Function != nil {
		if len(_gddbg.Function) == 1 {
			_egged.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _gddbg.Function[0].ToPdfObject())
		} else {
			_ggecg := _ebb.MakeArray()
			for _, _ebbgad := range _gddbg.Function {
				_ggecg.Append(_ebbgad.ToPdfObject())
			}
			_egged.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _ggecg)
		}
	}
	return _gddbg._fbfae
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a CalGray device.
func (_dgaaf *PdfColorspaceCalGray) GetNumComponents() int { return 1 }

// HasExtGState checks whether a font is defined by the specified keyName.
func (_gebbd *PdfPageResources) HasExtGState(keyName _ebb.PdfObjectName) bool {
	_, _gdcge := _gebbd.GetFontByName(keyName)
	return _gdcge
}

// NewPdfColorDeviceGray returns a new grayscale color based on an input grayscale float value in range [0-1].
func NewPdfColorDeviceGray(grayVal float64) *PdfColorDeviceGray {
	_eaeef := PdfColorDeviceGray(grayVal)
	return &_eaeef
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_beeb *PdfColorspaceSpecialSeparation) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bafdd := vals[0]
	_eecfbd := []float64{_bafdd}
	_gcfgb, _egced := _beeb.TintTransform.Evaluate(_eecfbd)
	if _egced != nil {
		_eg.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0065: \u0025\u0076", _egced)
		_eg.Log.Trace("\u0054\u0069\u006e\u0074 t\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _beeb.TintTransform)
		return nil, _egced
	}
	_eg.Log.Trace("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0043\u006f\u006c\u006fr\u0046\u0072\u006f\u006d\u0046\u006c\u006f\u0061\u0074\u0073\u0028\u0025\u002bv\u0029\u0020\u006f\u006e\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061te\u0053\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0023\u0076", _gcfgb, _beeb.AlternateSpace)
	_gddb, _egced := _beeb.AlternateSpace.ColorFromFloats(_gcfgb)
	if _egced != nil {
		_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u002c\u0020\u0066a\u0069\u006c\u0065d \u0074\u006f\u0020\u0065\u0076\u0061l\u0075\u0061\u0074\u0065\u0020\u0069\u006e\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u003a \u0025\u0076", _egced)
		return nil, _egced
	}
	return _gddb, nil
}

// SetCatalogMetadata sets the catalog metadata (XMP) stream object.
func (_feegg *PdfWriter) SetCatalogMetadata(meta _ebb.PdfObject) error {
	if meta == nil {
		_feegg._dffegd.Remove("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
		return nil
	}
	_abbee, _aegfe := _ebb.GetStream(meta)
	if !_aegfe {
		return _gf.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006d\u0065\u0074\u0061\u0064a\u0074\u0061\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0073t\u0072\u0065\u0061\u006d")
	}
	_feegg.addObject(_abbee)
	_feegg._dffegd.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _abbee)
	return nil
}
func (_cafeg *PdfPattern) getDict() *_ebb.PdfObjectDictionary {
	if _fdfaeb, _gfbbdd := _cafeg._dcddc.(*_ebb.PdfIndirectObject); _gfbbdd {
		_ccdc, _acebd := _fdfaeb.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_acebd {
			return nil
		}
		return _ccdc
	} else if _dbbbd, _fbcfe := _cafeg._dcddc.(*_ebb.PdfObjectStream); _fbcfe {
		return _dbbbd.PdfObjectDictionary
	} else {
		_eg.Log.Debug("\u0054r\u0079\u0069\u006e\u0067\u0020\u0074\u006f a\u0063\u0063\u0065\u0073\u0073\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063t \u0074\u0079\u0070e\u0020\u0028\u0025\u0054\u0029", _cafeg._dcddc)
		return nil
	}
}

// GetAsShadingPattern returns a shading pattern. Check with IsShading() prior to using this.
func (_dede *PdfPattern) GetAsShadingPattern() *PdfShadingPattern {
	return _dede._ffagg.(*PdfShadingPattern)
}

// NewPdfAnnotation3D returns a new 3d annotation.
func NewPdfAnnotation3D() *PdfAnnotation3D {
	_bec := NewPdfAnnotation()
	_fgad := &PdfAnnotation3D{}
	_fgad.PdfAnnotation = _bec
	_bec.SetContext(_fgad)
	return _fgad
}

// ToInteger convert to an integer format.
func (_eadf *PdfColorDeviceGray) ToInteger(bits int) uint32 {
	_eeeff := _cbg.Pow(2, float64(bits)) - 1
	return uint32(_eeeff * _eadf.Val())
}

// PdfPageResourcesColorspaces contains the colorspace in the PdfPageResources.
// Needs to have matching name and colorspace map entry. The Names define the order.
type PdfPageResourcesColorspaces struct {
	Names       []string
	Colorspaces map[string]PdfColorspace
	_ddffd      *_ebb.PdfIndirectObject
}

// PdfAnnotationMarkup represents additional fields for mark-up annotations.
// (Section 12.5.6.2 p. 399).
type PdfAnnotationMarkup struct {
	T            _ebb.PdfObject
	Popup        *PdfAnnotationPopup
	CA           _ebb.PdfObject
	RC           _ebb.PdfObject
	CreationDate _ebb.PdfObject
	IRT          _ebb.PdfObject
	Subj         _ebb.PdfObject
	RT           _ebb.PdfObject
	IT           _ebb.PdfObject
	ExData       _ebb.PdfObject
}

// UpdatePage updates the `page` in the new revision if it has changed.
func (_cecb *PdfAppender) UpdatePage(page *PdfPage) { _cecb.updateObjectsDeep(page.ToPdfObject(), nil) }
func _edfcb() _f.Time                               { _daddc.Lock(); defer _daddc.Unlock(); return _gecg }
func _abgfe(_gagf *_ebb.PdfObjectDictionary) bool {
	for _, _eefc := range _gagf.Keys() {
		if _, _ccee := _gdcdfg[_eefc.String()]; _ccee {
			return true
		}
	}
	return false
}
func (_eafg *PdfAcroForm) filteredFields(_dfgb FieldFilterFunc, _fbbeg bool) []*PdfField {
	if _eafg == nil {
		return nil
	}
	return _cgcga(_eafg.Fields, _dfgb, _fbbeg)
}

// ToPdfObject implements interface PdfModel.
func (_ebfg *PdfAnnotationLink) ToPdfObject() _ebb.PdfObject {
	_ebfg.PdfAnnotation.ToPdfObject()
	_ebad := _ebfg._bdcd
	_bedc := _ebad.PdfObject.(*_ebb.PdfObjectDictionary)
	_bedc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u004c\u0069\u006e\u006b"))
	if _ebfg._ffea != nil && _ebfg._ffea._ad != nil {
		_bedc.Set("\u0041", _ebfg._ffea._ad.ToPdfObject())
	} else if _ebfg.A != nil {
		_bedc.Set("\u0041", _ebfg.A)
	}
	_bedc.SetIfNotNil("\u0044\u0065\u0073\u0074", _ebfg.Dest)
	_bedc.SetIfNotNil("\u0048", _ebfg.H)
	_bedc.SetIfNotNil("\u0050\u0041", _ebfg.PA)
	_bedc.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _ebfg.QuadPoints)
	_bedc.SetIfNotNil("\u0042\u0053", _ebfg.BS)
	return _ebad
}

// GetAllContentStreams gets all the content streams for a page as one string.
func (_feefa *PdfPage) GetAllContentStreams() (string, error) {
	_aedcf, _eaaea := _feefa.GetContentStreams()
	if _eaaea != nil {
		return "", _eaaea
	}
	return _ee.Join(_aedcf, "\u0020"), nil
}

// PdfColorspaceDeviceNAttributes contains additional information about the components of colour space that
// conforming readers may use. Conforming readers need not use the alternateSpace and tintTransform parameters,
// and may instead use a custom blending algorithms, along with other information provided in the attributes
// dictionary if present.
type PdfColorspaceDeviceNAttributes struct {
	Subtype     *_ebb.PdfObjectName
	Colorants   _ebb.PdfObject
	Process     _ebb.PdfObject
	MixingHints _ebb.PdfObject
	_afca       *_ebb.PdfIndirectObject
}

// ToPdfObject implements interface PdfModel.
// Note: Call the sub-annotation's ToPdfObject to set both the generic and non-generic information.
func (_gggd *PdfAnnotation) ToPdfObject() _ebb.PdfObject {
	_bdcda := _gggd._bdcd
	_eegcd := _bdcda.PdfObject.(*_ebb.PdfObjectDictionary)
	_eegcd.Clear()
	_eegcd.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0041\u006e\u006eo\u0074"))
	_eegcd.SetIfNotNil("\u0052\u0065\u0063\u0074", _gggd.Rect)
	_eegcd.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _gggd.Contents)
	_eegcd.SetIfNotNil("\u0050", _gggd.P)
	_eegcd.SetIfNotNil("\u004e\u004d", _gggd.NM)
	_eegcd.SetIfNotNil("\u004d", _gggd.M)
	_eegcd.SetIfNotNil("\u0046", _gggd.F)
	_eegcd.SetIfNotNil("\u0041\u0050", _gggd.AP)
	_eegcd.SetIfNotNil("\u0041\u0053", _gggd.AS)
	_eegcd.SetIfNotNil("\u0042\u006f\u0072\u0064\u0065\u0072", _gggd.Border)
	_eegcd.SetIfNotNil("\u0043", _gggd.C)
	_eegcd.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _gggd.StructParent)
	_eegcd.SetIfNotNil("\u004f\u0043", _gggd.OC)
	return _bdcda
}
func _gggf(_bba _ebb.PdfObject) (*PdfFilespec, error) {
	if _bba == nil {
		return nil, nil
	}
	return NewPdfFilespecFromObj(_bba)
}

// PdfActionImportData represents a importData action.
type PdfActionImportData struct {
	*PdfAction
	F *PdfFilespec
}

func (_ebff *PdfReader) newPdfAnnotationRichMediaFromDict(_bbcf *_ebb.PdfObjectDictionary) (*PdfAnnotationRichMedia, error) {
	_cgda := &PdfAnnotationRichMedia{}
	_cgda.RichMediaSettings = _bbcf.Get("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073")
	_cgda.RichMediaContent = _bbcf.Get("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074")
	return _cgda, nil
}

// Evaluate runs the function. Input is [x1 x2 x3].
func (_cegd *PdfFunctionType4) Evaluate(xVec []float64) ([]float64, error) {
	if _cegd._bdbae == nil {
		_cegd._bdbae = _bc.NewPSExecutor(_cegd.Program)
	}
	var _bcce []_bc.PSObject
	for _, _fageaf := range xVec {
		_bcce = append(_bcce, _bc.MakeReal(_fageaf))
	}
	_dfced, _bgcge := _cegd._bdbae.Execute(_bcce)
	if _bgcge != nil {
		return nil, _bgcge
	}
	_bccbg, _bgcge := _bc.PSObjectArrayToFloat64Array(_dfced)
	if _bgcge != nil {
		return nil, _bgcge
	}
	return _bccbg, nil
}

// GetNumComponents returns the number of input color components, i.e. that are input to the tint transform.
func (_addf *PdfColorspaceDeviceN) GetNumComponents() int { return _addf.ColorantNames.Len() }

// AcroFormNeedsRepair returns true if the document contains widget annotations
// linked to fields which are not referenced in the AcroForm. The AcroForm can
// be repaired using the RepairAcroForm method of the reader.
func (_efgdb *PdfReader) AcroFormNeedsRepair() (bool, error) {
	var _cabb []*PdfField
	if _efgdb.AcroForm != nil {
		_cabb = _efgdb.AcroForm.AllFields()
	}
	_bgag := make(map[*PdfField]struct{}, len(_cabb))
	for _, _aggfa := range _cabb {
		_bgag[_aggfa] = struct{}{}
	}
	for _, _deagg := range _efgdb.PageList {
		_abgdc, _bccf := _deagg.GetAnnotations()
		if _bccf != nil {
			return false, _bccf
		}
		for _, _dcabe := range _abgdc {
			_egeb, _gedac := _dcabe.GetContext().(*PdfAnnotationWidget)
			if !_gedac {
				continue
			}
			_bgdgg := _egeb.Field()
			if _bgdgg == nil {
				return true, nil
			}
			if _, _fbgc := _bgag[_bgdgg]; !_fbgc {
				return true, nil
			}
		}
	}
	return false, nil
}

// PdfColorCalRGB represents a color in the Colorimetric CIE RGB colorspace.
// A, B, C components
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorCalRGB [3]float64

// NewPdfWriter initializes a new PdfWriter.
func NewPdfWriter() PdfWriter {
	_deaafa := PdfWriter{}
	_deaafa._ffffd = map[_ebb.PdfObject]struct{}{}
	_deaafa._ebdgg = []_ebb.PdfObject{}
	_deaafa._eefeb = map[_ebb.PdfObject][]*_ebb.PdfObjectDictionary{}
	_deaafa._dcfg = map[_ebb.PdfObject]struct{}{}
	_deaafa._efcge.Major = 1
	_deaafa._efcge.Minor = 3
	_afebb := _ebb.MakeDict()
	_gdec := []struct {
		_ddffa  _ebb.PdfObjectName
		_ecedfa string
	}{{
		"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _gagfgf()},
		{"\u0043r\u0065\u0061\u0074\u006f\u0072", _feaae()},
		{"\u0041\u0075\u0074\u0068\u006f\u0072", _dccgg()},
		{"\u0053u\u0062\u006a\u0065\u0063\u0074", _bdfea()},
		{"\u0054\u0069\u0074l\u0065", _bggeg()},
		{"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _fggcd()},
	}
	for _, _dggbbf := range _gdec {
		if _dggbbf._ecedfa != "" {
			_afebb.Set(_dggbbf._ddffa, _ebb.MakeString(_dggbbf._ecedfa))
		}
	}
	if _afccd := _edfcb(); !_afccd.IsZero() {
		if _cbgab, _eafac := NewPdfDateFromTime(_afccd); _eafac == nil {
			_afebb.Set("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _cbgab.ToPdfObject())
		}
	}
	if _abeeg := _eacbd(); !_abeeg.IsZero() {
		if _gdbed, _abagg := NewPdfDateFromTime(_abeeg); _abagg == nil {
			_afebb.Set("\u004do\u0064\u0044\u0061\u0074\u0065", _gdbed.ToPdfObject())
		}
	}
	_abde := _ebb.PdfIndirectObject{}
	_abde.PdfObject = _afebb
	_deaafa._eadfd = &_abde
	_deaafa.addObject(&_abde)
	_bdecc := _ebb.PdfIndirectObject{}
	_eafgf := _ebb.MakeDict()
	_eafgf.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0043a\u0074\u0061\u006c\u006f\u0067"))
	_bdecc.PdfObject = _eafgf
	_deaafa._gegba = &_bdecc
	_deaafa.addObject(_deaafa._gegba)
	_edcbca, _cbcb := _bafec("\u0077")
	if _cbcb != nil {
		_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cbcb)
	}
	_deaafa._cfecg = _edcbca
	_bcgaeb := _ebb.PdfIndirectObject{}
	_bcbcbb := _ebb.MakeDict()
	_bcbcbb.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0050\u0061\u0067e\u0073"))
	_fbfac := _ebb.PdfObjectArray{}
	_bcbcbb.Set("\u004b\u0069\u0064\u0073", &_fbfac)
	_bcbcbb.Set("\u0043\u006f\u0075n\u0074", _ebb.MakeInteger(0))
	_bcgaeb.PdfObject = _bcbcbb
	_deaafa._dggbf = &_bcgaeb
	_deaafa._afbdd = map[_ebb.PdfObject]struct{}{}
	_deaafa.addObject(_deaafa._dggbf)
	_eafgf.Set("\u0050\u0061\u0067e\u0073", &_bcgaeb)
	_deaafa._dffegd = _eafgf
	_eg.Log.Trace("\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0025\u0073", _bdecc)
	return _deaafa
}
func (_efg *PdfReader) newPdfActionFromIndirectObject(_fga *_ebb.PdfIndirectObject) (*PdfAction, error) {
	_bbe, _bccb := _fga.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_bccb {
		return nil, _bg.Errorf("\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u006e\u006f\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _efg._abbaca.GetModelFromPrimitive(_bbe); model != nil {
		_feac, _fbc := model.(*PdfAction)
		if !_fbc {
			return nil, _bg.Errorf("\u0063\u0061c\u0068\u0065\u0064\u0020\u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0061\u0063ti\u006f\u006e")
		}
		return _feac, nil
	}
	_dgf := &PdfAction{}
	_dgf._abe = _fga
	_efg._abbaca.Register(_bbe, _dgf)
	if _cfd := _bbe.Get("\u0054\u0079\u0070\u0065"); _cfd != nil {
		_gbd, _bade := _cfd.(*_ebb.PdfObjectName)
		if !_bade {
			_eg.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _cfd)
		} else {
			if *_gbd != "\u0041\u0063\u0074\u0069\u006f\u006e" {
				_eg.Log.Trace("\u0055\u006e\u0073u\u0073\u0070\u0065\u0063t\u0065\u0064\u0020\u0054\u0079\u0070\u0065 \u0021\u003d\u0020\u0041\u0063\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", *_gbd)
			}
			_dgf.Type = _gbd
		}
	}
	if _fgbd := _bbe.Get("\u004e\u0065\u0078\u0074"); _fgbd != nil {
		_dgf.Next = _fgbd
	}
	if _faad := _bbe.Get("\u0053"); _faad != nil {
		_dgf.S = _faad
	}
	_edef, _gfb := _dgf.S.(*_ebb.PdfObjectName)
	if !_gfb {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _dgf.S)
		return nil, _bg.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0053\u0020\u006f\u0062\u006a\u0065c\u0074 \u0074y\u0070e\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _dgf.S)
	}
	_fbda := PdfActionType(_edef.String())
	switch _fbda {
	case ActionTypeGoTo:
		_gba, _fde := _efg.newPdfActionGotoFromDict(_bbe)
		if _fde != nil {
			return nil, _fde
		}
		_gba.PdfAction = _dgf
		_dgf._ad = _gba
		return _dgf, nil
	case ActionTypeGoToR:
		_dbg, _gga := _efg.newPdfActionGotoRFromDict(_bbe)
		if _gga != nil {
			return nil, _gga
		}
		_dbg.PdfAction = _dgf
		_dgf._ad = _dbg
		return _dgf, nil
	case ActionTypeGoToE:
		_gdcd, _eed := _efg.newPdfActionGotoEFromDict(_bbe)
		if _eed != nil {
			return nil, _eed
		}
		_gdcd.PdfAction = _dgf
		_dgf._ad = _gdcd
		return _dgf, nil
	case ActionTypeLaunch:
		_dgg, _bce := _efg.newPdfActionLaunchFromDict(_bbe)
		if _bce != nil {
			return nil, _bce
		}
		_dgg.PdfAction = _dgf
		_dgf._ad = _dgg
		return _dgf, nil
	case ActionTypeThread:
		_bag, _fdf := _efg.newPdfActionThreadFromDict(_bbe)
		if _fdf != nil {
			return nil, _fdf
		}
		_bag.PdfAction = _dgf
		_dgf._ad = _bag
		return _dgf, nil
	case ActionTypeURI:
		_dbed, _eecf := _efg.newPdfActionURIFromDict(_bbe)
		if _eecf != nil {
			return nil, _eecf
		}
		_dbed.PdfAction = _dgf
		_dgf._ad = _dbed
		return _dgf, nil
	case ActionTypeSound:
		_fgbg, _bgb := _efg.newPdfActionSoundFromDict(_bbe)
		if _bgb != nil {
			return nil, _bgb
		}
		_fgbg.PdfAction = _dgf
		_dgf._ad = _fgbg
		return _dgf, nil
	case ActionTypeMovie:
		_ggg, _gfcf := _efg.newPdfActionMovieFromDict(_bbe)
		if _gfcf != nil {
			return nil, _gfcf
		}
		_ggg.PdfAction = _dgf
		_dgf._ad = _ggg
		return _dgf, nil
	case ActionTypeHide:
		_eba, _ggac := _efg.newPdfActionHideFromDict(_bbe)
		if _ggac != nil {
			return nil, _ggac
		}
		_eba.PdfAction = _dgf
		_dgf._ad = _eba
		return _dgf, nil
	case ActionTypeNamed:
		_cdc, _gbea := _efg.newPdfActionNamedFromDict(_bbe)
		if _gbea != nil {
			return nil, _gbea
		}
		_cdc.PdfAction = _dgf
		_dgf._ad = _cdc
		return _dgf, nil
	case ActionTypeSubmitForm:
		_bbg, _cccb := _efg.newPdfActionSubmitFormFromDict(_bbe)
		if _cccb != nil {
			return nil, _cccb
		}
		_bbg.PdfAction = _dgf
		_dgf._ad = _bbg
		return _dgf, nil
	case ActionTypeResetForm:
		_caf, _ggb := _efg.newPdfActionResetFormFromDict(_bbe)
		if _ggb != nil {
			return nil, _ggb
		}
		_caf.PdfAction = _dgf
		_dgf._ad = _caf
		return _dgf, nil
	case ActionTypeImportData:
		_efa, _bdd := _efg.newPdfActionImportDataFromDict(_bbe)
		if _bdd != nil {
			return nil, _bdd
		}
		_efa.PdfAction = _dgf
		_dgf._ad = _efa
		return _dgf, nil
	case ActionTypeSetOCGState:
		_dedg, _edg := _efg.newPdfActionSetOCGStateFromDict(_bbe)
		if _edg != nil {
			return nil, _edg
		}
		_dedg.PdfAction = _dgf
		_dgf._ad = _dedg
		return _dgf, nil
	case ActionTypeRendition:
		_fgde, _ag := _efg.newPdfActionRenditionFromDict(_bbe)
		if _ag != nil {
			return nil, _ag
		}
		_fgde.PdfAction = _dgf
		_dgf._ad = _fgde
		return _dgf, nil
	case ActionTypeTrans:
		_egb, _ebg := _efg.newPdfActionTransFromDict(_bbe)
		if _ebg != nil {
			return nil, _ebg
		}
		_egb.PdfAction = _dgf
		_dgf._ad = _egb
		return _dgf, nil
	case ActionTypeGoTo3DView:
		_ddd, _cde := _efg.newPdfActionGoTo3DViewFromDict(_bbe)
		if _cde != nil {
			return nil, _cde
		}
		_ddd.PdfAction = _dgf
		_dgf._ad = _ddd
		return _dgf, nil
	case ActionTypeJavaScript:
		_cec, _fac := _efg.newPdfActionJavaScriptFromDict(_bbe)
		if _fac != nil {
			return nil, _fac
		}
		_cec.PdfAction = _dgf
		_dgf._ad = _cec
		return _dgf, nil
	}
	_eg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006eg\u0020u\u006ek\u006eo\u0077\u006e\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073", _fbda)
	return nil, nil
}

// SetType sets the field button's type.  Can be one of:
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
// This sets the field's flag appropriately.
func (_bafdbb *PdfFieldButton) SetType(btype ButtonType) {
	_gdbf := uint32(0)
	if _bafdbb.Ff != nil {
		_gdbf = uint32(*_bafdbb.Ff)
	}
	switch btype {
	case ButtonTypePush:
		_gdbf |= FieldFlagPushbutton.Mask()
	case ButtonTypeRadio:
		_gdbf |= FieldFlagRadio.Mask()
	}
	_bafdbb.Ff = _ebb.MakeInteger(int64(_gdbf))
}
func (_cgcab *fontFile) loadFromSegments(_ffac, _edbga []byte) error {
	_eg.Log.Trace("\u006c\u006f\u0061dF\u0072\u006f\u006d\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u0020\u0025\u0064", len(_ffac), len(_edbga))
	_bgba := _cgcab.parseASCIIPart(_ffac)
	if _bgba != nil {
		return _bgba
	}
	_eg.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _cgcab)
	if len(_edbga) == 0 {
		return nil
	}
	_eg.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _cgcab)
	return nil
}
func (_gfff *PdfColorspaceICCBased) String() string {
	return "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_beggf *PdfShadingType2) ToPdfObject() _ebb.PdfObject {
	_beggf.PdfShading.ToPdfObject()
	_dabg, _aface := _beggf.getShadingDict()
	if _aface != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _dabg == nil {
		_eg.Log.Error("\u0053\u0068\u0061\u0064in\u0067\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		return nil
	}
	if _beggf.Coords != nil {
		_dabg.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _beggf.Coords)
	}
	if _beggf.Domain != nil {
		_dabg.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _beggf.Domain)
	}
	if _beggf.Function != nil {
		if len(_beggf.Function) == 1 {
			_dabg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _beggf.Function[0].ToPdfObject())
		} else {
			_cdcab := _ebb.MakeArray()
			for _, _gdgb := range _beggf.Function {
				_cdcab.Append(_gdgb.ToPdfObject())
			}
			_dabg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _cdcab)
		}
	}
	if _beggf.Extend != nil {
		_dabg.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _beggf.Extend)
	}
	return _beggf._fbfae
}
func (_edbd *PdfAnnotation) String() string {
	_aac := ""
	_ggbg, _fdg := _edbd.ToPdfObject().(*_ebb.PdfIndirectObject)
	if _fdg {
		_aac = _bg.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _edbd._efd, _ggbg.PdfObject.String())
	}
	return _aac
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

// ToPdfObject returns the choice field dictionary within an indirect object (container).
func (_cdbdg *PdfFieldChoice) ToPdfObject() _ebb.PdfObject {
	_cdbdg.PdfField.ToPdfObject()
	_ffeg := _cdbdg._cdfd
	_gbacb := _ffeg.PdfObject.(*_ebb.PdfObjectDictionary)
	_gbacb.Set("\u0046\u0054", _ebb.MakeName("\u0043\u0068"))
	if _cdbdg.Opt != nil {
		_gbacb.Set("\u004f\u0070\u0074", _cdbdg.Opt)
	}
	if _cdbdg.TI != nil {
		_gbacb.Set("\u0054\u0049", _cdbdg.TI)
	}
	if _cdbdg.I != nil {
		_gbacb.Set("\u0049", _cdbdg.I)
	}
	return _ffeg
}

// Has checks if flag fl is set in flag and returns true if so, false otherwise.
func (_cagg FieldFlag) Has(fl FieldFlag) bool { return (_cagg.Mask() & fl.Mask()) > 0 }

// PdfAnnotationRichMedia represents Rich Media annotations.
type PdfAnnotationRichMedia struct {
	*PdfAnnotation
	RichMediaSettings _ebb.PdfObject
	RichMediaContent  _ebb.PdfObject
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for an RGB device.
func (_gea *PdfColorspaceDeviceRGB) GetNumComponents() int { return 3 }

// StandardApplier is the interface that performs optimization of the whole PDF document.
// As a result an input document is being changed by the optimizer.
// The writer than takes back all it's parts and overwrites it.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
type StandardApplier interface {
	ApplyStandard(_gegab *_bda.Document) error
}

// ImageToRGB returns an error since an image cannot be defined in a pattern colorspace.
func (_ddbc *PdfColorspaceSpecialPattern) ImageToRGB(img Image) (Image, error) {
	_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0061\u0074\u0074\u0065\u0072n \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	return img, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0066\u006f\u0072\u0020\u0069m\u0061\u0067\u0065\u0020\u0028p\u0061\u0074t\u0065\u0072\u006e\u0029")
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_ecbf *PdfFunctionType0) Evaluate(x []float64) ([]float64, error) {
	if len(x) != _ecbf.NumInputs {
		_eg.Log.Error("\u004eu\u006d\u0062e\u0072\u0020\u006f\u0066 \u0069\u006e\u0070u\u0074\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061tc\u0068\u0069\u006eg\u0020\u0077h\u0061\u0074\u0020\u0069\u0073\u0020n\u0065\u0065d\u0065\u0064")
		return nil, _gf.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _ecbf._egdg == nil {
		_gcda := _ecbf.processSamples()
		if _gcda != nil {
			return nil, _gcda
		}
	}
	_edaa := _ecbf.Encode
	if _edaa == nil {
		_edaa = []float64{}
		for _eadbf := 0; _eadbf < len(_ecbf.Size); _eadbf++ {
			_edaa = append(_edaa, 0)
			_edaa = append(_edaa, float64(_ecbf.Size[_eadbf]-1))
		}
	}
	_gafe := _ecbf.Decode
	if _gafe == nil {
		_gafe = _ecbf.Range
	}
	_bcbec := make([]int, len(x))
	for _dcde := 0; _dcde < len(x); _dcde++ {
		_dccd := x[_dcde]
		_ffgec := _cbg.Min(_cbg.Max(_dccd, _ecbf.Domain[2*_dcde]), _ecbf.Domain[2*_dcde+1])
		_bgegd := _dg.LinearInterpolate(_ffgec, _ecbf.Domain[2*_dcde], _ecbf.Domain[2*_dcde+1], _edaa[2*_dcde], _edaa[2*_dcde+1])
		_gbdcd := _cbg.Min(_cbg.Max(_bgegd, 0), float64(_ecbf.Size[_dcde]-1))
		_dcec := int(_cbg.Floor(_gbdcd + 0.5))
		if _dcec < 0 {
			_dcec = 0
		} else if _dcec > _ecbf.Size[_dcde] {
			_dcec = _ecbf.Size[_dcde] - 1
		}
		_bcbec[_dcde] = _dcec
	}
	_gccad := _bcbec[0]
	for _dbef := 1; _dbef < _ecbf.NumInputs; _dbef++ {
		_eeddc := _bcbec[_dbef]
		for _bbefe := 0; _bbefe < _dbef; _bbefe++ {
			_eeddc *= _ecbf.Size[_bbefe]
		}
		_gccad += _eeddc
	}
	_gccad *= _ecbf.NumOutputs
	var _cfeg []float64
	for _ggdde := 0; _ggdde < _ecbf.NumOutputs; _ggdde++ {
		_fbgfe := _gccad + _ggdde
		if _fbgfe >= len(_ecbf._egdg) {
			_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u006e\u006ft\u0020\u0065\u006eo\u0075\u0067\u0068\u0020\u0069\u006ep\u0075\u0074\u0020sa\u006dp\u006c\u0065\u0073\u0020\u0074\u006f\u0020d\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0076\u0061lu\u0065\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			continue
		}
		_gaaga := _ecbf._egdg[_fbgfe]
		_baed := _dg.LinearInterpolate(float64(_gaaga), 0, _cbg.Pow(2, float64(_ecbf.BitsPerSample)), _gafe[2*_ggdde], _gafe[2*_ggdde+1])
		_ebadg := _cbg.Min(_cbg.Max(_baed, _ecbf.Range[2*_ggdde]), _ecbf.Range[2*_ggdde+1])
		_cfeg = append(_cfeg, _ebadg)
	}
	return _cfeg, nil
}

// String implements interface PdfObject.
func (_bdg *PdfAction) String() string {
	_cd, _badd := _bdg.ToPdfObject().(*_ebb.PdfIndirectObject)
	if _badd {
		return _bg.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _bdg._ad, _cd.PdfObject.String())
	}
	return ""
}

// EnableAll LTV enables all signatures in the PDF document.
// The signing certificate chain is extracted from each signature dictionary.
// Optionally, additional certificates can be specified through the
// `extraCerts` parameter. The LTV client attempts to build the certificate
// chain up to a trusted root by downloading any missing certificates.
func (_gdac *LTV) EnableAll(extraCerts []*_g.Certificate) error {
	_aeff := _gdac._ggdbg._acfe.AcroForm
	for _, _gebc := range _aeff.AllFields() {
		_fgbcea, _ := _gebc.GetContext().(*PdfFieldSignature)
		if _fgbcea == nil {
			continue
		}
		_gcdgb := _fgbcea.V
		if _bgbcc := _gdac.validateSig(_gcdgb); _bgbcc != nil {
			_eg.Log.Debug("\u0057\u0041\u0052N\u003a\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _bgbcc)
		}
		if _eecb := _gdac.Enable(_gcdgb, extraCerts); _eecb != nil {
			return _eecb
		}
	}
	return nil
}
func (_acf *PdfReader) newPdfActionSubmitFormFromDict(_fad *_ebb.PdfObjectDictionary) (*PdfActionSubmitForm, error) {
	_aba, _ebc := _gggf(_fad.Get("\u0046"))
	if _ebc != nil {
		return nil, _ebc
	}
	return &PdfActionSubmitForm{F: _aba, Fields: _fad.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _fad.Get("\u0046\u006c\u0061g\u0073")}, nil
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
	_fcgb  *_ebb.PdfIndirectObject
	Certs  []*_ebb.PdfObjectStream
	OCSPs  []*_ebb.PdfObjectStream
	CRLs   []*_ebb.PdfObjectStream
	VRI    map[string]*VRI
	_aeag  map[string]*_ebb.PdfObjectStream
	_cadd  map[string]*_ebb.PdfObjectStream
	_fafgb map[string]*_ebb.PdfObjectStream
}

// PageProcessCallback callback function used in page loading
// that could be used to modify the page content.
//
// If an error is returned, the `ToWriter` process would fail.
//
// This callback, if defined, will take precedence over `PageCallback` callback.
type PageProcessCallback func(_aaefe int, _babd *PdfPage) error

// ToPdfObject implements interface PdfModel.
func (_ddb *PdfAnnotationFreeText) ToPdfObject() _ebb.PdfObject {
	_ddb.PdfAnnotation.ToPdfObject()
	_ccf := _ddb._bdcd
	_gbacg := _ccf.PdfObject.(*_ebb.PdfObjectDictionary)
	_ddb.PdfAnnotationMarkup.appendToPdfDictionary(_gbacg)
	_gbacg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074"))
	_gbacg.SetIfNotNil("\u0044\u0041", _ddb.DA)
	_gbacg.SetIfNotNil("\u0051", _ddb.Q)
	_gbacg.SetIfNotNil("\u0052\u0043", _ddb.RC)
	_gbacg.SetIfNotNil("\u0044\u0053", _ddb.DS)
	_gbacg.SetIfNotNil("\u0043\u004c", _ddb.CL)
	_gbacg.SetIfNotNil("\u0049\u0054", _ddb.IT)
	_gbacg.SetIfNotNil("\u0042\u0045", _ddb.BE)
	_gbacg.SetIfNotNil("\u0052\u0044", _ddb.RD)
	_gbacg.SetIfNotNil("\u0042\u0053", _ddb.BS)
	_gbacg.SetIfNotNil("\u004c\u0045", _ddb.LE)
	return _ccf
}

// ToPdfObject returns the button field dictionary within an indirect object.
func (_ccgcc *PdfFieldButton) ToPdfObject() _ebb.PdfObject {
	_ccgcc.PdfField.ToPdfObject()
	_bdfcc := _ccgcc._cdfd
	_ccceg := _bdfcc.PdfObject.(*_ebb.PdfObjectDictionary)
	_ccceg.Set("\u0046\u0054", _ebb.MakeName("\u0042\u0074\u006e"))
	if _ccgcc.Opt != nil {
		_ccceg.Set("\u004f\u0070\u0074", _ccgcc.Opt)
	}
	return _bdfcc
}

// Decrypt decrypts the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful,
// false otherwise.
func (_bbbf *PdfReader) Decrypt(password []byte) (bool, error) {
	_ccggb, _aeee := _bbbf._cafdf.Decrypt(password)
	if _aeee != nil {
		return false, _aeee
	}
	if !_ccggb {
		return false, nil
	}
	_aeee = _bbbf.loadStructure()
	if _aeee != nil {
		_eg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f \u006co\u0061d\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", _aeee)
		return false, _aeee
	}
	return true, nil
}

// Insert adds an outline item as a child of the current outline item,
// at the specified index.
func (_cggfb *OutlineItem) Insert(index uint, item *OutlineItem) {
	_dfgcc := uint(len(_cggfb.Entries))
	if index > _dfgcc {
		index = _dfgcc
	}
	_cggfb.Entries = append(_cggfb.Entries[:index], append([]*OutlineItem{item}, _cggfb.Entries[index:]...)...)
}

// GetPageDict converts the Page to a PDF object dictionary.
func (_ddbbd *PdfPage) GetPageDict() *_ebb.PdfObjectDictionary {
	_fabgc := _ddbbd._cdbfde
	_fabgc.Clear()
	_fabgc.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0050\u0061\u0067\u0065"))
	_fabgc.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _ddbbd.Parent)
	if _ddbbd.LastModified != nil {
		_fabgc.Set("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _ddbbd.LastModified.ToPdfObject())
	}
	if _ddbbd.Resources != nil {
		_fabgc.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _ddbbd.Resources.ToPdfObject())
	}
	if _ddbbd.CropBox != nil {
		_fabgc.Set("\u0043r\u006f\u0070\u0042\u006f\u0078", _ddbbd.CropBox.ToPdfObject())
	}
	if _ddbbd.MediaBox != nil {
		_fabgc.Set("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", _ddbbd.MediaBox.ToPdfObject())
	}
	if _ddbbd.BleedBox != nil {
		_fabgc.Set("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078", _ddbbd.BleedBox.ToPdfObject())
	}
	if _ddbbd.TrimBox != nil {
		_fabgc.Set("\u0054r\u0069\u006d\u0042\u006f\u0078", _ddbbd.TrimBox.ToPdfObject())
	}
	if _ddbbd.ArtBox != nil {
		_fabgc.Set("\u0041\u0072\u0074\u0042\u006f\u0078", _ddbbd.ArtBox.ToPdfObject())
	}
	_fabgc.SetIfNotNil("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f", _ddbbd.BoxColorInfo)
	_fabgc.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _ddbbd.Contents)
	if _ddbbd.Rotate != nil {
		_fabgc.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _ebb.MakeInteger(*_ddbbd.Rotate))
	}
	_fabgc.SetIfNotNil("\u0047\u0072\u006fu\u0070", _ddbbd.Group)
	_fabgc.SetIfNotNil("\u0054\u0068\u0075m\u0062", _ddbbd.Thumb)
	_fabgc.SetIfNotNil("\u0042", _ddbbd.B)
	_fabgc.SetIfNotNil("\u0044\u0075\u0072", _ddbbd.Dur)
	_fabgc.SetIfNotNil("\u0054\u0072\u0061n\u0073", _ddbbd.Trans)
	_fabgc.SetIfNotNil("\u0041\u0041", _ddbbd.AA)
	_fabgc.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _ddbbd.Metadata)
	_fabgc.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _ddbbd.PieceInfo)
	_fabgc.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _ddbbd.StructParents)
	_fabgc.SetIfNotNil("\u0049\u0044", _ddbbd.ID)
	_fabgc.SetIfNotNil("\u0050\u005a", _ddbbd.PZ)
	_fabgc.SetIfNotNil("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f", _ddbbd.SeparationInfo)
	_fabgc.SetIfNotNil("\u0054\u0061\u0062\u0073", _ddbbd.Tabs)
	_fabgc.SetIfNotNil("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064", _ddbbd.TemplateInstantiated)
	_fabgc.SetIfNotNil("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s", _ddbbd.PresSteps)
	_fabgc.SetIfNotNil("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074", _ddbbd.UserUnit)
	_fabgc.SetIfNotNil("\u0056\u0050", _ddbbd.VP)
	if _ddbbd._bbfed != nil {
		_deecf := _ebb.MakeArray()
		for _, _cfcfg := range _ddbbd._bbfed {
			if _dcfdca := _cfcfg.GetContext(); _dcfdca != nil {
				_deecf.Append(_dcfdca.ToPdfObject())
			} else {
				_deecf.Append(_cfcfg.ToPdfObject())
			}
		}
		if _deecf.Len() > 0 {
			_fabgc.Set("\u0041\u006e\u006e\u006f\u0074\u0073", _deecf)
		}
	} else if _ddbbd.Annots != nil {
		_fabgc.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0073", _ddbbd.Annots)
	}
	return _fabgc
}
func _gbgb(_gffdd _ebb.PdfObject) (*PdfColorspaceDeviceNAttributes, error) {
	_dgcca := &PdfColorspaceDeviceNAttributes{}
	var _gece *_ebb.PdfObjectDictionary
	switch _ccgb := _gffdd.(type) {
	case *_ebb.PdfIndirectObject:
		_dgcca._afca = _ccgb
		var _dcbca bool
		_gece, _dcbca = _ccgb.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_dcbca {
			_eg.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
			return nil, _gf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	case *_ebb.PdfObjectDictionary:
		_gece = _ccgb
	case *_ebb.PdfObjectReference:
		_accgb := _ccgb.Resolve()
		return _gbgb(_accgb)
	default:
		_eg.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
		return nil, _gf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _gfcb := _gece.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _gfcb != nil {
		_fcee, _bage := _ebb.TraceToDirectObject(_gfcb).(*_ebb.PdfObjectName)
		if !_bage {
			_eg.Log.Error("\u0044\u0065vi\u0063\u0065\u004e \u0061\u0074\u0074\u0072ibu\u0074e \u0053\u0075\u0062\u0074\u0079\u0070\u0065 t\u0079\u0070\u0065\u0020\u0065\u0072\u0072o\u0072")
			return nil, _gf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_dgcca.Subtype = _fcee
	}
	if _abfb := _gece.Get("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s"); _abfb != nil {
		_dgcca.Colorants = _abfb
	}
	if _cedgb := _gece.Get("\u0050r\u006f\u0063\u0065\u0073\u0073"); _cedgb != nil {
		_dgcca.Process = _cedgb
	}
	if _afeb := _gece.Get("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073"); _afeb != nil {
		_dgcca.MixingHints = _afeb
	}
	return _dgcca, nil
}

// NewCompositePdfFontFromTTFFile loads a composite font from a TTF font file. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTFFile.
func NewCompositePdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_aabcf, _ccgg := _ed.Open(filePath)
	if _ccgg != nil {
		_eg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006f\u0070\u0065\u006e\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _ccgg)
		return nil, _ccgg
	}
	defer _aabcf.Close()
	return NewCompositePdfFontFromTTF(_aabcf)
}

// NewDSS returns a new DSS dictionary.
func NewDSS() *DSS {
	return &DSS{_fcgb: _ebb.MakeIndirectObject(_ebb.MakeDict()), VRI: map[string]*VRI{}}
}

// String returns a string representation of what flags are set.
func (_cafgf FieldFlag) String() string {
	_edda := ""
	if _cafgf == FieldFlagClear {
		_edda = "\u0043\u006c\u0065a\u0072"
		return _edda
	}
	if _cafgf&FieldFlagReadOnly > 0 {
		_edda += "\u007cR\u0065\u0061\u0064\u004f\u006e\u006cy"
	}
	if _cafgf&FieldFlagRequired > 0 {
		_edda += "\u007cR\u0065\u0061\u0064\u004f\u006e\u006cy"
	}
	if _cafgf&FieldFlagNoExport > 0 {
		_edda += "\u007cN\u006f\u0045\u0078\u0070\u006f\u0072t"
	}
	if _cafgf&FieldFlagNoToggleToOff > 0 {
		_edda += "\u007c\u004e\u006f\u0054\u006f\u0067\u0067\u006c\u0065T\u006f\u004f\u0066\u0066"
	}
	if _cafgf&FieldFlagRadio > 0 {
		_edda += "\u007c\u0052\u0061\u0064\u0069\u006f"
	}
	if _cafgf&FieldFlagPushbutton > 0 {
		_edda += "|\u0050\u0075\u0073\u0068\u0062\u0075\u0074\u0074\u006f\u006e"
	}
	if _cafgf&FieldFlagRadiosInUnision > 0 {
		_edda += "\u007c\u0052a\u0064\u0069\u006fs\u0049\u006e\u0055\u006e\u0069\u0073\u0069\u006f\u006e"
	}
	if _cafgf&FieldFlagMultiline > 0 {
		_edda += "\u007c\u004d\u0075\u006c\u0074\u0069\u006c\u0069\u006e\u0065"
	}
	if _cafgf&FieldFlagPassword > 0 {
		_edda += "\u007cP\u0061\u0073\u0073\u0077\u006f\u0072d"
	}
	if _cafgf&FieldFlagFileSelect > 0 {
		_edda += "|\u0046\u0069\u006c\u0065\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _cafgf&FieldFlagDoNotScroll > 0 {
		_edda += "\u007c\u0044\u006fN\u006f\u0074\u0053\u0063\u0072\u006f\u006c\u006c"
	}
	if _cafgf&FieldFlagComb > 0 {
		_edda += "\u007c\u0043\u006fm\u0062"
	}
	if _cafgf&FieldFlagRichText > 0 {
		_edda += "\u007cR\u0069\u0063\u0068\u0054\u0065\u0078t"
	}
	if _cafgf&FieldFlagDoNotSpellCheck > 0 {
		_edda += "\u007c\u0044o\u004e\u006f\u0074S\u0070\u0065\u006c\u006c\u0043\u0068\u0065\u0063\u006b"
	}
	if _cafgf&FieldFlagCombo > 0 {
		_edda += "\u007c\u0043\u006f\u006d\u0062\u006f"
	}
	if _cafgf&FieldFlagEdit > 0 {
		_edda += "\u007c\u0045\u0064i\u0074"
	}
	if _cafgf&FieldFlagSort > 0 {
		_edda += "\u007c\u0053\u006fr\u0074"
	}
	if _cafgf&FieldFlagMultiSelect > 0 {
		_edda += "\u007c\u004d\u0075l\u0074\u0069\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _cafgf&FieldFlagCommitOnSelChange > 0 {
		_edda += "\u007cC\u006fm\u006d\u0069\u0074\u004f\u006eS\u0065\u006cC\u0068\u0061\u006e\u0067\u0065"
	}
	return _ee.Trim(_edda, "\u007c")
}
func (_gbada *pdfFontType0) baseFields() *fontCommon { return &_gbada.fontCommon }

type modelManager struct {
	_bgfdb   map[PdfModel]_ebb.PdfObject
	_aabcbdd map[_ebb.PdfObject]PdfModel
}

var _ pdfFont = (*pdfFontType0)(nil)

// AppendContentBytes creates a PDF stream from `cs` and appends it to the
// array of streams specified by the pages's Contents entry.
// If `wrapContents` is true, the content stream of the page is wrapped using
// a `q/Q` operator pair, so that its state does not affect the appended
// content stream.
func (_cccfg *PdfPage) AppendContentBytes(cs []byte, wrapContents bool) error {
	_cdfff := _cccfg.GetContentStreamObjs()
	wrapContents = wrapContents && len(_cdfff) > 0
	_aecde := _ebb.NewFlateEncoder()
	_cdfcd := _ebb.MakeArray()
	if wrapContents {
		_dagae, _fdegd := _ebb.MakeStream([]byte("\u0071\u000a"), _aecde)
		if _fdegd != nil {
			return _fdegd
		}
		_cdfcd.Append(_dagae)
	}
	_cdfcd.Append(_cdfff...)
	if wrapContents {
		_afdca, _cccgf := _ebb.MakeStream([]byte("\u000a\u0051\u000a"), _aecde)
		if _cccgf != nil {
			return _cccgf
		}
		_cdfcd.Append(_afdca)
	}
	_cfge, _dfdcd := _ebb.MakeStream(cs, _aecde)
	if _dfdcd != nil {
		return _dfdcd
	}
	_cdfcd.Append(_cfge)
	_cccfg.Contents = _cdfcd
	return nil
}

// GetRuneMetrics returns the character metrics for the rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_babc pdfFontSimple) GetRuneMetrics(r rune) (_bad.CharMetrics, bool) {
	if _babc._ddgd != nil {
		_edbc, _edgee := _babc._ddgd.Read(r)
		if _edgee {
			return _edbc, true
		}
	}
	_fdbe := _babc.Encoder()
	if _fdbe == nil {
		_eg.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _babc)
		return _bad.CharMetrics{}, false
	}
	_dcac, _cdfga := _fdbe.RuneToCharcode(r)
	if !_cdfga {
		if r != ' ' {
			_eg.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _babc)
		}
		return _bad.CharMetrics{}, false
	}
	_ffaga, _eacb := _babc.GetCharMetrics(_dcac)
	return _ffaga, _eacb
}

// NewOutlineBookmark returns an initialized PdfOutlineItem for a given bookmark title and page.
func NewOutlineBookmark(title string, page *_ebb.PdfIndirectObject) *PdfOutlineItem {
	_ddcg := PdfOutlineItem{}
	_ddcg._geeee = &_ddcg
	_ddcg.Title = _ebb.MakeString(title)
	_cfbgg := _ebb.MakeArray()
	_cfbgg.Append(page)
	_cfbgg.Append(_ebb.MakeName("\u0046\u0069\u0074"))
	_ddcg.Dest = _cfbgg
	return &_ddcg
}

// ToPdfOutlineItem returns a low level PdfOutlineItem object,
// based on the current instance.
func (_bcbfa *OutlineItem) ToPdfOutlineItem() (*PdfOutlineItem, int64) {
	_cbec := NewPdfOutlineItem()
	_cbec.Title = _ebb.MakeEncodedString(_bcbfa.Title, true)
	_cbec.Dest = _bcbfa.Dest.ToPdfObject()
	var _ffcf []*PdfOutlineItem
	var _gbcgg int64
	var _ecceb *PdfOutlineItem
	for _, _bcbcd := range _bcbfa.Entries {
		_bbcef, _dggcb := _bcbcd.ToPdfOutlineItem()
		_bbcef.Parent = &_cbec.PdfOutlineTreeNode
		if _ecceb != nil {
			_ecceb.Next = &_bbcef.PdfOutlineTreeNode
			_bbcef.Prev = &_ecceb.PdfOutlineTreeNode
		}
		_ffcf = append(_ffcf, _bbcef)
		_gbcgg += _dggcb
		_ecceb = _bbcef
	}
	_agbd := len(_ffcf)
	_gbcgg += int64(_agbd)
	if _agbd > 0 {
		_cbec.First = &_ffcf[0].PdfOutlineTreeNode
		_cbec.Last = &_ffcf[_agbd-1].PdfOutlineTreeNode
		_cbec.Count = &_gbcgg
	}
	return _cbec, _gbcgg
}

type pdfFont interface {
	_bad.Font

	// ToPdfObject returns a PDF representation of the font and implements interface Model.
	ToPdfObject() _ebb.PdfObject
	getFontDescriptor() *PdfFontDescriptor
	baseFields() *fontCommon
}

// ToPdfObject implements interface PdfModel.
func (_bgc *PdfAnnotationScreen) ToPdfObject() _ebb.PdfObject {
	_bgc.PdfAnnotation.ToPdfObject()
	_ceee := _bgc._bdcd
	_fdgg := _ceee.PdfObject.(*_ebb.PdfObjectDictionary)
	_fdgg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0053\u0063\u0072\u0065\u0065\u006e"))
	_fdgg.SetIfNotNil("\u0054", _bgc.T)
	_fdgg.SetIfNotNil("\u004d\u004b", _bgc.MK)
	_fdgg.SetIfNotNil("\u0041", _bgc.A)
	_fdgg.SetIfNotNil("\u0041\u0041", _bgc.AA)
	return _ceee
}
func (_fgadb *DSS) add(_gbaca *[]*_ebb.PdfObjectStream, _cgbe map[string]*_ebb.PdfObjectStream, _edad [][]byte) ([]*_ebb.PdfObjectStream, error) {
	_gaaec := make([]*_ebb.PdfObjectStream, 0, len(_edad))
	for _, _gfeba := range _edad {
		_ggec, _bbbb := _eaef(_gfeba)
		if _bbbb != nil {
			return nil, _bbbb
		}
		_cgbb, _dacc := _cgbe[string(_ggec)]
		if !_dacc {
			_cgbb, _bbbb = _ebb.MakeStream(_gfeba, _ebb.NewRawEncoder())
			if _bbbb != nil {
				return nil, _bbbb
			}
			_cgbe[string(_ggec)] = _cgbb
			*_gbaca = append(*_gbaca, _cgbb)
		}
		_gaaec = append(_gaaec, _cgbb)
	}
	return _gaaec, nil
}

// String returns a string representation of the field.
func (_fgcfa *PdfField) String() string {
	if _aaag, _ffdf := _fgcfa.ToPdfObject().(*_ebb.PdfIndirectObject); _ffdf {
		return _bg.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _fgcfa._cada, _aaag.PdfObject.String())
	}
	return ""
}

// Subtype returns the font's "Subtype" field.
func (_dacdcd *PdfFont) Subtype() string {
	_cdeg := _dacdcd.baseFields()._dfbf
	if _bddc, _cdgc := _dacdcd._ebcad.(*pdfFontType0); _cdgc {
		_cdeg = _cdeg + "\u003a" + _bddc.DescendantFont.Subtype()
	}
	return _cdeg
}
func (_bceaf *pdfCIDFontType0) baseFields() *fontCommon { return &_bceaf.fontCommon }
func (_fddg *PdfWriter) flushWriter() error {
	if _fddg._bgef == nil {
		_fddg._bgef = _fddg._cbabb.Flush()
	}
	return _fddg._bgef
}

// GetNumComponents returns the number of color components (3 for RGB).
func (_ddbe *PdfColorDeviceRGB) GetNumComponents() int { return 3 }

// NewPdfAnnotationLink returns a new link annotation.
func NewPdfAnnotationLink() *PdfAnnotationLink {
	_cadf := NewPdfAnnotation()
	_daf := &PdfAnnotationLink{}
	_daf.PdfAnnotation = _cadf
	_cadf.SetContext(_daf)
	return _daf
}

// GetContainingPdfObject implements interface PdfModel.
func (_dbdfg *PdfFilespec) GetContainingPdfObject() _ebb.PdfObject { return _dbdfg._gcge }

// ToPdfObject implements interface PdfModel.
func (_dadg *PdfAnnotationPolygon) ToPdfObject() _ebb.PdfObject {
	_dadg.PdfAnnotation.ToPdfObject()
	_cdeb := _dadg._bdcd
	_aadc := _cdeb.PdfObject.(*_ebb.PdfObjectDictionary)
	_dadg.PdfAnnotationMarkup.appendToPdfDictionary(_aadc)
	_aadc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0050o\u006c\u0079\u0067\u006f\u006e"))
	_aadc.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _dadg.Vertices)
	_aadc.SetIfNotNil("\u004c\u0045", _dadg.LE)
	_aadc.SetIfNotNil("\u0042\u0053", _dadg.BS)
	_aadc.SetIfNotNil("\u0049\u0043", _dadg.IC)
	_aadc.SetIfNotNil("\u0042\u0045", _dadg.BE)
	_aadc.SetIfNotNil("\u0049\u0054", _dadg.IT)
	_aadc.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _dadg.Measure)
	return _cdeb
}
func _feef(_efdd _ebb.PdfObject, _ebef *fontCommon) (*_ebe.CMap, error) {
	_gefb, _ggfc := _ebb.GetStream(_efdd)
	if !_ggfc {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0074\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0054\u006f\u0043m\u0061\u0070\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0054\u0029", _efdd)
		return nil, _ebb.ErrTypeError
	}
	_aeaga, _cdfeb := _ebb.DecodeStream(_gefb)
	if _cdfeb != nil {
		return nil, _cdfeb
	}
	_deeda, _cdfeb := _ebe.LoadCmapFromData(_aeaga, !_ebef.isCIDFont())
	if _cdfeb != nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u004e\u0075\u006d\u0062\u0065\u0072\u003d\u0025\u0064\u0020\u0065\u0072r=\u0025\u0076", _gefb.ObjectNumber, _cdfeb)
	}
	return _deeda, _cdfeb
}

const (
	BorderStyleSolid     BorderStyle = iota
	BorderStyleDashed    BorderStyle = iota
	BorderStyleBeveled   BorderStyle = iota
	BorderStyleInset     BorderStyle = iota
	BorderStyleUnderline BorderStyle = iota
)

func (_dfd *PdfReader) newPdfActionLaunchFromDict(_bfg *_ebb.PdfObjectDictionary) (*PdfActionLaunch, error) {
	_egaa, _bgf := _gggf(_bfg.Get("\u0046"))
	if _bgf != nil {
		return nil, _bgf
	}
	return &PdfActionLaunch{Win: _bfg.Get("\u0057\u0069\u006e"), Mac: _bfg.Get("\u004d\u0061\u0063"), Unix: _bfg.Get("\u0055\u006e\u0069\u0078"), NewWindow: _bfg.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _egaa}, nil
}

// NewPdfColorCalRGB returns a new CalRBG color.
func NewPdfColorCalRGB(a, b, c float64) *PdfColorCalRGB {
	_dadc := PdfColorCalRGB{a, b, c}
	return &_dadc
}

// PageCallback callback function used in page loading
// that could be used to modify the page content.
//
// Deprecated: will be removed in v4. Use PageProcessCallback instead.
type PageCallback func(_egeed int, _cgggc *PdfPage)

// ColorToRGB converts a CalRGB color to an RGB color.
func (_dcge *PdfColorspaceCalRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_gded, _fecb := color.(*PdfColorCalRGB)
	if !_fecb {
		_eg.Log.Debug("\u0049\u006e\u0070ut\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0063\u0061\u006c\u0020\u0072\u0067\u0062")
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_cfcg := _gded.A()
	_gdea := _gded.B()
	_fffga := _gded.C()
	X := _dcge.Matrix[0]*_cbg.Pow(_cfcg, _dcge.Gamma[0]) + _dcge.Matrix[3]*_cbg.Pow(_gdea, _dcge.Gamma[1]) + _dcge.Matrix[6]*_cbg.Pow(_fffga, _dcge.Gamma[2])
	Y := _dcge.Matrix[1]*_cbg.Pow(_cfcg, _dcge.Gamma[0]) + _dcge.Matrix[4]*_cbg.Pow(_gdea, _dcge.Gamma[1]) + _dcge.Matrix[7]*_cbg.Pow(_fffga, _dcge.Gamma[2])
	Z := _dcge.Matrix[2]*_cbg.Pow(_cfcg, _dcge.Gamma[0]) + _dcge.Matrix[5]*_cbg.Pow(_gdea, _dcge.Gamma[1]) + _dcge.Matrix[8]*_cbg.Pow(_fffga, _dcge.Gamma[2])
	_ddfba := 3.240479*X + -1.537150*Y + -0.498535*Z
	_edba := -0.969256*X + 1.875992*Y + 0.041556*Z
	_adbc := 0.055648*X + -0.204043*Y + 1.057311*Z
	_ddfba = _cbg.Min(_cbg.Max(_ddfba, 0), 1.0)
	_edba = _cbg.Min(_cbg.Max(_edba, 0), 1.0)
	_adbc = _cbg.Min(_cbg.Max(_adbc, 0), 1.0)
	return NewPdfColorDeviceRGB(_ddfba, _edba, _adbc), nil
}

// SetShadingByName sets a shading resource specified by keyName.
func (_agfeb *PdfPageResources) SetShadingByName(keyName _ebb.PdfObjectName, shadingObj _ebb.PdfObject) error {
	if _agfeb.Shading == nil {
		_agfeb.Shading = _ebb.MakeDict()
	}
	_gbdff, _ebedbf := _agfeb.Shading.(*_ebb.PdfObjectDictionary)
	if !_ebedbf {
		return _ebb.ErrTypeError
	}
	_gbdff.Set(keyName, shadingObj)
	return nil
}

// PdfShadingType4 is a Free-form Gouraud-shaded triangle mesh.
type PdfShadingType4 struct {
	*PdfShading
	BitsPerCoordinate *_ebb.PdfObjectInteger
	BitsPerComponent  *_ebb.PdfObjectInteger
	BitsPerFlag       *_ebb.PdfObjectInteger
	Decode            *_ebb.PdfObjectArray
	Function          []PdfFunction
}

// WatermarkImageOptions contains options for configuring the watermark process.
type WatermarkImageOptions struct {
	Alpha               float64
	FitToWidth          bool
	PreserveAspectRatio bool
}

// NewPdfAnnotation returns an initialized generic PDF annotation model.
func NewPdfAnnotation() *PdfAnnotation {
	_gbdd := &PdfAnnotation{}
	_gbdd._bdcd = _ebb.MakeIndirectObject(_ebb.MakeDict())
	return _gbdd
}

// C returns the value of the cyan component of the color.
func (_bbacg *PdfColorDeviceCMYK) C() float64 { return _bbacg[0] }

// NewPdfAnnotationScreen returns a new screen annotation.
func NewPdfAnnotationScreen() *PdfAnnotationScreen {
	_bab := NewPdfAnnotation()
	_cffb := &PdfAnnotationScreen{}
	_cffb.PdfAnnotation = _bab
	_bab.SetContext(_cffb)
	return _cffb
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_ggedf pdfCIDFontType2) GetCharMetrics(code _da.CharCode) (_bad.CharMetrics, bool) {
	if _afea, _gcdf := _ggedf._dgbc[code]; _gcdf {
		return _bad.CharMetrics{Wx: _afea}, true
	}
	_adbbe := rune(code)
	_cdfcc, _fcfg := _ggedf._dceb[_adbbe]
	if !_fcfg {
		_cdfcc = int(_ggedf._bagcb)
	}
	return _bad.CharMetrics{Wx: float64(_cdfcc)}, true
}

// ColorToRGB verifies that the input color is an RGB color. Method exists in
// order to satisfy the PdfColorspace interface.
func (_efaab *PdfColorspaceDeviceRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_gcffa, _ebfe := color.(*PdfColorDeviceRGB)
	if !_ebfe {
		_eg.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072 \u006e\u006f\u0074\u0020\u0064\u0065\u0076\u0069\u0063\u0065 \u0052\u0047\u0042")
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return _gcffa, nil
}

// ParsePdfObject parses input pdf object into given output intent.
func (_eefce *PdfOutputIntent) ParsePdfObject(object _ebb.PdfObject) error {
	_fgfda, _afccb := _ebb.GetDict(object)
	if !_afccb {
		_eg.Log.Error("\u0055\u006e\u006bno\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020%\u0054 \u0066o\u0072 \u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074", object)
		return _gf.New("\u0075\u006e\u006b\u006e\u006fw\u006e\u0020\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074")
	}
	_eefce._faeb = _fgfda
	_eefce.Type, _ = _fgfda.GetString("\u0054\u0079\u0070\u0065")
	_agefe, _afccb := _fgfda.GetString("\u0053")
	if _afccb {
		switch _agefe {
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411":
			_eefce.S = PdfOutputIntentTypeA1
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00412":
			_eefce.S = PdfOutputIntentTypeA2
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00413":
			_eefce.S = PdfOutputIntentTypeA3
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00414":
			_eefce.S = PdfOutputIntentTypeA4
		case "\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0058":
			_eefce.S = PdfOutputIntentTypeX
		}
	}
	_eefce.OutputCondition, _ = _fgfda.GetString("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e")
	_eefce.OutputConditionIdentifier, _ = _fgfda.GetString("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r")
	_eefce.RegistryName, _ = _fgfda.GetString("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065")
	_eefce.Info, _ = _fgfda.GetString("\u0049\u006e\u0066\u006f")
	if _gabgbb, _edeg := _ebb.GetStream(_fgfda.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _edeg {
		_eefce.ColorComponents, _ = _ebb.GetIntVal(_gabgbb.Get("\u004e"))
		_aefbf, _fgadd := _ebb.DecodeStream(_gabgbb)
		if _fgadd != nil {
			return _fgadd
		}
		_eefce.DestOutputProfile = _aefbf
	}
	return nil
}

// EncryptionAlgorithm is used in EncryptOptions to change the default algorithm used to encrypt the document.
type EncryptionAlgorithm int

func _egbeb(_aaaec _ebb.PdfObject) (map[_da.CharCode]float64, error) {
	if _aaaec == nil {
		return nil, nil
	}
	_gbbg, _cgeef := _ebb.GetArray(_aaaec)
	if !_cgeef {
		return nil, nil
	}
	_egceg := map[_da.CharCode]float64{}
	_eccg := _gbbg.Len()
	for _cggc := 0; _cggc < _eccg-1; _cggc++ {
		_abead := _ebb.TraceToDirectObject(_gbbg.Get(_cggc))
		_ggce, _dafd := _ebb.GetIntVal(_abead)
		if !_dafd {
			return nil, _bg.Errorf("\u0042a\u0064\u0020\u0066\u006fn\u0074\u0020\u0057\u0020\u006fb\u006a0\u003a \u0069\u003d\u0025\u0064\u0020\u0025\u0023v", _cggc, _abead)
		}
		_cggc++
		if _cggc > _eccg-1 {
			return nil, _bg.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _gbbg)
		}
		_cccg := _ebb.TraceToDirectObject(_gbbg.Get(_cggc))
		switch _cccg.(type) {
		case *_ebb.PdfObjectArray:
			_gdca, _ := _ebb.GetArray(_cccg)
			if _fecfc, _edfde := _gdca.ToFloat64Array(); _edfde == nil {
				for _ceea := 0; _ceea < len(_fecfc); _ceea++ {
					_egceg[_da.CharCode(_ggce+_ceea)] = _fecfc[_ceea]
				}
			} else {
				return nil, _bg.Errorf("\u0042\u0061\u0064 \u0066\u006f\u006e\u0074 \u0057\u0020\u0061\u0072\u0072\u0061\u0079 \u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _cggc, _cccg)
			}
		case *_ebb.PdfObjectInteger:
			_ecgg, _bdge := _ebb.GetIntVal(_cccg)
			if !_bdge {
				return nil, _bg.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _cggc, _cccg)
			}
			_cggc++
			if _cggc > _eccg-1 {
				return nil, _bg.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _gbbg)
			}
			_eeage := _gbbg.Get(_cggc)
			_dcabf, _fafcd := _ebb.GetNumberAsFloat(_eeage)
			if _fafcd != nil {
				return nil, _bg.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0032\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _cggc, _eeage)
			}
			for _efbaa := _ggce; _efbaa <= _ecgg; _efbaa++ {
				_egceg[_da.CharCode(_efbaa)] = _dcabf
			}
		default:
			return nil, _bg.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057 \u006f\u0062\u006a\u0031\u0020\u0074\u0079p\u0065\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _cggc, _cccg)
		}
	}
	return _egceg, nil
}

type fontFile struct {
	_bgbcg string
	_badbf string
	_gega  _da.SimpleEncoder
}

// SetNameDictionary sets the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_dbebf *PdfWriter) SetNameDictionary(names _ebb.PdfObject) error {
	if names == nil {
		return nil
	}
	_eg.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u004e\u0061\u006d\u0065\u0073\u002e\u002e\u002e")
	_dbebf._dffegd.Set("\u004e\u0061\u006de\u0073", names)
	return _dbebf.addObjects(names)
}

// ToPdfObject implements interface PdfModel.
func (_ceg *PdfActionThread) ToPdfObject() _ebb.PdfObject {
	_ceg.PdfAction.ToPdfObject()
	_cfab := _ceg._abe
	_fc := _cfab.PdfObject.(*_ebb.PdfObjectDictionary)
	_fc.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeThread)))
	if _ceg.F != nil {
		_fc.Set("\u0046", _ceg.F.ToPdfObject())
	}
	_fc.SetIfNotNil("\u0044", _ceg.D)
	_fc.SetIfNotNil("\u0042", _ceg.B)
	return _cfab
}
func _aedce(_fbee _ebb.PdfObject) (*PdfColorspaceICCBased, error) {
	_ccfc := &PdfColorspaceICCBased{}
	if _bga, _cbccb := _fbee.(*_ebb.PdfIndirectObject); _cbccb {
		_ccfc._dagdd = _bga
	}
	_fbee = _ebb.TraceToDirectObject(_fbee)
	_dadb, _gcdb := _fbee.(*_ebb.PdfObjectArray)
	if !_gcdb {
		return nil, _bg.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _dadb.Len() != 2 {
		return nil, _bg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020c\u006f\u006c\u006fr\u0073p\u0061\u0063\u0065")
	}
	_fbee = _ebb.TraceToDirectObject(_dadb.Get(0))
	_eeec, _gcdb := _fbee.(*_ebb.PdfObjectName)
	if !_gcdb {
		return nil, _bg.Errorf("\u0049\u0043\u0043B\u0061\u0073\u0065\u0064 \u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_eeec != "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064" {
		return nil, _bg.Errorf("\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0049\u0043\u0043\u0042a\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_fbee = _dadb.Get(1)
	_dcd, _gcdb := _ebb.GetStream(_fbee)
	if !_gcdb {
		_eg.Log.Error("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006e\u006f\u0074\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067\u0020t\u006f\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020%\u0054", _fbee)
		return nil, _bg.Errorf("\u0049\u0043\u0043Ba\u0073\u0065\u0064\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_fcbg := _dcd.PdfObjectDictionary
	_ecaec, _gcdb := _fcbg.Get("\u004e").(*_ebb.PdfObjectInteger)
	if !_gcdb {
		return nil, _bg.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006d\u0069\u0073\u0073\u0069n\u0067\u0020\u004e\u0020\u0066\u0072\u006fm\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074")
	}
	if *_ecaec != 1 && *_ecaec != 3 && *_ecaec != 4 {
		return nil, _bg.Errorf("\u0049\u0043\u0043\u0042\u0061s\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u006e\u006f\u0074\u0020\u0031\u002c\u0033\u002c\u0034\u0029")
	}
	_ccfc.N = int(*_ecaec)
	if _gbbb := _fcbg.Get("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e"); _gbbb != nil {
		_aaae, _gebec := NewPdfColorspaceFromPdfObject(_gbbb)
		if _gebec != nil {
			return nil, _gebec
		}
		_ccfc.Alternate = _aaae
	}
	if _adbcc := _fcbg.Get("\u0052\u0061\u006eg\u0065"); _adbcc != nil {
		_adbcc = _ebb.TraceToDirectObject(_adbcc)
		_edgg, _bbgc := _adbcc.(*_ebb.PdfObjectArray)
		if !_bbgc {
			return nil, _bg.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0061\u006e \u0061r\u0072\u0061\u0079")
		}
		if _edgg.Len() != 2*_ccfc.N {
			return nil, _bg.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0052\u0061\u006e\u0067e\u0020\u0077\u0072\u006f\u006e\u0067 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073")
		}
		_fdeaf, _bdbfc := _edgg.GetAsFloat64Slice()
		if _bdbfc != nil {
			return nil, _bdbfc
		}
		_ccfc.Range = _fdeaf
	} else {
		_ccfc.Range = make([]float64, 2*_ccfc.N)
		for _faggf := 0; _faggf < _ccfc.N; _faggf++ {
			_ccfc.Range[2*_faggf] = 0.0
			_ccfc.Range[2*_faggf+1] = 1.0
		}
	}
	if _agce := _fcbg.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _agce != nil {
		_fbcgb, _eacea := _agce.(*_ebb.PdfObjectStream)
		if !_eacea {
			return nil, _bg.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u004de\u0074\u0061\u0064\u0061\u0074\u0061\u0020n\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		_ccfc.Metadata = _fbcgb
	}
	_fgda, _ggga := _ebb.DecodeStream(_dcd)
	if _ggga != nil {
		return nil, _ggga
	}
	_ccfc.Data = _fgda
	_ccfc._dfff = _dcd
	return _ccfc, nil
}
func (_eeea *PdfFilespec) getDict() *_ebb.PdfObjectDictionary {
	if _fgbc, _bedf := _eeea._gcge.(*_ebb.PdfIndirectObject); _bedf {
		_dbffa, _ddaa := _fgbc.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_ddaa {
			return nil
		}
		return _dbffa
	} else if _fabec, _faae := _eeea._gcge.(*_ebb.PdfObjectDictionary); _faae {
		return _fabec
	} else {
		_eg.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074 \u0074\u0079p\u0065\u0020(\u0025T\u0029", _eeea._gcge)
		return nil
	}
}

// PdfAnnotationScreen represents Screen annotations.
// (Section 12.5.6.18).
type PdfAnnotationScreen struct {
	*PdfAnnotation
	T  _ebb.PdfObject
	MK _ebb.PdfObject
	A  _ebb.PdfObject
	AA _ebb.PdfObject
}

// PdfAnnotationFileAttachment represents FileAttachment annotations.
// (Section 12.5.6.15).
type PdfAnnotationFileAttachment struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	FS   _ebb.PdfObject
	Name _ebb.PdfObject
}

func (_cgag *PdfReader) newPdfAnnotationTrapNetFromDict(_cfgd *_ebb.PdfObjectDictionary) (*PdfAnnotationTrapNet, error) {
	_fefe := PdfAnnotationTrapNet{}
	return &_fefe, nil
}

// ToPdfObject implements interface PdfModel.
func (_bebgd *PdfSignatureReference) ToPdfObject() _ebb.PdfObject {
	_gbdg := _ebb.MakeDict()
	_gbdg.SetIfNotNil("\u0054\u0079\u0070\u0065", _bebgd.Type)
	_gbdg.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064", _bebgd.TransformMethod)
	_gbdg.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073", _bebgd.TransformParams)
	_gbdg.SetIfNotNil("\u0044\u0061\u0074\u0061", _bebgd.Data)
	_gbdg.SetIfNotNil("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064", _bebgd.DigestMethod)
	return _gbdg
}

// C returns the value of the C component of the color.
func (_cffc *PdfColorCalRGB) C() float64 { return _cffc[2] }

// GetPerms returns the Permissions dictionary
func (_dcfcc *PdfReader) GetPerms() *Permissions { return _dcfcc._fdca }

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
	Metadata *_ebb.PdfObjectStream
	Data     []byte
	_dagdd   *_ebb.PdfIndirectObject
	_dfff    *_ebb.PdfObjectStream
}

func (_cegga *PdfReader) newPdfAnnotationPolygonFromDict(_bdfc *_ebb.PdfObjectDictionary) (*PdfAnnotationPolygon, error) {
	_cbeb := PdfAnnotationPolygon{}
	_fgf, _fafg := _cegga.newPdfAnnotationMarkupFromDict(_bdfc)
	if _fafg != nil {
		return nil, _fafg
	}
	_cbeb.PdfAnnotationMarkup = _fgf
	_cbeb.Vertices = _bdfc.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_cbeb.LE = _bdfc.Get("\u004c\u0045")
	_cbeb.BS = _bdfc.Get("\u0042\u0053")
	_cbeb.IC = _bdfc.Get("\u0049\u0043")
	_cbeb.BE = _bdfc.Get("\u0042\u0045")
	_cbeb.IT = _bdfc.Get("\u0049\u0054")
	_cbeb.Measure = _bdfc.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_cbeb, nil
}
func _dccgg() string {
	_daddc.Lock()
	defer _daddc.Unlock()
	return _ecgdd
}
func (_eedd *PdfReader) newPdfActionHideFromDict(_cda *_ebb.PdfObjectDictionary) (*PdfActionHide, error) {
	return &PdfActionHide{T: _cda.Get("\u0054"), H: _cda.Get("\u0048")}, nil
}

// PdfOutline represents a PDF outline dictionary (Table 152 - p. 376).
type PdfOutline struct {
	PdfOutlineTreeNode
	Parent *PdfOutlineTreeNode
	Count  *int64
	_egee  *_ebb.PdfIndirectObject
}

// HasExtGState checks if ExtGState name is available.
func (_beeebc *PdfPage) HasExtGState(name _ebb.PdfObjectName) bool {
	if _beeebc.Resources == nil {
		return false
	}
	if _beeebc.Resources.ExtGState == nil {
		return false
	}
	_gbaef, _ccfcf := _ebb.TraceToDirectObject(_beeebc.Resources.ExtGState).(*_ebb.PdfObjectDictionary)
	if !_ccfcf {
		_eg.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _ebb.TraceToDirectObject(_beeebc.Resources.ExtGState))
		return false
	}
	_fabbb := _gbaef.Get(name)
	_bbgab := _fabbb != nil
	return _bbgab
}

// ColorFromPdfObjects loads the color from PDF objects.
// The first objects (if present) represent the color in underlying colorspace.  The last one represents
// the name of the pattern.
func (_ace *PdfColorspaceSpecialPattern) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) < 1 {
		return nil, _gf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_gbeec := &PdfColorPattern{}
	_ccbd, _bdbb := objects[len(objects)-1].(*_ebb.PdfObjectName)
	if !_bdbb {
		_eg.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020a\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", objects[len(objects)-1])
		return nil, ErrTypeCheck
	}
	_gbeec.PatternName = *_ccbd
	if len(objects) > 1 {
		_ecbc := objects[0 : len(objects)-1]
		if _ace.UnderlyingCS == nil {
			_eg.Log.Debug("P\u0061\u0074t\u0065\u0072\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0077\u0069\u0074\u0068\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u0020\u0062\u0075\u0074\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0063\u0073\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
			return nil, _gf.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
		}
		_abbdf, _cgecd := _ace.UnderlyingCS.ColorFromPdfObjects(_ecbc)
		if _cgecd != nil {
			_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0076\u0069\u0061\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0063\u0073\u003a\u0020\u0025\u0076", _cgecd)
			return nil, _cgecd
		}
		_gbeec.Color = _abbdf
	}
	return _gbeec, nil
}
func (_geeb *PdfWriter) setWriter(_aaadc _ab.Writer) {
	_geeb._afedd = _geeb._fgdce
	_geeb._cbabb = _ba.NewWriter(_aaadc)
}
func (_ecaa *pdfFontType3) baseFields() *fontCommon { return &_ecaa.fontCommon }

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the A, B and C components of the color.
func (_fdda *PdfColorspaceCalRGB) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_dbff, _fecg := _ebb.GetNumbersAsFloat(objects)
	if _fecg != nil {
		return nil, _fecg
	}
	return _fdda.ColorFromFloats(_dbff)
}

// GetOCProperties returns the optional content properties PdfObject.
func (_bbadf *PdfReader) GetOCProperties() (_ebb.PdfObject, error) {
	_fcagb := _bbadf._fdgda
	_ebfba := _fcagb.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	_ebfba = _ebb.ResolveReference(_ebfba)
	if !_bbadf._ceefa {
		_acbd := _bbadf.traverseObjectData(_ebfba)
		if _acbd != nil {
			return nil, _acbd
		}
	}
	return _ebfba, nil
}
func (_abdgag *PdfWriter) checkCrossReferenceStream() bool {
	_ffcgc := _abdgag._efcge.Major > 1 || (_abdgag._efcge.Major == 1 && _abdgag._efcge.Minor > 4)
	if _abdgag._adeff != nil {
		_ffcgc = *_abdgag._adeff
	}
	return _ffcgc
}

// GetXObjectByName returns the XObject with the specified keyName and the object type.
func (_efgg *PdfPageResources) GetXObjectByName(keyName _ebb.PdfObjectName) (*_ebb.PdfObjectStream, XObjectType) {
	if _efgg.XObject == nil {
		return nil, XObjectTypeUndefined
	}
	_cdaaf, _eccef := _ebb.TraceToDirectObject(_efgg.XObject).(*_ebb.PdfObjectDictionary)
	if !_eccef {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ebb.TraceToDirectObject(_efgg.XObject))
		return nil, XObjectTypeUndefined
	}
	if _beegc := _cdaaf.Get(keyName); _beegc != nil {
		_bdgdg, _ddegc := _ebb.GetStream(_beegc)
		if !_ddegc {
			_eg.Log.Debug("X\u004f\u0062\u006a\u0065\u0063\u0074 \u006e\u006f\u0074\u0020\u0070\u006fi\u006e\u0074\u0069\u006e\u0067\u0020\u0074o\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020%\u0054", _beegc)
			return nil, XObjectTypeUndefined
		}
		_aeabd := _bdgdg.PdfObjectDictionary
		_dgaeg, _ddegc := _ebb.TraceToDirectObject(_aeabd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")).(*_ebb.PdfObjectName)
		if !_ddegc {
			_eg.Log.Debug("\u0058\u004fbj\u0065\u0063\u0074 \u0053\u0075\u0062\u0074ype\u0020no\u0074\u0020\u0061\u0020\u004e\u0061\u006de,\u0020\u0064\u0069\u0063\u0074\u003a\u0020%\u0073", _aeabd.String())
			return nil, XObjectTypeUndefined
		}
		if *_dgaeg == "\u0049\u006d\u0061g\u0065" {
			return _bdgdg, XObjectTypeImage
		} else if *_dgaeg == "\u0046\u006f\u0072\u006d" {
			return _bdgdg, XObjectTypeForm
		} else if *_dgaeg == "\u0050\u0053" {
			return _bdgdg, XObjectTypePS
		} else {
			_eg.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0053\u0075b\u0074\u0079\u0070\u0065\u0020\u006e\u006ft\u0020\u006b\u006e\u006f\u0077\u006e\u0020\u0028\u0025\u0073\u0029", *_dgaeg)
			return nil, XObjectTypeUndefined
		}
	} else {
		return nil, XObjectTypeUndefined
	}
}

// SetName sets the `Name` field of the signature.
func (_egbga *PdfSignature) SetName(name string) { _egbga.Name = _ebb.MakeString(name) }

// PdfColorDeviceGray represents a grayscale color value that shall be represented by a single number in the
// range 0.0 to 1.0 where 0.0 corresponds to black and 1.0 to white.
type PdfColorDeviceGray float64

func (_fafb *pdfFontType3) getFontDescriptor() *PdfFontDescriptor { return _fafb._fbbd }
func (_cgbbg fontCommon) fontFlags() int {
	if _cgbbg._fbbd == nil {
		return 0
	}
	return _cgbbg._fbbd._gfbge
}

// NewOutlineDest returns a new outline destination which can be used
// with outline items.
func NewOutlineDest(page int64, x, y float64) OutlineDest {
	return OutlineDest{Page: page, Mode: "\u0058\u0059\u005a", X: x, Y: y}
}

// GetMediaBox gets the inheritable media box value, either from the page
// or a higher up page/pages struct.
func (_bdafbc *PdfPage) GetMediaBox() (*PdfRectangle, error) {
	if _bdafbc.MediaBox != nil {
		return _bdafbc.MediaBox, nil
	}
	_eeddca := _bdafbc.Parent
	for _eeddca != nil {
		_fabgb, _edde := _ebb.GetDict(_eeddca)
		if !_edde {
			return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _dfcfg := _fabgb.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _dfcfg != nil {
			_feffe, _cgfgcg := _ebb.GetArray(_dfcfg)
			if !_cgfgcg {
				return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0064\u0069a\u0020\u0062\u006f\u0078")
			}
			_cgade, _dfafe := NewPdfRectangle(*_feffe)
			if _dfafe != nil {
				return nil, _dfafe
			}
			return _cgade, nil
		}
		_eeddca = _fabgb.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, _gf.New("m\u0065\u0064\u0069\u0061 b\u006fx\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
}
func (_dgaafe *XObjectImage) getParamsDict() *_ebb.PdfObjectDictionary {
	_gffed := _ebb.MakeDict()
	_gffed.Set("\u0057\u0069\u0064t\u0068", _ebb.MakeInteger(*_dgaafe.Width))
	_gffed.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _ebb.MakeInteger(*_dgaafe.Height))
	_gffed.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _ebb.MakeInteger(int64(_dgaafe.ColorSpace.GetNumComponents())))
	_gffed.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _ebb.MakeInteger(*_dgaafe.BitsPerComponent))
	return _gffed
}

// SetFontByName sets the font specified by keyName to the given object.
func (_fcfff *PdfPageResources) SetFontByName(keyName _ebb.PdfObjectName, obj _ebb.PdfObject) error {
	if _fcfff.Font == nil {
		_fcfff.Font = _ebb.MakeDict()
	}
	_ffaeg, _ddcbff := _ebb.TraceToDirectObject(_fcfff.Font).(*_ebb.PdfObjectDictionary)
	if !_ddcbff {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _ebb.TraceToDirectObject(_fcfff.Font))
		return _ebb.ErrTypeError
	}
	_ffaeg.Set(keyName, obj)
	return nil
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
type pdfCIDFontType0 struct {
	fontCommon
	_cgebd *_ebb.PdfIndirectObject
	_gefge _da.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// (Required) Dictionary that defines the character collection of the CIDFont.
	// See Table 116.
	CIDSystemInfo *_ebb.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW     _ebb.PdfObject
	W      _ebb.PdfObject
	DW2    _ebb.PdfObject
	W2     _ebb.PdfObject
	_afcac map[_da.CharCode]float64
	_gbdb  float64
}

// GetExtGState gets the ExtGState specified by keyName. Returns a bool
// indicating whether it was found or not.
func (_bagga *PdfPageResources) GetExtGState(keyName _ebb.PdfObjectName) (_ebb.PdfObject, bool) {
	if _bagga.ExtGState == nil {
		return nil, false
	}
	_eegbc, _edged := _ebb.TraceToDirectObject(_bagga.ExtGState).(*_ebb.PdfObjectDictionary)
	if !_edged {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bagga.ExtGState)
		return nil, false
	}
	if _bgbcab := _eegbc.Get(keyName); _bgbcab != nil {
		return _bgbcab, true
	}
	return nil, false
}
func _fgdad(_fddcb *_ebb.PdfObjectDictionary, _cdbe *fontCommon) (*pdfCIDFontType2, error) {
	if _cdbe._dfbf != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_eg.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0032\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _cdbe)
		return nil, _ebb.ErrRangeError
	}
	_aggcd := _gfdbb(_cdbe)
	_acba, _gaeae := _ebb.GetDict(_fddcb.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_gaeae {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _cdbe)
		return nil, ErrRequiredAttributeMissing
	}
	_aggcd.CIDSystemInfo = _acba
	_aggcd.DW = _fddcb.Get("\u0044\u0057")
	_aggcd.W = _fddcb.Get("\u0057")
	_aggcd.DW2 = _fddcb.Get("\u0044\u0057\u0032")
	_aggcd.W2 = _fddcb.Get("\u0057\u0032")
	_aggcd.CIDToGIDMap = _fddcb.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070")
	_aggcd._bagcb = 1000.0
	if _fbeba, _fedad := _ebb.GetNumberAsFloat(_aggcd.DW); _fedad == nil {
		_aggcd._bagcb = _fbeba
	}
	_ccaf, _fcfbb := _egbeb(_aggcd.W)
	if _fcfbb != nil {
		return nil, _fcfbb
	}
	if _ccaf == nil {
		_ccaf = map[_da.CharCode]float64{}
	}
	_aggcd._dgbc = _ccaf
	return _aggcd, nil
}
func (_eebda *PdfWriter) addObjects(_fgaebc _ebb.PdfObject) error {
	_eg.Log.Trace("\u0041d\u0064i\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073\u0021")
	if _fedfg, _aadb := _fgaebc.(*_ebb.PdfIndirectObject); _aadb {
		_eg.Log.Trace("\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074")
		_eg.Log.Trace("\u002d \u0025\u0073\u0020\u0028\u0025\u0070)", _fgaebc, _fedfg)
		_eg.Log.Trace("\u002d\u0020\u0025\u0073", _fedfg.PdfObject)
		if _eebda.addObject(_fedfg) {
			_bcede := _eebda.addObjects(_fedfg.PdfObject)
			if _bcede != nil {
				return _bcede
			}
		}
		return nil
	}
	if _fdgad, _ccdb := _fgaebc.(*_ebb.PdfObjectStream); _ccdb {
		_eg.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d")
		_eg.Log.Trace("\u002d \u0025\u0073\u0020\u0025\u0070", _fgaebc, _fgaebc)
		if _eebda.addObject(_fdgad) {
			_edcg := _eebda.addObjects(_fdgad.PdfObjectDictionary)
			if _edcg != nil {
				return _edcg
			}
		}
		return nil
	}
	if _bffg, _aecfc := _fgaebc.(*_ebb.PdfObjectDictionary); _aecfc {
		_eg.Log.Trace("\u0044\u0069\u0063\u0074")
		_eg.Log.Trace("\u002d\u0020\u0025\u0073", _fgaebc)
		for _, _dfgfc := range _bffg.Keys() {
			_dggde := _bffg.Get(_dfgfc)
			if _ggaag, _ddbce := _dggde.(*_ebb.PdfObjectReference); _ddbce {
				_dggde = _ggaag.Resolve()
				_bffg.Set(_dfgfc, _dggde)
			}
			if _dfgfc != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _befgeg := _eebda.addObjects(_dggde); _befgeg != nil {
					return _befgeg
				}
			} else {
				if _, _adgdf := _dggde.(*_ebb.PdfObjectNull); _adgdf {
					continue
				}
				if _ebdbgf := _eebda.hasObject(_dggde); !_ebdbgf {
					_eg.Log.Debug("P\u0061\u0072\u0065\u006e\u0074\u0020o\u0062\u006a\u0020\u006e\u006f\u0074 \u0061\u0064\u0064\u0065\u0064\u0020\u0079e\u0074\u0021\u0021\u0020\u0025\u0054\u0020\u0025\u0070\u0020%\u0076", _dggde, _dggde, _dggde)
					_eebda._eefeb[_dggde] = append(_eebda._eefeb[_dggde], _bffg)
				}
			}
		}
		return nil
	}
	if _agcg, _fccff := _fgaebc.(*_ebb.PdfObjectArray); _fccff {
		_eg.Log.Trace("\u0041\u0072\u0072a\u0079")
		_eg.Log.Trace("\u002d\u0020\u0025\u0073", _fgaebc)
		if _agcg == nil {
			return _gf.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _bedag, _aefdb := range _agcg.Elements() {
			if _caadd, _bfcfb := _aefdb.(*_ebb.PdfObjectReference); _bfcfb {
				_aefdb = _caadd.Resolve()
				_agcg.Set(_bedag, _aefdb)
			}
			if _bfdfb := _eebda.addObjects(_aefdb); _bfdfb != nil {
				return _bfdfb
			}
		}
		return nil
	}
	if _, _feagb := _fgaebc.(*_ebb.PdfObjectReference); _feagb {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0062\u0065\u0020\u0061\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u002d\u0020\u0067\u006f\u0074 \u0025\u0023\u0076\u0021", _fgaebc)
		return _gf.New("r\u0065\u0066\u0065\u0072en\u0063e\u0020\u006e\u006f\u0074\u0020a\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	return nil
}
func (_gebd *PdfReader) newPdfAnnotationLinkFromDict(_fgce *_ebb.PdfObjectDictionary) (*PdfAnnotationLink, error) {
	_abf := PdfAnnotationLink{}
	_abf.A = _fgce.Get("\u0041")
	_abf.Dest = _fgce.Get("\u0044\u0065\u0073\u0074")
	_abf.H = _fgce.Get("\u0048")
	_abf.PA = _fgce.Get("\u0050\u0041")
	_abf.QuadPoints = _fgce.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_abf.BS = _fgce.Get("\u0042\u0053")
	return &_abf, nil
}

// NewPdfAnnotationPolyLine returns a new polyline annotation.
func NewPdfAnnotationPolyLine() *PdfAnnotationPolyLine {
	_ccg := NewPdfAnnotation()
	_aabe := &PdfAnnotationPolyLine{}
	_aabe.PdfAnnotation = _ccg
	_aabe.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ccg.SetContext(_aabe)
	return _aabe
}

// ToPdfObject implements interface PdfModel.
func (_df *PdfActionImportData) ToPdfObject() _ebb.PdfObject {
	_df.PdfAction.ToPdfObject()
	_ffe := _df._abe
	_gbef := _ffe.PdfObject.(*_ebb.PdfObjectDictionary)
	_gbef.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeImportData)))
	if _df.F != nil {
		_gbef.Set("\u0046", _df.F.ToPdfObject())
	}
	return _ffe
}

// PdfAcroForm represents the AcroForm dictionary used for representation of form data in PDF.
type PdfAcroForm struct {
	Fields          *[]*PdfField
	NeedAppearances *_ebb.PdfObjectBool
	SigFlags        *_ebb.PdfObjectInteger
	CO              *_ebb.PdfObjectArray
	DR              *PdfPageResources
	DA              *_ebb.PdfObjectString
	Q               *_ebb.PdfObjectInteger
	XFA             _ebb.PdfObject
	_adcg           *_ebb.PdfIndirectObject
}

// NewPdfSignature creates a new PdfSignature object.
func NewPdfSignature(handler SignatureHandler) *PdfSignature {
	_faefg := &PdfSignature{Type: _ebb.MakeName("\u0053\u0069\u0067"), Handler: handler}
	_agacb := &pdfSignDictionary{PdfObjectDictionary: _ebb.MakeDict(), _dcfab: &handler, _bead: _faefg}
	_faefg._ffbgc = _ebb.MakeIndirectObject(_agacb)
	return _faefg
}
func _adg(_bcgg *PdfPage) map[_ebb.PdfObjectName]_ebb.PdfObject {
	_abfc := make(map[_ebb.PdfObjectName]_ebb.PdfObject)
	if _bcgg.Resources == nil {
		return _abfc
	}
	if _bcgg.Resources.Font != nil {
		if _ddc, _bada := _ebb.GetDict(_bcgg.Resources.Font); _bada {
			for _, _cgad := range _ddc.Keys() {
				_abfc[_cgad] = _ddc.Get(_cgad)
			}
		}
	}
	if _bcgg.Resources.ExtGState != nil {
		if _abbg, _bfb := _ebb.GetDict(_bcgg.Resources.ExtGState); _bfb {
			for _, _fffg := range _abbg.Keys() {
				_abfc[_fffg] = _abbg.Get(_fffg)
			}
		}
	}
	if _bcgg.Resources.XObject != nil {
		if _dedd, _dec := _ebb.GetDict(_bcgg.Resources.XObject); _dec {
			for _, _eeaf := range _dedd.Keys() {
				_abfc[_eeaf] = _dedd.Get(_eeaf)
			}
		}
	}
	if _bcgg.Resources.Pattern != nil {
		if _cfaa, _geba := _ebb.GetDict(_bcgg.Resources.Pattern); _geba {
			for _, _ecgb := range _cfaa.Keys() {
				_abfc[_ecgb] = _cfaa.Get(_ecgb)
			}
		}
	}
	if _bcgg.Resources.Shading != nil {
		if _ddee, _bcac := _ebb.GetDict(_bcgg.Resources.Shading); _bcac {
			for _, _dfa := range _ddee.Keys() {
				_abfc[_dfa] = _ddee.Get(_dfa)
			}
		}
	}
	if _bcgg.Resources.ProcSet != nil {
		if _dbba, _fccg := _ebb.GetDict(_bcgg.Resources.ProcSet); _fccg {
			for _, _gdae := range _dbba.Keys() {
				_abfc[_gdae] = _dbba.Get(_gdae)
			}
		}
	}
	if _bcgg.Resources.Properties != nil {
		if _badb, _agdf := _ebb.GetDict(_bcgg.Resources.Properties); _agdf {
			for _, _dgdc := range _badb.Keys() {
				_abfc[_dgdc] = _badb.Get(_dgdc)
			}
		}
	}
	return _abfc
}

// ImageToRGB convert 1-component grayscale data to 3-component RGB.
func (_bfda *PdfColorspaceDeviceGray) ImageToRGB(img Image) (Image, error) {
	if img.ColorComponents != 1 {
		return img, _gf.New("\u0074\u0068e \u0070\u0072\u006fv\u0069\u0064\u0065\u0064 im\u0061ge\u0020\u0069\u0073\u0020\u006e\u006f\u0074 g\u0072\u0061\u0079\u0020\u0073\u0063\u0061l\u0065")
	}
	_dbgc, _beed := _dg.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._dagcb, img._dgcea)
	if _beed != nil {
		return img, _beed
	}
	_ccead, _beed := _dg.NRGBAConverter.Convert(_dbgc)
	if _beed != nil {
		return img, _beed
	}
	_adgd := _afacb(_ccead.Base())
	_eg.Log.Trace("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079\u0020\u002d>\u0020\u0052\u0047\u0042")
	_eg.Log.Trace("s\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", img.Data)
	_eg.Log.Trace("\u0052G\u0042 \u0073\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", _adgd.Data)
	_eg.Log.Trace("\u0025\u0076\u0020\u002d\u003e\u0020\u0025\u0076", img, _adgd)
	return _adgd, nil
}

// ToPdfObject implements interface PdfModel.
func (_adce *PdfAnnotationMovie) ToPdfObject() _ebb.PdfObject {
	_adce.PdfAnnotation.ToPdfObject()
	_ffeaf := _adce._bdcd
	_gaae := _ffeaf.PdfObject.(*_ebb.PdfObjectDictionary)
	_gaae.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u004d\u006f\u0076i\u0065"))
	_gaae.SetIfNotNil("\u0054", _adce.T)
	_gaae.SetIfNotNil("\u004d\u006f\u0076i\u0065", _adce.Movie)
	_gaae.SetIfNotNil("\u0041", _adce.A)
	return _ffeaf
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_dded pdfFontType3) GetRuneMetrics(r rune) (_bad.CharMetrics, bool) {
	_fgbea := _dded.Encoder()
	if _fgbea == nil {
		_eg.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _dded)
		return _bad.CharMetrics{}, false
	}
	_cgbegf, _aecfg := _fgbea.RuneToCharcode(r)
	if !_aecfg {
		if r != ' ' {
			_eg.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _dded)
		}
		return _bad.CharMetrics{}, false
	}
	_dcege, _gffa := _dded.GetCharMetrics(_cgbegf)
	return _dcege, _gffa
}

// ToPdfObject converts the font to a PDF representation.
func (_begb *pdfFontType0) ToPdfObject() _ebb.PdfObject {
	if _begb._dfffc == nil {
		_begb._dfffc = &_ebb.PdfIndirectObject{}
	}
	_acdgf := _begb.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0030")
	_begb._dfffc.PdfObject = _acdgf
	if _begb.Encoding != nil {
		_acdgf.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _begb.Encoding)
	} else if _begb._bfdgc != nil {
		_acdgf.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _begb._bfdgc.ToPdfObject())
	}
	if _begb.DescendantFont != nil {
		_acdgf.Set("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073", _ebb.MakeArray(_begb.DescendantFont.ToPdfObject()))
	}
	return _begb._dfffc
}

// PdfActionSubmitForm represents a submitForm action.
type PdfActionSubmitForm struct {
	*PdfAction
	F      *PdfFilespec
	Fields _ebb.PdfObject
	Flags  _ebb.PdfObject
}

// PdfAnnotationPolyLine represents PolyLine annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolyLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _ebb.PdfObject
	LE       _ebb.PdfObject
	BS       _ebb.PdfObject
	IC       _ebb.PdfObject
	BE       _ebb.PdfObject
	IT       _ebb.PdfObject
	Measure  _ebb.PdfObject
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_dggdfb *PdfColorspaceSpecialIndexed) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_adcff, _decc := _ebb.GetNumbersAsFloat(objects)
	if _decc != nil {
		return nil, _decc
	}
	return _dggdfb.ColorFromFloats(_adcff)
}
func _ecaeg(_aaac []byte) []byte {
	const _ebbfd = 52845
	const _acga = 22719
	_ecbba := 55665
	for _, _acbad := range _aaac[:4] {
		_ecbba = (int(_acbad)+_ecbba)*_ebbfd + _acga
	}
	_acde := make([]byte, len(_aaac)-4)
	for _badbfd, _cgbcb := range _aaac[4:] {
		_acde[_badbfd] = byte(int(_cgbcb) ^ _ecbba>>8)
		_ecbba = (int(_cgbcb)+_ecbba)*_ebbfd + _acga
	}
	return _acde
}

// NewPdfAnnotationFileAttachment returns a new file attachment annotation.
func NewPdfAnnotationFileAttachment() *PdfAnnotationFileAttachment {
	_cegc := NewPdfAnnotation()
	_ced := &PdfAnnotationFileAttachment{}
	_ced.PdfAnnotation = _cegc
	_ced.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cegc.SetContext(_ced)
	return _ced
}

// IsColored specifies if the pattern is colored.
func (_faacc *PdfTilingPattern) IsColored() bool {
	if _faacc.PaintType != nil && *_faacc.PaintType == 1 {
		return true
	}
	return false
}

// PdfAnnotationPolygon represents Polygon annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolygon struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _ebb.PdfObject
	LE       _ebb.PdfObject
	BS       _ebb.PdfObject
	IC       _ebb.PdfObject
	BE       _ebb.PdfObject
	IT       _ebb.PdfObject
	Measure  _ebb.PdfObject
}

// NewPdfAnnotationHighlight returns a new text highlight annotation.
func NewPdfAnnotationHighlight() *PdfAnnotationHighlight {
	_cae := NewPdfAnnotation()
	_fbcf := &PdfAnnotationHighlight{}
	_fbcf.PdfAnnotation = _cae
	_fbcf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cae.SetContext(_fbcf)
	return _fbcf
}

// Encoder returns the font's text encoder.
func (_dcfe *pdfFontSimple) Encoder() _da.TextEncoder {
	if _dcfe._ebcb != nil {
		return _dcfe._ebcb
	}
	if _dcfe._dacee != nil {
		return _dcfe._dacee
	}
	_efae, _ := _da.NewSimpleTextEncoder("\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", nil)
	return _efae
}

// AlphaMap performs mapping of alpha data for transformations. Allows custom filtering of alpha data etc.
func (_afgf *Image) AlphaMap(mapFunc AlphaMapFunc) {
	for _fabg, _ebaeg := range _afgf._dagcb {
		_afgf._dagcb[_fabg] = mapFunc(_ebaeg)
	}
}

// NewPdfAnnotationCaret returns a new caret annotation.
func NewPdfAnnotationCaret() *PdfAnnotationCaret {
	_cce := NewPdfAnnotation()
	_ffad := &PdfAnnotationCaret{}
	_ffad.PdfAnnotation = _cce
	_ffad.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cce.SetContext(_ffad)
	return _ffad
}
func (_dacfe *pdfFontSimple) updateStandard14Font() {
	_gabe, _aceb := _dacfe.Encoder().(_da.SimpleEncoder)
	if !_aceb {
		_eg.Log.Error("\u0057\u0072\u006f\u006e\u0067\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074y\u0070e\u003a\u0020\u0025\u0054\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u002e", _dacfe.Encoder(), _dacfe)
		return
	}
	_cdgg := _gabe.Charcodes()
	_dacfe._cdff = make(map[_da.CharCode]float64, len(_cdgg))
	for _, _eaeg := range _cdgg {
		_bgeafe, _ := _gabe.CharcodeToRune(_eaeg)
		_affg, _ := _dacfe._ddgd.Read(_bgeafe)
		_dacfe._cdff[_eaeg] = _affg.Wx
	}
}

// ValidateSignatures validates digital signatures in the document.
func (_egaag *PdfReader) ValidateSignatures(handlers []SignatureHandler) ([]SignatureValidationResult, error) {
	if _egaag.AcroForm == nil {
		return nil, nil
	}
	if _egaag.AcroForm.Fields == nil {
		return nil, nil
	}
	type sigFieldPair struct {
		_daecec *PdfSignature
		_aaada  *PdfField
		_gccc   SignatureHandler
	}
	var _adcfda []*sigFieldPair
	for _, _aegc := range _egaag.AcroForm.AllFields() {
		if _aegc.V == nil {
			continue
		}
		if _dadbg, _ffege := _ebb.GetDict(_aegc.V); _ffege {
			if _daecbg, _bfddc := _ebb.GetNameVal(_dadbg.Get("\u0054\u0079\u0070\u0065")); _bfddc && _daecbg == "\u0053\u0069\u0067" {
				_edcecc, _ffgg := _ebb.GetIndirect(_aegc.V)
				if !_ffgg {
					_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0020\u0069s\u0020\u006e\u0069\u006c")
					return nil, ErrTypeCheck
				}
				_cacb, _bega := _egaag.newPdfSignatureFromIndirect(_edcecc)
				if _bega != nil {
					return nil, _bega
				}
				var _cfdg SignatureHandler
				for _, _fgbeec := range handlers {
					if _fgbeec.IsApplicable(_cacb) {
						_cfdg = _fgbeec
						break
					}
				}
				_adcfda = append(_adcfda, &sigFieldPair{_daecec: _cacb, _aaada: _aegc, _gccc: _cfdg})
			}
		}
	}
	var _bffedb []SignatureValidationResult
	for _, _bcbeb := range _adcfda {
		_ccdfe := SignatureValidationResult{IsSigned: true, Fields: []*PdfField{_bcbeb._aaada}}
		if _bcbeb._gccc == nil {
			_ccdfe.Errors = append(_ccdfe.Errors, "\u0068a\u006ed\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
			_bffedb = append(_bffedb, _ccdfe)
			continue
		}
		_fcfaf, _bfcfd := _bcbeb._gccc.NewDigest(_bcbeb._daecec)
		if _bfcfd != nil {
			_ccdfe.Errors = append(_ccdfe.Errors, "\u0064\u0069\u0067e\u0073\u0074\u0020\u0065\u0072\u0072\u006f\u0072", _bfcfd.Error())
			_bffedb = append(_bffedb, _ccdfe)
			continue
		}
		_caabf := _bcbeb._daecec.ByteRange
		if _caabf == nil {
			_ccdfe.Errors = append(_ccdfe.Errors, "\u0042\u0079\u0074\u0065\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
			_bffedb = append(_bffedb, _ccdfe)
			continue
		}
		for _bbcb := 0; _bbcb < _caabf.Len(); _bbcb = _bbcb + 2 {
			_fcgeb, _ := _ebb.GetNumberAsInt64(_caabf.Get(_bbcb))
			_dfcag, _ := _ebb.GetIntVal(_caabf.Get(_bbcb + 1))
			if _, _aabge := _egaag._ggdg.Seek(_fcgeb, _ab.SeekStart); _aabge != nil {
				return nil, _aabge
			}
			_begeg := make([]byte, _dfcag)
			if _, _ggcdb := _egaag._ggdg.Read(_begeg); _ggcdb != nil {
				return nil, _ggcdb
			}
			_fcfaf.Write(_begeg)
		}
		var _bcceg SignatureValidationResult
		if _fegec, _cdagd := _bcbeb._gccc.(SignatureHandlerDocMDP); _cdagd {
			_bcceg, _bfcfd = _fegec.ValidateWithOpts(_bcbeb._daecec, _fcfaf, SignatureHandlerDocMDPParams{Parser: _egaag._cafdf})
		} else {
			_bcceg, _bfcfd = _bcbeb._gccc.Validate(_bcbeb._daecec, _fcfaf)
		}
		if _bfcfd != nil {
			_eg.Log.Debug("E\u0052\u0052\u004f\u0052: \u0025v\u0020\u0028\u0025\u0054\u0029 \u002d\u0020\u0073\u006b\u0069\u0070", _bfcfd, _bcbeb._gccc)
			_bcceg.Errors = append(_bcceg.Errors, _bfcfd.Error())
		}
		_bcceg.Name = _bcbeb._daecec.Name.Decoded()
		_bcceg.Reason = _bcbeb._daecec.Reason.Decoded()
		if _bcbeb._daecec.M != nil {
			_efbgb, _beaag := NewPdfDate(_bcbeb._daecec.M.String())
			if _beaag != nil {
				_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _beaag)
				_bcceg.Errors = append(_bcceg.Errors, _beaag.Error())
				continue
			}
			_bcceg.Date = _efbgb
		}
		_bcceg.ContactInfo = _bcbeb._daecec.ContactInfo.Decoded()
		_bcceg.Location = _bcbeb._daecec.Location.Decoded()
		_bcceg.Fields = _ccdfe.Fields
		_bffedb = append(_bffedb, _bcceg)
	}
	return _bffedb, nil
}

// GetOutlines returns a high-level Outline object, based on the outline tree
// of the reader.
func (_gcfbc *PdfReader) GetOutlines() (*Outline, error) {
	if _gcfbc == nil {
		return nil, _gf.New("\u0063\u0061n\u006e\u006f\u0074\u0020c\u0072\u0065a\u0074\u0065\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u006e\u0069\u006c\u0020\u0072e\u0061\u0064\u0065\u0072")
	}
	_afcb := _gcfbc.GetOutlineTree()
	if _afcb == nil {
		return nil, _gf.New("\u0074\u0068\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u0064\u006f\u0065\u0073\u0020n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065")
	}
	var _ebega func(_efeeb *PdfOutlineTreeNode, _bgadd *[]*OutlineItem)
	_ebega = func(_gabbb *PdfOutlineTreeNode, _cbgbg *[]*OutlineItem) {
		if _gabbb == nil {
			return
		}
		if _gabbb._geeee == nil {
			_eg.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020m\u0069\u0073\u0073\u0069ng \u006fut\u006c\u0069\u006e\u0065\u0020\u0065\u006etr\u0079\u0020\u0063\u006f\u006e\u0074\u0065x\u0074")
			return
		}
		var _afbea *OutlineItem
		if _ccba, _cacecb := _gabbb._geeee.(*PdfOutlineItem); _cacecb {
			_aefc := _ccba.Dest
			if (_aefc == nil || _ebb.IsNullObject(_aefc)) && _ccba.A != nil {
				if _gcab, _degf := _ebb.GetDict(_ccba.A); _degf {
					if _cgggb, _facaa := _ebb.GetArray(_gcab.Get("\u0044")); _facaa {
						_aefc = _cgggb
					} else {
						_aecfe, _abacg := _ebb.GetString(_gcab.Get("\u0044"))
						if !_abacg {
							return
						}
						_eaeca, _abacg := _gcfbc._fdgda.Get("\u004e\u0061\u006de\u0073").(*_ebb.PdfObjectReference)
						if !_abacg {
							return
						}
						_egfdd, _cbgg := _gcfbc._cafdf.LookupByReference(*_eaeca)
						if _cbgg != nil {
							_eg.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0028\u0025\u0073\u0029", _cbgg.Error())
							return
						}
						_dccab, _abacg := _egfdd.(*_ebb.PdfIndirectObject)
						if !_abacg {
							return
						}
						_fgfc := map[_ebb.PdfObject]struct{}{}
						_cbgg = _gcfbc.buildNameNodes(_dccab, _fgfc)
						if _cbgg != nil {
							_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006ea\u006d\u0065\u0020\u006e\u006fd\u0065\u0073 \u0028\u0025\u0073\u0029", _cbgg.Error())
							return
						}
						for _gcbed := range _fgfc {
							_geaf, _agbe := _ebb.GetDict(_gcbed)
							if !_agbe {
								continue
							}
							_bebbd, _agbe := _ebb.GetArray(_geaf.Get("\u004e\u0061\u006de\u0073"))
							if !_agbe {
								continue
							}
							for _dead, _eacagc := range _bebbd.Elements() {
								switch _eacagc.(type) {
								case *_ebb.PdfObjectString:
									if _eacagc.String() == _aecfe.String() {
										if _gcbeg := _bebbd.Get(_dead + 1); _gcbeg != nil {
											if _bdefe, _gdfg := _ebb.GetDict(_gcbeg); _gdfg {
												_aefc = _bdefe.Get("\u0044")
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
			var _dcdfgg OutlineDest
			if _aefc != nil && !_ebb.IsNullObject(_aefc) {
				if _fdee, _ffdee := _fbga(_aefc, _gcfbc); _ffdee == nil {
					_dcdfgg = *_fdee
				} else {
					_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020p\u0061\u0072\u0073\u0065\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065\u0073\u0074\u0020\u0028\u0025\u0076\u0029\u003a\u0020\u0025\u0076", _aefc, _ffdee)
				}
			}
			_afbea = NewOutlineItem(_ccba.Title.Decoded(), _dcdfgg)
			*_cbgbg = append(*_cbgbg, _afbea)
			if _ccba.Next != nil {
				_ebega(_ccba.Next, _cbgbg)
			}
		}
		if _gabbb.First != nil {
			if _afbea != nil {
				_cbgbg = &_afbea.Entries
			}
			_ebega(_gabbb.First, _cbgbg)
		}
	}
	_fdfac := NewOutline()
	_ebega(_afcb, &_fdfac.Entries)
	return _fdfac, nil
}

// SetSubtype sets the Subtype S for given PdfOutputIntent.
func (_gdbeg *PdfOutputIntent) SetSubtype(subtype PdfOutputIntentType) error {
	if !subtype.IsValid() {
		return _gf.New("\u0070\u0072o\u0076\u0069\u0064\u0065d\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u004f\u0075t\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0053\u0075b\u0054\u0079\u0070\u0065")
	}
	_gdbeg.S = subtype
	return nil
}

// SetFillImage attach a model.Image to push button.
func (_ddgba *PdfFieldButton) SetFillImage(image *Image) {
	if _ddgba.IsPush() {
		_ddgba._cdbf = image
	}
}

// GetFillImage get attached model.Image in push button.
func (_gbfg *PdfFieldButton) GetFillImage() *Image {
	if _gbfg.IsPush() {
		return _gbfg._cdbf
	}
	return nil
}

// NewPdfFilespec returns an initialized generic PDF filespec model.
func NewPdfFilespec() *PdfFilespec {
	_ffdg := &PdfFilespec{}
	_ffdg._gcge = _ebb.MakeIndirectObject(_ebb.MakeDict())
	return _ffdg
}

// NewCompliancePdfReader creates a PdfReader or an input io.ReadSeeker that during reading will scan the files for the
// metadata details. It could be used for the PDF standard implementations like PDF/A or PDF/X.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
func NewCompliancePdfReader(rs _ab.ReadSeeker) (*CompliancePdfReader, error) {
	const _ggaf = "\u006d\u006f\u0064\u0065l\u003a\u004e\u0065\u0077\u0043\u006f\u006d\u0070\u006c\u0069a\u006ec\u0065\u0050\u0064\u0066\u0052\u0065\u0061d\u0065\u0072"
	_ffcg, _gabg := _dcbd(rs, &ReaderOpts{ComplianceMode: true}, false, _ggaf)
	if _gabg != nil {
		return nil, _gabg
	}
	return &CompliancePdfReader{PdfReader: _ffcg}, nil
}
func (_gdde *pdfFontType0) getFontDescriptor() *PdfFontDescriptor {
	if _gdde._fbbd == nil && _gdde.DescendantFont != nil {
		return _gdde.DescendantFont.FontDescriptor()
	}
	return _gdde._fbbd
}
func (_gecf fontCommon) asPdfObjectDictionary(_aecf string) *_ebb.PdfObjectDictionary {
	if _aecf != "" && _gecf._dfbf != "" && _aecf != _gecf._dfbf {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061\u0073\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020O\u0076\u0065\u0072\u0072\u0069\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0062t\u0079\u0070\u0065\u0020\u0074\u006f \u0025\u0023\u0071 \u0025\u0073", _aecf, _gecf)
	} else if _aecf == "" && _gecf._dfbf == "" {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061s\u0050\u0064\u0066Ob\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074y\u0070\u0065\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _gecf)
	} else if _gecf._dfbf == "" {
		_gecf._dfbf = _aecf
	}
	_dcbad := _ebb.MakeDict()
	_dcbad.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0046\u006f\u006e\u0074"))
	_dcbad.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _ebb.MakeName(_gecf._fdacg))
	_dcbad.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName(_gecf._dfbf))
	if _gecf._fbbd != nil {
		_dcbad.Set("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072", _gecf._fbbd.ToPdfObject())
	}
	if _gecf._baag != nil {
		_dcbad.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _gecf._baag)
	} else if _gecf._dcdd != nil {
		_cgfda, _ggadd := _gecf._dcdd.Stream()
		if _ggadd != nil {
			_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0067\u0065\u0074\u0020C\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0065r\u0072\u003d\u0025\u0076", _ggadd)
		} else {
			_dcbad.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _cgfda)
		}
	}
	return _dcbad
}
func (_cfc *PdfReader) newPdfActionImportDataFromDict(_agg *_ebb.PdfObjectDictionary) (*PdfActionImportData, error) {
	_eef, _gfe := _gggf(_agg.Get("\u0046"))
	if _gfe != nil {
		return nil, _gfe
	}
	return &PdfActionImportData{F: _eef}, nil
}

var _dbagf = _a.MustCompile("\u005b\\\u006e\u005c\u0072\u005d\u002b")

// NewPdfOutlineItem returns an initialized PdfOutlineItem.
func NewPdfOutlineItem() *PdfOutlineItem {
	_efaf := &PdfOutlineItem{_cacdf: _ebb.MakeIndirectObject(_ebb.MakeDict())}
	_efaf._geeee = _efaf
	return _efaf
}

// PdfActionHide represents a hide action.
type PdfActionHide struct {
	*PdfAction
	T _ebb.PdfObject
	H _ebb.PdfObject
}

// NewPdfColorspaceFromPdfObject loads a PdfColorspace from a PdfObject.  Returns an error if there is
// a failure in loading.
func NewPdfColorspaceFromPdfObject(obj _ebb.PdfObject) (PdfColorspace, error) {
	if obj == nil {
		return nil, nil
	}
	var _bcag *_ebb.PdfIndirectObject
	var _daec *_ebb.PdfObjectName
	var _fecd *_ebb.PdfObjectArray
	if _eedf, _ebgd := obj.(*_ebb.PdfIndirectObject); _ebgd {
		_bcag = _eedf
	}
	obj = _ebb.TraceToDirectObject(obj)
	switch _fcbf := obj.(type) {
	case *_ebb.PdfObjectArray:
		_fecd = _fcbf
	case *_ebb.PdfObjectName:
		_daec = _fcbf
	}
	if _daec != nil {
		switch *_daec {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			return NewPdfColorspaceDeviceGray(), nil
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			return NewPdfColorspaceDeviceRGB(), nil
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return NewPdfColorspaceDeviceCMYK(), nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return NewPdfColorspaceSpecialPattern(), nil
		default:
			_eg.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0025\u0073", *_daec)
			return nil, _fddb
		}
	}
	if _fecd != nil && _fecd.Len() > 0 {
		var _eaag _ebb.PdfObject = _bcag
		if _bcag == nil {
			_eaag = _fecd
		}
		if _ebga, _cacd := _ebb.GetName(_fecd.Get(0)); _cacd {
			switch _ebga.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				if _fecd.Len() == 1 {
					return NewPdfColorspaceDeviceGray(), nil
				}
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				if _fecd.Len() == 1 {
					return NewPdfColorspaceDeviceRGB(), nil
				}
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _fecd.Len() == 1 {
					return NewPdfColorspaceDeviceCMYK(), nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079":
				return _ageed(_eaag)
			case "\u0043\u0061\u006c\u0052\u0047\u0042":
				return _fcgc(_eaag)
			case "\u004c\u0061\u0062":
				return _efad(_eaag)
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064":
				return _aedce(_eaag)
			case "\u0050a\u0074\u0074\u0065\u0072\u006e":
				return _abba(_eaag)
			case "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return _dffb(_eaag)
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e":
				return _agcbd(_eaag)
			case "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return _egaaa(_eaag)
			default:
				_eg.Log.Debug("A\u0072\u0072\u0061\u0079\u0020\u0077i\u0074\u0068\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0061m\u0065:\u0020\u0025\u0073", *_ebga)
			}
		}
	}
	_eg.Log.Debug("\u0050\u0044\u0046\u0020\u0046i\u006c\u0065\u0020\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", obj.String())
	return nil, ErrTypeCheck
}

// NewPdfActionSound returns a new "sound" action.
func NewPdfActionSound() *PdfActionSound {
	_ec := NewPdfAction()
	_gc := &PdfActionSound{}
	_gc.PdfAction = _ec
	_ec.SetContext(_gc)
	return _gc
}
func (_dge *PdfReader) newPdfAnnotationHighlightFromDict(_bbbd *_ebb.PdfObjectDictionary) (*PdfAnnotationHighlight, error) {
	_bebf := PdfAnnotationHighlight{}
	_begf, _bbf := _dge.newPdfAnnotationMarkupFromDict(_bbbd)
	if _bbf != nil {
		return nil, _bbf
	}
	_bebf.PdfAnnotationMarkup = _begf
	_bebf.QuadPoints = _bbbd.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_bebf, nil
}

// SubsetRegistered subsets the font to only the glyphs that have been registered by the encoder.
// NOTE: This only works on fonts that support subsetting. For unsupported fonts this is a no-op, although a debug
//   message is emitted.  Currently supported fonts are embedded Truetype CID fonts (type 0).
// NOTE: Make sure to call this soon before writing (once all needed runes have been registered).
// If using package creator, use its EnableFontSubsetting method instead.
func (_gddbb *PdfFont) SubsetRegistered() error {
	switch _fabag := _gddbb._ebcad.(type) {
	case *pdfFontType0:
		_ebfec := _fabag.subsetRegistered()
		if _ebfec != nil {
			_eg.Log.Debug("\u0053\u0075b\u0073\u0065\u0074 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _ebfec)
			return _ebfec
		}
		if _fabag._dfffc != nil {
			if _fabag._bfdgc != nil {
				_fabag._bfdgc.ToPdfObject()
			}
			_fabag.ToPdfObject()
		}
	default:
		_eg.Log.Debug("F\u006f\u006e\u0074\u0020\u0025\u0054 \u0064\u006f\u0065\u0073\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006fr\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069n\u0067", _fabag)
	}
	return nil
}

// Encoder returns the font's text encoder.
func (_fbgff pdfCIDFontType0) Encoder() _da.TextEncoder { return _fbgff._gefge }

var _ pdfFont = (*pdfFontSimple)(nil)

// Encoder returns the font's text encoder.
func (_dffc pdfFontType0) Encoder() _da.TextEncoder { return _dffc._bfdgc }

// R returns the value of the red component of the color.
func (_aade *PdfColorDeviceRGB) R() float64      { return _aade[0] }
func _gfdbb(_aacab *fontCommon) *pdfCIDFontType2 { return &pdfCIDFontType2{fontCommon: *_aacab} }

// FieldFilterFunc represents a PDF field filtering function. If the function
// returns true, the PDF field is kept, otherwise it is discarded.
type FieldFilterFunc func(*PdfField) bool

// OutlineItem represents a PDF outline item dictionary (Table 153 - pp. 376 - 377).
type OutlineItem struct {
	Title   string         `json:"title"`
	Dest    OutlineDest    `json:"dest"`
	Entries []*OutlineItem `json:"entries,omitempty"`
}

func (_cccbf *PdfFunctionType0) processSamples() error {
	_eddb := _abg.ResampleBytes(_cccbf._aeecg, _cccbf.BitsPerSample)
	_cccbf._egdg = _eddb
	return nil
}

// PdfFilespec represents a file specification which can either refer to an external or embedded file.
type PdfFilespec struct {
	Type  _ebb.PdfObject
	FS    _ebb.PdfObject
	F     _ebb.PdfObject
	UF    _ebb.PdfObject
	DOS   _ebb.PdfObject
	Mac   _ebb.PdfObject
	Unix  _ebb.PdfObject
	ID    _ebb.PdfObject
	V     _ebb.PdfObject
	EF    _ebb.PdfObject
	RF    _ebb.PdfObject
	Desc  _ebb.PdfObject
	CI    _ebb.PdfObject
	_gcge _ebb.PdfObject
}

func (_bcgga *PdfWriter) setHashIDs(_bcdcc _c.Hash) error {
	_dfgbe := _bcdcc.Sum(nil)
	if _bcgga._gfdea == "" {
		_bcgga._gfdea = _d.EncodeToString(_dfgbe[:8])
	}
	_bcgga.setDocumentIDs(_bcgga._gfdea, _d.EncodeToString(_dfgbe[8:]))
	return nil
}

// String returns a string that describes `font`.
func (_eeda *PdfFont) String() string {
	_ggf := ""
	if _eeda._ebcad.Encoder() != nil {
		_ggf = _eeda._ebcad.Encoder().String()
	}
	return _bg.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0054\u0020\u0025s\u0020\u0025\u0073\u007d", _eeda._ebcad, _eeda.baseFields().coreString(), _ggf)
}
func (_gdgf *PdfReader) newPdfFieldFromIndirectObject(_ddfd *_ebb.PdfIndirectObject, _cege *PdfField) (*PdfField, error) {
	if _aaed, _cfga := _gdgf._abbaca.GetModelFromPrimitive(_ddfd).(*PdfField); _cfga {
		return _aaed, nil
	}
	_bgbb, _fefec := _ebb.GetDict(_ddfd)
	if !_fefec {
		return nil, _bg.Errorf("\u0050\u0064f\u0046\u0069\u0065\u006c\u0064 \u0069\u006e\u0064\u0069\u0072e\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_acgf := NewPdfField()
	_acgf._cdfd = _ddfd
	_acgf._cdfd.PdfObject = _bgbb
	if _dcgd, _ecdd := _ebb.GetName(_bgbb.Get("\u0046\u0054")); _ecdd {
		_acgf.FT = _dcgd
	}
	if _cege != nil {
		_acgf.Parent = _cege
	}
	_acgf.T, _ = _bgbb.Get("\u0054").(*_ebb.PdfObjectString)
	_acgf.TU, _ = _bgbb.Get("\u0054\u0055").(*_ebb.PdfObjectString)
	_acgf.TM, _ = _bgbb.Get("\u0054\u004d").(*_ebb.PdfObjectString)
	_acgf.Ff, _ = _bgbb.Get("\u0046\u0066").(*_ebb.PdfObjectInteger)
	_acgf.V = _bgbb.Get("\u0056")
	_acgf.DV = _bgbb.Get("\u0044\u0056")
	_acgf.AA = _bgbb.Get("\u0041\u0041")
	if DA := _bgbb.Get("\u0044\u0041"); DA != nil {
		DA, _ := _ebb.GetString(DA)
		_acgf.VariableText = &VariableText{DA: DA}
		Q, _ := _bgbb.Get("\u0051").(*_ebb.PdfObjectInteger)
		DS, _ := _bgbb.Get("\u0044\u0053").(*_ebb.PdfObjectString)
		RV := _bgbb.Get("\u0052\u0056")
		_acgf.VariableText.Q = Q
		_acgf.VariableText.DS = DS
		_acgf.VariableText.RV = RV
	}
	_gbgbc := _acgf.FT
	if _gbgbc == nil && _cege != nil {
		_gbgbc = _cege.FT
	}
	if _gbgbc != nil {
		switch *_gbgbc {
		case "\u0054\u0078":
			_gace, _aadd := _cfgdf(_bgbb)
			if _aadd != nil {
				return nil, _aadd
			}
			_gace.PdfField = _acgf
			_acgf._cada = _gace
		case "\u0043\u0068":
			_caae, _bcecc := _egcdd(_bgbb)
			if _bcecc != nil {
				return nil, _bcecc
			}
			_caae.PdfField = _acgf
			_acgf._cada = _caae
		case "\u0042\u0074\u006e":
			_gggff, _edeae := _gccbf(_bgbb)
			if _edeae != nil {
				return nil, _edeae
			}
			_gggff.PdfField = _acgf
			_acgf._cada = _gggff
		case "\u0053\u0069\u0067":
			_cgcc, _fgccc := _gdgf.newPdfFieldSignatureFromDict(_bgbb)
			if _fgccc != nil {
				return nil, _fgccc
			}
			_cgcc.PdfField = _acgf
			_acgf._cada = _cgcc
		default:
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", *_acgf.FT)
			return nil, _gf.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079p\u0065")
		}
	}
	if _fadc, _deef := _ebb.GetName(_bgbb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _deef {
		if *_fadc == "\u0057\u0069\u0064\u0067\u0065\u0074" {
			_abbe, _egabg := _gdgf.newPdfAnnotationFromIndirectObject(_ddfd)
			if _egabg != nil {
				return nil, _egabg
			}
			_adda, _bdfd := _abbe.GetContext().(*PdfAnnotationWidget)
			if !_bdfd {
				return nil, _gf.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0067e\u0074 \u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006fn")
			}
			_adda._gce = _acgf
			_adda.Parent = _acgf._cdfd
			_acgf.Annotations = append(_acgf.Annotations, _adda)
			return _acgf, nil
		}
	}
	_aebg := true
	if _dcee, _dfagfg := _ebb.GetArray(_bgbb.Get("\u004b\u0069\u0064\u0073")); _dfagfg {
		_fcfea := make([]*_ebb.PdfIndirectObject, 0, _dcee.Len())
		for _, _gaaf := range _dcee.Elements() {
			_fgeg, _afba := _ebb.GetIndirect(_gaaf)
			if !_afba {
				_cffbc, _cgdfg := _ebb.GetStream(_gaaf)
				if _cgdfg && _cffbc.PdfObjectDictionary != nil {
					_efeee, _feda := _ebb.GetNameVal(_cffbc.Get("\u0054\u0079\u0070\u0065"))
					if _feda && _efeee == "\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061" {
						_eg.Log.Debug("E\u0052RO\u0052:\u0020f\u006f\u0072\u006d\u0020\u0066i\u0065\u006c\u0064 \u004b\u0069\u0064\u0073\u0020a\u0072\u0072\u0061y\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0069n\u0076\u0061\u006cid \u004d\u0065\u0074\u0061\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e")
						continue
					}
				}
				return nil, _gf.New("n\u006f\u0074\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0028\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065\u006cd\u0029")
			}
			_dgdee, _agdca := _ebb.GetDict(_fgeg)
			if !_agdca {
				return nil, ErrTypeCheck
			}
			if _aebg {
				_aebg = !_abgfe(_dgdee)
			}
			_fcfea = append(_fcfea, _fgeg)
		}
		for _, _ceab := range _fcfea {
			if _aebg {
				_begd, _eedfa := _gdgf.newPdfAnnotationFromIndirectObject(_ceab)
				if _eedfa != nil {
					_eg.Log.Debug("\u0045r\u0072\u006fr\u0020\u006c\u006fa\u0064\u0069\u006e\u0067\u0020\u0077\u0069d\u0067\u0065\u0074\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0072 \u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _eedfa)
					return nil, _eedfa
				}
				_bbfc, _gaeca := _begd._efd.(*PdfAnnotationWidget)
				if !_gaeca {
					return nil, ErrTypeCheck
				}
				_bbfc._gce = _acgf
				_acgf.Annotations = append(_acgf.Annotations, _bbfc)
			} else {
				_eedff, _adfgg := _gdgf.newPdfFieldFromIndirectObject(_ceab, _acgf)
				if _adfgg != nil {
					_eg.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0068\u0069\u006c\u0064\u0020\u0066\u0069\u0065\u006c\u0064: \u0025\u0076", _adfgg)
					return nil, _adfgg
				}
				_acgf.Kids = append(_acgf.Kids, _eedff)
			}
		}
	}
	return _acgf, nil
}
func _bcbe(_dbca *_ebb.PdfObjectDictionary) *VRI {
	_fbde, _ := _ebb.GetString(_dbca.Get("\u0054\u0055"))
	_cccba, _ := _ebb.GetString(_dbca.Get("\u0054\u0053"))
	return &VRI{Cert: _ffaa(_dbca.Get("\u0043\u0065\u0072\u0074")), OCSP: _ffaa(_dbca.Get("\u004f\u0043\u0053\u0050")), CRL: _ffaa(_dbca.Get("\u0043\u0052\u004c")), TU: _fbde, TS: _cccba}
}

// NewPdfRectangle creates a PDF rectangle object based on an input array of 4 integers.
// Defining the lower left (LL) and upper right (UR) corners with
// floating point numbers.
func NewPdfRectangle(arr _ebb.PdfObjectArray) (*PdfRectangle, error) {
	_efabe := PdfRectangle{}
	if arr.Len() != 4 {
		return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065\u0020\u0061\u0072r\u0061\u0079\u002c\u0020\u006c\u0065\u006e \u0021\u003d\u0020\u0034")
	}
	var _bacc error
	_efabe.Llx, _bacc = _ebb.GetNumberAsFloat(arr.Get(0))
	if _bacc != nil {
		return nil, _bacc
	}
	_efabe.Lly, _bacc = _ebb.GetNumberAsFloat(arr.Get(1))
	if _bacc != nil {
		return nil, _bacc
	}
	_efabe.Urx, _bacc = _ebb.GetNumberAsFloat(arr.Get(2))
	if _bacc != nil {
		return nil, _bacc
	}
	_efabe.Ury, _bacc = _ebb.GetNumberAsFloat(arr.Get(3))
	if _bacc != nil {
		return nil, _bacc
	}
	return &_efabe, nil
}

// PdfAnnotationFreeText represents FreeText annotations.
// (Section 12.5.6.6).
type PdfAnnotationFreeText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	DA _ebb.PdfObject
	Q  _ebb.PdfObject
	RC _ebb.PdfObject
	DS _ebb.PdfObject
	CL _ebb.PdfObject
	IT _ebb.PdfObject
	BE _ebb.PdfObject
	RD _ebb.PdfObject
	BS _ebb.PdfObject
	LE _ebb.PdfObject
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_aacag *PdfShadingType4) ToPdfObject() _ebb.PdfObject {
	_aacag.PdfShading.ToPdfObject()
	_gbacbc, _fcbe := _aacag.getShadingDict()
	if _fcbe != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _aacag.BitsPerCoordinate != nil {
		_gbacbc.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _aacag.BitsPerCoordinate)
	}
	if _aacag.BitsPerComponent != nil {
		_gbacbc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _aacag.BitsPerComponent)
	}
	if _aacag.BitsPerFlag != nil {
		_gbacbc.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _aacag.BitsPerFlag)
	}
	if _aacag.Decode != nil {
		_gbacbc.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _aacag.Decode)
	}
	if _aacag.Function != nil {
		if len(_aacag.Function) == 1 {
			_gbacbc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _aacag.Function[0].ToPdfObject())
		} else {
			_fedec := _ebb.MakeArray()
			for _, _ebac := range _aacag.Function {
				_fedec.Append(_ebac.ToPdfObject())
			}
			_gbacbc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _fedec)
		}
	}
	return _aacag._fbfae
}

// XObjectImage (Table 89 in 8.9.5.1).
// Implements PdfModel interface.
type XObjectImage struct {

	//ColorSpace       PdfObject
	Width            *int64
	Height           *int64
	ColorSpace       PdfColorspace
	BitsPerComponent *int64
	Filter           _ebb.StreamEncoder
	Intent           _ebb.PdfObject
	ImageMask        _ebb.PdfObject
	Mask             _ebb.PdfObject
	Matte            _ebb.PdfObject
	Decode           _ebb.PdfObject
	Interpolate      _ebb.PdfObject
	Alternatives     _ebb.PdfObject
	SMask            _ebb.PdfObject
	SMaskInData      _ebb.PdfObject
	Name             _ebb.PdfObject
	StructParent     _ebb.PdfObject
	ID               _ebb.PdfObject
	OPI              _ebb.PdfObject
	Metadata         _ebb.PdfObject
	OC               _ebb.PdfObject
	Stream           []byte
	_fbeec           *_ebb.PdfObjectStream
}

// NewPdfAnnotationRedact returns a new redact annotation.
func NewPdfAnnotationRedact() *PdfAnnotationRedact {
	_fcb := NewPdfAnnotation()
	_egc := &PdfAnnotationRedact{}
	_egc.PdfAnnotation = _fcb
	_egc.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fcb.SetContext(_egc)
	return _egc
}

// PdfShadingType5 is a Lattice-form Gouraud-shaded triangle mesh.
type PdfShadingType5 struct {
	*PdfShading
	BitsPerCoordinate *_ebb.PdfObjectInteger
	BitsPerComponent  *_ebb.PdfObjectInteger
	VerticesPerRow    *_ebb.PdfObjectInteger
	Decode            *_ebb.PdfObjectArray
	Function          []PdfFunction
}

// Insert adds a top level outline item in the outline,
// at the specified index.
func (_ebdcg *Outline) Insert(index uint, item *OutlineItem) {
	_bffc := uint(len(_ebdcg.Entries))
	if index > _bffc {
		index = _bffc
	}
	_ebdcg.Entries = append(_ebdcg.Entries[:index], append([]*OutlineItem{item}, _ebdcg.Entries[index:]...)...)
}

// NewPdfAnnotationUnderline returns a new text underline annotation.
func NewPdfAnnotationUnderline() *PdfAnnotationUnderline {
	_gec := NewPdfAnnotation()
	_fdea := &PdfAnnotationUnderline{}
	_fdea.PdfAnnotation = _gec
	_fdea.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gec.SetContext(_fdea)
	return _fdea
}

// SetPdfTitle sets the Title attribute of the output PDF.
func SetPdfTitle(title string) { _daddc.Lock(); defer _daddc.Unlock(); _fead = title }

// Optimizer is the interface that performs optimization of PDF object structure for output writing.
//
// Optimize receives a slice of input `objects`, performs optimization, including removing, replacing objects and
// output the optimized slice of objects.
type Optimizer interface {
	Optimize(_edead []_ebb.PdfObject) ([]_ebb.PdfObject, error)
}

// ToInteger convert to an integer format.
func (_gdcdc *PdfColorCalRGB) ToInteger(bits int) [3]uint32 {
	_fgfea := _cbg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_fgfea * _gdcdc.A()), uint32(_fgfea * _gdcdc.B()), uint32(_fgfea * _gdcdc.C())}
}
func (_gabgb *LTV) buildCertChain(_egdf, _bceafa []*_g.Certificate) ([]*_g.Certificate, map[string]*_g.Certificate, error) {
	_gdbcg := map[string]*_g.Certificate{}
	for _, _faege := range _egdf {
		_gdbcg[_faege.Subject.CommonName] = _faege
	}
	_bbbee := _egdf
	for _, _gfbgee := range _bceafa {
		_ggbf := _gfbgee.Subject.CommonName
		if _, _fbec := _gdbcg[_ggbf]; _fbec {
			continue
		}
		_gdbcg[_ggbf] = _gfbgee
		_bbbee = append(_bbbee, _gfbgee)
	}
	if len(_bbbee) == 0 {
		return nil, nil, ErrSignNoCertificates
	}
	var _afga error
	for _gbafd := _bbbee[0]; _gbafd != nil && !_gabgb.CertClient.IsCA(_gbafd); {
		_cggbe, _ccacd := _gdbcg[_gbafd.Issuer.CommonName]
		if !_ccacd {
			if _cggbe, _afga = _gabgb.CertClient.GetIssuer(_gbafd); _afga != nil {
				_eg.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _afga)
				break
			}
			_gdbcg[_gbafd.Issuer.CommonName] = _cggbe
			_bbbee = append(_bbbee, _cggbe)
		}
		_gbafd = _cggbe
	}
	return _bbbee, _gdbcg, nil
}

// StandardValidator is the interface that is used for the PDF StandardImplementer validation for the PDF document.
// It is using a CompliancePdfReader which is expected to give more Metadata during reading process.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
type StandardValidator interface {

	// ValidateStandard checks if the input reader
	ValidateStandard(_feebc *CompliancePdfReader) error
}

func (_bbfce *PdfPage) getParentResources() (*PdfPageResources, error) {
	_dafef := _bbfce.Parent
	for _dafef != nil {
		_dgff, _fggdd := _ebb.GetDict(_dafef)
		if !_fggdd {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020n\u006f\u0064\u0065")
			return nil, _gf.New("i\u006e\u0076\u0061\u006cid\u0020p\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		if _cagbg := _dgff.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _cagbg != nil {
			_bafdbbg, _gbebe := _ebb.GetDict(_cagbg)
			if !_gbebe {
				return nil, _gf.New("i\u006e\u0076\u0061\u006cid\u0020r\u0065\u0073\u006f\u0075\u0072c\u0065\u0020\u0064\u0069\u0063\u0074")
			}
			_addba, _eecgd := NewPdfPageResourcesFromDict(_bafdbbg)
			if _eecgd != nil {
				return nil, _eecgd
			}
			return _addba, nil
		}
		_dafef = _dgff.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, nil
}
func _aagg(_ddafe _ebb.PdfObject) (PdfFunction, error) {
	_ddafe = _ebb.ResolveReference(_ddafe)
	if _caeb, _gaca := _ddafe.(*_ebb.PdfObjectStream); _gaca {
		_gcfaf := _caeb.PdfObjectDictionary
		_fgbdbe, _dceaf := _gcfaf.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_ebb.PdfObjectInteger)
		if !_dceaf {
			_eg.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _gf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_fgbdbe == 0 {
			return _gfcg(_caeb)
		} else if *_fgbdbe == 4 {
			return _gbab(_caeb)
		} else {
			return nil, _gf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _eaeea, _cfecb := _ddafe.(*_ebb.PdfIndirectObject); _cfecb {
		_geea, _acag := _eaeea.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_acag {
			_eg.Log.Error("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006eg\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _gf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		_efff, _acag := _geea.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_ebb.PdfObjectInteger)
		if !_acag {
			_eg.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _gf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_efff == 2 {
			return _fafaa(_eaeea)
		} else if *_efff == 3 {
			return _ggbgd(_eaeea)
		} else {
			return nil, _gf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _baeb, _eccb := _ddafe.(*_ebb.PdfObjectDictionary); _eccb {
		_fgbce, _cdbef := _baeb.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_ebb.PdfObjectInteger)
		if !_cdbef {
			_eg.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _gf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_fgbce == 2 {
			return _fafaa(_baeb)
		} else if *_fgbce == 3 {
			return _ggbgd(_baeb)
		} else {
			return nil, _gf.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else {
		_eg.Log.Debug("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0054\u0079\u0070\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0023\u0076", _ddafe)
		return nil, _gf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
}

// EnableByName LTV enables the signature dictionary of the PDF AcroForm
// field identified the specified name. The signing certificate chain is
// extracted from the signature dictionary. Optionally, additional certificates
// can be specified through the `extraCerts` parameter. The LTV client attempts
// to build the certificate chain up to a trusted root by downloading any
// missing certificates.
func (_fdefe *LTV) EnableByName(name string, extraCerts []*_g.Certificate) error {
	_dbcbf := _fdefe._ggdbg._acfe.AcroForm
	for _, _adeb := range _dbcbf.AllFields() {
		_ffca, _ := _adeb.GetContext().(*PdfFieldSignature)
		if _ffca == nil {
			continue
		}
		if _aaedg := _ffca.PartialName(); _aaedg != name {
			continue
		}
		return _fdefe.Enable(_ffca.V, extraCerts)
	}
	return nil
}
func _gegg(_effe *fontCommon) *pdfFontType0 { return &pdfFontType0{fontCommon: *_effe} }

// PdfShadingType1 is a Function-based shading.
type PdfShadingType1 struct {
	*PdfShading
	Domain   *_ebb.PdfObjectArray
	Matrix   *_ebb.PdfObjectArray
	Function []PdfFunction
}

func (_ggba *PdfReader) loadAction(_eeb _ebb.PdfObject) (*PdfAction, error) {
	if _geed, _eegg := _ebb.GetIndirect(_eeb); _eegg {
		_bff, _efe := _ggba.newPdfActionFromIndirectObject(_geed)
		if _efe != nil {
			return nil, _efe
		}
		return _bff, nil
	} else if !_ebb.IsNullObject(_eeb) {
		return nil, _gf.New("\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074 \u0074\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, nil
}

// IsSimple returns true if `font` is a simple font.
func (_abdd *PdfFont) IsSimple() bool { _, _gbba := _abdd._ebcad.(*pdfFontSimple); return _gbba }

var _gdcdfg = map[string]struct{}{"\u0046\u0054": {}, "\u004b\u0069\u0064\u0073": {}, "\u0054": {}, "\u0054\u0055": {}, "\u0054\u004d": {}, "\u0046\u0066": {}, "\u0056": {}, "\u0044\u0056": {}, "\u0041\u0041": {}, "\u0044\u0041": {}, "\u0051": {}, "\u0044\u0053": {}, "\u0052\u0056": {}}

// GetAction returns the PDF action for the annotation link.
func (_fdd *PdfAnnotationLink) GetAction() (*PdfAction, error) {
	if _fdd._ffea != nil {
		return _fdd._ffea, nil
	}
	if _fdd.A == nil {
		return nil, nil
	}
	if _fdd._cfbg == nil {
		return nil, nil
	}
	_beb, _dagf := _fdd._cfbg.loadAction(_fdd.A)
	if _dagf != nil {
		return nil, _dagf
	}
	_fdd._ffea = _beb
	return _fdd._ffea, nil
}

// NewPdfActionSubmitForm returns a new "submit form" action.
func NewPdfActionSubmitForm() *PdfActionSubmitForm {
	_bdc := NewPdfAction()
	_ccb := &PdfActionSubmitForm{}
	_ccb.PdfAction = _bdc
	_bdc.SetContext(_ccb)
	return _ccb
}

// IsCID returns true if the underlying font is CID.
func (_egeda *PdfFont) IsCID() bool { return _egeda.baseFields().isCIDFont() }

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
func (_fgfd *Image) Resample(targetBitsPerComponent int64) {
	if _fgfd.BitsPerComponent == targetBitsPerComponent {
		return
	}
	_dfbdb := _fgfd.GetSamples()
	if targetBitsPerComponent < _fgfd.BitsPerComponent {
		_faec := _fgfd.BitsPerComponent - targetBitsPerComponent
		for _bcccb := range _dfbdb {
			_dfbdb[_bcccb] >>= uint(_faec)
		}
	} else if targetBitsPerComponent > _fgfd.BitsPerComponent {
		_efdcg := targetBitsPerComponent - _fgfd.BitsPerComponent
		for _cbebb := range _dfbdb {
			_dfbdb[_cbebb] <<= uint(_efdcg)
		}
	}
	_fgfd.BitsPerComponent = targetBitsPerComponent
	if _fgfd.BitsPerComponent < 8 {
		_fgfd.resampleLowBits(_dfbdb)
		return
	}
	_cedcc := _dg.BytesPerLine(int(_fgfd.Width), int(_fgfd.BitsPerComponent), _fgfd.ColorComponents)
	_eecac := make([]byte, _cedcc*int(_fgfd.Height))
	var (
		_eeeag, _gcee, _deecg, _cafebf int
		_cfbce                         uint32
	)
	for _deecg = 0; _deecg < int(_fgfd.Height); _deecg++ {
		_eeeag = _deecg * _cedcc
		_gcee = (_deecg+1)*_cedcc - 1
		_dbcad := _abg.ResampleUint32(_dfbdb[_eeeag:_gcee], int(targetBitsPerComponent), 8)
		for _cafebf, _cfbce = range _dbcad {
			_eecac[_cafebf+_eeeag] = byte(_cfbce)
		}
	}
	_fgfd.Data = _eecac
}

// SetContentStream updates the content stream with specified encoding.
// If encoding is null, will use the xform.Filter object or Raw encoding if not set.
func (_eagdc *XObjectForm) SetContentStream(content []byte, encoder _ebb.StreamEncoder) error {
	_adefa := content
	if encoder == nil {
		if _eagdc.Filter != nil {
			encoder = _eagdc.Filter
		} else {
			encoder = _ebb.NewRawEncoder()
		}
	}
	_affbcb, _degcfg := encoder.EncodeBytes(_adefa)
	if _degcfg != nil {
		return _degcfg
	}
	_adefa = _affbcb
	_eagdc.Stream = _adefa
	_eagdc.Filter = encoder
	return nil
}

// XObjectForm (Table 95 in 8.10.2).
type XObjectForm struct {
	Filter        _ebb.StreamEncoder
	FormType      _ebb.PdfObject
	BBox          _ebb.PdfObject
	Matrix        _ebb.PdfObject
	Resources     *PdfPageResources
	Group         _ebb.PdfObject
	Ref           _ebb.PdfObject
	MetaData      _ebb.PdfObject
	PieceInfo     _ebb.PdfObject
	LastModified  _ebb.PdfObject
	StructParent  _ebb.PdfObject
	StructParents _ebb.PdfObject
	OPI           _ebb.PdfObject
	OC            _ebb.PdfObject
	Name          _ebb.PdfObject

	// Stream data.
	Stream []byte
	_gebcd *_ebb.PdfObjectStream
}

// PdfDate represents a date, which is a PDF string of the form:
// (D:YYYYMMDDHHmmSSOHH'mm)
type PdfDate struct {
	_dacdd  int64
	_agaba  int64
	_edcfe  int64
	_aedbe  int64
	_cfaba  int64
	_cgbcd  int64
	_degeb  byte
	_gcccb  int64
	_fddcbg int64
}

// Encrypt encrypts the output file with a specified user/owner password.
func (_aagad *PdfWriter) Encrypt(userPass, ownerPass []byte, options *EncryptOptions) error {
	_fcbaf := RC4_128bit
	if options != nil {
		_fcbaf = options.Algorithm
	}
	_aaebb := _fe.PermOwner
	if options != nil {
		_aaebb = options.Permissions
	}
	var _egad _fa.Filter
	switch _fcbaf {
	case RC4_128bit:
		_egad = _fa.NewFilterV2(16)
	case AES_128bit:
		_egad = _fa.NewFilterAESV2()
	case AES_256bit:
		_egad = _fa.NewFilterAESV3()
	default:
		return _bg.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020a\u006cg\u006fr\u0069\u0074\u0068\u006d\u003a\u0020\u0025v", options.Algorithm)
	}
	_gbgba, _cfffd, _dbfba := _ebb.PdfCryptNewEncrypt(_egad, userPass, ownerPass, _aaebb)
	if _dbfba != nil {
		return _dbfba
	}
	_aagad._cgfde = _gbgba
	if _cfffd.Major != 0 {
		_aagad.SetVersion(_cfffd.Major, _cfffd.Minor)
	}
	_aagad._gaccf = _cfffd.Encrypt
	_aagad._gfdea, _aagad._gffb = _cfffd.ID0, _cfffd.ID1
	_bbecb := _ebb.MakeIndirectObject(_cfffd.Encrypt)
	_aagad._cbcaa = _bbecb
	_aagad.addObject(_bbecb)
	return nil
}

// HasXObjectByName checks if has XObject resource by name.
func (_eafd *PdfPage) HasXObjectByName(name _ebb.PdfObjectName) bool {
	_ebece, _dffcb := _eafd.Resources.XObject.(*_ebb.PdfObjectDictionary)
	if !_dffcb {
		return false
	}
	if _eegb := _ebece.Get(name); _eegb != nil {
		return true
	}
	return false
}
func (_dbdg *PdfReader) newPdfActionNamedFromDict(_fdef *_ebb.PdfObjectDictionary) (*PdfActionNamed, error) {
	return &PdfActionNamed{N: _fdef.Get("\u004e")}, nil
}
func _eega(_ebeg *fontCommon) *pdfFontType3 { return &pdfFontType3{fontCommon: *_ebeg} }

// Fill populates `form` with values provided by `provider`.
func (_dfgge *PdfAcroForm) Fill(provider FieldValueProvider) error { return _dfgge.fill(provider, nil) }
func _ggdfc(_agcdd _ebb.PdfObject) (*PdfShading, error) {
	_cfgad := &PdfShading{}
	var _caaad *_ebb.PdfObjectDictionary
	if _aced, _dbddf := _ebb.GetIndirect(_agcdd); _dbddf {
		_cfgad._fbfae = _aced
		_eabfb, _gbdfe := _aced.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_gbdfe {
			_eg.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _ebb.ErrTypeError
		}
		_caaad = _eabfb
	} else if _cffdc, _ebeaf := _ebb.GetStream(_agcdd); _ebeaf {
		_cfgad._fbfae = _cffdc
		_caaad = _cffdc.PdfObjectDictionary
	} else if _befef, _cfabb := _ebb.GetDict(_agcdd); _cfabb {
		_cfgad._fbfae = _befef
		_caaad = _befef
	} else {
		_eg.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", _agcdd)
		return nil, _ebb.ErrTypeError
	}
	if _caaad == nil {
		_eg.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _gf.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_agcdd = _caaad.Get("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _agcdd == nil {
		_eg.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073si\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_agcdd = _ebb.TraceToDirectObject(_agcdd)
	_dafea, _efcf := _agcdd.(*_ebb.PdfObjectInteger)
	if !_efcf {
		_eg.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066o\u0072 \u0073h\u0061d\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _agcdd)
		return nil, _ebb.ErrTypeError
	}
	if *_dafea < 1 || *_dafea > 7 {
		_eg.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u006e\u006ft\u0020\u0031\u002d\u0037\u0020(\u0067\u006ft\u0020\u0025\u0064\u0029", *_dafea)
		return nil, _ebb.ErrTypeError
	}
	_cfgad.ShadingType = _dafea
	_agcdd = _caaad.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")
	if _agcdd == nil {
		_eg.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0020\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065\u0020e\u006e\u0074\u0072\u0079\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_abeef, _ecdcf := NewPdfColorspaceFromPdfObject(_agcdd)
	if _ecdcf != nil {
		_eg.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065: \u0025\u0076", _ecdcf)
		return nil, _ecdcf
	}
	_cfgad.ColorSpace = _abeef
	_agcdd = _caaad.Get("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064")
	if _agcdd != nil {
		_agcdd = _ebb.TraceToDirectObject(_agcdd)
		_decab, _gegeea := _agcdd.(*_ebb.PdfObjectArray)
		if !_gegeea {
			_eg.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _agcdd)
			return nil, _ebb.ErrTypeError
		}
		_cfgad.Background = _decab
	}
	_agcdd = _caaad.Get("\u0042\u0042\u006f\u0078")
	if _agcdd != nil {
		_agcdd = _ebb.TraceToDirectObject(_agcdd)
		_fggc, _abffc := _agcdd.(*_ebb.PdfObjectArray)
		if !_abffc {
			_eg.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _agcdd)
			return nil, _ebb.ErrTypeError
		}
		_bgbad, _agacf := NewPdfRectangle(*_fggc)
		if _agacf != nil {
			_eg.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _agacf)
			return nil, _agacf
		}
		_cfgad.BBox = _bgbad
	}
	_agcdd = _caaad.Get("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s")
	if _agcdd != nil {
		_agcdd = _ebb.TraceToDirectObject(_agcdd)
		_baada, _ddggg := _agcdd.(*_ebb.PdfObjectBool)
		if !_ddggg {
			_eg.Log.Debug("A\u006e\u0074\u0069\u0041\u006c\u0069\u0061\u0073\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020ty\u0070\u0065\u002c\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020bo\u006f\u006c \u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _agcdd)
			return nil, _ebb.ErrTypeError
		}
		_cfgad.AntiAlias = _baada
	}
	switch *_dafea {
	case 1:
		_ddgcf, _cdcb := _adaaf(_caaad)
		if _cdcb != nil {
			return nil, _cdcb
		}
		_ddgcf.PdfShading = _cfgad
		_cfgad._edgag = _ddgcf
		return _cfgad, nil
	case 2:
		_fdcbg, _bbdda := _adcga(_caaad)
		if _bbdda != nil {
			return nil, _bbdda
		}
		_fdcbg.PdfShading = _cfgad
		_cfgad._edgag = _fdcbg
		return _cfgad, nil
	case 3:
		_eadbac, _cgabg := _bcafa(_caaad)
		if _cgabg != nil {
			return nil, _cgabg
		}
		_eadbac.PdfShading = _cfgad
		_cfgad._edgag = _eadbac
		return _cfgad, nil
	case 4:
		_fcbc, _befag := _ccgda(_caaad)
		if _befag != nil {
			return nil, _befag
		}
		_fcbc.PdfShading = _cfgad
		_cfgad._edgag = _fcbc
		return _cfgad, nil
	case 5:
		_fdabbf, _gggdd := _ebcfc(_caaad)
		if _gggdd != nil {
			return nil, _gggdd
		}
		_fdabbf.PdfShading = _cfgad
		_cfgad._edgag = _fdabbf
		return _cfgad, nil
	case 6:
		_afbed, _gaga := _gcabc(_caaad)
		if _gaga != nil {
			return nil, _gaga
		}
		_afbed.PdfShading = _cfgad
		_cfgad._edgag = _afbed
		return _cfgad, nil
	case 7:
		_ddagg, _bbccc := _gaab(_caaad)
		if _bbccc != nil {
			return nil, _bbccc
		}
		_ddagg.PdfShading = _cfgad
		_cfgad._edgag = _ddagg
		return _cfgad, nil
	}
	return nil, _gf.New("u\u006ek\u006e\u006f\u0077\u006e\u0020\u0073\u0068\u0061d\u0069\u006e\u0067\u0020ty\u0070\u0065")
}

// GetContainingPdfObject implements interface PdfModel.
func (_ddegcb *PdfSignatureReference) GetContainingPdfObject() _ebb.PdfObject { return _ddegcb._afacea }

// WriteString outputs the object as it is to be written to file.
func (_ebcde *PdfTransformParamsDocMDP) WriteString() string {
	return _ebcde.ToPdfObject().WriteString()
}
func (_facce *PdfReader) newPdfAcroFormFromDict(_aaddd *_ebb.PdfIndirectObject, _fddd *_ebb.PdfObjectDictionary) (*PdfAcroForm, error) {
	_dddd := NewPdfAcroForm()
	if _aaddd != nil {
		_dddd._adcg = _aaddd
		_aaddd.PdfObject = _ebb.MakeDict()
	}
	if _bfcc := _fddd.Get("\u0046\u0069\u0065\u006c\u0064\u0073"); _bfcc != nil && !_ebb.IsNullObject(_bfcc) {
		_gcca, _cdaa := _ebb.GetArray(_bfcc)
		if !_cdaa {
			return nil, _bg.Errorf("\u0066i\u0065\u006c\u0064\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _bfcc)
		}
		var _ebdca []*PdfField
		for _, _aefaa := range _gcca.Elements() {
			_aagcc, _ffaf := _ebb.GetIndirect(_aefaa)
			if !_ffaf {
				if _, _bffd := _aefaa.(*_ebb.PdfObjectNull); _bffd {
					_eg.Log.Trace("\u0053k\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072 \u006e\u0075\u006c\u006c\u0020\u0066\u0069\u0065\u006c\u0064")
					continue
				}
				_eg.Log.Debug("\u0046\u0069\u0065\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0064 \u0069\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _aefaa)
				return nil, _bg.Errorf("\u0066\u0069\u0065l\u0064\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_aeagg, _bdabd := _facce.newPdfFieldFromIndirectObject(_aagcc, nil)
			if _bdabd != nil {
				return nil, _bdabd
			}
			_eg.Log.Trace("\u0041\u0063\u0072\u006fFo\u0072\u006d\u0020\u0046\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u002b\u0076", *_aeagg)
			_ebdca = append(_ebdca, _aeagg)
		}
		_dddd.Fields = &_ebdca
	}
	if _abggd := _fddd.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _abggd != nil {
		_bgegg, _gaebc := _ebb.GetBool(_abggd)
		if _gaebc {
			_dddd.NeedAppearances = _bgegg
		} else {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u0065\u0065\u0064\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063e\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _abggd)
		}
	}
	if _eggdf := _fddd.Get("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073"); _eggdf != nil {
		_ffgea, _bfca := _ebb.GetInt(_eggdf)
		if _bfca {
			_dddd.SigFlags = _ffgea
		} else {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073 \u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _eggdf)
		}
	}
	if _dfdag := _fddd.Get("\u0043\u004f"); _dfdag != nil {
		_feaa, _beaaf := _ebb.GetArray(_dfdag)
		if _beaaf {
			_dddd.CO = _feaa
		} else {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u004f\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _dfdag)
		}
	}
	if _fccgg := _fddd.Get("\u0044\u0052"); _fccgg != nil {
		if _cdffc, _gccgb := _ebb.GetDict(_fccgg); _gccgb {
			_caedda, _cgcgb := NewPdfPageResourcesFromDict(_cdffc)
			if _cgcgb != nil {
				_eg.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0044R\u003a\u0020\u0025\u0076", _cgcgb)
				return nil, _cgcgb
			}
			_dddd.DR = _caedda
		} else {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0052\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _fccgg)
		}
	}
	if _egfec := _fddd.Get("\u0044\u0041"); _egfec != nil {
		_cacf, _ggcg := _ebb.GetString(_egfec)
		if _ggcg {
			_dddd.DA = _cacf
		} else {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0041\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _egfec)
		}
	}
	if _afdbd := _fddd.Get("\u0051"); _afdbd != nil {
		_dafda, _gcffc := _ebb.GetInt(_afdbd)
		if _gcffc {
			_dddd.Q = _dafda
		} else {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0051\u0020\u0069\u006e\u0076a\u006ci\u0064 \u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _afdbd)
		}
	}
	if _dcdfg := _fddd.Get("\u0058\u0046\u0041"); _dcdfg != nil {
		_dddd.XFA = _dcdfg
	}
	_dddd.ToPdfObject()
	return _dddd, nil
}

// ToPdfObject returns the PDF representation of the function.
func (_fcfee *PdfFunctionType2) ToPdfObject() _ebb.PdfObject {
	_bdgb := _ebb.MakeDict()
	_bdgb.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _ebb.MakeInteger(2))
	_bgadf := &_ebb.PdfObjectArray{}
	for _, _aggbd := range _fcfee.Domain {
		_bgadf.Append(_ebb.MakeFloat(_aggbd))
	}
	_bdgb.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _bgadf)
	if _fcfee.Range != nil {
		_befge := &_ebb.PdfObjectArray{}
		for _, _egcgc := range _fcfee.Range {
			_befge.Append(_ebb.MakeFloat(_egcgc))
		}
		_bdgb.Set("\u0052\u0061\u006eg\u0065", _befge)
	}
	if _fcfee.C0 != nil {
		_gcbb := &_ebb.PdfObjectArray{}
		for _, _ebab := range _fcfee.C0 {
			_gcbb.Append(_ebb.MakeFloat(_ebab))
		}
		_bdgb.Set("\u0043\u0030", _gcbb)
	}
	if _fcfee.C1 != nil {
		_bfdge := &_ebb.PdfObjectArray{}
		for _, _gfebc := range _fcfee.C1 {
			_bfdge.Append(_ebb.MakeFloat(_gfebc))
		}
		_bdgb.Set("\u0043\u0031", _bfdge)
	}
	_bdgb.Set("\u004e", _ebb.MakeFloat(_fcfee.N))
	if _fcfee._bcbfd != nil {
		_fcfee._bcbfd.PdfObject = _bdgb
		return _fcfee._bcbfd
	}
	return _bdgb
}
func _fgag(_cafgd map[_bad.GID]int, _ggfb uint16) *_ebb.PdfObjectArray {
	_ddbd := &_ebb.PdfObjectArray{}
	_aeae := _bad.GID(_ggfb)
	for _gbdc := _bad.GID(0); _gbdc < _aeae; {
		_cceg, _gdcb := _cafgd[_gbdc]
		if !_gdcb {
			_gbdc++
			continue
		}
		_afbb := _gbdc
		for _cbdda := _afbb + 1; _cbdda < _aeae; _cbdda++ {
			if _aefe, _begfa := _cafgd[_cbdda]; !_begfa || _cceg != _aefe {
				break
			}
			_afbb = _cbdda
		}
		_ddbd.Append(_ebb.MakeInteger(int64(_gbdc)))
		_ddbd.Append(_ebb.MakeInteger(int64(_afbb)))
		_ddbd.Append(_ebb.MakeInteger(int64(_cceg)))
		_gbdc = _afbb + 1
	}
	return _ddbd
}

// PdfShadingPattern is a Shading patterns that provide a smooth transition between colors across an area to be painted,
// i.e. color(x,y) = f(x,y) at each point.
// It is a type 2 pattern (PatternType = 2).
type PdfShadingPattern struct {
	*PdfPattern
	Shading   *PdfShading
	Matrix    *_ebb.PdfObjectArray
	ExtGState _ebb.PdfObject
}

// RemovePage removes a page by number.
func (_faca *PdfAppender) RemovePage(pageNum int) {
	_bcdg := pageNum - 1
	_faca._dfbg = append(_faca._dfbg[0:_bcdg], _faca._dfbg[pageNum:]...)
}
func (_ggbfd *PdfWriter) writeObject(_abegg int, _abbf _ebb.PdfObject) {
	_eg.Log.Trace("\u0057\u0072\u0069\u0074\u0065\u0020\u006f\u0062\u006a \u0023\u0025\u0064\u000a", _abegg)
	if _ebacf, _gabad := _abbf.(*_ebb.PdfIndirectObject); _gabad {
		_ggbfd._bedfc[_abegg] = crossReference{Type: 1, Offset: _ggbfd._afedd, Generation: _ebacf.GenerationNumber}
		_ggbdb := _bg.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _abegg)
		if _eaaagg, _dffga := _ebacf.PdfObject.(*pdfSignDictionary); _dffga {
			_eaaagg._cbfg = _ggbfd._afedd + int64(len(_ggbdb))
		}
		if _ebacf.PdfObject == nil {
			_eg.Log.Debug("E\u0072\u0072\u006fr\u003a\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0027\u0073\u0020\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020b\u0065\u0020\u006e\u0069l\u0020\u002d\u0020\u0073e\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063t\u004e\u0075\u006c\u006c")
			_ebacf.PdfObject = _ebb.MakeNull()
		}
		_ggbdb += _ebacf.PdfObject.WriteString()
		_ggbdb += "\u000a\u0065\u006e\u0064\u006f\u0062\u006a\u000a"
		_ggbfd.writeString(_ggbdb)
		return
	}
	if _fbegb, _ageead := _abbf.(*_ebb.PdfObjectStream); _ageead {
		_ggbfd._bedfc[_abegg] = crossReference{Type: 1, Offset: _ggbfd._afedd, Generation: _fbegb.GenerationNumber}
		_afagc := _bg.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _abegg)
		_afagc += _fbegb.PdfObjectDictionary.WriteString()
		_afagc += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_ggbfd.writeString(_afagc)
		_ggbfd.writeBytes(_fbegb.Stream)
		_ggbfd.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	if _feaef, _caeac := _abbf.(*_ebb.PdfObjectStreams); _caeac {
		_ggbfd._bedfc[_abegg] = crossReference{Type: 1, Offset: _ggbfd._afedd, Generation: _feaef.GenerationNumber}
		_fcdfg := _bg.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _abegg)
		var _fgcge []string
		var _gbbgf string
		var _fgddda int64
		for _fcaee, _dadcc := range _feaef.Elements() {
			_fbfb, _deebe := _dadcc.(*_ebb.PdfIndirectObject)
			if !_deebe {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065am\u0073 \u004e\u0020\u0025\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006es\u0020\u006e\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u0070\u0064\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0076", _abegg, _dadcc)
				continue
			}
			_bcgdb := _fbfb.PdfObject.WriteString() + "\u0020"
			_gbbgf = _gbbgf + _bcgdb
			_fgcge = append(_fgcge, _bg.Sprintf("\u0025\u0064\u0020%\u0064", _fbfb.ObjectNumber, _fgddda))
			_ggbfd._bedfc[int(_fbfb.ObjectNumber)] = crossReference{Type: 2, ObjectNumber: _abegg, Index: _fcaee}
			_fgddda = _fgddda + int64(len([]byte(_bcgdb)))
		}
		_eccf := _ee.Join(_fgcge, "\u0020") + "\u0020"
		_dbaab := _ebb.NewFlateEncoder()
		_bdbac := _dbaab.MakeStreamDict()
		_bdbac.Set(_ebb.PdfObjectName("\u0054\u0079\u0070\u0065"), _ebb.MakeName("\u004f\u0062\u006a\u0053\u0074\u006d"))
		_bgeae := int64(_feaef.Len())
		_bdbac.Set(_ebb.PdfObjectName("\u004e"), _ebb.MakeInteger(_bgeae))
		_bfgece := int64(len(_eccf))
		_bdbac.Set(_ebb.PdfObjectName("\u0046\u0069\u0072s\u0074"), _ebb.MakeInteger(_bfgece))
		_gcafe, _ := _dbaab.EncodeBytes([]byte(_eccf + _gbbgf))
		_adcd := int64(len(_gcafe))
		_bdbac.Set(_ebb.PdfObjectName("\u004c\u0065\u006e\u0067\u0074\u0068"), _ebb.MakeInteger(_adcd))
		_fcdfg += _bdbac.WriteString()
		_fcdfg += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_ggbfd.writeString(_fcdfg)
		_ggbfd.writeBytes(_gcafe)
		_ggbfd.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	_ggbfd.writeString(_abbf.WriteString())
}
func (_dafc fontCommon) coreString() string {
	_bfgb := ""
	if _dafc._fbbd != nil {
		_bfgb = _dafc._fbbd.String()
	}
	return _bg.Sprintf("\u0025#\u0071\u0020%\u0023\u0071\u0020%\u0071\u0020\u006f\u0062\u006a\u003d\u0025d\u0020\u0054\u006f\u0055\u006e\u0069c\u006f\u0064\u0065\u003d\u0025\u0074\u0020\u0066\u006c\u0061\u0067s\u003d\u0030\u0078\u0025\u0030\u0078\u0020\u0025\u0073", _dafc._dfbf, _dafc._fdacg, _dafc._efge, _dafc._efbg, _dafc._baag != nil, _dafc.fontFlags(), _bfgb)
}

// NewPdfAnnotationStrikeOut returns a new text strikeout annotation.
func NewPdfAnnotationStrikeOut() *PdfAnnotationStrikeOut {
	_acg := NewPdfAnnotation()
	_cgb := &PdfAnnotationStrikeOut{}
	_cgb.PdfAnnotation = _acg
	_cgb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_acg.SetContext(_cgb)
	return _cgb
}

var _ pdfFont = (*pdfCIDFontType0)(nil)

// PdfShadingType3 is a Radial shading.
type PdfShadingType3 struct {
	*PdfShading
	Coords   *_ebb.PdfObjectArray
	Domain   *_ebb.PdfObjectArray
	Function []PdfFunction
	Extend   *_ebb.PdfObjectArray
}

// ToPdfObject implements interface PdfModel.
func (_gda *PdfAction) ToPdfObject() _ebb.PdfObject {
	_fb := _gda._abe
	_baa := _fb.PdfObject.(*_ebb.PdfObjectDictionary)
	_baa.Clear()
	_baa.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0041\u0063\u0074\u0069\u006f\u006e"))
	_baa.SetIfNotNil("\u0053", _gda.S)
	_baa.SetIfNotNil("\u004e\u0065\u0078\u0074", _gda.Next)
	return _fb
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_defd *PdfColorspaceDeviceCMYK) ToPdfObject() _ebb.PdfObject {
	return _ebb.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b")
}

// GetNumComponents returns the number of color components (1 for Indexed).
func (_ccag *PdfColorspaceSpecialIndexed) GetNumComponents() int { return 1 }

// ToPdfObject converts the PdfFont object to its PDF representation.
func (_addb *PdfFont) ToPdfObject() _ebb.PdfObject {
	if _addb._ebcad == nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0066\u006f\u006e\u0074 \u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return _ebb.MakeNull()
	}
	return _addb._ebcad.ToPdfObject()
}

// CompliancePdfReader is a wrapper over PdfReader that is used for verifying if the input Pdf document matches the
// compliance rules of standards like PDF/A.
// NOTE: This implementation is in experimental development state.
// 	Keep in mind that it might change in the subsequent minor versions.
type CompliancePdfReader struct {
	*PdfReader
	_ecbbg _ebb.ParserMetadata
}

func (_fggde *PdfField) inherit(_gfcac func(*PdfField) bool) (bool, error) {
	_eaeba := map[*PdfField]bool{}
	_affea := false
	_becf := _fggde
	for _becf != nil {
		if _, _acfcd := _eaeba[_becf]; _acfcd {
			return false, _gf.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		_eebcb := _gfcac(_becf)
		if _eebcb {
			_affea = true
			break
		}
		_eaeba[_becf] = true
		_becf = _becf.Parent
	}
	return _affea, nil
}

// GetAnnotations returns the list of page annotations for `page`. If not loaded attempts to load the
// annotations, otherwise returns the loaded list.
func (_eefa *PdfPage) GetAnnotations() ([]*PdfAnnotation, error) {
	if _eefa._bbfed != nil {
		return _eefa._bbfed, nil
	}
	if _eefa.Annots == nil {
		_eefa._bbfed = []*PdfAnnotation{}
		return nil, nil
	}
	if _eefa._ddab == nil {
		_eefa._bbfed = []*PdfAnnotation{}
		return nil, nil
	}
	_bdee, _agbgb := _eefa._ddab.loadAnnotations(_eefa.Annots)
	if _agbgb != nil {
		return nil, _agbgb
	}
	if _bdee == nil {
		_eefa._bbfed = []*PdfAnnotation{}
	}
	_eefa._bbfed = _bdee
	return _eefa._bbfed, nil
}
func _eeafg(_adgga _ebb.PdfObject) (*PdfPageResourcesColorspaces, error) {
	_dgffa := &PdfPageResourcesColorspaces{}
	if _fbeed, _fdffd := _adgga.(*_ebb.PdfIndirectObject); _fdffd {
		_dgffa._ddffd = _fbeed
		_adgga = _fbeed.PdfObject
	}
	_cccgg, _degb := _ebb.GetDict(_adgga)
	if !_degb {
		return nil, _gf.New("\u0043\u0053\u0020at\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_dgffa.Names = []string{}
	_dgffa.Colorspaces = map[string]PdfColorspace{}
	for _, _bgdg := range _cccgg.Keys() {
		_efgf := _cccgg.Get(_bgdg)
		_dgffa.Names = append(_dgffa.Names, string(_bgdg))
		_aabd, _fcgda := NewPdfColorspaceFromPdfObject(_efgf)
		if _fcgda != nil {
			return nil, _fcgda
		}
		_dgffa.Colorspaces[string(_bgdg)] = _aabd
	}
	return _dgffa, nil
}

// GetXObjectByName gets XObject by name.
func (_eeggg *PdfPage) GetXObjectByName(name _ebb.PdfObjectName) (_ebb.PdfObject, bool) {
	_gfgcd, _ffcfe := _eeggg.Resources.XObject.(*_ebb.PdfObjectDictionary)
	if !_ffcfe {
		return nil, false
	}
	if _gbacbb := _gfgcd.Get(name); _gbacbb != nil {
		return _gbacbb, true
	}
	return nil, false
}

// GetNumComponents returns the number of color components (3 for Lab).
func (_abfac *PdfColorLab) GetNumComponents() int { return 3 }
func (_egcedc *pdfFontSimple) addEncoding() error {
	var (
		_abaf  string
		_aadef map[_da.CharCode]_da.GlyphName
		_cfgab _da.SimpleEncoder
	)
	if _egcedc.Encoder() != nil {
		_bcgfcc, _gfaec := _egcedc.Encoder().(_da.SimpleEncoder)
		if _gfaec && _bcgfcc != nil {
			_abaf = _bcgfcc.BaseName()
		}
	}
	if _egcedc.Encoding != nil {
		_bafff, _affa, _aefa := _egcedc.getFontEncoding()
		if _aefa != nil {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0073\u0065F\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0053u\u0062t\u0079\u0070\u0065\u003d\u0025\u0071\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003d\u0025\u0073 \u0028\u0025\u0054\u0029\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _egcedc._fdacg, _egcedc._dfbf, _egcedc.Encoding, _egcedc.Encoding, _aefa)
			return _aefa
		}
		if _bafff != "" {
			_abaf = _bafff
		}
		_aadef = _affa
		_cfgab, _aefa = _da.NewSimpleTextEncoder(_abaf, _aadef)
		if _aefa != nil {
			return _aefa
		}
	}
	if _cfgab == nil {
		_ecdbf := _egcedc._fbbd
		if _ecdbf != nil {
			switch _egcedc._dfbf {
			case "\u0054\u0079\u0070e\u0031":
				if _ecdbf.fontFile != nil && _ecdbf.fontFile._gega != nil {
					_eg.Log.Debug("\u0055\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006et\u0046\u0069\u006c\u0065")
					_cfgab = _ecdbf.fontFile._gega
				}
			case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
				if _ecdbf._aeeb != nil {
					_eg.Log.Debug("\u0055s\u0069n\u0067\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032")
					_fffgb, _cfgag := _ecdbf._aeeb.MakeEncoder()
					if _cfgag == nil {
						_cfgab = _fffgb
					}
				}
			}
		}
	}
	if _cfgab != nil {
		if _aadef != nil {
			_eg.Log.Trace("\u0064\u0069\u0066fe\u0072\u0065\u006e\u0063\u0065\u0073\u003d\u0025\u002b\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _aadef, _egcedc.baseFields())
			_cfgab = _da.ApplyDifferences(_cfgab, _aadef)
		}
		_egcedc.SetEncoder(_cfgab)
	}
	return nil
}

// ToPdfObject converts date to a PDF string object.
func (_agag *PdfDate) ToPdfObject() _ebb.PdfObject {
	_gafdf := _bg.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _agag._dacdd, _agag._agaba, _agag._edcfe, _agag._aedbe, _agag._cfaba, _agag._cgbcd, _agag._degeb, _agag._gcccb, _agag._fddcbg)
	return _ebb.MakeString(_gafdf)
}

// BorderStyle defines border type, typically used for annotations.
type BorderStyle int

// ToPdfObject implements interface PdfModel.
func (_baeeb *PdfAnnotationPrinterMark) ToPdfObject() _ebb.PdfObject {
	_baeeb.PdfAnnotation.ToPdfObject()
	_ecga := _baeeb._bdcd
	_ecd := _ecga.PdfObject.(*_ebb.PdfObjectDictionary)
	_ecd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b"))
	_ecd.SetIfNotNil("\u004d\u004e", _baeeb.MN)
	return _ecga
}
func (_aaga *PdfReader) loadAnnotations(_abade _ebb.PdfObject) ([]*PdfAnnotation, error) {
	_eddd, _bgcf := _ebb.GetArray(_abade)
	if !_bgcf {
		return nil, _bg.Errorf("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	var _efffg []*PdfAnnotation
	for _, _dbagdc := range _eddd.Elements() {
		_dbagdc = _ebb.ResolveReference(_dbagdc)
		if _, _dfeb := _dbagdc.(*_ebb.PdfObjectNull); _dfeb {
			continue
		}
		_befbc, _beebg := _dbagdc.(*_ebb.PdfObjectDictionary)
		_aabca, _aacaa := _dbagdc.(*_ebb.PdfIndirectObject)
		if _beebg {
			_aabca = &_ebb.PdfIndirectObject{}
			_aabca.PdfObject = _befbc
		} else {
			if !_aacaa {
				return nil, _bg.Errorf("\u0061\u006eno\u0074\u0061\u0074i\u006f\u006e\u0020\u006eot \u0069n \u0061\u006e\u0020\u0069\u006e\u0064\u0069re\u0063\u0074\u0020\u006f\u0062\u006a\u0065c\u0074")
			}
		}
		_cbedb, _gdeff := _aaga.newPdfAnnotationFromIndirectObject(_aabca)
		if _gdeff != nil {
			return nil, _gdeff
		}
		switch _bfefa := _cbedb.GetContext().(type) {
		case *PdfAnnotationWidget:
			for _, _dffca := range _aaga.AcroForm.AllFields() {
				if _dffca._cdfd == _bfefa.Parent {
					_bfefa._gce = _dffca
					break
				}
			}
		}
		if _cbedb != nil {
			_efffg = append(_efffg, _cbedb)
		}
	}
	return _efffg, nil
}
func (_caef *PdfReader) newPdfAnnotationTextFromDict(_gcg *_ebb.PdfObjectDictionary) (*PdfAnnotationText, error) {
	_acd := PdfAnnotationText{}
	_aebd, _dcf := _caef.newPdfAnnotationMarkupFromDict(_gcg)
	if _dcf != nil {
		return nil, _dcf
	}
	_acd.PdfAnnotationMarkup = _aebd
	_acd.Open = _gcg.Get("\u004f\u0070\u0065\u006e")
	_acd.Name = _gcg.Get("\u004e\u0061\u006d\u0065")
	_acd.State = _gcg.Get("\u0053\u0074\u0061t\u0065")
	_acd.StateModel = _gcg.Get("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c")
	return &_acd, nil
}

// AddExtGState adds a graphics state to the XObject resources.
func (_baedc *PdfPage) AddExtGState(name _ebb.PdfObjectName, egs *_ebb.PdfObjectDictionary) error {
	if _baedc.Resources == nil {
		_baedc.Resources = NewPdfPageResources()
	}
	if _baedc.Resources.ExtGState == nil {
		_baedc.Resources.ExtGState = _ebb.MakeDict()
	}
	_aaeb, _beaef := _ebb.TraceToDirectObject(_baedc.Resources.ExtGState).(*_ebb.PdfObjectDictionary)
	if !_beaef {
		_eg.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _ebb.TraceToDirectObject(_baedc.Resources.ExtGState))
		return _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_aaeb.Set(name, egs)
	return nil
}
func (_fggdb *PdfWriter) setDocInfo(_bddaa _ebb.PdfObject) {
	if _fggdb.hasObject(_fggdb._eadfd) {
		delete(_fggdb._ffffd, _fggdb._eadfd)
		delete(_fggdb._dcfg, _fggdb._eadfd)
		for _aafgb, _fcfcb := range _fggdb._ebdgg {
			if _fcfcb == _fggdb._eadfd {
				copy(_fggdb._ebdgg[_aafgb:], _fggdb._ebdgg[_aafgb+1:])
				_fggdb._ebdgg[len(_fggdb._ebdgg)-1] = nil
				_fggdb._ebdgg = _fggdb._ebdgg[:len(_fggdb._ebdgg)-1]
				break
			}
		}
	}
	_aafb := _ebb.PdfIndirectObject{}
	_aafb.PdfObject = _bddaa
	_fggdb._eadfd = &_aafb
	_fggdb.addObject(&_aafb)
}

// ImageToRGB returns the passed in image. Method exists in order to satisfy
// the PdfColorspace interface.
func (_ebcd *PdfColorspaceDeviceRGB) ImageToRGB(img Image) (Image, error) { return img, nil }

// PdfAnnotationProjection represents Projection annotations.
type PdfAnnotationProjection struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
}

// GetObjectNums returns the object numbers of the PDF objects in the file
// Numbered objects are either indirect objects or stream objects.
// e.g. objNums := pdfReader.GetObjectNums()
// The underlying objects can then be accessed with
// pdfReader.GetIndirectObjectByNumber(objNums[0]) for the first available object.
func (_adcbb *PdfReader) GetObjectNums() []int { return _adcbb._cafdf.GetObjectNums() }

// GetColorspaceByName returns the colorspace with the specified name from the page resources.
func (_cdfgd *PdfPageResources) GetColorspaceByName(keyName _ebb.PdfObjectName) (PdfColorspace, bool) {
	_aeccb, _deedg := _cdfgd.GetColorspaces()
	if _deedg != nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _deedg)
		return nil, false
	}
	if _aeccb == nil {
		return nil, false
	}
	_efgfa, _dgedd := _aeccb.Colorspaces[string(keyName)]
	if !_dgedd {
		return nil, false
	}
	return _efgfa, true
}

// PdfModel is a higher level PDF construct which can be collapsed into a PdfObject.
// Each PdfModel has an underlying PdfObject and vice versa (one-to-one).
// Under normal circumstances there should only be one copy of each.
// Copies can be made, but care must be taken to do it properly.
type PdfModel interface {
	ToPdfObject() _ebb.PdfObject
	GetContainingPdfObject() _ebb.PdfObject
}

func (_aega *PdfWriter) seekByName(_fabbc _ebb.PdfObject, _defad []string, _cddcd string) ([]_ebb.PdfObject, error) {
	_eg.Log.Trace("\u0053\u0065\u0065\u006b\u0020\u0062\u0079\u0020\u006e\u0061\u006d\u0065.\u002e\u0020\u0025\u0054", _fabbc)
	var _cgfcdd []_ebb.PdfObject
	if _ddde, _fdcggc := _fabbc.(*_ebb.PdfIndirectObject); _fdcggc {
		return _aega.seekByName(_ddde.PdfObject, _defad, _cddcd)
	}
	if _abfcc, _dgcbe := _fabbc.(*_ebb.PdfObjectStream); _dgcbe {
		return _aega.seekByName(_abfcc.PdfObjectDictionary, _defad, _cddcd)
	}
	if _gfgef, _fdgbf := _fabbc.(*_ebb.PdfObjectDictionary); _fdgbf {
		_eg.Log.Trace("\u0044\u0069\u0063\u0074")
		for _, _eegcef := range _gfgef.Keys() {
			_bccff := _gfgef.Get(_eegcef)
			if string(_eegcef) == _cddcd {
				_cgfcdd = append(_cgfcdd, _bccff)
			}
			for _, _adff := range _defad {
				if string(_eegcef) == _adff {
					_eg.Log.Trace("\u0046\u006f\u006c\u006c\u006f\u0077\u0020\u006b\u0065\u0079\u0020\u0025\u0073", _adff)
					_bbedg, _cgbd := _aega.seekByName(_bccff, _defad, _cddcd)
					if _cgbd != nil {
						return _cgfcdd, _cgbd
					}
					_cgfcdd = append(_cgfcdd, _bbedg...)
					break
				}
			}
		}
		return _cgfcdd, nil
	}
	return _cgfcdd, nil
}

// ToPdfObject converts rectangle to a PDF object.
func (_bedcc *PdfRectangle) ToPdfObject() _ebb.PdfObject {
	return _ebb.MakeArray(_ebb.MakeFloat(_bedcc.Llx), _ebb.MakeFloat(_bedcc.Lly), _ebb.MakeFloat(_bedcc.Urx), _ebb.MakeFloat(_bedcc.Ury))
}

// GetCatalogMetadata gets the catalog defined XMP Metadata.
func (_gbcc *PdfReader) GetCatalogMetadata() (_ebb.PdfObject, bool) {
	if _gbcc._fdgda == nil {
		return nil, false
	}
	_aafce := _gbcc._fdgda.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _aafce, _aafce != nil
}

// ToPdfObject returns the PDF representation of the function.
func (_ggafb *PdfFunctionType3) ToPdfObject() _ebb.PdfObject {
	_bbadc := _ebb.MakeDict()
	_bbadc.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _ebb.MakeInteger(3))
	_fddba := &_ebb.PdfObjectArray{}
	for _, _gcgfa := range _ggafb.Domain {
		_fddba.Append(_ebb.MakeFloat(_gcgfa))
	}
	_bbadc.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _fddba)
	if _ggafb.Range != nil {
		_ccgd := &_ebb.PdfObjectArray{}
		for _, _fbab := range _ggafb.Range {
			_ccgd.Append(_ebb.MakeFloat(_fbab))
		}
		_bbadc.Set("\u0052\u0061\u006eg\u0065", _ccgd)
	}
	if _ggafb.Functions != nil {
		_ggecc := &_ebb.PdfObjectArray{}
		for _, _gffc := range _ggafb.Functions {
			_ggecc.Append(_gffc.ToPdfObject())
		}
		_bbadc.Set("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es", _ggecc)
	}
	if _ggafb.Bounds != nil {
		_ebcc := &_ebb.PdfObjectArray{}
		for _, _abeaf := range _ggafb.Bounds {
			_ebcc.Append(_ebb.MakeFloat(_abeaf))
		}
		_bbadc.Set("\u0042\u006f\u0075\u006e\u0064\u0073", _ebcc)
	}
	if _ggafb.Encode != nil {
		_ddeda := &_ebb.PdfObjectArray{}
		for _, _gebga := range _ggafb.Encode {
			_ddeda.Append(_ebb.MakeFloat(_gebga))
		}
		_bbadc.Set("\u0045\u006e\u0063\u006f\u0064\u0065", _ddeda)
	}
	if _ggafb._acac != nil {
		_ggafb._acac.PdfObject = _bbadc
		return _ggafb._acac
	}
	return _bbadc
}

var ImageHandling ImageHandler = DefaultImageHandler{}

// GetPatternByName gets the pattern specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_cdgb *PdfPageResources) GetPatternByName(keyName _ebb.PdfObjectName) (*PdfPattern, bool) {
	if _cdgb.Pattern == nil {
		return nil, false
	}
	_cbgcg, _bacgg := _ebb.TraceToDirectObject(_cdgb.Pattern).(*_ebb.PdfObjectDictionary)
	if !_bacgg {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074t\u0065\u0072\u006e\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cdgb.Pattern)
		return nil, false
	}
	if _cdde := _cbgcg.Get(keyName); _cdde != nil {
		_ddcd, _gbce := _ccef(_cdde)
		if _gbce != nil {
			_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0070\u0061\u0074\u0074\u0065\u0072n\u003a\u0020\u0025\u0076", _gbce)
			return nil, false
		}
		return _ddcd, true
	}
	return nil, false
}
func (_dfbca *PdfWriter) hasObject(_adcad _ebb.PdfObject) bool {
	_, _ecdgeg := _dfbca._ffffd[_adcad]
	return _ecdgeg
}
func (_cccfb *PdfReader) loadForms() (*PdfAcroForm, error) {
	if _cccfb._cafdf.GetCrypter() != nil && !_cccfb._cafdf.IsAuthenticated() {
		return nil, _bg.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_fegfd := _cccfb._fdgda
	_ggfdg := _fegfd.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")
	if _ggfdg == nil {
		return nil, nil
	}
	_dccdd, _ := _ebb.GetIndirect(_ggfdg)
	_ggfdg = _ebb.TraceToDirectObject(_ggfdg)
	if _ebb.IsNullObject(_ggfdg) {
		_eg.Log.Trace("\u0041\u0063\u0072of\u006f\u0072\u006d\u0020\u0069\u0073\u0020\u0061\u0020n\u0075l\u006c \u006fb\u006a\u0065\u0063\u0074\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029\u000a")
		return nil, nil
	}
	_gfcd, _bceb := _ebb.GetDict(_ggfdg)
	if !_bceb {
		_eg.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0041\u0063\u0072\u006fF\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _ggfdg)
		_eg.Log.Debug("\u0044\u006f\u0065\u0073 n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0073")
		return nil, _bg.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0061\u0063\u0072\u006ff\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _ggfdg)
	}
	_eg.Log.Trace("\u0048\u0061\u0073\u0020\u0041\u0063\u0072\u006f\u0020f\u006f\u0072\u006d\u0073")
	_eg.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0041\u0063r\u006ff\u006f\u0072\u006d\u0073\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065")
	if !_cccfb._ceefa {
		_beedb := _cccfb.traverseObjectData(_gfcd)
		if _beedb != nil {
			_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0074\u0072a\u0076\u0065\u0072\u0073\u0065\u0020\u0041\u0063\u0072\u006fFo\u0072\u006d\u0073 \u0028%\u0073\u0029", _beedb)
			return nil, _beedb
		}
	}
	_cbfbfb, _cdagg := _cccfb.newPdfAcroFormFromDict(_dccdd, _gfcd)
	if _cdagg != nil {
		return nil, _cdagg
	}
	return _cbfbfb, nil
}

// SetImage updates XObject Image with new image data.
func (_aagdg *XObjectImage) SetImage(img *Image, cs PdfColorspace) error {
	_aagdg.Filter.UpdateParams(img.GetParamsDict())
	_dbeg, _gbdce := _aagdg.Filter.EncodeBytes(img.Data)
	if _gbdce != nil {
		return _gbdce
	}
	_aagdg.Stream = _dbeg
	_egddcd := img.Width
	_aagdg.Width = &_egddcd
	_bcbafa := img.Height
	_aagdg.Height = &_bcbafa
	_dbdgff := img.BitsPerComponent
	_aagdg.BitsPerComponent = &_dbdgff
	if cs == nil {
		if img.ColorComponents == 1 {
			_aagdg.ColorSpace = NewPdfColorspaceDeviceGray()
		} else if img.ColorComponents == 3 {
			_aagdg.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			_aagdg.ColorSpace = NewPdfColorspaceDeviceCMYK()
		} else {
			return _gf.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_aagdg.ColorSpace = cs
	}
	return nil
}
func (_defbf *PdfColorspaceSpecialIndexed) String() string {
	return "\u0049n\u0064\u0065\u0078\u0065\u0064"
}

var (
	ErrRequiredAttributeMissing = _gf.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074t\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
	ErrInvalidAttribute         = _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	ErrTypeCheck                = _gf.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	_fddb                       = _gf.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrEncrypted                = _gf.New("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	ErrNoFont                   = _gf.New("\u0066\u006fn\u0074\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	ErrFontNotSupported         = _bfc.Errorf("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u006fn\u0074\u0020\u0028\u0025\u0077\u0029", _ebb.ErrNotSupported)
	ErrType1CFontNotSupported   = _bfc.Errorf("\u0054y\u0070\u00651\u0043\u0020\u0066o\u006e\u0074\u0073\u0020\u0061\u0072\u0065 \u006e\u006f\u0074\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0028\u0025\u0077\u0029", _ebb.ErrNotSupported)
	ErrType3FontNotSupported    = _bfc.Errorf("\u0054y\u0070\u00653\u0020\u0066\u006f\u006et\u0073\u0020\u0061r\u0065\u0020\u006e\u006f\u0074\u0020\u0063\u0075\u0072re\u006e\u0074\u006cy\u0020\u0073u\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0028%\u0077\u0029", _ebb.ErrNotSupported)
	ErrTTCmapNotSupported       = _bfc.Errorf("\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065d\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070e\u0020\u0063\u006d\u0061\u0070\u0020\u0066\u006f\u0072\u006d\u0061\u0074 \u0028\u0025\u0077\u0029", _ebb.ErrNotSupported)
	ErrSignNotEnoughSpace       = _bfc.Errorf("\u0069\u006e\u0073\u0075\u0066\u0066\u0069c\u0069\u0065\u006et\u0020\u0073\u0070a\u0063\u0065 \u0061\u006c\u006c\u006f\u0063\u0061t\u0065d \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
	ErrSignNoCertificates       = _bfc.Errorf("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065tri\u0065ve\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0063\u0068\u0061i\u006e")
)

// PdfAnnotationSound represents Sound annotations.
// (Section 12.5.6.16).
type PdfAnnotationSound struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Sound _ebb.PdfObject
	Name  _ebb.PdfObject
}

func _bdgae(_cbac *_ebb.PdfObjectDictionary) {
	_fadfb, _ccfdc := _ebb.GetArray(_cbac.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
	_caea, _fecf := _ebb.GetIntVal(_cbac.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
	_gebbg, _gacbe := _ebb.GetIntVal(_cbac.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
	if _ccfdc && _fecf && _gacbe {
		_daee := _fadfb.Len()
		if _daee != _gebbg-_caea+1 {
			_eg.Log.Debug("\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076\u002c\u0020\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u003a\u0020\u0025\u0076", _daee, _gebbg)
			_dgdd := _ebb.PdfObjectInteger(_caea + _daee - 1)
			_cbac.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", &_dgdd)
		}
	}
}

// ToPdfObject implements interface PdfModel.
func (_afg *PdfAnnotationRedact) ToPdfObject() _ebb.PdfObject {
	_afg.PdfAnnotation.ToPdfObject()
	_cfbd := _afg._bdcd
	_fcfb := _cfbd.PdfObject.(*_ebb.PdfObjectDictionary)
	_afg.PdfAnnotationMarkup.appendToPdfDictionary(_fcfb)
	_fcfb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0052\u0065\u0064\u0061\u0063\u0074"))
	_fcfb.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _afg.QuadPoints)
	_fcfb.SetIfNotNil("\u0049\u0043", _afg.IC)
	_fcfb.SetIfNotNil("\u0052\u004f", _afg.RO)
	_fcfb.SetIfNotNil("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074", _afg.OverlayText)
	_fcfb.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _afg.Repeat)
	_fcfb.SetIfNotNil("\u0044\u0041", _afg.DA)
	_fcfb.SetIfNotNil("\u0051", _afg.Q)
	return _cfbd
}

var _eacf = false

// NewPdfAction returns an initialized generic PDF action model.
func NewPdfAction() *PdfAction {
	_af := &PdfAction{}
	_af._abe = _ebb.MakeIndirectObject(_ebb.MakeDict())
	return _af
}

// PdfAnnotationHighlight represents Highlight annotations.
// (Section 12.5.6.10).
type PdfAnnotationHighlight struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _ebb.PdfObject
}

// NewStandardPdfOutputIntent creates a new standard PdfOutputIntent.
func NewStandardPdfOutputIntent(outputCondition, outputConditionIdentifier, registryName string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, RegistryName: registryName, DestOutputProfile: destOutputProfile, ColorComponents: colorComponents, _faeb: _ebb.MakeDict()}
}
func (_debd *pdfFontType0) subsetRegistered() error {
	_bcccg, _fecbc := _debd.DescendantFont._ebcad.(*pdfCIDFontType2)
	if !_fecbc {
		_eg.Log.Debug("\u0046\u006fnt\u0020\u006e\u006ft\u0020\u0073\u0075\u0070por\u0074ed\u0020\u0066\u006f\u0072\u0020\u0073\u0075bs\u0065\u0074\u0074\u0069\u006e\u0067\u0020%\u0054", _debd.DescendantFont)
		return nil
	}
	if _bcccg == nil {
		return nil
	}
	if _bcccg._fbbd == nil {
		_eg.Log.Debug("\u004d\u0069\u0073si\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return nil
	}
	if _debd._bfdgc == nil {
		_eg.Log.Debug("\u004e\u006f\u0020e\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0073\u0075\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0067\u006e\u006f\u0072\u0065\u0064")
		return nil
	}
	_gddfb, _fecbc := _ebb.GetStream(_bcccg._fbbd.FontFile2)
	if !_fecbc {
		_eg.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _gf.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_bcbaf, _gcffb := _ebb.DecodeStream(_gddfb)
	if _gcffb != nil {
		_eg.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _gcffb)
		return _gcffb
	}
	_adge, _gcffb := _gbg.Parse(_ca.NewReader(_bcbaf))
	if _gcffb != nil {
		_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_gddfb.Stream))
		return _gcffb
	}
	var _cccbd []rune
	var _begg *_gbg.Font
	switch _fcgbb := _debd._bfdgc.(type) {
	case *_da.TrueTypeFontEncoder:
		_cccbd = _fcgbb.RegisteredRunes()
		_begg, _gcffb = _adge.SubsetKeepRunes(_cccbd)
		if _gcffb != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gcffb)
			return _gcffb
		}
		_fcgbb.SubsetRegistered()
	case *_da.IdentityEncoder:
		_cccbd = _fcgbb.RegisteredRunes()
		_baff := make([]_gbg.GlyphIndex, len(_cccbd))
		for _dfffcf, _cafeb := range _cccbd {
			_baff[_dfffcf] = _gbg.GlyphIndex(_cafeb)
		}
		_begg, _gcffb = _adge.SubsetKeepIndices(_baff)
		if _gcffb != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gcffb)
			return _gcffb
		}
	case _da.SimpleEncoder:
		_dcab := _fcgbb.Charcodes()
		for _, _caeffc := range _dcab {
			_cdadd, _bgbe := _fcgbb.CharcodeToRune(_caeffc)
			if !_bgbe {
				_eg.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0020\u0074\u006f \u0072\u0075\u006e\u0065\u003a\u0020\u0025\u0064", _caeffc)
				continue
			}
			_cccbd = append(_cccbd, _cdadd)
		}
	default:
		return _bg.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069\u006eg\u003a\u0020\u0025\u0054", _debd._bfdgc)
	}
	var _efedd _ca.Buffer
	_gcffb = _begg.Write(&_efedd)
	if _gcffb != nil {
		_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gcffb)
		return _gcffb
	}
	if _debd._dcdd != nil {
		_gegd := make(map[_ebe.CharCode]rune, len(_cccbd))
		for _, _edefd := range _cccbd {
			_dgeab, _gbae := _debd._bfdgc.RuneToCharcode(_edefd)
			if !_gbae {
				continue
			}
			_gegd[_ebe.CharCode(_dgeab)] = _edefd
		}
		_debd._dcdd = _ebe.NewToUnicodeCMap(_gegd)
	}
	_gddfb, _gcffb = _ebb.MakeStream(_efedd.Bytes(), _ebb.NewFlateEncoder())
	if _gcffb != nil {
		_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gcffb)
		return _gcffb
	}
	_gddfb.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _ebb.MakeInteger(int64(_efedd.Len())))
	if _aage, _ffce := _ebb.GetStream(_bcccg._fbbd.FontFile2); _ffce {
		*_aage = *_gddfb
	} else {
		_bcccg._fbbd.FontFile2 = _gddfb
	}
	_adgg := _bdec()
	if len(_debd._fdacg) > 0 {
		_debd._fdacg = _gfed(_debd._fdacg, _adgg)
	}
	if len(_bcccg._fdacg) > 0 {
		_bcccg._fdacg = _gfed(_bcccg._fdacg, _adgg)
	}
	if len(_debd._efge) > 0 {
		_debd._efge = _gfed(_debd._efge, _adgg)
	}
	if _bcccg._fbbd != nil {
		_dggba, _aaaed := _ebb.GetName(_bcccg._fbbd.FontName)
		if _aaaed && len(_dggba.String()) > 0 {
			_edced := _gfed(_dggba.String(), _adgg)
			_bcccg._fbbd.FontName = _ebb.MakeName(_edced)
		}
	}
	return nil
}

// GetContainingPdfObject returns the page as a dictionary within an PdfIndirectObject.
func (_ecaf *PdfPage) GetContainingPdfObject() _ebb.PdfObject { return _ecaf._defbb }

// PdfActionSound represents a sound action.
type PdfActionSound struct {
	*PdfAction
	Sound       _ebb.PdfObject
	Volume      _ebb.PdfObject
	Synchronous _ebb.PdfObject
	Repeat      _ebb.PdfObject
	Mix         _ebb.PdfObject
}

func (_gfge *PdfPage) setContainer(_gcbef *_ebb.PdfIndirectObject) {
	_gcbef.PdfObject = _gfge._cdbfde
	_gfge._defbb = _gcbef
}

// NewPdfAnnotationTrapNet returns a new trapnet annotation.
func NewPdfAnnotationTrapNet() *PdfAnnotationTrapNet {
	_gae := NewPdfAnnotation()
	_aaf := &PdfAnnotationTrapNet{}
	_aaf.PdfAnnotation = _gae
	_gae.SetContext(_aaf)
	return _aaf
}

// ContentStreamWrapper wraps the Page's contentstream into q ... Q blocks.
type ContentStreamWrapper interface{ WrapContentStream(_dgdba *PdfPage) error }

// CharcodesToUnicodeWithStats is identical to CharcodesToUnicode except it returns more statistical
// information about hits and misses from the reverse mapping process.
// NOTE: The number of runes returned may be greater than the number of charcodes.
// TODO(peterwilliams97): Deprecate in v4 and use only CharcodesToStrings()
func (_dffef *PdfFont) CharcodesToUnicodeWithStats(charcodes []_da.CharCode) (_bdef []rune, _cgggg, _gfda int) {
	_cgdb, _cgggg, _gfda := _dffef.CharcodesToStrings(charcodes)
	return []rune(_ee.Join(_cgdb, "")), _cgggg, _gfda
}

// NewPdfAnnotationLine returns a new line annotation.
func NewPdfAnnotationLine() *PdfAnnotationLine {
	_ddg := NewPdfAnnotation()
	_fgbf := &PdfAnnotationLine{}
	_fgbf.PdfAnnotation = _ddg
	_fgbf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ddg.SetContext(_fgbf)
	return _fgbf
}

// NewOutline returns a new outline instance.
func NewOutline() *Outline { return &Outline{} }

// Encoder returns the font's text encoder.
func (_bbdc pdfCIDFontType2) Encoder() _da.TextEncoder { return _bbdc._aacb }

// GetContainingPdfObject returns the container of the image object (indirect object).
func (_egba *XObjectImage) GetContainingPdfObject() _ebb.PdfObject { return _egba._fbeec }

// NewPdfActionTrans returns a new "trans" action.
func NewPdfActionTrans() *PdfActionTrans {
	_fg := NewPdfAction()
	_cad := &PdfActionTrans{}
	_cad.PdfAction = _fg
	_fg.SetContext(_cad)
	return _cad
}

// DecodeArray returns the range of color component values in DeviceCMYK colorspace.
func (_bcca *PdfColorspaceDeviceCMYK) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}
func _adaaf(_gfcdd *_ebb.PdfObjectDictionary) (*PdfShadingType1, error) {
	_bbddf := PdfShadingType1{}
	if _cffaf := _gfcdd.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _cffaf != nil {
		_cffaf = _ebb.TraceToDirectObject(_cffaf)
		_cecff, _gfgfd := _cffaf.(*_ebb.PdfObjectArray)
		if !_gfgfd {
			_eg.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cffaf)
			return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_bbddf.Domain = _cecff
	}
	if _gceea := _gfcdd.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _gceea != nil {
		_gceea = _ebb.TraceToDirectObject(_gceea)
		_gecb, _ggfca := _gceea.(*_ebb.PdfObjectArray)
		if !_ggfca {
			_eg.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _gceea)
			return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_bbddf.Matrix = _gecb
	}
	_gfad := _gfcdd.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _gfad == nil {
		_eg.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_bbddf.Function = []PdfFunction{}
	if _cebfa, _decf := _gfad.(*_ebb.PdfObjectArray); _decf {
		for _, _bafgc := range _cebfa.Elements() {
			_bagf, _afbfg := _aagg(_bafgc)
			if _afbfg != nil {
				_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _afbfg)
				return nil, _afbfg
			}
			_bbddf.Function = append(_bbddf.Function, _bagf)
		}
	} else {
		_ecgge, _affeaf := _aagg(_gfad)
		if _affeaf != nil {
			_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _affeaf)
			return nil, _affeaf
		}
		_bbddf.Function = append(_bbddf.Function, _ecgge)
	}
	return &_bbddf, nil
}
func _bfegc(_bbccg *PdfPage) {
	_abcag := _ebb.PdfObjectName("\u0055\u0046\u0031")
	if !_bbccg.Resources.HasFontByName(_abcag) {
		_bbccg.Resources.SetFontByName(_abcag, DefaultFont().ToPdfObject())
	}
	var _bgcfg []string
	_bgcfg = append(_bgcfg, "\u0071")
	_bgcfg = append(_bgcfg, "\u0042\u0054")
	_bgcfg = append(_bgcfg, _bg.Sprintf("\u002f%\u0073\u0020\u0031\u0034\u0020\u0054f", _abcag.String()))
	_bgcfg = append(_bgcfg, "\u0031\u0020\u0030\u0020\u0030\u0020\u0072\u0067")
	_bgcfg = append(_bgcfg, "\u0031\u0030\u0020\u0031\u0030\u0020\u0054\u0064")
	_bgcfg = append(_bgcfg, "\u0045\u0054")
	_bgcfg = append(_bgcfg, "\u0051")
	_bedbe := _ee.Join(_bgcfg, "\u000a")
	_bbccg.AddContentStreamByString(_bedbe)
	_bbccg.ToPdfObject()
}
func _efad(_cdfg _ebb.PdfObject) (*PdfColorspaceLab, error) {
	_aabga := NewPdfColorspaceLab()
	if _acdcc, _afgd := _cdfg.(*_ebb.PdfIndirectObject); _afgd {
		_aabga._bdga = _acdcc
	}
	_cdfg = _ebb.TraceToDirectObject(_cdfg)
	_fcdb, _acff := _cdfg.(*_ebb.PdfObjectArray)
	if !_acff {
		return nil, _bg.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _fcdb.Len() != 2 {
		return nil, _bg.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_cdfg = _ebb.TraceToDirectObject(_fcdb.Get(0))
	_badab, _acff := _cdfg.(*_ebb.PdfObjectName)
	if !_acff {
		return nil, _bg.Errorf("\u006c\u0061\u0062\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if *_badab != "\u004c\u0061\u0062" {
		return nil, _bg.Errorf("n\u006ft\u0020\u0061\u0020\u004c\u0061\u0062\u0020\u0063o\u006c\u006f\u0072\u0073pa\u0063\u0065")
	}
	_cdfg = _ebb.TraceToDirectObject(_fcdb.Get(1))
	_gacfb, _acff := _cdfg.(*_ebb.PdfObjectDictionary)
	if !_acff {
		return nil, _bg.Errorf("c\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006ev\u0061l\u0069\u0064")
	}
	_cdfg = _gacfb.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_cdfg = _ebb.TraceToDirectObject(_cdfg)
	_gcef, _acff := _cdfg.(*_ebb.PdfObjectArray)
	if !_acff {
		return nil, _bg.Errorf("\u004c\u0061\u0062\u0020In\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069n\u0074")
	}
	if _gcef.Len() != 3 {
		return nil, _bg.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
	}
	_edbg, _agbg := _gcef.GetAsFloat64Slice()
	if _agbg != nil {
		return nil, _agbg
	}
	_aabga.WhitePoint = _edbg
	_cdfg = _gacfb.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _cdfg != nil {
		_cdfg = _ebb.TraceToDirectObject(_cdfg)
		_dfcd, _ggcd := _cdfg.(*_ebb.PdfObjectArray)
		if !_ggcd {
			return nil, _bg.Errorf("\u004c\u0061\u0062: \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
		}
		if _dfcd.Len() != 3 {
			return nil, _bg.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
		}
		_bfga, _gdag := _dfcd.GetAsFloat64Slice()
		if _gdag != nil {
			return nil, _gdag
		}
		_aabga.BlackPoint = _bfga
	}
	_cdfg = _gacfb.Get("\u0052\u0061\u006eg\u0065")
	if _cdfg != nil {
		_cdfg = _ebb.TraceToDirectObject(_cdfg)
		_agaa, _bgeb := _cdfg.(*_ebb.PdfObjectArray)
		if !_bgeb {
			_eg.Log.Error("\u0052\u0061n\u0067\u0065\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
			return nil, _bg.Errorf("\u004ca\u0062:\u0020\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if _agaa.Len() != 4 {
			_eg.Log.Error("\u0052\u0061\u006e\u0067\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020e\u0072\u0072\u006f\u0072")
			return nil, _bg.Errorf("\u004c\u0061b\u003a\u0020\u0052a\u006e\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_ggcf, _faadc := _agaa.GetAsFloat64Slice()
		if _faadc != nil {
			return nil, _faadc
		}
		_aabga.Range = _ggcf
	}
	return _aabga, nil
}

// NewXObjectFormFromStream builds the Form XObject from a stream object.
// TODO: Should this be exposed? Consider different access points.
func NewXObjectFormFromStream(stream *_ebb.PdfObjectStream) (*XObjectForm, error) {
	_ceggb := &XObjectForm{}
	_ceggb._gebcd = stream
	_cdbcb := *(stream.PdfObjectDictionary)
	_bfcgc, _acggg := _ebb.NewEncoderFromStream(stream)
	if _acggg != nil {
		return nil, _acggg
	}
	_ceggb.Filter = _bfcgc
	if _eeeaad := _cdbcb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _eeeaad != nil {
		_dagcd, _daddd := _eeeaad.(*_ebb.PdfObjectName)
		if !_daddd {
			return nil, _gf.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if *_dagcd != "\u0046\u006f\u0072\u006d" {
			_eg.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
			return nil, _gf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
		}
	}
	if _dddf := _cdbcb.Get("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065"); _dddf != nil {
		_ceggb.FormType = _dddf
	}
	if _fgecg := _cdbcb.Get("\u0042\u0042\u006f\u0078"); _fgecg != nil {
		_ceggb.BBox = _fgecg
	}
	if _ccfcff := _cdbcb.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _ccfcff != nil {
		_ceggb.Matrix = _ccfcff
	}
	if _fdcff := _cdbcb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _fdcff != nil {
		_fdcff = _ebb.TraceToDirectObject(_fdcff)
		_cafee, _gbbfb := _fdcff.(*_ebb.PdfObjectDictionary)
		if !_gbbfb {
			_eg.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0058\u004f\u0062j\u0065c\u0074\u0020\u0046\u006f\u0072\u006d\u0020\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u002c\u0020\u0070\u006f\u0069\u006e\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _ebb.ErrTypeError
		}
		_babec, _fbdff := NewPdfPageResourcesFromDict(_cafee)
		if _fbdff != nil {
			_eg.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0020\u0072\u0065\u0073\u006f\u0075rc\u0065\u0073")
			return nil, _fbdff
		}
		_ceggb.Resources = _babec
		_eg.Log.Trace("\u0046\u006f\u0072\u006d r\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u003a\u0020\u0025\u0023\u0076", _ceggb.Resources)
	}
	_ceggb.Group = _cdbcb.Get("\u0047\u0072\u006fu\u0070")
	_ceggb.Ref = _cdbcb.Get("\u0052\u0065\u0066")
	_ceggb.MetaData = _cdbcb.Get("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061")
	_ceggb.PieceInfo = _cdbcb.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o")
	_ceggb.LastModified = _cdbcb.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064")
	_ceggb.StructParent = _cdbcb.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_ceggb.StructParents = _cdbcb.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073")
	_ceggb.OPI = _cdbcb.Get("\u004f\u0050\u0049")
	_ceggb.OC = _cdbcb.Get("\u004f\u0043")
	_ceggb.Name = _cdbcb.Get("\u004e\u0061\u006d\u0065")
	_ceggb.Stream = stream.Stream
	return _ceggb, nil
}

// EnableChain adds the specified certificate chain and validation data (OCSP
// and CRL information) for it to the global scope of the document DSS. The
// added data is used for validating any of the signatures present in the
// document. The LTV client attempts to build the certificate chain up to a
// trusted root by downloading any missing certificates.
func (_eadbg *LTV) EnableChain(chain []*_g.Certificate) error { return _eadbg.enable(nil, chain, "") }

// ToPdfObject returns colorspace in a PDF object format [name stream]
func (_cdg *PdfColorspaceICCBased) ToPdfObject() _ebb.PdfObject {
	_eabf := &_ebb.PdfObjectArray{}
	_eabf.Append(_ebb.MakeName("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"))
	var _dacdc *_ebb.PdfObjectStream
	if _cdg._dfff != nil {
		_dacdc = _cdg._dfff
	} else {
		_dacdc = &_ebb.PdfObjectStream{}
	}
	_fbfeb := _ebb.MakeDict()
	_fbfeb.Set("\u004e", _ebb.MakeInteger(int64(_cdg.N)))
	if _cdg.Alternate != nil {
		_fbfeb.Set("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e", _cdg.Alternate.ToPdfObject())
	}
	if _cdg.Metadata != nil {
		_fbfeb.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _cdg.Metadata)
	}
	if _cdg.Range != nil {
		var _fbfa []_ebb.PdfObject
		for _, _efadc := range _cdg.Range {
			_fbfa = append(_fbfa, _ebb.MakeFloat(_efadc))
		}
		_fbfeb.Set("\u0052\u0061\u006eg\u0065", _ebb.MakeArray(_fbfa...))
	}
	_fbfeb.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ebb.MakeInteger(int64(len(_cdg.Data))))
	_dacdc.Stream = _cdg.Data
	_dacdc.PdfObjectDictionary = _fbfeb
	_eabf.Append(_dacdc)
	if _cdg._dagdd != nil {
		_cdg._dagdd.PdfObject = _eabf
		return _cdg._dagdd
	}
	return _eabf
}

// ToPdfObject returns a PDF object representation of the outline.
func (_acccd *Outline) ToPdfObject() _ebb.PdfObject { return _acccd.ToPdfOutline().ToPdfObject() }

// ImageToRGB converts image in CalGray color space to RGB (A, B, C -> X, Y, Z).
func (_cdfb *PdfColorspaceCalGray) ImageToRGB(img Image) (Image, error) {
	_febae := _abg.NewReader(img.getBase())
	_gaegf := _dg.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_caecf := _abg.NewWriter(_gaegf)
	_faag := _cbg.Pow(2, float64(img.BitsPerComponent)) - 1
	_aea := make([]uint32, 3)
	var (
		_dagd                               uint32
		ANorm, X, Y, Z, _faga, _aedb, _gdgd float64
		_dggdf                              error
	)
	for {
		_dagd, _dggdf = _febae.ReadSample()
		if _dggdf == _ab.EOF {
			break
		} else if _dggdf != nil {
			return img, _dggdf
		}
		ANorm = float64(_dagd) / _faag
		X = _cdfb.WhitePoint[0] * _cbg.Pow(ANorm, _cdfb.Gamma)
		Y = _cdfb.WhitePoint[1] * _cbg.Pow(ANorm, _cdfb.Gamma)
		Z = _cdfb.WhitePoint[2] * _cbg.Pow(ANorm, _cdfb.Gamma)
		_faga = 3.240479*X + -1.537150*Y + -0.498535*Z
		_aedb = -0.969256*X + 1.875992*Y + 0.041556*Z
		_gdgd = 0.055648*X + -0.204043*Y + 1.057311*Z
		_faga = _cbg.Min(_cbg.Max(_faga, 0), 1.0)
		_aedb = _cbg.Min(_cbg.Max(_aedb, 0), 1.0)
		_gdgd = _cbg.Min(_cbg.Max(_gdgd, 0), 1.0)
		_aea[0] = uint32(_faga * _faag)
		_aea[1] = uint32(_aedb * _faag)
		_aea[2] = uint32(_gdgd * _faag)
		if _dggdf = _caecf.WriteSamples(_aea); _dggdf != nil {
			return img, _dggdf
		}
	}
	return _afacb(&_gaegf), nil
}

// FieldAppearanceGenerator generates appearance stream for a given field.
type FieldAppearanceGenerator interface {
	ContentStreamWrapper
	GenerateAppearanceDict(_dagg *PdfAcroForm, _accde *PdfField, _cbga *PdfAnnotationWidget) (*_ebb.PdfObjectDictionary, error)
}

// PdfColorspaceCalGray represents CalGray color space.
type PdfColorspaceCalGray struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      float64
	_bebfd     *_ebb.PdfIndirectObject
}

// CharcodesToUnicode converts the character codes `charcodes` to a slice of runes.
// How it works:
//  1) Use the ToUnicode CMap if there is one.
//  2) Use the underlying font's encoding.
func (_bggg *PdfFont) CharcodesToUnicode(charcodes []_da.CharCode) []rune {
	_cbff, _, _ := _bggg.CharcodesToUnicodeWithStats(charcodes)
	return _cbff
}

// SetContentStream sets the pattern cell's content stream.
func (_ecabf *PdfTilingPattern) SetContentStream(content []byte, encoder _ebb.StreamEncoder) error {
	_abfe, _bffa := _ecabf._dcddc.(*_ebb.PdfObjectStream)
	if !_bffa {
		_eg.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ecabf._dcddc)
		return _ebb.ErrTypeError
	}
	if encoder == nil {
		encoder = _ebb.NewRawEncoder()
	}
	_bgegc := _abfe.PdfObjectDictionary
	_gcfbd := encoder.MakeStreamDict()
	_bgegc.Merge(_gcfbd)
	_cddcf, _agggg := encoder.EncodeBytes(content)
	if _agggg != nil {
		return _agggg
	}
	_bgegc.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ebb.MakeInteger(int64(len(_cddcf))))
	_abfe.Stream = _cddcf
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_cfb *PdfActionMovie) ToPdfObject() _ebb.PdfObject {
	_cfb.PdfAction.ToPdfObject()
	_fgc := _cfb._abe
	_gfc := _fgc.PdfObject.(*_ebb.PdfObjectDictionary)
	_gfc.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeMovie)))
	_gfc.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e", _cfb.Annotation)
	_gfc.SetIfNotNil("\u0054", _cfb.T)
	_gfc.SetIfNotNil("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn", _cfb.Operation)
	return _fgc
}

// GetNumPages returns the number of pages in the document.
func (_cggdf *PdfReader) GetNumPages() (int, error) {
	if _cggdf._cafdf.GetCrypter() != nil && !_cggdf._cafdf.IsAuthenticated() {
		return 0, _bg.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	return len(_cggdf._faebb), nil
}

// IsTerminal returns true for terminal fields, false otherwise.
// Terminal fields are fields whose descendants are only widget annotations.
func (_bccac *PdfField) IsTerminal() bool { return len(_bccac.Kids) == 0 }

// ToGray returns a PdfColorDeviceGray color based on the current RGB color.
func (_daeg *PdfColorDeviceRGB) ToGray() *PdfColorDeviceGray {
	_efca := 0.3*_daeg.R() + 0.59*_daeg.G() + 0.11*_daeg.B()
	_efca = _cbg.Min(_cbg.Max(_efca, 0.0), 1.0)
	return NewPdfColorDeviceGray(_efca)
}

// NewReaderOpts generates a default `ReaderOpts` instance.
func NewReaderOpts() *ReaderOpts { return &ReaderOpts{Password: "", LazyLoad: true} }

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_dfba *PdfColorspaceDeviceGray) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_adcfg, _dcca := _ebb.GetNumbersAsFloat(objects)
	if _dcca != nil {
		return nil, _dcca
	}
	return _dfba.ColorFromFloats(_adcfg)
}
func (_agea *PdfWriter) mapObjectStreams(_eagf bool) (map[_ebb.PdfObject]bool, bool) {
	_cfgec := make(map[_ebb.PdfObject]bool)
	for _, _fbcgg := range _agea._ebdgg {
		if _ebfacf, _efece := _fbcgg.(*_ebb.PdfObjectStreams); _efece {
			_eagf = true
			for _, _eeab := range _ebfacf.Elements() {
				_cfgec[_eeab] = true
				if _adgaf, _geeed := _eeab.(*_ebb.PdfIndirectObject); _geeed {
					_cfgec[_adgaf.PdfObject] = true
				}
			}
		}
	}
	return _cfgec, _eagf
}

// NewPdfActionGoTo returns a new "go to" action.
func NewPdfActionGoTo() *PdfActionGoTo {
	_bb := NewPdfAction()
	_acb := &PdfActionGoTo{}
	_acb.PdfAction = _bb
	_bb.SetContext(_acb)
	return _acb
}

// NewPdfFontFromTTF loads a TTF font and returns a PdfFont type that can be
// used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTF.
func NewPdfFontFromTTF(r _ab.ReadSeeker) (*PdfFont, error) {
	const _fdbc = _da.CharCode(32)
	const _dcad = _da.CharCode(255)
	_gcfb, _fcfed := _ef.ReadAll(r)
	if _fcfed != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _fcfed)
		return nil, _fcfed
	}
	_cfbdd, _fcfed := _bad.TtfParse(_ca.NewReader(_gcfb))
	if _fcfed != nil {
		_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0054\u0054F\u0020\u0066\u006fn\u0074:\u0020\u0025\u0076", _fcfed)
		return nil, _fcfed
	}
	_fcfgg := &pdfFontSimple{_cdff: make(map[_da.CharCode]float64), fontCommon: fontCommon{_dfbf: "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065"}}
	_fcfgg._ebcb = _da.NewWinAnsiEncoder()
	_fcfgg._fdacg = _cfbdd.PostScriptName
	_fcfgg.FirstChar = _ebb.MakeInteger(int64(_fdbc))
	_fcfgg.LastChar = _ebb.MakeInteger(int64(_dcad))
	_daaad := 1000.0 / float64(_cfbdd.UnitsPerEm)
	if len(_cfbdd.Widths) <= 0 {
		return nil, _gf.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_aaef := _daaad * float64(_cfbdd.Widths[0])
	_dgfb := make([]float64, 0, _dcad-_fdbc+1)
	for _gcbc := _fdbc; _gcbc <= _dcad; _gcbc++ {
		_dcfed, _gacea := _fcfgg.Encoder().CharcodeToRune(_gcbc)
		if !_gacea {
			_eg.Log.Debug("\u0052u\u006e\u0065\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u0028c\u006f\u0064\u0065\u003a\u0020\u0025\u0064)", _gcbc)
			_dgfb = append(_dgfb, _aaef)
			continue
		}
		_cfbc, _ggcfc := _cfbdd.Chars[_dcfed]
		if !_ggcfc {
			_eg.Log.Debug("R\u0075\u006e\u0065\u0020no\u0074 \u0069\u006e\u0020\u0054\u0054F\u0020\u0043\u0068\u0061\u0072\u0073")
			_dgfb = append(_dgfb, _aaef)
			continue
		}
		_adaed := _daaad * float64(_cfbdd.Widths[_cfbc])
		_dgfb = append(_dgfb, _adaed)
	}
	_fcfgg.Widths = _ebb.MakeIndirectObject(_ebb.MakeArrayFromFloats(_dgfb))
	if len(_dgfb) < int(_dcad-_fdbc+1) {
		_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u006f\u0066\u0020\u0077\u0069\u0064\u0074\u0068s,\u0020\u0025\u0064 \u003c \u0025\u0064", len(_dgfb), 255-32+1)
		return nil, _ebb.ErrRangeError
	}
	for _ffef := _fdbc; _ffef <= _dcad; _ffef++ {
		_fcfgg._cdff[_ffef] = _dgfb[_ffef-_fdbc]
	}
	_fcfgg.Encoding = _ebb.MakeName("\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	_efcdg := &PdfFontDescriptor{}
	_efcdg.FontName = _ebb.MakeName(_cfbdd.PostScriptName)
	_efcdg.Ascent = _ebb.MakeFloat(_daaad * float64(_cfbdd.TypoAscender))
	_efcdg.Descent = _ebb.MakeFloat(_daaad * float64(_cfbdd.TypoDescender))
	_efcdg.CapHeight = _ebb.MakeFloat(_daaad * float64(_cfbdd.CapHeight))
	_efcdg.FontBBox = _ebb.MakeArrayFromFloats([]float64{_daaad * float64(_cfbdd.Xmin), _daaad * float64(_cfbdd.Ymin), _daaad * float64(_cfbdd.Xmax), _daaad * float64(_cfbdd.Ymax)})
	_efcdg.ItalicAngle = _ebb.MakeFloat(_cfbdd.ItalicAngle)
	_efcdg.MissingWidth = _ebb.MakeFloat(_daaad * float64(_cfbdd.Widths[0]))
	_egbg, _fcfed := _ebb.MakeStream(_gcfb, _ebb.NewFlateEncoder())
	if _fcfed != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _fcfed)
		return nil, _fcfed
	}
	_egbg.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _ebb.MakeInteger(int64(len(_gcfb))))
	_efcdg.FontFile2 = _egbg
	if _cfbdd.Bold {
		_efcdg.StemV = _ebb.MakeInteger(120)
	} else {
		_efcdg.StemV = _ebb.MakeInteger(70)
	}
	_bdde := _cbfd
	if _cfbdd.IsFixedPitch {
		_bdde |= _gfega
	}
	if _cfbdd.ItalicAngle != 0 {
		_bdde |= _ccgbe
	}
	_efcdg.Flags = _ebb.MakeInteger(int64(_bdde))
	_fcfgg._fbbd = _efcdg
	_gfde := &PdfFont{_ebcad: _fcfgg}
	return _gfde, nil
}

// NewOutlineItem returns a new outline item instance.
func NewOutlineItem(title string, dest OutlineDest) *OutlineItem {
	return &OutlineItem{Title: title, Dest: dest}
}

// SetContext sets the sub pattern (context).  Either PdfTilingPattern or PdfShadingPattern.
func (_eabe *PdfPattern) SetContext(ctx PdfModel) { _eabe._ffagg = ctx }

// ToPdfObject returns the PDF representation of the colorspace.
func (_beac *PdfColorspaceSpecialSeparation) ToPdfObject() _ebb.PdfObject {
	_fcda := _ebb.MakeArray(_ebb.MakeName("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"))
	_fcda.Append(_beac.ColorantName)
	_fcda.Append(_beac.AlternateSpace.ToPdfObject())
	_fcda.Append(_beac.TintTransform.ToPdfObject())
	if _beac._cded != nil {
		_beac._cded.PdfObject = _fcda
		return _beac._cded
	}
	return _fcda
}

// GetEncryptionMethod returns a descriptive information string about the encryption method used.
func (_ccgdc *PdfReader) GetEncryptionMethod() string {
	_cgcac := _ccgdc._cafdf.GetCrypter()
	return _cgcac.String()
}
func (_cfebb *PdfSignature) extractChainFromPKCS7() ([]*_g.Certificate, error) {
	_ggebd, _cfca := _gb.Parse(_cfebb.Contents.Bytes())
	if _cfca != nil {
		return nil, _cfca
	}
	return _ggebd.Certificates, nil
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
func (_aabcbd *PdfFont) GetCharMetrics(code _da.CharCode) (CharMetrics, bool) {
	var _baaedf _bad.CharMetrics
	switch _egcca := _aabcbd._ebcad.(type) {
	case *pdfFontSimple:
		if _cgaca, _eafc := _egcca.GetCharMetrics(code); _eafc {
			return _cgaca, _eafc
		}
	case *pdfFontType0:
		if _gbfc, _edcb := _egcca.GetCharMetrics(code); _edcb {
			return _gbfc, _edcb
		}
	case *pdfCIDFontType0:
		if _abgda, _dafa := _egcca.GetCharMetrics(code); _dafa {
			return _abgda, _dafa
		}
	case *pdfCIDFontType2:
		if _addbb, _ebdg := _egcca.GetCharMetrics(code); _ebdg {
			return _addbb, _ebdg
		}
	case *pdfFontType3:
		if _agceaf, _aaeg := _egcca.GetCharMetrics(code); _aaeg {
			return _agceaf, _aaeg
		}
	default:
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020G\u0065\u0074\u0043h\u0061\u0072\u004de\u0074\u0072i\u0063\u0073\u0020\u006e\u006f\u0074 \u0069mp\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d\u0025\u0054\u002e", _aabcbd._ebcad)
		return _baaedf, false
	}
	if _cedc, _ccda := _aabcbd.GetFontDescriptor(); _ccda == nil && _cedc != nil {
		return _bad.CharMetrics{Wx: _cedc._gbfgb}, true
	}
	_eg.Log.Debug("\u0047\u0065\u0074\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069c\u0073\u0020\u0066\u006f\u0072 \u0066\u006fn\u0074\u003d\u0025\u0073", _aabcbd)
	return _baaedf, false
}
func (_dgc *PdfReader) newPdfActionRenditionFromDict(_dgb *_ebb.PdfObjectDictionary) (*PdfActionRendition, error) {
	return &PdfActionRendition{R: _dgb.Get("\u0052"), AN: _dgb.Get("\u0041\u004e"), OP: _dgb.Get("\u004f\u0050"), JS: _dgb.Get("\u004a\u0053")}, nil
}

// GetContainingPdfObject implements interface PdfModel.
func (_deefa *PdfSignature) GetContainingPdfObject() _ebb.PdfObject { return _deefa._ffbgc }

// SetContext sets the specific fielddata type, e.g. would be PdfFieldButton for a button field.
func (_bfdfa *PdfField) SetContext(ctx PdfModel) { _bfdfa._cada = ctx }

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_cbca *PdfShading) ToPdfObject() _ebb.PdfObject {
	_gegb := _cbca._fbfae
	_afgba, _baded := _cbca.getShadingDict()
	if _baded != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _cbca.ShadingType != nil {
		_afgba.Set("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _cbca.ShadingType)
	}
	if _cbca.ColorSpace != nil {
		_afgba.Set("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _cbca.ColorSpace.ToPdfObject())
	}
	if _cbca.Background != nil {
		_afgba.Set("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064", _cbca.Background)
	}
	if _cbca.BBox != nil {
		_afgba.Set("\u0042\u0042\u006f\u0078", _cbca.BBox.ToPdfObject())
	}
	if _cbca.AntiAlias != nil {
		_afgba.Set("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s", _cbca.AntiAlias)
	}
	return _gegb
}
func _ddacd(_cefgb _ebb.PdfObject, _bddde bool) (*PdfFont, error) {
	_gbcb, _bdabe, _dbdb := _efec(_cefgb)
	if _gbcb != nil {
		_bdgae(_gbcb)
	}
	if _dbdb != nil {
		if _dbdb == ErrType1CFontNotSupported {
			_eefe, _fbeg := _bcdgc(_gbcb, _bdabe, nil)
			if _fbeg != nil {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _bdabe, _fbeg)
				return nil, _dbdb
			}
			return &PdfFont{_ebcad: _eefe}, _dbdb
		}
		return nil, _dbdb
	}
	_gccg := &PdfFont{}
	switch _bdabe._dfbf {
	case "\u0054\u0079\u0070e\u0030":
		if !_bddde {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u004c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u00650\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _bdabe)
			return nil, _gf.New("\u0063\u0079\u0063\u006cic\u0061\u006c\u0020\u0074\u0079\u0070\u0065\u0030\u0020\u006c\u006f\u0061\u0064\u0069n\u0067")
		}
		_eage, _dbfc := _aggb(_gbcb, _bdabe)
		if _dbfc != nil {
			_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0057\u0068\u0069l\u0065\u0020\u006c\u006f\u0061\u0064\u0069ng\u0020\u0054\u0079\u0070e\u0030\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066on\u0074\u003d%\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bdabe, _dbfc)
			return nil, _dbfc
		}
		_gccg._ebcad = _eage
	case "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		var _ceffc *pdfFontSimple
		_cgaeb, _cgeb := _bad.NewStdFontByName(_bad.StdFontName(_bdabe._fdacg))
		if _cgeb {
			_bbdd := _fcgbc(_cgaeb)
			_gccg._ebcad = &_bbdd
			_bebee := _ebb.TraceToDirectObject(_bbdd.ToPdfObject())
			_faaf, _defag, _bbdg := _efec(_bebee)
			if _bbdg != nil {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _bdabe, _bbdd)
				return nil, _bbdg
			}
			for _, _dfca := range _gbcb.Keys() {
				_faaf.Set(_dfca, _gbcb.Get(_dfca))
			}
			_ceffc, _bbdg = _bcdgc(_faaf, _defag, _bbdd._dacee)
			if _bbdg != nil {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _bdabe, _bbdd)
				return nil, _bbdg
			}
			_ceffc._cdff = _bbdd._cdff
			_ceffc._ddgd = _bbdd._ddgd
			if _ceffc._adbd == nil {
				_ceffc._adbd = _bbdd._adbd
			}
		} else {
			_ceffc, _dbdb = _bcdgc(_gbcb, _bdabe, nil)
			if _dbdb != nil {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _bdabe, _dbdb)
				return nil, _dbdb
			}
		}
		_dbdb = _ceffc.addEncoding()
		if _dbdb != nil {
			return nil, _dbdb
		}
		if _cgeb {
			_ceffc.updateStandard14Font()
		}
		if _cgeb && _ceffc._ebcb == nil && _ceffc._dacee == nil {
			_eg.Log.Error("\u0073\u0069\u006d\u0070\u006c\u0065\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _ceffc)
			_eg.Log.Error("\u0066n\u0074\u003d\u0025\u002b\u0076", _cgaeb)
		}
		if len(_ceffc._cdff) == 0 {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u006f\u0020\u0077\u0069d\u0074h\u0073.\u0020\u0066\u006f\u006e\u0074\u003d\u0025s", _ceffc)
		}
		_gccg._ebcad = _ceffc
	case "\u0054\u0079\u0070e\u0033":
		_bgfb, _dbdba := _febaee(_gbcb, _bdabe)
		if _dbdba != nil {
			_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020W\u0068\u0069\u006c\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0074y\u0070\u0065\u0033\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _dbdba)
			return nil, _dbdba
		}
		_gccg._ebcad = _bgfb
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030":
		_agcbe, _afgg := _fagea(_gbcb, _bdabe)
		if _afgg != nil {
			_eg.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0057\u0068i\u006c\u0065\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069d \u0066\u006f\u006et\u0020\u0074y\u0070\u0065\u0030\u0020\u0066\u006fn\u0074\u003a \u0025\u0076", _afgg)
			return nil, _afgg
		}
		_gccg._ebcad = _agcbe
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_eabc, _bbeag := _fgdad(_gbcb, _bdabe)
		if _bbeag != nil {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0057\u0068\u0069l\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069\u0064\u0020f\u006f\u006e\u0074\u0020\u0074yp\u0065\u0032\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bdabe, _bbeag)
			return nil, _bbeag
		}
		_gccg._ebcad = _eabc
	default:
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020U\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020f\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073", _bdabe)
		return nil, _bg.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0066\u006f\u006e\u0074\u0020\u0074y\u0070\u0065\u003a\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _bdabe)
	}
	return _gccg, nil
}

// PdfShadingType6 is a Coons patch mesh.
type PdfShadingType6 struct {
	*PdfShading
	BitsPerCoordinate *_ebb.PdfObjectInteger
	BitsPerComponent  *_ebb.PdfObjectInteger
	BitsPerFlag       *_ebb.PdfObjectInteger
	Decode            *_ebb.PdfObjectArray
	Function          []PdfFunction
}

func (_feeae *PdfWriter) writeAcroFormFields() error {
	if _feeae._dbea == nil {
		return nil
	}
	_eg.Log.Trace("\u0057r\u0069t\u0069\u006e\u0067\u0020\u0061c\u0072\u006f \u0066\u006f\u0072\u006d\u0073")
	_cbece := _feeae._dbea.ToPdfObject()
	_eg.Log.Trace("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _cbece)
	_feeae._dffegd.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _cbece)
	_bbgeg := _feeae.addObjects(_cbece)
	if _bbgeg != nil {
		return _bbgeg
	}
	return nil
}
func _bfebg(_dgba _ebb.PdfObject) (string, error) {
	_dgba = _ebb.TraceToDirectObject(_dgba)
	switch _gdeab := _dgba.(type) {
	case *_ebb.PdfObjectString:
		return _gdeab.Str(), nil
	case *_ebb.PdfObjectStream:
		_cbegf, _eggde := _ebb.DecodeStream(_gdeab)
		if _eggde != nil {
			return "", _eggde
		}
		return string(_cbegf), nil
	}
	return "", _bg.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _dgba)
}

// AddImageResource adds an image to the XObject resources.
func (_dbgb *PdfPage) AddImageResource(name _ebb.PdfObjectName, ximg *XObjectImage) error {
	var _fdfae *_ebb.PdfObjectDictionary
	if _dbgb.Resources.XObject == nil {
		_fdfae = _ebb.MakeDict()
		_dbgb.Resources.XObject = _fdfae
	} else {
		var _dcdffe bool
		_fdfae, _dcdffe = (_dbgb.Resources.XObject).(*_ebb.PdfObjectDictionary)
		if !_dcdffe {
			return _gf.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0078\u0072\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020\u0074\u0079p\u0065")
		}
	}
	_fdfae.Set(name, ximg.ToPdfObject())
	return nil
}

// PdfFieldChoice represents a choice field which includes scrollable list boxes and combo boxes.
type PdfFieldChoice struct {
	*PdfField
	Opt *_ebb.PdfObjectArray
	TI  *_ebb.PdfObjectInteger
	I   *_ebb.PdfObjectArray
}

func _aafe(_ebafg _ebb.PdfObject) (*fontFile, error) {
	_eg.Log.Trace("\u006e\u0065\u0077\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0046\u0072\u006f\u006dP\u0064f\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u006f\u0062\u006a\u003d\u0025\u0073", _ebafg)
	_cged := &fontFile{}
	_ebafg = _ebb.TraceToDirectObject(_ebafg)
	_bdaaa, _aeabc := _ebafg.(*_ebb.PdfObjectStream)
	if !_aeabc {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020F\u006f\u006et\u0046\u0069\u006c\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0028\u0025\u0054\u0029", _ebafg)
		return nil, _ebb.ErrTypeError
	}
	_edcfg := _bdaaa.PdfObjectDictionary
	_gcded, _fcdfb := _ebb.DecodeStream(_bdaaa)
	if _fcdfb != nil {
		return nil, _fcdfb
	}
	_bfcf, _aeabc := _ebb.GetNameVal(_edcfg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_aeabc {
		_cged._badbf = _bfcf
		if _bfcf == "\u0054\u0079\u0070\u0065\u0031\u0043" {
			_eg.Log.Debug("T\u0079\u0070\u0065\u0031\u0043\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u006e\u006f\u0074 s\u0075\u0070\u0070o\u0072t\u0065\u0064")
			return nil, ErrType1CFontNotSupported
		}
	}
	_acgfc, _ := _ebb.GetIntVal(_edcfg.Get("\u004ce\u006e\u0067\u0074\u0068\u0031"))
	_accf, _ := _ebb.GetIntVal(_edcfg.Get("\u004ce\u006e\u0067\u0074\u0068\u0032"))
	if _acgfc > len(_gcded) {
		_acgfc = len(_gcded)
	}
	if _acgfc+_accf > len(_gcded) {
		_accf = len(_gcded) - _acgfc
	}
	_aggcb := _gcded[:_acgfc]
	var _bebff []byte
	if _accf > 0 {
		_bebff = _gcded[_acgfc : _acgfc+_accf]
	}
	if _acgfc > 0 && _accf > 0 {
		_efgac := _cged.loadFromSegments(_aggcb, _bebff)
		if _efgac != nil {
			return nil, _efgac
		}
	}
	return _cged, nil
}
func (_fcgbf PdfFont) actualFont() pdfFont {
	if _fcgbf._ebcad == nil {
		_eg.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0061\u0063\u0074\u0075\u0061\u006c\u0046\u006f\u006e\u0074\u002e\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c.\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _fcgbf)
	}
	return _fcgbf._ebcad
}

// PdfAnnotationWidget represents Widget annotations.
// Note: Widget annotations are used to display form fields.
// (Section 12.5.6.19).
type PdfAnnotationWidget struct {
	*PdfAnnotation
	H      _ebb.PdfObject
	MK     _ebb.PdfObject
	A      _ebb.PdfObject
	AA     _ebb.PdfObject
	BS     _ebb.PdfObject
	Parent _ebb.PdfObject
	_gce   *PdfField
	_gdga  bool
}

// GetXObjectFormByName returns the XObjectForm with the specified name from the
// page resources, if it exists.
func (_eefb *PdfPageResources) GetXObjectFormByName(keyName _ebb.PdfObjectName) (*XObjectForm, error) {
	_egfae, _fcbbe := _eefb.GetXObjectByName(keyName)
	if _egfae == nil {
		return nil, nil
	}
	if _fcbbe != XObjectTypeForm {
		return nil, _gf.New("\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006f\u0072\u006d")
	}
	_bdbab, _cadaag := NewXObjectFormFromStream(_egfae)
	if _cadaag != nil {
		return nil, _cadaag
	}
	return _bdbab, nil
}
func _egcdd(_ccfa *_ebb.PdfObjectDictionary) (*PdfFieldChoice, error) {
	_defae := &PdfFieldChoice{}
	_defae.Opt, _ = _ebb.GetArray(_ccfa.Get("\u004f\u0070\u0074"))
	_defae.TI, _ = _ebb.GetInt(_ccfa.Get("\u0054\u0049"))
	_defae.I, _ = _ebb.GetArray(_ccfa.Get("\u0049"))
	return _defae, nil
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_cagf *PdfColorspaceCalRGB) ToPdfObject() _ebb.PdfObject {
	_ecabb := &_ebb.PdfObjectArray{}
	_ecabb.Append(_ebb.MakeName("\u0043\u0061\u006c\u0052\u0047\u0042"))
	_cgce := _ebb.MakeDict()
	if _cagf.WhitePoint != nil {
		_ebfd := _ebb.MakeArray(_ebb.MakeFloat(_cagf.WhitePoint[0]), _ebb.MakeFloat(_cagf.WhitePoint[1]), _ebb.MakeFloat(_cagf.WhitePoint[2]))
		_cgce.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _ebfd)
	} else {
		_eg.Log.Error("\u0043\u0061l\u0052\u0047\u0042\u003a \u004d\u0069s\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _cagf.BlackPoint != nil {
		_ggadf := _ebb.MakeArray(_ebb.MakeFloat(_cagf.BlackPoint[0]), _ebb.MakeFloat(_cagf.BlackPoint[1]), _ebb.MakeFloat(_cagf.BlackPoint[2]))
		_cgce.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _ggadf)
	}
	if _cagf.Gamma != nil {
		_fafd := _ebb.MakeArray(_ebb.MakeFloat(_cagf.Gamma[0]), _ebb.MakeFloat(_cagf.Gamma[1]), _ebb.MakeFloat(_cagf.Gamma[2]))
		_cgce.Set("\u0047\u0061\u006dm\u0061", _fafd)
	}
	if _cagf.Matrix != nil {
		_bedgb := _ebb.MakeArray(_ebb.MakeFloat(_cagf.Matrix[0]), _ebb.MakeFloat(_cagf.Matrix[1]), _ebb.MakeFloat(_cagf.Matrix[2]), _ebb.MakeFloat(_cagf.Matrix[3]), _ebb.MakeFloat(_cagf.Matrix[4]), _ebb.MakeFloat(_cagf.Matrix[5]), _ebb.MakeFloat(_cagf.Matrix[6]), _ebb.MakeFloat(_cagf.Matrix[7]), _ebb.MakeFloat(_cagf.Matrix[8]))
		_cgce.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _bedgb)
	}
	_ecabb.Append(_cgce)
	if _cagf._aeac != nil {
		_cagf._aeac.PdfObject = _ecabb
		return _cagf._aeac
	}
	return _ecabb
}

// ToPdfObject returns a stream object.
func (_fafdfe *XObjectImage) ToPdfObject() _ebb.PdfObject {
	_ggca := _fafdfe._fbeec
	_agbed := _ggca.PdfObjectDictionary
	if _fafdfe.Filter != nil {
		_agbed = _fafdfe.Filter.MakeStreamDict()
		_ggca.PdfObjectDictionary = _agbed
	}
	_agbed.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_agbed.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0049\u006d\u0061g\u0065"))
	_agbed.Set("\u0057\u0069\u0064t\u0068", _ebb.MakeInteger(*(_fafdfe.Width)))
	_agbed.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _ebb.MakeInteger(*(_fafdfe.Height)))
	if _fafdfe.BitsPerComponent != nil {
		_agbed.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _ebb.MakeInteger(*(_fafdfe.BitsPerComponent)))
	}
	if _fafdfe.ColorSpace != nil {
		_agbed.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _fafdfe.ColorSpace.ToPdfObject())
	}
	_agbed.SetIfNotNil("\u0049\u006e\u0074\u0065\u006e\u0074", _fafdfe.Intent)
	_agbed.SetIfNotNil("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k", _fafdfe.ImageMask)
	_agbed.SetIfNotNil("\u004d\u0061\u0073\u006b", _fafdfe.Mask)
	_adcge := _agbed.Get("\u0044\u0065\u0063\u006f\u0064\u0065") != nil
	if _fafdfe.Decode == nil && _adcge {
		_agbed.Remove("\u0044\u0065\u0063\u006f\u0064\u0065")
	} else if _fafdfe.Decode != nil {
		_agbed.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _fafdfe.Decode)
	}
	_agbed.SetIfNotNil("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _fafdfe.Interpolate)
	_agbed.SetIfNotNil("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073", _fafdfe.Alternatives)
	_agbed.SetIfNotNil("\u0053\u004d\u0061s\u006b", _fafdfe.SMask)
	_agbed.SetIfNotNil("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061", _fafdfe.SMaskInData)
	_agbed.SetIfNotNil("\u004d\u0061\u0074t\u0065", _fafdfe.Matte)
	_agbed.SetIfNotNil("\u004e\u0061\u006d\u0065", _fafdfe.Name)
	_agbed.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _fafdfe.StructParent)
	_agbed.SetIfNotNil("\u0049\u0044", _fafdfe.ID)
	_agbed.SetIfNotNil("\u004f\u0050\u0049", _fafdfe.OPI)
	_agbed.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _fafdfe.Metadata)
	_agbed.SetIfNotNil("\u004f\u0043", _fafdfe.OC)
	_agbed.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ebb.MakeInteger(int64(len(_fafdfe.Stream))))
	_ggca.Stream = _fafdfe.Stream
	return _ggca
}

// ToPdfObject returns a *PdfIndirectObject containing a *PdfObjectArray representation of the DeviceN colorspace.
// Format: [/DeviceN names alternateSpace tintTransform]
//     or: [/DeviceN names alternateSpace tintTransform attributes]
func (_ffga *PdfColorspaceDeviceN) ToPdfObject() _ebb.PdfObject {
	_adbb := _ebb.MakeArray(_ebb.MakeName("\u0044e\u0076\u0069\u0063\u0065\u004e"))
	_adbb.Append(_ffga.ColorantNames)
	_adbb.Append(_ffga.AlternateSpace.ToPdfObject())
	_adbb.Append(_ffga.TintTransform.ToPdfObject())
	if _ffga.Attributes != nil {
		_adbb.Append(_ffga.Attributes.ToPdfObject())
	}
	if _ffga._gebb != nil {
		_ffga._gebb.PdfObject = _adbb
		return _ffga._gebb
	}
	return _adbb
}

// ToWriter creates a new writer from the current reader, based on the specified options.
// If no options are provided, all reader properties are copied to the writer.
func (_egccc *PdfReader) ToWriter(opts *ReaderToWriterOpts) (*PdfWriter, error) {
	_fccbcc := NewPdfWriter()
	if opts == nil {
		opts = &ReaderToWriterOpts{}
	}
	_eccee, _dccgb := _egccc.GetNumPages()
	if _dccgb != nil {
		_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dccgb)
		return nil, _dccgb
	}
	for _efdaf := 1; _efdaf <= _eccee; _efdaf++ {
		_ffab, _abff := _egccc.GetPage(_efdaf)
		if _abff != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _abff)
			return nil, _abff
		}
		if opts.PageProcessCallback != nil {
			_abff = opts.PageProcessCallback(_efdaf, _ffab)
			if _abff != nil {
				_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _abff)
				return nil, _abff
			}
		} else if opts.PageCallback != nil {
			opts.PageCallback(_efdaf, _ffab)
		}
		_abff = _fccbcc.AddPage(_ffab)
		if _abff != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _abff)
			return nil, _abff
		}
	}
	_fccbcc._efcge = _egccc.PdfVersion()
	if !opts.SkipInfo {
		_agafg, _bebcg := _egccc.GetPdfInfo()
		if _bebcg != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bebcg)
		} else {
			_fccbcc._eadfd.PdfObject = _agafg.ToPdfObject()
		}
	}
	if !opts.SkipMetadata {
		if _bbec := _egccc._fdgda.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _bbec != nil {
			if _bebd := _fccbcc.SetCatalogMetadata(_bbec); _bebd != nil {
				return nil, _bebd
			}
		}
	}
	if !opts.SkipAcroForm {
		_ceagd := _fccbcc.SetForms(_egccc.AcroForm)
		if _ceagd != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ceagd)
			return nil, _ceagd
		}
	}
	if !opts.SkipOutlines {
		_fccbcc.AddOutlineTree(_egccc.GetOutlineTree())
	}
	if !opts.SkipOCProperties {
		_gafgd, _fdaac := _egccc.GetOCProperties()
		if _fdaac != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fdaac)
		} else {
			_fdaac = _fccbcc.SetOCProperties(_gafgd)
			if _fdaac != nil {
				_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fdaac)
			}
		}
	}
	if !opts.SkipPageLabels {
		_fgdc, _cdaeeg := _egccc.GetPageLabels()
		if _cdaeeg != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cdaeeg)
		} else {
			_cdaeeg = _fccbcc.SetPageLabels(_fgdc)
			if _cdaeeg != nil {
				_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cdaeeg)
			}
		}
	}
	if !opts.SkipNamedDests {
		_dafed, _eadba := _egccc.GetNamedDestinations()
		if _eadba != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _eadba)
		} else {
			_eadba = _fccbcc.SetNamedDestinations(_dafed)
			if _eadba != nil {
				_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _eadba)
			}
		}
	}
	if !opts.SkipNameDictionary {
		_fbdaa, _cccc := _egccc.GetNameDictionary()
		if _cccc != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cccc)
		} else {
			_cccc = _fccbcc.SetNameDictionary(_fbdaa)
			if _cccc != nil {
				_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cccc)
			}
		}
	}
	if !opts.SkipRotation && _egccc.Rotate != nil {
		if _beeea := _fccbcc.SetRotation(*_egccc.Rotate); _beeea != nil {
			_eg.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _beeea)
		}
	}
	return &_fccbcc, nil
}

// ButtonType represents the subtype of a button field, can be one of:
// - Checkbox (ButtonTypeCheckbox)
// - PushButton (ButtonTypePushButton)
// - RadioButton (ButtonTypeRadioButton)
type ButtonType int

var _eefgb = map[string]struct{}{"\u0054\u0069\u0074l\u0065": {}, "\u0041\u0075\u0074\u0068\u006f\u0072": {}, "\u0053u\u0062\u006a\u0065\u0063\u0074": {}, "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073": {}, "\u0043r\u0065\u0061\u0074\u006f\u0072": {}, "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072": {}, "\u0054r\u0061\u0070\u0070\u0065\u0064": {}, "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065": {}, "\u004do\u0064\u0044\u0061\u0074\u0065": {}}
var (
	_daddc _bf.Mutex
	_ecgdd = ""
	_gecg  _f.Time
	_afedf = ""
	_fceef = ""
	_ccdff _f.Time
	_aaaae = ""
	_feaca = ""
	_fead  = ""
)

// ColorFromPdfObjects returns a new PdfColor based on input color components. The input PdfObjects should
// be numeric.
func (_bgeaa *PdfColorspaceDeviceN) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != _bgeaa.GetNumComponents() {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_eaff, _gca := _ebb.GetNumbersAsFloat(objects)
	if _gca != nil {
		return nil, _gca
	}
	return _bgeaa.ColorFromFloats(_eaff)
}

// IsTiling specifies if the pattern is a tiling pattern.
func (_cgdcg *PdfPattern) IsTiling() bool { return _cgdcg.PatternType == 1 }

// PdfBorderStyle represents a border style dictionary (12.5.4 Border Styles p. 394).
type PdfBorderStyle struct {
	W     *float64
	S     *BorderStyle
	D     *[]int
	_dgaa _ebb.PdfObject
}

// AddPage adds a page to the PDF file. The new page should be an indirect object.
func (_acabe *PdfWriter) AddPage(page *PdfPage) error {
	const _efbcf = "\u006d\u006f\u0064el\u003a\u0050\u0064\u0066\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_bfegc(page)
	_bebdg := page.ToPdfObject()
	_eg.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
	_eg.Log.Trace("\u0041p\u0070\u0065\u006e\u0064i\u006e\u0067\u0020\u0074\u006f \u0070a\u0067e\u0020\u006c\u0069\u0073\u0074\u0020\u0025T", _bebdg)
	_eafbf, _eddcf := _ebb.GetIndirect(_bebdg)
	if !_eddcf {
		return _gf.New("\u0070\u0061\u0067\u0065\u0020\u0073h\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_eg.Log.Trace("\u0025\u0073", _eafbf)
	_eg.Log.Trace("\u0025\u0073", _eafbf.PdfObject)
	_fbefd, _eddcf := _ebb.GetDict(_eafbf.PdfObject)
	if !_eddcf {
		return _gf.New("\u0070\u0061\u0067e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068o\u0075l\u0064 \u0062e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_cgdd, _eddcf := _ebb.GetName(_fbefd.Get("\u0054\u0079\u0070\u0065"))
	if !_eddcf {
		return _bg.Errorf("\u0070\u0061\u0067\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054y\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u006e\u0061m\u0065\u0020\u0028%\u0054\u0029", _fbefd.Get("\u0054\u0079\u0070\u0065"))
	}
	if _cgdd.String() != "\u0050\u0061\u0067\u0065" {
		return _gf.New("\u0066\u0069e\u006c\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u0050\u0061\u0067\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069re\u0064\u0029")
	}
	_degfe := []_ebb.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	_gebca, _afffg := _ebb.GetIndirect(_fbefd.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	_eg.Log.Trace("P\u0061g\u0065\u0020\u0050\u0061\u0072\u0065\u006e\u0074:\u0020\u0025\u0054\u0020(%\u0076\u0029", _fbefd.Get("\u0050\u0061\u0072\u0065\u006e\u0074"), _afffg)
	for _afffg {
		_eg.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _gebca)
		_fcdfe, _deea := _ebb.GetDict(_gebca.PdfObject)
		if !_deea {
			return _gf.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		for _, _accdb := range _degfe {
			_eg.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _accdb)
			if _fbefd.Get(_accdb) != nil {
				_eg.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
				continue
			}
			if _acccf := _fcdfe.Get(_accdb); _acccf != nil {
				_eg.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _accdb)
				_fbefd.Set(_accdb, _acccf)
			}
		}
		_gebca, _afffg = _ebb.GetIndirect(_fcdfe.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
		_eg.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _fcdfe.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	}
	_eg.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c \u0064\u006f\u006e\u0065")
	_fbefd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _acabe._dggbf)
	_eafbf.PdfObject = _fbefd
	_dgbbd, _eddcf := _ebb.GetDict(_acabe._dggbf.PdfObject)
	if !_eddcf {
		return _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_cgdgd, _eddcf := _ebb.GetArray(_dgbbd.Get("\u004b\u0069\u0064\u0073"))
	if !_eddcf {
		return _gf.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_cgdgd.Append(_eafbf)
	_acabe._afbdd[_fbefd] = struct{}{}
	_cgaee, _eddcf := _ebb.GetInt(_dgbbd.Get("\u0043\u006f\u0075n\u0074"))
	if !_eddcf {
		return _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	*_cgaee = *_cgaee + 1
	_acabe.addObject(_eafbf)
	_bddab := _acabe.addObjects(_fbefd)
	if _bddab != nil {
		return _bddab
	}
	return nil
}

// ToPdfObject returns the PdfFontDescriptor as a PDF dictionary inside an indirect object.
func (_gadad *PdfFontDescriptor) ToPdfObject() _ebb.PdfObject {
	_ddbeg := _ebb.MakeDict()
	if _gadad._ccfcg == nil {
		_gadad._ccfcg = &_ebb.PdfIndirectObject{}
	}
	_gadad._ccfcg.PdfObject = _ddbeg
	_ddbeg.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072"))
	if _gadad.FontName != nil {
		_ddbeg.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _gadad.FontName)
	}
	if _gadad.FontFamily != nil {
		_ddbeg.Set("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079", _gadad.FontFamily)
	}
	if _gadad.FontStretch != nil {
		_ddbeg.Set("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068", _gadad.FontStretch)
	}
	if _gadad.FontWeight != nil {
		_ddbeg.Set("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074", _gadad.FontWeight)
	}
	if _gadad.Flags != nil {
		_ddbeg.Set("\u0046\u006c\u0061g\u0073", _gadad.Flags)
	}
	if _gadad.FontBBox != nil {
		_ddbeg.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _gadad.FontBBox)
	}
	if _gadad.ItalicAngle != nil {
		_ddbeg.Set("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065", _gadad.ItalicAngle)
	}
	if _gadad.Ascent != nil {
		_ddbeg.Set("\u0041\u0073\u0063\u0065\u006e\u0074", _gadad.Ascent)
	}
	if _gadad.Descent != nil {
		_ddbeg.Set("\u0044e\u0073\u0063\u0065\u006e\u0074", _gadad.Descent)
	}
	if _gadad.Leading != nil {
		_ddbeg.Set("\u004ce\u0061\u0064\u0069\u006e\u0067", _gadad.Leading)
	}
	if _gadad.CapHeight != nil {
		_ddbeg.Set("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t", _gadad.CapHeight)
	}
	if _gadad.XHeight != nil {
		_ddbeg.Set("\u0058H\u0065\u0069\u0067\u0068\u0074", _gadad.XHeight)
	}
	if _gadad.StemV != nil {
		_ddbeg.Set("\u0053\u0074\u0065m\u0056", _gadad.StemV)
	}
	if _gadad.StemH != nil {
		_ddbeg.Set("\u0053\u0074\u0065m\u0048", _gadad.StemH)
	}
	if _gadad.AvgWidth != nil {
		_ddbeg.Set("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068", _gadad.AvgWidth)
	}
	if _gadad.MaxWidth != nil {
		_ddbeg.Set("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068", _gadad.MaxWidth)
	}
	if _gadad.MissingWidth != nil {
		_ddbeg.Set("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068", _gadad.MissingWidth)
	}
	if _gadad.FontFile != nil {
		_ddbeg.Set("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065", _gadad.FontFile)
	}
	if _gadad.FontFile2 != nil {
		_ddbeg.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652", _gadad.FontFile2)
	}
	if _gadad.FontFile3 != nil {
		_ddbeg.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653", _gadad.FontFile3)
	}
	if _gadad.CharSet != nil {
		_ddbeg.Set("\u0043h\u0061\u0072\u0053\u0065\u0074", _gadad.CharSet)
	}
	if _gadad.Style != nil {
		_ddbeg.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _gadad.FontName)
	}
	if _gadad.Lang != nil {
		_ddbeg.Set("\u004c\u0061\u006e\u0067", _gadad.Lang)
	}
	if _gadad.FD != nil {
		_ddbeg.Set("\u0046\u0044", _gadad.FD)
	}
	if _gadad.CIDSet != nil {
		_ddbeg.Set("\u0043\u0049\u0044\u0053\u0065\u0074", _gadad.CIDSet)
	}
	return _gadad._ccfcg
}

// PdfAnnotationText represents Text annotations.
// (Section 12.5.6.4 p. 402).
type PdfAnnotationText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Open       _ebb.PdfObject
	Name       _ebb.PdfObject
	State      _ebb.PdfObject
	StateModel _ebb.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_ccaaf *PdfAnnotationText) ToPdfObject() _ebb.PdfObject {
	_ccaaf.PdfAnnotation.ToPdfObject()
	_afe := _ccaaf._bdcd
	_adba := _afe.PdfObject.(*_ebb.PdfObjectDictionary)
	if _ccaaf.PdfAnnotationMarkup != nil {
		_ccaaf.PdfAnnotationMarkup.appendToPdfDictionary(_adba)
	}
	_adba.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0054\u0065\u0078\u0074"))
	_adba.SetIfNotNil("\u004f\u0070\u0065\u006e", _ccaaf.Open)
	_adba.SetIfNotNil("\u004e\u0061\u006d\u0065", _ccaaf.Name)
	_adba.SetIfNotNil("\u0053\u0074\u0061t\u0065", _ccaaf.State)
	_adba.SetIfNotNil("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c", _ccaaf.StateModel)
	return _afe
}

// PdfFieldText represents a text field where user can enter text.
type PdfFieldText struct {
	*PdfField
	DA     *_ebb.PdfObjectString
	Q      *_ebb.PdfObjectInteger
	DS     *_ebb.PdfObjectString
	RV     _ebb.PdfObject
	MaxLen *_ebb.PdfObjectInteger
}

func (_faadd *PdfColorspaceDeviceCMYK) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
}

// SignatureHandlerDocMDPParams describe the specific parameters for the SignatureHandlerEx
// These parameters describe how to check the difference between revisions.
// Revisions of the document get from the PdfParser.
type SignatureHandlerDocMDPParams struct {
	Parser     *_ebb.PdfParser
	DiffPolicy _ac.DiffPolicy
}

// GetContentStreamWithEncoder returns the pattern cell's content stream and its encoder
func (_ceeff *PdfTilingPattern) GetContentStreamWithEncoder() ([]byte, _ebb.StreamEncoder, error) {
	_befa, _fdfe := _ceeff._dcddc.(*_ebb.PdfObjectStream)
	if !_fdfe {
		_eg.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ceeff._dcddc)
		return nil, nil, _ebb.ErrTypeError
	}
	_abcf, _gdeed := _ebb.DecodeStream(_befa)
	if _gdeed != nil {
		_eg.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020s\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0065\u0072\u0072\u003a\u0020\u0025\u0076", _gdeed)
		return nil, nil, _gdeed
	}
	_cgcgd, _gdeed := _ebb.NewEncoderFromStream(_befa)
	if _gdeed != nil {
		_eg.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020f\u0069\u006e\u0064\u0069\u006e\u0067 \u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _gdeed)
		return nil, nil, _gdeed
	}
	return _abcf, _cgcgd, nil
}

// Items returns all children outline items.
func (_dbbb *OutlineItem) Items() []*OutlineItem { return _dbbb.Entries }

// ToPdfObject returns the PDF representation of the tiling pattern.
func (_ccefa *PdfTilingPattern) ToPdfObject() _ebb.PdfObject {
	_ccefa.PdfPattern.ToPdfObject()
	_abbac := _ccefa.getDict()
	if _ccefa.PaintType != nil {
		_abbac.Set("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e", _ccefa.PaintType)
	}
	if _ccefa.TilingType != nil {
		_abbac.Set("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _ccefa.TilingType)
	}
	if _ccefa.BBox != nil {
		_abbac.Set("\u0042\u0042\u006f\u0078", _ccefa.BBox.ToPdfObject())
	}
	if _ccefa.XStep != nil {
		_abbac.Set("\u0058\u0053\u0074e\u0070", _ccefa.XStep)
	}
	if _ccefa.YStep != nil {
		_abbac.Set("\u0059\u0053\u0074e\u0070", _ccefa.YStep)
	}
	if _ccefa.Resources != nil {
		_abbac.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _ccefa.Resources.ToPdfObject())
	}
	if _ccefa.Matrix != nil {
		_abbac.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _ccefa.Matrix)
	}
	return _ccefa._dcddc
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

func (_cbee *PdfReader) newPdfAnnotationFromIndirectObject(_cca *_ebb.PdfIndirectObject) (*PdfAnnotation, error) {
	_fcc, _dad := _cca.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_dad {
		return nil, _bg.Errorf("\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _cbee._abbaca.GetModelFromPrimitive(_fcc); model != nil {
		_gdbd, _bebb := model.(*PdfAnnotation)
		if !_bebb {
			return nil, _bg.Errorf("\u0063\u0061\u0063\u0068\u0065\u0064 \u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050D\u0046\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e")
		}
		return _gdbd, nil
	}
	_bgea := &PdfAnnotation{}
	_bgea._bdcd = _cca
	_cbee._abbaca.Register(_fcc, _bgea)
	if _cgf := _fcc.Get("\u0054\u0079\u0070\u0065"); _cgf != nil {
		_abc, _dba := _cgf.(*_ebb.PdfObjectName)
		if !_dba {
			_eg.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _cgf)
		} else {
			if *_abc != "\u0041\u006e\u006eo\u0074" {
				_eg.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070\u0065\u0063\u0074\u0065d\u0020\u0054\u0079\u0070\u0065\u0020\u0021=\u0020\u0041\u006e\u006e\u006f\u0074\u0020\u0028\u0025\u0073\u0029", *_abc)
			}
		}
	}
	if _dgd := _fcc.Get("\u0052\u0065\u0063\u0074"); _dgd != nil {
		_bgea.Rect = _dgd
	}
	if _fefa := _fcc.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _fefa != nil {
		_bgea.Contents = _fefa
	}
	if _dgbf := _fcc.Get("\u0050"); _dgbf != nil {
		_bgea.P = _dgbf
	}
	if _ddac := _fcc.Get("\u004e\u004d"); _ddac != nil {
		_bgea.NM = _ddac
	}
	if _gaa := _fcc.Get("\u004d"); _gaa != nil {
		_bgea.M = _gaa
	}
	if _beg := _fcc.Get("\u0046"); _beg != nil {
		_bgea.F = _beg
	}
	if _fgae := _fcc.Get("\u0041\u0050"); _fgae != nil {
		_bgea.AP = _fgae
	}
	if _fcf := _fcc.Get("\u0041\u0053"); _fcf != nil {
		_bgea.AS = _fcf
	}
	if _daa := _fcc.Get("\u0042\u006f\u0072\u0064\u0065\u0072"); _daa != nil {
		_bgea.Border = _daa
	}
	if _fcfd := _fcc.Get("\u0043"); _fcfd != nil {
		_bgea.C = _fcfd
	}
	if _agd := _fcc.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074"); _agd != nil {
		_bgea.StructParent = _agd
	}
	if _dea := _fcc.Get("\u004f\u0043"); _dea != nil {
		_bgea.OC = _dea
	}
	_fccb := _fcc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")
	if _fccb == nil {
		_eg.Log.Debug("\u0057\u0041\u0052\u004e\u0049\u004e\u0047:\u0020\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079 \u0069s\u0073\u0075\u0065\u0020\u002d\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073u\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074\u0079p\u0065")
		_bgea._efd = nil
		return _bgea, nil
	}
	_gbeg, _aedg := _fccb.(*_ebb.PdfObjectName)
	if !_aedg {
		_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065 !\u003d\u0020n\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _fccb)
		return nil, _bg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d n\u0061\u006d\u0065 \u0028%\u0054\u0029", _fccb)
	}
	switch *_gbeg {
	case "\u0054\u0065\u0078\u0074":
		_ddgb, _dgcc := _cbee.newPdfAnnotationTextFromDict(_fcc)
		if _dgcc != nil {
			return nil, _dgcc
		}
		_ddgb.PdfAnnotation = _bgea
		_bgea._efd = _ddgb
		return _bgea, nil
	case "\u004c\u0069\u006e\u006b":
		_adf, _bdf := _cbee.newPdfAnnotationLinkFromDict(_fcc)
		if _bdf != nil {
			return nil, _bdf
		}
		_adf.PdfAnnotation = _bgea
		_bgea._efd = _adf
		return _bgea, nil
	case "\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074":
		_abcc, _dggd := _cbee.newPdfAnnotationFreeTextFromDict(_fcc)
		if _dggd != nil {
			return nil, _dggd
		}
		_abcc.PdfAnnotation = _bgea
		_bgea._efd = _abcc
		return _bgea, nil
	case "\u004c\u0069\u006e\u0065":
		_fdefc, _edf := _cbee.newPdfAnnotationLineFromDict(_fcc)
		if _edf != nil {
			return nil, _edf
		}
		_fdefc.PdfAnnotation = _bgea
		_bgea._efd = _fdefc
		_eg.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041N\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0061\u006e\u006eo\u0074\u0020\u0028\u0025\u0054\u0029\u003a \u0025\u002b\u0076\u000a", _bgea, _bgea)
		_eg.Log.Trace("\u004c\u0049\u004eE\u0020\u0041\u004e\u004eO\u0054\u0041\u0054\u0049\u004f\u004e\u003a \u0063\u0074\u0078\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u002b\u0076\u000a", _fdefc, _fdefc)
		_eg.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041\u004e\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u0020\u004d\u0061\u0072\u006b\u0075\u0070\u003a\u0020c\u0074\u0078\u0020\u0028\u0025T\u0029\u003a \u0025\u002b\u0076\u000a", _fdefc.PdfAnnotationMarkup, _fdefc.PdfAnnotationMarkup)
		return _bgea, nil
	case "\u0053\u0071\u0075\u0061\u0072\u0065":
		_dbb, _cbaa := _cbee.newPdfAnnotationSquareFromDict(_fcc)
		if _cbaa != nil {
			return nil, _cbaa
		}
		_dbb.PdfAnnotation = _bgea
		_bgea._efd = _dbb
		return _bgea, nil
	case "\u0043\u0069\u0072\u0063\u006c\u0065":
		_dadd, _ccea := _cbee.newPdfAnnotationCircleFromDict(_fcc)
		if _ccea != nil {
			return nil, _ccea
		}
		_dadd.PdfAnnotation = _bgea
		_bgea._efd = _dadd
		return _bgea, nil
	case "\u0050o\u006c\u0079\u0067\u006f\u006e":
		_gdcdf, _dff := _cbee.newPdfAnnotationPolygonFromDict(_fcc)
		if _dff != nil {
			return nil, _dff
		}
		_gdcdf.PdfAnnotation = _bgea
		_bgea._efd = _gdcdf
		return _bgea, nil
	case "\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065":
		_gab, _bcd := _cbee.newPdfAnnotationPolyLineFromDict(_fcc)
		if _bcd != nil {
			return nil, _bcd
		}
		_gab.PdfAnnotation = _bgea
		_bgea._efd = _gab
		return _bgea, nil
	case "\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t":
		_bgd, _fae := _cbee.newPdfAnnotationHighlightFromDict(_fcc)
		if _fae != nil {
			return nil, _fae
		}
		_bgd.PdfAnnotation = _bgea
		_bgea._efd = _bgd
		return _bgea, nil
	case "\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee":
		_ddfa, _abea := _cbee.newPdfAnnotationUnderlineFromDict(_fcc)
		if _abea != nil {
			return nil, _abea
		}
		_ddfa.PdfAnnotation = _bgea
		_bgea._efd = _ddfa
		return _bgea, nil
	case "\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079":
		_bdb, _cfad := _cbee.newPdfAnnotationSquigglyFromDict(_fcc)
		if _cfad != nil {
			return nil, _cfad
		}
		_bdb.PdfAnnotation = _bgea
		_bgea._efd = _bdb
		return _bgea, nil
	case "\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t":
		_afb, _cfe := _cbee.newPdfAnnotationStrikeOut(_fcc)
		if _cfe != nil {
			return nil, _cfe
		}
		_afb.PdfAnnotation = _bgea
		_bgea._efd = _afb
		return _bgea, nil
	case "\u0043\u0061\u0072e\u0074":
		_abgf, _ceff := _cbee.newPdfAnnotationCaretFromDict(_fcc)
		if _ceff != nil {
			return nil, _ceff
		}
		_abgf.PdfAnnotation = _bgea
		_bgea._efd = _abgf
		return _bgea, nil
	case "\u0053\u0074\u0061m\u0070":
		_dbad, _ada := _cbee.newPdfAnnotationStampFromDict(_fcc)
		if _ada != nil {
			return nil, _ada
		}
		_dbad.PdfAnnotation = _bgea
		_bgea._efd = _dbad
		return _bgea, nil
	case "\u0049\u006e\u006b":
		_fba, _gcf := _cbee.newPdfAnnotationInkFromDict(_fcc)
		if _gcf != nil {
			return nil, _gcf
		}
		_fba.PdfAnnotation = _bgea
		_bgea._efd = _fba
		return _bgea, nil
	case "\u0050\u006f\u0070u\u0070":
		_aga, _cbcc := _cbee.newPdfAnnotationPopupFromDict(_fcc)
		if _cbcc != nil {
			return nil, _cbcc
		}
		_aga.PdfAnnotation = _bgea
		_bgea._efd = _aga
		return _bgea, nil
	case "\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074":
		_fag, _abcg := _cbee.newPdfAnnotationFileAttachmentFromDict(_fcc)
		if _abcg != nil {
			return nil, _abcg
		}
		_fag.PdfAnnotation = _bgea
		_bgea._efd = _fag
		return _bgea, nil
	case "\u0053\u006f\u0075n\u0064":
		_ddgc, _dbag := _cbee.newPdfAnnotationSoundFromDict(_fcc)
		if _dbag != nil {
			return nil, _dbag
		}
		_ddgc.PdfAnnotation = _bgea
		_bgea._efd = _ddgc
		return _bgea, nil
	case "\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a":
		_ffg, _dabc := _cbee.newPdfAnnotationRichMediaFromDict(_fcc)
		if _dabc != nil {
			return nil, _dabc
		}
		_ffg.PdfAnnotation = _bgea
		_bgea._efd = _ffg
		return _bgea, nil
	case "\u004d\u006f\u0076i\u0065":
		_adb, _afbe := _cbee.newPdfAnnotationMovieFromDict(_fcc)
		if _afbe != nil {
			return nil, _afbe
		}
		_adb.PdfAnnotation = _bgea
		_bgea._efd = _adb
		return _bgea, nil
	case "\u0053\u0063\u0072\u0065\u0065\u006e":
		_cece, _cab := _cbee.newPdfAnnotationScreenFromDict(_fcc)
		if _cab != nil {
			return nil, _cab
		}
		_cece.PdfAnnotation = _bgea
		_bgea._efd = _cece
		return _bgea, nil
	case "\u0057\u0069\u0064\u0067\u0065\u0074":
		_cafg, _dbeb := _cbee.newPdfAnnotationWidgetFromDict(_fcc)
		if _dbeb != nil {
			return nil, _dbeb
		}
		_cafg.PdfAnnotation = _bgea
		_bgea._efd = _cafg
		return _bgea, nil
	case "P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b":
		_eaee, _faee := _cbee.newPdfAnnotationPrinterMarkFromDict(_fcc)
		if _faee != nil {
			return nil, _faee
		}
		_eaee.PdfAnnotation = _bgea
		_bgea._efd = _eaee
		return _bgea, nil
	case "\u0054r\u0061\u0070\u004e\u0065\u0074":
		_gfac, _bgg := _cbee.newPdfAnnotationTrapNetFromDict(_fcc)
		if _bgg != nil {
			return nil, _bgg
		}
		_gfac.PdfAnnotation = _bgea
		_bgea._efd = _gfac
		return _bgea, nil
	case "\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k":
		_fdb, _ccce := _cbee.newPdfAnnotationWatermarkFromDict(_fcc)
		if _ccce != nil {
			return nil, _ccce
		}
		_fdb.PdfAnnotation = _bgea
		_bgea._efd = _fdb
		return _bgea, nil
	case "\u0033\u0044":
		_fbfd, _eaa := _cbee.newPdfAnnotation3DFromDict(_fcc)
		if _eaa != nil {
			return nil, _eaa
		}
		_fbfd.PdfAnnotation = _bgea
		_bgea._efd = _fbfd
		return _bgea, nil
	case "\u0050\u0072\u006f\u006a\u0065\u0063\u0074\u0069\u006f\u006e":
		_adag, _gbaf := _cbee.newPdfAnnotationProjectionFromDict(_fcc)
		if _gbaf != nil {
			return nil, _gbaf
		}
		_adag.PdfAnnotation = _bgea
		_bgea._efd = _adag
		return _bgea, nil
	case "\u0052\u0065\u0064\u0061\u0063\u0074":
		_abeac, _cdbd := _cbee.newPdfAnnotationRedactFromDict(_fcc)
		if _cdbd != nil {
			return nil, _cdbd
		}
		_abeac.PdfAnnotation = _bgea
		_bgea._efd = _abeac
		return _bgea, nil
	}
	_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069\u006fn\u003a\u0020\u0025\u0073", *_gbeg)
	return nil, nil
}

// DecodeArray returns the component range values for the Separation colorspace.
func (_aeaf *PdfColorspaceSpecialSeparation) DecodeArray() []float64 { return []float64{0, 1.0} }

// ToPdfObject returns the PDF representation of the VRI dictionary.
func (_edcf *VRI) ToPdfObject() *_ebb.PdfObjectDictionary {
	_dbcg := _ebb.MakeDict()
	_dbcg.SetIfNotNil(_ebb.PdfObjectName("\u0043\u0065\u0072\u0074"), _agdbf(_edcf.Cert))
	_dbcg.SetIfNotNil(_ebb.PdfObjectName("\u004f\u0043\u0053\u0050"), _agdbf(_edcf.OCSP))
	_dbcg.SetIfNotNil(_ebb.PdfObjectName("\u0043\u0052\u004c"), _agdbf(_edcf.CRL))
	_dbcg.SetIfNotNil("\u0054\u0055", _edcf.TU)
	_dbcg.SetIfNotNil("\u0054\u0053", _edcf.TS)
	return _dbcg
}
func (_bcef *PdfReader) newPdfAnnotationWidgetFromDict(_bgfc *_ebb.PdfObjectDictionary) (*PdfAnnotationWidget, error) {
	_deac := PdfAnnotationWidget{}
	_deac.H = _bgfc.Get("\u0048")
	_deac.MK = _bgfc.Get("\u004d\u004b")
	_deac.A = _bgfc.Get("\u0041")
	_deac.AA = _bgfc.Get("\u0041\u0041")
	_deac.BS = _bgfc.Get("\u0042\u0053")
	_deac.Parent = _bgfc.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	return &_deac, nil
}

// PdfAnnotationCaret represents Caret annotations.
// (Section 12.5.6.11).
type PdfAnnotationCaret struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	RD _ebb.PdfObject
	Sy _ebb.PdfObject
}

func (_ggbef *PdfShading) getShadingDict() (*_ebb.PdfObjectDictionary, error) {
	_dddcg := _ggbef._fbfae
	if _cegfc, _bedb := _dddcg.(*_ebb.PdfIndirectObject); _bedb {
		_gfdcf, _cecd := _cegfc.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_cecd {
			return nil, _ebb.ErrTypeError
		}
		return _gfdcf, nil
	} else if _ggee, _defee := _dddcg.(*_ebb.PdfObjectStream); _defee {
		return _ggee.PdfObjectDictionary, nil
	} else if _dcbcad, _fefd := _dddcg.(*_ebb.PdfObjectDictionary); _fefd {
		return _dcbcad, nil
	} else {
		_eg.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0063\u0063\u0065s\u0073\u0020\u0073\u0068\u0061\u0064\u0069n\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079")
		return nil, _ebb.ErrTypeError
	}
}

// String returns a string representation of PdfTransformParamsDocMDP.
func (_egecf *PdfTransformParamsDocMDP) String() string {
	return _bg.Sprintf("\u0025\u0073\u0020\u0050\u003a\u0020\u0025\u0073\u0020V\u003a\u0020\u0025\u0073", _egecf.Type, _egecf.P, _egecf.V)
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

func _gfcg(_abafc *_ebb.PdfObjectStream) (*PdfFunctionType0, error) {
	_cbebf := &PdfFunctionType0{}
	_cbebf._afgcc = _abafc
	_gdaga := _abafc.PdfObjectDictionary
	_geca, _fdbbb := _ebb.TraceToDirectObject(_gdaga.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_ebb.PdfObjectArray)
	if !_fdbbb {
		_eg.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _geca.Len() < 0 || _geca.Len()%2 != 0 {
		_eg.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _gf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_cbebf.NumInputs = _geca.Len() / 2
	_bfacf, _cgcbc := _geca.ToFloat64Array()
	if _cgcbc != nil {
		return nil, _cgcbc
	}
	_cbebf.Domain = _bfacf
	_geca, _fdbbb = _ebb.TraceToDirectObject(_gdaga.Get("\u0052\u0061\u006eg\u0065")).(*_ebb.PdfObjectArray)
	if !_fdbbb {
		_eg.Log.Error("\u0052\u0061\u006e\u0067e \u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _geca.Len() < 0 || _geca.Len()%2 != 0 {
		return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_cbebf.NumOutputs = _geca.Len() / 2
	_abada, _cgcbc := _geca.ToFloat64Array()
	if _cgcbc != nil {
		return nil, _cgcbc
	}
	_cbebf.Range = _abada
	_geca, _fdbbb = _ebb.TraceToDirectObject(_gdaga.Get("\u0053\u0069\u007a\u0065")).(*_ebb.PdfObjectArray)
	if !_fdbbb {
		_eg.Log.Error("\u0053i\u007ae\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_dcae, _cgcbc := _geca.ToIntegerArray()
	if _cgcbc != nil {
		return nil, _cgcbc
	}
	if len(_dcae) != _cbebf.NumInputs {
		_eg.Log.Error("T\u0061\u0062\u006c\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063h\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072 o\u0066\u0020\u0069n\u0070u\u0074\u0073")
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cbebf.Size = _dcae
	_efdbe, _fdbbb := _ebb.TraceToDirectObject(_gdaga.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065")).(*_ebb.PdfObjectInteger)
	if !_fdbbb {
		_eg.Log.Error("B\u0069\u0074\u0073\u0050\u0065\u0072S\u0061\u006d\u0070\u006c\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if *_efdbe != 1 && *_efdbe != 2 && *_efdbe != 4 && *_efdbe != 8 && *_efdbe != 12 && *_efdbe != 16 && *_efdbe != 24 && *_efdbe != 32 {
		_eg.Log.Error("\u0042\u0069\u0074s \u0070\u0065\u0072\u0020\u0073\u0061\u006d\u0070\u006ce\u0020o\u0075t\u0073i\u0064\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064\u0029", *_efdbe)
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cbebf.BitsPerSample = int(*_efdbe)
	_cbebf.Order = 1
	_bcgb, _fdbbb := _ebb.TraceToDirectObject(_gdaga.Get("\u004f\u0072\u0064e\u0072")).(*_ebb.PdfObjectInteger)
	if _fdbbb {
		if *_bcgb != 1 && *_bcgb != 3 {
			_eg.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u006fr\u0064\u0065r\u0020\u0028\u0025\u0064\u0029", *_bcgb)
			return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		_cbebf.Order = int(*_bcgb)
	}
	_geca, _fdbbb = _ebb.TraceToDirectObject(_gdaga.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_ebb.PdfObjectArray)
	if _fdbbb {
		_aebb, _bdac := _geca.ToFloat64Array()
		if _bdac != nil {
			return nil, _bdac
		}
		_cbebf.Encode = _aebb
	}
	_geca, _fdbbb = _ebb.TraceToDirectObject(_gdaga.Get("\u0044\u0065\u0063\u006f\u0064\u0065")).(*_ebb.PdfObjectArray)
	if _fdbbb {
		_cfcgaf, _aaegc := _geca.ToFloat64Array()
		if _aaegc != nil {
			return nil, _aaegc
		}
		_cbebf.Decode = _cfcgaf
	}
	_fegad, _cgcbc := _ebb.DecodeStream(_abafc)
	if _cgcbc != nil {
		return nil, _cgcbc
	}
	_cbebf._aeecg = _fegad
	return _cbebf, nil
}

// PdfColorCalGray represents a CalGray colorspace.
type PdfColorCalGray float64

func (_gcfcd *PdfReader) newPdfSignatureFromIndirect(_bddba *_ebb.PdfIndirectObject) (*PdfSignature, error) {
	_ecacg, _cbgce := _bddba.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_cbgce {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		return nil, ErrTypeCheck
	}
	if _acdgb, _debff := _gcfcd._abbaca.GetModelFromPrimitive(_bddba).(*PdfSignature); _debff {
		return _acdgb, nil
	}
	_eggcbg := &PdfSignature{}
	_eggcbg._ffbgc = _bddba
	_eggcbg.Type, _ = _ebb.GetName(_ecacg.Get("\u0054\u0079\u0070\u0065"))
	_eggcbg.Filter, _cbgce = _ebb.GetName(_ecacg.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if !_cbgce {
		_eg.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053i\u0067\u006e\u0061\u0074\u0075r\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_eggcbg.SubFilter, _ = _ebb.GetName(_ecacg.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r"))
	_eggcbg.Contents, _cbgce = _ebb.GetString(_ecacg.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_cbgce {
		_eg.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	if _cbdbb, _cadgd := _ebb.GetArray(_ecacg.Get("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e")); _cadgd {
		_eggcbg.Reference = _ebb.MakeArray()
		for _, _egacg := range _cbdbb.Elements() {
			_geead, _eddc := _ebb.GetDict(_egacg)
			if !_eddc {
				_eg.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020R\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0061\u0074\u0065\u0064")
				return nil, ErrInvalidAttribute
			}
			_ffgd, _bcfb := _gcfcd.newPdfSignatureReferenceFromDict(_geead)
			if _bcfb != nil {
				return nil, _bcfb
			}
			_eggcbg.Reference.Append(_ffgd.ToPdfObject())
		}
	}
	_eggcbg.Cert = _ecacg.Get("\u0043\u0065\u0072\u0074")
	_eggcbg.ByteRange, _ = _ebb.GetArray(_ecacg.Get("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e"))
	_eggcbg.Changes, _ = _ebb.GetArray(_ecacg.Get("\u0043h\u0061\u006e\u0067\u0065\u0073"))
	_eggcbg.Name, _ = _ebb.GetString(_ecacg.Get("\u004e\u0061\u006d\u0065"))
	_eggcbg.M, _ = _ebb.GetString(_ecacg.Get("\u004d"))
	_eggcbg.Location, _ = _ebb.GetString(_ecacg.Get("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"))
	_eggcbg.Reason, _ = _ebb.GetString(_ecacg.Get("\u0052\u0065\u0061\u0073\u006f\u006e"))
	_eggcbg.ContactInfo, _ = _ebb.GetString(_ecacg.Get("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f"))
	_eggcbg.R, _ = _ebb.GetInt(_ecacg.Get("\u0052"))
	_eggcbg.V, _ = _ebb.GetInt(_ecacg.Get("\u0056"))
	_eggcbg.PropBuild, _ = _ebb.GetDict(_ecacg.Get("\u0050\u0072\u006f\u0070\u005f\u0042\u0075\u0069\u006c\u0064"))
	_eggcbg.PropAuthTime, _ = _ebb.GetInt(_ecacg.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0069\u006d\u0065"))
	_eggcbg.PropAuthType, _ = _ebb.GetName(_ecacg.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0079\u0070\u0065"))
	_gcfcd._abbaca.Register(_bddba, _eggcbg)
	return _eggcbg, nil
}
func (_gfaeb *DSS) addCerts(_aafcg [][]byte) ([]*_ebb.PdfObjectStream, error) {
	return _gfaeb.add(&_gfaeb.Certs, _gfaeb._aeag, _aafcg)
}

// GetFontByName gets the font specified by keyName. Returns the PdfObject which
// the entry refers to. Returns a bool value indicating whether or not the entry was found.
func (_ddgcd *PdfPageResources) GetFontByName(keyName _ebb.PdfObjectName) (_ebb.PdfObject, bool) {
	if _ddgcd.Font == nil {
		return nil, false
	}
	_aebge, _gefca := _ebb.TraceToDirectObject(_ddgcd.Font).(*_ebb.PdfObjectDictionary)
	if !_gefca {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _ebb.TraceToDirectObject(_ddgcd.Font))
		return nil, false
	}
	if _gdgcb := _aebge.Get(keyName); _gdgcb != nil {
		return _gdgcb, true
	}
	return nil, false
}

// NewPdfAnnotationPrinterMark returns a new printermark annotation.
func NewPdfAnnotationPrinterMark() *PdfAnnotationPrinterMark {
	_dda := NewPdfAnnotation()
	_edd := &PdfAnnotationPrinterMark{}
	_edd.PdfAnnotation = _dda
	_dda.SetContext(_edd)
	return _edd
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_daabe *PdfShadingType1) ToPdfObject() _ebb.PdfObject {
	_daabe.PdfShading.ToPdfObject()
	_dgfbd, _dcdea := _daabe.getShadingDict()
	if _dcdea != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _daabe.Domain != nil {
		_dgfbd.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _daabe.Domain)
	}
	if _daabe.Matrix != nil {
		_dgfbd.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _daabe.Matrix)
	}
	if _daabe.Function != nil {
		if len(_daabe.Function) == 1 {
			_dgfbd.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _daabe.Function[0].ToPdfObject())
		} else {
			_acfa := _ebb.MakeArray()
			for _, _agffe := range _daabe.Function {
				_acfa.Append(_agffe.ToPdfObject())
			}
			_dgfbd.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _acfa)
		}
	}
	return _daabe._fbfae
}

// NewBorderStyle returns an initialized PdfBorderStyle.
func NewBorderStyle() *PdfBorderStyle { _ccbe := &PdfBorderStyle{}; return _ccbe }

// PdfColorspaceDeviceN represents a DeviceN color space. DeviceN color spaces are similar to Separation color
// spaces, except they can contain an arbitrary number of color components.
//
// Format: [/DeviceN names alternateSpace tintTransform]
//     or: [/DeviceN names alternateSpace tintTransform attributes]
type PdfColorspaceDeviceN struct {
	ColorantNames  *_ebb.PdfObjectArray
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	Attributes     *PdfColorspaceDeviceNAttributes
	_gebb          *_ebb.PdfIndirectObject
}

// AddWatermarkImage adds a watermark to the page.
func (_dffcag *PdfPage) AddWatermarkImage(ximg *XObjectImage, opt WatermarkImageOptions) error {
	_ddag, _bfebe := _dffcag.GetMediaBox()
	if _bfebe != nil {
		return _bfebe
	}
	_ebdcaa := _ddag.Urx - _ddag.Llx
	_gbggg := _ddag.Ury - _ddag.Lly
	_bbeaeb := float64(*ximg.Width)
	_ebaef := (_ebdcaa - _bbeaeb) / 2
	if opt.FitToWidth {
		_bbeaeb = _ebdcaa
		_ebaef = 0
	}
	_cagbc := _gbggg
	_gddbaa := float64(0)
	if opt.PreserveAspectRatio {
		_cagbc = _bbeaeb * float64(*ximg.Height) / float64(*ximg.Width)
		_gddbaa = (_gbggg - _cagbc) / 2
	}
	if _dffcag.Resources == nil {
		_dffcag.Resources = NewPdfPageResources()
	}
	_bcab := 0
	_aadce := _ebb.PdfObjectName(_bg.Sprintf("\u0049\u006d\u0077%\u0064", _bcab))
	for _dffcag.Resources.HasXObjectByName(_aadce) {
		_bcab++
		_aadce = _ebb.PdfObjectName(_bg.Sprintf("\u0049\u006d\u0077%\u0064", _bcab))
	}
	_bfebe = _dffcag.AddImageResource(_aadce, ximg)
	if _bfebe != nil {
		return _bfebe
	}
	_bcab = 0
	_gfbde := _ebb.PdfObjectName(_bg.Sprintf("\u0047\u0053\u0025\u0064", _bcab))
	for _dffcag.HasExtGState(_gfbde) {
		_bcab++
		_gfbde = _ebb.PdfObjectName(_bg.Sprintf("\u0047\u0053\u0025\u0064", _bcab))
	}
	_bbcg := _ebb.MakeDict()
	_bbcg.Set("\u0042\u004d", _ebb.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	_bbcg.Set("\u0043\u0041", _ebb.MakeFloat(opt.Alpha))
	_bbcg.Set("\u0063\u0061", _ebb.MakeFloat(opt.Alpha))
	_bfebe = _dffcag.AddExtGState(_gfbde, _bbcg)
	if _bfebe != nil {
		return _bfebe
	}
	_baggg := _bg.Sprintf("\u0071\u000a"+"\u002f%\u0073\u0020\u0067\u0073\u000a"+"%\u002e\u0030\u0066\u0020\u0030\u00200\u0020\u0025\u002e\u0030\u0066\u0020\u0025\u002e\u0034f\u0020\u0025\u002e4\u0066 \u0063\u006d\u000a"+"\u002f%\u0073\u0020\u0044\u006f\u000a"+"\u0051", _gfbde, _bbeaeb, _cagbc, _ebaef, _gddbaa, _aadce)
	_dffcag.AddContentStreamByString(_baggg)
	return nil
}

// PdfAnnotationPrinterMark represents PrinterMark annotations.
// (Section 12.5.6.20).
type PdfAnnotationPrinterMark struct {
	*PdfAnnotation
	MN _ebb.PdfObject
}

func (_agba *PdfColorspaceCalGray) String() string { return "\u0043a\u006c\u0047\u0072\u0061\u0079" }

// ToPdfObject converts the PdfPage to a dictionary within an indirect object container.
func (_bgbd *PdfPage) ToPdfObject() _ebb.PdfObject {
	_adgfa := _bgbd._defbb
	_bgbd.GetPageDict()
	return _adgfa
}

// GetContext returns the annotation context which contains the specific type-dependent context.
// The context represents the subannotation.
func (_bfdb *PdfAnnotation) GetContext() PdfModel {
	if _bfdb == nil {
		return nil
	}
	return _bfdb._efd
}
func _efec(_ggdb _ebb.PdfObject) (*_ebb.PdfObjectDictionary, *fontCommon, error) {
	_dcdf := &fontCommon{}
	if _ebeef, _cgfb := _ggdb.(*_ebb.PdfIndirectObject); _cgfb {
		_dcdf._efbg = _ebeef.ObjectNumber
	}
	_ccdg, _gebgf := _ebb.GetDict(_ggdb)
	if !_gebgf {
		_eg.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _ggdb)
		return nil, nil, ErrFontNotSupported
	}
	_cacec, _gebgf := _ebb.GetNameVal(_ccdg.Get("\u0054\u0079\u0070\u0065"))
	if !_gebgf {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	if _cacec != "\u0046\u006f\u006e\u0074" {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u0020\u0049\u006e\u0063\u006f\u006d\u0070\u0061t\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u003d\u0025\u0071\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0025\u0071.", _cacec, "\u0046\u006f\u006e\u0074")
		return nil, nil, _ebb.ErrTypeError
	}
	_dbbg, _gebgf := _ebb.GetNameVal(_ccdg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_gebgf {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020F\u006f\u006e\u0074 \u0049\u006e\u0063o\u006d\u0070a\u0074\u0069\u0062\u0069\u006c\u0069t\u0079. \u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	_dcdf._dfbf = _dbbg
	_cfbe, _gebgf := _ebb.GetNameVal(_ccdg.Get("\u004e\u0061\u006d\u0065"))
	if _gebgf {
		_dcdf._efge = _cfbe
	}
	_fgbdb := _ccdg.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _fgbdb != nil {
		_dcdf._baag = _ebb.TraceToDirectObject(_fgbdb)
		_bgdc, _dbgg := _feef(_dcdf._baag, _dcdf)
		if _dbgg != nil {
			return _ccdg, _dcdf, _dbgg
		}
		_dcdf._dcdd = _bgdc
	} else if _dbbg == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" || _dbbg == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_bbgac, _ffed := _ebe.NewCIDSystemInfo(_ccdg.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
		if _ffed != nil {
			return _ccdg, _dcdf, _ffed
		}
		_acfcad := _bg.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0055\u0043\u0053\u0032", _bbgac.Registry, _bbgac.Ordering)
		if _ebe.IsPredefinedCMap(_acfcad) {
			_dcdf._dcdd, _ffed = _ebe.LoadPredefinedCMap(_acfcad)
			if _ffed != nil {
				_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _acfcad, _ffed)
			}
		}
	}
	_beaed := _ccdg.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
	if _beaed != nil {
		_feacg, _caaf := _acdf(_beaed)
		if _caaf != nil {
			_eg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0042\u0061\u0064\u0020\u0066\u006f\u006et\u0020d\u0065s\u0063r\u0069\u0070\u0074\u006f\u0072\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _caaf)
			return _ccdg, _dcdf, _caaf
		}
		_dcdf._fbbd = _feacg
	}
	if _dbbg != "\u0054\u0079\u0070e\u0033" {
		_dbdbf, _acccb := _ebb.GetNameVal(_ccdg.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
		if !_acccb {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u006f\u006et\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069t\u0079\u002e\u0020\u0042\u0061se\u0046\u006f\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
			return _ccdg, _dcdf, ErrRequiredAttributeMissing
		}
		_dcdf._fdacg = _dbdbf
	}
	return _ccdg, _dcdf, nil
}

// PdfAnnotationInk represents Ink annotations.
// (Section 12.5.6.13).
type PdfAnnotationInk struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	InkList _ebb.PdfObject
	BS      _ebb.PdfObject
}

// ToPdfObject converts the pdfFontSimple to its PDF representation for outputting.
func (_gdcbd *pdfFontSimple) ToPdfObject() _ebb.PdfObject {
	if _gdcbd._fafaf == nil {
		_gdcbd._fafaf = &_ebb.PdfIndirectObject{}
	}
	_fbbb := _gdcbd.baseFields().asPdfObjectDictionary("")
	_gdcbd._fafaf.PdfObject = _fbbb
	if _gdcbd.FirstChar != nil {
		_fbbb.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _gdcbd.FirstChar)
	}
	if _gdcbd.LastChar != nil {
		_fbbb.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _gdcbd.LastChar)
	}
	if _gdcbd.Widths != nil {
		_fbbb.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _gdcbd.Widths)
	}
	if _gdcbd.Encoding != nil {
		_fbbb.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _gdcbd.Encoding)
	} else if _gdcbd._ebcb != nil {
		_dbced := _gdcbd._ebcb.ToPdfObject()
		if _dbced != nil {
			_fbbb.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _dbced)
		}
	}
	return _gdcbd._fafaf
}

// ToJBIG2Image converts current image to the core.JBIG2Image.
func (_bagb *Image) ToJBIG2Image() (*_ebb.JBIG2Image, error) {
	_fafdf, _cffg := _bagb.ToGoImage()
	if _cffg != nil {
		return nil, _cffg
	}
	return _ebb.GoImageToJBIG2(_fafdf, _ebb.JB2ImageAutoThreshold)
}

// GetNumComponents returns the number of color components (1 for Separation).
func (_dagfb *PdfColorspaceSpecialSeparation) GetNumComponents() int { return 1 }

// ToPdfObject returns the PDF representation of the page resources.
func (_aaedd *PdfPageResources) ToPdfObject() _ebb.PdfObject {
	_ggcea := _aaedd._efbed
	_ggcea.SetIfNotNil("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _aaedd.ExtGState)
	if _aaedd._aaee != nil {
		_aaedd.ColorSpace = _aaedd._aaee.ToPdfObject()
	}
	_ggcea.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _aaedd.ColorSpace)
	_ggcea.SetIfNotNil("\u0050a\u0074\u0074\u0065\u0072\u006e", _aaedd.Pattern)
	_ggcea.SetIfNotNil("\u0053h\u0061\u0064\u0069\u006e\u0067", _aaedd.Shading)
	_ggcea.SetIfNotNil("\u0058O\u0062\u006a\u0065\u0063\u0074", _aaedd.XObject)
	_ggcea.SetIfNotNil("\u0046\u006f\u006e\u0074", _aaedd.Font)
	_ggcea.SetIfNotNil("\u0050r\u006f\u0063\u0053\u0065\u0074", _aaedd.ProcSet)
	_ggcea.SetIfNotNil("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", _aaedd.Properties)
	return _ggcea
}

// NewPdfColorspaceDeviceRGB returns a new RGB colorspace object.
func NewPdfColorspaceDeviceRGB() *PdfColorspaceDeviceRGB { return &PdfColorspaceDeviceRGB{} }

// NewPdfActionJavaScript returns a new "javaScript" action.
func NewPdfActionJavaScript() *PdfActionJavaScript {
	_aab := NewPdfAction()
	_efc := &PdfActionJavaScript{}
	_efc.PdfAction = _aab
	_aab.SetContext(_efc)
	return _efc
}

// SetNamedDestinations sets the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_dbaac *PdfWriter) SetNamedDestinations(dests _ebb.PdfObject) error {
	if dests == nil {
		return nil
	}
	_eg.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0044\u0065\u0073\u0074\u0073\u002e\u002e\u002e")
	_dbaac._dffegd.Set("\u0044\u0065\u0073t\u0073", dests)
	return _dbaac.addObjects(dests)
}

// ReplacePage replaces the original page to a new page.
func (_aegbg *PdfAppender) ReplacePage(pageNum int, page *PdfPage) {
	_ecca := pageNum - 1
	for _fefag := range _aegbg._dfbg {
		if _fefag == _ecca {
			_bacd := page.Duplicate()
			_bfegc(_bacd)
			_aegbg._dfbg[_fefag] = _bacd
		}
	}
}

// GetContext returns the action context which contains the specific type-dependent context.
// The context represents the subaction.
func (_ega *PdfAction) GetContext() PdfModel {
	if _ega == nil {
		return nil
	}
	return _ega._ad
}

// PdfColorspaceSpecialSeparation is a Separation colorspace.
// At the moment the colour space is set to a Separation space, the conforming reader shall determine whether the
// device has an available colorant (e.g. dye) corresponding to the name of the requested space. If so, the conforming
// reader shall ignore the alternateSpace and tintTransform parameters; subsequent painting operations within the
// space shall apply the designated colorant directly, according to the tint values supplied.
//
// Format: [/Separation name alternateSpace tintTransform]
type PdfColorspaceSpecialSeparation struct {
	ColorantName   *_ebb.PdfObjectName
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	_cded          *_ebb.PdfIndirectObject
}

func (_gedcb *pdfCIDFontType2) baseFields() *fontCommon { return &_gedcb.fontCommon }

// ToPdfObject convert PdfInfo to pdf object.
func (_geab *PdfInfo) ToPdfObject() _ebb.PdfObject {
	_dfgc := _ebb.MakeDict()
	_dfgc.SetIfNotNil("\u0054\u0069\u0074l\u0065", _geab.Title)
	_dfgc.SetIfNotNil("\u0041\u0075\u0074\u0068\u006f\u0072", _geab.Author)
	_dfgc.SetIfNotNil("\u0053u\u0062\u006a\u0065\u0063\u0074", _geab.Subject)
	_dfgc.SetIfNotNil("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _geab.Keywords)
	_dfgc.SetIfNotNil("\u0043r\u0065\u0061\u0074\u006f\u0072", _geab.Creator)
	_dfgc.SetIfNotNil("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _geab.Producer)
	_dfgc.SetIfNotNil("\u0054r\u0061\u0070\u0070\u0065\u0064", _geab.Trapped)
	if _geab.CreationDate != nil {
		_dfgc.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _geab.CreationDate.ToPdfObject())
	}
	if _geab.ModifiedDate != nil {
		_dfgc.SetIfNotNil("\u004do\u0064\u0044\u0061\u0074\u0065", _geab.ModifiedDate.ToPdfObject())
	}
	for _, _abef := range _geab._gcgf.Keys() {
		_dfgc.SetIfNotNil(_abef, _geab._gcgf.Get(_abef))
	}
	return _dfgc
}

// AnnotFilterFunc represents a PDF annotation filtering function. If the function
// returns true, the annotation is kept, otherwise it is discarded.
type AnnotFilterFunc func(*PdfAnnotation) bool

// ConvertToBinary converts current image into binary (bi-level) format.
// Binary images are composed of single bits per pixel (only black or white).
// If provided image has more color components, then it would be converted into binary image using
// histogram auto threshold function.
func (_cffcf *Image) ConvertToBinary() error {
	if _cffcf.ColorComponents == 1 && _cffcf.BitsPerComponent == 1 {
		return nil
	}
	_gdfc, _cfgabd := _cffcf.ToGoImage()
	if _cfgabd != nil {
		return _cfgabd
	}
	_edbcd, _cfgabd := _dg.MonochromeConverter.Convert(_gdfc)
	if _cfgabd != nil {
		return _cfgabd
	}
	_cffcf.Data = _edbcd.Base().Data
	_cffcf._dagcb, _cfgabd = _dg.ScaleAlphaToMonochrome(_cffcf._dagcb, int(_cffcf.Width), int(_cffcf.Height))
	if _cfgabd != nil {
		return _cfgabd
	}
	_cffcf.BitsPerComponent = 1
	_cffcf.ColorComponents = 1
	_cffcf._dgcea = nil
	return nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// A, B and C components of the color. The values of the elements should be
// between 0 and 1.
func (_efde *PdfColorspaceCalRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bddg := vals[0]
	if _bddg < 0.0 || _bddg > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bddg)
		return nil, ErrColorOutOfRange
	}
	_fgadaf := vals[1]
	if _fgadaf < 0.0 || _fgadaf > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fgadaf)
		return nil, ErrColorOutOfRange
	}
	_eebb := vals[2]
	if _eebb < 0.0 || _eebb > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _eebb)
		return nil, ErrColorOutOfRange
	}
	_edaec := NewPdfColorCalRGB(_bddg, _fgadaf, _eebb)
	return _edaec, nil
}
func _ebcfc(_bdce *_ebb.PdfObjectDictionary) (*PdfShadingType5, error) {
	_bgeea := PdfShadingType5{}
	_ecfe := _bdce.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _ecfe == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_afdad, _dbcc := _ecfe.(*_ebb.PdfObjectInteger)
	if !_dbcc {
		_eg.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _ecfe)
		return nil, _ebb.ErrTypeError
	}
	_bgeea.BitsPerCoordinate = _afdad
	_ecfe = _bdce.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _ecfe == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_afdad, _dbcc = _ecfe.(*_ebb.PdfObjectInteger)
	if !_dbcc {
		_eg.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _ecfe)
		return nil, _ebb.ErrTypeError
	}
	_bgeea.BitsPerComponent = _afdad
	_ecfe = _bdce.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077")
	if _ecfe == nil {
		_eg.Log.Debug("\u0052\u0065\u0071u\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0056\u0065\u0072\u0074\u0069c\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077")
		return nil, ErrRequiredAttributeMissing
	}
	_afdad, _dbcc = _ecfe.(*_ebb.PdfObjectInteger)
	if !_dbcc {
		_eg.Log.Debug("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006et\u0065\u0067\u0065\u0072\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _ecfe)
		return nil, _ebb.ErrTypeError
	}
	_bgeea.VerticesPerRow = _afdad
	_ecfe = _bdce.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _ecfe == nil {
		_eg.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_bgdagc, _dbcc := _ecfe.(*_ebb.PdfObjectArray)
	if !_dbcc {
		_eg.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ecfe)
		return nil, _ebb.ErrTypeError
	}
	_bgeea.Decode = _bgdagc
	if _bebgg := _bdce.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _bebgg != nil {
		_bgeea.Function = []PdfFunction{}
		if _aecg, _eccbg := _bebgg.(*_ebb.PdfObjectArray); _eccbg {
			for _, _fgcda := range _aecg.Elements() {
				_gcce, _bdgaf := _aagg(_fgcda)
				if _bdgaf != nil {
					_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bdgaf)
					return nil, _bdgaf
				}
				_bgeea.Function = append(_bgeea.Function, _gcce)
			}
		} else {
			_bfbgfd, _agefg := _aagg(_bebgg)
			if _agefg != nil {
				_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _agefg)
				return nil, _agefg
			}
			_bgeea.Function = append(_bgeea.Function, _bfbgfd)
		}
	}
	return &_bgeea, nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_ecfc *PdfColorspaceDeviceGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gdaa := vals[0]
	if _gdaa < 0.0 || _gdaa > 1.0 {
		_eg.Log.Debug("\u0049\u006eco\u006d\u0070\u0061t\u0069\u0062\u0069\u006city\u003a R\u0061\u006e\u0067\u0065\u0020\u006f\u0075ts\u0069\u0064\u0065\u0020\u005b\u0030\u002c1\u005d")
	}
	if _gdaa < 0.0 {
		_gdaa = 0.0
	} else if _gdaa > 1.0 {
		_gdaa = 1.0
	}
	return NewPdfColorDeviceGray(_gdaa), nil
}

// NewPdfActionURI returns a new "Uri" action.
func NewPdfActionURI() *PdfActionURI {
	_ccc := NewPdfAction()
	_cac := &PdfActionURI{}
	_cac.PdfAction = _ccc
	_ccc.SetContext(_cac)
	return _cac
}

// Items returns all children outline items.
func (_agfd *Outline) Items() []*OutlineItem { return _agfd.Entries }

// NewLTV returns a new LTV client.
func NewLTV(appender *PdfAppender) (*LTV, error) {
	_dfcda := appender.Reader.DSS
	if _dfcda == nil {
		_dfcda = NewDSS()
	}
	if _bdage := _dfcda.generateHashMaps(); _bdage != nil {
		return nil, _bdage
	}
	return &LTV{CertClient: _cg.NewCertClient(), OCSPClient: _cg.NewOCSPClient(), CRLClient: _cg.NewCRLClient(), SkipExisting: true, _ggdbg: appender, _dfdgf: _dfcda}, nil
}

// NewPdfColorDeviceCMYK returns a new CMYK32 color.
func NewPdfColorDeviceCMYK(c, m, y, k float64) *PdfColorDeviceCMYK {
	_gacf := PdfColorDeviceCMYK{c, m, y, k}
	return &_gacf
}

// GetRevisionNumber returns the version of the current Pdf document
func (_gafgdc *PdfReader) GetRevisionNumber() int { return _gafgdc._cafdf.GetRevisionNumber() }

// PdfActionGoToE represents a GoToE action.
type PdfActionGoToE struct {
	*PdfAction
	F         *PdfFilespec
	D         _ebb.PdfObject
	NewWindow _ebb.PdfObject
	T         _ebb.PdfObject
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a grayscale device.
func (_ffb *PdfColorspaceDeviceGray) GetNumComponents() int { return 1 }

// SetAnnotations sets the annotations list.
func (_bbcff *PdfPage) SetAnnotations(annotations []*PdfAnnotation) { _bbcff._bbfed = annotations }
func _ebafc() string                                                { return _eg.Version }

// EncryptOptions represents encryption options for an output PDF.
type EncryptOptions struct {
	Permissions _fe.Permissions
	Algorithm   EncryptionAlgorithm
}

// NewPdfColorspaceCalGray returns a new CalGray colorspace object.
func NewPdfColorspaceCalGray() *PdfColorspaceCalGray {
	_ecdg := &PdfColorspaceCalGray{}
	_ecdg.BlackPoint = []float64{0.0, 0.0, 0.0}
	_ecdg.Gamma = 1
	return _ecdg
}
func (_gbace *LTV) getCRLs(_gfbbd []*_g.Certificate) ([][]byte, error) {
	_dbcgb := make([][]byte, 0, len(_gfbbd))
	for _, _fbaag := range _gfbbd {
		for _, _gggffg := range _fbaag.CRLDistributionPoints {
			if _gbace.CertClient.IsCA(_fbaag) {
				continue
			}
			_gfgf, _fgab := _gbace.CRLClient.MakeRequest(_gggffg, _fbaag)
			if _fgab != nil {
				_eg.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _fgab)
				continue
			}
			_dbcgb = append(_dbcgb, _gfgf)
		}
	}
	return _dbcgb, nil
}

// NewPdfReader returns a new PdfReader for an input io.ReadSeeker interface. Can be used to read PDF from
// memory or file. Immediately loads and traverses the PDF structure including pages and page contents (if
// not encrypted). Loads entire document structure into memory.
// Alternatively a lazy-loading reader can be created with NewPdfReaderLazy which loads only references,
// and references are loaded from disk into memory on an as-needed basis.
func NewPdfReader(rs _ab.ReadSeeker) (*PdfReader, error) {
	const _baaedfc = "\u006do\u0064e\u006c\u003a\u004e\u0065\u0077P\u0064\u0066R\u0065\u0061\u0064\u0065\u0072"
	return _dcbd(rs, &ReaderOpts{}, false, _baaedfc)
}

// SetDocInfo set document info.
// This will overwrite any globally declared document info.
func (_deggb *PdfWriter) SetDocInfo(info *PdfInfo) { _deggb.setDocInfo(info.ToPdfObject()) }

// ToPdfObject return the CalGray colorspace as a PDF object (name dictionary).
func (_agec *PdfColorspaceCalGray) ToPdfObject() _ebb.PdfObject {
	_adfca := &_ebb.PdfObjectArray{}
	_adfca.Append(_ebb.MakeName("\u0043a\u006c\u0047\u0072\u0061\u0079"))
	_fgcc := _ebb.MakeDict()
	if _agec.WhitePoint != nil {
		_fgcc.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _ebb.MakeArray(_ebb.MakeFloat(_agec.WhitePoint[0]), _ebb.MakeFloat(_agec.WhitePoint[1]), _ebb.MakeFloat(_agec.WhitePoint[2])))
	} else {
		_eg.Log.Error("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006fi\u006e\u0074\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _agec.BlackPoint != nil {
		_fgcc.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _ebb.MakeArray(_ebb.MakeFloat(_agec.BlackPoint[0]), _ebb.MakeFloat(_agec.BlackPoint[1]), _ebb.MakeFloat(_agec.BlackPoint[2])))
	}
	_fgcc.Set("\u0047\u0061\u006dm\u0061", _ebb.MakeFloat(_agec.Gamma))
	_adfca.Append(_fgcc)
	if _agec._bebfd != nil {
		_agec._bebfd.PdfObject = _adfca
		return _agec._bebfd
	}
	return _adfca
}
func (_fadge *PdfWriter) writeObjectsInStreams(_gfaa map[_ebb.PdfObject]bool) error {
	for _, _ebegab := range _fadge._ebdgg {
		if _ggfcc := _gfaa[_ebegab]; _ggfcc {
			continue
		}
		_faaea := int64(0)
		switch _fgddb := _ebegab.(type) {
		case *_ebb.PdfIndirectObject:
			_faaea = _fgddb.ObjectNumber
		case *_ebb.PdfObjectStream:
			_faaea = _fgddb.ObjectNumber
		case *_ebb.PdfObjectStreams:
			_faaea = _fgddb.ObjectNumber
		default:
			_eg.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0055n\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0074\u0079\u0070\u0065 \u0069\u006e\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0054", _ebegab)
			return ErrTypeCheck
		}
		if _fadge._cgfde != nil && _ebegab != _fadge._cbcaa {
			_fegab := _fadge._cgfde.Encrypt(_ebegab, _faaea, 0)
			if _fegab != nil {
				_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067\u0020(%\u0073\u0029", _fegab)
				return _fegab
			}
		}
		_fadge.writeObject(int(_faaea), _ebegab)
	}
	return nil
}

// PdfReader represents a PDF file reader. It is a frontend to the lower level parsing mechanism and provides
// a higher level access to work with PDF structure and information, such as the page structure etc.
type PdfReader struct {
	_cafdf   *_ebb.PdfParser
	_agbbe   _ebb.PdfObject
	_eedbb   *_ebb.PdfIndirectObject
	_egea    *_ebb.PdfObjectDictionary
	_faebb   []*_ebb.PdfIndirectObject
	PageList []*PdfPage
	_aadcb   int
	_fdgda   *_ebb.PdfObjectDictionary
	_fgbcg   *PdfOutlineTreeNode
	AcroForm *PdfAcroForm
	DSS      *DSS
	Rotate   *int64
	_fdca    *Permissions
	_cacfc   map[*PdfReader]*PdfReader
	_face    []*PdfReader
	_abbaca  *modelManager
	_ceefa   bool
	_dfadc   map[_ebb.PdfObject]struct{}
	_ggdg    _ab.ReadSeeker
	_decdd   string
	_cfbgga  bool
	_edcbc   *ReaderOpts
	_abadec  bool
}

// ToPdfObject implements interface PdfModel.
func (_bgbg *PdfAnnotationStamp) ToPdfObject() _ebb.PdfObject {
	_bgbg.PdfAnnotation.ToPdfObject()
	_fbad := _bgbg._bdcd
	_adcf := _fbad.PdfObject.(*_ebb.PdfObjectDictionary)
	_bgbg.PdfAnnotationMarkup.appendToPdfDictionary(_adcf)
	_adcf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0053\u0074\u0061m\u0070"))
	_adcf.SetIfNotNil("\u004e\u0061\u006d\u0065", _bgbg.Name)
	return _fbad
}

// ToPdfObject returns the PDF representation of the DSS dictionary.
func (_acdae *DSS) ToPdfObject() _ebb.PdfObject {
	_ffgaf := _acdae._fcgb.PdfObject.(*_ebb.PdfObjectDictionary)
	_ffgaf.Clear()
	_adae := _ebb.MakeDict()
	for _cdee, _ecfb := range _acdae.VRI {
		_adae.Set(*_ebb.MakeName(_cdee), _ecfb.ToPdfObject())
	}
	_ffgaf.SetIfNotNil("\u0043\u0065\u0072t\u0073", _agdbf(_acdae.Certs))
	_ffgaf.SetIfNotNil("\u004f\u0043\u0053P\u0073", _agdbf(_acdae.OCSPs))
	_ffgaf.SetIfNotNil("\u0043\u0052\u004c\u0073", _agdbf(_acdae.CRLs))
	_ffgaf.Set("\u0056\u0052\u0049", _adae)
	return _acdae._fcgb
}
func (_fcgfe *pdfFontSimple) getFontDescriptor() *PdfFontDescriptor {
	if _bggba := _fcgfe._fbbd; _bggba != nil {
		return _bggba
	}
	return _fcgfe._adbd
}

// NewPdfActionThread returns a new "thread" action.
func NewPdfActionThread() *PdfActionThread {
	_bcc := NewPdfAction()
	_cgd := &PdfActionThread{}
	_cgd.PdfAction = _bcc
	_bcc.SetContext(_cgd)
	return _cgd
}
func _bgfag(_agdb *PdfField, _ddfe _ebb.PdfObject) {
	for _, _bafdf := range _agdb.Annotations {
		_bafdf.AS = _ddfe
		_bafdf.ToPdfObject()
	}
}

// String returns a string describing the font descriptor.
func (_dagc *PdfFontDescriptor) String() string {
	var _bfffd []string
	if _dagc.FontName != nil {
		_bfffd = append(_bfffd, _dagc.FontName.String())
	}
	if _dagc.FontFamily != nil {
		_bfffd = append(_bfffd, _dagc.FontFamily.String())
	}
	if _dagc.fontFile != nil {
		_bfffd = append(_bfffd, _dagc.fontFile.String())
	}
	if _dagc._aeeb != nil {
		_bfffd = append(_bfffd, _dagc._aeeb.String())
	}
	_bfffd = append(_bfffd, _bg.Sprintf("\u0046\u006f\u006et\u0046\u0069\u006c\u0065\u0033\u003d\u0025\u0074", _dagc.FontFile3 != nil))
	return _bg.Sprintf("\u0046\u004f\u004e\u0054_D\u0045\u0053\u0043\u0052\u0049\u0050\u0054\u004f\u0052\u007b\u0025\u0073\u007d", _ee.Join(_bfffd, "\u002c\u0020"))
}

// GetContainingPdfObject returns the container of the pattern object (indirect object).
func (_bdafbb *PdfPattern) GetContainingPdfObject() _ebb.PdfObject { return _bdafbb._dcddc }

// ToImage converts an object to an Image which can be transformed or saved out.
// The image data is decoded and the Image returned.
func (_befed *XObjectImage) ToImage() (*Image, error) {
	_defagc := &Image{}
	if _befed.Height == nil {
		return nil, _gf.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_defagc.Height = *_befed.Height
	if _befed.Width == nil {
		return nil, _gf.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_defagc.Width = *_befed.Width
	if _befed.BitsPerComponent == nil {
		switch _befed.Filter.(type) {
		case *_ebb.CCITTFaxEncoder, *_ebb.JBIG2Encoder:
			_defagc.BitsPerComponent = 1
		case *_ebb.LZWEncoder, *_ebb.RunLengthEncoder:
			_defagc.BitsPerComponent = 8
		default:
			return nil, _gf.New("\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006fm\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
	} else {
		_defagc.BitsPerComponent = *_befed.BitsPerComponent
	}
	_defagc.ColorComponents = _befed.ColorSpace.GetNumComponents()
	_befed._fbeec.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _ebb.MakeInteger(int64(_defagc.ColorComponents)))
	_fdedc, _badffe := _ebb.DecodeStream(_befed._fbeec)
	if _badffe != nil {
		return nil, _badffe
	}
	_defagc.Data = _fdedc
	if _befed.Decode != nil {
		_agbf, _egbf := _befed.Decode.(*_ebb.PdfObjectArray)
		if !_egbf {
			_eg.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020D\u0065\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074")
			return nil, _gf.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_cbedc, _fdfdc := _agbf.ToFloat64Array()
		if _fdfdc != nil {
			return nil, _fdfdc
		}
		_defagc._dgcea = _cbedc
	}
	return _defagc, nil
}
func (_cabfg *PdfReader) resolveReference(_fdabb *_ebb.PdfObjectReference) (_ebb.PdfObject, bool, error) {
	_dged, _bdda := _cabfg._cafdf.ObjCache[int(_fdabb.ObjectNumber)]
	if !_bdda {
		_eg.Log.Trace("R\u0065\u0061\u0064\u0065r \u004co\u006f\u006b\u0075\u0070\u0020r\u0065\u0066\u003a\u0020\u0025\u0073", _fdabb)
		_aaafcb, _ceceb := _cabfg._cafdf.LookupByReference(*_fdabb)
		if _ceceb != nil {
			return nil, false, _ceceb
		}
		_cabfg._cafdf.ObjCache[int(_fdabb.ObjectNumber)] = _aaafcb
		return _aaafcb, false, nil
	}
	return _dged, true, nil
}

// Val returns the value of the color.
func (_dcfdg *PdfColorCalGray) Val() float64 { return float64(*_dcfdg) }
func _fagea(_ceeb *_ebb.PdfObjectDictionary, _gggc *fontCommon) (*pdfCIDFontType0, error) {
	if _gggc._dfbf != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" {
		_eg.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0030\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _gggc)
		return nil, _ebb.ErrRangeError
	}
	_gebdb := _afdc(_gggc)
	_afcgf, _gaccg := _ebb.GetDict(_ceeb.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_gaccg {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _gggc)
		return nil, ErrRequiredAttributeMissing
	}
	_gebdb.CIDSystemInfo = _afcgf
	_gebdb.DW = _ceeb.Get("\u0044\u0057")
	_gebdb.W = _ceeb.Get("\u0057")
	_gebdb.DW2 = _ceeb.Get("\u0044\u0057\u0032")
	_gebdb.W2 = _ceeb.Get("\u0057\u0032")
	_gebdb._gbdb = 1000.0
	if _ccecb, _fedc := _ebb.GetNumberAsFloat(_gebdb.DW); _fedc == nil {
		_gebdb._gbdb = _ccecb
	}
	_dffeg, _fbdg := _egbeb(_gebdb.W)
	if _fbdg != nil {
		return nil, _fbdg
	}
	if _dffeg == nil {
		_dffeg = map[_da.CharCode]float64{}
	}
	_gebdb._afcac = _dffeg
	return _gebdb, nil
}
func (_fceg *PdfReader) loadOutlines() (*PdfOutlineTreeNode, error) {
	if _fceg._cafdf.GetCrypter() != nil && !_fceg._cafdf.IsAuthenticated() {
		return nil, _bg.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_daece := _fceg._fdgda
	_febec := _daece.Get("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	if _febec == nil {
		return nil, nil
	}
	_eg.Log.Trace("\u002d\u0048\u0061\u0073\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	_edfc := _ebb.ResolveReference(_febec)
	_eg.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0020\u0072\u006f\u006f\u0074\u003a\u0020\u0025\u0076", _edfc)
	if _dadf := _ebb.IsNullObject(_edfc); _dadf {
		_eg.Log.Trace("\u004f\u0075\u0074li\u006e\u0065\u0020\u0072\u006f\u006f\u0074\u0020\u0069s\u0020n\u0075l\u006c \u002d\u0020\u006e\u006f\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
		return nil, nil
	}
	_fedb, _ecff := _edfc.(*_ebb.PdfIndirectObject)
	if !_ecff {
		if _, _fadcb := _ebb.GetDict(_edfc); !_fadcb {
			_eg.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			return nil, nil
		}
		_eg.Log.Debug("\u004f\u0075t\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u0069s\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		_fedb = _ebb.MakeIndirectObject(_edfc)
	}
	_eedc, _ecff := _fedb.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_ecff {
		return nil, _gf.New("\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y")
	}
	_eg.Log.Trace("O\u0075\u0074\u006c\u0069ne\u0020r\u006f\u006f\u0074\u0020\u0064i\u0063\u0074\u003a\u0020\u0025\u0076", _eedc)
	_febc, _, _gcfd := _fceg.buildOutlineTree(_fedb, nil, nil, nil)
	if _gcfd != nil {
		return nil, _gcfd
	}
	_eg.Log.Trace("\u0052\u0065\u0073\u0075\u006c\u0074\u0069\u006e\u0067\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065:\u0020\u0025\u0076", _febc)
	return _febc, nil
}

// NewPdfAnnotationWidget returns an initialized annotation widget.
func NewPdfAnnotationWidget() *PdfAnnotationWidget {
	_badg := NewPdfAnnotation()
	_ccec := &PdfAnnotationWidget{}
	_ccec.PdfAnnotation = _badg
	_badg.SetContext(_ccec)
	return _ccec
}

// PdfColorspaceSpecialPattern is a Pattern colorspace.
// Can be defined either as /Pattern or with an underlying colorspace [/Pattern cs].
type PdfColorspaceSpecialPattern struct {
	UnderlyingCS PdfColorspace
	_adegb       *_ebb.PdfIndirectObject
}

// AcroFormRepairOptions contains options for rebuilding the AcroForm.
type AcroFormRepairOptions struct{}

// GetNamedDestinations returns the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_begda *PdfReader) GetNamedDestinations() (_ebb.PdfObject, error) {
	_fgaa := _ebb.ResolveReference(_begda._fdgda.Get("\u0044\u0065\u0073t\u0073"))
	if _fgaa == nil {
		return nil, nil
	}
	if !_begda._ceefa {
		_bbgdb := _begda.traverseObjectData(_fgaa)
		if _bbgdb != nil {
			return nil, _bbgdb
		}
	}
	return _fgaa, nil
}

// NewPdfAnnotationStamp returns a new stamp annotation.
func NewPdfAnnotationStamp() *PdfAnnotationStamp {
	_ga := NewPdfAnnotation()
	_efaa := &PdfAnnotationStamp{}
	_efaa.PdfAnnotation = _ga
	_efaa.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ga.SetContext(_efaa)
	return _efaa
}

// NewPdfAnnotationFreeText returns a new free text annotation.
func NewPdfAnnotationFreeText() *PdfAnnotationFreeText {
	_fgec := NewPdfAnnotation()
	_caag := &PdfAnnotationFreeText{}
	_caag.PdfAnnotation = _fgec
	_caag.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fgec.SetContext(_caag)
	return _caag
}

// NewPdfPageResources returns a new PdfPageResources object.
func NewPdfPageResources() *PdfPageResources {
	_fdeea := &PdfPageResources{}
	_fdeea._efbed = _ebb.MakeDict()
	return _fdeea
}
func (_deaf *PdfReader) newPdfFieldSignatureFromDict(_cgffc *_ebb.PdfObjectDictionary) (*PdfFieldSignature, error) {
	_abda := &PdfFieldSignature{}
	_eefge, _cafgg := _ebb.GetIndirect(_cgffc.Get("\u0056"))
	if _cafgg {
		var _egcddb error
		_abda.V, _egcddb = _deaf.newPdfSignatureFromIndirect(_eefge)
		if _egcddb != nil {
			return nil, _egcddb
		}
	}
	_abda.Lock, _ = _ebb.GetIndirect(_cgffc.Get("\u004c\u006f\u0063\u006b"))
	_abda.SV, _ = _ebb.GetIndirect(_cgffc.Get("\u0053\u0056"))
	return _abda, nil
}

// Clear clears flag fl from the flag and returns the resulting flag.
func (_efdb FieldFlag) Clear(fl FieldFlag) FieldFlag { return FieldFlag(_efdb.Mask() &^ fl.Mask()) }

// GetContainingPdfObject returns the container of the outline tree node (indirect object).
func (_ageg *PdfOutlineTreeNode) GetContainingPdfObject() _ebb.PdfObject {
	return _ageg.GetContext().GetContainingPdfObject()
}

// GetNumComponents returns the number of color components (1 for CalGray).
func (_acda *PdfColorCalGray) GetNumComponents() int { return 1 }

// Encoder returns the font's text encoder.
func (_ecgd pdfFontType3) Encoder() _da.TextEncoder { return _ecgd._abcba }

// SetXObjectImageByName adds the provided XObjectImage to the page resources.
// The added XObjectImage is identified by the specified name.
func (_fgcff *PdfPageResources) SetXObjectImageByName(keyName _ebb.PdfObjectName, ximg *XObjectImage) error {
	_eafef := ximg.ToPdfObject().(*_ebb.PdfObjectStream)
	_baefe := _fgcff.SetXObjectByName(keyName, _eafef)
	return _baefe
}

// NewPdfColorCalGray returns a new CalGray color.
func NewPdfColorCalGray(grayVal float64) *PdfColorCalGray {
	_cffee := PdfColorCalGray(grayVal)
	return &_cffee
}
func (_bbef *PdfReader) newPdfActionMovieFromDict(_cee *_ebb.PdfObjectDictionary) (*PdfActionMovie, error) {
	return &PdfActionMovie{Annotation: _cee.Get("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e"), T: _cee.Get("\u0054"), Operation: _cee.Get("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn")}, nil
}

// SetContext set the sub annotation (context).
func (_bcbcb *PdfShading) SetContext(ctx PdfModel) { _bcbcb._edgag = ctx }

// String returns string value of output intent for given type
// ISO_19005-2 6.2.3: GTS_PDFA1 value should be used for PDF/A-1, A-2 and A-3 at least
func (_eeeef PdfOutputIntentType) String() string {
	switch _eeeef {
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
func (_dcdb *DSS) addCRLs(_feffc [][]byte) ([]*_ebb.PdfObjectStream, error) {
	return _dcdb.add(&_dcdb.CRLs, _dcdb._fafgb, _feffc)
}

// HasColorspaceByName checks if the colorspace with the specified name exists in the page resources.
func (_gaffg *PdfPageResources) HasColorspaceByName(keyName _ebb.PdfObjectName) bool {
	_dfffb, _ebgaf := _gaffg.GetColorspaces()
	if _ebgaf != nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _ebgaf)
		return false
	}
	if _dfffb == nil {
		return false
	}
	_, _ebfag := _dfffb.Colorspaces[string(keyName)]
	return _ebfag
}

// ToPdfObject implements interface PdfModel.
func (_ddfb *PdfAnnotationPolyLine) ToPdfObject() _ebb.PdfObject {
	_ddfb.PdfAnnotation.ToPdfObject()
	_bbeb := _ddfb._bdcd
	_aff := _bbeb.PdfObject.(*_ebb.PdfObjectDictionary)
	_ddfb.PdfAnnotationMarkup.appendToPdfDictionary(_aff)
	_aff.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065"))
	_aff.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _ddfb.Vertices)
	_aff.SetIfNotNil("\u004c\u0045", _ddfb.LE)
	_aff.SetIfNotNil("\u0042\u0053", _ddfb.BS)
	_aff.SetIfNotNil("\u0049\u0043", _ddfb.IC)
	_aff.SetIfNotNil("\u0042\u0045", _ddfb.BE)
	_aff.SetIfNotNil("\u0049\u0054", _ddfb.IT)
	_aff.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _ddfb.Measure)
	return _bbeb
}

// Enable LTV enables the specified signature. The signing certificate
// chain is extracted from the signature dictionary. Optionally, additional
// certificates can be specified through the `extraCerts` parameter.
// The LTV client attempts to build the certificate chain up to a trusted root
// by downloading any missing certificates.
func (_gcbca *LTV) Enable(sig *PdfSignature, extraCerts []*_g.Certificate) error {
	if _cdedg := _gcbca.validateSig(sig); _cdedg != nil {
		return _cdedg
	}
	_cabc, _defg := _gcbca.generateVRIKey(sig)
	if _defg != nil {
		return _defg
	}
	if _, _cadec := _gcbca._dfdgf.VRI[_cabc]; _cadec && _gcbca.SkipExisting {
		return nil
	}
	_cdab, _defg := sig.GetCerts()
	if _defg != nil {
		return _defg
	}
	return _gcbca.enable(_cdab, extraCerts, _cabc)
}

// NewImageFromGoImage creates a new NRGBA32 unidoc Image from a golang Image.
// If `goimg` is grayscale (*goimage.Gray8) then calls NewGrayImageFromGoImage instead.
func (_eadd DefaultImageHandler) NewImageFromGoImage(goimg _gdc.Image) (*Image, error) {
	_ebdcf, _cgbf := _dg.FromGoImage(goimg)
	if _cgbf != nil {
		return nil, _cgbf
	}
	_gdaf := _afacb(_ebdcf.Base())
	return &_gdaf, nil
}

// ToUnicode returns the name of the font's "ToUnicode" field if there is one, or "" if there isn't.
func (_eafag *PdfFont) ToUnicode() string {
	if _eafag.baseFields()._dcdd == nil {
		return ""
	}
	return _eafag.baseFields()._dcdd.Name()
}

// PdfFieldButton represents a button field which includes push buttons, checkboxes, and radio buttons.
type PdfFieldButton struct {
	*PdfField
	Opt   *_ebb.PdfObjectArray
	_cdbf *Image
}

// ToPdfOutline returns a low level PdfOutline object, based on the current
// instance.
func (_dcdfe *Outline) ToPdfOutline() *PdfOutline {
	_ecgf := NewPdfOutline()
	var _dagga []*PdfOutlineItem
	var _eggge int64
	var _dbdbc *PdfOutlineItem
	for _, _dbfed := range _dcdfe.Entries {
		_gbfaf, _baab := _dbfed.ToPdfOutlineItem()
		_gbfaf.Parent = &_ecgf.PdfOutlineTreeNode
		if _dbdbc != nil {
			_dbdbc.Next = &_gbfaf.PdfOutlineTreeNode
			_gbfaf.Prev = &_dbdbc.PdfOutlineTreeNode
		}
		_dagga = append(_dagga, _gbfaf)
		_eggge += _baab
		_dbdbc = _gbfaf
	}
	_caff := int64(len(_dagga))
	_eggge += _caff
	if _caff > 0 {
		_ecgf.First = &_dagga[0].PdfOutlineTreeNode
		_ecgf.Last = &_dagga[_caff-1].PdfOutlineTreeNode
		_ecgf.Count = &_eggge
	}
	return _ecgf
}

// ColorToRGB converts a DeviceN color to an RGB color.
func (_afdb *PdfColorspaceDeviceN) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _afdb.AlternateSpace == nil {
		return nil, _gf.New("\u0044\u0065\u0076\u0069\u0063\u0065N\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _afdb.AlternateSpace.ColorToRGB(color)
}

// ToPdfObject implements interface PdfModel.
func (_ded *PdfActionJavaScript) ToPdfObject() _ebb.PdfObject {
	_ded.PdfAction.ToPdfObject()
	_accd := _ded._abe
	_fgb := _accd.PdfObject.(*_ebb.PdfObjectDictionary)
	_fgb.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeJavaScript)))
	_fgb.SetIfNotNil("\u004a\u0053", _ded.JS)
	return _accd
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

// Write writes out the PDF.
func (_bdcab *PdfWriter) Write(writer _ab.Writer) error {
	if err := _bdcab.writeOutlines(); err != nil {
		return err
	}
	if err := _bdcab.writeAcroFormFields(); err != nil {
		return err
	}
	_bdcab.checkPendingObjects()
	if err := _bdcab.writeOutputIntents(); err != nil {
		return err
	}
	_bdcab.setCatalogVersion()
	_bdcab.copyObjects()
	if err := _bdcab.optimize(); err != nil {
		return err
	}
	if err := _bdcab.optimizeDocument(); err != nil {
		return err
	}
	var _bbddfd _c.Hash
	if _bdcab._abgcg {
		_bbddfd = _dd.New()
		writer = _ab.MultiWriter(_bbddfd, writer)
	}
	_bdcab.setWriter(writer)
	_edfb := _bdcab.checkCrossReferenceStream()
	_bcfdb, _edfb := _bdcab.mapObjectStreams(_edfb)
	_bdcab.adjustXRefAffectedVersion(_edfb)
	_bdcab.writeDocumentVersion()
	_bdcab.updateObjectNumbers()
	_bdcab.writeObjects()
	if err := _bdcab.writeObjectsInStreams(_bcfdb); err != nil {
		return err
	}
	_fbdcg := _bdcab._afedd
	var _gabcd int
	for _dfaga := range _bdcab._bedfc {
		if _dfaga > _gabcd {
			_gabcd = _dfaga
		}
	}
	if _bdcab._abgcg {
		if err := _bdcab.setHashIDs(_bbddfd); err != nil {
			return err
		}
	}
	if _edfb {
		if err := _bdcab.writeXRefStreams(_gabcd, _fbdcg); err != nil {
			return err
		}
	} else {
		_bdcab.writeTrailer(_gabcd)
	}
	_bdcab.makeOffSetReference(_fbdcg)
	if err := _bdcab.flushWriter(); err != nil {
		return err
	}
	return nil
}

// PdfOutlineItem represents an outline item dictionary (Table 153 - pp. 376 - 377).
type PdfOutlineItem struct {
	PdfOutlineTreeNode
	Title  *_ebb.PdfObjectString
	Parent *PdfOutlineTreeNode
	Prev   *PdfOutlineTreeNode
	Next   *PdfOutlineTreeNode
	Count  *int64
	Dest   _ebb.PdfObject
	A      _ebb.PdfObject
	SE     _ebb.PdfObject
	C      _ebb.PdfObject
	F      _ebb.PdfObject
	_cacdf *_ebb.PdfIndirectObject
}

// GetPrimitiveFromModel returns the primitive object corresponding to the input `model`.
func (_bfgeg *modelManager) GetPrimitiveFromModel(model PdfModel) _ebb.PdfObject {
	_cbaaf, _cgcfb := _bfgeg._bgfdb[model]
	if !_cgcfb {
		return nil
	}
	return _cbaaf
}

// PdfColorspaceCalRGB stores A, B, C components
type PdfColorspaceCalRGB struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      []float64
	Matrix     []float64
	_cbcd      *_ebb.PdfObjectDictionary
	_aeac      *_ebb.PdfIndirectObject
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_abbbg *PdfShadingPattern) ToPdfObject() _ebb.PdfObject {
	_abbbg.PdfPattern.ToPdfObject()
	_gadg := _abbbg.getDict()
	if _abbbg.Shading != nil {
		_gadg.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _abbbg.Shading.ToPdfObject())
	}
	if _abbbg.Matrix != nil {
		_gadg.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _abbbg.Matrix)
	}
	if _abbbg.ExtGState != nil {
		_gadg.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _abbbg.ExtGState)
	}
	return _abbbg._dcddc
}

// DecodeArray returns the range of color component values in DeviceRGB colorspace.
func (_badc *PdfColorspaceDeviceRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// PdfInfoTrapped specifies pdf trapped information.
type PdfInfoTrapped string

// NewPdfColorspaceSpecialPattern returns a new pattern color.
func NewPdfColorspaceSpecialPattern() *PdfColorspaceSpecialPattern {
	return &PdfColorspaceSpecialPattern{}
}

// GetContext returns a reference to the subshading entry as represented by PdfShadingType1-7.
func (_ccfe *PdfShading) GetContext() PdfModel { return _ccfe._edgag }

// AddAnnotation appends `annot` to the list of page annotations.
func (_fegf *PdfPage) AddAnnotation(annot *PdfAnnotation) {
	if _fegf._bbfed == nil {
		_fegf.GetAnnotations()
	}
	_fegf._bbfed = append(_fegf._bbfed, annot)
}
func _fafaa(_ecdce _ebb.PdfObject) (*PdfFunctionType2, error) {
	_bfge := &PdfFunctionType2{}
	var _cfee *_ebb.PdfObjectDictionary
	if _dcdg, _gbaec := _ecdce.(*_ebb.PdfIndirectObject); _gbaec {
		_fbff, _dcede := _dcdg.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_dcede {
			return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_bfge._bcbfd = _dcdg
		_cfee = _fbff
	} else if _fbeac, _dcaa := _ecdce.(*_ebb.PdfObjectDictionary); _dcaa {
		_cfee = _fbeac
	} else {
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_eg.Log.Trace("\u0046U\u004e\u0043\u0032\u003a\u0020\u0025s", _cfee.String())
	_dafbd, _gdfdb := _ebb.TraceToDirectObject(_cfee.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_ebb.PdfObjectArray)
	if !_gdfdb {
		_eg.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _dafbd.Len() < 0 || _dafbd.Len()%2 != 0 {
		_eg.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u0072\u0061\u006e\u0067e\u0020\u0069\u006e\u0076al\u0069\u0064")
		return nil, _gf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_fecgg, _aagccg := _dafbd.ToFloat64Array()
	if _aagccg != nil {
		return nil, _aagccg
	}
	_bfge.Domain = _fecgg
	_dafbd, _gdfdb = _ebb.TraceToDirectObject(_cfee.Get("\u0052\u0061\u006eg\u0065")).(*_ebb.PdfObjectArray)
	if _gdfdb {
		if _dafbd.Len() < 0 || _dafbd.Len()%2 != 0 {
			return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_cgfeb, _dffd := _dafbd.ToFloat64Array()
		if _dffd != nil {
			return nil, _dffd
		}
		_bfge.Range = _cgfeb
	}
	_dafbd, _gdfdb = _ebb.TraceToDirectObject(_cfee.Get("\u0043\u0030")).(*_ebb.PdfObjectArray)
	if _gdfdb {
		_daddb, _cgcbcc := _dafbd.ToFloat64Array()
		if _cgcbcc != nil {
			return nil, _cgcbcc
		}
		_bfge.C0 = _daddb
	}
	_dafbd, _gdfdb = _ebb.TraceToDirectObject(_cfee.Get("\u0043\u0031")).(*_ebb.PdfObjectArray)
	if _gdfdb {
		_ccdf, _eafcc := _dafbd.ToFloat64Array()
		if _eafcc != nil {
			return nil, _eafcc
		}
		_bfge.C1 = _ccdf
	}
	if len(_bfge.C0) != len(_bfge.C1) {
		_eg.Log.Error("\u0043\u0030\u0020\u0061nd\u0020\u0043\u0031\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069n\u0067")
		return nil, _ebb.ErrRangeError
	}
	N, _aagccg := _ebb.GetNumberAsFloat(_ebb.TraceToDirectObject(_cfee.Get("\u004e")))
	if _aagccg != nil {
		_eg.Log.Error("\u004e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020o\u0072\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u002c\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _cfee.String())
		return nil, _aagccg
	}
	_bfge.N = N
	return _bfge, nil
}

// GetPage returns the PdfPage model for the specified page number.
func (_ddfdcd *PdfReader) GetPage(pageNumber int) (*PdfPage, error) {
	if _ddfdcd._cafdf.GetCrypter() != nil && !_ddfdcd._cafdf.IsAuthenticated() {
		return nil, _bg.Errorf("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	if len(_ddfdcd._faebb) < pageNumber {
		return nil, _gf.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0028\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0074o\u006f\u0020\u0073\u0068\u006f\u0072\u0074\u0029")
	}
	_bfacfb := pageNumber - 1
	if _bfacfb < 0 {
		return nil, _bg.Errorf("\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0069\u006e\u0067\u0020\u006d\u0075\u0073t\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0061\u0074\u0020\u0031")
	}
	_aebfc := _ddfdcd.PageList[_bfacfb]
	return _aebfc, nil
}

// GenerateXObjectName generates an unused XObject name that can be used for
// adding new XObjects. Uses format XObj1, XObj2, ...
func (_aabgc *PdfPageResources) GenerateXObjectName() _ebb.PdfObjectName {
	_dfbge := 1
	for {
		_aeabb := _ebb.MakeName(_bg.Sprintf("\u0058\u004f\u0062\u006a\u0025\u0064", _dfbge))
		if !_aabgc.HasXObjectByName(*_aeabb) {
			return *_aeabb
		}
		_dfbge++
	}
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_afad pdfFontType3) GetCharMetrics(code _da.CharCode) (_bad.CharMetrics, bool) {
	if _beda, _bbege := _afad._aggbb[code]; _bbege {
		return _bad.CharMetrics{Wx: _beda}, true
	}
	if _bad.IsStdFont(_bad.StdFontName(_afad._fdacg)) {
		return _bad.CharMetrics{Wx: 250}, true
	}
	return _bad.CharMetrics{}, false
}
func (_eddfc *PdfAcroForm) fillImageWithAppearance(_bdfcg FieldImageProvider, _fbaab FieldAppearanceGenerator) error {
	if _eddfc == nil {
		return nil
	}
	_eebdf, _gbeea := _bdfcg.FieldImageValues()
	if _gbeea != nil {
		return _gbeea
	}
	for _, _fcfgd := range _eddfc.AllFields() {
		_cbbcc := _fcfgd.PartialName()
		_bdfbc, _dfdab := _eebdf[_cbbcc]
		if !_dfdab {
			if _gbcg, _gacead := _fcfgd.FullName(); _gacead == nil {
				_bdfbc, _dfdab = _eebdf[_gbcg]
			}
		}
		if !_dfdab {
			_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _cbbcc)
			continue
		}
		switch _gfdf := _fcfgd.GetContext().(type) {
		case *PdfFieldButton:
			if _gfdf.IsPush() {
				_gfdf.SetFillImage(_bdfbc)
			}
		}
		if _fbaab == nil {
			continue
		}
		for _, _abdgae := range _fcfgd.Annotations {
			_defc, _dfee := _fbaab.GenerateAppearanceDict(_eddfc, _fcfgd, _abdgae)
			if _dfee != nil {
				return _dfee
			}
			_abdgae.AP = _defc
			_abdgae.ToPdfObject()
		}
	}
	return nil
}

// RunesToCharcodeBytes maps the provided runes to charcode bytes and it
// returns the resulting slice of bytes, along with the number of runes which
// could not be converted. If the number of misses is 0, all runes were
// successfully converted.
func (_dcag *PdfFont) RunesToCharcodeBytes(data []rune) ([]byte, int) {
	var _gcad []_da.TextEncoder
	var _degd _da.CMapEncoder
	if _cebf := _dcag.baseFields()._dcdd; _cebf != nil {
		_degd = _da.NewCMapEncoder("", nil, _cebf)
	}
	_cgfg := _dcag.Encoder()
	if _cgfg != nil {
		switch _bfcg := _cgfg.(type) {
		case _da.SimpleEncoder:
			_adgfe := _bfcg.BaseName()
			if _, _ecdc := _efdc[_adgfe]; _ecdc {
				_gcad = append(_gcad, _cgfg)
			}
		}
	}
	if len(_gcad) == 0 {
		if _dcag.baseFields()._dcdd != nil {
			_gcad = append(_gcad, _degd)
		}
		if _cgfg != nil {
			_gcad = append(_gcad, _cgfg)
		}
	}
	var _fddc _ca.Buffer
	var _bdcdg int
	for _, _dffg := range data {
		var _ceabf bool
		for _, _edgfe := range _gcad {
			if _fade := _edgfe.Encode(string(_dffg)); len(_fade) > 0 {
				_fddc.Write(_fade)
				_ceabf = true
				break
			}
		}
		if !_ceabf {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020f\u0061\u0069\u006ce\u0064\u0020\u0074\u006f \u006d\u0061\u0070\u0020\u0072\u0075\u006e\u0065\u0020\u0060\u0025\u002b\u0071\u0060\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065", _dffg)
			_bdcdg++
		}
	}
	if _bdcdg != 0 {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0061\u006cl\u0020\u0072\u0075\u006e\u0065\u0073\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072c\u006fd\u0065\u0073\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0052\u0075\u006e\u0065\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073 \u0065\u006e\u0063\u006f\u0064e\u0072\u0073\u003d\u0025\u002b\u0076", len(data), _bdcdg, _dcag, _gcad)
	}
	return _fddc.Bytes(), _bdcdg
}

// NewXObjectImageFromStream builds the image xobject from a stream object.
// An image dictionary is the dictionary portion of a stream object representing an image XObject.
func NewXObjectImageFromStream(stream *_ebb.PdfObjectStream) (*XObjectImage, error) {
	_dccdb := &XObjectImage{}
	_dccdb._fbeec = stream
	_fedeb := *(stream.PdfObjectDictionary)
	_fcdd, _cfdf := _ebb.NewEncoderFromStream(stream)
	if _cfdf != nil {
		return nil, _cfdf
	}
	_dccdb.Filter = _fcdd
	if _fadfa := _ebb.TraceToDirectObject(_fedeb.Get("\u0057\u0069\u0064t\u0068")); _fadfa != nil {
		_bdadd, _adcgf := _fadfa.(*_ebb.PdfObjectInteger)
		if !_adcgf {
			return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_acffc := int64(*_bdadd)
		_dccdb.Width = &_acffc
	} else {
		return nil, _gf.New("\u0077\u0069\u0064\u0074\u0068\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _fdfgef := _ebb.TraceToDirectObject(_fedeb.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _fdfgef != nil {
		_aedd, _feegf := _fdfgef.(*_ebb.PdfObjectInteger)
		if !_feegf {
			return nil, _gf.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_dedaf := int64(*_aedd)
		_dccdb.Height = &_dedaf
	} else {
		return nil, _gf.New("\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	if _effge := _ebb.TraceToDirectObject(_fedeb.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _effge != nil {
		_ggfe, _dggdd := NewPdfColorspaceFromPdfObject(_effge)
		if _dggdd != nil {
			return nil, _dggdd
		}
		_dccdb.ColorSpace = _ggfe
	} else {
		_eg.Log.Debug("\u0058O\u0062\u006a\u0065c\u0074\u0020\u0049m\u0061ge\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067 1\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065n\u0074\u0020\u002d\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047r\u0061\u0079")
		_dccdb.ColorSpace = NewPdfColorspaceDeviceGray()
	}
	if _agecg := _ebb.TraceToDirectObject(_fedeb.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _agecg != nil {
		_aebag, _afffb := _agecg.(*_ebb.PdfObjectInteger)
		if !_afffb {
			return nil, _gf.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_agdd := int64(*_aebag)
		_dccdb.BitsPerComponent = &_agdd
	}
	_dccdb.Intent = _fedeb.Get("\u0049\u006e\u0074\u0065\u006e\u0074")
	_dccdb.ImageMask = _fedeb.Get("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k")
	_dccdb.Mask = _fedeb.Get("\u004d\u0061\u0073\u006b")
	_dccdb.Decode = _fedeb.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	_dccdb.Interpolate = _fedeb.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065")
	_dccdb.Alternatives = _fedeb.Get("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073")
	_dccdb.SMask = _fedeb.Get("\u0053\u004d\u0061s\u006b")
	_dccdb.SMaskInData = _fedeb.Get("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061")
	_dccdb.Matte = _fedeb.Get("\u004d\u0061\u0074t\u0065")
	_dccdb.Name = _fedeb.Get("\u004e\u0061\u006d\u0065")
	_dccdb.StructParent = _fedeb.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_dccdb.ID = _fedeb.Get("\u0049\u0044")
	_dccdb.OPI = _fedeb.Get("\u004f\u0050\u0049")
	_dccdb.Metadata = _fedeb.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	_dccdb.OC = _fedeb.Get("\u004f\u0043")
	_dccdb.Stream = stream.Stream
	return _dccdb, nil
}

// AddPages adds pages to be appended to the end of the source PDF.
func (_faddb *PdfAppender) AddPages(pages ...*PdfPage) {
	for _, _fdba := range pages {
		_fdba = _fdba.Duplicate()
		_bfegc(_fdba)
		_faddb._dfbg = append(_faddb._dfbg, _fdba)
	}
}

// Write writes the Appender output to io.Writer.
// It can only be called once and further invocations will result in an error.
func (_bcga *PdfAppender) Write(w _ab.Writer) error {
	if _bcga._adca {
		return _gf.New("\u0061\u0070\u0070\u0065\u006e\u0064\u0065\u0072\u0020\u0077\u0072\u0069\u0074e\u0020\u0063\u0061\u006e\u0020\u006fn\u006c\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0076\u006f\u006b\u0065\u0064 \u006f\u006e\u0063\u0065")
	}
	_edbe := NewPdfWriter()
	_dgag, _fdbd := _ebb.GetDict(_edbe._dggbf)
	if !_fdbd {
		return _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_ecad, _fdbd := _dgag.Get("\u004b\u0069\u0064\u0073").(*_ebb.PdfObjectArray)
	if !_fdbd {
		return _gf.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_eace, _fdbd := _dgag.Get("\u0043\u006f\u0075n\u0074").(*_ebb.PdfObjectInteger)
	if !_fdbd {
		return _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	_eddff := _bcga._acfe._cafdf
	_agga := _eddff.GetTrailer()
	if _agga == nil {
		return _gf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_dafg, _fdbd := _ebb.GetIndirect(_agga.Get("\u0052\u006f\u006f\u0074"))
	if !_fdbd {
		return _gf.New("c\u0061\u0074\u0061\u006c\u006f\u0067 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064")
	}
	_dfbc, _fdbd := _ebb.GetDict(_dafg)
	if !_fdbd {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _dafg, *_agga)
		return _gf.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	for _, _bbaf := range _dfbc.Keys() {
		if _edbe._dffegd.Get(_bbaf) == nil {
			_dbacb := _dfbc.Get(_bbaf)
			_edbe._dffegd.Set(_bbaf, _dbacb)
		}
	}
	if _bcga._bfef != nil {
		_edbe._dffegd.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _bcga._bfef.ToPdfObject())
		_bcga.updateObjectsDeep(_bcga._bfef.ToPdfObject(), nil)
	}
	if _bcga._eged != nil {
		_bcga.updateObjectsDeep(_bcga._eged.ToPdfObject(), nil)
		_edbe._dffegd.Set("\u0044\u0053\u0053", _bcga._eged.GetContainingPdfObject())
	}
	if _bcga._eaaa != nil {
		_edbe._dffegd.Set("\u0050\u0065\u0072m\u0073", _bcga._eaaa.ToPdfObject())
		_bcga.updateObjectsDeep(_bcga._eaaa.ToPdfObject(), nil)
	}
	if _edbe._efcge.Major < 2 {
		_edbe.AddExtension("\u0045\u0053\u0049\u0043", "\u0031\u002e\u0037", 5)
		_edbe.AddExtension("\u0041\u0044\u0042\u0045", "\u0031\u002e\u0037", 8)
	}
	if _cefa, _caga := _ebb.GetDict(_agga.Get("\u0049\u006e\u0066\u006f")); _caga {
		if _gdfe, _fgcf := _ebb.GetDict(_edbe._eadfd); _fgcf {
			for _, _cdfe := range _cefa.Keys() {
				if _gdfe.Get(_cdfe) == nil {
					_gdfe.Set(_cdfe, _cefa.Get(_cdfe))
				}
			}
		}
	}
	if _bcga._eeee != nil {
		_edbe._eadfd = _ebb.MakeIndirectObject(_bcga._eeee.ToPdfObject())
	}
	_bcga.addNewObject(_edbe._eadfd)
	_bcga.addNewObject(_edbe._gegba)
	_dggdg := false
	if len(_bcga._acfe.PageList) != len(_bcga._dfbg) {
		_dggdg = true
	} else {
		for _fdaa := range _bcga._acfe.PageList {
			switch {
			case _bcga._dfbg[_fdaa] == _bcga._acfe.PageList[_fdaa]:
			case _bcga._dfbg[_fdaa] == _bcga.Reader.PageList[_fdaa]:
			default:
				_dggdg = true
			}
			if _dggdg {
				break
			}
		}
	}
	if _dggdg {
		_bcga.updateObjectsDeep(_edbe._dggbf, nil)
	} else {
		_bcga._eebc[_edbe._dggbf] = struct{}{}
	}
	_edbe._dggbf.ObjectNumber = _bcga.Reader._eedbb.ObjectNumber
	_bcga._gbfa[_edbe._dggbf] = _bcga.Reader._eedbb.ObjectNumber
	_bcgfc := []_ebb.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	for _, _cefg := range _bcga._dfbg {
		_efga := _cefg.ToPdfObject()
		*_eace = *_eace + 1
		if _aacc, _fgfe := _efga.(*_ebb.PdfIndirectObject); _fgfe && _aacc.GetParser() == _bcga._acfe._cafdf {
			_ecad.Append(&_aacc.PdfObjectReference)
			continue
		}
		if _afac, _fafa := _ebb.GetDict(_efga); _fafa {
			_fdfa, _cegbd := _afac.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_ebb.PdfIndirectObject)
			for _cegbd {
				_eg.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _fdfa)
				_ebaf, _bcbc := _fdfa.PdfObject.(*_ebb.PdfObjectDictionary)
				if !_bcbc {
					return _gf.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
				}
				for _, _adfg := range _bcgfc {
					_eg.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _adfg)
					if _afac.Get(_adfg) != nil {
						_eg.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
						continue
					}
					if _ebee := _ebaf.Get(_adfg); _ebee != nil {
						_eg.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _adfg)
						_afac.Set(_adfg, _ebee)
					}
				}
				_fdfa, _cegbd = _ebaf.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_ebb.PdfIndirectObject)
				_eg.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _ebaf.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
			}
			_afac.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _edbe._dggbf)
		}
		_bcga.updateObjectsDeep(_efga, nil)
		_ecad.Append(_efga)
	}
	if _, _gdcg := _bcga._ecce.Seek(0, _ab.SeekStart); _gdcg != nil {
		return _gdcg
	}
	_dbf := make(map[SignatureHandler]_ab.Writer)
	_bcde := _ebb.MakeArray()
	for _, _daaa := range _bcga._bfeg {
		if _cecec, _feea := _ebb.GetIndirect(_daaa); _feea {
			if _gfce, _bfaf := _cecec.PdfObject.(*pdfSignDictionary); _bfaf {
				_deaa := *_gfce._dcfab
				var _adcfd error
				_dbf[_deaa], _adcfd = _deaa.NewDigest(_gfce._bead)
				if _adcfd != nil {
					return _adcfd
				}
				_bcde.Append(_ebb.MakeInteger(0xfffff), _ebb.MakeInteger(0xfffff))
			}
		}
	}
	if _bcde.Len() > 0 {
		_bcde.Append(_ebb.MakeInteger(0xfffff), _ebb.MakeInteger(0xfffff))
	}
	for _, _ecb := range _bcga._bfeg {
		if _bgda, _cgdg := _ebb.GetIndirect(_ecb); _cgdg {
			if _gfae, _acccc := _bgda.PdfObject.(*pdfSignDictionary); _acccc {
				_gfae.Set("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _bcde)
			}
		}
	}
	_gdfd := len(_dbf) > 0
	var _aec _ab.Reader = _bcga._ecce
	if _gdfd {
		_dggf := make([]_ab.Writer, 0, len(_dbf))
		for _, _cade := range _dbf {
			_dggf = append(_dggf, _cade)
		}
		_aec = _ab.TeeReader(_bcga._ecce, _ab.MultiWriter(_dggf...))
	}
	_defa, _aebe := _ab.Copy(w, _aec)
	if _aebe != nil {
		return _aebe
	}
	if len(_bcga._bfeg) == 0 {
		return nil
	}
	_edbe._fgdce = _defa
	_edbe.ObjNumOffset = _bcga._gbddb
	_edbe._abffb = true
	_edbe._efdega = _bcga._acfd
	_edbe._bcage = _bcga._cfag
	_edbe._ggbfg = _bcga._bee
	_edbe._efcge = _bcga._acfe.PdfVersion()
	_edbe._cdgd = _bcga._gbfa
	_edbe._cgfde = _bcga._gege.GetCrypter()
	_edbe._cbcaa = _bcga._gege.GetEncryptObj()
	_eade := _bcga._gege.GetXrefType()
	if _eade != nil {
		_gaeg := *_eade == _ebb.XrefTypeObjectStream
		_edbe._adeff = &_gaeg
	}
	_edbe._ffffd = map[_ebb.PdfObject]struct{}{}
	_edbe._ebdgg = []_ebb.PdfObject{}
	for _, _bgbc := range _bcga._bfeg {
		if _, _bafd := _bcga._eebc[_bgbc]; _bafd {
			continue
		}
		_edbe.addObject(_bgbc)
	}
	_fgada := w
	if _gdfd {
		_fgada = _ca.NewBuffer(nil)
	}
	if _bcga._accg != "" && _edbe._cgfde == nil {
		_edbe.Encrypt([]byte(_bcga._accg), []byte(_bcga._accg), _bcga._gfba)
	}
	if _eecd := _agga.Get("\u0049\u0044"); _eecd != nil {
		if _fadf, _egab := _ebb.GetArray(_eecd); _egab {
			_edbe._eecfe = _fadf
		}
	}
	if _daab := _edbe.Write(_fgada); _daab != nil {
		return _daab
	}
	if _gdfd {
		_gfag := _fgada.(*_ca.Buffer).Bytes()
		_agbb := _ebb.MakeArray()
		var _fdfca []*pdfSignDictionary
		var _adad int64
		for _, _facb := range _edbe._ebdgg {
			if _gaecc, _bfae := _ebb.GetIndirect(_facb); _bfae {
				if _bdbf, _bcfc := _gaecc.PdfObject.(*pdfSignDictionary); _bcfc {
					_fdfca = append(_fdfca, _bdbf)
					_aeca := _bdbf._cbfg + int64(_bdbf._eefbe)
					_agbb.Append(_ebb.MakeInteger(_adad), _ebb.MakeInteger(_aeca-_adad))
					_adad = _bdbf._cbfg + int64(_bdbf._adggf)
				}
			}
		}
		_agbb.Append(_ebb.MakeInteger(_adad), _ebb.MakeInteger(_defa+int64(len(_gfag))-_adad))
		_aeda := []byte(_agbb.WriteString())
		for _, _gde := range _fdfca {
			_dgea := int(_gde._cbfg - _defa)
			for _eeef := _gde._decgd; _eeef < _gde._ddceg; _eeef++ {
				_gfag[_dgea+_eeef] = ' '
			}
			_ecab := _gfag[_dgea+_gde._decgd : _dgea+_gde._ddceg]
			copy(_ecab, _aeda)
		}
		var _bafdb int
		for _, _abdc := range _fdfca {
			_dbce := int(_abdc._cbfg - _defa)
			_dacd := _gfag[_bafdb : _dbce+_abdc._eefbe]
			_cdfef := *_abdc._dcfab
			_dbf[_cdfef].Write(_dacd)
			_bafdb = _dbce + _abdc._adggf
		}
		for _, _caed := range _fdfca {
			_bfba := _gfag[_bafdb:]
			_adga := *_caed._dcfab
			_dbf[_adga].Write(_bfba)
		}
		for _, _gggea := range _fdfca {
			_cggd := int(_gggea._cbfg - _defa)
			_egef := *_gggea._dcfab
			_gcff := _dbf[_egef]
			if _dfda := _egef.Sign(_gggea._bead, _gcff); _dfda != nil {
				return _dfda
			}
			_gggea._bead.ByteRange = _agbb
			_cfce := []byte(_gggea._bead.Contents.WriteString())
			for _dgfg := _gggea._decgd; _dgfg < _gggea._ddceg; _dgfg++ {
				_gfag[_cggd+_dgfg] = ' '
			}
			for _defb := _gggea._eefbe; _defb < _gggea._adggf; _defb++ {
				_gfag[_cggd+_defb] = ' '
			}
			_ffc := _gfag[_cggd+_gggea._decgd : _cggd+_gggea._ddceg]
			copy(_ffc, _aeda)
			_ffc = _gfag[_cggd+_gggea._eefbe : _cggd+_gggea._adggf]
			copy(_ffc, _cfce)
		}
		_gebg := _ca.NewBuffer(_gfag)
		_, _aebe = _ab.Copy(w, _gebg)
		if _aebe != nil {
			return _aebe
		}
	}
	_bcga._adca = true
	return nil
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_decaf *PdfShadingType7) ToPdfObject() _ebb.PdfObject {
	_decaf.PdfShading.ToPdfObject()
	_dadgc, _aaff := _decaf.getShadingDict()
	if _aaff != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _decaf.BitsPerCoordinate != nil {
		_dadgc.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _decaf.BitsPerCoordinate)
	}
	if _decaf.BitsPerComponent != nil {
		_dadgc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _decaf.BitsPerComponent)
	}
	if _decaf.BitsPerFlag != nil {
		_dadgc.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _decaf.BitsPerFlag)
	}
	if _decaf.Decode != nil {
		_dadgc.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _decaf.Decode)
	}
	if _decaf.Function != nil {
		if len(_decaf.Function) == 1 {
			_dadgc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _decaf.Function[0].ToPdfObject())
		} else {
			_gbacge := _ebb.MakeArray()
			for _, _gcbg := range _decaf.Function {
				_gbacge.Append(_gcbg.ToPdfObject())
			}
			_dadgc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _gbacge)
		}
	}
	return _decaf._fbfae
}

// NewPdfAnnotationSquare returns a new square annotation.
func NewPdfAnnotationSquare() *PdfAnnotationSquare {
	_aaca := NewPdfAnnotation()
	_bge := &PdfAnnotationSquare{}
	_bge.PdfAnnotation = _aaca
	_bge.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_aaca.SetContext(_bge)
	return _bge
}

// GetColorspaces loads PdfPageResourcesColorspaces from `r.ColorSpace` and returns an error if there
// is a problem loading. Once loaded, the same object is returned on multiple calls.
func (_eaebe *PdfPageResources) GetColorspaces() (*PdfPageResourcesColorspaces, error) {
	if _eaebe._aaee != nil {
		return _eaebe._aaee, nil
	}
	if _eaebe.ColorSpace == nil {
		return nil, nil
	}
	_gfcda, _dgaac := _eeafg(_eaebe.ColorSpace)
	if _dgaac != nil {
		return nil, _dgaac
	}
	_eaebe._aaee = _gfcda
	return _eaebe._aaee, nil
}

// PdfFont represents an underlying font structure which can be of type:
// - Type0
// - Type1
// - TrueType
// etc.
type PdfFont struct{ _ebcad pdfFont }
type pdfSignDictionary struct {
	*_ebb.PdfObjectDictionary
	_dcfab *SignatureHandler
	_bead  *PdfSignature
	_cbfg  int64
	_eefbe int
	_adggf int
	_decgd int
	_ddceg int
}

// SetEncoder sets the encoding for the underlying font.
// TODO(peterwilliams97): Change function signature to SetEncoder(encoder *textencoding.simpleEncoder).
// TODO(gunnsth): Makes sense if SetEncoder is removed from the interface fonts.Font as proposed in PR #260.
func (_abgb *pdfFontSimple) SetEncoder(encoder _da.TextEncoder) { _abgb._ebcb = encoder }

// Width returns the width of `rect`.
func (_cfdb *PdfRectangle) Width() float64 { return _cbg.Abs(_cfdb.Urx - _cfdb.Llx) }
func _acdf(_feeb _ebb.PdfObject) (*PdfFontDescriptor, error) {
	_efbab := &PdfFontDescriptor{}
	_feeb = _ebb.ResolveReference(_feeb)
	if _fbac, _agef := _feeb.(*_ebb.PdfIndirectObject); _agef {
		_efbab._ccfcg = _fbac
		_feeb = _fbac.PdfObject
	}
	_addaa, _dgfgf := _ebb.GetDict(_feeb)
	if !_dgfgf {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0044\u0065\u0073c\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _feeb)
		return nil, _ebb.ErrTypeError
	}
	if _ffadg := _addaa.Get("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"); _ffadg != nil {
		_efbab.FontName = _ffadg
	} else {
		_eg.Log.Debug("\u0049n\u0063\u006fm\u0070\u0061\u0074\u0069b\u0069\u006c\u0069t\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u004e\u0061me\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006di\u0073\u0073i\u006e\u0067")
	}
	_gbcfc, _ := _ebb.GetName(_efbab.FontName)
	if _dcdff := _addaa.Get("\u0054\u0079\u0070\u0065"); _dcdff != nil {
		_badff, _gagfg := _dcdff.(*_ebb.PdfObjectName)
		if !_gagfg || string(*_badff) != "\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072" {
			_eg.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072i\u0070t\u006f\u0072\u0020\u0054y\u0070\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0054\u0029\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _dcdff, _gbcfc, _efbab.FontName)
		}
	} else {
		_eg.Log.Trace("\u0049\u006ec\u006f\u006d\u0070\u0061\u0074i\u0062\u0069\u006c\u0069\u0074y\u003a\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _gbcfc, _efbab.FontName)
	}
	_efbab.FontFamily = _addaa.Get("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079")
	_efbab.FontStretch = _addaa.Get("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068")
	_efbab.FontWeight = _addaa.Get("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074")
	_efbab.Flags = _addaa.Get("\u0046\u006c\u0061g\u0073")
	_efbab.FontBBox = _addaa.Get("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078")
	_efbab.ItalicAngle = _addaa.Get("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065")
	_efbab.Ascent = _addaa.Get("\u0041\u0073\u0063\u0065\u006e\u0074")
	_efbab.Descent = _addaa.Get("\u0044e\u0073\u0063\u0065\u006e\u0074")
	_efbab.Leading = _addaa.Get("\u004ce\u0061\u0064\u0069\u006e\u0067")
	_efbab.CapHeight = _addaa.Get("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t")
	_efbab.XHeight = _addaa.Get("\u0058H\u0065\u0069\u0067\u0068\u0074")
	_efbab.StemV = _addaa.Get("\u0053\u0074\u0065m\u0056")
	_efbab.StemH = _addaa.Get("\u0053\u0074\u0065m\u0048")
	_efbab.AvgWidth = _addaa.Get("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068")
	_efbab.MaxWidth = _addaa.Get("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068")
	_efbab.MissingWidth = _addaa.Get("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068")
	_efbab.FontFile = _addaa.Get("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065")
	_efbab.FontFile2 = _addaa.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652")
	_efbab.FontFile3 = _addaa.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653")
	_efbab.CharSet = _addaa.Get("\u0043h\u0061\u0072\u0053\u0065\u0074")
	_efbab.Style = _addaa.Get("\u0053\u0074\u0079l\u0065")
	_efbab.Lang = _addaa.Get("\u004c\u0061\u006e\u0067")
	_efbab.FD = _addaa.Get("\u0046\u0044")
	_efbab.CIDSet = _addaa.Get("\u0043\u0049\u0044\u0053\u0065\u0074")
	if _efbab.Flags != nil {
		if _efbb, _fbgf := _ebb.GetIntVal(_efbab.Flags); _fbgf {
			_efbab._gfbge = _efbb
		}
	}
	if _efbab.MissingWidth != nil {
		if _egcg, _gfage := _ebb.GetNumberAsFloat(_efbab.MissingWidth); _gfage == nil {
			_efbab._gbfgb = _egcg
		}
	}
	if _efbab.FontFile != nil {
		_dcbga, _bdaa := _aafe(_efbab.FontFile)
		if _bdaa != nil {
			return _efbab, _bdaa
		}
		_eg.Log.Trace("f\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u003d\u0025\u0073", _dcbga)
		_efbab.fontFile = _dcbga
	}
	if _efbab.FontFile2 != nil {
		_cgegf, _edbb := _bad.NewFontFile2FromPdfObject(_efbab.FontFile2)
		if _edbb != nil {
			return _efbab, _edbb
		}
		_eg.Log.Trace("\u0066\u006f\u006et\u0046\u0069\u006c\u0065\u0032\u003d\u0025\u0073", _cgegf.String())
		_efbab._aeeb = &_cgegf
	}
	return _efbab, nil
}
func (_ecegg *PdfWriter) optimizeDocument() error {
	if _ecegg._cafac == nil {
		return nil
	}
	_ebbac, _faebg := _ebb.GetDict(_ecegg._eadfd)
	if !_faebg {
		return _gf.New("\u0061\u006e\u0020in\u0066\u006f\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_beeab := _bda.Document{ID: [2]string{_ecegg._gfdea, _ecegg._gffb}, Version: _ecegg._efcge, Objects: _ecegg._ebdgg, Info: _ebbac, Crypt: _ecegg._cgfde, UseHashBasedID: _ecegg._abgcg}
	if _fefeca := _ecegg._cafac.ApplyStandard(&_beeab); _fefeca != nil {
		return _fefeca
	}
	_ecegg._gfdea, _ecegg._gffb = _beeab.ID[0], _beeab.ID[1]
	_ecegg._efcge = _beeab.Version
	_ecegg._ebdgg = _beeab.Objects
	_ecegg._eadfd.PdfObject = _beeab.Info
	_ecegg._abgcg = _beeab.UseHashBasedID
	_ecegg._cgfde = _beeab.Crypt
	_cggfc := make(map[_ebb.PdfObject]struct{}, len(_ecegg._ebdgg))
	for _, _eegd := range _ecegg._ebdgg {
		_cggfc[_eegd] = struct{}{}
	}
	_ecegg._ffffd = _cggfc
	return nil
}

// Flags returns the field flags for the field accounting for any inherited flags.
func (_ageea *PdfField) Flags() FieldFlag {
	var _dddb FieldFlag
	_bfdd, _bcba := _ageea.inherit(func(_fdac *PdfField) bool {
		if _fdac.Ff != nil {
			_dddb = FieldFlag(*_fdac.Ff)
			return true
		}
		return false
	})
	if _bcba != nil {
		_eg.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0069\u006e\u0067\u0020\u0066\u006c\u0061\u0067\u0073\u0020\u0076\u0069\u0061\u0020\u0069\u006e\u0068\u0065\u0072\u0069t\u0061\u006e\u0063\u0065\u003a\u0020\u0025\u0076", _bcba)
	}
	if !_bfdd {
		_eg.Log.Trace("N\u006f\u0020\u0066\u0069\u0065\u006cd\u0020\u0066\u006c\u0061\u0067\u0073 \u0066\u006f\u0075\u006e\u0064\u0020\u002d \u0061\u0073\u0073\u0075\u006d\u0065\u0020\u0063\u006c\u0065a\u0072")
	}
	return _dddb
}

// PdfColorspaceDeviceRGB represents an RGB colorspace.
type PdfColorspaceDeviceRGB struct{}

// AddContentStreamByString adds content stream by string. Puts the content
// string into a stream object and points the content stream towards it.
func (_bfgec *PdfPage) AddContentStreamByString(contentStr string) error {
	_ebabe, _acad := _ebb.MakeStream([]byte(contentStr), _ebb.NewFlateEncoder())
	if _acad != nil {
		return _acad
	}
	if _bfgec.Contents == nil {
		_bfgec.Contents = _ebabe
	} else {
		_bgdfc := _ebb.TraceToDirectObject(_bfgec.Contents)
		_fbfg, _agfcbf := _bgdfc.(*_ebb.PdfObjectArray)
		if !_agfcbf {
			_fbfg = _ebb.MakeArray(_bgdfc)
		}
		_fbfg.Append(_ebabe)
		_bfgec.Contents = _fbfg
	}
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_afa *PdfActionGoTo3DView) ToPdfObject() _ebb.PdfObject {
	_afa.PdfAction.ToPdfObject()
	_cge := _afa._abe
	_cba := _cge.PdfObject.(*_ebb.PdfObjectDictionary)
	_cba.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeGoTo3DView)))
	_cba.SetIfNotNil("\u0054\u0041", _afa.TA)
	_cba.SetIfNotNil("\u0056", _afa.V)
	return _cge
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_gacc *PdfColorspaceCalGray) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_aaadd, _cbab := _ebb.GetNumbersAsFloat(objects)
	if _cbab != nil {
		return nil, _cbab
	}
	return _gacc.ColorFromFloats(_aaadd)
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
	_bcbfd *_ebb.PdfIndirectObject
}

// ToPdfObject implements interface PdfModel.
func (_agc *PdfAnnotationProjection) ToPdfObject() _ebb.PdfObject {
	_agc.PdfAnnotation.ToPdfObject()
	_abga := _agc._bdcd
	_dafb := _abga.PdfObject.(*_ebb.PdfObjectDictionary)
	_agc.PdfAnnotationMarkup.appendToPdfDictionary(_dafb)
	return _abga
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_dggdc pdfFontType0) GetCharMetrics(code _da.CharCode) (_bad.CharMetrics, bool) {
	if _dggdc.DescendantFont == nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _dggdc)
		return _bad.CharMetrics{}, false
	}
	return _dggdc.DescendantFont.GetCharMetrics(code)
}
func _ffaba(_bbffd *_ebb.PdfObjectArray) (float64, error) {
	_dgffb, _ddcdd := _bbffd.ToFloat64Array()
	if _ddcdd != nil {
		_eg.Log.Debug("\u0042\u0061\u0064\u0020\u004d\u0061\u0074\u0074\u0065\u0020\u0061\u0072\u0072\u0061\u0079:\u0020m\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bbffd, _ddcdd)
	}
	switch len(_dgffb) {
	case 1:
		return _dgffb[0], nil
	case 3:
		_fbbgc := PdfColorspaceDeviceRGB{}
		_cfcaf, _bfgff := _fbbgc.ColorFromFloats(_dgffb)
		if _bfgff != nil {
			return 0.0, _bfgff
		}
		return _cfcaf.(*PdfColorDeviceRGB).ToGray().Val(), nil
	case 4:
		_faddeb := PdfColorspaceDeviceCMYK{}
		_dcebf, _dcacg := _faddeb.ColorFromFloats(_dgffb)
		if _dcacg != nil {
			return 0.0, _dcacg
		}
		_ebcdf, _dcacg := _faddeb.ColorToRGB(_dcebf.(*PdfColorDeviceCMYK))
		if _dcacg != nil {
			return 0.0, _dcacg
		}
		return _ebcdf.(*PdfColorDeviceRGB).ToGray().Val(), nil
	}
	_ddcdd = _gf.New("\u0062a\u0064 \u004d\u0061\u0074\u0074\u0065\u0020\u0063\u006f\u006c\u006f\u0072")
	_eg.Log.Error("\u0074\u006f\u0047ra\u0079\u003a\u0020\u006d\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bbffd, _ddcdd)
	return 0.0, _ddcdd
}

// GetPageLabels returns the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_degcf *PdfReader) GetPageLabels() (_ebb.PdfObject, error) {
	_cfebd := _ebb.ResolveReference(_degcf._fdgda.Get("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073"))
	if _cfebd == nil {
		return nil, nil
	}
	if !_degcf._ceefa {
		_afabg := _degcf.traverseObjectData(_cfebd)
		if _afabg != nil {
			return nil, _afabg
		}
	}
	return _cfebd, nil
}

// PdfActionURI represents an URI action.
type PdfActionURI struct {
	*PdfAction
	URI   _ebb.PdfObject
	IsMap _ebb.PdfObject
}

func (_ddgbg *PdfReader) buildOutlineTree(_eead _ebb.PdfObject, _fabf *PdfOutlineTreeNode, _cfgdbg *PdfOutlineTreeNode, _cdbcf map[_ebb.PdfObject]struct{}) (*PdfOutlineTreeNode, *PdfOutlineTreeNode, error) {
	if _cdbcf == nil {
		_cdbcf = map[_ebb.PdfObject]struct{}{}
	}
	_cdbcf[_eead] = struct{}{}
	_dfeeg, _bdbaef := _eead.(*_ebb.PdfIndirectObject)
	if !_bdbaef {
		return nil, nil, _bg.Errorf("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0063\u006f\u006et\u0061\u0069\u006e\u0065r\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _eead)
	}
	_fdggb, _aeabca := _dfeeg.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_aeabca {
		return nil, nil, _gf.New("\u006e\u006f\u0074 a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_eg.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0074\u0072\u0065\u0065\u003a\u0020d\u0069\u0063\u0074\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0076\u0029\u0020p\u003a\u0020\u0025\u0070", _fdggb, _dfeeg, _dfeeg)
	if _cgeea := _fdggb.Get("\u0054\u0069\u0074l\u0065"); _cgeea != nil {
		_fgfdb, _bdggg := _ddgbg.newPdfOutlineItemFromIndirectObject(_dfeeg)
		if _bdggg != nil {
			return nil, nil, _bdggg
		}
		_fgfdb.Parent = _fabf
		_fgfdb.Prev = _cfgdbg
		_febde := _ebb.ResolveReference(_fdggb.Get("\u0046\u0069\u0072s\u0074"))
		if _, _fcgbbg := _cdbcf[_febde]; _febde != nil && _febde != _dfeeg && !_fcgbbg {
			if !_ebb.IsNullObject(_febde) {
				_gcdga, _ggdc, _cfdce := _ddgbg.buildOutlineTree(_febde, &_fgfdb.PdfOutlineTreeNode, nil, _cdbcf)
				if _cfdce != nil {
					_eg.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0069\u0074\u0065\u006d\u0020\u0074\u0072\u0065\u0065\u003a \u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020n\u006f\u0064\u0065\u0020\u0063\u0068\u0069\u006c\u0064\u0072\u0065n\u002e", _cfdce)
				} else {
					_fgfdb.First = _gcdga
					_fgfdb.Last = _ggdc
				}
			}
		}
		_ggab := _ebb.ResolveReference(_fdggb.Get("\u004e\u0065\u0078\u0074"))
		if _, _fabd := _cdbcf[_ggab]; _ggab != nil && _ggab != _dfeeg && !_fabd {
			if !_ebb.IsNullObject(_ggab) {
				_bfaef, _daed, _dcef := _ddgbg.buildOutlineTree(_ggab, _fabf, &_fgfdb.PdfOutlineTreeNode, _cdbcf)
				if _dcef != nil {
					_eg.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065\u0020\u0066\u006f\u0072\u0020\u004ee\u0078\u0074\u0020\u006e\u006f\u0064\u0065\u003a\u0020\u0025\u0076\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006e\u006f\u0064e\u002e", _dcef)
				} else {
					_fgfdb.Next = _bfaef
					return &_fgfdb.PdfOutlineTreeNode, _daed, nil
				}
			}
		}
		return &_fgfdb.PdfOutlineTreeNode, &_fgfdb.PdfOutlineTreeNode, nil
	}
	_cdbda, _edfga := _fgff(_dfeeg)
	if _edfga != nil {
		return nil, nil, _edfga
	}
	_cdbda.Parent = _fabf
	if _cdfgea := _fdggb.Get("\u0046\u0069\u0072s\u0074"); _cdfgea != nil {
		_cdfgea = _ebb.ResolveReference(_cdfgea)
		if _, _gbegd := _cdbcf[_cdfgea]; _cdfgea != nil && _cdfgea != _dfeeg && !_gbegd {
			_cbbfe := _ebb.TraceToDirectObject(_cdfgea)
			if _, _gfcef := _cbbfe.(*_ebb.PdfObjectNull); !_gfcef && _cbbfe != nil {
				_eaddd, _bcaga, _bgae := _ddgbg.buildOutlineTree(_cdfgea, &_cdbda.PdfOutlineTreeNode, nil, _cdbcf)
				if _bgae != nil {
					_eg.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020b\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0074\u0072\u0065\u0065\u003a\u0020\u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006e\u006f\u0064\u0065 \u0063\u0068i\u006c\u0064r\u0065n\u002e", _bgae)
				} else {
					_cdbda.First = _eaddd
					_cdbda.Last = _bcaga
				}
			}
		}
	}
	return &_cdbda.PdfOutlineTreeNode, &_cdbda.PdfOutlineTreeNode, nil
}

// ToPdfObject implements interface PdfModel.
func (_abgg *PdfAnnotationUnderline) ToPdfObject() _ebb.PdfObject {
	_abgg.PdfAnnotation.ToPdfObject()
	_afd := _abgg._bdcd
	_ebada := _afd.PdfObject.(*_ebb.PdfObjectDictionary)
	_abgg.PdfAnnotationMarkup.appendToPdfDictionary(_ebada)
	_ebada.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee"))
	_ebada.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _abgg.QuadPoints)
	return _afd
}
func _ffaa(_eafcb _ebb.PdfObject) []*_ebb.PdfObjectStream {
	if _eafcb == nil {
		return nil
	}
	_acefg, _dgbcc := _ebb.GetArray(_eafcb)
	if !_dgbcc || _acefg.Len() == 0 {
		return nil
	}
	_fddae := make([]*_ebb.PdfObjectStream, 0, _acefg.Len())
	for _, _cgbge := range _acefg.Elements() {
		if _bggd, _fdgc := _ebb.GetStream(_cgbge); _fdgc {
			_fddae = append(_fddae, _bggd)
		}
	}
	return _fddae
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_cfaag *PdfColorspaceCalGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_afbc := vals[0]
	if _afbc < 0.0 || _afbc > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _afbc)
		return nil, ErrColorOutOfRange
	}
	_effd := NewPdfColorCalGray(_afbc)
	return _effd, nil
}
func (_gfbd *PdfColorspaceLab) String() string { return "\u004c\u0061\u0062" }

// BorderEffect represents a border effect (Table 167 p. 395).
type BorderEffect int

func (_aggg *PdfAcroForm) signatureFields() []*PdfFieldSignature {
	var _ecedb []*PdfFieldSignature
	for _, _gafd := range _aggg.AllFields() {
		switch _eeegd := _gafd.GetContext().(type) {
		case *PdfFieldSignature:
			_cafgdg := _eeegd
			_ecedb = append(_ecedb, _cafgdg)
		}
	}
	return _ecedb
}

// ToInteger convert to an integer format.
func (_ccfb *PdfColorLab) ToInteger(bits int) [3]uint32 {
	_eaf := _cbg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_eaf * _ccfb.L()), uint32(_eaf * _ccfb.A()), uint32(_eaf * _ccfb.B())}
}
func _cfacd(_bceda StdFontName) (pdfFontSimple, error) {
	_bebc, _bafa := _bad.NewStdFontByName(_bceda)
	if !_bafa {
		return pdfFontSimple{}, ErrFontNotSupported
	}
	_begc := _fcgbc(_bebc)
	return _begc, nil
}

// PartialName returns the partial name of the field.
func (_adfed *PdfField) PartialName() string {
	_fbbgb := ""
	if _adfed.T != nil {
		_fbbgb = _adfed.T.Decoded()
	} else {
		_eg.Log.Debug("\u0046\u0069el\u0064\u0020\u006di\u0073\u0073\u0069\u006eg T\u0020fi\u0065\u006c\u0064\u0020\u0028\u0069\u006eco\u006d\u0070\u0061\u0074\u0069\u0062\u006ce\u0029")
	}
	return _fbbgb
}

// GetAscent returns the Ascent of the font `descriptor`.
func (_dddc *PdfFontDescriptor) GetAscent() (float64, error) {
	return _ebb.GetNumberAsFloat(_dddc.Ascent)
}
func _bfbb(_gfab []byte) bool {
	if len(_gfab) < 4 {
		return true
	}
	for _afgc := range _gfab[:4] {
		_gaaa := rune(_afgc)
		if !_cc.Is(_cc.ASCII_Hex_Digit, _gaaa) && !_cc.IsSpace(_gaaa) {
			return true
		}
	}
	return false
}

// ToGoTime returns the date in time.Time format.
func (_ceecg PdfDate) ToGoTime() _f.Time {
	_babde := int(_ceecg._gcccb*60*60 + _ceecg._fddcbg*60)
	switch _ceecg._degeb {
	case '-':
		_babde = -_babde
	case 'Z':
		_babde = 0
	}
	_aagf := _bg.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _ceecg._degeb, _ceecg._gcccb, _ceecg._fddcbg)
	_dgcg := _f.FixedZone(_aagf, _babde)
	return _f.Date(int(_ceecg._dacdd), _f.Month(_ceecg._agaba), int(_ceecg._edcfe), int(_ceecg._aedbe), int(_ceecg._cfaba), int(_ceecg._cgbcd), 0, _dgcg)
}

// PdfShadingType7 is a Tensor-product patch mesh.
type PdfShadingType7 struct {
	*PdfShading
	BitsPerCoordinate *_ebb.PdfObjectInteger
	BitsPerComponent  *_ebb.PdfObjectInteger
	BitsPerFlag       *_ebb.PdfObjectInteger
	Decode            *_ebb.PdfObjectArray
	Function          []PdfFunction
}

func _bdfea() string {
	_daddc.Lock()
	defer _daddc.Unlock()
	return _feaca
}

// Hasher is the interface that wraps the basic Write method.
type Hasher interface {
	Write(_cfeba []byte) (_debdd int, _abeadc error)
}

// SetAction sets the PDF action for the annotation link.
func (_add *PdfAnnotationLink) SetAction(action *PdfAction) {
	_add._ffea = action
	if action == nil {
		_add.A = nil
	}
}

// NewPdfReaderLazy creates a new PdfReader for `rs` in lazy-loading mode. The difference
// from NewPdfReader is that in lazy-loading mode, objects are only loaded into memory when needed
// rather than entire structure being loaded into memory on reader creation.
// Note that it may make sense to use the lazy-load reader when processing only parts of files,
// rather than loading entire file into memory. Example: splitting a few pages from a large PDF file.
func NewPdfReaderLazy(rs _ab.ReadSeeker) (*PdfReader, error) {
	const _acbg = "\u006d\u006f\u0064\u0065l:\u004e\u0065\u0077\u0050\u0064\u0066\u0052\u0065\u0061\u0064\u0065\u0072\u004c\u0061z\u0079"
	return _dcbd(rs, &ReaderOpts{LazyLoad: true}, false, _acbg)
}

// PdfActionLaunch represents a launch action.
type PdfActionLaunch struct {
	*PdfAction
	F         *PdfFilespec
	Win       _ebb.PdfObject
	Mac       _ebb.PdfObject
	Unix      _ebb.PdfObject
	NewWindow _ebb.PdfObject
}

// SetPdfAuthor sets the Author attribute of the output PDF.
func SetPdfAuthor(author string) { _daddc.Lock(); defer _daddc.Unlock(); _ecgdd = author }

// SetSamples convert samples to byte-data and sets for the image.
// NOTE: The method resamples the data and this could lead to high memory usage,
// especially on large images. It should be used only when it is not possible
// to work with the image byte data directly.
func (_bgge *Image) SetSamples(samples []uint32) {
	if _bgge.BitsPerComponent < 8 {
		samples = _bgge.samplesAddPadding(samples)
	}
	_beea := _abg.ResampleUint32(samples, int(_bgge.BitsPerComponent), 8)
	_acgbf := make([]byte, len(_beea))
	for _ccbg, _fcgbg := range _beea {
		_acgbf[_ccbg] = byte(_fcgbg)
	}
	_bgge.Data = _acgbf
}

// PdfPage represents a page in a PDF document. (7.7.3.3 - Table 30).
type PdfPage struct {
	Parent               _ebb.PdfObject
	LastModified         *PdfDate
	Resources            *PdfPageResources
	CropBox              *PdfRectangle
	MediaBox             *PdfRectangle
	BleedBox             *PdfRectangle
	TrimBox              *PdfRectangle
	ArtBox               *PdfRectangle
	BoxColorInfo         _ebb.PdfObject
	Contents             _ebb.PdfObject
	Rotate               *int64
	Group                _ebb.PdfObject
	Thumb                _ebb.PdfObject
	B                    _ebb.PdfObject
	Dur                  _ebb.PdfObject
	Trans                _ebb.PdfObject
	AA                   _ebb.PdfObject
	Metadata             _ebb.PdfObject
	PieceInfo            _ebb.PdfObject
	StructParents        _ebb.PdfObject
	ID                   _ebb.PdfObject
	PZ                   _ebb.PdfObject
	SeparationInfo       _ebb.PdfObject
	Tabs                 _ebb.PdfObject
	TemplateInstantiated _ebb.PdfObject
	PresSteps            _ebb.PdfObject
	UserUnit             _ebb.PdfObject
	VP                   _ebb.PdfObject
	Annots               _ebb.PdfObject
	_bbfed               []*PdfAnnotation
	_cdbfde              *_ebb.PdfObjectDictionary
	_defbb               *_ebb.PdfIndirectObject
	_ddab                *PdfReader
}

// Inspect inspects the object types, subtypes and content in the PDF file returning a map of
// object type to number of instances of each.
func (_edbdg *PdfReader) Inspect() (map[string]int, error) { return _edbdg._cafdf.Inspect() }
func (_aabcbe *LTV) getOCSPs(_ccab []*_g.Certificate, _cgecc map[string]*_g.Certificate) ([][]byte, error) {
	_aecd := make([][]byte, 0, len(_ccab))
	for _, _acgc := range _ccab {
		for _, _cbad := range _acgc.OCSPServer {
			if _aabcbe.CertClient.IsCA(_acgc) {
				continue
			}
			_ggfg, _degg := _cgecc[_acgc.Issuer.CommonName]
			if !_degg {
				_eg.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _bacae, _bbcc := _aabcbe.OCSPClient.MakeRequest(_cbad, _acgc, _ggfg)
			if _bbcc != nil {
				_eg.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _bbcc)
				continue
			}
			_aecd = append(_aecd, _bacae)
		}
	}
	return _aecd, nil
}

// SetColorspaceByName adds the provided colorspace to the page resources.
func (_eegea *PdfPageResources) SetColorspaceByName(keyName _ebb.PdfObjectName, cs PdfColorspace) error {
	_bbbc, _fcffb := _eegea.GetColorspaces()
	if _fcffb != nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _fcffb)
		return _fcffb
	}
	if _bbbc == nil {
		_bbbc = NewPdfPageResourcesColorspaces()
		_eegea.SetColorSpace(_bbbc)
	}
	_bbbc.Set(keyName, cs)
	return nil
}

// ToGoImage converts the unidoc Image to a golang Image structure.
func (_adab *Image) ToGoImage() (_gdc.Image, error) {
	_eg.Log.Trace("\u0043\u006f\u006e\u0076er\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u006f\u0020\u0069\u006d\u0061g\u0065")
	_agebf, _ffbgb := _dg.NewImage(int(_adab.Width), int(_adab.Height), int(_adab.BitsPerComponent), _adab.ColorComponents, _adab.Data, _adab._dagcb, _adab._dgcea)
	if _ffbgb != nil {
		return nil, _ffbgb
	}
	return _agebf, nil
}

// NewPdfAnnotationPopup returns a new popup annotation.
func NewPdfAnnotationPopup() *PdfAnnotationPopup {
	_dgcf := NewPdfAnnotation()
	_ebf := &PdfAnnotationPopup{}
	_ebf.PdfAnnotation = _dgcf
	_dgcf.SetContext(_ebf)
	return _ebf
}

// ToPdfObject implements interface PdfModel.
func (_abgc *Permissions) ToPdfObject() _ebb.PdfObject { return _abgc._bdgdgf }
func (_bdfe *pdfFontSimple) baseFields() *fontCommon   { return &_bdfe.fontCommon }

// PdfAnnotationLink represents Link annotations.
// (Section 12.5.6.5 p. 403).
type PdfAnnotationLink struct {
	*PdfAnnotation
	A          _ebb.PdfObject
	Dest       _ebb.PdfObject
	H          _ebb.PdfObject
	PA         _ebb.PdfObject
	QuadPoints _ebb.PdfObject
	BS         _ebb.PdfObject
	_ffea      *PdfAction
	_cfbg      *PdfReader
}

// GetContainingPdfObject implements interface PdfModel.
func (_edabf *Permissions) GetContainingPdfObject() _ebb.PdfObject { return _edabf._bdgdgf }

// GetContext returns a reference to the subpattern entry: either PdfTilingPattern or PdfShadingPattern.
func (_abec *PdfPattern) GetContext() PdfModel { return _abec._ffagg }
func _fcgbc(_adgfb _bad.StdFont) pdfFontSimple {
	_eaabd := _adgfb.Descriptor()
	return pdfFontSimple{fontCommon: fontCommon{_dfbf: "\u0054\u0079\u0070e\u0031", _fdacg: _adgfb.Name()}, _ddgd: _adgfb.GetMetricsTable(), _adbd: &PdfFontDescriptor{FontName: _ebb.MakeName(string(_eaabd.Name)), FontFamily: _ebb.MakeName(_eaabd.Family), FontWeight: _ebb.MakeFloat(float64(_eaabd.Weight)), Flags: _ebb.MakeInteger(int64(_eaabd.Flags)), FontBBox: _ebb.MakeArrayFromFloats(_eaabd.BBox[:]), ItalicAngle: _ebb.MakeFloat(_eaabd.ItalicAngle), Ascent: _ebb.MakeFloat(_eaabd.Ascent), Descent: _ebb.MakeFloat(_eaabd.Descent), CapHeight: _ebb.MakeFloat(_eaabd.CapHeight), XHeight: _ebb.MakeFloat(_eaabd.XHeight), StemV: _ebb.MakeFloat(_eaabd.StemV), StemH: _ebb.MakeFloat(_eaabd.StemH)}, _dacee: _adgfb.Encoder()}
}

// NewXObjectForm creates a brand new XObject Form. Creates a new underlying PDF object stream primitive.
func NewXObjectForm() *XObjectForm {
	_abfcg := &XObjectForm{}
	_acggfa := &_ebb.PdfObjectStream{}
	_acggfa.PdfObjectDictionary = _ebb.MakeDict()
	_abfcg._gebcd = _acggfa
	return _abfcg
}

var (
	CourierName              = _bad.CourierName
	CourierBoldName          = _bad.CourierBoldName
	CourierObliqueName       = _bad.CourierObliqueName
	CourierBoldObliqueName   = _bad.CourierBoldObliqueName
	HelveticaName            = _bad.HelveticaName
	HelveticaBoldName        = _bad.HelveticaBoldName
	HelveticaObliqueName     = _bad.HelveticaObliqueName
	HelveticaBoldObliqueName = _bad.HelveticaBoldObliqueName
	SymbolName               = _bad.SymbolName
	ZapfDingbatsName         = _bad.ZapfDingbatsName
	TimesRomanName           = _bad.TimesRomanName
	TimesBoldName            = _bad.TimesBoldName
	TimesItalicName          = _bad.TimesItalicName
	TimesBoldItalicName      = _bad.TimesBoldItalicName
)

// Set sets the colorspace corresponding to key. Add to Names if not set.
func (_ggbae *PdfPageResourcesColorspaces) Set(key _ebb.PdfObjectName, val PdfColorspace) {
	if _, _gefgf := _ggbae.Colorspaces[string(key)]; !_gefgf {
		_ggbae.Names = append(_ggbae.Names, string(key))
	}
	_ggbae.Colorspaces[string(key)] = val
}

// GetContext returns the PdfField context which is the more specific field data type, e.g. PdfFieldButton
// for a button field.
func (_bbgd *PdfField) GetContext() PdfModel { return _bbgd._cada }

// NewPdfTransformParamsDocMDP create a PdfTransformParamsDocMDP with the specific permissions.
func NewPdfTransformParamsDocMDP(permission _ac.DocMDPPermission) *PdfTransformParamsDocMDP {
	return &PdfTransformParamsDocMDP{Type: _ebb.MakeName("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"), P: _ebb.MakeInteger(int64(permission)), V: _ebb.MakeName("\u0031\u002e\u0032")}
}

// NewPdfInfoFromObject creates a new PdfInfo from the input core.PdfObject.
func NewPdfInfoFromObject(obj _ebb.PdfObject) (*PdfInfo, error) {
	var _dfbgf PdfInfo
	_dddg, _abaa := obj.(*_ebb.PdfObjectDictionary)
	if !_abaa {
		return nil, _bg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", obj)
	}
	for _, _ggea := range _dddg.Keys() {
		switch _ggea {
		case "\u0054\u0069\u0074l\u0065":
			_dfbgf.Title, _ = _ebb.GetString(_dddg.Get("\u0054\u0069\u0074l\u0065"))
		case "\u0041\u0075\u0074\u0068\u006f\u0072":
			_dfbgf.Author, _ = _ebb.GetString(_dddg.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
		case "\u0053u\u0062\u006a\u0065\u0063\u0074":
			_dfbgf.Subject, _ = _ebb.GetString(_dddg.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
		case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
			_dfbgf.Keywords, _ = _ebb.GetString(_dddg.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
		case "\u0043r\u0065\u0061\u0074\u006f\u0072":
			_dfbgf.Creator, _ = _ebb.GetString(_dddg.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
		case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
			_dfbgf.Producer, _ = _ebb.GetString(_dddg.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
		case "\u0054r\u0061\u0070\u0070\u0065\u0064":
			_dfbgf.Trapped, _ = _ebb.GetName(_dddg.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
		case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
			if _gddf, _gddba := _ebb.GetString(_dddg.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _gddba && _gddf.String() != "" {
				_agbge, _dedc := NewPdfDate(_gddf.String())
				if _dedc != nil {
					return nil, _bg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _dedc)
				}
				_dfbgf.CreationDate = &_agbge
			}
		case "\u004do\u0064\u0044\u0061\u0074\u0065":
			if _afee, _fdbf := _ebb.GetString(_dddg.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _fdbf && _afee.String() != "" {
				_gcec, _cgaed := NewPdfDate(_afee.String())
				if _cgaed != nil {
					return nil, _bg.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _cgaed)
				}
				_dfbgf.ModifiedDate = &_gcec
			}
		default:
			_egdb, _ := _ebb.GetString(_dddg.Get(_ggea))
			if _dfbgf._gcgf == nil {
				_dfbgf._gcgf = _ebb.MakeDict()
			}
			_dfbgf._gcgf.Set(_ggea, _egdb)
		}
	}
	return &_dfbgf, nil
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_cegee *PdfFunctionType3) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_eg.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return nil, _gf.New("\u006e\u006f\u0074\u0020im\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
}
func (_dcg *PdfReader) newPdfActionSetOCGStateFromDict(_dag *_ebb.PdfObjectDictionary) (*PdfActionSetOCGState, error) {
	return &PdfActionSetOCGState{State: _dag.Get("\u0053\u0074\u0061t\u0065"), PreserveRB: _dag.Get("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042")}, nil
}

// NewPdfPage returns a new PDF page.
func NewPdfPage() *PdfPage {
	_cbfdf := PdfPage{}
	_cbfdf._cdbfde = _ebb.MakeDict()
	_cbfdf.Resources = NewPdfPageResources()
	_ccgf := _ebb.PdfIndirectObject{}
	_ccgf.PdfObject = _cbfdf._cdbfde
	_cbfdf._defbb = &_ccgf
	return &_cbfdf
}

// Y returns the value of the yellow component of the color.
func (_caeea *PdfColorDeviceCMYK) Y() float64 { return _caeea[2] }

// GetContainingPdfObject gets the primitive used to parse the color space.
func (_bccd *PdfColorspaceICCBased) GetContainingPdfObject() _ebb.PdfObject { return _bccd._dfff }

// DecodeArray returns the component range values for the Indexed colorspace.
func (_gaegg *PdfColorspaceSpecialIndexed) DecodeArray() []float64 {
	return []float64{0, float64(_gaegg.HiVal)}
}
func _gccbf(_cbcdc *_ebb.PdfObjectDictionary) (*PdfFieldButton, error) {
	_effg := &PdfFieldButton{}
	_effg.PdfField = NewPdfField()
	_effg.PdfField.SetContext(_effg)
	_effg.Opt, _ = _ebb.GetArray(_cbcdc.Get("\u004f\u0070\u0074"))
	_gbefa := NewPdfAnnotationWidget()
	_gbefa.A, _ = _ebb.GetDict(_cbcdc.Get("\u0041"))
	_gbefa.AP, _ = _ebb.GetDict(_cbcdc.Get("\u0041\u0050"))
	_gbefa.SetContext(_effg)
	_effg.PdfField.Annotations = append(_effg.PdfField.Annotations, _gbefa)
	return _effg, nil
}

// GetCatalogStructTreeRoot gets the catalog StructTreeRoot object.
func (_dgcdg *PdfReader) GetCatalogStructTreeRoot() (_ebb.PdfObject, bool) {
	if _dgcdg._fdgda == nil {
		return nil, false
	}
	_cggcf := _dgcdg._fdgda.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074")
	return _cggcf, _cggcf != nil
}

// AddCustomInfo adds a custom info into document info dictionary.
func (_gead *PdfInfo) AddCustomInfo(name string, value string) error {
	if _gead._gcgf == nil {
		_gead._gcgf = _ebb.MakeDict()
	}
	if _, _afcg := _eefgb[name]; _afcg {
		return _bg.Errorf("\u0063\u0061\u006e\u006e\u006ft\u0020\u0075\u0073\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072\u0064 \u0069\u006e\u0066\u006f\u0020\u006b\u0065\u0079\u0020\u0025\u0073\u0020\u0061\u0073\u0020\u0063\u0075\u0073\u0074\u006f\u006d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u006b\u0065y", name)
	}
	_gead._gcgf.SetIfNotNil(*_ebb.MakeName(name), _ebb.MakeString(value))
	return nil
}

// DecodeArray returns the range of color component values in DeviceGray colorspace.
func (_aggf *PdfColorspaceDeviceGray) DecodeArray() []float64 { return []float64{0, 1.0} }

// OutlineDest represents the destination of an outline item.
// It holds the page and the position on the page an outline item points to.
type OutlineDest struct {
	PageObj *_ebb.PdfIndirectObject `json:"-"`
	Page    int64                   `json:"page"`
	Mode    string                  `json:"mode"`
	X       float64                 `json:"x"`
	Y       float64                 `json:"y"`
	Zoom    float64                 `json:"zoom"`
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
func (_ebddf *PdfFont) CharcodeBytesToUnicode(data []byte) (string, int, int) {
	_gfeg, _, _ggae := _ebddf.CharcodesToUnicodeWithStats(_ebddf.BytesToCharcodes(data))
	_agff := _da.ExpandLigatures(_gfeg)
	return _agff, _de.RuneCountInString(_agff), _ggae
}
func (_fed *PdfReader) newPdfAnnotationSquareFromDict(_bbab *_ebb.PdfObjectDictionary) (*PdfAnnotationSquare, error) {
	_gbad := PdfAnnotationSquare{}
	_gfbc, _egf := _fed.newPdfAnnotationMarkupFromDict(_bbab)
	if _egf != nil {
		return nil, _egf
	}
	_gbad.PdfAnnotationMarkup = _gfbc
	_gbad.BS = _bbab.Get("\u0042\u0053")
	_gbad.IC = _bbab.Get("\u0049\u0043")
	_gbad.BE = _bbab.Get("\u0042\u0045")
	_gbad.RD = _bbab.Get("\u0052\u0044")
	return &_gbad, nil
}

// PdfTransformParamsDocMDP represents a transform parameters dictionary for the DocMDP method and is used to detect
// modifications relative to a signature field that is signed by the author of a document.
// (Section 12.8.2.2, Table 254 - Entries in the DocMDP transform parameters dictionary p. 471 in PDF32000_2008).
type PdfTransformParamsDocMDP struct {
	Type *_ebb.PdfObjectName
	P    *_ebb.PdfObjectInteger
	V    *_ebb.PdfObjectName
}

func (_bdabeb *PdfWriter) writeDocumentVersion() {
	if _bdabeb._abffb {
		_bdabeb.writeString("\u000a")
	} else {
		_bdabeb.writeString(_bg.Sprintf("\u0025\u0025\u0050D\u0046\u002d\u0025\u0064\u002e\u0025\u0064\u000a", _bdabeb._efcge.Major, _bdabeb._efcge.Minor))
		_bdabeb.writeString("\u0025\u00e2\u00e3\u00cf\u00d3\u000a")
	}
}

// ToPdfObject returns the text field dictionary within an indirect object (container).
func (_cgab *PdfFieldText) ToPdfObject() _ebb.PdfObject {
	_cgab.PdfField.ToPdfObject()
	_ffcc := _cgab._cdfd
	_fdcf := _ffcc.PdfObject.(*_ebb.PdfObjectDictionary)
	_fdcf.Set("\u0046\u0054", _ebb.MakeName("\u0054\u0078"))
	if _cgab.DA != nil {
		_fdcf.Set("\u0044\u0041", _cgab.DA)
	}
	if _cgab.Q != nil {
		_fdcf.Set("\u0051", _cgab.Q)
	}
	if _cgab.DS != nil {
		_fdcf.Set("\u0044\u0053", _cgab.DS)
	}
	if _cgab.RV != nil {
		_fdcf.Set("\u0052\u0056", _cgab.RV)
	}
	if _cgab.MaxLen != nil {
		_fdcf.Set("\u004d\u0061\u0078\u004c\u0065\u006e", _cgab.MaxLen)
	}
	return _ffcc
}

// SetPatternByName sets a pattern resource specified by keyName.
func (_geae *PdfPageResources) SetPatternByName(keyName _ebb.PdfObjectName, pattern _ebb.PdfObject) error {
	if _geae.Pattern == nil {
		_geae.Pattern = _ebb.MakeDict()
	}
	_geedg, _cfgcd := _geae.Pattern.(*_ebb.PdfObjectDictionary)
	if !_cfgcd {
		return _ebb.ErrTypeError
	}
	_geedg.Set(keyName, pattern)
	return nil
}

// PdfAction represents an action in PDF (section 12.6 p. 412).
type PdfAction struct {
	_ad  PdfModel
	Type _ebb.PdfObject
	S    _ebb.PdfObject
	Next _ebb.PdfObject
	_abe *_ebb.PdfIndirectObject
}

// ToPdfObject converts the pdfCIDFontType2 to a PDF representation.
func (_cfcga *pdfCIDFontType2) ToPdfObject() _ebb.PdfObject {
	if _cfcga._ddea == nil {
		_cfcga._ddea = &_ebb.PdfIndirectObject{}
	}
	_dggc := _cfcga.baseFields().asPdfObjectDictionary("\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032")
	_cfcga._ddea.PdfObject = _dggc
	if _cfcga.CIDSystemInfo != nil {
		_dggc.Set("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f", _cfcga.CIDSystemInfo)
	}
	if _cfcga.DW != nil {
		_dggc.Set("\u0044\u0057", _cfcga.DW)
	}
	if _cfcga.DW2 != nil {
		_dggc.Set("\u0044\u0057\u0032", _cfcga.DW2)
	}
	if _cfcga.W != nil {
		_dggc.Set("\u0057", _cfcga.W)
	}
	if _cfcga.W2 != nil {
		_dggc.Set("\u0057\u0032", _cfcga.W2)
	}
	if _cfcga.CIDToGIDMap != nil {
		_dggc.Set("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070", _cfcga.CIDToGIDMap)
	}
	return _cfcga._ddea
}

// GetContainingPdfObject returns the container of the resources object (indirect object).
func (_agfg *PdfPageResources) GetContainingPdfObject() _ebb.PdfObject { return _agfg._efbed }
func _ggbgd(_faac _ebb.PdfObject) (*PdfFunctionType3, error) {
	_aegfb := &PdfFunctionType3{}
	var _cffec *_ebb.PdfObjectDictionary
	if _fcffd, _fcea := _faac.(*_ebb.PdfIndirectObject); _fcea {
		_ebfecc, _egac := _fcffd.PdfObject.(*_ebb.PdfObjectDictionary)
		if !_egac {
			return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_aegfb._acac = _fcffd
		_cffec = _ebfecc
	} else if _gcaf, _fcaa := _faac.(*_ebb.PdfObjectDictionary); _fcaa {
		_cffec = _gcaf
	} else {
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gfdba, _gagb := _ebb.TraceToDirectObject(_cffec.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_ebb.PdfObjectArray)
	if !_gagb {
		_eg.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _gfdba.Len() != 2 {
		_eg.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _gf.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_gdaeg, _dfdb := _gfdba.ToFloat64Array()
	if _dfdb != nil {
		return nil, _dfdb
	}
	_aegfb.Domain = _gdaeg
	_gfdba, _gagb = _ebb.TraceToDirectObject(_cffec.Get("\u0052\u0061\u006eg\u0065")).(*_ebb.PdfObjectArray)
	if _gagb {
		if _gfdba.Len() < 0 || _gfdba.Len()%2 != 0 {
			return nil, _gf.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_dfbee, _afcfc := _gfdba.ToFloat64Array()
		if _afcfc != nil {
			return nil, _afcfc
		}
		_aegfb.Range = _dfbee
	}
	_gfdba, _gagb = _ebb.TraceToDirectObject(_cffec.Get("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es")).(*_ebb.PdfObjectArray)
	if !_gagb {
		_eg.Log.Error("\u0046\u0075\u006ect\u0069\u006f\u006e\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_aegfb.Functions = []PdfFunction{}
	for _, _fffc := range _gfdba.Elements() {
		_gbdf, _edcef := _aagg(_fffc)
		if _edcef != nil {
			return nil, _edcef
		}
		_aegfb.Functions = append(_aegfb.Functions, _gbdf)
	}
	_gfdba, _gagb = _ebb.TraceToDirectObject(_cffec.Get("\u0042\u006f\u0075\u006e\u0064\u0073")).(*_ebb.PdfObjectArray)
	if !_gagb {
		_eg.Log.Error("B\u006fu\u006e\u0064\u0073\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_ebfbg, _dfdb := _gfdba.ToFloat64Array()
	if _dfdb != nil {
		return nil, _dfdb
	}
	_aegfb.Bounds = _ebfbg
	if len(_aegfb.Bounds) != len(_aegfb.Functions)-1 {
		_eg.Log.Error("B\u006f\u0075\u006e\u0064\u0073\u0020\u0028\u0025\u0064)\u0020\u0061\u006e\u0064\u0020\u006e\u0075m \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029 n\u006f\u0074 \u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067", len(_aegfb.Bounds), len(_aegfb.Functions))
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gfdba, _gagb = _ebb.TraceToDirectObject(_cffec.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_ebb.PdfObjectArray)
	if !_gagb {
		_eg.Log.Error("E\u006ec\u006f\u0064\u0065\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _gf.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_dffgg, _dfdb := _gfdba.ToFloat64Array()
	if _dfdb != nil {
		return nil, _dfdb
	}
	_aegfb.Encode = _dffgg
	if len(_aegfb.Encode) != 2*len(_aegfb.Functions) {
		_eg.Log.Error("\u004c\u0065\u006e\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0028\u0025\u0064\u0029 \u0061\u006e\u0064\u0020\u006e\u0075\u006d\u0020\u0066\u0075\u006e\u0063\u0074i\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029\u0020\u006e\u006f\u0074 m\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u0075\u0070", len(_aegfb.Encode), len(_aegfb.Functions))
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return _aegfb, nil
}

// ToPdfObject implements interface PdfModel.
func (_eab *PdfAnnotationSound) ToPdfObject() _ebb.PdfObject {
	_eab.PdfAnnotation.ToPdfObject()
	_eddf := _eab._bdcd
	_bfe := _eddf.PdfObject.(*_ebb.PdfObjectDictionary)
	_eab.PdfAnnotationMarkup.appendToPdfDictionary(_bfe)
	_bfe.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0053\u006f\u0075n\u0064"))
	_bfe.SetIfNotNil("\u0053\u006f\u0075n\u0064", _eab.Sound)
	_bfe.SetIfNotNil("\u004e\u0061\u006d\u0065", _eab.Name)
	return _eddf
}

// NewStandard14FontMustCompile returns the standard 14 font named `basefont` as a *PdfFont.
// If `basefont` is one of the 14 Standard14Font values defined above then NewStandard14FontMustCompile
// is guaranteed to succeed.
func NewStandard14FontMustCompile(basefont StdFontName) *PdfFont {
	_dgggc, _gagd := NewStandard14Font(basefont)
	if _gagd != nil {
		panic(_bg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0031\u0034\u0046\u006f\u006e\u0074 \u0025\u0023\u0071", basefont))
	}
	return _dgggc
}

// PdfOutlineTreeNode contains common fields used by the outline and outline
// item objects.
type PdfOutlineTreeNode struct {
	_geeee interface{}
	First  *PdfOutlineTreeNode
	Last   *PdfOutlineTreeNode
}

func _fbgbf(_dcabfc _ebb.PdfObject) {
	_eg.Log.Debug("\u006f\u0062\u006a\u003a\u0020\u0025\u0054\u0020\u0025\u0073", _dcabfc, _dcabfc.String())
	if _dgccf, _dffce := _dcabfc.(*_ebb.PdfObjectStream); _dffce {
		_adafc, _addfe := _ebb.DecodeStream(_dgccf)
		if _addfe != nil {
			_eg.Log.Debug("\u0045r\u0072\u006f\u0072\u003a\u0020\u0025v", _addfe)
			return
		}
		_eg.Log.Debug("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _adafc)
	} else if _gcbf, _eegee := _dcabfc.(*_ebb.PdfIndirectObject); _eegee {
		_eg.Log.Debug("\u0025\u0054\u0020%\u0076", _gcbf.PdfObject, _gcbf.PdfObject)
		_eg.Log.Debug("\u0025\u0073", _gcbf.PdfObject.String())
	}
}

// GetNameDictionary returns the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_afegb *PdfReader) GetNameDictionary() (_ebb.PdfObject, error) {
	_ebefa := _ebb.ResolveReference(_afegb._fdgda.Get("\u004e\u0061\u006de\u0073"))
	if _ebefa == nil {
		return nil, nil
	}
	if !_afegb._ceefa {
		_accdd := _afegb.traverseObjectData(_ebefa)
		if _accdd != nil {
			return nil, _accdd
		}
	}
	return _ebefa, nil
}

// StringToCharcodeBytes maps the provided string runes to charcode bytes and
// it returns the resulting slice of bytes, along with the number of runes
// which could not be converted. If the number of misses is 0, all string runes
// were successfully converted.
func (_abbc *PdfFont) StringToCharcodeBytes(str string) ([]byte, int) {
	return _abbc.RunesToCharcodeBytes([]rune(str))
}

// Outline represents a PDF outline dictionary (Table 152 - p. 376).
// Currently, the Outline object can only be used to construct PDF outlines.
type Outline struct {
	Entries []*OutlineItem `json:"entries,omitempty"`
}

// ImageHandler interface implements common image loading and processing tasks.
// Implementing as an interface allows for the possibility to use non-standard libraries for faster
// loading and processing of images.
type ImageHandler interface {

	// Read any image type and load into a new Image object.
	Read(_fdgd _ab.Reader) (*Image, error)

	// NewImageFromGoImage loads a NRGBA32 unidoc Image from a standard Go image structure.
	NewImageFromGoImage(_cccf _gdc.Image) (*Image, error)

	// NewGrayImageFromGoImage loads a grayscale unidoc Image from a standard Go image structure.
	NewGrayImageFromGoImage(_dedde _gdc.Image) (*Image, error)

	// Compress an image.
	Compress(_dbcb *Image, _cdfae int64) (*Image, error)
}

// ToPdfObject returns a PDF object representation of the outline item.
func (_cgagd *OutlineItem) ToPdfObject() _ebb.PdfObject {
	_edaaf, _ := _cgagd.ToPdfOutlineItem()
	return _edaaf.ToPdfObject()
}

// ToPdfObject converts the font to a PDF representation.
func (_abbgd *pdfFontType3) ToPdfObject() _ebb.PdfObject {
	if _abbgd._adec == nil {
		_abbgd._adec = &_ebb.PdfIndirectObject{}
	}
	_beebd := _abbgd.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0033")
	_abbgd._adec.PdfObject = _beebd
	if _abbgd.FirstChar != nil {
		_beebd.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _abbgd.FirstChar)
	}
	if _abbgd.LastChar != nil {
		_beebd.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _abbgd.LastChar)
	}
	if _abbgd.Widths != nil {
		_beebd.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _abbgd.Widths)
	}
	if _abbgd.Encoding != nil {
		_beebd.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _abbgd.Encoding)
	} else if _abbgd._abcba != nil {
		_efag := _abbgd._abcba.ToPdfObject()
		if _efag != nil {
			_beebd.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _efag)
		}
	}
	if _abbgd.FontBBox != nil {
		_beebd.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _abbgd.FontBBox)
	}
	if _abbgd.FontMatrix != nil {
		_beebd.Set("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0069\u0072\u0078", _abbgd.FontMatrix)
	}
	if _abbgd.CharProcs != nil {
		_beebd.Set("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s", _abbgd.CharProcs)
	}
	if _abbgd.Resources != nil {
		_beebd.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _abbgd.Resources)
	}
	return _abbgd._adec
}
func (_afcf *PdfReader) flattenFieldsWithOpts(_gada bool, _gfbb FieldAppearanceGenerator, _dbfe *FieldFlattenOpts) error {
	if _dbfe == nil {
		_dbfe = &FieldFlattenOpts{}
	}
	var _badf bool
	_affeb := map[*PdfAnnotation]bool{}
	{
		var _febed []*PdfField
		_cagb := _afcf.AcroForm
		if _cagb != nil {
			if _dbfe.FilterFunc != nil {
				_febed = _cagb.filteredFields(_dbfe.FilterFunc, true)
				_badf = _cagb.Fields != nil && len(*_cagb.Fields) > 0
			} else {
				_febed = _cagb.AllFields()
			}
		}
		for _, _eggbe := range _febed {
			for _, _ceabg := range _eggbe.Annotations {
				_affeb[_ceabg.PdfAnnotation] = _eggbe.V != nil
				if _gfbb != nil {
					_gfbcg, _aedab := _gfbb.GenerateAppearanceDict(_cagb, _eggbe, _ceabg)
					if _aedab != nil {
						return _aedab
					}
					_ceabg.AP = _gfbcg
				}
			}
		}
	}
	if _gada {
		for _, _fcbb := range _afcf.PageList {
			_gdagg, _ebec := _fcbb.GetAnnotations()
			if _ebec != nil {
				return _ebec
			}
			for _, _dcfc := range _gdagg {
				_affeb[_dcfc] = true
			}
		}
	}
	for _, _egag := range _afcf.PageList {
		var _dcfbb []*PdfAnnotation
		if _gfbb != nil {
			if _fage := _gfbb.WrapContentStream(_egag); _fage != nil {
				return _fage
			}
		}
		_ecfd, _adfedf := _egag.GetAnnotations()
		if _adfedf != nil {
			return _adfedf
		}
		for _, _cbgf := range _ecfd {
			_ceb, _dfbd := _affeb[_cbgf]
			if !_dfbd && _dbfe.AnnotFilterFunc != nil {
				if _, _fadg := _cbgf.GetContext().(*PdfAnnotationWidget); !_fadg {
					_dfbd = _dbfe.AnnotFilterFunc(_cbgf)
				}
			}
			if !_dfbd {
				_dcfbb = append(_dcfbb, _cbgf)
				continue
			}
			switch _cbgf.GetContext().(type) {
			case *PdfAnnotationPopup:
				continue
			case *PdfAnnotationLink:
				continue
			case *PdfAnnotationProjection:
				continue
			}
			_deebg, _gdbg, _eccaa := _bdgcb(_cbgf)
			if _eccaa != nil {
				if !_ceb {
					_eg.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0056\u0020\u002d\u003e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075t\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065am\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072")
					continue
				}
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d,\u0020\u0065\u0072\u0072\u0020\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006bi\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072", _eccaa)
				continue
			}
			if _deebg == nil {
				continue
			}
			_dfce := _egag.Resources.GenerateXObjectName()
			_egag.Resources.SetXObjectFormByName(_dfce, _deebg)
			_gfgc := _cbg.Min(_gdbg.Llx, _gdbg.Urx)
			_ecfbb := _cbg.Min(_gdbg.Lly, _gdbg.Ury)
			var _bdgab []string
			_bdgab = append(_bdgab, "\u0071")
			_bdgab = append(_bdgab, _bg.Sprintf("\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020%\u002e6\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0063\u006d", 1.0, 0.0, 0.0, 1.0, _gfgc, _ecfbb))
			_bdgab = append(_bdgab, _bg.Sprintf("\u002f\u0025\u0073\u0020\u0044\u006f", _dfce.String()))
			_bdgab = append(_bdgab, "\u0051")
			_fdefd := _ee.Join(_bdgab, "\u000a")
			_eccaa = _egag.AppendContentStream(_fdefd)
			if _eccaa != nil {
				return _eccaa
			}
			if _deebg.Resources != nil {
				_ffecg, _addc := _ebb.GetDict(_deebg.Resources.Font)
				if _addc {
					for _, _daegg := range _ffecg.Keys() {
						if !_egag.Resources.HasFontByName(_daegg) {
							_egag.Resources.SetFontByName(_daegg, _ffecg.Get(_daegg))
						}
					}
				}
			}
		}
		if len(_dcfbb) > 0 {
			_egag._bbfed = _dcfbb
		} else {
			_egag._bbfed = []*PdfAnnotation{}
		}
	}
	if !_badf {
		_afcf.AcroForm = nil
	}
	return nil
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_fadb *PdfColorspaceSpecialSeparation) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ecbb, _fdcb := _ebb.GetNumbersAsFloat(objects)
	if _fdcb != nil {
		return nil, _fdcb
	}
	return _fadb.ColorFromFloats(_ecbb)
}

// ToInteger convert to an integer format.
func (_aefb *PdfColorDeviceRGB) ToInteger(bits int) [3]uint32 {
	_efccd := _cbg.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_efccd * _aefb.R()), uint32(_efccd * _aefb.G()), uint32(_efccd * _aefb.B())}
}

// NewPdfFieldSignature returns an initialized signature field.
func NewPdfFieldSignature(signature *PdfSignature) *PdfFieldSignature {
	_cgcf := &PdfFieldSignature{}
	_cgcf.PdfField = NewPdfField()
	_cgcf.PdfField.SetContext(_cgcf)
	_cgcf.PdfAnnotationWidget = NewPdfAnnotationWidget()
	_cgcf.PdfAnnotationWidget.SetContext(_cgcf)
	_cgcf.PdfAnnotationWidget._bdcd = _cgcf.PdfField._cdfd
	_cgcf.T = _ebb.MakeString("")
	_cgcf.F = _ebb.MakeInteger(132)
	_cgcf.V = signature
	return _cgcf
}
func (_gdg *PdfReader) newPdfActionSoundFromDict(_fbg *_ebb.PdfObjectDictionary) (*PdfActionSound, error) {
	return &PdfActionSound{Sound: _fbg.Get("\u0053\u006f\u0075n\u0064"), Volume: _fbg.Get("\u0056\u006f\u006c\u0075\u006d\u0065"), Synchronous: _fbg.Get("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073"), Repeat: _fbg.Get("\u0052\u0065\u0070\u0065\u0061\u0074"), Mix: _fbg.Get("\u004d\u0069\u0078")}, nil
}

// GetContainingPdfObject returns the container of the shading object (indirect object).
func (_ebead *PdfShading) GetContainingPdfObject() _ebb.PdfObject { return _ebead._fbfae }

// Mask returns the uin32 bitmask for the specific flag.
func (_dfagf FieldFlag) Mask() uint32 { return uint32(_dfagf) }

// NewPdfSignatureReferenceDocMDP returns PdfSignatureReference for the transformParams.
func NewPdfSignatureReferenceDocMDP(transformParams *PdfTransformParamsDocMDP) *PdfSignatureReference {
	return &PdfSignatureReference{Type: _ebb.MakeName("\u0053\u0069\u0067\u0052\u0065\u0066"), TransformMethod: _ebb.MakeName("\u0044\u006f\u0063\u004d\u0044\u0050"), TransformParams: transformParams.ToPdfObject()}
}

// GetXObjectImageByName returns the XObjectImage with the specified name from the
// page resources, if it exists.
func (_bdbe *PdfPageResources) GetXObjectImageByName(keyName _ebb.PdfObjectName) (*XObjectImage, error) {
	_bcfca, _dbdd := _bdbe.GetXObjectByName(keyName)
	if _bcfca == nil {
		return nil, nil
	}
	if _dbdd != XObjectTypeImage {
		return nil, _gf.New("\u006e\u006f\u0074 \u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065")
	}
	_baea, _bbece := NewXObjectImageFromStream(_bcfca)
	if _bbece != nil {
		return nil, _bbece
	}
	return _baea, nil
}

// GetContainingPdfObject implements model.PdfModel interface.
func (_bbebf *PdfOutputIntent) GetContainingPdfObject() _ebb.PdfObject { return _bbebf._faeb }

// A PdfPattern can represent a Pattern, either a tiling pattern or a shading pattern.
// Note that all patterns shall be treated as colours; a Pattern colour space shall be established with the CS or cs
// operator just like other colour spaces, and a particular pattern shall be installed as the current colour with the
// SCN or scn operator.
type PdfPattern struct {

	// Type: Pattern
	PatternType int64
	_ffagg      PdfModel
	_dcddc      _ebb.PdfObject
}

// HasFontByName checks whether a font is defined by the specified keyName.
func (_egccf *PdfPageResources) HasFontByName(keyName _ebb.PdfObjectName) bool {
	_, _cfeee := _egccf.GetFontByName(keyName)
	return _cfeee
}

// GetType returns the button field type which returns one of the following
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
func (_agdg *PdfFieldButton) GetType() ButtonType {
	_ebea := ButtonTypeCheckbox
	if _agdg.Ff != nil {
		if (uint32(*_agdg.Ff) & FieldFlagPushbutton.Mask()) > 0 {
			_ebea = ButtonTypePush
		} else if (uint32(*_agdg.Ff) & FieldFlagRadio.Mask()) > 0 {
			_ebea = ButtonTypeRadio
		}
	}
	return _ebea
}
func _aggb(_gbga *_ebb.PdfObjectDictionary, _eaec *fontCommon) (*pdfFontType0, error) {
	_bebfdb, _gbddg := _ebb.GetArray(_gbga.Get("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073"))
	if !_gbddg {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006cid\u0020\u0044\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073\u0020\u002d\u0020\u006e\u006f\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079 \u0025\u0073", _eaec)
		return nil, _ebb.ErrRangeError
	}
	if _bebfdb.Len() != 1 {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0041\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020(%\u0064\u0029", _bebfdb.Len())
		return nil, _ebb.ErrRangeError
	}
	_baeg, _bbgaa := _ddacd(_bebfdb.Get(0), false)
	if _bbgaa != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046a\u0069\u006c\u0065d \u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u0020\u0066\u006f\u006et\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076 \u0025\u0073", _bbgaa, _eaec)
		return nil, _bbgaa
	}
	_cbbc := _gegg(_eaec)
	_cbbc.DescendantFont = _baeg
	_ebbee, _gbddg := _ebb.GetNameVal(_gbga.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if _gbddg {
		if _ebbee == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048" || _ebbee == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056" {
			_cbbc._bfdgc = _da.NewIdentityTextEncoder(_ebbee)
		} else if _ebe.IsPredefinedCMap(_ebbee) {
			_cbbc._efeb, _bbgaa = _ebe.LoadPredefinedCMap(_ebbee)
			if _bbgaa != nil {
				_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _ebbee, _bbgaa)
			}
		} else {
			_eg.Log.Debug("\u0055\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0063\u006da\u0070\u0020\u0025\u0071", _ebbee)
		}
	}
	if _gaef := _baeg.baseFields()._dcdd; _gaef != nil {
		if _bcbaa := _gaef.Name(); _bcbaa == "\u0041d\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002d\u0055\u0043\u0053\u0032" || _bcbaa == "\u0041\u0064\u006f\u0062\u0065\u002d\u0047\u0042\u0031-\u0055\u0043\u0053\u0032" || _bcbaa == "\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031-\u0055\u0043\u0053\u0032" || _bcbaa == "\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006f\u0072\u0065\u0061\u0031-\u0055\u0043\u0053\u0032" {
			_cbbc._bfdgc = _da.NewCMapEncoder(_ebbee, _cbbc._efeb, _gaef)
		}
	}
	return _cbbc, nil
}

// Read reads an image and loads into a new Image object with an RGB
// colormap and 8 bits per component.
func (_affac DefaultImageHandler) Read(reader _ab.Reader) (*Image, error) {
	_daef, _, _bdfg := _gdc.Decode(reader)
	if _bdfg != nil {
		_eg.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _bdfg)
		return nil, _bdfg
	}
	return _affac.NewImageFromGoImage(_daef)
}

// ToPdfObject implements interface PdfModel.
func (_abfa *PdfAnnotationWatermark) ToPdfObject() _ebb.PdfObject {
	_abfa.PdfAnnotation.ToPdfObject()
	_ebdc := _abfa._bdcd
	_gaec := _ebdc.PdfObject.(*_ebb.PdfObjectDictionary)
	_gaec.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k"))
	_gaec.SetIfNotNil("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074", _abfa.FixedPrint)
	return _ebdc
}
func (_degdc fontCommon) isCIDFont() bool {
	if _degdc._dfbf == "" {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0069\u0073\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u002e\u0020\u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _degdc)
	}
	_gbbd := false
	switch _degdc._dfbf {
	case "\u0054\u0079\u0070e\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_gbbd = true
	}
	_eg.Log.Trace("i\u0073\u0043\u0049\u0044\u0046\u006fn\u0074\u003a\u0020\u0069\u0073\u0043\u0049\u0044\u003d%\u0074\u0020\u0066o\u006et\u003d\u0025\u0073", _gbbd, _degdc)
	return _gbbd
}

// PdfColorLab represents a color in the L*, a*, b* 3 component colorspace.
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorLab [3]float64

// ToInteger convert to an integer format.
func (_ggbb *PdfColorCalGray) ToInteger(bits int) uint32 {
	_cddg := _cbg.Pow(2, float64(bits)) - 1
	return uint32(_cddg * _ggbb.Val())
}

// SetXObjectByName adds the XObject from the passed in stream to the page resources.
// The added XObject is identified by the specified name.
func (_bafg *PdfPageResources) SetXObjectByName(keyName _ebb.PdfObjectName, stream *_ebb.PdfObjectStream) error {
	if _bafg.XObject == nil {
		_bafg.XObject = _ebb.MakeDict()
	}
	_eedge := _ebb.TraceToDirectObject(_bafg.XObject)
	_abgge, _bfaefg := _eedge.(*_ebb.PdfObjectDictionary)
	if !_bfaefg {
		_eg.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0058\u004f\u0062j\u0065\u0063\u0074\u002c\u0020\u0067\u006f\u0074\u0020\u0025T\u002f\u0025\u0054", _bafg.XObject, _eedge)
		return _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_abgge.Set(keyName, stream)
	return nil
}
func (_dege *PdfAppender) replaceObject(_efcg, _fdad _ebb.PdfObject) {
	switch _fgeee := _efcg.(type) {
	case *_ebb.PdfIndirectObject:
		_dege._gbfa[_fdad] = _fgeee.ObjectNumber
	case *_ebb.PdfObjectStream:
		_dege._gbfa[_fdad] = _fgeee.ObjectNumber
	}
}

// NewPdfAnnotationRichMedia returns a new rich media annotation.
func NewPdfAnnotationRichMedia() *PdfAnnotationRichMedia {
	_dab := NewPdfAnnotation()
	_afc := &PdfAnnotationRichMedia{}
	_afc.PdfAnnotation = _dab
	_dab.SetContext(_afc)
	return _afc
}
func (_fgg *PdfReader) newPdfActionGotoRFromDict(_eae *_ebb.PdfObjectDictionary) (*PdfActionGoToR, error) {
	_cfbf, _gdab := _gggf(_eae.Get("\u0046"))
	if _gdab != nil {
		return nil, _gdab
	}
	return &PdfActionGoToR{D: _eae.Get("\u0044"), NewWindow: _eae.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _cfbf}, nil
}
func (_fgggde *PdfWriter) writeBytes(_abdbb []byte) {
	if _fgggde._bgef != nil {
		return
	}
	_acdbg, _fegg := _fgggde._cbabb.Write(_abdbb)
	_fgggde._afedd += int64(_acdbg)
	_fgggde._bgef = _fegg
}
func (_cgcg *PdfAppender) mergeResources(_cegb, _cage _ebb.PdfObject, _gge map[_ebb.PdfObjectName]_ebb.PdfObjectName) _ebb.PdfObject {
	if _cage == nil && _cegb == nil {
		return nil
	}
	if _cage == nil {
		return _cegb
	}
	_geff, _acdd := _ebb.GetDict(_cage)
	if !_acdd {
		return _cegb
	}
	if _cegb == nil {
		_dgde := _ebb.MakeDict()
		_dgde.Merge(_geff)
		return _cage
	}
	_dgdb, _acdd := _ebb.GetDict(_cegb)
	if !_acdd {
		_eg.Log.Error("\u0045\u0072\u0072or\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		_dgdb = _ebb.MakeDict()
	}
	for _, _fbeb := range _geff.Keys() {
		if _deb, _dgcb := _gge[_fbeb]; _dgcb {
			_dgdb.Set(_deb, _geff.Get(_fbeb))
		} else {
			_dgdb.Set(_fbeb, _geff.Get(_fbeb))
		}
	}
	return _dgdb
}

// NewPdfPageResourcesColorspaces returns a new PdfPageResourcesColorspaces object.
func NewPdfPageResourcesColorspaces() *PdfPageResourcesColorspaces {
	_acbc := &PdfPageResourcesColorspaces{}
	_acbc.Names = []string{}
	_acbc.Colorspaces = map[string]PdfColorspace{}
	_acbc._ddffd = &_ebb.PdfIndirectObject{}
	return _acbc
}

// GetCerts returns the signature certificate chain.
func (_fadac *PdfSignature) GetCerts() ([]*_g.Certificate, error) {
	var _dabcd []func() ([]*_g.Certificate, error)
	switch _cbfc, _ := _ebb.GetNameVal(_fadac.SubFilter); _cbfc {
	case "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064", "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064":
		_dabcd = append(_dabcd, _fadac.extractChainFromPKCS7, _fadac.extractChainFromCert)
	case "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031":
		_dabcd = append(_dabcd, _fadac.extractChainFromCert)
	case "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031":
		_dabcd = append(_dabcd, _fadac.extractChainFromPKCS7)
	default:
		return nil, _bg.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020S\u0075b\u0046i\u006c\u0074\u0065\u0072\u003a\u0020\u0025s", _cbfc)
	}
	for _, _eadcd := range _dabcd {
		_dbaad, _aadf := _eadcd()
		if _aadf != nil {
			return nil, _aadf
		}
		if len(_dbaad) > 0 {
			return _dbaad, nil
		}
	}
	return nil, ErrSignNoCertificates
}
func (_bedg *PdfReader) newPdfAnnotationInkFromDict(_abac *_ebb.PdfObjectDictionary) (*PdfAnnotationInk, error) {
	_eece := PdfAnnotationInk{}
	_aedc, _aad := _bedg.newPdfAnnotationMarkupFromDict(_abac)
	if _aad != nil {
		return nil, _aad
	}
	_eece.PdfAnnotationMarkup = _aedc
	_eece.InkList = _abac.Get("\u0049n\u006b\u004c\u0069\u0073\u0074")
	_eece.BS = _abac.Get("\u0042\u0053")
	return &_eece, nil
}

// SetPdfModifiedDate sets the ModDate attribute of the output PDF.
func SetPdfModifiedDate(modifiedDate _f.Time) {
	_daddc.Lock()
	defer _daddc.Unlock()
	_ccdff = modifiedDate
}

// ToPdfObject implements interface PdfModel.
func (_fgee *PdfAnnotationLine) ToPdfObject() _ebb.PdfObject {
	_fgee.PdfAnnotation.ToPdfObject()
	_cabf := _fgee._bdcd
	_cea := _cabf.PdfObject.(*_ebb.PdfObjectDictionary)
	_fgee.PdfAnnotationMarkup.appendToPdfDictionary(_cea)
	_cea.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u004c\u0069\u006e\u0065"))
	_cea.SetIfNotNil("\u004c", _fgee.L)
	_cea.SetIfNotNil("\u0042\u0053", _fgee.BS)
	_cea.SetIfNotNil("\u004c\u0045", _fgee.LE)
	_cea.SetIfNotNil("\u0049\u0043", _fgee.IC)
	_cea.SetIfNotNil("\u004c\u004c", _fgee.LL)
	_cea.SetIfNotNil("\u004c\u004c\u0045", _fgee.LLE)
	_cea.SetIfNotNil("\u0043\u0061\u0070", _fgee.Cap)
	_cea.SetIfNotNil("\u0049\u0054", _fgee.IT)
	_cea.SetIfNotNil("\u004c\u004c\u004f", _fgee.LLO)
	_cea.SetIfNotNil("\u0043\u0050", _fgee.CP)
	_cea.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _fgee.Measure)
	_cea.SetIfNotNil("\u0043\u004f", _fgee.CO)
	return _cabf
}

var _ pdfFont = (*pdfCIDFontType2)(nil)

// ApplyStandard is used to apply changes required on the document to match the rules required by the input standard.
// The writer's content would be changed after all the document parts are already established during the Write method.
// A good example of the StandardApplier could be a PDF/A Profile (i.e.: pdfa.Profile1A). In such a case PdfWriter would
// set up all rules required by that Profile.
func (_bgdaae *PdfWriter) ApplyStandard(optimizer StandardApplier) { _bgdaae._cafac = optimizer }

// ReplaceAcroForm replaces the acrobat form. It appends a new form to the Pdf which
// replaces the original AcroForm.
func (_afcd *PdfAppender) ReplaceAcroForm(acroForm *PdfAcroForm) {
	if acroForm != nil {
		_afcd.updateObjectsDeep(acroForm.ToPdfObject(), nil)
	}
	_afcd._bfef = acroForm
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// red, green and blue components of the color. The values of the elements
// should be between 0 and 1.
func (_dbcf *PdfColorspaceDeviceRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fcgf := vals[0]
	if _fcgf < 0.0 || _fcgf > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fcgf)
		return nil, ErrColorOutOfRange
	}
	_dfgg := vals[1]
	if _dfgg < 0.0 || _dfgg > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fcgf)
		return nil, ErrColorOutOfRange
	}
	_bfad := vals[2]
	if _bfad < 0.0 || _bfad > 1.0 {
		_eg.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _fcgf)
		return nil, ErrColorOutOfRange
	}
	_aecb := NewPdfColorDeviceRGB(_fcgf, _dfgg, _bfad)
	return _aecb, nil
}

// ColorFromFloats returns a new PdfColor based on input color components.
func (_afef *PdfColorspaceDeviceN) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != _afef.GetNumComponents() {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_becea, _efcaa := _afef.TintTransform.Evaluate(vals)
	if _efcaa != nil {
		return nil, _efcaa
	}
	_dfad, _efcaa := _afef.AlternateSpace.ColorFromFloats(_becea)
	if _efcaa != nil {
		return nil, _efcaa
	}
	return _dfad, nil
}

// PdfFunctionType4 is a Postscript calculator functions.
type PdfFunctionType4 struct {
	Domain  []float64
	Range   []float64
	Program *_bc.PSProgram
	_bdbae  *_bc.PSExecutor
	_abbbb  []byte
	_gbgd   *_ebb.PdfObjectStream
}

// PdfActionGoToR represents a GoToR action.
type PdfActionGoToR struct {
	*PdfAction
	F         *PdfFilespec
	D         _ebb.PdfObject
	NewWindow _ebb.PdfObject
}

func (_gdgc *PdfReader) newPdfAnnotationSquigglyFromDict(_eegc *_ebb.PdfObjectDictionary) (*PdfAnnotationSquiggly, error) {
	_fcfdd := PdfAnnotationSquiggly{}
	_abdg, _eecg := _gdgc.newPdfAnnotationMarkupFromDict(_eegc)
	if _eecg != nil {
		return nil, _eecg
	}
	_fcfdd.PdfAnnotationMarkup = _abdg
	_fcfdd.QuadPoints = _eegc.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_fcfdd, nil
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_fgef *PdfFunctionType2) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_eg.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_dfffa := []float64{0.0}
	if _fgef.C0 != nil {
		_dfffa = _fgef.C0
	}
	_gbeef := []float64{1.0}
	if _fgef.C1 != nil {
		_gbeef = _fgef.C1
	}
	var _abfab []float64
	for _ffgce := 0; _ffgce < len(_dfffa); _ffgce++ {
		_fcgg := _dfffa[_ffgce] + _cbg.Pow(x[0], _fgef.N)*(_gbeef[_ffgce]-_dfffa[_ffgce])
		_abfab = append(_abfab, _fcgg)
	}
	return _abfab, nil
}
func (_cagfb *PdfAcroForm) fill(_ccbfd FieldValueProvider, _bdddeb FieldAppearanceGenerator) error {
	if _cagfb == nil {
		return nil
	}
	_ecgbe, _fbaad := _ccbfd.FieldValues()
	if _fbaad != nil {
		return _fbaad
	}
	for _, _bbfe := range _cagfb.AllFields() {
		_gfabe := _bbfe.PartialName()
		_dfada, _ggdbb := _ecgbe[_gfabe]
		if !_ggdbb {
			if _dfde, _aafee := _bbfe.FullName(); _aafee == nil {
				_dfada, _ggdbb = _ecgbe[_dfde]
			}
		}
		if !_ggdbb {
			_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _gfabe)
			continue
		}
		if _agbbd := _abeb(_bbfe, _dfada); _agbbd != nil {
			return _agbbd
		}
		if _bdddeb == nil {
			continue
		}
		for _, _bbaba := range _bbfe.Annotations {
			_cbdf, _fggb := _bdddeb.GenerateAppearanceDict(_cagfb, _bbfe, _bbaba)
			if _fggb != nil {
				return _fggb
			}
			_bbaba.AP = _cbdf
			_bbaba.ToPdfObject()
		}
	}
	return nil
}

// NewPdfActionLaunch returns a new "launch" action.
func NewPdfActionLaunch() *PdfActionLaunch {
	_gbe := NewPdfAction()
	_cbf := &PdfActionLaunch{}
	_cbf.PdfAction = _gbe
	_gbe.SetContext(_cbf)
	return _cbf
}
func (_febaa *PdfWriter) writeOutlines() error {
	if _febaa._bcbee == nil {
		return nil
	}
	_eg.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0054\u0072\u0065\u0065\u003a\u0020\u0025\u002b\u0076", _febaa._bcbee)
	_efbgd := _febaa._bcbee.ToPdfObject()
	_eg.Log.Trace("\u004fu\u0074\u006c\u0069\u006e\u0065\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0025\u0054\u002c\u0020\u0070\u003a\u0025\u0070\u0029", _efbgd, _efbgd, _efbgd)
	_febaa._dffegd.Set("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073", _efbgd)
	_dcfbf := _febaa.addObjects(_efbgd)
	if _dcfbf != nil {
		return _dcfbf
	}
	return nil
}
func (_afebc *PdfReader) newPdfSignatureReferenceFromDict(_bebed *_ebb.PdfObjectDictionary) (*PdfSignatureReference, error) {
	if _acfae, _gdce := _afebc._abbaca.GetModelFromPrimitive(_bebed).(*PdfSignatureReference); _gdce {
		return _acfae, nil
	}
	_faff := &PdfSignatureReference{_afacea: _bebed, Data: _bebed.Get("\u0044\u0061\u0074\u0061")}
	var _adbf bool
	_faff.Type, _ = _ebb.GetName(_bebed.Get("\u0054\u0079\u0070\u0065"))
	_faff.TransformMethod, _adbf = _ebb.GetName(_bebed.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"))
	if !_adbf {
		_eg.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0072\u0061\u006e\u0073\u0066o\u0072\u006dM\u0065\u0074h\u006f\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020in\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020m\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_faff.TransformParams, _ = _ebb.GetDict(_bebed.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"))
	_faff.DigestMethod, _ = _ebb.GetName(_bebed.Get("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064"))
	return _faff, nil
}
func (_gdfb *PdfWriter) setCatalogVersion() {
	_gdfb._dffegd.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _ebb.MakeName(_bg.Sprintf("\u0025\u0064\u002e%\u0064", _gdfb._efcge.Major, _gdfb._efcge.Minor)))
}

// GetCharMetrics returns the character metrics for the specified character code.  A bool flag is
// returned to indicate whether or not the entry was found in the glyph to charcode mapping.
// How it works:
//  1) Return a value the /Widths array (charWidths) if there is one.
//  2) If the font has the same name as a standard 14 font then return width=250.
//  3) Otherwise return no match and let the caller substitute a default.
func (_cfebe pdfFontSimple) GetCharMetrics(code _da.CharCode) (_bad.CharMetrics, bool) {
	if _ggdf, _ebddc := _cfebe._cdff[code]; _ebddc {
		return _bad.CharMetrics{Wx: _ggdf}, true
	}
	if _bad.IsStdFont(_bad.StdFontName(_cfebe._fdacg)) {
		return _bad.CharMetrics{Wx: 250}, true
	}
	return _bad.CharMetrics{}, false
}
func (_dage *PdfReader) newPdfAnnotationPopupFromDict(_ege *_ebb.PdfObjectDictionary) (*PdfAnnotationPopup, error) {
	_def := PdfAnnotationPopup{}
	_def.Parent = _ege.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	_def.Open = _ege.Get("\u004f\u0070\u0065\u006e")
	return &_def, nil
}

// ToPdfObject implements interface PdfModel.
func (_faab *PdfActionSound) ToPdfObject() _ebb.PdfObject {
	_faab.PdfAction.ToPdfObject()
	_ffa := _faab._abe
	_eeg := _ffa.PdfObject.(*_ebb.PdfObjectDictionary)
	_eeg.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeSound)))
	_eeg.SetIfNotNil("\u0053\u006f\u0075n\u0064", _faab.Sound)
	_eeg.SetIfNotNil("\u0056\u006f\u006c\u0075\u006d\u0065", _faab.Volume)
	_eeg.SetIfNotNil("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073", _faab.Synchronous)
	_eeg.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _faab.Repeat)
	_eeg.SetIfNotNil("\u004d\u0069\u0078", _faab.Mix)
	return _ffa
}

// SetPageLabels sets the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_acbfd *PdfWriter) SetPageLabels(pageLabels _ebb.PdfObject) error {
	if pageLabels == nil {
		return nil
	}
	_eg.Log.Trace("\u0053\u0065t\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006cs.\u002e\u002e")
	_acbfd._dffegd.Set("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073", pageLabels)
	return _acbfd.addObjects(pageLabels)
}
func _cdbaf(_ccaea *_ebb.PdfObjectDictionary) (*PdfShadingPattern, error) {
	_cabac := &PdfShadingPattern{}
	_cfaf := _ccaea.Get("\u0053h\u0061\u0064\u0069\u006e\u0067")
	if _cfaf == nil {
		_eg.Log.Debug("\u0053h\u0061d\u0069\u006e\u0067\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_beabd, _cffgf := _ggdfc(_cfaf)
	if _cffgf != nil {
		_eg.Log.Debug("\u0045r\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006eg\u0020s\u0068a\u0064\u0069\u006e\u0067\u003a\u0020\u0025v", _cffgf)
		return nil, _cffgf
	}
	_cabac.Shading = _beabd
	if _egdfe := _ccaea.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _egdfe != nil {
		_geadb, _acafe := _egdfe.(*_ebb.PdfObjectArray)
		if !_acafe {
			_eg.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _egdfe)
			return nil, _ebb.ErrTypeError
		}
		_cabac.Matrix = _geadb
	}
	if _cgfca := _ccaea.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _cgfca != nil {
		_cabac.ExtGState = _cgfca
	}
	return _cabac, nil
}
func (_fbf *PdfReader) newPdfActionResetFormFromDict(_dee *_ebb.PdfObjectDictionary) (*PdfActionResetForm, error) {
	return &PdfActionResetForm{Fields: _dee.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _dee.Get("\u0046\u006c\u0061g\u0073")}, nil
}
func (_cdbfd *pdfCIDFontType2) getFontDescriptor() *PdfFontDescriptor { return _cdbfd._fbbd }

// ToPdfObject returns the PDF representation of the colorspace.
func (_egce *PdfColorspaceSpecialPattern) ToPdfObject() _ebb.PdfObject {
	if _egce.UnderlyingCS == nil {
		return _ebb.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e")
	}
	_gafb := _ebb.MakeArray(_ebb.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_gafb.Append(_egce.UnderlyingCS.ToPdfObject())
	if _egce._adegb != nil {
		_egce._adegb.PdfObject = _gafb
		return _egce._adegb
	}
	return _gafb
}

// NewPdfActionHide returns a new "hide" action.
func NewPdfActionHide() *PdfActionHide {
	_bfd := NewPdfAction()
	_db := &PdfActionHide{}
	_db.PdfAction = _bfd
	_bfd.SetContext(_db)
	return _db
}
func (_acea *PdfWriter) copyObjects() {
	_abgga := make(map[_ebb.PdfObject]_ebb.PdfObject)
	_fbacc := make([]_ebb.PdfObject, 0, len(_acea._ebdgg))
	_efedf := make(map[_ebb.PdfObject]struct{}, len(_acea._ebdgg))
	_badbff := make(map[_ebb.PdfObject]struct{})
	for _, _cfaae := range _acea._ebdgg {
		_cefecg := _acea.copyObject(_cfaae, _abgga, _badbff, false)
		if _, _efdba := _badbff[_cfaae]; _efdba {
			continue
		}
		_fbacc = append(_fbacc, _cefecg)
		_efedf[_cefecg] = struct{}{}
	}
	_acea._ebdgg = _fbacc
	_acea._ffffd = _efedf
	_acea._eadfd = _acea.copyObject(_acea._eadfd, _abgga, nil, false).(*_ebb.PdfIndirectObject)
	_acea._gegba = _acea.copyObject(_acea._gegba, _abgga, nil, false).(*_ebb.PdfIndirectObject)
	if _acea._cbcaa != nil {
		_acea._cbcaa = _acea.copyObject(_acea._cbcaa, _abgga, nil, false).(*_ebb.PdfIndirectObject)
	}
	if _acea._abffb {
		_bdddc := make(map[_ebb.PdfObject]int64)
		for _gdcbe, _cbdba := range _acea._cdgd {
			if _gggg, _fadbed := _abgga[_gdcbe]; _fadbed {
				_bdddc[_gggg] = _cbdba
			} else {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020a\u0070\u0070\u0065n\u0064\u0020\u006d\u006fd\u0065\u0020\u002d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u0070\u0079\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070")
			}
		}
		_acea._cdgd = _bdddc
	}
}

// WriteToFile writes the output PDF to file.
func (_efdcb *PdfWriter) WriteToFile(outputFilePath string) error {
	_effc, _cbcaf := _ed.Create(outputFilePath)
	if _cbcaf != nil {
		return _cbcaf
	}
	defer _effc.Close()
	return _efdcb.Write(_effc)
}
func _bggeg() string {
	_daddc.Lock()
	defer _daddc.Unlock()
	return _fead
}

// ToPdfObject sets the common field elements.
// Note: Call the more field context's ToPdfObject to set both the generic and
// non-generic information.
func (_cgca *PdfField) ToPdfObject() _ebb.PdfObject {
	_cgee := _cgca._cdfd
	_badcb := _cgee.PdfObject.(*_ebb.PdfObjectDictionary)
	_fgdd := _ebb.MakeArray()
	for _, _dacf := range _cgca.Kids {
		_fgdd.Append(_dacf.ToPdfObject())
	}
	for _, _abgd := range _cgca.Annotations {
		if _abgd._bdcd != _cgca._cdfd {
			_fgdd.Append(_abgd.GetContext().ToPdfObject())
		}
	}
	if _cgca.Parent != nil {
		_badcb.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _cgca.Parent.GetContainingPdfObject())
	}
	if _fgdd.Len() > 0 {
		_badcb.Set("\u004b\u0069\u0064\u0073", _fgdd)
	}
	_badcb.SetIfNotNil("\u0046\u0054", _cgca.FT)
	_badcb.SetIfNotNil("\u0054", _cgca.T)
	_badcb.SetIfNotNil("\u0054\u0055", _cgca.TU)
	_badcb.SetIfNotNil("\u0054\u004d", _cgca.TM)
	_badcb.SetIfNotNil("\u0046\u0066", _cgca.Ff)
	_badcb.SetIfNotNil("\u0056", _cgca.V)
	_badcb.SetIfNotNil("\u0044\u0056", _cgca.DV)
	_badcb.SetIfNotNil("\u0041\u0041", _cgca.AA)
	if _cgca.VariableText != nil {
		_badcb.SetIfNotNil("\u0044\u0041", _cgca.VariableText.DA)
		_badcb.SetIfNotNil("\u0051", _cgca.VariableText.Q)
		_badcb.SetIfNotNil("\u0044\u0053", _cgca.VariableText.DS)
		_badcb.SetIfNotNil("\u0052\u0056", _cgca.VariableText.RV)
	}
	return _cgee
}

// SetBorderWidth sets the style's border width.
func (_daba *PdfBorderStyle) SetBorderWidth(width float64) { _daba.W = &width }

// FillWithAppearance populates `form` with values provided by `provider`.
// If not nil, `appGen` is used to generate appearance dictionaries for the
// field annotations, based on the specified settings. Otherwise, appearance
// generation is skipped.
// e.g.: appGen := annotator.FieldAppearance{OnlyIfMissing: true, RegenerateTextFields: true}
// NOTE: In next major version this functionality will be part of Fill. (v4)
func (_aebac *PdfAcroForm) FillWithAppearance(provider FieldValueProvider, appGen FieldAppearanceGenerator) error {
	_bggab := _aebac.fill(provider, appGen)
	if _bggab != nil {
		return _bggab
	}
	if _, _bdafc := provider.(FieldImageProvider); _bdafc {
		_bggab = _aebac.fillImageWithAppearance(provider.(FieldImageProvider), appGen)
	}
	return _bggab
}

// PdfActionTrans represents a trans action.
type PdfActionTrans struct {
	*PdfAction
	Trans _ebb.PdfObject
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_afgdb *PdfOutline) ToPdfObject() _ebb.PdfObject {
	_ceeac := _afgdb._egee
	_beede := _ceeac.PdfObject.(*_ebb.PdfObjectDictionary)
	_beede.Set("\u0054\u0079\u0070\u0065", _ebb.MakeName("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073"))
	if _afgdb.First != nil {
		_beede.Set("\u0046\u0069\u0072s\u0074", _afgdb.First.ToPdfObject())
	}
	if _afgdb.Last != nil {
		_beede.Set("\u004c\u0061\u0073\u0074", _afgdb.Last.GetContext().GetContainingPdfObject())
	}
	if _afgdb.Parent != nil {
		_beede.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _afgdb.Parent.GetContext().GetContainingPdfObject())
	}
	if _afgdb.Count != nil {
		_beede.Set("\u0043\u006f\u0075n\u0074", _ebb.MakeInteger(*_afgdb.Count))
	}
	return _ceeac
}

// NewPdfOutputIntentFromPdfObject creates a new PdfOutputIntent from the input core.PdfObject.
func NewPdfOutputIntentFromPdfObject(object _ebb.PdfObject) (*PdfOutputIntent, error) {
	_befc := &PdfOutputIntent{}
	if _beeeb := _befc.ParsePdfObject(object); _beeeb != nil {
		return nil, _beeeb
	}
	return _befc, nil
}

// GetSubFilter returns SubFilter value or empty string.
func (_fgaeg *pdfSignDictionary) GetSubFilter() string {
	_eebbf := _fgaeg.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r")
	if _eebbf == nil {
		return ""
	}
	if _fcgcg, _adgfc := _ebb.GetNameVal(_eebbf); _adgfc {
		return _fcgcg
	}
	return ""
}

// FieldValueProvider provides field values from a data source such as FDF, JSON or any other.
type FieldValueProvider interface {
	FieldValues() (map[string]_ebb.PdfObject, error)
}

func (_cff *PdfReader) newPdfActionGotoEFromDict(_fce *_ebb.PdfObjectDictionary) (*PdfActionGoToE, error) {
	_gee, _fda := _gggf(_fce.Get("\u0046"))
	if _fda != nil {
		return nil, _fda
	}
	return &PdfActionGoToE{D: _fce.Get("\u0044"), NewWindow: _fce.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), T: _fce.Get("\u0054"), F: _gee}, nil
}
func (_dega *PdfReader) newPdfAnnotationCaretFromDict(_cgc *_ebb.PdfObjectDictionary) (*PdfAnnotationCaret, error) {
	_cffe := PdfAnnotationCaret{}
	_aabg, _bcfd := _dega.newPdfAnnotationMarkupFromDict(_cgc)
	if _bcfd != nil {
		return nil, _bcfd
	}
	_cffe.PdfAnnotationMarkup = _aabg
	_cffe.RD = _cgc.Get("\u0052\u0044")
	_cffe.Sy = _cgc.Get("\u0053\u0079")
	return &_cffe, nil
}

// NewPdfActionGoToR returns a new "go to remote" action.
func NewPdfActionGoToR() *PdfActionGoToR {
	_gfd := NewPdfAction()
	_cgg := &PdfActionGoToR{}
	_cgg.PdfAction = _gfd
	_gfd.SetContext(_cgg)
	return _cgg
}
func (_bgfa *PdfReader) newPdfActionURIFromDict(_fbe *_ebb.PdfObjectDictionary) (*PdfActionURI, error) {
	return &PdfActionURI{URI: _fbe.Get("\u0055\u0052\u0049"), IsMap: _fbe.Get("\u0049\u0073\u004da\u0070")}, nil
}
func (_ecc *PdfReader) newPdfActionGoTo3DViewFromDict(_dacb *_ebb.PdfObjectDictionary) (*PdfActionGoTo3DView, error) {
	return &PdfActionGoTo3DView{TA: _dacb.Get("\u0054\u0041"), V: _dacb.Get("\u0056")}, nil
}

// ToPdfObject returns a PdfObject representation of PdfColorspaceDeviceNAttributes as a PdfObjectDictionary directly
// or indirectly within an indirect object container.
func (_gccb *PdfColorspaceDeviceNAttributes) ToPdfObject() _ebb.PdfObject {
	_afgbb := _ebb.MakeDict()
	if _gccb.Subtype != nil {
		_afgbb.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _gccb.Subtype)
	}
	_afgbb.SetIfNotNil("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s", _gccb.Colorants)
	_afgbb.SetIfNotNil("\u0050r\u006f\u0063\u0065\u0073\u0073", _gccb.Process)
	_afgbb.SetIfNotNil("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073", _gccb.MixingHints)
	if _gccb._afca != nil {
		_gccb._afca.PdfObject = _afgbb
		return _gccb._afca
	}
	return _afgbb
}

// GetContainingPdfObject returns the container of the outline item (indirect object).
func (_dedce *PdfOutlineItem) GetContainingPdfObject() _ebb.PdfObject { return _dedce._cacdf }

// NewXObjectImageFromImage creates a new XObject Image from an image object
// with default options. If encoder is nil, uses raw encoding (none).
func NewXObjectImageFromImage(img *Image, cs PdfColorspace, encoder _ebb.StreamEncoder) (*XObjectImage, error) {
	_gebce := NewXObjectImage()
	return UpdateXObjectImageFromImage(_gebce, img, cs, encoder)
}

// Initialize initializes the PdfSignature.
func (_ebefb *PdfSignature) Initialize() error {
	if _ebefb.Handler == nil {
		return _gf.New("\u0073\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0068\u0061\u006e\u0064\u006c\u0065r\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	return _ebefb.Handler.InitSignature(_ebefb)
}
func _fbga(_eeeee _ebb.PdfObject, _ggfbb *PdfReader) (*OutlineDest, error) {
	_ffagd, _beceab := _ebb.GetArray(_eeeee)
	if !_beceab {
		return nil, _gf.New("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0064\u0065\u0073t\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_aagd := _ffagd.Len()
	if _aagd < 2 {
		return nil, _bg.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u006fu\u0074\u006c\u0069n\u0065\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061ti\u006f\u006e\u0020a\u0072\u0072a\u0079\u0020\u006c\u0065\u006e\u0067t\u0068\u003a \u0025\u0064", _aagd)
	}
	_daaf := &OutlineDest{Mode: "\u0046\u0069\u0074"}
	_dccg := _ffagd.Get(0)
	if _dfac, _baaf := _ebb.GetIndirect(_dccg); _baaf {
		if _, _dcce, _cbdfg := _ggfbb.PageFromIndirectObject(_dfac); _cbdfg == nil {
			_daaf.Page = int64(_dcce - 1)
		} else {
			_eg.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020g\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u002b\u0076", _dfac)
		}
		_daaf.PageObj = _dfac
	} else if _eeebb, _gggca := _ebb.GetIntVal(_dccg); _gggca {
		if _eeebb >= 0 && _eeebb < len(_ggfbb.PageList) {
			_daaf.PageObj = _ggfbb.PageList[_eeebb].GetPageAsIndirectObject()
		} else {
			_eg.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _eeebb)
		}
		_daaf.Page = int64(_eeebb)
	} else {
		return nil, _bg.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006f\u0075\u0074\u006cine\u0020de\u0073\u0074\u0069\u006e\u0061\u0074\u0069on\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0054", _dccg)
	}
	_afeg, _beceab := _ebb.GetNameVal(_ffagd.Get(1))
	if !_beceab {
		_eg.Log.Debug("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006d\u0061\u0067\u006e\u0069\u0066\u0069\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _ffagd.Get(1))
		return _daaf, nil
	}
	switch _afeg {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		if _aagd > 2 {
			_daaf.Y, _ = _ebb.GetNumberAsFloat(_ebb.TraceToDirectObject(_ffagd.Get(2)))
		}
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		if _aagd > 2 {
			_daaf.X, _ = _ebb.GetNumberAsFloat(_ebb.TraceToDirectObject(_ffagd.Get(2)))
		}
	case "\u0058\u0059\u005a":
		if _aagd > 4 {
			_daaf.X, _ = _ebb.GetNumberAsFloat(_ebb.TraceToDirectObject(_ffagd.Get(2)))
			_daaf.Y, _ = _ebb.GetNumberAsFloat(_ebb.TraceToDirectObject(_ffagd.Get(3)))
			_daaf.Zoom, _ = _ebb.GetNumberAsFloat(_ebb.TraceToDirectObject(_ffagd.Get(4)))
		}
	default:
		_afeg = "\u0046\u0069\u0074"
	}
	_daaf.Mode = _afeg
	return _daaf, nil
}

// SetReason sets the `Reason` field of the signature.
func (_cbfga *PdfSignature) SetReason(reason string) { _cbfga.Reason = _ebb.MakeString(reason) }

// NewStandard14FontWithEncoding returns the standard 14 font named `basefont` as a *PdfFont and
// a TextEncoder that encodes all the runes in `alphabet`, or an error if this is not possible.
// An error can occur if `basefont` is not one the standard 14 font names.
func NewStandard14FontWithEncoding(basefont StdFontName, alphabet map[rune]int) (*PdfFont, _da.SimpleEncoder, error) {
	_ceag, _ageb := _cfacd(basefont)
	if _ageb != nil {
		return nil, nil, _ageb
	}
	_aebaa, _cefd := _ceag.Encoder().(_da.SimpleEncoder)
	if !_cefd {
		return nil, nil, _bg.Errorf("\u006f\u006e\u006c\u0079\u0020s\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u002c\u0020\u0067\u006f\u0074\u0020\u0025\u0054", _ceag.Encoder())
	}
	_abfbg := make(map[rune]_da.GlyphName)
	for _fgcccc := range alphabet {
		if _, _ggged := _aebaa.RuneToCharcode(_fgcccc); !_ggged {
			_, _aaec := _ceag._ddgd.Read(_fgcccc)
			if !_aaec {
				_eg.Log.Trace("r\u0075\u006e\u0065\u0020\u0025\u0023x\u003d\u0025\u0071\u0020\u006e\u006f\u0074\u0020\u0069n\u0020\u0074\u0068e\u0020f\u006f\u006e\u0074", _fgcccc, _fgcccc)
				continue
			}
			_dgca, _aaec := _da.RuneToGlyph(_fgcccc)
			if !_aaec {
				_eg.Log.Debug("\u006eo\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u0020\u0025\u0023\u0078\u003d\u0025\u0071", _fgcccc, _fgcccc)
				continue
			}
			if len(_abfbg) >= 255 {
				return nil, nil, _gf.New("\u0074\u006f\u006f\u0020\u006d\u0061\u006e\u0079\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020f\u006f\u0072\u0020\u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067")
			}
			_abfbg[_fgcccc] = _dgca
		}
	}
	var (
		_fgaeb []_da.CharCode
		_cbgc  []_da.CharCode
	)
	for _bfee := _da.CharCode(1); _bfee <= 0xff; _bfee++ {
		_cgcaa, _fcag := _aebaa.CharcodeToRune(_bfee)
		if !_fcag {
			_fgaeb = append(_fgaeb, _bfee)
			continue
		}
		if _, _fcag = alphabet[_cgcaa]; !_fcag {
			_cbgc = append(_cbgc, _bfee)
		}
	}
	_bcbf := append(_fgaeb, _cbgc...)
	if len(_bcbf) < len(_abfbg) {
		return nil, nil, _bg.Errorf("n\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u006ec\u006f\u0064\u0065\u0020\u0025\u0064\u0020ru\u006e\u0065\u0073\u002c \u0062\u0075\u0074\u0020\u0068\u0061\u0076\u0065\u0020on\u006c\u0079 \u0025\u0064\u0020\u0073\u006c\u006f\u0074\u0073", len(_abfbg), len(_bcbf))
	}
	_edga := make([]rune, 0, len(_abfbg))
	for _aaccd := range _abfbg {
		_edga = append(_edga, _aaccd)
	}
	_ae.Slice(_edga, func(_gcb, _fega int) bool { return _edga[_gcb] < _edga[_fega] })
	_eadb := make(map[_da.CharCode]_da.GlyphName, len(_edga))
	for _, _gdbbd := range _edga {
		_gbefg := _bcbf[0]
		_bcbf = _bcbf[1:]
		_eadb[_gbefg] = _abfbg[_gdbbd]
	}
	_aebaa = _da.ApplyDifferences(_aebaa, _eadb)
	_ceag.SetEncoder(_aebaa)
	return &PdfFont{_ebcad: &_ceag}, _aebaa, nil
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
	DiffResults *_ac.DiffResults

	// GeneralizedTime is the time at which the time-stamp token has been created by the TSA (RFC 3161).
	GeneralizedTime _f.Time
}

// SetOpenAction sets the OpenAction in the PDF catalog.
// The value shall be either an array defining a destination (12.3.2 "Destinations" PDF32000_2008),
// or an action dictionary representing an action (12.6 "Actions" PDF32000_2008).
func (_beceb *PdfWriter) SetOpenAction(dest _ebb.PdfObject) error {
	if dest == nil || _ebb.IsNullObject(dest) {
		return nil
	}
	_beceb._dffegd.Set("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e", dest)
	return _beceb.addObjects(dest)
}
func _gagfgf() string {
	_daddc.Lock()
	defer _daddc.Unlock()
	if len(_aaaae) > 0 {
		return _aaaae
	}
	return "Go PDF"
}
func _bede(_agfe *fontCommon) *pdfFontSimple { return &pdfFontSimple{fontCommon: *_agfe} }

// NewStandard14Font returns the standard 14 font named `basefont` as a *PdfFont, or an error if it
// `basefont` is not one of the standard 14 font names.
func NewStandard14Font(basefont StdFontName) (*PdfFont, error) {
	_ceef, _fdbdf := _cfacd(basefont)
	if _fdbdf != nil {
		return nil, _fdbdf
	}
	if basefont != SymbolName && basefont != ZapfDingbatsName {
		_ceef._ebcb = _da.NewWinAnsiEncoder()
	}
	return &PdfFont{_ebcad: &_ceef}, nil
}

// PdfActionResetForm represents a resetForm action.
type PdfActionResetForm struct {
	*PdfAction
	Fields _ebb.PdfObject
	Flags  _ebb.PdfObject
}

// Size returns the width and the height of the page. The method reports
// the page dimensions as displayed by a PDF viewer (i.e. page rotation is
// taken into account).
func (_cbdfgg *PdfPage) Size() (float64, float64, error) {
	_ffeff, _afdff := _cbdfgg.GetMediaBox()
	if _afdff != nil {
		return 0, 0, _afdff
	}
	_fcgfd, _cgbbc := _ffeff.Width(), _ffeff.Height()
	_fdga, _afdff := _cbdfgg.GetRotate()
	if _afdff != nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _afdff.Error())
	}
	if _cdba := _fdga; _cdba%360 != 0 && _cdba%90 == 0 {
		if _ebeab := (360 + _cdba%360) % 360; _ebeab == 90 || _ebeab == 270 {
			_fcgfd, _cgbbc = _cgbbc, _fcgfd
		}
	}
	return _fcgfd, _cgbbc, nil
}

// Sign signs a specific page with a digital signature.
// The signature field parameter must have a valid signature dictionary
// specified by its V field.
func (_bceg *PdfAppender) Sign(pageNum int, field *PdfFieldSignature) error {
	if field == nil {
		return _gf.New("\u0073\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 n\u0069\u006c")
	}
	_fcd := field.V
	if _fcd == nil {
		return _gf.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061r\u0079 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_ggbec := pageNum - 1
	if _ggbec < 0 || _ggbec > len(_bceg._dfbg)-1 {
		return _bg.Errorf("\u0070\u0061\u0067\u0065\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064", pageNum)
	}
	_fded := _bceg.Reader.PageList[_ggbec]
	field.P = _fded.ToPdfObject()
	if field.T == nil || field.T.String() == "" {
		field.T = _ebb.MakeString(_bg.Sprintf("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072\u0065\u0020\u0025\u0064", pageNum))
	}
	_fded.AddAnnotation(field.PdfAnnotationWidget.PdfAnnotation)
	if _bceg._bfef == _bceg._acfe.AcroForm {
		_bceg._bfef = _bceg.Reader.AcroForm
	}
	_eeddb := _bceg._bfef
	if _eeddb == nil {
		_eeddb = NewPdfAcroForm()
	}
	_eeddb.SigFlags = _ebb.MakeInteger(3)
	_dgadf := append(_eeddb.AllFields(), field.PdfField)
	_eeddb.Fields = &_dgadf
	_bceg.ReplaceAcroForm(_eeddb)
	_bceg.UpdatePage(_fded)
	_bceg._dfbg[_ggbec] = _fded
	if _, _afgb := field.V.GetDocMDPPermission(); _afgb {
		_bceg._eaaa = NewPermissions(field.V)
	}
	return nil
}

// NewPdfAnnotationWatermark returns a new watermark annotation.
func NewPdfAnnotationWatermark() *PdfAnnotationWatermark {
	_abab := NewPdfAnnotation()
	_ebfa := &PdfAnnotationWatermark{}
	_ebfa.PdfAnnotation = _abab
	_abab.SetContext(_ebfa)
	return _ebfa
}

// PdfAnnotation3D represents 3D annotations.
// (Section 13.6.2).
type PdfAnnotation3D struct {
	*PdfAnnotation
	T3DD _ebb.PdfObject
	T3DV _ebb.PdfObject
	T3DA _ebb.PdfObject
	T3DI _ebb.PdfObject
	T3DB _ebb.PdfObject
}

// NewPdfAnnotationInk returns a new ink annotation.
func NewPdfAnnotationInk() *PdfAnnotationInk {
	_dggg := NewPdfAnnotation()
	_efba := &PdfAnnotationInk{}
	_efba.PdfAnnotation = _dggg
	_efba.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dggg.SetContext(_efba)
	return _efba
}

// GetOutlineTree returns the outline tree.
func (_faef *PdfReader) GetOutlineTree() *PdfOutlineTreeNode { return _faef._fgbcg }

// NewPdfActionSetOCGState returns a new "named" action.
func NewPdfActionSetOCGState() *PdfActionSetOCGState {
	_cfa := NewPdfAction()
	_gdb := &PdfActionSetOCGState{}
	_gdb.PdfAction = _cfa
	_cfa.SetContext(_gdb)
	return _gdb
}

// GetOutlinesFlattened returns a flattened list of tree nodes and titles.
// NOTE: for most use cases, it is recommended to use the high-level GetOutlines
// method instead, which also provides information regarding the destination
// of the outline items.
func (_eddfb *PdfReader) GetOutlinesFlattened() ([]*PdfOutlineTreeNode, []string, error) {
	var _caagc []*PdfOutlineTreeNode
	var _bffbb []string
	var _bbfgd func(*PdfOutlineTreeNode, *[]*PdfOutlineTreeNode, *[]string, int)
	_bbfgd = func(_dcda *PdfOutlineTreeNode, _afce *[]*PdfOutlineTreeNode, _dgfc *[]string, _gecdba int) {
		if _dcda == nil {
			return
		}
		if _dcda._geeee == nil {
			_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020M\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006e\u006fd\u0065\u002e\u0063o\u006et\u0065\u0078\u0074")
			return
		}
		_gfagf, _ddfdc := _dcda._geeee.(*PdfOutlineItem)
		if _ddfdc {
			*_afce = append(*_afce, &_gfagf.PdfOutlineTreeNode)
			_eggf := _ee.Repeat("\u0020", _gecdba*2) + _gfagf.Title.Decoded()
			*_dgfc = append(*_dgfc, _eggf)
		}
		if _dcda.First != nil {
			_eegceb := _ee.Repeat("\u0020", _gecdba*2) + "\u002b"
			*_dgfc = append(*_dgfc, _eegceb)
			_bbfgd(_dcda.First, _afce, _dgfc, _gecdba+1)
		}
		if _ddfdc && _gfagf.Next != nil {
			_bbfgd(_gfagf.Next, _afce, _dgfc, _gecdba)
		}
	}
	_bbfgd(_eddfb._fgbcg, &_caagc, &_bffbb, 0)
	return _caagc, _bffbb, nil
}

// IsValid checks if the given pdf output intent type is valid.
func (_dffge PdfOutputIntentType) IsValid() bool {
	return _dffge >= PdfOutputIntentTypeA1 && _dffge <= PdfOutputIntentTypeX
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_cgde pdfCIDFontType2) GetRuneMetrics(r rune) (_bad.CharMetrics, bool) {
	_ebbd, _baeed := _cgde._dceb[r]
	if !_baeed {
		_caaa, _acaa := _ebb.GetInt(_cgde.DW)
		if !_acaa {
			return _bad.CharMetrics{}, false
		}
		_ebbd = int(*_caaa)
	}
	return _bad.CharMetrics{Wx: float64(_ebbd)}, true
}

// PdfRectangle is a definition of a rectangle.
type PdfRectangle struct {
	Llx float64
	Lly float64
	Urx float64
	Ury float64
}

const (
	BorderEffectNoEffect BorderEffect = iota
	BorderEffectCloudy   BorderEffect = iota
)

var _efdc = map[string]struct{}{"\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074\u0045\u006e\u0063o\u0064\u0069\u006e\u0067": {}, "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}}

// GetContainingPdfObject returns the containing object for the PdfField, i.e. an indirect object
// containing the field dictionary.
func (_ceaf *PdfField) GetContainingPdfObject() _ebb.PdfObject { return _ceaf._cdfd }

// ToPdfObject implements interface PdfModel.
func (_dedb *PdfAnnotationHighlight) ToPdfObject() _ebb.PdfObject {
	_dedb.PdfAnnotation.ToPdfObject()
	_ggaa := _dedb._bdcd
	_fbbf := _ggaa.PdfObject.(*_ebb.PdfObjectDictionary)
	_dedb.PdfAnnotationMarkup.appendToPdfDictionary(_fbbf)
	_fbbf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t"))
	_fbbf.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _dedb.QuadPoints)
	return _ggaa
}
func (_fccf *PdfReader) newPdfAnnotationWatermarkFromDict(_cgff *_ebb.PdfObjectDictionary) (*PdfAnnotationWatermark, error) {
	_gdbb := PdfAnnotationWatermark{}
	_gdbb.FixedPrint = _cgff.Get("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074")
	return &_gdbb, nil
}
func (_cgac *PdfReader) newPdfAnnotationSoundFromDict(_bbee *_ebb.PdfObjectDictionary) (*PdfAnnotationSound, error) {
	_cdcg := PdfAnnotationSound{}
	_gbc, _efed := _cgac.newPdfAnnotationMarkupFromDict(_bbee)
	if _efed != nil {
		return nil, _efed
	}
	_cdcg.PdfAnnotationMarkup = _gbc
	_cdcg.Name = _bbee.Get("\u004e\u0061\u006d\u0065")
	_cdcg.Sound = _bbee.Get("\u0053\u006f\u0075n\u0064")
	return &_cdcg, nil
}
func _bdgcb(_ecged *PdfAnnotation) (*XObjectForm, *PdfRectangle, error) {
	_dafbf, _gebaag := _ebb.GetDict(_ecged.AP)
	if !_gebaag {
		return nil, nil, _gf.New("f\u0069\u0065\u006c\u0064\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0041\u0050\u0020d\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079")
	}
	if _dafbf == nil {
		return nil, nil, nil
	}
	_gccf, _gebaag := _ebb.GetArray(_ecged.Rect)
	if !_gebaag || _gccf.Len() != 4 {
		return nil, nil, _gf.New("\u0072\u0065\u0063t\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_bgbgc, _cgfd := NewPdfRectangle(*_gccf)
	if _cgfd != nil {
		return nil, nil, _cgfd
	}
	_decg := _ebb.TraceToDirectObject(_dafbf.Get("\u004e"))
	switch _dbcgd := _decg.(type) {
	case *_ebb.PdfObjectStream:
		_gfagd := _dbcgd
		_egbc, _beee := NewXObjectFormFromStream(_gfagd)
		return _egbc, _bgbgc, _beee
	case *_ebb.PdfObjectDictionary:
		_gbcfe := _dbcgd
		_gecea, _abge := _ebb.GetName(_ecged.AS)
		if !_abge {
			return nil, nil, nil
		}
		if _gbcfe.Get(*_gecea) == nil {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0041\u0053\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
			return nil, nil, nil
		}
		_cace, _abge := _ebb.GetStream(_gbcfe.Get(*_gecea))
		if !_abge {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006ce \u0074\u006f\u0020\u0061\u0063\u0063e\u0073\u0073\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0066\u006f\u0072 \u0025\u0076", _gecea)
			return nil, nil, _gf.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		}
		_ggefc, _bcaa := NewXObjectFormFromStream(_cace)
		return _ggefc, _bgbgc, _bcaa
	}
	_eg.Log.Debug("\u0049\u006e\u0076\u0061li\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u004e\u003a\u0020%\u0054", _decg)
	return nil, nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
}

// ColorToRGB converts a CMYK32 color to an RGB color.
func (_efcaf *PdfColorspaceDeviceCMYK) ColorToRGB(color PdfColor) (PdfColor, error) {
	_gdee, _bdba := color.(*PdfColorDeviceCMYK)
	if !_bdba {
		_eg.Log.Debug("I\u006e\u0070\u0075\u0074\u0020\u0063o\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0064e\u0076\u0069\u0063e\u0020c\u006d\u0079\u006b")
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_fgbfc := _gdee.C()
	_afda := _gdee.M()
	_egbbc := _gdee.Y()
	_gffd := _gdee.K()
	_fgbfc = _fgbfc*(1-_gffd) + _gffd
	_afda = _afda*(1-_gffd) + _gffd
	_egbbc = _egbbc*(1-_gffd) + _gffd
	_cbaad := 1 - _fgbfc
	_afde := 1 - _afda
	_dfbb := 1 - _egbbc
	return NewPdfColorDeviceRGB(_cbaad, _afde, _dfbb), nil
}

// PdfColorspaceSpecialIndexed is an indexed color space is a lookup table, where the input element
// is an index to the lookup table and the output is a color defined in the lookup table in the Base
// colorspace.
// [/Indexed base hival lookup]
type PdfColorspaceSpecialIndexed struct {
	Base   PdfColorspace
	HiVal  int
	Lookup _ebb.PdfObject
	_eabg  []byte
	_gdaef *_ebb.PdfIndirectObject
}

func _eaef(_bcdb []byte) ([]byte, error) {
	_befae := _bd.New()
	if _, _cgbed := _ab.Copy(_befae, _ca.NewReader(_bcdb)); _cgbed != nil {
		return nil, _cgbed
	}
	return _befae.Sum(nil), nil
}
func _gaab(_aadcg *_ebb.PdfObjectDictionary) (*PdfShadingType7, error) {
	_ebdcad := PdfShadingType7{}
	_gaeee := _aadcg.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _gaeee == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_deefb, _efdbf := _gaeee.(*_ebb.PdfObjectInteger)
	if !_efdbf {
		_eg.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gaeee)
		return nil, _ebb.ErrTypeError
	}
	_ebdcad.BitsPerCoordinate = _deefb
	_gaeee = _aadcg.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _gaeee == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_deefb, _efdbf = _gaeee.(*_ebb.PdfObjectInteger)
	if !_efdbf {
		_eg.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _gaeee)
		return nil, _ebb.ErrTypeError
	}
	_ebdcad.BitsPerComponent = _deefb
	_gaeee = _aadcg.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _gaeee == nil {
		_eg.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_deefb, _efdbf = _gaeee.(*_ebb.PdfObjectInteger)
	if !_efdbf {
		_eg.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _gaeee)
		return nil, _ebb.ErrTypeError
	}
	_ebdcad.BitsPerComponent = _deefb
	_gaeee = _aadcg.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _gaeee == nil {
		_eg.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_bfffg, _efdbf := _gaeee.(*_ebb.PdfObjectArray)
	if !_efdbf {
		_eg.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _gaeee)
		return nil, _ebb.ErrTypeError
	}
	_ebdcad.Decode = _bfffg
	if _fdgbg := _aadcg.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _fdgbg != nil {
		_ebdcad.Function = []PdfFunction{}
		if _geac, _adcbc := _fdgbg.(*_ebb.PdfObjectArray); _adcbc {
			for _, _dgbd := range _geac.Elements() {
				_dgbfe, _gecec := _aagg(_dgbd)
				if _gecec != nil {
					_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _gecec)
					return nil, _gecec
				}
				_ebdcad.Function = append(_ebdcad.Function, _dgbfe)
			}
		} else {
			_bedfd, _affbc := _aagg(_fdgbg)
			if _affbc != nil {
				_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _affbc)
				return nil, _affbc
			}
			_ebdcad.Function = append(_ebdcad.Function, _bedfd)
		}
	}
	return &_ebdcad, nil
}

// NewPdfColorspaceSpecialIndexed returns a new Indexed color.
func NewPdfColorspaceSpecialIndexed() *PdfColorspaceSpecialIndexed {
	return &PdfColorspaceSpecialIndexed{HiVal: 255}
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_gcfa *PdfColorspaceLab) ToPdfObject() _ebb.PdfObject {
	_gdge := _ebb.MakeArray()
	_gdge.Append(_ebb.MakeName("\u004c\u0061\u0062"))
	_gbge := _ebb.MakeDict()
	if _gcfa.WhitePoint != nil {
		_daaba := _ebb.MakeArray(_ebb.MakeFloat(_gcfa.WhitePoint[0]), _ebb.MakeFloat(_gcfa.WhitePoint[1]), _ebb.MakeFloat(_gcfa.WhitePoint[2]))
		_gbge.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _daaba)
	} else {
		_eg.Log.Error("\u004c\u0061\u0062: \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057h\u0069t\u0065P\u006fi\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	if _gcfa.BlackPoint != nil {
		_egg := _ebb.MakeArray(_ebb.MakeFloat(_gcfa.BlackPoint[0]), _ebb.MakeFloat(_gcfa.BlackPoint[1]), _ebb.MakeFloat(_gcfa.BlackPoint[2]))
		_gbge.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _egg)
	}
	if _gcfa.Range != nil {
		_cfac := _ebb.MakeArray(_ebb.MakeFloat(_gcfa.Range[0]), _ebb.MakeFloat(_gcfa.Range[1]), _ebb.MakeFloat(_gcfa.Range[2]), _ebb.MakeFloat(_gcfa.Range[3]))
		_gbge.Set("\u0052\u0061\u006eg\u0065", _cfac)
	}
	_gdge.Append(_gbge)
	if _gcfa._bdga != nil {
		_gcfa._bdga.PdfObject = _gdge
		return _gcfa._bdga
	}
	return _gdge
}

// M returns the value of the magenta component of the color.
func (_cafe *PdfColorDeviceCMYK) M() float64 { return _cafe[1] }

// DefaultFont returns the default font, which is currently the built in Helvetica.
func DefaultFont() *PdfFont {
	_baeee, _ggccg := _bad.NewStdFontByName(HelveticaName)
	if !_ggccg {
		panic("\u0048\u0065lv\u0065\u0074\u0069c\u0061\u0020\u0073\u0068oul\u0064 a\u006c\u0077\u0061\u0079\u0073\u0020\u0062e \u0061\u0076\u0061\u0069\u006c\u0061\u0062l\u0065")
	}
	_ebbg := _fcgbc(_baeee)
	return &PdfFont{_ebcad: &_ebbg}
}

// PdfAnnotationLine represents Line annotations.
// (Section 12.5.6.7).
type PdfAnnotationLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	L       _ebb.PdfObject
	BS      _ebb.PdfObject
	LE      _ebb.PdfObject
	IC      _ebb.PdfObject
	LL      _ebb.PdfObject
	LLE     _ebb.PdfObject
	Cap     _ebb.PdfObject
	IT      _ebb.PdfObject
	LLO     _ebb.PdfObject
	CP      _ebb.PdfObject
	Measure _ebb.PdfObject
	CO      _ebb.PdfObject
}

// CharcodesToStrings returns the unicode strings corresponding to `charcodes`.
// The int returns are the number of strings and the number of unconvereted codes.
// NOTE: The number of strings returned is equal to the number of charcodes
func (_aedfb *PdfFont) CharcodesToStrings(charcodes []_da.CharCode) ([]string, int, int) {
	_bfab := _aedfb.baseFields()
	_gaba := make([]string, 0, len(charcodes))
	_daad := 0
	_cfaadg := _aedfb.Encoder()
	_abefa := _bfab._dcdd != nil && _aedfb.IsSimple() && _aedfb.Subtype() == "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" && !_ee.Contains(_bfab._dcdd.Name(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")
	if !_abefa && _cfaadg != nil {
		switch _cfaadc := _cfaadg.(type) {
		case _da.SimpleEncoder:
			_abcb := _cfaadc.BaseName()
			if _, _ccgbb := _efdc[_abcb]; _ccgbb {
				for _, _aada := range charcodes {
					if _bfeec, _ffecb := _cfaadg.CharcodeToRune(_aada); _ffecb {
						_gaba = append(_gaba, string(_bfeec))
					} else {
						_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _aada, charcodes, _bfab.isCIDFont(), _aedfb, _cfaadg)
						_daad++
						_gaba = append(_gaba, _ebe.MissingCodeString)
					}
				}
				return _gaba, len(_gaba), _daad
			}
		}
	}
	for _, _dgdeg := range charcodes {
		if _bfab._dcdd != nil {
			if _gfebg, _dfdga := _bfab._dcdd.CharcodeToUnicode(_ebe.CharCode(_dgdeg)); _dfdga {
				_gaba = append(_gaba, _gfebg)
				continue
			}
		}
		if _cfaadg != nil {
			if _gdef, _eced := _cfaadg.CharcodeToRune(_dgdeg); _eced {
				_gaba = append(_gaba, string(_gdef))
				continue
			}
		}
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _dgdeg, charcodes, _bfab.isCIDFont(), _aedfb, _cfaadg)
		_daad++
		_gaba = append(_gaba, _ebe.MissingCodeString)
	}
	if _daad != 0 {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0043\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0074\u006f\u0020u\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020i\u006ep\u0075\u0074\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0043\u0068\u0061\u0072\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073", len(charcodes), _daad, _aedfb)
	}
	return _gaba, len(_gaba), _daad
}

// SetContentStreams sets the content streams based on a string array. Will make
// 1 object stream for each string and reference from the page Contents.
// Each stream will be encoded using the encoding specified by the StreamEncoder,
// if empty, will use identity encoding (raw data).
func (_gaceg *PdfPage) SetContentStreams(cStreams []string, encoder _ebb.StreamEncoder) error {
	if len(cStreams) == 0 {
		_gaceg.Contents = nil
		return nil
	}
	if encoder == nil {
		encoder = _ebb.NewRawEncoder()
	}
	var _daae []*_ebb.PdfObjectStream
	for _, _fdff := range cStreams {
		_fgdde := &_ebb.PdfObjectStream{}
		_cfgb := encoder.MakeStreamDict()
		_feebg, _ceffg := encoder.EncodeBytes([]byte(_fdff))
		if _ceffg != nil {
			return _ceffg
		}
		_cfgb.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ebb.MakeInteger(int64(len(_feebg))))
		_fgdde.PdfObjectDictionary = _cfgb
		_fgdde.Stream = _feebg
		_daae = append(_daae, _fgdde)
	}
	if len(_daae) == 1 {
		_gaceg.Contents = _daae[0]
	} else {
		_dedf := _ebb.MakeArray()
		for _, _acaaa := range _daae {
			_dedf.Append(_acaaa)
		}
		_gaceg.Contents = _dedf
	}
	return nil
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a CalRGB device.
func (_acca *PdfColorspaceCalRGB) GetNumComponents() int { return 3 }

// PdfFieldSignature signature field represents digital signatures and optional data for authenticating
// the name of the signer and verifying document contents.
type PdfFieldSignature struct {
	*PdfField
	*PdfAnnotationWidget
	V    *PdfSignature
	Lock *_ebb.PdfIndirectObject
	SV   *_ebb.PdfIndirectObject
}

func (_bddb *PdfReader) newPdfAnnotationUnderlineFromDict(_bggf *_ebb.PdfObjectDictionary) (*PdfAnnotationUnderline, error) {
	_eecfb := PdfAnnotationUnderline{}
	_abbb, _gaee := _bddb.newPdfAnnotationMarkupFromDict(_bggf)
	if _gaee != nil {
		return nil, _gaee
	}
	_eecfb.PdfAnnotationMarkup = _abbb
	_eecfb.QuadPoints = _bggf.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_eecfb, nil
}
func _bcdgc(_fgaf *_ebb.PdfObjectDictionary, _baca *fontCommon, _bffb _da.TextEncoder) (*pdfFontSimple, error) {
	_dggce := _bede(_baca)
	_dggce._dacee = _bffb
	if _bffb == nil {
		_dgdbdf := _fgaf.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
		if _dgdbdf == nil {
			_dgdbdf = _ebb.MakeInteger(0)
		}
		_dggce.FirstChar = _dgdbdf
		_cefec, _geee := _ebb.GetIntVal(_dgdbdf)
		if !_geee {
			_eg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _dgdbdf)
			return nil, _ebb.ErrTypeError
		}
		_ggbbb := _da.CharCode(_cefec)
		_dgdbdf = _fgaf.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
		if _dgdbdf == nil {
			_dgdbdf = _ebb.MakeInteger(255)
		}
		_dggce.LastChar = _dgdbdf
		_cefec, _geee = _ebb.GetIntVal(_dgdbdf)
		if !_geee {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _dgdbdf)
			return nil, _ebb.ErrTypeError
		}
		_acgb := _da.CharCode(_cefec)
		_dggce._cdff = make(map[_da.CharCode]float64)
		_dgdbdf = _fgaf.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _dgdbdf != nil {
			_dggce.Widths = _dgdbdf
			_acdfd, _ebae := _ebb.GetArray(_dgdbdf)
			if !_ebae {
				_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _dgdbdf)
				return nil, _ebb.ErrTypeError
			}
			_cbgea, _aaaf := _acdfd.ToFloat64Array()
			if _aaaf != nil {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
				return nil, _aaaf
			}
			if len(_cbgea) != int(_acgb-_ggbbb+1) {
				_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _acgb-_ggbbb+1, len(_cbgea))
				return nil, _ebb.ErrRangeError
			}
			for _adcb, _acfdb := range _cbgea {
				_dggce._cdff[_ggbbb+_da.CharCode(_adcb)] = _acfdb
			}
		}
	}
	_dggce.Encoding = _ebb.TraceToDirectObject(_fgaf.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	return _dggce, nil
}

// NewPdfReaderWithOpts creates a new PdfReader for an input io.ReadSeeker interface
// with a ReaderOpts.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderWithOpts(rs _ab.ReadSeeker, opts *ReaderOpts) (*PdfReader, error) {
	const _ccca = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0057\u0069\u0074\u0068O\u0070\u0074\u0073"
	return _dcbd(rs, opts, true, _ccca)
}

// ColorToRGB converts a Lab color to an RGB color.
func (_dgead *PdfColorspaceLab) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ddeg := func(_ffdc float64) float64 {
		if _ffdc >= 6.0/29 {
			return _ffdc * _ffdc * _ffdc
		}
		return 108.0 / 841 * (_ffdc - 4/29)
	}
	_egff, _afdf := color.(*PdfColorLab)
	if !_afdf {
		_eg.Log.Debug("\u0069\u006e\u0070\u0075t \u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u006c\u0061\u0062")
		return nil, _gf.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	LStar := _egff.L()
	AStar := _egff.A()
	BStar := _egff.B()
	L := (LStar+16)/116 + AStar/500
	M := (LStar + 16) / 116
	N := (LStar+16)/116 - BStar/200
	X := _dgead.WhitePoint[0] * _ddeg(L)
	Y := _dgead.WhitePoint[1] * _ddeg(M)
	Z := _dgead.WhitePoint[2] * _ddeg(N)
	_aag := 3.240479*X + -1.537150*Y + -0.498535*Z
	_ddba := -0.969256*X + 1.875992*Y + 0.041556*Z
	_fccc := 0.055648*X + -0.204043*Y + 1.057311*Z
	_aag = _cbg.Min(_cbg.Max(_aag, 0), 1.0)
	_ddba = _cbg.Min(_cbg.Max(_ddba, 0), 1.0)
	_fccc = _cbg.Min(_cbg.Max(_fccc, 0), 1.0)
	return NewPdfColorDeviceRGB(_aag, _ddba, _fccc), nil
}

// ToPdfObject returns an indirect object containing the signature field dictionary.
func (_egec *PdfFieldSignature) ToPdfObject() _ebb.PdfObject {
	if _egec.PdfAnnotationWidget != nil {
		_egec.PdfAnnotationWidget.ToPdfObject()
	}
	_egec.PdfField.ToPdfObject()
	_dfbcc := _egec._cdfd
	_deda := _dfbcc.PdfObject.(*_ebb.PdfObjectDictionary)
	_deda.SetIfNotNil("\u0046\u0054", _ebb.MakeName("\u0053\u0069\u0067"))
	_deda.SetIfNotNil("\u004c\u006f\u0063\u006b", _egec.Lock)
	_deda.SetIfNotNil("\u0053\u0056", _egec.SV)
	if _egec.V != nil {
		_deda.SetIfNotNil("\u0056", _egec.V.ToPdfObject())
	}
	return _dfbcc
}

// NewPdfActionMovie returns a new "movie" action.
func NewPdfActionMovie() *PdfActionMovie {
	_dgab := NewPdfAction()
	_ede := &PdfActionMovie{}
	_ede.PdfAction = _dgab
	_dgab.SetContext(_ede)
	return _ede
}
func (_edec *PdfReader) newPdfActionThreadFromDict(_faf *_ebb.PdfObjectDictionary) (*PdfActionThread, error) {
	_caa, _abd := _gggf(_faf.Get("\u0046"))
	if _abd != nil {
		return nil, _abd
	}
	return &PdfActionThread{D: _faf.Get("\u0044"), B: _faf.Get("\u0042"), F: _caa}, nil
}

// NewPdfDateFromTime will create a PdfDate based on the given time
func NewPdfDateFromTime(timeObj _f.Time) (PdfDate, error) {
	_fadff := timeObj.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_faddf, _ := _aa.ParseInt(_fadff[1:3], 10, 32)
	_aaadg, _ := _aa.ParseInt(_fadff[4:6], 10, 32)
	return PdfDate{_dacdd: int64(timeObj.Year()), _agaba: int64(timeObj.Month()), _edcfe: int64(timeObj.Day()), _aedbe: int64(timeObj.Hour()), _cfaba: int64(timeObj.Minute()), _cgbcd: int64(timeObj.Second()), _degeb: _fadff[0], _gcccb: _faddf, _fddcbg: _aaadg}, nil
}
func _gefgb(_eaccf []byte) (_fcad, _gcbd string, _bgad error) {
	_eg.Log.Trace("g\u0065\u0074\u0041\u0053CI\u0049S\u0065\u0063\u0074\u0069\u006fn\u0073\u003a\u0020\u0025\u0064\u0020", len(_eaccf))
	_cfff := _aabec.FindIndex(_eaccf)
	if _cfff == nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0067\u0065\u0074\u0041\u0053\u0043\u0049\u0049\u0053\u0065\u0063\u0074\u0069o\u006e\u0073\u002e\u0020\u004e\u006f\u0020d\u0069\u0063\u0074\u002e")
		return "", "", _ebb.ErrTypeError
	}
	_ebfb := _cfff[1]
	_dgabg := _ee.Index(string(_eaccf[_ebfb:]), _aeecf)
	if _dgabg < 0 {
		_fcad = string(_eaccf[_ebfb:])
		return _fcad, "", nil
	}
	_gfege := _ebfb + _dgabg
	_fcad = string(_eaccf[_ebfb:_gfege])
	_egdd := _gfege
	_dgabg = _ee.Index(string(_eaccf[_egdd:]), _ccdgc)
	if _dgabg < 0 {
		_eg.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0067e\u0074\u0041\u0053\u0043\u0049\u0049\u0053e\u0063\u0074\u0069\u006f\u006e\u0073\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bgad)
		return "", "", _ebb.ErrTypeError
	}
	_cfef := _egdd + _dgabg
	_gcbd = string(_eaccf[_egdd:_cfef])
	return _fcad, _gcbd, nil
}
func _ageed(_cgbc _ebb.PdfObject) (*PdfColorspaceCalGray, error) {
	_agab := NewPdfColorspaceCalGray()
	if _ffbf, _bagd := _cgbc.(*_ebb.PdfIndirectObject); _bagd {
		_agab._bebfd = _ffbf
	}
	_cgbc = _ebb.TraceToDirectObject(_cgbc)
	_ddaf, _ecae := _cgbc.(*_ebb.PdfObjectArray)
	if !_ecae {
		return nil, _bg.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _ddaf.Len() != 2 {
		return nil, _bg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0061\u006cG\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_cgbc = _ebb.TraceToDirectObject(_ddaf.Get(0))
	_aggc, _ecae := _cgbc.(*_ebb.PdfObjectName)
	if !_ecae {
		return nil, _bg.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u0020\u006e\u0061m\u0065\u0020\u006e\u006f\u0074\u0020\u0061 \u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_aggc != "\u0043a\u006c\u0047\u0072\u0061\u0079" {
		return nil, _bg.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0043\u0061\u006c\u0047\u0072\u0061\u0079 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_cgbc = _ebb.TraceToDirectObject(_ddaf.Get(1))
	_eeaab, _ecae := _cgbc.(*_ebb.PdfObjectDictionary)
	if !_ecae {
		return nil, _bg.Errorf("\u0043\u0061lG\u0072\u0061\u0079 \u0064\u0069\u0063\u0074 no\u0074 a\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u006f\u0062\u006a\u0065c\u0074")
	}
	_cgbc = _eeaab.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_cgbc = _ebb.TraceToDirectObject(_cgbc)
	_eecc, _ecae := _cgbc.(*_ebb.PdfObjectArray)
	if !_ecae {
		return nil, _bg.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020W\u0068\u0069\u0074e\u0050o\u0069\u006e\u0074")
	}
	if _eecc.Len() != 3 {
		return nil, _bg.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
	}
	_fceb, _ffag := _eecc.GetAsFloat64Slice()
	if _ffag != nil {
		return nil, _ffag
	}
	_agab.WhitePoint = _fceb
	_cgbc = _eeaab.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _cgbc != nil {
		_cgbc = _ebb.TraceToDirectObject(_cgbc)
		_beab, _dbfb := _cgbc.(*_ebb.PdfObjectArray)
		if !_dbfb {
			return nil, _bg.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020B\u006c\u0061\u0063k\u0050o\u0069\u006e\u0074")
		}
		if _beab.Len() != 3 {
			return nil, _bg.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061c\u006b\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
		}
		_ebfac, _eege := _beab.GetAsFloat64Slice()
		if _eege != nil {
			return nil, _eege
		}
		_agab.BlackPoint = _ebfac
	}
	_cgbc = _eeaab.Get("\u0047\u0061\u006dm\u0061")
	if _cgbc != nil {
		_cgbc = _ebb.TraceToDirectObject(_cgbc)
		_edcd, _bafe := _ebb.GetNumberAsFloat(_cgbc)
		if _bafe != nil {
			return nil, _bg.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0067\u0061\u006d\u006d\u0061\u0020\u006e\u006ft\u0020\u0061\u0020n\u0075m\u0062\u0065\u0072")
		}
		_agab.Gamma = _edcd
	}
	return _agab, nil
}

// DecodeArray returns the range of color component values in the Lab colorspace.
func (_dfdc *PdfColorspaceLab) DecodeArray() []float64 {
	_cagfc := []float64{0, 100}
	if _dfdc.Range != nil && len(_dfdc.Range) == 4 {
		_cagfc = append(_cagfc, _dfdc.Range...)
	} else {
		_cagfc = append(_cagfc, -100, 100, -100, 100)
	}
	return _cagfc
}

// PdfColorPattern represents a pattern color.
type PdfColorPattern struct {
	Color       PdfColor
	PatternName _ebb.PdfObjectName
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 4 for a CMYK32 device.
func (_dgfga *PdfColorspaceDeviceCMYK) GetNumComponents() int { return 4 }
func (_debba *PdfWriter) makeOffSetReference(_gabff int64) {
	_cbcg := _bg.Sprintf("\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u000a\u0025\u0064\u000a", _gabff)
	_debba.writeString(_cbcg)
	_debba.writeString("\u0025\u0025\u0045\u004f\u0046\u000a")
}

// SetRotation sets the rotation of all pages added to writer. The rotation is
// specified in degrees and must be a multiple of 90.
// The Rotate field of individual pages has priority over the global rotation.
func (_cebe *PdfWriter) SetRotation(rotate int64) error {
	_gdgda, _eaccc := _ebb.GetDict(_cebe._dggbf)
	if !_eaccc {
		return ErrTypeCheck
	}
	_gdgda.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _ebb.MakeInteger(rotate))
	return nil
}
func _fbea(_dbadc *PdfField) []*PdfField {
	_badac := []*PdfField{_dbadc}
	for _, _dfbbg := range _dbadc.Kids {
		_badac = append(_badac, _fbea(_dfbbg)...)
	}
	return _badac
}

// IsCheckbox returns true if the button field represents a checkbox, false otherwise.
func (_gcdg *PdfFieldButton) IsCheckbox() bool { return _gcdg.GetType() == ButtonTypeCheckbox }

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_acef *PdfShadingType5) ToPdfObject() _ebb.PdfObject {
	_acef.PdfShading.ToPdfObject()
	_cebfc, _cacgg := _acef.getShadingDict()
	if _cacgg != nil {
		_eg.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _acef.BitsPerCoordinate != nil {
		_cebfc.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _acef.BitsPerCoordinate)
	}
	if _acef.BitsPerComponent != nil {
		_cebfc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _acef.BitsPerComponent)
	}
	if _acef.VerticesPerRow != nil {
		_cebfc.Set("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077", _acef.VerticesPerRow)
	}
	if _acef.Decode != nil {
		_cebfc.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _acef.Decode)
	}
	if _acef.Function != nil {
		if len(_acef.Function) == 1 {
			_cebfc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _acef.Function[0].ToPdfObject())
		} else {
			_eafb := _ebb.MakeArray()
			for _, _agbac := range _acef.Function {
				_eafb.Append(_agbac.ToPdfObject())
			}
			_cebfc.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _eafb)
		}
	}
	return _acef._fbfae
}

type pdfFontType3 struct {
	fontCommon
	_adec *_ebb.PdfIndirectObject

	// These fields are specific to Type 3 fonts.
	CharProcs  _ebb.PdfObject
	Encoding   _ebb.PdfObject
	FontBBox   _ebb.PdfObject
	FontMatrix _ebb.PdfObject
	FirstChar  _ebb.PdfObject
	LastChar   _ebb.PdfObject
	Widths     _ebb.PdfObject
	Resources  _ebb.PdfObject
	_aggbb     map[_da.CharCode]float64
	_abcba     _da.TextEncoder
}

// ImageToRGB converts an image with samples in Separation CS to an image with samples specified in
// DeviceRGB CS.
func (_cgeg *PdfColorspaceSpecialSeparation) ImageToRGB(img Image) (Image, error) {
	_eaab := _abg.NewReader(img.getBase())
	_cfec := _dg.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), _cgeg.AlternateSpace.GetNumComponents(), nil, img._dagcb, nil)
	_debb := _abg.NewWriter(_cfec)
	_eaaf := _cbg.Pow(2, float64(img.BitsPerComponent)) - 1
	_eg.Log.Trace("\u0053\u0065\u0070a\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u002d\u003e\u0020\u0054\u006f\u0052\u0047\u0042\u0020\u0063o\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	_eg.Log.Trace("\u0054i\u006et\u0054\u0072\u0061\u006e\u0073f\u006f\u0072m\u003a\u0020\u0025\u002b\u0076", _cgeg.TintTransform)
	_bdff := _cgeg.AlternateSpace.DecodeArray()
	var (
		_eebd uint32
		_babe error
	)
	for {
		_eebd, _babe = _eaab.ReadSample()
		if _babe == _ab.EOF {
			break
		}
		if _babe != nil {
			return img, _babe
		}
		_ebaad := float64(_eebd) / _eaaf
		_gebee, _eeega := _cgeg.TintTransform.Evaluate([]float64{_ebaad})
		if _eeega != nil {
			return img, _eeega
		}
		for _cbb, _acddd := range _gebee {
			_dgdbd := _dg.LinearInterpolate(_acddd, _bdff[_cbb*2], _bdff[_cbb*2+1], 0, 1)
			if _eeega = _debb.WriteSample(uint32(_dgdbd * _eaaf)); _eeega != nil {
				return img, _eeega
			}
		}
	}
	return _cgeg.AlternateSpace.ImageToRGB(_afacb(&_cfec))
}

// GetParamsDict returns *core.PdfObjectDictionary with a set of basic image parameters.
func (_dgfega *Image) GetParamsDict() *_ebb.PdfObjectDictionary {
	_bbggd := _ebb.MakeDict()
	_bbggd.Set("\u0057\u0069\u0064t\u0068", _ebb.MakeInteger(_dgfega.Width))
	_bbggd.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _ebb.MakeInteger(_dgfega.Height))
	_bbggd.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _ebb.MakeInteger(int64(_dgfega.ColorComponents)))
	_bbggd.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _ebb.MakeInteger(_dgfega.BitsPerComponent))
	return _bbggd
}

// NewPdfActionRendition returns a new "rendition" action.
func NewPdfActionRendition() *PdfActionRendition {
	_dbd := NewPdfAction()
	_bbc := &PdfActionRendition{}
	_bbc.PdfAction = _dbd
	_dbd.SetContext(_bbc)
	return _bbc
}
func (_dfge *PdfWriter) writeOutputIntents() error {
	if len(_dfge._geced) == 0 {
		return nil
	}
	_acgfg := make([]_ebb.PdfObject, len(_dfge._geced))
	for _dfafbf, _bgbfd := range _dfge._geced {
		_cadb := _bgbfd.ToPdfObject()
		_acgfg[_dfafbf] = _ebb.MakeIndirectObject(_cadb)
	}
	_dbee := _ebb.MakeIndirectObject(_ebb.MakeArray(_acgfg...))
	_dfge._dffegd.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _dbee)
	if _abgag := _dfge.addObjects(_dbee); _abgag != nil {
		return _abgag
	}
	return nil
}

// NewPdfColorLab returns a new Lab color.
func NewPdfColorLab(l, a, b float64) *PdfColorLab { _abeag := PdfColorLab{l, a, b}; return &_abeag }

// ToPdfObject implements interface PdfModel.
func (_bgbfe *PdfSignature) ToPdfObject() _ebb.PdfObject {
	_gbgc := _bgbfe._ffbgc
	var _dfcfd *_ebb.PdfObjectDictionary
	if _aagee, _fbdcb := _gbgc.PdfObject.(*pdfSignDictionary); _fbdcb {
		_dfcfd = _aagee.PdfObjectDictionary
	} else {
		_dfcfd = _gbgc.PdfObject.(*_ebb.PdfObjectDictionary)
	}
	_dfcfd.SetIfNotNil("\u0054\u0079\u0070\u0065", _bgbfe.Type)
	_dfcfd.SetIfNotNil("\u0046\u0069\u006c\u0074\u0065\u0072", _bgbfe.Filter)
	_dfcfd.SetIfNotNil("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r", _bgbfe.SubFilter)
	_dfcfd.SetIfNotNil("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _bgbfe.ByteRange)
	_dfcfd.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _bgbfe.Contents)
	_dfcfd.SetIfNotNil("\u0043\u0065\u0072\u0074", _bgbfe.Cert)
	_dfcfd.SetIfNotNil("\u004e\u0061\u006d\u0065", _bgbfe.Name)
	_dfcfd.SetIfNotNil("\u0052\u0065\u0061\u0073\u006f\u006e", _bgbfe.Reason)
	_dfcfd.SetIfNotNil("\u004d", _bgbfe.M)
	_dfcfd.SetIfNotNil("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e", _bgbfe.Reference)
	_dfcfd.SetIfNotNil("\u0043h\u0061\u006e\u0067\u0065\u0073", _bgbfe.Changes)
	_dfcfd.SetIfNotNil("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f", _bgbfe.ContactInfo)
	return _gbgc
}

// NewPdfDate returns a new PdfDate object from a PDF date string (see 7.9.4 Dates).
// format: "D: YYYYMMDDHHmmSSOHH'mm"
func NewPdfDate(dateStr string) (PdfDate, error) {
	_cffed, _dfdgfe := _fd.ParsePdfTime(dateStr)
	if _dfdgfe != nil {
		return PdfDate{}, _dfdgfe
	}
	return NewPdfDateFromTime(_cffed)
}

// NewPdfAppender creates a new Pdf appender from a Pdf reader.
func NewPdfAppender(reader *PdfReader) (*PdfAppender, error) {
	_fgbb := &PdfAppender{_ecce: reader._ggdg, Reader: reader, _gege: reader._cafdf, _agb: reader._dfadc}
	_acfca, _bbge := _fgbb._ecce.Seek(0, _ab.SeekEnd)
	if _bbge != nil {
		return nil, _bbge
	}
	_fgbb._bee = _acfca
	if _, _bbge = _fgbb._ecce.Seek(0, _ab.SeekStart); _bbge != nil {
		return nil, _bbge
	}
	_fgbb._acfe, _bbge = NewPdfReader(_fgbb._ecce)
	if _bbge != nil {
		return nil, _bbge
	}
	for _, _gcde := range _fgbb.Reader.GetObjectNums() {
		if _fgbb._gbddb < _gcde {
			_fgbb._gbddb = _gcde
		}
	}
	_fgbb._acfd = _fgbb._gege.GetXrefTable()
	_fgbb._cfag = _fgbb._gege.GetXrefOffset()
	_fgbb._dfbg = append(_fgbb._dfbg, _fgbb._acfe.PageList...)
	_fgbb._ddfg = make(map[_ebb.PdfObject]struct{})
	_fgbb._gbfa = make(map[_ebb.PdfObject]int64)
	_fgbb._eebc = make(map[_ebb.PdfObject]struct{})
	_fgbb._bfef = _fgbb._acfe.AcroForm
	_fgbb._eged = _fgbb._acfe.DSS
	return _fgbb, nil
}

// PdfAnnotationWatermark represents Watermark annotations.
// (Section 12.5.6.22).
type PdfAnnotationWatermark struct {
	*PdfAnnotation
	FixedPrint _ebb.PdfObject
}

// GetAsTilingPattern returns a tiling pattern. Check with IsTiling() prior to using this.
func (_eggc *PdfPattern) GetAsTilingPattern() *PdfTilingPattern {
	return _eggc._ffagg.(*PdfTilingPattern)
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_beccf pdfFontType0) GetRuneMetrics(r rune) (_bad.CharMetrics, bool) {
	if _beccf.DescendantFont == nil {
		_eg.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _beccf)
		return _bad.CharMetrics{}, false
	}
	return _beccf.DescendantFont.GetRuneMetrics(r)
}

// GetContainingPdfObject returns the container of the PdfAcroForm (indirect object).
func (_ccfac *PdfAcroForm) GetContainingPdfObject() _ebb.PdfObject { return _ccfac._adcg }

// PdfAnnotationRedact represents Redact annotations.
// (Section 12.5.6.23).
type PdfAnnotationRedact struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints  _ebb.PdfObject
	IC          _ebb.PdfObject
	RO          _ebb.PdfObject
	OverlayText _ebb.PdfObject
	Repeat      _ebb.PdfObject
	DA          _ebb.PdfObject
	Q           _ebb.PdfObject
}

// Encoder returns the font's text encoder.
func (_febedc *PdfFont) Encoder() _da.TextEncoder {
	_edac := _febedc.actualFont()
	if _edac == nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0045n\u0063\u006f\u0064er\u0020\u006e\u006f\u0074\u0020\u0069m\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066o\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d%\u0023\u0054", _febedc._ebcad)
		return nil
	}
	return _edac.Encoder()
}

// GetDescent returns the Descent of the font `descriptor`.
func (_eegcf *PdfFontDescriptor) GetDescent() (float64, error) {
	return _ebb.GetNumberAsFloat(_eegcf.Descent)
}
func (_gccaa *PdfReader) traverseObjectData(_cfabd _ebb.PdfObject) error {
	return _ebb.ResolveReferencesDeep(_cfabd, _gccaa._dfadc)
}

// WriteString outputs the object as it is to be written to file.
func (_ddeaf *pdfSignDictionary) WriteString() string {
	_ddeaf._eefbe = 0
	_ddeaf._adggf = 0
	_ddeaf._decgd = 0
	_ddeaf._ddceg = 0
	_eaagd := _ca.NewBuffer(nil)
	_eaagd.WriteString("\u003c\u003c")
	for _, _eaaeac := range _ddeaf.Keys() {
		_fgccf := _ddeaf.Get(_eaaeac)
		switch _eaaeac {
		case "\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e":
			_eaagd.WriteString(_eaaeac.WriteString())
			_eaagd.WriteString("\u0020")
			_ddeaf._decgd = _eaagd.Len()
			_eaagd.WriteString(_fgccf.WriteString())
			_eaagd.WriteString("\u0020")
			_ddeaf._ddceg = _eaagd.Len() - 1
		case "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073":
			_eaagd.WriteString(_eaaeac.WriteString())
			_eaagd.WriteString("\u0020")
			_ddeaf._eefbe = _eaagd.Len()
			_eaagd.WriteString(_fgccf.WriteString())
			_eaagd.WriteString("\u0020")
			_ddeaf._adggf = _eaagd.Len() - 1
		default:
			_eaagd.WriteString(_eaaeac.WriteString())
			_eaagd.WriteString("\u0020")
			_eaagd.WriteString(_fgccf.WriteString())
		}
	}
	_eaagd.WriteString("\u003e\u003e")
	return _eaagd.String()
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the L, A and B components of the color.
func (_efcd *PdfColorspaceLab) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_geec, _bccba := _ebb.GetNumbersAsFloat(objects)
	if _bccba != nil {
		return nil, _bccba
	}
	return _efcd.ColorFromFloats(_geec)
}

// ToPdfObject implements interface PdfModel.
func (_bfdg *PdfAnnotationWidget) ToPdfObject() _ebb.PdfObject {
	_bfdg.PdfAnnotation.ToPdfObject()
	_bfa := _bfdg._bdcd
	_beaa := _bfa.PdfObject.(*_ebb.PdfObjectDictionary)
	if _bfdg._gdga {
		return _bfa
	}
	_bfdg._gdga = true
	_beaa.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0057\u0069\u0064\u0067\u0065\u0074"))
	_beaa.SetIfNotNil("\u0048", _bfdg.H)
	_beaa.SetIfNotNil("\u004d\u004b", _bfdg.MK)
	_beaa.SetIfNotNil("\u0041", _bfdg.A)
	_beaa.SetIfNotNil("\u0041\u0041", _bfdg.AA)
	_beaa.SetIfNotNil("\u0042\u0053", _bfdg.BS)
	_gcd := _bfdg.Parent
	if _bfdg._gce != nil {
		if _bfdg._gce._cdfd == _bfdg._bdcd {
			_bfdg._gce.ToPdfObject()
		}
		_gcd = _bfdg._gce.GetContainingPdfObject()
	}
	if _gcd != _bfa {
		_beaa.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _gcd)
	}
	_bfdg._gdga = false
	return _bfa
}

// ColorFromPdfObjects gets the color from a series of pdf objects (4 for cmyk).
func (_ffeb *PdfColorspaceDeviceCMYK) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 4 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_feba, _bded := _ebb.GetNumbersAsFloat(objects)
	if _bded != nil {
		return nil, _bded
	}
	return _ffeb.ColorFromFloats(_feba)
}
func (_ffgc *PdfColorspaceSpecialSeparation) String() string {
	return "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"
}
func (_cbe *PdfReader) newPdfActionTransFromDict(_efbe *_ebb.PdfObjectDictionary) (*PdfActionTrans, error) {
	return &PdfActionTrans{Trans: _efbe.Get("\u0054\u0072\u0061n\u0073")}, nil
}

// NewCompositePdfFontFromTTF loads a composite TTF font. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTF.
func NewCompositePdfFontFromTTF(r _ab.ReadSeeker) (*PdfFont, error) {
	_efaaa, _gfedg := _ef.ReadAll(r)
	if _gfedg != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _gfedg)
		return nil, _gfedg
	}
	_afcabd, _gfedg := _bad.TtfParse(_ca.NewReader(_efaaa))
	if _gfedg != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067 \u0074\u0074\u0066\u0020\u0066\u006f\u006et\u003a\u0020\u0025\u0076", _gfedg)
		return nil, _gfedg
	}
	_aafcc := &pdfCIDFontType2{fontCommon: fontCommon{_dfbf: "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032"}, CIDToGIDMap: _ebb.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079")}
	if len(_afcabd.Widths) <= 0 {
		return nil, _gf.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_dafee := 1000.0 / float64(_afcabd.UnitsPerEm)
	_bfeb := _dafee * float64(_afcabd.Widths[0])
	_ddecg := make(map[rune]int)
	_aedfbc := make(map[_bad.GID]int)
	_fdggd := _bad.GID(len(_afcabd.Widths))
	for _fadbe, _befd := range _afcabd.Chars {
		if _befd > _fdggd-1 {
			continue
		}
		_dgcd := int(_dafee * float64(_afcabd.Widths[_befd]))
		_ddecg[_fadbe] = _dgcd
		_aedfbc[_befd] = _dgcd
	}
	_aafcc._dceb = _ddecg
	_aafcc.DW = _ebb.MakeInteger(int64(_bfeb))
	_acgge := _fgag(_aedfbc, uint16(_fdggd))
	_aafcc.W = _ebb.MakeIndirectObject(_acgge)
	_fadaa := _ebb.MakeDict()
	_fadaa.Set("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067", _ebb.MakeString("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"))
	_fadaa.Set("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079", _ebb.MakeString("\u0041\u0064\u006fb\u0065"))
	_fadaa.Set("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074", _ebb.MakeInteger(0))
	_aafcc.CIDSystemInfo = _fadaa
	_ffbe := &PdfFontDescriptor{FontName: _ebb.MakeName(_afcabd.PostScriptName), Ascent: _ebb.MakeFloat(_dafee * float64(_afcabd.TypoAscender)), Descent: _ebb.MakeFloat(_dafee * float64(_afcabd.TypoDescender)), CapHeight: _ebb.MakeFloat(_dafee * float64(_afcabd.CapHeight)), FontBBox: _ebb.MakeArrayFromFloats([]float64{_dafee * float64(_afcabd.Xmin), _dafee * float64(_afcabd.Ymin), _dafee * float64(_afcabd.Xmax), _dafee * float64(_afcabd.Ymax)}), ItalicAngle: _ebb.MakeFloat(_afcabd.ItalicAngle), MissingWidth: _ebb.MakeFloat(_bfeb)}
	_cbge, _gfedg := _ebb.MakeStream(_efaaa, _ebb.NewFlateEncoder())
	if _gfedg != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _gfedg)
		return nil, _gfedg
	}
	_cbge.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _ebb.MakeInteger(int64(len(_efaaa))))
	_ffbe.FontFile2 = _cbge
	if _afcabd.Bold {
		_ffbe.StemV = _ebb.MakeInteger(120)
	} else {
		_ffbe.StemV = _ebb.MakeInteger(70)
	}
	_fdfgd := _dffee
	if _afcabd.IsFixedPitch {
		_fdfgd |= _gfega
	}
	if _afcabd.ItalicAngle != 0 {
		_fdfgd |= _ccgbe
	}
	_ffbe.Flags = _ebb.MakeInteger(int64(_fdfgd))
	_aafcc._fdacg = _afcabd.PostScriptName
	_aafcc._fbbd = _ffbe
	_cbged := pdfFontType0{fontCommon: fontCommon{_dfbf: "\u0054\u0079\u0070e\u0030", _fdacg: _afcabd.PostScriptName}, DescendantFont: &PdfFont{_ebcad: _aafcc}, Encoding: _ebb.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048"), _bfdgc: _afcabd.NewEncoder()}
	if len(_afcabd.Chars) > 0 {
		_bgeg := make(map[_ebe.CharCode]rune, len(_afcabd.Chars))
		for _aecc, _caagd := range _afcabd.Chars {
			_fcgfg := _ebe.CharCode(_caagd)
			if _eebde, _aadac := _bgeg[_fcgfg]; !_aadac || (_aadac && _eebde > _aecc) {
				_bgeg[_fcgfg] = _aecc
			}
		}
		_cbged._dcdd = _ebe.NewToUnicodeCMap(_bgeg)
	}
	_bdgd := PdfFont{_ebcad: &_cbged}
	return &_bdgd, nil
}
func _fgff(_bgcd *_ebb.PdfIndirectObject) (*PdfOutline, error) {
	_bfgbc, _cedb := _bgcd.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_cedb {
		return nil, _bg.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_dgda := NewPdfOutline()
	if _aedfgd := _bfgbc.Get("\u0054\u0079\u0070\u0065"); _aedfgd != nil {
		_cbgd, _ddad := _aedfgd.(*_ebb.PdfObjectName)
		if _ddad {
			if *_cbgd != "\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073" {
				_eg.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u004f\u0075\u0074l\u0069\u006e\u0065s\u0020(\u0025\u0073\u0029", *_cbgd)
			}
		}
	}
	if _bdecd := _bfgbc.Get("\u0043\u006f\u0075n\u0074"); _bdecd != nil {
		_dcddf, _cddf := _ebb.GetNumberAsInt64(_bdecd)
		if _cddf != nil {
			return nil, _cddf
		}
		_dgda.Count = &_dcddf
	}
	return _dgda, nil
}

// UpdateXObjectImageFromImage creates a new XObject Image from an
// Image object `img` and default masks from xobjIn.
// The default masks are overridden if img.hasAlpha
// If `encoder` is nil, uses raw encoding (none).
func UpdateXObjectImageFromImage(xobjIn *XObjectImage, img *Image, cs PdfColorspace, encoder _ebb.StreamEncoder) (*XObjectImage, error) {
	if encoder == nil {
		encoder = _ebb.NewRawEncoder()
	}
	encoder.UpdateParams(img.GetParamsDict())
	_dgbg, _edgdc := encoder.EncodeBytes(img.Data)
	if _edgdc != nil {
		_eg.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _edgdc)
		return nil, _edgdc
	}
	_abfbgf := NewXObjectImage()
	_agdgb := img.Width
	_geedc := img.Height
	_abfbgf.Width = &_agdgb
	_abfbgf.Height = &_geedc
	_bcbba := img.BitsPerComponent
	_abfbgf.BitsPerComponent = &_bcbba
	_abfbgf.Filter = encoder
	_abfbgf.Stream = _dgbg
	if cs == nil {
		if img.ColorComponents == 1 {
			_abfbgf.ColorSpace = NewPdfColorspaceDeviceGray()
		} else if img.ColorComponents == 3 {
			_abfbgf.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			_abfbgf.ColorSpace = NewPdfColorspaceDeviceCMYK()
		} else {
			return nil, _gf.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_abfbgf.ColorSpace = cs
	}
	if len(img._dagcb) != 0 {
		_ddgbdc := NewXObjectImage()
		_ddgbdc.Filter = encoder
		_aafga, _bfdaa := encoder.EncodeBytes(img._dagcb)
		if _bfdaa != nil {
			_eg.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _bfdaa)
			return nil, _bfdaa
		}
		_ddgbdc.Stream = _aafga
		_ddgbdc.BitsPerComponent = _abfbgf.BitsPerComponent
		_ddgbdc.Width = &img.Width
		_ddgbdc.Height = &img.Height
		_ddgbdc.ColorSpace = NewPdfColorspaceDeviceGray()
		_abfbgf.SMask = _ddgbdc.ToPdfObject()
	} else {
		_abfbgf.SMask = xobjIn.SMask
		_abfbgf.ImageMask = xobjIn.ImageMask
		if _abfbgf.ColorSpace.GetNumComponents() == 1 {
			_dddfe(_abfbgf)
		}
	}
	return _abfbgf, nil
}

// GetModelFromPrimitive returns the model corresponding to the `primitive` PdfObject.
func (_ffdgf *modelManager) GetModelFromPrimitive(primitive _ebb.PdfObject) PdfModel {
	model, _fagb := _ffdgf._aabcbdd[primitive]
	if !_fagb {
		return nil
	}
	return model
}

// ToPdfObject implements interface PdfModel.
func (_baee *PdfActionTrans) ToPdfObject() _ebb.PdfObject {
	_baee.PdfAction.ToPdfObject()
	_gfa := _baee._abe
	_bfdf := _gfa.PdfObject.(*_ebb.PdfObjectDictionary)
	_bfdf.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeTrans)))
	_bfdf.SetIfNotNil("\u0054\u0072\u0061n\u0073", _baee.Trans)
	return _gfa
}

// PdfActionNamed represents a named action.
type PdfActionNamed struct {
	*PdfAction
	N _ebb.PdfObject
}

// SetDecode sets the decode image float slice.
func (_deag *Image) SetDecode(decode []float64) { _deag._dgcea = decode }

// PdfAnnotationCircle represents Circle annotations.
// (Section 12.5.6.8).
type PdfAnnotationCircle struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _ebb.PdfObject
	IC _ebb.PdfObject
	BE _ebb.PdfObject
	RD _ebb.PdfObject
}

// Field returns the parent form field of the widget annotation, if one exists.
// NOTE: the method returns nil if the parent form field has not been parsed.
func (_bbb *PdfAnnotationWidget) Field() *PdfField { return _bbb._gce }

// AddOutlineTree adds outlines to a PDF file.
func (_eggeb *PdfWriter) AddOutlineTree(outlineTree *PdfOutlineTreeNode) { _eggeb._bcbee = outlineTree }
func (_gefe *DSS) generateHashMaps() error {
	_fcdcf, _befe := _gefe.generateHashMap(_gefe.Certs)
	if _befe != nil {
		return _befe
	}
	_facc, _befe := _gefe.generateHashMap(_gefe.OCSPs)
	if _befe != nil {
		return _befe
	}
	_caad, _befe := _gefe.generateHashMap(_gefe.CRLs)
	if _befe != nil {
		return _befe
	}
	_gefe._aeag = _fcdcf
	_gefe._cadd = _facc
	_gefe._fafgb = _caad
	return nil
}

// ImageToRGB converts Lab colorspace image to RGB and returns the result.
func (_dcbc *PdfColorspaceLab) ImageToRGB(img Image) (Image, error) {
	_febe := func(_bbdf float64) float64 {
		if _bbdf >= 6.0/29 {
			return _bbdf * _bbdf * _bbdf
		}
		return 108.0 / 841 * (_bbdf - 4/29)
	}
	_fbge := img._dgcea
	if len(_fbge) != 6 {
		_eg.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u002d\u0020\u004c\u0061\u0062\u0020\u0044e\u0063\u006f\u0064\u0065\u0020\u0072\u0061\u006e\u0067e\u0020\u0021\u003d\u0020\u0036\u002e\u002e\u002e\u0020\u0075\u0073\u0065\u0020\u005b0\u0020\u0031\u0030\u0030\u0020\u0061\u006d\u0069\u006e\u0020\u0061\u006d\u0061\u0078\u0020\u0062\u006d\u0069\u006e\u0020\u0062\u006d\u0061\u0078\u005d\u0020\u0064\u0065\u0066\u0061u\u006c\u0074\u0020\u0064\u0065\u0063\u006f\u0064\u0065 \u0061\u0072r\u0061\u0079")
		_fbge = _dcbc.DecodeArray()
	}
	_aggfe := _abg.NewReader(img.getBase())
	_edeag := _dg.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, img._dagcb, img._dgcea)
	_fca := _abg.NewWriter(_edeag)
	_adeg := _cbg.Pow(2, float64(img.BitsPerComponent)) - 1
	_ggadb := make([]uint32, 3)
	var (
		_bfdc                                               error
		Ls, As, Bs, L, M, N, X, Y, Z, _bbeae, _gegee, _dece float64
	)
	for {
		_bfdc = _aggfe.ReadSamples(_ggadb)
		if _bfdc == _ab.EOF {
			break
		} else if _bfdc != nil {
			return img, _bfdc
		}
		Ls = float64(_ggadb[0]) / _adeg
		As = float64(_ggadb[1]) / _adeg
		Bs = float64(_ggadb[2]) / _adeg
		Ls = _dg.LinearInterpolate(Ls, 0.0, 1.0, _fbge[0], _fbge[1])
		As = _dg.LinearInterpolate(As, 0.0, 1.0, _fbge[2], _fbge[3])
		Bs = _dg.LinearInterpolate(Bs, 0.0, 1.0, _fbge[4], _fbge[5])
		L = (Ls+16)/116 + As/500
		M = (Ls + 16) / 116
		N = (Ls+16)/116 - Bs/200
		X = _dcbc.WhitePoint[0] * _febe(L)
		Y = _dcbc.WhitePoint[1] * _febe(M)
		Z = _dcbc.WhitePoint[2] * _febe(N)
		_bbeae = 3.240479*X + -1.537150*Y + -0.498535*Z
		_gegee = -0.969256*X + 1.875992*Y + 0.041556*Z
		_dece = 0.055648*X + -0.204043*Y + 1.057311*Z
		_bbeae = _cbg.Min(_cbg.Max(_bbeae, 0), 1.0)
		_gegee = _cbg.Min(_cbg.Max(_gegee, 0), 1.0)
		_dece = _cbg.Min(_cbg.Max(_dece, 0), 1.0)
		_ggadb[0] = uint32(_bbeae * _adeg)
		_ggadb[1] = uint32(_gegee * _adeg)
		_ggadb[2] = uint32(_dece * _adeg)
		if _bfdc = _fca.WriteSamples(_ggadb); _bfdc != nil {
			return img, _bfdc
		}
	}
	return _afacb(&_edeag), nil
}

// ColorToRGB only converts color used with uncolored patterns (defined in underlying colorspace).  Does not go into the
// pattern objects and convert those.  If that is desired, needs to be done separately.  See for example
// grayscale conversion example in unidoc-examples repo.
func (_dbaef *PdfColorspaceSpecialPattern) ColorToRGB(color PdfColor) (PdfColor, error) {
	_fabce, _gfdc := color.(*PdfColorPattern)
	if !_gfdc {
		_eg.Log.Debug("\u0043\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0070a\u0074\u0074\u0065\u0072\u006e\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", color)
		return nil, ErrTypeCheck
	}
	if _fabce.Color == nil {
		return color, nil
	}
	if _dbaef.UnderlyingCS == nil {
		return nil, _gf.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	return _dbaef.UnderlyingCS.ColorToRGB(_fabce.Color)
}

// ToPdfObject returns the PDF representation of the function.
func (_gbebc *PdfFunctionType4) ToPdfObject() _ebb.PdfObject {
	_fdfaa := _gbebc._gbgd
	if _fdfaa == nil {
		_gbebc._gbgd = &_ebb.PdfObjectStream{}
		_fdfaa = _gbebc._gbgd
	}
	_fffb := _ebb.MakeDict()
	_fffb.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _ebb.MakeInteger(4))
	_agaag := &_ebb.PdfObjectArray{}
	for _, _bcggc := range _gbebc.Domain {
		_agaag.Append(_ebb.MakeFloat(_bcggc))
	}
	_fffb.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _agaag)
	_dbaa := &_ebb.PdfObjectArray{}
	for _, _bcgd := range _gbebc.Range {
		_dbaa.Append(_ebb.MakeFloat(_bcgd))
	}
	_fffb.Set("\u0052\u0061\u006eg\u0065", _dbaa)
	if _gbebc._abbbb == nil && _gbebc.Program != nil {
		_gbebc._abbbb = []byte(_gbebc.Program.String())
	}
	_fffb.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ebb.MakeInteger(int64(len(_gbebc._abbbb))))
	_fdfaa.Stream = _gbebc._abbbb
	_fdfaa.PdfObjectDictionary = _fffb
	return _fdfaa
}
func (_acfc *PdfReader) newPdfAnnotationFileAttachmentFromDict(_ggbc *_ebb.PdfObjectDictionary) (*PdfAnnotationFileAttachment, error) {
	_aaa := PdfAnnotationFileAttachment{}
	_ggde, _ggdd := _acfc.newPdfAnnotationMarkupFromDict(_ggbc)
	if _ggdd != nil {
		return nil, _ggdd
	}
	_aaa.PdfAnnotationMarkup = _ggde
	_aaa.FS = _ggbc.Get("\u0046\u0053")
	_aaa.Name = _ggbc.Get("\u004e\u0061\u006d\u0065")
	return &_aaa, nil
}

// NewPdfAppenderWithOpts creates a new Pdf appender from a Pdf reader with options.
func NewPdfAppenderWithOpts(reader *PdfReader, opts *ReaderOpts, encryptOptions *EncryptOptions) (*PdfAppender, error) {
	_eff := &PdfAppender{_ecce: reader._ggdg, Reader: reader, _gege: reader._cafdf, _agb: reader._dfadc}
	_bcgf, _babb := _eff._ecce.Seek(0, _ab.SeekEnd)
	if _babb != nil {
		return nil, _babb
	}
	_eff._bee = _bcgf
	if _, _babb = _eff._ecce.Seek(0, _ab.SeekStart); _babb != nil {
		return nil, _babb
	}
	_eff._acfe, _babb = NewPdfReaderWithOpts(_eff._ecce, opts)
	if _babb != nil {
		return nil, _babb
	}
	for _, _dce := range _eff.Reader.GetObjectNums() {
		if _eff._gbddb < _dce {
			_eff._gbddb = _dce
		}
	}
	_eff._acfd = _eff._gege.GetXrefTable()
	_eff._cfag = _eff._gege.GetXrefOffset()
	_eff._dfbg = append(_eff._dfbg, _eff._acfe.PageList...)
	_eff._ddfg = make(map[_ebb.PdfObject]struct{})
	_eff._gbfa = make(map[_ebb.PdfObject]int64)
	_eff._eebc = make(map[_ebb.PdfObject]struct{})
	_eff._bfef = _eff._acfe.AcroForm
	_eff._eged = _eff._acfe.DSS
	if opts != nil {
		_eff._accg = opts.Password
	}
	if encryptOptions != nil {
		_eff._gfba = encryptOptions
	}
	return _eff, nil
}
func (_ebfae *Image) resampleLowBits(_ccbdd []uint32) {
	_egddc := _dg.BytesPerLine(int(_ebfae.Width), int(_ebfae.BitsPerComponent), _ebfae.ColorComponents)
	_gfec := make([]byte, _ebfae.ColorComponents*_egddc*int(_ebfae.Height))
	_eaffg := int(_ebfae.BitsPerComponent) * _ebfae.ColorComponents * int(_ebfae.Width)
	_aeccf := uint8(8)
	var (
		_dagb, _ffeac int
		_agdgf        uint32
	)
	for _afcgd := 0; _afcgd < int(_ebfae.Height); _afcgd++ {
		_ffeac = _afcgd * _egddc
		for _fcfa := 0; _fcfa < _eaffg; _fcfa++ {
			_agdgf = _ccbdd[_dagb]
			_aeccf -= uint8(_ebfae.BitsPerComponent)
			_gfec[_ffeac] |= byte(_agdgf) << _aeccf
			if _aeccf == 0 {
				_aeccf = 8
				_ffeac++
			}
			_dagb++
		}
	}
	_ebfae.Data = _gfec
}

// NewPdfColorDeviceRGB returns a new PdfColorDeviceRGB based on the r,g,b component values.
func NewPdfColorDeviceRGB(r, g, b float64) *PdfColorDeviceRGB {
	_aggd := PdfColorDeviceRGB{r, g, b}
	return &_aggd
}

// GetPdfVersion gets the version of the PDF used within this document.
func (_ffgcfe *PdfWriter) GetPdfVersion() string { return _ffgcfe.getPdfVersion() }

// GetContentStreamObjs returns a slice of PDF objects containing the content
// streams of the page.
func (_aafg *PdfPage) GetContentStreamObjs() []_ebb.PdfObject {
	if _aafg.Contents == nil {
		return nil
	}
	_bfgd := _ebb.TraceToDirectObject(_aafg.Contents)
	if _aeaba, _eeba := _bfgd.(*_ebb.PdfObjectArray); _eeba {
		return _aeaba.Elements()
	}
	return []_ebb.PdfObject{_bfgd}
}
func _agdbf(_cbffe []*_ebb.PdfObjectStream) *_ebb.PdfObjectArray {
	if len(_cbffe) == 0 {
		return nil
	}
	_gdcfa := make([]_ebb.PdfObject, 0, len(_cbffe))
	for _, _gdcae := range _cbffe {
		_gdcfa = append(_gdcfa, _gdcae)
	}
	return _ebb.MakeArray(_gdcfa...)
}

// NewPdfPageResourcesFromDict creates and returns a new PdfPageResources object
// from the input dictionary.
func NewPdfPageResourcesFromDict(dict *_ebb.PdfObjectDictionary) (*PdfPageResources, error) {
	_afecd := NewPdfPageResources()
	if _fbfdd := dict.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _fbfdd != nil {
		_afecd.ExtGState = _fbfdd
	}
	if _bbdde := dict.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"); _bbdde != nil && !_ebb.IsNullObject(_bbdde) {
		_afecd.ColorSpace = _bbdde
	}
	if _cfgea := dict.Get("\u0050a\u0074\u0074\u0065\u0072\u006e"); _cfgea != nil {
		_afecd.Pattern = _cfgea
	}
	if _cbeef := dict.Get("\u0053h\u0061\u0064\u0069\u006e\u0067"); _cbeef != nil {
		_afecd.Shading = _cbeef
	}
	if _ecedf := dict.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"); _ecedf != nil {
		_afecd.XObject = _ecedf
	}
	if _efffc := _ebb.ResolveReference(dict.Get("\u0046\u006f\u006e\u0074")); _efffc != nil {
		_afecd.Font = _efffc
	}
	if _acdgfc := dict.Get("\u0050r\u006f\u0063\u0053\u0065\u0074"); _acdgfc != nil {
		_afecd.ProcSet = _acdgfc
	}
	if _acdfa := dict.Get("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"); _acdfa != nil {
		_afecd.Properties = _acdfa
	}
	return _afecd, nil
}

// NewPdfColorspaceCalRGB returns a new CalRGB colorspace object.
func NewPdfColorspaceCalRGB() *PdfColorspaceCalRGB {
	_adfe := &PdfColorspaceCalRGB{}
	_adfe.BlackPoint = []float64{0.0, 0.0, 0.0}
	_adfe.Gamma = []float64{1.0, 1.0, 1.0}
	_adfe.Matrix = []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	return _adfe
}

// NewPdfAnnotationText returns a new text annotation.
func NewPdfAnnotationText() *PdfAnnotationText {
	_dca := NewPdfAnnotation()
	_ddf := &PdfAnnotationText{}
	_ddf.PdfAnnotation = _dca
	_ddf.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dca.SetContext(_ddf)
	return _ddf
}

// GetContentStream returns the XObject Form's content stream.
func (_facd *XObjectForm) GetContentStream() ([]byte, error) {
	_gbgbd, _eaea := _ebb.DecodeStream(_facd._gebcd)
	if _eaea != nil {
		return nil, _eaea
	}
	return _gbgbd, nil
}

// SetDSS sets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_dgce *PdfAppender) SetDSS(dss *DSS) {
	if dss != nil {
		_dgce.updateObjectsDeep(dss.ToPdfObject(), nil)
	}
	_dgce._eged = dss
}

// PdfAnnotationStamp represents Stamp annotations.
// (Section 12.5.6.12).
type PdfAnnotationStamp struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Name _ebb.PdfObject
}

// ToPdfObject converts colorspace to a PDF object. [/Indexed base hival lookup]
func (_bcge *PdfColorspaceSpecialIndexed) ToPdfObject() _ebb.PdfObject {
	_bgee := _ebb.MakeArray(_ebb.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"))
	_bgee.Append(_bcge.Base.ToPdfObject())
	_bgee.Append(_ebb.MakeInteger(int64(_bcge.HiVal)))
	_bgee.Append(_bcge.Lookup)
	if _bcge._gdaef != nil {
		_bcge._gdaef.PdfObject = _bgee
		return _bcge._gdaef
	}
	return _bgee
}
func (_cdede *PdfWriter) copyObject(_ffcfb _ebb.PdfObject, _fdfb map[_ebb.PdfObject]_ebb.PdfObject, _gaecbf map[_ebb.PdfObject]struct{}, _fgcaf bool) _ebb.PdfObject {
	_ggedd := !_cdede._abffb && _gaecbf != nil
	if _bcdge, _dabdc := _fdfb[_ffcfb]; _dabdc {
		if _ggedd && !_fgcaf {
			delete(_gaecbf, _ffcfb)
		}
		return _bcdge
	}
	_efbcd := _ffcfb
	switch _addcg := _ffcfb.(type) {
	case *_ebb.PdfObjectArray:
		_fadde := _ebb.MakeArray()
		_efbcd = _fadde
		_fdfb[_ffcfb] = _efbcd
		for _, _gdcbdg := range _addcg.Elements() {
			_fadde.Append(_cdede.copyObject(_gdcbdg, _fdfb, _gaecbf, _fgcaf))
		}
	case *_ebb.PdfObjectStreams:
		_agdcf := &_ebb.PdfObjectStreams{PdfObjectReference: _addcg.PdfObjectReference}
		_efbcd = _agdcf
		_fdfb[_ffcfb] = _efbcd
		for _, _bdacb := range _addcg.Elements() {
			_agdcf.Append(_cdede.copyObject(_bdacb, _fdfb, _gaecbf, _fgcaf))
		}
	case *_ebb.PdfObjectStream:
		_fgfde := &_ebb.PdfObjectStream{Stream: _addcg.Stream, PdfObjectReference: _addcg.PdfObjectReference}
		_efbcd = _fgfde
		_fdfb[_ffcfb] = _efbcd
		_fgfde.PdfObjectDictionary = _cdede.copyObject(_addcg.PdfObjectDictionary, _fdfb, _gaecbf, _fgcaf).(*_ebb.PdfObjectDictionary)
	case *_ebb.PdfObjectDictionary:
		var _aefec bool
		if _ggedd && !_fgcaf {
			if _ggeccg, _ := _ebb.GetNameVal(_addcg.Get("\u0054\u0079\u0070\u0065")); _ggeccg == "\u0050\u0061\u0067\u0065" {
				_, _fffde := _cdede._afbdd[_addcg]
				_fgcaf = !_fffde
				_aefec = _fgcaf
			}
		}
		_eeeaa := _ebb.MakeDict()
		_efbcd = _eeeaa
		_fdfb[_ffcfb] = _efbcd
		for _, _fbbc := range _addcg.Keys() {
			_eeeaa.Set(_fbbc, _cdede.copyObject(_addcg.Get(_fbbc), _fdfb, _gaecbf, _fgcaf))
		}
		if _aefec {
			_efbcd = _ebb.MakeNull()
			_fgcaf = false
		}
	case *_ebb.PdfIndirectObject:
		_gggbb := &_ebb.PdfIndirectObject{PdfObjectReference: _addcg.PdfObjectReference}
		_efbcd = _gggbb
		_fdfb[_ffcfb] = _efbcd
		_gggbb.PdfObject = _cdede.copyObject(_addcg.PdfObject, _fdfb, _gaecbf, _fgcaf)
	case *_ebb.PdfObjectString:
		_dcaec := *_addcg
		_efbcd = &_dcaec
		_fdfb[_ffcfb] = _efbcd
	case *_ebb.PdfObjectName:
		_adfd := *_addcg
		_efbcd = &_adfd
		_fdfb[_ffcfb] = _efbcd
	case *_ebb.PdfObjectNull:
		_efbcd = _ebb.MakeNull()
		_fdfb[_ffcfb] = _efbcd
	case *_ebb.PdfObjectInteger:
		_ggfgc := *_addcg
		_efbcd = &_ggfgc
		_fdfb[_ffcfb] = _efbcd
	case *_ebb.PdfObjectReference:
		_cgbeb := *_addcg
		_efbcd = &_cgbeb
		_fdfb[_ffcfb] = _efbcd
	case *_ebb.PdfObjectFloat:
		_gabd := *_addcg
		_efbcd = &_gabd
		_fdfb[_ffcfb] = _efbcd
	case *_ebb.PdfObjectBool:
		_eaadd := *_addcg
		_efbcd = &_eaadd
		_fdfb[_ffcfb] = _efbcd
	case *pdfSignDictionary:
		_gacbgf := &pdfSignDictionary{PdfObjectDictionary: _ebb.MakeDict(), _dcfab: _addcg._dcfab, _bead: _addcg._bead}
		_efbcd = _gacbgf
		_fdfb[_ffcfb] = _efbcd
		for _, _gabac := range _addcg.Keys() {
			_gacbgf.Set(_gabac, _cdede.copyObject(_addcg.Get(_gabac), _fdfb, _gaecbf, _fgcaf))
		}
	default:
		_eg.Log.Info("\u0054\u004f\u0044\u004f\u0028\u0061\u0035\u0069\u0029\u003a\u0020\u0069\u006dp\u006c\u0065\u006d\u0065\u006e\u0074 \u0063\u006f\u0070\u0079\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0066\u006fr\u0020\u0025\u002b\u0076", _ffcfb)
	}
	if _ggedd && _fgcaf {
		_gaecbf[_ffcfb] = struct{}{}
	}
	return _efbcd
}

// NewPdfAcroForm returns a new PdfAcroForm with an intialized container (indirect object).
func NewPdfAcroForm() *PdfAcroForm {
	return &PdfAcroForm{Fields: &[]*PdfField{}, _adcg: _ebb.MakeIndirectObject(_ebb.MakeDict())}
}

// UpdateObject marks `obj` as updated and to be included in the following revision.
func (_fdfg *PdfAppender) UpdateObject(obj _ebb.PdfObject) {
	_fdfg.replaceObject(obj, obj)
	if _, _caee := _fdfg._ddfg[obj]; !_caee {
		_fdfg._bfeg = append(_fdfg._bfeg, obj)
		_fdfg._ddfg[obj] = struct{}{}
	}
}
func _cfgdf(_acdg *_ebb.PdfObjectDictionary) (*PdfFieldText, error) {
	_babeb := &PdfFieldText{}
	_babeb.DA, _ = _ebb.GetString(_acdg.Get("\u0044\u0041"))
	_babeb.Q, _ = _ebb.GetInt(_acdg.Get("\u0051"))
	_babeb.DS, _ = _ebb.GetString(_acdg.Get("\u0044\u0053"))
	_babeb.RV = _acdg.Get("\u0052\u0056")
	_babeb.MaxLen, _ = _ebb.GetInt(_acdg.Get("\u004d\u0061\u0078\u004c\u0065\u006e"))
	return _babeb, nil
}

// ToPdfObject implements interface PdfModel.
func (_dgfa *PdfTransformParamsDocMDP) ToPdfObject() _ebb.PdfObject {
	_aefdg := _ebb.MakeDict()
	_aefdg.SetIfNotNil("\u0054\u0079\u0070\u0065", _dgfa.Type)
	_aefdg.SetIfNotNil("\u0056", _dgfa.V)
	_aefdg.SetIfNotNil("\u0050", _dgfa.P)
	return _aefdg
}
func (_bbadg *PdfWriter) checkPendingObjects() {
	for _gbfff, _gcedb := range _bbadg._eefeb {
		if !_bbadg.hasObject(_gbfff) {
			_eg.Log.Debug("\u0057\u0041\u0052\u004e\u0020\u0050\u0065n\u0064\u0069\u006eg\u0020\u006f\u0062j\u0065\u0063t\u0020\u0025\u002b\u0076\u0020\u0025T\u0020(%\u0070\u0029\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0077\u0072\u0069\u0074\u0069\u006e\u0067", _gbfff, _gbfff, _gbfff)
			for _, _dggdfbc := range _gcedb {
				for _, _efcfc := range _dggdfbc.Keys() {
					_babdf := _dggdfbc.Get(_efcfc)
					if _babdf == _gbfff {
						_eg.Log.Debug("\u0050e\u006e\u0064i\u006e\u0067\u0020\u006fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0061nd\u0020\u0072\u0065p\u006c\u0061c\u0065\u0064\u0020\u0077\u0069\u0074h\u0020\u006eu\u006c\u006c")
						_dggdfbc.Set(_efcfc, _ebb.MakeNull())
						break
					}
				}
			}
		}
	}
}

// CharMetrics represents width and height metrics of a glyph.
type CharMetrics = _bad.CharMetrics

func (_fgca *PdfSignature) extractChainFromCert() ([]*_g.Certificate, error) {
	var _cafgga *_ebb.PdfObjectArray
	switch _fffad := _fgca.Cert.(type) {
	case *_ebb.PdfObjectString:
		_cafgga = _ebb.MakeArray(_fffad)
	case *_ebb.PdfObjectArray:
		_cafgga = _fffad
	default:
		return nil, _bg.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _fffad)
	}
	var _bedcd _ca.Buffer
	for _, _ccdga := range _cafgga.Elements() {
		_fgcac, _beba := _ebb.GetString(_ccdga)
		if !_beba {
			return nil, _bg.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054", _ccdga)
		}
		if _, _bcfae := _bedcd.Write(_fgcac.Bytes()); _bcfae != nil {
			return nil, _bcfae
		}
	}
	return _g.ParseCertificates(_bedcd.Bytes())
}
func (_ggbe *PdfReader) newPdfAnnotationPolyLineFromDict(_gbac *_ebb.PdfObjectDictionary) (*PdfAnnotationPolyLine, error) {
	_acggf := PdfAnnotationPolyLine{}
	_dgfd, _bcb := _ggbe.newPdfAnnotationMarkupFromDict(_gbac)
	if _bcb != nil {
		return nil, _bcb
	}
	_acggf.PdfAnnotationMarkup = _dgfd
	_acggf.Vertices = _gbac.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_acggf.LE = _gbac.Get("\u004c\u0045")
	_acggf.BS = _gbac.Get("\u0042\u0053")
	_acggf.IC = _gbac.Get("\u0049\u0043")
	_acggf.BE = _gbac.Get("\u0042\u0045")
	_acggf.IT = _gbac.Get("\u0049\u0054")
	_acggf.Measure = _gbac.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_acggf, nil
}
func (_gceb *PdfReader) buildPageList(_gdafg *_ebb.PdfIndirectObject, _cbaadb *_ebb.PdfIndirectObject, _gbggf map[_ebb.PdfObject]struct{}) error {
	if _gdafg == nil {
		return nil
	}
	if _, _dceef := _gbggf[_gdafg]; _dceef {
		_eg.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _gdafg.ObjectNumber)
		return nil
	}
	_gbggf[_gdafg] = struct{}{}
	_bcaaa, _cbegc := _gdafg.PdfObject.(*_ebb.PdfObjectDictionary)
	if !_cbegc {
		return _gf.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_afcabe, _cbegc := (*_bcaaa).Get("\u0054\u0079\u0070\u0065").(*_ebb.PdfObjectName)
	if !_cbegc {
		if _bcaaa.Get("\u004b\u0069\u0064\u0073") == nil {
			return _gf.New("\u006e\u006f\u0064\u0065 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0054\u0079p\u0065 \u0028\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0029")
		}
		_eg.Log.Debug("ER\u0052\u004fR\u003a\u0020\u006e\u006f\u0064\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067\u0020\u0054\u0079\u0070\u0065\u002c\u0020\u0062\u0075\u0074\u0020\u0068\u0061\u0073\u0020\u004b\u0069\u0064\u0073\u002e\u0020\u0041\u0073\u0073u\u006di\u006e\u0067\u0020\u0050\u0061\u0067\u0065\u0073 \u006eo\u0064\u0065.")
		_afcabe = _ebb.MakeName("\u0050\u0061\u0067e\u0073")
		_bcaaa.Set("\u0054\u0079\u0070\u0065", _afcabe)
	}
	_eg.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0050a\u0067\u0065\u004c\u0069\u0073\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0074y\u0070\u0065\u003a\u0020\u0025\u0073\u0020(\u0025\u002b\u0076\u0029", *_afcabe, _gdafg)
	if *_afcabe == "\u0050\u0061\u0067\u0065" {
		_bgfade, _efbc := _gceb.newPdfPageFromDict(_bcaaa)
		if _efbc != nil {
			return _efbc
		}
		_bgfade.setContainer(_gdafg)
		if _cbaadb != nil {
			_bcaaa.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _cbaadb)
		}
		_gceb._faebb = append(_gceb._faebb, _gdafg)
		_gceb.PageList = append(_gceb.PageList, _bgfade)
		return nil
	}
	if *_afcabe != "\u0050\u0061\u0067e\u0073" {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0063\u006fnt\u0065n\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u0021\u0020\u0028\u0025\u0073\u0029", _afcabe)
		return _gf.New("\u0074\u0061\u0062\u006c\u0065\u0020o\u0066\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073 \u006fb\u006a\u0065\u0063\u0074")
	}
	if _cbaadb != nil {
		_bcaaa.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _cbaadb)
	}
	if !_gceb._ceefa {
		_fgcca := _gceb.traverseObjectData(_gdafg)
		if _fgcca != nil {
			return _fgcca
		}
	}
	_gfede, _dddce := _gceb._cafdf.Resolve(_bcaaa.Get("\u004b\u0069\u0064\u0073"))
	if _dddce != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u004b\u0069\u0064\u0073\u0020\u006fb\u006a\u0065\u0063\u0074")
		return _dddce
	}
	var _fadfc *_ebb.PdfObjectArray
	_fadfc, _cbegc = _gfede.(*_ebb.PdfObjectArray)
	if !_cbegc {
		_cbgge, _efab := _gfede.(*_ebb.PdfIndirectObject)
		if !_efab {
			return _gf.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_fadfc, _cbegc = _cbgge.PdfObject.(*_ebb.PdfObjectArray)
		if !_cbegc {
			return _gf.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u0069\u006ed\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
	}
	_eg.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _fadfc)
	for _affge, _edcbbc := range _fadfc.Elements() {
		_edefe, _fdfaf := _ebb.GetIndirect(_edcbbc)
		if !_fdfaf {
			_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 \u006f\u0062\u006a\u0065\u0063t\u0020\u002d \u0028\u0025\u0073\u0029", _edefe)
			return _gf.New("\u0070a\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_fadfc.Set(_affge, _edefe)
		_dddce = _gceb.buildPageList(_edefe, _gdafg, _gbggf)
		if _dddce != nil {
			return _dddce
		}
	}
	return nil
}

// ImageToRGB converts CalRGB colorspace image to RGB and returns the result.
func (_eaed *PdfColorspaceCalRGB) ImageToRGB(img Image) (Image, error) {
	_dcba := _abg.NewReader(img.getBase())
	_acdc := _dg.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_ffd := _abg.NewWriter(_acdc)
	_gebe := _cbg.Pow(2, float64(img.BitsPerComponent)) - 1
	_gbca := make([]uint32, 3)
	var (
		_fbaa                                      error
		_eeeg, _aee, _bgeaf, _eeddbe, _gcfg, _feff float64
	)
	for {
		_fbaa = _dcba.ReadSamples(_gbca)
		if _fbaa == _ab.EOF {
			break
		} else if _fbaa != nil {
			return img, _fbaa
		}
		_eeeg = float64(_gbca[0]) / _gebe
		_aee = float64(_gbca[1]) / _gebe
		_bgeaf = float64(_gbca[2]) / _gebe
		_eeddbe = _eaed.Matrix[0]*_cbg.Pow(_eeeg, _eaed.Gamma[0]) + _eaed.Matrix[3]*_cbg.Pow(_aee, _eaed.Gamma[1]) + _eaed.Matrix[6]*_cbg.Pow(_bgeaf, _eaed.Gamma[2])
		_gcfg = _eaed.Matrix[1]*_cbg.Pow(_eeeg, _eaed.Gamma[0]) + _eaed.Matrix[4]*_cbg.Pow(_aee, _eaed.Gamma[1]) + _eaed.Matrix[7]*_cbg.Pow(_bgeaf, _eaed.Gamma[2])
		_feff = _eaed.Matrix[2]*_cbg.Pow(_eeeg, _eaed.Gamma[0]) + _eaed.Matrix[5]*_cbg.Pow(_aee, _eaed.Gamma[1]) + _eaed.Matrix[8]*_cbg.Pow(_bgeaf, _eaed.Gamma[2])
		_eeeg = 3.240479*_eeddbe + -1.537150*_gcfg + -0.498535*_feff
		_aee = -0.969256*_eeddbe + 1.875992*_gcfg + 0.041556*_feff
		_bgeaf = 0.055648*_eeddbe + -0.204043*_gcfg + 1.057311*_feff
		_eeeg = _cbg.Min(_cbg.Max(_eeeg, 0), 1.0)
		_aee = _cbg.Min(_cbg.Max(_aee, 0), 1.0)
		_bgeaf = _cbg.Min(_cbg.Max(_bgeaf, 0), 1.0)
		_gbca[0] = uint32(_eeeg * _gebe)
		_gbca[1] = uint32(_aee * _gebe)
		_gbca[2] = uint32(_bgeaf * _gebe)
		if _fbaa = _ffd.WriteSamples(_gbca); _fbaa != nil {
			return img, _fbaa
		}
	}
	return _afacb(&_acdc), nil
}

// AppendContentStream adds content stream by string.  Appends to the last
// contentstream instance if many.
func (_bgec *PdfPage) AppendContentStream(contentStr string) error {
	_cfae, _cdadf := _bgec.GetContentStreams()
	if _cdadf != nil {
		return _cdadf
	}
	if len(_cfae) == 0 {
		_cfae = []string{contentStr}
		return _bgec.SetContentStreams(_cfae, _ebb.NewFlateEncoder())
	}
	var _fdec _ca.Buffer
	_fdec.WriteString(_cfae[len(_cfae)-1])
	_fdec.WriteString("\u000a")
	_fdec.WriteString(contentStr)
	_cfae[len(_cfae)-1] = _fdec.String()
	return _bgec.SetContentStreams(_cfae, _ebb.NewFlateEncoder())
}
func (_gdedc *PdfWriter) addObject(_cdcgd _ebb.PdfObject) bool {
	_bdfga := _gdedc.hasObject(_cdcgd)
	if !_bdfga {
		_aeaa := _ebb.ResolveReferencesDeep(_cdcgd, _gdedc._dcfg)
		if _aeaa != nil {
			_eg.Log.Debug("E\u0052R\u004f\u0052\u003a\u0020\u0025\u0076\u0020\u002d \u0073\u006b\u0069\u0070pi\u006e\u0067", _aeaa)
		}
		_gdedc._ebdgg = append(_gdedc._ebdgg, _cdcgd)
		_gdedc._ffffd[_cdcgd] = struct{}{}
		return true
	}
	return false
}

// AllFields returns a flattened list of all fields in the form.
func (_cgbbga *PdfAcroForm) AllFields() []*PdfField {
	if _cgbbga == nil {
		return nil
	}
	var _baad []*PdfField
	if _cgbbga.Fields != nil {
		for _, _ebgb := range *_cgbbga.Fields {
			_baad = append(_baad, _fbea(_ebgb)...)
		}
	}
	return _baad
}

// PdfColorspaceDeviceGray represents a grayscale colorspace.
type PdfColorspaceDeviceGray struct{}

// PdfTilingPattern is a Tiling pattern that consists of repetitions of a pattern cell with defined intervals.
// It is a type 1 pattern. (PatternType = 1).
// A tiling pattern is represented by a stream object, where the stream content is
// a content stream that describes the pattern cell.
type PdfTilingPattern struct {
	*PdfPattern
	PaintType  *_ebb.PdfObjectInteger
	TilingType *_ebb.PdfObjectInteger
	BBox       *PdfRectangle
	XStep      *_ebb.PdfObjectFloat
	YStep      *_ebb.PdfObjectFloat
	Resources  *PdfPageResources
	Matrix     *_ebb.PdfObjectArray
}

func (_bdbc *PdfReader) newPdfAnnotationFreeTextFromDict(_ababc *_ebb.PdfObjectDictionary) (*PdfAnnotationFreeText, error) {
	_ccbf := PdfAnnotationFreeText{}
	_adfb, _fdc := _bdbc.newPdfAnnotationMarkupFromDict(_ababc)
	if _fdc != nil {
		return nil, _fdc
	}
	_ccbf.PdfAnnotationMarkup = _adfb
	_ccbf.DA = _ababc.Get("\u0044\u0041")
	_ccbf.Q = _ababc.Get("\u0051")
	_ccbf.RC = _ababc.Get("\u0052\u0043")
	_ccbf.DS = _ababc.Get("\u0044\u0053")
	_ccbf.CL = _ababc.Get("\u0043\u004c")
	_ccbf.IT = _ababc.Get("\u0049\u0054")
	_ccbf.BE = _ababc.Get("\u0042\u0045")
	_ccbf.RD = _ababc.Get("\u0052\u0044")
	_ccbf.BS = _ababc.Get("\u0042\u0053")
	_ccbf.LE = _ababc.Get("\u004c\u0045")
	return &_ccbf, nil
}

// ToPdfObject implements interface PdfModel.
func (_efb *PdfActionResetForm) ToPdfObject() _ebb.PdfObject {
	_efb.PdfAction.ToPdfObject()
	_aeb := _efb._abe
	_bcf := _aeb.PdfObject.(*_ebb.PdfObjectDictionary)
	_bcf.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeResetForm)))
	_bcf.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _efb.Fields)
	_bcf.SetIfNotNil("\u0046\u006c\u0061g\u0073", _efb.Flags)
	return _aeb
}

var ErrColorOutOfRange = _gf.New("\u0063o\u006co\u0072\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")

// ToPdfObject implements interface PdfModel.
func (_edae *PdfAnnotationStrikeOut) ToPdfObject() _ebb.PdfObject {
	_edae.PdfAnnotation.ToPdfObject()
	_ffeag := _edae._bdcd
	_abad := _ffeag.PdfObject.(*_ebb.PdfObjectDictionary)
	_edae.PdfAnnotationMarkup.appendToPdfDictionary(_abad)
	_abad.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _ebb.MakeName("\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t"))
	_abad.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _edae.QuadPoints)
	return _ffeag
}

// CheckAccessRights checks access rights and permissions for a specified password.  If either user/owner
// password is specified,  full rights are granted, otherwise the access rights are specified by the
// Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_eceg *PdfReader) CheckAccessRights(password []byte) (bool, _fe.Permissions, error) {
	return _eceg._cafdf.CheckAccessRights(password)
}
func _dddfe(_gbecb *XObjectImage) error {
	if _gbecb.SMask == nil {
		return nil
	}
	_gfdag, _gddeb := _gbecb.SMask.(*_ebb.PdfObjectStream)
	if !_gddeb {
		_eg.Log.Debug("\u0053\u004da\u0073\u006b\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0053\u0074re\u0061\u006d")
		return _ebb.ErrTypeError
	}
	_acgce := _gfdag.PdfObjectDictionary
	_bbbga := _acgce.Get("\u004d\u0061\u0074t\u0065")
	if _bbbga == nil {
		return nil
	}
	_agaad, _eeddd := _ffaba(_bbbga.(*_ebb.PdfObjectArray))
	if _eeddd != nil {
		return _eeddd
	}
	_dbgcbc := _ebb.MakeArrayFromFloats([]float64{_agaad})
	_acgce.SetIfNotNil("\u004d\u0061\u0074t\u0065", _dbgcbc)
	return nil
}
func _gcabc(_aacea *_ebb.PdfObjectDictionary) (*PdfShadingType6, error) {
	_bgac := PdfShadingType6{}
	_dfef := _aacea.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _dfef == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_ebgaa, _fcffdb := _dfef.(*_ebb.PdfObjectInteger)
	if !_fcffdb {
		_eg.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _dfef)
		return nil, _ebb.ErrTypeError
	}
	_bgac.BitsPerCoordinate = _ebgaa
	_dfef = _aacea.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _dfef == nil {
		_eg.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_ebgaa, _fcffdb = _dfef.(*_ebb.PdfObjectInteger)
	if !_fcffdb {
		_eg.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _dfef)
		return nil, _ebb.ErrTypeError
	}
	_bgac.BitsPerComponent = _ebgaa
	_dfef = _aacea.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _dfef == nil {
		_eg.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_ebgaa, _fcffdb = _dfef.(*_ebb.PdfObjectInteger)
	if !_fcffdb {
		_eg.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _dfef)
		return nil, _ebb.ErrTypeError
	}
	_bgac.BitsPerComponent = _ebgaa
	_dfef = _aacea.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _dfef == nil {
		_eg.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_dcaea, _fcffdb := _dfef.(*_ebb.PdfObjectArray)
	if !_fcffdb {
		_eg.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _dfef)
		return nil, _ebb.ErrTypeError
	}
	_bgac.Decode = _dcaea
	if _gabf := _aacea.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _gabf != nil {
		_bgac.Function = []PdfFunction{}
		if _gabgc, _feebd := _gabf.(*_ebb.PdfObjectArray); _feebd {
			for _, _cgaded := range _gabgc.Elements() {
				_abebg, _ffefg := _aagg(_cgaded)
				if _ffefg != nil {
					_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _ffefg)
					return nil, _ffefg
				}
				_bgac.Function = append(_bgac.Function, _abebg)
			}
		} else {
			_ggbgb, _gfee := _aagg(_gabf)
			if _gfee != nil {
				_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _gfee)
				return nil, _gfee
			}
			_bgac.Function = append(_bgac.Function, _ggbgb)
		}
	}
	return &_bgac, nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_bbggb *PdfColorspaceDeviceGray) ToPdfObject() _ebb.PdfObject {
	return _ebb.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079")
}
func (_eedbf *LTV) validateSig(_cgfga *PdfSignature) error {
	if _cgfga == nil || _cgfga.Contents == nil || len(_cgfga.Contents.Bytes()) == 0 {
		return _bg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0066\u0069\u0065l\u0064:\u0020\u0025\u0076", _cgfga)
	}
	return nil
}

// PdfBorderEffect represents a PDF border effect.
type PdfBorderEffect struct {
	S *BorderEffect
	I *float64
}

// WriteToFile writes the Appender output to file specified by path.
func (_caba *PdfAppender) WriteToFile(outputPath string) error {
	_gggdb, _edca := _ed.Create(outputPath)
	if _edca != nil {
		return _edca
	}
	defer _gggdb.Close()
	return _caba.Write(_gggdb)
}
func _dcbd(_ffcfea _ab.ReadSeeker, _abaee *ReaderOpts, _dfdac bool, _dgabf string) (*PdfReader, error) {
	if _abaee == nil {
		_abaee = NewReaderOpts()
	}
	_aeebf := *_abaee
	_edeaec := &PdfReader{_ggdg: _ffcfea, _dfadc: map[_ebb.PdfObject]struct{}{}, _abbaca: _fadcd(), _ceefa: _abaee.LazyLoad, _cfbgga: _abaee.ComplianceMode, _abadec: _dfdac, _edcbc: &_aeebf}
	_ffcge, _aacce := _bafec("\u0072")
	if _aacce != nil {
		return nil, _aacce
	}
	_edeaec._decdd = _ffcge
	var _bgbeb *_ebb.PdfParser
	if !_edeaec._cfbgga {
		_bgbeb, _aacce = _ebb.NewParser(_ffcfea)
	} else {
		_bgbeb, _aacce = _ebb.NewCompliancePdfParser(_ffcfea)
	}
	if _aacce != nil {
		return nil, _aacce
	}
	_edeaec._cafdf = _bgbeb
	_agafc, _aacce := _edeaec.IsEncrypted()
	if _aacce != nil {
		return nil, _aacce
	}
	if !_agafc {
		_aacce = _edeaec.loadStructure()
		if _aacce != nil {
			return nil, _aacce
		}
	} else if _dfdac {
		_efddf, _cfafd := _edeaec.Decrypt([]byte(_abaee.Password))
		if _cfafd != nil {
			return nil, _cfafd
		}
		if !_efddf {
			return nil, _gf.New("\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0064\u0065c\u0072\u0079\u0070\u0074\u0020\u0070\u0061\u0073\u0073w\u006f\u0072\u0064\u0020p\u0072\u006f\u0074\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u002d\u0020\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0073\u0070\u0065\u0063\u0069\u0066y\u0020\u0070\u0061s\u0073\u0020\u0074\u006f\u0020\u0044\u0065\u0063\u0072\u0079\u0070\u0074")
		}
	}
	_edeaec._cacfc = make(map[*PdfReader]*PdfReader)
	_edeaec._face = make([]*PdfReader, _bgbeb.GetRevisionNumber())
	return _edeaec, nil
}

// IsPush returns true if the button field represents a push button, false otherwise.
func (_abdga *PdfFieldButton) IsPush() bool { return _abdga.GetType() == ButtonTypePush }

// NewPdfFontFromTTFFile loads a TTF font file and returns a PdfFont type
// that can be used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTFFile.
func NewPdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_egbd, _cbfbf := _ed.Open(filePath)
	if _cbfbf != nil {
		_eg.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020T\u0054F\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _cbfbf)
		return nil, _cbfbf
	}
	defer _egbd.Close()
	return NewPdfFontFromTTF(_egbd)
}

// VariableText contains the common attributes of a variable text.
// The VariableText is typically not used directly, but is can encapsulate by PdfField
// See section 12.7.3.3 "Variable Text" and Table 222 (pp. 434-436 PDF32000_2008).
type VariableText struct {
	DA *_ebb.PdfObjectString
	Q  *_ebb.PdfObjectInteger
	DS *_ebb.PdfObjectString
	RV _ebb.PdfObject
}

func (_cedgbg *PdfFont) baseFields() *fontCommon {
	if _cedgbg._ebcad == nil {
		_eg.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0062\u0061\u0073\u0065\u0046\u0069\u0065l\u0064s\u002e \u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e")
		return nil
	}
	return _cedgbg._ebcad.baseFields()
}

// PdfActionGoTo represents a GoTo action.
type PdfActionGoTo struct {
	*PdfAction
	D _ebb.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_ead *PdfActionSetOCGState) ToPdfObject() _ebb.PdfObject {
	_ead.PdfAction.ToPdfObject()
	_bea := _ead._abe
	_gff := _bea.PdfObject.(*_ebb.PdfObjectDictionary)
	_gff.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeSetOCGState)))
	_gff.SetIfNotNil("\u0053\u0074\u0061t\u0065", _ead.State)
	_gff.SetIfNotNil("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042", _ead.PreserveRB)
	return _bea
}

// ToPdfObject implements interface PdfModel.
func (_dae *PdfActionURI) ToPdfObject() _ebb.PdfObject {
	_dae.PdfAction.ToPdfObject()
	_cbc := _dae._abe
	_aeg := _cbc.PdfObject.(*_ebb.PdfObjectDictionary)
	_aeg.SetIfNotNil("\u0053", _ebb.MakeName(string(ActionTypeURI)))
	_aeg.SetIfNotNil("\u0055\u0052\u0049", _dae.URI)
	_aeg.SetIfNotNil("\u0049\u0073\u004da\u0070", _dae.IsMap)
	return _cbc
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_dfea *PdfOutlineItem) ToPdfObject() _ebb.PdfObject {
	_cbacc := _dfea._cacdf
	_fbddd := _cbacc.PdfObject.(*_ebb.PdfObjectDictionary)
	_fbddd.Set("\u0054\u0069\u0074l\u0065", _dfea.Title)
	if _dfea.A != nil {
		_fbddd.Set("\u0041", _dfea.A)
	}
	if _gdgac := _fbddd.Get("\u0053\u0045"); _gdgac != nil {
		_fbddd.Remove("\u0053\u0045")
	}
	if _dfea.C != nil {
		_fbddd.Set("\u0043", _dfea.C)
	}
	if _dfea.Dest != nil {
		_fbddd.Set("\u0044\u0065\u0073\u0074", _dfea.Dest)
	}
	if _dfea.F != nil {
		_fbddd.Set("\u0046", _dfea.F)
	}
	if _dfea.Count != nil {
		_fbddd.Set("\u0043\u006f\u0075n\u0074", _ebb.MakeInteger(*_dfea.Count))
	}
	if _dfea.Next != nil {
		_fbddd.Set("\u004e\u0065\u0078\u0074", _dfea.Next.ToPdfObject())
	}
	if _dfea.First != nil {
		_fbddd.Set("\u0046\u0069\u0072s\u0074", _dfea.First.ToPdfObject())
	}
	if _dfea.Prev != nil {
		_fbddd.Set("\u0050\u0072\u0065\u0076", _dfea.Prev.GetContext().GetContainingPdfObject())
	}
	if _dfea.Last != nil {
		_fbddd.Set("\u004c\u0061\u0073\u0074", _dfea.Last.GetContext().GetContainingPdfObject())
	}
	if _dfea.Parent != nil {
		_fbddd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _dfea.Parent.GetContext().GetContainingPdfObject())
	}
	return _cbacc
}
func _bcafa(_cfgda *_ebb.PdfObjectDictionary) (*PdfShadingType3, error) {
	_fcbbg := PdfShadingType3{}
	_eaaecg := _cfgda.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _eaaecg == nil {
		_eg.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0043\u006f\u006f\u0072\u0064\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_eagb, _gdfeg := _eaaecg.(*_ebb.PdfObjectArray)
	if !_gdfeg {
		_eg.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _eaaecg)
		return nil, _ebb.ErrTypeError
	}
	if _eagb.Len() != 6 {
		_eg.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0036\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _eagb.Len())
		return nil, ErrInvalidAttribute
	}
	_fcbbg.Coords = _eagb
	if _faagf := _cfgda.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _faagf != nil {
		_faagf = _ebb.TraceToDirectObject(_faagf)
		_acffd, _cgfbe := _faagf.(*_ebb.PdfObjectArray)
		if !_cgfbe {
			_eg.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _faagf)
			return nil, _ebb.ErrTypeError
		}
		_fcbbg.Domain = _acffd
	}
	_eaaecg = _cfgda.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _eaaecg == nil {
		_eg.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_fcbbg.Function = []PdfFunction{}
	if _gdbdc, _afegbb := _eaaecg.(*_ebb.PdfObjectArray); _afegbb {
		for _, _dbgd := range _gdbdc.Elements() {
			_agfb, _dcdcd := _aagg(_dbgd)
			if _dcdcd != nil {
				_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dcdcd)
				return nil, _dcdcd
			}
			_fcbbg.Function = append(_fcbbg.Function, _agfb)
		}
	} else {
		_fabbd, _dfafbe := _aagg(_eaaecg)
		if _dfafbe != nil {
			_eg.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dfafbe)
			return nil, _dfafbe
		}
		_fcbbg.Function = append(_fcbbg.Function, _fabbd)
	}
	if _facbe := _cfgda.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _facbe != nil {
		_facbe = _ebb.TraceToDirectObject(_facbe)
		_bffad, _ceda := _facbe.(*_ebb.PdfObjectArray)
		if !_ceda {
			_eg.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _facbe)
			return nil, _ebb.ErrTypeError
		}
		if _bffad.Len() != 2 {
			_eg.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _bffad.Len())
			return nil, ErrInvalidAttribute
		}
		_fcbbg.Extend = _bffad
	}
	return &_fcbbg, nil
}

type fontCommon struct {
	_fdacg string
	_dfbf  string
	_efge  string
	_baag  _ebb.PdfObject
	_dcdd  *_ebe.CMap
	_fbbd  *PdfFontDescriptor
	_efbg  int64
}

// PdfAnnotationSquiggly represents Squiggly annotations.
// (Section 12.5.6.10).
type PdfAnnotationSquiggly struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _ebb.PdfObject
}

// PdfVersion returns version of the PDF file.
func (_fedf *PdfReader) PdfVersion() _ebb.Version { return _fedf._cafdf.PdfVersion() }

// ColorFromPdfObjects gets the color from a series of pdf objects (3 for rgb).
func (_aebde *PdfColorspaceDeviceRGB) ColorFromPdfObjects(objects []_ebb.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _gf.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gggb, _abeed := _ebb.GetNumbersAsFloat(objects)
	if _abeed != nil {
		return nil, _abeed
	}
	return _aebde.ColorFromFloats(_gggb)
}

// NewPermissions returns a new permissions object.
func NewPermissions(docMdp *PdfSignature) *Permissions {
	_cbdg := Permissions{}
	_cbdg.DocMDP = docMdp
	_gbbf := _ebb.MakeDict()
	_gbbf.Set("\u0044\u006f\u0063\u004d\u0044\u0050", docMdp.ToPdfObject())
	_cbdg._bdgdgf = _gbbf
	return &_cbdg
}

// ToInteger convert to an integer format.
func (_ggdee *PdfColorDeviceCMYK) ToInteger(bits int) [4]uint32 {
	_fced := _cbg.Pow(2, float64(bits)) - 1
	return [4]uint32{uint32(_fced * _ggdee.C()), uint32(_fced * _ggdee.M()), uint32(_fced * _ggdee.Y()), uint32(_fced * _ggdee.K())}
}

// String returns the name of the colorspace (DeviceN).
func (_cbdb *PdfColorspaceDeviceN) String() string { return "\u0044e\u0076\u0069\u0063\u0065\u004e" }

// Height returns the height of `rect`.
func (_bcdeb *PdfRectangle) Height() float64 { return _cbg.Abs(_bcdeb.Ury - _bcdeb.Lly) }
