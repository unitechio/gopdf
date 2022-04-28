package core

import (
	_acg "bufio"
	_fd "bytes"
	_a "compress/lzw"
	_bc "compress/zlib"
	_eae "crypto/md5"
	_ag "crypto/rand"
	_cab "encoding/hex"
	_d "errors"
	_ac "fmt"
	_ea "image"
	_ed "image/color"
	_f "image/jpeg"
	_dgf "io"
	_cg "io/ioutil"
	_c "reflect"
	_ba "regexp"
	_ca "sort"
	_dg "strconv"
	_agg "strings"
	_e "sync"
	_ga "time"
	_b "unicode"

	_ae "bitbucket.org/shenghui0779/gopdf/common"
	_geg "bitbucket.org/shenghui0779/gopdf/core/security"
	_ebd "bitbucket.org/shenghui0779/gopdf/core/security/crypt"
	_ge "bitbucket.org/shenghui0779/gopdf/internal/ccittfax"
	_ce "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_df "bitbucket.org/shenghui0779/gopdf/internal/jbig2"
	_gg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_gb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder"
	_da "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_eb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_ef "bitbucket.org/shenghui0779/gopdf/internal/strutils"
	_cge "golang.org/x/image/tiff/lzw"
	_bae "golang.org/x/xerrors"
)

// Len returns the number of elements in the array.
func (_ddaf *PdfObjectArray) Len() int {
	if _ddaf == nil {
		return 0
	}
	return len(_ddaf._fffab)
}

// IsWhiteSpace checks if byte represents a white space character.
func IsWhiteSpace(ch byte) bool {
	if (ch == 0x00) || (ch == 0x09) || (ch == 0x0A) || (ch == 0x0C) || (ch == 0x0D) || (ch == 0x20) {
		return true
	}
	return false
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_fbf *JPXEncoder) MakeDecodeParams() PdfObject { return nil }

type cryptFilters map[string]_ebd.Filter

