package fdf

import (
	_ca "bufio"
	_aa "bytes"
	_f "encoding/hex"
	_c "errors"
	_be "fmt"
	_e "io"
	_ba "os"
	_b "regexp"
	_d "sort"
	_ag "strconv"
	_a "strings"

	_cd "bitbucket.org/shenghui0779/gopdf/common"
	_age "bitbucket.org/shenghui0779/gopdf/core"
)

var _bc = _b.MustCompile("\u0025\u0025\u0045O\u0046")

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath(fdfPath string) (*Data, error) {
	_eg, _ac := _ba.Open(fdfPath)
	if _ac != nil {
		return nil, _ac
	}
	defer _eg.Close()
	return Load(_eg)
}
func (_cbaf *fdfParser) seekFdfVersionTopDown() (int, int, error) {
	_cbaf._gaf.Seek(0, _e.SeekStart)
	_cbaf._dge = _ca.NewReader(_cbaf._gaf)
	_ddc := 20
	_ecg := make([]byte, _ddc)
	for {
		_gce, _gee := _cbaf._dge.ReadByte()
		if _gee != nil {
			if _gee == _e.EOF {
				break
			} else {
				return 0, 0, _gee
			}
		}
		if _age.IsDecimalDigit(_gce) && _ecg[_ddc-1] == '.' && _age.IsDecimalDigit(_ecg[_ddc-2]) && _ecg[_ddc-3] == '-' && _ecg[_ddc-4] == 'F' && _ecg[_ddc-5] == 'D' && _ecg[_ddc-6] == 'P' {
			_gedb := int(_ecg[_ddc-2] - '0')
			_dff := int(_gce - '0')
			return _gedb, _dff, nil
		}
		_ecg = append(_ecg[1:_ddc], _gce)
	}
	return 0, 0, _c.New("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}
func (_bdb *fdfParser) parseArray() (*_age.PdfObjectArray, error) {
	_fcd := _age.MakeArray()
	_bdb._dge.ReadByte()
	for {
		_bdb.skipSpaces()
		_eac, _aed := _bdb._dge.Peek(1)
		if _aed != nil {
			return _fcd, _aed
		}
		if _eac[0] == ']' {
			_bdb._dge.ReadByte()
			break
		}
		_ccg, _aed := _bdb.parseObject()
		if _aed != nil {
			return _fcd, _aed
		}
		_fcd.Append(_ccg)
	}
	return _fcd, nil
}
func (_gbgee *fdfParser) parseObject() (_age.PdfObject, error) {
	_cd.Log.Trace("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074")
	_gbgee.skipSpaces()
	for {
		_abba, _gaab := _gbgee._dge.Peek(2)
		if _gaab != nil {
			return nil, _gaab
		}
		_cd.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_abba))
		if _abba[0] == '/' {
			_egda, _fff := _gbgee.parseName()
			_cd.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _egda)
			return &_egda, _fff
		} else if _abba[0] == '(' {
			_cd.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			return _gbgee.parseString()
		} else if _abba[0] == '[' {
			_cd.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			return _gbgee.parseArray()
		} else if (_abba[0] == '<') && (_abba[1] == '<') {
			_cd.Log.Trace("\u002d>\u0044\u0069\u0063\u0074\u0021")
			return _gbgee.parseDict()
		} else if _abba[0] == '<' {
			_cd.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021")
			return _gbgee.parseHexString()
		} else if _abba[0] == '%' {
			_gbgee.readComment()
			_gbgee.skipSpaces()
		} else {
			_cd.Log.Trace("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f")
			_abba, _ = _gbgee._dge.Peek(15)
			_deba := string(_abba)
			_cd.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _deba)
			if (len(_deba) > 3) && (_deba[:4] == "\u006e\u0075\u006c\u006c") {
				_ecd, _gad := _gbgee.parseNull()
				return &_ecd, _gad
			} else if (len(_deba) > 4) && (_deba[:5] == "\u0066\u0061\u006cs\u0065") {
				_ef, _dgaa := _gbgee.parseBool()
				return &_ef, _dgaa
			} else if (len(_deba) > 3) && (_deba[:4] == "\u0074\u0072\u0075\u0065") {
				_bae, _dde := _gbgee.parseBool()
				return &_bae, _dde
			}
			_acda := _gaa.FindStringSubmatch(_deba)
			if len(_acda) > 1 {
				_abba, _ = _gbgee._dge.ReadBytes('R')
				_cd.Log.Trace("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027", string(_abba[:]))
				_ege, _dedb := _gdg(string(_abba))
				return &_ege, _dedb
			}
			_fce := _bd.FindStringSubmatch(_deba)
			if len(_fce) > 1 {
				_cd.Log.Trace("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021")
				return _gbgee.parseNumber()
			}
			_fce = _gg.FindStringSubmatch(_deba)
			if len(_fce) > 1 {
				_cd.Log.Trace("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021")
				_cd.Log.Trace("\u0025\u0020\u0073", _fce)
				return _gbgee.parseNumber()
			}
			_cd.Log.Debug("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)", _deba)
			return nil, _c.New("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e")
		}
	}
}
func (_gdf *fdfParser) readTextLine() (string, error) {
	var _geg _aa.Buffer
	for {
		_adg, _gbge := _gdf._dge.Peek(1)
		if _gbge != nil {
			_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gbge.Error())
			return _geg.String(), _gbge
		}
		if (_adg[0] != '\r') && (_adg[0] != '\n') {
			_ged, _ := _gdf._dge.ReadByte()
			_geg.WriteByte(_ged)
		} else {
			break
		}
	}
	return _geg.String(), nil
}
func _gdg(_cac string) (_age.PdfObjectReference, error) {
	_bcc := _age.PdfObjectReference{}
	_ebc := _gaa.FindStringSubmatch(_cac)
	if len(_ebc) < 3 {
		_cd.Log.Debug("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		return _bcc, _c.New("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_cag, _aec := _ag.Atoi(_ebc[1])
	if _aec != nil {
		_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030", _ebc[1])
		return _bcc, nil
	}
	_bcc.ObjectNumber = int64(_cag)
	_fcbc, _aec := _ag.Atoi(_ebc[2])
	if _aec != nil {
		_cd.Log.Debug("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030", _ebc[2])
		return _bcc, nil
	}
	_bcc.GenerationNumber = int64(_fcbc)
	return _bcc, nil
}
func (_eacg *fdfParser) parseBool() (_age.PdfObjectBool, error) {
	_ffa, _ccgf := _eacg._dge.Peek(4)
	if _ccgf != nil {
		return _age.PdfObjectBool(false), _ccgf
	}
	if (len(_ffa) >= 4) && (string(_ffa[:4]) == "\u0074\u0072\u0075\u0065") {
		_eacg._dge.Discard(4)
		return _age.PdfObjectBool(true), nil
	}
	_ffa, _ccgf = _eacg._dge.Peek(5)
	if _ccgf != nil {
		return _age.PdfObjectBool(false), _ccgf
	}
	if (len(_ffa) >= 5) && (string(_ffa[:5]) == "\u0066\u0061\u006cs\u0065") {
		_eacg._dge.Discard(5)
		return _age.PdfObjectBool(false), nil
	}
	return _age.PdfObjectBool(false), _c.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}
func (_gcg *fdfParser) trace(_accd _age.PdfObject) _age.PdfObject {
	switch _fgd := _accd.(type) {
	case *_age.PdfObjectReference:
		_cga, _bde := _gcg._agd[_fgd.ObjectNumber].(*_age.PdfIndirectObject)
		if _bde {
			return _cga.PdfObject
		}
		_cd.Log.Debug("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		return nil
	case *_age.PdfIndirectObject:
		return _fgd.PdfObject
	}
	return _accd
}
func (_gfd *fdfParser) parseIndirectObject() (_age.PdfObject, error) {
	_dbb := _age.PdfIndirectObject{}
	_cd.Log.Trace("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a")
	_bebd, _dbc := _gfd._dge.Peek(20)
	if _dbc != nil {
		_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
		return &_dbb, _dbc
	}
	_cd.Log.Trace("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022", string(_bebd))
	_ffad := _eed.FindStringSubmatchIndex(string(_bebd))
	if len(_ffad) < 6 {
		_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_bebd))
		return &_dbb, _c.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_gfd._dge.Discard(_ffad[0])
	_cd.Log.Trace("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064", _ffad)
	_agdg := _ffad[1] - _ffad[0]
	_aaf := make([]byte, _agdg)
	_, _dbc = _gfd.readAtLeast(_aaf, _agdg)
	if _dbc != nil {
		_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073", _dbc)
		return nil, _dbc
	}
	_cd.Log.Trace("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073", _aaf)
	_daee := _eed.FindStringSubmatch(string(_aaf))
	if len(_daee) < 3 {
		_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_aaf))
		return &_dbb, _c.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_gae, _ := _ag.Atoi(_daee[1])
	_ddd, _ := _ag.Atoi(_daee[2])
	_dbb.ObjectNumber = int64(_gae)
	_dbb.GenerationNumber = int64(_ddd)
	for {
		_bcg, _aag := _gfd._dge.Peek(2)
		if _aag != nil {
			return &_dbb, _aag
		}
		_cd.Log.Trace("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_bcg), string(_bcg))
		if _age.IsWhiteSpace(_bcg[0]) {
			_gfd.skipSpaces()
		} else if _bcg[0] == '%' {
			_gfd.skipComments()
		} else if (_bcg[0] == '<') && (_bcg[1] == '<') {
			_cd.Log.Trace("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074")
			_dbb.PdfObject, _aag = _gfd.parseDict()
			_cd.Log.Trace("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076", _aag)
			if _aag != nil {
				return &_dbb, _aag
			}
			_cd.Log.Trace("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		} else if (_bcg[0] == '/') || (_bcg[0] == '(') || (_bcg[0] == '[') || (_bcg[0] == '<') {
			_dbb.PdfObject, _aag = _gfd.parseObject()
			if _aag != nil {
				return &_dbb, _aag
			}
			_cd.Log.Trace("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e")
		} else {
			if _bcg[0] == 'e' {
				_cgbb, _dgbeb := _gfd.readTextLine()
				if _dgbeb != nil {
					return nil, _dgbeb
				}
				if len(_cgbb) >= 6 && _cgbb[0:6] == "\u0065\u006e\u0064\u006f\u0062\u006a" {
					break
				}
			} else if _bcg[0] == 's' {
				_bcg, _ = _gfd._dge.Peek(10)
				if string(_bcg[:6]) == "\u0073\u0074\u0072\u0065\u0061\u006d" {
					_fggg := 6
					if len(_bcg) > 6 {
						if _age.IsWhiteSpace(_bcg[_fggg]) && _bcg[_fggg] != '\r' && _bcg[_fggg] != '\n' {
							_cd.Log.Debug("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072")
							_fggg++
						}
						if _bcg[_fggg] == '\r' {
							_fggg++
							if _bcg[_fggg] == '\n' {
								_fggg++
							}
						} else if _bcg[_fggg] == '\n' {
							_fggg++
						}
					}
					_gfd._dge.Discard(_fggg)
					_bfbf, _dad := _dbb.PdfObject.(*_age.PdfObjectDictionary)
					if !_dad {
						return nil, _c.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
					}
					_cd.Log.Trace("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073", _bfbf)
					_ade, _cdg := _bfbf.Get("\u004c\u0065\u006e\u0067\u0074\u0068").(*_age.PdfObjectInteger)
					if !_cdg {
						return nil, _c.New("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072")
					}
					_dca := *_ade
					if _dca < 0 {
						return nil, _c.New("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
					}
					if int64(_dca) > _gfd._cbb {
						_cd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065")
						return nil, _c.New("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065")
					}
					_fgb := make([]byte, _dca)
					_, _aag = _gfd.readAtLeast(_fgb, int(_dca))
					if _aag != nil {
						_cd.Log.Debug("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058", len(_fgb), _fgb)
						_cd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aag)
						return nil, _aag
					}
					_fffd := _age.PdfObjectStream{}
					_fffd.Stream = _fgb
					_fffd.PdfObjectDictionary = _dbb.PdfObject.(*_age.PdfObjectDictionary)
					_fffd.ObjectNumber = _dbb.ObjectNumber
					_fffd.GenerationNumber = _dbb.GenerationNumber
					_gfd.skipSpaces()
					_gfd._dge.Discard(9)
					_gfd.skipSpaces()
					return &_fffd, nil
				}
			}
			_dbb.PdfObject, _aag = _gfd.parseObject()
			return &_dbb, _aag
		}
	}
	_cd.Log.Trace("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021")
	return &_dbb, nil
}
func (_daa *fdfParser) skipSpaces() (int, error) {
	_ec := 0
	for {
		_eb, _gd := _daa._dge.ReadByte()
		if _gd != nil {
			return 0, _gd
		}
		if _age.IsWhiteSpace(_eb) {
			_ec++
		} else {
			_daa._dge.UnreadByte()
			break
		}
	}
	return _ec, nil
}

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data) FieldValues() (map[string]_age.PdfObject, error) {
	_bb, _cce := fdf.FieldDictionaries()
	if _cce != nil {
		return nil, _cce
	}
	var _cb []string
	for _gbc := range _bb {
		_cb = append(_cb, _gbc)
	}
	_d.Strings(_cb)
	_egc := map[string]_age.PdfObject{}
	for _, _fe := range _cb {
		_ede := _bb[_fe]
		_bec := _age.TraceToDirectObject(_ede.Get("\u0056"))
		_egc[_fe] = _bec
	}
	return _egc, nil
}
func (_fd *fdfParser) parseName() (_age.PdfObjectName, error) {
	var _dga _aa.Buffer
	_abb := false
	for {
		_df, _acg := _fd._dge.Peek(1)
		if _acg == _e.EOF {
			break
		}
		if _acg != nil {
			return _age.PdfObjectName(_dga.String()), _acg
		}
		if !_abb {
			if _df[0] == '/' {
				_abb = true
				_fd._dge.ReadByte()
			} else if _df[0] == '%' {
				_fd.readComment()
				_fd.skipSpaces()
			} else {
				_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029", _df, _df)
				return _age.PdfObjectName(_dga.String()), _be.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _df[0])
			}
		} else {
			if _age.IsWhiteSpace(_df[0]) {
				break
			} else if (_df[0] == '/') || (_df[0] == '[') || (_df[0] == '(') || (_df[0] == ']') || (_df[0] == '<') || (_df[0] == '>') {
				break
			} else if _df[0] == '#' {
				_dcg, _fgf := _fd._dge.Peek(3)
				if _fgf != nil {
					return _age.PdfObjectName(_dga.String()), _fgf
				}
				_fd._dge.Discard(3)
				_daab, _fgf := _f.DecodeString(string(_dcg[1:3]))
				if _fgf != nil {
					return _age.PdfObjectName(_dga.String()), _fgf
				}
				_dga.Write(_daab)
			} else {
				_dab, _ := _fd._dge.ReadByte()
				_dga.WriteByte(_dab)
			}
		}
	}
	return _age.PdfObjectName(_dga.String()), nil
}

// Root returns the Root of the FDF document.
func (_bega *fdfParser) Root() (*_age.PdfObjectDictionary, error) {
	if _bega._ae != nil {
		if _gafc, _acc := _bega.trace(_bega._ae.Get("\u0052\u006f\u006f\u0074")).(*_age.PdfObjectDictionary); _acc {
			if _eae, _afa := _bega.trace(_gafc.Get("\u0046\u0044\u0046")).(*_age.PdfObjectDictionary); _afa {
				return _eae, nil
			}
		}
	}
	var _agb []int64
	for _cbaa := range _bega._agd {
		_agb = append(_agb, _cbaa)
	}
	_d.Slice(_agb, func(_dbe, _edb int) bool { return _agb[_dbe] < _agb[_edb] })
	for _, _cf := range _agb {
		_bbd := _bega._agd[_cf]
		if _cdad, _bgaa := _bega.trace(_bbd).(*_age.PdfObjectDictionary); _bgaa {
			if _cbdf, _fcda := _bega.trace(_cdad.Get("\u0046\u0044\u0046")).(*_age.PdfObjectDictionary); _fcda {
				return _cbdf, nil
			}
		}
	}
	return nil, _c.New("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

// Load loads FDF form data from `r`.
func Load(r _e.ReadSeeker) (*Data, error) {
	_ea, _cc := _bed(r)
	if _cc != nil {
		return nil, _cc
	}
	_ab, _cc := _ea.Root()
	if _cc != nil {
		return nil, _cc
	}
	_abe, _gb := _age.GetArray(_ab.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_gb {
		return nil, _c.New("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	return &Data{_ed: _abe, _cda: _ab}, nil
}

var _gg = _b.MustCompile("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029")

func (_gfb *fdfParser) parseHexString() (*_age.PdfObjectString, error) {
	_gfb._dge.ReadByte()
	var _cdec _aa.Buffer
	for {
		_afg, _ded := _gfb._dge.Peek(1)
		if _ded != nil {
			return _age.MakeHexString(""), _ded
		}
		if _afg[0] == '>' {
			_gfb._dge.ReadByte()
			break
		}
		_bdf, _ := _gfb._dge.ReadByte()
		if !_age.IsWhiteSpace(_bdf) {
			_cdec.WriteByte(_bdf)
		}
	}
	if _cdec.Len()%2 == 1 {
		_cdec.WriteRune('0')
	}
	_dae, _bac := _f.DecodeString(_cdec.String())
	if _bac != nil {
		_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067", _cdec.String())
		return _age.MakeHexString(""), nil
	}
	return _age.MakeHexString(string(_dae)), nil
}
func (_bbf *fdfParser) parseString() (*_age.PdfObjectString, error) {
	_bbf._dge.ReadByte()
	var _ead _aa.Buffer
	_ceg := 1
	for {
		_agc, _gf := _bbf._dge.Peek(1)
		if _gf != nil {
			return _age.MakeString(_ead.String()), _gf
		}
		if _agc[0] == '\\' {
			_bbf._dge.ReadByte()
			_adgd, _acd := _bbf._dge.ReadByte()
			if _acd != nil {
				return _age.MakeString(_ead.String()), _acd
			}
			if _age.IsOctalDigit(_adgd) {
				_gge, _bga := _bbf._dge.Peek(2)
				if _bga != nil {
					return _age.MakeString(_ead.String()), _bga
				}
				var _af []byte
				_af = append(_af, _adgd)
				for _, _eca := range _gge {
					if _age.IsOctalDigit(_eca) {
						_af = append(_af, _eca)
					} else {
						break
					}
				}
				_bbf._dge.Discard(len(_af) - 1)
				_cd.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _af)
				_ffb, _bga := _ag.ParseUint(string(_af), 8, 32)
				if _bga != nil {
					return _age.MakeString(_ead.String()), _bga
				}
				_ead.WriteByte(byte(_ffb))
				continue
			}
			switch _adgd {
			case 'n':
				_ead.WriteRune('\n')
			case 'r':
				_ead.WriteRune('\r')
			case 't':
				_ead.WriteRune('\t')
			case 'b':
				_ead.WriteRune('\b')
			case 'f':
				_ead.WriteRune('\f')
			case '(':
				_ead.WriteRune('(')
			case ')':
				_ead.WriteRune(')')
			case '\\':
				_ead.WriteRune('\\')
			}
			continue
		} else if _agc[0] == '(' {
			_ceg++
		} else if _agc[0] == ')' {
			_ceg--
			if _ceg == 0 {
				_bbf._dge.ReadByte()
				break
			}
		}
		_ccd, _ := _bbf._dge.ReadByte()
		_ead.WriteByte(_ccd)
	}
	return _age.MakeString(_ead.String()), nil
}

// Data represents forms data format (FDF) file data.
type Data struct {
	_cda *_age.PdfObjectDictionary
	_ed  *_age.PdfObjectArray
}

func (_deb *fdfParser) skipComments() error {
	if _, _eba := _deb.skipSpaces(); _eba != nil {
		return _eba
	}
	_ff := true
	for {
		_egd, _bg := _deb._dge.Peek(1)
		if _bg != nil {
			_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _bg.Error())
			return _bg
		}
		if _ff && _egd[0] != '%' {
			return nil
		}
		_ff = false
		if (_egd[0] != '\r') && (_egd[0] != '\n') {
			_deb._dge.ReadByte()
		} else {
			break
		}
	}
	return _deb.skipComments()
}
func (_gag *fdfParser) setFileOffset(_ge int64) {
	_gag._gaf.Seek(_ge, _e.SeekStart)
	_gag._dge = _ca.NewReader(_gag._gaf)
}

var _bd = _b.MustCompile("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029")
var _eed = _b.MustCompile("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a")

func (_bcf *fdfParser) parseDict() (*_age.PdfObjectDictionary, error) {
	_cd.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_fgg := _age.MakeDict()
	_ecc, _ := _bcf._dge.ReadByte()
	if _ecc != '<' {
		return nil, _c.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_ecc, _ = _bcf._dge.ReadByte()
	if _ecc != '<' {
		return nil, _c.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_bcf.skipSpaces()
		_bcf.skipComments()
		_fa, _gdgb := _bcf._dge.Peek(2)
		if _gdgb != nil {
			return nil, _gdgb
		}
		_cd.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_fa), string(_fa))
		if (_fa[0] == '>') && (_fa[1] == '>') {
			_cd.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_bcf._dge.ReadByte()
			_bcf._dge.ReadByte()
			break
		}
		_cd.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_dec, _gdgb := _bcf.parseName()
		_cd.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _dec)
		if _gdgb != nil {
			_cd.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _gdgb)
			return nil, _gdgb
		}
		if len(_dec) > 4 && _dec[len(_dec)-4:] == "\u006e\u0075\u006c\u006c" {
			_cadb := _dec[0 : len(_dec)-4]
			_cd.Log.Debug("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _dec)
			_cd.Log.Debug("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _cadb)
			_bcf.skipSpaces()
			_dfa, _ := _bcf._dge.Peek(1)
			if _dfa[0] == '/' {
				_fgg.Set(_cadb, _age.MakeNull())
				continue
			}
		}
		_bcf.skipSpaces()
		_egf, _gdgb := _bcf.parseObject()
		if _gdgb != nil {
			return nil, _gdgb
		}
		_fgg.Set(_dec, _egf)
		_cd.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _dec, _egf.String())
	}
	_cd.Log.Trace("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021")
	return _fgg, nil
}
func (_cgd *fdfParser) seekToEOFMarker(_ffg int64) error {
	_fga := int64(0)
	_dee := int64(1000)
	for _fga < _ffg {
		if _ffg <= (_dee + _fga) {
			_dee = _ffg - _fga
		}
		_, _ccdd := _cgd._gaf.Seek(-_fga-_dee, _e.SeekEnd)
		if _ccdd != nil {
			return _ccdd
		}
		_edef := make([]byte, _dee)
		_cgd._gaf.Read(_edef)
		_cd.Log.Trace("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022", string(_edef))
		_fdg := _bc.FindAllStringIndex(string(_edef), -1)
		if _fdg != nil {
			_dfaf := _fdg[len(_fdg)-1]
			_cd.Log.Trace("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064", _fdg)
			_cgd._gaf.Seek(-_fga-_dee+int64(_dfaf[0]), _e.SeekEnd)
			return nil
		}
		_cd.Log.Debug("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067")
		_fga += _dee
	}
	_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	return _c.New("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

var _gaa = _b.MustCompile("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052")

func (_ccf *fdfParser) getFileOffset() int64 {
	_de, _ := _ccf._gaf.Seek(0, _e.SeekCurrent)
	_de -= int64(_ccf._dge.Buffered())
	return _de
}
func (_acb *fdfParser) parse() error {
	_acb._gaf.Seek(0, _e.SeekStart)
	_acb._dge = _ca.NewReader(_acb._gaf)
	for {
		_acb.skipComments()
		_fdga, _ccb := _acb._dge.Peek(20)
		if _ccb != nil {
			_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a")
			return _ccb
		}
		if _a.HasPrefix(string(_fdga), "\u0074r\u0061\u0069\u006c\u0065\u0072") {
			_acb._dge.Discard(7)
			_acb.skipSpaces()
			_acb.skipComments()
			_fdeb, _ := _acb.parseDict()
			_acb._ae = _fdeb
			break
		}
		_bfc := _eed.FindStringSubmatchIndex(string(_fdga))
		if len(_bfc) < 6 {
			_cd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029", string(_fdga))
			return _c.New("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
		}
		_eecc, _ccb := _acb.parseIndirectObject()
		if _ccb != nil {
			return _ccb
		}
		switch _aecd := _eecc.(type) {
		case *_age.PdfIndirectObject:
			_acb._agd[_aecd.ObjectNumber] = _aecd
		case *_age.PdfObjectStream:
			_acb._agd[_aecd.ObjectNumber] = _aecd
		default:
			return _c.New("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
	}
	return nil
}
func (_fg *fdfParser) readComment() (string, error) {
	var _aaa _aa.Buffer
	_, _bdg := _fg.skipSpaces()
	if _bdg != nil {
		return _aaa.String(), _bdg
	}
	_bad := true
	for {
		_cgf, _beb := _fg._dge.Peek(1)
		if _beb != nil {
			_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _beb.Error())
			return _aaa.String(), _beb
		}
		if _bad && _cgf[0] != '%' {
			return _aaa.String(), _c.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
		}
		_bad = false
		if (_cgf[0] != '\r') && (_cgf[0] != '\n') {
			_caf, _ := _fg._dge.ReadByte()
			_aaa.WriteByte(_caf)
		} else {
			break
		}
	}
	return _aaa.String(), nil
}
func (_fcb *fdfParser) readAtLeast(_ad []byte, _aga int) (int, error) {
	_ce := _aga
	_cg := 0
	_adf := 0
	for _ce > 0 {
		_egb, _cgb := _fcb._dge.Read(_ad[_cg:])
		if _cgb != nil {
			_cd.Log.Debug("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073", _egb, _adf, _cgb.Error())
			return _cg, _c.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067")
		}
		_adf++
		_cg += _egb
		_ce -= _egb
	}
	return _cg, nil
}
func _cee(_ffbb string) (*fdfParser, error) {
	_fdf := fdfParser{}
	_bdgd := []byte(_ffbb)
	_bce := _aa.NewReader(_bdgd)
	_fdf._gaf = _bce
	_fdf._agd = map[int64]_age.PdfObject{}
	_decd := _ca.NewReader(_bce)
	_fdf._dge = _decd
	_fdf._cbb = int64(len(_ffbb))
	return &_fdf, _fdf.parse()
}
func (_dd *fdfParser) parseNull() (_age.PdfObjectNull, error) {
	_, _bf := _dd._dge.Discard(4)
	return _age.PdfObjectNull{}, _bf
}
func (_beg *fdfParser) parseNumber() (_age.PdfObject, error) { return _age.ParseNumber(_beg._dge) }
func (_bfb *fdfParser) parseFdfVersion() (int, int, error) {
	_bfb._gaf.Seek(0, _e.SeekStart)
	_aca := 20
	_fb := make([]byte, _aca)
	_bfb._gaf.Read(_fb)
	_aef := _dc.FindStringSubmatch(string(_fb))
	if len(_aef) < 3 {
		_ebcf, _fde, _dag := _bfb.seekFdfVersionTopDown()
		if _dag != nil {
			_cd.Log.Debug("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e")
			return 0, 0, _dag
		}
		return _ebcf, _fde, nil
	}
	_gc, _edg := _ag.Atoi(_aef[1])
	if _edg != nil {
		return 0, 0, _edg
	}
	_becc, _edg := _ag.Atoi(_aef[2])
	if _edg != nil {
		return 0, 0, _edg
	}
	_cd.Log.Debug("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064", _gc, _becc)
	return _gc, _becc, nil
}

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data) FieldDictionaries() (map[string]*_age.PdfObjectDictionary, error) {
	_fc := map[string]*_age.PdfObjectDictionary{}
	for _fcc := 0; _fcc < fdf._ed.Len(); _fcc++ {
		_cad, _dg := _age.GetDict(fdf._ed.Get(_fcc))
		if _dg {
			_ee, _ := _age.GetString(_cad.Get("\u0054"))
			if _ee != nil {
				_fc[_ee.Str()] = _cad
			}
		}
	}
	return _fc, nil
}

type fdfParser struct {
	_cef int
	_db  int
	_agd map[int64]_age.PdfObject
	_gaf _e.ReadSeeker
	_dge *_ca.Reader
	_cbb int64
	_ae  *_age.PdfObjectDictionary
}

func _bed(_bfe _e.ReadSeeker) (*fdfParser, error) {
	_feb := &fdfParser{}
	_feb._gaf = _bfe
	_feb._agd = map[int64]_age.PdfObject{}
	_caga, _bcd, _fbb := _feb.parseFdfVersion()
	if _fbb != nil {
		_cd.Log.Error("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076", _fbb)
		return nil, _fbb
	}
	_feb._cef = _caga
	_feb._db = _bcd
	_fbb = _feb.parse()
	return _feb, _fbb
}

var _dc = _b.MustCompile("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029")
