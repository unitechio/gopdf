package fonts

import (
	_da "bytes"
	_gd "encoding/binary"
	_dg "errors"
	_g "fmt"
	_d "io"
	_ff "os"
	_fb "regexp"
	_f "sort"
	_ga "strings"
	_fba "sync"

	_bd "bitbucket.org/shenghui0779/gopdf/common"
	_af "bitbucket.org/shenghui0779/gopdf/core"
	_c "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_e "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_a "golang.org/x/xerrors"
)

var _baga = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 556, 722, 722, 611, 611, 611, 611, 611, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 722, 556, 556, 584, 584, 389, 975, 556, 611, 278, 280, 389, 389, 333, 333, 333, 280, 350, 556, 556, 333, 556, 556, 333, 556, 333, 333, 278, 250, 737, 556, 611, 556, 556, 743, 611, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 611, 333, 333, 333, 611, 556, 611, 556, 556, 167, 611, 611, 611, 611, 333, 584, 549, 556, 556, 333, 333, 611, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 556, 556, 278, 278, 400, 278, 584, 549, 584, 494, 278, 889, 333, 584, 611, 584, 611, 611, 611, 611, 556, 549, 611, 556, 611, 611, 611, 611, 944, 333, 611, 611, 611, 556, 834, 834, 333, 370, 365, 611, 611, 611, 556, 333, 333, 494, 889, 278, 278, 1000, 584, 584, 611, 611, 611, 474, 500, 500, 500, 278, 278, 278, 238, 389, 389, 549, 389, 389, 737, 333, 556, 556, 556, 556, 556, 556, 333, 556, 556, 278, 278, 556, 600, 333, 389, 333, 611, 556, 834, 333, 333, 1000, 556, 333, 611, 611, 611, 611, 611, 611, 611, 556, 611, 611, 556, 778, 556, 556, 556, 556, 556, 500, 500, 500, 500, 556}

type StdFontName string

