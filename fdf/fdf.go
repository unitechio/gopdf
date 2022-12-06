package fdf

import (
	_e "bufio"
	_ae "bytes"
	_a "encoding/hex"
	_d "errors"
	_de "fmt"
	_db "io"
	_af "os"
	_f "regexp"
	_fe "sort"
	_g "strconv"
	_dg "strings"

	_c "bitbucket.org/shenghui0779/gopdf/common"
	_ff "bitbucket.org/shenghui0779/gopdf/core"
)

func (_aef *fdfParser) setFileOffset(_aec int64) {
	_aef._bfb.Seek(_aec, _db.SeekStart)
	_aef._gec = _e.NewReader(_aef._bfb)
}

var _fgcd = _f.MustCompile("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

func (_ccf *fdfParser) skipSpaces() (int, error) {
	_ceb := 0
	for {
		_gd, _feb := _ccf._gec.ReadByte()
		if _feb != nil {
			return 0, _feb
		}
		if _ff.IsWhiteSpace(_gd) {
			_ceb++
		} else {
			_ccf._gec.UnreadByte()
			break
		}
	}
	return _ceb, nil
}
func (_gdf *fdfParser) parseFdfVersion() (int, int, error) {
	_gdf._bfb.Seek(0, _db.SeekStart)
	_cgc := 20
	_fdb := make([]byte, _cgc)
	_gdf._bfb.Read(_fdb)
	_aaa := _fgcd.FindStringSubmatch(string(_fdb))
	if len(_aaa) < 3 {
		_cgg, _caag, _fceb := _gdf.seekFdfVersionTopDown()
		if _fceb != nil {
			_c.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _fceb
		}
		return _cgg, _caag, nil
	}
	_eef, _fag := _g.Atoi(_aaa[1])
	if _fag != nil {
		return 0, 0, _fag
	}
	_fefb, _fag := _g.Atoi(_aaa[2])
	if _fag != nil {
		return 0, 0, _fag
	}
	_c.Log.Debug("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _eef, _fefb)
	return _eef, _fefb, nil
}
func (_ecg *fdfParser) readComment() (string, error) {
	var _ee _ae.Buffer
	_, _be := _ecg.skipSpaces()
	if _be != nil {
		return _ee.String(), _be
	}
	_gfc := true
	for {
		_efd, _ade := _ecg._gec.Peek(1)
		if _ade != nil {
			_c.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _ade.Error())
			return _ee.String(), _ade
		}
		if _gfc && _efd[0] != '%' {
			return _ee.String(), _d.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_gfc = false
		if (_efd[0] != '\r') && (_efd[0] != '\n') {
			_bd, _ := _ecg._gec.ReadByte()
			_ee.WriteByte(_bd)
		} else {
			break
		}
	}
	return _ee.String(), nil
}

type fdfParser struct {
	_fbb int
	_abd int
	_cg  map[int64]_ff.PdfObject
	_bfb _db.ReadSeeker
	_gec *_e.Reader
	_fc  int64
	_ac  *_ff.PdfObjectDictionary
}

var _bf = _f.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

func (_fbc *fdfParser) parseObject() (_ff.PdfObject, error) {
	_c.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_fbc.skipSpaces()
	for {
		_ebe, _fda := _fbc._gec.Peek(2)
		if _fda != nil {
			return nil, _fda
		}
		_c.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_ebe))
		if _ebe[0] == '/' {
			_bdd, _fcb := _fbc.parseName()
			_c.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _bdd)
			return &_bdd, _fcb
		} else if _ebe[0] == '(' {
			_c.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			return _fbc.parseString()
		} else if _ebe[0] == '[' {
			_c.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			return _fbc.parseArray()
		} else if (_ebe[0] == '<') && (_ebe[1] == '<') {
			_c.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			return _fbc.parseDict()
		} else if _ebe[0] == '<' {
			_c.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			return _fbc.parseHexString()
		} else if _ebe[0] == '%' {
			_fbc.readComment()
			_fbc.skipSpaces()
		} else {
			_c.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_ebe, _ = _fbc._gec.Peek(15)
			_gbd := string(_ebe)
			_c.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _gbd)
			if (len(_gbd) > 3) && (_gbd[:4] == "\u006e\u0075\u006c\u006c") {
				_dfe, _ecgbe := _fbc.parseNull()
				return &_dfe, _ecgbe
			} else if (len(_gbd) > 4) && (_gbd[:5] == "\u0066\u0061\u006cs\u0065") {
				_eeg, _bbf := _fbc.parseBool()
				return &_eeg, _bbf
			} else if (len(_gbd) > 3) && (_gbd[:4] == "\u0074\u0072\u0075\u0065") {
				_ccc, _dga := _fbc.parseBool()
				return &_ccc, _dga
			}
			_aag := _fb.FindStringSubmatch(_gbd)
			if len(_aag) > 1 {
				_ebe, _ = _fbc._gec.ReadBytes('R')
				_c.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_ebe[:]))
				_ged, _egg := _aab(string(_ebe))
				return &_ged, _egg
			}
			_eee := _eaf.FindStringSubmatch(_gbd)
			if len(_eee) > 1 {
				_c.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				return _fbc.parseNumber()
			}
			_eee = _fab.FindStringSubmatch(_gbd)
			if len(_eee) > 1 {
				_c.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_c.Log.Trace("\u0025\u0020\u0073", _eee)
				return _fbc.parseNumber()
			}
			_c.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _gbd)
			return nil, _d.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}
