package fdf

import (
	_ef "bufio"
	_f "bytes"
	_cdc "encoding/hex"
	_e "errors"
	_cd "fmt"
	_d "io"
	_g "os"
	_c "regexp"
	_b "sort"
	_af "strconv"
	_ec "strings"

	_gf "unitechio/gopdf/gopdf/common"
	_be "unitechio/gopdf/gopdf/core"
)

// Load loads FDF form data from `r`.
func Load(r _d.ReadSeeker) (*Data, error) {
	_fd, _cc := _bdd(r)
	if _cc != nil {
		return nil, _cc
	}
	_afb, _cc := _fd.Root()
	if _cc != nil {
		return nil, _cc
	}
	_bd, _eg := _be.GetArray(_afb.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_eg {
		return nil, _e.New("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	return &Data{_ea: _bd, _df: _afb}, nil
}

var _ceca = _c.MustCompile("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

func (_egd *fdfParser) parseNumber() (_be.PdfObject, error) { return _be.ParseNumber(_egd._dgf) }
func (_egc *fdfParser) readComment() (string, error) {
	var _eb _f.Buffer
	_, _dee := _egc.skipSpaces()
	if _dee != nil {
		return _eb.String(), _dee
	}
	_fcb := true
	for {
		_ee, _gd := _egc._dgf.Peek(1)
		if _gd != nil {
			_gf.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gd.Error())
			return _eb.String(), _gd
		}
		if _fcb && _ee[0] != '%' {
			return _eb.String(), _e.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_fcb = false
		if (_ee[0] != '\r') && (_ee[0] != '\n') {
			_fdd, _ := _egc._dgf.ReadByte()
			_eb.WriteByte(_fdd)
		} else {
			break
		}
	}
	return _eb.String(), nil
}

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data) FieldValues() (map[string]_be.PdfObject, error) {
	_fe, _cec := fdf.FieldDictionaries()
	if _cec != nil {
		return nil, _cec
	}
	var _dff []string
	for _de := range _fe {
		_dff = append(_dff, _de)
	}
	_b.Strings(_dff)
	_bda := map[string]_be.PdfObject{}
	for _, _fc := range _dff {
		_dcd := _fe[_fc]
		_gbc := _be.TraceToDirectObject(_dcd.Get("\u0056"))
		_bda[_fc] = _gbc
	}
	return _bda, nil
}

func (_gfd *fdfParser) seekToEOFMarker(_eea int64) error {
	_ggc := int64(0)
	_fgba := int64(1000)
	for _ggc < _eea {
		if _eea <= (_fgba + _ggc) {
			_fgba = _eea - _ggc
		}
		_, _gbga := _gfd._beg.Seek(-_ggc-_fgba, _d.SeekEnd)
		if _gbga != nil {
			return _gbga
		}
		_dbf := make([]byte, _fgba)
		_gfd._beg.Read(_dbf)
		_gf.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_dbf))
		_bgd := _gc.FindAllStringIndex(string(_dbf), -1)
		if _bgd != nil {
			_aad := _bgd[len(_bgd)-1]
			_gf.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _bgd)
			_gfd._beg.Seek(-_ggc-_fgba+int64(_aad[0]), _d.SeekEnd)
			return nil
		}
		_gf.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_ggc += _fgba
	}
	_gf.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _e.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

func _ecab(_afca string) (_be.PdfObjectReference, error) {
	_gcd := _be.PdfObjectReference{}
	_dfag := _ga.FindStringSubmatch(_afca)
	if len(_dfag) < 3 {
		_gf.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _gcd, _e.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_fcbf, _ccg := _af.Atoi(_dfag[1])
	if _ccg != nil {
		_gf.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030", _dfag[1])
		return _gcd, nil
	}
	_gcd.ObjectNumber = int64(_fcbf)
	_cfg, _ccg := _af.Atoi(_dfag[2])
	if _ccg != nil {
		_gf.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030", _dfag[2])
		return _gcd, nil
	}
	_gcd.GenerationNumber = int64(_cfg)
	return _gcd, nil
}

func (_da *fdfParser) skipSpaces() (int, error) {
	_fdg := 0
	for {
		_agc, _agcb := _da._dgf.ReadByte()
		if _agcb != nil {
			return 0, _agcb
		}
		if _be.IsWhiteSpace(_agc) {
			_fdg++
		} else {
			_da._dgf.UnreadByte()
			break
		}
	}
	return _fdg, nil
}

func (_fbgb *fdfParser) parseDict() (*_be.PdfObjectDictionary, error) {
	_gf.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_fgb := _be.MakeDict()
	_dbd, _ := _fbgb._dgf.ReadByte()
	if _dbd != '<' {
		return nil, _e.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_dbd, _ = _fbgb._dgf.ReadByte()
	if _dbd != '<' {
		return nil, _e.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_fbgb.skipSpaces()
		_fbgb.skipComments()
		_ebd, _cca := _fbgb._dgf.Peek(2)
		if _cca != nil {
			return nil, _cca
		}
		_gf.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_ebd), string(_ebd))
		if (_ebd[0] == '>') && (_ebd[1] == '>') {
			_gf.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_fbgb._dgf.ReadByte()
			_fbgb._dgf.ReadByte()
			break
		}
		_gf.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_dde, _cca := _fbgb.parseName()
		_gf.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _dde)
		if _cca != nil {
			_gf.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _cca)
			return nil, _cca
		}
		if len(_dde) > 4 && _dde[len(_dde)-4:] == "\u006e\u0075\u006c\u006c" {
			_acdg := _dde[0 : len(_dde)-4]
			_gf.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _dde)
			_gf.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _acdg)
			_fbgb.skipSpaces()
			_eee, _ := _fbgb._dgf.Peek(1)
			if _eee[0] == '/' {
				_fgb.Set(_acdg, _be.MakeNull())
				continue
			}
		}
		_fbgb.skipSpaces()
		_gbg, _cca := _fbgb.parseObject()
		if _cca != nil {
			return nil, _cca
		}
		_fgb.Set(_dde, _gbg)
		_gf.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _dde, _gbg.String())
	}
	_gf.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _fgb, nil
}

