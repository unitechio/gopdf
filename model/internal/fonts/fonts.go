package fonts

import (
	_bf "bytes"
	_e "encoding/binary"
	_b "errors"
	_d "fmt"
	_fd "io"
	_f "os"
	_gd "regexp"
	_da "sort"
	_a "strings"
	_eb "sync"

	_fe "unitechio/gopdf/gopdf/common"
	_c "unitechio/gopdf/gopdf/core"
	_be "unitechio/gopdf/gopdf/internal/cmap"
	_bfg "unitechio/gopdf/gopdf/internal/textencoding"
)

func (_de *RuneCharSafeMap) Copy() *RuneCharSafeMap {
	_cbg := MakeRuneCharSafeMap(_de.Length())
	_de.Range(func(_ebc rune, _bb CharMetrics) (_ba bool) { _cbg._cb[_ebc] = _bb; return false })
	return _cbg
}

func _edf() StdFont {
	_fecg := _bfg.NewSymbolEncoder()
	_fee := Descriptor{Name: SymbolName, Family: string(SymbolName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-180, -293, 1090, 1010}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 85, StemH: 92}
	return NewStdFontWithEncoding(_fee, _dbd, _fecg)
}

type ttfParser struct {
	_gbedf TtfType
	_fdbe  _fd.ReadSeeker
	_fddc  map[string]uint32
	_dac   uint16
	_aegc  uint16
}
type StdFont struct {
	_gb Descriptor
	_fa *RuneCharSafeMap
	_fc _bfg.TextEncoder
}

func (_ffc *ttfParser) ParseCmap() error {
	var _bge int64
	if _beac := _ffc.Seek("\u0063\u006d\u0061\u0070"); _beac != nil {
		return _beac
	}
	_ffc.ReadUShort()
	_dcaf := int(_ffc.ReadUShort())
	_cdb := int64(0)
	_eabe := int64(0)
	_cfef := int64(0)
	for _bbba := 0; _bbba < _dcaf; _bbba++ {
		_cdf := _ffc.ReadUShort()
		_ebbf := _ffc.ReadUShort()
		_bge = int64(_ffc.ReadULong())
		if _cdf == 3 && _ebbf == 1 {
			_eabe = _bge
		} else if _cdf == 3 && _ebbf == 10 {
			_cfef = _bge
		} else if _cdf == 1 && _ebbf == 0 {
			_cdb = _bge
		}
	}
	if _cdb != 0 {
		if _fege := _ffc.parseCmapVersion(_cdb); _fege != nil {
			return _fege
		}
	}
	if _eabe != 0 {
		if _ceaf := _ffc.parseCmapSubtable31(_eabe); _ceaf != nil {
			return _ceaf
		}
	}
	if _cfef != 0 {
		if _bdb := _ffc.parseCmapVersion(_cfef); _bdb != nil {
			return _bdb
		}
	}
	if _eabe == 0 && _cdb == 0 && _cfef == 0 {
		_fe.Log.Debug("\u0074\u0074\u0066P\u0061\u0072\u0073\u0065\u0072\u002e\u0050\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u002e\u0020\u004e\u006f\u0020\u0033\u0031\u002c\u0020\u0031\u0030\u002c\u0020\u00331\u0030\u0020\u0074\u0061\u0062\u006c\u0065\u002e")
	}
	return nil
}

var _df *RuneCharSafeMap

func (_cge *ttfParser) ReadULong() (_geca uint32) {
	_e.Read(_cge._fdbe, _e.BigEndian, &_geca)
	return _geca
}

func (_fcc *ttfParser) ParseHmtx() error {
	if _gfcc := _fcc.Seek("\u0068\u006d\u0074\u0078"); _gfcc != nil {
		return _gfcc
	}
	_fcc._gbedf.Widths = make([]uint16, 0, 8)
	for _ffe := uint16(0); _ffe < _fcc._dac; _ffe++ {
		_fcc._gbedf.Widths = append(_fcc._gbedf.Widths, _fcc.ReadUShort())
		_fcc.Skip(2)
	}
	if _fcc._dac < _fcc._aegc && _fcc._dac > 0 {
		_edfd := _fcc._gbedf.Widths[_fcc._dac-1]
		for _gcf := _fcc._dac; _gcf < _fcc._aegc; _gcf++ {
			_fcc._gbedf.Widths = append(_fcc._gbedf.Widths, _edfd)
		}
	}
	return nil
}

func _ccg() StdFont {
	_dcf.Do(_ddd)
	_ga := Descriptor{Name: CourierName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0021, BBox: [4]float64{-23, -250, 715, 805}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_ga, _ebbg)
}

func MakeRuneCharSafeMap(length int) *RuneCharSafeMap {
	return &RuneCharSafeMap{_cb: make(map[rune]CharMetrics, length)}
}

func NewStdFont(desc Descriptor, metrics *RuneCharSafeMap) StdFont {
	return NewStdFontWithEncoding(desc, metrics, _bfg.NewStandardEncoder())
}

type Font interface {
	Encoder() _bfg.TextEncoder
	GetRuneMetrics(_ce rune) (CharMetrics, bool)
}

func NewStdFontWithEncoding(desc Descriptor, metrics *RuneCharSafeMap, encoder _bfg.TextEncoder) StdFont {
	var _ca rune = 0xA0
	if _, _bfgd := metrics.Read(_ca); !_bfgd {
		_bcc, _ := metrics.Read(0x20)
		metrics.Write(_ca, _bcc)
	}
	return StdFont{_gb: desc, _fa: metrics, _fc: encoder}
}
func (_bg StdFont) Name() string            { return string(_bg._gb.Name) }
func (_ddg *ttfParser) Skip(n int)          { _ddg._fdbe.Seek(int64(n), _fd.SeekCurrent) }
func (_eda StdFont) Descriptor() Descriptor { return _eda._gb }
func _cdg() StdFont {
	_dfb.Do(_dae)
	_ade := Descriptor{Name: TimesRomanName, Family: _bed, Weight: FontWeightRoman, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 898}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 662, XHeight: 450, StemV: 84, StemH: 28}
	return NewStdFont(_ade, _ge)
}

func (_fff *ttfParser) parseCmapSubtable31(_abbe int64) error {
	_gfeb := make([]rune, 0, 8)
	_af := make([]rune, 0, 8)
	_dbgb := make([]int16, 0, 8)
	_aff := make([]uint16, 0, 8)
	_fff._gbedf.Chars = make(map[rune]GID)
	_fff._fdbe.Seek(int64(_fff._fddc["\u0063\u006d\u0061\u0070"])+_abbe, _fd.SeekStart)
	_gcd := _fff.ReadUShort()
	if _gcd != 4 {
		_fe.Log.Debug("u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0073\u0075\u0062t\u0061\u0062\u006c\u0065\u0020\u0066\u006fr\u006d\u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0028\u0025w\u0029", _gcd)
		return nil
	}
	_fff.Skip(2 * 2)
	_dga := int(_fff.ReadUShort() / 2)
	_fff.Skip(3 * 2)
	for _dec := 0; _dec < _dga; _dec++ {
		_af = append(_af, rune(_fff.ReadUShort()))
	}
	_fff.Skip(2)
	for _ddfg := 0; _ddfg < _dga; _ddfg++ {
		_gfeb = append(_gfeb, rune(_fff.ReadUShort()))
	}
	for _aad := 0; _aad < _dga; _aad++ {
		_dbgb = append(_dbgb, _fff.ReadShort())
	}
	_feb, _ := _fff._fdbe.Seek(int64(0), _fd.SeekCurrent)
	for _dfg := 0; _dfg < _dga; _dfg++ {
		_aff = append(_aff, _fff.ReadUShort())
	}
	for _bbab := 0; _bbab < _dga; _bbab++ {
		_dgb := _gfeb[_bbab]
		_cfeb := _af[_bbab]
		_faag := _dbgb[_bbab]
		_ada := _aff[_bbab]
		if _ada > 0 {
			_fff._fdbe.Seek(_feb+2*int64(_bbab)+int64(_ada), _fd.SeekStart)
		}
		for _agf := _dgb; _agf <= _cfeb; _agf++ {
			if _agf == 0xFFFF {
				break
			}
			var _egafd int32
			if _ada > 0 {
				_egafd = int32(_fff.ReadUShort())
				if _egafd > 0 {
					_egafd += int32(_faag)
				}
			} else {
				_egafd = _agf + int32(_faag)
			}
			if _egafd >= 65536 {
				_egafd -= 65536
			}
			if _egafd > 0 {
				_fff._gbedf.Chars[_agf] = GID(_egafd)
			}
		}
	}
	return nil
}
func (_bfb StdFont) GetMetricsTable() *RuneCharSafeMap { return _bfb._fa }
func (_gbdd *ttfParser) ReadUShort() (_gac uint16) {
	_e.Read(_gbdd._fdbe, _e.BigEndian, &_gac)
	return _gac
}

func (_ed StdFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_cag, _gbe := _ed._fa.Read(r)
	return _cag, _gbe
}

type RuneCharSafeMap struct {
	_cb map[rune]CharMetrics
	_ec _eb.RWMutex
}

func TtfParse(r _fd.ReadSeeker) (TtfType, error) { _gca := &ttfParser{_fdbe: r}; return _gca.Parse() }
func _ae() StdFont {
	_cfa.Do(_cfg)
	_faa := Descriptor{Name: HelveticaObliqueName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-170, -225, 1116, 931}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_faa, _df)
}

func (_gba StdFont) ToPdfObject() _c.PdfObject {
	_ebb := _c.MakeDict()
	_ebb.Set("\u0054\u0079\u0070\u0065", _c.MakeName("\u0046\u006f\u006e\u0074"))
	_ebb.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _c.MakeName("\u0054\u0079\u0070e\u0031"))
	_ebb.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _c.MakeName(_gba.Name()))
	_ebb.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _gba._fc.ToPdfObject())
	return _c.MakeIndirectObject(_ebb)
}

