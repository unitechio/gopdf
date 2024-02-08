package core

import (
	_bfc "bufio"
	_fd "bytes"
	_bf "compress/lzw"
	_cea "compress/zlib"
	_bfd "crypto/md5"
	_f "crypto/rand"
	_bdc "encoding/hex"
	_c "errors"
	_gf "fmt"
	_cg "image"
	_ga "image/color"
	_eb "image/jpeg"
	_fg "io"
	_bc "reflect"
	_ce "regexp"
	_gg "sort"
	_bd "strconv"
	_gd "strings"
	_g "sync"
	_ba "time"
	_cd "unicode"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_dc "bitbucket.org/shenghui0779/gopdf/core/security"
	_efd "bitbucket.org/shenghui0779/gopdf/core/security/crypt"
	_gfa "bitbucket.org/shenghui0779/gopdf/internal/ccittfax"
	_be "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_fb "bitbucket.org/shenghui0779/gopdf/internal/jbig2"
	_ef "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_eg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder"
	_de "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_dd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_da "bitbucket.org/shenghui0779/gopdf/internal/strutils"
	_d "golang.org/x/image/tiff/lzw"
	_fgb "golang.org/x/xerrors"
)

func _egg(_fgbf *PdfObjectStream, _ade *PdfObjectDictionary) (*LZWEncoder, error) {
	_abef := NewLZWEncoder()
	_ebbc := _fgbf.PdfObjectDictionary
	if _ebbc == nil {
		return _abef, nil
	}
	if _ade == nil {
		_ageg := TraceToDirectObject(_ebbc.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _ageg != nil {
			if _eca, _ebc := _ageg.(*PdfObjectDictionary); _ebc {
				_ade = _eca
			} else if _bcfg, _ecd := _ageg.(*PdfObjectArray); _ecd {
				if _bcfg.Len() == 1 {
					if _dcf, _cfgb := GetDict(_bcfg.Get(0)); _cfgb {
						_ade = _dcf
					}
				}
			}
			if _ade == nil {
				_a.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _ageg)
				return nil, _gf.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
	}
	_gfce := _ebbc.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _gfce != nil {
		_bafg, _fccf := _gfce.(*PdfObjectInteger)
		if !_fccf {
			_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _gfce)
			return nil, _gf.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_bafg != 0 && *_bafg != 1 {
			return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_abef.EarlyChange = int(*_bafg)
	} else {
		_abef.EarlyChange = 1
	}
	if _ade == nil {
		return _abef, nil
	}
	if _babc, _cebf := GetIntVal(_ade.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")); _cebf {
		if _babc == 0 || _babc == 1 {
			_abef.EarlyChange = _babc
		} else {
			_a.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0069n\u0076\u0061\u006c\u0069\u0064\u0020E\u0061\u0072\u006c\u0079\u0043\u0068\u0061n\u0067\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020%\u0064", _babc)
		}
	}
	_gfce = _ade.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _gfce != nil {
		_abdg, _gdb := _gfce.(*PdfObjectInteger)
		if !_gdb {
			_a.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _gfce)
			return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_abef.Predictor = int(*_abdg)
	}
	_gfce = _ade.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _gfce != nil {
		_cgb, _bafb := _gfce.(*PdfObjectInteger)
		if !_bafb {
			_a.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _gf.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_abef.BitsPerComponent = int(*_cgb)
	}
	if _abef.Predictor > 1 {
		_abef.Columns = 1
		_gfce = _ade.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _gfce != nil {
			_cegf, _cbca := _gfce.(*PdfObjectInteger)
			if !_cbca {
				return nil, _gf.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_abef.Columns = int(*_cegf)
		}
		_abef.Colors = 1
		_gfce = _ade.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _gfce != nil {
			_babe, _fcdaf := _gfce.(*PdfObjectInteger)
			if !_fcdaf {
				return nil, _gf.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_abef.Colors = int(*_babe)
		}
	}
	_a.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _ade.String())
	return _abef, nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_gggfa *RunLengthEncoder) MakeStreamDict() *PdfObjectDictionary {
	_gbeg := MakeDict()
	_gbeg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_gggfa.GetFilterName()))
	return _gbeg
}

// AddPageImage adds the page with the image 'img' to the encoder context in order to encode it jbig2 document.
// The 'settings' defines what encoding type should be used by the encoder.
func (_dfga *JBIG2Encoder) AddPageImage(img *JBIG2Image, settings *JBIG2EncoderSettings) (_eaaf error) {
	const _dfca = "\u004a\u0042\u0049\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065\u0049m\u0061\u0067\u0065"
	if _dfga == nil {
		return _dd.Error(_dfca, "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u0069\u0073 n\u0069\u006c")
	}
	if settings == nil {
		settings = &_dfga.DefaultPageSettings
	}
	if _dfga._feg == nil {
		_dfga._feg = _de.InitEncodeDocument(settings.FileMode)
	}
	if _eaaf = settings.Validate(); _eaaf != nil {
		return _dd.Wrap(_eaaf, _dfca, "")
	}
	_dfbc, _eaaf := img.toBitmap()
	if _eaaf != nil {
		return _dd.Wrap(_eaaf, _dfca, "")
	}
	switch settings.Compression {
	case JB2Generic:
		if _eaaf = _dfga._feg.AddGenericPage(_dfbc, settings.DuplicatedLinesRemoval); _eaaf != nil {
			return _dd.Wrap(_eaaf, _dfca, "")
		}
	case JB2SymbolCorrelation:
		return _dd.Error(_dfca, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return _dd.Error(_dfca, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return _dd.Error(_dfca, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return nil
}
func (_ggagb *PdfParser) parseObject() (PdfObject, error) {
	_a.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_ggagb.skipSpaces()
	for {
		_gafd, _gcab := _ggagb._ffbg.Peek(2)
		if _gcab != nil {
			if _gcab != _fg.EOF || len(_gafd) == 0 {
				return nil, _gcab
			}
			if len(_gafd) == 1 {
				_gafd = append(_gafd, ' ')
			}
		}
		_a.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_gafd))
		if _gafd[0] == '/' {
			_dfbg, _aadd := _ggagb.parseName()
			_a.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _dfbg)
			return &_dfbg, _aadd
		} else if _gafd[0] == '(' {
			_a.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_dfbca, _eace := _ggagb.parseString()
			return _dfbca, _eace
		} else if _gafd[0] == '[' {
			_a.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_bffb, _fbba := _ggagb.parseArray()
			return _bffb, _fbba
		} else if (_gafd[0] == '<') && (_gafd[1] == '<') {
			_a.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			_bdcf, _gcgc := _ggagb.ParseDict()
			return _bdcf, _gcgc
		} else if _gafd[0] == '<' {
			_a.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			_dcbf, _caad := _ggagb.parseHexString()
			return _dcbf, _caad
		} else if _gafd[0] == '%' {
			_ggagb.readComment()
			_ggagb.skipSpaces()
		} else {
			_a.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_gafd, _ = _ggagb._ffbg.Peek(15)
			_bgba := string(_gafd)
			_a.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _bgba)
			if (len(_bgba) > 3) && (_bgba[:4] == "\u006e\u0075\u006c\u006c") {
				_afb, _fddc := _ggagb.parseNull()
				return &_afb, _fddc
			} else if (len(_bgba) > 4) && (_bgba[:5] == "\u0066\u0061\u006cs\u0065") {
				_baae, _eage := _ggagb.parseBool()
				return &_baae, _eage
			} else if (len(_bgba) > 3) && (_bgba[:4] == "\u0074\u0072\u0075\u0065") {
				_dcfg, _gega := _ggagb.parseBool()
				return &_dcfg, _gega
			}
			_aefeb := _fafa.FindStringSubmatch(_bgba)
			if len(_aefeb) > 1 {
				_gafd, _ = _ggagb._ffbg.ReadBytes('R')
				_a.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_gafd[:]))
				_feca, _dcdg := _gbgg(string(_gafd))
				_feca._cfada = _ggagb
				return &_feca, _dcdg
			}
			_dgae := _faaga.FindStringSubmatch(_bgba)
			if len(_dgae) > 1 {
				_a.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				_edge, _abeae := _ggagb.parseNumber()
				return _edge, _abeae
			}
			_dgae = _gadc.FindStringSubmatch(_bgba)
			if len(_dgae) > 1 {
				_a.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_a.Log.Trace("\u0025\u0020\u0073", _dgae)
				_dcdfg, _fcbaf := _ggagb.parseNumber()
				return _dcdfg, _fcbaf
			}
			_a.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _bgba)
			return nil, _c.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
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

// ResolveReference resolves reference if `o` is a *PdfObjectReference and returns the object referenced to.
// Otherwise returns back `o`.
func ResolveReference(obj PdfObject) PdfObject {
	if _bcbf, _ddac := obj.(*PdfObjectReference); _ddac {
		return _bcbf.Resolve()
	}
	return obj
}

// WriteString outputs the object as it is to be written to file.
func (_fggba *PdfObjectString) WriteString() string {
	var _fgccb _fd.Buffer
	if _fggba._gead {
		_dfda := _bdc.EncodeToString(_fggba.Bytes())
		_fgccb.WriteString("\u003c")
		_fgccb.WriteString(_dfda)
		_fgccb.WriteString("\u003e")
		return _fgccb.String()
	}
	_bgcf := map[byte]string{'\n': "\u005c\u006e", '\r': "\u005c\u0072", '\t': "\u005c\u0074", '\b': "\u005c\u0062", '\f': "\u005c\u0066", '(': "\u005c\u0028", ')': "\u005c\u0029", '\\': "\u005c\u005c"}
	_fgccb.WriteString("\u0028")
	for _bgfgd := 0; _bgfgd < len(_fggba._aaca); _bgfgd++ {
		_fcdfd := _fggba._aaca[_bgfgd]
		if _edfa, _caea := _bgcf[_fcdfd]; _caea {
			_fgccb.WriteString(_edfa)
		} else {
			_fgccb.WriteByte(_fcdfd)
		}
	}
	_fgccb.WriteString("\u0029")
	return _fgccb.String()
}

// ReadAtLeast reads at least n bytes into slice p.
// Returns the number of bytes read (should always be == n), and an error on failure.
func (_cgbf *PdfParser) ReadAtLeast(p []byte, n int) (int, error) {
	_bcdf := n
	_dagcf := 0
	_eede := 0
	for _bcdf > 0 {
		_gagg, _ccab := _cgbf._ffbg.Read(p[_dagcf:])
		if _ccab != nil {
			_a.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _gagg, _eede, _ccab.Error())
			return _dagcf, _c.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_eede++
		_dagcf += _gagg
		_bcdf -= _gagg
	}
	return _dagcf, nil
}

// String returns a string describing `array`.
func (_ceab *PdfObjectArray) String() string {
	_dfeg := "\u005b"
	for _aece, _aagb := range _ceab.Elements() {
		_dfeg += _aagb.String()
		if _aece < (_ceab.Len() - 1) {
			_dfeg += "\u002c\u0020"
		}
	}
	_dfeg += "\u005d"
	return _dfeg
}

// GetFileOffset returns the current file offset, accounting for buffered position.
func (_dadg *PdfParser) GetFileOffset() int64 {
	_bgcb, _ := _dadg._dfcdg.Seek(0, _fg.SeekCurrent)
	_bgcb -= int64(_dadg._ffbg.Buffered())
	return _bgcb
}

// GetFilterName returns the names of the underlying encoding filters,
// separated by spaces.
// Note: This is just a string, should not be used in /Filter dictionary entry. Use GetFilterArray for that.
// TODO(v4): Refactor to GetFilter() which can be used for /Filter (either Name or Array), this can be
//
//	renamed to String() as a pretty string to use in debugging etc.
func (_fabb *MultiEncoder) GetFilterName() string {
	_ccfa := ""
	for _agc, _aaac := range _fabb._bdfb {
		_ccfa += _aaac.GetFilterName()
		if _agc < len(_fabb._bdfb)-1 {
			_ccfa += "\u0020"
		}
	}
	return _ccfa
}

// ToInt64Slice returns a slice of all array elements as an int64 slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_ddae *PdfObjectArray) ToInt64Slice() ([]int64, error) {
	var _eegcd []int64
	for _, _gddg := range _ddae.Elements() {
		if _feff, _addbb := _gddg.(*PdfObjectInteger); _addbb {
			_eegcd = append(_eegcd, int64(*_feff))
		} else {
			return nil, ErrTypeError
		}
	}
	return _eegcd, nil
}

// GetAsFloat64Slice returns the array as []float64 slice.
// Returns an error if not entirely numeric (only PdfObjectIntegers, PdfObjectFloats).
func (_egfac *PdfObjectArray) GetAsFloat64Slice() ([]float64, error) {
	var _ecca []float64
	for _, _gagee := range _egfac.Elements() {
		_bffba, _gbbf := GetNumberAsFloat(TraceToDirectObject(_gagee))
		if _gbbf != nil {
			return nil, _gf.Errorf("\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0075m\u0062\u0065\u0072")
		}
		_ecca = append(_ecca, _bffba)
	}
	return _ecca, nil
}

// GetFilterName returns the name of the encoding filter.
func (_ceaa *ASCIIHexEncoder) GetFilterName() string { return StreamEncodingFilterNameASCIIHex }

// DecodeBytes decodes a slice of Flate encoded bytes and returns the result.
func (_afa *FlateEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_a.Log.Trace("\u0046\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065\u0020b\u0079\u0074\u0065\u0073")
	if len(encoded) == 0 {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0046\u006c\u0061\u0074\u0065 e\u006ec\u006f\u0064\u0065\u0064\u0020\u0062\u0075\u0066\u0066\u0065\u0072\u002e \u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0065\u006d\u0070\u0074\u0079\u0020\u0062y\u0074\u0065\u0020\u0073\u006c\u0069\u0063\u0065\u002e")
		return []byte{}, nil
	}
	_afg := _fd.NewReader(encoded)
	_deag, _cbceg := _cea.NewReader(_afg)
	if _cbceg != nil {
		_a.Log.Debug("\u0044e\u0063o\u0064\u0069\u006e\u0067\u0020e\u0072\u0072o\u0072\u0020\u0025\u0076\u000a", _cbceg)
		_a.Log.Debug("\u0053t\u0072e\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(encoded), encoded)
		return nil, _cbceg
	}
	defer _deag.Close()
	var _eedf _fd.Buffer
	_eedf.ReadFrom(_deag)
	return _eedf.Bytes(), nil
}
func (_eaccf *PdfCrypt) generateParams(_bbb, _geg []byte) error {
	_gbc := _eaccf.securityHandler()
	_dea, _gedd := _gbc.GenerateParams(&_eaccf._aec, _geg, _bbb)
	if _gedd != nil {
		return _gedd
	}
	_eaccf._baa = _dea
	return nil
}
func (_edefc *PdfParser) parseBool() (PdfObjectBool, error) {
	_bdee, _eagce := _edefc._ffbg.Peek(4)
	if _eagce != nil {
		return PdfObjectBool(false), _eagce
	}
	if (len(_bdee) >= 4) && (string(_bdee[:4]) == "\u0074\u0072\u0075\u0065") {
		_edefc._ffbg.Discard(4)
		return PdfObjectBool(true), nil
	}
	_bdee, _eagce = _edefc._ffbg.Peek(5)
	if _eagce != nil {
		return PdfObjectBool(false), _eagce
	}
	if (len(_bdee) >= 5) && (string(_bdee[:5]) == "\u0066\u0061\u006cs\u0065") {
		_edefc._ffbg.Discard(5)
		return PdfObjectBool(false), nil
	}
	return PdfObjectBool(false), _c.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// Append appends PdfObject(s) to the array.
func (_dgaeca *PdfObjectArray) Append(objects ...PdfObject) {
	if _dgaeca == nil {
		_a.Log.Debug("\u0057\u0061\u0072\u006e\u0020\u002d\u0020\u0041\u0074\u0074\u0065\u006d\u0070t\u0020\u0074\u006f\u0020\u0061\u0070p\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020a\u0072\u0072\u0061\u0079")
		return
	}
	_dgaeca._fabc = append(_dgaeca._fabc, objects...)
}
func _adcf(_baaf *PdfObjectStream, _afgaf *PdfObjectDictionary) (*CCITTFaxEncoder, error) {
	_fada := NewCCITTFaxEncoder()
	_abg := _baaf.PdfObjectDictionary
	if _abg == nil {
		return _fada, nil
	}
	if _afgaf == nil {
		_cafd := TraceToDirectObject(_abg.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		if _cafd != nil {
			switch _ebgf := _cafd.(type) {
			case *PdfObjectDictionary:
				_afgaf = _ebgf
			case *PdfObjectArray:
				if _ebgf.Len() == 1 {
					if _fcff, _aage := GetDict(_ebgf.Get(0)); _aage {
						_afgaf = _fcff
					}
				}
			default:
				_a.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074 \u0061 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0025\u0023\u0076", _cafd)
				return nil, _c.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
		}
		if _afgaf == nil {
			_a.Log.Error("\u0044\u0065c\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064 %\u0023\u0076", _cafd)
			return nil, _c.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _ace, _aegg := GetNumberAsInt64(_afgaf.Get("\u004b")); _aegg == nil {
		_fada.K = int(_ace)
	}
	if _adfb, _fdbgf := GetNumberAsInt64(_afgaf.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _fdbgf == nil {
		_fada.Columns = int(_adfb)
	} else {
		_fada.Columns = 1728
	}
	if _dede, _ccbd := GetNumberAsInt64(_afgaf.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _ccbd == nil {
		_fada.BlackIs1 = _dede > 0
	} else {
		if _cece, _ecgg := GetBoolVal(_afgaf.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _ecgg {
			_fada.BlackIs1 = _cece
		} else {
			if _gcga, _ecegfd := GetArray(_afgaf.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _ecegfd {
				_ffe, _eaed := _gcga.ToIntegerArray()
				if _eaed == nil {
					_fada.BlackIs1 = _ffe[0] == 1 && _ffe[1] == 0
				}
			}
		}
	}
	if _bgcc, _cfab := GetNumberAsInt64(_afgaf.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _cfab == nil {
		_fada.EncodedByteAlign = _bgcc > 0
	} else {
		if _ccde, _feaa := GetBoolVal(_afgaf.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _feaa {
			_fada.EncodedByteAlign = _ccde
		}
	}
	if _ddcd, _bbeg := GetNumberAsInt64(_afgaf.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _bbeg == nil {
		_fada.EndOfLine = _ddcd > 0
	} else {
		if _bcgc, _gcdd := GetBoolVal(_afgaf.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _gcdd {
			_fada.EndOfLine = _bcgc
		}
	}
	if _cdgee, _ebea := GetNumberAsInt64(_afgaf.Get("\u0052\u006f\u0077\u0073")); _ebea == nil {
		_fada.Rows = int(_cdgee)
	}
	_fada.EndOfBlock = true
	if _eegc, _eccc := GetNumberAsInt64(_afgaf.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _eccc == nil {
		_fada.EndOfBlock = _eegc > 0
	} else {
		if _ggdg, _cce := GetBoolVal(_afgaf.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _cce {
			_fada.EndOfBlock = _ggdg
		}
	}
	if _ccc, _aafc := GetNumberAsInt64(_afgaf.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _aafc != nil {
		_fada.DamagedRowsBeforeError = int(_ccc)
	}
	_a.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _afgaf.String())
	return _fada, nil
}

var _gfcec _g.Map

// UpdateParams updates the parameter values of the encoder.
func (_ggfbd *JPXEncoder) UpdateParams(params *PdfObjectDictionary) {}
func (_gfac *ASCII85Encoder) base256Tobase85(_dabd uint32) [5]byte {
	_afd := [5]byte{0, 0, 0, 0, 0}
	_cedg := _dabd
	for _baced := 0; _baced < 5; _baced++ {
		_eaab := uint32(1)
		for _bgfg := 0; _bgfg < 4-_baced; _bgfg++ {
			_eaab *= 85
		}
		_eccd := _cedg / _eaab
		_cedg = _cedg % _eaab
		_afd[_baced] = byte(_eccd)
	}
	return _afd
}

// XrefTable represents the cross references in a PDF, i.e. the table of objects and information
// where to access within the PDF file.
type XrefTable struct {
	ObjectMap map[int]XrefObject
	_cf       []XrefObject
}

// GetArray returns the *PdfObjectArray represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetArray(obj PdfObject) (_ageb *PdfObjectArray, _dbed bool) {
	_ageb, _dbed = TraceToDirectObject(obj).(*PdfObjectArray)
	return _ageb, _dbed
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_aeg *FlateEncoder) MakeDecodeParams() PdfObject {
	if _aeg.Predictor > 1 {
		_abeb := MakeDict()
		_abeb.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_aeg.Predictor)))
		if _aeg.BitsPerComponent != 8 {
			_abeb.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_aeg.BitsPerComponent)))
		}
		if _aeg.Columns != 1 {
			_abeb.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_aeg.Columns)))
		}
		if _aeg.Colors != 1 {
			_abeb.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_aeg.Colors)))
		}
		return _abeb
	}
	return nil
}

// HasInvalidSubsectionHeader implements core.ParserMetadata interface.
func (_bgfc ParserMetadata) HasInvalidSubsectionHeader() bool { return _bgfc._gfe }

// FlattenObject returns the contents of `obj`. In other words, `obj` with indirect objects replaced
// by their values.
// The replacements are made recursively to a depth of traceMaxDepth.
// NOTE: Dicts are sorted to make objects with same contents have the same PDF object strings.
func FlattenObject(obj PdfObject) PdfObject { return _gbbga(obj, 0) }

// Resolve resolves the reference and returns the indirect or stream object.
// If the reference cannot be resolved, a *PdfObjectNull object is returned.
func (_gfaff *PdfObjectReference) Resolve() PdfObject {
	if _gfaff._cfada == nil {
		return MakeNull()
	}
	_gfgbb, _, _cagb := _gfaff._cfada.resolveReference(_gfaff)
	if _cagb != nil {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0072\u0065\u0073\u006f\u006cv\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065r\u0065n\u0063\u0065\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074", _cagb)
		return MakeNull()
	}
	if _gfgbb == nil {
		_a.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0072\u0065\u0073ol\u0076\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065:\u0020\u006ei\u006c\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0061\u0020nu\u006c\u006c\u0020o\u0062\u006a\u0065\u0063\u0074")
		return MakeNull()
	}
	return _gfgbb
}

// DecodeStream decodes a DCT encoded stream and returns the result as a
// slice of bytes.
func (_ebbd *DCTEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _ebbd.DecodeBytes(streamObj.Stream)
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on the current encoder settings.
func (_gdab *JBIG2Encoder) MakeDecodeParams() PdfObject { return MakeDict() }

// NewCCITTFaxEncoder makes a new CCITTFax encoder.
func NewCCITTFaxEncoder() *CCITTFaxEncoder { return &CCITTFaxEncoder{Columns: 1728, EndOfBlock: true} }

// UpdateParams updates the parameter values of the encoder.
func (_abdd *RawEncoder) UpdateParams(params *PdfObjectDictionary) {}

// HasInvalidHexRunes implements core.ParserMetadata interface.
func (_fef ParserMetadata) HasInvalidHexRunes() bool { return _fef._cde }
func (_ccac *PdfParser) parseXrefTable() (*PdfObjectDictionary, error) {
	var _cedc *PdfObjectDictionary
	_dggf, _geba := _ccac.readTextLine()
	if _geba != nil {
		return nil, _geba
	}
	if _ccac._ecgd && _gd.Count(_gd.TrimPrefix(_dggf, "\u0078\u0072\u0065\u0066"), "\u0020") > 0 {
		_ccac._gadge._bdf = true
	}
	_a.Log.Trace("\u0078\u0072\u0065\u0066 f\u0069\u0072\u0073\u0074\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _dggf)
	_edgec := -1
	_aggg := 0
	_fcdg := false
	_aabc := ""
	for {
		_ccac.skipSpaces()
		_, _bfeb := _ccac._ffbg.Peek(1)
		if _bfeb != nil {
			return nil, _bfeb
		}
		_dggf, _bfeb = _ccac.readTextLine()
		if _bfeb != nil {
			return nil, _bfeb
		}
		_fedc := _cac.FindStringSubmatch(_dggf)
		if len(_fedc) == 0 {
			_dbgg := len(_aabc) > 0
			_aabc += _dggf + "\u000a"
			if _dbgg {
				_fedc = _cac.FindStringSubmatch(_aabc)
			}
		}
		if len(_fedc) == 3 {
			if _ccac._ecgd && !_ccac._gadge._gfe {
				var (
					_gafa bool
					_fdag int
				)
				for _, _bcfb := range _dggf {
					if _cd.IsDigit(_bcfb) {
						if _gafa {
							break
						}
						continue
					}
					if !_gafa {
						_gafa = true
					}
					_fdag++
				}
				if _fdag > 1 {
					_ccac._gadge._gfe = true
				}
			}
			_dabf, _ := _bd.Atoi(_fedc[1])
			_dfaa, _ := _bd.Atoi(_fedc[2])
			_edgec = _dabf
			_aggg = _dfaa
			_fcdg = true
			_aabc = ""
			_a.Log.Trace("\u0078r\u0065\u0066 \u0073\u0075\u0062s\u0065\u0063\u0074\u0069\u006f\u006e\u003a \u0066\u0069\u0072\u0073\u0074\u0020o\u0062\u006a\u0065\u0063\u0074\u003a\u0020\u0025\u0064\u0020\u006fb\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _edgec, _aggg)
			continue
		}
		_dfd := _fgdd.FindStringSubmatch(_dggf)
		if len(_dfd) == 4 {
			if !_fcdg {
				_a.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0058r\u0065\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006fr\u006da\u0074\u0021\u000a")
				return nil, _c.New("\u0078\u0072\u0065\u0066 i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_fege, _ := _bd.ParseInt(_dfd[1], 10, 64)
			_fdbc, _ := _bd.Atoi(_dfd[2])
			_cdacd := _dfd[3]
			_aabc = ""
			if _gd.ToLower(_cdacd) == "\u006e" && _fege > 1 {
				_gadcg, _dgaec := _ccac._bbdf.ObjectMap[_edgec]
				if !_dgaec || _fdbc > _gadcg.Generation {
					_ebeb := XrefObject{ObjectNumber: _edgec, XType: XrefTypeTableEntry, Offset: _fege, Generation: _fdbc}
					_ccac._bbdf.ObjectMap[_edgec] = _ebeb
				}
			}
			_edgec++
			continue
		}
		if (len(_dggf) > 6) && (_dggf[:7] == "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_a.Log.Trace("\u0046o\u0075n\u0064\u0020\u0074\u0072\u0061i\u006c\u0065r\u0020\u002d\u0020\u0025\u0073", _dggf)
			if len(_dggf) > 9 {
				_gffdc := _ccac.GetFileOffset()
				_ccac.SetFileOffset(_gffdc - int64(len(_dggf)) + 7)
			}
			_ccac.skipSpaces()
			_ccac.skipComments()
			_a.Log.Trace("R\u0065\u0061\u0064\u0069ng\u0020t\u0072\u0061\u0069\u006c\u0065r\u0020\u0064\u0069\u0063\u0074\u0021")
			_a.Log.Trace("\u0070\u0065\u0065\u006b\u003a\u0020\u0022\u0025\u0073\u0022", _dggf)
			_cedc, _bfeb = _ccac.ParseDict()
			_a.Log.Trace("\u0045O\u0046\u0020\u0072\u0065a\u0064\u0069\u006e\u0067\u0020t\u0072a\u0069l\u0065\u0072\u0020\u0064\u0069\u0063\u0074!")
			if _bfeb != nil {
				_a.Log.Debug("\u0045\u0072\u0072o\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _bfeb)
				return nil, _bfeb
			}
			break
		}
		if _dggf == "\u0025\u0025\u0045O\u0046" {
			_a.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u006e\u0064 \u006f\u0066\u0020\u0066\u0069\u006c\u0065 -\u0020\u0074\u0072\u0061i\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066ou\u006e\u0064 \u002d\u0020\u0065\u0072\u0072\u006f\u0072\u0021")
			return nil, _c.New("\u0065\u006e\u0064 \u006f\u0066\u0020\u0066i\u006c\u0065\u0020\u002d\u0020\u0074\u0072a\u0069\u006c\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_a.Log.Trace("\u0078\u0072\u0065\u0066\u0020\u006d\u006f\u0072\u0065 \u003a\u0020\u0025\u0073", _dggf)
	}
	_a.Log.Trace("\u0045\u004f\u0046 p\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0021")
	if _ccac._fddf == nil {
		_feda := XrefTypeTableEntry
		_ccac._fddf = &_feda
	}
	return _cedc, nil
}

const _edg = "\u0053\u0074\u0064C\u0046"

// DecodeStream returns the passed in stream as a slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_dgcg *RawEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return streamObj.Stream, nil
}
func (_fgdg *PdfParser) lookupByNumberWrapper(_eaa int, _faa bool) (PdfObject, bool, error) {
	_fdg, _cda, _ed := _fgdg.lookupByNumber(_eaa, _faa)
	if _ed != nil {
		return nil, _cda, _ed
	}
	if !_cda && _fgdg._abae != nil && _fgdg._abae._dfe && !_fgdg._abae.isDecrypted(_fdg) {
		_cfc := _fgdg._abae.Decrypt(_fdg, 0, 0)
		if _cfc != nil {
			return nil, _cda, _cfc
		}
	}
	return _fdg, _cda, nil
}

// MakeStringFromBytes creates an PdfObjectString from a byte array.
// This is more natural than MakeString as `data` is usually not utf-8 encoded.
func MakeStringFromBytes(data []byte) *PdfObjectString { return MakeString(string(data)) }

// GetFilterName returns the name of the encoding filter.
func (_acfg *LZWEncoder) GetFilterName() string { return StreamEncodingFilterNameLZW }
func _gdce(_cgaa string) (int, int, error) {
	_ggce := _ebceg.FindStringSubmatch(_cgaa)
	if len(_ggce) < 3 {
		return 0, 0, _c.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_aafd, _ := _bd.Atoi(_ggce[1])
	_fdecg, _ := _bd.Atoi(_ggce[2])
	return _aafd, _fdecg, nil
}
func (_gbedf *PdfObjectFloat) String() string { return _gf.Sprintf("\u0025\u0066", *_gbedf) }

// JBIG2CompressionType defines the enum compression type used by the JBIG2Encoder.
type JBIG2CompressionType int

// ParserMetadata gets the pdf parser metadata.
func (_caee *PdfParser) ParserMetadata() (ParserMetadata, error) {
	if !_caee._ecgd {
		return ParserMetadata{}, _gf.Errorf("\u0070\u0061\u0072\u0073\u0065r\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0072\u006be\u0064\u0020\u0066\u006f\u0072\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0064\u0065\u0074\u0061\u0069\u006c\u0065\u0064\u0020\u006d\u0065\u0074\u0061\u0064\u0061\u0074a")
	}
	return _caee._gadge, nil
}

// String returns a string describing `streams`.
func (_afdeb *PdfObjectStreams) String() string {
	return _gf.Sprintf("\u004f\u0062j\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0025\u0064", _afdeb.ObjectNumber)
}

// Resolve resolves a PdfObject to direct object, looking up and resolving references as needed (unlike TraceToDirect).
func (_fgc *PdfParser) Resolve(obj PdfObject) (PdfObject, error) {
	_ag, _gfd := obj.(*PdfObjectReference)
	if !_gfd {
		return obj, nil
	}
	_fba := _fgc.GetFileOffset()
	defer func() { _fgc.SetFileOffset(_fba) }()
	_ggg, _add := _fgc.LookupByReference(*_ag)
	if _add != nil {
		return nil, _add
	}
	_ddb, _debb := _ggg.(*PdfIndirectObject)
	if !_debb {
		return _ggg, nil
	}
	_ggg = _ddb.PdfObject
	_, _gfd = _ggg.(*PdfObjectReference)
	if _gfd {
		return _ddb, _c.New("\u006d\u0075lt\u0069\u0020\u0064e\u0070\u0074\u0068\u0020tra\u0063e \u0070\u006f\u0069\u006e\u0074\u0065\u0072 t\u006f\u0020\u0070\u006f\u0069\u006e\u0074e\u0072")
	}
	return _ggg, nil
}

// GetRevision returns PdfParser for the specific version of the Pdf document.
func (_ggdcd *PdfParser) GetRevision(revisionNumber int) (*PdfParser, error) {
	_dbbbb := _ggdcd._efce
	if _dbbbb == revisionNumber {
		return _ggdcd, nil
	}
	if _dbbbb < revisionNumber {
		return nil, _c.New("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0072\u0065\u0076\u0069\u0073i\u006fn\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e")
	}
	if _ggdcd._gfdfd[revisionNumber] != nil {
		return _ggdcd._gfdfd[revisionNumber], nil
	}
	_gfgb := _ggdcd
	for ; _dbbbb > revisionNumber; _dbbbb-- {
		_deeg, _ecaa := _gfgb.GetPreviousRevisionParser()
		if _ecaa != nil {
			return nil, _ecaa
		}
		_ggdcd._gfdfd[_dbbbb-1] = _deeg
		_ggdcd._beaf[_gfgb] = _deeg
		_gfgb = _deeg
	}
	return _gfgb, nil
}
func (_dcgad *PdfParser) parseNumber() (PdfObject, error) { return ParseNumber(_dcgad._ffbg) }

const _bca = 32 << (^uint(0) >> 63)

func (_ac *PdfParser) lookupObjectViaOS(_efg int, _dcc int) (PdfObject, error) {
	var _cdd *_fd.Reader
	var _fcg objectStream
	var _cc bool
	_fcg, _cc = _ac._fcba[_efg]
	if !_cc {
		_caf, _fcd := _ac.LookupByNumber(_efg)
		if _fcd != nil {
			_a.Log.Debug("\u004d\u0069ss\u0069\u006e\u0067 \u006f\u0062\u006a\u0065ct \u0073tr\u0065\u0061\u006d\u0020\u0077\u0069\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u0020%\u0064", _efg)
			return nil, _fcd
		}
		_cdg, _fe := _caf.(*PdfObjectStream)
		if !_fe {
			return nil, _c.New("i\u006e\u0076\u0061\u006cid\u0020o\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _ac._abae != nil && !_ac._abae.isDecrypted(_cdg) {
			return nil, _c.New("\u006e\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0064\u0065\u0063r\u0079\u0070\u0074\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d")
		}
		_ec := _cdg.PdfObjectDictionary
		_a.Log.Trace("\u0073o\u0020\u0064\u003a\u0020\u0025\u0073\n", _ec.String())
		_cae, _fe := _ec.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
		if !_fe {
			_a.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0061\u006c\u0077\u0061\u0079\u0073\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0054\u0079\u0070\u0065")
			return nil, _c.New("\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020T\u0079\u0070\u0065")
		}
		if _gd.ToLower(string(*_cae)) != "\u006f\u0062\u006a\u0073\u0074\u006d" {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0074\u0079\u0070\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0077\u0061\u0079\u0073 \u0062\u0065\u0020\u004f\u0062\u006a\u0053\u0074\u006d\u0020\u0021")
			return nil, _c.New("\u006f\u0062\u006a\u0065c\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0074y\u0070e\u0020\u0021\u003d\u0020\u004f\u0062\u006aS\u0074\u006d")
		}
		N, _fe := _ec.Get("\u004e").(*PdfObjectInteger)
		if !_fe {
			return nil, _c.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u004e\u0020i\u006e\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_cgg, _fe := _ec.Get("\u0046\u0069\u0072s\u0074").(*PdfObjectInteger)
		if !_fe {
			return nil, _c.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0046\u0069\u0072\u0073\u0074\u0020i\u006e \u0073t\u0072e\u0061\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
		}
		_a.Log.Trace("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066 \u006f\u0062\u006a\u0065\u0063\u0074\u0073\u003a\u0020\u0025\u0064", _cae, *N)
		_fgd, _fcd := DecodeStream(_cdg)
		if _fcd != nil {
			return nil, _fcd
		}
		_a.Log.Trace("D\u0065\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0073", _fgd)
		_ccb := _ac.GetFileOffset()
		defer func() { _ac.SetFileOffset(_ccb) }()
		_cdd = _fd.NewReader(_fgd)
		_ac._ffbg = _bfc.NewReader(_cdd)
		_a.Log.Trace("\u0050a\u0072s\u0069\u006e\u0067\u0020\u006ff\u0066\u0073e\u0074\u0020\u006d\u0061\u0070")
		_df := map[int]int64{}
		for _ggf := 0; _ggf < int(*N); _ggf++ {
			_ac.skipSpaces()
			_cb, _bdb := _ac.parseNumber()
			if _bdb != nil {
				return nil, _bdb
			}
			_fa, _bb := _cb.(*PdfObjectInteger)
			if !_bb {
				return nil, _c.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_ac.skipSpaces()
			_cb, _bdb = _ac.parseNumber()
			if _bdb != nil {
				return nil, _bdb
			}
			_ea, _bb := _cb.(*PdfObjectInteger)
			if !_bb {
				return nil, _c.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073t\u0072e\u0061m\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0074\u0061\u0062\u006c\u0065")
			}
			_a.Log.Trace("\u006f\u0062j\u0020\u0025\u0064 \u006f\u0066\u0066\u0073\u0065\u0074\u0020\u0025\u0064", *_fa, *_ea)
			_df[int(*_fa)] = int64(*_cgg + *_ea)
		}
		_fcg = objectStream{N: int(*N), _fc: _fgd, _ca: _df}
		_ac._fcba[_efg] = _fcg
	} else {
		_deb := _ac.GetFileOffset()
		defer func() { _ac.SetFileOffset(_deb) }()
		_cdd = _fd.NewReader(_fcg._fc)
		_ac._ffbg = _bfc.NewReader(_cdd)
	}
	_bae := _fcg._ca[_dcc]
	_a.Log.Trace("\u0041\u0043\u0054\u0055AL\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u005b\u0025\u0064\u005d\u0020\u003d\u0020%\u0064", _dcc, _bae)
	_cdd.Seek(_bae, _fg.SeekStart)
	_ac._ffbg = _bfc.NewReader(_cdd)
	_gb, _ := _ac._ffbg.Peek(100)
	_a.Log.Trace("\u004f\u0042\u004a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_gb))
	_fac, _dg := _ac.parseObject()
	if _dg != nil {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0072\u0065\u0061\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0028\u0025\u0073\u0029", _dg)
		return nil, _dg
	}
	if _fac == nil {
		return nil, _c.New("o\u0062\u006a\u0065\u0063t \u0063a\u006e\u006e\u006f\u0074\u0020b\u0065\u0020\u006e\u0075\u006c\u006c")
	}
	_dgb := PdfIndirectObject{}
	_dgb.ObjectNumber = int64(_dcc)
	_dgb.PdfObject = _fac
	_dgb._cfada = _ac
	return &_dgb, nil
}

// ASCIIHexEncoder implements ASCII hex encoder/decoder.
type ASCIIHexEncoder struct{}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_eead *CCITTFaxEncoder) MakeStreamDict() *PdfObjectDictionary {
	_fdbb := MakeDict()
	_fdbb.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_eead.GetFilterName()))
	_fdbb.SetIfNotNil("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _eead.MakeDecodeParams())
	return _fdbb
}
func (_ddfe *PdfParser) parsePdfVersion() (int, int, error) {
	var _baff int64 = 20
	_ffbda := make([]byte, _baff)
	_ddfe._dfcdg.Seek(0, _fg.SeekStart)
	_ddfe._dfcdg.Read(_ffbda)
	var _fdcgb error
	var _gagca, _efceg int
	if _eabf := _dbef.FindStringSubmatch(string(_ffbda)); len(_eabf) < 3 {
		if _gagca, _efceg, _fdcgb = _ddfe.seekPdfVersionTopDown(); _fdcgb != nil {
			_a.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _fdcgb
		}
		_ddfe._dfcdg, _fdcgb = _dbaeec(_ddfe._dfcdg, _ddfe.GetFileOffset()-8)
		if _fdcgb != nil {
			return 0, 0, _fdcgb
		}
	} else {
		if _gagca, _fdcgb = _bd.Atoi(_eabf[1]); _fdcgb != nil {
			return 0, 0, _fdcgb
		}
		if _efceg, _fdcgb = _bd.Atoi(_eabf[2]); _fdcgb != nil {
			return 0, 0, _fdcgb
		}
		_ddfe.SetFileOffset(0)
	}
	_ddfe._ffbg = _bfc.NewReader(_ddfe._dfcdg)
	_a.Log.Debug("\u0050\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _gagca, _efceg)
	return _gagca, _efceg, nil
}

// ToGoImage converts the JBIG2Image to the golang image.Image.
func (_ebf *JBIG2Image) ToGoImage() (_cg.Image, error) {
	const _gfcea = "J\u0042I\u0047\u0032\u0049\u006d\u0061\u0067\u0065\u002eT\u006f\u0047\u006f\u0049ma\u0067\u0065"
	if _ebf.Data == nil {
		return nil, _dd.Error(_gfcea, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _ebf.Width == 0 || _ebf.Height == 0 {
		return nil, _dd.Error(_gfcea, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	_ebaf, _acdab := _be.NewImage(_ebf.Width, _ebf.Height, 1, 1, _ebf.Data, nil, nil)
	if _acdab != nil {
		return nil, _acdab
	}
	return _ebaf, nil
}

// DecodeBytes decodes the CCITTFax encoded image data.
func (_bbec *CCITTFaxEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_fefc, _gfcd := _gfa.NewDecoder(encoded, _gfa.DecodeOptions{Columns: _bbec.Columns, Rows: _bbec.Rows, K: _bbec.K, EncodedByteAligned: _bbec.EncodedByteAlign, BlackIsOne: _bbec.BlackIs1, EndOfBlock: _bbec.EndOfBlock, EndOfLine: _bbec.EndOfLine, DamagedRowsBeforeError: _bbec.DamagedRowsBeforeError})
	if _gfcd != nil {
		return nil, _gfcd
	}
	_bcfd, _gfcd := _fg.ReadAll(_fefc)
	if _gfcd != nil {
		return nil, _gfcd
	}
	return _bcfd, nil
}

// GetFilterName returns the name of the encoding filter.
func (_afcag *CCITTFaxEncoder) GetFilterName() string { return StreamEncodingFilterNameCCITTFax }
func (_cfdd *PdfParser) parseLinearizedDictionary() (*PdfObjectDictionary, error) {
	_afdg, _fdde := _cfdd._dfcdg.Seek(0, _fg.SeekEnd)
	if _fdde != nil {
		return nil, _fdde
	}
	var _bacc int64
	var _dabfe int64 = 2048
	for _bacc < _afdg-4 {
		if _afdg <= (_dabfe + _bacc) {
			_dabfe = _afdg - _bacc
		}
		_, _edag := _cfdd._dfcdg.Seek(_bacc, _fg.SeekStart)
		if _edag != nil {
			return nil, _edag
		}
		_cgddb := make([]byte, _dabfe)
		_, _edag = _cfdd._dfcdg.Read(_cgddb)
		if _edag != nil {
			return nil, _edag
		}
		_a.Log.Trace("\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0066i\u0072\u0073\u0074\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0022\u0025\u0073\u0022", string(_cgddb))
		_dcfb := _ebceg.FindAllStringIndex(string(_cgddb), -1)
		if _dcfb != nil {
			_fded := _dcfb[0]
			_a.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _dcfb)
			_, _ecbef := _cfdd._dfcdg.Seek(int64(_fded[0]), _fg.SeekStart)
			if _ecbef != nil {
				return nil, _ecbef
			}
			_cfdd._ffbg = _bfc.NewReader(_cfdd._dfcdg)
			_bdac, _ecbef := _cfdd.ParseIndirectObject()
			if _ecbef != nil {
				return nil, nil
			}
			if _aegge, _dfcg := GetIndirect(_bdac); _dfcg {
				if _eabfg, _bgcdb := GetDict(_aegge.PdfObject); _bgcdb {
					if _eeeb := _eabfg.Get("\u004c\u0069\u006e\u0065\u0061\u0072\u0069\u007a\u0065\u0064"); _eeeb != nil {
						return _eabfg, nil
					}
					return nil, nil
				}
			}
			return nil, nil
		}
		_bacc += _dabfe - 4
	}
	return nil, _c.New("\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u006fb\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064")
}

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
func ParseNumber(buf *_bfc.Reader) (PdfObject, error) {
	_baafd := false
	_fedac := true
	var _bcacf _fd.Buffer
	for {
		if _a.Log.IsLogLevel(_a.LogLevelTrace) {
			_a.Log.Trace("\u0050\u0061\u0072\u0073in\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0022\u0025\u0073\u0022", _bcacf.String())
		}
		_befag, _bbdaf := buf.Peek(1)
		if _bbdaf == _fg.EOF {
			break
		}
		if _bbdaf != nil {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0025\u0073", _bbdaf)
			return nil, _bbdaf
		}
		if _fedac && (_befag[0] == '-' || _befag[0] == '+') {
			_adga, _ := buf.ReadByte()
			_bcacf.WriteByte(_adga)
			_fedac = false
		} else if IsDecimalDigit(_befag[0]) {
			_ecda, _ := buf.ReadByte()
			_bcacf.WriteByte(_ecda)
		} else if _befag[0] == '.' {
			_eaaae, _ := buf.ReadByte()
			_bcacf.WriteByte(_eaaae)
			_baafd = true
		} else if _befag[0] == 'e' || _befag[0] == 'E' {
			_dafd, _ := buf.ReadByte()
			_bcacf.WriteByte(_dafd)
			_baafd = true
			_fedac = true
		} else {
			break
		}
	}
	var _gceg PdfObject
	if _baafd {
		_eefd, _cfaa := _bd.ParseFloat(_bcacf.String(), 64)
		if _cfaa != nil {
			_a.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025v\u0020\u0065\u0072\u0072\u003d\u0025v\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0030\u002e\u0020\u004fu\u0074\u0070u\u0074\u0020\u006d\u0061y\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074", _bcacf.String(), _cfaa)
			_eefd = 0.0
		}
		_aabb := PdfObjectFloat(_eefd)
		_gceg = &_aabb
	} else {
		_bbab, _eegee := _bd.ParseInt(_bcacf.String(), 10, 64)
		if _eegee != nil {
			_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0025\u0076\u0020\u0065\u0072\u0072\u003d%\u0076\u002e\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0030\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 \u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006ec\u006f\u0072r\u0065c\u0074", _bcacf.String(), _eegee)
			_bbab = 0
		}
		_aadbb := PdfObjectInteger(_bbab)
		_gceg = &_aadbb
	}
	return _gceg, nil
}

var _cgcd = _ce.MustCompile("\u005c\u0073\u002a\u0078\u0072\u0065\u0066\u005c\u0073\u002a")

func _bgaa(_bbcd _fg.ReadSeeker, _cddb int64) (*limitedReadSeeker, error) {
	_, _ddf := _bbcd.Seek(0, _fg.SeekStart)
	if _ddf != nil {
		return nil, _ddf
	}
	return &limitedReadSeeker{_fbge: _bbcd, _bcgd: _cddb}, nil
}

// DecodeStream implements ASCII hex decoding.
func (_fdd *ASCIIHexEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _fdd.DecodeBytes(streamObj.Stream)
}

// GetFloatVal returns the float64 value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetFloatVal(obj PdfObject) (_gbca float64, _cfdc bool) {
	_cadfc, _cfdc := TraceToDirectObject(obj).(*PdfObjectFloat)
	if _cfdc {
		return float64(*_cadfc), true
	}
	return 0, false
}

// IsWhiteSpace checks if byte represents a white space character.
func IsWhiteSpace(ch byte) bool {
	if (ch == 0x00) || (ch == 0x09) || (ch == 0x0A) || (ch == 0x0C) || (ch == 0x0D) || (ch == 0x20) {
		return true
	}
	return false
}
func _ddgff(_accef PdfObject) (*float64, error) {
	switch _bada := _accef.(type) {
	case *PdfObjectFloat:
		_bdfd := float64(*_bada)
		return &_bdfd, nil
	case *PdfObjectInteger:
		_gacbd := float64(*_bada)
		return &_gacbd, nil
	case *PdfObjectNull:
		return nil, nil
	}
	return nil, ErrNotANumber
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
	_feg   *_de.Document

	// Globals are the JBIG2 global segments.
	Globals _fb.Globals

	// IsChocolateData defines if the data is encoded such that
	// binary data '1' means black and '0' white.
	// otherwise the data is called vanilla.
	// Naming convention taken from: 'https://en.wikipedia.org/wiki/Binary_image#Interpretation'
	IsChocolateData bool

	// DefaultPageSettings are the settings parameters used by the jbig2 encoder.
	DefaultPageSettings JBIG2EncoderSettings
}

func (_aaeab *PdfParser) repairRebuildXrefsTopDown() (*XrefTable, error) {
	if _aaeab._cged {
		return nil, _gf.Errorf("\u0072\u0065\u0070\u0061\u0069\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_aaeab._cged = true
	_aaeab._dfcdg.Seek(0, _fg.SeekStart)
	_aaeab._ffbg = _bfc.NewReader(_aaeab._dfcdg)
	_dbca := 20
	_bceca := make([]byte, _dbca)
	_aeag := XrefTable{}
	_aeag.ObjectMap = make(map[int]XrefObject)
	for {
		_bfac, _acdae := _aaeab._ffbg.ReadByte()
		if _acdae != nil {
			if _acdae == _fg.EOF {
				break
			} else {
				return nil, _acdae
			}
		}
		if _bfac == 'j' && _bceca[_dbca-1] == 'b' && _bceca[_dbca-2] == 'o' && IsWhiteSpace(_bceca[_dbca-3]) {
			_egcg := _dbca - 4
			for IsWhiteSpace(_bceca[_egcg]) && _egcg > 0 {
				_egcg--
			}
			if _egcg == 0 || !IsDecimalDigit(_bceca[_egcg]) {
				continue
			}
			for IsDecimalDigit(_bceca[_egcg]) && _egcg > 0 {
				_egcg--
			}
			if _egcg == 0 || !IsWhiteSpace(_bceca[_egcg]) {
				continue
			}
			for IsWhiteSpace(_bceca[_egcg]) && _egcg > 0 {
				_egcg--
			}
			if _egcg == 0 || !IsDecimalDigit(_bceca[_egcg]) {
				continue
			}
			for IsDecimalDigit(_bceca[_egcg]) && _egcg > 0 {
				_egcg--
			}
			if _egcg == 0 {
				continue
			}
			_gfbdg := _aaeab.GetFileOffset() - int64(_dbca-_egcg)
			_gfgf := append(_bceca[_egcg+1:], _bfac)
			_gadbg, _cbgc, _dcgc := _gdce(string(_gfgf))
			if _dcgc != nil {
				_a.Log.Debug("\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u003a\u0020\u0025\u0076", _dcgc)
				return nil, _dcgc
			}
			if _aacb, _fgfb := _aeag.ObjectMap[_gadbg]; !_fgfb || _aacb.Generation < _cbgc {
				_ebdg := XrefObject{}
				_ebdg.XType = XrefTypeTableEntry
				_ebdg.ObjectNumber = _gadbg
				_ebdg.Generation = _cbgc
				_ebdg.Offset = _gfbdg
				_aeag.ObjectMap[_gadbg] = _ebdg
			}
		}
		_bceca = append(_bceca[1:_dbca], _bfac)
	}
	_aaeab._eeaa = nil
	return &_aeag, nil
}

// NewDCTEncoder makes a new DCT encoder with default parameters.
func NewDCTEncoder() *DCTEncoder {
	_bbd := &DCTEncoder{}
	_bbd.ColorComponents = 3
	_bbd.BitsPerComponent = 8
	_bbd.Quality = DefaultJPEGQuality
	_bbd.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	return _bbd
}

// WriteString outputs the object as it is to be written to file.
func (_fega *PdfObjectInteger) WriteString() string { return _bd.FormatInt(int64(*_fega), 10) }

// IsDelimiter checks if a character represents a delimiter.
func IsDelimiter(c byte) bool {
	return c == '(' || c == ')' || c == '<' || c == '>' || c == '[' || c == ']' || c == '{' || c == '}' || c == '/' || c == '%'
}
func (_acdb *PdfCrypt) decryptBytes(_dfc []byte, _gba string, _gbea []byte) ([]byte, error) {
	_a.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_ecec, _ecc := _acdb._deg[_gba]
	if !_ecc {
		return nil, _gf.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _gba)
	}
	return _ecec.DecryptBytes(_dfc, _gbea)
}

// WriteString outputs the object as it is to be written to file.
func (_fdagb *PdfObjectArray) WriteString() string {
	var _eceee _gd.Builder
	_eceee.WriteString("\u005b")
	for _bcgbc, _eafac := range _fdagb.Elements() {
		_eceee.WriteString(_eafac.WriteString())
		if _bcgbc < (_fdagb.Len() - 1) {
			_eceee.WriteString("\u0020")
		}
	}
	_eceee.WriteString("\u005d")
	return _eceee.String()
}
func (_fed *PdfCrypt) loadCryptFilters(_gaeb *PdfObjectDictionary) error {
	_fed._deg = cryptFilters{}
	_bde := _gaeb.Get("\u0043\u0046")
	_bde = TraceToDirectObject(_bde)
	if _dbf, _dcd := _bde.(*PdfObjectReference); _dcd {
		_adg, _ffd := _fed._ada.LookupByReference(*_dbf)
		if _ffd != nil {
			_a.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0075\u0070\u0020\u0043\u0046\u0020\u0072\u0065\u0066\u0065\u0072en\u0063\u0065")
			return _ffd
		}
		_bde = TraceToDirectObject(_adg)
	}
	_dgc, _gdc := _bde.(*PdfObjectDictionary)
	if !_gdc {
		_a.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0043\u0046\u002c \u0074\u0079\u0070\u0065: \u0025\u0054", _bde)
		return _c.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0046")
	}
	for _, _baec := range _dgc.Keys() {
		_cag := _dgc.Get(_baec)
		if _cbf, _edb := _cag.(*PdfObjectReference); _edb {
			_adb, _eef := _fed._ada.LookupByReference(*_cbf)
			if _eef != nil {
				_a.Log.Debug("\u0045\u0072ro\u0072\u0020\u006co\u006f\u006b\u0075\u0070 up\u0020di\u0063\u0074\u0069\u006f\u006e\u0061\u0072y \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065")
				return _eef
			}
			_cag = TraceToDirectObject(_adb)
		}
		_gbb, _ffc := _cag.(*PdfObjectDictionary)
		if !_ffc {
			return _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006e \u0043\u0046\u0020\u0028\u006e\u0061\u006d\u0065\u0020\u0025\u0073\u0029\u0020-\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0062\u0075\u0074\u0020\u0025\u0054", _baec, _cag)
		}
		if _baec == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u002d\u0020\u0043\u0061\u006e\u006e\u006f\u0074\u0020\u006f\u0076\u0065\u0072\u0077r\u0069\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0069d\u0065\u006e\u0074\u0069\u0074\u0079\u0020\u0066\u0069\u006c\u0074\u0065\u0072 \u002d\u0020\u0054\u0072\u0079\u0069n\u0067\u0020\u006ee\u0078\u0074")
			continue
		}
		var _fcb _efd.FilterDict
		if _adbc := _db(&_fcb, _gbb); _adbc != nil {
			return _adbc
		}
		_gdf, _ggc := _efd.NewFilter(_fcb)
		if _ggc != nil {
			return _ggc
		}
		_fed._deg[string(_baec)] = _gdf
	}
	_fed._deg["\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"] = _efd.NewIdentity()
	_fed._dfb = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _gea, _ddd := _gaeb.Get("\u0053\u0074\u0072\u0046").(*PdfObjectName); _ddd {
		if _, _cec := _fed._deg[string(*_gea)]; !_cec {
			return _gf.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u0072\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_gea)
		}
		_fed._dfb = string(*_gea)
	}
	_fed._abdf = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
	if _aeb, _bea := _gaeb.Get("\u0053\u0074\u006d\u0046").(*PdfObjectName); _bea {
		if _, _dfae := _fed._deg[string(*_aeb)]; !_dfae {
			return _gf.Errorf("\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0066o\u0072\u0020\u0053\u0074\u006d\u0046\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069e\u0064\u0020\u0069\u006e\u0020C\u0046\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0028\u0025\u0073\u0029", *_aeb)
		}
		_fed._abdf = string(*_aeb)
	}
	return nil
}

// ParseDict reads and parses a PDF dictionary object enclosed with '<<' and '>>'
func (_cdfbb *PdfParser) ParseDict() (*PdfObjectDictionary, error) {
	_a.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_afdf := MakeDict()
	_afdf._gdcf = _cdfbb
	_ggaf, _ := _cdfbb._ffbg.ReadByte()
	if _ggaf != '<' {
		return nil, _c.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_ggaf, _ = _cdfbb._ffbg.ReadByte()
	if _ggaf != '<' {
		return nil, _c.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_cdfbb.skipSpaces()
		_cdfbb.skipComments()
		_gfaa, _eagcf := _cdfbb._ffbg.Peek(2)
		if _eagcf != nil {
			return nil, _eagcf
		}
		_a.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_gfaa), string(_gfaa))
		if (_gfaa[0] == '>') && (_gfaa[1] == '>') {
			_a.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_cdfbb._ffbg.ReadByte()
			_cdfbb._ffbg.ReadByte()
			break
		}
		_a.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_cdgg, _eagcf := _cdfbb.parseName()
		_a.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _cdgg)
		if _eagcf != nil {
			_a.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _eagcf)
			return nil, _eagcf
		}
		if len(_cdgg) > 4 && _cdgg[len(_cdgg)-4:] == "\u006e\u0075\u006c\u006c" {
			_aaaa := _cdgg[0 : len(_cdgg)-4]
			_a.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _cdgg)
			_a.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _aaaa)
			_cdfbb.skipSpaces()
			_ebfg, _ := _cdfbb._ffbg.Peek(1)
			if _ebfg[0] == '/' {
				_afdf.Set(_aaaa, MakeNull())
				continue
			}
		}
		_cdfbb.skipSpaces()
		_cbg, _eagcf := _cdfbb.parseObject()
		if _eagcf != nil {
			return nil, _eagcf
		}
		_afdf.Set(_cdgg, _cbg)
		if _a.Log.IsLogLevel(_a.LogLevelTrace) {
			_a.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _cdgg, _cbg.String())
		}
	}
	_a.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _afdf, nil
}

// RawEncoder implements Raw encoder/decoder (no encoding, pass through)
type RawEncoder struct{}

func _agaf(_adfed PdfObject, _cdfeed int, _afag map[PdfObject]struct{}) error {
	_a.Log.Trace("\u0054\u0072\u0061\u0076\u0065\u0072s\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0061 \u0028\u0064\u0065\u0070\u0074\u0068\u0020=\u0020\u0025\u0064\u0029", _cdfeed)
	if _, _ecfb := _afag[_adfed]; _ecfb {
		_a.Log.Trace("-\u0041\u006c\u0072\u0065ad\u0079 \u0074\u0072\u0061\u0076\u0065r\u0073\u0065\u0064\u002e\u002e\u002e")
		return nil
	}
	_afag[_adfed] = struct{}{}
	switch _edagf := _adfed.(type) {
	case *PdfIndirectObject:
		_agfd := _edagf
		_a.Log.Trace("\u0069\u006f\u003a\u0020\u0025\u0073", _agfd)
		_a.Log.Trace("\u002d\u0020\u0025\u0073", _agfd.PdfObject)
		return _agaf(_agfd.PdfObject, _cdfeed+1, _afag)
	case *PdfObjectStream:
		_eeba := _edagf
		return _agaf(_eeba.PdfObjectDictionary, _cdfeed+1, _afag)
	case *PdfObjectDictionary:
		_acbfd := _edagf
		_a.Log.Trace("\u002d\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073", _acbfd)
		for _, _bafad := range _acbfd.Keys() {
			_ebfd := _acbfd.Get(_bafad)
			if _cfgbf, _faadb := _ebfd.(*PdfObjectReference); _faadb {
				_baffa := _cfgbf.Resolve()
				_acbfd.Set(_bafad, _baffa)
				_ageba := _agaf(_baffa, _cdfeed+1, _afag)
				if _ageba != nil {
					return _ageba
				}
			} else {
				_ffcge := _agaf(_ebfd, _cdfeed+1, _afag)
				if _ffcge != nil {
					return _ffcge
				}
			}
		}
		return nil
	case *PdfObjectArray:
		_gbebe := _edagf
		_a.Log.Trace("-\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0073", _gbebe)
		for _caef, _dabdf := range _gbebe.Elements() {
			if _bebb, _fgge := _dabdf.(*PdfObjectReference); _fgge {
				_dedg := _bebb.Resolve()
				_gbebe.Set(_caef, _dedg)
				_dbbea := _agaf(_dedg, _cdfeed+1, _afag)
				if _dbbea != nil {
					return _dbbea
				}
			} else {
				_gecb := _agaf(_dabdf, _cdfeed+1, _afag)
				if _gecb != nil {
					return _gecb
				}
			}
		}
		return nil
	case *PdfObjectReference:
		_a.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020T\u0072\u0061\u0063\u0069\u006e\u0067\u0020\u0061\u0020r\u0065\u0066\u0065r\u0065n\u0063\u0065\u0021")
		return _c.New("\u0065r\u0072\u006f\u0072\u0020t\u0072\u0061\u0063\u0069\u006eg\u0020a\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	return nil
}

// SetFileOffset sets the file to an offset position and resets buffer.
func (_afcf *PdfParser) SetFileOffset(offset int64) {
	if offset < 0 {
		offset = 0
	}
	_afcf._dfcdg.Seek(offset, _fg.SeekStart)
	_afcf._ffbg = _bfc.NewReader(_afcf._dfcdg)
}

// GetParser returns the parser for lazy-loading or compare references.
func (_bfaf *PdfObjectReference) GetParser() *PdfParser { return _bfaf._cfada }
func (_efae *PdfParser) checkPostEOFData() error {
	const _agfe = "\u0025\u0025\u0045O\u0046"
	_, _cgad := _efae._dfcdg.Seek(-int64(len([]byte(_agfe)))-1, _fg.SeekEnd)
	if _cgad != nil {
		return _cgad
	}
	_cddaa := make([]byte, len([]byte(_agfe))+1)
	_, _cgad = _efae._dfcdg.Read(_cddaa)
	if _cgad != nil {
		if _cgad != _fg.EOF {
			return _cgad
		}
	}
	if string(_cddaa) == _agfe || string(_cddaa) == _agfe+"\u000a" {
		_efae._gadge._ccbc = true
	}
	return nil
}

// HasInvalidSeparationAfterXRef implements core.ParserMetadata interface.
func (_eadg ParserMetadata) HasInvalidSeparationAfterXRef() bool { return _eadg._bdf }

// String returns the state of the bool as "true" or "false".
func (_faad *PdfObjectBool) String() string {
	if *_faad {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_efgcd *JPXEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }
func (_cgf *PdfCrypt) saveCryptFilters(_fbb *PdfObjectDictionary) error {
	if _cgf._ceae.V < 4 {
		return _c.New("\u0063\u0061\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020V\u003e\u003d\u0034")
	}
	_gaee := MakeDict()
	_fbb.Set("\u0043\u0046", _gaee)
	for _eda, _feb := range _cgf._deg {
		if _eda == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
			continue
		}
		_cfb := _gfdf(_feb, "")
		_gaee.Set(PdfObjectName(_eda), _cfb)
	}
	_fbb.Set("\u0053\u0074\u0072\u0046", MakeName(_cgf._dfb))
	_fbb.Set("\u0053\u0074\u006d\u0046", MakeName(_cgf._abdf))
	return nil
}

// DecodeBytes decodes a byte slice from Run length encoding.
//
// 7.4.5 RunLengthDecode Filter
// The RunLengthDecode filter decodes data that has been encoded in a simple byte-oriented format based on run length.
// The encoded data shall be a sequence of runs, where each run shall consist of a length byte followed by 1 to 128
// bytes of data. If the length byte is in the range 0 to 127, the following length + 1 (1 to 128) bytes shall be
// copied literally during decompression. If length is in the range 129 to 255, the following single byte shall be
// copied 257 - length (2 to 128) times during decompression. A length value of 128 shall denote EOD.
func (_edgc *RunLengthEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_gdfbg := _fd.NewReader(encoded)
	var _dacf []byte
	for {
		_bfgf, _ffg := _gdfbg.ReadByte()
		if _ffg != nil {
			return nil, _ffg
		}
		if _bfgf > 128 {
			_fgec, _bcg := _gdfbg.ReadByte()
			if _bcg != nil {
				return nil, _bcg
			}
			for _gee := 0; _gee < 257-int(_bfgf); _gee++ {
				_dacf = append(_dacf, _fgec)
			}
		} else if _bfgf < 128 {
			for _afgg := 0; _afgg < int(_bfgf)+1; _afgg++ {
				_faag, _gegc := _gdfbg.ReadByte()
				if _gegc != nil {
					return nil, _gegc
				}
				_dacf = append(_dacf, _faag)
			}
		} else {
			break
		}
	}
	return _dacf, nil
}

// MakeFloat creates an PdfObjectFloat from a float64.
func MakeFloat(val float64) *PdfObjectFloat { _fgcea := PdfObjectFloat(val); return &_fgcea }

// MakeBool creates a PdfObjectBool from a bool value.
func MakeBool(val bool) *PdfObjectBool { _gced := PdfObjectBool(val); return &_gced }

// ASCII85Encoder implements ASCII85 encoder/decoder.
type ASCII85Encoder struct{}

// PdfVersion returns version of the PDF file.
func (_bdgbe *PdfParser) PdfVersion() Version { return _bdgbe._eebec }

// NewCompliancePdfParser creates a new PdfParser that will parse input reader with the focus on extracting more metadata, which
// might affect performance of the regular PdfParser this function.
func NewCompliancePdfParser(rs _fg.ReadSeeker) (_afc *PdfParser, _ffb error) {
	_afc = &PdfParser{_dfcdg: rs, ObjCache: make(objectCache), _ffed: map[int64]bool{}, _ecgd: true, _beaf: make(map[*PdfParser]*PdfParser)}
	if _ffb = _afc.parseDetailedHeader(); _ffb != nil {
		return nil, _ffb
	}
	if _afc._dabde, _ffb = _afc.loadXrefs(); _ffb != nil {
		_a.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _ffb)
		return nil, _ffb
	}
	_a.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _afc._dabde)
	if len(_afc._bbdf.ObjectMap) == 0 {
		return nil, _gf.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	return _afc, nil
}

// DrawableImage is same as golang image/draw's Image interface that allow drawing images.
type DrawableImage interface {
	ColorModel() _ga.Model
	Bounds() _cg.Rectangle
	At(_dgag, _gedf int) _ga.Color
	Set(_bgbc, _bace int, _gggf _ga.Color)
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_ccf *RunLengthEncoder) MakeDecodeParams() PdfObject { return nil }

// UpdateParams updates the parameter values of the encoder.
// Implements StreamEncoder interface.
func (_bdbd *JBIG2Encoder) UpdateParams(params *PdfObjectDictionary) {
	_gdfbc, _fde := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _fde == nil {
		_bdbd.BitsPerComponent = int(_gdfbc)
	}
	_daca, _fde := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _fde == nil {
		_bdbd.Width = int(_daca)
	}
	_gfea, _fde := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _fde == nil {
		_bdbd.Height = int(_gfea)
	}
	_becg, _fde := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _fde == nil {
		_bdbd.ColorComponents = int(_becg)
	}
}

// Decoded returns the PDFDocEncoding or UTF-16BE decoded string contents.
// UTF-16BE is applied when the first two bytes are 0xFE, 0XFF, otherwise decoding of
// PDFDocEncoding is performed.
func (_bgfca *PdfObjectString) Decoded() string {
	if _bgfca == nil {
		return ""
	}
	_dfebe := []byte(_bgfca._aaca)
	if len(_dfebe) >= 2 && _dfebe[0] == 0xFE && _dfebe[1] == 0xFF {
		return _da.UTF16ToString(_dfebe[2:])
	}
	return _da.PDFDocEncodingToString(_dfebe)
}
func (_egbb *PdfObjectDictionary) setWithLock(_edbad PdfObjectName, _cfac PdfObject, _befb bool) {
	if _befb {
		_egbb._ccdg.Lock()
		defer _egbb._ccdg.Unlock()
	}
	_, _bfbef := _egbb._adbeb[_edbad]
	if !_bfbef {
		_egbb._caeg = append(_egbb._caeg, _edbad)
	}
	_egbb._adbeb[_edbad] = _cfac
}

// UpdateParams updates the parameter values of the encoder.
func (_fdf *FlateEncoder) UpdateParams(params *PdfObjectDictionary) {
	_cded, _dead := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _dead == nil {
		_fdf.Predictor = int(_cded)
	}
	_eaf, _dead := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _dead == nil {
		_fdf.BitsPerComponent = int(_eaf)
	}
	_babd, _dead := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _dead == nil {
		_fdf.Columns = int(_babd)
	}
	_ggad, _dead := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _dead == nil {
		_fdf.Colors = int(_ggad)
	}
}
func (_feged *PdfParser) inspect() (map[string]int, error) {
	_a.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0053P\u0045\u0043\u0054\u0020\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d")
	_a.Log.Trace("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_cadbc := map[string]int{}
	_efcdg := 0
	_cbae := 0
	var _baefb []int
	for _edcg := range _feged._bbdf.ObjectMap {
		_baefb = append(_baefb, _edcg)
	}
	_gg.Ints(_baefb)
	_gdea := 0
	for _, _gfca := range _baefb {
		_bffd := _feged._bbdf.ObjectMap[_gfca]
		if _bffd.ObjectNumber == 0 {
			continue
		}
		_efcdg++
		_a.Log.Trace("\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d")
		_a.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0075\u0070\u0020\u006f\u0062\u006a\u0065c\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0025\u0064", _bffd.ObjectNumber)
		_ccfbe, _cfda := _feged.LookupByNumber(_bffd.ObjectNumber)
		if _cfda != nil {
			_a.Log.Trace("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u006c\u006f\u006f\u006b\u0075p\u0020\u006f\u0062\u006a\u0020\u0025\u0064 \u0028\u0025\u0073\u0029", _bffd.ObjectNumber, _cfda)
			_cbae++
			continue
		}
		_a.Log.Trace("\u006fb\u006a\u003a\u0020\u0025\u0073", _ccfbe)
		_fggg, _ccgbbc := _ccfbe.(*PdfIndirectObject)
		if _ccgbbc {
			_a.Log.Trace("\u0049N\u0044 \u004f\u004f\u0042\u004a\u0020\u0025\u0064\u003a\u0020\u0025\u0073", _bffd.ObjectNumber, _fggg)
			_cgae, _ebeed := _fggg.PdfObject.(*PdfObjectDictionary)
			if _ebeed {
				if _bfag, _daad := _cgae.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _daad {
					_feegc := string(*_bfag)
					_a.Log.Trace("\u002d\u002d\u002d\u003e\u0020\u004f\u0062\u006a\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073", _feegc)
					_, _gdabf := _cadbc[_feegc]
					if _gdabf {
						_cadbc[_feegc]++
					} else {
						_cadbc[_feegc] = 1
					}
				} else if _fgfe, _ddfc := _cgae.Get("\u0053u\u0062\u0074\u0079\u0070\u0065").(*PdfObjectName); _ddfc {
					_cgce := string(*_fgfe)
					_a.Log.Trace("-\u002d-\u003e\u0020\u004f\u0062\u006a\u0020\u0073\u0075b\u0074\u0079\u0070\u0065: \u0025\u0073", _cgce)
					_, _ggeg := _cadbc[_cgce]
					if _ggeg {
						_cadbc[_cgce]++
					} else {
						_cadbc[_cgce] = 1
					}
				}
				if _edfgb, _gcef := _cgae.Get("\u0053").(*PdfObjectName); _gcef && *_edfgb == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
					_, _cfcg := _cadbc["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]
					if _cfcg {
						_cadbc["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
					} else {
						_cadbc["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"] = 1
					}
				}
			}
		} else if _gefa, _aaddg := _ccfbe.(*PdfObjectStream); _aaddg {
			if _ecfc, _cdbg := _gefa.PdfObjectDictionary.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _cdbg {
				_a.Log.Trace("\u002d\u002d\u003e\u0020\u0053\u0074\u0072\u0065\u0061\u006d\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073", *_ecfc)
				_cafcb := string(*_ecfc)
				_cadbc[_cafcb]++
			}
		} else {
			_ecfg, _dceeg := _ccfbe.(*PdfObjectDictionary)
			if _dceeg {
				_efdc, _ggaee := _ecfg.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName)
				if _ggaee {
					_bdad := string(*_efdc)
					_a.Log.Trace("\u002d-\u002d \u006f\u0062\u006a\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _bdad)
					_cadbc[_bdad]++
				}
			}
			_a.Log.Trace("\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u004f\u0042\u004a\u0020\u0025d\u003a\u0020\u0025\u0073", _bffd.ObjectNumber, _ccfbe)
		}
		_gdea++
	}
	_a.Log.Trace("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0045\u004fF\u0020\u0049\u004e\u0053\u0050\u0045\u0043T\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d")
	_a.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	_a.Log.Trace("\u004f\u0062j\u0065\u0063\u0074 \u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064", _efcdg)
	_a.Log.Trace("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u006f\u006b\u0075p\u003a\u0020\u0025\u0064", _cbae)
	for _feaca, _bdfc := range _cadbc {
		_a.Log.Trace("\u0025\u0073\u003a\u0020\u0025\u0064", _feaca, _bdfc)
	}
	_a.Log.Trace("\u003d=\u003d\u003d\u003d\u003d\u003d")
	if len(_feged._bbdf.ObjectMap) < 1 {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0068\u0069\u0073 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074 \u0069s\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0028\u0078\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0021\u0029")
		return nil, _gf.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0028\u0078r\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006d\u0069\u0073s\u0069\u006e\u0067\u0029")
	}
	_cfcbe, _agcg := _cadbc["\u0046\u006f\u006e\u0074"]
	if !_agcg || _cfcbe < 2 {
		_a.Log.Trace("\u0054\u0068\u0069s \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020i\u0073 \u0070r\u006fb\u0061\u0062\u006c\u0079\u0020\u0073\u0063\u0061\u006e\u006e\u0065\u0064\u0021")
	} else {
		_a.Log.Trace("\u0054\u0068\u0069\u0073\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0066o\u0072\u0020\u0065\u0078\u0074r\u0061\u0063t\u0069\u006f\u006e\u0021")
	}
	return _cadbc, nil
}

// Remove removes an element specified by key.
func (_fcfc *PdfObjectDictionary) Remove(key PdfObjectName) {
	_becb := -1
	for _cddac, _abeag := range _fcfc._caeg {
		if _abeag == key {
			_becb = _cddac
			break
		}
	}
	if _becb >= 0 {
		_fcfc._caeg = append(_fcfc._caeg[:_becb], _fcfc._caeg[_becb+1:]...)
		delete(_fcfc._adbeb, key)
	}
}

// PdfCrypt provides PDF encryption/decryption support.
// The PDF standard supports encryption of strings and streams (Section 7.6).
type PdfCrypt struct {
	_ceae encryptDict
	_aec  _dc.StdEncryptDict
	_abf  string
	_baa  []byte
	_ecee map[PdfObject]bool
	_fee  map[PdfObject]bool
	_dfe  bool
	_deg  cryptFilters
	_abdf string
	_dfb  string
	_ada  *PdfParser
	_cdda map[int]struct{}
}

// PdfObjectString represents the primitive PDF string object.
type PdfObjectString struct {
	_aaca string
	_gead bool
}
type encryptDict struct {
	Filter    string
	V         int
	SubFilter string
	Length    int
	StmF      string
	StrF      string
	EFF       string
	CF        map[string]_efd.FilterDict
}

// WriteString outputs the object as it is to be written to file.
func (_eacdd *PdfIndirectObject) WriteString() string {
	var _ggga _gd.Builder
	_ggga.WriteString(_bd.FormatInt(_eacdd.ObjectNumber, 10))
	_ggga.WriteString("\u0020\u0030\u0020\u0052")
	return _ggga.String()
}
func _db(_fbg *_efd.FilterDict, _bce *PdfObjectDictionary) error {
	if _cefc, _acd := _bce.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _acd {
		if _cbe := string(*_cefc); _cbe != "C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072" {
			_a.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020C\u0046\u0020\u0064ic\u0074\u0020\u0074\u0079\u0070\u0065:\u0020\u0025\u0073\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074e\u0072\u0029", _cbe)
		}
	}
	_dag, _dagg := _bce.Get("\u0043\u0046\u004d").(*PdfObjectName)
	if !_dagg {
		return _gf.Errorf("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0063\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0028\u004e\u006f\u006e\u0065\u0029")
	}
	_fbg.CFM = string(*_dag)
	if _bcd, _cbcg := _bce.Get("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et").(*PdfObjectName); _cbcg {
		_fbg.AuthEvent = _dc.AuthEvent(*_bcd)
	} else {
		_fbg.AuthEvent = _dc.EventDocOpen
	}
	if _fdb, _gacd := _bce.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _gacd {
		_fbg.Length = int(*_fdb)
	}
	return nil
}

var _dbdb = _ce.MustCompile("\u0025\u0025\u0045\u004f\u0046\u003f")

// String returns the PDF version as a string. Implements interface fmt.Stringer.
func (_caeb Version) String() string {
	return _gf.Sprintf("\u00250\u0064\u002e\u0025\u0030\u0064", _caeb.Major, _caeb.Minor)
}

// DecodeBytes decodes a slice of ASCII encoded bytes and returns the result.
func (_dbae *ASCIIHexEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_dfgc := _fd.NewReader(encoded)
	var _ffgf []byte
	for {
		_bdcad, _geb := _dfgc.ReadByte()
		if _geb != nil {
			return nil, _geb
		}
		if _bdcad == '>' {
			break
		}
		if IsWhiteSpace(_bdcad) {
			continue
		}
		if (_bdcad >= 'a' && _bdcad <= 'f') || (_bdcad >= 'A' && _bdcad <= 'F') || (_bdcad >= '0' && _bdcad <= '9') {
			_ffgf = append(_ffgf, _bdcad)
		} else {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0061\u0073\u0063\u0069\u0069 \u0068\u0065\u0078\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072 \u0028\u0025\u0063\u0029", _bdcad)
			return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0073\u0063\u0069\u0069\u0020\u0068e\u0078 \u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0028\u0025\u0063\u0029", _bdcad)
		}
	}
	if len(_ffgf)%2 == 1 {
		_ffgf = append(_ffgf, '0')
	}
	_a.Log.Trace("\u0049\u006e\u0062\u006f\u0075\u006e\u0064\u0020\u0025\u0073", _ffgf)
	_bdef := make([]byte, _bdc.DecodedLen(len(_ffgf)))
	_, _aefe := _bdc.Decode(_bdef, _ffgf)
	if _aefe != nil {
		return nil, _aefe
	}
	return _bdef, nil
}
func (_daga *PdfCrypt) newEncryptDict() *PdfObjectDictionary {
	_eee := MakeDict()
	_eee.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064"))
	_eee.Set("\u0056", MakeInteger(int64(_daga._ceae.V)))
	_eee.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_daga._ceae.Length)))
	return _eee
}

// NewRunLengthEncoder makes a new run length encoder
func NewRunLengthEncoder() *RunLengthEncoder { return &RunLengthEncoder{} }
func (_aafg *PdfParser) getNumbersOfUpdatedObjects(_cead *PdfParser) ([]int, error) {
	if _cead == nil {
		return nil, _c.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_cfbb := _cead._fbaaa
	_gbdea := make([]int, 0)
	_bbee := make(map[int]interface{})
	_fdef := make(map[int]int64)
	for _bacf, _fdaa := range _aafg._bbdf.ObjectMap {
		if _fdaa.Offset == 0 {
			if _fdaa.OsObjNumber != 0 {
				if _eegfg, _fddce := _aafg._bbdf.ObjectMap[_fdaa.OsObjNumber]; _fddce {
					_bbee[_fdaa.OsObjNumber] = struct{}{}
					_fdef[_bacf] = _eegfg.Offset
				} else {
					return nil, _c.New("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0078r\u0065\u0066\u0020\u0074ab\u006c\u0065")
				}
			}
		} else {
			_fdef[_bacf] = _fdaa.Offset
		}
	}
	for _dagae, _gaaf := range _fdef {
		if _, _cege := _bbee[_dagae]; _cege {
			continue
		}
		if _gaaf > _cfbb {
			_gbdea = append(_gbdea, _dagae)
		}
	}
	return _gbdea, nil
}

// String returns a string describing `d`.
func (_dcgaa *PdfObjectDictionary) String() string {
	var _ceed _gd.Builder
	_ceed.WriteString("\u0044\u0069\u0063t\u0028")
	for _, _ccbdc := range _dcgaa._caeg {
		_gbdf := _dcgaa._adbeb[_ccbdc]
		_ceed.WriteString("\u0022" + _ccbdc.String() + "\u0022\u003a\u0020")
		_ceed.WriteString(_gbdf.String())
		_ceed.WriteString("\u002c\u0020")
	}
	_ceed.WriteString("\u0029")
	return _ceed.String()
}

// UpdateParams updates the parameter values of the encoder.
func (_cffg *ASCIIHexEncoder) UpdateParams(params *PdfObjectDictionary) {}

// GetFilterName returns the name of the encoding filter.
func (_gaab *JBIG2Encoder) GetFilterName() string { return StreamEncodingFilterNameJBIG2 }

// EncodeImage encodes 'img' golang image.Image into jbig2 encoded bytes document using default encoder settings.
func (_beb *JBIG2Encoder) EncodeImage(img _cg.Image) ([]byte, error) { return _beb.encodeImage(img) }
func (_eegd *PdfParser) parseNull() (PdfObjectNull, error) {
	_, _eegf := _eegd._ffbg.Discard(4)
	return PdfObjectNull{}, _eegf
}

// MakeArrayFromIntegers64 creates an PdfObjectArray from a slice of int64s, where each array element
// is an PdfObjectInteger.
func MakeArrayFromIntegers64(vals []int64) *PdfObjectArray {
	_dbbd := MakeArray()
	for _, _cgdgc := range vals {
		_dbbd.Append(MakeInteger(_cgdgc))
	}
	return _dbbd
}

// Bytes returns the PdfObjectString content as a []byte array.
func (_acfcc *PdfObjectString) Bytes() []byte { return []byte(_acfcc._aaca) }

// NewLZWEncoder makes a new LZW encoder with default parameters.
func NewLZWEncoder() *LZWEncoder {
	_bdea := &LZWEncoder{}
	_bdea.Predictor = 1
	_bdea.BitsPerComponent = 8
	_bdea.Colors = 1
	_bdea.Columns = 1
	_bdea.EarlyChange = 1
	return _bdea
}

// EncodeBytes implements support for LZW encoding.  Currently not supporting predictors (raw compressed data only).
// Only supports the Early change = 1 algorithm (compress/lzw) as the other implementation
// does not have a write method.
// TODO: Consider refactoring compress/lzw to allow both.
func (_dafa *LZWEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _dafa.Predictor != 1 {
		return nil, _gf.Errorf("\u004c\u005aW \u0050\u0072\u0065d\u0069\u0063\u0074\u006fr =\u00201 \u006f\u006e\u006c\u0079\u0020\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0079e\u0074")
	}
	if _dafa.EarlyChange == 1 {
		return nil, _gf.Errorf("\u004c\u005a\u0057\u0020\u0045\u0061\u0072\u006c\u0079\u0020\u0043\u0068\u0061n\u0067\u0065\u0020\u003d\u0020\u0030 \u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0079\u0065\u0074")
	}
	var _fefg _fd.Buffer
	_ebad := _bf.NewWriter(&_fefg, _bf.MSB, 8)
	_ebad.Write(data)
	_ebad.Close()
	return _fefg.Bytes(), nil
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

func _bead(_eaca int) int {
	if _eaca < 0 {
		return -_eaca
	}
	return _eaca
}

// IsAuthenticated returns true if the PDF has already been authenticated for accessing.
func (_abgg *PdfParser) IsAuthenticated() bool { return _abgg._abae._dfe }

// GetXrefType returns the type of the first xref object (table or stream).
func (_cdfe *PdfParser) GetXrefType() *xrefType { return _cdfe._fddf }
func (_dca *PdfParser) rebuildXrefTable() error {
	_bcdfg := XrefTable{}
	_bcdfg.ObjectMap = map[int]XrefObject{}
	_gfdb := make([]int, 0, len(_dca._bbdf.ObjectMap))
	for _deec := range _dca._bbdf.ObjectMap {
		_gfdb = append(_gfdb, _deec)
	}
	_gg.Ints(_gfdb)
	for _, _caeda := range _gfdb {
		_fdcfd := _dca._bbdf.ObjectMap[_caeda]
		_dffc, _, _dbbbd := _dca.lookupByNumberWrapper(_caeda, false)
		if _dbbbd != nil {
			_a.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f l\u006f\u006f\u006b\u0020\u0075\u0070\u0020ob\u006a\u0065\u0063\u0074\u0020\u0028\u0025s\u0029", _dbbbd)
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u006fm\u0070\u006c\u0065\u0074\u0065\u006c\u0079\u0020\u0062\u0072\u006f\u006b\u0065\u006e\u0020\u002d\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0074\u006f \u0072\u0065\u0070\u0061\u0069r\u0020")
			_bcfgc, _deefc := _dca.repairRebuildXrefsTopDown()
			if _deefc != nil {
				_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0072\u0065\u0062\u0075\u0069l\u0064\u0020\u0072\u0065\u0070a\u0069\u0072 \u0028\u0025\u0073\u0029", _deefc)
				return _deefc
			}
			_dca._bbdf = *_bcfgc
			_a.Log.Debug("\u0052e\u0070\u0061\u0069\u0072e\u0064\u0020\u0078\u0072\u0065f\u0020t\u0061b\u006c\u0065\u0020\u0062\u0075\u0069\u006ct")
			return nil
		}
		_bbdc, _afddc, _dbbbd := _ecg(_dffc)
		if _dbbbd != nil {
			return _dbbbd
		}
		_fdcfd.ObjectNumber = int(_bbdc)
		_fdcfd.Generation = int(_afddc)
		_bcdfg.ObjectMap[int(_bbdc)] = _fdcfd
	}
	_dca._bbdf = _bcdfg
	_a.Log.Debug("N\u0065w\u0020\u0078\u0072\u0065\u0066\u0020\u0074\u0061b\u006c\u0065\u0020\u0062ui\u006c\u0074")
	_ff(_dca._bbdf)
	return nil
}
func _aceeg(_aeafd, _aagc PdfObject, _ccfg int) bool {
	if _ccfg > _ddfg {
		_a.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _ddfg)
		return false
	}
	if _aeafd == nil && _aagc == nil {
		return true
	} else if _aeafd == nil || _aagc == nil {
		return false
	}
	if _bc.TypeOf(_aeafd) != _bc.TypeOf(_aagc) {
		return false
	}
	switch _ggegf := _aeafd.(type) {
	case *PdfObjectNull, *PdfObjectReference:
		return true
	case *PdfObjectName:
		return *_ggegf == *(_aagc.(*PdfObjectName))
	case *PdfObjectString:
		return *_ggegf == *(_aagc.(*PdfObjectString))
	case *PdfObjectInteger:
		return *_ggegf == *(_aagc.(*PdfObjectInteger))
	case *PdfObjectBool:
		return *_ggegf == *(_aagc.(*PdfObjectBool))
	case *PdfObjectFloat:
		return *_ggegf == *(_aagc.(*PdfObjectFloat))
	case *PdfIndirectObject:
		return _aceeg(TraceToDirectObject(_aeafd), TraceToDirectObject(_aagc), _ccfg+1)
	case *PdfObjectArray:
		_fgga := _aagc.(*PdfObjectArray)
		if len((*_ggegf)._fabc) != len((*_fgga)._fabc) {
			return false
		}
		for _ccfff, _bddee := range (*_ggegf)._fabc {
			if !_aceeg(_bddee, (*_fgga)._fabc[_ccfff], _ccfg+1) {
				return false
			}
		}
		return true
	case *PdfObjectDictionary:
		_bdfaa := _aagc.(*PdfObjectDictionary)
		_bbbba, _dgafd := (*_ggegf)._adbeb, (*_bdfaa)._adbeb
		if len(_bbbba) != len(_dgafd) {
			return false
		}
		for _bbda, _dcac := range _bbbba {
			_ffbf, _ccfe := _dgafd[_bbda]
			if !_ccfe || !_aceeg(_dcac, _ffbf, _ccfg+1) {
				return false
			}
		}
		return true
	case *PdfObjectStream:
		_ceeed := _aagc.(*PdfObjectStream)
		return _aceeg((*_ggegf).PdfObjectDictionary, (*_ceeed).PdfObjectDictionary, _ccfg+1)
	default:
		_a.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", _aeafd)
	}
	return false
}

var _bcaf = _c.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")

// MakeName creates a PdfObjectName from a string.
func MakeName(s string) *PdfObjectName { _daed := PdfObjectName(s); return &_daed }
func (_cfee *PdfCrypt) encryptBytes(_bfcf []byte, _gbbd string, _dbd []byte) ([]byte, error) {
	_a.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0062\u0079\u0074\u0065\u0073")
	_gad, _fdc := _cfee._deg[_gbbd]
	if !_fdc {
		return nil, _gf.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _gbbd)
	}
	return _gad.EncryptBytes(_bfcf, _dbd)
}

// Validate validates the page settings for the JBIG2 encoder.
func (_gadg JBIG2EncoderSettings) Validate() error {
	const _cgde = "\u0076a\u006ci\u0064\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065\u0072"
	if _gadg.Threshold < 0 || _gadg.Threshold > 1.0 {
		return _dd.Errorf(_cgde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064\u0020\u0076a\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0076\u0027 \u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061n\u0067\u0065\u0020\u005b\u0030\u002e0\u002c\u0020\u0031.\u0030\u005d", _gadg.Threshold)
	}
	if _gadg.ResolutionX < 0 {
		return _dd.Errorf(_cgde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0078\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _gadg.ResolutionX)
	}
	if _gadg.ResolutionY < 0 {
		return _dd.Errorf(_cgde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0079\u0020\u0072\u0065\u0073\u006f\u006c\u0075\u0074\u0069\u006fn\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065 \u006f\u0072\u0020\u007a\u0065\u0072o\u0020\u0076\u0061l\u0075\u0065", _gadg.ResolutionY)
	}
	if _gadg.DefaultPixelValue != 0 && _gadg.DefaultPixelValue != 1 {
		return _dd.Errorf(_cgde, "de\u0066\u0061u\u006c\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072 \u0074\u0068\u0065\u0020\u0062\u0069\u0074\u003a \u007b0\u002c\u0031}", _gadg.DefaultPixelValue)
	}
	if _gadg.Compression != JB2Generic {
		return _dd.Errorf(_cgde, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0063\u006fm\u0070\u0072\u0065\u0073s\u0069\u006f\u006e\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	}
	return nil
}

// GetBool returns the *PdfObjectBool object that is represented by a PdfObject directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetBool(obj PdfObject) (_fdcf *PdfObjectBool, _efeb bool) {
	_fdcf, _efeb = TraceToDirectObject(obj).(*PdfObjectBool)
	return _fdcf, _efeb
}

// HasEOLAfterHeader gets information if there is a EOL after the version header.
func (_bcb ParserMetadata) HasEOLAfterHeader() bool { return _bcb._cgcg }

// GetStream returns the *PdfObjectStream represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetStream(obj PdfObject) (_efccf *PdfObjectStream, _bfba bool) {
	obj = ResolveReference(obj)
	_efccf, _bfba = obj.(*PdfObjectStream)
	return _efccf, _bfba
}

// Decrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_effa *PdfCrypt) Decrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _effa.isDecrypted(obj) {
		return nil
	}
	switch _gabf := obj.(type) {
	case *PdfIndirectObject:
		_effa._ecee[_gabf] = true
		_a.Log.Trace("\u0044\u0065\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _gabf.ObjectNumber, _gabf.GenerationNumber)
		_bac := _gabf.ObjectNumber
		_gfb := _gabf.GenerationNumber
		_cbcd := _effa.Decrypt(_gabf.PdfObject, _bac, _gfb)
		if _cbcd != nil {
			return _cbcd
		}
		return nil
	case *PdfObjectStream:
		_effa._ecee[_gabf] = true
		_dce := _gabf.PdfObjectDictionary
		if _effa._aec.R != 5 {
			if _ccdb, _dec := _dce.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _dec && *_ccdb == "\u0058\u0052\u0065\u0066" {
				return nil
			}
		}
		_addd := _gabf.ObjectNumber
		_aaba := _gabf.GenerationNumber
		_a.Log.Trace("\u0044e\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _addd, _aaba)
		_bfge := _edg
		if _effa._ceae.V >= 4 {
			_bfge = _effa._abdf
			_a.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _effa._abdf)
			if _efa, _abdb := _dce.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _abdb {
				if _ged, _ede := GetName(_efa.Get(0)); _ede {
					if *_ged == "\u0043\u0072\u0079p\u0074" {
						_bfge = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _abdc, _cdc := _dce.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _cdc {
							if _cdde, _cdgb := _abdc.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _cdgb {
								if _, _eefa := _effa._deg[string(*_cdde)]; _eefa {
									_a.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_cdde)
									_bfge = string(*_cdde)
								}
							}
						}
					}
				}
			}
			_a.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _bfge)
			if _bfge == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_gfaf := _effa.Decrypt(_dce, _addd, _aaba)
		if _gfaf != nil {
			return _gfaf
		}
		_edff, _gfaf := _effa.makeKey(_bfge, uint32(_addd), uint32(_aaba), _effa._baa)
		if _gfaf != nil {
			return _gfaf
		}
		_gabf.Stream, _gfaf = _effa.decryptBytes(_gabf.Stream, _bfge, _edff)
		if _gfaf != nil {
			return _gfaf
		}
		_dce.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_gabf.Stream))))
		return nil
	case *PdfObjectString:
		_a.Log.Trace("\u0044e\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_cbd := _edg
		if _effa._ceae.V >= 4 {
			_a.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _effa._dfb)
			if _effa._dfb == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_cbd = _effa._dfb
		}
		_gc, _agde := _effa.makeKey(_cbd, uint32(parentObjNum), uint32(parentGenNum), _effa._baa)
		if _agde != nil {
			return _agde
		}
		_eeg := _gabf.Str()
		_cbfc := make([]byte, len(_eeg))
		for _gggbb := 0; _gggbb < len(_eeg); _gggbb++ {
			_cbfc[_gggbb] = _eeg[_gggbb]
		}
		if len(_cbfc) > 0 {
			_a.Log.Trace("\u0044e\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _cbfc, _cbfc)
			_cbfc, _agde = _effa.decryptBytes(_cbfc, _cbd, _gc)
			if _agde != nil {
				return _agde
			}
		}
		_gabf._aaca = string(_cbfc)
		return nil
	case *PdfObjectArray:
		for _, _cfe := range _gabf.Elements() {
			_eeea := _effa.Decrypt(_cfe, parentObjNum, parentGenNum)
			if _eeea != nil {
				return _eeea
			}
		}
		return nil
	case *PdfObjectDictionary:
		_bbfa := false
		if _cfdb := _gabf.Get("\u0054\u0079\u0070\u0065"); _cfdb != nil {
			_aebd, _fce := _cfdb.(*PdfObjectName)
			if _fce && *_aebd == "\u0053\u0069\u0067" {
				_bbfa = true
			}
		}
		for _, _aeda := range _gabf.Keys() {
			_feeg := _gabf.Get(_aeda)
			if _bbfa && string(_aeda) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_aeda) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_aeda) != "\u0050\u0072\u0065\u0076" && string(_aeda) != "\u004c\u0061\u0073\u0074" {
				_cgc := _effa.Decrypt(_feeg, parentObjNum, parentGenNum)
				if _cgc != nil {
					return _cgc
				}
			}
		}
		return nil
	}
	return nil
}

// NewASCII85Encoder makes a new ASCII85 encoder.
func NewASCII85Encoder() *ASCII85Encoder { _ebbca := &ASCII85Encoder{}; return _ebbca }

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_ceba *CCITTFaxEncoder) MakeDecodeParams() PdfObject {
	_bcge := MakeDict()
	_bcge.Set("\u004b", MakeInteger(int64(_ceba.K)))
	_bcge.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_ceba.Columns)))
	if _ceba.BlackIs1 {
		_bcge.Set("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031", MakeBool(_ceba.BlackIs1))
	}
	if _ceba.EncodedByteAlign {
		_bcge.Set("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e", MakeBool(_ceba.EncodedByteAlign))
	}
	if _ceba.EndOfLine && _ceba.K >= 0 {
		_bcge.Set("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee", MakeBool(_ceba.EndOfLine))
	}
	if _ceba.Rows != 0 && !_ceba.EndOfBlock {
		_bcge.Set("\u0052\u006f\u0077\u0073", MakeInteger(int64(_ceba.Rows)))
	}
	if !_ceba.EndOfBlock {
		_bcge.Set("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b", MakeBool(_ceba.EndOfBlock))
	}
	if _ceba.DamagedRowsBeforeError != 0 {
		_bcge.Set("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072", MakeInteger(int64(_ceba.DamagedRowsBeforeError)))
	}
	return _bcge
}

// DecodeBytes decodes a slice of JBIG2 encoded bytes and returns the results.
func (_ggfd *JBIG2Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	return _fb.DecodeBytes(encoded, _eg.Parameters{}, _ggfd.Globals)
}

// GetFilterName returns the name of the encoding filter.
func (_geddb *DCTEncoder) GetFilterName() string { return StreamEncodingFilterNameDCT }

// DecodeBytes decodes a slice of JPX encoded bytes and returns the result.
func (_acbf *JPXEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _acbf.GetFilterName())
	return encoded, ErrNoJPXDecode
}
func (_fbca *PdfParser) checkLinearizedInformation(_ccbf *PdfObjectDictionary) (bool, error) {
	var _gegg error
	_fbca._baac, _gegg = GetNumberAsInt64(_ccbf.Get("\u004c"))
	if _gegg != nil {
		return false, _gegg
	}
	_gegg = _fbca.seekToEOFMarker(_fbca._baac)
	switch _gegg {
	case nil:
		return true, nil
	case _bcaf:
		return false, nil
	default:
		return false, _gegg
	}
}