var _abce = _ba.MustCompile("\u005e\u005b\\\u002b\u002d\u002e\u005d*\u0028\u005b0\u002d\u0039\u002e\u005d\u002b\u0029\u005b\u0065E\u005d\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

// NewDCTEncoder makes a new DCT encoder with default parameters.
func NewDCTEncoder() *DCTEncoder {
	_gddb := &DCTEncoder{}
	_gddb.ColorComponents = 3
	_gddb.BitsPerComponent = 8
	_gddb.Quality = DefaultJPEGQuality
	return _gddb
}

// GetFilterName returns the name of the encoding filter.
func (_gdfca *ASCIIHexEncoder) GetFilterName() string { return StreamEncodingFilterNameASCIIHex }

// HeaderCommentBytes gets the header comment bytes.
func (_bcbb ParserMetadata) HeaderCommentBytes() [4]byte { return _bcbb._fcbg }

var _dgec = _ba.MustCompile("\u005c\u0073\u002a\u0078\u0072\u0065\u0066\u005c\u0073\u002a")

// HasOddLengthHexStrings checks if the document has odd length hexadecimal strings.
func (_eead ParserMetadata) HasOddLengthHexStrings() bool { return _eead._gbcdc }

// GetIntVal returns the int value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetIntVal(obj PdfObject) (_bagcf int, _ggef bool) {
	_febb, _ggef := TraceToDirectObject(obj).(*PdfObjectInteger)
	if _ggef && _febb != nil {
		return int(*_febb), true
	}
	return 0, false
}
func _cbcf(_eecc *PdfObjectStream, _bda *PdfObjectDictionary) (*FlateEncoder, error) {
	_cfga := NewFlateEncoder()
	_ffgf := _eecc.PdfObjectDictionary
	if _ffgf == nil {
		return _cfga, nil
	}
	_cfga._dga = _eeaf(_ffgf)
	if _bda == nil {
		_cdfa := TraceToDirectObject(_ffgf.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		switch _acf := _cdfa.(type) {
		case *PdfObjectArray:
			if _acf.Len() != 1 {
				_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020\u0028\u0025\u0064\u0029", _acf.Len())
				return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			if _bff, _egbe := GetDict(_acf.Get(0)); _egbe {
				_bda = _bff
			}
		case *PdfObjectDictionary:
			_bda = _acf
		case *PdfObjectNull, nil:
		default:
			_ae.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _cdfa)
			return nil, _ac.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _bda == nil {
		return _cfga, nil
	}
	_ae.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _bda.String())
	_ccdg := _bda.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _ccdg == nil {
		_ae.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_dfbf, _edcc := _ccdg.(*PdfObjectInteger)
		if !_edcc {
			_ae.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _ccdg)
			return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_cfga.Predictor = int(*_dfbf)
	}
	_ccdg = _bda.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _ccdg != nil {
		_abge, _ffec := _ccdg.(*PdfObjectInteger)
		if !_ffec {
			_ae.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ac.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_cfga.BitsPerComponent = int(*_abge)
	}
	if _cfga.Predictor > 1 {
		_cfga.Columns = 1
		_ccdg = _bda.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _ccdg != nil {
			_ecf, _fac := _ccdg.(*PdfObjectInteger)
			if !_fac {
				return nil, _ac.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_cfga.Columns = int(*_ecf)
		}
		_cfga.Colors = 1
		_ccdg = _bda.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _ccdg != nil {
			_dbd, _ebfb := _ccdg.(*PdfObjectInteger)
			if !_ebfb {
				return nil, _ac.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_cfga.Colors = int(*_dbd)
		}
	}
	return _cfga, nil
}

// Update updates multiple keys and returns the dictionary back so can be used in a chained fashion.
func (_bbdf *PdfObjectDictionary) Update(objmap map[string]PdfObject) *PdfObjectDictionary {
	_bbdf._efce.Lock()
	defer _bbdf._efce.Unlock()
	for _fbbg, _dadaa := range objmap {
		_bbdf.setWithLock(PdfObjectName(_fbbg), _dadaa, false)
	}
	return _bbdf
}

// NewFlateEncoder makes a new flate encoder with default parameters, predictor 1 and bits per component 8.
func NewFlateEncoder() *FlateEncoder {
	_cggb := &FlateEncoder{}
	_cggb.Predictor = 1
	_cggb.BitsPerComponent = 8
	_cggb.Colors = 1
	_cggb.Columns = 1
	return _cggb
}

// WriteString outputs the object as it is to be written to file.
func (_cbefe *PdfObjectBool) WriteString() string {
	if *_cbefe {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}
func (_dcf *PdfCrypt) loadCryptFilters(_bgbf *PdfObjectDictionary) error {
	_dcf._agfd = cryptFilters{}
	_dabd := _bgbf.Get("\u0043\u0046")
	_dabd = TraceToDirectObject(_dabd)
	if _dgeb, _ebe := _dabd.(*PdfObjectReference); _ebe {
		_ecb, _ccfe := _dcf._cdg.LookupByReference(*_dgeb)
		if _ccfe != nil {
			_ae.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0075\u0070\u0020\u0043\u0046\u0020\u0072\u0065\u0066\u0065\u0072en\u0063\u0065")
			return _ccfe
		}
		_dabd = TraceToDirectObject(_ecb)
	}
	_fgc, _bfa := _dabd.(*PdfObjectDictionary)
	if !_bfa {
		_ae.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0043\u0046\u002c \u0074\u0079\u0070\u0065: \u0025\u0054", _dabd)
		return _d.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0046")
	}
	for _, _ddfb := range _fgc.Keys() {
		_cac := _fgc.Get(_ddfb)
		if _fea, _deb := _cac.(*PdfObjectReference); _deb {
			_bbaf, _cbd := _dcf._cdg.LookupByReference(*_fea)
			if _cbd != nil {
				_ae.Log.Debug("\u0045\u0072ro\u0072\u0020\u006co\u006f\u006b\u0075\u0070 up\u0020di\u0063\u0074\u0069\u006f\u006e\u0061\u0072y \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065")
				return _cbd
			}
			_cac = TraceToDirectObject(_bbaf)
		}
		_agcg, _cfa := _cac.(*PdfObjectDictionary)
		if !_cfa {
			return _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006e \u0043\u0046\u0020\u0028\u006e\u0061\u006d\u0065\u0020\u0025\u0073\u0029\u0020-\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0062\u0075\u0074\u0020\u0025\u0054", _ddfb, _cac)
		}
		if _ddfb == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u002d\u0020\u0043\u0061\u006e\u006e\u006f\u0074\u0020\u006f\u0076\u0065\u0072\u0077r\u0069\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0069d\u0065\u006e\u0074\u0069\u0074\u0079\u0020\u0066\u0069\u006c\u0074\u0065\u0072 \u002d\u0020\u0054\u0072\u0079\u0069n\u0067\u0020\u006ee\u0078\u0074")
			continue
		}
		var _gbc _ebd.FilterDict
		if _cacb := _afe(&_gbc, _agcg); _cacb != nil {
			return _cacb
		}
		_cgb, _ebb := _ebd.NewFilter(_gbc)
		if _ebb != nil {
			return _ebb
		}
		_dcf._agfd[string(_ddfb)] = _cgb
	}
	_dcf._agfd["\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"] = _ebd.NewIdentity()
	_dcf._dee = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _bbg, _ebee := _bgbf.Get("\u0053\u0074\u0072\u0046").(*PdfObjectName); _ebee {
		if _, _fgaf := _dcf._agfd[string(*_bbg)]; !_fgaf {
			return _ac.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u0072\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_bbg)
		}
		_dcf._dee = string(*_bbg)
	}
	_dcf._bba = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _cbef, _dcaa := _bgbf.Get("\u0053\u0074\u006d\u0046").(*PdfObjectName); _dcaa {
		if _, _bfac := _dcf._agfd[string(*_cbef)]; !_bfac {
			return _ac.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u006d\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_cbef)
		}
		_dcf._bba = string(*_cbef)
	}
	return nil
}

// DecodeStream decodes a JBIG2 encoded stream and returns the result as a slice of bytes.
func (_eafg *JBIG2Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _eafg.DecodeBytes(streamObj.Stream)
}
func (_dbgd *PdfParser) xrefNextObjectOffset(_gffde int64) int64 {
	_bccd := int64(0)
	if len(_dbgd._fbab.ObjectMap) == 0 {
		return 0
	}
	if len(_dbgd._fbab._bd) == 0 {
		_bbfg := 0
		for _, _gccf := range _dbgd._fbab.ObjectMap {
			if _gccf.Offset > 0 {
				_bbfg++
			}
		}
		if _bbfg == 0 {
			return 0
		}
		_dbgd._fbab._bd = make([]XrefObject, _bbfg)
		_cdce := 0
		for _, _egac := range _dbgd._fbab.ObjectMap {
			if _egac.Offset > 0 {
				_dbgd._fbab._bd[_cdce] = _egac
				_cdce++
			}
		}
		_ca.Slice(_dbgd._fbab._bd, func(_fcfe, _gced int) bool { return _dbgd._fbab._bd[_fcfe].Offset < _dbgd._fbab._bd[_gced].Offset })
	}
	_gcfea := _ca.Search(len(_dbgd._fbab._bd), func(_agfg int) bool { return _dbgd._fbab._bd[_agfg].Offset >= _gffde })
	if _gcfea < len(_dbgd._fbab._bd) {
		_bccd = _dbgd._fbab._bd[_gcfea].Offset
	}
	return _bccd
}

// Validate validates the page settings for the JBIG2 encoder.
func (_beff JBIG2EncoderSettings) Validate() error {
	const _gddg = "\u0076a\u006ci\u0064\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065\u0072"
	if _beff.Threshold < 0 || _beff.Threshold > 1.0 {
		return _eb.Errorf(_gddg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064\u0020\u0076a\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0076\u0027 \u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061n\u0067\u0065\u0020\u005b\u0030\u002e0\u002c\u0020\u0031.\u0030\u005d", _beff.Threshold)
	}
	if _beff.ResolutionX < 0 {
		return _eb.Errorf(_gddg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0078\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _beff.ResolutionX)
	}
	if _beff.ResolutionY < 0 {
		return _eb.Errorf(_gddg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0079\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _beff.ResolutionY)
	}
	if _beff.DefaultPixelValue != 0 && _beff.DefaultPixelValue != 1 {
		return _eb.Errorf(_gddg, "de\u0066\u0061u\u006c\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072 \u0074\u0068\u0065\u0020\u0062\u0069\u0074\u003a \u007b0\u002c\u0031}", _beff.DefaultPixelValue)
	}
	if _beff.Compression != JB2Generic {
		return _eb.Errorf(_gddg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0063\u006fm\u0070\u0072\u0065\u0073s\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	}
	return nil
}

// PdfObjectStreams represents the primitive PDF object streams.
// 7.5.7 Object Streams (page 45).
type PdfObjectStreams struct {
	PdfObjectReference
	_dgebc []PdfObject
}

// EncodeBytes encodes the passed in slice of bytes by passing it through the
// EncodeBytes method of the underlying encoders.
func (_eaf *MultiEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_bagd := data
	var _gfd error
	for _ebbb := len(_eaf._bfacg) - 1; _ebbb >= 0; _ebbb-- {
		_egfa := _eaf._bfacg[_ebbb]
		_bagd, _gfd = _egfa.EncodeBytes(_bagd)
		if _gfd != nil {
			return nil, _gfd
		}
	}
	return _bagd, nil
}
func _cgbc(_bdb *PdfObjectStream, _gab *PdfObjectDictionary) (*LZWEncoder, error) {
	_gdba := NewLZWEncoder()
	_dbad := _bdb.PdfObjectDictionary
	if _dbad == nil {
		return _gdba, nil
	}
	if _gab == nil {
		_dcd := TraceToDirectObject(_dbad.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _dcd != nil {
			if _gfgfd, _cfd := _dcd.(*PdfObjectDictionary); _cfd {
				_gab = _gfgfd
			} else if _bac, _fcc := _dcd.(*PdfObjectArray); _fcc {
				if _bac.Len() == 1 {
					if _bddg, _gccg := GetDict(_bac.Get(0)); _gccg {
						_gab = _bddg
					}
				}
			}
			if _gab == nil {
				_ae.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _dcd)
				return nil, _ac.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
	}
	_dccb := _dbad.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _dccb != nil {
		_ddg, _agad := _dccb.(*PdfObjectInteger)
		if !_agad {
			_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _dccb)
			return nil, _ac.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_ddg != 0 && *_ddg != 1 {
			return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_gdba.EarlyChange = int(*_ddg)
	} else {
		_gdba.EarlyChange = 1
	}
	if _gab == nil {
		return _gdba, nil
	}
	if _bag, _cggf := GetIntVal(_gab.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")); _cggf {
		if _bag == 0 || _bag == 1 {
			_gdba.EarlyChange = _bag
		} else {
			_ae.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020E\u0061\u0072\u006c\u0079\u0043\u0068\u0061n\u0067\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020%\u0064", _bag)
		}
	}
	_dccb = _gab.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _dccb != nil {
		_efcg, _gege := _dccb.(*PdfObjectInteger)
		if !_gege {
			_ae.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _dccb)
			return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_gdba.Predictor = int(*_efcg)
	}
	_dccb = _gab.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _dccb != nil {
		_fgcc, _bacc := _dccb.(*PdfObjectInteger)
		if !_bacc {
			_ae.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ac.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_gdba.BitsPerComponent = int(*_fgcc)
	}
	if _gdba.Predictor > 1 {
		_gdba.Columns = 1
		_dccb = _gab.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _dccb != nil {
			_ebed, _eggcf := _dccb.(*PdfObjectInteger)
			if !_eggcf {
				return nil, _ac.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_gdba.Columns = int(*_ebed)
		}
		_gdba.Colors = 1
		_dccb = _gab.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _dccb != nil {
			_beda, _ebda := _dccb.(*PdfObjectInteger)
			if !_ebda {
				return nil, _ac.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_gdba.Colors = int(*_beda)
		}
	}
	_ae.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _gab.String())
	return _gdba, nil
}

// MakeNull creates an PdfObjectNull.
func MakeNull() *PdfObjectNull { _addfc := PdfObjectNull{}; return &_addfc }
func _dgfa(_daa _ebd.Filter, _dda _geg.AuthEvent) *PdfObjectDictionary {
	if _dda == "" {
		_dda = _geg.EventDocOpen
	}
	_dbf := MakeDict()
	_dbf.Set("\u0054\u0079\u0070\u0065", MakeName("C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072"))
	_dbf.Set("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et", MakeName(string(_dda)))
	_dbf.Set("\u0043\u0046\u004d", MakeName(_daa.Name()))
	_dbf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_daa.KeyLength())))
	return _dbf
}

// PdfObjectArray represents the primitive PDF array object.
type PdfObjectArray struct{ _fffab []PdfObject }

// MultiEncoder supports serial encoding.
type MultiEncoder struct{ _bfacg []StreamEncoder }

// Encode encodes previously prepare jbig2 document and stores it as the byte slice.
func (_fbdb *JBIG2Encoder) Encode() (_faagf []byte, _dgaf error) {
	const _afcb = "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	if _fbdb._ebac == nil {
		return nil, _eb.Errorf(_afcb, "\u0064\u006f\u0063u\u006d\u0065\u006e\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0064a\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_fbdb._ebac.FullHeaders = _fbdb.DefaultPageSettings.FileMode
	_faagf, _dgaf = _fbdb._ebac.Encode()
	if _dgaf != nil {
		return nil, _eb.Wrap(_dgaf, _afcb, "")
	}
	return _faagf, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_feaf *ASCII85Encoder) MakeDecodeParams() PdfObject { return nil }
func (_feba *PdfParser) parseXref() (*PdfObjectDictionary, error) {
	const _bgaa = 20
	_dcdb, _ := _feba._eecea.Peek(_bgaa)
	for _ecbf := 0; _ecbf < 2; _ecbf++ {
		if _feba._egfag == 0 {
			_feba._egfag = _feba.GetFileOffset()
		}
		if _ddce.Match(_dcdb) {
			_ae.Log.Trace("\u0078\u0072e\u0066\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002e\u0020\u0050\u0072\u006f\u0062\u0061\u0062\u006c\u0079\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			_ae.Log.Debug("\u0073t\u0061r\u0074\u0069\u006e\u0067\u0020w\u0069\u0074h\u0020\u0022\u0025\u0073\u0022", string(_dcdb))
			return _feba.parseXrefStream(nil)
		}
		if _dgec.Match(_dcdb) {
			_ae.Log.Trace("\u0053\u0074\u0061\u006ed\u0061\u0072\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073e\u0063t\u0069\u006f\u006e\u0020\u0074\u0061\u0062l\u0065\u0021")
			return _feba.parseXrefTable()
		}
		_cagb := _feba.GetFileOffset()
		if _feba._egfag == 0 {
			_feba._egfag = _cagb
		}
		_feba.SetFileOffset(_cagb - _bgaa)
		defer _feba.SetFileOffset(_cagb)
		_gdea, _ := _feba._eecea.Peek(_bgaa)
		_dcdb = append(_gdea, _dcdb...)
	}
	_ae.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u0078\u0072\u0065f\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006fr\u0020\u0073\u0074\u0072\u0065\u0061\u006d.\u0020\u0052\u0065\u0070\u0061i\u0072\u0020\u0061\u0074\u0074e\u006d\u0070\u0074\u0065\u0064\u003a\u0020\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0065\u0061\u0072\u006c\u0069\u0065\u0073\u0074\u0020x\u0072\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u0062\u006f\u0074to\u006d\u002e")
	if _aebca := _feba.repairSeekXrefMarker(); _aebca != nil {
		_ae.Log.Debug("\u0052e\u0070a\u0069\u0072\u0020\u0066\u0061i\u006c\u0065d\u0020\u002d\u0020\u0025\u0076", _aebca)
		return nil, _aebca
	}
	return _feba.parseXrefTable()
}

// AddEncoder adds the passed in encoder to the underlying encoder slice.
func (_efff *MultiEncoder) AddEncoder(encoder StreamEncoder) {
	_efff._bfacg = append(_efff._bfacg, encoder)
}
func (_gccc *limitedReadSeeker) getError(_becca int64) error {
	switch {
	case _becca < 0:
		return _ac.Errorf("\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0065\u0067\u0061\u0074\u0069\u0076e\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0064", _becca)
	case _becca > _gccc._gfcg:
		return _ac.Errorf("u\u006e\u0065\u0078\u0070ec\u0074e\u0064\u0020\u006f\u0066\u0066s\u0065\u0074\u003a\u0020\u0025\u0064", _becca)
	}
	return nil
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_gcga *FlateEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _gcga.Predictor != 1 && _gcga.Predictor != 11 {
		_ae.Log.Debug("E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0046\u006c\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065r\u0020P\u0072\u0065\u0064\u0069c\u0074\u006fr\u0020\u003d\u0020\u0031\u002c\u0020\u0031\u0031\u0020\u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		return nil, ErrUnsupportedEncodingParameters
	}
	if _gcga.Predictor == 11 {
		_cgfe := _gcga.Columns
		_aabe := len(data) / _cgfe
		if len(data)%_cgfe != 0 {
			_ae.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
			return nil, _d.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
		}
		_geffg := _fd.NewBuffer(nil)
		_bafg := make([]byte, _cgfe)
		for _ebfd := 0; _ebfd < _aabe; _ebfd++ {
			_ebef := data[_cgfe*_ebfd : _cgfe*(_ebfd+1)]
			_bafg[0] = _ebef[0]
			for _afb := 1; _afb < _cgfe; _afb++ {
				_bafg[_afb] = byte(int(_ebef[_afb]-_ebef[_afb-1]) % 256)
			}
			_geffg.WriteByte(1)
			_geffg.Write(_bafg)
		}
		data = _geffg.Bytes()
	}
	var _abf _fd.Buffer
	_ade := _bc.NewWriter(&_abf)
	_ade.Write(data)
	_ade.Close()
	return _abf.Bytes(), nil
}

var _bffa = _ba.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064\u002b)\u005c\u0073\u002a\u0024")

// HasDataAfterEOF checks if there is some data after EOF marker.
func (_dfge ParserMetadata) HasDataAfterEOF() bool { return _dfge._fdgg }
func (_adfa *PdfCrypt) isDecrypted(_adg PdfObject) bool {
	_, _ddaa := _adfa._dgd[_adg]
	if _ddaa {
		_ae.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0064\u0065\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	switch _cgd := _adg.(type) {
	case *PdfObjectStream:
		if _adfa._gge.R != 5 {
			if _aff, _bgf := _cgd.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _bgf && *_aff == "\u0058\u0052\u0065\u0066" {
				return true
			}
		}
	case *PdfIndirectObject:
		if _, _ddaa = _adfa._dfg[int(_cgd.ObjectNumber)]; _ddaa {
			return true
		}
		switch _ggb := _cgd.PdfObject.(type) {
		case *PdfObjectDictionary:
			_gfgf := true
			for _, _fcf := range _ccd {
				if _ggb.Get(_fcf) == nil {
					_gfgf = false
					break
				}
			}
			if _gfgf {
				return true
			}
		}
	}
	_ae.Log.Trace("\u004e\u006f\u0074\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}

// AddPageImage adds the page with the image 'img' to the encoder context in order to encode it jbig2 document.
// The 'settings' defines what encoding type should be used by the encoder.
func (_edd *JBIG2Encoder) AddPageImage(img *JBIG2Image, settings *JBIG2EncoderSettings) (_dag error) {
	const _feab = "\u004a\u0042\u0049\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065\u0049m\u0061\u0067\u0065"
	if _edd == nil {
		return _eb.Error(_feab, "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u0069\u0073 n\u0069\u006c")
	}
	if settings == nil {
		settings = &_edd.DefaultPageSettings
	}
	if _edd._ebac == nil {
		_edd._ebac = _da.InitEncodeDocument(settings.FileMode)
	}
	if _dag = settings.Validate(); _dag != nil {
		return _eb.Wrap(_dag, _feab, "")
	}
	_bebc, _dag := img.toBitmap()
	if _dag != nil {
		return _eb.Wrap(_dag, _feab, "")
	}
	switch settings.Compression {
	case JB2Generic:
		if _dag = _edd._ebac.AddGenericPage(_bebc, settings.DuplicatedLinesRemoval); _dag != nil {
			return _eb.Wrap(_dag, _feab, "")
		}
	case JB2SymbolCorrelation:
		return _eb.Error(_feab, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return _eb.Error(_feab, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return _eb.Error(_feab, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return nil
}

// EncodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_gcdb *RawEncoder) EncodeBytes(data []byte) ([]byte, error) { return data, nil }

// WriteString outputs the object as it is to be written to file.
func (_ffdfe *PdfObjectStream) WriteString() string {
	var _eaeg _agg.Builder
	_eaeg.WriteString(_dg.FormatInt(_ffdfe.ObjectNumber, 10))
	_eaeg.WriteString("\u0020\u0030\u0020\u0052")
	return _eaeg.String()
}
func (_bdee *PdfParser) getNumbersOfUpdatedObjects(_ddcec *PdfParser) ([]int, error) {
	if _ddcec == nil {
		return nil, _d.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_dcgc := _ddcec._fcca
	_addg := make([]int, 0)
	_fbbb := make(map[int]interface{})
	_dega := make(map[int]int64)
	for _addfa, _dbcg := range _bdee._fbab.ObjectMap {
		if _dbcg.Offset == 0 {
			if _dbcg.OsObjNumber != 0 {
				if _ccgd, _dgafc := _bdee._fbab.ObjectMap[_dbcg.OsObjNumber]; _dgafc {
					_fbbb[_dbcg.OsObjNumber] = struct{}{}
					_dega[_addfa] = _ccgd.Offset
				} else {
					return nil, _d.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0078r\u0065\u0066\u0020\u0074ab\u006c\u0065")
				}
			}
		} else {
			_dega[_addfa] = _dbcg.Offset
		}
	}
	for _bcca, _beee := range _dega {
		if _, _fdae := _fbbb[_bcca]; _fdae {
			continue
		}
		if _beee > _dcgc {
			_addg = append(_addg, _bcca)
		}
	}
	return _addg, nil
}

// ReadBytesAt reads byte content at specific offset and length within the PDF.
func (_ced *PdfParser) ReadBytesAt(offset, len int64) ([]byte, error) {
	_fbec := _ced.GetFileOffset()
	_, _aaab := _ced._fdee.Seek(offset, _dgf.SeekStart)
	if _aaab != nil {
		return nil, _aaab
	}
	_affa := make([]byte, len)
	_, _aaab = _dgf.ReadAtLeast(_ced._fdee, _affa, int(len))
	if _aaab != nil {
		return nil, _aaab
	}
	_ced.SetFileOffset(_fbec)
	return _affa, nil
}

// GetFilterName returns the name of the encoding filter.
func (_baef *LZWEncoder) GetFilterName() string { return StreamEncodingFilterNameLZW }

// String returns a string representation of `name`.
func (_gecee *PdfObjectName) String() string { return string(*_gecee) }

var _cgbea = _ba.MustCompile("\u005b\\\u0072\u005c\u006e\u005d\u005c\u0073\u002a\u0028\u0078\u0072\u0065f\u0029\u005c\u0073\u002a\u005b\u005c\u0072\u005c\u006e\u005d")

const _fdbde = 32 << (^uint(0) >> 63)

// Decoded returns the PDFDocEncoding or UTF-16BE decoded string contents.
// UTF-16BE is applied when the first two bytes are 0xFE, 0XFF, otherwise decoding of
// PDFDocEncoding is performed.
func (_bgfae *PdfObjectString) Decoded() string {
	if _bgfae == nil {
		return ""
	}
	_ggdg := []byte(_bgfae._eeee)
	if len(_ggdg) >= 2 && _ggdg[0] == 0xFE && _ggdg[1] == 0xFF {
		return _ef.UTF16ToString(_ggdg[2:])
	}
	return _ef.PDFDocEncodingToString(_ggdg)
}

// GetFilterName returns the name of the encoding filter.
func (_gcgf *FlateEncoder) GetFilterName() string { return StreamEncodingFilterNameFlate }

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_caf *RawEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }
func (_gfbfd *PdfParser) skipSpaces() (int, error) {
	_ffaca := 0
	for {
		_cace, _ecdg := _gfbfd._eecea.ReadByte()
		if _ecdg != nil {
			return 0, _ecdg
		}
		if IsWhiteSpace(_cace) {
			_ffaca++
		} else {
			_gfbfd._eecea.UnreadByte()
			break
		}
	}
	return _ffaca, nil
}

// DecodeStream decodes a FlateEncoded stream object and give back decoded bytes.
func (_fgad *FlateEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_ae.Log.Trace("\u0046l\u0061t\u0065\u0044\u0065\u0063\u006fd\u0065\u0020s\u0074\u0072\u0065\u0061\u006d")
	_ae.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _fgad.Predictor)
	if _fgad.BitsPerComponent != 8 {
		return nil, _ac.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003d\u0025\u0064\u0020\u0028\u006f\u006e\u006c\u0079\u0020\u0038\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0029", _fgad.BitsPerComponent)
	}
	_deeb, _cbaa := _fgad.DecodeBytes(streamObj.Stream)
	if _cbaa != nil {
		return nil, _cbaa
	}
	_deeb, _cbaa = _fgad.postDecodePredict(_deeb)
	if _cbaa != nil {
		return nil, _cbaa
	}
	return _deeb, nil
}
func (_ec *PdfParser) lookupObjectViaOS(_fdc int, _bb int) (PdfObject, error) {
	var _gd *_fd.Reader
	var _cf objectStream
	var _de bool
	_cf, _de = _ec._acab[_fdc]
	if !_de {
		_gf, _cd := _ec.LookupByNumber(_fdc)
		if _cd != nil {
			_ae.Log.Debug("\u004d\u0069ss\u0069\u006e\u0067 \u006f\u0062\u006a\u0065ct \u0073tr\u0065\u0061\u006d\u0020\u0077\u0069\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u0020%\u0064", _fdc)
			return nil, _cd
		}
		_agd, _aca := _gf.(*PdfObjectStream)
		if !_aca {
			return nil, _d.New("i\u006e\u0076\u0061\u006cid\u0020o\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _ec._bffd != nil && !_ec._bffd.isDecrypted(_agd) {
			return nil, _d.New("\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0064\u0065\u0063r\u0079\u0070\u0074\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d")
		}
		_aa := _agd.PdfObjectDictionary
		_ae.Log.Trace("\u0073o\u0020\u0064\u003a\u0020\u0025\u0073\n", _aa.String())
		_dad, _aca := _aa.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
		if !_aca {
			_ae.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0061\u006c\u0077\u0061\u0079\u0073\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054\u0079\u0070\u0065")
			return nil, _d.New("\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020T\u0079\u0070\u0065")
		}
		if _agg.ToLower(string(*_dad)) != "\u006f\u0062\u006a\u0073\u0074\u006d" {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0074\u0079\u0070\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0077\u0061\u0079\u0073 \u0062\u0065\u0020\u004f\u0062\u006a\u0053\u0074\u006d\u0020\u0021")
			return nil, _d.New("\u006f\u0062\u006a\u0065c\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074y\u0070e\u0020\u0021\u003d\u0020\u004f\u0062\u006aS\u0074\u006d")
		}
		N, _aca := _aa.Get("\u004e").(*PdfObjectInteger)
		if !_aca {
			return nil, _d.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020i\u006e\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_aaf, _aca := _aa.Get("\u0046\u0069\u0072s\u0074").(*PdfObjectInteger)
		if !_aca {
			return nil, _d.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0046\u0069\u0072\u0073\u0074\u0020i\u006e \u0073t\u0072e\u0061\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_ae.Log.Trace("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _dad, *N)
		_af, _cd := DecodeStream(_agd)
		if _cd != nil {
			return nil, _cd
		}
		_ae.Log.Trace("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _af)
		_ad := _ec.GetFileOffset()
		defer func() { _ec.SetFileOffset(_ad) }()
		_gd = _fd.NewReader(_af)
		_ec._eecea = _acg.NewReader(_gd)
		_ae.Log.Trace("\u0050a\u0072s\u0069\u006e\u0067\u0020\u006ff\u0066\u0073e\u0074\u0020\u006d\u0061\u0070")
		_cc := map[int]int64{}
		for _deg := 0; _deg < int(*N); _deg++ {
			_ec.skipSpaces()
			_eg, _acd := _ec.parseNumber()
			if _acd != nil {
				return nil, _acd
			}
			_ceb, _fb := _eg.(*PdfObjectInteger)
			if !_fb {
				return nil, _d.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_ec.skipSpaces()
			_eg, _acd = _ec.parseNumber()
			if _acd != nil {
				return nil, _acd
			}
			_fe, _fb := _eg.(*PdfObjectInteger)
			if !_fb {
				return nil, _d.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_ae.Log.Trace("\u006f\u0062j\u0020\u0025\u0064 \u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0025\u0064", *_ceb, *_fe)
			_cc[int(*_ceb)] = int64(*_aaf + *_fe)
		}
		_cf = objectStream{N: int(*N), _fg: _af, _db: _cc}
		_ec._acab[_fdc] = _cf
	} else {
		_eaef := _ec.GetFileOffset()
		defer func() { _ec.SetFileOffset(_eaef) }()
		_gd = _fd.NewReader(_cf._fg)
		_ec._eecea = _acg.NewReader(_gd)
	}
	_fc := _cf._db[_bb]
	_ae.Log.Trace("\u0041\u0043\u0054\u0055AL\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u005b\u0025\u0064\u005d\u0020\u003d\u0020%\u0064", _bb, _fc)
	_gd.Seek(_fc, _dgf.SeekStart)
	_ec._eecea = _acg.NewReader(_gd)
	_eaa, _ := _ec._eecea.Peek(100)
	_ae.Log.Trace("\u004f\u0042\u004a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_eaa))
	_egf, _ab := _ec.parseObject()
	if _ab != nil {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0072\u0065\u0061\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _ab)
		return nil, _ab
	}
	if _egf == nil {
		return nil, _d.New("o\u0062\u006a\u0065\u0063t \u0063a\u006e\u006e\u006f\u0074\u0020b\u0065\u0020\u006e\u0075\u006c\u006c")
	}
	_bf := PdfIndirectObject{}
	_bf.ObjectNumber = int64(_bb)
	_bf.PdfObject = _egf
	_bf._ffgd = _ec
	return &_bf, nil
}

// EncryptInfo contains an information generated by the document encrypter.
type EncryptInfo struct {
	Version

	// Encrypt is an encryption dictionary that contains all necessary parameters.
	// It should be stored in all copies of the document trailer.
	Encrypt *PdfObjectDictionary

	// ID0 and ID1 are IDs used in the trailer. Older algorithms such as RC4 uses them for encryption.
	ID0, ID1 string
}

var _ddce = _ba.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

// GetFilterArray returns the names of the underlying encoding filters in an array that
// can be used as /Filter entry.
func (_ddfe *MultiEncoder) GetFilterArray() *PdfObjectArray {
	_fffb := make([]PdfObject, len(_ddfe._bfacg))
	for _dacf, _cfec := range _ddfe._bfacg {
		_fffb[_dacf] = MakeName(_cfec.GetFilterName())
	}
	return MakeArray(_fffb...)
}
func (_bcgg *PdfParser) traceStreamLength(_cefe PdfObject) (PdfObject, error) {
	_bcce, _dggg := _cefe.(*PdfObjectReference)
	if _dggg {
		_dgaa, _fgcag := _bcgg._gfdf[_bcce.ObjectNumber]
		if _fgcag && _dgaa {
			_ae.Log.Debug("\u0053t\u0072\u0065a\u006d\u0020\u004c\u0065n\u0067\u0074\u0068 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 u\u006e\u0072\u0065s\u006f\u006cv\u0065\u0064\u0020\u0028\u0069\u006cl\u0065\u0067a\u006c\u0029")
			return nil, _d.New("\u0069\u006c\u006c\u0065ga\u006c\u0020\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006fo\u0070")
		}
		_bcgg._gfdf[_bcce.ObjectNumber] = true
	}
	_gbbaf, _bggdg := _bcgg.Resolve(_cefe)
	if _bggdg != nil {
		return nil, _bggdg
	}
	_ae.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _gbbaf)
	if _dggg {
		_bcgg._gfdf[_bcce.ObjectNumber] = false
	}
	return _gbbaf, nil
}

// Bytes returns the PdfObjectString content as a []byte array.
func (_caaga *PdfObjectString) Bytes() []byte { return []byte(_caaga._eeee) }
func (_aegd *PdfParser) parseHexString() (*PdfObjectString, error) {
	_aegd._eecea.ReadByte()
	var _ageg _fd.Buffer
	for {
		_dbc, _egdba := _aegd._eecea.Peek(1)
		if _egdba != nil {
			return MakeString(""), _egdba
		}
		if _dbc[0] == '>' {
			_aegd._eecea.ReadByte()
			break
		}
		_bbga, _ := _aegd._eecea.ReadByte()
		if _aegd._ccce {
			if _fd.IndexByte(_bfce, _bbga) == -1 {
				_aegd._aecec._bbec = true
			}
		}
		if !IsWhiteSpace(_bbga) {
			_ageg.WriteByte(_bbga)
		}
	}
	if _ageg.Len()%2 == 1 {
		_aegd._aecec._gbcdc = true
		_ageg.WriteRune('0')
	}
	_fgec, _ := _cab.DecodeString(_ageg.String())
	return MakeHexString(string(_fgec)), nil
}

// DecodeStream decodes a multi-encoded stream by passing it through the
// DecodeStream method of the underlying encoders.
func (_add *MultiEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _add.DecodeBytes(streamObj.Stream)
}

// EncodeBytes DCT encodes the passed in slice of bytes.
func (_dac *DCTEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _bggd _ea.Image
	if _dac.ColorComponents == 1 && _dac.BitsPerComponent == 8 {
		_bggd = &_ea.Gray{Rect: _ea.Rect(0, 0, _dac.Width, _dac.Height), Pix: data, Stride: _ce.BytesPerLine(_dac.Width, _dac.BitsPerComponent, _dac.ColorComponents)}
	} else {
		var _bea error
		_bggd, _bea = _ce.NewImage(_dac.Width, _dac.Height, _dac.BitsPerComponent, _dac.ColorComponents, data, nil, nil)
		if _bea != nil {
			return nil, _bea
		}
	}
	_gabeb := _f.Options{}
	_gabeb.Quality = _dac.Quality
	var _cfb _fd.Buffer
	if _dceac := _f.Encode(&_cfb, _bggd, &_gabeb); _dceac != nil {
		return nil, _dceac
	}
	return _cfb.Bytes(), nil
}

// NewRawEncoder returns a new instace of RawEncoder.
func NewRawEncoder() *RawEncoder { return &RawEncoder{} }
func (_ceac *PdfParser) repairRebuildXrefsTopDown() (*XrefTable, error) {
	if _ceac._bgcf {
		return nil, _ac.Errorf("\u0072\u0065\u0070\u0061\u0069\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_ceac._bgcf = true
	_ceac._fdee.Seek(0, _dgf.SeekStart)
	_ceac._eecea = _acg.NewReader(_ceac._fdee)
	_egag := 20
	_eggbc := make([]byte, _egag)
	_cfef := XrefTable{}
	_cfef.ObjectMap = make(map[int]XrefObject)
	for {
		_fadb, _dbgc := _ceac._eecea.ReadByte()
		if _dbgc != nil {
			if _dbgc == _dgf.EOF {
				break
			} else {
				return nil, _dbgc
			}
		}
		if _fadb == 'j' && _eggbc[_egag-1] == 'b' && _eggbc[_egag-2] == 'o' && IsWhiteSpace(_eggbc[_egag-3]) {
			_dddga := _egag - 4
			for IsWhiteSpace(_eggbc[_dddga]) && _dddga > 0 {
				_dddga--
			}
			if _dddga == 0 || !IsDecimalDigit(_eggbc[_dddga]) {
				continue
			}
			for IsDecimalDigit(_eggbc[_dddga]) && _dddga > 0 {
				_dddga--
			}
			if _dddga == 0 || !IsWhiteSpace(_eggbc[_dddga]) {
				continue
			}
			for IsWhiteSpace(_eggbc[_dddga]) && _dddga > 0 {
				_dddga--
			}
			if _dddga == 0 || !IsDecimalDigit(_eggbc[_dddga]) {
				continue
			}
			for IsDecimalDigit(_eggbc[_dddga]) && _dddga > 0 {
				_dddga--
			}
			if _dddga == 0 {
				continue
			}
			_gadd := _ceac.GetFileOffset() - int64(_egag-_dddga)
			_agaa := append(_eggbc[_dddga+1:], _fadb)
			_fbcda, _gaafc, _ecgg := _bbdb(string(_agaa))
			if _ecgg != nil {
				_ae.Log.Debug("\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u003a\u0020\u0025\u0076", _ecgg)
				return nil, _ecgg
			}
			if _gdcca, _cgdc := _cfef.ObjectMap[_fbcda]; !_cgdc || _gdcca.Generation < _gaafc {
				_aaeg := XrefObject{}
				_aaeg.XType = XrefTypeTableEntry
				_aaeg.ObjectNumber = _fbcda
				_aaeg.Generation = _gaafc
				_aaeg.Offset = _gadd
				_cfef.ObjectMap[_fbcda] = _aaeg
			}
		}
		_eggbc = append(_eggbc[1:_egag], _fadb)
	}
	_ceac._ccge = nil
	return &_cfef, nil
}

// ToGoImage converts the JBIG2Image to the golang image.Image.
func (_ccc *JBIG2Image) ToGoImage() (_ea.Image, error) {
	const _gggg = "J\u0042I\u0047\u0032\u0049\u006d\u0061\u0067\u0065\u002eT\u006f\u0047\u006f\u0049ma\u0067\u0065"
	if _ccc.Data == nil {
		return nil, _eb.Error(_gggg, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _ccc.Width == 0 || _ccc.Height == 0 {
		return nil, _eb.Error(_gggg, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	_fffe, _fbba := _ce.NewImage(_ccc.Width, _ccc.Height, 1, 1, _ccc.Data, nil, nil)
	if _fbba != nil {
		return nil, _fbba
	}
	return _fffe, nil
}
func (_ddb *PdfCrypt) isEncrypted(_aegg PdfObject) bool {
	_, _fcff := _ddb._egg[_aegg]
	if _fcff {
		_ae.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0065\u006e\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	_ae.Log.Trace("\u004e\u006f\u0074\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}

// PdfVersion returns version of the PDF file.
func (_ebdd *PdfParser) PdfVersion() Version { return _ebdd._fadgc }

const (
	DefaultJPEGQuality = 75
)

// DecodeBytes decodes the CCITTFax encoded image data.
func (_cdde *CCITTFaxEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_cgdgg, _cadg := _ge.NewDecoder(encoded, _ge.DecodeOptions{Columns: _cdde.Columns, Rows: _cdde.Rows, K: _cdde.K, EncodedByteAligned: _cdde.EncodedByteAlign, BlackIsOne: _cdde.BlackIs1, EndOfBlock: _cdde.EndOfBlock, EndOfLine: _cdde.EndOfLine, DamagedRowsBeforeError: _cdde.DamagedRowsBeforeError})
	if _cadg != nil {
		return nil, _cadg
	}
	_efef, _cadg := _cg.ReadAll(_cgdgg)
	if _cadg != nil {
		return nil, _cadg
	}
	return _efef, nil
}

// ParserMetadata gets the pdf parser metadata.
func (_acdd *PdfParser) ParserMetadata() (ParserMetadata, error) {
	if !_acdd._ccce {
		return ParserMetadata{}, _ac.Errorf("\u0070\u0061\u0072\u0073\u0065r\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0072\u006be\u0064\u0020\u0066\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0064\u0065\u0074\u0061\u0069\u006c\u0065\u0064\u0020\u006d\u0065\u0074\u0061\u0064\u0061\u0074a")
	}
	return _acdd._aecec, nil
}

// DecodeStream decodes a JPX encoded stream and returns the result as a
// slice of bytes.
func (_cfea *JPXEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _cfea.GetFilterName())
	return streamObj.Stream, ErrNoJPXDecode
}

// Remove removes an element specified by key.
func (_aeff *PdfObjectDictionary) Remove(key PdfObjectName) {
	_ddcf := -1
	for _cdbd, _eaed := range _aeff._dgcd {
		if _eaed == key {
			_ddcf = _cdbd
			break
		}
	}
	if _ddcf >= 0 {
		_aeff._dgcd = append(_aeff._dgcd[:_ddcf], _aeff._dgcd[_ddcf+1:]...)
		delete(_aeff._gged, key)
	}
}

// GetString is a helper for Get that returns a string value.
// Returns false if the key is missing or a value is not a string.
func (_cfad *PdfObjectDictionary) GetString(key PdfObjectName) (string, bool) {
	_ddad := _cfad.Get(key)
	if _ddad == nil {
		return "", false
	}
	_aacc, _dbfde := _ddad.(*PdfObjectString)
	if !_dbfde {
		return "", false
	}
	return _aacc.Str(), true
}
func _cbbg(_ddcgg PdfObject, _febad int) PdfObject {
	if _febad > _dgag {
		_ae.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _dgag)
		return MakeNull()
	}
	switch _gcccb := _ddcgg.(type) {
	case *PdfIndirectObject:
		_ddcgg = _cbbg((*_gcccb).PdfObject, _febad+1)
	case *PdfObjectArray:
		for _edaf, _acacb := range (*_gcccb)._fffab {
			(*_gcccb)._fffab[_edaf] = _cbbg(_acacb, _febad+1)
		}
	case *PdfObjectDictionary:
		for _ccfce, _eddf := range (*_gcccb)._gged {
			(*_gcccb)._gged[_ccfce] = _cbbg(_eddf, _febad+1)
		}
		_ca.Slice((*_gcccb)._dgcd, func(_ggdee, _dcge int) bool { return (*_gcccb)._dgcd[_ggdee] < (*_gcccb)._dgcd[_dcge] })
	}
	return _ddcgg
}

// GetFloat returns the *PdfObjectFloat represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetFloat(obj PdfObject) (_bbba *PdfObjectFloat, _gaggg bool) {
	_bbba, _gaggg = TraceToDirectObject(obj).(*PdfObjectFloat)
	return _bbba, _gaggg
}

// PdfObjectBool represents the primitive PDF boolean object.
type PdfObjectBool bool

// IsHexadecimal checks if the PdfObjectString contains Hexadecimal data.
func (_gdab *PdfObjectString) IsHexadecimal() bool { return _gdab._fcge }

// IsDecimalDigit checks if the character is a part of a decimal number string.
func IsDecimalDigit(c byte) bool { return '0' <= c && c <= '9' }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_aage *MultiEncoder) MakeDecodeParams() PdfObject {
	if len(_aage._bfacg) == 0 {
		return nil
	}
	if len(_aage._bfacg) == 1 {
		return _aage._bfacg[0].MakeDecodeParams()
	}
	_baca := MakeArray()
	_ddfeg := true
	for _, _adcgf := range _aage._bfacg {
		_bfcd := _adcgf.MakeDecodeParams()
		if _bfcd == nil {
			_baca.Append(MakeNull())
		} else {
			_ddfeg = false
			_baca.Append(_bfcd)
		}
	}
	if _ddfeg {
		return nil
	}
	return _baca
}

// UpdateParams updates the parameter values of the encoder.
func (_ffcg *JPXEncoder) UpdateParams(params *PdfObjectDictionary) {}

// Len returns the number of elements in the streams.
func (_bcbbg *PdfObjectStreams) Len() int {
	if _bcbbg == nil {
		return 0
	}
	return len(_bcbbg._dgebc)
}

// MakeDictMap creates a PdfObjectDictionary initialized from a map of keys to values.
func MakeDictMap(objmap map[string]PdfObject) *PdfObjectDictionary {
	_gbfcc := MakeDict()
	return _gbfcc.Update(objmap)
}

// Append appends PdfObject(s) to the array.
func (_dgdbd *PdfObjectArray) Append(objects ...PdfObject) {
	if _dgdbd == nil {
		_ae.Log.Debug("\u0057\u0061\u0072\u006e\u0020\u002d\u0020\u0041\u0074\u0074\u0065\u006d\u0070t\u0020\u0074\u006f\u0020\u0061\u0070p\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020a\u0072\u0072\u0061\u0079")
		return
	}
	_dgdbd._fffab = append(_dgdbd._fffab, objects...)
}

// MakeObjectStreams creates an PdfObjectStreams from a list of PdfObjects.
func MakeObjectStreams(objects ...PdfObject) *PdfObjectStreams {
	return &PdfObjectStreams{_dgebc: objects}
}

// Elements returns a slice of the PdfObject elements in the array.
func (_cagg *PdfObjectArray) Elements() []PdfObject {
	if _cagg == nil {
		return nil
	}
	return _cagg._fffab
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_faaa *LZWEncoder) MakeStreamDict() *PdfObjectDictionary {
	_eggb := MakeDict()
	_eggb.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_faaa.GetFilterName()))
	_fda := _faaa.MakeDecodeParams()
	if _fda != nil {
		_eggb.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _fda)
	}
	_eggb.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(int64(_faaa.EarlyChange)))
	return _eggb
}

// GetAccessPermissions returns the PDF access permissions as an AccessPermissions object.
func (_bddb *PdfCrypt) GetAccessPermissions() _geg.Permissions { return _bddb._gge.P }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_bbd *ASCIIHexEncoder) MakeDecodeParams() PdfObject            { return nil }
func (_fbdgc *offsetReader) Read(p []byte) (_dedfc int, _ccac error) { return _fbdgc._agbf.Read(p) }

// DecodeStream implements ASCII85 stream decoding.
func (_faca *ASCII85Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _faca.DecodeBytes(streamObj.Stream)
}

// HasInvalidSubsectionHeader implements core.ParserMetadata interface.
func (_fbdg ParserMetadata) HasInvalidSubsectionHeader() bool { return _fbdg._bcg }

// GetFilterName returns the name of the encoding filter.
func (_cggg *JPXEncoder) GetFilterName() string { return StreamEncodingFilterNameJPX }

// MakeEncodedString creates a PdfObjectString with encoded content, which can be either
// UTF-16BE or PDFDocEncoding depending on whether `utf16BE` is true or false respectively.
func MakeEncodedString(s string, utf16BE bool) *PdfObjectString {
	if utf16BE {
		var _fbffe _fd.Buffer
		_fbffe.Write([]byte{0xFE, 0xFF})
		_fbffe.WriteString(_ef.StringToUTF16(s))
		return &PdfObjectString{_eeee: _fbffe.String(), _fcge: true}
	}
	return &PdfObjectString{_eeee: string(_ef.StringToPDFDocEncoding(s)), _fcge: false}
}

// UpdateParams updates the parameter values of the encoder.
func (_gcfe *ASCII85Encoder) UpdateParams(params *PdfObjectDictionary) {}

// MakeArrayFromFloats creates an PdfObjectArray from a slice of float64s, where each array element is an
// PdfObjectFloat.
func MakeArrayFromFloats(vals []float64) *PdfObjectArray {
	_cedb := MakeArray()
	for _, _ccfc := range vals {
		_cedb.Append(MakeFloat(_ccfc))
	}
	return _cedb
}

// String returns a string describing `array`.
func (_addgc *PdfObjectArray) String() string {
	_cafda := "\u005b"
	for _bebea, _bfed := range _addgc.Elements() {
		_cafda += _bfed.String()
		if _bebea < (_addgc.Len() - 1) {
			_cafda += "\u002c\u0020"
		}
	}
	_cafda += "\u005d"
	return _cafda
}

// DecodeBytes decodes a slice of DCT encoded bytes and returns the result.
func (_ebce *DCTEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_bgfg := _fd.NewReader(encoded)
	_feb, _bcf := _f.Decode(_bgfg)
	if _bcf != nil {
		_ae.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bcf)
		return nil, _bcf
	}
	_dccba := _feb.Bounds()
	var _dgc = make([]byte, _dccba.Dx()*_dccba.Dy()*_ebce.ColorComponents*_ebce.BitsPerComponent/8)
	_edac := 0
	for _degg := _dccba.Min.Y; _degg < _dccba.Max.Y; _degg++ {
		for _feeg := _dccba.Min.X; _feeg < _dccba.Max.X; _feeg++ {
			_acgdd := _feb.At(_feeg, _degg)
			if _ebce.ColorComponents == 1 {
				if _ebce.BitsPerComponent == 16 {
					_cdec, _geeb := _acgdd.(_ed.Gray16)
					if !_geeb {
						return nil, _d.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_dgc[_edac] = byte((_cdec.Y >> 8) & 0xff)
					_edac++
					_dgc[_edac] = byte(_cdec.Y & 0xff)
					_edac++
				} else {
					_cbcfa, _gcge := _acgdd.(_ed.Gray)
					if !_gcge {
						return nil, _d.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_dgc[_edac] = _cbcfa.Y & 0xff
					_edac++
				}
			} else if _ebce.ColorComponents == 3 {
				if _ebce.BitsPerComponent == 16 {
					_dbac, _ddgd := _acgdd.(_ed.RGBA64)
					if !_ddgd {
						return nil, _d.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_dgc[_edac] = byte((_dbac.R >> 8) & 0xff)
					_edac++
					_dgc[_edac] = byte(_dbac.R & 0xff)
					_edac++
					_dgc[_edac] = byte((_dbac.G >> 8) & 0xff)
					_edac++
					_dgc[_edac] = byte(_dbac.G & 0xff)
					_edac++
					_dgc[_edac] = byte((_dbac.B >> 8) & 0xff)
					_edac++
					_dgc[_edac] = byte(_dbac.B & 0xff)
					_edac++
				} else {
					_gac, _afdc := _acgdd.(_ed.RGBA)
					if _afdc {
						_dgc[_edac] = _gac.R & 0xff
						_edac++
						_dgc[_edac] = _gac.G & 0xff
						_edac++
						_dgc[_edac] = _gac.B & 0xff
						_edac++
					} else {
						_aacf, _gaabb := _acgdd.(_ed.YCbCr)
						if !_gaabb {
							return nil, _d.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
						}
						_ffc, _facf, _abd, _ := _aacf.RGBA()
						_dgc[_edac] = byte(_ffc >> 8)
						_edac++
						_dgc[_edac] = byte(_facf >> 8)
						_edac++
						_dgc[_edac] = byte(_abd >> 8)
						_edac++
					}
				}
			} else if _ebce.ColorComponents == 4 {
				_efcf, _daff := _acgdd.(_ed.CMYK)
				if !_daff {
					return nil, _d.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
				}
				_dgc[_edac] = 255 - _efcf.C&0xff
				_edac++
				_dgc[_edac] = 255 - _efcf.M&0xff
				_edac++
				_dgc[_edac] = 255 - _efcf.Y&0xff
				_edac++
				_dgc[_edac] = 255 - _efcf.K&0xff
				_edac++
			}
		}
	}
	return _dgc, nil
}
func (_dfece *PdfParser) parsePdfVersion() (int, int, error) {
	var _fagf int64 = 20
	_dbbc := make([]byte, _fagf)
	_dfece._fdee.Seek(0, _dgf.SeekStart)
	_dfece._fdee.Read(_dbbc)
	var _abdb error
	var _gegcg, _dbdf int
	if _gfgd := _fgadb.FindStringSubmatch(string(_dbbc)); len(_gfgd) < 3 {
		if _gegcg, _dbdf, _abdb = _dfece.seekPdfVersionTopDown(); _abdb != nil {
			_ae.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _abdb
		}
		_dfece._fdee, _abdb = _bfdd(_dfece._fdee, _dfece.GetFileOffset()-8)
		if _abdb != nil {
			return 0, 0, _abdb
		}
	} else {
		if _gegcg, _abdb = _dg.Atoi(_gfgd[1]); _abdb != nil {
			return 0, 0, _abdb
		}
		if _dbdf, _abdb = _dg.Atoi(_gfgd[2]); _abdb != nil {
			return 0, 0, _abdb
		}
		_dfece.SetFileOffset(0)
	}
	_dfece._eecea = _acg.NewReader(_dfece._fdee)
	_ae.Log.Debug("\u0050\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _gegcg, _dbdf)
	return _gegcg, _dbdf, nil
}

// GetFilterName returns the name of the encoding filter.
func (_cgef *RunLengthEncoder) GetFilterName() string { return StreamEncodingFilterNameRunLength }

// TraceToDirectObject traces a PdfObject to a direct object.  For example direct objects contained
// in indirect objects (can be double referenced even).
func TraceToDirectObject(obj PdfObject) PdfObject {
	if _dddfg, _gfgbd := obj.(*PdfObjectReference); _gfgbd {
		obj = _dddfg.Resolve()
	}
	_acbc, _ceaa := obj.(*PdfIndirectObject)
	_gggf := 0
	for _ceaa {
		obj = _acbc.PdfObject
		_acbc, _ceaa = GetIndirect(obj)
		_gggf++
		if _gggf > _dgag {
			_ae.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0061\u0063\u0065\u0020\u0064\u0065p\u0074\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u0062\u0065\u0079\u006fn\u0064\u0020\u0025\u0064\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0067oi\u006e\u0067\u0020\u0064\u0065\u0065\u0070\u0065\u0072\u0021", _dgag)
			return nil
		}
	}
	return obj
}

// GetObjectNums returns a sorted list of object numbers of the PDF objects in the file.
func (_eedf *PdfParser) GetObjectNums() []int {
	var _eefae []int
	for _, _fbae := range _eedf._fbab.ObjectMap {
		_eefae = append(_eefae, _fbae.ObjectNumber)
	}
	_ca.Ints(_eefae)
	return _eefae
}

// PdfObjectNull represents the primitive PDF null object.
type PdfObjectNull struct{}

// UpdateParams updates the parameter values of the encoder.
func (_fbeb *RunLengthEncoder) UpdateParams(params *PdfObjectDictionary) {}
func (_bgffd *PdfParser) parseBool() (PdfObjectBool, error) {
	_egca, _cfdd := _bgffd._eecea.Peek(4)
	if _cfdd != nil {
		return PdfObjectBool(false), _cfdd
	}
	if (len(_egca) >= 4) && (string(_egca[:4]) == "\u0074\u0072\u0075\u0065") {
		_bgffd._eecea.Discard(4)
		return PdfObjectBool(true), nil
	}
	_egca, _cfdd = _bgffd._eecea.Peek(5)
	if _cfdd != nil {
		return PdfObjectBool(false), _cfdd
	}
	if (len(_egca) >= 5) && (string(_egca[:5]) == "\u0066\u0061\u006cs\u0065") {
		_bgffd._eecea.Discard(5)
		return PdfObjectBool(false), nil
	}
	return PdfObjectBool(false), _d.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// GetBool returns the *PdfObjectBool object that is represented by a PdfObject directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetBool(obj PdfObject) (_cbebe *PdfObjectBool, _ggge bool) {
	_cbebe, _ggge = TraceToDirectObject(obj).(*PdfObjectBool)
	return _cbebe, _ggge
}

// FlateEncoder represents Flate encoding.
type FlateEncoder struct {
	Predictor        int
	BitsPerComponent int

	// For predictors
	Columns int
	Rows    int
	Colors  int
	_dga    *_ce.ImageBase
}

// Get returns the i-th element of the array or nil if out of bounds (by index).
func (_aeca *PdfObjectArray) Get(i int) PdfObject {
	if _aeca == nil || i >= len(_aeca._fffab) || i < 0 {
		return nil
	}
	return _aeca._fffab[i]
}

var (
	ErrUnsupportedEncodingParameters = _d.New("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	ErrNoCCITTFaxDecode              = _d.New("\u0043\u0043I\u0054\u0054\u0046\u0061\u0078\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJBIG2Decode                 = _d.New("\u004a\u0042\u0049\u0047\u0032\u0044\u0065c\u006f\u0064\u0065 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJPXDecode                   = _d.New("\u004a\u0050\u0058\u0044\u0065c\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoPdfVersion                  = _d.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	ErrTypeError                     = _d.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	ErrRangeError                    = _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrNotSupported                  = _bae.New("\u0066\u0065\u0061t\u0075\u0072\u0065\u0020n\u006f\u0074\u0020\u0063\u0075\u0072\u0072e\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
	ErrNotANumber                    = _d.New("\u006e\u006f\u0074 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
)

func (_ddff *PdfCrypt) encryptBytes(_fff []byte, _cdfd string, _cfce []byte) ([]byte, error) {
	_ae.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_eef, _dcec := _ddff._agfd[_cdfd]
	if !_dcec {
		return nil, _ac.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _cdfd)
	}
	return _eef.EncryptBytes(_fff, _cfce)
}

var _feffd = _ba.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064+\u0029\u005c\u0073\u002b\u0028\u005b\u006e\u0066\u005d\u0029\\\u0073\u002a\u0024")

// XrefObject defines a cross reference entry which is a map between object number (with generation number) and the
// location of the actual object, either as a file offset (xref table entry), or as a location within an xref
// stream object (xref object stream).
type XrefObject struct {
	XType        xrefType
	ObjectNumber int
	Generation   int

	// For normal xrefs (defined by OFFSET)
	Offset int64

	// For xrefs to object streams.
	OsObjNumber int
	OsObjIndex  int
}

// DecodeBytes decodes byte array with ASCII85. 5 ASCII characters -> 4 raw binary bytes
func (_cacg *ASCII85Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _beac []byte
	_ae.Log.Trace("\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0020\u0044e\u0063\u006f\u0064\u0065")
	_dcbf := 0
	_ddec := false
	for _dcbf < len(encoded) && !_ddec {
		_dacd := [5]byte{0, 0, 0, 0, 0}
		_bbfa := 0
		_dcbde := 0
		_ecg := 4
		for _dcbde < 5+_bbfa {
			if _dcbf+_dcbde == len(encoded) {
				break
			}
			_aacd := encoded[_dcbf+_dcbde]
			if IsWhiteSpace(_aacd) {
				_bbfa++
				_dcbde++
				continue
			} else if _aacd == '~' && _dcbf+_dcbde+1 < len(encoded) && encoded[_dcbf+_dcbde+1] == '>' {
				_ecg = (_dcbde - _bbfa) - 1
				if _ecg < 0 {
					_ecg = 0
				}
				_ddec = true
				break
			} else if _aacd >= '!' && _aacd <= 'u' {
				_aacd -= '!'
			} else if _aacd == 'z' && _dcbde-_bbfa == 0 {
				_ecg = 4
				_dcbde++
				break
			} else {
				_ae.Log.Error("\u0046\u0061i\u006c\u0065\u0064\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020co\u0064\u0065")
				return nil, _d.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u0064\u0065\u0020e\u006e\u0063\u006f\u0075\u006e\u0074\u0065\u0072\u0065\u0064")
			}
			_dacd[_dcbde-_bbfa] = _aacd
			_dcbde++
		}
		_dcbf += _dcbde
		for _abag := _ecg + 1; _abag < 5; _abag++ {
			_dacd[_abag] = 84
		}
		_aaa := uint32(_dacd[0])*85*85*85*85 + uint32(_dacd[1])*85*85*85 + uint32(_dacd[2])*85*85 + uint32(_dacd[3])*85 + uint32(_dacd[4])
		_agba := []byte{byte((_aaa >> 24) & 0xff), byte((_aaa >> 16) & 0xff), byte((_aaa >> 8) & 0xff), byte(_aaa & 0xff)}
		_beac = append(_beac, _agba[:_ecg]...)
	}
	_ae.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0065\u006e\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", encoded)
	_ae.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0064\u0065\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", _beac)
	return _beac, nil
}
func (_dadgd *PdfParser) repairLocateXref() (int64, error) {
	_bbfb := int64(1000)
	_dadgd._fdee.Seek(-_bbfb, _dgf.SeekCurrent)
	_gccgc, _bdfe := _dadgd._fdee.Seek(0, _dgf.SeekCurrent)
	if _bdfe != nil {
		return 0, _bdfe
	}
	_gada := make([]byte, _bbfb)
	_dadgd._fdee.Read(_gada)
	_ffdcg := _cgbea.FindAllStringIndex(string(_gada), -1)
	if len(_ffdcg) < 1 {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0065\u0070a\u0069\u0072\u003a\u0020\u0078\u0072\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021")
		return 0, _d.New("\u0072\u0065\u0070\u0061ir\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_fgaac := int64(_ffdcg[len(_ffdcg)-1][0])
	_ecefg := _gccgc + _fgaac
	return _ecefg, nil
}

// MakeString creates an PdfObjectString from a string.
// NOTE: PDF does not use utf-8 string encoding like Go so `s` will often not be a utf-8 encoded
// string.
func MakeString(s string) *PdfObjectString { _cggd := PdfObjectString{_eeee: s}; return &_cggd }

const (
	JB2Generic JBIG2CompressionType = iota
	JB2SymbolCorrelation
	JB2SymbolRankHaus
)

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_bgff *JPXEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }
func (_dggb *FlateEncoder) postDecodePredict(_adgc []byte) ([]byte, error) {
	if _dggb.Predictor > 1 {
		if _dggb.Predictor == 2 {
			_ae.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_ae.Log.Trace("\u0043\u006f\u006c\u006f\u0072\u0073\u003a\u0020\u0025\u0064", _dggb.Colors)
			_egbd := _dggb.Columns * _dggb.Colors
			if _egbd < 1 {
				return []byte{}, nil
			}
			_egae := len(_adgc) / _egbd
			if len(_adgc)%_egbd != 0 {
				_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_adgc), _egbd)
			}
			if _egbd%_dggb.Colors != 0 {
				return nil, _ac.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _egbd, _dggb.Colors)
			}
			if _egbd > len(_adgc) {
				_ae.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _egbd, len(_adgc))
				return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_ae.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_adgc), _adgc)
			_afda := _fd.NewBuffer(nil)
			for _gdc := 0; _gdc < _egae; _gdc++ {
				_aece := _adgc[_egbd*_gdc : _egbd*(_gdc+1)]
				for _gcc := _dggb.Colors; _gcc < _egbd; _gcc++ {
					_aece[_gcc] += _aece[_gcc-_dggb.Colors]
				}
				_afda.Write(_aece)
			}
			_faae := _afda.Bytes()
			_ae.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_faae), _faae)
			return _faae, nil
		} else if _dggb.Predictor >= 10 && _dggb.Predictor <= 15 {
			_ae.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_aae := _dggb.Columns*_dggb.Colors + 1
			_gegc := len(_adgc) / _aae
			if len(_adgc)%_aae != 0 {
				return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_adgc), _aae)
			}
			if _aae > len(_adgc) {
				_ae.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _aae, len(_adgc))
				return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_dae := _fd.NewBuffer(nil)
			_ae.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _dggb.Columns)
			_ae.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_adgc), _aae, _gegc)
			_eda := make([]byte, _aae)
			for _effc := 0; _effc < _aae; _effc++ {
				_eda[_effc] = 0
			}
			_afdb := _dggb.Colors
			for _fefe := 0; _fefe < _gegc; _fefe++ {
				_bbed := _adgc[_aae*_fefe : _aae*(_fefe+1)]
				_acdg := _bbed[0]
				switch _acdg {
				case _bgg:
				case _ebec:
					for _cff := 1 + _afdb; _cff < _aae; _cff++ {
						_bbed[_cff] += _bbed[_cff-_afdb]
					}
				case _ggf:
					for _aacb := 1; _aacb < _aae; _aacb++ {
						_bbed[_aacb] += _eda[_aacb]
					}
				case _ebca:
					for _cde := 1; _cde < _afdb+1; _cde++ {
						_bbed[_cde] += _eda[_cde] / 2
					}
					for _gdga := _afdb + 1; _gdga < _aae; _gdga++ {
						_bbed[_gdga] += byte((int(_bbed[_gdga-_afdb]) + int(_eda[_gdga])) / 2)
					}
				case _ffgg:
					for _egfg := 1; _egfg < _aae; _egfg++ {
						var _bgd, _daba, _gegd byte
						_daba = _eda[_egfg]
						if _egfg >= _afdb+1 {
							_bgd = _bbed[_egfg-_afdb]
							_gegd = _eda[_egfg-_afdb]
						}
						_bbed[_egfg] += _gdcg(_bgd, _daba, _gegd)
					}
				default:
					_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029\u0020\u0040\u0072o\u0077\u0020\u0025\u0064", _acdg, _fefe)
					return nil, _ac.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _acdg)
				}
				copy(_eda, _bbed)
				_dae.Write(_bbed[1:])
			}
			_cgbe := _dae.Bytes()
			return _cgbe, nil
		} else {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _dggb.Predictor)
			return nil, _ac.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _dggb.Predictor)
		}
	}
	return _adgc, nil
}

// UpdateParams updates the parameter values of the encoder.
func (_bge *ASCIIHexEncoder) UpdateParams(params *PdfObjectDictionary) {}

// GetNumberAsFloat returns the contents of `obj` as a float if it is an integer or float, or an
// error if it isn't.
func GetNumberAsFloat(obj PdfObject) (float64, error) {
	switch _fega := obj.(type) {
	case *PdfObjectFloat:
		return float64(*_fega), nil
	case *PdfObjectInteger:
		return float64(*_fega), nil
	}
	return 0, ErrNotANumber
}

// PdfObjectDictionary represents the primitive PDF dictionary/map object.
type PdfObjectDictionary struct {
	_gged  map[PdfObjectName]PdfObject
	_dgcd  []PdfObjectName
	_efce  *_e.Mutex
	_fdddg *PdfParser
}

// IsPrintable checks if a character is printable.
// Regular characters that are outside the range EXCLAMATION MARK(21h)
// (!) to TILDE (7Eh) (~) should be written using the hexadecimal notation.
func IsPrintable(c byte) bool { return 0x21 <= c && c <= 0x7E }

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_aebc *FlateEncoder) MakeStreamDict() *PdfObjectDictionary {
	_cbeg := MakeDict()
	_cbeg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_aebc.GetFilterName()))
	_dea := _aebc.MakeDecodeParams()
	if _dea != nil {
		_cbeg.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _dea)
	}
	return _cbeg
}
func _dfa(_cec *PdfObjectStream) (*MultiEncoder, error) {
	_daeg := NewMultiEncoder()
	_ecc := _cec.PdfObjectDictionary
	if _ecc == nil {
		return _daeg, nil
	}
	var _bdeac *PdfObjectDictionary
	var _dacb []PdfObject
	_eefg := _ecc.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
	if _eefg != nil {
		_dccg, _dgga := _eefg.(*PdfObjectDictionary)
		if _dgga {
			_bdeac = _dccg
		}
		_ggdb, _ddbg := _eefg.(*PdfObjectArray)
		if _ddbg {
			for _, _dgffe := range _ggdb.Elements() {
				_dgffe = TraceToDirectObject(_dgffe)
				if _fdad, _fce := _dgffe.(*PdfObjectDictionary); _fce {
					_dacb = append(_dacb, _fdad)
				} else {
					_dacb = append(_dacb, MakeDict())
				}
			}
		}
	}
	_eefg = _ecc.Get("\u0046\u0069\u006c\u0074\u0065\u0072")
	if _eefg == nil {
		return nil, _ac.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_gbe, _gcca := _eefg.(*PdfObjectArray)
	if !_gcca {
		return nil, _ac.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _bgga, _aed := range _gbe.Elements() {
		_caad, _fgcec := _aed.(*PdfObjectName)
		if !_fgcec {
			return nil, _ac.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _cdegc PdfObject
		if _bdeac != nil {
			_cdegc = _bdeac
		} else {
			if len(_dacb) > 0 {
				if _bgga >= len(_dacb) {
					return nil, _ac.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_cdegc = _dacb[_bgga]
			}
		}
		var _bdaa *PdfObjectDictionary
		if _gbge, _cfdg := _cdegc.(*PdfObjectDictionary); _cfdg {
			_bdaa = _gbge
		}
		_ae.Log.Trace("\u004e\u0065\u0078t \u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073\u002c \u0064p\u003a \u0025v\u002c\u0020\u0064\u0050\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u0076", *_caad, _cdegc, _bdaa)
		if *_caad == StreamEncodingFilterNameFlate {
			_fceb, _gfbf := _cbcf(_cec, _bdaa)
			if _gfbf != nil {
				return nil, _gfbf
			}
			_daeg.AddEncoder(_fceb)
		} else if *_caad == StreamEncodingFilterNameLZW {
			_egaf, _cdb := _cgbc(_cec, _bdaa)
			if _cdb != nil {
				return nil, _cdb
			}
			_daeg.AddEncoder(_egaf)
		} else if *_caad == StreamEncodingFilterNameASCIIHex {
			_afcg := NewASCIIHexEncoder()
			_daeg.AddEncoder(_afcg)
		} else if *_caad == StreamEncodingFilterNameASCII85 {
			_cdbf := NewASCII85Encoder()
			_daeg.AddEncoder(_cdbf)
		} else if *_caad == StreamEncodingFilterNameDCT {
			_ggde, _fgade := _faac(_cec, _daeg)
			if _fgade != nil {
				return nil, _fgade
			}
			_daeg.AddEncoder(_ggde)
			_ae.Log.Trace("A\u0064d\u0065\u0064\u0020\u0044\u0043\u0054\u0020\u0065n\u0063\u006f\u0064\u0065r.\u002e\u002e")
			_ae.Log.Trace("\u004du\u006ct\u0069\u0020\u0065\u006e\u0063o\u0064\u0065r\u003a\u0020\u0025\u0023\u0076", _daeg)
		} else if *_caad == StreamEncodingFilterNameCCITTFax {
			_cdge, _gccb := _acb(_cec, _bdaa)
			if _gccb != nil {
				return nil, _gccb
			}
			_daeg.AddEncoder(_cdge)
		} else {
			_ae.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_caad)
			return nil, _ac.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _daeg, nil
}

// EncodeBytes encodes data into ASCII85 encoded format.
func (_bdab *ASCII85Encoder) EncodeBytes(data []byte) ([]byte, error) {
	var _egadc _fd.Buffer
	for _cagd := 0; _cagd < len(data); _cagd += 4 {
		_gcbe := data[_cagd]
		_cfcf := 1
		_faag := byte(0)
		if _cagd+1 < len(data) {
			_faag = data[_cagd+1]
			_cfcf++
		}
		_effg := byte(0)
		if _cagd+2 < len(data) {
			_effg = data[_cagd+2]
			_cfcf++
		}
		_ceabf := byte(0)
		if _cagd+3 < len(data) {
			_ceabf = data[_cagd+3]
			_cfcf++
		}
		_aggc := (uint32(_gcbe) << 24) | (uint32(_faag) << 16) | (uint32(_effg) << 8) | uint32(_ceabf)
		if _aggc == 0 {
			_egadc.WriteByte('z')
		} else {
			_daaa := _bdab.base256Tobase85(_aggc)
			for _, _ceba := range _daaa[:_cfcf+1] {
				_egadc.WriteByte(_ceba + '!')
			}
		}
	}
	_egadc.WriteString("\u007e\u003e")
	return _egadc.Bytes(), nil
}

// GetFilterName returns the names of the underlying encoding filters,
// separated by spaces.
// Note: This is just a string, should not be used in /Filter dictionary entry. Use GetFilterArray for that.
// TODO(v4): Refactor to GetFilter() which can be used for /Filter (either Name or Array), this can be
//  renamed to String() as a pretty string to use in debugging etc.
func (_dace *MultiEncoder) GetFilterName() string {
	_dfe := ""
	for _eefe, _bgdb := range _dace._bfacg {
		_dfe += _bgdb.GetFilterName()
		if _eefe < len(_dace._bfacg)-1 {
			_dfe += "\u0020"
		}
	}
	return _dfe
}

// ASCII85Encoder implements ASCII85 encoder/decoder.
type ASCII85Encoder struct{}

// UpdateParams updates the parameter values of the encoder.
// Implements StreamEncoder interface.
func (_agff *JBIG2Encoder) UpdateParams(params *PdfObjectDictionary) {
	_bfdgc, _ecege := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _ecege == nil {
		_agff.BitsPerComponent = int(_bfdgc)
	}
	_aedg, _ecege := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _ecege == nil {
		_agff.Width = int(_aedg)
	}
	_accd, _ecege := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _ecege == nil {
		_agff.Height = int(_accd)
	}
	_eaec, _ecege := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _ecege == nil {
		_agff.ColorComponents = int(_eaec)
	}
}

// DecodeBytes decodes a slice of Flate encoded bytes and returns the result.
func (_fbe *FlateEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_ae.Log.Trace("\u0046\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065\u0020b\u0079\u0074\u0065\u0073")
	if len(encoded) == 0 {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0046\u006c\u0061\u0074\u0065 e\u006ec\u006f\u0064\u0065\u0064\u0020\u0062\u0075\u0066\u0066\u0065\u0072\u002e \u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0062y\u0074\u0065\u0020\u0073\u006c\u0069\u0063\u0065\u002e")
		return []byte{}, nil
	}
	_eece := _fd.NewReader(encoded)
	_gfcc, _ggc := _bc.NewReader(_eece)
	if _ggc != nil {
		_ae.Log.Debug("\u0044e\u0063o\u0064\u0069\u006e\u0067\u0020e\u0072\u0072o\u0072\u0020\u0025\u0076\u000a", _ggc)
		_ae.Log.Debug("\u0053t\u0072e\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(encoded), encoded)
		return nil, _ggc
	}
	defer _gfcc.Close()
	var _cdcf _fd.Buffer
	_cdcf.ReadFrom(_gfcc)
	return _cdcf.Bytes(), nil
}
func (_fdeea *PdfParser) seekPdfVersionTopDown() (int, int, error) {
	_fdeea._fdee.Seek(0, _dgf.SeekStart)
	_fdeea._eecea = _acg.NewReader(_fdeea._fdee)
	_bdcf := 20
	_dedga := make([]byte, _bdcf)
	for {
		_fcebe, _adac := _fdeea._eecea.ReadByte()
		if _adac != nil {
			if _adac == _dgf.EOF {
				break
			} else {
				return 0, 0, _adac
			}
		}
		if IsDecimalDigit(_fcebe) && _dedga[_bdcf-1] == '.' && IsDecimalDigit(_dedga[_bdcf-2]) && _dedga[_bdcf-3] == '-' && _dedga[_bdcf-4] == 'F' && _dedga[_bdcf-5] == 'D' && _dedga[_bdcf-6] == 'P' {
			_bedee := int(_dedga[_bdcf-2] - '0')
			_ffcea := int(_fcebe - '0')
			return _bedee, _ffcea, nil
		}
		_dedga = append(_dedga[1:_bdcf], _fcebe)
	}
	return 0, 0, _d.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// String returns a string representation of the *PdfObjectString.
func (_ggdef *PdfObjectString) String() string { return _ggdef._eeee }

// WriteString outputs the object as it is to be written to file.
func (_fbffc *PdfObjectArray) WriteString() string {
	var _bace _agg.Builder
	_bace.WriteString("\u005b")
	for _ebedf, _ddfef := range _fbffc.Elements() {
		_bace.WriteString(_ddfef.WriteString())
		if _ebedf < (_fbffc.Len() - 1) {
			_bace.WriteString("\u0020")
		}
	}
	_bace.WriteString("\u005d")
	return _bace.String()
}

// WriteString outputs the object as it is to be written to file.
func (_efdbf *PdfObjectFloat) WriteString() string {
	return _dg.FormatFloat(float64(*_efdbf), 'f', -1, 64)
}

// String returns a string describing `stream`.
func (_gadf *PdfObjectStream) String() string {
	return _ac.Sprintf("O\u0062j\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0025\u0064: \u0025\u0073", _gadf.ObjectNumber, _gadf.PdfObjectDictionary)
}

// PdfObjectName represents the primitive PDF name object.
type PdfObjectName string

// Merge merges in key/values from another dictionary. Overwriting if has same keys.
// The mutated dictionary (d) is returned in order to allow method chaining.
func (_decfa *PdfObjectDictionary) Merge(another *PdfObjectDictionary) *PdfObjectDictionary {
	if another != nil {
		for _, _cdege := range another.Keys() {
			_gebbg := another.Get(_cdege)
			_decfa.Set(_cdege, _gebbg)
		}
	}
	return _decfa
}

// MakeStreamDict make a new instance of an encoding dictionary for a stream object.
func (_caee *ASCII85Encoder) MakeStreamDict() *PdfObjectDictionary {
	_becc := MakeDict()
	_becc.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_caee.GetFilterName()))
	return _becc
}

type objectCache map[int]PdfObject

// GetFilterName returns the name of the encoding filter.
func (_agbd *CCITTFaxEncoder) GetFilterName() string { return StreamEncodingFilterNameCCITTFax }

// DecodeGlobals decodes 'encoded' byte stream and returns their Globally defined segments ('Globals').
func (_efaa *JBIG2Encoder) DecodeGlobals(encoded []byte) (_df.Globals, error) {
	return _df.DecodeGlobals(encoded)
}
func (_cdgcbd *PdfParser) resolveReference(_fada *PdfObjectReference) (PdfObject, bool, error) {
	_decf, _cffb := _cdgcbd.ObjCache[int(_fada.ObjectNumber)]
	if _cffb {
		return _decf, true, nil
	}
	_gcfg, _eeg := _cdgcbd.LookupByReference(*_fada)
	if _eeg != nil {
		return nil, false, _eeg
	}
	_cdgcbd.ObjCache[int(_fada.ObjectNumber)] = _gcfg
	return _gcfg, false, nil
}

const _egb = "\u0053\u0074\u0064C\u0046"

// Clear resets the array to an empty state.
func (_geebc *PdfObjectArray) Clear() { _geebc._fffab = []PdfObject{} }
func (_gdbf *offsetReader) Seek(offset int64, whence int) (int64, error) {
	if whence == _dgf.SeekStart {
		offset += _gdbf._ceca
	}
	_ddcd, _bddbc := _gdbf._agbf.Seek(offset, whence)
	if _bddbc != nil {
		return _ddcd, _bddbc
	}
	if whence == _dgf.SeekCurrent {
		_ddcd -= _gdbf._ceca
	}
	if _ddcd < 0 {
		return 0, _d.New("\u0063\u006f\u0072\u0065\u002eo\u0066\u0066\u0073\u0065\u0074\u0052\u0065\u0061\u0064\u0065\u0072\u002e\u0053e\u0065\u006b\u003a\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	return _ddcd, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_gcbf *LZWEncoder) MakeDecodeParams() PdfObject {
	if _gcbf.Predictor > 1 {
		_ecbd := MakeDict()
		_ecbd.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_gcbf.Predictor)))
		if _gcbf.BitsPerComponent != 8 {
			_ecbd.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_gcbf.BitsPerComponent)))
		}
		if _gcbf.Columns != 1 {
			_ecbd.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_gcbf.Columns)))
		}
		if _gcbf.Colors != 1 {
			_ecbd.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_gcbf.Colors)))
		}
		return _ecbd
	}
	return nil
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_eefd *RunLengthEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_gbcg := _fd.NewReader(data)
	var _dgac []byte
	var _cagf []byte
	_fdba, _efg := _gbcg.ReadByte()
	if _efg == _dgf.EOF {
		return []byte{}, nil
	} else if _efg != nil {
		return nil, _efg
	}
	_dgad := 1
	for {
		_bdea, _gceb := _gbcg.ReadByte()
		if _gceb == _dgf.EOF {
			break
		} else if _gceb != nil {
			return nil, _gceb
		}
		if _bdea == _fdba {
			if len(_cagf) > 0 {
				_cagf = _cagf[:len(_cagf)-1]
				if len(_cagf) > 0 {
					_dgac = append(_dgac, byte(len(_cagf)-1))
					_dgac = append(_dgac, _cagf...)
				}
				_dgad = 1
				_cagf = []byte{}
			}
			_dgad++
			if _dgad >= 127 {
				_dgac = append(_dgac, byte(257-_dgad), _fdba)
				_dgad = 0
			}
		} else {
			if _dgad > 0 {
				if _dgad == 1 {
					_cagf = []byte{_fdba}
				} else {
					_dgac = append(_dgac, byte(257-_dgad), _fdba)
				}
				_dgad = 0
			}
			_cagf = append(_cagf, _bdea)
			if len(_cagf) >= 127 {
				_dgac = append(_dgac, byte(len(_cagf)-1))
				_dgac = append(_dgac, _cagf...)
				_cagf = []byte{}
			}
		}
		_fdba = _bdea
	}
	if len(_cagf) > 0 {
		_dgac = append(_dgac, byte(len(_cagf)-1))
		_dgac = append(_dgac, _cagf...)
	} else if _dgad > 0 {
		_dgac = append(_dgac, byte(257-_dgad), _fdba)
	}
	_dgac = append(_dgac, 128)
	return _dgac, nil
}
func (_afefb *PdfParser) parseObject() (PdfObject, error) {
	_ae.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_afefb.skipSpaces()
	for {
		_gfga, _ecad := _afefb._eecea.Peek(2)
		if _ecad != nil {
			if _ecad != _dgf.EOF || len(_gfga) == 0 {
				return nil, _ecad
			}
			if len(_gfga) == 1 {
				_gfga = append(_gfga, ' ')
			}
		}
		_ae.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_gfga))
		if _gfga[0] == '/' {
			_fecf, _faaee := _afefb.parseName()
			_ae.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _fecf)
			return &_fecf, _faaee
		} else if _gfga[0] == '(' {
			_ae.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_fafdf, _febd := _afefb.parseString()
			return _fafdf, _febd
		} else if _gfga[0] == '[' {
			_ae.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_eefb, _fbc := _afefb.parseArray()
			return _eefb, _fbc
		} else if (_gfga[0] == '<') && (_gfga[1] == '<') {
			_ae.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			_dafa, _abcd := _afefb.ParseDict()
			return _dafa, _abcd
		} else if _gfga[0] == '<' {
			_ae.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			_adeb, _aeeg := _afefb.parseHexString()
			return _adeb, _aeeg
		} else if _gfga[0] == '%' {
			_afefb.readComment()
			_afefb.skipSpaces()
		} else {
			_ae.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_gfga, _ = _afefb._eecea.Peek(15)
			_cedg := string(_gfga)
			_ae.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _cedg)
			if (len(_cedg) > 3) && (_cedg[:4] == "\u006e\u0075\u006c\u006c") {
				_feef, _gaeb := _afefb.parseNull()
				return &_feef, _gaeb
			} else if (len(_cedg) > 4) && (_cedg[:5] == "\u0066\u0061\u006cs\u0065") {
				_eecd, _gece := _afefb.parseBool()
				return &_eecd, _gece
			} else if (len(_cedg) > 3) && (_cedg[:4] == "\u0074\u0072\u0075\u0065") {
				_deff, _fgae := _afefb.parseBool()
				return &_deff, _fgae
			}
			_fggc := _aea.FindStringSubmatch(_cedg)
			if len(_fggc) > 1 {
				_gfga, _ = _afefb._eecea.ReadBytes('R')
				_ae.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_gfga[:]))
				_edde, _bfff := _acbd(string(_gfga))
				_edde._ffgd = _afefb
				return &_edde, _bfff
			}
			_fcbf := _eecg.FindStringSubmatch(_cedg)
			if len(_fcbf) > 1 {
				_ae.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				_bdbe, _acfc := _afefb.parseNumber()
				return _bdbe, _acfc
			}
			_fcbf = _abce.FindStringSubmatch(_cedg)
			if len(_fcbf) > 1 {
				_ae.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_ae.Log.Trace("\u0025\u0020\u0073", _fcbf)
				_cgbf, _badd := _afefb.parseNumber()
				return _cgbf, _badd
			}
			_ae.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _cedg)
			return nil, _d.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}