func _dae() {
	_ge = MakeRuneCharSafeMap(len(_cfe))
	_dca = MakeRuneCharSafeMap(len(_cfe))
	_ega = MakeRuneCharSafeMap(len(_cfe))
	_bbe = MakeRuneCharSafeMap(len(_cfe))
	for _ceg, _cda := range _cfe {
		_ge.Write(_cda, CharMetrics{Wx: float64(_bea[_ceg])})
		_dca.Write(_cda, CharMetrics{Wx: float64(_cae[_ceg])})
		_ega.Write(_cda, CharMetrics{Wx: float64(_gce[_ceg])})
		_bbe.Write(_cda, CharMetrics{Wx: float64(_bba[_ceg])})
	}
}

const (
	SymbolName       = StdFontName("\u0053\u0079\u006d\u0062\u006f\u006c")
	ZapfDingbatsName = StdFontName("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
)

var _bba = []int16{611, 889, 611, 611, 611, 611, 611, 611, 611, 611, 611, 611, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 611, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 444, 667, 667, 556, 556, 611, 556, 556, 833, 667, 667, 667, 667, 667, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 611, 611, 611, 611, 500, 500, 500, 500, 500, 556, 556, 556, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 611, 833, 611, 556, 556, 556, 556, 556, 556, 556, 500, 500, 500, 500, 333, 500, 667, 500, 500, 778, 500, 500, 422, 541, 500, 920, 500, 500, 278, 275, 400, 400, 389, 389, 333, 275, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 760, 500, 500, 500, 500, 544, 500, 400, 333, 675, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 889, 444, 889, 500, 444, 675, 500, 333, 389, 278, 500, 500, 500, 500, 500, 167, 500, 500, 500, 500, 333, 675, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 444, 444, 278, 278, 300, 278, 675, 549, 675, 471, 278, 722, 333, 675, 500, 675, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 667, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 523, 333, 333, 476, 833, 250, 250, 1000, 675, 675, 500, 500, 500, 420, 556, 556, 556, 333, 333, 333, 214, 389, 389, 453, 389, 389, 760, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 300, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 444, 667, 444, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func (_ecfe *ttfParser) ParseOS2() error {
	if _dcg := _ecfe.Seek("\u004f\u0053\u002f\u0032"); _dcg != nil {
		return _dcg
	}
	_cfebf := _ecfe.ReadUShort()
	_ecfe.Skip(4 * 2)
	_ecfe.Skip(11*2 + 10 + 4*4 + 4)
	_ceaa := _ecfe.ReadUShort()
	_ecfe._gbedf.Bold = (_ceaa & 32) != 0
	_ecfe.Skip(2 * 2)
	_ecfe._gbedf.TypoAscender = _ecfe.ReadShort()
	_ecfe._gbedf.TypoDescender = _ecfe.ReadShort()
	if _cfebf >= 2 {
		_ecfe.Skip(3*2 + 2*4 + 2)
		_ecfe._gbedf.CapHeight = _ecfe.ReadShort()
	} else {
		_ecfe._gbedf.CapHeight = 0
	}
	return nil
}

func _ecg() StdFont {
	_aeg := _bfg.NewZapfDingbatsEncoder()
	_fdb := Descriptor{Name: ZapfDingbatsName, Family: string(ZapfDingbatsName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-1, -143, 981, 820}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 90, StemH: 28}
	return NewStdFontWithEncoding(_fdb, _gc, _aeg)
}

var _gc = &RuneCharSafeMap{_cb: map[rune]CharMetrics{' ': {Wx: 278}, '→': {Wx: 838}, '↔': {Wx: 1016}, '↕': {Wx: 458}, '①': {Wx: 788}, '②': {Wx: 788}, '③': {Wx: 788}, '④': {Wx: 788}, '⑤': {Wx: 788}, '⑥': {Wx: 788}, '⑦': {Wx: 788}, '⑧': {Wx: 788}, '⑨': {Wx: 788}, '⑩': {Wx: 788}, '■': {Wx: 761}, '▲': {Wx: 892}, '▼': {Wx: 892}, '◆': {Wx: 788}, '●': {Wx: 791}, '◗': {Wx: 438}, '★': {Wx: 816}, '☎': {Wx: 719}, '☛': {Wx: 960}, '☞': {Wx: 939}, '♠': {Wx: 626}, '♣': {Wx: 776}, '♥': {Wx: 694}, '♦': {Wx: 595}, '✁': {Wx: 974}, '✂': {Wx: 961}, '✃': {Wx: 974}, '✄': {Wx: 980}, '✆': {Wx: 789}, '✇': {Wx: 790}, '✈': {Wx: 791}, '✉': {Wx: 690}, '✌': {Wx: 549}, '✍': {Wx: 855}, '✎': {Wx: 911}, '✏': {Wx: 933}, '✐': {Wx: 911}, '✑': {Wx: 945}, '✒': {Wx: 974}, '✓': {Wx: 755}, '✔': {Wx: 846}, '✕': {Wx: 762}, '✖': {Wx: 761}, '✗': {Wx: 571}, '✘': {Wx: 677}, '✙': {Wx: 763}, '✚': {Wx: 760}, '✛': {Wx: 759}, '✜': {Wx: 754}, '✝': {Wx: 494}, '✞': {Wx: 552}, '✟': {Wx: 537}, '✠': {Wx: 577}, '✡': {Wx: 692}, '✢': {Wx: 786}, '✣': {Wx: 788}, '✤': {Wx: 788}, '✥': {Wx: 790}, '✦': {Wx: 793}, '✧': {Wx: 794}, '✩': {Wx: 823}, '✪': {Wx: 789}, '✫': {Wx: 841}, '✬': {Wx: 823}, '✭': {Wx: 833}, '✮': {Wx: 816}, '✯': {Wx: 831}, '✰': {Wx: 923}, '✱': {Wx: 744}, '✲': {Wx: 723}, '✳': {Wx: 749}, '✴': {Wx: 790}, '✵': {Wx: 792}, '✶': {Wx: 695}, '✷': {Wx: 776}, '✸': {Wx: 768}, '✹': {Wx: 792}, '✺': {Wx: 759}, '✻': {Wx: 707}, '✼': {Wx: 708}, '✽': {Wx: 682}, '✾': {Wx: 701}, '✿': {Wx: 826}, '❀': {Wx: 815}, '❁': {Wx: 789}, '❂': {Wx: 789}, '❃': {Wx: 707}, '❄': {Wx: 687}, '❅': {Wx: 696}, '❆': {Wx: 689}, '❇': {Wx: 786}, '❈': {Wx: 787}, '❉': {Wx: 713}, '❊': {Wx: 791}, '❋': {Wx: 785}, '❍': {Wx: 873}, '❏': {Wx: 762}, '❐': {Wx: 762}, '❑': {Wx: 759}, '❒': {Wx: 759}, '❖': {Wx: 784}, '❘': {Wx: 138}, '❙': {Wx: 277}, '❚': {Wx: 415}, '❛': {Wx: 392}, '❜': {Wx: 392}, '❝': {Wx: 668}, '❞': {Wx: 668}, '❡': {Wx: 732}, '❢': {Wx: 544}, '❣': {Wx: 544}, '❤': {Wx: 910}, '❥': {Wx: 667}, '❦': {Wx: 760}, '❧': {Wx: 760}, '❶': {Wx: 788}, '❷': {Wx: 788}, '❸': {Wx: 788}, '❹': {Wx: 788}, '❺': {Wx: 788}, '❻': {Wx: 788}, '❼': {Wx: 788}, '❽': {Wx: 788}, '❾': {Wx: 788}, '❿': {Wx: 788}, '➀': {Wx: 788}, '➁': {Wx: 788}, '➂': {Wx: 788}, '➃': {Wx: 788}, '➄': {Wx: 788}, '➅': {Wx: 788}, '➆': {Wx: 788}, '➇': {Wx: 788}, '➈': {Wx: 788}, '➉': {Wx: 788}, '➊': {Wx: 788}, '➋': {Wx: 788}, '➌': {Wx: 788}, '➍': {Wx: 788}, '➎': {Wx: 788}, '➏': {Wx: 788}, '➐': {Wx: 788}, '➑': {Wx: 788}, '➒': {Wx: 788}, '➓': {Wx: 788}, '➔': {Wx: 894}, '➘': {Wx: 748}, '➙': {Wx: 924}, '➚': {Wx: 748}, '➛': {Wx: 918}, '➜': {Wx: 927}, '➝': {Wx: 928}, '➞': {Wx: 928}, '➟': {Wx: 834}, '➠': {Wx: 873}, '➡': {Wx: 828}, '➢': {Wx: 924}, '➣': {Wx: 924}, '➤': {Wx: 917}, '➥': {Wx: 930}, '➦': {Wx: 931}, '➧': {Wx: 463}, '➨': {Wx: 883}, '➩': {Wx: 836}, '➪': {Wx: 836}, '➫': {Wx: 867}, '➬': {Wx: 867}, '➭': {Wx: 696}, '➮': {Wx: 696}, '➯': {Wx: 874}, '➱': {Wx: 874}, '➲': {Wx: 760}, '➳': {Wx: 946}, '➴': {Wx: 771}, '➵': {Wx: 865}, '➶': {Wx: 771}, '➷': {Wx: 888}, '➸': {Wx: 967}, '➹': {Wx: 888}, '➺': {Wx: 831}, '➻': {Wx: 873}, '➼': {Wx: 927}, '➽': {Wx: 970}, '➾': {Wx: 918}, '\uf8d7': {Wx: 390}, '\uf8d8': {Wx: 390}, '\uf8d9': {Wx: 317}, '\uf8da': {Wx: 317}, '\uf8db': {Wx: 276}, '\uf8dc': {Wx: 276}, '\uf8dd': {Wx: 509}, '\uf8de': {Wx: 509}, '\uf8df': {Wx: 410}, '\uf8e0': {Wx: 410}, '\uf8e1': {Wx: 234}, '\uf8e2': {Wx: 234}, '\uf8e3': {Wx: 334}, '\uf8e4': {Wx: 334}}}

func (_fca StdFont) Encoder() _bfg.TextEncoder { return _fca._fc }
func init() {
	RegisterStdFont(CourierName, _ccg, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u0043\u006f\u0075\u0072\u0069e\u0072\u004e\u0065\u0077", "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077")
	RegisterStdFont(CourierBoldName, _edb, "\u0043o\u0075r\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(CourierObliqueName, _cab, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(CourierBoldObliqueName, _ddc, "C\u006f\u0075\u0072\u0069er\u004ee\u0077\u002c\u0042\u006f\u006cd\u0049\u0074\u0061\u006c\u0069\u0063")
}

func (_fge *ttfParser) ParseHhea() error {
	if _caa := _fge.Seek("\u0068\u0068\u0065\u0061"); _caa != nil {
		return _caa
	}
	_fge.Skip(4 + 15*2)
	_fge._dac = _fge.ReadUShort()
	return nil
}
func (_ace *TtfType) NewEncoder() _bfg.TextEncoder { return _bfg.NewTrueTypeFontEncoder(_ace.Chars) }
func _ddc() StdFont {
	_dcf.Do(_ddd)
	_fdg := Descriptor{Name: CourierBoldObliqueName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0061, BBox: [4]float64{-57, -250, 869, 801}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_fdg, _cba)
}

var _bea = []int16{722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 556, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 389, 722, 722, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 556, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 556, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 611, 611, 611, 611, 444, 444, 444, 444, 333, 444, 667, 444, 444, 778, 444, 444, 469, 541, 500, 921, 444, 500, 278, 200, 480, 480, 333, 333, 333, 200, 350, 444, 444, 333, 444, 444, 333, 500, 333, 278, 250, 250, 760, 500, 500, 500, 500, 588, 500, 400, 333, 564, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 564, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 564, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 344, 278, 564, 549, 564, 471, 278, 778, 333, 564, 500, 564, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 453, 333, 333, 476, 833, 250, 250, 1000, 564, 564, 500, 444, 444, 408, 444, 444, 444, 333, 333, 333, 180, 333, 333, 453, 333, 333, 760, 333, 389, 389, 389, 389, 389, 500, 278, 500, 500, 278, 250, 500, 600, 278, 326, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func IsStdFont(name StdFontName) bool { _, _beb := _dd.read(name); return _beb }

var _ebbg *RuneCharSafeMap

func NewStdFontByName(name StdFontName) (StdFont, bool) {
	_cea, _gf := _dd.read(name)
	if !_gf {
		return StdFont{}, false
	}
	return _cea(), true
}

func (_gegf *ttfParser) Read32Fixed() float64 {
	_cgac := float64(_gegf.ReadShort())
	_bcf := float64(_gegf.ReadUShort()) / 65536.0
	return _cgac + _bcf
}

func (_aag *ttfParser) ParseHead() error {
	if _bab := _aag.Seek("\u0068\u0065\u0061\u0064"); _bab != nil {
		return _bab
	}
	_aag.Skip(3 * 4)
	_dgd := _aag.ReadULong()
	if _dgd != 0x5F0F3CF5 {
		_fe.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0063\u006fr\u0072e\u0063\u0074\u0020\u006d\u0061\u0067\u0069\u0063\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e\u0020\u0046\u006fn\u0074\u0020\u006d\u0061\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0069\u0073\u0070\u006c\u0061\u0079\u0020\u0063\u006f\u0072\u0072\u0065\u0063t\u006c\u0079\u002e\u0020\u0025\u0073", _aag)
	}
	_aag.Skip(2)
	_aag._gbedf.UnitsPerEm = _aag.ReadUShort()
	_aag.Skip(2 * 8)
	_aag._gbedf.Xmin = _aag.ReadShort()
	_aag._gbedf.Ymin = _aag.ReadShort()
	_aag._gbedf.Xmax = _aag.ReadShort()
	_aag._gbedf.Ymax = _aag.ReadShort()
	return nil
}

func (_ab CharMetrics) String() string {
	return _d.Sprintf("<\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u003e", _ab.Wx, _ab.Wy)
}

func _cee() StdFont {
	_cfa.Do(_cfg)
	_ebf := Descriptor{Name: HelveticaBoldObliqueName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-174, -228, 1114, 962}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_ebf, _abb)
}

func init() {
	RegisterStdFont(HelveticaName, _bccd, "\u0041\u0072\u0069a\u006c")
	RegisterStdFont(HelveticaBoldName, _fda, "\u0041\u0072\u0069\u0061\u006c\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(HelveticaObliqueName, _ae, "\u0041\u0072\u0069a\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(HelveticaBoldObliqueName, _cee, "\u0041\u0072i\u0061\u006c\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}

func _ddd() {
	const _gaa = 600
	_ebbg = MakeRuneCharSafeMap(len(_cfe))
	for _, _gbed := range _cfe {
		_ebbg.Write(_gbed, CharMetrics{Wx: _gaa})
	}
	_cga = _ebbg.Copy()
	_cba = _ebbg.Copy()
	_cac = _ebbg.Copy()
}

func (_adee *ttfParser) parseTTC() (TtfType, error) {
	_adee.Skip(2 * 2)
	_cbc := _adee.ReadULong()
	if _cbc < 1 {
		return TtfType{}, _b.New("N\u006f \u0066\u006f\u006e\u0074\u0073\u0020\u0069\u006e \u0054\u0054\u0043\u0020fi\u006c\u0065")
	}
	_dfd := _adee.ReadULong()
	_, _aeb := _adee._fdbe.Seek(int64(_dfd), _fd.SeekStart)
	if _aeb != nil {
		return TtfType{}, _aeb
	}
	return _adee.Parse()
}

func (_bae *ttfParser) Seek(tag string) error {
	_gcbg, _gecd := _bae._fddc[tag]
	if !_gecd {
		return _d.Errorf("\u0074\u0061\u0062\u006ce \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0073", tag)
	}
	_bae._fdbe.Seek(int64(_gcbg), _fd.SeekStart)
	return nil
}

func _cab() StdFont {
	_dcf.Do(_ddd)
	_ag := Descriptor{Name: CourierObliqueName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0061, BBox: [4]float64{-27, -250, 849, 805}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_ag, _cac)
}

type (
	FontWeight int
	Descriptor struct {
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
)

var _dcf _eb.Once

func _cfg() {
	_cfb = MakeRuneCharSafeMap(len(_cfe))
	_eba = MakeRuneCharSafeMap(len(_cfe))
	for _bebc, _cd := range _cfe {
		_cfb.Write(_cd, CharMetrics{Wx: float64(_egd[_bebc])})
		_eba.Write(_cd, CharMetrics{Wx: float64(_bda[_bebc])})
	}
	_df = _cfb.Copy()
	_abb = _eba.Copy()
}

var (
	_cac *RuneCharSafeMap
	_cae = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 611, 778, 778, 778, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 778, 778, 667, 667, 667, 667, 667, 944, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 611, 778, 722, 722, 722, 722, 556, 556, 556, 556, 556, 667, 667, 667, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 1000, 722, 722, 722, 722, 667, 667, 667, 667, 500, 500, 500, 500, 333, 500, 722, 500, 500, 833, 500, 500, 581, 520, 500, 930, 500, 556, 278, 220, 394, 394, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 556, 500, 500, 672, 556, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 556, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 333, 556, 556, 278, 278, 394, 278, 570, 549, 570, 494, 278, 833, 333, 570, 556, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 300, 330, 500, 500, 556, 540, 333, 333, 494, 1000, 250, 250, 1000, 570, 570, 556, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 444, 444, 549, 444, 444, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 333, 416, 333, 556, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}
)

type GlyphName = _bfg.GlyphName

func (_dfge *ttfParser) readByte() (_gea uint8) {
	_e.Read(_dfge._fdbe, _e.BigEndian, &_gea)
	return _gea
}

func NewFontFile2FromPdfObject(obj _c.PdfObject) (TtfType, error) {
	obj = _c.TraceToDirectObject(obj)
	_def, _bfa := obj.(*_c.PdfObjectStream)
	if !_bfa {
		_fe.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0073\u0074\u0072e\u0061\u006d \u0028\u0025\u0054\u0029", obj)
		return TtfType{}, _c.ErrTypeError
	}
	_aba, _ecfc := _c.DecodeStream(_def)
	if _ecfc != nil {
		return TtfType{}, _ecfc
	}
	_cdga := ttfParser{_fdbe: _bf.NewReader(_aba)}
	return _cdga.Parse()
}

func init() {
	RegisterStdFont(SymbolName, _edf, "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063", "S\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064", "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(ZapfDingbatsName, _ecg)
}

func (_cbgg *fontMap) read(_cbd StdFontName) (func() StdFont, bool) {
	_cbgg.Lock()
	defer _cbgg.Unlock()
	_bc, _bfe := _cbgg._cg[_cbd]
	return _bc, _bfe
}

func (_gbd *ttfParser) parseCmapFormat6() error {
	_gcag := int(_gbd.ReadUShort())
	_fcaa := int(_gbd.ReadUShort())
	_fe.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u0046o\u0072\u006d\u0061\u0074\u0036\u003a\u0020%s\u0020\u0066\u0069\u0072s\u0074\u0043\u006f\u0064\u0065\u003d\u0025\u0064\u0020en\u0074\u0072y\u0043\u006f\u0075\u006e\u0074\u003d\u0025\u0064", _gbd._gbedf.String(), _gcag, _fcaa)
	for _ffa := 0; _ffa < _fcaa; _ffa++ {
		_eadc := GID(_gbd.ReadUShort())
		_gbd._gbedf.Chars[rune(_ffa+_gcag)] = _eadc
	}
	return nil
}

func (_ff *ttfParser) ParseComponents() error {
	if _dbg := _ff.ParseHead(); _dbg != nil {
		return _dbg
	}
	if _fdgd := _ff.ParseHhea(); _fdgd != nil {
		return _fdgd
	}
	if _fbg := _ff.ParseMaxp(); _fbg != nil {
		return _fbg
	}
	if _dcd := _ff.ParseHmtx(); _dcd != nil {
		return _dcd
	}
	if _, _edg := _ff._fddc["\u006e\u0061\u006d\u0065"]; _edg {
		if _gfc := _ff.ParseName(); _gfc != nil {
			return _gfc
		}
	}
	if _, _bbd := _ff._fddc["\u004f\u0053\u002f\u0032"]; _bbd {
		if _ecb := _ff.ParseOS2(); _ecb != nil {
			return _ecb
		}
	}
	if _, _cdaa := _ff._fddc["\u0070\u006f\u0073\u0074"]; _cdaa {
		if _bce := _ff.ParsePost(); _bce != nil {
			return _bce
		}
	}
	if _, _edc := _ff._fddc["\u0063\u006d\u0061\u0070"]; _edc {
		if _gga := _ff.ParseCmap(); _gga != nil {
			return _gga
		}
	}
	return nil
}

var _bbe *RuneCharSafeMap

func (_ea *RuneCharSafeMap) Read(b rune) (CharMetrics, bool) {
	_ea._ec.RLock()
	defer _ea._ec.RUnlock()
	_cc, _ad := _ea._cb[b]
	return _cc, _ad
}

func init() {
	RegisterStdFont(TimesRomanName, _cdg, "\u0054\u0069\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e", "\u0054\u0069\u006de\u0073")
	RegisterStdFont(TimesBoldName, _efc, "\u0054i\u006de\u0073\u004e\u0065\u0077\u0052o\u006d\u0061n\u002c\u0042\u006f\u006c\u0064", "\u0054\u0069\u006d\u0065\u0073\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(TimesItalicName, _eab, "T\u0069m\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006da\u006e\u002c\u0049\u0074al\u0069\u0063", "\u0054\u0069\u006de\u0073\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(TimesBoldItalicName, _feg, "\u0054i\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e,\u0042\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063", "\u0054\u0069m\u0065\u0073\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}

func _eab() StdFont {
	_dfb.Do(_dae)
	_fdd := Descriptor{Name: TimesItalicName, Family: _bed, Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-169, -217, 1010, 883}, ItalicAngle: -15.5, Ascent: 683, Descent: -217, CapHeight: 653, XHeight: 441, StemV: 76, StemH: 32}
	return NewStdFont(_fdd, _bbe)
}

func RegisterStdFont(name StdFontName, fnc func() StdFont, aliases ...StdFontName) {
	if _, _abgg := _dd.read(name); _abgg {
		panic("\u0066o\u006e\u0074\u0020\u0061l\u0072\u0065\u0061\u0064\u0079 \u0072e\u0067i\u0073\u0074\u0065\u0072\u0065\u0064\u003a " + string(name))
	}
	_dd.write(name, fnc)
	for _, _ead := range aliases {
		RegisterStdFont(_ead, fnc)
	}
}

type GID = _bfg.GID

var _gce = []int16{667, 944, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 667, 722, 722, 722, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 667, 667, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 889, 667, 611, 611, 611, 611, 611, 611, 611, 500, 500, 500, 500, 333, 500, 722, 500, 500, 778, 500, 500, 570, 570, 500, 832, 500, 500, 278, 220, 348, 348, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 500, 500, 500, 608, 500, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 389, 389, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 382, 278, 570, 549, 606, 494, 278, 778, 333, 606, 576, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 266, 300, 500, 500, 500, 500, 333, 333, 494, 833, 250, 250, 1000, 570, 570, 500, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 389, 389, 549, 389, 389, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 366, 278, 500, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 444, 667, 500, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func (_gg *TtfType) MakeToUnicode() *_be.CMap {
	_fgg := make(map[_be.CharCode]rune)
	if len(_gg.GlyphNames) == 0 {
		for _cfd := range _gg.Chars {
			_fgg[_be.CharCode(_cfd)] = _cfd
		}
		return _be.NewToUnicodeCMap(_fgg)
	}
	for _fdgg, _fb := range _gg.Chars {
		_fdbf := _be.CharCode(_fdgg)
		_ddec := _gg.GlyphNames[_fb]
		_aegf, _dgc := _bfg.GlyphToRune(_ddec)
		if !_dgc {
			_fe.Log.Debug("\u004e\u006f \u0072\u0075\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u0020\u0067\u006c\u0079\u0070h=\u0025\u0071", _fdgg, _ddec)
			_aegf = _bfg.MissingCodeRune
		}
		_fgg[_fdbf] = _aegf
	}
	return _be.NewToUnicodeCMap(_fgg)
}

func (_fdef *ttfParser) ReadStr(length int) (string, error) {
	_ffbc := make([]byte, length)
	_dgbb, _gbbb := _fdef._fdbe.Read(_ffbc)
	if _gbbb != nil {
		return "", _gbbb
	} else if _dgbb != length {
		return "", _d.Errorf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073", length)
	}
	return string(_ffbc), nil
}

func (_fec *fontMap) write(_aa StdFontName, _cbgb func() StdFont) {
	_fec.Lock()
	defer _fec.Unlock()
	_fec._cg[_aa] = _cbgb
}

func (_dfa *ttfParser) parseCmapSubtable10(_bbf int64) error {
	if _dfa._gbedf.Chars == nil {
		_dfa._gbedf.Chars = make(map[rune]GID)
	}
	_dfa._fdbe.Seek(int64(_dfa._fddc["\u0063\u006d\u0061\u0070"])+_bbf, _fd.SeekStart)
	var _edbb, _cgf uint32
	_cbf := _dfa.ReadUShort()
	if _cbf < 8 {
		_edbb = uint32(_dfa.ReadUShort())
		_cgf = uint32(_dfa.ReadUShort())
	} else {
		_dfa.ReadUShort()
		_edbb = _dfa.ReadULong()
		_cgf = _dfa.ReadULong()
	}
	_fe.Log.Trace("\u0070\u0061r\u0073\u0065\u0043\u006d\u0061p\u0053\u0075\u0062\u0074\u0061b\u006c\u0065\u0031\u0030\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065\u003d\u0025\u0064", _cbf, _edbb, _cgf)
	if _cbf != 0 {
		return _b.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006d\u0061p\u0020s\u0075\u0062\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	_ebd, _eebb := _dfa.ReadStr(256)
	if _eebb != nil {
		return _eebb
	}
	_bcb := []byte(_ebd)
	for _ceeb, _geb := range _bcb {
		_dfa._gbedf.Chars[rune(_ceeb)] = GID(_geb)
		if _geb != 0 {
			_d.Printf("\u0009\u0030\u0078\u002502\u0078\u0020\u279e\u0020\u0030\u0078\u0025\u0030\u0032\u0078\u003d\u0025\u0063\u000a", _ceeb, _geb, rune(_geb))
		}
	}
	return nil
}

var _cfb *RuneCharSafeMap

func (_adg *ttfParser) ParseMaxp() error {
	if _bebg := _adg.Seek("\u006d\u0061\u0078\u0070"); _bebg != nil {
		return _bebg
	}
	_adg.Skip(4)
	_adg._aegc = _adg.ReadUShort()
	return nil
}

func (_eaba *ttfParser) ParseName() error {
	if _faf := _eaba.Seek("\u006e\u0061\u006d\u0065"); _faf != nil {
		return _faf
	}
	_cacc, _ := _eaba._fdbe.Seek(0, _fd.SeekCurrent)
	_eaba._gbedf.PostScriptName = ""
	_eaba.Skip(2)
	_dge := _eaba.ReadUShort()
	_bdgd := _eaba.ReadUShort()
	for _ebdf := uint16(0); _ebdf < _dge && _eaba._gbedf.PostScriptName == ""; _ebdf++ {
		_eaba.Skip(3 * 2)
		_gbdf := _eaba.ReadUShort()
		_fae := _eaba.ReadUShort()
		_bfad := _eaba.ReadUShort()
		if _gbdf == 6 {
			_eaba._fdbe.Seek(_cacc+int64(_bdgd)+int64(_bfad), _fd.SeekStart)
			_dbc, _dbfg := _eaba.ReadStr(int(_fae))
			if _dbfg != nil {
				return _dbfg
			}
			_dbc = _a.Replace(_dbc, "\u0000", "", -1)
			_fgb, _dbfg := _gd.Compile("\u005b\u0028\u0029\u007b\u007d\u003c\u003e\u0020\u002f%\u005b\u005c\u005d\u005d")
			if _dbfg != nil {
				return _dbfg
			}
			_eaba._gbedf.PostScriptName = _fgb.ReplaceAllString(_dbc, "")
		}
	}
	if _eaba._gbedf.PostScriptName == "" {
		_fe.Log.Debug("\u0050a\u0072\u0073e\u004e\u0061\u006de\u003a\u0020\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u0050\u006f\u0073t\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0077\u0061\u0073\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	}
	return nil
}

func _abf(_ddeb map[string]uint32) string {
	var _cfc []string
	for _fbd := range _ddeb {
		_cfc = append(_cfc, _fbd)
	}
	_da.Slice(_cfc, func(_efb, _aced int) bool { return _ddeb[_cfc[_efb]] < _ddeb[_cfc[_aced]] })
	_cadb := []string{_d.Sprintf("\u0054\u0072\u0075\u0065Ty\u0070\u0065\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_ddeb))}
	for _, _fddf := range _cfc {
		_cadb = append(_cadb, _d.Sprintf("\u0009%\u0071\u0020\u0025\u0035\u0064", _fddf, _ddeb[_fddf]))
	}
	return _a.Join(_cadb, "\u000a")
}

var (
	_egd = []int16{667, 1000, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 500, 667, 667, 556, 556, 556, 556, 556, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 667, 556, 556, 469, 584, 389, 1015, 556, 556, 278, 260, 334, 334, 278, 278, 333, 260, 350, 500, 500, 333, 500, 500, 333, 556, 333, 278, 278, 250, 737, 556, 556, 556, 556, 643, 556, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 556, 278, 333, 278, 500, 556, 500, 556, 556, 167, 556, 556, 556, 611, 333, 584, 549, 556, 556, 333, 333, 556, 333, 333, 222, 278, 278, 278, 278, 278, 222, 222, 500, 500, 222, 222, 299, 222, 584, 549, 584, 471, 222, 833, 333, 584, 556, 584, 556, 556, 556, 556, 556, 549, 556, 556, 556, 556, 556, 556, 944, 333, 556, 556, 556, 556, 834, 834, 333, 370, 365, 611, 556, 556, 537, 333, 333, 476, 889, 278, 278, 1000, 584, 584, 556, 556, 611, 355, 333, 333, 333, 222, 222, 222, 191, 333, 333, 453, 333, 333, 737, 333, 500, 500, 500, 500, 500, 556, 278, 556, 556, 278, 278, 556, 600, 278, 317, 278, 556, 556, 834, 333, 333, 1000, 556, 333, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 500, 722, 500, 500, 500, 500, 556, 500, 500, 500, 500, 556}
	_ge  *RuneCharSafeMap
)

const (
	CourierName            = StdFontName("\u0043o\u0075\u0072\u0069\u0065\u0072")
	CourierBoldName        = StdFontName("\u0043\u006f\u0075r\u0069\u0065\u0072\u002d\u0042\u006f\u006c\u0064")
	CourierObliqueName     = StdFontName("\u0043o\u0075r\u0069\u0065\u0072\u002d\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
	CourierBoldObliqueName = StdFontName("\u0043\u006f\u0075\u0072ie\u0072\u002d\u0042\u006f\u006c\u0064\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
)

var _ Font = StdFont{}

type CharMetrics struct {
	Wx float64
	Wy float64
}

func (_abd *ttfParser) parseCmapFormat0() error {
	_ggb, _cbfe := _abd.ReadStr(256)
	if _cbfe != nil {
		return _cbfe
	}
	_bef := []byte(_ggb)
	_fe.Log.Trace("\u0070a\u0072\u0073e\u0043\u006d\u0061p\u0046\u006f\u0072\u006d\u0061\u0074\u0030:\u0020\u0025\u0073\u000a\u0064\u0061t\u0061\u0053\u0074\u0072\u003d\u0025\u002b\u0071\u000a\u0064\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d", _abd._gbedf.String(), _ggb, _bef)
	for _eeaa, _ede := range _bef {
		_abd._gbedf.Chars[rune(_eeaa)] = GID(_ede)
	}
	return nil
}

var _cfe = []rune{'A', 'Æ', 'Á', 'Ă', 'Â', 'Ä', 'À', 'Ā', 'Ą', 'Å', 'Ã', 'B', 'C', 'Ć', 'Č', 'Ç', 'D', 'Ď', 'Đ', '∆', 'E', 'É', 'Ě', 'Ê', 'Ë', 'Ė', 'È', 'Ē', 'Ę', 'Ð', '€', 'F', 'G', 'Ğ', 'Ģ', 'H', 'I', 'Í', 'Î', 'Ï', 'İ', 'Ì', 'Ī', 'Į', 'J', 'K', 'Ķ', 'L', 'Ĺ', 'Ľ', 'Ļ', 'Ł', 'M', 'N', 'Ń', 'Ň', 'Ņ', 'Ñ', 'O', 'Œ', 'Ó', 'Ô', 'Ö', 'Ò', 'Ő', 'Ō', 'Ø', 'Õ', 'P', 'Q', 'R', 'Ŕ', 'Ř', 'Ŗ', 'S', 'Ś', 'Š', 'Ş', 'Ș', 'T', 'Ť', 'Ţ', 'Þ', 'U', 'Ú', 'Û', 'Ü', 'Ù', 'Ű', 'Ū', 'Ų', 'Ů', 'V', 'W', 'X', 'Y', 'Ý', 'Ÿ', 'Z', 'Ź', 'Ž', 'Ż', 'a', 'á', 'ă', 'â', '´', 'ä', 'æ', 'à', 'ā', '&', 'ą', 'å', '^', '~', '*', '@', 'ã', 'b', '\\', '|', '{', '}', '[', ']', '˘', '¦', '•', 'c', 'ć', 'ˇ', 'č', 'ç', '¸', '¢', 'ˆ', ':', ',', '\uf6c3', '©', '¤', 'd', '†', '‡', 'ď', 'đ', '°', '¨', '÷', '$', '˙', 'ı', 'e', 'é', 'ě', 'ê', 'ë', 'ė', 'è', '8', '…', 'ē', '—', '–', 'ę', '=', 'ð', '!', '¡', 'f', 'ﬁ', '5', 'ﬂ', 'ƒ', '4', '⁄', 'g', 'ğ', 'ģ', 'ß', '`', '>', '≥', '«', '»', '‹', '›', 'h', '˝', '-', 'i', 'í', 'î', 'ï', 'ì', 'ī', 'į', 'j', 'k', 'ķ', 'l', 'ĺ', 'ľ', 'ļ', '<', '≤', '¬', '◊', 'ł', 'm', '¯', '−', 'µ', '×', 'n', 'ń', 'ň', 'ņ', '9', '≠', 'ñ', '#', 'o', 'ó', 'ô', 'ö', 'œ', '˛', 'ò', 'ő', 'ō', '1', '½', '¼', '¹', 'ª', 'º', 'ø', 'õ', 'p', '¶', '(', ')', '∂', '%', '.', '·', '‰', '+', '±', 'q', '?', '¿', '"', '„', '“', '”', '‘', '’', '‚', '\'', 'r', 'ŕ', '√', 'ř', 'ŗ', '®', '˚', 's', 'ś', 'š', 'ş', 'ș', '§', ';', '7', '6', '/', ' ', '£', '∑', 't', 'ť', 'ţ', 'þ', '3', '¾', '³', '˜', '™', '2', '²', 'u', 'ú', 'û', 'ü', 'ù', 'ű', 'ū', '_', 'ų', 'ů', 'v', 'w', 'x', 'y', 'ý', 'ÿ', '¥', 'z', 'ź', 'ž', 'ż', '0'}

func _fda() StdFont {
	_cfa.Do(_cfg)
	_eef := Descriptor{Name: HelveticaBoldName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-170, -228, 1003, 962}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_eef, _eba)
}

func (_cbfd *ttfParser) ParsePost() error {
	if _egb := _cbfd.Seek("\u0070\u006f\u0073\u0074"); _egb != nil {
		return _egb
	}
	_fad := _cbfd.Read32Fixed()
	_cbfd._gbedf.ItalicAngle = _cbfd.Read32Fixed()
	_cbfd._gbedf.UnderlinePosition = _cbfd.ReadShort()
	_cbfd._gbedf.UnderlineThickness = _cbfd.ReadShort()
	_cbfd._gbedf.IsFixedPitch = _cbfd.ReadULong() != 0
	_cbfd.ReadULong()
	_cbfd.ReadULong()
	_cbfd.ReadULong()
	_cbfd.ReadULong()
	_fe.Log.Trace("\u0050a\u0072\u0073\u0065\u0050\u006f\u0073\u0074\u003a\u0020\u0066\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065\u003d\u0025\u0066", _fad)
	switch _fad {
	case 1.0:
		_cbfd._gbedf.GlyphNames = _ffbe
	case 2.0:
		_gcba := int(_cbfd.ReadUShort())
		_bbc := make([]int, _gcba)
		_cbfd._gbedf.GlyphNames = make([]GlyphName, _gcba)
		_adaa := -1
		for _fdec := 0; _fdec < _gcba; _fdec++ {
			_fgd := int(_cbfd.ReadUShort())
			_bbc[_fdec] = _fgd
			if _fgd <= 0x7fff && _fgd > _adaa {
				_adaa = _fgd
			}
		}
		var _gfg []GlyphName
		if _adaa >= len(_ffbe) {
			_gfg = make([]GlyphName, _adaa-len(_ffbe)+1)
			for _aab := 0; _aab < _adaa-len(_ffbe)+1; _aab++ {
				_aabf := int(_cbfd.readByte())
				_deb, _ggbg := _cbfd.ReadStr(_aabf)
				if _ggbg != nil {
					return _ggbg
				}
				_gfg[_aab] = GlyphName(_deb)
			}
		}
		for _badc := 0; _badc < _gcba; _badc++ {
			_dbdg := _bbc[_badc]
			if _dbdg < len(_ffbe) {
				_cbfd._gbedf.GlyphNames[_badc] = _ffbe[_dbdg]
			} else if _dbdg >= len(_ffbe) && _dbdg <= 32767 {
				_cbfd._gbedf.GlyphNames[_badc] = _gfg[_dbdg-len(_ffbe)]
			} else {
				_cbfd._gbedf.GlyphNames[_badc] = "\u002e\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064"
			}
		}
	case 2.5:
		_afd := make([]int, _cbfd._aegc)
		for _acee := 0; _acee < len(_afd); _acee++ {
			_eafb := int(_cbfd.ReadSByte())
			_afd[_acee] = _acee + 1 + _eafb
		}
		_cbfd._gbedf.GlyphNames = make([]GlyphName, len(_afd))
		for _ggc := 0; _ggc < len(_cbfd._gbedf.GlyphNames); _ggc++ {
			_dcac := _ffbe[_afd[_ggc]]
			_cbfd._gbedf.GlyphNames[_ggc] = _dcac
		}
	case 3.0:
		_fe.Log.Debug("\u004e\u006f\u0020\u0050\u006f\u0073t\u0053\u0063\u0072i\u0070\u0074\u0020n\u0061\u006d\u0065\u0020\u0069\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020is\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e")
	default:
		_fe.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020f\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065=\u0025\u0066", _fad)
	}
	return nil
}

func (_fdf *RuneCharSafeMap) Range(f func(_dg rune, _ef CharMetrics) (_eg bool)) {
	_fdf._ec.RLock()
	defer _fdf._ec.RUnlock()
	for _abg, _ee := range _fdf._cb {
		if f(_abg, _ee) {
			break
		}
	}
}

const (
	HelveticaName            = StdFontName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	HelveticaBoldName        = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064")
	HelveticaObliqueName     = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065")
	HelveticaBoldObliqueName = StdFontName("H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065")
)

type fontMap struct {
	_eb.Mutex
	_cg map[StdFontName]func() StdFont
}

var _cba *RuneCharSafeMap

func _efc() StdFont {
	_dfb.Do(_dae)
	_bfef := Descriptor{Name: TimesBoldName, Family: _bed, Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 935}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 676, XHeight: 461, StemV: 139, StemH: 44}
	return NewStdFont(_bfef, _dca)
}

func (_cce *RuneCharSafeMap) Write(b rune, r CharMetrics) {
	_cce._ec.Lock()
	defer _cce._ec.Unlock()
	_cce._cb[b] = r
}

var (
	_dd  = &fontMap{_cg: make(map[StdFontName]func() StdFont)}
	_dca *RuneCharSafeMap
	_dbd = &RuneCharSafeMap{_cb: map[rune]CharMetrics{' ': {Wx: 250}, '!': {Wx: 333}, '#': {Wx: 500}, '%': {Wx: 833}, '&': {Wx: 778}, '(': {Wx: 333}, ')': {Wx: 333}, '+': {Wx: 549}, ',': {Wx: 250}, '.': {Wx: 250}, '/': {Wx: 278}, '0': {Wx: 500}, '1': {Wx: 500}, '2': {Wx: 500}, '3': {Wx: 500}, '4': {Wx: 500}, '5': {Wx: 500}, '6': {Wx: 500}, '7': {Wx: 500}, '8': {Wx: 500}, '9': {Wx: 500}, ':': {Wx: 278}, ';': {Wx: 278}, '<': {Wx: 549}, '=': {Wx: 549}, '>': {Wx: 549}, '?': {Wx: 444}, '[': {Wx: 333}, ']': {Wx: 333}, '_': {Wx: 500}, '{': {Wx: 480}, '|': {Wx: 200}, '}': {Wx: 480}, '¬': {Wx: 713}, '°': {Wx: 400}, '±': {Wx: 549}, 'µ': {Wx: 576}, '×': {Wx: 549}, '÷': {Wx: 549}, 'ƒ': {Wx: 500}, 'Α': {Wx: 722}, 'Β': {Wx: 667}, 'Γ': {Wx: 603}, 'Ε': {Wx: 611}, 'Ζ': {Wx: 611}, 'Η': {Wx: 722}, 'Θ': {Wx: 741}, 'Ι': {Wx: 333}, 'Κ': {Wx: 722}, 'Λ': {Wx: 686}, 'Μ': {Wx: 889}, 'Ν': {Wx: 722}, 'Ξ': {Wx: 645}, 'Ο': {Wx: 722}, 'Π': {Wx: 768}, 'Ρ': {Wx: 556}, 'Σ': {Wx: 592}, 'Τ': {Wx: 611}, 'Υ': {Wx: 690}, 'Φ': {Wx: 763}, 'Χ': {Wx: 722}, 'Ψ': {Wx: 795}, 'α': {Wx: 631}, 'β': {Wx: 549}, 'γ': {Wx: 411}, 'δ': {Wx: 494}, 'ε': {Wx: 439}, 'ζ': {Wx: 494}, 'η': {Wx: 603}, 'θ': {Wx: 521}, 'ι': {Wx: 329}, 'κ': {Wx: 549}, 'λ': {Wx: 549}, 'ν': {Wx: 521}, 'ξ': {Wx: 493}, 'ο': {Wx: 549}, 'π': {Wx: 549}, 'ρ': {Wx: 549}, 'ς': {Wx: 439}, 'σ': {Wx: 603}, 'τ': {Wx: 439}, 'υ': {Wx: 576}, 'φ': {Wx: 521}, 'χ': {Wx: 549}, 'ψ': {Wx: 686}, 'ω': {Wx: 686}, 'ϑ': {Wx: 631}, 'ϒ': {Wx: 620}, 'ϕ': {Wx: 603}, 'ϖ': {Wx: 713}, '•': {Wx: 460}, '…': {Wx: 1000}, '′': {Wx: 247}, '″': {Wx: 411}, '⁄': {Wx: 167}, '€': {Wx: 750}, 'ℑ': {Wx: 686}, '℘': {Wx: 987}, 'ℜ': {Wx: 795}, 'Ω': {Wx: 768}, 'ℵ': {Wx: 823}, '←': {Wx: 987}, '↑': {Wx: 603}, '→': {Wx: 987}, '↓': {Wx: 603}, '↔': {Wx: 1042}, '↵': {Wx: 658}, '⇐': {Wx: 987}, '⇑': {Wx: 603}, '⇒': {Wx: 987}, '⇓': {Wx: 603}, '⇔': {Wx: 1042}, '∀': {Wx: 713}, '∂': {Wx: 494}, '∃': {Wx: 549}, '∅': {Wx: 823}, '∆': {Wx: 612}, '∇': {Wx: 713}, '∈': {Wx: 713}, '∉': {Wx: 713}, '∋': {Wx: 439}, '∏': {Wx: 823}, '∑': {Wx: 713}, '−': {Wx: 549}, '∗': {Wx: 500}, '√': {Wx: 549}, '∝': {Wx: 713}, '∞': {Wx: 713}, '∠': {Wx: 768}, '∧': {Wx: 603}, '∨': {Wx: 603}, '∩': {Wx: 768}, '∪': {Wx: 768}, '∫': {Wx: 274}, '∴': {Wx: 863}, '∼': {Wx: 549}, '≅': {Wx: 549}, '≈': {Wx: 549}, '≠': {Wx: 549}, '≡': {Wx: 549}, '≤': {Wx: 549}, '≥': {Wx: 549}, '⊂': {Wx: 713}, '⊃': {Wx: 713}, '⊄': {Wx: 713}, '⊆': {Wx: 713}, '⊇': {Wx: 713}, '⊕': {Wx: 768}, '⊗': {Wx: 768}, '⊥': {Wx: 658}, '⋅': {Wx: 250}, '⌠': {Wx: 686}, '⌡': {Wx: 686}, '〈': {Wx: 329}, '〉': {Wx: 329}, '◊': {Wx: 494}, '♠': {Wx: 753}, '♣': {Wx: 753}, '♥': {Wx: 753}, '♦': {Wx: 753}, '\uf6d9': {Wx: 790}, '\uf6da': {Wx: 790}, '\uf6db': {Wx: 890}, '\uf8e5': {Wx: 500}, '\uf8e6': {Wx: 603}, '\uf8e7': {Wx: 1000}, '\uf8e8': {Wx: 790}, '\uf8e9': {Wx: 790}, '\uf8ea': {Wx: 786}, '\uf8eb': {Wx: 384}, '\uf8ec': {Wx: 384}, '\uf8ed': {Wx: 384}, '\uf8ee': {Wx: 384}, '\uf8ef': {Wx: 384}, '\uf8f0': {Wx: 384}, '\uf8f1': {Wx: 494}, '\uf8f2': {Wx: 494}, '\uf8f3': {Wx: 494}, '\uf8f4': {Wx: 494}, '\uf8f5': {Wx: 686}, '\uf8f6': {Wx: 384}, '\uf8f7': {Wx: 384}, '\uf8f8': {Wx: 384}, '\uf8f9': {Wx: 384}, '\uf8fa': {Wx: 384}, '\uf8fb': {Wx: 384}, '\uf8fc': {Wx: 494}, '\uf8fd': {Wx: 494}, '\uf8fe': {Wx: 494}, '\uf8ff': {Wx: 790}}}
)

func (_ccgb *ttfParser) parseCmapVersion(_gcb int64) error {
	_fe.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006da\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a \u006f\u0066\u0066s\u0065t\u003d\u0025\u0064", _gcb)
	if _ccgb._gbedf.Chars == nil {
		_ccgb._gbedf.Chars = make(map[rune]GID)
	}
	_ccgb._fdbe.Seek(int64(_ccgb._fddc["\u0063\u006d\u0061\u0070"])+_gcb, _fd.SeekStart)
	var _beag, _beacf uint32
	_eaf := _ccgb.ReadUShort()
	if _eaf < 8 {
		_beag = uint32(_ccgb.ReadUShort())
		_beacf = uint32(_ccgb.ReadUShort())
	} else {
		_ccgb.ReadUShort()
		_beag = _ccgb.ReadULong()
		_beacf = _ccgb.ReadULong()
	}
	_fe.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0043m\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064 \u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067u\u0061g\u0065\u003d\u0025\u0064", _eaf, _beag, _beacf)
	switch _eaf {
	case 0:
		return _ccgb.parseCmapFormat0()
	case 6:
		return _ccgb.parseCmapFormat6()
	case 12:
		return _ccgb.parseCmapFormat12()
	default:
		_fe.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063m\u0061\u0070\u0020\u0066\u006f\u0072\u006da\u0074\u003d\u0025\u0064", _eaf)
		return nil
	}
}

var (
	_bda  = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 556, 722, 722, 611, 611, 611, 611, 611, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 722, 556, 556, 584, 584, 389, 975, 556, 611, 278, 280, 389, 389, 333, 333, 333, 280, 350, 556, 556, 333, 556, 556, 333, 556, 333, 333, 278, 250, 737, 556, 611, 556, 556, 743, 611, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 611, 333, 333, 333, 611, 556, 611, 556, 556, 167, 611, 611, 611, 611, 333, 584, 549, 556, 556, 333, 333, 611, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 556, 556, 278, 278, 400, 278, 584, 549, 584, 494, 278, 889, 333, 584, 611, 584, 611, 611, 611, 611, 556, 549, 611, 556, 611, 611, 611, 611, 944, 333, 611, 611, 611, 556, 834, 834, 333, 370, 365, 611, 611, 611, 556, 333, 333, 494, 889, 278, 278, 1000, 584, 584, 611, 611, 611, 474, 500, 500, 500, 278, 278, 278, 238, 389, 389, 549, 389, 389, 737, 333, 556, 556, 556, 556, 556, 556, 333, 556, 556, 278, 278, 556, 600, 333, 389, 333, 611, 556, 834, 333, 333, 1000, 556, 333, 611, 611, 611, 611, 611, 611, 611, 556, 611, 611, 556, 778, 556, 556, 556, 556, 556, 500, 500, 500, 500, 556}
	_ffbe = []GlyphName{"\u002en\u006f\u0074\u0064\u0065\u0066", "\u002e\u006e\u0075l\u006c", "\u006e\u006fn\u006d\u0061\u0072k\u0069\u006e\u0067\u0072\u0065\u0074\u0075\u0072\u006e", "\u0073\u0070\u0061c\u0065", "\u0065\u0078\u0063\u006c\u0061\u006d", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c", "\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u0069\u0067\u006e", "\u0064\u006f\u006c\u006c\u0061\u0072", "\u0070e\u0072\u0063\u0065\u006e\u0074", "\u0061m\u0070\u0065\u0072\u0073\u0061\u006ed", "q\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067\u006c\u0065", "\u0070a\u0072\u0065\u006e\u006c\u0065\u0066t", "\u0070\u0061\u0072\u0065\u006e\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0074\u0065\u0072\u0069\u0073\u006b", "\u0070\u006c\u0075\u0073", "\u0063\u006f\u006dm\u0061", "\u0068\u0079\u0070\u0068\u0065\u006e", "\u0070\u0065\u0072\u0069\u006f\u0064", "\u0073\u006c\u0061s\u0068", "\u007a\u0065\u0072\u006f", "\u006f\u006e\u0065", "\u0074\u0077\u006f", "\u0074\u0068\u0072e\u0065", "\u0066\u006f\u0075\u0072", "\u0066\u0069\u0076\u0065", "\u0073\u0069\u0078", "\u0073\u0065\u0076e\u006e", "\u0065\u0069\u0067h\u0074", "\u006e\u0069\u006e\u0065", "\u0063\u006f\u006co\u006e", "\u0073e\u006d\u0069\u0063\u006f\u006c\u006fn", "\u006c\u0065\u0073\u0073", "\u0065\u0071\u0075a\u006c", "\u0067r\u0065\u0061\u0074\u0065\u0072", "\u0071\u0075\u0065\u0073\u0074\u0069\u006f\u006e", "\u0061\u0074", "\u0041", "\u0042", "\u0043", "\u0044", "\u0045", "\u0046", "\u0047", "\u0048", "\u0049", "\u004a", "\u004b", "\u004c", "\u004d", "\u004e", "\u004f", "\u0050", "\u0051", "\u0052", "\u0053", "\u0054", "\u0055", "\u0056", "\u0057", "\u0058", "\u0059", "\u005a", "b\u0072\u0061\u0063\u006b\u0065\u0074\u006c\u0065\u0066\u0074", "\u0062a\u0063\u006b\u0073\u006c\u0061\u0073h", "\u0062\u0072\u0061c\u006b\u0065\u0074\u0072\u0069\u0067\u0068\u0074", "a\u0073\u0063\u0069\u0069\u0063\u0069\u0072\u0063\u0075\u006d", "\u0075\u006e\u0064\u0065\u0072\u0073\u0063\u006f\u0072\u0065", "\u0067\u0072\u0061v\u0065", "\u0061", "\u0062", "\u0063", "\u0064", "\u0065", "\u0066", "\u0067", "\u0068", "\u0069", "\u006a", "\u006b", "\u006c", "\u006d", "\u006e", "\u006f", "\u0070", "\u0071", "\u0072", "\u0073", "\u0074", "\u0075", "\u0076", "\u0077", "\u0078", "\u0079", "\u007a", "\u0062r\u0061\u0063\u0065\u006c\u0065\u0066t", "\u0062\u0061\u0072", "\u0062\u0072\u0061\u0063\u0065\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0063\u0069\u0069\u0074\u0069\u006c\u0064\u0065", "\u0041d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0041\u0072\u0069n\u0067", "\u0043\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0045\u0061\u0063\u0075\u0074\u0065", "\u004e\u0074\u0069\u006c\u0064\u0065", "\u004fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0055d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0061\u0063\u0075\u0074\u0065", "\u0061\u0067\u0072\u0061\u0076\u0065", "a\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0074\u0069\u006c\u0064\u0065", "\u0061\u0072\u0069n\u0067", "\u0063\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0065\u0061\u0063\u0075\u0074\u0065", "\u0065\u0067\u0072\u0061\u0076\u0065", "e\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0065d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0069\u0061\u0063\u0075\u0074\u0065", "\u0069\u0067\u0072\u0061\u0076\u0065", "i\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0069d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006e\u0074\u0069\u006c\u0064\u0065", "\u006f\u0061\u0063\u0075\u0074\u0065", "\u006f\u0067\u0072\u0061\u0076\u0065", "o\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u006fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006f\u0074\u0069\u006c\u0064\u0065", "\u0075\u0061\u0063\u0075\u0074\u0065", "\u0075\u0067\u0072\u0061\u0076\u0065", "u\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0075d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0064\u0061\u0067\u0067\u0065\u0072", "\u0064\u0065\u0067\u0072\u0065\u0065", "\u0063\u0065\u006e\u0074", "\u0073\u0074\u0065\u0072\u006c\u0069\u006e\u0067", "\u0073e\u0063\u0074\u0069\u006f\u006e", "\u0062\u0075\u006c\u006c\u0065\u0074", "\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h", "\u0067\u0065\u0072\u006d\u0061\u006e\u0064\u0062\u006c\u0073", "\u0072\u0065\u0067\u0069\u0073\u0074\u0065\u0072\u0065\u0064", "\u0063o\u0070\u0079\u0072\u0069\u0067\u0068t", "\u0074r\u0061\u0064\u0065\u006d\u0061\u0072k", "\u0061\u0063\u0075t\u0065", "\u0064\u0069\u0065\u0072\u0065\u0073\u0069\u0073", "\u006e\u006f\u0074\u0065\u0071\u0075\u0061\u006c", "\u0041\u0045", "\u004f\u0073\u006c\u0061\u0073\u0068", "\u0069\u006e\u0066\u0069\u006e\u0069\u0074\u0079", "\u0070l\u0075\u0073\u006d\u0069\u006e\u0075s", "\u006ce\u0073\u0073\u0065\u0071\u0075\u0061l", "\u0067\u0072\u0065a\u0074\u0065\u0072\u0065\u0071\u0075\u0061\u006c", "\u0079\u0065\u006e", "\u006d\u0075", "p\u0061\u0072\u0074\u0069\u0061\u006c\u0064\u0069\u0066\u0066", "\u0073u\u006d\u006d\u0061\u0074\u0069\u006fn", "\u0070r\u006f\u0064\u0075\u0063\u0074", "\u0070\u0069", "\u0069\u006e\u0074\u0065\u0067\u0072\u0061\u006c", "o\u0072\u0064\u0066\u0065\u006d\u0069\u006e\u0069\u006e\u0065", "\u006f\u0072\u0064m\u0061\u0073\u0063\u0075\u006c\u0069\u006e\u0065", "\u004f\u006d\u0065g\u0061", "\u0061\u0065", "\u006f\u0073\u006c\u0061\u0073\u0068", "\u0071\u0075\u0065s\u0074\u0069\u006f\u006e\u0064\u006f\u0077\u006e", "\u0065\u0078\u0063\u006c\u0061\u006d\u0064\u006f\u0077\u006e", "\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u006e\u006f\u0074", "\u0072a\u0064\u0069\u0063\u0061\u006c", "\u0066\u006c\u006f\u0072\u0069\u006e", "a\u0070\u0070\u0072\u006f\u0078\u0065\u0071\u0075\u0061\u006c", "\u0044\u0065\u006ct\u0061", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074r\u0069\u0067\u0068\u0074", "\u0065\u006c\u006c\u0069\u0070\u0073\u0069\u0073", "\u006e\u006fn\u0062\u0072\u0065a\u006b\u0069\u006e\u0067\u0073\u0070\u0061\u0063\u0065", "\u0041\u0067\u0072\u0061\u0076\u0065", "\u0041\u0074\u0069\u006c\u0064\u0065", "\u004f\u0074\u0069\u006c\u0064\u0065", "\u004f\u0045", "\u006f\u0065", "\u0065\u006e\u0064\u0061\u0073\u0068", "\u0065\u006d\u0064\u0061\u0073\u0068", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u006c\u0065\u0066\u0074", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c\u0072\u0069\u0067\u0068\u0074", "\u0071u\u006f\u0074\u0065\u006c\u0065\u0066t", "\u0071\u0075\u006f\u0074\u0065\u0072\u0069\u0067\u0068\u0074", "\u0064\u0069\u0076\u0069\u0064\u0065", "\u006co\u007a\u0065\u006e\u0067\u0065", "\u0079d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0059d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e", "\u0063\u0075\u0072\u0072\u0065\u006e\u0063\u0079", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006c\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006cr\u0069\u0067\u0068\u0074", "\u0066\u0069", "\u0066\u006c", "\u0064a\u0067\u0067\u0065\u0072\u0064\u0062l", "\u0070\u0065\u0072\u0069\u006f\u0064\u0063\u0065\u006et\u0065\u0072\u0065\u0064", "\u0071\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067l\u0062\u0061\u0073\u0065", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u0062\u0061\u0073\u0065", "p\u0065\u0072\u0074\u0068\u006f\u0075\u0073\u0061\u006e\u0064", "A\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "E\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0041\u0061\u0063\u0075\u0074\u0065", "\u0045d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0045\u0067\u0072\u0061\u0076\u0065", "\u0049\u0061\u0063\u0075\u0074\u0065", "I\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0049d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0049\u0067\u0072\u0061\u0076\u0065", "\u004f\u0061\u0063\u0075\u0074\u0065", "O\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061\u0070\u0070l\u0065", "\u004f\u0067\u0072\u0061\u0076\u0065", "\u0055\u0061\u0063\u0075\u0074\u0065", "U\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0055\u0067\u0072\u0061\u0076\u0065", "\u0064\u006f\u0074\u006c\u0065\u0073\u0073\u0069", "\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0074\u0069\u006cd\u0065", "\u006d\u0061\u0063\u0072\u006f\u006e", "\u0062\u0072\u0065v\u0065", "\u0064o\u0074\u0061\u0063\u0063\u0065\u006et", "\u0072\u0069\u006e\u0067", "\u0063e\u0064\u0069\u006c\u006c\u0061", "\u0068\u0075\u006eg\u0061\u0072\u0075\u006d\u006c\u0061\u0075\u0074", "\u006f\u0067\u006f\u006e\u0065\u006b", "\u0063\u0061\u0072o\u006e", "\u004c\u0073\u006c\u0061\u0073\u0068", "\u006c\u0073\u006c\u0061\u0073\u0068", "\u0053\u0063\u0061\u0072\u006f\u006e", "\u0073\u0063\u0061\u0072\u006f\u006e", "\u005a\u0063\u0061\u0072\u006f\u006e", "\u007a\u0063\u0061\u0072\u006f\u006e", "\u0062r\u006f\u006b\u0065\u006e\u0062\u0061r", "\u0045\u0074\u0068", "\u0065\u0074\u0068", "\u0059\u0061\u0063\u0075\u0074\u0065", "\u0079\u0061\u0063\u0075\u0074\u0065", "\u0054\u0068\u006fr\u006e", "\u0074\u0068\u006fr\u006e", "\u006d\u0069\u006eu\u0073", "\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "o\u006e\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "t\u0077\u006f\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u0074\u0068\u0072\u0065\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u006fn\u0065\u0068\u0061\u006c\u0066", "\u006f\u006e\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072", "\u0074\u0068\u0072\u0065\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072\u0073", "\u0066\u0072\u0061n\u0063", "\u0047\u0062\u0072\u0065\u0076\u0065", "\u0067\u0062\u0072\u0065\u0076\u0065", "\u0049\u0064\u006f\u0074\u0061\u0063\u0063\u0065\u006e\u0074", "\u0053\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0073\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0043\u0061\u0063\u0075\u0074\u0065", "\u0063\u0061\u0063\u0075\u0074\u0065", "\u0043\u0063\u0061\u0072\u006f\u006e", "\u0063\u0063\u0061\u0072\u006f\u006e", "\u0064\u0063\u0072\u006f\u0061\u0074"}
	_dfb  _eb.Once
)

func (_agff *ttfParser) ReadShort() (_ebfb int16) {
	_e.Read(_agff._fdbe, _e.BigEndian, &_ebfb)
	return _ebfb
}

var _cfa _eb.Once

func _feg() StdFont {
	_dfb.Do(_dae)
	_dde := Descriptor{Name: TimesBoldItalicName, Family: _bed, Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-200, -218, 996, 921}, ItalicAngle: -15, Ascent: 683, Descent: -217, CapHeight: 669, XHeight: 462, StemV: 121, StemH: 42}
	return NewStdFont(_dde, _ega)
}

var _abb *RuneCharSafeMap

func (_ccb *ttfParser) ReadSByte() (_adgb int8) {
	_e.Read(_ccb._fdbe, _e.BigEndian, &_adgb)
	return _adgb
}

func (_fcf *TtfType) String() string {
	return _d.Sprintf("\u0046\u004fN\u0054\u005f\u0046\u0049\u004cE\u0032\u007b\u0025\u0023\u0071 \u0055\u006e\u0069\u0074\u0073\u0050\u0065\u0072\u0045\u006d\u003d\u0025\u0064\u0020\u0042\u006f\u006c\u0064\u003d\u0025\u0074\u0020\u0049\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065\u003d\u0025\u0066\u0020"+"\u0043\u0061pH\u0065\u0069\u0067h\u0074\u003d\u0025\u0064 Ch\u0061rs\u003d\u0025\u0064\u0020\u0047\u006c\u0079ph\u004e\u0061\u006d\u0065\u0073\u003d\u0025d\u007d", _fcf.PostScriptName, _fcf.UnitsPerEm, _fcf.Bold, _fcf.ItalicAngle, _fcf.CapHeight, len(_fcf.Chars), len(_fcf.GlyphNames))
}

const (
	_bed                = "\u0054\u0069\u006de\u0073"
	TimesRomanName      = StdFontName("T\u0069\u006d\u0065\u0073\u002d\u0052\u006f\u006d\u0061\u006e")
	TimesBoldName       = StdFontName("\u0054\u0069\u006d\u0065\u0073\u002d\u0042\u006f\u006c\u0064")
	TimesItalicName     = StdFontName("\u0054\u0069\u006de\u0073\u002d\u0049\u0074\u0061\u006c\u0069\u0063")
	TimesBoldItalicName = StdFontName("\u0054\u0069m\u0065\u0073\u002dB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
)

var (
	_eba *RuneCharSafeMap
	_ega *RuneCharSafeMap
	_cga *RuneCharSafeMap
)

type StdFontName string

func (_fde *RuneCharSafeMap) Length() int {
	_fde._ec.RLock()
	defer _fde._ec.RUnlock()
	return len(_fde._cb)
}

func _edb() StdFont {
	_dcf.Do(_ddd)
	_dc := Descriptor{Name: CourierBoldName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0021, BBox: [4]float64{-113, -250, 749, 801}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_dc, _cga)
}

func _bccd() StdFont {
	_cfa.Do(_cfg)
	_dab := Descriptor{Name: HelveticaName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0020, BBox: [4]float64{-166, -225, 1000, 931}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_dab, _cfb)
}

func (_dba *ttfParser) Parse() (TtfType, error) {
	_gbedfg, _bbaf := _dba.ReadStr(4)
	if _bbaf != nil {
		return TtfType{}, _bbaf
	}
	if _gbedfg == "\u0074\u0074\u0063\u0066" {
		return _dba.parseTTC()
	} else if _gbedfg != "\u0000\u0001\u0000\u0000" && _gbedfg != "\u0074\u0072\u0075\u0065" {
		_fe.Log.Debug("\u0055n\u0072\u0065c\u006f\u0067\u006ei\u007a\u0065\u0064\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066i\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u002e\u0020v\u0065\u0072\u0073\u0069\u006f\u006e\u003d\u0025\u0071", _gbedfg)
	}
	_gdaf := int(_dba.ReadUShort())
	_dba.Skip(3 * 2)
	_dba._fddc = make(map[string]uint32)
	var _adf string
	for _ddf := 0; _ddf < _gdaf; _ddf++ {
		_adf, _bbaf = _dba.ReadStr(4)
		if _bbaf != nil {
			return TtfType{}, _bbaf
		}
		_dba.Skip(4)
		_dcc := _dba.ReadULong()
		_dba.Skip(4)
		_dba._fddc[_adf] = _dcc
	}
	_fe.Log.Trace(_abf(_dba._fddc))
	if _bbaf = _dba.ParseComponents(); _bbaf != nil {
		return TtfType{}, _bbaf
	}
	return _dba._gbedf, nil
}

func (_add *TtfType) MakeEncoder() (_bfg.SimpleEncoder, error) {
	_ecf := make(map[_bfg.CharCode]GlyphName)
	for _bad := _bfg.CharCode(0); _bad <= 256; _bad++ {
		_cbe := rune(_bad)
		_dbf, _cad := _add.Chars[_cbe]
		if !_cad {
			continue
		}
		var _ceb GlyphName
		if int(_dbf) >= 0 && int(_dbf) < len(_add.GlyphNames) {
			_ceb = _add.GlyphNames[_dbf]
		} else {
			_bbb := rune(_dbf)
			if _gfe, _fg := _bfg.RuneToGlyph(_bbb); _fg {
				_ceb = _gfe
			}
		}
		if _ceb != "" {
			_ecf[_bad] = _ceb
		}
	}
	if len(_ecf) == 0 {
		_fe.Log.Debug("WA\u0052\u004eI\u004e\u0047\u003a\u0020\u005a\u0065\u0072\u006f\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002e\u0020\u0074\u0074\u0066=\u0025s\u0020\u0043\u0068\u0061\u0072\u0073\u003d\u005b%\u00200\u0032\u0078]", _add, _add.Chars)
	}
	return _bfg.NewCustomSimpleTextEncoder(_ecf, nil)
}

func TtfParseFile(fileStr string) (TtfType, error) {
	_eea, _egaf := _f.Open(fileStr)
	if _egaf != nil {
		return TtfType{}, _egaf
	}
	defer _eea.Close()
	return TtfParse(_eea)
}

const (
	FontWeightMedium FontWeight = iota
	FontWeightBold
	FontWeightRoman
)

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

func (_gec *ttfParser) parseCmapFormat12() error {
	_dfgd := _gec.ReadULong()
	_fe.Log.Trace("\u0070\u0061\u0072se\u0043\u006d\u0061\u0070\u0046\u006f\u0072\u006d\u0061t\u00312\u003a \u0025s\u0020\u006e\u0075\u006d\u0047\u0072\u006f\u0075\u0070\u0073\u003d\u0025\u0064", _gec._gbedf.String(), _dfgd)
	for _adgf := uint32(0); _adgf < _dfgd; _adgf++ {
		_edaa := _gec.ReadULong()
		_eag := _gec.ReadULong()
		_cfgf := _gec.ReadULong()
		if _edaa > 0x0010FFFF || (0xD800 <= _edaa && _edaa <= 0xDFFF) {
			return _b.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		if _eag < _edaa || _eag > 0x0010FFFF || (0xD800 <= _eag && _eag <= 0xDFFF) {
			return _b.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		for _gebg := _edaa; _gebg <= _eag; _gebg++ {
			if _gebg > 0x10FFFF {
				_fe.Log.Debug("\u0046\u006fr\u006d\u0061\u0074\u0020\u0031\u0032\u0020\u0063\u006d\u0061\u0070\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0062\u0065\u0079\u006f\u006e\u0064\u0020\u0055\u0043\u0053\u002d\u0034")
			}
			_gec._gbedf.Chars[rune(_gebg)] = GID(_cfgf)
			_cfgf++
		}
	}
	return nil
}