func (_ebb *fdfParser) parseFdfVersion() (int, int, error) {
	_ebb._beg.Seek(0, _d.SeekStart)
	_bceag := 20
	_ecg := make([]byte, _bceag)
	_ebb._beg.Read(_ecg)
	_eebc := _ceca.FindStringSubmatch(string(_ecg))
	if len(_eebc) < 3 {
		_cae, _bgf, _gaa := _ebb.seekFdfVersionTopDown()
		if _gaa != nil {
			_gf.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _gaa
		}
		return _cae, _bgf, nil
	}
	_cgdb, _add := _af.Atoi(_eebc[1])
	if _add != nil {
		return 0, 0, _add
	}
	_cbc, _add := _af.Atoi(_eebc[2])
	if _add != nil {
		return 0, 0, _add
	}
	_gf.Log.Debug("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _cgdb, _cbc)
	return _cgdb, _cbc, nil
}

func (_egaa *fdfParser) parseString() (*_be.PdfObjectString, error) {
	_egaa._dgf.ReadByte()
	var _fdc _f.Buffer
	_bdc := 1
	for {
		_ffd, _eaf := _egaa._dgf.Peek(1)
		if _eaf != nil {
			return _be.MakeString(_fdc.String()), _eaf
		}
		if _ffd[0] == '\\' {
			_egaa._dgf.ReadByte()
			_afa, _fdgd := _egaa._dgf.ReadByte()
			if _fdgd != nil {
				return _be.MakeString(_fdc.String()), _fdgd
			}
			if _be.IsOctalDigit(_afa) {
				_bce, _ced := _egaa._dgf.Peek(2)
				if _ced != nil {
					return _be.MakeString(_fdc.String()), _ced
				}
				var _dbc []byte
				_dbc = append(_dbc, _afa)
				for _, _ba := range _bce {
					if _be.IsOctalDigit(_ba) {
						_dbc = append(_dbc, _ba)
					} else {
						break
					}
				}
				_egaa._dgf.Discard(len(_dbc) - 1)
				_gf.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _dbc)
				_eab, _ced := _af.ParseUint(string(_dbc), 8, 32)
				if _ced != nil {
					return _be.MakeString(_fdc.String()), _ced
				}
				_fdc.WriteByte(byte(_eab))
				continue
			}
			switch _afa {
			case 'n':
				_fdc.WriteRune('\n')
			case 'r':
				_fdc.WriteRune('\r')
			case 't':
				_fdc.WriteRune('\t')
			case 'b':
				_fdc.WriteRune('\b')
			case 'f':
				_fdc.WriteRune('\f')
			case '(':
				_fdc.WriteRune('(')
			case ')':
				_fdc.WriteRune(')')
			case '\\':
				_fdc.WriteRune('\\')
			}
			continue
		} else if _ffd[0] == '(' {
			_bdc++
		} else if _ffd[0] == ')' {
			_bdc--
			if _bdc == 0 {
				_egaa._dgf.ReadByte()
				break
			}
		}
		_cbg, _ := _egaa._dgf.ReadByte()
		_fdc.WriteByte(_cbg)
	}
	return _be.MakeString(_fdc.String()), nil
}