// MakeStream creates an PdfObjectStream with specified contents and encoding. If encoding is nil, then raw encoding
// will be used (i.e. no encoding applied).
func MakeStream(contents []byte, encoder StreamEncoder) (*PdfObjectStream, error) {
	_acae := &PdfObjectStream{}
	if encoder == nil {
		encoder = NewRawEncoder()
	}
	_acae.PdfObjectDictionary = encoder.MakeStreamDict()
	_cfdga, _addge := encoder.EncodeBytes(contents)
	if _addge != nil {
		return nil, _addge
	}
	_acae.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_cfdga))))
	_acae.Stream = _cfdga
	return _acae, nil
}

var _ecee = _d.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")

func (_ffea *PdfParser) parseArray() (*PdfObjectArray, error) {
	_afcd := MakeArray()
	_ffea._eecea.ReadByte()
	for {
		_ffea.skipSpaces()
		_ecdge, _gaff := _ffea._eecea.Peek(1)
		if _gaff != nil {
			return _afcd, _gaff
		}
		if _ecdge[0] == ']' {
			_ffea._eecea.ReadByte()
			break
		}
		_ggbfg, _gaff := _ffea.parseObject()
		if _gaff != nil {
			return _afcd, _gaff
		}
		_afcd.Append(_ggbfg)
	}
	return _afcd, nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_aabdf *JBIG2Encoder) MakeStreamDict() *PdfObjectDictionary {
	_bdge := MakeDict()
	_bdge.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_aabdf.GetFilterName()))
	return _bdge
}
func (_gfc *PdfCrypt) checkAccessRights(_gbd []byte) (bool, _geg.Permissions, error) {
	_fggd := _gfc.securityHandler()
	_cdd, _fec, _fdgb := _fggd.Authenticate(&_gfc._gge, _gbd)
	if _fdgb != nil {
		return false, 0, _fdgb
	} else if _fec == 0 || len(_cdd) == 0 {
		return false, 0, nil
	}
	return true, _fec, nil
}

// GetFileOffset returns the current file offset, accounting for buffered position.
func (_bcgc *PdfParser) GetFileOffset() int64 {
	_cdcg, _ := _bcgc._fdee.Seek(0, _dgf.SeekCurrent)
	_cdcg -= int64(_bcgc._eecea.Buffered())
	return _cdcg
}

// CCITTFaxEncoder implements Group3 and Group4 facsimile (fax) encoder/decoder.
type CCITTFaxEncoder struct {
	K                      int
	EndOfLine              bool
	EncodedByteAlign       bool
	Columns                int
	Rows                   int
	EndOfBlock             bool
	BlackIs1               bool
	DamagedRowsBeforeError int
}

func _gba(_bdd *_geg.StdEncryptDict, _dgda *PdfObjectDictionary) error {
	R, _eac := _dgda.Get("\u0052").(*PdfObjectInteger)
	if !_eac {
		return _d.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0052")
	}
	if *R < 2 || *R > 6 {
		return _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052 \u0028\u0025\u0064\u0029", *R)
	}
	_bdd.R = int(*R)
	O, _eac := _dgda.GetString("\u004f")
	if !_eac {
		return _d.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u004f")
	}
	if _bdd.R == 5 || _bdd.R == 6 {
		if len(O) < 48 {
			return _ac.Errorf("\u004c\u0065\u006e\u0067th\u0028\u004f\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(O))
		}
	} else if len(O) != 32 {
		return _ac.Errorf("L\u0065n\u0067\u0074\u0068\u0028\u004f\u0029\u0020\u0021=\u0020\u0033\u0032\u0020(%\u0064\u0029", len(O))
	}
	_bdd.O = []byte(O)
	U, _eac := _dgda.GetString("\u0055")
	if !_eac {
		return _d.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0055")
	}
	if _bdd.R == 5 || _bdd.R == 6 {
		if len(U) < 48 {
			return _ac.Errorf("\u004c\u0065\u006e\u0067th\u0028\u0055\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(U))
		}
	} else if len(U) != 32 {
		_ae.Log.Debug("\u0057\u0061r\u006e\u0069\u006e\u0067\u003a\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0028\u0055\u0029\u0020\u0021\u003d\u0020\u0033\u0032\u0020(%\u0064\u0029", len(U))
	}
	_bdd.U = []byte(U)
	if _bdd.R >= 5 {
		OE, _cga := _dgda.GetString("\u004f\u0045")
		if !_cga {
			return _d.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u004f\u0045")
		} else if len(OE) != 32 {
			return _ac.Errorf("L\u0065\u006e\u0067\u0074h(\u004fE\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(OE))
		}
		_bdd.OE = []byte(OE)
		UE, _cga := _dgda.GetString("\u0055\u0045")
		if !_cga {
			return _d.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u0055\u0045")
		} else if len(UE) != 32 {
			return _ac.Errorf("L\u0065\u006e\u0067\u0074h(\u0055E\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(UE))
		}
		_bdd.UE = []byte(UE)
	}
	P, _eac := _dgda.Get("\u0050").(*PdfObjectInteger)
	if !_eac {
		return _d.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074 \u0064\u0069\u0063t\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0061\u0074\u0074\u0072")
	}
	_bdd.P = _geg.Permissions(*P)
	if _bdd.R == 6 {
		Perms, _fef := _dgda.GetString("\u0050\u0065\u0072m\u0073")
		if !_fef {
			return _d.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0050\u0065\u0072\u006d\u0073")
		} else if len(Perms) != 16 {
			return _ac.Errorf("\u004ce\u006e\u0067\u0074\u0068\u0028\u0050\u0065\u0072\u006d\u0073\u0029 \u0021\u003d\u0020\u0031\u0036\u0020\u0028\u0025\u0064\u0029", len(Perms))
		}
		_bdd.Perms = []byte(Perms)
	}
	if _caa, _ddf := _dgda.Get("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061").(*PdfObjectBool); _ddf {
		_bdd.EncryptMetadata = bool(*_caa)
	} else {
		_bdd.EncryptMetadata = true
	}
	return nil
}

// Set sets the dictionary's key -> val mapping entry. Overwrites if key already set.
func (_aada *PdfObjectDictionary) Set(key PdfObjectName, val PdfObject) {
	_aada.setWithLock(key, val, true)
}

// GetObjectStreams returns the *PdfObjectStreams represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetObjectStreams(obj PdfObject) (_cagfe *PdfObjectStreams, _bbag bool) {
	_cagfe, _bbag = obj.(*PdfObjectStreams)
	return _cagfe, _bbag
}
func (_gbefb *PdfParser) rebuildXrefTable() error {
	_dgfg := XrefTable{}
	_dgfg.ObjectMap = map[int]XrefObject{}
	_beec := make([]int, 0, len(_gbefb._fbab.ObjectMap))
	for _facg := range _gbefb._fbab.ObjectMap {
		_beec = append(_beec, _facg)
	}
	_ca.Ints(_beec)
	for _, _egadd := range _beec {
		_fcfa := _gbefb._fbab.ObjectMap[_egadd]
		_ffad, _, _dfbg := _gbefb.lookupByNumberWrapper(_egadd, false)
		if _dfbg != nil {
			_ae.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f l\u006f\u006f\u006b\u0020\u0075\u0070\u0020ob\u006a\u0065\u0063\u0074\u0020\u0028\u0025s\u0029", _dfbg)
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u006fm\u0070\u006c\u0065\u0074\u0065\u006c\u0079\u0020\u0062\u0072\u006f\u006b\u0065\u006e\u0020\u002d\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0074\u006f \u0072\u0065\u0070\u0061\u0069r\u0020")
			_gbbd, _dff := _gbefb.repairRebuildXrefsTopDown()
			if _dff != nil {
				_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0072\u0065\u0062\u0075\u0069l\u0064\u0020\u0072\u0065\u0070a\u0069\u0072 \u0028\u0025\u0073\u0029", _dff)
				return _dff
			}
			_gbefb._fbab = *_gbbd
			_ae.Log.Debug("\u0052e\u0070\u0061\u0069\u0072e\u0064\u0020\u0078\u0072\u0065f\u0020t\u0061b\u006c\u0065\u0020\u0062\u0075\u0069\u006ct")
			return nil
		}
		_bddgd, _baaf, _dfbg := _def(_ffad)
		if _dfbg != nil {
			return _dfbg
		}
		_fcfa.ObjectNumber = int(_bddgd)
		_fcfa.Generation = int(_baaf)
		_dgfg.ObjectMap[int(_bddgd)] = _fcfa
	}
	_gbefb._fbab = _dgfg
	_ae.Log.Debug("N\u0065w\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061b\u006c\u0065\u0020\u0062ui\u006c\u0074")
	_eec(_gbefb._fbab)
	return nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_acddg *DCTEncoder) MakeDecodeParams() PdfObject { return nil }

// RunLengthEncoder represents Run length encoding.
type RunLengthEncoder struct{}

// DecodeBytes decodes a byte slice from Run length encoding.
//
// 7.4.5 RunLengthDecode Filter
// The RunLengthDecode filter decodes data that has been encoded in a simple byte-oriented format based on run length.
// The encoded data shall be a sequence of runs, where each run shall consist of a length byte followed by 1 to 128
// bytes of data. If the length byte is in the range 0 to 127, the following length + 1 (1 to 128) bytes shall be
// copied literally during decompression. If length is in the range 129 to 255, the following single byte shall be
// copied 257 - length (2 to 128) times during decompression. A length value of 128 shall denote EOD.
func (_gebd *RunLengthEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_abbe := _fd.NewReader(encoded)
	var _deag []byte
	for {
		_fggb, _dfd := _abbe.ReadByte()
		if _dfd != nil {
			return nil, _dfd
		}
		if _fggb > 128 {
			_ffcf, _acfe := _abbe.ReadByte()
			if _acfe != nil {
				return nil, _acfe
			}
			for _abdc := 0; _abdc < 257-int(_fggb); _abdc++ {
				_deag = append(_deag, _ffcf)
			}
		} else if _fggb < 128 {
			for _beeg := 0; _beeg < int(_fggb)+1; _beeg++ {
				_fafdd, _dggf := _abbe.ReadByte()
				if _dggf != nil {
					return nil, _dggf
				}
				_deag = append(_deag, _fafdd)
			}
		} else {
			break
		}
	}
	return _deag, nil
}

// ToInt64Slice returns a slice of all array elements as an int64 slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_ffde *PdfObjectArray) ToInt64Slice() ([]int64, error) {
	var _gdcf []int64
	for _, _fgaa := range _ffde.Elements() {
		if _cfeb, _efcd := _fgaa.(*PdfObjectInteger); _efcd {
			_gdcf = append(_gdcf, int64(*_cfeb))
		} else {
			return nil, ErrTypeError
		}
	}
	return _gdcf, nil
}

// DecodeBytes decodes a multi-encoded slice of bytes by passing it through the
// DecodeBytes method of the underlying encoders.
func (_dcaac *MultiEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_bagc := encoded
	var _ffdf error
	for _, _geec := range _dcaac._bfacg {
		_ae.Log.Trace("\u004du\u006c\u0074i\u0020\u0045\u006e\u0063o\u0064\u0065\u0072 \u0044\u0065\u0063\u006f\u0064\u0065\u003a\u0020\u0041pp\u006c\u0079\u0069n\u0067\u0020F\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0076 \u0025\u0054", _geec, _geec)
		_bagc, _ffdf = _geec.DecodeBytes(_bagc)
		if _ffdf != nil {
			return nil, _ffdf
		}
	}
	return _bagc, nil
}

const (
	StreamEncodingFilterNameFlate     = "F\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameLZW       = "\u004cZ\u0057\u0044\u0065\u0063\u006f\u0064e"
	StreamEncodingFilterNameDCT       = "\u0044C\u0054\u0044\u0065\u0063\u006f\u0064e"
	StreamEncodingFilterNameRunLength = "\u0052u\u006eL\u0065\u006e\u0067\u0074\u0068\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameASCIIHex  = "\u0041\u0053\u0043\u0049\u0049\u0048\u0065\u0078\u0044e\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameASCII85   = "\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameCCITTFax  = "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameJBIG2     = "J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065"
	StreamEncodingFilterNameJPX       = "\u004aP\u0058\u0044\u0065\u0063\u006f\u0064e"
	StreamEncodingFilterNameRaw       = "\u0052\u0061\u0077"
)

// Append appends PdfObject(s) to the streams.
func (_fcgd *PdfObjectStreams) Append(objects ...PdfObject) {
	if _fcgd == nil {
		_ae.Log.Debug("\u0057\u0061\u0072\u006e\u0020-\u0020\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0074\u006f\u0020\u0061p\u0070\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073")
		return
	}
	_fcgd._dgebc = append(_fcgd._dgebc, objects...)
}

// Decrypt attempts to decrypt the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful, false otherwise.
// An error is returned when there is a problem with decrypting.
func (_ecdc *PdfParser) Decrypt(password []byte) (bool, error) {
	if _ecdc._bffd == nil {
		return false, _d.New("\u0063\u0068\u0065\u0063k \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u0072s\u0074")
	}
	_baae, _bdeg := _ecdc._bffd.authenticate(password)
	if _bdeg != nil {
		return false, _bdeg
	}
	if !_baae {
		_baae, _bdeg = _ecdc._bffd.authenticate([]byte(""))
	}
	return _baae, _bdeg
}
func (_ceda *PdfParser) inspect() (map[string]int, error) {
	_ae.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0053P\u0045\u0043\u0054\u0020\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d")
	_ae.Log.Trace("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_fbbff := map[string]int{}
	_dfag := 0
	_ffcc := 0
	var _abeg []int
	for _bagb := range _ceda._fbab.ObjectMap {
		_abeg = append(_abeg, _bagb)
	}
	_ca.Ints(_abeg)
	_fggag := 0
	for _, _cdded := range _abeg {
		_beaa := _ceda._fbab.ObjectMap[_cdded]
		if _beaa.ObjectNumber == 0 {
			continue
		}
		_dfag++
		_ae.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
		_ae.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0075\u0070\u0020\u006f\u0062\u006a\u0065c\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0025\u0064", _beaa.ObjectNumber)
		_eadf, _eaga := _ceda.LookupByNumber(_beaa.ObjectNumber)
		if _eaga != nil {
			_ae.Log.Trace("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u006c\u006f\u006f\u006b\u0075p\u0020\u006f\u0062\u006a\u0020\u0025\u0064 \u0028\u0025\u0073\u0029", _beaa.ObjectNumber, _eaga)
			_ffcc++
			continue
		}
		_ae.Log.Trace("\u006fb\u006a\u003a\u0020\u0025\u0073", _eadf)
		_geee, _acbe := _eadf.(*PdfIndirectObject)
		if _acbe {
			_ae.Log.Trace("\u0049N\u0044 \u004f\u004f\u0042\u004a\u0020\u0025\u0064\u003a\u0020\u0025\u0073", _beaa.ObjectNumber, _geee)
			_gbdg, _gaac := _geee.PdfObject.(*PdfObjectDictionary)
			if _gaac {
				if _efcb, _fegf := _gbdg.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _fegf {
					_agcaa := string(*_efcb)
					_ae.Log.Trace("\u002d\u002d\u002d\u003e\u0020\u004f\u0062\u006a\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _agcaa)
					_, _gaabf := _fbbff[_agcaa]
					if _gaabf {
						_fbbff[_agcaa]++
					} else {
						_fbbff[_agcaa] = 1
					}
				} else if _addfcb, _cfafa := _gbdg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065").(*PdfObjectName); _cfafa {
					_dgdge := string(*_addfcb)
					_ae.Log.Trace("-\u002d-\u003e\u0020\u004f\u0062\u006a\u0020\u0073\u0075b\u0074\u0079\u0070\u0065: \u0025\u0073", _dgdge)
					_, _acfa := _fbbff[_dgdge]
					if _acfa {
						_fbbff[_dgdge]++
					} else {
						_fbbff[_dgdge] = 1
					}
				}
				if _geeed, _dfgf := _gbdg.Get("\u0053").(*PdfObjectName); _dfgf && *_geeed == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
					_, _eccb := _fbbff["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]
					if _eccb {
						_fbbff["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
					} else {
						_fbbff["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"] = 1
					}
				}
			}
		} else if _acee, _dgee := _eadf.(*PdfObjectStream); _dgee {
			if _ffda, _ebba := _acee.PdfObjectDictionary.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _ebba {
				_ae.Log.Trace("\u002d\u002d\u003e\u0020\u0053\u0074\u0072\u0065\u0061\u006d\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073", *_ffda)
				_faef := string(*_ffda)
				_fbbff[_faef]++
			}
		} else {
			_fcfb, _baea := _eadf.(*PdfObjectDictionary)
			if _baea {
				_aaabe, _bggf := _fcfb.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
				if _bggf {
					_dddfgf := string(*_aaabe)
					_ae.Log.Trace("\u002d-\u002d \u006f\u0062\u006a\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _dddfgf)
					_fbbff[_dddfgf]++
				}
			}
			_ae.Log.Trace("\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u004f\u0042\u004a\u0020\u0025d\u003a\u0020\u0025\u0073", _beaa.ObjectNumber, _eadf)
		}
		_fggag++
	}
	_ae.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0045\u004fF\u0020\u0049\u004e\u0053\u0050\u0045\u0043T\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d")
	_ae.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	_ae.Log.Trace("\u004f\u0062j\u0065\u0063\u0074 \u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064", _dfag)
	_ae.Log.Trace("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u006f\u006b\u0075p\u003a\u0020\u0025\u0064", _ffcc)
	for _ccae, _befb := range _fbbff {
		_ae.Log.Trace("\u0025\u0073\u003a\u0020\u0025\u0064", _ccae, _befb)
	}
	_ae.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	if len(_ceda._fbab.ObjectMap) < 1 {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0068\u0069\u0073 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074 \u0069s\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0021\u0029")
		return nil, _ac.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0028\u0078r\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u0029")
	}
	_afcf, _baffe := _fbbff["\u0046\u006f\u006e\u0074"]
	if !_baffe || _afcf < 2 {
		_ae.Log.Trace("\u0054\u0068\u0069s \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073 \u0070r\u006fb\u0061\u0062\u006c\u0079\u0020\u0073\u0063\u0061\u006e\u006e\u0065\u0064\u0021")
	} else {
		_ae.Log.Trace("\u0054\u0068\u0069\u0073\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0066o\u0072\u0020\u0065\u0078\u0074r\u0061\u0063t\u0069\u006f\u006e\u0021")
	}
	return _fbbff, nil
}

// GetFilterName returns the name of the encoding filter.
func (_ffdg *DCTEncoder) GetFilterName() string { return StreamEncodingFilterNameDCT }

// IsOctalDigit checks if a character can be part of an octal digit string.
func IsOctalDigit(c byte) bool { return '0' <= c && c <= '7' }

// WriteString outputs the object as it is to be written to file.
func (_dcagg *PdfObjectNull) WriteString() string { return "\u006e\u0075\u006c\u006c" }
func (_agca *PdfCrypt) newEncryptDict() *PdfObjectDictionary {
	_fdg := MakeDict()
	_fdg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064"))
	_fdg.Set("\u0056", MakeInteger(int64(_agca._ddc.V)))
	_fdg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_agca._ddc.Length)))
	return _fdg
}

// EncodeBytes encodes the image data using either Group3 or Group4 CCITT facsimile (fax) encoding.
// `data` is expected to be 1 color component, 1 bit per component. It is also valid to provide 8 BPC, 1 CC image like
// a standard go image Gray data.
func (_bgbb *CCITTFaxEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _ffgge _ce.Gray
	switch len(data) {
	case _bgbb.Rows * _bgbb.Columns:
		_bebe, _aaaa := _ce.NewImage(_bgbb.Columns, _bgbb.Rows, 8, 1, data, nil, nil)
		if _aaaa != nil {
			return nil, _aaaa
		}
		_ffgge = _bebe.(_ce.Gray)
	case (_bgbb.Columns * _bgbb.Rows) + 7>>3:
		_cbdc, _fcac := _ce.NewImage(_bgbb.Columns, _bgbb.Rows, 1, 1, data, nil, nil)
		if _fcac != nil {
			return nil, _fcac
		}
		_bfacd := _cbdc.(*_ce.Monochrome)
		if _fcac = _bfacd.AddPadding(); _fcac != nil {
			return nil, _fcac
		}
		_ffgge = _bfacd
	default:
		if len(data) < _ce.BytesPerLine(_bgbb.Columns, 1, 1)*_bgbb.Rows {
			return nil, _d.New("p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020i\u006e\u0070\u0075t\u0020d\u0061\u0074\u0061")
		}
		_eaab, _cda := _ce.NewImage(_bgbb.Columns, _bgbb.Rows, 1, 1, data, nil, nil)
		if _cda != nil {
			return nil, _cda
		}
		_adcg := _eaab.(*_ce.Monochrome)
		_ffgge = _adcg
	}
	_acff := make([][]byte, _bgbb.Rows)
	for _ebfbg := 0; _ebfbg < _bgbb.Rows; _ebfbg++ {
		_daae := make([]byte, _bgbb.Columns)
		for _abc := 0; _abc < _bgbb.Columns; _abc++ {
			_ffgb := _ffgge.GrayAt(_abc, _ebfbg)
			_daae[_abc] = _ffgb.Y >> 7
		}
		_acff[_ebfbg] = _daae
	}
	_ffbf := &_ge.Encoder{K: _bgbb.K, Columns: _bgbb.Columns, EndOfLine: _bgbb.EndOfLine, EndOfBlock: _bgbb.EndOfBlock, BlackIs1: _bgbb.BlackIs1, DamagedRowsBeforeError: _bgbb.DamagedRowsBeforeError, Rows: _bgbb.Rows, EncodedByteAlign: _bgbb.EncodedByteAlign}
	return _ffbf.Encode(_acff), nil
}
func (_afdf *PdfParser) seekToEOFMarker(_gcdf int64) error {
	var _dgfd int64
	var _adaa int64 = 2048
	for _dgfd < _gcdf-4 {
		if _gcdf <= (_adaa + _dgfd) {
			_adaa = _gcdf - _dgfd
		}
		_, _daadb := _afdf._fdee.Seek(_gcdf-_dgfd-_adaa, _dgf.SeekStart)
		if _daadb != nil {
			return _daadb
		}
		_cgea := make([]byte, _adaa)
		_afdf._fdee.Read(_cgea)
		_ae.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_cgea))
		_dddfb := _cbae.FindAllStringIndex(string(_cgea), -1)
		if _dddfb != nil {
			_dgdcb := _dddfb[len(_dddfb)-1]
			_ae.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _dddfb)
			_ffeae := _gcdf - _dgfd - _adaa + int64(_dgdcb[0])
			_afdf._fdee.Seek(_ffeae, _dgf.SeekStart)
			return nil
		}
		_ae.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_dgfd += _adaa - 4
	}
	_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _ecee
}