func (_ecd *fdfParser) readTextLine() (string, error) {
	var _ffa _ae.Buffer
	for {
		_eg, _ege := _ecd._gec.Peek(1)
		if _ege != nil {
			_c.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _ege.Error())
			return _ffa.String(), _ege
		}
		if (_eg[0] != '\r') && (_eg[0] != '\n') {
			_ffd, _ := _ecd._gec.ReadByte()
			_ffa.WriteByte(_ffd)
		} else {
			break
		}
	}
	return _ffa.String(), nil
}
func _daeb(_edbg _db.ReadSeeker) (*fdfParser, error) {
	_gab := &fdfParser{}
	_gab._bfb = _edbg
	_gab._cg = map[int64]_ff.PdfObject{}
	_bcf, _aff, _dgf := _gab.parseFdfVersion()
	if _dgf != nil {
		_c.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _dgf)
		return nil, _dgf
	}
	_gab._fbb = _bcf
	_gab._abd = _aff
	_dgf = _gab.parse()
	return _gab, _dgf
}

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath(fdfPath string) (*Data, error) {
	_fg, _gc := _af.Open(fdfPath)
	if _gc != nil {
		return nil, _gc
	}
	defer _fg.Close()
	return Load(_fg)
}
func (_gdc *fdfParser) seekFdfVersionTopDown() (int, int, error) {
	_gdc._bfb.Seek(0, _db.SeekStart)
	_gdc._gec = _e.NewReader(_gdc._bfb)
	_efa := 20
	_add := make([]byte, _efa)
	for {
		_bce, _fae := _gdc._gec.ReadByte()
		if _fae != nil {
			if _fae == _db.EOF {
				break
			} else {
				return 0, 0, _fae
			}
		}
		if _ff.IsDecimalDigit(_bce) && _add[_efa-1] == '.' && _ff.IsDecimalDigit(_add[_efa-2]) && _add[_efa-3] == '-' && _add[_efa-4] == 'F' && _add[_efa-5] == 'D' && _add[_efa-6] == 'P' {
			_bgd := int(_add[_efa-2] - '0')
			_dff := int(_bce - '0')
			return _bgd, _dff, nil
		}
		_add = append(_add[1:_efa], _bce)
	}
	return 0, 0, _d.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// Load loads FDF form data from `r`.
func Load(r _db.ReadSeeker) (*Data, error) {
	_ga, _cc := _daeb(r)
	if _cc != nil {
		return nil, _cc
	}
	_cd, _cc := _ga.Root()
	if _cc != nil {
		return nil, _cc
	}
	_ag, _ca := _ff.GetArray(_cd.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_ca {
		return nil, _d.New("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	return &Data{_ge: _ag, _bc: _cd}, nil
}
func (_gaa *fdfParser) parseArray() (*_ff.PdfObjectArray, error) {
	_eadc := _ff.MakeArray()
	_gaa._gec.ReadByte()
	for {
		_gaa.skipSpaces()
		_gca, _beab := _gaa._gec.Peek(1)
		if _beab != nil {
			return _eadc, _beab
		}
		if _gca[0] == ']' {
			_gaa._gec.ReadByte()
			break
		}
		_ebb, _beab := _gaa.parseObject()
		if _beab != nil {
			return _eadc, _beab
		}
		_eadc.Append(_ebb)
	}
	return _eadc, nil
}
func (_dc *fdfParser) skipComments() error {
	if _, _fce := _dc.skipSpaces(); _fce != nil {
		return _fce
	}
	_gdb := true
	for {
		_dge, _adf := _dc._gec.Peek(1)
		if _adf != nil {
			_c.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _adf.Error())
			return _adf
		}
		if _gdb && _dge[0] != '%' {
			return nil
		}
		_gdb = false
		if (_dge[0] != '\r') && (_dge[0] != '\n') {
			_dc._gec.ReadByte()
		} else {
			break
		}
	}
	return _dc.skipComments()
}
func (_dgc *fdfParser) parseName() (_ff.PdfObjectName, error) {
	var _adg _ae.Buffer
	_agfd := false
	for {
		_abe, _ecf := _dgc._gec.Peek(1)
		if _ecf == _db.EOF {
			break
		}
		if _ecf != nil {
			return _ff.PdfObjectName(_adg.String()), _ecf
		}
		if !_agfd {
			if _abe[0] == '/' {
				_agfd = true
				_dgc._gec.ReadByte()
			} else if _abe[0] == '%' {
				_dgc.readComment()
				_dgc.skipSpaces()
			} else {
				_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _abe, _abe)
				return _ff.PdfObjectName(_adg.String()), _de.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _abe[0])
			}
		} else {
			if _ff.IsWhiteSpace(_abe[0]) {
				break
			} else if (_abe[0] == '/') || (_abe[0] == '[') || (_abe[0] == '(') || (_abe[0] == ']') || (_abe[0] == '<') || (_abe[0] == '>') {
				break
			} else if _abe[0] == '#' {
				_adgb, _gga := _dgc._gec.Peek(3)
				if _gga != nil {
					return _ff.PdfObjectName(_adg.String()), _gga
				}
				_dgc._gec.Discard(3)
				_bg, _gga := _a.DecodeString(string(_adgb[1:3]))
				if _gga != nil {
					return _ff.PdfObjectName(_adg.String()), _gga
				}
				_adg.Write(_bg)
			} else {
				_bfbd, _ := _dgc._gec.ReadByte()
				_adg.WriteByte(_bfbd)
			}
		}
	}
	return _ff.PdfObjectName(_adg.String()), nil
}
func (_ggae *fdfParser) parseNull() (_ff.PdfObjectNull, error) {
	_, _bga := _ggae._gec.Discard(4)
	return _ff.PdfObjectNull{}, _bga
}
func (_fec *fdfParser) parseString() (*_ff.PdfObjectString, error) {
	_fec._gec.ReadByte()
	var _eab _ae.Buffer
	_abc := 1
	for {
		_afg, _caa := _fec._gec.Peek(1)
		if _caa != nil {
			return _ff.MakeString(_eab.String()), _caa
		}
		if _afg[0] == '\\' {
			_fec._gec.ReadByte()
			_cb, _aa := _fec._gec.ReadByte()
			if _aa != nil {
				return _ff.MakeString(_eab.String()), _aa
			}
			if _ff.IsOctalDigit(_cb) {
				_aad, _ba := _fec._gec.Peek(2)
				if _ba != nil {
					return _ff.MakeString(_eab.String()), _ba
				}
				var _ada []byte
				_ada = append(_ada, _cb)
				for _, _fbbd := range _aad {
					if _ff.IsOctalDigit(_fbbd) {
						_ada = append(_ada, _fbbd)
					} else {
						break
					}
				}
				_fec._gec.Discard(len(_ada) - 1)
				_c.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _ada)
				_gdbg, _ba := _g.ParseUint(string(_ada), 8, 32)
				if _ba != nil {
					return _ff.MakeString(_eab.String()), _ba
				}
				_eab.WriteByte(byte(_gdbg))
				continue
			}
			switch _cb {
			case 'n':
				_eab.WriteRune('\n')
			case 'r':
				_eab.WriteRune('\r')
			case 't':
				_eab.WriteRune('\t')
			case 'b':
				_eab.WriteRune('\b')
			case 'f':
				_eab.WriteRune('\f')
			case '(':
				_eab.WriteRune('(')
			case ')':
				_eab.WriteRune(')')
			case '\\':
				_eab.WriteRune('\\')
			}
			continue
		} else if _afg[0] == '(' {
			_abc++
		} else if _afg[0] == ')' {
			_abc--
			if _abc == 0 {
				_fec._gec.ReadByte()
				break
			}
		}
		_bec, _ := _fec._gec.ReadByte()
		_eab.WriteByte(_bec)
	}
	return _ff.MakeString(_eab.String()), nil
}

