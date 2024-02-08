package cmap

import (
	_a "bufio"
	_ec "bytes"
	_df "encoding/hex"
	_be "errors"
	_g "fmt"
	_c "io"
	_ee "sort"
	_b "strconv"
	_ag "strings"
	_e "unicode/utf16"

	_bc "bitbucket.org/shenghui0779/gopdf/common"
	_ba "bitbucket.org/shenghui0779/gopdf/core"
	_ga "bitbucket.org/shenghui0779/gopdf/internal/cmap/bcmaps"
)

func LoadCmapFromDataCID(data []byte) (*CMap, error) { return LoadCmapFromData(data, false) }
func NewToUnicodeCMap(codeToRune map[CharCode]rune) *CMap {
	_dg := make(map[CharCode]string, len(codeToRune))
	for _fae, _geg := range codeToRune {
		_dg[_fae] = string(_geg)
	}
	cmap := &CMap{_ded: "\u0041d\u006fb\u0065\u002d\u0049\u0064\u0065n\u0074\u0069t\u0079\u002d\u0055\u0043\u0053", _beb: 2, _bf: 16, _dd: CIDSystemInfo{Registry: "\u0041\u0064\u006fb\u0065", Ordering: "\u0055\u0043\u0053", Supplement: 0}, _dbb: []Codespace{{Low: 0, High: 0xffff}}, _cbg: _dg, _gcf: make(map[string]CharCode, len(codeToRune)), _fa: make(map[CharCode]CharCode, len(codeToRune)), _bd: make(map[CharCode]CharCode, len(codeToRune))}
	cmap.computeInverseMappings()
	return cmap
}

const (
	_efa = 100
	_gfb = "\u000a\u002f\u0043\u0049\u0044\u0049\u006e\u0069\u0074\u0020\u002f\u0050\u0072\u006fc\u0053\u0065\u0074\u0020\u0066\u0069\u006e\u0064\u0072es\u006fu\u0072c\u0065 \u0062\u0065\u0067\u0069\u006e\u000a\u0031\u0032\u0020\u0064\u0069\u0063\u0074\u0020\u0062\u0065\u0067\u0069n\u000a\u0062\u0065\u0067\u0069\u006e\u0063\u006d\u0061\u0070\n\u002f\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u003c\u003c\u0020\u002f\u0052\u0065\u0067\u0069\u0073t\u0072\u0079\u0020\u0028\u0041\u0064\u006f\u0062\u0065\u0029\u0020\u002f\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u0028\u0055\u0043\u0053)\u0020\u002f\u0053\u0075\u0070p\u006c\u0065\u006d\u0065\u006et\u0020\u0030\u0020\u003e\u003e\u0020\u0064\u0065\u0066\u000a\u002f\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u002f\u0041\u0064\u006f\u0062\u0065-\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0055\u0043\u0053\u0020\u0064\u0065\u0066\u000a\u002fC\u004d\u0061\u0070\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0064\u0065\u0066\u000a\u0031\u0020\u0062\u0065\u0067\u0069\u006e\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063e\u0072\u0061n\u0067\u0065\n\u003c\u0030\u0030\u0030\u0030\u003e\u0020<\u0046\u0046\u0046\u0046\u003e\u000a\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065\u000a"
	_dag = "\u0065\u006e\u0064\u0063\u006d\u0061\u0070\u000a\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065\u0020\u0063ur\u0072e\u006e\u0074\u0064\u0069\u0063\u0074\u0020\u002f\u0043\u004d\u0061\u0070 \u0064\u0065\u0066\u0069\u006e\u0065\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0070\u006fp\u000a\u0065\u006e\u0064\u000a\u0065\u006e\u0064\u000a"
)

type cmapInt struct{ _dbae int64 }

func (cmap *CMap) String() string {
	_gf := cmap._dd
	_dc := []string{_g.Sprintf("\u006e\u0062\u0069\u0074\u0073\u003a\u0025\u0064", cmap._bf), _g.Sprintf("\u0074y\u0070\u0065\u003a\u0025\u0064", cmap._beb)}
	if cmap._gg != "" {
		_dc = append(_dc, _g.Sprintf("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0025\u0073", cmap._gg))
	}
	if cmap._db != "" {
		_dc = append(_dc, _g.Sprintf("u\u0073\u0065\u0063\u006d\u0061\u0070\u003a\u0025\u0023\u0071", cmap._db))
	}
	_dc = append(_dc, _g.Sprintf("\u0073\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u003a\u0025\u0073", _gf.String()))
	if len(cmap._dbb) > 0 {
		_dc = append(_dc, _g.Sprintf("\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u003a\u0025\u0064", len(cmap._dbb)))
	}
	if len(cmap._cbg) > 0 {
		_dc = append(_dc, _g.Sprintf("\u0063\u006fd\u0065\u0054\u006fU\u006e\u0069\u0063\u006f\u0064\u0065\u003a\u0025\u0064", len(cmap._cbg)))
	}
	return _g.Sprintf("\u0043\u004d\u0041P\u007b\u0025\u0023\u0071\u0020\u0025\u0073\u007d", cmap._ded, _ag.Join(_dc, "\u0020"))
}
func _gga(_dge string) (*CMap, error) {
	_bed, _da := _ga.Asset(_dge)
	if _da != nil {
		return nil, _da
	}
	return LoadCmapFromDataCID(_bed)
}
func (cmap *CMap) WMode() (int, bool) { return cmap._eb._beee, cmap._eb._efaa }

type cMapParser struct{ _cfg *_a.Reader }

func _ebe(_efc string) string {
	_cg := []rune(_efc)
	_dcc := make([]string, len(_cg))
	for _abe, _fbba := range _cg {
		_dcc[_abe] = _g.Sprintf("\u0025\u0030\u0034\u0078", _fbba)
	}
	return _g.Sprintf("\u003c\u0025\u0073\u003e", _ag.Join(_dcc, ""))
}

const (
	_bb               = 4
	MissingCodeRune   = '\ufffd'
	MissingCodeString = string(MissingCodeRune)
)

type CIDSystemInfo struct {
	Registry   string
	Ordering   string
	Supplement int
}