// MakeInteger creates a PdfObjectInteger from an int64.
func MakeInteger(val int64) *PdfObjectInteger { _cbacf := PdfObjectInteger(val); return &_cbacf }

// Resolve resolves a PdfObject to direct object, looking up and resolving references as needed (unlike TraceToDirect).
func (_ccf *PdfParser) Resolve(obj PdfObject) (PdfObject, error) {
	_bec, _fa := obj.(*PdfObjectReference)
	if !_fa {
		return obj, nil
	}
	_cgc := _ccf.GetFileOffset()
	defer func() { _ccf.SetFileOffset(_cgc) }()
	_fde, _dgeg := _ccf.LookupByReference(*_bec)
	if _dgeg != nil {
		return nil, _dgeg
	}
	_beg, _effb := _fde.(*PdfIndirectObject)
	if !_effb {
		return _fde, nil
	}
	_fde = _beg.PdfObject
	_, _fa = _fde.(*PdfObjectReference)
	if _fa {
		return _beg, _d.New("\u006d\u0075lt\u0069\u0020\u0064e\u0070\u0074\u0068\u0020tra\u0063e \u0070\u006f\u0069\u006e\u0074\u0065\u0072 t\u006f\u0020\u0070\u006f\u0069\u006e\u0074e\u0072")
	}
	return _fde, nil
}

// GetName returns the *PdfObjectName represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetName(obj PdfObject) (_fdeec *PdfObjectName, _fdaa bool) {
	_fdeec, _fdaa = TraceToDirectObject(obj).(*PdfObjectName)
	return _fdeec, _fdaa
}
func _ecca(_gfgg _ce.Image) *JBIG2Image {
	_fddf := _gfgg.Base()
	return &JBIG2Image{Data: _fddf.Data, Width: _fddf.Width, Height: _fddf.Height, HasPadding: true}
}
func (_aec *PdfParser) checkPostEOFData() error {
	const _dfgef = "\u0025\u0025\u0045O\u0046"
	_, _gdfd := _aec._fdee.Seek(-int64(len([]byte(_dfgef)))-1, _dgf.SeekEnd)
	if _gdfd != nil {
		return _gdfd
	}
	_dedf := make([]byte, len([]byte(_dfgef))+1)
	_, _gdfd = _aec._fdee.Read(_dedf)
	if _gdfd != nil {
		if _gdfd != _dgf.EOF {
			return _gdfd
		}
	}
	if string(_dedf) == _dfgef || string(_dedf) == _dfgef+"\u000a" {
		_aec._aecec._fdgg = true
	}
	return nil
}
func _afe(_bef *_ebd.FilterDict, _aee *PdfObjectDictionary) error {
	if _agc, _fefg := _aee.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _fefg {
		if _faa := string(*_agc); _faa != "C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072" {
			_ae.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020C\u0046\u0020\u0064ic\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074e\u0072\u0029", _faa)
		}
	}
	_gfa, _fca := _aee.Get("\u0043\u0046\u004d").(*PdfObjectName)
	if !_fca {
		return _ac.Errorf("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0028\u004e\u006f\u006e\u0065\u0029")
	}
	_bef.CFM = string(*_gfa)
	if _fed, _faf := _aee.Get("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et").(*PdfObjectName); _faf {
		_bef.AuthEvent = _geg.AuthEvent(*_fed)
	} else {
		_bef.AuthEvent = _geg.EventDocOpen
	}
	if _adff, _gec := _aee.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _gec {
		_bef.Length = int(*_adff)
	}
	return nil
}

// NewASCII85Encoder makes a new ASCII85 encoder.
func NewASCII85Encoder() *ASCII85Encoder { _efbc := &ASCII85Encoder{}; return _efbc }

// NewRunLengthEncoder makes a new run length encoder
func NewRunLengthEncoder() *RunLengthEncoder { return &RunLengthEncoder{} }

// Inspect analyzes the document object structure. Returns a map of object types (by name) with the instance count
// as value.
func (_edgc *PdfParser) Inspect() (map[string]int, error) { return _edgc.inspect() }

// PdfObjectInteger represents the primitive PDF integer numerical object.
type PdfObjectInteger int64

// GetIndirect returns the *PdfIndirectObject represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetIndirect(obj PdfObject) (_efge *PdfIndirectObject, _cbaed bool) {
	obj = ResolveReference(obj)
	_efge, _cbaed = obj.(*PdfIndirectObject)
	return _efge, _cbaed
}
func _acb(_bdcb *PdfObjectStream, _badf *PdfObjectDictionary) (*CCITTFaxEncoder, error) {
	_bedc := NewCCITTFaxEncoder()
	_dadgb := _bdcb.PdfObjectDictionary
	if _dadgb == nil {
		return _bedc, nil
	}
	if _badf == nil {
		_dded := TraceToDirectObject(_dadgb.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _dded != nil {
			switch _ceff := _dded.(type) {
			case *PdfObjectDictionary:
				_badf = _ceff
			case *PdfObjectArray:
				if _ceff.Len() == 1 {
					if _cadd, _eaefa := GetDict(_ceff.Get(0)); _eaefa {
						_badf = _cadd
					}
				}
			default:
				_ae.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _dded)
				return nil, _d.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
		if _badf == nil {
			_ae.Log.Error("\u0044\u0065c\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064 %\u0023\u0076", _dded)
			return nil, _d.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _ede, _gfef := GetNumberAsInt64(_badf.Get("\u004b")); _gfef == nil {
		_bedc.K = int(_ede)
	}
	if _gfb, _geea := GetNumberAsInt64(_badf.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _geea == nil {
		_bedc.Columns = int(_gfb)
	} else {
		_bedc.Columns = 1728
	}
	if _fdddc, _abef := GetNumberAsInt64(_badf.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _abef == nil {
		_bedc.BlackIs1 = _fdddc > 0
	} else {
		if _bfae, _cabaa := GetBoolVal(_badf.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _cabaa {
			_bedc.BlackIs1 = _bfae
		} else {
			if _caec, _aecg := GetArray(_badf.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _aecg {
				_aafd, _ffed := _caec.ToIntegerArray()
				if _ffed == nil {
					_bedc.BlackIs1 = _aafd[0] == 1 && _aafd[1] == 0
				}
			}
		}
	}
	if _bede, _dfbc := GetNumberAsInt64(_badf.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _dfbc == nil {
		_bedc.EncodedByteAlign = _bede > 0
	} else {
		if _eefdd, _eggaa := GetBoolVal(_badf.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _eggaa {
			_bedc.EncodedByteAlign = _eefdd
		}
	}
	if _aegf, _edfdd := GetNumberAsInt64(_badf.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _edfdd == nil {
		_bedc.EndOfLine = _aegf > 0
	} else {
		if _dgdc, _eaee := GetBoolVal(_badf.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _eaee {
			_bedc.EndOfLine = _dgdc
		}
	}
	if _fgcd, _ccg := GetNumberAsInt64(_badf.Get("\u0052\u006f\u0077\u0073")); _ccg == nil {
		_bedc.Rows = int(_fgcd)
	}
	_bedc.EndOfBlock = true
	if _cdfce, _gaae := GetNumberAsInt64(_badf.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _gaae == nil {
		_bedc.EndOfBlock = _cdfce > 0
	} else {
		if _ccgg, _ddfg := GetBoolVal(_badf.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _ddfg {
			_bedc.EndOfBlock = _ccgg
		}
	}
	if _abfd, _bgca := GetNumberAsInt64(_badf.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _bgca != nil {
		_bedc.DamagedRowsBeforeError = int(_abfd)
	}
	_ae.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _badf.String())
	return _bedc, nil
}

// LookupByReference looks up a PdfObject by a reference.
func (_bfc *PdfParser) LookupByReference(ref PdfObjectReference) (PdfObject, error) {
	_ae.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0075\u0070\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0025\u0073", ref.String())
	return _bfc.LookupByNumber(int(ref.ObjectNumber))
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_gfe *RunLengthEncoder) MakeDecodeParams() PdfObject { return nil }

const _dgag = 10

// String returns a string describing `streams`.
func (_cbab *PdfObjectStreams) String() string {
	return _ac.Sprintf("\u004f\u0062j\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0025\u0064", _cbab.ObjectNumber)
}
func _gfca(_dbb *PdfObjectStream, _bee *PdfObjectDictionary) (*RunLengthEncoder, error) {
	return NewRunLengthEncoder(), nil
}

// NewParser creates a new parser for a PDF file via ReadSeeker. Loads the cross reference stream and trailer.
// An error is returned on failure.
func NewParser(rs _dgf.ReadSeeker) (*PdfParser, error) {
	_fgcf := &PdfParser{_fdee: rs, ObjCache: make(objectCache), _gfdf: map[int64]bool{}, _edcac: make([]int64, 0), _bgcdd: make(map[*PdfParser]*PdfParser)}
	_cgba, _eggg, _cbga := _fgcf.parsePdfVersion()
	if _cbga != nil {
		_ae.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _cbga)
		return nil, _cbga
	}
	_fgcf._fadgc.Major = _cgba
	_fgcf._fadgc.Minor = _eggg
	if _fgcf._dfec, _cbga = _fgcf.loadXrefs(); _cbga != nil {
		_ae.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _cbga)
		return nil, _cbga
	}
	_ae.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _fgcf._dfec)
	_bgeg, _cbga := _fgcf.parseLinearizedDictionary()
	if _cbga != nil {
		return nil, _cbga
	}
	if _bgeg != nil {
		_fgcf._gfggc, _cbga = _fgcf.checkLinearizedInformation(_bgeg)
		if _cbga != nil {
			return nil, _cbga
		}
	}
	if len(_fgcf._fbab.ObjectMap) == 0 {
		return nil, _ac.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	_fgcf._fdbdc = len(_fgcf._edcac)
	if _fgcf._gfggc && _fgcf._fdbdc != 0 {
		_fgcf._fdbdc--
	}
	_fgcf._ddcg = make([]*PdfParser, _fgcf._fdbdc)
	return _fgcf, nil
}
func (_efdb *PdfParser) parseName() (PdfObjectName, error) {
	var _aedf _fd.Buffer
	_gagff := false
	for {
		_egdg, _egdc := _efdb._eecea.Peek(1)
		if _egdc == _dgf.EOF {
			break
		}
		if _egdc != nil {
			return PdfObjectName(_aedf.String()), _egdc
		}
		if !_gagff {
			if _egdg[0] == '/' {
				_gagff = true
				_efdb._eecea.ReadByte()
			} else if _egdg[0] == '%' {
				_efdb.readComment()
				_efdb.skipSpaces()
			} else {
				_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _egdg, _egdg)
				return PdfObjectName(_aedf.String()), _ac.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _egdg[0])
			}
		} else {
			if IsWhiteSpace(_egdg[0]) {
				break
			} else if (_egdg[0] == '/') || (_egdg[0] == '[') || (_egdg[0] == '(') || (_egdg[0] == ']') || (_egdg[0] == '<') || (_egdg[0] == '>') {
				break
			} else if _egdg[0] == '#' {
				_fcef, _aaabd := _efdb._eecea.Peek(3)
				if _aaabd != nil {
					return PdfObjectName(_aedf.String()), _aaabd
				}
				_decd, _aaabd := _cab.DecodeString(string(_fcef[1:3]))
				if _aaabd != nil {
					_ae.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0066o\u006c\u006co\u0077\u0069\u006e\u0067 \u0027\u0023\u0027\u002c \u0063\u006f\u006e\u0074\u0069n\u0075\u0069\u006e\u0067\u0020\u0075\u0073i\u006e\u0067\u0020\u006c\u0069t\u0065\u0072\u0061\u006c\u0020\u002d\u0020\u004f\u0075t\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074")
					_aedf.WriteByte('#')
					_efdb._eecea.Discard(1)
					continue
				}
				_efdb._eecea.Discard(3)
				_aedf.Write(_decd)
			} else {
				_agea, _ := _efdb._eecea.ReadByte()
				_aedf.WriteByte(_agea)
			}
		}
	}
	return PdfObjectName(_aedf.String()), nil
}
func _faac(_cabge *PdfObjectStream, _egee *MultiEncoder) (*DCTEncoder, error) {
	_caef := NewDCTEncoder()
	_bcgfa := _cabge.PdfObjectDictionary
	if _bcgfa == nil {
		return _caef, nil
	}
	_daea := _cabge.Stream
	if _egee != nil {
		_gbbc, _bcbe := _egee.DecodeBytes(_daea)
		if _bcbe != nil {
			return nil, _bcbe
		}
		_daea = _gbbc
	}
	_ebbe := _fd.NewReader(_daea)
	_gga, _aba := _f.DecodeConfig(_ebbe)
	if _aba != nil {
		_ae.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _aba)
		return nil, _aba
	}
	switch _gga.ColorModel {
	case _ed.RGBAModel:
		_caef.BitsPerComponent = 8
		_caef.ColorComponents = 3
	case _ed.RGBA64Model:
		_caef.BitsPerComponent = 16
		_caef.ColorComponents = 3
	case _ed.GrayModel:
		_caef.BitsPerComponent = 8
		_caef.ColorComponents = 1
	case _ed.Gray16Model:
		_caef.BitsPerComponent = 16
		_caef.ColorComponents = 1
	case _ed.CMYKModel:
		_caef.BitsPerComponent = 8
		_caef.ColorComponents = 4
	case _ed.YCbCrModel:
		_caef.BitsPerComponent = 8
		_caef.ColorComponents = 3
	default:
		return nil, _d.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_caef.Width = _gga.Width
	_caef.Height = _gga.Height
	_ae.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _caef)
	_caef.Quality = DefaultJPEGQuality
	return _caef, nil
}

const _fba = 6

func _adef(_fccc *PdfObjectStream, _acda *PdfObjectDictionary) (*JBIG2Encoder, error) {
	const _eaad = "\u006ee\u0077\u004a\u0042\u0049G\u0032\u0044\u0065\u0063\u006fd\u0065r\u0046r\u006f\u006d\u0053\u0074\u0072\u0065\u0061m"
	_egbb := NewJBIG2Encoder()
	_fbebb := _fccc.PdfObjectDictionary
	if _fbebb == nil {
		return _egbb, nil
	}
	if _acda == nil {
		_bcde := _fbebb.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		if _bcde != nil {
			switch _cgbd := _bcde.(type) {
			case *PdfObjectDictionary:
				_acda = _cgbd
			case *PdfObjectArray:
				if _cgbd.Len() == 1 {
					if _eecce, _edfa := GetDict(_cgbd.Get(0)); _edfa {
						_acda = _eecce
					}
				}
			default:
				_ae.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065P\u0061\u0072\u0061\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0025\u0023\u0076", _bcde)
				return nil, _eb.Errorf(_eaad, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050a\u0072m\u0073\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _cgbd)
			}
		}
	}
	if _acda == nil {
		return _egbb, nil
	}
	_egbb.UpdateParams(_acda)
	_cgfb, _edb := GetStream(_acda.Get("\u004a\u0042\u0049G\u0032\u0047\u006c\u006f\u0062\u0061\u006c\u0073"))
	if !_edb {
		return _egbb, nil
	}
	var _caca error
	_egbb.Globals, _caca = _df.DecodeGlobals(_cgfb.Stream)
	if _caca != nil {
		_caca = _eb.Wrap(_caca, _eaad, "\u0063\u006f\u0072\u0072u\u0070\u0074\u0065\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0065n\u0063\u006f\u0064\u0065\u0064\u0020\u0064a\u0074\u0061")
		_ae.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _caca)
		return nil, _caca
	}
	return _egbb, nil
}

// SetPredictor sets the predictor function.  Specify the number of columns per row.
// The columns indicates the number of samples per row.
// Used for grouping data together for compression.
func (_daf *FlateEncoder) SetPredictor(columns int) { _daf.Predictor = 11; _daf.Columns = columns }

// RawEncoder implements Raw encoder/decoder (no encoding, pass through)
type RawEncoder struct{}

// String returns a string describing `d`.
func (_cbdd *PdfObjectDictionary) String() string {
	var _agfa _agg.Builder
	_agfa.WriteString("\u0044\u0069\u0063t\u0028")
	for _, _cebb := range _cbdd._dgcd {
		_aafcg := _cbdd._gged[_cebb]
		_agfa.WriteString("\u0022" + _cebb.String() + "\u0022\u003a\u0020")
		_agfa.WriteString(_aafcg.String())
		_agfa.WriteString("\u002c\u0020")
	}
	_agfa.WriteString("\u0029")
	return _agfa.String()
}

// Decrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_abeb *PdfCrypt) Decrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _abeb.isDecrypted(obj) {
		return nil
	}
	switch _cae := obj.(type) {
	case *PdfIndirectObject:
		_abeb._dgd[_cae] = true
		_ae.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _cae.ObjectNumber, _cae.GenerationNumber)
		_ccaa := _cae.ObjectNumber
		_cgcd := _cae.GenerationNumber
		_fead := _abeb.Decrypt(_cae.PdfObject, _ccaa, _cgcd)
		if _fead != nil {
			return _fead
		}
		return nil
	case *PdfObjectStream:
		_abeb._dgd[_cae] = true
		_gcae := _cae.PdfObjectDictionary
		if _abeb._gge.R != 5 {
			if _dfbb, _bad := _gcae.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _bad && *_dfbb == "\u0058\u0052\u0065\u0066" {
				return nil
			}
		}
		_adfff := _cae.ObjectNumber
		_agdd := _cae.GenerationNumber
		_ae.Log.Trace("\u0044e\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _adfff, _agdd)
		_eadb := _egb
		if _abeb._ddc.V >= 4 {
			_eadb = _abeb._bba
			_ae.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _abeb._bba)
			if _adb, _bca := _gcae.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _bca {
				if _ffb, _bde := GetName(_adb.Get(0)); _bde {
					if *_ffb == "\u0043\u0072\u0079p\u0074" {
						_eadb = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _egc, _baee := _gcae.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _baee {
							if _egd, _fdb := _egc.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _fdb {
								if _, _fgfd := _abeb._agfd[string(*_egd)]; _fgfd {
									_ae.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_egd)
									_eadb = string(*_egd)
								}
							}
						}
					}
				}
			}
			_ae.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _eadb)
			if _eadb == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_aegc := _abeb.Decrypt(_gcae, _adfff, _agdd)
		if _aegc != nil {
			return _aegc
		}
		_caae, _aegc := _abeb.makeKey(_eadb, uint32(_adfff), uint32(_agdd), _abeb._dadd)
		if _aegc != nil {
			return _aegc
		}
		_cae.Stream, _aegc = _abeb.decryptBytes(_cae.Stream, _eadb, _caae)
		if _aegc != nil {
			return _aegc
		}
		_gcae.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_cae.Stream))))
		return nil
	case *PdfObjectString:
		_ae.Log.Trace("\u0044e\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_cag := _egb
		if _abeb._ddc.V >= 4 {
			_ae.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _abeb._dee)
			if _abeb._dee == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_cag = _abeb._dee
		}
		_gbb, _fdd := _abeb.makeKey(_cag, uint32(parentObjNum), uint32(parentGenNum), _abeb._dadd)
		if _fdd != nil {
			return _fdd
		}
		_fbd := _cae.Str()
		_eag := make([]byte, len(_fbd))
		for _bfb := 0; _bfb < len(_fbd); _bfb++ {
			_eag[_bfb] = _fbd[_bfb]
		}
		if len(_eag) > 0 {
			_ae.Log.Trace("\u0044e\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _eag, _eag)
			_eag, _fdd = _abeb.decryptBytes(_eag, _cag, _gbb)
			if _fdd != nil {
				return _fdd
			}
		}
		_cae._eeee = string(_eag)
		return nil
	case *PdfObjectArray:
		for _, _geaf := range _cae.Elements() {
			_dgg := _abeb.Decrypt(_geaf, parentObjNum, parentGenNum)
			if _dgg != nil {
				return _dgg
			}
		}
		return nil
	case *PdfObjectDictionary:
		_fgd := false
		if _dcc := _cae.Get("\u0054\u0079\u0070\u0065"); _dcc != nil {
			_ecde, _abb := _dcc.(*PdfObjectName)
			if _abb && *_ecde == "\u0053\u0069\u0067" {
				_fgd = true
			}
		}
		for _, _gdf := range _cae.Keys() {
			_gdd := _cae.Get(_gdf)
			if _fgd && string(_gdf) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_gdf) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_gdf) != "\u0050\u0072\u0065\u0076" && string(_gdf) != "\u004c\u0061\u0073\u0074" {
				_ace := _abeb.Decrypt(_gdd, parentObjNum, parentGenNum)
				if _ace != nil {
					return _ace
				}
			}
		}
		return nil
	}
	return nil
}

const (
	XrefTypeTableEntry   xrefType = iota
	XrefTypeObjectStream xrefType = iota
)

// NewCompliancePdfParser creates a new PdfParser that will parse input reader with the focus on extracting more metadata, which
// might affect performance of the regular PdfParser this function.
func NewCompliancePdfParser(rs _dgf.ReadSeeker) (_dcbd *PdfParser, _cgac error) {
	_dcbd = &PdfParser{_fdee: rs, ObjCache: make(objectCache), _gfdf: map[int64]bool{}, _ccce: true, _bgcdd: make(map[*PdfParser]*PdfParser)}
	if _cgac = _dcbd.parseDetailedHeader(); _cgac != nil {
		return nil, _cgac
	}
	if _dcbd._dfec, _cgac = _dcbd.loadXrefs(); _cgac != nil {
		_ae.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _cgac)
		return nil, _cgac
	}
	_ae.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _dcbd._dfec)
	if len(_dcbd._fbab.ObjectMap) == 0 {
		return nil, _ac.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	return _dcbd, nil
}
func _cdcd(_bcgca int) int {
	if _bcgca < 0 {
		return -_bcgca
	}
	return _bcgca
}

// String returns a string describing `ind`.
func (_gbagb *PdfIndirectObject) String() string {
	return _ac.Sprintf("\u0049\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0025\u0064", (*_gbagb).ObjectNumber)
}

// Seek implementation of Seek interface.
func (_gfdc *limitedReadSeeker) Seek(offset int64, whence int) (int64, error) {
	var _dccf int64
	switch whence {
	case _dgf.SeekStart:
		_dccf = offset
	case _dgf.SeekCurrent:
		_fbdf, _fegcc := _gfdc._cgfa.Seek(0, _dgf.SeekCurrent)
		if _fegcc != nil {
			return 0, _fegcc
		}
		_dccf = _fbdf + offset
	case _dgf.SeekEnd:
		_dccf = _gfdc._gfcg + offset
	}
	if _ecgb := _gfdc.getError(_dccf); _ecgb != nil {
		return 0, _ecgb
	}
	if _, _ada := _gfdc._cgfa.Seek(_dccf, _dgf.SeekStart); _ada != nil {
		return 0, _ada
	}
	return _dccf, nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set.  Some other parameters are generated elsewhere.
func (_dcbb *DCTEncoder) MakeStreamDict() *PdfObjectDictionary {
	_gde := MakeDict()
	_gde.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_dcbb.GetFilterName()))
	return _gde
}

// Get returns the PdfObject corresponding to the specified key.
// Returns a nil value if the key is not set.
func (_ffdff *PdfObjectDictionary) Get(key PdfObjectName) PdfObject {
	_ffdff._efce.Lock()
	defer _ffdff._efce.Unlock()
	_ccfee, _aceg := _ffdff._gged[key]
	if !_aceg {
		return nil
	}
	return _ccfee
}

// DecodeStream returns the passed in stream as a slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_feec *RawEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return streamObj.Stream, nil
}
func (_dgae *PdfParser) parseString() (*PdfObjectString, error) {
	_dgae._eecea.ReadByte()
	var _degge _fd.Buffer
	_gaaff := 1
	for {
		_ebgd, _eddg := _dgae._eecea.Peek(1)
		if _eddg != nil {
			return MakeString(_degge.String()), _eddg
		}
		if _ebgd[0] == '\\' {
			_dgae._eecea.ReadByte()
			_edeg, _bfaca := _dgae._eecea.ReadByte()
			if _bfaca != nil {
				return MakeString(_degge.String()), _bfaca
			}
			if IsOctalDigit(_edeg) {
				_ccgb, _adbg := _dgae._eecea.Peek(2)
				if _adbg != nil {
					return MakeString(_degge.String()), _adbg
				}
				var _adbga []byte
				_adbga = append(_adbga, _edeg)
				for _, _cfeab := range _ccgb {
					if IsOctalDigit(_cfeab) {
						_adbga = append(_adbga, _cfeab)
					} else {
						break
					}
				}
				_dgae._eecea.Discard(len(_adbga) - 1)
				_ae.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _adbga)
				_bdgb, _adbg := _dg.ParseUint(string(_adbga), 8, 32)
				if _adbg != nil {
					return MakeString(_degge.String()), _adbg
				}
				_degge.WriteByte(byte(_bdgb))
				continue
			}
			switch _edeg {
			case 'n':
				_degge.WriteRune('\n')
			case 'r':
				_degge.WriteRune('\r')
			case 't':
				_degge.WriteRune('\t')
			case 'b':
				_degge.WriteRune('\b')
			case 'f':
				_degge.WriteRune('\f')
			case '(':
				_degge.WriteRune('(')
			case ')':
				_degge.WriteRune(')')
			case '\\':
				_degge.WriteRune('\\')
			}
			continue
		} else if _ebgd[0] == '(' {
			_gaaff++
		} else if _ebgd[0] == ')' {
			_gaaff--
			if _gaaff == 0 {
				_dgae._eecea.ReadByte()
				break
			}
		}
		_ggfd, _ := _dgae._eecea.ReadByte()
		_degge.WriteByte(_ggfd)
	}
	return MakeString(_degge.String()), nil
}

// UpdateParams updates the parameter values of the encoder.
func (_bga *DCTEncoder) UpdateParams(params *PdfObjectDictionary) {
	_abgb, _cfdb := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _cfdb == nil {
		_bga.ColorComponents = int(_abgb)
	}
	_bfdf, _cfdb := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _cfdb == nil {
		_bga.BitsPerComponent = int(_bfdf)
	}
	_gae, _cfdb := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _cfdb == nil {
		_bga.Width = int(_gae)
	}
	_efa, _cfdb := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _cfdb == nil {
		_bga.Height = int(_efa)
	}
	_fdbd, _cfdb := GetNumberAsInt64(params.Get("\u0051u\u0061\u006c\u0069\u0074\u0079"))
	if _cfdb == nil {
		_bga.Quality = int(_fdbd)
	}
}

// GetXrefOffset returns the offset of the xref table.
func (_ggfa *PdfParser) GetXrefOffset() int64 { return _ggfa._egfag }

// PdfObjectFloat represents the primitive PDF floating point numerical object.
type PdfObjectFloat float64

// GetFilterName returns the name of the encoding filter.
func (_ggcg *JBIG2Encoder) GetFilterName() string { return StreamEncodingFilterNameJBIG2 }
func (_afag *ASCII85Encoder) base256Tobase85(_eged uint32) [5]byte {
	_eceg := [5]byte{0, 0, 0, 0, 0}
	_dadg := _eged
	for _ebcf := 0; _ebcf < 5; _ebcf++ {
		_dgfae := uint32(1)
		for _ebcgd := 0; _ebcgd < 4-_ebcf; _ebcgd++ {
			_dgfae *= 85
		}
		_gaaf := _dadg / _dgfae
		_dadg = _dadg % _dgfae
		_eceg[_ebcf] = byte(_gaaf)
	}
	return _eceg
}

// EncodeBytes ASCII encodes the passed in slice of bytes.
func (_dbec *ASCIIHexEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _ccec _fd.Buffer
	for _, _fefa := range data {
		_ccec.WriteString(_ac.Sprintf("\u0025\u002e\u0032X\u0020", _fefa))
	}
	_ccec.WriteByte('>')
	return _ccec.Bytes(), nil
}

// MakeStringFromBytes creates an PdfObjectString from a byte array.
// This is more natural than MakeString as `data` is usually not utf-8 encoded.
func MakeStringFromBytes(data []byte) *PdfObjectString { return MakeString(string(data)) }
func _acbd(_ggfe string) (PdfObjectReference, error) {
	_fgca := PdfObjectReference{}
	_dedg := _aea.FindStringSubmatch(_ggfe)
	if len(_dedg) < 3 {
		_ae.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _fgca, _d.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_cbbd, _ := _dg.Atoi(_dedg[1])
	_gfad, _ := _dg.Atoi(_dedg[2])
	_fgca.ObjectNumber = int64(_cbbd)
	_fgca.GenerationNumber = int64(_gfad)
	return _fgca, nil
}

// DecodeStream decodes the stream data and returns the decoded data.
// An error is returned upon failure.
func DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_ae.Log.Trace("\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_baag, _bcabd := NewEncoderFromStream(streamObj)
	if _bcabd != nil {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _bcabd)
		return nil, _bcabd
	}
	_ae.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076\u000a", _baag)
	_egbef, _bcabd := _baag.DecodeStream(streamObj)
	if _bcabd != nil {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _bcabd)
		return nil, _bcabd
	}
	return _egbef, nil
}

// NewLZWEncoder makes a new LZW encoder with default parameters.
func NewLZWEncoder() *LZWEncoder {
	_gfcd := &LZWEncoder{}
	_gfcd.Predictor = 1
	_gfcd.BitsPerComponent = 8
	_gfcd.Colors = 1
	_gfcd.Columns = 1
	_gfcd.EarlyChange = 1
	return _gfcd
}

type objectStream struct {
	N   int
	_fg []byte
	_db map[int]int64
}

func _eeaf(_egcag *PdfObjectDictionary) (_ebcee *_ce.ImageBase) {
	var (
		_efac  *PdfObjectInteger
		_fegcg bool
	)
	if _efac, _fegcg = _egcag.Get("\u0057\u0069\u0064t\u0068").(*PdfObjectInteger); _fegcg {
		_ebcee = &_ce.ImageBase{Width: int(*_efac)}
	} else {
		return nil
	}
	if _efac, _fegcg = _egcag.Get("\u0048\u0065\u0069\u0067\u0068\u0074").(*PdfObjectInteger); _fegcg {
		_ebcee.Height = int(*_efac)
	}
	if _efac, _fegcg = _egcag.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074").(*PdfObjectInteger); _fegcg {
		_ebcee.BitsPerComponent = int(*_efac)
	}
	if _efac, _fegcg = _egcag.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073").(*PdfObjectInteger); _fegcg {
		_ebcee.ColorComponents = int(*_efac)
	}
	return _ebcee
}

// HasEOLAfterHeader gets information if there is a EOL after the version header.
func (_egdb ParserMetadata) HasEOLAfterHeader() bool { return _egdb._egad }