var _eaf = _f.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func (_fac *fdfParser) parseDict() (*_ff.PdfObjectDictionary, error) {
	_c.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_bbc := _ff.MakeDict()
	_dfa, _ := _fac._gec.ReadByte()
	if _dfa != '<' {
		return nil, _d.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_dfa, _ = _fac._gec.ReadByte()
	if _dfa != '<' {
		return nil, _d.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_fac.skipSpaces()
		_fac.skipComments()
		_cag, _efc := _fac._gec.Peek(2)
		if _efc != nil {
			return nil, _efc
		}
		_c.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_cag), string(_cag))
		if (_cag[0] == '>') && (_cag[1] == '>') {
			_c.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_fac._gec.ReadByte()
			_fac._gec.ReadByte()
			break
		}
		_c.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_def, _efc := _fac.parseName()
		_c.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _def)
		if _efc != nil {
			_c.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _efc)
			return nil, _efc
		}
		if len(_def) > 4 && _def[len(_def)-4:] == "\u006e\u0075\u006c\u006c" {
			_bae := _def[0 : len(_def)-4]
			_c.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _def)
			_c.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _bae)
			_fac.skipSpaces()
			_bdf, _ := _fac._gec.Peek(1)
			if _bdf[0] == '/' {
				_bbc.Set(_bae, _ff.MakeNull())
				continue
			}
		}
		_fac.skipSpaces()
		_gaab, _efc := _fac.parseObject()
		if _efc != nil {
			return nil, _efc
		}
		_bbc.Set(_def, _gaab)
		_c.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _def, _gaab.String())
	}
	_c.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _bbc, nil
}

