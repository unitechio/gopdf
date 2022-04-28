package fdf

import (
	_eb "bufio"
	_ea "bytes"
	_df "encoding/hex"
	_d "errors"
	_a "fmt"
	_gc "io"
	_cf "os"
	_f "regexp"
	_g "sort"
	_cg "strconv"
	_c "strings"

	_ac "bitbucket.org/shenghui0779/gopdf/common"
	_ebc "bitbucket.org/shenghui0779/gopdf/core"
)

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath(fdfPath string) (*Data, error) {
	_cfb, _ge := _cf.Open(fdfPath)
	if _ge != nil {
		return nil, _ge
	}
	defer _cfb.Close()
	return Load(_cfb)
}

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data) FieldDictionaries() (map[string]*_ebc.PdfObjectDictionary, error) {
	_bd := map[string]*_ebc.PdfObjectDictionary{}
	for _bf := 0; _bf < fdf._ff.Len(); _bf++ {
		_ba, _bda := _ebc.GetDict(fdf._ff.Get(_bf))
		if _bda {
			_gb, _ := _ebc.GetString(_ba.Get("\u0054"))
			if _gb != nil {
				_bd[_gb.Str()] = _ba
			}
		}
	}
	return _bd, nil
}
func (_gea *fdfParser) readComment() (string, error) {
	var _bag _ea.Buffer
	_, _bfg := _gea.skipSpaces()
	if _bfg != nil {
		return _bag.String(), _bfg
	}
	_ddd := true
	for {
		_db, _cbe := _gea._gac.Peek(1)
		if _cbe != nil {
			_ac.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _cbe.Error())
			return _bag.String(), _cbe
		}
		if _ddd && _db[0] != '%' {
			return _bag.String(), _d.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_ddd = false
		if (_db[0] != '\r') && (_db[0] != '\n') {
			_ega, _ := _gea._gac.ReadByte()
			_bag.WriteByte(_ega)
		} else {
			break
		}
	}
	return _bag.String(), nil
}
func (_ec *fdfParser) parseNumber() (_ebc.PdfObject, error) { return _ebc.ParseNumber(_ec._gac) }
func (_eaf *fdfParser) parseBool() (_ebc.PdfObjectBool, error) {
	_bff, _gfa := _eaf._gac.Peek(4)
	if _gfa != nil {
		return _ebc.PdfObjectBool(false), _gfa
	}
	if (len(_bff) >= 4) && (string(_bff[:4]) == "\u0074\u0072\u0075\u0065") {
		_eaf._gac.Discard(4)
		return _ebc.PdfObjectBool(true), nil
	}
	_bff, _gfa = _eaf._gac.Peek(5)
	if _gfa != nil {
		return _ebc.PdfObjectBool(false), _gfa
	}
	if (len(_bff) >= 5) && (string(_bff[:5]) == "\u0066\u0061\u006cs\u0065") {
		_eaf._gac.Discard(5)
		return _ebc.PdfObjectBool(false), nil
	}
	return _ebc.PdfObjectBool(false), _d.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

var _af = _f.MustCompile("\u0025\u0025\u0045O\u0046")
var _bb = _f.MustCompile("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029")

func (_badc *fdfParser) parseString() (*_ebc.PdfObjectString, error) {
	_badc._gac.ReadByte()
	var _bbf _ea.Buffer
	_bdab := 1
	for {
		_fgf, _ddda := _badc._gac.Peek(1)
		if _ddda != nil {
			return _ebc.MakeString(_bbf.String()), _ddda
		}
		if _fgf[0] == '\\' {
			_badc._gac.ReadByte()
			_fbg, _cag := _badc._gac.ReadByte()
			if _cag != nil {
				return _ebc.MakeString(_bbf.String()), _cag
			}
			if _ebc.IsOctalDigit(_fbg) {
				_gg, _gcf := _badc._gac.Peek(2)
				if _gcf != nil {
					return _ebc.MakeString(_bbf.String()), _gcf
				}
				var _baa []byte
				_baa = append(_baa, _fbg)
				for _, _dbec := range _gg {
					if _ebc.IsOctalDigit(_dbec) {
						_baa = append(_baa, _dbec)
					} else {
						break
					}
				}
				_badc._gac.Discard(len(_baa) - 1)
				_ac.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _baa)
				_daa, _gcf := _cg.ParseUint(string(_baa), 8, 32)
				if _gcf != nil {
					return _ebc.MakeString(_bbf.String()), _gcf
				}
				_bbf.WriteByte(byte(_daa))
				continue
			}
			switch _fbg {
			case 'n':
				_bbf.WriteRune('\n')
			case 'r':
				_bbf.WriteRune('\r')
			case 't':
				_bbf.WriteRune('\t')
			case 'b':
				_bbf.WriteRune('\b')
			case 'f':
				_bbf.WriteRune('\f')
			case '(':
				_bbf.WriteRune('(')
			case ')':
				_bbf.WriteRune(')')
			case '\\':
				_bbf.WriteRune('\\')
			}
			continue
		} else if _fgf[0] == '(' {
			_bdab++
		} else if _fgf[0] == ')' {
			_bdab--
			if _bdab == 0 {
				_badc._gac.ReadByte()
				break
			}
		}
		_daab, _ := _badc._gac.ReadByte()
		_bbf.WriteByte(_daab)
	}
	return _ebc.MakeString(_bbf.String()), nil
}

