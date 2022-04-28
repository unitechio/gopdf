package fonts

import (
	_fe "bytes"
	_eaf "encoding/binary"
	_c "errors"
	_f "fmt"
	_d "io"
	_g "os"
	_be "regexp"
	_ea "sort"
	_b "strings"
	_fa "sync"

	_gb "bitbucket.org/shenghui0779/gopdf/common"
	_bc "bitbucket.org/shenghui0779/gopdf/core"
	_ff "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_ca "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_bb "golang.org/x/xerrors"
)

var _cfc *RuneCharSafeMap
var _cd = &fontMap{_ecf: make(map[StdFontName]func() StdFont)}

type ttfParser struct {
	_afe  TtfType
	_dada _d.ReadSeeker
	_fgb  map[string]uint32
	_edbd uint16
	_cbc  uint16
}

func (_cgda *ttfParser) readByte() (_gaca uint8) {
	_eaf.Read(_cgda._dada, _eaf.BigEndian, &_gaca)
	return _gaca
}
func (_ffce *ttfParser) Read32Fixed() float64 {
	_bfc := float64(_ffce.ReadShort())
	_fbeg := float64(_ffce.ReadUShort()) / 65536.0
	return _bfc + _fbeg
}

type CharMetrics struct {
	Wx float64
	Wy float64
}

func (_ffca *ttfParser) ReadUShort() (_fgga uint16) {
	_eaf.Read(_ffca._dada, _eaf.BigEndian, &_fgga)
	return _fgga
}
func (_cb StdFont) Descriptor() Descriptor { return _cb._eab }

var _cab _fa.Once

func (_gbc StdFont) GetMetricsTable() *RuneCharSafeMap { return _gbc._eee }
func (_ddec *ttfParser) parseCmapFormat0() error {
	_efd, _fed := _ddec.ReadStr(256)
	if _fed != nil {
		return _fed
	}
	_cea := []byte(_efd)
	_gb.Log.Trace("\u0070a\u0072\u0073e\u0043\u006d\u0061p\u0046\u006f\u0072\u006d\u0061\u0074\u0030:\u0020\u0025\u0073\u000a\u0064\u0061t\u0061\u0053\u0074\u0072\u003d\u0025\u002b\u0071\u000a\u0064\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d", _ddec._afe.String(), _efd, _cea)
	for _aef, _bdd := range _cea {
		_ddec._afe.Chars[rune(_aef)] = GID(_bdd)
	}
	return nil
}
func _cbd() StdFont {
	_cab.Do(_fd)
	_gge := Descriptor{Name: CourierObliqueName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0061, BBox: [4]float64{-27, -250, 849, 805}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_gge, _ccd)
}

var _cbag *RuneCharSafeMap
var _cff = &RuneCharSafeMap{_dd: map[rune]CharMetrics{' ': {Wx: 250}, '!': {Wx: 333}, '#': {Wx: 500}, '%': {Wx: 833}, '&': {Wx: 778}, '(': {Wx: 333}, ')': {Wx: 333}, '+': {Wx: 549}, ',': {Wx: 250}, '.': {Wx: 250}, '/': {Wx: 278}, '0': {Wx: 500}, '1': {Wx: 500}, '2': {Wx: 500}, '3': {Wx: 500}, '4': {Wx: 500}, '5': {Wx: 500}, '6': {Wx: 500}, '7': {Wx: 500}, '8': {Wx: 500}, '9': {Wx: 500}, ':': {Wx: 278}, ';': {Wx: 278}, '<': {Wx: 549}, '=': {Wx: 549}, '>': {Wx: 549}, '?': {Wx: 444}, '[': {Wx: 333}, ']': {Wx: 333}, '_': {Wx: 500}, '{': {Wx: 480}, '|': {Wx: 200}, '}': {Wx: 480}, '¬': {Wx: 713}, '°': {Wx: 400}, '±': {Wx: 549}, 'µ': {Wx: 576}, '×': {Wx: 549}, '÷': {Wx: 549}, 'ƒ': {Wx: 500}, 'Α': {Wx: 722}, 'Β': {Wx: 667}, 'Γ': {Wx: 603}, 'Ε': {Wx: 611}, 'Ζ': {Wx: 611}, 'Η': {Wx: 722}, 'Θ': {Wx: 741}, 'Ι': {Wx: 333}, 'Κ': {Wx: 722}, 'Λ': {Wx: 686}, 'Μ': {Wx: 889}, 'Ν': {Wx: 722}, 'Ξ': {Wx: 645}, 'Ο': {Wx: 722}, 'Π': {Wx: 768}, 'Ρ': {Wx: 556}, 'Σ': {Wx: 592}, 'Τ': {Wx: 611}, 'Υ': {Wx: 690}, 'Φ': {Wx: 763}, 'Χ': {Wx: 722}, 'Ψ': {Wx: 795}, 'α': {Wx: 631}, 'β': {Wx: 549}, 'γ': {Wx: 411}, 'δ': {Wx: 494}, 'ε': {Wx: 439}, 'ζ': {Wx: 494}, 'η': {Wx: 603}, 'θ': {Wx: 521}, 'ι': {Wx: 329}, 'κ': {Wx: 549}, 'λ': {Wx: 549}, 'ν': {Wx: 521}, 'ξ': {Wx: 493}, 'ο': {Wx: 549}, 'π': {Wx: 549}, 'ρ': {Wx: 549}, 'ς': {Wx: 439}, 'σ': {Wx: 603}, 'τ': {Wx: 439}, 'υ': {Wx: 576}, 'φ': {Wx: 521}, 'χ': {Wx: 549}, 'ψ': {Wx: 686}, 'ω': {Wx: 686}, 'ϑ': {Wx: 631}, 'ϒ': {Wx: 620}, 'ϕ': {Wx: 603}, 'ϖ': {Wx: 713}, '•': {Wx: 460}, '…': {Wx: 1000}, '′': {Wx: 247}, '″': {Wx: 411}, '⁄': {Wx: 167}, '€': {Wx: 750}, 'ℑ': {Wx: 686}, '℘': {Wx: 987}, 'ℜ': {Wx: 795}, 'Ω': {Wx: 768}, 'ℵ': {Wx: 823}, '←': {Wx: 987}, '↑': {Wx: 603}, '→': {Wx: 987}, '↓': {Wx: 603}, '↔': {Wx: 1042}, '↵': {Wx: 658}, '⇐': {Wx: 987}, '⇑': {Wx: 603}, '⇒': {Wx: 987}, '⇓': {Wx: 603}, '⇔': {Wx: 1042}, '∀': {Wx: 713}, '∂': {Wx: 494}, '∃': {Wx: 549}, '∅': {Wx: 823}, '∆': {Wx: 612}, '∇': {Wx: 713}, '∈': {Wx: 713}, '∉': {Wx: 713}, '∋': {Wx: 439}, '∏': {Wx: 823}, '∑': {Wx: 713}, '−': {Wx: 549}, '∗': {Wx: 500}, '√': {Wx: 549}, '∝': {Wx: 713}, '∞': {Wx: 713}, '∠': {Wx: 768}, '∧': {Wx: 603}, '∨': {Wx: 603}, '∩': {Wx: 768}, '∪': {Wx: 768}, '∫': {Wx: 274}, '∴': {Wx: 863}, '∼': {Wx: 549}, '≅': {Wx: 549}, '≈': {Wx: 549}, '≠': {Wx: 549}, '≡': {Wx: 549}, '≤': {Wx: 549}, '≥': {Wx: 549}, '⊂': {Wx: 713}, '⊃': {Wx: 713}, '⊄': {Wx: 713}, '⊆': {Wx: 713}, '⊇': {Wx: 713}, '⊕': {Wx: 768}, '⊗': {Wx: 768}, '⊥': {Wx: 658}, '⋅': {Wx: 250}, '⌠': {Wx: 686}, '⌡': {Wx: 686}, '〈': {Wx: 329}, '〉': {Wx: 329}, '◊': {Wx: 494}, '♠': {Wx: 753}, '♣': {Wx: 753}, '♥': {Wx: 753}, '♦': {Wx: 753}, '\uf6d9': {Wx: 790}, '\uf6da': {Wx: 790}, '\uf6db': {Wx: 890}, '\uf8e5': {Wx: 500}, '\uf8e6': {Wx: 603}, '\uf8e7': {Wx: 1000}, '\uf8e8': {Wx: 790}, '\uf8e9': {Wx: 790}, '\uf8ea': {Wx: 786}, '\uf8eb': {Wx: 384}, '\uf8ec': {Wx: 384}, '\uf8ed': {Wx: 384}, '\uf8ee': {Wx: 384}, '\uf8ef': {Wx: 384}, '\uf8f0': {Wx: 384}, '\uf8f1': {Wx: 494}, '\uf8f2': {Wx: 494}, '\uf8f3': {Wx: 494}, '\uf8f4': {Wx: 494}, '\uf8f5': {Wx: 686}, '\uf8f6': {Wx: 384}, '\uf8f7': {Wx: 384}, '\uf8f8': {Wx: 384}, '\uf8f9': {Wx: 384}, '\uf8fa': {Wx: 384}, '\uf8fb': {Wx: 384}, '\uf8fc': {Wx: 494}, '\uf8fd': {Wx: 494}, '\uf8fe': {Wx: 494}, '\uf8ff': {Wx: 790}}}

type RuneCharSafeMap struct {
	_dd map[rune]CharMetrics
	_df _fa.RWMutex
}

func NewFontFile2FromPdfObject(obj _bc.PdfObject) (TtfType, error) {
	obj = _bc.TraceToDirectObject(obj)
	_fbf, _gcd := obj.(*_bc.PdfObjectStream)
	if !_gcd {
		_gb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0073\u0074\u0072e\u0061\u006d \u0028\u0025\u0054\u0029", obj)
		return TtfType{}, _bc.ErrTypeError
	}
	_fab, _ggd := _bc.DecodeStream(_fbf)
	if _ggd != nil {
		return TtfType{}, _ggd
	}
	_cgb := ttfParser{_dada: _fe.NewReader(_fab)}
	return _cgb.Parse()
}

