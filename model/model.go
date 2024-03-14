package model

import (
	_ac "bufio"
	_dd "bytes"
	_ag "crypto/md5"
	_g "crypto/rand"
	_eg "crypto/sha1"
	_fa "crypto/x509"
	_bg "encoding/binary"
	_cb "encoding/hex"
	_fd "errors"
	_e "fmt"
	_a "hash"
	_aa "image"
	_ga "image/color"
	_ "image/gif"
	_ "image/png"
	_gc "io"
	_ge "math"
	_aaf "math/rand"
	_cf "os"
	_af "regexp"
	_bb "sort"
	_gb "strconv"
	_be "strings"
	_c "sync"
	_f "time"
	_gg "unicode"
	_bc "unicode/utf8"

	_acd "bitbucket.org/shenghui0779/gopdf/common"
	_abf "bitbucket.org/shenghui0779/gopdf/core"
	_bga "bitbucket.org/shenghui0779/gopdf/core/security"
	_bf "bitbucket.org/shenghui0779/gopdf/core/security/crypt"
	_bd "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_gca "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_gf "bitbucket.org/shenghui0779/gopdf/internal/sampling"
	_cbb "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_fae "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_ad "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_bbf "bitbucket.org/shenghui0779/gopdf/model/internal/docutil"
	_gbe "bitbucket.org/shenghui0779/gopdf/model/internal/fonts"
	_df "bitbucket.org/shenghui0779/gopdf/model/mdp"
	_fe "bitbucket.org/shenghui0779/gopdf/model/sigutil"
	_ae "bitbucket.org/shenghui0779/gopdf/ps"
	_eb "github.com/unidoc/pkcs7"
	_ab "github.com/unidoc/unitype"
	_ddd "golang.org/x/xerrors"
)

// PdfModel is a higher level PDF construct which can be collapsed into a PdfObject.
// Each PdfModel has an underlying PdfObject and vice versa (one-to-one).
// Under normal circumstances there should only be one copy of each.
// Copies can be made, but care must be taken to do it properly.
type PdfModel interface {
	ToPdfObject() _abf.PdfObject
	GetContainingPdfObject() _abf.PdfObject
}

// GetContext returns the context of the outline tree node, which is either a
// *PdfOutline or a *PdfOutlineItem. The method returns nil for uninitialized
// tree nodes.
func (_bdgca *PdfOutlineTreeNode) GetContext() PdfModel {
	if _bfcaf, _effdd := _bdgca._aecec.(*PdfOutline); _effdd {
		return _bfcaf
	}
	if _cffgg, _eafc := _bdgca._aecec.(*PdfOutlineItem); _eafc {
		return _cffgg
	}
	_acd.Log.Debug("\u0045\u0052RO\u0052\u0020\u0049n\u0076\u0061\u006c\u0069d o\u0075tl\u0069\u006e\u0065\u0020\u0074\u0072\u0065e \u006e\u006f\u0064\u0065\u0020\u0069\u0074e\u006d")
	return nil
}

// ToPdfOutlineItem returns a low level PdfOutlineItem object,
// based on the current instance.
func (_adae *OutlineItem) ToPdfOutlineItem() (*PdfOutlineItem, int64) {
	_dbbde := NewPdfOutlineItem()
	_dbbde.Title = _abf.MakeEncodedString(_adae.Title, true)
	_dbbde.Dest = _adae.Dest.ToPdfObject()
	var _ebefdc []*PdfOutlineItem
	var _cbfa int64
	var _gdfad *PdfOutlineItem
	for _, _geebc := range _adae.Entries {
		_gcdf, _fgff := _geebc.ToPdfOutlineItem()
		_gcdf.Parent = &_dbbde.PdfOutlineTreeNode
		if _gdfad != nil {
			_gdfad.Next = &_gcdf.PdfOutlineTreeNode
			_gcdf.Prev = &_gdfad.PdfOutlineTreeNode
		}
		_ebefdc = append(_ebefdc, _gcdf)
		_cbfa += _fgff
		_gdfad = _gcdf
	}
	_beaae := len(_ebefdc)
	_cbfa += int64(_beaae)
	if _beaae > 0 {
		_dbbde.First = &_ebefdc[0].PdfOutlineTreeNode
		_dbbde.Last = &_ebefdc[_beaae-1].PdfOutlineTreeNode
		_dbbde.Count = &_cbfa
	}
	return _dbbde, _cbfa
}

// ImageToRGB converts CalRGB colorspace image to RGB and returns the result.
func (_bdeg *PdfColorspaceCalRGB) ImageToRGB(img Image) (Image, error) {
	_cgcb := _gf.NewReader(img.getBase())
	_fafe := _gca.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_bebf := _gf.NewWriter(_fafe)
	_gdcg := _ge.Pow(2, float64(img.BitsPerComponent)) - 1
	_dbbb := make([]uint32, 3)
	var (
		_ggbcc                                     error
		_cebe, _cagbe, _abcg, _fbgc, _bgadc, _gafb float64
	)
	for {
		_ggbcc = _cgcb.ReadSamples(_dbbb)
		if _ggbcc == _gc.EOF {
			break
		} else if _ggbcc != nil {
			return img, _ggbcc
		}
		_cebe = float64(_dbbb[0]) / _gdcg
		_cagbe = float64(_dbbb[1]) / _gdcg
		_abcg = float64(_dbbb[2]) / _gdcg
		_fbgc = _bdeg.Matrix[0]*_ge.Pow(_cebe, _bdeg.Gamma[0]) + _bdeg.Matrix[3]*_ge.Pow(_cagbe, _bdeg.Gamma[1]) + _bdeg.Matrix[6]*_ge.Pow(_abcg, _bdeg.Gamma[2])
		_bgadc = _bdeg.Matrix[1]*_ge.Pow(_cebe, _bdeg.Gamma[0]) + _bdeg.Matrix[4]*_ge.Pow(_cagbe, _bdeg.Gamma[1]) + _bdeg.Matrix[7]*_ge.Pow(_abcg, _bdeg.Gamma[2])
		_gafb = _bdeg.Matrix[2]*_ge.Pow(_cebe, _bdeg.Gamma[0]) + _bdeg.Matrix[5]*_ge.Pow(_cagbe, _bdeg.Gamma[1]) + _bdeg.Matrix[8]*_ge.Pow(_abcg, _bdeg.Gamma[2])
		_cebe = 3.240479*_fbgc + -1.537150*_bgadc + -0.498535*_gafb
		_cagbe = -0.969256*_fbgc + 1.875992*_bgadc + 0.041556*_gafb
		_abcg = 0.055648*_fbgc + -0.204043*_bgadc + 1.057311*_gafb
		_cebe = _ge.Min(_ge.Max(_cebe, 0), 1.0)
		_cagbe = _ge.Min(_ge.Max(_cagbe, 0), 1.0)
		_abcg = _ge.Min(_ge.Max(_abcg, 0), 1.0)
		_dbbb[0] = uint32(_cebe * _gdcg)
		_dbbb[1] = uint32(_cagbe * _gdcg)
		_dbbb[2] = uint32(_abcg * _gdcg)
		if _ggbcc = _bebf.WriteSamples(_dbbb); _ggbcc != nil {
			return img, _ggbcc
		}
	}
	return _cega(&_fafe), nil
}

func (_bfb *PdfReader) loadAction(_ddgb _abf.PdfObject) (*PdfAction, error) {
	if _fcbb, _acgg := _abf.GetIndirect(_ddgb); _acgg {
		_gge, _gbgf := _bfb.newPdfActionFromIndirectObject(_fcbb)
		if _gbgf != nil {
			return nil, _gbgf
		}
		return _gge, nil
	} else if !_abf.IsNullObject(_ddgb) {
		return nil, _fd.New("\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074 \u0074\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, nil
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

// OutlineItem represents a PDF outline item dictionary (Table 153 - pp. 376 - 377).
type OutlineItem struct {
	Title   string         `json:"title"`
	Dest    OutlineDest    `json:"dest"`
	Entries []*OutlineItem `json:"entries,omitempty"`
}

// Write writes the Appender output to io.Writer.
// It can only be called once and further invocations will result in an error.
func (_bde *PdfAppender) Write(w _gc.Writer) error {
	if _bde._ccaf {
		return _fd.New("\u0061\u0070\u0070\u0065\u006e\u0064\u0065\u0072\u0020\u0077\u0072\u0069\u0074e\u0020\u0063\u0061\u006e\u0020\u006fn\u006c\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0076\u006f\u006b\u0065\u0064 \u006f\u006e\u0063\u0065")
	}
	_dfga := NewPdfWriter()
	_bcge, _cece := _abf.GetDict(_dfga._cgeed)
	if !_cece {
		return _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_ccdc, _cece := _bcge.Get("\u004b\u0069\u0064\u0073").(*_abf.PdfObjectArray)
	if !_cece {
		return _fd.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_ecb, _cece := _bcge.Get("\u0043\u006f\u0075n\u0074").(*_abf.PdfObjectInteger)
	if !_cece {
		return _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	_cbg := _bde._agda._bebc
	_dffb := _cbg.GetTrailer()
	if _dffb == nil {
		return _fd.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_fcggd, _cece := _abf.GetIndirect(_dffb.Get("\u0052\u006f\u006f\u0074"))
	if !_cece {
		return _fd.New("c\u0061\u0074\u0061\u006c\u006f\u0067 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064")
	}
	_ebbb, _cece := _abf.GetDict(_fcggd)
	if !_cece {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _fcggd, *_dffb)
		return _fd.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	for _, _gfdf := range _ebbb.Keys() {
		if _dfga._ddffc.Get(_gfdf) == nil {
			_gacb := _ebbb.Get(_gfdf)
			_dfga._ddffc.Set(_gfdf, _gacb)
		}
	}
	if _bde._ffbb != nil {
		if _bde._ffbb._dfebf {
			if _degf := _abf.TraceToDirectObject(_bde._ffbb.ToPdfObject()); !_abf.IsNullObject(_degf) {
				_dfga._ddffc.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _degf)
				_bde.updateObjectsDeep(_degf, nil)
			} else {
				_acd.Log.Debug("\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020t\u0072\u0061\u0063e\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0020o\u0062\u006a\u0065\u0063\u0074, \u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u002e")
			}
		} else {
			_dfga._ddffc.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _bde._ffbb.ToPdfObject())
			_bde.updateObjectsDeep(_bde._ffbb.ToPdfObject(), nil)
		}
	}
	if _bde._ffbe != nil {
		_bde.updateObjectsDeep(_bde._ffbe.ToPdfObject(), nil)
		_dfga._ddffc.Set("\u0044\u0053\u0053", _bde._ffbe.GetContainingPdfObject())
	}
	if _bde._edcbe != nil {
		_dfga._ddffc.Set("\u0050\u0065\u0072m\u0073", _bde._edcbe.ToPdfObject())
		_bde.updateObjectsDeep(_bde._edcbe.ToPdfObject(), nil)
	}
	if _dfga._ecfa.Major < 2 {
		_dfga.AddExtension("\u0045\u0053\u0049\u0043", "\u0031\u002e\u0037", 5)
		_dfga.AddExtension("\u0041\u0044\u0042\u0045", "\u0031\u002e\u0037", 8)
	}
	if _faeb, _daca := _abf.GetDict(_dffb.Get("\u0049\u006e\u0066\u006f")); _daca {
		if _edff, _ffg := _abf.GetDict(_dfga._ddegc); _ffg {
			for _, _ccbe := range _faeb.Keys() {
				if _edff.Get(_ccbe) == nil {
					_edff.Set(_ccbe, _faeb.Get(_ccbe))
				}
			}
		}
	}
	if _bde._acff != nil {
		_dfga._ddegc = _abf.MakeIndirectObject(_bde._acff.ToPdfObject())
	}
	_bde.addNewObject(_dfga._ddegc)
	_bde.addNewObject(_dfga._cfdde)
	_ceed := false
	if len(_bde._agda.PageList) != len(_bde._cggfa) {
		_ceed = true
	} else {
		for _dfea := range _bde._agda.PageList {
			switch {
			case _bde._cggfa[_dfea] == _bde._agda.PageList[_dfea]:
			case _bde._cggfa[_dfea] == _bde.Reader.PageList[_dfea]:
			default:
				_ceed = true
			}
			if _ceed {
				break
			}
		}
	}
	if _ceed {
		_bde.updateObjectsDeep(_dfga._cgeed, nil)
	} else {
		_bde._cdbbg[_dfga._cgeed] = struct{}{}
	}
	_dfga._cgeed.ObjectNumber = _bde.Reader._bfdff.ObjectNumber
	_bde._bge[_dfga._cgeed] = _bde.Reader._bfdff.ObjectNumber
	_geeaa := []_abf.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	for _, _bcgb := range _bde._cggfa {
		_fea := _bcgb.ToPdfObject()
		*_ecb = *_ecb + 1
		if _fdfe, _gbeg := _fea.(*_abf.PdfIndirectObject); _gbeg && _fdfe.GetParser() == _bde._agda._bebc {
			_ccdc.Append(&_fdfe.PdfObjectReference)
			continue
		}
		if _faeaa, _dege := _abf.GetDict(_fea); _dege {
			_gdgc, _aafbd := _faeaa.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_abf.PdfIndirectObject)
			for _aafbd {
				_acd.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _gdgc)
				_aecc, _fdfc := _gdgc.PdfObject.(*_abf.PdfObjectDictionary)
				if !_fdfc {
					return _fd.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
				}
				for _, _fed := range _geeaa {
					_acd.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _fed)
					if _adbb := _faeaa.Get(_fed); _adbb != nil {
						_acd.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
						if len(_bcgb._efca.Keys()) > 0 && !_ceed {
							_abceb := _bcgb._efca
							if _dgff := _abceb.Get(_fed); _dgff != nil {
								if _adbb != _dgff {
									_acd.Log.Trace("\u0049\u006e\u0068\u0065\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u006f\u0072\u0069\u0067i\u006ea\u006c\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073\u002c\u0020\u0025\u0054", _fed, _dgff)
									_faeaa.Set(_fed, _dgff)
								}
							}
						}
						continue
					}
					if _gfag := _aecc.Get(_fed); _gfag != nil {
						_acd.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _fed)
						_faeaa.Set(_fed, _gfag)
					}
				}
				_gdgc, _aafbd = _aecc.Get("\u0050\u0061\u0072\u0065\u006e\u0074").(*_abf.PdfIndirectObject)
				_acd.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _aecc.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
			}
			if _ceed {
				_faeaa.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _dfga._cgeed)
			}
		}
		_bde.updateObjectsDeep(_fea, nil)
		_ccdc.Append(_fea)
	}
	if _, _cbef := _bde._eeded.Seek(0, _gc.SeekStart); _cbef != nil {
		return _cbef
	}
	_gcaa := make(map[SignatureHandler]_gc.Writer)
	_fgdc := _abf.MakeArray()
	for _, _dgegb := range _bde._ffcf {
		if _ecd, _bedb := _abf.GetIndirect(_dgegb); _bedb {
			if _egdc, _fbdf := _ecd.PdfObject.(*pdfSignDictionary); _fbdf {
				_bcefe := *_egdc._fafgf
				var _cbed error
				_gcaa[_bcefe], _cbed = _bcefe.NewDigest(_egdc._dcbed)
				if _cbed != nil {
					return _cbed
				}
				_fgdc.Append(_abf.MakeInteger(0xfffff), _abf.MakeInteger(0xfffff))
			}
		}
	}
	if _fgdc.Len() > 0 {
		_fgdc.Append(_abf.MakeInteger(0xfffff), _abf.MakeInteger(0xfffff))
	}
	for _, _cdaf := range _bde._ffcf {
		if _efda, _gccc := _abf.GetIndirect(_cdaf); _gccc {
			if _aeed, _ggg := _efda.PdfObject.(*pdfSignDictionary); _ggg {
				_aeed.Set("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _fgdc)
			}
		}
	}
	_cddb := len(_gcaa) > 0
	var _aaee _gc.Reader = _bde._eeded
	if _cddb {
		_agg := make([]_gc.Writer, 0, len(_gcaa))
		for _, _bafcg := range _gcaa {
			_agg = append(_agg, _bafcg)
		}
		_aaee = _gc.TeeReader(_bde._eeded, _gc.MultiWriter(_agg...))
	}
	_ffda, _cdfgc := _gc.Copy(w, _aaee)
	if _cdfgc != nil {
		return _cdfgc
	}
	if len(_bde._ffcf) == 0 {
		return nil
	}
	_dfga._cgded = _ffda
	_dfga.ObjNumOffset = _bde._ffc
	_dfga._aegbd = true
	_dfga._cagaf = _bde._abce
	_dfga._ffgf = _bde._dac
	_dfga._cfecga = _bde._cfga
	_dfga._ecfa = _bde._agda.PdfVersion()
	_dfga._deff = _bde._bge
	_dfga._ddbgd = _bde._bdcd.GetCrypter()
	_dfga._dcdbb = _bde._bdcd.GetEncryptObj()
	_dgcfd := _bde._bdcd.GetXrefType()
	if _dgcfd != nil {
		_cfdc := *_dgcfd == _abf.XrefTypeObjectStream
		_dfga._adceg = &_cfdc
	}
	_dfga._fdgae = map[_abf.PdfObject]struct{}{}
	_dfga._edcgc = []_abf.PdfObject{}
	for _, _becgd := range _bde._ffcf {
		if _, _ddfb := _bde._cdbbg[_becgd]; _ddfb {
			continue
		}
		_dfga.addObject(_becgd)
	}
	_eaca := w
	if _cddb {
		_eaca = _dd.NewBuffer(nil)
	}
	if _bde._fcfb != "" && _dfga._ddbgd == nil {
		_dfga.Encrypt([]byte(_bde._fcfb), []byte(_bde._fcfb), _bde._bbag)
	}
	if _edcg := _dffb.Get("\u0049\u0044"); _edcg != nil {
		if _dcaa, _gdea := _abf.GetArray(_edcg); _gdea {
			_dfga._dedfdf = _dcaa
		}
	}
	if _abed := _dfga.Write(_eaca); _abed != nil {
		return _abed
	}
	if _cddb {
		_fbed := _eaca.(*_dd.Buffer).Bytes()
		_cfdg := _abf.MakeArray()
		var _gggb []*pdfSignDictionary
		var _cafe int64
		for _, _dcca := range _dfga._edcgc {
			if _caeef, _dabfb := _abf.GetIndirect(_dcca); _dabfb {
				if _aefb, _fagg := _caeef.PdfObject.(*pdfSignDictionary); _fagg {
					_gggb = append(_gggb, _aefb)
					_gfggc := _aefb._eefbf + int64(_aefb._dgfdf)
					_cfdg.Append(_abf.MakeInteger(_cafe), _abf.MakeInteger(_gfggc-_cafe))
					_cafe = _aefb._eefbf + int64(_aefb._afgef)
				}
			}
		}
		_cfdg.Append(_abf.MakeInteger(_cafe), _abf.MakeInteger(_ffda+int64(len(_fbed))-_cafe))
		_aabf := []byte(_cfdg.WriteString())
		for _, _caba := range _gggb {
			_agbef := int(_caba._eefbf - _ffda)
			for _fcbg := _caba._edcbf; _fcbg < _caba._bcbcg; _fcbg++ {
				_fbed[_agbef+_fcbg] = ' '
			}
			_dbcd := _fbed[_agbef+_caba._edcbf : _agbef+_caba._bcbcg]
			copy(_dbcd, _aabf)
		}
		var _bdg int
		for _, _fffe := range _gggb {
			_eeda := int(_fffe._eefbf - _ffda)
			_fffef := _fbed[_bdg : _eeda+_fffe._dgfdf]
			_afc := *_fffe._fafgf
			_gcaa[_afc].Write(_fffef)
			_bdg = _eeda + _fffe._afgef
		}
		for _, _cgaef := range _gggb {
			_gfca := _fbed[_bdg:]
			_feae := *_cgaef._fafgf
			_gcaa[_feae].Write(_gfca)
		}
		for _, _dffa := range _gggb {
			_febcf := int(_dffa._eefbf - _ffda)
			_beeee := *_dffa._fafgf
			_cabc := _gcaa[_beeee]
			if _cebag := _beeee.Sign(_dffa._dcbed, _cabc); _cebag != nil {
				return _cebag
			}
			_dffa._dcbed.ByteRange = _cfdg
			_acbdf := []byte(_dffa._dcbed.Contents.WriteString())
			for _aaadf := _dffa._edcbf; _aaadf < _dffa._bcbcg; _aaadf++ {
				_fbed[_febcf+_aaadf] = ' '
			}
			for _ageb := _dffa._dgfdf; _ageb < _dffa._afgef; _ageb++ {
				_fbed[_febcf+_ageb] = ' '
			}
			_begc := _fbed[_febcf+_dffa._edcbf : _febcf+_dffa._bcbcg]
			copy(_begc, _aabf)
			_begc = _fbed[_febcf+_dffa._dgfdf : _febcf+_dffa._afgef]
			copy(_begc, _acbdf)
		}
		_efdad := _dd.NewBuffer(_fbed)
		_, _cdfgc = _gc.Copy(w, _efdad)
		if _cdfgc != nil {
			return _cdfgc
		}
	}
	_bde._ccaf = true
	return nil
}
func _eggec() string { return _acd.Version }

// NewPdfColorspaceDeviceRGB returns a new RGB colorspace object.
func NewPdfColorspaceDeviceRGB() *PdfColorspaceDeviceRGB { return &PdfColorspaceDeviceRGB{} }

// NewPdfAnnotation returns an initialized generic PDF annotation model.
func NewPdfAnnotation() *PdfAnnotation {
	_dcb := &PdfAnnotation{}
	_dcb._dbc = _abf.MakeIndirectObject(_abf.MakeDict())
	return _dcb
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_gdgee *PdfPageResourcesColorspaces) ToPdfObject() _abf.PdfObject {
	_aacfc := _abf.MakeDict()
	for _, _acbffc := range _gdgee.Names {
		_aacfc.Set(_abf.PdfObjectName(_acbffc), _gdgee.Colorspaces[_acbffc].ToPdfObject())
	}
	if _gdgee._cebc != nil {
		_gdgee._cebc.PdfObject = _aacfc
		return _gdgee._cebc
	}
	return _aacfc
}

// SetAlpha sets the alpha layer for the image.
func (_abbbc *Image) SetAlpha(alpha []byte) { _abbbc._gedg = alpha }

// NewPdfColorspaceFromPdfObject loads a PdfColorspace from a PdfObject.  Returns an error if there is
// a failure in loading.
func NewPdfColorspaceFromPdfObject(obj _abf.PdfObject) (PdfColorspace, error) {
	if obj == nil {
		return nil, nil
	}
	var _fcba *_abf.PdfIndirectObject
	var _edef *_abf.PdfObjectName
	var _efbd *_abf.PdfObjectArray
	if _ggeb, _cecb := obj.(*_abf.PdfIndirectObject); _cecb {
		_fcba = _ggeb
	}
	obj = _abf.TraceToDirectObject(obj)
	switch _fbafg := obj.(type) {
	case *_abf.PdfObjectArray:
		_efbd = _fbafg
	case *_abf.PdfObjectName:
		_edef = _fbafg
	}
	if _edef != nil {
		switch *_edef {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			return NewPdfColorspaceDeviceGray(), nil
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			return NewPdfColorspaceDeviceRGB(), nil
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return NewPdfColorspaceDeviceCMYK(), nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return NewPdfColorspaceSpecialPattern(), nil
		default:
			_acd.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0025\u0073", *_edef)
			return nil, _bgaaa
		}
	}
	if _efbd != nil && _efbd.Len() > 0 {
		var _ccfa _abf.PdfObject = _fcba
		if _fcba == nil {
			_ccfa = _efbd
		}
		if _fbbd, _fcbdb := _abf.GetName(_efbd.Get(0)); _fcbdb {
			switch _fbbd.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				if _efbd.Len() == 1 {
					return NewPdfColorspaceDeviceGray(), nil
				}
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				if _efbd.Len() == 1 {
					return NewPdfColorspaceDeviceRGB(), nil
				}
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _efbd.Len() == 1 {
					return NewPdfColorspaceDeviceCMYK(), nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079":
				return _gcfd(_ccfa)
			case "\u0043\u0061\u006c\u0052\u0047\u0042":
				return _bfbg(_ccfa)
			case "\u004c\u0061\u0062":
				return _agcb(_ccfa)
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064":
				return _becc(_ccfa)
			case "\u0050a\u0074\u0074\u0065\u0072\u006e":
				return _fcce(_ccfa)
			case "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return _acffa(_ccfa)
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e":
				return _cecag(_ccfa)
			case "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return _egeeb(_ccfa)
			default:
				_acd.Log.Debug("A\u0072\u0072\u0061\u0079\u0020\u0077i\u0074\u0068\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0061m\u0065:\u0020\u0025\u0073", *_fbbd)
			}
		}
	}
	_acd.Log.Debug("\u0050\u0044\u0046\u0020\u0046i\u006c\u0065\u0020\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", obj.String())
	return nil, ErrTypeCheck
}

// PdfActionSound represents a sound action.
type PdfActionSound struct {
	*PdfAction
	Sound       _abf.PdfObject
	Volume      _abf.PdfObject
	Synchronous _abf.PdfObject
	Repeat      _abf.PdfObject
	Mix         _abf.PdfObject
}

// SetContext set the sub annotation (context).
func (_dgeec *PdfShading) SetContext(ctx PdfModel) { _dgeec._eabd = ctx }

// ToPdfObject implements interface PdfModel.
func (_gdeg *PdfAnnotationRedact) ToPdfObject() _abf.PdfObject {
	_gdeg.PdfAnnotation.ToPdfObject()
	_ebbe := _gdeg._dbc
	_ddfd := _ebbe.PdfObject.(*_abf.PdfObjectDictionary)
	_gdeg.PdfAnnotationMarkup.appendToPdfDictionary(_ddfd)
	_ddfd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0052\u0065\u0064\u0061\u0063\u0074"))
	_ddfd.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _gdeg.QuadPoints)
	_ddfd.SetIfNotNil("\u0049\u0043", _gdeg.IC)
	_ddfd.SetIfNotNil("\u0052\u004f", _gdeg.RO)
	_ddfd.SetIfNotNil("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074", _gdeg.OverlayText)
	_ddfd.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _gdeg.Repeat)
	_ddfd.SetIfNotNil("\u0044\u0041", _gdeg.DA)
	_ddfd.SetIfNotNil("\u0051", _gdeg.Q)
	return _ebbe
}

// ToPdfObject implements interface PdfModel.
func (_eeg *PdfActionLaunch) ToPdfObject() _abf.PdfObject {
	_eeg.PdfAction.ToPdfObject()
	_gec := _eeg._egg
	_fdc := _gec.PdfObject.(*_abf.PdfObjectDictionary)
	_fdc.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeLaunch)))
	if _eeg.F != nil {
		_fdc.Set("\u0046", _eeg.F.ToPdfObject())
	}
	_fdc.SetIfNotNil("\u0057\u0069\u006e", _eeg.Win)
	_fdc.SetIfNotNil("\u004d\u0061\u0063", _eeg.Mac)
	_fdc.SetIfNotNil("\u0055\u006e\u0069\u0078", _eeg.Unix)
	_fdc.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _eeg.NewWindow)
	return _gec
}

// GetNumComponents returns the number of color components.
func (_cbbed *PdfColorspaceICCBased) GetNumComponents() int { return _cbbed.N }

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
	ColorToRGB(_fcbd PdfColor) (PdfColor, error)

	// GetNumComponents returns the number of components in the PdfColorspace.
	GetNumComponents() int

	// ToPdfObject returns a PdfObject representation of the PdfColorspace.
	ToPdfObject() _abf.PdfObject

	// ColorFromPdfObjects returns a PdfColor in the given PdfColorspace from an array of PdfObject where each
	// PdfObject represents a numeric value.
	ColorFromPdfObjects(_efa []_abf.PdfObject) (PdfColor, error)

	// ColorFromFloats returns a new PdfColor based on input color components for a given PdfColorspace.
	ColorFromFloats(_ebbc []float64) (PdfColor, error)

	// DecodeArray returns the Decode array for the PdfColorSpace, i.e. the range of each component.
	DecodeArray() []float64
}

func _geead() string {
	_acef := "\u0051\u0057\u0045\u0052\u0054\u0059\u0055\u0049\u004f\u0050\u0041S\u0044\u0046\u0047\u0048\u004a\u004b\u004c\u005a\u0058\u0043V\u0042\u004e\u004d"
	var _efgb _dd.Buffer
	for _fbdbg := 0; _fbdbg < 6; _fbdbg++ {
		_efgb.WriteRune(rune(_acef[_aaf.Intn(len(_acef))]))
	}
	return _efgb.String()
}

func _caece(_fcag _abf.PdfObject, _dggb bool) (*PdfFont, error) {
	_gacbd, _dcdb, _fcced := _addf(_fcag)
	if _gacbd != nil {
		_gedcb(_gacbd)
	}
	if _fcced != nil {
		if _fcced == ErrType1CFontNotSupported {
			_cacd, _abcbe := _fggeg(_gacbd, _dcdb, nil)
			if _abcbe != nil {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _dcdb, _abcbe)
				return nil, _fcced
			}
			return &PdfFont{_gedca: _cacd}, _fcced
		}
		return nil, _fcced
	}
	_dead := &PdfFont{}
	switch _dcdb._aacbc {
	case "\u0054\u0079\u0070e\u0030":
		if !_dggb {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u004c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u00650\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _dcdb)
			return nil, _fd.New("\u0063\u0079\u0063\u006cic\u0061\u006c\u0020\u0074\u0079\u0070\u0065\u0030\u0020\u006c\u006f\u0061\u0064\u0069n\u0067")
		}
		_deaa, _gdegb := _gdcdc(_gacbd, _dcdb)
		if _gdegb != nil {
			_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0057\u0068\u0069l\u0065\u0020\u006c\u006f\u0061\u0064\u0069ng\u0020\u0054\u0079\u0070e\u0030\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066on\u0074\u003d%\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dcdb, _gdegb)
			return nil, _gdegb
		}
		_dead._gedca = _deaa
	case "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		var _ecgg *pdfFontSimple
		_aggb, _bdfe := _gbe.NewStdFontByName(_gbe.StdFontName(_dcdb._ecggf))
		if _bdfe {
			_faggc := _bcee(_aggb)
			_dead._gedca = &_faggc
			_eaacb := _abf.TraceToDirectObject(_faggc.ToPdfObject())
			_dcfg, _dcgg, _bccb := _addf(_eaacb)
			if _bccb != nil {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _dcdb, _faggc)
				return nil, _bccb
			}
			for _, _afefg := range _gacbd.Keys() {
				_dcfg.Set(_afefg, _gacbd.Get(_afefg))
			}
			_ecgg, _bccb = _fggeg(_dcfg, _dcgg, _faggc._edabc)
			if _bccb != nil {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0064\u0020\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u00314\u000a\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u000a\u0009\u0073\u0074d\u003d\u0025\u002b\u0076", _dcdb, _faggc)
				return nil, _bccb
			}
			_ecgg._aadgb = _faggc._aadgb
			_ecgg._aecd = _faggc._aecd
			if _ecgg._abeb == nil {
				_ecgg._abeb = _faggc._abeb
			}
		} else {
			_ecgg, _fcced = _fggeg(_gacbd, _dcdb, nil)
			if _fcced != nil {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0057h\u0069\u006c\u0065 l\u006f\u0061\u0064\u0069\u006e\u0067 \u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072=\u0025\u0076", _dcdb, _fcced)
				return nil, _fcced
			}
		}
		_fcced = _ecgg.addEncoding()
		if _fcced != nil {
			return nil, _fcced
		}
		if _bdfe {
			_ecgg.updateStandard14Font()
		}
		if _bdfe && _ecgg._ebada == nil && _ecgg._edabc == nil {
			_acd.Log.Error("\u0073\u0069\u006d\u0070\u006c\u0065\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _ecgg)
			_acd.Log.Error("\u0066n\u0074\u003d\u0025\u002b\u0076", _aggb)
		}
		if len(_ecgg._aadgb) == 0 {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u006f\u0020\u0077\u0069d\u0074h\u0073.\u0020\u0066\u006f\u006e\u0074\u003d\u0025s", _ecgg)
		}
		_dead._gedca = _ecgg
	case "\u0054\u0079\u0070e\u0033":
		_dbad, _fccc := _bddec(_gacbd, _dcdb)
		if _fccc != nil {
			_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020W\u0068\u0069\u006c\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0074y\u0070\u0065\u0033\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _fccc)
			return nil, _fccc
		}
		_dead._gedca = _dbad
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030":
		_geccf, _ceedc := _edde(_gacbd, _dcdb)
		if _ceedc != nil {
			_acd.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0057\u0068i\u006c\u0065\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069d \u0066\u006f\u006et\u0020\u0074y\u0070\u0065\u0030\u0020\u0066\u006fn\u0074\u003a \u0025\u0076", _ceedc)
			return nil, _ceedc
		}
		_dead._gedca = _geccf
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_cabe, _abfed := _fccda(_gacbd, _dcdb)
		if _abfed != nil {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0057\u0068\u0069l\u0065\u0020\u006co\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0069\u0064\u0020f\u006f\u006e\u0074\u0020\u0074yp\u0065\u0032\u0020\u0066\u006f\u006e\u0074\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dcdb, _abfed)
			return nil, _abfed
		}
		_dead._gedca = _cabe
	default:
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020U\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020f\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0066\u006fn\u0074\u003d\u0025\u0073", _dcdb)
		return nil, _e.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0066\u006f\u006e\u0074\u0020\u0074y\u0070\u0065\u003a\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _dcdb)
	}
	return _dead, nil
}

// L returns the value of the L component of the color.
func (_gfbbc *PdfColorLab) L() float64 { return _gfbbc[0] }

// ToInteger convert to an integer format.
func (_gabaa *PdfColorDeviceCMYK) ToInteger(bits int) [4]uint32 {
	_abgd := _ge.Pow(2, float64(bits)) - 1
	return [4]uint32{uint32(_abgd * _gabaa.C()), uint32(_abgd * _gabaa.M()), uint32(_abgd * _gabaa.Y()), uint32(_abgd * _gabaa.K())}
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

// NewStandard14FontMustCompile returns the standard 14 font named `basefont` as a *PdfFont.
// If `basefont` is one of the 14 Standard14Font values defined above then NewStandard14FontMustCompile
// is guaranteed to succeed.
func NewStandard14FontMustCompile(basefont StdFontName) *PdfFont {
	_bddcf, _ggcc := NewStandard14Font(basefont)
	if _ggcc != nil {
		panic(_e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0031\u0034\u0046\u006f\u006e\u0074 \u0025\u0023\u0071", basefont))
	}
	return _bddcf
}

// ToPdfObject implements interface PdfModel.
func (_aed *PdfActionGoToE) ToPdfObject() _abf.PdfObject {
	_aed.PdfAction.ToPdfObject()
	_agb := _aed._egg
	_ebd := _agb.PdfObject.(*_abf.PdfObjectDictionary)
	_ebd.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeGoToE)))
	if _aed.F != nil {
		_ebd.Set("\u0046", _aed.F.ToPdfObject())
	}
	_ebd.SetIfNotNil("\u0044", _aed.D)
	_ebd.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _aed.NewWindow)
	_ebd.SetIfNotNil("\u0054", _aed.T)
	return _agb
}
func _bcce(_bdfbb *fontCommon) *pdfCIDFontType0 { return &pdfCIDFontType0{fontCommon: *_bdfbb} }

// FullName returns the full name of the field as in rootname.parentname.partialname.
func (_gafe *PdfField) FullName() (string, error) {
	var _fcbe _dd.Buffer
	_agbec := []string{}
	if _gafe.T != nil {
		_agbec = append(_agbec, _gafe.T.Decoded())
	}
	_dfede := map[*PdfField]bool{}
	_dfede[_gafe] = true
	_afda := _gafe.Parent
	for _afda != nil {
		if _, _fgccd := _dfede[_afda]; _fgccd {
			return _fcbe.String(), _fd.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		if _afda.T == nil {
			return _fcbe.String(), _fd.New("\u0066\u0069el\u0064\u0020\u0070a\u0072\u0074\u0069\u0061l n\u0061me\u0020\u0028\u0054\u0029\u0020\u006e\u006ft \u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064")
		}
		_agbec = append(_agbec, _afda.T.Decoded())
		_dfede[_afda] = true
		_afda = _afda.Parent
	}
	for _eeaf := len(_agbec) - 1; _eeaf >= 0; _eeaf-- {
		_fcbe.WriteString(_agbec[_eeaf])
		if _eeaf > 0 {
			_fcbe.WriteString("\u002e")
		}
	}
	return _fcbe.String(), nil
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

// NewPdfAnnotationProjection returns a new projection annotation.
func NewPdfAnnotationProjection() *PdfAnnotationProjection {
	_fffd := NewPdfAnnotation()
	_ggcfg := &PdfAnnotationProjection{}
	_ggcfg.PdfAnnotation = _fffd
	_ggcfg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fffd.SetContext(_ggcfg)
	return _ggcfg
}

// Insert adds a top level outline item in the outline,
// at the specified index.
func (_ddea *Outline) Insert(index uint, item *OutlineItem) {
	_dfeg := uint(len(_ddea.Entries))
	if index > _dfeg {
		index = _dfeg
	}
	_ddea.Entries = append(_ddea.Entries[:index], append([]*OutlineItem{item}, _ddea.Entries[index:]...)...)
}

// ToPdfObject implements interface PdfModel.
func (_ca *PdfActionGoTo) ToPdfObject() _abf.PdfObject {
	_ca.PdfAction.ToPdfObject()
	_bfe := _ca._egg
	_aab := _bfe.PdfObject.(*_abf.PdfObjectDictionary)
	_aab.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeGoTo)))
	_aab.SetIfNotNil("\u0044", _ca.D)
	return _bfe
}

func (_dggbcb *PdfWriter) setCatalogVersion() {
	_dggbcb._ddffc.Set("\u0056e\u0072\u0073\u0069\u006f\u006e", _abf.MakeName(_e.Sprintf("\u0025\u0064\u002e%\u0064", _dggbcb._ecfa.Major, _dggbcb._ecfa.Minor)))
}

// GetNumComponents returns the number of input color components, i.e. that are input to the tint transform.
func (_dgeea *PdfColorspaceDeviceN) GetNumComponents() int { return _dgeea.ColorantNames.Len() }

type pdfFontSimple struct {
	fontCommon
	_ddddaf *_abf.PdfIndirectObject
	_aadgb  map[_cbb.CharCode]float64
	_ebada  _cbb.TextEncoder
	_edabc  _cbb.TextEncoder
	_abeb   *PdfFontDescriptor

	// Encoding is subject to limitations that are described in 9.6.6, "Character Encoding".
	// BaseFont is derived differently.
	FirstChar _abf.PdfObject
	LastChar  _abf.PdfObject
	Widths    _abf.PdfObject
	Encoding  _abf.PdfObject
	_aecd     *_gbe.RuneCharSafeMap
}

// SetCatalogMetadata sets the catalog metadata (XMP) stream object.
func (_bdeadc *PdfWriter) SetCatalogMetadata(meta _abf.PdfObject) error {
	if meta == nil {
		_bdeadc._ddffc.Remove("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
		return nil
	}
	_fedbg, _dfaee := _abf.GetStream(meta)
	if !_dfaee {
		return _fd.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006d\u0065\u0074\u0061\u0064a\u0074\u0061\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0073t\u0072\u0065\u0061\u006d")
	}
	_bdeadc.addObject(_fedbg)
	_bdeadc._ddffc.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _fedbg)
	return nil
}

func (_aeffd *DSS) generateHashMap(_dddd []*_abf.PdfObjectStream) (map[string]*_abf.PdfObjectStream, error) {
	_bbbg := map[string]*_abf.PdfObjectStream{}
	for _, _ebbed := range _dddd {
		_fccfa, _cdef := _abf.DecodeStream(_ebbed)
		if _cdef != nil {
			return nil, _cdef
		}
		_gbafc, _cdef := _fdbbe(_fccfa)
		if _cdef != nil {
			return nil, _cdef
		}
		_bbbg[string(_gbafc)] = _ebbed
	}
	return _bbbg, nil
}

// IsColored specifies if the pattern is colored.
func (_fgdgg *PdfTilingPattern) IsColored() bool {
	if _fgdgg.PaintType != nil && *_fgdgg.PaintType == 1 {
		return true
	}
	return false
}

// PdfActionImportData represents a importData action.
type PdfActionImportData struct {
	*PdfAction
	F *PdfFilespec
}

func (_aebefc *PdfWriter) writeOutputIntents() error {
	if len(_aebefc._dgfea) == 0 {
		return nil
	}
	_ebgf := make([]_abf.PdfObject, len(_aebefc._dgfea))
	for _bdgee, _caeb := range _aebefc._dgfea {
		_ecfef := _caeb.ToPdfObject()
		_ebgf[_bdgee] = _abf.MakeIndirectObject(_ecfef)
	}
	_befbc := _abf.MakeIndirectObject(_abf.MakeArray(_ebgf...))
	_aebefc._ddffc.Set("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073", _befbc)
	if _aabeb := _aebefc.addObjects(_befbc); _aabeb != nil {
		return _aabeb
	}
	return nil
}

func (_ceb *PdfReader) newPdfAnnotationPopupFromDict(_daf *_abf.PdfObjectDictionary) (*PdfAnnotationPopup, error) {
	_ddbb := PdfAnnotationPopup{}
	_ddbb.Parent = _daf.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	_ddbb.Open = _daf.Get("\u004f\u0070\u0065\u006e")
	return &_ddbb, nil
}

// ToInteger convert to an integer format.
func (_efbc *PdfColorCalGray) ToInteger(bits int) uint32 {
	_ecec := _ge.Pow(2, float64(bits)) - 1
	return uint32(_ecec * _efbc.Val())
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_eegff *PdfColorspaceCalGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cfec := vals[0]
	if _cfec < 0.0 || _cfec > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _cfec)
		return nil, ErrColorOutOfRange
	}
	_fgce := NewPdfColorCalGray(_cfec)
	return _fgce, nil
}

// NewPdfShadingPatternType3 creates an empty shading pattern type 3 object.
func NewPdfShadingPatternType3() *PdfShadingPatternType3 {
	_edgd := &PdfShadingPatternType3{}
	_edgd.Matrix = _abf.MakeArrayFromIntegers([]int{1, 0, 0, 1, 0, 0})
	_edgd.PdfPattern = &PdfPattern{}
	_edgd.PdfPattern.PatternType = int64(*_abf.MakeInteger(2))
	_edgd.PdfPattern._bgafe = _edgd
	_edgd.PdfPattern._bcfca = _abf.MakeIndirectObject(_abf.MakeDict())
	return _edgd
}

// Set sets the colorspace corresponding to key. Add to Names if not set.
func (_gaegbd *PdfPageResourcesColorspaces) Set(key _abf.PdfObjectName, val PdfColorspace) {
	if _, _afbg := _gaegbd.Colorspaces[string(key)]; !_afbg {
		_gaegbd.Names = append(_gaegbd.Names, string(key))
	}
	_gaegbd.Colorspaces[string(key)] = val
}

// ImageToRGB converts an image in CMYK32 colorspace to an RGB image.
func (_abaaa *PdfColorspaceDeviceCMYK) ImageToRGB(img Image) (Image, error) {
	_acd.Log.Trace("\u0043\u004d\u0059\u004b\u0033\u0032\u0020\u002d\u003e\u0020\u0052\u0047\u0042")
	_acd.Log.Trace("I\u006d\u0061\u0067\u0065\u0020\u0042P\u0043\u003a\u0020\u0025\u0064\u002c \u0043\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020%\u0064", img.BitsPerComponent, img.ColorComponents)
	_acd.Log.Trace("\u004c\u0065\u006e \u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0064", len(img.Data))
	_acd.Log.Trace("H\u0065\u0069\u0067\u0068t:\u0020%\u0064\u002c\u0020\u0057\u0069d\u0074\u0068\u003a\u0020\u0025\u0064", img.Height, img.Width)
	_dgfg, _ecbc := _gca.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._gedg, img._ceeag)
	if _ecbc != nil {
		return Image{}, _ecbc
	}
	_ceeg, _ecbc := _gca.NRGBAConverter.Convert(_dgfg)
	if _ecbc != nil {
		return Image{}, _ecbc
	}
	return _cega(_ceeg.Base()), nil
}

func (_caf *PdfReader) newPdfAnnotationMovieFromDict(_fgada *_abf.PdfObjectDictionary) (*PdfAnnotationMovie, error) {
	_ddaa := PdfAnnotationMovie{}
	_ddaa.T = _fgada.Get("\u0054")
	_ddaa.Movie = _fgada.Get("\u004d\u006f\u0076i\u0065")
	_ddaa.A = _fgada.Get("\u0041")
	return &_ddaa, nil
}

// PdfInfoTrapped specifies pdf trapped information.
type PdfInfoTrapped string

func (_ccegeg *PdfWriter) setHashIDs(_aebg _a.Hash) error {
	_cdce := _aebg.Sum(nil)
	if _ccegeg._aefff == "" {
		_ccegeg._aefff = _cb.EncodeToString(_cdce[:8])
	}
	_ccegeg.setDocumentIDs(_ccegeg._aefff, _cb.EncodeToString(_cdce[8:]))
	return nil
}

// ToPdfObject implements interface PdfModel.
func (_bae *PdfActionJavaScript) ToPdfObject() _abf.PdfObject {
	_bae.PdfAction.ToPdfObject()
	_eddc := _bae._egg
	_abe := _eddc.PdfObject.(*_abf.PdfObjectDictionary)
	_abe.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeJavaScript)))
	_abe.SetIfNotNil("\u004a\u0053", _bae.JS)
	return _eddc
}

func (_aca *PdfReader) newPdfActionGotoEFromDict(_ef *_abf.PdfObjectDictionary) (*PdfActionGoToE, error) {
	_cgb, _efg := _dgf(_ef.Get("\u0046"))
	if _efg != nil {
		return nil, _efg
	}
	return &PdfActionGoToE{D: _ef.Get("\u0044"), NewWindow: _ef.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), T: _ef.Get("\u0054"), F: _cgb}, nil
}

// NewPdfActionGoTo returns a new "go to" action.
func NewPdfActionGoTo() *PdfActionGoTo {
	_bgb := NewPdfAction()
	_bce := &PdfActionGoTo{}
	_bce.PdfAction = _bgb
	_bgb.SetContext(_bce)
	return _bce
}

// IsValid checks if the given pdf output intent type is valid.
func (_fbced PdfOutputIntentType) IsValid() bool {
	return _fbced >= PdfOutputIntentTypeA1 && _fbced <= PdfOutputIntentTypeX
}

// GetContext returns the PdfField context which is the more specific field data type, e.g. PdfFieldButton
// for a button field.
func (_edeee *PdfField) GetContext() PdfModel { return _edeee._ffea }

// UpdatePage updates the `page` in the new revision if it has changed.
func (_dcff *PdfAppender) UpdatePage(page *PdfPage) { _dcff.updateObjectsDeep(page.ToPdfObject(), nil) }

// ToInteger convert to an integer format.
func (_gbdef *PdfColorLab) ToInteger(bits int) [3]uint32 {
	_ccafd := _ge.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_ccafd * _gbdef.L()), uint32(_ccafd * _gbdef.A()), uint32(_ccafd * _gbdef.B())}
}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_dfed *PdfColorspaceCalRGB) ToPdfObject() _abf.PdfObject {
	_gabce := &_abf.PdfObjectArray{}
	_gabce.Append(_abf.MakeName("\u0043\u0061\u006c\u0052\u0047\u0042"))
	_dfbc := _abf.MakeDict()
	if _dfed.WhitePoint != nil {
		_aafbg := _abf.MakeArray(_abf.MakeFloat(_dfed.WhitePoint[0]), _abf.MakeFloat(_dfed.WhitePoint[1]), _abf.MakeFloat(_dfed.WhitePoint[2]))
		_dfbc.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _aafbg)
	} else {
		_acd.Log.Error("\u0043\u0061l\u0052\u0047\u0042\u003a \u004d\u0069s\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _dfed.BlackPoint != nil {
		_cfba := _abf.MakeArray(_abf.MakeFloat(_dfed.BlackPoint[0]), _abf.MakeFloat(_dfed.BlackPoint[1]), _abf.MakeFloat(_dfed.BlackPoint[2]))
		_dfbc.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _cfba)
	}
	if _dfed.Gamma != nil {
		_abad := _abf.MakeArray(_abf.MakeFloat(_dfed.Gamma[0]), _abf.MakeFloat(_dfed.Gamma[1]), _abf.MakeFloat(_dfed.Gamma[2]))
		_dfbc.Set("\u0047\u0061\u006dm\u0061", _abad)
	}
	if _dfed.Matrix != nil {
		_fgbfb := _abf.MakeArray(_abf.MakeFloat(_dfed.Matrix[0]), _abf.MakeFloat(_dfed.Matrix[1]), _abf.MakeFloat(_dfed.Matrix[2]), _abf.MakeFloat(_dfed.Matrix[3]), _abf.MakeFloat(_dfed.Matrix[4]), _abf.MakeFloat(_dfed.Matrix[5]), _abf.MakeFloat(_dfed.Matrix[6]), _abf.MakeFloat(_dfed.Matrix[7]), _abf.MakeFloat(_dfed.Matrix[8]))
		_dfbc.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _fgbfb)
	}
	_gabce.Append(_dfbc)
	if _dfed._bdfg != nil {
		_dfed._bdfg.PdfObject = _gabce
		return _dfed._bdfg
	}
	return _gabce
}

func (_gdcd *PdfReader) newPdfAnnotationFreeTextFromDict(_cag *_abf.PdfObjectDictionary) (*PdfAnnotationFreeText, error) {
	_cddfd := PdfAnnotationFreeText{}
	_aead, _adce := _gdcd.newPdfAnnotationMarkupFromDict(_cag)
	if _adce != nil {
		return nil, _adce
	}
	_cddfd.PdfAnnotationMarkup = _aead
	_cddfd.DA = _cag.Get("\u0044\u0041")
	_cddfd.Q = _cag.Get("\u0051")
	_cddfd.RC = _cag.Get("\u0052\u0043")
	_cddfd.DS = _cag.Get("\u0044\u0053")
	_cddfd.CL = _cag.Get("\u0043\u004c")
	_cddfd.IT = _cag.Get("\u0049\u0054")
	_cddfd.BE = _cag.Get("\u0042\u0045")
	_cddfd.RD = _cag.Get("\u0052\u0044")
	_cddfd.BS = _cag.Get("\u0042\u0053")
	_cddfd.LE = _cag.Get("\u004c\u0045")
	return &_cddfd, nil
}

// PdfAnnotationWidget represents Widget annotations.
// Note: Widget annotations are used to display form fields.
// (Section 12.5.6.19).
type PdfAnnotationWidget struct {
	*PdfAnnotation
	H      _abf.PdfObject
	MK     _abf.PdfObject
	A      _abf.PdfObject
	AA     _abf.PdfObject
	BS     _abf.PdfObject
	Parent _abf.PdfObject
	_agdc  *PdfField
	_gbga  bool
}

// ImageToRGB convert 1-component grayscale data to 3-component RGB.
func (_ageac *PdfColorspaceDeviceGray) ImageToRGB(img Image) (Image, error) {
	if img.ColorComponents != 1 {
		return img, _fd.New("\u0074\u0068e \u0070\u0072\u006fv\u0069\u0064\u0065\u0064 im\u0061ge\u0020\u0069\u0073\u0020\u006e\u006f\u0074 g\u0072\u0061\u0079\u0020\u0073\u0063\u0061l\u0065")
	}
	_beebb, _bddc := _gca.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._gedg, img._ceeag)
	if _bddc != nil {
		return img, _bddc
	}
	_fcgf, _bddc := _gca.NRGBAConverter.Convert(_beebb)
	if _bddc != nil {
		return img, _bddc
	}
	_gbff := _cega(_fcgf.Base())
	_acd.Log.Trace("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079\u0020\u002d>\u0020\u0052\u0047\u0042")
	_acd.Log.Trace("s\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", img.Data)
	_acd.Log.Trace("\u0052G\u0042 \u0073\u0061\u006d\u0070\u006c\u0065\u0073\u003a\u0020\u0025\u0076", _gbff.Data)
	_acd.Log.Trace("\u0025\u0076\u0020\u002d\u003e\u0020\u0025\u0076", img, _gbff)
	return _gbff, nil
}

// PdfColorPattern represents a pattern color.
type PdfColorPattern struct {
	Color       PdfColor
	PatternName _abf.PdfObjectName
}

// PdfActionJavaScript represents a javaScript action.
type PdfActionJavaScript struct {
	*PdfAction
	JS _abf.PdfObject
}

// PdfAnnotationInk represents Ink annotations.
// (Section 12.5.6.13).
type PdfAnnotationInk struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	InkList _abf.PdfObject
	BS      _abf.PdfObject
}

// Val returns the value of the color.
func (_agga *PdfColorCalGray) Val() float64 { return float64(*_agga) }

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_fadcf *PdfColorspaceCalGray) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cgcd, _egbeb := _abf.GetNumbersAsFloat(objects)
	if _egbeb != nil {
		return nil, _egbeb
	}
	return _fadcf.ColorFromFloats(_cgcd)
}

func (_ggbc *PdfReader) newPdfAnnotationCircleFromDict(_gbcb *_abf.PdfObjectDictionary) (*PdfAnnotationCircle, error) {
	_adag := PdfAnnotationCircle{}
	_efe, _dcga := _ggbc.newPdfAnnotationMarkupFromDict(_gbcb)
	if _dcga != nil {
		return nil, _dcga
	}
	_adag.PdfAnnotationMarkup = _efe
	_adag.BS = _gbcb.Get("\u0042\u0053")
	_adag.IC = _gbcb.Get("\u0049\u0043")
	_adag.BE = _gbcb.Get("\u0042\u0045")
	_adag.RD = _gbcb.Get("\u0052\u0044")
	return &_adag, nil
}

// BytesToCharcodes converts the bytes in a PDF string to character codes.
func (_bgcg *PdfFont) BytesToCharcodes(data []byte) []_cbb.CharCode {
	_acd.Log.Trace("\u0042\u0079\u0074es\u0054\u006f\u0043\u0068\u0061\u0072\u0063\u006f\u0064e\u0073:\u0020d\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", data, data)
	if _gfbe, _baec := _bgcg._gedca.(*pdfFontType0); _baec && _gfbe._fcfg != nil {
		if _ddegf, _dcee := _gfbe.bytesToCharcodes(data); _dcee {
			return _ddegf
		}
	}
	var (
		_baadf = make([]_cbb.CharCode, 0, len(data)+len(data)%2)
		_bcbgb = _bgcg.baseFields()
	)
	if _bcbgb._aabfe != nil {
		if _aggab, _gfba := _bcbgb._aabfe.BytesToCharcodes(data); _gfba {
			for _, _gcgd := range _aggab {
				_baadf = append(_baadf, _cbb.CharCode(_gcgd))
			}
			return _baadf
		}
	}
	if _bcbgb.isCIDFont() {
		if len(data) == 1 {
			data = []byte{0, data[0]}
		}
		if len(data)%2 != 0 {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0064\u0061\u0074\u0061\u003d\u0025\u002b\u0076\u0020t\u006f\u0020\u0065\u0076\u0065n\u0020\u006ce\u006e\u0067\u0074\u0068", data)
			data = append(data, 0)
		}
		for _gbaadg := 0; _gbaadg < len(data); _gbaadg += 2 {
			_bcede := uint16(data[_gbaadg])<<8 | uint16(data[_gbaadg+1])
			_baadf = append(_baadf, _cbb.CharCode(_bcede))
		}
	} else {
		for _, _cdeab := range data {
			_baadf = append(_baadf, _cbb.CharCode(_cdeab))
		}
	}
	return _baadf
}

// NewPdfActionGoTo3DView returns a new "goTo3DView" action.
func NewPdfActionGoTo3DView() *PdfActionGoTo3DView {
	_bgg := NewPdfAction()
	_adg := &PdfActionGoTo3DView{}
	_adg.PdfAction = _bgg
	_bgg.SetContext(_adg)
	return _adg
}

// AddPage adds a page to the PDF file. The new page should be an indirect object.
func (_eaegf *PdfWriter) AddPage(page *PdfPage) error {
	_eadfe := page.ToPdfObject()
	_acd.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
	_acd.Log.Trace("\u0041p\u0070\u0065\u006e\u0064i\u006e\u0067\u0020\u0074\u006f \u0070a\u0067e\u0020\u006c\u0069\u0073\u0074\u0020\u0025T", _eadfe)
	_fgeg, _cafce := _abf.GetIndirect(_eadfe)
	if !_cafce {
		return _fd.New("\u0070\u0061\u0067\u0065\u0020\u0073h\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_acd.Log.Trace("\u0025\u0073", _fgeg)
	_acd.Log.Trace("\u0025\u0073", _fgeg.PdfObject)
	_dbdef, _cafce := _abf.GetDict(_fgeg.PdfObject)
	if !_cafce {
		return _fd.New("\u0070\u0061\u0067e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068o\u0075l\u0064 \u0062e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_deddf, _cafce := _abf.GetName(_dbdef.Get("\u0054\u0079\u0070\u0065"))
	if !_cafce {
		return _e.Errorf("\u0070\u0061\u0067\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054y\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u006e\u0061m\u0065\u0020\u0028%\u0054\u0029", _dbdef.Get("\u0054\u0079\u0070\u0065"))
	}
	if _deddf.String() != "\u0050\u0061\u0067\u0065" {
		return _fd.New("\u0066\u0069e\u006c\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u0050\u0061\u0067\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069re\u0064\u0029")
	}
	_bbbdf := []_abf.PdfObjectName{"\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", "\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", "\u0043r\u006f\u0070\u0042\u006f\u0078", "\u0052\u006f\u0074\u0061\u0074\u0065"}
	_adagga, _cggdgb := _abf.GetIndirect(_dbdef.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	_acd.Log.Trace("P\u0061g\u0065\u0020\u0050\u0061\u0072\u0065\u006e\u0074:\u0020\u0025\u0054\u0020(%\u0076\u0029", _dbdef.Get("\u0050\u0061\u0072\u0065\u006e\u0074"), _cggdgb)
	for _cggdgb {
		_acd.Log.Trace("\u0050a\u0067e\u0020\u0050\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _adagga)
		_fbfac, _cbdfeb := _abf.GetDict(_adagga.PdfObject)
		if !_cbdfeb {
			return _fd.New("i\u006e\u0076\u0061\u006cid\u0020P\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		for _, _fbeda := range _bbbdf {
			_acd.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _fbeda)
			if _dbdef.Get(_fbeda) != nil {
				_acd.Log.Trace("\u002d \u0070a\u0067\u0065\u0020\u0068\u0061s\u0020\u0061l\u0072\u0065\u0061\u0064\u0079")
				continue
			}
			if _bbfg := _fbfac.Get(_fbeda); _bbfg != nil {
				_acd.Log.Trace("\u0049\u006e\u0068\u0065ri\u0074\u0069\u006e\u0067\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0025\u0073", _fbeda)
				_dbdef.Set(_fbeda, _bbfg)
			}
		}
		_adagga, _cggdgb = _abf.GetIndirect(_fbfac.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
		_acd.Log.Trace("\u004ee\u0078t\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u003a\u0020\u0025\u0054", _fbfac.Get("\u0050\u0061\u0072\u0065\u006e\u0074"))
	}
	_acd.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c \u0064\u006f\u006e\u0065")
	_dbdef.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _eaegf._cgeed)
	_fgeg.PdfObject = _dbdef
	_dcbede, _cafce := _abf.GetDict(_eaegf._cgeed.PdfObject)
	if !_cafce {
		return _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0020(\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0029")
	}
	_gbecfe, _cafce := _abf.GetArray(_dcbede.Get("\u004b\u0069\u0064\u0073"))
	if !_cafce {
		return _fd.New("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0050\u0061g\u0065\u0073\u0020\u004b\u0069\u0064\u0073\u0020o\u0062\u006a\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079\u0029")
	}
	_gbecfe.Append(_fgeg)
	_eaegf._aadb[_dbdef] = struct{}{}
	_eebab, _cafce := _abf.GetInt(_dcbede.Get("\u0043\u006f\u0075n\u0074"))
	if !_cafce {
		return _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u0050\u0061\u0067e\u0073\u0020\u0043\u006fu\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0029")
	}
	*_eebab = *_eebab + 1
	_eaegf.addObject(_fgeg)
	_gaafg := _eaegf.addObjects(_dbdef)
	if _gaafg != nil {
		return _gaafg
	}
	return nil
}

// SetAction sets the PDF action for the annotation link.
func (_cdegc *PdfAnnotationLink) SetAction(action *PdfAction) {
	_cdegc._bgad = action
	if action == nil {
		_cdegc.A = nil
	}
}

// GetNumComponents returns the number of color components (1 for CalGray).
func (_baed *PdfColorCalGray) GetNumComponents() int { return 1 }

type pdfCIDFontType2 struct {
	fontCommon
	_cfbae *_abf.PdfIndirectObject
	_geaca _cbb.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// Dictionary that defines the character collection of the CIDFont (required).
	// See Table 116.
	CIDSystemInfo *_abf.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW  _abf.PdfObject
	W   _abf.PdfObject
	DW2 _abf.PdfObject
	W2  _abf.PdfObject

	// CIDs to glyph indices mapping (optional).
	CIDToGIDMap _abf.PdfObject
	_ddeea      map[_cbb.CharCode]float64
	_cecdg      float64
	_dffcb      map[rune]int
}

// GetNumComponents returns the number of color components of the underlying
// colorspace device.
func (_gffd *PdfColorspaceSpecialPattern) GetNumComponents() int {
	return _gffd.UnderlyingCS.GetNumComponents()
}

// PdfColorspaceCalGray represents CalGray color space.
type PdfColorspaceCalGray struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      float64
	_dgcg      *_abf.PdfIndirectObject
}

// SetXObjectFormByName adds the provided XObjectForm to the page resources.
// The added XObjectForm is identified by the specified name.
func (_eggeg *PdfPageResources) SetXObjectFormByName(keyName _abf.PdfObjectName, xform *XObjectForm) error {
	_cccd := xform.ToPdfObject().(*_abf.PdfObjectStream)
	_efeea := _eggeg.SetXObjectByName(keyName, _cccd)
	return _efeea
}

func _bcee(_effd _gbe.StdFont) pdfFontSimple {
	_fabfa := _effd.Descriptor()
	return pdfFontSimple{fontCommon: fontCommon{_aacbc: "\u0054\u0079\u0070e\u0031", _ecggf: _effd.Name()}, _aecd: _effd.GetMetricsTable(), _abeb: &PdfFontDescriptor{FontName: _abf.MakeName(string(_fabfa.Name)), FontFamily: _abf.MakeName(_fabfa.Family), FontWeight: _abf.MakeFloat(float64(_fabfa.Weight)), Flags: _abf.MakeInteger(int64(_fabfa.Flags)), FontBBox: _abf.MakeArrayFromFloats(_fabfa.BBox[:]), ItalicAngle: _abf.MakeFloat(_fabfa.ItalicAngle), Ascent: _abf.MakeFloat(_fabfa.Ascent), Descent: _abf.MakeFloat(_fabfa.Descent), CapHeight: _abf.MakeFloat(_fabfa.CapHeight), XHeight: _abf.MakeFloat(_fabfa.XHeight), StemV: _abf.MakeFloat(_fabfa.StemV), StemH: _abf.MakeFloat(_fabfa.StemH)}, _edabc: _effd.Encoder()}
}

// IsCheckbox returns true if the button field represents a checkbox, false otherwise.
func (_bgcc *PdfFieldButton) IsCheckbox() bool { return _bgcc.GetType() == ButtonTypeCheckbox }

// ToPdfObject implements interface PdfModel.
func (_gaad *PdfAnnotationSquare) ToPdfObject() _abf.PdfObject {
	_gaad.PdfAnnotation.ToPdfObject()
	_eddd := _gaad._dbc
	_adee := _eddd.PdfObject.(*_abf.PdfObjectDictionary)
	if _gaad.PdfAnnotationMarkup != nil {
		_gaad.PdfAnnotationMarkup.appendToPdfDictionary(_adee)
	}
	_adee.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0053\u0071\u0075\u0061\u0072\u0065"))
	_adee.SetIfNotNil("\u0042\u0053", _gaad.BS)
	_adee.SetIfNotNil("\u0049\u0043", _gaad.IC)
	_adee.SetIfNotNil("\u0042\u0045", _gaad.BE)
	_adee.SetIfNotNil("\u0052\u0044", _gaad.RD)
	return _eddd
}

func (_agce *PdfReader) newPdfActionMovieFromDict(_dfd *_abf.PdfObjectDictionary) (*PdfActionMovie, error) {
	return &PdfActionMovie{Annotation: _dfd.Get("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e"), T: _dfd.Get("\u0054"), Operation: _dfd.Get("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn")}, nil
}

func (_dgcb *PdfColorspaceDeviceRGB) String() string {
	return "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
}

var _ pdfFont = (*pdfCIDFontType2)(nil)

// SetXObjectImageByName adds the provided XObjectImage to the page resources.
// The added XObjectImage is identified by the specified name.
func (_accf *PdfPageResources) SetXObjectImageByName(keyName _abf.PdfObjectName, ximg *XObjectImage) error {
	_dfgbe := ximg.ToPdfObject().(*_abf.PdfObjectStream)
	_eafdg := _accf.SetXObjectByName(keyName, _dfgbe)
	return _eafdg
}

func (_fdae *PdfAppender) updateObjectsDeep(_fgaa _abf.PdfObject, _fbe map[_abf.PdfObject]struct{}) {
	if _fbe == nil {
		_fbe = map[_abf.PdfObject]struct{}{}
	}
	if _, _gegc := _fbe[_fgaa]; _gegc || _fgaa == nil {
		return
	}
	_fbe[_fgaa] = struct{}{}
	_ccb := _abf.ResolveReferencesDeep(_fgaa, _fdae._gfeg)
	if _ccb != nil {
		_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ccb)
	}
	switch _gfab := _fgaa.(type) {
	case *_abf.PdfIndirectObject:
		switch {
		case _gfab.GetParser() == _fdae._agda._bebc:
			return
		case _gfab.GetParser() == _fdae.Reader._bebc:
			_fafa, _ := _fdae._agda.GetIndirectObjectByNumber(int(_gfab.ObjectNumber))
			_dfgb, _cccf := _fafa.(*_abf.PdfIndirectObject)
			if _cccf && _dfgb != nil {
				if _dfgb.PdfObject != _gfab.PdfObject && _dfgb.PdfObject.WriteString() != _gfab.PdfObject.WriteString() {
					if _be.Contains(_gfab.PdfObject.WriteString(), "\u002f\u0053\u0069\u0067") && _be.Contains(_gfab.PdfObject.WriteString(), "\u002f\u0053\u0075\u0062\u0074\u0079\u0070\u0065") {
						return
					}
					_fdae.addNewObject(_fgaa)
					_fdae._bge[_fgaa] = _gfab.ObjectNumber
				}
			}
		default:
			_fdae.addNewObject(_fgaa)
		}
		_fdae.updateObjectsDeep(_gfab.PdfObject, _fbe)
	case *_abf.PdfObjectArray:
		for _, _cdg := range _gfab.Elements() {
			_fdae.updateObjectsDeep(_cdg, _fbe)
		}
	case *_abf.PdfObjectDictionary:
		for _, _ffdc := range _gfab.Keys() {
			_fdae.updateObjectsDeep(_gfab.Get(_ffdc), _fbe)
		}
	case *_abf.PdfObjectStreams:
		if _gfab.GetParser() != _fdae._agda._bebc {
			for _, _ccbc := range _gfab.Elements() {
				_fdae.updateObjectsDeep(_ccbc, _fbe)
			}
		}
	case *_abf.PdfObjectStream:
		switch {
		case _gfab.GetParser() == _fdae._agda._bebc:
			return
		case _gfab.GetParser() == _fdae.Reader._bebc:
			if _ceeab, _dggdf := _fdae._agda._bebc.LookupByReference(_gfab.PdfObjectReference); _dggdf == nil {
				var _aeec bool
				if _adaa, _gfff := _abf.GetStream(_ceeab); _gfff && _dd.Equal(_adaa.Stream, _gfab.Stream) {
					_aeec = true
				}
				if _bedg, _gbab := _abf.GetDict(_ceeab); _aeec && _gbab {
					_aeec = _bedg.WriteString() == _gfab.PdfObjectDictionary.WriteString()
				}
				if _aeec {
					return
				}
			}
			if _gfab.ObjectNumber != 0 {
				_fdae._bge[_fgaa] = _gfab.ObjectNumber
			}
		default:
			if _, _dabfc := _fdae._gcba[_fgaa]; !_dabfc {
				_fdae.addNewObject(_fgaa)
			}
		}
		_fdae.updateObjectsDeep(_gfab.PdfObjectDictionary, _fbe)
	}
}

func _edagf(_fefef []byte) []byte {
	const _eebb = 52845
	const _acdba = 22719
	_dfdb := 55665
	for _, _bdbbc := range _fefef[:4] {
		_dfdb = (int(_bdbbc)+_dfdb)*_eebb + _acdba
	}
	_feefc := make([]byte, len(_fefef)-4)
	for _ecce, _dcceg := range _fefef[4:] {
		_feefc[_ecce] = byte(int(_dcceg) ^ _dfdb>>8)
		_dfdb = (int(_dcceg)+_dfdb)*_eebb + _acdba
	}
	return _feefc
}

func _cgefc(_cfbbe []byte) (_fadfa, _gbed string, _eggb error) {
	_acd.Log.Trace("g\u0065\u0074\u0041\u0053CI\u0049S\u0065\u0063\u0074\u0069\u006fn\u0073\u003a\u0020\u0025\u0064\u0020", len(_cfbbe))
	_ccgag := _fabag.FindIndex(_cfbbe)
	if _ccgag == nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0067\u0065\u0074\u0041\u0053\u0043\u0049\u0049\u0053\u0065\u0063\u0074\u0069o\u006e\u0073\u002e\u0020\u004e\u006f\u0020d\u0069\u0063\u0074\u002e")
		return "", "", _abf.ErrTypeError
	}
	_ggef := _ccgag[1]
	_bddea := _be.Index(string(_cfbbe[_ggef:]), _ffed)
	if _bddea < 0 {
		_fadfa = string(_cfbbe[_ggef:])
		return _fadfa, "", nil
	}
	_ggafc := _ggef + _bddea
	_fadfa = string(_cfbbe[_ggef:_ggafc])
	_dbbdd := _ggafc
	_bddea = _be.Index(string(_cfbbe[_dbbdd:]), _bdec)
	if _bddea < 0 {
		_acd.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0067e\u0074\u0041\u0053\u0043\u0049\u0049\u0053e\u0063\u0074\u0069\u006f\u006e\u0073\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eggb)
		return "", "", _abf.ErrTypeError
	}
	_bcddb := _dbbdd + _bddea
	_gbed = string(_cfbbe[_dbbdd:_bcddb])
	return _fadfa, _gbed, nil
}

// Field returns the parent form field of the widget annotation, if one exists.
// NOTE: the method returns nil if the parent form field has not been parsed.
func (_bcd *PdfAnnotationWidget) Field() *PdfField { return _bcd._agdc }

func (_ecg *PdfReader) newPdfAnnotationTrapNetFromDict(_dcfd *_abf.PdfObjectDictionary) (*PdfAnnotationTrapNet, error) {
	_gagggg := PdfAnnotationTrapNet{}
	return &_gagggg, nil
}

// DecodeArray returns the range of color component values in CalRGB colorspace.
func (_ggf *PdfColorspaceCalRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// DecodeArray returns the component range values for the Indexed colorspace.
func (_dcgf *PdfColorspaceSpecialIndexed) DecodeArray() []float64 {
	return []float64{0, float64(_dcgf.HiVal)}
}

func _bddec(_ecfc *_abf.PdfObjectDictionary, _deggg *fontCommon) (*pdfFontType3, error) {
	_cbbcb := _gcag(_deggg)
	_fefe := _ecfc.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
	if _fefe == nil {
		_fefe = _abf.MakeInteger(0)
	}
	_cbbcb.FirstChar = _fefe
	_fdac, _debb := _abf.GetIntVal(_fefe)
	if !_debb {
		_acd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _fefe)
		return nil, _abf.ErrTypeError
	}
	_bgacd := _cbb.CharCode(_fdac)
	_fefe = _ecfc.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
	if _fefe == nil {
		_fefe = _abf.MakeInteger(255)
	}
	_cbbcb.LastChar = _fefe
	_fdac, _debb = _abf.GetIntVal(_fefe)
	if !_debb {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _fefe)
		return nil, _abf.ErrTypeError
	}
	_dged := _cbb.CharCode(_fdac)
	_fefe = _ecfc.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _fefe != nil {
		_cbbcb.Resources = _fefe
	}
	_fefe = _ecfc.Get("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s")
	if _fefe == nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0068\u0061\u0072\u0050\u0072\u006f\u0063\u0073\u0020(%\u0076\u0029", _fefe)
		return nil, _abf.ErrNotSupported
	}
	_cbbcb.CharProcs = _fefe
	_fefe = _ecfc.Get("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078")
	if _fefe == nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0046\u006f\u006et\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0028\u0025\u0076\u0029", _fefe)
		return nil, _abf.ErrNotSupported
	}
	_cbbcb.FontMatrix = _fefe
	_cbbcb._ecgf = make(map[_cbb.CharCode]float64)
	_fefe = _ecfc.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
	if _fefe != nil {
		_cbbcb.Widths = _fefe
		_cfbc, _eacf := _abf.GetArray(_fefe)
		if !_eacf {
			_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _fefe)
			return nil, _abf.ErrTypeError
		}
		_eggf, _gecea := _cfbc.ToFloat64Array()
		if _gecea != nil {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _gecea
		}
		if len(_eggf) != int(_dged-_bgacd+1) {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _dged-_bgacd+1, len(_eggf))
			return nil, _abf.ErrRangeError
		}
		_ccbcd, _eacf := _abf.GetArray(_cbbcb.FontMatrix)
		if !_eacf {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _ccbcd)
			return nil, _gecea
		}
		_dcgec, _gecea := _ccbcd.ToFloat64Array()
		if _gecea != nil {
			_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020c\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0046o\u006e\u0074\u004d\u0061\u0074\u0072\u0069\u0078\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
			return nil, _gecea
		}
		_bgdcc := _ad.NewMatrix(_dcgec[0], _dcgec[1], _dcgec[2], _dcgec[3], _dcgec[4], _dcgec[5])
		for _affdcb, _gcef := range _eggf {
			_cccfd, _ := _bgdcc.Transform(_gcef, _gcef)
			_cbbcb._ecgf[_bgacd+_cbb.CharCode(_affdcb)] = _cccfd
		}
	}
	_cbbcb.Encoding = _abf.TraceToDirectObject(_ecfc.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	_edgc := _ecfc.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _edgc != nil {
		_cbbcb._dabca = _abf.TraceToDirectObject(_edgc)
		_ageed, _ggbdd := _cebb(_cbbcb._dabca, &_cbbcb.fontCommon)
		if _ggbdd != nil {
			return nil, _ggbdd
		}
		_cbbcb._aabfe = _ageed
	}
	if _edffd := _cbbcb._aabfe; _edffd != nil {
		_cbbcb._dgbd = _cbb.NewCMapEncoder("", nil, _edffd)
	} else {
		_cbbcb._dgbd = _cbb.NewPdfDocEncoder()
	}
	return _cbbcb, nil
}

func _eeggg(_egcfe _abf.PdfObject) (string, error) {
	_egcfe = _abf.TraceToDirectObject(_egcfe)
	switch _fcgbd := _egcfe.(type) {
	case *_abf.PdfObjectString:
		return _fcgbd.Str(), nil
	case *_abf.PdfObjectStream:
		_cbeeb, _cbcgbe := _abf.DecodeStream(_fcgbd)
		if _cbcgbe != nil {
			return "", _cbcgbe
		}
		return string(_cbeeb), nil
	}
	return "", _e.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _egcfe)
}

// GetShadingByName gets the shading specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_bgfe *PdfPageResources) GetShadingByName(keyName _abf.PdfObjectName) (*PdfShading, bool) {
	if _bgfe.Shading == nil {
		return nil, false
	}
	_aafa, _adcde := _abf.TraceToDirectObject(_bgfe.Shading).(*_abf.PdfObjectDictionary)
	if !_adcde {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0068\u0061d\u0069\u006e\u0067\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bgfe.Shading)
		return nil, false
	}
	if _bdfdc := _aafa.Get(keyName); _bdfdc != nil {
		_ecgad, _babde := _abaef(_bdfdc)
		if _babde != nil {
			_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0073\u0068\u0061\u0064\u0069\u006eg\u003a\u0020\u0025\u0076", _babde)
			return nil, false
		}
		return _ecgad, true
	}
	return nil, false
}

func (_gacc *PdfAnnotation) String() string {
	_fbg := ""
	_cea, _agd := _gacc.ToPdfObject().(*_abf.PdfIndirectObject)
	if _agd {
		_fbg = _e.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _gacc._edg, _cea.PdfObject.String())
	}
	return _fbg
}

// GetContainingPdfObject returns the containing object for the PdfField, i.e. an indirect object
// containing the field dictionary.
func (_aedg *PdfField) GetContainingPdfObject() _abf.PdfObject { return _aedg._dgdc }

func (_agbg *PdfReader) newPdfActionURIFromDict(_eebf *_abf.PdfObjectDictionary) (*PdfActionURI, error) {
	return &PdfActionURI{URI: _eebf.Get("\u0055\u0052\u0049"), IsMap: _eebf.Get("\u0049\u0073\u004da\u0070")}, nil
}

func (_eecc *PdfAnnotationMarkup) appendToPdfDictionary(_acfg *_abf.PdfObjectDictionary) {
	_acfg.SetIfNotNil("\u0054", _eecc.T)
	if _eecc.Popup != nil {
		_acfg.Set("\u0050\u006f\u0070u\u0070", _eecc.Popup.ToPdfObject())
	}
	_acfg.SetIfNotNil("\u0043\u0041", _eecc.CA)
	_acfg.SetIfNotNil("\u0052\u0043", _eecc.RC)
	_acfg.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _eecc.CreationDate)
	_acfg.SetIfNotNil("\u0049\u0052\u0054", _eecc.IRT)
	_acfg.SetIfNotNil("\u0053\u0075\u0062\u006a", _eecc.Subj)
	_acfg.SetIfNotNil("\u0052\u0054", _eecc.RT)
	_acfg.SetIfNotNil("\u0049\u0054", _eecc.IT)
	_acfg.SetIfNotNil("\u0045\u0078\u0044\u0061\u0074\u0061", _eecc.ExData)
}

// PdfAnnotation3D represents 3D annotations.
// (Section 13.6.2).
type PdfAnnotation3D struct {
	*PdfAnnotation
	T3DD _abf.PdfObject
	T3DV _abf.PdfObject
	T3DA _abf.PdfObject
	T3DI _abf.PdfObject
	T3DB _abf.PdfObject
}

// AddContentStreamByString adds content stream by string. Puts the content
// string into a stream object and points the content stream towards it.
func (_beedb *PdfPage) AddContentStreamByString(contentStr string) error {
	_cbgec, _ggcce := _abf.MakeStream([]byte(contentStr), _abf.NewFlateEncoder())
	if _ggcce != nil {
		return _ggcce
	}
	if _beedb.Contents == nil {
		_beedb.Contents = _cbgec
	} else {
		_addbg := _abf.TraceToDirectObject(_beedb.Contents)
		_bcffa, _dbddf := _addbg.(*_abf.PdfObjectArray)
		if !_dbddf {
			_bcffa = _abf.MakeArray(_addbg)
		}
		_bcffa.Append(_cbgec)
		_beedb.Contents = _bcffa
	}
	return nil
}

// NewPdfActionResetForm returns a new "reset form" action.
func NewPdfActionResetForm() *PdfActionResetForm {
	_eec := NewPdfAction()
	_cge := &PdfActionResetForm{}
	_cge.PdfAction = _eec
	_eec.SetContext(_cge)
	return _cge
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_gcefe *PdfShadingType1) ToPdfObject() _abf.PdfObject {
	_gcefe.PdfShading.ToPdfObject()
	_gcaee, _ebbge := _gcefe.getShadingDict()
	if _ebbge != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _gcefe.Domain != nil {
		_gcaee.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _gcefe.Domain)
	}
	if _gcefe.Matrix != nil {
		_gcaee.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _gcefe.Matrix)
	}
	if _gcefe.Function != nil {
		if len(_gcefe.Function) == 1 {
			_gcaee.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _gcefe.Function[0].ToPdfObject())
		} else {
			_ffgg := _abf.MakeArray()
			for _, _abaac := range _gcefe.Function {
				_ffgg.Append(_abaac.ToPdfObject())
			}
			_gcaee.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _ffgg)
		}
	}
	return _gcefe._eabcgc
}

// SetVersion sets the PDF version of the output file.
func (_fbbeg *PdfWriter) SetVersion(majorVersion, minorVersion int) {
	_fbbeg._ecfa.Major = majorVersion
	_fbbeg._ecfa.Minor = minorVersion
}

// SetRotation sets the rotation of all pages added to writer. The rotation is
// specified in degrees and must be a multiple of 90.
// The Rotate field of individual pages has priority over the global rotation.
func (_ggfcg *PdfWriter) SetRotation(rotate int64) error {
	_gaaaed, _gagef := _abf.GetDict(_ggfcg._cgeed)
	if !_gagef {
		return ErrTypeCheck
	}
	_gaaaed.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _abf.MakeInteger(rotate))
	return nil
}

// BorderEffect represents a border effect (Table 167 p. 395).
type BorderEffect int

// BorderStyle defines border type, typically used for annotations.
type BorderStyle int

// GetCatalogMarkInfo gets catalog MarkInfo object.
func (_gebee *PdfReader) GetCatalogMarkInfo() (_abf.PdfObject, bool) {
	if _gebee._dagde == nil {
		return nil, false
	}
	_fdfbg := _gebee._dagde.Get("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f")
	return _fdfbg, _fdfbg != nil
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 4 for a CMYK32 device.
func (_adbbf *PdfColorspaceDeviceCMYK) GetNumComponents() int { return 4 }

func (_bffa fontCommon) coreString() string {
	_abgaef := ""
	if _bffa._dcbaf != nil {
		_abgaef = _bffa._dcbaf.String()
	}
	return _e.Sprintf("\u0025#\u0071\u0020%\u0023\u0071\u0020%\u0071\u0020\u006f\u0062\u006a\u003d\u0025d\u0020\u0054\u006f\u0055\u006e\u0069c\u006f\u0064\u0065\u003d\u0025\u0074\u0020\u0066\u006c\u0061\u0067s\u003d\u0030\u0078\u0025\u0030\u0078\u0020\u0025\u0073", _bffa._aacbc, _bffa._ecggf, _bffa._dddac, _bffa._bgbd, _bffa._dabca != nil, _bffa.fontFlags(), _abgaef)
}

// SignatureHandlerDocMDP extends SignatureHandler with the ValidateWithOpts method for checking the DocMDP policy.
type SignatureHandlerDocMDP interface {
	SignatureHandler

	// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser
	// ValidateWithOpts shall contain Validate call
	ValidateWithOpts(_ddffb *PdfSignature, _fgcd Hasher, _caca SignatureHandlerDocMDPParams) (SignatureValidationResult, error)
}

// GetContainingPdfObject implements interface PdfModel.
func (_ffeg *PdfAnnotation) GetContainingPdfObject() _abf.PdfObject { return _ffeg._dbc }

// NewPdfAnnotationPolyLine returns a new polyline annotation.
func NewPdfAnnotationPolyLine() *PdfAnnotationPolyLine {
	_cff := NewPdfAnnotation()
	_dbb := &PdfAnnotationPolyLine{}
	_dbb.PdfAnnotation = _cff
	_dbb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_cff.SetContext(_dbb)
	return _dbb
}

// PdfColorspaceDeviceN represents a DeviceN color space. DeviceN color spaces are similar to Separation color
// spaces, except they can contain an arbitrary number of color components.
/*
	Format: [/DeviceN names alternateSpace tintTransform]
        or: [/DeviceN names alternateSpace tintTransform attributes]
*/
type PdfColorspaceDeviceN struct {
	ColorantNames  *_abf.PdfObjectArray
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	Attributes     *PdfColorspaceDeviceNAttributes
	_ddee          *_abf.PdfIndirectObject
}

// NewPdfAnnotationPrinterMark returns a new printermark annotation.
func NewPdfAnnotationPrinterMark() *PdfAnnotationPrinterMark {
	_ead := NewPdfAnnotation()
	_ebed := &PdfAnnotationPrinterMark{}
	_ebed.PdfAnnotation = _ead
	_ead.SetContext(_ebed)
	return _ebed
}

// B returns the value of the B component of the color.
func (_dbdg *PdfColorCalRGB) B() float64 { return _dbdg[1] }

// NewXObjectImageFromImage creates a new XObject Image from an image object
// with default options. If encoder is nil, uses raw encoding (none).
func NewXObjectImageFromImage(img *Image, cs PdfColorspace, encoder _abf.StreamEncoder) (*XObjectImage, error) {
	_dfbdab := NewXObjectImage()
	return UpdateXObjectImageFromImage(_dfbdab, img, cs, encoder)
}

func (_dage *PdfReader) newPdfOutlineItemFromIndirectObject(_bdfbba *_abf.PdfIndirectObject) (*PdfOutlineItem, error) {
	_dbccd, _aegba := _bdfbba.PdfObject.(*_abf.PdfObjectDictionary)
	if !_aegba {
		return nil, _e.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_ggeaa := NewPdfOutlineItem()
	_gbdee := _dbccd.Get("\u0054\u0069\u0074l\u0065")
	if _gbdee == nil {
		return nil, _e.Errorf("\u006d\u0069\u0073s\u0069\u006e\u0067\u0020\u0054\u0069\u0074\u006c\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0049\u0074\u0065\u006d\u0020\u0028r\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	_gccbb, _baefb := _abf.GetString(_gbdee)
	if !_baefb {
		return nil, _e.Errorf("\u0074\u0069\u0074le\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0054\u0029", _gbdee)
	}
	_ggeaa.Title = _gccbb
	if _bcec := _dbccd.Get("\u0043\u006f\u0075n\u0074"); _bcec != nil {
		_ecef, _efgcd := _bcec.(*_abf.PdfObjectInteger)
		if !_efgcd {
			return nil, _e.Errorf("\u0063o\u0075\u006e\u0074\u0020n\u006f\u0074\u0020\u0061\u006e \u0069n\u0074e\u0067\u0065\u0072\u0020\u0028\u0025\u0054)", _bcec)
		}
		_eadca := int64(*_ecef)
		_ggeaa.Count = &_eadca
	}
	if _egfda := _dbccd.Get("\u0044\u0065\u0073\u0074"); _egfda != nil {
		_ggeaa.Dest = _abf.ResolveReference(_egfda)
		if !_dage._abgge {
			_acbgf := _dage.traverseObjectData(_ggeaa.Dest)
			if _acbgf != nil {
				return nil, _acbgf
			}
		}
	}
	if _aebdb := _dbccd.Get("\u0041"); _aebdb != nil {
		_ggeaa.A = _abf.ResolveReference(_aebdb)
		if !_dage._abgge {
			_gdbf := _dage.traverseObjectData(_ggeaa.A)
			if _gdbf != nil {
				return nil, _gdbf
			}
		}
	}
	if _ccbfb := _dbccd.Get("\u0053\u0045"); _ccbfb != nil {
		_ggeaa.SE = nil
	}
	if _agfge := _dbccd.Get("\u0043"); _agfge != nil {
		_ggeaa.C = _abf.ResolveReference(_agfge)
	}
	if _dbce := _dbccd.Get("\u0046"); _dbce != nil {
		_ggeaa.F = _abf.ResolveReference(_dbce)
	}
	return _ggeaa, nil
}

func (_ebg *PdfReader) newPdfActionLaunchFromDict(_bda *_abf.PdfObjectDictionary) (*PdfActionLaunch, error) {
	_dee, _gga := _dgf(_bda.Get("\u0046"))
	if _gga != nil {
		return nil, _gga
	}
	return &PdfActionLaunch{Win: _bda.Get("\u0057\u0069\u006e"), Mac: _bda.Get("\u004d\u0061\u0063"), Unix: _bda.Get("\u0055\u006e\u0069\u0078"), NewWindow: _bda.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _dee}, nil
}

// SetAnnotations sets the annotations list.
func (_gbgba *PdfPage) SetAnnotations(annotations []*PdfAnnotation) { _gbgba._baagf = annotations }

var _ pdfFont = (*pdfFontType0)(nil)

// NewPdfFontFromTTFFile loads a TTF font file and returns a PdfFont type
// that can be used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTFFile.
func NewPdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_ecadc, _bbcg := _cf.Open(filePath)
	if _bbcg != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020T\u0054F\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _bbcg)
		return nil, _bbcg
	}
	defer _ecadc.Close()
	return NewPdfFontFromTTF(_ecadc)
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_abdfd *PdfShadingType2) ToPdfObject() _abf.PdfObject {
	_abdfd.PdfShading.ToPdfObject()
	_bgag, _edbcb := _abdfd.getShadingDict()
	if _edbcb != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _bgag == nil {
		_acd.Log.Error("\u0053\u0068\u0061\u0064in\u0067\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		return nil
	}
	if _abdfd.Coords != nil {
		_bgag.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _abdfd.Coords)
	}
	if _abdfd.Domain != nil {
		_bgag.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _abdfd.Domain)
	}
	if _abdfd.Function != nil {
		if len(_abdfd.Function) == 1 {
			_bgag.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _abdfd.Function[0].ToPdfObject())
		} else {
			_agddef := _abf.MakeArray()
			for _, _fcdc := range _abdfd.Function {
				_agddef.Append(_fcdc.ToPdfObject())
			}
			_bgag.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _agddef)
		}
	}
	if _abdfd.Extend != nil {
		_bgag.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _abdfd.Extend)
	}
	return _abdfd._eabcgc
}

func (_eaccc *PdfWriter) setDocumentIDs(_eacea, _gbaac string) {
	_eaccc._dedfdf = _abf.MakeArray(_abf.MakeHexString(_eacea), _abf.MakeHexString(_gbaac))
}

const (
	_ PdfOutputIntentType = iota
	PdfOutputIntentTypeA1
	PdfOutputIntentTypeA2
	PdfOutputIntentTypeA3
	PdfOutputIntentTypeA4
	PdfOutputIntentTypeX
)

func _bdda(_dffcd *PdfField, _gbbef _abf.PdfObject) {
	for _, _ecda := range _dffcd.Annotations {
		_ecda.AS = _gbbef
		_ecda.ToPdfObject()
	}
}

// ImageToRGB converts an Image in a given PdfColorspace to an RGB image.
func (_edfe *PdfColorspaceDeviceN) ImageToRGB(img Image) (Image, error) {
	_gdbaf := _gf.NewReader(img.getBase())
	_afagc := _gca.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, nil, img._gedg, img._ceeag)
	_eeff := _gf.NewWriter(_afagc)
	_gabee := _ge.Pow(2, float64(img.BitsPerComponent)) - 1
	_dfge := _edfe.GetNumComponents()
	_cggbb := make([]uint32, _dfge)
	_eadaa := make([]float64, _dfge)
	for {
		_gaga := _gdbaf.ReadSamples(_cggbb)
		if _gaga == _gc.EOF {
			break
		} else if _gaga != nil {
			return img, _gaga
		}
		for _afbda := 0; _afbda < _dfge; _afbda++ {
			_edfa := float64(_cggbb[_afbda]) / _gabee
			_eadaa[_afbda] = _edfa
		}
		_ebggb, _gaga := _edfe.TintTransform.Evaluate(_eadaa)
		if _gaga != nil {
			return img, _gaga
		}
		for _, _abade := range _ebggb {
			_abade = _ge.Min(_ge.Max(0, _abade), 1.0)
			if _gaga = _eeff.WriteSample(uint32(_abade * _gabee)); _gaga != nil {
				return img, _gaga
			}
		}
	}
	return _edfe.AlternateSpace.ImageToRGB(_cega(&_afagc))
}

// NewPdfActionJavaScript returns a new "javaScript" action.
func NewPdfActionJavaScript() *PdfActionJavaScript {
	_fcg := NewPdfAction()
	_eba := &PdfActionJavaScript{}
	_eba.PdfAction = _fcg
	_fcg.SetContext(_eba)
	return _eba
}

func (_bbfa *PdfReader) newPdfActionResetFormFromDict(_bebd *_abf.PdfObjectDictionary) (*PdfActionResetForm, error) {
	return &PdfActionResetForm{Fields: _bebd.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _bebd.Get("\u0046\u006c\u0061g\u0073")}, nil
}

// ImageHandler interface implements common image loading and processing tasks.
// Implementing as an interface allows for the possibility to use non-standard libraries for faster
// loading and processing of images.
type ImageHandler interface {
	// Read any image type and load into a new Image object.
	Read(_ccecf _gc.Reader) (*Image, error)

	// NewImageFromGoImage loads a NRGBA32 unidoc Image from a standard Go image structure.
	NewImageFromGoImage(_ddbgg _aa.Image) (*Image, error)

	// NewGrayImageFromGoImage loads a grayscale unidoc Image from a standard Go image structure.
	NewGrayImageFromGoImage(_adbee _aa.Image) (*Image, error)

	// Compress an image.
	Compress(_acce *Image, _gaadf int64) (*Image, error)
}

// PdfWriter handles outputing PDF content.
type PdfWriter struct {
	_cfdde         *_abf.PdfIndirectObject
	_cgeed         *_abf.PdfIndirectObject
	_aadb          map[_abf.PdfObject]struct{}
	_edcgc         []_abf.PdfObject
	_fdgae         map[_abf.PdfObject]struct{}
	_abcfb         []*_abf.PdfIndirectObject
	_gbcge         *PdfOutlineTreeNode
	_ddffc         *_abf.PdfObjectDictionary
	_daaae         []_abf.PdfObject
	_ddegc         *_abf.PdfIndirectObject
	_agfba         *_ac.Writer
	_dbfaad        int64
	_dacaeg        error
	_ddbgd         *_abf.PdfCrypt
	_cebae         *_abf.PdfObjectDictionary
	_dcdbb         *_abf.PdfIndirectObject
	_dedfdf        *_abf.PdfObjectArray
	_ecfa          _abf.Version
	_adceg         *bool
	_fadb          map[_abf.PdfObject][]*_abf.PdfObjectDictionary
	_bdgeb         *PdfAcroForm
	_cacbf         Optimizer
	_adgdc         StandardApplier
	_becfc         map[int]crossReference
	_cgded         int64
	ObjNumOffset   int
	_aegbd         bool
	_cagaf         _abf.XrefTable
	_ffgf          int64
	_cfecga        int64
	_deff          map[_abf.PdfObject]int64
	_dbdcg         map[_abf.PdfObject]struct{}
	_ceega         string
	_dgfea         []*PdfOutputIntent
	_fegae         bool
	_aefff, _cfbce string
}

// AddCerts adds certificates to DSS.
func (_gfcee *DSS) AddCerts(certs [][]byte) ([]*_abf.PdfObjectStream, error) {
	return _gfcee.add(&_gfcee.Certs, _gfcee._gcee, certs)
}

// ColorToRGB converts a ICCBased color to an RGB color.
func (_cabd *PdfColorspaceICCBased) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _cabd.Alternate == nil {
		_acd.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _cabd.N == 1 {
			_acd.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_acege := NewPdfColorspaceDeviceGray()
			return _acege.ColorToRGB(color)
		} else if _cabd.N == 3 {
			_acd.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return color, nil
		} else if _cabd.N == 4 {
			_acd.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_aeeee := NewPdfColorspaceDeviceCMYK()
			return _aeeee.ColorToRGB(color)
		} else {
			return nil, _fd.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_acd.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _cabd)
	return _cabd.Alternate.ColorToRGB(color)
}

// NewPdfAnnotationRedact returns a new redact annotation.
func NewPdfAnnotationRedact() *PdfAnnotationRedact {
	_ddcg := NewPdfAnnotation()
	_ggca := &PdfAnnotationRedact{}
	_ggca.PdfAnnotation = _ddcg
	_ggca.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ddcg.SetContext(_ggca)
	return _ggca
}

func (_eefeg *LTV) buildCertChain(_bcdgab, _eebcf []*_fa.Certificate) ([]*_fa.Certificate, map[string]*_fa.Certificate, error) {
	_cfdd := map[string]*_fa.Certificate{}
	for _, _cfecc := range _bcdgab {
		_cfdd[_cfecc.Subject.CommonName] = _cfecc
	}
	_cbedb := _bcdgab
	for _, _deaag := range _eebcf {
		_dggee := _deaag.Subject.CommonName
		if _, _ddabg := _cfdd[_dggee]; _ddabg {
			continue
		}
		_cfdd[_dggee] = _deaag
		_cbedb = append(_cbedb, _deaag)
	}
	if len(_cbedb) == 0 {
		return nil, nil, ErrSignNoCertificates
	}
	var _bdbc error
	for _fege := _cbedb[0]; _fege != nil && !_eefeg.CertClient.IsCA(_fege); {
		_gcadf, _cdecg := _cfdd[_fege.Issuer.CommonName]
		if !_cdecg {
			if _gcadf, _bdbc = _eefeg.CertClient.GetIssuer(_fege); _bdbc != nil {
				_acd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _bdbc)
				break
			}
			_cfdd[_fege.Issuer.CommonName] = _gcadf
			_cbedb = append(_cbedb, _gcadf)
		}
		_fege = _gcadf
	}
	return _cbedb, _cfdd, nil
}

func (_gac *PdfReader) newPdfActionSetOCGStateFromDict(_eff *_abf.PdfObjectDictionary) (*PdfActionSetOCGState, error) {
	return &PdfActionSetOCGState{State: _eff.Get("\u0053\u0074\u0061t\u0065"), PreserveRB: _eff.Get("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042")}, nil
}

// IsSimple returns true if `font` is a simple font.
func (_babcc *PdfFont) IsSimple() bool { _, _fdbf := _babcc._gedca.(*pdfFontSimple); return _fdbf }

// ToPdfObject converts the pdfCIDFontType2 to a PDF representation.
func (_abgaeg *pdfCIDFontType2) ToPdfObject() _abf.PdfObject {
	if _abgaeg._cfbae == nil {
		_abgaeg._cfbae = &_abf.PdfIndirectObject{}
	}
	_ffce := _abgaeg.baseFields().asPdfObjectDictionary("\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032")
	_abgaeg._cfbae.PdfObject = _ffce
	if _abgaeg.CIDSystemInfo != nil {
		_ffce.Set("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f", _abgaeg.CIDSystemInfo)
	}
	if _abgaeg.DW != nil {
		_ffce.Set("\u0044\u0057", _abgaeg.DW)
	}
	if _abgaeg.DW2 != nil {
		_ffce.Set("\u0044\u0057\u0032", _abgaeg.DW2)
	}
	if _abgaeg.W != nil {
		_ffce.Set("\u0057", _abgaeg.W)
	}
	if _abgaeg.W2 != nil {
		_ffce.Set("\u0057\u0032", _abgaeg.W2)
	}
	if _abgaeg.CIDToGIDMap != nil {
		_ffce.Set("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070", _abgaeg.CIDToGIDMap)
	}
	return _abgaeg._cfbae
}

func _gdcdc(_egff *_abf.PdfObjectDictionary, _ecdda *fontCommon) (*pdfFontType0, error) {
	_dggbd, _cbgbc := _abf.GetArray(_egff.Get("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073"))
	if !_cbgbc {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006cid\u0020\u0044\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073\u0020\u002d\u0020\u006e\u006f\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079 \u0025\u0073", _ecdda)
		return nil, _abf.ErrRangeError
	}
	if _dggbd.Len() != 1 {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0041\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020(%\u0064\u0029", _dggbd.Len())
		return nil, _abf.ErrRangeError
	}
	_eagfe, _ccfcd := _caece(_dggbd.Get(0), false)
	if _ccfcd != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046a\u0069\u006c\u0065d \u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u0020\u0066\u006f\u006et\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076 \u0025\u0073", _ccfcd, _ecdda)
		return nil, _ccfcd
	}
	_gbfec := _bedce(_ecdda)
	_gbfec.DescendantFont = _eagfe
	_bgbfa, _cbgbc := _abf.GetNameVal(_egff.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if _cbgbc {
		if _bgbfa == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048" || _bgbfa == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056" {
			_gbfec._edeaf = _cbb.NewIdentityTextEncoder(_bgbfa)
		} else if _bd.IsPredefinedCMap(_bgbfa) {
			_gbfec._fcfg, _ccfcd = _bd.LoadPredefinedCMap(_bgbfa)
			if _ccfcd != nil {
				_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _bgbfa, _ccfcd)
			}
		} else {
			_acd.Log.Debug("\u0055\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0063\u006da\u0070\u0020\u0025\u0071", _bgbfa)
		}
	}
	if _gdfb := _eagfe.baseFields()._aabfe; _gdfb != nil {
		if _aefad := _gdfb.Name(); _aefad == "\u0041d\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002d\u0055\u0043\u0053\u0032" || _aefad == "\u0041\u0064\u006f\u0062\u0065\u002d\u0047\u0042\u0031-\u0055\u0043\u0053\u0032" || _aefad == "\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031-\u0055\u0043\u0053\u0032" || _aefad == "\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006f\u0072\u0065\u0061\u0031-\u0055\u0043\u0053\u0032" {
			_gbfec._edeaf = _cbb.NewCMapEncoder(_bgbfa, _gbfec._fcfg, _gdfb)
		}
	}
	return _gbfec, nil
}

// ToPdfObject converts the font to a PDF representation.
func (_baag *pdfFontType0) ToPdfObject() _abf.PdfObject {
	if _baag._bgefb == nil {
		_baag._bgefb = &_abf.PdfIndirectObject{}
	}
	_adeca := _baag.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0030")
	_baag._bgefb.PdfObject = _adeca
	if _baag.Encoding != nil {
		_adeca.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _baag.Encoding)
	} else if _baag._edeaf != nil {
		_adeca.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _baag._edeaf.ToPdfObject())
	}
	if _baag.DescendantFont != nil {
		_adeca.Set("\u0044e\u0073c\u0065\u006e\u0064\u0061\u006e\u0074\u0046\u006f\u006e\u0074\u0073", _abf.MakeArray(_baag.DescendantFont.ToPdfObject()))
	}
	return _baag._bgefb
}

// AllFields returns a flattened list of all fields in the form.
func (_bedag *PdfAcroForm) AllFields() []*PdfField {
	if _bedag == nil {
		return nil
	}
	var _efdac []*PdfField
	if _bedag.Fields != nil {
		for _, _egdg := range *_bedag.Fields {
			_efdac = append(_efdac, _ddadg(_egdg)...)
		}
	}
	return _efdac
}

// SignatureHandlerDocMDPParams describe the specific parameters for the SignatureHandlerEx
// These parameters describe how to check the difference between revisions.
// Revisions of the document get from the PdfParser.
type SignatureHandlerDocMDPParams struct {
	Parser     *_abf.PdfParser
	DiffPolicy _df.DiffPolicy
}

// ToPdfObject implements interface PdfModel.
func (_cdfgd *PdfAnnotationRichMedia) ToPdfObject() _abf.PdfObject {
	_cdfgd.PdfAnnotation.ToPdfObject()
	_egafa := _cdfgd._dbc
	_ebcg := _egafa.PdfObject.(*_abf.PdfObjectDictionary)
	_ebcg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a"))
	_ebcg.SetIfNotNil("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073", _cdfgd.RichMediaSettings)
	_ebcg.SetIfNotNil("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074", _cdfgd.RichMediaContent)
	return _egafa
}

// ColorToRGB converts a color in Separation colorspace to RGB colorspace.
func (_eafa *PdfColorspaceSpecialSeparation) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _eafa.AlternateSpace == nil {
		return nil, _fd.New("\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020c\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	return _eafa.AlternateSpace.ColorToRGB(color)
}

// CharMetrics represents width and height metrics of a glyph.
type CharMetrics = _gbe.CharMetrics

func (_adaee *PdfWriter) seekByName(_adffaf _abf.PdfObject, _dbfd []string, _gccbg string) ([]_abf.PdfObject, error) {
	_acd.Log.Trace("\u0053\u0065\u0065\u006b\u0020\u0062\u0079\u0020\u006e\u0061\u006d\u0065.\u002e\u0020\u0025\u0054", _adffaf)
	var _fecg []_abf.PdfObject
	if _agdbc, _caafea := _adffaf.(*_abf.PdfIndirectObject); _caafea {
		return _adaee.seekByName(_agdbc.PdfObject, _dbfd, _gccbg)
	}
	if _dagcd, _aacbcc := _adffaf.(*_abf.PdfObjectStream); _aacbcc {
		return _adaee.seekByName(_dagcd.PdfObjectDictionary, _dbfd, _gccbg)
	}
	if _eagg, _bddfg := _adffaf.(*_abf.PdfObjectDictionary); _bddfg {
		_acd.Log.Trace("\u0044\u0069\u0063\u0074")
		for _, _daeea := range _eagg.Keys() {
			_gfdgb := _eagg.Get(_daeea)
			if string(_daeea) == _gccbg {
				_fecg = append(_fecg, _gfdgb)
			}
			for _, _debef := range _dbfd {
				if string(_daeea) == _debef {
					_acd.Log.Trace("\u0046\u006f\u006c\u006c\u006f\u0077\u0020\u006b\u0065\u0079\u0020\u0025\u0073", _debef)
					_ffeed, _bcagbb := _adaee.seekByName(_gfdgb, _dbfd, _gccbg)
					if _bcagbb != nil {
						return _fecg, _bcagbb
					}
					_fecg = append(_fecg, _ffeed...)
					break
				}
			}
		}
		return _fecg, nil
	}
	return _fecg, nil
}

const (
	TrappedUnknown PdfInfoTrapped = "\u0055n\u006b\u006e\u006f\u0077\u006e"
	TrappedTrue    PdfInfoTrapped = "\u0054\u0072\u0075\u0065"
	TrappedFalse   PdfInfoTrapped = "\u0046\u0061\u006cs\u0065"
)

func _bfdc(_dfgd *_abf.PdfObjectDictionary) (*PdfFieldText, error) {
	_ecag := &PdfFieldText{}
	_ecag.DA, _ = _abf.GetString(_dfgd.Get("\u0044\u0041"))
	_ecag.Q, _ = _abf.GetInt(_dfgd.Get("\u0051"))
	_ecag.DS, _ = _abf.GetString(_dfgd.Get("\u0044\u0053"))
	_ecag.RV = _dfgd.Get("\u0052\u0056")
	_ecag.MaxLen, _ = _abf.GetInt(_dfgd.Get("\u004d\u0061\u0078\u004c\u0065\u006e"))
	return _ecag, nil
}

// FieldValueProvider provides field values from a data source such as FDF, JSON or any other.
type FieldValueProvider interface {
	FieldValues() (map[string]_abf.PdfObject, error)
}

// PdfColorspaceSpecialPattern is a Pattern colorspace.
// Can be defined either as /Pattern or with an underlying colorspace [/Pattern cs].
type PdfColorspaceSpecialPattern struct {
	UnderlyingCS PdfColorspace
	_afca        *_abf.PdfIndirectObject
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_bbfde *PdfColorspaceSpecialSeparation) ToPdfObject() _abf.PdfObject {
	_gadg := _abf.MakeArray(_abf.MakeName("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"))
	_gadg.Append(_bbfde.ColorantName)
	_gadg.Append(_bbfde.AlternateSpace.ToPdfObject())
	_gadg.Append(_bbfde.TintTransform.ToPdfObject())
	if _bbfde._bbed != nil {
		_bbfde._bbed.PdfObject = _gadg
		return _bbfde._bbed
	}
	return _gadg
}

// PdfPageResources is a Page resources model.
// Implements PdfModel.
type PdfPageResources struct {
	ExtGState  _abf.PdfObject
	ColorSpace _abf.PdfObject
	Pattern    _abf.PdfObject
	Shading    _abf.PdfObject
	XObject    _abf.PdfObject
	Font       _abf.PdfObject
	ProcSet    _abf.PdfObject
	Properties _abf.PdfObject
	_gagb      *_abf.PdfObjectDictionary
	_aafff     *PdfPageResourcesColorspaces
}

func (_cdee *PdfColorspaceDeviceCMYK) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
}

// ToPdfObject implements interface PdfModel.
func (_ccg *PdfAnnotationText) ToPdfObject() _abf.PdfObject {
	_ccg.PdfAnnotation.ToPdfObject()
	_egbe := _ccg._dbc
	_gde := _egbe.PdfObject.(*_abf.PdfObjectDictionary)
	if _ccg.PdfAnnotationMarkup != nil {
		_ccg.PdfAnnotationMarkup.appendToPdfDictionary(_gde)
	}
	_gde.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0054\u0065\u0078\u0074"))
	_gde.SetIfNotNil("\u004f\u0070\u0065\u006e", _ccg.Open)
	_gde.SetIfNotNil("\u004e\u0061\u006d\u0065", _ccg.Name)
	_gde.SetIfNotNil("\u0053\u0074\u0061t\u0065", _ccg.State)
	_gde.SetIfNotNil("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c", _ccg.StateModel)
	return _egbe
}

// GetAlphabet returns a map of the runes in `text` and their frequencies.
func GetAlphabet(text string) map[rune]int {
	_ceaae := map[rune]int{}
	for _, _cdga := range text {
		_ceaae[_cdga]++
	}
	return _ceaae
}

// NewPdfAnnotationRichMedia returns a new rich media annotation.
func NewPdfAnnotationRichMedia() *PdfAnnotationRichMedia {
	_edeb := NewPdfAnnotation()
	_ddgc := &PdfAnnotationRichMedia{}
	_ddgc.PdfAnnotation = _edeb
	_edeb.SetContext(_ddgc)
	return _ddgc
}
func (_cbge *PdfColorspaceCalGray) String() string { return "\u0043a\u006c\u0047\u0072\u0061\u0079" }
func (_eccg *PdfColorspaceSpecialIndexed) String() string {
	return "\u0049n\u0064\u0065\u0078\u0065\u0064"
}

// PdfAnnotationPopup represents Popup annotations.
// (Section 12.5.6.14).
type PdfAnnotationPopup struct {
	*PdfAnnotation
	Parent _abf.PdfObject
	Open   _abf.PdfObject
}

// GetContext returns the action context which contains the specific type-dependent context.
// The context represents the subaction.
func (_dc *PdfAction) GetContext() PdfModel {
	if _dc == nil {
		return nil
	}
	return _dc._gfg
}

// NewPdfColorDeviceCMYK returns a new CMYK32 color.
func NewPdfColorDeviceCMYK(c, m, y, k float64) *PdfColorDeviceCMYK {
	_bdb := PdfColorDeviceCMYK{c, m, y, k}
	return &_bdb
}

// ToPdfObject implements interface PdfModel.
func (_dgc *PdfAnnotationCircle) ToPdfObject() _abf.PdfObject {
	_dgc.PdfAnnotation.ToPdfObject()
	_gabe := _dgc._dbc
	_fgga := _gabe.PdfObject.(*_abf.PdfObjectDictionary)
	_dgc.PdfAnnotationMarkup.appendToPdfDictionary(_fgga)
	_fgga.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0043\u0069\u0072\u0063\u006c\u0065"))
	_fgga.SetIfNotNil("\u0042\u0053", _dgc.BS)
	_fgga.SetIfNotNil("\u0049\u0043", _dgc.IC)
	_fgga.SetIfNotNil("\u0042\u0045", _dgc.BE)
	_fgga.SetIfNotNil("\u0052\u0044", _dgc.RD)
	return _gabe
}

func _gadf() *modelManager {
	_gfcge := modelManager{}
	_gfcge._baecg = map[PdfModel]_abf.PdfObject{}
	_gfcge._addgc = map[_abf.PdfObject]PdfModel{}
	return &_gfcge
}

// NewPdfColorDeviceGray returns a new grayscale color based on an input grayscale float value in range [0-1].
func NewPdfColorDeviceGray(grayVal float64) *PdfColorDeviceGray {
	_dceca := PdfColorDeviceGray(grayVal)
	return &_dceca
}

// ColorFromPdfObjects loads the color from PDF objects.
// The first objects (if present) represent the color in underlying colorspace.  The last one represents
// the name of the pattern.
func (_acc *PdfColorspaceSpecialPattern) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) < 1 {
		return nil, _fd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_bgee := &PdfColorPattern{}
	_bgac, _deaf := objects[len(objects)-1].(*_abf.PdfObjectName)
	if !_deaf {
		_acd.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020a\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", objects[len(objects)-1])
		return nil, ErrTypeCheck
	}
	_bgee.PatternName = *_bgac
	if len(objects) > 1 {
		_eceaf := objects[0 : len(objects)-1]
		if _acc.UnderlyingCS == nil {
			_acd.Log.Debug("P\u0061\u0074t\u0065\u0072\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0077\u0069\u0074\u0068\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006et\u0073\u0020\u0062\u0075\u0074\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0063\u0073\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
			return nil, _fd.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
		}
		_gdba, _cbcd := _acc.UnderlyingCS.ColorFromPdfObjects(_eceaf)
		if _cbcd != nil {
			_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0076\u0069\u0061\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0063\u0073\u003a\u0020\u0025\u0076", _cbcd)
			return nil, _cbcd
		}
		_bgee.Color = _gdba
	}
	return _bgee, nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_gdcc *PdfColorspaceDeviceGray) ToPdfObject() _abf.PdfObject {
	return _abf.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079")
}

// OutlineDest represents the destination of an outline item.
// It holds the page and the position on the page an outline item points to.
type OutlineDest struct {
	PageObj *_abf.PdfIndirectObject `json:"-"`
	Page    int64                   `json:"page"`
	Mode    string                  `json:"mode"`
	X       float64                 `json:"x"`
	Y       float64                 `json:"y"`
	Zoom    float64                 `json:"zoom"`
}

func (_ggcbd *Image) getSuitableEncoder() (_abf.StreamEncoder, error) {
	var (
		_acac, _bdfga = int(_ggcbd.Width), int(_ggcbd.Height)
		_fcbgc        = make(map[string]bool)
		_caafc        = true
		_gfee         = false
		_ebgc         = func() *_abf.DCTEncoder { return _abf.NewDCTEncoder() }
		_bcdgf        = func() *_abf.DCTEncoder { _bafcd := _abf.NewDCTEncoder(); _bafcd.BitsPerComponent = 16; return _bafcd }
	)
	for _gaced := 0; _gaced < _bdfga; _gaced++ {
		for _adfd := 0; _adfd < _acac; _adfd++ {
			_cffag, _caddc := _ggcbd.ColorAt(_adfd, _gaced)
			if _caddc != nil {
				return nil, _caddc
			}
			_addeg, _fced, _gbfbb, _fbec := _cffag.RGBA()
			if _caafc && (_addeg != _fced || _addeg != _gbfbb || _fced != _gbfbb) {
				_caafc = false
			}
			if !_gfee {
				switch _cffag.(type) {
				case _ga.NRGBA:
					_gfee = _fbec > 0
				}
			}
			_fcbgc[_e.Sprintf("\u0025\u0064\u002c\u0025\u0064\u002c\u0025\u0064", _addeg, _fced, _gbfbb)] = true
			if len(_fcbgc) > 2 && _gfee {
				return _bcdgf(), nil
			}
		}
	}
	if _gfee || len(_ggcbd._gedg) > 0 {
		return _abf.NewFlateEncoder(), nil
	}
	if len(_fcbgc) <= 2 {
		_dbgc := _ggcbd.ConvertToBinary()
		if _dbgc != nil {
			return nil, _dbgc
		}
		return _abf.NewJBIG2Encoder(), nil
	}
	if _caafc {
		return _ebgc(), nil
	}
	if _ggcbd.ColorComponents == 1 {
		if _ggcbd.BitsPerComponent == 1 {
			return _abf.NewJBIG2Encoder(), nil
		} else if _ggcbd.BitsPerComponent == 8 {
			_acggeg := _abf.NewDCTEncoder()
			_acggeg.ColorComponents = 1
			return _acggeg, nil
		}
	} else if _ggcbd.ColorComponents == 3 {
		if _ggcbd.BitsPerComponent == 8 {
			return _ebgc(), nil
		} else if _ggcbd.BitsPerComponent == 16 {
			return _bcdgf(), nil
		}
	} else if _ggcbd.ColorComponents == 4 {
		_gdgg := _bcdgf()
		_gdgg.ColorComponents = 4
		return _gdgg, nil
	}
	return _bcdgf(), nil
}

// ToPdfObject implements interface PdfModel.
func (_abb *PdfAnnotationStrikeOut) ToPdfObject() _abf.PdfObject {
	_abb.PdfAnnotation.ToPdfObject()
	_faad := _abb._dbc
	_eac := _faad.PdfObject.(*_abf.PdfObjectDictionary)
	_abb.PdfAnnotationMarkup.appendToPdfDictionary(_eac)
	_eac.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t"))
	_eac.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _abb.QuadPoints)
	return _faad
}

// PdfColorPatternType3 represents a color shading pattern type 3 (Radial).
type PdfColorPatternType3 struct {
	Color       PdfColor
	PatternName _abf.PdfObjectName
}

// WriteToFile writes the output PDF to file.
func (_gaffg *PdfWriter) WriteToFile(outputFilePath string) error {
	_eegc, _bfaaa := _cf.Create(outputFilePath)
	if _bfaaa != nil {
		return _bfaaa
	}
	defer _eegc.Close()
	return _gaffg.Write(_eegc)
}

var ErrColorOutOfRange = _fd.New("\u0063o\u006co\u0072\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")

// PdfShadingType5 is a Lattice-form Gouraud-shaded triangle mesh.
type PdfShadingType5 struct {
	*PdfShading
	BitsPerCoordinate *_abf.PdfObjectInteger
	BitsPerComponent  *_abf.PdfObjectInteger
	VerticesPerRow    *_abf.PdfObjectInteger
	Decode            *_abf.PdfObjectArray
	Function          []PdfFunction
}

// ContentStreamWrapper wraps the Page's contentstream into q ... Q blocks.
type ContentStreamWrapper interface{ WrapContentStream(_deeg *PdfPage) error }

// GetOutlines returns a high-level Outline object, based on the outline tree
// of the reader.
func (_dedfc *PdfReader) GetOutlines() (*Outline, error) {
	if _dedfc == nil {
		return nil, _fd.New("\u0063\u0061n\u006e\u006f\u0074\u0020c\u0072\u0065a\u0074\u0065\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u006e\u0069\u006c\u0020\u0072e\u0061\u0064\u0065\u0072")
	}
	_eabcg := _dedfc.GetOutlineTree()
	if _eabcg == nil {
		return nil, _fd.New("\u0074\u0068\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u0064\u006f\u0065\u0073\u0020n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065")
	}
	var _efafg func(_edgdb *PdfOutlineTreeNode, _fbfed *[]*OutlineItem)
	_efafg = func(_cdecf *PdfOutlineTreeNode, _bgdcg *[]*OutlineItem) {
		if _cdecf == nil {
			return
		}
		if _cdecf._aecec == nil {
			_acd.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020m\u0069\u0073\u0073\u0069ng \u006fut\u006c\u0069\u006e\u0065\u0020\u0065\u006etr\u0079\u0020\u0063\u006f\u006e\u0074\u0065x\u0074")
			return
		}
		var _cbag *OutlineItem
		if _cgdeg, _dggag := _cdecf._aecec.(*PdfOutlineItem); _dggag {
			_gddc := _cgdeg.Dest
			if (_gddc == nil || _abf.IsNullObject(_gddc)) && _cgdeg.A != nil {
				if _cefed, _egffd := _abf.GetDict(_cgdeg.A); _egffd {
					if _dacae, _cecdc := _abf.GetArray(_cefed.Get("\u0044")); _cecdc {
						_gddc = _dacae
					} else {
						_aebef, _ecbfe := _abf.GetString(_cefed.Get("\u0044"))
						if !_ecbfe {
							return
						}
						_egcdg, _ecbfe := _dedfc._dagde.Get("\u004e\u0061\u006de\u0073").(*_abf.PdfObjectReference)
						if !_ecbfe {
							return
						}
						_ebccec, _bebdff := _dedfc._bebc.LookupByReference(*_egcdg)
						if _bebdff != nil {
							_acd.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0072\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0028\u0025\u0073\u0029", _bebdff.Error())
							return
						}
						_bfde, _ecbfe := _ebccec.(*_abf.PdfIndirectObject)
						if !_ecbfe {
							return
						}
						_dfbda := map[_abf.PdfObject]struct{}{}
						_bebdff = _dedfc.buildNameNodes(_bfde, _dfbda)
						if _bebdff != nil {
							_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006ea\u006d\u0065\u0020\u006e\u006fd\u0065\u0073 \u0028\u0025\u0073\u0029", _bebdff.Error())
							return
						}
						for _fbfaa := range _dfbda {
							_afgf, _ggfbf := _abf.GetDict(_fbfaa)
							if !_ggfbf {
								continue
							}
							_ebcdf, _ggfbf := _abf.GetArray(_afgf.Get("\u004e\u0061\u006de\u0073"))
							if !_ggfbf {
								continue
							}
							for _aedc, _gaefe := range _ebcdf.Elements() {
								switch _gaefe.(type) {
								case *_abf.PdfObjectString:
									if _gaefe.String() == _aebef.String() {
										if _bcfad := _ebcdf.Get(_aedc + 1); _bcfad != nil {
											if _abgad, _bffec := _abf.GetDict(_bcfad); _bffec {
												_gddc = _abgad.Get("\u0044")
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
			var _cgaca OutlineDest
			if _gddc != nil && !_abf.IsNullObject(_gddc) {
				if _bddeab, _faaab := _aaagb(_gddc, _dedfc); _faaab == nil {
					_cgaca = *_bddeab
				} else {
					_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020p\u0061\u0072\u0073\u0065\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065\u0073\u0074\u0020\u0028\u0025\u0076\u0029\u003a\u0020\u0025\u0076", _gddc, _faaab)
				}
			}
			_cbag = NewOutlineItem(_cgdeg.Title.Decoded(), _cgaca)
			*_bgdcg = append(*_bgdcg, _cbag)
			if _cgdeg.Next != nil {
				_efafg(_cgdeg.Next, _bgdcg)
			}
		}
		if _cdecf.First != nil {
			if _cbag != nil {
				_bgdcg = &_cbag.Entries
			}
			_efafg(_cdecf.First, _bgdcg)
		}
	}
	_caga := NewOutline()
	_efafg(_eabcg, &_caga.Entries)
	return _caga, nil
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a grayscale device.
func (_cffb *PdfColorspaceDeviceGray) GetNumComponents() int { return 1 }

// SetFlag sets the flag for the field.
func (_ceebd *PdfField) SetFlag(flag FieldFlag) { _ceebd.Ff = _abf.MakeInteger(int64(flag)) }

// SetShadingByName sets a shading resource specified by keyName.
func (_cdadf *PdfPageResources) SetShadingByName(keyName _abf.PdfObjectName, shadingObj _abf.PdfObject) error {
	if _cdadf.Shading == nil {
		_cdadf.Shading = _abf.MakeDict()
	}
	_bdbgd, _gagege := _abf.GetDict(_cdadf.Shading)
	if !_gagege {
		return _abf.ErrTypeError
	}
	_bdbgd.Set(keyName, shadingObj)
	return nil
}

func _abaef(_fafga _abf.PdfObject) (*PdfShading, error) {
	_bcbgbe := &PdfShading{}
	var _ffdad *_abf.PdfObjectDictionary
	if _deafd, _cfbfe := _abf.GetIndirect(_fafga); _cfbfe {
		_bcbgbe._eabcgc = _deafd
		_fffcg, _cdffd := _deafd.PdfObject.(*_abf.PdfObjectDictionary)
		if !_cdffd {
			_acd.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _abf.ErrTypeError
		}
		_ffdad = _fffcg
	} else if _debdc, _adgfc := _abf.GetStream(_fafga); _adgfc {
		_bcbgbe._eabcgc = _debdc
		_ffdad = _debdc.PdfObjectDictionary
	} else if _dcbc, _cgbgf := _abf.GetDict(_fafga); _cgbgf {
		_bcbgbe._eabcgc = _dcbc
		_ffdad = _dcbc
	} else {
		_acd.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", _fafga)
		return nil, _abf.ErrTypeError
	}
	if _ffdad == nil {
		_acd.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _fd.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_fafga = _ffdad.Get("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _fafga == nil {
		_acd.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073si\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fafga = _abf.TraceToDirectObject(_fafga)
	_faegf, _cddgg := _fafga.(*_abf.PdfObjectInteger)
	if !_cddgg {
		_acd.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066o\u0072 \u0073h\u0061d\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _fafga)
		return nil, _abf.ErrTypeError
	}
	if *_faegf < 1 || *_faegf > 7 {
		_acd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u006e\u006ft\u0020\u0031\u002d\u0037\u0020(\u0067\u006ft\u0020\u0025\u0064\u0029", *_faegf)
		return nil, _abf.ErrTypeError
	}
	_bcbgbe.ShadingType = _faegf
	_fafga = _ffdad.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")
	if _fafga == nil {
		_acd.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0020\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065\u0020e\u006e\u0074\u0072\u0079\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bcaf, _cgga := NewPdfColorspaceFromPdfObject(_fafga)
	if _cgga != nil {
		_acd.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065: \u0025\u0076", _cgga)
		return nil, _cgga
	}
	_bcbgbe.ColorSpace = _bcaf
	_fafga = _ffdad.Get("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064")
	if _fafga != nil {
		_fafga = _abf.TraceToDirectObject(_fafga)
		_abbdb, _fbfca := _fafga.(*_abf.PdfObjectArray)
		if !_fbfca {
			_acd.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _fafga)
			return nil, _abf.ErrTypeError
		}
		_bcbgbe.Background = _abbdb
	}
	_fafga = _ffdad.Get("\u0042\u0042\u006f\u0078")
	if _fafga != nil {
		_fafga = _abf.TraceToDirectObject(_fafga)
		_eafeb, _daabf := _fafga.(*_abf.PdfObjectArray)
		if !_daabf {
			_acd.Log.Debug("\u0042\u0061\u0063\u006b\u0067r\u006f\u0075\u006e\u0064\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _fafga)
			return nil, _abf.ErrTypeError
		}
		_gbcff, _eaebg := NewPdfRectangle(*_eafeb)
		if _eaebg != nil {
			_acd.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _eaebg)
			return nil, _eaebg
		}
		_bcbgbe.BBox = _gbcff
	}
	_fafga = _ffdad.Get("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s")
	if _fafga != nil {
		_fafga = _abf.TraceToDirectObject(_fafga)
		_gbcg, _cadg := _fafga.(*_abf.PdfObjectBool)
		if !_cadg {
			_acd.Log.Debug("A\u006e\u0074\u0069\u0041\u006c\u0069\u0061\u0073\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020ty\u0070\u0065\u002c\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020bo\u006f\u006c \u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _fafga)
			return nil, _abf.ErrTypeError
		}
		_bcbgbe.AntiAlias = _gbcg
	}
	switch *_faegf {
	case 1:
		_egbebd, _deac := _eccc(_ffdad)
		if _deac != nil {
			return nil, _deac
		}
		_egbebd.PdfShading = _bcbgbe
		_bcbgbe._eabd = _egbebd
		return _bcbgbe, nil
	case 2:
		_gecfc, _cadfa := _eacca(_ffdad)
		if _cadfa != nil {
			return nil, _cadfa
		}
		_gecfc.PdfShading = _bcbgbe
		_bcbgbe._eabd = _gecfc
		return _bcbgbe, nil
	case 3:
		_faaed, _ccbdcb := _fecfd(_ffdad)
		if _ccbdcb != nil {
			return nil, _ccbdcb
		}
		_faaed.PdfShading = _bcbgbe
		_bcbgbe._eabd = _faaed
		return _bcbgbe, nil
	case 4:
		_caceg, _bbddcb := _faefbc(_ffdad)
		if _bbddcb != nil {
			return nil, _bbddcb
		}
		_caceg.PdfShading = _bcbgbe
		_bcbgbe._eabd = _caceg
		return _bcbgbe, nil
	case 5:
		_fggag, _afbfc := _daacfg(_ffdad)
		if _afbfc != nil {
			return nil, _afbfc
		}
		_fggag.PdfShading = _bcbgbe
		_bcbgbe._eabd = _fggag
		return _bcbgbe, nil
	case 6:
		_ecgec, _aabec := _gabff(_ffdad)
		if _aabec != nil {
			return nil, _aabec
		}
		_ecgec.PdfShading = _bcbgbe
		_bcbgbe._eabd = _ecgec
		return _bcbgbe, nil
	case 7:
		_ffbcd, _cebea := _fdade(_ffdad)
		if _cebea != nil {
			return nil, _cebea
		}
		_ffbcd.PdfShading = _bcbgbe
		_bcbgbe._eabd = _ffbcd
		return _bcbgbe, nil
	}
	return nil, _fd.New("u\u006ek\u006e\u006f\u0077\u006e\u0020\u0073\u0068\u0061d\u0069\u006e\u0067\u0020ty\u0070\u0065")
}

// AlphaMap performs mapping of alpha data for transformations. Allows custom filtering of alpha data etc.
func (_bcceg *Image) AlphaMap(mapFunc AlphaMapFunc) {
	for _ggfca, _gggbg := range _bcceg._gedg {
		_bcceg._gedg[_ggfca] = mapFunc(_gggbg)
	}
}

// Flags returns the field flags for the field accounting for any inherited flags.
func (_fbgce *PdfField) Flags() FieldFlag {
	var _gbfdd FieldFlag
	_ccgf, _gefe := _fbgce.inherit(func(_cffec *PdfField) bool {
		if _cffec.Ff != nil {
			_gbfdd = FieldFlag(*_cffec.Ff)
			return true
		}
		return false
	})
	if _gefe != nil {
		_acd.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0069\u006e\u0067\u0020\u0066\u006c\u0061\u0067\u0073\u0020\u0076\u0069\u0061\u0020\u0069\u006e\u0068\u0065\u0072\u0069t\u0061\u006e\u0063\u0065\u003a\u0020\u0025\u0076", _gefe)
	}
	if !_ccgf {
		_acd.Log.Trace("N\u006f\u0020\u0066\u0069\u0065\u006cd\u0020\u0066\u006c\u0061\u0067\u0073 \u0066\u006f\u0075\u006e\u0064\u0020\u002d \u0061\u0073\u0073\u0075\u006d\u0065\u0020\u0063\u006c\u0065a\u0072")
	}
	return _gbfdd
}

func (_aafc *PdfReader) newPdfAnnotationScreenFromDict(_gba *_abf.PdfObjectDictionary) (*PdfAnnotationScreen, error) {
	_fdbc := PdfAnnotationScreen{}
	_fdbc.T = _gba.Get("\u0054")
	_fdbc.MK = _gba.Get("\u004d\u004b")
	_fdbc.A = _gba.Get("\u0041")
	_fdbc.AA = _gba.Get("\u0041\u0041")
	return &_fdbc, nil
}

// Size returns the width and the height of the page. The method reports
// the page dimensions as displayed by a PDF viewer (i.e. page rotation is
// taken into account).
func (_fgdea *PdfPage) Size() (float64, float64, error) {
	_ccea, _dbdga := _fgdea.GetMediaBox()
	if _dbdga != nil {
		return 0, 0, _dbdga
	}
	_bacbd, _ebffc := _ccea.Width(), _ccea.Height()
	_bgccg, _dbdga := _fgdea.GetRotate()
	if _dbdga != nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _dbdga.Error())
	}
	if _dfagg := _bgccg; _dfagg%360 != 0 && _dfagg%90 == 0 {
		if _eggga := (360 + _dfagg%360) % 360; _eggga == 90 || _eggga == 270 {
			_bacbd, _ebffc = _ebffc, _bacbd
		}
	}
	return _bacbd, _ebffc, nil
}

// String returns a string representation of what flags are set.
func (_gbgb FieldFlag) String() string {
	_dedae := ""
	if _gbgb == FieldFlagClear {
		_dedae = "\u0043\u006c\u0065a\u0072"
		return _dedae
	}
	if _gbgb&FieldFlagReadOnly > 0 {
		_dedae += "\u007cR\u0065\u0061\u0064\u004f\u006e\u006cy"
	}
	if _gbgb&FieldFlagRequired > 0 {
		_dedae += "\u007cR\u0065\u0071\u0075\u0069\u0072\u0065d"
	}
	if _gbgb&FieldFlagNoExport > 0 {
		_dedae += "\u007cN\u006f\u0045\u0078\u0070\u006f\u0072t"
	}
	if _gbgb&FieldFlagNoToggleToOff > 0 {
		_dedae += "\u007c\u004e\u006f\u0054\u006f\u0067\u0067\u006c\u0065T\u006f\u004f\u0066\u0066"
	}
	if _gbgb&FieldFlagRadio > 0 {
		_dedae += "\u007c\u0052\u0061\u0064\u0069\u006f"
	}
	if _gbgb&FieldFlagPushbutton > 0 {
		_dedae += "|\u0050\u0075\u0073\u0068\u0062\u0075\u0074\u0074\u006f\u006e"
	}
	if _gbgb&FieldFlagRadiosInUnision > 0 {
		_dedae += "\u007c\u0052a\u0064\u0069\u006fs\u0049\u006e\u0055\u006e\u0069\u0073\u0069\u006f\u006e"
	}
	if _gbgb&FieldFlagMultiline > 0 {
		_dedae += "\u007c\u004d\u0075\u006c\u0074\u0069\u006c\u0069\u006e\u0065"
	}
	if _gbgb&FieldFlagPassword > 0 {
		_dedae += "\u007cP\u0061\u0073\u0073\u0077\u006f\u0072d"
	}
	if _gbgb&FieldFlagFileSelect > 0 {
		_dedae += "|\u0046\u0069\u006c\u0065\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _gbgb&FieldFlagDoNotScroll > 0 {
		_dedae += "\u007c\u0044\u006fN\u006f\u0074\u0053\u0063\u0072\u006f\u006c\u006c"
	}
	if _gbgb&FieldFlagComb > 0 {
		_dedae += "\u007c\u0043\u006fm\u0062"
	}
	if _gbgb&FieldFlagRichText > 0 {
		_dedae += "\u007cR\u0069\u0063\u0068\u0054\u0065\u0078t"
	}
	if _gbgb&FieldFlagDoNotSpellCheck > 0 {
		_dedae += "\u007c\u0044o\u004e\u006f\u0074S\u0070\u0065\u006c\u006c\u0043\u0068\u0065\u0063\u006b"
	}
	if _gbgb&FieldFlagCombo > 0 {
		_dedae += "\u007c\u0043\u006f\u006d\u0062\u006f"
	}
	if _gbgb&FieldFlagEdit > 0 {
		_dedae += "\u007c\u0045\u0064i\u0074"
	}
	if _gbgb&FieldFlagSort > 0 {
		_dedae += "\u007c\u0053\u006fr\u0074"
	}
	if _gbgb&FieldFlagMultiSelect > 0 {
		_dedae += "\u007c\u004d\u0075l\u0074\u0069\u0053\u0065\u006c\u0065\u0063\u0074"
	}
	if _gbgb&FieldFlagCommitOnSelChange > 0 {
		_dedae += "\u007cC\u006fm\u006d\u0069\u0074\u004f\u006eS\u0065\u006cC\u0068\u0061\u006e\u0067\u0065"
	}
	return _be.Trim(_dedae, "\u007c")
}

// NewPdfColorLab returns a new Lab color.
func NewPdfColorLab(l, a, b float64) *PdfColorLab { _dfagb := PdfColorLab{l, a, b}; return &_dfagb }

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_fcge *PdfColorspaceSpecialIndexed) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_ebgg, _aadg := _abf.GetNumbersAsFloat(objects)
	if _aadg != nil {
		return nil, _aadg
	}
	return _fcge.ColorFromFloats(_ebgg)
}

func _fccaa(_dafe *[]*PdfField, _bcaab FieldFilterFunc, _bcdgc bool) []*PdfField {
	if _dafe == nil {
		return nil
	}
	_bdgdg := *_dafe
	if len(*_dafe) == 0 {
		return nil
	}
	_dffe := _bdgdg[:0]
	if _bcaab == nil {
		_bcaab = func(*PdfField) bool { return true }
	}
	var _ecddf []*PdfField
	for _, _dcbgg := range _bdgdg {
		_cfbgd := _bcaab(_dcbgg)
		if _cfbgd {
			_ecddf = append(_ecddf, _dcbgg)
			if len(_dcbgg.Kids) > 0 {
				_ecddf = append(_ecddf, _fccaa(&_dcbgg.Kids, _bcaab, _bcdgc)...)
			}
		}
		if !_bcdgc || !_cfbgd || len(_dcbgg.Kids) > 0 {
			_dffe = append(_dffe, _dcbgg)
		}
	}
	*_dafe = _dffe
	return _ecddf
}

var _becf = _af.MustCompile("\u005b\\\u006e\u005c\u0072\u005d\u002b")

// M returns the value of the magenta component of the color.
func (_dbcb *PdfColorDeviceCMYK) M() float64 { return _dbcb[1] }

// ToInteger convert to an integer format.
func (_gfbb *PdfColorDeviceRGB) ToInteger(bits int) [3]uint32 {
	_aagg := _ge.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_aagg * _gfbb.R()), uint32(_aagg * _gfbb.G()), uint32(_aagg * _gfbb.B())}
}

// PdfFieldSignature signature field represents digital signatures and optional data for authenticating
// the name of the signer and verifying document contents.
type PdfFieldSignature struct {
	*PdfField
	*PdfAnnotationWidget
	V    *PdfSignature
	Lock *_abf.PdfIndirectObject
	SV   *_abf.PdfIndirectObject
}

func _efcef() string { _gaabd.Lock(); defer _gaabd.Unlock(); return _geggga }

// PdfPageResourcesColorspaces contains the colorspace in the PdfPageResources.
// Needs to have matching name and colorspace map entry. The Names define the order.
type PdfPageResourcesColorspaces struct {
	Names       []string
	Colorspaces map[string]PdfColorspace
	_cebc       *_abf.PdfIndirectObject
}

// SetSubtype sets the Subtype S for given PdfOutputIntent.
func (_fdge *PdfOutputIntent) SetSubtype(subtype PdfOutputIntentType) error {
	if !subtype.IsValid() {
		return _fd.New("\u0070\u0072o\u0076\u0069\u0064\u0065d\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u004f\u0075t\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0053\u0075b\u0054\u0079\u0070\u0065")
	}
	_fdge.S = subtype
	return nil
}

// PdfAnnotationStrikeOut represents StrikeOut annotations.
// (Section 12.5.6.10).
type PdfAnnotationStrikeOut struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _abf.PdfObject
}

// GetRuneMetrics returns the character metrics for the rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_eaag pdfFontSimple) GetRuneMetrics(r rune) (_gbe.CharMetrics, bool) {
	if _eaag._aecd != nil {
		_agcec, _aada := _eaag._aecd.Read(r)
		if _aada {
			return _agcec, true
		}
	}
	_cdfe := _eaag.Encoder()
	if _cdfe == nil {
		_acd.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _eaag)
		return _gbe.CharMetrics{}, false
	}
	_efab, _dfagf := _cdfe.RuneToCharcode(r)
	if !_dfagf {
		if r != ' ' {
			_acd.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _eaag)
		}
		return _gbe.CharMetrics{}, false
	}
	_ggdbf, _efga := _eaag.GetCharMetrics(_efab)
	return _ggdbf, _efga
}

// Insert adds an outline item as a child of the current outline item,
// at the specified index.
func (_bbde *OutlineItem) Insert(index uint, item *OutlineItem) {
	_affa := uint(len(_bbde.Entries))
	if index > _affa {
		index = _affa
	}
	_bbde.Entries = append(_bbde.Entries[:index], append([]*OutlineItem{item}, _bbde.Entries[index:]...)...)
}

func _egeeb(_degd _abf.PdfObject) (*PdfColorspaceDeviceN, error) {
	_daag := NewPdfColorspaceDeviceN()
	if _dbac, _gcbd := _degd.(*_abf.PdfIndirectObject); _gcbd {
		_daag._ddee = _dbac
	}
	_degd = _abf.TraceToDirectObject(_degd)
	_edea, _bagd := _degd.(*_abf.PdfObjectArray)
	if !_bagd {
		return nil, _e.Errorf("\u0064\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if _edea.Len() != 4 && _edea.Len() != 5 {
		return nil, _e.Errorf("\u0064\u0065\u0076ic\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0049\u006ec\u006fr\u0072e\u0063t\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_degd = _edea.Get(0)
	_ddcb, _bagd := _degd.(*_abf.PdfObjectName)
	if !_bagd {
		return nil, _e.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_ddcb != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
		return nil, _e.Errorf("\u0064\u0065v\u0069\u0063\u0065\u004e\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020na\u006d\u0065")
	}
	_degd = _edea.Get(1)
	_degd = _abf.TraceToDirectObject(_degd)
	_decd, _bagd := _degd.(*_abf.PdfObjectArray)
	if !_bagd {
		return nil, _e.Errorf("\u0064\u0065\u0076i\u0063\u0065\u004e\u0020C\u0053\u003a\u0020\u0049\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0061\u006d\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_daag.ColorantNames = _decd
	_degd = _edea.Get(2)
	_dfadg, _gcff := NewPdfColorspaceFromPdfObject(_degd)
	if _gcff != nil {
		return nil, _gcff
	}
	_daag.AlternateSpace = _dfadg
	_eega, _gcff := _ebedg(_edea.Get(3))
	if _gcff != nil {
		return nil, _gcff
	}
	_daag.TintTransform = _eega
	if _edea.Len() == 5 {
		_gdcf, _cffa := _bgab(_edea.Get(4))
		if _cffa != nil {
			return nil, _cffa
		}
		_daag.Attributes = _gdcf
	}
	return _daag, nil
}

// PdfActionResetForm represents a resetForm action.
type PdfActionResetForm struct {
	*PdfAction
	Fields _abf.PdfObject
	Flags  _abf.PdfObject
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

// SetBorderWidth sets the style's border width.
func (_ebdca *PdfBorderStyle) SetBorderWidth(width float64) { _ebdca.W = &width }

var _bdgdc = map[string]struct{}{"\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}, "\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074\u0045\u006e\u0063o\u0064\u0069\u006e\u0067": {}, "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067": {}}

// StdFontName represents name of a standard font.
type StdFontName = _gbe.StdFontName

func _eacca(_bfegc *_abf.PdfObjectDictionary) (*PdfShadingType2, error) {
	_fgdd := PdfShadingType2{}
	_dabba := _bfegc.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _dabba == nil {
		_acd.Log.Debug("R\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u003a\u0020\u0020\u0043\u006f\u006f\u0072d\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_bcdef, _fbffeg := _dabba.(*_abf.PdfObjectArray)
	if !_fbffeg {
		_acd.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _dabba)
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _bcdef.Len() != 4 {
		_acd.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0034\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _bcdef.Len())
		return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	}
	_fgdd.Coords = _bcdef
	if _babga := _bfegc.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _babga != nil {
		_babga = _abf.TraceToDirectObject(_babga)
		_aagff, _agbfg := _babga.(*_abf.PdfObjectArray)
		if !_agbfg {
			_acd.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _babga)
			return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fgdd.Domain = _aagff
	}
	_dabba = _bfegc.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _dabba == nil {
		_acd.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_fgdd.Function = []PdfFunction{}
	if _afgbd, _adbdb := _dabba.(*_abf.PdfObjectArray); _adbdb {
		for _, _gdcee := range _afgbd.Elements() {
			_dccgb, _acbc := _ebedg(_gdcee)
			if _acbc != nil {
				_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _acbc)
				return nil, _acbc
			}
			_fgdd.Function = append(_fgdd.Function, _dccgb)
		}
	} else {
		_agab, _dedg := _ebedg(_dabba)
		if _dedg != nil {
			_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dedg)
			return nil, _dedg
		}
		_fgdd.Function = append(_fgdd.Function, _agab)
	}
	if _aedb := _bfegc.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _aedb != nil {
		_aedb = _abf.TraceToDirectObject(_aedb)
		_dfff, _dfedc := _aedb.(*_abf.PdfObjectArray)
		if !_dfedc {
			_acd.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _aedb)
			return nil, _abf.ErrTypeError
		}
		if _dfff.Len() != 2 {
			_acd.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _dfff.Len())
			return nil, ErrInvalidAttribute
		}
		_fgdd.Extend = _dfff
	}
	return &_fgdd, nil
}

func (_affe fontCommon) asPdfObjectDictionary(_cgde string) *_abf.PdfObjectDictionary {
	if _cgde != "" && _affe._aacbc != "" && _cgde != _affe._aacbc {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061\u0073\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020O\u0076\u0065\u0072\u0072\u0069\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0062t\u0079\u0070\u0065\u0020\u0074\u006f \u0025\u0023\u0071 \u0025\u0073", _cgde, _affe)
	} else if _cgde == "" && _affe._aacbc == "" {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0061s\u0050\u0064\u0066Ob\u006a\u0065\u0063\u0074\u0044\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074y\u0070\u0065\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _affe)
	} else if _affe._aacbc == "" {
		_affe._aacbc = _cgde
	}
	_fdea := _abf.MakeDict()
	_fdea.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0046\u006f\u006e\u0074"))
	_fdea.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _abf.MakeName(_affe._ecggf))
	_fdea.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName(_affe._aacbc))
	if _affe._dcbaf != nil {
		_fdea.Set("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072", _affe._dcbaf.ToPdfObject())
	}
	if _affe._dabca != nil {
		_fdea.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _affe._dabca)
	} else if _affe._aabfe != nil {
		_geda, _eedf := _affe._aabfe.Stream()
		if _eedf != nil {
			_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0067\u0065\u0074\u0020C\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0065r\u0072\u003d\u0025\u0076", _eedf)
		} else {
			_fdea.Set("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e", _geda)
		}
	}
	return _fdea
}

// ImageToRGB converts ICCBased colorspace image to RGB and returns the result.
func (_degge *PdfColorspaceICCBased) ImageToRGB(img Image) (Image, error) {
	if _degge.Alternate == nil {
		_acd.Log.Debug("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		if _degge.N == 1 {
			_acd.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061y\u0020\u0028\u004e\u003d\u0031\u0029")
			_ecdc := NewPdfColorspaceDeviceGray()
			return _ecdc.ImageToRGB(img)
		} else if _degge.N == 3 {
			_acd.Log.Debug("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006eg\u0020\u0044\u0065\u0076\u0069\u0063e\u0052\u0047B\u0020\u0028N\u003d3\u0029")
			return img, nil
		} else if _degge.N == 4 {
			_acd.Log.Debug("\u0049\u0043\u0043\u0020\u0042a\u0073\u0065\u0064\u0020\u0063o\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061\u006c\u0074\u0065r\u006e\u0061\u0074\u0069\u0076\u0065\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059K\u0020\u0028\u004e\u003d\u0034\u0029")
			_bfbba := NewPdfColorspaceDeviceCMYK()
			return _bfbba.ImageToRGB(img)
		} else {
			return img, _fd.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	_acd.Log.Trace("\u0049\u0043\u0043 \u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0069\u0076\u0065\u003a\u0020\u0025\u0023\u0076", _degge)
	_cfbe, _baff := _degge.Alternate.ImageToRGB(img)
	_acd.Log.Trace("I\u0043C\u0020\u0049\u006e\u0070\u0075\u0074\u0020\u0069m\u0061\u0067\u0065\u003a %\u002b\u0076", img)
	_acd.Log.Trace("I\u0043\u0043\u0020\u004fut\u0070u\u0074\u0020\u0069\u006d\u0061g\u0065\u003a\u0020\u0025\u002b\u0076", _cfbe)
	return _cfbe, _baff
}

// NewCustomPdfOutputIntent creates a new custom PdfOutputIntent.
func NewCustomPdfOutputIntent(outputCondition, outputConditionIdentifier, info string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, Info: info, DestOutputProfile: destOutputProfile, _dcfb: _abf.MakeDict(), ColorComponents: colorComponents}
}

// PdfAnnotationSound represents Sound annotations.
// (Section 12.5.6.16).
type PdfAnnotationSound struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Sound _abf.PdfObject
	Name  _abf.PdfObject
}

// GetAsTilingPattern returns a tiling pattern. Check with IsTiling() prior to using this.
func (_gaedf *PdfPattern) GetAsTilingPattern() *PdfTilingPattern {
	return _gaedf._bgafe.(*PdfTilingPattern)
}
func _fcfeb() _f.Time { _gaabd.Lock(); defer _gaabd.Unlock(); return _edfdc }

// Encoder returns the font's text encoder.
func (_bfcgg pdfFontType0) Encoder() _cbb.TextEncoder { return _bfcgg._edeaf }

// GetObjectNums returns the object numbers of the PDF objects in the file
// Numbered objects are either indirect objects or stream objects.
// e.g. objNums := pdfReader.GetObjectNums()
// The underlying objects can then be accessed with
// pdfReader.GetIndirectObjectByNumber(objNums[0]) for the first available object.
func (_gaffd *PdfReader) GetObjectNums() []int { return _gaffd._bebc.GetObjectNums() }

func (_acgcc *fontFile) parseASCIIPart(_gfdb []byte) error {
	if len(_gfdb) < 2 || string(_gfdb[:2]) != "\u0025\u0021" {
		return _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0074a\u0072\u0074\u0020\u006f\u0066\u0020\u0041S\u0043\u0049\u0049\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_cdgf, _ddcgca, _egce := _cgefc(_gfdb)
	if _egce != nil {
		return _egce
	}
	_gcgdb := _ceeabe(_cdgf)
	_acgcc._gadc = _gcgdb["\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"]
	if _acgcc._gadc == "" {
		_acd.Log.Debug("\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0020\u0068a\u0073\u0020\u006e\u006f\u0020\u002f\u0046\u006f\u006e\u0074N\u0061\u006d\u0065")
	}
	if _ddcgca != "" {
		_daaa, _fegc := _becce(_ddcgca)
		if _fegc != nil {
			return _fegc
		}
		_dgbgb, _fegc := _cbb.NewCustomSimpleTextEncoder(_daaa, nil)
		if _fegc != nil {
			_acd.Log.Debug("\u0045\u0052\u0052\u004fR\u0020\u003a\u0055\u004e\u004b\u004e\u004f\u0057\u004e\u0020G\u004cY\u0050\u0048\u003a\u0020\u0065\u0072\u0072=\u0025\u0076", _fegc)
			return nil
		}
		_acgcc._eedb = _dgbgb
	}
	return nil
}

// PdfActionSubmitForm represents a submitForm action.
type PdfActionSubmitForm struct {
	*PdfAction
	F      *PdfFilespec
	Fields _abf.PdfObject
	Flags  _abf.PdfObject
}

func (_decdg *PdfWriter) writeAcroFormFields() error {
	if _decdg._bdgeb == nil {
		return nil
	}
	_acd.Log.Trace("\u0057r\u0069t\u0069\u006e\u0067\u0020\u0061c\u0072\u006f \u0066\u006f\u0072\u006d\u0073")
	_cfbaa := _decdg._bdgeb.ToPdfObject()
	_acd.Log.Trace("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _cfbaa)
	_decdg._ddffc.Set("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d", _cfbaa)
	_cfcbg := _decdg.addObjects(_cfbaa)
	if _cfcbg != nil {
		return _cfcbg
	}
	return nil
}

// GetCatalogStructTreeRoot gets the catalog StructTreeRoot object.
func (_deeea *PdfReader) GetCatalogStructTreeRoot() (_abf.PdfObject, bool) {
	if _deeea._dagde == nil {
		return nil, false
	}
	_ddbf := _deeea._dagde.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074")
	return _ddbf, _ddbf != nil
}

// ToPdfObject implements interface PdfModel.
func (_ec *PdfActionHide) ToPdfObject() _abf.PdfObject {
	_ec.PdfAction.ToPdfObject()
	_bcg := _ec._egg
	_ede := _bcg.PdfObject.(*_abf.PdfObjectDictionary)
	_ede.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeHide)))
	_ede.SetIfNotNil("\u0054", _ec.T)
	_ede.SetIfNotNil("\u0048", _ec.H)
	return _bcg
}

// GetRevisionNumber returns the version of the current Pdf document
func (_fecb *PdfReader) GetRevisionNumber() int { return _fecb._bebc.GetRevisionNumber() }

// NewPdfOutputIntentFromPdfObject creates a new PdfOutputIntent from the input core.PdfObject.
func NewPdfOutputIntentFromPdfObject(object _abf.PdfObject) (*PdfOutputIntent, error) {
	_cgbgg := &PdfOutputIntent{}
	if _bffdgb := _cgbgg.ParsePdfObject(object); _bffdgb != nil {
		return nil, _bffdgb
	}
	return _cgbgg, nil
}

// ToPdfObject implements interface PdfModel.
func (_fgbf *PdfAnnotationPopup) ToPdfObject() _abf.PdfObject {
	_fgbf.PdfAnnotation.ToPdfObject()
	_fcef := _fgbf._dbc
	_gfdcd := _fcef.PdfObject.(*_abf.PdfObjectDictionary)
	_gfdcd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0050\u006f\u0070u\u0070"))
	_gfdcd.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _fgbf.Parent)
	_gfdcd.SetIfNotNil("\u004f\u0070\u0065\u006e", _fgbf.Open)
	return _fcef
}

// SetImage updates XObject Image with new image data.
func (_ecagc *XObjectImage) SetImage(img *Image, cs PdfColorspace) error {
	_ecagc.Filter.UpdateParams(img.GetParamsDict())
	_fgdcg, _cgdgf := _ecagc.Filter.EncodeBytes(img.Data)
	if _cgdgf != nil {
		return _cgdgf
	}
	_ecagc.Stream = _fgdcg
	_eeef := img.Width
	_ecagc.Width = &_eeef
	_ggfce := img.Height
	_ecagc.Height = &_ggfce
	_cbfaf := img.BitsPerComponent
	_ecagc.BitsPerComponent = &_cbfaf
	if cs == nil {
		if img.ColorComponents == 1 {
			_ecagc.ColorSpace = NewPdfColorspaceDeviceGray()
		} else if img.ColorComponents == 3 {
			_ecagc.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			_ecagc.ColorSpace = NewPdfColorspaceDeviceCMYK()
		} else {
			return _fd.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_ecagc.ColorSpace = cs
	}
	return nil
}

// NewPdfColorspaceSpecialPattern returns a new pattern color.
func NewPdfColorspaceSpecialPattern() *PdfColorspaceSpecialPattern {
	return &PdfColorspaceSpecialPattern{}
}

// NewPdfColorspaceDeviceGray returns a new grayscale colorspace.
func NewPdfColorspaceDeviceGray() *PdfColorspaceDeviceGray { return &PdfColorspaceDeviceGray{} }

// SetContext sets the specific fielddata type, e.g. would be PdfFieldButton for a button field.
func (_caeg *PdfField) SetContext(ctx PdfModel) { _caeg._ffea = ctx }

// WatermarkImageOptions contains options for configuring the watermark process.
type WatermarkImageOptions struct {
	Alpha               float64
	FitToWidth          bool
	PreserveAspectRatio bool
}

// NewPdfActionURI returns a new "Uri" action.
func NewPdfActionURI() *PdfActionURI {
	_adb := NewPdfAction()
	_gagf := &PdfActionURI{}
	_gagf.PdfAction = _adb
	_adb.SetContext(_gagf)
	return _gagf
}

const (
	BorderStyleSolid     BorderStyle = iota
	BorderStyleDashed    BorderStyle = iota
	BorderStyleBeveled   BorderStyle = iota
	BorderStyleInset     BorderStyle = iota
	BorderStyleUnderline BorderStyle = iota
)

func (_gcc *PdfReader) newPdfAnnotationPolyLineFromDict(_cfaag *_abf.PdfObjectDictionary) (*PdfAnnotationPolyLine, error) {
	_gdg := PdfAnnotationPolyLine{}
	_geef, _ggcaf := _gcc.newPdfAnnotationMarkupFromDict(_cfaag)
	if _ggcaf != nil {
		return nil, _ggcaf
	}
	_gdg.PdfAnnotationMarkup = _geef
	_gdg.Vertices = _cfaag.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_gdg.LE = _cfaag.Get("\u004c\u0045")
	_gdg.BS = _cfaag.Get("\u0042\u0053")
	_gdg.IC = _cfaag.Get("\u0049\u0043")
	_gdg.BE = _cfaag.Get("\u0042\u0045")
	_gdg.IT = _cfaag.Get("\u0049\u0054")
	_gdg.Measure = _cfaag.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_gdg, nil
}

// Items returns all children outline items.
func (_cbcgb *Outline) Items() []*OutlineItem { return _cbcgb.Entries }

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
func (_cfff *Image) Resample(targetBitsPerComponent int64) {
	if _cfff.BitsPerComponent == targetBitsPerComponent {
		return
	}
	_fdbce := _cfff.GetSamples()
	if targetBitsPerComponent < _cfff.BitsPerComponent {
		_cbdbc := _cfff.BitsPerComponent - targetBitsPerComponent
		for _ddebe := range _fdbce {
			_fdbce[_ddebe] >>= uint(_cbdbc)
		}
	} else if targetBitsPerComponent > _cfff.BitsPerComponent {
		_decg := targetBitsPerComponent - _cfff.BitsPerComponent
		for _affg := range _fdbce {
			_fdbce[_affg] <<= uint(_decg)
		}
	}
	_cfff.BitsPerComponent = targetBitsPerComponent
	if _cfff.BitsPerComponent < 8 {
		_cfff.resampleLowBits(_fdbce)
		return
	}
	_egcd := _gca.BytesPerLine(int(_cfff.Width), int(_cfff.BitsPerComponent), _cfff.ColorComponents)
	_bfaa := make([]byte, _egcd*int(_cfff.Height))
	var (
		_ggfb, _daacf, _afee, _dfccd int
		_ddcgf                       uint32
	)
	for _afee = 0; _afee < int(_cfff.Height); _afee++ {
		_ggfb = _afee * _egcd
		_daacf = (_afee+1)*_egcd - 1
		_aeffb := _gf.ResampleUint32(_fdbce[_ggfb:_daacf], int(targetBitsPerComponent), 8)
		for _dfccd, _ddcgf = range _aeffb {
			_bfaa[_dfccd+_ggfb] = byte(_ddcgf)
		}
	}
	_cfff.Data = _bfaa
}

// GetAsShadingPattern returns a shading pattern. Check with IsShading() prior to using this.
func (_gdaf *PdfPattern) GetAsShadingPattern() *PdfShadingPattern {
	return _gdaf._bgafe.(*PdfShadingPattern)
}

// GetContentStreams returns the content stream as an array of strings.
func (_ffedg *PdfPage) GetContentStreams() ([]string, error) {
	_fabbce := _ffedg.GetContentStreamObjs()
	var _cedd []string
	for _, _feee := range _fabbce {
		_eefc, _gfdbd := _eeggg(_feee)
		if _gfdbd != nil {
			return nil, _gfdbd
		}
		_cedd = append(_cedd, _eefc)
	}
	return _cedd, nil
}

func (_eebg *PdfReader) newPdfAnnotationPrinterMarkFromDict(_gdb *_abf.PdfObjectDictionary) (*PdfAnnotationPrinterMark, error) {
	_cef := PdfAnnotationPrinterMark{}
	_cef.MN = _gdb.Get("\u004d\u004e")
	return &_cef, nil
}

// PdfFieldButton represents a button field which includes push buttons, checkboxes, and radio buttons.
type PdfFieldButton struct {
	*PdfField
	Opt   *_abf.PdfObjectArray
	_ccdd *Image
}

// ToPdfObject converts date to a PDF string object.
func (_aegdg *PdfDate) ToPdfObject() _abf.PdfObject {
	_acfd := _e.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _aegdg._fabd, _aegdg._fcdacf, _aegdg._gecdc, _aegdg._ebda, _aegdg._efba, _aegdg._fgddf, _aegdg._aggabc, _aegdg._dbgccd, _aegdg._ccfca)
	return _abf.MakeString(_acfd)
}

func (_acfbc *PdfAcroForm) fill(_geec FieldValueProvider, _bdfde FieldAppearanceGenerator) error {
	if _acfbc == nil {
		return nil
	}
	_gfecbg, _cdeb := _geec.FieldValues()
	if _cdeb != nil {
		return _cdeb
	}
	for _, _geceag := range _acfbc.AllFields() {
		_bbdf := _geceag.PartialName()
		_faec, _eegae := _gfecbg[_bbdf]
		if !_eegae {
			if _bgfca, _efafc := _geceag.FullName(); _efafc == nil {
				_faec, _eegae = _gfecbg[_bgfca]
			}
		}
		if !_eegae {
			_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _bbdf)
			continue
		}
		if _gbae := _bage(_geceag, _faec); _gbae != nil {
			return _gbae
		}
		if _bdfde == nil {
			continue
		}
		for _, _bdfag := range _geceag.Annotations {
			_egbg, _acddc := _bdfde.GenerateAppearanceDict(_acfbc, _geceag, _bdfag)
			if _acddc != nil {
				return _acddc
			}
			_bdfag.AP = _egbg
			_bdfag.ToPdfObject()
		}
	}
	return nil
}

// DefaultFont returns the default font, which is currently the built in Helvetica.
func DefaultFont() *PdfFont {
	_ffff, _dabfbe := _gbe.NewStdFontByName(HelveticaName)
	if !_dabfbe {
		panic("\u0048\u0065lv\u0065\u0074\u0069c\u0061\u0020\u0073\u0068oul\u0064 a\u006c\u0077\u0061\u0079\u0073\u0020\u0062e \u0061\u0076\u0061\u0069\u006c\u0061\u0062l\u0065")
	}
	_fgag := _bcee(_ffff)
	return &PdfFont{_gedca: &_fgag}
}

// CustomKeys returns all custom info keys as list.
func (_cbca *PdfInfo) CustomKeys() []string {
	if _cbca._cbf == nil {
		return nil
	}
	_gbcf := make([]string, len(_cbca._cbf.Keys()))
	for _, _eagfa := range _cbca._cbf.Keys() {
		_gbcf = append(_gbcf, _eagfa.String())
	}
	return _gbcf
}

const (
	_becb  = 0x00001
	_aabab = 0x00002
	_eceag = 0x00004
	_afde  = 0x00008
	_bbadf = 0x00020
	_bacb  = 0x00040
	_bbbee = 0x10000
	_dbff  = 0x20000
	_geba  = 0x40000
)

func _gedcb(_ebcff *_abf.PdfObjectDictionary) {
	_aafd, _ccbdf := _abf.GetArray(_ebcff.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
	_egefb, _cfdf := _abf.GetIntVal(_ebcff.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
	_fbfa, _ggafd := _abf.GetIntVal(_ebcff.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
	if _ccbdf && _cfdf && _ggafd {
		_cbfc := _aafd.Len()
		if _cbfc != _fbfa-_egefb+1 {
			_acd.Log.Debug("\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076\u002c\u0020\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u003a\u0020\u0025\u0076", _cbfc, _fbfa)
			_dfbcg := _abf.PdfObjectInteger(_egefb + _cbfc - 1)
			_ebcff.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", &_dfbcg)
		}
	}
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_gefea pdfCIDFontType2) GetCharMetrics(code _cbb.CharCode) (_gbe.CharMetrics, bool) {
	if _fbaff, _edbda := _gefea._ddeea[code]; _edbda {
		return _gbe.CharMetrics{Wx: _fbaff}, true
	}
	_gbabf := rune(code)
	_ebede, _edfec := _gefea._dffcb[_gbabf]
	if !_edfec {
		_ebede = int(_gefea._cecdg)
	}
	return _gbe.CharMetrics{Wx: float64(_ebede)}, true
}

// GetNumComponents returns the number of color components (1 for grayscale).
func (_fafcf *PdfColorDeviceGray) GetNumComponents() int { return 1 }

func (_eef *PdfReader) newPdfAnnotationFromIndirectObject(_gce *_abf.PdfIndirectObject) (*PdfAnnotation, error) {
	_ade, _dccdf := _gce.PdfObject.(*_abf.PdfObjectDictionary)
	if !_dccdf {
		return nil, _e.Errorf("\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _eef._ceecd.GetModelFromPrimitive(_ade); model != nil {
		_abc, _abfe := model.(*PdfAnnotation)
		if !_abfe {
			return nil, _e.Errorf("\u0063\u0061\u0063\u0068\u0065\u0064 \u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050D\u0046\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e")
		}
		return _abc, nil
	}
	_gfbf := &PdfAnnotation{}
	_gfbf._dbc = _gce
	_eef._ceecd.Register(_ade, _gfbf)
	if _bag := _ade.Get("\u0054\u0079\u0070\u0065"); _bag != nil {
		_cgae, _ceae := _bag.(*_abf.PdfObjectName)
		if !_ceae {
			_acd.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _bag)
		} else {
			if *_cgae != "\u0041\u006e\u006eo\u0074" {
				_acd.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070\u0065\u0063\u0074\u0065d\u0020\u0054\u0079\u0070\u0065\u0020\u0021=\u0020\u0041\u006e\u006e\u006f\u0074\u0020\u0028\u0025\u0073\u0029", *_cgae)
			}
		}
	}
	if _fdb := _ade.Get("\u0052\u0065\u0063\u0074"); _fdb != nil {
		_gfbf.Rect = _fdb
	}
	if _edc := _ade.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _edc != nil {
		_gfbf.Contents = _edc
	}
	if _gffc := _ade.Get("\u0050"); _gffc != nil {
		_gfbf.P = _gffc
	}
	if _ebdcb := _ade.Get("\u004e\u004d"); _ebdcb != nil {
		_gfbf.NM = _ebdcb
	}
	if _aaef := _ade.Get("\u004d"); _aaef != nil {
		_gfbf.M = _aaef
	}
	if _bgce := _ade.Get("\u0046"); _bgce != nil {
		_gfbf.F = _bgce
	}
	if _ded := _ade.Get("\u0041\u0050"); _ded != nil {
		_gfbf.AP = _ded
	}
	if _ddb := _ade.Get("\u0041\u0053"); _ddb != nil {
		_gfbf.AS = _ddb
	}
	if _defa := _ade.Get("\u0042\u006f\u0072\u0064\u0065\u0072"); _defa != nil {
		_gfbf.Border = _defa
	}
	if _aedd := _ade.Get("\u0043"); _aedd != nil {
		_gfbf.C = _aedd
	}
	if _dfaf := _ade.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074"); _dfaf != nil {
		_gfbf.StructParent = _dfaf
	}
	if _aceg := _ade.Get("\u004f\u0043"); _aceg != nil {
		_gfbf.OC = _aceg
	}
	_cda := _ade.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")
	if _cda == nil {
		_acd.Log.Debug("\u0057\u0041\u0052\u004e\u0049\u004e\u0047:\u0020\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079 \u0069s\u0073\u0075\u0065\u0020\u002d\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073u\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0073\u0075\u0062\u0074\u0079p\u0065")
		_gfbf._edg = nil
		return _gfbf, nil
	}
	_acgc, _fgg := _cda.(*_abf.PdfObjectName)
	if !_fgg {
		_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065 !\u003d\u0020n\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _cda)
		return nil, _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d n\u0061\u006d\u0065 \u0028%\u0054\u0029", _cda)
	}
	switch *_acgc {
	case "\u0054\u0065\u0078\u0074":
		_cec, _dcec := _eef.newPdfAnnotationTextFromDict(_ade)
		if _dcec != nil {
			return nil, _dcec
		}
		_cec.PdfAnnotation = _gfbf
		_gfbf._edg = _cec
		return _gfbf, nil
	case "\u004c\u0069\u006e\u006b":
		_dceb, _eaba := _eef.newPdfAnnotationLinkFromDict(_ade)
		if _eaba != nil {
			return nil, _eaba
		}
		_dceb.PdfAnnotation = _gfbf
		_gfbf._edg = _dceb
		return _gfbf, nil
	case "\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074":
		_dfgc, _gbbd := _eef.newPdfAnnotationFreeTextFromDict(_ade)
		if _gbbd != nil {
			return nil, _gbbd
		}
		_dfgc.PdfAnnotation = _gfbf
		_gfbf._edg = _dfgc
		return _gfbf, nil
	case "\u004c\u0069\u006e\u0065":
		_dag, _daa := _eef.newPdfAnnotationLineFromDict(_ade)
		if _daa != nil {
			return nil, _daa
		}
		_dag.PdfAnnotation = _gfbf
		_gfbf._edg = _dag
		_acd.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041N\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0061\u006e\u006eo\u0074\u0020\u0028\u0025\u0054\u0029\u003a \u0025\u002b\u0076\u000a", _gfbf, _gfbf)
		_acd.Log.Trace("\u004c\u0049\u004eE\u0020\u0041\u004e\u004eO\u0054\u0041\u0054\u0049\u004f\u004e\u003a \u0063\u0074\u0078\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u002b\u0076\u000a", _dag, _dag)
		_acd.Log.Trace("\u004c\u0049\u004e\u0045\u0020\u0041\u004e\u004e\u004f\u0054\u0041\u0054\u0049\u004f\u004e\u0020\u004d\u0061\u0072\u006b\u0075\u0070\u003a\u0020c\u0074\u0078\u0020\u0028\u0025T\u0029\u003a \u0025\u002b\u0076\u000a", _dag.PdfAnnotationMarkup, _dag.PdfAnnotationMarkup)
		return _gfbf, nil
	case "\u0053\u0071\u0075\u0061\u0072\u0065":
		_gee, _fcc := _eef.newPdfAnnotationSquareFromDict(_ade)
		if _fcc != nil {
			return nil, _fcc
		}
		_gee.PdfAnnotation = _gfbf
		_gfbf._edg = _gee
		return _gfbf, nil
	case "\u0043\u0069\u0072\u0063\u006c\u0065":
		_gead, _efgg := _eef.newPdfAnnotationCircleFromDict(_ade)
		if _efgg != nil {
			return nil, _efgg
		}
		_gead.PdfAnnotation = _gfbf
		_gfbf._edg = _gead
		return _gfbf, nil
	case "\u0050o\u006c\u0079\u0067\u006f\u006e":
		_ccff, _gfgg := _eef.newPdfAnnotationPolygonFromDict(_ade)
		if _gfgg != nil {
			return nil, _gfgg
		}
		_ccff.PdfAnnotation = _gfbf
		_gfbf._edg = _ccff
		return _gfbf, nil
	case "\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065":
		_feg, _cfaaf := _eef.newPdfAnnotationPolyLineFromDict(_ade)
		if _cfaaf != nil {
			return nil, _cfaaf
		}
		_feg.PdfAnnotation = _gfbf
		_gfbf._edg = _feg
		return _gfbf, nil
	case "\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t":
		_bgfa, _aaff := _eef.newPdfAnnotationHighlightFromDict(_ade)
		if _aaff != nil {
			return nil, _aaff
		}
		_bgfa.PdfAnnotation = _gfbf
		_gfbf._edg = _bgfa
		return _gfbf, nil
	case "\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee":
		_cddf, _ddaf := _eef.newPdfAnnotationUnderlineFromDict(_ade)
		if _ddaf != nil {
			return nil, _ddaf
		}
		_cddf.PdfAnnotation = _gfbf
		_gfbf._edg = _cddf
		return _gfbf, nil
	case "\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079":
		_aefe, _efb := _eef.newPdfAnnotationSquigglyFromDict(_ade)
		if _efb != nil {
			return nil, _efb
		}
		_aefe.PdfAnnotation = _gfbf
		_gfbf._edg = _aefe
		return _gfbf, nil
	case "\u0053t\u0072\u0069\u006b\u0065\u004f\u0075t":
		_ebbd, _gacg := _eef.newPdfAnnotationStrikeOut(_ade)
		if _gacg != nil {
			return nil, _gacg
		}
		_ebbd.PdfAnnotation = _gfbf
		_gfbf._edg = _ebbd
		return _gfbf, nil
	case "\u0043\u0061\u0072e\u0074":
		_agec, _aacb := _eef.newPdfAnnotationCaretFromDict(_ade)
		if _aacb != nil {
			return nil, _aacb
		}
		_agec.PdfAnnotation = _gfbf
		_gfbf._edg = _agec
		return _gfbf, nil
	case "\u0053\u0074\u0061m\u0070":
		_acda, _dga := _eef.newPdfAnnotationStampFromDict(_ade)
		if _dga != nil {
			return nil, _dga
		}
		_acda.PdfAnnotation = _gfbf
		_gfbf._edg = _acda
		return _gfbf, nil
	case "\u0049\u006e\u006b":
		_ffaf, _caee := _eef.newPdfAnnotationInkFromDict(_ade)
		if _caee != nil {
			return nil, _caee
		}
		_ffaf.PdfAnnotation = _gfbf
		_gfbf._edg = _ffaf
		return _gfbf, nil
	case "\u0050\u006f\u0070u\u0070":
		_faff, _gagc := _eef.newPdfAnnotationPopupFromDict(_ade)
		if _gagc != nil {
			return nil, _gagc
		}
		_faff.PdfAnnotation = _gfbf
		_gfbf._edg = _faff
		return _gfbf, nil
	case "\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074":
		_bdaa, _geea := _eef.newPdfAnnotationFileAttachmentFromDict(_ade)
		if _geea != nil {
			return nil, _geea
		}
		_bdaa.PdfAnnotation = _gfbf
		_gfbf._edg = _bdaa
		return _gfbf, nil
	case "\u0053\u006f\u0075n\u0064":
		_dge, _aece := _eef.newPdfAnnotationSoundFromDict(_ade)
		if _aece != nil {
			return nil, _aece
		}
		_dge.PdfAnnotation = _gfbf
		_gfbf._edg = _dge
		return _gfbf, nil
	case "\u0052i\u0063\u0068\u004d\u0065\u0064\u0069a":
		_dccf, _dgg := _eef.newPdfAnnotationRichMediaFromDict(_ade)
		if _dgg != nil {
			return nil, _dgg
		}
		_dccf.PdfAnnotation = _gfbf
		_gfbf._edg = _dccf
		return _gfbf, nil
	case "\u004d\u006f\u0076i\u0065":
		_edeg, _fbac := _eef.newPdfAnnotationMovieFromDict(_ade)
		if _fbac != nil {
			return nil, _fbac
		}
		_edeg.PdfAnnotation = _gfbf
		_gfbf._edg = _edeg
		return _gfbf, nil
	case "\u0053\u0063\u0072\u0065\u0065\u006e":
		_fdaa, _fbge := _eef.newPdfAnnotationScreenFromDict(_ade)
		if _fbge != nil {
			return nil, _fbge
		}
		_fdaa.PdfAnnotation = _gfbf
		_gfbf._edg = _fdaa
		return _gfbf, nil
	case "\u0057\u0069\u0064\u0067\u0065\u0074":
		_ddbd, _bacd := _eef.newPdfAnnotationWidgetFromDict(_ade)
		if _bacd != nil {
			return nil, _bacd
		}
		_ddbd.PdfAnnotation = _gfbf
		_gfbf._edg = _ddbd
		return _gfbf, nil
	case "P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b":
		_cdf, _cgg := _eef.newPdfAnnotationPrinterMarkFromDict(_ade)
		if _cgg != nil {
			return nil, _cgg
		}
		_cdf.PdfAnnotation = _gfbf
		_gfbf._edg = _cdf
		return _gfbf, nil
	case "\u0054r\u0061\u0070\u004e\u0065\u0074":
		_gfgf, _fbgec := _eef.newPdfAnnotationTrapNetFromDict(_ade)
		if _fbgec != nil {
			return nil, _fbgec
		}
		_gfgf.PdfAnnotation = _gfbf
		_gfbf._edg = _gfgf
		return _gfbf, nil
	case "\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k":
		_bcfc, _cce := _eef.newPdfAnnotationWatermarkFromDict(_ade)
		if _cce != nil {
			return nil, _cce
		}
		_bcfc.PdfAnnotation = _gfbf
		_gfbf._edg = _bcfc
		return _gfbf, nil
	case "\u0033\u0044":
		_aeddg, _fafc := _eef.newPdfAnnotation3DFromDict(_ade)
		if _fafc != nil {
			return nil, _fafc
		}
		_aeddg.PdfAnnotation = _gfbf
		_gfbf._edg = _aeddg
		return _gfbf, nil
	case "\u0050\u0072\u006f\u006a\u0065\u0063\u0074\u0069\u006f\u006e":
		_ceaea, _acad := _eef.newPdfAnnotationProjectionFromDict(_ade)
		if _acad != nil {
			return nil, _acad
		}
		_ceaea.PdfAnnotation = _gfbf
		_gfbf._edg = _ceaea
		return _gfbf, nil
	case "\u0052\u0065\u0064\u0061\u0063\u0074":
		_fgeb, _efd := _eef.newPdfAnnotationRedactFromDict(_ade)
		if _efd != nil {
			return nil, _efd
		}
		_fgeb.PdfAnnotation = _gfbf
		_gfbf._edg = _fgeb
		return _gfbf, nil
	}
	_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069\u006fn\u003a\u0020\u0025\u0073", *_acgc)
	return nil, nil
}

// NewCompositePdfFontFromTTF loads a composite TTF font. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTF.
func NewCompositePdfFontFromTTF(r _gc.ReadSeeker) (*PdfFont, error) {
	_edfda, _bcba := _gc.ReadAll(r)
	if _bcba != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _bcba)
		return nil, _bcba
	}
	_efag, _bcba := _gbe.TtfParse(_dd.NewReader(_edfda))
	if _bcba != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067 \u0074\u0074\u0066\u0020\u0066\u006f\u006et\u003a\u0020\u0025\u0076", _bcba)
		return nil, _bcba
	}
	_efaae := &pdfCIDFontType2{fontCommon: fontCommon{_aacbc: "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032"}, CIDToGIDMap: _abf.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079")}
	if len(_efag.Widths) <= 0 {
		return nil, _fd.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_afce := 1000.0 / float64(_efag.UnitsPerEm)
	_dfbe := _afce * float64(_efag.Widths[0])
	_egfaa := make(map[rune]int)
	_faace := make(map[_gbe.GID]int)
	_fecae := _gbe.GID(len(_efag.Widths))
	for _gcbca, _adaac := range _efag.Chars {
		if _adaac > _fecae-1 {
			continue
		}
		_ecfe := int(_afce * float64(_efag.Widths[_adaac]))
		_egfaa[_gcbca] = _ecfe
		_faace[_adaac] = _ecfe
	}
	_efaae._dffcb = _egfaa
	_efaae.DW = _abf.MakeInteger(int64(_dfbe))
	_gceee := _aabg(_faace, uint16(_fecae))
	_efaae.W = _abf.MakeIndirectObject(_gceee)
	_egagf := _abf.MakeDict()
	_egagf.Set("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067", _abf.MakeString("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"))
	_egagf.Set("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079", _abf.MakeString("\u0041\u0064\u006fb\u0065"))
	_egagf.Set("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074", _abf.MakeInteger(0))
	_efaae.CIDSystemInfo = _egagf
	_cedcd := &PdfFontDescriptor{FontName: _abf.MakeName(_efag.PostScriptName), Ascent: _abf.MakeFloat(_afce * float64(_efag.TypoAscender)), Descent: _abf.MakeFloat(_afce * float64(_efag.TypoDescender)), CapHeight: _abf.MakeFloat(_afce * float64(_efag.CapHeight)), FontBBox: _abf.MakeArrayFromFloats([]float64{_afce * float64(_efag.Xmin), _afce * float64(_efag.Ymin), _afce * float64(_efag.Xmax), _afce * float64(_efag.Ymax)}), ItalicAngle: _abf.MakeFloat(_efag.ItalicAngle), MissingWidth: _abf.MakeFloat(_dfbe)}
	_bfbgb, _bcba := _abf.MakeStream(_edfda, _abf.NewFlateEncoder())
	if _bcba != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _bcba)
		return nil, _bcba
	}
	_bfbgb.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _abf.MakeInteger(int64(len(_edfda))))
	_cedcd.FontFile2 = _bfbgb
	if _efag.Bold {
		_cedcd.StemV = _abf.MakeInteger(120)
	} else {
		_cedcd.StemV = _abf.MakeInteger(70)
	}
	_fggaa := _eceag
	if _efag.IsFixedPitch {
		_fggaa |= _becb
	}
	if _efag.ItalicAngle != 0 {
		_fggaa |= _bacb
	}
	_cedcd.Flags = _abf.MakeInteger(int64(_fggaa))
	_efaae._ecggf = _efag.PostScriptName
	_efaae._dcbaf = _cedcd
	_edegc := pdfFontType0{fontCommon: fontCommon{_aacbc: "\u0054\u0079\u0070e\u0030", _ecggf: _efag.PostScriptName}, DescendantFont: &PdfFont{_gedca: _efaae}, Encoding: _abf.MakeName("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048"), _edeaf: _efag.NewEncoder()}
	if len(_efag.Chars) > 0 {
		_gbefd := make(map[_bd.CharCode]rune, len(_efag.Chars))
		for _ccbbe, _dcbff := range _efag.Chars {
			_ecafb := _bd.CharCode(_dcbff)
			if _fgedc, _gaeee := _gbefd[_ecafb]; !_gaeee || (_gaeee && _fgedc > _ccbbe) {
				_gbefd[_ecafb] = _ccbbe
			}
		}
		_edegc._aabfe = _bd.NewToUnicodeCMap(_gbefd)
	}
	_ccbf := PdfFont{_gedca: &_edegc}
	return &_ccbf, nil
}

// NewPdfColorPattern returns an empty color pattern.
func NewPdfColorPattern() *PdfColorPattern { _ecee := &PdfColorPattern{}; return _ecee }

// NewPdfColorspaceICCBased returns a new ICCBased colorspace object.
func NewPdfColorspaceICCBased(N int) (*PdfColorspaceICCBased, error) {
	_dgbc := &PdfColorspaceICCBased{}
	if N != 1 && N != 3 && N != 4 {
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u0031/\u0033\u002f\u0034\u0029")
	}
	_dgbc.N = N
	return _dgbc, nil
}

// K returns the value of the key component of the color.
func (_edaac *PdfColorDeviceCMYK) K() float64 { return _edaac[3] }

// NewPdfActionSound returns a new "sound" action.
func NewPdfActionSound() *PdfActionSound {
	_bed := NewPdfAction()
	_gef := &PdfActionSound{}
	_gef.PdfAction = _bed
	_bed.SetContext(_gef)
	return _gef
}

// A PdfPattern can represent a Pattern, either a tiling pattern or a shading pattern.
// Note that all patterns shall be treated as colours; a Pattern colour space shall be established with the CS or cs
// operator just like other colour spaces, and a particular pattern shall be installed as the current colour with the
// SCN or scn operator.
type PdfPattern struct {
	// Type: Pattern
	PatternType int64
	_bgafe      PdfModel
	_bcfca      _abf.PdfObject
}

// GetVersion gets the document version.
func (_ffagcg *PdfWriter) GetVersion() _abf.Version { return _ffagcg._ecfa }

func (_dgfa *PdfPage) getParentResources() (*PdfPageResources, error) {
	_bfedd := _dgfa.Parent
	for _bfedd != nil {
		_edeaa, _cfcg := _abf.GetDict(_bfedd)
		if !_cfcg {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020n\u006f\u0064\u0065")
			return nil, _fd.New("i\u006e\u0076\u0061\u006cid\u0020p\u0061\u0072\u0065\u006e\u0074 \u006f\u0062\u006a\u0065\u0063\u0074")
		}
		if _bgeaa := _edeaa.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _bgeaa != nil {
			_cfab, _afacgg := _abf.GetDict(_bgeaa)
			if !_afacgg {
				return nil, _fd.New("i\u006e\u0076\u0061\u006cid\u0020r\u0065\u0073\u006f\u0075\u0072c\u0065\u0020\u0064\u0069\u0063\u0074")
			}
			_bebea, _dbcef := NewPdfPageResourcesFromDict(_cfab)
			if _dbcef != nil {
				return nil, _dbcef
			}
			return _bebea, nil
		}
		_bfedd = _edeaa.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, nil
}

// ToWriter creates a new writer from the current reader, based on the specified options.
// If no options are provided, all reader properties are copied to the writer.
func (_fade *PdfReader) ToWriter(opts *ReaderToWriterOpts) (*PdfWriter, error) {
	_aade := NewPdfWriter()
	if opts == nil {
		opts = &ReaderToWriterOpts{}
	}
	_gfead, _effcc := _fade.GetNumPages()
	if _effcc != nil {
		_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _effcc)
		return nil, _effcc
	}
	for _abcced := 1; _abcced <= _gfead; _abcced++ {
		_egfg, _gagec := _fade.GetPage(_abcced)
		if _gagec != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gagec)
			return nil, _gagec
		}
		if opts.PageProcessCallback != nil {
			_gagec = opts.PageProcessCallback(_abcced, _egfg)
			if _gagec != nil {
				_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gagec)
				return nil, _gagec
			}
		} else if opts.PageCallback != nil {
			opts.PageCallback(_abcced, _egfg)
		}
		_gagec = _aade.AddPage(_egfg)
		if _gagec != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gagec)
			return nil, _gagec
		}
	}
	_aade._ecfa = _fade.PdfVersion()
	if !opts.SkipInfo {
		_fgdef, _cfgfg := _fade.GetPdfInfo()
		if _cfgfg != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cfgfg)
		} else {
			_aade._ddegc.PdfObject = _fgdef.ToPdfObject()
		}
	}
	if !opts.SkipMetadata {
		if _eecce := _fade._dagde.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _eecce != nil {
			if _bfgbf := _aade.SetCatalogMetadata(_eecce); _bfgbf != nil {
				return nil, _bfgbf
			}
		}
	}
	if !opts.SkipAcroForm {
		_acabdd := _aade.SetForms(_fade.AcroForm)
		if _acabdd != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _acabdd)
			return nil, _acabdd
		}
	}
	if !opts.SkipOutlines {
		_aade.AddOutlineTree(_fade.GetOutlineTree())
	}
	if !opts.SkipOCProperties {
		_aadge, _dcbaec := _fade.GetOCProperties()
		if _dcbaec != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dcbaec)
		} else {
			_dcbaec = _aade.SetOCProperties(_aadge)
			if _dcbaec != nil {
				_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dcbaec)
			}
		}
	}
	if !opts.SkipPageLabels {
		_fgabe, _afbaf := _fade.GetPageLabels()
		if _afbaf != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _afbaf)
		} else {
			_afbaf = _aade.SetPageLabels(_fgabe)
			if _afbaf != nil {
				_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _afbaf)
			}
		}
	}
	if !opts.SkipNamedDests {
		_eacbd, _fabc := _fade.GetNamedDestinations()
		if _fabc != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fabc)
		} else {
			_fabc = _aade.SetNamedDestinations(_eacbd)
			if _fabc != nil {
				_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fabc)
			}
		}
	}
	if !opts.SkipNameDictionary {
		_dcgdf, _gfbfb := _fade.GetNameDictionary()
		if _gfbfb != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gfbfb)
		} else {
			_gfbfb = _aade.SetNameDictionary(_dcgdf)
			if _gfbfb != nil {
				_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gfbfb)
			}
		}
	}
	if !opts.SkipRotation && _fade.Rotate != nil {
		if _acgfa := _aade.SetRotation(*_fade.Rotate); _acgfa != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _acgfa)
		}
	}
	return &_aade, nil
}

// FontDescriptor returns font's PdfFontDescriptor. This may be a builtin descriptor for standard 14
// fonts but must be an explicit descriptor for other fonts.
func (_feda *PdfFont) FontDescriptor() *PdfFontDescriptor {
	if _feda.baseFields()._dcbaf != nil {
		return _feda.baseFields()._dcbaf
	}
	if _bafb := _feda._gedca.getFontDescriptor(); _bafb != nil {
		return _bafb
	}
	_acd.Log.Error("\u0041\u006cl \u0066\u006f\u006et\u0073\u0020\u0068\u0061ve \u0061 D\u0065\u0073\u0063\u0072\u0069\u0070\u0074or\u002e\u0020\u0066\u006f\u006e\u0074\u003d%\u0073", _feda)
	return nil
}

func (_gedc *PdfReader) newPdfAnnotationSquigglyFromDict(_feca *_abf.PdfObjectDictionary) (*PdfAnnotationSquiggly, error) {
	_bfef := PdfAnnotationSquiggly{}
	_baae, _aebb := _gedc.newPdfAnnotationMarkupFromDict(_feca)
	if _aebb != nil {
		return nil, _aebb
	}
	_bfef.PdfAnnotationMarkup = _baae
	_bfef.QuadPoints = _feca.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_bfef, nil
}
func _bdfef() string { _gaabd.Lock(); defer _gaabd.Unlock(); return _efdg }
func (_faeaf *PdfColorspaceSpecialPattern) String() string {
	return "\u0050a\u0074\u0074\u0065\u0072\u006e"
}

// HasExtGState checks whether a font is defined by the specified keyName.
func (_edaeb *PdfPageResources) HasExtGState(keyName _abf.PdfObjectName) bool {
	_, _bbdef := _edaeb.GetFontByName(keyName)
	return _bbdef
}

// ToPdfObject implements interface PdfModel.
func (_eagb *PdfAnnotationStamp) ToPdfObject() _abf.PdfObject {
	_eagb.PdfAnnotation.ToPdfObject()
	_dcfc := _eagb._dbc
	_ccc := _dcfc.PdfObject.(*_abf.PdfObjectDictionary)
	_eagb.PdfAnnotationMarkup.appendToPdfDictionary(_ccc)
	_ccc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0053\u0074\u0061m\u0070"))
	_ccc.SetIfNotNil("\u004e\u0061\u006d\u0065", _eagb.Name)
	return _dcfc
}

// PdfAnnotationLink represents Link annotations.
// (Section 12.5.6.5 p. 403).
type PdfAnnotationLink struct {
	*PdfAnnotation
	A          _abf.PdfObject
	Dest       _abf.PdfObject
	H          _abf.PdfObject
	PA         _abf.PdfObject
	QuadPoints _abf.PdfObject
	BS         _abf.PdfObject
	_bgad      *PdfAction
	_aefa      *PdfReader
}

// GetContainingPdfObject returns the page as a dictionary within an PdfIndirectObject.
func (_bbbf *PdfPage) GetContainingPdfObject() _abf.PdfObject { return _bbbf._gefee }

func (_eddbc *PdfWriter) getPdfVersion() string {
	return _e.Sprintf("\u0025\u0064\u002e%\u0064", _eddbc._ecfa.Major, _eddbc._ecfa.Minor)
}

// SetColorspaceByName adds the provided colorspace to the page resources.
func (_ecgbg *PdfPageResources) SetColorspaceByName(keyName _abf.PdfObjectName, cs PdfColorspace) error {
	_cfcaf, _gadgc := _ecgbg.GetColorspaces()
	if _gadgc != nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _gadgc)
		return _gadgc
	}
	if _cfcaf == nil {
		_cfcaf = NewPdfPageResourcesColorspaces()
		_ecgbg.SetColorSpace(_cfcaf)
	}
	_cfcaf.Set(keyName, cs)
	return nil
}

func _agcb(_adabd _abf.PdfObject) (*PdfColorspaceLab, error) {
	_edaaca := NewPdfColorspaceLab()
	if _deggf, _gcbb := _adabd.(*_abf.PdfIndirectObject); _gcbb {
		_edaaca._aaec = _deggf
	}
	_adabd = _abf.TraceToDirectObject(_adabd)
	_ddfg, _bfbbf := _adabd.(*_abf.PdfObjectArray)
	if !_bfbbf {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _ddfg.Len() != 2 {
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_adabd = _abf.TraceToDirectObject(_ddfg.Get(0))
	_eagc, _bfbbf := _adabd.(*_abf.PdfObjectName)
	if !_bfbbf {
		return nil, _e.Errorf("\u006c\u0061\u0062\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	if *_eagc != "\u004c\u0061\u0062" {
		return nil, _e.Errorf("n\u006ft\u0020\u0061\u0020\u004c\u0061\u0062\u0020\u0063o\u006c\u006f\u0072\u0073pa\u0063\u0065")
	}
	_adabd = _abf.TraceToDirectObject(_ddfg.Get(1))
	_begg, _bfbbf := _adabd.(*_abf.PdfObjectDictionary)
	if !_bfbbf {
		return nil, _e.Errorf("c\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006ev\u0061l\u0069\u0064")
	}
	_adabd = _begg.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_adabd = _abf.TraceToDirectObject(_adabd)
	_ggdae, _bfbbf := _adabd.(*_abf.PdfObjectArray)
	if !_bfbbf {
		return nil, _e.Errorf("\u004c\u0061\u0062\u0020In\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069n\u0074")
	}
	if _ggdae.Len() != 3 {
		return nil, _e.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
	}
	_aff, _badeg := _ggdae.GetAsFloat64Slice()
	if _badeg != nil {
		return nil, _badeg
	}
	_edaaca.WhitePoint = _aff
	_adabd = _begg.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _adabd != nil {
		_adabd = _abf.TraceToDirectObject(_adabd)
		_bfcg, _agfea := _adabd.(*_abf.PdfObjectArray)
		if !_agfea {
			return nil, _e.Errorf("\u004c\u0061\u0062: \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
		}
		if _bfcg.Len() != 3 {
			return nil, _e.Errorf("\u004c\u0061b\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061rr\u0061\u0079")
		}
		_cbce, _egcf := _bfcg.GetAsFloat64Slice()
		if _egcf != nil {
			return nil, _egcf
		}
		_edaaca.BlackPoint = _cbce
	}
	_adabd = _begg.Get("\u0052\u0061\u006eg\u0065")
	if _adabd != nil {
		_adabd = _abf.TraceToDirectObject(_adabd)
		_feacg, _egga := _adabd.(*_abf.PdfObjectArray)
		if !_egga {
			_acd.Log.Error("\u0052\u0061n\u0067\u0065\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
			return nil, _e.Errorf("\u004ca\u0062:\u0020\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if _feacg.Len() != 4 {
			_acd.Log.Error("\u0052\u0061\u006e\u0067\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020e\u0072\u0072\u006f\u0072")
			return nil, _e.Errorf("\u004c\u0061b\u003a\u0020\u0052a\u006e\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_gggf, _edee := _feacg.GetAsFloat64Slice()
		if _edee != nil {
			return nil, _edee
		}
		_edaaca.Range = _gggf
	}
	return _edaaca, nil
}

// NewPdfAction returns an initialized generic PDF action model.
func NewPdfAction() *PdfAction {
	_ee := &PdfAction{}
	_ee._egg = _abf.MakeIndirectObject(_abf.MakeDict())
	return _ee
}

func _fggeg(_agecbc *_abf.PdfObjectDictionary, _egaff *fontCommon, _eggg _cbb.TextEncoder) (*pdfFontSimple, error) {
	_cgac := _dedf(_egaff)
	_cgac._edabc = _eggg
	if _eggg == nil {
		_fdcc := _agecbc.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r")
		if _fdcc == nil {
			_fdcc = _abf.MakeInteger(0)
		}
		_cgac.FirstChar = _fdcc
		_aedea, _agfef := _abf.GetIntVal(_fdcc)
		if !_agfef {
			_acd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0046i\u0072s\u0074C\u0068\u0061\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029", _fdcc)
			return nil, _abf.ErrTypeError
		}
		_bcaag := _cbb.CharCode(_aedea)
		_fdcc = _agecbc.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072")
		if _fdcc == nil {
			_fdcc = _abf.MakeInteger(255)
		}
		_cgac.LastChar = _fdcc
		_aedea, _agfef = _abf.GetIntVal(_fdcc)
		if !_agfef {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004c\u0061\u0073\u0074\u0043h\u0061\u0072\u0020\u0074\u0079\u0070\u0065 \u0028\u0025\u0054\u0029", _fdcc)
			return nil, _abf.ErrTypeError
		}
		_aecg := _cbb.CharCode(_aedea)
		_cgac._aadgb = make(map[_cbb.CharCode]float64)
		_fdcc = _agecbc.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _fdcc != nil {
			_cgac.Widths = _fdcc
			_geacb, _degde := _abf.GetArray(_fdcc)
			if !_degde {
				_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020W\u0069\u0064t\u0068\u0073\u0020\u0061\u0074\u0074\u0072\u0069b\u0075\u0074\u0065\u0020\u0021\u003d\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _fdcc)
				return nil, _abf.ErrTypeError
			}
			_gfbg, _gcbgf := _geacb.ToFloat64Array()
			if _gcbgf != nil {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069d\u0074\u0068\u0073\u0020\u0074\u006f\u0020a\u0072\u0072\u0061\u0079")
				return nil, _gcbgf
			}
			if len(_gfbg) != int(_aecg-_bcaag+1) {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0074\u0068s\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0025\u0064 \u0028\u0025\u0064\u0029", _aecg-_bcaag+1, len(_gfbg))
				return nil, _abf.ErrRangeError
			}
			for _cgfa, _gfbc := range _gfbg {
				_cgac._aadgb[_bcaag+_cbb.CharCode(_cgfa)] = _gfbc
			}
		}
	}
	_cgac.Encoding = _abf.TraceToDirectObject(_agecbc.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	return _cgac, nil
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_bffc *PdfColorspaceSpecialPattern) ToPdfObject() _abf.PdfObject {
	if _bffc.UnderlyingCS == nil {
		return _abf.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e")
	}
	_cecd := _abf.MakeArray(_abf.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_cecd.Append(_bffc.UnderlyingCS.ToPdfObject())
	if _bffc._afca != nil {
		_bffc._afca.PdfObject = _cecd
		return _bffc._afca
	}
	return _cecd
}

// GetFillImage get attached model.Image in push button.
func (_faab *PdfFieldButton) GetFillImage() *Image {
	if _faab.IsPush() {
		return _faab._ccdd
	}
	return nil
}

func (_caea *PdfWriter) optimizeDocument() error {
	if _caea._adgdc == nil {
		return nil
	}
	_dbabf, _dceee := _abf.GetDict(_caea._ddegc)
	if !_dceee {
		return _fd.New("\u0061\u006e\u0020in\u0066\u006f\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_eeca := _bbf.Document{ID: [2]string{_caea._aefff, _caea._cfbce}, Version: _caea._ecfa, Objects: _caea._edcgc, Info: _dbabf, Crypt: _caea._ddbgd, UseHashBasedID: _caea._fegae}
	if _aeddb := _caea._adgdc.ApplyStandard(&_eeca); _aeddb != nil {
		return _aeddb
	}
	_caea._aefff, _caea._cfbce = _eeca.ID[0], _eeca.ID[1]
	_caea._ecfa = _eeca.Version
	_caea._edcgc = _eeca.Objects
	_caea._ddegc.PdfObject = _eeca.Info
	_caea._fegae = _eeca.UseHashBasedID
	_caea._ddbgd = _eeca.Crypt
	_cafbg := make(map[_abf.PdfObject]struct{}, len(_caea._edcgc))
	for _, _abdef := range _caea._edcgc {
		_cafbg[_abdef] = struct{}{}
	}
	_caea._fdgae = _cafbg
	return nil
}

// String implements interface PdfObject.
func (_fef *PdfAction) String() string {
	_aaa, _gcg := _fef.ToPdfObject().(*_abf.PdfIndirectObject)
	if _gcg {
		return _e.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _fef._gfg, _aaa.PdfObject.String())
	}
	return ""
}

// GetModelFromPrimitive returns the model corresponding to the `primitive` PdfObject.
func (_gcac *modelManager) GetModelFromPrimitive(primitive _abf.PdfObject) PdfModel {
	model, _gdedf := _gcac._addgc[primitive]
	if !_gdedf {
		return nil
	}
	return model
}

// GetCerts returns the signature certificate chain.
func (_agagec *PdfSignature) GetCerts() ([]*_fa.Certificate, error) {
	var _gcfc []func() ([]*_fa.Certificate, error)
	switch _bccegf, _ := _abf.GetNameVal(_agagec.SubFilter); _bccegf {
	case "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064", "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064":
		_gcfc = append(_gcfc, _agagec.extractChainFromPKCS7, _agagec.extractChainFromCert)
	case "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031":
		_gcfc = append(_gcfc, _agagec.extractChainFromCert)
	case "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031":
		_gcfc = append(_gcfc, _agagec.extractChainFromPKCS7)
	default:
		return nil, _e.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020S\u0075b\u0046i\u006c\u0074\u0065\u0072\u003a\u0020\u0025s", _bccegf)
	}
	for _, _caffa := range _gcfc {
		_bdff, _bbgdff := _caffa()
		if _bbgdff != nil {
			return nil, _bbgdff
		}
		if len(_bdff) > 0 {
			return _bdff, nil
		}
	}
	return nil, ErrSignNoCertificates
}

// ToPdfObject returns the button field dictionary within an indirect object.
func (_ecbf *PdfFieldButton) ToPdfObject() _abf.PdfObject {
	_ecbf.PdfField.ToPdfObject()
	_dgegc := _ecbf._dgdc
	_feccg := _dgegc.PdfObject.(*_abf.PdfObjectDictionary)
	_feccg.Set("\u0046\u0054", _abf.MakeName("\u0042\u0074\u006e"))
	if _ecbf.Opt != nil {
		_feccg.Set("\u004f\u0070\u0074", _ecbf.Opt)
	}
	return _dgegc
}

// NewReaderOpts generates a default `ReaderOpts` instance.
func NewReaderOpts() *ReaderOpts { return &ReaderOpts{Password: "", LazyLoad: true} }

// NewPdfColorCalGray returns a new CalGray color.
func NewPdfColorCalGray(grayVal float64) *PdfColorCalGray {
	_bgdc := PdfColorCalGray(grayVal)
	return &_bgdc
}

func (_adfbe *PdfWriter) mapObjectStreams(_dgeab bool) (map[_abf.PdfObject]bool, bool) {
	_fgecf := make(map[_abf.PdfObject]bool)
	for _, _fagef := range _adfbe._edcgc {
		if _egfgc, _gdbab := _fagef.(*_abf.PdfObjectStreams); _gdbab {
			_dgeab = true
			for _, _gdefc := range _egfgc.Elements() {
				_fgecf[_gdefc] = true
				if _gede, _bbccg := _gdefc.(*_abf.PdfIndirectObject); _bbccg {
					_fgecf[_gede.PdfObject] = true
				}
			}
		}
	}
	return _fgecf, _dgeab
}

// PdfColorspaceDeviceCMYK represents a CMYK32 colorspace.
type PdfColorspaceDeviceCMYK struct{}

// ToPdfObject returns colorspace in a PDF object format [name dictionary]
func (_afgd *PdfColorspaceLab) ToPdfObject() _abf.PdfObject {
	_ccae := _abf.MakeArray()
	_ccae.Append(_abf.MakeName("\u004c\u0061\u0062"))
	_febcg := _abf.MakeDict()
	if _afgd.WhitePoint != nil {
		_ffba := _abf.MakeArray(_abf.MakeFloat(_afgd.WhitePoint[0]), _abf.MakeFloat(_afgd.WhitePoint[1]), _abf.MakeFloat(_afgd.WhitePoint[2]))
		_febcg.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _ffba)
	} else {
		_acd.Log.Error("\u004c\u0061\u0062: \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057h\u0069t\u0065P\u006fi\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029")
	}
	if _afgd.BlackPoint != nil {
		_aggc := _abf.MakeArray(_abf.MakeFloat(_afgd.BlackPoint[0]), _abf.MakeFloat(_afgd.BlackPoint[1]), _abf.MakeFloat(_afgd.BlackPoint[2]))
		_febcg.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _aggc)
	}
	if _afgd.Range != nil {
		_edaae := _abf.MakeArray(_abf.MakeFloat(_afgd.Range[0]), _abf.MakeFloat(_afgd.Range[1]), _abf.MakeFloat(_afgd.Range[2]), _abf.MakeFloat(_afgd.Range[3]))
		_febcg.Set("\u0052\u0061\u006eg\u0065", _edaae)
	}
	_ccae.Append(_febcg)
	if _afgd._aaec != nil {
		_afgd._aaec.PdfObject = _ccae
		return _afgd._aaec
	}
	return _ccae
}

// Outline represents a PDF outline dictionary (Table 152 - p. 376).
// Currently, the Outline object can only be used to construct PDF outlines.
type Outline struct {
	Entries []*OutlineItem `json:"entries,omitempty"`
}

// NewPdfDateFromTime will create a PdfDate based on the given time
func NewPdfDateFromTime(timeObj _f.Time) (PdfDate, error) {
	_cdfbcc := timeObj.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_cfabe, _ := _gb.ParseInt(_cdfbcc[1:3], 10, 32)
	_bdfdf, _ := _gb.ParseInt(_cdfbcc[4:6], 10, 32)
	return PdfDate{_fabd: int64(timeObj.Year()), _fcdacf: int64(timeObj.Month()), _gecdc: int64(timeObj.Day()), _ebda: int64(timeObj.Hour()), _efba: int64(timeObj.Minute()), _fgddf: int64(timeObj.Second()), _aggabc: _cdfbcc[0], _dbgccd: _cfabe, _ccfca: _bdfdf}, nil
}

// ToPdfObject implements interface PdfModel.
func (_bedae *PdfTransformParamsDocMDP) ToPdfObject() _abf.PdfObject {
	_geebea := _abf.MakeDict()
	_geebea.SetIfNotNil("\u0054\u0079\u0070\u0065", _bedae.Type)
	_geebea.SetIfNotNil("\u0056", _bedae.V)
	_geebea.SetIfNotNil("\u0050", _bedae.P)
	return _geebea
}

func (_gbeab *PdfReader) newPdfFieldSignatureFromDict(_gfbbe *_abf.PdfObjectDictionary) (*PdfFieldSignature, error) {
	_dccae := &PdfFieldSignature{}
	_dcge, _aaba := _abf.GetIndirect(_gfbbe.Get("\u0056"))
	if _aaba {
		var _gdddc error
		_dccae.V, _gdddc = _gbeab.newPdfSignatureFromIndirect(_dcge)
		if _gdddc != nil {
			return nil, _gdddc
		}
	}
	_dccae.Lock, _ = _abf.GetIndirect(_gfbbe.Get("\u004c\u006f\u0063\u006b"))
	_dccae.SV, _ = _abf.GetIndirect(_gfbbe.Get("\u0053\u0056"))
	return _dccae, nil
}

func _gebbf(_fcgfg _abf.PdfObject) (*PdfPageResourcesColorspaces, error) {
	_gbac := &PdfPageResourcesColorspaces{}
	if _addbfb, _gedaf := _fcgfg.(*_abf.PdfIndirectObject); _gedaf {
		_gbac._cebc = _addbfb
		_fcgfg = _addbfb.PdfObject
	}
	_dbadf, _egdca := _abf.GetDict(_fcgfg)
	if !_egdca {
		return nil, _fd.New("\u0043\u0053\u0020at\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gbac.Names = []string{}
	_gbac.Colorspaces = map[string]PdfColorspace{}
	for _, _acbbd := range _dbadf.Keys() {
		_addgb := _dbadf.Get(_acbbd)
		_gbac.Names = append(_gbac.Names, string(_acbbd))
		_acfbcd, _deeab := NewPdfColorspaceFromPdfObject(_addgb)
		if _deeab != nil {
			return nil, _deeab
		}
		_gbac.Colorspaces[string(_acbbd)] = _acfbcd
	}
	return _gbac, nil
}

// DecodeArray returns the range of color component values in DeviceCMYK colorspace.
func (_facca *PdfColorspaceDeviceCMYK) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// ToPdfObject implements interface PdfModel.
func (_cggf *PdfAnnotationSound) ToPdfObject() _abf.PdfObject {
	_cggf.PdfAnnotation.ToPdfObject()
	_dadd := _cggf._dbc
	_adff := _dadd.PdfObject.(*_abf.PdfObjectDictionary)
	_cggf.PdfAnnotationMarkup.appendToPdfDictionary(_adff)
	_adff.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0053\u006f\u0075n\u0064"))
	_adff.SetIfNotNil("\u0053\u006f\u0075n\u0064", _cggf.Sound)
	_adff.SetIfNotNil("\u004e\u0061\u006d\u0065", _cggf.Name)
	return _dadd
}

// ToPdfObject implements interface PdfModel.
func (_ggbd *PdfAnnotationInk) ToPdfObject() _abf.PdfObject {
	_ggbd.PdfAnnotation.ToPdfObject()
	_bdfa := _ggbd._dbc
	_ceeb := _bdfa.PdfObject.(*_abf.PdfObjectDictionary)
	_ggbd.PdfAnnotationMarkup.appendToPdfDictionary(_ceeb)
	_ceeb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0049\u006e\u006b"))
	_ceeb.SetIfNotNil("\u0049n\u006b\u004c\u0069\u0073\u0074", _ggbd.InkList)
	_ceeb.SetIfNotNil("\u0042\u0053", _ggbd.BS)
	return _bdfa
}

// NewPdfAnnotationInk returns a new ink annotation.
func NewPdfAnnotationInk() *PdfAnnotationInk {
	_afea := NewPdfAnnotation()
	_bffb := &PdfAnnotationInk{}
	_bffb.PdfAnnotation = _afea
	_bffb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_afea.SetContext(_bffb)
	return _bffb
}

var (
	_fabag = _af.MustCompile("\u005cd\u002b\u0020\u0064\u0069c\u0074\u005c\u0073\u002b\u0028d\u0075p\u005cs\u002b\u0029\u003f\u0062\u0065\u0067\u0069n")
	_geaa  = _af.MustCompile("\u005e\u005cs\u002a\u002f\u0028\u005c\u0053\u002b\u003f\u0029\u005c\u0073\u002b\u0028\u002e\u002b\u003f\u0029\u005c\u0073\u002b\u0064\u0065\u0066\\s\u002a\u0024")
	_gffgf = _af.MustCompile("\u005e\u005c\u0073*\u0064\u0075\u0070\u005c\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002a\u002f\u0028\u005c\u0077\u002b\u003f\u0029\u0028\u003f\u003a\u005c\u002e\u005c\u0064\u002b)\u003f\u005c\u0073\u002b\u0070\u0075\u0074\u0024")
	_ffed  = "\u002f\u0045\u006e\u0063od\u0069\u006e\u0067\u0020\u0032\u0035\u0036\u0020\u0061\u0072\u0072\u0061\u0079"
	_bdec  = "\u0072\u0065\u0061d\u006f\u006e\u006c\u0079\u0020\u0064\u0065\u0066"
	_gabdf = "\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u0066\u0069\u006c\u0065\u0020e\u0065\u0078\u0065\u0063"
)

func _aaagb(_fdff _abf.PdfObject, _gcbba *PdfReader) (*OutlineDest, error) {
	_bdbdb, _badef := _abf.GetArray(_fdff)
	if !_badef {
		return nil, _fd.New("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0064\u0065\u0073t\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_afefa := _bdbdb.Len()
	if _afefa < 2 {
		return nil, _e.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u006fu\u0074\u006c\u0069n\u0065\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061ti\u006f\u006e\u0020a\u0072\u0072a\u0079\u0020\u006c\u0065\u006e\u0067t\u0068\u003a \u0025\u0064", _afefa)
	}
	_ecgd := &OutlineDest{Mode: "\u0046\u0069\u0074"}
	_dbgadc := _bdbdb.Get(0)
	if _dbcfg, _cgafa := _abf.GetIndirect(_dbgadc); _cgafa {
		if _, _aabgf, _bfgbe := _gcbba.PageFromIndirectObject(_dbcfg); _bfgbe == nil {
			_ecgd.Page = int64(_aabgf - 1)
		} else {
			_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020g\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u002b\u0076", _dbcfg)
		}
		_ecgd.PageObj = _dbcfg
	} else if _ecdbd, _cddgb := _abf.GetIntVal(_dbgadc); _cddgb {
		if _ecdbd >= 0 && _ecdbd < len(_gcbba.PageList) {
			_ecgd.PageObj = _gcbba.PageList[_ecdbd].GetPageAsIndirectObject()
		} else {
			_acd.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _ecdbd)
		}
		_ecgd.Page = int64(_ecdbd)
	} else {
		return nil, _e.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006f\u0075\u0074\u006cine\u0020de\u0073\u0074\u0069\u006e\u0061\u0074\u0069on\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0054", _dbgadc)
	}
	_eafb, _badef := _abf.GetNameVal(_bdbdb.Get(1))
	if !_badef {
		_acd.Log.Debug("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006d\u0061\u0067\u006e\u0069\u0066\u0069\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _bdbdb.Get(1))
		return _ecgd, nil
	}
	switch _eafb {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		if _afefa > 2 {
			_ecgd.Y, _ = _abf.GetNumberAsFloat(_abf.TraceToDirectObject(_bdbdb.Get(2)))
		}
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		if _afefa > 2 {
			_ecgd.X, _ = _abf.GetNumberAsFloat(_abf.TraceToDirectObject(_bdbdb.Get(2)))
		}
	case "\u0058\u0059\u005a":
		if _afefa > 4 {
			_ecgd.X, _ = _abf.GetNumberAsFloat(_abf.TraceToDirectObject(_bdbdb.Get(2)))
			_ecgd.Y, _ = _abf.GetNumberAsFloat(_abf.TraceToDirectObject(_bdbdb.Get(3)))
			_ecgd.Zoom, _ = _abf.GetNumberAsFloat(_abf.TraceToDirectObject(_bdbdb.Get(4)))
		}
	default:
		_eafb = "\u0046\u0069\u0074"
	}
	_ecgd.Mode = _eafb
	return _ecgd, nil
}

// GetPdfName returns the PDF name used to indicate the border style.
// (Table 166 p. 395).
func (_ced *BorderStyle) GetPdfName() string {
	switch *_ced {
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

// Evaluate runs the function on the passed in slice and returns the results.
func (_abdfb *PdfFunctionType3) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_acd.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return nil, _fd.New("\u006e\u006f\u0074\u0020im\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
}

func (_cgba *PdfReader) newPdfAnnotationUnderlineFromDict(_eca *_abf.PdfObjectDictionary) (*PdfAnnotationUnderline, error) {
	_fad := PdfAnnotationUnderline{}
	_gae, _ggec := _cgba.newPdfAnnotationMarkupFromDict(_eca)
	if _ggec != nil {
		return nil, _ggec
	}
	_fad.PdfAnnotationMarkup = _gae
	_fad.QuadPoints = _eca.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_fad, nil
}

func _dgf(_ddg _abf.PdfObject) (*PdfFilespec, error) {
	if _ddg == nil {
		return nil, nil
	}
	return NewPdfFilespecFromObj(_ddg)
}

// SetFontByName sets the font specified by keyName to the given object.
func (_fafg *PdfPageResources) SetFontByName(keyName _abf.PdfObjectName, obj _abf.PdfObject) error {
	if _fafg.Font == nil {
		_fafg.Font = _abf.MakeDict()
	}
	_gaggf, _efdf := _abf.TraceToDirectObject(_fafg.Font).(*_abf.PdfObjectDictionary)
	if !_efdf {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _abf.TraceToDirectObject(_fafg.Font))
		return _abf.ErrTypeError
	}
	_gaggf.Set(keyName, obj)
	return nil
}
func (_ccfb *pdfCIDFontType0) baseFields() *fontCommon { return &_ccfb.fontCommon }

// AddCRLs adds CRLs to DSS.
func (_ebcdc *DSS) AddCRLs(crls [][]byte) ([]*_abf.PdfObjectStream, error) {
	return _ebcdc.add(&_ebcdc.CRLs, _ebcdc._daee, crls)
}

func _fdaf(_ecefg []*_abf.PdfObjectStream) *_abf.PdfObjectArray {
	if len(_ecefg) == 0 {
		return nil
	}
	_acfcc := make([]_abf.PdfObject, 0, len(_ecefg))
	for _, _efdae := range _ecefg {
		_acfcc = append(_acfcc, _efdae)
	}
	return _abf.MakeArray(_acfcc...)
}

func (_affda *PdfWriter) writeObject(_eagce int, _cbdgaa _abf.PdfObject) {
	_acd.Log.Trace("\u0057\u0072\u0069\u0074\u0065\u0020\u006f\u0062\u006a \u0023\u0025\u0064\u000a", _eagce)
	if _eefgb, _ceacea := _cbdgaa.(*_abf.PdfIndirectObject); _ceacea {
		_affda._becfc[_eagce] = crossReference{Type: 1, Offset: _affda._dbfaad, Generation: _eefgb.GenerationNumber}
		_ffgd := _e.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _eagce)
		if _bgbeca, _caagb := _eefgb.PdfObject.(*pdfSignDictionary); _caagb {
			_bgbeca._eefbf = _affda._dbfaad + int64(len(_ffgd))
		}
		if _eefgb.PdfObject == nil {
			_acd.Log.Debug("E\u0072\u0072\u006fr\u003a\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0027\u0073\u0020\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020b\u0065\u0020\u006e\u0069l\u0020\u002d\u0020\u0073e\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063t\u004e\u0075\u006c\u006c")
			_eefgb.PdfObject = _abf.MakeNull()
		}
		_ffgd += _eefgb.PdfObject.WriteString()
		_ffgd += "\u000a\u0065\u006e\u0064\u006f\u0062\u006a\u000a"
		_affda.writeString(_ffgd)
		return
	}
	if _dgedg, _eccegf := _cbdgaa.(*_abf.PdfObjectStream); _eccegf {
		_affda._becfc[_eagce] = crossReference{Type: 1, Offset: _affda._dbfaad, Generation: _dgedg.GenerationNumber}
		_bceca := _e.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _eagce)
		_bceca += _dgedg.PdfObjectDictionary.WriteString()
		_bceca += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_affda.writeString(_bceca)
		_affda.writeBytes(_dgedg.Stream)
		_affda.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	if _decegc, _dageg := _cbdgaa.(*_abf.PdfObjectStreams); _dageg {
		_affda._becfc[_eagce] = crossReference{Type: 1, Offset: _affda._dbfaad, Generation: _decegc.GenerationNumber}
		_fdacc := _e.Sprintf("\u0025d\u0020\u0030\u0020\u006f\u0062\u006a\n", _eagce)
		var _gdbaa []string
		var _gbdde string
		var _cebda int64
		for _bedcee, _bcafg := range _decegc.Elements() {
			_gbgac, _efaca := _bcafg.(*_abf.PdfIndirectObject)
			if !_efaca {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065am\u0073 \u004e\u0020\u0025\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006es\u0020\u006e\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u0070\u0064\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0076", _eagce, _bcafg)
				continue
			}
			_ffeac := _gbgac.PdfObject.WriteString() + "\u0020"
			_gbdde = _gbdde + _ffeac
			_gdbaa = append(_gdbaa, _e.Sprintf("\u0025\u0064\u0020%\u0064", _gbgac.ObjectNumber, _cebda))
			_affda._becfc[int(_gbgac.ObjectNumber)] = crossReference{Type: 2, ObjectNumber: _eagce, Index: _bedcee}
			_cebda = _cebda + int64(len([]byte(_ffeac)))
		}
		_cfgb := _be.Join(_gdbaa, "\u0020") + "\u0020"
		_dgfdd := _abf.NewFlateEncoder()
		_dbeef := _dgfdd.MakeStreamDict()
		_dbeef.Set(_abf.PdfObjectName("\u0054\u0079\u0070\u0065"), _abf.MakeName("\u004f\u0062\u006a\u0053\u0074\u006d"))
		_eeaac := int64(_decegc.Len())
		_dbeef.Set(_abf.PdfObjectName("\u004e"), _abf.MakeInteger(_eeaac))
		_gdgaf := int64(len(_cfgb))
		_dbeef.Set(_abf.PdfObjectName("\u0046\u0069\u0072s\u0074"), _abf.MakeInteger(_gdgaf))
		_geeeg, _ := _dgfdd.EncodeBytes([]byte(_cfgb + _gbdde))
		_ffebb := int64(len(_geeeg))
		_dbeef.Set(_abf.PdfObjectName("\u004c\u0065\u006e\u0067\u0074\u0068"), _abf.MakeInteger(_ffebb))
		_fdacc += _dbeef.WriteString()
		_fdacc += "\u000a\u0073\u0074\u0072\u0065\u0061\u006d\u000a"
		_affda.writeString(_fdacc)
		_affda.writeBytes(_geeeg)
		_affda.writeString("\u000ae\u006ed\u0073\u0074\u0072\u0065\u0061m\u000a\u0065n\u0064\u006f\u0062\u006a\u000a")
		return
	}
	_affda.writeString(_cbdgaa.WriteString())
}

// PdfFieldText represents a text field where user can enter text.
type PdfFieldText struct {
	*PdfField
	DA     *_abf.PdfObjectString
	Q      *_abf.PdfObjectInteger
	DS     *_abf.PdfObjectString
	RV     _abf.PdfObject
	MaxLen *_abf.PdfObjectInteger
}

var _gcgde = map[string]string{"\u0053\u0079\u006d\u0062\u006f\u006c": "\u0053\u0079\u006d\u0062\u006f\u006c\u0045\u006e\u0063o\u0064\u0069\u006e\u0067", "\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073": "Z\u0061p\u0066\u0044\u0069\u006e\u0067\u0062\u0061\u0074s\u0045\u006e\u0063\u006fdi\u006e\u0067"}

// NewPdfAnnotationWidget returns an initialized annotation widget.
func NewPdfAnnotationWidget() *PdfAnnotationWidget {
	_ggaf := NewPdfAnnotation()
	_gfge := &PdfAnnotationWidget{}
	_gfge.PdfAnnotation = _ggaf
	_ggaf.SetContext(_gfge)
	return _gfge
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_bgef *PdfColorspaceDeviceRGB) ToPdfObject() _abf.PdfObject {
	return _abf.MakeName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_fdef *PdfColorspaceSpecialIndexed) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	N := _fdef.Base.GetNumComponents()
	_edace := int(vals[0]) * N
	if _edace < 0 || (_edace+N-1) >= len(_fdef._bcdf) {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _edace)
		return nil, ErrColorOutOfRange
	}
	_fged := _fdef._bcdf[_edace : _edace+N]
	var _bdea []float64
	for _, _cagd := range _fged {
		_bdea = append(_bdea, float64(_cagd)/255.0)
	}
	_ffab, _eagf := _fdef.Base.ColorFromFloats(_bdea)
	if _eagf != nil {
		return nil, _eagf
	}
	return _ffab, nil
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the A, B and C components of the color.
func (_ceca *PdfColorspaceCalRGB) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_eccf, _ecge := _abf.GetNumbersAsFloat(objects)
	if _ecge != nil {
		return nil, _ecge
	}
	return _ceca.ColorFromFloats(_eccf)
}

func (_bbdeb *PdfReader) loadPerms() (*Permissions, error) {
	if _aadc := _bbdeb._dagde.Get("\u0050\u0065\u0072m\u0073"); _aadc != nil {
		if _abcf, _bbfdee := _abf.GetDict(_aadc); _bbfdee {
			_eddaf := _abcf.Get("\u0044\u006f\u0063\u004d\u0044\u0050")
			if _eddaf == nil {
				return nil, nil
			}
			if _gfcgec, _fddf := _abf.GetIndirect(_eddaf); _fddf {
				_ebfa, _gfddc := _bbdeb.newPdfSignatureFromIndirect(_gfcgec)
				if _gfddc != nil {
					return nil, _gfddc
				}
				return NewPermissions(_ebfa), nil
			}
			return nil, _e.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u006f\u0063M\u0044\u0050\u0020\u0065nt\u0072\u0079")
		}
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0050\u0065\u0072\u006d\u0073\u0020\u0065\u006e\u0074\u0072\u0079")
	}
	return nil, nil
}

// PdfAnnotationFreeText represents FreeText annotations.
// (Section 12.5.6.6).
type PdfAnnotationFreeText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	DA _abf.PdfObject
	Q  _abf.PdfObject
	RC _abf.PdfObject
	DS _abf.PdfObject
	CL _abf.PdfObject
	IT _abf.PdfObject
	BE _abf.PdfObject
	RD _abf.PdfObject
	BS _abf.PdfObject
	LE _abf.PdfObject
}

func (_bgggd *PdfWriter) setWriter(_bcaeg _gc.Writer) {
	_bgggd._dbfaad = _bgggd._cgded
	_bgggd._agfba = _ac.NewWriter(_bcaeg)
}

// PdfOutlineTreeNode contains common fields used by the outline and outline
// item objects.
type PdfOutlineTreeNode struct {
	_aecec interface{}
	First  *PdfOutlineTreeNode
	Last   *PdfOutlineTreeNode
}

func (_fgf *PdfColorspaceCalRGB) String() string { return "\u0043\u0061\u006c\u0052\u0047\u0042" }

// NewPdfPage returns a new PDF page.
func NewPdfPage() *PdfPage {
	_edaf := PdfPage{}
	_edaf._bdbfa = _abf.MakeDict()
	_edaf.Resources = NewPdfPageResources()
	_fegcd := _abf.PdfIndirectObject{}
	_fegcd.PdfObject = _edaf._bdbfa
	_edaf._gefee = &_fegcd
	_edaf._efca = *_edaf._bdbfa
	return &_edaf
}
func _dedf(_eeba *fontCommon) *pdfFontSimple { return &pdfFontSimple{fontCommon: *_eeba} }
func (_bbcfe *PdfWriter) copyObject(_abecdg _abf.PdfObject, _ggfd map[_abf.PdfObject]_abf.PdfObject, _febaf map[_abf.PdfObject]struct{}, _ccba bool) _abf.PdfObject {
	_aaeggg := !_bbcfe._aegbd && _febaf != nil
	if _gfegg, _agbae := _ggfd[_abecdg]; _agbae {
		if _aaeggg && !_ccba {
			delete(_febaf, _abecdg)
		}
		return _gfegg
	}
	if _abecdg == nil {
		_cdfbe := _abf.MakeNull()
		return _cdfbe
	}
	_gbbdb := _abecdg
	switch _efcg := _abecdg.(type) {
	case *_abf.PdfObjectArray:
		_gcbfa := _abf.MakeArray()
		_gbbdb = _gcbfa
		_ggfd[_abecdg] = _gbbdb
		for _, _gffagd := range _efcg.Elements() {
			_gcbfa.Append(_bbcfe.copyObject(_gffagd, _ggfd, _febaf, _ccba))
		}
	case *_abf.PdfObjectStreams:
		_dcdg := &_abf.PdfObjectStreams{PdfObjectReference: _efcg.PdfObjectReference}
		_gbbdb = _dcdg
		_ggfd[_abecdg] = _gbbdb
		for _, _dcfbc := range _efcg.Elements() {
			_dcdg.Append(_bbcfe.copyObject(_dcfbc, _ggfd, _febaf, _ccba))
		}
	case *_abf.PdfObjectStream:
		_gfaf := &_abf.PdfObjectStream{Stream: _efcg.Stream, PdfObjectReference: _efcg.PdfObjectReference}
		_gbbdb = _gfaf
		_ggfd[_abecdg] = _gbbdb
		_gfaf.PdfObjectDictionary = _bbcfe.copyObject(_efcg.PdfObjectDictionary, _ggfd, _febaf, _ccba).(*_abf.PdfObjectDictionary)
	case *_abf.PdfObjectDictionary:
		var _cddc bool
		if _aaeggg && !_ccba {
			if _ccbgd, _ := _abf.GetNameVal(_efcg.Get("\u0054\u0079\u0070\u0065")); _ccbgd == "\u0050\u0061\u0067\u0065" {
				_, _fcbga := _bbcfe._aadb[_efcg]
				_ccba = !_fcbga
				_cddc = _ccba
			}
		}
		_efae := _abf.MakeDict()
		_gbbdb = _efae
		_ggfd[_abecdg] = _gbbdb
		for _, _eagde := range _efcg.Keys() {
			_efae.Set(_eagde, _bbcfe.copyObject(_efcg.Get(_eagde), _ggfd, _febaf, _ccba))
		}
		if _cddc {
			_gbbdb = _abf.MakeNull()
			_ccba = false
		}
	case *_abf.PdfIndirectObject:
		_cbdga := &_abf.PdfIndirectObject{PdfObjectReference: _efcg.PdfObjectReference}
		_gbbdb = _cbdga
		_ggfd[_abecdg] = _gbbdb
		_cbdga.PdfObject = _bbcfe.copyObject(_efcg.PdfObject, _ggfd, _febaf, _ccba)
	case *_abf.PdfObjectString:
		_badcd := *_efcg
		_gbbdb = &_badcd
		_ggfd[_abecdg] = _gbbdb
	case *_abf.PdfObjectName:
		_dffcfc := *_efcg
		_gbbdb = &_dffcfc
		_ggfd[_abecdg] = _gbbdb
	case *_abf.PdfObjectNull:
		_gbbdb = _abf.MakeNull()
		_ggfd[_abecdg] = _gbbdb
	case *_abf.PdfObjectInteger:
		_adfcf := *_efcg
		_gbbdb = &_adfcf
		_ggfd[_abecdg] = _gbbdb
	case *_abf.PdfObjectReference:
		_bcbf := *_efcg
		_gbbdb = &_bcbf
		_ggfd[_abecdg] = _gbbdb
	case *_abf.PdfObjectFloat:
		_gfegc := *_efcg
		_gbbdb = &_gfegc
		_ggfd[_abecdg] = _gbbdb
	case *_abf.PdfObjectBool:
		_fgba := *_efcg
		_gbbdb = &_fgba
		_ggfd[_abecdg] = _gbbdb
	case *pdfSignDictionary:
		_baaf := &pdfSignDictionary{PdfObjectDictionary: _abf.MakeDict(), _fafgf: _efcg._fafgf, _dcbed: _efcg._dcbed}
		_gbbdb = _baaf
		_ggfd[_abecdg] = _gbbdb
		for _, _ebdg := range _efcg.Keys() {
			_baaf.Set(_ebdg, _bbcfe.copyObject(_efcg.Get(_ebdg), _ggfd, _febaf, _ccba))
		}
	default:
		_acd.Log.Info("\u0054\u004f\u0044\u004f\u0028\u0061\u0035\u0069\u0029\u003a\u0020\u0069\u006dp\u006c\u0065\u006d\u0065\u006e\u0074 \u0063\u006f\u0070\u0079\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0066\u006fr\u0020\u0025\u002b\u0076", _abecdg)
	}
	if _aaeggg && _ccba {
		_febaf[_abecdg] = struct{}{}
	}
	return _gbbdb
}

// NewPdfActionHide returns a new "hide" action.
func NewPdfActionHide() *PdfActionHide {
	_fcd := NewPdfAction()
	_cga := &PdfActionHide{}
	_cga.PdfAction = _fcd
	_fcd.SetContext(_cga)
	return _cga
}

// PdfVersion returns version of the PDF file.
func (_bdbe *PdfReader) PdfVersion() _abf.Version { return _bdbe._bebc.PdfVersion() }

// NewPdfAnnotationPopup returns a new popup annotation.
func NewPdfAnnotationPopup() *PdfAnnotationPopup {
	_ace := NewPdfAnnotation()
	_ccf := &PdfAnnotationPopup{}
	_ccf.PdfAnnotation = _ace
	_ace.SetContext(_ccf)
	return _ccf
}

// FieldFlag represents form field flags. Some of the flags can apply to all types of fields whereas other
// flags are specific.
type FieldFlag uint32

// NewPdfAnnotationScreen returns a new screen annotation.
func NewPdfAnnotationScreen() *PdfAnnotationScreen {
	_geag := NewPdfAnnotation()
	_efgc := &PdfAnnotationScreen{}
	_efgc.PdfAnnotation = _geag
	_geag.SetContext(_efgc)
	return _efgc
}

// HasShadingByName checks whether a shading is defined by the specified keyName.
func (_fafee *PdfPageResources) HasShadingByName(keyName _abf.PdfObjectName) bool {
	_, _cgcab := _fafee.GetShadingByName(keyName)
	return _cgcab
}

// SetFillImage attach a model.Image to push button.
func (_bcfdb *PdfFieldButton) SetFillImage(image *Image) {
	if _bcfdb.IsPush() {
		_bcfdb._ccdd = image
	}
}

// ToPdfObject returns the choice field dictionary within an indirect object (container).
func (_gdgcg *PdfFieldChoice) ToPdfObject() _abf.PdfObject {
	_gdgcg.PdfField.ToPdfObject()
	_ddedf := _gdgcg._dgdc
	_ggce := _ddedf.PdfObject.(*_abf.PdfObjectDictionary)
	_ggce.Set("\u0046\u0054", _abf.MakeName("\u0043\u0068"))
	if _gdgcg.Opt != nil {
		_ggce.Set("\u004f\u0070\u0074", _gdgcg.Opt)
	}
	if _gdgcg.TI != nil {
		_ggce.Set("\u0054\u0049", _gdgcg.TI)
	}
	if _gdgcg.I != nil {
		_ggce.Set("\u0049", _gdgcg.I)
	}
	return _ddedf
}

// Has checks if flag fl is set in flag and returns true if so, false otherwise.
func (_ffdd FieldFlag) Has(fl FieldFlag) bool { return (_ffdd.Mask() & fl.Mask()) > 0 }

// NewPdfSignature creates a new PdfSignature object.
func NewPdfSignature(handler SignatureHandler) *PdfSignature {
	_bacea := &PdfSignature{Type: _abf.MakeName("\u0053\u0069\u0067"), Handler: handler}
	_egafe := &pdfSignDictionary{PdfObjectDictionary: _abf.MakeDict(), _fafgf: &handler, _dcbed: _bacea}
	_bacea._geebd = _abf.MakeIndirectObject(_egafe)
	return _bacea
}
func (_fedd *pdfFontType0) baseFields() *fontCommon { return &_fedd.fontCommon }

// ColorToRGB converts a CalRGB color to an RGB color.
func (_aeeg *PdfColorspaceCalRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_ccbd, _ffeb := color.(*PdfColorCalRGB)
	if !_ffeb {
		_acd.Log.Debug("\u0049\u006e\u0070ut\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0063\u0061\u006c\u0020\u0072\u0067\u0062")
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_cgdb := _ccbd.A()
	_caed := _ccbd.B()
	_ffbc := _ccbd.C()
	X := _aeeg.Matrix[0]*_ge.Pow(_cgdb, _aeeg.Gamma[0]) + _aeeg.Matrix[3]*_ge.Pow(_caed, _aeeg.Gamma[1]) + _aeeg.Matrix[6]*_ge.Pow(_ffbc, _aeeg.Gamma[2])
	Y := _aeeg.Matrix[1]*_ge.Pow(_cgdb, _aeeg.Gamma[0]) + _aeeg.Matrix[4]*_ge.Pow(_caed, _aeeg.Gamma[1]) + _aeeg.Matrix[7]*_ge.Pow(_ffbc, _aeeg.Gamma[2])
	Z := _aeeg.Matrix[2]*_ge.Pow(_cgdb, _aeeg.Gamma[0]) + _aeeg.Matrix[5]*_ge.Pow(_caed, _aeeg.Gamma[1]) + _aeeg.Matrix[8]*_ge.Pow(_ffbc, _aeeg.Gamma[2])
	_ddec := 3.240479*X + -1.537150*Y + -0.498535*Z
	_ebfb := -0.969256*X + 1.875992*Y + 0.041556*Z
	_dfbcc := 0.055648*X + -0.204043*Y + 1.057311*Z
	_ddec = _ge.Min(_ge.Max(_ddec, 0), 1.0)
	_ebfb = _ge.Min(_ge.Max(_ebfb, 0), 1.0)
	_dfbcc = _ge.Min(_ge.Max(_dfbcc, 0), 1.0)
	return NewPdfColorDeviceRGB(_ddec, _ebfb, _dfbcc), nil
}

// IsEncrypted returns true if the PDF file is encrypted.
func (_ddedd *PdfReader) IsEncrypted() (bool, error) { return _ddedd._bebc.IsEncrypted() }

// FieldImageProvider provides fields images for specified fields.
type FieldImageProvider interface {
	FieldImageValues() (map[string]*Image, error)
}

// Read reads an image and loads into a new Image object with an RGB
// colormap and 8 bits per component.
func (_defcee DefaultImageHandler) Read(reader _gc.Reader) (*Image, error) {
	_dbfeb, _, _ffbg := _aa.Decode(reader)
	if _ffbg != nil {
		_acd.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _ffbg)
		return nil, _ffbg
	}
	return _defcee.NewImageFromGoImage(_dbfeb)
}

// CharcodesToStrings returns the unicode strings corresponding to `charcodes`.
// The int returns are the number of strings and the number of unconvereted codes.
// NOTE: The number of strings returned is equal to the number of charcodes
func (_bafd *PdfFont) CharcodesToStrings(charcodes []_cbb.CharCode) ([]string, int, int) {
	_ceef := _bafd.baseFields()
	_geeff := make([]string, 0, len(charcodes))
	_afbb := 0
	_edec := _bafd.Encoder()
	_bgbe := _ceef._aabfe != nil && _bafd.IsSimple() && _bafd.Subtype() == "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" && !_be.Contains(_ceef._aabfe.Name(), "\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")
	if !_bgbe && _edec != nil {
		switch _ebfg := _edec.(type) {
		case _cbb.SimpleEncoder:
			_efgcb := _ebfg.BaseName()
			if _, _afgg := _bdgdc[_efgcb]; _afgg {
				for _, _fedb := range charcodes {
					if _ddagc, _aeab := _edec.CharcodeToRune(_fedb); _aeab {
						_geeff = append(_geeff, string(_ddagc))
					} else {
						_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _fedb, charcodes, _ceef.isCIDFont(), _bafd, _edec)
						_afbb++
						_geeff = append(_geeff, _bd.MissingCodeString)
					}
				}
				return _geeff, len(_geeff), _afbb
			}
		}
	}
	for _, _ggbf := range charcodes {
		if _ceef._aabfe != nil {
			if _ecab, _dfcf := _ceef._aabfe.CharcodeToUnicode(_bd.CharCode(_ggbf)); _dfcf {
				_geeff = append(_geeff, _ecab)
				continue
			}
		}
		if _edec != nil {
			if _gdabad, _geafc := _edec.CharcodeToRune(_ggbf); _geafc {
				_geeff = append(_geeff, string(_gdabad))
				continue
			}
		}
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0072u\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065=\u0030x\u0025\u0030\u0034\u0078\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0073\u003d\u005b\u0025\u00200\u0034\u0078\u005d\u0020\u0043\u0049\u0044\u003d\u0025\u0074\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073\n\u0009\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003d\u0025\u0073", _ggbf, charcodes, _ceef.isCIDFont(), _bafd, _edec)
		_afbb++
		_geeff = append(_geeff, _bd.MissingCodeString)
	}
	if _afbb != 0 {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0043\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0074\u006f\u0020u\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020i\u006ep\u0075\u0074\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0043\u0068\u0061\u0072\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\u0009\u0066\u006f\u006e\u0074\u003d\u0025\u0073", len(charcodes), _afbb, _bafd)
	}
	return _geeff, len(_geeff), _afbb
}

func (_aabb *PdfReader) newPdfActionGotoRFromDict(_ddf *_abf.PdfObjectDictionary) (*PdfActionGoToR, error) {
	_dfa, _cdd := _dgf(_ddf.Get("\u0046"))
	if _cdd != nil {
		return nil, _cdd
	}
	return &PdfActionGoToR{D: _ddf.Get("\u0044"), NewWindow: _ddf.Get("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw"), F: _dfa}, nil
}

// ToPdfObject implements interface PdfModel.
func (_geb *PdfActionGoToR) ToPdfObject() _abf.PdfObject {
	_geb.PdfAction.ToPdfObject()
	_ba := _geb._egg
	_ggc := _ba.PdfObject.(*_abf.PdfObjectDictionary)
	_ggc.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeGoToR)))
	if _geb.F != nil {
		_ggc.Set("\u0046", _geb.F.ToPdfObject())
	}
	_ggc.SetIfNotNil("\u0044", _geb.D)
	_ggc.SetIfNotNil("\u004ee\u0077\u0057\u0069\u006e\u0064\u006fw", _geb.NewWindow)
	return _ba
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the reader using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
// At the end of the process, the AcroForm contains all the fields which were
// not flattened. If all fields are flattened, the reader's AcroForm field
// is set to nil.
func (_geed *PdfReader) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	return _geed.flattenFieldsWithOpts(false, appgen, opts)
}

// ToPdfObject returns the PDF representation of the outline tree node.
func (_dfaa *PdfOutlineTreeNode) ToPdfObject() _abf.PdfObject {
	return _dfaa.GetContext().ToPdfObject()
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
	_dcfb           *_abf.PdfObjectDictionary
}

func _ddadg(_daea *PdfField) []*PdfField {
	_abdcg := []*PdfField{_daea}
	for _, _bgefd := range _daea.Kids {
		_abdcg = append(_abdcg, _ddadg(_bgefd)...)
	}
	return _abdcg
}

// GetContainingPdfObject returns the container of the DSS (indirect object).
func (_cagg *DSS) GetContainingPdfObject() _abf.PdfObject { return _cagg._gffg }

func _cecag(_abab _abf.PdfObject) (*PdfColorspaceSpecialSeparation, error) {
	_caddd := NewPdfColorspaceSpecialSeparation()
	if _effe, _efdcb := _abab.(*_abf.PdfIndirectObject); _efdcb {
		_caddd._bbed = _effe
	}
	_abab = _abf.TraceToDirectObject(_abab)
	_eace, _fbfg := _abab.(*_abf.PdfObjectArray)
	if !_fbfg {
		return nil, _e.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062je\u0063\u0074")
	}
	if _eace.Len() != 4 {
		return nil, _e.Errorf("\u0073\u0065p\u0061\u0072\u0061\u0074i\u006f\u006e \u0043\u0053\u003a\u0020\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006ce\u006e\u0067\u0074\u0068")
	}
	_abab = _eace.Get(0)
	_dabad, _fbfg := _abab.(*_abf.PdfObjectName)
	if !_fbfg {
		return nil, _e.Errorf("\u0073\u0065\u0070ar\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a \u0069n\u0076a\u006ci\u0064\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	if *_dabad != "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e" {
		return nil, _e.Errorf("\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0043\u0053\u003a\u0020w\u0072o\u006e\u0067\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u006e\u0061\u006d\u0065")
	}
	_abab = _eace.Get(1)
	_dabad, _fbfg = _abab.(*_abf.PdfObjectName)
	if !_fbfg {
		return nil, _e.Errorf("\u0073\u0065pa\u0072\u0061\u0074i\u006f\u006e\u0020\u0043S: \u0049nv\u0061\u006c\u0069\u0064\u0020\u0063\u006flo\u0072\u0061\u006e\u0074\u0020\u006e\u0061m\u0065")
	}
	_caddd.ColorantName = _dabad
	_abab = _eace.Get(2)
	_bbeb, _bbgf := NewPdfColorspaceFromPdfObject(_abab)
	if _bbgf != nil {
		return nil, _bbgf
	}
	_caddd.AlternateSpace = _bbeb
	_bbdd, _bbgf := _ebedg(_eace.Get(3))
	if _bbgf != nil {
		return nil, _bbgf
	}
	_caddd.TintTransform = _bbdd
	return _caddd, nil
}

// NewPdfColorspaceSpecialSeparation returns a new separation color.
func NewPdfColorspaceSpecialSeparation() *PdfColorspaceSpecialSeparation {
	_aebf := &PdfColorspaceSpecialSeparation{}
	return _aebf
}

// String returns a string representation of PdfTransformParamsDocMDP.
func (_fcbbe *PdfTransformParamsDocMDP) String() string {
	return _e.Sprintf("\u0025\u0073\u0020\u0050\u003a\u0020\u0025\u0073\u0020V\u003a\u0020\u0025\u0073", _fcbbe.Type, _fcbbe.P, _fcbbe.V)
}

// NewPdfPageResourcesColorspaces returns a new PdfPageResourcesColorspaces object.
func NewPdfPageResourcesColorspaces() *PdfPageResourcesColorspaces {
	_agage := &PdfPageResourcesColorspaces{}
	_agage.Names = []string{}
	_agage.Colorspaces = map[string]PdfColorspace{}
	_agage._cebc = &_abf.PdfIndirectObject{}
	return _agage
}

// ToPdfObject implements interface PdfModel.
func (_cafb *PdfAnnotationWidget) ToPdfObject() _abf.PdfObject {
	_cafb.PdfAnnotation.ToPdfObject()
	_aaadc := _cafb._dbc
	_dbd := _aaadc.PdfObject.(*_abf.PdfObjectDictionary)
	if _cafb._gbga {
		return _aaadc
	}
	_cafb._gbga = true
	_dbd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0057\u0069\u0064\u0067\u0065\u0074"))
	_dbd.SetIfNotNil("\u0048", _cafb.H)
	_dbd.SetIfNotNil("\u004d\u004b", _cafb.MK)
	_dbd.SetIfNotNil("\u0041", _cafb.A)
	_dbd.SetIfNotNil("\u0041\u0041", _cafb.AA)
	_dbd.SetIfNotNil("\u0042\u0053", _cafb.BS)
	_fccd := _cafb.Parent
	if _cafb._agdc != nil {
		if _cafb._agdc._dgdc == _cafb._dbc {
			_cafb._agdc.ToPdfObject()
		}
		_fccd = _cafb._agdc.GetContainingPdfObject()
	}
	if _fccd != _aaadc {
		_dbd.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _fccd)
	}
	_cafb._gbga = false
	return _aaadc
}

// GetContentStreamObjs returns a slice of PDF objects containing the content
// streams of the page.
func (_bgfcg *PdfPage) GetContentStreamObjs() []_abf.PdfObject {
	if _bgfcg.Contents == nil {
		return nil
	}
	_fdfg := _abf.TraceToDirectObject(_bgfcg.Contents)
	if _agbc, _bdbgg := _fdfg.(*_abf.PdfObjectArray); _bdbgg {
		return _agbc.Elements()
	}
	return []_abf.PdfObject{_fdfg}
}

// NewPdfAnnotation3D returns a new 3d annotation.
func NewPdfAnnotation3D() *PdfAnnotation3D {
	_afbd := NewPdfAnnotation()
	_acdfe := &PdfAnnotation3D{}
	_acdfe.PdfAnnotation = _afbd
	_afbd.SetContext(_acdfe)
	return _acdfe
}

// NewPdfShadingType3 creates an empty shading type 3 dictionary.
func NewPdfShadingType3() *PdfShadingType3 {
	_caddb := &PdfShadingType3{}
	_caddb.PdfShading = &PdfShading{}
	_caddb.PdfShading._eabcgc = _abf.MakeIndirectObject(_abf.MakeDict())
	_caddb.PdfShading._eabd = _caddb
	return _caddb
}

// NewPdfActionGoToR returns a new "go to remote" action.
func NewPdfActionGoToR() *PdfActionGoToR {
	_ed := NewPdfAction()
	_gag := &PdfActionGoToR{}
	_gag.PdfAction = _ed
	_ed.SetContext(_gag)
	return _gag
}

var _ pdfFont = (*pdfFontType3)(nil)

// GenerateHashMaps generates DSS hashmaps for Certificates, OCSPs and CRLs to make sure they are unique.
func (_eaeg *DSS) GenerateHashMaps() error {
	_degc, _ceac := _eaeg.generateHashMap(_eaeg.Certs)
	if _ceac != nil {
		return _ceac
	}
	_ceegg, _ceac := _eaeg.generateHashMap(_eaeg.OCSPs)
	if _ceac != nil {
		return _ceac
	}
	_adgc, _ceac := _eaeg.generateHashMap(_eaeg.CRLs)
	if _ceac != nil {
		return _ceac
	}
	_eaeg._gcee = _degc
	_eaeg._ggfg = _ceegg
	_eaeg._daee = _adgc
	return nil
}

// PdfAnnotationSquiggly represents Squiggly annotations.
// (Section 12.5.6.10).
type PdfAnnotationSquiggly struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _abf.PdfObject
}

// NewPdfFilespec returns an initialized generic PDF filespec model.
func NewPdfFilespec() *PdfFilespec {
	_fbga := &PdfFilespec{}
	_fbga._badbg = _abf.MakeIndirectObject(_abf.MakeDict())
	return _fbga
}

// PdfAnnotationMarkup represents additional fields for mark-up annotations.
// (Section 12.5.6.2 p. 399).
type PdfAnnotationMarkup struct {
	T            _abf.PdfObject
	Popup        *PdfAnnotationPopup
	CA           _abf.PdfObject
	RC           _abf.PdfObject
	CreationDate _abf.PdfObject
	IRT          _abf.PdfObject
	Subj         _abf.PdfObject
	RT           _abf.PdfObject
	IT           _abf.PdfObject
	ExData       _abf.PdfObject
}

// ImageToRGB converts Lab colorspace image to RGB and returns the result.
func (_acaf *PdfColorspaceLab) ImageToRGB(img Image) (Image, error) {
	_afba := func(_befd float64) float64 {
		if _befd >= 6.0/29 {
			return _befd * _befd * _befd
		}
		return 108.0 / 841 * (_befd - 4.0/29.0)
	}
	_dbbe := img._ceeag
	if len(_dbbe) != 6 {
		_acd.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u002d\u0020\u004c\u0061\u0062\u0020\u0044e\u0063\u006f\u0064\u0065\u0020\u0072\u0061\u006e\u0067e\u0020\u0021\u003d\u0020\u0036\u002e\u002e\u002e\u0020\u0075\u0073\u0065\u0020\u005b0\u0020\u0031\u0030\u0030\u0020\u0061\u006d\u0069\u006e\u0020\u0061\u006d\u0061\u0078\u0020\u0062\u006d\u0069\u006e\u0020\u0062\u006d\u0061\u0078\u005d\u0020\u0064\u0065\u0066\u0061u\u006c\u0074\u0020\u0064\u0065\u0063\u006f\u0064\u0065 \u0061\u0072r\u0061\u0079")
		_dbbe = _acaf.DecodeArray()
	}
	_gfccb := _gf.NewReader(img.getBase())
	_cdgb := _gca.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, img._gedg, img._ceeag)
	_dfaea := _gf.NewWriter(_cdgb)
	_efff := _ge.Pow(2, float64(img.BitsPerComponent)) - 1
	_ebcda := make([]uint32, 3)
	var (
		_eafe                                             error
		Ls, As, Bs, L, M, N, X, Y, Z, _dcbg, _efaf, _gbad float64
	)
	for {
		_eafe = _gfccb.ReadSamples(_ebcda)
		if _eafe == _gc.EOF {
			break
		} else if _eafe != nil {
			return img, _eafe
		}
		Ls = float64(_ebcda[0]) / _efff
		As = float64(_ebcda[1]) / _efff
		Bs = float64(_ebcda[2]) / _efff
		Ls = _gca.LinearInterpolate(Ls, 0.0, 1.0, _dbbe[0], _dbbe[1])
		As = _gca.LinearInterpolate(As, 0.0, 1.0, _dbbe[2], _dbbe[3])
		Bs = _gca.LinearInterpolate(Bs, 0.0, 1.0, _dbbe[4], _dbbe[5])
		L = (Ls+16)/116 + As/500
		M = (Ls + 16) / 116
		N = (Ls+16)/116 - Bs/200
		X = _acaf.WhitePoint[0] * _afba(L)
		Y = _acaf.WhitePoint[1] * _afba(M)
		Z = _acaf.WhitePoint[2] * _afba(N)
		_dcbg = 3.240479*X + -1.537150*Y + -0.498535*Z
		_efaf = -0.969256*X + 1.875992*Y + 0.041556*Z
		_gbad = 0.055648*X + -0.204043*Y + 1.057311*Z
		_dcbg = _ge.Min(_ge.Max(_dcbg, 0), 1.0)
		_efaf = _ge.Min(_ge.Max(_efaf, 0), 1.0)
		_gbad = _ge.Min(_ge.Max(_gbad, 0), 1.0)
		_ebcda[0] = uint32(_dcbg * _efff)
		_ebcda[1] = uint32(_efaf * _efff)
		_ebcda[2] = uint32(_gbad * _efff)
		if _eafe = _dfaea.WriteSamples(_ebcda); _eafe != nil {
			return img, _eafe
		}
	}
	return _cega(&_cdgb), nil
}

// Width returns the width of `rect`.
func (_ccbge *PdfRectangle) Width() float64 { return _ge.Abs(_ccbge.Urx - _ccbge.Llx) }

// ColorToRGB converts an Indexed color to an RGB color.
func (_cfee *PdfColorspaceSpecialIndexed) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _cfee.Base == nil {
		return nil, _fd.New("\u0069\u006e\u0064\u0065\u0078\u0065d\u0020\u0062\u0061\u0073\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _cfee.Base.ColorToRGB(color)
}

type pdfFontType3 struct {
	fontCommon
	_baee *_abf.PdfIndirectObject

	// These fields are specific to Type 3 fonts.
	CharProcs  _abf.PdfObject
	Encoding   _abf.PdfObject
	FontBBox   _abf.PdfObject
	FontMatrix _abf.PdfObject
	FirstChar  _abf.PdfObject
	LastChar   _abf.PdfObject
	Widths     _abf.PdfObject
	Resources  _abf.PdfObject
	_ecgf      map[_cbb.CharCode]float64
	_dgbd      _cbb.TextEncoder
}

func (_cggec *pdfFontType0) bytesToCharcodes(_aebeb []byte) ([]_cbb.CharCode, bool) {
	if _cggec._fcfg == nil {
		return nil, false
	}
	_accc, _bfae := _cggec._fcfg.BytesToCharcodes(_aebeb)
	if !_bfae {
		return nil, false
	}
	_bdac := make([]_cbb.CharCode, len(_accc))
	for _agbbg, _addbd := range _accc {
		_bdac[_agbbg] = _cbb.CharCode(_addbd)
	}
	return _bdac, true
}

// NewPermissions returns a new permissions object.
func NewPermissions(docMdp *PdfSignature) *Permissions {
	_gdead := Permissions{}
	_gdead.DocMDP = docMdp
	_fcdgb := _abf.MakeDict()
	_fcdgb.Set("\u0044\u006f\u0063\u004d\u0044\u0050", docMdp.ToPdfObject())
	_gdead._deefb = _fcdgb
	return &_gdead
}

// NewPdfColorspaceSpecialIndexed returns a new Indexed color.
func NewPdfColorspaceSpecialIndexed() *PdfColorspaceSpecialIndexed {
	return &PdfColorspaceSpecialIndexed{HiVal: 255}
}

// NewPdfWriter initializes a new PdfWriter.
func NewPdfWriter() PdfWriter {
	_cffff := PdfWriter{}
	_cffff._fdgae = map[_abf.PdfObject]struct{}{}
	_cffff._edcgc = []_abf.PdfObject{}
	_cffff._fadb = map[_abf.PdfObject][]*_abf.PdfObjectDictionary{}
	_cffff._dbdcg = map[_abf.PdfObject]struct{}{}
	_cffff._ecfa.Major = 1
	_cffff._ecfa.Minor = 3
	_dgdcd := _abf.MakeDict()
	_eccfc := []struct {
		_ggac  _abf.PdfObjectName
		_egaeb string
	}{{"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _gacgg()}, {"\u0043r\u0065\u0061\u0074\u006f\u0072", _aacbg()}, {"\u0041\u0075\u0074\u0068\u006f\u0072", _bdfef()}, {"\u0053u\u0062\u006a\u0065\u0063\u0074", _ebgb()}, {"\u0054\u0069\u0074l\u0065", _edcbb()}, {"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _efcef()}}
	for _, _facd := range _eccfc {
		if _facd._egaeb != "" {
			_dgdcd.Set(_facd._ggac, _abf.MakeString(_facd._egaeb))
		}
	}
	if _efgfe := _dgdfd(); !_efgfe.IsZero() {
		if _fegb, _ccdf := NewPdfDateFromTime(_efgfe); _ccdf == nil {
			_dgdcd.Set("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _fegb.ToPdfObject())
		}
	}
	if _afceb := _fcfeb(); !_afceb.IsZero() {
		if _becd, _egead := NewPdfDateFromTime(_afceb); _egead == nil {
			_dgdcd.Set("\u004do\u0064\u0044\u0061\u0074\u0065", _becd.ToPdfObject())
		}
	}
	_cebaea := _abf.PdfIndirectObject{}
	_cebaea.PdfObject = _dgdcd
	_cffff._ddegc = &_cebaea
	_cffff.addObject(&_cebaea)
	_bgagd := _abf.PdfIndirectObject{}
	_agcdd := _abf.MakeDict()
	_agcdd.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0043a\u0074\u0061\u006c\u006f\u0067"))
	_bgagd.PdfObject = _agcdd
	_cffff._cfdde = &_bgagd
	_cffff.addObject(_cffff._cfdde)
	_edca, _befgb := _addec("\u0077")
	if _befgb != nil {
		_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _befgb)
	}
	_cffff._ceega = _edca
	_cgcbc := _abf.PdfIndirectObject{}
	_defdc := _abf.MakeDict()
	_defdc.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0050\u0061\u0067e\u0073"))
	_ecgcg := _abf.PdfObjectArray{}
	_defdc.Set("\u004b\u0069\u0064\u0073", &_ecgcg)
	_defdc.Set("\u0043\u006f\u0075n\u0074", _abf.MakeInteger(0))
	_cgcbc.PdfObject = _defdc
	_cffff._cgeed = &_cgcbc
	_cffff._aadb = map[_abf.PdfObject]struct{}{}
	_cffff.addObject(_cffff._cgeed)
	_agcdd.Set("\u0050\u0061\u0067e\u0073", &_cgcbc)
	_cffff._ddffc = _agcdd
	_acd.Log.Trace("\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0025\u0073", _bgagd)
	return _cffff
}

// ColorFromFloats returns a new PdfColor based on input color components.
func (_gadb *PdfColorspaceDeviceN) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != _gadb.GetNumComponents() {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fgbff, _cecff := _gadb.TintTransform.Evaluate(vals)
	if _cecff != nil {
		return nil, _cecff
	}
	_bbda, _cecff := _gadb.AlternateSpace.ColorFromFloats(_fgbff)
	if _cecff != nil {
		return nil, _cecff
	}
	return _bbda, nil
}

func (_fceb *PdfReader) newPdfFieldFromIndirectObject(_faabd *_abf.PdfIndirectObject, _adgg *PdfField) (*PdfField, error) {
	if _dffd, _ddcgc := _fceb._ceecd.GetModelFromPrimitive(_faabd).(*PdfField); _ddcgc {
		return _dffd, nil
	}
	_bdgb, _gcfa := _abf.GetDict(_faabd)
	if !_gcfa {
		return nil, _e.Errorf("\u0050\u0064f\u0046\u0069\u0065\u006c\u0064 \u0069\u006e\u0064\u0069\u0072e\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_aeafb := NewPdfField()
	_aeafb._dgdc = _faabd
	_aeafb._dgdc.PdfObject = _bdgb
	if _bagcc, _bdeb := _abf.GetName(_bdgb.Get("\u0046\u0054")); _bdeb {
		_aeafb.FT = _bagcc
	}
	if _adgg != nil {
		_aeafb.Parent = _adgg
	}
	_aeafb.T, _ = _bdgb.Get("\u0054").(*_abf.PdfObjectString)
	_aeafb.TU, _ = _bdgb.Get("\u0054\u0055").(*_abf.PdfObjectString)
	_aeafb.TM, _ = _bdgb.Get("\u0054\u004d").(*_abf.PdfObjectString)
	_aeafb.Ff, _ = _bdgb.Get("\u0046\u0066").(*_abf.PdfObjectInteger)
	_aeafb.V = _bdgb.Get("\u0056")
	_aeafb.DV = _bdgb.Get("\u0044\u0056")
	_aeafb.AA = _bdgb.Get("\u0041\u0041")
	if DA := _bdgb.Get("\u0044\u0041"); DA != nil {
		DA, _ := _abf.GetString(DA)
		_aeafb.VariableText = &VariableText{DA: DA}
		Q, _ := _bdgb.Get("\u0051").(*_abf.PdfObjectInteger)
		DS, _ := _bdgb.Get("\u0044\u0053").(*_abf.PdfObjectString)
		RV := _bdgb.Get("\u0052\u0056")
		_aeafb.VariableText.Q = Q
		_aeafb.VariableText.DS = DS
		_aeafb.VariableText.RV = RV
	}
	_gbecf := _aeafb.FT
	if _gbecf == nil && _adgg != nil {
		_gbecf = _adgg.FT
	}
	if _gbecf != nil {
		switch *_gbecf {
		case "\u0054\u0078":
			_edab, _abec := _bfdc(_bdgb)
			if _abec != nil {
				return nil, _abec
			}
			_edab.PdfField = _aeafb
			_aeafb._ffea = _edab
		case "\u0043\u0068":
			_gbaad, _aaga := _cbea(_bdgb)
			if _aaga != nil {
				return nil, _aaga
			}
			_gbaad.PdfField = _aeafb
			_aeafb._ffea = _gbaad
		case "\u0042\u0074\u006e":
			_ffaa, _ggbb := _dgbgc(_bdgb)
			if _ggbb != nil {
				return nil, _ggbb
			}
			_ffaa.PdfField = _aeafb
			_aeafb._ffea = _ffaa
		case "\u0053\u0069\u0067":
			_cgff, _fcdag := _fceb.newPdfFieldSignatureFromDict(_bdgb)
			if _fcdag != nil {
				return nil, _fcdag
			}
			_cgff.PdfField = _aeafb
			_aeafb._ffea = _cgff
		default:
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", *_aeafb.FT)
			return nil, _fd.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074\u0079p\u0065")
		}
	}
	if _dddda, _fbcf := _abf.GetName(_bdgb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fbcf {
		if *_dddda == "\u0057\u0069\u0064\u0067\u0065\u0074" {
			_ggga, _gecd := _fceb.newPdfAnnotationFromIndirectObject(_faabd)
			if _gecd != nil {
				return nil, _gecd
			}
			_debf, _gdddb := _ggga.GetContext().(*PdfAnnotationWidget)
			if !_gdddb {
				return nil, _fd.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0077\u0069\u0064\u0067e\u0074 \u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006fn")
			}
			_debf._agdc = _aeafb
			_debf.Parent = _aeafb._dgdc
			_aeafb.Annotations = append(_aeafb.Annotations, _debf)
			return _aeafb, nil
		}
	}
	_bfged := true
	if _ecgb, _bbcd := _abf.GetArray(_bdgb.Get("\u004b\u0069\u0064\u0073")); _bbcd {
		_affc := make([]*_abf.PdfIndirectObject, 0, _ecgb.Len())
		for _, _fefc := range _ecgb.Elements() {
			_bebe, _afcd := _abf.GetIndirect(_fefc)
			if !_afcd {
				_cbba, _aeaa := _abf.GetStream(_fefc)
				if _aeaa && _cbba.PdfObjectDictionary != nil {
					_gade, _dagdf := _abf.GetNameVal(_cbba.Get("\u0054\u0079\u0070\u0065"))
					if _dagdf && _gade == "\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061" {
						_acd.Log.Debug("E\u0052RO\u0052:\u0020f\u006f\u0072\u006d\u0020\u0066i\u0065\u006c\u0064 \u004b\u0069\u0064\u0073\u0020a\u0072\u0072\u0061y\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0069n\u0076\u0061\u006cid \u004d\u0065\u0074\u0061\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e")
						continue
					}
				}
				return nil, _fd.New("n\u006f\u0074\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0028\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065\u006cd\u0029")
			}
			_bfea, _ecf := _abf.GetDict(_bebe)
			if !_ecf {
				return nil, ErrTypeCheck
			}
			if _bfged {
				_bfged = !_gdcaf(_bfea)
			}
			_affc = append(_affc, _bebe)
		}
		for _, _aecfa := range _affc {
			if _bfged {
				_gabb, _dedab := _fceb.newPdfAnnotationFromIndirectObject(_aecfa)
				if _dedab != nil {
					_acd.Log.Debug("\u0045r\u0072\u006fr\u0020\u006c\u006fa\u0064\u0069\u006e\u0067\u0020\u0077\u0069d\u0067\u0065\u0074\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0072 \u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _dedab)
					return nil, _dedab
				}
				_edeed, _begd := _gabb._edg.(*PdfAnnotationWidget)
				if !_begd {
					return nil, ErrTypeCheck
				}
				_edeed._agdc = _aeafb
				_aeafb.Annotations = append(_aeafb.Annotations, _edeed)
			} else {
				_fcabg, _cfbba := _fceb.newPdfFieldFromIndirectObject(_aecfa, _aeafb)
				if _cfbba != nil {
					_acd.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0063\u0068\u0069\u006c\u0064\u0020\u0066\u0069\u0065\u006c\u0064: \u0025\u0076", _cfbba)
					return nil, _cfbba
				}
				_aeafb.Kids = append(_aeafb.Kids, _fcabg)
			}
		}
	}
	return _aeafb, nil
}

// ToInteger convert to an integer format.
func (_ecgc *PdfColorDeviceGray) ToInteger(bits int) uint32 {
	_addef := _ge.Pow(2, float64(bits)) - 1
	return uint32(_addef * _ecgc.Val())
}

// ToPdfObject converts rectangle to a PDF object.
func (_cdgfb *PdfRectangle) ToPdfObject() _abf.PdfObject {
	return _abf.MakeArray(_abf.MakeFloat(_cdgfb.Llx), _abf.MakeFloat(_cdgfb.Lly), _abf.MakeFloat(_cdgfb.Urx), _abf.MakeFloat(_cdgfb.Ury))
}

func (_gbeed *XObjectImage) getParamsDict() *_abf.PdfObjectDictionary {
	_face := _abf.MakeDict()
	_face.Set("\u0057\u0069\u0064t\u0068", _abf.MakeInteger(*_gbeed.Width))
	_face.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _abf.MakeInteger(*_gbeed.Height))
	_face.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _abf.MakeInteger(int64(_gbeed.ColorSpace.GetNumComponents())))
	_face.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _abf.MakeInteger(*_gbeed.BitsPerComponent))
	return _face
}

// NewPdfFieldSignature returns an initialized signature field.
func NewPdfFieldSignature(signature *PdfSignature) *PdfFieldSignature {
	_ffec := &PdfFieldSignature{}
	_ffec.PdfField = NewPdfField()
	_ffec.PdfField.SetContext(_ffec)
	_ffec.PdfAnnotationWidget = NewPdfAnnotationWidget()
	_ffec.PdfAnnotationWidget.SetContext(_ffec)
	_ffec.PdfAnnotationWidget._dbc = _ffec.PdfField._dgdc
	_ffec.T = _abf.MakeString("")
	_ffec.F = _abf.MakeInteger(132)
	_ffec.V = signature
	return _ffec
}

// Compress is yet to be implemented.
// Should be able to compress in terms of JPEG quality parameter,
// and DPI threshold (need to know bounding area dimensions).
func (_fddg DefaultImageHandler) Compress(input *Image, quality int64) (*Image, error) {
	return input, nil
}

// String returns a string representation of the field.
func (_ccca *PdfField) String() string {
	if _ffbcb, _debeg := _ccca.ToPdfObject().(*_abf.PdfIndirectObject); _debeg {
		return _e.Sprintf("\u0025\u0054\u003a\u0020\u0025\u0073", _ccca._ffea, _ffbcb.PdfObject.String())
	}
	return ""
}

// GetPageDict converts the Page to a PDF object dictionary.
func (_cabed *PdfPage) GetPageDict() *_abf.PdfObjectDictionary {
	_gaae := _cabed._bdbfa
	_gaae.Clear()
	_gaae.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0050\u0061\u0067\u0065"))
	_gaae.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _cabed.Parent)
	if _cabed.LastModified != nil {
		_gaae.Set("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _cabed.LastModified.ToPdfObject())
	}
	if _cabed.Resources != nil {
		_gaae.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _cabed.Resources.ToPdfObject())
	}
	if _cabed.CropBox != nil {
		_gaae.Set("\u0043r\u006f\u0070\u0042\u006f\u0078", _cabed.CropBox.ToPdfObject())
	}
	if _cabed.MediaBox != nil {
		_gaae.Set("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078", _cabed.MediaBox.ToPdfObject())
	}
	if _cabed.BleedBox != nil {
		_gaae.Set("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078", _cabed.BleedBox.ToPdfObject())
	}
	if _cabed.TrimBox != nil {
		_gaae.Set("\u0054r\u0069\u006d\u0042\u006f\u0078", _cabed.TrimBox.ToPdfObject())
	}
	if _cabed.ArtBox != nil {
		_gaae.Set("\u0041\u0072\u0074\u0042\u006f\u0078", _cabed.ArtBox.ToPdfObject())
	}
	_gaae.SetIfNotNil("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f", _cabed.BoxColorInfo)
	_gaae.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _cabed.Contents)
	if _cabed.Rotate != nil {
		_gaae.Set("\u0052\u006f\u0074\u0061\u0074\u0065", _abf.MakeInteger(*_cabed.Rotate))
	}
	_gaae.SetIfNotNil("\u0047\u0072\u006fu\u0070", _cabed.Group)
	_gaae.SetIfNotNil("\u0054\u0068\u0075m\u0062", _cabed.Thumb)
	_gaae.SetIfNotNil("\u0042", _cabed.B)
	_gaae.SetIfNotNil("\u0044\u0075\u0072", _cabed.Dur)
	_gaae.SetIfNotNil("\u0054\u0072\u0061n\u0073", _cabed.Trans)
	_gaae.SetIfNotNil("\u0041\u0041", _cabed.AA)
	_gaae.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _cabed.Metadata)
	_gaae.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _cabed.PieceInfo)
	_gaae.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _cabed.StructParents)
	_gaae.SetIfNotNil("\u0049\u0044", _cabed.ID)
	_gaae.SetIfNotNil("\u0050\u005a", _cabed.PZ)
	_gaae.SetIfNotNil("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f", _cabed.SeparationInfo)
	_gaae.SetIfNotNil("\u0054\u0061\u0062\u0073", _cabed.Tabs)
	_gaae.SetIfNotNil("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064", _cabed.TemplateInstantiated)
	_gaae.SetIfNotNil("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s", _cabed.PresSteps)
	_gaae.SetIfNotNil("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074", _cabed.UserUnit)
	_gaae.SetIfNotNil("\u0056\u0050", _cabed.VP)
	if _cabed._baagf != nil {
		_decf := _abf.MakeArray()
		for _, _eegfg := range _cabed._baagf {
			if _bbfdf := _eegfg.GetContext(); _bbfdf != nil {
				_decf.Append(_bbfdf.ToPdfObject())
			} else {
				_decf.Append(_eegfg.ToPdfObject())
			}
		}
		if _decf.Len() > 0 {
			_gaae.Set("\u0041\u006e\u006e\u006f\u0074\u0073", _decf)
		}
	} else if _cabed.Annots != nil {
		_gaae.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0073", _cabed.Annots)
	}
	return _gaae
}

// ToPdfObject returns the PDF representation of the VRI dictionary.
func (_cdegg *VRI) ToPdfObject() *_abf.PdfObjectDictionary {
	_dbbgb := _abf.MakeDict()
	_dbbgb.SetIfNotNil(_abf.PdfObjectName("\u0043\u0065\u0072\u0074"), _fdaf(_cdegg.Cert))
	_dbbgb.SetIfNotNil(_abf.PdfObjectName("\u004f\u0043\u0053\u0050"), _fdaf(_cdegg.OCSP))
	_dbbgb.SetIfNotNil(_abf.PdfObjectName("\u0043\u0052\u004c"), _fdaf(_cdegg.CRL))
	_dbbgb.SetIfNotNil("\u0054\u0055", _cdegg.TU)
	_dbbgb.SetIfNotNil("\u0054\u0053", _cdegg.TS)
	return _dbbgb
}

// NewXObjectForm creates a brand new XObject Form. Creates a new underlying PDF object stream primitive.
func NewXObjectForm() *XObjectForm {
	_fbdac := &XObjectForm{}
	_gcccg := &_abf.PdfObjectStream{}
	_gcccg.PdfObjectDictionary = _abf.MakeDict()
	_fbdac._dbba = _gcccg
	return _fbdac
}

func (_acf *PdfReader) newPdfAnnotationStrikeOut(_gacd *_abf.PdfObjectDictionary) (*PdfAnnotationStrikeOut, error) {
	_eed := PdfAnnotationStrikeOut{}
	_aafb, _ggad := _acf.newPdfAnnotationMarkupFromDict(_gacd)
	if _ggad != nil {
		return nil, _ggad
	}
	_eed.PdfAnnotationMarkup = _aafb
	_eed.QuadPoints = _gacd.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_eed, nil
}

// String returns a string that describes `font`.
func (_ccbcf *PdfFont) String() string {
	_agecb := ""
	if _ccbcf._gedca.Encoder() != nil {
		_agecb = _ccbcf._gedca.Encoder().String()
	}
	return _e.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0054\u0020\u0025s\u0020\u0025\u0073\u007d", _ccbcf._gedca, _ccbcf.baseFields().coreString(), _agecb)
}

// ColorAt returns the color of the image pixel specified by the x and y coordinates.
func (_fggee *Image) ColorAt(x, y int) (_ga.Color, error) {
	_cdaba := _gca.BytesPerLine(int(_fggee.Width), int(_fggee.BitsPerComponent), _fggee.ColorComponents)
	switch _fggee.ColorComponents {
	case 1:
		return _gca.ColorAtGrayscale(x, y, int(_fggee.BitsPerComponent), _cdaba, _fggee.Data, _fggee._ceeag)
	case 3:
		return _gca.ColorAtNRGBA(x, y, int(_fggee.Width), _cdaba, int(_fggee.BitsPerComponent), _fggee.Data, _fggee._gedg, _fggee._ceeag)
	case 4:
		return _gca.ColorAtCMYK(x, y, int(_fggee.Width), _fggee.Data, _fggee._ceeag)
	}
	_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 i\u006da\u0067\u0065\u002e\u0020\u0025\u0064\u0020\u0063\u006f\u006d\u0070\u006fn\u0065\u006e\u0074\u0073\u002c\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072 \u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _fggee.ColorComponents, _fggee.BitsPerComponent)
	return nil, _fd.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061g\u0065 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
}

// Encoder returns the font's text encoder.
func (_cbcda pdfCIDFontType0) Encoder() _cbb.TextEncoder { return _cbcda._aefc }

// SetDSS sets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_aag *PdfAppender) SetDSS(dss *DSS) {
	if dss != nil {
		_aag.updateObjectsDeep(dss.ToPdfObject(), nil)
	}
	_aag._ffbe = dss
}

func _cbea(_efde *_abf.PdfObjectDictionary) (*PdfFieldChoice, error) {
	_ccdbc := &PdfFieldChoice{}
	_ccdbc.Opt, _ = _abf.GetArray(_efde.Get("\u004f\u0070\u0074"))
	_ccdbc.TI, _ = _abf.GetInt(_efde.Get("\u0054\u0049"))
	_ccdbc.I, _ = _abf.GetArray(_efde.Get("\u0049"))
	return _ccdbc, nil
}

// PdfFont represents an underlying font structure which can be of type:
// - Type0
// - Type1
// - TrueType
// etc.
type PdfFont struct{ _gedca pdfFont }

// UpdateXObjectImageFromImage creates a new XObject Image from an
// Image object `img` and default masks from xobjIn.
// The default masks are overridden if img.hasAlpha
// If `encoder` is nil, uses raw encoding (none).
func UpdateXObjectImageFromImage(xobjIn *XObjectImage, img *Image, cs PdfColorspace, encoder _abf.StreamEncoder) (*XObjectImage, error) {
	if encoder == nil {
		var _acgbe error
		encoder, _acgbe = img.getSuitableEncoder()
		if _acgbe != nil {
			_acd.Log.Debug("F\u0061\u0069l\u0075\u0072\u0065\u0020\u006f\u006e\u0020\u0066\u0069\u006e\u0064\u0069\u006e\u0067\u0020\u0073\u0075\u0069\u0074\u0061b\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072,\u0020\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b\u0020\u0074\u006f\u0020R\u0061\u0077\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _acgbe)
			encoder = _abf.NewRawEncoder()
		}
	}
	encoder.UpdateParams(img.GetParamsDict())
	_efgcc, _gcffe := encoder.EncodeBytes(img.Data)
	if _gcffe != nil {
		_acd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _gcffe)
		return nil, _gcffe
	}
	_geccb := NewXObjectImage()
	_dfdce := img.Width
	_adafd := img.Height
	_geccb.Width = &_dfdce
	_geccb.Height = &_adafd
	_cccea := img.BitsPerComponent
	_geccb.BitsPerComponent = &_cccea
	_geccb.Filter = encoder
	_geccb.Stream = _efgcc
	if cs == nil {
		if img.ColorComponents == 1 {
			_geccb.ColorSpace = NewPdfColorspaceDeviceGray()
			if img.BitsPerComponent == 16 {
				switch encoder.(type) {
				case *_abf.DCTEncoder:
					_geccb.ColorSpace = NewPdfColorspaceDeviceRGB()
					_cccea = 8
					_geccb.BitsPerComponent = &_cccea
				}
			}
		} else if img.ColorComponents == 3 {
			_geccb.ColorSpace = NewPdfColorspaceDeviceRGB()
		} else if img.ColorComponents == 4 {
			switch encoder.(type) {
			case *_abf.DCTEncoder:
				_geccb.ColorSpace = NewPdfColorspaceDeviceRGB()
			default:
				_geccb.ColorSpace = NewPdfColorspaceDeviceCMYK()
			}
		} else {
			return nil, _fd.New("c\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020u\u006e\u0064\u0065\u0066in\u0065\u0064")
		}
	} else {
		_geccb.ColorSpace = cs
	}
	if len(img._gedg) != 0 {
		_acadfe := NewXObjectImage()
		_acadfe.Filter = encoder
		_efdfb, _dcdcg := encoder.EncodeBytes(img._gedg)
		if _dcdcg != nil {
			_acd.Log.Debug("\u0045\u0072\u0072or\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0076", _dcdcg)
			return nil, _dcdcg
		}
		_acadfe.Stream = _efdfb
		_acadfe.BitsPerComponent = _geccb.BitsPerComponent
		_acadfe.Width = &img.Width
		_acadfe.Height = &img.Height
		_acadfe.ColorSpace = NewPdfColorspaceDeviceGray()
		_geccb.SMask = _acadfe.ToPdfObject()
	} else {
		_geccb.SMask = xobjIn.SMask
		_geccb.ImageMask = xobjIn.ImageMask
		if _geccb.ColorSpace.GetNumComponents() == 1 {
			_bffaa(_geccb)
		}
	}
	return _geccb, nil
}

// String returns string value of output intent for given type
// ISO_19005-2 6.2.3: GTS_PDFA1 value should be used for PDF/A-1, A-2 and A-3 at least
func (_gbcd PdfOutputIntentType) String() string {
	switch _gbcd {
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

// PdfColorspaceSpecialSeparation is a Separation colorspace.
// At the moment the colour space is set to a Separation space, the conforming reader shall determine whether the
// device has an available colorant (e.g. dye) corresponding to the name of the requested space. If so, the conforming
// reader shall ignore the alternateSpace and tintTransform parameters; subsequent painting operations within the
// space shall apply the designated colorant directly, according to the tint values supplied.
//
// Format: [/Separation name alternateSpace tintTransform]
type PdfColorspaceSpecialSeparation struct {
	ColorantName   *_abf.PdfObjectName
	AlternateSpace PdfColorspace
	TintTransform  PdfFunction
	_bbed          *_abf.PdfIndirectObject
}

func (_dedb *PdfReader) newPdfAnnotationLinkFromDict(_dcbe *_abf.PdfObjectDictionary) (*PdfAnnotationLink, error) {
	_eefb := PdfAnnotationLink{}
	_eefb.A = _dcbe.Get("\u0041")
	_eefb.Dest = _dcbe.Get("\u0044\u0065\u0073\u0074")
	_eefb.H = _dcbe.Get("\u0048")
	_eefb.PA = _dcbe.Get("\u0050\u0041")
	_eefb.QuadPoints = _dcbe.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_eefb.BS = _dcbe.Get("\u0042\u0053")
	return &_eefb, nil
}

// NewOutlineDest returns a new outline destination which can be used
// with outline items.
func NewOutlineDest(page int64, x, y float64) OutlineDest {
	return OutlineDest{Page: page, Mode: "\u0058\u0059\u005a", X: x, Y: y}
}

// GetAction returns the PDF action for the annotation link.
func (_ege *PdfAnnotationLink) GetAction() (*PdfAction, error) {
	if _ege._bgad != nil {
		return _ege._bgad, nil
	}
	if _ege.A == nil {
		return nil, nil
	}
	if _ege._aefa == nil {
		return nil, nil
	}
	_effg, _afe := _ege._aefa.loadAction(_ege.A)
	if _afe != nil {
		return nil, _afe
	}
	_ege._bgad = _effg
	return _ege._bgad, nil
}

// ColorFromPdfObjects returns a new PdfColor based on input color components. The input PdfObjects should
// be numeric.
func (_faebc *PdfColorspaceDeviceN) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != _faebc.GetNumComponents() {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_eccgd, _gafcc := _abf.GetNumbersAsFloat(objects)
	if _gafcc != nil {
		return nil, _gafcc
	}
	return _faebc.ColorFromFloats(_eccgd)
}

func (_fgbde *PdfReader) buildNameNodes(_aagfe *_abf.PdfIndirectObject, _beef map[_abf.PdfObject]struct{}) error {
	if _aagfe == nil {
		return nil
	}
	if _, _dbgcf := _beef[_aagfe]; _dbgcf {
		_acd.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _aagfe.ObjectNumber)
		return nil
	}
	_beef[_aagfe] = struct{}{}
	_bfccg, _cgada := _aagfe.PdfObject.(*_abf.PdfObjectDictionary)
	if !_cgada {
		return _fd.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if _ebggd, _bdbdc := _abf.GetDict(_bfccg.Get("\u0044\u0065\u0073t\u0073")); _bdbdc {
		_ffdba, _ccbdfc := _abf.GetArray(_ebggd.Get("\u004b\u0069\u0064\u0073"))
		if !_ccbdfc {
			return _fd.New("\u0049n\u0076\u0061\u006c\u0069d\u0020\u004b\u0069\u0064\u0073 \u0061r\u0072a\u0079\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_acd.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _ffdba)
		for _aeea, _gffag := range _ffdba.Elements() {
			_afdab, _faeab := _abf.GetIndirect(_gffag)
			if !_faeab {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u0068\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002d \u0028\u0025\u0073\u0029", _afdab)
				return _fd.New("\u0063h\u0069\u006c\u0064\u0020n\u006f\u0074\u0020\u0069\u006ed\u0069r\u0065c\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
			}
			_ffdba.Set(_aeea, _afdab)
			_eaec := _fgbde.buildNameNodes(_afdab, _beef)
			if _eaec != nil {
				return _eaec
			}
		}
	}
	if _facf, _adeeg := _abf.GetDict(_bfccg); _adeeg {
		if !_abf.IsNullObject(_facf.Get("\u004b\u0069\u0064\u0073")) {
			if _fbgfa, _adbgc := _abf.GetArray(_facf.Get("\u004b\u0069\u0064\u0073")); _adbgc {
				for _agcgb, _agfcc := range _fbgfa.Elements() {
					if _gcaf, _ffcd := _abf.GetIndirect(_agfcc); _ffcd {
						_fbgfa.Set(_agcgb, _gcaf)
						_ddgba := _fgbde.buildNameNodes(_gcaf, _beef)
						if _ddgba != nil {
							return _ddgba
						}
					}
				}
			}
		}
	}
	return nil
}

func (_cfca *PdfReader) newPdfAnnotationRichMediaFromDict(_dff *_abf.PdfObjectDictionary) (*PdfAnnotationRichMedia, error) {
	_gaggg := &PdfAnnotationRichMedia{}
	_gaggg.RichMediaSettings = _dff.Get("\u0052\u0069\u0063\u0068\u004d\u0065\u0064\u0069\u0061\u0053\u0065\u0074t\u0069\u006e\u0067\u0073")
	_gaggg.RichMediaContent = _dff.Get("\u0052\u0069c\u0068\u004d\u0065d\u0069\u0061\u0043\u006f\u006e\u0074\u0065\u006e\u0074")
	return _gaggg, nil
}

// ButtonType represents the subtype of a button field, can be one of:
// - Checkbox (ButtonTypeCheckbox)
// - PushButton (ButtonTypePushButton)
// - RadioButton (ButtonTypeRadioButton)
type ButtonType int

func (_bffg *PdfReader) newPdfAnnotationFileAttachmentFromDict(_bcaa *_abf.PdfObjectDictionary) (*PdfAnnotationFileAttachment, error) {
	_cgd := PdfAnnotationFileAttachment{}
	_gbd, _eade := _bffg.newPdfAnnotationMarkupFromDict(_bcaa)
	if _eade != nil {
		return nil, _eade
	}
	_cgd.PdfAnnotationMarkup = _gbd
	_cgd.FS = _bcaa.Get("\u0046\u0053")
	_cgd.Name = _bcaa.Get("\u004e\u0061\u006d\u0065")
	return &_cgd, nil
}

// AddExtension adds the specified extension to the Extensions dictionary.
// See section 7.1.2 "Extensions Dictionary" (pp. 108-109 PDF32000_2008).
func (_cbcddb *PdfWriter) AddExtension(extName, baseVersion string, extLevel int) {
	_dcgcg, _abeed := _abf.GetDict(_cbcddb._ddffc.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073"))
	if !_abeed {
		_dcgcg = _abf.MakeDict()
		_cbcddb._ddffc.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073", _dcgcg)
	}
	_bdgbg, _abeed := _abf.GetDict(_dcgcg.Get(_abf.PdfObjectName(extName)))
	if !_abeed {
		_bdgbg = _abf.MakeDict()
		_dcgcg.Set(_abf.PdfObjectName(extName), _bdgbg)
	}
	if _dbcdd, _ := _abf.GetNameVal(_bdgbg.Get("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e")); _dbcdd != baseVersion {
		_bdgbg.Set("B\u0061\u0073\u0065\u0056\u0065\u0072\u0073\u0069\u006f\u006e", _abf.MakeName(baseVersion))
	}
	if _ecbb, _ := _abf.GetIntVal(_bdgbg.Get("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c")); _ecbb != extLevel {
		_bdgbg.Set("\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006eL\u0065\u0076\u0065\u006c", _abf.MakeInteger(int64(extLevel)))
	}
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// component PDF objects.
func (_bcgc *PdfColorspaceICCBased) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if _bcgc.Alternate == nil {
		if _bcgc.N == 1 {
			_cadd := NewPdfColorspaceDeviceGray()
			return _cadd.ColorFromPdfObjects(objects)
		} else if _bcgc.N == 3 {
			_gbgg := NewPdfColorspaceDeviceRGB()
			return _gbgg.ColorFromPdfObjects(objects)
		} else if _bcgc.N == 4 {
			_feea := NewPdfColorspaceDeviceCMYK()
			return _feea.ColorFromPdfObjects(objects)
		} else {
			return nil, _fd.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _bcgc.Alternate.ColorFromPdfObjects(objects)
}

// ToPdfObject converts the font to a PDF representation.
func (_gedb *pdfFontType3) ToPdfObject() _abf.PdfObject {
	if _gedb._baee == nil {
		_gedb._baee = &_abf.PdfIndirectObject{}
	}
	_aefee := _gedb.baseFields().asPdfObjectDictionary("\u0054\u0079\u0070e\u0033")
	_gedb._baee.PdfObject = _aefee
	if _gedb.FirstChar != nil {
		_aefee.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _gedb.FirstChar)
	}
	if _gedb.LastChar != nil {
		_aefee.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _gedb.LastChar)
	}
	if _gedb.Widths != nil {
		_aefee.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _gedb.Widths)
	}
	if _gedb.Encoding != nil {
		_aefee.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _gedb.Encoding)
	} else if _gedb._dgbd != nil {
		_abbe := _gedb._dgbd.ToPdfObject()
		if _abbe != nil {
			_aefee.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _abbe)
		}
	}
	if _gedb.FontBBox != nil {
		_aefee.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _gedb.FontBBox)
	}
	if _gedb.FontMatrix != nil {
		_aefee.Set("\u0046\u006f\u006e\u0074\u004d\u0061\u0074\u0069\u0072\u0078", _gedb.FontMatrix)
	}
	if _gedb.CharProcs != nil {
		_aefee.Set("\u0043h\u0061\u0072\u0050\u0072\u006f\u0063s", _gedb.CharProcs)
	}
	if _gedb.Resources != nil {
		_aefee.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _gedb.Resources)
	}
	return _gedb._baee
}

func (_eefe *pdfFontSimple) updateStandard14Font() {
	_fded, _baadfd := _eefe.Encoder().(_cbb.SimpleEncoder)
	if !_baadfd {
		_acd.Log.Error("\u0057\u0072\u006f\u006e\u0067\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074y\u0070e\u003a\u0020\u0025\u0054\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073\u002e", _eefe.Encoder(), _eefe)
		return
	}
	_bgfdd := _fded.Charcodes()
	_eefe._aadgb = make(map[_cbb.CharCode]float64, len(_bgfdd))
	for _, _fggfa := range _bgfdd {
		_cbdbb, _ := _fded.CharcodeToRune(_fggfa)
		_eecde, _ := _eefe._aecd.Read(_cbdbb)
		_eefe._aadgb[_fggfa] = _eecde.Wx
	}
}

// ToPdfObject implements interface PdfModel.
func (_bdf *PdfAnnotationUnderline) ToPdfObject() _abf.PdfObject {
	_bdf.PdfAnnotation.ToPdfObject()
	_gafg := _bdf._dbc
	_bebg := _gafg.PdfObject.(*_abf.PdfObjectDictionary)
	_bdf.PdfAnnotationMarkup.appendToPdfDictionary(_bebg)
	_bebg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0055n\u0064\u0065\u0072\u006c\u0069\u006ee"))
	_bebg.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _bdf.QuadPoints)
	return _gafg
}

// PdfColorspaceDeviceNAttributes contains additional information about the components of colour space that
// conforming readers may use. Conforming readers need not use the alternateSpace and tintTransform parameters,
// and may instead use a custom blending algorithms, along with other information provided in the attributes
// dictionary if present.
type PdfColorspaceDeviceNAttributes struct {
	Subtype     *_abf.PdfObjectName
	Colorants   _abf.PdfObject
	Process     _abf.PdfObject
	MixingHints _abf.PdfObject
	_ddbdd      *_abf.PdfIndirectObject
}

func (_dgdae *LTV) getCerts(_agfg []*_fa.Certificate) ([][]byte, error) {
	_cbffd := make([][]byte, 0, len(_agfg))
	for _, _fdda := range _agfg {
		_cbffd = append(_cbffd, _fdda.Raw)
	}
	return _cbffd, nil
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 1 for a CalGray device.
func (_cead *PdfColorspaceCalGray) GetNumComponents() int { return 1 }

func (_aacfe *pdfFontSimple) addEncoding() error {
	var (
		_fbab  string
		_gaggc map[_cbb.CharCode]_cbb.GlyphName
		_gbfc  _cbb.SimpleEncoder
	)
	if _aacfe.Encoder() != nil {
		_eedede, _agacf := _aacfe.Encoder().(_cbb.SimpleEncoder)
		if _agacf && _eedede != nil {
			_fbab = _eedede.BaseName()
		}
	}
	if _aacfe.Encoding != nil {
		_abgaeb, _caaa, _daffa := _aacfe.getFontEncoding()
		if _daffa != nil {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042\u0061\u0073\u0065F\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0053u\u0062t\u0079\u0070\u0065\u003d\u0025\u0071\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003d\u0025\u0073 \u0028\u0025\u0054\u0029\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _aacfe._ecggf, _aacfe._aacbc, _aacfe.Encoding, _aacfe.Encoding, _daffa)
			return _daffa
		}
		if _abgaeb != "" {
			_fbab = _abgaeb
		}
		_gaggc = _caaa
		_gbfc, _daffa = _cbb.NewSimpleTextEncoder(_fbab, _gaggc)
		if _daffa != nil {
			return _daffa
		}
	}
	if _gbfc == nil {
		_gbbcg := _aacfe._dcbaf
		if _gbbcg != nil {
			switch _aacfe._aacbc {
			case "\u0054\u0079\u0070e\u0031":
				if _gbbcg.fontFile != nil && _gbbcg.fontFile._eedb != nil {
					_acd.Log.Debug("\u0055\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006et\u0046\u0069\u006c\u0065")
					_gbfc = _gbbcg.fontFile._eedb
				}
			case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
				if _gbbcg._fcdf != nil {
					_acd.Log.Debug("\u0055s\u0069n\u0067\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032")
					_bfcf, _fdag := _gbbcg._fcdf.MakeEncoder()
					if _fdag == nil {
						_gbfc = _bfcf
					}
					if _aacfe._aabfe == nil {
						_aacfe._aabfe = _gbbcg._fcdf.MakeToUnicode()
					}
				}
			}
		}
	}
	if _gbfc != nil {
		if _gaggc != nil {
			_acd.Log.Trace("\u0064\u0069\u0066fe\u0072\u0065\u006e\u0063\u0065\u0073\u003d\u0025\u002b\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gaggc, _aacfe.baseFields())
			_gbfc = _cbb.ApplyDifferences(_gbfc, _gaggc)
		}
		_aacfe.SetEncoder(_gbfc)
	}
	return nil
}

// ToPdfObject returns a PdfObject representation of PdfColorspaceDeviceNAttributes as a PdfObjectDictionary directly
// or indirectly within an indirect object container.
func (_fcaa *PdfColorspaceDeviceNAttributes) ToPdfObject() _abf.PdfObject {
	_gcbcd := _abf.MakeDict()
	if _fcaa.Subtype != nil {
		_gcbcd.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _fcaa.Subtype)
	}
	_gcbcd.SetIfNotNil("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s", _fcaa.Colorants)
	_gcbcd.SetIfNotNil("\u0050r\u006f\u0063\u0065\u0073\u0073", _fcaa.Process)
	_gcbcd.SetIfNotNil("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073", _fcaa.MixingHints)
	if _fcaa._ddbdd != nil {
		_fcaa._ddbdd.PdfObject = _gcbcd
		return _fcaa._ddbdd
	}
	return _gcbcd
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_ccebg *PdfColorspaceSpecialPattern) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _ccebg.UnderlyingCS == nil {
		return nil, _fd.New("u\u006e\u0064\u0065\u0072\u006c\u0079i\u006e\u0067\u0020\u0043\u0053\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
	}
	return _ccebg.UnderlyingCS.ColorFromFloats(vals)
}

// FlattenFieldsWithOpts flattens the AcroForm fields of the page using the
// provided field appearance generator and the specified options. If no options
// are specified, all form fields are flattened for the page.
// If a filter function is provided using the opts parameter, only the filtered
// fields are flattened. Otherwise, all form fields are flattened.
func (_cbbc *PdfPage) FlattenFieldsWithOpts(appgen FieldAppearanceGenerator, opts *FieldFlattenOpts) error {
	_ddcbd := map[*PdfAnnotation]bool{}
	_fbbb, _begdc := _cbbc.GetAnnotations()
	if _begdc != nil {
		return _begdc
	}
	_gdde := false
	for _, _bdgf := range _fbbb {
		if opts.AnnotFilterFunc != nil {
			_ddcbd[_bdgf] = opts.AnnotFilterFunc(_bdgf)
		} else {
			_ddcbd[_bdgf] = true
		}
		if _ddcbd[_bdgf] {
			_gdde = true
		}
	}
	if !_gdde {
		return nil
	}
	return _cbbc.flattenFieldsWithOpts(appgen, opts, _ddcbd)
}

// AlphaMapFunc represents a alpha mapping function: byte -> byte. Can be used for
// thresholding the alpha channel, i.e. setting all alpha values below threshold to transparent.
type AlphaMapFunc func(_edda byte) byte

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

func (_cae *PdfReader) newPdfActionSubmitFormFromDict(_beb *_abf.PdfObjectDictionary) (*PdfActionSubmitForm, error) {
	_bfc, _dbe := _dgf(_beb.Get("\u0046"))
	if _dbe != nil {
		return nil, _dbe
	}
	return &PdfActionSubmitForm{F: _bfc, Fields: _beb.Get("\u0046\u0069\u0065\u006c\u0064\u0073"), Flags: _beb.Get("\u0046\u006c\u0061g\u0073")}, nil
}

// ToPdfObject implements interface PdfModel.
func (_eaa *PdfActionURI) ToPdfObject() _abf.PdfObject {
	_eaa.PdfAction.ToPdfObject()
	_fb := _eaa._egg
	_dg := _fb.PdfObject.(*_abf.PdfObjectDictionary)
	_dg.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeURI)))
	_dg.SetIfNotNil("\u0055\u0052\u0049", _eaa.URI)
	_dg.SetIfNotNil("\u0049\u0073\u004da\u0070", _eaa.IsMap)
	return _fb
}

func (_fccf *PdfReader) newPdfAnnotationProjectionFromDict(_gaff *_abf.PdfObjectDictionary) (*PdfAnnotationProjection, error) {
	_dggac := &PdfAnnotationProjection{}
	_ece, _fcea := _fccf.newPdfAnnotationMarkupFromDict(_gaff)
	if _fcea != nil {
		return nil, _fcea
	}
	_dggac.PdfAnnotationMarkup = _ece
	return _dggac, nil
}

// AddAnnotation appends `annot` to the list of page annotations.
func (_bbcda *PdfPage) AddAnnotation(annot *PdfAnnotation) {
	if _bbcda._baagf == nil {
		_bbcda.GetAnnotations()
	}
	_bbcda._baagf = append(_bbcda._baagf, annot)
}

// PdfAnnotationText represents Text annotations.
// (Section 12.5.6.4 p. 402).
type PdfAnnotationText struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Open       _abf.PdfObject
	Name       _abf.PdfObject
	State      _abf.PdfObject
	StateModel _abf.PdfObject
}

// PdfBorderEffect represents a PDF border effect.
type PdfBorderEffect struct {
	S *BorderEffect
	I *float64
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a CalRGB device.
func (_cgdg *PdfColorspaceCalRGB) GetNumComponents() int { return 3 }

// ToPdfObject convert PdfInfo to pdf object.
func (_ddge *PdfInfo) ToPdfObject() _abf.PdfObject {
	_fead := _abf.MakeDict()
	_fead.SetIfNotNil("\u0054\u0069\u0074l\u0065", _ddge.Title)
	_fead.SetIfNotNil("\u0041\u0075\u0074\u0068\u006f\u0072", _ddge.Author)
	_fead.SetIfNotNil("\u0053u\u0062\u006a\u0065\u0063\u0074", _ddge.Subject)
	_fead.SetIfNotNil("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _ddge.Keywords)
	_fead.SetIfNotNil("\u0043r\u0065\u0061\u0074\u006f\u0072", _ddge.Creator)
	_fead.SetIfNotNil("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _ddge.Producer)
	_fead.SetIfNotNil("\u0054r\u0061\u0070\u0070\u0065\u0064", _ddge.Trapped)
	if _ddge.CreationDate != nil {
		_fead.SetIfNotNil("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _ddge.CreationDate.ToPdfObject())
	}
	if _ddge.ModifiedDate != nil {
		_fead.SetIfNotNil("\u004do\u0064\u0044\u0061\u0074\u0065", _ddge.ModifiedDate.ToPdfObject())
	}
	for _, _cafc := range _ddge._cbf.Keys() {
		_fead.SetIfNotNil(_cafc, _ddge._cbf.Get(_cafc))
	}
	return _fead
}

// ParserMetadata gets the parser  metadata.
func (_ccdbe *CompliancePdfReader) ParserMetadata() _abf.ParserMetadata {
	if _ccdbe._fcgbc == (_abf.ParserMetadata{}) {
		_ccdbe._fcgbc, _ = _ccdbe._bebc.ParserMetadata()
	}
	return _ccdbe._fcgbc
}

// GetContainingPdfObject returns the container of the pattern object (indirect object).
func (_afec *PdfPattern) GetContainingPdfObject() _abf.PdfObject { return _afec._bcfca }

func _aagc(_egacb _abf.PdfObject) (*PdfPattern, error) {
	_eceeg := &PdfPattern{}
	var _cfbagd *_abf.PdfObjectDictionary
	if _fagf, _beeab := _abf.GetIndirect(_egacb); _beeab {
		_eceeg._bcfca = _fagf
		_ecbec, _deaac := _fagf.PdfObject.(*_abf.PdfObjectDictionary)
		if !_deaac {
			_acd.Log.Debug("\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0028g\u006f\u0074\u0020%\u0054\u0029", _fagf.PdfObject)
			return nil, _abf.ErrTypeError
		}
		_cfbagd = _ecbec
	} else if _bfcgb, _fbffe := _abf.GetStream(_egacb); _fbffe {
		_eceeg._bcfca = _bfcgb
		_cfbagd = _bfcgb.PdfObjectDictionary
	} else {
		_acd.Log.Debug("\u0050a\u0074\u0074e\u0072\u006e\u0020\u006eo\u0074\u0020\u0061n\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 o\u0062\u006a\u0065c\u0074\u0020o\u0072\u0020\u0073\u0074\u0072\u0065a\u006d\u002e \u0025\u0054", _egacb)
		return nil, _abf.ErrTypeError
	}
	_afcbg := _cfbagd.Get("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065")
	if _afcbg == nil {
		_acd.Log.Debug("\u0050\u0064\u0066\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069n\u0067\u0020\u0050\u0061\u0074t\u0065\u0072n\u0054\u0079\u0070\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fcgbe, _ccegef := _afcbg.(*_abf.PdfObjectInteger)
	if !_ccegef {
		_acd.Log.Debug("\u0050\u0061tt\u0065\u0072\u006e \u0074\u0079\u0070\u0065 no\u0074 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067er\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _afcbg)
		return nil, _abf.ErrTypeError
	}
	if *_fcgbe != 1 && *_fcgbe != 2 {
		_acd.Log.Debug("\u0050\u0061\u0074\u0074e\u0072\u006e\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d \u0031/\u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", *_fcgbe)
		return nil, _abf.ErrRangeError
	}
	_eceeg.PatternType = int64(*_fcgbe)
	switch *_fcgbe {
	case 1:
		_aafee, _baccb := _dfaag(_cfbagd)
		if _baccb != nil {
			return nil, _baccb
		}
		_aafee.PdfPattern = _eceeg
		_eceeg._bgafe = _aafee
		return _eceeg, nil
	case 2:
		_dgacd, _cggfd := _dfdga(_cfbagd)
		if _cggfd != nil {
			return nil, _cggfd
		}
		_dgacd.PdfPattern = _eceeg
		_eceeg._bgafe = _dgacd
		return _eceeg, nil
	}
	return nil, _fd.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
}

func _dbde(_acba *XObjectForm) (*PdfRectangle, bool, error) {
	if _bdbg, _gfeb := _acba.BBox.(*_abf.PdfObjectArray); _gfeb {
		_dagf, _badc := NewPdfRectangle(*_bdbg)
		if _badc != nil {
			return nil, false, _badc
		}
		if _cfeea, _ccce := _acba.Matrix.(*_abf.PdfObjectArray); _ccce {
			_bead, _eaebe := _cfeea.ToFloat64Array()
			if _eaebe != nil {
				return nil, false, _eaebe
			}
			_bdbb := _ad.IdentityMatrix()
			if len(_bead) == 6 {
				_bdbb = _ad.NewMatrix(_bead[0], _bead[1], _bead[2], _bead[3], _bead[4], _bead[5])
			}
			_dagf.Transform(_bdbb)
			return _dagf, true, nil
		}
		return _dagf, false, nil
	}
	return nil, false, _fd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063e\u0020\u0042\u0042\u006f\u0078\u0020\u0074y\u0070\u0065")
}

// HasFontByName checks whether a font is defined by the specified keyName.
func (_fbcg *PdfPageResources) HasFontByName(keyName _abf.PdfObjectName) bool {
	_, _ffdbg := _fbcg.GetFontByName(keyName)
	return _ffdbg
}

// PdfAnnotationProjection represents Projection annotations.
type PdfAnnotationProjection struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
}

func (_aggge *PdfReader) loadForms() (*PdfAcroForm, error) {
	if _aggge._bebc.GetCrypter() != nil && !_aggge._bebc.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_cbcfc := _aggge._dagde
	_gecfg := _cbcfc.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d")
	if _gecfg == nil {
		return nil, nil
	}
	_ecceg, _aedeg := _abf.GetIndirect(_gecfg)
	_gecfg = _abf.TraceToDirectObject(_gecfg)
	if _abf.IsNullObject(_gecfg) {
		_acd.Log.Trace("\u0041\u0063\u0072of\u006f\u0072\u006d\u0020\u0069\u0073\u0020\u0061\u0020n\u0075l\u006c \u006fb\u006a\u0065\u0063\u0074\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029\u000a")
		return nil, nil
	}
	_ffeef, _bebef := _abf.GetDict(_gecfg)
	if !_bebef {
		_acd.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0041\u0063\u0072\u006fF\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _gecfg)
		_acd.Log.Debug("\u0044\u006f\u0065\u0073 n\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0073")
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0061\u0063\u0072\u006ff\u006fr\u006d \u0065\u006e\u0074\u0072\u0079\u0020\u0025T", _gecfg)
	}
	_acd.Log.Trace("\u0048\u0061\u0073\u0020\u0041\u0063\u0072\u006f\u0020f\u006f\u0072\u006d\u0073")
	_acd.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0041\u0063r\u006ff\u006f\u0072\u006d\u0073\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065")
	if !_aggge._abgge {
		_abccd := _aggge.traverseObjectData(_ffeef)
		if _abccd != nil {
			_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0074\u0072a\u0076\u0065\u0072\u0073\u0065\u0020\u0041\u0063\u0072\u006fFo\u0072\u006d\u0073 \u0028%\u0073\u0029", _abccd)
			return nil, _abccd
		}
	}
	_fcbdbd, _gbcad := _aggge.newPdfAcroFormFromDict(_ecceg, _ffeef)
	if _gbcad != nil {
		return nil, _gbcad
	}
	_fcbdbd._dfebf = !_aedeg
	return _fcbdbd, nil
}

// GetContainingPdfObject returns the container of the shading object (indirect object).
func (_fdfec *PdfShading) GetContainingPdfObject() _abf.PdfObject { return _fdfec._eabcgc }

// ToPdfObject implements interface PdfModel.
func (_cgea *PdfActionImportData) ToPdfObject() _abf.PdfObject {
	_cgea.PdfAction.ToPdfObject()
	_bee := _cgea._egg
	_gfdc := _bee.PdfObject.(*_abf.PdfObjectDictionary)
	_gfdc.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeImportData)))
	if _cgea.F != nil {
		_gfdc.Set("\u0046", _cgea.F.ToPdfObject())
	}
	return _bee
}

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
func (_ebbba *PdfFont) GetCharMetrics(code _cbb.CharCode) (CharMetrics, bool) {
	var _cdbe _gbe.CharMetrics
	switch _dbdc := _ebbba._gedca.(type) {
	case *pdfFontSimple:
		if _ccddd, _bcaac := _dbdc.GetCharMetrics(code); _bcaac {
			return _ccddd, _bcaac
		}
	case *pdfFontType0:
		if _bedbf, _cfge := _dbdc.GetCharMetrics(code); _cfge {
			return _bedbf, _cfge
		}
	case *pdfCIDFontType0:
		if _cgcba, _gbba := _dbdc.GetCharMetrics(code); _gbba {
			return _cgcba, _gbba
		}
	case *pdfCIDFontType2:
		if _bcgea, _bacgf := _dbdc.GetCharMetrics(code); _bacgf {
			return _bcgea, _bacgf
		}
	case *pdfFontType3:
		if _daff, _efbcb := _dbdc.GetCharMetrics(code); _efbcb {
			return _daff, _efbcb
		}
	default:
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020G\u0065\u0074\u0043h\u0061\u0072\u004de\u0074\u0072i\u0063\u0073\u0020\u006e\u006f\u0074 \u0069mp\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d\u0025\u0054\u002e", _ebbba._gedca)
		return _cdbe, false
	}
	if _adeg, _ddgbc := _ebbba.GetFontDescriptor(); _ddgbc == nil && _adeg != nil {
		return _gbe.CharMetrics{Wx: _adeg._fgccc}, true
	}
	_acd.Log.Debug("\u0047\u0065\u0074\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069c\u0073\u0020\u0066\u006f\u0072 \u0066\u006fn\u0074\u003d\u0025\u0073", _ebbba)
	return _cdbe, false
}

// ToPdfObject implements interface PdfModel.
func (_fac *PdfAnnotationFreeText) ToPdfObject() _abf.PdfObject {
	_fac.PdfAnnotation.ToPdfObject()
	_dafd := _fac._dbc
	_abga := _dafd.PdfObject.(*_abf.PdfObjectDictionary)
	_fac.PdfAnnotationMarkup.appendToPdfDictionary(_abga)
	_abga.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0046\u0072\u0065\u0065\u0054\u0065\u0078\u0074"))
	_abga.SetIfNotNil("\u0044\u0041", _fac.DA)
	_abga.SetIfNotNil("\u0051", _fac.Q)
	_abga.SetIfNotNil("\u0052\u0043", _fac.RC)
	_abga.SetIfNotNil("\u0044\u0053", _fac.DS)
	_abga.SetIfNotNil("\u0043\u004c", _fac.CL)
	_abga.SetIfNotNil("\u0049\u0054", _fac.IT)
	_abga.SetIfNotNil("\u0042\u0045", _fac.BE)
	_abga.SetIfNotNil("\u0052\u0044", _fac.RD)
	_abga.SetIfNotNil("\u0042\u0053", _fac.BS)
	_abga.SetIfNotNil("\u004c\u0045", _fac.LE)
	return _dafd
}

func (_fccdae *PdfReader) resolveReference(_bbedc *_abf.PdfObjectReference) (_abf.PdfObject, bool, error) {
	_gcfac, _cfcab := _fccdae._bebc.ObjCache[int(_bbedc.ObjectNumber)]
	if !_cfcab {
		_acd.Log.Trace("R\u0065\u0061\u0064\u0065r \u004co\u006f\u006b\u0075\u0070\u0020r\u0065\u0066\u003a\u0020\u0025\u0073", _bbedc)
		_dcbac, _daggg := _fccdae._bebc.LookupByReference(*_bbedc)
		if _daggg != nil {
			return nil, false, _daggg
		}
		_fccdae._bebc.ObjCache[int(_bbedc.ObjectNumber)] = _dcbac
		return _dcbac, false, nil
	}
	return _gcfac, true, nil
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
	ShadingType *_abf.PdfObjectInteger
	ColorSpace  PdfColorspace
	Background  *_abf.PdfObjectArray
	BBox        *PdfRectangle
	AntiAlias   *_abf.PdfObjectBool
	_eabd       PdfModel
	_eabcgc     _abf.PdfObject
}

// ColorToRGB converts a CalGray color to an RGB color.
func (_edfff *PdfColorspaceCalGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_egfa, _gcbab := color.(*PdfColorCalGray)
	if !_gcbab {
		_acd.Log.Debug("\u0049n\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006eo\u0074\u0020\u0063\u0061\u006c\u0020\u0067\u0072\u0061\u0079")
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	ANorm := _egfa.Val()
	X := _edfff.WhitePoint[0] * _ge.Pow(ANorm, _edfff.Gamma)
	Y := _edfff.WhitePoint[1] * _ge.Pow(ANorm, _edfff.Gamma)
	Z := _edfff.WhitePoint[2] * _ge.Pow(ANorm, _edfff.Gamma)
	_dea := 3.240479*X + -1.537150*Y + -0.498535*Z
	_bdfaf := -0.969256*X + 1.875992*Y + 0.041556*Z
	_gecee := 0.055648*X + -0.204043*Y + 1.057311*Z
	_dea = _ge.Min(_ge.Max(_dea, 0), 1.0)
	_bdfaf = _ge.Min(_ge.Max(_bdfaf, 0), 1.0)
	_gecee = _ge.Min(_ge.Max(_gecee, 0), 1.0)
	return NewPdfColorDeviceRGB(_dea, _bdfaf, _gecee), nil
}

// GetContainingPdfObject returns the container of the PdfAcroForm (indirect object).
func (_fcagf *PdfAcroForm) GetContainingPdfObject() _abf.PdfObject { return _fcagf._bgfc }

// GetExtGState gets the ExtGState specified by keyName. Returns a bool
// indicating whether it was found or not.
func (_gdgb *PdfPageResources) GetExtGState(keyName _abf.PdfObjectName) (_abf.PdfObject, bool) {
	if _gdgb.ExtGState == nil {
		return nil, false
	}
	_cbbdc, _bbbdb := _abf.TraceToDirectObject(_gdgb.ExtGState).(*_abf.PdfObjectDictionary)
	if !_bbbdb {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gdgb.ExtGState)
		return nil, false
	}
	if _cded := _cbbdc.Get(keyName); _cded != nil {
		return _cded, true
	}
	return nil, false
}

var (
	CourierName              = _gbe.CourierName
	CourierBoldName          = _gbe.CourierBoldName
	CourierObliqueName       = _gbe.CourierObliqueName
	CourierBoldObliqueName   = _gbe.CourierBoldObliqueName
	HelveticaName            = _gbe.HelveticaName
	HelveticaBoldName        = _gbe.HelveticaBoldName
	HelveticaObliqueName     = _gbe.HelveticaObliqueName
	HelveticaBoldObliqueName = _gbe.HelveticaBoldObliqueName
	SymbolName               = _gbe.SymbolName
	ZapfDingbatsName         = _gbe.ZapfDingbatsName
	TimesRomanName           = _gbe.TimesRomanName
	TimesBoldName            = _gbe.TimesBoldName
	TimesItalicName          = _gbe.TimesItalicName
	TimesBoldItalicName      = _gbe.TimesBoldItalicName
)

// StandardApplier is the interface that performs optimization of the whole PDF document.
// As a result an input document is being changed by the optimizer.
// The writer than takes back all it's parts and overwrites it.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardApplier interface {
	ApplyStandard(_ebfbfg *_bbf.Document) error
}

// NewLTV returns a new LTV client.
func NewLTV(appender *PdfAppender) (*LTV, error) {
	_aabce := appender.Reader.DSS
	if _aabce == nil {
		_aabce = NewDSS()
	}
	if _befgc := _aabce.GenerateHashMaps(); _befgc != nil {
		return nil, _befgc
	}
	return &LTV{CertClient: _fe.NewCertClient(), OCSPClient: _fe.NewOCSPClient(), CRLClient: _fe.NewCRLClient(), SkipExisting: true, _bfed: appender, _dgfe: _aabce}, nil
}

// SetPdfAuthor sets the Author attribute of the output PDF.
func SetPdfAuthor(author string) { _gaabd.Lock(); defer _gaabd.Unlock(); _efdg = author }

// GetAscent returns the Ascent of the font `descriptor`.
func (_cgcf *PdfFontDescriptor) GetAscent() (float64, error) {
	return _abf.GetNumberAsFloat(_cgcf.Ascent)
}

func _dfdga(_fdged *_abf.PdfObjectDictionary) (*PdfShadingPattern, error) {
	_bgdeb := &PdfShadingPattern{}
	_ceacg := _fdged.Get("\u0053h\u0061\u0064\u0069\u006e\u0067")
	if _ceacg == nil {
		_acd.Log.Debug("\u0053h\u0061d\u0069\u006e\u0067\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_fdgd, _fbgab := _abaef(_ceacg)
	if _fbgab != nil {
		_acd.Log.Debug("\u0045r\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006eg\u0020s\u0068a\u0064\u0069\u006e\u0067\u003a\u0020\u0025v", _fbgab)
		return nil, _fbgab
	}
	_bgdeb.Shading = _fdgd
	if _cegad := _fdged.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _cegad != nil {
		_ebfe, _aceged := _cegad.(*_abf.PdfObjectArray)
		if !_aceged {
			_acd.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _cegad)
			return nil, _abf.ErrTypeError
		}
		_bgdeb.Matrix = _ebfe
	}
	if _daaca := _fdged.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _daaca != nil {
		_bgdeb.ExtGState = _daaca
	}
	return _bgdeb, nil
}

func (_daddf *LTV) validateSig(_bgga *PdfSignature) error {
	if _bgga == nil || _bgga.Contents == nil || len(_bgga.Contents.Bytes()) == 0 {
		return _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0066\u0069\u0065l\u0064:\u0020\u0025\u0076", _bgga)
	}
	return nil
}

func _fccda(_cbceg *_abf.PdfObjectDictionary, _fabb *fontCommon) (*pdfCIDFontType2, error) {
	if _fabb._aacbc != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_acd.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0032\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _fabb)
		return nil, _abf.ErrRangeError
	}
	_ddad := _gbbbf(_fabb)
	_gecb, _dfcg := _abf.GetDict(_cbceg.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_dfcg {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _fabb)
		return nil, ErrRequiredAttributeMissing
	}
	_ddad.CIDSystemInfo = _gecb
	_ddad.DW = _cbceg.Get("\u0044\u0057")
	_ddad.W = _cbceg.Get("\u0057")
	_ddad.DW2 = _cbceg.Get("\u0044\u0057\u0032")
	_ddad.W2 = _cbceg.Get("\u0057\u0032")
	_ddad.CIDToGIDMap = _cbceg.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070")
	_ddad._cecdg = 1000.0
	if _cfbfd, _ceaeab := _abf.GetNumberAsFloat(_ddad.DW); _ceaeab == nil {
		_ddad._cecdg = _cfbfd
	}
	_bcdaf, _cfgf := _fecf(_ddad.W)
	if _cfgf != nil {
		return nil, _cfgf
	}
	if _bcdaf == nil {
		_bcdaf = map[_cbb.CharCode]float64{}
	}
	_ddad._ddeea = _bcdaf
	return _ddad, nil
}

func (_faea *PdfReader) newPdfActionImportDataFromDict(_agea *_abf.PdfObjectDictionary) (*PdfActionImportData, error) {
	_gagg, _aad := _dgf(_agea.Get("\u0046"))
	if _aad != nil {
		return nil, _aad
	}
	return &PdfActionImportData{F: _gagg}, nil
}

var _ pdfFont = (*pdfFontSimple)(nil)

// C returns the value of the cyan component of the color.
func (_beda *PdfColorDeviceCMYK) C() float64 { return _beda[0] }

func (_dffdg *pdfFontSimple) getFontEncoding() (_gecba string, _fbcff map[_cbb.CharCode]_cbb.GlyphName, _gfecb error) {
	_gecba = "\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"
	if _baca, _acffd := _gcgde[_dffdg._ecggf]; _acffd {
		_gecba = _baca
	} else if _dffdg.fontFlags()&_eceag != 0 {
		for _fdbgc, _gbbg := range _gcgde {
			if _be.Contains(_dffdg._ecggf, _fdbgc) {
				_gecba = _gbbg
				break
			}
		}
	}
	if _dffdg.Encoding == nil {
		return _gecba, nil, nil
	}
	switch _eefbe := _dffdg.Encoding.(type) {
	case *_abf.PdfObjectName:
		return string(*_eefbe), nil, nil
	case *_abf.PdfObjectDictionary:
		_gada, _fabbf := _abf.GetName(_eefbe.Get("\u0042\u0061\u0073e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
		if _fabbf {
			_gecba = _gada.String()
		}
		if _addgac := _eefbe.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"); _addgac != nil {
			_deae, _edeae := _abf.GetArray(_addgac)
			if !_edeae {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0042a\u0064\u0020\u0066on\u0074\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u003d\u0025\u002b\u0076\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073=\u0025\u0054", _eefbe, _eefbe.Get("D\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0073"))
				return "", nil, _abf.ErrTypeError
			}
			_fbcff, _gfecb = _cbb.FromFontDifferences(_deae)
		}
		return _gecba, _fbcff, _gfecb
	default:
		_acd.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0072\u0020\u0064\u0069\u0063t\u0020\u0028\u0025\u0054\u0029\u0020\u0025\u0073", _dffdg.Encoding, _dffdg.Encoding)
		return "", nil, _abf.ErrTypeError
	}
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_gaca *PdfOutlineItem) ToPdfObject() _abf.PdfObject {
	_daaba := _gaca._ceegc
	_bbdc := _daaba.PdfObject.(*_abf.PdfObjectDictionary)
	_bbdc.Set("\u0054\u0069\u0074l\u0065", _gaca.Title)
	if _gaca.A != nil {
		_bbdc.Set("\u0041", _gaca.A)
	}
	if _afcb := _bbdc.Get("\u0053\u0045"); _afcb != nil {
		_bbdc.Remove("\u0053\u0045")
	}
	if _gaca.C != nil {
		_bbdc.Set("\u0043", _gaca.C)
	}
	if _gaca.Dest != nil {
		_bbdc.Set("\u0044\u0065\u0073\u0074", _gaca.Dest)
	}
	if _gaca.F != nil {
		_bbdc.Set("\u0046", _gaca.F)
	}
	if _gaca.Count != nil {
		_bbdc.Set("\u0043\u006f\u0075n\u0074", _abf.MakeInteger(*_gaca.Count))
	}
	if _gaca.Next != nil {
		_bbdc.Set("\u004e\u0065\u0078\u0074", _gaca.Next.ToPdfObject())
	}
	if _gaca.First != nil {
		_bbdc.Set("\u0046\u0069\u0072s\u0074", _gaca.First.ToPdfObject())
	}
	if _gaca.Prev != nil {
		_bbdc.Set("\u0050\u0072\u0065\u0076", _gaca.Prev.GetContext().GetContainingPdfObject())
	}
	if _gaca.Last != nil {
		_bbdc.Set("\u004c\u0061\u0073\u0074", _gaca.Last.GetContext().GetContainingPdfObject())
	}
	if _gaca.Parent != nil {
		_bbdc.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _gaca.Parent.GetContext().GetContainingPdfObject())
	}
	return _daaba
}

// PartialName returns the partial name of the field.
func (_fggd *PdfField) PartialName() string {
	_dbga := ""
	if _fggd.T != nil {
		_dbga = _fggd.T.Decoded()
	} else {
		_acd.Log.Debug("\u0046\u0069el\u0064\u0020\u006di\u0073\u0073\u0069\u006eg T\u0020fi\u0065\u006c\u0064\u0020\u0028\u0069\u006eco\u006d\u0070\u0061\u0074\u0069\u0062\u006ce\u0029")
	}
	return _dbga
}

func _acgge(_dadf _abf.PdfObject) (*PdfFunctionType3, error) {
	_gcbaf := &PdfFunctionType3{}
	var _dagg *_abf.PdfObjectDictionary
	if _cagba, _fafff := _dadf.(*_abf.PdfIndirectObject); _fafff {
		_dbdfd, _ecaa := _cagba.PdfObject.(*_abf.PdfObjectDictionary)
		if !_ecaa {
			return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_gcbaf._edacd = _cagba
		_dagg = _dbdfd
	} else if _ecbgg, _becff := _dadf.(*_abf.PdfObjectDictionary); _becff {
		_dagg = _ecbgg
	} else {
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_bddef, _ccfbb := _abf.TraceToDirectObject(_dagg.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_abf.PdfObjectArray)
	if !_ccfbb {
		_acd.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _bddef.Len() != 2 {
		_acd.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _fd.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_degdd, _ecgcbe := _bddef.ToFloat64Array()
	if _ecgcbe != nil {
		return nil, _ecgcbe
	}
	_gcbaf.Domain = _degdd
	_bddef, _ccfbb = _abf.TraceToDirectObject(_dagg.Get("\u0052\u0061\u006eg\u0065")).(*_abf.PdfObjectArray)
	if _ccfbb {
		if _bddef.Len() < 0 || _bddef.Len()%2 != 0 {
			return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_deaba, _cgfcf := _bddef.ToFloat64Array()
		if _cgfcf != nil {
			return nil, _cgfcf
		}
		_gcbaf.Range = _deaba
	}
	_bddef, _ccfbb = _abf.TraceToDirectObject(_dagg.Get("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es")).(*_abf.PdfObjectArray)
	if !_ccfbb {
		_acd.Log.Error("\u0046\u0075\u006ect\u0069\u006f\u006e\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_gcbaf.Functions = []PdfFunction{}
	for _, _cbgag := range _bddef.Elements() {
		_dedfa, _ggeec := _ebedg(_cbgag)
		if _ggeec != nil {
			return nil, _ggeec
		}
		_gcbaf.Functions = append(_gcbaf.Functions, _dedfa)
	}
	_bddef, _ccfbb = _abf.TraceToDirectObject(_dagg.Get("\u0042\u006f\u0075\u006e\u0064\u0073")).(*_abf.PdfObjectArray)
	if !_ccfbb {
		_acd.Log.Error("B\u006fu\u006e\u0064\u0073\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_egca, _ecgcbe := _bddef.ToFloat64Array()
	if _ecgcbe != nil {
		return nil, _ecgcbe
	}
	_gcbaf.Bounds = _egca
	if len(_gcbaf.Bounds) != len(_gcbaf.Functions)-1 {
		_acd.Log.Error("B\u006f\u0075\u006e\u0064\u0073\u0020\u0028\u0025\u0064)\u0020\u0061\u006e\u0064\u0020\u006e\u0075m \u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029 n\u006f\u0074 \u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067", len(_gcbaf.Bounds), len(_gcbaf.Functions))
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bddef, _ccfbb = _abf.TraceToDirectObject(_dagg.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_abf.PdfObjectArray)
	if !_ccfbb {
		_acd.Log.Error("E\u006ec\u006f\u0064\u0065\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_fgafe, _ecgcbe := _bddef.ToFloat64Array()
	if _ecgcbe != nil {
		return nil, _ecgcbe
	}
	_gcbaf.Encode = _fgafe
	if len(_gcbaf.Encode) != 2*len(_gcbaf.Functions) {
		_acd.Log.Error("\u004c\u0065\u006e\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0028\u0025\u0064\u0029 \u0061\u006e\u0064\u0020\u006e\u0075\u006d\u0020\u0066\u0075\u006e\u0063\u0074i\u006f\u006e\u0073\u0020\u0028\u0025\u0064\u0029\u0020\u006e\u006f\u0074 m\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u0075\u0070", len(_gcbaf.Encode), len(_gcbaf.Functions))
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	return _gcbaf, nil
}

// GetContainingPdfObject returns the container of the outline tree node (indirect object).
func (_bfaea *PdfOutlineTreeNode) GetContainingPdfObject() _abf.PdfObject {
	return _bfaea.GetContext().GetContainingPdfObject()
}

func (_cdbfc *Image) samplesAddPadding(_dbgd []uint32) []uint32 {
	_ecdg := _gca.BytesPerLine(int(_cdbfc.Width), int(_cdbfc.BitsPerComponent), _cdbfc.ColorComponents) * (8 / int(_cdbfc.BitsPerComponent))
	_gdbdd := _ecdg * int(_cdbfc.Height)
	if len(_dbgd) == _gdbdd {
		return _dbgd
	}
	_gfffed := make([]uint32, _gdbdd)
	_fegcf := int(_cdbfc.Width) * _cdbfc.ColorComponents
	for _dcde := 0; _dcde < int(_cdbfc.Height); _dcde++ {
		_daced := _dcde * int(_cdbfc.Width)
		_ggebd := _dcde * _ecdg
		for _bfga := 0; _bfga < _fegcf; _bfga++ {
			_gfffed[_ggebd+_bfga] = _dbgd[_daced+_bfga]
		}
	}
	return _gfffed
}

// EnableChain adds the specified certificate chain and validation data (OCSP
// and CRL information) for it to the global scope of the document DSS. The
// added data is used for validating any of the signatures present in the
// document. The LTV client attempts to build the certificate chain up to a
// trusted root by downloading any missing certificates.
func (_adbbc *LTV) EnableChain(chain []*_fa.Certificate) error { return _adbbc.enable(nil, chain, "") }

// GetPageAsIndirectObject returns the page as a dictionary within an PdfIndirectObject.
func (_ceace *PdfPage) GetPageAsIndirectObject() *_abf.PdfIndirectObject { return _ceace._gefee }

// GetContainingPdfObject returns the container of the image object (indirect object).
func (_dfbbb *XObjectImage) GetContainingPdfObject() _abf.PdfObject { return _dfbbb._ccbad }

func (_aecbd *PdfWriter) adjustXRefAffectedVersion(_gabbb bool) {
	if _gabbb && _aecbd._ecfa.Major == 1 && _aecbd._ecfa.Minor < 5 {
		_aecbd._ecfa.Minor = 5
	}
}

func (_efggd *PdfWriter) writeXRefStreams(_cgeg int, _ffcfg int64) error {
	_bcafc := _cgeg + 1
	_efggd._becfc[_bcafc] = crossReference{Type: 1, ObjectNumber: _bcafc, Offset: _ffcfg}
	_edbbf := _dd.NewBuffer(nil)
	_badce := _abf.MakeArray()
	for _bcbe := 0; _bcbe <= _cgeg; {
		for ; _bcbe <= _cgeg; _bcbe++ {
			_gbgfc, _dedfe := _efggd._becfc[_bcbe]
			if _dedfe && (!_efggd._aegbd || _efggd._aegbd && (_gbgfc.Type == 1 && _gbgfc.Offset >= _efggd._cfecga || _gbgfc.Type == 0)) {
				break
			}
		}
		var _fdee int
		for _fdee = _bcbe + 1; _fdee <= _cgeg; _fdee++ {
			_bdcgd, _gdbad := _efggd._becfc[_fdee]
			if _gdbad && (!_efggd._aegbd || _efggd._aegbd && (_bdcgd.Type == 1 && _bdcgd.Offset > _efggd._cfecga)) {
				continue
			}
			break
		}
		_badce.Append(_abf.MakeInteger(int64(_bcbe)), _abf.MakeInteger(int64(_fdee-_bcbe)))
		for _bcfef := _bcbe; _bcfef < _fdee; _bcfef++ {
			_fedc := _efggd._becfc[_bcfef]
			switch _fedc.Type {
			case 0:
				_bg.Write(_edbbf, _bg.BigEndian, byte(0))
				_bg.Write(_edbbf, _bg.BigEndian, uint32(0))
				_bg.Write(_edbbf, _bg.BigEndian, uint16(0xFFFF))
			case 1:
				_bg.Write(_edbbf, _bg.BigEndian, byte(1))
				_bg.Write(_edbbf, _bg.BigEndian, uint32(_fedc.Offset))
				_bg.Write(_edbbf, _bg.BigEndian, uint16(_fedc.Generation))
			case 2:
				_bg.Write(_edbbf, _bg.BigEndian, byte(2))
				_bg.Write(_edbbf, _bg.BigEndian, uint32(_fedc.ObjectNumber))
				_bg.Write(_edbbf, _bg.BigEndian, uint16(_fedc.Index))
			}
		}
		_bcbe = _fdee + 1
	}
	_dbfdd, _badcdg := _abf.MakeStream(_edbbf.Bytes(), _abf.NewFlateEncoder())
	if _badcdg != nil {
		return _badcdg
	}
	_dbfdd.ObjectNumber = int64(_bcafc)
	_dbfdd.PdfObjectDictionary.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0058\u0052\u0065\u0066"))
	_dbfdd.PdfObjectDictionary.Set("\u0057", _abf.MakeArray(_abf.MakeInteger(1), _abf.MakeInteger(4), _abf.MakeInteger(2)))
	_dbfdd.PdfObjectDictionary.Set("\u0049\u006e\u0064e\u0078", _badce)
	_dbfdd.PdfObjectDictionary.Set("\u0053\u0069\u007a\u0065", _abf.MakeInteger(int64(_bcafc)))
	_dbfdd.PdfObjectDictionary.Set("\u0049\u006e\u0066\u006f", _efggd._ddegc)
	_dbfdd.PdfObjectDictionary.Set("\u0052\u006f\u006f\u0074", _efggd._cfdde)
	if _efggd._aegbd && _efggd._ffgf > 0 {
		_dbfdd.PdfObjectDictionary.Set("\u0050\u0072\u0065\u0076", _abf.MakeInteger(_efggd._ffgf))
	}
	if _efggd._ddbgd != nil {
		_dbfdd.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _efggd._dcdbb)
	}
	if _efggd._dedfdf == nil && _efggd._aefff != "" && _efggd._cfbce != "" {
		_efggd._dedfdf = _abf.MakeArray(_abf.MakeHexString(_efggd._aefff), _abf.MakeHexString(_efggd._cfbce))
	}
	if _efggd._dedfdf != nil {
		_acd.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _efggd._dedfdf)
		_dbfdd.Set("\u0049\u0044", _efggd._dedfdf)
	}
	_efggd.writeObject(int(_dbfdd.ObjectNumber), _dbfdd)
	return nil
}

// GetSubFilter returns SubFilter value or empty string.
func (_addgg *pdfSignDictionary) GetSubFilter() string {
	_ccfedf := _addgg.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r")
	if _ccfedf == nil {
		return ""
	}
	if _cggdg, _gaeeef := _abf.GetNameVal(_ccfedf); _gaeeef {
		return _cggdg
	}
	return ""
}

// ImageToRGB returns the passed in image. Method exists in order to satisfy
// the PdfColorspace interface.
func (_ggee *PdfColorspaceDeviceRGB) ImageToRGB(img Image) (Image, error) { return img, nil }

func _gggfec(_cbgaa _abf.PdfObject) []*_abf.PdfObjectStream {
	if _cbgaa == nil {
		return nil
	}
	_fbea, _ceagb := _abf.GetArray(_cbgaa)
	if !_ceagb || _fbea.Len() == 0 {
		return nil
	}
	_eebca := make([]*_abf.PdfObjectStream, 0, _fbea.Len())
	for _, _gdbb := range _fbea.Elements() {
		if _gdbfa, _gbag := _abf.GetStream(_gdbb); _gbag {
			_eebca = append(_eebca, _gdbfa)
		}
	}
	return _eebca
}

// GetNumComponents returns the number of color components (3 for CalRGB).
func (_dgec *PdfColorCalRGB) GetNumComponents() int { return 3 }

// CharcodesToUnicodeWithStats is identical to CharcodesToUnicode except it returns more statistical
// information about hits and misses from the reverse mapping process.
// NOTE: The number of runes returned may be greater than the number of charcodes.
// TODO(peterwilliams97): Deprecate in v4 and use only CharcodesToStrings()
func (_abecf *PdfFont) CharcodesToUnicodeWithStats(charcodes []_cbb.CharCode) (_gbca []rune, _addga, _agded int) {
	_aedgg, _addga, _agded := _abecf.CharcodesToStrings(charcodes)
	return []rune(_be.Join(_aedgg, "")), _addga, _agded
}

// SetForms sets the Acroform for a PDF file.
func (_ggcef *PdfWriter) SetForms(form *PdfAcroForm) error { _ggcef._bdgeb = form; return nil }

func (_cfbaee *PdfWriter) flushWriter() error {
	if _cfbaee._dacaeg == nil {
		_cfbaee._dacaeg = _cfbaee._agfba.Flush()
	}
	return _cfbaee._dacaeg
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_adfg pdfFontType0) GetRuneMetrics(r rune) (_gbe.CharMetrics, bool) {
	if _adfg.DescendantFont == nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _adfg)
		return _gbe.CharMetrics{}, false
	}
	return _adfg.DescendantFont.GetRuneMetrics(r)
}

func _edde(_egab *_abf.PdfObjectDictionary, _bgcda *fontCommon) (*pdfCIDFontType0, error) {
	if _bgcda._aacbc != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" {
		_acd.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0046\u006fn\u0074\u0020\u0053u\u0062\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020CI\u0044\u0046\u006fn\u0074\u0054y\u0070\u0065\u0030\u002e\u0020\u0066o\u006e\u0074=\u0025\u0073", _bgcda)
		return nil, _abf.ErrRangeError
	}
	_bccf := _bcce(_bgcda)
	_fbcfa, _ggag := _abf.GetDict(_egab.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
	if !_ggag {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043I\u0044\u0053\u0079st\u0065\u006d\u0049\u006e\u0066\u006f \u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074=\u0025\u0073", _bgcda)
		return nil, ErrRequiredAttributeMissing
	}
	_bccf.CIDSystemInfo = _fbcfa
	_bccf.DW = _egab.Get("\u0044\u0057")
	_bccf.W = _egab.Get("\u0057")
	_bccf.DW2 = _egab.Get("\u0044\u0057\u0032")
	_bccf.W2 = _egab.Get("\u0057\u0032")
	_bccf._bdced = 1000.0
	if _ffga, _egffg := _abf.GetNumberAsFloat(_bccf.DW); _egffg == nil {
		_bccf._bdced = _ffga
	}
	_ecabg, _bcead := _fecf(_bccf.W)
	if _bcead != nil {
		return nil, _bcead
	}
	if _ecabg == nil {
		_ecabg = map[_cbb.CharCode]float64{}
	}
	_bccf._fbcfb = _ecabg
	return _bccf, nil
}

// ToPdfObject returns the PDF representation of the DSS dictionary.
func (_bacf *DSS) ToPdfObject() _abf.PdfObject {
	_ebce := _bacf._gffg.PdfObject.(*_abf.PdfObjectDictionary)
	_ebce.Clear()
	_ccbdc := _abf.MakeDict()
	for _gfcb, _agff := range _bacf.VRI {
		_ccbdc.Set(*_abf.MakeName(_gfcb), _agff.ToPdfObject())
	}
	_ebce.SetIfNotNil("\u0043\u0065\u0072t\u0073", _fdaf(_bacf.Certs))
	_ebce.SetIfNotNil("\u004f\u0043\u0053P\u0073", _fdaf(_bacf.OCSPs))
	_ebce.SetIfNotNil("\u0043\u0052\u004c\u0073", _fdaf(_bacf.CRLs))
	_ebce.Set("\u0056\u0052\u0049", _ccbdc)
	return _bacf._gffg
}

// SetLocation sets the `Location` field of the signature.
func (_bbgdf *PdfSignature) SetLocation(location string) { _bbgdf.Location = _abf.MakeString(location) }

func (_adagg *PdfPattern) getDict() *_abf.PdfObjectDictionary {
	if _gefb, _bcdbe := _adagg._bcfca.(*_abf.PdfIndirectObject); _bcdbe {
		_dccea, _bbdeg := _gefb.PdfObject.(*_abf.PdfObjectDictionary)
		if !_bbdeg {
			return nil
		}
		return _dccea
	} else if _ggcgf, _ffdbc := _adagg._bcfca.(*_abf.PdfObjectStream); _ffdbc {
		return _ggcgf.PdfObjectDictionary
	} else {
		_acd.Log.Debug("\u0054r\u0079\u0069\u006e\u0067\u0020\u0074\u006f a\u0063\u0063\u0065\u0073\u0073\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0062j\u0065\u0063t \u0074\u0079\u0070e\u0020\u0028\u0025\u0054\u0029", _adagg._bcfca)
		return nil
	}
}

// NewPdfTransformParamsDocMDP create a PdfTransformParamsDocMDP with the specific permissions.
func NewPdfTransformParamsDocMDP(permission _df.DocMDPPermission) *PdfTransformParamsDocMDP {
	return &PdfTransformParamsDocMDP{Type: _abf.MakeName("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"), P: _abf.MakeInteger(int64(permission)), V: _abf.MakeName("\u0031\u002e\u0032")}
}

// ToPdfObject converts PdfAcroForm to a PdfObject, i.e. an indirect object containing the
// AcroForm dictionary.
func (_cecba *PdfAcroForm) ToPdfObject() _abf.PdfObject {
	_beccd := _cecba._bgfc
	_cgcae := _beccd.PdfObject.(*_abf.PdfObjectDictionary)
	if _cecba.Fields != nil {
		_edagg := _abf.PdfObjectArray{}
		for _, _cgbg := range *_cecba.Fields {
			_cdbd := _cgbg.GetContext()
			if _cdbd != nil {
				_edagg.Append(_cdbd.ToPdfObject())
			} else {
				_edagg.Append(_cgbg.ToPdfObject())
			}
		}
		_cgcae.Set("\u0046\u0069\u0065\u006c\u0064\u0073", &_edagg)
	}
	if _cecba.NeedAppearances != nil {
		_cgcae.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cecba.NeedAppearances)
	} else {
		if _beadf := _cgcae.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _beadf != nil {
			_cgcae.Remove("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073")
		}
	}
	if _cecba.SigFlags != nil {
		_cgcae.Set("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073", _cecba.SigFlags)
	}
	if _cecba.CO != nil {
		_cgcae.Set("\u0043\u004f", _cecba.CO)
	}
	if _cecba.DR != nil {
		_cgcae.Set("\u0044\u0052", _cecba.DR.ToPdfObject())
	}
	if _cecba.DA != nil {
		_cgcae.Set("\u0044\u0041", _cecba.DA)
	}
	if _cecba.Q != nil {
		_cgcae.Set("\u0051", _cecba.Q)
	}
	if _cecba.XFA != nil {
		_cgcae.Set("\u0058\u0046\u0041", _cecba.XFA)
	}
	if _cecba.ADBEEchoSign != nil {
		_cgcae.Set("\u0041\u0044\u0042\u0045\u005f\u0045\u0063\u0068\u006f\u0053\u0069\u0067\u006e", _cecba.ADBEEchoSign)
	}
	return _beccd
}

// ToPdfObject returns an indirect object containing the signature field dictionary.
func (_ccfgb *PdfFieldSignature) ToPdfObject() _abf.PdfObject {
	if _ccfgb.PdfAnnotationWidget != nil {
		_ccfgb.PdfAnnotationWidget.ToPdfObject()
	}
	_ccfgb.PdfField.ToPdfObject()
	_gdef := _ccfgb._dgdc
	_bbee := _gdef.PdfObject.(*_abf.PdfObjectDictionary)
	_bbee.SetIfNotNil("\u0046\u0054", _abf.MakeName("\u0053\u0069\u0067"))
	_bbee.SetIfNotNil("\u004c\u006f\u0063\u006b", _ccfgb.Lock)
	_bbee.SetIfNotNil("\u0053\u0056", _ccfgb.SV)
	if _ccfgb.V != nil {
		_bbee.SetIfNotNil("\u0056", _ccfgb.V.ToPdfObject())
	}
	return _gdef
}

// ToPdfObject returns the PDF representation of the tiling pattern.
func (_bedeg *PdfTilingPattern) ToPdfObject() _abf.PdfObject {
	_bedeg.PdfPattern.ToPdfObject()
	_ebcba := _bedeg.getDict()
	if _bedeg.PaintType != nil {
		_ebcba.Set("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e", _bedeg.PaintType)
	}
	if _bedeg.TilingType != nil {
		_ebcba.Set("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _bedeg.TilingType)
	}
	if _bedeg.BBox != nil {
		_ebcba.Set("\u0042\u0042\u006f\u0078", _bedeg.BBox.ToPdfObject())
	}
	if _bedeg.XStep != nil {
		_ebcba.Set("\u0058\u0053\u0074e\u0070", _bedeg.XStep)
	}
	if _bedeg.YStep != nil {
		_ebcba.Set("\u0059\u0053\u0074e\u0070", _bedeg.YStep)
	}
	if _bedeg.Resources != nil {
		_ebcba.Set("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _bedeg.Resources.ToPdfObject())
	}
	if _bedeg.Matrix != nil {
		_ebcba.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _bedeg.Matrix)
	}
	return _bedeg._bcfca
}

// PdfColorDeviceRGB represents a color in DeviceRGB colorspace with R, G, B components, where component is
// defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceRGB [3]float64

// ToPdfObject converts the pdfCIDFontType0 to a PDF representation.
func (_gfdg *pdfCIDFontType0) ToPdfObject() _abf.PdfObject { return _abf.MakeNull() }

// Y returns the value of the yellow component of the color.
func (_addbf *PdfColorDeviceCMYK) Y() float64 { return _addbf[2] }

// HasPatternByName checks whether a pattern object is defined by the specified keyName.
func (_ddaad *PdfPageResources) HasPatternByName(keyName _abf.PdfObjectName) bool {
	_, _egacf := _ddaad.GetPatternByName(keyName)
	return _egacf
}

var _dadge = map[string]struct{}{"\u0046\u0054": {}, "\u004b\u0069\u0064\u0073": {}, "\u0054": {}, "\u0054\u0055": {}, "\u0054\u004d": {}, "\u0046\u0066": {}, "\u0056": {}, "\u0044\u0056": {}, "\u0041\u0041": {}, "\u0044\u0041": {}, "\u0051": {}, "\u0044\u0053": {}, "\u0052\u0056": {}}

// ToPdfObject implements interface PdfModel.
func (_cdbfe *Permissions) ToPdfObject() _abf.PdfObject { return _cdbfe._deefb }

// Encrypt encrypts the output file with a specified user/owner password.
func (_geadc *PdfWriter) Encrypt(userPass, ownerPass []byte, options *EncryptOptions) error {
	_aabge := RC4_128bit
	if options != nil {
		_aabge = options.Algorithm
	}
	_ffbf := _bga.PermOwner
	if options != nil {
		_ffbf = options.Permissions
	}
	var _eegab _bf.Filter
	switch _aabge {
	case RC4_128bit:
		_eegab = _bf.NewFilterV2(16)
	case AES_128bit:
		_eegab = _bf.NewFilterAESV2()
	case AES_256bit:
		_eegab = _bf.NewFilterAESV3()
	default:
		return _e.Errorf("\u0075n\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020a\u006cg\u006fr\u0069\u0074\u0068\u006d\u003a\u0020\u0025v", options.Algorithm)
	}
	_fabab, _gdeb, _daecg := _abf.PdfCryptNewEncrypt(_eegab, userPass, ownerPass, _ffbf)
	if _daecg != nil {
		return _daecg
	}
	_geadc._ddbgd = _fabab
	if _gdeb.Major != 0 {
		_geadc.SetVersion(_gdeb.Major, _gdeb.Minor)
	}
	_geadc._cebae = _gdeb.Encrypt
	_geadc._aefff, _geadc._cfbce = _gdeb.ID0, _gdeb.ID1
	_geegc := _abf.MakeIndirectObject(_gdeb.Encrypt)
	_geadc._dcdbb = _geegc
	_geadc.addObject(_geegc)
	return nil
}

// Evaluate runs the function. Input is [x1 x2 x3].
func (_fbacg *PdfFunctionType4) Evaluate(xVec []float64) ([]float64, error) {
	if _fbacg._fggda == nil {
		_fbacg._fggda = _ae.NewPSExecutor(_fbacg.Program)
	}
	var _cebed []_ae.PSObject
	for _, _cdcb := range xVec {
		_cebed = append(_cebed, _ae.MakeReal(_cdcb))
	}
	_gfcbb, _bfcad := _fbacg._fggda.Execute(_cebed)
	if _bfcad != nil {
		return nil, _bfcad
	}
	_aeaae, _bfcad := _ae.PSObjectArrayToFloat64Array(_gfcbb)
	if _bfcad != nil {
		return nil, _bfcad
	}
	return _aeaae, nil
}

// GetContainingPdfObject returns the container of the outline item (indirect object).
func (_cagc *PdfOutlineItem) GetContainingPdfObject() _abf.PdfObject { return _cagc._ceegc }

type modelManager struct {
	_baecg map[PdfModel]_abf.PdfObject
	_addgc map[_abf.PdfObject]PdfModel
}

func (_eaab *PdfReader) buildOutlineTree(_abbea _abf.PdfObject, _ddcac *PdfOutlineTreeNode, _ffcbd *PdfOutlineTreeNode, _fdba map[_abf.PdfObject]struct{}) (*PdfOutlineTreeNode, *PdfOutlineTreeNode, error) {
	if _fdba == nil {
		_fdba = map[_abf.PdfObject]struct{}{}
	}
	_fdba[_abbea] = struct{}{}
	_ggcbe, _bfbac := _abbea.(*_abf.PdfIndirectObject)
	if !_bfbac {
		return nil, nil, _e.Errorf("\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0063\u006f\u006et\u0061\u0069\u006e\u0065r\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _abbea)
	}
	_afcfa, _daddg := _ggcbe.PdfObject.(*_abf.PdfObjectDictionary)
	if !_daddg {
		return nil, nil, _fd.New("\u006e\u006f\u0074 a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	_acd.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 \u0074\u0072\u0065\u0065\u003a\u0020d\u0069\u0063\u0074\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0076\u0029\u0020p\u003a\u0020\u0025\u0070", _afcfa, _ggcbe, _ggcbe)
	if _efbbg := _afcfa.Get("\u0054\u0069\u0074l\u0065"); _efbbg != nil {
		_gdeec, _gdce := _eaab.newPdfOutlineItemFromIndirectObject(_ggcbe)
		if _gdce != nil {
			return nil, nil, _gdce
		}
		_gdeec.Parent = _ddcac
		_gdeec.Prev = _ffcbd
		_ffde := _abf.ResolveReference(_afcfa.Get("\u0046\u0069\u0072s\u0074"))
		if _, _adbce := _fdba[_ffde]; _ffde != nil && _ffde != _ggcbe && !_adbce {
			if !_abf.IsNullObject(_ffde) {
				_eeaaf, _bbdffa, _gdbec := _eaab.buildOutlineTree(_ffde, &_gdeec.PdfOutlineTreeNode, nil, _fdba)
				if _gdbec != nil {
					_acd.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0069\u0074\u0065\u006d\u0020\u0074\u0072\u0065\u0065\u003a \u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020n\u006f\u0064\u0065\u0020\u0063\u0068\u0069\u006c\u0064\u0072\u0065n\u002e", _gdbec)
				} else {
					_gdeec.First = _eeaaf
					_gdeec.Last = _bbdffa
				}
			}
		}
		_caafd := _abf.ResolveReference(_afcfa.Get("\u004e\u0065\u0078\u0074"))
		if _, _cgeeb := _fdba[_caafd]; _caafd != nil && _caafd != _ggcbe && !_cgeeb {
			if !_abf.IsNullObject(_caafd) {
				_ddfge, _adcff, _adgcd := _eaab.buildOutlineTree(_caafd, _ddcac, &_gdeec.PdfOutlineTreeNode, _fdba)
				if _adgcd != nil {
					_acd.Log.Debug("D\u0045\u0042U\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0075\u0069\u006c\u0064\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065\u0020\u0066\u006f\u0072\u0020\u004ee\u0078\u0074\u0020\u006e\u006f\u0064\u0065\u003a\u0020\u0025\u0076\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006e\u006f\u0064e\u002e", _adgcd)
				} else {
					_gdeec.Next = _ddfge
					return &_gdeec.PdfOutlineTreeNode, _adcff, nil
				}
			}
		}
		return &_gdeec.PdfOutlineTreeNode, &_gdeec.PdfOutlineTreeNode, nil
	}
	_bbgc, _agfa := _gabdad(_ggcbe)
	if _agfa != nil {
		return nil, nil, _agfa
	}
	_bbgc.Parent = _ddcac
	if _acafd := _afcfa.Get("\u0046\u0069\u0072s\u0074"); _acafd != nil {
		_acafd = _abf.ResolveReference(_acafd)
		if _, _eccd := _fdba[_acafd]; _acafd != nil && _acafd != _ggcbe && !_eccd {
			_cafeee := _abf.TraceToDirectObject(_acafd)
			if _, _cccgf := _cafeee.(*_abf.PdfObjectNull); !_cccgf && _cafeee != nil {
				_aebdf, _cacdd, _ddbe := _eaab.buildOutlineTree(_acafd, &_bbgc.PdfOutlineTreeNode, nil, _fdba)
				if _ddbe != nil {
					_acd.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020b\u0075\u0069\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0074\u0072\u0065\u0065\u003a\u0020\u0025\u0076\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006e\u006f\u0064\u0065 \u0063\u0068i\u006c\u0064r\u0065n\u002e", _ddbe)
				} else {
					_bbgc.First = _aebdf
					_bbgc.Last = _cacdd
				}
			}
		}
	}
	return &_bbgc.PdfOutlineTreeNode, &_bbgc.PdfOutlineTreeNode, nil
}

// SetPdfCreationDate sets the CreationDate attribute of the output PDF.
func SetPdfCreationDate(creationDate _f.Time) {
	_gaabd.Lock()
	defer _gaabd.Unlock()
	_egdgg = creationDate
}

// ToPdfObject returns the PDF representation of the function.
func (_gfcfa *PdfFunctionType2) ToPdfObject() _abf.PdfObject {
	_aegc := _abf.MakeDict()
	_aegc.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _abf.MakeInteger(2))
	_daafb := &_abf.PdfObjectArray{}
	for _, _dbgae := range _gfcfa.Domain {
		_daafb.Append(_abf.MakeFloat(_dbgae))
	}
	_aegc.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _daafb)
	if _gfcfa.Range != nil {
		_ffbba := &_abf.PdfObjectArray{}
		for _, _gfbee := range _gfcfa.Range {
			_ffbba.Append(_abf.MakeFloat(_gfbee))
		}
		_aegc.Set("\u0052\u0061\u006eg\u0065", _ffbba)
	}
	if _gfcfa.C0 != nil {
		_dbaga := &_abf.PdfObjectArray{}
		for _, _ecacb := range _gfcfa.C0 {
			_dbaga.Append(_abf.MakeFloat(_ecacb))
		}
		_aegc.Set("\u0043\u0030", _dbaga)
	}
	if _gfcfa.C1 != nil {
		_fafba := &_abf.PdfObjectArray{}
		for _, _egbcc := range _gfcfa.C1 {
			_fafba.Append(_abf.MakeFloat(_egbcc))
		}
		_aegc.Set("\u0043\u0031", _fafba)
	}
	_aegc.Set("\u004e", _abf.MakeFloat(_gfcfa.N))
	if _gfcfa._gaaae != nil {
		_gfcfa._gaaae.PdfObject = _aegc
		return _gfcfa._gaaae
	}
	return _aegc
}

// GetContentStreamWithEncoder returns the pattern cell's content stream and its encoder
func (_gabbgf *PdfTilingPattern) GetContentStreamWithEncoder() ([]byte, _abf.StreamEncoder, error) {
	_eggc, _gbddc := _gabbgf._bcfca.(*_abf.PdfObjectStream)
	if !_gbddc {
		_acd.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _gabbgf._bcfca)
		return nil, nil, _abf.ErrTypeError
	}
	_gdgce, _egddf := _abf.DecodeStream(_eggc)
	if _egddf != nil {
		_acd.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020s\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0065\u0072\u0072\u003a\u0020\u0025\u0076", _egddf)
		return nil, nil, _egddf
	}
	_acggec, _egddf := _abf.NewEncoderFromStream(_eggc)
	if _egddf != nil {
		_acd.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020f\u0069\u006e\u0064\u0069\u006e\u0067 \u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020%\u0076", _egddf)
		return nil, nil, _egddf
	}
	return _gdgce, _acggec, nil
}

// Encoder returns the font's text encoder.
func (_cace pdfCIDFontType2) Encoder() _cbb.TextEncoder { return _cace._geaca }

// PdfAnnotationRichMedia represents Rich Media annotations.
type PdfAnnotationRichMedia struct {
	*PdfAnnotation
	RichMediaSettings _abf.PdfObject
	RichMediaContent  _abf.PdfObject
}

// Encoder returns the font's text encoder.
func (_gedcc *PdfFont) Encoder() _cbb.TextEncoder {
	_ddca := _gedcc.actualFont()
	if _ddca == nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0045n\u0063\u006f\u0064er\u0020\u006e\u006f\u0074\u0020\u0069m\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0066o\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u003d%\u0023\u0054", _gedcc._gedca)
		return nil
	}
	return _ddca.Encoder()
}

func _becc(_abdg _abf.PdfObject) (*PdfColorspaceICCBased, error) {
	_gcad := &PdfColorspaceICCBased{}
	if _deeb, _adaf := _abdg.(*_abf.PdfIndirectObject); _adaf {
		_gcad._afcc = _deeb
	}
	_abdg = _abf.TraceToDirectObject(_abdg)
	_dbabd, _cabcc := _abdg.(*_abf.PdfObjectArray)
	if !_cabcc {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _dbabd.Len() != 2 {
		return nil, _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020c\u006f\u006c\u006fr\u0073p\u0061\u0063\u0065")
	}
	_abdg = _abf.TraceToDirectObject(_dbabd.Get(0))
	_bcagbg, _cabcc := _abdg.(*_abf.PdfObjectName)
	if !_cabcc {
		return nil, _e.Errorf("\u0049\u0043\u0043B\u0061\u0073\u0065\u0064 \u006e\u0061\u006d\u0065\u0020\u006e\u006ft\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_bcagbg != "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064" {
		return nil, _e.Errorf("\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0049\u0043\u0043\u0042a\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_abdg = _dbabd.Get(1)
	_acgca, _cabcc := _abf.GetStream(_abdg)
	if !_cabcc {
		_acd.Log.Error("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006e\u006f\u0074\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067\u0020t\u006f\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020%\u0054", _abdg)
		return nil, _e.Errorf("\u0049\u0043\u0043Ba\u0073\u0065\u0064\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_efee := _acgca.PdfObjectDictionary
	_gfac, _cabcc := _efee.Get("\u004e").(*_abf.PdfObjectInteger)
	if !_cabcc {
		return nil, _e.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u006d\u0069\u0073\u0073\u0069n\u0067\u0020\u004e\u0020\u0066\u0072\u006fm\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074")
	}
	if *_gfac != 1 && *_gfac != 3 && *_gfac != 4 {
		return nil, _e.Errorf("\u0049\u0043\u0043\u0042\u0061s\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020\u0028\u006e\u006f\u0074\u0020\u0031\u002c\u0033\u002c\u0034\u0029")
	}
	_gcad.N = int(*_gfac)
	if _ebad := _efee.Get("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e"); _ebad != nil {
		_aceee, _gfce := NewPdfColorspaceFromPdfObject(_ebad)
		if _gfce != nil {
			return nil, _gfce
		}
		_gcad.Alternate = _aceee
	}
	if _ceag := _efee.Get("\u0052\u0061\u006eg\u0065"); _ceag != nil {
		_ceag = _abf.TraceToDirectObject(_ceag)
		_edac, _afbdee := _ceag.(*_abf.PdfObjectArray)
		if !_afbdee {
			return nil, _e.Errorf("I\u0043\u0043\u0042\u0061\u0073\u0065d\u0020\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0061\u006e \u0061r\u0072\u0061\u0079")
		}
		if _edac.Len() != 2*_gcad.N {
			return nil, _e.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0052\u0061\u006e\u0067e\u0020\u0077\u0072\u006f\u006e\u0067 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073")
		}
		_gdbg, _bbcb := _edac.GetAsFloat64Slice()
		if _bbcb != nil {
			return nil, _bbcb
		}
		_gcad.Range = _gdbg
	} else {
		_gcad.Range = make([]float64, 2*_gcad.N)
		for _dcbaa := 0; _dcbaa < _gcad.N; _dcbaa++ {
			_gcad.Range[2*_dcbaa] = 0.0
			_gcad.Range[2*_dcbaa+1] = 1.0
		}
	}
	if _cdea := _efee.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _cdea != nil {
		_aefab, _abfcf := _cdea.(*_abf.PdfObjectStream)
		if !_abfcf {
			return nil, _e.Errorf("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u004de\u0074\u0061\u0064\u0061\u0074\u0061\u0020n\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		_gcad.Metadata = _aefab
	}
	_ggff, _affd := _abf.DecodeStream(_acgca)
	if _affd != nil {
		return nil, _affd
	}
	_gcad.Data = _ggff
	_gcad._bfgc = _acgca
	return _gcad, nil
}

// PdfAnnotationCaret represents Caret annotations.
// (Section 12.5.6.11).
type PdfAnnotationCaret struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	RD _abf.PdfObject
	Sy _abf.PdfObject
}

// SetOCProperties sets the optional content properties.
func (_eface *PdfWriter) SetOCProperties(ocProperties _abf.PdfObject) error {
	_bgcb := _eface._ddffc
	if ocProperties != nil {
		_acd.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u004f\u0043\u0020\u0050\u0072o\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002e\u002e\u002e")
		_bgcb.Set("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", ocProperties)
		return _eface.addObjects(ocProperties)
	}
	return nil
}

// ColorToRGB verifies that the input color is an RGB color. Method exists in
// order to satisfy the PdfColorspace interface.
func (_ggge *PdfColorspaceDeviceRGB) ColorToRGB(color PdfColor) (PdfColor, error) {
	_dagb, _ggdf := color.(*PdfColorDeviceRGB)
	if !_ggdf {
		_acd.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006f\u0072 \u006e\u006f\u0074\u0020\u0064\u0065\u0076\u0069\u0063\u0065 \u0052\u0047\u0042")
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return _dagb, nil
}

func (_gggbd *PdfFont) baseFields() *fontCommon {
	if _gggbd._gedca == nil {
		_acd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0062\u0061\u0073\u0065\u0046\u0069\u0065l\u0064s\u002e \u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e")
		return nil
	}
	return _gggbd._gedca.baseFields()
}

var (
	_gaabd  _c.Mutex
	_efdg   = ""
	_egdgg  _f.Time
	_edead  = ""
	_geggga = ""
	_edfdc  _f.Time
	_babfc  = ""
	_cgaaa  = ""
	_eabe   = ""
)

func (_adf *PdfReader) newPdfActionTransFromDict(_aef *_abf.PdfObjectDictionary) (*PdfActionTrans, error) {
	return &PdfActionTrans{Trans: _aef.Get("\u0054\u0072\u0061n\u0073")}, nil
}

// DecodeArray returns the range of color component values in CalGray colorspace.
func (_faeac *PdfColorspaceCalGray) DecodeArray() []float64 { return []float64{0.0, 1.0} }

// IsTerminal returns true for terminal fields, false otherwise.
// Terminal fields are fields whose descendants are only widget annotations.
func (_caa *PdfField) IsTerminal() bool { return len(_caa.Kids) == 0 }

// NewOutlineItem returns a new outline item instance.
func NewOutlineItem(title string, dest OutlineDest) *OutlineItem {
	return &OutlineItem{Title: title, Dest: dest}
}

// NewPdfAnnotationFileAttachment returns a new file attachment annotation.
func NewPdfAnnotationFileAttachment() *PdfAnnotationFileAttachment {
	_ccfc := NewPdfAnnotation()
	_gegg := &PdfAnnotationFileAttachment{}
	_gegg.PdfAnnotation = _ccfc
	_gegg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ccfc.SetContext(_gegg)
	return _gegg
}

func (_geeac *PdfReader) newPdfSignatureReferenceFromDict(_deegg *_abf.PdfObjectDictionary) (*PdfSignatureReference, error) {
	if _efafb, _fabg := _geeac._ceecd.GetModelFromPrimitive(_deegg).(*PdfSignatureReference); _fabg {
		return _efafb, nil
	}
	_ddgfg := &PdfSignatureReference{_bfbaf: _deegg, Data: _deegg.Get("\u0044\u0061\u0074\u0061")}
	var _dcceab bool
	_ddgfg.Type, _ = _abf.GetName(_deegg.Get("\u0054\u0079\u0070\u0065"))
	_ddgfg.TransformMethod, _dcceab = _abf.GetName(_deegg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"))
	if !_dcceab {
		_acd.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0072\u0061\u006e\u0073\u0066o\u0072\u006dM\u0065\u0074h\u006f\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020in\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020m\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_ddgfg.TransformParams, _ = _abf.GetDict(_deegg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"))
	_ddgfg.DigestMethod, _ = _abf.GetName(_deegg.Get("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064"))
	return _ddgfg, nil
}

// GetStructRoot gets the StructTreeRoot object
func (_fdad *PdfPage) GetStructTreeRoot() (*_abf.PdfObject, bool) {
	_cbfdc, _ebbg := _fdad._dbaef.GetCatalogStructTreeRoot()
	return &_cbfdc, _ebbg
}

// PdfActionSetOCGState represents a SetOCGState action.
type PdfActionSetOCGState struct {
	*PdfAction
	State      _abf.PdfObject
	PreserveRB _abf.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_debe *PdfAnnotationWatermark) ToPdfObject() _abf.PdfObject {
	_debe.PdfAnnotation.ToPdfObject()
	_fadc := _debe._dbc
	_agbf := _fadc.PdfObject.(*_abf.PdfObjectDictionary)
	_agbf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0057a\u0074\u0065\u0072\u006d\u0061\u0072k"))
	_agbf.SetIfNotNil("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074", _debe.FixedPrint)
	return _fadc
}

func (_dbdb *PdfWriter) copyObjects() {
	_gcfb := make(map[_abf.PdfObject]_abf.PdfObject)
	_babgad := make([]_abf.PdfObject, 0, len(_dbdb._edcgc))
	_cbfda := make(map[_abf.PdfObject]struct{}, len(_dbdb._edcgc))
	_ggecc := make(map[_abf.PdfObject]struct{})
	for _, _aaeb := range _dbdb._edcgc {
		_fefgd := _dbdb.copyObject(_aaeb, _gcfb, _ggecc, false)
		if _, _gdacc := _ggecc[_aaeb]; _gdacc {
			continue
		}
		_babgad = append(_babgad, _fefgd)
		_cbfda[_fefgd] = struct{}{}
	}
	_dbdb._edcgc = _babgad
	_dbdb._fdgae = _cbfda
	_dbdb._ddegc = _dbdb.copyObject(_dbdb._ddegc, _gcfb, nil, false).(*_abf.PdfIndirectObject)
	_dbdb._cfdde = _dbdb.copyObject(_dbdb._cfdde, _gcfb, nil, false).(*_abf.PdfIndirectObject)
	if _dbdb._dcdbb != nil {
		_dbdb._dcdbb = _dbdb.copyObject(_dbdb._dcdbb, _gcfb, nil, false).(*_abf.PdfIndirectObject)
	}
	if _dbdb._aegbd {
		_abde := make(map[_abf.PdfObject]int64)
		for _baebb, _eabg := range _dbdb._deff {
			if _ceee, _fbgga := _gcfb[_baebb]; _fbgga {
				_abde[_ceee] = _eabg
			} else {
				_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020a\u0070\u0070\u0065n\u0064\u0020\u006d\u006fd\u0065\u0020\u002d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u0070\u0079\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070")
			}
		}
		_dbdb._deff = _abde
	}
}

func _fdbbe(_geab []byte) ([]byte, error) {
	_bgfddc := _eg.New()
	if _, _gfdgcf := _gc.Copy(_bgfddc, _dd.NewReader(_geab)); _gfdgcf != nil {
		return nil, _gfdgcf
	}
	return _bgfddc.Sum(nil), nil
}

func (_cdbdb *PdfWriter) checkPendingObjects() {
	for _cdcfa, _eabgg := range _cdbdb._fadb {
		if !_cdbdb.hasObject(_cdcfa) {
			_acd.Log.Debug("\u0057\u0041\u0052\u004e\u0020\u0050\u0065n\u0064\u0069\u006eg\u0020\u006f\u0062j\u0065\u0063t\u0020\u0025\u002b\u0076\u0020\u0025T\u0020(%\u0070\u0029\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0077\u0072\u0069\u0074\u0069\u006e\u0067", _cdcfa, _cdcfa, _cdcfa)
			for _, _cebff := range _eabgg {
				for _, _dcffa := range _cebff.Keys() {
					_ebcfg := _cebff.Get(_dcffa)
					if _ebcfg == _cdcfa {
						_acd.Log.Debug("\u0050e\u006e\u0064i\u006e\u0067\u0020\u006fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0061nd\u0020\u0072\u0065p\u006c\u0061c\u0065\u0064\u0020\u0077\u0069\u0074h\u0020\u006eu\u006c\u006c")
						_cebff.Set(_dcffa, _abf.MakeNull())
						break
					}
				}
			}
		}
	}
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
	DiffResults *_df.DiffResults
	IsCrlFound  bool
	IsOcspFound bool

	// GeneralizedTime is the time at which the time-stamp token has been created by the TSA (RFC 3161).
	GeneralizedTime _f.Time
}

// PdfAnnotationUnderline represents Underline annotations.
// (Section 12.5.6.10).
type PdfAnnotationUnderline struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _abf.PdfObject
}

func (_gcb *PdfReader) newPdfAnnotation3DFromDict(_eede *_abf.PdfObjectDictionary) (*PdfAnnotation3D, error) {
	_acec := PdfAnnotation3D{}
	_acec.T3DD = _eede.Get("\u0033\u0044\u0044")
	_acec.T3DV = _eede.Get("\u0033\u0044\u0056")
	_acec.T3DA = _eede.Get("\u0033\u0044\u0041")
	_acec.T3DI = _eede.Get("\u0033\u0044\u0049")
	_acec.T3DB = _eede.Get("\u0033\u0044\u0042")
	return &_acec, nil
}

func (_febc *PdfReader) newPdfAnnotationInkFromDict(_bgd *_abf.PdfObjectDictionary) (*PdfAnnotationInk, error) {
	_dfe := PdfAnnotationInk{}
	_defc, _eege := _febc.newPdfAnnotationMarkupFromDict(_bgd)
	if _eege != nil {
		return nil, _eege
	}
	_dfe.PdfAnnotationMarkup = _defc
	_dfe.InkList = _bgd.Get("\u0049n\u006b\u004c\u0069\u0073\u0074")
	_dfe.BS = _bgd.Get("\u0042\u0053")
	return &_dfe, nil
}

func _becce(_bbdab string) (map[_cbb.CharCode]_cbb.GlyphName, error) {
	_cfdb := _be.Split(_bbdab, "\u000a")
	_ffbaf := make(map[_cbb.CharCode]_cbb.GlyphName)
	for _, _adcg := range _cfdb {
		_degeb := _gffgf.FindStringSubmatch(_adcg)
		if _degeb == nil {
			continue
		}
		_agdde, _gddafd := _degeb[1], _degeb[2]
		_gccd, _acbbc := _gb.Atoi(_agdde)
		if _acbbc != nil {
			_acd.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0042\u0061\u0064\u0020\u0065\u006e\u0063\u006fd\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u002e \u0025\u0071", _adcg)
			return nil, _abf.ErrTypeError
		}
		_ffbaf[_cbb.CharCode(_gccd)] = _cbb.GlyphName(_gddafd)
	}
	_acd.Log.Trace("g\u0065\u0074\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0073\u003a\u0020\u006b\u0065\u0079V\u0061\u006c\u0075e\u0073=\u0025\u0023\u0076", _ffbaf)
	return _ffbaf, nil
}

// NewOutline returns a new outline instance.
func NewOutline() *Outline { return &Outline{} }

// ToPdfObject returns a PDF object representation of the outline.
func (_abba *Outline) ToPdfObject() _abf.PdfObject { return _abba.ToPdfOutline().ToPdfObject() }

// Permissions specify a permissions dictionary (PDF 1.5).
// (Section 12.8.4, Table 258 - Entries in a permissions dictionary p. 477 in PDF32000_2008).
type Permissions struct {
	DocMDP *PdfSignature
	_deefb *_abf.PdfObjectDictionary
}

func (_cddg *PdfReader) newPdfAnnotationTextFromDict(_gfbd *_abf.PdfObjectDictionary) (*PdfAnnotationText, error) {
	_adad := PdfAnnotationText{}
	_gfcd, _bbb := _cddg.newPdfAnnotationMarkupFromDict(_gfbd)
	if _bbb != nil {
		return nil, _bbb
	}
	_adad.PdfAnnotationMarkup = _gfcd
	_adad.Open = _gfbd.Get("\u004f\u0070\u0065\u006e")
	_adad.Name = _gfbd.Get("\u004e\u0061\u006d\u0065")
	_adad.State = _gfbd.Get("\u0053\u0074\u0061t\u0065")
	_adad.StateModel = _gfbd.Get("\u0053\u0074\u0061\u0074\u0065\u004d\u006f\u0064\u0065\u006c")
	return &_adad, nil
}
func (_ddcfg *pdfCIDFontType2) baseFields() *fontCommon { return &_ddcfg.fontCommon }
func (_ffd *PdfReader) newPdfActionSoundFromDict(_dcg *_abf.PdfObjectDictionary) (*PdfActionSound, error) {
	return &PdfActionSound{Sound: _dcg.Get("\u0053\u006f\u0075n\u0064"), Volume: _dcg.Get("\u0056\u006f\u006c\u0075\u006d\u0065"), Synchronous: _dcg.Get("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073"), Repeat: _dcg.Get("\u0052\u0065\u0070\u0065\u0061\u0074"), Mix: _dcg.Get("\u004d\u0069\u0078")}, nil
}

// PdfField contains the common attributes of a form field. The context object contains the specific field data
// which can represent a button, text, choice or signature.
// The PdfField is typically not used directly, but is encapsulated by the more specific field types such as
// PdfFieldButton etc (i.e. the context attribute).
type PdfField struct {
	_ffea        PdfModel
	_dgdc        *_abf.PdfIndirectObject
	Parent       *PdfField
	Annotations  []*PdfAnnotationWidget
	Kids         []*PdfField
	FT           *_abf.PdfObjectName
	T            *_abf.PdfObjectString
	TU           *_abf.PdfObjectString
	TM           *_abf.PdfObjectString
	Ff           *_abf.PdfObjectInteger
	V            _abf.PdfObject
	DV           _abf.PdfObject
	AA           _abf.PdfObject
	VariableText *VariableText
}

// GetContainingPdfObject implements interface PdfModel.
func (_gaecc *Permissions) GetContainingPdfObject() _abf.PdfObject { return _gaecc._deefb }

func _cebb(_gfbad _abf.PdfObject, _dccge *fontCommon) (*_bd.CMap, error) {
	_ccee, _aefg := _abf.GetStream(_gfbad)
	if !_aefg {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0074\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0054\u006f\u0043m\u0061\u0070\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0054\u0029", _gfbad)
		return nil, _abf.ErrTypeError
	}
	_cccfa, _efdeb := _abf.DecodeStream(_ccee)
	if _efdeb != nil {
		return nil, _efdeb
	}
	_ebca, _efdeb := _bd.LoadCmapFromData(_cccfa, !_dccge.isCIDFont())
	if _efdeb != nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u004e\u0075\u006d\u0062\u0065\u0072\u003d\u0025\u0064\u0020\u0065\u0072r=\u0025\u0076", _ccee.ObjectNumber, _efdeb)
	}
	return _ebca, _efdeb
}

// ToPdfObject returns the PDF representation of the function.
func (_cffeb *PdfFunctionType3) ToPdfObject() _abf.PdfObject {
	_bdgc := _abf.MakeDict()
	_bdgc.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _abf.MakeInteger(3))
	_dbdd := &_abf.PdfObjectArray{}
	for _, _gcab := range _cffeb.Domain {
		_dbdd.Append(_abf.MakeFloat(_gcab))
	}
	_bdgc.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _dbdd)
	if _cffeb.Range != nil {
		_edcbg := &_abf.PdfObjectArray{}
		for _, _cdcf := range _cffeb.Range {
			_edcbg.Append(_abf.MakeFloat(_cdcf))
		}
		_bdgc.Set("\u0052\u0061\u006eg\u0065", _edcbg)
	}
	if _cffeb.Functions != nil {
		_gddec := &_abf.PdfObjectArray{}
		for _, _aggca := range _cffeb.Functions {
			_gddec.Append(_aggca.ToPdfObject())
		}
		_bdgc.Set("\u0046u\u006e\u0063\u0074\u0069\u006f\u006es", _gddec)
	}
	if _cffeb.Bounds != nil {
		_eadcf := &_abf.PdfObjectArray{}
		for _, _ebaag := range _cffeb.Bounds {
			_eadcf.Append(_abf.MakeFloat(_ebaag))
		}
		_bdgc.Set("\u0042\u006f\u0075\u006e\u0064\u0073", _eadcf)
	}
	if _cffeb.Encode != nil {
		_bfdcb := &_abf.PdfObjectArray{}
		for _, _acbff := range _cffeb.Encode {
			_bfdcb.Append(_abf.MakeFloat(_acbff))
		}
		_bdgc.Set("\u0045\u006e\u0063\u006f\u0064\u0065", _bfdcb)
	}
	if _cffeb._edacd != nil {
		_cffeb._edacd.PdfObject = _bdgc
		return _cffeb._edacd
	}
	return _bdgc
}

func _daacfg(_baaeg *_abf.PdfObjectDictionary) (*PdfShadingType5, error) {
	_cfbff := PdfShadingType5{}
	_abbc := _baaeg.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _abbc == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_egadc, _ggbcd := _abbc.(*_abf.PdfObjectInteger)
	if !_ggbcd {
		_acd.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _abbc)
		return nil, _abf.ErrTypeError
	}
	_cfbff.BitsPerCoordinate = _egadc
	_abbc = _baaeg.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _abbc == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_egadc, _ggbcd = _abbc.(*_abf.PdfObjectInteger)
	if !_ggbcd {
		_acd.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _abbc)
		return nil, _abf.ErrTypeError
	}
	_cfbff.BitsPerComponent = _egadc
	_abbc = _baaeg.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077")
	if _abbc == nil {
		_acd.Log.Debug("\u0052\u0065\u0071u\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0056\u0065\u0072\u0074\u0069c\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077")
		return nil, ErrRequiredAttributeMissing
	}
	_egadc, _ggbcd = _abbc.(*_abf.PdfObjectInteger)
	if !_ggbcd {
		_acd.Log.Debug("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050\u0065\u0072\u0052\u006f\u0077\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006et\u0065\u0067\u0065\u0072\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _abbc)
		return nil, _abf.ErrTypeError
	}
	_cfbff.VerticesPerRow = _egadc
	_abbc = _baaeg.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _abbc == nil {
		_acd.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_bceee, _ggbcd := _abbc.(*_abf.PdfObjectArray)
	if !_ggbcd {
		_acd.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _abbc)
		return nil, _abf.ErrTypeError
	}
	_cfbff.Decode = _bceee
	if _badbd := _baaeg.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _badbd != nil {
		_cfbff.Function = []PdfFunction{}
		if _degfd, _gdedc := _badbd.(*_abf.PdfObjectArray); _gdedc {
			for _, _ebefg := range _degfd.Elements() {
				_dcfdg, _cbcgf := _ebedg(_ebefg)
				if _cbcgf != nil {
					_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cbcgf)
					return nil, _cbcgf
				}
				_cfbff.Function = append(_cfbff.Function, _dcfdg)
			}
		} else {
			_dbceb, _bgdga := _ebedg(_badbd)
			if _bgdga != nil {
				_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bgdga)
				return nil, _bgdga
			}
			_cfbff.Function = append(_cfbff.Function, _dbceb)
		}
	}
	return &_cfbff, nil
}

// EncryptOptions represents encryption options for an output PDF.
type EncryptOptions struct {
	Permissions _bga.Permissions
	Algorithm   EncryptionAlgorithm
}

// GenerateXObjectName generates an unused XObject name that can be used for
// adding new XObjects. Uses format XObj1, XObj2, ...
func (_bcaba *PdfPageResources) GenerateXObjectName() _abf.PdfObjectName {
	_acdae := 1
	for {
		_cbddf := _abf.MakeName(_e.Sprintf("\u0058\u004f\u0062\u006a\u0025\u0064", _acdae))
		if !_bcaba.HasXObjectByName(*_cbddf) {
			return *_cbddf
		}
		_acdae++
	}
}

var _ pdfFont = (*pdfCIDFontType0)(nil)

// SubsetRegistered subsets the font to only the glyphs that have been registered by the encoder.
//
// NOTE: This only works on fonts that support subsetting. For unsupported fonts this is a no-op, although a debug
// message is emitted.  Currently supported fonts are embedded Truetype CID fonts (type 0).
//
// NOTE: Make sure to call this soon before writing (once all needed runes have been registered).
// If using package creator, use its EnableFontSubsetting method instead.
func (_cafa *PdfFont) SubsetRegistered() error {
	switch _fgbe := _cafa._gedca.(type) {
	case *pdfFontType0:
		_baacb := _fgbe.subsetRegistered()
		if _baacb != nil {
			_acd.Log.Debug("\u0053\u0075b\u0073\u0065\u0074 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _baacb)
			return _baacb
		}
		if _fgbe._bgefb != nil {
			if _fgbe._edeaf != nil {
				_fgbe._edeaf.ToPdfObject()
			}
			_fgbe.ToPdfObject()
		}
	default:
		_acd.Log.Debug("F\u006f\u006e\u0074\u0020\u0025\u0054 \u0064\u006f\u0065\u0073\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006fr\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069n\u0067", _fgbe)
	}
	return nil
}

func (_ggda *PdfAppender) replaceObject(_ddbg, _gafge _abf.PdfObject) {
	switch _dcba := _ddbg.(type) {
	case *_abf.PdfIndirectObject:
		_ggda._bge[_gafge] = _dcba.ObjectNumber
	case *_abf.PdfObjectStream:
		_ggda._bge[_gafge] = _dcba.ObjectNumber
	}
}

// GetNumPages returns the number of pages in the document.
func (_gfaaa *PdfReader) GetNumPages() (int, error) {
	if _gfaaa._bebc.GetCrypter() != nil && !_gfaaa._bebc.IsAuthenticated() {
		return 0, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	return len(_gfaaa._gbfaf), nil
}

// NewPdfAnnotationMovie returns a new movie annotation.
func NewPdfAnnotationMovie() *PdfAnnotationMovie {
	_dad := NewPdfAnnotation()
	_cbd := &PdfAnnotationMovie{}
	_cbd.PdfAnnotation = _dad
	_dad.SetContext(_cbd)
	return _cbd
}

// GetContainingPdfObject returns the container of the outline (indirect object).
func (_gcega *PdfOutline) GetContainingPdfObject() _abf.PdfObject { return _gcega._cgcg }

// Add appends a top level outline item to the outline.
func (_gfceea *Outline) Add(item *OutlineItem) { _gfceea.Entries = append(_gfceea.Entries, item) }

func _acffa(_efbdb _abf.PdfObject) (*PdfColorspaceSpecialIndexed, error) {
	_aaac := NewPdfColorspaceSpecialIndexed()
	if _gacbb, _ccfg := _efbdb.(*_abf.PdfIndirectObject); _ccfg {
		_aaac._acea = _gacbb
	}
	_efbdb = _abf.TraceToDirectObject(_efbdb)
	_bdge, _ddbbb := _efbdb.(*_abf.PdfObjectArray)
	if !_ddbbb {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _bdge.Len() != 4 {
		return nil, _e.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
	}
	_efbdb = _bdge.Get(0)
	_ggfaa, _ddbbb := _efbdb.(*_abf.PdfObjectName)
	if !_ddbbb {
		return nil, _e.Errorf("\u0069n\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
	}
	if *_ggfaa != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _e.Errorf("\u0069\u006e\u0064\u0065xe\u0064\u0020\u0043\u0053\u003a\u0020\u0077\u0072\u006f\u006e\u0067\u0020\u006e\u0061m\u0065")
	}
	_efbdb = _bdge.Get(1)
	_cbcbc, _aefeb := DetermineColorspaceNameFromPdfObject(_efbdb)
	if _aefeb != nil {
		return nil, _aefeb
	}
	if _cbcbc == "\u0049n\u0064\u0065\u0078\u0065\u0064" || _cbcbc == "\u0050a\u0074\u0074\u0065\u0072\u006e" {
		_acd.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0049\u006e\u0064e\u0078\u0065\u0064\u002f\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043S\u0020\u0061\u0073\u0020\u0062\u0061\u0073\u0065\u0020\u0028\u0025v\u0029", _cbcbc)
		return nil, _bgaaa
	}
	_fbedd, _aefeb := NewPdfColorspaceFromPdfObject(_efbdb)
	if _aefeb != nil {
		return nil, _aefeb
	}
	_aaac.Base = _fbedd
	_efbdb = _bdge.Get(2)
	_dcaac, _aefeb := _abf.GetNumberAsInt64(_efbdb)
	if _aefeb != nil {
		return nil, _aefeb
	}
	if _dcaac > 255 {
		return nil, _e.Errorf("\u0069n\u0064\u0065\u0078\u0065d\u0020\u0043\u0053\u003a\u0020I\u006ev\u0061l\u0069\u0064\u0020\u0068\u0069\u0076\u0061l")
	}
	_aaac.HiVal = int(_dcaac)
	_efbdb = _bdge.Get(3)
	_aaac.Lookup = _efbdb
	_efbdb = _abf.TraceToDirectObject(_efbdb)
	var _afag []byte
	if _cbcc, _fcgb := _efbdb.(*_abf.PdfObjectString); _fcgb {
		_afag = _cbcc.Bytes()
		_acd.Log.Trace("\u0049\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0020\u0064", _afag)
	} else if _gcbbd, _adgd := _efbdb.(*_abf.PdfObjectStream); _adgd {
		_acd.Log.Trace("\u0049n\u0064e\u0078\u0065\u0064\u0020\u0073t\u0072\u0065a\u006d\u003a\u0020\u0025\u0073", _efbdb.String())
		_acd.Log.Trace("\u0045\u006e\u0063\u006fde\u0064\u0020\u0028\u0025\u0064\u0029\u0020\u003a\u0020\u0025\u0023\u0020\u0078", len(_gcbbd.Stream), _gcbbd.Stream)
		_dabc, _ccge := _abf.DecodeStream(_gcbbd)
		if _ccge != nil {
			return nil, _ccge
		}
		_acd.Log.Trace("\u0044e\u0063o\u0064\u0065\u0064\u0020\u0028%\u0064\u0029 \u003a\u0020\u0025\u0020\u0058", len(_dabc), _dabc)
		_afag = _dabc
	} else {
		_acd.Log.Debug("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _efbdb)
		return nil, _e.Errorf("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0043\u0053\u003a\u0020\u0049\u006e\u0076a\u006ci\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if len(_afag) < _aaac.Base.GetNumComponents()*(_aaac.HiVal+1) {
		_acd.Log.Debug("\u0050\u0044\u0046\u0020\u0049\u006e\u0063o\u006d\u0070\u0061t\u0069\u0062\u0069\u006ci\u0074\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074\u006f\u006f\u0020\u0073\u0068\u006f\u0072\u0074")
		_acd.Log.Debug("\u0046\u0061i\u006c\u002c\u0020\u006c\u0065\u006e\u0028\u0064\u0061\u0074\u0061\u0029\u003a\u0020\u0025\u0064\u002c\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u002c\u0020\u0068\u0069\u0056\u0061\u006c\u003a\u0020\u0025\u0064", len(_afag), _aaac.Base.GetNumComponents(), _aaac.HiVal)
	} else {
		_afag = _afag[:_aaac.Base.GetNumComponents()*(_aaac.HiVal+1)]
	}
	_aaac._bcdf = _afag
	return _aaac, nil
}

// PdfColorspaceDeviceGray represents a grayscale colorspace.
type PdfColorspaceDeviceGray struct{}

// NewPdfAnnotationStamp returns a new stamp annotation.
func NewPdfAnnotationStamp() *PdfAnnotationStamp {
	_ageg := NewPdfAnnotation()
	_gbeb := &PdfAnnotationStamp{}
	_gbeb.PdfAnnotation = _ageg
	_gbeb.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ageg.SetContext(_gbeb)
	return _gbeb
}

// SetPdfSubject sets the Subject attribute of the output PDF.
func SetPdfSubject(subject string) { _gaabd.Lock(); defer _gaabd.Unlock(); _cgaaa = subject }

// EnableAll LTV enables all signatures in the PDF document.
// The signing certificate chain is extracted from each signature dictionary.
// Optionally, additional certificates can be specified through the
// `extraCerts` parameter. The LTV client attempts to build the certificate
// chain up to a trusted root by downloading any missing certificates.
func (_aggg *LTV) EnableAll(extraCerts []*_fa.Certificate) error {
	_fgaaee := _aggg._bfed._agda.AcroForm
	for _, _ddece := range _fgaaee.AllFields() {
		_fbgb, _ := _ddece.GetContext().(*PdfFieldSignature)
		if _fbgb == nil {
			continue
		}
		_cafgg := _fbgb.V
		if _ceebf := _aggg.validateSig(_cafgg); _ceebf != nil {
			_acd.Log.Debug("\u0057\u0041\u0052N\u003a\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076", _ceebf)
		}
		if _abdfbf := _aggg.Enable(_cafgg, extraCerts); _abdfbf != nil {
			return _abdfbf
		}
	}
	return nil
}

// DecodeArray returns the range of color component values in the ICCBased colorspace.
func (_dcae *PdfColorspaceICCBased) DecodeArray() []float64 { return _dcae.Range }

// NewPdfActionGoToE returns a new "go to embedded" action.
func NewPdfActionGoToE() *PdfActionGoToE {
	_cg := NewPdfAction()
	_dec := &PdfActionGoToE{}
	_dec.PdfAction = _cg
	_cg.SetContext(_dec)
	return _dec
}

// PdfAnnotationCircle represents Circle annotations.
// (Section 12.5.6.8).
type PdfAnnotationCircle struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _abf.PdfObject
	IC _abf.PdfObject
	BE _abf.PdfObject
	RD _abf.PdfObject
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

// PageFromIndirectObject returns the PdfPage and page number for a given indirect object.
func (_fafce *PdfReader) PageFromIndirectObject(ind *_abf.PdfIndirectObject) (*PdfPage, int, error) {
	if len(_fafce.PageList) != len(_fafce._gbfaf) {
		return nil, 0, _fd.New("\u0070\u0061\u0067\u0065\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	for _adffc, _cgfd := range _fafce._gbfaf {
		if _cgfd == ind {
			return _fafce.PageList[_adffc], _adffc + 1, nil
		}
	}
	return nil, 0, _fd.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// NewPdfColorPatternType3 returns an empty color shading pattern type 3 (Radial).
func NewPdfColorPatternType3() *PdfColorPatternType3 { _ggdd := &PdfColorPatternType3{}; return _ggdd }

// PdfSignatureReference represents a PDF signature reference dictionary and is used for signing via form signature fields.
// (Section 12.8.1, Table 253 - Entries in a signature reference dictionary p. 469 in PDF32000_2008).
type PdfSignatureReference struct {
	_bfbaf          *_abf.PdfObjectDictionary
	Type            *_abf.PdfObjectName
	TransformMethod *_abf.PdfObjectName
	TransformParams _abf.PdfObject
	Data            _abf.PdfObject
	DigestMethod    *_abf.PdfObjectName
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
	_gaaae *_abf.PdfIndirectObject
}

// ToPdfObject returns the PDF representation of the pattern.
func (_aaab *PdfPattern) ToPdfObject() _abf.PdfObject {
	_baeff := _aaab.getDict()
	_baeff.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
	_baeff.Set("P\u0061\u0074\u0074\u0065\u0072\u006e\u0054\u0079\u0070\u0065", _abf.MakeInteger(_aaab.PatternType))
	return _aaab._bcfca
}

// PdfOutline represents a PDF outline dictionary (Table 152 - p. 376).
type PdfOutline struct {
	PdfOutlineTreeNode
	Parent *PdfOutlineTreeNode
	Count  *int64
	_cgcg  *_abf.PdfIndirectObject
}

func _fffc(_aadga, _fcdac string) string {
	if _be.Contains(_aadga, "\u002b") {
		_gfec := _be.Split(_aadga, "\u002b")
		if len(_gfec) == 2 {
			_aadga = _gfec[1]
		}
	}
	return _fcdac + "\u002b" + _aadga
}
func _ebgb() string { _gaabd.Lock(); defer _gaabd.Unlock(); return _cgaaa }

// A returns the value of the A component of the color.
func (_edfc *PdfColorCalRGB) A() float64 { return _edfc[0] }

// PdfFunctionType4 is a Postscript calculator functions.
type PdfFunctionType4 struct {
	Domain  []float64
	Range   []float64
	Program *_ae.PSProgram
	_fggda  *_ae.PSExecutor
	_cddgf  []byte
	_dgbdc  *_abf.PdfObjectStream
}

// RunesToCharcodeBytes maps the provided runes to charcode bytes and it
// returns the resulting slice of bytes, along with the number of runes which
// could not be converted. If the number of misses is 0, all runes were
// successfully converted.
func (_ddcc *PdfFont) RunesToCharcodeBytes(data []rune) ([]byte, int) {
	var _aegd []_cbb.TextEncoder
	var _ceadc _cbb.CMapEncoder
	if _gaccb := _ddcc.baseFields()._aabfe; _gaccb != nil {
		_ceadc = _cbb.NewCMapEncoder("", nil, _gaccb)
	}
	_bdfd := _ddcc.Encoder()
	if _bdfd != nil {
		switch _cadf := _bdfd.(type) {
		case _cbb.SimpleEncoder:
			_dbcbb := _cadf.BaseName()
			if _, _fbcb := _bdgdc[_dbcbb]; _fbcb {
				_aegd = append(_aegd, _bdfd)
			}
		}
	}
	if len(_aegd) == 0 {
		if _ddcc.baseFields()._aabfe != nil {
			_aegd = append(_aegd, _ceadc)
		}
		if _bdfd != nil {
			_aegd = append(_aegd, _bdfd)
		}
	}
	var _fcaad _dd.Buffer
	var _eacc int
	for _, _dfdg := range data {
		var _gacba bool
		for _, _dfdac := range _aegd {
			if _ecff := _dfdac.Encode(string(_dfdg)); len(_ecff) > 0 {
				_fcaad.Write(_ecff)
				_gacba = true
				break
			}
		}
		if !_gacba {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020f\u0061\u0069\u006ce\u0064\u0020\u0074\u006f \u006d\u0061\u0070\u0020\u0072\u0075\u006e\u0065\u0020\u0060\u0025\u002b\u0071\u0060\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065", _dfdg)
			_eacc++
		}
	}
	if _eacc != 0 {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0061\u006cl\u0020\u0072\u0075\u006e\u0065\u0073\u0020\u0074\u006f\u0020\u0063\u0068\u0061\u0072c\u006fd\u0065\u0073\u002e\u000a"+"\u0009\u006e\u0075\u006d\u0052\u0075\u006e\u0065\u0073\u003d\u0025d\u0020\u006e\u0075\u006d\u004d\u0069\u0073\u0073\u0065\u0073=\u0025\u0064\u000a"+"\t\u0066\u006f\u006e\u0074=%\u0073 \u0065\u006e\u0063\u006f\u0064e\u0072\u0073\u003d\u0025\u002b\u0076", len(data), _eacc, _ddcc, _aegd)
	}
	return _fcaad.Bytes(), _eacc
}

// NewPdfShadingPatternType2 creates an empty shading pattern type 2 object.
func NewPdfShadingPatternType2() *PdfShadingPatternType2 {
	_aedgd := &PdfShadingPatternType2{}
	_aedgd.Matrix = _abf.MakeArrayFromIntegers([]int{1, 0, 0, 1, 0, 0})
	_aedgd.PdfPattern = &PdfPattern{}
	_aedgd.PdfPattern.PatternType = int64(*_abf.MakeInteger(2))
	_aedgd.PdfPattern._bgafe = _aedgd
	_aedgd.PdfPattern._bcfca = _abf.MakeIndirectObject(_abf.MakeDict())
	return _aedgd
}

// WriteString outputs the object as it is to be written to file.
func (_fddgdb *pdfSignDictionary) WriteString() string {
	_fddgdb._dgfdf = 0
	_fddgdb._afgef = 0
	_fddgdb._edcbf = 0
	_fddgdb._bcbcg = 0
	_ebebb := _dd.NewBuffer(nil)
	_ebebb.WriteString("\u003c\u003c")
	for _, _egcfb := range _fddgdb.Keys() {
		_dbecc := _fddgdb.Get(_egcfb)
		switch _egcfb {
		case "\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e":
			_ebebb.WriteString(_egcfb.WriteString())
			_ebebb.WriteString("\u0020")
			_fddgdb._edcbf = _ebebb.Len()
			_ebebb.WriteString(_dbecc.WriteString())
			_ebebb.WriteString("\u0020")
			_fddgdb._bcbcg = _ebebb.Len() - 1
		case "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073":
			_ebebb.WriteString(_egcfb.WriteString())
			_ebebb.WriteString("\u0020")
			_fddgdb._dgfdf = _ebebb.Len()
			_ebebb.WriteString(_dbecc.WriteString())
			_ebebb.WriteString("\u0020")
			_fddgdb._afgef = _ebebb.Len() - 1
		default:
			_ebebb.WriteString(_egcfb.WriteString())
			_ebebb.WriteString("\u0020")
			_ebebb.WriteString(_dbecc.WriteString())
		}
	}
	_ebebb.WriteString("\u003e\u003e")
	return _ebebb.String()
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element.
func (_dcbd *PdfColorspaceSpecialSeparation) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_dadea, _abcgf := _abf.GetNumbersAsFloat(objects)
	if _abcgf != nil {
		return nil, _abcgf
	}
	return _dcbd.ColorFromFloats(_dadea)
}

// NewXObjectImageFromStream builds the image xobject from a stream object.
// An image dictionary is the dictionary portion of a stream object representing an image XObject.
func NewXObjectImageFromStream(stream *_abf.PdfObjectStream) (*XObjectImage, error) {
	_ecdcf := &XObjectImage{}
	_ecdcf._ccbad = stream
	_aadef := *(stream.PdfObjectDictionary)
	_gggac, _fcfd := _abf.NewEncoderFromStream(stream)
	if _fcfd != nil {
		return nil, _fcfd
	}
	_ecdcf.Filter = _gggac
	if _abdca := _abf.TraceToDirectObject(_aadef.Get("\u0057\u0069\u0064t\u0068")); _abdca != nil {
		_ddcaf, _beeeb := _abdca.(*_abf.PdfObjectInteger)
		if !_beeeb {
			return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_eadgd := int64(*_ddcaf)
		_ecdcf.Width = &_eadgd
	} else {
		return nil, _fd.New("\u0077\u0069\u0064\u0074\u0068\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _bfag := _abf.TraceToDirectObject(_aadef.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _bfag != nil {
		_ecgdg, _bgeaab := _bfag.(*_abf.PdfObjectInteger)
		if !_bgeaab {
			return nil, _fd.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_abgab := int64(*_ecgdg)
		_ecdcf.Height = &_abgab
	} else {
		return nil, _fd.New("\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	if _fagee := _abf.TraceToDirectObject(_aadef.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065")); _fagee != nil {
		_gecaa, _acga := NewPdfColorspaceFromPdfObject(_fagee)
		if _acga != nil {
			return nil, _acga
		}
		_ecdcf.ColorSpace = _gecaa
	} else {
		_acd.Log.Debug("\u0058O\u0062\u006a\u0065c\u0074\u0020\u0049m\u0061ge\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067 1\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065n\u0074\u0020\u002d\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047r\u0061\u0079")
		_ecdcf.ColorSpace = NewPdfColorspaceDeviceGray()
	}
	if _gbace := _abf.TraceToDirectObject(_aadef.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _gbace != nil {
		_daddfg, _ebeae := _gbace.(*_abf.PdfObjectInteger)
		if !_ebeae {
			return nil, _fd.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0069\u006d\u0061\u0067\u0065\u0020\u0068\u0065\u0069g\u0068\u0074\u0020o\u0062j\u0065\u0063\u0074")
		}
		_accef := int64(*_daddfg)
		_ecdcf.BitsPerComponent = &_accef
	}
	_ecdcf.Intent = _aadef.Get("\u0049\u006e\u0074\u0065\u006e\u0074")
	_ecdcf.ImageMask = _aadef.Get("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k")
	_ecdcf.Mask = _aadef.Get("\u004d\u0061\u0073\u006b")
	_ecdcf.Decode = _aadef.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	_ecdcf.Interpolate = _aadef.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065")
	_ecdcf.Alternatives = _aadef.Get("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073")
	_ecdcf.SMask = _aadef.Get("\u0053\u004d\u0061s\u006b")
	_ecdcf.SMaskInData = _aadef.Get("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061")
	_ecdcf.Matte = _aadef.Get("\u004d\u0061\u0074t\u0065")
	_ecdcf.Name = _aadef.Get("\u004e\u0061\u006d\u0065")
	_ecdcf.StructParent = _aadef.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_ecdcf.ID = _aadef.Get("\u0049\u0044")
	_ecdcf.OPI = _aadef.Get("\u004f\u0050\u0049")
	_ecdcf.Metadata = _aadef.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	_ecdcf.OC = _aadef.Get("\u004f\u0043")
	_ecdcf.Stream = stream.Stream
	return _ecdcf, nil
}

func (_bbded *PdfWriter) makeOffSetReference(_addefc int64) {
	_egdggc := _e.Sprintf("\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u000a\u0025\u0064\u000a", _addefc)
	_bbded.writeString(_egdggc)
	_bbded.writeString("\u0025\u0025\u0045\u004f\u0046\u000a")
}

// ToPdfObject implements interface PdfModel.
func (_fddba *PdfSignatureReference) ToPdfObject() _abf.PdfObject {
	_deceb := _abf.MakeDict()
	_deceb.SetIfNotNil("\u0054\u0079\u0070\u0065", _fddba.Type)
	_deceb.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064", _fddba.TransformMethod)
	_deceb.SetIfNotNil("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073", _fddba.TransformParams)
	_deceb.SetIfNotNil("\u0044\u0061\u0074\u0061", _fddba.Data)
	_deceb.SetIfNotNil("\u0044\u0069\u0067e\u0073\u0074\u004d\u0065\u0074\u0068\u006f\u0064", _fddba.DigestMethod)
	return _deceb
}

func (_gd *PdfReader) newPdfActionRenditionFromDict(_fff *_abf.PdfObjectDictionary) (*PdfActionRendition, error) {
	return &PdfActionRendition{R: _fff.Get("\u0052"), AN: _fff.Get("\u0041\u004e"), OP: _fff.Get("\u004f\u0050"), JS: _fff.Get("\u004a\u0053")}, nil
}

// ImageToRGB returns an error since an image cannot be defined in a pattern colorspace.
func (_gcbc *PdfColorspaceSpecialPattern) ImageToRGB(img Image) (Image, error) {
	_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0061\u0074\u0074\u0065\u0072n \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	return img, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020\u0066\u006f\u0072\u0020\u0069m\u0061\u0067\u0065\u0020\u0028p\u0061\u0074t\u0065\u0072\u006e\u0029")
}

// ToPdfObject implements interface PdfModel.
func (_faef *PdfAnnotationCaret) ToPdfObject() _abf.PdfObject {
	_faef.PdfAnnotation.ToPdfObject()
	_dgb := _faef._dbc
	_fcdb := _dgb.PdfObject.(*_abf.PdfObjectDictionary)
	_faef.PdfAnnotationMarkup.appendToPdfDictionary(_fcdb)
	_fcdb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0043\u0061\u0072e\u0074"))
	_fcdb.SetIfNotNil("\u0052\u0044", _faef.RD)
	_fcdb.SetIfNotNil("\u0053\u0079", _faef.Sy)
	return _dgb
}

// GetNumComponents returns the number of color components (1 for Indexed).
func (_gcccc *PdfColorspaceSpecialIndexed) GetNumComponents() int { return 1 }

// PdfTilingPattern is a Tiling pattern that consists of repetitions of a pattern cell with defined intervals.
// It is a type 1 pattern. (PatternType = 1).
// A tiling pattern is represented by a stream object, where the stream content is
// a content stream that describes the pattern cell.
type PdfTilingPattern struct {
	*PdfPattern
	PaintType  *_abf.PdfObjectInteger
	TilingType *_abf.PdfObjectInteger
	BBox       *PdfRectangle
	XStep      *_abf.PdfObjectFloat
	YStep      *_abf.PdfObjectFloat
	Resources  *PdfPageResources
	Matrix     *_abf.PdfObjectArray
}

func _gbdd(_gafbe *PdfAnnotation) (*XObjectForm, *PdfRectangle, error) {
	_gbdfc, _gfad := _abf.GetDict(_gafbe.AP)
	if !_gfad {
		return nil, nil, _fd.New("f\u0069\u0065\u006c\u0064\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0041\u0050\u0020d\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079")
	}
	if _gbdfc == nil {
		return nil, nil, nil
	}
	_baeb, _gfad := _abf.GetArray(_gafbe.Rect)
	if !_gfad || _baeb.Len() != 4 {
		return nil, nil, _fd.New("\u0072\u0065\u0063t\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_ecddb, _cgdbe := NewPdfRectangle(*_baeb)
	if _cgdbe != nil {
		return nil, nil, _cgdbe
	}
	_dfda := _abf.TraceToDirectObject(_gbdfc.Get("\u004e"))
	switch _dcag := _dfda.(type) {
	case *_abf.PdfObjectStream:
		_gcdcg := _dcag
		_gfdcf, _edbde := NewXObjectFormFromStream(_gcdcg)
		return _gfdcf, _ecddb, _edbde
	case *_abf.PdfObjectDictionary:
		_fdbcab := _dcag
		_dcea, _bfcd := _abf.GetName(_gafbe.AS)
		if !_bfcd {
			return nil, nil, nil
		}
		if _fdbcab.Get(*_dcea) == nil {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0041\u0053\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
			return nil, nil, nil
		}
		_ggaa, _bfcd := _abf.GetStream(_fdbcab.Get(*_dcea))
		if !_bfcd {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055n\u0061\u0062\u006ce \u0074\u006f\u0020\u0061\u0063\u0063e\u0073\u0073\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0066\u006f\u0072 \u0025\u0076", _dcea)
			return nil, nil, _fd.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		}
		_edfd, _ggdgg := NewXObjectFormFromStream(_ggaa)
		return _edfd, _ecddb, _ggdgg
	}
	_acd.Log.Debug("\u0049\u006e\u0076\u0061li\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u004e\u003a\u0020%\u0054", _dfda)
	return nil, nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
}

// HasXObjectByName checks if an XObject with a specified keyName is defined.
func (_gbdac *PdfPageResources) HasXObjectByName(keyName _abf.PdfObjectName) bool {
	_fedbf, _ := _gbdac.GetXObjectByName(keyName)
	return _fedbf != nil
}

func (_ggefd *PdfReader) loadAnnotations(_ddfab _abf.PdfObject) ([]*PdfAnnotation, error) {
	_aafcb, _fdbcabd := _abf.GetArray(_ddfab)
	if !_fdbcabd {
		return nil, _e.Errorf("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	var _abaee []*PdfAnnotation
	for _, _dedd := range _aafcb.Elements() {
		_dedd = _abf.ResolveReference(_dedd)
		if _, _cafee := _dedd.(*_abf.PdfObjectNull); _cafee {
			continue
		}
		_afdf, _bfgg := _dedd.(*_abf.PdfObjectDictionary)
		_efdag, _cbcdd := _dedd.(*_abf.PdfIndirectObject)
		if _bfgg {
			_efdag = &_abf.PdfIndirectObject{}
			_efdag.PdfObject = _afdf
		} else {
			if !_cbcdd {
				return nil, _e.Errorf("\u0061\u006eno\u0074\u0061\u0074i\u006f\u006e\u0020\u006eot \u0069n \u0061\u006e\u0020\u0069\u006e\u0064\u0069re\u0063\u0074\u0020\u006f\u0062\u006a\u0065c\u0074")
			}
		}
		_edbdg, _bbacd := _ggefd.newPdfAnnotationFromIndirectObject(_efdag)
		if _bbacd != nil {
			return nil, _bbacd
		}
		switch _gegcgf := _edbdg.GetContext().(type) {
		case *PdfAnnotationWidget:
			for _, _abge := range _ggefd.AcroForm.AllFields() {
				if _abge._dgdc == _gegcgf.Parent {
					_gegcgf._agdc = _abge
					break
				}
			}
		}
		if _edbdg != nil {
			_abaee = append(_abaee, _edbdg)
		}
	}
	return _abaee, nil
}

// PageCallback callback function used in page loading
// that could be used to modify the page content.
//
// Deprecated: will be removed in v4. Use PageProcessCallback instead.
type PageCallback func(_fcgeg int, _cfdfc *PdfPage)

// HasExtGState checks if ExtGState name is available.
func (_edbgc *PdfPage) HasExtGState(name _abf.PdfObjectName) bool {
	if _edbgc.Resources == nil {
		return false
	}
	if _edbgc.Resources.ExtGState == nil {
		return false
	}
	_afga, _bbcfg := _abf.TraceToDirectObject(_edbgc.Resources.ExtGState).(*_abf.PdfObjectDictionary)
	if !_bbcfg {
		_acd.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _abf.TraceToDirectObject(_edbgc.Resources.ExtGState))
		return false
	}
	_fagecf := _afga.Get(name)
	_cdcd := _fagecf != nil
	return _cdcd
}

func _bgab(_cbab _abf.PdfObject) (*PdfColorspaceDeviceNAttributes, error) {
	_fbaeg := &PdfColorspaceDeviceNAttributes{}
	var _gfcef *_abf.PdfObjectDictionary
	switch _ccbg := _cbab.(type) {
	case *_abf.PdfIndirectObject:
		_fbaeg._ddbdd = _ccbg
		var _gbfda bool
		_gfcef, _gbfda = _ccbg.PdfObject.(*_abf.PdfObjectDictionary)
		if !_gbfda {
			_acd.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
			return nil, _fd.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	case *_abf.PdfObjectDictionary:
		_gfcef = _ccbg
	case *_abf.PdfObjectReference:
		_cbcdg := _ccbg.Resolve()
		return _bgab(_cbcdg)
	default:
		_acd.Log.Error("\u0044\u0065\u0076\u0069c\u0065\u004e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065 \u0074\u0079\u0070\u0065\u0020\u0065\u0072r\u006f\u0072")
		return nil, _fd.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _beaa := _gfcef.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _beaa != nil {
		_eaac, _edbb := _abf.TraceToDirectObject(_beaa).(*_abf.PdfObjectName)
		if !_edbb {
			_acd.Log.Error("\u0044\u0065vi\u0063\u0065\u004e \u0061\u0074\u0074\u0072ibu\u0074e \u0053\u0075\u0062\u0074\u0079\u0070\u0065 t\u0079\u0070\u0065\u0020\u0065\u0072\u0072o\u0072")
			return nil, _fd.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_fbaeg.Subtype = _eaac
	}
	if _dgfd := _gfcef.Get("\u0043o\u006c\u006f\u0072\u0061\u006e\u0074s"); _dgfd != nil {
		_fbaeg.Colorants = _dgfd
	}
	if _fgbc := _gfcef.Get("\u0050r\u006f\u0063\u0065\u0073\u0073"); _fgbc != nil {
		_fbaeg.Process = _fgbc
	}
	if _gaegb := _gfcef.Get("M\u0069\u0078\u0069\u006e\u0067\u0048\u0069\u006e\u0074\u0073"); _gaegb != nil {
		_fbaeg.MixingHints = _gaegb
	}
	return _fbaeg, nil
}

// Optimizer is the interface that performs optimization of PDF object structure for output writing.
//
// Optimize receives a slice of input `objects`, performs optimization, including removing, replacing objects and
// output the optimized slice of objects.
type Optimizer interface {
	Optimize(_fbfdc []_abf.PdfObject) ([]_abf.PdfObject, error)
}

func _dgbgc(_afacb *_abf.PdfObjectDictionary) (*PdfFieldButton, error) {
	_gecfa := &PdfFieldButton{}
	_gecfa.PdfField = NewPdfField()
	_gecfa.PdfField.SetContext(_gecfa)
	_gecfa.Opt, _ = _abf.GetArray(_afacb.Get("\u004f\u0070\u0074"))
	_gdab := NewPdfAnnotationWidget()
	_gdab.A, _ = _abf.GetDict(_afacb.Get("\u0041"))
	_gdab.AP, _ = _abf.GetDict(_afacb.Get("\u0041\u0050"))
	_gdab.SetContext(_gecfa)
	_gecfa.PdfField.Annotations = append(_gecfa.PdfField.Annotations, _gdab)
	return _gecfa, nil
}

// NewPdfColorspaceLab returns a new Lab colorspace object.
func NewPdfColorspaceLab() *PdfColorspaceLab {
	_fabf := &PdfColorspaceLab{}
	_fabf.BlackPoint = []float64{0.0, 0.0, 0.0}
	_fabf.Range = []float64{-100, 100, -100, 100}
	return _fabf
}

// GetAllContentStreams gets all the content streams for a page as one string.
func (_bebab *PdfPage) GetAllContentStreams() (string, error) {
	_afedg, _adffd := _bebab.GetContentStreams()
	if _adffd != nil {
		return "", _adffd
	}
	return _be.Join(_afedg, "\u0020"), nil
}

// PdfColorCalGray represents a CalGray colorspace.
type PdfColorCalGray float64

// NewPdfAnnotationHighlight returns a new text highlight annotation.
func NewPdfAnnotationHighlight() *PdfAnnotationHighlight {
	_geg := NewPdfAnnotation()
	_bac := &PdfAnnotationHighlight{}
	_bac.PdfAnnotation = _geg
	_bac.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_geg.SetContext(_bac)
	return _bac
}

// ToPdfObject returns a *PdfIndirectObject containing a *PdfObjectArray representation of the DeviceN colorspace.
/*
	Format: [/DeviceN names alternateSpace tintTransform]
	    or: [/DeviceN names alternateSpace tintTransform attributes]
*/
func (_ggdg *PdfColorspaceDeviceN) ToPdfObject() _abf.PdfObject {
	_cgf := _abf.MakeArray(_abf.MakeName("\u0044e\u0076\u0069\u0063\u0065\u004e"))
	_cgf.Append(_ggdg.ColorantNames)
	_cgf.Append(_ggdg.AlternateSpace.ToPdfObject())
	_cgf.Append(_ggdg.TintTransform.ToPdfObject())
	if _ggdg.Attributes != nil {
		_cgf.Append(_ggdg.Attributes.ToPdfObject())
	}
	if _ggdg._ddee != nil {
		_ggdg._ddee.PdfObject = _cgf
		return _ggdg._ddee
	}
	return _cgf
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
	_bfed        *PdfAppender
	_dgfe        *DSS
}

// DecodeArray returns the component range values for the DeviceN colorspace.
// [0 1.0 0 1.0 ...] for each color component.
func (_efbf *PdfColorspaceDeviceN) DecodeArray() []float64 {
	var _fbfb []float64
	for _bcbb := 0; _bcbb < _efbf.GetNumComponents(); _bcbb++ {
		_fbfb = append(_fbfb, 0.0, 1.0)
	}
	return _fbfb
}

// NewPdfFilespecFromObj creates and returns a new PdfFilespec object.
func NewPdfFilespecFromObj(obj _abf.PdfObject) (*PdfFilespec, error) {
	_fgde := &PdfFilespec{}
	var _acecf *_abf.PdfObjectDictionary
	if _gcdga, _acgcgc := _abf.GetIndirect(obj); _acgcgc {
		_fgde._badbg = _gcdga
		_gcegf, _fbdd := _abf.GetDict(_gcdga.PdfObject)
		if !_fbdd {
			_acd.Log.Debug("\u004f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0074y\u0070\u0065")
			return nil, _abf.ErrTypeError
		}
		_acecf = _gcegf
	} else if _bdcc, _ebcdg := _abf.GetDict(obj); _ebcdg {
		_fgde._badbg = _bdcc
		_acecf = _bdcc
	} else {
		_acd.Log.Debug("O\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0075\u006e\u0065\u0078\u0070e\u0063\u0074\u0065d\u0020(\u0025\u0054\u0029", obj)
		return nil, _abf.ErrTypeError
	}
	if _acecf == nil {
		_acd.Log.Debug("\u0044i\u0063t\u0069\u006f\u006e\u0061\u0072y\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, _fd.New("\u0064\u0069\u0063t\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	if _dgdf := _acecf.Get("\u0054\u0079\u0070\u0065"); _dgdf != nil {
		_ggdb, _ebga := _dgdf.(*_abf.PdfObjectName)
		if !_ebga {
			_acd.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _dgdf)
		} else {
			if *_ggdb != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
				_acd.Log.Trace("\u0055\u006e\u0073\u0075\u0073\u0070e\u0063\u0074\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063 \u0028\u0025\u0073\u0029", *_ggdb)
			}
		}
	}
	if _ebaad := _acecf.Get("\u0046\u0053"); _ebaad != nil {
		_fgde.FS = _ebaad
	}
	if _eadec := _acecf.Get("\u0046"); _eadec != nil {
		_fgde.F = _eadec
	}
	if _dcbfd := _acecf.Get("\u0055\u0046"); _dcbfd != nil {
		_fgde.UF = _dcbfd
	}
	if _dgfb := _acecf.Get("\u0044\u004f\u0053"); _dgfb != nil {
		_fgde.DOS = _dgfb
	}
	if _fbff := _acecf.Get("\u004d\u0061\u0063"); _fbff != nil {
		_fgde.Mac = _fbff
	}
	if _bacdd := _acecf.Get("\u0055\u006e\u0069\u0078"); _bacdd != nil {
		_fgde.Unix = _bacdd
	}
	if _agdec := _acecf.Get("\u0049\u0044"); _agdec != nil {
		_fgde.ID = _agdec
	}
	if _dbda := _acecf.Get("\u0056"); _dbda != nil {
		_fgde.V = _dbda
	}
	if _bebbc := _acecf.Get("\u0045\u0046"); _bebbc != nil {
		_fgde.EF = _bebbc
	}
	if _bebdf := _acecf.Get("\u0052\u0046"); _bebdf != nil {
		_fgde.RF = _bebdf
	}
	if _efce := _acecf.Get("\u0044\u0065\u0073\u0063"); _efce != nil {
		_fgde.Desc = _efce
	}
	if _bcbg := _acecf.Get("\u0043\u0049"); _bcbg != nil {
		_fgde.CI = _bcbg
	}
	return _fgde, nil
}

func (_dbf *PdfReader) newPdfActionJavaScriptFromDict(_aac *_abf.PdfObjectDictionary) (*PdfActionJavaScript, error) {
	return &PdfActionJavaScript{JS: _aac.Get("\u004a\u0053")}, nil
}

// NewPdfActionTrans returns a new "trans" action.
func NewPdfActionTrans() *PdfActionTrans {
	_cd := NewPdfAction()
	_cfbf := &PdfActionTrans{}
	_cfbf.PdfAction = _cd
	_cd.SetContext(_cfbf)
	return _cfbf
}

// GetDescent returns the Descent of the font `descriptor`.
func (_ffagg *PdfFontDescriptor) GetDescent() (float64, error) {
	return _abf.GetNumberAsFloat(_ffagg.Descent)
}

// AddImageResource adds an image to the XObject resources.
func (_dcbbc *PdfPage) AddImageResource(name _abf.PdfObjectName, ximg *XObjectImage) error {
	var _acfa *_abf.PdfObjectDictionary
	if _dcbbc.Resources.XObject == nil {
		_acfa = _abf.MakeDict()
		_dcbbc.Resources.XObject = _acfa
	} else {
		var _cgcca bool
		_acfa, _cgcca = (_dcbbc.Resources.XObject).(*_abf.PdfObjectDictionary)
		if !_cgcca {
			return _fd.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0078\u0072\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020\u0074\u0079p\u0065")
		}
	}
	_acfa.Set(name, ximg.ToPdfObject())
	return nil
}

// NewPdfPageResourcesFromDict creates and returns a new PdfPageResources object
// from the input dictionary.
func NewPdfPageResourcesFromDict(dict *_abf.PdfObjectDictionary) (*PdfPageResources, error) {
	_ddege := NewPdfPageResources()
	if _bcgbd := dict.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"); _bcgbd != nil {
		_ddege.ExtGState = _bcgbd
	}
	if _gbfbc := dict.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"); _gbfbc != nil && !_abf.IsNullObject(_gbfbc) {
		_ddege.ColorSpace = _gbfbc
	}
	if _ffbgb := dict.Get("\u0050a\u0074\u0074\u0065\u0072\u006e"); _ffbgb != nil {
		_ddege.Pattern = _ffbgb
	}
	if _geeg := dict.Get("\u0053h\u0061\u0064\u0069\u006e\u0067"); _geeg != nil {
		_ddege.Shading = _geeg
	}
	if _egabe := dict.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"); _egabe != nil {
		_ddege.XObject = _egabe
	}
	if _fdcce := _abf.ResolveReference(dict.Get("\u0046\u006f\u006e\u0074")); _fdcce != nil {
		_ddege.Font = _fdcce
	}
	if _dbbfe := dict.Get("\u0050r\u006f\u0063\u0053\u0065\u0074"); _dbbfe != nil {
		_ddege.ProcSet = _dbbfe
	}
	if _dfddf := dict.Get("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"); _dfddf != nil {
		_ddege.Properties = _dfddf
	}
	return _ddege, nil
}

const (
	XObjectTypeUndefined XObjectType = iota
	XObjectTypeImage
	XObjectTypeForm
	XObjectTypePS
	XObjectTypeUnknown
)

func (_eccab *pdfFontSimple) baseFields() *fontCommon { return &_eccab.fontCommon }
func (_ageaf *PdfShading) getShadingDict() (*_abf.PdfObjectDictionary, error) {
	_ceecc := _ageaf._eabcgc
	if _gegd, _efead := _ceecc.(*_abf.PdfIndirectObject); _efead {
		_eeffa, _gaaab := _gegd.PdfObject.(*_abf.PdfObjectDictionary)
		if !_gaaab {
			return nil, _abf.ErrTypeError
		}
		return _eeffa, nil
	} else if _geee, _dgcce := _ceecc.(*_abf.PdfObjectStream); _dgcce {
		return _geee.PdfObjectDictionary, nil
	} else if _bdcf, _dbge := _ceecc.(*_abf.PdfObjectDictionary); _dbge {
		return _bdcf, nil
	} else {
		_acd.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0063\u0063\u0065s\u0073\u0020\u0073\u0068\u0061\u0064\u0069n\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079")
		return nil, _abf.ErrTypeError
	}
}

// ToPdfOutline returns a low level PdfOutline object, based on the current
// instance.
func (_bccg *Outline) ToPdfOutline() *PdfOutline {
	_fdbeb := NewPdfOutline()
	var _dcacb []*PdfOutlineItem
	var _fbcc int64
	var _daafe *PdfOutlineItem
	for _, _aeadg := range _bccg.Entries {
		_fcfbd, _fabbc := _aeadg.ToPdfOutlineItem()
		_fcfbd.Parent = &_fdbeb.PdfOutlineTreeNode
		if _daafe != nil {
			_daafe.Next = &_fcfbd.PdfOutlineTreeNode
			_fcfbd.Prev = &_daafe.PdfOutlineTreeNode
		}
		_dcacb = append(_dcacb, _fcfbd)
		_fbcc += _fabbc
		_daafe = _fcfbd
	}
	_abca := int64(len(_dcacb))
	_fbcc += _abca
	if _abca > 0 {
		_fdbeb.First = &_dcacb[0].PdfOutlineTreeNode
		_fdbeb.Last = &_dcacb[_abca-1].PdfOutlineTreeNode
		_fdbeb.Count = &_fbcc
	}
	return _fdbeb
}

func (_gbcca *PdfReader) loadStructure() error {
	if _gbcca._bebc.GetCrypter() != nil && !_gbcca._bebc.IsAuthenticated() {
		return _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_cegd := _gbcca._bebc.GetTrailer()
	if _cegd == nil {
		return _e.Errorf("\u006di\u0073s\u0069\u006e\u0067\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
	}
	_fgfe, _gefeed := _cegd.Get("\u0052\u006f\u006f\u0074").(*_abf.PdfObjectReference)
	if !_gefeed {
		return _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052\u006f\u006ft\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a \u0025\u0073\u0029", _cegd)
	}
	_dggbc, _fbad := _gbcca._bebc.LookupByReference(*_fgfe)
	if _fbad != nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0072\u006f\u006f\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _fbad)
		return _fbad
	}
	_bdef, _gefeed := _dggbc.(*_abf.PdfIndirectObject)
	if !_gefeed {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0028\u0072\u006f\u006f\u0074\u0020\u0025\u0071\u0029\u0020\u0028\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0025\u0073\u0029", _dggbc, *_cegd)
		return _fd.New("\u006di\u0073s\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_dcbga, _gefeed := (*_bdef).PdfObject.(*_abf.PdfObjectDictionary)
	if !_gefeed {
		_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0061t\u0061\u006c\u006fg\u0020(\u0025\u0073\u0029", _bdef.PdfObject)
		return _fd.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067")
	}
	_acd.Log.Trace("C\u0061\u0074\u0061\u006c\u006f\u0067\u003a\u0020\u0025\u0073", _dcbga)
	_dfce, _gefeed := _dcbga.Get("\u0050\u0061\u0067e\u0073").(*_abf.PdfObjectReference)
	if !_gefeed {
		return _fd.New("\u0070\u0061\u0067\u0065\u0073\u0020\u0069\u006e\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0061\u0020\u0072e\u0066\u0065r\u0065\u006e\u0063\u0065")
	}
	_eefg, _fbad := _gbcca._bebc.LookupByReference(*_dfce)
	if _fbad != nil {
		_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020F\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020r\u0065\u0061\u0064 \u0070a\u0067\u0065\u0073")
		return _fbad
	}
	_acggd, _gefeed := _eefg.(*_abf.PdfIndirectObject)
	if !_gefeed {
		_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020P\u0061\u0067\u0065\u0073\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0069n\u0076a\u006c\u0069\u0064")
		_acd.Log.Debug("\u006f\u0070\u003a\u0020\u0025\u0070", _acggd)
		return _fd.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_edeaeb, _gefeed := _acggd.PdfObject.(*_abf.PdfObjectDictionary)
	if !_gefeed {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065c\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0073\u0029", _acggd)
		return _fd.New("p\u0061g\u0065\u0073\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0069\u006e\u0076al\u0069\u0064")
	}
	_bcgcf, _gefeed := _abf.GetInt(_edeaeb.Get("\u0043\u006f\u0075n\u0074"))
	if !_gefeed {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0050\u0061\u0067\u0065\u0073\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return _fd.New("\u0070\u0061\u0067\u0065s \u0063\u006f\u0075\u006e\u0074\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _, _gefeed = _abf.GetName(_edeaeb.Get("\u0054\u0079\u0070\u0065")); !_gefeed {
		_acd.Log.Debug("\u0050\u0061\u0067\u0065\u0073\u0020\u0064\u0069\u0063\u0074\u0020T\u0079\u0070\u0065\u0020\u0066\u0069\u0065\u006cd\u0020n\u006f\u0074\u0020\u0073\u0065\u0074\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0054\u0079p\u0065\u0020\u0074\u006f\u0020\u0050\u0061\u0067\u0065\u0073\u002e")
		_edeaeb.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0050\u0061\u0067e\u0073"))
	}
	if _efcf, _cbbca := _abf.GetInt(_edeaeb.Get("\u0052\u006f\u0074\u0061\u0074\u0065")); _cbbca {
		_fgbg := int64(*_efcf)
		_gbcca.Rotate = &_fgbg
	}
	_gbcca._afdaf = _fgfe
	_gbcca._dagde = _dcbga
	_gbcca._agbecg = _edeaeb
	_gbcca._bfdff = _acggd
	_gbcca._gcegc = int(*_bcgcf)
	_gbcca._gbfaf = []*_abf.PdfIndirectObject{}
	_bfege := map[_abf.PdfObject]struct{}{}
	_fbad = _gbcca.buildPageList(_acggd, nil, _bfege)
	if _fbad != nil {
		return _fbad
	}
	_acd.Log.Trace("\u002d\u002d\u002d")
	_acd.Log.Trace("\u0054\u004f\u0043")
	_acd.Log.Trace("\u0050\u0061\u0067e\u0073")
	_acd.Log.Trace("\u0025\u0064\u003a\u0020\u0025\u0073", len(_gbcca._gbfaf), _gbcca._gbfaf)
	_gbcca._cggee, _fbad = _gbcca.loadOutlines()
	if _fbad != nil {
		_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0062\u0075i\u006c\u0064\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065 t\u0072\u0065\u0065 \u0028%\u0073\u0029", _fbad)
		return _fbad
	}
	_gbcca.AcroForm, _fbad = _gbcca.loadForms()
	if _fbad != nil {
		return _fbad
	}
	_gbcca.DSS, _fbad = _gbcca.loadDSS()
	if _fbad != nil {
		return _fbad
	}
	_gbcca._gedbg, _fbad = _gbcca.loadPerms()
	if _fbad != nil {
		return _fbad
	}
	return nil
}

// DefaultImageHandler is the default implementation of the ImageHandler using the standard go library.
type DefaultImageHandler struct{}

// ColorFromPdfObjects gets the color from a series of pdf objects (4 for cmyk).
func (_fbc *PdfColorspaceDeviceCMYK) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 4 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cfbb, _cege := _abf.GetNumbersAsFloat(objects)
	if _cege != nil {
		return nil, _cege
	}
	return _fbc.ColorFromFloats(_cfbb)
}

// NewPdfActionLaunch returns a new "launch" action.
func NewPdfActionLaunch() *PdfActionLaunch {
	_gea := NewPdfAction()
	_gcf := &PdfActionLaunch{}
	_gcf.PdfAction = _gea
	_gea.SetContext(_gcf)
	return _gcf
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
func (_ceade *PdfReader) FlattenFields(allannots bool, appgen FieldAppearanceGenerator) error {
	return _ceade.flattenFieldsWithOpts(allannots, appgen, nil)
}

// NewPdfOutlineItem returns an initialized PdfOutlineItem.
func NewPdfOutlineItem() *PdfOutlineItem {
	_dbcgb := &PdfOutlineItem{_ceegc: _abf.MakeIndirectObject(_abf.MakeDict())}
	_dbcgb._aecec = _dbcgb
	return _dbcgb
}

// SetImageHandler sets the image handler used by the package.
func SetImageHandler(imgHandling ImageHandler) { ImageHandling = imgHandling }

// SetPdfCreator sets the Creator attribute of the output PDF.
func SetPdfCreator(creator string)            { _gaabd.Lock(); defer _gaabd.Unlock(); _edead = creator }
func _bedce(_ddcga *fontCommon) *pdfFontType0 { return &pdfFontType0{fontCommon: *_ddcga} }

// PdfActionRendition represents a Rendition action.
type PdfActionRendition struct {
	*PdfAction
	R  _abf.PdfObject
	AN _abf.PdfObject
	OP _abf.PdfObject
	JS _abf.PdfObject
}

// NewPdfAnnotationSound returns a new sound annotation.
func NewPdfAnnotationSound() *PdfAnnotationSound {
	_gdf := NewPdfAnnotation()
	_cfg := &PdfAnnotationSound{}
	_cfg.PdfAnnotation = _gdf
	_cfg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gdf.SetContext(_cfg)
	return _cfg
}

// SetOpenAction sets the OpenAction in the PDF catalog.
// The value shall be either an array defining a destination (12.3.2 "Destinations" PDF32000_2008),
// or an action dictionary representing an action (12.6 "Actions" PDF32000_2008).
func (_cgabc *PdfWriter) SetOpenAction(dest _abf.PdfObject) error {
	if dest == nil || _abf.IsNullObject(dest) {
		return nil
	}
	_cgabc._ddffc.Set("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e", dest)
	return _cgabc.addObjects(dest)
}

func (_geecc *PdfWriter) writeString(_cgccd string) {
	if _geecc._dacaeg != nil {
		return
	}
	_dcda, _cfegf := _geecc._agfba.WriteString(_cgccd)
	_geecc._dbfaad += int64(_dcda)
	_geecc._dacaeg = _cfegf
}

// GetPrimitiveFromModel returns the primitive object corresponding to the input `model`.
func (_agacc *modelManager) GetPrimitiveFromModel(model PdfModel) _abf.PdfObject {
	_bacc, _eecf := _agacc._baecg[model]
	if !_eecf {
		return nil
	}
	return _bacc
}

// ToPdfObject returns a stream object.
func (_efbfd *XObjectForm) ToPdfObject() _abf.PdfObject {
	_gfgca := _efbfd._dbba
	_fddbg := _gfgca.PdfObjectDictionary
	if _efbfd.Filter != nil {
		_fddbg = _efbfd.Filter.MakeStreamDict()
		_gfgca.PdfObjectDictionary = _fddbg
	}
	_fddbg.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_fddbg.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0046\u006f\u0072\u006d"))
	_fddbg.SetIfNotNil("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065", _efbfd.FormType)
	_fddbg.SetIfNotNil("\u0042\u0042\u006f\u0078", _efbfd.BBox)
	_fddbg.SetIfNotNil("\u004d\u0061\u0074\u0072\u0069\u0078", _efbfd.Matrix)
	if _efbfd.Resources != nil {
		_fddbg.SetIfNotNil("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s", _efbfd.Resources.ToPdfObject())
	}
	_fddbg.SetIfNotNil("\u0047\u0072\u006fu\u0070", _efbfd.Group)
	_fddbg.SetIfNotNil("\u0052\u0065\u0066", _efbfd.Ref)
	_fddbg.SetIfNotNil("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061", _efbfd.MetaData)
	_fddbg.SetIfNotNil("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o", _efbfd.PieceInfo)
	_fddbg.SetIfNotNil("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064", _efbfd.LastModified)
	_fddbg.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _efbfd.StructParent)
	_fddbg.SetIfNotNil("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073", _efbfd.StructParents)
	_fddbg.SetIfNotNil("\u004f\u0050\u0049", _efbfd.OPI)
	_fddbg.SetIfNotNil("\u004f\u0043", _efbfd.OC)
	_fddbg.SetIfNotNil("\u004e\u0061\u006d\u0065", _efbfd.Name)
	_fddbg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _abf.MakeInteger(int64(len(_efbfd.Stream))))
	_gfgca.Stream = _efbfd.Stream
	return _gfgca
}

// EnableByName LTV enables the signature dictionary of the PDF AcroForm
// field identified the specified name. The signing certificate chain is
// extracted from the signature dictionary. Optionally, additional certificates
// can be specified through the `extraCerts` parameter. The LTV client attempts
// to build the certificate chain up to a trusted root by downloading any
// missing certificates.
func (_cfecd *LTV) EnableByName(name string, extraCerts []*_fa.Certificate) error {
	_ceebdd := _cfecd._bfed._agda.AcroForm
	for _, _cdfbd := range _ceebdd.AllFields() {
		_edfcg, _ := _cdfbd.GetContext().(*PdfFieldSignature)
		if _edfcg == nil {
			continue
		}
		if _gfdbf := _edfcg.PartialName(); _gfdbf != name {
			continue
		}
		return _cfecd.Enable(_edfcg.V, extraCerts)
	}
	return nil
}

// PdfFontDescriptor specifies metrics and other attributes of a font and can refer to a FontFile
// for embedded fonts.
// 9.8 Font Descriptors (page 281)
type PdfFontDescriptor struct {
	FontName     _abf.PdfObject
	FontFamily   _abf.PdfObject
	FontStretch  _abf.PdfObject
	FontWeight   _abf.PdfObject
	Flags        _abf.PdfObject
	FontBBox     _abf.PdfObject
	ItalicAngle  _abf.PdfObject
	Ascent       _abf.PdfObject
	Descent      _abf.PdfObject
	Leading      _abf.PdfObject
	CapHeight    _abf.PdfObject
	XHeight      _abf.PdfObject
	StemV        _abf.PdfObject
	StemH        _abf.PdfObject
	AvgWidth     _abf.PdfObject
	MaxWidth     _abf.PdfObject
	MissingWidth _abf.PdfObject
	FontFile     _abf.PdfObject
	FontFile2    _abf.PdfObject
	FontFile3    _abf.PdfObject
	CharSet      _abf.PdfObject
	_bgbdf       int
	_fgccc       float64
	*fontFile
	_fcdf *_gbe.TtfType

	// Additional entries for CIDFonts
	Style  _abf.PdfObject
	Lang   _abf.PdfObject
	FD     _abf.PdfObject
	CIDSet _abf.PdfObject
	_aage  *_abf.PdfIndirectObject
}

// ToPdfObject implements interface PdfModel.
func (_bgf *PdfActionSubmitForm) ToPdfObject() _abf.PdfObject {
	_bgf.PdfAction.ToPdfObject()
	_age := _bgf._egg
	_gbf := _age.PdfObject.(*_abf.PdfObjectDictionary)
	_gbf.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeSubmitForm)))
	if _bgf.F != nil {
		_gbf.Set("\u0046", _bgf.F.ToPdfObject())
	}
	_gbf.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _bgf.Fields)
	_gbf.SetIfNotNil("\u0046\u006c\u0061g\u0073", _bgf.Flags)
	return _age
}

// ToPdfObject implements interface PdfModel.
func (_ccga *PdfFilespec) ToPdfObject() _abf.PdfObject {
	_gcfdd := _ccga.getDict()
	_gcfdd.Clear()
	_gcfdd.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063"))
	_gcfdd.SetIfNotNil("\u0046\u0053", _ccga.FS)
	_gcfdd.SetIfNotNil("\u0046", _ccga.F)
	_gcfdd.SetIfNotNil("\u0055\u0046", _ccga.UF)
	_gcfdd.SetIfNotNil("\u0044\u004f\u0053", _ccga.DOS)
	_gcfdd.SetIfNotNil("\u004d\u0061\u0063", _ccga.Mac)
	_gcfdd.SetIfNotNil("\u0055\u006e\u0069\u0078", _ccga.Unix)
	_gcfdd.SetIfNotNil("\u0049\u0044", _ccga.ID)
	_gcfdd.SetIfNotNil("\u0056", _ccga.V)
	_gcfdd.SetIfNotNil("\u0045\u0046", _ccga.EF)
	_gcfdd.SetIfNotNil("\u0052\u0046", _ccga.RF)
	_gcfdd.SetIfNotNil("\u0044\u0065\u0073\u0063", _ccga.Desc)
	_gcfdd.SetIfNotNil("\u0043\u0049", _ccga.CI)
	return _ccga._badbg
}

func (_aabcd *PdfWriter) writeDocumentVersion() {
	if _aabcd._aegbd {
		_aabcd.writeString("\u000a")
	} else {
		_aabcd.writeString(_e.Sprintf("\u0025\u0025\u0050D\u0046\u002d\u0025\u0064\u002e\u0025\u0064\u000a", _aabcd._ecfa.Major, _aabcd._ecfa.Minor))
		_aabcd.writeString("\u0025\u00e2\u00e3\u00cf\u00d3\u000a")
	}
}

// SetReason sets the `Reason` field of the signature.
func (_deagb *PdfSignature) SetReason(reason string) {
	_deagb.Reason = _abf.MakeEncodedString(reason, true)
}

var _abfb = map[string]struct{}{"\u0054\u0069\u0074l\u0065": {}, "\u0041\u0075\u0074\u0068\u006f\u0072": {}, "\u0053u\u0062\u006a\u0065\u0063\u0074": {}, "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073": {}, "\u0043r\u0065\u0061\u0074\u006f\u0072": {}, "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072": {}, "\u0054r\u0061\u0070\u0070\u0065\u0064": {}, "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065": {}, "\u004do\u0064\u0044\u0061\u0074\u0065": {}}

// IsShading specifies if the pattern is a shading pattern.
func (_degcc *PdfPattern) IsShading() bool { return _degcc.PatternType == 2 }

// B returns the value of the B component of the color.
func (_bcag *PdfColorLab) B() float64 { return _bcag[2] }

type fontCommon struct {
	_ecggf string
	_aacbc string
	_dddac string
	_dabca _abf.PdfObject
	_aabfe *_bd.CMap
	_dcbaf *PdfFontDescriptor
	_bgbd  int64
}

// PdfOutlineItem represents an outline item dictionary (Table 153 - pp. 376 - 377).
type PdfOutlineItem struct {
	PdfOutlineTreeNode
	Title  *_abf.PdfObjectString
	Parent *PdfOutlineTreeNode
	Prev   *PdfOutlineTreeNode
	Next   *PdfOutlineTreeNode
	Count  *int64
	Dest   _abf.PdfObject
	A      _abf.PdfObject
	SE     _abf.PdfObject
	C      _abf.PdfObject
	F      _abf.PdfObject
	_ceegc *_abf.PdfIndirectObject
}

// NewPdfShadingType2 creates an empty shading type 2 dictionary.
func NewPdfShadingType2() *PdfShadingType2 {
	_ebgac := &PdfShadingType2{}
	_ebgac.PdfShading = &PdfShading{}
	_ebgac.PdfShading._eabcgc = _abf.MakeIndirectObject(_abf.MakeDict())
	_ebgac.PdfShading._eabd = _ebgac
	return _ebgac
}

// VariableText contains the common attributes of a variable text.
// The VariableText is typically not used directly, but is can encapsulate by PdfField
// See section 12.7.3.3 "Variable Text" and Table 222 (pp. 434-436 PDF32000_2008).
type VariableText struct {
	DA *_abf.PdfObjectString
	Q  *_abf.PdfObjectInteger
	DS *_abf.PdfObjectString
	RV _abf.PdfObject
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_agafc *PdfShadingType4) ToPdfObject() _abf.PdfObject {
	_agafc.PdfShading.ToPdfObject()
	_daae, _cbdfe := _agafc.getShadingDict()
	if _cbdfe != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _agafc.BitsPerCoordinate != nil {
		_daae.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _agafc.BitsPerCoordinate)
	}
	if _agafc.BitsPerComponent != nil {
		_daae.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _agafc.BitsPerComponent)
	}
	if _agafc.BitsPerFlag != nil {
		_daae.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _agafc.BitsPerFlag)
	}
	if _agafc.Decode != nil {
		_daae.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _agafc.Decode)
	}
	if _agafc.Function != nil {
		if len(_agafc.Function) == 1 {
			_daae.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _agafc.Function[0].ToPdfObject())
		} else {
			_cegfd := _abf.MakeArray()
			for _, _abgbb := range _agafc.Function {
				_cegfd.Append(_abgbb.ToPdfObject())
			}
			_daae.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _cegfd)
		}
	}
	return _agafc._eabcgc
}

func (_gfgc *Image) resampleLowBits(_dbed []uint32) {
	_fbbbc := _gca.BytesPerLine(int(_gfgc.Width), int(_gfgc.BitsPerComponent), _gfgc.ColorComponents)
	_bcbc := make([]byte, _gfgc.ColorComponents*_fbbbc*int(_gfgc.Height))
	_gdgcc := int(_gfgc.BitsPerComponent) * _gfgc.ColorComponents * int(_gfgc.Width)
	_agcd := uint8(8)
	var (
		_aaagg, _bfdce int
		_gbfdc         uint32
	)
	for _cbdg := 0; _cbdg < int(_gfgc.Height); _cbdg++ {
		_bfdce = _cbdg * _fbbbc
		for _eddee := 0; _eddee < _gdgcc; _eddee++ {
			_gbfdc = _dbed[_aaagg]
			_agcd -= uint8(_gfgc.BitsPerComponent)
			_bcbc[_bfdce] |= byte(_gbfdc) << _agcd
			if _agcd == 0 {
				_agcd = 8
				_bfdce++
			}
			_aaagg++
		}
	}
	_gfgc.Data = _bcbc
}

// NewPdfColorspaceCalGray returns a new CalGray colorspace object.
func NewPdfColorspaceCalGray() *PdfColorspaceCalGray {
	_gaeb := &PdfColorspaceCalGray{}
	_gaeb.BlackPoint = []float64{0.0, 0.0, 0.0}
	_gaeb.Gamma = 1
	return _gaeb
}

// GetAnnotations returns the list of page annotations for `page`. If not loaded attempts to load the
// annotations, otherwise returns the loaded list.
func (_gccg *PdfPage) GetAnnotations() ([]*PdfAnnotation, error) {
	if _gccg._baagf != nil {
		return _gccg._baagf, nil
	}
	if _gccg.Annots == nil {
		_gccg._baagf = []*PdfAnnotation{}
		return nil, nil
	}
	if _gccg._dbaef == nil {
		_gccg._baagf = []*PdfAnnotation{}
		return nil, nil
	}
	_ebfbg, _aaaf := _gccg._dbaef.loadAnnotations(_gccg.Annots)
	if _aaaf != nil {
		return nil, _aaaf
	}
	if _ebfbg == nil {
		_gccg._baagf = []*PdfAnnotation{}
	}
	_gccg._baagf = _ebfbg
	return _gccg._baagf, nil
}

// SetContext sets the sub pattern (context).  Either PdfTilingPattern or PdfShadingPattern.
func (_ccab *PdfPattern) SetContext(ctx PdfModel) { _ccab._bgafe = ctx }

// AppendContentBytes creates a PDF stream from `cs` and appends it to the
// array of streams specified by the pages's Contents entry.
// If `wrapContents` is true, the content stream of the page is wrapped using
// a `q/Q` operator pair, so that its state does not affect the appended
// content stream.
func (_ggbde *PdfPage) AppendContentBytes(cs []byte, wrapContents bool) error {
	_beedeg := _ggbde.GetContentStreamObjs()
	wrapContents = wrapContents && len(_beedeg) > 0
	_ebfbga := _abf.NewFlateEncoder()
	_ffbec := _abf.MakeArray()
	if wrapContents {
		_ebbee, _cebgf := _abf.MakeStream([]byte("\u0071\u000a"), _ebfbga)
		if _cebgf != nil {
			return _cebgf
		}
		_ffbec.Append(_ebbee)
	}
	_ffbec.Append(_beedeg...)
	if wrapContents {
		_aacg, _fffdeb := _abf.MakeStream([]byte("\u000a\u0051\u000a"), _ebfbga)
		if _fffdeb != nil {
			return _fffdeb
		}
		_ffbec.Append(_aacg)
	}
	_aafg, _dffee := _abf.MakeStream(cs, _ebfbga)
	if _dffee != nil {
		return _dffee
	}
	_ffbec.Append(_aafg)
	_ggbde.Contents = _ffbec
	return nil
}

// PdfColorspaceDeviceRGB represents an RGB colorspace.
type PdfColorspaceDeviceRGB struct{}

func (_eda *PdfReader) newPdfAnnotationCaretFromDict(_gcea *_abf.PdfObjectDictionary) (*PdfAnnotationCaret, error) {
	_afef := PdfAnnotationCaret{}
	_gbgaa, _fcgg := _eda.newPdfAnnotationMarkupFromDict(_gcea)
	if _fcgg != nil {
		return nil, _fcgg
	}
	_afef.PdfAnnotationMarkup = _gbgaa
	_afef.RD = _gcea.Get("\u0052\u0044")
	_afef.Sy = _gcea.Get("\u0053\u0079")
	return &_afef, nil
}

// ToPdfObject implements interface PdfModel.
func (_gab *PdfActionResetForm) ToPdfObject() _abf.PdfObject {
	_gab.PdfAction.ToPdfObject()
	_ebf := _gab._egg
	_ebc := _ebf.PdfObject.(*_abf.PdfObjectDictionary)
	_ebc.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeResetForm)))
	_ebc.SetIfNotNil("\u0046\u0069\u0065\u006c\u0064\u0073", _gab.Fields)
	_ebc.SetIfNotNil("\u0046\u006c\u0061g\u0073", _gab.Flags)
	return _ebf
}

// ToJBIG2Image converts current image to the core.JBIG2Image.
func (_dgge *Image) ToJBIG2Image() (*_abf.JBIG2Image, error) {
	_baba, _adbc := _dgge.ToGoImage()
	if _adbc != nil {
		return nil, _adbc
	}
	return _abf.GoImageToJBIG2(_baba, _abf.JB2ImageAutoThreshold)
}

func _fcce(_gffb _abf.PdfObject) (*PdfColorspaceSpecialPattern, error) {
	_acd.Log.Trace("\u004e\u0065\u0077\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u0043\u0053\u0020\u0066\u0072\u006fm\u0020\u006f\u0062\u006a\u003a\u0020\u0025\u0073\u0020\u0025\u0054", _gffb.String(), _gffb)
	_gbgc := NewPdfColorspaceSpecialPattern()
	if _cfaff, _fcddg := _gffb.(*_abf.PdfIndirectObject); _fcddg {
		_gbgc._afca = _cfaff
	}
	_gffb = _abf.TraceToDirectObject(_gffb)
	if _ggfc, _cada := _gffb.(*_abf.PdfObjectName); _cada {
		if *_ggfc != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			return nil, _e.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
		return _gbgc, nil
	}
	_gddd, _acfgb := _gffb.(*_abf.PdfObjectArray)
	if !_acfgb {
		_acd.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061t\u0074\u0065\u0072\u006e\u0020\u0043\u0053 \u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0023\u0076", _gffb)
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0050\u0061\u0074\u0074e\u0072n\u0020C\u0053\u0020\u006f\u0062\u006a\u0065\u0063t")
	}
	if _gddd.Len() != 1 && _gddd.Len() != 2 {
		_acd.Log.Error("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079\u003a %\u0023\u0076", _gddd)
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065r\u006e\u0020\u0043\u0053\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_gffb = _gddd.Get(0)
	if _aegf, _debd := _gffb.(*_abf.PdfObjectName); _debd {
		if *_aegf != "\u0050a\u0074\u0074\u0065\u0072\u006e" {
			_acd.Log.Error("\u0049\u006e\u0076al\u0069\u0064\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020C\u0053 \u0061r\u0072\u0061\u0079\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0023\u0076", _aegf)
			return nil, _e.Errorf("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u006e\u0061\u006d\u0065")
		}
	}
	if _gddd.Len() > 1 {
		_gffb = _gddd.Get(1)
		_gffb = _abf.TraceToDirectObject(_gffb)
		_acdac, _gcda := NewPdfColorspaceFromPdfObject(_gffb)
		if _gcda != nil {
			return nil, _gcda
		}
		_gbgc.UnderlyingCS = _acdac
	}
	_acd.Log.Trace("R\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0050\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0077i\u0074\u0068\u0020\u0075\u006e\u0064\u0065\u0072\u006c\u0079in\u0067\u0020\u0063s\u003a \u0025\u0054", _gbgc.UnderlyingCS)
	return _gbgc, nil
}

// PdfActionGoToR represents a GoToR action.
type PdfActionGoToR struct {
	*PdfAction
	F         *PdfFilespec
	D         _abf.PdfObject
	NewWindow _abf.PdfObject
}

// ImageToRGB converts image in CalGray color space to RGB (A, B, C -> X, Y, Z).
func (_ebdd *PdfColorspaceCalGray) ImageToRGB(img Image) (Image, error) {
	_egbc := _gf.NewReader(img.getBase())
	_dded := _gca.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), 3, nil, nil, nil)
	_ccbb := _gf.NewWriter(_dded)
	_gdaa := _ge.Pow(2, float64(img.BitsPerComponent)) - 1
	_bfcc := make([]uint32, 3)
	var (
		_egae                                 uint32
		ANorm, X, Y, Z, _gfefd, _dbbd, _cfecf float64
		_baab                                 error
	)
	for {
		_egae, _baab = _egbc.ReadSample()
		if _baab == _gc.EOF {
			break
		} else if _baab != nil {
			return img, _baab
		}
		ANorm = float64(_egae) / _gdaa
		X = _ebdd.WhitePoint[0] * _ge.Pow(ANorm, _ebdd.Gamma)
		Y = _ebdd.WhitePoint[1] * _ge.Pow(ANorm, _ebdd.Gamma)
		Z = _ebdd.WhitePoint[2] * _ge.Pow(ANorm, _ebdd.Gamma)
		_gfefd = 3.240479*X + -1.537150*Y + -0.498535*Z
		_dbbd = -0.969256*X + 1.875992*Y + 0.041556*Z
		_cfecf = 0.055648*X + -0.204043*Y + 1.057311*Z
		_gfefd = _ge.Min(_ge.Max(_gfefd, 0), 1.0)
		_dbbd = _ge.Min(_ge.Max(_dbbd, 0), 1.0)
		_cfecf = _ge.Min(_ge.Max(_cfecf, 0), 1.0)
		_bfcc[0] = uint32(_gfefd * _gdaa)
		_bfcc[1] = uint32(_dbbd * _gdaa)
		_bfcc[2] = uint32(_cfecf * _gdaa)
		if _baab = _ccbb.WriteSamples(_bfcc); _baab != nil {
			return img, _baab
		}
	}
	return _cega(&_dded), nil
}

// PdfColorDeviceCMYK is a CMYK32 color, where each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorDeviceCMYK [4]float64

func _cgddd(_daagf *_abf.PdfObjectStream) (*PdfFunctionType4, error) {
	_agfb := &PdfFunctionType4{}
	_agfb._dgbdc = _daagf
	_befe := _daagf.PdfObjectDictionary
	_dgbca, _gggfc := _abf.TraceToDirectObject(_befe.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_abf.PdfObjectArray)
	if !_gggfc {
		_acd.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _dgbca.Len()%2 != 0 {
		_acd.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _fd.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_fegdc, _feece := _dgbca.ToFloat64Array()
	if _feece != nil {
		return nil, _feece
	}
	_agfb.Domain = _fegdc
	_dgbca, _gggfc = _abf.TraceToDirectObject(_befe.Get("\u0052\u0061\u006eg\u0065")).(*_abf.PdfObjectArray)
	if _gggfc {
		if _dgbca.Len() < 0 || _dgbca.Len()%2 != 0 {
			return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_ffabb, _eceg := _dgbca.ToFloat64Array()
		if _eceg != nil {
			return nil, _eceg
		}
		_agfb.Range = _ffabb
	}
	_bbdff, _feece := _abf.DecodeStream(_daagf)
	if _feece != nil {
		return nil, _feece
	}
	_agfb._cddgf = _bbdff
	_dbgf := _ae.NewPSParser(_bbdff)
	_dbdge, _feece := _dbgf.Parse()
	if _feece != nil {
		return nil, _feece
	}
	_agfb.Program = _dbdge
	return _agfb, nil
}

func (_bcdga *Image) samplesTrimPadding(_gfbada []uint32) []uint32 {
	_geeaf := _bcdga.ColorComponents * int(_bcdga.Width) * int(_bcdga.Height)
	if len(_gfbada) == _geeaf {
		return _gfbada
	}
	_bggd := make([]uint32, _geeaf)
	_ebefd := int(_bcdga.Width) * _bcdga.ColorComponents
	var _bdba, _adbfa, _deged, _bgbab int
	_dfeda := _gca.BytesPerLine(int(_bcdga.Width), int(_bcdga.BitsPerComponent), _bcdga.ColorComponents)
	for _bdba = 0; _bdba < int(_bcdga.Height); _bdba++ {
		_adbfa = _bdba * int(_bcdga.Width)
		_deged = _bdba * _dfeda
		for _bgbab = 0; _bgbab < _ebefd; _bgbab++ {
			_bggd[_adbfa+_bgbab] = _gfbada[_deged+_bgbab]
		}
	}
	return _bggd
}

// NewXObjectImage returns a new XObjectImage.
func NewXObjectImage() *XObjectImage {
	_cdbbgg := &XObjectImage{}
	_bfffa := &_abf.PdfObjectStream{}
	_bfffa.PdfObjectDictionary = _abf.MakeDict()
	_cdbbgg._ccbad = _bfffa
	return _cdbbgg
}

// AddOCSPs adds OCSPs to DSS.
func (_aagf *DSS) AddOCSPs(ocsps [][]byte) ([]*_abf.PdfObjectStream, error) {
	return _aagf.add(&_aagf.OCSPs, _aagf._ggfg, ocsps)
}

// Encoder returns the font's text encoder.
func (_dgagg *pdfFontSimple) Encoder() _cbb.TextEncoder {
	if _dgagg._ebada != nil {
		return _dgagg._ebada
	}
	if _dgagg._edabc != nil {
		return _dgagg._edabc
	}
	_cbdde, _ := _cbb.NewSimpleTextEncoder("\u0053\u0074a\u006e\u0064\u0061r\u0064\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", nil)
	return _cbdde
}

// ToPdfObject returns the PDF representation of the function.
func (_bdgff *PdfFunctionType4) ToPdfObject() _abf.PdfObject {
	_bafcb := _bdgff._dgbdc
	if _bafcb == nil {
		_bdgff._dgbdc = &_abf.PdfObjectStream{}
		_bafcb = _bdgff._dgbdc
	}
	_efabe := _abf.MakeDict()
	_efabe.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _abf.MakeInteger(4))
	_beaec := &_abf.PdfObjectArray{}
	for _, _cafg := range _bdgff.Domain {
		_beaec.Append(_abf.MakeFloat(_cafg))
	}
	_efabe.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _beaec)
	_ffdfg := &_abf.PdfObjectArray{}
	for _, _dcbec := range _bdgff.Range {
		_ffdfg.Append(_abf.MakeFloat(_dcbec))
	}
	_efabe.Set("\u0052\u0061\u006eg\u0065", _ffdfg)
	if _bdgff._cddgf == nil && _bdgff.Program != nil {
		_bdgff._cddgf = []byte(_bdgff.Program.String())
	}
	_efabe.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _abf.MakeInteger(int64(len(_bdgff._cddgf))))
	_bafcb.Stream = _bdgff._cddgf
	_bafcb.PdfObjectDictionary = _efabe
	return _bafcb
}

// ColorToRGB converts gray -> rgb for a single color component.
func (_aefed *PdfColorspaceDeviceGray) ColorToRGB(color PdfColor) (PdfColor, error) {
	_eecb, _edcdb := color.(*PdfColorDeviceGray)
	if !_edcdb {
		_acd.Log.Debug("\u0049\u006e\u0070\u0075\u0074\u0020\u0063\u006f\u006c\u006fr\u0020\u006e\u006f\u0074\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0067\u0072\u0061\u0079\u0020\u0025\u0054", color)
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return NewPdfColorDeviceRGB(float64(*_eecb), float64(*_eecb), float64(*_eecb)), nil
}

// ToPdfObject implements interface PdfModel.
func (_gbgag *PdfAnnotationSquiggly) ToPdfObject() _abf.PdfObject {
	_gbgag.PdfAnnotation.ToPdfObject()
	_bcgf := _gbgag._dbc
	_dgcf := _bcgf.PdfObject.(*_abf.PdfObjectDictionary)
	_gbgag.PdfAnnotationMarkup.appendToPdfDictionary(_dgcf)
	_dgcf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0053\u0071\u0075\u0069\u0067\u0067\u006c\u0079"))
	_dgcf.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _gbgag.QuadPoints)
	return _bcgf
}

// NewPdfAcroForm returns a new PdfAcroForm with an initialized container (indirect object).
func NewPdfAcroForm() *PdfAcroForm {
	return &PdfAcroForm{Fields: &[]*PdfField{}, _bgfc: _abf.MakeIndirectObject(_abf.MakeDict())}
}

// GetOutlineTree returns the outline tree.
func (_fcceb *PdfReader) GetOutlineTree() *PdfOutlineTreeNode { return _fcceb._cggee }

// ImageToRGB converts an image with samples in Separation CS to an image with samples specified in
// DeviceRGB CS.
func (_egfd *PdfColorspaceSpecialSeparation) ImageToRGB(img Image) (Image, error) {
	_fcbbd := _gf.NewReader(img.getBase())
	_dedbb := _gca.NewImageBase(int(img.Width), int(img.Height), int(img.BitsPerComponent), _egfd.AlternateSpace.GetNumComponents(), nil, img._gedg, nil)
	_fffde := _gf.NewWriter(_dedbb)
	_gagd := _ge.Pow(2, float64(img.BitsPerComponent)) - 1
	_acd.Log.Trace("\u0053\u0065\u0070a\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u002d\u003e\u0020\u0054\u006f\u0052\u0047\u0042\u0020\u0063o\u006e\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	_acd.Log.Trace("\u0054i\u006et\u0054\u0072\u0061\u006e\u0073f\u006f\u0072m\u003a\u0020\u0025\u002b\u0076", _egfd.TintTransform)
	_cbgf := _egfd.AlternateSpace.DecodeArray()
	var (
		_efef  uint32
		_dbabg error
	)
	for {
		_efef, _dbabg = _fcbbd.ReadSample()
		if _dbabg == _gc.EOF {
			break
		}
		if _dbabg != nil {
			return img, _dbabg
		}
		_aefac := float64(_efef) / _gagd
		_cegba, _ggcag := _egfd.TintTransform.Evaluate([]float64{_aefac})
		if _ggcag != nil {
			return img, _ggcag
		}
		for _fbda, _abea := range _cegba {
			_bbagc := _gca.LinearInterpolate(_abea, _cbgf[_fbda*2], _cbgf[_fbda*2+1], 0, 1)
			if _ggcag = _fffde.WriteSample(uint32(_bbagc * _gagd)); _ggcag != nil {
				return img, _ggcag
			}
		}
	}
	return _egfd.AlternateSpace.ImageToRGB(_cega(&_dedbb))
}

func (_affeb *PdfSignature) extractChainFromPKCS7() ([]*_fa.Certificate, error) {
	_fffb, _ebedd := _eb.Parse(_affeb.Contents.Bytes())
	if _ebedd != nil {
		return nil, _ebedd
	}
	return _fffb.Certificates, nil
}

// GetTrailer returns the PDF's trailer dictionary.
func (_ecgcc *PdfReader) GetTrailer() (*_abf.PdfObjectDictionary, error) {
	_afad := _ecgcc._bebc.GetTrailer()
	if _afad == nil {
		return nil, _fd.New("\u0074r\u0061i\u006c\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	return _afad, nil
}

// PdfActionHide represents a hide action.
type PdfActionHide struct {
	*PdfAction
	T _abf.PdfObject
	H _abf.PdfObject
}

// PdfAnnotationPolygon represents Polygon annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolygon struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _abf.PdfObject
	LE       _abf.PdfObject
	BS       _abf.PdfObject
	IC       _abf.PdfObject
	BE       _abf.PdfObject
	IT       _abf.PdfObject
	Measure  _abf.PdfObject
}

func _ebeg(_aeeca *_abf.PdfObjectStream) (*PdfFunctionType0, error) {
	_eaed := &PdfFunctionType0{}
	_eaed._cabaa = _aeeca
	_gdbd := _aeeca.PdfObjectDictionary
	_abgcg, _addbc := _abf.TraceToDirectObject(_gdbd.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_abf.PdfObjectArray)
	if !_addbc {
		_acd.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _abgcg.Len() < 0 || _abgcg.Len()%2 != 0 {
		_acd.Log.Error("\u0044\u006f\u006d\u0061\u0069\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
		return nil, _fd.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_eaed.NumInputs = _abgcg.Len() / 2
	_bgbfb, _bfce := _abgcg.ToFloat64Array()
	if _bfce != nil {
		return nil, _bfce
	}
	_eaed.Domain = _bgbfb
	_abgcg, _addbc = _abf.TraceToDirectObject(_gdbd.Get("\u0052\u0061\u006eg\u0065")).(*_abf.PdfObjectArray)
	if !_addbc {
		_acd.Log.Error("\u0052\u0061\u006e\u0067e \u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _abgcg.Len() < 0 || _abgcg.Len()%2 != 0 {
		return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_eaed.NumOutputs = _abgcg.Len() / 2
	_cebd, _bfce := _abgcg.ToFloat64Array()
	if _bfce != nil {
		return nil, _bfce
	}
	_eaed.Range = _cebd
	_abgcg, _addbc = _abf.TraceToDirectObject(_gdbd.Get("\u0053\u0069\u007a\u0065")).(*_abf.PdfObjectArray)
	if !_addbc {
		_acd.Log.Error("\u0053i\u007ae\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_cgda, _bfce := _abgcg.ToIntegerArray()
	if _bfce != nil {
		return nil, _bfce
	}
	if len(_cgda) != _eaed.NumInputs {
		_acd.Log.Error("T\u0061\u0062\u006c\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063h\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072 o\u0066\u0020\u0069n\u0070u\u0074\u0073")
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_eaed.Size = _cgda
	_eeaa, _addbc := _abf.TraceToDirectObject(_gdbd.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065")).(*_abf.PdfObjectInteger)
	if !_addbc {
		_acd.Log.Error("B\u0069\u0074\u0073\u0050\u0065\u0072S\u0061\u006d\u0070\u006c\u0065\u0020\u006e\u006f\u0074 \u0073\u0070\u0065c\u0069f\u0069\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if *_eeaa != 1 && *_eeaa != 2 && *_eeaa != 4 && *_eeaa != 8 && *_eeaa != 12 && *_eeaa != 16 && *_eeaa != 24 && *_eeaa != 32 {
		_acd.Log.Error("\u0042\u0069\u0074s \u0070\u0065\u0072\u0020\u0073\u0061\u006d\u0070\u006ce\u0020o\u0075t\u0073i\u0064\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0028\u0025\u0064\u0029", *_eeaa)
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_eaed.BitsPerSample = int(*_eeaa)
	_eaed.Order = 1
	_efgfb, _addbc := _abf.TraceToDirectObject(_gdbd.Get("\u004f\u0072\u0064e\u0072")).(*_abf.PdfObjectInteger)
	if _addbc {
		if *_efgfb != 1 && *_efgfb != 3 {
			_acd.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u006fr\u0064\u0065r\u0020\u0028\u0025\u0064\u0029", *_efgfb)
			return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		_eaed.Order = int(*_efgfb)
	}
	_abgcg, _addbc = _abf.TraceToDirectObject(_gdbd.Get("\u0045\u006e\u0063\u006f\u0064\u0065")).(*_abf.PdfObjectArray)
	if _addbc {
		_bfdf, _aegg := _abgcg.ToFloat64Array()
		if _aegg != nil {
			return nil, _aegg
		}
		_eaed.Encode = _bfdf
	}
	_abgcg, _addbc = _abf.TraceToDirectObject(_gdbd.Get("\u0044\u0065\u0063\u006f\u0064\u0065")).(*_abf.PdfObjectArray)
	if _addbc {
		_deef, _dcbeb := _abgcg.ToFloat64Array()
		if _dcbeb != nil {
			return nil, _dcbeb
		}
		_eaed.Decode = _deef
	}
	_fbbbf, _bfce := _abf.DecodeStream(_aeeca)
	if _bfce != nil {
		return nil, _bfce
	}
	_eaed._aefbg = _fbbbf
	return _eaed, nil
}

var ImageHandling ImageHandler = DefaultImageHandler{}

func _dfefe(_adgbd _abf.PdfObject) (*PdfFontDescriptor, error) {
	_ccgca := &PdfFontDescriptor{}
	_adgbd = _abf.ResolveReference(_adgbd)
	if _baef, _ddeb := _adgbd.(*_abf.PdfIndirectObject); _ddeb {
		_ccgca._aage = _baef
		_adgbd = _baef.PdfObject
	}
	_fbfee, _fadcff := _abf.GetDict(_adgbd)
	if !_fadcff {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0044\u0065\u0073c\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _adgbd)
		return nil, _abf.ErrTypeError
	}
	if _ecgbf := _fbfee.Get("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065"); _ecgbf != nil {
		_ccgca.FontName = _ecgbf
	} else {
		_acd.Log.Debug("\u0049n\u0063\u006fm\u0070\u0061\u0074\u0069b\u0069\u006c\u0069t\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u004e\u0061me\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006di\u0073\u0073i\u006e\u0067")
	}
	_feadc, _ := _abf.GetName(_ccgca.FontName)
	if _fabfc := _fbfee.Get("\u0054\u0079\u0070\u0065"); _fabfc != nil {
		_agadb, _acbg := _fabfc.(*_abf.PdfObjectName)
		if !_acbg || string(*_agadb) != "\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072" {
			_acd.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072i\u0070t\u006f\u0072\u0020\u0054y\u0070\u0065 \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0025\u0054\u0029\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _fabfc, _feadc, _ccgca.FontName)
		}
	} else {
		_acd.Log.Trace("\u0049\u006ec\u006f\u006d\u0070\u0061\u0074i\u0062\u0069\u006c\u0069\u0074y\u003a\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071\u0020\u0025\u0054", _feadc, _ccgca.FontName)
	}
	_ccgca.FontFamily = _fbfee.Get("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079")
	_ccgca.FontStretch = _fbfee.Get("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068")
	_ccgca.FontWeight = _fbfee.Get("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074")
	_ccgca.Flags = _fbfee.Get("\u0046\u006c\u0061g\u0073")
	_ccgca.FontBBox = _fbfee.Get("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078")
	_ccgca.ItalicAngle = _fbfee.Get("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065")
	_ccgca.Ascent = _fbfee.Get("\u0041\u0073\u0063\u0065\u006e\u0074")
	_ccgca.Descent = _fbfee.Get("\u0044e\u0073\u0063\u0065\u006e\u0074")
	_ccgca.Leading = _fbfee.Get("\u004ce\u0061\u0064\u0069\u006e\u0067")
	_ccgca.CapHeight = _fbfee.Get("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t")
	_ccgca.XHeight = _fbfee.Get("\u0058H\u0065\u0069\u0067\u0068\u0074")
	_ccgca.StemV = _fbfee.Get("\u0053\u0074\u0065m\u0056")
	_ccgca.StemH = _fbfee.Get("\u0053\u0074\u0065m\u0048")
	_ccgca.AvgWidth = _fbfee.Get("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068")
	_ccgca.MaxWidth = _fbfee.Get("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068")
	_ccgca.MissingWidth = _fbfee.Get("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068")
	_ccgca.FontFile = _fbfee.Get("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065")
	_ccgca.FontFile2 = _fbfee.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652")
	_ccgca.FontFile3 = _fbfee.Get("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653")
	_ccgca.CharSet = _fbfee.Get("\u0043h\u0061\u0072\u0053\u0065\u0074")
	_ccgca.Style = _fbfee.Get("\u0053\u0074\u0079l\u0065")
	_ccgca.Lang = _fbfee.Get("\u004c\u0061\u006e\u0067")
	_ccgca.FD = _fbfee.Get("\u0046\u0044")
	_ccgca.CIDSet = _fbfee.Get("\u0043\u0049\u0044\u0053\u0065\u0074")
	if _ccgca.Flags != nil {
		if _efdd, _faba := _abf.GetIntVal(_ccgca.Flags); _faba {
			_ccgca._bgbdf = _efdd
		}
	}
	if _ccgca.MissingWidth != nil {
		if _cfaea, _cbff := _abf.GetNumberAsFloat(_ccgca.MissingWidth); _cbff == nil {
			_ccgca._fgccc = _cfaea
		}
	}
	if _ccgca.FontFile != nil {
		_fggc, _cfde := _gbbcga(_ccgca.FontFile)
		if _cfde != nil {
			return _ccgca, _cfde
		}
		_acd.Log.Trace("f\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u003d\u0025\u0073", _fggc)
		_ccgca.fontFile = _fggc
	}
	if _ccgca.FontFile2 != nil {
		_edcf, _cgfb := _gbe.NewFontFile2FromPdfObject(_ccgca.FontFile2)
		if _cgfb != nil {
			return _ccgca, _cgfb
		}
		_acd.Log.Trace("\u0066\u006f\u006et\u0046\u0069\u006c\u0065\u0032\u003d\u0025\u0073", _edcf.String())
		_ccgca._fcdf = &_edcf
	}
	return _ccgca, nil
}

// NewPdfColorspaceDeviceCMYK returns a new CMYK32 colorspace object.
func NewPdfColorspaceDeviceCMYK() *PdfColorspaceDeviceCMYK { return &PdfColorspaceDeviceCMYK{} }

// PdfActionGoTo3DView represents a GoTo3DView action.
type PdfActionGoTo3DView struct {
	*PdfAction
	TA _abf.PdfObject
	V  _abf.PdfObject
}

// NewPdfActionSubmitForm returns a new "submit form" action.
func NewPdfActionSubmitForm() *PdfActionSubmitForm {
	_gcae := NewPdfAction()
	_ea := &PdfActionSubmitForm{}
	_ea.PdfAction = _gcae
	_gcae.SetContext(_ea)
	return _ea
}

// GetPerms returns the Permissions dictionary
func (_fccac *PdfReader) GetPerms() *Permissions { return _fccac._gedbg }

func _addf(_dabdb _abf.PdfObject) (*_abf.PdfObjectDictionary, *fontCommon, error) {
	_feec := &fontCommon{}
	if _egea, _gdabe := _dabdb.(*_abf.PdfIndirectObject); _gdabe {
		_feec._bgbd = _egea.ObjectNumber
	}
	_egeff, _gffda := _abf.GetDict(_dabdb)
	if !_gffda {
		_acd.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0062\u0079\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _dabdb)
		return nil, nil, ErrFontNotSupported
	}
	_fgfd, _gffda := _abf.GetNameVal(_egeff.Get("\u0054\u0079\u0070\u0065"))
	if !_gffda {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046o\u006e\u0074\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	if _fgfd != "\u0046\u006f\u006e\u0074" {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0046\u006f\u006e\u0074\u0020\u0049\u006e\u0063\u006f\u006d\u0070\u0061t\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u002e\u0020\u0054\u0079\u0070\u0065\u003d\u0025\u0071\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0025\u0071.", _fgfd, "\u0046\u006f\u006e\u0074")
		return nil, nil, _abf.ErrTypeError
	}
	_abdf, _gffda := _abf.GetNameVal(_egeff.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_gffda {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020F\u006f\u006e\u0074 \u0049\u006e\u0063o\u006d\u0070a\u0074\u0069\u0062\u0069\u006c\u0069t\u0079. \u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, nil, ErrRequiredAttributeMissing
	}
	_feec._aacbc = _abdf
	_bbcc, _gffda := _abf.GetNameVal(_egeff.Get("\u004e\u0061\u006d\u0065"))
	if _gffda {
		_feec._dddac = _bbcc
	}
	_bafe := _egeff.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e")
	if _bafe != nil {
		_feec._dabca = _abf.TraceToDirectObject(_bafe)
		_edbc, _adcc := _cebb(_feec._dabca, _feec)
		if _adcc != nil {
			return _egeff, _feec, _adcc
		}
		_feec._aabfe = _edbc
	} else if _abdf == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030" || _abdf == "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		_bcggg, _dfbb := _bd.NewCIDSystemInfo(_egeff.Get("\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"))
		if _dfbb != nil {
			return _egeff, _feec, _dfbb
		}
		_cfdfd := _e.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0055\u0043\u0053\u0032", _bcggg.Registry, _bcggg.Ordering)
		if _bd.IsPredefinedCMap(_cfdfd) {
			_feec._aabfe, _dfbb = _bd.LoadPredefinedCMap(_cfdfd)
			if _dfbb != nil {
				_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064 \u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0043\u004d\u0061\u0070\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _cfdfd, _dfbb)
			}
		}
	}
	_dbcf := _egeff.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
	if _dbcf != nil {
		_bbaf, _fega := _dfefe(_dbcf)
		if _fega != nil {
			_acd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0042\u0061\u0064\u0020\u0066\u006f\u006et\u0020d\u0065s\u0063r\u0069\u0070\u0074\u006f\u0072\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fega)
			return _egeff, _feec, _fega
		}
		_feec._dcbaf = _bbaf
	}
	if _abdf != "\u0054\u0079\u0070e\u0033" {
		_dcecac, _affdc := _abf.GetNameVal(_egeff.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
		if !_affdc {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u006f\u006et\u0020\u0049\u006ec\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069t\u0079\u002e\u0020\u0042\u0061se\u0046\u006f\u006e\u0074\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0029\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
			return _egeff, _feec, ErrRequiredAttributeMissing
		}
		_feec._ecggf = _dcecac
	}
	return _egeff, _feec, nil
}

func (_afa *PdfColorspaceICCBased) String() string {
	return "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"
}

// SetColorSpace sets `r` colorspace object to `colorspace`.
func (_cggc *PdfPageResources) SetColorSpace(colorspace *PdfPageResourcesColorspaces) {
	_cggc._aafff = colorspace
}

// NewPdfReaderLazy creates a new PdfReader for `rs` in lazy-loading mode. The difference
// from NewPdfReader is that in lazy-loading mode, objects are only loaded into memory when needed
// rather than entire structure being loaded into memory on reader creation.
// Note that it may make sense to use the lazy-load reader when processing only parts of files,
// rather than loading entire file into memory. Example: splitting a few pages from a large PDF file.
func NewPdfReaderLazy(rs _gc.ReadSeeker) (*PdfReader, error) {
	const _afbge = "\u006d\u006f\u0064\u0065l:\u004e\u0065\u0077\u0050\u0064\u0066\u0052\u0065\u0061\u0064\u0065\u0072\u004c\u0061z\u0079"
	return _fbaec(rs, &ReaderOpts{LazyLoad: true}, false, _afbge)
}

// Inspect inspects the object types, subtypes and content in the PDF file returning a map of
// object type to number of instances of each.
func (_beade *PdfReader) Inspect() (map[string]int, error) { return _beade._bebc.Inspect() }

// Subtype returns the font's "Subtype" field.
func (_bgba *PdfFont) Subtype() string {
	_gecc := _bgba.baseFields()._aacbc
	if _acbad, _bgcd := _bgba._gedca.(*pdfFontType0); _bgcd {
		_gecc = _gecc + "\u003a" + _acbad.DescendantFont.Subtype()
	}
	return _gecc
}

var (
	ErrRequiredAttributeMissing = _fd.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074t\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
	ErrInvalidAttribute         = _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065")
	ErrTypeCheck                = _fd.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	_bgaaa                      = _fd.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrEncrypted                = _fd.New("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	ErrNoFont                   = _fd.New("\u0066\u006fn\u0074\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	ErrFontNotSupported         = _ddd.Errorf("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u006fn\u0074\u0020\u0028\u0025\u0077\u0029", _abf.ErrNotSupported)
	ErrType1CFontNotSupported   = _ddd.Errorf("\u0054y\u0070\u00651\u0043\u0020\u0066o\u006e\u0074\u0073\u0020\u0061\u0072\u0065 \u006e\u006f\u0074\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0028\u0025\u0077\u0029", _abf.ErrNotSupported)
	ErrType3FontNotSupported    = _ddd.Errorf("\u0054y\u0070\u00653\u0020\u0066\u006f\u006et\u0073\u0020\u0061r\u0065\u0020\u006e\u006f\u0074\u0020\u0063\u0075\u0072re\u006e\u0074\u006cy\u0020\u0073u\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0028%\u0077\u0029", _abf.ErrNotSupported)
	ErrTTCmapNotSupported       = _ddd.Errorf("\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065d\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070e\u0020\u0063\u006d\u0061\u0070\u0020\u0066\u006f\u0072\u006d\u0061\u0074 \u0028\u0025\u0077\u0029", _abf.ErrNotSupported)
	ErrSignNotEnoughSpace       = _ddd.Errorf("\u0069\u006e\u0073\u0075\u0066\u0066\u0069c\u0069\u0065\u006et\u0020\u0073\u0070a\u0063\u0065 \u0061\u006c\u006c\u006f\u0063\u0061t\u0065d \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
	ErrSignNoCertificates       = _ddd.Errorf("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065tri\u0065ve\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0063\u0068\u0061i\u006e")
)

// SetNameDictionary sets the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_gccf *PdfWriter) SetNameDictionary(names _abf.PdfObject) error {
	if names == nil {
		return nil
	}
	_acd.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u004e\u0061\u006d\u0065\u0073\u002e\u002e\u002e")
	_gccf._ddffc.Set("\u004e\u0061\u006de\u0073", names)
	return _gccf.addObjects(names)
}

// NewPdfAnnotationTrapNet returns a new trapnet annotation.
func NewPdfAnnotationTrapNet() *PdfAnnotationTrapNet {
	_gfdd := NewPdfAnnotation()
	_bgaa := &PdfAnnotationTrapNet{}
	_bgaa.PdfAnnotation = _gfdd
	_gfdd.SetContext(_bgaa)
	return _bgaa
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for an RGB device.
func (_cageb *PdfColorspaceDeviceRGB) GetNumComponents() int { return 3 }

// Items returns all children outline items.
func (_acbdb *OutlineItem) Items() []*OutlineItem { return _acbdb.Entries }

// PdfInfo holds document information that will overwrite
// document information global variables defined above.
type PdfInfo struct {
	Title        *_abf.PdfObjectString
	Author       *_abf.PdfObjectString
	Subject      *_abf.PdfObjectString
	Keywords     *_abf.PdfObjectString
	Creator      *_abf.PdfObjectString
	Producer     *_abf.PdfObjectString
	CreationDate *PdfDate
	ModifiedDate *PdfDate
	Trapped      *_abf.PdfObjectName
	_cbf         *_abf.PdfObjectDictionary
}

// NewPdfAnnotationSquare returns a new square annotation.
func NewPdfAnnotationSquare() *PdfAnnotationSquare {
	_dcgd := NewPdfAnnotation()
	_dccd := &PdfAnnotationSquare{}
	_dccd.PdfAnnotation = _dcgd
	_dccd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dcgd.SetContext(_dccd)
	return _dccd
}

// String returns a human readable description of `fontfile`.
func (_gcdee *fontFile) String() string {
	_eged := "\u005b\u004e\u006f\u006e\u0065\u005d"
	if _gcdee._eedb != nil {
		_eged = _gcdee._eedb.String()
	}
	return _e.Sprintf("\u0046O\u004e\u0054\u0046\u0049\u004c\u0045\u007b\u0025\u0023\u0071\u0020e\u006e\u0063\u006f\u0064\u0065\u0072\u003d\u0025\u0073\u007d", _gcdee._gadc, _eged)
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element between 0 and 1.
func (_bacg *PdfColorspaceDeviceGray) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gcdc := vals[0]
	if _gcdc < 0.0 || _gcdc > 1.0 {
		_acd.Log.Debug("\u0049\u006eco\u006d\u0070\u0061t\u0069\u0062\u0069\u006city\u003a R\u0061\u006e\u0067\u0065\u0020\u006f\u0075ts\u0069\u0064\u0065\u0020\u005b\u0030\u002c1\u005d")
	}
	if _gcdc < 0.0 {
		_gcdc = 0.0
	} else if _gcdc > 1.0 {
		_gcdc = 1.0
	}
	return NewPdfColorDeviceGray(_gcdc), nil
}

// PdfRectangle is a definition of a rectangle.
type PdfRectangle struct {
	Llx float64
	Lly float64
	Urx float64
	Ury float64
}

// Duplicate creates a duplicate page based on the current one and returns it.
func (_beca *PdfPage) Duplicate() *PdfPage {
	_egggf := *_beca
	_egggf._bdbfa = _abf.MakeDict()
	_egggf._gefee = _abf.MakeIndirectObject(_egggf._bdbfa)
	_egggf._efca = *_egggf._bdbfa
	return &_egggf
}

func (_ffdgb *PdfReader) newPdfAcroFormFromDict(_afefc *_abf.PdfIndirectObject, _cbdfba *_abf.PdfObjectDictionary) (*PdfAcroForm, error) {
	_afacg := NewPdfAcroForm()
	if _afefc != nil {
		_afacg._bgfc = _afefc
		_afefc.PdfObject = _abf.MakeDict()
	}
	if _aecee := _cbdfba.Get("\u0046\u0069\u0065\u006c\u0064\u0073"); _aecee != nil && !_abf.IsNullObject(_aecee) {
		_bdcb, _daab := _abf.GetArray(_aecee)
		if !_daab {
			return nil, _e.Errorf("\u0066i\u0065\u006c\u0064\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u0028\u0025\u0054\u0029", _aecee)
		}
		var _gabbg []*PdfField
		for _, _bgega := range _bdcb.Elements() {
			_fffg, _fagec := _abf.GetIndirect(_bgega)
			if !_fagec {
				if _, _gbdb := _bgega.(*_abf.PdfObjectNull); _gbdb {
					_acd.Log.Trace("\u0053k\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072 \u006e\u0075\u006c\u006c\u0020\u0066\u0069\u0065\u006c\u0064")
					continue
				}
				_acd.Log.Debug("\u0046\u0069\u0065\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0064 \u0069\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0025\u0054", _bgega)
				return nil, _e.Errorf("\u0066\u0069\u0065l\u0064\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u006e\u0020i\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_cffd, _beggb := _ffdgb.newPdfFieldFromIndirectObject(_fffg, nil)
			if _beggb != nil {
				return nil, _beggb
			}
			_acd.Log.Trace("\u0041\u0063\u0072\u006fFo\u0072\u006d\u0020\u0046\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u002b\u0076", *_cffd)
			_gabbg = append(_gabbg, _cffd)
		}
		_afacg.Fields = &_gabbg
	}
	if _ecabb := _cbdfba.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"); _ecabb != nil {
		_egcga, _fdcae := _abf.GetBool(_ecabb)
		if _fdcae {
			_afacg.NeedAppearances = _egcga
		} else {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u0065\u0065\u0064\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063e\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _ecabb)
		}
	}
	if _cgeb := _cbdfba.Get("\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073"); _cgeb != nil {
		_abfa, _gabda := _abf.GetInt(_cgeb)
		if _gabda {
			_afacg.SigFlags = _abfa
		} else {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u0046\u006c\u0061\u0067\u0073 \u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _cgeb)
		}
	}
	if _caaf := _cbdfba.Get("\u0043\u004f"); _caaf != nil {
		_beafb, _agdb := _abf.GetArray(_caaf)
		if _agdb {
			_afacg.CO = _beafb
		} else {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u004f\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _caaf)
		}
	}
	if _aegb := _cbdfba.Get("\u0044\u0052"); _aegb != nil {
		if _bbacf, _feefd := _abf.GetDict(_aegb); _feefd {
			_fedadc, _fbgg := NewPdfPageResourcesFromDict(_bbacf)
			if _fbgg != nil {
				_acd.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0044R\u003a\u0020\u0025\u0076", _fbgg)
				return nil, _fbgg
			}
			_afacg.DR = _fedadc
		} else {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0052\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _aegb)
		}
	}
	if _dcef := _cbdfba.Get("\u0044\u0041"); _dcef != nil {
		_edgff, _gbfce := _abf.GetString(_dcef)
		if _gbfce {
			_afacg.DA = _edgff
		} else {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0044\u0041\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", _dcef)
		}
	}
	if _afcde := _cbdfba.Get("\u0051"); _afcde != nil {
		_babf, _bbcdf := _abf.GetInt(_afcde)
		if _bbcdf {
			_afacg.Q = _babf
		} else {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0051\u0020\u0069\u006e\u0076a\u006ci\u0064 \u0028\u0067\u006f\u0074\u0020\u0025\u0054)", _afcde)
		}
	}
	if _cgcee := _cbdfba.Get("\u0058\u0046\u0041"); _cgcee != nil {
		_afacg.XFA = _cgcee
	}
	if _cecgd := _cbdfba.Get("\u0041\u0044\u0042\u0045\u005f\u0045\u0063\u0068\u006f\u0053\u0069\u0067\u006e"); _cecgd != nil {
		_afacg.ADBEEchoSign = _cecgd
	}
	_afacg.ToPdfObject()
	return _afacg, nil
}

// DecodeArray returns an empty slice as there are no components associated with pattern colorspace.
func (_ffegb *PdfColorspaceSpecialPattern) DecodeArray() []float64 { return []float64{} }

func _aabg(_eeab map[_gbe.GID]int, _geacac uint16) *_abf.PdfObjectArray {
	_badg := &_abf.PdfObjectArray{}
	_caag := _gbe.GID(_geacac)
	for _eadb := _gbe.GID(0); _eadb < _caag; {
		_dfdcd, _bdcgf := _eeab[_eadb]
		if !_bdcgf {
			_eadb++
			continue
		}
		_fefb := _eadb
		for _ceege := _fefb + 1; _ceege < _caag; _ceege++ {
			if _fcefc, _dggg := _eeab[_ceege]; !_dggg || _dfdcd != _fcefc {
				break
			}
			_fefb = _ceege
		}
		_badg.Append(_abf.MakeInteger(int64(_eadb)))
		_badg.Append(_abf.MakeInteger(int64(_fefb)))
		_badg.Append(_abf.MakeInteger(int64(_dfdcd)))
		_eadb = _fefb + 1
	}
	return _badg
}

// GetFontByName gets the font specified by keyName. Returns the PdfObject which
// the entry refers to. Returns a bool value indicating whether or not the entry was found.
func (_begcb *PdfPageResources) GetFontByName(keyName _abf.PdfObjectName) (_abf.PdfObject, bool) {
	if _begcb.Font == nil {
		return nil, false
	}
	_eccdg, _fggff := _abf.TraceToDirectObject(_begcb.Font).(*_abf.PdfObjectDictionary)
	if !_fggff {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0021\u0020(\u0067\u006ft\u0020\u0025\u0054\u0029", _abf.TraceToDirectObject(_begcb.Font))
		return nil, false
	}
	if _ffge := _eccdg.Get(keyName); _ffge != nil {
		return _ffge, true
	}
	return nil, false
}

// SignatureHandler interface defines the common functionality for PDF signature handlers, which
// need to be capable of validating digital signatures and signing PDF documents.
type SignatureHandler interface {
	// IsApplicable checks if a given signature dictionary `sig` is applicable for the signature handler.
	// For example a signature of type `adbe.pkcs7.detached` might not fit for a rsa.sha1 handler.
	IsApplicable(_ffcaf *PdfSignature) bool

	// Validate validates a PDF signature against a given digest (hash) such as that determined
	// for an input file. Returns validation results.
	Validate(_bcddbe *PdfSignature, _dcggg Hasher) (SignatureValidationResult, error)

	// InitSignature prepares the signature dictionary for signing. This involves setting all
	// necessary fields, and also allocating sufficient space to the Contents so that the
	// finalized signature can be inserted once the hash is calculated.
	InitSignature(_geadf *PdfSignature) error

	// NewDigest creates a new digest/hasher based on the signature dictionary and handler.
	NewDigest(_gagaae *PdfSignature) (Hasher, error)

	// Sign receives the hash `digest` (for example hash of an input file), and signs based
	// on the signature dictionary `sig` and applies the signature data to the signature
	// dictionary Contents field.
	Sign(_dcbcb *PdfSignature, _eefbg Hasher) error
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_gddaf pdfFontType0) GetCharMetrics(code _cbb.CharCode) (_gbe.CharMetrics, bool) {
	if _gddaf.DescendantFont == nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0064\u0065\u0073\u0063\u0065\u006e\u0064\u0061\u006e\u0074\u002e\u0020\u0066\u006f\u006et=\u0025\u0073", _gddaf)
		return _gbe.CharMetrics{}, false
	}
	return _gddaf.DescendantFont.GetCharMetrics(code)
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_bdead *PdfFunctionType2) Evaluate(x []float64) ([]float64, error) {
	if len(x) != 1 {
		_acd.Log.Error("\u004f\u006e\u006c\u0079 o\u006e\u0065\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0061\u006c\u006c\u006f\u0077e\u0064")
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fgdee := []float64{0.0}
	if _bdead.C0 != nil {
		_fgdee = _bdead.C0
	}
	_aabe := []float64{1.0}
	if _bdead.C1 != nil {
		_aabe = _bdead.C1
	}
	var _cdbbeb []float64
	for _fdcag := 0; _fdcag < len(_fgdee); _fdcag++ {
		_cbfd := _fgdee[_fdcag] + _ge.Pow(x[0], _bdead.N)*(_aabe[_fdcag]-_fgdee[_fdcag])
		_cdbbeb = append(_cdbbeb, _cbfd)
	}
	return _cdbbeb, nil
}

// NewPdfAnnotationLink returns a new link annotation.
func NewPdfAnnotationLink() *PdfAnnotationLink {
	_aebe := NewPdfAnnotation()
	_ebef := &PdfAnnotationLink{}
	_ebef.PdfAnnotation = _aebe
	_aebe.SetContext(_ebef)
	return _ebef
}

// CheckAccessRights checks access rights and permissions for a specified password.  If either user/owner
// password is specified,  full rights are granted, otherwise the access rights are specified by the
// Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_gceda *PdfReader) CheckAccessRights(password []byte) (bool, _bga.Permissions, error) {
	return _gceda._bebc.CheckAccessRights(password)
}

// ConvertToBinary converts current image into binary (bi-level) format.
// Binary images are composed of single bits per pixel (only black or white).
// If provided image has more color components, then it would be converted into binary image using
// histogram auto threshold function.
func (_beeeg *Image) ConvertToBinary() error {
	if _beeeg.ColorComponents == 1 && _beeeg.BitsPerComponent == 1 {
		return nil
	}
	_addd, _fbfcb := _beeeg.ToGoImage()
	if _fbfcb != nil {
		return _fbfcb
	}
	_fdcdc, _fbfcb := _gca.MonochromeConverter.Convert(_addd)
	if _fbfcb != nil {
		return _fbfcb
	}
	_beeeg.Data = _fdcdc.Base().Data
	_beeeg._gedg, _fbfcb = _gca.ScaleAlphaToMonochrome(_beeeg._gedg, int(_beeeg.Width), int(_beeeg.Height))
	if _fbfcb != nil {
		return _fbfcb
	}
	_beeeg.BitsPerComponent = 1
	_beeeg.ColorComponents = 1
	_beeeg._ceeag = nil
	return nil
}

// GetContext returns a reference to the subpattern entry: either PdfTilingPattern or PdfShadingPattern.
func (_fbcac *PdfPattern) GetContext() PdfModel { return _fbcac._bgafe }

func (_ccfae *LTV) getCRLs(_dccb []*_fa.Certificate) ([][]byte, error) {
	_ebbbf := make([][]byte, 0, len(_dccb))
	for _, _dgbe := range _dccb {
		for _, _fgeaf := range _dgbe.CRLDistributionPoints {
			if _ccfae.CertClient.IsCA(_dgbe) {
				continue
			}
			_ebea, _bged := _ccfae.CRLClient.MakeRequest(_fgeaf, _dgbe)
			if _bged != nil {
				_acd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _bged)
				continue
			}
			_ebbbf = append(_ebbbf, _ebea)
		}
	}
	return _ebbbf, nil
}

const (
	ButtonTypeCheckbox ButtonType = iota
	ButtonTypePush     ButtonType = iota
	ButtonTypeRadio    ButtonType = iota
)

// CharcodesToUnicode converts the character codes `charcodes` to a slice of runes.
// How it works:
//  1. Use the ToUnicode CMap if there is one.
//  2. Use the underlying font's encoding.
func (_ecdb *PdfFont) CharcodesToUnicode(charcodes []_cbb.CharCode) []rune {
	_agaf, _, _ := _ecdb.CharcodesToUnicodeWithStats(charcodes)
	return _agaf
}

// IsCID returns true if the underlying font is CID.
func (_gaef *PdfFont) IsCID() bool { return _gaef.baseFields().isCIDFont() }

// ToPdfObject returns colorspace in a PDF object format [name stream]
func (_dfad *PdfColorspaceICCBased) ToPdfObject() _abf.PdfObject {
	_cdge := &_abf.PdfObjectArray{}
	_cdge.Append(_abf.MakeName("\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064"))
	var _cecf *_abf.PdfObjectStream
	if _dfad._bfgc != nil {
		_cecf = _dfad._bfgc
	} else {
		_cecf = &_abf.PdfObjectStream{}
	}
	_dbeg := _abf.MakeDict()
	_dbeg.Set("\u004e", _abf.MakeInteger(int64(_dfad.N)))
	if _dfad.Alternate != nil {
		_dbeg.Set("\u0041l\u0074\u0065\u0072\u006e\u0061\u0074e", _dfad.Alternate.ToPdfObject())
	}
	if _dfad.Metadata != nil {
		_dbeg.Set("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _dfad.Metadata)
	}
	if _dfad.Range != nil {
		var _baad []_abf.PdfObject
		for _, _cbcb := range _dfad.Range {
			_baad = append(_baad, _abf.MakeFloat(_cbcb))
		}
		_dbeg.Set("\u0052\u0061\u006eg\u0065", _abf.MakeArray(_baad...))
	}
	_dbeg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _abf.MakeInteger(int64(len(_dfad.Data))))
	_cecf.Stream = _dfad.Data
	_cecf.PdfObjectDictionary = _dbeg
	_cdge.Append(_cecf)
	if _dfad._afcc != nil {
		_dfad._afcc.PdfObject = _cdge
		return _dfad._afcc
	}
	return _cdge
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain three PdfObjectFloat elements representing
// the L, A and B components of the color.
func (_ffcb *PdfColorspaceLab) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_cgab, _abgc := _abf.GetNumbersAsFloat(objects)
	if _abgc != nil {
		return nil, _abgc
	}
	return _ffcb.ColorFromFloats(_cgab)
}

func (_abcb *PdfReader) newPdfAnnotationPolygonFromDict(_fee *_abf.PdfObjectDictionary) (*PdfAnnotationPolygon, error) {
	_ebefa := PdfAnnotationPolygon{}
	_gece, _dade := _abcb.newPdfAnnotationMarkupFromDict(_fee)
	if _dade != nil {
		return nil, _dade
	}
	_ebefa.PdfAnnotationMarkup = _gece
	_ebefa.Vertices = _fee.Get("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073")
	_ebefa.LE = _fee.Get("\u004c\u0045")
	_ebefa.BS = _fee.Get("\u0042\u0053")
	_ebefa.IC = _fee.Get("\u0049\u0043")
	_ebefa.BE = _fee.Get("\u0042\u0045")
	_ebefa.IT = _fee.Get("\u0049\u0054")
	_ebefa.Measure = _fee.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	return &_ebefa, nil
}

func (_dcf *PdfReader) newPdfActionFromIndirectObject(_ebcd *_abf.PdfIndirectObject) (*PdfAction, error) {
	_cde, _fafb := _ebcd.PdfObject.(*_abf.PdfObjectDictionary)
	if !_fafb {
		return nil, _e.Errorf("\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u006e\u006f\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	if model := _dcf._ceecd.GetModelFromPrimitive(_cde); model != nil {
		_eab, _aea := model.(*PdfAction)
		if !_aea {
			return nil, _e.Errorf("\u0063\u0061c\u0068\u0065\u0064\u0020\u006d\u006f\u0064\u0065\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0061\u0063ti\u006f\u006e")
		}
		return _eab, nil
	}
	_ebdc := &PdfAction{}
	_ebdc._egg = _ebcd
	_dcf._ceecd.Register(_cde, _ebdc)
	if _ega := _cde.Get("\u0054\u0079\u0070\u0065"); _ega != nil {
		_feb, _bbe := _ega.(*_abf.PdfObjectName)
		if !_bbe {
			_acd.Log.Trace("\u0049\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074\u0079\u0021\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0054\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u004e\u0061m\u0065", _ega)
		} else {
			if *_feb != "\u0041\u0063\u0074\u0069\u006f\u006e" {
				_acd.Log.Trace("\u0055\u006e\u0073u\u0073\u0070\u0065\u0063t\u0065\u0064\u0020\u0054\u0079\u0070\u0065 \u0021\u003d\u0020\u0041\u0063\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", *_feb)
			}
			_ebdc.Type = _feb
		}
	}
	if _ged := _cde.Get("\u004e\u0065\u0078\u0074"); _ged != nil {
		_ebdc.Next = _ged
	}
	if _gfa := _cde.Get("\u0053"); _gfa != nil {
		_ebdc.S = _gfa
	}
	_dab, _gfc := _ebdc.S.(*_abf.PdfObjectName)
	if !_gfc {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0053\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _ebdc.S)
		return nil, _e.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0053\u0020\u006f\u0062\u006a\u0065c\u0074 \u0074y\u0070e\u0020\u0021\u003d\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0025\u0054\u0029", _ebdc.S)
	}
	_bbfd := PdfActionType(_dab.String())
	switch _bbfd {
	case ActionTypeGoTo:
		_fda, _bcef := _dcf.newPdfActionGotoFromDict(_cde)
		if _bcef != nil {
			return nil, _bcef
		}
		_fda.PdfAction = _ebdc
		_ebdc._gfg = _fda
		return _ebdc, nil
	case ActionTypeGoToR:
		_abg, _gcd := _dcf.newPdfActionGotoRFromDict(_cde)
		if _gcd != nil {
			return nil, _gcd
		}
		_abg.PdfAction = _ebdc
		_ebdc._gfg = _abg
		return _ebdc, nil
	case ActionTypeGoToE:
		_bgc, _dce := _dcf.newPdfActionGotoEFromDict(_cde)
		if _dce != nil {
			return nil, _dce
		}
		_bgc.PdfAction = _ebdc
		_ebdc._gfg = _bgc
		return _ebdc, nil
	case ActionTypeLaunch:
		_agc, _gbfe := _dcf.newPdfActionLaunchFromDict(_cde)
		if _gbfe != nil {
			return nil, _gbfe
		}
		_agc.PdfAction = _ebdc
		_ebdc._gfg = _agc
		return _ebdc, nil
	case ActionTypeThread:
		_bef, _bad := _dcf.newPdfActionThreadFromDict(_cde)
		if _bad != nil {
			return nil, _bad
		}
		_bef.PdfAction = _ebdc
		_ebdc._gfg = _bef
		return _ebdc, nil
	case ActionTypeURI:
		_cc, _dfb := _dcf.newPdfActionURIFromDict(_cde)
		if _dfb != nil {
			return nil, _dfb
		}
		_cc.PdfAction = _ebdc
		_ebdc._gfg = _cc
		return _ebdc, nil
	case ActionTypeSound:
		_fca, _dfg := _dcf.newPdfActionSoundFromDict(_cde)
		if _dfg != nil {
			return nil, _dfg
		}
		_fca.PdfAction = _ebdc
		_ebdc._gfg = _fca
		return _ebdc, nil
	case ActionTypeMovie:
		_fab, _eeb := _dcf.newPdfActionMovieFromDict(_cde)
		if _eeb != nil {
			return nil, _eeb
		}
		_fab.PdfAction = _ebdc
		_ebdc._gfg = _fab
		return _ebdc, nil
	case ActionTypeHide:
		_egaf, _bea := _dcf.newPdfActionHideFromDict(_cde)
		if _bea != nil {
			return nil, _bea
		}
		_egaf.PdfAction = _ebdc
		_ebdc._gfg = _egaf
		return _ebdc, nil
	case ActionTypeNamed:
		_cdeg, _fecc := _dcf.newPdfActionNamedFromDict(_cde)
		if _fecc != nil {
			return nil, _fecc
		}
		_cdeg.PdfAction = _ebdc
		_ebdc._gfg = _cdeg
		return _ebdc, nil
	case ActionTypeSubmitForm:
		_bbc, _def := _dcf.newPdfActionSubmitFormFromDict(_cde)
		if _def != nil {
			return nil, _def
		}
		_bbc.PdfAction = _ebdc
		_ebdc._gfg = _bbc
		return _ebdc, nil
	case ActionTypeResetForm:
		_bff, _fgd := _dcf.newPdfActionResetFormFromDict(_cde)
		if _fgd != nil {
			return nil, _fgd
		}
		_bff.PdfAction = _ebdc
		_ebdc._gfg = _bff
		return _ebdc, nil
	case ActionTypeImportData:
		_afb, _eea := _dcf.newPdfActionImportDataFromDict(_cde)
		if _eea != nil {
			return nil, _eea
		}
		_afb.PdfAction = _ebdc
		_ebdc._gfg = _afb
		return _ebdc, nil
	case ActionTypeSetOCGState:
		_ddc, _gff := _dcf.newPdfActionSetOCGStateFromDict(_cde)
		if _gff != nil {
			return nil, _gff
		}
		_ddc.PdfAction = _ebdc
		_ebdc._gfg = _ddc
		return _ebdc, nil
	case ActionTypeRendition:
		_eaf, _fcab := _dcf.newPdfActionRenditionFromDict(_cde)
		if _fcab != nil {
			return nil, _fcab
		}
		_eaf.PdfAction = _ebdc
		_ebdc._gfg = _eaf
		return _ebdc, nil
	case ActionTypeTrans:
		_fcb, _gfb := _dcf.newPdfActionTransFromDict(_cde)
		if _gfb != nil {
			return nil, _gfb
		}
		_fcb.PdfAction = _ebdc
		_ebdc._gfg = _fcb
		return _ebdc, nil
	case ActionTypeGoTo3DView:
		_bbfe, _beea := _dcf.newPdfActionGoTo3DViewFromDict(_cde)
		if _beea != nil {
			return nil, _beea
		}
		_bbfe.PdfAction = _ebdc
		_ebdc._gfg = _bbfe
		return _ebdc, nil
	case ActionTypeJavaScript:
		_acdf, _fcf := _dcf.newPdfActionJavaScriptFromDict(_cde)
		if _fcf != nil {
			return nil, _fcf
		}
		_acdf.PdfAction = _ebdc
		_ebdc._gfg = _acdf
		return _ebdc, nil
	}
	_acd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006eg\u0020u\u006ek\u006eo\u0077\u006e\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073", _bbfd)
	return nil, nil
}

func _bgde(_gbfb *_abf.PdfObjectDictionary) *VRI {
	_dgdb, _ := _abf.GetString(_gbfb.Get("\u0054\u0055"))
	_fedf, _ := _abf.GetString(_gbfb.Get("\u0054\u0053"))
	return &VRI{Cert: _gggfec(_gbfb.Get("\u0043\u0065\u0072\u0074")), OCSP: _gggfec(_gbfb.Get("\u004f\u0043\u0053\u0050")), CRL: _gggfec(_gbfb.Get("\u0043\u0052\u004c")), TU: _dgdb, TS: _fedf}
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_cagee *PdfShading) ToPdfObject() _abf.PdfObject {
	_ecaed := _cagee._eabcgc
	_gbafd, _bbagf := _cagee.getShadingDict()
	if _bbagf != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _cagee.ShadingType != nil {
		_gbafd.Set("S\u0068\u0061\u0064\u0069\u006e\u0067\u0054\u0079\u0070\u0065", _cagee.ShadingType)
	}
	if _cagee.ColorSpace != nil {
		_gbafd.Set("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _cagee.ColorSpace.ToPdfObject())
	}
	if _cagee.Background != nil {
		_gbafd.Set("\u0042\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064", _cagee.Background)
	}
	if _cagee.BBox != nil {
		_gbafd.Set("\u0042\u0042\u006f\u0078", _cagee.BBox.ToPdfObject())
	}
	if _cagee.AntiAlias != nil {
		_gbafd.Set("\u0041n\u0074\u0069\u0041\u006c\u0069\u0061s", _cagee.AntiAlias)
	}
	return _ecaed
}

// Transform rectangle with the supplied matrix.
func (_bdgcg *PdfRectangle) Transform(transformMatrix _ad.Matrix) {
	_bdgcg.Llx, _bdgcg.Lly = transformMatrix.Transform(_bdgcg.Llx, _bdgcg.Lly)
	_bdgcg.Urx, _bdgcg.Ury = transformMatrix.Transform(_bdgcg.Urx, _bdgcg.Ury)
	_bdgcg.Normalize()
}

func (_afbaeg *PdfWriter) setDocInfo(_egec _abf.PdfObject) {
	if _afbaeg.hasObject(_afbaeg._ddegc) {
		delete(_afbaeg._fdgae, _afbaeg._ddegc)
		delete(_afbaeg._dbdcg, _afbaeg._ddegc)
		for _bgbda, _faadg := range _afbaeg._edcgc {
			if _faadg == _afbaeg._ddegc {
				copy(_afbaeg._edcgc[_bgbda:], _afbaeg._edcgc[_bgbda+1:])
				_afbaeg._edcgc[len(_afbaeg._edcgc)-1] = nil
				_afbaeg._edcgc = _afbaeg._edcgc[:len(_afbaeg._edcgc)-1]
				break
			}
		}
	}
	_fcacc := _abf.PdfIndirectObject{}
	_fcacc.PdfObject = _egec
	_afbaeg._ddegc = &_fcacc
	_afbaeg.addObject(&_fcacc)
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_egbd *PdfShadingType5) ToPdfObject() _abf.PdfObject {
	_egbd.PdfShading.ToPdfObject()
	_ddfbg, _ecadg := _egbd.getShadingDict()
	if _ecadg != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _egbd.BitsPerCoordinate != nil {
		_ddfbg.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _egbd.BitsPerCoordinate)
	}
	if _egbd.BitsPerComponent != nil {
		_ddfbg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _egbd.BitsPerComponent)
	}
	if _egbd.VerticesPerRow != nil {
		_ddfbg.Set("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073\u0050e\u0072\u0052\u006f\u0077", _egbd.VerticesPerRow)
	}
	if _egbd.Decode != nil {
		_ddfbg.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _egbd.Decode)
	}
	if _egbd.Function != nil {
		if len(_egbd.Function) == 1 {
			_ddfbg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _egbd.Function[0].ToPdfObject())
		} else {
			_dagbb := _abf.MakeArray()
			for _, _gebga := range _egbd.Function {
				_dagbb.Append(_gebga.ToPdfObject())
			}
			_ddfbg.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _dagbb)
		}
	}
	return _egbd._eabcgc
}

// GetNumComponents returns the number of color components (3 for Lab).
func (_bcfg *PdfColorLab) GetNumComponents() int { return 3 }

// GetXObjectByName gets XObject by name.
func (_edbg *PdfPage) GetXObjectByName(name _abf.PdfObjectName) (_abf.PdfObject, bool) {
	_bgae, _fabfb := _edbg.Resources.XObject.(*_abf.PdfObjectDictionary)
	if !_fabfb {
		return nil, false
	}
	if _dcad := _bgae.Get(name); _dcad != nil {
		return _dcad, true
	}
	return nil, false
}

// ToPdfObject implements interface PdfModel.
func (_bcdg *PdfAnnotationPrinterMark) ToPdfObject() _abf.PdfObject {
	_bcdg.PdfAnnotation.ToPdfObject()
	_agbe := _bcdg._dbc
	_efgcf := _agbe.PdfObject.(*_abf.PdfObjectDictionary)
	_efgcf.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("P\u0072\u0069\u006e\u0074\u0065\u0072\u004d\u0061\u0072\u006b"))
	_efgcf.SetIfNotNil("\u004d\u004e", _bcdg.MN)
	return _agbe
}

// ColorFromPdfObjects returns a new PdfColor based on the input slice of color
// components. The slice should contain a single PdfObjectFloat element in
// range 0-1.
func (_gbbb *PdfColorspaceDeviceGray) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_adcf, _bcfb := _abf.GetNumbersAsFloat(objects)
	if _bcfb != nil {
		return nil, _bcfb
	}
	return _gbbb.ColorFromFloats(_adcf)
}

// PdfSignature represents a PDF signature dictionary and is used for signing via form signature fields.
// (Section 12.8, Table 252 - Entries in a signature dictionary p. 475 in PDF32000_2008).
type PdfSignature struct {
	Handler SignatureHandler
	_geebd  *_abf.PdfIndirectObject

	// Type: Sig/DocTimeStamp
	Type         *_abf.PdfObjectName
	Filter       *_abf.PdfObjectName
	SubFilter    *_abf.PdfObjectName
	Contents     *_abf.PdfObjectString
	Cert         _abf.PdfObject
	ByteRange    *_abf.PdfObjectArray
	Reference    *_abf.PdfObjectArray
	Changes      *_abf.PdfObjectArray
	Name         *_abf.PdfObjectString
	M            *_abf.PdfObjectString
	Location     *_abf.PdfObjectString
	Reason       *_abf.PdfObjectString
	ContactInfo  *_abf.PdfObjectString
	R            *_abf.PdfObjectInteger
	V            *_abf.PdfObjectInteger
	PropBuild    *_abf.PdfObjectDictionary
	PropAuthTime *_abf.PdfObjectInteger
	PropAuthType *_abf.PdfObjectName
}

func (_fdbe *PdfReader) newPdfAnnotationStampFromDict(_cdbf *_abf.PdfObjectDictionary) (*PdfAnnotationStamp, error) {
	_dbfe := PdfAnnotationStamp{}
	_efdb, _bca := _fdbe.newPdfAnnotationMarkupFromDict(_cdbf)
	if _bca != nil {
		return nil, _bca
	}
	_dbfe.PdfAnnotationMarkup = _efdb
	_dbfe.Name = _cdbf.Get("\u004e\u0061\u006d\u0065")
	return &_dbfe, nil
}

// NewStandard14Font returns the standard 14 font named `basefont` as a *PdfFont, or an error if it
// `basefont` is not one of the standard 14 font names.
func NewStandard14Font(basefont StdFontName) (*PdfFont, error) {
	_bdfc, _eegef := _bfabe(basefont)
	if _eegef != nil {
		return nil, _eegef
	}
	if basefont != SymbolName && basefont != ZapfDingbatsName {
		_bdfc._ebada = _cbb.NewWinAnsiEncoder()
	}
	return &PdfFont{_gedca: &_bdfc}, nil
}

func (_ebebd *PdfWriter) hasObject(_aaeee _abf.PdfObject) bool {
	_, _eegb := _ebebd._fdgae[_aaeee]
	return _eegb
}

// PdfShadingType1 is a Function-based shading.
type PdfShadingType1 struct {
	*PdfShading
	Domain   *_abf.PdfObjectArray
	Matrix   *_abf.PdfObjectArray
	Function []PdfFunction
}

func (_gcfed *LTV) enable(_faead, _debfc []*_fa.Certificate, _abgcb string) error {
	_gbda, _ddadd, _adgcg := _gcfed.buildCertChain(_faead, _debfc)
	if _adgcg != nil {
		return _adgcg
	}
	_bcgca, _adgcg := _gcfed.getCerts(_gbda)
	if _adgcg != nil {
		return _adgcg
	}
	var _geca, _bddac [][]byte
	if _gcfed.OCSPClient != nil {
		_geca, _adgcg = _gcfed.getOCSPs(_gbda, _ddadd)
		if _adgcg != nil {
			return _adgcg
		}
	}
	if _gcfed.CRLClient != nil {
		_bddac, _adgcg = _gcfed.getCRLs(_gbda)
		if _adgcg != nil {
			return _adgcg
		}
	}
	_agag := _gcfed._dgfe
	_bdabg, _adgcg := _agag.AddCerts(_bcgca)
	if _adgcg != nil {
		return _adgcg
	}
	_dbbec, _adgcg := _agag.AddOCSPs(_geca)
	if _adgcg != nil {
		return _adgcg
	}
	_abgbf, _adgcg := _agag.AddCRLs(_bddac)
	if _adgcg != nil {
		return _adgcg
	}
	if _abgcb != "" {
		_agag.VRI[_abgcb] = &VRI{Cert: _bdabg, OCSP: _dbbec, CRL: _abgbf}
	}
	_gcfed._bfed.SetDSS(_agag)
	return nil
}

func (_cgfc fontCommon) fontFlags() int {
	if _cgfc._dcbaf == nil {
		return 0
	}
	return _cgfc._dcbaf._bgbdf
}

// ToGoImage converts the unidoc Image to a golang Image structure.
func (_fagbe *Image) ToGoImage() (_aa.Image, error) {
	_acd.Log.Trace("\u0043\u006f\u006e\u0076er\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u006f\u0020\u0069\u006d\u0061g\u0065")
	_cbfg, _dfcga := _gca.NewImage(int(_fagbe.Width), int(_fagbe.Height), int(_fagbe.BitsPerComponent), _fagbe.ColorComponents, _fagbe.Data, _fagbe._gedg, _fagbe._ceeag)
	if _dfcga != nil {
		return nil, _dfcga
	}
	return _cbfg, nil
}

// ToPdfObject implements interface PdfModel.
func (_gdd *PdfAnnotationHighlight) ToPdfObject() _abf.PdfObject {
	_gdd.PdfAnnotation.ToPdfObject()
	_bfa := _gdd._dbc
	_faa := _bfa.PdfObject.(*_abf.PdfObjectDictionary)
	_gdd.PdfAnnotationMarkup.appendToPdfDictionary(_faa)
	_faa.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0048i\u0067\u0068\u006c\u0069\u0067\u0068t"))
	_faa.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _gdd.QuadPoints)
	return _bfa
}

// PdfAnnotationMovie represents Movie annotations.
// (Section 12.5.6.17).
type PdfAnnotationMovie struct {
	*PdfAnnotation
	T     _abf.PdfObject
	Movie _abf.PdfObject
	A     _abf.PdfObject
}

// PageProcessCallback callback function used in page loading
// that could be used to modify the page content.
//
// If an error is returned, the `ToWriter` process would fail.
//
// This callback, if defined, will take precedence over `PageCallback` callback.
type PageProcessCallback func(_dcagf int, _affce *PdfPage) error

// NewPdfFontFromTTF loads a TTF font and returns a PdfFont type that can be
// used in text styling functions.
// Uses a WinAnsiTextEncoder and loads only character codes 32-255.
// NOTE: For composite fonts such as used in symbolic languages, use NewCompositePdfFontFromTTF.
func NewPdfFontFromTTF(r _gc.ReadSeeker) (*PdfFont, error) {
	const _fgdfe = _cbb.CharCode(32)
	const _cgdbg = _cbb.CharCode(255)
	_gage, _bcbbg := _gc.ReadAll(r)
	if _bcbbg != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _bcbbg)
		return nil, _bcbbg
	}
	_aaefa, _bcbbg := _gbe.TtfParse(_dd.NewReader(_gage))
	if _bcbbg != nil {
		_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0054\u0054F\u0020\u0066\u006fn\u0074:\u0020\u0025\u0076", _bcbbg)
		return nil, _bcbbg
	}
	_fafcd := &pdfFontSimple{_aadgb: make(map[_cbb.CharCode]float64), fontCommon: fontCommon{_aacbc: "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065"}}
	_fafcd._ebada = _cbb.NewWinAnsiEncoder()
	_fafcd._ecggf = _aaefa.PostScriptName
	_fafcd.FirstChar = _abf.MakeInteger(int64(_fgdfe))
	_fafcd.LastChar = _abf.MakeInteger(int64(_cgdbg))
	_fgea := 1000.0 / float64(_aaefa.UnitsPerEm)
	if len(_aaefa.Widths) <= 0 {
		return nil, _fd.New("\u0045\u0052\u0052O\u0052\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u0028\u0057\u0069\u0064\u0074\u0068\u0073\u0029")
	}
	_fgged := _fgea * float64(_aaefa.Widths[0])
	_ddcde := make([]float64, 0, _cgdbg-_fgdfe+1)
	for _gebgf := _fgdfe; _gebgf <= _cgdbg; _gebgf++ {
		_cfeca, _gefdc := _fafcd.Encoder().CharcodeToRune(_gebgf)
		if !_gefdc {
			_acd.Log.Debug("\u0052u\u006e\u0065\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u0028c\u006f\u0064\u0065\u003a\u0020\u0025\u0064)", _gebgf)
			_ddcde = append(_ddcde, _fgged)
			continue
		}
		_bcedf, _cacc := _aaefa.Chars[_cfeca]
		if !_cacc {
			_acd.Log.Debug("R\u0075\u006e\u0065\u0020no\u0074 \u0069\u006e\u0020\u0054\u0054F\u0020\u0043\u0068\u0061\u0072\u0073")
			_ddcde = append(_ddcde, _fgged)
			continue
		}
		_dafa := _fgea * float64(_aaefa.Widths[_bcedf])
		_ddcde = append(_ddcde, _dafa)
	}
	_fafcd.Widths = _abf.MakeIndirectObject(_abf.MakeArrayFromFloats(_ddcde))
	if len(_ddcde) < int(_cgdbg-_fgdfe+1) {
		_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u006f\u0066\u0020\u0077\u0069\u0064\u0074\u0068s,\u0020\u0025\u0064 \u003c \u0025\u0064", len(_ddcde), 255-32+1)
		return nil, _abf.ErrRangeError
	}
	for _edbbb := _fgdfe; _edbbb <= _cgdbg; _edbbb++ {
		_fafcd._aadgb[_edbbb] = _ddcde[_edbbb-_fgdfe]
	}
	_fafcd.Encoding = _abf.MakeName("\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	_bacfd := &PdfFontDescriptor{}
	_bacfd.FontName = _abf.MakeName(_aaefa.PostScriptName)
	_bacfd.Ascent = _abf.MakeFloat(_fgea * float64(_aaefa.TypoAscender))
	_bacfd.Descent = _abf.MakeFloat(_fgea * float64(_aaefa.TypoDescender))
	_bacfd.CapHeight = _abf.MakeFloat(_fgea * float64(_aaefa.CapHeight))
	_bacfd.FontBBox = _abf.MakeArrayFromFloats([]float64{_fgea * float64(_aaefa.Xmin), _fgea * float64(_aaefa.Ymin), _fgea * float64(_aaefa.Xmax), _fgea * float64(_aaefa.Ymax)})
	_bacfd.ItalicAngle = _abf.MakeFloat(_aaefa.ItalicAngle)
	_bacfd.MissingWidth = _abf.MakeFloat(_fgea * float64(_aaefa.Widths[0]))
	_fdfbd, _bcbbg := _abf.MakeStream(_gage, _abf.NewFlateEncoder())
	if _bcbbg != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020m\u0061\u006b\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _bcbbg)
		return nil, _bcbbg
	}
	_fdfbd.PdfObjectDictionary.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _abf.MakeInteger(int64(len(_gage))))
	_bacfd.FontFile2 = _fdfbd
	if _aaefa.Bold {
		_bacfd.StemV = _abf.MakeInteger(120)
	} else {
		_bacfd.StemV = _abf.MakeInteger(70)
	}
	_abbb := _bbadf
	if _aaefa.IsFixedPitch {
		_abbb |= _becb
	}
	if _aaefa.ItalicAngle != 0 {
		_abbb |= _bacb
	}
	_bacfd.Flags = _abf.MakeInteger(int64(_abbb))
	_fafcd._dcbaf = _bacfd
	_dfdd := &PdfFont{_gedca: _fafcd}
	return _dfdd, nil
}

// PdfPage represents a page in a PDF document. (7.7.3.3 - Table 30).
type PdfPage struct {
	Parent               _abf.PdfObject
	LastModified         *PdfDate
	Resources            *PdfPageResources
	CropBox              *PdfRectangle
	MediaBox             *PdfRectangle
	BleedBox             *PdfRectangle
	TrimBox              *PdfRectangle
	ArtBox               *PdfRectangle
	BoxColorInfo         _abf.PdfObject
	Contents             _abf.PdfObject
	Rotate               *int64
	Group                _abf.PdfObject
	Thumb                _abf.PdfObject
	B                    _abf.PdfObject
	Dur                  _abf.PdfObject
	Trans                _abf.PdfObject
	AA                   _abf.PdfObject
	Metadata             _abf.PdfObject
	PieceInfo            _abf.PdfObject
	StructParents        _abf.PdfObject
	ID                   _abf.PdfObject
	PZ                   _abf.PdfObject
	SeparationInfo       _abf.PdfObject
	Tabs                 _abf.PdfObject
	TemplateInstantiated _abf.PdfObject
	PresSteps            _abf.PdfObject
	UserUnit             _abf.PdfObject
	VP                   _abf.PdfObject
	Annots               _abf.PdfObject
	_baagf               []*PdfAnnotation
	_bdbfa               *_abf.PdfObjectDictionary
	_gefee               *_abf.PdfIndirectObject
	_efca                _abf.PdfObjectDictionary
	_dbaef               *PdfReader
}

// NewPdfActionSetOCGState returns a new "named" action.
func NewPdfActionSetOCGState() *PdfActionSetOCGState {
	_aeee := NewPdfAction()
	_ff := &PdfActionSetOCGState{}
	_ff.PdfAction = _aeee
	_aeee.SetContext(_ff)
	return _ff
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain a single element.
func (_deceg *PdfColorspaceSpecialSeparation) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 1 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_fegd := vals[0]
	_afbe := []float64{_fegd}
	_agad, _bfgf := _deceg.TintTransform.Evaluate(_afbe)
	if _bfgf != nil {
		_acd.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u0076\u0061\u006c\u0075\u0061\u0074\u0065: \u0025\u0076", _bfgf)
		_acd.Log.Trace("\u0054\u0069\u006e\u0074 t\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u002b\u0076", _deceg.TintTransform)
		return nil, _bfgf
	}
	_acd.Log.Trace("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0043\u006f\u006c\u006fr\u0046\u0072\u006f\u006d\u0046\u006c\u006f\u0061\u0074\u0073\u0028\u0025\u002bv\u0029\u0020\u006f\u006e\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061te\u0053\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0023\u0076", _agad, _deceg.AlternateSpace)
	_cbbf, _bfgf := _deceg.AlternateSpace.ColorFromFloats(_agad)
	if _bfgf != nil {
		_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u002c\u0020\u0066a\u0069\u006c\u0065d \u0074\u006f\u0020\u0065\u0076\u0061l\u0075\u0061\u0074\u0065\u0020\u0069\u006e\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u003a \u0025\u0076", _bfgf)
		return nil, _bfgf
	}
	return _cbbf, nil
}

func (_fdbb *PdfReader) lookupPageByObject(_dcbgac _abf.PdfObject) (*PdfPage, error) {
	return nil, _fd.New("\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// PdfColorspaceCalRGB stores A, B, C components
type PdfColorspaceCalRGB struct {
	WhitePoint []float64
	BlackPoint []float64
	Gamma      []float64
	Matrix     []float64
	_afd       *_abf.PdfObjectDictionary
	_bdfg      *_abf.PdfIndirectObject
}

func (_ecfb Image) getBase() _gca.ImageBase {
	return _gca.NewImageBase(int(_ecfb.Width), int(_ecfb.Height), int(_ecfb.BitsPerComponent), _ecfb.ColorComponents, _ecfb.Data, _ecfb._gedg, _ecfb._ceeag)
}

// PdfAnnotationScreen represents Screen annotations.
// (Section 12.5.6.18).
type PdfAnnotationScreen struct {
	*PdfAnnotation
	T  _abf.PdfObject
	MK _abf.PdfObject
	A  _abf.PdfObject
	AA _abf.PdfObject
}

// GetCharMetrics returns the character metrics for the specified character code.  A bool flag is
// returned to indicate whether or not the entry was found in the glyph to charcode mapping.
// How it works:
//  1. Return a value the /Widths array (charWidths) if there is one.
//  2. If the font has the same name as a standard 14 font then return width=250.
//  3. Otherwise return no match and let the caller substitute a default.
func (_dace pdfFontSimple) GetCharMetrics(code _cbb.CharCode) (_gbe.CharMetrics, bool) {
	if _bcgggd, _ecga := _dace._aadgb[code]; _ecga {
		return _gbe.CharMetrics{Wx: _bcgggd}, true
	}
	if _gbe.IsStdFont(_gbe.StdFontName(_dace._ecggf)) {
		return _gbe.CharMetrics{Wx: 250}, true
	}
	return _gbe.CharMetrics{}, false
}

// NewPdfReader returns a new PdfReader for an input io.ReadSeeker interface. Can be used to read PDF from
// memory or file. Immediately loads and traverses the PDF structure including pages and page contents (if
// not encrypted). Loads entire document structure into memory.
// Alternatively a lazy-loading reader can be created with NewPdfReaderLazy which loads only references,
// and references are loaded from disk into memory on an as-needed basis.
func NewPdfReader(rs _gc.ReadSeeker) (*PdfReader, error) {
	const _cbedce = "\u006do\u0064e\u006c\u003a\u004e\u0065\u0077P\u0064\u0066R\u0065\u0061\u0064\u0065\u0072"
	return _fbaec(rs, &ReaderOpts{}, false, _cbedce)
}

// ToPdfObject converts the PdfFont object to its PDF representation.
func (_cbbg *PdfFont) ToPdfObject() _abf.PdfObject {
	if _cbbg._gedca == nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0066\u006f\u006e\u0074 \u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return _abf.MakeNull()
	}
	return _cbbg._gedca.ToPdfObject()
}

// PdfColor interface represents a generic color in PDF.
type PdfColor interface{}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_gfabc pdfFontType3) GetRuneMetrics(r rune) (_gbe.CharMetrics, bool) {
	_ggcee := _gfabc.Encoder()
	if _ggcee == nil {
		_acd.Log.Debug("\u004e\u006f\u0020en\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u0073\u003d\u0025\u0073", _gfabc)
		return _gbe.CharMetrics{}, false
	}
	_ggafa, _gaged := _ggcee.RuneToCharcode(r)
	if !_gaged {
		if r != ' ' {
			_acd.Log.Trace("\u004e\u006f\u0020c\u0068\u0061\u0072\u0063o\u0064\u0065\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u003d\u0025\u0076\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", r, _gfabc)
		}
		return _gbe.CharMetrics{}, false
	}
	_bfabb, _bbeff := _gfabc.GetCharMetrics(_ggafa)
	return _bfabb, _bbeff
}

// PdfAnnotationRedact represents Redact annotations.
// (Section 12.5.6.23).
type PdfAnnotationRedact struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints  _abf.PdfObject
	IC          _abf.PdfObject
	RO          _abf.PdfObject
	OverlayText _abf.PdfObject
	Repeat      _abf.PdfObject
	DA          _abf.PdfObject
	Q           _abf.PdfObject
}

// GetType returns the button field type which returns one of the following
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
func (_cbeg *PdfFieldButton) GetType() ButtonType {
	_gbdec := ButtonTypeCheckbox
	if _cbeg.Ff != nil {
		if (uint32(*_cbeg.Ff) & FieldFlagPushbutton.Mask()) > 0 {
			_gbdec = ButtonTypePush
		} else if (uint32(*_cbeg.Ff) & FieldFlagRadio.Mask()) > 0 {
			_gbdec = ButtonTypeRadio
		}
	}
	return _gbdec
}

func (_ebbcb *PdfReader) traverseObjectData(_ebbca _abf.PdfObject) error {
	return _abf.ResolveReferencesDeep(_ebbca, _ebbcb._ggbccc)
}

func (_acefg *PdfReader) loadOutlines() (*PdfOutlineTreeNode, error) {
	if _acefg._bebc.GetCrypter() != nil && !_acefg._bebc.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_ebaac := _acefg._dagde
	_abgcd := _ebaac.Get("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	if _abgcd == nil {
		return nil, nil
	}
	_acd.Log.Trace("\u002d\u0048\u0061\u0073\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
	_daega := _abf.ResolveReference(_abgcd)
	_acd.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0020\u0072\u006f\u006f\u0074\u003a\u0020\u0025\u0076", _daega)
	if _bbadc := _abf.IsNullObject(_daega); _bbadc {
		_acd.Log.Trace("\u004f\u0075\u0074li\u006e\u0065\u0020\u0072\u006f\u006f\u0074\u0020\u0069s\u0020n\u0075l\u006c \u002d\u0020\u006e\u006f\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u0073")
		return nil, nil
	}
	_fbbag, _bfcac := _daega.(*_abf.PdfIndirectObject)
	if !_bfcac {
		if _, _bfcb := _abf.GetDict(_daega); !_bfcb {
			_acd.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020o\u0075\u0074\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			return nil, nil
		}
		_acd.Log.Debug("\u004f\u0075t\u006c\u0069\u006e\u0065\u0020r\u006f\u006f\u0074\u0020\u0069s\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u002e\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		_fbbag = _abf.MakeIndirectObject(_daega)
	}
	_fdeb, _bfcac := _fbbag.PdfObject.(*_abf.PdfObjectDictionary)
	if !_bfcac {
		return nil, _fd.New("\u006f\u0075\u0074\u006c\u0069n\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y")
	}
	_acd.Log.Trace("O\u0075\u0074\u006c\u0069ne\u0020r\u006f\u006f\u0074\u0020\u0064i\u0063\u0074\u003a\u0020\u0025\u0076", _fdeb)
	_gadbd, _, _fbeea := _acefg.buildOutlineTree(_fbbag, nil, nil, nil)
	if _fbeea != nil {
		return nil, _fbeea
	}
	_acd.Log.Trace("\u0052\u0065\u0073\u0075\u006c\u0074\u0069\u006e\u0067\u0020\u006fu\u0074\u006c\u0069\u006e\u0065\u0020\u0074\u0072\u0065\u0065:\u0020\u0025\u0076", _gadbd)
	return _gadbd, nil
}

// PdfAction represents an action in PDF (section 12.6 p. 412).
type PdfAction struct {
	_gfg PdfModel
	Type _abf.PdfObject
	S    _abf.PdfObject
	Next _abf.PdfObject
	_egg *_abf.PdfIndirectObject
}

// Add appends an outline item as a child of the current outline item.
func (_eacbc *OutlineItem) Add(item *OutlineItem) { _eacbc.Entries = append(_eacbc.Entries, item) }

func (_ada *PdfReader) newPdfAnnotationMarkupFromDict(_cdfg *_abf.PdfObjectDictionary) (*PdfAnnotationMarkup, error) {
	_egee := &PdfAnnotationMarkup{}
	if _cdfb := _cdfg.Get("\u0054"); _cdfb != nil {
		_egee.T = _cdfb
	}
	if _aabc := _cdfg.Get("\u0050\u006f\u0070u\u0070"); _aabc != nil {
		_fgad, _dcce := _aabc.(*_abf.PdfIndirectObject)
		if !_dcce {
			if _, _gfe := _aabc.(*_abf.PdfObjectNull); !_gfe {
				return nil, _fd.New("p\u006f\u0070\u0075\u0070\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0070\u006f\u0069\u006e\u0074\u0020t\u006f\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			}
		} else {
			_eegd, _ddff := _ada.newPdfAnnotationFromIndirectObject(_fgad)
			if _ddff != nil {
				return nil, _ddff
			}
			if _eegd != nil {
				_ggbe, _gbc := _eegd._edg.(*PdfAnnotationPopup)
				if !_gbc {
					return nil, _fd.New("\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0020\u0070\u006f\u0070\u0075\u0070\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e")
				}
				_egee.Popup = _ggbe
			}
		}
	}
	if _gaba := _cdfg.Get("\u0043\u0041"); _gaba != nil {
		_egee.CA = _gaba
	}
	if _agbb := _cdfg.Get("\u0052\u0043"); _agbb != nil {
		_egee.RC = _agbb
	}
	if _abaa := _cdfg.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"); _abaa != nil {
		_egee.CreationDate = _abaa
	}
	if _eecg := _cdfg.Get("\u0049\u0052\u0054"); _eecg != nil {
		_egee.IRT = _eecg
	}
	if _cfc := _cdfg.Get("\u0053\u0075\u0062\u006a"); _cfc != nil {
		_egee.Subj = _cfc
	}
	if _edcb := _cdfg.Get("\u0052\u0054"); _edcb != nil {
		_egee.RT = _edcb
	}
	if _ffe := _cdfg.Get("\u0049\u0054"); _ffe != nil {
		_egee.IT = _ffe
	}
	if _gbea := _cdfg.Get("\u0045\u0078\u0044\u0061\u0074\u0061"); _gbea != nil {
		_egee.ExData = _gbea
	}
	return _egee, nil
}

func (_fdec *pdfFontSimple) getFontDescriptor() *PdfFontDescriptor {
	if _fgfa := _fdec._dcbaf; _fgfa != nil {
		return _fgfa
	}
	return _fdec._abeb
}

// RepairAcroForm attempts to rebuild the AcroForm fields using the widget
// annotations present in the document pages. Pass nil for the opts parameter
// in order to use the default options.
// NOTE: Currently, the opts parameter is declared in order to enable adding
// future options, but passing nil will always result in the default options
// being used.
func (_fgcg *PdfReader) RepairAcroForm(opts *AcroFormRepairOptions) error {
	var _gfaa []*PdfField
	_egbb := map[*_abf.PdfIndirectObject]struct{}{}
	for _, _aedcb := range _fgcg.PageList {
		_edacf, _fceee := _aedcb.GetAnnotations()
		if _fceee != nil {
			return _fceee
		}
		for _, _dbegf := range _edacf {
			var _ggbfc *PdfField
			switch _edaag := _dbegf.GetContext().(type) {
			case *PdfAnnotationWidget:
				if _edaag._agdc != nil {
					_ggbfc = _edaag._agdc
					break
				}
				if _fgbfe, _bbcfa := _abf.GetIndirect(_edaag.Parent); _bbcfa {
					_ggbfc, _fceee = _fgcg.newPdfFieldFromIndirectObject(_fgbfe, nil)
					if _fceee == nil {
						break
					}
					_acd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _fgbfe, _fceee)
				}
				if _edaag._dbc != nil {
					_ggbfc, _fceee = _fgcg.newPdfFieldFromIndirectObject(_edaag._dbc, nil)
					if _fceee == nil {
						break
					}
					_acd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072s\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0066\u0069\u0065ld\u0020\u0025\u002bv\u003a \u0025\u0076", _edaag._dbc, _fceee)
				}
			}
			if _ggbfc == nil {
				continue
			}
			if _, _ecceb := _egbb[_ggbfc._dgdc]; _ecceb {
				continue
			}
			_egbb[_ggbfc._dgdc] = struct{}{}
			_gfaa = append(_gfaa, _ggbfc)
		}
	}
	if len(_gfaa) == 0 {
		return nil
	}
	if _fgcg.AcroForm == nil {
		_fgcg.AcroForm = NewPdfAcroForm()
	}
	_fgcg.AcroForm.Fields = &_gfaa
	return nil
}

// B returns the value of the blue component of the color.
func (_gfgd *PdfColorDeviceRGB) B() float64 { return _gfgd[2] }

// ToPdfObject returns the PDF representation of the shading pattern.
func (_feded *PdfShadingPatternType2) ToPdfObject() _abf.PdfObject {
	_feded.PdfPattern.ToPdfObject()
	_cggd := _feded.getDict()
	if _feded.Shading != nil {
		_cggd.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _feded.Shading.ToPdfObject())
	}
	if _feded.Matrix != nil {
		_cggd.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _feded.Matrix)
	}
	if _feded.ExtGState != nil {
		_cggd.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _feded.ExtGState)
	}
	return _feded._bcfca
}

// FieldFilterFunc represents a PDF field filtering function. If the function
// returns true, the PDF field is kept, otherwise it is discarded.
type FieldFilterFunc func(*PdfField) bool

// ToPdfObject implements interface PdfModel.
func (_acdd *PdfActionMovie) ToPdfObject() _abf.PdfObject {
	_acdd.PdfAction.ToPdfObject()
	_edb := _acdd._egg
	_edd := _edb.PdfObject.(*_abf.PdfObjectDictionary)
	_edd.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeMovie)))
	_edd.SetIfNotNil("\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e", _acdd.Annotation)
	_edd.SetIfNotNil("\u0054", _acdd.T)
	_edd.SetIfNotNil("\u004fp\u0065\u0072\u0061\u0074\u0069\u006fn", _acdd.Operation)
	return _edb
}

// GetSamples converts the raw byte slice into samples which are stored in a uint32 bit array.
// Each sample is represented by BitsPerComponent consecutive bits in the raw data.
// NOTE: The method resamples the image byte data before returning the result and
// this could lead to high memory usage, especially on large images. It should
// be avoided, when possible. It is recommended to access the Data field of the
// image directly or use the ColorAt method to extract individual pixels.
func (_dedfd *Image) GetSamples() []uint32 {
	_gfcfac := _gf.ResampleBytes(_dedfd.Data, int(_dedfd.BitsPerComponent))
	if _dedfd.BitsPerComponent < 8 {
		_gfcfac = _dedfd.samplesTrimPadding(_gfcfac)
	}
	_gfece := int(_dedfd.Width) * int(_dedfd.Height) * _dedfd.ColorComponents
	if len(_gfcfac) < _gfece {
		_acd.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u0066\u0065\u0077\u0020\u0073a\u006d\u0070\u006c\u0065\u0073\u0020(\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064\u0029", len(_gfcfac), _gfece)
		return _gfcfac
	} else if len(_gfcfac) > _gfece {
		_acd.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0054o\u006f\u0020\u006d\u0061\u006e\u0079\u0020s\u0061\u006d\u0070\u006c\u0065\u0073 \u0028\u0067\u006f\u0074\u0020\u0025\u0064\u002c\u0020\u0065\u0078p\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u0025\u0064", len(_gfcfac), _gfece)
		_gfcfac = _gfcfac[:_gfece]
	}
	return _gfcfac
}

// Height returns the height of `rect`.
func (_dacgg *PdfRectangle) Height() float64 { return _ge.Abs(_dacgg.Ury - _dacgg.Lly) }

// SetPageLabels sets the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_bcbcb *PdfWriter) SetPageLabels(pageLabels _abf.PdfObject) error {
	if pageLabels == nil {
		return nil
	}
	_acd.Log.Trace("\u0053\u0065t\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006cs.\u002e\u002e")
	_bcbcb._ddffc.Set("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073", pageLabels)
	return _bcbcb.addObjects(pageLabels)
}

// AddOutlineTree adds outlines to a PDF file.
func (_fbcgd *PdfWriter) AddOutlineTree(outlineTree *PdfOutlineTreeNode) { _fbcgd._gbcge = outlineTree }

func (_adca *PdfAcroForm) fillImageWithAppearance(_cadfe FieldImageProvider, _bfdd FieldAppearanceGenerator) error {
	if _adca == nil {
		return nil
	}
	_ggegf, _ggaga := _cadfe.FieldImageValues()
	if _ggaga != nil {
		return _ggaga
	}
	for _, _ggcfa := range _adca.AllFields() {
		_eabad := _ggcfa.PartialName()
		_feefg, _gageg := _ggegf[_eabad]
		if !_gageg {
			if _acgbf, _eggff := _ggcfa.FullName(); _eggff == nil {
				_feefg, _gageg = _ggegf[_acgbf]
			}
		}
		if !_gageg {
			_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020f\u006f\u0072\u006d \u0066\u0069\u0065l\u0064\u0020\u0025\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069n \u0074\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _eabad)
			continue
		}
		switch _fcfc := _ggcfa.GetContext().(type) {
		case *PdfFieldButton:
			if _fcfc.IsPush() {
				_fcfc.SetFillImage(_feefg)
			}
		}
		if _bfdd == nil {
			continue
		}
		for _, _cbda := range _ggcfa.Annotations {
			_cdgfa, _ddfdg := _bfdd.GenerateAppearanceDict(_adca, _ggcfa, _cbda)
			if _ddfdg != nil {
				return _ddfdg
			}
			_cbda.AP = _cdgfa
			_cbda.ToPdfObject()
		}
	}
	return nil
}

func _dfaag(_bcced *_abf.PdfObjectDictionary) (*PdfTilingPattern, error) {
	_daec := &PdfTilingPattern{}
	_bggbg := _bcced.Get("\u0050a\u0069\u006e\u0074\u0054\u0079\u0070e")
	if _bggbg == nil {
		_acd.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054\u0079\u0070\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_ecddc, _ffgc := _bggbg.(*_abf.PdfObjectInteger)
	if !_ffgc {
		_acd.Log.Debug("\u0050\u0061\u0069\u006e\u0074\u0054y\u0070\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006ft\u0020\u0025\u0054\u0029", _bggbg)
		return nil, _abf.ErrTypeError
	}
	_daec.PaintType = _ecddc
	_bggbg = _bcced.Get("\u0054\u0069\u006c\u0069\u006e\u0067\u0054\u0079\u0070\u0065")
	if _bggbg == nil {
		_acd.Log.Debug("\u0054i\u006ci\u006e\u0067\u0054\u0079\u0070e\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_caff, _ffgc := _bggbg.(*_abf.PdfObjectInteger)
	if !_ffgc {
		_acd.Log.Debug("\u0054\u0069\u006cin\u0067\u0054\u0079\u0070\u0065\u0020\u006e\u006f\u0074 \u0061n\u0020i\u006et\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bggbg)
		return nil, _abf.ErrTypeError
	}
	_daec.TilingType = _caff
	_bggbg = _bcced.Get("\u0042\u0042\u006f\u0078")
	if _bggbg == nil {
		_acd.Log.Debug("\u0042\u0042\u006fx\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bggbg = _abf.TraceToDirectObject(_bggbg)
	_dcgb, _ffgc := _bggbg.(*_abf.PdfObjectArray)
	if !_ffgc {
		_acd.Log.Debug("\u0042B\u006f\u0078 \u0073\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0079\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061y\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bggbg)
		return nil, _abf.ErrTypeError
	}
	_effb, _fcbcb := NewPdfRectangle(*_dcgb)
	if _fcbcb != nil {
		_acd.Log.Debug("\u0042\u0042\u006f\u0078\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _fcbcb)
		return nil, _fcbcb
	}
	_daec.BBox = _effb
	_bggbg = _bcced.Get("\u0058\u0053\u0074e\u0070")
	if _bggbg == nil {
		_acd.Log.Debug("\u0058\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_ceadg, _fcbcb := _abf.GetNumberAsFloat(_bggbg)
	if _fcbcb != nil {
		_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0058S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _ceadg)
		return nil, _fcbcb
	}
	_daec.XStep = _abf.MakeFloat(_ceadg)
	_bggbg = _bcced.Get("\u0059\u0053\u0074e\u0070")
	if _bggbg == nil {
		_acd.Log.Debug("\u0059\u0053\u0074\u0065\u0070\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_faeg, _fcbcb := _abf.GetNumberAsFloat(_bggbg)
	if _fcbcb != nil {
		_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0059S\u0074e\u0070\u0020\u0061\u0073\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _faeg)
		return nil, _fcbcb
	}
	_daec.YStep = _abf.MakeFloat(_faeg)
	_bggbg = _bcced.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s")
	if _bggbg == nil {
		_acd.Log.Debug("\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bcced, _ffgc = _abf.TraceToDirectObject(_bggbg).(*_abf.PdfObjectDictionary)
	if !_ffgc {
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _bggbg)
	}
	_begce, _fcbcb := NewPdfPageResourcesFromDict(_bcced)
	if _fcbcb != nil {
		return nil, _fcbcb
	}
	_daec.Resources = _begce
	if _ecafg := _bcced.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _ecafg != nil {
		_dcgae, _cfgfb := _ecafg.(*_abf.PdfObjectArray)
		if !_cfgfb {
			_acd.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ecafg)
			return nil, _abf.ErrTypeError
		}
		_daec.Matrix = _dcgae
	}
	return _daec, nil
}

// SetPatternByName sets a pattern resource specified by keyName.
func (_cbfcd *PdfPageResources) SetPatternByName(keyName _abf.PdfObjectName, pattern _abf.PdfObject) error {
	if _cbfcd.Pattern == nil {
		_cbfcd.Pattern = _abf.MakeDict()
	}
	_fcffb, _eaccf := _abf.GetDict(_cbfcd.Pattern)
	if !_eaccf {
		return _abf.ErrTypeError
	}
	_fcffb.Set(keyName, pattern)
	return nil
}

// GetRevision returns the specific version of the PdfReader for the current Pdf document
func (_eebaf *PdfReader) GetRevision(revisionNumber int) (*PdfReader, error) {
	_egcgg := _eebaf._bebc.GetRevisionNumber()
	if revisionNumber < 0 || revisionNumber > _egcgg {
		return nil, _fd.New("w\u0072\u006f\u006e\u0067 r\u0065v\u0069\u0073\u0069\u006f\u006e \u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if revisionNumber == _egcgg {
		return _eebaf, nil
	}
	if _eebaf._egade[revisionNumber] != nil {
		return _eebaf._egade[revisionNumber], nil
	}
	_dgacc := _eebaf
	for _aaefbd := _egcgg - 1; _aaefbd >= revisionNumber; _aaefbd-- {
		_bcdeaf, _dacc := _dgacc.GetPreviousRevision()
		if _dacc != nil {
			return nil, _dacc
		}
		_eebaf._egade[_aaefbd] = _bcdeaf
		_dgacc = _bcdeaf
	}
	return _dgacc, nil
}

func (_fdbg *PdfReader) newPdfAnnotationWidgetFromDict(_dbca *_abf.PdfObjectDictionary) (*PdfAnnotationWidget, error) {
	_bfg := PdfAnnotationWidget{}
	_bfg.H = _dbca.Get("\u0048")
	_bfg.MK = _dbca.Get("\u004d\u004b")
	_bfg.A = _dbca.Get("\u0041")
	_bfg.AA = _dbca.Get("\u0041\u0041")
	_bfg.BS = _dbca.Get("\u0042\u0053")
	_bfg.Parent = _dbca.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	return &_bfg, nil
}

type pdfCIDFontType0 struct {
	fontCommon
	_dgafee *_abf.PdfIndirectObject
	_aefc   _cbb.TextEncoder

	// Table 117 – Entries in a CIDFont dictionary (page 269)
	// (Required) Dictionary that defines the character collection of the CIDFont.
	// See Table 116.
	CIDSystemInfo *_abf.PdfObjectDictionary

	// Glyph metrics fields (optional).
	DW     _abf.PdfObject
	W      _abf.PdfObject
	DW2    _abf.PdfObject
	W2     _abf.PdfObject
	_fbcfb map[_cbb.CharCode]float64
	_bdced float64
}

// NewPdfAppenderWithOpts creates a new Pdf appender from a Pdf reader with options.
func NewPdfAppenderWithOpts(reader *PdfReader, opts *ReaderOpts, encryptOptions *EncryptOptions) (*PdfAppender, error) {
	_dggd := &PdfAppender{_eeded: reader._affbb, Reader: reader, _bdcd: reader._bebc, _gfeg: reader._ggbccc}
	_bdce, _dcgc := _dggd._eeded.Seek(0, _gc.SeekEnd)
	if _dcgc != nil {
		return nil, _dcgc
	}
	_dggd._cfga = _bdce
	if _, _dcgc = _dggd._eeded.Seek(0, _gc.SeekStart); _dcgc != nil {
		return nil, _dcgc
	}
	_dggd._agda, _dcgc = NewPdfReaderWithOpts(_dggd._eeded, opts)
	if _dcgc != nil {
		return nil, _dcgc
	}
	for _, _cgad := range _dggd.Reader.GetObjectNums() {
		if _dggd._ffc < _cgad {
			_dggd._ffc = _cgad
		}
	}
	_dggd._abce = _dggd._bdcd.GetXrefTable()
	_dggd._dac = _dggd._bdcd.GetXrefOffset()
	_dggd._cggfa = append(_dggd._cggfa, _dggd._agda.PageList...)
	_dggd._gcba = make(map[_abf.PdfObject]struct{})
	_dggd._bge = make(map[_abf.PdfObject]int64)
	_dggd._cdbbg = make(map[_abf.PdfObject]struct{})
	_dggd._ffbb = _dggd._agda.AcroForm
	_dggd._ffbe = _dggd._agda.DSS
	if opts != nil {
		_dggd._fcfb = opts.Password
	}
	if encryptOptions != nil {
		_dggd._bbag = encryptOptions
	}
	return _dggd, nil
}

// GetXObjectImageByName returns the XObjectImage with the specified name from the
// page resources, if it exists.
func (_gdcgb *PdfPageResources) GetXObjectImageByName(keyName _abf.PdfObjectName) (*XObjectImage, error) {
	_bbce, _ccfd := _gdcgb.GetXObjectByName(keyName)
	if _bbce == nil {
		return nil, nil
	}
	if _ccfd != XObjectTypeImage {
		return nil, _fd.New("\u006e\u006f\u0074 \u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065")
	}
	_fdfd, _bbcbe := NewXObjectImageFromStream(_bbce)
	if _bbcbe != nil {
		return nil, _bbcbe
	}
	return _fdfd, nil
}

// NewPdfAnnotationCircle returns a new circle annotation.
func NewPdfAnnotationCircle() *PdfAnnotationCircle {
	_dfdc := NewPdfAnnotation()
	_egd := &PdfAnnotationCircle{}
	_egd.PdfAnnotation = _dfdc
	_egd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dfdc.SetContext(_egd)
	return _egd
}

// PdfAnnotationStamp represents Stamp annotations.
// (Section 12.5.6.12).
type PdfAnnotationStamp struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Name _abf.PdfObject
}

// AcroFormRepairOptions contains options for rebuilding the AcroForm.
type AcroFormRepairOptions struct{}

// C returns the value of the C component of the color.
func (_dbfa *PdfColorCalRGB) C() float64 { return _dbfa[2] }

func (_ggacg *PdfWriter) writeObjectsInStreams(_gdcgc map[_abf.PdfObject]bool) error {
	for _, _cbdaa := range _ggacg._edcgc {
		if _aeege := _gdcgc[_cbdaa]; _aeege {
			continue
		}
		_cacac := int64(0)
		switch _dedff := _cbdaa.(type) {
		case *_abf.PdfIndirectObject:
			_cacac = _dedff.ObjectNumber
		case *_abf.PdfObjectStream:
			_cacac = _dedff.ObjectNumber
		case *_abf.PdfObjectStreams:
			_cacac = _dedff.ObjectNumber
		default:
			_acd.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0055n\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0074\u0079\u0070\u0065 \u0069\u006e\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0054", _cbdaa)
			return ErrTypeCheck
		}
		if _ggacg._ddbgd != nil && _cbdaa != _ggacg._dcdbb {
			_cfdff := _ggacg._ddbgd.Encrypt(_cbdaa, _cacac, 0)
			if _cfdff != nil {
				_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067\u0020(%\u0073\u0029", _cfdff)
				return _cfdff
			}
		}
		_ggacg.writeObject(int(_cacac), _cbdaa)
	}
	return nil
}

// GetContainingPdfObject implements interface PdfModel.
func (_accd *PdfSignatureReference) GetContainingPdfObject() _abf.PdfObject { return _accd._bfbaf }

// DecodeArray returns the range of color component values in DeviceRGB colorspace.
func (_gfcc *PdfColorspaceDeviceRGB) DecodeArray() []float64 {
	return []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
}

// Initialize initializes the PdfSignature.
func (_ddbdc *PdfSignature) Initialize() error {
	if _ddbdc.Handler == nil {
		return _fd.New("\u0073\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0068\u0061\u006e\u0064\u006c\u0065r\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	return _ddbdc.Handler.InitSignature(_ddbdc)
}

// Evaluate runs the function on the passed in slice and returns the results.
func (_bgfdc *PdfFunctionType0) Evaluate(x []float64) ([]float64, error) {
	if len(x) != _bgfdc.NumInputs {
		_acd.Log.Error("\u004eu\u006d\u0062e\u0072\u0020\u006f\u0066 \u0069\u006e\u0070u\u0074\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061tc\u0068\u0069\u006eg\u0020\u0077h\u0061\u0074\u0020\u0069\u0073\u0020n\u0065\u0065d\u0065\u0064")
		return nil, _fd.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _bgfdc._bega == nil {
		_ebgge := _bgfdc.processSamples()
		if _ebgge != nil {
			return nil, _ebgge
		}
	}
	_cbfcc := _bgfdc.Encode
	if _cbfcc == nil {
		_cbfcc = []float64{}
		for _gdge := 0; _gdge < len(_bgfdc.Size); _gdge++ {
			_cbfcc = append(_cbfcc, 0)
			_cbfcc = append(_cbfcc, float64(_bgfdc.Size[_gdge]-1))
		}
	}
	_dgce := _bgfdc.Decode
	if _dgce == nil {
		_dgce = _bgfdc.Range
	}
	_bfdcd := make([]int, len(x))
	for _gbcfe := 0; _gbcfe < len(x); _gbcfe++ {
		_cbbd := x[_gbcfe]
		_babd := _ge.Min(_ge.Max(_cbbd, _bgfdc.Domain[2*_gbcfe]), _bgfdc.Domain[2*_gbcfe+1])
		_bgcea := _gca.LinearInterpolate(_babd, _bgfdc.Domain[2*_gbcfe], _bgfdc.Domain[2*_gbcfe+1], _cbfcc[2*_gbcfe], _cbfcc[2*_gbcfe+1])
		_fgfag := _ge.Min(_ge.Max(_bgcea, 0), float64(_bgfdc.Size[_gbcfe]-1))
		_gbaag := int(_ge.Floor(_fgfag + 0.5))
		if _gbaag < 0 {
			_gbaag = 0
		} else if _gbaag > _bgfdc.Size[_gbcfe] {
			_gbaag = _bgfdc.Size[_gbcfe] - 1
		}
		_bfdcd[_gbcfe] = _gbaag
	}
	_eaebf := _bfdcd[0]
	for _dbeb := 1; _dbeb < _bgfdc.NumInputs; _dbeb++ {
		_befdc := _bfdcd[_dbeb]
		for _bcdea := 0; _bcdea < _dbeb; _bcdea++ {
			_befdc *= _bgfdc.Size[_bcdea]
		}
		_eaebf += _befdc
	}
	_eaebf *= _bgfdc.NumOutputs
	var _beab []float64
	for _babg := 0; _babg < _bgfdc.NumOutputs; _babg++ {
		_dgef := _eaebf + _babg
		if _dgef >= len(_bgfdc._bega) {
			_acd.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u006e\u006ft\u0020\u0065\u006eo\u0075\u0067\u0068\u0020\u0069\u006ep\u0075\u0074\u0020sa\u006dp\u006c\u0065\u0073\u0020\u0074\u006f\u0020d\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0076\u0061lu\u0065\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			continue
		}
		_ccddg := _bgfdc._bega[_dgef]
		_fggfaa := _gca.LinearInterpolate(float64(_ccddg), 0, _ge.Pow(2, float64(_bgfdc.BitsPerSample)), _dgce[2*_babg], _dgce[2*_babg+1])
		_adgfa := _ge.Min(_ge.Max(_fggfaa, _bgfdc.Range[2*_babg]), _bgfdc.Range[2*_babg+1])
		_beab = append(_beab, _adgfa)
	}
	return _beab, nil
}

// GetColorspaceByName returns the colorspace with the specified name from the page resources.
func (_bagbb *PdfPageResources) GetColorspaceByName(keyName _abf.PdfObjectName) (PdfColorspace, bool) {
	_adgde, _eecdd := _bagbb.GetColorspaces()
	if _eecdd != nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _eecdd)
		return nil, false
	}
	if _adgde == nil {
		return nil, false
	}
	_dgcba, _ddadgc := _adgde.Colorspaces[string(keyName)]
	if !_ddadgc {
		return nil, false
	}
	return _dgcba, true
}

// PdfAnnotationTrapNet represents TrapNet annotations.
// (Section 12.5.6.21).
type PdfAnnotationTrapNet struct{ *PdfAnnotation }

func (_feceb *pdfFontType0) subsetRegistered() error {
	_dfafe, _cccg := _feceb.DescendantFont._gedca.(*pdfCIDFontType2)
	if !_cccg {
		_acd.Log.Debug("\u0046\u006fnt\u0020\u006e\u006ft\u0020\u0073\u0075\u0070por\u0074ed\u0020\u0066\u006f\u0072\u0020\u0073\u0075bs\u0065\u0074\u0074\u0069\u006e\u0067\u0020%\u0054", _feceb.DescendantFont)
		return nil
	}
	if _dfafe == nil {
		return nil
	}
	if _dfafe._dcbaf == nil {
		_acd.Log.Debug("\u004d\u0069\u0073si\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return nil
	}
	if _feceb._edeaf == nil {
		_acd.Log.Debug("\u004e\u006f\u0020e\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0073\u0075\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0067\u006e\u006f\u0072\u0065\u0064")
		return nil
	}
	_befa, _cccg := _abf.GetStream(_dfafe._dcbaf.FontFile2)
	if !_cccg {
		_acd.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _fd.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_gegcg, _aaged := _abf.DecodeStream(_befa)
	if _aaged != nil {
		_acd.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _aaged)
		return _aaged
	}
	_faeae, _aaged := _ab.Parse(_dd.NewReader(_gegcg))
	if _aaged != nil {
		_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_befa.Stream))
		return _aaged
	}
	var _bffd []rune
	var _cgeea *_ab.Font
	switch _gfea := _feceb._edeaf.(type) {
	case *_cbb.TrueTypeFontEncoder:
		_bffd = _gfea.RegisteredRunes()
		_cgeea, _aaged = _faeae.SubsetKeepRunes(_bffd)
		if _aaged != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aaged)
			return _aaged
		}
		_gfea.SubsetRegistered()
	case *_cbb.IdentityEncoder:
		_bffd = _gfea.RegisteredRunes()
		_bfda := make([]_ab.GlyphIndex, len(_bffd))
		for _fdbge, _aaefb := range _bffd {
			_bfda[_fdbge] = _ab.GlyphIndex(_aaefb)
		}
		_cgeea, _aaged = _faeae.SubsetKeepIndices(_bfda)
		if _aaged != nil {
			_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aaged)
			return _aaged
		}
	case _cbb.SimpleEncoder:
		_aagb := _gfea.Charcodes()
		for _, _cccc := range _aagb {
			_cefe, _dgdbb := _gfea.CharcodeToRune(_cccc)
			if !_dgdbb {
				_acd.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0063\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0020\u0074\u006f \u0072\u0075\u006e\u0065\u003a\u0020\u0025\u0064", _cccc)
				continue
			}
			_bffd = append(_bffd, _cefe)
		}
	default:
		return _e.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u006f\u0072\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069\u006eg\u003a\u0020\u0025\u0054", _feceb._edeaf)
	}
	var _aceae _dd.Buffer
	_aaged = _cgeea.Write(&_aceae)
	if _aaged != nil {
		_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aaged)
		return _aaged
	}
	if _feceb._aabfe != nil {
		_bfeg := make(map[_bd.CharCode]rune, len(_bffd))
		for _, _cedg := range _bffd {
			_geeb, _cgce := _feceb._edeaf.RuneToCharcode(_cedg)
			if !_cgce {
				continue
			}
			_bfeg[_bd.CharCode(_geeb)] = _cedg
		}
		_feceb._aabfe = _bd.NewToUnicodeCMap(_bfeg)
	}
	_befa, _aaged = _abf.MakeStream(_aceae.Bytes(), _abf.NewFlateEncoder())
	if _aaged != nil {
		_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aaged)
		return _aaged
	}
	_befa.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _abf.MakeInteger(int64(_aceae.Len())))
	if _fdcfc, _facbd := _abf.GetStream(_dfafe._dcbaf.FontFile2); _facbd {
		*_fdcfc = *_befa
	} else {
		_dfafe._dcbaf.FontFile2 = _befa
	}
	_dbgb := _geead()
	if len(_feceb._ecggf) > 0 {
		_feceb._ecggf = _fffc(_feceb._ecggf, _dbgb)
	}
	if len(_dfafe._ecggf) > 0 {
		_dfafe._ecggf = _fffc(_dfafe._ecggf, _dbgb)
	}
	if len(_feceb._dddac) > 0 {
		_feceb._dddac = _fffc(_feceb._dddac, _dbgb)
	}
	if _dfafe._dcbaf != nil {
		_ecba, _dcffe := _abf.GetName(_dfafe._dcbaf.FontName)
		if _dcffe && len(_ecba.String()) > 0 {
			_efefc := _fffc(_ecba.String(), _dbgb)
			_dfafe._dcbaf.FontName = _abf.MakeName(_efefc)
		}
	}
	return nil
}

// Mask returns the uin32 bitmask for the specific flag.
func (_cfdgd FieldFlag) Mask() uint32 { return uint32(_cfdgd) }

// ToPdfObject converts the pdfFontSimple to its PDF representation for outputting.
func (_cbga *pdfFontSimple) ToPdfObject() _abf.PdfObject {
	if _cbga._ddddaf == nil {
		_cbga._ddddaf = &_abf.PdfIndirectObject{}
	}
	_egeef := _cbga.baseFields().asPdfObjectDictionary("")
	_cbga._ddddaf.PdfObject = _egeef
	if _cbga.FirstChar != nil {
		_egeef.Set("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r", _cbga.FirstChar)
	}
	if _cbga.LastChar != nil {
		_egeef.Set("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072", _cbga.LastChar)
	}
	if _cbga.Widths != nil {
		_egeef.Set("\u0057\u0069\u0064\u0074\u0068\u0073", _cbga.Widths)
	}
	if _cbga.Encoding != nil {
		_egeef.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _cbga.Encoding)
	} else if _cbga._ebada != nil {
		_dbgad := _cbga._ebada.ToPdfObject()
		if _dbgad != nil {
			_egeef.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _dbgad)
		}
	}
	return _cbga._ddddaf
}

func _eccc(_fgeae *_abf.PdfObjectDictionary) (*PdfShadingType1, error) {
	_faebd := PdfShadingType1{}
	if _gbdff := _fgeae.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _gbdff != nil {
		_gbdff = _abf.TraceToDirectObject(_gbdff)
		_fgdbg, _aagedd := _gbdff.(*_abf.PdfObjectArray)
		if !_aagedd {
			_acd.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _gbdff)
			return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_faebd.Domain = _fgdbg
	}
	if _ggbeg := _fgeae.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _ggbeg != nil {
		_ggbeg = _abf.TraceToDirectObject(_ggbeg)
		_fcffe, _addcg := _ggbeg.(*_abf.PdfObjectArray)
		if !_addcg {
			_acd.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ggbeg)
			return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_faebd.Matrix = _fcffe
	}
	_cgdc := _fgeae.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _cgdc == nil {
		_acd.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_faebd.Function = []PdfFunction{}
	if _aebbe, _addcf := _cgdc.(*_abf.PdfObjectArray); _addcf {
		for _, _geff := range _aebbe.Elements() {
			_ddfgb, _cafca := _ebedg(_geff)
			if _cafca != nil {
				_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _cafca)
				return nil, _cafca
			}
			_faebd.Function = append(_faebd.Function, _ddfgb)
		}
	} else {
		_ffaad, _gcegcd := _ebedg(_cgdc)
		if _gcegcd != nil {
			_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _gcegcd)
			return nil, _gcegcd
		}
		_faebd.Function = append(_faebd.Function, _ffaad)
	}
	return &_faebd, nil
}

// PdfColorDeviceGray represents a grayscale color value that shall be represented by a single number in the
// range 0.0 to 1.0 where 0.0 corresponds to black and 1.0 to white.
type PdfColorDeviceGray float64

// Decrypt decrypts the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful,
// false otherwise.
func (_cccb *PdfReader) Decrypt(password []byte) (bool, error) {
	_cgceeb, _beccc := _cccb._bebc.Decrypt(password)
	if _beccc != nil {
		return false, _beccc
	}
	if !_cgceeb {
		return false, nil
	}
	_beccc = _cccb.loadStructure()
	if _beccc != nil {
		_acd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f \u006co\u0061d\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", _beccc)
		return false, _beccc
	}
	return true, nil
}

// GetRotate gets the inheritable rotate value, either from the page
// or a higher up page/pages struct.
func (_gcdad *PdfPage) GetRotate() (int64, error) {
	if _gcdad.Rotate != nil {
		return *_gcdad.Rotate, nil
	}
	_egac := _gcdad.Parent
	for _egac != nil {
		_bfcff, _gcfg := _abf.GetDict(_egac)
		if !_gcfg {
			return 0, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _cacb := _bfcff.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _cacb != nil {
			_fgfb, _fcccb := _abf.GetInt(_cacb)
			if !_fcccb {
				return 0, _fd.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0074a\u0074\u0065\u0020\u0076al\u0075\u0065")
			}
			if _fgfb != nil {
				return int64(*_fgfb), nil
			}
			return 0, _fd.New("\u0072\u006f\u0074\u0061te\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		_egac = _bfcff.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return 0, _fd.New("\u0072o\u0074a\u0074\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
}

// ToPdfObject returns the PDF representation of the colorspace.
func (_abcce *PdfColorspaceDeviceCMYK) ToPdfObject() _abf.PdfObject {
	return _abf.MakeName("\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b")
}

// SetContentStreams sets the content streams based on a string array. Will make
// 1 object stream for each string and reference from the page Contents.
// Each stream will be encoded using the encoding specified by the StreamEncoder,
// if empty, will use identity encoding (raw data).
func (_adga *PdfPage) SetContentStreams(cStreams []string, encoder _abf.StreamEncoder) error {
	if len(cStreams) == 0 {
		_adga.Contents = nil
		return nil
	}
	if encoder == nil {
		encoder = _abf.NewRawEncoder()
	}
	var _ebba []*_abf.PdfObjectStream
	for _, _fdga := range cStreams {
		_cdff := &_abf.PdfObjectStream{}
		_gbdcd := encoder.MakeStreamDict()
		_fccab, _dfgge := encoder.EncodeBytes([]byte(_fdga))
		if _dfgge != nil {
			return _dfgge
		}
		_gbdcd.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _abf.MakeInteger(int64(len(_fccab))))
		_cdff.PdfObjectDictionary = _gbdcd
		_cdff.Stream = _fccab
		_ebba = append(_ebba, _cdff)
	}
	if len(_ebba) == 1 {
		_adga.Contents = _ebba[0]
	} else {
		_gadba := _abf.MakeArray()
		for _, _geecf := range _ebba {
			_gadba.Append(_geecf)
		}
		_adga.Contents = _gadba
	}
	return nil
}

// PdfFunctionType3 defines stitching of the subdomains of several 1-input functions to produce
// a single new 1-input function.
type PdfFunctionType3 struct {
	Domain    []float64
	Range     []float64
	Functions []PdfFunction
	Bounds    []float64
	Encode    []float64
	_edacd    *_abf.PdfIndirectObject
}

func (_efdc *PdfColorspaceLab) String() string { return "\u004c\u0061\u0062" }

// AnnotFilterFunc represents a PDF annotation filtering function. If the function
// returns true, the annotation is kept, otherwise it is discarded.
type AnnotFilterFunc func(*PdfAnnotation) bool

// GetContentStream returns the pattern cell's content stream
func (_deee *PdfTilingPattern) GetContentStream() ([]byte, error) {
	_deegd, _, _cabdf := _deee.GetContentStreamWithEncoder()
	return _deegd, _cabdf
}

func (_bdcg *PdfReader) newPdfAnnotationLineFromDict(_agf *_abf.PdfObjectDictionary) (*PdfAnnotationLine, error) {
	_cdbb := PdfAnnotationLine{}
	_bec, _dcfa := _bdcg.newPdfAnnotationMarkupFromDict(_agf)
	if _dcfa != nil {
		return nil, _dcfa
	}
	_cdbb.PdfAnnotationMarkup = _bec
	_cdbb.L = _agf.Get("\u004c")
	_cdbb.BS = _agf.Get("\u0042\u0053")
	_cdbb.LE = _agf.Get("\u004c\u0045")
	_cdbb.IC = _agf.Get("\u0049\u0043")
	_cdbb.LL = _agf.Get("\u004c\u004c")
	_cdbb.LLE = _agf.Get("\u004c\u004c\u0045")
	_cdbb.Cap = _agf.Get("\u0043\u0061\u0070")
	_cdbb.IT = _agf.Get("\u0049\u0054")
	_cdbb.LLO = _agf.Get("\u004c\u004c\u004f")
	_cdbb.CP = _agf.Get("\u0043\u0050")
	_cdbb.Measure = _agf.Get("\u004de\u0061\u0073\u0075\u0072\u0065")
	_cdbb.CO = _agf.Get("\u0043\u004f")
	return &_cdbb, nil
}

// PdfColorCalRGB represents a color in the Colorimetric CIE RGB colorspace.
// A, B, C components
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorCalRGB [3]float64

// AddFont adds a font dictionary to the Font resources.
func (_ggced *PdfPage) AddFont(name _abf.PdfObjectName, font _abf.PdfObject) error {
	if _ggced.Resources == nil {
		_ggced.Resources = NewPdfPageResources()
	}
	if _ggced.Resources.Font == nil {
		_ggced.Resources.Font = _abf.MakeDict()
	}
	_cbgee, _bfacd := _abf.TraceToDirectObject(_ggced.Resources.Font).(*_abf.PdfObjectDictionary)
	if !_bfacd {
		_acd.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0066\u006f\u006et \u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a \u0025\u0076", _abf.TraceToDirectObject(_ggced.Resources.Font))
		return _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_cbgee.Set(name, font)
	return nil
}

func _fdade(_agefa *_abf.PdfObjectDictionary) (*PdfShadingType7, error) {
	_dgcdb := PdfShadingType7{}
	_geaae := _agefa.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _geaae == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_bbgce, _ceefb := _geaae.(*_abf.PdfObjectInteger)
	if !_ceefb {
		_acd.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _geaae)
		return nil, _abf.ErrTypeError
	}
	_dgcdb.BitsPerCoordinate = _bbgce
	_geaae = _agefa.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _geaae == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_bbgce, _ceefb = _geaae.(*_abf.PdfObjectInteger)
	if !_ceefb {
		_acd.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _geaae)
		return nil, _abf.ErrTypeError
	}
	_dgcdb.BitsPerComponent = _bbgce
	_geaae = _agefa.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _geaae == nil {
		_acd.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_bbgce, _ceefb = _geaae.(*_abf.PdfObjectInteger)
	if !_ceefb {
		_acd.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _geaae)
		return nil, _abf.ErrTypeError
	}
	_dgcdb.BitsPerComponent = _bbgce
	_geaae = _agefa.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _geaae == nil {
		_acd.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_dbcde, _ceefb := _geaae.(*_abf.PdfObjectArray)
	if !_ceefb {
		_acd.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _geaae)
		return nil, _abf.ErrTypeError
	}
	_dgcdb.Decode = _dbcde
	if _eecbd := _agefa.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _eecbd != nil {
		_dgcdb.Function = []PdfFunction{}
		if _aedge, _bfdeb := _eecbd.(*_abf.PdfObjectArray); _bfdeb {
			for _, _bcbd := range _aedge.Elements() {
				_cgcfd, _dfedf := _ebedg(_bcbd)
				if _dfedf != nil {
					_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dfedf)
					return nil, _dfedf
				}
				_dgcdb.Function = append(_dgcdb.Function, _cgcfd)
			}
		} else {
			_edbbbf, _egbbb := _ebedg(_eecbd)
			if _egbbb != nil {
				_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _egbbb)
				return nil, _egbbb
			}
			_dgcdb.Function = append(_dgcdb.Function, _edbbbf)
		}
	}
	return &_dgcdb, nil
}

func _bfbg(_bbfc _abf.PdfObject) (*PdfColorspaceCalRGB, error) {
	_bcfe := NewPdfColorspaceCalRGB()
	if _cbdfb, _daba := _bbfc.(*_abf.PdfIndirectObject); _daba {
		_bcfe._bdfg = _cbdfb
	}
	_bbfc = _abf.TraceToDirectObject(_bbfc)
	_gecf, _gcde := _bbfc.(*_abf.PdfObjectArray)
	if !_gcde {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _gecf.Len() != 2 {
		return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0043\u0061\u006c\u0052G\u0042 \u0063o\u006c\u006f\u0072\u0073\u0070\u0061\u0063e")
	}
	_bbfc = _abf.TraceToDirectObject(_gecf.Get(0))
	_dacg, _gcde := _bbfc.(*_abf.PdfObjectName)
	if !_gcde {
		return nil, _e.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	if *_dacg != "\u0043\u0061\u006c\u0052\u0047\u0042" {
		return nil, _e.Errorf("\u006e\u006f\u0074 a\u0020\u0043\u0061\u006c\u0052\u0047\u0042\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_bbfc = _abf.TraceToDirectObject(_gecf.Get(1))
	_acfb, _gcde := _bbfc.(*_abf.PdfObjectDictionary)
	if !_gcde {
		return nil, _e.Errorf("\u0043\u0061l\u0052\u0047\u0042\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0062je\u0063\u0074")
	}
	_bbfc = _acfb.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_bbfc = _abf.TraceToDirectObject(_bbfc)
	_dbg, _gcde := _bbfc.(*_abf.PdfObjectArray)
	if !_gcde {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0057\u0068\u0069\u0074\u0065\u0050o\u0069\u006e\u0074")
	}
	if _dbg.Len() != 3 {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0057h\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_bgcec, _ecca := _dbg.GetAsFloat64Slice()
	if _ecca != nil {
		return nil, _ecca
	}
	_bcfe.WhitePoint = _bgcec
	_bbfc = _acfb.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _bbfc != nil {
		_bbfc = _abf.TraceToDirectObject(_bbfc)
		_geac, _efc := _bbfc.(*_abf.PdfObjectArray)
		if !_efc {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076a\u006c\u0069\u0064\u0020\u0042\u006c\u0061\u0063\u006b\u0050o\u0069\u006e\u0074")
		}
		if _geac.Len() != 3 {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052\u0047\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0042l\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		_dgag, _addg := _geac.GetAsFloat64Slice()
		if _addg != nil {
			return nil, _addg
		}
		_bcfe.BlackPoint = _dgag
	}
	_bbfc = _acfb.Get("\u0047\u0061\u006dm\u0061")
	if _bbfc != nil {
		_bbfc = _abf.TraceToDirectObject(_bbfc)
		_fde, _fbfe := _bbfc.(*_abf.PdfObjectArray)
		if !_fbfe {
			return nil, _e.Errorf("C\u0061\u006c\u0052\u0047B:\u0020I\u006e\u0076\u0061\u006c\u0069d\u0020\u0047\u0061\u006d\u006d\u0061")
		}
		if _fde.Len() != 3 {
			return nil, _e.Errorf("C\u0061\u006c\u0052\u0047\u0042\u003a \u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0047a\u006d\u006d\u0061 \u0061r\u0072\u0061\u0079")
		}
		_aeecg, _eagdg := _fde.GetAsFloat64Slice()
		if _eagdg != nil {
			return nil, _eagdg
		}
		_bcfe.Gamma = _aeecg
	}
	_bbfc = _acfb.Get("\u004d\u0061\u0074\u0072\u0069\u0078")
	if _bbfc != nil {
		_bbfc = _abf.TraceToDirectObject(_bbfc)
		_abae, _ebcf := _bbfc.(*_abf.PdfObjectArray)
		if !_ebcf {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052GB\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004d\u0061\u0074\u0072i\u0078")
		}
		if _abae.Len() != 9 {
			_acd.Log.Error("\u004d\u0061t\u0072\u0069\u0078 \u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _abae.String())
			return nil, _e.Errorf("\u0043\u0061\u006c\u0052G\u0042\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u004da\u0074\u0072\u0069\u0078\u0020\u0061\u0072r\u0061\u0079")
		}
		_eecd, _cfgde := _abae.GetAsFloat64Slice()
		if _cfgde != nil {
			return nil, _cfgde
		}
		_bcfe.Matrix = _eecd
	}
	return _bcfe, nil
}

// NewPdfActionMovie returns a new "movie" action.
func NewPdfActionMovie() *PdfActionMovie {
	_gaa := NewPdfAction()
	_cfb := &PdfActionMovie{}
	_cfb.PdfAction = _gaa
	_gaa.SetContext(_cfb)
	return _cfb
}

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
func (_fdfb *PdfFont) CharcodeBytesToUnicode(data []byte) (string, int, int) {
	_ecac, _, _aaefg := _fdfb.CharcodesToUnicodeWithStats(_fdfb.BytesToCharcodes(data))
	_edfac := _cbb.ExpandLigatures(_ecac)
	return _edfac, _bc.RuneCountInString(_edfac), _aaefg
}

// ToPdfObject sets the common field elements.
// Note: Call the more field context's ToPdfObject to set both the generic and
// non-generic information.
func (_ccfe *PdfField) ToPdfObject() _abf.PdfObject {
	_ffdf := _ccfe._dgdc
	_fbeg := _ffdf.PdfObject.(*_abf.PdfObjectDictionary)
	_eadc := _abf.MakeArray()
	for _, _feafb := range _ccfe.Kids {
		_eadc.Append(_feafb.ToPdfObject())
	}
	for _, _bdbf := range _ccfe.Annotations {
		if _bdbf._dbc != _ccfe._dgdc {
			_eadc.Append(_bdbf.GetContext().ToPdfObject())
		}
	}
	if _ccfe.Parent != nil {
		_fbeg.SetIfNotNil("\u0050\u0061\u0072\u0065\u006e\u0074", _ccfe.Parent.GetContainingPdfObject())
	}
	if _eadc.Len() > 0 {
		_fbeg.Set("\u004b\u0069\u0064\u0073", _eadc)
	}
	_fbeg.SetIfNotNil("\u0046\u0054", _ccfe.FT)
	_fbeg.SetIfNotNil("\u0054", _ccfe.T)
	_fbeg.SetIfNotNil("\u0054\u0055", _ccfe.TU)
	_fbeg.SetIfNotNil("\u0054\u004d", _ccfe.TM)
	_fbeg.SetIfNotNil("\u0046\u0066", _ccfe.Ff)
	_fbeg.SetIfNotNil("\u0056", _ccfe.V)
	_fbeg.SetIfNotNil("\u0044\u0056", _ccfe.DV)
	_fbeg.SetIfNotNil("\u0041\u0041", _ccfe.AA)
	if _ccfe.VariableText != nil {
		_fbeg.SetIfNotNil("\u0044\u0041", _ccfe.VariableText.DA)
		_fbeg.SetIfNotNil("\u0051", _ccfe.VariableText.Q)
		_fbeg.SetIfNotNil("\u0044\u0053", _ccfe.VariableText.DS)
		_fbeg.SetIfNotNil("\u0052\u0056", _ccfe.VariableText.RV)
	}
	return _ffdf
}

// NewPdfActionRendition returns a new "rendition" action.
func NewPdfActionRendition() *PdfActionRendition {
	_cfe := NewPdfAction()
	_adbf := &PdfActionRendition{}
	_adbf.PdfAction = _cfe
	_cfe.SetContext(_adbf)
	return _adbf
}

// ToPdfObject implements interface PdfModel.
func (_dced *PdfAnnotationPolygon) ToPdfObject() _abf.PdfObject {
	_dced.PdfAnnotation.ToPdfObject()
	_fbaf := _dced._dbc
	_agfc := _fbaf.PdfObject.(*_abf.PdfObjectDictionary)
	_dced.PdfAnnotationMarkup.appendToPdfDictionary(_agfc)
	_agfc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0050o\u006c\u0079\u0067\u006f\u006e"))
	_agfc.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _dced.Vertices)
	_agfc.SetIfNotNil("\u004c\u0045", _dced.LE)
	_agfc.SetIfNotNil("\u0042\u0053", _dced.BS)
	_agfc.SetIfNotNil("\u0049\u0043", _dced.IC)
	_agfc.SetIfNotNil("\u0042\u0045", _dced.BE)
	_agfc.SetIfNotNil("\u0049\u0054", _dced.IT)
	_agfc.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _dced.Measure)
	return _fbaf
}

// PdfShadingType6 is a Coons patch mesh.
type PdfShadingType6 struct {
	*PdfShading
	BitsPerCoordinate *_abf.PdfObjectInteger
	BitsPerComponent  *_abf.PdfObjectInteger
	BitsPerFlag       *_abf.PdfObjectInteger
	Decode            *_abf.PdfObjectArray
	Function          []PdfFunction
}

// ToPdfObject returns a stream object.
func (_affgg *XObjectImage) ToPdfObject() _abf.PdfObject {
	_edeeg := _affgg._ccbad
	_gdbc := _edeeg.PdfObjectDictionary
	if _affgg.Filter != nil {
		_gdbc = _affgg.Filter.MakeStreamDict()
		_edeeg.PdfObjectDictionary = _gdbc
	}
	_gdbc.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0058O\u0062\u006a\u0065\u0063\u0074"))
	_gdbc.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0049\u006d\u0061g\u0065"))
	_gdbc.Set("\u0057\u0069\u0064t\u0068", _abf.MakeInteger(*(_affgg.Width)))
	_gdbc.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _abf.MakeInteger(*(_affgg.Height)))
	if _affgg.BitsPerComponent != nil {
		_gdbc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _abf.MakeInteger(*(_affgg.BitsPerComponent)))
	}
	if _affgg.ColorSpace != nil {
		_gdbc.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _affgg.ColorSpace.ToPdfObject())
	}
	_gdbc.SetIfNotNil("\u0049\u006e\u0074\u0065\u006e\u0074", _affgg.Intent)
	_gdbc.SetIfNotNil("\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k", _affgg.ImageMask)
	_gdbc.SetIfNotNil("\u004d\u0061\u0073\u006b", _affgg.Mask)
	_cfegg := _gdbc.Get("\u0044\u0065\u0063\u006f\u0064\u0065") != nil
	if _affgg.Decode == nil && _cfegg {
		_gdbc.Remove("\u0044\u0065\u0063\u006f\u0064\u0065")
	} else if _affgg.Decode != nil {
		_gdbc.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _affgg.Decode)
	}
	_gdbc.SetIfNotNil("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _affgg.Interpolate)
	_gdbc.SetIfNotNil("\u0041\u006c\u0074e\u0072\u006e\u0061\u0074\u0069\u0076\u0065\u0073", _affgg.Alternatives)
	_gdbc.SetIfNotNil("\u0053\u004d\u0061s\u006b", _affgg.SMask)
	_gdbc.SetIfNotNil("S\u004d\u0061\u0073\u006b\u0049\u006e\u0044\u0061\u0074\u0061", _affgg.SMaskInData)
	_gdbc.SetIfNotNil("\u004d\u0061\u0074t\u0065", _affgg.Matte)
	_gdbc.SetIfNotNil("\u004e\u0061\u006d\u0065", _affgg.Name)
	_gdbc.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _affgg.StructParent)
	_gdbc.SetIfNotNil("\u0049\u0044", _affgg.ID)
	_gdbc.SetIfNotNil("\u004f\u0050\u0049", _affgg.OPI)
	_gdbc.SetIfNotNil("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", _affgg.Metadata)
	_gdbc.SetIfNotNil("\u004f\u0043", _affgg.OC)
	_gdbc.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _abf.MakeInteger(int64(len(_affgg.Stream))))
	_edeeg.Stream = _affgg.Stream
	return _edeeg
}

// NewGrayImageFromGoImage creates a new grayscale unidoc Image from a golang Image.
func (_cfecg DefaultImageHandler) NewGrayImageFromGoImage(goimg _aa.Image) (*Image, error) {
	_gaea := goimg.Bounds()
	_eaaf := &Image{Width: int64(_gaea.Dx()), Height: int64(_gaea.Dy()), ColorComponents: 1, BitsPerComponent: 8}
	switch _dbabb := goimg.(type) {
	case *_aa.Gray:
		if len(_dbabb.Pix) != _gaea.Dx()*_gaea.Dy() {
			_acgea, _ccede := _gca.GrayConverter.Convert(goimg)
			if _ccede != nil {
				return nil, _ccede
			}
			_eaaf.Data = _acgea.Pix()
		} else {
			_eaaf.Data = _dbabb.Pix
		}
	case *_aa.Gray16:
		_eaaf.BitsPerComponent = 16
		if len(_dbabb.Pix) != _gaea.Dx()*_gaea.Dy()*2 {
			_cfdea, _fccbf := _gca.Gray16Converter.Convert(goimg)
			if _fccbf != nil {
				return nil, _fccbf
			}
			_eaaf.Data = _cfdea.Pix()
		} else {
			_eaaf.Data = _dbabb.Pix
		}
	case _gca.Image:
		_dcaee := _dbabb.Base()
		if _dcaee.ColorComponents == 1 {
			_eaaf.BitsPerComponent = int64(_dcaee.BitsPerComponent)
			_eaaf.Data = _dcaee.Data
			return _eaaf, nil
		}
		_egbed, _acaeg := _gca.GrayConverter.Convert(goimg)
		if _acaeg != nil {
			return nil, _acaeg
		}
		_eaaf.Data = _egbed.Pix()
	default:
		_gafca, _fcec := _gca.GrayConverter.Convert(goimg)
		if _fcec != nil {
			return nil, _fcec
		}
		_eaaf.Data = _gafca.Pix()
	}
	return _eaaf, nil
}

func (_bgfg *PdfAppender) addNewObject(_efbg _abf.PdfObject) {
	if _, _bede := _bgfg._gcba[_efbg]; !_bede {
		_bgfg._ffcf = append(_bgfg._ffcf, _efbg)
		_bgfg._gcba[_efbg] = struct{}{}
	}
}

func (_dgcd *PdfReader) newPdfPageFromDict(_defe *_abf.PdfObjectDictionary) (*PdfPage, error) {
	_affb := NewPdfPage()
	_affb._bdbfa = _defe
	_affb._efca = *_defe
	_bbfbc := *_defe
	_edabe, _eaff := _bbfbc.Get("\u0054\u0079\u0070\u0065").(*_abf.PdfObjectName)
	if !_eaff {
		return nil, _fd.New("\u006d\u0069ss\u0069\u006e\u0067/\u0069\u006e\u0076\u0061lid\u0020Pa\u0067\u0065\u0020\u0064\u0069\u0063\u0074io\u006e\u0061\u0072\u0079\u0020\u0054\u0079p\u0065")
	}
	if *_edabe != "\u0050\u0061\u0067\u0065" {
		return nil, _fd.New("\u0070\u0061\u0067\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u0050a\u0067\u0065")
	}
	if _efgee := _bbfbc.Get("\u0050\u0061\u0072\u0065\u006e\u0074"); _efgee != nil {
		_affb.Parent = _efgee
	}
	if _ccbea := _bbfbc.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"); _ccbea != nil {
		_dgfee, _aefd := _abf.GetString(_ccbea)
		if !_aefd {
			return nil, _fd.New("\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u004c\u0061\u0073\u0074\u004d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0021=\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		_gbfgf, _gdac := NewPdfDate(_dgfee.Str())
		if _gdac != nil {
			return nil, _gdac
		}
		_affb.LastModified = &_gbfgf
	}
	if _fbba := _bbfbc.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _fbba != nil && !_abf.IsNullObject(_fbba) {
		_ecbe, _dgfbc := _abf.GetDict(_fbba)
		if !_dgfbc {
			return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0054\u0029", _fbba)
		}
		var _bgfcb error
		_affb.Resources, _bgfcb = NewPdfPageResourcesFromDict(_ecbe)
		if _bgfcb != nil {
			return nil, _bgfcb
		}
	} else {
		_bdaag, _baeab := _affb.getParentResources()
		if _baeab != nil {
			return nil, _baeab
		}
		if _bdaag == nil {
			_bdaag = NewPdfPageResources()
		}
		_affb.Resources = _bdaag
	}
	if _bgbg := _bbfbc.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _bgbg != nil {
		_dcfga, _ceacd := _abf.GetArray(_bgbg)
		if !_ceacd {
			return nil, _fd.New("\u0070\u0061\u0067\u0065\u0020\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _facgd error
		_affb.MediaBox, _facgd = NewPdfRectangle(*_dcfga)
		if _facgd != nil {
			return nil, _facgd
		}
	}
	if _gcbdd := _bbfbc.Get("\u0043r\u006f\u0070\u0042\u006f\u0078"); _gcbdd != nil {
		_cbgga, _gdggc := _abf.GetArray(_gcbdd)
		if !_gdggc {
			return nil, _fd.New("\u0070a\u0067\u0065\u0020\u0043r\u006f\u0070\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _ggbce error
		_affb.CropBox, _ggbce = NewPdfRectangle(*_cbgga)
		if _ggbce != nil {
			return nil, _ggbce
		}
	}
	if _beafbf := _bbfbc.Get("\u0042\u006c\u0065\u0065\u0064\u0042\u006f\u0078"); _beafbf != nil {
		_ceff, _caef := _abf.GetArray(_beafbf)
		if !_caef {
			return nil, _fd.New("\u0070\u0061\u0067\u0065\u0020\u0042\u006c\u0065\u0065\u0064\u0042o\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072\u0061\u0079")
		}
		var _cbbae error
		_affb.BleedBox, _cbbae = NewPdfRectangle(*_ceff)
		if _cbbae != nil {
			return nil, _cbbae
		}
	}
	if _afgbg := _bbfbc.Get("\u0054r\u0069\u006d\u0042\u006f\u0078"); _afgbg != nil {
		_degfg, _ddfga := _abf.GetArray(_afgbg)
		if !_ddfga {
			return nil, _fd.New("\u0070a\u0067\u0065\u0020\u0054r\u0069\u006d\u0042\u006f\u0078 \u006eo\u0074 \u0061\u006e\u0020\u0061\u0072\u0072\u0061y")
		}
		var _dbgfc error
		_affb.TrimBox, _dbgfc = NewPdfRectangle(*_degfg)
		if _dbgfc != nil {
			return nil, _dbgfc
		}
	}
	if _aeeb := _bbfbc.Get("\u0041\u0072\u0074\u0042\u006f\u0078"); _aeeb != nil {
		_gddg, _cagcc := _abf.GetArray(_aeeb)
		if !_cagcc {
			return nil, _fd.New("\u0070a\u0067\u0065\u0020\u0041\u0072\u0074\u0042\u006f\u0078\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		}
		var _fcaac error
		_affb.ArtBox, _fcaac = NewPdfRectangle(*_gddg)
		if _fcaac != nil {
			return nil, _fcaac
		}
	}
	if _fbca := _bbfbc.Get("\u0042\u006f\u0078C\u006f\u006c\u006f\u0072\u0049\u006e\u0066\u006f"); _fbca != nil {
		_affb.BoxColorInfo = _fbca
	}
	if _fgdeb := _bbfbc.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"); _fgdeb != nil {
		_affb.Contents = _fgdeb
	}
	if _cgbff := _bbfbc.Get("\u0052\u006f\u0074\u0061\u0074\u0065"); _cgbff != nil {
		_ebag, _ccaa := _abf.GetNumberAsInt64(_cgbff)
		if _ccaa != nil {
			return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0067e\u0020\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		_affb.Rotate = &_ebag
	}
	if _fcfcf := _bbfbc.Get("\u0047\u0072\u006fu\u0070"); _fcfcf != nil {
		_affb.Group = _fcfcf
	}
	if _fbdcg := _bbfbc.Get("\u0054\u0068\u0075m\u0062"); _fbdcg != nil {
		_affb.Thumb = _fbdcg
	}
	if _cecef := _bbfbc.Get("\u0042"); _cecef != nil {
		_affb.B = _cecef
	}
	if _ccaee := _bbfbc.Get("\u0044\u0075\u0072"); _ccaee != nil {
		_affb.Dur = _ccaee
	}
	if _ecegg := _bbfbc.Get("\u0054\u0072\u0061n\u0073"); _ecegg != nil {
		_affb.Trans = _ecegg
	}
	if _acggc := _bbfbc.Get("\u0041\u0041"); _acggc != nil {
		_affb.AA = _acggc
	}
	if _defdg := _bbfbc.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"); _defdg != nil {
		_affb.Metadata = _defdg
	}
	if _cfbac := _bbfbc.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o"); _cfbac != nil {
		_affb.PieceInfo = _cfbac
	}
	if _baceg := _bbfbc.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073"); _baceg != nil {
		_affb.StructParents = _baceg
	}
	if _eced := _bbfbc.Get("\u0049\u0044"); _eced != nil {
		_affb.ID = _eced
	}
	if _begbb := _bbfbc.Get("\u0050\u005a"); _begbb != nil {
		_affb.PZ = _begbb
	}
	if _egefde := _bbfbc.Get("\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006fn\u0049\u006e\u0066\u006f"); _egefde != nil {
		_affb.SeparationInfo = _egefde
	}
	if _dedc := _bbfbc.Get("\u0054\u0061\u0062\u0073"); _dedc != nil {
		_affb.Tabs = _dedc
	}
	if _cdgaa := _bbfbc.Get("T\u0065m\u0070\u006c\u0061\u0074\u0065\u0049\u006e\u0073t\u0061\u006e\u0074\u0069at\u0065\u0064"); _cdgaa != nil {
		_affb.TemplateInstantiated = _cdgaa
	}
	if _cafd := _bbfbc.Get("\u0050r\u0065\u0073\u0053\u0074\u0065\u0070s"); _cafd != nil {
		_affb.PresSteps = _cafd
	}
	if _cfccd := _bbfbc.Get("\u0055\u0073\u0065\u0072\u0055\u006e\u0069\u0074"); _cfccd != nil {
		_affb.UserUnit = _cfccd
	}
	if _cgdea := _bbfbc.Get("\u0056\u0050"); _cgdea != nil {
		_affb.VP = _cgdea
	}
	if _cfeg := _bbfbc.Get("\u0041\u006e\u006e\u006f\u0074\u0073"); _cfeg != nil {
		_affb.Annots = _cfeg
	}
	_affb._dbaef = _dgcd
	return _affb, nil
}

// SetType sets the field button's type.  Can be one of:
// - PdfFieldButtonPush for push button fields
// - PdfFieldButtonCheckbox for checkbox fields
// - PdfFieldButtonRadio for radio button fields
// This sets the field's flag appropriately.
func (_fdefc *PdfFieldButton) SetType(btype ButtonType) {
	_gbaa := uint32(0)
	if _fdefc.Ff != nil {
		_gbaa = uint32(*_fdefc.Ff)
	}
	switch btype {
	case ButtonTypePush:
		_gbaa |= FieldFlagPushbutton.Mask()
	case ButtonTypeRadio:
		_gbaa |= FieldFlagRadio.Mask()
	}
	_fdefc.Ff = _abf.MakeInteger(int64(_gbaa))
}

func _bcgcee(_bdgg _abf.PdfObject) {
	_acd.Log.Debug("\u006f\u0062\u006a\u003a\u0020\u0025\u0054\u0020\u0025\u0073", _bdgg, _bdgg.String())
	if _afaeg, _dgfae := _bdgg.(*_abf.PdfObjectStream); _dgfae {
		_ecgag, _aadd := _abf.DecodeStream(_afaeg)
		if _aadd != nil {
			_acd.Log.Debug("\u0045r\u0072\u006f\u0072\u003a\u0020\u0025v", _aadd)
			return
		}
		_acd.Log.Debug("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _ecgag)
	} else if _afbcd, _cfgfe := _bdgg.(*_abf.PdfIndirectObject); _cfgfe {
		_acd.Log.Debug("\u0025\u0054\u0020%\u0076", _afbcd.PdfObject, _afbcd.PdfObject)
		_acd.Log.Debug("\u0025\u0073", _afbcd.PdfObject.String())
	}
}

func (_dabbg *PdfAcroForm) filteredFields(_cgaf FieldFilterFunc, _ceced bool) []*PdfField {
	if _dabbg == nil {
		return nil
	}
	return _fccaa(_dabbg.Fields, _cgaf, _ceced)
}

// NewPdfActionImportData returns a new "import data" action.
func NewPdfActionImportData() *PdfActionImportData {
	_fce := NewPdfAction()
	_aeb := &PdfActionImportData{}
	_aeb.PdfAction = _fce
	_fce.SetContext(_aeb)
	return _aeb
}

// Enable LTV enables the specified signature. The signing certificate
// chain is extracted from the signature dictionary. Optionally, additional
// certificates can be specified through the `extraCerts` parameter.
// The LTV client attempts to build the certificate chain up to a trusted root
// by downloading any missing certificates.
func (_cdbc *LTV) Enable(sig *PdfSignature, extraCerts []*_fa.Certificate) error {
	if _dbcfe := _cdbc.validateSig(sig); _dbcfe != nil {
		return _dbcfe
	}
	_bcdc, _cfeec := _cdbc.generateVRIKey(sig)
	if _cfeec != nil {
		return _cfeec
	}
	if _, _dfbg := _cdbc._dgfe.VRI[_bcdc]; _dfbg && _cdbc.SkipExisting {
		return nil
	}
	_afacd, _cfeec := sig.GetCerts()
	if _cfeec != nil {
		return _cfeec
	}
	return _cdbc.enable(_afacd, extraCerts, _bcdc)
}

// Encoder returns the font's text encoder.
func (_efbfa pdfFontType3) Encoder() _cbb.TextEncoder { return _efbfa._dgbd }

// PdfActionType represents an action type in PDF (section 12.6.4 p. 417).
type PdfActionType string

// GetPreviousRevision returns the previous revision of PdfReader for the Pdf document
func (_fbdfc *PdfReader) GetPreviousRevision() (*PdfReader, error) {
	if _fbdfc._bebc.GetRevisionNumber() == 0 {
		return nil, _fd.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065xi\u0073\u0074")
	}
	if _cfcba, _fgfee := _fbdfc._bfced[_fbdfc]; _fgfee {
		return _cfcba, nil
	}
	_dggde, _fbdg := _fbdfc._bebc.GetPreviousRevisionReadSeeker()
	if _fbdg != nil {
		return nil, _fbdg
	}
	_ebaaa, _fbdg := _fbaec(_dggde, _fbdfc._gebfg, _fbdfc._dbgdg, "\u006do\u0064\u0065\u006c\u003aG\u0065\u0074\u0050\u0072\u0065v\u0069o\u0075s\u0052\u0065\u0076\u0069\u0073\u0069\u006fn")
	if _fbdg != nil {
		return nil, _fbdg
	}
	_fbdfc._egade[_fbdfc._bebc.GetRevisionNumber()-1] = _ebaaa
	_fbdfc._bfced[_fbdfc] = _ebaaa
	_ebaaa._bfced = _fbdfc._bfced
	return _ebaaa, nil
}

// IsPush returns true if the button field represents a push button, false otherwise.
func (_ebaa *PdfFieldButton) IsPush() bool { return _ebaa.GetType() == ButtonTypePush }

func (_gegce *LTV) getOCSPs(_bfddf []*_fa.Certificate, _ededb map[string]*_fa.Certificate) ([][]byte, error) {
	_ecdgg := make([][]byte, 0, len(_bfddf))
	for _, _ebfgf := range _bfddf {
		for _, _ebbbc := range _ebfgf.OCSPServer {
			if _gegce.CertClient.IsCA(_ebfgf) {
				continue
			}
			_fgbd, _bgbec := _ededb[_ebfgf.Issuer.CommonName]
			if !_bgbec {
				_acd.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _dgfgf, _ggae := _gegce.OCSPClient.MakeRequest(_ebbbc, _ebfgf, _fgbd)
			if _ggae != nil {
				_acd.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _ggae)
				continue
			}
			_ecdgg = append(_ecdgg, _dgfgf)
		}
	}
	return _ecdgg, nil
}

// PdfShadingPatternType2 is shading patterns that will use a Type 2 shading pattern (Axial).
type PdfShadingPatternType2 struct {
	*PdfPattern
	Shading   *PdfShadingType2
	Matrix    *_abf.PdfObjectArray
	ExtGState _abf.PdfObject
}

// StandardValidator is the interface that is used for the PDF StandardImplementer validation for the PDF document.
// It is using a CompliancePdfReader which is expected to give more Metadata during reading process.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type StandardValidator interface {
	// ValidateStandard checks if the input reader
	ValidateStandard(_afede *CompliancePdfReader) error
}

func _gabc(_cbdd *PdfPage) map[_abf.PdfObjectName]_abf.PdfObject {
	_bfeb := make(map[_abf.PdfObjectName]_abf.PdfObject)
	if _cbdd.Resources == nil {
		return _bfeb
	}
	if _cbdd.Resources.Font != nil {
		if _fadcg, _fbf := _abf.GetDict(_cbdd.Resources.Font); _fbf {
			for _, _bcgg := range _fadcg.Keys() {
				_bfeb[_bcgg] = _fadcg.Get(_bcgg)
			}
		}
	}
	if _cbdd.Resources.ExtGState != nil {
		if _dabd, _abbd := _abf.GetDict(_cbdd.Resources.ExtGState); _abbd {
			for _, _ceea := range _dabd.Keys() {
				_bfeb[_ceea] = _dabd.Get(_ceea)
			}
		}
	}
	if _cbdd.Resources.XObject != nil {
		if _adba, _defd := _abf.GetDict(_cbdd.Resources.XObject); _defd {
			for _, _beeb := range _adba.Keys() {
				_bfeb[_beeb] = _adba.Get(_beeb)
			}
		}
	}
	if _cbdd.Resources.Pattern != nil {
		if _becg, _fcca := _abf.GetDict(_cbdd.Resources.Pattern); _fcca {
			for _, _ddeg := range _becg.Keys() {
				_bfeb[_ddeg] = _becg.Get(_ddeg)
			}
		}
	}
	if _cbdd.Resources.Shading != nil {
		if _fdbcg, _ebcbg := _abf.GetDict(_cbdd.Resources.Shading); _ebcbg {
			for _, _ffae := range _fdbcg.Keys() {
				_bfeb[_ffae] = _fdbcg.Get(_ffae)
			}
		}
	}
	if _cbdd.Resources.ProcSet != nil {
		if _cdfbc, _edfb := _abf.GetDict(_cbdd.Resources.ProcSet); _edfb {
			for _, _deea := range _cdfbc.Keys() {
				_bfeb[_deea] = _cdfbc.Get(_deea)
			}
		}
	}
	if _cbdd.Resources.Properties != nil {
		if _eada, _ddffg := _abf.GetDict(_cbdd.Resources.Properties); _ddffg {
			for _, _ggde := range _eada.Keys() {
				_bfeb[_ggde] = _eada.Get(_ggde)
			}
		}
	}
	return _bfeb
}

// ToPdfObject returns the PdfFontDescriptor as a PDF dictionary inside an indirect object.
func (_fdcd *PdfFontDescriptor) ToPdfObject() _abf.PdfObject {
	_dcagb := _abf.MakeDict()
	if _fdcd._aage == nil {
		_fdcd._aage = &_abf.PdfIndirectObject{}
	}
	_fdcd._aage.PdfObject = _dcagb
	_dcagb.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072"))
	if _fdcd.FontName != nil {
		_dcagb.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _fdcd.FontName)
	}
	if _fdcd.FontFamily != nil {
		_dcagb.Set("\u0046\u006f\u006e\u0074\u0046\u0061\u006d\u0069\u006c\u0079", _fdcd.FontFamily)
	}
	if _fdcd.FontStretch != nil {
		_dcagb.Set("F\u006f\u006e\u0074\u0053\u0074\u0072\u0065\u0074\u0063\u0068", _fdcd.FontStretch)
	}
	if _fdcd.FontWeight != nil {
		_dcagb.Set("\u0046\u006f\u006e\u0074\u0057\u0065\u0069\u0067\u0068\u0074", _fdcd.FontWeight)
	}
	if _fdcd.Flags != nil {
		_dcagb.Set("\u0046\u006c\u0061g\u0073", _fdcd.Flags)
	}
	if _fdcd.FontBBox != nil {
		_dcagb.Set("\u0046\u006f\u006e\u0074\u0042\u0042\u006f\u0078", _fdcd.FontBBox)
	}
	if _fdcd.ItalicAngle != nil {
		_dcagb.Set("I\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065", _fdcd.ItalicAngle)
	}
	if _fdcd.Ascent != nil {
		_dcagb.Set("\u0041\u0073\u0063\u0065\u006e\u0074", _fdcd.Ascent)
	}
	if _fdcd.Descent != nil {
		_dcagb.Set("\u0044e\u0073\u0063\u0065\u006e\u0074", _fdcd.Descent)
	}
	if _fdcd.Leading != nil {
		_dcagb.Set("\u004ce\u0061\u0064\u0069\u006e\u0067", _fdcd.Leading)
	}
	if _fdcd.CapHeight != nil {
		_dcagb.Set("\u0043a\u0070\u0048\u0065\u0069\u0067\u0068t", _fdcd.CapHeight)
	}
	if _fdcd.XHeight != nil {
		_dcagb.Set("\u0058H\u0065\u0069\u0067\u0068\u0074", _fdcd.XHeight)
	}
	if _fdcd.StemV != nil {
		_dcagb.Set("\u0053\u0074\u0065m\u0056", _fdcd.StemV)
	}
	if _fdcd.StemH != nil {
		_dcagb.Set("\u0053\u0074\u0065m\u0048", _fdcd.StemH)
	}
	if _fdcd.AvgWidth != nil {
		_dcagb.Set("\u0041\u0076\u0067\u0057\u0069\u0064\u0074\u0068", _fdcd.AvgWidth)
	}
	if _fdcd.MaxWidth != nil {
		_dcagb.Set("\u004d\u0061\u0078\u0057\u0069\u0064\u0074\u0068", _fdcd.MaxWidth)
	}
	if _fdcd.MissingWidth != nil {
		_dcagb.Set("\u004d\u0069\u0073s\u0069\u006e\u0067\u0057\u0069\u0064\u0074\u0068", _fdcd.MissingWidth)
	}
	if _fdcd.FontFile != nil {
		_dcagb.Set("\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065", _fdcd.FontFile)
	}
	if _fdcd.FontFile2 != nil {
		_dcagb.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00652", _fdcd.FontFile2)
	}
	if _fdcd.FontFile3 != nil {
		_dcagb.Set("\u0046o\u006e\u0074\u0046\u0069\u006c\u00653", _fdcd.FontFile3)
	}
	if _fdcd.CharSet != nil {
		_dcagb.Set("\u0043h\u0061\u0072\u0053\u0065\u0074", _fdcd.CharSet)
	}
	if _fdcd.Style != nil {
		_dcagb.Set("\u0046\u006f\u006e\u0074\u004e\u0061\u006d\u0065", _fdcd.FontName)
	}
	if _fdcd.Lang != nil {
		_dcagb.Set("\u004c\u0061\u006e\u0067", _fdcd.Lang)
	}
	if _fdcd.FD != nil {
		_dcagb.Set("\u0046\u0044", _fdcd.FD)
	}
	if _fdcd.CIDSet != nil {
		_dcagb.Set("\u0043\u0049\u0044\u0053\u0065\u0074", _fdcd.CIDSet)
	}
	return _fdcd._aage
}

// ValidateSignatures validates digital signatures in the document.
func (_cbfff *PdfReader) ValidateSignatures(handlers []SignatureHandler) ([]SignatureValidationResult, error) {
	if _cbfff.AcroForm == nil {
		return nil, nil
	}
	if _cbfff.AcroForm.Fields == nil {
		return nil, nil
	}
	type sigFieldPair struct {
		_abaeea *PdfSignature
		_aegbab *PdfField
		_fdggd  SignatureHandler
	}
	var _egagfc []*sigFieldPair
	for _, _abggd := range _cbfff.AcroForm.AllFields() {
		if _abggd.V == nil {
			continue
		}
		if _dfebb, _eagaf := _abf.GetDict(_abggd.V); _eagaf {
			if _afgeff, _gacf := _abf.GetNameVal(_dfebb.Get("\u0054\u0079\u0070\u0065")); _gacf && (_afgeff == "\u0053\u0069\u0067" || _afgeff == "\u0044\u006f\u0063T\u0069\u006d\u0065\u0053\u0074\u0061\u006d\u0070") {
				_eacag, _gcgbd := _abf.GetIndirect(_abggd.V)
				if !_gcgbd {
					_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0020\u0069s\u0020\u006e\u0069\u006c")
					return nil, ErrTypeCheck
				}
				_effcg, _acebb := _cbfff.newPdfSignatureFromIndirect(_eacag)
				if _acebb != nil {
					return nil, _acebb
				}
				var _defadd SignatureHandler
				for _, _dcecf := range handlers {
					if _dcecf.IsApplicable(_effcg) {
						_defadd = _dcecf
						break
					}
				}
				_egagfc = append(_egagfc, &sigFieldPair{_abaeea: _effcg, _aegbab: _abggd, _fdggd: _defadd})
			}
		}
	}
	var _bbeba []SignatureValidationResult
	for _, _bbbfe := range _egagfc {
		_bfcca := SignatureValidationResult{IsSigned: true, Fields: []*PdfField{_bbbfe._aegbab}}
		if _bbbfe._fdggd == nil {
			_bfcca.Errors = append(_bfcca.Errors, "\u0068a\u006ed\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
			_bbeba = append(_bbeba, _bfcca)
			continue
		}
		_fcfa, _cebgfc := _bbbfe._fdggd.NewDigest(_bbbfe._abaeea)
		if _cebgfc != nil {
			_bfcca.Errors = append(_bfcca.Errors, "\u0064\u0069\u0067e\u0073\u0074\u0020\u0065\u0072\u0072\u006f\u0072", _cebgfc.Error())
			_bbeba = append(_bbeba, _bfcca)
			continue
		}
		_eddbe := _bbbfe._abaeea.ByteRange
		if _eddbe == nil {
			_bfcca.Errors = append(_bfcca.Errors, "\u0042\u0079\u0074\u0065\u0052\u0061\u006e\u0067\u0065\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
			_bbeba = append(_bbeba, _bfcca)
			continue
		}
		for _gbaee := 0; _gbaee < _eddbe.Len(); _gbaee = _gbaee + 2 {
			_cbdfg, _ := _abf.GetNumberAsInt64(_eddbe.Get(_gbaee))
			_ebfaa, _ := _abf.GetIntVal(_eddbe.Get(_gbaee + 1))
			if _, _dafgg := _cbfff._affbb.Seek(_cbdfg, _gc.SeekStart); _dafgg != nil {
				return nil, _dafgg
			}
			_efdagd := make([]byte, _ebfaa)
			if _, _ccgae := _cbfff._affbb.Read(_efdagd); _ccgae != nil {
				return nil, _ccgae
			}
			_fcfa.Write(_efdagd)
		}
		var _bdag SignatureValidationResult
		if _gaag, _baccd := _bbbfe._fdggd.(SignatureHandlerDocMDP); _baccd {
			_bdag, _cebgfc = _gaag.ValidateWithOpts(_bbbfe._abaeea, _fcfa, SignatureHandlerDocMDPParams{Parser: _cbfff._bebc})
		} else {
			_bdag, _cebgfc = _bbbfe._fdggd.Validate(_bbbfe._abaeea, _fcfa)
		}
		if _cebgfc != nil {
			_acd.Log.Debug("E\u0052\u0052\u004f\u0052: \u0025v\u0020\u0028\u0025\u0054\u0029 \u002d\u0020\u0073\u006b\u0069\u0070", _cebgfc, _bbbfe._fdggd)
			_bdag.Errors = append(_bdag.Errors, _cebgfc.Error())
		}
		_bdag.Name = _bbbfe._abaeea.Name.Decoded()
		_bdag.Reason = _bbbfe._abaeea.Reason.Decoded()
		if _bbbfe._abaeea.M != nil {
			_ccdg, _gfdgc := NewPdfDate(_bbbfe._abaeea.M.String())
			if _gfdgc != nil {
				_acd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gfdgc)
				_bdag.Errors = append(_bdag.Errors, _gfdgc.Error())
				continue
			}
			_bdag.Date = _ccdg
		}
		_bdag.ContactInfo = _bbbfe._abaeea.ContactInfo.Decoded()
		_bdag.Location = _bbbfe._abaeea.Location.Decoded()
		_bdag.Fields = _bfcca.Fields
		_bbeba = append(_bbeba, _bdag)
	}
	return _bbeba, nil
}

func _bffaa(_eagae *XObjectImage) error {
	if _eagae.SMask == nil {
		return nil
	}
	_beddgf, _fcdce := _eagae.SMask.(*_abf.PdfObjectStream)
	if !_fcdce {
		_acd.Log.Debug("\u0053\u004da\u0073\u006b\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0053\u0074re\u0061\u006d")
		return _abf.ErrTypeError
	}
	_gecbb := _beddgf.PdfObjectDictionary
	_bgfeb := _gecbb.Get("\u004d\u0061\u0074t\u0065")
	if _bgfeb == nil {
		return nil
	}
	_affdd, _bfaec := _efggg(_bgfeb.(*_abf.PdfObjectArray))
	if _bfaec != nil {
		return _bfaec
	}
	_feeff := _abf.MakeArrayFromFloats([]float64{_affdd})
	_gecbb.SetIfNotNil("\u004d\u0061\u0074t\u0065", _feeff)
	return nil
}

func (_ebdcaa *PdfReader) newPdfSignatureFromIndirect(_afdfg *_abf.PdfIndirectObject) (*PdfSignature, error) {
	_bbdca, _dada := _afdfg.PdfObject.(*_abf.PdfObjectDictionary)
	if !_dada {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020a \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		return nil, ErrTypeCheck
	}
	if _dgfgd, _ddbeb := _ebdcaa._ceecd.GetModelFromPrimitive(_afdfg).(*PdfSignature); _ddbeb {
		return _dgfgd, nil
	}
	_fedae := &PdfSignature{}
	_fedae._geebd = _afdfg
	_fedae.Type, _ = _abf.GetName(_bbdca.Get("\u0054\u0079\u0070\u0065"))
	_fedae.Filter, _dada = _abf.GetName(_bbdca.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if !_dada {
		_acd.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053i\u0067\u006e\u0061\u0074\u0075r\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	_fedae.SubFilter, _ = _abf.GetName(_bbdca.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r"))
	_fedae.Contents, _dada = _abf.GetString(_bbdca.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_dada {
		_acd.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
		return nil, ErrInvalidAttribute
	}
	if _cbefgg, _bafba := _abf.GetArray(_bbdca.Get("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e")); _bafba {
		_fedae.Reference = _abf.MakeArray()
		for _, _cbgc := range _cbefgg.Elements() {
			_egda, _cagdd := _abf.GetDict(_cbgc)
			if !_cagdd {
				_acd.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020R\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0061\u0074\u0065\u0064")
				return nil, ErrInvalidAttribute
			}
			_eeggf, _gdcca := _ebdcaa.newPdfSignatureReferenceFromDict(_egda)
			if _gdcca != nil {
				return nil, _gdcca
			}
			_fedae.Reference.Append(_eeggf.ToPdfObject())
		}
	}
	_fedae.Cert = _bbdca.Get("\u0043\u0065\u0072\u0074")
	_fedae.ByteRange, _ = _abf.GetArray(_bbdca.Get("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e"))
	_fedae.Changes, _ = _abf.GetArray(_bbdca.Get("\u0043h\u0061\u006e\u0067\u0065\u0073"))
	_fedae.Name, _ = _abf.GetString(_bbdca.Get("\u004e\u0061\u006d\u0065"))
	_fedae.M, _ = _abf.GetString(_bbdca.Get("\u004d"))
	_fedae.Location, _ = _abf.GetString(_bbdca.Get("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"))
	_fedae.Reason, _ = _abf.GetString(_bbdca.Get("\u0052\u0065\u0061\u0073\u006f\u006e"))
	_fedae.ContactInfo, _ = _abf.GetString(_bbdca.Get("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f"))
	_fedae.R, _ = _abf.GetInt(_bbdca.Get("\u0052"))
	_fedae.V, _ = _abf.GetInt(_bbdca.Get("\u0056"))
	_fedae.PropBuild, _ = _abf.GetDict(_bbdca.Get("\u0050\u0072\u006f\u0070\u005f\u0042\u0075\u0069\u006c\u0064"))
	_fedae.PropAuthTime, _ = _abf.GetInt(_bbdca.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0069\u006d\u0065"))
	_fedae.PropAuthType, _ = _abf.GetName(_bbdca.Get("\u0050\u0072\u006f\u0070\u005f\u0041\u0075\u0074\u0068\u0054\u0079\u0070\u0065"))
	_ebdcaa._ceecd.Register(_afdfg, _fedae)
	return _fedae, nil
}

// GetIndirectObjectByNumber retrieves and returns a specific PdfObject by object number.
func (_edafc *PdfReader) GetIndirectObjectByNumber(number int) (_abf.PdfObject, error) {
	_eebe, _abdd := _edafc._bebc.LookupByNumber(number)
	return _eebe, _abdd
}

// NewCompositePdfFontFromTTFFile loads a composite font from a TTF font file. Composite fonts can
// be used to represent unicode fonts which can have multi-byte character codes, representing a wide
// range of values. They are often used for symbolic languages, including Chinese, Japanese and Korean.
// It is represented by a Type0 Font with an underlying CIDFontType2 and an Identity-H encoding map.
// TODO: May be extended in the future to support a larger variety of CMaps and vertical fonts.
// NOTE: For simple fonts, use NewPdfFontFromTTFFile.
func NewCompositePdfFontFromTTFFile(filePath string) (*PdfFont, error) {
	_bdee, _aaefgb := _cf.Open(filePath)
	if _aaefgb != nil {
		_acd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006f\u0070\u0065\u006e\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0076", _aaefgb)
		return nil, _aaefgb
	}
	defer _bdee.Close()
	return NewCompositePdfFontFromTTF(_bdee)
}

// GetPatternByName gets the pattern specified by keyName. Returns nil if not existing.
// The bool flag indicated whether it was found or not.
func (_caafe *PdfPageResources) GetPatternByName(keyName _abf.PdfObjectName) (*PdfPattern, bool) {
	if _caafe.Pattern == nil {
		return nil, false
	}
	_ababf, _cbcbg := _abf.TraceToDirectObject(_caafe.Pattern).(*_abf.PdfObjectDictionary)
	if !_cbcbg {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0061\u0074t\u0065\u0072\u006e\u0020\u0065\u006e\u0074r\u0079\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _caafe.Pattern)
		return nil, false
	}
	if _dadeg := _ababf.Get(keyName); _dadeg != nil {
		_eegfbg, _abgda := _aagc(_dadeg)
		if _abgda != nil {
			_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020f\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u006c\u006fa\u0064\u0020\u0070\u0064\u0066\u0020\u0070\u0061\u0074\u0074\u0065\u0072n\u003a\u0020\u0025\u0076", _abgda)
			return nil, false
		}
		return _eegfbg, true
	}
	return nil, false
}

// ToPdfObject implements interface PdfModel.
func (_fcda *PdfBorderStyle) ToPdfObject() _abf.PdfObject {
	_beeg := _abf.MakeDict()
	if _fcda._gfcg != nil {
		if _cfcf, _ddag := _fcda._gfcg.(*_abf.PdfIndirectObject); _ddag {
			_cfcf.PdfObject = _beeg
		}
	}
	_beeg.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0042\u006f\u0072\u0064\u0065\u0072"))
	if _fcda.W != nil {
		_beeg.Set("\u0057", _abf.MakeFloat(*_fcda.W))
	}
	if _fcda.S != nil {
		_beeg.Set("\u0053", _abf.MakeName(_fcda.S.GetPdfName()))
	}
	if _fcda.D != nil {
		_beeg.Set("\u0044", _abf.MakeArrayFromIntegers(*_fcda.D))
	}
	if _fcda._gfcg != nil {
		return _fcda._gfcg
	}
	return _beeg
}

// GetOCProperties returns the optional content properties PdfObject.
func (_afdd *PdfReader) GetOCProperties() (_abf.PdfObject, error) {
	_cebf := _afdd._dagde
	_edbgg := _cebf.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	_edbgg = _abf.ResolveReference(_edbgg)
	if !_afdd._abgge {
		_ffcee := _afdd.traverseObjectData(_edbgg)
		if _ffcee != nil {
			return nil, _ffcee
		}
	}
	return _edbgg, nil
}

func (_cccaa *pdfFontType0) getFontDescriptor() *PdfFontDescriptor {
	if _cccaa._dcbaf == nil && _cccaa.DescendantFont != nil {
		return _cccaa.DescendantFont.FontDescriptor()
	}
	return _cccaa._dcbaf
}

func _efggg(_dgeed *_abf.PdfObjectArray) (float64, error) {
	_begcea, _aebff := _dgeed.ToFloat64Array()
	if _aebff != nil {
		_acd.Log.Debug("\u0042\u0061\u0064\u0020\u004d\u0061\u0074\u0074\u0065\u0020\u0061\u0072\u0072\u0061\u0079:\u0020m\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dgeed, _aebff)
	}
	switch len(_begcea) {
	case 1:
		return _begcea[0], nil
	case 3:
		_aegfa := PdfColorspaceDeviceRGB{}
		_fecd, _cbec := _aegfa.ColorFromFloats(_begcea)
		if _cbec != nil {
			return 0.0, _cbec
		}
		return _fecd.(*PdfColorDeviceRGB).ToGray().Val(), nil
	case 4:
		_cdfc := PdfColorspaceDeviceCMYK{}
		_agfgef, _cbcaa := _cdfc.ColorFromFloats(_begcea)
		if _cbcaa != nil {
			return 0.0, _cbcaa
		}
		_dabe, _cbcaa := _cdfc.ColorToRGB(_agfgef.(*PdfColorDeviceCMYK))
		if _cbcaa != nil {
			return 0.0, _cbcaa
		}
		return _dabe.(*PdfColorDeviceRGB).ToGray().Val(), nil
	}
	_aebff = _fd.New("\u0062a\u0064 \u004d\u0061\u0074\u0074\u0065\u0020\u0063\u006f\u006c\u006f\u0072")
	_acd.Log.Error("\u0074\u006f\u0047ra\u0079\u003a\u0020\u006d\u0061\u0074\u0074\u0065\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dgeed, _aebff)
	return 0.0, _aebff
}

func _faefbc(_dbgde *_abf.PdfObjectDictionary) (*PdfShadingType4, error) {
	_fafge := PdfShadingType4{}
	_eddg := _dbgde.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _eddg == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_dcega, _cgffbe := _eddg.(*_abf.PdfObjectInteger)
	if !_cgffbe {
		_acd.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _eddg)
		return nil, _abf.ErrTypeError
	}
	_fafge.BitsPerCoordinate = _dcega
	_eddg = _dbgde.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _eddg == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_dcega, _cgffbe = _eddg.(*_abf.PdfObjectInteger)
	if !_cgffbe {
		_acd.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _eddg)
		return nil, _abf.ErrTypeError
	}
	_fafge.BitsPerComponent = _dcega
	_eddg = _dbgde.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _eddg == nil {
		_acd.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_dcega, _cgffbe = _eddg.(*_abf.PdfObjectInteger)
	if !_cgffbe {
		_acd.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _eddg)
		return nil, _abf.ErrTypeError
	}
	_fafge.BitsPerComponent = _dcega
	_eddg = _dbgde.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _eddg == nil {
		_acd.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fdbgf, _cgffbe := _eddg.(*_abf.PdfObjectArray)
	if !_cgffbe {
		_acd.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _eddg)
		return nil, _abf.ErrTypeError
	}
	_fafge.Decode = _fdbgf
	_eddg = _dbgde.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _eddg == nil {
		_acd.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_fafge.Function = []PdfFunction{}
	if _acbfb, _dbgbe := _eddg.(*_abf.PdfObjectArray); _dbgbe {
		for _, _aagba := range _acbfb.Elements() {
			_dbec, _bbge := _ebedg(_aagba)
			if _bbge != nil {
				_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bbge)
				return nil, _bbge
			}
			_fafge.Function = append(_fafge.Function, _dbec)
		}
	} else {
		_adffa, _dfaae := _ebedg(_eddg)
		if _dfaae != nil {
			_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dfaae)
			return nil, _dfaae
		}
		_fafge.Function = append(_fafge.Function, _adffa)
	}
	return &_fafge, nil
}

// GetNumComponents returns the number of color components (1 for Separation).
func (_cebg *PdfColorspaceSpecialSeparation) GetNumComponents() int { return 1 }

// NewPdfInfoFromObject creates a new PdfInfo from the input core.PdfObject.
func NewPdfInfoFromObject(obj _abf.PdfObject) (*PdfInfo, error) {
	var _fcbae PdfInfo
	_gddf, _eeee := obj.(*_abf.PdfObjectDictionary)
	if !_eeee {
		return nil, _e.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", obj)
	}
	for _, _gabfe := range _gddf.Keys() {
		switch _gabfe {
		case "\u0054\u0069\u0074l\u0065":
			_fcbae.Title, _ = _abf.GetString(_gddf.Get("\u0054\u0069\u0074l\u0065"))
		case "\u0041\u0075\u0074\u0068\u006f\u0072":
			_fcbae.Author, _ = _abf.GetString(_gddf.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
		case "\u0053u\u0062\u006a\u0065\u0063\u0074":
			_fcbae.Subject, _ = _abf.GetString(_gddf.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
		case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
			_fcbae.Keywords, _ = _abf.GetString(_gddf.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
		case "\u0043r\u0065\u0061\u0074\u006f\u0072":
			_fcbae.Creator, _ = _abf.GetString(_gddf.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
		case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
			_fcbae.Producer, _ = _abf.GetString(_gddf.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
		case "\u0054r\u0061\u0070\u0070\u0065\u0064":
			_fcbae.Trapped, _ = _abf.GetName(_gddf.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
		case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
			if _fgaae, _cgca := _abf.GetString(_gddf.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _cgca && _fgaae.String() != "" {
				_acab, _gabca := NewPdfDate(_fgaae.String())
				if _gabca != nil {
					return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _gabca)
				}
				_fcbae.CreationDate = &_acab
			}
		case "\u004do\u0064\u0044\u0061\u0074\u0065":
			if _acdb, _dbdf := _abf.GetString(_gddf.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _dbdf && _acdb.String() != "" {
				_cgef, _fdd := NewPdfDate(_acdb.String())
				if _fdd != nil {
					return nil, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _fdd)
				}
				_fcbae.ModifiedDate = &_cgef
			}
		default:
			_gdfd, _ := _abf.GetString(_gddf.Get(_gabfe))
			if _fcbae._cbf == nil {
				_fcbae._cbf = _abf.MakeDict()
			}
			_fcbae._cbf.Set(_gabfe, _gdfd)
		}
	}
	return &_fcbae, nil
}

// SetPdfTitle sets the Title attribute of the output PDF.
func SetPdfTitle(title string) { _gaabd.Lock(); defer _gaabd.Unlock(); _eabe = title }

// ToPdfObject implements interface PdfModel.
func (_cba *PdfActionSetOCGState) ToPdfObject() _abf.PdfObject {
	_cba.PdfAction.ToPdfObject()
	_dda := _cba._egg
	_fge := _dda.PdfObject.(*_abf.PdfObjectDictionary)
	_fge.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeSetOCGState)))
	_fge.SetIfNotNil("\u0053\u0074\u0061t\u0065", _cba.State)
	_fge.SetIfNotNil("\u0050\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0052\u0042", _cba.PreserveRB)
	return _dda
}

func _efaa(_eeegd *_abf.PdfIndirectObject, _ffdg *_abf.PdfObjectDictionary) (*DSS, error) {
	if _eeegd == nil {
		_eeegd = _abf.MakeIndirectObject(nil)
	}
	_eeegd.PdfObject = _abf.MakeDict()
	_gegcd := map[string]*VRI{}
	if _gdca, _gebe := _abf.GetDict(_ffdg.Get("\u0056\u0052\u0049")); _gebe {
		for _, _cddga := range _gdca.Keys() {
			if _eacd, _gecfe := _abf.GetDict(_gdca.Get(_cddga)); _gecfe {
				_gegcd[_be.ToUpper(_cddga.String())] = _bgde(_eacd)
			}
		}
	}
	return &DSS{Certs: _gggfec(_ffdg.Get("\u0043\u0065\u0072t\u0073")), OCSPs: _gggfec(_ffdg.Get("\u004f\u0043\u0053P\u0073")), CRLs: _gggfec(_ffdg.Get("\u0043\u0052\u004c\u0073")), VRI: _gegcd, _gffg: _eeegd}, nil
}

// DecodeArray returns the range of color component values in DeviceGray colorspace.
func (_acbb *PdfColorspaceDeviceGray) DecodeArray() []float64 { return []float64{0, 1.0} }

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_eadg *PdfShadingType3) ToPdfObject() _abf.PdfObject {
	_eadg.PdfShading.ToPdfObject()
	_cbbfe, _fddb := _eadg.getShadingDict()
	if _fddb != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _eadg.Coords != nil {
		_cbbfe.Set("\u0043\u006f\u006f\u0072\u0064\u0073", _eadg.Coords)
	}
	if _eadg.Domain != nil {
		_cbbfe.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _eadg.Domain)
	}
	if _eadg.Function != nil {
		if len(_eadg.Function) == 1 {
			_cbbfe.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _eadg.Function[0].ToPdfObject())
		} else {
			_adgbc := _abf.MakeArray()
			for _, _dagfd := range _eadg.Function {
				_adgbc.Append(_dagfd.ToPdfObject())
			}
			_cbbfe.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _adgbc)
		}
	}
	if _eadg.Extend != nil {
		_cbbfe.Set("\u0045\u0078\u0074\u0065\u006e\u0064", _eadg.Extend)
	}
	return _eadg._eabcgc
}

// PdfShadingType4 is a Free-form Gouraud-shaded triangle mesh.
type PdfShadingType4 struct {
	*PdfShading
	BitsPerCoordinate *_abf.PdfObjectInteger
	BitsPerComponent  *_abf.PdfObjectInteger
	BitsPerFlag       *_abf.PdfObjectInteger
	Decode            *_abf.PdfObjectArray
	Function          []PdfFunction
}

// GetContainingPdfObject implements model.PdfModel interface.
func (_dgfeg *PdfOutputIntent) GetContainingPdfObject() _abf.PdfObject { return _dgfeg._dcfb }

// PdfBorderStyle represents a border style dictionary (12.5.4 Border Styles p. 394).
type PdfBorderStyle struct {
	W     *float64
	S     *BorderStyle
	D     *[]int
	_gfcg _abf.PdfObject
}

// PdfActionThread represents a thread action.
type PdfActionThread struct {
	*PdfAction
	F *PdfFilespec
	D _abf.PdfObject
	B _abf.PdfObject
}

// ColorToRGB converts a CMYK32 color to an RGB color.
func (_cbbe *PdfColorspaceDeviceCMYK) ColorToRGB(color PdfColor) (PdfColor, error) {
	_cad, _fagb := color.(*PdfColorDeviceCMYK)
	if !_fagb {
		_acd.Log.Debug("I\u006e\u0070\u0075\u0074\u0020\u0063o\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0064e\u0076\u0069\u0063e\u0020c\u006d\u0079\u006b")
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gfcae := _cad.C()
	_cbgg := _cad.M()
	_fbae := _cad.Y()
	_cdfbg := _cad.K()
	_gfcae = _gfcae*(1-_cdfbg) + _cdfbg
	_cbgg = _cbgg*(1-_cdfbg) + _cdfbg
	_fbae = _fbae*(1-_cdfbg) + _cdfbg
	_agcc := 1 - _gfcae
	_dffc := 1 - _cbgg
	_fbce := 1 - _fbae
	return NewPdfColorDeviceRGB(_agcc, _dffc, _fbce), nil
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// A, B and C components of the color. The values of the elements should be
// between 0 and 1.
func (_dgba *PdfColorspaceCalRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_gabg := vals[0]
	if _gabg < 0.0 || _gabg > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _gabg)
		return nil, ErrColorOutOfRange
	}
	_aeaf := vals[1]
	if _aeaf < 0.0 || _aeaf > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _aeaf)
		return nil, ErrColorOutOfRange
	}
	_cgbf := vals[2]
	if _cgbf < 0.0 || _cgbf > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _cgbf)
		return nil, ErrColorOutOfRange
	}
	_befg := NewPdfColorCalRGB(_gabg, _aeaf, _cgbf)
	return _befg, nil
}

func _ceeabe(_dfcc string) map[string]string {
	_fcdba := _becf.Split(_dfcc, -1)
	_gegf := map[string]string{}
	for _, _fbfbe := range _fcdba {
		_bacgc := _geaa.FindStringSubmatch(_fbfbe)
		if _bacgc == nil {
			continue
		}
		_fegg, _febgd := _bacgc[1], _bacgc[2]
		_gegf[_fegg] = _febgd
	}
	return _gegf
}

type pdfFont interface {
	_gbe.Font

	// ToPdfObject returns a PDF representation of the font and implements interface Model.
	ToPdfObject() _abf.PdfObject
	getFontDescriptor() *PdfFontDescriptor
	baseFields() *fontCommon
}

func _gbbcga(_gfefdg _abf.PdfObject) (*fontFile, error) {
	_acd.Log.Trace("\u006e\u0065\u0077\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0046\u0072\u006f\u006dP\u0064f\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u006f\u0062\u006a\u003d\u0025\u0073", _gfefdg)
	_gfdgf := &fontFile{}
	_gfefdg = _abf.TraceToDirectObject(_gfefdg)
	_fbfc, _dbee := _gfefdg.(*_abf.PdfObjectStream)
	if !_dbee {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020F\u006f\u006et\u0046\u0069\u006c\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0028\u0025\u0054\u0029", _gfefdg)
		return nil, _abf.ErrTypeError
	}
	_cbbef := _fbfc.PdfObjectDictionary
	_gbegd, _ggeg := _abf.DecodeStream(_fbfc)
	if _ggeg != nil {
		return nil, _ggeg
	}
	_ageef, _dbee := _abf.GetNameVal(_cbbef.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_dbee {
		_gfdgf._eadac = _ageef
		if _ageef == "\u0054\u0079\u0070\u0065\u0031\u0043" {
			_acd.Log.Debug("T\u0079\u0070\u0065\u0031\u0043\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0063\u0075r\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u006e\u006f\u0074 s\u0075\u0070\u0070o\u0072t\u0065\u0064")
			return nil, ErrType1CFontNotSupported
		}
	}
	_gfeaf, _ := _abf.GetIntVal(_cbbef.Get("\u004ce\u006e\u0067\u0074\u0068\u0031"))
	_gcaad, _ := _abf.GetIntVal(_cbbef.Get("\u004ce\u006e\u0067\u0074\u0068\u0032"))
	if _gfeaf > len(_gbegd) {
		_gfeaf = len(_gbegd)
	}
	if _gfeaf+_gcaad > len(_gbegd) {
		_gcaad = len(_gbegd) - _gfeaf
	}
	_ffafb := _gbegd[:_gfeaf]
	var _fdde []byte
	if _gcaad > 0 {
		_fdde = _gbegd[_gfeaf : _gfeaf+_gcaad]
	}
	if _gfeaf > 0 && _gcaad > 0 {
		_bedafc := _gfdgf.loadFromSegments(_ffafb, _fdde)
		if _bedafc != nil {
			return nil, _bedafc
		}
	}
	return _gfdgf, nil
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_efad *PdfShadingPattern) ToPdfObject() _abf.PdfObject {
	_efad.PdfPattern.ToPdfObject()
	_daeef := _efad.getDict()
	if _efad.Shading != nil {
		_daeef.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _efad.Shading.ToPdfObject())
	}
	if _efad.Matrix != nil {
		_daeef.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _efad.Matrix)
	}
	if _efad.ExtGState != nil {
		_daeef.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _efad.ExtGState)
	}
	return _efad._bcfca
}

// GetPageLabels returns the PageLabels entry in the PDF catalog.
// See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
func (_cfccc *PdfReader) GetPageLabels() (_abf.PdfObject, error) {
	_bbddc := _abf.ResolveReference(_cfccc._dagde.Get("\u0050\u0061\u0067\u0065\u004c\u0061\u0062\u0065\u006c\u0073"))
	if _bbddc == nil {
		return nil, nil
	}
	if !_cfccc._abgge {
		_bfebb := _cfccc.traverseObjectData(_bbddc)
		if _bfebb != nil {
			return nil, _bfebb
		}
	}
	return _bbddc, nil
}

// ToPdfObject implements interface PdfModel.
func (_fec *PdfActionTrans) ToPdfObject() _abf.PdfObject {
	_fec.PdfAction.ToPdfObject()
	_dca := _fec._egg
	_faf := _dca.PdfObject.(*_abf.PdfObjectDictionary)
	_faf.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeTrans)))
	_faf.SetIfNotNil("\u0054\u0072\u0061n\u0073", _fec.Trans)
	return _dca
}

// PdfActionGoTo represents a GoTo action.
type PdfActionGoTo struct {
	*PdfAction
	D _abf.PdfObject
}

// AppendContentStream adds content stream by string.  Appends to the last
// contentstream instance if many.
func (_dadgf *PdfPage) AppendContentStream(contentStr string) error {
	_bcgda, _fafaa := _dadgf.GetContentStreams()
	if _fafaa != nil {
		return _fafaa
	}
	if len(_bcgda) == 0 {
		_bcgda = []string{contentStr}
		return _dadgf.SetContentStreams(_bcgda, _abf.NewFlateEncoder())
	}
	var _ecefe _dd.Buffer
	_ecefe.WriteString(_bcgda[len(_bcgda)-1])
	_ecefe.WriteString("\u000a")
	_ecefe.WriteString(contentStr)
	_bcgda[len(_bcgda)-1] = _ecefe.String()
	return _dadgf.SetContentStreams(_bcgda, _abf.NewFlateEncoder())
}

func (_fgda *PdfReader) newPdfActionNamedFromDict(_ebcb *_abf.PdfObjectDictionary) (*PdfActionNamed, error) {
	return &PdfActionNamed{N: _ebcb.Get("\u004e")}, nil
}

// GetColorspaces loads PdfPageResourcesColorspaces from `r.ColorSpace` and returns an error if there
// is a problem loading. Once loaded, the same object is returned on multiple calls.
func (_fbbfc *PdfPageResources) GetColorspaces() (*PdfPageResourcesColorspaces, error) {
	if _fbbfc._aafff != nil {
		return _fbbfc._aafff, nil
	}
	if _fbbfc.ColorSpace == nil {
		return nil, nil
	}
	_dagce, _ecfeg := _gebbf(_fbbfc.ColorSpace)
	if _ecfeg != nil {
		return nil, _ecfeg
	}
	_fbbfc._aafff = _dagce
	return _fbbfc._aafff, nil
}

// GetNameDictionary returns the Names entry in the PDF catalog.
// See section 7.7.4 "Name Dictionary" (p. 80 PDF32000_2008).
func (_gffa *PdfReader) GetNameDictionary() (_abf.PdfObject, error) {
	_cbcfg := _abf.ResolveReference(_gffa._dagde.Get("\u004e\u0061\u006de\u0073"))
	if _cbcfg == nil {
		return nil, nil
	}
	if !_gffa._abgge {
		_bcdcb := _gffa.traverseObjectData(_cbcfg)
		if _bcdcb != nil {
			return nil, _bcdcb
		}
	}
	return _cbcfg, nil
}

// ImageToRGB convert an indexed image to RGB.
func (_gfdda *PdfColorspaceSpecialIndexed) ImageToRGB(img Image) (Image, error) {
	N := _gfdda.Base.GetNumComponents()
	if N < 1 {
		return Image{}, _e.Errorf("\u0062\u0061d \u0062\u0061\u0073e\u0020\u0063\u006f\u006cors\u0070ac\u0065\u0020\u004e\u0075\u006d\u0043\u006fmp\u006f\u006e\u0065\u006e\u0074\u0073\u003d%\u0064", N)
	}
	_dfeb := _gca.NewImageBase(int(img.Width), int(img.Height), 8, N, nil, img._gedg, img._ceeag)
	_fgcc := _gf.NewReader(img.getBase())
	_beedgd := _gf.NewWriter(_dfeb)
	var (
		_aacf uint32
		_cbde int
		_egfc error
	)
	for {
		_aacf, _egfc = _fgcc.ReadSample()
		if _egfc == _gc.EOF {
			break
		} else if _egfc != nil {
			return img, _egfc
		}
		_cbde = int(_aacf)
		_acd.Log.Trace("\u0049\u006ed\u0065\u0078\u0065\u0064\u003a\u0020\u0069\u006e\u0064\u0065\u0078\u003d\u0025\u0064\u0020\u004e\u003d\u0025\u0064\u0020\u006c\u0075t=\u0025\u0064", _cbde, N, len(_gfdda._bcdf))
		if (_cbde+1)*N > len(_gfdda._bcdf) {
			_cbde = len(_gfdda._bcdf)/N - 1
			_acd.Log.Trace("C\u006c\u0069\u0070\u0070in\u0067 \u0074\u006f\u0020\u0069\u006ed\u0065\u0078\u003a\u0020\u0025\u0064", _cbde)
			if _cbde < 0 {
				_acd.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069p\u0020\u0069\u006e\u0064\u0065\u0078.\u0020\u0049\u0073\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006ce\u0020\u0064\u0061\u006d\u0061\u0067\u0065\u0064\u003f")
				break
			}
		}
		for _acdda := _cbde * N; _acdda < (_cbde+1)*N; _acdda++ {
			if _egfc = _beedgd.WriteSample(uint32(_gfdda._bcdf[_acdda])); _egfc != nil {
				return img, _egfc
			}
		}
	}
	return _gfdda.Base.ImageToRGB(_cega(&_dfeb))
}

func _gdcaf(_feba *_abf.PdfObjectDictionary) bool {
	for _, _cdda := range _feba.Keys() {
		if _, _abef := _dadge[_cdda.String()]; _abef {
			return true
		}
	}
	return false
}

// NewPdfAnnotationUnderline returns a new text underline annotation.
func NewPdfAnnotationUnderline() *PdfAnnotationUnderline {
	_fdf := NewPdfAnnotation()
	_dfc := &PdfAnnotationUnderline{}
	_dfc.PdfAnnotation = _fdf
	_dfc.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_fdf.SetContext(_dfc)
	return _dfc
}

// IsRadio returns true if the button field represents a radio button, false otherwise.
func (_ageaca *PdfFieldButton) IsRadio() bool { return _ageaca.GetType() == ButtonTypeRadio }

// GetPdfVersion gets the version of the PDF used within this document.
func (_gggd *PdfWriter) GetPdfVersion() string { return _gggd.getPdfVersion() }

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_afcf pdfCIDFontType2) GetRuneMetrics(r rune) (_gbe.CharMetrics, bool) {
	_abecd, _agbgf := _afcf._dffcb[r]
	if !_agbgf {
		_bbgd, _cbcg := _abf.GetInt(_afcf.DW)
		if !_cbcg {
			return _gbe.CharMetrics{}, false
		}
		_abecd = int(*_bbgd)
	}
	return _gbe.CharMetrics{Wx: float64(_abecd)}, true
}

func (_efgef *PdfReader) loadDSS() (*DSS, error) {
	if _efgef._bebc.GetCrypter() != nil && !_efgef._bebc.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006ce\u0020\u006e\u0065\u0065d\u0020\u0074\u006f\u0020\u0062\u0065\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	_cgdfb := _efgef._dagde.Get("\u0044\u0053\u0053")
	if _cgdfb == nil {
		return nil, nil
	}
	_ccgfa, _ := _abf.GetIndirect(_cgdfb)
	_cgdfb = _abf.TraceToDirectObject(_cgdfb)
	switch _dgde := _cgdfb.(type) {
	case *_abf.PdfObjectNull:
		return nil, nil
	case *_abf.PdfObjectDictionary:
		return _efaa(_ccgfa, _dgde)
	}
	return nil, _e.Errorf("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0044\u0053\u0053 \u0065\u006e\u0074\u0072y \u0025\u0054", _cgdfb)
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// red, green and blue components of the color. The values of the elements
// should be between 0 and 1.
func (_afedff *PdfColorspaceDeviceRGB) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bafcgd := vals[0]
	if _bafcgd < 0.0 || _bafcgd > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bafcgd)
		return nil, ErrColorOutOfRange
	}
	_eabc := vals[1]
	if _eabc < 0.0 || _eabc > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bafcgd)
		return nil, ErrColorOutOfRange
	}
	_dcbf := vals[2]
	if _dcbf < 0.0 || _dcbf > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bafcgd)
		return nil, ErrColorOutOfRange
	}
	_dadg := NewPdfColorDeviceRGB(_bafcgd, _eabc, _dcbf)
	return _dadg, nil
}

// PdfShadingPattern is a Shading patterns that provide a smooth transition between colors across an area to be painted,
// i.e. color(x,y) = f(x,y) at each point.
// It is a type 2 pattern (PatternType = 2).
type PdfShadingPattern struct {
	*PdfPattern
	Shading   *PdfShading
	Matrix    *_abf.PdfObjectArray
	ExtGState _abf.PdfObject
}

func _addec(_cgbfbe string) (string, error) {
	var _fcfe _dd.Buffer
	_fcfe.WriteString(_cgbfbe)
	_gbggg := make([]byte, 8+16)
	_egagcf := _f.Now().UTC().UnixNano()
	_bg.BigEndian.PutUint64(_gbggg, uint64(_egagcf))
	_, _gfeadg := _g.Read(_gbggg[8:])
	if _gfeadg != nil {
		return "", _gfeadg
	}
	_fcfe.WriteString(_cb.EncodeToString(_gbggg))
	return _fcfe.String(), nil
}

var _gffad = false

// ColorToRGB converts a DeviceN color to an RGB color.
func (_cfbag *PdfColorspaceDeviceN) ColorToRGB(color PdfColor) (PdfColor, error) {
	if _cfbag.AlternateSpace == nil {
		return nil, _fd.New("\u0044\u0065\u0076\u0069\u0063\u0065N\u0020\u0061\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070a\u0063\u0065\u0020\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	return _cfbag.AlternateSpace.ColorToRGB(color)
}

// PdfActionMovie represents a movie action.
type PdfActionMovie struct {
	*PdfAction
	Annotation _abf.PdfObject
	T          _abf.PdfObject
	Operation  _abf.PdfObject
}

func (_aabdd SignatureValidationResult) String() string {
	var _bdefe _dd.Buffer
	_bdefe.WriteString(_e.Sprintf("\u004ea\u006d\u0065\u003a\u0020\u0025\u0073\n", _aabdd.Name))
	if _aabdd.Date._fabd > 0 {
		_bdefe.WriteString(_e.Sprintf("\u0044a\u0074\u0065\u003a\u0020\u0025\u0073\n", _aabdd.Date.ToGoTime().String()))
	} else {
		_bdefe.WriteString("\u0044\u0061\u0074\u0065 n\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_aabdd.Reason) > 0 {
		_bdefe.WriteString(_e.Sprintf("R\u0065\u0061\u0073\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _aabdd.Reason))
	} else {
		_bdefe.WriteString("N\u006f \u0072\u0065\u0061\u0073\u006f\u006e\u0020\u0073p\u0065\u0063\u0069\u0066ie\u0064\u000a")
	}
	if len(_aabdd.Location) > 0 {
		_bdefe.WriteString(_e.Sprintf("\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u000a", _aabdd.Location))
	} else {
		_bdefe.WriteString("\u004c\u006f\u0063at\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u000a")
	}
	if len(_aabdd.ContactInfo) > 0 {
		_bdefe.WriteString(_e.Sprintf("\u0043\u006f\u006e\u0074\u0061\u0063\u0074\u0020\u0049\u006e\u0066\u006f:\u0020\u0025\u0073\u000a", _aabdd.ContactInfo))
	} else {
		_bdefe.WriteString("C\u006f\u006e\u0074\u0061\u0063\u0074 \u0069\u006e\u0066\u006f\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063i\u0066i\u0065\u0064\u000a")
	}
	_bdefe.WriteString(_e.Sprintf("F\u0069\u0065\u006c\u0064\u0073\u003a\u0020\u0025\u0064\u000a", len(_aabdd.Fields)))
	if _aabdd.IsSigned {
		_bdefe.WriteString("S\u0069\u0067\u006e\u0065\u0064\u003a \u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073\u0020\u0073i\u0067n\u0065\u0064\u000a")
	} else {
		_bdefe.WriteString("\u0053\u0069\u0067\u006eed\u003a\u0020\u004e\u006f\u0074\u0020\u0073\u0069\u0067\u006e\u0065\u0064\u000a")
	}
	if _aabdd.IsVerified {
		_bdefe.WriteString("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072e\u0020\u0076\u0061\u006c\u0069\u0064\u0061t\u0069\u006f\u006e\u003a\u0020\u0049\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u000a")
	} else {
		_bdefe.WriteString("\u0053\u0069\u0067\u006e\u0061\u0074u\u0072\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e:\u0020\u0049\u0073\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u000a")
	}
	if _aabdd.IsTrusted {
		_bdefe.WriteString("\u0054\u0072\u0075\u0073\u0074\u0065\u0064\u003a\u0020\u0043\u0065\u0072\u0074\u0069\u0066i\u0063a\u0074\u0065\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0073\u0074\u0065\u0064\u000a")
	} else {
		_bdefe.WriteString("\u0054\u0072\u0075s\u0074\u0065\u0064\u003a \u0055\u006e\u0074\u0072\u0075\u0073\u0074e\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u000a")
	}
	if !_aabdd.GeneralizedTime.IsZero() {
		_bdefe.WriteString(_e.Sprintf("G\u0065n\u0065\u0072\u0061\u006c\u0069\u007a\u0065\u0064T\u0069\u006d\u0065\u003a %\u0073\u000a", _aabdd.GeneralizedTime.String()))
	}
	if _aabdd.DiffResults != nil {
		_bdefe.WriteString(_e.Sprintf("\u0064\u0069\u0066\u0066 i\u0073\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u003a\u0020\u0025v\u000a", _aabdd.DiffResults.IsPermitted()))
		if len(_aabdd.DiffResults.Warnings) > 0 {
			_bdefe.WriteString("\u004d\u0044\u0050\u0020\u0077\u0061\u0072\u006e\u0069n\u0067\u0073\u003a\u000a")
			for _, _efbbe := range _aabdd.DiffResults.Warnings {
				_bdefe.WriteString(_e.Sprintf("\u0009\u0025\u0073\u000a", _efbbe))
			}
		}
		if len(_aabdd.DiffResults.Errors) > 0 {
			_bdefe.WriteString("\u004d\u0044\u0050 \u0065\u0072\u0072\u006f\u0072\u0073\u003a\u000a")
			for _, _bdfgg := range _aabdd.DiffResults.Errors {
				_bdefe.WriteString(_e.Sprintf("\u0009\u0025\u0073\u000a", _bdfgg))
			}
		}
	}
	if _aabdd.IsCrlFound {
		_bdefe.WriteString("R\u0065\u0076\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0061\u0074\u0061\u003a\u0020C\u0052\u004c\u0020f\u006fu\u006e\u0064\u000a")
	} else {
		_bdefe.WriteString("\u0052\u0065\u0076o\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0061\u0074\u0061\u003a \u0043\u0052\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u000a")
	}
	if _aabdd.IsOcspFound {
		_bdefe.WriteString("\u0052\u0065\u0076\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0061\u0074\u0061:\u0020O\u0043\u0053\u0050\u0020\u0066\u006f\u0075n\u0064\u000a")
	} else {
		_bdefe.WriteString("\u0052\u0065\u0076\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0061\u0074\u0061:\u0020O\u0043\u0053\u0050\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u000a")
	}
	return _bdefe.String()
}

func (_egdge *PdfPage) setContainer(_aebfb *_abf.PdfIndirectObject) {
	_aebfb.PdfObject = _egdge._bdbfa
	_egdge._gefee = _aebfb
}

// NewPdfActionNamed returns a new "named" action.
func NewPdfActionNamed() *PdfActionNamed {
	_ce := NewPdfAction()
	_aae := &PdfActionNamed{}
	_aae.PdfAction = _ce
	_ce.SetContext(_aae)
	return _aae
}

// ToPdfObject implements interface PdfModel.
func (_db *PdfActionGoTo3DView) ToPdfObject() _abf.PdfObject {
	_db.PdfAction.ToPdfObject()
	_abfc := _db._egg
	_bcfa := _abfc.PdfObject.(*_abf.PdfObjectDictionary)
	_bcfa.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeGoTo3DView)))
	_bcfa.SetIfNotNil("\u0054\u0041", _db.TA)
	_bcfa.SetIfNotNil("\u0056", _db.V)
	return _abfc
}

// NewDSS returns a new DSS dictionary.
func NewDSS() *DSS {
	return &DSS{_gffg: _abf.MakeIndirectObject(_abf.MakeDict()), VRI: map[string]*VRI{}}
}

// GetContext returns a reference to the subshading entry as represented by PdfShadingType1-7.
func (_decc *PdfShading) GetContext() PdfModel { return _decc._eabd }

func (_gbde *PdfColorspaceDeviceGray) String() string {
	return "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_egfce *PdfShadingType7) ToPdfObject() _abf.PdfObject {
	_egfce.PdfShading.ToPdfObject()
	_bcae, _fcbgg := _egfce.getShadingDict()
	if _fcbgg != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _egfce.BitsPerCoordinate != nil {
		_bcae.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _egfce.BitsPerCoordinate)
	}
	if _egfce.BitsPerComponent != nil {
		_bcae.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _egfce.BitsPerComponent)
	}
	if _egfce.BitsPerFlag != nil {
		_bcae.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _egfce.BitsPerFlag)
	}
	if _egfce.Decode != nil {
		_bcae.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _egfce.Decode)
	}
	if _egfce.Function != nil {
		if len(_egfce.Function) == 1 {
			_bcae.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _egfce.Function[0].ToPdfObject())
		} else {
			_abfbg := _abf.MakeArray()
			for _, _edeeee := range _egfce.Function {
				_abfbg.Append(_edeeee.ToPdfObject())
			}
			_bcae.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _abfbg)
		}
	}
	return _egfce._eabcgc
}

// ToPdfObject returns the PDF representation of the function.
func (_deab *PdfFunctionType0) ToPdfObject() _abf.PdfObject {
	if _deab._cabaa == nil {
		_deab._cabaa = &_abf.PdfObjectStream{}
	}
	_dbadc := _abf.MakeDict()
	_dbadc.Set("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065", _abf.MakeInteger(0))
	_gccb := &_abf.PdfObjectArray{}
	for _, _daabd := range _deab.Domain {
		_gccb.Append(_abf.MakeFloat(_daabd))
	}
	_dbadc.Set("\u0044\u006f\u006d\u0061\u0069\u006e", _gccb)
	_bcff := &_abf.PdfObjectArray{}
	for _, _baga := range _deab.Range {
		_bcff.Append(_abf.MakeFloat(_baga))
	}
	_dbadc.Set("\u0052\u0061\u006eg\u0065", _bcff)
	_ffcc := &_abf.PdfObjectArray{}
	for _, _cedcde := range _deab.Size {
		_ffcc.Append(_abf.MakeInteger(int64(_cedcde)))
	}
	_dbadc.Set("\u0053\u0069\u007a\u0065", _ffcc)
	_dbadc.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0053\u0061\u006d\u0070\u006c\u0065", _abf.MakeInteger(int64(_deab.BitsPerSample)))
	if _deab.Order != 1 {
		_dbadc.Set("\u004f\u0072\u0064e\u0072", _abf.MakeInteger(int64(_deab.Order)))
	}
	_dbadc.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _abf.MakeInteger(int64(len(_deab._aefbg))))
	_deab._cabaa.Stream = _deab._aefbg
	_deab._cabaa.PdfObjectDictionary = _dbadc
	return _deab._cabaa
}

func (_adgb fontCommon) isCIDFont() bool {
	if _adgb._aacbc == "" {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0069\u0073\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u002e\u0020\u0063o\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u002e\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _adgb)
	}
	_cdac := false
	switch _adgb._aacbc {
	case "\u0054\u0079\u0070e\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_cdac = true
	}
	_acd.Log.Trace("i\u0073\u0043\u0049\u0044\u0046\u006fn\u0074\u003a\u0020\u0069\u0073\u0043\u0049\u0044\u003d%\u0074\u0020\u0066o\u006et\u003d\u0025\u0073", _cdac, _adgb)
	return _cdac
}

func (_bcefd *PdfColorspaceSpecialSeparation) String() string {
	return "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e"
}

func _adde(_bdd _abf.PdfObject) (*PdfBorderStyle, error) {
	_agac := &PdfBorderStyle{}
	_agac._gfcg = _bdd
	var _ebcc *_abf.PdfObjectDictionary
	_bdd = _abf.TraceToDirectObject(_bdd)
	_ebcc, _ddcd := _bdd.(*_abf.PdfObjectDictionary)
	if !_ddcd {
		return nil, _fd.New("\u0074\u0079\u0070\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	if _dgee := _ebcc.Get("\u0054\u0079\u0070\u0065"); _dgee != nil {
		_edcd, _dddc := _dgee.(*_abf.PdfObjectName)
		if !_dddc {
			_acd.Log.Debug("I\u006e\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u0077\u0069th\u0020\u0054\u0079\u0070e\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061me\u0020\u006fb\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0054", _dgee)
		} else {
			if *_edcd != "\u0042\u006f\u0072\u0064\u0065\u0072" {
				_acd.Log.Debug("W\u0061\u0072\u006e\u0069\u006e\u0067,\u0020\u0054\u0079\u0070\u0065\u0020\u0021\u003d\u0020B\u006f\u0072\u0064e\u0072:\u0020\u0025\u0073", *_edcd)
			}
		}
	}
	if _bcdb := _ebcc.Get("\u0057"); _bcdb != nil {
		_fbb, _ceg := _abf.GetNumberAsFloat(_bcdb)
		if _ceg != nil {
			_acd.Log.Debug("\u0045\u0072\u0072\u006fr \u0072\u0065\u0074\u0072\u0069\u0065\u0076\u0069\u006e\u0067\u0020\u0057\u003a\u0020%\u0076", _ceg)
			return nil, _ceg
		}
		_agac.W = &_fbb
	}
	if _cac := _ebcc.Get("\u0053"); _cac != nil {
		_afed, _acee := _cac.(*_abf.PdfObjectName)
		if !_acee {
			return nil, _fd.New("\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u0053\u0020\u006e\u006ft\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0062j\u0065\u0063\u0074")
		}
		var _aaffb BorderStyle
		switch *_afed {
		case "\u0053":
			_aaffb = BorderStyleSolid
		case "\u0044":
			_aaffb = BorderStyleDashed
		case "\u0042":
			_aaffb = BorderStyleBeveled
		case "\u0049":
			_aaffb = BorderStyleInset
		case "\u0055":
			_aaffb = BorderStyleUnderline
		default:
			_acd.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020s\u0074\u0079\u006c\u0065\u0020n\u0061\u006d\u0065\u0020\u0025\u0073", *_afed)
			return nil, _fd.New("\u0073\u0074\u0079\u006ce \u0074\u0079\u0070\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065c\u006b")
		}
		_agac.S = &_aaffb
	}
	if _abee := _ebcc.Get("\u0044"); _abee != nil {
		_eeeg, _bba := _abee.(*_abf.PdfObjectArray)
		if !_bba {
			_acd.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044\u0020\u0064a\u0073\u0068\u0020\u006e\u006f\u0074\u0020a\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0054", _abee)
			return nil, _fd.New("\u0062o\u0072\u0064\u0065\u0072 \u0044\u0020\u0074\u0079\u0070e\u0020c\u0068e\u0063\u006b\u0020\u0065\u0072\u0072\u006fr")
		}
		_bcfd, _gabd := _eeeg.ToIntegerArray()
		if _gabd != nil {
			_acd.Log.Debug("\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u0044 \u0050\u0072\u006fbl\u0065\u006d\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u006e\u0074\u0065\u0067e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u003a \u0025\u0076", _gabd)
			return nil, _gabd
		}
		_agac.D = &_bcfd
	}
	return _agac, nil
}

func (_ddfc *PdfPage) flattenFieldsWithOpts(_gbbe FieldAppearanceGenerator, _abeec *FieldFlattenOpts, _gaaf map[*PdfAnnotation]bool) error {
	var _egdd []*PdfAnnotation
	if _gbbe != nil {
		if _daef := _gbbe.WrapContentStream(_ddfc); _daef != nil {
			return _daef
		}
	}
	_gadge, _gbcc := _ddfc.GetAnnotations()
	if _gbcc != nil {
		return _gbcc
	}
	for _, _daeg := range _gadge {
		_fgaag, _beebf := _gaaf[_daeg]
		if !_beebf && _abeec.AnnotFilterFunc != nil {
			if _, _gcgb := _daeg.GetContext().(*PdfAnnotationWidget); !_gcgb {
				_beebf = _abeec.AnnotFilterFunc(_daeg)
			}
		}
		if !_beebf {
			_egdd = append(_egdd, _daeg)
			continue
		}
		switch _daeg.GetContext().(type) {
		case *PdfAnnotationPopup:
			continue
		case *PdfAnnotationLink:
			continue
		case *PdfAnnotationProjection:
			continue
		}
		_edgg, _agcg, _bgfd := _gbdd(_daeg)
		if _bgfd != nil {
			if !_fgaag {
				_acd.Log.Trace("\u0046\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0056\u0020\u002d\u003e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075t\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065am\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072")
				continue
			}
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0077\u0069\u0074h\u006f\u0075\u0074\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d,\u0020\u0065\u0072\u0072\u0020\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006bi\u0070\u0070\u0069n\u0067\u0020\u006f\u0076\u0065\u0072", _bgfd)
			continue
		}
		if _edgg == nil {
			continue
		}
		_cggba := _ddfc.Resources.GenerateXObjectName()
		_ddfc.Resources.SetXObjectFormByName(_cggba, _edgg)
		_dfbd, _cefd, _bgfd := _dbde(_edgg)
		if _bgfd != nil {
			_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0061\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u004d\u0061\u0074\u0072\u0069\u0078\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0078\u0066\u006f\u0072\u006d\u0020\u0062\u0062\u006f\u0078\u0020\u0061\u0064\u006a\u0075\u0073t\u006d\u0065\u006e\u0074\u003a \u0025\u0076", _bgfd)
		} else {
			_fccb := _ad.IdentityMatrix()
			_fccb = _fccb.Translate(-_dfbd.Llx, -_dfbd.Lly)
			if _cefd {
				_dbag := 0.0
				if _dfbd.Width() > 0 {
					_dbag = _agcg.Width() / _dfbd.Width()
				}
				_efgf := 0.0
				if _dfbd.Height() > 0 {
					_efgf = _agcg.Height() / _dfbd.Height()
				}
				_fccb = _fccb.Scale(_dbag, _efgf)
			}
			_agcg.Transform(_fccb)
		}
		_eaae := _ge.Min(_agcg.Llx, _agcg.Urx)
		_dceg := _ge.Min(_agcg.Lly, _agcg.Ury)
		var _cecg []string
		_cecg = append(_cecg, "\u0071")
		_cecg = append(_cecg, _e.Sprintf("\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020%\u002e6\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0025\u002e\u0036\u0066\u0020\u0063\u006d", 1.0, 0.0, 0.0, 1.0, _eaae, _dceg))
		_cecg = append(_cecg, _e.Sprintf("\u002f\u0025\u0073\u0020\u0044\u006f", _cggba.String()))
		_cecg = append(_cecg, "\u0051")
		_acfc := _be.Join(_cecg, "\u000a")
		_bgfd = _ddfc.AppendContentStream(_acfc)
		if _bgfd != nil {
			return _bgfd
		}
		if _edgg.Resources != nil {
			_cefg, _eecbc := _abf.GetDict(_edgg.Resources.Font)
			if _eecbc {
				for _, _gbfgd := range _cefg.Keys() {
					if !_ddfc.Resources.HasFontByName(_gbfgd) {
						_ddfc.Resources.SetFontByName(_gbfgd, _cefg.Get(_gbfgd))
					}
				}
			}
		}
	}
	if len(_egdd) > 0 {
		_ddfc._baagf = _egdd
	} else {
		_ddfc._baagf = []*PdfAnnotation{}
	}
	return nil
}

// PdfColorspaceLab is a L*, a*, b* 3 component colorspace.
type PdfColorspaceLab struct {
	WhitePoint []float64
	BlackPoint []float64
	Range      []float64
	_aaec      *_abf.PdfIndirectObject
}

// GetCatalogMetadata gets the catalog defined XMP Metadata.
func (_bcgce *PdfReader) GetCatalogMetadata() (_abf.PdfObject, bool) {
	if _bcgce._dagde == nil {
		return nil, false
	}
	_afgac := _bcgce._dagde.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	return _afgac, _afgac != nil
}

// GetStandardApplier gets currently used StandardApplier..
func (_gbegdg *PdfWriter) GetStandardApplier() StandardApplier { return _gbegdg._adgdc }

// GetRuneMetrics returns the char metrics for a rune.
// TODO(peterwilliams97) There is nothing callers can do if no CharMetrics are found so we might as
// well give them 0 width. There is no need for the bool return.
func (_gbcaa *PdfFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_begb := _gbcaa.actualFont()
	if _begb == nil {
		_acd.Log.Debug("ER\u0052\u004fR\u003a\u0020\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061\u0072\u004d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u004e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020f\u006fr\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079p\u0065=\u0025\u0023T", _gbcaa._gedca)
		return _gbe.CharMetrics{}, false
	}
	if _acbf, _aabd := _begb.GetRuneMetrics(r); _aabd {
		return _acbf, true
	}
	if _bbef, _adec := _gbcaa.GetFontDescriptor(); _adec == nil && _bbef != nil {
		return _gbe.CharMetrics{Wx: _bbef._fgccc}, true
	}
	_acd.Log.Debug("\u0047\u0065\u0074\u0047\u006c\u0079\u0070h\u0043\u0068\u0061r\u004d\u0065\u0074\u0072i\u0063\u0073\u003a\u0020\u004e\u006f\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gbcaa)
	return _gbe.CharMetrics{}, false
}

func (_ggeae *PdfWriter) checkCrossReferenceStream() bool {
	_abedcg := _ggeae._ecfa.Major > 1 || (_ggeae._ecfa.Major == 1 && _ggeae._ecfa.Minor > 4)
	if _ggeae._adceg != nil {
		_abedcg = *_ggeae._adceg
	}
	return _abedcg
}

// UpdateObject marks `obj` as updated and to be included in the following revision.
func (_dadc *PdfAppender) UpdateObject(obj _abf.PdfObject) {
	_dadc.replaceObject(obj, obj)
	if _, _eece := _dadc._gcba[obj]; !_eece {
		_dadc._ffcf = append(_dadc._ffcf, obj)
		_dadc._gcba[obj] = struct{}{}
	}
}

// ColorFromFloats returns a new PdfColorDevice based on the input slice of
// color components. The slice should contain four elements representing the
// cyan, magenta, yellow and key components of the color. The values of the
// elements should be between 0 and 1.
func (_ddcf *PdfColorspaceDeviceCMYK) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 4 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bbbe := vals[0]
	if _bbbe < 0.0 || _bbbe > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _bbbe)
		return nil, ErrColorOutOfRange
	}
	_gbbc := vals[1]
	if _gbbc < 0.0 || _gbbc > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _gbbc)
		return nil, ErrColorOutOfRange
	}
	_dbbf := vals[2]
	if _dbbf < 0.0 || _dbbf > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _dbbf)
		return nil, ErrColorOutOfRange
	}
	_ecea := vals[3]
	if _ecea < 0.0 || _ecea > 1.0 {
		_acd.Log.Debug("\u0063\u006f\u006cor\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0043\u0053\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020", _ecea)
		return nil, ErrColorOutOfRange
	}
	_gege := NewPdfColorDeviceCMYK(_bbbe, _gbbc, _dbbf, _ecea)
	return _gege, nil
}

// SetPdfProducer sets the Producer attribute of the output PDF.
func SetPdfProducer(producer string) { _gaabd.Lock(); defer _gaabd.Unlock(); _babfc = producer }

// PdfColorLab represents a color in the L*, a*, b* 3 component colorspace.
// Each component is defined in the range 0.0 - 1.0 where 1.0 is the primary intensity.
type PdfColorLab [3]float64

// ToPdfObject returns a PDF object representation of the outline item.
func (_ggead *OutlineItem) ToPdfObject() _abf.PdfObject {
	_geceg, _ := _ggead.ToPdfOutlineItem()
	return _geceg.ToPdfObject()
}

// NewOutlineBookmark returns an initialized PdfOutlineItem for a given bookmark title and page.
func NewOutlineBookmark(title string, page *_abf.PdfIndirectObject) *PdfOutlineItem {
	_debg := PdfOutlineItem{}
	_debg._aecec = &_debg
	_debg.Title = _abf.MakeString(title)
	_faaag := _abf.MakeArray()
	_faaag.Append(page)
	_faaag.Append(_abf.MakeName("\u0046\u0069\u0074"))
	_debg.Dest = _faaag
	return &_debg
}

// GetMediaBox gets the inheritable media box value, either from the page
// or a higher up page/pages struct.
func (_cagbaa *PdfPage) GetMediaBox() (*PdfRectangle, error) {
	if _cagbaa.MediaBox != nil {
		return _cagbaa.MediaBox, nil
	}
	_bceg := _cagbaa.Parent
	for _bceg != nil {
		_beba, _cfgfa := _abf.GetDict(_bceg)
		if !_cfgfa {
			return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		if _gfade := _beba.Get("\u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078"); _gfade != nil {
			_gacac, _bccgf := _abf.GetArray(_gfade)
			if !_bccgf {
				return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0064\u0069a\u0020\u0062\u006f\u0078")
			}
			_ggeac, _efcc := NewPdfRectangle(*_gacac)
			if _efcc != nil {
				return nil, _efcc
			}
			return _ggeac, nil
		}
		_bceg = _beba.Get("\u0050\u0061\u0072\u0065\u006e\u0074")
	}
	return nil, _fd.New("m\u0065\u0064\u0069\u0061 b\u006fx\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
}
func _gbbbf(_efac *fontCommon) *pdfCIDFontType2 { return &pdfCIDFontType2{fontCommon: *_efac} }

// PdfActionLaunch represents a launch action.
type PdfActionLaunch struct {
	*PdfAction
	F         *PdfFilespec
	Win       _abf.PdfObject
	Mac       _abf.PdfObject
	Unix      _abf.PdfObject
	NewWindow _abf.PdfObject
}

// GetBorderWidth returns the border style's width.
func (_febg *PdfBorderStyle) GetBorderWidth() float64 {
	if _febg.W == nil {
		return 1
	}
	return *_febg.W
}

func (_cgbfb *DSS) add(_fcga *[]*_abf.PdfObjectStream, _bdbd map[string]*_abf.PdfObjectStream, _adbae [][]byte) ([]*_abf.PdfObjectStream, error) {
	_fcff := make([]*_abf.PdfObjectStream, 0, len(_adbae))
	for _, _cacfc := range _adbae {
		_aede, _caec := _fdbbe(_cacfc)
		if _caec != nil {
			return nil, _caec
		}
		_gcdg, _bgeg := _bdbd[string(_aede)]
		if !_bgeg {
			_gcdg, _caec = _abf.MakeStream(_cacfc, _abf.NewRawEncoder())
			if _caec != nil {
				return nil, _caec
			}
			_bdbd[string(_aede)] = _gcdg
			*_fcga = append(*_fcga, _gcdg)
		}
		_fcff = append(_fcff, _gcdg)
	}
	return _fcff, nil
}

// SetDate sets the `M` field of the signature.
func (_cbfb *PdfSignature) SetDate(date _f.Time, format string) {
	if format == "" {
		format = "\u0044\u003a\u003200\u0036\u0030\u0031\u0030\u0032\u0031\u0035\u0030\u0034\u0030\u0035\u002d\u0030\u0037\u0027\u0030\u0030\u0027"
	}
	_cbfb.M = _abf.MakeString(date.Format(format))
}

func _fecfd(_fgdfc *_abf.PdfObjectDictionary) (*PdfShadingType3, error) {
	_abgac := PdfShadingType3{}
	_agba := _fgdfc.Get("\u0043\u006f\u006f\u0072\u0064\u0073")
	if _agba == nil {
		_acd.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0043\u006f\u006f\u0072\u0064\u0073")
		return nil, ErrRequiredAttributeMissing
	}
	_dbcbf, _ccedf := _agba.(*_abf.PdfObjectArray)
	if !_ccedf {
		_acd.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _agba)
		return nil, _abf.ErrTypeError
	}
	if _dbcbf.Len() != 6 {
		_acd.Log.Debug("\u0043\u006f\u006f\u0072d\u0073\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0036\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _dbcbf.Len())
		return nil, ErrInvalidAttribute
	}
	_abgac.Coords = _dbcbf
	if _ccac := _fgdfc.Get("\u0044\u006f\u006d\u0061\u0069\u006e"); _ccac != nil {
		_ccac = _abf.TraceToDirectObject(_ccac)
		_cgebb, _fabac := _ccac.(*_abf.PdfObjectArray)
		if !_fabac {
			_acd.Log.Debug("\u0044\u006f\u006d\u0061i\u006e\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _ccac)
			return nil, _abf.ErrTypeError
		}
		_abgac.Domain = _cgebb
	}
	_agba = _fgdfc.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	if _agba == nil {
		_acd.Log.Debug("\u0052\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0020\u0046\u0075\u006ec\u0074\u0069\u006f\u006e")
		return nil, ErrRequiredAttributeMissing
	}
	_abgac.Function = []PdfFunction{}
	if _cdcbc, _gafcd := _agba.(*_abf.PdfObjectArray); _gafcd {
		for _, _bggbd := range _cdcbc.Elements() {
			_agdf, _bcgbg := _ebedg(_bggbd)
			if _bcgbg != nil {
				_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bcgbg)
				return nil, _bcgbg
			}
			_abgac.Function = append(_abgac.Function, _agdf)
		}
	} else {
		_ebggec, _ccced := _ebedg(_agba)
		if _ccced != nil {
			_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _ccced)
			return nil, _ccced
		}
		_abgac.Function = append(_abgac.Function, _ebggec)
	}
	if _deddg := _fgdfc.Get("\u0045\u0078\u0074\u0065\u006e\u0064"); _deddg != nil {
		_deddg = _abf.TraceToDirectObject(_deddg)
		_faegd, _ccadd := _deddg.(*_abf.PdfObjectArray)
		if !_ccadd {
			_acd.Log.Debug("\u004d\u0061\u0074\u0072i\u0078\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _deddg)
			return nil, _abf.ErrTypeError
		}
		if _faegd.Len() != 2 {
			_acd.Log.Debug("\u0045\u0078\u0074\u0065n\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074 \u0032\u0020\u0028\u0067\u006f\u0074\u0020%\u0064\u0029", _faegd.Len())
			return nil, ErrInvalidAttribute
		}
		_abgac.Extend = _faegd
	}
	return &_abgac, nil
}

// SetFilter sets compression filter. Decodes with current filter sets and
// encodes the data with the new filter.
func (_bgdfe *XObjectImage) SetFilter(encoder _abf.StreamEncoder) error {
	_ebgd := _bgdfe.Stream
	_dbged, _dbdff := _bgdfe.Filter.DecodeBytes(_ebgd)
	if _dbdff != nil {
		return _dbdff
	}
	_bgdfe.Filter = encoder
	encoder.UpdateParams(_bgdfe.getParamsDict())
	_ebgd, _dbdff = encoder.EncodeBytes(_dbged)
	if _dbdff != nil {
		return _dbdff
	}
	_bgdfe.Stream = _ebgd
	return nil
}

func (_aace *PdfReader) buildPageList(_cgfee *_abf.PdfIndirectObject, _fbafgc *_abf.PdfIndirectObject, _feaa map[_abf.PdfObject]struct{}) error {
	if _cgfee == nil {
		return nil
	}
	if _, _ceffb := _feaa[_cgfee]; _ceffb {
		_acd.Log.Debug("\u0043\u0079\u0063l\u0069\u0063\u0020\u0072e\u0063\u0075\u0072\u0073\u0069\u006f\u006e,\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0028\u0025\u0076\u0029", _cgfee.ObjectNumber)
		return nil
	}
	_feaa[_cgfee] = struct{}{}
	_fddgd, _dfba := _cgfee.PdfObject.(*_abf.PdfObjectDictionary)
	if !_dfba {
		return _fd.New("n\u006f\u0064\u0065\u0020no\u0074 \u0061\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_adgee, _dfba := (*_fddgd).Get("\u0054\u0079\u0070\u0065").(*_abf.PdfObjectName)
	if !_dfba {
		if _fddgd.Get("\u004b\u0069\u0064\u0073") == nil {
			return _fd.New("\u006e\u006f\u0064\u0065 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0054\u0079p\u0065 \u0028\u0052\u0065\u0071\u0075\u0069\u0072e\u0064\u0029")
		}
		_acd.Log.Debug("ER\u0052\u004fR\u003a\u0020\u006e\u006f\u0064\u0065\u0020\u006d\u0069s\u0073\u0069\u006e\u0067\u0020\u0054\u0079\u0070\u0065\u002c\u0020\u0062\u0075\u0074\u0020\u0068\u0061\u0073\u0020\u004b\u0069\u0064\u0073\u002e\u0020\u0041\u0073\u0073u\u006di\u006e\u0067\u0020\u0050\u0061\u0067\u0065\u0073 \u006eo\u0064\u0065.")
		_adgee = _abf.MakeName("\u0050\u0061\u0067e\u0073")
		_fddgd.Set("\u0054\u0079\u0070\u0065", _adgee)
	}
	_acd.Log.Trace("\u0062\u0075\u0069\u006c\u0064\u0050a\u0067\u0065\u004c\u0069\u0073\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0074y\u0070\u0065\u003a\u0020\u0025\u0073\u0020(\u0025\u002b\u0076\u0029", *_adgee, _cgfee)
	if *_adgee == "\u0050\u0061\u0067\u0065" {
		_deefg, _dgfed := _aace.newPdfPageFromDict(_fddgd)
		if _dgfed != nil {
			return _dgfed
		}
		_deefg.setContainer(_cgfee)
		if _fbafgc != nil {
			_fddgd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _fbafgc)
		}
		_aace._gbfaf = append(_aace._gbfaf, _cgfee)
		_aace.PageList = append(_aace.PageList, _deefg)
		return nil
	}
	if *_adgee != "\u0050\u0061\u0067e\u0073" {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0063\u006fnt\u0065n\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u0021\u0020\u0028\u0025\u0073\u0029", _adgee)
		return _fd.New("\u0074\u0061\u0062\u006c\u0065\u0020o\u0066\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067 \u006e\u006f\u006e\u0020\u0050\u0061\u0067\u0065\u002f\u0050\u0061\u0067\u0065\u0073 \u006fb\u006a\u0065\u0063\u0074")
	}
	if _fbafgc != nil {
		_fddgd.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _fbafgc)
	}
	if !_aace._abgge {
		_aeba := _aace.traverseObjectData(_cgfee)
		if _aeba != nil {
			return _aeba
		}
	}
	_gafgb, _dggaf := _aace._bebc.Resolve(_fddgd.Get("\u004b\u0069\u0064\u0073"))
	if _dggaf != nil {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069\u006eg\u0020\u004b\u0069\u0064\u0073\u0020\u006fb\u006a\u0065\u0063\u0074")
		return _dggaf
	}
	var _adbeg *_abf.PdfObjectArray
	_adbeg, _dfba = _gafgb.(*_abf.PdfObjectArray)
	if !_dfba {
		_cgebe, _bcfgf := _gafgb.(*_abf.PdfIndirectObject)
		if !_bcfgf {
			return _fd.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_adbeg, _dfba = _cgebe.PdfObject.(*_abf.PdfObjectArray)
		if !_dfba {
			return _fd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u004b\u0069\u0064\u0073\u0020\u0069\u006ed\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
	}
	_acd.Log.Trace("\u004b\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _adbeg)
	for _geebe, _acggecb := range _adbeg.Elements() {
		_afcec, _afeef := _abf.GetIndirect(_acggecb)
		if !_afeef {
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074 \u006f\u0062\u006a\u0065\u0063t\u0020\u002d \u0028\u0025\u0073\u0029", _afcec)
			return _fd.New("\u0070a\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0064\u0069r\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		}
		_adbeg.Set(_geebe, _afcec)
		_dggaf = _aace.buildPageList(_afcec, _cgfee, _feaa)
		if _dggaf != nil {
			return _dggaf
		}
	}
	return nil
}

func (_fbede *PdfWriter) writeBytes(_aabca []byte) {
	if _fbede._dacaeg != nil {
		return
	}
	_gaeed, _dbdbg := _fbede._agfba.Write(_aabca)
	_fbede._dbfaad += int64(_gaeed)
	_fbede._dacaeg = _dbdbg
}

// ToPdfObject implements interface PdfModel.
func (_gbef *PdfAnnotation3D) ToPdfObject() _abf.PdfObject {
	_gbef.PdfAnnotation.ToPdfObject()
	_bbd := _gbef._dbc
	_cca := _bbd.PdfObject.(*_abf.PdfObjectDictionary)
	_cca.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0033\u0044"))
	_cca.SetIfNotNil("\u0033\u0044\u0044", _gbef.T3DD)
	_cca.SetIfNotNil("\u0033\u0044\u0056", _gbef.T3DV)
	_cca.SetIfNotNil("\u0033\u0044\u0041", _gbef.T3DA)
	_cca.SetIfNotNil("\u0033\u0044\u0049", _gbef.T3DI)
	_cca.SetIfNotNil("\u0033\u0044\u0042", _gbef.T3DB)
	return _bbd
}

// ToInteger convert to an integer format.
func (_acgb *PdfColorCalRGB) ToInteger(bits int) [3]uint32 {
	_feac := _ge.Pow(2, float64(bits)) - 1
	return [3]uint32{uint32(_feac * _acgb.A()), uint32(_feac * _acgb.B()), uint32(_feac * _acgb.C())}
}

// VRI represents a Validation-Related Information dictionary.
// The VRI dictionary contains validation data in the form of
// certificates, OCSP and CRL information, for a single signature.
// See ETSI TS 102 778-4 V1.1.1 for more information.
type VRI struct {
	Cert []*_abf.PdfObjectStream
	OCSP []*_abf.PdfObjectStream
	CRL  []*_abf.PdfObjectStream
	TU   *_abf.PdfObjectString
	TS   *_abf.PdfObjectString
}

// PdfAcroForm represents the AcroForm dictionary used for representation of form data in PDF.
type PdfAcroForm struct {
	Fields          *[]*PdfField
	NeedAppearances *_abf.PdfObjectBool
	SigFlags        *_abf.PdfObjectInteger
	CO              *_abf.PdfObjectArray
	DR              *PdfPageResources
	DA              *_abf.PdfObjectString
	Q               *_abf.PdfObjectInteger
	XFA             _abf.PdfObject

	// ADBEEchoSign extra objects from Adobe Acrobat, causing signature invalid if not exists.
	ADBEEchoSign _abf.PdfObject
	_bgfc        *_abf.PdfIndirectObject
	_dfebf       bool
}

// GetDSS gets the DSS dictionary (ETSI TS 102 778-4 V1.1.1) of the current
// document revision.
func (_aaeff *PdfAppender) GetDSS() (_bagc *DSS) { return _aaeff._ffbe }

// GetContainingPdfObject implements interface PdfModel.
func (_de *PdfAction) GetContainingPdfObject() _abf.PdfObject { return _de._egg }

// CompliancePdfReader is a wrapper over PdfReader that is used for verifying if the input Pdf document matches the
// compliance rules of standards like PDF/A.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
type CompliancePdfReader struct {
	*PdfReader
	_fcgbc _abf.ParserMetadata
}

// PdfDate represents a date, which is a PDF string of the form:
// (D:YYYYMMDDHHmmSSOHH'mm)
type PdfDate struct {
	_fabd   int64
	_fcdacf int64
	_gecdc  int64
	_ebda   int64
	_efba   int64
	_fgddf  int64
	_aggabc byte
	_dbgccd int64
	_ccfca  int64
}

// ToPdfObject implements interface PdfModel.
func (_bcb *PdfAnnotationTrapNet) ToPdfObject() _abf.PdfObject {
	_bcb.PdfAnnotation.ToPdfObject()
	_fceg := _bcb._dbc
	_feeb := _fceg.PdfObject.(*_abf.PdfObjectDictionary)
	_feeb.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0054r\u0061\u0070\u004e\u0065\u0074"))
	return _fceg
}

// AddExtGState adds a graphics state to the XObject resources.
func (_dggcd *PdfPage) AddExtGState(name _abf.PdfObjectName, egs *_abf.PdfObjectDictionary) error {
	if _dggcd.Resources == nil {
		_dggcd.Resources = NewPdfPageResources()
	}
	if _dggcd.Resources.ExtGState == nil {
		_dggcd.Resources.ExtGState = _abf.MakeDict()
	}
	_gcadd, _cdecgd := _abf.TraceToDirectObject(_dggcd.Resources.ExtGState).(*_abf.PdfObjectDictionary)
	if !_cdecgd {
		_acd.Log.Debug("\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064i\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u003a\u0020\u0025\u0076", _abf.TraceToDirectObject(_dggcd.Resources.ExtGState))
		return _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gcadd.Set(name, egs)
	return nil
}

// EncryptionAlgorithm is used in EncryptOptions to change the default algorithm used to encrypt the document.
type EncryptionAlgorithm int

// XObjectForm (Table 95 in 8.10.2).
type XObjectForm struct {
	Filter        _abf.StreamEncoder
	FormType      _abf.PdfObject
	BBox          _abf.PdfObject
	Matrix        _abf.PdfObject
	Resources     *PdfPageResources
	Group         _abf.PdfObject
	Ref           _abf.PdfObject
	MetaData      _abf.PdfObject
	PieceInfo     _abf.PdfObject
	LastModified  _abf.PdfObject
	StructParent  _abf.PdfObject
	StructParents _abf.PdfObject
	OPI           _abf.PdfObject
	OC            _abf.PdfObject
	Name          _abf.PdfObject

	// Stream data.
	Stream []byte
	_dbba  *_abf.PdfObjectStream
}

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components. The slice should contain three elements representing the
// L (range 0-100), A (range -100-100) and B (range -100-100) components of
// the color.
func (_ceaa *PdfColorspaceLab) ColorFromFloats(vals []float64) (PdfColor, error) {
	if len(vals) != 3 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_bfab := vals[0]
	if _bfab < 0.0 || _bfab > 100.0 {
		_acd.Log.Debug("\u004c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030-\u0031\u0030\u0030\u0029", _bfab)
		return nil, ErrColorOutOfRange
	}
	_gdbe := vals[1]
	_ebfbf := float64(-100)
	_feaga := float64(100)
	if len(_ceaa.Range) > 1 {
		_ebfbf = _ceaa.Range[0]
		_feaga = _ceaa.Range[1]
	}
	if _gdbe < _ebfbf || _gdbe > _feaga {
		_acd.Log.Debug("\u0041\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _gdbe, _ebfbf, _feaga)
		return nil, ErrColorOutOfRange
	}
	_abgdb := vals[2]
	_deda := float64(-100)
	_cegf := float64(100)
	if len(_ceaa.Range) > 3 {
		_deda = _ceaa.Range[2]
		_cegf = _ceaa.Range[3]
	}
	if _abgdb < _deda || _abgdb > _cegf {
		_acd.Log.Debug("\u0062\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067e\u0020\u0028\u0067\u006f\u0074\u0020%\u0076\u003b\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0025\u0076\u0020\u0074o\u0020\u0025\u0076\u0029", _abgdb, _deda, _cegf)
		return nil, ErrColorOutOfRange
	}
	_cgbdc := NewPdfColorLab(_bfab, _gdbe, _abgdb)
	return _cgbdc, nil
}

func (_acabd *PdfFilespec) getDict() *_abf.PdfObjectDictionary {
	if _gbeff, _ddcdb := _acabd._badbg.(*_abf.PdfIndirectObject); _ddcdb {
		_fefg, _eadd := _gbeff.PdfObject.(*_abf.PdfObjectDictionary)
		if !_eadd {
			return nil
		}
		return _fefg
	} else if _dcaad, _cffg := _acabd._badbg.(*_abf.PdfObjectDictionary); _cffg {
		return _dcaad
	} else {
		_acd.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020F\u0069\u006c\u0065\u0073\u0070\u0065\u0063\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074 \u0074\u0079p\u0065\u0020(\u0025T\u0029", _acabd._badbg)
		return nil
	}
}

// HasFontByName checks if has font resource by name.
func (_bddf *PdfPage) HasFontByName(name _abf.PdfObjectName) bool {
	_fccbd, _ffgad := _bddf.Resources.Font.(*_abf.PdfObjectDictionary)
	if !_ffgad {
		return false
	}
	if _bfcdc := _fccbd.Get(name); _bfcdc != nil {
		return true
	}
	return false
}

// SetName sets the `Name` field of the signature.
func (_gacgbe *PdfSignature) SetName(name string) { _gacgbe.Name = _abf.MakeEncodedString(name, true) }

// GetOptimizer returns current PDF optimizer.
func (_bffdb *PdfWriter) GetOptimizer() Optimizer { return _bffdb._cacbf }

// DetermineColorspaceNameFromPdfObject determines PDF colorspace from a PdfObject.  Returns the colorspace name and
// an error on failure. If the colorspace was not found, will return an empty string.
func DetermineColorspaceNameFromPdfObject(obj _abf.PdfObject) (_abf.PdfObjectName, error) {
	var _gedcg *_abf.PdfObjectName
	var _dccfg *_abf.PdfObjectArray
	if _bbg, _cage := obj.(*_abf.PdfIndirectObject); _cage {
		if _dfgad, _agee := _bbg.PdfObject.(*_abf.PdfObjectArray); _agee {
			_dccfg = _dfgad
		} else if _bcgeg, _cffe := _bbg.PdfObject.(*_abf.PdfObjectName); _cffe {
			_gedcg = _bcgeg
		}
	} else if _adbg, _acae := obj.(*_abf.PdfObjectArray); _acae {
		_dccfg = _adbg
	} else if _dbab, _facb := obj.(*_abf.PdfObjectName); _facb {
		_gedcg = _dbab
	}
	if _gedcg != nil {
		switch *_gedcg {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			return *_gedcg, nil
		case "\u0050a\u0074\u0074\u0065\u0072\u006e":
			return *_gedcg, nil
		}
	}
	if _dccfg != nil && _dccfg.Len() > 0 {
		if _decab, _gda := _dccfg.Get(0).(*_abf.PdfObjectName); _gda {
			switch *_decab {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				if _dccfg.Len() == 1 {
					return *_decab, nil
				}
			case "\u0043a\u006c\u0047\u0072\u0061\u0079", "\u0043\u0061\u006c\u0052\u0047\u0042", "\u004c\u0061\u0062":
				return *_decab, nil
			case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064", "\u0050a\u0074\u0074\u0065\u0072\u006e", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				return *_decab, nil
			case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e", "\u0044e\u0076\u0069\u0063\u0065\u004e":
				return *_decab, nil
			}
		}
	}
	return "", nil
}

// NewPdfActionThread returns a new "thread" action.
func NewPdfActionThread() *PdfActionThread {
	_eee := NewPdfAction()
	_gfd := &PdfActionThread{}
	_gfd.PdfAction = _eee
	_eee.SetContext(_gfd)
	return _gfd
}

// PdfReader represents a PDF file reader. It is a frontend to the lower level parsing mechanism and provides
// a higher level access to work with PDF structure and information, such as the page structure etc.
type PdfReader struct {
	_bebc    *_abf.PdfParser
	_afdaf   _abf.PdfObject
	_bfdff   *_abf.PdfIndirectObject
	_agbecg  *_abf.PdfObjectDictionary
	_gbfaf   []*_abf.PdfIndirectObject
	PageList []*PdfPage
	_gcegc   int
	_dagde   *_abf.PdfObjectDictionary
	_cggee   *PdfOutlineTreeNode
	AcroForm *PdfAcroForm
	DSS      *DSS
	Rotate   *int64
	_gedbg   *Permissions
	_bfced   map[*PdfReader]*PdfReader
	_egade   []*PdfReader
	_ceecd   *modelManager
	_abgge   bool
	_ggbccc  map[_abf.PdfObject]struct{}
	_affbb   _gc.ReadSeeker
	_bccga   string
	_dfafc   bool
	_gebfg   *ReaderOpts
	_dbgdg   bool
}

// GetNumComponents returns the number of color components (3 for RGB).
func (_adge *PdfColorDeviceRGB) GetNumComponents() int { return 3 }

// MergePageWith appends page content to source Pdf file page content.
func (_afedf *PdfAppender) MergePageWith(pageNum int, page *PdfPage) error {
	_fdca := pageNum - 1
	var _bada *PdfPage
	for _effc, _dgeg := range _afedf._cggfa {
		if _effc == _fdca {
			_bada = _dgeg
		}
	}
	if _bada == nil {
		return _e.Errorf("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073o\u0075\u0072\u0063\u0065\u0020\u0064o\u0063\u0075\u006de\u006e\u0074", pageNum)
	}
	if _bada._gefee != nil && _bada._gefee.GetParser() == _afedf._agda._bebc {
		_bada = _bada.Duplicate()
		_afedf._cggfa[_fdca] = _bada
	}
	page = page.Duplicate()
	_ccgc := _gabc(_bada)
	_ccege := _gabc(page)
	_bdcdg := make(map[_abf.PdfObjectName]_abf.PdfObjectName)
	for _cggb := range _ccege {
		if _, _dggc := _ccgc[_cggb]; _dggc {
			for _daaf := 1; true; _daaf++ {
				_fdcf := _abf.PdfObjectName(string(_cggb) + _gb.Itoa(_daaf))
				if _, _abgae := _ccgc[_fdcf]; !_abgae {
					_bdcdg[_cggb] = _fdcf
					break
				}
			}
		}
	}
	_egc, _beee := page.GetContentStreams()
	if _beee != nil {
		return _beee
	}
	_bbea, _beee := _bada.GetContentStreams()
	if _beee != nil {
		return _beee
	}
	for _fdbca, _eaeb := range _egc {
		for _cbe, _fbd := range _bdcdg {
			_eaeb = _be.Replace(_eaeb, "\u002f"+string(_cbe), "\u002f"+string(_fbd), -1)
		}
		_egc[_fdbca] = _eaeb
	}
	_bbea = append(_bbea, _egc...)
	if _gabf := _bada.SetContentStreams(_bbea, _abf.NewFlateEncoder()); _gabf != nil {
		return _gabf
	}
	_bada._baagf = append(_bada._baagf, page._baagf...)
	if _bada.Resources == nil {
		_bada.Resources = NewPdfPageResources()
	}
	if page.Resources != nil {
		_bada.Resources.Font = _afedf.mergeResources(_bada.Resources.Font, page.Resources.Font, _bdcdg)
		_bada.Resources.XObject = _afedf.mergeResources(_bada.Resources.XObject, page.Resources.XObject, _bdcdg)
		_bada.Resources.Properties = _afedf.mergeResources(_bada.Resources.Properties, page.Resources.Properties, _bdcdg)
		if _bada.Resources.ProcSet == nil {
			_bada.Resources.ProcSet = page.Resources.ProcSet
		}
		_bada.Resources.Shading = _afedf.mergeResources(_bada.Resources.Shading, page.Resources.Shading, _bdcdg)
		_bada.Resources.ExtGState = _afedf.mergeResources(_bada.Resources.ExtGState, page.Resources.ExtGState, _bdcdg)
	}
	_cfbd, _beee := _bada.GetMediaBox()
	if _beee != nil {
		return _beee
	}
	_egeb, _beee := page.GetMediaBox()
	if _beee != nil {
		return _beee
	}
	var _dagd bool
	if _cfbd.Llx > _egeb.Llx {
		_cfbd.Llx = _egeb.Llx
		_dagd = true
	}
	if _cfbd.Lly > _egeb.Lly {
		_cfbd.Lly = _egeb.Lly
		_dagd = true
	}
	if _cfbd.Urx < _egeb.Urx {
		_cfbd.Urx = _egeb.Urx
		_dagd = true
	}
	if _cfbd.Ury < _egeb.Ury {
		_cfbd.Ury = _egeb.Ury
		_dagd = true
	}
	if _dagd {
		_bada.MediaBox = _cfbd
	}
	return nil
}

// SetContentStream sets the pattern cell's content stream.
func (_bdcdef *PdfTilingPattern) SetContentStream(content []byte, encoder _abf.StreamEncoder) error {
	_fbddf, _bcdca := _bdcdef._bcfca.(*_abf.PdfObjectStream)
	if !_bcdca {
		_acd.Log.Debug("\u0054\u0069l\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _bdcdef._bcfca)
		return _abf.ErrTypeError
	}
	if encoder == nil {
		encoder = _abf.NewRawEncoder()
	}
	_adgca := _fbddf.PdfObjectDictionary
	_cbefg := encoder.MakeStreamDict()
	_adgca.Merge(_cbefg)
	_ffcgc, _dcbae := encoder.EncodeBytes(content)
	if _dcbae != nil {
		return _dcbae
	}
	_adgca.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _abf.MakeInteger(int64(len(_ffcgc))))
	_fbddf.Stream = _ffcgc
	return nil
}

// GetFontDescriptor returns the font descriptor for `font`.
func (_gdccd PdfFont) GetFontDescriptor() (*PdfFontDescriptor, error) {
	return _gdccd._gedca.getFontDescriptor(), nil
}

// ReplaceAcroForm replaces the acrobat form. It appends a new form to the Pdf which
// replaces the original AcroForm.
func (_cgaa *PdfAppender) ReplaceAcroForm(acroForm *PdfAcroForm) {
	if acroForm != nil {
		_cgaa.updateObjectsDeep(acroForm.ToPdfObject(), nil)
	}
	_cgaa._ffbb = acroForm
}

// NewPdfPageResources returns a new PdfPageResources object.
func NewPdfPageResources() *PdfPageResources {
	_bdbdf := &PdfPageResources{}
	_bdbdf._gagb = _abf.MakeDict()
	return _bdbdf
}

// HasXObjectByName checks if has XObject resource by name.
func (_ffeae *PdfPage) HasXObjectByName(name _abf.PdfObjectName) bool {
	_bfbd, _gbdg := _ffeae.Resources.XObject.(*_abf.PdfObjectDictionary)
	if !_gbdg {
		return false
	}
	if _abgg := _bfbd.Get(name); _abgg != nil {
		return true
	}
	return false
}

// GetOutlinesFlattened returns a flattened list of tree nodes and titles.
// NOTE: for most use cases, it is recommended to use the high-level GetOutlines
// method instead, which also provides information regarding the destination
// of the outline items.
func (_beegc *PdfReader) GetOutlinesFlattened() ([]*PdfOutlineTreeNode, []string, error) {
	var _cefge []*PdfOutlineTreeNode
	var _fabae []string
	var _daafbg func(*PdfOutlineTreeNode, *[]*PdfOutlineTreeNode, *[]string, int)
	_daafbg = func(_dfcfa *PdfOutlineTreeNode, _gdfge *[]*PdfOutlineTreeNode, _adege *[]string, _egde int) {
		if _dfcfa == nil {
			return
		}
		if _dfcfa._aecec == nil {
			_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020M\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006e\u006fd\u0065\u002e\u0063o\u006et\u0065\u0078\u0074")
			return
		}
		_gdeecf, _abedc := _dfcfa._aecec.(*PdfOutlineItem)
		if _abedc {
			*_gdfge = append(*_gdfge, &_gdeecf.PdfOutlineTreeNode)
			_agfag := _be.Repeat("\u0020", _egde*2) + _gdeecf.Title.Decoded()
			*_adege = append(*_adege, _agfag)
		}
		if _dfcfa.First != nil {
			_fgae := _be.Repeat("\u0020", _egde*2) + "\u002b"
			*_adege = append(*_adege, _fgae)
			_daafbg(_dfcfa.First, _gdfge, _adege, _egde+1)
		}
		if _abedc && _gdeecf.Next != nil {
			_daafbg(_gdeecf.Next, _gdfge, _adege, _egde)
		}
	}
	_daafbg(_beegc._cggee, &_cefge, &_fabae, 0)
	return _cefge, _fabae, nil
}

// ToPdfObject implements interface PdfModel.
func (_ffa *PdfActionRendition) ToPdfObject() _abf.PdfObject {
	_ffa.PdfAction.ToPdfObject()
	_deb := _ffa._egg
	_cdb := _deb.PdfObject.(*_abf.PdfObjectDictionary)
	_cdb.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeRendition)))
	_cdb.SetIfNotNil("\u0052", _ffa.R)
	_cdb.SetIfNotNil("\u0041\u004e", _ffa.AN)
	_cdb.SetIfNotNil("\u004f\u0050", _ffa.OP)
	_cdb.SetIfNotNil("\u004a\u0053", _ffa.JS)
	return _deb
}

func (_bfbf *PdfReader) newPdfAnnotationWatermarkFromDict(_cceb *_abf.PdfObjectDictionary) (*PdfAnnotationWatermark, error) {
	_dccee := PdfAnnotationWatermark{}
	_dccee.FixedPrint = _cceb.Get("\u0046\u0069\u0078\u0065\u0064\u0050\u0072\u0069\u006e\u0074")
	return &_dccee, nil
}

// SetPdfKeywords sets the Keywords attribute of the output PDF.
func SetPdfKeywords(keywords string) { _gaabd.Lock(); defer _gaabd.Unlock(); _geggga = keywords }

// PdfShadingType7 is a Tensor-product patch mesh.
type PdfShadingType7 struct {
	*PdfShading
	BitsPerCoordinate *_abf.PdfObjectInteger
	BitsPerComponent  *_abf.PdfObjectInteger
	BitsPerFlag       *_abf.PdfObjectInteger
	Decode            *_abf.PdfObjectArray
	Function          []PdfFunction
}

// ToPdfObject implements interface PdfModel.
func (_fgdb *PdfAnnotationProjection) ToPdfObject() _abf.PdfObject {
	_fgdb.PdfAnnotation.ToPdfObject()
	_gafc := _fgdb._dbc
	_bfd := _gafc.PdfObject.(*_abf.PdfObjectDictionary)
	_fgdb.PdfAnnotationMarkup.appendToPdfDictionary(_bfd)
	return _gafc
}

// ColorToRGB converts a Lab color to an RGB color.
func (_aebd *PdfColorspaceLab) ColorToRGB(color PdfColor) (PdfColor, error) {
	_eegg := func(_dgcc float64) float64 {
		if _dgcc >= 6.0/29 {
			return _dgcc * _dgcc * _dgcc
		}
		return 108.0 / 841 * (_dgcc - 4.0/29.0)
	}
	_bcagb, _agaa := color.(*PdfColorLab)
	if !_agaa {
		_acd.Log.Debug("\u0069\u006e\u0070\u0075t \u0063\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u006c\u0061\u0062")
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	LStar := _bcagb.L()
	AStar := _bcagb.A()
	BStar := _bcagb.B()
	L := (LStar+16)/116 + AStar/500
	M := (LStar + 16) / 116
	N := (LStar+16)/116 - BStar/200
	X := _aebd.WhitePoint[0] * _eegg(L)
	Y := _aebd.WhitePoint[1] * _eegg(M)
	Z := _aebd.WhitePoint[2] * _eegg(N)
	_ceec := 3.240479*X + -1.537150*Y + -0.498535*Z
	_abag := -0.969256*X + 1.875992*Y + 0.041556*Z
	_cfgaf := 0.055648*X + -0.204043*Y + 1.057311*Z
	_ceec = _ge.Min(_ge.Max(_ceec, 0), 1.0)
	_abag = _ge.Min(_ge.Max(_abag, 0), 1.0)
	_cfgaf = _ge.Min(_ge.Max(_cfgaf, 0), 1.0)
	return NewPdfColorDeviceRGB(_ceec, _abag, _cfgaf), nil
}

// ToPdfObject converts colorspace to a PDF object. [/Indexed base hival lookup]
func (_ecaf *PdfColorspaceSpecialIndexed) ToPdfObject() _abf.PdfObject {
	_bffcd := _abf.MakeArray(_abf.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"))
	_bffcd.Append(_ecaf.Base.ToPdfObject())
	_bffcd.Append(_abf.MakeInteger(int64(_ecaf.HiVal)))
	_bffcd.Append(_ecaf.Lookup)
	if _ecaf._acea != nil {
		_ecaf._acea.PdfObject = _bffcd
		return _ecaf._acea
	}
	return _bffcd
}
func _dgdfd() _f.Time { _gaabd.Lock(); defer _gaabd.Unlock(); return _egdgg }

// PdfColorPatternType2 represents a color shading pattern type 2 (Axial).
type PdfColorPatternType2 struct {
	Color       PdfColor
	PatternName _abf.PdfObjectName
}

// RemovePage removes a page by number.
func (_gaaa *PdfAppender) RemovePage(pageNum int) {
	_eegf := pageNum - 1
	_gaaa._cggfa = append(_gaaa._cggfa[0:_eegf], _gaaa._cggfa[pageNum:]...)
}

// GetCustomInfo returns a custom info value for the specified name.
func (_geaf *PdfInfo) GetCustomInfo(name string) *_abf.PdfObjectString {
	var _fgdf *_abf.PdfObjectString
	if _geaf._cbf == nil {
		return _fgdf
	}
	if _bdde, _gbec := _geaf._cbf.Get(*_abf.MakeName(name)).(*_abf.PdfObjectString); _gbec {
		_fgdf = _bdde
	}
	return _fgdf
}

func (_bbabe *PdfWriter) writeTrailer(_dfggfg int) {
	_bbabe.writeString("\u0078\u0072\u0065\u0066\u000d\u000a")
	for _ebbeb := 0; _ebbeb <= _dfggfg; {
		for ; _ebbeb <= _dfggfg; _ebbeb++ {
			_gcece, _gecaf := _bbabe._becfc[_ebbeb]
			if _gecaf && (!_bbabe._aegbd || _bbabe._aegbd && (_gcece.Type == 1 && _gcece.Offset >= _bbabe._cfecga || _gcece.Type == 0)) {
				break
			}
		}
		var _edga int
		for _edga = _ebbeb + 1; _edga <= _dfggfg; _edga++ {
			_bdcdf, _beeea := _bbabe._becfc[_edga]
			if _beeea && (!_bbabe._aegbd || _bbabe._aegbd && (_bdcdf.Type == 1 && _bdcdf.Offset > _bbabe._cfecga)) {
				continue
			}
			break
		}
		_fadef := _e.Sprintf("\u0025d\u0020\u0025\u0064\u000d\u000a", _ebbeb, _edga-_ebbeb)
		_bbabe.writeString(_fadef)
		for _bdgbf := _ebbeb; _bdgbf < _edga; _bdgbf++ {
			_ebbdc := _bbabe._becfc[_bdgbf]
			switch _ebbdc.Type {
			case 0:
				_fadef = _e.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u0066\u000d\u000a", 0, 65535)
				_bbabe.writeString(_fadef)
			case 1:
				_fadef = _e.Sprintf("\u0025\u002e\u0031\u0030\u0064\u0020\u0025\u002e\u0035d\u0020\u006e\u000d\u000a", _ebbdc.Offset, 0)
				_bbabe.writeString(_fadef)
			}
		}
		_ebbeb = _edga + 1
	}
	_babfa := _abf.MakeDict()
	_babfa.Set("\u0049\u006e\u0066\u006f", _bbabe._ddegc)
	_babfa.Set("\u0052\u006f\u006f\u0074", _bbabe._cfdde)
	_babfa.Set("\u0053\u0069\u007a\u0065", _abf.MakeInteger(int64(_dfggfg+1)))
	if _bbabe._aegbd && _bbabe._ffgf > 0 {
		_babfa.Set("\u0050\u0072\u0065\u0076", _abf.MakeInteger(_bbabe._ffgf))
	}
	if _bbabe._ddbgd != nil {
		_babfa.Set("\u0045n\u0063\u0072\u0079\u0070\u0074", _bbabe._dcdbb)
	}
	if _bbabe._dedfdf == nil && _bbabe._aefff != "" && _bbabe._cfbce != "" {
		_bbabe._dedfdf = _abf.MakeArray(_abf.MakeHexString(_bbabe._aefff), _abf.MakeHexString(_bbabe._cfbce))
	}
	if _bbabe._dedfdf != nil {
		_babfa.Set("\u0049\u0044", _bbabe._dedfdf)
		_acd.Log.Trace("\u0049d\u0073\u003a\u0020\u0025\u0073", _bbabe._dedfdf)
	}
	_bbabe.writeString("\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u000a")
	_bbabe.writeString(_babfa.WriteString())
	_bbabe.writeString("\u000a")
}

// ToPdfObject implements interface PdfModel.
func (_eae *PdfAnnotationFileAttachment) ToPdfObject() _abf.PdfObject {
	_eae.PdfAnnotation.ToPdfObject()
	_ccdb := _eae._dbc
	_abd := _ccdb.PdfObject.(*_abf.PdfObjectDictionary)
	_eae.PdfAnnotationMarkup.appendToPdfDictionary(_abd)
	_abd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0046\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063h\u006d\u0065\u006e\u0074"))
	_abd.SetIfNotNil("\u0046\u0053", _eae.FS)
	_abd.SetIfNotNil("\u004e\u0061\u006d\u0065", _eae.Name)
	return _ccdb
}

// NewXObjectFormFromStream builds the Form XObject from a stream object.
// TODO: Should this be exposed? Consider different access points.
func NewXObjectFormFromStream(stream *_abf.PdfObjectStream) (*XObjectForm, error) {
	_bdded := &XObjectForm{}
	_bdded._dbba = stream
	_bbaca := *(stream.PdfObjectDictionary)
	_fbeaf, _gdebg := _abf.NewEncoderFromStream(stream)
	if _gdebg != nil {
		return nil, _gdebg
	}
	_bdded.Filter = _fbeaf
	if _dbfb := _bbaca.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _dbfb != nil {
		_afdfe, _gafbc := _dbfb.(*_abf.PdfObjectName)
		if !_gafbc {
			return nil, _fd.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		if *_afdfe != "\u0046\u006f\u0072\u006d" {
			_acd.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
			return nil, _fd.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072m\u0020\u0073\u0075\u0062ty\u0070\u0065")
		}
	}
	if _aacef := _bbaca.Get("\u0046\u006f\u0072\u006d\u0054\u0079\u0070\u0065"); _aacef != nil {
		_bdded.FormType = _aacef
	}
	if _bbec := _bbaca.Get("\u0042\u0042\u006f\u0078"); _bbec != nil {
		_bdded.BBox = _bbec
	}
	if _ddgeb := _bbaca.Get("\u004d\u0061\u0074\u0072\u0069\u0078"); _ddgeb != nil {
		_bdded.Matrix = _ddgeb
	}
	if _aeeeed := _bbaca.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"); _aeeeed != nil {
		_aeeeed = _abf.TraceToDirectObject(_aeeeed)
		_cafdf, _aefge := _aeeeed.(*_abf.PdfObjectDictionary)
		if !_aefge {
			_acd.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0058\u004f\u0062j\u0065c\u0074\u0020\u0046\u006f\u0072\u006d\u0020\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006f\u0062j\u0065\u0063\u0074\u002c\u0020\u0070\u006f\u0069\u006e\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _abf.ErrTypeError
		}
		_ebfd, _gedea := NewPdfPageResourcesFromDict(_cafdf)
		if _gedea != nil {
			_acd.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0020\u0072\u0065\u0073\u006f\u0075rc\u0065\u0073")
			return nil, _gedea
		}
		_bdded.Resources = _ebfd
		_acd.Log.Trace("\u0046\u006f\u0072\u006d r\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u003a\u0020\u0025\u0023\u0076", _bdded.Resources)
	}
	_bdded.Group = _bbaca.Get("\u0047\u0072\u006fu\u0070")
	_bdded.Ref = _bbaca.Get("\u0052\u0065\u0066")
	_bdded.MetaData = _bbaca.Get("\u004d\u0065\u0074\u0061\u0044\u0061\u0074\u0061")
	_bdded.PieceInfo = _bbaca.Get("\u0050i\u0065\u0063\u0065\u0049\u006e\u0066o")
	_bdded.LastModified = _bbaca.Get("\u004c\u0061\u0073t\u004d\u006f\u0064\u0069\u0066\u0069\u0065\u0064")
	_bdded.StructParent = _bbaca.Get("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074")
	_bdded.StructParents = _bbaca.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074\u0073")
	_bdded.OPI = _bbaca.Get("\u004f\u0050\u0049")
	_bdded.OC = _bbaca.Get("\u004f\u0043")
	_bdded.Name = _bbaca.Get("\u004e\u0061\u006d\u0065")
	_bdded.Stream = stream.Stream
	return _bdded, nil
}

// PdfShadingType2 is an Axial shading.
type PdfShadingType2 struct {
	*PdfShading
	Coords   *_abf.PdfObjectArray
	Domain   *_abf.PdfObjectArray
	Function []PdfFunction
	Extend   *_abf.PdfObjectArray
}

// Set applies flag fl to the flag's bitmask and returns the combined flag.
func (_beddg FieldFlag) Set(fl FieldFlag) FieldFlag { return FieldFlag(_beddg.Mask() | fl.Mask()) }

// Hasher is the interface that wraps the basic Write method.
type Hasher interface {
	Write(_edgga []byte) (_ceaf int, _feaed error)
}

// ToPdfObject implements interface PdfModel.
func (_acbd *PdfAnnotationLink) ToPdfObject() _abf.PdfObject {
	_acbd.PdfAnnotation.ToPdfObject()
	_dae := _acbd._dbc
	_ggcg := _dae.PdfObject.(*_abf.PdfObjectDictionary)
	_ggcg.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u004c\u0069\u006e\u006b"))
	if _acbd._bgad != nil && _acbd._bgad._gfg != nil {
		_ggcg.Set("\u0041", _acbd._bgad._gfg.ToPdfObject())
	} else if _acbd.A != nil {
		_ggcg.Set("\u0041", _acbd.A)
	}
	_ggcg.SetIfNotNil("\u0044\u0065\u0073\u0074", _acbd.Dest)
	_ggcg.SetIfNotNil("\u0048", _acbd.H)
	_ggcg.SetIfNotNil("\u0050\u0041", _acbd.PA)
	_ggcg.SetIfNotNil("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073", _acbd.QuadPoints)
	_ggcg.SetIfNotNil("\u0042\u0053", _acbd.BS)
	return _dae
}

// AddCustomInfo adds a custom info into document info dictionary.
func (_gegcc *PdfInfo) AddCustomInfo(name string, value string) error {
	if _gegcc._cbf == nil {
		_gegcc._cbf = _abf.MakeDict()
	}
	if _, _bece := _abfb[name]; _bece {
		return _e.Errorf("\u0063\u0061\u006e\u006e\u006ft\u0020\u0075\u0073\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072\u0064 \u0069\u006e\u0066\u006f\u0020\u006b\u0065\u0079\u0020\u0025\u0073\u0020\u0061\u0073\u0020\u0063\u0075\u0073\u0074\u006f\u006d\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u006b\u0065y", name)
	}
	_gegcc._cbf.SetIfNotNil(*_abf.MakeName(name), _abf.MakeString(value))
	return nil
}

// NewPdfAnnotationText returns a new text annotation.
func NewPdfAnnotationText() *PdfAnnotationText {
	_edf := NewPdfAnnotation()
	_aba := &PdfAnnotationText{}
	_aba.PdfAnnotation = _edf
	_aba.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_edf.SetContext(_aba)
	return _aba
}

func (_afdfb *PdfWriter) optimize() error {
	if _afdfb._cacbf == nil {
		return nil
	}
	var _gfgdf error
	_afdfb._edcgc, _gfgdf = _afdfb._cacbf.Optimize(_afdfb._edcgc)
	if _gfgdf != nil {
		return _gfgdf
	}
	_bbfeb := make(map[_abf.PdfObject]struct{}, len(_afdfb._edcgc))
	for _, _gddea := range _afdfb._edcgc {
		_bbfeb[_gddea] = struct{}{}
	}
	_afdfb._fdgae = _bbfeb
	return nil
}

// AddPages adds pages to be appended to the end of the source PDF.
func (_fcdd *PdfAppender) AddPages(pages ...*PdfPage) {
	for _, _bafc := range pages {
		_bafc = _bafc.Duplicate()
		_fcdd._cggfa = append(_fcdd._cggfa, _bafc)
	}
}

// G returns the value of the green component of the color.
func (_bebb *PdfColorDeviceRGB) G() float64 { return _bebb[1] }

// SetContentStream updates the content stream with specified encoding.
// If encoding is null, will use the xform.Filter object or Raw encoding if not set.
func (_cddca *XObjectForm) SetContentStream(content []byte, encoder _abf.StreamEncoder) error {
	_cdaab := content
	if encoder == nil {
		if _cddca.Filter != nil {
			encoder = _cddca.Filter
		} else {
			encoder = _abf.NewRawEncoder()
		}
	}
	_afeee, _bgdeg := encoder.EncodeBytes(_cdaab)
	if _bgdeg != nil {
		return _bgdeg
	}
	_cdaab = _afeee
	_cddca.Stream = _cdaab
	_cddca.Filter = encoder
	return nil
}

// String returns the name of the colorspace (DeviceN).
func (_adcd *PdfColorspaceDeviceN) String() string { return "\u0044e\u0076\u0069\u0063\u0065\u004e" }

func (_aeecge *fontFile) loadFromSegments(_afbae, _dcac []byte) error {
	_acd.Log.Trace("\u006c\u006f\u0061dF\u0072\u006f\u006d\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0064\u0020\u0025\u0064", len(_afbae), len(_dcac))
	_badbe := _aeecge.parseASCIIPart(_afbae)
	if _badbe != nil {
		return _badbe
	}
	_acd.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _aeecge)
	if len(_dcac) == 0 {
		return nil
	}
	_acd.Log.Trace("f\u006f\u006e\u0074\u0066\u0069\u006c\u0065\u003d\u0025\u0073", _aeecge)
	return nil
}

func _ebedg(_defaa _abf.PdfObject) (PdfFunction, error) {
	_defaa = _abf.ResolveReference(_defaa)
	if _bfgb, _fgbeg := _defaa.(*_abf.PdfObjectStream); _fgbeg {
		_fbafgf := _bfgb.PdfObjectDictionary
		_caada, _efge := _fbafgf.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_abf.PdfObjectInteger)
		if !_efge {
			_acd.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _fd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_caada == 0 {
			return _ebeg(_bfgb)
		} else if *_caada == 4 {
			return _cgddd(_bfgb)
		} else {
			return nil, _fd.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _gfbec, _aggf := _defaa.(*_abf.PdfIndirectObject); _aggf {
		_aafe, _eebbe := _gfbec.PdfObject.(*_abf.PdfObjectDictionary)
		if !_eebbe {
			_acd.Log.Error("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006eg\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			return nil, _fd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		_eeccg, _eebbe := _aafe.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_abf.PdfObjectInteger)
		if !_eebbe {
			_acd.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _fd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_eeccg == 2 {
			return _fgeba(_gfbec)
		} else if *_eeccg == 3 {
			return _acgge(_gfbec)
		} else {
			return nil, _fd.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else if _aaaa, _cabda := _defaa.(*_abf.PdfObjectDictionary); _cabda {
		_dbgbb, _befab := _aaaa.Get("\u0046\u0075\u006ec\u0074\u0069\u006f\u006e\u0054\u0079\u0070\u0065").(*_abf.PdfObjectInteger)
		if !_befab {
			_acd.Log.Error("F\u0075\u006e\u0063\u0074\u0069\u006fn\u0054\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006di\u0073s\u0069\u006e\u0067")
			return nil, _fd.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072 \u006f\u0072\u0020\u006d\u0069\u0073\u0073i\u006e\u0067")
		}
		if *_dbgbb == 2 {
			return _fgeba(_aaaa)
		} else if *_dbgbb == 3 {
			return _acgge(_aaaa)
		} else {
			return nil, _fd.New("i\u006e\u0076\u0061\u006cid\u0020f\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0074\u0079\u0070\u0065")
		}
	} else {
		_acd.Log.Debug("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e\u0020\u0054\u0079\u0070\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0023\u0076", _defaa)
		return nil, _fd.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
}

// GetContainingPdfObject returns the container of the resources object (indirect object).
func (_ccbcb *PdfPageResources) GetContainingPdfObject() _abf.PdfObject { return _ccbcb._gagb }

// String returns a string that describes `base`.
func (_ebbdb fontCommon) String() string {
	return _e.Sprintf("\u0046\u004f\u004e\u0054\u007b\u0025\u0073\u007d", _ebbdb.coreString())
}

// GetPage returns the PdfPage model for the specified page number.
func (_gfcfacd *PdfReader) GetPage(pageNumber int) (*PdfPage, error) {
	if _gfcfacd._bebc.GetCrypter() != nil && !_gfcfacd._bebc.IsAuthenticated() {
		return nil, _e.Errorf("\u0066\u0069\u006c\u0065\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020d\u0065\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074")
	}
	if len(_gfcfacd._gbfaf) < pageNumber {
		return nil, _fd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0028\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0074o\u006f\u0020\u0073\u0068\u006f\u0072\u0074\u0029")
	}
	_edcc := pageNumber - 1
	if _edcc < 0 {
		return nil, _e.Errorf("\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0069\u006e\u0067\u0020\u006d\u0075\u0073t\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0061\u0074\u0020\u0031")
	}
	_gfgcb := _gfcfacd.PageList[_edcc]
	return _gfgcb, nil
}

func _bfabe(_bcde StdFontName) (pdfFontSimple, error) {
	_beggc, _eddb := _gbe.NewStdFontByName(_bcde)
	if !_eddb {
		return pdfFontSimple{}, ErrFontNotSupported
	}
	_ddgf := _bcee(_beggc)
	return _ddgf, nil
}

// GetEncryptionMethod returns a descriptive information string about the encryption method used.
func (_aaggd *PdfReader) GetEncryptionMethod() string {
	_bdace := _aaggd._bebc.GetCrypter()
	return _bdace.String()
}

// NewCompliancePdfReader creates a PdfReader or an input io.ReadSeeker that during reading will scan the files for the
// metadata details. It could be used for the PDF standard implementations like PDF/A or PDF/X.
// NOTE: This implementation is in experimental development state.
//
//	Keep in mind that it might change in the subsequent minor versions.
func NewCompliancePdfReader(rs _gc.ReadSeeker) (*CompliancePdfReader, error) {
	const _cacf = "\u006d\u006f\u0064\u0065l\u003a\u004e\u0065\u0077\u0043\u006f\u006d\u0070\u006c\u0069a\u006ec\u0065\u0050\u0064\u0066\u0052\u0065\u0061d\u0065\u0072"
	_gagda, _bbfb := _fbaec(rs, &ReaderOpts{ComplianceMode: true}, false, _cacf)
	if _bbfb != nil {
		return nil, _bbfb
	}
	return &CompliancePdfReader{PdfReader: _gagda}, nil
}

// PdfActionGoToE represents a GoToE action.
type PdfActionGoToE struct {
	*PdfAction
	F         *PdfFilespec
	D         _abf.PdfObject
	NewWindow _abf.PdfObject
	T         _abf.PdfObject
}

func (_gfbdb *PdfAcroForm) signatureFields() []*PdfFieldSignature {
	var _bbcgb []*PdfFieldSignature
	for _, _agbfc := range _gfbdb.AllFields() {
		switch _dbacf := _agbfc.GetContext().(type) {
		case *PdfFieldSignature:
			_fcbbdf := _dbacf
			_bbcgb = append(_bbcgb, _fcbbdf)
		}
	}
	return _bbcgb
}

// NewPdfAnnotationWatermark returns a new watermark annotation.
func NewPdfAnnotationWatermark() *PdfAnnotationWatermark {
	_baf := NewPdfAnnotation()
	_cgc := &PdfAnnotationWatermark{}
	_cgc.PdfAnnotation = _baf
	_baf.SetContext(_cgc)
	return _cgc
}

// ToGray returns a PdfColorDeviceGray color based on the current RGB color.
func (_gfef *PdfColorDeviceRGB) ToGray() *PdfColorDeviceGray {
	_fbdb := 0.3*_gfef.R() + 0.59*_gfef.G() + 0.11*_gfef.B()
	_fbdb = _ge.Min(_ge.Max(_fbdb, 0.0), 1.0)
	return NewPdfColorDeviceGray(_fbdb)
}

// NewPdfColorspaceDeviceN returns an initialized PdfColorspaceDeviceN.
func NewPdfColorspaceDeviceN() *PdfColorspaceDeviceN {
	_gfggca := &PdfColorspaceDeviceN{}
	return _gfggca
}

func (_gbfg *PdfReader) newPdfActionThreadFromDict(_ebe *_abf.PdfObjectDictionary) (*PdfActionThread, error) {
	_ggcf, _cfa := _dgf(_ebe.Get("\u0046"))
	if _cfa != nil {
		return nil, _cfa
	}
	return &PdfActionThread{D: _ebe.Get("\u0044"), B: _ebe.Get("\u0042"), F: _ggcf}, nil
}

// NewReaderForText makes a new PdfReader for an input PDF content string. For use in testing.
func NewReaderForText(txt string) *PdfReader {
	return &PdfReader{_ggbccc: map[_abf.PdfObject]struct{}{}, _ceecd: _gadf(), _bebc: _abf.NewParserFromString(txt)}
}

// GetRuneMetrics returns the character metrics for the specified rune.
// A bool flag is returned to indicate whether or not the entry was found.
func (_bcea pdfCIDFontType0) GetRuneMetrics(r rune) (_gbe.CharMetrics, bool) {
	return _gbe.CharMetrics{Wx: _bcea._bdced}, true
}

// PdfAnnotationFileAttachment represents FileAttachment annotations.
// (Section 12.5.6.15).
type PdfAnnotationFileAttachment struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	FS   _abf.PdfObject
	Name _abf.PdfObject
}

// ToPdfObject implements interface PdfModel.
func (_dgaggf *PdfSignature) ToPdfObject() _abf.PdfObject {
	_faega := _dgaggf._geebd
	var _cbdef *_abf.PdfObjectDictionary
	if _cagea, _eaee := _faega.PdfObject.(*pdfSignDictionary); _eaee {
		_cbdef = _cagea.PdfObjectDictionary
	} else {
		_cbdef = _faega.PdfObject.(*_abf.PdfObjectDictionary)
	}
	_cbdef.SetIfNotNil("\u0054\u0079\u0070\u0065", _dgaggf.Type)
	_cbdef.SetIfNotNil("\u0046\u0069\u006c\u0074\u0065\u0072", _dgaggf.Filter)
	_cbdef.SetIfNotNil("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r", _dgaggf.SubFilter)
	_cbdef.SetIfNotNil("\u0042y\u0074\u0065\u0052\u0061\u006e\u0067e", _dgaggf.ByteRange)
	_cbdef.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _dgaggf.Contents)
	_cbdef.SetIfNotNil("\u0043\u0065\u0072\u0074", _dgaggf.Cert)
	_cbdef.SetIfNotNil("\u004e\u0061\u006d\u0065", _dgaggf.Name)
	_cbdef.SetIfNotNil("\u0052\u0065\u0061\u0073\u006f\u006e", _dgaggf.Reason)
	_cbdef.SetIfNotNil("\u004d", _dgaggf.M)
	_cbdef.SetIfNotNil("\u0052e\u0066\u0065\u0072\u0065\u006e\u0063e", _dgaggf.Reference)
	_cbdef.SetIfNotNil("\u0043h\u0061\u006e\u0067\u0065\u0073", _dgaggf.Changes)
	_cbdef.SetIfNotNil("C\u006f\u006e\u0074\u0061\u0063\u0074\u0049\u006e\u0066\u006f", _dgaggf.ContactInfo)
	return _faega
}

// ToGoTime returns the date in time.Time format.
func (_deba PdfDate) ToGoTime() _f.Time {
	_gdaaa := int(_deba._dbgccd*60*60 + _deba._ccfca*60)
	switch _deba._aggabc {
	case '-':
		_gdaaa = -_gdaaa
	case 'Z':
		_gdaaa = 0
	}
	_fcdgd := _e.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _deba._aggabc, _deba._dbgccd, _deba._ccfca)
	_ebec := _f.FixedZone(_fcdgd, _gdaaa)
	return _f.Date(int(_deba._fabd), _f.Month(_deba._fcdacf), int(_deba._gecdc), int(_deba._ebda), int(_deba._efba), int(_deba._fgddf), 0, _ebec)
}

// ToPdfObject implements interface PdfModel.
func (_cdbg *PdfAnnotationScreen) ToPdfObject() _abf.PdfObject {
	_cdbg.PdfAnnotation.ToPdfObject()
	_cgcc := _cdbg._dbc
	_afbde := _cgcc.PdfObject.(*_abf.PdfObjectDictionary)
	_afbde.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0053\u0063\u0072\u0065\u0065\u006e"))
	_afbde.SetIfNotNil("\u0054", _cdbg.T)
	_afbde.SetIfNotNil("\u004d\u004b", _cdbg.MK)
	_afbde.SetIfNotNil("\u0041", _cdbg.A)
	_afbde.SetIfNotNil("\u0041\u0041", _cdbg.AA)
	return _cgcc
}

// NewPdfAnnotationStrikeOut returns a new text strikeout annotation.
func NewPdfAnnotationStrikeOut() *PdfAnnotationStrikeOut {
	_gdc := NewPdfAnnotation()
	_ccd := &PdfAnnotationStrikeOut{}
	_ccd.PdfAnnotation = _gdc
	_ccd.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gdc.SetContext(_ccd)
	return _ccd
}

// ToPdfObject returns the text field dictionary within an indirect object (container).
func (_cdbbe *PdfFieldText) ToPdfObject() _abf.PdfObject {
	_cdbbe.PdfField.ToPdfObject()
	_edae := _cdbbe._dgdc
	_bdgd := _edae.PdfObject.(*_abf.PdfObjectDictionary)
	_bdgd.Set("\u0046\u0054", _abf.MakeName("\u0054\u0078"))
	if _cdbbe.DA != nil {
		_bdgd.Set("\u0044\u0041", _cdbbe.DA)
	}
	if _cdbbe.Q != nil {
		_bdgd.Set("\u0051", _cdbbe.Q)
	}
	if _cdbbe.DS != nil {
		_bdgd.Set("\u0044\u0053", _cdbbe.DS)
	}
	if _cdbbe.RV != nil {
		_bdgd.Set("\u0052\u0056", _cdbbe.RV)
	}
	if _cdbbe.MaxLen != nil {
		_bdgd.Set("\u004d\u0061\u0078\u004c\u0065\u006e", _cdbbe.MaxLen)
	}
	return _edae
}

// ToPdfObject implements interface PdfModel.
func (_cfaf *PdfAnnotationPolyLine) ToPdfObject() _abf.PdfObject {
	_cfaf.PdfAnnotation.ToPdfObject()
	_eag := _cfaf._dbc
	_bcedd := _eag.PdfObject.(*_abf.PdfObjectDictionary)
	_cfaf.PdfAnnotationMarkup.appendToPdfDictionary(_bcedd)
	_bcedd.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u0050\u006f\u006c\u0079\u004c\u0069\u006e\u0065"))
	_bcedd.SetIfNotNil("\u0056\u0065\u0072\u0074\u0069\u0063\u0065\u0073", _cfaf.Vertices)
	_bcedd.SetIfNotNil("\u004c\u0045", _cfaf.LE)
	_bcedd.SetIfNotNil("\u0042\u0053", _cfaf.BS)
	_bcedd.SetIfNotNil("\u0049\u0043", _cfaf.IC)
	_bcedd.SetIfNotNil("\u0042\u0045", _cfaf.BE)
	_bcedd.SetIfNotNil("\u0049\u0054", _cfaf.IT)
	_bcedd.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _cfaf.Measure)
	return _eag
}

// ToPdfObject implements interface PdfModel.
func (_egb *PdfActionThread) ToPdfObject() _abf.PdfObject {
	_egb.PdfAction.ToPdfObject()
	_afg := _egb._egg
	_dece := _afg.PdfObject.(*_abf.PdfObjectDictionary)
	_dece.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeThread)))
	if _egb.F != nil {
		_dece.Set("\u0046", _egb.F.ToPdfObject())
	}
	_dece.SetIfNotNil("\u0044", _egb.D)
	_dece.SetIfNotNil("\u0042", _egb.B)
	return _afg
}

// HasColorspaceByName checks if the colorspace with the specified name exists in the page resources.
func (_gagfb *PdfPageResources) HasColorspaceByName(keyName _abf.PdfObjectName) bool {
	_ecae, _edba := _gagfb.GetColorspaces()
	if _edba != nil {
		_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0072\u0061\u0063\u0065: \u0025\u0076", _edba)
		return false
	}
	if _ecae == nil {
		return false
	}
	_, _bacgcd := _ecae.Colorspaces[string(keyName)]
	return _bacgcd
}

func (_beed *PdfReader) newPdfAnnotationSoundFromDict(_ecc *_abf.PdfObjectDictionary) (*PdfAnnotationSound, error) {
	_cagb := PdfAnnotationSound{}
	_dgga, _dgaf := _beed.newPdfAnnotationMarkupFromDict(_ecc)
	if _dgaf != nil {
		return nil, _dgaf
	}
	_cagb.PdfAnnotationMarkup = _dgga
	_cagb.Name = _ecc.Get("\u004e\u0061\u006d\u0065")
	_cagb.Sound = _ecc.Get("\u0053\u006f\u0075n\u0064")
	return &_cagb, nil
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
	Metadata *_abf.PdfObjectStream
	Data     []byte
	_afcc    *_abf.PdfIndirectObject
	_bfgc    *_abf.PdfObjectStream
}

// SetNamedDestinations sets the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_becgc *PdfWriter) SetNamedDestinations(dests _abf.PdfObject) error {
	if dests == nil {
		return nil
	}
	_acd.Log.Trace("\u0053e\u0074\u0074\u0069\u006e\u0067\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0044\u0065\u0073\u0074\u0073\u002e\u002e\u002e")
	_becgc._ddffc.Set("\u0044\u0065\u0073t\u0073", dests)
	return _becgc.addObjects(dests)
}

// ImageToGray returns a new grayscale image based on the passed in RGB image.
func (_edgf *PdfColorspaceDeviceRGB) ImageToGray(img Image) (Image, error) {
	if img.ColorComponents != 3 {
		return img, _fd.New("\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u0020\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042")
	}
	_adeb, _gfffe := _gca.NewImage(int(img.Width), int(img.Height), int(img.BitsPerComponent), img.ColorComponents, img.Data, img._gedg, img._ceeag)
	if _gfffe != nil {
		return img, _gfffe
	}
	_cggfc, _gfffe := _gca.GrayConverter.Convert(_adeb)
	if _gfffe != nil {
		return img, _gfffe
	}
	return _cega(_cggfc.Base()), nil
}
func (_cbeb *pdfCIDFontType0) getFontDescriptor() *PdfFontDescriptor { return _cbeb._dcbaf }

// NewStandard14FontWithEncoding returns the standard 14 font named `basefont` as a *PdfFont and
// a TextEncoder that encodes all the runes in `alphabet`, or an error if this is not possible.
// An error can occur if `basefont` is not one the standard 14 font names.
func NewStandard14FontWithEncoding(basefont StdFontName, alphabet map[rune]int) (*PdfFont, _cbb.SimpleEncoder, error) {
	_bffe, _agebb := _bfabe(basefont)
	if _agebb != nil {
		return nil, nil, _agebb
	}
	_fbdc, _egag := _bffe.Encoder().(_cbb.SimpleEncoder)
	if !_egag {
		return nil, nil, _e.Errorf("\u006f\u006e\u006c\u0079\u0020s\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u002c\u0020\u0067\u006f\u0074\u0020\u0025\u0054", _bffe.Encoder())
	}
	_ecged := make(map[rune]_cbb.GlyphName)
	for _cfbbc := range alphabet {
		if _, _gddb := _fbdc.RuneToCharcode(_cfbbc); !_gddb {
			_, _fece := _bffe._aecd.Read(_cfbbc)
			if !_fece {
				_acd.Log.Trace("r\u0075\u006e\u0065\u0020\u0025\u0023x\u003d\u0025\u0071\u0020\u006e\u006f\u0074\u0020\u0069n\u0020\u0074\u0068e\u0020f\u006f\u006e\u0074", _cfbbc, _cfbbc)
				continue
			}
			_bagb, _fece := _cbb.RuneToGlyph(_cfbbc)
			if !_fece {
				_acd.Log.Debug("\u006eo\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0066\u006f\u0072\u0020r\u0075\u006e\u0065\u0020\u0025\u0023\u0078\u003d\u0025\u0071", _cfbbc, _cfbbc)
				continue
			}
			if len(_ecged) >= 255 {
				return nil, nil, _fd.New("\u0074\u006f\u006f\u0020\u006d\u0061\u006e\u0079\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020f\u006f\u0072\u0020\u0073\u0069\u006d\u0070\u006c\u0065\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067")
			}
			_ecged[_cfbbc] = _bagb
		}
	}
	var (
		_eebd  []_cbb.CharCode
		_gadee []_cbb.CharCode
	)
	for _agdd := _cbb.CharCode(1); _agdd <= 0xff; _agdd++ {
		_ccfed, _bfcdd := _fbdc.CharcodeToRune(_agdd)
		if !_bfcdd {
			_eebd = append(_eebd, _agdd)
			continue
		}
		if _, _bfcdd = alphabet[_ccfed]; !_bfcdd {
			_gadee = append(_gadee, _agdd)
		}
	}
	_afae := append(_eebd, _gadee...)
	if len(_afae) < len(_ecged) {
		return nil, nil, _e.Errorf("n\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0065\u006ec\u006f\u0064\u0065\u0020\u0025\u0064\u0020ru\u006e\u0065\u0073\u002c \u0062\u0075\u0074\u0020\u0068\u0061\u0076\u0065\u0020on\u006c\u0079 \u0025\u0064\u0020\u0073\u006c\u006f\u0074\u0073", len(_ecged), len(_afae))
	}
	_acfbd := make([]rune, 0, len(_ecged))
	for _gaege := range _ecged {
		_acfbd = append(_acfbd, _gaege)
	}
	_bb.Slice(_acfbd, func(_feed, _bcgd int) bool { return _acfbd[_feed] < _acfbd[_bcgd] })
	_eacdb := make(map[_cbb.CharCode]_cbb.GlyphName, len(_acfbd))
	for _, _baeg := range _acfbd {
		_fedad := _afae[0]
		_afae = _afae[1:]
		_eacdb[_fedad] = _ecged[_baeg]
	}
	_fbdc = _cbb.ApplyDifferences(_fbdc, _eacdb)
	_bffe.SetEncoder(_fbdc)
	return &PdfFont{_gedca: &_bffe}, _fbdc, nil
}

// ToPdfObject implements interface PdfModel.
func (_fc *PdfAction) ToPdfObject() _abf.PdfObject {
	_fgb := _fc._egg
	_gbg := _fgb.PdfObject.(*_abf.PdfObjectDictionary)
	_gbg.Clear()
	_gbg.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0041\u0063\u0074\u0069\u006f\u006e"))
	_gbg.SetIfNotNil("\u0053", _fc.S)
	_gbg.SetIfNotNil("\u004e\u0065\u0078\u0074", _fc.Next)
	return _fgb
}

// AcroFormNeedsRepair returns true if the document contains widget annotations
// linked to fields which are not referenced in the AcroForm. The AcroForm can
// be repaired using the RepairAcroForm method of the reader.
func (_edebf *PdfReader) AcroFormNeedsRepair() (bool, error) {
	var _afdff []*PdfField
	if _edebf.AcroForm != nil {
		_afdff = _edebf.AcroForm.AllFields()
	}
	_agdg := make(map[*PdfField]struct{}, len(_afdff))
	for _, _gcccce := range _afdff {
		_agdg[_gcccce] = struct{}{}
	}
	for _, _gagac := range _edebf.PageList {
		_ecdcc, _bdccf := _gagac.GetAnnotations()
		if _bdccf != nil {
			return false, _bdccf
		}
		for _, _eceee := range _ecdcc {
			_decb, _eadf := _eceee.GetContext().(*PdfAnnotationWidget)
			if !_eadf {
				continue
			}
			_adcgb := _decb.Field()
			if _adcgb == nil {
				return true, nil
			}
			if _, _gegcf := _agdg[_adcgb]; !_gegcf {
				return true, nil
			}
		}
	}
	return false, nil
}

// GetNumComponents returns the number of color components (4 for CMYK32).
func (_bbac *PdfColorDeviceCMYK) GetNumComponents() int { return 4 }

// NewPdfReaderFromFile creates a new PdfReader from the speficied PDF file.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderFromFile(pdfFile string, opts *ReaderOpts) (*PdfReader, *_cf.File, error) {
	const _babb = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0046\u0072\u006f\u006dF\u0069\u006c\u0065"
	_afbc, _fcac := _cf.Open(pdfFile)
	if _fcac != nil {
		return nil, nil, _fcac
	}
	_bgacc, _fcac := _fbaec(_afbc, opts, true, _babb)
	if _fcac != nil {
		_afbc.Close()
		return nil, nil, _fcac
	}
	return _bgacc, _afbc, nil
}

// ParsePdfObject parses input pdf object into given output intent.
func (_gafd *PdfOutputIntent) ParsePdfObject(object _abf.PdfObject) error {
	_agef, _dgac := _abf.GetDict(object)
	if !_dgac {
		_acd.Log.Error("\u0055\u006e\u006bno\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020%\u0054 \u0066o\u0072 \u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074", object)
		return _fd.New("\u0075\u006e\u006b\u006e\u006fw\u006e\u0020\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0066\u006f\u0072\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006e\u0074")
	}
	_gafd._dcfb = _agef
	_gafd.Type, _ = _agef.GetString("\u0054\u0079\u0070\u0065")
	_gaebc, _dgac := _agef.GetString("\u0053")
	if _dgac {
		switch _gaebc {
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411":
			_gafd.S = PdfOutputIntentTypeA1
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00412":
			_gafd.S = PdfOutputIntentTypeA2
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00413":
			_gafd.S = PdfOutputIntentTypeA3
		case "\u0047T\u0053\u005f\u0050\u0044\u0046\u00414":
			_gafd.S = PdfOutputIntentTypeA4
		case "\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0058":
			_gafd.S = PdfOutputIntentTypeX
		}
	}
	_gafd.OutputCondition, _ = _agef.GetString("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e")
	_gafd.OutputConditionIdentifier, _ = _agef.GetString("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r")
	_gafd.RegistryName, _ = _agef.GetString("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065")
	_gafd.Info, _ = _agef.GetString("\u0049\u006e\u0066\u006f")
	if _feede, _bffdg := _abf.GetStream(_agef.Get("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065")); _bffdg {
		_gafd.ColorComponents, _ = _abf.GetIntVal(_feede.Get("\u004e"))
		_efgfc, _bcgab := _abf.DecodeStream(_feede)
		if _bcgab != nil {
			return _bcgab
		}
		_gafd.DestOutputProfile = _efgfc
	}
	return nil
}

// NewImageFromGoImage creates a new NRGBA32 unidoc Image from a golang Image.
// If `goimg` is grayscale (*goimage.Gray8) then calls NewGrayImageFromGoImage instead.
func (_dcgad DefaultImageHandler) NewImageFromGoImage(goimg _aa.Image) (*Image, error) {
	_dabg, _aefga := _gca.FromGoImage(goimg)
	if _aefga != nil {
		return nil, _aefga
	}
	_geeade := _cega(_dabg.Base())
	return &_geeade, nil
}

// NewPdfSignatureReferenceDocMDP returns PdfSignatureReference for the transformParams.
func NewPdfSignatureReferenceDocMDP(transformParams *PdfTransformParamsDocMDP) *PdfSignatureReference {
	return &PdfSignatureReference{Type: _abf.MakeName("\u0053\u0069\u0067\u0052\u0065\u0066"), TransformMethod: _abf.MakeName("\u0044\u006f\u0063\u004d\u0044\u0050"), TransformParams: transformParams.ToPdfObject()}
}

// NewPdfOutline returns an initialized PdfOutline.
func NewPdfOutline() *PdfOutline {
	_faffc := &PdfOutline{_cgcg: _abf.MakeIndirectObject(_abf.MakeDict())}
	_faffc._aecec = _faffc
	return _faffc
}

// SetSamples convert samples to byte-data and sets for the image.
// NOTE: The method resamples the data and this could lead to high memory usage,
// especially on large images. It should be used only when it is not possible
// to work with the image byte data directly.
func (_bdcde *Image) SetSamples(samples []uint32) {
	if _bdcde.BitsPerComponent < 8 {
		samples = _bdcde.samplesAddPadding(samples)
	}
	_ecggd := _gf.ResampleUint32(samples, int(_bdcde.BitsPerComponent), 8)
	_ebab := make([]byte, len(_ecggd))
	for _fbbfd, _abeg := range _ecggd {
		_ebab[_fbbfd] = byte(_abeg)
	}
	_bdcde.Data = _ebab
}

// NewPdfAnnotationFreeText returns a new free text annotation.
func NewPdfAnnotationFreeText() *PdfAnnotationFreeText {
	_acb := NewPdfAnnotation()
	_aec := &PdfAnnotationFreeText{}
	_aec.PdfAnnotation = _acb
	_aec.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_acb.SetContext(_aec)
	return _aec
}

// NewPdfDate returns a new PdfDate object from a PDF date string (see 7.9.4 Dates).
// format: "D: YYYYMMDDHHmmSSOHH'mm"
func NewPdfDate(dateStr string) (PdfDate, error) {
	_edddf, _dcdec := _fae.ParsePdfTime(dateStr)
	if _dcdec != nil {
		return PdfDate{}, _dcdec
	}
	return NewPdfDateFromTime(_edddf)
}

// GetParamsDict returns *core.PdfObjectDictionary with a set of basic image parameters.
func (_bace *Image) GetParamsDict() *_abf.PdfObjectDictionary {
	_agbgc := _abf.MakeDict()
	_agbgc.Set("\u0057\u0069\u0064t\u0068", _abf.MakeInteger(_bace.Width))
	_agbgc.Set("\u0048\u0065\u0069\u0067\u0068\u0074", _abf.MakeInteger(_bace.Height))
	_agbgc.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _abf.MakeInteger(int64(_bace.ColorComponents)))
	_agbgc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _abf.MakeInteger(_bace.BitsPerComponent))
	return _agbgc
}

const (
	BorderEffectNoEffect BorderEffect = iota
	BorderEffectCloudy   BorderEffect = iota
)

func (_eabge *PdfWriter) writeOutlines() error {
	if _eabge._gbcge == nil {
		return nil
	}
	_acd.Log.Trace("\u004f\u0075t\u006c\u0069\u006ee\u0054\u0072\u0065\u0065\u003a\u0020\u0025\u002b\u0076", _eabge._gbcge)
	_ggbefd := _eabge._gbcge.ToPdfObject()
	_acd.Log.Trace("\u004fu\u0074\u006c\u0069\u006e\u0065\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0025\u0054\u002c\u0020\u0070\u003a\u0025\u0070\u0029", _ggbefd, _ggbefd, _ggbefd)
	_eabge._ddffc.Set("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073", _ggbefd)
	_bccba := _eabge.addObjects(_ggbefd)
	if _bccba != nil {
		return _bccba
	}
	return nil
}
func _edcbb() string { _gaabd.Lock(); defer _gaabd.Unlock(); return _eabe }

// ToPdfObject returns the PDF representation of the page resources.
func (_bcegc *PdfPageResources) ToPdfObject() _abf.PdfObject {
	_dbgcc := _bcegc._gagb
	_dbgcc.SetIfNotNil("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _bcegc.ExtGState)
	if _bcegc._aafff != nil {
		_bcegc.ColorSpace = _bcegc._aafff.ToPdfObject()
	}
	_dbgcc.SetIfNotNil("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065", _bcegc.ColorSpace)
	_dbgcc.SetIfNotNil("\u0050a\u0074\u0074\u0065\u0072\u006e", _bcegc.Pattern)
	_dbgcc.SetIfNotNil("\u0053h\u0061\u0064\u0069\u006e\u0067", _bcegc.Shading)
	_dbgcc.SetIfNotNil("\u0058O\u0062\u006a\u0065\u0063\u0074", _bcegc.XObject)
	_dbgcc.SetIfNotNil("\u0046\u006f\u006e\u0074", _bcegc.Font)
	_dbgcc.SetIfNotNil("\u0050r\u006f\u0063\u0053\u0065\u0074", _bcegc.ProcSet)
	_dbgcc.SetIfNotNil("\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073", _bcegc.Properties)
	return _dbgcc
}

func _bage(_ccfge *PdfField, _gfefc _abf.PdfObject) error {
	switch _ccfge.GetContext().(type) {
	case *PdfFieldText:
		switch _fbega := _gfefc.(type) {
		case *_abf.PdfObjectName:
			_eefbeb := _fbega
			_acd.Log.Debug("\u0055\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u003a\u0020\u0047\u006f\u0074 \u0056\u0020\u0061\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u003e\u0020c\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0074\u006f s\u0074\u0072\u0069\u006e\u0067\u0020\u0027\u0025\u0073\u0027", _eefbeb.String())
			_ccfge.V = _abf.MakeEncodedString(_fbega.String(), true)
		case *_abf.PdfObjectString:
			_ccfge.V = _abf.MakeEncodedString(_fbega.String(), true)
		default:
			_acd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0056\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u0028\u0025\u0023\u0076\u0029", _fbega, _fbega)
		}
	case *PdfFieldButton:
		switch _gfefc.(type) {
		case *_abf.PdfObjectName:
			if len(_gfefc.String()) > 0 {
				_ccfge.V = _gfefc
				_bdda(_ccfge, _gfefc)
			}
		case *_abf.PdfObjectString:
			if len(_gfefc.String()) > 0 {
				_ccfge.V = _abf.MakeName(_gfefc.String())
				_bdda(_ccfge, _ccfge.V)
			}
		default:
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _ccfge.PartialName(), _gfefc)
			_ccfge.V = _gfefc
		}
	case *PdfFieldChoice:
		switch _gfefc.(type) {
		case *_abf.PdfObjectName:
			if len(_gfefc.String()) > 0 {
				_ccfge.V = _abf.MakeString(_gfefc.String())
				_bdda(_ccfge, _gfefc)
			}
		case *_abf.PdfObjectString:
			if len(_gfefc.String()) > 0 {
				_ccfge.V = _gfefc
				_bdda(_ccfge, _abf.MakeName(_gfefc.String()))
			}
		default:
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u004e\u0045\u0058P\u0045\u0043\u0054\u0045\u0044\u0020\u0025\u0073\u0020\u002d>\u0020\u0025\u0076", _ccfge.PartialName(), _gfefc)
			_ccfge.V = _gfefc
		}
	case *PdfFieldSignature:
		_acd.Log.Debug("\u0054\u004f\u0044\u004f\u003a \u0053\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0061\u0070\u0070e\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0079\u0065\u0074\u003a\u0020\u0025\u0073\u002f\u0025v", _ccfge.PartialName(), _gfefc)
	}
	return nil
}

// Register registers (caches) a model to primitive object relationship.
func (_afbf *modelManager) Register(primitive _abf.PdfObject, model PdfModel) {
	_afbf._baecg[model] = primitive
	_afbf._addgc[primitive] = model
}

// GetCapHeight returns the CapHeight of the font `descriptor`.
func (_ebcce *PdfFontDescriptor) GetCapHeight() (float64, error) {
	return _abf.GetNumberAsFloat(_ebcce.CapHeight)
}

// Val returns the color value.
func (_bcdd *PdfColorDeviceGray) Val() float64 { return float64(*_bcdd) }

// GetContentStream returns the XObject Form's content stream.
func (_fdaed *XObjectForm) GetContentStream() ([]byte, error) {
	_gbeae, _gegda := _abf.DecodeStream(_fdaed._dbba)
	if _gegda != nil {
		return nil, _gegda
	}
	return _gbeae, nil
}

// XObjectImage (Table 89 in 8.9.5.1).
// Implements PdfModel interface.
type XObjectImage struct {
	// ColorSpace       PdfObject
	Width            *int64
	Height           *int64
	ColorSpace       PdfColorspace
	BitsPerComponent *int64
	Filter           _abf.StreamEncoder
	Intent           _abf.PdfObject
	ImageMask        _abf.PdfObject
	Mask             _abf.PdfObject
	Matte            _abf.PdfObject
	Decode           _abf.PdfObject
	Interpolate      _abf.PdfObject
	Alternatives     _abf.PdfObject
	SMask            _abf.PdfObject
	SMaskInData      _abf.PdfObject
	Name             _abf.PdfObject
	StructParent     _abf.PdfObject
	ID               _abf.PdfObject
	OPI              _abf.PdfObject
	Metadata         _abf.PdfObject
	OC               _abf.PdfObject
	Stream           []byte
	_ccbad           *_abf.PdfObjectStream
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_faac pdfCIDFontType0) GetCharMetrics(code _cbb.CharCode) (_gbe.CharMetrics, bool) {
	_gbfa := _faac._bdced
	if _dagc, _afcac := _faac._fbcfb[code]; _afcac {
		_gbfa = _dagc
	}
	return _gbe.CharMetrics{Wx: _gbfa}, true
}

// NewPdfAppender creates a new Pdf appender from a Pdf reader.
func NewPdfAppender(reader *PdfReader) (*PdfAppender, error) {
	_gaee := &PdfAppender{_eeded: reader._affbb, Reader: reader, _bdcd: reader._bebc, _gfeg: reader._ggbccc}
	_aeff, _fbgf := _gaee._eeded.Seek(0, _gc.SeekEnd)
	if _fbgf != nil {
		return nil, _fbgf
	}
	_gaee._cfga = _aeff
	if _, _fbgf = _gaee._eeded.Seek(0, _gc.SeekStart); _fbgf != nil {
		return nil, _fbgf
	}
	_gaee._agda, _fbgf = NewPdfReader(_gaee._eeded)
	if _fbgf != nil {
		return nil, _fbgf
	}
	for _, _gbaf := range _gaee.Reader.GetObjectNums() {
		if _gaee._ffc < _gbaf {
			_gaee._ffc = _gbaf
		}
	}
	_gaee._abce = _gaee._bdcd.GetXrefTable()
	_gaee._dac = _gaee._bdcd.GetXrefOffset()
	_gaee._cggfa = append(_gaee._cggfa, _gaee._agda.PageList...)
	_gaee._gcba = make(map[_abf.PdfObject]struct{})
	_gaee._bge = make(map[_abf.PdfObject]int64)
	_gaee._cdbbg = make(map[_abf.PdfObject]struct{})
	_gaee._ffbb = _gaee._agda.AcroForm
	_gaee._ffbe = _gaee._agda.DSS
	return _gaee, nil
}

// GetContainingPdfObject returns the XObject Form's containing object (indirect object).
func (_fbdde *XObjectForm) GetContainingPdfObject() _abf.PdfObject { return _fbdde._dbba }

// ToPdfObject implements interface PdfModel.
func (_eaad *PdfActionNamed) ToPdfObject() _abf.PdfObject {
	_eaad.PdfAction.ToPdfObject()
	_adc := _eaad._egg
	_edbd := _adc.PdfObject.(*_abf.PdfObjectDictionary)
	_edbd.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeNamed)))
	_edbd.SetIfNotNil("\u004e", _eaad.N)
	return _adc
}

// PdfFieldChoice represents a choice field which includes scrollable list boxes and combo boxes.
type PdfFieldChoice struct {
	*PdfField
	Opt *_abf.PdfObjectArray
	TI  *_abf.PdfObjectInteger
	I   *_abf.PdfObjectArray
}

// Fill populates `form` with values provided by `provider`.
func (_aeda *PdfAcroForm) Fill(provider FieldValueProvider) error { return _aeda.fill(provider, nil) }

func _fbaec(_eadbb _gc.ReadSeeker, _ffee *ReaderOpts, _gceed bool, _defad string) (*PdfReader, error) {
	if _ffee == nil {
		_ffee = NewReaderOpts()
	}
	_cdefbd := *_ffee
	_gagaa := &PdfReader{_affbb: _eadbb, _ggbccc: map[_abf.PdfObject]struct{}{}, _ceecd: _gadf(), _abgge: _ffee.LazyLoad, _dfafc: _ffee.ComplianceMode, _dbgdg: _gceed, _gebfg: &_cdefbd}
	_gffcb, _gdfg := _addec("\u0072")
	if _gdfg != nil {
		return nil, _gdfg
	}
	_gagaa._bccga = _gffcb
	var _babeg *_abf.PdfParser
	if !_gagaa._dfafc {
		_babeg, _gdfg = _abf.NewParser(_eadbb)
	} else {
		_babeg, _gdfg = _abf.NewCompliancePdfParser(_eadbb)
	}
	if _gdfg != nil {
		return nil, _gdfg
	}
	_gagaa._bebc = _babeg
	_edaaed, _gdfg := _gagaa.IsEncrypted()
	if _gdfg != nil {
		return nil, _gdfg
	}
	if !_edaaed {
		_gdfg = _gagaa.loadStructure()
		if _gdfg != nil {
			return nil, _gdfg
		}
	} else if _gceed {
		_bfba, _aaffbb := _gagaa.Decrypt([]byte(_ffee.Password))
		if _aaffbb != nil {
			return nil, _aaffbb
		}
		if !_bfba {
			return nil, _fd.New("\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0064\u0065c\u0072\u0079\u0070\u0074\u0020\u0070\u0061\u0073\u0073w\u006f\u0072\u0064\u0020p\u0072\u006f\u0074\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u002d\u0020\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0073\u0070\u0065\u0063\u0069\u0066y\u0020\u0070\u0061s\u0073\u0020\u0074\u006f\u0020\u0044\u0065\u0063\u0072\u0079\u0070\u0074")
		}
	}
	_gagaa._bfced = make(map[*PdfReader]*PdfReader)
	_gagaa._egade = make([]*PdfReader, _babeg.GetRevisionNumber())
	return _gagaa, nil
}

// SetDocInfo set document info.
// This will overwrite any globally declared document info.
func (_fbeb *PdfWriter) SetDocInfo(info *PdfInfo) { _fbeb.setDocInfo(info.ToPdfObject()) }

// BaseFont returns the font's "BaseFont" field.
func (_fdg *PdfFont) BaseFont() string { return _fdg.baseFields()._ecggf }

// PdfAnnotationWatermark represents Watermark annotations.
// (Section 12.5.6.22).
type PdfAnnotationWatermark struct {
	*PdfAnnotation
	FixedPrint _abf.PdfObject
}

// NewStandardPdfOutputIntent creates a new standard PdfOutputIntent.
func NewStandardPdfOutputIntent(outputCondition, outputConditionIdentifier, registryName string, destOutputProfile []byte, colorComponents int) *PdfOutputIntent {
	return &PdfOutputIntent{Type: "\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074", OutputCondition: outputCondition, OutputConditionIdentifier: outputConditionIdentifier, RegistryName: registryName, DestOutputProfile: destOutputProfile, ColorComponents: colorComponents, _dcfb: _abf.MakeDict()}
}

// ToPdfObject recursively builds the Outline tree PDF object.
func (_efgce *PdfOutline) ToPdfObject() _abf.PdfObject {
	_egefa := _efgce._cgcg
	_cdcc := _egefa.PdfObject.(*_abf.PdfObjectDictionary)
	_cdcc.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073"))
	if _efgce.First != nil {
		_cdcc.Set("\u0046\u0069\u0072s\u0074", _efgce.First.ToPdfObject())
	}
	if _efgce.Last != nil {
		_cdcc.Set("\u004c\u0061\u0073\u0074", _efgce.Last.GetContext().GetContainingPdfObject())
	}
	if _efgce.Parent != nil {
		_cdcc.Set("\u0050\u0061\u0072\u0065\u006e\u0074", _efgce.Parent.GetContext().GetContainingPdfObject())
	}
	if _efgce.Count != nil {
		_cdcc.Set("\u0043\u006f\u0075n\u0074", _abf.MakeInteger(*_efgce.Count))
	}
	return _egefa
}

// PdfFilespec represents a file specification which can either refer to an external or embedded file.
type PdfFilespec struct {
	Type   _abf.PdfObject
	FS     _abf.PdfObject
	F      _abf.PdfObject
	UF     _abf.PdfObject
	DOS    _abf.PdfObject
	Mac    _abf.PdfObject
	Unix   _abf.PdfObject
	ID     _abf.PdfObject
	V      _abf.PdfObject
	EF     _abf.PdfObject
	RF     _abf.PdfObject
	Desc   _abf.PdfObject
	CI     _abf.PdfObject
	_badbg _abf.PdfObject
}

// PdfShadingType3 is a Radial shading.
type PdfShadingType3 struct {
	*PdfShading
	Coords   *_abf.PdfObjectArray
	Domain   *_abf.PdfObjectArray
	Function []PdfFunction
	Extend   *_abf.PdfObjectArray
}

func _gabdad(_bdeaa *_abf.PdfIndirectObject) (*PdfOutline, error) {
	_debdf, _ebff := _bdeaa.PdfObject.(*_abf.PdfObjectDictionary)
	if !_ebff {
		return nil, _e.Errorf("\u006f\u0075\u0074l\u0069\u006e\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
	}
	_dfbccb := NewPdfOutline()
	if _fagc := _debdf.Get("\u0054\u0079\u0070\u0065"); _fagc != nil {
		_abecg, _befdf := _fagc.(*_abf.PdfObjectName)
		if _befdf {
			if *_abecg != "\u004f\u0075\u0074\u006c\u0069\u006e\u0065\u0073" {
				_acd.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0054y\u0070\u0065\u0020\u0021\u003d\u0020\u004f\u0075\u0074l\u0069\u006e\u0065s\u0020(\u0025\u0073\u0029", *_abecg)
			}
		}
	}
	if _gbcac := _debdf.Get("\u0043\u006f\u0075n\u0074"); _gbcac != nil {
		_adgfg, _afgbb := _abf.GetNumberAsInt64(_gbcac)
		if _afgbb != nil {
			return nil, _afgbb
		}
		_dfbccb.Count = &_adgfg
	}
	return _dfbccb, nil
}

// PdfActionNamed represents a named action.
type PdfActionNamed struct {
	*PdfAction
	N _abf.PdfObject
}

// NewBorderStyle returns an initialized PdfBorderStyle.
func NewBorderStyle() *PdfBorderStyle { _bgbc := &PdfBorderStyle{}; return _bgbc }

func (_bbbb *PdfWriter) writeObjects() {
	_acd.Log.Trace("\u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0025d\u0020\u006f\u0062\u006a", len(_bbbb._edcgc))
	_bbbb._becfc = make(map[int]crossReference)
	_bbbb._becfc[0] = crossReference{Type: 0, ObjectNumber: 0, Generation: 0xFFFF}
	if _bbbb._cagaf.ObjectMap != nil {
		for _egcec, _cbefa := range _bbbb._cagaf.ObjectMap {
			if _egcec == 0 {
				continue
			}
			if _cbefa.XType == _abf.XrefTypeObjectStream {
				_gfdbe := crossReference{Type: 2, ObjectNumber: _cbefa.OsObjNumber, Index: _cbefa.OsObjIndex}
				_bbbb._becfc[_egcec] = _gfdbe
			}
			if _cbefa.XType == _abf.XrefTypeTableEntry {
				_dbgda := crossReference{Type: 1, ObjectNumber: _cbefa.ObjectNumber, Offset: _cbefa.Offset}
				_bbbb._becfc[_egcec] = _dbgda
			}
		}
	}
}

// ToPdfObject implements interface PdfModel.
func (_gefd *PdfAnnotationLine) ToPdfObject() _abf.PdfObject {
	_gefd.PdfAnnotation.ToPdfObject()
	_dgafe := _gefd._dbc
	_eebc := _dgafe.PdfObject.(*_abf.PdfObjectDictionary)
	_gefd.PdfAnnotationMarkup.appendToPdfDictionary(_eebc)
	_eebc.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u004c\u0069\u006e\u0065"))
	_eebc.SetIfNotNil("\u004c", _gefd.L)
	_eebc.SetIfNotNil("\u0042\u0053", _gefd.BS)
	_eebc.SetIfNotNil("\u004c\u0045", _gefd.LE)
	_eebc.SetIfNotNil("\u0049\u0043", _gefd.IC)
	_eebc.SetIfNotNil("\u004c\u004c", _gefd.LL)
	_eebc.SetIfNotNil("\u004c\u004c\u0045", _gefd.LLE)
	_eebc.SetIfNotNil("\u0043\u0061\u0070", _gefd.Cap)
	_eebc.SetIfNotNil("\u0049\u0054", _gefd.IT)
	_eebc.SetIfNotNil("\u004c\u004c\u004f", _gefd.LLO)
	_eebc.SetIfNotNil("\u0043\u0050", _gefd.CP)
	_eebc.SetIfNotNil("\u004de\u0061\u0073\u0075\u0072\u0065", _gefd.Measure)
	_eebc.SetIfNotNil("\u0043\u004f", _gefd.CO)
	return _dgafe
}

func (_afdb *PdfReader) flattenFieldsWithOpts(_babc bool, _bgdf FieldAppearanceGenerator, _gdaba *FieldFlattenOpts) error {
	if _gdaba == nil {
		_gdaba = &FieldFlattenOpts{}
	}
	var _fdcb bool
	_fdfeg := map[*PdfAnnotation]bool{}
	{
		var _dcebf []*PdfField
		_fbcfd := _afdb.AcroForm
		if _fbcfd != nil {
			if _gdaba.FilterFunc != nil {
				_dcebf = _fbcfd.filteredFields(_gdaba.FilterFunc, true)
				_fdcb = _fbcfd.Fields != nil && len(*_fbcfd.Fields) > 0
			} else {
				_dcebf = _fbcfd.AllFields()
			}
		}
		for _, _edfae := range _dcebf {
			if len(_edfae.Annotations) < 1 {
				_acd.Log.Debug("\u004e\u006f\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u006f\u0075\u006ed\u0020\u0066\u006f\u0072\u003a\u0020\u0025v\u002c\u0020\u006c\u006f\u006f\u006b\u0020\u0069\u006e\u0074\u006f \u004b\u0069\u0064\u0073\u0020\u004f\u0062\u006a\u0065\u0063\u0074", _edfae.PartialName())
				for _egcg, _daddd := range _edfae.Kids {
					for _, _bfee := range _daddd.Annotations {
						_fdfeg[_bfee.PdfAnnotation] = _edfae.V != nil
						if _daddd.V == nil {
							_daddd.V = _edfae.V
						}
						if _daddd.T == nil {
							_daddd.T = _abf.MakeString(_e.Sprintf("\u0025\u0073\u0023%\u0064", _edfae.PartialName(), _egcg))
						}
						if _bgdf != nil {
							_cgee, _feeg := _bgdf.GenerateAppearanceDict(_fbcfd, _daddd, _bfee)
							if _feeg != nil {
								return _feeg
							}
							_bfee.AP = _cgee
						}
					}
				}
			}
			for _, _gcfe := range _edfae.Annotations {
				_fdfeg[_gcfe.PdfAnnotation] = _edfae.V != nil
				if _bgdf != nil {
					_gadgee, _beae := _bgdf.GenerateAppearanceDict(_fbcfd, _edfae, _gcfe)
					if _beae != nil {
						return _beae
					}
					_gcfe.AP = _gadgee
				}
			}
		}
	}
	if _babc {
		for _, _dfec := range _afdb.PageList {
			_egge, _ddedg := _dfec.GetAnnotations()
			if _ddedg != nil {
				return _ddedg
			}
			for _, _egefd := range _egge {
				_fdfeg[_egefd] = true
			}
		}
	}
	for _, _dfef := range _afdb.PageList {
		_cggbe := _dfef.flattenFieldsWithOpts(_bgdf, _gdaba, _fdfeg)
		if _cggbe != nil {
			return _cggbe
		}
	}
	if !_fdcb {
		_afdb.AcroForm = nil
	}
	return nil
}

// GetNamedDestinations returns the Dests entry in the PDF catalog.
// See section 12.3.2.3 "Named Destinations" (p. 367 PDF32000_2008).
func (_bgfcd *PdfReader) GetNamedDestinations() (_abf.PdfObject, error) {
	_begcc := _abf.ResolveReference(_bgfcd._dagde.Get("\u0044\u0065\u0073t\u0073"))
	if _begcc == nil {
		return nil, nil
	}
	if !_bgfcd._abgge {
		_fdfbf := _bgfcd.traverseObjectData(_begcc)
		if _fdfbf != nil {
			return nil, _fdfbf
		}
	}
	return _begcc, nil
}

func (_eeag *PdfWriter) updateObjectNumbers() {
	_abaf := _eeag.ObjNumOffset
	_gaaac := 0
	for _, _ebbdd := range _eeag._edcgc {
		_afcg := int64(_gaaac + 1 + _abaf)
		_fdcg := true
		if _eeag._aegbd {
			if _ddfcb, _dgea := _eeag._deff[_ebbdd]; _dgea {
				_afcg = _ddfcb
				_fdcg = false
			}
		}
		switch _cacdc := _ebbdd.(type) {
		case *_abf.PdfIndirectObject:
			_cacdc.ObjectNumber = _afcg
			_cacdc.GenerationNumber = 0
		case *_abf.PdfObjectStream:
			_cacdc.ObjectNumber = _afcg
			_cacdc.GenerationNumber = 0
		case *_abf.PdfObjectStreams:
			_cacdc.ObjectNumber = _afcg
			_cacdc.GenerationNumber = 0
		default:
			_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u0020%\u0054\u0020\u002d\u0020\u0073\u006b\u0069p\u0070\u0069\u006e\u0067", _cacdc)
			continue
		}
		if _fdcg {
			_gaaac++
		}
	}
	_fgedf := func(_edcag _abf.PdfObject) int64 {
		switch _defeb := _edcag.(type) {
		case *_abf.PdfIndirectObject:
			return _defeb.ObjectNumber
		case *_abf.PdfObjectStream:
			return _defeb.ObjectNumber
		case *_abf.PdfObjectStreams:
			return _defeb.ObjectNumber
		}
		return 0
	}
	_bb.SliceStable(_eeag._edcgc, func(_dggbe, _bddd int) bool { return _fgedf(_eeag._edcgc[_dggbe]) < _fgedf(_eeag._edcgc[_bddd]) })
}

func (_ffedc *PdfWriter) addObject(_gbcgee _abf.PdfObject) bool {
	_cbced := _ffedc.hasObject(_gbcgee)
	if !_cbced {
		_cfeee := _abf.ResolveReferencesDeep(_gbcgee, _ffedc._dbdcg)
		if _cfeee != nil {
			_acd.Log.Debug("E\u0052R\u004f\u0052\u003a\u0020\u0025\u0076\u0020\u002d \u0073\u006b\u0069\u0070pi\u006e\u0067", _cfeee)
		}
		_ffedc._edcgc = append(_ffedc._edcgc, _gbcgee)
		_ffedc._fdgae[_gbcgee] = struct{}{}
		return true
	}
	return false
}

// PdfOutputIntentType is the subtype of the given PdfOutputIntent.
type (
	PdfOutputIntentType int
	crossReference      struct {
		Type int

		// Type 1
		Offset     int64
		Generation int64

		// Type 2
		ObjectNumber int
		Index        int
	}
)

// ColorFromFloats returns a new PdfColor based on the input slice of color
// components.
func (_beede *PdfColorspaceICCBased) ColorFromFloats(vals []float64) (PdfColor, error) {
	if _beede.Alternate == nil {
		if _beede.N == 1 {
			_cabaf := NewPdfColorspaceDeviceGray()
			return _cabaf.ColorFromFloats(vals)
		} else if _beede.N == 3 {
			_cged := NewPdfColorspaceDeviceRGB()
			return _cged.ColorFromFloats(vals)
		} else if _beede.N == 4 {
			_fggf := NewPdfColorspaceDeviceCMYK()
			return _fggf.ColorFromFloats(vals)
		} else {
			return nil, _fd.New("I\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0061lt\u0065\u0072\u006ea\u0074i\u0076\u0065")
		}
	}
	return _beede.Alternate.ColorFromFloats(vals)
}

// XObjectType represents the type of an XObject.
type XObjectType int

func (_ggea *PdfReader) newPdfAnnotationRedactFromDict(_bggb *_abf.PdfObjectDictionary) (*PdfAnnotationRedact, error) {
	_dabb := PdfAnnotationRedact{}
	_fgab, _egf := _ggea.newPdfAnnotationMarkupFromDict(_bggb)
	if _egf != nil {
		return nil, _egf
	}
	_dabb.PdfAnnotationMarkup = _fgab
	_dabb.QuadPoints = _bggb.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	_dabb.IC = _bggb.Get("\u0049\u0043")
	_dabb.RO = _bggb.Get("\u0052\u004f")
	_dabb.OverlayText = _bggb.Get("O\u0076\u0065\u0072\u006c\u0061\u0079\u0054\u0065\u0078\u0074")
	_dabb.Repeat = _bggb.Get("\u0052\u0065\u0070\u0065\u0061\u0074")
	_dabb.DA = _bggb.Get("\u0044\u0041")
	_dabb.Q = _bggb.Get("\u0051")
	return &_dabb, nil
}

const (
	RC4_128bit = EncryptionAlgorithm(iota)
	AES_128bit
	AES_256bit
)

// SetContext sets the sub annotation (context).
func (_fga *PdfAnnotation) SetContext(ctx PdfModel)                  { _fga._edg = ctx }
func (_cced *pdfCIDFontType2) getFontDescriptor() *PdfFontDescriptor { return _cced._dcbaf }
func (_acde *PdfWriter) addObjects(_fcggf _abf.PdfObject) error {
	_acd.Log.Trace("\u0041d\u0064i\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073\u0021")
	if _bcfdc, _acdcg := _fcggf.(*_abf.PdfIndirectObject); _acdcg {
		_acd.Log.Trace("\u0049\u006e\u0064\u0069\u0072\u0065\u0063\u0074")
		_acd.Log.Trace("\u002d \u0025\u0073\u0020\u0028\u0025\u0070)", _fcggf, _bcfdc)
		_acd.Log.Trace("\u002d\u0020\u0025\u0073", _bcfdc.PdfObject)
		if _acde.addObject(_bcfdc) {
			_dfdde := _acde.addObjects(_bcfdc.PdfObject)
			if _dfdde != nil {
				return _dfdde
			}
		}
		return nil
	}
	if _efcd, _aegdb := _fcggf.(*_abf.PdfObjectStream); _aegdb {
		_acd.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d")
		_acd.Log.Trace("\u002d \u0025\u0073\u0020\u0025\u0070", _fcggf, _fcggf)
		if _acde.addObject(_efcd) {
			_daecgc := _acde.addObjects(_efcd.PdfObjectDictionary)
			if _daecgc != nil {
				return _daecgc
			}
		}
		return nil
	}
	if _dcaaa, _gdbdb := _fcggf.(*_abf.PdfObjectDictionary); _gdbdb {
		_acd.Log.Trace("\u0044\u0069\u0063\u0074")
		_acd.Log.Trace("\u002d\u0020\u0025\u0073", _fcggf)
		for _, _bgbgb := range _dcaaa.Keys() {
			_ecaeb := _dcaaa.Get(_bgbgb)
			if _gebaf, _fgac := _ecaeb.(*_abf.PdfObjectReference); _fgac {
				_ecaeb = _gebaf.Resolve()
				_dcaaa.Set(_bgbgb, _ecaeb)
			}
			if _bgbgb != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _gabec := _acde.addObjects(_ecaeb); _gabec != nil {
					return _gabec
				}
			} else {
				if _, _aafbb := _ecaeb.(*_abf.PdfObjectNull); _aafbb {
					continue
				}
				if _ebbcd := _acde.hasObject(_ecaeb); !_ebbcd {
					_acd.Log.Debug("P\u0061\u0072\u0065\u006e\u0074\u0020o\u0062\u006a\u0020\u006e\u006f\u0074 \u0061\u0064\u0064\u0065\u0064\u0020\u0079e\u0074\u0021\u0021\u0020\u0025\u0054\u0020\u0025\u0070\u0020%\u0076", _ecaeb, _ecaeb, _ecaeb)
					_acde._fadb[_ecaeb] = append(_acde._fadb[_ecaeb], _dcaaa)
				}
			}
		}
		return nil
	}
	if _cdeaf, _cagebd := _fcggf.(*_abf.PdfObjectArray); _cagebd {
		_acd.Log.Trace("\u0041\u0072\u0072a\u0079")
		_acd.Log.Trace("\u002d\u0020\u0025\u0073", _fcggf)
		if _cdeaf == nil {
			return _fd.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _bfgfc, _bgbgc := range _cdeaf.Elements() {
			if _afefd, _acccc := _bgbgc.(*_abf.PdfObjectReference); _acccc {
				_bgbgc = _afefd.Resolve()
				_cdeaf.Set(_bfgfc, _bgbgc)
			}
			if _dgdd := _acde.addObjects(_bgbgc); _dgdd != nil {
				return _dgdd
			}
		}
		return nil
	}
	if _, _gbefe := _fcggf.(*_abf.PdfObjectReference); _gbefe {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0062\u0065\u0020\u0061\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u002d\u0020\u0067\u006f\u0074 \u0025\u0023\u0076\u0021", _fcggf)
		return _fd.New("r\u0065\u0066\u0065\u0072en\u0063e\u0020\u006e\u006f\u0074\u0020a\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	return nil
}

func _gacgg() string {
	_gaabd.Lock()
	defer _gaabd.Unlock()
	return _babfc
}

type pdfSignDictionary struct {
	*_abf.PdfObjectDictionary
	_fafgf *SignatureHandler
	_dcbed *PdfSignature
	_eefbf int64
	_dgfdf int
	_afgef int
	_edcbf int
	_bcbcg int
}

// NewPdfColorCalRGB returns a new CalRBG color.
func NewPdfColorCalRGB(a, b, c float64) *PdfColorCalRGB {
	_dgab := PdfColorCalRGB{a, b, c}
	return &_dgab
}

// GetXHeight returns the XHeight of the font `descriptor`.
func (_eebcd *PdfFontDescriptor) GetXHeight() (float64, error) {
	return _abf.GetNumberAsFloat(_eebcd.XHeight)
}

// GetDocMDPPermission returns the DocMDP level of the restrictions
func (_fbecd *PdfSignature) GetDocMDPPermission() (_df.DocMDPPermission, bool) {
	for _, _ecaee := range _fbecd.Reference.Elements() {
		if _eecbg, _bdfda := _abf.GetDict(_ecaee); _bdfda {
			if _ggaeb, _cgade := _abf.GetNameVal(_eecbg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _cgade && _ggaeb == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				if _cfac, _fdbfg := _abf.GetDict(_eecbg.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _fdbfg {
					if P, _dcece := _abf.GetIntVal(_cfac.Get("\u0050")); _dcece {
						return _df.DocMDPPermission(P), true
					}
				}
			}
		}
	}
	return 0, false
}

func (_dgeb *PdfReader) newPdfAnnotationSquareFromDict(_deca *_abf.PdfObjectDictionary) (*PdfAnnotationSquare, error) {
	_cbdb := PdfAnnotationSquare{}
	_cgge, _cee := _dgeb.newPdfAnnotationMarkupFromDict(_deca)
	if _cee != nil {
		return nil, _cee
	}
	_cbdb.PdfAnnotationMarkup = _cgge
	_cbdb.BS = _deca.Get("\u0042\u0053")
	_cbdb.IC = _deca.Get("\u0049\u0043")
	_cbdb.BE = _deca.Get("\u0042\u0045")
	_cbdb.RD = _deca.Get("\u0052\u0044")
	return &_cbdb, nil
}

// PdfAppender appends new PDF content to an existing PDF document via incremental updates.
type PdfAppender struct {
	_eeded _gc.ReadSeeker
	_bdcd  *_abf.PdfParser
	_agda  *PdfReader
	Reader *PdfReader
	_cggfa []*PdfPage
	_ffbb  *PdfAcroForm
	_ffbe  *DSS
	_edcbe *Permissions
	_abce  _abf.XrefTable
	_dac   int64
	_ffc   int
	_ffcf  []_abf.PdfObject
	_gcba  map[_abf.PdfObject]struct{}
	_bge   map[_abf.PdfObject]int64
	_cdbbg map[_abf.PdfObject]struct{}
	_gfeg  map[_abf.PdfObject]struct{}
	_cfga  int64
	_ccaf  bool
	_fcfb  string
	_bbag  *EncryptOptions
	_acff  *PdfInfo
}

// NewPdfAnnotationLine returns a new line annotation.
func NewPdfAnnotationLine() *PdfAnnotationLine {
	_dde := NewPdfAnnotation()
	_aaad := &PdfAnnotationLine{}
	_aaad.PdfAnnotation = _dde
	_aaad.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_dde.SetContext(_aaad)
	return _aaad
}

// Normalize swaps (Llx,Urx) if Urx < Llx, and (Lly,Ury) if Ury < Lly.
func (_gbge *PdfRectangle) Normalize() {
	if _gbge.Llx > _gbge.Urx {
		_gbge.Llx, _gbge.Urx = _gbge.Urx, _gbge.Llx
	}
	if _gbge.Lly > _gbge.Ury {
		_gbge.Lly, _gbge.Ury = _gbge.Ury, _gbge.Lly
	}
}

// ApplyStandard is used to apply changes required on the document to match the rules required by the input standard.
// The writer's content would be changed after all the document parts are already established during the Write method.
// A good example of the StandardApplier could be a PDF/A Profile (i.e.: pdfa.Profile1A). In such a case PdfWriter would
// set up all rules required by that Profile.
func (_agffa *PdfWriter) ApplyStandard(optimizer StandardApplier) { _agffa._adgdc = optimizer }

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
	_gffg *_abf.PdfIndirectObject
	Certs []*_abf.PdfObjectStream
	OCSPs []*_abf.PdfObjectStream
	CRLs  []*_abf.PdfObjectStream
	VRI   map[string]*VRI
	_gcee map[string]*_abf.PdfObjectStream
	_ggfg map[string]*_abf.PdfObjectStream
	_daee map[string]*_abf.PdfObjectStream
}

// NewPdfReaderWithOpts creates a new PdfReader for an input io.ReadSeeker interface
// with a ReaderOpts.
// If ReaderOpts is nil it will be set to default value from NewReaderOpts.
func NewPdfReaderWithOpts(rs _gc.ReadSeeker, opts *ReaderOpts) (*PdfReader, error) {
	const _dcfdf = "\u006d\u006f\u0064\u0065\u006c\u003a\u004e\u0065\u0077\u0050\u0064f\u0052\u0065\u0061\u0064\u0065\u0072\u0057\u0069\u0074\u0068O\u0070\u0074\u0073"
	return _fbaec(rs, opts, true, _dcfdf)
}

// AddWatermarkImage adds a watermark to the page.
func (_bgabg *PdfPage) AddWatermarkImage(ximg *XObjectImage, opt WatermarkImageOptions) error {
	_aggda, _ffegba := _bgabg.GetMediaBox()
	if _ffegba != nil {
		return _ffegba
	}
	_ddfff := _aggda.Urx - _aggda.Llx
	_cdbbd := _aggda.Ury - _aggda.Lly
	_cecgg := float64(*ximg.Width)
	_fceff := (_ddfff - _cecgg) / 2
	if opt.FitToWidth {
		_cecgg = _ddfff
		_fceff = 0
	}
	_debff := _cdbbd
	_gcdb := float64(0)
	if opt.PreserveAspectRatio {
		_debff = _cecgg * float64(*ximg.Height) / float64(*ximg.Width)
		_gcdb = (_cdbbd - _debff) / 2
	}
	if _bgabg.Resources == nil {
		_bgabg.Resources = NewPdfPageResources()
	}
	_cggea := 0
	_cfcb := _abf.PdfObjectName(_e.Sprintf("\u0049\u006d\u0077%\u0064", _cggea))
	for _bgabg.Resources.HasXObjectByName(_cfcb) {
		_cggea++
		_cfcb = _abf.PdfObjectName(_e.Sprintf("\u0049\u006d\u0077%\u0064", _cggea))
	}
	_ffegba = _bgabg.AddImageResource(_cfcb, ximg)
	if _ffegba != nil {
		return _ffegba
	}
	_cggea = 0
	_ebagf := _abf.PdfObjectName(_e.Sprintf("\u0047\u0053\u0025\u0064", _cggea))
	for _bgabg.HasExtGState(_ebagf) {
		_cggea++
		_ebagf = _abf.PdfObjectName(_e.Sprintf("\u0047\u0053\u0025\u0064", _cggea))
	}
	_edbf := _abf.MakeDict()
	_edbf.Set("\u0042\u004d", _abf.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	_edbf.Set("\u0043\u0041", _abf.MakeFloat(opt.Alpha))
	_edbf.Set("\u0063\u0061", _abf.MakeFloat(opt.Alpha))
	_ffegba = _bgabg.AddExtGState(_ebagf, _edbf)
	if _ffegba != nil {
		return _ffegba
	}
	_decaf := _e.Sprintf("\u0071\u000a"+"\u002f%\u0073\u0020\u0067\u0073\u000a"+"%\u002e\u0030\u0066\u0020\u0030\u00200\u0020\u0025\u002e\u0030\u0066\u0020\u0025\u002e\u0034f\u0020\u0025\u002e4\u0066 \u0063\u006d\u000a"+"\u002f%\u0073\u0020\u0044\u006f\u000a"+"\u0051", _ebagf, _cecgg, _debff, _fceff, _gcdb, _cfcb)
	_bgabg.AddContentStreamByString(_decaf)
	return nil
}

// PdfAnnotationPolyLine represents PolyLine annotations.
// (Section 12.5.6.9).
type PdfAnnotationPolyLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	Vertices _abf.PdfObject
	LE       _abf.PdfObject
	BS       _abf.PdfObject
	IC       _abf.PdfObject
	BE       _abf.PdfObject
	IT       _abf.PdfObject
	Measure  _abf.PdfObject
}

// WriteString outputs the object as it is to be written to file.
func (_caae *PdfTransformParamsDocMDP) WriteString() string { return _caae.ToPdfObject().WriteString() }

// Sign signs a specific page with a digital signature.
// The signature field parameter must have a valid signature dictionary
// specified by its V field.
func (_ceba *PdfAppender) Sign(pageNum int, field *PdfFieldSignature) error {
	if field == nil {
		return _fd.New("\u0073\u0069g\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 n\u0069\u006c")
	}
	_cfgd := field.V
	if _cfgd == nil {
		return _fd.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061r\u0079 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_agbd := pageNum - 1
	if _agbd < 0 || _agbd > len(_ceba._cggfa)-1 {
		return _e.Errorf("\u0070\u0061\u0067\u0065\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064", pageNum)
	}
	_gcbf := _ceba.Reader.PageList[_agbd]
	field.P = _gcbf.ToPdfObject()
	if field.T == nil || field.T.String() == "" {
		field.T = _abf.MakeString(_e.Sprintf("\u0053\u0069\u0067n\u0061\u0074\u0075\u0072\u0065\u0020\u0025\u0064", pageNum))
	}
	_gcbf.AddAnnotation(field.PdfAnnotationWidget.PdfAnnotation)
	if _ceba._ffbb == _ceba._agda.AcroForm {
		_ceba._ffbb = _ceba.Reader.AcroForm
	}
	_cdad := _ceba._ffbb
	if _cdad == nil {
		_cdad = NewPdfAcroForm()
	}
	_cdad.SigFlags = _abf.MakeInteger(3)
	if _cdad.NeedAppearances != nil {
		_cdad.NeedAppearances = nil
	}
	_aaeaa := append(_cdad.AllFields(), field.PdfField)
	_cdad.Fields = &_aaeaa
	_ceba.ReplaceAcroForm(_cdad)
	_ceba.UpdatePage(_gcbf)
	_ceba._cggfa[_agbd] = _gcbf
	if _, _gfda := field.V.GetDocMDPPermission(); _gfda {
		_ceba._edcbe = NewPermissions(field.V)
	}
	return nil
}

// ColorToRGB only converts color used with uncolored patterns (defined in underlying colorspace).  Does not go into the
// pattern objects and convert those.  If that is desired, needs to be done separately.  See for example
// grayscale conversion example in unidoc-examples repo.
func (_bcdbf *PdfColorspaceSpecialPattern) ColorToRGB(color PdfColor) (PdfColor, error) {
	_bddb, _fbbe := color.(*PdfColorPattern)
	if !_fbbe {
		_acd.Log.Debug("\u0043\u006f\u006c\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0070a\u0074\u0074\u0065\u0072\u006e\u0020\u0028\u0067\u006f\u0074 \u0025\u0054\u0029", color)
		return nil, ErrTypeCheck
	}
	if _bddb.Color == nil {
		return color, nil
	}
	if _bcdbf.UnderlyingCS == nil {
		return nil, _fd.New("\u0075n\u0064\u0065\u0072\u006cy\u0069\u006e\u0067\u0020\u0043S\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	return _bcdbf.UnderlyingCS.ColorToRGB(_bddb.Color)
}

// NewPdfField returns an initialized PdfField.
func NewPdfField() *PdfField { return &PdfField{_dgdc: _abf.MakeIndirectObject(_abf.MakeDict())} }

func (_edfcf *PdfField) inherit(_gbdc func(*PdfField) bool) (bool, error) {
	_gafa := map[*PdfField]bool{}
	_fbfd := false
	_acge := _edfcf
	for _acge != nil {
		if _, _dfgf := _gafa[_acge]; _dfgf {
			return false, _fd.New("\u0072\u0065\u0063\u0075rs\u0069\u0076\u0065\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c")
		}
		_dgda := _gbdc(_acge)
		if _dgda {
			_fbfd = true
			break
		}
		_gafa[_acge] = true
		_acge = _acge.Parent
	}
	return _fbfd, nil
}

// NewPdfRectangle creates a PDF rectangle object based on an input array of 4 integers.
// Defining the lower left (LL) and upper right (UR) corners with
// floating point numbers.
func NewPdfRectangle(arr _abf.PdfObjectArray) (*PdfRectangle, error) {
	_cbbcg := PdfRectangle{}
	if arr.Len() != 4 {
		return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065\u0020\u0061\u0072r\u0061\u0079\u002c\u0020\u006c\u0065\u006e \u0021\u003d\u0020\u0034")
	}
	var _cadddg error
	_cbbcg.Llx, _cadddg = _abf.GetNumberAsFloat(arr.Get(0))
	if _cadddg != nil {
		return nil, _cadddg
	}
	_cbbcg.Lly, _cadddg = _abf.GetNumberAsFloat(arr.Get(1))
	if _cadddg != nil {
		return nil, _cadddg
	}
	_cbbcg.Urx, _cadddg = _abf.GetNumberAsFloat(arr.Get(2))
	if _cadddg != nil {
		return nil, _cadddg
	}
	_cbbcg.Ury, _cadddg = _abf.GetNumberAsFloat(arr.Get(3))
	if _cadddg != nil {
		return nil, _cadddg
	}
	return &_cbbcg, nil
}
func (_bgca *pdfFontType3) baseFields() *fontCommon                { return &_bgca.fontCommon }
func (_eafag *pdfFontType3) getFontDescriptor() *PdfFontDescriptor { return _eafag._dcbaf }
func _aacbg() string {
	_gaabd.Lock()
	defer _gaabd.Unlock()
	if len(_edead) > 0 {
		return _edead
	}
	return "\u0055n\u0069\u0044\u006f\u0063 \u002d\u0020\u0068\u0074\u0074p\u003a/\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069o"
}

func _fgeba(_bgdg _abf.PdfObject) (*PdfFunctionType2, error) {
	_geaaa := &PdfFunctionType2{}
	var _faae *_abf.PdfObjectDictionary
	if _bcbge, _bffde := _bgdg.(*_abf.PdfIndirectObject); _bffde {
		_adbd, _geggg := _bcbge.PdfObject.(*_abf.PdfObjectDictionary)
		if !_geggg {
			return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_geaaa._gaaae = _bcbge
		_faae = _adbd
	} else if _dffcg, _egaef := _bgdg.(*_abf.PdfObjectDictionary); _egaef {
		_faae = _dffcg
	} else {
		return nil, _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_acd.Log.Trace("\u0046U\u004e\u0043\u0032\u003a\u0020\u0025s", _faae.String())
	_eeafe, _eefea := _abf.TraceToDirectObject(_faae.Get("\u0044\u006f\u006d\u0061\u0069\u006e")).(*_abf.PdfObjectArray)
	if !_eefea {
		_acd.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065\u0063\u0069fi\u0065\u0064")
		return nil, _fd.New("\u0072\u0065q\u0075\u0069\u0072\u0065d\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020m\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _eeafe.Len() < 0 || _eeafe.Len()%2 != 0 {
		_acd.Log.Error("D\u006fm\u0061\u0069\u006e\u0020\u0072\u0061\u006e\u0067e\u0020\u0069\u006e\u0076al\u0069\u0064")
		return nil, _fd.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0064\u006f\u006da\u0069\u006e\u0020\u0072an\u0067\u0065")
	}
	_faggf, _ecbg := _eeafe.ToFloat64Array()
	if _ecbg != nil {
		return nil, _ecbg
	}
	_geaaa.Domain = _faggf
	_eeafe, _eefea = _abf.TraceToDirectObject(_faae.Get("\u0052\u0061\u006eg\u0065")).(*_abf.PdfObjectArray)
	if _eefea {
		if _eeafe.Len() < 0 || _eeafe.Len()%2 != 0 {
			return nil, _fd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_bbadb, _cbaa := _eeafe.ToFloat64Array()
		if _cbaa != nil {
			return nil, _cbaa
		}
		_geaaa.Range = _bbadb
	}
	_eeafe, _eefea = _abf.TraceToDirectObject(_faae.Get("\u0043\u0030")).(*_abf.PdfObjectArray)
	if _eefea {
		_geadd, _ecgbb := _eeafe.ToFloat64Array()
		if _ecgbb != nil {
			return nil, _ecgbb
		}
		_geaaa.C0 = _geadd
	}
	_eeafe, _eefea = _abf.TraceToDirectObject(_faae.Get("\u0043\u0031")).(*_abf.PdfObjectArray)
	if _eefea {
		_eagfg, _bbfbd := _eeafe.ToFloat64Array()
		if _bbfbd != nil {
			return nil, _bbfbd
		}
		_geaaa.C1 = _eagfg
	}
	if len(_geaaa.C0) != len(_geaaa.C1) {
		_acd.Log.Error("\u0043\u0030\u0020\u0061nd\u0020\u0043\u0031\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069n\u0067")
		return nil, _abf.ErrRangeError
	}
	N, _ecbg := _abf.GetNumberAsFloat(_abf.TraceToDirectObject(_faae.Get("\u004e")))
	if _ecbg != nil {
		_acd.Log.Error("\u004e\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020o\u0072\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u002c\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _faae.String())
		return nil, _ecbg
	}
	_geaaa.N = N
	return _geaaa, nil
}

// Clear clears flag fl from the flag and returns the resulting flag.
func (_dcbfe FieldFlag) Clear(fl FieldFlag) FieldFlag { return FieldFlag(_dcbfe.Mask() &^ fl.Mask()) }

// A returns the value of the A component of the color.
func (_dfabc *PdfColorLab) A() float64 { return _dfabc[1] }

// ToPdfObject converts the PdfPage to a dictionary within an indirect object container.
func (_bdaf *PdfPage) ToPdfObject() _abf.PdfObject {
	_bcfdd := _bdaf._gefee
	_bdaf.GetPageDict()
	return _bcfdd
}

// SetEncoder sets the encoding for the underlying font.
// TODO(peterwilliams97): Change function signature to SetEncoder(encoder *textencoding.simpleEncoder).
// TODO(gunnsth): Makes sense if SetEncoder is removed from the interface fonts.Font as proposed in PR #260.
func (_bgeea *pdfFontSimple) SetEncoder(encoder _cbb.TextEncoder) { _bgeea._ebada = encoder }

// Write writes out the PDF.
func (_adbed *PdfWriter) Write(writer _gc.Writer) error {
	_acd.Log.Trace("\u0057r\u0069\u0074\u0065\u0028\u0029")
	if _gcdea := _adbed.writeOutlines(); _gcdea != nil {
		return _gcdea
	}
	if _gcdea := _adbed.writeAcroFormFields(); _gcdea != nil {
		return _gcdea
	}
	_adbed.checkPendingObjects()
	if _gcdea := _adbed.writeOutputIntents(); _gcdea != nil {
		return _gcdea
	}
	_adbed.setCatalogVersion()
	_adbed.copyObjects()
	if _gcdea := _adbed.optimize(); _gcdea != nil {
		return _gcdea
	}
	if _gcdea := _adbed.optimizeDocument(); _gcdea != nil {
		return _gcdea
	}
	var _aadbc _a.Hash
	if _adbed._fegae {
		_aadbc = _ag.New()
		writer = _gc.MultiWriter(_aadbc, writer)
	}
	_adbed.setWriter(writer)
	_aeedf := _adbed.checkCrossReferenceStream()
	_gggbf, _aeedf := _adbed.mapObjectStreams(_aeedf)
	_adbed.adjustXRefAffectedVersion(_aeedf)
	_adbed.writeDocumentVersion()
	_adbed.updateObjectNumbers()
	_adbed.writeObjects()
	if _gcdea := _adbed.writeObjectsInStreams(_gggbf); _gcdea != nil {
		return _gcdea
	}
	_deeae := _adbed._dbfaad
	var _fbaea int
	for _agcdc := range _adbed._becfc {
		if _agcdc > _fbaea {
			_fbaea = _agcdc
		}
	}
	if _adbed._fegae {
		if _gcdea := _adbed.setHashIDs(_aadbc); _gcdea != nil {
			return _gcdea
		}
	}
	if _aeedf {
		if _gcdea := _adbed.writeXRefStreams(_fbaea, _deeae); _gcdea != nil {
			return _gcdea
		}
	} else {
		_adbed.writeTrailer(_fbaea)
	}
	_adbed.makeOffSetReference(_deeae)
	if _gcdea := _adbed.flushWriter(); _gcdea != nil {
		return _gcdea
	}
	return nil
}

// NewPdfColorspaceCalRGB returns a new CalRGB colorspace object.
func NewPdfColorspaceCalRGB() *PdfColorspaceCalRGB {
	_aecf := &PdfColorspaceCalRGB{}
	_aecf.BlackPoint = []float64{0.0, 0.0, 0.0}
	_aecf.Gamma = []float64{1.0, 1.0, 1.0}
	_aecf.Matrix = []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	return _aecf
}

// NewPdfColorDeviceRGB returns a new PdfColorDeviceRGB based on the r,g,b component values.
func NewPdfColorDeviceRGB(r, g, b float64) *PdfColorDeviceRGB {
	_cbc := PdfColorDeviceRGB{r, g, b}
	return &_cbc
}

// PdfAnnotation represents an annotation in PDF (section 12.5 p. 389).
type PdfAnnotation struct {
	_edg         PdfModel
	Rect         _abf.PdfObject
	Contents     _abf.PdfObject
	P            _abf.PdfObject
	NM           _abf.PdfObject
	M            _abf.PdfObject
	F            _abf.PdfObject
	AP           _abf.PdfObject
	AS           _abf.PdfObject
	Border       _abf.PdfObject
	C            _abf.PdfObject
	StructParent _abf.PdfObject
	OC           _abf.PdfObject
	_dbc         *_abf.PdfIndirectObject
}

// ToPdfObject returns the PDF representation of the shading dictionary.
func (_bcgbe *PdfShadingType6) ToPdfObject() _abf.PdfObject {
	_bcgbe.PdfShading.ToPdfObject()
	_geaab, _gbbgb := _bcgbe.getShadingDict()
	if _gbbgb != nil {
		_acd.Log.Error("\u0055\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073\u0020\u0073\u0068\u0061\u0064\u0069\u006e\u0067\u0020di\u0063\u0074")
		return nil
	}
	if _bcgbe.BitsPerCoordinate != nil {
		_geaab.Set("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065", _bcgbe.BitsPerCoordinate)
	}
	if _bcgbe.BitsPerComponent != nil {
		_geaab.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", _bcgbe.BitsPerComponent)
	}
	if _bcgbe.BitsPerFlag != nil {
		_geaab.Set("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067", _bcgbe.BitsPerFlag)
	}
	if _bcgbe.Decode != nil {
		_geaab.Set("\u0044\u0065\u0063\u006f\u0064\u0065", _bcgbe.Decode)
	}
	if _bcgbe.Function != nil {
		if len(_bcgbe.Function) == 1 {
			_geaab.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _bcgbe.Function[0].ToPdfObject())
		} else {
			_fgdeag := _abf.MakeArray()
			for _, _fbgbc := range _bcgbe.Function {
				_fgdeag.Append(_fbgbc.ToPdfObject())
			}
			_geaab.Set("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e", _fgdeag)
		}
	}
	return _bcgbe._eabcgc
}

// IsTiling specifies if the pattern is a tiling pattern.
func (_cefa *PdfPattern) IsTiling() bool { return _cefa.PatternType == 1 }

// StringToCharcodeBytes maps the provided string runes to charcode bytes and
// it returns the resulting slice of bytes, along with the number of runes
// which could not be converted. If the number of misses is 0, all string runes
// were successfully converted.
func (_aaegg *PdfFont) StringToCharcodeBytes(str string) ([]byte, int) {
	return _aaegg.RunesToCharcodeBytes([]rune(str))
}

func _efea(_fbbf []byte) bool {
	if len(_fbbf) < 4 {
		return true
	}
	for _aebea := range _fbbf[:4] {
		_bbgfe := rune(_aebea)
		if !_gg.Is(_gg.ASCII_Hex_Digit, _bbgfe) && !_gg.IsSpace(_bbgfe) {
			return true
		}
	}
	return false
}

// NewPdfFontFromPdfObject loads a PdfFont from the dictionary `fontObj`.  If there is a problem an
// error is returned.
func NewPdfFontFromPdfObject(fontObj _abf.PdfObject) (*PdfFont, error) { return _caece(fontObj, true) }

// PdfActionTrans represents a trans action.
type PdfActionTrans struct {
	*PdfAction
	Trans _abf.PdfObject
}

// FillWithAppearance populates `form` with values provided by `provider`.
// If not nil, `appGen` is used to generate appearance dictionaries for the
// field annotations, based on the specified settings. Otherwise, appearance
// generation is skipped.
// e.g.: appGen := annotator.FieldAppearance{OnlyIfMissing: true, RegenerateTextFields: true}
// NOTE: In next major version this functionality will be part of Fill. (v4)
func (_aeeef *PdfAcroForm) FillWithAppearance(provider FieldValueProvider, appGen FieldAppearanceGenerator) error {
	_cgace := _aeeef.fill(provider, appGen)
	if _cgace != nil {
		return _cgace
	}
	if _, _dabbgf := provider.(FieldImageProvider); _dabbgf {
		_cgace = _aeeef.fillImageWithAppearance(provider.(FieldImageProvider), appGen)
	}
	return _cgace
}

// ToOutlineTree returns a low level PdfOutlineTreeNode object, based on
// the current instance.
func (_bfff *Outline) ToOutlineTree() *PdfOutlineTreeNode {
	return &_bfff.ToPdfOutline().PdfOutlineTreeNode
}

// PdfAnnotationLine represents Line annotations.
// (Section 12.5.6.7).
type PdfAnnotationLine struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	L       _abf.PdfObject
	BS      _abf.PdfObject
	LE      _abf.PdfObject
	IC      _abf.PdfObject
	LL      _abf.PdfObject
	LLE     _abf.PdfObject
	Cap     _abf.PdfObject
	IT      _abf.PdfObject
	LLO     _abf.PdfObject
	CP      _abf.PdfObject
	Measure _abf.PdfObject
	CO      _abf.PdfObject
}

// ToImage converts an object to an Image which can be transformed or saved out.
// The image data is decoded and the Image returned.
func (_aaddf *XObjectImage) ToImage() (*Image, error) {
	_edfbd := &Image{}
	if _aaddf.Height == nil {
		return nil, _fd.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_edfbd.Height = *_aaddf.Height
	if _aaddf.Width == nil {
		return nil, _fd.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_edfbd.Width = *_aaddf.Width
	if _aaddf.BitsPerComponent == nil {
		switch _aaddf.Filter.(type) {
		case *_abf.CCITTFaxEncoder, *_abf.JBIG2Encoder:
			_edfbd.BitsPerComponent = 1
		case *_abf.LZWEncoder, *_abf.RunLengthEncoder:
			_edfbd.BitsPerComponent = 8
		default:
			return nil, _fd.New("\u0062\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006fm\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
	} else {
		_edfbd.BitsPerComponent = *_aaddf.BitsPerComponent
	}
	_edfbd.ColorComponents = _aaddf.ColorSpace.GetNumComponents()
	_aaddf._ccbad.Set("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073", _abf.MakeInteger(int64(_edfbd.ColorComponents)))
	_adef, _adgbdg := _abf.DecodeStream(_aaddf._ccbad)
	if _adgbdg != nil {
		return nil, _adgbdg
	}
	_edfbd.Data = _adef
	if _aaddf.Decode != nil {
		_ggfe, _cccef := _aaddf.Decode.(*_abf.PdfObjectArray)
		if !_cccef {
			_acd.Log.Debug("I\u006e\u0076\u0061\u006cid\u0020D\u0065\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074")
			return nil, _fd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_fffbc, _ggfcgf := _ggfe.ToFloat64Array()
		if _ggfcgf != nil {
			return nil, _ggfcgf
		}
		switch _aaddf.ColorSpace.(type) {
		case *PdfColorspaceDeviceCMYK:
			_ecafbe := _aaddf.ColorSpace.DecodeArray()
			if _ecafbe[0] == _fffbc[0] && _ecafbe[1] == _fffbc[1] && _ecafbe[2] == _fffbc[2] && _ecafbe[3] == _fffbc[3] {
				_edfbd._ceeag = _fffbc
			}
		default:
			_edfbd._ceeag = _fffbc
		}
	}
	return _edfbd, nil
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
	_aefbg        []byte
	_bega         []uint32
	_cabaa        *_abf.PdfObjectStream
}

// GetContainingPdfObject gets the primitive used to parse the color space.
func (_bedd *PdfColorspaceICCBased) GetContainingPdfObject() _abf.PdfObject { return _bedd._bfgc }

func (_aaea *PdfReader) newPdfAnnotationHighlightFromDict(_baa *_abf.PdfObjectDictionary) (*PdfAnnotationHighlight, error) {
	_gceg := PdfAnnotationHighlight{}
	_bdab, _fage := _aaea.newPdfAnnotationMarkupFromDict(_baa)
	if _fage != nil {
		return nil, _fage
	}
	_gceg.PdfAnnotationMarkup = _bdab
	_gceg.QuadPoints = _baa.Get("\u0051\u0075\u0061\u0064\u0050\u006f\u0069\u006e\u0074\u0073")
	return &_gceg, nil
}

// ToPdfObject returns the PDF representation of the shading pattern.
func (_aeggg *PdfShadingPatternType3) ToPdfObject() _abf.PdfObject {
	_aeggg.PdfPattern.ToPdfObject()
	_bceaa := _aeggg.getDict()
	if _aeggg.Shading != nil {
		_bceaa.Set("\u0053h\u0061\u0064\u0069\u006e\u0067", _aeggg.Shading.ToPdfObject())
	}
	if _aeggg.Matrix != nil {
		_bceaa.Set("\u004d\u0061\u0074\u0072\u0069\u0078", _aeggg.Matrix)
	}
	if _aeggg.ExtGState != nil {
		_bceaa.Set("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e", _aeggg.ExtGState)
	}
	return _aeggg._bcfca
}

// PdfAnnotationSquare represents Square annotations.
// (Section 12.5.6.8).
type PdfAnnotationSquare struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	BS _abf.PdfObject
	IC _abf.PdfObject
	BE _abf.PdfObject
	RD _abf.PdfObject
}

// WriteToFile writes the Appender output to file specified by path.
func (_ccec *PdfAppender) WriteToFile(outputPath string) error {
	_facc, _addb := _cf.Create(outputPath)
	if _addb != nil {
		return _addb
	}
	defer _facc.Close()
	return _ccec.Write(_facc)
}

func _fecf(_cfafb _abf.PdfObject) (map[_cbb.CharCode]float64, error) {
	if _cfafb == nil {
		return nil, nil
	}
	_ebfc, _caad := _abf.GetArray(_cfafb)
	if !_caad {
		return nil, nil
	}
	_gdee := map[_cbb.CharCode]float64{}
	_cgfe := _ebfc.Len()
	for _ccad := 0; _ccad < _cgfe-1; _ccad++ {
		_ggcb := _abf.TraceToDirectObject(_ebfc.Get(_ccad))
		_faffg, _fgeda := _abf.GetIntVal(_ggcb)
		if !_fgeda {
			return nil, _e.Errorf("\u0042a\u0064\u0020\u0066\u006fn\u0074\u0020\u0057\u0020\u006fb\u006a0\u003a \u0069\u003d\u0025\u0064\u0020\u0025\u0023v", _ccad, _ggcb)
		}
		_ccad++
		if _ccad > _cgfe-1 {
			return nil, _e.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _ebfc)
		}
		_geeae := _abf.TraceToDirectObject(_ebfc.Get(_ccad))
		switch _geeae.(type) {
		case *_abf.PdfObjectArray:
			_fbdae, _ := _abf.GetArray(_geeae)
			if _ebeb, _bcefg := _fbdae.ToFloat64Array(); _bcefg == nil {
				for _cedc := 0; _cedc < len(_ebeb); _cedc++ {
					_gdee[_cbb.CharCode(_faffg+_cedc)] = _ebeb[_cedc]
				}
			} else {
				return nil, _e.Errorf("\u0042\u0061\u0064 \u0066\u006f\u006e\u0074 \u0057\u0020\u0061\u0072\u0072\u0061\u0079 \u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _ccad, _geeae)
			}
		case *_abf.PdfObjectInteger:
			_fgca, _gfagb := _abf.GetIntVal(_geeae)
			if !_gfagb {
				return nil, _e.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0031\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _ccad, _geeae)
			}
			_ccad++
			if _ccad > _cgfe-1 {
				return nil, _e.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020a\u0072\u0072\u0061\u0079\u003a\u0020\u0061\u0072\u0072\u0032=\u0025\u002b\u0076", _ebfc)
			}
			_eebgc := _ebfc.Get(_ccad)
			_gebg, _cdfba := _abf.GetNumberAsFloat(_eebgc)
			if _cdfba != nil {
				return nil, _e.Errorf("\u0042\u0061d\u0020\u0066\u006f\u006e\u0074\u0020\u0057\u0020\u0069\u006e\u0074\u0020\u006f\u0062\u006a\u0032\u003a\u0020\u0069\u003d\u0025\u0064 %\u0023\u0076", _ccad, _eebgc)
			}
			for _ggbbb := _faffg; _ggbbb <= _fgca; _ggbbb++ {
				_gdee[_cbb.CharCode(_ggbbb)] = _gebg
			}
		default:
			return nil, _e.Errorf("\u0042\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0057 \u006f\u0062\u006a\u0031\u0020\u0074\u0079p\u0065\u003a\u0020\u0069\u003d\u0025\u0064\u0020\u0025\u0023\u0076", _ccad, _geeae)
		}
	}
	return _gdee, nil
}

// ToUnicode returns the name of the font's "ToUnicode" field if there is one, or "" if there isn't.
func (_cbgb *PdfFont) ToUnicode() string {
	if _cbgb.baseFields()._aabfe == nil {
		return ""
	}
	return _cbgb.baseFields()._aabfe.Name()
}

// ToPdfObject implements model.PdfModel interface.
func (_fdecb *PdfOutputIntent) ToPdfObject() _abf.PdfObject {
	if _fdecb._dcfb == nil {
		_fdecb._dcfb = _abf.MakeDict()
	}
	_babe := _fdecb._dcfb
	if _fdecb.Type != "" {
		_babe.Set("\u0054\u0079\u0070\u0065", _abf.MakeName(_fdecb.Type))
	}
	_babe.Set("\u0053", _abf.MakeName(_fdecb.S.String()))
	if _fdecb.OutputCondition != "" {
		_babe.Set("\u004fu\u0074p\u0075\u0074\u0043\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e", _abf.MakeString(_fdecb.OutputCondition))
	}
	_babe.Set("\u004fu\u0074\u0070\u0075\u0074C\u006f\u006e\u0064\u0069\u0074i\u006fn\u0049d\u0065\u006e\u0074\u0069\u0066\u0069\u0065r", _abf.MakeString(_fdecb.OutputConditionIdentifier))
	_babe.Set("\u0052\u0065\u0067i\u0073\u0074\u0072\u0079\u004e\u0061\u006d\u0065", _abf.MakeString(_fdecb.RegistryName))
	if _fdecb.Info != "" {
		_babe.Set("\u0049\u006e\u0066\u006f", _abf.MakeString(_fdecb.Info))
	}
	if len(_fdecb.DestOutputProfile) != 0 {
		_dcbb, _caggb := _abf.MakeStream(_fdecb.DestOutputProfile, _abf.NewFlateEncoder())
		if _caggb != nil {
			_acd.Log.Error("\u004d\u0061\u006b\u0065\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0044\u0065s\u0074\u004f\u0075\u0074\u0070\u0075t\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _caggb)
		}
		_dcbb.PdfObjectDictionary.Set("\u004e", _abf.MakeInteger(int64(_fdecb.ColorComponents)))
		_bgada := make([]float64, _fdecb.ColorComponents*2)
		for _cfccgb := 0; _cfccgb < _fdecb.ColorComponents*2; _cfccgb++ {
			_faccf := 0.0
			if _cfccgb%2 != 0 {
				_faccf = 1.0
			}
			_bgada[_cfccgb] = _faccf
		}
		_dcbb.PdfObjectDictionary.Set("\u0052\u0061\u006eg\u0065", _abf.MakeArrayFromFloats(_bgada))
		_babe.Set("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065", _dcbb)
	}
	return _babe
}

// SetDecode sets the decode image float slice.
func (_bcab *Image) SetDecode(decode []float64) { _bcab._ceeag = decode }

func _gcfd(_egeg _abf.PdfObject) (*PdfColorspaceCalGray, error) {
	_cdfbcf := NewPdfColorspaceCalGray()
	if _beegb, _dccg := _egeg.(*_abf.PdfIndirectObject); _dccg {
		_cdfbcf._dgcg = _beegb
	}
	_egeg = _abf.TraceToDirectObject(_egeg)
	_dfag, _dgbf := _egeg.(*_abf.PdfObjectArray)
	if !_dgbf {
		return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _dfag.Len() != 2 {
		return nil, _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0061\u006cG\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0072\u0073p\u0061\u0063\u0065")
	}
	_egeg = _abf.TraceToDirectObject(_dfag.Get(0))
	_fgaf, _dgbf := _egeg.(*_abf.PdfObjectName)
	if !_dgbf {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u0020\u006e\u0061m\u0065\u0020\u006e\u006f\u0074\u0020\u0061 \u004e\u0061\u006d\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	if *_fgaf != "\u0043a\u006c\u0047\u0072\u0061\u0079" {
		return nil, _e.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0043\u0061\u006c\u0047\u0072\u0061\u0079 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065")
	}
	_egeg = _abf.TraceToDirectObject(_dfag.Get(1))
	_gdda, _dgbf := _egeg.(*_abf.PdfObjectDictionary)
	if !_dgbf {
		return nil, _e.Errorf("\u0043\u0061lG\u0072\u0061\u0079 \u0064\u0069\u0063\u0074 no\u0074 a\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u006f\u0062\u006a\u0065c\u0074")
	}
	_egeg = _gdda.Get("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074")
	_egeg = _abf.TraceToDirectObject(_egeg)
	_agde, _dgbf := _egeg.(*_abf.PdfObjectArray)
	if !_dgbf {
		return nil, _e.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020W\u0068\u0069\u0074e\u0050o\u0069\u006e\u0074")
	}
	if _agde.Len() != 3 {
		return nil, _e.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0068\u0069t\u0065\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
	}
	_ecgcb, _abaag := _agde.GetAsFloat64Slice()
	if _abaag != nil {
		return nil, _abaag
	}
	_cdfbcf.WhitePoint = _ecgcb
	_egeg = _gdda.Get("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074")
	if _egeg != nil {
		_egeg = _abf.TraceToDirectObject(_egeg)
		_feaf, _ccafe := _egeg.(*_abf.PdfObjectArray)
		if !_ccafe {
			return nil, _e.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020B\u006c\u0061\u0063k\u0050o\u0069\u006e\u0074")
		}
		if _feaf.Len() != 3 {
			return nil, _e.Errorf("\u0043\u0061\u006c\u0047\u0072\u0061y\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u006c\u0061c\u006b\u0050\u006f\u0069\u006e\u0074\u0020a\u0072\u0072\u0061\u0079")
		}
		_eded, _cdfa := _feaf.GetAsFloat64Slice()
		if _cdfa != nil {
			return nil, _cdfa
		}
		_cdfbcf.BlackPoint = _eded
	}
	_egeg = _gdda.Get("\u0047\u0061\u006dm\u0061")
	if _egeg != nil {
		_egeg = _abf.TraceToDirectObject(_egeg)
		_gacgb, _cbdf := _abf.GetNumberAsFloat(_egeg)
		if _cbdf != nil {
			return nil, _e.Errorf("C\u0061\u006c\u0047\u0072\u0061\u0079:\u0020\u0067\u0061\u006d\u006d\u0061\u0020\u006e\u006ft\u0020\u0061\u0020n\u0075m\u0062\u0065\u0072")
		}
		_cdfbcf.Gamma = _gacgb
	}
	return _cdfbcf, nil
}

// DecodeArray returns the component range values for the Separation colorspace.
func (_bbae *PdfColorspaceSpecialSeparation) DecodeArray() []float64 { return []float64{0, 1.0} }

// SetDocInfo sets the document /Info metadata.
// This will overwrite any globally declared document info.
func (_bfbb *PdfAppender) SetDocInfo(info *PdfInfo) { _bfbb._acff = info }

// PdfTransformParamsDocMDP represents a transform parameters dictionary for the DocMDP method and is used to detect
// modifications relative to a signature field that is signed by the author of a document.
// (Section 12.8.2.2, Table 254 - Entries in the DocMDP transform parameters dictionary p. 471 in PDF32000_2008).
type PdfTransformParamsDocMDP struct {
	Type *_abf.PdfObjectName
	P    *_abf.PdfObjectInteger
	V    *_abf.PdfObjectName
}

// PdfAnnotationHighlight represents Highlight annotations.
// (Section 12.5.6.10).
type PdfAnnotationHighlight struct {
	*PdfAnnotation
	*PdfAnnotationMarkup
	QuadPoints _abf.PdfObject
}

// PdfColorspaceSpecialIndexed is an indexed color space is a lookup table, where the input element
// is an index to the lookup table and the output is a color defined in the lookup table in the Base
// colorspace.
// [/Indexed base hival lookup]
type PdfColorspaceSpecialIndexed struct {
	Base   PdfColorspace
	HiVal  int
	Lookup _abf.PdfObject
	_bcdf  []byte
	_acea  *_abf.PdfIndirectObject
}

func (_adfb *PdfFunctionType0) processSamples() error {
	_fffdc := _gf.ResampleBytes(_adfb._aefbg, _adfb.BitsPerSample)
	_adfb._bega = _fffdc
	return nil
}

// FieldAppearanceGenerator generates appearance stream for a given field.
type FieldAppearanceGenerator interface {
	ContentStreamWrapper
	GenerateAppearanceDict(_gfaca *PdfAcroForm, _gdfc *PdfField, _egad *PdfAnnotationWidget) (*_abf.PdfObjectDictionary, error)
}

func (_dba *PdfReader) newPdfActionHideFromDict(_fbag *_abf.PdfObjectDictionary) (*PdfActionHide, error) {
	return &PdfActionHide{T: _fbag.Get("\u0054"), H: _fbag.Get("\u0048")}, nil
}

func (_dgbg *PdfAppender) mergeResources(_dbcc, _abdc _abf.PdfObject, _agfe map[_abf.PdfObjectName]_abf.PdfObjectName) _abf.PdfObject {
	if _abdc == nil && _dbcc == nil {
		return nil
	}
	if _abdc == nil {
		return _dbcc
	}
	_feccc, _egef := _abf.GetDict(_abdc)
	if !_egef {
		return _dbcc
	}
	if _dbcc == nil {
		_gbfd := _abf.MakeDict()
		_gbfd.Merge(_feccc)
		return _abdc
	}
	_abcc, _egef := _abf.GetDict(_dbcc)
	if !_egef {
		_acd.Log.Error("\u0045\u0072\u0072or\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065 \u0069s\u0020n\u006ft\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		_abcc = _abf.MakeDict()
	}
	for _, _gace := range _feccc.Keys() {
		if _dcbef, _fgc := _agfe[_gace]; _fgc {
			_abcc.Set(_dcbef, _feccc.Get(_gace))
		} else {
			_abcc.Set(_gace, _feccc.Get(_gace))
		}
	}
	return _abcc
}

// GetNumComponents returns the number of color components of the colorspace device.
// Returns 3 for a Lab device.
func (_cfaef *PdfColorspaceLab) GetNumComponents() int { return 3 }

func (_fba *PdfReader) newPdfActionGotoFromDict(_ffb *_abf.PdfObjectDictionary) (*PdfActionGoTo, error) {
	return &PdfActionGoTo{D: _ffb.Get("\u0044")}, nil
}

type fontFile struct {
	_gadc  string
	_eadac string
	_eedb  _cbb.SimpleEncoder
}

// GetContainingPdfObject implements interface PdfModel.
func (_bgbb *PdfFilespec) GetContainingPdfObject() _abf.PdfObject { return _bgbb._badbg }

// SetContext sets the sub action (context).
func (_aee *PdfAction) SetContext(ctx PdfModel) { _aee._gfg = ctx }

// NewPdfAnnotationCaret returns a new caret annotation.
func NewPdfAnnotationCaret() *PdfAnnotationCaret {
	_ggb := NewPdfAnnotation()
	_fag := &PdfAnnotationCaret{}
	_fag.PdfAnnotation = _ggb
	_fag.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_ggb.SetContext(_fag)
	return _fag
}

// PdfAnnotationPrinterMark represents PrinterMark annotations.
// (Section 12.5.6.20).
type PdfAnnotationPrinterMark struct {
	*PdfAnnotation
	MN _abf.PdfObject
}

// ToPdfObject implements interface PdfModel.
// Note: Call the sub-annotation's ToPdfObject to set both the generic and non-generic information.
func (_aga *PdfAnnotation) ToPdfObject() _abf.PdfObject {
	_acbe := _aga._dbc
	_add := _acbe.PdfObject.(*_abf.PdfObjectDictionary)
	_add.Clear()
	_add.Set("\u0054\u0079\u0070\u0065", _abf.MakeName("\u0041\u006e\u006eo\u0074"))
	_add.SetIfNotNil("\u0052\u0065\u0063\u0074", _aga.Rect)
	_add.SetIfNotNil("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _aga.Contents)
	_add.SetIfNotNil("\u0050", _aga.P)
	_add.SetIfNotNil("\u004e\u004d", _aga.NM)
	_add.SetIfNotNil("\u004d", _aga.M)
	_add.SetIfNotNil("\u0046", _aga.F)
	_add.SetIfNotNil("\u0041\u0050", _aga.AP)
	_add.SetIfNotNil("\u0041\u0053", _aga.AS)
	_add.SetIfNotNil("\u0042\u006f\u0072\u0064\u0065\u0072", _aga.Border)
	_add.SetIfNotNil("\u0043", _aga.C)
	_add.SetIfNotNil("\u0053\u0074\u0072u\u0063\u0074\u0050\u0061\u0072\u0065\u006e\u0074", _aga.StructParent)
	_add.SetIfNotNil("\u004f\u0043", _aga.OC)
	return _acbe
}

// ToPdfObject returns a PDF object representation of the outline destination.
func (_fedff OutlineDest) ToPdfObject() _abf.PdfObject {
	if (_fedff.PageObj == nil && _fedff.Page < 0) || _fedff.Mode == "" {
		return _abf.MakeNull()
	}
	_cfdbg := _abf.MakeArray()
	if _fedff.PageObj != nil {
		_cfdbg.Append(_fedff.PageObj)
	} else {
		_cfdbg.Append(_abf.MakeInteger(_fedff.Page))
	}
	_cfdbg.Append(_abf.MakeName(_fedff.Mode))
	switch _fedff.Mode {
	case "\u0046\u0069\u0074", "\u0046\u0069\u0074\u0042":
	case "\u0046\u0069\u0074\u0048", "\u0046\u0069\u0074B\u0048":
		_cfdbg.Append(_abf.MakeFloat(_fedff.Y))
	case "\u0046\u0069\u0074\u0056", "\u0046\u0069\u0074B\u0056":
		_cfdbg.Append(_abf.MakeFloat(_fedff.X))
	case "\u0058\u0059\u005a":
		_cfdbg.Append(_abf.MakeFloat(_fedff.X))
		_cfdbg.Append(_abf.MakeFloat(_fedff.Y))
		_cfdbg.Append(_abf.MakeFloat(_fedff.Zoom))
	default:
		_cfdbg.Set(1, _abf.MakeName("\u0046\u0069\u0074"))
	}
	return _cfdbg
}

// ColorFromPdfObjects gets the color from a series of pdf objects (3 for rgb).
func (_daac *PdfColorspaceDeviceRGB) ColorFromPdfObjects(objects []_abf.PdfObject) (PdfColor, error) {
	if len(objects) != 3 {
		return nil, _fd.New("r\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b")
	}
	_beedg, _gbdf := _abf.GetNumbersAsFloat(objects)
	if _gbdf != nil {
		return nil, _gbdf
	}
	return _daac.ColorFromFloats(_beedg)
}

// SetXObjectByName adds the XObject from the passed in stream to the page resources.
// The added XObject is identified by the specified name.
func (_becba *PdfPageResources) SetXObjectByName(keyName _abf.PdfObjectName, stream *_abf.PdfObjectStream) error {
	if _becba.XObject == nil {
		_becba.XObject = _abf.MakeDict()
	}
	_fefa := _abf.TraceToDirectObject(_becba.XObject)
	_dggba, _gadebf := _fefa.(*_abf.PdfObjectDictionary)
	if !_gadebf {
		_acd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0058\u004f\u0062j\u0065\u0063\u0074\u002c\u0020\u0067\u006f\u0074\u0020\u0025T\u002f\u0025\u0054", _becba.XObject, _fefa)
		return _fd.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_dggba.Set(keyName, stream)
	return nil
}

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_gcffb *PdfWriter) SetOptimizer(optimizer Optimizer) { _gcffb._cacbf = optimizer }

// String returns a string describing the font descriptor.
func (_efffg *PdfFontDescriptor) String() string {
	var _ddgff []string
	if _efffg.FontName != nil {
		_ddgff = append(_ddgff, _efffg.FontName.String())
	}
	if _efffg.FontFamily != nil {
		_ddgff = append(_ddgff, _efffg.FontFamily.String())
	}
	if _efffg.fontFile != nil {
		_ddgff = append(_ddgff, _efffg.fontFile.String())
	}
	if _efffg._fcdf != nil {
		_ddgff = append(_ddgff, _efffg._fcdf.String())
	}
	_ddgff = append(_ddgff, _e.Sprintf("\u0046\u006f\u006et\u0046\u0069\u006c\u0065\u0033\u003d\u0025\u0074", _efffg.FontFile3 != nil))
	return _e.Sprintf("\u0046\u004f\u004e\u0054_D\u0045\u0053\u0043\u0052\u0049\u0050\u0054\u004f\u0052\u007b\u0025\u0073\u007d", _be.Join(_ddgff, "\u002c\u0020"))
}

func _gabff(_aceb *_abf.PdfObjectDictionary) (*PdfShadingType6, error) {
	_ebdcd := PdfShadingType6{}
	_dcbgc := _aceb.Get("\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006f\u0072\u0064i\u006e\u0061\u0074\u0065")
	if _dcbgc == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064 \u0061\u0074\u0074\u0072\u0069\u0062\u0075t\u0065\u0020\u006d\u0069\u0073\u0073i\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072C\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_dfbde, _aggdf := _dcbgc.(*_abf.PdfObjectInteger)
	if !_aggdf {
		_acd.Log.Debug("\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _dcbgc)
		return nil, _abf.ErrTypeError
	}
	_ebdcd.BitsPerCoordinate = _dfbde
	_dcbgc = _aceb.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _dcbgc == nil {
		_acd.Log.Debug("\u0052e\u0071\u0075i\u0072\u0065\u0064\u0020a\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073in\u0067\u003a\u0020B\u0069\u0074s\u0050\u0065\u0072\u0043\u006f\u006dp\u006f\u006ee\u006e\u0074")
		return nil, ErrRequiredAttributeMissing
	}
	_dfbde, _aggdf = _dcbgc.(*_abf.PdfObjectInteger)
	if !_aggdf {
		_acd.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065r \u0028\u0067\u006ft\u0020%\u0054\u0029", _dcbgc)
		return nil, _abf.ErrTypeError
	}
	_ebdcd.BitsPerComponent = _dfbde
	_dcbgc = _aceb.Get("B\u0069\u0074\u0073\u0050\u0065\u0072\u0046\u006c\u0061\u0067")
	if _dcbgc == nil {
		_acd.Log.Debug("\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0042\u0069\u0074\u0073\u0050\u0065r\u0046\u006c\u0061\u0067")
		return nil, ErrRequiredAttributeMissing
	}
	_dfbde, _aggdf = _dcbgc.(*_abf.PdfObjectInteger)
	if !_aggdf {
		_acd.Log.Debug("B\u0069\u0074\u0073\u0050\u0065\u0072F\u006c\u0061\u0067\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025T\u0029", _dcbgc)
		return nil, _abf.ErrTypeError
	}
	_ebdcd.BitsPerComponent = _dfbde
	_dcbgc = _aceb.Get("\u0044\u0065\u0063\u006f\u0064\u0065")
	if _dcbgc == nil {
		_acd.Log.Debug("\u0052\u0065\u0071ui\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069b\u0075t\u0065 \u006di\u0073\u0073\u0069\u006e\u0067\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065")
		return nil, ErrRequiredAttributeMissing
	}
	_fdgg, _aggdf := _dcbgc.(*_abf.PdfObjectArray)
	if !_aggdf {
		_acd.Log.Debug("\u0044\u0065\u0063\u006fd\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061r\u0072a\u0079\u0020\u0028\u0067\u006f\u0074\u0020%\u0054\u0029", _dcbgc)
		return nil, _abf.ErrTypeError
	}
	_ebdcd.Decode = _fdgg
	if _cedb := _aceb.Get("\u0046\u0075\u006e\u0063\u0074\u0069\u006f\u006e"); _cedb != nil {
		_ebdcd.Function = []PdfFunction{}
		if _abfef, _beefa := _cedb.(*_abf.PdfObjectArray); _beefa {
			for _, _ecdf := range _abfef.Elements() {
				_bcffaa, _bbeea := _ebedg(_ecdf)
				if _bbeea != nil {
					_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _bbeea)
					return nil, _bbeea
				}
				_ebdcd.Function = append(_ebdcd.Function, _bcffaa)
			}
		} else {
			_dbbfd, _dcegf := _ebedg(_cedb)
			if _dcegf != nil {
				_acd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e:\u0020\u0025\u0076", _dcegf)
				return nil, _dcegf
			}
			_ebdcd.Function = append(_ebdcd.Function, _dbbfd)
		}
	}
	return &_ebdcd, nil
}

// NewPdfAnnotationSquiggly returns a new text squiggly annotation.
func NewPdfAnnotationSquiggly() *PdfAnnotationSquiggly {
	_gaf := NewPdfAnnotation()
	_bced := &PdfAnnotationSquiggly{}
	_bced.PdfAnnotation = _gaf
	_bced.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_gaf.SetContext(_bced)
	return _bced
}

// PdfFunction interface represents the common methods of a function in PDF.
type PdfFunction interface {
	Evaluate([]float64) ([]float64, error)
	ToPdfObject() _abf.PdfObject
}

func (_gegge PdfFont) actualFont() pdfFont {
	if _gegge._gedca == nil {
		_acd.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0061\u0063\u0074\u0075\u0061\u006c\u0046\u006f\u006e\u0074\u002e\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0069\u0073\u0020\u006e\u0069\u006c.\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0073", _gegge)
	}
	return _gegge._gedca
}

// NewPdfOutlineTree returns an initialized PdfOutline tree.
func NewPdfOutlineTree() *PdfOutline { _gced := NewPdfOutline(); _gced._aecec = &_gced; return _gced }

// ToPdfObject return the CalGray colorspace as a PDF object (name dictionary).
func (_gaeg *PdfColorspaceCalGray) ToPdfObject() _abf.PdfObject {
	_bab := &_abf.PdfObjectArray{}
	_bab.Append(_abf.MakeName("\u0043a\u006c\u0047\u0072\u0061\u0079"))
	_afeae := _abf.MakeDict()
	if _gaeg.WhitePoint != nil {
		_afeae.Set("\u0057\u0068\u0069\u0074\u0065\u0050\u006f\u0069\u006e\u0074", _abf.MakeArray(_abf.MakeFloat(_gaeg.WhitePoint[0]), _abf.MakeFloat(_gaeg.WhitePoint[1]), _abf.MakeFloat(_gaeg.WhitePoint[2])))
	} else {
		_acd.Log.Error("\u0043\u0061\u006c\u0047\u0072\u0061\u0079\u003a\u0020\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0057\u0068\u0069\u0074\u0065\u0050\u006fi\u006e\u0074\u0020\u0028\u0052e\u0071\u0075i\u0072\u0065\u0064\u0029")
	}
	if _gaeg.BlackPoint != nil {
		_afeae.Set("\u0042\u006c\u0061\u0063\u006b\u0050\u006f\u0069\u006e\u0074", _abf.MakeArray(_abf.MakeFloat(_gaeg.BlackPoint[0]), _abf.MakeFloat(_gaeg.BlackPoint[1]), _abf.MakeFloat(_gaeg.BlackPoint[2])))
	}
	_afeae.Set("\u0047\u0061\u006dm\u0061", _abf.MakeFloat(_gaeg.Gamma))
	_bab.Append(_afeae)
	if _gaeg._dgcg != nil {
		_gaeg._dgcg.PdfObject = _bab
		return _gaeg._dgcg
	}
	return _bab
}

// GetContainingPdfObject implements interface PdfModel.
func (_ecfg *PdfSignature) GetContainingPdfObject() _abf.PdfObject { return _ecfg._geebd }

// PdfActionURI represents an URI action.
type PdfActionURI struct {
	*PdfAction
	URI   _abf.PdfObject
	IsMap _abf.PdfObject
}

// GetXObjectFormByName returns the XObjectForm with the specified name from the
// page resources, if it exists.
func (_febca *PdfPageResources) GetXObjectFormByName(keyName _abf.PdfObjectName) (*XObjectForm, error) {
	_ecfcc, _cdcfg := _febca.GetXObjectByName(keyName)
	if _ecfcc == nil {
		return nil, nil
	}
	if _cdcfg != XObjectTypeForm {
		return nil, _fd.New("\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006f\u0072\u006d")
	}
	_ffagc, _adcab := NewXObjectFormFromStream(_ecfcc)
	if _adcab != nil {
		return nil, _adcab
	}
	return _ffagc, nil
}

// PdfShadingPatternType3 is shading patterns that will use a Type 3 shading pattern (Radial).
type PdfShadingPatternType3 struct {
	*PdfPattern
	Shading   *PdfShadingType3
	Matrix    *_abf.PdfObjectArray
	ExtGState _abf.PdfObject
}

// ReplacePage replaces the original page to a new page.
func (_deg *PdfAppender) ReplacePage(pageNum int, page *PdfPage) {
	_dfae := pageNum - 1
	for _degg := range _deg._cggfa {
		if _degg == _dfae {
			_dbae := page.Duplicate()
			_deg._cggfa[_degg] = _dbae
		}
	}
}

// NewPdfAnnotationPolygon returns a new polygon annotation.
func NewPdfAnnotationPolygon() *PdfAnnotationPolygon {
	_defb := NewPdfAnnotation()
	_dfgg := &PdfAnnotationPolygon{}
	_dfgg.PdfAnnotation = _defb
	_dfgg.PdfAnnotationMarkup = &PdfAnnotationMarkup{}
	_defb.SetContext(_dfgg)
	return _dfgg
}

// GetPdfInfo returns the PDF info dictionary.
func (_geada *PdfReader) GetPdfInfo() (*PdfInfo, error) {
	_cabg, _gbccae := _geada.GetTrailer()
	if _gbccae != nil {
		return nil, _gbccae
	}
	var _cefb *_abf.PdfObjectDictionary
	_faadd := _cabg.Get("\u0049\u006e\u0066\u006f")
	switch _cgbb := _faadd.(type) {
	case *_abf.PdfObjectReference:
		_gbedd := _cgbb
		_faadd, _gbccae = _geada.GetIndirectObjectByNumber(int(_gbedd.ObjectNumber))
		_faadd = _abf.TraceToDirectObject(_faadd)
		if _gbccae != nil {
			return nil, _gbccae
		}
		_cefb, _ = _faadd.(*_abf.PdfObjectDictionary)
	case *_abf.PdfObjectDictionary:
		_cefb = _cgbb
	}
	if _cefb == nil {
		return nil, _fd.New("I\u006e\u0066\u006f\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0070r\u0065s\u0065\u006e\u0074")
	}
	_caffe, _gbccae := NewPdfInfoFromObject(_cefb)
	if _gbccae != nil {
		return nil, _gbccae
	}
	return _caffe, nil
}

type pdfFontType0 struct {
	fontCommon
	_bgefb         *_abf.PdfIndirectObject
	_edeaf         _cbb.TextEncoder
	Encoding       _abf.PdfObject
	DescendantFont *PdfFont
	_fcfg          *_bd.CMap
}

// R returns the value of the red component of the color.
func (_ageea *PdfColorDeviceRGB) R() float64 { return _ageea[0] }

// ToPdfObject implements interface PdfModel.
func (_ggd *PdfAnnotationMovie) ToPdfObject() _abf.PdfObject {
	_ggd.PdfAnnotation.ToPdfObject()
	_baac := _ggd._dbc
	_abcba := _baac.PdfObject.(*_abf.PdfObjectDictionary)
	_abcba.SetIfNotNil("\u0053u\u0062\u0074\u0079\u0070\u0065", _abf.MakeName("\u004d\u006f\u0076i\u0065"))
	_abcba.SetIfNotNil("\u0054", _ggd.T)
	_abcba.SetIfNotNil("\u004d\u006f\u0076i\u0065", _ggd.Movie)
	_abcba.SetIfNotNil("\u0041", _ggd.A)
	return _baac
}

func _cega(_ggbca *_gca.ImageBase) (_fgcf Image) {
	_fgcf.Width = int64(_ggbca.Width)
	_fgcf.Height = int64(_ggbca.Height)
	_fgcf.BitsPerComponent = int64(_ggbca.BitsPerComponent)
	_fgcf.ColorComponents = _ggbca.ColorComponents
	_fgcf.Data = _ggbca.Data
	_fgcf._ceeag = _ggbca.Decode
	_fgcf._gedg = _ggbca.Alpha
	return _fgcf
}
func _gcag(_afacbc *fontCommon) *pdfFontType3 { return &pdfFontType3{fontCommon: *_afacbc} }

// DecodeArray returns the range of color component values in the Lab colorspace.
func (_feag *PdfColorspaceLab) DecodeArray() []float64 {
	_ddfa := []float64{0, 100}
	if _feag.Range != nil && len(_feag.Range) == 4 {
		_ddfa = append(_ddfa, _feag.Range...)
	} else {
		_ddfa = append(_ddfa, -100, 100, -100, 100)
	}
	return _ddfa
}

// SetPdfModifiedDate sets the ModDate attribute of the output PDF.
func SetPdfModifiedDate(modifiedDate _f.Time) {
	_gaabd.Lock()
	defer _gaabd.Unlock()
	_edfdc = modifiedDate
}

func (_aeae *LTV) generateVRIKey(_bcdbb *PdfSignature) (string, error) {
	_eaea, _bfac := _fdbbe(_bcdbb.Contents.Bytes())
	if _bfac != nil {
		return "", _bfac
	}
	return _be.ToUpper(_cb.EncodeToString(_eaea)), nil
}

// AddExtGState add External Graphics State (GState). The gsDict can be specified
// either directly as a dictionary or an indirect object containing a dictionary.
func (_gdgfb *PdfPageResources) AddExtGState(gsName _abf.PdfObjectName, gsDict _abf.PdfObject) error {
	if _gdgfb.ExtGState == nil {
		_gdgfb.ExtGState = _abf.MakeDict()
	}
	_ffaga := _gdgfb.ExtGState
	_beafe, _edfef := _abf.TraceToDirectObject(_ffaga).(*_abf.PdfObjectDictionary)
	if !_edfef {
		_acd.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0074\u0079\u0070\u0065\u0020e\u0072r\u006f\u0072\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u002f\u0025\u0054\u0029", _ffaga, _abf.TraceToDirectObject(_ffaga))
		return _abf.ErrTypeError
	}
	_beafe.Set(gsName, gsDict)
	return nil
}

func (_bgfcac *PdfSignature) extractChainFromCert() ([]*_fa.Certificate, error) {
	var _dbcgd *_abf.PdfObjectArray
	switch _edbdb := _bgfcac.Cert.(type) {
	case *_abf.PdfObjectString:
		_dbcgd = _abf.MakeArray(_edbdb)
	case *_abf.PdfObjectArray:
		_dbcgd = _edbdb
	default:
		return nil, _e.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _edbdb)
	}
	var _agadc _dd.Buffer
	for _, _gadbg := range _dbcgd.Elements() {
		_aafcba, _decgc := _abf.GetString(_gadbg)
		if !_decgc {
			return nil, _e.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054", _gadbg)
		}
		if _, _dbcda := _agadc.Write(_aafcba.Bytes()); _dbcda != nil {
			return nil, _dbcda
		}
	}
	return _fa.ParseCertificates(_agadc.Bytes())
}

// ToPdfObject implements interface PdfModel.
func (_dcc *PdfActionSound) ToPdfObject() _abf.PdfObject {
	_dcc.PdfAction.ToPdfObject()
	_da := _dcc._egg
	_bcf := _da.PdfObject.(*_abf.PdfObjectDictionary)
	_bcf.SetIfNotNil("\u0053", _abf.MakeName(string(ActionTypeSound)))
	_bcf.SetIfNotNil("\u0053\u006f\u0075n\u0064", _dcc.Sound)
	_bcf.SetIfNotNil("\u0056\u006f\u006c\u0075\u006d\u0065", _dcc.Volume)
	_bcf.SetIfNotNil("S\u0079\u006e\u0063\u0068\u0072\u006f\u006e\u006f\u0075\u0073", _dcc.Synchronous)
	_bcf.SetIfNotNil("\u0052\u0065\u0070\u0065\u0061\u0074", _dcc.Repeat)
	_bcf.SetIfNotNil("\u004d\u0069\u0078", _dcc.Mix)
	return _da
}

func (_bdc *PdfReader) newPdfActionGoTo3DViewFromDict(_cfaa *_abf.PdfObjectDictionary) (*PdfActionGoTo3DView, error) {
	return &PdfActionGoTo3DView{TA: _cfaa.Get("\u0054\u0041"), V: _cfaa.Get("\u0056")}, nil
}

// GetCharMetrics returns the char metrics for character code `code`.
func (_ffdb pdfFontType3) GetCharMetrics(code _cbb.CharCode) (_gbe.CharMetrics, bool) {
	if _bfca, _dffca := _ffdb._ecgf[code]; _dffca {
		return _gbe.CharMetrics{Wx: _bfca}, true
	}
	if _gbe.IsStdFont(_gbe.StdFontName(_ffdb._ecggf)) {
		return _gbe.CharMetrics{Wx: 250}, true
	}
	return _gbe.CharMetrics{}, false
}

// NewPdfColorPatternType2 returns an empty color shading pattern type 2 (Axial).
func NewPdfColorPatternType2() *PdfColorPatternType2 { _ggfa := &PdfColorPatternType2{}; return _ggfa }

// Image interface is a basic representation of an image used in PDF.
// The colorspace is not specified, but must be known when handling the image.
type Image struct {
	Width            int64
	Height           int64
	BitsPerComponent int64
	ColorComponents  int
	Data             []byte
	_gedg            []byte
	_ceeag           []float64
}

// GetXObjectByName returns the XObject with the specified keyName and the object type.
func (_efccg *PdfPageResources) GetXObjectByName(keyName _abf.PdfObjectName) (*_abf.PdfObjectStream, XObjectType) {
	if _efccg.XObject == nil {
		return nil, XObjectTypeUndefined
	}
	_acgec, _cdaa := _abf.TraceToDirectObject(_efccg.XObject).(*_abf.PdfObjectDictionary)
	if !_cdaa {
		_acd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0054\u0029", _abf.TraceToDirectObject(_efccg.XObject))
		return nil, XObjectTypeUndefined
	}
	if _daagfb := _acgec.Get(keyName); _daagfb != nil {
		_bbab, _afge := _abf.GetStream(_daagfb)
		if !_afge {
			_acd.Log.Debug("X\u004f\u0062\u006a\u0065\u0063\u0074 \u006e\u006f\u0074\u0020\u0070\u006fi\u006e\u0074\u0069\u006e\u0067\u0020\u0074o\u0020\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020%\u0054", _daagfb)
			return nil, XObjectTypeUndefined
		}
		_afcaa := _bbab.PdfObjectDictionary
		_ccgb, _afge := _abf.TraceToDirectObject(_afcaa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")).(*_abf.PdfObjectName)
		if !_afge {
			_acd.Log.Debug("\u0058\u004fbj\u0065\u0063\u0074 \u0053\u0075\u0062\u0074ype\u0020no\u0074\u0020\u0061\u0020\u004e\u0061\u006de,\u0020\u0064\u0069\u0063\u0074\u003a\u0020%\u0073", _afcaa.String())
			return nil, XObjectTypeUndefined
		}
		if *_ccgb == "\u0049\u006d\u0061g\u0065" {
			return _bbab, XObjectTypeImage
		} else if *_ccgb == "\u0046\u006f\u0072\u006d" {
			return _bbab, XObjectTypeForm
		} else if *_ccgb == "\u0050\u0053" {
			return _bbab, XObjectTypePS
		} else {
			_acd.Log.Debug("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0053\u0075b\u0074\u0079\u0070\u0065\u0020\u006e\u006ft\u0020\u006b\u006e\u006f\u0077\u006e\u0020\u0028\u0025\u0073\u0029", *_ccgb)
			return nil, XObjectTypeUndefined
		}
	} else {
		return nil, XObjectTypeUndefined
	}
}

// GetContext returns the annotation context which contains the specific type-dependent context.
// The context represents the subannotation.
func (_dabf *PdfAnnotation) GetContext() PdfModel {
	if _dabf == nil {
		return nil
	}
	return _dabf._edg
}