var _bg = _f.MustCompile("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052")
var _cb = _f.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

// Root returns the Root of the FDF document.
func (_egbf *fdfParser) Root() (*_ebc.PdfObjectDictionary, error) {
	if _egbf._gacd != nil {
		if _cae, _ebbc := _egbf.trace(_egbf._gacd.Get("\u0052\u006f\u006f\u0074")).(*_ebc.PdfObjectDictionary); _ebbc {
			if _aad, _dbc := _egbf.trace(_cae.Get("\u0046\u0044\u0046")).(*_ebc.PdfObjectDictionary); _dbc {
				return _aad, nil
			}
		}
	}
	var _cbd []int64
	for _ecge := range _egbf._efd {
		_cbd = append(_cbd, _ecge)
	}
	_g.Slice(_cbd, func(_dcc, _bacf int) bool { return _cbd[_dcc] < _cbd[_bacf] })
	for _, _edgag := range _cbd {
		_add := _egbf._efd[_edgag]
		if _daf, _fdb := _egbf.trace(_add).(*_ebc.PdfObjectDictionary); _fdb {
			if _gfg, _ecbf := _egbf.trace(_daf.Get("\u0046\u0044\u0046")).(*_ebc.PdfObjectDictionary); _ecbf {
				return _gfg, nil
			}
		}
	}
	return nil, _d.New("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}
func (_daae *fdfParser) parseHexString() (*_ebc.PdfObjectString, error) {
	_daae._gac.ReadByte()
	var _daec _ea.Buffer
	for {
		_fcf, _dgg := _daae._gac.Peek(1)
		if _dgg != nil {
			return _ebc.MakeHexString(""), _dgg
		}
		if _fcf[0] == '>' {
			_daae._gac.ReadByte()
			break
		}
		_fbc, _ := _daae._gac.ReadByte()
		if !_ebc.IsWhiteSpace(_fbc) {
			_daec.WriteByte(_fbc)
		}
	}
	if _daec.Len()%2 == 1 {
		_daec.WriteRune('0')
	}
	_acg, _daee := _df.DecodeString(_daec.String())
	if _daee != nil {
		_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067", _daec.String())
		return _ebc.MakeHexString(""), nil
	}
	return _ebc.MakeHexString(string(_acg)), nil
}
func (_cfd *fdfParser) parseDict() (*_ebc.PdfObjectDictionary, error) {
	_ac.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_dcfa := _ebc.MakeDict()
	_egb, _ := _cfd._gac.ReadByte()
	if _egb != '<' {
		return nil, _d.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_egb, _ = _cfd._gac.ReadByte()
	if _egb != '<' {
		return nil, _d.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_cfd.skipSpaces()
		_cfd.skipComments()
		_adf, _acf := _cfd._gac.Peek(2)
		if _acf != nil {
			return nil, _acf
		}
		_ac.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_adf), string(_adf))
		if (_adf[0] == '>') && (_adf[1] == '>') {
			_ac.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_cfd._gac.ReadByte()
			_cfd._gac.ReadByte()
			break
		}
		_ac.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_eag, _acf := _cfd.parseName()
		_ac.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _eag)
		if _acf != nil {
			_ac.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _acf)
			return nil, _acf
		}
		if len(_eag) > 4 && _eag[len(_eag)-4:] == "\u006e\u0075\u006c\u006c" {
			_adg := _eag[0 : len(_eag)-4]
			_ac.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _eag)
			_ac.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _adg)
			_cfd.skipSpaces()
			_fae, _ := _cfd._gac.Peek(1)
			if _fae[0] == '/' {
				_dcfa.Set(_adg, _ebc.MakeNull())
				continue
			}
		}
		_cfd.skipSpaces()
		_fge, _acf := _cfd.parseObject()
		if _acf != nil {
			return nil, _acf
		}
		_dcfa.Set(_eag, _fge)
		_ac.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _eag, _fge.String())
	}
	_ac.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _dcfa, nil
}
func (_daabe *fdfParser) parse() error {
	_daabe._fff.Seek(0, _gc.SeekStart)
	_daabe._gac = _eb.NewReader(_daabe._fff)
	for {
		_daabe.skipComments()
		_dcfb, _ecgg := _daabe._gac.Peek(20)
		if _ecgg != nil {
			_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return _ecgg
		}
		if _c.HasPrefix(string(_dcfb), "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_daabe._gac.Discard(7)
			_daabe.skipSpaces()
			_daabe.skipComments()
			_acc, _ := _daabe.parseDict()
			_daabe._gacd = _acc
			break
		}
		_cege := _cb.FindStringSubmatchIndex(string(_dcfb))
		if len(_cege) < 6 {
			_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_dcfb))
			return _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
		}
		_cee, _ecgg := _daabe.parseIndirectObject()
		if _ecgg != nil {
			return _ecgg
		}
		switch _edd := _cee.(type) {
		case *_ebc.PdfIndirectObject:
			_daabe._efd[_edd.ObjectNumber] = _edd
		case *_ebc.PdfObjectStream:
			_daabe._efd[_edd.ObjectNumber] = _edd
		default:
			return _d.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	}
	return nil
}