// UpdateParams updates the parameter values of the encoder.
func (_adbf *LZWEncoder) UpdateParams(params *PdfObjectDictionary) {
	_eggc, _gffd := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _gffd == nil {
		_adbf.Predictor = int(_eggc)
	}
	_fag, _gffd := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _gffd == nil {
		_adbf.BitsPerComponent = int(_fag)
	}
	_ffd, _gffd := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _gffd == nil {
		_adbf.Columns = int(_ffd)
	}
	_fgce, _gffd := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _gffd == nil {
		_adbf.Colors = int(_fgce)
	}
	_ddd, _gffd := GetNumberAsInt64(params.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065"))
	if _gffd == nil {
		_adbf.EarlyChange = int(_ddd)
	}
}

var _cacae = _ba.MustCompile("\u0073t\u0061r\u0074\u0078\u003f\u0072\u0065f\u005c\u0073*\u0028\u005c\u0064\u002b\u0029")

// ParseNumber parses a numeric objects from a buffered stream.
// Section 7.3.3.
// Integer or Float.
//
// An integer shall be written as one or more decimal digits optionally
// preceded by a sign. The value shall be interpreted as a signed
// decimal integer and shall be converted to an integer object.
//
// A real value shall be written as one or more decimal digits with an
// optional sign and a leading, trailing, or embedded PERIOD (2Eh)
// (decimal point). The value shall be interpreted as a real number
// and shall be converted to a real object.
//
// Regarding exponential numbers: 7.3.3 Numeric Objects:
// A conforming writer shall not use the PostScript syntax for numbers
// with non-decimal radices (such as 16#FFFE) or in exponential format
// (such as 6.02E23).
// Nonetheless, we sometimes get numbers with exponential format, so
// we will support it in the reader (no confusion with other types, so
// no compromise).
func ParseNumber(buf *_acg.Reader) (PdfObject, error) {
	_eefaee := false
	_dbcf := true
	var _deef _fd.Buffer
	for {
		if _ae.Log.IsLogLevel(_ae.LogLevelTrace) {
			_ae.Log.Trace("\u0050\u0061\u0072\u0073in\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0022\u0025\u0073\u0022", _deef.String())
		}
		_bbee, _cacgb := buf.Peek(1)
		if _cacgb == _dgf.EOF {
			break
		}
		if _cacgb != nil {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0025\u0073", _cacgb)
			return nil, _cacgb
		}
		if _dbcf && (_bbee[0] == '-' || _bbee[0] == '+') {
			_cgab, _ := buf.ReadByte()
			_deef.WriteByte(_cgab)
			_dbcf = false
		} else if IsDecimalDigit(_bbee[0]) {
			_efeb, _ := buf.ReadByte()
			_deef.WriteByte(_efeb)
		} else if _bbee[0] == '.' {
			_adbd, _ := buf.ReadByte()
			_deef.WriteByte(_adbd)
			_eefaee = true
		} else if _bbee[0] == 'e' || _bbee[0] == 'E' {
			_cdbe, _ := buf.ReadByte()
			_deef.WriteByte(_cdbe)
			_eefaee = true
			_dbcf = true
		} else {
			break
		}
	}
	var _gddf PdfObject
	if _eefaee {
		_acfcb, _caaee := _dg.ParseFloat(_deef.String(), 64)
		if _caaee != nil {
			_ae.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025v\u0020\u0065\u0072\u0072\u003d\u0025v\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0030\u002e\u0020\u004fu\u0074\u0070u\u0074\u0020\u006d\u0061y\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074", _deef.String(), _caaee)
			_acfcb = 0.0
		}
		_gcef := PdfObjectFloat(_acfcb)
		_gddf = &_gcef
	} else {
		_bdeb, _abdcd := _dg.ParseInt(_deef.String(), 10, 64)
		if _abdcd != nil {
			_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0025\u0076\u0020\u0065\u0072\u0072\u003d%\u0076\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 \u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006ec\u006f\u0072r\u0065c\u0074", _deef.String(), _abdcd)
			_bdeb = 0
		}
		_ebbeb := PdfObjectInteger(_bdeb)
		_gddf = &_ebbeb
	}
	return _gddf, nil
}

// LookupByNumber looks up a PdfObject by object number.  Returns an error on failure.
func (_ccb *PdfParser) LookupByNumber(objNumber int) (PdfObject, error) {
	_acgf, _, _eff := _ccb.lookupByNumberWrapper(objNumber, true)
	return _acgf, _eff
}

// GetFilterName returns the name of the encoding filter.
func (_fdca *ASCII85Encoder) GetFilterName() string { return StreamEncodingFilterNameASCII85 }
func _gcba(_aeed int) cryptFilters                  { return cryptFilters{_egb: _ebd.NewFilterV2(_aeed)} }

// GetNameVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetNameVal(obj PdfObject) (_gcag string, _gffda bool) {
	_gcbb, _gffda := TraceToDirectObject(obj).(*PdfObjectName)
	if _gffda {
		return string(*_gcbb), true
	}
	return
}

// GetTrailer returns the PDFs trailer dictionary. The trailer dictionary is typically the starting point for a PDF,
// referencing other key objects that are important in the document structure.
func (_deba *PdfParser) GetTrailer() *PdfObjectDictionary { return _deba._dfec }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_efea *RawEncoder) MakeDecodeParams() PdfObject { return nil }

// DecodeBytes decodes a slice of JPX encoded bytes and returns the result.
func (_gffb *JPXEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _gffb.GetFilterName())
	return encoded, ErrNoJPXDecode
}

var _cbae = _ba.MustCompile("\u0025\u0025\u0045\u004f\u0046\u003f")

func (_gdaf *JBIG2Image) toBitmap() (_afbf *_gg.Bitmap, _eabb error) {
	const _baed = "\u004a\u0042\u0049\u00472I\u006d\u0061\u0067\u0065\u002e\u0074\u006f\u0042\u0069\u0074\u006d\u0061\u0070"
	if _gdaf.Data == nil {
		return nil, _eb.Error(_baed, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _gdaf.Width == 0 || _gdaf.Height == 0 {
		return nil, _eb.Error(_baed, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if _gdaf.HasPadding {
		_afbf, _eabb = _gg.NewWithData(_gdaf.Width, _gdaf.Height, _gdaf.Data)
	} else {
		_afbf, _eabb = _gg.NewWithUnpaddedData(_gdaf.Width, _gdaf.Height, _gdaf.Data)
	}
	if _eabb != nil {
		return nil, _eb.Wrap(_eabb, _baed, "")
	}
	return _afbf, nil
}

// GetRevision returns PdfParser for the specific version of the Pdf document.
func (_ccfbf *PdfParser) GetRevision(revisionNumber int) (*PdfParser, error) {
	_fadgf := _ccfbf._fdbdc
	if _fadgf == revisionNumber {
		return _ccfbf, nil
	}
	if _fadgf < revisionNumber {
		return nil, _d.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0072\u0065\u0076\u0069\u0073i\u006fn\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	if _ccfbf._ddcg[revisionNumber] != nil {
		return _ccfbf._ddcg[revisionNumber], nil
	}
	_dccc := _ccfbf
	for ; _fadgf > revisionNumber; _fadgf-- {
		_cafe, _fddg := _dccc.GetPreviousRevisionParser()
		if _fddg != nil {
			return nil, _fddg
		}
		_ccfbf._ddcg[_fadgf-1] = _cafe
		_ccfbf._bgcdd[_dccc] = _cafe
		_dccc = _cafe
	}
	return _dccc, nil
}

// ParseIndirectObject parses an indirect object from the input stream. Can also be an object stream.
// Returns the indirect object (*PdfIndirectObject) or the stream object (*PdfObjectStream).
func (_dgggb *PdfParser) ParseIndirectObject() (PdfObject, error) {
	_gebb := PdfIndirectObject{}
	_gebb._ffgd = _dgggb
	_ae.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_acbge, _ecbfc := _dgggb._eecea.Peek(20)
	if _ecbfc != nil {
		if _ecbfc != _dgf.EOF {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return &_gebb, _ecbfc
		}
	}
	_ae.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_acbge))
	_cafc := _ddce.FindStringSubmatchIndex(string(_acbge))
	if len(_cafc) < 6 {
		if _ecbfc == _dgf.EOF {
			return nil, _ecbfc
		}
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_acbge))
		return &_gebb, _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_dgggb._eecea.Discard(_cafc[0])
	_ae.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _cafc)
	_cdea := _cafc[1] - _cafc[0]
	_dedge := make([]byte, _cdea)
	_, _ecbfc = _dgggb.ReadAtLeast(_dedge, _cdea)
	if _ecbfc != nil {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _ecbfc)
		return nil, _ecbfc
	}
	_ae.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _dedge)
	_afffe := _ddce.FindStringSubmatch(string(_dedge))
	if len(_afffe) < 3 {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_dedge))
		return &_gebb, _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_acgb, _ := _dg.Atoi(_afffe[1])
	_dbbaf, _ := _dg.Atoi(_afffe[2])
	_gebb.ObjectNumber = int64(_acgb)
	_gebb.GenerationNumber = int64(_dbbaf)
	for {
		_defb, _adde := _dgggb._eecea.Peek(2)
		if _adde != nil {
			return &_gebb, _adde
		}
		_ae.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_defb), string(_defb))
		if IsWhiteSpace(_defb[0]) {
			_dgggb.skipSpaces()
		} else if _defb[0] == '%' {
			_dgggb.skipComments()
		} else if (_defb[0] == '<') && (_defb[1] == '<') {
			_ae.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_gebb.PdfObject, _adde = _dgggb.ParseDict()
			_ae.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _adde)
			if _adde != nil {
				return &_gebb, _adde
			}
			_ae.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_defb[0] == '/') || (_defb[0] == '(') || (_defb[0] == '[') || (_defb[0] == '<') {
			_gebb.PdfObject, _adde = _dgggb.parseObject()
			if _adde != nil {
				return &_gebb, _adde
			}
			_ae.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else if _defb[0] == ']' {
			_ae.Log.Debug("\u0057\u0041\u0052\u004e\u0049N\u0047\u003a\u0020\u0027\u005d\u0027 \u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u006eo\u0074\u0020\u0062\u0065i\u006e\u0067\u0020\u0075\u0073\u0065d\u0020\u0061\u0073\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006e\u0064\u0069n\u0067\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			_dgggb._eecea.Discard(1)
		} else {
			if _defb[0] == 'e' {
				_aacg, _acgff := _dgggb.readTextLine()
				if _acgff != nil {
					return nil, _acgff
				}
				if len(_aacg) >= 6 && _aacg[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _defb[0] == 's' {
				_defb, _ = _dgggb._eecea.Peek(10)
				if string(_defb[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_fcgae := 6
					if len(_defb) > 6 {
						if IsWhiteSpace(_defb[_fcgae]) && _defb[_fcgae] != '\r' && _defb[_fcgae] != '\n' {
							_ae.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0050\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_dgggb._aecec._gbbe = true
							_fcgae++
						}
						if _defb[_fcgae] == '\r' {
							_fcgae++
							if _defb[_fcgae] == '\n' {
								_fcgae++
							}
						} else if _defb[_fcgae] == '\n' {
							_fcgae++
						} else {
							_dgggb._aecec._gbbe = true
						}
					}
					_dgggb._eecea.Discard(_fcgae)
					_ggdd, _agegg := _gebb.PdfObject.(*PdfObjectDictionary)
					if !_agegg {
						return nil, _d.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_ae.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _ggdd)
					_abdd, _afdfg := _dgggb.traceStreamLength(_ggdd.Get("\u004c\u0065\u006e\u0067\u0074\u0068"))
					if _afdfg != nil {
						_ae.Log.Debug("\u0046\u0061\u0069l\u0020\u0074\u006f\u0020t\u0072\u0061\u0063\u0065\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076", _afdfg)
						return nil, _afdfg
					}
					_ae.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _abdd)
					_aegb, _aeea := _abdd.(*PdfObjectInteger)
					if !_aeea {
						return nil, _d.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_cbgd := *_aegb
					if _cbgd < 0 {
						return nil, _d.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					_geeg := _dgggb.GetFileOffset()
					_cfcg := _dgggb.xrefNextObjectOffset(_geeg)
					if _geeg+int64(_cbgd) > _cfcg && _cfcg > _geeg {
						_ae.Log.Debug("E\u0078\u0070\u0065\u0063te\u0064 \u0065\u006e\u0064\u0069\u006eg\u0020\u0061\u0074\u0020\u0025\u0064", _geeg+int64(_cbgd))
						_ae.Log.Debug("\u004e\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020\u0061t\u0020\u0025\u0064", _cfcg)
						_baa := _cfcg - _geeg - 17
						if _baa < 0 {
							return nil, _d.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020go\u0069\u006e\u0067 \u0070\u0061s\u0074\u0020\u0062\u006f\u0075\u006ed\u0061\u0072i\u0065\u0073")
						}
						_ae.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u0020c\u006f\u0072\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020%\u0064\u002e\u002e\u002e", _baa)
						_cbgd = PdfObjectInteger(_baa)
						_ggdd.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(_baa))
					}
					if int64(_cbgd) > _dgggb._fcca {
						_ae.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _d.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_acac := make([]byte, _cbgd)
					_, _afdfg = _dgggb.ReadAtLeast(_acac, int(_cbgd))
					if _afdfg != nil {
						_ae.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_acac), _acac)
						_ae.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _afdfg)
						return nil, _afdfg
					}
					_fdcdc := PdfObjectStream{}
					_fdcdc.Stream = _acac
					_fdcdc.PdfObjectDictionary = _gebb.PdfObject.(*PdfObjectDictionary)
					_fdcdc.ObjectNumber = _gebb.ObjectNumber
					_fdcdc.GenerationNumber = _gebb.GenerationNumber
					_fdcdc.PdfObjectReference._ffgd = _dgggb
					_dgggb.skipSpaces()
					_dgggb._eecea.Discard(9)
					_dgggb.skipSpaces()
					return &_fdcdc, nil
				}
			}
			_gebb.PdfObject, _adde = _dgggb.parseObject()
			if _gebb.PdfObject == nil {
				_ae.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
				_gebb.PdfObject = MakeNull()
			}
			return &_gebb, _adde
		}
	}
	if _gebb.PdfObject == nil {
		_ae.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
		_gebb.PdfObject = MakeNull()
	}
	_ae.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_gebb, nil
}

// DrawableImage is same as golang image/draw's Image interface that allow drawing images.
type DrawableImage interface {
	ColorModel() _ed.Model
	Bounds() _ea.Rectangle
	At(_fbg, _adec int) _ed.Color
	Set(_bbgf, _bgdc int, _gdbcg _ed.Color)
}

// HasInvalidSeparationAfterXRef implements core.ParserMetadata interface.
func (_acde ParserMetadata) HasInvalidSeparationAfterXRef() bool { return _acde._bega }
func (_cdf *PdfCrypt) saveCryptFilters(_cdfc *PdfObjectDictionary) error {
	if _cdf._ddc.V < 4 {
		return _d.New("\u0063\u0061\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020V\u003e\u003d\u0034")
	}
	_gcd := MakeDict()
	_cdfc.Set("\u0043\u0046", _gcd)
	for _fbb, _ggea := range _cdf._agfd {
		if _fbb == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			continue
		}
		_cfc := _dgfa(_ggea, "")
		_gcd.Set(PdfObjectName(_fbb), _cfc)
	}
	_cdfc.Set("\u0053\u0074\u0072\u0046", MakeName(_cdf._dee))
	_cdfc.Set("\u0053\u0074\u006d\u0046", MakeName(_cdf._bba))
	return nil
}