func (cmap *CMap) parseBfrange() error {
	for {
		var _cfbd CharCode
		_beag, _bced := cmap.parseObject()
		if _bced != nil {
			if _bced == _c.EOF {
				break
			}
			return _bced
		}
		switch _cgc := _beag.(type) {
		case cmapOperand:
			if _cgc.Operand == _dgdd {
				return nil
			}
			return _be.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_cfbd = _ebba(_cgc)
		default:
			return _be.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		var _dec CharCode
		_beag, _bced = cmap.parseObject()
		if _bced != nil {
			if _bced == _c.EOF {
				break
			}
			return _bced
		}
		switch _ddg := _beag.(type) {
		case cmapOperand:
			_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0049\u006e\u0063\u006f\u006d\u0070\u006c\u0065\u0074\u0065\u0020\u0062\u0066r\u0061\u006e\u0067\u0065\u0020\u0074\u0072i\u0070\u006c\u0065\u0074")
			return ErrBadCMap
		case cmapHexString:
			_dec = _ebba(_ddg)
			if _dec > 0xffff {
				_dec = 0xffff
			}
		default:
			_bc.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _beag)
			return ErrBadCMap
		}
		_beag, _bced = cmap.parseObject()
		if _bced != nil {
			if _bced == _c.EOF {
				break
			}
			return _bced
		}
		switch _geef := _beag.(type) {
		case cmapArray:
			if len(_geef.Array) != int(_dec-_cfbd)+1 {
				_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0069\u0074\u0065\u006d\u0073\u0020\u0069\u006e\u0020a\u0072\u0072\u0061\u0079")
				return ErrBadCMap
			}
			for _bcg := _cfbd; _bcg <= _dec; _bcg++ {
				_baaa := _geef.Array[_bcg-_cfbd]
				_ebf, _abed := _baaa.(cmapHexString)
				if !_abed {
					return _be.New("\u006e\u006f\u006e-h\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				}
				_ggb := _ggee(_ebf)
				cmap._cbg[_bcg] = string(_ggb)
			}
		case cmapHexString:
			_dba := _ggee(_geef)
			_bbga := len(_dba)
			for _bcf := _cfbd; _bcf <= _dec; _bcf++ {
				cmap._cbg[_bcf] = string(_dba)
				if _bbga > 0 {
					_dba[_bbga-1]++
				} else {
					_bc.Log.Debug("\u004e\u006f\u0020c\u006d\u0061\u0070\u0020\u0074\u0061\u0072\u0067\u0065\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0066\u006f\u0072\u0020\u0025\u0023\u0076", _bcf)
				}
				if _bcf == 1<<32-1 {
					break
				}
			}
		default:
			_bc.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0055\u006e\u0065\u0078\u0070e\u0063t\u0065d\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _beag)
			return ErrBadCMap
		}
	}
	return nil
}
func (cmap *CMap) parseType() error {
	_ggfb := 0
	_eag := false
	for _cced := 0; _cced < 3 && !_eag; _cced++ {
		_eab, _fgg := cmap.parseObject()
		if _fgg != nil {
			return _fgg
		}
		switch _adb := _eab.(type) {
		case cmapOperand:
			switch _adb.Operand {
			case "\u0064\u0065\u0066":
				_eag = true
			default:
				_bc.Log.Error("\u0070\u0061r\u0073\u0065\u0054\u0079\u0070\u0065\u003a\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u006f=%\u0023\u0076", _eab)
				return ErrBadCMap
			}
		case cmapInt:
			_ggfb = int(_adb._dbae)
		}
	}
	cmap._beb = _ggfb
	return nil
}
func _fea(_eefc string) rune { _eeda := []rune(_eefc); return _eeda[len(_eeda)-1] }
func (cmap *CMap) parseVersion() error {
	_aacb := ""
	_cdd := false
	for _fdb := 0; _fdb < 3 && !_cdd; _fdb++ {
		_dcd, _cgg := cmap.parseObject()
		if _cgg != nil {
			return _cgg
		}
		switch _cbe := _dcd.(type) {
		case cmapOperand:
			switch _cbe.Operand {
			case "\u0064\u0065\u0066":
				_cdd = true
			default:
				_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0070\u0061\u0072\u0073\u0065\u0056e\u0072\u0073\u0069\u006f\u006e\u003a \u0073\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076", _dcd)
				return ErrBadCMap
			}
		case cmapInt:
			_aacb = _g.Sprintf("\u0025\u0064", _cbe._dbae)
		case cmapFloat:
			_aacb = _g.Sprintf("\u0025\u0066", _cbe._ege)
		case cmapString:
			_aacb = _cbe.String
		default:
			_bc.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020p\u0061\u0072\u0073\u0065Ver\u0073io\u006e\u003a\u0020\u0042\u0061\u0064\u0020ty\u0070\u0065\u002e\u0020\u006f\u003d\u0025#\u0076", _dcd)
		}
	}
	cmap._gg = _aacb
	return nil
}
func (_gb *CIDSystemInfo) String() string {
	return _g.Sprintf("\u0025\u0073\u002d\u0025\u0073\u002d\u0025\u0030\u0033\u0064", _gb.Registry, _gb.Ordering, _gb.Supplement)
}
func (cmap *CMap) parseCIDRange() error {
	for {
		_feef, _eac := cmap.parseObject()
		if _eac != nil {
			if _eac == _c.EOF {
				break
			}
			return _eac
		}
		_gbf, _af := _feef.(cmapHexString)
		if !_af {
			if _gfd, _ggae := _feef.(cmapOperand); _ggae {
				if _gfd.Operand == _fbdf {
					return nil
				}
				return _be.New("\u0063\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0072\u0076\u0061\u006c\u0020s\u0074\u0061\u0072\u0074\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
			}
		}
		_bbc := _ebba(_gbf)
		_feef, _eac = cmap.parseObject()
		if _eac != nil {
			if _eac == _c.EOF {
				break
			}
			return _eac
		}
		_gac, _af := _feef.(cmapHexString)
		if !_af {
			return _be.New("\u0063\u0069d\u0020\u0069\u006e\u0074e\u0072\u0076a\u006c\u0020\u0065\u006e\u0064\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0068\u0065\u0078\u0020\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_gbf._adac) != len(_gac._adac) {
			return _be.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_cddf := _ebba(_gac)
		if _bbc > _cddf {
			_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0049\u0044\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073t\u0061\u0072\u0074\u003d\u0030\u0078\u0025\u0030\u0032\u0078\u0020\u0065\u006e\u0064=\u0030x\u0025\u0030\u0032\u0078", _bbc, _cddf)
			return ErrBadCMap
		}
		_feef, _eac = cmap.parseObject()
		if _eac != nil {
			if _eac == _c.EOF {
				break
			}
			return _eac
		}
		_adf, _af := _feef.(cmapInt)
		if !_af {
			return _be.New("\u0063\u0069\u0064\u0020\u0073t\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		if _adf._dbae < 0 {
			return _be.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u0063\u0069\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0076\u0061\u006c\u0075\u0065")
		}
		_gca := _adf._dbae
		for _fag := _bbc; _fag <= _cddf; _fag++ {
			cmap._fa[_fag] = CharCode(_gca)
			_gca++
		}
		_bc.Log.Trace("C\u0049\u0044\u0020\u0072\u0061\u006eg\u0065\u003a\u0020\u003c\u0030\u0078\u0025\u0058\u003e \u003c\u0030\u0078%\u0058>\u0020\u0025\u0064", _bbc, _cddf, _adf._dbae)
	}
	return nil
}
func (_fbdc *cMapParser) parseHexString() (cmapHexString, error) {
	_fbdc._cfg.ReadByte()
	_fec := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	_gggd := _ec.Buffer{}
	for {
		_fbdc.skipSpaces()
		_eefb, _abc := _fbdc._cfg.Peek(1)
		if _abc != nil {
			return cmapHexString{}, _abc
		}
		if _eefb[0] == '>' {
			_fbdc._cfg.ReadByte()
			break
		}
		_dfcc, _ := _fbdc._cfg.ReadByte()
		if _ec.IndexByte(_fec, _dfcc) >= 0 {
			_gggd.WriteByte(_dfcc)
		}
	}
	if _gggd.Len()%2 == 1 {
		_bc.Log.Debug("\u0070\u0061rs\u0065\u0048\u0065x\u0053\u0074\u0072\u0069ng:\u0020ap\u0070\u0065\u006e\u0064\u0069\u006e\u0067 '\u0030\u0027\u0020\u0074\u006f\u0020\u0025#\u0071", _gggd.String())
		_gggd.WriteByte('0')
	}
	_fcad := _gggd.Len() / 2
	_adda, _ := _df.DecodeString(_gggd.String())
	return cmapHexString{_gbce: _fcad, _adac: _adda}, nil
}
func (_bcb *cMapParser) parseDict() (cmapDict, error) {
	_bc.Log.Trace("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020D\u0069\u0063\u0074\u0021")
	_ffa := _fggb()
	_cffe, _ := _bcb._cfg.ReadByte()
	if _cffe != '<' {
		return _ffa, ErrBadCMapDict
	}
	_cffe, _ = _bcb._cfg.ReadByte()
	if _cffe != '<' {
		return _ffa, ErrBadCMapDict
	}
	for {
		_bcb.skipSpaces()
		_acg, _caee := _bcb._cfg.Peek(2)
		if _caee != nil {
			return _ffa, _caee
		}
		if (_acg[0] == '>') && (_acg[1] == '>') {
			_bcb._cfg.ReadByte()
			_bcb._cfg.ReadByte()
			break
		}
		_eegf, _caee := _bcb.parseName()
		_bc.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _eegf.Name)
		if _caee != nil {
			_bc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u002e\u0020\u0065\u0072r=\u0025\u0076", _caee)
			return _ffa, _caee
		}
		_bcb.skipSpaces()
		_aggf, _caee := _bcb.parseObject()
		if _caee != nil {
			return _ffa, _caee
		}
		_ffa.Dict[_eegf.Name] = _aggf
		_bcb.skipSpaces()
		_acg, _caee = _bcb._cfg.Peek(3)
		if _caee != nil {
			return _ffa, _caee
		}
		if string(_acg) == "\u0064\u0065\u0066" {
			_bcb._cfg.Discard(3)
		}
	}
	return _ffa, nil
}
func (cmap *CMap) matchCode(_ebb []byte) (_cbgd CharCode, _gbc int, _dfcf bool) {
	for _aac := 0; _aac < _bb; _aac++ {
		if _aac < len(_ebb) {
			_cbgd = _cbgd<<8 | CharCode(_ebb[_aac])
			_gbc++
		}
		_dfcf = cmap.inCodespace(_cbgd, _aac+1)
		if _dfcf {
			return _cbgd, _gbc, true
		}
	}
	_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063o\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0020m\u0061t\u0063\u0068\u0065\u0073\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d=\u0025\u0023\u0071\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", _ebb, string(_ebb), cmap)
	return 0, 0, false
}
func _ggee(_gcb cmapHexString) []rune {
	if len(_gcb._adac) == 1 {
		return []rune{rune(_gcb._adac[0])}
	}
	_gdge := _gcb._adac
	if len(_gdge)%2 != 0 {
		_gdge = append(_gdge, 0)
		_bc.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u0073\u002e\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0073\u0068\u0065\u0078\u003d\u0025#\u0076\u0020\u0074\u006f\u0020\u0025\u002b\u0076", _gcb, _gdge)
	}
	_cdefc := len(_gdge) >> 1
	_gcg := make([]uint16, _cdefc)
	for _ecd := 0; _ecd < _cdefc; _ecd++ {
		_gcg[_ecd] = uint16(_gdge[_ecd<<1])<<8 + uint16(_gdge[_ecd<<1+1])
	}
	_ccedc := _e.Decode(_gcg)
	return _ccedc
}
func _fggb() cmapDict { return cmapDict{Dict: map[string]cmapObject{}} }

const (
	_ecbe  = "\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f"
	_cbc   = "\u0062e\u0067\u0069\u006e\u0063\u006d\u0061p"
	_gbe   = "\u0065n\u0064\u0063\u006d\u0061\u0070"
	_cff   = "\u0062\u0065\u0067\u0069nc\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0072\u0061\u006e\u0067\u0065"
	_add   = "\u0065\u006e\u0064\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065r\u0061\u006e\u0067\u0065"
	_egb   = "b\u0065\u0067\u0069\u006e\u0062\u0066\u0063\u0068\u0061\u0072"
	_agg   = "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r"
	_afe   = "\u0062\u0065\u0067i\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_dgdd  = "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065"
	_acc   = "\u0062\u0065\u0067\u0069\u006e\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_fbdf  = "e\u006e\u0064\u0063\u0069\u0064\u0072\u0061\u006e\u0067\u0065"
	_aaecb = "\u0075s\u0065\u0063\u006d\u0061\u0070"
	_acb   = "\u0057\u004d\u006fd\u0065"
	_gfe   = "\u0043\u004d\u0061\u0070\u004e\u0061\u006d\u0065"
	_aff   = "\u0043\u004d\u0061\u0070\u0054\u0079\u0070\u0065"
	_ddgg  = "C\u004d\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e"
)

func (cmap *CMap) parse() error {
	var _baag cmapObject
	for {
		_fdc, _cfc := cmap.parseObject()
		if _cfc != nil {
			if _cfc == _c.EOF {
				break
			}
			_bc.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u004d\u0061\u0070\u003a\u0020\u0025\u0076", _cfc)
			return _cfc
		}
		switch _ggff := _fdc.(type) {
		case cmapOperand:
			_aae := _ggff
			switch _aae.Operand {
			case _cff:
				_dda := cmap.parseCodespaceRange()
				if _dda != nil {
					return _dda
				}
			case _acc:
				_ecf := cmap.parseCIDRange()
				if _ecf != nil {
					return _ecf
				}
			case _egb:
				_agd := cmap.parseBfchar()
				if _agd != nil {
					return _agd
				}
			case _afe:
				_gddd := cmap.parseBfrange()
				if _gddd != nil {
					return _gddd
				}
			case _aaecb:
				if _baag == nil {
					_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u0073\u0065\u0063m\u0061\u0070\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f \u0061\u0072\u0067")
					return ErrBadCMap
				}
				_ddad, _eda := _baag.(cmapName)
				if !_eda {
					_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0075\u0073\u0065\u0063\u006d\u0061\u0070\u0020\u0061\u0072\u0067\u0020\u006eo\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u0020\u0025\u0023\u0076", _baag)
					return ErrBadCMap
				}
				cmap._db = _ddad.Name
			case _ecbe:
				_fee := cmap.parseSystemInfo()
				if _fee != nil {
					return _fee
				}
			}
		case cmapName:
			_cad := _ggff
			switch _cad.Name {
			case _ecbe:
				_aeb := cmap.parseSystemInfo()
				if _aeb != nil {
					return _aeb
				}
			case _gfe:
				_ddb := cmap.parseName()
				if _ddb != nil {
					return _ddb
				}
			case _aff:
				_dbbe := cmap.parseType()
				if _dbbe != nil {
					return _dbbe
				}
			case _ddgg:
				_egfc := cmap.parseVersion()
				if _egfc != nil {
					return _egfc
				}
			case _acb:
				if _cfc = cmap.parseWMode(); _cfc != nil {
					return _cfc
				}
			}
		}
		_baag = _fdc
	}
	return nil
}
func (_agga *cMapParser) parseName() (cmapName, error) {
	_ffd := ""
	_agc := false
	for {
		_gaa, _bdd := _agga._cfg.Peek(1)
		if _bdd == _c.EOF {
			break
		}
		if _bdd != nil {
			return cmapName{_ffd}, _bdd
		}
		if !_agc {
			if _gaa[0] == '/' {
				_agc = true
				_agga._cfg.ReadByte()
			} else {
				_bc.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u004e\u0061\u006d\u0065\u0020\u0073\u0074a\u0072t\u0069n\u0067 \u0077\u0069\u0074\u0068\u0020\u0025\u0073\u0020\u0028\u0025\u0020\u0078\u0029", _gaa, _gaa)
				return cmapName{_ffd}, _g.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _gaa[0])
			}
		} else {
			if _ba.IsWhiteSpace(_gaa[0]) {
				break
			} else if (_gaa[0] == '/') || (_gaa[0] == '[') || (_gaa[0] == '(') || (_gaa[0] == ']') || (_gaa[0] == '<') || (_gaa[0] == '>') {
				break
			} else if _gaa[0] == '#' {
				_afa, _fbc := _agga._cfg.Peek(3)
				if _fbc != nil {
					return cmapName{_ffd}, _fbc
				}
				_agga._cfg.Discard(3)
				_dga, _fbc := _df.DecodeString(string(_afa[1:3]))
				if _fbc != nil {
					return cmapName{_ffd}, _fbc
				}
				_ffd += string(_dga)
			} else {
				_efce, _ := _agga._cfg.ReadByte()
				_ffd += string(_efce)
			}
		}
	}
	return cmapName{_ffd}, nil
}

type fbRange struct {
	_ca CharCode
	_de CharCode
	_ge string
}

func (_fefg *cMapParser) parseString() (cmapString, error) {
	_fefg._cfg.ReadByte()
	_bbf := _ec.Buffer{}
	_cdfc := 1
	for {
		_eaef, _ddbgf := _fefg._cfg.Peek(1)
		if _ddbgf != nil {
			return cmapString{_bbf.String()}, _ddbgf
		}
		if _eaef[0] == '\\' {
			_fefg._cfg.ReadByte()
			_bbd, _eedf := _fefg._cfg.ReadByte()
			if _eedf != nil {
				return cmapString{_bbf.String()}, _eedf
			}
			if _ba.IsOctalDigit(_bbd) {
				_dbgg, _bca := _fefg._cfg.Peek(2)
				if _bca != nil {
					return cmapString{_bbf.String()}, _bca
				}
				var _fgd []byte
				_fgd = append(_fgd, _bbd)
				for _, _gfc := range _dbgg {
					if _ba.IsOctalDigit(_gfc) {
						_fgd = append(_fgd, _gfc)
					} else {
						break
					}
				}
				_fefg._cfg.Discard(len(_fgd) - 1)
				_bc.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _fgd)
				_cea, _bca := _b.ParseUint(string(_fgd), 8, 32)
				if _bca != nil {
					return cmapString{_bbf.String()}, _bca
				}
				_bbf.WriteByte(byte(_cea))
				continue
			}
			switch _bbd {
			case 'n':
				_bbf.WriteByte('\n')
			case 'r':
				_bbf.WriteByte('\r')
			case 't':
				_bbf.WriteByte('\t')
			case 'b':
				_bbf.WriteByte('\b')
			case 'f':
				_bbf.WriteByte('\f')
			case '(':
				_bbf.WriteByte('(')
			case ')':
				_bbf.WriteByte(')')
			case '\\':
				_bbf.WriteByte('\\')
			}
			continue
		} else if _eaef[0] == '(' {
			_cdfc++
		} else if _eaef[0] == ')' {
			_cdfc--
			if _cdfc == 0 {
				_fefg._cfg.ReadByte()
				break
			}
		}
		_gge, _ := _fefg._cfg.ReadByte()
		_bbf.WriteByte(_gge)
	}
	return cmapString{_bbf.String()}, nil
}
func (cmap *CMap) NBits() int { return cmap._bf }

var (
	ErrBadCMap        = _be.New("\u0062\u0061\u0064\u0020\u0063\u006d\u0061\u0070")
	ErrBadCMapComment = _be.New("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025")
	ErrBadCMapDict    = _be.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
)

func (cmap *CMap) CharcodeToUnicode(code CharCode) (string, bool) {
	if _ecb, _bfgb := cmap._cbg[code]; _bfgb {
		return _ecb, true
	}
	return MissingCodeString, false
}
func (_bgd *cMapParser) parseOperand() (cmapOperand, error) {
	_ebcc := cmapOperand{}
	_dada := _ec.Buffer{}
	for {
		_ada, _ccea := _bgd._cfg.Peek(1)
		if _ccea != nil {
			if _ccea == _c.EOF {
				break
			}
			return _ebcc, _ccea
		}
		if _ba.IsDelimiter(_ada[0]) {
			break
		}
		if _ba.IsWhiteSpace(_ada[0]) {
			break
		}
		_gae, _ := _bgd._cfg.ReadByte()
		_dada.WriteByte(_gae)
	}
	if _dada.Len() == 0 {
		return _ebcc, _g.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	_ebcc.Operand = _dada.String()
	return _ebcc, nil
}
func _bec(_deda bool) *CMap {
	_dgf := 16
	if _deda {
		_dgf = 8
	}
	return &CMap{_bf: _dgf, _fa: make(map[CharCode]CharCode), _bd: make(map[CharCode]CharCode), _cbg: make(map[CharCode]string), _gcf: make(map[string]CharCode)}
}

type cmapDict struct{ Dict map[string]cmapObject }

func (_beab *cMapParser) parseObject() (cmapObject, error) {
	_beab.skipSpaces()
	for {
		_age, _gcd := _beab._cfg.Peek(2)
		if _gcd != nil {
			return nil, _gcd
		}
		if _age[0] == '%' {
			_beab.parseComment()
			_beab.skipSpaces()
			continue
		} else if _age[0] == '/' {
			_badc, _ggc := _beab.parseName()
			return _badc, _ggc
		} else if _age[0] == '(' {
			_dad, _gddf := _beab.parseString()
			return _dad, _gddf
		} else if _age[0] == '[' {
			_bagd, _fce := _beab.parseArray()
			return _bagd, _fce
		} else if (_age[0] == '<') && (_age[1] == '<') {
			_gdc, _cdef := _beab.parseDict()
			return _gdc, _cdef
		} else if _age[0] == '<' {
			_aga, _fceb := _beab.parseHexString()
			return _aga, _fceb
		} else if _ba.IsDecimalDigit(_age[0]) || (_age[0] == '-' && _ba.IsDecimalDigit(_age[1])) {
			_dbba, _gegf := _beab.parseNumber()
			if _gegf != nil {
				return nil, _gegf
			}
			return _dbba, nil
		} else {
			_edeg, _feae := _beab.parseOperand()
			if _feae != nil {
				return nil, _feae
			}
			return _edeg, nil
		}
	}
}

type cmapFloat struct{ _ege float64 }

func (cmap *CMap) CIDSystemInfo() CIDSystemInfo { return cmap._dd }
func (cmap *CMap) Stream() (*_ba.PdfObjectStream, error) {
	if cmap._fg != nil {
		return cmap._fg, nil
	}
	_dac, _dff := _ba.MakeStream(cmap.Bytes(), _ba.NewFlateEncoder())
	if _dff != nil {
		return nil, _dff
	}
	cmap._fg = _dac
	return cmap._fg, nil
}

type Codespace struct {
	NumBytes int
	Low      CharCode
	High     CharCode
}
type charRange struct {
	_f  CharCode
	_ef CharCode
}

func _feg(_fbf []byte) *cMapParser {
	_ggffa := cMapParser{}
	_ffb := _ec.NewBuffer(_fbf)
	_ggffa._cfg = _a.NewReader(_ffb)
	return &_ggffa
}
func (cmap *CMap) parseCodespaceRange() error {
	for {
		_eae, _bgg := cmap.parseObject()
		if _bgg != nil {
			if _bgg == _c.EOF {
				break
			}
			return _bgg
		}
		_cef, _ccb := _eae.(cmapHexString)
		if !_ccb {
			if _bgc, _egfd := _eae.(cmapOperand); _egfd {
				if _bgc.Operand == _add {
					return nil
				}
				return _be.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
			}
		}
		_eae, _bgg = cmap.parseObject()
		if _bgg != nil {
			if _bgg == _c.EOF {
				break
			}
			return _bgg
		}
		_dea, _ccb := _eae.(cmapHexString)
		if !_ccb {
			return _be.New("\u006e\u006f\u006e-\u0068\u0065\u0078\u0020\u0068\u0069\u0067\u0068")
		}
		if len(_cef._adac) != len(_dea._adac) {
			return _be.New("\u0075\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0062\u0079\u0074\u0065\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065")
		}
		_acf := _ebba(_cef)
		_fbd := _ebba(_dea)
		if _fbd < _acf {
			_bc.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0061d\u0020\u0063\u006fd\u0065\u0073\u0070\u0061\u0063\u0065\u002e\u0020\u006cow\u003d\u0030\u0078%\u0030\u0032x\u0020\u0068\u0069\u0067\u0068\u003d0\u0078\u00250\u0032\u0078", _acf, _fbd)
			return ErrBadCMap
		}
		_eaa := _dea._gbce
		_geb := Codespace{NumBytes: _eaa, Low: _acf, High: _fbd}
		cmap._dbb = append(cmap._dbb, _geb)
		_bc.Log.Trace("\u0043\u006f\u0064e\u0073\u0070\u0061\u0063e\u0020\u006c\u006f\u0077\u003a\u0020\u0030x\u0025\u0058\u002c\u0020\u0068\u0069\u0067\u0068\u003a\u0020\u0030\u0078\u0025\u0058", _acf, _fbd)
	}
	if len(cmap._dbb) == 0 {
		_bc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u0020\u0069\u006e\u0020\u0063ma\u0070\u002e")
		return ErrBadCMap
	}
	return nil
}
func (cmap *CMap) parseSystemInfo() error {
	_bab := false
	_aaec := false
	_abea := ""
	_deg := false
	_cgb := CIDSystemInfo{}
	for _fca := 0; _fca < 50 && !_deg; _fca++ {
		_cde, _edb := cmap.parseObject()
		if _edb != nil {
			return _edb
		}
		switch _ade := _cde.(type) {
		case cmapDict:
			_bag := _ade.Dict
			_fef, _facc := _bag["\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"]
			if !_facc {
				_bc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ede, _facc := _fef.(cmapString)
			if !_facc {
				_bc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_cgb.Registry = _ede.String
			_fef, _facc = _bag["\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"]
			if !_facc {
				_bc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_ede, _facc = _fef.(cmapString)
			if !_facc {
				_bc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_cgb.Ordering = _ede.String
			_dcde, _facc := _bag["\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"]
			if !_facc {
				_bc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_babg, _facc := _dcde.(cmapInt)
			if !_facc {
				_bc.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0042\u0061\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006ef\u006f")
				return ErrBadCMap
			}
			_cgb.Supplement = int(_babg._dbae)
			_deg = true
		case cmapOperand:
			switch _ade.Operand {
			case "\u0062\u0065\u0067i\u006e":
				_bab = true
			case "\u0065\u006e\u0064":
				_deg = true
			case "\u0064\u0065\u0066":
				_aaec = false
			}
		case cmapName:
			if _bab {
				_abea = _ade.Name
				_aaec = true
			}
		case cmapString:
			if _aaec {
				switch _abea {
				case "\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079":
					_cgb.Registry = _ade.String
				case "\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067":
					_cgb.Ordering = _ade.String
				}
			}
		case cmapInt:
			if _aaec {
				switch _abea {
				case "\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074":
					_cgb.Supplement = int(_ade._dbae)
				}
			}
		}
	}
	if !_deg {
		_bc.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u0020\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0020\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u006c\u0079")
		return ErrBadCMap
	}
	cmap._dd = _cgb
	return nil
}
func (_dbe *cMapParser) parseNumber() (cmapObject, error) {
	_fbg, _cda := _ba.ParseNumber(_dbe._cfg)
	if _cda != nil {
		return nil, _cda
	}
	switch _eaba := _fbg.(type) {
	case *_ba.PdfObjectFloat:
		return cmapFloat{float64(*_eaba)}, nil
	case *_ba.PdfObjectInteger:
		return cmapInt{int64(*_eaba)}, nil
	}
	return nil, _g.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _fbg)
}
func (cmap *CMap) CharcodeBytesToUnicode(data []byte) (string, int) {
	_cce, _fgf := cmap.BytesToCharcodes(data)
	if !_fgf {
		_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u0068\u0061\u0072\u0063\u006f\u0064\u0065\u0042\u0079\u0074\u0065s\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069n\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0064\u0061\u0074\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078]\u0020\u0063\u006d\u0061\u0070=\u0025\u0073", data, cmap)
		return "", 0
	}
	_ega := make([]string, len(_cce))
	var _cd []CharCode
	for _gde, _bg := range _cce {
		_gbgc, _cdf := cmap._cbg[_bg]
		if !_cdf {
			_cd = append(_cd, _bg)
			_gbgc = MissingCodeString
		}
		_ega[_gde] = _gbgc
	}
	_baf := _ag.Join(_ega, "")
	if len(_cd) > 0 {
		_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0068\u0061\u0072c\u006f\u0064\u0065\u0042y\u0074\u0065\u0073\u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u002e\u0020\u004e\u006f\u0074\u0020\u0069\u006e\u0020\u006d\u0061\u0070\u002e\u000a"+"\u0009d\u0061t\u0061\u003d\u005b\u0025\u00200\u0032\u0078]\u003d\u0025\u0023\u0071\u000a"+"\u0009\u0063h\u0061\u0072\u0063o\u0064\u0065\u0073\u003d\u0025\u0030\u0032\u0078\u000a"+"\u0009\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u003d\u0025\u0064\u0020%\u0030\u0032\u0078\u000a"+"\u0009\u0075\u006e\u0069\u0063\u006f\u0064\u0065\u003d`\u0025\u0073\u0060\u000a"+"\u0009\u0063\u006d\u0061\u0070\u003d\u0025\u0073", data, string(data), _cce, len(_cd), _cd, _baf, cmap)
	}
	return _baf, len(_cd)
}
func (cmap *CMap) parseBfchar() error {
	for {
		_eacb, _gfbg := cmap.parseObject()
		if _gfbg != nil {
			if _gfbg == _c.EOF {
				break
			}
			return _gfbg
		}
		var _ddbb CharCode
		switch _dca := _eacb.(type) {
		case cmapOperand:
			if _dca.Operand == _agg {
				return nil
			}
			return _be.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065d\u0020\u006fp\u0065\u0072\u0061\u006e\u0064")
		case cmapHexString:
			_ddbb = _ebba(_dca)
		default:
			return _be.New("\u0075n\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065")
		}
		_eacb, _gfbg = cmap.parseObject()
		if _gfbg != nil {
			if _gfbg == _c.EOF {
				break
			}
			return _gfbg
		}
		var _cbd []rune
		switch _eabc := _eacb.(type) {
		case cmapOperand:
			if _eabc.Operand == _agg {
				return nil
			}
			_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020o\u0070\u0065\u0072\u0061\u006e\u0064\u002e\u0020\u0025\u0023\u0076", _eabc)
			return ErrBadCMap
		case cmapHexString:
			_cbd = _ggee(_eabc)
		case cmapName:
			_bc.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0061\u006de\u002e \u0025\u0023\u0076", _eabc)
			_cbd = []rune{MissingCodeRune}
		default:
			_bc.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064 \u0074\u0079\u0070e\u002e \u0025\u0023\u0076", _eacb)
			return ErrBadCMap
		}
		cmap._cbg[_ddbb] = string(_cbd)
	}
	return nil
}

type CharCode uint32

func _ecca(_bdgf cmapHexString) rune {
	_edaf := _ggee(_bdgf)
	if _gace := len(_edaf); _gace == 0 {
		_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054o\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065c\u0074\u0065\u0064\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0072u\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023\u0076", _bdgf)
		return MissingCodeRune
	}
	if len(_edaf) > 1 {
		_bc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0068\u0065\u0078\u0054\u006f\u0052\u0075\u006e\u0065\u002e\u0020\u0045\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0078\u0061\u0063\u0074\u006c\u0079\u0020\u006f\u006e\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u0073\u0068\u0065\u0078\u003d\u0025\u0023v\u0020\u002d\u003e\u0020\u0025#\u0076", _bdgf, _edaf)
	}
	return _edaf[0]
}
func _caea(_bedc, _ccbg int) int {
	if _bedc < _ccbg {
		return _bedc
	}
	return _ccbg
}

type cmapObject interface{}
type CMap struct {
	*cMapParser
	_ded string
	_bf  int
	_beb int
	_gg  string
	_db  string
	_dd  CIDSystemInfo
	_dbb []Codespace
	_fa  map[CharCode]CharCode
	_bd  map[CharCode]CharCode
	_cbg map[CharCode]string
	_gcf map[string]CharCode
	_gd  []byte
	_fg  *_ba.PdfObjectStream
	_eb  integer
}
type cmapArray struct{ Array []cmapObject }

func (cmap *CMap) parseWMode() error {
	var _bae int
	_bbg := false
	for _dce := 0; _dce < 3 && !_bbg; _dce++ {
		_edcc, _ddbg := cmap.parseObject()
		if _ddbg != nil {
			return _ddbg
		}
		switch _fac := _edcc.(type) {
		case cmapOperand:
			switch _fac.Operand {
			case "\u0064\u0065\u0066":
				_bbg = true
			default:
				_bc.Log.Error("\u0070\u0061\u0072\u0073\u0065\u0057\u004d\u006f\u0064\u0065:\u0020\u0073\u0074\u0061\u0074\u0065\u0020e\u0072\u0072\u006f\u0072\u002e\u0020\u006f\u003d\u0025\u0023\u0076", _edcc)
				return ErrBadCMap
			}
		case cmapInt:
			_bae = int(_fac._dbae)
		}
	}
	cmap._eb = integer{_efaa: true, _beee: _bae}
	return nil
}

type cmapName struct{ Name string }
type cmapOperand struct{ Operand string }

func (cmap *CMap) toBfData() string {
	if len(cmap._cbg) == 0 {
		return ""
	}
	_edc := make([]CharCode, 0, len(cmap._cbg))
	for _ea := range cmap._cbg {
		_edc = append(_edc, _ea)
	}
	_ee.Slice(_edc, func(_cfa, _bee int) bool { return _edc[_cfa] < _edc[_bee] })
	var _abd []charRange
	_bcc := charRange{_edc[0], _edc[0]}
	_ebc := cmap._cbg[_edc[0]]
	for _, _ceeg := range _edc[1:] {
		_cbbb := cmap._cbg[_ceeg]
		if _ceeg == _bcc._ef+1 && _fea(_cbbb) == _fea(_ebc)+1 {
			_bcc._ef = _ceeg
		} else {
			_abd = append(_abd, _bcc)
			_bcc._f, _bcc._ef = _ceeg, _ceeg
		}
		_ebc = _cbbb
	}
	_abd = append(_abd, _bcc)
	var _bdaa []CharCode
	var _cdb []fbRange
	for _, _dgd := range _abd {
		if _dgd._f == _dgd._ef {
			_bdaa = append(_bdaa, _dgd._f)
		} else {
			_cdb = append(_cdb, fbRange{_ca: _dgd._f, _de: _dgd._ef, _ge: cmap._cbg[_dgd._f]})
		}
	}
	_bc.Log.Trace("\u0063\u0068ar\u0052\u0061\u006eg\u0065\u0073\u003d\u0025d f\u0062Ch\u0061\u0072\u0073\u003d\u0025\u0064\u0020fb\u0052\u0061\u006e\u0067\u0065\u0073\u003d%\u0064", len(_abd), len(_bdaa), len(_cdb))
	var _eeb []string
	if len(_bdaa) > 0 {
		_ggg := (len(_bdaa) + _efa - 1) / _efa
		for _ecc := 0; _ecc < _ggg; _ecc++ {
			_abdb := _caea(len(_bdaa)-_ecc*_efa, _efa)
			_eeb = append(_eeb, _g.Sprintf("\u0025\u0064\u0020\u0062\u0065\u0067\u0069\u006e\u0062f\u0063\u0068\u0061\u0072", _abdb))
			for _gab := 0; _gab < _abdb; _gab++ {
				_fdf := _bdaa[_ecc*_efa+_gab]
				_caa := cmap._cbg[_fdf]
				_eeb = append(_eeb, _g.Sprintf("\u003c%\u0030\u0034\u0078\u003e\u0020\u0025s", _fdf, _ebe(_caa)))
			}
			_eeb = append(_eeb, "\u0065n\u0064\u0062\u0066\u0063\u0068\u0061r")
		}
	}
	if len(_cdb) > 0 {
		_fgfd := (len(_cdb) + _efa - 1) / _efa
		for _gbd := 0; _gbd < _fgfd; _gbd++ {
			_bcdc := _caea(len(_cdb)-_gbd*_efa, _efa)
			_eeb = append(_eeb, _g.Sprintf("\u0025d\u0020b\u0065\u0067\u0069\u006e\u0062\u0066\u0072\u0061\u006e\u0067\u0065", _bcdc))
			for _cab := 0; _cab < _bcdc; _cab++ {
				_aaa := _cdb[_gbd*_efa+_cab]
				_eeb = append(_eeb, _g.Sprintf("\u003c%\u00304\u0078\u003e\u003c\u0025\u0030\u0034\u0078\u003e\u0020\u0025\u0073", _aaa._ca, _aaa._de, _ebe(_aaa._ge)))
			}
			_eeb = append(_eeb, "\u0065\u006e\u0064\u0062\u0066\u0072\u0061\u006e\u0067\u0065")
		}
	}
	return _ag.Join(_eeb, "\u000a")
}
func LoadPredefinedCMap(name string) (*CMap, error) {
	cmap, _bda := _gga(name)
	if _bda != nil {
		return nil, _bda
	}
	if cmap._db == "" {
		cmap.computeInverseMappings()
		return cmap, nil
	}
	_bfg, _bda := _gga(cmap._db)
	if _bda != nil {
		return nil, _bda
	}
	for _cbb, _ceb := range _bfg._fa {
		if _, _eec := cmap._fa[_cbb]; !_eec {
			cmap._fa[_cbb] = _ceb
		}
	}
	cmap._dbb = append(cmap._dbb, _bfg._dbb...)
	cmap.computeInverseMappings()
	return cmap, nil
}
func LoadCmapFromData(data []byte, isSimple bool) (*CMap, error) {
	_bc.Log.Trace("\u004c\u006fa\u0064\u0043\u006d\u0061\u0070\u0046\u0072\u006f\u006d\u0044\u0061\u0074\u0061\u003a\u0020\u0069\u0073\u0053\u0069\u006d\u0070\u006ce=\u0025\u0074", isSimple)
	cmap := _bec(isSimple)
	cmap.cMapParser = _feg(data)
	_agf := cmap.parse()
	if _agf != nil {
		return nil, _agf
	}
	if len(cmap._dbb) == 0 {
		if cmap._db != "" {
			return cmap, nil
		}
		_bc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0073\u0070\u0061\u0063\u0065\u0073\u002e\u0020\u0063\u006d\u0061p=\u0025\u0073", cmap)
	}
	cmap.computeInverseMappings()
	return cmap, nil
}
func (_fege *cMapParser) skipSpaces() (int, error) {
	_acfg := 0
	for {
		_gfg, _bcedb := _fege._cfg.Peek(1)
		if _bcedb != nil {
			return 0, _bcedb
		}
		if _ba.IsWhiteSpace(_gfg[0]) {
			_fege._cfg.ReadByte()
			_acfg++
		} else {
			break
		}
	}
	return _acfg, nil
}
func (cmap *CMap) Name() string { return cmap._ded }
func (_dcg *cMapParser) parseComment() (string, error) {
	var _cabe _ec.Buffer
	_, _ccdb := _dcg.skipSpaces()
	if _ccdb != nil {
		return _cabe.String(), _ccdb
	}
	_gacc := true
	for {
		_cgfe, _addb := _dcg._cfg.Peek(1)
		if _addb != nil {
			_bc.Log.Debug("p\u0061r\u0073\u0065\u0043\u006f\u006d\u006d\u0065\u006et\u003a\u0020\u0065\u0072r=\u0025\u0076", _addb)
			return _cabe.String(), _addb
		}
		if _gacc && _cgfe[0] != '%' {
			return _cabe.String(), ErrBadCMapComment
		}
		_gacc = false
		if (_cgfe[0] != '\r') && (_cgfe[0] != '\n') {
			_fdcb, _ := _dcg._cfg.ReadByte()
			_cabe.WriteByte(_fdcb)
		} else {
			break
		}
	}
	return _cabe.String(), nil
}

type cmapHexString struct {
	_gbce int
	_adac []byte
}

func (cmap *CMap) BytesToCharcodes(data []byte) ([]CharCode, bool) {
	var _dgc []CharCode
	if cmap._bf == 8 {
		for _, _eeg := range data {
			_dgc = append(_dgc, CharCode(_eeg))
		}
		return _dgc, true
	}
	for _bdc := 0; _bdc < len(data); {
		_dab, _bfe, _ggf := cmap.matchCode(data[_bdc:])
		if !_ggf {
			_bc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0063\u006f\u0064\u0065\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0074\u0020\u0069\u003d\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d\u003d\u0025\u0023\u0071", _bdc, data, string(data))
			return _dgc, false
		}
		_dgc = append(_dgc, _dab)
		_bdc += _bfe
	}
	return _dgc, true
}
func (cmap *CMap) inCodespace(_dfa CharCode, _bad int) bool {
	for _, _ed := range cmap._dbb {
		if _ed.Low <= _dfa && _dfa <= _ed.High && _bad == _ed.NumBytes {
			return true
		}
	}
	return false
}
func (cmap *CMap) computeInverseMappings() {
	for _eca, _fe := range cmap._fa {
		if _cc, _eg := cmap._bd[_fe]; !_eg || (_eg && _cc > _eca) {
			cmap._bd[_fe] = _eca
		}
	}
	for _gbg, _fcb := range cmap._cbg {
		if _gda, _gdd := cmap._gcf[_fcb]; !_gdd || (_gdd && _gda > _gbg) {
			cmap._gcf[_fcb] = _gbg
		}
	}
	_ee.Slice(cmap._dbb, func(_bea, _ab int) bool { return cmap._dbb[_bea].Low < cmap._dbb[_ab].Low })
}
func (cmap *CMap) StringToCID(s string) (CharCode, bool) {
	_dbg, _baa := cmap._gcf[s]
	return _dbg, _baa
}

type integer struct {
	_efaa bool
	_beee int
}
type cmapString struct{ String string }

func IsPredefinedCMap(name string) bool { return _ga.AssetExists(name) }
func (cmap *CMap) CIDToCharcode(cid CharCode) (CharCode, bool) {
	_gec, _aa := cmap._bd[cid]
	return _gec, _aa
}
func (cmap *CMap) Type() int { return cmap._beb }
func (cmap *CMap) parseName() error {
	_cae := ""
	_dgb := false
	for _ccee := 0; _ccee < 20 && !_dgb; _ccee++ {
		_fge, _ad := cmap.parseObject()
		if _ad != nil {
			return _ad
		}
		switch _fcg := _fge.(type) {
		case cmapOperand:
			switch _fcg.Operand {
			case "\u0064\u0065\u0066":
				_dgb = true
			default:
				_bc.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a\u0020\u0053\u0074\u0061\u0074\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020o\u003d\u0025\u0023\u0076\u0020n\u0061\u006de\u003d\u0025\u0023\u0071", _fge, _cae)
				if _cae != "" {
					_cae = _g.Sprintf("\u0025\u0073\u0020%\u0073", _cae, _fcg.Operand)
				}
				_bc.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u004e\u0061\u006d\u0065\u003a \u0052\u0065\u0063\u006f\u0076\u0065\u0072e\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025\u0023\u0071", _cae)
			}
		case cmapName:
			_cae = _fcg.Name
		}
	}
	if !_dgb {
		_bc.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0070\u0061\u0072\u0073\u0065N\u0061m\u0065:\u0020\u004e\u006f\u0020\u0064\u0065\u0066 ")
		return ErrBadCMap
	}
	cmap._ded = _cae
	return nil
}
func (cmap *CMap) CharcodeToCID(code CharCode) (CharCode, bool) {
	_bebg, _fcf := cmap._fa[code]
	return _bebg, _fcf
}
func (_edbd *cMapParser) parseArray() (cmapArray, error) {
	_gfcd := cmapArray{}
	_gfcd.Array = []cmapObject{}
	_edbd._cfg.ReadByte()
	for {
		_edbd.skipSpaces()
		_feed, _dfcd := _edbd._cfg.Peek(1)
		if _dfcd != nil {
			return _gfcd, _dfcd
		}
		if _feed[0] == ']' {
			_edbd._cfg.ReadByte()
			break
		}
		_gdg, _dfcd := _edbd.parseObject()
		if _dfcd != nil {
			return _gfcd, _dfcd
		}
		_gfcd.Array = append(_gfcd.Array, _gdg)
	}
	return _gfcd, nil
}
func NewCIDSystemInfo(obj _ba.PdfObject) (_cb CIDSystemInfo, _ce error) {
	_ac, _gc := _ba.GetDict(obj)
	if !_gc {
		return CIDSystemInfo{}, _ba.ErrTypeError
	}
	_ff, _gc := _ba.GetStringVal(_ac.Get("\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079"))
	if !_gc {
		return CIDSystemInfo{}, _ba.ErrTypeError
	}
	_fb, _gc := _ba.GetStringVal(_ac.Get("\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067"))
	if !_gc {
		return CIDSystemInfo{}, _ba.ErrTypeError
	}
	_cee, _gc := _ba.GetIntVal(_ac.Get("\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074"))
	if !_gc {
		return CIDSystemInfo{}, _ba.ErrTypeError
	}
	return CIDSystemInfo{Registry: _ff, Ordering: _fb, Supplement: _cee}, nil
}
func _ebba(_ced cmapHexString) CharCode {
	_cceb := CharCode(0)
	for _, _cefa := range _ced._adac {
		_cceb <<= 8
		_cceb |= CharCode(_cefa)
	}
	return _cceb
}
func (cmap *CMap) Bytes() []byte {
	_bc.Log.Trace("\u0063\u006d\u0061\u0070.B\u0079\u0074\u0065\u0073\u003a\u0020\u0063\u006d\u0061\u0070\u003d\u0025\u0073", cmap.String())
	if len(cmap._gd) > 0 {
		return cmap._gd
	}
	cmap._gd = []byte(_ag.Join([]string{_gfb, cmap.toBfData(), _dag}, "\u000a"))
	return cmap._gd
}