var _fb = _f.MustCompile("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052")

func _aab(_afb string) (_ff.PdfObjectReference, error) {
	_cf := _ff.PdfObjectReference{}
	_gfg := _fb.FindStringSubmatch(_afb)
	if len(_gfg) < 3 {
		_c.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _cf, _d.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_cdf, _fdg := _g.Atoi(_gfg[1])
	if _fdg != nil {
		_c.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030", _gfg[1])
		return _cf, nil
	}
	_cf.ObjectNumber = int64(_cdf)
	_cda, _fdg := _g.Atoi(_gfg[2])
	if _fdg != nil {
		_c.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030", _gfg[2])
		return _cf, nil
	}
	_cf.GenerationNumber = int64(_cda)
	return _cf, nil
}

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data) FieldDictionaries() (map[string]*_ff.PdfObjectDictionary, error) {
	_fgc := map[string]*_ff.PdfObjectDictionary{}
	for _ad := 0; _ad < fdf._ge.Len(); _ad++ {
		_fed, _gef := _ff.GetDict(fdf._ge.Get(_ad))
		if _gef {
			_ec, _ := _ff.GetString(_fed.Get("\u0054"))
			if _ec != nil {
				_fgc[_ec.Str()] = _fed
			}
		}
	}
	return _fgc, nil
}
func (_baa *fdfParser) seekToEOFMarker(_agdf int64) error {
	_dcfc := int64(0)
	_abcf := int64(1000)
	for _dcfc < _agdf {
		if _agdf <= (_abcf + _dcfc) {
			_abcf = _agdf - _dcfc
		}
		_, _geg := _baa._bfb.Seek(-_dcfc-_abcf, _db.SeekEnd)
		if _geg != nil {
			return _geg
		}
		_gce := make([]byte, _abcf)
		_baa._bfb.Read(_gce)
		_c.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_gce))
		_dee := _bb.FindAllStringIndex(string(_gce), -1)
		if _dee != nil {
			_cbf := _dee[len(_dee)-1]
			_c.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _dee)
			_baa._bfb.Seek(-_dcfc-_abcf+int64(_cbf[0]), _db.SeekEnd)
			return nil
		}
		_c.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_dcfc += _abcf
	}
	_c.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _d.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

