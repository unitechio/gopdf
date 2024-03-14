package core

import (
	_fd "bufio"
	_d "bytes"
	_ag "compress/lzw"
	_ecd "compress/zlib"
	_fff "crypto/md5"
	_fc "crypto/rand"
	_ac "encoding/hex"
	_a "errors"
	_ea "fmt"
	_ff "image"
	_fe "image/color"
	_bg "image/jpeg"
	_gd "io"
	_e "reflect"
	_f "regexp"
	_g "sort"
	_be "strconv"
	_cb "strings"
	_c "sync"
	_eca "time"
	_ec "unicode"

	_df "bitbucket.org/shenghui0779/gopdf/common"
	_gb "bitbucket.org/shenghui0779/gopdf/core/security"
	_gc "bitbucket.org/shenghui0779/gopdf/core/security/crypt"
	_dc "bitbucket.org/shenghui0779/gopdf/internal/ccittfax"
	_cf "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_db "bitbucket.org/shenghui0779/gopdf/internal/jbig2"
	_dg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_gg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder"
	_ecdf "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_dd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_gf "bitbucket.org/shenghui0779/gopdf/internal/strutils"
	_cd "golang.org/x/image/tiff/lzw"
	_aa "golang.org/x/xerrors"
)

// PdfIndirectObject represents the primitive PDF indirect object.
type PdfIndirectObject struct {
	PdfObjectReference
	PdfObject
}

const (
	_ega  = 0
	_bgec = 1
	_dace = 2
	_bfc  = 3
	_agce = 4
)
const _gfb = "\u0053\u0074\u0064C\u0046"

func (_dbcaa *JBIG2Image) toBitmap() (_cgfe *_dg.Bitmap, _egc error) {
	const _aea = "\u004a\u0042\u0049\u00472I\u006d\u0061\u0067\u0065\u002e\u0074\u006f\u0042\u0069\u0074\u006d\u0061\u0070"
	if _dbcaa.Data == nil {
		return nil, _dd.Error(_aea, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _dbcaa.Width == 0 || _dbcaa.Height == 0 {
		return nil, _dd.Error(_aea, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if _dbcaa.HasPadding {
		_cgfe, _egc = _dg.NewWithData(_dbcaa.Width, _dbcaa.Height, _dbcaa.Data)
	} else {
		_cgfe, _egc = _dg.NewWithUnpaddedData(_dbcaa.Width, _dbcaa.Height, _dbcaa.Data)
	}
	if _egc != nil {
		return nil, _dd.Wrap(_egc, _aea, "")
	}
	return _cgfe, nil
}

// NewASCIIHexEncoder makes a new ASCII hex encoder.
func NewASCIIHexEncoder() *ASCIIHexEncoder { _egab := &ASCIIHexEncoder{}; return _egab }

// Resolve resolves a PdfObject to direct object, looking up and resolving references as needed (unlike TraceToDirect).
func (_ged *PdfParser) Resolve(obj PdfObject) (PdfObject, error) {
	_bac, _dbc := obj.(*PdfObjectReference)
	if !_dbc {
		return obj, nil
	}
	_egb := _ged.GetFileOffset()
	defer func() { _ged.SetFileOffset(_egb) }()
	_gcf, _deef := _ged.LookupByReference(*_bac)
	if _deef != nil {
		return nil, _deef
	}
	_gge, _cde := _gcf.(*PdfIndirectObject)
	if !_cde {
		return _gcf, nil
	}
	_gcf = _gge.PdfObject
	_, _dbc = _gcf.(*PdfObjectReference)
	if _dbc {
		return _gge, _a.New("\u006d\u0075lt\u0069\u0020\u0064e\u0070\u0074\u0068\u0020tra\u0063e \u0070\u006f\u0069\u006e\u0074\u0065\u0072 t\u006f\u0020\u0070\u006f\u0069\u006e\u0074e\u0072")
	}
	return _gcf, nil
}

func _cdacc(_dcagg _gd.ReadSeeker, _fdbff int64) (*offsetReader, error) {
	_aebe := &offsetReader{_eacg: _dcagg, _fefc: _fdbff}
	_, _dbd := _aebe.Seek(0, _gd.SeekStart)
	return _aebe, _dbd
}

// DecodeStream decodes a FlateEncoded stream object and give back decoded bytes.
func (_gdaa *FlateEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_df.Log.Trace("\u0046l\u0061t\u0065\u0044\u0065\u0063\u006fd\u0065\u0020s\u0074\u0072\u0065\u0061\u006d")
	_df.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _gdaa.Predictor)
	if _gdaa.BitsPerComponent != 8 {
		return nil, _ea.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003d\u0025\u0064\u0020\u0028\u006f\u006e\u006c\u0079\u0020\u0038\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0029", _gdaa.BitsPerComponent)
	}
	_feea, _ggge := _gdaa.DecodeBytes(streamObj.Stream)
	if _ggge != nil {
		return nil, _ggge
	}
	_feea, _ggge = _gdaa.postDecodePredict(_feea)
	if _ggge != nil {
		return nil, _ggge
	}
	return _feea, nil
}

// DecodeStream decodes a JBIG2 encoded stream and returns the result as a slice of bytes.
func (_gggd *JBIG2Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _gggd.DecodeBytes(streamObj.Stream)
}

// DecodeStream implements ASCII85 stream decoding.
func (_gfef *ASCII85Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _gfef.DecodeBytes(streamObj.Stream)
}

// NewParser creates a new parser for a PDF file via ReadSeeker. Loads the cross reference stream and trailer.
// An error is returned on failure.
func NewParser(rs _gd.ReadSeeker) (*PdfParser, error) {
	_dcbfc := &PdfParser{_abdga: rs, ObjCache: make(objectCache), _dbaad: map[int64]bool{}, _decff: make([]int64, 0), _dgef: make(map[*PdfParser]*PdfParser)}
	_gadc, _bgef, _gdcc := _dcbfc.parsePdfVersion()
	if _gdcc != nil {
		_df.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _gdcc)
		return nil, _gdcc
	}
	_dcbfc._bdbfe.Major = _gadc
	_dcbfc._bdbfe.Minor = _bgef
	if _dcbfc._aagb, _gdcc = _dcbfc.loadXrefs(); _gdcc != nil {
		_df.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _gdcc)
		return nil, _gdcc
	}
	_df.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _dcbfc._aagb)
	_bcfg, _gdcc := _dcbfc.parseLinearizedDictionary()
	if _gdcc != nil {
		return nil, _gdcc
	}
	if _bcfg != nil {
		_dcbfc._cegb, _gdcc = _dcbfc.checkLinearizedInformation(_bcfg)
		if _gdcc != nil {
			return nil, _gdcc
		}
	}
	if len(_dcbfc._ggaf.ObjectMap) == 0 {
		return nil, _ea.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	_dcbfc._eaae = len(_dcbfc._decff)
	if _dcbfc._cegb && _dcbfc._eaae != 0 {
		_dcbfc._eaae--
	}
	_dcbfc._bddb = make([]*PdfParser, _dcbfc._eaae)
	return _dcbfc, nil
}

var _gedee = _f.MustCompile("\u005b\\\u0072\u005c\u006e\u005d\u005c\u0073\u002a\u0028\u0078\u0072\u0065f\u0029\u005c\u0073\u002a\u005b\u005c\u0072\u005c\u006e\u005d")

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_gceg *FlateEncoder) MakeDecodeParams() PdfObject {
	if _gceg.Predictor > 1 {
		_eafg := MakeDict()
		_eafg.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_gceg.Predictor)))
		if _gceg.BitsPerComponent != 8 {
			_eafg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_gceg.BitsPerComponent)))
		}
		if _gceg.Columns != 1 {
			_eafg.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_gceg.Columns)))
		}
		if _gceg.Colors != 1 {
			_eafg.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_gceg.Colors)))
		}
		return _eafg
	}
	return nil
}

var (
	_gbaad = []PdfObjectName{"\u0056", "\u0052", "\u004f", "\u0055", "\u0050"}
	_aefgf = _f.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064+\u0029\u005c\u0073\u002b\u0028\u005b\u006e\u0066\u005d\u0029\\\u0073\u002a\u0024")
)

type objectStreams map[int]objectStream

// DecodeStream decodes a DCT encoded stream and returns the result as a
// slice of bytes.
func (_gdb *DCTEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _gdb.DecodeBytes(streamObj.Stream)
}

// GetXrefType returns the type of the first xref object (table or stream).
func (_fefae *PdfParser) GetXrefType() *xrefType { return _fefae._dfbg }

// String returns a string describing `d`.
func (_aecbe *PdfObjectDictionary) String() string {
	var _aegb _cb.Builder
	_aegb.WriteString("\u0044\u0069\u0063t\u0028")
	for _, _gecfa := range _aecbe._aggf {
		_agbd := _aecbe._ccfa[_gecfa]
		_aegb.WriteString("\u0022" + _gecfa.String() + "\u0022\u003a\u0020")
		_aegb.WriteString(_agbd.String())
		_aegb.WriteString("\u002c\u0020")
	}
	_aegb.WriteString("\u0029")
	return _aegb.String()
}

// GetNumberAsInt64 returns the contents of `obj` as an int64 if it is an integer or float, or an
// error if it isn't. This is for cases where expecting an integer, but some implementations
// actually store the number in a floating point format.
func GetNumberAsInt64(obj PdfObject) (int64, error) {
	switch _gdab := obj.(type) {
	case *PdfObjectFloat:
		_df.Log.Debug("\u004e\u0075m\u0062\u0065\u0072\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u0073\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0077\u0061s\u0020\u0073\u0074\u006f\u0072\u0065\u0064\u0020\u0061\u0073\u0020\u0066\u006c\u006fa\u0074\u0020(\u0074\u0079\u0070\u0065 \u0063\u0061\u0073\u0074\u0069n\u0067\u0020\u0075\u0073\u0065\u0064\u0029")
		return int64(*_gdab), nil
	case *PdfObjectInteger:
		return int64(*_gdab), nil
	case *PdfObjectReference:
		_ggdf := TraceToDirectObject(obj)
		return GetNumberAsInt64(_ggdf)
	case *PdfIndirectObject:
		return GetNumberAsInt64(_gdab.PdfObject)
	}
	return 0, ErrNotANumber
}

// Merge merges in key/values from another dictionary. Overwriting if has same keys.
// The mutated dictionary (d) is returned in order to allow method chaining.
func (_adbeca *PdfObjectDictionary) Merge(another *PdfObjectDictionary) *PdfObjectDictionary {
	if another != nil {
		for _, _bcbdc := range another.Keys() {
			_faad := another.Get(_bcbdc)
			_adbeca.Set(_bcbdc, _faad)
		}
	}
	return _adbeca
}
func _dafa() string { return _df.Version }

// HasInvalidHexRunes implements core.ParserMetadata interface.
func (_ggec ParserMetadata) HasInvalidHexRunes() bool { return _ggec._feed }

// GetNumberAsFloat returns the contents of `obj` as a float if it is an integer or float, or an
// error if it isn't.
func GetNumberAsFloat(obj PdfObject) (float64, error) {
	switch _aaea := obj.(type) {
	case *PdfObjectFloat:
		return float64(*_aaea), nil
	case *PdfObjectInteger:
		return float64(*_aaea), nil
	case *PdfObjectReference:
		_ggbeg := TraceToDirectObject(obj)
		return GetNumberAsFloat(_ggbeg)
	case *PdfIndirectObject:
		return GetNumberAsFloat(_aaea.PdfObject)
	}
	return 0, ErrNotANumber
}

var _dfff = _f.MustCompile("\u005e\u005b\\\u002b\u002d\u002e\u005d*\u0028\u005b0\u002d\u0039\u002e\u005d\u002b\u0029\u005b\u0065E\u005d\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func (_adaa *PdfParser) loadXrefs() (*PdfObjectDictionary, error) {
	_adaa._ggaf.ObjectMap = make(map[int]XrefObject)
	_adaa._aeede = make(objectStreams)
	_adcf, _addb := _adaa._abdga.Seek(0, _gd.SeekEnd)
	if _addb != nil {
		return nil, _addb
	}
	_df.Log.Trace("\u0066s\u0069\u007a\u0065\u003a\u0020\u0025d", _adcf)
	_adaa._gccgc = _adcf
	_addb = _adaa.seekToEOFMarker(_adcf)
	if _addb != nil {
		_df.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0073\u0065\u0065\u006b\u0020\u0074\u006f\u0020\u0065\u006f\u0066\u0020\u006d\u0061\u0072\u006b\u0065\u0072: \u0025\u0076", _addb)
		return nil, _addb
	}
	_efbc, _addb := _adaa._abdga.Seek(0, _gd.SeekCurrent)
	if _addb != nil {
		return nil, _addb
	}
	var _ecaa int64 = 64
	_cffg := _efbc - _ecaa
	if _cffg < 0 {
		_cffg = 0
	}
	_, _addb = _adaa._abdga.Seek(_cffg, _gd.SeekStart)
	if _addb != nil {
		return nil, _addb
	}
	_dgfege := make([]byte, _ecaa)
	_, _addb = _adaa._abdga.Read(_dgfege)
	if _addb != nil {
		_df.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u003a\u0020\u0025\u0076", _addb)
		return nil, _addb
	}
	_efbgd := _gagg.FindStringSubmatch(string(_dgfege))
	if len(_efbgd) < 2 {
		_df.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020n\u006f\u0074\u0020f\u006fu\u006e\u0064\u0021")
		return nil, _a.New("\u0073\u0074\u0061\u0072tx\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	if len(_efbgd) > 2 {
		_df.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u004du\u006c\u0074\u0069\u0070\u006c\u0065\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029\u0021", _dgfege)
		return nil, _a.New("m\u0075\u006c\u0074\u0069\u0070\u006ce\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065f\u0020\u0065\u006et\u0072i\u0065\u0073\u003f")
	}
	_gece, _ := _be.ParseInt(_efbgd[1], 10, 64)
	_df.Log.Trace("\u0073t\u0061r\u0074\u0078\u0072\u0065\u0066\u0020\u0061\u0074\u0020\u0025\u0064", _gece)
	if _gece > _adcf {
		_df.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u006f\u0066f\u0073e\u0074 \u006fu\u0074\u0073\u0069\u0064\u0065\u0020\u006f\u0066\u0020\u0066\u0069\u006c\u0065")
		_df.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0072e\u0070\u0061\u0069\u0072")
		_gece, _addb = _adaa.repairLocateXref()
		if _addb != nil {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0065\u0070\u0061\u0069\u0072\u0020\u0061\u0074\u0074\u0065\u006d\u0070t\u0020\u0066\u0061\u0069\u006c\u0065\u0064 \u0028\u0025\u0073\u0029")
			return nil, _addb
		}
	}
	_adaa._abdga.Seek(_gece, _gd.SeekStart)
	_adaa._gcec = _fd.NewReader(_adaa._abdga)
	_ffee, _addb := _adaa.parseXref()
	if _addb != nil {
		return nil, _addb
	}
	_aefc := _ffee.Get("\u0058R\u0065\u0066\u0053\u0074\u006d")
	if _aefc != nil {
		_ecdb, _cgbd := _aefc.(*PdfObjectInteger)
		if !_cgbd {
			return nil, _a.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0021=\u0020\u0069\u006e\u0074")
		}
		_, _addb = _adaa.parseXrefStream(_ecdb)
		if _addb != nil {
			return nil, _addb
		}
	}
	var _gaff []int64
	_cfaac := func(_egeb int64, _adg []int64) bool {
		for _, _afcb := range _adg {
			if _afcb == _egeb {
				return true
			}
		}
		return false
	}
	_aefc = _ffee.Get("\u0050\u0072\u0065\u0076")
	for _aefc != nil {
		_gbgf, _fdcfb := _aefc.(*PdfObjectInteger)
		if !_fdcfb {
			_df.Log.Debug("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020P\u0072\u0065\u0076\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _aefc)
			return _ffee, nil
		}
		_ddafa := *_gbgf
		_df.Log.Trace("\u0041\u006eot\u0068\u0065\u0072 \u0050\u0072\u0065\u0076 xr\u0065f \u0074\u0061\u0062\u006c\u0065\u0020\u006fbj\u0065\u0063\u0074\u0020\u0061\u0074\u0020%\u0064", _ddafa)
		_adaa._abdga.Seek(int64(_ddafa), _gd.SeekStart)
		_adaa._gcec = _fd.NewReader(_adaa._abdga)
		_cdfg, _gdbed := _adaa.parseXref()
		if _gdbed != nil {
			_df.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0045\u0072\u0072\u006f\u0072\u0020-\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069n\u0067\u0020\u0061\u006e\u006f\u0074\u0068\u0065\u0072\u0020\u0028\u0050re\u0076\u0029\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
			_df.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0063\u006f\u006e\u0074\u0069n\u0075\u0065\u0020\u0062\u0079\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067\u0020\u0069\u0074")
			break
		}
		_adaa._decff = append(_adaa._decff, int64(_ddafa))
		_aefc = _cdfg.Get("\u0050\u0072\u0065\u0076")
		if _aefc != nil {
			_cgbde := *(_aefc.(*PdfObjectInteger))
			if _cfaac(int64(_cgbde), _gaff) {
				_df.Log.Debug("\u0050\u0072ev\u0065\u006e\u0074i\u006e\u0067\u0020\u0063irc\u0075la\u0072\u0020\u0078\u0072\u0065\u0066\u0020re\u0066\u0065\u0072\u0065\u006e\u0063\u0069n\u0067")
				break
			}
			_gaff = append(_gaff, int64(_cgbde))
		}
	}
	return _ffee, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_ggeg *DCTEncoder) MakeDecodeParams() PdfObject { return nil }

// MakeFloat creates an PdfObjectFloat from a float64.
func MakeFloat(val float64) *PdfObjectFloat { _bcbaf := PdfObjectFloat(val); return &_bcbaf }

// HasInvalidSubsectionHeader implements core.ParserMetadata interface.
func (_cbe ParserMetadata) HasInvalidSubsectionHeader() bool { return _cbe._bfa }

// ParserMetadata is the parser based metadata information about document.
// The data here could be used on document verification.
type ParserMetadata struct {
	_ecef  int
	_eddgg bool
	_cdef  [4]byte
	_dacf  bool
	_faae  bool
	_feed  bool
	_ecg   bool
	_bfa   bool
	_fdg   bool
}

// Decrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_ada *PdfCrypt) Decrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _ada.isDecrypted(obj) {
		return nil
	}
	switch _fgab := obj.(type) {
	case *PdfIndirectObject:
		_ada._gee[_fgab] = true
		_df.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _fgab.ObjectNumber, _fgab.GenerationNumber)
		_fgb := _fgab.ObjectNumber
		_cffc := _fgab.GenerationNumber
		_dge := _ada.Decrypt(_fgab.PdfObject, _fgb, _cffc)
		if _dge != nil {
			return _dge
		}
		return nil
	case *PdfObjectStream:
		_ada._gee[_fgab] = true
		_ggbg := _fgab.PdfObjectDictionary
		if _ada._bga.R != 5 {
			if _ede, _cag := _ggbg.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _cag && *_ede == "\u0058\u0052\u0065\u0066" {
				return nil
			}
		}
		_ffefd := _fgab.ObjectNumber
		_gea := _fgab.GenerationNumber
		_df.Log.Trace("\u0044e\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _ffefd, _gea)
		_egbb := _gfb
		if _ada._dgc.V >= 4 {
			_egbb = _ada._ddfb
			_df.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _ada._ddfb)
			if _gcaf, _fgac := _ggbg.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _fgac {
				if _aaf, _faf := GetName(_gcaf.Get(0)); _faf {
					if *_aaf == "\u0043\u0072\u0079p\u0074" {
						_egbb = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _gdc, _dfdb := _ggbg.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _dfdb {
							if _febc, _fdeg := _gdc.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _fdeg {
								if _, _adea := _ada._bde[string(*_febc)]; _adea {
									_df.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_febc)
									_egbb = string(*_febc)
								}
							}
						}
					}
				}
			}
			_df.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _egbb)
			if _egbb == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_dda := _ada.Decrypt(_ggbg, _ffefd, _gea)
		if _dda != nil {
			return _dda
		}
		_bbb, _dda := _ada.makeKey(_egbb, uint32(_ffefd), uint32(_gea), _ada._edb)
		if _dda != nil {
			return _dda
		}
		_fgab.Stream, _dda = _ada.decryptBytes(_fgab.Stream, _egbb, _bbb)
		if _dda != nil {
			return _dda
		}
		_ggbg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_fgab.Stream))))
		return nil
	case *PdfObjectString:
		_df.Log.Trace("\u0044e\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_bfeg := _gfb
		if _ada._dgc.V >= 4 {
			_df.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _ada._fbe)
			if _ada._fbe == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_bfeg = _ada._fbe
		}
		_egd, _bggd := _ada.makeKey(_bfeg, uint32(parentObjNum), uint32(parentGenNum), _ada._edb)
		if _bggd != nil {
			return _bggd
		}
		_gcg := _fgab.Str()
		_bcb := make([]byte, len(_gcg))
		for _aafc := 0; _aafc < len(_gcg); _aafc++ {
			_bcb[_aafc] = _gcg[_aafc]
		}
		if len(_bcb) > 0 {
			_df.Log.Trace("\u0044e\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _bcb, _bcb)
			_bcb, _bggd = _ada.decryptBytes(_bcb, _bfeg, _egd)
			if _bggd != nil {
				return _bggd
			}
		}
		_fgab._bcfef = string(_bcb)
		return nil
	case *PdfObjectArray:
		for _, _efe := range _fgab.Elements() {
			_dagg := _ada.Decrypt(_efe, parentObjNum, parentGenNum)
			if _dagg != nil {
				return _dagg
			}
		}
		return nil
	case *PdfObjectDictionary:
		_fcb := false
		if _dcec := _fgab.Get("\u0054\u0079\u0070\u0065"); _dcec != nil {
			_bbae, _adab := _dcec.(*PdfObjectName)
			if _adab && *_bbae == "\u0053\u0069\u0067" {
				_fcb = true
			}
		}
		for _, _dbgf := range _fgab.Keys() {
			_fagd := _fgab.Get(_dbgf)
			if _fcb && string(_dbgf) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_dbgf) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_dbgf) != "\u0050\u0072\u0065\u0076" && string(_dbgf) != "\u004c\u0061\u0073\u0074" {
				_adfe := _ada.Decrypt(_fagd, parentObjNum, parentGenNum)
				if _adfe != nil {
					return _adfe
				}
			}
		}
		return nil
	}
	return nil
}

func (_gedg *limitedReadSeeker) getError(_cgefe int64) error {
	switch {
	case _cgefe < 0:
		return _ea.Errorf("\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0065\u0067\u0061\u0074\u0069\u0076e\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0064", _cgefe)
	case _cgefe > _gedg._dgfg:
		return _ea.Errorf("u\u006e\u0065\u0078\u0070ec\u0074e\u0064\u0020\u006f\u0066\u0066s\u0065\u0074\u003a\u0020\u0025\u0064", _cgefe)
	}
	return nil
}

// DecodeStream decodes the stream data and returns the decoded data.
// An error is returned upon failure.
func DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_df.Log.Trace("\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_cbcg, _fbec := NewEncoderFromStream(streamObj)
	if _fbec != nil {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fbec)
		return nil, _fbec
	}
	_df.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076\u000a", _cbcg)
	_fbfe, _fbec := _cbcg.DecodeStream(streamObj)
	if _fbec != nil {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fbec)
		return nil, _fbec
	}
	return _fbfe, nil
}

