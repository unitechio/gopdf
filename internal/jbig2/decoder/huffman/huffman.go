package huffman

import (
	_f "errors"
	_cf "fmt"
	_a "math"
	_ca "strings"

	_c "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_g "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
)

func (_ac *InternalNode) String() string {
	_adb := &_ca.Builder{}
	_adb.WriteString("\u000a")
	_ac.pad(_adb)
	_adb.WriteString("\u0030\u003a\u0020")
	_adb.WriteString(_ac._dd.String() + "\u000a")
	_ac.pad(_adb)
	_adb.WriteString("\u0031\u003a\u0020")
	_adb.WriteString(_ac._fb.String() + "\u000a")
	return _adb.String()
}
func (_gb *EncodedTable) Decode(r _c.StreamReader) (int64, error) { return _gb._fe.Decode(r) }
func GetStandardTable(number int) (Tabler, error) {
	if number <= 0 || number > len(_bgb) {
		return nil, _f.New("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bfa := _bgb[number-1]
	if _bfa == nil {
		var _fff error
		_bfa, _fff = _aff(_cbc[number-1])
		if _fff != nil {
			return nil, _fff
		}
		_bgb[number-1] = _bfa
	}
	return _bfa, nil
}

var _ Node = &InternalNode{}

func NewEncodedTable(table BasicTabler) (*EncodedTable, error) {
	_cff := &EncodedTable{_fe: &InternalNode{}, BasicTabler: table}
	if _ce := _cff.parseTable(); _ce != nil {
		return nil, _ce
	}
	return _cff, nil
}
func (_bg *EncodedTable) parseTable() error {
	var (
		_ba            []*Code
		_aa, _bge, _gd int32
		_cfg           uint64
		_dc            error
	)
	_be := _bg.StreamReader()
	_gda := _bg.HtLow()
	for _gda < _bg.HtHigh() {
		_cfg, _dc = _be.ReadBits(byte(_bg.HtPS()))
		if _dc != nil {
			return _dc
		}
		_aa = int32(_cfg)
		_cfg, _dc = _be.ReadBits(byte(_bg.HtRS()))
		if _dc != nil {
			return _dc
		}
		_bge = int32(_cfg)
		_ba = append(_ba, NewCode(_aa, _bge, _gd, false))
		_gda += 1 << uint(_bge)
	}
	_cfg, _dc = _be.ReadBits(byte(_bg.HtPS()))
	if _dc != nil {
		return _dc
	}
	_aa = int32(_cfg)
	_bge = 32
	_gd = _bg.HtLow() - 1
	_ba = append(_ba, NewCode(_aa, _bge, _gd, true))
	_cfg, _dc = _be.ReadBits(byte(_bg.HtPS()))
	if _dc != nil {
		return _dc
	}
	_aa = int32(_cfg)
	_bge = 32
	_gd = _bg.HtHigh()
	_ba = append(_ba, NewCode(_aa, _bge, _gd, false))
	if _bg.HtOOB() == 1 {
		_cfg, _dc = _be.ReadBits(byte(_bg.HtPS()))
		if _dc != nil {
			return _dc
		}
		_aa = int32(_cfg)
		_ba = append(_ba, NewCode(_aa, -1, -1, false))
	}
	if _dc = _bg.InitTree(_ba); _dc != nil {
		return _dc
	}
	return nil
}

type BasicTabler interface {
	HtHigh() int32
	HtLow() int32
	StreamReader() _c.StreamReader
	HtPS() int32
	HtRS() int32
	HtOOB() int32
}

var _cbc = [][][]int32{{{1, 4, 0}, {2, 8, 16}, {3, 16, 272}, {3, 32, 65808}}, {{1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {6, 32, 75}, {6, -1, 0}}, {{8, 8, -256}, {1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {8, 32, -257, 999}, {7, 32, 75}, {6, -1, 0}}, {{1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {5, 32, 76}}, {{7, 8, -255}, {1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {7, 32, -256, 999}, {6, 32, 76}}, {{5, 10, -2048}, {4, 9, -1024}, {4, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {2, 7, 0}, {3, 7, 128}, {3, 8, 256}, {4, 9, 512}, {4, 10, 1024}, {6, 32, -2049, 999}, {6, 32, 2048}}, {{4, 9, -1024}, {3, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {4, 5, 0}, {5, 5, 32}, {5, 6, 64}, {4, 7, 128}, {3, 8, 256}, {3, 9, 512}, {3, 10, 1024}, {5, 32, -1025, 999}, {5, 32, 2048}}, {{8, 3, -15}, {9, 1, -7}, {8, 1, -5}, {9, 0, -3}, {7, 0, -2}, {4, 0, -1}, {2, 1, 0}, {5, 0, 2}, {6, 0, 3}, {3, 4, 4}, {6, 1, 20}, {4, 4, 22}, {4, 5, 38}, {5, 6, 70}, {5, 7, 134}, {6, 7, 262}, {7, 8, 390}, {6, 10, 646}, {9, 32, -16, 999}, {9, 32, 1670}, {2, -1, 0}}, {{8, 4, -31}, {9, 2, -15}, {8, 2, -11}, {9, 1, -7}, {7, 1, -5}, {4, 1, -3}, {3, 1, -1}, {3, 1, 1}, {5, 1, 3}, {6, 1, 5}, {3, 5, 7}, {6, 2, 39}, {4, 5, 43}, {4, 6, 75}, {5, 7, 139}, {5, 8, 267}, {6, 8, 523}, {7, 9, 779}, {6, 11, 1291}, {9, 32, -32, 999}, {9, 32, 3339}, {2, -1, 0}}, {{7, 4, -21}, {8, 0, -5}, {7, 0, -4}, {5, 0, -3}, {2, 2, -2}, {5, 0, 2}, {6, 0, 3}, {7, 0, 4}, {8, 0, 5}, {2, 6, 6}, {5, 5, 70}, {6, 5, 102}, {6, 6, 134}, {6, 7, 198}, {6, 8, 326}, {6, 9, 582}, {6, 10, 1094}, {7, 11, 2118}, {8, 32, -22, 999}, {8, 32, 4166}, {2, -1, 0}}, {{1, 0, 1}, {2, 1, 2}, {4, 0, 4}, {4, 1, 5}, {5, 1, 7}, {5, 2, 9}, {6, 2, 13}, {7, 2, 17}, {7, 3, 21}, {7, 4, 29}, {7, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{1, 0, 1}, {2, 0, 2}, {3, 1, 3}, {5, 0, 5}, {5, 1, 6}, {6, 1, 8}, {7, 0, 10}, {7, 1, 11}, {7, 2, 13}, {7, 3, 17}, {7, 4, 25}, {8, 5, 41}, {8, 32, 73}}, {{1, 0, 1}, {3, 0, 2}, {4, 0, 3}, {5, 0, 4}, {4, 1, 5}, {3, 3, 7}, {6, 1, 15}, {6, 2, 17}, {6, 3, 21}, {6, 4, 29}, {6, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{3, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {3, 0, 2}}, {{7, 4, -24}, {6, 2, -8}, {5, 1, -4}, {4, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {4, 0, 2}, {5, 1, 3}, {6, 2, 5}, {7, 4, 9}, {7, 32, -25, 999}, {7, 32, 25}}}

func (_ggf *EncodedTable) String() string { return _ggf._fe.String() + "\u000a" }
func (_dfa *InternalNode) pad(_cfe *_ca.Builder) {
	for _de := int32(0); _de < _dfa._dac; _de++ {
		_cfe.WriteString("\u0020\u0020\u0020")
	}
}
func NewCode(prefixLength, rangeLength, rangeLow int32, isLowerRange bool) *Code {
	return &Code{_geb: prefixLength, _fa: rangeLength, _dff: rangeLow, _egg: isLowerRange, _cga: -1}
}
func (_fc *OutOfBandNode) String() string {
	return _cf.Sprintf("\u0025\u0030\u00364\u0062", int64(_a.MaxInt64))
}
func (_eeb *Code) String() string {
	var _cc string
	if _eeb._cga != -1 {
		_cc = _fce(_eeb._cga, _eeb._geb)
	} else {
		_cc = "\u003f"
	}
	return _cf.Sprintf("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _cc, _eeb._geb, _eeb._fa, _eeb._dff)
}
func _bc(_dcd *Code) *OutOfBandNode { return &OutOfBandNode{} }
func (_gdc *InternalNode) Decode(r _c.StreamReader) (int64, error) {
	_fef, _dbg := r.ReadBit()
	if _dbg != nil {
		return 0, _dbg
	}
	if _fef == 0 {
		return _gdc._dd.Decode(r)
	}
	return _gdc._fb.Decode(r)
}

type InternalNode struct {
	_dac int32
	_dd  Node
	_fb  Node
}

var _ Node = &ValueNode{}

func _aff(_fdf [][]int32) (*StandardTable, error) {
	var _fba []*Code
	for _efa := 0; _efa < len(_fdf); _efa++ {
		_dacf := _fdf[_efa][0]
		_egf := _fdf[_efa][1]
		_geg := _fdf[_efa][2]
		var _egb bool
		if len(_fdf[_efa]) > 3 {
			_egb = true
		}
		_fba = append(_fba, NewCode(_dacf, _egf, _geg, _egb))
	}
	_aaf := &StandardTable{_cgd: _bgg(0)}
	if _ga := _aaf.InitTree(_fba); _ga != nil {
		return nil, _ga
	}
	return _aaf, nil
}
func (_adbd *StandardTable) Decode(r _c.StreamReader) (int64, error) { return _adbd._cgd.Decode(r) }

type OutOfBandNode struct{}

func (_da *FixedSizeTable) Decode(r _c.StreamReader) (int64, error) { return _da._ceg.Decode(r) }
func (_ada *ValueNode) String() string {
	return _cf.Sprintf("\u0025\u0064\u002f%\u0064", _ada._bb, _ada._bgd)
}
func NewFixedSizeTable(codeTable []*Code) (*FixedSizeTable, error) {
	_bed := &FixedSizeTable{_ceg: &InternalNode{}}
	if _cg := _bed.InitTree(codeTable); _cg != nil {
		return nil, _cg
	}
	return _bed, nil
}

type FixedSizeTable struct{ _ceg *InternalNode }

func (_b *EncodedTable) InitTree(codeTable []*Code) error {
	_fea(codeTable)
	for _, _dg := range codeTable {
		if _ae := _b._fe.append(_dg); _ae != nil {
			return _ae
		}
	}
	return nil
}

type StandardTable struct{ _cgd *InternalNode }

func (_ed *StandardTable) InitTree(codeTable []*Code) error {
	_fea(codeTable)
	for _, _fdd := range codeTable {
		if _egd := _ed._cgd.append(_fdd); _egd != nil {
			return _egd
		}
	}
	return nil
}

type ValueNode struct {
	_bb  int32
	_bgd int32
	_ff  bool
}
type Tabler interface {
	Decode(_bad _c.StreamReader) (int64, error)
	InitTree(_cgg []*Code) error
	String() string
	RootNode() *InternalNode
}

func (_cef *ValueNode) Decode(r _c.StreamReader) (int64, error) {
	_bcf, _fde := r.ReadBits(byte(_cef._bb))
	if _fde != nil {
		return 0, _fde
	}
	if _cef._ff {
		_bcf = -_bcf
	}
	return int64(_cef._bgd) + int64(_bcf), nil
}

var _ Node = &OutOfBandNode{}

type EncodedTable struct {
	BasicTabler
	_fe *InternalNode
}

func _af(_gbd *Code) *ValueNode { return &ValueNode{_bb: _gbd._fa, _bgd: _gbd._dff, _ff: _gbd._egg} }
func _fea(_dgd []*Code) {
	var _ggd int32
	for _, _dbd := range _dgd {
		_ggd = _fcd(_ggd, _dbd._geb)
	}
	_ddb := make([]int32, _ggd+1)
	for _, _fdcc := range _dgd {
		_ddb[_fdcc._geb]++
	}
	var _dec int32
	_bgga := make([]int32, len(_ddb)+1)
	_ddb[0] = 0
	for _dgc := int32(1); _dgc <= int32(len(_ddb)); _dgc++ {
		_bgga[_dgc] = (_bgga[_dgc-1] + (_ddb[_dgc-1])) << 1
		_dec = _bgga[_dgc]
		for _, _bfd := range _dgd {
			if _bfd._geb == _dgc {
				_bfd._cga = _dec
				_dec++
			}
		}
	}
}
func (_df *FixedSizeTable) InitTree(codeTable []*Code) error {
	_fea(codeTable)
	for _, _ef := range codeTable {
		_ee := _df._ceg.append(_ef)
		if _ee != nil {
			return _ee
		}
	}
	return nil
}
func (_ggfe *InternalNode) append(_bf *Code) (_ag error) {
	if _bf._geb == 0 {
		return nil
	}
	_cb := _bf._geb - 1 - _ggfe._dac
	if _cb < 0 {
		return _f.New("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	_fdc := (_bf._cga >> uint(_cb)) & 0x1
	if _cb == 0 {
		if _bf._fa == -1 {
			if _fdc == 1 {
				if _ggfe._fb != nil {
					return _cf.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _bf)
				}
				_ggfe._fb = _bc(_bf)
			} else {
				if _ggfe._dd != nil {
					return _cf.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _bf)
				}
				_ggfe._dd = _bc(_bf)
			}
		} else {
			if _fdc == 1 {
				if _ggfe._fb != nil {
					return _cf.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _bf)
				}
				_ggfe._fb = _af(_bf)
			} else {
				if _ggfe._dd != nil {
					return _cf.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _bf)
				}
				_ggfe._dd = _af(_bf)
			}
		}
	} else {
		if _fdc == 1 {
			if _ggfe._fb == nil {
				_ggfe._fb = _bgg(_ggfe._dac + 1)
			}
			if _ag = _ggfe._fb.(*InternalNode).append(_bf); _ag != nil {
				return _ag
			}
		} else {
			if _ggfe._dd == nil {
				_ggfe._dd = _bgg(_ggfe._dac + 1)
			}
			if _ag = _ggfe._dd.(*InternalNode).append(_bf); _ag != nil {
				return _ag
			}
		}
	}
	return nil
}
func (_ad *OutOfBandNode) Decode(r _c.StreamReader) (int64, error) { return 0, _g.ErrOOB }

var _ Tabler = &EncodedTable{}

func _fce(_daf, _aaa int32) string {
	var _bff int32
	_bd := make([]rune, _aaa)
	for _efag := int32(1); _efag <= _aaa; _efag++ {
		_bff = _daf >> uint(_aaa-_efag) & 1
		if _bff != 0 {
			_bd[_efag-1] = '1'
		} else {
			_bd[_efag-1] = '0'
		}
	}
	return string(_bd)
}
func _bgg(_cba int32) *InternalNode { return &InternalNode{_dac: _cba} }

type Node interface {
	Decode(_bea _c.StreamReader) (int64, error)
	String() string
}
type Code struct {
	_geb int32
	_fa  int32
	_dff int32
	_egg bool
	_cga int32
}

func (_caa *EncodedTable) RootNode() *InternalNode  { return _caa._fe }
func (_eg *FixedSizeTable) RootNode() *InternalNode { return _eg._ceg }
func _fcd(_eb, _dge int32) int32 {
	if _eb > _dge {
		return _eb
	}
	return _dge
}
func (_eec *StandardTable) RootNode() *InternalNode { return _eec._cgd }
func (_fg *StandardTable) String() string           { return _fg._cgd.String() + "\u000a" }
func (_eea *FixedSizeTable) String() string         { return _eea._ceg.String() + "\u000a" }

var _bgb = make([]Tabler, len(_cbc))