// Data represents forms data format (FDF) file data.
type Data struct {
	_bc *_ff.PdfObjectDictionary
	_ge *_ff.PdfObjectArray
}

func (_ce *fdfParser) getFileOffset() int64 {
	_afc, _ := _ce._bfb.Seek(0, _db.SeekCurrent)
	_afc -= int64(_ce._gec.Buffered())
	return _afc
}

var _bb = _f.MustCompile("\u0025\u0025\u0045O\u0046")

func (_aae *fdfParser) trace(_ecgf _ff.PdfObject) _ff.PdfObject {
	switch _ecc := _ecgf.(type) {
	case *_ff.PdfObjectReference:
		_agdg, _fbbf := _aae._cg[_ecc.ObjectNumber].(*_ff.PdfIndirectObject)
		if _fbbf {
			return _agdg.PdfObject
		}
		_c.Log.Debug("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		return nil
	case *_ff.PdfIndirectObject:
		return _ecc.PdfObject
	}
	return _ecgf
}
func (_aba *fdfParser) parseHexString() (*_ff.PdfObjectString, error) {
	_aba._gec.ReadByte()
	var _fba _ae.Buffer
	for {
		_bea, _ddd := _aba._gec.Peek(1)
		if _ddd != nil {
			return _ff.MakeHexString(""), _ddd
		}
		if _bea[0] == '>' {
			_aba._gec.ReadByte()
			break
		}
		_ed, _ := _aba._gec.ReadByte()
		if !_ff.IsWhiteSpace(_ed) {
			_fba.WriteByte(_ed)
		}
	}
	if _fba.Len()%2 == 1 {
		_fba.WriteRune('0')
	}
	_fad, _dea := _a.DecodeString(_fba.String())
	if _dea != nil {
		_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067", _fba.String())
		return _ff.MakeHexString(""), nil
	}
	return _ff.MakeHexString(string(_fad)), nil
}
func (_geff *fdfParser) parseBool() (_ff.PdfObjectBool, error) {
	_fbd, _fef := _geff._gec.Peek(4)
	if _fef != nil {
		return _ff.PdfObjectBool(false), _fef
	}
	if (len(_fbd) >= 4) && (string(_fbd[:4]) == "\u0074\u0072\u0075\u0065") {
		_geff._gec.Discard(4)
		return _ff.PdfObjectBool(true), nil
	}
	_fbd, _fef = _geff._gec.Peek(5)
	if _fef != nil {
		return _ff.PdfObjectBool(false), _fef
	}
	if (len(_fbd) >= 5) && (string(_fbd[:5]) == "\u0066\u0061\u006cs\u0065") {
		_geff._gec.Discard(5)
		return _ff.PdfObjectBool(false), nil
	}
	return _ff.PdfObjectBool(false), _d.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

var _fab = _f.MustCompile("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029")

func _fca(_cbfg string) (*fdfParser, error) {
	_cac := fdfParser{}
	_bfce := []byte(_cbfg)
	_ecgg := _ae.NewReader(_bfce)
	_cac._bfb = _ecgg
	_cac._cg = map[int64]_ff.PdfObject{}
	_afa := _e.NewReader(_ecgg)
	_cac._gec = _afa
	_cac._fc = int64(len(_cbfg))
	return &_cac, _cac.parse()
}
func (_bdda *fdfParser) parseIndirectObject() (_ff.PdfObject, error) {
	_ffda := _ff.PdfIndirectObject{}
	_c.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_dfae, _gda := _bdda._gec.Peek(20)
	if _gda != nil {
		_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
		return &_ffda, _gda
	}
	_c.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_dfae))
	_cggg := _bf.FindStringSubmatchIndex(string(_dfae))
	if len(_cggg) < 6 {
		_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_dfae))
		return &_ffda, _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_bdda._gec.Discard(_cggg[0])
	_c.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _cggg)
	_bgc := _cggg[1] - _cggg[0]
	_cdfc := make([]byte, _bgc)
	_, _gda = _bdda.readAtLeast(_cdfc, _bgc)
	if _gda != nil {
		_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _gda)
		return nil, _gda
	}
	_c.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _cdfc)
	_cgd := _bf.FindStringSubmatch(string(_cdfc))
	if len(_cgd) < 3 {
		_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_cdfc))
		return &_ffda, _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_aee, _ := _g.Atoi(_cgd[1])
	_cfb, _ := _g.Atoi(_cgd[2])
	_ffda.ObjectNumber = int64(_aee)
	_ffda.GenerationNumber = int64(_cfb)
	for {
		_baed, _fee := _bdda._gec.Peek(2)
		if _fee != nil {
			return &_ffda, _fee
		}
		_c.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_baed), string(_baed))
		if _ff.IsWhiteSpace(_baed[0]) {
			_bdda.skipSpaces()
		} else if _baed[0] == '%' {
			_bdda.skipComments()
		} else if (_baed[0] == '<') && (_baed[1] == '<') {
			_c.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_ffda.PdfObject, _fee = _bdda.parseDict()
			_c.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _fee)
			if _fee != nil {
				return &_ffda, _fee
			}
			_c.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_baed[0] == '/') || (_baed[0] == '(') || (_baed[0] == '[') || (_baed[0] == '<') {
			_ffda.PdfObject, _fee = _bdda.parseObject()
			if _fee != nil {
				return &_ffda, _fee
			}
			_c.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else {
			if _baed[0] == 'e' {
				_ddg, _feee := _bdda.readTextLine()
				if _feee != nil {
					return nil, _feee
				}
				if len(_ddg) >= 6 && _ddg[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _baed[0] == 's' {
				_baed, _ = _bdda._gec.Peek(10)
				if string(_baed[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_ded := 6
					if len(_baed) > 6 {
						if _ff.IsWhiteSpace(_baed[_ded]) && _baed[_ded] != '\r' && _baed[_ded] != '\n' {
							_c.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_ded++
						}
						if _baed[_ded] == '\r' {
							_ded++
							if _baed[_ded] == '\n' {
								_ded++
							}
						} else if _baed[_ded] == '\n' {
							_ded++
						}
					}
					_bdda._gec.Discard(_ded)
					_facc, _acca := _ffda.PdfObject.(*_ff.PdfObjectDictionary)
					if !_acca {
						return nil, _d.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_c.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _facc)
					_dbg, _acb := _facc.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*_ff.PdfObjectInteger)
					if !_acb {
						return nil, _d.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_daf := *_dbg
					if _daf < 0 {
						return nil, _d.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					if int64(_daf) > _bdda._fc {
						_c.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _d.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_bbcg := make([]byte, _daf)
					_, _fee = _bdda.readAtLeast(_bbcg, int(_daf))
					if _fee != nil {
						_c.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_bbcg), _bbcg)
						_c.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fee)
						return nil, _fee
					}
					_gefa := _ff.PdfObjectStream{}
					_gefa.Stream = _bbcg
					_gefa.PdfObjectDictionary = _ffda.PdfObject.(*_ff.PdfObjectDictionary)
					_gefa.ObjectNumber = _ffda.ObjectNumber
					_gefa.GenerationNumber = _ffda.GenerationNumber
					_bdda.skipSpaces()
					_bdda._gec.Discard(9)
					_bdda.skipSpaces()
					return &_gefa, nil
				}
			}
			_ffda.PdfObject, _fee = _bdda.parseObject()
			return &_ffda, _fee
		}
	}
	_c.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_ffda, nil
}

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data) FieldValues() (map[string]_ff.PdfObject, error) {
	_bcb, _dec := fdf.FieldDictionaries()
	if _dec != nil {
		return nil, _dec
	}
	var _caf []string
	for _feg := range _bcb {
		_caf = append(_caf, _feg)
	}
	_fe.Strings(_caf)
	_ab := map[string]_ff.PdfObject{}
	for _, _ea := range _caf {
		_fa := _bcb[_ea]
		_dd := _ff.TraceToDirectObject(_fa.Get("\u0056"))
		_ab[_ea] = _dd
	}
	return _ab, nil
}