func _daee(_dfdf *PdfObjectStream) (*MultiEncoder, error) {
	_aggg := NewMultiEncoder()
	_egfb := _dfdf.PdfObjectDictionary
	if _egfb == nil {
		return _aggg, nil
	}
	var _egga *PdfObjectDictionary
	var _dgff []PdfObject
	_fefb := _egfb.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
	if _fefb != nil {
		_fbad, _bbde := _fefb.(*PdfObjectDictionary)
		if _bbde {
			_egga = _fbad
		}
		_ecac, _dffe := _fefb.(*PdfObjectArray)
		if _dffe {
			for _, _fadb := range _ecac.Elements() {
				_fadb = TraceToDirectObject(_fadb)
				if _dcgb, _gcdd := _fadb.(*PdfObjectDictionary); _gcdd {
					_dgff = append(_dgff, _dcgb)
				} else {
					_dgff = append(_dgff, MakeDict())
				}
			}
		}
	}
	_fefb = _egfb.Get("\u0046\u0069\u006c\u0074\u0065\u0072")
	if _fefb == nil {
		return nil, _ea.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_abbb, _eabac := _fefb.(*PdfObjectArray)
	if !_eabac {
		return nil, _ea.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _eccdb, _feba := range _abbb.Elements() {
		_beaf, _gefa := _feba.(*PdfObjectName)
		if !_gefa {
			return nil, _ea.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _eafa PdfObject
		if _egga != nil {
			_eafa = _egga
		} else {
			if len(_dgff) > 0 {
				if _eccdb >= len(_dgff) {
					return nil, _ea.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_eafa = _dgff[_eccdb]
			}
		}
		var _bfeb *PdfObjectDictionary
		if _bded, _ddcge := _eafa.(*PdfObjectDictionary); _ddcge {
			_bfeb = _bded
		}
		_df.Log.Trace("\u004e\u0065\u0078t \u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073\u002c \u0064p\u003a \u0025v\u002c\u0020\u0064\u0050\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u0076", *_beaf, _eafa, _bfeb)
		if *_beaf == StreamEncodingFilterNameFlate {
			_cege, _gdgdb := _daca(_dfdf, _bfeb)
			if _gdgdb != nil {
				return nil, _gdgdb
			}
			_aggg.AddEncoder(_cege)
		} else if *_beaf == StreamEncodingFilterNameLZW {
			_efbg, _abdc := _bdbf(_dfdf, _bfeb)
			if _abdc != nil {
				return nil, _abdc
			}
			_aggg.AddEncoder(_efbg)
		} else if *_beaf == StreamEncodingFilterNameASCIIHex {
			_dbcb := NewASCIIHexEncoder()
			_aggg.AddEncoder(_dbcb)
		} else if *_beaf == StreamEncodingFilterNameASCII85 {
			_gbaaab := NewASCII85Encoder()
			_aggg.AddEncoder(_gbaaab)
		} else if *_beaf == StreamEncodingFilterNameDCT {
			_beee, _gegb := _defa(_dfdf, _aggg)
			if _gegb != nil {
				return nil, _gegb
			}
			_aggg.AddEncoder(_beee)
			_df.Log.Trace("A\u0064d\u0065\u0064\u0020\u0044\u0043\u0054\u0020\u0065n\u0063\u006f\u0064\u0065r.\u002e\u002e")
			_df.Log.Trace("\u004du\u006ct\u0069\u0020\u0065\u006e\u0063o\u0064\u0065r\u003a\u0020\u0025\u0023\u0076", _aggg)
		} else if *_beaf == StreamEncodingFilterNameCCITTFax {
			_dfee, _dggg := _dbad(_dfdf, _bfeb)
			if _dggg != nil {
				return nil, _dggg
			}
			_aggg.AddEncoder(_dfee)
		} else {
			_df.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_beaf)
			return nil, _ea.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _aggg, nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_ege *FlateEncoder) MakeStreamDict() *PdfObjectDictionary {
	_abde := MakeDict()
	_abde.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_ege.GetFilterName()))
	_ecag := _ege.MakeDecodeParams()
	if _ecag != nil {
		_abde.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _ecag)
	}
	return _abde
}

// GetFilterName returns the name of the encoding filter.
func (_cdgf *LZWEncoder) GetFilterName() string { return StreamEncodingFilterNameLZW }

// EqualObjects returns true if `obj1` and `obj2` have the same contents.
//
// NOTE: It is a good idea to flatten obj1 and obj2 with FlattenObject before calling this function
// so that contents, rather than references, can be compared.
func EqualObjects(obj1, obj2 PdfObject) bool { return _edege(obj1, obj2, 0) }

// PdfObjectInteger represents the primitive PDF integer numerical object.
type PdfObjectInteger int64

// Keys returns the list of keys in the dictionary.
// If `d` is nil returns a nil slice.
func (_abee *PdfObjectDictionary) Keys() []PdfObjectName {
	if _abee == nil {
		return nil
	}
	return _abee._aggf
}

// IsDelimiter checks if a character represents a delimiter.
func IsDelimiter(c byte) bool {
	return c == '(' || c == ')' || c == '<' || c == '>' || c == '[' || c == ']' || c == '{' || c == '}' || c == '/' || c == '%'
}

// MakeInteger creates a PdfObjectInteger from an int64.
func MakeInteger(val int64) *PdfObjectInteger { _cgcc := PdfObjectInteger(val); return &_cgcc }

// UpdateParams updates the parameter values of the encoder.
func (_gae *ASCIIHexEncoder) UpdateParams(params *PdfObjectDictionary) {}

func _cecd(_deag, _bbef, _dcbf uint8) uint8 {
	_acebc := int(_dcbf)
	_ccaa := int(_bbef) - _acebc
	_gbcd := int(_deag) - _acebc
	_acebc = _bca(_ccaa + _gbcd)
	_ccaa = _bca(_ccaa)
	_gbcd = _bca(_gbcd)
	if _ccaa <= _gbcd && _ccaa <= _acebc {
		return _deag
	} else if _gbcd <= _acebc {
		return _bbef
	}
	return _dcbf
}

// EncodeImage encodes 'img' golang image.Image into jbig2 encoded bytes document using default encoder settings.
func (_daac *JBIG2Encoder) EncodeImage(img _ff.Image) ([]byte, error) { return _daac.encodeImage(img) }

// IsFloatDigit checks if a character can be a part of a float number string.
func IsFloatDigit(c byte) bool { return ('0' <= c && c <= '9') || c == '.' }

func (_afc *PdfParser) readComment() (string, error) {
	var _dffb _d.Buffer
	_, _ddaf := _afc.skipSpaces()
	if _ddaf != nil {
		return _dffb.String(), _ddaf
	}
	_fgbb := true
	for {
		_gbda, _dcde := _afc._gcec.Peek(1)
		if _dcde != nil {
			_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _dcde.Error())
			return _dffb.String(), _dcde
		}
		if _fgbb && _gbda[0] != '%' {
			return _dffb.String(), _a.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_fgbb = false
		if (_gbda[0] != '\r') && (_gbda[0] != '\n') {
			_dfcb, _ := _afc._gcec.ReadByte()
			_dffb.WriteByte(_dfcb)
		} else {
			break
		}
	}
	return _dffb.String(), nil
}

// DecodeStream decodes the stream containing CCITTFax encoded image data.
func (_ceac *CCITTFaxEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _ceac.DecodeBytes(streamObj.Stream)
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

func (_efbge *JBIG2Encoder) encodeImage(_aebb _ff.Image) ([]byte, error) {
	const _gfeff = "e\u006e\u0063\u006f\u0064\u0065\u0049\u006d\u0061\u0067\u0065"
	_fdfe, _fbga := GoImageToJBIG2(_aebb, JB2ImageAutoThreshold)
	if _fbga != nil {
		return nil, _dd.Wrap(_fbga, _gfeff, "\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069m\u0061g\u0065\u0020\u0074\u006f\u0020\u006a\u0062\u0069\u0067\u0032\u0020\u0069\u006d\u0067")
	}
	if _fbga = _efbge.AddPageImage(_fdfe, &_efbge.DefaultPageSettings); _fbga != nil {
		return nil, _dd.Wrap(_fbga, _gfeff, "")
	}
	return _efbge.Encode()
}

// Str returns the string value of the PdfObjectString. Defined in addition to String() function to clarify that
// this function returns the underlying string directly, whereas the String function technically could include
// debug info.
func (_edgb *PdfObjectString) Str() string { return _edgb._bcfef }

// UpdateParams updates the parameter values of the encoder.
func (_ffb *DCTEncoder) UpdateParams(params *PdfObjectDictionary) {
	_ecea, _ccgd := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _ccgd == nil {
		_ffb.ColorComponents = int(_ecea)
	}
	_aefa, _ccgd := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _ccgd == nil {
		_ffb.BitsPerComponent = int(_aefa)
	}
	_bdc, _ccgd := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _ccgd == nil {
		_ffb.Width = int(_bdc)
	}
	_dcce, _ccgd := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _ccgd == nil {
		_ffb.Height = int(_dcce)
	}
	_ffde, _ccgd := GetNumberAsInt64(params.Get("\u0051u\u0061\u006c\u0069\u0074\u0079"))
	if _ccgd == nil {
		_ffb.Quality = int(_ffde)
	}
	_bgbb, _gdgf := GetArray(params.Get("\u0044\u0065\u0063\u006f\u0064\u0065"))
	if _gdgf {
		_ffb.Decode, _ccgd = _bgbb.ToFloat64Array()
		if _ccgd != nil {
			_df.Log.Error("F\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020de\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006eto\u0020\u0061r\u0072\u0061\u0079\u0073\u003a\u0020\u0025\u0076", _ccgd)
		}
	}
}

// Resolve resolves the reference and returns the indirect or stream object.
// If the reference cannot be resolved, a *PdfObjectNull object is returned.
func (_gaece *PdfObjectReference) Resolve() PdfObject {
	if _gaece._egcg == nil {
		return MakeNull()
	}
	_dgde, _, _daadb := _gaece._egcg.resolveReference(_gaece)
	if _daadb != nil {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0072\u0065\u0073\u006f\u006cv\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065r\u0065n\u0063\u0065\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074", _daadb)
		return MakeNull()
	}
	if _dgde == nil {
		_df.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0072\u0065\u0073ol\u0076\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065:\u0020\u006ei\u006c\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0061\u0020nu\u006c\u006c\u0020o\u0062\u006a\u0065\u0063\u0074")
		return MakeNull()
	}
	return _dgde
}

// LookupByNumber looks up a PdfObject by object number.  Returns an error on failure.
func (_gdg *PdfParser) LookupByNumber(objNumber int) (PdfObject, error) {
	_eg, _, _abcg := _gdg.lookupByNumberWrapper(objNumber, true)
	return _eg, _abcg
}

// DecodeBytes decodes a slice of Flate encoded bytes and returns the result.
func (_agaa *FlateEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_df.Log.Trace("\u0046\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065\u0020b\u0079\u0074\u0065\u0073")
	if len(encoded) == 0 {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0046\u006c\u0061\u0074\u0065 e\u006ec\u006f\u0064\u0065\u0064\u0020\u0062\u0075\u0066\u0066\u0065\u0072\u002e \u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0062y\u0074\u0065\u0020\u0073\u006c\u0069\u0063\u0065\u002e")
		return []byte{}, nil
	}
	_fbbf := _d.NewReader(encoded)
	_baeb, _degeb := _ecd.NewReader(_fbbf)
	if _degeb != nil {
		_df.Log.Debug("\u0044e\u0063o\u0064\u0069\u006e\u0067\u0020e\u0072\u0072o\u0072\u0020\u0025\u0076\u000a", _degeb)
		_df.Log.Debug("\u0053t\u0072e\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(encoded), encoded)
		return nil, _degeb
	}
	defer _baeb.Close()
	var _fbc _d.Buffer
	_fbc.ReadFrom(_baeb)
	return _fbc.Bytes(), nil
}

// String returns a string describing `ind`.
func (_gfacc *PdfIndirectObject) String() string {
	return _ea.Sprintf("\u0049\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0025\u0064", (*_gfacc).ObjectNumber)
}

// GetUpdatedObjects returns pdf objects which were updated from the specific version (from prevParser).
func (_fceb *PdfParser) GetUpdatedObjects(prevParser *PdfParser) (map[int64]PdfObject, error) {
	if prevParser == nil {
		return nil, _a.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_gebed, _gfbbf := _fceb.getNumbersOfUpdatedObjects(prevParser)
	if _gfbbf != nil {
		return nil, _gfbbf
	}
	_efbe := make(map[int64]PdfObject)
	for _, _dcfe := range _gebed {
		if _begb, _gdbc := _fceb.LookupByNumber(_dcfe); _gdbc == nil {
			_efbe[int64(_dcfe)] = _begb
		} else {
			return nil, _gdbc
		}
	}
	return _efbe, nil
}

// IsAuthenticated returns true if the PDF has already been authenticated for accessing.
func (_ccf *PdfParser) IsAuthenticated() bool { return _ccf._acg._age }

func (_dacad *PdfParser) inspect() (map[string]int, error) {
	_df.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0053P\u0045\u0043\u0054\u0020\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d")
	_df.Log.Trace("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_fbeaa := map[string]int{}
	_ggcb := 0
	_fgce := 0
	var _cgab []int
	for _dfgd := range _dacad._ggaf.ObjectMap {
		_cgab = append(_cgab, _dfgd)
	}
	_g.Ints(_cgab)
	_egag := 0
	for _, _fcfe := range _cgab {
		_cgeac := _dacad._ggaf.ObjectMap[_fcfe]
		if _cgeac.ObjectNumber == 0 {
			continue
		}
		_ggcb++
		_df.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
		_df.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0075\u0070\u0020\u006f\u0062\u006a\u0065c\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0025\u0064", _cgeac.ObjectNumber)
		_gbdgf, _egdae := _dacad.LookupByNumber(_cgeac.ObjectNumber)
		if _egdae != nil {
			_df.Log.Trace("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u006c\u006f\u006f\u006b\u0075p\u0020\u006f\u0062\u006a\u0020\u0025\u0064 \u0028\u0025\u0073\u0029", _cgeac.ObjectNumber, _egdae)
			_fgce++
			continue
		}
		_df.Log.Trace("\u006fb\u006a\u003a\u0020\u0025\u0073", _gbdgf)
		_ffdd, _bccc := _gbdgf.(*PdfIndirectObject)
		if _bccc {
			_df.Log.Trace("\u0049N\u0044 \u004f\u004f\u0042\u004a\u0020\u0025\u0064\u003a\u0020\u0025\u0073", _cgeac.ObjectNumber, _ffdd)
			_cdfbg, _gacc := _ffdd.PdfObject.(*PdfObjectDictionary)
			if _gacc {
				if _gaaca, _gddg := _cdfbg.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _gddg {
					_deca := string(*_gaaca)
					_df.Log.Trace("\u002d\u002d\u002d\u003e\u0020\u004f\u0062\u006a\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _deca)
					_, _bffc := _fbeaa[_deca]
					if _bffc {
						_fbeaa[_deca]++
					} else {
						_fbeaa[_deca] = 1
					}
				} else if _geea, _aceg := _cdfbg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065").(*PdfObjectName); _aceg {
					_fcba := string(*_geea)
					_df.Log.Trace("-\u002d-\u003e\u0020\u004f\u0062\u006a\u0020\u0073\u0075b\u0074\u0079\u0070\u0065: \u0025\u0073", _fcba)
					_, _dcbbg := _fbeaa[_fcba]
					if _dcbbg {
						_fbeaa[_fcba]++
					} else {
						_fbeaa[_fcba] = 1
					}
				}
				if _eaaa, _fbgg := _cdfbg.Get("\u0053").(*PdfObjectName); _fbgg && *_eaaa == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
					_, _bcdeeb := _fbeaa["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]
					if _bcdeeb {
						_fbeaa["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
					} else {
						_fbeaa["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"] = 1
					}
				}
			}
		} else if _eafd, _bebb := _gbdgf.(*PdfObjectStream); _bebb {
			if _gfffg, _efbgeb := _eafd.PdfObjectDictionary.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _efbgeb {
				_df.Log.Trace("\u002d\u002d\u003e\u0020\u0053\u0074\u0072\u0065\u0061\u006d\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073", *_gfffg)
				_aefd := string(*_gfffg)
				_fbeaa[_aefd]++
			}
		} else {
			_degg, _bebba := _gbdgf.(*PdfObjectDictionary)
			if _bebba {
				_fgdea, _bafea := _degg.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
				if _bafea {
					_bdee := string(*_fgdea)
					_df.Log.Trace("\u002d-\u002d \u006f\u0062\u006a\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _bdee)
					_fbeaa[_bdee]++
				}
			}
			_df.Log.Trace("\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u004f\u0042\u004a\u0020\u0025d\u003a\u0020\u0025\u0073", _cgeac.ObjectNumber, _gbdgf)
		}
		_egag++
	}
	_df.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0045\u004fF\u0020\u0049\u004e\u0053\u0050\u0045\u0043T\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d")
	_df.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	_df.Log.Trace("\u004f\u0062j\u0065\u0063\u0074 \u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064", _ggcb)
	_df.Log.Trace("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u006f\u006b\u0075p\u003a\u0020\u0025\u0064", _fgce)
	for _fagce, _agcg := range _fbeaa {
		_df.Log.Trace("\u0025\u0073\u003a\u0020\u0025\u0064", _fagce, _agcg)
	}
	_df.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	if len(_dacad._ggaf.ObjectMap) < 1 {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0068\u0069\u0073 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074 \u0069s\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0021\u0029")
		return nil, _ea.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0028\u0078r\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u0029")
	}
	_dedf, _caeb := _fbeaa["\u0046\u006f\u006e\u0074"]
	if !_caeb || _dedf < 2 {
		_df.Log.Trace("\u0054\u0068\u0069s \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073 \u0070r\u006fb\u0061\u0062\u006c\u0079\u0020\u0073\u0063\u0061\u006e\u006e\u0065\u0064\u0021")
	} else {
		_df.Log.Trace("\u0054\u0068\u0069\u0073\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0066o\u0072\u0020\u0065\u0078\u0074r\u0061\u0063t\u0069\u006f\u006e\u0021")
	}
	return _fbeaa, nil
}

// WriteString outputs the object as it is to be written to file.
func (_cdfe *PdfObjectArray) WriteString() string {
	var _dgbde _cb.Builder
	_dgbde.WriteString("\u005b")
	for _gdbda, _cgbf := range _cdfe.Elements() {
		_dgbde.WriteString(_cgbf.WriteString())
		if _gdbda < (_cdfe.Len() - 1) {
			_dgbde.WriteString("\u0020")
		}
	}
	_dgbde.WriteString("\u005d")
	return _dgbde.String()
}

// NewDCTEncoder makes a new DCT encoder with default parameters.
func NewDCTEncoder() *DCTEncoder {
	_cfe := &DCTEncoder{}
	_cfe.ColorComponents = 3
	_cfe.BitsPerComponent = 8
	_cfe.Quality = DefaultJPEGQuality
	_cfe.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	return _cfe
}

// PdfObjectStream represents the primitive PDF Object stream.
type PdfObjectStream struct {
	PdfObjectReference
	*PdfObjectDictionary
	Stream []byte
}

// EncodeBytes implements support for LZW encoding.  Currently not supporting predictors (raw compressed data only).
// Only supports the Early change = 1 algorithm (compress/lzw) as the other implementation
// does not have a write method.
// TODO: Consider refactoring compress/lzw to allow both.
func (_aab *LZWEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _aab.Predictor != 1 {
		return nil, _ea.Errorf("\u004c\u005aW \u0050\u0072\u0065d\u0069\u0063\u0074\u006fr =\u00201 \u006f\u006e\u006c\u0079\u0020\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0079e\u0074")
	}
	if _aab.EarlyChange == 1 {
		return nil, _ea.Errorf("\u004c\u005a\u0057\u0020\u0045\u0061\u0072\u006c\u0079\u0020\u0043\u0068\u0061n\u0067\u0065\u0020\u003d\u0020\u0030 \u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0079\u0065\u0074")
	}
	var _dde _d.Buffer
	_dbf := _ag.NewWriter(&_dde, _ag.MSB, 8)
	_dbf.Write(data)
	_dbf.Close()
	return _dde.Bytes(), nil
}

// RunLengthEncoder represents Run length encoding.
type RunLengthEncoder struct{}

// ReadAtLeast reads at least n bytes into slice p.
// Returns the number of bytes read (should always be == n), and an error on failure.
func (_ebb *PdfParser) ReadAtLeast(p []byte, n int) (int, error) {
	_agad := n
	_efeff := 0
	_dcac := 0
	for _agad > 0 {
		_aff, _fggg := _ebb._gcec.Read(p[_efeff:])
		if _fggg != nil {
			_df.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _aff, _dcac, _fggg.Error())
			return _efeff, _a.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_dcac++
		_efeff += _aff
		_agad -= _aff
	}
	return _efeff, nil
}

// DecodeStream returns the passed in stream as a slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_cagfd *RawEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return streamObj.Stream, nil
}

// PdfObject is an interface which all primitive PDF objects must implement.
type PdfObject interface {
	// String outputs a string representation of the primitive (for debugging).
	String() string

	// WriteString outputs the PDF primitive as written to file as expected by the standard.
	// TODO(dennwc): it should return a byte slice, or accept a writer
	WriteString() string
}

// SetIfNotNil sets the dictionary's key -> val mapping entry -IF- val is not nil.
// Note that we take care to perform a type switch.  Otherwise if we would supply a nil value
// of another type, e.g. (PdfObjectArray*)(nil), then it would not be a PdfObject(nil) and thus
// would get set.
func (_abaec *PdfObjectDictionary) SetIfNotNil(key PdfObjectName, val PdfObject) {
	if val != nil {
		switch _ggeae := val.(type) {
		case *PdfObjectName:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectDictionary:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectStream:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectString:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectNull:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectInteger:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectArray:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectBool:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectFloat:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfObjectReference:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		case *PdfIndirectObject:
			if _ggeae != nil {
				_abaec.Set(key, val)
			}
		default:
			_df.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", val)
		}
	}
}

// PdfObjectNull represents the primitive PDF null object.
type PdfObjectNull struct{}

func (_cgf *PdfParser) lookupByNumber(_ef int, _dce bool) (PdfObject, bool, error) {
	_eb, _ced := _cgf.ObjCache[_ef]
	if _ced {
		_df.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0063a\u0063\u0068\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0025\u0064", _ef)
		return _eb, false, nil
	}
	if _cgf._dbeed == nil {
		_cgf._dbeed = map[int]bool{}
	}
	if _cgf._dbeed[_ef] {
		_df.Log.Debug("ER\u0052\u004f\u0052\u003a\u0020\u004c\u006fok\u0075\u0070\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0069\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0069\u006e\u0020\u0070\u0072\u006f\u0067\u0072\u0065\u0073\u0073\u0020\u002d\u0020\u0072\u0065c\u0075\u0072\u0073\u0069\u0076\u0065 \u006c\u006f\u006f\u006b\u0075\u0070\u0020\u0061\u0074t\u0065m\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064", _ef)
		return nil, false, _a.New("\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006f\u006f\u006b\u0075p\u0020a\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064")
	}
	_cgf._dbeed[_ef] = true
	defer delete(_cgf._dbeed, _ef)
	_fad, _ced := _cgf._ggaf.ObjectMap[_ef]
	if !_ced {
		_df.Log.Trace("\u0055\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u006c\u006f\u0063\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006e\u0020\u0078\u0072\u0065\u0066\u0073\u0021 \u002d\u0020\u0052\u0065\u0074u\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		var _dcb PdfObjectNull
		return &_dcb, false, nil
	}
	_df.Log.Trace("L\u006fo\u006b\u0075\u0070\u0020\u006f\u0062\u006a\u0020n\u0075\u006d\u0062\u0065r \u0025\u0064", _ef)
	if _fad.XType == XrefTypeTableEntry {
		_df.Log.Trace("\u0078r\u0065f\u006f\u0062\u006a\u0020\u006fb\u006a\u0020n\u0075\u006d\u0020\u0025\u0064", _fad.ObjectNumber)
		_df.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u0067e\u006e\u0020\u0025\u0064", _fad.Generation)
		_df.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0064", _fad.Offset)
		_cgf._abdga.Seek(_fad.Offset, _gd.SeekStart)
		_cgf._gcec = _fd.NewReader(_cgf._abdga)
		_cce, _dfe := _cgf.ParseIndirectObject()
		if _dfe != nil {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006ce\u0064\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029", _dfe)
			if _dce {
				_df.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0072\u0065\u0070\u0061\u0069r\u0020\u0078\u0072\u0065\u0066\u0073\u0020\u0028\u0074\u006f\u0070\u0020d\u006f\u0077\u006e\u0029")
				_fdc, _fgc := _cgf.repairRebuildXrefsTopDown()
				if _fgc != nil {
					_df.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020r\u0065\u0070\u0061\u0069\u0072\u0020\u0028\u0025\u0073\u0029", _fgc)
					return nil, false, _fgc
				}
				_cgf._ggaf = *_fdc
				return _cgf.lookupByNumber(_ef, false)
			}
			return nil, false, _dfe
		}
		if _dce {
			_ad, _, _ := _ffe(_cce)
			if int(_ad) != _ef {
				_df.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0078\u0072\u0065\u0066s\u003a \u0052e\u0062\u0075\u0069\u006c\u0064\u0069\u006eg")
				_fadf := _cgf.rebuildXrefTable()
				if _fadf != nil {
					return nil, false, _fadf
				}
				_cgf.ObjCache = objectCache{}
				return _cgf.lookupByNumberWrapper(_ef, false)
			}
		}
		_df.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006f\u0062\u006a")
		_cgf.ObjCache[_ef] = _cce
		return _cce, false, nil
	} else if _fad.XType == XrefTypeObjectStream {
		_df.Log.Trace("\u0078r\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0021")
		_df.Log.Trace("\u003e\u004c\u006f\u0061\u0064\u0020\u0076\u0069\u0061\u0020\u004f\u0053\u0021")
		_df.Log.Trace("\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0069\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020%\u0064\u002f\u0025\u0064", _fad.OsObjNumber, _fad.OsObjIndex)
		if _fad.OsObjNumber == _ef {
			_df.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0043i\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065f\u0065\u0072\u0065n\u0063e\u0021\u003f\u0021")
			return nil, true, _a.New("\u0078\u0072\u0065f \u0063\u0069\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		if _, _fga := _cgf._ggaf.ObjectMap[_fad.OsObjNumber]; _fga {
			_bf, _ace := _cgf.lookupObjectViaOS(_fad.OsObjNumber, _ef)
			if _ace != nil {
				_df.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069n\u0067\u0020\u0045\u0052\u0052\u0020\u0028\u0025\u0073\u0029", _ace)
				return nil, true, _ace
			}
			_df.Log.Trace("\u003c\u004c\u006f\u0061\u0064\u0065\u0064\u0020\u0076i\u0061\u0020\u004f\u0053")
			_cgf.ObjCache[_ef] = _bf
			if _cgf._acg != nil {
				_cgf._acg._gee[_bf] = true
			}
			return _bf, true, nil
		}
		_df.Log.Debug("\u003f\u003f\u0020\u0042\u0065\u006c\u006f\u006eg\u0073\u0020\u0074o \u0061\u0020\u006e\u006f\u006e\u002dc\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002e.\u002e\u0021")
		return nil, true, _a.New("\u006f\u0073\u0020\u0062\u0065\u006c\u006fn\u0067\u0073\u0020t\u006f\u0020\u0061\u0020n\u006f\u006e\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, false, _a.New("\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0078\u0072\u0065\u0066 \u0074\u0079\u0070\u0065")
}

// UpdateParams updates the parameter values of the encoder.
func (_cfce *JPXEncoder) UpdateParams(params *PdfObjectDictionary) {}

func (_egcf *PdfParser) repairSeekXrefMarker() error {
	_fedg, _fgbcf := _egcf._abdga.Seek(0, _gd.SeekEnd)
	if _fgbcf != nil {
		return _fgbcf
	}
	_cddcgd := _f.MustCompile("\u005cs\u0078\u0072\u0065\u0066\u005c\u0073*")
	var _eeec int64
	var _edggg int64 = 1000
	for _eeec < _fedg {
		if _fedg <= (_edggg + _eeec) {
			_edggg = _fedg - _eeec
		}
		_, _ddbbe := _egcf._abdga.Seek(-_eeec-_edggg, _gd.SeekEnd)
		if _ddbbe != nil {
			return _ddbbe
		}
		_ebgf := make([]byte, _edggg)
		_egcf._abdga.Read(_ebgf)
		_df.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0078\u0072\u0065\u0066\u0020\u003a\u0020\u0022\u0025\u0073\u0022", string(_ebgf))
		_cdebf := _cddcgd.FindAllStringIndex(string(_ebgf), -1)
		if _cdebf != nil {
			_cagea := _cdebf[len(_cdebf)-1]
			_df.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _cdebf)
			_egcf._abdga.Seek(-_eeec-_edggg+int64(_cagea[0]), _gd.SeekEnd)
			_egcf._gcec = _fd.NewReader(_egcf._abdga)
			for {
				_acfe, _abfb := _egcf._gcec.Peek(1)
				if _abfb != nil {
					return _abfb
				}
				_df.Log.Trace("\u0042\u003a\u0020\u0025\u0064\u0020\u0025\u0063", _acfe[0], _acfe[0])
				if !IsWhiteSpace(_acfe[0]) {
					break
				}
				_egcf._gcec.Discard(1)
			}
			return nil
		}
		_df.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_eeec += _edggg
	}
	_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074a\u0062\u006c\u0065\u0020\u006d\u0061r\u006b\u0065\u0072\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u002e")
	return _a.New("\u0078r\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020")
}

// String returns a string describing `ref`.
func (_fgdf *PdfObjectReference) String() string {
	return _ea.Sprintf("\u0052\u0065\u0066\u0028\u0025\u0064\u0020\u0025\u0064\u0029", _fgdf.ObjectNumber, _fgdf.GenerationNumber)
}

// Set sets the PdfObject at index i of the streams. An error is returned if the index is outside bounds.
func (_edff *PdfObjectStreams) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_edff._dgbb) {
		return _a.New("\u004f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_edff._dgbb[i] = obj
	return nil
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

func _daca(_aacd *PdfObjectStream, _aca *PdfObjectDictionary) (*FlateEncoder, error) {
	_gadd := NewFlateEncoder()
	_cabf := _aacd.PdfObjectDictionary
	if _cabf == nil {
		return _gadd, nil
	}
	_gadd._dade = _cfcd(_cabf)
	if _aca == nil {
		_ggbe := TraceToDirectObject(_cabf.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		switch _dgfb := _ggbe.(type) {
		case *PdfObjectArray:
			if _dgfb.Len() != 1 {
				_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020\u0028\u0025\u0064\u0029", _dgfb.Len())
				return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			if _gegd, _eafe := GetDict(_dgfb.Get(0)); _eafe {
				_aca = _gegd
			}
		case *PdfObjectDictionary:
			_aca = _dgfb
		case *PdfObjectNull, nil:
		default:
			_df.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _ggbe)
			return nil, _ea.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _aca == nil {
		return _gadd, nil
	}
	_df.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _aca.String())
	_gfg := _aca.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _gfg == nil {
		_df.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_dba, _daa := _gfg.(*PdfObjectInteger)
		if !_daa {
			_df.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _gfg)
			return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_gadd.Predictor = int(*_dba)
	}
	_gfg = _aca.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _gfg != nil {
		_baf, _gag := _gfg.(*PdfObjectInteger)
		if !_gag {
			_df.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ea.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_gadd.BitsPerComponent = int(*_baf)
	}
	if _gadd.Predictor > 1 {
		_gadd.Columns = 1
		_gfg = _aca.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _gfg != nil {
			_cfff, _dafg := _gfg.(*PdfObjectInteger)
			if !_dafg {
				return nil, _ea.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_gadd.Columns = int(*_cfff)
		}
		_gadd.Colors = 1
		_gfg = _aca.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _gfg != nil {
			_fbb, _gadb := _gfg.(*PdfObjectInteger)
			if !_gadb {
				return nil, _ea.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_gadd.Colors = int(*_fbb)
		}
	}
	return _gadd, nil
}

// GetStream returns the *PdfObjectStream represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetStream(obj PdfObject) (_bdgbd *PdfObjectStream, _fbddf bool) {
	obj = ResolveReference(obj)
	_bdgbd, _fbddf = obj.(*PdfObjectStream)
	return _bdgbd, _fbddf
}

// WriteString outputs the object as it is to be written to file.
func (_dbcag *PdfObjectInteger) WriteString() string { return _be.FormatInt(int64(*_dbcag), 10) }

// PdfCryptNewEncrypt makes the document crypt handler based on a specified crypt filter.
func PdfCryptNewEncrypt(cf _gc.Filter, userPass, ownerPass []byte, perm _gb.Permissions) (*PdfCrypt, *EncryptInfo, error) {
	_ddf := &PdfCrypt{_fag: make(map[PdfObject]bool), _bde: make(cryptFilters), _bga: _gb.StdEncryptDict{P: perm, EncryptMetadata: true}}
	var _bdg Version
	if cf != nil {
		_bef := cf.PDFVersion()
		_bdg.Major, _bdg.Minor = _bef[0], _bef[1]
		V, R := cf.HandlerVersion()
		_ddf._dgc.V = V
		_ddf._bga.R = R
		_ddf._dgc.Length = cf.KeyLength() * 8
	}
	const (
		_cdb = _gfb
	)
	_ddf._bde[_cdb] = cf
	if _ddf._dgc.V >= 4 {
		_ddf._ddfb = _cdb
		_ddf._fbe = _cdb
	}
	_bbd := _ddf.newEncryptDict()
	_ee := _fff.Sum([]byte(_eca.Now().Format(_eca.RFC850)))
	_ffef := string(_ee[:])
	_dcf := make([]byte, 100)
	_fc.Read(_dcf)
	_ee = _fff.Sum(_dcf)
	_fbd := string(_ee[:])
	_df.Log.Trace("\u0052\u0061\u006e\u0064\u006f\u006d\u0020\u0062\u003a\u0020\u0025\u0020\u0078", _dcf)
	_df.Log.Trace("\u0047\u0065\u006e\u0020\u0049\u0064\u0020\u0030\u003a\u0020\u0025\u0020\u0078", _ffef)
	_ddf._geg = _ffef
	_cdeb := _ddf.generateParams(userPass, ownerPass)
	if _cdeb != nil {
		return nil, nil, _cdeb
	}
	_ggd(&_ddf._bga, _bbd)
	if _ddf._dgc.V >= 4 {
		if _ccd := _ddf.saveCryptFilters(_bbd); _ccd != nil {
			return nil, nil, _ccd
		}
	}
	return _ddf, &EncryptInfo{Version: _bdg, Encrypt: _bbd, ID0: _ffef, ID1: _fbd}, nil
}

// MakeArrayFromIntegers64 creates an PdfObjectArray from a slice of int64s, where each array element
// is an PdfObjectInteger.
func MakeArrayFromIntegers64(vals []int64) *PdfObjectArray {
	_dgdf := MakeArray()
	for _, _fdabe := range vals {
		_dgdf.Append(MakeInteger(_fdabe))
	}
	return _dgdf
}

func _edege(_eaaef, _degebc PdfObject, _aece int) bool {
	if _aece > _abedf {
		_df.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _abedf)
		return false
	}
	if _eaaef == nil && _degebc == nil {
		return true
	} else if _eaaef == nil || _degebc == nil {
		return false
	}
	if _e.TypeOf(_eaaef) != _e.TypeOf(_degebc) {
		return false
	}
	switch _gfcb := _eaaef.(type) {
	case *PdfObjectNull, *PdfObjectReference:
		return true
	case *PdfObjectName:
		return *_gfcb == *(_degebc.(*PdfObjectName))
	case *PdfObjectString:
		return *_gfcb == *(_degebc.(*PdfObjectString))
	case *PdfObjectInteger:
		return *_gfcb == *(_degebc.(*PdfObjectInteger))
	case *PdfObjectBool:
		return *_gfcb == *(_degebc.(*PdfObjectBool))
	case *PdfObjectFloat:
		return *_gfcb == *(_degebc.(*PdfObjectFloat))
	case *PdfIndirectObject:
		return _edege(TraceToDirectObject(_eaaef), TraceToDirectObject(_degebc), _aece+1)
	case *PdfObjectArray:
		_ddab := _degebc.(*PdfObjectArray)
		if len((*_gfcb)._cdea) != len((*_ddab)._cdea) {
			return false
		}
		for _eged, _bcbafe := range (*_gfcb)._cdea {
			if !_edege(_bcbafe, (*_ddab)._cdea[_eged], _aece+1) {
				return false
			}
		}
		return true
	case *PdfObjectDictionary:
		_cdeggg := _degebc.(*PdfObjectDictionary)
		_fbac, _cbbg := (*_gfcb)._ccfa, (*_cdeggg)._ccfa
		if len(_fbac) != len(_cbbg) {
			return false
		}
		for _cfdb, _ccee := range _fbac {
			_cbfdc, _ffca := _cbbg[_cfdb]
			if !_ffca || !_edege(_ccee, _cbfdc, _aece+1) {
				return false
			}
		}
		return true
	case *PdfObjectStream:
		_eabaf := _degebc.(*PdfObjectStream)
		return _edege((*_gfcb).PdfObjectDictionary, (*_eabaf).PdfObjectDictionary, _aece+1)
	default:
		_df.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", _eaaef)
	}
	return false
}

func (_bec *PdfParser) lookupObjectViaOS(_dff int, _ab int) (PdfObject, error) {
	var _ge *_d.Reader
	var _dgg objectStream
	var _ce bool
	_dgg, _ce = _bec._aeede[_dff]
	if !_ce {
		_gca, _cdc := _bec.LookupByNumber(_dff)
		if _cdc != nil {
			_df.Log.Debug("\u004d\u0069ss\u0069\u006e\u0067 \u006f\u0062\u006a\u0065ct \u0073tr\u0065\u0061\u006d\u0020\u0077\u0069\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u0020%\u0064", _dff)
			return nil, _cdc
		}
		_fee, _da := _gca.(*PdfObjectStream)
		if !_da {
			return nil, _a.New("i\u006e\u0076\u0061\u006cid\u0020o\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _bec._acg != nil && !_bec._acg.isDecrypted(_fee) {
			return nil, _a.New("\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0064\u0065\u0063r\u0079\u0070\u0074\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d")
		}
		_aaa := _fee.PdfObjectDictionary
		_df.Log.Trace("\u0073o\u0020\u0064\u003a\u0020\u0025\u0073\n", _aaa.String())
		_dfd, _da := _aaa.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
		if !_da {
			_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0061\u006c\u0077\u0061\u0079\u0073\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054\u0079\u0070\u0065")
			return nil, _a.New("\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020T\u0079\u0070\u0065")
		}
		if _cb.ToLower(string(*_dfd)) != "\u006f\u0062\u006a\u0073\u0074\u006d" {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0074\u0079\u0070\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0077\u0061\u0079\u0073 \u0062\u0065\u0020\u004f\u0062\u006a\u0053\u0074\u006d\u0020\u0021")
			return nil, _a.New("\u006f\u0062\u006a\u0065c\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074y\u0070e\u0020\u0021\u003d\u0020\u004f\u0062\u006aS\u0074\u006d")
		}
		N, _da := _aaa.Get("\u004e").(*PdfObjectInteger)
		if !_da {
			return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020i\u006e\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_fa, _da := _aaa.Get("\u0046\u0069\u0072s\u0074").(*PdfObjectInteger)
		if !_da {
			return nil, _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0046\u0069\u0072\u0073\u0074\u0020i\u006e \u0073t\u0072e\u0061\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_df.Log.Trace("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _dfd, *N)
		_abc, _cdc := DecodeStream(_fee)
		if _cdc != nil {
			return nil, _cdc
		}
		_df.Log.Trace("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _abc)
		_acf := _bec.GetFileOffset()
		defer func() { _bec.SetFileOffset(_acf) }()
		_ge = _d.NewReader(_abc)
		_bec._gcec = _fd.NewReader(_ge)
		_df.Log.Trace("\u0050a\u0072s\u0069\u006e\u0067\u0020\u006ff\u0066\u0073e\u0074\u0020\u006d\u0061\u0070")
		_dad := map[int]int64{}
		for _aac := 0; _aac < int(*N); _aac++ {
			_bec.skipSpaces()
			_ecaf, _dbg := _bec.parseNumber()
			if _dbg != nil {
				return nil, _dbg
			}
			_ecf, _abd := _ecaf.(*PdfObjectInteger)
			if !_abd {
				return nil, _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_bec.skipSpaces()
			_ecaf, _dbg = _bec.parseNumber()
			if _dbg != nil {
				return nil, _dbg
			}
			_ffd, _abd := _ecaf.(*PdfObjectInteger)
			if !_abd {
				return nil, _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_df.Log.Trace("\u006f\u0062j\u0020\u0025\u0064 \u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0025\u0064", *_ecf, *_ffd)
			_dad[int(*_ecf)] = int64(*_fa + *_ffd)
		}
		_dgg = objectStream{N: int(*N), _dcc: _abc, _ece: _dad}
		_bec._aeede[_dff] = _dgg
	} else {
		_beca := _bec.GetFileOffset()
		defer func() { _bec.SetFileOffset(_beca) }()
		_ge = _d.NewReader(_dgg._dcc)
		_bec._gcec = _fd.NewReader(_ge)
	}
	_beb := _dgg._ece[_ab]
	_df.Log.Trace("\u0041\u0043\u0054\u0055AL\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u005b\u0025\u0064\u005d\u0020\u003d\u0020%\u0064", _ab, _beb)
	_ge.Seek(_beb, _gd.SeekStart)
	_bec._gcec = _fd.NewReader(_ge)
	_gcc, _ := _bec._gcec.Peek(100)
	_df.Log.Trace("\u004f\u0042\u004a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_gcc))
	_cec, _bb := _bec.parseObject()
	if _bb != nil {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0072\u0065\u0061\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _bb)
		return nil, _bb
	}
	if _cec == nil {
		return nil, _a.New("o\u0062\u006a\u0065\u0063t \u0063a\u006e\u006e\u006f\u0074\u0020b\u0065\u0020\u006e\u0075\u006c\u006c")
	}
	_fda := PdfIndirectObject{}
	_fda.ObjectNumber = int64(_ab)
	_fda.PdfObject = _cec
	_fda._egcg = _bec
	return &_fda, nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_cebe *ASCIIHexEncoder) MakeStreamDict() *PdfObjectDictionary {
	_cdcb := MakeDict()
	_cdcb.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_cebe.GetFilterName()))
	return _cdcb
}

func _cgc(_gccg *_gc.FilterDict, _ebf *PdfObjectDictionary) error {
	if _bgae, _bgf := _ebf.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _bgf {
		if _abdg := string(*_bgae); _abdg != "C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072" {
			_df.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020C\u0046\u0020\u0064ic\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074e\u0072\u0029", _abdg)
		}
	}
	_bge, _eag := _ebf.Get("\u0043\u0046\u004d").(*PdfObjectName)
	if !_eag {
		return _ea.Errorf("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0028\u004e\u006f\u006e\u0065\u0029")
	}
	_gccg.CFM = string(*_bge)
	if _egbd, _cga := _ebf.Get("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et").(*PdfObjectName); _cga {
		_gccg.AuthEvent = _gb.AuthEvent(*_egbd)
	} else {
		_gccg.AuthEvent = _gb.EventDocOpen
	}
	if _ffc, _eee := _ebf.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _eee {
		_gccg.Length = int(*_ffc)
	}
	return nil
}

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

func (_gffe *PdfParser) parseHexString() (*PdfObjectString, error) {
	_gffe._gcec.ReadByte()
	var _cecgd _d.Buffer
	for {
		_eecd, _cgeg := _gffe._gcec.Peek(1)
		if _cgeg != nil {
			return MakeString(""), _cgeg
		}
		if _eecd[0] == '>' {
			_gffe._gcec.ReadByte()
			break
		}
		_dagd, _ := _gffe._gcec.ReadByte()
		if _gffe._dcad {
			if _d.IndexByte(_dfde, _dagd) == -1 {
				_gffe._ffge._feed = true
			}
		}
		if !IsWhiteSpace(_dagd) {
			_cecgd.WriteByte(_dagd)
		}
	}
	if _cecgd.Len()%2 == 1 {
		_gffe._ffge._faae = true
		_cecgd.WriteRune('0')
	}
	_afegc, _ := _ac.DecodeString(_cecgd.String())
	return MakeHexString(string(_afegc)), nil
}
func (_gfdg *PdfParser) parseNumber() (PdfObject, error) { return ParseNumber(_gfdg._gcec) }

// GetEncryptObj returns the PdfIndirectObject which has information about the PDFs encryption details.
func (_ggcg *PdfParser) GetEncryptObj() *PdfIndirectObject { return _ggcg._bgafa }

func _defa(_bad *PdfObjectStream, _gbgb *MultiEncoder) (*DCTEncoder, error) {
	_fcd := NewDCTEncoder()
	_bbg := _bad.PdfObjectDictionary
	if _bbg == nil {
		return _fcd, nil
	}
	_gagf := _bad.Stream
	if _gbgb != nil {
		_faac, _dgdc := _gbgb.DecodeBytes(_gagf)
		if _dgdc != nil {
			return nil, _dgdc
		}
		_gagf = _faac
	}
	_gbbc := _d.NewReader(_gagf)
	_baag, _bgaae := _bg.DecodeConfig(_gbbc)
	if _bgaae != nil {
		_df.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _bgaae)
		return nil, _bgaae
	}
	switch _baag.ColorModel {
	case _fe.RGBAModel:
		_fcd.BitsPerComponent = 8
		_fcd.ColorComponents = 3
		_fcd.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _fe.RGBA64Model:
		_fcd.BitsPerComponent = 16
		_fcd.ColorComponents = 3
		_fcd.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _fe.GrayModel:
		_fcd.BitsPerComponent = 8
		_fcd.ColorComponents = 1
		_fcd.Decode = []float64{0.0, 1.0}
	case _fe.Gray16Model:
		_fcd.BitsPerComponent = 16
		_fcd.ColorComponents = 1
		_fcd.Decode = []float64{0.0, 1.0}
	case _fe.CMYKModel:
		_fcd.BitsPerComponent = 8
		_fcd.ColorComponents = 4
		_fcd.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _fe.YCbCrModel:
		_fcd.BitsPerComponent = 8
		_fcd.ColorComponents = 3
		_fcd.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	default:
		return nil, _a.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_fcd.Width = _baag.Width
	_fcd.Height = _baag.Height
	_df.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _fcd)
	_fcd.Quality = DefaultJPEGQuality
	_eaea, _gde := GetArray(_bbg.Get("\u0044\u0065\u0063\u006f\u0064\u0065"))
	if _gde {
		_geag, _dfea := _eaea.ToFloat64Array()
		if _dfea != nil {
			return _fcd, _dfea
		}
		_fcd.Decode = _geag
	}
	return _fcd, nil
}

func _bdbf(_cfaa *PdfObjectStream, _ead *PdfObjectDictionary) (*LZWEncoder, error) {
	_ebgef := NewLZWEncoder()
	_bgag := _cfaa.PdfObjectDictionary
	if _bgag == nil {
		return _ebgef, nil
	}
	if _ead == nil {
		_gdge := TraceToDirectObject(_bgag.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _gdge != nil {
			if _agcc, _fce := _gdge.(*PdfObjectDictionary); _fce {
				_ead = _agcc
			} else if _dfa, _fcef := _gdge.(*PdfObjectArray); _fcef {
				if _dfa.Len() == 1 {
					if _gcag, _cgfa := GetDict(_dfa.Get(0)); _cgfa {
						_ead = _gcag
					}
				}
			}
			if _ead == nil {
				_df.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _gdge)
				return nil, _ea.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
	}
	_dgfe := _bgag.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _dgfe != nil {
		_eddag, _eea := _dgfe.(*PdfObjectInteger)
		if !_eea {
			_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _dgfe)
			return nil, _ea.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_eddag != 0 && *_eddag != 1 {
			return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_ebgef.EarlyChange = int(*_eddag)
	} else {
		_ebgef.EarlyChange = 1
	}
	if _ead == nil {
		return _ebgef, nil
	}
	if _geafb, _cfca := GetIntVal(_ead.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")); _cfca {
		if _geafb == 0 || _geafb == 1 {
			_ebgef.EarlyChange = _geafb
		} else {
			_df.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020E\u0061\u0072\u006c\u0079\u0043\u0068\u0061n\u0067\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020%\u0064", _geafb)
		}
	}
	_dgfe = _ead.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _dgfe != nil {
		_fdabf, _gbbe := _dgfe.(*PdfObjectInteger)
		if !_gbbe {
			_df.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _dgfe)
			return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_ebgef.Predictor = int(*_fdabf)
	}
	_dgfe = _ead.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _dgfe != nil {
		_bcgf, _cebc := _dgfe.(*PdfObjectInteger)
		if !_cebc {
			_df.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ea.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_ebgef.BitsPerComponent = int(*_bcgf)
	}
	if _ebgef.Predictor > 1 {
		_ebgef.Columns = 1
		_dgfe = _ead.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _dgfe != nil {
			_gfab, _dbec := _dgfe.(*PdfObjectInteger)
			if !_dbec {
				return nil, _ea.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_ebgef.Columns = int(*_gfab)
		}
		_ebgef.Colors = 1
		_dgfe = _ead.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _dgfe != nil {
			_eaba, _cad := _dgfe.(*PdfObjectInteger)
			if !_cad {
				return nil, _ea.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_ebgef.Colors = int(*_eaba)
		}
	}
	_df.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _ead.String())
	return _ebgef, nil
}
func _bca(_gfgc int) int { _edbee := _gfgc >> (_dffc - 1); return (_gfgc ^ _edbee) - _edbee }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_dfcg *MultiEncoder) MakeDecodeParams() PdfObject {
	if len(_dfcg._gbgba) == 0 {
		return nil
	}
	if len(_dfcg._gbgba) == 1 {
		return _dfcg._gbgba[0].MakeDecodeParams()
	}
	_fgcb := MakeArray()
	_gaeg := true
	for _, _dcfg := range _dfcg._gbgba {
		_cggb := _dcfg.MakeDecodeParams()
		if _cggb == nil {
			_fgcb.Append(MakeNull())
		} else {
			_gaeg = false
			_fgcb.Append(_cggb)
		}
	}
	if _gaeg {
		return nil
	}
	return _fgcb
}

// NewJBIG2Encoder creates a new JBIG2Encoder.
func NewJBIG2Encoder() *JBIG2Encoder { return &JBIG2Encoder{_efaa: _ecdf.InitEncodeDocument(false)} }

// Set sets the dictionary's key -> val mapping entry. Overwrites if key already set.
func (_abdgf *PdfObjectDictionary) Set(key PdfObjectName, val PdfObject) {
	_abdgf.setWithLock(key, val, true)
}

// FlateEncoder represents Flate encoding.
type FlateEncoder struct {
	Predictor        int
	BitsPerComponent int

	// For predictors
	Columns int
	Rows    int
	Colors  int
	_dade   *_cf.ImageBase
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_ddd *FlateEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _ddd.Predictor != 1 && _ddd.Predictor != 11 {
		_df.Log.Debug("E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0046\u006c\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065r\u0020P\u0072\u0065\u0064\u0069c\u0074\u006fr\u0020\u003d\u0020\u0031\u002c\u0020\u0031\u0031\u0020\u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		return nil, ErrUnsupportedEncodingParameters
	}
	if _ddd.Predictor == 11 {
		_eage := _ddd.Columns
		_acfb := len(data) / _eage
		if len(data)%_eage != 0 {
			_df.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
			return nil, _a.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
		}
		_cgef := _d.NewBuffer(nil)
		_cgca := make([]byte, _eage)
		for _faef := 0; _faef < _acfb; _faef++ {
			_gabd := data[_eage*_faef : _eage*(_faef+1)]
			_cgca[0] = _gabd[0]
			for _aeee := 1; _aeee < _eage; _aeee++ {
				_cgca[_aeee] = byte(int(_gabd[_aeee]-_gabd[_aeee-1]) % 256)
			}
			_cgef.WriteByte(1)
			_cgef.Write(_cgca)
		}
		data = _cgef.Bytes()
	}
	var _abf _d.Buffer
	_egbf := _ecd.NewWriter(&_abf)
	_egbf.Write(data)
	_egbf.Close()
	return _abf.Bytes(), nil
}

// GetFilterName returns the name of the encoding filter.
func (_bdgf *CCITTFaxEncoder) GetFilterName() string { return StreamEncodingFilterNameCCITTFax }

// DecodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_ddbe *RawEncoder) DecodeBytes(encoded []byte) ([]byte, error) { return encoded, nil }

// Inspect analyzes the document object structure. Returns a map of object types (by name) with the instance count
// as value.
func (_bgcf *PdfParser) Inspect() (map[string]int, error) { return _bgcf.inspect() }

// String returns a string describing `streams`.
func (_fbgfa *PdfObjectStreams) String() string {
	return _ea.Sprintf("\u004f\u0062j\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0025\u0064", _fbgfa.ObjectNumber)
}

func (_dac *PdfCrypt) newEncryptDict() *PdfObjectDictionary {
	_egf := MakeDict()
	_egf.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064"))
	_egf.Set("\u0056", MakeInteger(int64(_dac._dgc.V)))
	_egf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_dac._dgc.Length)))
	return _egf
}

var _cadf = _f.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func (_agf *PdfParser) skipSpaces() (int, error) {
	_cabc := 0
	for {
		_afga, _dfba := _agf._gcec.ReadByte()
		if _dfba != nil {
			return 0, _dfba
		}
		if IsWhiteSpace(_afga) {
			_cabc++
		} else {
			_agf._gcec.UnreadByte()
			break
		}
	}
	return _cabc, nil
}

func (_dfb *PdfCrypt) generateParams(_eaf, _acb []byte) error {
	_eeg := _dfb.securityHandler()
	_fae, _cage := _eeg.GenerateParams(&_dfb._bga, _acb, _eaf)
	if _cage != nil {
		return _cage
	}
	_dfb._edb = _fae
	return nil
}

// EncodeBytes DCT encodes the passed in slice of bytes.
func (_dedb *DCTEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _ccdd _ff.Image
	if _dedb.ColorComponents == 1 && _dedb.BitsPerComponent == 8 {
		_ccdd = &_ff.Gray{Rect: _ff.Rect(0, 0, _dedb.Width, _dedb.Height), Pix: data, Stride: _cf.BytesPerLine(_dedb.Width, _dedb.BitsPerComponent, _dedb.ColorComponents)}
	} else {
		var _bfab error
		_ccdd, _bfab = _cf.NewImage(_dedb.Width, _dedb.Height, _dedb.BitsPerComponent, _dedb.ColorComponents, data, nil, nil)
		if _bfab != nil {
			return nil, _bfab
		}
	}
	_bcfd := _bg.Options{}
	_bcfd.Quality = _dedb.Quality
	var _bdca _d.Buffer
	if _bcfa := _bg.Encode(&_bdca, _ccdd, &_bcfd); _bcfa != nil {
		return nil, _bcfa
	}
	return _bdca.Bytes(), nil
}

// String returns a descriptive information string about the encryption method used.
func (_agc *PdfCrypt) String() string {
	if _agc == nil {
		return ""
	}
	_fde := _agc._dgc.Filter + "\u0020\u002d\u0020"
	if _agc._dgc.V == 0 {
		_fde += "\u0055\u006e\u0064\u006fcu\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0061\u006c\u0067\u006f\u0072\u0069\u0074h\u006d"
	} else if _agc._dgc.V == 1 {
		_fde += "\u0052\u0043\u0034:\u0020\u0034\u0030\u0020\u0062\u0069\u0074\u0073"
	} else if _agc._dgc.V == 2 {
		_fde += _ea.Sprintf("\u0052\u0043\u0034:\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073", _agc._dgc.Length)
	} else if _agc._dgc.V == 3 {
		_fde += "U\u006e\u0070\u0075\u0062li\u0073h\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d"
	} else if _agc._dgc.V >= 4 {
		_fde += _ea.Sprintf("\u0053\u0074r\u0065\u0061\u006d\u0020f\u0069\u006ct\u0065\u0072\u003a\u0020\u0025\u0073\u0020\u002d \u0053\u0074\u0072\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074\u0065r\u003a\u0020\u0025\u0073", _agc._ddfb, _agc._fbe)
		_fde += "\u003b\u0020C\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0073\u003a"
		for _fbf, _cbf := range _agc._bde {
			_fde += _ea.Sprintf("\u0020\u002d\u0020\u0025\u0073\u003a\u0020\u0025\u0073 \u0028\u0025\u0064\u0029", _fbf, _cbf.Name(), _cbf.KeyLength())
		}
	}
	_cdcc := _agc.GetAccessPermissions()
	_fde += _ea.Sprintf("\u0020\u002d\u0020\u0025\u0023\u0076", _cdcc)
	return _fde
}

// MakeString creates an PdfObjectString from a string.
// NOTE: PDF does not use utf-8 string encoding like Go so `s` will often not be a utf-8 encoded
// string.
func MakeString(s string) *PdfObjectString { _egffg := PdfObjectString{_bcfef: s}; return &_egffg }

// UpdateParams updates the parameter values of the encoder.
func (_adde *FlateEncoder) UpdateParams(params *PdfObjectDictionary) {
	_aee, _bgba := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _bgba == nil {
		_adde.Predictor = int(_aee)
	}
	_abeg, _bgba := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _bgba == nil {
		_adde.BitsPerComponent = int(_abeg)
	}
	_ggfb, _bgba := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _bgba == nil {
		_adde.Columns = int(_ggfb)
	}
	_dbcg, _bgba := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _bgba == nil {
		_adde.Colors = int(_dbcg)
	}
}

// GetFilterName returns the name of the encoding filter.
func (_ggfe *ASCII85Encoder) GetFilterName() string { return StreamEncodingFilterNameASCII85 }

func _fcfa(_gefe PdfObject, _aafdc int, _fcbfd map[PdfObject]struct{}) error {
	_df.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072s\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0061 \u0028\u0064\u0065\u0070\u0074\u0068\u0020=\u0020\u0025\u0064\u0029", _aafdc)
	if _, _gffb := _fcbfd[_gefe]; _gffb {
		_df.Log.Trace("-\u0041\u006c\u0072\u0065ad\u0079 \u0074\u0072\u0061\u0076\u0065r\u0073\u0065\u0064\u002e\u002e\u002e")
		return nil
	}
	_fcbfd[_gefe] = struct{}{}
	switch _gbgd := _gefe.(type) {
	case *PdfIndirectObject:
		_aebeg := _gbgd
		_df.Log.Trace("\u0069\u006f\u003a\u0020\u0025\u0073", _aebeg)
		_df.Log.Trace("\u002d\u0020\u0025\u0073", _aebeg.PdfObject)
		return _fcfa(_aebeg.PdfObject, _aafdc+1, _fcbfd)
	case *PdfObjectStream:
		_egcb := _gbgd
		return _fcfa(_egcb.PdfObjectDictionary, _aafdc+1, _fcbfd)
	case *PdfObjectDictionary:
		_ebab := _gbgd
		_df.Log.Trace("\u002d\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _ebab)
		for _, _fecd := range _ebab.Keys() {
			_egbe := _ebab.Get(_fecd)
			if _abeba, _fabb := _egbe.(*PdfObjectReference); _fabb {
				_decd := _abeba.Resolve()
				_ebab.Set(_fecd, _decd)
				_efcff := _fcfa(_decd, _aafdc+1, _fcbfd)
				if _efcff != nil {
					return _efcff
				}
			} else {
				_bfca := _fcfa(_egbe, _aafdc+1, _fcbfd)
				if _bfca != nil {
					return _bfca
				}
			}
		}
		return nil
	case *PdfObjectArray:
		_bbaea := _gbgd
		_df.Log.Trace("-\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _bbaea)
		for _acbd, _fcaa := range _bbaea.Elements() {
			if _adbf, _ecdbd := _fcaa.(*PdfObjectReference); _ecdbd {
				_bfdbd := _adbf.Resolve()
				_bbaea.Set(_acbd, _bfdbd)
				_ddgc := _fcfa(_bfdbd, _aafdc+1, _fcbfd)
				if _ddgc != nil {
					return _ddgc
				}
			} else {
				_fabea := _fcfa(_fcaa, _aafdc+1, _fcbfd)
				if _fabea != nil {
					return _fabea
				}
			}
		}
		return nil
	case *PdfObjectReference:
		_df.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020T\u0072\u0061\u0063\u0069\u006e\u0067\u0020\u0061\u0020r\u0065\u0066\u0065r\u0065n\u0063\u0065\u0021")
		return _a.New("\u0065r\u0072\u006f\u0072\u0020t\u0072\u0061\u0063\u0069\u006eg\u0020a\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	return nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_dcea *RunLengthEncoder) MakeStreamDict() *PdfObjectDictionary {
	_egeg := MakeDict()
	_egeg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_dcea.GetFilterName()))
	return _egeg
}

// GetAccessPermissions returns the PDF access permissions as an AccessPermissions object.
func (_dege *PdfCrypt) GetAccessPermissions() _gb.Permissions { return _dege._bga.P }

type cryptFilters map[string]_gc.Filter

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_eecg *MultiEncoder) MakeStreamDict() *PdfObjectDictionary {
	_dgfc := MakeDict()
	_dgfc.Set("\u0046\u0069\u006c\u0074\u0065\u0072", _eecg.GetFilterArray())
	for _, _cbfe := range _eecg._gbgba {
		_aeed := _cbfe.MakeStreamDict()
		for _, _aag := range _aeed.Keys() {
			_ffdb := _aeed.Get(_aag)
			if _aag != "\u0046\u0069\u006c\u0074\u0065\u0072" && _aag != "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073" {
				_dgfc.Set(_aag, _ffdb)
			}
		}
	}
	_fdbd := _eecg.MakeDecodeParams()
	if _fdbd != nil {
		_dgfc.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _fdbd)
	}
	return _dgfc
}

// GetDict returns the *PdfObjectDictionary represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetDict(obj PdfObject) (_agdb *PdfObjectDictionary, _acae bool) {
	_agdb, _acae = TraceToDirectObject(obj).(*PdfObjectDictionary)
	return _agdb, _acae
}

// GetFilterName returns the names of the underlying encoding filters,
// separated by spaces.
// Note: This is just a string, should not be used in /Filter dictionary entry. Use GetFilterArray for that.
// TODO(v4): Refactor to GetFilter() which can be used for /Filter (either Name or Array), this can be
//
//	renamed to String() as a pretty string to use in debugging etc.
func (_cdab *MultiEncoder) GetFilterName() string {
	_cdac := ""
	for _fcge, _cagfc := range _cdab._gbgba {
		_cdac += _cagfc.GetFilterName()
		if _fcge < len(_cdab._gbgba)-1 {
			_cdac += "\u0020"
		}
	}
	return _cdac
}

// DCTEncoder provides a DCT (JPG) encoding/decoding functionality for images.
type DCTEncoder struct {
	ColorComponents  int
	BitsPerComponent int
	Width            int
	Height           int
	Quality          int
	Decode           []float64
}

func _gcda(_fdcaba *PdfObjectStream, _cggc *PdfObjectDictionary) (*JBIG2Encoder, error) {
	const _fbfa = "\u006ee\u0077\u004a\u0042\u0049G\u0032\u0044\u0065\u0063\u006fd\u0065r\u0046r\u006f\u006d\u0053\u0074\u0072\u0065\u0061m"
	_faaef := NewJBIG2Encoder()
	_eggf := _fdcaba.PdfObjectDictionary
	if _eggf == nil {
		return _faaef, nil
	}
	if _cggc == nil {
		_dgbf := _eggf.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		if _dgbf != nil {
			switch _cfeeg := _dgbf.(type) {
			case *PdfObjectDictionary:
				_cggc = _cfeeg
			case *PdfObjectArray:
				if _cfeeg.Len() == 1 {
					if _ecfb, _febb := GetDict(_cfeeg.Get(0)); _febb {
						_cggc = _ecfb
					}
				}
			default:
				_df.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065P\u0061\u0072\u0061\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0025\u0023\u0076", _dgbf)
				return nil, _dd.Errorf(_fbfa, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050a\u0072m\u0073\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _cfeeg)
			}
		}
	}
	if _cggc == nil {
		return _faaef, nil
	}
	_faaef.UpdateParams(_cggc)
	_dbbb, _bfda := GetStream(_cggc.Get("\u004a\u0042\u0049G\u0032\u0047\u006c\u006f\u0062\u0061\u006c\u0073"))
	if !_bfda {
		return _faaef, nil
	}
	var _aefac error
	_faaef.Globals, _aefac = _db.DecodeGlobals(_dbbb.Stream)
	if _aefac != nil {
		_aefac = _dd.Wrap(_aefac, _fbfa, "\u0063\u006f\u0072\u0072u\u0070\u0074\u0065\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0065n\u0063\u006f\u0064\u0065\u0064\u0020\u0064a\u0074\u0061")
		_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aefac)
		return nil, _aefac
	}
	return _faaef, nil
}

// GetFilterName returns the name of the encoding filter.
func (_efcd *ASCIIHexEncoder) GetFilterName() string { return StreamEncodingFilterNameASCIIHex }

func (_geac *PdfParser) parseLinearizedDictionary() (*PdfObjectDictionary, error) {
	_ceae, _ffgab := _geac._abdga.Seek(0, _gd.SeekEnd)
	if _ffgab != nil {
		return nil, _ffgab
	}
	var _bafg int64
	var _gfcfd int64 = 2048
	for _bafg < _ceae-4 {
		if _ceae <= (_gfcfd + _bafg) {
			_gfcfd = _ceae - _bafg
		}
		_, _gfda := _geac._abdga.Seek(_bafg, _gd.SeekStart)
		if _gfda != nil {
			return nil, _gfda
		}
		_fcdf := make([]byte, _gfcfd)
		_, _gfda = _geac._abdga.Read(_fcdf)
		if _gfda != nil {
			return nil, _gfda
		}
		_df.Log.Trace("\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0066i\u0072\u0073\u0074\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0022\u0025\u0073\u0022", string(_fcdf))
		_fgda := _bfdf.FindAllStringIndex(string(_fcdf), -1)
		if _fgda != nil {
			_dgbdb := _fgda[0]
			_df.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _fgda)
			_, _deeg := _geac._abdga.Seek(int64(_dgbdb[0]), _gd.SeekStart)
			if _deeg != nil {
				return nil, _deeg
			}
			_geac._gcec = _fd.NewReader(_geac._abdga)
			_bcae, _deeg := _geac.ParseIndirectObject()
			if _deeg != nil {
				return nil, nil
			}
			if _eeaa, _ffbc := GetIndirect(_bcae); _ffbc {
				if _aafcc, _gadge := GetDict(_eeaa.PdfObject); _gadge {
					if _ecdc := _aafcc.Get("\u004c\u0069\u006e\u0065\u0061\u0072\u0069\u007a\u0065\u0064"); _ecdc != nil {
						return _aafcc, nil
					}
					return nil, nil
				}
			}
			return nil, nil
		}
		_bafg += _gfcfd - 4
	}
	return nil, _a.New("\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064")
}

// Encode encodes previously prepare jbig2 document and stores it as the byte slice.
func (_eaab *JBIG2Encoder) Encode() (_agac []byte, _ddg error) {
	const _cgcb = "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	if _eaab._efaa == nil {
		return nil, _dd.Errorf(_cgcb, "\u0064\u006f\u0063u\u006d\u0065\u006e\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0064a\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_eaab._efaa.FullHeaders = _eaab.DefaultPageSettings.FileMode
	_agac, _ddg = _eaab._efaa.Encode()
	if _ddg != nil {
		return nil, _dd.Wrap(_ddg, _cgcb, "")
	}
	return _agac, nil
}

// SetPredictor sets the predictor function.  Specify the number of columns per row.
// The columns indicates the number of samples per row.
// Used for grouping data together for compression.
func (_add *FlateEncoder) SetPredictor(columns int) { _add.Predictor = 11; _add.Columns = columns }

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

// DecodeGlobals decodes 'encoded' byte stream and returns their Globally defined segments ('Globals').
func (_fbaf *JBIG2Encoder) DecodeGlobals(encoded []byte) (_db.Globals, error) {
	return _db.DecodeGlobals(encoded)
}

func _ccdddf(_bfaeb, _caga, _bdea int) error {
	if _caga < 0 || _caga > _bfaeb {
		return _a.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0061\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	if _bdea < _caga {
		return _a.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0073\u006c\u0069\u0063e\u0020i\u006ed\u0065\u0078\u0020\u0062\u0020\u003c\u0020a")
	}
	if _bdea > _bfaeb {
		return _a.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0062\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	return nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_aef *LZWEncoder) MakeStreamDict() *PdfObjectDictionary {
	_ffg := MakeDict()
	_ffg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_aef.GetFilterName()))
	_adee := _aef.MakeDecodeParams()
	if _adee != nil {
		_ffg.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _adee)
	}
	_ffg.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(int64(_aef.EarlyChange)))
	return _ffg
}

func (_agb *PdfParser) parseBool() (PdfObjectBool, error) {
	_cadc, _defc := _agb._gcec.Peek(4)
	if _defc != nil {
		return PdfObjectBool(false), _defc
	}
	if (len(_cadc) >= 4) && (string(_cadc[:4]) == "\u0074\u0072\u0075\u0065") {
		_agb._gcec.Discard(4)
		return PdfObjectBool(true), nil
	}
	_cadc, _defc = _agb._gcec.Peek(5)
	if _defc != nil {
		return PdfObjectBool(false), _defc
	}
	if (len(_cadc) >= 5) && (string(_cadc[:5]) == "\u0066\u0061\u006cs\u0065") {
		_agb._gcec.Discard(5)
		return PdfObjectBool(false), nil
	}
	return PdfObjectBool(false), _a.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// DecodeImages decodes the page images from the jbig2 'encoded' data input.
// The jbig2 document may contain multiple pages, thus the function can return multiple
// images. The images order corresponds to the page number.
func (_daggc *JBIG2Encoder) DecodeImages(encoded []byte) ([]_ff.Image, error) {
	const _cbfc = "\u004aB\u0049\u0047\u0032\u0045n\u0063\u006f\u0064\u0065\u0072.\u0044e\u0063o\u0064\u0065\u0049\u006d\u0061\u0067\u0065s"
	_ecgbg, _afge := _gg.Decode(encoded, _gg.Parameters{}, _daggc.Globals.ToDocumentGlobals())
	if _afge != nil {
		return nil, _dd.Wrap(_afge, _cbfc, "")
	}
	_eaca, _afge := _ecgbg.PageNumber()
	if _afge != nil {
		return nil, _dd.Wrap(_afge, _cbfc, "")
	}
	_faba := []_ff.Image{}
	var _agdc _ff.Image
	for _gfd := 1; _gfd <= _eaca; _gfd++ {
		_agdc, _afge = _ecgbg.DecodePageImage(_gfd)
		if _afge != nil {
			return nil, _dd.Wrapf(_afge, _cbfc, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _gfd)
		}
		_faba = append(_faba, _agdc)
	}
	return _faba, nil
}

// MakeStream creates an PdfObjectStream with specified contents and encoding. If encoding is nil, then raw encoding
// will be used (i.e. no encoding applied).
func MakeStream(contents []byte, encoder StreamEncoder) (*PdfObjectStream, error) {
	_ccbff := &PdfObjectStream{}
	if encoder == nil {
		encoder = NewRawEncoder()
	}
	_ccbff.PdfObjectDictionary = encoder.MakeStreamDict()
	_aegc, _dbgeab := encoder.EncodeBytes(contents)
	if _dbgeab != nil {
		return nil, _dbgeab
	}
	_ccbff.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_aegc))))
	_ccbff.Stream = _aegc
	return _ccbff, nil
}

func _eegb(_aebfd int) int {
	if _aebfd < 0 {
		return -_aebfd
	}
	return _aebfd
}

func (_gffd *PdfParser) parseXref() (*PdfObjectDictionary, error) {
	_gffd.skipSpaces()
	const _acgff = 20
	_dbaadb, _ := _gffd._gcec.Peek(_acgff)
	for _dgfeg := 0; _dgfeg < 2; _dgfeg++ {
		if _gffd._cdca == 0 {
			_gffd._cdca = _gffd.GetFileOffset()
		}
		if _bfdf.Match(_dbaadb) {
			_df.Log.Trace("\u0078\u0072e\u0066\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002e\u0020\u0050\u0072\u006f\u0062\u0061\u0062\u006c\u0079\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			_df.Log.Debug("\u0073t\u0061r\u0074\u0069\u006e\u0067\u0020w\u0069\u0074h\u0020\u0022\u0025\u0073\u0022", string(_dbaadb))
			return _gffd.parseXrefStream(nil)
		}
		if _bggfa.Match(_dbaadb) {
			_df.Log.Trace("\u0053\u0074\u0061\u006ed\u0061\u0072\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073e\u0063t\u0069\u006f\u006e\u0020\u0074\u0061\u0062l\u0065\u0021")
			return _gffd.parseXrefTable()
		}
		_dedd := _gffd.GetFileOffset()
		if _gffd._cdca == 0 {
			_gffd._cdca = _dedd
		}
		_gffd.SetFileOffset(_dedd - _acgff)
		defer _gffd.SetFileOffset(_dedd)
		_gebc, _ := _gffd._gcec.Peek(_acgff)
		_dbaadb = append(_gebc, _dbaadb...)
	}
	_df.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u0078\u0072\u0065f\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006fr\u0020\u0073\u0074\u0072\u0065\u0061\u006d.\u0020\u0052\u0065\u0070\u0061i\u0072\u0020\u0061\u0074\u0074e\u006d\u0070\u0074\u0065\u0064\u003a\u0020\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0065\u0061\u0072\u006c\u0069\u0065\u0073\u0074\u0020x\u0072\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u0062\u006f\u0074to\u006d\u002e")
	if _eabad := _gffd.repairSeekXrefMarker(); _eabad != nil {
		_df.Log.Debug("\u0052e\u0070a\u0069\u0072\u0020\u0066\u0061i\u006c\u0065d\u0020\u002d\u0020\u0025\u0076", _eabad)
		return nil, _eabad
	}
	return _gffd.parseXrefTable()
}

var _bfdf = _f.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

func (_fefad *PdfParser) repairRebuildXrefsTopDown() (*XrefTable, error) {
	if _fefad._fbae {
		return nil, _ea.Errorf("\u0072\u0065\u0070\u0061\u0069\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_fefad._fbae = true
	_fefad._abdga.Seek(0, _gd.SeekStart)
	_fefad._gcec = _fd.NewReader(_fefad._abdga)
	_fbcb := 20
	_faddd := make([]byte, _fbcb)
	_eccce := XrefTable{}
	_eccce.ObjectMap = make(map[int]XrefObject)
	for {
		_ccebg, _gdgbc := _fefad._gcec.ReadByte()
		if _gdgbc != nil {
			if _gdgbc == _gd.EOF {
				break
			} else {
				return nil, _gdgbc
			}
		}
		if _ccebg == 'j' && _faddd[_fbcb-1] == 'b' && _faddd[_fbcb-2] == 'o' && IsWhiteSpace(_faddd[_fbcb-3]) {
			_bdaac := _fbcb - 4
			for IsWhiteSpace(_faddd[_bdaac]) && _bdaac > 0 {
				_bdaac--
			}
			if _bdaac == 0 || !IsDecimalDigit(_faddd[_bdaac]) {
				continue
			}
			for IsDecimalDigit(_faddd[_bdaac]) && _bdaac > 0 {
				_bdaac--
			}
			if _bdaac == 0 || !IsWhiteSpace(_faddd[_bdaac]) {
				continue
			}
			for IsWhiteSpace(_faddd[_bdaac]) && _bdaac > 0 {
				_bdaac--
			}
			if _bdaac == 0 || !IsDecimalDigit(_faddd[_bdaac]) {
				continue
			}
			for IsDecimalDigit(_faddd[_bdaac]) && _bdaac > 0 {
				_bdaac--
			}
			if _bdaac == 0 {
				continue
			}
			_ebeg := _fefad.GetFileOffset() - int64(_fbcb-_bdaac)
			_dffdd := append(_faddd[_bdaac+1:], _ccebg)
			_cdae, _deagf, _befdc := _fagc(string(_dffdd))
			if _befdc != nil {
				_df.Log.Debug("\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u003a\u0020\u0025\u0076", _befdc)
				return nil, _befdc
			}
			if _bbaa, _ggcd := _eccce.ObjectMap[_cdae]; !_ggcd || _bbaa.Generation < _deagf {
				_efbf := XrefObject{}
				_efbf.XType = XrefTypeTableEntry
				_efbf.ObjectNumber = _cdae
				_efbf.Generation = _deagf
				_efbf.Offset = _ebeg
				_eccce.ObjectMap[_cdae] = _efbf
			}
		}
		_faddd = append(_faddd[1:_fbcb], _ccebg)
	}
	_fefad._dbeed = nil
	return &_eccce, nil
}

// GetName returns the *PdfObjectName represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetName(obj PdfObject) (_bbaeb *PdfObjectName, _ccag bool) {
	_bbaeb, _ccag = TraceToDirectObject(obj).(*PdfObjectName)
	return _bbaeb, _ccag
}

var _dcef = _a.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_dae *CCITTFaxEncoder) MakeDecodeParams() PdfObject {
	_bfaa := MakeDict()
	_bfaa.Set("\u004b", MakeInteger(int64(_dae.K)))
	_bfaa.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_dae.Columns)))
	if _dae.BlackIs1 {
		_bfaa.Set("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031", MakeBool(_dae.BlackIs1))
	}
	if _dae.EncodedByteAlign {
		_bfaa.Set("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e", MakeBool(_dae.EncodedByteAlign))
	}
	if _dae.EndOfLine && _dae.K >= 0 {
		_bfaa.Set("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee", MakeBool(_dae.EndOfLine))
	}
	if _dae.Rows != 0 && !_dae.EndOfBlock {
		_bfaa.Set("\u0052\u006f\u0077\u0073", MakeInteger(int64(_dae.Rows)))
	}
	if !_dae.EndOfBlock {
		_bfaa.Set("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b", MakeBool(_dae.EndOfBlock))
	}
	if _dae.DamagedRowsBeforeError != 0 {
		_bfaa.Set("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072", MakeInteger(int64(_dae.DamagedRowsBeforeError)))
	}
	return _bfaa
}

// ResolveReferencesDeep recursively traverses through object `o`, looking up and replacing
// references with indirect objects.
// Optionally a map of already deep-resolved objects can be provided via `traversed`. The `traversed` map
// is updated while traversing the objects to avoid traversing same objects multiple times.
func ResolveReferencesDeep(o PdfObject, traversed map[PdfObject]struct{}) error {
	if traversed == nil {
		traversed = map[PdfObject]struct{}{}
	}
	return _fcfa(o, 0, traversed)
}

// UpdateParams updates the parameter values of the encoder.
func (_fgabd *RunLengthEncoder) UpdateParams(params *PdfObjectDictionary) {}

func (_dbb *PdfCrypt) loadCryptFilters(_bgg *PdfObjectDictionary) error {
	_dbb._bde = cryptFilters{}
	_cgcg := _bgg.Get("\u0043\u0046")
	_cgcg = TraceToDirectObject(_cgcg)
	if _af, _fec := _cgcg.(*PdfObjectReference); _fec {
		_ccb, _ageb := _dbb._edd.LookupByReference(*_af)
		if _ageb != nil {
			_df.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0075\u0070\u0020\u0043\u0046\u0020\u0072\u0065\u0066\u0065\u0072en\u0063\u0065")
			return _ageb
		}
		_cgcg = TraceToDirectObject(_ccb)
	}
	_afb, _eaa := _cgcg.(*PdfObjectDictionary)
	if !_eaa {
		_df.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0043\u0046\u002c \u0074\u0079\u0070\u0065: \u0025\u0054", _cgcg)
		return _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0046")
	}
	for _, _gba := range _afb.Keys() {
		_gaa := _afb.Get(_gba)
		if _cgd, _fbdd := _gaa.(*PdfObjectReference); _fbdd {
			_ae, _gbaa := _dbb._edd.LookupByReference(*_cgd)
			if _gbaa != nil {
				_df.Log.Debug("\u0045\u0072ro\u0072\u0020\u006co\u006f\u006b\u0075\u0070 up\u0020di\u0063\u0074\u0069\u006f\u006e\u0061\u0072y \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065")
				return _gbaa
			}
			_gaa = TraceToDirectObject(_ae)
		}
		_cdbe, _edc := _gaa.(*PdfObjectDictionary)
		if !_edc {
			return _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006e \u0043\u0046\u0020\u0028\u006e\u0061\u006d\u0065\u0020\u0025\u0073\u0029\u0020-\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0062\u0075\u0074\u0020\u0025\u0054", _gba, _gaa)
		}
		if _gba == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u002d\u0020\u0043\u0061\u006e\u006e\u006f\u0074\u0020\u006f\u0076\u0065\u0072\u0077r\u0069\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0069d\u0065\u006e\u0074\u0069\u0074\u0079\u0020\u0066\u0069\u006c\u0074\u0065\u0072 \u002d\u0020\u0054\u0072\u0079\u0069n\u0067\u0020\u006ee\u0078\u0074")
			continue
		}
		var _cee _gc.FilterDict
		if _fece := _cgc(&_cee, _cdbe); _fece != nil {
			return _fece
		}
		_bda, _bab := _gc.NewFilter(_cee)
		if _bab != nil {
			return _bab
		}
		_dbb._bde[string(_gba)] = _bda
	}
	_dbb._bde["\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"] = _gc.NewIdentity()
	_dbb._fbe = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _cdgg, _dcbb := _bgg.Get("\u0053\u0074\u0072\u0046").(*PdfObjectName); _dcbb {
		if _, _ddb := _dbb._bde[string(*_cdgg)]; !_ddb {
			return _ea.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u0072\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_cdgg)
		}
		_dbb._fbe = string(*_cdgg)
	}
	_dbb._ddfb = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _bba, _dadc := _bgg.Get("\u0053\u0074\u006d\u0046").(*PdfObjectName); _dadc {
		if _, _gbe := _dbb._bde[string(*_bba)]; !_gbe {
			return _ea.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u006d\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_bba)
		}
		_dbb._ddfb = string(*_bba)
	}
	return nil
}

// Read implementation of Read interface.
func (_deab *limitedReadSeeker) Read(p []byte) (_acc int, _dbaa error) {
	_ddba, _dbaa := _deab._cgfaf.Seek(0, _gd.SeekCurrent)
	if _dbaa != nil {
		return 0, _dbaa
	}
	_fgde := _deab._dgfg - _ddba
	if _fgde == 0 {
		return 0, _gd.EOF
	}
	if _acbf := int64(len(p)); _acbf < _fgde {
		_fgde = _acbf
	}
	_bbfb := make([]byte, _fgde)
	_acc, _dbaa = _deab._cgfaf.Read(_bbfb)
	copy(p, _bbfb)
	return _acc, _dbaa
}

// GetRevision returns PdfParser for the specific version of the Pdf document.
func (_egfbg *PdfParser) GetRevision(revisionNumber int) (*PdfParser, error) {
	_dedg := _egfbg._eaae
	if _dedg == revisionNumber {
		return _egfbg, nil
	}
	if _dedg < revisionNumber {
		return nil, _a.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0072\u0065\u0076\u0069\u0073i\u006fn\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	if _egfbg._bddb[revisionNumber] != nil {
		return _egfbg._bddb[revisionNumber], nil
	}
	_ggad := _egfbg
	for ; _dedg > revisionNumber; _dedg-- {
		_fbda, _cdff := _ggad.GetPreviousRevisionParser()
		if _cdff != nil {
			return nil, _cdff
		}
		_egfbg._bddb[_dedg-1] = _fbda
		_egfbg._dgef[_ggad] = _fbda
		_ggad = _fbda
	}
	return _ggad, nil
}

// GetFileOffset returns the current file offset, accounting for buffered position.
func (_cffd *PdfParser) GetFileOffset() int64 {
	_aedbe, _ := _cffd._abdga.Seek(0, _gd.SeekCurrent)
	_aedbe -= int64(_cffd._gcec.Buffered())
	return _aedbe
}

// IsEncrypted checks if the document is encrypted. A bool flag is returned indicating the result.
// First time when called, will check if the Encrypt dictionary is accessible through the trailer dictionary.
// If encrypted, prepares a crypt datastructure which can be used to authenticate and decrypt the document.
// On failure, an error is returned.
func (_cegd *PdfParser) IsEncrypted() (bool, error) {
	if _cegd._acg != nil {
		return true, nil
	} else if _cegd._aagb == nil {
		return false, nil
	}
	_df.Log.Trace("\u0043\u0068\u0065c\u006b\u0069\u006e\u0067 \u0065\u006e\u0063\u0072\u0079\u0070\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021")
	_dabb := _cegd._aagb.Get("\u0045n\u0063\u0072\u0079\u0070\u0074")
	if _dabb == nil {
		return false, nil
	}
	_df.Log.Trace("\u0049\u0073\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0021")
	var _gfea *PdfObjectDictionary
	switch _ceba := _dabb.(type) {
	case *PdfObjectDictionary:
		_gfea = _ceba
	case *PdfObjectReference:
		_df.Log.Trace("\u0030\u003a\u0020\u004c\u006f\u006f\u006b\u0020\u0075\u0070\u0020\u0072e\u0066\u0020\u0025\u0071", _ceba)
		_cbd, _gbade := _cegd.LookupByReference(*_ceba)
		_df.Log.Trace("\u0031\u003a\u0020%\u0071", _cbd)
		if _gbade != nil {
			return false, _gbade
		}
		_eddc, _ecb := _cbd.(*PdfIndirectObject)
		if !_ecb {
			_df.Log.Debug("E\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			return false, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_bgd, _ecb := _eddc.PdfObject.(*PdfObjectDictionary)
		_cegd._bgafa = _eddc
		_df.Log.Trace("\u0032\u003a\u0020%\u0071", _bgd)
		if !_ecb {
			return false, _a.New("\u0074\u0072a\u0069\u006c\u0065\u0072 \u0045\u006ec\u0072\u0079\u0070\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		_gfea = _bgd
	case *PdfObjectNull:
		_df.Log.Debug("\u0045\u006e\u0063\u0072\u0079\u0070\u0074 \u0069\u0073\u0020a\u0020\u006e\u0075l\u006c\u0020o\u0062\u006a\u0065\u0063\u0074\u002e \u0046il\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u002e")
		return false, nil
	default:
		return false, _ea.Errorf("u\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065: \u0025\u0054", _ceba)
	}
	_dbfea, _fdfg := PdfCryptNewDecrypt(_cegd, _gfea, _cegd._aagb)
	if _fdfg != nil {
		return false, _fdfg
	}
	for _, _ebgefc := range []string{"\u0045n\u0063\u0072\u0079\u0070\u0074"} {
		_gfae := _cegd._aagb.Get(PdfObjectName(_ebgefc))
		if _gfae == nil {
			continue
		}
		switch _abeb := _gfae.(type) {
		case *PdfObjectReference:
			_dbfea._becc[int(_abeb.ObjectNumber)] = struct{}{}
		case *PdfIndirectObject:
			_dbfea._gee[_abeb] = true
			_dbfea._becc[int(_abeb.ObjectNumber)] = struct{}{}
		}
	}
	_cegd._acg = _dbfea
	_df.Log.Trace("\u0043\u0072\u0079\u0070\u0074\u0065\u0072\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0025\u0062", _dbfea)
	return true, nil
}

// PdfObjectArray represents the primitive PDF array object.
type PdfObjectArray struct{ _cdea []PdfObject }

func (_gfcf *PdfParser) getNumbersOfUpdatedObjects(_ebae *PdfParser) ([]int, error) {
	if _ebae == nil {
		return nil, _a.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_cabe := _ebae._gccgc
	_deac := make([]int, 0)
	_decb := make(map[int]interface{})
	_gaeb := make(map[int]int64)
	for _acag, _ccbgf := range _gfcf._ggaf.ObjectMap {
		if _ccbgf.Offset == 0 {
			if _ccbgf.OsObjNumber != 0 {
				if _ccca, _gfgg := _gfcf._ggaf.ObjectMap[_ccbgf.OsObjNumber]; _gfgg {
					_decb[_ccbgf.OsObjNumber] = struct{}{}
					_gaeb[_acag] = _ccca.Offset
				} else {
					return nil, _a.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0078r\u0065\u0066\u0020\u0074ab\u006c\u0065")
				}
			}
		} else {
			_gaeb[_acag] = _ccbgf.Offset
		}
	}
	for _gcfc, _gfbf := range _gaeb {
		if _, _eddad := _decb[_gcfc]; _eddad {
			continue
		}
		if _gfbf > _cabe {
			_deac = append(_deac, _gcfc)
		}
	}
	return _deac, nil
}

const _adeg = 6

type objectStream struct {
	N    int
	_dcc []byte
	_ece map[int]int64
}

var _dfde = []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")

// EncodeBytes ASCII encodes the passed in slice of bytes.
func (_abca *ASCIIHexEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _agg _d.Buffer
	for _, _bgfc := range data {
		_agg.WriteString(_ea.Sprintf("\u0025\u002e\u0032X\u0020", _bgfc))
	}
	_agg.WriteByte('>')
	return _agg.Bytes(), nil
}

// UpdateParams updates the parameter values of the encoder.
func (_dffd *CCITTFaxEncoder) UpdateParams(params *PdfObjectDictionary) {
	if _cfbdf, _beg := GetNumberAsInt64(params.Get("\u004b")); _beg == nil {
		_dffd.K = int(_cfbdf)
	}
	if _bcded, _ffea := GetNumberAsInt64(params.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _ffea == nil {
		_dffd.Columns = int(_bcded)
	} else if _bcded, _ffea = GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068")); _ffea == nil {
		_dffd.Columns = int(_bcded)
	}
	if _gdcf, _gdea := GetNumberAsInt64(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _gdea == nil {
		_dffd.BlackIs1 = _gdcf > 0
	} else {
		if _begd, _fbabd := GetBoolVal(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _fbabd {
			_dffd.BlackIs1 = _begd
		} else {
			if _bgca, _acdgb := GetArray(params.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _acdgb {
				_decc, _gagfd := _bgca.ToIntegerArray()
				if _gagfd == nil {
					_dffd.BlackIs1 = _decc[0] == 1 && _decc[1] == 0
				}
			}
		}
	}
	if _beff, _gfge := GetNumberAsInt64(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _gfge == nil {
		_dffd.EncodedByteAlign = _beff > 0
	} else {
		if _fab, _fefa := GetBoolVal(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _fefa {
			_dffd.EncodedByteAlign = _fab
		}
	}
	if _bff, _ffaf := GetNumberAsInt64(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _ffaf == nil {
		_dffd.EndOfLine = _bff > 0
	} else {
		if _gbdg, _bdd := GetBoolVal(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _bdd {
			_dffd.EndOfLine = _gbdg
		}
	}
	if _bebf, _bbc := GetNumberAsInt64(params.Get("\u0052\u006f\u0077\u0073")); _bbc == nil {
		_dffd.Rows = int(_bebf)
	} else if _bebf, _bbc = GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _bbc == nil {
		_dffd.Rows = int(_bebf)
	}
	if _bcfe, _afd := GetNumberAsInt64(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _afd == nil {
		_dffd.EndOfBlock = _bcfe > 0
	} else {
		if _cbaf, _bcdee := GetBoolVal(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _bcdee {
			_dffd.EndOfBlock = _cbaf
		}
	}
	if _bddf, _daga := GetNumberAsInt64(params.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _daga != nil {
		_dffd.DamagedRowsBeforeError = int(_bddf)
	}
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_acaf *RunLengthEncoder) MakeDecodeParams() PdfObject { return nil }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_bbe *LZWEncoder) MakeDecodeParams() PdfObject {
	if _bbe.Predictor > 1 {
		_afeg := MakeDict()
		_afeg.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_bbe.Predictor)))
		if _bbe.BitsPerComponent != 8 {
			_afeg.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_bbe.BitsPerComponent)))
		}
		if _bbe.Columns != 1 {
			_afeg.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_bbe.Columns)))
		}
		if _bbe.Colors != 1 {
			_afeg.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_bbe.Colors)))
		}
		return _afeg
	}
	return nil
}

// NewRunLengthEncoder makes a new run length encoder
func NewRunLengthEncoder() *RunLengthEncoder { return &RunLengthEncoder{} }

// MakeArray creates an PdfObjectArray from a list of PdfObjects.
func MakeArray(objects ...PdfObject) *PdfObjectArray { return &PdfObjectArray{_cdea: objects} }

// WriteString outputs the object as it is to be written to file.
func (_gbadc *PdfObjectBool) WriteString() string {
	if *_gbadc {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}

// DrawableImage is same as golang image/draw's Image interface that allow drawing images.
type DrawableImage interface {
	ColorModel() _fe.Model
	Bounds() _ff.Rectangle
	At(_fdca, _fbg int) _fe.Color
	Set(_gdbd, _ccbc int, _bbaf _fe.Color)
}

// GetFilterName returns the name of the encoding filter.
func (_cabb *JPXEncoder) GetFilterName() string { return StreamEncodingFilterNameJPX }

func (_agef *PdfCrypt) securityHandler() _gb.StdHandler {
	if _agef._bga.R >= 5 {
		return _gb.NewHandlerR6()
	}
	return _gb.NewHandlerR4(_agef._geg, _agef._dgc.Length)
}

// DecodeBytes decodes a slice of LZW encoded bytes and returns the result.
func (_fecb *LZWEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _efge _d.Buffer
	_bfag := _d.NewReader(encoded)
	var _cgea _gd.ReadCloser
	if _fecb.EarlyChange == 1 {
		_cgea = _cd.NewReader(_bfag, _cd.MSB, 8)
	} else {
		_cgea = _ag.NewReader(_bfag, _ag.MSB, 8)
	}
	defer _cgea.Close()
	if _, _cced := _efge.ReadFrom(_cgea); _cced != nil {
		if _cced != _gd.ErrUnexpectedEOF || _efge.Len() == 0 {
			return nil, _cced
		}
		_df.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u004c\u005a\u0057\u0020\u0064\u0065\u0063\u006f\u0064i\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076\u002e \u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062e \u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _cced)
	}
	return _efge.Bytes(), nil
}

func (_fdega *PdfParser) parseArray() (*PdfObjectArray, error) {
	_cegg := MakeArray()
	_fdega._gcec.ReadByte()
	for {
		_fdega.skipSpaces()
		_gabb, _fdef := _fdega._gcec.Peek(1)
		if _fdef != nil {
			return _cegg, _fdef
		}
		if _gabb[0] == ']' {
			_fdega._gcec.ReadByte()
			break
		}
		_bfg, _fdef := _fdega.parseObject()
		if _fdef != nil {
			return _cegg, _fdef
		}
		_cegg.Append(_bfg)
	}
	return _cegg, nil
}

// EncodeBytes encodes the passed in slice of bytes by passing it through the
// EncodeBytes method of the underlying encoders.
func (_cgb *MultiEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_beaa := data
	var _gagfa error
	for _facb := len(_cgb._gbgba) - 1; _facb >= 0; _facb-- {
		_daeg := _cgb._gbgba[_facb]
		_beaa, _gagfa = _daeg.EncodeBytes(_beaa)
		if _gagfa != nil {
			return nil, _gagfa
		}
	}
	return _beaa, nil
}

// SetFileOffset sets the file to an offset position and resets buffer.
func (_cecg *PdfParser) SetFileOffset(offset int64) {
	if offset < 0 {
		offset = 0
	}
	_cecg._abdga.Seek(offset, _gd.SeekStart)
	_cecg._gcec = _fd.NewReader(_cecg._abdga)
}

// ToInt64Slice returns a slice of all array elements as an int64 slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_afab *PdfObjectArray) ToInt64Slice() ([]int64, error) {
	var _cfcc []int64
	for _, _afdcd := range _afab.Elements() {
		if _fege, _dbaab := _afdcd.(*PdfObjectInteger); _dbaab {
			_cfcc = append(_cfcc, int64(*_fege))
		} else {
			return nil, ErrTypeError
		}
	}
	return _cfcc, nil
}

// MakeHexString creates an PdfObjectString from a string intended for output as a hexadecimal string.
func MakeHexString(s string) *PdfObjectString {
	_cgeae := PdfObjectString{_bcfef: s, _aae: true}
	return &_cgeae
}

// Seek implementation of Seek interface.
func (_feee *limitedReadSeeker) Seek(offset int64, whence int) (int64, error) {
	var _gabcd int64
	switch whence {
	case _gd.SeekStart:
		_gabcd = offset
	case _gd.SeekCurrent:
		_bdaa, _eedd := _feee._cgfaf.Seek(0, _gd.SeekCurrent)
		if _eedd != nil {
			return 0, _eedd
		}
		_gabcd = _bdaa + offset
	case _gd.SeekEnd:
		_gabcd = _feee._dgfg + offset
	}
	if _gcbe := _feee.getError(_gabcd); _gcbe != nil {
		return 0, _gcbe
	}
	if _, _gefb := _feee._cgfaf.Seek(_gabcd, _gd.SeekStart); _gefb != nil {
		return 0, _gefb
	}
	return _gabcd, nil
}

// NewRawEncoder returns a new instace of RawEncoder.
func NewRawEncoder() *RawEncoder { return &RawEncoder{} }

// GetNameVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetNameVal(obj PdfObject) (_gegdg string, _eefe bool) {
	_bceb, _eefe := TraceToDirectObject(obj).(*PdfObjectName)
	if _eefe {
		return string(*_bceb), true
	}
	return
}

var _bggfa = _f.MustCompile("\u005c\u0073\u002a\u0078\u0072\u0065\u0066\u005c\u0073\u002a")

// MakeArrayFromFloats creates an PdfObjectArray from a slice of float64s, where each array element is an
// PdfObjectFloat.
func MakeArrayFromFloats(vals []float64) *PdfObjectArray {
	_ggea := MakeArray()
	for _, _dgaf := range vals {
		_ggea.Append(MakeFloat(_dgaf))
	}
	return _ggea
}

type offsetReader struct {
	_eacg _gd.ReadSeeker
	_fefc int64
}

// ToGoImage converts the JBIG2Image to the golang image.Image.
func (_adec *JBIG2Image) ToGoImage() (_ff.Image, error) {
	const _aafa = "J\u0042I\u0047\u0032\u0049\u006d\u0061\u0067\u0065\u002eT\u006f\u0047\u006f\u0049ma\u0067\u0065"
	if _adec.Data == nil {
		return nil, _dd.Error(_aafa, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _adec.Width == 0 || _adec.Height == 0 {
		return nil, _dd.Error(_aafa, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	_eeacg, _dbgd := _cf.NewImage(_adec.Width, _adec.Height, 1, 1, _adec.Data, nil, nil)
	if _dbgd != nil {
		return nil, _dbgd
	}
	return _eeacg, nil
}

func _fagc(_bgggc string) (int, int, error) {
	_cdgb := _bfdf.FindStringSubmatch(_bgggc)
	if len(_cdgb) < 3 {
		return 0, 0, _a.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_ebbg, _ := _be.Atoi(_cdgb[1])
	_cebcd, _ := _be.Atoi(_cdgb[2])
	return _ebbg, _cebcd, nil
}

// Clear resets the array to an empty state.
func (_abba *PdfObjectArray) Clear() { _abba._cdea = []PdfObject{} }

// DecodeStream decodes a JPX encoded stream and returns the result as a
// slice of bytes.
func (_aeca *JPXEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _aeca.GetFilterName())
	return streamObj.Stream, ErrNoJPXDecode
}

// NewMultiEncoder returns a new instance of MultiEncoder.
func NewMultiEncoder() *MultiEncoder {
	_bfcf := MultiEncoder{}
	_bfcf._gbgba = []StreamEncoder{}
	return &_bfcf
}

func (_gfbb *PdfCrypt) isEncrypted(_dfeg PdfObject) bool {
	_, _gccb := _gfbb._fag[_dfeg]
	if _gccb {
		_df.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0065\u006e\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	_df.Log.Trace("\u004e\u006f\u0074\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}

// Remove removes an element specified by key.
func (_gdaaf *PdfObjectDictionary) Remove(key PdfObjectName) {
	_feggf := -1
	for _aega, _bcad := range _gdaaf._aggf {
		if _bcad == key {
			_feggf = _aega
			break
		}
	}
	if _feggf >= 0 {
		_gdaaf._aggf = append(_gdaaf._aggf[:_feggf], _gdaaf._aggf[_feggf+1:]...)
		delete(_gdaaf._ccfa, key)
	}
}

// Decoded returns the PDFDocEncoding or UTF-16BE decoded string contents.
// UTF-16BE is applied when the first two bytes are 0xFE, 0XFF, otherwise decoding of
// PDFDocEncoding is performed.
func (_gdege *PdfObjectString) Decoded() string {
	if _gdege == nil {
		return ""
	}
	_dgdaf := []byte(_gdege._bcfef)
	if len(_dgdaf) >= 2 && _dgdaf[0] == 0xFE && _dgdaf[1] == 0xFF {
		return _gf.UTF16ToString(_dgdaf[2:])
	}
	return _gf.PDFDocEncodingToString(_dgdaf)
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_bcbg *ASCIIHexEncoder) MakeDecodeParams() PdfObject { return nil }

// DecodeStream implements ASCII hex decoding.
func (_adafd *ASCIIHexEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _adafd.DecodeBytes(streamObj.Stream)
}

// GetPreviousRevisionReadSeeker returns ReadSeeker for the previous version of the Pdf document.
func (_cfcgg *PdfParser) GetPreviousRevisionReadSeeker() (_gd.ReadSeeker, error) {
	if _fbgfc := _cfcgg.seekToEOFMarker(_cfcgg._gccgc - _adeg); _fbgfc != nil {
		return nil, _fbgfc
	}
	_dbeg, _ccbb := _cfcgg._abdga.Seek(0, _gd.SeekCurrent)
	if _ccbb != nil {
		return nil, _ccbb
	}
	_dbeg += _adeg
	return _afbe(_cfcgg._abdga, _dbeg)
}

// DecodeStream decodes a LZW encoded stream and returns the result as a
// slice of bytes.
func (_eggc *LZWEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_df.Log.Trace("\u004c\u005a\u0057 \u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067")
	_df.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _eggc.Predictor)
	_eceda, _afec := _eggc.DecodeBytes(streamObj.Stream)
	if _afec != nil {
		return nil, _afec
	}
	_df.Log.Trace("\u0020\u0049\u004e\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(streamObj.Stream), streamObj.Stream)
	_df.Log.Trace("\u004f\u0055\u0054\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(_eceda), _eceda)
	if _eggc.Predictor > 1 {
		if _eggc.Predictor == 2 {
			_df.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_dbca := _eggc.Columns * _eggc.Colors
			if _dbca < 1 {
				return []byte{}, nil
			}
			_dfg := len(_eceda) / _dbca
			if len(_eceda)%_dbca != 0 {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_eceda), _dbca)
			}
			if _dbca%_eggc.Colors != 0 {
				return nil, _ea.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _dbca, _eggc.Colors)
			}
			if _dbca > len(_eceda) {
				_df.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _dbca, len(_eceda))
				return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_df.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_eceda), _eceda)
			_gdce := _d.NewBuffer(nil)
			for _gdf := 0; _gdf < _dfg; _gdf++ {
				_gdac := _eceda[_dbca*_gdf : _dbca*(_gdf+1)]
				for _ddc := _eggc.Colors; _ddc < _dbca; _ddc++ {
					_gdac[_ddc] = byte(int(_gdac[_ddc]+_gdac[_ddc-_eggc.Colors]) % 256)
				}
				_gdce.Write(_gdac)
			}
			_efb := _gdce.Bytes()
			_df.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_efb), _efb)
			return _efb, nil
		} else if _eggc.Predictor >= 10 && _eggc.Predictor <= 15 {
			_df.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_bbdc := _eggc.Columns*_eggc.Colors + 1
			if _bbdc < 1 {
				return []byte{}, nil
			}
			_ffag := len(_eceda) / _bbdc
			if len(_eceda)%_bbdc != 0 {
				return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_eceda), _bbdc)
			}
			if _bbdc > len(_eceda) {
				_df.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _bbdc, len(_eceda))
				return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_eda := _d.NewBuffer(nil)
			_df.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _eggc.Columns)
			_df.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_eceda), _bbdc, _ffag)
			_fdgb := make([]byte, _bbdc)
			for _fgd := 0; _fgd < _bbdc; _fgd++ {
				_fdgb[_fgd] = 0
			}
			for _deeb := 0; _deeb < _ffag; _deeb++ {
				_fafg := _eceda[_bbdc*_deeb : _bbdc*(_deeb+1)]
				_fca := _fafg[0]
				switch _fca {
				case 0:
				case 1:
					for _cddc := 2; _cddc < _bbdc; _cddc++ {
						_fafg[_cddc] = byte(int(_fafg[_cddc]+_fafg[_cddc-1]) % 256)
					}
				case 2:
					for _fagdc := 1; _fagdc < _bbdc; _fagdc++ {
						_fafg[_fagdc] = byte(int(_fafg[_fagdc]+_fdgb[_fagdc]) % 256)
					}
				default:
					_df.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0066i\u006c\u0074\u0065\u0072\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _fca)
					return nil, _ea.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _fca)
				}
				for _ecgb := 0; _ecgb < _bbdc; _ecgb++ {
					_fdgb[_ecgb] = _fafg[_ecgb]
				}
				_eda.Write(_fafg[1:])
			}
			_fddf := _eda.Bytes()
			return _fddf, nil
		} else {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _eggc.Predictor)
			return nil, _ea.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _eggc.Predictor)
		}
	}
	return _eceda, nil
}

// GetString returns the *PdfObjectString represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetString(obj PdfObject) (_acfbee *PdfObjectString, _accf bool) {
	_acfbee, _accf = TraceToDirectObject(obj).(*PdfObjectString)
	return _acfbee, _accf
}