// UpdateParams updates the parameter values of the encoder.
func (_dgf *ASCII85Encoder) UpdateParams(params *PdfObjectDictionary) {}

// DecodeStream decodes a LZW encoded stream and returns the result as a
// slice of bytes.
func (_accf *LZWEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_a.Log.Trace("\u004c\u005a\u0057 \u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067")
	_a.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _accf.Predictor)
	_dbc, _aagf := _accf.DecodeBytes(streamObj.Stream)
	if _aagf != nil {
		return nil, _aagf
	}
	_a.Log.Trace("\u0020\u0049\u004e\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(streamObj.Stream), streamObj.Stream)
	_a.Log.Trace("\u004f\u0055\u0054\u003a\u0020\u0028\u0025\u0064\u0029\u0020\u0025\u0020\u0078", len(_dbc), _dbc)
	if _accf.Predictor > 1 {
		if _accf.Predictor == 2 {
			_a.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_bced := _accf.Columns * _accf.Colors
			if _bced < 1 {
				return []byte{}, nil
			}
			_cgdb := len(_dbc) / _bced
			if len(_dbc)%_bced != 0 {
				_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_dbc), _bced)
			}
			if _bced%_accf.Colors != 0 {
				return nil, _gf.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _bced, _accf.Colors)
			}
			if _bced > len(_dbc) {
				_a.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _bced, len(_dbc))
				return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_a.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_dbc), _dbc)
			_feee := _fd.NewBuffer(nil)
			for _gce := 0; _gce < _cgdb; _gce++ {
				_eagc := _dbc[_bced*_gce : _bced*(_gce+1)]
				for _acfb := _accf.Colors; _acfb < _bced; _acfb++ {
					_eagc[_acfb] = byte(int(_eagc[_acfb]+_eagc[_acfb-_accf.Colors]) % 256)
				}
				_feee.Write(_eagc)
			}
			_edgb := _feee.Bytes()
			_a.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_edgb), _edgb)
			return _edgb, nil
		} else if _accf.Predictor >= 10 && _accf.Predictor <= 15 {
			_a.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_cca := _accf.Columns*_accf.Colors + 1
			if _cca < 1 {
				return []byte{}, nil
			}
			_aae := len(_dbc) / _cca
			if len(_dbc)%_cca != 0 {
				return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_dbc), _cca)
			}
			if _cca > len(_dbc) {
				_a.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _cca, len(_dbc))
				return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_gcdg := _fd.NewBuffer(nil)
			_a.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _accf.Columns)
			_a.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_dbc), _cca, _aae)
			_bdbe := make([]byte, _cca)
			for _ecf := 0; _ecf < _cca; _ecf++ {
				_bdbe[_ecf] = 0
			}
			for _fbe := 0; _fbe < _aae; _fbe++ {
				_dac := _dbc[_cca*_fbe : _cca*(_fbe+1)]
				_adff := _dac[0]
				switch _adff {
				case 0:
				case 1:
					for _ebgb := 2; _ebgb < _cca; _ebgb++ {
						_dac[_ebgb] = byte(int(_dac[_ebgb]+_dac[_ebgb-1]) % 256)
					}
				case 2:
					for _eedff := 1; _eedff < _cca; _eedff++ {
						_dac[_eedff] = byte(int(_dac[_eedff]+_bdbe[_eedff]) % 256)
					}
				default:
					_a.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0066i\u006c\u0074\u0065\u0072\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _adff)
					return nil, _gf.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _adff)
				}
				for _efgf := 0; _efgf < _cca; _efgf++ {
					_bdbe[_efgf] = _dac[_efgf]
				}
				_gcdg.Write(_dac[1:])
			}
			_bge := _gcdg.Bytes()
			return _bge, nil
		} else {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _accf.Predictor)
			return nil, _gf.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _accf.Predictor)
		}
	}
	return _dbc, nil
}