func (_afc *fdfParser) parseBool() (_be.PdfObjectBool, error) {
	_cfe, _ae := _afc._dgf.Peek(4)
	if _ae != nil {
		return _be.PdfObjectBool(false), _ae
	}
	if (len(_cfe) >= 4) && (string(_cfe[:4]) == "\u0074\u0072\u0075\u0065") {
		_afc._dgf.Discard(4)
		return _be.PdfObjectBool(true), nil
	}
	_cfe, _ae = _afc._dgf.Peek(5)
	if _ae != nil {
		return _be.PdfObjectBool(false), _ae
	}
	if (len(_cfe) >= 5) && (string(_cfe[:5]) == "\u0066\u0061\u006cs\u0065") {
		_afc._dgf.Discard(5)
		return _be.PdfObjectBool(false), nil
	}
	return _be.PdfObjectBool(false), _e.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

func (_dcc *fdfParser) parseNull() (_be.PdfObjectNull, error) {
	_, _aee := _dcc._dgf.Discard(4)
	return _be.PdfObjectNull{}, _aee
}

func (_ece *fdfParser) readAtLeast(_ag []byte, _ega int) (int, error) {
	_cg := _ega
	_deg := 0
	_acg := 0
	for _cg > 0 {
		_fb, _ca := _ece._dgf.Read(_ag[_deg:])
		if _ca != nil {
			_gf.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _fb, _acg, _ca.Error())
			return _deg, _e.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_acg++
		_deg += _fb
		_cg -= _fb
	}
	return _deg, nil
}

var _gc = _c.MustCompile("\u0025\u0025\u0045O\u0046")

func _bdd(_cded _d.ReadSeeker) (*fdfParser, error) {
	_bad := &fdfParser{}
	_bad._beg = _cded
	_bad._gg = map[int64]_be.PdfObject{}
	_fce, _egg, _ggb := _bad.parseFdfVersion()
	if _ggb != nil {
		_gf.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _ggb)
		return nil, _ggb
	}
	_bad._bdb = _fce
	_bad._cgd = _egg
	_ggb = _bad.parse()
	return _bad, _ggb
}

func (_cgdg *fdfParser) parseArray() (*_be.PdfObjectArray, error) {
	_gdc := _be.MakeArray()
	_cgdg._dgf.ReadByte()
	for {
		_cgdg.skipSpaces()
		_dd, _gfe := _cgdg._dgf.Peek(1)
		if _gfe != nil {
			return _gdc, _gfe
		}
		if _dd[0] == ']' {
			_cgdg._dgf.ReadByte()
			break
		}
		_afad, _gfe := _cgdg.parseObject()
		if _gfe != nil {
			return _gdc, _gfe
		}
		_gdc.Append(_afad)
	}
	return _gdc, nil
}