// XrefTable represents the cross references in a PDF, i.e. the table of objects and information
// where to access within the PDF file.
type XrefTable struct {
	ObjectMap map[int]XrefObject
	_ggf      []XrefObject
}

func (_eebe *PdfParser) parseXrefTable() (*PdfObjectDictionary, error) {
	var _cade *PdfObjectDictionary
	_aeab, _aaaab := _eebe.readTextLine()
	if _aaaab != nil {
		return nil, _aaaab
	}
	if _eebe._dcad && _cb.Count(_cb.TrimPrefix(_aeab, "\u0078\u0072\u0065\u0066"), "\u0020") > 0 {
		_eebe._ffge._fdg = true
	}
	_df.Log.Trace("\u0078\u0072\u0065\u0066 f\u0069\u0072\u0073\u0074\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _aeab)
	_eebc := -1
	_ddde := 0
	_aefab := false
	_dabgc := ""
	for {
		_eebe.skipSpaces()
		_, _abdd := _eebe._gcec.Peek(1)
		if _abdd != nil {
			return nil, _abdd
		}
		_aeab, _abdd = _eebe.readTextLine()
		if _abdd != nil {
			return nil, _abdd
		}
		_efbb := _bcbac.FindStringSubmatch(_aeab)
		if len(_efbb) == 0 {
			_agab := len(_dabgc) > 0
			_dabgc += _aeab + "\u000a"
			if _agab {
				_efbb = _bcbac.FindStringSubmatch(_dabgc)
			}
		}
		if len(_efbb) == 3 {
			if _eebe._dcad && !_eebe._ffge._bfa {
				var (
					_ecfbc bool
					_ggffe int
				)
				for _, _eff := range _aeab {
					if _ec.IsDigit(_eff) {
						if _ecfbc {
							break
						}
						continue
					}
					if !_ecfbc {
						_ecfbc = true
					}
					_ggffe++
				}
				if _ggffe > 1 {
					_eebe._ffge._bfa = true
				}
			}
			_gebg, _ := _be.Atoi(_efbb[1])
			_eaeb, _ := _be.Atoi(_efbb[2])
			_eebc = _gebg
			_ddde = _eaeb
			_aefab = true
			_dabgc = ""
			_df.Log.Trace("\u0078r\u0065\u0066 \u0073\u0075\u0062s\u0065\u0063\u0074\u0069\u006f\u006e\u003a \u0066\u0069\u0072\u0073\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0064\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _eebc, _ddde)
			continue
		}
		_gbce := _aefgf.FindStringSubmatch(_aeab)
		if len(_gbce) == 4 {
			if !_aefab {
				_df.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0058r\u0065\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006fr\u006da\u0074\u0021\u000a")
				return nil, _a.New("\u0078\u0072\u0065\u0066 i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_ddbgf, _ := _be.ParseInt(_gbce[1], 10, 64)
			_eabda, _ := _be.Atoi(_gbce[2])
			_fgbd := _gbce[3]
			_dabgc = ""
			if _cb.ToLower(_fgbd) == "\u006e" && _ddbgf > 1 {
				_fgf, _ffcf := _eebe._ggaf.ObjectMap[_eebc]
				if !_ffcf || _eabda > _fgf.Generation {
					_acca := XrefObject{ObjectNumber: _eebc, XType: XrefTypeTableEntry, Offset: _ddbgf, Generation: _eabda}
					_eebe._ggaf.ObjectMap[_eebc] = _acca
				}
			}
			_eebc++
			continue
		}
		if (len(_aeab) > 6) && (_aeab[:7] == "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_df.Log.Trace("\u0046o\u0075n\u0064\u0020\u0074\u0072\u0061i\u006c\u0065r\u0020\u002d\u0020\u0025\u0073", _aeab)
			if len(_aeab) > 9 {
				_baed := _eebe.GetFileOffset()
				_eebe.SetFileOffset(_baed - int64(len(_aeab)) + 7)
			}
			_eebe.skipSpaces()
			_eebe.skipComments()
			_df.Log.Trace("R\u0065\u0061\u0064\u0069ng\u0020t\u0072\u0061\u0069\u006c\u0065r\u0020\u0064\u0069\u0063\u0074\u0021")
			_df.Log.Trace("\u0070\u0065\u0065\u006b\u003a\u0020\u0022\u0025\u0073\u0022", _aeab)
			_cade, _abdd = _eebe.ParseDict()
			_df.Log.Trace("\u0045O\u0046\u0020\u0072\u0065a\u0064\u0069\u006e\u0067\u0020t\u0072a\u0069l\u0065\u0072\u0020\u0064\u0069\u0063\u0074!")
			if _abdd != nil {
				_df.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _abdd)
				return nil, _abdd
			}
			break
		}
		if _aeab == "\u0025\u0025\u0045O\u0046" {
			_df.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006e\u0064 \u006f\u0066\u0020\u0066\u0069\u006c\u0065 -\u0020\u0074\u0072\u0061i\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066ou\u006e\u0064 \u002d\u0020\u0065\u0072\u0072\u006f\u0072\u0021")
			return nil, _a.New("\u0065\u006e\u0064 \u006f\u0066\u0020\u0066i\u006c\u0065\u0020\u002d\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_df.Log.Trace("\u0078\u0072\u0065\u0066\u0020\u006d\u006f\u0072\u0065 \u003a\u0020\u0025\u0073", _aeab)
	}
	_df.Log.Trace("\u0045\u004f\u0046 p\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0021")
	if _eebe._dfbg == nil {
		_fefd := XrefTypeTableEntry
		_eebe._dfbg = &_fefd
	}
	return _cade, nil
}

// LookupByReference looks up a PdfObject by a reference.
func (_bea *PdfParser) LookupByReference(ref PdfObjectReference) (PdfObject, error) {
	_df.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0075\u0070\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0025\u0073", ref.String())
	return _bea.LookupByNumber(int(ref.ObjectNumber))
}

// PdfObjectString represents the primitive PDF string object.
type PdfObjectString struct {
	_bcfef string
	_aae   bool
}

// GetFloatVal returns the float64 value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetFloatVal(obj PdfObject) (_abbc float64, _deabe bool) {
	_ecbe, _deabe := TraceToDirectObject(obj).(*PdfObjectFloat)
	if _deabe {
		return float64(*_ecbe), true
	}
	return 0, false
}

func (_afdc *PdfParser) seekToEOFMarker(_dfbd int64) error {
	var _gecf int64
	var _gcac int64 = 2048
	for _gecf < _dfbd-4 {
		if _dfbd <= (_gcac + _gecf) {
			_gcac = _dfbd - _gecf
		}
		_, _aeag := _afdc._abdga.Seek(_dfbd-_gecf-_gcac, _gd.SeekStart)
		if _aeag != nil {
			return _aeag
		}
		_cbec := make([]byte, _gcac)
		_afdc._abdga.Read(_cbec)
		_df.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_cbec))
		_egfg := _eggcd.FindAllStringIndex(string(_cbec), -1)
		if _egfg != nil {
			_acgd := _egfg[len(_egfg)-1]
			_df.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _egfg)
			_dffef := _dfbd - _gecf - _gcac + int64(_acgd[0])
			_afdc._abdga.Seek(_dffef, _gd.SeekStart)
			return nil
		}
		_df.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_gecf += _gcac - 4
	}
	_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _dcef
}

// ParseIndirectObject parses an indirect object from the input stream. Can also be an object stream.
// Returns the indirect object (*PdfIndirectObject) or the stream object (*PdfObjectStream).
func (_ggcf *PdfParser) ParseIndirectObject() (PdfObject, error) {
	_abed := PdfIndirectObject{}
	_abed._egcg = _ggcf
	_df.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_agfc, _febaf := _ggcf._gcec.Peek(20)
	if _febaf != nil {
		if _febaf != _gd.EOF {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return &_abed, _febaf
		}
	}
	_df.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_agfc))
	_fcc := _bfdf.FindStringSubmatchIndex(string(_agfc))
	if len(_fcc) < 6 {
		if _febaf == _gd.EOF {
			return nil, _febaf
		}
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_agfc))
		return &_abed, _a.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_ggcf._gcec.Discard(_fcc[0])
	_df.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _fcc)
	_fabe := _fcc[1] - _fcc[0]
	_fdac := make([]byte, _fabe)
	_, _febaf = _ggcf.ReadAtLeast(_fdac, _fabe)
	if _febaf != nil {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _febaf)
		return nil, _febaf
	}
	_df.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _fdac)
	_ffgaf := _bfdf.FindStringSubmatch(string(_fdac))
	if len(_ffgaf) < 3 {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_fdac))
		return &_abed, _a.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_bdfc, _ := _be.Atoi(_ffgaf[1])
	_ffege, _ := _be.Atoi(_ffgaf[2])
	_abed.ObjectNumber = int64(_bdfc)
	_abed.GenerationNumber = int64(_ffege)
	for {
		_dcbfb, _ceedb := _ggcf._gcec.Peek(2)
		if _ceedb != nil {
			return &_abed, _ceedb
		}
		_df.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_dcbfb), string(_dcbfb))
		if IsWhiteSpace(_dcbfb[0]) {
			_ggcf.skipSpaces()
		} else if _dcbfb[0] == '%' {
			_ggcf.skipComments()
		} else if (_dcbfb[0] == '<') && (_dcbfb[1] == '<') {
			_df.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_abed.PdfObject, _ceedb = _ggcf.ParseDict()
			_df.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _ceedb)
			if _ceedb != nil {
				return &_abed, _ceedb
			}
			_df.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_dcbfb[0] == '/') || (_dcbfb[0] == '(') || (_dcbfb[0] == '[') || (_dcbfb[0] == '<') {
			_abed.PdfObject, _ceedb = _ggcf.parseObject()
			if _ceedb != nil {
				return &_abed, _ceedb
			}
			_df.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else if _dcbfb[0] == ']' {
			_df.Log.Debug("\u0057\u0041\u0052\u004e\u0049N\u0047\u003a\u0020\u0027\u005d\u0027 \u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u006eo\u0074\u0020\u0062\u0065i\u006e\u0067\u0020\u0075\u0073\u0065d\u0020\u0061\u0073\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006e\u0064\u0069n\u0067\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			_ggcf._gcec.Discard(1)
		} else {
			if _dcbfb[0] == 'e' {
				_afedc, _cdefc := _ggcf.readTextLine()
				if _cdefc != nil {
					return nil, _cdefc
				}
				if len(_afedc) >= 6 && _afedc[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _dcbfb[0] == 's' {
				_dcbfb, _ = _ggcf._gcec.Peek(10)
				if string(_dcbfb[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_gfc := 6
					if len(_dcbfb) > 6 {
						if IsWhiteSpace(_dcbfb[_gfc]) && _dcbfb[_gfc] != '\r' && _dcbfb[_gfc] != '\n' {
							_df.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0050\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_ggcf._ffge._ecg = true
							_gfc++
						}
						if _dcbfb[_gfc] == '\r' {
							_gfc++
							if _dcbfb[_gfc] == '\n' {
								_gfc++
							}
						} else if _dcbfb[_gfc] == '\n' {
							_gfc++
						} else {
							_ggcf._ffge._ecg = true
						}
					}
					_ggcf._gcec.Discard(_gfc)
					_dcccf, _abfd := _abed.PdfObject.(*PdfObjectDictionary)
					if !_abfd {
						return nil, _a.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_df.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _dcccf)
					_effg, _ggac := _ggcf.traceStreamLength(_dcccf.Get("\u004c\u0065\u006e\u0067\u0074\u0068"))
					if _ggac != nil {
						_df.Log.Debug("\u0046\u0061\u0069l\u0020\u0074\u006f\u0020t\u0072\u0061\u0063\u0065\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076", _ggac)
						return nil, _ggac
					}
					_df.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _effg)
					_aggbg, _gbcca := _effg.(*PdfObjectInteger)
					if !_gbcca {
						return nil, _a.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_cfbbb := *_aggbg
					if _cfbbb < 0 {
						return nil, _a.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					_beafg := _ggcf.GetFileOffset()
					_gbead := _ggcf.xrefNextObjectOffset(_beafg)
					if _beafg+int64(_cfbbb) > _gbead && _gbead > _beafg {
						_df.Log.Debug("E\u0078\u0070\u0065\u0063te\u0064 \u0065\u006e\u0064\u0069\u006eg\u0020\u0061\u0074\u0020\u0025\u0064", _beafg+int64(_cfbbb))
						_df.Log.Debug("\u004e\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020\u0061t\u0020\u0025\u0064", _gbead)
						_eabf := _gbead - _beafg - 17
						if _eabf < 0 {
							return nil, _a.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020go\u0069\u006e\u0067 \u0070\u0061s\u0074\u0020\u0062\u006f\u0075\u006ed\u0061\u0072i\u0065\u0073")
						}
						_df.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u0020c\u006f\u0072\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020%\u0064\u002e\u002e\u002e", _eabf)
						_cfbbb = PdfObjectInteger(_eabf)
						_dcccf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(_eabf))
					}
					if int64(_cfbbb) > _ggcf._gccgc {
						_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _a.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_bgbdb := make([]byte, _cfbbb)
					_, _ggac = _ggcf.ReadAtLeast(_bgbdb, int(_cfbbb))
					if _ggac != nil {
						_df.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_bgbdb), _bgbdb)
						_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ggac)
						return nil, _ggac
					}
					_adac := PdfObjectStream{}
					_adac.Stream = _bgbdb
					_adac.PdfObjectDictionary = _abed.PdfObject.(*PdfObjectDictionary)
					_adac.ObjectNumber = _abed.ObjectNumber
					_adac.GenerationNumber = _abed.GenerationNumber
					_adac.PdfObjectReference._egcg = _ggcf
					_ggcf.skipSpaces()
					_ggcf._gcec.Discard(9)
					_ggcf.skipSpaces()
					return &_adac, nil
				}
			}
			_abed.PdfObject, _ceedb = _ggcf.parseObject()
			if _abed.PdfObject == nil {
				_df.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
				_abed.PdfObject = MakeNull()
			}
			return &_abed, _ceedb
		}
	}
	if _abed.PdfObject == nil {
		_df.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
		_abed.PdfObject = MakeNull()
	}
	_df.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_abed, nil
}

// WriteString outputs the object as it is to be written to file.
func (_cfdfd *PdfObjectString) WriteString() string {
	var _afbb _d.Buffer
	if _cfdfd._aae {
		_bdcg := _ac.EncodeToString(_cfdfd.Bytes())
		_afbb.WriteString("\u003c")
		_afbb.WriteString(_bdcg)
		_afbb.WriteString("\u003e")
		return _afbb.String()
	}
	_fadd := map[byte]string{'\n': "\u005c\u006e", '\r': "\u005c\u0072", '\t': "\u005c\u0074", '\b': "\u005c\u0062", '\f': "\u005c\u0066", '(': "\u005c\u0028", ')': "\u005c\u0029", '\\': "\u005c\u005c"}
	_afbb.WriteString("\u0028")
	for _cfag := 0; _cfag < len(_cfdfd._bcfef); _cfag++ {
		_dgea := _cfdfd._bcfef[_cfag]
		if _gaffg, _bdcfa := _fadd[_dgea]; _bdcfa {
			_afbb.WriteString(_gaffg)
		} else {
			_afbb.WriteByte(_dgea)
		}
	}
	_afbb.WriteString("\u0029")
	return _afbb.String()
}

// ASCIIHexEncoder implements ASCII hex encoder/decoder.
type ASCIIHexEncoder struct{}

// NewJPXEncoder returns a new instance of JPXEncoder.
func NewJPXEncoder() *JPXEncoder { return &JPXEncoder{} }

// DecodeBytes decodes byte array with ASCII85. 5 ASCII characters -> 4 raw binary bytes
func (_acfd *ASCII85Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _gccd []byte
	_df.Log.Trace("\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0020\u0044e\u0063\u006f\u0064\u0065")
	_dfgg := 0
	_bfcd := false
	for _dfgg < len(encoded) && !_bfcd {
		_beeg := [5]byte{0, 0, 0, 0, 0}
		_cdba := 0
		_cabfe := 0
		_ddcd := 4
		for _cabfe < 5+_cdba {
			if _dfgg+_cabfe == len(encoded) {
				break
			}
			_ddfd := encoded[_dfgg+_cabfe]
			if IsWhiteSpace(_ddfd) {
				_cdba++
				_cabfe++
				continue
			} else if _ddfd == '~' && _dfgg+_cabfe+1 < len(encoded) && encoded[_dfgg+_cabfe+1] == '>' {
				_ddcd = (_cabfe - _cdba) - 1
				if _ddcd < 0 {
					_ddcd = 0
				}
				_bfcd = true
				break
			} else if _ddfd >= '!' && _ddfd <= 'u' {
				_ddfd -= '!'
			} else if _ddfd == 'z' && _cabfe-_cdba == 0 {
				_ddcd = 4
				_cabfe++
				break
			} else {
				_df.Log.Error("\u0046\u0061i\u006c\u0065\u0064\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020co\u0064\u0065")
				return nil, _a.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u0064\u0065\u0020e\u006e\u0063\u006f\u0075\u006e\u0074\u0065\u0072\u0065\u0064")
			}
			_beeg[_cabfe-_cdba] = _ddfd
			_cabfe++
		}
		_dfgg += _cabfe
		for _cbg := _ddcd + 1; _cbg < 5; _cbg++ {
			_beeg[_cbg] = 84
		}
		_bcdb := uint32(_beeg[0])*85*85*85*85 + uint32(_beeg[1])*85*85*85 + uint32(_beeg[2])*85*85 + uint32(_beeg[3])*85 + uint32(_beeg[4])
		_adbec := []byte{byte((_bcdb >> 24) & 0xff), byte((_bcdb >> 16) & 0xff), byte((_bcdb >> 8) & 0xff), byte(_bcdb & 0xff)}
		_gccd = append(_gccd, _adbec[:_ddcd]...)
	}
	_df.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0065\u006e\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", encoded)
	_df.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0064\u0065\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", _gccd)
	return _gccd, nil
}

// UpdateParams updates the parameter values of the encoder.
func (_faga *MultiEncoder) UpdateParams(params *PdfObjectDictionary) {
	for _, _bbea := range _faga._gbgba {
		_bbea.UpdateParams(params)
	}
}

var _efdd = _f.MustCompile("\u0025P\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

// HasInvalidSeparationAfterXRef implements core.ParserMetadata interface.
func (_bgga ParserMetadata) HasInvalidSeparationAfterXRef() bool { return _bgga._fdg }

// Decrypt attempts to decrypt the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful, false otherwise.
// An error is returned when there is a problem with decrypting.
func (_affb *PdfParser) Decrypt(password []byte) (bool, error) {
	if _affb._acg == nil {
		return false, _a.New("\u0063\u0068\u0065\u0063k \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u0072s\u0074")
	}
	_bggdg, _agggf := _affb._acg.authenticate(password)
	if _agggf != nil {
		return false, _agggf
	}
	if !_bggdg {
		_bggdg, _agggf = _affb._acg.authenticate([]byte(""))
	}
	return _bggdg, _agggf
}

func _gfabf(_bdgd string) (PdfObjectReference, error) {
	_edcgb := PdfObjectReference{}
	_feca := _dbgee.FindStringSubmatch(_bdgd)
	if len(_feca) < 3 {
		_df.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _edcgb, _a.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_cgdc, _ := _be.Atoi(_feca[1])
	_abcgd, _ := _be.Atoi(_feca[2])
	_edcgb.ObjectNumber = int64(_cgdc)
	_edcgb.GenerationNumber = int64(_abcgd)
	return _edcgb, nil
}

// NewCCITTFaxEncoder makes a new CCITTFax encoder.
func NewCCITTFaxEncoder() *CCITTFaxEncoder { return &CCITTFaxEncoder{Columns: 1728, EndOfBlock: true} }

// DecodeBytes decodes the CCITTFax encoded image data.
func (_fgef *CCITTFaxEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_egdb, _eccd := _dc.NewDecoder(encoded, _dc.DecodeOptions{Columns: _fgef.Columns, Rows: _fgef.Rows, K: _fgef.K, EncodedByteAligned: _fgef.EncodedByteAlign, BlackIsOne: _fgef.BlackIs1, EndOfBlock: _fgef.EndOfBlock, EndOfLine: _fgef.EndOfLine, DamagedRowsBeforeError: _fgef.DamagedRowsBeforeError})
	if _eccd != nil {
		return nil, _eccd
	}
	_dcba, _eccd := _gd.ReadAll(_egdb)
	if _eccd != nil {
		return nil, _eccd
	}
	return _dcba, nil
}

// StreamEncoder represents the interface for all PDF stream encoders.
type StreamEncoder interface {
	GetFilterName() string
	MakeDecodeParams() PdfObject
	MakeStreamDict() *PdfObjectDictionary
	UpdateParams(_eab *PdfObjectDictionary)
	EncodeBytes(_dcg []byte) ([]byte, error)
	DecodeBytes(_bfec []byte) ([]byte, error)
	DecodeStream(_gce *PdfObjectStream) ([]byte, error)
}

const _dffc = 32 << (^uint(0) >> 63)

// GetFilterName returns the name of the encoding filter.
func (_fcg *RawEncoder) GetFilterName() string { return StreamEncodingFilterNameRaw }

// NewCompliancePdfParser creates a new PdfParser that will parse input reader with the focus on extracting more metadata, which
// might affect performance of the regular PdfParser this function.
func NewCompliancePdfParser(rs _gd.ReadSeeker) (_dbgb *PdfParser, _baec error) {
	_dbgb = &PdfParser{_abdga: rs, ObjCache: make(objectCache), _dbaad: map[int64]bool{}, _dcad: true, _dgef: make(map[*PdfParser]*PdfParser)}
	if _baec = _dbgb.parseDetailedHeader(); _baec != nil {
		return nil, _baec
	}
	if _dbgb._aagb, _baec = _dbgb.loadXrefs(); _baec != nil {
		_df.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _baec)
		return nil, _baec
	}
	_df.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _dbgb._aagb)
	if len(_dbgb._ggaf.ObjectMap) == 0 {
		return nil, _ea.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	return _dbgb, nil
}

// Len returns the number of elements in the array.
func (_dcge *PdfObjectArray) Len() int {
	if _dcge == nil {
		return 0
	}
	return len(_dcge._cdea)
}

// String returns the PDF version as a string. Implements interface fmt.Stringer.
func (_fgbg Version) String() string {
	return _ea.Sprintf("\u00250\u0064\u002e\u0025\u0030\u0064", _fgbg.Major, _fgbg.Minor)
}

// JPXEncoder implements JPX encoder/decoder (dummy, for now)
// FIXME: implement
type JPXEncoder struct{}

const _abedf = 10

// NewParserFromString is used for testing purposes.
func NewParserFromString(txt string) *PdfParser {
	_abge := _d.NewReader([]byte(txt))
	_agabg := &PdfParser{ObjCache: objectCache{}, _abdga: _abge, _gcec: _fd.NewReader(_abge), _gccgc: int64(len(txt)), _dbaad: map[int64]bool{}, _dgef: make(map[*PdfParser]*PdfParser)}
	_agabg._ggaf.ObjectMap = make(map[int]XrefObject)
	return _agabg
}

// GetInt returns the *PdfObjectBool object that is represented by a PdfObject either directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetInt(obj PdfObject) (_abegg *PdfObjectInteger, _deccf bool) {
	_abegg, _deccf = TraceToDirectObject(obj).(*PdfObjectInteger)
	return _abegg, _deccf
}

// GetXrefTable returns the PDFs xref table.
func (_dbdc *PdfParser) GetXrefTable() XrefTable { return _dbdc._ggaf }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_cfcg *JPXEncoder) MakeDecodeParams() PdfObject { return nil }