var _ddc = &RuneCharSafeMap{_cc: map[rune]CharMetrics{' ': {Wx: 278}, '→': {Wx: 838}, '↔': {Wx: 1016}, '↕': {Wx: 458}, '①': {Wx: 788}, '②': {Wx: 788}, '③': {Wx: 788}, '④': {Wx: 788}, '⑤': {Wx: 788}, '⑥': {Wx: 788}, '⑦': {Wx: 788}, '⑧': {Wx: 788}, '⑨': {Wx: 788}, '⑩': {Wx: 788}, '■': {Wx: 761}, '▲': {Wx: 892}, '▼': {Wx: 892}, '◆': {Wx: 788}, '●': {Wx: 791}, '◗': {Wx: 438}, '★': {Wx: 816}, '☎': {Wx: 719}, '☛': {Wx: 960}, '☞': {Wx: 939}, '♠': {Wx: 626}, '♣': {Wx: 776}, '♥': {Wx: 694}, '♦': {Wx: 595}, '✁': {Wx: 974}, '✂': {Wx: 961}, '✃': {Wx: 974}, '✄': {Wx: 980}, '✆': {Wx: 789}, '✇': {Wx: 790}, '✈': {Wx: 791}, '✉': {Wx: 690}, '✌': {Wx: 549}, '✍': {Wx: 855}, '✎': {Wx: 911}, '✏': {Wx: 933}, '✐': {Wx: 911}, '✑': {Wx: 945}, '✒': {Wx: 974}, '✓': {Wx: 755}, '✔': {Wx: 846}, '✕': {Wx: 762}, '✖': {Wx: 761}, '✗': {Wx: 571}, '✘': {Wx: 677}, '✙': {Wx: 763}, '✚': {Wx: 760}, '✛': {Wx: 759}, '✜': {Wx: 754}, '✝': {Wx: 494}, '✞': {Wx: 552}, '✟': {Wx: 537}, '✠': {Wx: 577}, '✡': {Wx: 692}, '✢': {Wx: 786}, '✣': {Wx: 788}, '✤': {Wx: 788}, '✥': {Wx: 790}, '✦': {Wx: 793}, '✧': {Wx: 794}, '✩': {Wx: 823}, '✪': {Wx: 789}, '✫': {Wx: 841}, '✬': {Wx: 823}, '✭': {Wx: 833}, '✮': {Wx: 816}, '✯': {Wx: 831}, '✰': {Wx: 923}, '✱': {Wx: 744}, '✲': {Wx: 723}, '✳': {Wx: 749}, '✴': {Wx: 790}, '✵': {Wx: 792}, '✶': {Wx: 695}, '✷': {Wx: 776}, '✸': {Wx: 768}, '✹': {Wx: 792}, '✺': {Wx: 759}, '✻': {Wx: 707}, '✼': {Wx: 708}, '✽': {Wx: 682}, '✾': {Wx: 701}, '✿': {Wx: 826}, '❀': {Wx: 815}, '❁': {Wx: 789}, '❂': {Wx: 789}, '❃': {Wx: 707}, '❄': {Wx: 687}, '❅': {Wx: 696}, '❆': {Wx: 689}, '❇': {Wx: 786}, '❈': {Wx: 787}, '❉': {Wx: 713}, '❊': {Wx: 791}, '❋': {Wx: 785}, '❍': {Wx: 873}, '❏': {Wx: 762}, '❐': {Wx: 762}, '❑': {Wx: 759}, '❒': {Wx: 759}, '❖': {Wx: 784}, '❘': {Wx: 138}, '❙': {Wx: 277}, '❚': {Wx: 415}, '❛': {Wx: 392}, '❜': {Wx: 392}, '❝': {Wx: 668}, '❞': {Wx: 668}, '❡': {Wx: 732}, '❢': {Wx: 544}, '❣': {Wx: 544}, '❤': {Wx: 910}, '❥': {Wx: 667}, '❦': {Wx: 760}, '❧': {Wx: 760}, '❶': {Wx: 788}, '❷': {Wx: 788}, '❸': {Wx: 788}, '❹': {Wx: 788}, '❺': {Wx: 788}, '❻': {Wx: 788}, '❼': {Wx: 788}, '❽': {Wx: 788}, '❾': {Wx: 788}, '❿': {Wx: 788}, '➀': {Wx: 788}, '➁': {Wx: 788}, '➂': {Wx: 788}, '➃': {Wx: 788}, '➄': {Wx: 788}, '➅': {Wx: 788}, '➆': {Wx: 788}, '➇': {Wx: 788}, '➈': {Wx: 788}, '➉': {Wx: 788}, '➊': {Wx: 788}, '➋': {Wx: 788}, '➌': {Wx: 788}, '➍': {Wx: 788}, '➎': {Wx: 788}, '➏': {Wx: 788}, '➐': {Wx: 788}, '➑': {Wx: 788}, '➒': {Wx: 788}, '➓': {Wx: 788}, '➔': {Wx: 894}, '➘': {Wx: 748}, '➙': {Wx: 924}, '➚': {Wx: 748}, '➛': {Wx: 918}, '➜': {Wx: 927}, '➝': {Wx: 928}, '➞': {Wx: 928}, '➟': {Wx: 834}, '➠': {Wx: 873}, '➡': {Wx: 828}, '➢': {Wx: 924}, '➣': {Wx: 924}, '➤': {Wx: 917}, '➥': {Wx: 930}, '➦': {Wx: 931}, '➧': {Wx: 463}, '➨': {Wx: 883}, '➩': {Wx: 836}, '➪': {Wx: 836}, '➫': {Wx: 867}, '➬': {Wx: 867}, '➭': {Wx: 696}, '➮': {Wx: 696}, '➯': {Wx: 874}, '➱': {Wx: 874}, '➲': {Wx: 760}, '➳': {Wx: 946}, '➴': {Wx: 771}, '➵': {Wx: 865}, '➶': {Wx: 771}, '➷': {Wx: 888}, '➸': {Wx: 967}, '➹': {Wx: 888}, '➺': {Wx: 831}, '➻': {Wx: 873}, '➼': {Wx: 927}, '➽': {Wx: 970}, '➾': {Wx: 918}, '\uf8d7': {Wx: 390}, '\uf8d8': {Wx: 390}, '\uf8d9': {Wx: 317}, '\uf8da': {Wx: 317}, '\uf8db': {Wx: 276}, '\uf8dc': {Wx: 276}, '\uf8dd': {Wx: 509}, '\uf8de': {Wx: 509}, '\uf8df': {Wx: 410}, '\uf8e0': {Wx: 410}, '\uf8e1': {Wx: 234}, '\uf8e2': {Wx: 234}, '\uf8e3': {Wx: 334}, '\uf8e4': {Wx: 334}}}
var _edc = []int16{722, 1000, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 611, 778, 778, 778, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 778, 778, 667, 667, 667, 667, 667, 944, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 611, 778, 722, 722, 722, 722, 556, 556, 556, 556, 556, 667, 667, 667, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 1000, 722, 722, 722, 722, 667, 667, 667, 667, 500, 500, 500, 500, 333, 500, 722, 500, 500, 833, 500, 500, 581, 520, 500, 930, 500, 556, 278, 220, 394, 394, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 556, 500, 500, 672, 556, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 556, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 333, 556, 556, 278, 278, 394, 278, 570, 549, 570, 494, 278, 833, 333, 570, 556, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 300, 330, 500, 500, 556, 540, 333, 333, 494, 1000, 250, 250, 1000, 570, 570, 556, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 444, 444, 549, 444, 444, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 333, 416, 333, 556, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func (_gacf *ttfParser) parseCmapFormat6() error {
	_geb := int(_gacf.ReadUShort())
	_dfg := int(_gacf.ReadUShort())
	_bd.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u0046o\u0072\u006d\u0061\u0074\u0036\u003a\u0020%s\u0020\u0066\u0069\u0072s\u0074\u0043\u006f\u0064\u0065\u003d\u0025\u0064\u0020en\u0074\u0072y\u0043\u006f\u0075\u006e\u0074\u003d\u0025\u0064", _gacf._cgc.String(), _geb, _dfg)
	for _cec := 0; _cec < _dfg; _cec++ {
		_eef := GID(_gacf.ReadUShort())
		_gacf._cgc.Chars[rune(_cec+_geb)] = _eef
	}
	return nil
}

const (
	FontWeightMedium FontWeight = iota
	FontWeightBold
	FontWeightRoman
)

func (_afdc *TtfType) String() string {
	return _g.Sprintf("\u0046\u004fN\u0054\u005f\u0046\u0049\u004cE\u0032\u007b\u0025\u0023\u0071 \u0055\u006e\u0069\u0074\u0073\u0050\u0065\u0072\u0045\u006d\u003d\u0025\u0064\u0020\u0042\u006f\u006c\u0064\u003d\u0025\u0074\u0020\u0049\u0074\u0061\u006c\u0069\u0063\u0041\u006e\u0067\u006c\u0065\u003d\u0025\u0066\u0020"+"\u0043\u0061pH\u0065\u0069\u0067h\u0074\u003d\u0025\u0064 Ch\u0061rs\u003d\u0025\u0064\u0020\u0047\u006c\u0079ph\u004e\u0061\u006d\u0065\u0073\u003d\u0025d\u007d", _afdc.PostScriptName, _afdc.UnitsPerEm, _afdc.Bold, _afdc.ItalicAngle, _afdc.CapHeight, len(_afdc.Chars), len(_afdc.GlyphNames))
}

var _ffa *RuneCharSafeMap

func NewStdFontByName(name StdFontName) (StdFont, bool) {
	_fec, _cea := _afd.read(name)
	if !_cea {
		return StdFont{}, false
	}
	return _fec(), true
}
func (_afge *ttfParser) Parse() (TtfType, error) {
	_ggg, _dec := _afge.ReadStr(4)
	if _dec != nil {
		return TtfType{}, _dec
	}
	if _ggg == "\u004f\u0054\u0054\u004f" {
		return TtfType{}, _a.Errorf("\u0066\u006f\u006e\u0074s\u0020\u0062\u0061\u0073\u0065\u0064\u0020\u006f\u006e \u0050\u006f\u0073\u0074\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u006f\u0075\u0074\u006c\u0069\u006e\u0065s\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0028\u0025\u0077\u0029", _af.ErrNotSupported)
	}
	if _ggg != "\u0000\u0001\u0000\u0000" && _ggg != "\u0074\u0072\u0075\u0065" {
		_bd.Log.Debug("\u0055n\u0072\u0065c\u006f\u0067\u006ei\u007a\u0065\u0064\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066i\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u002e\u0020v\u0065\u0072\u0073\u0069\u006f\u006e\u003d\u0025\u0071", _ggg)
	}
	_gfc := int(_afge.ReadUShort())
	_afge.Skip(3 * 2)
	_afge._ebba = make(map[string]uint32)
	var _bgb string
	for _eaeg := 0; _eaeg < _gfc; _eaeg++ {
		_bgb, _dec = _afge.ReadStr(4)
		if _dec != nil {
			return TtfType{}, _dec
		}
		_afge.Skip(4)
		_acf := _afge.ReadULong()
		_afge.Skip(4)
		_afge._ebba[_bgb] = _acf
	}
	_bd.Log.Trace(_bbb(_afge._ebba))
	if _dec = _afge.ParseComponents(); _dec != nil {
		return TtfType{}, _dec
	}
	return _afge._cgc, nil
}
func (_bfef *ttfParser) ReadStr(length int) (string, error) {
	_abb := make([]byte, length)
	_bege, _aaab := _bfef._cbee.Read(_abb)
	if _aaab != nil {
		return "", _aaab
	} else if _bege != length {
		return "", _g.Errorf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0073", length)
	}
	return string(_abb), nil
}
func (_aa *RuneCharSafeMap) Copy() *RuneCharSafeMap {
	_bg := MakeRuneCharSafeMap(_aa.Length())
	_aa.Range(func(_gf rune, _ca CharMetrics) (_ba bool) { _bg._cc[_gf] = _ca; return false })
	return _bg
}

var _bef *RuneCharSafeMap

type fontMap struct {
	_fba.Mutex
	_be map[StdFontName]func() StdFont
}

func (_ega *ttfParser) ReadSByte() (_eaf int8) {
	_gd.Read(_ega._cbee, _gd.BigEndian, &_eaf)
	return _eaf
}
func (_ace *ttfParser) Seek(tag string) error {
	_bcg, _cdad := _ace._ebba[tag]
	if !_cdad {
		return _g.Errorf("\u0074\u0061\u0062\u006ce \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0073", tag)
	}
	_ace._cbee.Seek(int64(_bcg), _d.SeekStart)
	return nil
}

var _gb = []int16{667, 1000, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 722, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 556, 611, 778, 778, 778, 722, 278, 278, 278, 278, 278, 278, 278, 278, 500, 667, 667, 556, 556, 556, 556, 556, 833, 722, 722, 722, 722, 722, 778, 1000, 778, 778, 778, 778, 778, 778, 778, 778, 667, 778, 722, 722, 722, 722, 667, 667, 667, 667, 667, 611, 611, 611, 667, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 944, 667, 667, 667, 667, 611, 611, 611, 611, 556, 556, 556, 556, 333, 556, 889, 556, 556, 667, 556, 556, 469, 584, 389, 1015, 556, 556, 278, 260, 334, 334, 278, 278, 333, 260, 350, 500, 500, 333, 500, 500, 333, 556, 333, 278, 278, 250, 737, 556, 556, 556, 556, 643, 556, 400, 333, 584, 556, 333, 278, 556, 556, 556, 556, 556, 556, 556, 556, 1000, 556, 1000, 556, 556, 584, 556, 278, 333, 278, 500, 556, 500, 556, 556, 167, 556, 556, 556, 611, 333, 584, 549, 556, 556, 333, 333, 556, 333, 333, 222, 278, 278, 278, 278, 278, 222, 222, 500, 500, 222, 222, 299, 222, 584, 549, 584, 471, 222, 833, 333, 584, 556, 584, 556, 556, 556, 556, 556, 549, 556, 556, 556, 556, 556, 556, 944, 333, 556, 556, 556, 556, 834, 834, 333, 370, 365, 611, 556, 556, 537, 333, 333, 476, 889, 278, 278, 1000, 584, 584, 556, 556, 611, 355, 333, 333, 333, 222, 222, 222, 191, 333, 333, 453, 333, 333, 737, 333, 500, 500, 500, 500, 500, 556, 278, 556, 556, 278, 278, 556, 600, 278, 317, 278, 556, 556, 834, 333, 333, 1000, 556, 333, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 500, 722, 500, 500, 500, 500, 556, 500, 500, 500, 500, 556}

func (_ccb *ttfParser) ParseHhea() error {
	if _bba := _ccb.Seek("\u0068\u0068\u0065\u0061"); _bba != nil {
		return _bba
	}
	_ccb.Skip(4 + 15*2)
	_ccb._dbd = _ccb.ReadUShort()
	return nil
}

var _ec = []rune{'A', 'Æ', 'Á', 'Ă', 'Â', 'Ä', 'À', 'Ā', 'Ą', 'Å', 'Ã', 'B', 'C', 'Ć', 'Č', 'Ç', 'D', 'Ď', 'Đ', '∆', 'E', 'É', 'Ě', 'Ê', 'Ë', 'Ė', 'È', 'Ē', 'Ę', 'Ð', '€', 'F', 'G', 'Ğ', 'Ģ', 'H', 'I', 'Í', 'Î', 'Ï', 'İ', 'Ì', 'Ī', 'Į', 'J', 'K', 'Ķ', 'L', 'Ĺ', 'Ľ', 'Ļ', 'Ł', 'M', 'N', 'Ń', 'Ň', 'Ņ', 'Ñ', 'O', 'Œ', 'Ó', 'Ô', 'Ö', 'Ò', 'Ő', 'Ō', 'Ø', 'Õ', 'P', 'Q', 'R', 'Ŕ', 'Ř', 'Ŗ', 'S', 'Ś', 'Š', 'Ş', 'Ș', 'T', 'Ť', 'Ţ', 'Þ', 'U', 'Ú', 'Û', 'Ü', 'Ù', 'Ű', 'Ū', 'Ų', 'Ů', 'V', 'W', 'X', 'Y', 'Ý', 'Ÿ', 'Z', 'Ź', 'Ž', 'Ż', 'a', 'á', 'ă', 'â', '´', 'ä', 'æ', 'à', 'ā', '&', 'ą', 'å', '^', '~', '*', '@', 'ã', 'b', '\\', '|', '{', '}', '[', ']', '˘', '¦', '•', 'c', 'ć', 'ˇ', 'č', 'ç', '¸', '¢', 'ˆ', ':', ',', '\uf6c3', '©', '¤', 'd', '†', '‡', 'ď', 'đ', '°', '¨', '÷', '$', '˙', 'ı', 'e', 'é', 'ě', 'ê', 'ë', 'ė', 'è', '8', '…', 'ē', '—', '–', 'ę', '=', 'ð', '!', '¡', 'f', 'ﬁ', '5', 'ﬂ', 'ƒ', '4', '⁄', 'g', 'ğ', 'ģ', 'ß', '`', '>', '≥', '«', '»', '‹', '›', 'h', '˝', '-', 'i', 'í', 'î', 'ï', 'ì', 'ī', 'į', 'j', 'k', 'ķ', 'l', 'ĺ', 'ľ', 'ļ', '<', '≤', '¬', '◊', 'ł', 'm', '¯', '−', 'µ', '×', 'n', 'ń', 'ň', 'ņ', '9', '≠', 'ñ', '#', 'o', 'ó', 'ô', 'ö', 'œ', '˛', 'ò', 'ő', 'ō', '1', '½', '¼', '¹', 'ª', 'º', 'ø', 'õ', 'p', '¶', '(', ')', '∂', '%', '.', '·', '‰', '+', '±', 'q', '?', '¿', '"', '„', '“', '”', '‘', '’', '‚', '\'', 'r', 'ŕ', '√', 'ř', 'ŗ', '®', '˚', 's', 'ś', 'š', 'ş', 'ș', '§', ';', '7', '6', '/', ' ', '£', '∑', 't', 'ť', 'ţ', 'þ', '3', '¾', '³', '˜', '™', '2', '²', 'u', 'ú', 'û', 'ü', 'ù', 'ű', 'ū', '_', 'ų', 'ů', 'v', 'w', 'x', 'y', 'ý', 'ÿ', '¥', 'z', 'ź', 'ž', 'ż', '0'}

func (_bc *RuneCharSafeMap) Write(b rune, r CharMetrics) {
	_bc._gdc.Lock()
	defer _bc._gdc.Unlock()
	_bc._cc[b] = r
}
func _fbfe() StdFont {
	_fdd.Do(_gcdb)
	_cf := Descriptor{Name: CourierBoldName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0021, BBox: [4]float64{-113, -250, 749, 801}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_cf, _afe)
}

var _afe *RuneCharSafeMap

const (
	SymbolName       = StdFontName("\u0053\u0079\u006d\u0062\u006f\u006c")
	ZapfDingbatsName = StdFontName("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
)

type StdFont struct {
	_ea  Descriptor
	_eec *RuneCharSafeMap
	_gfb _e.TextEncoder
}

func _bb() StdFont {
	_fdd.Do(_gcdb)
	_ceaf := Descriptor{Name: CourierName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0021, BBox: [4]float64{-23, -250, 715, 805}, ItalicAngle: 0, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_ceaf, _afc)
}
func (_dge *RuneCharSafeMap) Read(b rune) (CharMetrics, bool) {
	_dge._gdc.RLock()
	defer _dge._gdc.RUnlock()
	_ed, _cd := _dge._cc[b]
	return _ed, _cd
}

var _fdd _fba.Once
var _gde *RuneCharSafeMap

func (_de StdFont) ToPdfObject() _af.PdfObject {
	_bag := _af.MakeDict()
	_bag.Set("\u0054\u0079\u0070\u0065", _af.MakeName("\u0046\u006f\u006e\u0074"))
	_bag.Set("\u0053u\u0062\u0074\u0079\u0070\u0065", _af.MakeName("\u0054\u0079\u0070e\u0031"))
	_bag.Set("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074", _af.MakeName(_de.Name()))
	_bag.Set("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", _de._gfb.ToPdfObject())
	return _af.MakeIndirectObject(_bag)
}
func _dgc() StdFont {
	_fdd.Do(_gcdb)
	_edfa := Descriptor{Name: CourierObliqueName, Family: string(CourierName), Weight: FontWeightMedium, Flags: 0x0061, BBox: [4]float64{-27, -250, 849, 805}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 426, StemV: 51, StemH: 51}
	return NewStdFont(_edfa, _bce)
}

type GID = _e.GID

func (_fbaf StdFont) GetRuneMetrics(r rune) (CharMetrics, bool) {
	_eg, _ge := _fbaf._eec.Read(r)
	return _eg, _ge
}
func init() {
	RegisterStdFont(CourierName, _bb, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u0043\u006f\u0075\u0072\u0069e\u0072\u004e\u0065\u0077", "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077")
	RegisterStdFont(CourierBoldName, _fbfe, "\u0043o\u0075r\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(CourierObliqueName, _dgc, "\u0043\u006f\u0075\u0072\u0069\u0065\u0072\u004e\u0065\u0077\u002c\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(CourierBoldObliqueName, _dd, "C\u006f\u0075\u0072\u0069er\u004ee\u0077\u002c\u0042\u006f\u006cd\u0049\u0074\u0061\u006c\u0069\u0063")
}
func _cfg() StdFont {
	_gdeb := _e.NewSymbolEncoder()
	_cdee := Descriptor{Name: SymbolName, Family: string(SymbolName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-180, -293, 1090, 1010}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 85, StemH: 92}
	return NewStdFontWithEncoding(_cdee, _dff, _gdeb)
}

const (
	HelveticaName            = StdFontName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	HelveticaBoldName        = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064")
	HelveticaObliqueName     = StdFontName("\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065")
	HelveticaBoldObliqueName = StdFontName("H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065")
)

var _gbdg = []GlyphName{"\u002en\u006f\u0074\u0064\u0065\u0066", "\u002e\u006e\u0075l\u006c", "\u006e\u006fn\u006d\u0061\u0072k\u0069\u006e\u0067\u0072\u0065\u0074\u0075\u0072\u006e", "\u0073\u0070\u0061c\u0065", "\u0065\u0078\u0063\u006c\u0061\u006d", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c", "\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u0069\u0067\u006e", "\u0064\u006f\u006c\u006c\u0061\u0072", "\u0070e\u0072\u0063\u0065\u006e\u0074", "\u0061m\u0070\u0065\u0072\u0073\u0061\u006ed", "q\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067\u006c\u0065", "\u0070a\u0072\u0065\u006e\u006c\u0065\u0066t", "\u0070\u0061\u0072\u0065\u006e\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0074\u0065\u0072\u0069\u0073\u006b", "\u0070\u006c\u0075\u0073", "\u0063\u006f\u006dm\u0061", "\u0068\u0079\u0070\u0068\u0065\u006e", "\u0070\u0065\u0072\u0069\u006f\u0064", "\u0073\u006c\u0061s\u0068", "\u007a\u0065\u0072\u006f", "\u006f\u006e\u0065", "\u0074\u0077\u006f", "\u0074\u0068\u0072e\u0065", "\u0066\u006f\u0075\u0072", "\u0066\u0069\u0076\u0065", "\u0073\u0069\u0078", "\u0073\u0065\u0076e\u006e", "\u0065\u0069\u0067h\u0074", "\u006e\u0069\u006e\u0065", "\u0063\u006f\u006co\u006e", "\u0073e\u006d\u0069\u0063\u006f\u006c\u006fn", "\u006c\u0065\u0073\u0073", "\u0065\u0071\u0075a\u006c", "\u0067r\u0065\u0061\u0074\u0065\u0072", "\u0071\u0075\u0065\u0073\u0074\u0069\u006f\u006e", "\u0061\u0074", "\u0041", "\u0042", "\u0043", "\u0044", "\u0045", "\u0046", "\u0047", "\u0048", "\u0049", "\u004a", "\u004b", "\u004c", "\u004d", "\u004e", "\u004f", "\u0050", "\u0051", "\u0052", "\u0053", "\u0054", "\u0055", "\u0056", "\u0057", "\u0058", "\u0059", "\u005a", "b\u0072\u0061\u0063\u006b\u0065\u0074\u006c\u0065\u0066\u0074", "\u0062a\u0063\u006b\u0073\u006c\u0061\u0073h", "\u0062\u0072\u0061c\u006b\u0065\u0074\u0072\u0069\u0067\u0068\u0074", "a\u0073\u0063\u0069\u0069\u0063\u0069\u0072\u0063\u0075\u006d", "\u0075\u006e\u0064\u0065\u0072\u0073\u0063\u006f\u0072\u0065", "\u0067\u0072\u0061v\u0065", "\u0061", "\u0062", "\u0063", "\u0064", "\u0065", "\u0066", "\u0067", "\u0068", "\u0069", "\u006a", "\u006b", "\u006c", "\u006d", "\u006e", "\u006f", "\u0070", "\u0071", "\u0072", "\u0073", "\u0074", "\u0075", "\u0076", "\u0077", "\u0078", "\u0079", "\u007a", "\u0062r\u0061\u0063\u0065\u006c\u0065\u0066t", "\u0062\u0061\u0072", "\u0062\u0072\u0061\u0063\u0065\u0072\u0069\u0067\u0068\u0074", "\u0061\u0073\u0063\u0069\u0069\u0074\u0069\u006c\u0064\u0065", "\u0041d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0041\u0072\u0069n\u0067", "\u0043\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0045\u0061\u0063\u0075\u0074\u0065", "\u004e\u0074\u0069\u006c\u0064\u0065", "\u004fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0055d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0061\u0063\u0075\u0074\u0065", "\u0061\u0067\u0072\u0061\u0076\u0065", "a\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0061\u0074\u0069\u006c\u0064\u0065", "\u0061\u0072\u0069n\u0067", "\u0063\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0065\u0061\u0063\u0075\u0074\u0065", "\u0065\u0067\u0072\u0061\u0076\u0065", "e\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0065d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0069\u0061\u0063\u0075\u0074\u0065", "\u0069\u0067\u0072\u0061\u0076\u0065", "i\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0069d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006e\u0074\u0069\u006c\u0064\u0065", "\u006f\u0061\u0063\u0075\u0074\u0065", "\u006f\u0067\u0072\u0061\u0076\u0065", "o\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u006fd\u0069\u0065\u0072\u0065\u0073\u0069s", "\u006f\u0074\u0069\u006c\u0064\u0065", "\u0075\u0061\u0063\u0075\u0074\u0065", "\u0075\u0067\u0072\u0061\u0076\u0065", "u\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0075d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0064\u0061\u0067\u0067\u0065\u0072", "\u0064\u0065\u0067\u0072\u0065\u0065", "\u0063\u0065\u006e\u0074", "\u0073\u0074\u0065\u0072\u006c\u0069\u006e\u0067", "\u0073e\u0063\u0074\u0069\u006f\u006e", "\u0062\u0075\u006c\u006c\u0065\u0074", "\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h", "\u0067\u0065\u0072\u006d\u0061\u006e\u0064\u0062\u006c\u0073", "\u0072\u0065\u0067\u0069\u0073\u0074\u0065\u0072\u0065\u0064", "\u0063o\u0070\u0079\u0072\u0069\u0067\u0068t", "\u0074r\u0061\u0064\u0065\u006d\u0061\u0072k", "\u0061\u0063\u0075t\u0065", "\u0064\u0069\u0065\u0072\u0065\u0073\u0069\u0073", "\u006e\u006f\u0074\u0065\u0071\u0075\u0061\u006c", "\u0041\u0045", "\u004f\u0073\u006c\u0061\u0073\u0068", "\u0069\u006e\u0066\u0069\u006e\u0069\u0074\u0079", "\u0070l\u0075\u0073\u006d\u0069\u006e\u0075s", "\u006ce\u0073\u0073\u0065\u0071\u0075\u0061l", "\u0067\u0072\u0065a\u0074\u0065\u0072\u0065\u0071\u0075\u0061\u006c", "\u0079\u0065\u006e", "\u006d\u0075", "p\u0061\u0072\u0074\u0069\u0061\u006c\u0064\u0069\u0066\u0066", "\u0073u\u006d\u006d\u0061\u0074\u0069\u006fn", "\u0070r\u006f\u0064\u0075\u0063\u0074", "\u0070\u0069", "\u0069\u006e\u0074\u0065\u0067\u0072\u0061\u006c", "o\u0072\u0064\u0066\u0065\u006d\u0069\u006e\u0069\u006e\u0065", "\u006f\u0072\u0064m\u0061\u0073\u0063\u0075\u006c\u0069\u006e\u0065", "\u004f\u006d\u0065g\u0061", "\u0061\u0065", "\u006f\u0073\u006c\u0061\u0073\u0068", "\u0071\u0075\u0065s\u0074\u0069\u006f\u006e\u0064\u006f\u0077\u006e", "\u0065\u0078\u0063\u006c\u0061\u006d\u0064\u006f\u0077\u006e", "\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u006e\u006f\u0074", "\u0072a\u0064\u0069\u0063\u0061\u006c", "\u0066\u006c\u006f\u0072\u0069\u006e", "a\u0070\u0070\u0072\u006f\u0078\u0065\u0071\u0075\u0061\u006c", "\u0044\u0065\u006ct\u0061", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u006c\u0065\u006d\u006f\u0074r\u0069\u0067\u0068\u0074", "\u0065\u006c\u006c\u0069\u0070\u0073\u0069\u0073", "\u006e\u006fn\u0062\u0072\u0065a\u006b\u0069\u006e\u0067\u0073\u0070\u0061\u0063\u0065", "\u0041\u0067\u0072\u0061\u0076\u0065", "\u0041\u0074\u0069\u006c\u0064\u0065", "\u004f\u0074\u0069\u006c\u0064\u0065", "\u004f\u0045", "\u006f\u0065", "\u0065\u006e\u0064\u0061\u0073\u0068", "\u0065\u006d\u0064\u0061\u0073\u0068", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u006c\u0065\u0066\u0074", "\u0071\u0075\u006f\u0074\u0065\u0064\u0062\u006c\u0072\u0069\u0067\u0068\u0074", "\u0071u\u006f\u0074\u0065\u006c\u0065\u0066t", "\u0071\u0075\u006f\u0074\u0065\u0072\u0069\u0067\u0068\u0074", "\u0064\u0069\u0076\u0069\u0064\u0065", "\u006co\u007a\u0065\u006e\u0067\u0065", "\u0079d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0059d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e", "\u0063\u0075\u0072\u0072\u0065\u006e\u0063\u0079", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006c\u006c\u0065\u0066\u0074", "\u0067\u0075\u0069\u006c\u0073\u0069\u006e\u0067\u006cr\u0069\u0067\u0068\u0074", "\u0066\u0069", "\u0066\u006c", "\u0064a\u0067\u0067\u0065\u0072\u0064\u0062l", "\u0070\u0065\u0072\u0069\u006f\u0064\u0063\u0065\u006et\u0065\u0072\u0065\u0064", "\u0071\u0075\u006f\u0074\u0065\u0073\u0069\u006e\u0067l\u0062\u0061\u0073\u0065", "\u0071\u0075\u006ft\u0065\u0064\u0062\u006c\u0062\u0061\u0073\u0065", "p\u0065\u0072\u0074\u0068\u006f\u0075\u0073\u0061\u006e\u0064", "A\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "E\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0041\u0061\u0063\u0075\u0074\u0065", "\u0045d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0045\u0067\u0072\u0061\u0076\u0065", "\u0049\u0061\u0063\u0075\u0074\u0065", "I\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0049d\u0069\u0065\u0072\u0065\u0073\u0069s", "\u0049\u0067\u0072\u0061\u0076\u0065", "\u004f\u0061\u0063\u0075\u0074\u0065", "O\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0061\u0070\u0070l\u0065", "\u004f\u0067\u0072\u0061\u0076\u0065", "\u0055\u0061\u0063\u0075\u0074\u0065", "U\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0055\u0067\u0072\u0061\u0076\u0065", "\u0064\u006f\u0074\u006c\u0065\u0073\u0073\u0069", "\u0063\u0069\u0072\u0063\u0075\u006d\u0066\u006c\u0065\u0078", "\u0074\u0069\u006cd\u0065", "\u006d\u0061\u0063\u0072\u006f\u006e", "\u0062\u0072\u0065v\u0065", "\u0064o\u0074\u0061\u0063\u0063\u0065\u006et", "\u0072\u0069\u006e\u0067", "\u0063e\u0064\u0069\u006c\u006c\u0061", "\u0068\u0075\u006eg\u0061\u0072\u0075\u006d\u006c\u0061\u0075\u0074", "\u006f\u0067\u006f\u006e\u0065\u006b", "\u0063\u0061\u0072o\u006e", "\u004c\u0073\u006c\u0061\u0073\u0068", "\u006c\u0073\u006c\u0061\u0073\u0068", "\u0053\u0063\u0061\u0072\u006f\u006e", "\u0073\u0063\u0061\u0072\u006f\u006e", "\u005a\u0063\u0061\u0072\u006f\u006e", "\u007a\u0063\u0061\u0072\u006f\u006e", "\u0062r\u006f\u006b\u0065\u006e\u0062\u0061r", "\u0045\u0074\u0068", "\u0065\u0074\u0068", "\u0059\u0061\u0063\u0075\u0074\u0065", "\u0079\u0061\u0063\u0075\u0074\u0065", "\u0054\u0068\u006fr\u006e", "\u0074\u0068\u006fr\u006e", "\u006d\u0069\u006eu\u0073", "\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "o\u006e\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "t\u0077\u006f\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u0074\u0068\u0072\u0065\u0065\u0073\u0075\u0070\u0065\u0072\u0069\u006f\u0072", "\u006fn\u0065\u0068\u0061\u006c\u0066", "\u006f\u006e\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072", "\u0074\u0068\u0072\u0065\u0065\u0071\u0075\u0061\u0072\u0074\u0065\u0072\u0073", "\u0066\u0072\u0061n\u0063", "\u0047\u0062\u0072\u0065\u0076\u0065", "\u0067\u0062\u0072\u0065\u0076\u0065", "\u0049\u0064\u006f\u0074\u0061\u0063\u0063\u0065\u006e\u0074", "\u0053\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0073\u0063\u0065\u0064\u0069\u006c\u006c\u0061", "\u0043\u0061\u0063\u0075\u0074\u0065", "\u0063\u0061\u0063\u0075\u0074\u0065", "\u0043\u0063\u0061\u0072\u006f\u006e", "\u0063\u0063\u0061\u0072\u006f\u006e", "\u0064\u0063\u0072\u006f\u0061\u0074"}

func _dd() StdFont {
	_fdd.Do(_gcdb)
	_cdd := Descriptor{Name: CourierBoldObliqueName, Family: string(CourierName), Weight: FontWeightBold, Flags: 0x0061, BBox: [4]float64{-57, -250, 869, 801}, ItalicAngle: -12, Ascent: 629, Descent: -157, CapHeight: 562, XHeight: 439, StemV: 106, StemH: 84}
	return NewStdFont(_cdd, _beg)
}
func (_cdb *fontMap) write(_ab StdFontName, _ce func() StdFont) {
	_cdb.Lock()
	defer _cdb.Unlock()
	_cdb._be[_ab] = _ce
}

var _fea *RuneCharSafeMap
var _ggaa = []int16{722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 556, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 389, 722, 722, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 889, 722, 722, 722, 722, 722, 722, 722, 722, 556, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 556, 722, 722, 722, 722, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 611, 611, 611, 611, 444, 444, 444, 444, 333, 444, 667, 444, 444, 778, 444, 444, 469, 541, 500, 921, 444, 500, 278, 200, 480, 480, 333, 333, 333, 200, 350, 444, 444, 333, 444, 444, 333, 500, 333, 278, 250, 250, 760, 500, 500, 500, 500, 588, 500, 400, 333, 564, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 564, 500, 333, 333, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 564, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 344, 278, 564, 549, 564, 471, 278, 778, 333, 564, 500, 564, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 453, 333, 333, 476, 833, 250, 250, 1000, 564, 564, 500, 444, 444, 408, 444, 444, 444, 333, 333, 333, 180, 333, 333, 453, 333, 333, 760, 333, 389, 389, 389, 389, 389, 500, 278, 500, 500, 278, 250, 500, 600, 278, 326, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 722, 500, 500, 500, 500, 500, 444, 444, 444, 444, 500}

func _cbef() {
	_fea = MakeRuneCharSafeMap(len(_ec))
	_ffa = MakeRuneCharSafeMap(len(_ec))
	_abf = MakeRuneCharSafeMap(len(_ec))
	_gcae = MakeRuneCharSafeMap(len(_ec))
	for _ebf, _cgf := range _ec {
		_fea.Write(_cgf, CharMetrics{Wx: float64(_ggaa[_ebf])})
		_ffa.Write(_cgf, CharMetrics{Wx: float64(_edc[_ebf])})
		_abf.Write(_cgf, CharMetrics{Wx: float64(_bfd[_ebf])})
		_gcae.Write(_cgf, CharMetrics{Wx: float64(_efb[_ebf])})
	}
}

type GlyphName = _e.GlyphName
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

func (_dcbc *ttfParser) ParseOS2() error {
	if _dab := _dcbc.Seek("\u004f\u0053\u002f\u0032"); _dab != nil {
		return _dab
	}
	_bfe := _dcbc.ReadUShort()
	_dcbc.Skip(4 * 2)
	_dcbc.Skip(11*2 + 10 + 4*4 + 4)
	_febd := _dcbc.ReadUShort()
	_dcbc._cgc.Bold = (_febd & 32) != 0
	_dcbc.Skip(2 * 2)
	_dcbc._cgc.TypoAscender = _dcbc.ReadShort()
	_dcbc._cgc.TypoDescender = _dcbc.ReadShort()
	if _bfe >= 2 {
		_dcbc.Skip(3*2 + 2*4 + 2)
		_dcbc._cgc.CapHeight = _dcbc.ReadShort()
	} else {
		_dcbc._cgc.CapHeight = 0
	}
	return nil
}
func init() {
	RegisterStdFont(HelveticaName, _bf, "\u0041\u0072\u0069a\u006c")
	RegisterStdFont(HelveticaBoldName, _gcg, "\u0041\u0072\u0069\u0061\u006c\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(HelveticaObliqueName, _cda, "\u0041\u0072\u0069a\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(HelveticaBoldObliqueName, _cga, "\u0041\u0072i\u0061\u006c\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}

var _beg *RuneCharSafeMap

type CharMetrics struct {
	Wx float64
	Wy float64
}

func (_gdcg *ttfParser) parseCmapSubtable10(_bed int64) error {
	if _gdcg._cgc.Chars == nil {
		_gdcg._cgc.Chars = make(map[rune]GID)
	}
	_gdcg._cbee.Seek(int64(_gdcg._ebba["\u0063\u006d\u0061\u0070"])+_bed, _d.SeekStart)
	var _cddf, _gad uint32
	_cbfd := _gdcg.ReadUShort()
	if _cbfd < 8 {
		_cddf = uint32(_gdcg.ReadUShort())
		_gad = uint32(_gdcg.ReadUShort())
	} else {
		_gdcg.ReadUShort()
		_cddf = _gdcg.ReadULong()
		_gad = _gdcg.ReadULong()
	}
	_bd.Log.Trace("\u0070\u0061r\u0073\u0065\u0043\u006d\u0061p\u0053\u0075\u0062\u0074\u0061b\u006c\u0065\u0031\u0030\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065\u003d\u0025\u0064", _cbfd, _cddf, _gad)
	if _cbfd != 0 {
		return _dg.New("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006d\u0061p\u0020s\u0075\u0062\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	_gafc, _agbf := _gdcg.ReadStr(256)
	if _agbf != nil {
		return _agbf
	}
	_cbg := []byte(_gafc)
	for _gac, _cbge := range _cbg {
		_gdcg._cgc.Chars[rune(_gac)] = GID(_cbge)
		if _cbge != 0 {
			_g.Printf("\u0009\u0030\u0078\u002502\u0078\u0020\u279e\u0020\u0030\u0078\u0025\u0030\u0032\u0078\u003d\u0025\u0063\u000a", _gac, _cbge, rune(_cbge))
		}
	}
	return nil
}
func (_gcab *ttfParser) parseCmapSubtable31(_ecb int64) error {
	_ebd := make([]rune, 0, 8)
	_feg := make([]rune, 0, 8)
	_fgcg := make([]int16, 0, 8)
	_fga := make([]uint16, 0, 8)
	_gcab._cgc.Chars = make(map[rune]GID)
	_gcab._cbee.Seek(int64(_gcab._ebba["\u0063\u006d\u0061\u0070"])+_ecb, _d.SeekStart)
	_agbe := _gcab.ReadUShort()
	if _agbe != 4 {
		return _a.Errorf("u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0073\u0075\u0062t\u0061\u0062\u006c\u0065\u0020\u0066\u006fr\u006d\u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0028\u0025w\u0029", _agbe, _af.ErrNotSupported)
	}
	_gcab.Skip(2 * 2)
	_agf := int(_gcab.ReadUShort() / 2)
	_gcab.Skip(3 * 2)
	for _cff := 0; _cff < _agf; _cff++ {
		_feg = append(_feg, rune(_gcab.ReadUShort()))
	}
	_gcab.Skip(2)
	for _gae := 0; _gae < _agf; _gae++ {
		_ebd = append(_ebd, rune(_gcab.ReadUShort()))
	}
	for _gba := 0; _gba < _agf; _gba++ {
		_fgcg = append(_fgcg, _gcab.ReadShort())
	}
	_cbb, _ := _gcab._cbee.Seek(int64(0), _d.SeekCurrent)
	for _gec := 0; _gec < _agf; _gec++ {
		_fga = append(_fga, _gcab.ReadUShort())
	}
	for _fef := 0; _fef < _agf; _fef++ {
		_bdd := _ebd[_fef]
		_aab := _feg[_fef]
		_befd := _fgcg[_fef]
		_fece := _fga[_fef]
		if _fece > 0 {
			_gcab._cbee.Seek(_cbb+2*int64(_fef)+int64(_fece), _d.SeekStart)
		}
		for _agg := _bdd; _agg <= _aab; _agg++ {
			if _agg == 0xFFFF {
				break
			}
			var _dae int32
			if _fece > 0 {
				_dae = int32(_gcab.ReadUShort())
				if _dae > 0 {
					_dae += int32(_befd)
				}
			} else {
				_dae = _agg + int32(_befd)
			}
			if _dae >= 65536 {
				_dae -= 65536
			}
			if _dae > 0 {
				_gcab._cgc.Chars[_agg] = GID(_dae)
			}
		}
	}
	return nil
}

var _bfd = []int16{667, 944, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 722, 722, 612, 667, 667, 667, 667, 667, 667, 667, 667, 667, 722, 500, 667, 722, 722, 722, 778, 389, 389, 389, 389, 389, 389, 389, 389, 500, 667, 667, 611, 611, 611, 611, 611, 889, 722, 722, 722, 722, 722, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 667, 667, 667, 667, 556, 556, 556, 556, 556, 611, 611, 611, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 667, 889, 667, 611, 611, 611, 611, 611, 611, 611, 500, 500, 500, 500, 333, 500, 722, 500, 500, 778, 500, 500, 570, 570, 500, 832, 500, 500, 278, 220, 348, 348, 333, 333, 333, 220, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 747, 500, 500, 500, 500, 608, 500, 400, 333, 570, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 1000, 444, 1000, 500, 444, 570, 500, 389, 389, 333, 556, 500, 556, 500, 500, 167, 500, 500, 500, 500, 333, 570, 549, 500, 500, 333, 333, 556, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 500, 500, 278, 278, 382, 278, 570, 549, 606, 494, 278, 778, 333, 606, 576, 570, 556, 556, 556, 556, 500, 549, 556, 500, 500, 500, 500, 500, 722, 333, 500, 500, 500, 500, 750, 750, 300, 266, 300, 500, 500, 500, 500, 333, 333, 494, 833, 250, 250, 1000, 570, 570, 500, 500, 500, 555, 500, 500, 500, 333, 333, 333, 278, 389, 389, 549, 389, 389, 747, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 366, 278, 500, 500, 750, 300, 333, 1000, 500, 300, 556, 556, 556, 556, 556, 556, 556, 500, 556, 556, 444, 667, 500, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func _febe() StdFont {
	_daa.Do(_cbef)
	_dde := Descriptor{Name: TimesBoldItalicName, Family: _ccd, Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-200, -218, 996, 921}, ItalicAngle: -15, Ascent: 683, Descent: -217, CapHeight: 669, XHeight: 462, StemV: 121, StemH: 42}
	return NewStdFont(_dde, _abf)
}
func (_df *RuneCharSafeMap) Length() int {
	_df._gdc.RLock()
	defer _df._gdc.RUnlock()
	return len(_df._cc)
}
func (_febc *TtfType) NewEncoder() _e.TextEncoder { return _e.NewTrueTypeFontEncoder(_febc.Chars) }

type Font interface {
	Encoder() _e.TextEncoder
	GetRuneMetrics(_gg rune) (CharMetrics, bool)
}

func (_gff *ttfParser) parseCmapVersion(_bgg int64) error {
	_bd.Log.Trace("p\u0061\u0072\u0073\u0065\u0043\u006da\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a \u006f\u0066\u0066s\u0065t\u003d\u0025\u0064", _bgg)
	if _gff._cgc.Chars == nil {
		_gff._cgc.Chars = make(map[rune]GID)
	}
	_gff._cbee.Seek(int64(_gff._ebba["\u0063\u006d\u0061\u0070"])+_bgg, _d.SeekStart)
	var _fee, _bdea uint32
	_ccf := _gff.ReadUShort()
	if _ccf < 8 {
		_fee = uint32(_gff.ReadUShort())
		_bdea = uint32(_gff.ReadUShort())
	} else {
		_gff.ReadUShort()
		_fee = _gff.ReadULong()
		_bdea = _gff.ReadULong()
	}
	_bd.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0043m\u0061\u0070\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u003a\u0020\u0066\u006f\u0072\u006d\u0061\u0074\u003d\u0025\u0064 \u006c\u0065\u006e\u0067\u0074\u0068\u003d\u0025\u0064\u0020\u006c\u0061\u006e\u0067u\u0061g\u0065\u003d\u0025\u0064", _ccf, _fee, _bdea)
	switch _ccf {
	case 0:
		return _gff.parseCmapFormat0()
	case 6:
		return _gff.parseCmapFormat6()
	case 12:
		return _gff.parseCmapFormat12()
	default:
		_bd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063m\u0061\u0070\u0020\u0066\u006f\u0072\u006da\u0074\u003d\u0025\u0064", _ccf)
		return nil
	}
}
func _cda() StdFont {
	_aaf.Do(_cgag)
	_bgd := Descriptor{Name: HelveticaObliqueName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-170, -225, 1116, 931}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_bgd, _gde)
}
func _bf() StdFont {
	_aaf.Do(_cgag)
	_dgg := Descriptor{Name: HelveticaName, Family: string(HelveticaName), Weight: FontWeightMedium, Flags: 0x0020, BBox: [4]float64{-166, -225, 1000, 931}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 523, StemV: 88, StemH: 76}
	return NewStdFont(_dgg, _bef)
}
func (_cfd *ttfParser) ParseName() error {
	if _ebff := _cfd.Seek("\u006e\u0061\u006d\u0065"); _ebff != nil {
		return _ebff
	}
	_edg, _ := _cfd._cbee.Seek(0, _d.SeekCurrent)
	_cfd._cgc.PostScriptName = ""
	_cfd.Skip(2)
	_cdc := _cfd.ReadUShort()
	_ggd := _cfd.ReadUShort()
	for _bdec := uint16(0); _bdec < _cdc && _cfd._cgc.PostScriptName == ""; _bdec++ {
		_cfd.Skip(3 * 2)
		_faa := _cfd.ReadUShort()
		_ccba := _cfd.ReadUShort()
		_cfdg := _cfd.ReadUShort()
		if _faa == 6 {
			_cfd._cbee.Seek(_edg+int64(_ggd)+int64(_cfdg), _d.SeekStart)
			_beb, _gdccb := _cfd.ReadStr(int(_ccba))
			if _gdccb != nil {
				return _gdccb
			}
			_beb = _ga.Replace(_beb, "\u0000", "", -1)
			_bdc, _gdccb := _fb.Compile("\u005b\u0028\u0029\u007b\u007d\u003c\u003e\u0020\u002f%\u005b\u005c\u005d\u005d")
			if _gdccb != nil {
				return _gdccb
			}
			_cfd._cgc.PostScriptName = _bdc.ReplaceAllString(_beb, "")
		}
	}
	if _cfd._cgc.PostScriptName == "" {
		_bd.Log.Debug("\u0050a\u0072\u0073e\u004e\u0061\u006de\u003a\u0020\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u0050\u006f\u0073t\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0077\u0061\u0073\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
	}
	return nil
}
func init() {
	RegisterStdFont(SymbolName, _cfg, "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0049\u0074\u0061\u006c\u0069\u0063", "S\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064", "\u0053\u0079\u006d\u0062\u006f\u006c\u002c\u0042\u006f\u006c\u0064\u0049t\u0061\u006c\u0069\u0063")
	RegisterStdFont(ZapfDingbatsName, _cead)
}
func (_eba *ttfParser) parseCmapFormat12() error {
	_gcf := _eba.ReadULong()
	_bd.Log.Trace("\u0070\u0061\u0072se\u0043\u006d\u0061\u0070\u0046\u006f\u0072\u006d\u0061t\u00312\u003a \u0025s\u0020\u006e\u0075\u006d\u0047\u0072\u006f\u0075\u0070\u0073\u003d\u0025\u0064", _eba._cgc.String(), _gcf)
	for _fda := uint32(0); _fda < _gcf; _fda++ {
		_febb := _eba.ReadULong()
		_gfcg := _eba.ReadULong()
		_cgcd := _eba.ReadULong()
		if _febb > 0x0010FFFF || (0xD800 <= _febb && _febb <= 0xDFFF) {
			return _dg.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		if _gfcg < _febb || _gfcg > 0x0010FFFF || (0xD800 <= _gfcg && _gfcg <= 0xDFFF) {
			return _dg.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0068\u0061\u0072\u0061c\u0074\u0065\u0072\u0073\u0020\u0063\u006f\u0064\u0065\u0073")
		}
		for _gce := _febb; _gce <= _gfcg; _gce++ {
			if _gce > 0x10FFFF {
				_bd.Log.Debug("\u0046\u006fr\u006d\u0061\u0074\u0020\u0031\u0032\u0020\u0063\u006d\u0061\u0070\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0062\u0065\u0079\u006f\u006e\u0064\u0020\u0055\u0043\u0053\u002d\u0034")
			}
			_eba._cgc.Chars[rune(_gce)] = GID(_cgcd)
			_cgcd++
		}
	}
	return nil
}
func (_fe CharMetrics) String() string {
	return _g.Sprintf("<\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u003e", _fe.Wx, _fe.Wy)
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

var _ffe *RuneCharSafeMap

func IsStdFont(name StdFontName) bool {
	_, _eb := _afd.read(name)
	return _eb
}
func init() {
	RegisterStdFont(TimesRomanName, _gedc, "\u0054\u0069\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e", "\u0054\u0069\u006de\u0073")
	RegisterStdFont(TimesBoldName, _eae, "\u0054i\u006de\u0073\u004e\u0065\u0077\u0052o\u006d\u0061n\u002c\u0042\u006f\u006c\u0064", "\u0054\u0069\u006d\u0065\u0073\u002c\u0042\u006f\u006c\u0064")
	RegisterStdFont(TimesItalicName, _agb, "T\u0069m\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006da\u006e\u002c\u0049\u0074al\u0069\u0063", "\u0054\u0069\u006de\u0073\u002c\u0049\u0074\u0061\u006c\u0069\u0063")
	RegisterStdFont(TimesBoldItalicName, _febe, "\u0054i\u006d\u0065\u0073\u004e\u0065\u0077\u0052\u006f\u006d\u0061\u006e,\u0042\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063", "\u0054\u0069m\u0065\u0073\u002cB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
}
func (_efdb *ttfParser) Read32Fixed() float64 {
	_gecba := float64(_efdb.ReadShort())
	_cecb := float64(_efdb.ReadUShort()) / 65536.0
	return _gecba + _cecb
}

var _afd = &fontMap{_be: make(map[StdFontName]func() StdFont)}
var _ged *RuneCharSafeMap

func (_ege *ttfParser) ReadULong() (_fdc uint32) {
	_gd.Read(_ege._cbee, _gd.BigEndian, &_fdc)
	return _fdc
}
func _gcg() StdFont {
	_aaf.Do(_cgag)
	_gfe := Descriptor{Name: HelveticaBoldName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-170, -228, 1003, 962}, ItalicAngle: 0, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_gfe, _ffe)
}

var _abf *RuneCharSafeMap

func (_ecbc *ttfParser) parseCmapFormat0() error {
	_adc, _bbe := _ecbc.ReadStr(256)
	if _bbe != nil {
		return _bbe
	}
	_dgae := []byte(_adc)
	_bd.Log.Trace("\u0070a\u0072\u0073e\u0043\u006d\u0061p\u0046\u006f\u0072\u006d\u0061\u0074\u0030:\u0020\u0025\u0073\u000a\u0064\u0061t\u0061\u0053\u0074\u0072\u003d\u0025\u002b\u0071\u000a\u0064\u0061t\u0061\u003d\u005b\u0025\u0020\u0030\u0032\u0078\u005d", _ecbc._cgc.String(), _adc, _dgae)
	for _fac, _bfcf := range _dgae {
		_ecbc._cgc.Chars[rune(_fac)] = GID(_bfcf)
	}
	return nil
}
func (_cb StdFont) Name() string { return string(_cb._ea.Name) }
func RegisterStdFont(name StdFontName, fnc func() StdFont, aliases ...StdFontName) {
	if _, _edf := _afd.read(name); _edf {
		panic("\u0066o\u006e\u0074\u0020\u0061l\u0072\u0065\u0061\u0064\u0079 \u0072e\u0067i\u0073\u0074\u0065\u0072\u0065\u0064\u003a " + string(name))
	}
	_afd.write(name, fnc)
	for _, _gcd := range aliases {
		RegisterStdFont(_gcd, fnc)
	}
}

var _ Font = StdFont{}

func _cead() StdFont {
	_eeb := _e.NewZapfDingbatsEncoder()
	_daf := Descriptor{Name: ZapfDingbatsName, Family: string(ZapfDingbatsName), Weight: FontWeightMedium, Flags: 0x0004, BBox: [4]float64{-1, -143, 981, 820}, ItalicAngle: 0, Ascent: 0, Descent: 0, CapHeight: 0, XHeight: 0, StemV: 90, StemH: 28}
	return NewStdFontWithEncoding(_daf, _ddc, _eeb)
}
func (_bbf *ttfParser) ReadShort() (_adg int16) {
	_gd.Read(_bbf._cbee, _gd.BigEndian, &_adg)
	return _adg
}
func (_fgf *ttfParser) readByte() (_afda uint8) {
	_gd.Read(_fgf._cbee, _gd.BigEndian, &_afda)
	return _afda
}
func _cgag() {
	_bef = MakeRuneCharSafeMap(len(_ec))
	_ffe = MakeRuneCharSafeMap(len(_ec))
	for _ef, _cefd := range _ec {
		_bef.Write(_cefd, CharMetrics{Wx: float64(_gb[_ef])})
		_ffe.Write(_cefd, CharMetrics{Wx: float64(_baga[_ef])})
	}
	_gde = _bef.Copy()
	_ged = _ffe.Copy()
}
func (_abc *ttfParser) ReadUShort() (_eabe uint16) {
	_gd.Read(_abc._cbee, _gd.BigEndian, &_eabe)
	return _eabe
}
func (_fce *ttfParser) ParseCmap() error {
	var _acd int64
	if _abeg := _fce.Seek("\u0063\u006d\u0061\u0070"); _abeg != nil {
		return _abeg
	}
	_bd.Log.Trace("\u0050a\u0072\u0073\u0065\u0043\u006d\u0061p")
	_fce.ReadUShort()
	_dee := int(_fce.ReadUShort())
	_fegc := int64(0)
	_gafd := int64(0)
	_fed := int64(0)
	for _badb := 0; _badb < _dee; _badb++ {
		_eaec := _fce.ReadUShort()
		_dag := _fce.ReadUShort()
		_acd = int64(_fce.ReadULong())
		if _eaec == 3 && _dag == 1 {
			_gafd = _acd
		} else if _eaec == 3 && _dag == 10 {
			_fed = _acd
		} else if _eaec == 1 && _dag == 0 {
			_fegc = _acd
		}
	}
	if _fegc != 0 {
		if _cce := _fce.parseCmapVersion(_fegc); _cce != nil {
			return _cce
		}
	}
	if _gafd != 0 {
		if _gecb := _fce.parseCmapSubtable31(_gafd); _gecb != nil {
			return _gecb
		}
	}
	if _fed != 0 {
		if _aba := _fce.parseCmapVersion(_fed); _aba != nil {
			return _aba
		}
	}
	if _gafd == 0 && _fegc == 0 && _fed == 0 {
		_bd.Log.Debug("\u0074\u0074\u0066P\u0061\u0072\u0073\u0065\u0072\u002e\u0050\u0061\u0072\u0073\u0065\u0043\u006d\u0061\u0070\u002e\u0020\u004e\u006f\u0020\u0033\u0031\u002c\u0020\u0031\u0030\u002c\u0020\u00331\u0030\u0020\u0074\u0061\u0062\u006c\u0065\u002e")
	}
	return nil
}

var _dff = &RuneCharSafeMap{_cc: map[rune]CharMetrics{' ': {Wx: 250}, '!': {Wx: 333}, '#': {Wx: 500}, '%': {Wx: 833}, '&': {Wx: 778}, '(': {Wx: 333}, ')': {Wx: 333}, '+': {Wx: 549}, ',': {Wx: 250}, '.': {Wx: 250}, '/': {Wx: 278}, '0': {Wx: 500}, '1': {Wx: 500}, '2': {Wx: 500}, '3': {Wx: 500}, '4': {Wx: 500}, '5': {Wx: 500}, '6': {Wx: 500}, '7': {Wx: 500}, '8': {Wx: 500}, '9': {Wx: 500}, ':': {Wx: 278}, ';': {Wx: 278}, '<': {Wx: 549}, '=': {Wx: 549}, '>': {Wx: 549}, '?': {Wx: 444}, '[': {Wx: 333}, ']': {Wx: 333}, '_': {Wx: 500}, '{': {Wx: 480}, '|': {Wx: 200}, '}': {Wx: 480}, '¬': {Wx: 713}, '°': {Wx: 400}, '±': {Wx: 549}, 'µ': {Wx: 576}, '×': {Wx: 549}, '÷': {Wx: 549}, 'ƒ': {Wx: 500}, 'Α': {Wx: 722}, 'Β': {Wx: 667}, 'Γ': {Wx: 603}, 'Ε': {Wx: 611}, 'Ζ': {Wx: 611}, 'Η': {Wx: 722}, 'Θ': {Wx: 741}, 'Ι': {Wx: 333}, 'Κ': {Wx: 722}, 'Λ': {Wx: 686}, 'Μ': {Wx: 889}, 'Ν': {Wx: 722}, 'Ξ': {Wx: 645}, 'Ο': {Wx: 722}, 'Π': {Wx: 768}, 'Ρ': {Wx: 556}, 'Σ': {Wx: 592}, 'Τ': {Wx: 611}, 'Υ': {Wx: 690}, 'Φ': {Wx: 763}, 'Χ': {Wx: 722}, 'Ψ': {Wx: 795}, 'α': {Wx: 631}, 'β': {Wx: 549}, 'γ': {Wx: 411}, 'δ': {Wx: 494}, 'ε': {Wx: 439}, 'ζ': {Wx: 494}, 'η': {Wx: 603}, 'θ': {Wx: 521}, 'ι': {Wx: 329}, 'κ': {Wx: 549}, 'λ': {Wx: 549}, 'ν': {Wx: 521}, 'ξ': {Wx: 493}, 'ο': {Wx: 549}, 'π': {Wx: 549}, 'ρ': {Wx: 549}, 'ς': {Wx: 439}, 'σ': {Wx: 603}, 'τ': {Wx: 439}, 'υ': {Wx: 576}, 'φ': {Wx: 521}, 'χ': {Wx: 549}, 'ψ': {Wx: 686}, 'ω': {Wx: 686}, 'ϑ': {Wx: 631}, 'ϒ': {Wx: 620}, 'ϕ': {Wx: 603}, 'ϖ': {Wx: 713}, '•': {Wx: 460}, '…': {Wx: 1000}, '′': {Wx: 247}, '″': {Wx: 411}, '⁄': {Wx: 167}, '€': {Wx: 750}, 'ℑ': {Wx: 686}, '℘': {Wx: 987}, 'ℜ': {Wx: 795}, 'Ω': {Wx: 768}, 'ℵ': {Wx: 823}, '←': {Wx: 987}, '↑': {Wx: 603}, '→': {Wx: 987}, '↓': {Wx: 603}, '↔': {Wx: 1042}, '↵': {Wx: 658}, '⇐': {Wx: 987}, '⇑': {Wx: 603}, '⇒': {Wx: 987}, '⇓': {Wx: 603}, '⇔': {Wx: 1042}, '∀': {Wx: 713}, '∂': {Wx: 494}, '∃': {Wx: 549}, '∅': {Wx: 823}, '∆': {Wx: 612}, '∇': {Wx: 713}, '∈': {Wx: 713}, '∉': {Wx: 713}, '∋': {Wx: 439}, '∏': {Wx: 823}, '∑': {Wx: 713}, '−': {Wx: 549}, '∗': {Wx: 500}, '√': {Wx: 549}, '∝': {Wx: 713}, '∞': {Wx: 713}, '∠': {Wx: 768}, '∧': {Wx: 603}, '∨': {Wx: 603}, '∩': {Wx: 768}, '∪': {Wx: 768}, '∫': {Wx: 274}, '∴': {Wx: 863}, '∼': {Wx: 549}, '≅': {Wx: 549}, '≈': {Wx: 549}, '≠': {Wx: 549}, '≡': {Wx: 549}, '≤': {Wx: 549}, '≥': {Wx: 549}, '⊂': {Wx: 713}, '⊃': {Wx: 713}, '⊄': {Wx: 713}, '⊆': {Wx: 713}, '⊇': {Wx: 713}, '⊕': {Wx: 768}, '⊗': {Wx: 768}, '⊥': {Wx: 658}, '⋅': {Wx: 250}, '⌠': {Wx: 686}, '⌡': {Wx: 686}, '〈': {Wx: 329}, '〉': {Wx: 329}, '◊': {Wx: 494}, '♠': {Wx: 753}, '♣': {Wx: 753}, '♥': {Wx: 753}, '♦': {Wx: 753}, '\uf6d9': {Wx: 790}, '\uf6da': {Wx: 790}, '\uf6db': {Wx: 890}, '\uf8e5': {Wx: 500}, '\uf8e6': {Wx: 603}, '\uf8e7': {Wx: 1000}, '\uf8e8': {Wx: 790}, '\uf8e9': {Wx: 790}, '\uf8ea': {Wx: 786}, '\uf8eb': {Wx: 384}, '\uf8ec': {Wx: 384}, '\uf8ed': {Wx: 384}, '\uf8ee': {Wx: 384}, '\uf8ef': {Wx: 384}, '\uf8f0': {Wx: 384}, '\uf8f1': {Wx: 494}, '\uf8f2': {Wx: 494}, '\uf8f3': {Wx: 494}, '\uf8f4': {Wx: 494}, '\uf8f5': {Wx: 686}, '\uf8f6': {Wx: 384}, '\uf8f7': {Wx: 384}, '\uf8f8': {Wx: 384}, '\uf8f9': {Wx: 384}, '\uf8fa': {Wx: 384}, '\uf8fb': {Wx: 384}, '\uf8fc': {Wx: 494}, '\uf8fd': {Wx: 494}, '\uf8fe': {Wx: 494}, '\uf8ff': {Wx: 790}}}

type RuneCharSafeMap struct {
	_cc  map[rune]CharMetrics
	_gdc _fba.RWMutex
}

func (_efd *ttfParser) ParseMaxp() error {
	if _fc := _efd.Seek("\u006d\u0061\u0078\u0070"); _fc != nil {
		return _fc
	}
	_efd.Skip(4)
	_efd._gbg = _efd.ReadUShort()
	return nil
}
func (_gdcc *TtfType) MakeEncoder() (_e.SimpleEncoder, error) {
	_dea := make(map[_e.CharCode]GlyphName)
	for _cbeff := _e.CharCode(0); _cbeff <= 256; _cbeff++ {
		_egg := rune(_cbeff)
		_dbb, _bae := _gdcc.Chars[_egg]
		if !_bae {
			continue
		}
		var _ddg GlyphName
		if int(_dbb) >= 0 && int(_dbb) < len(_gdcc.GlyphNames) {
			_ddg = _gdcc.GlyphNames[_dbb]
		} else {
			_dc := rune(_dbb)
			if _gdf, _ade := _e.RuneToGlyph(_dc); _ade {
				_ddg = _gdf
			}
		}
		if _ddg != "" {
			_dea[_cbeff] = _ddg
		}
	}
	if len(_dea) == 0 {
		_bd.Log.Debug("WA\u0052\u004eI\u004e\u0047\u003a\u0020\u005a\u0065\u0072\u006f\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002e\u0020\u0074\u0074\u0066=\u0025s\u0020\u0043\u0068\u0061\u0072\u0073\u003d\u005b%\u00200\u0032\u0078]", _gdcc, _gdcc.Chars)
	}
	return _e.NewCustomSimpleTextEncoder(_dea, nil)
}
func MakeRuneCharSafeMap(length int) *RuneCharSafeMap {
	return &RuneCharSafeMap{_cc: make(map[rune]CharMetrics, length)}
}
func (_gga *RuneCharSafeMap) Range(f func(_fbd rune, _dga CharMetrics) (_afg bool)) {
	_gga._gdc.RLock()
	defer _gga._gdc.RUnlock()
	for _gc, _eda := range _gga._cc {
		if f(_gc, _eda) {
			break
		}
	}
}

var _afc *RuneCharSafeMap

func TtfParseFile(fileStr string) (TtfType, error) {
	_gbe, _bac := _ff.Open(fileStr)
	if _bac != nil {
		return TtfType{}, _bac
	}
	defer _gbe.Close()
	return TtfParse(_gbe)
}
func (_bfcd *ttfParser) ParsePost() error {
	if _caca := _bfcd.Seek("\u0070\u006f\u0073\u0074"); _caca != nil {
		return _caca
	}
	_gbd := _bfcd.Read32Fixed()
	_bfcd._cgc.ItalicAngle = _bfcd.Read32Fixed()
	_bfcd._cgc.UnderlinePosition = _bfcd.ReadShort()
	_bfcd._cgc.UnderlineThickness = _bfcd.ReadShort()
	_bfcd._cgc.IsFixedPitch = _bfcd.ReadULong() != 0
	_bfcd.ReadULong()
	_bfcd.ReadULong()
	_bfcd.ReadULong()
	_bfcd.ReadULong()
	_bd.Log.Trace("\u0050a\u0072\u0073\u0065\u0050\u006f\u0073\u0074\u003a\u0020\u0066\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065\u003d\u0025\u0066", _gbd)
	switch _gbd {
	case 1.0:
		_bfcd._cgc.GlyphNames = _gbdg
	case 2.0:
		_eac := int(_bfcd.ReadUShort())
		_edd := make([]int, _eac)
		_bfcd._cgc.GlyphNames = make([]GlyphName, _eac)
		_ddca := -1
		for _fcd := 0; _fcd < _eac; _fcd++ {
			_deg := int(_bfcd.ReadUShort())
			_edd[_fcd] = _deg
			if _deg <= 0x7fff && _deg > _ddca {
				_ddca = _deg
			}
		}
		var _ggdc []GlyphName
		if _ddca >= len(_gbdg) {
			_ggdc = make([]GlyphName, _ddca-len(_gbdg)+1)
			for _aea := 0; _aea < _ddca-len(_gbdg)+1; _aea++ {
				_dbg := int(_bfcd.readByte())
				_abega, _fcg := _bfcd.ReadStr(_dbg)
				if _fcg != nil {
					return _fcg
				}
				_ggdc[_aea] = GlyphName(_abega)
			}
		}
		for _fbdd := 0; _fbdd < _eac; _fbdd++ {
			_cdf := _edd[_fbdd]
			if _cdf < len(_gbdg) {
				_bfcd._cgc.GlyphNames[_fbdd] = _gbdg[_cdf]
			} else if _cdf >= len(_gbdg) && _cdf <= 32767 {
				_bfcd._cgc.GlyphNames[_fbdd] = _ggdc[_cdf-len(_gbdg)]
			} else {
				_bfcd._cgc.GlyphNames[_fbdd] = "\u002e\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064"
			}
		}
	case 2.5:
		_eea := make([]int, _bfcd._gbg)
		for _gaee := 0; _gaee < len(_eea); _gaee++ {
			_fdaa := int(_bfcd.ReadSByte())
			_eea[_gaee] = _gaee + 1 + _fdaa
		}
		_bfcd._cgc.GlyphNames = make([]GlyphName, len(_eea))
		for _dcf := 0; _dcf < len(_bfcd._cgc.GlyphNames); _dcf++ {
			_gbc := _gbdg[_eea[_dcf]]
			_bfcd._cgc.GlyphNames[_dcf] = _gbc
		}
	case 3.0:
		_bd.Log.Debug("\u004e\u006f\u0020\u0050\u006f\u0073t\u0053\u0063\u0072i\u0070\u0074\u0020n\u0061\u006d\u0065\u0020\u0069\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020is\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e")
	default:
		_bd.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020f\u006fr\u006d\u0061\u0074\u0054\u0079\u0070\u0065=\u0025\u0066", _gbd)
	}
	return nil
}
func (_fag *TtfType) MakeToUnicode() *_c.CMap {
	_gfd := make(map[_c.CharCode]rune)
	if len(_fag.GlyphNames) == 0 {
		return _c.NewToUnicodeCMap(_gfd)
	}
	for _dcb, _bde := range _fag.Chars {
		_eab := _c.CharCode(_dcb)
		_cddc := _fag.GlyphNames[_bde]
		_fgc, _cgd := _e.GlyphToRune(_cddc)
		if !_cgd {
			_bd.Log.Debug("\u004e\u006f \u0072\u0075\u006e\u0065\u002e\u0020\u0063\u006f\u0064\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u0020\u0067\u006c\u0079\u0070h=\u0025\u0071", _dcb, _cddc)
			_fgc = _e.MissingCodeRune
		}
		_gfd[_eab] = _fgc
	}
	return _c.NewToUnicodeCMap(_gfd)
}

var _daa _fba.Once
var _gcae *RuneCharSafeMap

func (_bad *ttfParser) ParseHead() error {
	if _ggga := _bad.Seek("\u0068\u0065\u0061\u0064"); _ggga != nil {
		return _ggga
	}
	_bad.Skip(3 * 4)
	_fead := _bad.ReadULong()
	if _fead != 0x5F0F3CF5 {
		_bd.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0049\u006e\u0063\u006fr\u0072e\u0063\u0074\u0020\u006d\u0061\u0067\u0069\u0063\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e\u0020\u0046\u006fn\u0074\u0020\u006d\u0061\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0069\u0073\u0070\u006c\u0061\u0079\u0020\u0063\u006f\u0072\u0072\u0065\u0063t\u006c\u0079\u002e\u0020\u0025\u0073", _bad)
	}
	_bad.Skip(2)
	_bad._cgc.UnitsPerEm = _bad.ReadUShort()
	_bad.Skip(2 * 8)
	_bad._cgc.Xmin = _bad.ReadShort()
	_bad._cgc.Ymin = _bad.ReadShort()
	_bad._cgc.Xmax = _bad.ReadShort()
	_bad._cgc.Ymax = _bad.ReadShort()
	return nil
}
func (_cbe StdFont) Encoder() _e.TextEncoder { return _cbe._gfb }

var _efb = []int16{611, 889, 611, 611, 611, 611, 611, 611, 611, 611, 611, 611, 667, 667, 667, 667, 722, 722, 722, 612, 611, 611, 611, 611, 611, 611, 611, 611, 611, 722, 500, 611, 722, 722, 722, 722, 333, 333, 333, 333, 333, 333, 333, 333, 444, 667, 667, 556, 556, 611, 556, 556, 833, 667, 667, 667, 667, 667, 722, 944, 722, 722, 722, 722, 722, 722, 722, 722, 611, 722, 611, 611, 611, 611, 500, 500, 500, 500, 500, 556, 556, 556, 611, 722, 722, 722, 722, 722, 722, 722, 722, 722, 611, 833, 611, 556, 556, 556, 556, 556, 556, 556, 500, 500, 500, 500, 333, 500, 667, 500, 500, 778, 500, 500, 422, 541, 500, 920, 500, 500, 278, 275, 400, 400, 389, 389, 333, 275, 350, 444, 444, 333, 444, 444, 333, 500, 333, 333, 250, 250, 760, 500, 500, 500, 500, 544, 500, 400, 333, 675, 500, 333, 278, 444, 444, 444, 444, 444, 444, 444, 500, 889, 444, 889, 500, 444, 675, 500, 333, 389, 278, 500, 500, 500, 500, 500, 167, 500, 500, 500, 500, 333, 675, 549, 500, 500, 333, 333, 500, 333, 333, 278, 278, 278, 278, 278, 278, 278, 278, 444, 444, 278, 278, 300, 278, 675, 549, 675, 471, 278, 722, 333, 675, 500, 675, 500, 500, 500, 500, 500, 549, 500, 500, 500, 500, 500, 500, 667, 333, 500, 500, 500, 500, 750, 750, 300, 276, 310, 500, 500, 500, 523, 333, 333, 476, 833, 250, 250, 1000, 675, 675, 500, 500, 500, 420, 556, 556, 556, 333, 333, 333, 214, 389, 389, 453, 389, 389, 760, 333, 389, 389, 389, 389, 389, 500, 333, 500, 500, 278, 250, 500, 600, 278, 300, 278, 500, 500, 750, 300, 333, 980, 500, 300, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 444, 667, 444, 444, 444, 444, 500, 389, 389, 389, 389, 500}

func NewStdFontWithEncoding(desc Descriptor, metrics *RuneCharSafeMap, encoder _e.TextEncoder) StdFont {
	var _cca rune = 0xA0
	if _, _fd := metrics.Read(_cca); !_fd {
		_feb, _ := metrics.Read(0x20)
		metrics.Write(_cca, _feb)
	}
	return StdFont{_ea: desc, _eec: metrics, _gfb: encoder}
}
func (_cdbc *ttfParser) Skip(n int) { _cdbc._cbee.Seek(int64(n), _d.SeekCurrent) }
func _gedc() StdFont {
	_daa.Do(_cbef)
	_baf := Descriptor{Name: TimesRomanName, Family: _ccd, Weight: FontWeightRoman, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 898}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 662, XHeight: 450, StemV: 84, StemH: 28}
	return NewStdFont(_baf, _fea)
}
func (_gdce *ttfParser) ParseComponents() error {
	if _bfc := _gdce.ParseHead(); _bfc != nil {
		return _bfc
	}
	if _ddcg := _gdce.ParseHhea(); _ddcg != nil {
		return _ddcg
	}
	if _abe := _gdce.ParseMaxp(); _abe != nil {
		return _abe
	}
	if _age := _gdce.ParseHmtx(); _age != nil {
		return _age
	}
	if _, _ece := _gdce._ebba["\u006e\u0061\u006d\u0065"]; _ece {
		if _fae := _gdce.ParseName(); _fae != nil {
			return _fae
		}
	}
	if _, _fbc := _gdce._ebba["\u004f\u0053\u002f\u0032"]; _fbc {
		if _eaed := _gdce.ParseOS2(); _eaed != nil {
			return _eaed
		}
	}
	if _, _ddgc := _gdce._ebba["\u0070\u006f\u0073\u0074"]; _ddgc {
		if _egd := _gdce.ParsePost(); _egd != nil {
			return _egd
		}
	}
	if _, _ceg := _gdce._ebba["\u0063\u006d\u0061\u0070"]; _ceg {
		if _cgg := _gdce.ParseCmap(); _cgg != nil {
			return _cgg
		}
	}
	return nil
}
func (_dfb *ttfParser) ParseHmtx() error {
	if _dded := _dfb.Seek("\u0068\u006d\u0074\u0078"); _dded != nil {
		return _dded
	}
	_dfb._cgc.Widths = make([]uint16, 0, 8)
	for _bgf := uint16(0); _bgf < _dfb._dbd; _bgf++ {
		_dfb._cgc.Widths = append(_dfb._cgc.Widths, _dfb.ReadUShort())
		_dfb.Skip(2)
	}
	if _dfb._dbd < _dfb._gbg && _dfb._dbd > 0 {
		_ecf := _dfb._cgc.Widths[_dfb._dbd-1]
		for _fdg := _dfb._dbd; _fdg < _dfb._gbg; _fdg++ {
			_dfb._cgc.Widths = append(_dfb._cgc.Widths, _ecf)
		}
	}
	return nil
}

var _bce *RuneCharSafeMap

func NewFontFile2FromPdfObject(obj _af.PdfObject) (TtfType, error) {
	obj = _af.TraceToDirectObject(obj)
	_bfg, _aaa := obj.(*_af.PdfObjectStream)
	if !_aaa {
		_bd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0046\u0069\u006c\u0065\u0032\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0073\u0074\u0072e\u0061\u006d \u0028\u0025\u0054\u0029", obj)
		return TtfType{}, _af.ErrTypeError
	}
	_cag, _afcb := _af.DecodeStream(_bfg)
	if _afcb != nil {
		return TtfType{}, _afcb
	}
	_aad := ttfParser{_cbee: _da.NewReader(_cag)}
	return _aad.Parse()
}
func NewStdFont(desc Descriptor, metrics *RuneCharSafeMap) StdFont {
	return NewStdFontWithEncoding(desc, metrics, _e.NewStandardEncoder())
}

const (
	_ccd                = "\u0054\u0069\u006de\u0073"
	TimesRomanName      = StdFontName("T\u0069\u006d\u0065\u0073\u002d\u0052\u006f\u006d\u0061\u006e")
	TimesBoldName       = StdFontName("\u0054\u0069\u006d\u0065\u0073\u002d\u0042\u006f\u006c\u0064")
	TimesItalicName     = StdFontName("\u0054\u0069\u006de\u0073\u002d\u0049\u0074\u0061\u006c\u0069\u0063")
	TimesBoldItalicName = StdFontName("\u0054\u0069m\u0065\u0073\u002dB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063")
)

type FontWeight int

func _cga() StdFont {
	_aaf.Do(_cgag)
	_gca := Descriptor{Name: HelveticaBoldObliqueName, Family: string(HelveticaName), Weight: FontWeightBold, Flags: 0x0060, BBox: [4]float64{-174, -228, 1114, 962}, ItalicAngle: -12, Ascent: 718, Descent: -207, CapHeight: 718, XHeight: 532, StemV: 140, StemH: 118}
	return NewStdFont(_gca, _ged)
}
func (_caa StdFont) Descriptor() Descriptor { return _caa._ea }

var _aaf _fba.Once

func _gcdb() {
	const _cg = 600
	_afc = MakeRuneCharSafeMap(len(_ec))
	for _, _db := range _ec {
		_afc.Write(_db, CharMetrics{Wx: _cg})
	}
	_afe = _afc.Copy()
	_beg = _afc.Copy()
	_bce = _afc.Copy()
}

type ttfParser struct {
	_cgc  TtfType
	_cbee _d.ReadSeeker
	_ebba map[string]uint32
	_dbd  uint16
	_gbg  uint16
}

const (
	CourierName            = StdFontName("\u0043o\u0075\u0072\u0069\u0065\u0072")
	CourierBoldName        = StdFontName("\u0043\u006f\u0075r\u0069\u0065\u0072\u002d\u0042\u006f\u006c\u0064")
	CourierObliqueName     = StdFontName("\u0043o\u0075r\u0069\u0065\u0072\u002d\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
	CourierBoldObliqueName = StdFontName("\u0043\u006f\u0075\u0072ie\u0072\u002d\u0042\u006f\u006c\u0064\u004f\u0062\u006c\u0069\u0071\u0075\u0065")
)

func TtfParse(r _d.ReadSeeker) (TtfType, error) { _fgca := &ttfParser{_cbee: r}; return _fgca.Parse() }
func _eae() StdFont {
	_daa.Do(_cbef)
	_cgagf := Descriptor{Name: TimesBoldName, Family: _ccd, Weight: FontWeightBold, Flags: 0x0020, BBox: [4]float64{-168, -218, 1000, 935}, ItalicAngle: 0, Ascent: 683, Descent: -217, CapHeight: 676, XHeight: 461, StemV: 139, StemH: 44}
	return NewStdFont(_cgagf, _ffa)
}
func _agb() StdFont {
	_daa.Do(_cbef)
	_eag := Descriptor{Name: TimesItalicName, Family: _ccd, Weight: FontWeightMedium, Flags: 0x0060, BBox: [4]float64{-169, -217, 1010, 883}, ItalicAngle: -15.5, Ascent: 683, Descent: -217, CapHeight: 653, XHeight: 441, StemV: 76, StemH: 32}
	return NewStdFont(_eag, _gcae)
}
func _bbb(_ccaf map[string]uint32) string {
	var _bea []string
	for _bfb := range _ccaf {
		_bea = append(_bea, _bfb)
	}
	_f.Slice(_bea, func(_bafb, _bcd int) bool { return _ccaf[_bea[_bafb]] < _ccaf[_bea[_bcd]] })
	_bbbf := []string{_g.Sprintf("\u0054\u0072\u0075\u0065Ty\u0070\u0065\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_ccaf))}
	for _, _aff := range _bea {
		_bbbf = append(_bbbf, _g.Sprintf("\u0009%\u0071\u0020\u0025\u0035\u0064", _aff, _ccaf[_aff]))
	}
	return _ga.Join(_bbbf, "\u000a")
}
func (_gaf StdFont) GetMetricsTable() *RuneCharSafeMap { return _gaf._eec }
func (_ac *fontMap) read(_ad StdFontName) (func() StdFont, bool) {
	_ac.Lock()
	defer _ac.Unlock()
	_fbf, _ag := _ac._be[_ad]
	return _fbf, _ag
}
