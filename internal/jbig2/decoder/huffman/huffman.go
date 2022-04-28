package huffman

import (
	_d "errors"
	_fd "fmt"
	_aa "math"
	_c "strings"

	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_a "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
)

func (_bc *EncodedTable) InitTree(codeTable []*Code) error {
	_faf(codeTable)
	for _, _g := range codeTable {
		if _bb := _bc._fg.append(_g); _bb != nil {
			return _bb
		}
	}
	return nil
}
func (_gd *EncodedTable) String() string                            { return _gd._fg.String() + "\u000a" }
func (_df *FixedSizeTable) Decode(r _f.StreamReader) (int64, error) { return _df._bba.Decode(r) }
func (_efb *FixedSizeTable) RootNode() *InternalNode                { return _efb._bba }
func (_ge *FixedSizeTable) InitTree(codeTable []*Code) error {
	_faf(codeTable)
	for _, _ea := range codeTable {
		_fdg := _ge._bba.append(_ea)
		if _fdg != nil {
			return _fdg
		}
	}
	return nil
}

var _ Node = &InternalNode{}

func (_cb *InternalNode) append(_agg *Code) (_aba error) {
	if _agg._gff == 0 {
		return nil
	}
	_gbb := _agg._gff - 1 - _cb._bd
	if _gbb < 0 {
		return _d.New("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	_fgf := (_agg._fed >> uint(_gbb)) & 0x1
	if _gbb == 0 {
		if _agg._fga == -1 {
			if _fgf == 1 {
				if _cb._gfc != nil {
					return _fd.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _agg)
				}
				_cb._gfc = _bg(_agg)
			} else {
				if _cb._cgb != nil {
					return _fd.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _agg)
				}
				_cb._cgb = _bg(_agg)
			}
		} else {
			if _fgf == 1 {
				if _cb._gfc != nil {
					return _fd.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _agg)
				}
				_cb._gfc = _gdd(_agg)
			} else {
				if _cb._cgb != nil {
					return _fd.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _agg)
				}
				_cb._cgb = _gdd(_agg)
			}
		}
	} else {
		if _fgf == 1 {
			if _cb._gfc == nil {
				_cb._gfc = _gfa(_cb._bd + 1)
			}
			if _aba = _cb._gfc.(*InternalNode).append(_agg); _aba != nil {
				return _aba
			}
		} else {
			if _cb._cgb == nil {
				_cb._cgb = _gfa(_cb._bd + 1)
			}
			if _aba = _cb._cgb.(*InternalNode).append(_agg); _aba != nil {
				return _aba
			}
		}
	}
	return nil
}
func (_ef *EncodedTable) RootNode() *InternalNode                   { return _ef._fg }
func (_ccd *StandardTable) Decode(r _f.StreamReader) (int64, error) { return _ccd._fe.Decode(r) }
func NewEncodedTable(table BasicTabler) (*EncodedTable, error) {
	_b := &EncodedTable{_fg: &InternalNode{}, BasicTabler: table}
	if _dd := _b.parseTable(); _dd != nil {
		return nil, _dd
	}
	return _b, nil
}
func (_ec *OutOfBandNode) Decode(r _f.StreamReader) (int64, error) { return 0, _a.ErrOOB }
func (_debb *Code) String() string {
	var _dfd string
	if _debb._fed != -1 {
		_dfd = _fb(_debb._fed, _debb._gff)
	} else {
		_dfd = "\u003f"
	}
	return _fd.Sprintf("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _dfd, _debb._gff, _debb._fga, _debb._eda)
}

type ValueNode struct {
	_ceg int32
	_be  int32
	_dg  bool
}

func (_ee *StandardTable) String() string { return _ee._fe.String() + "\u000a" }
func _cdgc(_ebc, _adda int32) int32 {
	if _ebc > _adda {
		return _ebc
	}
	return _adda
}

var _ Node = &OutOfBandNode{}

func _gfa(_dga int32) *InternalNode { return &InternalNode{_bd: _dga} }
func NewFixedSizeTable(codeTable []*Code) (*FixedSizeTable, error) {
	_ded := &FixedSizeTable{_bba: &InternalNode{}}
	if _gf := _ded.InitTree(codeTable); _gf != nil {
		return nil, _gf
	}
	return _ded, nil
}

type OutOfBandNode struct{}

func (_bbac *StandardTable) InitTree(codeTable []*Code) error {
	_faf(codeTable)
	for _, _ced := range codeTable {
		if _dee := _bbac._fe.append(_ced); _dee != nil {
			return _dee
		}
	}
	return nil
}

var _ Node = &ValueNode{}

func (_deb *ValueNode) Decode(r _f.StreamReader) (int64, error) {
	_bfe, _dc := r.ReadBits(byte(_deb._ceg))
	if _dc != nil {
		return 0, _dc
	}
	if _deb._dg {
		_bfe = -_bfe
	}
	return int64(_deb._be) + int64(_bfe), nil
}

type FixedSizeTable struct{ _bba *InternalNode }

var _gbc = make([]Tabler, len(_ead))

func _fb(_fbg, _gab int32) string {
	var _eff int32
	_fbd := make([]rune, _gab)
	for _cec := int32(1); _cec <= _gab; _cec++ {
		_eff = _fbg >> uint(_gab-_cec) & 1
		if _eff != 0 {
			_fbd[_cec-1] = '1'
		} else {
			_fbd[_cec-1] = '0'
		}
	}
	return string(_fbd)
}
func (_eae *FixedSizeTable) String() string { return _eae._bba.String() + "\u000a" }
func (_cd *ValueNode) String() string {
	return _fd.Sprintf("\u0025\u0064\u002f%\u0064", _cd._ceg, _cd._be)
}

type BasicTabler interface {
	HtHigh() int32
	HtLow() int32
	StreamReader() _f.StreamReader
	HtPS() int32
	HtRS() int32
	HtOOB() int32
}

func _ade(_dbf [][]int32) (*StandardTable, error) {
	var _abd []*Code
	for _add := 0; _add < len(_dbf); _add++ {
		_ege := _dbf[_add][0]
		_dbe := _dbf[_add][1]
		_eba := _dbf[_add][2]
		var _adg bool
		if len(_dbf[_add]) > 3 {
			_adg = true
		}
		_abd = append(_abd, NewCode(_ege, _dbe, _eba, _adg))
	}
	_ff := &StandardTable{_fe: _gfa(0)}
	if _gee := _ff.InitTree(_abd); _gee != nil {
		return nil, _gee
	}
	return _ff, nil
}
func (_cg *EncodedTable) Decode(r _f.StreamReader) (int64, error) { return _cg._fg.Decode(r) }
func (_cbc *InternalNode) pad(_ga *_c.Builder) {
	for _dgf := int32(0); _dgf < _cbc._bd; _dgf++ {
		_ga.WriteString("\u0020\u0020\u0020")
	}
}

var _ead = [][][]int32{{{1, 4, 0}, {2, 8, 16}, {3, 16, 272}, {3, 32, 65808}}, {{1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {6, 32, 75}, {6, -1, 0}}, {{8, 8, -256}, {1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {8, 32, -257, 999}, {7, 32, 75}, {6, -1, 0}}, {{1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {5, 32, 76}}, {{7, 8, -255}, {1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {7, 32, -256, 999}, {6, 32, 76}}, {{5, 10, -2048}, {4, 9, -1024}, {4, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {2, 7, 0}, {3, 7, 128}, {3, 8, 256}, {4, 9, 512}, {4, 10, 1024}, {6, 32, -2049, 999}, {6, 32, 2048}}, {{4, 9, -1024}, {3, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {4, 5, 0}, {5, 5, 32}, {5, 6, 64}, {4, 7, 128}, {3, 8, 256}, {3, 9, 512}, {3, 10, 1024}, {5, 32, -1025, 999}, {5, 32, 2048}}, {{8, 3, -15}, {9, 1, -7}, {8, 1, -5}, {9, 0, -3}, {7, 0, -2}, {4, 0, -1}, {2, 1, 0}, {5, 0, 2}, {6, 0, 3}, {3, 4, 4}, {6, 1, 20}, {4, 4, 22}, {4, 5, 38}, {5, 6, 70}, {5, 7, 134}, {6, 7, 262}, {7, 8, 390}, {6, 10, 646}, {9, 32, -16, 999}, {9, 32, 1670}, {2, -1, 0}}, {{8, 4, -31}, {9, 2, -15}, {8, 2, -11}, {9, 1, -7}, {7, 1, -5}, {4, 1, -3}, {3, 1, -1}, {3, 1, 1}, {5, 1, 3}, {6, 1, 5}, {3, 5, 7}, {6, 2, 39}, {4, 5, 43}, {4, 6, 75}, {5, 7, 139}, {5, 8, 267}, {6, 8, 523}, {7, 9, 779}, {6, 11, 1291}, {9, 32, -32, 999}, {9, 32, 3339}, {2, -1, 0}}, {{7, 4, -21}, {8, 0, -5}, {7, 0, -4}, {5, 0, -3}, {2, 2, -2}, {5, 0, 2}, {6, 0, 3}, {7, 0, 4}, {8, 0, 5}, {2, 6, 6}, {5, 5, 70}, {6, 5, 102}, {6, 6, 134}, {6, 7, 198}, {6, 8, 326}, {6, 9, 582}, {6, 10, 1094}, {7, 11, 2118}, {8, 32, -22, 999}, {8, 32, 4166}, {2, -1, 0}}, {{1, 0, 1}, {2, 1, 2}, {4, 0, 4}, {4, 1, 5}, {5, 1, 7}, {5, 2, 9}, {6, 2, 13}, {7, 2, 17}, {7, 3, 21}, {7, 4, 29}, {7, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{1, 0, 1}, {2, 0, 2}, {3, 1, 3}, {5, 0, 5}, {5, 1, 6}, {6, 1, 8}, {7, 0, 10}, {7, 1, 11}, {7, 2, 13}, {7, 3, 17}, {7, 4, 25}, {8, 5, 41}, {8, 32, 73}}, {{1, 0, 1}, {3, 0, 2}, {4, 0, 3}, {5, 0, 4}, {4, 1, 5}, {3, 3, 7}, {6, 1, 15}, {6, 2, 17}, {6, 3, 21}, {6, 4, 29}, {6, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{3, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {3, 0, 2}}, {{7, 4, -24}, {6, 2, -8}, {5, 1, -4}, {4, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {4, 0, 2}, {5, 1, 3}, {6, 2, 5}, {7, 4, 9}, {7, 32, -25, 999}, {7, 32, 25}}}

func (_gbg *InternalNode) Decode(r _f.StreamReader) (int64, error) {
	_bge, _eac := r.ReadBit()
	if _eac != nil {
		return 0, _eac
	}
	if _bge == 0 {
		return _gbg._cgb.Decode(r)
	}
	return _gbg._gfc.Decode(r)
}
func (_ag *InternalNode) String() string {
	_acd := &_c.Builder{}
	_acd.WriteString("\u000a")
	_ag.pad(_acd)
	_acd.WriteString("\u0030\u003a\u0020")
	_acd.WriteString(_ag._cgb.String() + "\u000a")
	_ag.pad(_acd)
	_acd.WriteString("\u0031\u003a\u0020")
	_acd.WriteString(_ag._gfc.String() + "\u000a")
	return _acd.String()
}

type Code struct {
	_gff int32
	_fga int32
	_eda int32
	_dbg bool
	_fed int32
}

func (_fc *EncodedTable) parseTable() error {
	var (
		_bbc           []*Code
		_cf, _fdb, _cc int32
		_de            uint64
		_ce            error
	)
	_ac := _fc.StreamReader()
	_ad := _fc.HtLow()
	for _ad < _fc.HtHigh() {
		_de, _ce = _ac.ReadBits(byte(_fc.HtPS()))
		if _ce != nil {
			return _ce
		}
		_cf = int32(_de)
		_de, _ce = _ac.ReadBits(byte(_fc.HtRS()))
		if _ce != nil {
			return _ce
		}
		_fdb = int32(_de)
		_bbc = append(_bbc, NewCode(_cf, _fdb, _cc, false))
		_ad += 1 << uint(_fdb)
	}
	_de, _ce = _ac.ReadBits(byte(_fc.HtPS()))
	if _ce != nil {
		return _ce
	}
	_cf = int32(_de)
	_fdb = 32
	_cc = _fc.HtLow() - 1
	_bbc = append(_bbc, NewCode(_cf, _fdb, _cc, true))
	_de, _ce = _ac.ReadBits(byte(_fc.HtPS()))
	if _ce != nil {
		return _ce
	}
	_cf = int32(_de)
	_fdb = 32
	_cc = _fc.HtHigh()
	_bbc = append(_bbc, NewCode(_cf, _fdb, _cc, false))
	if _fc.HtOOB() == 1 {
		_de, _ce = _ac.ReadBits(byte(_fc.HtPS()))
		if _ce != nil {
			return _ce
		}
		_cf = int32(_de)
		_bbc = append(_bbc, NewCode(_cf, -1, -1, false))
	}
	if _ce = _fc.InitTree(_bbc); _ce != nil {
		return _ce
	}
	return nil
}

type Tabler interface {
	Decode(_fab _f.StreamReader) (int64, error)
	InitTree(_ged []*Code) error
	String() string
	RootNode() *InternalNode
}

func (_gb *OutOfBandNode) String() string {
	return _fd.Sprintf("\u0025\u0030\u00364\u0062", int64(_aa.MaxInt64))
}
func _bg(_bgb *Code) *OutOfBandNode { return &OutOfBandNode{} }

type Node interface {
	Decode(_bfg _f.StreamReader) (int64, error)
	String() string
}
type EncodedTable struct {
	BasicTabler
	_fg *InternalNode
}

func _gdd(_eb *Code) *ValueNode { return &ValueNode{_ceg: _eb._fga, _be: _eb._eda, _dg: _eb._dbg} }
func NewCode(prefixLength, rangeLength, rangeLow int32, isLowerRange bool) *Code {
	return &Code{_gff: prefixLength, _fga: rangeLength, _eda: rangeLow, _dbg: isLowerRange, _fed: -1}
}

type InternalNode struct {
	_bd  int32
	_cgb Node
	_gfc Node
}

func GetStandardTable(number int) (Tabler, error) {
	if number <= 0 || number > len(_gbc) {
		return nil, _d.New("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_gbbf := _gbc[number-1]
	if _gbbf == nil {
		var _ed error
		_gbbf, _ed = _ade(_ead[number-1])
		if _ed != nil {
			return nil, _ed
		}
		_gbc[number-1] = _gbbf
	}
	return _gbbf, nil
}

var _ Tabler = &EncodedTable{}

func (_ae *StandardTable) RootNode() *InternalNode { return _ae._fe }

type StandardTable struct{ _fe *InternalNode }

func _faf(_cbf []*Code) {
	var _geea int32
	for _, _dbc := range _cbf {
		_geea = _cdgc(_geea, _dbc._gff)
	}
	_agc := make([]int32, _geea+1)
	for _, _adc := range _cbf {
		_agc[_adc._gff]++
	}
	var _ccg int32
	_da := make([]int32, len(_agc)+1)
	_agc[0] = 0
	for _dec := int32(1); _dec <= int32(len(_agc)); _dec++ {
		_da[_dec] = (_da[_dec-1] + (_agc[_dec-1])) << 1
		_ccg = _da[_dec]
		for _, _aed := range _cbf {
			if _aed._gff == _dec {
				_aed._fed = _ccg
				_ccg++
			}
		}
	}
}