// EncodeBytes JPX encodes the passed in slice of bytes.
func (_abeg *JPXEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _abeg.GetFilterName())
	return data, ErrNoJPXDecode
}

// DecodeBytes decodes a multi-encoded slice of bytes by passing it through the
// DecodeBytes method of the underlying encoders.
func (_bgd *MultiEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_bdfa := encoded
	var _bbca error
	for _, _geef := range _bgd._bdfb {
		_a.Log.Trace("\u004du\u006c\u0074i\u0020\u0045\u006e\u0063o\u0064\u0065\u0072 \u0044\u0065\u0063\u006f\u0064\u0065\u003a\u0020\u0041pp\u006c\u0079\u0069n\u0067\u0020F\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0076 \u0025\u0054", _geef, _geef)
		_bdfa, _bbca = _geef.DecodeBytes(_bdfa)
		if _bbca != nil {
			return nil, _bbca
		}
	}
	return _bdfa, nil
}

// NewParser creates a new parser for a PDF file via ReadSeeker. Loads the cross reference stream and trailer.
// An error is returned on failure.
func NewParser(rs _fg.ReadSeeker) (*PdfParser, error) {
	_ebfgd := &PdfParser{_dfcdg: rs, ObjCache: make(objectCache), _ffed: map[int64]bool{}, _fbaae: make([]int64, 0), _beaf: make(map[*PdfParser]*PdfParser)}
	_gaabe, _dgbaf, _cba := _ebfgd.parsePdfVersion()
	if _cba != nil {
		_a.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _cba)
		return nil, _cba
	}
	_ebfgd._eebec.Major = _gaabe
	_ebfgd._eebec.Minor = _dgbaf
	if _ebfgd._dabde, _cba = _ebfgd.loadXrefs(); _cba != nil {
		_a.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020F\u0061\u0069\u006c\u0065d t\u006f l\u006f\u0061\u0064\u0020\u0078\u0072\u0065f \u0074\u0061\u0062\u006c\u0065\u0021\u0020%\u0073", _cba)
		return nil, _cba
	}
	_a.Log.Trace("T\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0073", _ebfgd._dabde)
	_bddg, _cba := _ebfgd.parseLinearizedDictionary()
	if _cba != nil {
		return nil, _cba
	}
	if _bddg != nil {
		_ebfgd._dgcba, _cba = _ebfgd.checkLinearizedInformation(_bddg)
		if _cba != nil {
			return nil, _cba
		}
	}
	if len(_ebfgd._bbdf.ObjectMap) == 0 {
		return nil, _gf.Errorf("\u0065\u006d\u0070\u0074\u0079\u0020\u0058\u0052\u0045\u0046\u0020t\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0049\u006e\u0076a\u006c\u0069\u0064")
	}
	_ebfgd._efce = len(_ebfgd._fbaae)
	if _ebfgd._dgcba && _ebfgd._efce != 0 {
		_ebfgd._efce--
	}
	_ebfgd._gfdfd = make([]*PdfParser, _ebfgd._efce)
	return _ebfgd, nil
}