var _gd = _f.MustCompile("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")

func _bage(_gaa string) (*fdfParser, error) {
	_befc := fdfParser{}
	_dbee := []byte(_gaa)
	_eab := _ea.NewReader(_dbee)
	_befc._fff = _eab
	_befc._efd = map[int64]_ebc.PdfObject{}
	_gaca := _eb.NewReader(_eab)
	_befc._gac = _gaca
	_befc._ada = int64(len(_gaa))
	return &_befc, _befc.parse()
}

// Data represents forms data format (FDF) file data.
type Data struct {
	_b  *_ebc.PdfObjectDictionary
	_ff *_ebc.PdfObjectArray
}

func (_dag *fdfParser) parseObject() (_ebc.PdfObject, error) {
	_ac.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_dag.skipSpaces()
	for {
		_cea, _cfa := _dag._gac.Peek(2)
		if _cfa != nil {
			return nil, _cfa
		}
		_ac.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_cea))
		if _cea[0] == '/' {
			_de, _deb := _dag.parseName()
			_ac.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _de)
			return &_de, _deb
		} else if _cea[0] == '(' {
			_ac.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			return _dag.parseString()
		} else if _cea[0] == '[' {
			_ac.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			return _dag.parseArray()
		} else if (_cea[0] == '<') && (_cea[1] == '<') {
			_ac.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			return _dag.parseDict()
		} else if _cea[0] == '<' {
			_ac.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			return _dag.parseHexString()
		} else if _cea[0] == '%' {
			_dag.readComment()
			_dag.skipSpaces()
		} else {
			_ac.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_cea, _ = _dag._gac.Peek(15)
			_bagad := string(_cea)
			_ac.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _bagad)
			if (len(_bagad) > 3) && (_bagad[:4] == "\u006e\u0075\u006c\u006c") {
				_fgb, _cbb := _dag.parseNull()
				return &_fgb, _cbb
			} else if (len(_bagad) > 4) && (_bagad[:5] == "\u0066\u0061\u006cs\u0065") {
				_ggc, _aef := _dag.parseBool()
				return &_ggc, _aef
			} else if (len(_bagad) > 3) && (_bagad[:4] == "\u0074\u0072\u0075\u0065") {
				_ecb, _gebg := _dag.parseBool()
				return &_ecb, _gebg
			}
			_dcf := _bg.FindStringSubmatch(_bagad)
			if len(_dcf) > 1 {
				_cea, _ = _dag._gac.ReadBytes('R')
				_ac.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_cea[:]))
				_dce, _ebf := _bfga(string(_cea))
				return &_dce, _ebf
			}
			_agf := _fb.FindStringSubmatch(_bagad)
			if len(_agf) > 1 {
				_ac.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				return _dag.parseNumber()
			}
			_agf = _bb.FindStringSubmatch(_bagad)
			if len(_agf) > 1 {
				_ac.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_ac.Log.Trace("\u0025\u0020\u0073", _agf)
				return _dag.parseNumber()
			}
			_ac.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _bagad)
			return nil, _d.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}