func (_bgdb *PdfParser) checkLinearizedInformation(_acbc *PdfObjectDictionary) (bool, error) {
	var _bdgb error
	_bgdb._adcg, _bdgb = GetNumberAsInt64(_acbc.Get("\u004c"))
	if _bdgb != nil {
		return false, _bdgb
	}
	_bdgb = _bgdb.seekToEOFMarker(_bgdb._adcg)
	switch _bdgb {
	case nil:
		return true, nil
	case _dcef:
		return false, nil
	default:
		return false, _bdgb
	}
}

// DecodeBytes decodes a byte slice from Run length encoding.
//
// 7.4.5 RunLengthDecode Filter
// The RunLengthDecode filter decodes data that has been encoded in a simple byte-oriented format based on run length.
// The encoded data shall be a sequence of runs, where each run shall consist of a length byte followed by 1 to 128
// bytes of data. If the length byte is in the range 0 to 127, the following length + 1 (1 to 128) bytes shall be
// copied literally during decompression. If length is in the range 129 to 255, the following single byte shall be
// copied 257 - length (2 to 128) times during decompression. A length value of 128 shall denote EOD.
func (_ebaa *RunLengthEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_dacb := _d.NewReader(encoded)
	var _gbbg []byte
	for {
		_fagb, _aadc := _dacb.ReadByte()
		if _aadc != nil {
			return nil, _aadc
		}
		if _fagb > 128 {
			_fefg, _gef := _dacb.ReadByte()
			if _gef != nil {
				return nil, _gef
			}
			for _eddb := 0; _eddb < 257-int(_fagb); _eddb++ {
				_gbbg = append(_gbbg, _fefg)
			}
		} else if _fagb < 128 {
			for _gdag := 0; _gdag < int(_fagb)+1; _gdag++ {
				_cdga, _geba := _dacb.ReadByte()
				if _geba != nil {
					return nil, _geba
				}
				_gbbg = append(_gbbg, _cdga)
			}
		} else {
			break
		}
	}
	return _gbbg, nil
}