// DecodeImages decodes the page images from the jbig2 'encoded' data input.
// The jbig2 document may contain multiple pages, thus the function can return multiple
// images. The images order corresponds to the page number.
func (_dabad *JBIG2Encoder) DecodeImages(encoded []byte) ([]_ea.Image, error) {
	const _dbea = "\u004aB\u0049\u0047\u0032\u0045n\u0063\u006f\u0064\u0065\u0072.\u0044e\u0063o\u0064\u0065\u0049\u006d\u0061\u0067\u0065s"
	_dacea, _cbgg := _gb.Decode(encoded, _gb.Parameters{}, _dabad.Globals.ToDocumentGlobals())
	if _cbgg != nil {
		return nil, _eb.Wrap(_cbgg, _dbea, "")
	}
	_aacde, _cbgg := _dacea.PageNumber()
	if _cbgg != nil {
		return nil, _eb.Wrap(_cbgg, _dbea, "")
	}
	_ecbe := []_ea.Image{}
	var _daad _ea.Image
	for _bbeb := 1; _bbeb <= _aacde; _bbeb++ {
		_daad, _cbgg = _dacea.DecodePageImage(_bbeb)
		if _cbgg != nil {
			return nil, _eb.Wrapf(_cbgg, _dbea, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _bbeb)
		}
		_ecbe = append(_ecbe, _daad)
	}
	return _ecbe, nil
}
func (_eccf *PdfParser) loadXrefs() (*PdfObjectDictionary, error) {
	_eccf._fbab.ObjectMap = make(map[int]XrefObject)
	_eccf._acab = make(objectStreams)
	_adca, _feacf := _eccf._fdee.Seek(0, _dgf.SeekEnd)
	if _feacf != nil {
		return nil, _feacf
	}
	_ae.Log.Trace("\u0066s\u0069\u007a\u0065\u003a\u0020\u0025d", _adca)
	_eccf._fcca = _adca
	_feacf = _eccf.seekToEOFMarker(_adca)
	if _feacf != nil {
		_ae.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0073\u0065\u0065\u006b\u0020\u0074\u006f\u0020\u0065\u006f\u0066\u0020\u006d\u0061\u0072\u006b\u0065\u0072: \u0025\u0076", _feacf)
		return nil, _feacf
	}
	_afce, _feacf := _eccf._fdee.Seek(0, _dgf.SeekCurrent)
	if _feacf != nil {
		return nil, _feacf
	}
	var _cfccf int64 = 64
	_egbf := _afce - _cfccf
	if _egbf < 0 {
		_egbf = 0
	}
	_, _feacf = _eccf._fdee.Seek(_egbf, _dgf.SeekStart)
	if _feacf != nil {
		return nil, _feacf
	}
	_ebeed := make([]byte, _cfccf)
	_, _feacf = _eccf._fdee.Read(_ebeed)
	if _feacf != nil {
		_ae.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u003a\u0020\u0025\u0076", _feacf)
		return nil, _feacf
	}
	_abee := _cacae.FindStringSubmatch(string(_ebeed))
	if len(_abee) < 2 {
		_ae.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020n\u006f\u0074\u0020f\u006fu\u006e\u0064\u0021")
		return nil, _d.New("\u0073\u0074\u0061\u0072tx\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	if len(_abee) > 2 {
		_ae.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u004du\u006c\u0074\u0069\u0070\u006c\u0065\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029\u0021", _ebeed)
		return nil, _d.New("m\u0075\u006c\u0074\u0069\u0070\u006ce\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065f\u0020\u0065\u006et\u0072i\u0065\u0073\u003f")
	}
	_dcdf, _ := _dg.ParseInt(_abee[1], 10, 64)
	_ae.Log.Trace("\u0073t\u0061r\u0074\u0078\u0072\u0065\u0066\u0020\u0061\u0074\u0020\u0025\u0064", _dcdf)
	if _dcdf > _adca {
		_ae.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u006f\u0066f\u0073e\u0074 \u006fu\u0074\u0073\u0069\u0064\u0065\u0020\u006f\u0066\u0020\u0066\u0069\u006c\u0065")
		_ae.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0072e\u0070\u0061\u0069\u0072")
		_dcdf, _feacf = _eccf.repairLocateXref()
		if _feacf != nil {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0065\u0070\u0061\u0069\u0072\u0020\u0061\u0074\u0074\u0065\u006d\u0070t\u0020\u0066\u0061\u0069\u006c\u0065\u0064 \u0028\u0025\u0073\u0029")
			return nil, _feacf
		}
	}
	_eccf._fdee.Seek(_dcdf, _dgf.SeekStart)
	_eccf._eecea = _acg.NewReader(_eccf._fdee)
	_fceg, _feacf := _eccf.parseXref()
	if _feacf != nil {
		return nil, _feacf
	}
	_bgeb := _fceg.Get("\u0058R\u0065\u0066\u0053\u0074\u006d")
	if _bgeb != nil {
		_eafb, _bdec := _bgeb.(*PdfObjectInteger)
		if !_bdec {
			return nil, _d.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0021=\u0020\u0069\u006e\u0074")
		}
		_, _feacf = _eccf.parseXrefStream(_eafb)
		if _feacf != nil {
			return nil, _feacf
		}
	}
	var _eabbe []int64
	_cdgcb := func(_facc int64, _gaba []int64) bool {
		for _, _gbaf := range _gaba {
			if _gbaf == _facc {
				return true
			}
		}
		return false
	}
	_bgeb = _fceg.Get("\u0050\u0072\u0065\u0076")
	for _bgeb != nil {
		_cbf, _ebgf := _bgeb.(*PdfObjectInteger)
		if !_ebgf {
			_ae.Log.Debug("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020P\u0072\u0065\u0076\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _bgeb)
			return _fceg, nil
		}
		_cfgd := *_cbf
		_ae.Log.Trace("\u0041\u006eot\u0068\u0065\u0072 \u0050\u0072\u0065\u0076 xr\u0065f \u0074\u0061\u0062\u006c\u0065\u0020\u006fbj\u0065\u0063\u0074\u0020\u0061\u0074\u0020%\u0064", _cfgd)
		_eccf._fdee.Seek(int64(_cfgd), _dgf.SeekStart)
		_eccf._eecea = _acg.NewReader(_eccf._fdee)
		_cgbce, _ebdf := _eccf.parseXref()
		if _ebdf != nil {
			_ae.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0045\u0072\u0072\u006f\u0072\u0020-\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069n\u0067\u0020\u0061\u006e\u006f\u0074\u0068\u0065\u0072\u0020\u0028\u0050re\u0076\u0029\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
			_ae.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0063\u006f\u006e\u0074\u0069n\u0075\u0065\u0020\u0062\u0079\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067\u0020\u0069\u0074")
			break
		}
		_eccf._edcac = append(_eccf._edcac, int64(_cfgd))
		_bgeb = _cgbce.Get("\u0050\u0072\u0065\u0076")
		if _bgeb != nil {
			_fdggb := *(_bgeb.(*PdfObjectInteger))
			if _cdgcb(int64(_fdggb), _eabbe) {
				_ae.Log.Debug("\u0050\u0072ev\u0065\u006e\u0074i\u006e\u0067\u0020\u0063irc\u0075la\u0072\u0020\u0078\u0072\u0065\u0066\u0020re\u0066\u0065\u0072\u0065\u006e\u0063\u0069n\u0067")
				break
			}
			_eabbe = append(_eabbe, int64(_fdggb))
		}
	}
	return _fceg, nil
}

// PdfIndirectObject represents the primitive PDF indirect object.
type PdfIndirectObject struct {
	PdfObjectReference
	PdfObject
}

func (_fcg *PdfCrypt) authenticate(_eab []byte) (bool, error) {
	_fcg._gfg = false
	_bfdg := _fcg.securityHandler()
	_fgf, _gca, _cba := _bfdg.Authenticate(&_fcg._gge, _eab)
	if _cba != nil {
		return false, _cba
	} else if _gca == 0 || len(_fgf) == 0 {
		return false, nil
	}
	_fcg._gfg = true
	_fcg._dadd = _fgf
	return true, nil
}

// Set sets the PdfObject at index i of the streams. An error is returned if the index is outside bounds.
func (_cfge *PdfObjectStreams) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_cfge._dgebc) {
		return _d.New("\u004f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_cfge._dgebc[i] = obj
	return nil
}
func (_gdb *PdfParser) lookupByNumber(_fgg int, _dba bool) (PdfObject, bool, error) {
	_bbe, _caba := _gdb.ObjCache[_fgg]
	if _caba {
		_ae.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0063a\u0063\u0068\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0025\u0064", _fgg)
		return _bbe, false, nil
	}
	if _gdb._ccge == nil {
		_gdb._ccge = map[int]bool{}
	}
	if _gdb._ccge[_fgg] {
		_ae.Log.Debug("ER\u0052\u004f\u0052\u003a\u0020\u004c\u006fok\u0075\u0070\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0069\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0069\u006e\u0020\u0070\u0072\u006f\u0067\u0072\u0065\u0073\u0073\u0020\u002d\u0020\u0072\u0065c\u0075\u0072\u0073\u0069\u0076\u0065 \u006c\u006f\u006f\u006b\u0075\u0070\u0020\u0061\u0074t\u0065m\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064", _fgg)
		return nil, false, _d.New("\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006f\u006f\u006b\u0075p\u0020a\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064")
	}
	_gdb._ccge[_fgg] = true
	defer delete(_gdb._ccge, _fgg)
	_fcb, _caba := _gdb._fbab.ObjectMap[_fgg]
	if !_caba {
		_ae.Log.Trace("\u0055\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u006c\u006f\u0063\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006e\u0020\u0078\u0072\u0065\u0066\u0073\u0021 \u002d\u0020\u0052\u0065\u0074u\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		var _gdbc PdfObjectNull
		return &_gdbc, false, nil
	}
	_ae.Log.Trace("L\u006fo\u006b\u0075\u0070\u0020\u006f\u0062\u006a\u0020n\u0075\u006d\u0062\u0065r \u0025\u0064", _fgg)
	if _fcb.XType == XrefTypeTableEntry {
		_ae.Log.Trace("\u0078r\u0065f\u006f\u0062\u006a\u0020\u006fb\u006a\u0020n\u0075\u006d\u0020\u0025\u0064", _fcb.ObjectNumber)
		_ae.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u0067e\u006e\u0020\u0025\u0064", _fcb.Generation)
		_ae.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0064", _fcb.Offset)
		_gdb._fdee.Seek(_fcb.Offset, _dgf.SeekStart)
		_gdb._eecea = _acg.NewReader(_gdb._fdee)
		_cca, _ebf := _gdb.ParseIndirectObject()
		if _ebf != nil {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006ce\u0064\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029", _ebf)
			if _dba {
				_ae.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0072\u0065\u0070\u0061\u0069r\u0020\u0078\u0072\u0065\u0066\u0073\u0020\u0028\u0074\u006f\u0070\u0020d\u006f\u0077\u006e\u0029")
				_gcf, _ecd := _gdb.repairRebuildXrefsTopDown()
				if _ecd != nil {
					_ae.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020r\u0065\u0070\u0061\u0069\u0072\u0020\u0028\u0025\u0073\u0029", _ecd)
					return nil, false, _ecd
				}
				_gdb._fbab = *_gcf
				return _gdb.lookupByNumber(_fgg, false)
			}
			return nil, false, _ebf
		}
		if _dba {
			_bgb, _, _ := _def(_cca)
			if int(_bgb) != _fgg {
				_ae.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0078\u0072\u0065\u0066s\u003a \u0052e\u0062\u0075\u0069\u006c\u0064\u0069\u006eg")
				_gcb := _gdb.rebuildXrefTable()
				if _gcb != nil {
					return nil, false, _gcb
				}
				_gdb.ObjCache = objectCache{}
				return _gdb.lookupByNumberWrapper(_fgg, false)
			}
		}
		_ae.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006f\u0062\u006a")
		_gdb.ObjCache[_fgg] = _cca
		return _cca, false, nil
	} else if _fcb.XType == XrefTypeObjectStream {
		_ae.Log.Trace("\u0078r\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0021")
		_ae.Log.Trace("\u003e\u004c\u006f\u0061\u0064\u0020\u0076\u0069\u0061\u0020\u004f\u0053\u0021")
		_ae.Log.Trace("\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0069\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020%\u0064\u002f\u0025\u0064", _fcb.OsObjNumber, _fcb.OsObjIndex)
		if _fcb.OsObjNumber == _fgg {
			_ae.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0043i\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065f\u0065\u0072\u0065n\u0063e\u0021\u003f\u0021")
			return nil, true, _d.New("\u0078\u0072\u0065f \u0063\u0069\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		if _, _cb := _gdb._fbab.ObjectMap[_fcb.OsObjNumber]; _cb {
			_bbf, _agf := _gdb.lookupObjectViaOS(_fcb.OsObjNumber, _fgg)
			if _agf != nil {
				_ae.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069n\u0067\u0020\u0045\u0052\u0052\u0020\u0028\u0025\u0073\u0029", _agf)
				return nil, true, _agf
			}
			_ae.Log.Trace("\u003c\u004c\u006f\u0061\u0064\u0065\u0064\u0020\u0076i\u0061\u0020\u004f\u0053")
			_gdb.ObjCache[_fgg] = _bbf
			if _gdb._bffd != nil {
				_gdb._bffd._dgd[_bbf] = true
			}
			return _bbf, true, nil
		}
		_ae.Log.Debug("\u003f\u003f\u0020\u0042\u0065\u006c\u006f\u006eg\u0073\u0020\u0074o \u0061\u0020\u006e\u006f\u006e\u002dc\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002e.\u002e\u0021")
		return nil, true, _d.New("\u006f\u0073\u0020\u0062\u0065\u006c\u006fn\u0067\u0073\u0020t\u006f\u0020\u0061\u0020n\u006f\u006e\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, false, _d.New("\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0078\u0072\u0065\u0066 \u0074\u0079\u0070\u0065")
}

// GetNumbersAsFloat converts a list of pdf objects representing floats or integers to a slice of
// float64 values.
func GetNumbersAsFloat(objects []PdfObject) (_fbbd []float64, _ddee error) {
	for _, _bgedb := range objects {
		_bcbf, _cbbe := GetNumberAsFloat(_bgedb)
		if _cbbe != nil {
			return nil, _cbbe
		}
		_fbbd = append(_fbbd, _bcbf)
	}
	return _fbbd, nil
}

// WriteString outputs the object as it is to be written to file.
func (_daead *PdfObjectStreams) WriteString() string {
	var _gedd _agg.Builder
	_gedd.WriteString(_dg.FormatInt(_daead.ObjectNumber, 10))
	_gedd.WriteString("\u0020\u0030\u0020\u0052")
	return _gedd.String()
}
func _eec(_beb XrefTable) {
	_ae.Log.Debug("\u003dX\u003d\u0058\u003d\u0058\u003d")
	_ae.Log.Debug("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_gef := 0
	for _, _dec := range _beb.ObjectMap {
		_ae.Log.Debug("i\u002b\u0031\u003a\u0020\u0025\u0064 \u0028\u006f\u0062\u006a\u0020\u006eu\u006d\u003a\u0020\u0025\u0064\u0020\u0067e\u006e\u003a\u0020\u0025\u0064\u0029\u0020\u002d\u003e\u0020%\u0064", _gef+1, _dec.ObjectNumber, _dec.Generation, _dec.Offset)
		_gef++
	}
}

// GetDict returns the *PdfObjectDictionary represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetDict(obj PdfObject) (_cfcag *PdfObjectDictionary, _bagf bool) {
	_cfcag, _bagf = TraceToDirectObject(obj).(*PdfObjectDictionary)
	return _cfcag, _bagf
}

// NewEncoderFromStream creates a StreamEncoder based on the stream's dictionary.
func NewEncoderFromStream(streamObj *PdfObjectStream) (StreamEncoder, error) {
	_ddeg := TraceToDirectObject(streamObj.PdfObjectDictionary.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if _ddeg == nil {
		return NewRawEncoder(), nil
	}
	if _, _faagb := _ddeg.(*PdfObjectNull); _faagb {
		return NewRawEncoder(), nil
	}
	_adad, _dcdbb := _ddeg.(*PdfObjectName)
	if !_dcdbb {
		_eada, _bedd := _ddeg.(*PdfObjectArray)
		if !_bedd {
			return nil, _ac.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _eada.Len() == 0 {
			return NewRawEncoder(), nil
		}
		if _eada.Len() != 1 {
			_dabg, _fccac := _dfa(streamObj)
			if _fccac != nil {
				_ae.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _fccac)
				return nil, _fccac
			}
			_ae.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _dabg)
			return _dabg, nil
		}
		_ddeg = _eada.Get(0)
		_adad, _bedd = _ddeg.(*PdfObjectName)
		if !_bedd {
			return nil, _ac.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	if _cfac, _eagg := _dggc.Load(_adad.String()); _eagg {
		return _cfac.(StreamEncoder), nil
	}
	switch *_adad {
	case StreamEncodingFilterNameFlate:
		return _cbcf(streamObj, nil)
	case StreamEncodingFilterNameLZW:
		return _cgbc(streamObj, nil)
	case StreamEncodingFilterNameDCT:
		return _faac(streamObj, nil)
	case StreamEncodingFilterNameRunLength:
		return _gfca(streamObj, nil)
	case StreamEncodingFilterNameASCIIHex:
		return NewASCIIHexEncoder(), nil
	case StreamEncodingFilterNameASCII85, "\u0041\u0038\u0035":
		return NewASCII85Encoder(), nil
	case StreamEncodingFilterNameCCITTFax:
		return _acb(streamObj, nil)
	case StreamEncodingFilterNameJBIG2:
		return _adef(streamObj, nil)
	case StreamEncodingFilterNameJPX:
		return NewJPXEncoder(), nil
	}
	_ae.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006fd\u0021")
	return nil, _ac.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064i\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u0020\u0028\u0025\u0073\u0029", *_adad)
}

// Encrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_bcb *PdfCrypt) Encrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _bcb.isEncrypted(obj) {
		return nil
	}
	switch _bdf := obj.(type) {
	case *PdfIndirectObject:
		_bcb._egg[_bdf] = true
		_ae.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _bdf.ObjectNumber, _bdf.GenerationNumber)
		_edf := _bdf.ObjectNumber
		_acc := _bdf.GenerationNumber
		_affd := _bcb.Encrypt(_bdf.PdfObject, _edf, _acc)
		if _affd != nil {
			return _affd
		}
		return nil
	case *PdfObjectStream:
		_bcb._egg[_bdf] = true
		_gfce := _bdf.PdfObjectDictionary
		if _fcga, _feda := _gfce.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _feda && *_fcga == "\u0058\u0052\u0065\u0066" {
			return nil
		}
		_cce := _bdf.ObjectNumber
		_edca := _bdf.GenerationNumber
		_ae.Log.Trace("\u0045n\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _cce, _edca)
		_gaa := _egb
		if _bcb._ddc.V >= 4 {
			_gaa = _bcb._bba
			_ae.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _bcb._bba)
			if _geff, _ggd := _gfce.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _ggd {
				if _efb, _aef := GetName(_geff.Get(0)); _aef {
					if *_efb == "\u0043\u0072\u0079p\u0074" {
						_gaa = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _dgdf, _fcd := _gfce.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _fcd {
							if _gcg, _dcg := _dgdf.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _dcg {
								if _, _decc := _bcb._agfd[string(*_gcg)]; _decc {
									_ae.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_gcg)
									_gaa = string(*_gcg)
								}
							}
						}
					}
				}
			}
			_ae.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _gaa)
			if _gaa == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_afff := _bcb.Encrypt(_bdf.PdfObjectDictionary, _cce, _edca)
		if _afff != nil {
			return _afff
		}
		_cbac, _afff := _bcb.makeKey(_gaa, uint32(_cce), uint32(_edca), _bcb._dadd)
		if _afff != nil {
			return _afff
		}
		_bdf.Stream, _afff = _bcb.encryptBytes(_bdf.Stream, _gaa, _cbac)
		if _afff != nil {
			return _afff
		}
		_gfce.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_bdf.Stream))))
		return nil
	case *PdfObjectString:
		_ae.Log.Trace("\u0045n\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_ege := _egb
		if _bcb._ddc.V >= 4 {
			_ae.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _bcb._dee)
			if _bcb._dee == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_ege = _bcb._dee
		}
		_cgf, _ebfg := _bcb.makeKey(_ege, uint32(parentObjNum), uint32(parentGenNum), _bcb._dadd)
		if _ebfg != nil {
			return _ebfg
		}
		_ebc := _bdf.Str()
		_gagf := make([]byte, len(_ebc))
		for _gad := 0; _gad < len(_ebc); _gad++ {
			_gagf[_gad] = _ebc[_gad]
		}
		_ae.Log.Trace("\u0045n\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _gagf, _gagf)
		_gagf, _ebfg = _bcb.encryptBytes(_gagf, _ege, _cgf)
		if _ebfg != nil {
			return _ebfg
		}
		_bdf._eeee = string(_gagf)
		return nil
	case *PdfObjectArray:
		for _, _dgb := range _bdf.Elements() {
			_fafd := _bcb.Encrypt(_dgb, parentObjNum, parentGenNum)
			if _fafd != nil {
				return _fafd
			}
		}
		return nil
	case *PdfObjectDictionary:
		_afd := false
		if _gecc := _bdf.Get("\u0054\u0079\u0070\u0065"); _gecc != nil {
			_eggac, _ega := _gecc.(*PdfObjectName)
			if _ega && *_eggac == "\u0053\u0069\u0067" {
				_afd = true
			}
		}
		for _, _cea := range _bdf.Keys() {
			_cbc := _bdf.Get(_cea)
			if _afd && string(_cea) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_cea) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_cea) != "\u0050\u0072\u0065\u0076" && string(_cea) != "\u004c\u0061\u0073\u0074" {
				_fafb := _bcb.Encrypt(_cbc, parentObjNum, parentGenNum)
				if _fafb != nil {
					return _fafb
				}
			}
		}
		return nil
	}
	return nil
}
func _adfaf(_bbcbf, _gfacf, _bbde int) error {
	if _gfacf < 0 || _gfacf > _bbcbf {
		return _d.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0061\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	if _bbde < _gfacf {
		return _d.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0073\u006c\u0069\u0063e\u0020i\u006ed\u0065\u0078\u0020\u0062\u0020\u003c\u0020a")
	}
	if _bbde > _bbcbf {
		return _d.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0062\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	return nil
}
func (_fcgb *PdfParser) parseLinearizedDictionary() (*PdfObjectDictionary, error) {
	_fgea, _cgbb := _fcgb._fdee.Seek(0, _dgf.SeekEnd)
	if _cgbb != nil {
		return nil, _cgbb
	}
	var _gfba int64
	var _bgba int64 = 2048
	for _gfba < _fgea-4 {
		if _fgea <= (_bgba + _gfba) {
			_bgba = _fgea - _gfba
		}
		_, _fffef := _fcgb._fdee.Seek(_gfba, _dgf.SeekStart)
		if _fffef != nil {
			return nil, _fffef
		}
		_faaef := make([]byte, _bgba)
		_, _fffef = _fcgb._fdee.Read(_faaef)
		if _fffef != nil {
			return nil, _fffef
		}
		_ae.Log.Trace("\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0066i\u0072\u0073\u0074\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0022\u0025\u0073\u0022", string(_faaef))
		_bdbb := _ddce.FindAllStringIndex(string(_faaef), -1)
		if _bdbb != nil {
			_bfg := _bdbb[0]
			_ae.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _bdbb)
			_, _ffgbc := _fcgb._fdee.Seek(int64(_bfg[0]), _dgf.SeekStart)
			if _ffgbc != nil {
				return nil, _ffgbc
			}
			_fcgb._eecea = _acg.NewReader(_fcgb._fdee)
			_ffgc, _ffgbc := _fcgb.ParseIndirectObject()
			if _ffgbc != nil {
				return nil, nil
			}
			if _dgfb, _edg := GetIndirect(_ffgc); _edg {
				if _gafa, _gbdb := GetDict(_dgfb.PdfObject); _gbdb {
					if _cdac := _gafa.Get("\u004c\u0069\u006e\u0065\u0061\u0072\u0069\u007a\u0065\u0064"); _cdac != nil {
						return _gafa, nil
					}
					return nil, nil
				}
			}
			return nil, nil
		}
		_gfba += _bgba - 4
	}
	return nil, _d.New("\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064")
}

// IsNullObject returns true if `obj` is a PdfObjectNull.
func IsNullObject(obj PdfObject) bool {
	_, _ffge := TraceToDirectObject(obj).(*PdfObjectNull)
	return _ffge
}

var _dggc _e.Map

// DecodeBytes decodes a slice of JBIG2 encoded bytes and returns the results.
func (_ccga *JBIG2Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	return _df.DecodeBytes(encoded, _gb.Parameters{}, _ccga.Globals)
}
func _def(_gea PdfObject) (int64, int64, error) {
	if _aeg, _dc := _gea.(*PdfIndirectObject); _dc {
		return _aeg.ObjectNumber, _aeg.GenerationNumber, nil
	}
	if _cabg, _dfb := _gea.(*PdfObjectStream); _dfb {
		return _cabg.ObjectNumber, _cabg.GenerationNumber, nil
	}
	return 0, 0, _d.New("\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u002f\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062je\u0063\u0074")
}

const JB2ImageAutoThreshold = -1.0

// DecodeBytes decodes a slice of ASCII encoded bytes and returns the result.
func (_gda *ASCIIHexEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_fdcd := _fd.NewReader(encoded)
	var _afde []byte
	for {
		_cdeg, _egcdg := _fdcd.ReadByte()
		if _egcdg != nil {
			return nil, _egcdg
		}
		if _cdeg == '>' {
			break
		}
		if IsWhiteSpace(_cdeg) {
			continue
		}
		if (_cdeg >= 'a' && _cdeg <= 'f') || (_cdeg >= 'A' && _cdeg <= 'F') || (_cdeg >= '0' && _cdeg <= '9') {
			_afde = append(_afde, _cdeg)
		} else {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0061\u0073\u0063\u0069\u0069 \u0068\u0065\u0078\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072 \u0028\u0025\u0063\u0029", _cdeg)
			return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0073\u0063\u0069\u0069\u0020\u0068e\u0078 \u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0028\u0025\u0063\u0029", _cdeg)
		}
	}
	if len(_afde)%2 == 1 {
		_afde = append(_afde, '0')
	}
	_ae.Log.Trace("\u0049\u006e\u0062\u006f\u0075\u006e\u0064\u0020\u0025\u0073", _afde)
	_bgfa := make([]byte, _cab.DecodedLen(len(_afde)))
	_, _cfca := _cab.Decode(_bgfa, _afde)
	if _cfca != nil {
		return nil, _cfca
	}
	return _bgfa, nil
}
func _bbgd(_egcb PdfObject, _aggcc int, _dafae map[PdfObject]struct{}) error {
	_ae.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072s\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0061 \u0028\u0064\u0065\u0070\u0074\u0068\u0020=\u0020\u0025\u0064\u0029", _aggcc)
	if _, _cbcd := _dafae[_egcb]; _cbcd {
		_ae.Log.Trace("-\u0041\u006c\u0072\u0065ad\u0079 \u0074\u0072\u0061\u0076\u0065r\u0073\u0065\u0064\u002e\u002e\u002e")
		return nil
	}
	_dafae[_egcb] = struct{}{}
	switch _fggdf := _egcb.(type) {
	case *PdfIndirectObject:
		_ebeg := _fggdf
		_ae.Log.Trace("\u0069\u006f\u003a\u0020\u0025\u0073", _ebeg)
		_ae.Log.Trace("\u002d\u0020\u0025\u0073", _ebeg.PdfObject)
		return _bbgd(_ebeg.PdfObject, _aggcc+1, _dafae)
	case *PdfObjectStream:
		_gdbfc := _fggdf
		return _bbgd(_gdbfc.PdfObjectDictionary, _aggcc+1, _dafae)
	case *PdfObjectDictionary:
		_gfed := _fggdf
		_ae.Log.Trace("\u002d\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _gfed)
		for _, _dffc := range _gfed.Keys() {
			_cfeae := _gfed.Get(_dffc)
			if _dcbbe, _gffgd := _cfeae.(*PdfObjectReference); _gffgd {
				_gdfg := _dcbbe.Resolve()
				_gfed.Set(_dffc, _gdfg)
				_bbae := _bbgd(_gdfg, _aggcc+1, _dafae)
				if _bbae != nil {
					return _bbae
				}
			} else {
				_cageg := _bbgd(_cfeae, _aggcc+1, _dafae)
				if _cageg != nil {
					return _cageg
				}
			}
		}
		return nil
	case *PdfObjectArray:
		_dece := _fggdf
		_ae.Log.Trace("-\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _dece)
		for _fcded, _afga := range _dece.Elements() {
			if _gbdf, _fdgf := _afga.(*PdfObjectReference); _fdgf {
				_cbcfe := _gbdf.Resolve()
				_dece.Set(_fcded, _cbcfe)
				_cgaff := _bbgd(_cbcfe, _aggcc+1, _dafae)
				if _cgaff != nil {
					return _cgaff
				}
			} else {
				_gbgf := _bbgd(_afga, _aggcc+1, _dafae)
				if _gbgf != nil {
					return _gbgf
				}
			}
		}
		return nil
	case *PdfObjectReference:
		_ae.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020T\u0072\u0061\u0063\u0069\u006e\u0067\u0020\u0061\u0020r\u0065\u0066\u0065r\u0065n\u0063\u0065\u0021")
		return _d.New("\u0065r\u0072\u006f\u0072\u0020t\u0072\u0061\u0063\u0069\u006eg\u0020a\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	return nil
}

// EncodeBytes encodes slice of bytes into JBIG2 encoding format.
// The input 'data' must be an image. In order to Decode it a user is responsible to
// load the codec ('png', 'jpg').
// Returns jbig2 single page encoded document byte slice. The encoder uses DefaultPageSettings
// to encode given image.
func (_fdce *JBIG2Encoder) EncodeBytes(data []byte) ([]byte, error) {
	const _gfaf = "\u004aB\u0049\u0047\u0032\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002eE\u006e\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065\u0073"
	if _fdce.ColorComponents != 1 || _fdce.BitsPerComponent != 1 {
		return nil, _eb.Errorf(_gfaf, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u002e\u0020\u004a\u0042\u0049G\u0032\u0020E\u006e\u0063o\u0064\u0065\u0072\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020bi\u006e\u0061\u0072\u0079\u0020\u0069\u006d\u0061\u0067e\u0073\u0020\u0064\u0061\u0074\u0061")
	}
	var (
		_ddgde *_gg.Bitmap
		_debf  error
	)
	_gbfg := (_fdce.Width * _fdce.Height) == len(data)
	if _gbfg {
		_ddgde, _debf = _gg.NewWithUnpaddedData(_fdce.Width, _fdce.Height, data)
	} else {
		_ddgde, _debf = _gg.NewWithData(_fdce.Width, _fdce.Height, data)
	}
	if _debf != nil {
		return nil, _debf
	}
	_fcda := _fdce.DefaultPageSettings
	if _debf = _fcda.Validate(); _debf != nil {
		return nil, _eb.Wrap(_debf, _gfaf, "")
	}
	if _fdce._ebac == nil {
		_fdce._ebac = _da.InitEncodeDocument(_fcda.FileMode)
	}
	switch _fcda.Compression {
	case JB2Generic:
		if _debf = _fdce._ebac.AddGenericPage(_ddgde, _fcda.DuplicatedLinesRemoval); _debf != nil {
			return nil, _eb.Wrap(_debf, _gfaf, "")
		}
	case JB2SymbolCorrelation:
		return nil, _eb.Error(_gfaf, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return nil, _eb.Error(_gfaf, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return nil, _eb.Error(_gfaf, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return _fdce.Encode()
}

// LZWEncoder provides LZW encoding/decoding functionality.
type LZWEncoder struct {
	Predictor        int
	BitsPerComponent int

	// For predictors
	Columns int
	Colors  int

	// LZW algorithm setting.
	EarlyChange int
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_fgafb *CCITTFaxEncoder) MakeDecodeParams() PdfObject {
	_fefb := MakeDict()
	_fefb.Set("\u004b", MakeInteger(int64(_fgafb.K)))
	_fefb.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_fgafb.Columns)))
	if _fgafb.BlackIs1 {
		_fefb.Set("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031", MakeBool(_fgafb.BlackIs1))
	}
	if _fgafb.EncodedByteAlign {
		_fefb.Set("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e", MakeBool(_fgafb.EncodedByteAlign))
	}
	if _fgafb.EndOfLine && _fgafb.K >= 0 {
		_fefb.Set("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee", MakeBool(_fgafb.EndOfLine))
	}
	if _fgafb.Rows != 0 && !_fgafb.EndOfBlock {
		_fefb.Set("\u0052\u006f\u0077\u0073", MakeInteger(int64(_fgafb.Rows)))
	}
	if !_fgafb.EndOfBlock {
		_fefb.Set("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b", MakeBool(_fgafb.EndOfBlock))
	}
	if _fgafb.DamagedRowsBeforeError != 0 {
		_fefb.Set("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072", MakeInteger(int64(_fgafb.DamagedRowsBeforeError)))
	}
	return _fefb
}

// StreamEncoder represents the interface for all PDF stream encoders.
type StreamEncoder interface {
	GetFilterName() string
	MakeDecodeParams() PdfObject
	MakeStreamDict() *PdfObjectDictionary
	UpdateParams(_bbce *PdfObjectDictionary)
	EncodeBytes(_cgg []byte) ([]byte, error)
	DecodeBytes(_aab []byte) ([]byte, error)
	DecodeStream(_eefc *PdfObjectStream) ([]byte, error)
}
type objectStreams map[int]objectStream
type limitedReadSeeker struct {
	_cgfa _dgf.ReadSeeker
	_gfcg int64
}

// WriteString outputs the object as it is to be written to file.
func (_edce *PdfObjectDictionary) WriteString() string {
	var _fbgb _agg.Builder
	_fbgb.WriteString("\u003c\u003c")
	for _, _adag := range _edce._dgcd {
		_faff := _edce._gged[_adag]
		_fbgb.WriteString(_adag.WriteString())
		_fbgb.WriteString("\u0020")
		_fbgb.WriteString(_faff.WriteString())
	}
	_fbgb.WriteString("\u003e\u003e")
	return _fbgb.String()
}

// XrefTable represents the cross references in a PDF, i.e. the table of objects and information
// where to access within the PDF file.
type XrefTable struct {
	ObjectMap map[int]XrefObject
	_bd       []XrefObject
}

func (_dabf *PdfObjectInteger) String() string { return _ac.Sprintf("\u0025\u0064", *_dabf) }

// EncodeBytes implements support for LZW encoding.  Currently not supporting predictors (raw compressed data only).
// Only supports the Early change = 1 algorithm (compress/lzw) as the other implementation
// does not have a write method.
// TODO: Consider refactoring compress/lzw to allow both.
func (_dcea *LZWEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _dcea.Predictor != 1 {
		return nil, _ac.Errorf("\u004c\u005aW \u0050\u0072\u0065d\u0069\u0063\u0074\u006fr =\u00201 \u006f\u006e\u006c\u0079\u0020\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0079e\u0074")
	}
	if _dcea.EarlyChange == 1 {
		return nil, _ac.Errorf("\u004c\u005a\u0057\u0020\u0045\u0061\u0072\u006c\u0079\u0020\u0043\u0068\u0061n\u0067\u0065\u0020\u003d\u0020\u0030 \u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0079\u0065\u0074")
	}
	var _egcd _fd.Buffer
	_aabd := _a.NewWriter(&_egcd, _a.MSB, 8)
	_aabd.Write(data)
	_aabd.Close()
	return _egcd.Bytes(), nil
}

// MakeArrayFromIntegers creates an PdfObjectArray from a slice of ints, where each array element is
// an PdfObjectInteger.
func MakeArrayFromIntegers(vals []int) *PdfObjectArray {
	_afbg := MakeArray()
	for _, _cbbf := range vals {
		_afbg.Append(MakeInteger(int64(_cbbf)))
	}
	return _afbg
}

// GetBoolVal returns the bool value within a *PdObjectBool represented by an PdfObject interface directly or indirectly.
// If the PdfObject does not represent a bool value, a default value of false is returned (found = false also).
func GetBoolVal(obj PdfObject) (_eaaf bool, _fdcef bool) {
	_dbade, _fdcef := TraceToDirectObject(obj).(*PdfObjectBool)
	if _fdcef {
		return bool(*_dbade), true
	}
	return false, false
}
func _gdcg(_agdg, _ebga, _fdef uint8) uint8 {
	_cgcc := int(_fdef)
	_aabb := int(_ebga) - _cgcc
	_fdefg := int(_agdg) - _cgcc
	_cgcc = _bcc(_aabb + _fdefg)
	_aabb = _bcc(_aabb)
	_fdefg = _bcc(_fdefg)
	if _aabb <= _fdefg && _aabb <= _cgcc {
		return _agdg
	} else if _fdefg <= _cgcc {
		return _ebga
	}
	return _fdef
}

// HeaderPosition gets the file header position.
func (_dcgb ParserMetadata) HeaderPosition() int { return _dcgb._cfg }

// PdfObject is an interface which all primitive PDF objects must implement.
type PdfObject interface {

	// String outputs a string representation of the primitive (for debugging).
	String() string

	// WriteString outputs the PDF primitive as written to file as expected by the standard.
	// TODO(dennwc): it should return a byte slice, or accept a writer
	WriteString() string
}

// MakeArray creates an PdfObjectArray from a list of PdfObjects.
func MakeArray(objects ...PdfObject) *PdfObjectArray { return &PdfObjectArray{_fffab: objects} }
func (_dcac *PdfCrypt) makeKey(_bbc string, _ebg, _ggg uint32, _dgff []byte) ([]byte, error) {
	_bbb, _ddae := _dcac._agfd[_bbc]
	if !_ddae {
		return nil, _ac.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _bbc)
	}
	return _bbb.MakeKey(_ebg, _ggg, _dgff)
}

// HasNonConformantStream implements core.ParserMetadata.
func (_ded ParserMetadata) HasNonConformantStream() bool { return _ded._gbbe }

// ReadAtLeast reads at least n bytes into slice p.
// Returns the number of bytes read (should always be == n), and an error on failure.
func (_accb *PdfParser) ReadAtLeast(p []byte, n int) (int, error) {
	_feff := n
	_fdge := 0
	_ggcb := 0
	for _feff > 0 {
		_dbfd, _bfe := _accb._eecea.Read(p[_fdge:])
		if _bfe != nil {
			_ae.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _dbfd, _ggcb, _bfe.Error())
			return _fdge, _d.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_ggcb++
		_fdge += _dbfd
		_feff -= _dbfd
	}
	return _fdge, nil
}
func (_dbgb *PdfParser) parseNumber() (PdfObject, error) { return ParseNumber(_dbgb._eecea) }

// String returns a descriptive information string about the encryption method used.
func (_efe *PdfCrypt) String() string {
	if _efe == nil {
		return ""
	}
	_defg := _efe._ddc.Filter + "\u0020\u002d\u0020"
	if _efe._ddc.V == 0 {
		_defg += "\u0055\u006e\u0064\u006fcu\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0061\u006c\u0067\u006f\u0072\u0069\u0074h\u006d"
	} else if _efe._ddc.V == 1 {
		_defg += "\u0052\u0043\u0034:\u0020\u0034\u0030\u0020\u0062\u0069\u0074\u0073"
	} else if _efe._ddc.V == 2 {
		_defg += _ac.Sprintf("\u0052\u0043\u0034:\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073", _efe._ddc.Length)
	} else if _efe._ddc.V == 3 {
		_defg += "U\u006e\u0070\u0075\u0062li\u0073h\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d"
	} else if _efe._ddc.V >= 4 {
		_defg += _ac.Sprintf("\u0053\u0074r\u0065\u0061\u006d\u0020f\u0069\u006ct\u0065\u0072\u003a\u0020\u0025\u0073\u0020\u002d \u0053\u0074\u0072\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074\u0065r\u003a\u0020\u0025\u0073", _efe._bba, _efe._dee)
		_defg += "\u003b\u0020C\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0073\u003a"
		for _adffg, _cad := range _efe._agfd {
			_defg += _ac.Sprintf("\u0020\u002d\u0020\u0025\u0073\u003a\u0020\u0025\u0073 \u0028\u0025\u0064\u0029", _adffg, _cad.Name(), _cad.KeyLength())
		}
	}
	_fga := _efe.GetAccessPermissions()
	_defg += _ac.Sprintf("\u0020\u002d\u0020\u0025\u0023\u0076", _fga)
	return _defg
}

// GetStream returns the *PdfObjectStream represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetStream(obj PdfObject) (_fccb *PdfObjectStream, _ddcc bool) {
	obj = ResolveReference(obj)
	_fccb, _ddcc = obj.(*PdfObjectStream)
	return _fccb, _ddcc
}

// WriteString outputs the object as it is to be written to file.
func (_abab *PdfObjectReference) WriteString() string {
	var _aeae _agg.Builder
	_aeae.WriteString(_dg.FormatInt(_abab.ObjectNumber, 10))
	_aeae.WriteString("\u0020")
	_aeae.WriteString(_dg.FormatInt(_abab.GenerationNumber, 10))
	_aeae.WriteString("\u0020\u0052")
	return _aeae.String()
}

// String returns a string describing `null`.
func (_gebbe *PdfObjectNull) String() string { return "\u006e\u0075\u006c\u006c" }

// GetStringVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetStringVal(obj PdfObject) (_baaec string, _cgcda bool) {
	_fcde, _cgcda := TraceToDirectObject(obj).(*PdfObjectString)
	if _cgcda {
		return _fcde.Str(), true
	}
	return
}

// Set sets the PdfObject at index i of the array. An error is returned if the index is outside bounds.
func (_acbf *PdfObjectArray) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_acbf._fffab) {
		return _d.New("\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_acbf._fffab[i] = obj
	return nil
}

// MakeName creates a PdfObjectName from a string.
func MakeName(s string) *PdfObjectName { _aecee := PdfObjectName(s); return &_aecee }

// GetPreviousRevisionReadSeeker returns ReadSeeker for the previous version of the Pdf document.
func (_gbfc *PdfParser) GetPreviousRevisionReadSeeker() (_dgf.ReadSeeker, error) {
	if _bafd := _gbfc.seekToEOFMarker(_gbfc._fcca - _fba); _bafd != nil {
		return nil, _bafd
	}
	_cagc, _abagc := _gbfc._fdee.Seek(0, _dgf.SeekCurrent)
	if _abagc != nil {
		return nil, _abagc
	}
	_cagc += _fba
	return _afef(_gbfc._fdee, _cagc)
}
func (_dce *PdfCrypt) securityHandler() _geg.StdHandler {
	if _dce._gge.R >= 5 {
		return _geg.NewHandlerR6()
	}
	return _geg.NewHandlerR4(_dce._dfc, _dce._ddc.Length)
}

// EncodeImage encodes 'img' golang image.Image into jbig2 encoded bytes document using default encoder settings.
func (_fgfg *JBIG2Encoder) EncodeImage(img _ea.Image) ([]byte, error) { return _fgfg.encodeImage(img) }

// Str returns the string value of the PdfObjectString. Defined in addition to String() function to clarify that
// this function returns the underlying string directly, whereas the String function technically could include
// debug info.
func (_cfbfg *PdfObjectString) Str() string { return _cfbfg._eeee }

// GetNumberAsInt64 returns the contents of `obj` as an int64 if it is an integer or float, or an
// error if it isn't. This is for cases where expecting an integer, but some implementations
// actually store the number in a floating point format.
func GetNumberAsInt64(obj PdfObject) (int64, error) {
	_babaf, _dgbc := obj.(*PdfObjectReference)
	if _dgbc {
		obj = TraceToDirectObject(_babaf)
	} else if _fcaf, _egfb := obj.(*PdfIndirectObject); _egfb {
		obj = _fcaf.PdfObject
	}
	switch _acbb := obj.(type) {
	case *PdfObjectFloat:
		_ae.Log.Debug("\u004e\u0075m\u0062\u0065\u0072\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u0073\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0077\u0061s\u0020\u0073\u0074\u006f\u0072\u0065\u0064\u0020\u0061\u0073\u0020\u0066\u006c\u006fa\u0074\u0020(\u0074\u0079\u0070\u0065 \u0063\u0061\u0073\u0074\u0069n\u0067\u0020\u0075\u0073\u0065\u0064\u0029")
		return int64(*_acbb), nil
	case *PdfObjectInteger:
		return int64(*_acbb), nil
	}
	return 0, ErrNotANumber
}

// ResolveReference resolves reference if `o` is a *PdfObjectReference and returns the object referenced to.
// Otherwise returns back `o`.
func ResolveReference(obj PdfObject) PdfObject {
	if _fcgab, _dcefc := obj.(*PdfObjectReference); _dcefc {
		return _fcgab.Resolve()
	}
	return obj
}

// Clear resets the dictionary to an empty state.
func (_dedc *PdfObjectDictionary) Clear() {
	_dedc._dgcd = []PdfObjectName{}
	_dedc._gged = map[PdfObjectName]PdfObject{}
	_dedc._efce = &_e.Mutex{}
}

// Keys returns the list of keys in the dictionary.
// If `d` is nil returns a nil slice.
func (_cfdfc *PdfObjectDictionary) Keys() []PdfObjectName {
	if _cfdfc == nil {
		return nil
	}
	return _cfdfc._dgcd
}
func (_cdgc *PdfParser) readComment() (string, error) {
	var _dceag _fd.Buffer
	_, _bedcf := _cdgc.skipSpaces()
	if _bedcf != nil {
		return _dceag.String(), _bedcf
	}
	_aaaf := true
	for {
		_fbgc, _adaf := _cdgc._eecea.Peek(1)
		if _adaf != nil {
			_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _adaf.Error())
			return _dceag.String(), _adaf
		}
		if _aaaf && _fbgc[0] != '%' {
			return _dceag.String(), _d.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_aaaf = false
		if (_fbgc[0] != '\r') && (_fbgc[0] != '\n') {
			_dfeg, _ := _cdgc._eecea.ReadByte()
			_dceag.WriteByte(_dfeg)
		} else {
			break
		}
	}
	return _dceag.String(), nil
}

// JBIG2Image is the image structure used by the jbig2 encoder. Its Data must be in a
// 1 bit per component and 1 component per pixel (1bpp). In order to create binary image
// use GoImageToJBIG2 function. If the image data contains the row bytes padding set the HasPadding to true.
type JBIG2Image struct {

	// Width and Height defines the image boundaries.
	Width, Height int

	// Data is the byte slice data for the input image
	Data []byte

	// HasPadding is the attribute that defines if the last byte of the data in the row contains
	// 0 bits padding.
	HasPadding bool
}

// IsFloatDigit checks if a character can be a part of a float number string.
func IsFloatDigit(c byte) bool { return ('0' <= c && c <= '9') || c == '.' }

// GoImageToJBIG2 creates a binary image on the base of 'i' golang image.Image.
// If the image is not a black/white image then the function converts provided input into
// JBIG2Image with 1bpp. For non grayscale images the function performs the conversion to the grayscale temp image.
// Then it checks the value of the gray image value if it's within bounds of the black white threshold.
// This 'bwThreshold' value should be in range (0.0, 1.0). The threshold checks if the grayscale pixel (uint) value
// is greater or smaller than 'bwThreshold' * 255. Pixels inside the range will be white, and the others will be black.
// If the 'bwThreshold' is equal to -1.0 - JB2ImageAutoThreshold then it's value would be set on the base of
// it's histogram using Triangle method. For more information go to:
// 	https://www.mathworks.com/matlabcentral/fileexchange/28047-gray-image-thresholding-using-the-triangle-method
func GoImageToJBIG2(i _ea.Image, bwThreshold float64) (*JBIG2Image, error) {
	const _ggdc = "\u0047\u006f\u0049\u006d\u0061\u0067\u0065\u0054\u006fJ\u0042\u0049\u0047\u0032"
	if i == nil {
		return nil, _eb.Error(_ggdc, "i\u006d\u0061\u0067\u0065 '\u0069'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	var (
		_ggaa uint8
		_fbff _ce.Image
		_ffac error
	)
	if bwThreshold == JB2ImageAutoThreshold {
		_fbff, _ffac = _ce.MonochromeConverter.Convert(i)
	} else if bwThreshold > 1.0 || bwThreshold < 0.0 {
		return nil, _eb.Error(_ggdc, "p\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074h\u0072\u0065\u0073\u0068\u006f\u006c\u0064 i\u0073\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u0020\u0072\u0061\u006ege\u0020\u007b0\u002e\u0030\u002c\u0020\u0031\u002e\u0030\u007d")
	} else {
		_ggaa = uint8(255 * bwThreshold)
		_fbff, _ffac = _ce.MonochromeThresholdConverter(_ggaa).Convert(i)
	}
	if _ffac != nil {
		return nil, _ffac
	}
	return _ecca(_fbff), nil
}
func (_bcgff *PdfObjectDictionary) setWithLock(_ebbd PdfObjectName, _aeaf PdfObject, _gebe bool) {
	if _gebe {
		_bcgff._efce.Lock()
		defer _bcgff._efce.Unlock()
	}
	_, _ccgc := _bcgff._gged[_ebbd]
	if !_ccgc {
		_bcgff._dgcd = append(_bcgff._dgcd, _ebbd)
	}
	_bcgff._gged[_ebbd] = _aeaf
}

// Version represents a version of a PDF standard.
type Version struct {
	Major int
	Minor int
}

// JBIG2Encoder implements both jbig2 encoder and the decoder. The encoder allows to encode
// provided images (best used document scans) in multiple way. By default it uses single page generic
// encoder. It allows to store lossless data as a single segment.
// In order to store multiple image pages use the 'FileMode' which allows to store more pages within single jbig2 document.
// WIP: In order to obtain better compression results the encoder would allow to encode the input in a
// lossy or lossless way with a component (symbol) mode. It divides the image into components.
// Then checks if any component is 'similar' to the others and maps them together. The symbol classes are stored
// in the dictionary. Then the encoder creates text regions which uses the related symbol classes to fill it's space.
// The similarity is defined by the 'Threshold' variable (default: 0.95). The less the value is, the more components
// matches to single class, thus the compression is better, but the result might become lossy.
type JBIG2Encoder struct {

	// These values are required to be set for the 'EncodeBytes' method.
	// ColorComponents defines the number of color components for provided image.
	ColorComponents int

	// BitsPerComponent is the number of bits that stores per color component
	BitsPerComponent int

	// Width is the width of the image to encode
	Width int

	// Height is the height of the image to encode.
	Height int
	_ebac  *_da.Document

	// Globals are the JBIG2 global segments.
	Globals _df.Globals

	// IsChocolateData defines if the data is encoded such that
	// binary data '1' means black and '0' white.
	// otherwise the data is called vanilla.
	// Naming convention taken from: 'https://en.wikipedia.org/wiki/Binary_image#Interpretation'
	IsChocolateData bool

	// DefaultPageSettings are the settings parameters used by the jbig2 encoder.
	DefaultPageSettings JBIG2EncoderSettings
}

// ToIntegerArray returns a slice of all array elements as an int slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_acddc *PdfObjectArray) ToIntegerArray() ([]int, error) {
	var _gffba []int
	for _, _cfcd := range _acddc.Elements() {
		if _cdcfd, _ceec := _cfcd.(*PdfObjectInteger); _ceec {
			_gffba = append(_gffba, int(*_cdcfd))
		} else {
			return nil, ErrTypeError
		}
	}
	return _gffba, nil
}

// GetStringBytes is like GetStringVal except that it returns the string as a []byte.
// It is for convenience.
func GetStringBytes(obj PdfObject) (_gfae []byte, _ecef bool) {
	_gdde, _ecef := TraceToDirectObject(obj).(*PdfObjectString)
	if _ecef {
		return _gdde.Bytes(), true
	}
	return
}

// JBIG2CompressionType defines the enum compression type used by the JBIG2Encoder.
type JBIG2CompressionType int

// GetInt returns the *PdfObjectBool object that is represented by a PdfObject either directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetInt(obj PdfObject) (_facd *PdfObjectInteger, _edfg bool) {
	_facd, _edfg = TraceToDirectObject(obj).(*PdfObjectInteger)
	return _facd, _edfg
}

// PdfCryptNewDecrypt makes the document crypt handler based on the encryption dictionary
// and trailer dictionary. Returns an error on failure to process.
func PdfCryptNewDecrypt(parser *PdfParser, ed, trailer *PdfObjectDictionary) (*PdfCrypt, error) {
	_bgbg := &PdfCrypt{_gfg: false, _dgd: make(map[PdfObject]bool), _egg: make(map[PdfObject]bool), _dfg: make(map[int]struct{}), _cdg: parser}
	_egga, _afeg := ed.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectName)
	if !_afeg {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0043\u0072\u0079\u0070\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u0066\u0069\u0065\u006c\u0064\u0021")
		return _bgbg, _d.New("r\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0063\u0072\u0079\u0070\u0074 \u0066\u0069\u0065\u006c\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
	}
	if *_egga != "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064" {
		_ae.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020(%\u0073\u0029", *_egga)
		return _bgbg, _d.New("\u0075n\u0073u\u0070\u0070\u006f\u0072\u0074e\u0064\u0020F\u0069\u006c\u0074\u0065\u0072")
	}
	_bgbg._ddc.Filter = string(*_egga)
	if _cabf, _cfcc := ed.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r").(*PdfObjectString); _cfcc {
		_bgbg._ddc.SubFilter = _cabf.Str()
		_ae.Log.Debug("\u0055s\u0069n\u0067\u0020\u0073\u0075\u0062f\u0069\u006ct\u0065\u0072\u0020\u0025\u0073", _cabf)
	}
	if L, _cgcf := ed.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _cgcf {
		if (*L % 8) != 0 {
			_ae.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0065\u006ec\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
			return _bgbg, _d.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0065\u006e\u0063\u0072y\u0070t\u0069o\u006e\u0020\u006c\u0065\u006e\u0067\u0074h")
		}
		_bgbg._ddc.Length = int(*L)
	} else {
		_bgbg._ddc.Length = 40
	}
	_bgbg._ddc.V = 0
	if _abe, _ff := ed.Get("\u0056").(*PdfObjectInteger); _ff {
		V := int(*_abe)
		_bgbg._ddc.V = V
		if V >= 1 && V <= 2 {
			_bgbg._agfd = _gcba(_bgbg._ddc.Length)
		} else if V >= 4 && V <= 5 {
			if _dde := _bgbg.loadCryptFilters(ed); _dde != nil {
				return _bgbg, _dde
			}
		} else {
			_ae.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0061lg\u006f\u0020\u0056 \u003d \u0025\u0064", V)
			return _bgbg, _d.New("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d")
		}
	}
	if _gee := _gba(&_bgbg._gge, ed); _gee != nil {
		return _bgbg, _gee
	}
	_gbae := ""
	if _gbcd, _eee := trailer.Get("\u0049\u0044").(*PdfObjectArray); _eee && _gbcd.Len() >= 1 {
		_efc, _dcb := GetString(_gbcd.Get(0))
		if !_dcb {
			return _bgbg, _d.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0074r\u0061\u0069l\u0065\u0072\u0020\u0049\u0044")
		}
		_gbae = _efc.Str()
	} else {
		_ae.Log.Debug("\u0054\u0072ai\u006c\u0065\u0072 \u0049\u0044\u0020\u0061rra\u0079 m\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0021")
	}
	_bgbg._dfc = _gbae
	return _bgbg, nil
}

// NewCCITTFaxEncoder makes a new CCITTFax encoder.
func NewCCITTFaxEncoder() *CCITTFaxEncoder { return &CCITTFaxEncoder{Columns: 1728, EndOfBlock: true} }

// String returns the state of the bool as "true" or "false".
func (_feafe *PdfObjectBool) String() string {
	if *_feafe {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}

// UpdateParams updates the parameter values of the encoder.
func (_ccfg *RawEncoder) UpdateParams(params *PdfObjectDictionary) {}

// PdfObjectReference represents the primitive PDF reference object.
type PdfObjectReference struct {
	_ffgd            *PdfParser
	ObjectNumber     int64
	GenerationNumber int64
}

func (_cdgd *PdfParser) readTextLine() (string, error) {
	var _cfba _fd.Buffer
	for {
		_aebbc, _cccee := _cdgd._eecea.Peek(1)
		if _cccee != nil {
			_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _cccee.Error())
			return _cfba.String(), _cccee
		}
		if (_aebbc[0] != '\r') && (_aebbc[0] != '\n') {
			_abfb, _ := _cdgd._eecea.ReadByte()
			_cfba.WriteByte(_abfb)
		} else {
			break
		}
	}
	return _cfba.String(), nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_bbecc *FlateEncoder) MakeDecodeParams() PdfObject {
	if _bbecc.Predictor > 1 {
		_ggeb := MakeDict()
		_ggeb.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_bbecc.Predictor)))
		if _bbecc.BitsPerComponent != 8 {
			_ggeb.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_bbecc.BitsPerComponent)))
		}
		if _bbecc.Columns != 1 {
			_ggeb.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_bbecc.Columns)))
		}
		if _bbecc.Colors != 1 {
			_ggeb.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_bbecc.Colors)))
		}
		return _ggeb
	}
	return nil
}