// NewRawEncoder returns a new instace of RawEncoder.
func NewRawEncoder() *RawEncoder { return &RawEncoder{} }
func (_aaaf *PdfParser) xrefNextObjectOffset(_fgea int64) int64 {
	_cebd := int64(0)
	if len(_aaaf._bbdf.ObjectMap) == 0 {
		return 0
	}
	if len(_aaaf._bbdf._cf) == 0 {
		_abff := 0
		for _, _ceccd := range _aaaf._bbdf.ObjectMap {
			if _ceccd.Offset > 0 {
				_abff++
			}
		}
		if _abff == 0 {
			return 0
		}
		_aaaf._bbdf._cf = make([]XrefObject, _abff)
		_bcgb := 0
		for _, _dgbd := range _aaaf._bbdf.ObjectMap {
			if _dgbd.Offset > 0 {
				_aaaf._bbdf._cf[_bcgb] = _dgbd
				_bcgb++
			}
		}
		_gg.Slice(_aaaf._bbdf._cf, func(_bgef, _ffdg int) bool { return _aaaf._bbdf._cf[_bgef].Offset < _aaaf._bbdf._cf[_ffdg].Offset })
	}
	_cdea := _gg.Search(len(_aaaf._bbdf._cf), func(_eggf int) bool { return _aaaf._bbdf._cf[_eggf].Offset >= _fgea })
	if _cdea < len(_aaaf._bbdf._cf) {
		_cebd = _aaaf._bbdf._cf[_cdea].Offset
	}
	return _cebd
}
func _egdfc(_bbfd *PdfObjectStream, _eacce *PdfObjectDictionary) (*JBIG2Encoder, error) {
	const _bgeag = "\u006ee\u0077\u004a\u0042\u0049G\u0032\u0044\u0065\u0063\u006fd\u0065r\u0046r\u006f\u006d\u0053\u0074\u0072\u0065\u0061m"
	_dbaee := NewJBIG2Encoder()
	_ddgf := _bbfd.PdfObjectDictionary
	if _ddgf == nil {
		return _dbaee, nil
	}
	if _eacce == nil {
		_baafg := _ddgf.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		if _baafg != nil {
			switch _gged := _baafg.(type) {
			case *PdfObjectDictionary:
				_eacce = _gged
			case *PdfObjectArray:
				if _gged.Len() == 1 {
					if _bfbc, _cgea := GetDict(_gged.Get(0)); _cgea {
						_eacce = _bfbc
					}
				}
			default:
				_a.Log.Error("\u0044\u0065\u0063\u006f\u0064\u0065P\u0061\u0072\u0061\u006d\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0025\u0023\u0076", _baafg)
				return nil, _dd.Errorf(_bgeag, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050a\u0072m\u0073\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _gged)
			}
		}
	}
	if _eacce == nil {
		return _dbaee, nil
	}
	_dbaee.UpdateParams(_eacce)
	_befda, _afgf := GetStream(_eacce.Get("\u004a\u0042\u0049G\u0032\u0047\u006c\u006f\u0062\u0061\u006c\u0073"))
	if !_afgf {
		return _dbaee, nil
	}
	var _ccad error
	_dbaee.Globals, _ccad = _fb.DecodeGlobals(_befda.Stream)
	if _ccad != nil {
		_ccad = _dd.Wrap(_ccad, _bgeag, "\u0063\u006f\u0072\u0072u\u0070\u0074\u0065\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0065n\u0063\u006f\u0064\u0065\u0064\u0020\u0064a\u0074\u0061")
		_a.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ccad)
		return nil, _ccad
	}
	return _dbaee, nil
}
func (_bcde *PdfParser) parseXrefStream(_egec *PdfObjectInteger) (*PdfObjectDictionary, error) {
	if _egec != nil {
		_a.Log.Trace("\u0058\u0052\u0065f\u0053\u0074\u006d\u0020x\u0072\u0065\u0066\u0020\u0074\u0061\u0062l\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0061\u0074\u0020\u0025\u0064", _egec)
		_bcde._dfcdg.Seek(int64(*_egec), _fg.SeekStart)
		_bcde._ffbg = _bfc.NewReader(_bcde._dfcdg)
	}
	_fgce := _bcde.GetFileOffset()
	_fbgg, _bbcc := _bcde.ParseIndirectObject()
	if _bbcc != nil {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072\u0065\u0061d\u0020\u0078\u0072\u0065\u0066\u0020\u006fb\u006a\u0065\u0063\u0074")
		return nil, _c.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0072e\u0061\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062j\u0065\u0063\u0074")
	}
	_a.Log.Trace("\u0058R\u0065f\u0053\u0074\u006d\u0020\u006fb\u006a\u0065c\u0074\u003a\u0020\u0025\u0073", _fbgg)
	_fcag, _aded := _fbgg.(*PdfObjectStream)
	if !_aded {
		_a.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0058R\u0065\u0066\u0053\u0074\u006d\u0020\u0070o\u0069\u006e\u0074\u0069\u006e\u0067 \u0074\u006f\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0021")
		return nil, _c.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0070\u006f\u0069\u006e\u0074i\u006e\u0067\u0020\u0074\u006f\u0020a\u0020\u006e\u006f\u006e\u002d\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006fb\u006a\u0065\u0063\u0074")
	}
	_cefa := _fcag.PdfObjectDictionary
	_dbbfe, _aded := _fcag.PdfObjectDictionary.Get("\u0053\u0069\u007a\u0065").(*PdfObjectInteger)
	if !_aded {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0073\u0069\u007a\u0065\u0020f\u0072\u006f\u006d\u0020\u0078\u0072\u0065f\u0020\u0073\u0074\u006d")
		return nil, _c.New("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0053\u0069\u007ae\u0020\u0066\u0072\u006f\u006d\u0020\u0078\u0072\u0065\u0066 \u0073\u0074\u006d")
	}
	if int64(*_dbbfe) > 8388607 {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u0053\u0069\u007a\u0065\u0020\u0065x\u0063\u0065\u0065\u0064\u0065\u0064\u0020l\u0069\u006d\u0069\u0074\u002c\u0020\u006f\u0076\u0065\u0072\u00208\u0033\u0038\u0038\u0036\u0030\u0037\u0020\u0028\u0025\u0064\u0029", *_dbbfe)
		return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_cfba := _fcag.PdfObjectDictionary.Get("\u0057")
	_abfg, _aded := _cfba.(*PdfObjectArray)
	if !_aded {
		return nil, _c.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0057\u0020\u0069\u006e\u0020x\u0072\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	}
	_ecacb := _abfg.Len()
	if _ecacb != 3 {
		_a.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020\u0028\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033\u0020\u002d\u0020\u0025\u0064\u0029", _ecacb)
		return nil, _c.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0078\u0072\u0065f\u0020s\u0074\u006d\u0020\u006c\u0065\u006e\u0028\u0057\u0029\u0020\u0021\u003d\u0020\u0033")
	}
	var _gedb []int64
	for _gbce := 0; _gbce < 3; _gbce++ {
		_acfd, _dfag := GetInt(_abfg.Get(_gbce))
		if !_dfag {
			return nil, _c.New("i\u006e\u0076\u0061\u006cid\u0020w\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		_gedb = append(_gedb, int64(*_acfd))
	}
	_ecfe, _bbcc := DecodeStream(_fcag)
	if _bbcc != nil {
		_a.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f \u0064e\u0063o\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _bbcc)
		return nil, _bbcc
	}
	_cdfee := int(_gedb[0])
	_eafae := int(_gedb[0] + _gedb[1])
	_dcdd := int(_gedb[0] + _gedb[1] + _gedb[2])
	_dafc := int(_gedb[0] + _gedb[1] + _gedb[2])
	if _cdfee < 0 || _eafae < 0 || _dcdd < 0 {
		_a.Log.Debug("\u0045\u0072\u0072\u006fr\u0020\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u003c \u0030 \u0028\u0025\u0064\u002c\u0025\u0064\u002c%\u0064\u0029", _cdfee, _eafae, _dcdd)
		return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	if _dafc == 0 {
		_a.Log.Debug("\u004e\u006f\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0069\u006e\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028\u0064\u0065\u006c\u0074\u0061\u0062\u0020=\u003d\u0020\u0030\u0029")
		return _cefa, nil
	}
	_gcea := len(_ecfe) / _dafc
	_dffba := 0
	_aaaeb := _fcag.PdfObjectDictionary.Get("\u0049\u006e\u0064e\u0078")
	var _eedea []int
	if _aaaeb != nil {
		_a.Log.Trace("\u0049n\u0064\u0065\u0078\u003a\u0020\u0025b", _aaaeb)
		_dcegf, _bgbg := _aaaeb.(*PdfObjectArray)
		if !_bgbg {
			_a.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0029")
			return nil, _c.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0049\u006e\u0064e\u0078\u0020\u006f\u0062je\u0063\u0074")
		}
		if _dcegf.Len()%2 != 0 {
			_a.Log.Debug("\u0057\u0041\u0052\u004eI\u004e\u0047\u0020\u0046\u0061\u0069\u006c\u0075\u0072e\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0073\u0074\u006d\u0020i\u006e\u0064\u0065\u0078\u0020n\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006f\u0066\u0020\u0032\u002e")
			return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
		}
		_dffba = 0
		_gdda, _dgbb := _dcegf.ToIntegerArray()
		if _dgbb != nil {
			_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0065\u0078 \u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0073\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0073\u003a\u0020\u0025\u0076", _dgbb)
			return nil, _dgbb
		}
		for _gdge := 0; _gdge < len(_gdda); _gdge += 2 {
			_acag := _gdda[_gdge]
			_cgdbe := _gdda[_gdge+1]
			for _dadgc := 0; _dadgc < _cgdbe; _dadgc++ {
				_eedea = append(_eedea, _acag+_dadgc)
			}
			_dffba += _cgdbe
		}
	} else {
		for _ecdb := 0; _ecdb < int(*_dbbfe); _ecdb++ {
			_eedea = append(_eedea, _ecdb)
		}
		_dffba = int(*_dbbfe)
	}
	if _gcea == _dffba+1 {
		_a.Log.Debug("\u0049n\u0063\u006f\u006d\u0070ati\u0062\u0069\u006c\u0069t\u0079\u003a\u0020\u0049\u006e\u0064\u0065\u0078\u0020\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u0076\u0065\u0072\u0061\u0067\u0065\u0020\u006f\u0066\u0020\u0031\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u002d\u0020\u0061\u0070\u0070en\u0064\u0069\u006eg\u0020\u006f\u006e\u0065\u0020-\u0020M\u0061\u0079\u0020\u006c\u0065\u0061\u0064\u0020\u0074o\u0020\u0070\u0072\u006f\u0062\u006c\u0065\u006d\u0073")
		_gbeb := _dffba - 1
		for _, _ebcb := range _eedea {
			if _ebcb > _gbeb {
				_gbeb = _ebcb
			}
		}
		_eedea = append(_eedea, _gbeb+1)
		_dffba++
	}
	if _gcea != len(_eedea) {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020x\u0072\u0065\u0066 \u0073\u0074\u006d:\u0020\u006eu\u006d\u0020\u0065\u006e\u0074\u0072i\u0065s \u0021\u003d\u0020\u006c\u0065\u006e\u0028\u0069\u006e\u0064\u0069\u0063\u0065\u0073\u0029\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", _gcea, len(_eedea))
		return nil, _c.New("\u0078\u0072ef\u0020\u0073\u0074m\u0020\u006e\u0075\u006d en\u0074ri\u0065\u0073\u0020\u0021\u003d\u0020\u006cen\u0028\u0069\u006e\u0064\u0069\u0063\u0065s\u0029")
	}
	_a.Log.Trace("\u004f\u0062j\u0065\u0063\u0074s\u0020\u0063\u006f\u0075\u006e\u0074\u0020\u0025\u0064", _dffba)
	_a.Log.Trace("\u0049\u006e\u0064i\u0063\u0065\u0073\u003a\u0020\u0025\u0020\u0064", _eedea)
	_beda := func(_ecccg []byte) int64 {
		var _eafab int64
		for _faae := 0; _faae < len(_ecccg); _faae++ {
			_eafab += int64(_ecccg[_faae]) * (1 << uint(8*(len(_ecccg)-_faae-1)))
		}
		return _eafab
	}
	_a.Log.Trace("\u0044e\u0063\u006f\u0064\u0065d\u0020\u0073\u0074\u0072\u0065a\u006d \u006ce\u006e\u0067\u0074\u0068\u003a\u0020\u0025d", len(_ecfe))
	_cefe := 0
	for _fadggb := 0; _fadggb < len(_ecfe); _fadggb += _dafc {
		_geeg := _agba(len(_ecfe), _fadggb, _fadggb+_cdfee)
		if _geeg != nil {
			_a.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _geeg)
			return nil, _geeg
		}
		_bfed := _ecfe[_fadggb : _fadggb+_cdfee]
		_geeg = _agba(len(_ecfe), _fadggb+_cdfee, _fadggb+_eafae)
		if _geeg != nil {
			_a.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _geeg)
			return nil, _geeg
		}
		_acfc := _ecfe[_fadggb+_cdfee : _fadggb+_eafae]
		_geeg = _agba(len(_ecfe), _fadggb+_eafae, _fadggb+_dcdd)
		if _geeg != nil {
			_a.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020\u0025\u0076", _geeg)
			return nil, _geeg
		}
		_dbab := _ecfe[_fadggb+_eafae : _fadggb+_dcdd]
		_aadca := _beda(_bfed)
		_dcff := _beda(_acfc)
		_eadfb := _beda(_dbab)
		if _gedb[0] == 0 {
			_aadca = 1
		}
		if _cefe >= len(_eedea) {
			_a.Log.Debug("X\u0052\u0065\u0066\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u002d\u0020\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0061\u0063\u0063e\u0073s\u0020\u0069\u006e\u0064e\u0078\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073\u0020\u002d\u0020\u0062\u0072\u0065\u0061\u006b\u0069\u006e\u0067")
			break
		}
		_edca := _eedea[_cefe]
		_cefe++
		_a.Log.Trace("%\u0064\u002e\u0020\u0070\u0031\u003a\u0020\u0025\u0020\u0078", _edca, _bfed)
		_a.Log.Trace("%\u0064\u002e\u0020\u0070\u0032\u003a\u0020\u0025\u0020\u0078", _edca, _acfc)
		_a.Log.Trace("%\u0064\u002e\u0020\u0070\u0033\u003a\u0020\u0025\u0020\u0078", _edca, _dbab)
		_a.Log.Trace("\u0025d\u002e \u0078\u0072\u0065\u0066\u003a \u0025\u0064 \u0025\u0064\u0020\u0025\u0064", _edca, _aadca, _dcff, _eadfb)
		if _aadca == 0 {
			_a.Log.Trace("-\u0020\u0046\u0072\u0065\u0065\u0020o\u0062\u006a\u0065\u0063\u0074\u0020-\u0020\u0063\u0061\u006e\u0020\u0070\u0072o\u0062\u0061\u0062\u006c\u0079\u0020\u0069\u0067\u006e\u006fr\u0065")
		} else if _aadca == 1 {
			_a.Log.Trace("\u002d\u0020I\u006e\u0020\u0075\u0073e\u0020\u002d \u0075\u006e\u0063\u006f\u006d\u0070\u0072\u0065s\u0073\u0065\u0064\u0020\u0076\u0069\u0061\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0062", _acfc)
			if _dcff == _fgce {
				_a.Log.Debug("\u0055\u0070d\u0061\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0058\u0052\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064", _edca, _fcag.ObjectNumber)
				_edca = int(_fcag.ObjectNumber)
			}
			if _edde, _bcae := _bcde._bbdf.ObjectMap[_edca]; !_bcae || int(_eadfb) > _edde.Generation {
				_afacd := XrefObject{ObjectNumber: _edca, XType: XrefTypeTableEntry, Offset: _dcff, Generation: int(_eadfb)}
				_bcde._bbdf.ObjectMap[_edca] = _afacd
			}
		} else if _aadca == 2 {
			_a.Log.Trace("\u002d\u0020\u0049\u006e \u0075\u0073\u0065\u0020\u002d\u0020\u0063\u006f\u006d\u0070r\u0065s\u0073\u0065\u0064\u0020\u006f\u0062\u006ae\u0063\u0074")
			if _, _ccgb := _bcde._bbdf.ObjectMap[_edca]; !_ccgb {
				_gabg := XrefObject{ObjectNumber: _edca, XType: XrefTypeObjectStream, OsObjNumber: int(_dcff), OsObjIndex: int(_eadfb)}
				_bcde._bbdf.ObjectMap[_edca] = _gabg
				_a.Log.Trace("\u0065\u006e\u0074\u0072\u0079\u003a\u0020\u0025\u002b\u0076", _gabg)
			}
		} else {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u0049\u004e\u0056\u0041L\u0049\u0044\u0020\u0054\u0059\u0050\u0045\u0020\u0058\u0072\u0065\u0066\u0053\u0074\u006d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u003f\u002d\u002d\u002d\u002d\u002d\u002d-")
			continue
		}
	}
	if _bcde._fddf == nil {
		_cddf := XrefTypeObjectStream
		_bcde._fddf = &_cddf
	}
	return _cefa, nil
}

// EncodeStream encodes the stream data using the encoded specified by the stream's dictionary.
func EncodeStream(streamObj *PdfObjectStream) error {
	_a.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_dfgb, _fdbgab := NewEncoderFromStream(streamObj)
	if _fdbgab != nil {
		_a.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _fdbgab)
		return _fdbgab
	}
	if _bdcd, _gabfc := _dfgb.(*LZWEncoder); _gabfc {
		_bdcd.EarlyChange = 0
		streamObj.PdfObjectDictionary.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(0))
	}
	_a.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u000a", _dfgb)
	_dgdc, _fdbgab := _dfgb.EncodeBytes(streamObj.Stream)
	if _fdbgab != nil {
		_a.Log.Debug("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _fdbgab)
		return _fdbgab
	}
	streamObj.Stream = _dgdc
	streamObj.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_dgdc))))
	return nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_ffce *ASCIIHexEncoder) MakeStreamDict() *PdfObjectDictionary {
	_fdbg := MakeDict()
	_fdbg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_ffce.GetFilterName()))
	return _fdbg
}
func (_dgccd *PdfParser) loadXrefs() (*PdfObjectDictionary, error) {
	_dgccd._bbdf.ObjectMap = make(map[int]XrefObject)
	_dgccd._fcba = make(objectStreams)
	_abdcf, _ddbe := _dgccd._dfcdg.Seek(0, _fg.SeekEnd)
	if _ddbe != nil {
		return nil, _ddbe
	}
	_a.Log.Trace("\u0066s\u0069\u007a\u0065\u003a\u0020\u0025d", _abdcf)
	_dgccd._fbaaa = _abdcf
	_ddbe = _dgccd.seekToEOFMarker(_abdcf)
	if _ddbe != nil {
		_a.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0073\u0065\u0065\u006b\u0020\u0074\u006f\u0020\u0065\u006f\u0066\u0020\u006d\u0061\u0072\u006b\u0065\u0072: \u0025\u0076", _ddbe)
		return nil, _ddbe
	}
	_bccf, _ddbe := _dgccd._dfcdg.Seek(0, _fg.SeekCurrent)
	if _ddbe != nil {
		return nil, _ddbe
	}
	var _fgca int64 = 64
	_gbad := _bccf - _fgca
	if _gbad < 0 {
		_gbad = 0
	}
	_, _ddbe = _dgccd._dfcdg.Seek(_gbad, _fg.SeekStart)
	if _ddbe != nil {
		return nil, _ddbe
	}
	_cedce := make([]byte, _fgca)
	_, _ddbe = _dgccd._dfcdg.Read(_cedce)
	if _ddbe != nil {
		_a.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u006c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u003a\u0020\u0025\u0076", _ddbe)
		return nil, _ddbe
	}
	_cadb := _adgg.FindStringSubmatch(string(_cedce))
	if len(_cadb) < 2 {
		_a.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020n\u006f\u0074\u0020f\u006fu\u006e\u0064\u0021")
		return nil, _c.New("\u0073\u0074\u0061\u0072tx\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	if len(_cadb) > 2 {
		_a.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u004du\u006c\u0074\u0069\u0070\u006c\u0065\u0020s\u0074\u0061\u0072\u0074\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029\u0021", _cedce)
		return nil, _c.New("m\u0075\u006c\u0074\u0069\u0070\u006ce\u0020\u0073\u0074\u0061\u0072\u0074\u0078\u0072\u0065f\u0020\u0065\u006et\u0072i\u0065\u0073\u003f")
	}
	_cbcgc, _ := _bd.ParseInt(_cadb[1], 10, 64)
	_a.Log.Trace("\u0073t\u0061r\u0074\u0078\u0072\u0065\u0066\u0020\u0061\u0074\u0020\u0025\u0064", _cbcgc)
	if _cbcgc > _abdcf {
		_a.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u006f\u0066f\u0073e\u0074 \u006fu\u0074\u0073\u0069\u0064\u0065\u0020\u006f\u0066\u0020\u0066\u0069\u006c\u0065")
		_a.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0072e\u0070\u0061\u0069\u0072")
		_cbcgc, _ddbe = _dgccd.repairLocateXref()
		if _ddbe != nil {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0065\u0070\u0061\u0069\u0072\u0020\u0061\u0074\u0074\u0065\u006d\u0070t\u0020\u0066\u0061\u0069\u006c\u0065\u0064 \u0028\u0025\u0073\u0029")
			return nil, _ddbe
		}
	}
	_dgccd._dfcdg.Seek(_cbcgc, _fg.SeekStart)
	_dgccd._ffbg = _bfc.NewReader(_dgccd._dfcdg)
	_degc, _ddbe := _dgccd.parseXref()
	if _ddbe != nil {
		return nil, _ddbe
	}
	_cgbd := _degc.Get("\u0058R\u0065\u0066\u0053\u0074\u006d")
	if _cgbd != nil {
		_bcfbf, _ccff := _cgbd.(*PdfObjectInteger)
		if !_ccff {
			return nil, _c.New("\u0058\u0052\u0065\u0066\u0053\u0074\u006d\u0020\u0021=\u0020\u0069\u006e\u0074")
		}
		_, _ddbe = _dgccd.parseXrefStream(_bcfbf)
		if _ddbe != nil {
			return nil, _ddbe
		}
	}
	var _abec []int64
	_ccbe := func(_ebfc int64, _gacbc []int64) bool {
		for _, _abgb := range _gacbc {
			if _abgb == _ebfc {
				return true
			}
		}
		return false
	}
	_cgbd = _degc.Get("\u0050\u0072\u0065\u0076")
	for _cgbd != nil {
		_gebe, _gfbdc := _cgbd.(*PdfObjectInteger)
		if !_gfbdc {
			_a.Log.Debug("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020P\u0072\u0065\u0076\u0020\u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u003a\u0020\u004e\u006f\u0074\u0020\u0061\u0020\u002a\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _cgbd)
			return _degc, nil
		}
		_fbgd := *_gebe
		_a.Log.Trace("\u0041\u006eot\u0068\u0065\u0072 \u0050\u0072\u0065\u0076 xr\u0065f \u0074\u0061\u0062\u006c\u0065\u0020\u006fbj\u0065\u0063\u0074\u0020\u0061\u0074\u0020%\u0064", _fbgd)
		_dgccd._dfcdg.Seek(int64(_fbgd), _fg.SeekStart)
		_dgccd._ffbg = _bfc.NewReader(_dgccd._dfcdg)
		_bfbe, _dggg := _dgccd.parseXref()
		if _dggg != nil {
			_a.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0045\u0072\u0072\u006f\u0072\u0020-\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u006c\u006f\u0061\u0064\u0069n\u0067\u0020\u0061\u006e\u006f\u0074\u0068\u0065\u0072\u0020\u0028\u0050re\u0076\u0029\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072")
			_a.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0063\u006f\u006e\u0074\u0069n\u0075\u0065\u0020\u0062\u0079\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067\u0020\u0069\u0074")
			break
		}
		_dgccd._fbaae = append(_dgccd._fbaae, int64(_fbgd))
		_cgbd = _bfbe.Get("\u0050\u0072\u0065\u0076")
		if _cgbd != nil {
			_acgc := *(_cgbd.(*PdfObjectInteger))
			if _ccbe(int64(_acgc), _abec) {
				_a.Log.Debug("\u0050\u0072ev\u0065\u006e\u0074i\u006e\u0067\u0020\u0063irc\u0075la\u0072\u0020\u0078\u0072\u0065\u0066\u0020re\u0066\u0065\u0072\u0065\u006e\u0063\u0069n\u0067")
				break
			}
			_abec = append(_abec, int64(_acgc))
		}
	}
	return _degc, nil
}

// NewJBIG2Encoder creates a new JBIG2Encoder.
func NewJBIG2Encoder() *JBIG2Encoder { return &JBIG2Encoder{_feg: _de.InitEncodeDocument(false)} }

// GetRevisionNumber returns the current version of the Pdf document.
func (_cfeb *PdfParser) GetRevisionNumber() int { return _cfeb._efce }

// PdfObjectFloat represents the primitive PDF floating point numerical object.
type PdfObjectFloat float64

const (
	JB2Generic JBIG2CompressionType = iota
	JB2SymbolCorrelation
	JB2SymbolRankHaus
)

// MakeStreamDict make a new instance of an encoding dictionary for a stream object.
func (_edc *ASCII85Encoder) MakeStreamDict() *PdfObjectDictionary {
	_gfgafd := MakeDict()
	_gfgafd.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_edc.GetFilterName()))
	return _gfgafd
}
func _deage(_dccdd uint, _ggfc, _afac float64) float64 {
	return (_ggfc + (float64(_dccdd) * (_afac - _ggfc) / 255)) * 255
}

// Version represents a version of a PDF standard.
type Version struct {
	Major int
	Minor int
}

func _ecg(_ggd PdfObject) (int64, int64, error) {
	if _ddg, _efc := _ggd.(*PdfIndirectObject); _efc {
		return _ddg.ObjectNumber, _ddg.GenerationNumber, nil
	}
	if _gge, _cee := _ggd.(*PdfObjectStream); _cee {
		return _gge.ObjectNumber, _gge.GenerationNumber, nil
	}
	return 0, 0, _c.New("\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u002f\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062je\u0063\u0074")
}

// GetUpdatedObjects returns pdf objects which were updated from the specific version (from prevParser).
func (_fdbda *PdfParser) GetUpdatedObjects(prevParser *PdfParser) (map[int64]PdfObject, error) {
	if prevParser == nil {
		return nil, _c.New("\u0070\u0072e\u0076\u0069\u006f\u0075\u0073\u0020\u0070\u0061\u0072\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	_egacc, _bbfe := _fdbda.getNumbersOfUpdatedObjects(prevParser)
	if _bbfe != nil {
		return nil, _bbfe
	}
	_decg := make(map[int64]PdfObject)
	for _, _addgb := range _egacc {
		if _cfea, _gadb := _fdbda.LookupByNumber(_addgb); _gadb == nil {
			_decg[int64(_addgb)] = _cfea
		} else {
			return nil, _gadb
		}
	}
	return _decg, nil
}

// GetObjectStreams returns the *PdfObjectStreams represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetObjectStreams(obj PdfObject) (_ffcg *PdfObjectStreams, _fcdd bool) {
	_ffcg, _fcdd = obj.(*PdfObjectStreams)
	return _ffcg, _fcdd
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_aecf *MultiEncoder) MakeDecodeParams() PdfObject {
	if len(_aecf._bdfb) == 0 {
		return nil
	}
	if len(_aecf._bdfb) == 1 {
		return _aecf._bdfb[0].MakeDecodeParams()
	}
	_bfda := MakeArray()
	_edgae := true
	for _, _gfbc := range _aecf._bdfb {
		_egadb := _gfbc.MakeDecodeParams()
		if _egadb == nil {
			_bfda.Append(MakeNull())
		} else {
			_edgae = false
			_bfda.Append(_egadb)
		}
	}
	if _edgae {
		return nil
	}
	return _bfda
}

// Seek implementation of Seek interface.
func (_bcdd *limitedReadSeeker) Seek(offset int64, whence int) (int64, error) {
	var _gcff int64
	switch whence {
	case _fg.SeekStart:
		_gcff = offset
	case _fg.SeekCurrent:
		_bbef, _gfbcd := _bcdd._fbge.Seek(0, _fg.SeekCurrent)
		if _gfbcd != nil {
			return 0, _gfbcd
		}
		_gcff = _bbef + offset
	case _fg.SeekEnd:
		_gcff = _bcdd._bcgd + offset
	}
	if _badfb := _bcdd.getError(_gcff); _badfb != nil {
		return 0, _badfb
	}
	if _, _feaaa := _bcdd._fbge.Seek(_gcff, _fg.SeekStart); _feaaa != nil {
		return 0, _feaaa
	}
	return _gcff, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_gdcb *JPXEncoder) MakeDecodeParams() PdfObject { return nil }

// PdfCryptNewDecrypt makes the document crypt handler based on the encryption dictionary
// and trailer dictionary. Returns an error on failure to process.
func PdfCryptNewDecrypt(parser *PdfParser, ed, trailer *PdfObjectDictionary) (*PdfCrypt, error) {
	_dcca := &PdfCrypt{_dfe: false, _ecee: make(map[PdfObject]bool), _fee: make(map[PdfObject]bool), _cdda: make(map[int]struct{}), _ada: parser}
	_bg, _dab := ed.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectName)
	if !_dab {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0043\u0072\u0079\u0070\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u0066\u0069\u0065\u006c\u0064\u0021")
		return _dcca, _c.New("r\u0065\u0071\u0075\u0069\u0072\u0065d\u0020\u0063\u0072\u0079\u0070\u0074 \u0066\u0069\u0065\u006c\u0064\u0020\u0046i\u006c\u0074\u0065\u0072\u0020\u006d\u0069\u0073\u0073\u0069n\u0067")
	}
	if *_bg != "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064" {
		_a.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020(%\u0073\u0029", *_bg)
		return _dcca, _c.New("\u0075n\u0073u\u0070\u0070\u006f\u0072\u0074e\u0064\u0020F\u0069\u006c\u0074\u0065\u0072")
	}
	_dcca._ceae.Filter = string(*_bg)
	if _dde, _edbe := ed.Get("\u0053u\u0062\u0046\u0069\u006c\u0074\u0065r").(*PdfObjectString); _edbe {
		_dcca._ceae.SubFilter = _dde.Str()
		_a.Log.Debug("\u0055s\u0069n\u0067\u0020\u0073\u0075\u0062f\u0069\u006ct\u0065\u0072\u0020\u0025\u0073", _dde)
	}
	if L, _abe := ed.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*PdfObjectInteger); _abe {
		if (*L % 8) != 0 {
			_a.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0065\u006ec\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u006e\u0067\u0074\u0068")
			return _dcca, _c.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0065\u006e\u0063\u0072y\u0070t\u0069o\u006e\u0020\u006c\u0065\u006e\u0067\u0074h")
		}
		_dcca._ceae.Length = int(*L)
	} else {
		_dcca._ceae.Length = 40
	}
	_dcca._ceae.V = 0
	if _fgg, _egaf := ed.Get("\u0056").(*PdfObjectInteger); _egaf {
		V := int(*_fgg)
		_dcca._ceae.V = V
		if V >= 1 && V <= 2 {
			_dcca._deg = _aea(_dcca._ceae.Length)
		} else if V >= 4 && V <= 5 {
			if _eab := _dcca.loadCryptFilters(ed); _eab != nil {
				return _dcca, _eab
			}
		} else {
			_a.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0061lg\u006f\u0020\u0056 \u003d \u0025\u0064", V)
			return _dcca, _c.New("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d")
		}
	}
	if _aeae := _eceb(&_dcca._aec, ed); _aeae != nil {
		return _dcca, _aeae
	}
	_baef := ""
	if _ebg, _fadd := trailer.Get("\u0049\u0044").(*PdfObjectArray); _fadd && _ebg.Len() >= 1 {
		_ead, _fea := GetString(_ebg.Get(0))
		if !_fea {
			return _dcca, _c.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0074r\u0061\u0069l\u0065\u0072\u0020\u0049\u0044")
		}
		_baef = _ead.Str()
	} else {
		_a.Log.Debug("\u0054\u0072ai\u006c\u0065\u0072 \u0049\u0044\u0020\u0061rra\u0079 m\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0021")
	}
	_dcca._abf = _baef
	return _dcca, nil
}

// UpdateParams updates the parameter values of the encoder.
func (_bgcd *RunLengthEncoder) UpdateParams(params *PdfObjectDictionary) {}

// MakeDictMap creates a PdfObjectDictionary initialized from a map of keys to values.
func MakeDictMap(objmap map[string]PdfObject) *PdfObjectDictionary {
	_dgdfe := MakeDict()
	return _dgdfe.Update(objmap)
}

// LookupByNumber looks up a PdfObject by object number.  Returns an error on failure.
func (_dad *PdfParser) LookupByNumber(objNumber int) (PdfObject, error) {
	_ad, _, _dada := _dad.lookupByNumberWrapper(objNumber, true)
	return _ad, _dada
}

// GetNameVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetNameVal(obj PdfObject) (_cddeb string, _aegb bool) {
	_ebge, _aegb := TraceToDirectObject(obj).(*PdfObjectName)
	if _aegb {
		return string(*_ebge), true
	}
	return
}

// RegisterCustomStreamEncoder register a custom encoder handler for certain filter.
func RegisterCustomStreamEncoder(filterName string, customStreamEncoder StreamEncoder) {
	_gfcec.Store(filterName, customStreamEncoder)
}

// Encode encodes previously prepare jbig2 document and stores it as the byte slice.
func (_gabd *JBIG2Encoder) Encode() (_afcae []byte, _ffad error) {
	const _gffd = "J\u0042I\u0047\u0032\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	if _gabd._feg == nil {
		return nil, _dd.Errorf(_gffd, "\u0064\u006f\u0063u\u006d\u0065\u006e\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0064a\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_gabd._feg.FullHeaders = _gabd.DefaultPageSettings.FileMode
	_afcae, _ffad = _gabd._feg.Encode()
	if _ffad != nil {
		return nil, _dd.Wrap(_ffad, _gffd, "")
	}
	return _afcae, nil
}

var _faaga = _ce.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

// MakeInteger creates a PdfObjectInteger from an int64.
func MakeInteger(val int64) *PdfObjectInteger                       { _adfe := PdfObjectInteger(val); return &_adfe }
func (_bbgb *offsetReader) Read(p []byte) (_cdad int, _afdbd error) { return _bbgb._gfgc.Read(p) }

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

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_edgaa *RawEncoder) MakeDecodeParams() PdfObject { return nil }

// HasDataAfterEOF checks if there is some data after EOF marker.
func (_gga ParserMetadata) HasDataAfterEOF() bool { return _gga._ccbc }