// PdfObjectFloat represents the primitive PDF floating point numerical object.
type PdfObjectFloat float64

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_babg *ASCII85Encoder) MakeDecodeParams() PdfObject { return nil }

// String returns a string describing `array`.
func (_eeba *PdfObjectArray) String() string {
	_efaab := "\u005b"
	for _eccf, _ebaf := range _eeba.Elements() {
		_efaab += _ebaf.String()
		if _eccf < (_eeba.Len() - 1) {
			_efaab += "\u002c\u0020"
		}
	}
	_efaab += "\u005d"
	return _efaab
}

func (_eedb *PdfCrypt) makeKey(_ffeg string, _caba, _fded uint32, _fba []byte) ([]byte, error) {
	_gabc, _edda := _eedb._bde[_ffeg]
	if !_edda {
		return nil, _ea.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _ffeg)
	}
	return _gabc.MakeKey(_caba, _fded, _fba)
}

// Get returns the i-th element of the array or nil if out of bounds (by index).
func (_eef *PdfObjectArray) Get(i int) PdfObject {
	if _eef == nil || i >= len(_eef._cdea) || i < 0 {
		return nil
	}
	return _eef._cdea[i]
}

// GetRevisionNumber returns the current version of the Pdf document.
func (_fdece *PdfParser) GetRevisionNumber() int { return _fdece._eaae }

// AddPageImage adds the page with the image 'img' to the encoder context in order to encode it jbig2 document.
// The 'settings' defines what encoding type should be used by the encoder.
func (_ffac *JBIG2Encoder) AddPageImage(img *JBIG2Image, settings *JBIG2EncoderSettings) (_caa error) {
	const _agd = "\u004a\u0042\u0049\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065\u0049m\u0061\u0067\u0065"
	if _ffac == nil {
		return _dd.Error(_agd, "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u0069\u0073 n\u0069\u006c")
	}
	if settings == nil {
		settings = &_ffac.DefaultPageSettings
	}
	if _ffac._efaa == nil {
		_ffac._efaa = _ecdf.InitEncodeDocument(settings.FileMode)
	}
	if _caa = settings.Validate(); _caa != nil {
		return _dd.Wrap(_caa, _agd, "")
	}
	_ddbg, _caa := img.toBitmap()
	if _caa != nil {
		return _dd.Wrap(_caa, _agd, "")
	}
	switch settings.Compression {
	case JB2Generic:
		if _caa = _ffac._efaa.AddGenericPage(_ddbg, settings.DuplicatedLinesRemoval); _caa != nil {
			return _dd.Wrap(_caa, _agd, "")
		}
	case JB2SymbolCorrelation:
		return _dd.Error(_agd, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return _dd.Error(_agd, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return _dd.Error(_agd, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return nil
}

// GetBoolVal returns the bool value within a *PdObjectBool represented by an PdfObject interface directly or indirectly.
// If the PdfObject does not represent a bool value, a default value of false is returned (found = false also).
func GetBoolVal(obj PdfObject) (_cbde bool, _bfecf bool) {
	_babf, _bfecf := TraceToDirectObject(obj).(*PdfObjectBool)
	if _bfecf {
		return bool(*_babf), true
	}
	return false, false
}

// Set sets the PdfObject at index i of the array. An error is returned if the index is outside bounds.
func (_dfce *PdfObjectArray) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_dfce._cdea) {
		return _a.New("\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_dfce._cdea[i] = obj
	return nil
}

// GetObjectStreams returns the *PdfObjectStreams represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetObjectStreams(obj PdfObject) (_bbbc *PdfObjectStreams, _ceag bool) {
	_bbbc, _ceag = obj.(*PdfObjectStreams)
	return _bbbc, _ceag
}

// MakeArrayFromIntegers creates an PdfObjectArray from a slice of ints, where each array element is
// an PdfObjectInteger.
func MakeArrayFromIntegers(vals []int) *PdfObjectArray {
	_gacdc := MakeArray()
	for _, _gbga := range vals {
		_gacdc.Append(MakeInteger(int64(_gbga)))
	}
	return _gacdc
}

// DecodeBytes decodes a slice of JPX encoded bytes and returns the result.
func (_eegf *JPXEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _eegf.GetFilterName())
	return encoded, ErrNoJPXDecode
}

// EncodeStream encodes the stream data using the encoded specified by the stream's dictionary.
func EncodeStream(streamObj *PdfObjectStream) error {
	_df.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_cgcad, _agfd := NewEncoderFromStream(streamObj)
	if _agfd != nil {
		_df.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _agfd)
		return _agfd
	}
	if _ecba, _cdegd := _cgcad.(*LZWEncoder); _cdegd {
		_ecba.EarlyChange = 0
		streamObj.PdfObjectDictionary.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(0))
	}
	_df.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u000a", _cgcad)
	_gaffe, _agfd := _cgcad.EncodeBytes(streamObj.Stream)
	if _agfd != nil {
		_df.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _agfd)
		return _agfd
	}
	streamObj.Stream = _gaffe
	streamObj.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_gaffe))))
	return nil
}

// PdfCryptNewDecrypt makes the document crypt handler based on the encryption dictionary
// and trailer dictionary. Returns an error on failure to process.
func PdfCryptNewDecrypt(parser *PdfParser, ed, trailer *PdfObjectDictionary) (*PdfCrypt, error) {
	_dbe := &PdfCrypt{_age: false, _gee: make(map[PdfObject]bool), _fag: make(map[PdfObject]bool), _becc: make(map[int]struct{}), _edd: parser}
	_dgd, _acd := ed.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectName)
	if !_acd {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0043\u0072\u0079\u0070\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u0066\u0069\u0065\u006c\u0064\u0021")
		return _dbe, _a.New("r\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0063\u0072\u0079\u0070\u0074 \u0066\u0069\u0065\u006c\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
	}
	if *_dgd != "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064" {
		_df.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020(%\u0073\u0029", *_dgd)
		return _dbe, _a.New("\u0075n\u0073u\u0070\u0070\u006f\u0072\u0074e\u0064\u0020F\u0069\u006c\u0074\u0065\u0072")
	}
	_dbe._dgc.Filter = string(*_dgd)
	if _cac, _cbfd := ed.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r").(*PdfObjectString); _cbfd {
		_dbe._dgc.SubFilter = _cac.Str()
		_df.Log.Debug("\u0055s\u0069n\u0067\u0020\u0073\u0075\u0062f\u0069\u006ct\u0065\u0072\u0020\u0025\u0073", _cac)
	}
	if L, _eed := ed.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _eed {
		if (*L % 8) != 0 {
			_df.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0065\u006ec\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
			return _dbe, _a.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0065\u006e\u0063\u0072y\u0070t\u0069o\u006e\u0020\u006c\u0065\u006e\u0067\u0074h")
		}
		_dbe._dgc.Length = int(*L)
	} else {
		_dbe._dgc.Length = 40
	}
	_dbe._dgc.V = 0
	if _bdb, _gbea := ed.Get("\u0056").(*PdfObjectInteger); _gbea {
		V := int(*_bdb)
		_dbe._dgc.V = V
		if V >= 1 && V <= 2 {
			_dbe._bde = _gff(_dbe._dgc.Length)
		} else if V >= 4 && V <= 5 {
			if _efd := _dbe.loadCryptFilters(ed); _efd != nil {
				return _dbe, _efd
			}
		} else {
			_df.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0061lg\u006f\u0020\u0056 \u003d \u0025\u0064", V)
			return _dbe, _a.New("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d")
		}
	}
	if _cbag := _fbea(&_dbe._bga, ed); _cbag != nil {
		return _dbe, _cbag
	}
	_fagg := ""
	if _egg, _aaca := trailer.Get("\u0049\u0044").(*PdfObjectArray); _aaca && _egg.Len() >= 1 {
		_gab, _ceb := GetString(_egg.Get(0))
		if !_ceb {
			return _dbe, _a.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0074r\u0061\u0069l\u0065\u0072\u0020\u0049\u0044")
		}
		_fagg = _gab.Str()
	} else {
		_df.Log.Debug("\u0054\u0072ai\u006c\u0065\u0072 \u0049\u0044\u0020\u0061rra\u0079 m\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0021")
	}
	_dbe._geg = _fagg
	return _dbe, nil
}

func (_aabb *PdfParser) parseXrefStream(_gcegf *PdfObjectInteger) (*PdfObjectDictionary, error) {
	if _gcegf != nil {
		_df.Log.Trace("\u0058\u0052\u0065f\u0053\u0074\u006d\u0020x\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0061\u0074\u0020\u0025\u0064", _gcegf)
		_aabb._abdga.Seek(int64(*_gcegf), _gd.SeekStart)
		_aabb._gcec = _fd.NewReader(_aabb._abdga)
	}
	_agccg := _aabb.GetFileOffset()
	_egbbf, _gafg := _aabb.ParseIndirectObject()
	if _gafg != nil {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0078\u0072\u0065\u0066\u0020\u006fb\u006a\u0065\u0063\u0074")
		return nil, _a.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072e\u0061\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	_df.Log.Trace("\u0058R\u0065f\u0053\u0074\u006d\u0020\u006fb\u006a\u0065c\u0074\u003a\u0020\u0025\u0073", _egbbf)
	_fdae, _eddd := _egbbf.(*PdfObjectStream)
	if !_eddd {
		_df.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0058R\u0065\u0066\u0053\u0074\u006d\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067 \u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0021")
		return nil, _a.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0070\u006f\u0069\u006e\u0074i\u006e\u0067\u0020\u0074\u006f\u0020a\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_caab := _fdae.PdfObjectDictionary
	_fcbf, _eddd := _fdae.PdfObjectDictionary.Get("\u0053\u0069\u007a\u0065").(*PdfObjectInteger)
	if !_eddd {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0073\u0069\u007a\u0065\u0020f\u0072\u006f\u006d\u0020\u0078\u0072\u0065f\u0020\u0073\u0074\u006d")
		return nil, _a.New("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0053\u0069\u007ae\u0020\u0066\u0072\u006f\u006d\u0020\u0078\u0072\u0065\u0066 \u0073\u0074\u006d")
	}
	if int64(*_fcbf) > 8388607 {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u0053\u0069\u007a\u0065\u0020\u0065x\u0063\u0065\u0065\u0064\u0065\u0064\u0020l\u0069\u006d\u0069\u0074\u002c\u0020\u006f\u0076\u0065\u0072\u00208\u0033\u0038\u0038\u0036\u0030\u0037\u0020\u0028\u0025\u0064\u0029", *_fcbf)
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_caae := _fdae.PdfObjectDictionary.Get("\u0057")
	_eecgf, _eddd := _caae.(*PdfObjectArray)
	if !_eddd {
		return nil, _a.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0020\u0069\u006e\u0020x\u0072\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_fgee := _eecgf.Len()
	if _fgee != 3 {
		_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020\u0028\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033\u0020\u002d\u0020\u0025\u0064\u0029", _fgee)
		return nil, _a.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065f\u0020s\u0074\u006d\u0020\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033")
	}
	var _adfc []int64
	for _aged := 0; _aged < 3; _aged++ {
		_edcd, _ffdc := GetInt(_eecgf.Get(_aged))
		if !_ffdc {
			return nil, _a.New("i\u006e\u0076\u0061\u006cid\u0020w\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		_adfc = append(_adfc, int64(*_edcd))
	}
	_feeac, _gafg := DecodeStream(_fdae)
	if _gafg != nil {
		_df.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f \u0064e\u0063o\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _gafg)
		return nil, _gafg
	}
	_eaee := int(_adfc[0])
	_cddf := int(_adfc[0] + _adfc[1])
	_ccbg := int(_adfc[0] + _adfc[1] + _adfc[2])
	_agggd := int(_adfc[0] + _adfc[1] + _adfc[2])
	if _eaee < 0 || _cddf < 0 || _ccbg < 0 {
		_df.Log.Debug("\u0045\u0072\u0072\u006fr\u0020\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u003c \u0030 \u0028\u0025\u0064\u002c\u0025\u0064\u002c%\u0064\u0029", _eaee, _cddf, _ccbg)
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _agggd == 0 {
		_df.Log.Debug("\u004e\u006f\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0069\u006e\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028\u0064\u0065\u006c\u0074\u0061\u0062\u0020=\u003d\u0020\u0030\u0029")
		return _caab, nil
	}
	_ddad := len(_feeac) / _agggd
	_cgeb := 0
	_dcgc := _fdae.PdfObjectDictionary.Get("\u0049\u006e\u0064e\u0078")
	var _gdca []int
	if _dcgc != nil {
		_df.Log.Trace("\u0049n\u0064\u0065\u0078\u003a\u0020\u0025b", _dcgc)
		_feaf, _gagd := _dcgc.(*PdfObjectArray)
		if !_gagd {
			_df.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0029")
			return nil, _a.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0049\u006e\u0064e\u0078\u0020\u006f\u0062je\u0063\u0074")
		}
		if _feaf.Len()%2 != 0 {
			_df.Log.Debug("\u0057\u0041\u0052\u004eI\u004e\u0047\u0020\u0046\u0061\u0069\u006c\u0075\u0072e\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020i\u006e\u0064\u0065\u0078\u0020n\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006f\u0066\u0020\u0032\u002e")
			return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
		}
		_cgeb = 0
		_fcdc, _egde := _feaf.ToIntegerArray()
		if _egde != nil {
			_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0065\u0078 \u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0073\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0073\u003a\u0020\u0025\u0076", _egde)
			return nil, _egde
		}
		for _gaef := 0; _gaef < len(_fcdc); _gaef += 2 {
			_aagbf := _fcdc[_gaef]
			_efeb := _fcdc[_gaef+1]
			for _dgee := 0; _dgee < _efeb; _dgee++ {
				_gdca = append(_gdca, _aagbf+_dgee)
			}
			_cgeb += _efeb
		}
	} else {
		for _dcefd := 0; _dcefd < int(*_fcbf); _dcefd++ {
			_gdca = append(_gdca, _dcefd)
		}
		_cgeb = int(*_fcbf)
	}
	if _ddad == _cgeb+1 {
		_df.Log.Debug("\u0049n\u0063\u006f\u006d\u0070ati\u0062\u0069\u006c\u0069t\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u0076\u0065\u0072\u0061\u0067\u0065\u0020\u006f\u0066\u0020\u0031\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u002d\u0020\u0061\u0070\u0070en\u0064\u0069\u006eg\u0020\u006f\u006e\u0065\u0020-\u0020M\u0061\u0079\u0020\u006c\u0065\u0061\u0064\u0020\u0074o\u0020\u0070\u0072\u006f\u0062\u006c\u0065\u006d\u0073")
		_fdgg := _cgeb - 1
		for _, _ccdec := range _gdca {
			if _ccdec > _fdgg {
				_fdgg = _ccdec
			}
		}
		_gdca = append(_gdca, _fdgg+1)
		_cgeb++
	}
	if _ddad != len(_gdca) {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020x\u0072\u0065\u0066 \u0073\u0074\u006d:\u0020\u006eu\u006d\u0020\u0065\u006e\u0074\u0072i\u0065s \u0021\u003d\u0020\u006c\u0065\u006e\u0028\u0069\u006e\u0064\u0069\u0063\u0065\u0073\u0029\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", _ddad, len(_gdca))
		return nil, _a.New("\u0078\u0072ef\u0020\u0073\u0074m\u0020\u006e\u0075\u006d en\u0074ri\u0065\u0073\u0020\u0021\u003d\u0020\u006cen\u0028\u0069\u006e\u0064\u0069\u0063\u0065s\u0029")
	}
	_df.Log.Trace("\u004f\u0062j\u0065\u0063\u0074s\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0025\u0064", _cgeb)
	_df.Log.Trace("\u0049\u006e\u0064i\u0063\u0065\u0073\u003a\u0020\u0025\u0020\u0064", _gdca)
	_gcbf := func(_dbgec []byte) int64 {
		var _dccc int64
		for _dfcd := 0; _dfcd < len(_dbgec); _dfcd++ {
			_dccc += int64(_dbgec[_dfcd]) * (1 << uint(8*(len(_dbgec)-_dfcd-1)))
		}
		return _dccc
	}
	_df.Log.Trace("\u0044e\u0063\u006f\u0064\u0065d\u0020\u0073\u0074\u0072\u0065a\u006d \u006ce\u006e\u0067\u0074\u0068\u003a\u0020\u0025d", len(_feeac))
	_babc := 0
	for _gegg := 0; _gegg < len(_feeac); _gegg += _agggd {
		_ffcef := _ccdddf(len(_feeac), _gegg, _gegg+_eaee)
		if _ffcef != nil {
			_df.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _ffcef)
			return nil, _ffcef
		}
		_efcf := _feeac[_gegg : _gegg+_eaee]
		_ffcef = _ccdddf(len(_feeac), _gegg+_eaee, _gegg+_cddf)
		if _ffcef != nil {
			_df.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _ffcef)
			return nil, _ffcef
		}
		_egdc := _feeac[_gegg+_eaee : _gegg+_cddf]
		_ffcef = _ccdddf(len(_feeac), _gegg+_cddf, _gegg+_ccbg)
		if _ffcef != nil {
			_df.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _ffcef)
			return nil, _ffcef
		}
		_bbda := _feeac[_gegg+_cddf : _gegg+_ccbg]
		_bdbb := _gcbf(_efcf)
		_cgfg := _gcbf(_egdc)
		_agbe := _gcbf(_bbda)
		if _adfc[0] == 0 {
			_bdbb = 1
		}
		if _babc >= len(_gdca) {
			_df.Log.Debug("X\u0052\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u002d\u0020\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063e\u0073s\u0020\u0069\u006e\u0064e\u0078\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073\u0020\u002d\u0020\u0062\u0072\u0065\u0061\u006b\u0069\u006e\u0067")
			break
		}
		_feafc := _gdca[_babc]
		_babc++
		_df.Log.Trace("%\u0064\u002e\u0020\u0070\u0031\u003a\u0020\u0025\u0020\u0078", _feafc, _efcf)
		_df.Log.Trace("%\u0064\u002e\u0020\u0070\u0032\u003a\u0020\u0025\u0020\u0078", _feafc, _egdc)
		_df.Log.Trace("%\u0064\u002e\u0020\u0070\u0033\u003a\u0020\u0025\u0020\u0078", _feafc, _bbda)
		_df.Log.Trace("\u0025d\u002e \u0078\u0072\u0065\u0066\u003a \u0025\u0064 \u0025\u0064\u0020\u0025\u0064", _feafc, _bdbb, _cgfg, _agbe)
		if _bdbb == 0 {
			_df.Log.Trace("-\u0020\u0046\u0072\u0065\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020-\u0020\u0063\u0061\u006e\u0020\u0070\u0072o\u0062\u0061\u0062\u006c\u0079\u0020\u0069\u0067\u006e\u006fr\u0065")
		} else if _bdbb == 1 {
			_df.Log.Trace("\u002d\u0020I\u006e\u0020\u0075\u0073e\u0020\u002d \u0075\u006e\u0063\u006f\u006d\u0070\u0072\u0065s\u0073\u0065\u0064\u0020\u0076\u0069\u0061\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0062", _egdc)
			if _cgfg == _agccg {
				_df.Log.Debug("\u0055\u0070d\u0061\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0058\u0052\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064", _feafc, _fdae.ObjectNumber)
				_feafc = int(_fdae.ObjectNumber)
			}
			if _fceg, _eeggc := _aabb._ggaf.ObjectMap[_feafc]; !_eeggc || int(_agbe) > _fceg.Generation {
				_daae := XrefObject{ObjectNumber: _feafc, XType: XrefTypeTableEntry, Offset: _cgfg, Generation: int(_agbe)}
				_aabb._ggaf.ObjectMap[_feafc] = _daae
			}
		} else if _bdbb == 2 {
			_df.Log.Trace("\u002d\u0020\u0049\u006e \u0075\u0073\u0065\u0020\u002d\u0020\u0063\u006f\u006d\u0070r\u0065s\u0073\u0065\u0064\u0020\u006f\u0062\u006ae\u0063\u0074")
			if _, _daggg := _aabb._ggaf.ObjectMap[_feafc]; !_daggg {
				_afa := XrefObject{ObjectNumber: _feafc, XType: XrefTypeObjectStream, OsObjNumber: int(_cgfg), OsObjIndex: int(_agbe)}
				_aabb._ggaf.ObjectMap[_feafc] = _afa
				_df.Log.Trace("\u0065\u006e\u0074\u0072\u0079\u003a\u0020\u0025\u002b\u0076", _afa)
			}
		} else {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0056\u0041L\u0049\u0044\u0020\u0054\u0059\u0050\u0045\u0020\u0058\u0072\u0065\u0066\u0053\u0074\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u003f\u002d\u002d\u002d\u002d\u002d\u002d-")
			continue
		}
	}
	if _aabb._dfbg == nil {
		_ecagf := XrefTypeObjectStream
		_aabb._dfbg = &_ecagf
	}
	return _caab, nil
}
func (_bgaff *PdfObjectInteger) String() string { return _ea.Sprintf("\u0025\u0064", *_bgaff) }

// UpdateParams updates the parameter values of the encoder.
// Implements StreamEncoder interface.
func (_eccc *JBIG2Encoder) UpdateParams(params *PdfObjectDictionary) {
	_ccefc, _aagf := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _aagf == nil {
		_eccc.BitsPerComponent = int(_ccefc)
	}
	_aefg, _aagf := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _aagf == nil {
		_eccc.Width = int(_aefg)
	}
	_acec, _aagf := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _aagf == nil {
		_eccc.Height = int(_acec)
	}
	_eeac, _aagf := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _aagf == nil {
		_eccc.ColorComponents = int(_eeac)
	}
}

// PdfObjectDictionary represents the primitive PDF dictionary/map object.
type PdfObjectDictionary struct {
	_ccfa map[PdfObjectName]PdfObject
	_aggf []PdfObjectName
	_gfff *_c.Mutex
	_fdad *PdfParser
}

// GetTrailer returns the PDFs trailer dictionary. The trailer dictionary is typically the starting point for a PDF,
// referencing other key objects that are important in the document structure.
func (_gebe *PdfParser) GetTrailer() *PdfObjectDictionary { return _gebe._aagb }

// MakeNull creates an PdfObjectNull.
func MakeNull() *PdfObjectNull { _daef := PdfObjectNull{}; return &_daef }

// MultiEncoder supports serial encoding.
type MultiEncoder struct{ _gbgba []StreamEncoder }

// Bytes returns the PdfObjectString content as a []byte array.
func (_gbcbg *PdfObjectString) Bytes() []byte { return []byte(_gbcbg._bcfef) }

func (_dea *PdfParser) checkPostEOFData() error {
	const _dgfd = "\u0025\u0025\u0045O\u0046"
	_, _geb := _dea._abdga.Seek(-int64(len([]byte(_dgfd)))-1, _gd.SeekEnd)
	if _geb != nil {
		return _geb
	}
	_ccg := make([]byte, len([]byte(_dgfd))+1)
	_, _geb = _dea._abdga.Read(_ccg)
	if _geb != nil {
		if _geb != _gd.EOF {
			return _geb
		}
	}
	if string(_ccg) == _dgfd || string(_ccg) == _dgfd+"\u000a" {
		_dea._ffge._dacf = true
	}
	return nil
}

// TraceToDirectObject traces a PdfObject to a direct object.  For example direct objects contained
// in indirect objects (can be double referenced even).
func TraceToDirectObject(obj PdfObject) PdfObject {
	if _daba, _faag := obj.(*PdfObjectReference); _faag {
		obj = _daba.Resolve()
	}
	_fagae, _eeff := obj.(*PdfIndirectObject)
	_edddd := 0
	for _eeff {
		obj = _fagae.PdfObject
		_fagae, _eeff = GetIndirect(obj)
		_edddd++
		if _edddd > _abedf {
			_df.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0061\u0063\u0065\u0020\u0064\u0065p\u0074\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u0062\u0065\u0079\u006fn\u0064\u0020\u0025\u0064\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0067oi\u006e\u0067\u0020\u0064\u0065\u0065\u0070\u0065\u0072\u0021", _abedf)
			return nil
		}
	}
	return obj
}

// WriteString outputs the object as it is to be written to file.
func (_bcca *PdfObjectStreams) WriteString() string {
	var _gddb _cb.Builder
	_gddb.WriteString(_be.FormatInt(_bcca.ObjectNumber, 10))
	_gddb.WriteString("\u0020\u0030\u0020\u0052")
	return _gddb.String()
}

// GetBool returns the *PdfObjectBool object that is represented by a PdfObject directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetBool(obj PdfObject) (_fgeec *PdfObjectBool, _ffega bool) {
	_fgeec, _ffega = TraceToDirectObject(obj).(*PdfObjectBool)
	return _fgeec, _ffega
}

func (_fbaa *PdfParser) parseObject() (PdfObject, error) {
	_df.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_fbaa.skipSpaces()
	for {
		_fedd, _ddbb := _fbaa._gcec.Peek(2)
		if _ddbb != nil {
			if _ddbb != _gd.EOF || len(_fedd) == 0 {
				return nil, _ddbb
			}
			if len(_fedd) == 1 {
				_fedd = append(_fedd, ' ')
			}
		}
		_df.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_fedd))
		if _fedd[0] == '/' {
			_ffcb, _cadb := _fbaa.parseName()
			_df.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _ffcb)
			return &_ffcb, _cadb
		} else if _fedd[0] == '(' {
			_df.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_ebbf, _gdda := _fbaa.parseString()
			return _ebbf, _gdda
		} else if _fedd[0] == '[' {
			_df.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_bdcf, _aeedb := _fbaa.parseArray()
			return _bdcf, _aeedb
		} else if (_fedd[0] == '<') && (_fedd[1] == '<') {
			_df.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			_fcgee, _affa := _fbaa.ParseDict()
			return _fcgee, _affa
		} else if _fedd[0] == '<' {
			_df.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			_abgaf, _eccdd := _fbaa.parseHexString()
			return _abgaf, _eccdd
		} else if _fedd[0] == '%' {
			_fbaa.readComment()
			_fbaa.skipSpaces()
		} else {
			_df.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_fedd, _ = _fbaa._gcec.Peek(15)
			_cddbd := string(_fedd)
			_df.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _cddbd)
			if (len(_cddbd) > 3) && (_cddbd[:4] == "\u006e\u0075\u006c\u006c") {
				_bcda, _ggcgb := _fbaa.parseNull()
				return &_bcda, _ggcgb
			} else if (len(_cddbd) > 4) && (_cddbd[:5] == "\u0066\u0061\u006cs\u0065") {
				_fdcff, _gfgd := _fbaa.parseBool()
				return &_fdcff, _gfgd
			} else if (len(_cddbd) > 3) && (_cddbd[:4] == "\u0074\u0072\u0075\u0065") {
				_gadg, _dcgg := _fbaa.parseBool()
				return &_gadg, _dcgg
			}
			_bdga := _dbgee.FindStringSubmatch(_cddbd)
			if len(_bdga) > 1 {
				_fedd, _ = _fbaa._gcec.ReadBytes('R')
				_df.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_fedd[:]))
				_egdd, _baaa := _gfabf(string(_fedd))
				_egdd._egcg = _fbaa
				return &_egdd, _baaa
			}
			_cabdb := _cadf.FindStringSubmatch(_cddbd)
			if len(_cabdb) > 1 {
				_df.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				_gfbbe, _fdcc := _fbaa.parseNumber()
				return _gfbbe, _fdcc
			}
			_cabdb = _dfff.FindStringSubmatch(_cddbd)
			if len(_cabdb) > 1 {
				_df.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_df.Log.Trace("\u0025\u0020\u0073", _cabdb)
				_ebc, _fcgd := _fbaa.parseNumber()
				return _ebc, _fcgd
			}
			_df.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _cddbd)
			return nil, _a.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}