// String returns the PDF version as a string. Implements interface fmt.Stringer.
func (_edaca Version) String() string {
	return _ac.Sprintf("\u00250\u0064\u002e\u0025\u0030\u0064", _edaca.Major, _edaca.Minor)
}

// WriteString outputs the object as it is to be written to file.
func (_efdda *PdfObjectName) WriteString() string {
	var _aeceg _fd.Buffer
	if len(*_efdda) > 127 {
		_ae.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u0061\u006d\u0065\u0020t\u006fo\u0020l\u006f\u006e\u0067\u0020\u0028\u0025\u0073)", *_efdda)
	}
	_aeceg.WriteString("\u002f")
	for _ecac := 0; _ecac < len(*_efdda); _ecac++ {
		_cbgf := (*_efdda)[_ecac]
		if !IsPrintable(_cbgf) || _cbgf == '#' || IsDelimiter(_cbgf) {
			_aeceg.WriteString(_ac.Sprintf("\u0023\u0025\u002e2\u0078", _cbgf))
		} else {
			_aeceg.WriteByte(_cbgf)
		}
	}
	return _aeceg.String()
}
func _bdff(_fbcdc, _dfae PdfObject, _efaac int) bool {
	if _efaac > _dgag {
		_ae.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _dgag)
		return false
	}
	if _fbcdc == nil && _dfae == nil {
		return true
	} else if _fbcdc == nil || _dfae == nil {
		return false
	}
	if _c.TypeOf(_fbcdc) != _c.TypeOf(_dfae) {
		return false
	}
	switch _afbfa := _fbcdc.(type) {
	case *PdfObjectNull, *PdfObjectReference:
		return true
	case *PdfObjectName:
		return *_afbfa == *(_dfae.(*PdfObjectName))
	case *PdfObjectString:
		return *_afbfa == *(_dfae.(*PdfObjectString))
	case *PdfObjectInteger:
		return *_afbfa == *(_dfae.(*PdfObjectInteger))
	case *PdfObjectBool:
		return *_afbfa == *(_dfae.(*PdfObjectBool))
	case *PdfObjectFloat:
		return *_afbfa == *(_dfae.(*PdfObjectFloat))
	case *PdfIndirectObject:
		return _bdff(TraceToDirectObject(_fbcdc), TraceToDirectObject(_dfae), _efaac+1)
	case *PdfObjectArray:
		_cebe := _dfae.(*PdfObjectArray)
		if len((*_afbfa)._fffab) != len((*_cebe)._fffab) {
			return false
		}
		for _dgfdg, _gfgc := range (*_afbfa)._fffab {
			if !_bdff(_gfgc, (*_cebe)._fffab[_dgfdg], _efaac+1) {
				return false
			}
		}
		return true
	case *PdfObjectDictionary:
		_fbfcf := _dfae.(*PdfObjectDictionary)
		_gedc, _bcda := (*_afbfa)._gged, (*_fbfcf)._gged
		if len(_gedc) != len(_bcda) {
			return false
		}
		for _cacd, _cgcb := range _gedc {
			_aebbcg, _cbee := _bcda[_cacd]
			if !_cbee || !_bdff(_cgcb, _aebbcg, _efaac+1) {
				return false
			}
		}
		return true
	case *PdfObjectStream:
		_dfff := _dfae.(*PdfObjectStream)
		return _bdff((*_afbfa).PdfObjectDictionary, (*_dfff).PdfObjectDictionary, _efaac+1)
	default:
		_ae.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", _fbcdc)
	}
	return false
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_gcea *MultiEncoder) MakeStreamDict() *PdfObjectDictionary {
	_facac := MakeDict()
	_facac.Set("\u0046\u0069\u006c\u0074\u0065\u0072", _gcea.GetFilterArray())
	for _, _gbef := range _gcea._bfacg {
		_dbfb := _gbef.MakeStreamDict()
		for _, _dgcf := range _dbfb.Keys() {
			_cged := _dbfb.Get(_dgcf)
			if _dgcf != "\u0046\u0069\u006c\u0074\u0065\u0072" && _dgcf != "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073" {
				_facac.Set(_dgcf, _cged)
			}
		}
	}
	_bcab := _gcea.MakeDecodeParams()
	if _bcab != nil {
		_facac.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _bcab)
	}
	return _facac
}

// DCTEncoder provides a DCT (JPG) encoding/decoding functionality for images.
type DCTEncoder struct {
	ColorComponents  int
	BitsPerComponent int
	Width            int
	Height           int
	Quality          int
}

// UpdateParams updates the parameter values of the encoder.
func (_bgcd *MultiEncoder) UpdateParams(params *PdfObjectDictionary) {
	for _, _egec := range _bgcd._bfacg {
		_egec.UpdateParams(params)
	}
}

// ASCIIHexEncoder implements ASCII hex encoder/decoder.
type ASCIIHexEncoder struct{}

// ParserMetadata is the parser based metadata information about document.
// The data here could be used on document verification.
type ParserMetadata struct {
	_cfg   int
	_egad  bool
	_fcbg  [4]byte
	_fdgg  bool
	_gbcdc bool
	_bbec  bool
	_gbbe  bool
	_bcg   bool
	_bega  bool
}

// SetImage sets the image base for given flate encoder.
func (_egdf *FlateEncoder) SetImage(img *_ce.ImageBase) { _egdf._dga = img }

// UpdateParams updates the parameter values of the encoder.
func (_ggbf *FlateEncoder) UpdateParams(params *PdfObjectDictionary) {
	_cbeb, _dcab := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _dcab == nil {
		_ggbf.Predictor = int(_cbeb)
	}
	_fad, _dcab := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _dcab == nil {
		_ggbf.BitsPerComponent = int(_fad)
	}
	_cddb, _dcab := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _dcab == nil {
		_ggbf.Columns = int(_cddb)
	}
	_eba, _dcab := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _dcab == nil {
		_ggbf.Colors = int(_eba)
	}
}
func _efdg() string { return _ae.Version }

// NewJPXEncoder returns a new instance of JPXEncoder.
func NewJPXEncoder() *JPXEncoder { return &JPXEncoder{} }
func (_abba *PdfParser) parseDetailedHeader() (_agda error) {
	_abba._fdee.Seek(0, _dgf.SeekStart)
	_abba._eecea = _acg.NewReader(_abba._fdee)
	_dcfa := 20
	_aced := make([]byte, _dcfa)
	var (
		_bdc  bool
		_ceab int
	)
	for {
		_ebcg, _gbf := _abba._eecea.ReadByte()
		if _gbf != nil {
			if _gbf == _dgf.EOF {
				break
			} else {
				return _gbf
			}
		}
		if IsDecimalDigit(_ebcg) && _aced[_dcfa-1] == '.' && IsDecimalDigit(_aced[_dcfa-2]) && _aced[_dcfa-3] == '-' && _aced[_dcfa-4] == 'F' && _aced[_dcfa-5] == 'D' && _aced[_dcfa-6] == 'P' && _aced[_dcfa-7] == '%' {
			_abba._fadgc = Version{Major: int(_aced[_dcfa-2] - '0'), Minor: int(_ebcg - '0')}
			_abba._aecec._cfg = _ceab - 7
			_bdc = true
			break
		}
		_ceab++
		_aced = append(_aced[1:_dcfa], _ebcg)
	}
	if !_bdc {
		return _ac.Errorf("n\u006f \u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066ou\u006e\u0064")
	}
	_cbea, _agda := _abba._eecea.ReadByte()
	if _agda == _dgf.EOF {
		return _ac.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0050d\u0066\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _agda != nil {
		return _agda
	}
	_abba._aecec._egad = _cbea == '\n'
	_cbea, _agda = _abba._eecea.ReadByte()
	if _agda != nil {
		return _ac.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _agda)
	}
	if _cbea != '%' {
		return nil
	}
	_afc := make([]byte, 4)
	_, _agda = _abba._eecea.Read(_afc)
	if _agda != nil {
		return _ac.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _agda)
	}
	_abba._aecec._fcbg = [4]byte{_afc[0], _afc[1], _afc[2], _afc[3]}
	return nil
}
func (_egff *PdfCrypt) generateParams(_cbg, _aac []byte) error {
	_bfbf := _egff.securityHandler()
	_afa, _ffe := _bfbf.GenerateParams(&_egff._gge, _aac, _cbg)
	if _ffe != nil {
		return _ffe
	}
	_egff._dadd = _afa
	return nil
}

// String returns a string describing `ref`.
func (_bacg *PdfObjectReference) String() string {
	return _ac.Sprintf("\u0052\u0065\u0066\u0028\u0025\u0064\u0020\u0025\u0064\u0029", _bacg.ObjectNumber, _bacg.GenerationNumber)
}

// PdfParser parses a PDF file and provides access to the object structure of the PDF.
type PdfParser struct {
	_fadgc   Version
	_fdee    _dgf.ReadSeeker
	_eecea   *_acg.Reader
	_fcca    int64
	_fbab    XrefTable
	_egfag   int64
	_dfcgd   *xrefType
	_acab    objectStreams
	_dfec    *PdfObjectDictionary
	_bffd    *PdfCrypt
	_bcdd    *PdfIndirectObject
	_bgcf    bool
	ObjCache objectCache
	_ccge    map[int]bool
	_gfdf    map[int64]bool
	_aecec   ParserMetadata
	_ccce    bool
	_edcac   []int64
	_fdbdc   int
	_gfggc   bool
	_dada    int64
	_bgcdd   map[*PdfParser]*PdfParser
	_ddcg    []*PdfParser
}

// DecodeStream decodes the stream containing CCITTFax encoded image data.
func (_ebfbb *CCITTFaxEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _ebfbb.DecodeBytes(streamObj.Stream)
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_cgdg *ASCIIHexEncoder) MakeStreamDict() *PdfObjectDictionary {
	_geae := MakeDict()
	_geae.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_cgdg.GetFilterName()))
	return _geae
}
func (_ebea *PdfParser) checkLinearizedInformation(_gbab *PdfObjectDictionary) (bool, error) {
	var _geegd error
	_ebea._dada, _geegd = GetNumberAsInt64(_gbab.Get("\u004c"))
	if _geegd != nil {
		return false, _geegd
	}
	_geegd = _ebea.seekToEOFMarker(_ebea._dada)
	switch _geegd {
	case nil:
		return true, nil
	case _ecee:
		return false, nil
	default:
		return false, _geegd
	}
}

// FlattenObject returns the contents of `obj`. In other words, `obj` with indirect objects replaced
// by their values.
// The replacements are made recursively to a depth of traceMaxDepth.
// NOTE: Dicts are sorted to make objects with same contents have the same PDF object strings.
func FlattenObject(obj PdfObject) PdfObject { return _cbbg(obj, 0) }

// GetPreviousRevisionParser returns PdfParser for the previous version of the Pdf document.
func (_fbeg *PdfParser) GetPreviousRevisionParser() (*PdfParser, error) {
	if _fbeg._fdbdc == 0 {
		return nil, _d.New("\u0074\u0068\u0069\u0073 i\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0072\u0065\u0076\u0069\u0073\u0069o\u006e")
	}
	if _bgfd, _eegd := _fbeg._bgcdd[_fbeg]; _eegd {
		return _bgfd, nil
	}
	_bdaf, _bbbf := _fbeg.GetPreviousRevisionReadSeeker()
	if _bbbf != nil {
		return nil, _bbbf
	}
	_adgb, _bbbf := NewParser(_bdaf)
	_adgb._bgcdd = _fbeg._bgcdd
	if _bbbf != nil {
		return nil, _bbbf
	}
	_fbeg._bgcdd[_fbeg] = _adgb
	return _adgb, nil
}

// EqualObjects returns true if `obj1` and `obj2` have the same contents.
//
// NOTE: It is a good idea to flatten obj1 and obj2 with FlattenObject before calling this function
// so that contents, rather than references, can be compared.
func EqualObjects(obj1, obj2 PdfObject) bool { return _bdff(obj1, obj2, 0) }

// UpdateParams updates the parameter values of the encoder.
func (_fedb *CCITTFaxEncoder) UpdateParams(params *PdfObjectDictionary) {
	if _abgef, _caac := GetNumberAsInt64(params.Get("\u004b")); _caac == nil {
		_fedb.K = int(_abgef)
	}
	if _dcag, _eeff := GetNumberAsInt64(params.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _eeff == nil {
		_fedb.Columns = int(_dcag)
	} else if _dcag, _eeff = GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068")); _eeff == nil {
		_fedb.Columns = int(_dcag)
	}
	if _afdea, _bdg := GetNumberAsInt64(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _bdg == nil {
		_fedb.BlackIs1 = _afdea > 0
	} else {
		if _dddg, _ebfga := GetBoolVal(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _ebfga {
			_fedb.BlackIs1 = _dddg
		} else {
			if _cfbd, _fede := GetArray(params.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _fede {
				_ebefe, _gbg := _cfbd.ToIntegerArray()
				if _gbg == nil {
					_fedb.BlackIs1 = _ebefe[0] == 1 && _ebefe[1] == 0
				}
			}
		}
	}
	if _dbg, _aebd := GetNumberAsInt64(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _aebd == nil {
		_fedb.EncodedByteAlign = _dbg > 0
	} else {
		if _cdda, _cfe := GetBoolVal(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _cfe {
			_fedb.EncodedByteAlign = _cdda
		}
	}
	if _fcag, _bged := GetNumberAsInt64(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _bged == nil {
		_fedb.EndOfLine = _fcag > 0
	} else {
		if _fgb, _aag := GetBoolVal(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _aag {
			_fedb.EndOfLine = _fgb
		}
	}
	if _cadf, _ageb := GetNumberAsInt64(params.Get("\u0052\u006f\u0077\u0073")); _ageb == nil {
		_fedb.Rows = int(_cadf)
	} else if _cadf, _ageb = GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _ageb == nil {
		_fedb.Rows = int(_cadf)
	}
	if _adc, _gaaa := GetNumberAsInt64(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _gaaa == nil {
		_fedb.EndOfBlock = _adc > 0
	} else {
		if _cdca, _begb := GetBoolVal(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _begb {
			_fedb.EndOfBlock = _cdca
		}
	}
	if _gcdg, _ebbc := GetNumberAsInt64(params.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _ebbc != nil {
		_fedb.DamagedRowsBeforeError = int(_gcdg)
	}
}
func _bfdd(_gegea _dgf.ReadSeeker, _dgdg int64) (*offsetReader, error) {
	_dccd := &offsetReader{_agbf: _gegea, _ceca: _dgdg}
	_, _gfbg := _dccd.Seek(0, _dgf.SeekStart)
	return _dccd, _gfbg
}

// JPXEncoder implements JPX encoder/decoder (dummy, for now)
// FIXME: implement
type JPXEncoder struct{}

// IsDelimiter checks if a character represents a delimiter.
func IsDelimiter(c byte) bool {
	return c == '(' || c == ')' || c == '<' || c == '>' || c == '[' || c == ']' || c == '{' || c == '}' || c == '/' || c == '%'
}

// EncodeStream encodes the stream data using the encoded specified by the stream's dictionary.
func EncodeStream(streamObj *PdfObjectStream) error {
	_ae.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_aeafa, _fbca := NewEncoderFromStream(streamObj)
	if _fbca != nil {
		_ae.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _fbca)
		return _fbca
	}
	if _bcaa, _acabf := _aeafa.(*LZWEncoder); _acabf {
		_bcaa.EarlyChange = 0
		streamObj.PdfObjectDictionary.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(0))
	}
	_ae.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u000a", _aeafa)
	_bgea, _fbca := _aeafa.EncodeBytes(streamObj.Stream)
	if _fbca != nil {
		_ae.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _fbca)
		return _fbca
	}
	streamObj.Stream = _bgea
	streamObj.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_bgea))))
	return nil
}

// MakeBool creates a PdfObjectBool from a bool value.
func MakeBool(val bool) *PdfObjectBool { _gbgec := PdfObjectBool(val); return &_gbgec }

var _aea = _ba.MustCompile("\u005e\\\u0073\u002a\u005b\u002d]\u002a\u0028\u005c\u0064\u002b)\u005cs\u002b(\u005c\u0064\u002b\u0029\u005c\u0073\u002bR")

// DecodeStream implements ASCII hex decoding.
func (_ddea *ASCIIHexEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _ddea.DecodeBytes(streamObj.Stream)
}
func (_febcc *PdfParser) repairSeekXrefMarker() error {
	_fadag, _ffgba := _febcc._fdee.Seek(0, _dgf.SeekEnd)
	if _ffgba != nil {
		return _ffgba
	}
	_ceffe := _ba.MustCompile("\u005cs\u0078\u0072\u0065\u0066\u005c\u0073*")
	var _cccc int64
	var _bacd int64 = 1000
	for _cccc < _fadag {
		if _fadag <= (_bacd + _cccc) {
			_bacd = _fadag - _cccc
		}
		_, _caaeb := _febcc._fdee.Seek(-_cccc-_bacd, _dgf.SeekEnd)
		if _caaeb != nil {
			return _caaeb
		}
		_fdbe := make([]byte, _bacd)
		_febcc._fdee.Read(_fdbe)
		_ae.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0078\u0072\u0065\u0066\u0020\u003a\u0020\u0022\u0025\u0073\u0022", string(_fdbe))
		_ebde := _ceffe.FindAllStringIndex(string(_fdbe), -1)
		if _ebde != nil {
			_gdfa := _ebde[len(_ebde)-1]
			_ae.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _ebde)
			_febcc._fdee.Seek(-_cccc-_bacd+int64(_gdfa[0]), _dgf.SeekEnd)
			_febcc._eecea = _acg.NewReader(_febcc._fdee)
			for {
				_ebcd, _cccb := _febcc._eecea.Peek(1)
				if _cccb != nil {
					return _cccb
				}
				_ae.Log.Trace("\u0042\u003a\u0020\u0025\u0064\u0020\u0025\u0063", _ebcd[0], _ebcd[0])
				if !IsWhiteSpace(_ebcd[0]) {
					break
				}
				_febcc._eecea.Discard(1)
			}
			return nil
		}
		_ae.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_cccc += _bacd
	}
	_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074a\u0062\u006c\u0065\u0020\u006d\u0061r\u006b\u0065\u0072\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u002e")
	return _d.New("\u0078r\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020")
}
func (_gag *PdfParser) lookupByNumberWrapper(_abg int, _dd bool) (PdfObject, bool, error) {
	_gc, _agb, _be := _gag.lookupByNumber(_abg, _dd)
	if _be != nil {
		return nil, _agb, _be
	}
	if !_agb && _gag._bffd != nil && _gag._bffd._gfg && !_gag._bffd.isDecrypted(_gc) {
		_dge := _gag._bffd.Decrypt(_gc, 0, 0)
		if _dge != nil {
			return nil, _agb, _dge
		}
	}
	return _gc, _agb, nil
}

// DecodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_edfd *RawEncoder) DecodeBytes(encoded []byte) ([]byte, error) { return encoded, nil }

const (
	_bgg  = 0
	_ebec = 1
	_ggf  = 2
	_ebca = 3
	_ffgg = 4
)

// GetRevisionNumber returns the current version of the Pdf document.
func (_bfcef *PdfParser) GetRevisionNumber() int { return _bfcef._fdbdc }
func (_dfce *PdfObjectFloat) String() string     { return _ac.Sprintf("\u0025\u0066", *_dfce) }

// WriteString outputs the object as it is to be written to file.
func (_efda *PdfIndirectObject) WriteString() string {
	var _edag _agg.Builder
	_edag.WriteString(_dg.FormatInt(_efda.ObjectNumber, 10))
	_edag.WriteString("\u0020\u0030\u0020\u0052")
	return _edag.String()
}

// GetAsFloat64Slice returns the array as []float64 slice.
// Returns an error if not entirely numeric (only PdfObjectIntegers, PdfObjectFloats).
func (_eade *PdfObjectArray) GetAsFloat64Slice() ([]float64, error) {
	var _bccac []float64
	for _, _cgbfd := range _eade.Elements() {
		_gbgb, _gabf := GetNumberAsFloat(TraceToDirectObject(_cgbfd))
		if _gabf != nil {
			return nil, _ac.Errorf("\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0075m\u0062\u0065\u0072")
		}
		_bccac = append(_bccac, _gbgb)
	}
	return _bccac, nil
}

// ResolveReferencesDeep recursively traverses through object `o`, looking up and replacing
// references with indirect objects.
// Optionally a map of already deep-resolved objects can be provided via `traversed`. The `traversed` map
// is updated while traversing the objects to avoid traversing same objects multiple times.
func ResolveReferencesDeep(o PdfObject, traversed map[PdfObject]struct{}) error {
	if traversed == nil {
		traversed = map[PdfObject]struct{}{}
	}
	return _bbgd(o, 0, traversed)
}
func _gcaf(_cabb PdfObject) (*float64, error) {
	switch _gabg := _cabb.(type) {
	case *PdfObjectFloat:
		_bggaa := float64(*_gabg)
		return &_bggaa, nil
	case *PdfObjectInteger:
		_bfag := float64(*_gabg)
		return &_bfag, nil
	case *PdfObjectNull:
		return nil, nil
	}
	return nil, ErrNotANumber
}

// ToFloat64Array returns a slice of all elements in the array as a float64 slice.  An error is
// returned if the array contains non-numeric objects (each element can be either PdfObjectInteger
// or PdfObjectFloat).
func (_eaea *PdfObjectArray) ToFloat64Array() ([]float64, error) {
	var _dbee []float64
	for _, _ffee := range _eaea.Elements() {
		switch _bbcb := _ffee.(type) {
		case *PdfObjectInteger:
			_dbee = append(_dbee, float64(*_bbcb))
		case *PdfObjectFloat:
			_dbee = append(_dbee, float64(*_bbcb))
		default:
			return nil, ErrTypeError
		}
	}
	return _dbee, nil
}

// JBIG2EncoderSettings contains the parameters and settings used by the JBIG2Encoder.
// Current version works only on JB2Generic compression.
type JBIG2EncoderSettings struct {

	// FileMode defines if the jbig2 encoder should return full jbig2 file instead of
	// shortened pdf mode. This adds the file header to the jbig2 definition.
	FileMode bool

	// Compression is the setting that defines the compression type used for encoding the page.
	Compression JBIG2CompressionType

	// DuplicatedLinesRemoval code generic region in a way such that if the lines are duplicated the encoder
	// doesn't store it twice.
	DuplicatedLinesRemoval bool

	// DefaultPixelValue is the bit value initial for every pixel in the page.
	DefaultPixelValue uint8

	// ResolutionX optional setting that defines the 'x' axis input image resolution - used for single page encoding.
	ResolutionX int

	// ResolutionY optional setting that defines the 'y' axis input image resolution - used for single page encoding.
	ResolutionY int

	// Threshold defines the threshold of the image correlation for
	// non Generic compression.
	// User only for JB2SymbolCorrelation and JB2SymbolRankHaus methods.
	// Best results in range [0.7 - 0.98] - the less the better the compression would be
	// but the more lossy.
	// Default value: 0.95
	Threshold float64
}

func (_eea *PdfCrypt) decryptBytes(_acgd []byte, _ggbc string, _dfcg []byte) ([]byte, error) {
	_ae.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_bfab, _feea := _eea._agfd[_ggbc]
	if !_feea {
		return nil, _ac.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _ggbc)
	}
	return _bfab.DecryptBytes(_acgd, _dfcg)
}

// SetFileOffset sets the file to an offset position and resets buffer.
func (_dfbfe *PdfParser) SetFileOffset(offset int64) {
	if offset < 0 {
		offset = 0
	}
	_dfbfe._fdee.Seek(offset, _dgf.SeekStart)
	_dfbfe._eecea = _acg.NewReader(_dfbfe._fdee)
}

// DecodeStream decodes a LZW encoded stream and returns the result as a
// slice of bytes.
func (_bcec *LZWEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_ae.Log.Trace("\u004c\u005a\u0057 \u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067")
	_ae.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _bcec.Predictor)
	_bab, _ffa := _bcec.DecodeBytes(streamObj.Stream)
	if _ffa != nil {
		return nil, _ffa
	}
	_ae.Log.Trace("\u0020\u0049\u004e\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(streamObj.Stream), streamObj.Stream)
	_ae.Log.Trace("\u004f\u0055\u0054\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(_bab), _bab)
	if _bcec.Predictor > 1 {
		if _bcec.Predictor == 2 {
			_ae.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_cgce := _bcec.Columns * _bcec.Colors
			if _cgce < 1 {
				return []byte{}, nil
			}
			_gaab := len(_bab) / _cgce
			if len(_bab)%_cgce != 0 {
				_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_bab), _cgce)
			}
			if _cgce%_bcec.Colors != 0 {
				return nil, _ac.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _cgce, _bcec.Colors)
			}
			if _cgce > len(_bab) {
				_ae.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _cgce, len(_bab))
				return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_ae.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_bab), _bab)
			_faga := _fd.NewBuffer(nil)
			for _geb := 0; _geb < _gaab; _geb++ {
				_adeg := _bab[_cgce*_geb : _cgce*(_geb+1)]
				for _eeaa := _bcec.Colors; _eeaa < _cgce; _eeaa++ {
					_adeg[_eeaa] = byte(int(_adeg[_eeaa]+_adeg[_eeaa-_bcec.Colors]) % 256)
				}
				_faga.Write(_adeg)
			}
			_fdab := _faga.Bytes()
			_ae.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_fdab), _fdab)
			return _fdab, nil
		} else if _bcec.Predictor >= 10 && _bcec.Predictor <= 15 {
			_ae.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_gcee := _bcec.Columns*_bcec.Colors + 1
			if _gcee < 1 {
				return []byte{}, nil
			}
			_debd := len(_bab) / _gcee
			if len(_bab)%_gcee != 0 {
				return nil, _ac.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_bab), _gcee)
			}
			if _gcee > len(_bab) {
				_ae.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _gcee, len(_bab))
				return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_cbcg := _fd.NewBuffer(nil)
			_ae.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _bcec.Columns)
			_ae.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_bab), _gcee, _debd)
			_dcef := make([]byte, _gcee)
			for _gdcc := 0; _gdcc < _gcee; _gdcc++ {
				_dcef[_gdcc] = 0
			}
			for _cef := 0; _cef < _debd; _cef++ {
				_fddd := _bab[_gcee*_cef : _gcee*(_cef+1)]
				_bcgf := _fddd[0]
				switch _bcgf {
				case 0:
				case 1:
					for _ece := 2; _ece < _gcee; _ece++ {
						_fddd[_ece] = byte(int(_fddd[_ece]+_fddd[_ece-1]) % 256)
					}
				case 2:
					for _gabe := 1; _gabe < _gcee; _gabe++ {
						_fddd[_gabe] = byte(int(_fddd[_gabe]+_dcef[_gabe]) % 256)
					}
				default:
					_ae.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0066i\u006c\u0074\u0065\u0072\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _bcgf)
					return nil, _ac.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _bcgf)
				}
				for _baeb := 0; _baeb < _gcee; _baeb++ {
					_dcef[_baeb] = _fddd[_baeb]
				}
				_cbcg.Write(_fddd[1:])
			}
			_gdfc := _cbcg.Bytes()
			return _gdfc, nil
		} else {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _bcec.Predictor)
			return nil, _ac.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _bcec.Predictor)
		}
	}
	return _bab, nil
}

// PdfObjectStream represents the primitive PDF Object stream.
type PdfObjectStream struct {
	PdfObjectReference
	*PdfObjectDictionary
	Stream []byte
}

// MakeFloat creates an PdfObjectFloat from a float64.
func MakeFloat(val float64) *PdfObjectFloat { _dbcgb := PdfObjectFloat(val); return &_dbcgb }

// GetFilterName returns the name of the encoding filter.
func (_dbba *RawEncoder) GetFilterName() string { return StreamEncodingFilterNameRaw }

// GetFloatVal returns the float64 value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetFloatVal(obj PdfObject) (_fbecd float64, _defc bool) {
	_deeg, _defc := TraceToDirectObject(obj).(*PdfObjectFloat)
	if _defc {
		return float64(*_deeg), true
	}
	return 0, false
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_feg *RunLengthEncoder) MakeStreamDict() *PdfObjectDictionary {
	_dbbe := MakeDict()
	_dbbe.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_feg.GetFilterName()))
	return _dbbe
}

// HasInvalidHexRunes implements core.ParserMetadata interface.
func (_ffg ParserMetadata) HasInvalidHexRunes() bool { return _ffg._bbec }

// PdfCrypt provides PDF encryption/decryption support.
// The PDF standard supports encryption of strings and streams (Section 7.6).
type PdfCrypt struct {
	_ddc  encryptDict
	_gge  _geg.StdEncryptDict
	_dfc  string
	_dadd []byte
	_dgd  map[PdfObject]bool
	_egg  map[PdfObject]bool
	_gfg  bool
	_agfd cryptFilters
	_bba  string
	_dee  string
	_cdg  *PdfParser
	_dfg  map[int]struct{}
}

// MakeIndirectObject creates an PdfIndirectObject with a specified direct object PdfObject.
func MakeIndirectObject(obj PdfObject) *PdfIndirectObject {
	_fcdf := &PdfIndirectObject{}
	_fcdf.PdfObject = obj
	return _fcdf
}

// DecodeStream decodes a DCT encoded stream and returns the result as a
// slice of bytes.
func (_efd *DCTEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _efd.DecodeBytes(streamObj.Stream)
}

// Elements returns a slice of the PdfObject elements in the array.
// Preferred over accessing the array directly as type may be changed in future major versions (v3).
func (_cece *PdfObjectStreams) Elements() []PdfObject {
	if _cece == nil {
		return nil
	}
	return _cece._dgebc
}

// Resolve resolves the reference and returns the indirect or stream object.
// If the reference cannot be resolved, a *PdfObjectNull object is returned.
func (_ebcec *PdfObjectReference) Resolve() PdfObject {
	if _ebcec._ffgd == nil {
		return MakeNull()
	}
	_bfec, _, _cgcea := _ebcec._ffgd.resolveReference(_ebcec)
	if _cgcea != nil {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0072\u0065\u0073\u006f\u006cv\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065r\u0065n\u0063\u0065\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074", _cgcea)
		return MakeNull()
	}
	if _bfec == nil {
		_ae.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0072\u0065\u0073ol\u0076\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065:\u0020\u006ei\u006c\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0061\u0020nu\u006c\u006c\u0020o\u0062\u006a\u0065\u0063\u0074")
		return MakeNull()
	}
	return _bfec
}
func _bbdb(_gagd string) (int, int, error) {
	_beaf := _ddce.FindStringSubmatch(_gagd)
	if len(_beaf) < 3 {
		return 0, 0, _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_abfbc, _ := _dg.Atoi(_beaf[1])
	_gabd, _ := _dg.Atoi(_beaf[2])
	return _abfbc, _gabd, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on the current encoder settings.
func (_ggda *JBIG2Encoder) MakeDecodeParams() PdfObject { return MakeDict() }
func _afef(_fadg _dgf.ReadSeeker, _beef int64) (*limitedReadSeeker, error) {
	_, _abae := _fadg.Seek(0, _dgf.SeekStart)
	if _abae != nil {
		return nil, _abae
	}
	return &limitedReadSeeker{_cgfa: _fadg, _gfcg: _beef}, nil
}

// PdfCryptNewEncrypt makes the document crypt handler based on a specified crypt filter.
func PdfCryptNewEncrypt(cf _ebd.Filter, userPass, ownerPass []byte, perm _geg.Permissions) (*PdfCrypt, *EncryptInfo, error) {
	_dca := &PdfCrypt{_egg: make(map[PdfObject]bool), _agfd: make(cryptFilters), _gge: _geg.StdEncryptDict{P: perm, EncryptMetadata: true}}
	var _acgc Version
	if cf != nil {
		_ceg := cf.PDFVersion()
		_acgc.Major, _acgc.Minor = _ceg[0], _ceg[1]
		V, R := cf.HandlerVersion()
		_dca._ddc.V = V
		_dca._gge.R = R
		_dca._ddc.Length = cf.KeyLength() * 8
	}
	const (
		_aeb = _egb
	)
	_dca._agfd[_aeb] = cf
	if _dca._ddc.V >= 4 {
		_dca._bba = _aeb
		_dca._dee = _aeb
	}
	_agbb := _dca.newEncryptDict()
	_gaf := _eae.Sum([]byte(_ga.Now().Format(_ga.RFC850)))
	_gff := string(_gaf[:])
	_afg := make([]byte, 100)
	_ag.Read(_afg)
	_gaf = _eae.Sum(_afg)
	_cbe := string(_gaf[:])
	_ae.Log.Trace("\u0052\u0061\u006e\u0064\u006f\u006d\u0020\u0062\u003a\u0020\u0025\u0020\u0078", _afg)
	_ae.Log.Trace("\u0047\u0065\u006e\u0020\u0049\u0064\u0020\u0030\u003a\u0020\u0025\u0020\u0078", _gff)
	_dca._dfc = _gff
	_bce := _dca.generateParams(userPass, ownerPass)
	if _bce != nil {
		return nil, nil, _bce
	}
	_fge(&_dca._gge, _agbb)
	if _dca._ddc.V >= 4 {
		if _baff := _dca.saveCryptFilters(_agbb); _baff != nil {
			return nil, nil, _baff
		}
	}
	return _dca, &EncryptInfo{Version: _acgc, Encrypt: _agbb, ID0: _gff, ID1: _cbe}, nil
}

type encryptDict struct {
	Filter    string
	V         int
	SubFilter string
	Length    int
	StmF      string
	StrF      string
	EFF       string
	CF        map[string]_ebd.FilterDict
}

// GetArray returns the *PdfObjectArray represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetArray(obj PdfObject) (_cfeg *PdfObjectArray, _fdeb bool) {
	_cfeg, _fdeb = TraceToDirectObject(obj).(*PdfObjectArray)
	return _cfeg, _fdeb
}

// EncodeBytes JPX encodes the passed in slice of bytes.
func (_bgge *JPXEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _bgge.GetFilterName())
	return data, ErrNoJPXDecode
}

// GetParser returns the parser for lazy-loading or compare references.
func (_ecbc *PdfObjectReference) GetParser() *PdfParser { return _ecbc._ffgd }

// MakeHexString creates an PdfObjectString from a string intended for output as a hexadecimal string.
func MakeHexString(s string) *PdfObjectString {
	_ebgdc := PdfObjectString{_eeee: s, _fcge: true}
	return &_ebgdc
}