func (_ade *fdfParser) readAtLeast(_ffb []byte, _da int) (int, error) {
	_bdd := _da
	_fc := 0
	_eg := 0
	for _bdd > 0 {
		_geb, _be := _ade._gac.Read(_ffb[_fc:])
		if _be != nil {
			_ac.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _geb, _eg, _be.Error())
			return _fc, _d.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_eg++
		_fc += _geb
		_bdd -= _geb
	}
	return _fc, nil
}
func (_baga *fdfParser) parseNull() (_ebc.PdfObjectNull, error) {
	_, _cge := _baga._gac.Discard(4)
	return _ebc.PdfObjectNull{}, _cge
}
func (_ga *fdfParser) getFileOffset() int64 {
	_fcg, _ := _ga._fff.Seek(0, _gc.SeekCurrent)
	_fcg -= int64(_ga._gac.Buffered())
	return _fcg
}
func (_eefa *fdfParser) trace(_bebe _ebc.PdfObject) _ebc.PdfObject {
	switch _ecbfe := _bebe.(type) {
	case *_ebc.PdfObjectReference:
		_gfd, _fdg := _eefa._efd[_ecbfe.ObjectNumber].(*_ebc.PdfIndirectObject)
		if _fdg {
			return _gfd.PdfObject
		}
		_ac.Log.Debug("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		return nil
	case *_ebc.PdfIndirectObject:
		return _ecbfe.PdfObject
	}
	return _bebe
}
func (_feb *fdfParser) parseIndirectObject() (_ebc.PdfObject, error) {
	_ced := _ebc.PdfIndirectObject{}
	_ac.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_aaa, _gcg := _feb._gac.Peek(20)
	if _gcg != nil {
		_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
		return &_ced, _gcg
	}
	_ac.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_aaa))
	_gcb := _cb.FindStringSubmatchIndex(string(_aaa))
	if len(_gcb) < 6 {
		_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_aaa))
		return &_ced, _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_feb._gac.Discard(_gcb[0])
	_ac.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _gcb)
	_dgge := _gcb[1] - _gcb[0]
	_fea := make([]byte, _dgge)
	_, _gcg = _feb.readAtLeast(_fea, _dgge)
	if _gcg != nil {
		_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _gcg)
		return nil, _gcg
	}
	_ac.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _fea)
	_fee := _cb.FindStringSubmatch(string(_fea))
	if len(_fee) < 3 {
		_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_fea))
		return &_ced, _d.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_adga, _ := _cg.Atoi(_fee[1])
	_cac, _ := _cg.Atoi(_fee[2])
	_ced.ObjectNumber = int64(_adga)
	_ced.GenerationNumber = int64(_cac)
	for {
		_acb, _beb := _feb._gac.Peek(2)
		if _beb != nil {
			return &_ced, _beb
		}
		_ac.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_acb), string(_acb))
		if _ebc.IsWhiteSpace(_acb[0]) {
			_feb.skipSpaces()
		} else if _acb[0] == '%' {
			_feb.skipComments()
		} else if (_acb[0] == '<') && (_acb[1] == '<') {
			_ac.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_ced.PdfObject, _beb = _feb.parseDict()
			_ac.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _beb)
			if _beb != nil {
				return &_ced, _beb
			}
			_ac.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_acb[0] == '/') || (_acb[0] == '(') || (_acb[0] == '[') || (_acb[0] == '<') {
			_ced.PdfObject, _beb = _feb.parseObject()
			if _beb != nil {
				return &_ced, _beb
			}
			_ac.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else {
			if _acb[0] == 'e' {
				_baec, _edga := _feb.readTextLine()
				if _edga != nil {
					return nil, _edga
				}
				if len(_baec) >= 6 && _baec[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _acb[0] == 's' {
				_acb, _ = _feb._gac.Peek(10)
				if string(_acb[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_fgd := 6
					if len(_acb) > 6 {
						if _ebc.IsWhiteSpace(_acb[_fgd]) && _acb[_fgd] != '\r' && _acb[_fgd] != '\n' {
							_ac.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_fgd++
						}
						if _acb[_fgd] == '\r' {
							_fgd++
							if _acb[_fgd] == '\n' {
								_fgd++
							}
						} else if _acb[_fgd] == '\n' {
							_fgd++
						}
					}
					_feb._gac.Discard(_fgd)
					_efce, _gcgb := _ced.PdfObject.(*_ebc.PdfObjectDictionary)
					if !_gcgb {
						return nil, _d.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_ac.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _efce)
					_aecg, _fed := _efce.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*_ebc.PdfObjectInteger)
					if !_fed {
						return nil, _d.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_fbag := *_aecg
					if _fbag < 0 {
						return nil, _d.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					if int64(_fbag) > _feb._ada {
						_ac.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _d.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_ceg := make([]byte, _fbag)
					_, _beb = _feb.readAtLeast(_ceg, int(_fbag))
					if _beb != nil {
						_ac.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_ceg), _ceg)
						_ac.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _beb)
						return nil, _beb
					}
					_aae := _ebc.PdfObjectStream{}
					_aae.Stream = _ceg
					_aae.PdfObjectDictionary = _ced.PdfObject.(*_ebc.PdfObjectDictionary)
					_aae.ObjectNumber = _ced.ObjectNumber
					_aae.GenerationNumber = _ced.GenerationNumber
					_feb.skipSpaces()
					_feb._gac.Discard(9)
					_feb.skipSpaces()
					return &_aae, nil
				}
			}
			_ced.PdfObject, _beb = _feb.parseObject()
			return &_ced, _beb
		}
	}
	_ac.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_ced, nil
}
func _bea(_cbg _gc.ReadSeeker) (*fdfParser, error) {
	_cgcd := &fdfParser{}
	_cgcd._fff = _cbg
	_cgcd._efd = map[int64]_ebc.PdfObject{}
	_cfed, _ace, _fbd := _cgcd.parseFdfVersion()
	if _fbd != nil {
		_ac.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _fbd)
		return nil, _fbd
	}
	_cgcd._fbb = _cfed
	_cgcd._cfbg = _ace
	_fbd = _cgcd.parse()
	return _cgcd, _fbd
}
func (_cga *fdfParser) readTextLine() (string, error) {
	var _gdf _ea.Buffer
	for {
		_bfgc, _aec := _cga._gac.Peek(1)
		if _aec != nil {
			_ac.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _aec.Error())
			return _gdf.String(), _aec
		}
		if (_bfgc[0] != '\r') && (_bfgc[0] != '\n') {
			_gf, _ := _cga._gac.ReadByte()
			_gdf.WriteByte(_gf)
		} else {
			break
		}
	}
	return _gdf.String(), nil
}