// GoImageToJBIG2 creates a binary image on the base of 'i' golang image.Image.
// If the image is not a black/white image then the function converts provided input into
// JBIG2Image with 1bpp. For non grayscale images the function performs the conversion to the grayscale temp image.
// Then it checks the value of the gray image value if it's within bounds of the black white threshold.
// This 'bwThreshold' value should be in range (0.0, 1.0). The threshold checks if the grayscale pixel (uint) value
// is greater or smaller than 'bwThreshold' * 255. Pixels inside the range will be white, and the others will be black.
// If the 'bwThreshold' is equal to -1.0 - JB2ImageAutoThreshold then it's value would be set on the base of
// it's histogram using Triangle method. For more information go to:
//
//	https://www.mathworks.com/matlabcentral/fileexchange/28047-gray-image-thresholding-using-the-triangle-method
func GoImageToJBIG2(i _ff.Image, bwThreshold float64) (*JBIG2Image, error) {
	const _gcfb = "\u0047\u006f\u0049\u006d\u0061\u0067\u0065\u0054\u006fJ\u0042\u0049\u0047\u0032"
	if i == nil {
		return nil, _dd.Error(_gcfb, "i\u006d\u0061\u0067\u0065 '\u0069'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	var (
		_aeec uint8
		_bgaf _cf.Image
		_gbag error
	)
	if bwThreshold == JB2ImageAutoThreshold {
		_bgaf, _gbag = _cf.MonochromeConverter.Convert(i)
	} else if bwThreshold > 1.0 || bwThreshold < 0.0 {
		return nil, _dd.Error(_gcfb, "p\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074h\u0072\u0065\u0073\u0068\u006f\u006c\u0064 i\u0073\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u0020\u0072\u0061\u006ege\u0020\u007b0\u002e\u0030\u002c\u0020\u0031\u002e\u0030\u007d")
	} else {
		_aeec = uint8(255 * bwThreshold)
		_bgaf, _gbag = _cf.MonochromeThresholdConverter(_aeec).Convert(i)
	}
	if _gbag != nil {
		return nil, _gbag
	}
	return _dbfe(_bgaf), nil
}

// EncodeBytes encodes data into ASCII85 encoded format.
func (_acfbe *ASCII85Encoder) EncodeBytes(data []byte) ([]byte, error) {
	var _fcf _d.Buffer
	for _facg := 0; _facg < len(data); _facg += 4 {
		_ggdg := data[_facg]
		_cffcc := 1
		_dgb := byte(0)
		if _facg+1 < len(data) {
			_dgb = data[_facg+1]
			_cffcc++
		}
		_adc := byte(0)
		if _facg+2 < len(data) {
			_adc = data[_facg+2]
			_cffcc++
		}
		_fgbe := byte(0)
		if _facg+3 < len(data) {
			_fgbe = data[_facg+3]
			_cffcc++
		}
		_cea := (uint32(_ggdg) << 24) | (uint32(_dgb) << 16) | (uint32(_adc) << 8) | uint32(_fgbe)
		if _cea == 0 {
			_fcf.WriteByte('z')
		} else {
			_agag := _acfbe.base256Tobase85(_cea)
			for _, _efgfd := range _agag[:_cffcc+1] {
				_fcf.WriteByte(_efgfd + '!')
			}
		}
	}
	_fcf.WriteString("\u007e\u003e")
	return _fcf.Bytes(), nil
}
func (_aafd *offsetReader) Read(p []byte) (_adag int, _edf error) { return _aafd._eacg.Read(p) }

// UpdateParams updates the parameter values of the encoder.
func (_ggecb *LZWEncoder) UpdateParams(params *PdfObjectDictionary) {
	_baa, _def := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _def == nil {
		_ggecb.Predictor = int(_baa)
	}
	_deff, _def := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _def == nil {
		_ggecb.BitsPerComponent = int(_deff)
	}
	_eced, _def := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _def == nil {
		_ggecb.Columns = int(_eced)
	}
	_abea, _def := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _def == nil {
		_ggecb.Colors = int(_abea)
	}
	_gaf, _def := GetNumberAsInt64(params.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065"))
	if _def == nil {
		_ggecb.EarlyChange = int(_gaf)
	}
}

// GetStringVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetStringVal(obj PdfObject) (_efgc string, _fdcd bool) {
	_bdde, _fdcd := TraceToDirectObject(obj).(*PdfObjectString)
	if _fdcd {
		return _bdde.Str(), true
	}
	return
}

// DecodeStream decodes RunLengthEncoded stream object and give back decoded bytes.
func (_gdee *RunLengthEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _gdee.DecodeBytes(streamObj.Stream)
}

// HasOddLengthHexStrings checks if the document has odd length hexadecimal strings.
func (_gfe ParserMetadata) HasOddLengthHexStrings() bool { return _gfe._faae }

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
	_efaa  *_ecdf.Document

	// Globals are the JBIG2 global segments.
	Globals _db.Globals

	// IsChocolateData defines if the data is encoded such that
	// binary data '1' means black and '0' white.
	// otherwise the data is called vanilla.
	// Naming convention taken from: 'https://en.wikipedia.org/wiki/Binary_image#Interpretation'
	IsChocolateData bool

	// DefaultPageSettings are the settings parameters used by the jbig2 encoder.
	DefaultPageSettings JBIG2EncoderSettings
}

// ASCII85Encoder implements ASCII85 encoder/decoder.
type ASCII85Encoder struct{}

var (
	ErrUnsupportedEncodingParameters = _a.New("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	ErrNoCCITTFaxDecode              = _a.New("\u0043\u0043I\u0054\u0054\u0046\u0061\u0078\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJBIG2Decode                 = _a.New("\u004a\u0042\u0049\u0047\u0032\u0044\u0065c\u006f\u0064\u0065 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJPXDecode                   = _a.New("\u004a\u0050\u0058\u0044\u0065c\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoPdfVersion                  = _a.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	ErrTypeError                     = _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	ErrRangeError                    = _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrNotSupported                  = _aa.New("\u0066\u0065\u0061t\u0075\u0072\u0065\u0020n\u006f\u0074\u0020\u0063\u0075\u0072\u0072e\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
	ErrNotANumber                    = _a.New("\u006e\u006f\u0074 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
)

// PdfObjectBool represents the primitive PDF boolean object.
type PdfObjectBool bool

func (_acgf *PdfParser) parseName() (PdfObjectName, error) {
	var _dedc _d.Buffer
	_gcgg := false
	for {
		_bbgb, _cabfb := _acgf._gcec.Peek(1)
		if _cabfb == _gd.EOF {
			break
		}
		if _cabfb != nil {
			return PdfObjectName(_dedc.String()), _cabfb
		}
		if !_gcgg {
			if _bbgb[0] == '/' {
				_gcgg = true
				_acgf._gcec.ReadByte()
			} else if _bbgb[0] == '%' {
				_acgf.readComment()
				_acgf.skipSpaces()
			} else {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _bbgb, _bbgb)
				return PdfObjectName(_dedc.String()), _ea.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _bbgb[0])
			}
		} else {
			if IsWhiteSpace(_bbgb[0]) {
				break
			} else if (_bbgb[0] == '/') || (_bbgb[0] == '[') || (_bbgb[0] == '(') || (_bbgb[0] == ']') || (_bbgb[0] == '<') || (_bbgb[0] == '>') {
				break
			} else if _bbgb[0] == '#' {
				_baad, _fcda := _acgf._gcec.Peek(3)
				if _fcda != nil {
					return PdfObjectName(_dedc.String()), _fcda
				}
				_fbbbf, _fcda := _ac.DecodeString(string(_baad[1:3]))
				if _fcda != nil {
					_df.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0066o\u006c\u006co\u0077\u0069\u006e\u0067 \u0027\u0023\u0027\u002c \u0063\u006f\u006e\u0074\u0069n\u0075\u0069\u006e\u0067\u0020\u0075\u0073i\u006e\u0067\u0020\u006c\u0069t\u0065\u0072\u0061\u006c\u0020\u002d\u0020\u004f\u0075t\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074")
					_dedc.WriteByte('#')
					_acgf._gcec.Discard(1)
					continue
				}
				_acgf._gcec.Discard(3)
				_dedc.Write(_fbbbf)
			} else {
				_bacc, _ := _acgf._gcec.ReadByte()
				_dedc.WriteByte(_bacc)
			}
		}
	}
	return PdfObjectName(_dedc.String()), nil
}

func (_geafg *PdfParser) repairLocateXref() (int64, error) {
	_gabe := int64(1000)
	_geafg._abdga.Seek(-_gabe, _gd.SeekCurrent)
	_gegf, _efdc := _geafg._abdga.Seek(0, _gd.SeekCurrent)
	if _efdc != nil {
		return 0, _efdc
	}
	_feac := make([]byte, _gabe)
	_geafg._abdga.Read(_feac)
	_aggc := _gedee.FindAllStringIndex(string(_feac), -1)
	if len(_aggc) < 1 {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0065\u0070a\u0069\u0072\u003a\u0020\u0078\u0072\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021")
		return 0, _a.New("\u0072\u0065\u0070\u0061ir\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_efbef := int64(_aggc[len(_aggc)-1][0])
	_ggefa := _gegf + _efbef
	return _ggefa, nil
}

// GetFilterName returns the name of the encoding filter.
func (_gfac *RunLengthEncoder) GetFilterName() string { return StreamEncodingFilterNameRunLength }

// GetFilterName returns the name of the encoding filter.
func (_cbab *FlateEncoder) GetFilterName() string { return StreamEncodingFilterNameFlate }

// HasEOLAfterHeader gets information if there is a EOL after the version header.
func (_eba ParserMetadata) HasEOLAfterHeader() bool { return _eba._eddgg }

var _gagg = _f.MustCompile("\u0073t\u0061r\u0074\u0078\u003f\u0072\u0065f\u005c\u0073*\u0028\u005c\u0064\u002b\u0029")

// MakeDictMap creates a PdfObjectDictionary initialized from a map of keys to values.
func MakeDictMap(objmap map[string]PdfObject) *PdfObjectDictionary {
	_febg := MakeDict()
	return _febg.Update(objmap)
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_cfge *CCITTFaxEncoder) MakeStreamDict() *PdfObjectDictionary {
	_cfbb := MakeDict()
	_cfbb.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_cfge.GetFilterName()))
	_cfbb.SetIfNotNil("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _cfge.MakeDecodeParams())
	return _cfbb
}

// IsWhiteSpace checks if byte represents a white space character.
func IsWhiteSpace(ch byte) bool {
	if (ch == 0x00) || (ch == 0x09) || (ch == 0x0A) || (ch == 0x0C) || (ch == 0x0D) || (ch == 0x20) {
		return true
	}
	return false
}

// String returns a string describing `stream`.
func (_eeef *PdfObjectStream) String() string {
	return _ea.Sprintf("O\u0062j\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0025\u0064: \u0025\u0073", _eeef.ObjectNumber, _eeef.PdfObjectDictionary)
}

// Len returns the number of elements in the streams.
func (_bdegf *PdfObjectStreams) Len() int {
	if _bdegf == nil {
		return 0
	}
	return len(_bdegf._dgbb)
}

func _gded(_ebbe PdfObject, _egfbd int) PdfObject {
	if _egfbd > _abedf {
		_df.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _abedf)
		return MakeNull()
	}
	switch _adbfe := _ebbe.(type) {
	case *PdfIndirectObject:
		_ebbe = _gded((*_adbfe).PdfObject, _egfbd+1)
	case *PdfObjectArray:
		for _bfebc, _gecg := range (*_adbfe)._cdea {
			(*_adbfe)._cdea[_bfebc] = _gded(_gecg, _egfbd+1)
		}
	case *PdfObjectDictionary:
		for _ffafg, _fccf := range (*_adbfe)._ccfa {
			(*_adbfe)._ccfa[_ffafg] = _gded(_fccf, _egfbd+1)
		}
		_g.Slice((*_adbfe)._aggf, func(_afegbb, _cfbfc int) bool { return (*_adbfe)._aggf[_afegbb] < (*_adbfe)._aggf[_cfbfc] })
	}
	return _ebbe
}

func (_gad *PdfCrypt) saveCryptFilters(_dab *PdfObjectDictionary) error {
	if _gad._dgc.V < 4 {
		return _a.New("\u0063\u0061\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020V\u003e\u003d\u0034")
	}
	_gbg := MakeDict()
	_dab.Set("\u0043\u0046", _gbg)
	for _cfb, _bggg := range _gad._bde {
		if _cfb == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			continue
		}
		_ccde := _ceda(_bggg, "")
		_gbg.Set(PdfObjectName(_cfb), _ccde)
	}
	_dab.Set("\u0053\u0074\u0072\u0046", MakeName(_gad._fbe))
	_dab.Set("\u0053\u0074\u006d\u0046", MakeName(_gad._ddfb))
	return nil
}

// MakeEncodedString creates a PdfObjectString with encoded content, which can be either
// UTF-16BE or PDFDocEncoding depending on whether `utf16BE` is true or false respectively.
func MakeEncodedString(s string, utf16BE bool) *PdfObjectString {
	if utf16BE {
		var _beeeb _d.Buffer
		_beeeb.Write([]byte{0xFE, 0xFF})
		_beeeb.WriteString(_gf.StringToUTF16(s))
		return &PdfObjectString{_bcfef: _beeeb.String(), _aae: true}
	}
	return &PdfObjectString{_bcfef: string(_gf.StringToPDFDocEncoding(s)), _aae: false}
}

// Validate validates the page settings for the JBIG2 encoder.
func (_fbgb JBIG2EncoderSettings) Validate() error {
	const _geeeg = "\u0076a\u006ci\u0064\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065\u0072"
	if _fbgb.Threshold < 0 || _fbgb.Threshold > 1.0 {
		return _dd.Errorf(_geeeg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064\u0020\u0076a\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0076\u0027 \u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061n\u0067\u0065\u0020\u005b\u0030\u002e0\u002c\u0020\u0031.\u0030\u005d", _fbgb.Threshold)
	}
	if _fbgb.ResolutionX < 0 {
		return _dd.Errorf(_geeeg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0078\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _fbgb.ResolutionX)
	}
	if _fbgb.ResolutionY < 0 {
		return _dd.Errorf(_geeeg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0079\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _fbgb.ResolutionY)
	}
	if _fbgb.DefaultPixelValue != 0 && _fbgb.DefaultPixelValue != 1 {
		return _dd.Errorf(_geeeg, "de\u0066\u0061u\u006c\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072 \u0074\u0068\u0065\u0020\u0062\u0069\u0074\u003a \u007b0\u002c\u0031}", _fbgb.DefaultPixelValue)
	}
	if _fbgb.Compression != JB2Generic {
		return _dd.Errorf(_geeeg, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0063\u006fm\u0070\u0072\u0065\u0073s\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	}
	return nil
}

// UpdateParams updates the parameter values of the encoder.
func (_bdcb *ASCII85Encoder) UpdateParams(params *PdfObjectDictionary) {}

// MakeDict creates and returns an empty PdfObjectDictionary.
func MakeDict() *PdfObjectDictionary {
	_abce := &PdfObjectDictionary{}
	_abce._ccfa = map[PdfObjectName]PdfObject{}
	_abce._aggf = []PdfObjectName{}
	_abce._gfff = &_c.Mutex{}
	return _abce
}

// Append appends PdfObject(s) to the array.
func (_cgce *PdfObjectArray) Append(objects ...PdfObject) {
	if _cgce == nil {
		_df.Log.Debug("\u0057\u0061\u0072\u006e\u0020\u002d\u0020\u0041\u0074\u0074\u0065\u006d\u0070t\u0020\u0074\u006f\u0020\u0061\u0070p\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020a\u0072\u0072\u0061\u0079")
		return
	}
	_cgce._cdea = append(_cgce._cdea, objects...)
}

// MakeName creates a PdfObjectName from a string.
func MakeName(s string) *PdfObjectName { _fbgc := PdfObjectName(s); return &_fbgc }

// MakeBool creates a PdfObjectBool from a bool value.
func MakeBool(val bool) *PdfObjectBool { _cfef := PdfObjectBool(val); return &_cfef }

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
func ParseNumber(buf *_fd.Reader) (PdfObject, error) {
	_dbbf := false
	_fdbfc := true
	var _egedc _d.Buffer
	for {
		if _df.Log.IsLogLevel(_df.LogLevelTrace) {
			_df.Log.Trace("\u0050\u0061\u0072\u0073in\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0022\u0025\u0073\u0022", _egedc.String())
		}
		_acfa, _beccg := buf.Peek(1)
		if _beccg == _gd.EOF {
			break
		}
		if _beccg != nil {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0025\u0073", _beccg)
			return nil, _beccg
		}
		if _fdbfc && (_acfa[0] == '-' || _acfa[0] == '+') {
			_cdbac, _ := buf.ReadByte()
			_egedc.WriteByte(_cdbac)
			_fdbfc = false
		} else if IsDecimalDigit(_acfa[0]) {
			_fggc, _ := buf.ReadByte()
			_egedc.WriteByte(_fggc)
		} else if _acfa[0] == '.' {
			_afgg, _ := buf.ReadByte()
			_egedc.WriteByte(_afgg)
			_dbbf = true
		} else if _acfa[0] == 'e' || _acfa[0] == 'E' {
			_fgcc, _ := buf.ReadByte()
			_egedc.WriteByte(_fgcc)
			_dbbf = true
			_fdbfc = true
		} else {
			break
		}
	}
	var _bead PdfObject
	if _dbbf {
		_dgeg, _aaggf := _be.ParseFloat(_egedc.String(), 64)
		if _aaggf != nil {
			_df.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025v\u0020\u0065\u0072\u0072\u003d\u0025v\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0030\u002e\u0020\u004fu\u0074\u0070u\u0074\u0020\u006d\u0061y\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074", _egedc.String(), _aaggf)
			_dgeg = 0.0
		}
		_aaeg := PdfObjectFloat(_dgeg)
		_bead = &_aaeg
	} else {
		_eebf, _dafb := _be.ParseInt(_egedc.String(), 10, 64)
		if _dafb != nil {
			_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0025\u0076\u0020\u0065\u0072\u0072\u003d%\u0076\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 \u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006ec\u006f\u0072r\u0065c\u0074", _egedc.String(), _dafb)
			_eebf = 0
		}
		_baee := PdfObjectInteger(_eebf)
		_bead = &_baee
	}
	return _bead, nil
}

// IsDecimalDigit checks if the character is a part of a decimal number string.
func IsDecimalDigit(c byte) bool { return '0' <= c && c <= '9' }

// IsPrintable checks if a character is printable.
// Regular characters that are outside the range EXCLAMATION MARK(21h)
// (!) to TILDE (7Eh) (~) should be written using the hexadecimal notation.
func IsPrintable(c byte) bool { return 0x21 <= c && c <= 0x7E }

func _ggd(_cba *_gb.StdEncryptDict, _dfdg *PdfObjectDictionary) {
	_dfdg.Set("\u0052", MakeInteger(int64(_cba.R)))
	_dfdg.Set("\u0050", MakeInteger(int64(_cba.P)))
	_dfdg.Set("\u004f", MakeStringFromBytes(_cba.O))
	_dfdg.Set("\u0055", MakeStringFromBytes(_cba.U))
	if _cba.R >= 5 {
		_dfdg.Set("\u004f\u0045", MakeStringFromBytes(_cba.OE))
		_dfdg.Set("\u0055\u0045", MakeStringFromBytes(_cba.UE))
		_dfdg.Set("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", MakeBool(_cba.EncryptMetadata))
		if _cba.R > 5 {
			_dfdg.Set("\u0050\u0065\u0072m\u0073", MakeStringFromBytes(_cba.Perms))
		}
	}
}

// PdfVersion returns version of the PDF file.
func (_cddcg *PdfParser) PdfVersion() Version { return _cddcg._bdbfe }

// IsHexadecimal checks if the PdfObjectString contains Hexadecimal data.
func (_dfdfg *PdfObjectString) IsHexadecimal() bool { return _dfdfg._aae }

// ResolveReference resolves reference if `o` is a *PdfObjectReference and returns the object referenced to.
// Otherwise returns back `o`.
func ResolveReference(obj PdfObject) PdfObject {
	if _cfeb, _fbcg := obj.(*PdfObjectReference); _fbcg {
		return _cfeb.Resolve()
	}
	return obj
}

// RegisterCustomStreamEncoder register a custom encoder handler for certain filter.
func RegisterCustomStreamEncoder(filterName string, customStreamEncoder StreamEncoder) {
	_faed.Store(filterName, customStreamEncoder)
}

// IsNullObject returns true if `obj` is a PdfObjectNull.
func IsNullObject(obj PdfObject) bool {
	_, _babe := TraceToDirectObject(obj).(*PdfObjectNull)
	return _babe
}

// WriteString outputs the object as it is to be written to file.
func (_abae *PdfObjectName) WriteString() string {
	var _cdaa _d.Buffer
	if len(*_abae) > 127 {
		_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u0061\u006d\u0065\u0020t\u006fo\u0020l\u006f\u006e\u0067\u0020\u0028\u0025\u0073)", *_abae)
	}
	_cdaa.WriteString("\u002f")
	for _ccedc := 0; _ccedc < len(*_abae); _ccedc++ {
		_ecdbg := (*_abae)[_ccedc]
		if !IsPrintable(_ecdbg) || _ecdbg == '#' || IsDelimiter(_ecdbg) {
			_cdaa.WriteString(_ea.Sprintf("\u0023\u0025\u002e2\u0078", _ecdbg))
		} else {
			_cdaa.WriteByte(_ecdbg)
		}
	}
	return _cdaa.String()
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

// DecodeBytes decodes a slice of ASCII encoded bytes and returns the result.
func (_bgbd *ASCIIHexEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_geeg := _d.NewReader(encoded)
	var _gdbe []byte
	for {
		_ebfe, _feceg := _geeg.ReadByte()
		if _feceg != nil {
			return nil, _feceg
		}
		if _ebfe == '>' {
			break
		}
		if IsWhiteSpace(_ebfe) {
			continue
		}
		if (_ebfe >= 'a' && _ebfe <= 'f') || (_ebfe >= 'A' && _ebfe <= 'F') || (_ebfe >= '0' && _ebfe <= '9') {
			_gdbe = append(_gdbe, _ebfe)
		} else {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0061\u0073\u0063\u0069\u0069 \u0068\u0065\u0078\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072 \u0028\u0025\u0063\u0029", _ebfe)
			return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0073\u0063\u0069\u0069\u0020\u0068e\u0078 \u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0028\u0025\u0063\u0029", _ebfe)
		}
	}
	if len(_gdbe)%2 == 1 {
		_gdbe = append(_gdbe, '0')
	}
	_df.Log.Trace("\u0049\u006e\u0062\u006f\u0075\u006e\u0064\u0020\u0025\u0073", _gdbe)
	_gbae := make([]byte, _ac.DecodedLen(len(_gdbe)))
	_, _cabfc := _ac.Decode(_gbae, _gdbe)
	if _cabfc != nil {
		return nil, _cabfc
	}
	return _gbae, nil
}

// GetXrefOffset returns the offset of the xref table.
func (_fbbb *PdfParser) GetXrefOffset() int64 { return _fbbb._cdca }

// String returns a string describing `null`.
func (_fccb *PdfObjectNull) String() string { return "\u006e\u0075\u006c\u006c" }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_cdeg *RawEncoder) MakeDecodeParams() PdfObject { return nil }

// MakeDecodeParams makes a new instance of an encoding dictionary based on the current encoder settings.
func (_cagc *JBIG2Encoder) MakeDecodeParams() PdfObject { return MakeDict() }

// EncodeBytes encodes slice of bytes into JBIG2 encoding format.
// The input 'data' must be an image. In order to Decode it a user is responsible to
// load the codec ('png', 'jpg').
// Returns jbig2 single page encoded document byte slice. The encoder uses DefaultPageSettings
// to encode given image.
func (_aada *JBIG2Encoder) EncodeBytes(data []byte) ([]byte, error) {
	const _dacbb = "\u004aB\u0049\u0047\u0032\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002eE\u006e\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065\u0073"
	if _aada.ColorComponents != 1 || _aada.BitsPerComponent != 1 {
		return nil, _dd.Errorf(_dacbb, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u002e\u0020\u004a\u0042\u0049G\u0032\u0020E\u006e\u0063o\u0064\u0065\u0072\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020bi\u006e\u0061\u0072\u0079\u0020\u0069\u006d\u0061\u0067e\u0073\u0020\u0064\u0061\u0074\u0061")
	}
	var (
		_feaa *_dg.Bitmap
		_dfgb error
	)
	_decf := (_aada.Width * _aada.Height) == len(data)
	if _decf {
		_feaa, _dfgb = _dg.NewWithUnpaddedData(_aada.Width, _aada.Height, data)
	} else {
		_feaa, _dfgb = _dg.NewWithData(_aada.Width, _aada.Height, data)
	}
	if _dfgb != nil {
		return nil, _dfgb
	}
	_bfebb := _aada.DefaultPageSettings
	if _dfgb = _bfebb.Validate(); _dfgb != nil {
		return nil, _dd.Wrap(_dfgb, _dacbb, "")
	}
	if _aada._efaa == nil {
		_aada._efaa = _ecdf.InitEncodeDocument(_bfebb.FileMode)
	}
	switch _bfebb.Compression {
	case JB2Generic:
		if _dfgb = _aada._efaa.AddGenericPage(_feaa, _bfebb.DuplicatedLinesRemoval); _dfgb != nil {
			return nil, _dd.Wrap(_dfgb, _dacbb, "")
		}
	case JB2SymbolCorrelation:
		return nil, _dd.Error(_dacbb, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return nil, _dd.Error(_dacbb, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return nil, _dd.Error(_dacbb, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return _aada.Encode()
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
type xrefType int

// WriteString outputs the object as it is to be written to file.
func (_fgded *PdfIndirectObject) WriteString() string {
	var _bbeb _cb.Builder
	_bbeb.WriteString(_be.FormatInt(_fgded.ObjectNumber, 10))
	_bbeb.WriteString("\u0020\u0030\u0020\u0052")
	return _bbeb.String()
}

var _dbgee = _f.MustCompile("\u005e\\\u0073\u002a\u005b\u002d]\u002a\u0028\u005c\u0064\u002b)\u005cs\u002b(\u005c\u0064\u002b\u0029\u005c\u0073\u002bR")

// Elements returns a slice of the PdfObject elements in the array.
// Preferred over accessing the array directly as type may be changed in future major versions (v3).
func (_dfad *PdfObjectStreams) Elements() []PdfObject {
	if _dfad == nil {
		return nil
	}
	return _dfad._dgbb
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_agcec *RunLengthEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_gedb := _d.NewReader(data)
	var _bbfd []byte
	var _fbbg []byte
	_cda, _baca := _gedb.ReadByte()
	if _baca == _gd.EOF {
		return []byte{}, nil
	} else if _baca != nil {
		return nil, _baca
	}
	_bee := 1
	for {
		_fed, _gfbe := _gedb.ReadByte()
		if _gfbe == _gd.EOF {
			break
		} else if _gfbe != nil {
			return nil, _gfbe
		}
		if _fed == _cda {
			if len(_fbbg) > 0 {
				_fbbg = _fbbg[:len(_fbbg)-1]
				if len(_fbbg) > 0 {
					_bbfd = append(_bbfd, byte(len(_fbbg)-1))
					_bbfd = append(_bbfd, _fbbg...)
				}
				_bee = 1
				_fbbg = []byte{}
			}
			_bee++
			if _bee >= 127 {
				_bbfd = append(_bbfd, byte(257-_bee), _cda)
				_bee = 0
			}
		} else {
			if _bee > 0 {
				if _bee == 1 {
					_fbbg = []byte{_cda}
				} else {
					_bbfd = append(_bbfd, byte(257-_bee), _cda)
				}
				_bee = 0
			}
			_fbbg = append(_fbbg, _fed)
			if len(_fbbg) >= 127 {
				_bbfd = append(_bbfd, byte(len(_fbbg)-1))
				_bbfd = append(_bbfd, _fbbg...)
				_fbbg = []byte{}
			}
		}
		_cda = _fed
	}
	if len(_fbbg) > 0 {
		_bbfd = append(_bbfd, byte(len(_fbbg)-1))
		_bbfd = append(_bbfd, _fbbg...)
	} else if _bee > 0 {
		_bbfd = append(_bbfd, byte(257-_bee), _cda)
	}
	_bbfd = append(_bbfd, 128)
	return _bbfd, nil
}

// FlattenObject returns the contents of `obj`. In other words, `obj` with indirect objects replaced
// by their values.
// The replacements are made recursively to a depth of traceMaxDepth.
// NOTE: Dicts are sorted to make objects with same contents have the same PDF object strings.
func FlattenObject(obj PdfObject) PdfObject { return _gded(obj, 0) }

// String returns a string representation of the *PdfObjectString.
func (_fgbf *PdfObjectString) String() string { return _fgbf._bcfef }

// MakeIndirectObject creates an PdfIndirectObject with a specified direct object PdfObject.
func MakeIndirectObject(obj PdfObject) *PdfIndirectObject {
	_bffb := &PdfIndirectObject{}
	_bffb.PdfObject = obj
	return _bffb
}

func (_fefbd *PdfParser) skipComments() error {
	if _, _gcgf := _fefbd.skipSpaces(); _gcgf != nil {
		return _gcgf
	}
	_bccg := true
	for {
		_bafd, _gdgc := _fefbd._gcec.Peek(1)
		if _gdgc != nil {
			_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gdgc.Error())
			return _gdgc
		}
		if _bccg && _bafd[0] != '%' {
			return nil
		}
		_bccg = false
		if (_bafd[0] != '\r') && (_bafd[0] != '\n') {
			_fefbd._gcec.ReadByte()
		} else {
			break
		}
	}
	return _fefbd.skipComments()
}

// GetFilterArray returns the names of the underlying encoding filters in an array that
// can be used as /Filter entry.
func (_bfff *MultiEncoder) GetFilterArray() *PdfObjectArray {
	_eddf := make([]PdfObject, len(_bfff._gbgba))
	for _ebe, _bgagg := range _bfff._gbgba {
		_eddf[_ebe] = MakeName(_bgagg.GetFilterName())
	}
	return MakeArray(_eddf...)
}

// IsOctalDigit checks if a character can be part of an octal digit string.
func IsOctalDigit(c byte) bool { return '0' <= c && c <= '7' }

func (_bfdd *PdfParser) traceStreamLength(_fgad PdfObject) (PdfObject, error) {
	_dcbag, _eadb := _fgad.(*PdfObjectReference)
	if _eadb {
		_aeae, _gefbb := _bfdd._dbaad[_dcbag.ObjectNumber]
		if _gefbb && _aeae {
			_df.Log.Debug("\u0053t\u0072\u0065a\u006d\u0020\u004c\u0065n\u0067\u0074\u0068 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 u\u006e\u0072\u0065s\u006f\u006cv\u0065\u0064\u0020\u0028\u0069\u006cl\u0065\u0067a\u006c\u0029")
			return nil, _a.New("\u0069\u006c\u006c\u0065ga\u006c\u0020\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006fo\u0070")
		}
		_bfdd._dbaad[_dcbag.ObjectNumber] = true
	}
	_aecb, _bccf := _bfdd.Resolve(_fgad)
	if _bccf != nil {
		return nil, _bccf
	}
	_df.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _aecb)
	if _eadb {
		_bfdd._dbaad[_dcbag.ObjectNumber] = false
	}
	return _aecb, nil
}

func (_ccgf *PdfParser) readTextLine() (string, error) {
	var _dddc _d.Buffer
	for {
		_dcga, _bdegg := _ccgf._gcec.Peek(1)
		if _bdegg != nil {
			_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _bdegg.Error())
			return _dddc.String(), _bdegg
		}
		if (_dcga[0] != '\r') && (_dcga[0] != '\n') {
			_agadf, _ := _ccgf._gcec.ReadByte()
			_dddc.WriteByte(_agadf)
		} else {
			break
		}
	}
	return _dddc.String(), nil
}
func _gff(_fdcf int) cryptFilters { return cryptFilters{_gfb: _gc.NewFilterV2(_fdcf)} }

// GetNumbersAsFloat converts a list of pdf objects representing floats or integers to a slice of
// float64 values.
func GetNumbersAsFloat(objects []PdfObject) (_fbed []float64, _cbdd error) {
	for _, _caea := range objects {
		_bfae, _aaff := GetNumberAsFloat(_caea)
		if _aaff != nil {
			return nil, _aaff
		}
		_fbed = append(_fbed, _bfae)
	}
	return _fbed, nil
}

func (_gda *FlateEncoder) postDecodePredict(_dabg []byte) ([]byte, error) {
	if _gda.Predictor > 1 {
		if _gda.Predictor == 2 {
			_df.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_df.Log.Trace("\u0043\u006f\u006c\u006f\u0072\u0073\u003a\u0020\u0025\u0064", _gda.Colors)
			_gdgd := _gda.Columns * _gda.Colors
			if _gdgd < 1 {
				return []byte{}, nil
			}
			_dga := len(_dabg) / _gdgd
			if len(_dabg)%_gdgd != 0 {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_dabg), _gdgd)
			}
			if _gdgd%_gda.Colors != 0 {
				return nil, _ea.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _gdgd, _gda.Colors)
			}
			if _gdgd > len(_dabg) {
				_df.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _gdgd, len(_dabg))
				return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_df.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_dabg), _dabg)
			_gfeg := _d.NewBuffer(nil)
			for _adbe := 0; _adbe < _dga; _adbe++ {
				_dfbb := _dabg[_gdgd*_adbe : _gdgd*(_adbe+1)]
				for _gdgb := _gda.Colors; _gdgb < _gdgd; _gdgb++ {
					_dfbb[_gdgb] += _dfbb[_gdgb-_gda.Colors]
				}
				_gfeg.Write(_dfbb)
			}
			_bbff := _gfeg.Bytes()
			_df.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_bbff), _bbff)
			return _bbff, nil
		} else if _gda.Predictor >= 10 && _gda.Predictor <= 15 {
			_df.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_caed := _gda.Columns*_gda.Colors + 1
			_cfc := len(_dabg) / _caed
			if len(_dabg)%_caed != 0 {
				return nil, _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_dabg), _caed)
			}
			if _caed > len(_dabg) {
				_df.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _caed, len(_dabg))
				return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_feedb := _d.NewBuffer(nil)
			_df.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _gda.Columns)
			_df.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_dabg), _caed, _cfc)
			_dacg := make([]byte, _caed)
			for _fdb := 0; _fdb < _caed; _fdb++ {
				_dacg[_fdb] = 0
			}
			_gec := _gda.Colors
			for _egda := 0; _egda < _cfc; _egda++ {
				_cbc := _dabg[_caed*_egda : _caed*(_egda+1)]
				_dcda := _cbc[0]
				switch _dcda {
				case _ega:
				case _bgec:
					for _gfa := 1 + _gec; _gfa < _caed; _gfa++ {
						_cbc[_gfa] += _cbc[_gfa-_gec]
					}
				case _dace:
					for _cfgb := 1; _cfgb < _caed; _cfgb++ {
						_cbc[_cfgb] += _dacg[_cfgb]
					}
				case _bfc:
					for _fdaf := 1; _fdaf < _gec+1; _fdaf++ {
						_cbc[_fdaf] += _dacg[_fdaf] / 2
					}
					for _fgacd := _gec + 1; _fgacd < _caed; _fgacd++ {
						_cbc[_fgacd] += byte((int(_cbc[_fgacd-_gec]) + int(_dacg[_fgacd])) / 2)
					}
				case _agce:
					for _afed := 1; _afed < _caed; _afed++ {
						var _dca, _gebd, _dfbbc byte
						_gebd = _dacg[_afed]
						if _afed >= _gec+1 {
							_dca = _cbc[_afed-_gec]
							_dfbbc = _dacg[_afed-_gec]
						}
						_cbc[_afed] += _cecd(_dca, _gebd, _dfbbc)
					}
				default:
					_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029\u0020\u0040\u0072o\u0077\u0020\u0025\u0064", _dcda, _egda)
					return nil, _ea.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _dcda)
				}
				copy(_dacg, _cbc)
				_feedb.Write(_cbc[1:])
			}
			_aec := _feedb.Bytes()
			return _aec, nil
		} else {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _gda.Predictor)
			return nil, _ea.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _gda.Predictor)
		}
	}
	return _dabg, nil
}