// WriteString outputs the object as it is to be written to file.
func (_ffcd *PdfObjectString) WriteString() string {
	var _gafd _fd.Buffer
	if _ffcd._fcge {
		_cgacb := _cab.EncodeToString(_ffcd.Bytes())
		_gafd.WriteString("\u003c")
		_gafd.WriteString(_cgacb)
		_gafd.WriteString("\u003e")
		return _gafd.String()
	}
	_ceed := map[byte]string{'\n': "\u005c\u006e", '\r': "\u005c\u0072", '\t': "\u005c\u0074", '\b': "\u005c\u0062", '\f': "\u005c\u0066", '(': "\u005c\u0028", ')': "\u005c\u0029", '\\': "\u005c\u005c"}
	_gafd.WriteString("\u0028")
	for _ffceg := 0; _ffceg < len(_ffcd._eeee); _ffceg++ {
		_dbbd := _ffcd._eeee[_ffceg]
		if _acdc, _bfgc := _ceed[_dbbd]; _bfgc {
			_gafd.WriteString(_acdc)
		} else {
			_gafd.WriteByte(_dbbd)
		}
	}
	_gafd.WriteString("\u0029")
	return _gafd.String()
}

// EncodeJBIG2Image encodes 'img' into jbig2 encoded bytes stream, using default encoder settings.
func (_bade *JBIG2Encoder) EncodeJBIG2Image(img *JBIG2Image) ([]byte, error) {
	const _cbb = "c\u006f\u0072\u0065\u002eEn\u0063o\u0064\u0065\u004a\u0042\u0049G\u0032\u0049\u006d\u0061\u0067\u0065"
	if _gdbcc := _bade.AddPageImage(img, &_bade.DefaultPageSettings); _gdbcc != nil {
		return nil, _eb.Wrap(_gdbcc, _cbb, "")
	}
	return _bade.Encode()
}

type xrefType int

var _ccd = []PdfObjectName{"\u0056", "\u0052", "\u004f", "\u0055", "\u0050"}
var _eecg = _ba.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

// DecodeBytes decodes a slice of LZW encoded bytes and returns the result.
func (_gaad *LZWEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _fddc _fd.Buffer
	_age := _fd.NewReader(encoded)
	var _fae _dgf.ReadCloser
	if _gaad.EarlyChange == 1 {
		_fae = _cge.NewReader(_age, _cge.MSB, 8)
	} else {
		_fae = _a.NewReader(_age, _a.MSB, 8)
	}
	defer _fae.Close()
	if _, _eefcd := _fddc.ReadFrom(_fae); _eefcd != nil {
		if _eefcd != _dgf.ErrUnexpectedEOF || _fddc.Len() == 0 {
			return nil, _eefcd
		}
		_ae.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u004c\u005a\u0057\u0020\u0064\u0065\u0063\u006f\u0064i\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076\u002e \u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062e \u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _eefcd)
	}
	return _fddc.Bytes(), nil
}

// ParseDict reads and parses a PDF dictionary object enclosed with '<<' and '>>'
func (_ebgad *PdfParser) ParseDict() (*PdfObjectDictionary, error) {
	_ae.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_cdff := MakeDict()
	_cdff._fdddg = _ebgad
	_fcee, _ := _ebgad._eecea.ReadByte()
	if _fcee != '<' {
		return nil, _d.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_fcee, _ = _ebgad._eecea.ReadByte()
	if _fcee != '<' {
		return nil, _d.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_ebgad.skipSpaces()
		_ebgad.skipComments()
		_bdda, _bcdec := _ebgad._eecea.Peek(2)
		if _bcdec != nil {
			return nil, _bcdec
		}
		_ae.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_bdda), string(_bdda))
		if (_bdda[0] == '>') && (_bdda[1] == '>') {
			_ae.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_ebgad._eecea.ReadByte()
			_ebgad._eecea.ReadByte()
			break
		}
		_ae.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_fgcaf, _bcdec := _ebgad.parseName()
		_ae.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _fgcaf)
		if _bcdec != nil {
			_ae.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _bcdec)
			return nil, _bcdec
		}
		if len(_fgcaf) > 4 && _fgcaf[len(_fgcaf)-4:] == "\u006e\u0075\u006c\u006c" {
			_gddba := _fgcaf[0 : len(_fgcaf)-4]
			_ae.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _fgcaf)
			_ae.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _gddba)
			_ebgad.skipSpaces()
			_bbab, _ := _ebgad._eecea.Peek(1)
			if _bbab[0] == '/' {
				_cdff.Set(_gddba, MakeNull())
				continue
			}
		}
		_ebgad.skipSpaces()
		_fgcb, _bcdec := _ebgad.parseObject()
		if _bcdec != nil {
			return nil, _bcdec
		}
		_cdff.Set(_fgcaf, _fgcb)
		if _ae.Log.IsLogLevel(_ae.LogLevelTrace) {
			_ae.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _fgcaf, _fgcb.String())
		}
	}
	_ae.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _cdff, nil
}

// GetXrefTable returns the PDFs xref table.
func (_aebb *PdfParser) GetXrefTable() XrefTable { return _aebb._fbab }

// IsAuthenticated returns true if the PDF has already been authenticated for accessing.
func (_ddedb *PdfParser) IsAuthenticated() bool { return _ddedb._bffd._gfg }

// CheckAccessRights checks access rights and permissions for a specified password. If either user/owner password is
// specified, full rights are granted, otherwise the access rights are specified by the Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_dgace *PdfParser) CheckAccessRights(password []byte) (bool, _geg.Permissions, error) {
	if _dgace._bffd == nil {
		return true, _geg.PermOwner, nil
	}
	return _dgace._bffd.checkAccessRights(password)
}

// Read implementation of Read interface.
func (_ggdbb *limitedReadSeeker) Read(p []byte) (_fedc int, _eggaca error) {
	_bgdcc, _eggaca := _ggdbb._cgfa.Seek(0, _dgf.SeekCurrent)
	if _eggaca != nil {
		return 0, _eggaca
	}
	_bacaf := _ggdbb._gfcg - _bgdcc
	if _bacaf == 0 {
		return 0, _dgf.EOF
	}
	if _dgdb := int64(len(p)); _dgdb < _bacaf {
		_bacaf = _dgdb
	}
	_dcbe := make([]byte, _bacaf)
	_fedc, _eggaca = _ggdbb._cgfa.Read(_dcbe)
	copy(p, _dcbe)
	return _fedc, _eggaca
}

// DecodeStream decodes RunLengthEncoded stream object and give back decoded bytes.
func (_fcfd *RunLengthEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _fcfd.DecodeBytes(streamObj.Stream)
}

var _fgadb = _ba.MustCompile("\u0025P\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

// GetEncryptObj returns the PdfIndirectObject which has information about the PDFs encryption details.
func (_gdgc *PdfParser) GetEncryptObj() *PdfIndirectObject { return _gdgc._bcdd }

// RegisterCustomStreamEncoder register a custom encoder handler for certain filter.
func RegisterCustomStreamEncoder(filterName string, customStreamEncoder StreamEncoder) {
	_dggc.Store(filterName, customStreamEncoder)
}
func (_dabb *PdfParser) parseXrefStream(_agcb *PdfObjectInteger) (*PdfObjectDictionary, error) {
	if _agcb != nil {
		_ae.Log.Trace("\u0058\u0052\u0065f\u0053\u0074\u006d\u0020x\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0061\u0074\u0020\u0025\u0064", _agcb)
		_dabb._fdee.Seek(int64(*_agcb), _dgf.SeekStart)
		_dabb._eecea = _acg.NewReader(_dabb._fdee)
	}
	_caacb := _dabb.GetFileOffset()
	_gdca, _fbcd := _dabb.ParseIndirectObject()
	if _fbcd != nil {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0078\u0072\u0065\u0066\u0020\u006fb\u006a\u0065\u0063\u0074")
		return nil, _d.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072e\u0061\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	_ae.Log.Trace("\u0058R\u0065f\u0053\u0074\u006d\u0020\u006fb\u006a\u0065c\u0074\u003a\u0020\u0025\u0073", _gdca)
	_cacf, _adcf := _gdca.(*PdfObjectStream)
	if !_adcf {
		_ae.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0058R\u0065\u0066\u0053\u0074\u006d\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067 \u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0021")
		return nil, _d.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0070\u006f\u0069\u006e\u0074i\u006e\u0067\u0020\u0074\u006f\u0020a\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_feed := _cacf.PdfObjectDictionary
	_fgdef, _adcf := _cacf.PdfObjectDictionary.Get("\u0053\u0069\u007a\u0065").(*PdfObjectInteger)
	if !_adcf {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0073\u0069\u007a\u0065\u0020f\u0072\u006f\u006d\u0020\u0078\u0072\u0065f\u0020\u0073\u0074\u006d")
		return nil, _d.New("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0053\u0069\u007ae\u0020\u0066\u0072\u006f\u006d\u0020\u0078\u0072\u0065\u0066 \u0073\u0074\u006d")
	}
	if int64(*_fgdef) > 8388607 {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u0053\u0069\u007a\u0065\u0020\u0065x\u0063\u0065\u0065\u0064\u0065\u0064\u0020l\u0069\u006d\u0069\u0074\u002c\u0020\u006f\u0076\u0065\u0072\u00208\u0033\u0038\u0038\u0036\u0030\u0037\u0020\u0028\u0025\u0064\u0029", *_fgdef)
		return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_caecb := _cacf.PdfObjectDictionary.Get("\u0057")
	_dfea, _adcf := _caecb.(*PdfObjectArray)
	if !_adcf {
		return nil, _d.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0020\u0069\u006e\u0020x\u0072\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_bcdf := _dfea.Len()
	if _bcdf != 3 {
		_ae.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020\u0028\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033\u0020\u002d\u0020\u0025\u0064\u0029", _bcdf)
		return nil, _d.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065f\u0020s\u0074\u006d\u0020\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033")
	}
	var _ffce []int64
	for _aead := 0; _aead < 3; _aead++ {
		_baba, _ffbb := GetInt(_dfea.Get(_aead))
		if !_ffbb {
			return nil, _d.New("i\u006e\u0076\u0061\u006cid\u0020w\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		_ffce = append(_ffce, int64(*_baba))
	}
	_adbc, _fbcd := DecodeStream(_cacf)
	if _fbcd != nil {
		_ae.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f \u0064e\u0063o\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _fbcd)
		return nil, _fbcd
	}
	_dfbfb := int(_ffce[0])
	_bcfa := int(_ffce[0] + _ffce[1])
	_feaa := int(_ffce[0] + _ffce[1] + _ffce[2])
	_ccfgc := int(_ffce[0] + _ffce[1] + _ffce[2])
	if _dfbfb < 0 || _bcfa < 0 || _feaa < 0 {
		_ae.Log.Debug("\u0045\u0072\u0072\u006fr\u0020\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u003c \u0030 \u0028\u0025\u0064\u002c\u0025\u0064\u002c%\u0064\u0029", _dfbfb, _bcfa, _feaa)
		return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _ccfgc == 0 {
		_ae.Log.Debug("\u004e\u006f\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0069\u006e\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028\u0064\u0065\u006c\u0074\u0061\u0062\u0020=\u003d\u0020\u0030\u0029")
		return _feed, nil
	}
	_fccg := len(_adbc) / _ccfgc
	_dcgbg := 0
	_dfbbe := _cacf.PdfObjectDictionary.Get("\u0049\u006e\u0064e\u0078")
	var _bbad []int
	if _dfbbe != nil {
		_ae.Log.Trace("\u0049n\u0064\u0065\u0078\u003a\u0020\u0025b", _dfbbe)
		_aeec, _ccfed := _dfbbe.(*PdfObjectArray)
		if !_ccfed {
			_ae.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0029")
			return nil, _d.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0049\u006e\u0064e\u0078\u0020\u006f\u0062je\u0063\u0074")
		}
		if _aeec.Len()%2 != 0 {
			_ae.Log.Debug("\u0057\u0041\u0052\u004eI\u004e\u0047\u0020\u0046\u0061\u0069\u006c\u0075\u0072e\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020i\u006e\u0064\u0065\u0078\u0020n\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006f\u0066\u0020\u0032\u002e")
			return nil, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
		}
		_dcgbg = 0
		_fffa, _ddca := _aeec.ToIntegerArray()
		if _ddca != nil {
			_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0065\u0078 \u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0073\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0073\u003a\u0020\u0025\u0076", _ddca)
			return nil, _ddca
		}
		for _abfc := 0; _abfc < len(_fffa); _abfc += 2 {
			_fcdb := _fffa[_abfc]
			_eeeb := _fffa[_abfc+1]
			for _dead := 0; _dead < _eeeb; _dead++ {
				_bbad = append(_bbad, _fcdb+_dead)
			}
			_dcgbg += _eeeb
		}
	} else {
		for _adebd := 0; _adebd < int(*_fgdef); _adebd++ {
			_bbad = append(_bbad, _adebd)
		}
		_dcgbg = int(*_fgdef)
	}
	if _fccg == _dcgbg+1 {
		_ae.Log.Debug("\u0049n\u0063\u006f\u006d\u0070ati\u0062\u0069\u006c\u0069t\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u0076\u0065\u0072\u0061\u0067\u0065\u0020\u006f\u0066\u0020\u0031\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u002d\u0020\u0061\u0070\u0070en\u0064\u0069\u006eg\u0020\u006f\u006e\u0065\u0020-\u0020M\u0061\u0079\u0020\u006c\u0065\u0061\u0064\u0020\u0074o\u0020\u0070\u0072\u006f\u0062\u006c\u0065\u006d\u0073")
		_feac := _dcgbg - 1
		for _, _dbdfc := range _bbad {
			if _dbdfc > _feac {
				_feac = _dbdfc
			}
		}
		_bbad = append(_bbad, _feac+1)
		_dcgbg++
	}
	if _fccg != len(_bbad) {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020x\u0072\u0065\u0066 \u0073\u0074\u006d:\u0020\u006eu\u006d\u0020\u0065\u006e\u0074\u0072i\u0065s \u0021\u003d\u0020\u006c\u0065\u006e\u0028\u0069\u006e\u0064\u0069\u0063\u0065\u0073\u0029\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", _fccg, len(_bbad))
		return nil, _d.New("\u0078\u0072ef\u0020\u0073\u0074m\u0020\u006e\u0075\u006d en\u0074ri\u0065\u0073\u0020\u0021\u003d\u0020\u006cen\u0028\u0069\u006e\u0064\u0069\u0063\u0065s\u0029")
	}
	_ae.Log.Trace("\u004f\u0062j\u0065\u0063\u0074s\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0025\u0064", _dcgbg)
	_ae.Log.Trace("\u0049\u006e\u0064i\u0063\u0065\u0073\u003a\u0020\u0025\u0020\u0064", _bbad)
	_daed := func(_agef []byte) int64 {
		var _gdge int64
		for _gbgg := 0; _gbgg < len(_agef); _gbgg++ {
			_gdge += int64(_agef[_gbgg]) * (1 << uint(8*(len(_agef)-_gbgg-1)))
		}
		return _gdge
	}
	_ae.Log.Trace("\u0044e\u0063\u006f\u0064\u0065d\u0020\u0073\u0074\u0072\u0065a\u006d \u006ce\u006e\u0067\u0074\u0068\u003a\u0020\u0025d", len(_adbc))
	_aggg := 0
	for _gebda := 0; _gebda < len(_adbc); _gebda += _ccfgc {
		_bfacde := _adfaf(len(_adbc), _gebda, _gebda+_dfbfb)
		if _bfacde != nil {
			_ae.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _bfacde)
			return nil, _bfacde
		}
		_gfea := _adbc[_gebda : _gebda+_dfbfb]
		_bfacde = _adfaf(len(_adbc), _gebda+_dfbfb, _gebda+_bcfa)
		if _bfacde != nil {
			_ae.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _bfacde)
			return nil, _bfacde
		}
		_efca := _adbc[_gebda+_dfbfb : _gebda+_bcfa]
		_bfacde = _adfaf(len(_adbc), _gebda+_bcfa, _gebda+_feaa)
		if _bfacde != nil {
			_ae.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _bfacde)
			return nil, _bfacde
		}
		_daaaf := _adbc[_gebda+_bcfa : _gebda+_feaa]
		_addb := _daed(_gfea)
		_ebbba := _daed(_efca)
		_degea := _daed(_daaaf)
		if _ffce[0] == 0 {
			_addb = 1
		}
		if _aggg >= len(_bbad) {
			_ae.Log.Debug("X\u0052\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u002d\u0020\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063e\u0073s\u0020\u0069\u006e\u0064e\u0078\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073\u0020\u002d\u0020\u0062\u0072\u0065\u0061\u006b\u0069\u006e\u0067")
			break
		}
		_bdeaa := _bbad[_aggg]
		_aggg++
		_ae.Log.Trace("%\u0064\u002e\u0020\u0070\u0031\u003a\u0020\u0025\u0020\u0078", _bdeaa, _gfea)
		_ae.Log.Trace("%\u0064\u002e\u0020\u0070\u0032\u003a\u0020\u0025\u0020\u0078", _bdeaa, _efca)
		_ae.Log.Trace("%\u0064\u002e\u0020\u0070\u0033\u003a\u0020\u0025\u0020\u0078", _bdeaa, _daaaf)
		_ae.Log.Trace("\u0025d\u002e \u0078\u0072\u0065\u0066\u003a \u0025\u0064 \u0025\u0064\u0020\u0025\u0064", _bdeaa, _addb, _ebbba, _degea)
		if _addb == 0 {
			_ae.Log.Trace("-\u0020\u0046\u0072\u0065\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020-\u0020\u0063\u0061\u006e\u0020\u0070\u0072o\u0062\u0061\u0062\u006c\u0079\u0020\u0069\u0067\u006e\u006fr\u0065")
		} else if _addb == 1 {
			_ae.Log.Trace("\u002d\u0020I\u006e\u0020\u0075\u0073e\u0020\u002d \u0075\u006e\u0063\u006f\u006d\u0070\u0072\u0065s\u0073\u0065\u0064\u0020\u0076\u0069\u0061\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0062", _efca)
			if _ebbba == _caacb {
				_ae.Log.Debug("\u0055\u0070d\u0061\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0058\u0052\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064", _bdeaa, _cacf.ObjectNumber)
				_bdeaa = int(_cacf.ObjectNumber)
			}
			if _dade, _ffbc := _dabb._fbab.ObjectMap[_bdeaa]; !_ffbc || int(_degea) > _dade.Generation {
				_fbbf := XrefObject{ObjectNumber: _bdeaa, XType: XrefTypeTableEntry, Offset: _ebbba, Generation: int(_degea)}
				_dabb._fbab.ObjectMap[_bdeaa] = _fbbf
			}
		} else if _addb == 2 {
			_ae.Log.Trace("\u002d\u0020\u0049\u006e \u0075\u0073\u0065\u0020\u002d\u0020\u0063\u006f\u006d\u0070r\u0065s\u0073\u0065\u0064\u0020\u006f\u0062\u006ae\u0063\u0074")
			if _, _faagc := _dabb._fbab.ObjectMap[_bdeaa]; !_faagc {
				_edbb := XrefObject{ObjectNumber: _bdeaa, XType: XrefTypeObjectStream, OsObjNumber: int(_ebbba), OsObjIndex: int(_degea)}
				_dabb._fbab.ObjectMap[_bdeaa] = _edbb
				_ae.Log.Trace("\u0065\u006e\u0074\u0072\u0079\u003a\u0020\u0025\u002b\u0076", _edbb)
			}
		} else {
			_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0056\u0041L\u0049\u0044\u0020\u0054\u0059\u0050\u0045\u0020\u0058\u0072\u0065\u0066\u0053\u0074\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u003f\u002d\u002d\u002d\u002d\u002d\u002d-")
			continue
		}
	}
	if _dabb._dfcgd == nil {
		_ecfd := XrefTypeObjectStream
		_dabb._dfcgd = &_ecfd
	}
	return _feed, nil
}

// NewParserFromString is used for testing purposes.
func NewParserFromString(txt string) *PdfParser {
	_afgb := _fd.NewReader([]byte(txt))
	_fdcaa := &PdfParser{ObjCache: objectCache{}, _fdee: _afgb, _eecea: _acg.NewReader(_afgb), _fcca: int64(len(txt)), _gfdf: map[int64]bool{}, _bgcdd: make(map[*PdfParser]*PdfParser)}
	_fdcaa._fbab.ObjectMap = make(map[int]XrefObject)
	return _fdcaa
}

// NewJBIG2Encoder creates a new JBIG2Encoder.
func NewJBIG2Encoder() *JBIG2Encoder { return &JBIG2Encoder{_ebac: _da.InitEncodeDocument(false)} }

// GetString returns the *PdfObjectString represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetString(obj PdfObject) (_agddb *PdfObjectString, _ddbb bool) {
	_agddb, _ddbb = TraceToDirectObject(obj).(*PdfObjectString)
	return _agddb, _ddbb
}

// MakeDict creates and returns an empty PdfObjectDictionary.
func MakeDict() *PdfObjectDictionary {
	_caaf := &PdfObjectDictionary{}
	_caaf._gged = map[PdfObjectName]PdfObject{}
	_caaf._dgcd = []PdfObjectName{}
	_caaf._efce = &_e.Mutex{}
	return _caaf
}

// PdfObjectString represents the primitive PDF string object.
type PdfObjectString struct {
	_eeee string
	_fcge bool
}

func (_fecc *PdfParser) parseXrefTable() (*PdfObjectDictionary, error) {
	var _afac *PdfObjectDictionary
	_dcabb, _efdd := _fecc.readTextLine()
	if _efdd != nil {
		return nil, _efdd
	}
	if _fecc._ccce && _agg.Count(_agg.TrimPrefix(_dcabb, "\u0078\u0072\u0065\u0066"), "\u0020") > 0 {
		_fecc._aecec._bega = true
	}
	_ae.Log.Trace("\u0078\u0072\u0065\u0066 f\u0069\u0072\u0073\u0074\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _dcabb)
	_befa := -1
	_dcda := 0
	_cfdf := false
	_fecfb := ""
	for {
		_fecc.skipSpaces()
		_, _ffdc := _fecc._eecea.Peek(1)
		if _ffdc != nil {
			return nil, _ffdc
		}
		_dcabb, _ffdc = _fecc.readTextLine()
		if _ffdc != nil {
			return nil, _ffdc
		}
		_acbg := _bffa.FindStringSubmatch(_dcabb)
		if len(_acbg) == 0 {
			_cffd := len(_fecfb) > 0
			_fecfb += _dcabb + "\u000a"
			if _cffd {
				_acbg = _bffa.FindStringSubmatch(_fecfb)
			}
		}
		if len(_acbg) == 3 {
			if _fecc._ccce && !_fecc._aecec._bcg {
				var (
					_dcfab bool
					_gagg  int
				)
				for _, _cedf := range _dcabb {
					if _b.IsDigit(_cedf) {
						if _dcfab {
							break
						}
						continue
					}
					if !_dcfab {
						_dcfab = true
					}
					_gagg++
				}
				if _gagg > 1 {
					_fecc._aecec._bcg = true
				}
			}
			_badfb, _ := _dg.Atoi(_acbg[1])
			_bfdb, _ := _dg.Atoi(_acbg[2])
			_befa = _badfb
			_dcda = _bfdb
			_cfdf = true
			_fecfb = ""
			_ae.Log.Trace("\u0078r\u0065\u0066 \u0073\u0075\u0062s\u0065\u0063\u0074\u0069\u006f\u006e\u003a \u0066\u0069\u0072\u0073\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0064\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _befa, _dcda)
			continue
		}
		_cggfb := _feffd.FindStringSubmatch(_dcabb)
		if len(_cggfb) == 4 {
			if !_cfdf {
				_ae.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0058r\u0065\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006fr\u006da\u0074\u0021\u000a")
				return nil, _d.New("\u0078\u0072\u0065\u0066 i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_fgfa, _ := _dg.ParseInt(_cggfb[1], 10, 64)
			_ddbc, _ := _dg.Atoi(_cggfb[2])
			_gbag := _cggfb[3]
			_fecfb = ""
			if _agg.ToLower(_gbag) == "\u006e" && _fgfa > 1 {
				_fbecb, _gfgb := _fecc._fbab.ObjectMap[_befa]
				if !_gfgb || _ddbc > _fbecb.Generation {
					_caeb := XrefObject{ObjectNumber: _befa, XType: XrefTypeTableEntry, Offset: _fgfa, Generation: _ddbc}
					_fecc._fbab.ObjectMap[_befa] = _caeb
				}
			}
			_befa++
			continue
		}
		if (len(_dcabb) > 6) && (_dcabb[:7] == "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_ae.Log.Trace("\u0046o\u0075n\u0064\u0020\u0074\u0072\u0061i\u006c\u0065r\u0020\u002d\u0020\u0025\u0073", _dcabb)
			if len(_dcabb) > 9 {
				_gffg := _fecc.GetFileOffset()
				_fecc.SetFileOffset(_gffg - int64(len(_dcabb)) + 7)
			}
			_fecc.skipSpaces()
			_fecc.skipComments()
			_ae.Log.Trace("R\u0065\u0061\u0064\u0069ng\u0020t\u0072\u0061\u0069\u006c\u0065r\u0020\u0064\u0069\u0063\u0074\u0021")
			_ae.Log.Trace("\u0070\u0065\u0065\u006b\u003a\u0020\u0022\u0025\u0073\u0022", _dcabb)
			_afac, _ffdc = _fecc.ParseDict()
			_ae.Log.Trace("\u0045O\u0046\u0020\u0072\u0065a\u0064\u0069\u006e\u0067\u0020t\u0072a\u0069l\u0065\u0072\u0020\u0064\u0069\u0063\u0074!")
			if _ffdc != nil {
				_ae.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _ffdc)
				return nil, _ffdc
			}
			break
		}
		if _dcabb == "\u0025\u0025\u0045O\u0046" {
			_ae.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006e\u0064 \u006f\u0066\u0020\u0066\u0069\u006c\u0065 -\u0020\u0074\u0072\u0061i\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066ou\u006e\u0064 \u002d\u0020\u0065\u0072\u0072\u006f\u0072\u0021")
			return nil, _d.New("\u0065\u006e\u0064 \u006f\u0066\u0020\u0066i\u006c\u0065\u0020\u002d\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_ae.Log.Trace("\u0078\u0072\u0065\u0066\u0020\u006d\u006f\u0072\u0065 \u003a\u0020\u0025\u0073", _dcabb)
	}
	_ae.Log.Trace("\u0045\u004f\u0046 p\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0021")
	if _fecc._dfcgd == nil {
		_fced := XrefTypeTableEntry
		_fecc._dfcgd = &_fced
	}
	return _afac, nil
}

// GetXrefType returns the type of the first xref object (table or stream).
func (_afbc *PdfParser) GetXrefType() *xrefType { return _afbc._dfcgd }

type offsetReader struct {
	_agbf _dgf.ReadSeeker
	_ceca int64
}

// SetIfNotNil sets the dictionary's key -> val mapping entry -IF- val is not nil.
// Note that we take care to perform a type switch.  Otherwise if we would supply a nil value
// of another type, e.g. (PdfObjectArray*)(nil), then it would not be a PdfObject(nil) and thus
// would get set.
func (_eecdb *PdfObjectDictionary) SetIfNotNil(key PdfObjectName, val PdfObject) {
	if val != nil {
		switch _bead := val.(type) {
		case *PdfObjectName:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectDictionary:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectStream:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectString:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectNull:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectInteger:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectArray:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectBool:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectFloat:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfObjectReference:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		case *PdfIndirectObject:
			if _bead != nil {
				_eecdb.Set(key, val)
			}
		default:
			_ae.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", val)
		}
	}
}

// NewASCIIHexEncoder makes a new ASCII hex encoder.
func NewASCIIHexEncoder() *ASCIIHexEncoder { _gfge := &ASCIIHexEncoder{}; return _gfge }

// MakeArrayFromIntegers64 creates an PdfObjectArray from a slice of int64s, where each array element
// is an PdfObjectInteger.
func MakeArrayFromIntegers64(vals []int64) *PdfObjectArray {
	_gfdgb := MakeArray()
	for _, _bacf := range vals {
		_gfdgb.Append(MakeInteger(_bacf))
	}
	return _gfdgb
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_ggbb *CCITTFaxEncoder) MakeStreamDict() *PdfObjectDictionary {
	_bgc := MakeDict()
	_bgc.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_ggbb.GetFilterName()))
	_bgc.SetIfNotNil("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _ggbb.MakeDecodeParams())
	return _bgc
}

// GetUpdatedObjects returns pdf objects which were updated from the specific version (from prevParser).
func (_cbdb *PdfParser) GetUpdatedObjects(prevParser *PdfParser) (map[int64]PdfObject, error) {
	if prevParser == nil {
		return nil, _d.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_dedfg, _bcfe := _cbdb.getNumbersOfUpdatedObjects(prevParser)
	if _bcfe != nil {
		return nil, _bcfe
	}
	_gfdfb := make(map[int64]PdfObject)
	for _, _agbc := range _dedfg {
		if _dfgd, _gcdd := _cbdb.LookupByNumber(_agbc); _gcdd == nil {
			_gfdfb[int64(_agbc)] = _dfgd
		} else {
			return nil, _gcdd
		}
	}
	return _gfdfb, nil
}
func (_gbdc *PdfParser) parseNull() (PdfObjectNull, error) {
	_, _ffeb := _gbdc._eecea.Discard(4)
	return PdfObjectNull{}, _ffeb
}
func _bcc(_adba int) int { _abed := _adba >> (_fdbde - 1); return (_adba ^ _abed) - _abed }
func _fge(_ead *_geg.StdEncryptDict, _aga *PdfObjectDictionary) {
	_aga.Set("\u0052", MakeInteger(int64(_ead.R)))
	_aga.Set("\u0050", MakeInteger(int64(_ead.P)))
	_aga.Set("\u004f", MakeStringFromBytes(_ead.O))
	_aga.Set("\u0055", MakeStringFromBytes(_ead.U))
	if _ead.R >= 5 {
		_aga.Set("\u004f\u0045", MakeStringFromBytes(_ead.OE))
		_aga.Set("\u0055\u0045", MakeStringFromBytes(_ead.UE))
		_aga.Set("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", MakeBool(_ead.EncryptMetadata))
		if _ead.R > 5 {
			_aga.Set("\u0050\u0065\u0072m\u0073", MakeStringFromBytes(_ead.Perms))
		}
	}
}

// IsEncrypted checks if the document is encrypted. A bool flag is returned indicating the result.
// First time when called, will check if the Encrypt dictionary is accessible through the trailer dictionary.
// If encrypted, prepares a crypt datastructure which can be used to authenticate and decrypt the document.
// On failure, an error is returned.
func (_daef *PdfParser) IsEncrypted() (bool, error) {
	if _daef._bffd != nil {
		return true, nil
	} else if _daef._dfec == nil {
		return false, nil
	}
	_ae.Log.Trace("\u0043\u0068\u0065c\u006b\u0069\u006e\u0067 \u0065\u006e\u0063\u0072\u0079\u0070\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021")
	_bdce := _daef._dfec.Get("\u0045n\u0063\u0072\u0079\u0070\u0074")
	if _bdce == nil {
		return false, nil
	}
	_ae.Log.Trace("\u0049\u0073\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0021")
	var (
		_bcad *PdfObjectDictionary
	)
	switch _gfbb := _bdce.(type) {
	case *PdfObjectDictionary:
		_bcad = _gfbb
	case *PdfObjectReference:
		_ae.Log.Trace("\u0030\u003a\u0020\u004c\u006f\u006f\u006b\u0020\u0075\u0070\u0020\u0072e\u0066\u0020\u0025\u0071", _gfbb)
		_agag, _gaee := _daef.LookupByReference(*_gfbb)
		_ae.Log.Trace("\u0031\u003a\u0020%\u0071", _agag)
		if _gaee != nil {
			return false, _gaee
		}
		_afad, _dgcfe := _agag.(*PdfIndirectObject)
		if !_dgcfe {
			_ae.Log.Debug("E\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			return false, _d.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_gcda, _dgcfe := _afad.PdfObject.(*PdfObjectDictionary)
		_daef._bcdd = _afad
		_ae.Log.Trace("\u0032\u003a\u0020%\u0071", _gcda)
		if !_dgcfe {
			return false, _d.New("\u0074\u0072a\u0069\u006c\u0065\u0072 \u0045\u006ec\u0072\u0079\u0070\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		_bcad = _gcda
	case *PdfObjectNull:
		_ae.Log.Debug("\u0045\u006e\u0063\u0072\u0079\u0070\u0074 \u0069\u0073\u0020a\u0020\u006e\u0075l\u006c\u0020o\u0062\u006a\u0065\u0063\u0074\u002e \u0046il\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u002e")
		return false, nil
	default:
		return false, _ac.Errorf("u\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065: \u0025\u0054", _gfbb)
	}
	_gdcaa, _gfdb := PdfCryptNewDecrypt(_daef, _bcad, _daef._dfec)
	if _gfdb != nil {
		return false, _gfdb
	}
	for _, _affg := range []string{"\u0045n\u0063\u0072\u0079\u0070\u0074"} {
		_begbg := _daef._dfec.Get(PdfObjectName(_affg))
		if _begbg == nil {
			continue
		}
		switch _fbcb := _begbg.(type) {
		case *PdfObjectReference:
			_gdcaa._dfg[int(_fbcb.ObjectNumber)] = struct{}{}
		case *PdfIndirectObject:
			_gdcaa._dgd[_fbcb] = true
			_gdcaa._dfg[int(_fbcb.ObjectNumber)] = struct{}{}
		}
	}
	_daef._bffd = _gdcaa
	_ae.Log.Trace("\u0043\u0072\u0079\u0070\u0074\u0065\u0072\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0025\u0062", _gdcaa)
	return true, nil
}

// NewMultiEncoder returns a new instance of MultiEncoder.
func NewMultiEncoder() *MultiEncoder {
	_cegg := MultiEncoder{}
	_cegg._bfacg = []StreamEncoder{}
	return &_cegg
}

// GetCrypter returns the PdfCrypt instance which has information about the PDFs encryption.
func (_eeaaa *PdfParser) GetCrypter() *PdfCrypt { return _eeaaa._bffd }

// WriteString outputs the object as it is to be written to file.
func (_cafd *PdfObjectInteger) WriteString() string { return _dg.FormatInt(int64(*_cafd), 10) }
func (_ffedd *JBIG2Encoder) encodeImage(_cfbf _ea.Image) ([]byte, error) {
	const _bcgb = "e\u006e\u0063\u006f\u0064\u0065\u0049\u006d\u0061\u0067\u0065"
	_gfbe, _bcd := GoImageToJBIG2(_cfbf, JB2ImageAutoThreshold)
	if _bcd != nil {
		return nil, _eb.Wrap(_bcd, _bcgb, "\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069m\u0061g\u0065\u0020\u0074\u006f\u0020\u006a\u0062\u0069\u0067\u0032\u0020\u0069\u006d\u0067")
	}
	if _bcd = _ffedd.AddPageImage(_gfbe, &_ffedd.DefaultPageSettings); _bcd != nil {
		return nil, _eb.Wrap(_bcd, _bcgb, "")
	}
	return _ffedd.Encode()
}
func (_gfdg *PdfParser) skipComments() error {
	if _, _eecb := _gfdg.skipSpaces(); _eecb != nil {
		return _eecb
	}
	_fbdc := true
	for {
		_addf, _dcgf := _gfdg._eecea.Peek(1)
		if _dcgf != nil {
			_ae.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _dcgf.Error())
			return _dcgf
		}
		if _fbdc && _addf[0] != '%' {
			return nil
		}
		_fbdc = false
		if (_addf[0] != '\r') && (_addf[0] != '\n') {
			_gfdg._eecea.ReadByte()
		} else {
			break
		}
	}
	return _gfdg.skipComments()
}

var _bfce = []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
