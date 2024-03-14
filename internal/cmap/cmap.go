package cmap

import (
	_e "bufio"
	_bb "bytes"
	_d "encoding/hex"
	_cb "errors"
	_be "fmt"
	_c "io"
	_bc "sort"
	_ba "strconv"
	_eg "strings"
	_bg "unicode/utf16"

	_da "bitbucket.org/shenghui0779/gopdf/common"
	_bce "bitbucket.org/shenghui0779/gopdf/core"
	_g "bitbucket.org/shenghui0779/gopdf/internal/cmap/bcmaps"
)

func LoadPredefinedCMap(name string) (*CMap, error) {
	cmap, _fc := _abe(name)
	if _fc != nil {
		return nil, _fc
	}
	if cmap._f == "" {
		cmap.computeInverseMappings()
		return cmap, nil
	}
	_bed, _fc := _abe(cmap._f)
	if _fc != nil {
		return nil, _fc
	}
	for _fdd, _ab := range _bed._ff {
		if _, _ecg := cmap._ff[_fdd]; !_ecg {
			cmap._ff[_fdd] = _ab
		}
	}
	cmap._ec = append(cmap._ec, _bed._ec...)
	cmap.computeInverseMappings()
	return cmap, nil
}

func (cmap *CMap) parseCodespaceRange() error {
	for {
		_aaec, _dbf := cmap.parseObject()
		if _dbf != nil {
			if _dbf == _c.EOF {
				break
			}
			return _dbf
		}
		_aag, _fdef := _aaec.(cmapHexString)
		if !_fdef {
			if _dcda, _gaa := _aaec.(cmapOperand); _gaa {
				if _dcda.Operand == _ecbc {
					return nil
				}
				return _cb.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
			}
		}
		_aaec, _dbf = cmap.parseObject()
		if _dbf != nil {
			if _dbf == _c.EOF {
				break
			}
			return _dbf
		}
		_aege, _fdef := _aaec.(cmapHexString)
		if !_fdef {
			return _cb.New("\u006e\u006f\u006e-\u0068\u0065\u0078\u0020\u0068\u0069\u0067\u0068")
		}
		if len(_aag._eeea) != len(_aege._eeea) {
			return _cb.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_dbd := _ebb(_aag)
		_dagc := _ebb(_aege)
		if _dagc < _dbd {
			_da.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0061d\u0020\u0063\u006fd\u0065\u0073\u0070\u0061\u0063\u0065\u002e\u0020\u006cow\u003d\u0030\u0078%\u0030\u0032x\u0020\u0068\u0069\u0067\u0068\u003d0\u0078\u00250\u0032\u0078", _dbd, _dagc)
			return ErrBadCMap
		}
		_acg := _aege._ccgc
		_feb := Codespace{NumBytes: _acg, Low: _dbd, High: _dagc}
		cmap._ec = append(cmap._ec, _feb)
		_da.Log.Trace("\u0043\u006f\u0064e\u0073\u0070\u0061\u0063e\u0020\u006c\u006f\u0077\u003a\u0020\u0030x\u0025\u0058\u002c\u0020\u0068\u0069\u0067\u0068\u003a\u0020\u0030\u0078\u0025\u0058", _dbd, _dagc)
	}
	if len(cmap._ec) == 0 {
		_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u0020\u0069\u006e\u0020\u0063ma\u0070\u002e")
		return ErrBadCMap
	}
	return nil
}

func (_bbge *cMapParser) parseComment() (string, error) {
	var _cad _bb.Buffer
	_, _gag := _bbge.skipSpaces()
	if _gag != nil {
		return _cad.String(), _gag
	}
	_fbf := true
	for {
		_eacc, _abd := _bbge._fbeb.Peek(1)
		if _abd != nil {
			_da.Log.Debug("p\u0061r\u0073\u0065\u0043\u006f\u006d\u006d\u0065\u006et\u003a\u0020\u0065\u0072r=\u0025\u0076", _abd)
			return _cad.String(), _abd
		}
		if _fbf && _eacc[0] != '%' {
			return _cad.String(), ErrBadCMapComment
		}
		_fbf = false
		if (_eacc[0] != '\r') && (_eacc[0] != '\n') {
			_cfac, _ := _bbge._fbeb.ReadByte()
			_cad.WriteByte(_cfac)
		} else {
			break
		}
	}
	return _cad.String(), nil
}

func _abe(_bbg string) (*CMap, error) {
	_gg, _aae := _g.Asset(_bbg)
	if _aae != nil {
		return nil, _aae
	}
	return LoadCmapFromDataCID(_gg)
}

type (
	cMapParser struct{ _fbeb *_e.Reader }
	fbRange    struct {
		_cd CharCode
		_bf CharCode
		_a  string
	}
)

func (cmap *CMap) String() string {
	_fgb := cmap._dfd
	_cae := []string{_be.Sprintf("\u006e\u0062\u0069\u0074\u0073\u003a\u0025\u0064", cmap._dc), _be.Sprintf("\u0074y\u0070\u0065\u003a\u0025\u0064", cmap._bd)}
	if cmap._gc != "" {
		_cae = append(_cae, _be.Sprintf("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0025\u0073", cmap._gc))
	}
	if cmap._f != "" {
		_cae = append(_cae, _be.Sprintf("u\u0073\u0065\u0063\u006d\u0061\u0070\u003a\u0025\u0023\u0071", cmap._f))
	}
	_cae = append(_cae, _be.Sprintf("\u0073\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u003a\u0025\u0073", _fgb.String()))
	if len(cmap._ec) > 0 {
		_cae = append(_cae, _be.Sprintf("\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u003a\u0025\u0064", len(cmap._ec)))
	}
	if len(cmap._ce) > 0 {
		_cae = append(_cae, _be.Sprintf("\u0063\u006fd\u0065\u0054\u006fU\u006e\u0069\u0063\u006f\u0064\u0065\u003a\u0025\u0064", len(cmap._ce)))
	}
	return _be.Sprintf("\u0043\u004d\u0041P\u007b\u0025\u0023\u0071\u0020\u0025\u0073\u007d", cmap._cdb, _eg.Join(_cae, "\u0020"))
}

func (_bbbb *cMapParser) parseName() (cmapName, error) {
	_ffeb := ""
	_fec := false
	for {
		_ggg, _dcc := _bbbb._fbeb.Peek(1)
		if _dcc == _c.EOF {
			break
		}
		if _dcc != nil {
			return cmapName{_ffeb}, _dcc
		}
		if !_fec {
			if _ggg[0] == '/' {
				_fec = true
				_bbbb._fbeb.ReadByte()
			} else {
				_da.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u004e\u0061\u006d\u0065\u0020\u0073\u0074a\u0072t\u0069n\u0067 \u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0028\u0025\u0020\u0078\u0029", _ggg, _ggg)
				return cmapName{_ffeb}, _be.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _ggg[0])
			}
		} else {
			if _bce.IsWhiteSpace(_ggg[0]) {
				break
			} else if (_ggg[0] == '/') || (_ggg[0] == '[') || (_ggg[0] == '(') || (_ggg[0] == ']') || (_ggg[0] == '<') || (_ggg[0] == '>') {
				break
			} else if _ggg[0] == '#' {
				_gbb, _aba := _bbbb._fbeb.Peek(3)
				if _aba != nil {
					return cmapName{_ffeb}, _aba
				}
				_bbbb._fbeb.Discard(3)
				_dbca, _aba := _d.DecodeString(string(_gbb[1:3]))
				if _aba != nil {
					return cmapName{_ffeb}, _aba
				}
				_ffeb += string(_dbca)
			} else {
				_agf, _ := _bbbb._fbeb.ReadByte()
				_ffeb += string(_agf)
			}
		}
	}
	return cmapName{_ffeb}, nil
}