type fdfParser struct {
	_fbb  int
	_cfbg int
	_efd  map[int64]_ebc.PdfObject
	_fff  _gc.ReadSeeker
	_gac  *_eb.Reader
	_ada  int64
	_gacd *_ebc.PdfObjectDictionary
}

// Load loads FDF form data from `r`.
func Load(r _gc.ReadSeeker) (*Data, error) {
	_gca, _dc := _bea(r)
	if _dc != nil {
		return nil, _dc
	}
	_ed, _dc := _gca.Root()
	if _dc != nil {
		return nil, _dc
	}
	_ef, _cfe := _ebc.GetArray(_ed.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_cfe {
		return nil, _d.New("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	return &Data{_ff: _ef, _b: _ed}, nil
}
func (_ce *fdfParser) setFileOffset(_dff int64) {
	_ce._fff.Seek(_dff, _gc.SeekStart)
	_ce._gac = _eb.NewReader(_ce._fff)
}
func (_egac *fdfParser) parseArray() (*_ebc.PdfObjectArray, error) {
	_ab := _ebc.MakeArray()
	_egac._gac.ReadByte()
	for {
		_egac.skipSpaces()
		_ece, _bbc := _egac._gac.Peek(1)
		if _bbc != nil {
			return _ab, _bbc
		}
		if _ece[0] == ']' {
			_egac._gac.ReadByte()
			break
		}
		_ecg, _bbc := _egac.parseObject()
		if _bbc != nil {
			return _ab, _bbc
		}
		_ab.Append(_ecg)
	}
	return _ab, nil
}

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data) FieldValues() (map[string]_ebc.PdfObject, error) {
	_bac, _fa := fdf.FieldDictionaries()
	if _fa != nil {
		return nil, _fa
	}
	var _dg []string
	for _dd := range _bac {
		_dg = append(_dg, _dd)
	}
	_g.Strings(_dg)
	_cd := map[string]_ebc.PdfObject{}
	for _, _efg := range _dg {
		_cgb := _bac[_efg]
		_ad := _ebc.TraceToDirectObject(_cgb.Get("\u0056"))
		_cd[_efg] = _ad
	}
	return _cd, nil
}
func (_egd *fdfParser) seekFdfVersionTopDown() (int, int, error) {
	_egd._fff.Seek(0, _gc.SeekStart)
	_egd._gac = _eb.NewReader(_egd._fff)
	_dbg := 20
	_ffd := make([]byte, _dbg)
	for {
		_aaaa, _cba := _egd._gac.ReadByte()
		if _cba != nil {
			if _cba == _gc.EOF {
				break
			} else {
				return 0, 0, _cba
			}
		}
		if _ebc.IsDecimalDigit(_aaaa) && _ffd[_dbg-1] == '.' && _ebc.IsDecimalDigit(_ffd[_dbg-2]) && _ffd[_dbg-3] == '-' && _ffd[_dbg-4] == 'F' && _ffd[_dbg-5] == 'D' && _ffd[_dbg-6] == 'P' {
			_gbb := int(_ffd[_dbg-2] - '0')
			_dea := int(_aaaa - '0')
			return _gbb, _dea, nil
		}
		_ffd = append(_ffd[1:_dbg], _aaaa)
	}
	return 0, 0, _d.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}
func (_bgb *fdfParser) seekToEOFMarker(_bc int64) error {
	_abe := int64(0)
	_edg := int64(1000)
	for _abe < _bc {
		if _bc <= (_edg + _abe) {
			_edg = _bc - _abe
		}
		_, _fcd := _bgb._fff.Seek(-_abe-_edg, _gc.SeekEnd)
		if _fcd != nil {
			return _fcd
		}
		_fffd := make([]byte, _edg)
		_bgb._fff.Read(_fffd)
		_ac.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_fffd))
		_gge := _af.FindAllStringIndex(string(_fffd), -1)
		if _gge != nil {
			_bae := _gge[len(_gge)-1]
			_ac.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _gge)
			_bgb._fff.Seek(-_abe-_edg+int64(_bae[0]), _gc.SeekEnd)
			return nil
		}
		_ac.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_abe += _edg
	}
	_ac.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _d.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