func (_bc *fdfParser) parseName() (_be.PdfObjectName, error) {
	var _bec _f.Buffer
	_db := false
	for {
		_fgf, _ff := _bc._dgf.Peek(1)
		if _ff == _d.EOF {
			break
		}
		if _ff != nil {
			return _be.PdfObjectName(_bec.String()), _ff
		}
		if !_db {
			if _fgf[0] == '/' {
				_db = true
				_bc._dgf.ReadByte()
			} else if _fgf[0] == '%' {
				_bc.readComment()
				_bc.skipSpaces()
			} else {
				_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _fgf, _fgf)
				return _be.PdfObjectName(_bec.String()), _cd.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _fgf[0])
			}
		} else {
			if _be.IsWhiteSpace(_fgf[0]) {
				break
			} else if (_fgf[0] == '/') || (_fgf[0] == '[') || (_fgf[0] == '(') || (_fgf[0] == ']') || (_fgf[0] == '<') || (_fgf[0] == '>') {
				break
			} else if _fgf[0] == '#' {
				_dag, _cf := _bc._dgf.Peek(3)
				if _cf != nil {
					return _be.PdfObjectName(_bec.String()), _cf
				}
				_bc._dgf.Discard(3)
				_eeb, _cf := _cdc.DecodeString(string(_dag[1:3]))
				if _cf != nil {
					return _be.PdfObjectName(_bec.String()), _cf
				}
				_bec.Write(_eeb)
			} else {
				_cdb, _ := _bc._dgf.ReadByte()
				_bec.WriteByte(_cdb)
			}
		}
	}
	return _be.PdfObjectName(_bec.String()), nil
}

var (
	_eca = _c.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")
	_ga  = _c.MustCompile("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052")
)

func (_afd *fdfParser) getFileOffset() int64 {
	_agf, _ := _afd._beg.Seek(0, _d.SeekCurrent)
	_agf -= int64(_afd._dgf.Buffered())
	return _agf
}

var _bf = _c.MustCompile("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029")

// Data represents forms data format (FDF) file data.
type Data struct {
	_df *_be.PdfObjectDictionary
	_ea *_be.PdfObjectArray
}

func (_eadg *fdfParser) parse() error {
	_eadg._beg.Seek(0, _d.SeekStart)
	_eadg._dgf = _ef.NewReader(_eadg._beg)
	for {
		_eadg.skipComments()
		_ggeg, _ddb := _eadg._dgf.Peek(20)
		if _ddb != nil {
			_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return _ddb
		}
		if _ec.HasPrefix(string(_ggeg), "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_eadg._dgf.Discard(7)
			_eadg.skipSpaces()
			_eadg.skipComments()
			_fdba, _ := _eadg.parseDict()
			_eadg._gga = _fdba
			break
		}
		_gag := _eca.FindStringSubmatchIndex(string(_ggeg))
		if len(_gag) < 6 {
			_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_ggeg))
			return _e.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
		}
		_cdf, _ddb := _eadg.parseIndirectObject()
		if _ddb != nil {
			return _ddb
		}
		switch _cbf := _cdf.(type) {
		case *_be.PdfIndirectObject:
			_eadg._gg[_cbf.ObjectNumber] = _cbf
		case *_be.PdfObjectStream:
			_eadg._gg[_cbf.ObjectNumber] = _cbf
		default:
			return _e.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	}
	return nil
}