// LookupByReference looks up a PdfObject by a reference.
func (_fad *PdfParser) LookupByReference(ref PdfObjectReference) (PdfObject, error) {
	_a.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0075\u0070\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0025\u0073", ref.String())
	return _fad.LookupByNumber(int(ref.ObjectNumber))
}
func (_bdg *PdfParser) lookupByNumber(_cafa int, _gae bool) (PdfObject, bool, error) {
	_bfg, _ee := _bdg.ObjCache[_cafa]
	if _ee {
		_a.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0063a\u0063\u0068\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0025\u0064", _cafa)
		return _bfg, false, nil
	}
	if _bdg._eeaa == nil {
		_bdg._eeaa = map[int]bool{}
	}
	if _bdg._eeaa[_cafa] {
		_a.Log.Debug("ER\u0052\u004f\u0052\u003a\u0020\u004c\u006fok\u0075\u0070\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0069\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0069\u006e\u0020\u0070\u0072\u006f\u0067\u0072\u0065\u0073\u0073\u0020\u002d\u0020\u0072\u0065c\u0075\u0072\u0073\u0069\u0076\u0065 \u006c\u006f\u006f\u006b\u0075\u0070\u0020\u0061\u0074t\u0065m\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064", _cafa)
		return nil, false, _c.New("\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006f\u006f\u006b\u0075p\u0020a\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0062\u006c\u006f\u0063\u006b\u0065\u0064")
	}
	_bdg._eeaa[_cafa] = true
	defer delete(_bdg._eeaa, _cafa)
	_cef, _ee := _bdg._bbdf.ObjectMap[_cafa]
	if !_ee {
		_a.Log.Trace("\u0055\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u006c\u006f\u0063\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006e\u0020\u0078\u0072\u0065\u0066\u0073\u0021 \u002d\u0020\u0052\u0065\u0074u\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
		var _dfg PdfObjectNull
		return &_dfg, false, nil
	}
	_a.Log.Trace("L\u006fo\u006b\u0075\u0070\u0020\u006f\u0062\u006a\u0020n\u0075\u006d\u0062\u0065r \u0025\u0064", _cafa)
	if _cef.XType == XrefTypeTableEntry {
		_a.Log.Trace("\u0078r\u0065f\u006f\u0062\u006a\u0020\u006fb\u006a\u0020n\u0075\u006d\u0020\u0025\u0064", _cef.ObjectNumber)
		_a.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u0067e\u006e\u0020\u0025\u0064", _cef.Generation)
		_a.Log.Trace("\u0078\u0072\u0065\u0066\u006f\u0062\u006a\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0025\u0064", _cef.Offset)
		_bdg._dfcdg.Seek(_cef.Offset, _fg.SeekStart)
		_bdg._ffbg = _bfc.NewReader(_bdg._dfcdg)
		_aa, _dfa := _bdg.ParseIndirectObject()
		if _dfa != nil {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006ce\u0064\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0078\u0072\u0065\u0066\u0020\u0028\u0025\u0073\u0029", _dfa)
			if _gae {
				_a.Log.Debug("\u0041\u0074t\u0065\u006d\u0070\u0074i\u006e\u0067 \u0074\u006f\u0020\u0072\u0065\u0070\u0061\u0069r\u0020\u0078\u0072\u0065\u0066\u0073\u0020\u0028\u0074\u006f\u0070\u0020d\u006f\u0077\u006e\u0029")
				_cad, _bbg := _bdg.repairRebuildXrefsTopDown()
				if _bbg != nil {
					_a.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020r\u0065\u0070\u0061\u0069\u0072\u0020\u0028\u0025\u0073\u0029", _bbg)
					return nil, false, _bbg
				}
				_bdg._bbdf = *_cad
				return _bdg.lookupByNumber(_cafa, false)
			}
			return nil, false, _dfa
		}
		if _gae {
			_gbd, _, _ := _ecg(_aa)
			if int(_gbd) != _cafa {
				_a.Log.Debug("\u0049n\u0076\u0061\u006c\u0069d\u0020\u0078\u0072\u0065\u0066s\u003a \u0052e\u0062\u0075\u0069\u006c\u0064\u0069\u006eg")
				_edf := _bdg.rebuildXrefTable()
				if _edf != nil {
					return nil, false, _edf
				}
				_bdg.ObjCache = objectCache{}
				return _bdg.lookupByNumberWrapper(_cafa, false)
			}
		}
		_a.Log.Trace("\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006f\u0062\u006a")
		_bdg.ObjCache[_cafa] = _aa
		return _aa, false, nil
	} else if _cef.XType == XrefTypeObjectStream {
		_a.Log.Trace("\u0078r\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0021")
		_a.Log.Trace("\u003e\u004c\u006f\u0061\u0064\u0020\u0076\u0069\u0061\u0020\u004f\u0053\u0021")
		_a.Log.Trace("\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0069\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020%\u0064\u002f\u0025\u0064", _cef.OsObjNumber, _cef.OsObjIndex)
		if _cef.OsObjNumber == _cafa {
			_a.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0043i\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065f\u0065\u0072\u0065n\u0063e\u0021\u003f\u0021")
			return nil, true, _c.New("\u0078\u0072\u0065f \u0063\u0069\u0072\u0063\u0075\u006c\u0061\u0072\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		if _, _gfc := _bdg._bbdf.ObjectMap[_cef.OsObjNumber]; _gfc {
			_gff, _fab := _bdg.lookupObjectViaOS(_cef.OsObjNumber, _cafa)
			if _fab != nil {
				_a.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069n\u0067\u0020\u0045\u0052\u0052\u0020\u0028\u0025\u0073\u0029", _fab)
				return nil, true, _fab
			}
			_a.Log.Trace("\u003c\u004c\u006f\u0061\u0064\u0065\u0064\u0020\u0076i\u0061\u0020\u004f\u0053")
			_bdg.ObjCache[_cafa] = _gff
			if _bdg._abae != nil {
				_bdg._abae._ecee[_gff] = true
			}
			return _gff, true, nil
		}
		_a.Log.Debug("\u003f\u003f\u0020\u0042\u0065\u006c\u006f\u006eg\u0073\u0020\u0074o \u0061\u0020\u006e\u006f\u006e\u002dc\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u002e.\u002e\u0021")
		return nil, true, _c.New("\u006f\u0073\u0020\u0062\u0065\u006c\u006fn\u0067\u0073\u0020t\u006f\u0020\u0061\u0020n\u006f\u006e\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
	}
	return nil, false, _c.New("\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0078\u0072\u0065\u0066 \u0074\u0079\u0070\u0065")
}

const (
	XrefTypeTableEntry   xrefType = iota
	XrefTypeObjectStream xrefType = iota
)

// PdfObjectDictionary represents the primitive PDF dictionary/map object.
type PdfObjectDictionary struct {
	_adbeb map[PdfObjectName]PdfObject
	_caeg  []PdfObjectName
	_ccdg  *_g.Mutex
	_gdcf  *PdfParser
}

func (_caddd *PdfParser) repairSeekXrefMarker() error {
	_edbb, _eadb := _caddd._dfcdg.Seek(0, _fg.SeekEnd)
	if _eadb != nil {
		return _eadb
	}
	_ecbff := _ce.MustCompile("\u005cs\u0078\u0072\u0065\u0066\u005c\u0073*")
	var _cgff int64
	var _bged int64 = 1000
	for _cgff < _edbb {
		if _edbb <= (_bged + _cgff) {
			_bged = _edbb - _cgff
		}
		_, _edbda := _caddd._dfcdg.Seek(-_cgff-_bged, _fg.SeekEnd)
		if _edbda != nil {
			return _edbda
		}
		_dbag := make([]byte, _bged)
		_caddd._dfcdg.Read(_dbag)
		_a.Log.Trace("\u004c\u006f\u006fki\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0078\u0072\u0065\u0066\u0020\u003a\u0020\u0022\u0025\u0073\u0022", string(_dbag))
		_ccgbb := _ecbff.FindAllStringIndex(string(_dbag), -1)
		if _ccgbb != nil {
			_faaee := _ccgbb[len(_ccgbb)-1]
			_a.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _ccgbb)
			_caddd._dfcdg.Seek(-_cgff-_bged+int64(_faaee[0]), _fg.SeekEnd)
			_caddd._ffbg = _bfc.NewReader(_caddd._dfcdg)
			for {
				_fgfc, _aeaf := _caddd._ffbg.Peek(1)
				if _aeaf != nil {
					return _aeaf
				}
				_a.Log.Trace("\u0042\u003a\u0020\u0025\u0064\u0020\u0025\u0063", _fgfc[0], _fgfc[0])
				if !IsWhiteSpace(_fgfc[0]) {
					break
				}
				_caddd._ffbg.Discard(1)
			}
			return nil
		}
		_a.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_cgff += _bged
	}
	_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0058\u0072\u0065\u0066\u0020\u0074a\u0062\u006c\u0065\u0020\u006d\u0061r\u006b\u0065\u0072\u0020\u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u002e")
	return _c.New("\u0078r\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020")
}

// Clear resets the dictionary to an empty state.
func (_gccca *PdfObjectDictionary) Clear() {
	_gccca._caeg = []PdfObjectName{}
	_gccca._adbeb = map[PdfObjectName]PdfObject{}
	_gccca._ccdg = &_g.Mutex{}
}
func (_fbec *JBIG2Image) toBitmap() (_abda *_ef.Bitmap, _ccg error) {
	const _gfec = "\u004a\u0042\u0049\u00472I\u006d\u0061\u0067\u0065\u002e\u0074\u006f\u0042\u0069\u0074\u006d\u0061\u0070"
	if _fbec.Data == nil {
		return nil, _dd.Error(_gfec, "\u0069\u006d\u0061\u0067e \u0064\u0061\u0074\u0061\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _fbec.Width == 0 || _fbec.Height == 0 {
		return nil, _dd.Error(_gfec, "\u0069\u006d\u0061\u0067\u0065\u0020h\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u0077\u0069\u0064\u0074h\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if _fbec.HasPadding {
		_abda, _ccg = _ef.NewWithData(_fbec.Width, _fbec.Height, _fbec.Data)
	} else {
		_abda, _ccg = _ef.NewWithUnpaddedData(_fbec.Width, _fbec.Height, _fbec.Data)
	}
	if _ccg != nil {
		return nil, _dd.Wrap(_ccg, _gfec, "")
	}
	return _abda, nil
}
func (_gccbg *PdfParser) parseHexString() (*PdfObjectString, error) {
	_gccbg._ffbg.ReadByte()
	var _fdec _fd.Buffer
	for {
		_ffda, _egeae := _gccbg._ffbg.Peek(1)
		if _egeae != nil {
			return MakeString(""), _egeae
		}
		if _ffda[0] == '>' {
			_gccbg._ffbg.ReadByte()
			break
		}
		_bcafb, _ := _gccbg._ffbg.ReadByte()
		if _gccbg._ecgd {
			if _fd.IndexByte(_bbbe, _bcafb) == -1 {
				_gccbg._gadge._cde = true
			}
		}
		if !IsWhiteSpace(_bcafb) {
			_fdec.WriteByte(_bcafb)
		}
	}
	if _fdec.Len()%2 == 1 {
		_gccbg._gadge._gcg = true
		_fdec.WriteRune('0')
	}
	_dcba, _ := _bdc.DecodeString(_fdec.String())
	return MakeHexString(string(_dcba)), nil
}
func (_ddaa *PdfParser) readTextLine() (string, error) {
	var _gfbd _fd.Buffer
	for {
		_baca, _cccd := _ddaa._ffbg.Peek(1)
		if _cccd != nil {
			_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _cccd.Error())
			return _gfbd.String(), _cccd
		}
		if (_baca[0] != '\r') && (_baca[0] != '\n') {
			_eafb, _ := _ddaa._ffbg.ReadByte()
			_gfbd.WriteByte(_eafb)
		} else {
			break
		}
	}
	return _gfbd.String(), nil
}

// NewJPXEncoder returns a new instance of JPXEncoder.
func NewJPXEncoder() *JPXEncoder { return &JPXEncoder{} }

// ParseIndirectObject parses an indirect object from the input stream. Can also be an object stream.
// Returns the indirect object (*PdfIndirectObject) or the stream object (*PdfObjectStream).
func (_beeb *PdfParser) ParseIndirectObject() (PdfObject, error) {
	_cafb := PdfIndirectObject{}
	_cafb._cfada = _beeb
	_a.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_dcegb, _adggg := _beeb._ffbg.Peek(20)
	if _adggg != nil {
		if _adggg != _fg.EOF {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return &_cafb, _adggg
		}
	}
	_a.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_dcegb))
	_bgfd := _ebceg.FindStringSubmatchIndex(string(_dcegb))
	if len(_bgfd) < 6 {
		if _adggg == _fg.EOF {
			return nil, _adggg
		}
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_dcegb))
		return &_cafb, _c.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_beeb._ffbg.Discard(_bgfd[0])
	_a.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _bgfd)
	_bbbf := _bgfd[1] - _bgfd[0]
	_dfebd := make([]byte, _bbbf)
	_, _adggg = _beeb.ReadAtLeast(_dfebd, _bbbf)
	if _adggg != nil {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _adggg)
		return nil, _adggg
	}
	_a.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _dfebd)
	_bgag := _ebceg.FindStringSubmatch(string(_dfebd))
	if len(_bgag) < 3 {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_dfebd))
		return &_cafb, _c.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_abfa, _ := _bd.Atoi(_bgag[1])
	_ffdf, _ := _bd.Atoi(_bgag[2])
	_cafb.ObjectNumber = int64(_abfa)
	_cafb.GenerationNumber = int64(_ffdf)
	for {
		_dge, _eacb := _beeb._ffbg.Peek(2)
		if _eacb != nil {
			return &_cafb, _eacb
		}
		_a.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_dge), string(_dge))
		if IsWhiteSpace(_dge[0]) {
			_beeb.skipSpaces()
		} else if _dge[0] == '%' {
			_beeb.skipComments()
		} else if (_dge[0] == '<') && (_dge[1] == '<') {
			_a.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_cafb.PdfObject, _eacb = _beeb.ParseDict()
			_a.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _eacb)
			if _eacb != nil {
				return &_cafb, _eacb
			}
			_a.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_dge[0] == '/') || (_dge[0] == '(') || (_dge[0] == '[') || (_dge[0] == '<') {
			_cafb.PdfObject, _eacb = _beeb.parseObject()
			if _eacb != nil {
				return &_cafb, _eacb
			}
			_a.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else if _dge[0] == ']' {
			_a.Log.Debug("\u0057\u0041\u0052\u004e\u0049N\u0047\u003a\u0020\u0027\u005d\u0027 \u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u006eo\u0074\u0020\u0062\u0065i\u006e\u0067\u0020\u0075\u0073\u0065d\u0020\u0061\u0073\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0065\u006e\u0064\u0069n\u0067\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			_beeb._ffbg.Discard(1)
		} else {
			if _dge[0] == 'e' {
				_cafc, _ggfdd := _beeb.readTextLine()
				if _ggfdd != nil {
					return nil, _ggfdd
				}
				if len(_cafc) >= 6 && _cafc[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _dge[0] == 's' {
				_dge, _ = _beeb._ffbg.Peek(10)
				if string(_dge[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_geaa := 6
					if len(_dge) > 6 {
						if IsWhiteSpace(_dge[_geaa]) && _dge[_geaa] != '\r' && _dge[_geaa] != '\n' {
							_a.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0050\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_beeb._gadge._efcd = true
							_geaa++
						}
						if _dge[_geaa] == '\r' {
							_geaa++
							if _dge[_geaa] == '\n' {
								_geaa++
							}
						} else if _dge[_geaa] == '\n' {
							_geaa++
						} else {
							_beeb._gadge._efcd = true
						}
					}
					_beeb._ffbg.Discard(_geaa)
					_beeg, _gdbg := _cafb.PdfObject.(*PdfObjectDictionary)
					if !_gdbg {
						return nil, _c.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_a.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _beeg)
					_afgd, _gaef := _beeb.traceStreamLength(_beeg.Get("\u004c\u0065\u006e\u0067\u0074\u0068"))
					if _gaef != nil {
						_a.Log.Debug("\u0046\u0061\u0069l\u0020\u0074\u006f\u0020t\u0072\u0061\u0063\u0065\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0025\u0076", _gaef)
						return nil, _gaef
					}
					_a.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _afgd)
					_aecaa, _cabd := _afgd.(*PdfObjectInteger)
					if !_cabd {
						return nil, _c.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_afec := *_aecaa
					if _afec < 0 {
						return nil, _c.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					_eeag := _beeb.GetFileOffset()
					_decf := _beeb.xrefNextObjectOffset(_eeag)
					if _eeag+int64(_afec) > _decf && _decf > _eeag {
						_a.Log.Debug("E\u0078\u0070\u0065\u0063te\u0064 \u0065\u006e\u0064\u0069\u006eg\u0020\u0061\u0074\u0020\u0025\u0064", _eeag+int64(_afec))
						_a.Log.Debug("\u004e\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020\u0061t\u0020\u0025\u0064", _decf)
						_efaeb := _decf - _eeag - 17
						if _efaeb < 0 {
							return nil, _c.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020go\u0069\u006e\u0067 \u0070\u0061s\u0074\u0020\u0062\u006f\u0075\u006ed\u0061\u0072i\u0065\u0073")
						}
						_a.Log.Debug("\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u0020c\u006f\u0072\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020%\u0064\u002e\u002e\u002e", _efaeb)
						_afec = PdfObjectInteger(_efaeb)
						_beeg.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(_efaeb))
					}
					if int64(_afec) > _beeb._fbaaa {
						_a.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _c.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_dbac := make([]byte, _afec)
					_, _gaef = _beeb.ReadAtLeast(_dbac, int(_afec))
					if _gaef != nil {
						_a.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_dbac), _dbac)
						_a.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gaef)
						return nil, _gaef
					}
					_afde := PdfObjectStream{}
					_afde.Stream = _dbac
					_afde.PdfObjectDictionary = _cafb.PdfObject.(*PdfObjectDictionary)
					_afde.ObjectNumber = _cafb.ObjectNumber
					_afde.GenerationNumber = _cafb.GenerationNumber
					_afde.PdfObjectReference._cfada = _beeb
					_beeb.skipSpaces()
					_beeb._ffbg.Discard(9)
					_beeb.skipSpaces()
					return &_afde, nil
				}
			}
			_cafb.PdfObject, _eacb = _beeb.parseObject()
			if _cafb.PdfObject == nil {
				_a.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
				_cafb.PdfObject = MakeNull()
			}
			return &_cafb, _eacb
		}
	}
	if _cafb.PdfObject == nil {
		_a.Log.Debug("\u0049N\u0043\u004f\u004dP\u0041\u0054\u0049B\u0049LI\u0054\u0059\u003a\u0020\u0049\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061n \u006fb\u006a\u0065\u0063\u0074\u0020\u002d \u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u006e\u0075\u006c\u006c\u0020\u006f\u0062\u006ae\u0063\u0074")
		_cafb.PdfObject = MakeNull()
	}
	_a.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_cafb, nil
}

type objectStream struct {
	N   int
	_fc []byte
	_ca map[int]int64
}

var _bbbe = []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")

// ToIntegerArray returns a slice of all array elements as an int slice. An error is returned if the
// array non-integer objects. Each element can only be PdfObjectInteger.
func (_dgdff *PdfObjectArray) ToIntegerArray() ([]int, error) {
	var _ceceb []int
	for _, _fcbad := range _dgdff.Elements() {
		if _afdd, _fbbe := _fcbad.(*PdfObjectInteger); _fbbe {
			_ceceb = append(_ceceb, int(*_afdd))
		} else {
			return nil, ErrTypeError
		}
	}
	return _ceceb, nil
}

// DecodeStream decodes a JBIG2 encoded stream and returns the result as a slice of bytes.
func (_fdgaf *JBIG2Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _fdgaf.DecodeBytes(streamObj.Stream)
}
func (_bbf *PdfCrypt) checkAccessRights(_feac []byte) (bool, _dc.Permissions, error) {
	_bfe := _bbf.securityHandler()
	_agd, _bcec, _eed := _bfe.Authenticate(&_bbf._aec, _feac)
	if _eed != nil {
		return false, 0, _eed
	} else if _bcec == 0 || len(_agd) == 0 {
		return false, 0, nil
	}
	return true, _bcec, nil
}

// DecodeStream decodes RunLengthEncoded stream object and give back decoded bytes.
func (_eadf *RunLengthEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _eadf.DecodeBytes(streamObj.Stream)
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set.  Some other parameters are generated elsewhere.
func (_agda *DCTEncoder) MakeStreamDict() *PdfObjectDictionary {
	_eeeaf := MakeDict()
	_eeeaf.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_agda.GetFilterName()))
	return _eeeaf
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_gfee *MultiEncoder) MakeStreamDict() *PdfObjectDictionary {
	_ecgff := MakeDict()
	_ecgff.Set("\u0046\u0069\u006c\u0074\u0065\u0072", _gfee.GetFilterArray())
	for _, _bffc := range _gfee._bdfb {
		_cffb := _bffc.MakeStreamDict()
		for _, _ebgde := range _cffb.Keys() {
			_aaeae := _cffb.Get(_ebgde)
			if _ebgde != "\u0046\u0069\u006c\u0074\u0065\u0072" && _ebgde != "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073" {
				_ecgff.Set(_ebgde, _aaeae)
			}
		}
	}
	_aceg := _gfee.MakeDecodeParams()
	if _aceg != nil {
		_ecgff.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _aceg)
	}
	return _ecgff
}

// CheckAccessRights checks access rights and permissions for a specified password. If either user/owner password is
// specified, full rights are granted, otherwise the access rights are specified by the Permissions flag.
//
// The bool flag indicates that the user can access and view the file.
// The AccessPermissions shows what access the user has for editing etc.
// An error is returned if there was a problem performing the authentication.
func (_fdfb *PdfParser) CheckAccessRights(password []byte) (bool, _dc.Permissions, error) {
	if _fdfb._abae == nil {
		return true, _dc.PermOwner, nil
	}
	return _fdfb._abae.checkAccessRights(password)
}

// GetObjectNums returns a sorted list of object numbers of the PDF objects in the file.
func (_badbg *PdfParser) GetObjectNums() []int {
	var _bfaa []int
	for _, _geed := range _badbg._bbdf.ObjectMap {
		_bfaa = append(_bfaa, _geed.ObjectNumber)
	}
	_gg.Ints(_bfaa)
	return _bfaa
}

// DecodeStream decodes a multi-encoded stream by passing it through the
// DecodeStream method of the underlying encoders.
func (_agcb *MultiEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _agcb.DecodeBytes(streamObj.Stream)
}
func _feae(_dbbe *PdfObjectDictionary) (_cfce *_be.ImageBase) {
	var (
		_aabcf *PdfObjectInteger
		_cdeae bool
	)
	if _aabcf, _cdeae = _dbbe.Get("\u0057\u0069\u0064t\u0068").(*PdfObjectInteger); _cdeae {
		_cfce = &_be.ImageBase{Width: int(*_aabcf)}
	} else {
		return nil
	}
	if _aabcf, _cdeae = _dbbe.Get("\u0048\u0065\u0069\u0067\u0068\u0074").(*PdfObjectInteger); _cdeae {
		_cfce.Height = int(*_aabcf)
	}
	if _aabcf, _cdeae = _dbbe.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074").(*PdfObjectInteger); _cdeae {
		_cfce.BitsPerComponent = int(*_aabcf)
	}
	if _aabcf, _cdeae = _dbbe.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073").(*PdfObjectInteger); _cdeae {
		_cfce.ColorComponents = int(*_aabcf)
	}
	return _cfce
}

// Keys returns the list of keys in the dictionary.
// If `d` is nil returns a nil slice.
func (_badc *PdfObjectDictionary) Keys() []PdfObjectName {
	if _badc == nil {
		return nil
	}
	return _badc._caeg
}

// GetPreviousRevisionParser returns PdfParser for the previous version of the Pdf document.
func (_aabg *PdfParser) GetPreviousRevisionParser() (*PdfParser, error) {
	if _aabg._efce == 0 {
		return nil, _c.New("\u0074\u0068\u0069\u0073 i\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0072\u0065\u0076\u0069\u0073\u0069o\u006e")
	}
	if _eacd, _dfgcc := _aabg._beaf[_aabg]; _dfgcc {
		return _eacd, nil
	}
	_fcffb, _cbedb := _aabg.GetPreviousRevisionReadSeeker()
	if _cbedb != nil {
		return nil, _cbedb
	}
	_fdbd, _cbedb := NewParser(_fcffb)
	_fdbd._beaf = _aabg._beaf
	if _cbedb != nil {
		return nil, _cbedb
	}
	_aabg._beaf[_aabg] = _fdbd
	return _fdbd, nil
}

// EncodeBytes DCT encodes the passed in slice of bytes.
func (_aecg *DCTEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _gbg _cg.Image
	if _aecg.ColorComponents == 1 && _aecg.BitsPerComponent == 8 {
		_gbg = &_cg.Gray{Rect: _cg.Rect(0, 0, _aecg.Width, _aecg.Height), Pix: data, Stride: _be.BytesPerLine(_aecg.Width, _aecg.BitsPerComponent, _aecg.ColorComponents)}
	} else {
		var _gedff error
		_gbg, _gedff = _be.NewImage(_aecg.Width, _aecg.Height, _aecg.BitsPerComponent, _aecg.ColorComponents, data, nil, nil)
		if _gedff != nil {
			return nil, _gedff
		}
	}
	_egfa := _eb.Options{}
	_egfa.Quality = _aecg.Quality
	var _ecbe _fd.Buffer
	if _ced := _eb.Encode(&_ecbe, _gbg, &_egfa); _ced != nil {
		return nil, _ced
	}
	return _ecbe.Bytes(), nil
}

// IsHexadecimal checks if the PdfObjectString contains Hexadecimal data.
func (_bdbfe *PdfObjectString) IsHexadecimal() bool { return _bdbfe._gead }

// EqualObjects returns true if `obj1` and `obj2` have the same contents.
//
// NOTE: It is a good idea to flatten obj1 and obj2 with FlattenObject before calling this function
// so that contents, rather than references, can be compared.
func EqualObjects(obj1, obj2 PdfObject) bool { return _aceeg(obj1, obj2, 0) }

// ReadBytesAt reads byte content at specific offset and length within the PDF.
func (_aabe *PdfParser) ReadBytesAt(offset, len int64) ([]byte, error) {
	_fcef := _aabe.GetFileOffset()
	_, _bee := _aabe._dfcdg.Seek(offset, _fg.SeekStart)
	if _bee != nil {
		return nil, _bee
	}
	_cfed := make([]byte, len)
	_, _bee = _fg.ReadAtLeast(_aabe._dfcdg, _cfed, int(len))
	if _bee != nil {
		return nil, _bee
	}
	_aabe.SetFileOffset(_fcef)
	return _cfed, nil
}
func (_cdbf *PdfParser) parseArray() (*PdfObjectArray, error) {
	_bbefc := MakeArray()
	_cdbf._ffbg.ReadByte()
	for {
		_cdbf.skipSpaces()
		_dffb, _dccb := _cdbf._ffbg.Peek(1)
		if _dccb != nil {
			return _bbefc, _dccb
		}
		if _dffb[0] == ']' {
			_cdbf._ffbg.ReadByte()
			break
		}
		_fdae, _dccb := _cdbf.parseObject()
		if _dccb != nil {
			return _bbefc, _dccb
		}
		_bbefc.Append(_fdae)
	}
	return _bbefc, nil
}
func _eceb(_gac *_dc.StdEncryptDict, _bff *PdfObjectDictionary) error {
	R, _cfd := _bff.Get("\u0052").(*PdfObjectInteger)
	if !_cfd {
		return _c.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0052")
	}
	if *R < 2 || *R > 6 {
		return _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0052 \u0028\u0025\u0064\u0029", *R)
	}
	_gac.R = int(*R)
	O, _cfd := _bff.GetString("\u004f")
	if !_cfd {
		return _c.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u004f")
	}
	if _gac.R == 5 || _gac.R == 6 {
		if len(O) < 48 {
			return _gf.Errorf("\u004c\u0065\u006e\u0067th\u0028\u004f\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(O))
		}
	} else if len(O) != 32 {
		return _gf.Errorf("L\u0065n\u0067\u0074\u0068\u0028\u004f\u0029\u0020\u0021=\u0020\u0033\u0032\u0020(%\u0064\u0029", len(O))
	}
	_gac.O = []byte(O)
	U, _cfd := _bff.GetString("\u0055")
	if !_cfd {
		return _c.New("\u0065\u006e\u0063\u0072y\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006eg\u0020\u0055")
	}
	if _gac.R == 5 || _gac.R == 6 {
		if len(U) < 48 {
			return _gf.Errorf("\u004c\u0065\u006e\u0067th\u0028\u0055\u0029\u0020\u003c\u0020\u0034\u0038\u0020\u0028\u0025\u0064\u0029", len(U))
		}
	} else if len(U) != 32 {
		_a.Log.Debug("\u0057\u0061r\u006e\u0069\u006e\u0067\u003a\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0028\u0055\u0029\u0020\u0021\u003d\u0020\u0033\u0032\u0020(%\u0064\u0029", len(U))
	}
	_gac.U = []byte(U)
	if _gac.R >= 5 {
		OE, _bfcd := _bff.GetString("\u004f\u0045")
		if !_bfcd {
			return _c.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u004f\u0045")
		} else if len(OE) != 32 {
			return _gf.Errorf("L\u0065\u006e\u0067\u0074h(\u004fE\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(OE))
		}
		_gac.OE = []byte(OE)
		UE, _bfcd := _bff.GetString("\u0055\u0045")
		if !_bfcd {
			return _c.New("\u0065\u006ec\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006eg \u0055\u0045")
		} else if len(UE) != 32 {
			return _gf.Errorf("L\u0065\u006e\u0067\u0074h(\u0055E\u0029\u0020\u0021\u003d\u00203\u0032\u0020\u0028\u0025\u0064\u0029", len(UE))
		}
		_gac.UE = []byte(UE)
	}
	P, _cfd := _bff.Get("\u0050").(*PdfObjectInteger)
	if !_cfd {
		return _c.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074 \u0064\u0069\u0063t\u0069\u006f\u006e\u0061r\u0079\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0061\u0074\u0074\u0072")
	}
	_gac.P = _dc.Permissions(*P)
	if _gac.R == 6 {
		Perms, _cadd := _bff.GetString("\u0050\u0065\u0072m\u0073")
		if !_cadd {
			return _c.New("\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0050\u0065\u0072\u006d\u0073")
		} else if len(Perms) != 16 {
			return _gf.Errorf("\u004ce\u006e\u0067\u0074\u0068\u0028\u0050\u0065\u0072\u006d\u0073\u0029 \u0021\u003d\u0020\u0031\u0036\u0020\u0028\u0025\u0064\u0029", len(Perms))
		}
		_gac.Perms = []byte(Perms)
	}
	if _cgd, _caa := _bff.Get("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061").(*PdfObjectBool); _caa {
		_gac.EncryptMetadata = bool(*_cgd)
	} else {
		_gac.EncryptMetadata = true
	}
	return nil
}

// PdfObjectInteger represents the primitive PDF integer numerical object.
type PdfObjectInteger int64
type limitedReadSeeker struct {
	_fbge _fg.ReadSeeker
	_bcgd int64
}

func _aafcf(_decdf int) int { _cgbb := _decdf >> (_bca - 1); return (_decdf ^ _cgbb) - _cgbb }

// String returns a string representation of the *PdfObjectString.
func (_befa *PdfObjectString) String() string { return _befa._aaca }

// GetFilterName returns the name of the encoding filter.
func (_defd *ASCII85Encoder) GetFilterName() string { return StreamEncodingFilterNameASCII85 }

// GetNumbersAsFloat converts a list of pdf objects representing floats or integers to a slice of
// float64 values.
func GetNumbersAsFloat(objects []PdfObject) (_bafcf []float64, _efad error) {
	for _, _bdcb := range objects {
		_dgea, _aggc := GetNumberAsFloat(_bdcb)
		if _aggc != nil {
			return nil, _aggc
		}
		_bafcf = append(_bafcf, _dgea)
	}
	return _bafcf, nil
}

// DecodeGlobals decodes 'encoded' byte stream and returns their Globally defined segments ('Globals').
func (_efdf *JBIG2Encoder) DecodeGlobals(encoded []byte) (_fb.Globals, error) {
	return _fb.DecodeGlobals(encoded)
}
func _eecf() string { return _a.Version }

// Elements returns a slice of the PdfObject elements in the array.
// Preferred over accessing the array directly as type may be changed in future major versions (v3).
func (_cbedbf *PdfObjectStreams) Elements() []PdfObject {
	if _cbedbf == nil {
		return nil
	}
	return _cbedbf._cegef
}

const (
	_agdd = 0
	_bfa  = 1
	_fca  = 2
	_daef = 3
	_gabc = 4
)

// DecodeStream decodes the stream data and returns the decoded data.
// An error is returned upon failure.
func DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_a.Log.Trace("\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
	_eadbd, _fddfd := NewEncoderFromStream(streamObj)
	if _fddfd != nil {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fddfd)
		return nil, _fddfd
	}
	_a.Log.Trace("\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076\u000a", _eadbd)
	_cbfea, _fddfd := _eadbd.DecodeStream(streamObj)
	if _fddfd != nil {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0065\u0063\u006f\u0064\u0069n\u0067\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fddfd)
		return nil, _fddfd
	}
	return _cbfea, nil
}

const JB2ImageAutoThreshold = -1.0

// HasNonConformantStream implements core.ParserMetadata.
func (_cgdg ParserMetadata) HasNonConformantStream() bool { return _cgdg._efcd }

// GetFloat returns the *PdfObjectFloat represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetFloat(obj PdfObject) (_egdb *PdfObjectFloat, _aebg bool) {
	_egdb, _aebg = TraceToDirectObject(obj).(*PdfObjectFloat)
	return _egdb, _aebg
}
func _gfdf(_bdgb _efd.Filter, _aee _dc.AuthEvent) *PdfObjectDictionary {
	if _aee == "" {
		_aee = _dc.EventDocOpen
	}
	_gdde := MakeDict()
	_gdde.Set("\u0054\u0079\u0070\u0065", MakeName("C\u0072\u0079\u0070\u0074\u0046\u0069\u006c\u0074\u0065\u0072"))
	_gdde.Set("\u0041u\u0074\u0068\u0045\u0076\u0065\u006et", MakeName(string(_aee)))
	_gdde.Set("\u0043\u0046\u004d", MakeName(_bdgb.Name()))
	_gdde.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(_bdgb.KeyLength())))
	return _gdde
}

// Elements returns a slice of the PdfObject elements in the array.
func (_ceaad *PdfObjectArray) Elements() []PdfObject {
	if _ceaad == nil {
		return nil
	}
	return _ceaad._fabc
}

