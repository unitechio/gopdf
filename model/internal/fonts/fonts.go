package fonts

import (
	_af "bytes"
	_fg "encoding/binary"
	_dd "errors"
	_eg "fmt"
	_dc "io"
	_d "os"
	_f "regexp"
	_b "sort"
	_a "strings"
	_ad "sync"

	_ae "bitbucket.org/shenghui0779/gopdf/common"
	_gc "bitbucket.org/shenghui0779/gopdf/core"
	_g "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_dg "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
)

const (
	HelveticaName            = StdFontName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	HelveticaBoldName        = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064")
	HelveticaObliqueName     = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065")
	HelveticaBoldObliqueName = StdFontName("H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065")
)

func (_bef StdFont) Name() string { return string(_bef._be.Name) }
func _dag(_afed map[string]uint32) string {
	var _ea []string
	for _gbcc := range _afed {
		_ea = append(_ea, _gbcc)
	}
	_b.Slice(_ea, func(_caeg, _agf int) bool { return _afed[_ea[_caeg]] < _afed[_ea[_agf]] })
	_eaa := []string{_eg.Sprintf("\u0054\u0072\u0075\u0065Ty\u0070\u0065\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_afed))}
	for _, _cdc := range _ea {
		_eaa = append(_eaa, _eg.Sprintf("\u0009%\u0071\u0020\u0025\u0035\u0064", _cdc, _afed[_cdc]))
	}
	return _a.Join(_eaa, "\u000a")
}
func TtfParse(r _dc.ReadSeeker) (TtfType, error) { _ced := &ttfParser{_cdb: r}; return _ced.Parse() }
func (_ca *RuneCharSafeMap) Range(f func(_adc rune, _cb CharMetrics) (_fb bool)) {
	_ca._fc.RLock()
	defer _ca._fc.RUnlock()
	for _fbc, _bb := range _ca._ee {
		if f(_fbc, _bb) {
			break
		}
	}
}
func (_cdaa *ttfParser) Seek(tag string) error {
	_dfee, _adg := _cdaa._eef[tag]
	if !_adg {
		return _eg.Errorf("\u0074\u0061\u0062\u006ce \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0073", tag)
	}
	_cdaa._cdb.Seek(int64(_dfee), _dc.SeekStart)
	return nil
}

var _gcgg *RuneCharSafeMap

func (_ffa StdFont) GetMetricsTable() *RuneCharSafeMap { return _ffa._ge }

type GlyphName = _dg.GlyphName

func _afc() StdFont {
	_bbf.Do(_gaa)
	_cd := Descriptor{Name: CourierBoldObliqueName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0061, BBox: [4]float64{-57, -250, 869, 801}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_cd, _gag)
}
func MakeRuneCharSafeMap(length int) *RuneCharSafeMap {
	return &RuneCharSafeMap{_ee: make(map[rune]CharMetrics, length)}
}

type fontMap struct {
	_ad.Mutex
	_gca map[StdFontName]func() StdFont
}

func (_bbg *ttfParser) ParseComponents() error {
	if _eaf := _bbg.ParseHead(); _eaf != nil {
		return _eaf
	}
	if _cf := _bbg.ParseHhea(); _cf != nil {
		return _cf
	}
	if _ddd := _bbg.ParseMaxp(); _ddd != nil {
		return _ddd
	}
	if _dfff := _bbg.ParseHmtx(); _dfff != nil {
		return _dfff
	}
	if _, _dfe := _bbg._eef["\u006e\u0061\u006d\u0065"]; _dfe {
		if _ba := _bbg.ParseName(); _ba != nil {
			return _ba
		}
	}
	if _, _fbg := _bbg._eef["\u004f\u0053\u002f\u0032"]; _fbg {
		if _bdgg := _bbg.ParseOS2(); _bdgg != nil {
			return _bdgg
		}
	}
	if _, _add := _bbg._eef["\u0070\u006f\u0073\u0074"]; _add {
		if _ada := _bbg.ParsePost(); _ada != nil {
			return _ada
		}
	}
	if _, _eeece := _bbg._eef["\u0063\u006d\u0061\u0070"]; _eeece {
		if _gga := _bbg.ParseCmap(); _gga != nil {
			return _gga
		}
	}
	return nil
}
func (_dec *ttfParser) parseCmapFormat6() error {
	_gbg := int(_dec.ReadUShort())
	_gbf := int(_dec.ReadUShort())
	_ae.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u0046o\u0072\u006d\u0061\u0074\u0036\u003a\u0020%s\u0020\u0066\u0069\u0072s\u0074\u0043\u006f\u0064\u0065\u003d\u0025\u0064\u0020en\u0074\u0072y\u0043\u006f\u0075\u006e\u0074\u003d\u0025\u0064", _dec._fce.String(), _gbg, _gbf)
	for _gdc := 0; _gdc < _gbf; _gdc++ {
		_cdfb := GID(_dec.ReadUShort())
		_dec._fce.Chars[rune(_gdc+_gbg)] = _cdfb
	}
	return nil
}
func (_fea StdFont) Encoder() _dg.TextEncoder { return _fea._df }
func (_cc *RuneCharSafeMap) Read(b rune) (CharMetrics, bool) {
	_cc._fc.RLock()
	defer _cc._fc.RUnlock()
	_ef, _ec := _cc._ee[b]
	return _ef, _ec
}
func _ebd() StdFont {
	_cbdc.Do(_dcaf)
	_dge := Descriptor{Name: TimesRomanName, Family: _ebe, Weight: FontWeightRoman, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 898}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 662, XHeight: 450, StemV: 84, StemH: 28}
	return NewStdFont(_dge, _ecg)
}
func NewStdFontByName(name StdFontName) (StdFont, bool) {
	_caf, _fe := _ed.read(name)
	if !_fe {
		return StdFont{}, false
	}
	return _caf(), true
}
func (_fcgc *ttfParser) parseCmapFormat12() error {
	_dgfg := _fcgc.ReadULong()
	_ae.Log.Trace("\u0070\u0061\u0072se\u0043\u006d\u0061\u0070\u0046\u006f\u0072\u006d\u0061t\u00312\u003a \u0025s\u0020\u006e\u0075\u006d\u0047\u0072\u006f\u0075\u0070\u0073\u003d\u0025\u0064", _fcgc._fce.String(), _dgfg)
	for _cdbd := uint32(0); _cdbd < _dgfg; _cdbd++ {
		_cea := _fcgc.ReadULong()
		_fgea := _fcgc.ReadULong()
		_cab := _fcgc.ReadULong()
		if _cea > 0x0010FFFF || (0xD800 <= _cea && _cea <= 0xDFFF) {
			return _dd.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		if _fgea < _cea || _fgea > 0x0010FFFF || (0xD800 <= _fgea && _fgea <= 0xDFFF) {
			return _dd.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		for _gcd := _cea; _gcd <= _fgea; _gcd++ {
			if _gcd > 0x10FFFF {
				_ae.Log.Debug("\u0046\u006fr\u006d\u0061\u0074\u0020\u0031\u0032\u0020\u0063\u006d\u0061\u0070\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0062\u0065\u0079\u006f\u006e\u0064\u0020\u0055\u0043\u0053\u002d\u0034")
			}
			_fcgc._fce.Chars[rune(_gcd)] = GID(_cab)
			_cab++
		}
	}
	return nil
}

type RuneCharSafeMap struct {
	_ee map[rune]CharMetrics
	_fc _ad.RWMutex
}

func _eec() StdFont {
	_ffc := _dg.NewSymbolEncoder()
	_eee := Descriptor{Name: SymbolName, Family: string(SymbolName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-180, -293, 1090, 1010}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 85, StemH: 92}
	return NewStdFontWithEncoding(_eee, _fbab, _ffc)
}

var _gcaa *RuneCharSafeMap

func _efa() StdFont {
	_fcg.Do(_dea)
	_deb := Descriptor{Name: HelveticaName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0020, BBox: [4]float64{-166, -225, 1000, 931}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_deb, _gcaa)
}

var _bdb *RuneCharSafeMap
var _eba = []int16{722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 556, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 389, 722, 722, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 556, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 556, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 611, 611, 611, 611, 444, 444, 444, 444, 333, 444, 667, 444, 444, 778, 444, 444, 469, 541, 500, 921, 444, 500, 278, 200, 480, 480, 333, 333, 333, 200, 350, 444, 444, 333, 444, 444, 333, 500, 333, 278, 250, 250, 760, 500, 500, 500, 500, 588, 500, 400, 333, 564, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 564, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 564, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 344, 278, 564, 549, 564, 471, 278, 778, 333, 564, 500, 564, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 453, 333, 333, 476, 833, 250, 250, 1000, 564, 564, 500, 444, 444, 408, 444, 444, 444, 333, 333, 333, 180, 333, 333, 453, 333, 333, 760, 333, 389, 389, 389, 389, 389, 500, 278, 500, 500, 278, 250, 500, 600, 278, 326, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func _gcc() StdFont {
	_cbdc.Do(_dcaf)
	_fdc := Descriptor{Name: TimesBoldName, Family: _ebe, Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 935}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 676, XHeight: 461, StemV: 139, StemH: 44}
	return NewStdFont(_fdc, _gcgg)
}

type FontWeight int

func _ded() StdFont {
	_fcg.Do(_dea)
	_db := Descriptor{Name: HelveticaObliqueName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-170, -225, 1116, 931}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_db, _ebb)
}
func _efg() StdFont {
	_fcg.Do(_dea)
	_ag := Descriptor{Name: HelveticaBoldName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-170, -228, 1003, 962}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_ag, _ebg)
}
func _gd() StdFont {
	_aebf := _dg.NewZapfDingbatsEncoder()
	_fag := Descriptor{Name: ZapfDingbatsName, Family: string(ZapfDingbatsName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-1, -143, 981, 820}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 90, StemH: 28}
	return NewStdFontWithEncoding(_fag, _gec, _aebf)
}

var _ce = []int16{611, 889, 611, 611, 611, 611, 611, 611, 611, 611, 611, 611, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 611, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 444, 667, 667, 556, 556, 611, 556, 556, 833, 667, 667, 667, 667, 667, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 611, 611, 611, 611, 500, 500, 500, 500, 500, 556, 556, 556, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 611, 833, 611, 556, 556, 556, 556, 556, 556, 556, 500, 500, 500, 500, 333, 500, 667, 500, 500, 778, 500, 500, 422, 541, 500, 920, 500, 500, 278, 275, 400, 400, 389, 389, 333, 275, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 760, 500, 500, 500, 500, 544, 500, 400, 333, 675, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 889, 444, 889, 500, 444, 675, 500, 333, 389, 278, 500, 500, 500, 500, 500, 167, 500, 500, 500, 500, 333, 675, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 444, 444, 278, 278, 300, 278, 675, 549, 675, 471, 278, 722, 333, 675, 500, 675, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 667, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 523, 333, 333, 476, 833, 250, 250, 1000, 675, 675, 500, 500, 500, 420, 556, 556, 556, 333, 333, 333, 214, 389, 389, 453, 389, 389, 760, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 300, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 444, 667, 444, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func _geg() StdFont {
	_fcg.Do(_dea)
	_ac := Descriptor{Name: HelveticaBoldObliqueName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-174, -228, 1114, 962}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_ac, _fff)
}
func init() {
	RegisterStdFont(CourierName, _fgc, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u0043\u006f\u0075\u0072\u0069e\u0072\u004e\u0065\u0077", "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077")
	RegisterStdFont(CourierBoldName, _cbd, "\u0043o\u0075r\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(CourierObliqueName, _dff, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(CourierBoldObliqueName, _afc, "C\u006f\u0075\u0072\u0069er\u004ee\u0077\u002c\u0042\u006f\u006cd\u0049\u0074\u0061\u006c\u0069\u0063")
}

type StdFontName string

var _aga = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 556, 722, 722, 611, 611, 611, 611, 611, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 722, 556, 556, 584, 584, 389, 975, 556, 611, 278, 280, 389, 389, 333, 333, 333, 280, 350, 556, 556, 333, 556, 556, 333, 556, 333, 333, 278, 250, 737, 556, 611, 556, 556, 743, 611, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 611, 333, 333, 333, 611, 556, 611, 556, 556, 167, 611, 611, 611, 611, 333, 584, 549, 556, 556, 333, 333, 611, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 556, 556, 278, 278, 400, 278, 584, 549, 584, 494, 278, 889, 333, 584, 611, 584, 611, 611, 611, 611, 556, 549, 611, 556, 611, 611, 611, 611, 944, 333, 611, 611, 611, 556, 834, 834, 333, 370, 365, 611, 611, 611, 556, 333, 333, 494, 889, 278, 278, 1000, 584, 584, 611, 611, 611, 474, 500, 500, 500, 278, 278, 278, 238, 389, 389, 549, 389, 389, 737, 333, 556, 556, 556, 556, 556, 556, 333, 556, 556, 278, 278, 556, 600, 333, 389, 333, 611, 556, 834, 333, 333, 1000, 556, 333, 611, 611, 611, 611, 611, 611, 611, 556, 611, 611, 556, 778, 556, 556, 556, 556, 556, 500, 500, 500, 500, 556}
var _ddc *RuneCharSafeMap

func NewStdFontWithEncoding(desc Descriptor, metrics *RuneCharSafeMap, encoder _dg.TextEncoder) StdFont {
	var _gbc rune = 0xA0
	if _, _gbb := metrics.Read(_gbc); !_gbb {
		_fec, _ := metrics.Read(0x20)
		metrics.Write(_gbc, _fec)
	}
	return StdFont{_be: desc, _ge: metrics, _df: encoder}
}

type StdFont struct {
	_be Descriptor
	_ge *RuneCharSafeMap
	_df _dg.TextEncoder
}

func (_bebc *ttfParser) ParseHhea() error {
	if _edag := _bebc.Seek("\u0068\u0068\u0065\u0061"); _edag != nil {
		return _edag
	}
	_bebc.Skip(4 + 15*2)
	_bebc._adf = _bebc.ReadUShort()
	return nil
}
func TtfParseFile(fileStr string) (TtfType, error) {
	_bdf, _dfbef := _d.Open(fileStr)
	if _dfbef != nil {
		return TtfType{}, _dfbef
	}
	defer _bdf.Close()
	return TtfParse(_bdf)
}
func (_fgd *ttfParser) ReadUShort() (_adcg uint16) {
	_fg.Read(_fgd._cdb, _fg.BigEndian, &_adcg)
	return _adcg
}

var _fbab = &RuneCharSafeMap{_ee: map[rune]CharMetrics{' ': {Wx: 250}, '!': {Wx: 333}, '#': {Wx: 500}, '%': {Wx: 833}, '&': {Wx: 778}, '(': {Wx: 333}, ')': {Wx: 333}, '+': {Wx: 549}, ',': {Wx: 250}, '.': {Wx: 250}, '/': {Wx: 278}, '0': {Wx: 500}, '1': {Wx: 500}, '2': {Wx: 500}, '3': {Wx: 500}, '4': {Wx: 500}, '5': {Wx: 500}, '6': {Wx: 500}, '7': {Wx: 500}, '8': {Wx: 500}, '9': {Wx: 500}, ':': {Wx: 278}, ';': {Wx: 278}, '<': {Wx: 549}, '=': {Wx: 549}, '>': {Wx: 549}, '?': {Wx: 444}, '[': {Wx: 333}, ']': {Wx: 333}, '_': {Wx: 500}, '{': {Wx: 480}, '|': {Wx: 200}, '}': {Wx: 480}, '¬': {Wx: 713}, '°': {Wx: 400}, '±': {Wx: 549}, 'µ': {Wx: 576}, '×': {Wx: 549}, '÷': {Wx: 549}, 'ƒ': {Wx: 500}, 'Α': {Wx: 722}, 'Β': {Wx: 667}, 'Γ': {Wx: 603}, 'Ε': {Wx: 611}, 'Ζ': {Wx: 611}, 'Η': {Wx: 722}, 'Θ': {Wx: 741}, 'Ι': {Wx: 333}, 'Κ': {Wx: 722}, 'Λ': {Wx: 686}, 'Μ': {Wx: 889}, 'Ν': {Wx: 722}, 'Ξ': {Wx: 645}, 'Ο': {Wx: 722}, 'Π': {Wx: 768}, 'Ρ': {Wx: 556}, 'Σ': {Wx: 592}, 'Τ': {Wx: 611}, 'Υ': {Wx: 690}, 'Φ': {Wx: 763}, 'Χ': {Wx: 722}, 'Ψ': {Wx: 795}, 'α': {Wx: 631}, 'β': {Wx: 549}, 'γ': {Wx: 411}, 'δ': {Wx: 494}, 'ε': {Wx: 439}, 'ζ': {Wx: 494}, 'η': {Wx: 603}, 'θ': {Wx: 521}, 'ι': {Wx: 329}, 'κ': {Wx: 549}, 'λ': {Wx: 549}, 'ν': {Wx: 521}, 'ξ': {Wx: 493}, 'ο': {Wx: 549}, 'π': {Wx: 549}, 'ρ': {Wx: 549}, 'ς': {Wx: 439}, 'σ': {Wx: 603}, 'τ': {Wx: 439}, 'υ': {Wx: 576}, 'φ': {Wx: 521}, 'χ': {Wx: 549}, 'ψ': {Wx: 686}, 'ω': {Wx: 686}, 'ϑ': {Wx: 631}, 'ϒ': {Wx: 620}, 'ϕ': {Wx: 603}, 'ϖ': {Wx: 713}, '•': {Wx: 460}, '…': {Wx: 1000}, '′': {Wx: 247}, '″': {Wx: 411}, '⁄': {Wx: 167}, '€': {Wx: 750}, 'ℑ': {Wx: 686}, '℘': {Wx: 987}, 'ℜ': {Wx: 795}, 'Ω': {Wx: 768}, 'ℵ': {Wx: 823}, '←': {Wx: 987}, '↑': {Wx: 603}, '→': {Wx: 987}, '↓': {Wx: 603}, '↔': {Wx: 1042}, '↵': {Wx: 658}, '⇐': {Wx: 987}, '⇑': {Wx: 603}, '⇒': {Wx: 987}, '⇓': {Wx: 603}, '⇔': {Wx: 1042}, '∀': {Wx: 713}, '∂': {Wx: 494}, '∃': {Wx: 549}, '∅': {Wx: 823}, '∆': {Wx: 612}, '∇': {Wx: 713}, '∈': {Wx: 713}, '∉': {Wx: 713}, '∋': {Wx: 439}, '∏': {Wx: 823}, '∑': {Wx: 713}, '−': {Wx: 549}, '∗': {Wx: 500}, '√': {Wx: 549}, '∝': {Wx: 713}, '∞': {Wx: 713}, '∠': {Wx: 768}, '∧': {Wx: 603}, '∨': {Wx: 603}, '∩': {Wx: 768}, '∪': {Wx: 768}, '∫': {Wx: 274}, '∴': {Wx: 863}, '∼': {Wx: 549}, '≅': {Wx: 549}, '≈': {Wx: 549}, '≠': {Wx: 549}, '≡': {Wx: 549}, '≤': {Wx: 549}, '≥': {Wx: 549}, '⊂': {Wx: 713}, '⊃': {Wx: 713}, '⊄': {Wx: 713}, '⊆': {Wx: 713}, '⊇': {Wx: 713}, '⊕': {Wx: 768}, '⊗': {Wx: 768}, '⊥': {Wx: 658}, '⋅': {Wx: 250}, '⌠': {Wx: 686}, '⌡': {Wx: 686}, '〈': {Wx: 329}, '〉': {Wx: 329}, '◊': {Wx: 494}, '♠': {Wx: 753}, '♣': {Wx: 753}, '♥': {Wx: 753}, '♦': {Wx: 753}, '\uf6d9': {Wx: 790}, '\uf6da': {Wx: 790}, '\uf6db': {Wx: 890}, '\uf8e5': {Wx: 500}, '\uf8e6': {Wx: 603}, '\uf8e7': {Wx: 1000}, '\uf8e8': {Wx: 790}, '\uf8e9': {Wx: 790}, '\uf8ea': {Wx: 786}, '\uf8eb': {Wx: 384}, '\uf8ec': {Wx: 384}, '\uf8ed': {Wx: 384}, '\uf8ee': {Wx: 384}, '\uf8ef': {Wx: 384}, '\uf8f0': {Wx: 384}, '\uf8f1': {Wx: 494}, '\uf8f2': {Wx: 494}, '\uf8f3': {Wx: 494}, '\uf8f4': {Wx: 494}, '\uf8f5': {Wx: 686}, '\uf8f6': {Wx: 384}, '\uf8f7': {Wx: 384}, '\uf8f8': {Wx: 384}, '\uf8f9': {Wx: 384}, '\uf8fa': {Wx: 384}, '\uf8fb': {Wx: 384}, '\uf8fc': {Wx: 494}, '\uf8fd': {Wx: 494}, '\uf8fe': {Wx: 494}, '\uf8ff': {Wx: 790}}}

const (
	_ebe                = "\u0054\u0069\u006de\u0073"
	TimesRomanName      = StdFontName("T\u0069\u006d\u0065\u0073\u002d\u0052\u006f\u006d\u0061\u006e")
	TimesBoldName       = StdFontName("\u0054\u0069\u006d\u0065\u0073\u002d\u0042\u006f\u006c\u0064")
	TimesItalicName     = StdFontName("\u0054\u0069\u006de\u0073\u002d\u0049\u0074\u0061\u006c\u0069\u0063")
	TimesBoldItalicName = StdFontName("\u0054\u0069m\u0065\u0073\u002dB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
)

func (_adb *RuneCharSafeMap) Write(b rune, r CharMetrics) {
	_adb._fc.Lock()
	defer _adb._fc.Unlock()
	_adb._ee[b] = r
}
func init() {
	RegisterStdFont(HelveticaName, _efa, "\u0041\u0072\u0069a\u006c")
	RegisterStdFont(HelveticaBoldName, _efg, "\u0041\u0072\u0069\u0061\u006c\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(HelveticaObliqueName, _ded, "\u0041\u0072\u0069a\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(HelveticaBoldObliqueName, _geg, "\u0041\u0072i\u0061\u006c\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}
func (_gabf *ttfParser) ReadStr(length int) (string, error) {
	_abe := make([]byte, length)
	_caa, _becb := _gabf._cdb.Read(_abe)
	if _becb != nil {
		return "", _becb
	} else if _caa != length {
		return "", _eg.Errorf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073", length)
	}
	return string(_abe), nil
}
func (_dca StdFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_abc, _egd := _dca._ge.Read(r)
	return _abc, _egd
}

var _gec = &RuneCharSafeMap{_ee: map[rune]CharMetrics{' ': {Wx: 278}, '→': {Wx: 838}, '↔': {Wx: 1016}, '↕': {Wx: 458}, '①': {Wx: 788}, '②': {Wx: 788}, '③': {Wx: 788}, '④': {Wx: 788}, '⑤': {Wx: 788}, '⑥': {Wx: 788}, '⑦': {Wx: 788}, '⑧': {Wx: 788}, '⑨': {Wx: 788}, '⑩': {Wx: 788}, '■': {Wx: 761}, '▲': {Wx: 892}, '▼': {Wx: 892}, '◆': {Wx: 788}, '●': {Wx: 791}, '◗': {Wx: 438}, '★': {Wx: 816}, '☎': {Wx: 719}, '☛': {Wx: 960}, '☞': {Wx: 939}, '♠': {Wx: 626}, '♣': {Wx: 776}, '♥': {Wx: 694}, '♦': {Wx: 595}, '✁': {Wx: 974}, '✂': {Wx: 961}, '✃': {Wx: 974}, '✄': {Wx: 980}, '✆': {Wx: 789}, '✇': {Wx: 790}, '✈': {Wx: 791}, '✉': {Wx: 690}, '✌': {Wx: 549}, '✍': {Wx: 855}, '✎': {Wx: 911}, '✏': {Wx: 933}, '✐': {Wx: 911}, '✑': {Wx: 945}, '✒': {Wx: 974}, '✓': {Wx: 755}, '✔': {Wx: 846}, '✕': {Wx: 762}, '✖': {Wx: 761}, '✗': {Wx: 571}, '✘': {Wx: 677}, '✙': {Wx: 763}, '✚': {Wx: 760}, '✛': {Wx: 759}, '✜': {Wx: 754}, '✝': {Wx: 494}, '✞': {Wx: 552}, '✟': {Wx: 537}, '✠': {Wx: 577}, '✡': {Wx: 692}, '✢': {Wx: 786}, '✣': {Wx: 788}, '✤': {Wx: 788}, '✥': {Wx: 790}, '✦': {Wx: 793}, '✧': {Wx: 794}, '✩': {Wx: 823}, '✪': {Wx: 789}, '✫': {Wx: 841}, '✬': {Wx: 823}, '✭': {Wx: 833}, '✮': {Wx: 816}, '✯': {Wx: 831}, '✰': {Wx: 923}, '✱': {Wx: 744}, '✲': {Wx: 723}, '✳': {Wx: 749}, '✴': {Wx: 790}, '✵': {Wx: 792}, '✶': {Wx: 695}, '✷': {Wx: 776}, '✸': {Wx: 768}, '✹': {Wx: 792}, '✺': {Wx: 759}, '✻': {Wx: 707}, '✼': {Wx: 708}, '✽': {Wx: 682}, '✾': {Wx: 701}, '✿': {Wx: 826}, '❀': {Wx: 815}, '❁': {Wx: 789}, '❂': {Wx: 789}, '❃': {Wx: 707}, '❄': {Wx: 687}, '❅': {Wx: 696}, '❆': {Wx: 689}, '❇': {Wx: 786}, '❈': {Wx: 787}, '❉': {Wx: 713}, '❊': {Wx: 791}, '❋': {Wx: 785}, '❍': {Wx: 873}, '❏': {Wx: 762}, '❐': {Wx: 762}, '❑': {Wx: 759}, '❒': {Wx: 759}, '❖': {Wx: 784}, '❘': {Wx: 138}, '❙': {Wx: 277}, '❚': {Wx: 415}, '❛': {Wx: 392}, '❜': {Wx: 392}, '❝': {Wx: 668}, '❞': {Wx: 668}, '❡': {Wx: 732}, '❢': {Wx: 544}, '❣': {Wx: 544}, '❤': {Wx: 910}, '❥': {Wx: 667}, '❦': {Wx: 760}, '❧': {Wx: 760}, '❶': {Wx: 788}, '❷': {Wx: 788}, '❸': {Wx: 788}, '❹': {Wx: 788}, '❺': {Wx: 788}, '❻': {Wx: 788}, '❼': {Wx: 788}, '❽': {Wx: 788}, '❾': {Wx: 788}, '❿': {Wx: 788}, '➀': {Wx: 788}, '➁': {Wx: 788}, '➂': {Wx: 788}, '➃': {Wx: 788}, '➄': {Wx: 788}, '➅': {Wx: 788}, '➆': {Wx: 788}, '➇': {Wx: 788}, '➈': {Wx: 788}, '➉': {Wx: 788}, '➊': {Wx: 788}, '➋': {Wx: 788}, '➌': {Wx: 788}, '➍': {Wx: 788}, '➎': {Wx: 788}, '➏': {Wx: 788}, '➐': {Wx: 788}, '➑': {Wx: 788}, '➒': {Wx: 788}, '➓': {Wx: 788}, '➔': {Wx: 894}, '➘': {Wx: 748}, '➙': {Wx: 924}, '➚': {Wx: 748}, '➛': {Wx: 918}, '➜': {Wx: 927}, '➝': {Wx: 928}, '➞': {Wx: 928}, '➟': {Wx: 834}, '➠': {Wx: 873}, '➡': {Wx: 828}, '➢': {Wx: 924}, '➣': {Wx: 924}, '➤': {Wx: 917}, '➥': {Wx: 930}, '➦': {Wx: 931}, '➧': {Wx: 463}, '➨': {Wx: 883}, '➩': {Wx: 836}, '➪': {Wx: 836}, '➫': {Wx: 867}, '➬': {Wx: 867}, '➭': {Wx: 696}, '➮': {Wx: 696}, '➯': {Wx: 874}, '➱': {Wx: 874}, '➲': {Wx: 760}, '➳': {Wx: 946}, '➴': {Wx: 771}, '➵': {Wx: 865}, '➶': {Wx: 771}, '➷': {Wx: 888}, '➸': {Wx: 967}, '➹': {Wx: 888}, '➺': {Wx: 831}, '➻': {Wx: 873}, '➼': {Wx: 927}, '➽': {Wx: 970}, '➾': {Wx: 918}, '\uf8d7': {Wx: 390}, '\uf8d8': {Wx: 390}, '\uf8d9': {Wx: 317}, '\uf8da': {Wx: 317}, '\uf8db': {Wx: 276}, '\uf8dc': {Wx: 276}, '\uf8dd': {Wx: 509}, '\uf8de': {Wx: 509}, '\uf8df': {Wx: 410}, '\uf8e0': {Wx: 410}, '\uf8e1': {Wx: 234}, '\uf8e2': {Wx: 234}, '\uf8e3': {Wx: 334}, '\uf8e4': {Wx: 334}}}
var _ffe = []GlyphName{"\u002en\u006f\u0074\u0064\u0065\u0066", "\u002e\u006e\u0075l\u006c", "\u006e\u006fn\u006d\u0061\u0072k\u0069\u006e\u0067\u0072\u0065\u0074\u0075\u0072\u006e", "\u0073\u0070\u0061c\u0065", "\u0065\u0078\u0063\u006c\u0061\u006d", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c", "\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u0069\u0067\u006e", "\u0064\u006f\u006c\u006c\u0061\u0072", "\u0070e\u0072\u0063\u0065\u006e\u0074", "\u0061m\u0070\u0065\u0072\u0073\u0061\u006ed", "q\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067\u006c\u0065", "\u0070a\u0072\u0065\u006e\u006c\u0065\u0066t", "\u0070\u0061\u0072\u0065\u006e\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0074\u0065\u0072\u0069\u0073\u006b", "\u0070\u006c\u0075\u0073", "\u0063\u006f\u006dm\u0061", "\u0068\u0079\u0070\u0068\u0065\u006e", "\u0070\u0065\u0072\u0069\u006f\u0064", "\u0073\u006c\u0061s\u0068", "\u007a\u0065\u0072\u006f", "\u006f\u006e\u0065", "\u0074\u0077\u006f", "\u0074\u0068\u0072e\u0065", "\u0066\u006f\u0075\u0072", "\u0066\u0069\u0076\u0065", "\u0073\u0069\u0078", "\u0073\u0065\u0076e\u006e", "\u0065\u0069\u0067h\u0074", "\u006e\u0069\u006e\u0065", "\u0063\u006f\u006co\u006e", "\u0073e\u006d\u0069\u0063\u006f\u006c\u006fn", "\u006c\u0065\u0073\u0073", "\u0065\u0071\u0075a\u006c", "\u0067r\u0065\u0061\u0074\u0065\u0072", "\u0071\u0075\u0065\u0073\u0074\u0069\u006f\u006e", "\u0061\u0074", "\u0041", "\u0042", "\u0043", "\u0044", "\u0045", "\u0046", "\u0047", "\u0048", "\u0049", "\u004a", "\u004b", "\u004c", "\u004d", "\u004e", "\u004f", "\u0050", "\u0051", "\u0052", "\u0053", "\u0054", "\u0055", "\u0056", "\u0057", "\u0058", "\u0059", "\u005a", "b\u0072\u0061\u0063\u006b\u0065\u0074\u006c\u0065\u0066\u0074", "\u0062a\u0063\u006b\u0073\u006c\u0061\u0073h", "\u0062\u0072\u0061c\u006b\u0065\u0074\u0072\u0069\u0067\u0068\u0074", "a\u0073\u0063\u0069\u0069\u0063\u0069\u0072\u0063\u0075\u006d", "\u0075\u006e\u0064\u0065\u0072\u0073\u0063\u006f\u0072\u0065", "\u0067\u0072\u0061v\u0065", "\u0061", "\u0062", "\u0063", "\u0064", "\u0065", "\u0066", "\u0067", "\u0068", "\u0069", "\u006a", "\u006b", "\u006c", "\u006d", "\u006e", "\u006f", "\u0070", "\u0071", "\u0072", "\u0073", "\u0074", "\u0075", "\u0076", "\u0077", "\u0078", "\u0079", "\u007a", "\u0062r\u0061\u0063\u0065\u006c\u0065\u0066t", "\u0062\u0061\u0072", "\u0062\u0072\u0061\u0063\u0065\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0063\u0069\u0069\u0074\u0069\u006c\u0064\u0065", "\u0041d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0041\u0072\u0069n\u0067", "\u0043\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0045\u0061\u0063\u0075\u0074\u0065", "\u004e\u0074\u0069\u006c\u0064\u0065", "\u004fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0055d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0061\u0063\u0075\u0074\u0065", "\u0061\u0067\u0072\u0061\u0076\u0065", "a\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0074\u0069\u006c\u0064\u0065", "\u0061\u0072\u0069n\u0067", "\u0063\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0065\u0061\u0063\u0075\u0074\u0065", "\u0065\u0067\u0072\u0061\u0076\u0065", "e\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0065d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0069\u0061\u0063\u0075\u0074\u0065", "\u0069\u0067\u0072\u0061\u0076\u0065", "i\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0069d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006e\u0074\u0069\u006c\u0064\u0065", "\u006f\u0061\u0063\u0075\u0074\u0065", "\u006f\u0067\u0072\u0061\u0076\u0065", "o\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u006fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006f\u0074\u0069\u006c\u0064\u0065", "\u0075\u0061\u0063\u0075\u0074\u0065", "\u0075\u0067\u0072\u0061\u0076\u0065", "u\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0075d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0064\u0061\u0067\u0067\u0065\u0072", "\u0064\u0065\u0067\u0072\u0065\u0065", "\u0063\u0065\u006e\u0074", "\u0073\u0074\u0065\u0072\u006c\u0069\u006e\u0067", "\u0073e\u0063\u0074\u0069\u006f\u006e", "\u0062\u0075\u006c\u006c\u0065\u0074", "\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h", "\u0067\u0065\u0072\u006d\u0061\u006e\u0064\u0062\u006c\u0073", "\u0072\u0065\u0067\u0069\u0073\u0074\u0065\u0072\u0065\u0064", "\u0063o\u0070\u0079\u0072\u0069\u0067\u0068t", "\u0074r\u0061\u0064\u0065\u006d\u0061\u0072k", "\u0061\u0063\u0075t\u0065", "\u0064\u0069\u0065\u0072\u0065\u0073\u0069\u0073", "\u006e\u006f\u0074\u0065\u0071\u0075\u0061\u006c", "\u0041\u0045", "\u004f\u0073\u006c\u0061\u0073\u0068", "\u0069\u006e\u0066\u0069\u006e\u0069\u0074\u0079", "\u0070l\u0075\u0073\u006d\u0069\u006e\u0075s", "\u006ce\u0073\u0073\u0065\u0071\u0075\u0061l", "\u0067\u0072\u0065a\u0074\u0065\u0072\u0065\u0071\u0075\u0061\u006c", "\u0079\u0065\u006e", "\u006d\u0075", "p\u0061\u0072\u0074\u0069\u0061\u006c\u0064\u0069\u0066\u0066", "\u0073u\u006d\u006d\u0061\u0074\u0069\u006fn", "\u0070r\u006f\u0064\u0075\u0063\u0074", "\u0070\u0069", "\u0069\u006e\u0074\u0065\u0067\u0072\u0061\u006c", "o\u0072\u0064\u0066\u0065\u006d\u0069\u006e\u0069\u006e\u0065", "\u006f\u0072\u0064m\u0061\u0073\u0063\u0075\u006c\u0069\u006e\u0065", "\u004f\u006d\u0065g\u0061", "\u0061\u0065", "\u006f\u0073\u006c\u0061\u0073\u0068", "\u0071\u0075\u0065s\u0074\u0069\u006f\u006e\u0064\u006f\u0077\u006e", "\u0065\u0078\u0063\u006c\u0061\u006d\u0064\u006f\u0077\u006e", "\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u006e\u006f\u0074", "\u0072a\u0064\u0069\u0063\u0061\u006c", "\u0066\u006c\u006f\u0072\u0069\u006e", "a\u0070\u0070\u0072\u006f\u0078\u0065\u0071\u0075\u0061\u006c", "\u0044\u0065\u006ct\u0061", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074r\u0069\u0067\u0068\u0074", "\u0065\u006c\u006c\u0069\u0070\u0073\u0069\u0073", "\u006e\u006fn\u0062\u0072\u0065a\u006b\u0069\u006e\u0067\u0073\u0070\u0061\u0063\u0065", "\u0041\u0067\u0072\u0061\u0076\u0065", "\u0041\u0074\u0069\u006c\u0064\u0065", "\u004f\u0074\u0069\u006c\u0064\u0065", "\u004f\u0045", "\u006f\u0065", "\u0065\u006e\u0064\u0061\u0073\u0068", "\u0065\u006d\u0064\u0061\u0073\u0068", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u006c\u0065\u0066\u0074", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c\u0072\u0069\u0067\u0068\u0074", "\u0071u\u006f\u0074\u0065\u006c\u0065\u0066t", "\u0071\u0075\u006f\u0074\u0065\u0072\u0069\u0067\u0068\u0074", "\u0064\u0069\u0076\u0069\u0064\u0065", "\u006co\u007a\u0065\u006e\u0067\u0065", "\u0079d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0059d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e", "\u0063\u0075\u0072\u0072\u0065\u006e\u0063\u0079", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006c\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006cr\u0069\u0067\u0068\u0074", "\u0066\u0069", "\u0066\u006c", "\u0064a\u0067\u0067\u0065\u0072\u0064\u0062l", "\u0070\u0065\u0072\u0069\u006f\u0064\u0063\u0065\u006et\u0065\u0072\u0065\u0064", "\u0071\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067l\u0062\u0061\u0073\u0065", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u0062\u0061\u0073\u0065", "p\u0065\u0072\u0074\u0068\u006f\u0075\u0073\u0061\u006e\u0064", "A\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "E\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0041\u0061\u0063\u0075\u0074\u0065", "\u0045d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0045\u0067\u0072\u0061\u0076\u0065", "\u0049\u0061\u0063\u0075\u0074\u0065", "I\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0049d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0049\u0067\u0072\u0061\u0076\u0065", "\u004f\u0061\u0063\u0075\u0074\u0065", "O\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061\u0070\u0070l\u0065", "\u004f\u0067\u0072\u0061\u0076\u0065", "\u0055\u0061\u0063\u0075\u0074\u0065", "U\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0055\u0067\u0072\u0061\u0076\u0065", "\u0064\u006f\u0074\u006c\u0065\u0073\u0073\u0069", "\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0074\u0069\u006cd\u0065", "\u006d\u0061\u0063\u0072\u006f\u006e", "\u0062\u0072\u0065v\u0065", "\u0064o\u0074\u0061\u0063\u0063\u0065\u006et", "\u0072\u0069\u006e\u0067", "\u0063e\u0064\u0069\u006c\u006c\u0061", "\u0068\u0075\u006eg\u0061\u0072\u0075\u006d\u006c\u0061\u0075\u0074", "\u006f\u0067\u006f\u006e\u0065\u006b", "\u0063\u0061\u0072o\u006e", "\u004c\u0073\u006c\u0061\u0073\u0068", "\u006c\u0073\u006c\u0061\u0073\u0068", "\u0053\u0063\u0061\u0072\u006f\u006e", "\u0073\u0063\u0061\u0072\u006f\u006e", "\u005a\u0063\u0061\u0072\u006f\u006e", "\u007a\u0063\u0061\u0072\u006f\u006e", "\u0062r\u006f\u006b\u0065\u006e\u0062\u0061r", "\u0045\u0074\u0068", "\u0065\u0074\u0068", "\u0059\u0061\u0063\u0075\u0074\u0065", "\u0079\u0061\u0063\u0075\u0074\u0065", "\u0054\u0068\u006fr\u006e", "\u0074\u0068\u006fr\u006e", "\u006d\u0069\u006eu\u0073", "\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "o\u006e\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "t\u0077\u006f\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u0074\u0068\u0072\u0065\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u006fn\u0065\u0068\u0061\u006c\u0066", "\u006f\u006e\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072", "\u0074\u0068\u0072\u0065\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072\u0073", "\u0066\u0072\u0061n\u0063", "\u0047\u0062\u0072\u0065\u0076\u0065", "\u0067\u0062\u0072\u0065\u0076\u0065", "\u0049\u0064\u006f\u0074\u0061\u0063\u0063\u0065\u006e\u0074", "\u0053\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0073\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0043\u0061\u0063\u0075\u0074\u0065", "\u0063\u0061\u0063\u0075\u0074\u0065", "\u0043\u0063\u0061\u0072\u006f\u006e", "\u0063\u0063\u0061\u0072\u006f\u006e", "\u0064\u0063\u0072\u006f\u0061\u0074"}

func (_bagf *ttfParser) ReadULong() (_dbf uint32) {
	_fg.Read(_bagf._cdb, _fg.BigEndian, &_dbf)
	return _dbf
}

var _ggd *RuneCharSafeMap

type Font interface {
	Encoder() _dg.TextEncoder
	GetRuneMetrics(_bf rune) (CharMetrics, bool)
}

func (_ebbf *TtfType) String() string {
	return _eg.Sprintf("\u0046\u004fN\u0054\u005f\u0046\u0049\u004cE\u0032\u007b\u0025\u0023\u0071 \u0055\u006e\u0069\u0074\u0073\u0050\u0065\u0072\u0045\u006d\u003d\u0025\u0064\u0020\u0042\u006f\u006c\u0064\u003d\u0025\u0074\u0020\u0049\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065\u003d\u0025\u0066\u0020"+"\u0043\u0061pH\u0065\u0069\u0067h\u0074\u003d\u0025\u0064 Ch\u0061rs\u003d\u0025\u0064\u0020\u0047\u006c\u0079ph\u004e\u0061\u006d\u0065\u0073\u003d\u0025d\u007d", _ebbf.PostScriptName, _ebbf.UnitsPerEm, _ebbf.Bold, _ebbf.ItalicAngle, _ebbf.CapHeight, len(_ebbf.Chars), len(_ebbf.GlyphNames))
}
func _dea() {
	_gcaa = MakeRuneCharSafeMap(len(_bec))
	_ebg = MakeRuneCharSafeMap(len(_bec))
	for _fcd, _gg := range _bec {
		_gcaa.Write(_gg, CharMetrics{Wx: float64(_gcg[_fcd])})
		_ebg.Write(_gg, CharMetrics{Wx: float64(_aga[_fcd])})
	}
	_ebb = _gcaa.Copy()
	_fff = _ebg.Copy()
}

type CharMetrics struct {
	Wx float64
	Wy float64
}

func _efc() StdFont {
	_cbdc.Do(_dcaf)
	_edb := Descriptor{Name: TimesBoldItalicName, Family: _ebe, Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-200, -218, 996, 921}, ItalicAngle: -15, Ascent: 683, Descent: -217, CapHeight: 669, XHeight: 462, StemV: 121, StemH: 42}
	return NewStdFont(_edb, _ggd)
}
func (_bfb *ttfParser) parseCmapSubtable31(_fbd int64) error {
	_gagg := make([]rune, 0, 8)
	_dada := make([]rune, 0, 8)
	_eff := make([]int16, 0, 8)
	_faca := make([]uint16, 0, 8)
	_bfb._fce.Chars = make(map[rune]GID)
	_bfb._cdb.Seek(int64(_bfb._eef["\u0063\u006d\u0061\u0070"])+_fbd, _dc.SeekStart)
	_debd := _bfb.ReadUShort()
	if _debd != 4 {
		_ae.Log.Debug("u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0073\u0075\u0062t\u0061\u0062\u006c\u0065\u0020\u0066\u006fr\u006d\u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0028\u0025w\u0029", _debd)
		return nil
	}
	_bfb.Skip(2 * 2)
	_eca := int(_bfb.ReadUShort() / 2)
	_bfb.Skip(3 * 2)
	for _abd := 0; _abd < _eca; _abd++ {
		_dada = append(_dada, rune(_bfb.ReadUShort()))
	}
	_bfb.Skip(2)
	for _gea := 0; _gea < _eca; _gea++ {
		_gagg = append(_gagg, rune(_bfb.ReadUShort()))
	}
	for _feeb := 0; _feeb < _eca; _feeb++ {
		_eff = append(_eff, _bfb.ReadShort())
	}
	_eafe, _ := _bfb._cdb.Seek(int64(0), _dc.SeekCurrent)
	for _gaf := 0; _gaf < _eca; _gaf++ {
		_faca = append(_faca, _bfb.ReadUShort())
	}
	for _bfa := 0; _bfa < _eca; _bfa++ {
		_bbfb := _gagg[_bfa]
		_dffc := _dada[_bfa]
		_eag := _eff[_bfa]
		_ebgd := _faca[_bfa]
		if _ebgd > 0 {
			_bfb._cdb.Seek(_eafe+2*int64(_bfa)+int64(_ebgd), _dc.SeekStart)
		}
		for _gce := _bbfb; _gce <= _dffc; _gce++ {
			if _gce == 0xFFFF {
				break
			}
			var _acf int32
			if _ebgd > 0 {
				_acf = int32(_bfb.ReadUShort())
				if _acf > 0 {
					_acf += int32(_eag)
				}
			} else {
				_acf = _gce + int32(_eag)
			}
			if _acf >= 65536 {
				_acf -= 65536
			}
			if _acf > 0 {
				_bfb._fce.Chars[_gce] = GID(_acf)
			}
		}
	}
	return nil
}

var _ebb *RuneCharSafeMap

func init() {
	RegisterStdFont(SymbolName, _eec, "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063", "S\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064", "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(ZapfDingbatsName, _gd)
}

var _ed = &fontMap{_gca: make(map[StdFontName]func() StdFont)}

func (_aeb *fontMap) read(_dgb StdFontName) (func() StdFont, bool) {
	_aeb.Lock()
	defer _aeb.Unlock()
	_bd, _fd := _aeb._gca[_dgb]
	return _bd, _fd
}
func _dcaf() {
	_ecg = MakeRuneCharSafeMap(len(_bec))
	_gcgg = MakeRuneCharSafeMap(len(_bec))
	_ggd = MakeRuneCharSafeMap(len(_bec))
	_ddc = MakeRuneCharSafeMap(len(_bec))
	for _beb, _eeec := range _bec {
		_ecg.Write(_eeec, CharMetrics{Wx: float64(_eba[_beb])})
		_gcgg.Write(_eeec, CharMetrics{Wx: float64(_gee[_beb])})
		_ggd.Write(_eeec, CharMetrics{Wx: float64(_geca[_beb])})
		_ddc.Write(_eeec, CharMetrics{Wx: float64(_ce[_beb])})
	}
}
func NewFontFile2FromPdfObject(obj _gc.PdfObject) (TtfType, error) {
	obj = _gc.TraceToDirectObject(obj)
	_cgd, _eefc := obj.(*_gc.PdfObjectStream)
	if !_eefc {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0073\u0074\u0072e\u0061\u006d \u0028\u0025\u0054\u0029", obj)
		return TtfType{}, _gc.ErrTypeError
	}
	_eea, _dbe := _gc.DecodeStream(_cgd)
	if _dbe != nil {
		return TtfType{}, _dbe
	}
	_bbfec := ttfParser{_cdb: _af.NewReader(_eea)}
	return _bbfec.Parse()
}
func (_feac *TtfType) NewEncoder() _dg.TextEncoder { return _dg.NewTrueTypeFontEncoder(_feac.Chars) }
func init() {
	RegisterStdFont(TimesRomanName, _ebd, "\u0054\u0069\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e", "\u0054\u0069\u006de\u0073")
	RegisterStdFont(TimesBoldName, _gcc, "\u0054i\u006de\u0073\u004e\u0065\u0077\u0052o\u006d\u0061n\u002c\u0042\u006f\u006c\u0064", "\u0054\u0069\u006d\u0065\u0073\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(TimesItalicName, _bdba, "T\u0069m\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006da\u006e\u002c\u0049\u0074al\u0069\u0063", "\u0054\u0069\u006de\u0073\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(TimesBoldItalicName, _efc, "\u0054i\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e,\u0042\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063", "\u0054\u0069m\u0065\u0073\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}

const (
	CourierName            = StdFontName("\u0043o\u0075\u0072\u0069\u0065\u0072")
	CourierBoldName        = StdFontName("\u0043\u006f\u0075r\u0069\u0065\u0072\u002d\u0042\u006f\u006c\u0064")
	CourierObliqueName     = StdFontName("\u0043o\u0075r\u0069\u0065\u0072\u002d\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
	CourierBoldObliqueName = StdFontName("\u0043\u006f\u0075\u0072ie\u0072\u002d\u0042\u006f\u006c\u0064\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
)

func _cbd() StdFont {
	_bbf.Do(_gaa)
	_cag := Descriptor{Name: CourierBoldName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0021, BBox: [4]float64{-113, -250, 749, 801}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_cag, _fa)
}
func NewStdFont(desc Descriptor, metrics *RuneCharSafeMap) StdFont {
	return NewStdFontWithEncoding(desc, metrics, _dg.NewStandardEncoder())
}

type GID = _dg.GID

var _ Font = StdFont{}
var _gag *RuneCharSafeMap

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

var _gee = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 611, 778, 778, 778, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 778, 778, 667, 667, 667, 667, 667, 944, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 611, 778, 722, 722, 722, 722, 556, 556, 556, 556, 556, 667, 667, 667, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 1000, 722, 722, 722, 722, 667, 667, 667, 667, 500, 500, 500, 500, 333, 500, 722, 500, 500, 833, 500, 500, 581, 520, 500, 930, 500, 556, 278, 220, 394, 394, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 556, 500, 500, 672, 556, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 556, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 333, 556, 556, 278, 278, 394, 278, 570, 549, 570, 494, 278, 833, 333, 570, 556, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 300, 330, 500, 500, 556, 540, 333, 333, 494, 1000, 250, 250, 1000, 570, 570, 556, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 444, 444, 549, 444, 444, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 333, 416, 333, 556, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func _dff() StdFont {
	_bbf.Do(_gaa)
	_bgg := Descriptor{Name: CourierObliqueName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0061, BBox: [4]float64{-27, -250, 849, 805}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_bgg, _eb)
}
func (_cbb *ttfParser) readByte() (_bbfa uint8) {
	_fg.Read(_cbb._cdb, _fg.BigEndian, &_bbfa)
	return _bbfa
}
func RegisterStdFont(name StdFontName, fnc func() StdFont, aliases ...StdFontName) {
	if _, _bdd := _ed.read(name); _bdd {
		panic("\u0066o\u006e\u0074\u0020\u0061l\u0072\u0065\u0061\u0064\u0079 \u0072e\u0067i\u0073\u0074\u0065\u0072\u0065\u0064\u003a " + string(name))
	}
	_ed.write(name, fnc)
	for _, _bge := range aliases {
		RegisterStdFont(_bge, fnc)
	}
}
func (_afe StdFont) Descriptor() Descriptor { return _afe._be }
func (_bgge *ttfParser) Parse() (TtfType, error) {
	_eed, _ccb := _bgge.ReadStr(4)
	if _ccb != nil {
		return TtfType{}, _ccb
	}
	if _eed == "\u0074\u0074\u0063\u0066" {
		return _bgge.parseTTC()
	} else if _eed != "\u0000\u0001\u0000\u0000" && _eed != "\u0074\u0072\u0075\u0065" {
		_ae.Log.Debug("\u0055n\u0072\u0065c\u006f\u0067\u006ei\u007a\u0065\u0064\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066i\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u002e\u0020v\u0065\u0072\u0073\u0069\u006f\u006e\u003d\u0025\u0071", _eed)
	}
	_adcc := int(_bgge.ReadUShort())
	_bgge.Skip(3 * 2)
	_bgge._eef = make(map[string]uint32)
	var _gcaf string
	for _aa := 0; _aa < _adcc; _aa++ {
		_gcaf, _ccb = _bgge.ReadStr(4)
		if _ccb != nil {
			return TtfType{}, _ccb
		}
		_bgge.Skip(4)
		_dgd := _bgge.ReadULong()
		_bgge.Skip(4)
		_bgge._eef[_gcaf] = _dgd
	}
	_ae.Log.Trace(_dag(_bgge._eef))
	if _ccb = _bgge.ParseComponents(); _ccb != nil {
		return TtfType{}, _ccb
	}
	return _bgge._fce, nil
}
func (_dgg CharMetrics) String() string {
	return _eg.Sprintf("<\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u003e", _dgg.Wx, _dgg.Wy)
}
func (_bg *RuneCharSafeMap) Length() int {
	_bg._fc.RLock()
	defer _bg._fc.RUnlock()
	return len(_bg._ee)
}
func (_fceg *ttfParser) ParsePost() error {
	if _fcdg := _fceg.Seek("\u0070\u006f\u0073\u0074"); _fcdg != nil {
		return _fcdg
	}
	_becd := _fceg.Read32Fixed()
	_fceg._fce.ItalicAngle = _fceg.Read32Fixed()
	_fceg._fce.UnderlinePosition = _fceg.ReadShort()
	_fceg._fce.UnderlineThickness = _fceg.ReadShort()
	_fceg._fce.IsFixedPitch = _fceg.ReadULong() != 0
	_fceg.ReadULong()
	_fceg.ReadULong()
	_fceg.ReadULong()
	_fceg.ReadULong()
	_ae.Log.Trace("\u0050a\u0072\u0073\u0065\u0050\u006f\u0073\u0074\u003a\u0020\u0066\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065\u003d\u0025\u0066", _becd)
	switch _becd {
	case 1.0:
		_fceg._fce.GlyphNames = _ffe
	case 2.0:
		_gge := int(_fceg.ReadUShort())
		_eeeaa := make([]int, _gge)
		_fceg._fce.GlyphNames = make([]GlyphName, _gge)
		_debdc := -1
		for _bee := 0; _bee < _gge; _bee++ {
			_afa := int(_fceg.ReadUShort())
			_eeeaa[_bee] = _afa
			if _afa <= 0x7fff && _afa > _debdc {
				_debdc = _afa
			}
		}
		var _geee []GlyphName
		if _debdc >= len(_ffe) {
			_geee = make([]GlyphName, _debdc-len(_ffe)+1)
			for _agd := 0; _agd < _debdc-len(_ffe)+1; _agd++ {
				_edcb := int(_fceg.readByte())
				_efgb, _aag := _fceg.ReadStr(_edcb)
				if _aag != nil {
					return _aag
				}
				_geee[_agd] = GlyphName(_efgb)
			}
		}
		for _faf := 0; _faf < _gge; _faf++ {
			_acagb := _eeeaa[_faf]
			if _acagb < len(_ffe) {
				_fceg._fce.GlyphNames[_faf] = _ffe[_acagb]
			} else if _acagb >= len(_ffe) && _acagb <= 32767 {
				_fceg._fce.GlyphNames[_faf] = _geee[_acagb-len(_ffe)]
			} else {
				_fceg._fce.GlyphNames[_faf] = "\u002e\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064"
			}
		}
	case 2.5:
		_egdc := make([]int, _fceg._gac)
		for _cba := 0; _cba < len(_egdc); _cba++ {
			_aba := int(_fceg.ReadSByte())
			_egdc[_cba] = _cba + 1 + _aba
		}
		_fceg._fce.GlyphNames = make([]GlyphName, len(_egdc))
		for _edg := 0; _edg < len(_fceg._fce.GlyphNames); _edg++ {
			_cbf := _ffe[_egdc[_edg]]
			_fceg._fce.GlyphNames[_edg] = _cbf
		}
	case 3.0:
		_ae.Log.Debug("\u004e\u006f\u0020\u0050\u006f\u0073t\u0053\u0063\u0072i\u0070\u0074\u0020n\u0061\u006d\u0065\u0020\u0069\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020is\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e")
	default:
		_ae.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020f\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065=\u0025\u0066", _becd)
	}
	return nil
}
func (_ffb *ttfParser) ParseOS2() error {
	if _abcb := _ffb.Seek("\u004f\u0053\u002f\u0032"); _abcb != nil {
		return _abcb
	}
	_dcb := _ffb.ReadUShort()
	_ffb.Skip(4 * 2)
	_ffb.Skip(11*2 + 10 + 4*4 + 4)
	_baa := _ffb.ReadUShort()
	_ffb._fce.Bold = (_baa & 32) != 0
	_ffb.Skip(2 * 2)
	_ffb._fce.TypoAscender = _ffb.ReadShort()
	_ffb._fce.TypoDescender = _ffb.ReadShort()
	if _dcb >= 2 {
		_ffb.Skip(3*2 + 2*4 + 2)
		_ffb._fce.CapHeight = _ffb.ReadShort()
	} else {
		_ffb._fce.CapHeight = 0
	}
	return nil
}

var _geca = []int16{667, 944, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 667, 722, 722, 722, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 667, 667, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 889, 667, 611, 611, 611, 611, 611, 611, 611, 500, 500, 500, 500, 333, 500, 722, 500, 500, 778, 500, 500, 570, 570, 500, 832, 500, 500, 278, 220, 348, 348, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 500, 500, 500, 608, 500, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 389, 389, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 382, 278, 570, 549, 606, 494, 278, 778, 333, 606, 576, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 266, 300, 500, 500, 500, 500, 333, 333, 494, 833, 250, 250, 1000, 570, 570, 500, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 389, 389, 549, 389, 389, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 366, 278, 500, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 444, 667, 500, 444, 444, 444, 500, 389, 389, 389, 389, 500}
var _eb *RuneCharSafeMap
var _fcg _ad.Once

func (_bea *ttfParser) Read32Fixed() float64 {
	_ege := float64(_bea.ReadShort())
	_afeb := float64(_bea.ReadUShort()) / 65536.0
	return _ege + _afeb
}
func (_cce *ttfParser) parseCmapSubtable10(_aaa int64) error {
	if _cce._fce.Chars == nil {
		_cce._fce.Chars = make(map[rune]GID)
	}
	_cce._cdb.Seek(int64(_cce._eef["\u0063\u006d\u0061\u0070"])+_aaa, _dc.SeekStart)
	var _edc, _bgf uint32
	_edbb := _cce.ReadUShort()
	if _edbb < 8 {
		_edc = uint32(_cce.ReadUShort())
		_bgf = uint32(_cce.ReadUShort())
	} else {
		_cce.ReadUShort()
		_edc = _cce.ReadULong()
		_bgf = _cce.ReadULong()
	}
	_ae.Log.Trace("\u0070\u0061r\u0073\u0065\u0043\u006d\u0061p\u0053\u0075\u0062\u0074\u0061b\u006c\u0065\u0031\u0030\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065\u003d\u0025\u0064", _edbb, _edc, _bgf)
	if _edbb != 0 {
		return _dd.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006d\u0061p\u0020s\u0075\u0062\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	_dbc, _eab := _cce.ReadStr(256)
	if _eab != nil {
		return _eab
	}
	_fage := []byte(_dbc)
	for _bbd, _ffcc := range _fage {
		_cce._fce.Chars[rune(_bbd)] = GID(_ffcc)
		if _ffcc != 0 {
			_eg.Printf("\u0009\u0030\u0078\u002502\u0078\u0020\u279e\u0020\u0030\u0078\u0025\u0030\u0032\u0078\u003d\u0025\u0063\u000a", _bbd, _ffcc, rune(_ffcc))
		}
	}
	return nil
}

var _bbf _ad.Once

const (
	SymbolName       = StdFontName("\u0053\u0079\u006d\u0062\u006f\u006c")
	ZapfDingbatsName = StdFontName("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
)

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

func _fgc() StdFont {
	_bbf.Do(_gaa)
	_afea := Descriptor{Name: CourierName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0021, BBox: [4]float64{-23, -250, 715, 805}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_afea, _bdb)
}
func (_daf *ttfParser) ParseMaxp() error {
	if _adba := _daf.Seek("\u006d\u0061\u0078\u0070"); _adba != nil {
		return _adba
	}
	_daf.Skip(4)
	_daf._gac = _daf.ReadUShort()
	return nil
}
func _bdba() StdFont {
	_cbdc.Do(_dcaf)
	_fee := Descriptor{Name: TimesItalicName, Family: _ebe, Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-169, -217, 1010, 883}, ItalicAngle: -15.5, Ascent: 683, Descent: -217, CapHeight: 653, XHeight: 441, StemV: 76, StemH: 32}
	return NewStdFont(_fee, _ddc)
}
func (_edbd *ttfParser) parseCmapVersion(_cfe int64) error {
	_ae.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006da\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a \u006f\u0066\u0066s\u0065t\u003d\u0025\u0064", _cfe)
	if _edbd._fce.Chars == nil {
		_edbd._fce.Chars = make(map[rune]GID)
	}
	_edbd._cdb.Seek(int64(_edbd._eef["\u0063\u006d\u0061\u0070"])+_cfe, _dc.SeekStart)
	var _gff, _fga uint32
	_gfb := _edbd.ReadUShort()
	if _gfb < 8 {
		_gff = uint32(_edbd.ReadUShort())
		_fga = uint32(_edbd.ReadUShort())
	} else {
		_edbd.ReadUShort()
		_gff = _edbd.ReadULong()
		_fga = _edbd.ReadULong()
	}
	_ae.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0043m\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064 \u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067u\u0061g\u0065\u003d\u0025\u0064", _gfb, _gff, _fga)
	switch _gfb {
	case 0:
		return _edbd.parseCmapFormat0()
	case 6:
		return _edbd.parseCmapFormat6()
	case 12:
		return _edbd.parseCmapFormat12()
	default:
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063m\u0061\u0070\u0020\u0066\u006f\u0072\u006da\u0074\u003d\u0025\u0064", _gfb)
		return nil
	}
}
func (_dfd *ttfParser) ParseCmap() error {
	var _cda int64
	if _fcc := _dfd.Seek("\u0063\u006d\u0061\u0070"); _fcc != nil {
		return _fcc
	}
	_dfd.ReadUShort()
	_fed := int(_dfd.ReadUShort())
	_bgga := int64(0)
	_ecgg := int64(0)
	_facb := int64(0)
	for _ggg := 0; _ggg < _fed; _ggg++ {
		_bag := _dfd.ReadUShort()
		_gaga := _dfd.ReadUShort()
		_cda = int64(_dfd.ReadULong())
		if _bag == 3 && _gaga == 1 {
			_ecgg = _cda
		} else if _bag == 3 && _gaga == 10 {
			_facb = _cda
		} else if _bag == 1 && _gaga == 0 {
			_bgga = _cda
		}
	}
	if _bgga != 0 {
		if _ggdc := _dfd.parseCmapVersion(_bgga); _ggdc != nil {
			return _ggdc
		}
	}
	if _ecgg != 0 {
		if _fae := _dfd.parseCmapSubtable31(_ecgg); _fae != nil {
			return _fae
		}
	}
	if _facb != 0 {
		if _cgb := _dfd.parseCmapVersion(_facb); _cgb != nil {
			return _cgb
		}
	}
	if _ecgg == 0 && _bgga == 0 && _facb == 0 {
		_ae.Log.Debug("\u0074\u0074\u0066P\u0061\u0072\u0073\u0065\u0072\u002e\u0050\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u002e\u0020\u004e\u006f\u0020\u0033\u0031\u002c\u0020\u0031\u0030\u002c\u0020\u00331\u0030\u0020\u0074\u0061\u0062\u006c\u0065\u002e")
	}
	return nil
}
func (_cee *ttfParser) ParseHmtx() error {
	if _aad := _cee.Seek("\u0068\u006d\u0074\u0078"); _aad != nil {
		return _aad
	}
	_cee._fce.Widths = make([]uint16, 0, 8)
	for _aca := uint16(0); _aca < _cee._adf; _aca++ {
		_cee._fce.Widths = append(_cee._fce.Widths, _cee.ReadUShort())
		_cee.Skip(2)
	}
	if _cee._adf < _cee._gac && _cee._adf > 0 {
		_eeff := _cee._fce.Widths[_cee._adf-1]
		for _eeea := _cee._adf; _eeea < _cee._gac; _eeea++ {
			_cee._fce.Widths = append(_cee._fce.Widths, _eeff)
		}
	}
	return nil
}

const (
	FontWeightMedium FontWeight = iota
	FontWeightBold
	FontWeightRoman
)

func (_cae StdFont) ToPdfObject() _gc.PdfObject {
	_ga := _gc.MakeDict()
	_ga.Set("\u0054\u0079\u0070\u0065", _gc.MakeName("\u0046\u006f\u006e\u0074"))
	_ga.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _gc.MakeName("\u0054\u0079\u0070e\u0031"))
	_ga.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _gc.MakeName(_cae.Name()))
	_ga.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _cae._df.ToPdfObject())
	return _gc.MakeIndirectObject(_ga)
}
func (_adfd *ttfParser) ParseName() error {
	if _dffg := _adfd.Seek("\u006e\u0061\u006d\u0065"); _dffg != nil {
		return _dffg
	}
	_cdba, _ := _adfd._cdb.Seek(0, _dc.SeekCurrent)
	_adfd._fce.PostScriptName = ""
	_adfd.Skip(2)
	_cgc := _adfd.ReadUShort()
	_bca := _adfd.ReadUShort()
	for _gcad := uint16(0); _gcad < _cgc && _adfd._fce.PostScriptName == ""; _gcad++ {
		_adfd.Skip(3 * 2)
		_bae := _adfd.ReadUShort()
		_dee := _adfd.ReadUShort()
		_cef := _adfd.ReadUShort()
		if _bae == 6 {
			_adfd._cdb.Seek(_cdba+int64(_bca)+int64(_cef), _dc.SeekStart)
			_fdg, _afff := _adfd.ReadStr(int(_dee))
			if _afff != nil {
				return _afff
			}
			_fdg = _a.Replace(_fdg, "\u0000", "", -1)
			_acag, _afff := _f.Compile("\u005b\u0028\u0029\u007b\u007d\u003c\u003e\u0020\u002f%\u005b\u005c\u005d\u005d")
			if _afff != nil {
				return _afff
			}
			_adfd._fce.PostScriptName = _acag.ReplaceAllString(_fdg, "")
		}
	}
	if _adfd._fce.PostScriptName == "" {
		_ae.Log.Debug("\u0050a\u0072\u0073e\u004e\u0061\u006de\u003a\u0020\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u0050\u006f\u0073t\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0077\u0061\u0073\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	}
	return nil
}

var _ecg *RuneCharSafeMap

func (_cdf *TtfType) MakeEncoder() (_dg.SimpleEncoder, error) {
	_ffg := make(map[_dg.CharCode]GlyphName)
	for _bdg := _dg.CharCode(0); _bdg <= 256; _bdg++ {
		_bga := rune(_bdg)
		_gfg, _dfb := _cdf.Chars[_bga]
		if !_dfb {
			continue
		}
		var _ggc GlyphName
		if int(_gfg) >= 0 && int(_gfg) < len(_cdf.GlyphNames) {
			_ggc = _cdf.GlyphNames[_gfg]
		} else {
			_gad := rune(_gfg)
			if _ccf, _dfbe := _dg.RuneToGlyph(_gad); _dfbe {
				_ggc = _ccf
			}
		}
		if _ggc != "" {
			_ffg[_bdg] = _ggc
		}
	}
	if len(_ffg) == 0 {
		_ae.Log.Debug("WA\u0052\u004eI\u004e\u0047\u003a\u0020\u005a\u0065\u0072\u006f\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002e\u0020\u0074\u0074\u0066=\u0025s\u0020\u0043\u0068\u0061\u0072\u0073\u003d\u005b%\u00200\u0032\u0078]", _cdf, _cdf.Chars)
	}
	return _dg.NewCustomSimpleTextEncoder(_ffg, nil)
}
func (_c *RuneCharSafeMap) Copy() *RuneCharSafeMap {
	_gb := MakeRuneCharSafeMap(_c.Length())
	_c.Range(func(_gf rune, _cg CharMetrics) (_gba bool) { _gb._ee[_gf] = _cg; return false })
	return _gb
}

var _cbdc _ad.Once

func _gaa() {
	const _egdd = 600
	_bdb = MakeRuneCharSafeMap(len(_bec))
	for _, _eda := range _bec {
		_bdb.Write(_eda, CharMetrics{Wx: _egdd})
	}
	_fa = _bdb.Copy()
	_gag = _bdb.Copy()
	_eb = _bdb.Copy()
}
func IsStdFont(name StdFontName) bool {
	_, _dgf := _ed.read(name)
	return _dgf
}

type ttfParser struct {
	_fce TtfType
	_cdb _dc.ReadSeeker
	_eef map[string]uint32
	_adf uint16
	_gac uint16
}

var _bec = []rune{'A', 'Æ', 'Á', 'Ă', 'Â', 'Ä', 'À', 'Ā', 'Ą', 'Å', 'Ã', 'B', 'C', 'Ć', 'Č', 'Ç', 'D', 'Ď', 'Đ', '∆', 'E', 'É', 'Ě', 'Ê', 'Ë', 'Ė', 'È', 'Ē', 'Ę', 'Ð', '€', 'F', 'G', 'Ğ', 'Ģ', 'H', 'I', 'Í', 'Î', 'Ï', 'İ', 'Ì', 'Ī', 'Į', 'J', 'K', 'Ķ', 'L', 'Ĺ', 'Ľ', 'Ļ', 'Ł', 'M', 'N', 'Ń', 'Ň', 'Ņ', 'Ñ', 'O', 'Œ', 'Ó', 'Ô', 'Ö', 'Ò', 'Ő', 'Ō', 'Ø', 'Õ', 'P', 'Q', 'R', 'Ŕ', 'Ř', 'Ŗ', 'S', 'Ś', 'Š', 'Ş', 'Ș', 'T', 'Ť', 'Ţ', 'Þ', 'U', 'Ú', 'Û', 'Ü', 'Ù', 'Ű', 'Ū', 'Ų', 'Ů', 'V', 'W', 'X', 'Y', 'Ý', 'Ÿ', 'Z', 'Ź', 'Ž', 'Ż', 'a', 'á', 'ă', 'â', '´', 'ä', 'æ', 'à', 'ā', '&', 'ą', 'å', '^', '~', '*', '@', 'ã', 'b', '\\', '|', '{', '}', '[', ']', '˘', '¦', '•', 'c', 'ć', 'ˇ', 'č', 'ç', '¸', '¢', 'ˆ', ':', ',', '\uf6c3', '©', '¤', 'd', '†', '‡', 'ď', 'đ', '°', '¨', '÷', '$', '˙', 'ı', 'e', 'é', 'ě', 'ê', 'ë', 'ė', 'è', '8', '…', 'ē', '—', '–', 'ę', '=', 'ð', '!', '¡', 'f', 'ﬁ', '5', 'ﬂ', 'ƒ', '4', '⁄', 'g', 'ğ', 'ģ', 'ß', '`', '>', '≥', '«', '»', '‹', '›', 'h', '˝', '-', 'i', 'í', 'î', 'ï', 'ì', 'ī', 'į', 'j', 'k', 'ķ', 'l', 'ĺ', 'ľ', 'ļ', '<', '≤', '¬', '◊', 'ł', 'm', '¯', '−', 'µ', '×', 'n', 'ń', 'ň', 'ņ', '9', '≠', 'ñ', '#', 'o', 'ó', 'ô', 'ö', 'œ', '˛', 'ò', 'ő', 'ō', '1', '½', '¼', '¹', 'ª', 'º', 'ø', 'õ', 'p', '¶', '(', ')', '∂', '%', '.', '·', '‰', '+', '±', 'q', '?', '¿', '"', '„', '“', '”', '‘', '’', '‚', '\'', 'r', 'ŕ', '√', 'ř', 'ŗ', '®', '˚', 's', 'ś', 'š', 'ş', 'ș', '§', ';', '7', '6', '/', ' ', '£', '∑', 't', 'ť', 'ţ', 'þ', '3', '¾', '³', '˜', '™', '2', '²', 'u', 'ú', 'û', 'ü', 'ù', 'ű', 'ū', '_', 'ų', 'ů', 'v', 'w', 'x', 'y', 'ý', 'ÿ', '¥', 'z', 'ź', 'ž', 'ż', '0'}

func (_de *fontMap) write(_aec StdFontName, _da func() StdFont) {
	_de.Lock()
	defer _de.Unlock()
	_de._gca[_aec] = _da
}
func (_fgee *ttfParser) parseTTC() (TtfType, error) {
	_fgee.Skip(2 * 2)
	_aff := _fgee.ReadULong()
	if _aff < 1 {
		return TtfType{}, _dd.New("N\u006f \u0066\u006f\u006e\u0074\u0073\u0020\u0069\u006e \u0054\u0054\u0043\u0020fi\u006c\u0065")
	}
	_agc := _fgee.ReadULong()
	_, _feaf := _fgee._cdb.Seek(int64(_agc), _dc.SeekStart)
	if _feaf != nil {
		return TtfType{}, _feaf
	}
	return _fgee.Parse()
}

var _fff *RuneCharSafeMap
var _gcg = []int16{667, 1000, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 500, 667, 667, 556, 556, 556, 556, 556, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 667, 556, 556, 469, 584, 389, 1015, 556, 556, 278, 260, 334, 334, 278, 278, 333, 260, 350, 500, 500, 333, 500, 500, 333, 556, 333, 278, 278, 250, 737, 556, 556, 556, 556, 643, 556, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 556, 278, 333, 278, 500, 556, 500, 556, 556, 167, 556, 556, 556, 611, 333, 584, 549, 556, 556, 333, 333, 556, 333, 333, 222, 278, 278, 278, 278, 278, 222, 222, 500, 500, 222, 222, 299, 222, 584, 549, 584, 471, 222, 833, 333, 584, 556, 584, 556, 556, 556, 556, 556, 549, 556, 556, 556, 556, 556, 556, 944, 333, 556, 556, 556, 556, 834, 834, 333, 370, 365, 611, 556, 556, 537, 333, 333, 476, 889, 278, 278, 1000, 584, 584, 556, 556, 611, 355, 333, 333, 333, 222, 222, 222, 191, 333, 333, 453, 333, 333, 737, 333, 500, 500, 500, 500, 500, 556, 278, 556, 556, 278, 278, 556, 600, 278, 317, 278, 556, 556, 834, 333, 333, 1000, 556, 333, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 500, 722, 500, 500, 500, 500, 556, 500, 500, 500, 500, 556}

func (_bcaf *ttfParser) ReadShort() (_aeca int16) {
	_fg.Read(_bcaf._cdb, _fg.BigEndian, &_aeca)
	return _aeca
}
func (_efbc *ttfParser) ReadSByte() (_dbdg int8) {
	_fg.Read(_efbc._cdb, _fg.BigEndian, &_dbdg)
	return _dbdg
}
func (_cgdd *ttfParser) Skip(n int) { _cgdd._cdb.Seek(int64(n), _dc.SeekCurrent) }

var _fa *RuneCharSafeMap

func (_fcb *TtfType) MakeToUnicode() *_g.CMap {
	_bgd := make(map[_g.CharCode]rune)
	if len(_fcb.GlyphNames) == 0 {
		for _def := range _fcb.Chars {
			_bgd[_g.CharCode(_def)] = _def
		}
		return _g.NewToUnicodeCMap(_bgd)
	}
	for _fbf, _gcf := range _fcb.Chars {
		_agea := _g.CharCode(_fbf)
		_ecf := _fcb.GlyphNames[_gcf]
		_fac, _deg := _dg.GlyphToRune(_ecf)
		if !_deg {
			_ae.Log.Debug("\u004e\u006f \u0072\u0075\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u0020\u0067\u006c\u0079\u0070h=\u0025\u0071", _fbf, _ecf)
			_fac = _dg.MissingCodeRune
		}
		_bgd[_agea] = _fac
	}
	return _g.NewToUnicodeCMap(_bgd)
}

var _ebg *RuneCharSafeMap

func (_gfa *ttfParser) parseCmapFormat0() error {
	_ggdb, _dbd := _gfa.ReadStr(256)
	if _dbd != nil {
		return _dbd
	}
	_efb := []byte(_ggdb)
	_ae.Log.Trace("\u0070a\u0072\u0073e\u0043\u006d\u0061p\u0046\u006f\u0072\u006d\u0061\u0074\u0030:\u0020\u0025\u0073\u000a\u0064\u0061t\u0061\u0053\u0074\u0072\u003d\u0025\u002b\u0071\u000a\u0064\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d", _gfa._fce.String(), _ggdb, _efb)
	for _bfc, _gffa := range _efb {
		_gfa._fce.Chars[rune(_bfc)] = GID(_gffa)
	}
	return nil
}
func (_bcf *ttfParser) ParseHead() error {
	if _gab := _bcf.Seek("\u0068\u0065\u0061\u0064"); _gab != nil {
		return _gab
	}
	_bcf.Skip(3 * 4)
	_eeaa := _bcf.ReadULong()
	if _eeaa != 0x5F0F3CF5 {
		_ae.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0063\u006fr\u0072e\u0063\u0074\u0020\u006d\u0061\u0067\u0069\u0063\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e\u0020\u0046\u006fn\u0074\u0020\u006d\u0061\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0069\u0073\u0070\u006c\u0061\u0079\u0020\u0063\u006f\u0072\u0072\u0065\u0063t\u006c\u0079\u002e\u0020\u0025\u0073", _bcf)
	}
	_bcf.Skip(2)
	_bcf._fce.UnitsPerEm = _bcf.ReadUShort()
	_bcf.Skip(2 * 8)
	_bcf._fce.Xmin = _bcf.ReadShort()
	_bcf._fce.Ymin = _bcf.ReadShort()
	_bcf._fce.Xmax = _bcf.ReadShort()
	_bcf._fce.Ymax = _bcf.ReadShort()
	return nil
}