var _fb = _f.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")

func (_gefd *fdfParser) parseFdfVersion() (int, int, error) {
	_gefd._fff.Seek(0, _gc.SeekStart)
	_egg := 20
	_efc := make([]byte, _egg)
	_gefd._fff.Read(_efc)
	_gbf := _gd.FindStringSubmatch(string(_efc))
	if len(_gbf) < 3 {
		_eef, _agdf, _edc := _gefd.seekFdfVersionTopDown()
		if _edc != nil {
			_ac.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _edc
		}
		return _eef, _agdf, nil
	}
	_daed, _caf := _cg.Atoi(_gbf[1])
	if _caf != nil {
		return 0, 0, _caf
	}
	_abf, _caf := _cg.Atoi(_gbf[2])
	if _caf != nil {
		return 0, 0, _caf
	}
	_ac.Log.Debug("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _daed, _abf)
	return _daed, _abf, nil
}
func (_aa *fdfParser) skipSpaces() (int, error) {
	_gce := 0
	for {
		_fg, _ddc := _aa._gac.ReadByte()
		if _ddc != nil {
			return 0, _ddc
		}
		if _ebc.IsWhiteSpace(_fg) {
			_gce++
		} else {
			_aa._gac.UnreadByte()
			break
		}
	}
	return _gce, nil
}
func (_bfd *fdfParser) parseName() (_ebc.PdfObjectName, error) {
	var _adc _ea.Buffer
	_gaf := false
	for {
		_fffc, _dbe := _bfd._gac.Peek(1)
		if _dbe == _gc.EOF {
			break
		}
		if _dbe != nil {
			return _ebc.PdfObjectName(_adc.String()), _dbe
		}
		if !_gaf {
			if _fffc[0] == '/' {
				_gaf = true
				_bfd._gac.ReadByte()
			} else if _fffc[0] == '%' {
				_bfd.readComment()
				_bfd.skipSpaces()
			} else {
				_ac.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _fffc, _fffc)
				return _ebc.PdfObjectName(_adc.String()), _a.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _fffc[0])
			}
		} else {
			if _ebc.IsWhiteSpace(_fffc[0]) {
				break
			} else if (_fffc[0] == '/') || (_fffc[0] == '[') || (_fffc[0] == '(') || (_fffc[0] == ']') || (_fffc[0] == '<') || (_fffc[0] == '>') {
				break
			} else if _fffc[0] == '#' {
				_bad, _gef := _bfd._gac.Peek(3)
				if _gef != nil {
					return _ebc.PdfObjectName(_adc.String()), _gef
				}
				_bfd._gac.Discard(3)
				_cc, _gef := _df.DecodeString(string(_bad[1:3]))
				if _gef != nil {
					return _ebc.PdfObjectName(_adc.String()), _gef
				}
				_adc.Write(_cc)
			} else {
				_cbf, _ := _bfd._gac.ReadByte()
				_adc.WriteByte(_cbf)
			}
		}
	}
	return _ebc.PdfObjectName(_adc.String()), nil
}
func _bfga(_fgc string) (_ebc.PdfObjectReference, error) {
	_ebb := _ebc.PdfObjectReference{}
	_bef := _bg.FindStringSubmatch(_fgc)
	if len(_bef) < 3 {
		_ac.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _ebb, _d.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_gba, _efdg := _cg.Atoi(_bef[1])
	if _efdg != nil {
		_ac.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030", _bef[1])
		return _ebb, nil
	}
	_ebb.ObjectNumber = int64(_gba)
	_beee, _efdg := _cg.Atoi(_bef[2])
	if _efdg != nil {
		_ac.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030", _bef[2])
		return _ebb, nil
	}
	_ebb.GenerationNumber = int64(_beee)
	return _ebb, nil
}
func (_ag *fdfParser) skipComments() error {
	if _, _ca := _ag.skipSpaces(); _ca != nil {
		return _ca
	}
	_dae := true
	for {
		_bbd, _dac := _ag._gac.Peek(1)
		if _dac != nil {
			_ac.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _dac.Error())
			return _dac
		}
		if _dae && _bbd[0] != '%' {
			return nil
		}
		_dae = false
		if (_bbd[0] != '\r') && (_bbd[0] != '\n') {
			_ag._gac.ReadByte()
		} else {
			break
		}
	}
	return _ag.skipComments()
}