func (_cac *fdfParser) parseObject() (_be.PdfObject, error) {
	_gf.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_cac.skipSpaces()
	for {
		_ab, _bcea := _cac._dgf.Peek(2)
		if _bcea != nil {
			return nil, _bcea
		}
		_gf.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_ab))
		if _ab[0] == '/' {
			_gaf, _agfc := _cac.parseName()
			_gf.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _gaf)
			return &_gaf, _agfc
		} else if _ab[0] == '(' {
			_gf.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			return _cac.parseString()
		} else if _ab[0] == '[' {
			_gf.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			return _cac.parseArray()
		} else if (_ab[0] == '<') && (_ab[1] == '<') {
			_gf.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			return _cac.parseDict()
		} else if _ab[0] == '<' {
			_gf.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			return _cac.parseHexString()
		} else if _ab[0] == '%' {
			_cac.readComment()
			_cac.skipSpaces()
		} else {
			_gf.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_ab, _ = _cac._dgf.Peek(15)
			_dad := string(_ab)
			_gf.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _dad)
			if (len(_dad) > 3) && (_dad[:4] == "\u006e\u0075\u006c\u006c") {
				_ad, _ddd := _cac.parseNull()
				return &_ad, _ddd
			} else if (len(_dad) > 4) && (_dad[:5] == "\u0066\u0061\u006cs\u0065") {
				_ddc, _egdb := _cac.parseBool()
				return &_ddc, _egdb
			} else if (len(_dad) > 3) && (_dad[:4] == "\u0074\u0072\u0075\u0065") {
				_gbcf, _dca := _cac.parseBool()
				return &_gbcf, _dca
			}
			_fga := _ga.FindStringSubmatch(_dad)
			if len(_fga) > 1 {
				_ab, _ = _cac._dgf.ReadBytes('R')
				_gf.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_ab[:]))
				_bac, _fab := _ecab(string(_ab))
				return &_bac, _fab
			}
			_aa := _bb.FindStringSubmatch(_dad)
			if len(_aa) > 1 {
				_gf.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				return _cac.parseNumber()
			}
			_aa = _bf.FindStringSubmatch(_dad)
			if len(_aa) > 1 {
				_gf.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_gf.Log.Trace("\u0025\u0020\u0073", _aa)
				return _cac.parseNumber()
			}
			_gf.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _dad)
			return nil, _e.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}

// Root returns the Root of the FDF document.
func (_edg *fdfParser) Root() (*_be.PdfObjectDictionary, error) {
	if _edg._gga != nil {
		if _gge, _dddd := _edg.trace(_edg._gga.Get("\u0052\u006f\u006f\u0074")).(*_be.PdfObjectDictionary); _dddd {
			if _gaaa, _fee := _edg.trace(_gge.Get("\u0046\u0044\u0046")).(*_be.PdfObjectDictionary); _fee {
				return _gaaa, nil
			}
		}
	}
	var _bca []int64
	for _cbe := range _edg._gg {
		_bca = append(_bca, _cbe)
	}
	_b.Slice(_bca, func(_bdg, _cdbd int) bool { return _bca[_bdg] < _bca[_cdbd] })
	for _, _eaa := range _bca {
		_dda := _edg._gg[_eaa]
		if _agcd, _ecf := _edg.trace(_dda).(*_be.PdfObjectDictionary); _ecf {
			if _fabc, _fgbg := _edg.trace(_agcd.Get("\u0046\u0044\u0046")).(*_be.PdfObjectDictionary); _fgbg {
				return _fabc, nil
			}
		}
	}
	return nil, _e.New("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

var _bb = _c.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func (_acf *fdfParser) readTextLine() (string, error) {
	var _eac _f.Buffer
	for {
		_fbg, _ggg := _acf._dgf.Peek(1)
		if _ggg != nil {
			_gf.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _ggg.Error())
			return _eac.String(), _ggg
		}
		if (_fbg[0] != '\r') && (_fbg[0] != '\n') {
			_fcc, _ := _acf._dgf.ReadByte()
			_eac.WriteByte(_fcc)
		} else {
			break
		}
	}
	return _eac.String(), nil
}

func (_agd *fdfParser) setFileOffset(_agfe int64) {
	_agd._beg.Seek(_agfe, _d.SeekStart)
	_agd._dgf = _ef.NewReader(_agd._beg)
}

func _feg(_ecb string) (*fdfParser, error) {
	_ggab := fdfParser{}
	_cba := []byte(_ecb)
	_ceb := _f.NewReader(_cba)
	_ggab._beg = _ceb
	_ggab._gg = map[int64]_be.PdfObject{}
	_aeed := _ef.NewReader(_ceb)
	_ggab._dgf = _aeed
	_ggab._egb = int64(len(_ecb))
	return &_ggab, _ggab.parse()
}

func (_agg *fdfParser) parseHexString() (*_be.PdfObjectString, error) {
	_agg._dgf.ReadByte()
	var _ggge _f.Buffer
	for {
		_bdbb, _ecae := _agg._dgf.Peek(1)
		if _ecae != nil {
			return _be.MakeHexString(""), _ecae
		}
		if _bdbb[0] == '>' {
			_agg._dgf.ReadByte()
			break
		}
		_aggb, _ := _agg._dgf.ReadByte()
		if !_be.IsWhiteSpace(_aggb) {
			_ggge.WriteByte(_aggb)
		}
	}
	if _ggge.Len()%2 == 1 {
		_ggge.WriteRune('0')
	}
	_dfa, _gbb := _cdc.DecodeString(_ggge.String())
	if _gbb != nil {
		_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067", _ggge.String())
		return _be.MakeHexString(""), nil
	}
	return _be.MakeHexString(string(_dfa)), nil
}

func (_bg *fdfParser) skipComments() error {
	if _, _cb := _bg.skipSpaces(); _cb != nil {
		return _cb
	}
	_dae := true
	for {
		_dfc, _ded := _bg._dgf.Peek(1)
		if _ded != nil {
			_gf.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _ded.Error())
			return _ded
		}
		if _dae && _dfc[0] != '%' {
			return nil
		}
		_dae = false
		if (_dfc[0] != '\r') && (_dfc[0] != '\n') {
			_bg._dgf.ReadByte()
		} else {
			break
		}
	}
	return _bg.skipComments()
}

