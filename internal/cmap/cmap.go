package cmap

import (
	_c "bufio"
	_g "bytes"
	_e "encoding/hex"
	_ef "errors"
	_f "fmt"
	_ga "io"
	_bg "sort"
	_fc "strconv"
	_fcg "strings"
	_a "unicode/utf16"

	_gc "bitbucket.org/shenghui0779/gopdf/common"
	_gg "bitbucket.org/shenghui0779/gopdf/core"
	_gf "bitbucket.org/shenghui0779/gopdf/internal/cmap/bcmaps"
)

func (cmap *CMap) parseCIDRange() error {
	for {
		_dfa, _cbe := cmap.parseObject()
		if _cbe != nil {
			if _cbe == _ga.EOF {
				break
			}
			return _cbe
		}
		_dad, _afge := _dfa.(cmapHexString)
		if !_afge {
			if _bcadb, _bgdc := _dfa.(cmapOperand); _bgdc {
				if _bcadb.Operand == _bec {
					return nil
				}
				return _ef.New("\u0063\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0072\u0076\u0061\u006c\u0020s\u0074\u0061\u0072\u0074\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
			}
		}
		_dgb := _gfdd(_dad)
		_dfa, _cbe = cmap.parseObject()
		if _cbe != nil {
			if _cbe == _ga.EOF {
				break
			}
			return _cbe
		}
		_bgdf, _afge := _dfa.(cmapHexString)
		if !_afge {
			return _ef.New("\u0063\u0069d\u0020\u0069\u006e\u0074e\u0072\u0076a\u006c\u0020\u0065\u006e\u0064\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_dad._gec) != len(_bgdf._gec) {
			return _ef.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_gaf := _gfdd(_bgdf)
		if _dgb > _gaf {
			_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0049\u0044\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073t\u0061\u0072\u0074\u003d\u0030\u0078\u0025\u0030\u0032\u0078\u0020\u0065\u006e\u0064=\u0030x\u0025\u0030\u0032\u0078", _dgb, _gaf)
			return ErrBadCMap
		}
		_dfa, _cbe = cmap.parseObject()
		if _cbe != nil {
			if _cbe == _ga.EOF {
				break
			}
			return _cbe
		}
		_bbcc, _afge := _dfa.(cmapInt)
		if !_afge {
			return _ef.New("\u0063\u0069\u0064\u0020\u0073t\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		if _bbcc._dgfg < 0 {
			return _ef.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0063\u0069\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065")
		}
		_gbbc := _bbcc._dgfg
		for _bfa := _dgb; _bfa <= _gaf; _bfa++ {
			cmap._aaa[_bfa] = CharCode(_gbbc)
			_gbbc++
		}
		_gc.Log.Trace("C\u0049\u0044\u0020\u0072\u0061\u006eg\u0065\u003a\u0020\u003c\u0030\u0078\u0025\u0058\u003e \u003c\u0030\u0078%\u0058>\u0020\u0025\u0064", _dgb, _gaf, _bbcc._dgfg)
	}
	return nil
}
func (_dgdg *cMapParser) parseArray() (cmapArray, error) {
	_fcd := cmapArray{}
	_fcd.Array = []cmapObject{}
	_dgdg._defg.ReadByte()
	for {
		_dgdg.skipSpaces()
		_fbcb, _fcac := _dgdg._defg.Peek(1)
		if _fcac != nil {
			return _fcd, _fcac
		}
		if _fbcb[0] == ']' {
			_dgdg._defg.ReadByte()
			break
		}
		_aafb, _fcac := _dgdg.parseObject()
		if _fcac != nil {
			return _fcd, _fcac
		}
		_fcd.Array = append(_fcd.Array, _aafb)
	}
	return _fcd, nil
}
func NewToUnicodeCMap(codeToRune map[CharCode]rune) *CMap {
	_cdg := make(map[CharCode]string, len(codeToRune))
	for _ed, _cc := range codeToRune {
		_cdg[_ed] = string(_cc)
	}
	cmap := &CMap{_fd: "\u0041d\u006fb\u0065\u002d\u0049\u0064\u0065n\u0074\u0069t\u0079\u002d\u0055\u0043\u0053", _gad: 2, _be: 16, _gd: CIDSystemInfo{Registry: "\u0041\u0064\u006fb\u0065", Ordering: "\u0055\u0043\u0053", Supplement: 0}, _gga: []Codespace{{Low: 0, High: 0xffff}}, _ab: _cdg, _cg: make(map[string]CharCode, len(codeToRune)), _aaa: make(map[CharCode]CharCode, len(codeToRune)), _efd: make(map[CharCode]CharCode, len(codeToRune))}
	cmap.computeInverseMappings()
	return cmap
}

type fbRange struct {
	_ca  CharCode
	_d   CharCode
	_ffa string
}
type CIDSystemInfo struct {
	Registry   string
	Ordering   string
	Supplement int
}

func (_bbff *cMapParser) parseString() (cmapString, error) {
	_bbff._defg.ReadByte()
	_ggf := _g.Buffer{}
	_bed := 1
	for {
		_cca, _fedgg := _bbff._defg.Peek(1)
		if _fedgg != nil {
			return cmapString{_ggf.String()}, _fedgg
		}
		if _cca[0] == '\\' {
			_bbff._defg.ReadByte()
			_acda, _fegc := _bbff._defg.ReadByte()
			if _fegc != nil {
				return cmapString{_ggf.String()}, _fegc
			}
			if _gg.IsOctalDigit(_acda) {
				_afag, _bdce := _bbff._defg.Peek(2)
				if _bdce != nil {
					return cmapString{_ggf.String()}, _bdce
				}
				var _gbdd []byte
				_gbdd = append(_gbdd, _acda)
				for _, _acdd := range _afag {
					if _gg.IsOctalDigit(_acdd) {
						_gbdd = append(_gbdd, _acdd)
					} else {
						break
					}
				}
				_bbff._defg.Discard(len(_gbdd) - 1)
				_gc.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _gbdd)
				_eea, _bdce := _fc.ParseUint(string(_gbdd), 8, 32)
				if _bdce != nil {
					return cmapString{_ggf.String()}, _bdce
				}
				_ggf.WriteByte(byte(_eea))
				continue
			}
			switch _acda {
			case 'n':
				_ggf.WriteByte('\n')
			case 'r':
				_ggf.WriteByte('\r')
			case 't':
				_ggf.WriteByte('\t')
			case 'b':
				_ggf.WriteByte('\b')
			case 'f':
				_ggf.WriteByte('\f')
			case '(':
				_ggf.WriteByte('(')
			case ')':
				_ggf.WriteByte(')')
			case '\\':
				_ggf.WriteByte('\\')
			}
			continue
		} else if _cca[0] == '(' {
			_bed++
		} else if _cca[0] == ')' {
			_bed--
			if _bed == 0 {
				_bbff._defg.ReadByte()
				break
			}
		}
		_fdac, _ := _bbff._defg.ReadByte()
		_ggf.WriteByte(_fdac)
	}
	return cmapString{_ggf.String()}, nil
}
func (cmap *CMap) CharcodeToCID(code CharCode) (CharCode, bool) {
	_eac, _fdae := cmap._aaa[code]
	return _eac, _fdae
}
func _dc(_ade bool) *CMap {
	_gdb := 16
	if _ade {
		_gdb = 8
	}
	return &CMap{_be: _gdb, _aaa: make(map[CharCode]CharCode), _efd: make(map[CharCode]CharCode), _ab: make(map[CharCode]string), _cg: make(map[string]CharCode)}
}

type cmapHexString struct {
	_caed int
	_gec  []byte
}

func (cmap *CMap) matchCode(_acd []byte) (_fba CharCode, _fbd int, _gbff bool) {
	for _cbg := 0; _cbg < _ff; _cbg++ {
		if _cbg < len(_acd) {
			_fba = _fba<<8 | CharCode(_acd[_cbg])
			_fbd++
		}
		_gbff = cmap.inCodespace(_fba, _cbg+1)
		if _gbff {
			return _fba, _fbd, true
		}
	}
	_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063o\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0020m\u0061t\u0063\u0068\u0065\u0073\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d=\u0025\u0023\u0071\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", _acd, string(_acd), cmap)
	return 0, 0, false
}
func (_gbd *cMapParser) parseComment() (string, error) {
	var _fdb _g.Buffer
	_, _ffg := _gbd.skipSpaces()
	if _ffg != nil {
		return _fdb.String(), _ffg
	}
	_ggd := true
	for {
		_geab, _afaa := _gbd._defg.Peek(1)
		if _afaa != nil {
			_gc.Log.Debug("p\u0061r\u0073\u0065\u0043\u006f\u006d\u006d\u0065\u006et\u003a\u0020\u0065\u0072r=\u0025\u0076", _afaa)
			return _fdb.String(), _afaa
		}
		if _ggd && _geab[0] != '%' {
			return _fdb.String(), ErrBadCMapComment
		}
		_ggd = false
		if (_geab[0] != '\r') && (_geab[0] != '\n') {
			_bcd, _ := _gbd._defg.ReadByte()
			_fdb.WriteByte(_bcd)
		} else {
			break
		}
	}
	return _fdb.String(), nil
}
func (cmap *CMap) CharcodeToUnicode(code CharCode) (string, bool) {
	if _gee, _gbfe := cmap._ab[code]; _gbfe {
		return _gee, true
	}
	return MissingCodeString, false
}
func (cmap *CMap) parseBfrange() error {
	for {
		var _bcada CharCode
		_afc, _gcfg := cmap.parseObject()
		if _gcfg != nil {
			if _gcfg == _ga.EOF {
				break
			}
			return _gcfg
		}
		switch _bebg := _afc.(type) {
		case cmapOperand:
			if _bebg.Operand == _cdga {
				return nil
			}
			return _ef.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_bcada = _gfdd(_bebg)
		default:
			return _ef.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		var _gde CharCode
		_afc, _gcfg = cmap.parseObject()
		if _gcfg != nil {
			if _gcfg == _ga.EOF {
				break
			}
			return _gcfg
		}
		switch _dff := _afc.(type) {
		case cmapOperand:
			_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065\u0020\u0062\u0066r\u0061\u006e\u0067\u0065\u0020\u0074\u0072i\u0070\u006c\u0065\u0074")
			return ErrBadCMap
		case cmapHexString:
			_gde = _gfdd(_dff)
			if _gde > 0xffff {
				_gde = 0xffff
			}
		default:
			_gc.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _afc)
			return ErrBadCMap
		}
		_afc, _gcfg = cmap.parseObject()
		if _gcfg != nil {
			if _gcfg == _ga.EOF {
				break
			}
			return _gcfg
		}
		switch _gfecg := _afc.(type) {
		case cmapArray:
			if len(_gfecg.Array) != int(_gde-_bcada)+1 {
				_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0069\u0074\u0065\u006d\u0073\u0020\u0069\u006e\u0020a\u0072\u0072\u0061\u0079")
				return ErrBadCMap
			}
			for _bfdc := _bcada; _bfdc <= _gde; _bfdc++ {
				_gcd := _gfecg.Array[_bfdc-_bcada]
				_fcc, _bac := _gcd.(cmapHexString)
				if !_bac {
					return _ef.New("\u006e\u006f\u006e-h\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				}
				_acde := _ebd(_fcc)
				cmap._ab[_bfdc] = string(_acde)
			}
		case cmapHexString:
			_ede := _ebd(_gfecg)
			_fbdd := len(_ede)
			for _dfgg := _bcada; _dfgg <= _gde; _dfgg++ {
				cmap._ab[_dfgg] = string(_ede)
				if _fbdd > 0 {
					_ede[_fbdd-1]++
				} else {
					_gc.Log.Debug("\u004e\u006f\u0020c\u006d\u0061\u0070\u0020\u0074\u0061\u0072\u0067\u0065\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0066\u006f\u0072\u0020\u0025\u0023\u0076", _dfgg)
				}
				if _dfgg == 1<<32-1 {
					break
				}
			}
		default:
			_gc.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _afc)
			return ErrBadCMap
		}
	}
	return nil
}
func (cmap *CMap) CIDSystemInfo() CIDSystemInfo { return cmap._gd }

type cmapArray struct{ Array []cmapObject }

func IsPredefinedCMap(name string) bool { return _gf.AssetExists(name) }
func (cmap *CMap) parseBfchar() error {
	for {
		_bgc, _faa := cmap.parseObject()
		if _faa != nil {
			if _faa == _ga.EOF {
				break
			}
			return _faa
		}
		var _gcce CharCode
		switch _gbc := _bgc.(type) {
		case cmapOperand:
			if _gbc.Operand == _eca {
				return nil
			}
			return _ef.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_gcce = _gfdd(_gbc)
		default:
			return _ef.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_bgc, _faa = cmap.parseObject()
		if _faa != nil {
			if _faa == _ga.EOF {
				break
			}
			return _faa
		}
		var _egec []rune
		switch _bcge := _bgc.(type) {
		case cmapOperand:
			if _bcge.Operand == _eca {
				return nil
			}
			_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020o\u0070\u0065\u0072\u0061\u006e\u0064\u002e\u0020\u0025\u0023\u0076", _bcge)
			return ErrBadCMap
		case cmapHexString:
			_egec = _ebd(_bcge)
		case cmapName:
			_gc.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0061\u006de\u002e \u0025\u0023\u0076", _bcge)
			_egec = []rune{MissingCodeRune}
		default:
			_gc.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0074\u0079\u0070e\u002e \u0025\u0023\u0076", _bgc)
			return ErrBadCMap
		}
		cmap._ab[_gcce] = string(_egec)
	}
	return nil
}
func (cmap *CMap) WMode() (int, bool) { return cmap._de._gbg, cmap._de._eae }
func (cmap *CMap) Type() int          { return cmap._gad }

const (
	_bba = 100
	_gce = "\u000a\u002f\u0043\u0049\u0044\u0049\u006e\u0069\u0074\u0020\u002f\u0050\u0072\u006fc\u0053\u0065\u0074\u0020\u0066\u0069\u006e\u0064\u0072es\u006fu\u0072c\u0065 \u0062\u0065\u0067\u0069\u006e\u000a\u0031\u0032\u0020\u0064\u0069\u0063\u0074\u0020\u0062\u0065\u0067\u0069n\u000a\u0062\u0065\u0067\u0069\u006e\u0063\u006d\u0061\u0070\n\u002f\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u003c\u003c\u0020\u002f\u0052\u0065\u0067\u0069\u0073t\u0072\u0079\u0020\u0028\u0041\u0064\u006f\u0062\u0065\u0029\u0020\u002f\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0028\u0055\u0043\u0053)\u0020\u002f\u0053\u0075\u0070p\u006c\u0065\u006d\u0065\u006et\u0020\u0030\u0020\u003e\u003e\u0020\u0064\u0065\u0066\u000a\u002f\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u002f\u0041\u0064\u006f\u0062\u0065-\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0055\u0043\u0053\u0020\u0064\u0065\u0066\u000a\u002fC\u004d\u0061\u0070\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0064\u0065\u0066\u000a\u0031\u0020\u0062\u0065\u0067\u0069\u006e\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063e\u0072\u0061n\u0067\u0065\n\u003c\u0030\u0030\u0030\u0030\u003e\u0020<\u0046\u0046\u0046\u0046\u003e\u000a\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065\u000a"
	_gcc = "\u0065\u006e\u0064\u0063\u006d\u0061\u0070\u000a\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u0063ur\u0072e\u006e\u0074\u0064\u0069\u0063\u0074\u0020\u002f\u0043\u004d\u0061\u0070 \u0064\u0065\u0066\u0069\u006e\u0065\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0070\u006fp\u000a\u0065\u006e\u0064\u000a\u0065\u006e\u0064\u000a"
)

type cmapFloat struct{ _bfb float64 }

func (cmap *CMap) inCodespace(_cfe CharCode, _fdab int) bool {
	for _, _baf := range cmap._gga {
		if _baf.Low <= _cfe && _cfe <= _baf.High && _fdab == _baf.NumBytes {
			return true
		}
	}
	return false
}
func (cmap *CMap) toBfData() string {
	if len(cmap._ab) == 0 {
		return ""
	}
	_def := make([]CharCode, 0, len(cmap._ab))
	for _ddg := range cmap._ab {
		_def = append(_def, _ddg)
	}
	_bg.Slice(_def, func(_dgd, _ffd int) bool { return _def[_dgd] < _def[_ffd] })
	var _feg []charRange
	_bbge := charRange{_def[0], _def[0]}
	_aafa := cmap._ab[_def[0]]
	for _, _bgb := range _def[1:] {
		_fgc := cmap._ab[_bgb]
		if _bgb == _bbge._bb+1 && _bcf(_fgc) == _bcf(_aafa)+1 {
			_bbge._bb = _bgb
		} else {
			_feg = append(_feg, _bbge)
			_bbge._bd, _bbge._bb = _bgb, _bgb
		}
		_aafa = _fgc
	}
	_feg = append(_feg, _bbge)
	var _bag []CharCode
	var _gfgf []fbRange
	for _, _ce := range _feg {
		if _ce._bd == _ce._bb {
			_bag = append(_bag, _ce._bd)
		} else {
			_gfgf = append(_gfgf, fbRange{_ca: _ce._bd, _d: _ce._bb, _ffa: cmap._ab[_ce._bd]})
		}
	}
	_gc.Log.Trace("\u0063\u0068ar\u0052\u0061\u006eg\u0065\u0073\u003d\u0025d f\u0062Ch\u0061\u0072\u0073\u003d\u0025\u0064\u0020fb\u0052\u0061\u006e\u0067\u0065\u0073\u003d%\u0064", len(_feg), len(_bag), len(_gfgf))
	var _age []string
	if len(_bag) > 0 {
		_cde := (len(_bag) + _bba - 1) / _bba
		for _gbb := 0; _gbb < _cde; _gbb++ {
			_gff := _dcd(len(_bag)-_gbb*_bba, _bba)
			_age = append(_age, _f.Sprintf("\u0025\u0064\u0020\u0062\u0065\u0067\u0069\u006e\u0062f\u0063\u0068\u0061\u0072", _gff))
			for _afe := 0; _afe < _gff; _afe++ {
				_aee := _bag[_gbb*_bba+_afe]
				_dcf := cmap._ab[_aee]
				_age = append(_age, _f.Sprintf("\u003c%\u0030\u0034\u0078\u003e\u0020\u0025s", _aee, _gcf(_dcf)))
			}
			_age = append(_age, "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r")
		}
	}
	if len(_gfgf) > 0 {
		_dfg := (len(_gfgf) + _bba - 1) / _bba
		for _gcg := 0; _gcg < _dfg; _gcg++ {
			_fcf := _dcd(len(_gfgf)-_gcg*_bba, _bba)
			_age = append(_age, _f.Sprintf("\u0025d\u0020b\u0065\u0067\u0069\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065", _fcf))
			for _dee := 0; _dee < _fcf; _dee++ {
				_fgg := _gfgf[_gcg*_bba+_dee]
				_age = append(_age, _f.Sprintf("\u003c%\u00304\u0078\u003e\u003c\u0025\u0030\u0034\u0078\u003e\u0020\u0025\u0073", _fgg._ca, _fgg._d, _gcf(_fgg._ffa)))
			}
			_age = append(_age, "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065")
		}
	}
	return _fcg.Join(_age, "\u000a")
}

type cmapOperand struct{ Operand string }

func (_cea *cMapParser) skipSpaces() (int, error) {
	_bebe := 0
	for {
		_dac, _bee := _cea._defg.Peek(1)
		if _bee != nil {
			return 0, _bee
		}
		if _gg.IsWhiteSpace(_dac[0]) {
			_cea._defg.ReadByte()
			_bebe++
		} else {
			break
		}
	}
	return _bebe, nil
}
func (cmap *CMap) Name() string                          { return cmap._fd }
func (cmap *CMap) StringToCID(s string) (CharCode, bool) { _abf, _ea := cmap._cg[s]; return _abf, _ea }

type charRange struct {
	_bd CharCode
	_bb CharCode
}

func (_ag *CIDSystemInfo) String() string {
	return _f.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0025\u0030\u0033\u0064", _ag.Registry, _ag.Ordering, _ag.Supplement)
}
func (cmap *CMap) String() string {
	_dfb := cmap._gd
	_add := []string{_f.Sprintf("\u006e\u0062\u0069\u0074\u0073\u003a\u0025\u0064", cmap._be), _f.Sprintf("\u0074y\u0070\u0065\u003a\u0025\u0064", cmap._gad)}
	if cmap._ffc != "" {
		_add = append(_add, _f.Sprintf("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0025\u0073", cmap._ffc))
	}
	if cmap._fda != "" {
		_add = append(_add, _f.Sprintf("u\u0073\u0065\u0063\u006d\u0061\u0070\u003a\u0025\u0023\u0071", cmap._fda))
	}
	_add = append(_add, _f.Sprintf("\u0073\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u003a\u0025\u0073", _dfb.String()))
	if len(cmap._gga) > 0 {
		_add = append(_add, _f.Sprintf("\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u003a\u0025\u0064", len(cmap._gga)))
	}
	if len(cmap._ab) > 0 {
		_add = append(_add, _f.Sprintf("\u0063\u006fd\u0065\u0054\u006fU\u006e\u0069\u0063\u006f\u0064\u0065\u003a\u0025\u0064", len(cmap._ab)))
	}
	return _f.Sprintf("\u0043\u004d\u0041P\u007b\u0025\u0023\u0071\u0020\u0025\u0073\u007d", cmap._fd, _fcg.Join(_add, "\u0020"))
}

type integer struct {
	_eae bool
	_gbg int
}

func (cmap *CMap) computeInverseMappings() {
	for _cad, _aaf := range cmap._aaa {
		if _aaag, _cf := cmap._efd[_aaf]; !_cf || (_cf && _aaag > _cad) {
			cmap._efd[_aaf] = _cad
		}
	}
	for _gb, _bbg := range cmap._ab {
		if _cae, _abg := cmap._cg[_bbg]; !_abg || (_abg && _cae > _gb) {
			cmap._cg[_bbg] = _gb
		}
	}
	_bg.Slice(cmap._gga, func(_cfa, _gadf int) bool { return cmap._gga[_cfa].Low < cmap._gga[_gadf].Low })
}
func _bcf(_beb string) rune { _aeb := []rune(_beb); return _aeb[len(_aeb)-1] }
func (cmap *CMap) CIDToCharcode(cid CharCode) (CharCode, bool) {
	_ee, _bbf := cmap._efd[cid]
	return _ee, _bbf
}

const (
	_dfe  = "\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"
	_edf  = "\u0062e\u0067\u0069\u006e\u0063\u006d\u0061p"
	_fbcc = "\u0065n\u0064\u0063\u006d\u0061\u0070"
	_fee  = "\u0062\u0065\u0067\u0069nc\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0072\u0061\u006e\u0067\u0065"
	_bfdb = "\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065"
	_afec = "b\u0065\u0067\u0069\u006e\u0062\u0066\u0063\u0068\u0061\u0072"
	_eca  = "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r"
	_fgdg = "\u0062\u0065\u0067i\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_cdga = "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_gae  = "\u0062\u0065\u0067\u0069\u006e\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_bec  = "e\u006e\u0064\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_adce = "\u0075s\u0065\u0063\u006d\u0061\u0070"
	_ddd  = "\u0057\u004d\u006fd\u0065"
	_fbac = "\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065"
	_bddb = "\u0043\u004d\u0061\u0070\u0054\u0079\u0070\u0065"
	_eef  = "C\u004d\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e"
)

func NewCIDSystemInfo(obj _gg.PdfObject) (_cd CIDSystemInfo, _ba error) {
	_aa, _cb := _gg.GetDict(obj)
	if !_cb {
		return CIDSystemInfo{}, _gg.ErrTypeError
	}
	_ad, _cb := _gg.GetStringVal(_aa.Get("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"))
	if !_cb {
		return CIDSystemInfo{}, _gg.ErrTypeError
	}
	_fg, _cb := _gg.GetStringVal(_aa.Get("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"))
	if !_cb {
		return CIDSystemInfo{}, _gg.ErrTypeError
	}
	_dg, _cb := _gg.GetIntVal(_aa.Get("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"))
	if !_cb {
		return CIDSystemInfo{}, _gg.ErrTypeError
	}
	return CIDSystemInfo{Registry: _ad, Ordering: _fg, Supplement: _dg}, nil
}

type cmapInt struct{ _dgfg int64 }

func (cmap *CMap) Bytes() []byte {
	_gc.Log.Trace("\u0063\u006d\u0061\u0070.B\u0079\u0074\u0065\u0073\u003a\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", cmap.String())
	if len(cmap._ac) > 0 {
		return cmap._ac
	}
	cmap._ac = []byte(_fcg.Join([]string{_gce, cmap.toBfData(), _gcc}, "\u000a"))
	return cmap._ac
}
func (cmap *CMap) parseType() error {
	_bce := 0
	_ggac := false
	for _fedg := 0; _fedg < 3 && !_ggac; _fedg++ {
		_aeea, _afg := cmap.parseObject()
		if _afg != nil {
			return _afg
		}
		switch _aga := _aeea.(type) {
		case cmapOperand:
			switch _aga.Operand {
			case "\u0064\u0065\u0066":
				_ggac = true
			default:
				_gc.Log.Error("\u0070\u0061r\u0073\u0065\u0054\u0079\u0070\u0065\u003a\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u006f=%\u0023\u0076", _aeea)
				return ErrBadCMap
			}
		case cmapInt:
			_bce = int(_aga._dgfg)
		}
	}
	cmap._gad = _bce
	return nil
}
func _gfdd(_afed cmapHexString) CharCode {
	_fdbg := CharCode(0)
	for _, _dfff := range _afed._gec {
		_fdbg <<= 8
		_fdbg |= CharCode(_dfff)
	}
	return _fdbg
}

type cmapObject interface{}

func _cdd(_deg []byte) *cMapParser {
	_deea := cMapParser{}
	_acg := _g.NewBuffer(_deg)
	_deea._defg = _c.NewReader(_acg)
	return &_deea
}
func (_bace *cMapParser) parseDict() (cmapDict, error) {
	_gc.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_gac := _dgfd()
	_ddf, _ := _bace._defg.ReadByte()
	if _ddf != '<' {
		return _gac, ErrBadCMapDict
	}
	_ddf, _ = _bace._defg.ReadByte()
	if _ddf != '<' {
		return _gac, ErrBadCMapDict
	}
	for {
		_bace.skipSpaces()
		_baff, _edgcd := _bace._defg.Peek(2)
		if _edgcd != nil {
			return _gac, _edgcd
		}
		if (_baff[0] == '>') && (_baff[1] == '>') {
			_bace._defg.ReadByte()
			_bace._defg.ReadByte()
			break
		}
		_gdd, _edgcd := _bace.parseName()
		_gc.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _gdd.Name)
		if _edgcd != nil {
			_gc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u002e\u0020\u0065\u0072r=\u0025\u0076", _edgcd)
			return _gac, _edgcd
		}
		_bace.skipSpaces()
		_edab, _edgcd := _bace.parseObject()
		if _edgcd != nil {
			return _gac, _edgcd
		}
		_gac.Dict[_gdd.Name] = _edab
		_bace.skipSpaces()
		_baff, _edgcd = _bace._defg.Peek(3)
		if _edgcd != nil {
			return _gac, _edgcd
		}
		if string(_baff) == "\u0064\u0065\u0066" {
			_bace._defg.Discard(3)
		}
	}
	return _gac, nil
}

type CMap struct {
	*cMapParser
	_fd  string
	_be  int
	_gad int
	_ffc string
	_fda string
	_gd  CIDSystemInfo
	_gga []Codespace
	_aaa map[CharCode]CharCode
	_efd map[CharCode]CharCode
	_ab  map[CharCode]string
	_cg  map[string]CharCode
	_ac  []byte
	_cgg *_gg.PdfObjectStream
	_de  integer
}
type cmapName struct{ Name string }

func _dcd(_cec, _ffe int) int {
	if _cec < _ffe {
		return _cec
	}
	return _ffe
}
func LoadPredefinedCMap(name string) (*CMap, error) {
	cmap, _adf := _bdc(name)
	if _adf != nil {
		return nil, _adf
	}
	if cmap._fda == "" {
		cmap.computeInverseMappings()
		return cmap, nil
	}
	_dcb, _adf := _bdc(cmap._fda)
	if _adf != nil {
		return nil, _adf
	}
	for _bdg, _af := range _dcb._aaa {
		if _, _ae := cmap._aaa[_bdg]; !_ae {
			cmap._aaa[_bdg] = _af
		}
	}
	cmap._gga = append(cmap._gga, _dcb._gga...)
	cmap.computeInverseMappings()
	return cmap, nil
}
func _ccac(_fbf cmapHexString) rune {
	_gdg := _ebd(_fbf)
	if _dbgc := len(_gdg); _dbgc == 0 {
		_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054o\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065c\u0074\u0065\u0064\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0072u\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023\u0076", _fbf)
		return MissingCodeRune
	}
	if len(_gdg) > 1 {
		_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0078\u0061\u0063\u0074\u006c\u0079\u0020\u006f\u006e\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023v\u0020\u002d\u003e\u0020\u0025#\u0076", _fbf, _gdg)
	}
	return _gdg[0]
}
func (cmap *CMap) parseWMode() error {
	var _cfg int
	_gcge := false
	for _eeb := 0; _eeb < 3 && !_gcge; _eeb++ {
		_gfef, _gag := cmap.parseObject()
		if _gag != nil {
			return _gag
		}
		switch _cafc := _gfef.(type) {
		case cmapOperand:
			switch _cafc.Operand {
			case "\u0064\u0065\u0066":
				_gcge = true
			default:
				_gc.Log.Error("\u0070\u0061\u0072\u0073\u0065\u0057\u004d\u006f\u0064\u0065:\u0020\u0073\u0074\u0061\u0074\u0065\u0020e\u0072\u0072\u006f\u0072\u002e\u0020\u006f\u003d\u0025\u0023\u0076", _gfef)
				return ErrBadCMap
			}
		case cmapInt:
			_cfg = int(_cafc._dgfg)
		}
	}
	cmap._de = integer{_eae: true, _gbg: _cfg}
	return nil
}
func (cmap *CMap) NBits() int { return cmap._be }

const (
	_ff               = 4
	MissingCodeRune   = '\ufffd'
	MissingCodeString = string(MissingCodeRune)
)

func (cmap *CMap) parse() error {
	var _fgca cmapObject
	for {
		_ead, _bca := cmap.parseObject()
		if _bca != nil {
			if _bca == _ga.EOF {
				break
			}
			_gc.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u004d\u0061\u0070\u003a\u0020\u0025\u0076", _bca)
			return _bca
		}
		switch _cgf := _ead.(type) {
		case cmapOperand:
			_dgdf := _cgf
			switch _dgdf.Operand {
			case _fee:
				_abd := cmap.parseCodespaceRange()
				if _abd != nil {
					return _abd
				}
			case _gae:
				_aad := cmap.parseCIDRange()
				if _aad != nil {
					return _aad
				}
			case _afec:
				_eec := cmap.parseBfchar()
				if _eec != nil {
					return _eec
				}
			case _fgdg:
				_cgc := cmap.parseBfrange()
				if _cgc != nil {
					return _cgc
				}
			case _adce:
				if _fgca == nil {
					_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u0073\u0065\u0063m\u0061\u0070\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f \u0061\u0072\u0067")
					return ErrBadCMap
				}
				_fcae, _bde := _fgca.(cmapName)
				if !_bde {
					_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0075\u0073\u0065\u0063\u006d\u0061\u0070\u0020\u0061\u0072\u0067\u0020\u006eo\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u0020\u0025\u0023\u0076", _fgca)
					return ErrBadCMap
				}
				cmap._fda = _fcae.Name
			case _dfe:
				_fbc := cmap.parseSystemInfo()
				if _fbc != nil {
					return _fbc
				}
			}
		case cmapName:
			_agf := _cgf
			switch _agf.Name {
			case _dfe:
				_dcg := cmap.parseSystemInfo()
				if _dcg != nil {
					return _dcg
				}
			case _fbac:
				_db := cmap.parseName()
				if _db != nil {
					return _db
				}
			case _bddb:
				_cee := cmap.parseType()
				if _cee != nil {
					return _cee
				}
			case _eef:
				_gfd := cmap.parseVersion()
				if _gfd != nil {
					return _gfd
				}
			case _ddd:
				if _bca = cmap.parseWMode(); _bca != nil {
					return _bca
				}
			}
		}
		_fgca = _ead
	}
	return nil
}
func (cmap *CMap) parseCodespaceRange() error {
	for {
		_adc, _gdc := cmap.parseObject()
		if _gdc != nil {
			if _gdc == _ga.EOF {
				break
			}
			return _gdc
		}
		_cab, _fdabd := _adc.(cmapHexString)
		if !_fdabd {
			if _caeb, _ada := _adc.(cmapOperand); _ada {
				if _caeb.Operand == _bfdb {
					return nil
				}
				return _ef.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
			}
		}
		_adc, _gdc = cmap.parseObject()
		if _gdc != nil {
			if _gdc == _ga.EOF {
				break
			}
			return _gdc
		}
		_bbea, _fdabd := _adc.(cmapHexString)
		if !_fdabd {
			return _ef.New("\u006e\u006f\u006e-\u0068\u0065\u0078\u0020\u0068\u0069\u0067\u0068")
		}
		if len(_cab._gec) != len(_bbea._gec) {
			return _ef.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_debb := _gfdd(_cab)
		_ged := _gfdd(_bbea)
		if _ged < _debb {
			_gc.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0061d\u0020\u0063\u006fd\u0065\u0073\u0070\u0061\u0063\u0065\u002e\u0020\u006cow\u003d\u0030\u0078%\u0030\u0032x\u0020\u0068\u0069\u0067\u0068\u003d0\u0078\u00250\u0032\u0078", _debb, _ged)
			return ErrBadCMap
		}
		_bbee := _bbea._caed
		_dbf := Codespace{NumBytes: _bbee, Low: _debb, High: _ged}
		cmap._gga = append(cmap._gga, _dbf)
		_gc.Log.Trace("\u0043\u006f\u0064e\u0073\u0070\u0061\u0063e\u0020\u006c\u006f\u0077\u003a\u0020\u0030x\u0025\u0058\u002c\u0020\u0068\u0069\u0067\u0068\u003a\u0020\u0030\u0078\u0025\u0058", _debb, _ged)
	}
	if len(cmap._gga) == 0 {
		_gc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u0020\u0069\u006e\u0020\u0063ma\u0070\u002e")
		return ErrBadCMap
	}
	return nil
}
func (cmap *CMap) Stream() (*_gg.PdfObjectStream, error) {
	if cmap._cgg != nil {
		return cmap._cgg, nil
	}
	_fec, _fgda := _gg.MakeStream(cmap.Bytes(), _gg.NewFlateEncoder())
	if _fgda != nil {
		return nil, _fgda
	}
	cmap._cgg = _fec
	return cmap._cgg, nil
}
func (cmap *CMap) parseSystemInfo() error {
	_gfec := false
	_dcgb := false
	_fdfb := ""
	_cbag := false
	_ega := CIDSystemInfo{}
	for _gab := 0; _gab < 50 && !_cbag; _gab++ {
		_dec, _gdbd := cmap.parseObject()
		if _gdbd != nil {
			return _gdbd
		}
		switch _bfd := _dec.(type) {
		case cmapDict:
			_eaga := _bfd.Dict
			_fbb, _ege := _eaga["\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"]
			if !_ege {
				_gc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_bcgd, _ege := _fbb.(cmapString)
			if !_ege {
				_gc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ega.Registry = _bcgd.String
			_fbb, _ege = _eaga["\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"]
			if !_ege {
				_gc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_bcgd, _ege = _fbb.(cmapString)
			if !_ege {
				_gc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ega.Ordering = _bcgd.String
			_aedf, _ege := _eaga["\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"]
			if !_ege {
				_gc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_bdd, _ege := _aedf.(cmapInt)
			if !_ege {
				_gc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ega.Supplement = int(_bdd._dgfg)
			_cbag = true
		case cmapOperand:
			switch _bfd.Operand {
			case "\u0062\u0065\u0067i\u006e":
				_gfec = true
			case "\u0065\u006e\u0064":
				_cbag = true
			case "\u0064\u0065\u0066":
				_dcgb = false
			}
		case cmapName:
			if _gfec {
				_fdfb = _bfd.Name
				_dcgb = true
			}
		case cmapString:
			if _dcgb {
				switch _fdfb {
				case "\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079":
					_ega.Registry = _bfd.String
				case "\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067":
					_ega.Ordering = _bfd.String
				}
			}
		case cmapInt:
			if _dcgb {
				switch _fdfb {
				case "\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074":
					_ega.Supplement = int(_bfd._dgfg)
				}
			}
		}
	}
	if !_cbag {
		_gc.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u006c\u0079")
		return ErrBadCMap
	}
	cmap._gd = _ega
	return nil
}
func (_egaa *cMapParser) parseNumber() (cmapObject, error) {
	_bdec, _egb := _gg.ParseNumber(_egaa._defg)
	if _egb != nil {
		return nil, _egb
	}
	switch _edeb := _bdec.(type) {
	case *_gg.PdfObjectFloat:
		return cmapFloat{float64(*_edeb)}, nil
	case *_gg.PdfObjectInteger:
		return cmapInt{int64(*_edeb)}, nil
	}
	return nil, _f.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _bdec)
}
func _ebd(_ddc cmapHexString) []rune {
	if len(_ddc._gec) == 1 {
		return []rune{rune(_ddc._gec[0])}
	}
	_fdad := _ddc._gec
	if len(_fdad)%2 != 0 {
		_fdad = append(_fdad, 0)
		_gc.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u0073\u002e\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0073\u0068\u0065\u0078\u003d\u0025#\u0076\u0020\u0074\u006f\u0020\u0025\u002b\u0076", _ddc, _fdad)
	}
	_dedc := len(_fdad) >> 1
	_eaed := make([]uint16, _dedc)
	for _cbfb := 0; _cbfb < _dedc; _cbfb++ {
		_eaed[_cbfb] = uint16(_fdad[_cbfb<<1])<<8 + uint16(_fdad[_cbfb<<1+1])
	}
	_dddf := _a.Decode(_eaed)
	return _dddf
}
func (cmap *CMap) parseVersion() error {
	_dbb := ""
	_fcgb := false
	for _bcad := 0; _bcad < 3 && !_fcgb; _bcad++ {
		_fdaa, _aag := cmap.parseObject()
		if _aag != nil {
			return _aag
		}
		switch _ceeb := _fdaa.(type) {
		case cmapOperand:
			switch _ceeb.Operand {
			case "\u0064\u0065\u0066":
				_fcgb = true
			default:
				_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0070\u0061\u0072\u0073\u0065\u0056e\u0072\u0073\u0069\u006f\u006e\u003a \u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076", _fdaa)
				return ErrBadCMap
			}
		case cmapInt:
			_dbb = _f.Sprintf("\u0025\u0064", _ceeb._dgfg)
		case cmapFloat:
			_dbb = _f.Sprintf("\u0025\u0066", _ceeb._bfb)
		case cmapString:
			_dbb = _ceeb.String
		default:
			_gc.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020p\u0061\u0072\u0073\u0065Ver\u0073io\u006e\u003a\u0020\u0042\u0061\u0064\u0020ty\u0070\u0065\u002e\u0020\u006f\u003d\u0025#\u0076", _fdaa)
		}
	}
	cmap._ffc = _dbb
	return nil
}

type cmapDict struct{ Dict map[string]cmapObject }
type Codespace struct {
	NumBytes int
	Low      CharCode
	High     CharCode
}
type cmapString struct{ String string }

func (_abb *cMapParser) parseHexString() (cmapHexString, error) {
	_abb._defg.ReadByte()
	_fbg := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	_cefg := _g.Buffer{}
	for {
		_abb.skipSpaces()
		_bbeb, _gfda := _abb._defg.Peek(1)
		if _gfda != nil {
			return cmapHexString{}, _gfda
		}
		if _bbeb[0] == '>' {
			_abb._defg.ReadByte()
			break
		}
		_bafc, _ := _abb._defg.ReadByte()
		if _g.IndexByte(_fbg, _bafc) >= 0 {
			_cefg.WriteByte(_bafc)
		}
	}
	if _cefg.Len()%2 == 1 {
		_gc.Log.Debug("\u0070\u0061rs\u0065\u0048\u0065x\u0053\u0074\u0072\u0069ng:\u0020ap\u0070\u0065\u006e\u0064\u0069\u006e\u0067 '\u0030\u0027\u0020\u0074\u006f\u0020\u0025#\u0071", _cefg.String())
		_cefg.WriteByte('0')
	}
	_eecg := _cefg.Len() / 2
	_fbbc, _ := _e.DecodeString(_cefg.String())
	return cmapHexString{_caed: _eecg, _gec: _fbbc}, nil
}
func (_efg *cMapParser) parseOperand() (cmapOperand, error) {
	_eff := cmapOperand{}
	_aaca := _g.Buffer{}
	for {
		_faba, _bcaf := _efg._defg.Peek(1)
		if _bcaf != nil {
			if _bcaf == _ga.EOF {
				break
			}
			return _eff, _bcaf
		}
		if _gg.IsDelimiter(_faba[0]) {
			break
		}
		if _gg.IsWhiteSpace(_faba[0]) {
			break
		}
		_ccde, _ := _efg._defg.ReadByte()
		_aaca.WriteByte(_ccde)
	}
	if _aaca.Len() == 0 {
		return _eff, _f.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	_eff.Operand = _aaca.String()
	return _eff, nil
}

type cMapParser struct{ _defg *_c.Reader }

func _bdc(_fb string) (*CMap, error) {
	_cbf, _adea := _gf.Asset(_fb)
	if _adea != nil {
		return nil, _adea
	}
	return LoadCmapFromDataCID(_cbf)
}
func LoadCmapFromDataCID(data []byte) (*CMap, error) { return LoadCmapFromData(data, false) }

var (
	ErrBadCMap        = _ef.New("\u0062\u0061\u0064\u0020\u0063\u006d\u0061\u0070")
	ErrBadCMapComment = _ef.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
	ErrBadCMapDict    = _ef.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
)

func _gcf(_afb string) string {
	_cada := []rune(_afb)
	_dce := make([]string, len(_cada))
	for _dfc, _aed := range _cada {
		_dce[_dfc] = _f.Sprintf("\u0025\u0030\u0034\u0078", _aed)
	}
	return _f.Sprintf("\u003c\u0025\u0073\u003e", _fcg.Join(_dce, ""))
}

type CharCode uint32

func (_abfg *cMapParser) parseObject() (cmapObject, error) {
	_abfg.skipSpaces()
	for {
		_fab, _ecg := _abfg._defg.Peek(2)
		if _ecg != nil {
			return nil, _ecg
		}
		if _fab[0] == '%' {
			_abfg.parseComment()
			_abfg.skipSpaces()
			continue
		} else if _fab[0] == '/' {
			_gea, _dbg := _abfg.parseName()
			return _gea, _dbg
		} else if _fab[0] == '(' {
			_dcc, _ded := _abfg.parseString()
			return _dcc, _ded
		} else if _fab[0] == '[' {
			_aeee, _gaa := _abfg.parseArray()
			return _aeee, _gaa
		} else if (_fab[0] == '<') && (_fab[1] == '<') {
			_cccg, _ffcb := _abfg.parseDict()
			return _cccg, _ffcb
		} else if _fab[0] == '<' {
			_caebg, _gda := _abfg.parseHexString()
			return _caebg, _gda
		} else if _gg.IsDecimalDigit(_fab[0]) || (_fab[0] == '-' && _gg.IsDecimalDigit(_fab[1])) {
			_cbd, _edgc := _abfg.parseNumber()
			if _edgc != nil {
				return nil, _edgc
			}
			return _cbd, nil
		} else {
			_fgdc, _fce := _abfg.parseOperand()
			if _fce != nil {
				return nil, _fce
			}
			return _fgdc, nil
		}
	}
}
func _dgfd() cmapDict { return cmapDict{Dict: map[string]cmapObject{}} }
func (cmap *CMap) parseName() error {
	_bf := ""
	_cfd := false
	for _caf := 0; _caf < 20 && !_cfd; _caf++ {
		_aeec, _cga := cmap.parseObject()
		if _cga != nil {
			return _cga
		}
		switch _cef := _aeec.(type) {
		case cmapOperand:
			switch _cef.Operand {
			case "\u0064\u0065\u0066":
				_cfd = true
			default:
				_gc.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a\u0020\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076\u0020n\u0061\u006de\u003d\u0025\u0023\u0071", _aeec, _bf)
				if _bf != "" {
					_bf = _f.Sprintf("\u0025\u0073\u0020%\u0073", _bf, _cef.Operand)
				}
				_gc.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a \u0052\u0065\u0063\u006f\u0076\u0065\u0072e\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025\u0023\u0071", _bf)
			}
		case cmapName:
			_bf = _cef.Name
		}
	}
	if !_cfd {
		_gc.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0070\u0061\u0072\u0073\u0065N\u0061m\u0065:\u0020\u004e\u006f\u0020\u0064\u0065\u0066 ")
		return ErrBadCMap
	}
	cmap._fd = _bf
	return nil
}
func (cmap *CMap) BytesToCharcodes(data []byte) ([]CharCode, bool) {
	var _edg []CharCode
	if cmap._be == 8 {
		for _, _fdf := range data {
			_edg = append(_edg, CharCode(_fdf))
		}
		return _edg, true
	}
	for _fff := 0; _fff < len(data); {
		_eda, _eag, _fed := cmap.matchCode(data[_fff:])
		if !_fed {
			_gc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0074\u0020\u0069\u003d\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", _fff, data, string(data))
			return _edg, false
		}
		_edg = append(_edg, _eda)
		_fff += _eag
	}
	return _edg, true
}
func (cmap *CMap) CharcodeBytesToUnicode(data []byte) (string, int) {
	_fgd, _eb := cmap.BytesToCharcodes(data)
	if !_eb {
		_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065s\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069n\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0064\u0061\u0074\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078]\u0020\u0063\u006d\u0061\u0070=\u0025\u0073", data, cmap)
		return "", 0
	}
	_gbf := make([]string, len(_fgd))
	var _gfe []CharCode
	for _fa, _ccd := range _fgd {
		_ge, _bbe := cmap._ab[_ccd]
		if !_bbe {
			_gfe = append(_gfe, _ccd)
			_ge = MissingCodeString
		}
		_gbf[_fa] = _ge
	}
	_gfg := _fcg.Join(_gbf, "")
	if len(_gfe) > 0 {
		_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0068\u0061\u0072c\u006f\u0064\u0065\u0042y\u0074\u0065\u0073\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070\u002e\u000a"+"\u0009d\u0061t\u0061\u003d\u005b\u0025\u00200\u0032\u0078]\u003d\u0025\u0023\u0071\u000a"+"\u0009\u0063h\u0061\u0072\u0063o\u0064\u0065\u0073\u003d\u0025\u0030\u0032\u0078\u000a"+"\u0009\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003d\u0025\u0064\u0020%\u0030\u0032\u0078\u000a"+"\u0009\u0075\u006e\u0069\u0063\u006f\u0064\u0065\u003d`\u0025\u0073\u0060\u000a"+"\u0009\u0063\u006d\u0061\u0070\u003d\u0025\u0073", data, string(data), _fgd, len(_gfe), _gfe, _gfg, cmap)
	}
	return _gfg, len(_gfe)
}
func LoadCmapFromData(data []byte, isSimple bool) (*CMap, error) {
	_gc.Log.Trace("\u004c\u006fa\u0064\u0043\u006d\u0061\u0070\u0046\u0072\u006f\u006d\u0044\u0061\u0074\u0061\u003a\u0020\u0069\u0073\u0053\u0069\u006d\u0070\u006ce=\u0025\u0074", isSimple)
	cmap := _dc(isSimple)
	cmap.cMapParser = _cdd(data)
	_fe := cmap.parse()
	if _fe != nil {
		return nil, _fe
	}
	if len(cmap._gga) == 0 {
		if cmap._fda != "" {
			return cmap, nil
		}
		_gc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0063\u006d\u0061p=\u0025\u0073", cmap)
		return nil, ErrBadCMap
	}
	cmap.computeInverseMappings()
	return cmap, nil
}
func (_bcfa *cMapParser) parseName() (cmapName, error) {
	_ffgd := ""
	_fad := false
	for {
		_bcb, _deca := _bcfa._defg.Peek(1)
		if _deca == _ga.EOF {
			break
		}
		if _deca != nil {
			return cmapName{_ffgd}, _deca
		}
		if !_fad {
			if _bcb[0] == '/' {
				_fad = true
				_bcfa._defg.ReadByte()
			} else {
				_gc.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u004e\u0061\u006d\u0065\u0020\u0073\u0074a\u0072t\u0069n\u0067 \u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0028\u0025\u0020\u0078\u0029", _bcb, _bcb)
				return cmapName{_ffgd}, _f.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _bcb[0])
			}
		} else {
			if _gg.IsWhiteSpace(_bcb[0]) {
				break
			} else if (_bcb[0] == '/') || (_bcb[0] == '[') || (_bcb[0] == '(') || (_bcb[0] == ']') || (_bcb[0] == '<') || (_bcb[0] == '>') {
				break
			} else if _bcb[0] == '#' {
				_bgbe, _caff := _bcfa._defg.Peek(3)
				if _caff != nil {
					return cmapName{_ffgd}, _caff
				}
				_bcfa._defg.Discard(3)
				_ecb, _caff := _e.DecodeString(string(_bgbe[1:3]))
				if _caff != nil {
					return cmapName{_ffgd}, _caff
				}
				_ffgd += string(_ecb)
			} else {
				_abdg, _ := _bcfa._defg.ReadByte()
				_ffgd += string(_abdg)
			}
		}
	}
	return cmapName{_ffgd}, nil
}
