package huffman

import (
	_ed "errors"
	_ba "fmt"
	_b "math"
	_c "strings"

	_a "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_ab "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
)

func NewFixedSizeTable(codeTable []*Code) (*FixedSizeTable, error) {
	_egc := &FixedSizeTable{_ecf: &InternalNode{}}
	if _bede := _egc.InitTree(codeTable); _bede != nil {
		return nil, _bede
	}
	return _egc, nil
}

type BasicTabler interface {
	HtHigh() int32
	HtLow() int32
	StreamReader() *_a.Reader
	HtPS() int32
	HtRS() int32
	HtOOB() int32
}

func GetStandardTable(number int) (Tabler, error) {
	if number <= 0 || number > len(_cec) {
		return nil, _ed.New("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_eaa := _cec[number-1]
	if _eaa == nil {
		var _dgb error
		_eaa, _dgb = _bge(_ade[number-1])
		if _dgb != nil {
			return nil, _dgb
		}
		_cec[number-1] = _eaa
	}
	return _eaa, nil
}

var _ Tabler = &EncodedTable{}

type Tabler interface {
	Decode(_fcd *_a.Reader) (int64, error)
	InitTree(_gdb []*Code) error
	String() string
	RootNode() *InternalNode
}

func (_ga *FixedSizeTable) Decode(r *_a.Reader) (int64, error) { return _ga._ecf.Decode(r) }
func (_ce *StandardTable) String() string                      { return _ce._dcb.String() + "\u000a" }

var _ Node = &OutOfBandNode{}

type EncodedTable struct {
	BasicTabler
	_g *InternalNode
}
type OutOfBandNode struct{}

func _dc(_faa *Code) *OutOfBandNode { return &OutOfBandNode{} }
func _ebce(_fbe, _bdc int32) int32 {
	if _fbe > _bdc {
		return _fbe
	}
	return _bdc
}
func (_ec *EncodedTable) Decode(r *_a.Reader) (int64, error) { return _ec._g.Decode(r) }

var _cec = make([]Tabler, len(_ade))

func (_acd *ValueNode) String() string {
	return _ba.Sprintf("\u0025\u0064\u002f%\u0064", _acd._aa, _acd._eed)
}
func (_bd *OutOfBandNode) Decode(r *_a.Reader) (int64, error) { return 0, _ab.ErrOOB }
func (_bae *InternalNode) pad(_dad *_c.Builder) {
	for _gadd := int32(0); _gadd < _bae._gc; _gadd++ {
		_dad.WriteString("\u0020\u0020\u0020")
	}
}
func (_gad *InternalNode) String() string {
	_ecg := &_c.Builder{}
	_ecg.WriteString("\u000a")
	_gad.pad(_ecg)
	_ecg.WriteString("\u0030\u003a\u0020")
	_ecg.WriteString(_gad._fdb.String() + "\u000a")
	_gad.pad(_ecg)
	_ecg.WriteString("\u0031\u003a\u0020")
	_ecg.WriteString(_gad._dgf.String() + "\u000a")
	return _ecg.String()
}
func (_acc *FixedSizeTable) InitTree(codeTable []*Code) error {
	_adef(codeTable)
	for _, _bb := range codeTable {
		_acg := _acc._ecf.append(_bb)
		if _acg != nil {
			return _acg
		}
	}
	return nil
}

type FixedSizeTable struct{ _ecf *InternalNode }

func (_bg *ValueNode) Decode(r *_a.Reader) (int64, error) {
	_gd, _da := r.ReadBits(byte(_bg._aa))
	if _da != nil {
		return 0, _da
	}
	if _bg._aae {
		_gd = -_gd
	}
	return int64(_bg._eed) + int64(_gd), nil
}

var _ Node = &ValueNode{}

func _beg(_fdf *Code) *ValueNode { return &ValueNode{_aa: _fdf._eac, _eed: _fdf._ge, _aae: _fdf._fg} }
func _adef(_cbb []*Code) {
	var _cf int32
	for _, _dcg := range _cbb {
		_cf = _ebce(_cf, _dcg._cba)
	}
	_cfa := make([]int32, _cf+1)
	for _, _ebc := range _cbb {
		_cfa[_ebc._cba]++
	}
	var _bac int32
	_bgb := make([]int32, len(_cfa)+1)
	_cfa[0] = 0
	for _bab := int32(1); _bab <= int32(len(_cfa)); _bab++ {
		_bgb[_bab] = (_bgb[_bab-1] + (_cfa[_bab-1])) << 1
		_bac = _bgb[_bab]
		for _, _fgcc := range _cbb {
			if _fgcc._cba == _bab {
				_fgcc._de = _bac
				_bac++
			}
		}
	}
}

var _ade = [][][]int32{{{1, 4, 0}, {2, 8, 16}, {3, 16, 272}, {3, 32, 65808}}, {{1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {6, 32, 75}, {6, -1, 0}}, {{8, 8, -256}, {1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {8, 32, -257, 999}, {7, 32, 75}, {6, -1, 0}}, {{1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {5, 32, 76}}, {{7, 8, -255}, {1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {7, 32, -256, 999}, {6, 32, 76}}, {{5, 10, -2048}, {4, 9, -1024}, {4, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {2, 7, 0}, {3, 7, 128}, {3, 8, 256}, {4, 9, 512}, {4, 10, 1024}, {6, 32, -2049, 999}, {6, 32, 2048}}, {{4, 9, -1024}, {3, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {4, 5, 0}, {5, 5, 32}, {5, 6, 64}, {4, 7, 128}, {3, 8, 256}, {3, 9, 512}, {3, 10, 1024}, {5, 32, -1025, 999}, {5, 32, 2048}}, {{8, 3, -15}, {9, 1, -7}, {8, 1, -5}, {9, 0, -3}, {7, 0, -2}, {4, 0, -1}, {2, 1, 0}, {5, 0, 2}, {6, 0, 3}, {3, 4, 4}, {6, 1, 20}, {4, 4, 22}, {4, 5, 38}, {5, 6, 70}, {5, 7, 134}, {6, 7, 262}, {7, 8, 390}, {6, 10, 646}, {9, 32, -16, 999}, {9, 32, 1670}, {2, -1, 0}}, {{8, 4, -31}, {9, 2, -15}, {8, 2, -11}, {9, 1, -7}, {7, 1, -5}, {4, 1, -3}, {3, 1, -1}, {3, 1, 1}, {5, 1, 3}, {6, 1, 5}, {3, 5, 7}, {6, 2, 39}, {4, 5, 43}, {4, 6, 75}, {5, 7, 139}, {5, 8, 267}, {6, 8, 523}, {7, 9, 779}, {6, 11, 1291}, {9, 32, -32, 999}, {9, 32, 3339}, {2, -1, 0}}, {{7, 4, -21}, {8, 0, -5}, {7, 0, -4}, {5, 0, -3}, {2, 2, -2}, {5, 0, 2}, {6, 0, 3}, {7, 0, 4}, {8, 0, 5}, {2, 6, 6}, {5, 5, 70}, {6, 5, 102}, {6, 6, 134}, {6, 7, 198}, {6, 8, 326}, {6, 9, 582}, {6, 10, 1094}, {7, 11, 2118}, {8, 32, -22, 999}, {8, 32, 4166}, {2, -1, 0}}, {{1, 0, 1}, {2, 1, 2}, {4, 0, 4}, {4, 1, 5}, {5, 1, 7}, {5, 2, 9}, {6, 2, 13}, {7, 2, 17}, {7, 3, 21}, {7, 4, 29}, {7, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{1, 0, 1}, {2, 0, 2}, {3, 1, 3}, {5, 0, 5}, {5, 1, 6}, {6, 1, 8}, {7, 0, 10}, {7, 1, 11}, {7, 2, 13}, {7, 3, 17}, {7, 4, 25}, {8, 5, 41}, {8, 32, 73}}, {{1, 0, 1}, {3, 0, 2}, {4, 0, 3}, {5, 0, 4}, {4, 1, 5}, {3, 3, 7}, {6, 1, 15}, {6, 2, 17}, {6, 3, 21}, {6, 4, 29}, {6, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{3, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {3, 0, 2}}, {{7, 4, -24}, {6, 2, -8}, {5, 1, -4}, {4, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {4, 0, 2}, {5, 1, 3}, {6, 2, 5}, {7, 4, 9}, {7, 32, -25, 999}, {7, 32, 25}}}

type StandardTable struct{ _dcb *InternalNode }

func (_ef *InternalNode) Decode(r *_a.Reader) (int64, error) {
	_cg, _ca := r.ReadBit()
	if _ca != nil {
		return 0, _ca
	}
	if _cg == 0 {
		return _ef._fdb.Decode(r)
	}
	return _ef._dgf.Decode(r)
}
func NewEncodedTable(table BasicTabler) (*EncodedTable, error) {
	_ea := &EncodedTable{_g: &InternalNode{}, BasicTabler: table}
	if _f := _ea.parseTable(); _f != nil {
		return nil, _f
	}
	return _ea, nil
}
func (_dd *Code) String() string {
	var _df string
	if _dd._de != -1 {
		_df = _cbe(_dd._de, _dd._cba)
	} else {
		_df = "\u003f"
	}
	return _ba.Sprintf("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _df, _dd._cba, _dd._eac, _dd._ge)
}
func (_ceb *StandardTable) RootNode() *InternalNode { return _ceb._dcb }
func (_gb *FixedSizeTable) String() string          { return _gb._ecf.String() + "\u000a" }
func (_bgg *InternalNode) append(_ae *Code) (_fdg error) {
	if _ae._cba == 0 {
		return nil
	}
	_gcd := _ae._cba - 1 - _bgg._gc
	if _gcd < 0 {
		return _ed.New("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	_dac := (_ae._de >> uint(_gcd)) & 0x1
	if _gcd == 0 {
		if _ae._eac == -1 {
			if _dac == 1 {
				if _bgg._dgf != nil {
					return _ba.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _ae)
				}
				_bgg._dgf = _dc(_ae)
			} else {
				if _bgg._fdb != nil {
					return _ba.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _ae)
				}
				_bgg._fdb = _dc(_ae)
			}
		} else {
			if _dac == 1 {
				if _bgg._dgf != nil {
					return _ba.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _ae)
				}
				_bgg._dgf = _beg(_ae)
			} else {
				if _bgg._fdb != nil {
					return _ba.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _ae)
				}
				_bgg._fdb = _beg(_ae)
			}
		}
	} else {
		if _dac == 1 {
			if _bgg._dgf == nil {
				_bgg._dgf = _cb(_bgg._gc + 1)
			}
			if _fdg = _bgg._dgf.(*InternalNode).append(_ae); _fdg != nil {
				return _fdg
			}
		} else {
			if _bgg._fdb == nil {
				_bgg._fdb = _cb(_bgg._gc + 1)
			}
			if _fdg = _bgg._fdb.(*InternalNode).append(_ae); _fdg != nil {
				return _fdg
			}
		}
	}
	return nil
}
func _cbe(_aaea, _ddd int32) string {
	var _ddf int32
	_bga := make([]rune, _ddd)
	for _cbag := int32(1); _cbag <= _ddd; _cbag++ {
		_ddf = _aaea >> uint(_ddd-_cbag) & 1
		if _ddf != 0 {
			_bga[_cbag-1] = '1'
		} else {
			_bga[_cbag-1] = '0'
		}
	}
	return string(_bga)
}
func _cb(_ggf int32) *InternalNode { return &InternalNode{_gc: _ggf} }
func (_bed *EncodedTable) parseTable() error {
	var (
		_bad           []*Code
		_abe, _fc, _ad int32
		_abd           uint64
		_ac            error
	)
	_d := _bed.StreamReader()
	_dg := _bed.HtLow()
	for _dg < _bed.HtHigh() {
		_abd, _ac = _d.ReadBits(byte(_bed.HtPS()))
		if _ac != nil {
			return _ac
		}
		_abe = int32(_abd)
		_abd, _ac = _d.ReadBits(byte(_bed.HtRS()))
		if _ac != nil {
			return _ac
		}
		_fc = int32(_abd)
		_bad = append(_bad, NewCode(_abe, _fc, _ad, false))
		_dg += 1 << uint(_fc)
	}
	_abd, _ac = _d.ReadBits(byte(_bed.HtPS()))
	if _ac != nil {
		return _ac
	}
	_abe = int32(_abd)
	_fc = 32
	_ad = _bed.HtLow() - 1
	_bad = append(_bad, NewCode(_abe, _fc, _ad, true))
	_abd, _ac = _d.ReadBits(byte(_bed.HtPS()))
	if _ac != nil {
		return _ac
	}
	_abe = int32(_abd)
	_fc = 32
	_ad = _bed.HtHigh()
	_bad = append(_bad, NewCode(_abe, _fc, _ad, false))
	if _bed.HtOOB() == 1 {
		_abd, _ac = _d.ReadBits(byte(_bed.HtPS()))
		if _ac != nil {
			return _ac
		}
		_abe = int32(_abd)
		_bad = append(_bad, NewCode(_abe, -1, -1, false))
	}
	if _ac = _bed.InitTree(_bad); _ac != nil {
		return _ac
	}
	return nil
}

type Node interface {
	Decode(_ee *_a.Reader) (int64, error)
	String() string
}

var _ Node = &InternalNode{}

type Code struct {
	_cba int32
	_eac int32
	_ge  int32
	_fg  bool
	_de  int32
}

func (_bc *FixedSizeTable) RootNode() *InternalNode { return _bc._ecf }
func (_fd *EncodedTable) InitTree(codeTable []*Code) error {
	_adef(codeTable)
	for _, _eg := range codeTable {
		if _eca := _fd._g.append(_eg); _eca != nil {
			return _eca
		}
	}
	return nil
}
func (_gga *StandardTable) Decode(r *_a.Reader) (int64, error) { return _gga._dcb.Decode(r) }
func (_fb *EncodedTable) String() string                       { return _fb._g.String() + "\u000a" }
func (_bba *OutOfBandNode) String() string {
	return _ba.Sprintf("\u0025\u0030\u00364\u0062", int64(_b.MaxInt64))
}

type ValueNode struct {
	_aa  int32
	_eed int32
	_aae bool
}

func (_gf *EncodedTable) RootNode() *InternalNode { return _gf._g }
func NewCode(prefixLength, rangeLength, rangeLow int32, isLowerRange bool) *Code {
	return &Code{_cba: prefixLength, _eac: rangeLength, _ge: rangeLow, _fg: isLowerRange, _de: -1}
}
func _bge(_efd [][]int32) (*StandardTable, error) {
	var _eag []*Code
	for _eaf := 0; _eaf < len(_efd); _eaf++ {
		_abb := _efd[_eaf][0]
		_afd := _efd[_eaf][1]
		_acf := _efd[_eaf][2]
		var _eae bool
		if len(_efd[_eaf]) > 3 {
			_eae = true
		}
		_eag = append(_eag, NewCode(_abb, _afd, _acf, _eae))
	}
	_gfa := &StandardTable{_dcb: _cb(0)}
	if _gaa := _gfa.InitTree(_eag); _gaa != nil {
		return nil, _gaa
	}
	return _gfa, nil
}

type InternalNode struct {
	_gc  int32
	_fdb Node
	_dgf Node
}

func (_egf *StandardTable) InitTree(codeTable []*Code) error {
	_adef(codeTable)
	for _, _fca := range codeTable {
		if _aca := _egf._dcb.append(_fca); _aca != nil {
			return _aca
		}
	}
	return nil
}