const (
	DefaultJPEGQuality = 75
)

func (_fdf *PdfCrypt) encryptBytes(_efc []byte, _fea string, _cceb []byte) ([]byte, error) {
	_df.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_dgca, _abga := _fdf._bde[_fea]
	if !_abga {
		return nil, _ea.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _fea)
	}
	return _dgca.EncryptBytes(_efc, _cceb)
}

// JBIG2CompressionType defines the enum compression type used by the JBIG2Encoder.
type JBIG2CompressionType int

func (_cc *PdfParser) lookupByNumberWrapper(_gcb int, _fb bool) (PdfObject, bool, error) {
	_ga, _bcg, _cg := _cc.lookupByNumber(_gcb, _fb)
	if _cg != nil {
		return nil, _bcg, _cg
	}
	if !_bcg && _cc._acg != nil && _cc._acg._age && !_cc._acg.isDecrypted(_ga) {
		_ba := _cc._acg.Decrypt(_ga, 0, 0)
		if _ba != nil {
			return nil, _bcg, _ba
		}
	}
	return _ga, _bcg, nil
}

func _ffe(_de PdfObject) (int64, int64, error) {
	if _deg, _fg := _de.(*PdfIndirectObject); _fg {
		return _deg.ObjectNumber, _deg.GenerationNumber, nil
	}
	if _ggb, _fac := _de.(*PdfObjectStream); _fac {
		return _ggb.ObjectNumber, _ggb.GenerationNumber, nil
	}
	return 0, 0, _a.New("\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u002f\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062je\u0063\u0074")
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_egfe *RawEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }
func (_cccb *PdfObjectFloat) String() string                   { return _ea.Sprintf("\u0025\u0066", *_cccb) }

// WriteString outputs the object as it is to be written to file.
func (_cfbda *PdfObjectStream) WriteString() string {
	var _dfcac _cb.Builder
	_dfcac.WriteString(_be.FormatInt(_cfbda.ObjectNumber, 10))
	_dfcac.WriteString("\u0020\u0030\u0020\u0052")
	return _dfcac.String()
}

func (_ccbgd *PdfParser) seekPdfVersionTopDown() (int, int, error) {
	_ccbgd._abdga.Seek(0, _gd.SeekStart)
	_ccbgd._gcec = _fd.NewReader(_ccbgd._abdga)
	_ccdea := 20
	_ebd := make([]byte, _ccdea)
	for {
		_dgec, _dbegc := _ccbgd._gcec.ReadByte()
		if _dbegc != nil {
			if _dbegc == _gd.EOF {
				break
			} else {
				return 0, 0, _dbegc
			}
		}
		if IsDecimalDigit(_dgec) && _ebd[_ccdea-1] == '.' && IsDecimalDigit(_ebd[_ccdea-2]) && _ebd[_ccdea-3] == '-' && _ebd[_ccdea-4] == 'F' && _ebd[_ccdea-5] == 'D' && _ebd[_ccdea-6] == 'P' {
			_gbcf := int(_ebd[_ccdea-2] - '0')
			_fcgc := int(_dgec - '0')
			return _gbcf, _fcgc, nil
		}
		_ebd = append(_ebd[1:_ccdea], _dgec)
	}
	return 0, 0, _a.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// HasDataAfterEOF checks if there is some data after EOF marker.
func (_ffa ParserMetadata) HasDataAfterEOF() bool { return _ffa._dacf }

func (_dgga *offsetReader) Seek(offset int64, whence int) (int64, error) {
	if whence == _gd.SeekStart {
		offset += _dgga._fefc
	}
	_gcab, _dcaa := _dgga._eacg.Seek(offset, whence)
	if _dcaa != nil {
		return _gcab, _dcaa
	}
	if whence == _gd.SeekCurrent {
		_gcab -= _dgga._fefc
	}
	if _gcab < 0 {
		return 0, _a.New("\u0063\u006f\u0072\u0065\u002eo\u0066\u0066\u0073\u0065\u0074\u0052\u0065\u0061\u0064\u0065\u0072\u002e\u0053e\u0065\u006b\u003a\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	return _gcab, nil
}

func (_eaaf *PdfCrypt) authenticate(_abg []byte) (bool, error) {
	_eaaf._age = false
	_ceed := _eaaf.securityHandler()
	_ebg, _ggff, _dbee := _ceed.Authenticate(&_eaaf._bga, _abg)
	if _dbee != nil {
		return false, _dbee
	} else if _ggff == 0 || len(_ebg) == 0 {
		return false, nil
	}
	_eaaf._age = true
	_eaaf._edb = _ebg
	return true, nil
}

// GetIntVal returns the int value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetIntVal(obj PdfObject) (_cacc int, _edgg bool) {
	_bfddf, _edgg := TraceToDirectObject(obj).(*PdfObjectInteger)
	if _edgg && _bfddf != nil {
		return int(*_bfddf), true
	}
	return 0, false
}

// GetArray returns the *PdfObjectArray represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetArray(obj PdfObject) (_dbcdd *PdfObjectArray, _fffe bool) {
	_dbcdd, _fffe = TraceToDirectObject(obj).(*PdfObjectArray)
	return _dbcdd, _fffe
}

func (_bcbae *PdfParser) parseNull() (PdfObjectNull, error) {
	_, _bgbdc := _bcbae._gcec.Discard(4)
	return PdfObjectNull{}, _bgbdc
}

func (_aeb *PdfCrypt) checkAccessRights(_dcd []byte) (bool, _gb.Permissions, error) {
	_aad := _aeb.securityHandler()
	_bbfa, _cdee, _gbaaa := _aad.Authenticate(&_aeb._bga, _dcd)
	if _gbaaa != nil {
		return false, 0, _gbaaa
	} else if _cdee == 0 || len(_bbfa) == 0 {
		return false, 0, nil
	}
	return true, _cdee, nil
}

// Elements returns a slice of the PdfObject elements in the array.
func (_efaf *PdfObjectArray) Elements() []PdfObject {
	if _efaf == nil {
		return nil
	}
	return _efaf._cdea
}

func (_bgde *PdfObjectDictionary) setWithLock(_aebga PdfObjectName, _cded PdfObject, _dfcbd bool) {
	if _dfcbd {
		_bgde._gfff.Lock()
		defer _bgde._gfff.Unlock()
	}
	_, _bdcbb := _bgde._ccfa[_aebga]
	if !_bdcbb {
		_bgde._aggf = append(_bgde._aggf, _aebga)
	}
	_bgde._ccfa[_aebga] = _cded
}

func (_dabf *PdfParser) parseDetailedHeader() (_ccc error) {
	_dabf._abdga.Seek(0, _gd.SeekStart)
	_dabf._gcec = _fd.NewReader(_dabf._abdga)
	_efed := 20
	_eaag := make([]byte, _efed)
	var (
		_dgcg bool
		_dcdb int
	)
	for {
		_ebfg, _cagf := _dabf._gcec.ReadByte()
		if _cagf != nil {
			if _cagf == _gd.EOF {
				break
			} else {
				return _cagf
			}
		}
		if IsDecimalDigit(_ebfg) && _eaag[_efed-1] == '.' && IsDecimalDigit(_eaag[_efed-2]) && _eaag[_efed-3] == '-' && _eaag[_efed-4] == 'F' && _eaag[_efed-5] == 'D' && _eaag[_efed-6] == 'P' && _eaag[_efed-7] == '%' {
			_dabf._bdbfe = Version{Major: int(_eaag[_efed-2] - '0'), Minor: int(_ebfg - '0')}
			_dabf._ffge._ecef = _dcdb - 7
			_dgcg = true
			break
		}
		_dcdb++
		_eaag = append(_eaag[1:_efed], _ebfg)
	}
	if !_dgcg {
		return _ea.Errorf("n\u006f \u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066ou\u006e\u0064")
	}
	_geaf, _ccc := _dabf._gcec.ReadByte()
	if _ccc == _gd.EOF {
		return _ea.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0050d\u0066\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _ccc != nil {
		return _ccc
	}
	_dabf._ffge._eddgg = _geaf == '\n'
	_geaf, _ccc = _dabf._gcec.ReadByte()
	if _ccc != nil {
		return _ea.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _ccc)
	}
	if _geaf != '%' {
		return nil
	}
	_bggaf := make([]byte, 4)
	_, _ccc = _dabf._gcec.Read(_bggaf)
	if _ccc != nil {
		return _ea.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _ccc)
	}
	_dabf._ffge._cdef = [4]byte{_bggaf[0], _bggaf[1], _bggaf[2], _bggaf[3]}
	return nil
}

// ToFloat64Array returns a slice of all elements in the array as a float64 slice.  An error is
// returned if the array contains non-numeric objects (each element can be either PdfObjectInteger
// or PdfObjectFloat).
func (_bgafb *PdfObjectArray) ToFloat64Array() ([]float64, error) {
	var _fbcd []float64
	for _, _gdfb := range _bgafb.Elements() {
		switch _aegg := _gdfb.(type) {
		case *PdfObjectInteger:
			_fbcd = append(_fbcd, float64(*_aegg))
		case *PdfObjectFloat:
			_fbcd = append(_fbcd, float64(*_aegg))
		default:
			return nil, ErrTypeError
		}
	}
	return _fbcd, nil
}

func (_egff *PdfParser) parseString() (*PdfObjectString, error) {
	_egff._gcec.ReadByte()
	var _debd _d.Buffer
	_ffgd := 1
	for {
		_ccgb, _caf := _egff._gcec.Peek(1)
		if _caf != nil {
			return MakeString(_debd.String()), _caf
		}
		if _ccgb[0] == '\\' {
			_egff._gcec.ReadByte()
			_gdec, _bgcb := _egff._gcec.ReadByte()
			if _bgcb != nil {
				return MakeString(_debd.String()), _bgcb
			}
			if IsOctalDigit(_gdec) {
				_acef, _fbddg := _egff._gcec.Peek(2)
				if _fbddg != nil {
					return MakeString(_debd.String()), _fbddg
				}
				var _fegf []byte
				_fegf = append(_fegf, _gdec)
				for _, _cdec := range _acef {
					if IsOctalDigit(_cdec) {
						_fegf = append(_fegf, _cdec)
					} else {
						break
					}
				}
				_egff._gcec.Discard(len(_fegf) - 1)
				_df.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _fegf)
				_dcacf, _fbddg := _be.ParseUint(string(_fegf), 8, 32)
				if _fbddg != nil {
					return MakeString(_debd.String()), _fbddg
				}
				_debd.WriteByte(byte(_dcacf))
				continue
			}
			switch _gdec {
			case 'n':
				_debd.WriteRune('\n')
			case 'r':
				_debd.WriteRune('\r')
			case 't':
				_debd.WriteRune('\t')
			case 'b':
				_debd.WriteRune('\b')
			case 'f':
				_debd.WriteRune('\f')
			case '(':
				_debd.WriteRune('(')
			case ')':
				_debd.WriteRune(')')
			case '\\':
				_debd.WriteRune('\\')
			}
			continue
		} else if _ccgb[0] == '(' {
			_ffgd++
		} else if _ccgb[0] == ')' {
			_ffgd--
			if _ffgd == 0 {
				_egff._gcec.ReadByte()
				break
			}
		}
		_agga, _ := _egff._gcec.ReadByte()
		_debd.WriteByte(_agga)
	}
	return MakeString(_debd.String()), nil
}

// ParseDict reads and parses a PDF dictionary object enclosed with '<<' and '>>'
func (_agfb *PdfParser) ParseDict() (*PdfObjectDictionary, error) {
	_df.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_cfdec := MakeDict()
	_cfdec._fdad = _agfb
	_edg, _ := _agfb._gcec.ReadByte()
	if _edg != '<' {
		return nil, _a.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_edg, _ = _agfb._gcec.ReadByte()
	if _edg != '<' {
		return nil, _a.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_agfb.skipSpaces()
		_agfb.skipComments()
		_ffcc, _ccddb := _agfb._gcec.Peek(2)
		if _ccddb != nil {
			return nil, _ccddb
		}
		_df.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_ffcc), string(_ffcc))
		if (_ffcc[0] == '>') && (_ffcc[1] == '>') {
			_df.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_agfb._gcec.ReadByte()
			_agfb._gcec.ReadByte()
			break
		}
		_df.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_bgbf, _ccddb := _agfb.parseName()
		_df.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _bgbf)
		if _ccddb != nil {
			_df.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _ccddb)
			return nil, _ccddb
		}
		if len(_bgbf) > 4 && _bgbf[len(_bgbf)-4:] == "\u006e\u0075\u006c\u006c" {
			_edeb := _bgbf[0 : len(_bgbf)-4]
			_df.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _bgbf)
			_df.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _edeb)
			_agfb.skipSpaces()
			_dbgea, _ := _agfb._gcec.Peek(1)
			if _dbgea[0] == '/' {
				_cfdec.Set(_edeb, MakeNull())
				continue
			}
		}
		_agfb.skipSpaces()
		_ccec, _ccddb := _agfb.parseObject()
		if _ccddb != nil {
			return nil, _ccddb
		}
		_cfdec.Set(_bgbf, _ccec)
		if _df.Log.IsLogLevel(_df.LogLevelTrace) {
			_df.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _bgbf, _ccec.String())
		}
	}
	_df.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _cfdec, nil
}

// MakeStreamDict make a new instance of an encoding dictionary for a stream object.
func (_acff *ASCII85Encoder) MakeStreamDict() *PdfObjectDictionary {
	_cdefe := MakeDict()
	_cdefe.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_acff.GetFilterName()))
	return _cdefe
}

type objectCache map[int]PdfObject

// GetFloat returns the *PdfObjectFloat represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetFloat(obj PdfObject) (_cfae *PdfObjectFloat, _dabc bool) {
	_cfae, _dabc = TraceToDirectObject(obj).(*PdfObjectFloat)
	return _cfae, _dabc
}

// ToIntegerArray returns a slice of all array elements as an int slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_cbdg *PdfObjectArray) ToIntegerArray() ([]int, error) {
	var _fabg []int
	for _, _fdcbb := range _cbdg.Elements() {
		if _gggc, _cdaaf := _fdcbb.(*PdfObjectInteger); _cdaaf {
			_fabg = append(_fabg, int(*_gggc))
		} else {
			return nil, ErrTypeError
		}
	}
	return _fabg, nil
}

// Version represents a version of a PDF standard.
type Version struct {
	Major int
	Minor int
}

func _afbe(_adff _gd.ReadSeeker, _efbgf int64) (*limitedReadSeeker, error) {
	_, _gaab := _adff.Seek(0, _gd.SeekStart)
	if _gaab != nil {
		return nil, _gaab
	}
	return &limitedReadSeeker{_cgfaf: _adff, _dgfg: _efbgf}, nil
}

// NewEncoderFromStream creates a StreamEncoder based on the stream's dictionary.
func NewEncoderFromStream(streamObj *PdfObjectStream) (StreamEncoder, error) {
	_cdcda := TraceToDirectObject(streamObj.PdfObjectDictionary.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if _cdcda == nil {
		return NewRawEncoder(), nil
	}
	if _, _cdegg := _cdcda.(*PdfObjectNull); _cdegg {
		return NewRawEncoder(), nil
	}
	_egdf, _dfcab := _cdcda.(*PdfObjectName)
	if !_dfcab {
		_gbagf, _agfa := _cdcda.(*PdfObjectArray)
		if !_agfa {
			return nil, _ea.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _gbagf.Len() == 0 {
			return NewRawEncoder(), nil
		}
		if _gbagf.Len() != 1 {
			_fabgc, _becd := _daee(streamObj)
			if _becd != nil {
				_df.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _becd)
				return nil, _becd
			}
			_df.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _fabgc)
			return _fabgc, nil
		}
		_cdcda = _gbagf.Get(0)
		_egdf, _agfa = _cdcda.(*PdfObjectName)
		if !_agfa {
			return nil, _ea.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	if _bfdb, _edfb := _faed.Load(_egdf.String()); _edfb {
		return _bfdb.(StreamEncoder), nil
	}
	switch *_egdf {
	case StreamEncodingFilterNameFlate:
		return _daca(streamObj, nil)
	case StreamEncodingFilterNameLZW:
		return _bdbf(streamObj, nil)
	case StreamEncodingFilterNameDCT:
		return _defa(streamObj, nil)
	case StreamEncodingFilterNameRunLength:
		return _efba(streamObj, nil)
	case StreamEncodingFilterNameASCIIHex:
		return NewASCIIHexEncoder(), nil
	case StreamEncodingFilterNameASCII85, "\u0041\u0038\u0035":
		return NewASCII85Encoder(), nil
	case StreamEncodingFilterNameCCITTFax:
		return _dbad(streamObj, nil)
	case StreamEncodingFilterNameJBIG2:
		return _gcda(streamObj, nil)
	case StreamEncodingFilterNameJPX:
		return NewJPXEncoder(), nil
	}
	_df.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006fd\u0021")
	return nil, _ea.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064i\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u0020\u0028\u0025\u0073\u0029", *_egdf)
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set.  Some other parameters are generated elsewhere.
func (_acdc *DCTEncoder) MakeStreamDict() *PdfObjectDictionary {
	_dada := MakeDict()
	_dada.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_acdc.GetFilterName()))
	return _dada
}

// WriteString outputs the object as it is to be written to file.
func (_ffgabe *PdfObjectReference) WriteString() string {
	var _bgcc _cb.Builder
	_bgcc.WriteString(_be.FormatInt(_ffgabe.ObjectNumber, 10))
	_bgcc.WriteString("\u0020")
	_bgcc.WriteString(_be.FormatInt(_ffgabe.GenerationNumber, 10))
	_bgcc.WriteString("\u0020\u0052")
	return _bgcc.String()
}

type limitedReadSeeker struct {
	_cgfaf _gd.ReadSeeker
	_dgfg  int64
}

// WriteString outputs the object as it is to be written to file.
func (_adeac *PdfObjectFloat) WriteString() string {
	return _be.FormatFloat(float64(*_adeac), 'f', -1, 64)
}

// RawEncoder implements Raw encoder/decoder (no encoding, pass through)
type RawEncoder struct{}

// NewASCII85Encoder makes a new ASCII85 encoder.
func NewASCII85Encoder() *ASCII85Encoder { _cdcf := &ASCII85Encoder{}; return _cdcf }

// GetIndirect returns the *PdfIndirectObject represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetIndirect(obj PdfObject) (_aacag *PdfIndirectObject, _ageg bool) {
	obj = ResolveReference(obj)
	_aacag, _ageg = obj.(*PdfIndirectObject)
	return _aacag, _ageg
}

var _bcbac = _f.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064\u002b)\u005c\u0073\u002a\u0024")

func _dagdd(_edegf PdfObject) (*float64, error) {
	switch _bgcbb := _edegf.(type) {
	case *PdfObjectFloat:
		_bgfe := float64(*_bgcbb)
		return &_bgfe, nil
	case *PdfObjectInteger:
		_aaaae := float64(*_bgcbb)
		return &_aaaae, nil
	case *PdfObjectNull:
		return nil, nil
	}
	return nil, ErrNotANumber
}

// GetCrypter returns the PdfCrypt instance which has information about the PDFs encryption.
func (_cdbb *PdfParser) GetCrypter() *PdfCrypt { return _cdbb._acg }

// CheckAccessRights checks access rights and permissions for a specified password. If either user/owner password is
// specified, full rights are granted, otherwise the access rights are specified by the Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_dbff *PdfParser) CheckAccessRights(password []byte) (bool, _gb.Permissions, error) {
	if _dbff._acg == nil {
		return true, _gb.PermOwner, nil
	}
	return _dbff._acg.checkAccessRights(password)
}

func (_fbcf *PdfParser) xrefNextObjectOffset(_adca int64) int64 {
	_dfbfd := int64(0)
	if len(_fbcf._ggaf.ObjectMap) == 0 {
		return 0
	}
	if len(_fbcf._ggaf._ggf) == 0 {
		_ccae := 0
		for _, _cbfeb := range _fbcf._ggaf.ObjectMap {
			if _cbfeb.Offset > 0 {
				_ccae++
			}
		}
		if _ccae == 0 {
			return 0
		}
		_fbcf._ggaf._ggf = make([]XrefObject, _ccae)
		_bbgd := 0
		for _, _bgbfg := range _fbcf._ggaf.ObjectMap {
			if _bgbfg.Offset > 0 {
				_fbcf._ggaf._ggf[_bbgd] = _bgbfg
				_bbgd++
			}
		}
		_g.Slice(_fbcf._ggaf._ggf, func(_aceba, _cabae int) bool {
			return _fbcf._ggaf._ggf[_aceba].Offset < _fbcf._ggaf._ggf[_cabae].Offset
		})
	}
	_gaad := _g.Search(len(_fbcf._ggaf._ggf), func(_gfefd int) bool { return _fbcf._ggaf._ggf[_gfefd].Offset >= _adca })
	if _gaad < len(_fbcf._ggaf._ggf) {
		_dfbfd = _fbcf._ggaf._ggf[_gaad].Offset
	}
	return _dfbfd
}

func _dbfe(_gaga _cf.Image) *JBIG2Image {
	_edbe := _gaga.Base()
	return &JBIG2Image{Data: _edbe.Data, Width: _edbe.Width, Height: _edbe.Height, HasPadding: true}
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_gacd *JBIG2Encoder) MakeStreamDict() *PdfObjectDictionary {
	_begde := MakeDict()
	_begde.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_gacd.GetFilterName()))
	return _begde
}