func (_efg *fdfParser) trace(_bba _be.PdfObject) _be.PdfObject {
	switch _gac := _bba.(type) {
	case *_be.PdfObjectReference:
		_dcg, _feb := _efg._gg[_gac.ObjectNumber].(*_be.PdfIndirectObject)
		if _feb {
			return _dcg.PdfObject
		}
		_gf.Log.Debug("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		return nil
	case *_be.PdfIndirectObject:
		return _gac.PdfObject
	}
	return _bba
}

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data) FieldDictionaries() (map[string]*_be.PdfObjectDictionary, error) {
	_fa := map[string]*_be.PdfObjectDictionary{}
	for _cdd := 0; _cdd < fdf._ea.Len(); _cdd++ {
		_cea, _eag := _be.GetDict(fdf._ea.Get(_cdd))
		if _eag {
			_cdg, _ := _be.GetString(_cea.Get("\u0054"))
			if _cdg != nil {
				_fa[_cdg.Str()] = _cea
			}
		}
	}
	return _fa, nil
}

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath(fdfPath string) (*Data, error) {
	_ce, _ac := _g.Open(fdfPath)
	if _ac != nil {
		return nil, _ac
	}
	defer _ce.Close()
	return Load(_ce)
}

func (_bcad *fdfParser) seekFdfVersionTopDown() (int, int, error) {
	_bcad._beg.Seek(0, _d.SeekStart)
	_bcad._dgf = _ef.NewReader(_bcad._beg)
	_gba := 20
	_gee := make([]byte, _gba)
	for {
		_aefa, _ecaef := _bcad._dgf.ReadByte()
		if _ecaef != nil {
			if _ecaef == _d.EOF {
				break
			} else {
				return 0, 0, _ecaef
			}
		}
		if _be.IsDecimalDigit(_aefa) && _gee[_gba-1] == '.' && _be.IsDecimalDigit(_gee[_gba-2]) && _gee[_gba-3] == '-' && _gee[_gba-4] == 'F' && _gee[_gba-5] == 'D' && _gee[_gba-6] == 'P' {
			_edd := int(_gee[_gba-2] - '0')
			_aacd := int(_aefa - '0')
			return _edd, _aacd, nil
		}
		_gee = append(_gee[1:_gba], _aefa)
	}
	return 0, 0, _e.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

func (_ecag *fdfParser) parseIndirectObject() (_be.PdfObject, error) {
	_fgd := _be.PdfIndirectObject{}
	_gf.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_acga, _abf := _ecag._dgf.Peek(20)
	if _abf != nil {
		_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
		return &_fgd, _abf
	}
	_gf.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_acga))
	_ecd := _eca.FindStringSubmatchIndex(string(_acga))
	if len(_ecd) < 6 {
		_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_acga))
		return &_fgd, _e.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_ecag._dgf.Discard(_ecd[0])
	_gf.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _ecd)
	_dbe := _ecd[1] - _ecd[0]
	_bdf := make([]byte, _dbe)
	_, _abf = _ecag.readAtLeast(_bdf, _dbe)
	if _abf != nil {
		_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _abf)
		return nil, _abf
	}
	_gf.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _bdf)
	_cde := _eca.FindStringSubmatch(string(_bdf))
	if len(_cde) < 3 {
		_gf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_bdf))
		return &_fgd, _e.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_fbfe, _ := _af.Atoi(_cde[1])
	_ebg, _ := _af.Atoi(_cde[2])
	_fgd.ObjectNumber = int64(_fbfe)
	_fgd.GenerationNumber = int64(_ebg)
	for {
		_degd, _bge := _ecag._dgf.Peek(2)
		if _bge != nil {
			return &_fgd, _bge
		}
		_gf.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_degd), string(_degd))
		if _be.IsWhiteSpace(_degd[0]) {
			_ecag.skipSpaces()
		} else if _degd[0] == '%' {
			_ecag.skipComments()
		} else if (_degd[0] == '<') && (_degd[1] == '<') {
			_gf.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_fgd.PdfObject, _bge = _ecag.parseDict()
			_gf.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _bge)
			if _bge != nil {
				return &_fgd, _bge
			}
			_gf.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_degd[0] == '/') || (_degd[0] == '(') || (_degd[0] == '[') || (_degd[0] == '<') {
			_fgd.PdfObject, _bge = _ecag.parseObject()
			if _bge != nil {
				return &_fgd, _bge
			}
			_gf.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else {
			if _degd[0] == 'e' {
				_dfe, _aac := _ecag.readTextLine()
				if _aac != nil {
					return nil, _aac
				}
				if len(_dfe) >= 6 && _dfe[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _degd[0] == 's' {
				_degd, _ = _ecag._dgf.Peek(10)
				if string(_degd[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_aag := 6
					if len(_degd) > 6 {
						if _be.IsWhiteSpace(_degd[_aag]) && _degd[_aag] != '\r' && _degd[_aag] != '\n' {
							_gf.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_aag++
						}
						if _degd[_aag] == '\r' {
							_aag++
							if _degd[_aag] == '\n' {
								_aag++
							}
						} else if _degd[_aag] == '\n' {
							_aag++
						}
					}
					_ecag._dgf.Discard(_aag)
					_fec, _daa := _fgd.PdfObject.(*_be.PdfObjectDictionary)
					if !_daa {
						return nil, _e.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_gf.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _fec)
					_bgc, _ge := _fec.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*_be.PdfObjectInteger)
					if !_ge {
						return nil, _e.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_cab := *_bgc
					if _cab < 0 {
						return nil, _e.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					if int64(_cab) > _ecag._egb {
						_gf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _e.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_fgc := make([]byte, _cab)
					_, _bge = _ecag.readAtLeast(_fgc, int(_cab))
					if _bge != nil {
						_gf.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_fgc), _fgc)
						_gf.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bge)
						return nil, _bge
					}
					_aef := _be.PdfObjectStream{}
					_aef.Stream = _fgc
					_aef.PdfObjectDictionary = _fgd.PdfObject.(*_be.PdfObjectDictionary)
					_aef.ObjectNumber = _fgd.ObjectNumber
					_aef.GenerationNumber = _fgd.GenerationNumber
					_ecag.skipSpaces()
					_ecag._dgf.Discard(9)
					_ecag.skipSpaces()
					return &_aef, nil
				}
			}
			_fgd.PdfObject, _bge = _ecag.parseObject()
			return &_fgd, _bge
		}
	}
	_gf.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_fgd, nil
}

type fdfParser struct {
	_bdb int
	_cgd int
	_gg  map[int64]_be.PdfObject
	_beg _d.ReadSeeker
	_dgf *_ef.Reader
	_egb int64
	_gga *_be.PdfObjectDictionary
}