func (_adc *cMapParser) skipSpaces() (int, error) {
	_age := 0
	for {
		_faa, _cde := _adc._fbeb.Peek(1)
		if _cde != nil {
			return 0, _cde
		}
		if _bce.IsWhiteSpace(_faa[0]) {
			_adc._fbeb.ReadByte()
			_age++
		} else {
			break
		}
	}
	return _age, nil
}

type cmapName struct{ Name string }

func (cmap *CMap) parse() error {
	var _gac cmapObject
	for {
		_cddc, _eaeb := cmap.parseObject()
		if _eaeb != nil {
			if _eaeb == _c.EOF {
				break
			}
			_da.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u004d\u0061\u0070\u003a\u0020\u0025\u0076", _eaeb)
			return _eaeb
		}
		switch _dbe := _cddc.(type) {
		case cmapOperand:
			_cfgf := _dbe
			switch _cfgf.Operand {
			case _babb:
				_caff := cmap.parseCodespaceRange()
				if _caff != nil {
					return _caff
				}
			case _bcab:
				_cge := cmap.parseCIDRange()
				if _cge != nil {
					return _cge
				}
			case _ee:
				_dgb := cmap.parseBfchar()
				if _dgb != nil {
					return _dgb
				}
			case _eee:
				_efe := cmap.parseBfrange()
				if _efe != nil {
					return _efe
				}
			case _cab:
				if _gac == nil {
					_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u0073\u0065\u0063m\u0061\u0070\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f \u0061\u0072\u0067")
					return ErrBadCMap
				}
				_cgg, _deg := _gac.(cmapName)
				if !_deg {
					_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0075\u0073\u0065\u0063\u006d\u0061\u0070\u0020\u0061\u0072\u0067\u0020\u006eo\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u0020\u0025\u0023\u0076", _gac)
					return ErrBadCMap
				}
				cmap._f = _cgg.Name
			case _bdb:
				_daf := cmap.parseSystemInfo()
				if _daf != nil {
					return _daf
				}
			}
		case cmapName:
			_cea := _dbe
			switch _cea.Name {
			case _bdb:
				_dee := cmap.parseSystemInfo()
				if _dee != nil {
					return _dee
				}
			case _gfbb:
				_ad := cmap.parseName()
				if _ad != nil {
					return _ad
				}
			case _dcgc:
				_dcd := cmap.parseType()
				if _dcd != nil {
					return _dcd
				}
			case _gcgd:
				_ddc := cmap.parseVersion()
				if _ddc != nil {
					return _ddc
				}
			case _acbe:
				if _eaeb = cmap.parseWMode(); _eaeb != nil {
					return _eaeb
				}
			}
		}
		_gac = _cddc
	}
	return nil
}

func LoadCmapFromData(data []byte, isSimple bool) (*CMap, error) {
	_da.Log.Trace("\u004c\u006fa\u0064\u0043\u006d\u0061\u0070\u0046\u0072\u006f\u006d\u0044\u0061\u0074\u0061\u003a\u0020\u0069\u0073\u0053\u0069\u006d\u0070\u006ce=\u0025\u0074", isSimple)
	cmap := _fbg(isSimple)
	cmap.cMapParser = _aab(data)
	_de := cmap.parse()
	if _de != nil {
		return nil, _de
	}
	if len(cmap._ec) == 0 {
		if cmap._f != "" {
			return cmap, nil
		}
		_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0063\u006d\u0061p=\u0025\u0073", cmap)
	}
	cmap.computeInverseMappings()
	return cmap, nil
}

type CIDSystemInfo struct {
	Registry   string
	Ordering   string
	Supplement int
}
type cmapOperand struct{ Operand string }