// ToFloat64Array returns a slice of all elements in the array as a float64 slice.  An error is
// returned if the array contains non-numeric objects (each element can be either PdfObjectInteger
// or PdfObjectFloat).
func (_baaca *PdfObjectArray) ToFloat64Array() ([]float64, error) {
	var _fedfb []float64
	for _, _agdad := range _baaca.Elements() {
		switch _eade := _agdad.(type) {
		case *PdfObjectInteger:
			_fedfb = append(_fedfb, float64(*_eade))
		case *PdfObjectFloat:
			_fedfb = append(_fedfb, float64(*_eade))
		default:
			return nil, ErrTypeError
		}
	}
	return _fedfb, nil
}
func (_bbcdg *PdfParser) repairLocateXref() (int64, error) {
	_fcccc := int64(1000)
	_bbcdg._dfcdg.Seek(-_fcccc, _fg.SeekCurrent)
	_aaef, _bggf := _bbcdg._dfcdg.Seek(0, _fg.SeekCurrent)
	if _bggf != nil {
		return 0, _bggf
	}
	_fcbae := make([]byte, _fcccc)
	_bbcdg._dfcdg.Read(_fcbae)
	_begc := _dgbg.FindAllStringIndex(string(_fcbae), -1)
	if len(_begc) < 1 {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0065\u0070a\u0069\u0072\u003a\u0020\u0078\u0072\u0065f\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021")
		return 0, _c.New("\u0072\u0065\u0070\u0061ir\u003a\u0020\u0078\u0072\u0065\u0066\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_ggff := int64(_begc[len(_begc)-1][0])
	_ggdb := _aaef + _ggff
	return _ggdb, nil
}

// GetFilterName returns the name of the encoding filter.
func (_fceb *JPXEncoder) GetFilterName() string { return StreamEncodingFilterNameJPX }

// PdfObjectArray represents the primitive PDF array object.
type PdfObjectArray struct{ _fabc []PdfObject }

// DecodeBytes decodes a slice of DCT encoded bytes and returns the result.
func (_ebce *DCTEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	_fabd := _fd.NewReader(encoded)
	_efaac, _aaea := _eb.Decode(_fabd)
	if _aaea != nil {
		_a.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006eg\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _aaea)
		return nil, _aaea
	}
	_abc := _efaac.Bounds()
	var _eege = make([]byte, _abc.Dx()*_abc.Dy()*_ebce.ColorComponents*_ebce.BitsPerComponent/8)
	_cfbc := 0
	switch _ebce.ColorComponents {
	case 1:
		_gdfb := []float64{_ebce.Decode[0], _ebce.Decode[1]}
		for _cgade := _abc.Min.Y; _cgade < _abc.Max.Y; _cgade++ {
			for _ebcg := _abc.Min.X; _ebcg < _abc.Max.X; _ebcg++ {
				_fbed := _efaac.At(_ebcg, _cgade)
				if _ebce.BitsPerComponent == 16 {
					_fced, _bfab := _fbed.(_ga.Gray16)
					if !_bfab {
						return nil, _c.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_dbfe := _deage(uint(_fced.Y>>8), _gdfb[0], _gdfb[1])
					_fcf := _deage(uint(_fced.Y), _gdfb[0], _gdfb[1])
					_eege[_cfbc] = byte(_dbfe)
					_cfbc++
					_eege[_cfbc] = byte(_fcf)
					_cfbc++
				} else {
					_ffbb, _eecd := _fbed.(_ga.Gray)
					if !_eecd {
						return nil, _c.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_eege[_cfbc] = byte(_deage(uint(_ffbb.Y), _gdfb[0], _gdfb[1]))
					_cfbc++
				}
			}
		}
	case 3:
		_bgfb := []float64{_ebce.Decode[0], _ebce.Decode[1]}
		_cdeg := []float64{_ebce.Decode[2], _ebce.Decode[3]}
		_dgdf := []float64{_ebce.Decode[4], _ebce.Decode[5]}
		for _ddc := _abc.Min.Y; _ddc < _abc.Max.Y; _ddc++ {
			for _adad := _abc.Min.X; _adad < _abc.Max.X; _adad++ {
				_acce := _efaac.At(_adad, _ddc)
				if _ebce.BitsPerComponent == 16 {
					_fbbd, _afca := _acce.(_ga.RGBA64)
					if !_afca {
						return nil, _c.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
					}
					_dcdf := _deage(uint(_fbbd.R>>8), _bgfb[0], _bgfb[1])
					_adbea := _deage(uint(_fbbd.R), _bgfb[0], _bgfb[1])
					_cdb := _deage(uint(_fbbd.G>>8), _cdeg[0], _cdeg[1])
					_acaa := _deage(uint(_fbbd.G), _cdeg[0], _cdeg[1])
					_fedf := _deage(uint(_fbbd.B>>8), _dgdf[0], _dgdf[1])
					_cgdd := _deage(uint(_fbbd.B), _dgdf[0], _dgdf[1])
					_eege[_cfbc] = byte(_dcdf)
					_cfbc++
					_eege[_cfbc] = byte(_adbea)
					_cfbc++
					_eege[_cfbc] = byte(_cdb)
					_cfbc++
					_eege[_cfbc] = byte(_acaa)
					_cfbc++
					_eege[_cfbc] = byte(_fedf)
					_cfbc++
					_eege[_cfbc] = byte(_cgdd)
					_cfbc++
				} else {
					_ddge, _bgea := _acce.(_ga.RGBA)
					if _bgea {
						_ded := _deage(uint(_ddge.R), _bgfb[0], _bgfb[1])
						_cff := _deage(uint(_ddge.G), _cdeg[0], _cdeg[1])
						_gcf := _deage(uint(_ddge.B), _dgdf[0], _dgdf[1])
						_eege[_cfbc] = byte(_ded)
						_cfbc++
						_eege[_cfbc] = byte(_cff)
						_cfbc++
						_eege[_cfbc] = byte(_gcf)
						_cfbc++
					} else {
						_bbe, _dedf := _acce.(_ga.YCbCr)
						if !_dedf {
							return nil, _c.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
						}
						_dddg, _egad, _dbbf, _ := _bbe.RGBA()
						_agee := _deage(uint(_dddg>>8), _bgfb[0], _bgfb[1])
						_ceefg := _deage(uint(_egad>>8), _cdeg[0], _cdeg[1])
						_gfgaf := _deage(uint(_dbbf>>8), _dgdf[0], _dgdf[1])
						_eege[_cfbc] = byte(_agee)
						_cfbc++
						_eege[_cfbc] = byte(_ceefg)
						_cfbc++
						_eege[_cfbc] = byte(_gfgaf)
						_cfbc++
					}
				}
			}
		}
	case 4:
		_ebbg := []float64{_ebce.Decode[0], _ebce.Decode[1]}
		_cbfb := []float64{_ebce.Decode[2], _ebce.Decode[3]}
		_edba := []float64{_ebce.Decode[4], _ebce.Decode[5]}
		_egdg := []float64{_ebce.Decode[6], _ebce.Decode[7]}
		for _cegg := _abc.Min.Y; _cegg < _abc.Max.Y; _cegg++ {
			for _bcdc := _abc.Min.X; _bcdc < _abc.Max.X; _bcdc++ {
				_dba := _efaac.At(_bcdc, _cegg)
				_edeb, _fdcg := _dba.(_ga.CMYK)
				if !_fdcg {
					return nil, _c.New("\u0063\u006fl\u006f\u0072\u0020t\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
				}
				_fgcg := 255 - _deage(uint(_edeb.C), _ebbg[0], _ebbg[1])
				_gcb := 255 - _deage(uint(_edeb.M), _cbfb[0], _cbfb[1])
				_gde := 255 - _deage(uint(_edeb.Y), _edba[0], _edba[1])
				_eccb := 255 - _deage(uint(_edeb.K), _egdg[0], _egdg[1])
				_eege[_cfbc] = byte(_fgcg)
				_cfbc++
				_eege[_cfbc] = byte(_gcb)
				_cfbc++
				_eege[_cfbc] = byte(_gde)
				_cfbc++
				_eege[_cfbc] = byte(_eccb)
				_cfbc++
			}
		}
	}
	return _eege, nil
}

// StreamEncoder represents the interface for all PDF stream encoders.
type StreamEncoder interface {
	GetFilterName() string
	MakeDecodeParams() PdfObject
	MakeStreamDict() *PdfObjectDictionary
	UpdateParams(_fcea *PdfObjectDictionary)
	EncodeBytes(_facd []byte) ([]byte, error)
	DecodeBytes(_cge []byte) ([]byte, error)
	DecodeStream(_acbg *PdfObjectStream) ([]byte, error)
}

// WriteString outputs the object as it is to be written to file.
func (_cbdfa *PdfObjectDictionary) WriteString() string {
	var _ggec _gd.Builder
	_ggec.WriteString("\u003c\u003c")
	for _, _bcecd := range _cbdfa._caeg {
		_dabdee := _cbdfa._adbeb[_bcecd]
		_ggec.WriteString(_bcecd.WriteString())
		_ggec.WriteString("\u0020")
		_ggec.WriteString(_dabdee.WriteString())
	}
	_ggec.WriteString("\u003e\u003e")
	return _ggec.String()
}

// WriteString outputs the object as it is to be written to file.
func (_ccgba *PdfObjectReference) WriteString() string {
	var _dfaag _gd.Builder
	_dfaag.WriteString(_bd.FormatInt(_ccgba.ObjectNumber, 10))
	_dfaag.WriteString("\u0020")
	_dfaag.WriteString(_bd.FormatInt(_ccgba.GenerationNumber, 10))
	_dfaag.WriteString("\u0020\u0052")
	return _dfaag.String()
}
func _efb(_dff *_dc.StdEncryptDict, _efgc *PdfObjectDictionary) {
	_efgc.Set("\u0052", MakeInteger(int64(_dff.R)))
	_efgc.Set("\u0050", MakeInteger(int64(_dff.P)))
	_efgc.Set("\u004f", MakeStringFromBytes(_dff.O))
	_efgc.Set("\u0055", MakeStringFromBytes(_dff.U))
	if _dff.R >= 5 {
		_efgc.Set("\u004f\u0045", MakeStringFromBytes(_dff.OE))
		_efgc.Set("\u0055\u0045", MakeStringFromBytes(_dff.UE))
		_efgc.Set("\u0045n\u0063r\u0079\u0070\u0074\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061", MakeBool(_dff.EncryptMetadata))
		if _dff.R > 5 {
			_efgc.Set("\u0050\u0065\u0072m\u0073", MakeStringFromBytes(_dff.Perms))
		}
	}
}
func (_cdcdd *PdfParser) traceStreamLength(_aabec PdfObject) (PdfObject, error) {
	_agdcc, _cfff := _aabec.(*PdfObjectReference)
	if _cfff {
		_agdb, _gcfc := _cdcdd._ffed[_agdcc.ObjectNumber]
		if _gcfc && _agdb {
			_a.Log.Debug("\u0053t\u0072\u0065a\u006d\u0020\u004c\u0065n\u0067\u0074\u0068 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 u\u006e\u0072\u0065s\u006f\u006cv\u0065\u0064\u0020\u0028\u0069\u006cl\u0065\u0067a\u006c\u0029")
			return nil, _c.New("\u0069\u006c\u006c\u0065ga\u006c\u0020\u0072\u0065\u0063\u0075\u0072\u0073\u0069\u0076\u0065\u0020\u006c\u006fo\u0070")
		}
		_cdcdd._ffed[_agdcc.ObjectNumber] = true
	}
	_gbec, _ddfb := _cdcdd.Resolve(_aabec)
	if _ddfb != nil {
		return nil, _ddfb
	}
	_a.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0065\u006e\u0067\u0074h\u003f\u0020\u0025\u0073", _gbec)
	if _cfff {
		_cdcdd._ffed[_agdcc.ObjectNumber] = false
	}
	return _gbec, nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_ccba *DCTEncoder) MakeDecodeParams() PdfObject { return nil }

// WriteString outputs the object as it is to be written to file.
func (_dcee *PdfObjectStream) WriteString() string {
	var _geff _gd.Builder
	_geff.WriteString(_bd.FormatInt(_dcee.ObjectNumber, 10))
	_geff.WriteString("\u0020\u0030\u0020\u0052")
	return _geff.String()
}

// EncodeJBIG2Image encodes 'img' into jbig2 encoded bytes stream, using default encoder settings.
func (_egdf *JBIG2Encoder) EncodeJBIG2Image(img *JBIG2Image) ([]byte, error) {
	const _aaae = "c\u006f\u0072\u0065\u002eEn\u0063o\u0064\u0065\u004a\u0042\u0049G\u0032\u0049\u006d\u0061\u0067\u0065"
	if _egaa := _egdf.AddPageImage(img, &_egdf.DefaultPageSettings); _egaa != nil {
		return nil, _dd.Wrap(_egaa, _aaae, "")
	}
	return _egdf.Encode()
}

// NewASCIIHexEncoder makes a new ASCII hex encoder.
func NewASCIIHexEncoder() *ASCIIHexEncoder { _cdfa := &ASCIIHexEncoder{}; return _cdfa }
func (_aeac *PdfParser) seekToEOFMarker(_cegfe int64) error {
	var _gcgf int64
	var _bcac int64 = 2048
	for _gcgf < _cegfe-4 {
		if _cegfe <= (_bcac + _gcgf) {
			_bcac = _cegfe - _gcgf
		}
		_, _afgfc := _aeac._dfcdg.Seek(_cegfe-_gcgf-_bcac, _fg.SeekStart)
		if _afgfc != nil {
			return _afgfc
		}
		_bcfda := make([]byte, _bcac)
		_aeac._dfcdg.Read(_bcfda)
		_a.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_bcfda))
		_aegd := _dbdb.FindAllStringIndex(string(_bcfda), -1)
		if _aegd != nil {
			_dbbb := _aegd[len(_aegd)-1]
			_a.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _aegd)
			_bgaab := _cegfe - _gcgf - _bcac + int64(_dbbb[0])
			_aeac._dfcdg.Seek(_bgaab, _fg.SeekStart)
			return nil
		}
		_a.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_gcgf += _bcac - 4
	}
	_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _bcaf
}

const _gbed = 6

// EncodeBytes encodes the passed in slice of bytes by passing it through the
// EncodeBytes method of the underlying encoders.
func (_bacg *MultiEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_cgbc := data
	var _fccd error
	for _fdba := len(_bacg._bdfb) - 1; _fdba >= 0; _fdba-- {
		_gef := _bacg._bdfb[_fdba]
		_cgbc, _fccd = _gef.EncodeBytes(_cgbc)
		if _fccd != nil {
			return nil, _fccd
		}
	}
	return _cgbc, nil
}

// Decrypt attempts to decrypt the PDF file with a specified password.  Also tries to
// decrypt with an empty password.  Returns true if successful, false otherwise.
// An error is returned when there is a problem with decrypting.
func (_cdfba *PdfParser) Decrypt(password []byte) (bool, error) {
	if _cdfba._abae == nil {
		return false, _c.New("\u0063\u0068\u0065\u0063k \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u0072s\u0074")
	}
	_aaafe, _cffe := _cdfba._abae.authenticate(password)
	if _cffe != nil {
		return false, _cffe
	}
	if !_aaafe {
		_aaafe, _cffe = _cdfba._abae.authenticate([]byte(""))
	}
	return _aaafe, _cffe
}

type xrefType int

// EncryptInfo contains an information generated by the document encrypter.
type EncryptInfo struct {
	Version

	// Encrypt is an encryption dictionary that contains all necessary parameters.
	// It should be stored in all copies of the document trailer.
	Encrypt *PdfObjectDictionary

	// ID0 and ID1 are IDs used in the trailer. Older algorithms such as RC4 uses them for encryption.
	ID0, ID1 string
}

func _gbgg(_efafb string) (PdfObjectReference, error) {
	_eeggb := PdfObjectReference{}
	_bfeg := _fafa.FindStringSubmatch(_efafb)
	if len(_bfeg) < 3 {
		_a.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _eeggb, _c.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_eadga, _ := _bd.Atoi(_bfeg[1])
	_egacd, _ := _bd.Atoi(_bfeg[2])
	_eeggb.ObjectNumber = int64(_eadga)
	_eeggb.GenerationNumber = int64(_egacd)
	return _eeggb, nil
}
func (_bba *FlateEncoder) postDecodePredict(_cbb []byte) ([]byte, error) {
	if _bba.Predictor > 1 {
		if _bba.Predictor == 2 {
			_a.Log.Trace("\u0054\u0069\u0066\u0066\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_a.Log.Trace("\u0043\u006f\u006c\u006f\u0072\u0073\u003a\u0020\u0025\u0064", _bba.Colors)
			_afe := _bba.Columns * _bba.Colors
			if _afe < 1 {
				return []byte{}, nil
			}
			_efcfd := len(_cbb) / _afe
			if len(_cbb)%_afe != 0 {
				_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020T\u0049\u0046\u0046 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002e\u002e\u002e")
				return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_cbb), _afe)
			}
			if _afe%_bba.Colors != 0 {
				return nil, _gf.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0072\u006fw\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020(\u0025\u0064\u0029\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006fr\u0073\u0020\u0025\u0064", _afe, _bba.Colors)
			}
			if _afe > len(_cbb) {
				_a.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _afe, len(_cbb))
				return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_a.Log.Trace("i\u006e\u0070\u0020\u006fut\u0044a\u0074\u0061\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(_cbb), _cbb)
			_ddeea := _fd.NewBuffer(nil)
			for _ebgd := 0; _ebgd < _efcfd; _ebgd++ {
				_bda := _cbb[_afe*_ebgd : _afe*(_ebgd+1)]
				for _bgc := _bba.Colors; _bgc < _afe; _bgc++ {
					_bda[_bgc] += _bda[_bgc-_bba.Colors]
				}
				_ddeea.Write(_bda)
			}
			_eec := _ddeea.Bytes()
			_a.Log.Trace("\u0050O\u0075t\u0044\u0061\u0074\u0061\u0020(\u0025\u0064)\u003a\u0020\u0025\u0020\u0078", len(_eec), _eec)
			return _eec, nil
		} else if _bba.Predictor >= 10 && _bba.Predictor <= 15 {
			_a.Log.Trace("\u0050\u004e\u0047 \u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
			_aecd := _bba.Columns*_bba.Colors + 1
			_cebe := len(_cbb) / _aecd
			if len(_cbb)%_aecd != 0 {
				return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u006f\u0077 \u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064/\u0025\u0064\u0029", len(_cbb), _aecd)
			}
			if _aecd > len(_cbb) {
				_a.Log.Debug("\u0052\u006fw\u0020\u006c\u0065\u006e\u0067t\u0068\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0028\u0025\u0064\u002f\u0025\u0064\u0029", _aecd, len(_cbb))
				return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			_gag := _fd.NewBuffer(nil)
			_a.Log.Trace("P\u0072\u0065\u0064\u0069ct\u006fr\u0020\u0063\u006f\u006c\u0075m\u006e\u0073\u003a\u0020\u0025\u0064", _bba.Columns)
			_a.Log.Trace("\u004ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064\u0020\u002f \u0025d\u0020=\u0020\u0025\u0064\u0020\u0072\u006f\u0077s", len(_cbb), _aecd, _cebe)
			_bfdd := make([]byte, _aecd)
			for _bbad := 0; _bbad < _aecd; _bbad++ {
				_bfdd[_bbad] = 0
			}
			_gcde := _bba.Colors
			for _dddb := 0; _dddb < _cebe; _dddb++ {
				_gfga := _cbb[_aecd*_dddb : _aecd*(_dddb+1)]
				_fbc := _gfga[0]
				switch _fbc {
				case _agdd:
				case _bfa:
					for _bgb := 1 + _gcde; _bgb < _aecd; _bgb++ {
						_gfga[_bgb] += _gfga[_bgb-_gcde]
					}
				case _fca:
					for _edec := 1; _edec < _aecd; _edec++ {
						_gfga[_edec] += _bfdd[_edec]
					}
				case _daef:
					for _ceg := 1; _ceg < _gcde+1; _ceg++ {
						_gfga[_ceg] += _bfdd[_ceg] / 2
					}
					for _fbac := _gcde + 1; _fbac < _aecd; _fbac++ {
						_gfga[_fbac] += byte((int(_gfga[_fbac-_gcde]) + int(_bfdd[_fbac])) / 2)
					}
				case _gabc:
					for _adbac := 1; _adbac < _aecd; _adbac++ {
						var _ddbc, _effb, _gccb byte
						_effb = _bfdd[_adbac]
						if _adbac >= _gcde+1 {
							_ddbc = _gfga[_adbac-_gcde]
							_gccb = _bfdd[_adbac-_gcde]
						}
						_gfga[_adbac] += _dbge(_ddbc, _effb, _gccb)
					}
				default:
					_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029\u0020\u0040\u0072o\u0077\u0020\u0025\u0064", _fbc, _dddb)
					return nil, _gf.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065r\u0020\u0062\u0079\u0074\u0065\u0020\u0028\u0025\u0064\u0029", _fbc)
				}
				copy(_bfdd, _gfga)
				_gag.Write(_gfga[1:])
			}
			_cfcb := _gag.Bytes()
			return _cfcb, nil
		} else {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072 \u0028\u0025\u0064\u0029", _bba.Predictor)
			return nil, _gf.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0070\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020(\u0025\u0064\u0029", _bba.Predictor)
		}
	}
	return _cbb, nil
}

// MakeArray creates an PdfObjectArray from a list of PdfObjects.
func MakeArray(objects ...PdfObject) *PdfObjectArray { return &PdfObjectArray{_fabc: objects} }

// RunLengthEncoder represents Run length encoding.
type RunLengthEncoder struct{}

// Set sets the PdfObject at index i of the array. An error is returned if the index is outside bounds.
func (_acee *PdfObjectArray) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_acee._fabc) {
		return _c.New("\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_acee._fabc[i] = obj
	return nil
}

// SetPredictor sets the predictor function.  Specify the number of columns per row.
// The columns indicates the number of samples per row.
// Used for grouping data together for compression.
func (_eea *FlateEncoder) SetPredictor(columns int) { _eea.Predictor = 11; _eea.Columns = columns }

// GetStringBytes is like GetStringVal except that it returns the string as a []byte.
// It is for convenience.
func GetStringBytes(obj PdfObject) (_gedfd []byte, _dfeba bool) {
	_befdf, _dfeba := TraceToDirectObject(obj).(*PdfObjectString)
	if _dfeba {
		return _befdf.Bytes(), true
	}
	return
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

var (
	ErrUnsupportedEncodingParameters = _c.New("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	ErrNoCCITTFaxDecode              = _c.New("\u0043\u0043I\u0054\u0054\u0046\u0061\u0078\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJBIG2Decode                 = _c.New("\u004a\u0042\u0049\u0047\u0032\u0044\u0065c\u006f\u0064\u0065 \u0065\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoJPXDecode                   = _c.New("\u004a\u0050\u0058\u0044\u0065c\u006f\u0064\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0079\u0065\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0065\u0064")
	ErrNoPdfVersion                  = _c.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	ErrTypeError                     = _c.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	ErrRangeError                    = _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	ErrNotSupported                  = _fgb.New("\u0066\u0065\u0061t\u0075\u0072\u0065\u0020n\u006f\u0074\u0020\u0063\u0075\u0072\u0072e\u006e\u0074\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
	ErrNotANumber                    = _c.New("\u006e\u006f\u0074 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
)

// SetIfNotNil sets the dictionary's key -> val mapping entry -IF- val is not nil.
// Note that we take care to perform a type switch.  Otherwise if we would supply a nil value
// of another type, e.g. (PdfObjectArray*)(nil), then it would not be a PdfObject(nil) and thus
// would get set.
func (_aefg *PdfObjectDictionary) SetIfNotNil(key PdfObjectName, val PdfObject) {
	if val != nil {
		switch _cacbf := val.(type) {
		case *PdfObjectName:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectDictionary:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectStream:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectString:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectNull:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectInteger:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectArray:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectBool:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectFloat:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfObjectReference:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		case *PdfIndirectObject:
			if _cacbf != nil {
				_aefg.Set(key, val)
			}
		default:
			_a.Log.Error("\u0045\u0052R\u004f\u0052\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065\u0072\u0020\u0068\u0061\u0070\u0070\u0065\u006e\u0021", val)
		}
	}
}
func (_aff *PdfParser) resolveReference(_gggd *PdfObjectReference) (PdfObject, bool, error) {
	_bdeee, _acdbg := _aff.ObjCache[int(_gggd.ObjectNumber)]
	if _acdbg {
		return _bdeee, true, nil
	}
	_bcgg, _eeed := _aff.LookupByReference(*_gggd)
	if _eeed != nil {
		return nil, false, _eeed
	}
	_aff.ObjCache[int(_gggd.ObjectNumber)] = _bcgg
	return _bcgg, false, nil
}

var _adgg = _ce.MustCompile("\u0073t\u0061r\u0074\u0078\u003f\u0072\u0065f\u005c\u0073*\u0028\u005c\u0064\u002b\u0029")

type objectStreams map[int]objectStream

// GetIndirect returns the *PdfIndirectObject represented by the PdfObject. On type mismatch the found bool flag is
// false and a nil pointer is returned.
func GetIndirect(obj PdfObject) (_dbaaf *PdfIndirectObject, _ebeg bool) {
	obj = ResolveReference(obj)
	_dbaaf, _ebeg = obj.(*PdfIndirectObject)
	return _dbaaf, _ebeg
}
func _bcee(_gage *PdfObjectStream, _aac *MultiEncoder) (*DCTEncoder, error) {
	_gdgb := NewDCTEncoder()
	_cbfe := _gage.PdfObjectDictionary
	if _cbfe == nil {
		return _gdgb, nil
	}
	_gfef := _gage.Stream
	if _aac != nil {
		_cdge, _dddf := _aac.DecodeBytes(_gfef)
		if _dddf != nil {
			return nil, _dddf
		}
		_gfef = _cdge
	}
	_bfb := _fd.NewReader(_gfef)
	_agac, _bacd := _eb.DecodeConfig(_bfb)
	if _bacd != nil {
		_a.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _bacd)
		return nil, _bacd
	}
	switch _agac.ColorModel {
	case _ga.RGBAModel:
		_gdgb.BitsPerComponent = 8
		_gdgb.ColorComponents = 3
		_gdgb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _ga.RGBA64Model:
		_gdgb.BitsPerComponent = 16
		_gdgb.ColorComponents = 3
		_gdgb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _ga.GrayModel:
		_gdgb.BitsPerComponent = 8
		_gdgb.ColorComponents = 1
		_gdgb.Decode = []float64{0.0, 1.0}
	case _ga.Gray16Model:
		_gdgb.BitsPerComponent = 16
		_gdgb.ColorComponents = 1
		_gdgb.Decode = []float64{0.0, 1.0}
	case _ga.CMYKModel:
		_gdgb.BitsPerComponent = 8
		_gdgb.ColorComponents = 4
		_gdgb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	case _ga.YCbCrModel:
		_gdgb.BitsPerComponent = 8
		_gdgb.ColorComponents = 3
		_gdgb.Decode = []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}
	default:
		return nil, _c.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_gdgb.Width = _agac.Width
	_gdgb.Height = _agac.Height
	_a.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _gdgb)
	_gdgb.Quality = DefaultJPEGQuality
	_agdc, _dda := GetArray(_cbfe.Get("\u0044\u0065\u0063\u006f\u0064\u0065"))
	if _dda {
		_aaa, _ebagb := _agdc.ToFloat64Array()
		if _ebagb != nil {
			return _gdgb, _ebagb
		}
		_gdgb.Decode = _aaa
	}
	return _gdgb, nil
}

// GetXrefTable returns the PDFs xref table.
func (_cgddg *PdfParser) GetXrefTable() XrefTable { return _cgddg._bbdf }

// Merge merges in key/values from another dictionary. Overwriting if has same keys.
// The mutated dictionary (d) is returned in order to allow method chaining.
func (_cgfd *PdfObjectDictionary) Merge(another *PdfObjectDictionary) *PdfObjectDictionary {
	if another != nil {
		for _, _abbe := range another.Keys() {
			_fdac := another.Get(_abbe)
			_cgfd.Set(_abbe, _fdac)
		}
	}
	return _cgfd
}

// NewEncoderFromStream creates a StreamEncoder based on the stream's dictionary.
func NewEncoderFromStream(streamObj *PdfObjectStream) (StreamEncoder, error) {
	_gfcc := TraceToDirectObject(streamObj.PdfObjectDictionary.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
	if _gfcc == nil {
		return NewRawEncoder(), nil
	}
	if _, _fbbg := _gfcc.(*PdfObjectNull); _fbbg {
		return NewRawEncoder(), nil
	}
	_dgaf, _bgeagd := _gfcc.(*PdfObjectName)
	if !_bgeagd {
		_decb, _bbae := _gfcc.(*PdfObjectArray)
		if !_bbae {
			return nil, _gf.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _decb.Len() == 0 {
			return NewRawEncoder(), nil
		}
		if _decb.Len() != 1 {
			_efca, _fcffbe := _bece(streamObj)
			if _fcffbe != nil {
				_a.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _fcffbe)
				return nil, _fcffbe
			}
			_a.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _efca)
			return _efca, nil
		}
		_gfcc = _decb.Get(0)
		_dgaf, _bbae = _gfcc.(*PdfObjectName)
		if !_bbae {
			return nil, _gf.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	if _dadb, _eedc := _gfcec.Load(_dgaf.String()); _eedc {
		return _dadb.(StreamEncoder), nil
	}
	switch *_dgaf {
	case StreamEncodingFilterNameFlate:
		return _age(streamObj, nil)
	case StreamEncodingFilterNameLZW:
		return _egg(streamObj, nil)
	case StreamEncodingFilterNameDCT:
		return _bcee(streamObj, nil)
	case StreamEncodingFilterNameRunLength:
		return _bcda(streamObj, nil)
	case StreamEncodingFilterNameASCIIHex:
		return NewASCIIHexEncoder(), nil
	case StreamEncodingFilterNameASCII85, "\u0041\u0038\u0035":
		return NewASCII85Encoder(), nil
	case StreamEncodingFilterNameCCITTFax:
		return _adcf(streamObj, nil)
	case StreamEncodingFilterNameJBIG2:
		return _egdfc(streamObj, nil)
	case StreamEncodingFilterNameJPX:
		return NewJPXEncoder(), nil
	}
	_a.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006fd\u0021")
	return nil, _gf.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006e\u0063o\u0064i\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u0020\u0028\u0025\u0073\u0029", *_dgaf)
}

type objectCache map[int]PdfObject

// GetString is a helper for Get that returns a string value.
// Returns false if the key is missing or a value is not a string.
func (_ebdb *PdfObjectDictionary) GetString(key PdfObjectName) (string, bool) {
	_bcafbd := _ebdb.Get(key)
	if _bcafbd == nil {
		return "", false
	}
	_daea, _ecdd := _bcafbd.(*PdfObjectString)
	if !_ecdd {
		return "", false
	}
	return _daea.Str(), true
}

// MakeArrayFromIntegers creates an PdfObjectArray from a slice of ints, where each array element is
// an PdfObjectInteger.
func MakeArrayFromIntegers(vals []int) *PdfObjectArray {
	_febe := MakeArray()
	for _, _effag := range vals {
		_febe.Append(MakeInteger(int64(_effag)))
	}
	return _febe
}

// EncodeBytes encodes data into ASCII85 encoded format.
func (_ebca *ASCII85Encoder) EncodeBytes(data []byte) ([]byte, error) {
	var _dbcb _fd.Buffer
	for _acaac := 0; _acaac < len(data); _acaac += 4 {
		_gcgg := data[_acaac]
		_febg := 1
		_gdgd := byte(0)
		if _acaac+1 < len(data) {
			_gdgd = data[_acaac+1]
			_febg++
		}
		_ecac := byte(0)
		if _acaac+2 < len(data) {
			_ecac = data[_acaac+2]
			_febg++
		}
		_fcgc := byte(0)
		if _acaac+3 < len(data) {
			_fcgc = data[_acaac+3]
			_febg++
		}
		_ceda := (uint32(_gcgg) << 24) | (uint32(_gdgd) << 16) | (uint32(_ecac) << 8) | uint32(_fcgc)
		if _ceda == 0 {
			_dbcb.WriteByte('z')
		} else {
			_geab := _ebca.base256Tobase85(_ceda)
			for _, _fec := range _geab[:_febg+1] {
				_dbcb.WriteByte(_fec + '!')
			}
		}
	}
	_dbcb.WriteString("\u007e\u003e")
	return _dbcb.Bytes(), nil
}
func _aea(_agf int) cryptFilters { return cryptFilters{_edg: _efd.NewFilterV2(_agf)} }

// GetInt returns the *PdfObjectBool object that is represented by a PdfObject either directly or indirectly
// within an indirect object. The bool flag indicates whether a match was found.
func GetInt(obj PdfObject) (_egc *PdfObjectInteger, _feaf bool) {
	_egc, _feaf = TraceToDirectObject(obj).(*PdfObjectInteger)
	return _egc, _feaf
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_baee *FlateEncoder) MakeStreamDict() *PdfObjectDictionary {
	_gfg := MakeDict()
	_gfg.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_baee.GetFilterName()))
	_ddba := _baee.MakeDecodeParams()
	if _ddba != nil {
		_gfg.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _ddba)
	}
	return _gfg
}

// Set sets the dictionary's key -> val mapping entry. Overwrites if key already set.
func (_ggdd *PdfObjectDictionary) Set(key PdfObjectName, val PdfObject) {
	_ggdd.setWithLock(key, val, true)
}

// DecodeStream decodes a JPX encoded stream and returns the result as a
// slice of bytes.
func (_bdag *JPXEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0041t\u0074\u0065\u006dpt\u0069\u006e\u0067\u0020\u0074\u006f \u0075\u0073\u0065\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067 \u0025\u0073", _bdag.GetFilterName())
	return streamObj.Stream, ErrNoJPXDecode
}

// IsOctalDigit checks if a character can be part of an octal digit string.
func IsOctalDigit(c byte) bool { return '0' <= c && c <= '7' }
func _bcda(_bcc *PdfObjectStream, _fgee *PdfObjectDictionary) (*RunLengthEncoder, error) {
	return NewRunLengthEncoder(), nil
}

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_debc *RawEncoder) MakeStreamDict() *PdfObjectDictionary { return MakeDict() }

// GetPreviousRevisionReadSeeker returns ReadSeeker for the previous version of the Pdf document.
func (_acfcd *PdfParser) GetPreviousRevisionReadSeeker() (_fg.ReadSeeker, error) {
	if _efbb := _acfcd.seekToEOFMarker(_acfcd._fbaaa - _gbed); _efbb != nil {
		return nil, _efbb
	}
	_gcggba, _eaag := _acfcd._dfcdg.Seek(0, _fg.SeekCurrent)
	if _eaag != nil {
		return nil, _eaag
	}
	_gcggba += _gbed
	return _bgaa(_acfcd._dfcdg, _gcggba)
}

// GetIntVal returns the int value represented by the PdfObject directly or indirectly if contained within an
// indirect object. On type mismatch the found bool flag returned is false and a nil pointer is returned.
func GetIntVal(obj PdfObject) (_cbfcb int, _dcef bool) {
	_dbea, _dcef := TraceToDirectObject(obj).(*PdfObjectInteger)
	if _dcef && _dbea != nil {
		return int(*_dbea), true
	}
	return 0, false
}
func (_gdded *PdfCrypt) securityHandler() _dc.StdHandler {
	if _gdded._aec.R >= 5 {
		return _dc.NewHandlerR6()
	}
	return _dc.NewHandlerR4(_gdded._abf, _gdded._ceae.Length)
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

// PdfCryptNewEncrypt makes the document crypt handler based on a specified crypt filter.
func PdfCryptNewEncrypt(cf _efd.Filter, userPass, ownerPass []byte, perm _dc.Permissions) (*PdfCrypt, *EncryptInfo, error) {
	_abd := &PdfCrypt{_fee: make(map[PdfObject]bool), _deg: make(cryptFilters), _aec: _dc.StdEncryptDict{P: perm, EncryptMetadata: true}}
	var _gbe Version
	if cf != nil {
		_ece := cf.PDFVersion()
		_gbe.Major, _gbe.Minor = _ece[0], _ece[1]
		V, R := cf.HandlerVersion()
		_abd._ceae.V = V
		_abd._aec.R = R
		_abd._ceae.Length = cf.KeyLength() * 8
	}
	const (
		_gggb = _edg
	)
	_abd._deg[_gggb] = cf
	if _abd._ceae.V >= 4 {
		_abd._abdf = _gggb
		_abd._dfb = _gggb
	}
	_gab := _abd.newEncryptDict()
	_cdf := _bfd.Sum([]byte(_ba.Now().Format(_ba.RFC850)))
	_eac := string(_cdf[:])
	_dee := make([]byte, 100)
	_f.Read(_dee)
	_cdf = _bfd.Sum(_dee)
	_bab := string(_cdf[:])
	_a.Log.Trace("\u0052\u0061\u006e\u0064\u006f\u006d\u0020\u0062\u003a\u0020\u0025\u0020\u0078", _dee)
	_a.Log.Trace("\u0047\u0065\u006e\u0020\u0049\u0064\u0020\u0030\u003a\u0020\u0025\u0020\u0078", _eac)
	_abd._abf = _eac
	_deea := _abd.generateParams(userPass, ownerPass)
	if _deea != nil {
		return nil, nil, _deea
	}
	_efb(&_abd._aec, _gab)
	if _abd._ceae.V >= 4 {
		if _eff := _abd.saveCryptFilters(_gab); _eff != nil {
			return nil, nil, _eff
		}
	}
	return _abd, &EncryptInfo{Version: _gbe, Encrypt: _gab, ID0: _eac, ID1: _bab}, nil
}

// Set sets the PdfObject at index i of the streams. An error is returned if the index is outside bounds.
func (_badb *PdfObjectStreams) Set(i int, obj PdfObject) error {
	if i < 0 || i >= len(_badb._cegef) {
		return _c.New("\u004f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0062o\u0075\u006e\u0064\u0073")
	}
	_badb._cegef[i] = obj
	return nil
}

// UpdateParams updates the parameter values of the encoder.
func (_gaaa *CCITTFaxEncoder) UpdateParams(params *PdfObjectDictionary) {
	if _dcge, _dfcf := GetNumberAsInt64(params.Get("\u004b")); _dfcf == nil {
		_gaaa.K = int(_dcge)
	}
	if _gffe, _bdcadg := GetNumberAsInt64(params.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")); _bdcadg == nil {
		_gaaa.Columns = int(_gffe)
	} else if _gffe, _bdcadg = GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068")); _bdcadg == nil {
		_gaaa.Columns = int(_gffe)
	}
	if _ecfa, _bdcc := GetNumberAsInt64(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _bdcc == nil {
		_gaaa.BlackIs1 = _ecfa > 0
	} else {
		if _dace, _dbaa := GetBoolVal(params.Get("\u0042\u006c\u0061\u0063\u006b\u0049\u0073\u0031")); _dbaa {
			_gaaa.BlackIs1 = _dace
		} else {
			if _bafc, _edefd := GetArray(params.Get("\u0044\u0065\u0063\u006f\u0064\u0065")); _edefd {
				_eafa, _gaed := _bafc.ToIntegerArray()
				if _gaed == nil {
					_gaaa.BlackIs1 = _eafa[0] == 1 && _eafa[1] == 0
				}
			}
		}
	}
	if _cfbe, _eaef := GetNumberAsInt64(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _eaef == nil {
		_gaaa.EncodedByteAlign = _cfbe > 0
	} else {
		if _gagc, _gbga := GetBoolVal(params.Get("\u0045\u006ec\u006f\u0064\u0065d\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e")); _gbga {
			_gaaa.EncodedByteAlign = _gagc
		}
	}
	if _fbef, _fda := GetNumberAsInt64(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _fda == nil {
		_gaaa.EndOfLine = _fbef > 0
	} else {
		if _efgb, _gdee := GetBoolVal(params.Get("\u0045n\u0064\u004f\u0066\u004c\u0069\u006ee")); _gdee {
			_gaaa.EndOfLine = _efgb
		}
	}
	if _cbed, _aadc := GetNumberAsInt64(params.Get("\u0052\u006f\u0077\u0073")); _aadc == nil {
		_gaaa.Rows = int(_cbed)
	} else if _cbed, _aadc = GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _aadc == nil {
		_gaaa.Rows = int(_cbed)
	}
	if _acfbc, _agag := GetNumberAsInt64(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _agag == nil {
		_gaaa.EndOfBlock = _acfbc > 0
	} else {
		if _edab, _dgcd := GetBoolVal(params.Get("\u0045\u006e\u0064\u004f\u0066\u0042\u006c\u006f\u0063\u006b")); _dgcd {
			_gaaa.EndOfBlock = _edab
		}
	}
	if _agb, _cccg := GetNumberAsInt64(params.Get("\u0044\u0061\u006d\u0061ge\u0064\u0052\u006f\u0077\u0073\u0042\u0065\u0066\u006f\u0072\u0065\u0045\u0072\u0072o\u0072")); _cccg != nil {
		_gaaa.DamagedRowsBeforeError = int(_agb)
	}
}
func _dbge(_ffga, _gbgb, _efaf uint8) uint8 {
	_dedeb := int(_efaf)
	_efgbb := int(_gbgb) - _dedeb
	_badfd := int(_ffga) - _dedeb
	_dedeb = _aafcf(_efgbb + _badfd)
	_efgbb = _aafcf(_efgbb)
	_badfd = _aafcf(_badfd)
	if _efgbb <= _badfd && _efgbb <= _dedeb {
		return _ffga
	} else if _badfd <= _dedeb {
		return _gbgb
	}
	return _efaf
}
func (_gfdg *PdfParser) skipSpaces() (int, error) {
	_bcfa := 0
	for {
		_abdbf, _dafg := _gfdg._ffbg.ReadByte()
		if _dafg != nil {
			return 0, _dafg
		}
		if IsWhiteSpace(_abdbf) {
			_bcfa++
		} else {
			_gfdg._ffbg.UnreadByte()
			break
		}
	}
	return _bcfa, nil
}

// MultiEncoder supports serial encoding.
type MultiEncoder struct{ _bdfb []StreamEncoder }

// DecodeStream decodes a FlateEncoded stream object and give back decoded bytes.
func (_ebb *FlateEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	_a.Log.Trace("\u0046l\u0061t\u0065\u0044\u0065\u0063\u006fd\u0065\u0020s\u0074\u0072\u0065\u0061\u006d")
	_a.Log.Trace("\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u003a\u0020\u0025\u0064", _ebb.Predictor)
	if _ebb.BitsPerComponent != 8 {
		return nil, _gf.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u003d\u0025\u0064\u0020\u0028\u006f\u006e\u006c\u0079\u0020\u0038\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0029", _ebb.BitsPerComponent)
	}
	_cecf, _faddf := _ebb.DecodeBytes(streamObj.Stream)
	if _faddf != nil {
		return nil, _faddf
	}
	_cecf, _faddf = _ebb.postDecodePredict(_cecf)
	if _faddf != nil {
		return nil, _faddf
	}
	return _cecf, nil
}

// PdfObjectStream represents the primitive PDF Object stream.
type PdfObjectStream struct {
	PdfObjectReference
	*PdfObjectDictionary
	Stream []byte
}

const _ddfg = 10

// SetImage sets the image base for given flate encoder.
func (_egb *FlateEncoder) SetImage(img *_be.ImageBase) { _egb._dagc = img }

// Read implementation of Read interface.
func (_dgcc *limitedReadSeeker) Read(p []byte) (_dfce int, _defe error) {
	_adac, _defe := _dgcc._fbge.Seek(0, _fg.SeekCurrent)
	if _defe != nil {
		return 0, _defe
	}
	_dfcd := _dgcc._bcgd - _adac
	if _dfcd == 0 {
		return 0, _fg.EOF
	}
	if _gdfd := int64(len(p)); _gdfd < _dfcd {
		_dfcd = _gdfd
	}
	_fggb := make([]byte, _dfcd)
	_dfce, _defe = _dgcc._fbge.Read(_fggb)
	copy(p, _fggb)
	return _dfce, _defe
}

// Len returns the number of elements in the array.
func (_cfcf *PdfObjectArray) Len() int {
	if _cfcf == nil {
		return 0
	}
	return len(_cfcf._fabc)
}

// MakeEncodedString creates a PdfObjectString with encoded content, which can be either
// UTF-16BE or PDFDocEncoding depending on whether `utf16BE` is true or false respectively.
func MakeEncodedString(s string, utf16BE bool) *PdfObjectString {
	if utf16BE {
		var _fgae _fd.Buffer
		_fgae.Write([]byte{0xFE, 0xFF})
		_fgae.WriteString(_da.StringToUTF16(s))
		return &PdfObjectString{_aaca: _fgae.String(), _gead: true}
	}
	return &PdfObjectString{_aaca: string(_da.StringToPDFDocEncoding(s)), _gead: false}
}
func (_cecc *PdfParser) skipComments() error {
	if _, _ccbae := _cecc.skipSpaces(); _ccbae != nil {
		return _ccbae
	}
	_dffg := true
	for {
		_cbcgg, _gegb := _cecc._ffbg.Peek(1)
		if _gegb != nil {
			_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gegb.Error())
			return _gegb
		}
		if _dffg && _cbcgg[0] != '%' {
			return nil
		}
		_dffg = false
		if (_cbcgg[0] != '\r') && (_cbcgg[0] != '\n') {
			_cecc._ffbg.ReadByte()
		} else {
			break
		}
	}
	return _cecc.skipComments()
}
func (_bafa *PdfParser) parseDetailedHeader() (_gaa error) {
	_bafa._dfcdg.Seek(0, _fg.SeekStart)
	_bafa._ffbg = _bfc.NewReader(_bafa._dfcdg)
	_cdcf := 20
	_ddgb := make([]byte, _cdcf)
	var (
		_aefa bool
		_cdac int
	)
	for {
		_edd, _def := _bafa._ffbg.ReadByte()
		if _def != nil {
			if _def == _fg.EOF {
				break
			} else {
				return _def
			}
		}
		if IsDecimalDigit(_edd) && _ddgb[_cdcf-1] == '.' && IsDecimalDigit(_ddgb[_cdcf-2]) && _ddgb[_cdcf-3] == '-' && _ddgb[_cdcf-4] == 'F' && _ddgb[_cdcf-5] == 'D' && _ddgb[_cdcf-6] == 'P' && _ddgb[_cdcf-7] == '%' {
			_bafa._eebec = Version{Major: int(_ddgb[_cdcf-2] - '0'), Minor: int(_edd - '0')}
			_bafa._gadge._bef = _cdac - 7
			_aefa = true
			break
		}
		_cdac++
		_ddgb = append(_ddgb[1:_cdcf], _edd)
	}
	if !_aefa {
		return _gf.Errorf("n\u006f \u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066ou\u006e\u0064")
	}
	_daec, _gaa := _bafa._ffbg.ReadByte()
	if _gaa == _fg.EOF {
		return _gf.Errorf("\u006eo\u0074\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0050d\u0066\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _gaa != nil {
		return _gaa
	}
	_bafa._gadge._cgcg = _daec == '\n'
	_daec, _gaa = _bafa._ffbg.ReadByte()
	if _gaa != nil {
		return _gf.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _gaa)
	}
	if _daec != '%' {
		return nil
	}
	_bfgd := make([]byte, 4)
	_, _gaa = _bafa._ffbg.Read(_bfgd)
	if _gaa != nil {
		return _gf.Errorf("\u006e\u006f\u0074\u0020a\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0064\u0066 \u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u003a \u0025\u0077", _gaa)
	}
	_bafa._gadge._egf = [4]byte{_bfgd[0], _bfgd[1], _bfgd[2], _bfgd[3]}
	return nil
}

// String returns a string describing `ref`.
func (_cebc *PdfObjectReference) String() string {
	return _gf.Sprintf("\u0052\u0065\u0066\u0028\u0025\u0064\u0020\u0025\u0064\u0029", _cebc.ObjectNumber, _cebc.GenerationNumber)
}

const (
	DefaultJPEGQuality = 75
)

// UpdateParams updates the parameter values of the encoder.
func (_dbaeb *MultiEncoder) UpdateParams(params *PdfObjectDictionary) {
	for _, _abge := range _dbaeb._bdfb {
		_abge.UpdateParams(params)
	}
}

// IsFloatDigit checks if a character can be a part of a float number string.
func IsFloatDigit(c byte) bool { return ('0' <= c && c <= '9') || c == '.' }

// PdfObjectReference represents the primitive PDF reference object.
type PdfObjectReference struct {
	_cfada           *PdfParser
	ObjectNumber     int64
	GenerationNumber int64
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
func GoImageToJBIG2(i _cg.Image, bwThreshold float64) (*JBIG2Image, error) {
	const _gfde = "\u0047\u006f\u0049\u006d\u0061\u0067\u0065\u0054\u006fJ\u0042\u0049\u0047\u0032"
	if i == nil {
		return nil, _dd.Error(_gfde, "i\u006d\u0061\u0067\u0065 '\u0069'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	var (
		_gfad  uint8
		_acceb _be.Image
		_fadgg error
	)
	if bwThreshold == JB2ImageAutoThreshold {
		_acceb, _fadgg = _be.MonochromeConverter.Convert(i)
	} else if bwThreshold > 1.0 || bwThreshold < 0.0 {
		return nil, _dd.Error(_gfde, "p\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074h\u0072\u0065\u0073\u0068\u006f\u006c\u0064 i\u0073\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u0061\u0020\u0072\u0061\u006ege\u0020\u007b0\u002e\u0030\u002c\u0020\u0031\u002e\u0030\u007d")
	} else {
		_gfad = uint8(255 * bwThreshold)
		_acceb, _fadgg = _be.MonochromeThresholdConverter(_gfad).Convert(i)
	}
	if _fadgg != nil {
		return nil, _fadgg
	}
	return _cbcac(_acceb), nil
}

// ParserMetadata is the parser based metadata information about document.
// The data here could be used on document verification.
type ParserMetadata struct {
	_bef  int
	_cgcg bool
	_egf  [4]byte
	_ccbc bool
	_gcg  bool
	_cde  bool
	_efcd bool
	_gfe  bool
	_bdf  bool
}

func (_bbbb *PdfParser) seekPdfVersionTopDown() (int, int, error) {
	_bbbb._dfcdg.Seek(0, _fg.SeekStart)
	_bbbb._ffbg = _bfc.NewReader(_bbbb._dfcdg)
	_ffba := 20
	_cabde := make([]byte, _ffba)
	for {
		_bcdfgc, _aced := _bbbb._ffbg.ReadByte()
		if _aced != nil {
			if _aced == _fg.EOF {
				break
			} else {
				return 0, 0, _aced
			}
		}
		if IsDecimalDigit(_bcdfgc) && _cabde[_ffba-1] == '.' && IsDecimalDigit(_cabde[_ffba-2]) && _cabde[_ffba-3] == '-' && _cabde[_ffba-4] == 'F' && _cabde[_ffba-5] == 'D' && _cabde[_ffba-6] == 'P' {
			_agbd := int(_cabde[_ffba-2] - '0')
			_fabe := int(_bcdfgc - '0')
			return _agbd, _fabe, nil
		}
		_cabde = append(_cabde[1:_ffba], _bcdfgc)
	}
	return 0, 0, _c.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// GetEncryptObj returns the PdfIndirectObject which has information about the PDFs encryption details.
func (_debbc *PdfParser) GetEncryptObj() *PdfIndirectObject { return _debbc._fafd }

// DecodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_dadd *RawEncoder) DecodeBytes(encoded []byte) ([]byte, error) { return encoded, nil }

// NewFlateEncoder makes a new flate encoder with default parameters, predictor 1 and bits per component 8.
func NewFlateEncoder() *FlateEncoder {
	_fbbf := &FlateEncoder{}
	_fbbf.Predictor = 1
	_fbbf.BitsPerComponent = 8
	_fbbf.Colors = 1
	_fbbf.Columns = 1
	return _fbbf
}

// DecodeStream decodes the stream containing CCITTFax encoded image data.
func (_bade *CCITTFaxEncoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _bade.DecodeBytes(streamObj.Stream)
}

// MakeObjectStreams creates an PdfObjectStreams from a list of PdfObjects.
func MakeObjectStreams(objects ...PdfObject) *PdfObjectStreams {
	return &PdfObjectStreams{_cegef: objects}
}
func _bece(_gbff *PdfObjectStream) (*MultiEncoder, error) {
	_ege := NewMultiEncoder()
	_dgfg := _gbff.PdfObjectDictionary
	if _dgfg == nil {
		return _ege, nil
	}
	var _dcfa *PdfObjectDictionary
	var _bbegg []PdfObject
	_eefc := _dgfg.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
	if _eefc != nil {
		_bgg, _gdfc := _eefc.(*PdfObjectDictionary)
		if _gdfc {
			_dcfa = _bgg
		}
		_edgbg, _ffbd := _eefc.(*PdfObjectArray)
		if _ffbd {
			for _, _cdeeg := range _edgbg.Elements() {
				_cdeeg = TraceToDirectObject(_cdeeg)
				if _eaaa, _ggcb := _cdeeg.(*PdfObjectDictionary); _ggcb {
					_bbegg = append(_bbegg, _eaaa)
				} else {
					_bbegg = append(_bbegg, MakeDict())
				}
			}
		}
	}
	_eefc = _dgfg.Get("\u0046\u0069\u006c\u0074\u0065\u0072")
	if _eefc == nil {
		return nil, _gf.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_acec, _badf := _eefc.(*PdfObjectArray)
	if !_badf {
		return nil, _gf.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _afdb, _aedc := range _acec.Elements() {
		_eega, _cdfc := _aedc.(*PdfObjectName)
		if !_cdfc {
			return nil, _gf.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _aaab PdfObject
		if _dcfa != nil {
			_aaab = _dcfa
		} else {
			if len(_bbegg) > 0 {
				if _afdb >= len(_bbegg) {
					return nil, _gf.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_aaab = _bbegg[_afdb]
			}
		}
		var _bbgf *PdfObjectDictionary
		if _gaeg, _agae := _aaab.(*PdfObjectDictionary); _agae {
			_bbgf = _gaeg
		}
		_a.Log.Trace("\u004e\u0065\u0078t \u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073\u002c \u0064p\u003a \u0025v\u002c\u0020\u0064\u0050\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u0076", *_eega, _aaab, _bbgf)
		if *_eega == StreamEncodingFilterNameFlate {
			_deda, _debcc := _age(_gbff, _bbgf)
			if _debcc != nil {
				return nil, _debcc
			}
			_ege.AddEncoder(_deda)
		} else if *_eega == StreamEncodingFilterNameLZW {
			_cbdf, _fadb := _egg(_gbff, _bbgf)
			if _fadb != nil {
				return nil, _fadb
			}
			_ege.AddEncoder(_cbdf)
		} else if *_eega == StreamEncodingFilterNameASCIIHex {
			_efgd := NewASCIIHexEncoder()
			_ege.AddEncoder(_efgd)
		} else if *_eega == StreamEncodingFilterNameASCII85 {
			_ggae := NewASCII85Encoder()
			_ege.AddEncoder(_ggae)
		} else if *_eega == StreamEncodingFilterNameDCT {
			_cccb, _ggdc := _bcee(_gbff, _ege)
			if _ggdc != nil {
				return nil, _ggdc
			}
			_ege.AddEncoder(_cccb)
			_a.Log.Trace("A\u0064d\u0065\u0064\u0020\u0044\u0043\u0054\u0020\u0065n\u0063\u006f\u0064\u0065r.\u002e\u002e")
			_a.Log.Trace("\u004du\u006ct\u0069\u0020\u0065\u006e\u0063o\u0064\u0065r\u003a\u0020\u0025\u0023\u0076", _ege)
		} else if *_eega == StreamEncodingFilterNameCCITTFax {
			_ececc, _dbdd := _adcf(_gbff, _bbgf)
			if _dbdd != nil {
				return nil, _dbdd
			}
			_ege.AddEncoder(_ececc)
		} else {
			_a.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_eega)
			return nil, _gf.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _ege, nil
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_ceee *RunLengthEncoder) EncodeBytes(data []byte) ([]byte, error) {
	_ecebb := _fd.NewReader(data)
	var _cgee []byte
	var _ccae []byte
	_ecgf, _fga := _ecebb.ReadByte()
	if _fga == _fg.EOF {
		return []byte{}, nil
	} else if _fga != nil {
		return nil, _fga
	}
	_dgca := 1
	for {
		_dgcb, _cbfbf := _ecebb.ReadByte()
		if _cbfbf == _fg.EOF {
			break
		} else if _cbfbf != nil {
			return nil, _cbfbf
		}
		if _dgcb == _ecgf {
			if len(_ccae) > 0 {
				_ccae = _ccae[:len(_ccae)-1]
				if len(_ccae) > 0 {
					_cgee = append(_cgee, byte(len(_ccae)-1))
					_cgee = append(_cgee, _ccae...)
				}
				_dgca = 1
				_ccae = []byte{}
			}
			_dgca++
			if _dgca >= 127 {
				_cgee = append(_cgee, byte(257-_dgca), _ecgf)
				_dgca = 0
			}
		} else {
			if _dgca > 0 {
				if _dgca == 1 {
					_ccae = []byte{_ecgf}
				} else {
					_cgee = append(_cgee, byte(257-_dgca), _ecgf)
				}
				_dgca = 0
			}
			_ccae = append(_ccae, _dgcb)
			if len(_ccae) >= 127 {
				_cgee = append(_cgee, byte(len(_ccae)-1))
				_cgee = append(_cgee, _ccae...)
				_ccae = []byte{}
			}
		}
		_ecgf = _dgcb
	}
	if len(_ccae) > 0 {
		_cgee = append(_cgee, byte(len(_ccae)-1))
		_cgee = append(_cgee, _ccae...)
	} else if _dgca > 0 {
		_cgee = append(_cgee, byte(257-_dgca), _ecgf)
	}
	_cgee = append(_cgee, 128)
	return _cgee, nil
}

// MakeNull creates an PdfObjectNull.
func MakeNull() *PdfObjectNull { _affc := PdfObjectNull{}; return &_affc }

// String returns a string describing `stream`.
func (_cbec *PdfObjectStream) String() string {
	return _gf.Sprintf("O\u0062j\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u0025\u0064: \u0025\u0073", _cbec.ObjectNumber, _cbec.PdfObjectDictionary)
}

// IsDecimalDigit checks if the character is a part of a decimal number string.
func IsDecimalDigit(c byte) bool { return '0' <= c && c <= '9' }
func _agba(_efed, _efged, _abfd int) error {
	if _efged < 0 || _efged > _efed {
		return _c.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0061\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	if _abfd < _efged {
		return _c.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0073\u006c\u0069\u0063e\u0020i\u006ed\u0065\u0078\u0020\u0062\u0020\u003c\u0020a")
	}
	if _abfd > _efed {
		return _c.New("s\u006c\u0069\u0063\u0065\u0020\u0069n\u0064\u0065\u0078\u0020\u0062\u0020\u006f\u0075\u0074 \u006f\u0066\u0020b\u006fu\u006e\u0064\u0073")
	}
	return nil
}

// GetTrailer returns the PDFs trailer dictionary. The trailer dictionary is typically the starting point for a PDF,
// referencing other key objects that are important in the document structure.
func (_gdcaf *PdfParser) GetTrailer() *PdfObjectDictionary { return _gdcaf._dabde }

// DecodeBytes decodes byte array with ASCII85. 5 ASCII characters -> 4 raw binary bytes
func (_afga *ASCII85Encoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _deaf []byte
	_a.Log.Trace("\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0020\u0044e\u0063\u006f\u0064\u0065")
	_egbf := 0
	_bed := false
	for _egbf < len(encoded) && !_bed {
		_fdga := [5]byte{0, 0, 0, 0, 0}
		_eebe := 0
		_cgcc := 0
		_fag := 4
		for _cgcc < 5+_eebe {
			if _egbf+_cgcc == len(encoded) {
				break
			}
			_acaae := encoded[_egbf+_cgcc]
			if IsWhiteSpace(_acaae) {
				_eebe++
				_cgcc++
				continue
			} else if _acaae == '~' && _egbf+_cgcc+1 < len(encoded) && encoded[_egbf+_cgcc+1] == '>' {
				_fag = (_cgcc - _eebe) - 1
				if _fag < 0 {
					_fag = 0
				}
				_bed = true
				break
			} else if _acaae >= '!' && _acaae <= 'u' {
				_acaae -= '!'
			} else if _acaae == 'z' && _cgcc-_eebe == 0 {
				_fag = 4
				_cgcc++
				break
			} else {
				_a.Log.Error("\u0046\u0061i\u006c\u0065\u0064\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020co\u0064\u0065")
				return nil, _c.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u0064\u0065\u0020e\u006e\u0063\u006f\u0075\u006e\u0074\u0065\u0072\u0065\u0064")
			}
			_fdga[_cgcc-_eebe] = _acaae
			_cgcc++
		}
		_egbf += _cgcc
		for _aagd := _fag + 1; _aagd < 5; _aagd++ {
			_fdga[_aagd] = 84
		}
		_fbgc := uint32(_fdga[0])*85*85*85*85 + uint32(_fdga[1])*85*85*85 + uint32(_fdga[2])*85*85 + uint32(_fdga[3])*85 + uint32(_fdga[4])
		_ffdc := []byte{byte((_fbgc >> 24) & 0xff), byte((_fbgc >> 16) & 0xff), byte((_fbgc >> 8) & 0xff), byte(_fbgc & 0xff)}
		_deaf = append(_deaf, _ffdc[:_fag]...)
	}
	_a.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0065\u006e\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", encoded)
	_a.Log.Trace("A\u0053\u0043\u0049\u004985\u002c \u0064\u0065\u0063\u006f\u0064e\u0064\u003a\u0020\u0025\u0020\u0058", _deaf)
	return _deaf, nil
}

// WriteString outputs the object as it is to be written to file.
func (_deab *PdfObjectBool) WriteString() string {
	if *_deab {
		return "\u0074\u0072\u0075\u0065"
	}
	return "\u0066\u0061\u006cs\u0065"
}

// GetName returns the *PdfObjectName represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetName(obj PdfObject) (_ebdc *PdfObjectName, _cfbce bool) {
	_ebdc, _cfbce = TraceToDirectObject(obj).(*PdfObjectName)
	return _ebdc, _cfbce
}
func (_ggcf *PdfCrypt) authenticate(_fcda []byte) (bool, error) {
	_ggcf._dfe = false
	_aad := _ggcf.securityHandler()
	_gbdc, _dabb, _eceg := _aad.Authenticate(&_ggcf._aec, _fcda)
	if _eceg != nil {
		return false, _eceg
	} else if _dabb == 0 || len(_gbdc) == 0 {
		return false, nil
	}
	_ggcf._dfe = true
	_ggcf._baa = _gbdc
	return true, nil
}

// PdfIndirectObject represents the primitive PDF indirect object.
type PdfIndirectObject struct {
	PdfObjectReference
	PdfObject
}

// MakeString creates an PdfObjectString from a string.
// NOTE: PDF does not use utf-8 string encoding like Go so `s` will often not be a utf-8 encoded
// string.
func MakeString(s string) *PdfObjectString { _cfcbb := PdfObjectString{_aaca: s}; return &_cfcbb }

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
func (_edce *JBIG2Encoder) MakeStreamDict() *PdfObjectDictionary {
	_ecbf := MakeDict()
	_ecbf.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_edce.GetFilterName()))
	return _ecbf
}

// PdfObject is an interface which all primitive PDF objects must implement.
type PdfObject interface {

	// String outputs a string representation of the primitive (for debugging).
	String() string

	// WriteString outputs the PDF primitive as written to file as expected by the standard.
	// TODO(dennwc): it should return a byte slice, or accept a writer
	WriteString() string
}

// EncodeBytes returns the passed in slice of bytes.
// The purpose of the method is to satisfy the StreamEncoder interface.
func (_efggf *RawEncoder) EncodeBytes(data []byte) ([]byte, error) { return data, nil }

var _fgdd = _ce.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064+\u0029\u005c\u0073\u002b\u0028\u005b\u006e\u0066\u005d\u0029\\\u0073\u002a\u0024")

func _cbcac(_edgbe _be.Image) *JBIG2Image {
	_dccac := _edgbe.Base()
	return &JBIG2Image{Data: _dccac.Data, Width: _dccac.Width, Height: _dccac.Height, HasPadding: true}
}

// EncodeBytes encodes a bytes array and return the encoded value based on the encoder parameters.
func (_bdd *FlateEncoder) EncodeBytes(data []byte) ([]byte, error) {
	if _bdd.Predictor != 1 && _bdd.Predictor != 11 {
		_a.Log.Debug("E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0046\u006c\u0061\u0074\u0065\u0045\u006e\u0063\u006f\u0064\u0065r\u0020P\u0072\u0065\u0064\u0069c\u0074\u006fr\u0020\u003d\u0020\u0031\u002c\u0020\u0031\u0031\u0020\u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064")
		return nil, ErrUnsupportedEncodingParameters
	}
	if _bdd.Predictor == 11 {
		_gaag := _bdd.Columns
		_eeac := len(data) / _gaag
		if len(data)%_gaag != 0 {
			_a.Log.Error("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
			return nil, _c.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0072o\u0077\u0020l\u0065\u006e\u0067\u0074\u0068")
		}
		_fbbc := _fd.NewBuffer(nil)
		_fbgb := make([]byte, _gaag)
		for _dfaf := 0; _dfaf < _eeac; _dfaf++ {
			_abfb := data[_gaag*_dfaf : _gaag*(_dfaf+1)]
			_fbgb[0] = _abfb[0]
			for _egae := 1; _egae < _gaag; _egae++ {
				_fbgb[_egae] = byte(int(_abfb[_egae]-_abfb[_egae-1]) % 256)
			}
			_fbbc.WriteByte(1)
			_fbbc.Write(_fbgb)
		}
		data = _fbbc.Bytes()
	}
	var _bddb _fd.Buffer
	_adbe := _cea.NewWriter(&_bddb)
	_adbe.Write(data)
	_adbe.Close()
	return _bddb.Bytes(), nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_fdbf *ASCIIHexEncoder) MakeDecodeParams() PdfObject { return nil }

var _gadc = _ce.MustCompile("\u005e\u005b\\\u002b\u002d\u002e\u005d*\u0028\u005b0\u002d\u0039\u002e\u005d\u002b\u0029\u005b\u0065E\u005d\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

// Len returns the number of elements in the streams.
func (_gbegg *PdfObjectStreams) Len() int {
	if _gbegg == nil {
		return 0
	}
	return len(_gbegg._cegef)
}

// IsPrintable checks if a character is printable.
// Regular characters that are outside the range EXCLAMATION MARK(21h)
// (!) to TILDE (7Eh) (~) should be written using the hexadecimal notation.
func IsPrintable(c byte) bool { return 0x21 <= c && c <= 0x7E }

type offsetReader struct {
	_gfgc _fg.ReadSeeker
	_fegd int64
}

// Str returns the string value of the PdfObjectString. Defined in addition to String() function to clarify that
// this function returns the underlying string directly, whereas the String function technically could include
// debug info.
func (_bdgd *PdfObjectString) Str() string { return _bdgd._aaca }

var _ebceg = _ce.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

func (_acb *PdfCrypt) isEncrypted(_cbce PdfObject) bool {
	_, _bdce := _acb._fee[_cbce]
	if _bdce {
		_a.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0065\u006e\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	_a.Log.Trace("\u004e\u006f\u0074\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}
func (_acegf *JBIG2Encoder) encodeImage(_cffc _cg.Image) ([]byte, error) {
	const _cbef = "e\u006e\u0063\u006f\u0064\u0065\u0049\u006d\u0061\u0067\u0065"
	_addg, _bfcc := GoImageToJBIG2(_cffc, JB2ImageAutoThreshold)
	if _bfcc != nil {
		return nil, _dd.Wrap(_bfcc, _cbef, "\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069m\u0061g\u0065\u0020\u0074\u006f\u0020\u006a\u0062\u0069\u0067\u0032\u0020\u0069\u006d\u0067")
	}
	if _bfcc = _acegf.AddPageImage(_addg, &_acegf.DefaultPageSettings); _bfcc != nil {
		return nil, _dd.Wrap(_bfcc, _cbef, "")
	}
	return _acegf.Encode()
}

// WriteString outputs the object as it is to be written to file.
func (_ggda *PdfObjectStreams) WriteString() string {
	var _baecd _gd.Builder
	_baecd.WriteString(_bd.FormatInt(_ggda.ObjectNumber, 10))
	_baecd.WriteString("\u0020\u0030\u0020\u0052")
	return _baecd.String()
}

// GetNumberAsFloat returns the contents of `obj` as a float if it is an integer or float, or an
// error if it isn't.
func GetNumberAsFloat(obj PdfObject) (float64, error) {
	switch _deef := obj.(type) {
	case *PdfObjectFloat:
		return float64(*_deef), nil
	case *PdfObjectInteger:
		return float64(*_deef), nil
	case *PdfObjectReference:
		_bcggg := TraceToDirectObject(obj)
		return GetNumberAsFloat(_bcggg)
	case *PdfIndirectObject:
		return GetNumberAsFloat(_deef.PdfObject)
	}
	return 0, ErrNotANumber
}
func (_baeb *PdfCrypt) makeKey(_fcc string, _eacc, _eag uint32, _edfg []byte) ([]byte, error) {
	_fgf, _ecb := _baeb._deg[_fcc]
	if !_ecb {
		return nil, _gf.Errorf("\u0075n\u006b\u006e\u006f\u0077n\u0020\u0063\u0072\u0079\u0070t\u0020f\u0069l\u0074\u0065\u0072\u0020\u0028\u0025\u0073)", _fcc)
	}
	return _fgf.MakeKey(_eacc, _eag, _edfg)
}

// DecodeBytes decodes a slice of LZW encoded bytes and returns the result.
func (_egd *LZWEncoder) DecodeBytes(encoded []byte) ([]byte, error) {
	var _ffa _fd.Buffer
	_ccdd := _fd.NewReader(encoded)
	var _adgf _fg.ReadCloser
	if _egd.EarlyChange == 1 {
		_adgf = _d.NewReader(_ccdd, _d.MSB, 8)
	} else {
		_adgf = _bf.NewReader(_ccdd, _bf.MSB, 8)
	}
	defer _adgf.Close()
	if _, _aebdc := _ffa.ReadFrom(_adgf); _aebdc != nil {
		if _aebdc != _fg.ErrUnexpectedEOF || _ffa.Len() == 0 {
			return nil, _aebdc
		}
		_a.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u004c\u005a\u0057\u0020\u0064\u0065\u0063\u006f\u0064i\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076\u002e \u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062e \u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _aebdc)
	}
	return _ffa.Bytes(), nil
}

// WriteString outputs the object as it is to be written to file.
func (_aggb *PdfObjectName) WriteString() string {
	var _fddb _fd.Buffer
	if len(*_aggb) > 127 {
		_a.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u004e\u0061\u006d\u0065\u0020t\u006fo\u0020l\u006f\u006e\u0067\u0020\u0028\u0025\u0073)", *_aggb)
	}
	_fddb.WriteString("\u002f")
	for _edgaaf := 0; _edgaaf < len(*_aggb); _edgaaf++ {
		_acbff := (*_aggb)[_edgaaf]
		if !IsPrintable(_acbff) || _acbff == '#' || IsDelimiter(_acbff) {
			_fddb.WriteString(_gf.Sprintf("\u0023\u0025\u002e2\u0078", _acbff))
		} else {
			_fddb.WriteByte(_acbff)
		}
	}
	return _fddb.String()
}
func _age(_efgg *PdfObjectStream, _gfdfc *PdfObjectDictionary) (*FlateEncoder, error) {
	_deadc := NewFlateEncoder()
	_eafc := _efgg.PdfObjectDictionary
	if _eafc == nil {
		return _deadc, nil
	}
	_deadc._dagc = _feae(_eafc)
	if _gfdfc == nil {
		_eegg := TraceToDirectObject(_eafc.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073"))
		switch _aeff := _eegg.(type) {
		case *PdfObjectArray:
			if _aeff.Len() != 1 {
				_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0031\u0020\u0028\u0025\u0064\u0029", _aeff.Len())
				return nil, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
			}
			if _bdca, _cdee := GetDict(_aeff.Get(0)); _cdee {
				_gfdfc = _bdca
			}
		case *PdfObjectDictionary:
			_gfdfc = _aeff
		case *PdfObjectNull, nil:
		default:
			_a.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _eegg)
			return nil, _gf.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
		}
	}
	if _gfdfc == nil {
		return _deadc, nil
	}
	_a.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _gfdfc.String())
	_fadg := _gfdfc.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _fadg == nil {
		_a.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_gcc, _bad := _fadg.(*PdfObjectInteger)
		if !_bad {
			_a.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _fadg)
			return nil, _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_deadc.Predictor = int(*_gcc)
	}
	_fadg = _gfdfc.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _fadg != nil {
		_baece, _aag := _fadg.(*PdfObjectInteger)
		if !_aag {
			_a.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _gf.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_deadc.BitsPerComponent = int(*_baece)
	}
	if _deadc.Predictor > 1 {
		_deadc.Columns = 1
		_fadg = _gfdfc.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _fadg != nil {
			_bec, _abea := _fadg.(*PdfObjectInteger)
			if !_abea {
				return nil, _gf.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_deadc.Columns = int(*_bec)
		}
		_deadc.Colors = 1
		_fadg = _gfdfc.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _fadg != nil {
			_ddee, _decd := _fadg.(*PdfObjectInteger)
			if !_decd {
				return nil, _gf.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_deadc.Colors = int(*_ddee)
		}
	}
	return _deadc, nil
}

type cryptFilters map[string]_efd.Filter

// MakeStreamDict makes a new instance of an encoding dictionary for a stream object.
// Has the Filter set and the DecodeParms.
func (_adf *LZWEncoder) MakeStreamDict() *PdfObjectDictionary {
	_cfdf := MakeDict()
	_cfdf.Set("\u0046\u0069\u006c\u0074\u0065\u0072", MakeName(_adf.GetFilterName()))
	_faf := _adf.MakeDecodeParams()
	if _faf != nil {
		_cfdf.Set("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073", _faf)
	}
	_cfdf.Set("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065", MakeInteger(int64(_adf.EarlyChange)))
	return _cfdf
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

func (_gdcad *limitedReadSeeker) getError(_dcgg int64) error {
	switch {
	case _dcgg < 0:
		return _gf.Errorf("\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0065\u0067\u0061\u0074\u0069\u0076e\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0064", _dcgg)
	case _dcgg > _gdcad._bcgd:
		return _gf.Errorf("u\u006e\u0065\u0078\u0070ec\u0074e\u0064\u0020\u006f\u0066\u0066s\u0065\u0074\u003a\u0020\u0025\u0064", _dcgg)
	}
	return nil
}

// GetFilterName returns the name of the encoding filter.
func (_gaf *RunLengthEncoder) GetFilterName() string { return StreamEncodingFilterNameRunLength }
func _gbbga(_dggge PdfObject, _fgbd int) PdfObject {
	if _fgbd > _ddfg {
		_a.Log.Error("\u0054\u0072ac\u0065\u0020\u0064e\u0070\u0074\u0068\u0020lev\u0065l \u0062\u0065\u0079\u006f\u006e\u0064\u0020%d\u0020\u002d\u0020\u0065\u0072\u0072\u006fr\u0021", _ddfg)
		return MakeNull()
	}
	switch _ccfc := _dggge.(type) {
	case *PdfIndirectObject:
		_dggge = _gbbga((*_ccfc).PdfObject, _fgbd+1)
	case *PdfObjectArray:
		for _gfeb, _abcf := range (*_ccfc)._fabc {
			(*_ccfc)._fabc[_gfeb] = _gbbga(_abcf, _fgbd+1)
		}
	case *PdfObjectDictionary:
		for _cbcce, _dfbfa := range (*_ccfc)._adbeb {
			(*_ccfc)._adbeb[_cbcce] = _gbbga(_dfbfa, _fgbd+1)
		}
		_gg.Slice((*_ccfc)._caeg, func(_aada, _egef int) bool { return (*_ccfc)._caeg[_aada] < (*_ccfc)._caeg[_egef] })
	}
	return _dggge
}

var _dgbg = _ce.MustCompile("\u005b\\\u0072\u005c\u006e\u005d\u005c\u0073\u002a\u0028\u0078\u0072\u0065f\u0029\u005c\u0073\u002a\u005b\u005c\u0072\u005c\u006e\u005d")

// IsEncrypted checks if the document is encrypted. A bool flag is returned indicating the result.
// First time when called, will check if the Encrypt dictionary is accessible through the trailer dictionary.
// If encrypted, prepares a crypt datastructure which can be used to authenticate and decrypt the document.
// On failure, an error is returned.
func (_bcdad *PdfParser) IsEncrypted() (bool, error) {
	if _bcdad._abae != nil {
		return true, nil
	} else if _bcdad._dabde == nil {
		return false, nil
	}
	_a.Log.Trace("\u0043\u0068\u0065c\u006b\u0069\u006e\u0067 \u0065\u006e\u0063\u0072\u0079\u0070\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0021")
	_ccaed := _bcdad._dabde.Get("\u0045n\u0063\u0072\u0079\u0070\u0074")
	if _ccaed == nil {
		return false, nil
	}
	_a.Log.Trace("\u0049\u0073\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u0021")
	var (
		_dbec *PdfObjectDictionary
	)
	switch _fcce := _ccaed.(type) {
	case *PdfObjectDictionary:
		_dbec = _fcce
	case *PdfObjectReference:
		_a.Log.Trace("\u0030\u003a\u0020\u004c\u006f\u006f\u006b\u0020\u0075\u0070\u0020\u0072e\u0066\u0020\u0025\u0071", _fcce)
		_dbad, _eddd := _bcdad.LookupByReference(*_fcce)
		_a.Log.Trace("\u0031\u003a\u0020%\u0071", _dbad)
		if _eddd != nil {
			return false, _eddd
		}
		_dbba, _cgcgf := _dbad.(*PdfIndirectObject)
		if !_cgcgf {
			_a.Log.Debug("E\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006eo\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0064\u0069\u0072ec\u0074\u0020\u006fb\u006ae\u0063\u0074")
			return false, _c.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_abcb, _cgcgf := _dbba.PdfObject.(*PdfObjectDictionary)
		_bcdad._fafd = _dbba
		_a.Log.Trace("\u0032\u003a\u0020%\u0071", _abcb)
		if !_cgcgf {
			return false, _c.New("\u0074\u0072a\u0069\u006c\u0065\u0072 \u0045\u006ec\u0072\u0079\u0070\u0074\u0020\u006f\u0062\u006ae\u0063\u0074\u0020\u006e\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		}
		_dbec = _abcb
	case *PdfObjectNull:
		_a.Log.Debug("\u0045\u006e\u0063\u0072\u0079\u0070\u0074 \u0069\u0073\u0020a\u0020\u006e\u0075l\u006c\u0020o\u0062\u006a\u0065\u0063\u0074\u002e \u0046il\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0065\u0064\u002e")
		return false, nil
	default:
		return false, _gf.Errorf("u\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065: \u0025\u0054", _fcce)
	}
	_cdga, _fgdc := PdfCryptNewDecrypt(_bcdad, _dbec, _bcdad._dabde)
	if _fgdc != nil {
		return false, _fgdc
	}
	for _, _eabc := range []string{"\u0045n\u0063\u0072\u0079\u0070\u0074"} {
		_gdgc := _bcdad._dabde.Get(PdfObjectName(_eabc))
		if _gdgc == nil {
			continue
		}
		switch _dgbafa := _gdgc.(type) {
		case *PdfObjectReference:
			_cdga._cdda[int(_dgbafa.ObjectNumber)] = struct{}{}
		case *PdfIndirectObject:
			_cdga._ecee[_dgbafa] = true
			_cdga._cdda[int(_dgbafa.ObjectNumber)] = struct{}{}
		}
	}
	_bcdad._abae = _cdga
	_a.Log.Trace("\u0043\u0072\u0079\u0070\u0074\u0065\u0072\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0025\u0062", _cdga)
	return true, nil
}

// String returns a string describing `ind`.
func (_fdfbd *PdfIndirectObject) String() string {
	return _gf.Sprintf("\u0049\u004f\u0062\u006a\u0065\u0063\u0074\u003a\u0025\u0064", (*_fdfbd).ObjectNumber)
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_acg *LZWEncoder) MakeDecodeParams() PdfObject {
	if _acg.Predictor > 1 {
		_fgcc := MakeDict()
		_fgcc.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", MakeInteger(int64(_acg.Predictor)))
		if _acg.BitsPerComponent != 8 {
			_fgcc.Set("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074", MakeInteger(int64(_acg.BitsPerComponent)))
		}
		if _acg.Columns != 1 {
			_fgcc.Set("\u0043o\u006c\u0075\u006d\u006e\u0073", MakeInteger(int64(_acg.Columns)))
		}
		if _acg.Colors != 1 {
			_fgcc.Set("\u0043\u006f\u006c\u006f\u0072\u0073", MakeInteger(int64(_acg.Colors)))
		}
		return _fgcc
	}
	return nil
}

// MakeArrayFromFloats creates an PdfObjectArray from a slice of float64s, where each array element is an
// PdfObjectFloat.
func MakeArrayFromFloats(vals []float64) *PdfObjectArray {
	_edgd := MakeArray()
	for _, _ebfe := range vals {
		_edgd.Append(MakeFloat(_ebfe))
	}
	return _edgd
}

// PdfObjectBool represents the primitive PDF boolean object.
type PdfObjectBool bool

// GetString returns the *PdfObjectString represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetString(obj PdfObject) (_fcgf *PdfObjectString, _bfegf bool) {
	_fcgf, _bfegf = TraceToDirectObject(obj).(*PdfObjectString)
	return _fcgf, _bfegf
}

// MakeStream creates an PdfObjectStream with specified contents and encoding. If encoding is nil, then raw encoding
// will be used (i.e. no encoding applied).
func MakeStream(contents []byte, encoder StreamEncoder) (*PdfObjectStream, error) {
	_efcc := &PdfObjectStream{}
	if encoder == nil {
		encoder = NewRawEncoder()
	}
	_efcc.PdfObjectDictionary = encoder.MakeStreamDict()
	_bcedb, _bdfe := encoder.EncodeBytes(contents)
	if _bdfe != nil {
		return nil, _bdfe
	}
	_efcc.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_bcedb))))
	_efcc.Stream = _bcedb
	return _efcc, nil
}

// Update updates multiple keys and returns the dictionary back so can be used in a chained fashion.
func (_gafc *PdfObjectDictionary) Update(objmap map[string]PdfObject) *PdfObjectDictionary {
	_gafc._ccdg.Lock()
	defer _gafc._ccdg.Unlock()
	for _defb, _bcfga := range objmap {
		_gafc.setWithLock(PdfObjectName(_defb), _bcfga, false)
	}
	return _gafc
}

// UpdateParams updates the parameter values of the encoder.
func (_agfa *DCTEncoder) UpdateParams(params *PdfObjectDictionary) {
	_baedg, _dbb := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _dbb == nil {
		_agfa.ColorComponents = int(_baedg)
	}
	_dafb, _dbb := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _dbb == nil {
		_agfa.BitsPerComponent = int(_dafb)
	}
	_acda, _dbb := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _dbb == nil {
		_agfa.Width = int(_acda)
	}
	_fdgg, _dbb := GetNumberAsInt64(params.Get("\u0048\u0065\u0069\u0067\u0068\u0074"))
	if _dbb == nil {
		_agfa.Height = int(_fdgg)
	}
	_dcga, _dbb := GetNumberAsInt64(params.Get("\u0051u\u0061\u006c\u0069\u0074\u0079"))
	if _dbb == nil {
		_agfa.Quality = int(_dcga)
	}
	_befd, _abb := GetArray(params.Get("\u0044\u0065\u0063\u006f\u0064\u0065"))
	if _abb {
		_agfa.Decode, _dbb = _befd.ToFloat64Array()
		if _dbb != nil {
			_a.Log.Error("F\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u006f\u006ev\u0065\u0072\u0074\u0069\u006e\u0067\u0020de\u0063\u006f\u0064\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006eto\u0020\u0061r\u0072\u0061\u0079\u0073\u003a\u0020\u0025\u0076", _dbb)
		}
	}
}

// WriteString outputs the object as it is to be written to file.
func (_gcgcf *PdfObjectNull) WriteString() string { return "\u006e\u0075\u006c\u006c" }

// DecodeImages decodes the page images from the jbig2 'encoded' data input.
// The jbig2 document may contain multiple pages, thus the function can return multiple
// images. The images order corresponds to the page number.
func (_aaacc *JBIG2Encoder) DecodeImages(encoded []byte) ([]_cg.Image, error) {
	const _ddgeb = "\u004aB\u0049\u0047\u0032\u0045n\u0063\u006f\u0064\u0065\u0072.\u0044e\u0063o\u0064\u0065\u0049\u006d\u0061\u0067\u0065s"
	_cab, _aadb := _eg.Decode(encoded, _eg.Parameters{}, _aaacc.Globals.ToDocumentGlobals())
	if _aadb != nil {
		return nil, _dd.Wrap(_aadb, _ddgeb, "")
	}
	_dccc, _aadb := _cab.PageNumber()
	if _aadb != nil {
		return nil, _dd.Wrap(_aadb, _ddgeb, "")
	}
	_eded := []_cg.Image{}
	var _dgga _cg.Image
	for _dgba := 1; _dgba <= _dccc; _dgba++ {
		_dgga, _aadb = _cab.DecodePageImage(_dgba)
		if _aadb != nil {
			return nil, _dd.Wrapf(_aadb, _ddgeb, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _dgba)
		}
		_eded = append(_eded, _dgga)
	}
	return _eded, nil
}

// GetFilterName returns the name of the encoding filter.
func (_dgg *RawEncoder) GetFilterName() string { return StreamEncodingFilterNameRaw }

// NewParserFromString is used for testing purposes.
func NewParserFromString(txt string) *PdfParser {
	_gfda := _fd.NewReader([]byte(txt))
	_ggfbb := &PdfParser{ObjCache: objectCache{}, _dfcdg: _gfda, _ffbg: _bfc.NewReader(_gfda), _fbaaa: int64(len(txt)), _ffed: map[int64]bool{}, _beaf: make(map[*PdfParser]*PdfParser)}
	_ggfbb._bbdf.ObjectMap = make(map[int]XrefObject)
	return _ggfbb
}

// MakeHexString creates an PdfObjectString from a string intended for output as a hexadecimal string.
func MakeHexString(s string) *PdfObjectString {
	_fdbfa := PdfObjectString{_aaca: s, _gead: true}
	return &_fdbfa
}
func (_dbga *PdfParser) readComment() (string, error) {
	var _dcbe _fd.Buffer
	_, _fbf := _dbga.skipSpaces()
	if _fbf != nil {
		return _dcbe.String(), _fbf
	}
	_baeba := true
	for {
		_eecg, _fccg := _dbga._ffbg.Peek(1)
		if _fccg != nil {
			_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _fccg.Error())
			return _dcbe.String(), _fccg
		}
		if _baeba && _eecg[0] != '%' {
			return _dcbe.String(), _c.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_baeba = false
		if (_eecg[0] != '\r') && (_eecg[0] != '\n') {
			_gbbc, _ := _dbga._ffbg.ReadByte()
			_dcbe.WriteByte(_gbbc)
		} else {
			break
		}
	}
	return _dcbe.String(), nil
}

// MakeDecodeParams makes a new instance of an encoding dictionary based on
// the current encoder settings.
func (_febb *ASCII85Encoder) MakeDecodeParams() PdfObject { return nil }

// PdfObjectName represents the primitive PDF name object.
type PdfObjectName string

// DecodeStream implements ASCII85 stream decoding.
func (_ebade *ASCII85Encoder) DecodeStream(streamObj *PdfObjectStream) ([]byte, error) {
	return _ebade.DecodeBytes(streamObj.Stream)
}

var _cac = _ce.MustCompile("\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u0028\u005c\u0064\u002b)\u005c\u0073\u002a\u0024")
var _dbef = _ce.MustCompile("\u0025P\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

func (_dadf *PdfCrypt) isDecrypted(_bagd PdfObject) bool {
	_, _adbd := _dadf._ecee[_bagd]
	if _adbd {
		_a.Log.Trace("\u0041\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0064\u0065\u0063\u0072y\u0070\u0074\u0065\u0064")
		return true
	}
	switch _dcb := _bagd.(type) {
	case *PdfObjectStream:
		if _dadf._aec.R != 5 {
			if _gacb, _dga := _dcb.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _dga && *_gacb == "\u0058\u0052\u0065\u0066" {
				return true
			}
		}
	case *PdfIndirectObject:
		if _, _adbd = _dadf._cdda[int(_dcb.ObjectNumber)]; _adbd {
			return true
		}
		switch _dccd := _dcb.PdfObject.(type) {
		case *PdfObjectDictionary:
			_eae := true
			for _, _fae := range _aeec {
				if _dccd.Get(_fae) == nil {
					_eae = false
					break
				}
			}
			if _eae {
				return true
			}
		}
	}
	_a.Log.Trace("\u004e\u006f\u0074\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u0065d\u0020\u0079\u0065\u0074")
	return false
}

// Append appends PdfObject(s) to the streams.
func (_dadgcc *PdfObjectStreams) Append(objects ...PdfObject) {
	if _dadgcc == nil {
		_a.Log.Debug("\u0057\u0061\u0072\u006e\u0020-\u0020\u0041\u0074\u0074\u0065\u006d\u0070\u0074\u0020\u0074\u006f\u0020\u0061p\u0070\u0065\u006e\u0064\u0020\u0074\u006f\u0020\u0061\u0020\u006e\u0069\u006c\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073")
		return
	}
	_dadgcc._cegef = append(_dadgcc._cegef, objects...)
}

// GetXrefOffset returns the offset of the xref table.
func (_ddgbg *PdfParser) GetXrefOffset() int64 { return _ddgbg._caed }

// GetNumberAsInt64 returns the contents of `obj` as an int64 if it is an integer or float, or an
// error if it isn't. This is for cases where expecting an integer, but some implementations
// actually store the number in a floating point format.
func GetNumberAsInt64(obj PdfObject) (int64, error) {
	switch _fadgge := obj.(type) {
	case *PdfObjectFloat:
		_a.Log.Debug("\u004e\u0075m\u0062\u0065\u0072\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u0073\u0020\u0069\u006e\u0074e\u0067\u0065\u0072\u0020\u0077\u0061s\u0020\u0073\u0074\u006f\u0072\u0065\u0064\u0020\u0061\u0073\u0020\u0066\u006c\u006fa\u0074\u0020(\u0074\u0079\u0070\u0065 \u0063\u0061\u0073\u0074\u0069n\u0067\u0020\u0075\u0073\u0065\u0064\u0029")
		return int64(*_fadgge), nil
	case *PdfObjectInteger:
		return int64(*_fadgge), nil
	case *PdfObjectReference:
		_fbbfg := TraceToDirectObject(obj)
		return GetNumberAsInt64(_fbbfg)
	case *PdfIndirectObject:
		return GetNumberAsInt64(_fadgge.PdfObject)
	}
	return 0, ErrNotANumber
}

// PdfObjectNull represents the primitive PDF null object.
type PdfObjectNull struct{}

func (_fbcd *offsetReader) Seek(offset int64, whence int) (int64, error) {
	if whence == _fg.SeekStart {
		offset += _fbcd._fegd
	}
	_bagb, _gaaga := _fbcd._gfgc.Seek(offset, whence)
	if _gaaga != nil {
		return _bagb, _gaaga
	}
	if whence == _fg.SeekCurrent {
		_bagb -= _fbcd._fegd
	}
	if _bagb < 0 {
		return 0, _c.New("\u0063\u006f\u0072\u0065\u002eo\u0066\u0066\u0073\u0065\u0074\u0052\u0065\u0061\u0064\u0065\u0072\u002e\u0053e\u0065\u006b\u003a\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	return _bagb, nil
}

// EncodeBytes encodes slice of bytes into JBIG2 encoding format.
// The input 'data' must be an image. In order to Decode it a user is responsible to
// load the codec ('png', 'jpg').
// Returns jbig2 single page encoded document byte slice. The encoder uses DefaultPageSettings
// to encode given image.
func (_agad *JBIG2Encoder) EncodeBytes(data []byte) ([]byte, error) {
	const _gagbd = "\u004aB\u0049\u0047\u0032\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002eE\u006e\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065\u0073"
	if _agad.ColorComponents != 1 || _agad.BitsPerComponent != 1 {
		return nil, _dd.Errorf(_gagbd, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u002e\u0020\u004a\u0042\u0049G\u0032\u0020E\u006e\u0063o\u0064\u0065\u0072\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020bi\u006e\u0061\u0072\u0079\u0020\u0069\u006d\u0061\u0067e\u0073\u0020\u0064\u0061\u0074\u0061")
	}
	var (
		_efdgf *_ef.Bitmap
		_bga   error
	)
	_adce := (_agad.Width * _agad.Height) == len(data)
	if _adce {
		_efdgf, _bga = _ef.NewWithUnpaddedData(_agad.Width, _agad.Height, data)
	} else {
		_efdgf, _bga = _ef.NewWithData(_agad.Width, _agad.Height, data)
	}
	if _bga != nil {
		return nil, _bga
	}
	_eaeb := _agad.DefaultPageSettings
	if _bga = _eaeb.Validate(); _bga != nil {
		return nil, _dd.Wrap(_bga, _gagbd, "")
	}
	if _agad._feg == nil {
		_agad._feg = _de.InitEncodeDocument(_eaeb.FileMode)
	}
	switch _eaeb.Compression {
	case JB2Generic:
		if _bga = _agad._feg.AddGenericPage(_efdgf, _eaeb.DuplicatedLinesRemoval); _bga != nil {
			return nil, _dd.Wrap(_bga, _gagbd, "")
		}
	case JB2SymbolCorrelation:
		return nil, _dd.Error(_gagbd, "s\u0079\u006d\u0062\u006f\u006c\u0020\u0063\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e e\u006e\u0063\u006f\u0064i\u006e\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006dpl\u0065\u006de\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	case JB2SymbolRankHaus:
		return nil, _dd.Error(_gagbd, "\u0073y\u006d\u0062o\u006c\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073 \u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0020\u006e\u006f\u0074\u0020\u0069\u006d\u0070\u006c\u0065m\u0065\u006e\u0074\u0065\u0064\u0020\u0079\u0065\u0074")
	default:
		return nil, _dd.Error(_gagbd, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020c\u006f\u006d\u0070\u0072\u0065\u0073\u0073i\u006f\u006e")
	}
	return _agad.Encode()
}

// EncodeBytes encodes the image data using either Group3 or Group4 CCITT facsimile (fax) encoding.
// `data` is expected to be 1 color component, 1 bit per component. It is also valid to provide 8 BPC, 1 CC image like
// a standard go image Gray data.
func (_afef *CCITTFaxEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _dfeb _be.Gray
	switch len(data) {
	case _afef.Rows * _afef.Columns:
		_efge, _fdaf := _be.NewImage(_afef.Columns, _afef.Rows, 8, 1, data, nil, nil)
		if _fdaf != nil {
			return nil, _fdaf
		}
		_dfeb = _efge.(_be.Gray)
	case (_afef.Columns * _afef.Rows) + 7>>3:
		_eabg, _dceg := _be.NewImage(_afef.Columns, _afef.Rows, 1, 1, data, nil, nil)
		if _dceg != nil {
			return nil, _dceg
		}
		_gcgd := _eabg.(*_be.Monochrome)
		if _dceg = _gcgd.AddPadding(); _dceg != nil {
			return nil, _dceg
		}
		_dfeb = _gcgd
	default:
		if len(data) < _be.BytesPerLine(_afef.Columns, 1, 1)*_afef.Rows {
			return nil, _c.New("p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020i\u006e\u0070\u0075t\u0020d\u0061\u0074\u0061")
		}
		_gfbg, _fbd := _be.NewImage(_afef.Columns, _afef.Rows, 1, 1, data, nil, nil)
		if _fbd != nil {
			return nil, _fbd
		}
		_gccc := _gfbg.(*_be.Monochrome)
		_dfeb = _gccc
	}
	_cdgc := make([][]byte, _afef.Rows)
	for _cdege := 0; _cdege < _afef.Rows; _cdege++ {
		_badg := make([]byte, _afef.Columns)
		for _accfb := 0; _accfb < _afef.Columns; _accfb++ {
			_adag := _dfeb.GrayAt(_accfb, _cdege)
			_badg[_accfb] = _adag.Y >> 7
		}
		_cdgc[_cdege] = _badg
	}
	_gbf := &_gfa.Encoder{K: _afef.K, Columns: _afef.Columns, EndOfLine: _afef.EndOfLine, EndOfBlock: _afef.EndOfBlock, BlackIs1: _afef.BlackIs1, DamagedRowsBeforeError: _afef.DamagedRowsBeforeError, Rows: _afef.Rows, EncodedByteAlign: _afef.EncodedByteAlign}
	return _gbf.Encode(_cdgc), nil
}

// String returns a string representation of `name`.
func (_gace *PdfObjectName) String() string { return string(*_gace) }

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

// FlateEncoder represents Flate encoding.
type FlateEncoder struct {
	Predictor        int
	BitsPerComponent int

	// For predictors
	Columns int
	Rows    int
	Colors  int
	_dagc   *_be.ImageBase
}

// AddEncoder adds the passed in encoder to the underlying encoder slice.
func (_cgaf *MultiEncoder) AddEncoder(encoder StreamEncoder) {
	_cgaf._bdfb = append(_cgaf._bdfb, encoder)
}

// GetBoolVal returns the bool value within a *PdObjectBool represented by an PdfObject interface directly or indirectly.
// If the PdfObject does not represent a bool value, a default value of false is returned (found = false also).
func GetBoolVal(obj PdfObject) (_aacf bool, _cgfg bool) {
	_cfgd, _cgfg := TraceToDirectObject(obj).(*PdfObjectBool)
	if _cgfg {
		return bool(*_cfgd), true
	}
	return false, false
}
func _dbaeec(_edgaf _fg.ReadSeeker, _egfd int64) (*offsetReader, error) {
	_cdbc := &offsetReader{_gfgc: _edgaf, _fegd: _egfd}
	_, _cbcc := _cdbc.Seek(0, _fg.SeekStart)
	return _cdbc, _cbcc
}

// Encrypt an object with specified key. For numbered objects,
// the key argument is not used and a new one is generated based
// on the object and generation number.
// Traverses through all the subobjects (recursive).
//
// Does not look up references..  That should be done prior to calling.
func (_fccc *PdfCrypt) Encrypt(obj PdfObject, parentObjNum, parentGenNum int64) error {
	if _fccc.isEncrypted(obj) {
		return nil
	}
	switch _baed := obj.(type) {
	case *PdfIndirectObject:
		_fccc._fee[_baed] = true
		_a.Log.Trace("\u0045\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006e\u0067 \u0069\u006e\u0064\u0069\u0072\u0065\u0063t\u0020\u0025\u0064\u0020\u0025\u0064\u0020\u006f\u0062\u006a\u0021", _baed.ObjectNumber, _baed.GenerationNumber)
		_cfg := _baed.ObjectNumber
		_dae := _baed.GenerationNumber
		_bdbf := _fccc.Encrypt(_baed.PdfObject, _cfg, _dae)
		if _bdbf != nil {
			return _bdbf
		}
		return nil
	case *PdfObjectStream:
		_fccc._fee[_baed] = true
		_adda := _baed.PdfObjectDictionary
		if _ebd, _efaa := _adda.Get("\u0054\u0079\u0070\u0065").(*PdfObjectName); _efaa && *_ebd == "\u0058\u0052\u0065\u0066" {
			return nil
		}
		_aga := _baed.ObjectNumber
		_fbaa := _baed.GenerationNumber
		_a.Log.Trace("\u0045n\u0063\u0072\u0079\u0070t\u0069\u006e\u0067\u0020\u0073t\u0072e\u0061m\u0020\u0025\u0064\u0020\u0025\u0064\u0020!", _aga, _fbaa)
		_bgf := _edg
		if _fccc._ceae.V >= 4 {
			_bgf = _fccc._abdf
			_a.Log.Trace("\u0074\u0068\u0069\u0073.s\u0074\u0072\u0065\u0061\u006d\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u003d\u0020%\u0073", _fccc._abdf)
			if _efcf, _acc := _adda.Get("\u0046\u0069\u006c\u0074\u0065\u0072").(*PdfObjectArray); _acc {
				if _bdgg, _eadc := GetName(_efcf.Get(0)); _eadc {
					if *_bdgg == "\u0043\u0072\u0079p\u0074" {
						_bgf = "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"
						if _gacdg, _fgcd := _adda.Get("D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073").(*PdfObjectDictionary); _fgcd {
							if _edef, _daf := _gacdg.Get("\u004e\u0061\u006d\u0065").(*PdfObjectName); _daf {
								if _, _gdg := _fccc._deg[string(*_edef)]; _gdg {
									_a.Log.Trace("\u0055\u0073\u0069\u006eg \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020%\u0073", *_edef)
									_bgf = string(*_edef)
								}
							}
						}
					}
				}
			}
			_a.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _bgf)
			if _bgf == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
		}
		_agg := _fccc.Encrypt(_baed.PdfObjectDictionary, _aga, _fbaa)
		if _agg != nil {
			return _agg
		}
		_dgd, _agg := _fccc.makeKey(_bgf, uint32(_aga), uint32(_fbaa), _fccc._baa)
		if _agg != nil {
			return _agg
		}
		_baed.Stream, _agg = _fccc.encryptBytes(_baed.Stream, _bgf, _dgd)
		if _agg != nil {
			return _agg
		}
		_adda.Set("\u004c\u0065\u006e\u0067\u0074\u0068", MakeInteger(int64(len(_baed.Stream))))
		return nil
	case *PdfObjectString:
		_a.Log.Trace("\u0045n\u0063r\u0079\u0070\u0074\u0069\u006eg\u0020\u0073t\u0072\u0069\u006e\u0067\u0021")
		_acf := _edg
		if _fccc._ceae.V >= 4 {
			_a.Log.Trace("\u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0066i\u006c\u0074\u0065\u0072", _fccc._dfb)
			if _fccc._dfb == "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" {
				return nil
			}
			_acf = _fccc._dfb
		}
		_eba, _aef := _fccc.makeKey(_acf, uint32(parentObjNum), uint32(parentGenNum), _fccc._baa)
		if _aef != nil {
			return _aef
		}
		_dbg := _baed.Str()
		_adc := make([]byte, len(_dbg))
		for _dfea := 0; _dfea < len(_dbg); _dfea++ {
			_adc[_dfea] = _dbg[_dfea]
		}
		_a.Log.Trace("\u0045n\u0063\u0072\u0079\u0070\u0074\u0020\u0073\u0074\u0072\u0069\u006eg\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025\u0020\u0078", _adc, _adc)
		_adc, _aef = _fccc.encryptBytes(_adc, _acf, _eba)
		if _aef != nil {
			return _aef
		}
		_baed._aaca = string(_adc)
		return nil
	case *PdfObjectArray:
		for _, _deee := range _baed.Elements() {
			_egac := _fccc.Encrypt(_deee, parentObjNum, parentGenNum)
			if _egac != nil {
				return _egac
			}
		}
		return nil
	case *PdfObjectDictionary:
		_af := false
		if _ceef := _baed.Get("\u0054\u0079\u0070\u0065"); _ceef != nil {
			_fbaf, _efe := _ceef.(*PdfObjectName)
			if _efe && *_fbaf == "\u0053\u0069\u0067" {
				_af = true
			}
		}
		for _, _dcg := range _baed.Keys() {
			_aaf := _baed.Get(_dcg)
			if _af && string(_dcg) == "\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073" {
				continue
			}
			if string(_dcg) != "\u0050\u0061\u0072\u0065\u006e\u0074" && string(_dcg) != "\u0050\u0072\u0065\u0076" && string(_dcg) != "\u004c\u0061\u0073\u0074" {
				_ceb := _fccc.Encrypt(_aaf, parentObjNum, parentGenNum)
				if _ceb != nil {
					return _ceb
				}
			}
		}
		return nil
	}
	return nil
}

// PdfParser parses a PDF file and provides access to the object structure of the PDF.
type PdfParser struct {
	_eebec   Version
	_dfcdg   _fg.ReadSeeker
	_ffbg    *_bfc.Reader
	_fbaaa   int64
	_bbdf    XrefTable
	_caed    int64
	_fddf    *xrefType
	_fcba    objectStreams
	_dabde   *PdfObjectDictionary
	_abae    *PdfCrypt
	_fafd    *PdfIndirectObject
	_cged    bool
	ObjCache objectCache
	_eeaa    map[int]bool
	_ffed    map[int64]bool
	_gadge   ParserMetadata
	_ecgd    bool
	_fbaae   []int64
	_efce    int
	_dgcba   bool
	_baac    int64
	_beaf    map[*PdfParser]*PdfParser
	_gfdfd   []*PdfParser
}

// GetFilterArray returns the names of the underlying encoding filters in an array that
// can be used as /Filter entry.
func (_fgeef *MultiEncoder) GetFilterArray() *PdfObjectArray {
	_efab := make([]PdfObject, len(_fgeef._bdfb))
	for _bdde, _daa := range _fgeef._bdfb {
		_efab[_bdde] = MakeName(_daa.GetFilterName())
	}
	return MakeArray(_efab...)
}

// GetCrypter returns the PdfCrypt instance which has information about the PDFs encryption.
func (_ggag *PdfParser) GetCrypter() *PdfCrypt { return _ggag._abae }

// TraceToDirectObject traces a PdfObject to a direct object.  For example direct objects contained
// in indirect objects (can be double referenced even).
func TraceToDirectObject(obj PdfObject) PdfObject {
	if _dfbcaf, _ggeb := obj.(*PdfObjectReference); _ggeb {
		obj = _dfbcaf.Resolve()
	}
	_fccff, _efeg := obj.(*PdfIndirectObject)
	_egaec := 0
	for _efeg {
		obj = _fccff.PdfObject
		_fccff, _efeg = GetIndirect(obj)
		_egaec++
		if _egaec > _ddfg {
			_a.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0061\u0063\u0065\u0020\u0064\u0065p\u0074\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u0062\u0065\u0079\u006fn\u0064\u0020\u0025\u0064\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0067oi\u006e\u0067\u0020\u0064\u0065\u0065\u0070\u0065\u0072\u0021", _ddfg)
			return nil
		}
	}
	return obj
}

// Inspect analyzes the document object structure. Returns a map of object types (by name) with the instance count
// as value.
func (_ggdde *PdfParser) Inspect() (map[string]int, error) { return _ggdde.inspect() }

// Get returns the i-th element of the array or nil if out of bounds (by index).
func (_cedd *PdfObjectArray) Get(i int) PdfObject {
	if _cedd == nil || i >= len(_cedd._fabc) || i < 0 {
		return nil
	}
	return _cedd._fabc[i]
}

var _aeec = []PdfObjectName{"\u0056", "\u0052", "\u004f", "\u0055", "\u0050"}

// String returns a string describing `null`.
func (_fcagb *PdfObjectNull) String() string { return "\u006e\u0075\u006c\u006c" }

// EncodeBytes ASCII encodes the passed in slice of bytes.
func (_bbaa *ASCIIHexEncoder) EncodeBytes(data []byte) ([]byte, error) {
	var _gagb _fd.Buffer
	for _, _becc := range data {
		_gagb.WriteString(_gf.Sprintf("\u0025\u002e\u0032X\u0020", _becc))
	}
	_gagb.WriteByte('>')
	return _gagb.Bytes(), nil
}

// HeaderCommentBytes gets the header comment bytes.
func (_bbc ParserMetadata) HeaderCommentBytes() [4]byte { return _bbc._egf }

// GetStringVal returns the string value represented by the PdfObject directly or indirectly if
// contained within an indirect object. On type mismatch the found bool flag returned is false and
// an empty string is returned.
func GetStringVal(obj PdfObject) (_aaad string, _fdbga bool) {
	_gbge, _fdbga := TraceToDirectObject(obj).(*PdfObjectString)
	if _fdbga {
		return _gbge.Str(), true
	}
	return
}

// String returns a descriptive information string about the encryption method used.
func (_cga *PdfCrypt) String() string {
	if _cga == nil {
		return ""
	}
	_ega := _cga._ceae.Filter + "\u0020\u002d\u0020"
	if _cga._ceae.V == 0 {
		_ega += "\u0055\u006e\u0064\u006fcu\u006d\u0065\u006e\u0074\u0065\u0064\u0020\u0061\u006c\u0067\u006f\u0072\u0069\u0074h\u006d"
	} else if _cga._ceae.V == 1 {
		_ega += "\u0052\u0043\u0034:\u0020\u0034\u0030\u0020\u0062\u0069\u0074\u0073"
	} else if _cga._ceae.V == 2 {
		_ega += _gf.Sprintf("\u0052\u0043\u0034:\u0020\u0025\u0064\u0020\u0062\u0069\u0074\u0073", _cga._ceae.Length)
	} else if _cga._ceae.V == 3 {
		_ega += "U\u006e\u0070\u0075\u0062li\u0073h\u0065\u0064\u0020\u0061\u006cg\u006f\u0072\u0069\u0074\u0068\u006d"
	} else if _cga._ceae.V >= 4 {
		_ega += _gf.Sprintf("\u0053\u0074r\u0065\u0061\u006d\u0020f\u0069\u006ct\u0065\u0072\u003a\u0020\u0025\u0073\u0020\u002d \u0053\u0074\u0072\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074\u0065r\u003a\u0020\u0025\u0073", _cga._abdf, _cga._dfb)
		_ega += "\u003b\u0020C\u0072\u0079\u0070t\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0073\u003a"
		for _ggfb, _ge := range _cga._deg {
			_ega += _gf.Sprintf("\u0020\u002d\u0020\u0025\u0073\u003a\u0020\u0025\u0073 \u0028\u0025\u0064\u0029", _ggfb, _ge.Name(), _ge.KeyLength())
		}
	}
	_eeb := _cga.GetAccessPermissions()
	_ega += _gf.Sprintf("\u0020\u002d\u0020\u0025\u0023\u0076", _eeb)
	return _ega
}
func (_aace *PdfParser) parseXref() (*PdfObjectDictionary, error) {
	_aace.skipSpaces()
	const _fagb = 20
	_bggb, _ := _aace._ffbg.Peek(_fagb)
	for _bdggb := 0; _bdggb < 2; _bdggb++ {
		if _aace._caed == 0 {
			_aace._caed = _aace.GetFileOffset()
		}
		if _ebceg.Match(_bggb) {
			_a.Log.Trace("\u0078\u0072e\u0066\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002e\u0020\u0050\u0072\u006f\u0062\u0061\u0062\u006c\u0079\u0020\u0078\u0072\u0065\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			_a.Log.Debug("\u0073t\u0061r\u0074\u0069\u006e\u0067\u0020w\u0069\u0074h\u0020\u0022\u0025\u0073\u0022", string(_bggb))
			return _aace.parseXrefStream(nil)
		}
		if _cgcd.Match(_bggb) {
			_a.Log.Trace("\u0053\u0074\u0061\u006ed\u0061\u0072\u0064\u0020\u0078\u0072\u0065\u0066\u0020\u0073e\u0063t\u0069\u006f\u006e\u0020\u0074\u0061\u0062l\u0065\u0021")
			return _aace.parseXrefTable()
		}
		_cdgf := _aace.GetFileOffset()
		if _aace._caed == 0 {
			_aace._caed = _cdgf
		}
		_aace.SetFileOffset(_cdgf - _fagb)
		defer _aace.SetFileOffset(_cdgf)
		_ccfb, _ := _aace._ffbg.Peek(_fagb)
		_bggb = append(_ccfb, _bggb...)
	}
	_a.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006e\u0067\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u0078\u0072\u0065f\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u006fr\u0020\u0073\u0074\u0072\u0065\u0061\u006d.\u0020\u0052\u0065\u0070\u0061i\u0072\u0020\u0061\u0074\u0074e\u006d\u0070\u0074\u0065\u0064\u003a\u0020\u004c\u006f\u006f\u006b\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0065\u0061\u0072\u006c\u0069\u0065\u0073\u0074\u0020x\u0072\u0065\u0066\u0020\u0066\u0072\u006f\u006d\u0020\u0062\u006f\u0074to\u006d\u002e")
	if _addf := _aace.repairSeekXrefMarker(); _addf != nil {
		_a.Log.Debug("\u0052e\u0070a\u0069\u0072\u0020\u0066\u0061i\u006c\u0065d\u0020\u002d\u0020\u0025\u0076", _addf)
		return nil, _addf
	}
	return _aace.parseXrefTable()
}

// HasOddLengthHexStrings checks if the document has odd length hexadecimal strings.
func (_ebag ParserMetadata) HasOddLengthHexStrings() bool { return _ebag._gcg }
func (_aafb *PdfParser) parseName() (PdfObjectName, error) {
	var _eccbe _fd.Buffer
	_fcca := false
	for {
		_efgcb, _adadb := _aafb._ffbg.Peek(1)
		if _adadb == _fg.EOF {
			break
		}
		if _adadb != nil {
			return PdfObjectName(_eccbe.String()), _adadb
		}
		if !_fcca {
			if _efgcb[0] == '/' {
				_fcca = true
				_aafb._ffbg.ReadByte()
			} else if _efgcb[0] == '%' {
				_aafb.readComment()
				_aafb.skipSpaces()
			} else {
				_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _efgcb, _efgcb)
				return PdfObjectName(_eccbe.String()), _gf.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _efgcb[0])
			}
		} else {
			if IsWhiteSpace(_efgcb[0]) {
				break
			} else if (_efgcb[0] == '/') || (_efgcb[0] == '[') || (_efgcb[0] == '(') || (_efgcb[0] == ']') || (_efgcb[0] == '<') || (_efgcb[0] == '>') {
				break
			} else if _efgcb[0] == '#' {
				_gbgd, _gcggb := _aafb._ffbg.Peek(3)
				if _gcggb != nil {
					return PdfObjectName(_eccbe.String()), _gcggb
				}
				_adge, _gcggb := _bdc.DecodeString(string(_gbgd[1:3]))
				if _gcggb != nil {
					_a.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0066o\u006c\u006co\u0077\u0069\u006e\u0067 \u0027\u0023\u0027\u002c \u0063\u006f\u006e\u0074\u0069n\u0075\u0069\u006e\u0067\u0020\u0075\u0073i\u006e\u0067\u0020\u006c\u0069t\u0065\u0072\u0061\u006c\u0020\u002d\u0020\u004f\u0075t\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074")
					_eccbe.WriteByte('#')
					_aafb._ffbg.Discard(1)
					continue
				}
				_aafb._ffbg.Discard(3)
				_eccbe.Write(_adge)
			} else {
				_aefef, _ := _aafb._ffbg.ReadByte()
				_eccbe.WriteByte(_aefef)
			}
		}
	}
	return PdfObjectName(_eccbe.String()), nil
}

// Get returns the PdfObject corresponding to the specified key.
// Returns a nil value if the key is not set.
func (_aebe *PdfObjectDictionary) Get(key PdfObjectName) PdfObject {
	_aebe._ccdg.Lock()
	defer _aebe._ccdg.Unlock()
	_eeeac, _bcbe := _aebe._adbeb[key]
	if !_bcbe {
		return nil
	}
	return _eeeac
}
func (_dfaad *PdfObjectInteger) String() string { return _gf.Sprintf("\u0025\u0064", *_dfaad) }

// PdfObjectStreams represents the primitive PDF object streams.
// 7.5.7 Object Streams (page 45).
type PdfObjectStreams struct {
	PdfObjectReference
	_cegef []PdfObject
}

// GetAccessPermissions returns the PDF access permissions as an AccessPermissions object.
func (_ebe *PdfCrypt) GetAccessPermissions() _dc.Permissions { return _ebe._aec.P }

// WriteString outputs the object as it is to be written to file.
func (_abcd *PdfObjectFloat) WriteString() string {
	return _bd.FormatFloat(float64(*_abcd), 'f', -1, 64)
}

// Clear resets the array to an empty state.
func (_edcd *PdfObjectArray) Clear() { _edcd._fabc = []PdfObject{} }

// GetFilterName returns the name of the encoding filter.
func (_gda *FlateEncoder) GetFilterName() string { return StreamEncodingFilterNameFlate }

// HeaderPosition gets the file header position.
func (_ccdc ParserMetadata) HeaderPosition() int { return _ccdc._bef }

// NewMultiEncoder returns a new instance of MultiEncoder.
func NewMultiEncoder() *MultiEncoder {
	_aba := MultiEncoder{}
	_aba._bdfb = []StreamEncoder{}
	return &_aba
}

var _fafa = _ce.MustCompile("\u005e\\\u0073\u002a\u005b\u002d]\u002a\u0028\u005c\u0064\u002b)\u005cs\u002b(\u005c\u0064\u002b\u0029\u005c\u0073\u002bR")

// MakeDict creates and returns an empty PdfObjectDictionary.
func MakeDict() *PdfObjectDictionary {
	_geefb := &PdfObjectDictionary{}
	_geefb._adbeb = map[PdfObjectName]PdfObject{}
	_geefb._caeg = []PdfObjectName{}
	_geefb._ccdg = &_g.Mutex{}
	return _geefb
}
func (_fbea *PdfParser) parseString() (*PdfObjectString, error) {
	_fbea._ffbg.ReadByte()
	var _fede _fd.Buffer
	_bbfda := 1
	for {
		_gbbg, _aaee := _fbea._ffbg.Peek(1)
		if _aaee != nil {
			return MakeString(_fede.String()), _aaee
		}
		if _gbbg[0] == '\\' {
			_fbea._ffbg.ReadByte()
			_addb, _dfff := _fbea._ffbg.ReadByte()
			if _dfff != nil {
				return MakeString(_fede.String()), _dfff
			}
			if IsOctalDigit(_addb) {
				_dccf, _eedg := _fbea._ffbg.Peek(2)
				if _eedg != nil {
					return MakeString(_fede.String()), _eedg
				}
				var _aagfe []byte
				_aagfe = append(_aagfe, _addb)
				for _, _dgdd := range _dccf {
					if IsOctalDigit(_dgdd) {
						_aagfe = append(_aagfe, _dgdd)
					} else {
						break
					}
				}
				_fbea._ffbg.Discard(len(_aagfe) - 1)
				_a.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _aagfe)
				_dgab, _eedg := _bd.ParseUint(string(_aagfe), 8, 32)
				if _eedg != nil {
					return MakeString(_fede.String()), _eedg
				}
				_fede.WriteByte(byte(_dgab))
				continue
			}
			switch _addb {
			case 'n':
				_fede.WriteRune('\n')
			case 'r':
				_fede.WriteRune('\r')
			case 't':
				_fede.WriteRune('\t')
			case 'b':
				_fede.WriteRune('\b')
			case 'f':
				_fede.WriteRune('\f')
			case '(':
				_fede.WriteRune('(')
			case ')':
				_fede.WriteRune(')')
			case '\\':
				_fede.WriteRune('\\')
			}
			continue
		} else if _gbbg[0] == '(' {
			_bbfda++
		} else if _gbbg[0] == ')' {
			_bbfda--
			if _bbfda == 0 {
				_fbea._ffbg.ReadByte()
				break
			}
		}
		_ffde, _ := _fbea._ffbg.ReadByte()
		_fede.WriteByte(_ffde)
	}
	return MakeString(_fede.String()), nil
}

// JPXEncoder implements JPX encoder/decoder (dummy, for now)
// FIXME: implement
type JPXEncoder struct{}

// ResolveReferencesDeep recursively traverses through object `o`, looking up and replacing
// references with indirect objects.
// Optionally a map of already deep-resolved objects can be provided via `traversed`. The `traversed` map
// is updated while traversing the objects to avoid traversing same objects multiple times.
func ResolveReferencesDeep(o PdfObject, traversed map[PdfObject]struct{}) error {
	if traversed == nil {
		traversed = map[PdfObject]struct{}{}
	}
	return _agaf(o, 0, traversed)
}
func _ff(_cbc XrefTable) {
	_a.Log.Debug("\u003dX\u003d\u0058\u003d\u0058\u003d")
	_a.Log.Debug("X\u0072\u0065\u0066\u0020\u0074\u0061\u0062\u006c\u0065\u003a")
	_beg := 0
	for _, _aed := range _cbc.ObjectMap {
		_a.Log.Debug("i\u002b\u0031\u003a\u0020\u0025\u0064 \u0028\u006f\u0062\u006a\u0020\u006eu\u006d\u003a\u0020\u0025\u0064\u0020\u0067e\u006e\u003a\u0020\u0025\u0064\u0029\u0020\u002d\u003e\u0020%\u0064", _beg+1, _aed.ObjectNumber, _aed.Generation, _aed.Offset)
		_beg++
	}
}

// GetDict returns the *PdfObjectDictionary represented by the PdfObject directly or indirectly within an indirect
// object. On type mismatch the found bool flag is false and a nil pointer is returned.
func GetDict(obj PdfObject) (_gfdef *PdfObjectDictionary, _cbde bool) {
	_gfdef, _cbde = TraceToDirectObject(obj).(*PdfObjectDictionary)
	return _gfdef, _cbde
}

// MakeIndirectObject creates an PdfIndirectObject with a specified direct object PdfObject.
func MakeIndirectObject(obj PdfObject) *PdfIndirectObject {
	_gfdfe := &PdfIndirectObject{}
	_gfdfe.PdfObject = obj
	return _gfdfe
}

// UpdateParams updates the parameter values of the encoder.
func (_deeec *LZWEncoder) UpdateParams(params *PdfObjectDictionary) {
	_aca, _cdcd := GetNumberAsInt64(params.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr"))
	if _cdcd == nil {
		_deeec.Predictor = int(_aca)
	}
	_cdfb, _cdcd := GetNumberAsInt64(params.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"))
	if _cdcd == nil {
		_deeec.BitsPerComponent = int(_cdfb)
	}
	_efdg, _cdcd := GetNumberAsInt64(params.Get("\u0057\u0069\u0064t\u0068"))
	if _cdcd == nil {
		_deeec.Columns = int(_efdg)
	}
	_edga, _cdcd := GetNumberAsInt64(params.Get("\u0043o\u006co\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073"))
	if _cdcd == nil {
		_deeec.Colors = int(_edga)
	}
	_edeg, _cdcd := GetNumberAsInt64(params.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065"))
	if _cdcd == nil {
		_deeec.EarlyChange = int(_edeg)
	}
}

// IsNullObject returns true if `obj` is a PdfObjectNull.
func IsNullObject(obj PdfObject) bool {
	_, _dgad := TraceToDirectObject(obj).(*PdfObjectNull)
	return _dgad
}
