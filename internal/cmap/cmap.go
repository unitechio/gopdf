package cmap

import (
	_b "bufio"
	_fd "bytes"
	_d "encoding/hex"
	_bf "errors"
	_f "fmt"
	_fc "io"
	_db "sort"
	_g "strconv"
	_da "strings"
	_e "unicode/utf16"

	_dbf "bitbucket.org/shenghui0779/gopdf/common"
	_fg "bitbucket.org/shenghui0779/gopdf/core"
	_bc "bitbucket.org/shenghui0779/gopdf/internal/cmap/bcmaps"
)

func NewToUnicodeCMap(codeToRune map[CharCode]rune) *CMap {
	_bbg := make(map[CharCode]string, len(codeToRune))
	for _ca, _ccc := range codeToRune {
		_bbg[_ca] = string(_ccc)
	}
	cmap := &CMap{_dd: "\u0041d\u006fb\u0065\u002d\u0049\u0064\u0065n\u0074\u0069t\u0079\u002d\u0055\u0043\u0053", _cc: 2, _dba: 16, _gbf: CIDSystemInfo{Registry: "\u0041\u0064\u006fb\u0065", Ordering: "\u0055\u0043\u0053", Supplement: 0}, _ga: []Codespace{{Low: 0, High: 0xffff}}, _ed: _bbg, _cfb: make(map[string]CharCode, len(codeToRune)), _dc: make(map[CharCode]CharCode, len(codeToRune)), _gec: make(map[CharCode]CharCode, len(codeToRune))}
	cmap.computeInverseMappings()
	return cmap
}

type cmapString struct{ String string }

func _fde(_gg bool) *CMap {
	_fgc := 16
	if _gg {
		_fgc = 8
	}
	return &CMap{_dba: _fgc, _dc: make(map[CharCode]CharCode), _gec: make(map[CharCode]CharCode), _ed: make(map[CharCode]string), _cfb: make(map[string]CharCode)}
}
func (cmap *CMap) NBits() int { return cmap._dba }

const (
	_dgd = 100
	_geb = "\u000a\u002f\u0043\u0049\u0044\u0049\u006e\u0069\u0074\u0020\u002f\u0050\u0072\u006fc\u0053\u0065\u0074\u0020\u0066\u0069\u006e\u0064\u0072es\u006fu\u0072c\u0065 \u0062\u0065\u0067\u0069\u006e\u000a\u0031\u0032\u0020\u0064\u0069\u0063\u0074\u0020\u0062\u0065\u0067\u0069n\u000a\u0062\u0065\u0067\u0069\u006e\u0063\u006d\u0061\u0070\n\u002f\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u003c\u003c\u0020\u002f\u0052\u0065\u0067\u0069\u0073t\u0072\u0079\u0020\u0028\u0041\u0064\u006f\u0062\u0065\u0029\u0020\u002f\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0028\u0055\u0043\u0053)\u0020\u002f\u0053\u0075\u0070p\u006c\u0065\u006d\u0065\u006et\u0020\u0030\u0020\u003e\u003e\u0020\u0064\u0065\u0066\u000a\u002f\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u002f\u0041\u0064\u006f\u0062\u0065-\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0055\u0043\u0053\u0020\u0064\u0065\u0066\u000a\u002fC\u004d\u0061\u0070\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0064\u0065\u0066\u000a\u0031\u0020\u0062\u0065\u0067\u0069\u006e\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063e\u0072\u0061n\u0067\u0065\n\u003c\u0030\u0030\u0030\u0030\u003e\u0020<\u0046\u0046\u0046\u0046\u003e\u000a\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065\u000a"
	_acb = "\u0065\u006e\u0064\u0063\u006d\u0061\u0070\u000a\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u0063ur\u0072e\u006e\u0074\u0064\u0069\u0063\u0074\u0020\u002f\u0043\u004d\u0061\u0070 \u0064\u0065\u0066\u0069\u006e\u0065\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0070\u006fp\u000a\u0065\u006e\u0064\u000a\u0065\u006e\u0064\u000a"
)

func (cmap *CMap) CharcodeToUnicode(code CharCode) (string, bool) {
	if _be, _de := cmap._ed[code]; _de {
		return _be, true
	}
	return MissingCodeString, false
}
func (cmap *CMap) Bytes() []byte {
	_dbf.Log.Trace("\u0063\u006d\u0061\u0070.B\u0079\u0074\u0065\u0073\u003a\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", cmap.String())
	if len(cmap._gaa) > 0 {
		return cmap._gaa
	}
	cmap._gaa = []byte(_da.Join([]string{_geb, cmap.toBfData(), _acb}, "\u000a"))
	return cmap._gaa
}
func LoadPredefinedCMap(name string) (*CMap, error) {
	cmap, _dcf := _aae(name)
	if _dcf != nil {
		return nil, _dcf
	}
	if cmap._gb == "" {
		cmap.computeInverseMappings()
		return cmap, nil
	}
	_gfa, _dcf := _aae(cmap._gb)
	if _dcf != nil {
		return nil, _dcf
	}
	for _caf, _fgb := range _gfa._dc {
		if _, _ab := cmap._dc[_caf]; !_ab {
			cmap._dc[_caf] = _fgb
		}
	}
	cmap._ga = append(cmap._ga, _gfa._ga...)
	cmap.computeInverseMappings()
	return cmap, nil
}
func _cef(_bdb string) string {
	_gecc := []rune(_bdb)
	_daa := make([]string, len(_gecc))
	for _gce, _ceg := range _gecc {
		_daa[_gce] = _f.Sprintf("\u0025\u0030\u0034\u0078", _ceg)
	}
	return _f.Sprintf("\u003c\u0025\u0073\u003e", _da.Join(_daa, ""))
}
func (cmap *CMap) computeInverseMappings() {
	for _aaf, _egg := range cmap._dc {
		if _gfc, _fec := cmap._gec[_egg]; !_fec || (_fec && _gfc > _aaf) {
			cmap._gec[_egg] = _aaf
		}
	}
	for _aab, _cad := range cmap._ed {
		if _ccf, _gag := cmap._cfb[_cad]; !_gag || (_gag && _ccf > _aab) {
			cmap._cfb[_cad] = _aab
		}
	}
	_db.Slice(cmap._ga, func(_cfg, _ccd int) bool { return cmap._ga[_cfg].Low < cmap._ga[_ccd].Low })
}
func (_ge *CIDSystemInfo) String() string {
	return _f.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0025\u0030\u0033\u0064", _ge.Registry, _ge.Ordering, _ge.Supplement)
}
func (cmap *CMap) parseVersion() error {
	_dfg := ""
	_bga := false
	for _eba := 0; _eba < 3 && !_bga; _eba++ {
		_dfeb, _eea := cmap.parseObject()
		if _eea != nil {
			return _eea
		}
		switch _acd := _dfeb.(type) {
		case cmapOperand:
			switch _acd.Operand {
			case "\u0064\u0065\u0066":
				_bga = true
			default:
				_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0070\u0061\u0072\u0073\u0065\u0056e\u0072\u0073\u0069\u006f\u006e\u003a \u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076", _dfeb)
				return ErrBadCMap
			}
		case cmapInt:
			_dfg = _f.Sprintf("\u0025\u0064", _acd._dgca)
		case cmapFloat:
			_dfg = _f.Sprintf("\u0025\u0066", _acd._cbb)
		case cmapString:
			_dfg = _acd.String
		default:
			_dbf.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020p\u0061\u0072\u0073\u0065Ver\u0073io\u006e\u003a\u0020\u0042\u0061\u0064\u0020ty\u0070\u0065\u002e\u0020\u006f\u003d\u0025#\u0076", _dfeb)
		}
	}
	cmap._eb = _dfg
	return nil
}
func (cmap *CMap) WMode() (int, bool) { return cmap._bb._ebb, cmap._bb._bdf }