func (cmap *CMap) CharcodeBytesToUnicode(data []byte) (string, int) {
	_cf, _cg := cmap.BytesToCharcodes(data)
	if !_cg {
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065s\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069n\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0064\u0061\u0074\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078]\u0020\u0063\u006d\u0061\u0070=\u0025\u0073", data, cmap)
		return "", 0
	}
	_cfg := make([]string, len(_cf))
	var _ggb []CharCode
	for _egb, _dg := range _cf {
		_fg, _bae := cmap._ce[_dg]
		if !_bae {
			_ggb = append(_ggb, _dg)
			_fg = MissingCodeString
		}
		_cfg[_egb] = _fg
	}
	_ef := _eg.Join(_cfg, "")
	if len(_ggb) > 0 {
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0068\u0061\u0072c\u006f\u0064\u0065\u0042y\u0074\u0065\u0073\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070\u002e\u000a"+"\u0009d\u0061t\u0061\u003d\u005b\u0025\u00200\u0032\u0078]\u003d\u0025\u0023\u0071\u000a"+"\u0009\u0063h\u0061\u0072\u0063o\u0064\u0065\u0073\u003d\u0025\u0030\u0032\u0078\u000a"+"\u0009\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003d\u0025\u0064\u0020%\u0030\u0032\u0078\u000a"+"\u0009\u0075\u006e\u0069\u0063\u006f\u0064\u0065\u003d`\u0025\u0073\u0060\u000a"+"\u0009\u0063\u006d\u0061\u0070\u003d\u0025\u0073", data, string(data), _cf, len(_ggb), _ggb, _ef, cmap)
	}
	return _ef, len(_ggb)
}
func (cmap *CMap) Type() int { return cmap._bd }
func (cmap *CMap) matchCode(_abb []byte) (_eaed CharCode, _cee int, _fe bool) {
	for _cdcg := 0; _cdcg < _egc; _cdcg++ {
		if _cdcg < len(_abb) {
			_eaed = _eaed<<8 | CharCode(_abb[_cdcg])
			_cee++
		}
		_fe = cmap.inCodespace(_eaed, _cdcg+1)
		if _fe {
			return _eaed, _cee, true
		}
	}
	_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063o\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0020m\u0061t\u0063\u0068\u0065\u0073\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d=\u0025\u0023\u0071\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", _abb, string(_abb), cmap)
	return 0, 0, false
}
func (cmap *CMap) NBits() int { return cmap._dc }

type cmapFloat struct{ _fdegf float64 }

func (_cdee *cMapParser) parseOperand() (cmapOperand, error) {
	_dga := cmapOperand{}
	_fbea := _bb.Buffer{}
	for {
		_bacd, _bceg := _cdee._fbeb.Peek(1)
		if _bceg != nil {
			if _bceg == _c.EOF {
				break
			}
			return _dga, _bceg
		}
		if _bce.IsDelimiter(_bacd[0]) {
			break
		}
		if _bce.IsWhiteSpace(_bacd[0]) {
			break
		}
		_befb, _ := _cdee._fbeb.ReadByte()
		_fbea.WriteByte(_befb)
	}
	if _fbea.Len() == 0 {
		return _dga, _be.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	_dga.Operand = _fbea.String()
	return _dga, nil
}

type cmapDict struct{ Dict map[string]cmapObject }

const (
	_egc              = 4
	MissingCodeRune   = '\ufffd'
	MissingCodeString = string(MissingCodeRune)
)

func _cfb(_gccf cmapHexString) rune {
	_aggb := _acc(_gccf)
	if _ageb := len(_aggb); _ageb == 0 {
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054o\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065c\u0074\u0065\u0064\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0072u\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023\u0076", _gccf)
		return MissingCodeRune
	}
	if len(_aggb) > 1 {
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0078\u0061\u0063\u0074\u006c\u0079\u0020\u006f\u006e\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023v\u0020\u002d\u003e\u0020\u0025#\u0076", _gccf, _aggb)
	}
	return _aggb[0]
}

func _ebb(_geba cmapHexString) CharCode {
	_edgga := CharCode(0)
	for _, _dccc := range _geba._eeea {
		_edgga <<= 8
		_edgga |= CharCode(_dccc)
	}
	return _edgga
}

const (
	_bdb  = "\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"
	_cdda = "\u0062e\u0067\u0069\u006e\u0063\u006d\u0061p"
	_cdf  = "\u0065n\u0064\u0063\u006d\u0061\u0070"
	_babb = "\u0062\u0065\u0067\u0069nc\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0072\u0061\u006e\u0067\u0065"
	_ecbc = "\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065"
	_ee   = "b\u0065\u0067\u0069\u006e\u0062\u0066\u0063\u0068\u0061\u0072"
	_ecdc = "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r"
	_eee  = "\u0062\u0065\u0067i\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_afb  = "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_bcab = "\u0062\u0065\u0067\u0069\u006e\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_eaec = "e\u006e\u0064\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_cab  = "\u0075s\u0065\u0063\u006d\u0061\u0070"
	_acbe = "\u0057\u004d\u006fd\u0065"
	_gfbb = "\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065"
	_dcgc = "\u0043\u004d\u0061\u0070\u0054\u0079\u0070\u0065"
	_gcgd = "C\u004d\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e"
)

func (cmap *CMap) Bytes() []byte {
	_da.Log.Trace("\u0063\u006d\u0061\u0070.B\u0079\u0074\u0065\u0073\u003a\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", cmap.String())
	if len(cmap._cec) > 0 {
		return cmap._cec
	}
	cmap._cec = []byte(_eg.Join([]string{_ace, cmap.toBfData(), _dad}, "\u000a"))
	return cmap._cec
}

type cmapHexString struct {
	_ccgc int
	_eeea []byte
}

func (cmap *CMap) parseName() error {
	_ecba := ""
	_gefa := false
	for _efd := 0; _efd < 20 && !_gefa; _efd++ {
		_cbd, _cca := cmap.parseObject()
		if _cca != nil {
			return _cca
		}
		switch _gee := _cbd.(type) {
		case cmapOperand:
			switch _gee.Operand {
			case "\u0064\u0065\u0066":
				_gefa = true
			default:
				_da.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a\u0020\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076\u0020n\u0061\u006de\u003d\u0025\u0023\u0071", _cbd, _ecba)
				if _ecba != "" {
					_ecba = _be.Sprintf("\u0025\u0073\u0020%\u0073", _ecba, _gee.Operand)
				}
				_da.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a \u0052\u0065\u0063\u006f\u0076\u0065\u0072e\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025\u0023\u0071", _ecba)
			}
		case cmapName:
			_ecba = _gee.Name
		}
	}
	if !_gefa {
		_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0070\u0061\u0072\u0073\u0065N\u0061m\u0065:\u0020\u004e\u006f\u0020\u0064\u0065\u0066 ")
		return ErrBadCMap
	}
	cmap._cdb = _ecba
	return nil
}

type cmapString struct {
	String string
}

func (_gadf *cMapParser) parseArray() (cmapArray, error) {
	_eeeb := cmapArray{}
	_eeeb.Array = []cmapObject{}
	_gadf._fbeb.ReadByte()
	for {
		_gadf.skipSpaces()
		_dcfc, _ecbg := _gadf._fbeb.Peek(1)
		if _ecbg != nil {
			return _eeeb, _ecbg
		}
		if _dcfc[0] == ']' {
			_gadf._fbeb.ReadByte()
			break
		}
		_cgd, _ecbg := _gadf.parseObject()
		if _ecbg != nil {
			return _eeeb, _ecbg
		}
		_eeeb.Array = append(_eeeb.Array, _cgd)
	}
	return _eeeb, nil
}
func _fa(_fdg string) rune { _acbf := []rune(_fdg); return _acbf[len(_acbf)-1] }
func (cmap *CMap) toBfData() string {
	if len(cmap._ce) == 0 {
		return ""
	}
	_cc := make([]CharCode, 0, len(cmap._ce))
	for _fbca := range cmap._ce {
		_cc = append(_cc, _fbca)
	}
	_bc.Slice(_cc, func(_ddbg, _bbdg int) bool { return _cc[_ddbg] < _cc[_bbdg] })
	var _gbd []charRange
	_bag := charRange{_cc[0], _cc[0]}
	_ddg := cmap._ce[_cc[0]]
	for _, _ceef := range _cc[1:] {
		_ege := cmap._ce[_ceef]
		if _ceef == _bag._gf+1 && _fa(_ege) == _fa(_ddg)+1 {
			_bag._gf = _ceef
		} else {
			_gbd = append(_gbd, _bag)
			_bag._dd, _bag._gf = _ceef, _ceef
		}
		_ddg = _ege
	}
	_gbd = append(_gbd, _bag)
	var _cfa []CharCode
	var _ecb []fbRange
	for _, _gbe := range _gbd {
		if _gbe._dd == _gbe._gf {
			_cfa = append(_cfa, _gbe._dd)
		} else {
			_ecb = append(_ecb, fbRange{_cd: _gbe._dd, _bf: _gbe._gf, _a: cmap._ce[_gbe._dd]})
		}
	}
	_da.Log.Trace("\u0063\u0068ar\u0052\u0061\u006eg\u0065\u0073\u003d\u0025d f\u0062Ch\u0061\u0072\u0073\u003d\u0025\u0064\u0020fb\u0052\u0061\u006e\u0067\u0065\u0073\u003d%\u0064", len(_gbd), len(_cfa), len(_ecb))
	var _acb []string
	if len(_cfa) > 0 {
		_ebc := (len(_cfa) + _bab - 1) / _bab
		for _gcc := 0; _gcc < _ebc; _gcc++ {
			_ceg := _fbab(len(_cfa)-_gcc*_bab, _bab)
			_acb = append(_acb, _be.Sprintf("\u0025\u0064\u0020\u0062\u0065\u0067\u0069\u006e\u0062f\u0063\u0068\u0061\u0072", _ceg))
			for _gef := 0; _gef < _ceg; _gef++ {
				_ebf := _cfa[_gcc*_bab+_gef]
				_dfg := cmap._ce[_ebf]
				_acb = append(_acb, _be.Sprintf("\u003c%\u0030\u0034\u0078\u003e\u0020\u0025s", _ebf, _dfe(_dfg)))
			}
			_acb = append(_acb, "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r")
		}
	}
	if len(_ecb) > 0 {
		_gegf := (len(_ecb) + _bab - 1) / _bab
		for _fgd := 0; _fgd < _gegf; _fgd++ {
			_aeb := _fbab(len(_ecb)-_fgd*_bab, _bab)
			_acb = append(_acb, _be.Sprintf("\u0025d\u0020b\u0065\u0067\u0069\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065", _aeb))
			for _gfb := 0; _gfb < _aeb; _gfb++ {
				_ga := _ecb[_fgd*_bab+_gfb]
				_acb = append(_acb, _be.Sprintf("\u003c%\u00304\u0078\u003e\u003c\u0025\u0030\u0034\u0078\u003e\u0020\u0025\u0073", _ga._cd, _ga._bf, _dfe(_ga._a)))
			}
			_acb = append(_acb, "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065")
		}
	}
	return _eg.Join(_acb, "\u000a")
}
func IsPredefinedCMap(name string) bool                  { return _g.AssetExists(name) }
func (cmap *CMap) StringToCID(s string) (CharCode, bool) { _gb, _fcf := cmap._aa[s]; return _gb, _fcf }
func (cmap *CMap) parseVersion() error {
	_abg := ""
	_fbga := false
	for _egad := 0; _egad < 3 && !_fbga; _egad++ {
		_cdbb, _ecc := cmap.parseObject()
		if _ecc != nil {
			return _ecc
		}
		switch _fde := _cdbb.(type) {
		case cmapOperand:
			switch _fde.Operand {
			case "\u0064\u0065\u0066":
				_fbga = true
			default:
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0070\u0061\u0072\u0073\u0065\u0056e\u0072\u0073\u0069\u006f\u006e\u003a \u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076", _cdbb)
				return ErrBadCMap
			}
		case cmapInt:
			_abg = _be.Sprintf("\u0025\u0064", _fde._abbg)
		case cmapFloat:
			_abg = _be.Sprintf("\u0025\u0066", _fde._fdegf)
		case cmapString:
			_abg = _fde.String
		default:
			_da.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020p\u0061\u0072\u0073\u0065Ver\u0073io\u006e\u003a\u0020\u0042\u0061\u0064\u0020ty\u0070\u0065\u002e\u0020\u006f\u003d\u0025#\u0076", _cdbb)
		}
	}
	cmap._gc = _abg
	return nil
}

func _aab(_fdf []byte) *cMapParser {
	_gae := cMapParser{}
	_fcb := _bb.NewBuffer(_fdf)
	_gae._fbeb = _e.NewReader(_fcb)
	return &_gae
}

type Codespace struct {
	NumBytes int
	Low      CharCode
	High     CharCode
}

func (_fdda *cMapParser) parseNumber() (cmapObject, error) {
	_cbfa, _cdfc := _bce.ParseNumber(_fdda._fbeb)
	if _cdfc != nil {
		return nil, _cdfc
	}
	switch _eebg := _cbfa.(type) {
	case *_bce.PdfObjectFloat:
		return cmapFloat{float64(*_eebg)}, nil
	case *_bce.PdfObjectInteger:
		return cmapInt{int64(*_eebg)}, nil
	}
	return nil, _be.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _cbfa)
}

func (cmap *CMap) computeInverseMappings() {
	for _ddb, _gd := range cmap._ff {
		if _ebd, _dbb := cmap._fb[_gd]; !_dbb || (_dbb && _ebd > _ddb) {
			cmap._fb[_gd] = _ddb
		}
	}
	for _cag, _ed := range cmap._ce {
		if _bfb, _agg := cmap._aa[_ed]; !_agg || (_agg && _bfb > _cag) {
			cmap._aa[_ed] = _cag
		}
	}
	_bc.Slice(cmap._ec, func(_gff, _fbb int) bool { return cmap._ec[_gff].Low < cmap._ec[_fbb].Low })
}

func (cmap *CMap) parseCIDRange() error {
	for {
		_fca, _baeb := cmap.parseObject()
		if _baeb != nil {
			if _baeb == _c.EOF {
				break
			}
			return _baeb
		}
		_acgc, _ccb := _fca.(cmapHexString)
		if !_ccb {
			if _ccbe, _cgb := _fca.(cmapOperand); _cgb {
				if _ccbe.Operand == _eaec {
					return nil
				}
				return _cb.New("\u0063\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0072\u0076\u0061\u006c\u0020s\u0074\u0061\u0072\u0074\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
			}
		}
		_edd := _ebb(_acgc)
		_fca, _baeb = cmap.parseObject()
		if _baeb != nil {
			if _baeb == _c.EOF {
				break
			}
			return _baeb
		}
		_ecga, _ccb := _fca.(cmapHexString)
		if !_ccb {
			return _cb.New("\u0063\u0069d\u0020\u0069\u006e\u0074e\u0072\u0076a\u006c\u0020\u0065\u006e\u0064\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_acgc._eeea) != len(_ecga._eeea) {
			return _cb.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_cbb := _ebb(_ecga)
		if _edd > _cbb {
			_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0049\u0044\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073t\u0061\u0072\u0074\u003d\u0030\u0078\u0025\u0030\u0032\u0078\u0020\u0065\u006e\u0064=\u0030x\u0025\u0030\u0032\u0078", _edd, _cbb)
			return ErrBadCMap
		}
		_fca, _baeb = cmap.parseObject()
		if _baeb != nil {
			if _baeb == _c.EOF {
				break
			}
			return _baeb
		}
		_efb, _ccb := _fca.(cmapInt)
		if !_ccb {
			return _cb.New("\u0063\u0069\u0064\u0020\u0073t\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		if _efb._abbg < 0 {
			return _cb.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0063\u0069\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065")
		}
		_fad := _efb._abbg
		for _cgeg := _edd; _cgeg <= _cbb; _cgeg++ {
			cmap._ff[_cgeg] = CharCode(_fad)
			_fad++
		}
		_da.Log.Trace("C\u0049\u0044\u0020\u0072\u0061\u006eg\u0065\u003a\u0020\u003c\u0030\u0078\u0025\u0058\u003e \u003c\u0030\u0078%\u0058>\u0020\u0025\u0064", _edd, _cbb, _efb._abbg)
	}
	return nil
}
func _dafe() cmapDict { return cmapDict{Dict: map[string]cmapObject{}} }
func (cmap *CMap) parseBfrange() error {
	for {
		var _gfgb CharCode
		_cbfb, _ced := cmap.parseObject()
		if _ced != nil {
			if _ced == _c.EOF {
				break
			}
			return _ced
		}
		switch _ffe := _cbfb.(type) {
		case cmapOperand:
			if _ffe.Operand == _afb {
				return nil
			}
			return _cb.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_gfgb = _ebb(_ffe)
		default:
			return _cb.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		var _daa CharCode
		_cbfb, _ced = cmap.parseObject()
		if _ced != nil {
			if _ced == _c.EOF {
				break
			}
			return _ced
		}
		switch _fbef := _cbfb.(type) {
		case cmapOperand:
			_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065\u0020\u0062\u0066r\u0061\u006e\u0067\u0065\u0020\u0074\u0072i\u0070\u006c\u0065\u0074")
			return ErrBadCMap
		case cmapHexString:
			_daa = _ebb(_fbef)
			if _daa > 0xffff {
				_daa = 0xffff
			}
		default:
			_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _cbfb)
			return ErrBadCMap
		}
		_cbfb, _ced = cmap.parseObject()
		if _ced != nil {
			if _ced == _c.EOF {
				break
			}
			return _ced
		}
		switch _dbc := _cbfb.(type) {
		case cmapArray:
			if len(_dbc.Array) != int(_daa-_gfgb)+1 {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0069\u0074\u0065\u006d\u0073\u0020\u0069\u006e\u0020a\u0072\u0072\u0061\u0079")
				return ErrBadCMap
			}
			for _bge := _gfgb; _bge <= _daa; _bge++ {
				_febc := _dbc.Array[_bge-_gfgb]
				_befd, _ada := _febc.(cmapHexString)
				if !_ada {
					return _cb.New("\u006e\u006f\u006e-h\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				}
				_gcaa := _acc(_befd)
				cmap._ce[_bge] = string(_gcaa)
			}
		case cmapHexString:
			_gbg := _acc(_dbc)
			_eddc := len(_gbg)
			for _ecd := _gfgb; _ecd <= _daa; _ecd++ {
				cmap._ce[_ecd] = string(_gbg)
				if _eddc > 0 {
					_gbg[_eddc-1]++
				} else {
					_da.Log.Debug("\u004e\u006f\u0020c\u006d\u0061\u0070\u0020\u0074\u0061\u0072\u0067\u0065\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0066\u006f\u0072\u0020\u0025\u0023\u0076", _ecd)
				}
				if _ecd == 1<<32-1 {
					break
				}
			}
		default:
			_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _cbfb)
			return ErrBadCMap
		}
	}
	return nil
}

type cmapInt struct{ _abbg int64 }

func (cmap *CMap) inCodespace(_aga CharCode, _bdg int) bool {
	for _, _ecfc := range cmap._ec {
		if _ecfc.Low <= _aga && _aga <= _ecfc.High && _bdg == _ecfc.NumBytes {
			return true
		}
	}
	return false
}

var (
	ErrBadCMap        = _cb.New("\u0062\u0061\u0064\u0020\u0063\u006d\u0061\u0070")
	ErrBadCMapComment = _cb.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
	ErrBadCMapDict    = _cb.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
)

func (cmap *CMap) Stream() (*_bce.PdfObjectStream, error) {
	if cmap._eb != nil {
		return cmap._eb, nil
	}
	_dagf, _efc := _bce.MakeStream(cmap.Bytes(), _bce.NewFlateEncoder())
	if _efc != nil {
		return nil, _efc
	}
	cmap._eb = _dagf
	return cmap._eb, nil
}

type charRange struct {
	_dd CharCode
	_gf CharCode
}

func _fbg(_aea bool) *CMap {
	_db := 16
	if _aea {
		_db = 8
	}
	return &CMap{_dc: _db, _ff: make(map[CharCode]CharCode), _fb: make(map[CharCode]CharCode), _ce: make(map[CharCode]string), _aa: make(map[string]CharCode)}
}

func (cmap *CMap) parseType() error {
	_bde := 0
	_gffa := false
	for _ade := 0; _ade < 3 && !_gffa; _ade++ {
		_cff, _afee := cmap.parseObject()
		if _afee != nil {
			return _afee
		}
		switch _gffad := _cff.(type) {
		case cmapOperand:
			switch _gffad.Operand {
			case "\u0064\u0065\u0066":
				_gffa = true
			default:
				_da.Log.Error("\u0070\u0061r\u0073\u0065\u0054\u0079\u0070\u0065\u003a\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u006f=%\u0023\u0076", _cff)
				return ErrBadCMap
			}
		case cmapInt:
			_bde = int(_gffad._abbg)
		}
	}
	cmap._bd = _bde
	return nil
}

type CMap struct {
	*cMapParser
	_cdb string
	_dc  int
	_bd  int
	_gc  string
	_f   string
	_dfd CIDSystemInfo
	_ec  []Codespace
	_ff  map[CharCode]CharCode
	_fb  map[CharCode]CharCode
	_ce  map[CharCode]string
	_aa  map[string]CharCode
	_cec []byte
	_eb  *_bce.PdfObjectStream
	_bfc integer
}

func (cmap *CMap) Name() string { return cmap._cdb }

type CharCode uint32

func LoadCmapFromDataCID(data []byte) (*CMap, error) { return LoadCmapFromData(data, false) }
func (_cfgg *cMapParser) parseDict() (cmapDict, error) {
	_da.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_cbga := _dafe()
	_ggf, _ := _cfgg._fbeb.ReadByte()
	if _ggf != '<' {
		return _cbga, ErrBadCMapDict
	}
	_ggf, _ = _cfgg._fbeb.ReadByte()
	if _ggf != '<' {
		return _cbga, ErrBadCMapDict
	}
	for {
		_cfgg.skipSpaces()
		_ceda, _caa := _cfgg._fbeb.Peek(2)
		if _caa != nil {
			return _cbga, _caa
		}
		if (_ceda[0] == '>') && (_ceda[1] == '>') {
			_cfgg._fbeb.ReadByte()
			_cfgg._fbeb.ReadByte()
			break
		}
		_eedb, _caa := _cfgg.parseName()
		_da.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _eedb.Name)
		if _caa != nil {
			_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u002e\u0020\u0065\u0072r=\u0025\u0076", _caa)
			return _cbga, _caa
		}
		_cfgg.skipSpaces()
		_ece, _caa := _cfgg.parseObject()
		if _caa != nil {
			return _cbga, _caa
		}
		_cbga.Dict[_eedb.Name] = _ece
		_cfgg.skipSpaces()
		_ceda, _caa = _cfgg._fbeb.Peek(3)
		if _caa != nil {
			return _cbga, _caa
		}
		if string(_ceda) == "\u0064\u0065\u0066" {
			_cfgg._fbeb.Discard(3)
		}
	}
	return _cbga, nil
}

const (
	_bab = 100
	_ace = "\u000a\u002f\u0043\u0049\u0044\u0049\u006e\u0069\u0074\u0020\u002f\u0050\u0072\u006fc\u0053\u0065\u0074\u0020\u0066\u0069\u006e\u0064\u0072es\u006fu\u0072c\u0065 \u0062\u0065\u0067\u0069\u006e\u000a\u0031\u0032\u0020\u0064\u0069\u0063\u0074\u0020\u0062\u0065\u0067\u0069n\u000a\u0062\u0065\u0067\u0069\u006e\u0063\u006d\u0061\u0070\n\u002f\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u003c\u003c\u0020\u002f\u0052\u0065\u0067\u0069\u0073t\u0072\u0079\u0020\u0028\u0041\u0064\u006f\u0062\u0065\u0029\u0020\u002f\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0028\u0055\u0043\u0053)\u0020\u002f\u0053\u0075\u0070p\u006c\u0065\u006d\u0065\u006et\u0020\u0030\u0020\u003e\u003e\u0020\u0064\u0065\u0066\u000a\u002f\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u002f\u0041\u0064\u006f\u0062\u0065-\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0055\u0043\u0053\u0020\u0064\u0065\u0066\u000a\u002fC\u004d\u0061\u0070\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0064\u0065\u0066\u000a\u0031\u0020\u0062\u0065\u0067\u0069\u006e\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063e\u0072\u0061n\u0067\u0065\n\u003c\u0030\u0030\u0030\u0030\u003e\u0020<\u0046\u0046\u0046\u0046\u003e\u000a\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065\u000a"
	_dad = "\u0065\u006e\u0064\u0063\u006d\u0061\u0070\u000a\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u0063ur\u0072e\u006e\u0074\u0064\u0069\u0063\u0074\u0020\u002f\u0043\u004d\u0061\u0070 \u0064\u0065\u0066\u0069\u006e\u0065\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0070\u006fp\u000a\u0065\u006e\u0064\u000a\u0065\u006e\u0064\u000a"
)

func _fbab(_gdff, _dcdc int) int {
	if _gdff < _dcdc {
		return _gdff
	}
	return _dcdc
}

type cmapObject interface{}

func (_dadb *cMapParser) parseString() (cmapString, error) {
	_dadb._fbeb.ReadByte()
	_abeg := _bb.Buffer{}
	_cdcb := 1
	for {
		_fgga, _eed := _dadb._fbeb.Peek(1)
		if _eed != nil {
			return cmapString{_abeg.String()}, _eed
		}
		if _fgga[0] == '\\' {
			_dadb._fbeb.ReadByte()
			_cedb, _aebc := _dadb._fbeb.ReadByte()
			if _aebc != nil {
				return cmapString{_abeg.String()}, _aebc
			}
			if _bce.IsOctalDigit(_cedb) {
				_agea, _ccde := _dadb._fbeb.Peek(2)
				if _ccde != nil {
					return cmapString{_abeg.String()}, _ccde
				}
				var _bffa []byte
				_bffa = append(_bffa, _cedb)
				for _, _aagd := range _agea {
					if _bce.IsOctalDigit(_aagd) {
						_bffa = append(_bffa, _aagd)
					} else {
						break
					}
				}
				_dadb._fbeb.Discard(len(_bffa) - 1)
				_da.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _bffa)
				_gad, _ccde := _ba.ParseUint(string(_bffa), 8, 32)
				if _ccde != nil {
					return cmapString{_abeg.String()}, _ccde
				}
				_abeg.WriteByte(byte(_gad))
				continue
			}
			switch _cedb {
			case 'n':
				_abeg.WriteByte('\n')
			case 'r':
				_abeg.WriteByte('\r')
			case 't':
				_abeg.WriteByte('\t')
			case 'b':
				_abeg.WriteByte('\b')
			case 'f':
				_abeg.WriteByte('\f')
			case '(':
				_abeg.WriteByte('(')
			case ')':
				_abeg.WriteByte(')')
			case '\\':
				_abeg.WriteByte('\\')
			}
			continue
		} else if _fgga[0] == '(' {
			_cdcb++
		} else if _fgga[0] == ')' {
			_cdcb--
			if _cdcb == 0 {
				_dadb._fbeb.ReadByte()
				break
			}
		}
		_cbg, _ := _dadb._fbeb.ReadByte()
		_abeg.WriteByte(_cbg)
	}
	return cmapString{_abeg.String()}, nil
}

func (cmap *CMap) parseBfchar() error {
	for {
		_dcf, _dgf := cmap.parseObject()
		if _dgf != nil {
			if _dgf == _c.EOF {
				break
			}
			return _dgf
		}
		var _cfed CharCode
		switch _fae := _dcf.(type) {
		case cmapOperand:
			if _fae.Operand == _ecdc {
				return nil
			}
			return _cb.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_cfed = _ebb(_fae)
		default:
			return _cb.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_dcf, _dgf = cmap.parseObject()
		if _dgf != nil {
			if _dgf == _c.EOF {
				break
			}
			return _dgf
		}
		var _bdgg []rune
		switch _gfg := _dcf.(type) {
		case cmapOperand:
			if _gfg.Operand == _ecdc {
				return nil
			}
			_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020o\u0070\u0065\u0072\u0061\u006e\u0064\u002e\u0020\u0025\u0023\u0076", _gfg)
			return ErrBadCMap
		case cmapHexString:
			_bdgg = _acc(_gfg)
		case cmapName:
			_da.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0061\u006de\u002e \u0025\u0023\u0076", _gfg)
			_bdgg = []rune{MissingCodeRune}
		default:
			_da.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0074\u0079\u0070e\u002e \u0025\u0023\u0076", _dcf)
			return ErrBadCMap
		}
		cmap._ce[_cfed] = string(_bdgg)
	}
	return nil
}

func (_aeg *CIDSystemInfo) String() string {
	return _be.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0025\u0030\u0033\u0064", _aeg.Registry, _aeg.Ordering, _aeg.Supplement)
}

func NewToUnicodeCMap(codeToRune map[CharCode]rune) *CMap {
	_bda := make(map[CharCode]string, len(codeToRune))
	for _cdd, _bgd := range codeToRune {
		_bda[_cdd] = string(_bgd)
	}
	cmap := &CMap{_cdb: "\u0041d\u006fb\u0065\u002d\u0049\u0064\u0065n\u0074\u0069t\u0079\u002d\u0055\u0043\u0053", _bd: 2, _dc: 16, _dfd: CIDSystemInfo{Registry: "\u0041\u0064\u006fb\u0065", Ordering: "\u0055\u0043\u0053", Supplement: 0}, _ec: []Codespace{{Low: 0, High: 0xffff}}, _ce: _bda, _aa: make(map[string]CharCode, len(codeToRune)), _ff: make(map[CharCode]CharCode, len(codeToRune)), _fb: make(map[CharCode]CharCode, len(codeToRune))}
	cmap.computeInverseMappings()
	return cmap
}

type cmapArray struct{ Array []cmapObject }

func (cmap *CMap) CIDSystemInfo() CIDSystemInfo { return cmap._dfd }
func (cmap *CMap) WMode() (int, bool)           { return cmap._bfc._fba, cmap._bfc._acf }
func (cmap *CMap) BytesToCharcodes(data []byte) ([]CharCode, bool) {
	var _ac []CharCode
	if cmap._dc == 8 {
		for _, _bbd := range data {
			_ac = append(_ac, CharCode(_bbd))
		}
		return _ac, true
	}
	for _eae := 0; _eae < len(data); {
		_cbf, _cfe, _cdc := cmap.matchCode(data[_eae:])
		if !_cdc {
			_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0074\u0020\u0069\u003d\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", _eae, data, string(data))
			return _ac, false
		}
		_ac = append(_ac, _cbf)
		_eae += _cfe
	}
	return _ac, true
}

func _dfe(_dde string) string {
	_bca := []rune(_dde)
	_ebdf := make([]string, len(_bca))
	for _cfad, _fbe := range _bca {
		_ebdf[_cfad] = _be.Sprintf("\u0025\u0030\u0034\u0078", _fbe)
	}
	return _be.Sprintf("\u003c\u0025\u0073\u003e", _eg.Join(_ebdf, ""))
}

func (cmap *CMap) CharcodeToCID(code CharCode) (CharCode, bool) {
	_gcd, _edg := cmap._ff[code]
	return _gcd, _edg
}

func (_fdeg *cMapParser) parseHexString() (cmapHexString, error) {
	_fdeg._fbeb.ReadByte()
	_abeb := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	_gggd := _bb.Buffer{}
	for {
		_fdeg.skipSpaces()
		_caba, _efdd := _fdeg._fbeb.Peek(1)
		if _efdd != nil {
			return cmapHexString{}, _efdd
		}
		if _caba[0] == '>' {
			_fdeg._fbeb.ReadByte()
			break
		}
		_gacg, _ := _fdeg._fbeb.ReadByte()
		if _bb.IndexByte(_abeb, _gacg) >= 0 {
			_gggd.WriteByte(_gacg)
		}
	}
	if _gggd.Len()%2 == 1 {
		_da.Log.Debug("\u0070\u0061rs\u0065\u0048\u0065x\u0053\u0074\u0072\u0069ng:\u0020ap\u0070\u0065\u006e\u0064\u0069\u006e\u0067 '\u0030\u0027\u0020\u0074\u006f\u0020\u0025#\u0071", _gggd.String())
		_gggd.WriteByte('0')
	}
	_gcca := _gggd.Len() / 2
	_gcag, _ := _d.DecodeString(_gggd.String())
	return cmapHexString{_ccgc: _gcca, _eeea: _gcag}, nil
}

func (cmap *CMap) parseWMode() error {
	var _gga int
	_cggb := false
	for _eba := 0; _eba < 3 && !_cggb; _eba++ {
		_ccg, _baeg := cmap.parseObject()
		if _baeg != nil {
			return _baeg
		}
		switch _fgg := _ccg.(type) {
		case cmapOperand:
			switch _fgg.Operand {
			case "\u0064\u0065\u0066":
				_cggb = true
			default:
				_da.Log.Error("\u0070\u0061\u0072\u0073\u0065\u0057\u004d\u006f\u0064\u0065:\u0020\u0073\u0074\u0061\u0074\u0065\u0020e\u0072\u0072\u006f\u0072\u002e\u0020\u006f\u003d\u0025\u0023\u0076", _ccg)
				return ErrBadCMap
			}
		case cmapInt:
			_gga = int(_fgg._abbg)
		}
	}
	cmap._bfc = integer{_acf: true, _fba: _gga}
	return nil
}

func _acc(_ffd cmapHexString) []rune {
	if len(_ffd._eeea) == 1 {
		return []rune{rune(_ffd._eeea[0])}
	}
	_afa := _ffd._eeea
	if len(_afa)%2 != 0 {
		_afa = append(_afa, 0)
		_da.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u0073\u002e\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0073\u0068\u0065\u0078\u003d\u0025#\u0076\u0020\u0074\u006f\u0020\u0025\u002b\u0076", _ffd, _afa)
	}
	_fcaf := len(_afa) >> 1
	_acgb := make([]uint16, _fcaf)
	for _cfadd := 0; _cfadd < _fcaf; _cfadd++ {
		_acgb[_cfadd] = uint16(_afa[_cfadd<<1])<<8 + uint16(_afa[_cfadd<<1+1])
	}
	_eef := _bg.Decode(_acgb)
	return _eef
}

type integer struct {
	_acf bool
	_fba int
}

func NewCIDSystemInfo(obj _bce.PdfObject) (_df CIDSystemInfo, _ca error) {
	_ea, _ag := _bce.GetDict(obj)
	if !_ag {
		return CIDSystemInfo{}, _bce.ErrTypeError
	}
	_bff, _ag := _bce.GetStringVal(_ea.Get("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"))
	if !_ag {
		return CIDSystemInfo{}, _bce.ErrTypeError
	}
	_ae, _ag := _bce.GetStringVal(_ea.Get("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"))
	if !_ag {
		return CIDSystemInfo{}, _bce.ErrTypeError
	}
	_agc, _ag := _bce.GetIntVal(_ea.Get("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"))
	if !_ag {
		return CIDSystemInfo{}, _bce.ErrTypeError
	}
	return CIDSystemInfo{Registry: _bff, Ordering: _ae, Supplement: _agc}, nil
}

func (_adab *cMapParser) parseObject() (cmapObject, error) {
	_adab.skipSpaces()
	for {
		_dggg, _acd := _adab._fbeb.Peek(2)
		if _acd != nil {
			return nil, _acd
		}
		if _dggg[0] == '%' {
			_adab.parseComment()
			_adab.skipSpaces()
			continue
		} else if _dggg[0] == '/' {
			_ddeb, _eeb := _adab.parseName()
			return _ddeb, _eeb
		} else if _dggg[0] == '(' {
			_cfgd, _aff := _adab.parseString()
			return _cfgd, _aff
		} else if _dggg[0] == '[' {
			_acgd, _agb := _adab.parseArray()
			return _acgd, _agb
		} else if (_dggg[0] == '<') && (_dggg[1] == '<') {
			_bgc, _abf := _adab.parseDict()
			return _bgc, _abf
		} else if _dggg[0] == '<' {
			_bcad, _eaba := _adab.parseHexString()
			return _bcad, _eaba
		} else if _bce.IsDecimalDigit(_dggg[0]) || (_dggg[0] == '-' && _bce.IsDecimalDigit(_dggg[1])) {
			_dcdg, _edgg := _adab.parseNumber()
			if _edgg != nil {
				return nil, _edgg
			}
			return _dcdg, nil
		} else {
			_eacb, _caee := _adab.parseOperand()
			if _caee != nil {
				return nil, _caee
			}
			return _eacb, nil
		}
	}
}

func (cmap *CMap) CIDToCharcode(cid CharCode) (CharCode, bool) {
	_gcg, _eac := cmap._fb[cid]
	return _gcg, _eac
}

func (cmap *CMap) parseSystemInfo() error {
	_cbc := false
	_cffa := false
	_gbdc := ""
	_gdd := false
	_ebfg := CIDSystemInfo{}
	for _ddd := 0; _ddd < 50 && !_gdd; _ddd++ {
		_bef, _dcb := cmap.parseObject()
		if _dcb != nil {
			return _dcb
		}
		switch _dgd := _bef.(type) {
		case cmapDict:
			_fgba := _dgd.Dict
			_fee, _dbg := _fgba["\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"]
			if !_dbg {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_bbb, _dbg := _fee.(cmapString)
			if !_dbg {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ebfg.Registry = _bbb.String
			_fee, _dbg = _fgba["\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"]
			if !_dbg {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_bbb, _dbg = _fee.(cmapString)
			if !_dbg {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ebfg.Ordering = _bbb.String
			_gca, _dbg := _fgba["\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"]
			if !_dbg {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_fea, _dbg := _gca.(cmapInt)
			if !_dbg {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ebfg.Supplement = int(_fea._abbg)
			_gdd = true
		case cmapOperand:
			switch _dgd.Operand {
			case "\u0062\u0065\u0067i\u006e":
				_cbc = true
			case "\u0065\u006e\u0064":
				_gdd = true
			case "\u0064\u0065\u0066":
				_cffa = false
			}
		case cmapName:
			if _cbc {
				_gbdc = _dgd.Name
				_cffa = true
			}
		case cmapString:
			if _cffa {
				switch _gbdc {
				case "\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079":
					_ebfg.Registry = _dgd.String
				case "\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067":
					_ebfg.Ordering = _dgd.String
				}
			}
		case cmapInt:
			if _cffa {
				switch _gbdc {
				case "\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074":
					_ebfg.Supplement = int(_dgd._abbg)
				}
			}
		}
	}
	if !_gdd {
		_da.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u006c\u0079")
		return ErrBadCMap
	}
	cmap._dfd = _ebfg
	return nil
}

func (cmap *CMap) CharcodeToUnicode(code CharCode) (string, bool) {
	if _beb, _ecf := cmap._ce[code]; _ecf {
		return _beb, true
	}
	return MissingCodeString, false
}