func _ed(_feb XrefTable) {
	_df.Log.Debug("\u003dX\u003d\u0058\u003d\u0058\u003d")
	_df.Log.Debug("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_ggee := 0
	for _, _aga := range _feb.ObjectMap {
		_df.Log.Debug("i\u002b\u0031\u003a\u0020\u0025\u0064 \u0028\u006f\u0062\u006a\u0020\u006eu\u006d\u003a\u0020\u0025\u0064\u0020\u0067e\u006e\u003a\u0020\u0025\u0064\u0029\u0020\u002d\u003e\u0020%\u0064", _ggee+1, _aga.ObjectNumber, _aga.Generation, _aga.Offset)
		_ggee++
	}
}

// WriteString outputs the object as it is to be written to file.
func (_acebe *PdfObjectDictionary) WriteString() string {
	var _adbc _cb.Builder
	_adbc.WriteString("\u003c\u003c")
	for _, _eabcf := range _acebe._aggf {
		_ggdff := _acebe._ccfa[_eabcf]
		_adbc.WriteString(_eabcf.WriteString())
		_adbc.WriteString("\u0020")
		_adbc.WriteString(_ggdff.WriteString())
	}
	_adbc.WriteString("\u003e\u003e")
	return _adbc.String()
}

func _dbad(_cagec *PdfObjectStream, _fbca *PdfObjectDictionary) (*CCITTFaxEncoder, error) {
	_cdaf := NewCCITTFaxEncoder()
	_bcbb := _cagec.PdfObjectDictionary
	if _bcbb == nil {
		return _cdaf, nil
	}
	if _fbca == nil {
		_bdbd := TraceToDirectObject(_bcbb.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _bdbd != nil {
			switch _efgfdg := _bdbd.(type) {
			case *PdfObjectDictionary:
				_fbca = _efgfdg
			case *PdfObjectArray:
				if _efgfdg.Len() == 1 {
					if _gac, _efcdd := GetDict(_efgfdg.Get(0)); _efcdd {
						_fbca = _gac
					}
				}
			default:
				_df.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _bdbd)
				return nil, _a.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
		if _fbca == nil {
			_df.Log.Error("\u0044\u0065c\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064 %\u0023\u0076", _bdbd)
			return nil, _a.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _ccddd, _bgace := GetNumberAsInt64(_fbca.Get("\u004b")); _bgace == nil {
		_cdaf.K = int(_ccddd)
	}
	if _dec, _eedbg := GetNumberAsInt64(_fbca.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _eedbg == nil {
		_cdaf.Columns = int(_dec)
	} else {
		_cdaf.Columns = 1728
	}
	if _fcag, _cfed := GetNumberAsInt64(_fbca.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _cfed == nil {
		_cdaf.BlackIs1 = _fcag > 0
	} else {
		if _bdae, _fdcb := GetBoolVal(_fbca.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _fdcb {
			_cdaf.BlackIs1 = _bdae
		} else {
			if _cfbbf, _bfd := GetArray(_fbca.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _bfd {
				_gcfd, _dfdd := _cfbbf.ToIntegerArray()
				if _dfdd == nil {
					_cdaf.BlackIs1 = _gcfd[0] == 1 && _gcfd[1] == 0
				}
			}
		}
	}
	if _fecf, _cacd := GetNumberAsInt64(_fbca.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _cacd == nil {
		_cdaf.EncodedByteAlign = _fecf > 0
	} else {
		if _faff, _cbfa := GetBoolVal(_fbca.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _cbfa {
			_cdaf.EncodedByteAlign = _faff
		}
	}
	if _dgbd, _adba := GetNumberAsInt64(_fbca.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _adba == nil {
		_cdaf.EndOfLine = _dgbd > 0
	} else {
		if _gdgde, _cgdb := GetBoolVal(_fbca.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _cgdb {
			_cdaf.EndOfLine = _gdgde
		}
	}
	if _edbb, _bfed := GetNumberAsInt64(_fbca.Get("\u0052\u006f\u0077\u0073")); _bfed == nil {
		_cdaf.Rows = int(_edbb)
	}
	_cdaf.EndOfBlock = true
	if _fdcab, _dadaa := GetNumberAsInt64(_fbca.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _dadaa == nil {
		_cdaf.EndOfBlock = _fdcab > 0
	} else {
		if _bafe, _cgff := GetBoolVal(_fbca.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _cgff {
			_cdaf.EndOfBlock = _bafe
		}
	}
	if _cfcb, _gbed := GetNumberAsInt64(_fbca.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _gbed != nil {
		_cdaf.DamagedRowsBeforeError = int(_cfcb)
	}
	_df.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _fbca.String())
	return _cdaf, nil
}

// Clear resets the dictionary to an empty state.
func (_cebec *PdfObjectDictionary) Clear() {
	_cebec._aggf = []PdfObjectName{}
	_cebec._ccfa = map[PdfObjectName]PdfObject{}
	_cebec._gfff = &_c.Mutex{}
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

func (_dfc *PdfCrypt) decryptBytes(_fdab []byte, _bgfg string, _daf []byte) ([]byte, error) {
	_df.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_afe, _ccef := _dfc._bde[_bgfg]
	if !_ccef {
		return nil, _ea.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _bgfg)
	}
	return _afe.DecryptBytes(_fdab, _daf)
}

// EncodeBytes encodes the image data using either Group3 or Group4 CCITT facsimile (fax) encoding.
// `data` is expected to be 1 color component, 1 bit per component. It is also valid to provide 8 BPC, 1 CC image like
// a standard go image Gray data.
func (_gbcb *CCITTFaxEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _ceeg _cf.Gray
	switch len(data) {
	case _gbcb.Rows * _gbcb.Columns:
		_gagc, _edcc := _cf.NewImage(_gbcb.Columns, _gbcb.Rows, 8, 1, data, nil, nil)
		if _edcc != nil {
			return nil, _edcc
		}
		_ceeg = _gagc.(_cf.Gray)
	case (_gbcb.Columns * _gbcb.Rows) + 7>>3:
		_aebf, _cfee := _cf.NewImage(_gbcb.Columns, _gbcb.Rows, 1, 1, data, nil, nil)
		if _cfee != nil {
			return nil, _cfee
		}
		_gbbfg := _aebf.(*_cf.Monochrome)
		if _cfee = _gbbfg.AddPadding(); _cfee != nil {
			return nil, _cfee
		}
		_ceeg = _gbbfg
	default:
		if len(data) < _cf.BytesPerLine(_gbcb.Columns, 1, 1)*_gbcb.Rows {
			return nil, _a.New("p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020i\u006e\u0070\u0075t\u0020d\u0061\u0074\u0061")
		}
		_dcdc, _ebag := _cf.NewImage(_gbcb.Columns, _gbcb.Rows, 1, 1, data, nil, nil)
		if _ebag != nil {
			return nil, _ebag
		}
		_fffg := _dcdc.(*_cf.Monochrome)
		_ceeg = _fffg
	}
	_gga := make([][]byte, _gbcb.Rows)
	for _ecafd := 0; _ecafd < _gbcb.Rows; _ecafd++ {
		_fdff := make([]byte, _gbcb.Columns)
		for _dddg := 0; _dddg < _gbcb.Columns; _dddg++ {
			_feg := _ceeg.GrayAt(_dddg, _ecafd)
			_fdff[_dddg] = _feg.Y >> 7
		}
		_gga[_ecafd] = _fdff
	}
	_geee := &_dc.Encoder{K: _gbcb.K, Columns: _gbcb.Columns, EndOfLine: _gbcb.EndOfLine, EndOfBlock: _gbcb.EndOfBlock, BlackIs1: _gbcb.BlackIs1, DamagedRowsBeforeError: _gbcb.DamagedRowsBeforeError, Rows: _gbcb.Rows, EncodedByteAlign: _gbcb.EncodedByteAlign}
	return _geee.Encode(_gga), nil
}

// Append appends PdfObject(s) to the streams.
func (_cdcg *PdfObjectStreams) Append(objects ...PdfObject) {
	if _cdcg == nil {
		_df.Log.Debug("\u0057\u0061\u0072\u006e\u0020-\u0020\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0074\u006f\u0020\u0061p\u0070\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073")
		return
	}
	_cdcg._dgbb = append(_cdcg._dgbb, objects...)
}

// String returns the state of the bool as "true" or "false".
func (_dacc *PdfObjectBool) String() string {
	if *_dacc {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}

// GetObjectNums returns a sorted list of object numbers of the PDF objects in the file.
func (_ceee *PdfParser) GetObjectNums() []int {
	var _cagee []int
	for _, _gfacf := range _ceee._ggaf.ObjectMap {
		_cagee = append(_cagee, _gfacf.ObjectNumber)
	}
	_g.Ints(_cagee)
	return _cagee
}

// PdfObjectStreams represents the primitive PDF object streams.
// 7.5.7 Object Streams (page 45).
type PdfObjectStreams struct {
	PdfObjectReference
	_dgbb []PdfObject
}

// MakeObjectStreams creates an PdfObjectStreams from a list of PdfObjects.
func MakeObjectStreams(objects ...PdfObject) *PdfObjectStreams {
	return &PdfObjectStreams{_dgbb: objects}
}

func (_fbgf *PdfParser) parsePdfVersion() (int, int, error) {
	var _bgage int64 = 20
	_gccc := make([]byte, _bgage)
	_fbgf._abdga.Seek(0, _gd.SeekStart)
	_fbgf._abdga.Read(_gccc)
	var _ddfg error
	var _eabc, _bdcc int
	if _fegb := _efdd.FindStringSubmatch(string(_gccc)); len(_fegb) < 3 {
		if _eabc, _bdcc, _ddfg = _fbgf.seekPdfVersionTopDown(); _ddfg != nil {
			_df.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _ddfg
		}
		_fbgf._abdga, _ddfg = _cdacc(_fbgf._abdga, _fbgf.GetFileOffset()-8)
		if _ddfg != nil {
			return 0, 0, _ddfg
		}
	} else {
		if _eabc, _ddfg = _be.Atoi(_fegb[1]); _ddfg != nil {
			return 0, 0, _ddfg
		}
		if _bdcc, _ddfg = _be.Atoi(_fegb[2]); _ddfg != nil {
			return 0, 0, _ddfg
		}
		_fbgf.SetFileOffset(0)
	}
	_fbgf._gcec = _fd.NewReader(_fbgf._abdga)
	_df.Log.Debug("\u0050\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _eabc, _bdcc)
	return _eabc, _bdcc, nil
}

func _efba(_edaf *PdfObjectStream, _bbgg *PdfObjectDictionary) (*RunLengthEncoder, error) {
	return NewRunLengthEncoder(), nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_bbgc *JPXEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }

// MakeStringFromBytes creates an PdfObjectString from a byte array.
// This is more natural than MakeString as `data` is usually not utf-8 encoded.
func MakeStringFromBytes(data []byte) *PdfObjectString { return MakeString(string(data)) }

func _fbea(_ade *_gb.StdEncryptDict, _eac *PdfObjectDictionary) error {
	R, _bbf := _eac.Get("\u0052").(*PdfObjectInteger)
	if !_bbf {
		return _a.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0052")
	}
	if *R < 2 || *R > 6 {
		return _ea.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052 \u0028\u0025\u0064\u0029", *R)
	}
	_ade.R = int(*R)
	O, _bbf := _eac.GetString("\u004f")
	if !_bbf {
		return _a.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u004f")
	}
	if _ade.R == 5 || _ade.R == 6 {
		if len(O) < 48 {
			return _ea.Errorf("\u004c\u0065\u006e\u0067th\u0028\u004f\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(O))
		}
	} else if len(O) != 32 {
		return _ea.Errorf("L\u0065n\u0067\u0074\u0068\u0028\u004f\u0029\u0020\u0021=\u0020\u0033\u0032\u0020(%\u0064\u0029", len(O))
	}
	_ade.O = []byte(O)
	U, _bbf := _eac.GetString("\u0055")
	if !_bbf {
		return _a.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0055")
	}
	if _ade.R == 5 || _ade.R == 6 {
		if len(U) < 48 {
			return _ea.Errorf("\u004c\u0065\u006e\u0067th\u0028\u0055\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(U))
		}
	} else if len(U) != 32 {
		_df.Log.Debug("\u0057\u0061r\u006e\u0069\u006e\u0067\u003a\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0028\u0055\u0029\u0020\u0021\u003d\u0020\u0033\u0032\u0020(%\u0064\u0029", len(U))
	}
	_ade.U = []byte(U)
	if _ade.R >= 5 {
		OE, _bgc := _eac.GetString("\u004f\u0045")
		if !_bgc {
			return _a.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u004f\u0045")
		} else if len(OE) != 32 {
			return _ea.Errorf("L\u0065\u006e\u0067\u0074h(\u004fE\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(OE))
		}
		_ade.OE = []byte(OE)
		UE, _bgc := _eac.GetString("\u0055\u0045")
		if !_bgc {
			return _a.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u0055\u0045")
		} else if len(UE) != 32 {
			return _ea.Errorf("L\u0065\u006e\u0067\u0074h(\u0055E\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(UE))
		}
		_ade.UE = []byte(UE)
	}
	P, _bbf := _eac.Get("\u0050").(*PdfObjectInteger)
	if !_bbf {
		return _a.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074 \u0064\u0069\u0063t\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0061\u0074\u0074\u0072")
	}
	_ade.P = _gb.Permissions(*P)
	if _ade.R == 6 {
		Perms, _befd := _eac.GetString("\u0050\u0065\u0072m\u0073")
		if !_befd {
			return _a.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0050\u0065\u0072\u006d\u0073")
		} else if len(Perms) != 16 {
			return _ea.Errorf("\u004ce\u006e\u0067\u0074\u0068\u0028\u0050\u0065\u0072\u006d\u0073\u0029 \u0021\u003d\u0020\u0031\u0036\u0020\u0028\u0025\u0064\u0029", len(Perms))
		}
		_ade.Perms = []byte(Perms)
	}
	if _efa, _cab := _eac.Get("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061").(*PdfObjectBool); _cab {
		_ade.EncryptMetadata = bool(*_efa)
	} else {
		_ade.EncryptMetadata = true
	}
	return nil
}

type encryptDict struct {
	Filter    string
	V         int
	SubFilter string
	Length    int
	StmF      string
	StrF      string
	EFF       string
	CF        map[string]_gc.FilterDict
}

// GetFilterName returns the name of the encoding filter.
func (_eegg *JBIG2Encoder) GetFilterName() string { return StreamEncodingFilterNameJBIG2 }

// GetString is a helper for Get that returns a string value.
// Returns false if the key is missing or a value is not a string.
func (_ffebg *PdfObjectDictionary) GetString(key PdfObjectName) (string, bool) {
	_acfc := _ffebg.Get(key)
	if _acfc == nil {
		return "", false
	}
	_adead, _fdeca := _acfc.(*PdfObjectString)
	if !_fdeca {
		return "", false
	}
	return _adead.Str(), true
}

func (_eecde *PdfParser) rebuildXrefTable() error {
	_bfedg := XrefTable{}
	_bfedg.ObjectMap = map[int]XrefObject{}
	_ggda := make([]int, 0, len(_eecde._ggaf.ObjectMap))
	for _fabf := range _eecde._ggaf.ObjectMap {
		_ggda = append(_ggda, _fabf)
	}
	_g.Ints(_ggda)
	for _, _ddca := range _ggda {
		_ecdge := _eecde._ggaf.ObjectMap[_ddca]
		_adda, _, _cgda := _eecde.lookupByNumberWrapper(_ddca, false)
		if _cgda != nil {
			_df.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f l\u006f\u006f\u006b\u0020\u0075\u0070\u0020ob\u006a\u0065\u0063\u0074\u0020\u0028\u0025s\u0029", _cgda)
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u006fm\u0070\u006c\u0065\u0074\u0065\u006c\u0079\u0020\u0062\u0072\u006f\u006b\u0065\u006e\u0020\u002d\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0074\u006f \u0072\u0065\u0070\u0061\u0069r\u0020")
			_gdba, _eddge := _eecde.repairRebuildXrefsTopDown()
			if _eddge != nil {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0072\u0065\u0062\u0075\u0069l\u0064\u0020\u0072\u0065\u0070a\u0069\u0072 \u0028\u0025\u0073\u0029", _eddge)
				return _eddge
			}
			_eecde._ggaf = *_gdba
			_df.Log.Debug("\u0052e\u0070\u0061\u0069\u0072e\u0064\u0020\u0078\u0072\u0065f\u0020t\u0061b\u006c\u0065\u0020\u0062\u0075\u0069\u006ct")
			return nil
		}
		_cafc, _gefc, _cgda := _ffe(_adda)
		if _cgda != nil {
			return _cgda
		}
		_ecdge.ObjectNumber = int(_cafc)
		_ecdge.Generation = int(_gefc)
		_bfedg.ObjectMap[int(_cafc)] = _ecdge
	}
	_eecde._ggaf = _bfedg
	_df.Log.Debug("N\u0065w\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061b\u006c\u0065\u0020\u0062ui\u006c\u0074")
	_ed(_eecde._ggaf)
	return nil
}

// GetFilterName returns the name of the encoding filter.
func (_ded *DCTEncoder) GetFilterName() string { return StreamEncodingFilterNameDCT }

func _cfcd(_fafc *PdfObjectDictionary) (_edfa *_cf.ImageBase) {
	var (
		_becca *PdfObjectInteger
		_bfbb  bool
	)
	if _becca, _bfbb = _fafc.Get("\u0057\u0069\u0064t\u0068").(*PdfObjectInteger); _bfbb {
		_edfa = &_cf.ImageBase{Width: int(*_becca)}
	} else {
		return nil
	}
	if _becca, _bfbb = _fafc.Get("\u0048\u0065\u0069\u0067\u0068\u0074").(*PdfObjectInteger); _bfbb {
		_edfa.Height = int(*_becca)
	}
	if _becca, _bfbb = _fafc.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074").(*PdfObjectInteger); _bfbb {
		_edfa.BitsPerComponent = int(*_becca)
	}
	if _becca, _bfbb = _fafc.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073").(*PdfObjectInteger); _bfbb {
		_edfa.ColorComponents = int(*_becca)
	}
	return _edfa
}

// GetPreviousRevisionParser returns PdfParser for the previous version of the Pdf document.
func (_gcfde *PdfParser) GetPreviousRevisionParser() (*PdfParser, error) {
	if _gcfde._eaae == 0 {
		return nil, _a.New("\u0074\u0068\u0069\u0073 i\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0072\u0065\u0076\u0069\u0073\u0069o\u006e")
	}
	if _cgdd, _ebff := _gcfde._dgef[_gcfde]; _ebff {
		return _cgdd, nil
	}
	_eacf, _caee := _gcfde.GetPreviousRevisionReadSeeker()
	if _caee != nil {
		return nil, _caee
	}
	_dfda, _caee := NewParser(_eacf)
	_dfda._dgef = _gcfde._dgef
	if _caee != nil {
		return nil, _caee
	}
	_gcfde._dgef[_gcfde] = _dfda
	return _dfda, nil
}

func (_aagg *PdfParser) resolveReference(_dbgeae *PdfObjectReference) (PdfObject, bool, error) {
	_ggef, _bcag := _aagg.ObjCache[int(_dbgeae.ObjectNumber)]
	if _bcag {
		return _ggef, true, nil
	}
	_afad, _fgeb := _aagg.LookupByReference(*_dbgeae)
	if _fgeb != nil {
		return nil, false, _fgeb
	}
	_aagg.ObjCache[int(_dbgeae.ObjectNumber)] = _afad
	return _afad, false, nil
}

// PdfObjectReference represents the primitive PDF reference object.
type PdfObjectReference struct {
	_egcg            *PdfParser
	ObjectNumber     int64
	GenerationNumber int64
}

const (
	JB2Generic JBIG2CompressionType = iota
	JB2SymbolCorrelation
	JB2SymbolRankHaus
)

// EncodeJBIG2Image encodes 'img' into jbig2 encoded bytes stream, using default encoder settings.
func (_fbbe *JBIG2Encoder) EncodeJBIG2Image(img *JBIG2Image) ([]byte, error) {
	const _debf = "c\u006f\u0072\u0065\u002eEn\u0063o\u0064\u0065\u004a\u0042\u0049G\u0032\u0049\u006d\u0061\u0067\u0065"
	if _aaag := _fbbe.AddPageImage(img, &_fbbe.DefaultPageSettings); _aaag != nil {
		return nil, _dd.Wrap(_aaag, _debf, "")
	}
	return _fbbe.Encode()
}

func (_cge *PdfCrypt) isDecrypted(_adb PdfObject) bool {
	_, _cgg := _cge._gee[_adb]
	if _cgg {
		_df.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0064\u0065\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	switch _cff := _adb.(type) {
	case *PdfObjectStream:
		if _cge._bga.R != 5 {
			if _cdd, _eeb := _cff.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _eeb && *_cdd == "\u0058\u0052\u0065\u0066" {
				return true
			}
		}
	case *PdfIndirectObject:
		if _, _cgg = _cge._becc[int(_cff.ObjectNumber)]; _cgg {
			return true
		}
		switch _ebgd := _cff.PdfObject.(type) {
		case *PdfObjectDictionary:
			_efgb := true
			for _, _ceec := range _gbaad {
				if _ebgd.Get(_ceec) == nil {
					_efgb = false
					break
				}
			}
			if _efgb {
				return true
			}
		}
	}
	_df.Log.Trace("\u004e\u006f\u0074\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}

// UpdateParams updates the parameter values of the encoder.
func (_bcde *RawEncoder) UpdateParams(params *PdfObjectDictionary) {}

// WriteString outputs the object as it is to be written to file.
func (_cggd *PdfObjectNull) WriteString() string { return "\u006e\u0075\u006c\u006c" }

// NewFlateEncoder makes a new flate encoder with default parameters, predictor 1 and bits per component 8.
func NewFlateEncoder() *FlateEncoder {
	_cca := &FlateEncoder{}
	_cca.Predictor = 1
	_cca.BitsPerComponent = 8
	_cca.Colors = 1
	_cca.Columns = 1
	return _cca
}

var _faed _c.Map

// GetStringBytes is like GetStringVal except that it returns the string as a []byte.
// It is for convenience.
func GetStringBytes(obj PdfObject) (_cedaa []byte, _bce bool) {
	_adef, _bce := TraceToDirectObject(obj).(*PdfObjectString)
	if _bce {
		return _adef.Bytes(), true
	}
	return
}

var _eggcd = _f.MustCompile("\u0025\u0025\u0045\u004f\u0046\u003f")

// DecodeBytes decodes a multi-encoded slice of bytes by passing it through the
// DecodeBytes method of the underlying encoders.
func (_bcgd *MultiEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_edaa := encoded
	var _fdec error
	for _, _bace := range _bcgd._gbgba {
		_df.Log.Trace("\u004du\u006c\u0074i\u0020\u0045\u006e\u0063o\u0064\u0065\u0072 \u0044\u0065\u0063\u006f\u0064\u0065\u003a\u0020\u0041pp\u006c\u0079\u0069n\u0067\u0020F\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0076 \u0025\u0054", _bace, _bace)
		_edaa, _fdec = _bace.DecodeBytes(_edaa)
		if _fdec != nil {
			return nil, _fdec
		}
	}
	return _edaa, nil
}

const JB2ImageAutoThreshold = -1.0

// ReadBytesAt reads byte content at specific offset and length within the PDF.
func (_bfde *PdfParser) ReadBytesAt(offset, len int64) ([]byte, error) {
	_eccba := _bfde.GetFileOffset()
	_, _ddec := _bfde._abdga.Seek(offset, _gd.SeekStart)
	if _ddec != nil {
		return nil, _ddec
	}
	_cfdf := make([]byte, len)
	_, _ddec = _gd.ReadAtLeast(_bfde._abdga, _cfdf, int(len))
	if _ddec != nil {
		return nil, _ddec
	}
	_bfde.SetFileOffset(_eccba)
	return _cfdf, nil
}

// DecodeStream decodes a multi-encoded stream by passing it through the
// DecodeStream method of the underlying encoders.
func (_gdfg *MultiEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _gdfg.DecodeBytes(streamObj.Stream)
}

func _cebed(_geafe uint, _eddagc, _gdbg float64) float64 {
	return (_eddagc + (float64(_geafe) * (_gdbg - _eddagc) / 255)) * 255
}

// DecodeBytes decodes a slice of JBIG2 encoded bytes and returns the results.
func (_febce *JBIG2Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	return _db.DecodeBytes(encoded, _gg.Parameters{}, _febce.Globals)
}

// HeaderPosition gets the file header position.
func (_fbab ParserMetadata) HeaderPosition() int { return _fbab._ecef }

// HeaderCommentBytes gets the header comment bytes.
func (_bgaa ParserMetadata) HeaderCommentBytes() [4]byte { return _bgaa._cdef }

// DecodeBytes decodes a slice of DCT encoded bytes and returns the result.
func (_gdd *DCTEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_gcga := _d.NewReader(encoded)
	_daad, _fge := _bg.Decode(_gcga)
	if _fge != nil {
		_df.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _fge)
		return nil, _fge
	}
	_gbad := _daad.Bounds()
	_egfa := make([]byte, _gbad.Dx()*_gbad.Dy()*_gdd.ColorComponents*_gdd.BitsPerComponent/8)
	_aedb := 0
	switch _gdd.ColorComponents {
	case 1:
		_bcbaa := []float64{_gdd.Decode[0], _gdd.Decode[1]}
		for _fbfc := _gbad.Min.Y; _fbfc < _gbad.Max.Y; _fbfc++ {
			for _eccb := _gbad.Min.X; _eccb < _gbad.Max.X; _eccb++ {
				_fef := _daad.At(_eccb, _fbfc)
				if _gdd.BitsPerComponent == 16 {
					_afgb, _gfad := _fef.(_fe.Gray16)
					if !_gfad {
						return nil, _a.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_bggf := _cebed(uint(_afgb.Y>>8), _bcbaa[0], _bcbaa[1])
					_bafa := _cebed(uint(_afgb.Y), _bcbaa[0], _bcbaa[1])
					_egfa[_aedb] = byte(_bggf)
					_aedb++
					_egfa[_aedb] = byte(_bafa)
					_aedb++
				} else {
					_cdbc, _agcd := _fef.(_fe.Gray)
					if !_agcd {
						return nil, _a.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_egfa[_aedb] = byte(_cebed(uint(_cdbc.Y), _bcbaa[0], _bcbaa[1]))
					_aedb++
				}
			}
		}
	case 3:
		_fdce := []float64{_gdd.Decode[0], _gdd.Decode[1]}
		_ffcg := []float64{_gdd.Decode[2], _gdd.Decode[3]}
		_efdg := []float64{_gdd.Decode[4], _gdd.Decode[5]}
		for _ebfd := _gbad.Min.Y; _ebfd < _gbad.Max.Y; _ebfd++ {
			for _dfbf := _gbad.Min.X; _dfbf < _gbad.Max.X; _dfbf++ {
				_agebc := _daad.At(_dfbf, _ebfd)
				if _gdd.BitsPerComponent == 16 {
					_ddcg, _bdcd := _agebc.(_fe.RGBA64)
					if !_bdcd {
						return nil, _a.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_dcff := _cebed(uint(_ddcg.R>>8), _fdce[0], _fdce[1])
					_cfbd := _cebed(uint(_ddcg.R), _fdce[0], _fdce[1])
					_gaacf := _cebed(uint(_ddcg.G>>8), _ffcg[0], _ffcg[1])
					_aeg := _cebed(uint(_ddcg.G), _ffcg[0], _ffcg[1])
					_gdfd := _cebed(uint(_ddcg.B>>8), _efdg[0], _efdg[1])
					_eabd := _cebed(uint(_ddcg.B), _efdg[0], _efdg[1])
					_egfa[_aedb] = byte(_dcff)
					_aedb++
					_egfa[_aedb] = byte(_cfbd)
					_aedb++
					_egfa[_aedb] = byte(_gaacf)
					_aedb++
					_egfa[_aedb] = byte(_aeg)
					_aedb++
					_egfa[_aedb] = byte(_gdfd)
					_aedb++
					_egfa[_aedb] = byte(_eabd)
					_aedb++
				} else {
					_ecgg, _deb := _agebc.(_fe.RGBA)
					if _deb {
						_eace := _cebed(uint(_ecgg.R), _fdce[0], _fdce[1])
						_eaeg := _cebed(uint(_ecgg.G), _ffcg[0], _ffcg[1])
						_fdee := _cebed(uint(_ecgg.B), _efdg[0], _efdg[1])
						_egfa[_aedb] = byte(_eace)
						_aedb++
						_egfa[_aedb] = byte(_eaeg)
						_aedb++
						_egfa[_aedb] = byte(_fdee)
						_aedb++
					} else {
						_gbbf, _dbge := _agebc.(_fe.YCbCr)
						if !_dbge {
							return nil, _a.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
						}
						_ccbf, _dbgg, _cdcd, _ := _gbbf.RGBA()
						_cffe := _cebed(uint(_ccbf>>8), _fdce[0], _fdce[1])
						_cebd := _cebed(uint(_dbgg>>8), _ffcg[0], _ffcg[1])
						_ccbe := _cebed(uint(_cdcd>>8), _efdg[0], _efdg[1])
						_egfa[_aedb] = byte(_cffe)
						_aedb++
						_egfa[_aedb] = byte(_cebd)
						_aedb++
						_egfa[_aedb] = byte(_ccbe)
						_aedb++
					}
				}
			}
		}
	case 4:
		_bfb := []float64{_gdd.Decode[0], _gdd.Decode[1]}
		_gbc := []float64{_gdd.Decode[2], _gdd.Decode[3]}
		_abfg := []float64{_gdd.Decode[4], _gdd.Decode[5]}
		_aeba := []float64{_gdd.Decode[6], _gdd.Decode[7]}
		for _dadab := _gbad.Min.Y; _dadab < _gbad.Max.Y; _dadab++ {
			for _bdeg := _gbad.Min.X; _bdeg < _gbad.Max.X; _bdeg++ {
				_fafd := _daad.At(_bdeg, _dadab)
				_efda, _abb := _fafd.(_fe.CMYK)
				if !_abb {
					return nil, _a.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
				}
				_efgf := 255 - _cebed(uint(_efda.C), _bfb[0], _bfb[1])
				_cabd := 255 - _cebed(uint(_efda.M), _gbc[0], _gbc[1])
				_geaae := 255 - _cebed(uint(_efda.Y), _abfg[0], _abfg[1])
				_eeab := 255 - _cebed(uint(_efda.K), _aeba[0], _aeba[1])
				_egfa[_aedb] = byte(_efgf)
				_aedb++
				_egfa[_aedb] = byte(_cabd)
				_aedb++
				_egfa[_aedb] = byte(_geaae)
				_aedb++
				_egfa[_aedb] = byte(_eeab)
				_aedb++
			}
		}
	}
	return _egfa, nil
}

// HasNonConformantStream implements core.ParserMetadata.
func (_aed ParserMetadata) HasNonConformantStream() bool { return _aed._ecg }

// PdfObjectName represents the primitive PDF name object.
type PdfObjectName string

// Update updates multiple keys and returns the dictionary back so can be used in a chained fashion.
func (_gaebd *PdfObjectDictionary) Update(objmap map[string]PdfObject) *PdfObjectDictionary {
	_gaebd._gfff.Lock()
	defer _gaebd._gfff.Unlock()
	for _baba, _fegg := range objmap {
		_gaebd.setWithLock(PdfObjectName(_baba), _fegg, false)
	}
	return _gaebd
}

// EncodeBytes JPX encodes the passed in slice of bytes.
func (_cfde *JPXEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _cfde.GetFilterName())
	return data, ErrNoJPXDecode
}

// PdfCrypt provides PDF encryption/decryption support.
// The PDF standard supports encryption of strings and streams (Section 7.6).
type PdfCrypt struct {
	_dgc  encryptDict
	_bga  _gb.StdEncryptDict
	_geg  string
	_edb  []byte
	_gee  map[PdfObject]bool
	_fag  map[PdfObject]bool
	_age  bool
	_bde  cryptFilters
	_ddfb string
	_fbe  string
	_edd  *PdfParser
	_becc map[int]struct{}
}

func _ceda(_ggg _gc.Filter, _ccdb _gb.AuthEvent) *PdfObjectDictionary {
	if _ccdb == "" {
		_ccdb = _gb.EventDocOpen
	}
	_efg := MakeDict()
	_efg.Set("\u0054\u0079\u0070\u0065", MakeName("C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072"))
	_efg.Set("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et", MakeName(string(_ccdb)))
	_efg.Set("\u0043\u0046\u004d", MakeName(_ggg.Name()))
	_efg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_ggg.KeyLength())))
	return _efg
}

// NewLZWEncoder makes a new LZW encoder with default parameters.
func NewLZWEncoder() *LZWEncoder {
	_dega := &LZWEncoder{}
	_dega.Predictor = 1
	_dega.BitsPerComponent = 8
	_dega.Colors = 1
	_dega.Columns = 1
	_dega.EarlyChange = 1
	return _dega
}

// Get returns the PdfObject corresponding to the specified key.
// Returns a nil value if the key is not set.
func (_accbc *PdfObjectDictionary) Get(key PdfObjectName) PdfObject {
	_accbc._gfff.Lock()
	defer _accbc._gfff.Unlock()
	_cdfb, _ccce := _accbc._ccfa[key]
	if !_ccce {
		return nil
	}
	return _cdfb
}

// GetAsFloat64Slice returns the array as []float64 slice.
// Returns an error if not entirely numeric (only PdfObjectIntegers, PdfObjectFloats).
func (_gcdb *PdfObjectArray) GetAsFloat64Slice() ([]float64, error) {
	var _fgbc []float64
	for _, _daded := range _gcdb.Elements() {
		_abfe, _eccca := GetNumberAsFloat(TraceToDirectObject(_daded))
		if _eccca != nil {
			return nil, _ea.Errorf("\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0075m\u0062\u0065\u0072")
		}
		_fgbc = append(_fgbc, _abfe)
	}
	return _fgbc, nil
}

// ParserMetadata gets the pdf parser metadata.
func (_edeg *PdfParser) ParserMetadata() (ParserMetadata, error) {
	if !_edeg._dcad {
		return ParserMetadata{}, _ea.Errorf("\u0070\u0061\u0072\u0073\u0065r\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0072\u006be\u0064\u0020\u0066\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0064\u0065\u0074\u0061\u0069\u006c\u0065\u0064\u0020\u006d\u0065\u0074\u0061\u0064\u0061\u0074a")
	}
	return _edeg._ffge, nil
}

// GetParser returns the parser for lazy-loading or compare references.
func (_dcggc *PdfObjectReference) GetParser() *PdfParser { return _dcggc._egcg }

// EncodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_dfca *RawEncoder) EncodeBytes(data []byte) ([]byte, error) { return data, nil }

const (
	XrefTypeTableEntry   xrefType = iota
	XrefTypeObjectStream xrefType = iota
)

// String returns a string representation of `name`.
func (_cef *PdfObjectName) String() string { return string(*_cef) }

// AddEncoder adds the passed in encoder to the underlying encoder slice.
func (_aadb *MultiEncoder) AddEncoder(encoder StreamEncoder) {
	_aadb._gbgba = append(_aadb._gbgba, encoder)
}

// Encrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_bcba *PdfCrypt) Encrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _bcba.isEncrypted(obj) {
		return nil
	}
	switch _eddg := obj.(type) {
	case *PdfIndirectObject:
		_bcba._fag[_eddg] = true
		_df.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _eddg.ObjectNumber, _eddg.GenerationNumber)
		_dgda := _eddg.ObjectNumber
		_cacf := _eddg.GenerationNumber
		_bae := _bcba.Encrypt(_eddg.PdfObject, _dgda, _cacf)
		if _bae != nil {
			return _bae
		}
		return nil
	case *PdfObjectStream:
		_bcba._fag[_eddg] = true
		_geaa := _eddg.PdfObjectDictionary
		if _dgf, _aba := _geaa.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _aba && *_dgf == "\u0058\u0052\u0065\u0066" {
			return nil
		}
		_acdg := _eddg.ObjectNumber
		_efef := _eddg.GenerationNumber
		_df.Log.Trace("\u0045n\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _acdg, _efef)
		_bcc := _gfb
		if _bcba._dgc.V >= 4 {
			_bcc = _bcba._ddfb
			_df.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _bcba._ddfb)
			if _eae, _bdf := _geaa.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _bdf {
				if _dbcd, _cfa := GetName(_eae.Get(0)); _cfa {
					if *_dbcd == "\u0043\u0072\u0079p\u0074" {
						_bcc = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _eec, _adfb := _geaa.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _adfb {
							if _ddae, _gfbg := _eec.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _gfbg {
								if _, _cdf := _bcba._bde[string(*_ddae)]; _cdf {
									_df.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_ddae)
									_bcc = string(*_ddae)
								}
							}
						}
					}
				}
			}
			_df.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _bcc)
			if _bcc == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_efcg := _bcba.Encrypt(_eddg.PdfObjectDictionary, _acdg, _efef)
		if _efcg != nil {
			return _efcg
		}
		_cae, _efcg := _bcba.makeKey(_bcc, uint32(_acdg), uint32(_efef), _bcba._edb)
		if _efcg != nil {
			return _efcg
		}
		_eddg.Stream, _efcg = _bcba.encryptBytes(_eddg.Stream, _bcc, _cae)
		if _efcg != nil {
			return _efcg
		}
		_geaa.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_eddg.Stream))))
		return nil
	case *PdfObjectString:
		_df.Log.Trace("\u0045n\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_egba := _gfb
		if _bcba._dgc.V >= 4 {
			_df.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _bcba._fbe)
			if _bcba._fbe == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_egba = _bcba._fbe
		}
		_cfd, _abe := _bcba.makeKey(_egba, uint32(parentObjNum), uint32(parentGenNum), _bcba._edb)
		if _abe != nil {
			return _abe
		}
		_gcd := _eddg.Str()
		_ddaa := make([]byte, len(_gcd))
		for _bgb := 0; _bgb < len(_gcd); _bgb++ {
			_ddaa[_bgb] = _gcd[_bgb]
		}
		_df.Log.Trace("\u0045n\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _ddaa, _ddaa)
		_ddaa, _abe = _bcba.encryptBytes(_ddaa, _egba, _cfd)
		if _abe != nil {
			return _abe
		}
		_eddg._bcfef = string(_ddaa)
		return nil
	case *PdfObjectArray:
		for _, _baeg := range _eddg.Elements() {
			_dgcag := _bcba.Encrypt(_baeg, parentObjNum, parentGenNum)
			if _dgcag != nil {
				return _dgcag
			}
		}
		return nil
	case *PdfObjectDictionary:
		_aebg := false
		if _aacc := _eddg.Get("\u0054\u0079\u0070\u0065"); _aacc != nil {
			_ddfbf, _cbb := _aacc.(*PdfObjectName)
			if _cbb && *_ddfbf == "\u0053\u0069\u0067" {
				_aebg = true
			}
		}
		for _, _gaaf := range _eddg.Keys() {
			_faa := _eddg.Get(_gaaf)
			if _aebg && string(_gaaf) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_gaaf) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_gaaf) != "\u0050\u0072\u0065\u0076" && string(_gaaf) != "\u004c\u0061\u0073\u0074" {
				_gbb := _bcba.Encrypt(_faa, parentObjNum, parentGenNum)
				if _gbb != nil {
					return _gbb
				}
			}
		}
		return nil
	}
	return nil
}

// SetImage sets the image base for given flate encoder.
func (_cgac *FlateEncoder) SetImage(img *_cf.ImageBase) { _cgac._dade = img }

func (_abeaf *ASCII85Encoder) base256Tobase85(_addd uint32) [5]byte {
	_beba := [5]byte{0, 0, 0, 0, 0}
	_ceg := _addd
	for _abac := 0; _abac < 5; _abac++ {
		_bfba := uint32(1)
		for _aaaa := 0; _aaaa < 4-_abac; _aaaa++ {
			_bfba *= 85
		}
		_fdba := _ceg / _bfba
		_ceg = _ceg % _bfba
		_beba[_abac] = byte(_fdba)
	}
	return _beba
}

// PdfParser parses a PDF file and provides access to the object structure of the PDF.
type PdfParser struct {
	_bdbfe   Version
	_abdga   _gd.ReadSeeker
	_gcec    *_fd.Reader
	_gccgc   int64
	_ggaf    XrefTable
	_cdca    int64
	_dfbg    *xrefType
	_aeede   objectStreams
	_aagb    *PdfObjectDictionary
	_acg     *PdfCrypt
	_bgafa   *PdfIndirectObject
	_fbae    bool
	ObjCache objectCache
	_dbeed   map[int]bool
	_dbaad   map[int64]bool
	_ffge    ParserMetadata
	_dcad    bool
	_decff   []int64
	_eaae    int
	_cegb    bool
	_adcg    int64
	_dgef    map[*PdfParser]*PdfParser
	_bddb    []*PdfParser
}