type cmapOperand struct{ Operand string }

func (cmap *CMap) parseCIDRange() error {
	for {
		_gege, _gcaa := cmap.parseObject()
		if _gcaa != nil {
			if _gcaa == _fc.EOF {
				break
			}
			return _gcaa
		}
		_aad, _deg := _gege.(cmapHexString)
		if !_deg {
			if _ecbc, _ffc := _gege.(cmapOperand); _ffc {
				if _ecbc.Operand == _cedg {
					return nil
				}
				return _bf.New("\u0063\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0072\u0076\u0061\u006c\u0020s\u0074\u0061\u0072\u0074\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
			}
		}
		_feaa := _egfg(_aad)
		_gege, _gcaa = cmap.parseObject()
		if _gcaa != nil {
			if _gcaa == _fc.EOF {
				break
			}
			return _gcaa
		}
		_bgd, _deg := _gege.(cmapHexString)
		if !_deg {
			return _bf.New("\u0063\u0069d\u0020\u0069\u006e\u0074e\u0072\u0076a\u006c\u0020\u0065\u006e\u0064\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_aad._ddc) != len(_bgd._ddc) {
			return _bf.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_cge := _egfg(_bgd)
		if _feaa > _cge {
			_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0049\u0044\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073t\u0061\u0072\u0074\u003d\u0030\u0078\u0025\u0030\u0032\u0078\u0020\u0065\u006e\u0064=\u0030x\u0025\u0030\u0032\u0078", _feaa, _cge)
			return ErrBadCMap
		}
		_gege, _gcaa = cmap.parseObject()
		if _gcaa != nil {
			if _gcaa == _fc.EOF {
				break
			}
			return _gcaa
		}
		_gbe, _deg := _gege.(cmapInt)
		if !_deg {
			return _bf.New("\u0063\u0069\u0064\u0020\u0073t\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		if _gbe._dgca < 0 {
			return _bf.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0063\u0069\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065")
		}
		_agg := _gbe._dgca
		for _bec := _feaa; _bec <= _cge; _bec++ {
			cmap._dc[_bec] = CharCode(_agg)
			_agg++
		}
		_dbf.Log.Trace("C\u0049\u0044\u0020\u0072\u0061\u006eg\u0065\u003a\u0020\u003c\u0030\u0078\u0025\u0058\u003e \u003c\u0030\u0078%\u0058>\u0020\u0025\u0064", _feaa, _cge, _gbe._dgca)
	}
	return nil
}

type cmapFloat struct{ _cbb float64 }

func (cmap *CMap) BytesToCharcodes(data []byte) ([]CharCode, bool) {
	var _fdd []CharCode
	if cmap._dba == 8 {
		for _, _abgf := range data {
			_fdd = append(_fdd, CharCode(_abgf))
		}
		return _fdd, true
	}
	for _bd := 0; _bd < len(data); {
		_bcd, _ee, _dfe := cmap.matchCode(data[_bd:])
		if !_dfe {
			_dbf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0074\u0020\u0069\u003d\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", _bd, data, string(data))
			return _fdd, false
		}
		_fdd = append(_fdd, _bcd)
		_bd += _ee
	}
	return _fdd, true
}

type cmapInt struct{ _dgca int64 }

func (cmap *CMap) Stream() (*_fg.PdfObjectStream, error) {
	if cmap._dga != nil {
		return cmap._dga, nil
	}
	_dae, _bbc := _fg.MakeStream(cmap.Bytes(), _fg.NewFlateEncoder())
	if _bbc != nil {
		return nil, _bbc
	}
	cmap._dga = _dae
	return cmap._dga, nil
}
func (_cbe *cMapParser) parseArray() (cmapArray, error) {
	_agc := cmapArray{}
	_agc.Array = []cmapObject{}
	_cbe._dgae.ReadByte()
	for {
		_cbe.skipSpaces()
		_aced, _gee := _cbe._dgae.Peek(1)
		if _gee != nil {
			return _agc, _gee
		}
		if _aced[0] == ']' {
			_cbe._dgae.ReadByte()
			break
		}
		_dcga, _gee := _cbe.parseObject()
		if _gee != nil {
			return _agc, _gee
		}
		_agc.Array = append(_agc.Array, _dcga)
	}
	return _agc, nil
}

type cmapDict struct {
	Dict map[string]cmapObject
}

func _egfg(_dbgc cmapHexString) CharCode {
	_gaac := CharCode(0)
	for _, _bdc := range _dbgc._ddc {
		_gaac <<= 8
		_gaac |= CharCode(_bdc)
	}
	return _gaac
}
func (cmap *CMap) parseCodespaceRange() error {
	for {
		_dcec, _bcg := cmap.parseObject()
		if _bcg != nil {
			if _bcg == _fc.EOF {
				break
			}
			return _bcg
		}
		_cbg, _gdca := _dcec.(cmapHexString)
		if !_gdca {
			if _eebb, _cgg := _dcec.(cmapOperand); _cgg {
				if _eebb.Operand == _cfc {
					return nil
				}
				return _bf.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
			}
		}
		_dcec, _bcg = cmap.parseObject()
		if _bcg != nil {
			if _bcg == _fc.EOF {
				break
			}
			return _bcg
		}
		_gdd, _gdca := _dcec.(cmapHexString)
		if !_gdca {
			return _bf.New("\u006e\u006f\u006e-\u0068\u0065\u0078\u0020\u0068\u0069\u0067\u0068")
		}
		if len(_cbg._ddc) != len(_gdd._ddc) {
			return _bf.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_gfad := _egfg(_cbg)
		_edcd := _egfg(_gdd)
		if _edcd < _gfad {
			_dbf.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0061d\u0020\u0063\u006fd\u0065\u0073\u0070\u0061\u0063\u0065\u002e\u0020\u006cow\u003d\u0030\u0078%\u0030\u0032x\u0020\u0068\u0069\u0067\u0068\u003d0\u0078\u00250\u0032\u0078", _gfad, _edcd)
			return ErrBadCMap
		}
		_fff := _gdd._ebdf
		_adg := Codespace{NumBytes: _fff, Low: _gfad, High: _edcd}
		cmap._ga = append(cmap._ga, _adg)
		_dbf.Log.Trace("\u0043\u006f\u0064e\u0073\u0070\u0061\u0063e\u0020\u006c\u006f\u0077\u003a\u0020\u0030x\u0025\u0058\u002c\u0020\u0068\u0069\u0067\u0068\u003a\u0020\u0030\u0078\u0025\u0058", _gfad, _edcd)
	}
	if len(cmap._ga) == 0 {
		_dbf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u0020\u0069\u006e\u0020\u0063ma\u0070\u002e")
		return ErrBadCMap
	}
	return nil
}
func (_fag *cMapParser) parseNumber() (cmapObject, error) {
	_ddgf, _eed := _fg.ParseNumber(_fag._dgae)
	if _eed != nil {
		return nil, _eed
	}
	switch _eafg := _ddgf.(type) {
	case *_fg.PdfObjectFloat:
		return cmapFloat{float64(*_eafg)}, nil
	case *_fg.PdfObjectInteger:
		return cmapInt{int64(*_eafg)}, nil
	}
	return nil, _f.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _ddgf)
}
func (_ffff *cMapParser) parseName() (cmapName, error) {
	_gfce := ""
	_efd := false
	for {
		_accg, _aedf := _ffff._dgae.Peek(1)
		if _aedf == _fc.EOF {
			break
		}
		if _aedf != nil {
			return cmapName{_gfce}, _aedf
		}
		if !_efd {
			if _accg[0] == '/' {
				_efd = true
				_ffff._dgae.ReadByte()
			} else {
				_dbf.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u004e\u0061\u006d\u0065\u0020\u0073\u0074a\u0072t\u0069n\u0067 \u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0028\u0025\u0020\u0078\u0029", _accg, _accg)
				return cmapName{_gfce}, _f.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _accg[0])
			}
		} else {
			if _fg.IsWhiteSpace(_accg[0]) {
				break
			} else if (_accg[0] == '/') || (_accg[0] == '[') || (_accg[0] == '(') || (_accg[0] == ']') || (_accg[0] == '<') || (_accg[0] == '>') {
				break
			} else if _accg[0] == '#' {
				_fdfa, _bdbc := _ffff._dgae.Peek(3)
				if _bdbc != nil {
					return cmapName{_gfce}, _bdbc
				}
				_ffff._dgae.Discard(3)
				_geff, _bdbc := _d.DecodeString(string(_fdfa[1:3]))
				if _bdbc != nil {
					return cmapName{_gfce}, _bdbc
				}
				_gfce += string(_geff)
			} else {
				_dge, _ := _ffff._dgae.ReadByte()
				_gfce += string(_dge)
			}
		}
	}
	return cmapName{_gfce}, nil
}
func LoadCmapFromDataCID(data []byte) (*CMap, error) { return LoadCmapFromData(data, false) }

type CMap struct {
	*cMapParser
	_dd  string
	_dba int
	_cc  int
	_eb  string
	_gb  string
	_gbf CIDSystemInfo
	_ga  []Codespace
	_dc  map[CharCode]CharCode
	_gec map[CharCode]CharCode
	_ed  map[CharCode]string
	_cfb map[string]CharCode
	_gaa []byte
	_dga *_fg.PdfObjectStream
	_bb  integer
}

func (cmap *CMap) String() string {
	_beg := cmap._gbf
	_dfee := []string{_f.Sprintf("\u006e\u0062\u0069\u0074\u0073\u003a\u0025\u0064", cmap._dba), _f.Sprintf("\u0074y\u0070\u0065\u003a\u0025\u0064", cmap._cc)}
	if cmap._eb != "" {
		_dfee = append(_dfee, _f.Sprintf("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0025\u0073", cmap._eb))
	}
	if cmap._gb != "" {
		_dfee = append(_dfee, _f.Sprintf("u\u0073\u0065\u0063\u006d\u0061\u0070\u003a\u0025\u0023\u0071", cmap._gb))
	}
	_dfee = append(_dfee, _f.Sprintf("\u0073\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u003a\u0025\u0073", _beg.String()))
	if len(cmap._ga) > 0 {
		_dfee = append(_dfee, _f.Sprintf("\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u003a\u0025\u0064", len(cmap._ga)))
	}
	if len(cmap._ed) > 0 {
		_dfee = append(_dfee, _f.Sprintf("\u0063\u006fd\u0065\u0054\u006fU\u006e\u0069\u0063\u006f\u0064\u0065\u003a\u0025\u0064", len(cmap._ed)))
	}
	return _f.Sprintf("\u0043\u004d\u0041P\u007b\u0025\u0023\u0071\u0020\u0025\u0073\u007d", cmap._dd, _da.Join(_dfee, "\u0020"))
}
func (_bfg *cMapParser) parseHexString() (cmapHexString, error) {
	_bfg._dgae.ReadByte()
	_ecg := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	_bad := _fd.Buffer{}
	for {
		_bfg.skipSpaces()
		_cefe, _egbc := _bfg._dgae.Peek(1)
		if _egbc != nil {
			return cmapHexString{}, _egbc
		}
		if _cefe[0] == '>' {
			_bfg._dgae.ReadByte()
			break
		}
		_becd, _ := _bfg._dgae.ReadByte()
		if _fd.IndexByte(_ecg, _becd) >= 0 {
			_bad.WriteByte(_becd)
		}
	}
	if _bad.Len()%2 == 1 {
		_dbf.Log.Debug("\u0070\u0061rs\u0065\u0048\u0065x\u0053\u0074\u0072\u0069ng:\u0020ap\u0070\u0065\u006e\u0064\u0069\u006e\u0067 '\u0030\u0027\u0020\u0074\u006f\u0020\u0025#\u0071", _bad.String())
		_bad.WriteByte('0')
	}
	_fbff := _bad.Len() / 2
	_eee, _ := _d.DecodeString(_bad.String())
	return cmapHexString{_ebdf: _fbff, _ddc: _eee}, nil
}

type cmapObject interface{}
type charRange struct {
	_ec CharCode
	_gf CharCode
}

func (_ggf *cMapParser) skipSpaces() (int, error) {
	_dee := 0
	for {
		_cee, _gbea := _ggf._dgae.Peek(1)
		if _gbea != nil {
			return 0, _gbea
		}
		if _fg.IsWhiteSpace(_cee[0]) {
			_ggf._dgae.ReadByte()
			_dee++
		} else {
			break
		}
	}
	return _dee, nil
}
func (cmap *CMap) parseName() error {
	_fccg := ""
	_cab := false
	for _fdb := 0; _fdb < 20 && !_cab; _fdb++ {
		_dbc, _eeg := cmap.parseObject()
		if _eeg != nil {
			return _eeg
		}
		switch _edfa := _dbc.(type) {
		case cmapOperand:
			switch _edfa.Operand {
			case "\u0064\u0065\u0066":
				_cab = true
			default:
				_dbf.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a\u0020\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076\u0020n\u0061\u006de\u003d\u0025\u0023\u0071", _dbc, _fccg)
				if _fccg != "" {
					_fccg = _f.Sprintf("\u0025\u0073\u0020%\u0073", _fccg, _edfa.Operand)
				}
				_dbf.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a \u0052\u0065\u0063\u006f\u0076\u0065\u0072e\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025\u0023\u0071", _fccg)
			}
		case cmapName:
			_fccg = _edfa.Name
		}
	}
	if !_cab {
		_dbf.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0070\u0061\u0072\u0073\u0065N\u0061m\u0065:\u0020\u004e\u006f\u0020\u0064\u0065\u0066 ")
		return ErrBadCMap
	}
	cmap._dd = _fccg
	return nil
}

const (
	_dgg  = "\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"
	_dbgg = "\u0062e\u0067\u0069\u006e\u0063\u006d\u0061p"
	_add  = "\u0065n\u0064\u0063\u006d\u0061\u0070"
	_bbfd = "\u0062\u0065\u0067\u0069nc\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0072\u0061\u006e\u0067\u0065"
	_cfc  = "\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065"
	_cdd  = "b\u0065\u0067\u0069\u006e\u0062\u0066\u0063\u0068\u0061\u0072"
	_gfe  = "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r"
	_gbc  = "\u0062\u0065\u0067i\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_facf = "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_aefb = "\u0062\u0065\u0067\u0069\u006e\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_cedg = "e\u006e\u0064\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_gcf  = "\u0075s\u0065\u0063\u006d\u0061\u0070"
	_ffeb = "\u0057\u004d\u006fd\u0065"
	_cda  = "\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065"
	_ddag = "\u0043\u004d\u0061\u0070\u0054\u0079\u0070\u0065"
	_gcb  = "C\u004d\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e"
)

func (cmap *CMap) matchCode(_gga []byte) (_fad CharCode, _fgd int, _ecb bool) {
	for _fcc := 0; _fcc < _c; _fcc++ {
		if _fcc < len(_gga) {
			_fad = _fad<<8 | CharCode(_gga[_fcc])
			_fgd++
		}
		_ecb = cmap.inCodespace(_fad, _fcc+1)
		if _ecb {
			return _fad, _fgd, true
		}
	}
	_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063o\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0020m\u0061t\u0063\u0068\u0065\u0073\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d=\u0025\u0023\u0071\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", _gga, string(_gga), cmap)
	return 0, 0, false
}

type Codespace struct {
	NumBytes int
	Low      CharCode
	High     CharCode
}

func (cmap *CMap) parseBfchar() error {
	for {
		_ffe, _edcb := cmap.parseObject()
		if _edcb != nil {
			if _edcb == _fc.EOF {
				break
			}
			return _edcb
		}
		var _fce CharCode
		switch _abb := _ffe.(type) {
		case cmapOperand:
			if _abb.Operand == _gfe {
				return nil
			}
			return _bf.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_fce = _egfg(_abb)
		default:
			return _bf.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_ffe, _edcb = cmap.parseObject()
		if _edcb != nil {
			if _edcb == _fc.EOF {
				break
			}
			return _edcb
		}
		var _ffea []rune
		switch _cca := _ffe.(type) {
		case cmapOperand:
			if _cca.Operand == _gfe {
				return nil
			}
			_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020o\u0070\u0065\u0072\u0061\u006e\u0064\u002e\u0020\u0025\u0023\u0076", _cca)
			return ErrBadCMap
		case cmapHexString:
			_ffea = _gab(_cca)
		case cmapName:
			_dbf.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0061\u006de\u002e \u0025\u0023\u0076", _cca)
			_ffea = []rune{MissingCodeRune}
		default:
			_dbf.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0074\u0079\u0070e\u002e \u0025\u0023\u0076", _ffe)
			return ErrBadCMap
		}
		cmap._ed[_fce] = string(_ffea)
	}
	return nil
}
func (_bgg *cMapParser) parseObject() (cmapObject, error) {
	_bgg.skipSpaces()
	for {
		_bca, _gagd := _bgg._dgae.Peek(2)
		if _gagd != nil {
			return nil, _gagd
		}
		if _bca[0] == '%' {
			_bgg.parseComment()
			_bgg.skipSpaces()
			continue
		} else if _bca[0] == '/' {
			_ede, _gdce := _bgg.parseName()
			return _ede, _gdce
		} else if _bca[0] == '(' {
			_def, _fbc := _bgg.parseString()
			return _def, _fbc
		} else if _bca[0] == '[' {
			_bddd, _cfa := _bgg.parseArray()
			return _bddd, _cfa
		} else if (_bca[0] == '<') && (_bca[1] == '<') {
			_gae, _gfada := _bgg.parseDict()
			return _gae, _gfada
		} else if _bca[0] == '<' {
			_gbfe, _ace := _bgg.parseHexString()
			return _gbfe, _ace
		} else if _fg.IsDecimalDigit(_bca[0]) || (_bca[0] == '-' && _fg.IsDecimalDigit(_bca[1])) {
			_gffa, _gcg := _bgg.parseNumber()
			if _gcg != nil {
				return nil, _gcg
			}
			return _gffa, nil
		} else {
			_fga, _bbce := _bgg.parseOperand()
			if _bbce != nil {
				return nil, _bbce
			}
			return _fga, nil
		}
	}
}
func _gcc(_dcd string) rune { _ddf := []rune(_dcd); return _ddf[len(_ddf)-1] }
func (_gacf *cMapParser) parseString() (cmapString, error) {
	_gacf._dgae.ReadByte()
	_ead := _fd.Buffer{}
	_ded := 1
	for {
		_cdbe, _fcd := _gacf._dgae.Peek(1)
		if _fcd != nil {
			return cmapString{_ead.String()}, _fcd
		}
		if _cdbe[0] == '\\' {
			_gacf._dgae.ReadByte()
			_eec, _cce := _gacf._dgae.ReadByte()
			if _cce != nil {
				return cmapString{_ead.String()}, _cce
			}
			if _fg.IsOctalDigit(_eec) {
				_fgcb, _eege := _gacf._dgae.Peek(2)
				if _eege != nil {
					return cmapString{_ead.String()}, _eege
				}
				var _dfcg []byte
				_dfcg = append(_dfcg, _eec)
				for _, _ccda := range _fgcb {
					if _fg.IsOctalDigit(_ccda) {
						_dfcg = append(_dfcg, _ccda)
					} else {
						break
					}
				}
				_gacf._dgae.Discard(len(_dfcg) - 1)
				_dbf.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _dfcg)
				_aacc, _eege := _g.ParseUint(string(_dfcg), 8, 32)
				if _eege != nil {
					return cmapString{_ead.String()}, _eege
				}
				_ead.WriteByte(byte(_aacc))
				continue
			}
			switch _eec {
			case 'n':
				_ead.WriteByte('\n')
			case 'r':
				_ead.WriteByte('\r')
			case 't':
				_ead.WriteByte('\t')
			case 'b':
				_ead.WriteByte('\b')
			case 'f':
				_ead.WriteByte('\f')
			case '(':
				_ead.WriteByte('(')
			case ')':
				_ead.WriteByte(')')
			case '\\':
				_ead.WriteByte('\\')
			}
			continue
		} else if _cdbe[0] == '(' {
			_ded++
		} else if _cdbe[0] == ')' {
			_ded--
			if _ded == 0 {
				_gacf._dgae.ReadByte()
				break
			}
		}
		_agd, _ := _gacf._dgae.ReadByte()
		_ead.WriteByte(_agd)
	}
	return cmapString{_ead.String()}, nil
}

type cMapParser struct{ _dgae *_b.Reader }
type CharCode uint32

func (_bfgg *cMapParser) parseOperand() (cmapOperand, error) {
	_ddfa := cmapOperand{}
	_fgbc := _fd.Buffer{}
	for {
		_agcb, _aeg := _bfgg._dgae.Peek(1)
		if _aeg != nil {
			if _aeg == _fc.EOF {
				break
			}
			return _ddfa, _aeg
		}
		if _fg.IsDelimiter(_agcb[0]) {
			break
		}
		if _fg.IsWhiteSpace(_agcb[0]) {
			break
		}
		_fecb, _ := _bfgg._dgae.ReadByte()
		_fgbc.WriteByte(_fecb)
	}
	if _fgbc.Len() == 0 {
		return _ddfa, _f.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	_ddfa.Operand = _fgbc.String()
	return _ddfa, nil
}
func (cmap *CMap) parseWMode() error {
	var _gfda int
	_ggcc := false
	for _fged := 0; _fged < 3 && !_ggcc; _fged++ {
		_bgc, _ffb := cmap.parseObject()
		if _ffb != nil {
			return _ffb
		}
		switch _egb := _bgc.(type) {
		case cmapOperand:
			switch _egb.Operand {
			case "\u0064\u0065\u0066":
				_ggcc = true
			default:
				_dbf.Log.Error("\u0070\u0061\u0072\u0073\u0065\u0057\u004d\u006f\u0064\u0065:\u0020\u0073\u0074\u0061\u0074\u0065\u0020e\u0072\u0072\u006f\u0072\u002e\u0020\u006f\u003d\u0025\u0023\u0076", _bgc)
				return ErrBadCMap
			}
		case cmapInt:
			_gfda = int(_egb._dgca)
		}
	}
	cmap._bb = integer{_bdf: true, _ebb: _gfda}
	return nil
}

const (
	_c                = 4
	MissingCodeRune   = '\ufffd'
	MissingCodeString = string(MissingCodeRune)
)

type CIDSystemInfo struct {
	Registry   string
	Ordering   string
	Supplement int
}

func (cmap *CMap) parseSystemInfo() error {
	_dad := false
	_aaa := false
	_dda := ""
	_bbgg := false
	_cdb := CIDSystemInfo{}
	for _fac := 0; _fac < 50 && !_bbgg; _fac++ {
		_bbe, _gdf := cmap.parseObject()
		if _gdf != nil {
			return _gdf
		}
		switch _edb := _bbe.(type) {
		case cmapDict:
			_geg := _edb.Dict
			_facg, _ged := _geg["\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"]
			if !_ged {
				_dbf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_aef, _ged := _facg.(cmapString)
			if !_ged {
				_dbf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_cdb.Registry = _aef.String
			_facg, _ged = _geg["\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"]
			if !_ged {
				_dbf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_aef, _ged = _facg.(cmapString)
			if !_ged {
				_dbf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_cdb.Ordering = _aef.String
			_fadc, _ged := _geg["\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"]
			if !_ged {
				_dbf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_edcf, _ged := _fadc.(cmapInt)
			if !_ged {
				_dbf.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_cdb.Supplement = int(_edcf._dgca)
			_bbgg = true
		case cmapOperand:
			switch _edb.Operand {
			case "\u0062\u0065\u0067i\u006e":
				_dad = true
			case "\u0065\u006e\u0064":
				_bbgg = true
			case "\u0064\u0065\u0066":
				_aaa = false
			}
		case cmapName:
			if _dad {
				_dda = _edb.Name
				_aaa = true
			}
		case cmapString:
			if _aaa {
				switch _dda {
				case "\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079":
					_cdb.Registry = _edb.String
				case "\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067":
					_cdb.Ordering = _edb.String
				}
			}
		case cmapInt:
			if _aaa {
				switch _dda {
				case "\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074":
					_cdb.Supplement = int(_edb._dgca)
				}
			}
		}
	}
	if !_bbgg {
		_dbf.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u006c\u0079")
		return ErrBadCMap
	}
	cmap._gbf = _cdb
	return nil
}
func (cmap *CMap) CIDToCharcode(cid CharCode) (CharCode, bool) {
	_fgge, _cdg := cmap._gec[cid]
	return _fgge, _cdg
}
func (cmap *CMap) inCodespace(_cgd CharCode, _fdde int) bool {
	for _, _beb := range cmap._ga {
		if _beb.Low <= _cgd && _cgd <= _beb.High && _fdde == _beb.NumBytes {
			return true
		}
	}
	return false
}
func (_facd *cMapParser) parseComment() (string, error) {
	var _ffd _fd.Buffer
	_, _effg := _facd.skipSpaces()
	if _effg != nil {
		return _ffd.String(), _effg
	}
	_feac := true
	for {
		_afb, _dcb := _facd._dgae.Peek(1)
		if _dcb != nil {
			_dbf.Log.Debug("p\u0061r\u0073\u0065\u0043\u006f\u006d\u006d\u0065\u006et\u003a\u0020\u0065\u0072r=\u0025\u0076", _dcb)
			return _ffd.String(), _dcb
		}
		if _feac && _afb[0] != '%' {
			return _ffd.String(), ErrBadCMapComment
		}
		_feac = false
		if (_afb[0] != '\r') && (_afb[0] != '\n') {
			_ggb, _ := _facd._dgae.ReadByte()
			_ffd.WriteByte(_ggb)
		} else {
			break
		}
	}
	return _ffd.String(), nil
}
func _bef(_cff, _acgc int) int {
	if _cff < _acgc {
		return _cff
	}
	return _acgc
}
func (_aeb *cMapParser) parseDict() (cmapDict, error) {
	_dbf.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_dca := _bea()
	_ecf, _ := _aeb._dgae.ReadByte()
	if _ecf != '<' {
		return _dca, ErrBadCMapDict
	}
	_ecf, _ = _aeb._dgae.ReadByte()
	if _ecf != '<' {
		return _dca, ErrBadCMapDict
	}
	for {
		_aeb.skipSpaces()
		_facgf, _cgb := _aeb._dgae.Peek(2)
		if _cgb != nil {
			return _dca, _cgb
		}
		if (_facgf[0] == '>') && (_facgf[1] == '>') {
			_aeb._dgae.ReadByte()
			_aeb._dgae.ReadByte()
			break
		}
		_eaf, _cgb := _aeb.parseName()
		_dbf.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _eaf.Name)
		if _cgb != nil {
			_dbf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u002e\u0020\u0065\u0072r=\u0025\u0076", _cgb)
			return _dca, _cgb
		}
		_aeb.skipSpaces()
		_gbec, _cgb := _aeb.parseObject()
		if _cgb != nil {
			return _dca, _cgb
		}
		_dca.Dict[_eaf.Name] = _gbec
		_aeb.skipSpaces()
		_facgf, _cgb = _aeb._dgae.Peek(3)
		if _cgb != nil {
			return _dca, _cgb
		}
		if string(_facgf) == "\u0064\u0065\u0066" {
			_aeb._dgae.Discard(3)
		}
	}
	return _dca, nil
}
func _bea() cmapDict { return cmapDict{Dict: map[string]cmapObject{}} }

var (
	ErrBadCMap        = _bf.New("\u0062\u0061\u0064\u0020\u0063\u006d\u0061\u0070")
	ErrBadCMapComment = _bf.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
	ErrBadCMapDict    = _bf.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
)

func _gefb(_ebfb cmapHexString) rune {
	_afa := _gab(_ebfb)
	if _eda := len(_afa); _eda == 0 {
		_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054o\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065c\u0074\u0065\u0064\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0072u\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023\u0076", _ebfb)
		return MissingCodeRune
	}
	if len(_afa) > 1 {
		_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0078\u0061\u0063\u0074\u006c\u0079\u0020\u006f\u006e\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023v\u0020\u002d\u003e\u0020\u0025#\u0076", _ebfb, _afa)
	}
	return _afa[0]
}
func IsPredefinedCMap(name string) bool { return _bc.AssetExists(name) }
func LoadCmapFromData(data []byte, isSimple bool) (*CMap, error) {
	_dbf.Log.Trace("\u004c\u006fa\u0064\u0043\u006d\u0061\u0070\u0046\u0072\u006f\u006d\u0044\u0061\u0074\u0061\u003a\u0020\u0069\u0073\u0053\u0069\u006d\u0070\u006ce=\u0025\u0074", isSimple)
	cmap := _fde(isSimple)
	cmap.cMapParser = _dcg(data)
	_ggd := cmap.parse()
	if _ggd != nil {
		return nil, _ggd
	}
	if len(cmap._ga) == 0 {
		if cmap._gb != "" {
			return cmap, nil
		}
		_dbf.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0063\u006d\u0061p=\u0025\u0073", cmap)
		return nil, ErrBadCMap
	}
	cmap.computeInverseMappings()
	return cmap, nil
}

type integer struct {
	_bdf bool
	_ebb int
}

func (cmap *CMap) Name() string { return cmap._dd }
func (cmap *CMap) CharcodeToCID(code CharCode) (CharCode, bool) {
	_eff, _deb := cmap._dc[code]
	return _eff, _deb
}
func _aae(_fe string) (*CMap, error) {
	_df, _cg := _bc.Asset(_fe)
	if _cg != nil {
		return nil, _cg
	}
	return LoadCmapFromDataCID(_df)
}
func (cmap *CMap) CharcodeBytesToUnicode(data []byte) (string, int) {
	_gfcd, _fgg := cmap.BytesToCharcodes(data)
	if !_fgg {
		_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065s\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069n\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0064\u0061\u0074\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078]\u0020\u0063\u006d\u0061\u0070=\u0025\u0073", data, cmap)
		return "", 0
	}
	_cga := make([]string, len(_gfcd))
	var _dbff []CharCode
	for _gc, _ccg := range _gfcd {
		_ba, _ddg := cmap._ed[_ccg]
		if !_ddg {
			_dbff = append(_dbff, _ccg)
			_ba = MissingCodeString
		}
		_cga[_gc] = _ba
	}
	_gcd := _da.Join(_cga, "")
	if len(_dbff) > 0 {
		_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0068\u0061\u0072c\u006f\u0064\u0065\u0042y\u0074\u0065\u0073\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070\u002e\u000a"+"\u0009d\u0061t\u0061\u003d\u005b\u0025\u00200\u0032\u0078]\u003d\u0025\u0023\u0071\u000a"+"\u0009\u0063h\u0061\u0072\u0063o\u0064\u0065\u0073\u003d\u0025\u0030\u0032\u0078\u000a"+"\u0009\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003d\u0025\u0064\u0020%\u0030\u0032\u0078\u000a"+"\u0009\u0075\u006e\u0069\u0063\u006f\u0064\u0065\u003d`\u0025\u0073\u0060\u000a"+"\u0009\u0063\u006d\u0061\u0070\u003d\u0025\u0073", data, string(data), _gfcd, len(_dbff), _dbff, _gcd, cmap)
	}
	return _gcd, len(_dbff)
}
func (cmap *CMap) CIDSystemInfo() CIDSystemInfo { return cmap._gbf }
func NewCIDSystemInfo(obj _fg.PdfObject) (_fb CIDSystemInfo, _ad error) {
	_dg, _cf := _fg.GetDict(obj)
	if !_cf {
		return CIDSystemInfo{}, _fg.ErrTypeError
	}
	_cd, _cf := _fg.GetStringVal(_dg.Get("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"))
	if !_cf {
		return CIDSystemInfo{}, _fg.ErrTypeError
	}
	_cdc, _cf := _fg.GetStringVal(_dg.Get("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"))
	if !_cf {
		return CIDSystemInfo{}, _fg.ErrTypeError
	}
	_aac, _cf := _fg.GetIntVal(_dg.Get("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"))
	if !_cf {
		return CIDSystemInfo{}, _fg.ErrTypeError
	}
	return CIDSystemInfo{Registry: _cd, Ordering: _cdc, Supplement: _aac}, nil
}
func (cmap *CMap) Type() int { return cmap._cc }
func (cmap *CMap) StringToCID(s string) (CharCode, bool) {
	_gecg, _ac := cmap._cfb[s]
	return _gecg, _ac
}

type cmapHexString struct {
	_ebdf int
	_ddc  []byte
}

func _dcg(_caef []byte) *cMapParser {
	_aed := cMapParser{}
	_cfcg := _fd.NewBuffer(_caef)
	_aed._dgae = _b.NewReader(_cfcg)
	return &_aed
}
func _gab(_fgbdd cmapHexString) []rune {
	if len(_fgbdd._ddc) == 1 {
		return []rune{rune(_fgbdd._ddc[0])}
	}
	_cfbd := _fgbdd._ddc
	if len(_cfbd)%2 != 0 {
		_cfbd = append(_cfbd, 0)
		_dbf.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u0073\u002e\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0073\u0068\u0065\u0078\u003d\u0025#\u0076\u0020\u0074\u006f\u0020\u0025\u002b\u0076", _fgbdd, _cfbd)
	}
	_bde := len(_cfbd) >> 1
	_fcbb := make([]uint16, _bde)
	for _aaae := 0; _aaae < _bde; _aaae++ {
		_fcbb[_aaae] = uint16(_cfbd[_aaae<<1])<<8 + uint16(_cfbd[_aaae<<1+1])
	}
	_bdab := _e.Decode(_fcbb)
	return _bdab
}
func (cmap *CMap) parseType() error {
	_bge := 0
	_fea := false
	for _gfdd := 0; _gfdd < 3 && !_fea; _gfdd++ {
		_bdbe, _bgb := cmap.parseObject()
		if _bgb != nil {
			return _bgb
		}
		switch _efff := _bdbe.(type) {
		case cmapOperand:
			switch _efff.Operand {
			case "\u0064\u0065\u0066":
				_fea = true
			default:
				_dbf.Log.Error("\u0070\u0061r\u0073\u0065\u0054\u0079\u0070\u0065\u003a\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u006f=%\u0023\u0076", _bdbe)
				return ErrBadCMap
			}
		case cmapInt:
			_bge = int(_efff._dgca)
		}
	}
	cmap._cc = _bge
	return nil
}
func (cmap *CMap) parse() error {
	var _dcdd cmapObject
	for {
		_bda, _ggc := cmap.parseObject()
		if _ggc != nil {
			if _ggc == _fc.EOF {
				break
			}
			_dbf.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u004d\u0061\u0070\u003a\u0020\u0025\u0076", _ggc)
			return _ggc
		}
		switch _bcda := _bda.(type) {
		case cmapOperand:
			_fdg := _bcda
			switch _fdg.Operand {
			case _bbfd:
				_fdgf := cmap.parseCodespaceRange()
				if _fdgf != nil {
					return _fdgf
				}
			case _aefb:
				_ff := cmap.parseCIDRange()
				if _ff != nil {
					return _ff
				}
			case _cdd:
				_dbag := cmap.parseBfchar()
				if _dbag != nil {
					return _dbag
				}
			case _gbc:
				_bbf := cmap.parseBfrange()
				if _bbf != nil {
					return _bbf
				}
			case _gcf:
				if _dcdd == nil {
					_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u0073\u0065\u0063m\u0061\u0070\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f \u0061\u0072\u0067")
					return ErrBadCMap
				}
				_gfae, _geba := _dcdd.(cmapName)
				if !_geba {
					_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0075\u0073\u0065\u0063\u006d\u0061\u0070\u0020\u0061\u0072\u0067\u0020\u006eo\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u0020\u0025\u0023\u0076", _dcdd)
					return ErrBadCMap
				}
				cmap._gb = _gfae.Name
			case _dgg:
				_fadg := cmap.parseSystemInfo()
				if _fadg != nil {
					return _fadg
				}
			}
		case cmapName:
			_gff := _bcda
			switch _gff.Name {
			case _dgg:
				_gca := cmap.parseSystemInfo()
				if _gca != nil {
					return _gca
				}
			case _cda:
				_daaa := cmap.parseName()
				if _daaa != nil {
					return _daaa
				}
			case _ddag:
				_fecg := cmap.parseType()
				if _fecg != nil {
					return _fecg
				}
			case _gcb:
				_ebg := cmap.parseVersion()
				if _ebg != nil {
					return _ebg
				}
			case _ffeb:
				if _ggc = cmap.parseWMode(); _ggc != nil {
					return _ggc
				}
			}
		}
		_dcdd = _bda
	}
	return nil
}

type fbRange struct {
	_ag CharCode
	_aa CharCode
	_eg string
}
type cmapName struct{ Name string }

func (cmap *CMap) parseBfrange() error {
	for {
		var _efb CharCode
		_dgf, _aggd := cmap.parseObject()
		if _aggd != nil {
			if _aggd == _fc.EOF {
				break
			}
			return _aggd
		}
		switch _adf := _dgf.(type) {
		case cmapOperand:
			if _adf.Operand == _facf {
				return nil
			}
			return _bf.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_efb = _egfg(_adf)
		default:
			return _bf.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		var _ecc CharCode
		_dgf, _aggd = cmap.parseObject()
		if _aggd != nil {
			if _aggd == _fc.EOF {
				break
			}
			return _aggd
		}
		switch _fgbd := _dgf.(type) {
		case cmapOperand:
			_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065\u0020\u0062\u0066r\u0061\u006e\u0067\u0065\u0020\u0074\u0072i\u0070\u006c\u0065\u0074")
			return ErrBadCMap
		case cmapHexString:
			_ecc = _egfg(_fgbd)
			if _ecc > 0xffff {
				_ecc = 0xffff
			}
		default:
			_dbf.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _dgf)
			return ErrBadCMap
		}
		_dgf, _aggd = cmap.parseObject()
		if _aggd != nil {
			if _aggd == _fc.EOF {
				break
			}
			return _aggd
		}
		switch _af := _dgf.(type) {
		case cmapArray:
			if len(_af.Array) != int(_ecc-_efb)+1 {
				_dbf.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0069\u0074\u0065\u006d\u0073\u0020\u0069\u006e\u0020a\u0072\u0072\u0061\u0079")
				return ErrBadCMap
			}
			for _gffc := _efb; _gffc <= _ecc; _gffc++ {
				_acf := _af.Array[_gffc-_efb]
				_dgaf, _cegc := _acf.(cmapHexString)
				if !_cegc {
					return _bf.New("\u006e\u006f\u006e-h\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				}
				_eef := _gab(_dgaf)
				cmap._ed[_gffc] = string(_eef)
			}
		case cmapHexString:
			_gac := _gab(_af)
			_ced := len(_gac)
			for _ffa := _efb; _ffa <= _ecc; _ffa++ {
				cmap._ed[_ffa] = string(_gac)
				if _ced > 0 {
					_gac[_ced-1]++
				} else {
					_dbf.Log.Debug("\u004e\u006f\u0020c\u006d\u0061\u0070\u0020\u0074\u0061\u0072\u0067\u0065\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0066\u006f\u0072\u0020\u0025\u0023\u0076", _ffa)
				}
				if _ffa == 1<<32-1 {
					break
				}
			}
		default:
			_dbf.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _dgf)
			return ErrBadCMap
		}
	}
	return nil
}
func (cmap *CMap) toBfData() string {
	if len(cmap._ed) == 0 {
		return ""
	}
	_gd := make([]CharCode, 0, len(cmap._ed))
	for _egf := range cmap._ed {
		_gd = append(_gd, _egf)
	}
	_db.Slice(_gd, func(_ccgg, _dab int) bool { return _gd[_ccgg] < _gd[_dab] })
	var _gdc []charRange
	_gef := charRange{_gd[0], _gd[0]}
	_bab := cmap._ed[_gd[0]]
	for _, _agb := range _gd[1:] {
		_gaaf := cmap._ed[_agb]
		if _agb == _gef._gf+1 && _gcc(_gaaf) == _gcc(_bab)+1 {
			_gef._gf = _agb
		} else {
			_gdc = append(_gdc, _gef)
			_gef._ec, _gef._gf = _agb, _agb
		}
		_bab = _gaaf
	}
	_gdc = append(_gdc, _gef)
	var _dce []CharCode
	var _eag []fbRange
	for _, _ce := range _gdc {
		if _ce._ec == _ce._gf {
			_dce = append(_dce, _ce._ec)
		} else {
			_eag = append(_eag, fbRange{_ag: _ce._ec, _aa: _ce._gf, _eg: cmap._ed[_ce._ec]})
		}
	}
	_dbf.Log.Trace("\u0063\u0068ar\u0052\u0061\u006eg\u0065\u0073\u003d\u0025d f\u0062Ch\u0061\u0072\u0073\u003d\u0025\u0064\u0020fb\u0052\u0061\u006e\u0067\u0065\u0073\u003d%\u0064", len(_gdc), len(_dce), len(_eag))
	var _eeb []string
	if len(_dce) > 0 {
		_dbg := (len(_dce) + _dgd - 1) / _dgd
		for _fcb := 0; _fcb < _dbg; _fcb++ {
			_gadc := _bef(len(_dce)-_fcb*_dgd, _dgd)
			_eeb = append(_eeb, _f.Sprintf("\u0025\u0064\u0020\u0062\u0065\u0067\u0069\u006e\u0062f\u0063\u0068\u0061\u0072", _gadc))
			for _cbc := 0; _cbc < _gadc; _cbc++ {
				_bbgb := _dce[_fcb*_dgd+_cbc]
				_feca := cmap._ed[_bbgb]
				_eeb = append(_eeb, _f.Sprintf("\u003c%\u0030\u0034\u0078\u003e\u0020\u0025s", _bbgb, _cef(_feca)))
			}
			_eeb = append(_eeb, "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r")
		}
	}
	if len(_eag) > 0 {
		_bg := (len(_eag) + _dgd - 1) / _dgd
		for _daeg := 0; _daeg < _bg; _daeg++ {
			_ggg := _bef(len(_eag)-_daeg*_dgd, _dgd)
			_eeb = append(_eeb, _f.Sprintf("\u0025d\u0020b\u0065\u0067\u0069\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065", _ggg))
			for _bdd := 0; _bdd < _ggg; _bdd++ {
				_fdf := _eag[_daeg*_dgd+_bdd]
				_eeb = append(_eeb, _f.Sprintf("\u003c%\u00304\u0078\u003e\u003c\u0025\u0030\u0034\u0078\u003e\u0020\u0025\u0073", _fdf._ag, _fdf._aa, _cef(_fdf._eg)))
			}
			_eeb = append(_eeb, "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065")
		}
	}
	return _da.Join(_eeb, "\u000a")
}

type cmapArray struct{ Array []cmapObject }