// Root returns the Root of the FDF document.
func (_eca *fdfParser) Root() (*_ff.PdfObjectDictionary, error) {
	if _eca._ac != nil {
		if _gfd, _ccd := _eca.trace(_eca._ac.Get("\u0052\u006f\u006f\u0074")).(*_ff.PdfObjectDictionary); _ccd {
			if _bac, _fea := _eca.trace(_gfd.Get("\u0046\u0044\u0046")).(*_ff.PdfObjectDictionary); _fea {
				return _bac, nil
			}
		}
	}
	var _eec []int64
	for _deeb := range _eca._cg {
		_eec = append(_eec, _deeb)
	}
	_fe.Slice(_eec, func(_bbcb, _fcac int) bool { return _eec[_bbcb] < _eec[_fcac] })
	for _, _dag := range _eec {
		_edb := _eca._cg[_dag]
		if _faf, _abec := _eca.trace(_edb).(*_ff.PdfObjectDictionary); _abec {
			if _fcg, _fdf := _eca.trace(_faf.Get("\u0046\u0044\u0046")).(*_ff.PdfObjectDictionary); _fdf {
				return _fcg, nil
			}
		}
	}
	return nil, _d.New("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}
func (_agd *fdfParser) readAtLeast(_fd []byte, _agf int) (int, error) {
	_faa := _agf
	_cca := 0
	_df := 0
	for _faa > 0 {
		_gcg, _ead := _agd._gec.Read(_fd[_cca:])
		if _ead != nil {
			_c.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _gcg, _df, _ead.Error())
			return _cca, _d.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_df++
		_cca += _gcg
		_faa -= _gcg
	}
	return _cca, nil
}
func (_cdb *fdfParser) parse() error {
	_cdb._bfb.Seek(0, _db.SeekStart)
	_cdb._gec = _e.NewReader(_cdb._bfb)
	for {
		_cdb.skipComments()
		_cgga, _cbg := _cdb._gec.Peek(20)
		if _cbg != nil {
			_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return _cbg
		}
		if _dg.HasPrefix(string(_cgga), "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_cdb._gec.Discard(7)
			_cdb.skipSpaces()
			_cdb.skipComments()
			_eabb, _ := _cdb.parseDict()
			_cdb._ac = _eabb
			break
		}
		_ddb := _bf.FindStringSubmatchIndex(string(_cgga))
		if len(_ddb) < 6 {
			_c.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_cgga))
			return _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
		}
		_dab, _cbg := _cdb.parseIndirectObject()
		if _cbg != nil {
			return _cbg
		}
		switch _bba := _dab.(type) {
		case *_ff.PdfIndirectObject:
			_cdb._cg[_bba.ObjectNumber] = _bba
		case *_ff.PdfObjectStream:
			_cdb._cg[_bba.ObjectNumber] = _bba
		default:
			return _d.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	}
	return nil
}
func (_da *fdfParser) parseNumber() (_ff.PdfObject, error) { return _ff.ParseNumber(_da._gec) }