const (
	FontWeightMedium FontWeight = iota
	FontWeightBold
	FontWeightRoman
)

func _ad() StdFont {
	_aeb.Do(_aab)
	_egg := Descriptor{Name: HelveticaName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0020, BBox: [4]float64{-166, -225, 1000, 931}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_egg, _gffc)
}

var _ Font = StdFont{}

func (_dfd *ttfParser) ParseHmtx() error {
	if _gea := _dfd.Seek("\u0068\u006d\u0074\u0078"); _gea != nil {
		return _gea
	}
	_dfd._afe.Widths = make([]uint16, 0, 8)
	for _cfea := uint16(0); _cfea < _dfd._edbd; _cfea++ {
		_dfd._afe.Widths = append(_dfd._afe.Widths, _dfd.ReadUShort())
		_dfd.Skip(2)
	}
	if _dfd._edbd < _dfd._cbc && _dfd._edbd > 0 {
		_ged := _dfd._afe.Widths[_dfd._edbd-1]
		for _edfc := _dfd._edbd; _edfc < _dfd._cbc; _edfc++ {
			_dfd._afe.Widths = append(_dfd._afe.Widths, _ged)
		}
	}
	return nil
}
func (_gba *ttfParser) parseCmapFormat12() error {
	_dac := _gba.ReadULong()
	_gb.Log.Trace("\u0070\u0061\u0072se\u0043\u006d\u0061\u0070\u0046\u006f\u0072\u006d\u0061t\u00312\u003a \u0025s\u0020\u006e\u0075\u006d\u0047\u0072\u006f\u0075\u0070\u0073\u003d\u0025\u0064", _gba._afe.String(), _dac)
	for _gdb := uint32(0); _gdb < _dac; _gdb++ {
		_cebd := _gba.ReadULong()
		_ced := _gba.ReadULong()
		_gfb := _gba.ReadULong()
		if _cebd > 0x0010FFFF || (0xD800 <= _cebd && _cebd <= 0xDFFF) {
			return _c.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		if _ced < _cebd || _ced > 0x0010FFFF || (0xD800 <= _ced && _ced <= 0xDFFF) {
			return _c.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		for _faae := _cebd; _faae <= _ced; _faae++ {
			if _faae > 0x10FFFF {
				_gb.Log.Debug("\u0046\u006fr\u006d\u0061\u0074\u0020\u0031\u0032\u0020\u0063\u006d\u0061\u0070\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0062\u0065\u0079\u006f\u006e\u0064\u0020\u0055\u0043\u0053\u002d\u0034")
			}
			_gba._afe.Chars[rune(_faae)] = GID(_gfb)
			_gfb++
		}
	}
	return nil
}

type TtfType struct {
	UnitsPerEm             uint16
	PostScriptName         string
	Bold                   bool
	ItalicAngle            float64
	IsFixedPitch           bool
	TypoAscender           int16
	TypoDescender          int16
	UnderlinePosition      int16
	UnderlineThickness     int16
	Xmin, Ymin, Xmax, Ymax int16
	CapHeight              int16
	Widths                 []uint16
	Chars                  map[rune]GID
	GlyphNames             []GlyphName
}

func (_abe *ttfParser) ParseCmap() error {
	var _ddbc int64
	if _cec := _abe.Seek("\u0063\u006d\u0061\u0070"); _cec != nil {
		return _cec
	}
	_gb.Log.Trace("\u0050a\u0072\u0073\u0065\u0043\u006d\u0061p")
	_abe.ReadUShort()
	_dbgb := int(_abe.ReadUShort())
	_ddbf := int64(0)
	_dec := int64(0)
	_ggef := int64(0)
	for _cbb := 0; _cbb < _dbgb; _cbb++ {
		_fac := _abe.ReadUShort()
		_dcc := _abe.ReadUShort()
		_ddbc = int64(_abe.ReadULong())
		if _fac == 3 && _dcc == 1 {
			_dec = _ddbc
		} else if _fac == 3 && _dcc == 10 {
			_ggef = _ddbc
		} else if _fac == 1 && _dcc == 0 {
			_ddbf = _ddbc
		}
	}
	if _ddbf != 0 {
		if _dcf := _abe.parseCmapVersion(_ddbf); _dcf != nil {
			return _dcf
		}
	}
	if _dec != 0 {
		if _dbfd := _abe.parseCmapSubtable31(_dec); _dbfd != nil {
			return _dbfd
		}
	}
	if _ggef != 0 {
		if _gac := _abe.parseCmapVersion(_ggef); _gac != nil {
			return _gac
		}
	}
	if _dec == 0 && _ddbf == 0 && _ggef == 0 {
		_gb.Log.Debug("\u0074\u0074\u0066P\u0061\u0072\u0073\u0065\u0072\u002e\u0050\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u002e\u0020\u004e\u006f\u0020\u0033\u0031\u002c\u0020\u0031\u0030\u002c\u0020\u00331\u0030\u0020\u0074\u0061\u0062\u006c\u0065\u002e")
	}
	return nil
}

var _fba *RuneCharSafeMap
var _dab *RuneCharSafeMap

type GlyphName = _ca.GlyphName

func (_ccf *ttfParser) parseCmapSubtable31(_bcd int64) error {
	_aac := make([]rune, 0, 8)
	_add := make([]rune, 0, 8)
	_eged := make([]int16, 0, 8)
	_bd := make([]uint16, 0, 8)
	_ccf._afe.Chars = make(map[rune]GID)
	_ccf._dada.Seek(int64(_ccf._fgb["\u0063\u006d\u0061\u0070"])+_bcd, _d.SeekStart)
	_dc := _ccf.ReadUShort()
	if _dc != 4 {
		return _bb.Errorf("u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0073\u0075\u0062t\u0061\u0062\u006c\u0065\u0020\u0066\u006fr\u006d\u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0028\u0025w\u0029", _dc, _bc.ErrNotSupported)
	}
	_ccf.Skip(2 * 2)
	_fbge := int(_ccf.ReadUShort() / 2)
	_ccf.Skip(3 * 2)
	for _cbf := 0; _cbf < _fbge; _cbf++ {
		_add = append(_add, rune(_ccf.ReadUShort()))
	}
	_ccf.Skip(2)
	for _cef := 0; _cef < _fbge; _cef++ {
		_aac = append(_aac, rune(_ccf.ReadUShort()))
	}
	for _abd := 0; _abd < _fbge; _abd++ {
		_eged = append(_eged, _ccf.ReadShort())
	}
	_gfa, _ := _ccf._dada.Seek(int64(0), _d.SeekCurrent)
	for _fcc := 0; _fcc < _fbge; _fcc++ {
		_bd = append(_bd, _ccf.ReadUShort())
	}
	for _bgb := 0; _bgb < _fbge; _bgb++ {
		_gbg := _aac[_bgb]
		_gad := _add[_bgb]
		_ffb := _eged[_bgb]
		_ebf := _bd[_bgb]
		if _ebf > 0 {
			_ccf._dada.Seek(_gfa+2*int64(_bgb)+int64(_ebf), _d.SeekStart)
		}
		for _ecag := _gbg; _ecag <= _gad; _ecag++ {
			if _ecag == 0xFFFF {
				break
			}
			var _bac int32
			if _ebf > 0 {
				_bac = int32(_ccf.ReadUShort())
				if _bac > 0 {
					_bac += int32(_ffb)
				}
			} else {
				_bac = _ecag + int32(_ffb)
			}
			if _bac >= 65536 {
				_bac -= 65536
			}
			if _bac > 0 {
				_ccf._afe.Chars[_ecag] = GID(_bac)
			}
		}
	}
	return nil
}
func MakeRuneCharSafeMap(length int) *RuneCharSafeMap {
	return &RuneCharSafeMap{_dd: make(map[rune]CharMetrics, length)}
}

var _eea = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 611, 778, 778, 778, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 778, 778, 667, 667, 667, 667, 667, 944, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 611, 778, 722, 722, 722, 722, 556, 556, 556, 556, 556, 667, 667, 667, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 1000, 722, 722, 722, 722, 667, 667, 667, 667, 500, 500, 500, 500, 333, 500, 722, 500, 500, 833, 500, 500, 581, 520, 500, 930, 500, 556, 278, 220, 394, 394, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 556, 500, 500, 672, 556, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 556, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 333, 556, 556, 278, 278, 394, 278, 570, 549, 570, 494, 278, 833, 333, 570, 556, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 300, 330, 500, 500, 556, 540, 333, 333, 494, 1000, 250, 250, 1000, 570, 570, 556, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 444, 444, 549, 444, 444, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 333, 416, 333, 556, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func _fdg() StdFont {
	_aeb.Do(_aab)
	_bbc := Descriptor{Name: HelveticaObliqueName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-170, -225, 1116, 931}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_bbc, _dea)
}
func (_bdc *ttfParser) ParsePost() error {
	if _efcdd := _bdc.Seek("\u0070\u006f\u0073\u0074"); _efcdd != nil {
		return _efcdd
	}
	_ded := _bdc.Read32Fixed()
	_bdc._afe.ItalicAngle = _bdc.Read32Fixed()
	_bdc._afe.UnderlinePosition = _bdc.ReadShort()
	_bdc._afe.UnderlineThickness = _bdc.ReadShort()
	_bdc._afe.IsFixedPitch = _bdc.ReadULong() != 0
	_bdc.ReadULong()
	_bdc.ReadULong()
	_bdc.ReadULong()
	_bdc.ReadULong()
	_gb.Log.Trace("\u0050a\u0072\u0073\u0065\u0050\u006f\u0073\u0074\u003a\u0020\u0066\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065\u003d\u0025\u0066", _ded)
	switch _ded {
	case 1.0:
		_bdc._afe.GlyphNames = _gade
	case 2.0:
		_efcf := int(_bdc.ReadUShort())
		_feg := make([]int, _efcf)
		_bdc._afe.GlyphNames = make([]GlyphName, _efcf)
		_fgf := -1
		for _gda := 0; _gda < _efcf; _gda++ {
			_adgb := int(_bdc.ReadUShort())
			_feg[_gda] = _adgb
			if _adgb <= 0x7fff && _adgb > _fgf {
				_fgf = _adgb
			}
		}
		var _daa []GlyphName
		if _fgf >= len(_gade) {
			_daa = make([]GlyphName, _fgf-len(_gade)+1)
			for _dae := 0; _dae < _fgf-len(_gade)+1; _dae++ {
				_eaa := int(_bdc.readByte())
				_gdd, _fga := _bdc.ReadStr(_eaa)
				if _fga != nil {
					return _fga
				}
				_daa[_dae] = GlyphName(_gdd)
			}
		}
		for _caba := 0; _caba < _efcf; _caba++ {
			_egac := _feg[_caba]
			if _egac < len(_gade) {
				_bdc._afe.GlyphNames[_caba] = _gade[_egac]
			} else if _egac >= len(_gade) && _egac <= 32767 {
				_bdc._afe.GlyphNames[_caba] = _daa[_egac-len(_gade)]
			} else {
				_bdc._afe.GlyphNames[_caba] = "\u002e\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064"
			}
		}
	case 2.5:
		_gegg := make([]int, _bdc._cbc)
		for _eed := 0; _eed < len(_gegg); _eed++ {
			_aad := int(_bdc.ReadSByte())
			_gegg[_eed] = _eed + 1 + _aad
		}
		_bdc._afe.GlyphNames = make([]GlyphName, len(_gegg))
		for _gdfc := 0; _gdfc < len(_bdc._afe.GlyphNames); _gdfc++ {
			_dff := _gade[_gegg[_gdfc]]
			_bdc._afe.GlyphNames[_gdfc] = _dff
		}
	case 3.0:
		_gb.Log.Debug("\u004e\u006f\u0020\u0050\u006f\u0073t\u0053\u0063\u0072i\u0070\u0074\u0020n\u0061\u006d\u0065\u0020\u0069\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020is\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e")
	default:
		_gb.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020f\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065=\u0025\u0066", _ded)
	}
	return nil
}

var _fg = []rune{'A', 'Æ', 'Á', 'Ă', 'Â', 'Ä', 'À', 'Ā', 'Ą', 'Å', 'Ã', 'B', 'C', 'Ć', 'Č', 'Ç', 'D', 'Ď', 'Đ', '∆', 'E', 'É', 'Ě', 'Ê', 'Ë', 'Ė', 'È', 'Ē', 'Ę', 'Ð', '€', 'F', 'G', 'Ğ', 'Ģ', 'H', 'I', 'Í', 'Î', 'Ï', 'İ', 'Ì', 'Ī', 'Į', 'J', 'K', 'Ķ', 'L', 'Ĺ', 'Ľ', 'Ļ', 'Ł', 'M', 'N', 'Ń', 'Ň', 'Ņ', 'Ñ', 'O', 'Œ', 'Ó', 'Ô', 'Ö', 'Ò', 'Ő', 'Ō', 'Ø', 'Õ', 'P', 'Q', 'R', 'Ŕ', 'Ř', 'Ŗ', 'S', 'Ś', 'Š', 'Ş', 'Ș', 'T', 'Ť', 'Ţ', 'Þ', 'U', 'Ú', 'Û', 'Ü', 'Ù', 'Ű', 'Ū', 'Ų', 'Ů', 'V', 'W', 'X', 'Y', 'Ý', 'Ÿ', 'Z', 'Ź', 'Ž', 'Ż', 'a', 'á', 'ă', 'â', '´', 'ä', 'æ', 'à', 'ā', '&', 'ą', 'å', '^', '~', '*', '@', 'ã', 'b', '\\', '|', '{', '}', '[', ']', '˘', '¦', '•', 'c', 'ć', 'ˇ', 'č', 'ç', '¸', '¢', 'ˆ', ':', ',', '\uf6c3', '©', '¤', 'd', '†', '‡', 'ď', 'đ', '°', '¨', '÷', '$', '˙', 'ı', 'e', 'é', 'ě', 'ê', 'ë', 'ė', 'è', '8', '…', 'ē', '—', '–', 'ę', '=', 'ð', '!', '¡', 'f', 'ﬁ', '5', 'ﬂ', 'ƒ', '4', '⁄', 'g', 'ğ', 'ģ', 'ß', '`', '>', '≥', '«', '»', '‹', '›', 'h', '˝', '-', 'i', 'í', 'î', 'ï', 'ì', 'ī', 'į', 'j', 'k', 'ķ', 'l', 'ĺ', 'ľ', 'ļ', '<', '≤', '¬', '◊', 'ł', 'm', '¯', '−', 'µ', '×', 'n', 'ń', 'ň', 'ņ', '9', '≠', 'ñ', '#', 'o', 'ó', 'ô', 'ö', 'œ', '˛', 'ò', 'ő', 'ō', '1', '½', '¼', '¹', 'ª', 'º', 'ø', 'õ', 'p', '¶', '(', ')', '∂', '%', '.', '·', '‰', '+', '±', 'q', '?', '¿', '"', '„', '“', '”', '‘', '’', '‚', '\'', 'r', 'ŕ', '√', 'ř', 'ŗ', '®', '˚', 's', 'ś', 'š', 'ş', 'ș', '§', ';', '7', '6', '/', ' ', '£', '∑', 't', 'ť', 'ţ', 'þ', '3', '¾', '³', '˜', '™', '2', '²', 'u', 'ú', 'û', 'ü', 'ù', 'ű', 'ū', '_', 'ų', 'ů', 'v', 'w', 'x', 'y', 'ý', 'ÿ', '¥', 'z', 'ź', 'ž', 'ż', '0'}

func init() {
	RegisterStdFont(CourierName, _dbb, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u0043\u006f\u0075\u0072\u0069e\u0072\u004e\u0065\u0077", "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077")
	RegisterStdFont(CourierBoldName, _fcb, "\u0043o\u0075r\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(CourierObliqueName, _cbd, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(CourierBoldObliqueName, _agdg, "C\u006f\u0075\u0072\u0069er\u004ee\u0077\u002c\u0042\u006f\u006cd\u0049\u0074\u0061\u006c\u0069\u0063")
}

var _bg *RuneCharSafeMap

type Font interface {
	Encoder() _ca.TextEncoder
	GetRuneMetrics(_fc rune) (CharMetrics, bool)
}

func _ecg() StdFont {
	_fdge := _ca.NewSymbolEncoder()
	_ab := Descriptor{Name: SymbolName, Family: string(SymbolName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-180, -293, 1090, 1010}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 85, StemH: 92}
	return NewStdFontWithEncoding(_ab, _cff, _fdge)
}
func (_gc CharMetrics) String() string {
	return _f.Sprintf("<\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u003e", _gc.Wx, _gc.Wy)
}
func (_ccbg *ttfParser) ParseComponents() error {
	if _bbd := _ccbg.ParseHead(); _bbd != nil {
		return _bbd
	}
	if _abg := _ccbg.ParseHhea(); _abg != nil {
		return _abg
	}
	if _ebb := _ccbg.ParseMaxp(); _ebb != nil {
		return _ebb
	}
	if _ddbe := _ccbg.ParseHmtx(); _ddbe != nil {
		return _ddbe
	}
	if _, _age := _ccbg._fgb["\u006e\u0061\u006d\u0065"]; _age {
		if _fbg := _ccbg.ParseName(); _fbg != nil {
			return _fbg
		}
	}
	if _, _ebg := _ccbg._fgb["\u004f\u0053\u002f\u0032"]; _ebg {
		if _fbad := _ccbg.ParseOS2(); _fbad != nil {
			return _fbad
		}
	}
	if _, _gce := _ccbg._fgb["\u0070\u006f\u0073\u0074"]; _gce {
		if _defc := _ccbg.ParsePost(); _defc != nil {
			return _defc
		}
	}
	if _, _cfd := _ccbg._fgb["\u0063\u006d\u0061\u0070"]; _cfd {
		if _afg := _ccbg.ParseCmap(); _afg != nil {
			return _afg
		}
	}
	return nil
}

type FontWeight int

func TtfParseFile(fileStr string) (TtfType, error) {
	_cfad, _dde := _g.Open(fileStr)
	if _dde != nil {
		return TtfType{}, _dde
	}
	defer _cfad.Close()
	return TtfParse(_cfad)
}
func (_aa *RuneCharSafeMap) Range(f func(_gf rune, _eca CharMetrics) (_cg bool)) {
	_aa._df.RLock()
	defer _aa._df.RUnlock()
	for _ed, _ef := range _aa._dd {
		if f(_ed, _ef) {
			break
		}
	}
}

var _ecaf = []int16{611, 889, 611, 611, 611, 611, 611, 611, 611, 611, 611, 611, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 611, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 444, 667, 667, 556, 556, 611, 556, 556, 833, 667, 667, 667, 667, 667, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 611, 611, 611, 611, 500, 500, 500, 500, 500, 556, 556, 556, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 611, 833, 611, 556, 556, 556, 556, 556, 556, 556, 500, 500, 500, 500, 333, 500, 667, 500, 500, 778, 500, 500, 422, 541, 500, 920, 500, 500, 278, 275, 400, 400, 389, 389, 333, 275, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 760, 500, 500, 500, 500, 544, 500, 400, 333, 675, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 889, 444, 889, 500, 444, 675, 500, 333, 389, 278, 500, 500, 500, 500, 500, 167, 500, 500, 500, 500, 333, 675, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 444, 444, 278, 278, 300, 278, 675, 549, 675, 471, 278, 722, 333, 675, 500, 675, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 667, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 523, 333, 333, 476, 833, 250, 250, 1000, 675, 675, 500, 500, 500, 420, 556, 556, 556, 333, 333, 333, 214, 389, 389, 453, 389, 389, 760, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 300, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 444, 667, 444, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func _cba() StdFont {
	_bf.Do(_ggc)
	_edb := Descriptor{Name: TimesItalicName, Family: _caf, Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-169, -217, 1010, 883}, ItalicAngle: -15.5, Ascent: 683, Descent: -217, CapHeight: 653, XHeight: 441, StemV: 76, StemH: 32}
	return NewStdFont(_edb, _dbf)
}

var _eda = []int16{667, 944, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 667, 722, 722, 722, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 667, 667, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 889, 667, 611, 611, 611, 611, 611, 611, 611, 500, 500, 500, 500, 333, 500, 722, 500, 500, 778, 500, 500, 570, 570, 500, 832, 500, 500, 278, 220, 348, 348, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 500, 500, 500, 608, 500, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 389, 389, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 382, 278, 570, 549, 606, 494, 278, 778, 333, 606, 576, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 266, 300, 500, 500, 500, 500, 333, 333, 494, 833, 250, 250, 1000, 570, 570, 500, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 389, 389, 549, 389, 389, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 366, 278, 500, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 444, 667, 500, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func init() {
	RegisterStdFont(HelveticaName, _ad, "\u0041\u0072\u0069a\u006c")
	RegisterStdFont(HelveticaBoldName, _ddb, "\u0041\u0072\u0069\u0061\u006c\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(HelveticaObliqueName, _fdg, "\u0041\u0072\u0069a\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(HelveticaBoldObliqueName, _ege, "\u0041\u0072i\u0061\u006c\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}

var _de *RuneCharSafeMap

const (
	CourierName            = StdFontName("\u0043o\u0075\u0072\u0069\u0065\u0072")
	CourierBoldName        = StdFontName("\u0043\u006f\u0075r\u0069\u0065\u0072\u002d\u0042\u006f\u006c\u0064")
	CourierObliqueName     = StdFontName("\u0043o\u0075r\u0069\u0065\u0072\u002d\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
	CourierBoldObliqueName = StdFontName("\u0043\u006f\u0075\u0072ie\u0072\u002d\u0042\u006f\u006c\u0064\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
)

func (_caee *RuneCharSafeMap) Length() int {
	_caee._df.RLock()
	defer _caee._df.RUnlock()
	return len(_caee._dd)
}
func (_ecaa *ttfParser) parseCmapSubtable10(_fdgc int64) error {
	if _ecaa._afe.Chars == nil {
		_ecaa._afe.Chars = make(map[rune]GID)
	}
	_ecaa._dada.Seek(int64(_ecaa._fgb["\u0063\u006d\u0061\u0070"])+_fdgc, _d.SeekStart)
	var _eba, _dca uint32
	_cgde := _ecaa.ReadUShort()
	if _cgde < 8 {
		_eba = uint32(_ecaa.ReadUShort())
		_dca = uint32(_ecaa.ReadUShort())
	} else {
		_ecaa.ReadUShort()
		_eba = _ecaa.ReadULong()
		_dca = _ecaa.ReadULong()
	}
	_gb.Log.Trace("\u0070\u0061r\u0073\u0065\u0043\u006d\u0061p\u0053\u0075\u0062\u0074\u0061b\u006c\u0065\u0031\u0030\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065\u003d\u0025\u0064", _cgde, _eba, _dca)
	if _cgde != 0 {
		return _c.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006d\u0061p\u0020s\u0075\u0062\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	_cdcb, _dbac := _ecaa.ReadStr(256)
	if _dbac != nil {
		return _dbac
	}
	_defa := []byte(_cdcb)
	for _effg, _ccbgc := range _defa {
		_ecaa._afe.Chars[rune(_effg)] = GID(_ccbgc)
		if _ccbgc != 0 {
			_f.Printf("\u0009\u0030\u0078\u002502\u0078\u0020\u279e\u0020\u0030\u0078\u0025\u0030\u0032\u0078\u003d\u0025\u0063\u000a", _effg, _ccbgc, rune(_ccbgc))
		}
	}
	return nil
}
func (_faa *TtfType) String() string {
	return _f.Sprintf("\u0046\u004fN\u0054\u005f\u0046\u0049\u004cE\u0032\u007b\u0025\u0023\u0071 \u0055\u006e\u0069\u0074\u0073\u0050\u0065\u0072\u0045\u006d\u003d\u0025\u0064\u0020\u0042\u006f\u006c\u0064\u003d\u0025\u0074\u0020\u0049\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065\u003d\u0025\u0066\u0020"+"\u0043\u0061pH\u0065\u0069\u0067h\u0074\u003d\u0025\u0064 Ch\u0061rs\u003d\u0025\u0064\u0020\u0047\u006c\u0079ph\u004e\u0061\u006d\u0065\u0073\u003d\u0025d\u007d", _faa.PostScriptName, _faa.UnitsPerEm, _faa.Bold, _faa.ItalicAngle, _faa.CapHeight, len(_faa.Chars), len(_faa.GlyphNames))
}
func _dbe(_dafg map[string]uint32) string {
	var _fbeb []string
	for _aff := range _dafg {
		_fbeb = append(_fbeb, _aff)
	}
	_ea.Slice(_fbeb, func(_edf, _cbg int) bool { return _dafg[_fbeb[_edf]] < _dafg[_fbeb[_cbg]] })
	_fef := []string{_f.Sprintf("\u0054\u0072\u0075\u0065Ty\u0070\u0065\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_dafg))}
	for _, _fff := range _fbeb {
		_fef = append(_fef, _f.Sprintf("\u0009%\u0071\u0020\u0025\u0035\u0064", _fff, _dafg[_fff]))
	}
	return _b.Join(_fef, "\u000a")
}

var _ccd *RuneCharSafeMap

func (_daeb *ttfParser) ReadULong() (_ecd uint32) {
	_eaf.Read(_daeb._dada, _eaf.BigEndian, &_ecd)
	return _ecd
}
func (_gbcg *ttfParser) ReadStr(length int) (string, error) {
	_gfe := make([]byte, length)
	_ebe, _cfff := _gbcg._dada.Read(_gfe)
	if _cfff != nil {
		return "", _cfff
	} else if _ebe != length {
		return "", _f.Errorf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073", length)
	}
	return string(_gfe), nil
}

var _gffc *RuneCharSafeMap

const (
	HelveticaName            = StdFontName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	HelveticaBoldName        = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064")
	HelveticaObliqueName     = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065")
	HelveticaBoldObliqueName = StdFontName("H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065")
)

func (_cf *fontMap) read(_cgc StdFontName) (func() StdFont, bool) {
	_cf.Lock()
	defer _cf.Unlock()
	_fca, _ac := _cf._ecf[_cgc]
	return _fca, _ac
}
func (_efff *TtfType) MakeToUnicode() *_ff.CMap {
	_gag := make(map[_ff.CharCode]rune)
	if len(_efff.GlyphNames) == 0 {
		return _ff.NewToUnicodeCMap(_gag)
	}
	for _fbee, _egc := range _efff.Chars {
		_geb := _ff.CharCode(_fbee)
		_bee := _efff.GlyphNames[_egc]
		_ggca, _bge := _ca.GlyphToRune(_bee)
		if !_bge {
			_gb.Log.Debug("\u004e\u006f \u0072\u0075\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u0020\u0067\u006c\u0079\u0070h=\u0025\u0071", _fbee, _bee)
			_ggca = _ca.MissingCodeRune
		}
		_gag[_geb] = _ggca
	}
	return _ff.NewToUnicodeCMap(_gag)
}
func init() {
	RegisterStdFont(SymbolName, _ecg, "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063", "S\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064", "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(ZapfDingbatsName, _daf)
}

type Descriptor struct {
	Name        StdFontName
	Family      string
	Weight      FontWeight
	Flags       uint
	BBox        [4]float64
	ItalicAngle float64
	Ascent      float64
	Descent     float64
	CapHeight   float64
	XHeight     float64
	StemV       float64
	StemH       float64
}

const (
	_caf                = "\u0054\u0069\u006de\u0073"
	TimesRomanName      = StdFontName("T\u0069\u006d\u0065\u0073\u002d\u0052\u006f\u006d\u0061\u006e")
	TimesBoldName       = StdFontName("\u0054\u0069\u006d\u0065\u0073\u002d\u0042\u006f\u006c\u0064")
	TimesItalicName     = StdFontName("\u0054\u0069\u006de\u0073\u002d\u0049\u0074\u0061\u006c\u0069\u0063")
	TimesBoldItalicName = StdFontName("\u0054\u0069m\u0065\u0073\u002dB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
)

func _agdg() StdFont {
	_cab.Do(_fd)
	_geg := Descriptor{Name: CourierBoldObliqueName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0061, BBox: [4]float64{-57, -250, 869, 801}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_geg, _cfc)
}
func _ceb() StdFont {
	_bf.Do(_ggc)
	_eff := Descriptor{Name: TimesBoldItalicName, Family: _caf, Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-200, -218, 996, 921}, ItalicAngle: -15, Ascent: 683, Descent: -217, CapHeight: 669, XHeight: 462, StemV: 121, StemH: 42}
	return NewStdFont(_eff, _ccdb)
}
func (_ae StdFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_cag, _gff := _ae._eee.Read(r)
	return _cag, _gff
}
func init() {
	RegisterStdFont(TimesRomanName, _dba, "\u0054\u0069\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e", "\u0054\u0069\u006de\u0073")
	RegisterStdFont(TimesBoldName, _fgd, "\u0054i\u006de\u0073\u004e\u0065\u0077\u0052o\u006d\u0061n\u002c\u0042\u006f\u006c\u0064", "\u0054\u0069\u006d\u0065\u0073\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(TimesItalicName, _cba, "T\u0069m\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006da\u006e\u002c\u0049\u0074al\u0069\u0063", "\u0054\u0069\u006de\u0073\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(TimesBoldItalicName, _ceb, "\u0054i\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e,\u0042\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063", "\u0054\u0069m\u0065\u0073\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}
func _fgd() StdFont {
	_bf.Do(_ggc)
	_fee := Descriptor{Name: TimesBoldName, Family: _caf, Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 935}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 676, XHeight: 461, StemV: 139, StemH: 44}
	return NewStdFont(_fee, _cbag)
}
func (_bcgc *ttfParser) ReadShort() (_cdb int16) {
	_eaf.Read(_bcgc._dada, _eaf.BigEndian, &_cdb)
	return _cdb
}
func _ege() StdFont {
	_aeb.Do(_aab)
	_edd := Descriptor{Name: HelveticaBoldObliqueName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-174, -228, 1114, 962}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_edd, _ecc)
}
func _ddb() StdFont {
	_aeb.Do(_aab)
	_cfaf := Descriptor{Name: HelveticaBoldName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-170, -228, 1003, 962}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_cfaf, _fba)
}

const (
	SymbolName       = StdFontName("\u0053\u0079\u006d\u0062\u006f\u006c")
	ZapfDingbatsName = StdFontName("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
)

func (_ecb *fontMap) write(_ecff StdFontName, _fb func() StdFont) {
	_ecb.Lock()
	defer _ecb.Unlock()
	_ecb._ecf[_ecff] = _fb
}
func (_dda *ttfParser) ReadSByte() (_eae int8) {
	_eaf.Read(_dda._dada, _eaf.BigEndian, &_eae)
	return _eae
}
func _dbb() StdFont {
	_cab.Do(_fd)
	_ffe := Descriptor{Name: CourierName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0021, BBox: [4]float64{-23, -250, 715, 805}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_ffe, _bg)
}

var _gade = []GlyphName{"\u002en\u006f\u0074\u0064\u0065\u0066", "\u002e\u006e\u0075l\u006c", "\u006e\u006fn\u006d\u0061\u0072k\u0069\u006e\u0067\u0072\u0065\u0074\u0075\u0072\u006e", "\u0073\u0070\u0061c\u0065", "\u0065\u0078\u0063\u006c\u0061\u006d", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c", "\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u0069\u0067\u006e", "\u0064\u006f\u006c\u006c\u0061\u0072", "\u0070e\u0072\u0063\u0065\u006e\u0074", "\u0061m\u0070\u0065\u0072\u0073\u0061\u006ed", "q\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067\u006c\u0065", "\u0070a\u0072\u0065\u006e\u006c\u0065\u0066t", "\u0070\u0061\u0072\u0065\u006e\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0074\u0065\u0072\u0069\u0073\u006b", "\u0070\u006c\u0075\u0073", "\u0063\u006f\u006dm\u0061", "\u0068\u0079\u0070\u0068\u0065\u006e", "\u0070\u0065\u0072\u0069\u006f\u0064", "\u0073\u006c\u0061s\u0068", "\u007a\u0065\u0072\u006f", "\u006f\u006e\u0065", "\u0074\u0077\u006f", "\u0074\u0068\u0072e\u0065", "\u0066\u006f\u0075\u0072", "\u0066\u0069\u0076\u0065", "\u0073\u0069\u0078", "\u0073\u0065\u0076e\u006e", "\u0065\u0069\u0067h\u0074", "\u006e\u0069\u006e\u0065", "\u0063\u006f\u006co\u006e", "\u0073e\u006d\u0069\u0063\u006f\u006c\u006fn", "\u006c\u0065\u0073\u0073", "\u0065\u0071\u0075a\u006c", "\u0067r\u0065\u0061\u0074\u0065\u0072", "\u0071\u0075\u0065\u0073\u0074\u0069\u006f\u006e", "\u0061\u0074", "\u0041", "\u0042", "\u0043", "\u0044", "\u0045", "\u0046", "\u0047", "\u0048", "\u0049", "\u004a", "\u004b", "\u004c", "\u004d", "\u004e", "\u004f", "\u0050", "\u0051", "\u0052", "\u0053", "\u0054", "\u0055", "\u0056", "\u0057", "\u0058", "\u0059", "\u005a", "b\u0072\u0061\u0063\u006b\u0065\u0074\u006c\u0065\u0066\u0074", "\u0062a\u0063\u006b\u0073\u006c\u0061\u0073h", "\u0062\u0072\u0061c\u006b\u0065\u0074\u0072\u0069\u0067\u0068\u0074", "a\u0073\u0063\u0069\u0069\u0063\u0069\u0072\u0063\u0075\u006d", "\u0075\u006e\u0064\u0065\u0072\u0073\u0063\u006f\u0072\u0065", "\u0067\u0072\u0061v\u0065", "\u0061", "\u0062", "\u0063", "\u0064", "\u0065", "\u0066", "\u0067", "\u0068", "\u0069", "\u006a", "\u006b", "\u006c", "\u006d", "\u006e", "\u006f", "\u0070", "\u0071", "\u0072", "\u0073", "\u0074", "\u0075", "\u0076", "\u0077", "\u0078", "\u0079", "\u007a", "\u0062r\u0061\u0063\u0065\u006c\u0065\u0066t", "\u0062\u0061\u0072", "\u0062\u0072\u0061\u0063\u0065\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0063\u0069\u0069\u0074\u0069\u006c\u0064\u0065", "\u0041d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0041\u0072\u0069n\u0067", "\u0043\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0045\u0061\u0063\u0075\u0074\u0065", "\u004e\u0074\u0069\u006c\u0064\u0065", "\u004fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0055d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0061\u0063\u0075\u0074\u0065", "\u0061\u0067\u0072\u0061\u0076\u0065", "a\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0074\u0069\u006c\u0064\u0065", "\u0061\u0072\u0069n\u0067", "\u0063\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0065\u0061\u0063\u0075\u0074\u0065", "\u0065\u0067\u0072\u0061\u0076\u0065", "e\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0065d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0069\u0061\u0063\u0075\u0074\u0065", "\u0069\u0067\u0072\u0061\u0076\u0065", "i\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0069d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006e\u0074\u0069\u006c\u0064\u0065", "\u006f\u0061\u0063\u0075\u0074\u0065", "\u006f\u0067\u0072\u0061\u0076\u0065", "o\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u006fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006f\u0074\u0069\u006c\u0064\u0065", "\u0075\u0061\u0063\u0075\u0074\u0065", "\u0075\u0067\u0072\u0061\u0076\u0065", "u\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0075d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0064\u0061\u0067\u0067\u0065\u0072", "\u0064\u0065\u0067\u0072\u0065\u0065", "\u0063\u0065\u006e\u0074", "\u0073\u0074\u0065\u0072\u006c\u0069\u006e\u0067", "\u0073e\u0063\u0074\u0069\u006f\u006e", "\u0062\u0075\u006c\u006c\u0065\u0074", "\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h", "\u0067\u0065\u0072\u006d\u0061\u006e\u0064\u0062\u006c\u0073", "\u0072\u0065\u0067\u0069\u0073\u0074\u0065\u0072\u0065\u0064", "\u0063o\u0070\u0079\u0072\u0069\u0067\u0068t", "\u0074r\u0061\u0064\u0065\u006d\u0061\u0072k", "\u0061\u0063\u0075t\u0065", "\u0064\u0069\u0065\u0072\u0065\u0073\u0069\u0073", "\u006e\u006f\u0074\u0065\u0071\u0075\u0061\u006c", "\u0041\u0045", "\u004f\u0073\u006c\u0061\u0073\u0068", "\u0069\u006e\u0066\u0069\u006e\u0069\u0074\u0079", "\u0070l\u0075\u0073\u006d\u0069\u006e\u0075s", "\u006ce\u0073\u0073\u0065\u0071\u0075\u0061l", "\u0067\u0072\u0065a\u0074\u0065\u0072\u0065\u0071\u0075\u0061\u006c", "\u0079\u0065\u006e", "\u006d\u0075", "p\u0061\u0072\u0074\u0069\u0061\u006c\u0064\u0069\u0066\u0066", "\u0073u\u006d\u006d\u0061\u0074\u0069\u006fn", "\u0070r\u006f\u0064\u0075\u0063\u0074", "\u0070\u0069", "\u0069\u006e\u0074\u0065\u0067\u0072\u0061\u006c", "o\u0072\u0064\u0066\u0065\u006d\u0069\u006e\u0069\u006e\u0065", "\u006f\u0072\u0064m\u0061\u0073\u0063\u0075\u006c\u0069\u006e\u0065", "\u004f\u006d\u0065g\u0061", "\u0061\u0065", "\u006f\u0073\u006c\u0061\u0073\u0068", "\u0071\u0075\u0065s\u0074\u0069\u006f\u006e\u0064\u006f\u0077\u006e", "\u0065\u0078\u0063\u006c\u0061\u006d\u0064\u006f\u0077\u006e", "\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u006e\u006f\u0074", "\u0072a\u0064\u0069\u0063\u0061\u006c", "\u0066\u006c\u006f\u0072\u0069\u006e", "a\u0070\u0070\u0072\u006f\u0078\u0065\u0071\u0075\u0061\u006c", "\u0044\u0065\u006ct\u0061", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074r\u0069\u0067\u0068\u0074", "\u0065\u006c\u006c\u0069\u0070\u0073\u0069\u0073", "\u006e\u006fn\u0062\u0072\u0065a\u006b\u0069\u006e\u0067\u0073\u0070\u0061\u0063\u0065", "\u0041\u0067\u0072\u0061\u0076\u0065", "\u0041\u0074\u0069\u006c\u0064\u0065", "\u004f\u0074\u0069\u006c\u0064\u0065", "\u004f\u0045", "\u006f\u0065", "\u0065\u006e\u0064\u0061\u0073\u0068", "\u0065\u006d\u0064\u0061\u0073\u0068", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u006c\u0065\u0066\u0074", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c\u0072\u0069\u0067\u0068\u0074", "\u0071u\u006f\u0074\u0065\u006c\u0065\u0066t", "\u0071\u0075\u006f\u0074\u0065\u0072\u0069\u0067\u0068\u0074", "\u0064\u0069\u0076\u0069\u0064\u0065", "\u006co\u007a\u0065\u006e\u0067\u0065", "\u0079d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0059d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e", "\u0063\u0075\u0072\u0072\u0065\u006e\u0063\u0079", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006c\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006cr\u0069\u0067\u0068\u0074", "\u0066\u0069", "\u0066\u006c", "\u0064a\u0067\u0067\u0065\u0072\u0064\u0062l", "\u0070\u0065\u0072\u0069\u006f\u0064\u0063\u0065\u006et\u0065\u0072\u0065\u0064", "\u0071\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067l\u0062\u0061\u0073\u0065", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u0062\u0061\u0073\u0065", "p\u0065\u0072\u0074\u0068\u006f\u0075\u0073\u0061\u006e\u0064", "A\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "E\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0041\u0061\u0063\u0075\u0074\u0065", "\u0045d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0045\u0067\u0072\u0061\u0076\u0065", "\u0049\u0061\u0063\u0075\u0074\u0065", "I\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0049d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0049\u0067\u0072\u0061\u0076\u0065", "\u004f\u0061\u0063\u0075\u0074\u0065", "O\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061\u0070\u0070l\u0065", "\u004f\u0067\u0072\u0061\u0076\u0065", "\u0055\u0061\u0063\u0075\u0074\u0065", "U\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0055\u0067\u0072\u0061\u0076\u0065", "\u0064\u006f\u0074\u006c\u0065\u0073\u0073\u0069", "\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0074\u0069\u006cd\u0065", "\u006d\u0061\u0063\u0072\u006f\u006e", "\u0062\u0072\u0065v\u0065", "\u0064o\u0074\u0061\u0063\u0063\u0065\u006et", "\u0072\u0069\u006e\u0067", "\u0063e\u0064\u0069\u006c\u006c\u0061", "\u0068\u0075\u006eg\u0061\u0072\u0075\u006d\u006c\u0061\u0075\u0074", "\u006f\u0067\u006f\u006e\u0065\u006b", "\u0063\u0061\u0072o\u006e", "\u004c\u0073\u006c\u0061\u0073\u0068", "\u006c\u0073\u006c\u0061\u0073\u0068", "\u0053\u0063\u0061\u0072\u006f\u006e", "\u0073\u0063\u0061\u0072\u006f\u006e", "\u005a\u0063\u0061\u0072\u006f\u006e", "\u007a\u0063\u0061\u0072\u006f\u006e", "\u0062r\u006f\u006b\u0065\u006e\u0062\u0061r", "\u0045\u0074\u0068", "\u0065\u0074\u0068", "\u0059\u0061\u0063\u0075\u0074\u0065", "\u0079\u0061\u0063\u0075\u0074\u0065", "\u0054\u0068\u006fr\u006e", "\u0074\u0068\u006fr\u006e", "\u006d\u0069\u006eu\u0073", "\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "o\u006e\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "t\u0077\u006f\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u0074\u0068\u0072\u0065\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u006fn\u0065\u0068\u0061\u006c\u0066", "\u006f\u006e\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072", "\u0074\u0068\u0072\u0065\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072\u0073", "\u0066\u0072\u0061n\u0063", "\u0047\u0062\u0072\u0065\u0076\u0065", "\u0067\u0062\u0072\u0065\u0076\u0065", "\u0049\u0064\u006f\u0074\u0061\u0063\u0063\u0065\u006e\u0074", "\u0053\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0073\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0043\u0061\u0063\u0075\u0074\u0065", "\u0063\u0061\u0063\u0075\u0074\u0065", "\u0043\u0063\u0061\u0072\u006f\u006e", "\u0063\u0063\u0061\u0072\u006f\u006e", "\u0064\u0063\u0072\u006f\u0061\u0074"}

type StdFontName string

func (_gbd *ttfParser) ParseHead() error {
	if _ebc := _gbd.Seek("\u0068\u0065\u0061\u0064"); _ebc != nil {
		return _ebc
	}
	_gbd.Skip(3 * 4)
	_agc := _gbd.ReadULong()
	if _agc != 0x5F0F3CF5 {
		_gb.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0063\u006fr\u0072e\u0063\u0074\u0020\u006d\u0061\u0067\u0069\u0063\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e\u0020\u0046\u006fn\u0074\u0020\u006d\u0061\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0069\u0073\u0070\u006c\u0061\u0079\u0020\u0063\u006f\u0072\u0072\u0065\u0063t\u006c\u0079\u002e\u0020\u0025\u0073", _gbd)
	}
	_gbd.Skip(2)
	_gbd._afe.UnitsPerEm = _gbd.ReadUShort()
	_gbd.Skip(2 * 8)
	_gbd._afe.Xmin = _gbd.ReadShort()
	_gbd._afe.Ymin = _gbd.ReadShort()
	_gbd._afe.Xmax = _gbd.ReadShort()
	_gbd._afe.Ymax = _gbd.ReadShort()
	return nil
}
func (_ega StdFont) ToPdfObject() _bc.PdfObject {
	_gg := _bc.MakeDict()
	_gg.Set("\u0054\u0079\u0070\u0065", _bc.MakeName("\u0046\u006f\u006e\u0074"))
	_gg.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _bc.MakeName("\u0054\u0079\u0070e\u0031"))
	_gg.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _bc.MakeName(_ega.Name()))
	_gg.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _ega._cfe.ToPdfObject())
	return _bc.MakeIndirectObject(_gg)
}

var _aag = []int16{667, 1000, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 500, 667, 667, 556, 556, 556, 556, 556, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 667, 556, 556, 469, 584, 389, 1015, 556, 556, 278, 260, 334, 334, 278, 278, 333, 260, 350, 500, 500, 333, 500, 500, 333, 556, 333, 278, 278, 250, 737, 556, 556, 556, 556, 643, 556, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 556, 278, 333, 278, 500, 556, 500, 556, 556, 167, 556, 556, 556, 611, 333, 584, 549, 556, 556, 333, 333, 556, 333, 333, 222, 278, 278, 278, 278, 278, 222, 222, 500, 500, 222, 222, 299, 222, 584, 549, 584, 471, 222, 833, 333, 584, 556, 584, 556, 556, 556, 556, 556, 549, 556, 556, 556, 556, 556, 556, 944, 333, 556, 556, 556, 556, 834, 834, 333, 370, 365, 611, 556, 556, 537, 333, 333, 476, 889, 278, 278, 1000, 584, 584, 556, 556, 611, 355, 333, 333, 333, 222, 222, 222, 191, 333, 333, 453, 333, 333, 737, 333, 500, 500, 500, 500, 500, 556, 278, 556, 556, 278, 278, 556, 600, 278, 317, 278, 556, 556, 834, 333, 333, 1000, 556, 333, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 500, 722, 500, 500, 500, 500, 556, 500, 500, 500, 500, 556}

func (_bad *TtfType) NewEncoder() _ca.TextEncoder { return _ca.NewTrueTypeFontEncoder(_bad.Chars) }
func (_cca *ttfParser) ParseName() error {
	if _eef := _cca.Seek("\u006e\u0061\u006d\u0065"); _eef != nil {
		return _eef
	}
	_fefg, _ := _cca._dada.Seek(0, _d.SeekCurrent)
	_cca._afe.PostScriptName = ""
	_cca.Skip(2)
	_gdf := _cca.ReadUShort()
	_ddfd := _cca.ReadUShort()
	for _dcec := uint16(0); _dcec < _gdf && _cca._afe.PostScriptName == ""; _dcec++ {
		_cca.Skip(3 * 2)
		_ddc := _cca.ReadUShort()
		_cdcf := _cca.ReadUShort()
		_dfc := _cca.ReadUShort()
		if _ddc == 6 {
			_cca._dada.Seek(_fefg+int64(_ddfd)+int64(_dfc), _d.SeekStart)
			_dccg, _abdc := _cca.ReadStr(int(_cdcf))
			if _abdc != nil {
				return _abdc
			}
			_dccg = _b.Replace(_dccg, "\u0000", "", -1)
			_gdc, _abdc := _be.Compile("\u005b\u0028\u0029\u007b\u007d\u003c\u003e\u0020\u002f%\u005b\u005c\u005d\u005d")
			if _abdc != nil {
				return _abdc
			}
			_cca._afe.PostScriptName = _gdc.ReplaceAllString(_dccg, "")
		}
	}
	if _cca._afe.PostScriptName == "" {
		_gb.Log.Debug("\u0050a\u0072\u0073e\u004e\u0061\u006de\u003a\u0020\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u0050\u006f\u0073t\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0077\u0061\u0073\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	}
	return nil
}
func _fd() {
	const _gd = 600
	_bg = MakeRuneCharSafeMap(len(_fg))
	for _, _fbe := range _fg {
		_bg.Write(_fbe, CharMetrics{Wx: _gd})
	}
	_de = _bg.Copy()
	_cfc = _bg.Copy()
	_ccd = _bg.Copy()
}

type GID = _ca.GID

func (_bgdg *ttfParser) parseCmapVersion(_gagg int64) error {
	_gb.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006da\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a \u006f\u0066\u0066s\u0065t\u003d\u0025\u0064", _gagg)
	if _bgdg._afe.Chars == nil {
		_bgdg._afe.Chars = make(map[rune]GID)
	}
	_bgdg._dada.Seek(int64(_bgdg._fgb["\u0063\u006d\u0061\u0070"])+_gagg, _d.SeekStart)
	var _abef, _faf uint32
	_ecca := _bgdg.ReadUShort()
	if _ecca < 8 {
		_abef = uint32(_bgdg.ReadUShort())
		_faf = uint32(_bgdg.ReadUShort())
	} else {
		_bgdg.ReadUShort()
		_abef = _bgdg.ReadULong()
		_faf = _bgdg.ReadULong()
	}
	_gb.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0043m\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064 \u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067u\u0061g\u0065\u003d\u0025\u0064", _ecca, _abef, _faf)
	switch _ecca {
	case 0:
		return _bgdg.parseCmapFormat0()
	case 6:
		return _bgdg.parseCmapFormat6()
	case 12:
		return _bgdg.parseCmapFormat12()
	default:
		_gb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063m\u0061\u0070\u0020\u0066\u006f\u0072\u006da\u0074\u003d\u0025\u0064", _ecca)
		return nil
	}
}
func (_aae *ttfParser) ParseMaxp() error {
	if _bef := _aae.Seek("\u006d\u0061\u0078\u0070"); _bef != nil {
		return _bef
	}
	_aae.Skip(4)
	_aae._cbc = _aae.ReadUShort()
	return nil
}
func (_ag *RuneCharSafeMap) Write(b rune, r CharMetrics) {
	_ag._df.Lock()
	defer _ag._df.Unlock()
	_ag._dd[b] = r
}
func (_cdc StdFont) Encoder() _ca.TextEncoder { return _cdc._cfe }

var _dea *RuneCharSafeMap

func IsStdFont(name StdFontName) bool { _, _eg := _cd.read(name); return _eg }
func (_ebce *ttfParser) parseCmapFormat6() error {
	_dce := int(_ebce.ReadUShort())
	_cda := int(_ebce.ReadUShort())
	_gb.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u0046o\u0072\u006d\u0061\u0074\u0036\u003a\u0020%s\u0020\u0066\u0069\u0072s\u0074\u0043\u006f\u0064\u0065\u003d\u0025\u0064\u0020en\u0074\u0072y\u0043\u006f\u0075\u006e\u0074\u003d\u0025\u0064", _ebce._afe.String(), _dce, _cda)
	for _cagg := 0; _cagg < _cda; _cagg++ {
		_bga := GID(_ebce.ReadUShort())
		_ebce._afe.Chars[rune(_cagg+_dce)] = _bga
	}
	return nil
}
func RegisterStdFont(name StdFontName, fnc func() StdFont, aliases ...StdFontName) {
	if _, _beb := _cd.read(name); _beb {
		panic("\u0066o\u006e\u0074\u0020\u0061l\u0072\u0065\u0061\u0064\u0079 \u0072e\u0067i\u0073\u0074\u0065\u0072\u0065\u0064\u003a " + string(name))
	}
	_cd.write(name, fnc)
	for _, _bcf := range aliases {
		RegisterStdFont(_bcf, fnc)
	}
}
func TtfParse(r _d.ReadSeeker) (TtfType, error) { _cgg := &ttfParser{_dada: r}; return _cgg.Parse() }
func (_eafg *ttfParser) ParseOS2() error {
	if _edc := _eafg.Seek("\u004f\u0053\u002f\u0032"); _edc != nil {
		return _edc
	}
	_dabg := _eafg.ReadUShort()
	_eafg.Skip(4 * 2)
	_eafg.Skip(11*2 + 10 + 4*4 + 4)
	_beg := _eafg.ReadUShort()
	_eafg._afe.Bold = (_beg & 32) != 0
	_eafg.Skip(2 * 2)
	_eafg._afe.TypoAscender = _eafg.ReadShort()
	_eafg._afe.TypoDescender = _eafg.ReadShort()
	if _dabg >= 2 {
		_eafg.Skip(3*2 + 2*4 + 2)
		_eafg._afe.CapHeight = _eafg.ReadShort()
	} else {
		_eafg._afe.CapHeight = 0
	}
	return nil
}
func (_cae *RuneCharSafeMap) Read(b rune) (CharMetrics, bool) {
	_cae._df.RLock()
	defer _cae._df.RUnlock()
	_ec, _ece := _cae._dd[b]
	return _ec, _ece
}

var _ba = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 556, 722, 722, 611, 611, 611, 611, 611, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 722, 556, 556, 584, 584, 389, 975, 556, 611, 278, 280, 389, 389, 333, 333, 333, 280, 350, 556, 556, 333, 556, 556, 333, 556, 333, 333, 278, 250, 737, 556, 611, 556, 556, 743, 611, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 611, 333, 333, 333, 611, 556, 611, 556, 556, 167, 611, 611, 611, 611, 333, 584, 549, 556, 556, 333, 333, 611, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 556, 556, 278, 278, 400, 278, 584, 549, 584, 494, 278, 889, 333, 584, 611, 584, 611, 611, 611, 611, 556, 549, 611, 556, 611, 611, 611, 611, 944, 333, 611, 611, 611, 556, 834, 834, 333, 370, 365, 611, 611, 611, 556, 333, 333, 494, 889, 278, 278, 1000, 584, 584, 611, 611, 611, 474, 500, 500, 500, 278, 278, 278, 238, 389, 389, 549, 389, 389, 737, 333, 556, 556, 556, 556, 556, 556, 333, 556, 556, 278, 278, 556, 600, 333, 389, 333, 611, 556, 834, 333, 333, 1000, 556, 333, 611, 611, 611, 611, 611, 611, 611, 556, 611, 611, 556, 778, 556, 556, 556, 556, 556, 500, 500, 500, 500, 556}
var _ecc *RuneCharSafeMap
var _def = []int16{722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 556, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 389, 722, 722, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 556, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 556, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 611, 611, 611, 611, 444, 444, 444, 444, 333, 444, 667, 444, 444, 778, 444, 444, 469, 541, 500, 921, 444, 500, 278, 200, 480, 480, 333, 333, 333, 200, 350, 444, 444, 333, 444, 444, 333, 500, 333, 278, 250, 250, 760, 500, 500, 500, 500, 588, 500, 400, 333, 564, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 564, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 564, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 344, 278, 564, 549, 564, 471, 278, 778, 333, 564, 500, 564, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 453, 333, 333, 476, 833, 250, 250, 1000, 564, 564, 500, 444, 444, 408, 444, 444, 444, 333, 333, 333, 180, 333, 333, 453, 333, 333, 760, 333, 389, 389, 389, 389, 389, 500, 278, 500, 500, 278, 250, 500, 600, 278, 326, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func _daf() StdFont {
	_cgd := _ca.NewZapfDingbatsEncoder()
	_ce := Descriptor{Name: ZapfDingbatsName, Family: string(ZapfDingbatsName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-1, -143, 981, 820}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 90, StemH: 28}
	return NewStdFontWithEncoding(_ce, _gaf, _cgd)
}
func (_dad *TtfType) MakeEncoder() (_ca.SimpleEncoder, error) {
	_aebd := make(map[_ca.CharCode]GlyphName)
	for _af := _ca.CharCode(0); _af <= 256; _af++ {
		_efc := rune(_af)
		_fdf, _fbag := _dad.Chars[_efc]
		if !_fbag {
			continue
		}
		var _ffc GlyphName
		if int(_fdf) >= 0 && int(_fdf) < len(_dad.GlyphNames) {
			_ffc = _dad.GlyphNames[_fdf]
		} else {
			_fgg := rune(_fdf)
			if _gfgc, _fcbf := _ca.RuneToGlyph(_fgg); _fcbf {
				_ffc = _gfgc
			}
		}
		if _ffc != "" {
			_aebd[_af] = _ffc
		}
	}
	if len(_aebd) == 0 {
		_gb.Log.Debug("WA\u0052\u004eI\u004e\u0047\u003a\u0020\u005a\u0065\u0072\u006f\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002e\u0020\u0074\u0074\u0066=\u0025s\u0020\u0043\u0068\u0061\u0072\u0073\u003d\u005b%\u00200\u0032\u0078]", _dad, _dad.Chars)
	}
	return _ca.NewCustomSimpleTextEncoder(_aebd, nil)
}
func NewStdFontWithEncoding(desc Descriptor, metrics *RuneCharSafeMap, encoder _ca.TextEncoder) StdFont {
	var _bcg rune = 0xA0
	if _, _agg := metrics.Read(_bcg); !_agg {
		_cfed, _ := metrics.Read(0x20)
		metrics.Write(_bcg, _cfed)
	}
	return StdFont{_eab: desc, _eee: metrics, _cfe: encoder}
}
func (_eb *RuneCharSafeMap) Copy() *RuneCharSafeMap {
	_dg := MakeRuneCharSafeMap(_eb.Length())
	_eb.Range(func(_a rune, _ee CharMetrics) (_cc bool) {
		_dg._dd[_a] = _ee
		return false
	})
	return _dg
}
func _aab() {
	_gffc = MakeRuneCharSafeMap(len(_fg))
	_fba = MakeRuneCharSafeMap(len(_fg))
	for _ggf, _gfg := range _fg {
		_gffc.Write(_gfg, CharMetrics{Wx: float64(_aag[_ggf])})
		_fba.Write(_gfg, CharMetrics{Wx: float64(_ba[_ggf])})
	}
	_dea = _gffc.Copy()
	_ecc = _fba.Copy()
}
func (_gebc *ttfParser) Skip(n int) { _gebc._dada.Seek(int64(n), _d.SeekCurrent) }
func _fcb() StdFont {
	_cab.Do(_fd)
	_acf := Descriptor{Name: CourierBoldName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0021, BBox: [4]float64{-113, -250, 749, 801}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_acf, _de)
}

var _ccdb *RuneCharSafeMap
var _dbf *RuneCharSafeMap

type StdFont struct {
	_eab Descriptor
	_eee *RuneCharSafeMap
	_cfe _ca.TextEncoder
}

var _aeb _fa.Once

func (_eag *ttfParser) Parse() (TtfType, error) {
	_bcbf, _dbg := _eag.ReadStr(4)
	if _dbg != nil {
		return TtfType{}, _dbg
	}
	if _bcbf == "\u004f\u0054\u0054\u004f" {
		return TtfType{}, _bb.Errorf("\u0066\u006f\u006e\u0074s\u0020\u0062\u0061\u0073\u0065\u0064\u0020\u006f\u006e \u0050\u006f\u0073\u0074\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065s\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0028\u0025\u0077\u0029", _bc.ErrNotSupported)
	}
	if _bcbf != "\u0000\u0001\u0000\u0000" && _bcbf != "\u0074\u0072\u0075\u0065" {
		_gb.Log.Debug("\u0055n\u0072\u0065c\u006f\u0067\u006ei\u007a\u0065\u0064\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066i\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u002e\u0020v\u0065\u0072\u0073\u0069\u006f\u006e\u003d\u0025\u0071", _bcbf)
	}
	_dfe := int(_eag.ReadUShort())
	_eag.Skip(3 * 2)
	_eag._fgb = make(map[string]uint32)
	var _efcd string
	for _afd := 0; _afd < _dfe; _afd++ {
		_efcd, _dbg = _eag.ReadStr(4)
		if _dbg != nil {
			return TtfType{}, _dbg
		}
		_eag.Skip(4)
		_fcae := _eag.ReadULong()
		_eag.Skip(4)
		_eag._fgb[_efcd] = _fcae
	}
	_gb.Log.Trace(_dbe(_eag._fgb))
	if _dbg = _eag.ParseComponents(); _dbg != nil {
		return TtfType{}, _dbg
	}
	return _eag._afe, nil
}
func NewStdFont(desc Descriptor, metrics *RuneCharSafeMap) StdFont {
	return NewStdFontWithEncoding(desc, metrics, _ca.NewStandardEncoder())
}
func (_ge StdFont) Name() string { return string(_ge._eab.Name) }
func (_bgg *ttfParser) Seek(tag string) error {
	_bcdd, _bce := _bgg._fgb[tag]
	if !_bce {
		return _f.Errorf("\u0074\u0061\u0062\u006ce \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0073", tag)
	}
	_bgg._dada.Seek(int64(_bcdd), _d.SeekStart)
	return nil
}
func _ggc() {
	_dab = MakeRuneCharSafeMap(len(_fg))
	_cbag = MakeRuneCharSafeMap(len(_fg))
	_ccdb = MakeRuneCharSafeMap(len(_fg))
	_dbf = MakeRuneCharSafeMap(len(_fg))
	for _bfa, _adg := range _fg {
		_dab.Write(_adg, CharMetrics{Wx: float64(_def[_bfa])})
		_cbag.Write(_adg, CharMetrics{Wx: float64(_eea[_bfa])})
		_ccdb.Write(_adg, CharMetrics{Wx: float64(_eda[_bfa])})
		_dbf.Write(_adg, CharMetrics{Wx: float64(_ecaf[_bfa])})
	}
}

var _gaf = &RuneCharSafeMap{_dd: map[rune]CharMetrics{' ': {Wx: 278}, '→': {Wx: 838}, '↔': {Wx: 1016}, '↕': {Wx: 458}, '①': {Wx: 788}, '②': {Wx: 788}, '③': {Wx: 788}, '④': {Wx: 788}, '⑤': {Wx: 788}, '⑥': {Wx: 788}, '⑦': {Wx: 788}, '⑧': {Wx: 788}, '⑨': {Wx: 788}, '⑩': {Wx: 788}, '■': {Wx: 761}, '▲': {Wx: 892}, '▼': {Wx: 892}, '◆': {Wx: 788}, '●': {Wx: 791}, '◗': {Wx: 438}, '★': {Wx: 816}, '☎': {Wx: 719}, '☛': {Wx: 960}, '☞': {Wx: 939}, '♠': {Wx: 626}, '♣': {Wx: 776}, '♥': {Wx: 694}, '♦': {Wx: 595}, '✁': {Wx: 974}, '✂': {Wx: 961}, '✃': {Wx: 974}, '✄': {Wx: 980}, '✆': {Wx: 789}, '✇': {Wx: 790}, '✈': {Wx: 791}, '✉': {Wx: 690}, '✌': {Wx: 549}, '✍': {Wx: 855}, '✎': {Wx: 911}, '✏': {Wx: 933}, '✐': {Wx: 911}, '✑': {Wx: 945}, '✒': {Wx: 974}, '✓': {Wx: 755}, '✔': {Wx: 846}, '✕': {Wx: 762}, '✖': {Wx: 761}, '✗': {Wx: 571}, '✘': {Wx: 677}, '✙': {Wx: 763}, '✚': {Wx: 760}, '✛': {Wx: 759}, '✜': {Wx: 754}, '✝': {Wx: 494}, '✞': {Wx: 552}, '✟': {Wx: 537}, '✠': {Wx: 577}, '✡': {Wx: 692}, '✢': {Wx: 786}, '✣': {Wx: 788}, '✤': {Wx: 788}, '✥': {Wx: 790}, '✦': {Wx: 793}, '✧': {Wx: 794}, '✩': {Wx: 823}, '✪': {Wx: 789}, '✫': {Wx: 841}, '✬': {Wx: 823}, '✭': {Wx: 833}, '✮': {Wx: 816}, '✯': {Wx: 831}, '✰': {Wx: 923}, '✱': {Wx: 744}, '✲': {Wx: 723}, '✳': {Wx: 749}, '✴': {Wx: 790}, '✵': {Wx: 792}, '✶': {Wx: 695}, '✷': {Wx: 776}, '✸': {Wx: 768}, '✹': {Wx: 792}, '✺': {Wx: 759}, '✻': {Wx: 707}, '✼': {Wx: 708}, '✽': {Wx: 682}, '✾': {Wx: 701}, '✿': {Wx: 826}, '❀': {Wx: 815}, '❁': {Wx: 789}, '❂': {Wx: 789}, '❃': {Wx: 707}, '❄': {Wx: 687}, '❅': {Wx: 696}, '❆': {Wx: 689}, '❇': {Wx: 786}, '❈': {Wx: 787}, '❉': {Wx: 713}, '❊': {Wx: 791}, '❋': {Wx: 785}, '❍': {Wx: 873}, '❏': {Wx: 762}, '❐': {Wx: 762}, '❑': {Wx: 759}, '❒': {Wx: 759}, '❖': {Wx: 784}, '❘': {Wx: 138}, '❙': {Wx: 277}, '❚': {Wx: 415}, '❛': {Wx: 392}, '❜': {Wx: 392}, '❝': {Wx: 668}, '❞': {Wx: 668}, '❡': {Wx: 732}, '❢': {Wx: 544}, '❣': {Wx: 544}, '❤': {Wx: 910}, '❥': {Wx: 667}, '❦': {Wx: 760}, '❧': {Wx: 760}, '❶': {Wx: 788}, '❷': {Wx: 788}, '❸': {Wx: 788}, '❹': {Wx: 788}, '❺': {Wx: 788}, '❻': {Wx: 788}, '❼': {Wx: 788}, '❽': {Wx: 788}, '❾': {Wx: 788}, '❿': {Wx: 788}, '➀': {Wx: 788}, '➁': {Wx: 788}, '➂': {Wx: 788}, '➃': {Wx: 788}, '➄': {Wx: 788}, '➅': {Wx: 788}, '➆': {Wx: 788}, '➇': {Wx: 788}, '➈': {Wx: 788}, '➉': {Wx: 788}, '➊': {Wx: 788}, '➋': {Wx: 788}, '➌': {Wx: 788}, '➍': {Wx: 788}, '➎': {Wx: 788}, '➏': {Wx: 788}, '➐': {Wx: 788}, '➑': {Wx: 788}, '➒': {Wx: 788}, '➓': {Wx: 788}, '➔': {Wx: 894}, '➘': {Wx: 748}, '➙': {Wx: 924}, '➚': {Wx: 748}, '➛': {Wx: 918}, '➜': {Wx: 927}, '➝': {Wx: 928}, '➞': {Wx: 928}, '➟': {Wx: 834}, '➠': {Wx: 873}, '➡': {Wx: 828}, '➢': {Wx: 924}, '➣': {Wx: 924}, '➤': {Wx: 917}, '➥': {Wx: 930}, '➦': {Wx: 931}, '➧': {Wx: 463}, '➨': {Wx: 883}, '➩': {Wx: 836}, '➪': {Wx: 836}, '➫': {Wx: 867}, '➬': {Wx: 867}, '➭': {Wx: 696}, '➮': {Wx: 696}, '➯': {Wx: 874}, '➱': {Wx: 874}, '➲': {Wx: 760}, '➳': {Wx: 946}, '➴': {Wx: 771}, '➵': {Wx: 865}, '➶': {Wx: 771}, '➷': {Wx: 888}, '➸': {Wx: 967}, '➹': {Wx: 888}, '➺': {Wx: 831}, '➻': {Wx: 873}, '➼': {Wx: 927}, '➽': {Wx: 970}, '➾': {Wx: 918}, '\uf8d7': {Wx: 390}, '\uf8d8': {Wx: 390}, '\uf8d9': {Wx: 317}, '\uf8da': {Wx: 317}, '\uf8db': {Wx: 276}, '\uf8dc': {Wx: 276}, '\uf8dd': {Wx: 509}, '\uf8de': {Wx: 509}, '\uf8df': {Wx: 410}, '\uf8e0': {Wx: 410}, '\uf8e1': {Wx: 234}, '\uf8e2': {Wx: 234}, '\uf8e3': {Wx: 334}, '\uf8e4': {Wx: 334}}}

func NewStdFontByName(name StdFontName) (StdFont, bool) {
	_ffd, _efg := _cd.read(name)
	if !_efg {
		return StdFont{}, false
	}
	return _ffd(), true
}
func (_gfc *ttfParser) ParseHhea() error {
	if _bae := _gfc.Seek("\u0068\u0068\u0065\u0061"); _bae != nil {
		return _bae
	}
	_gfc.Skip(4 + 15*2)
	_gfc._edbd = _gfc.ReadUShort()
	return nil
}
func _dba() StdFont {
	_bf.Do(_ggc)
	_bgd := Descriptor{Name: TimesRomanName, Family: _caf, Weight: FontWeightRoman, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 898}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 662, XHeight: 450, StemV: 84, StemH: 28}
	return NewStdFont(_bgd, _dab)
}

type fontMap struct {
	_fa.Mutex
	_ecf map[StdFontName]func() StdFont
}

var _bf _fa.Once
