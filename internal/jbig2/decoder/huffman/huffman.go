package huffman

import (
	_g "errors"
	_af "fmt"
	_ffg "math"
	_ff "strings"

	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_a "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
)

type InternalNode struct {
	_gbf  int32
	_dfc  Node
	_bcda Node
}

func (_c *EncodedTable) InitTree(codeTable []*Code) error {
	_gda(codeTable)
	for _, _aa := range codeTable {
		if _e := _c._gb.append(_aa); _e != nil {
			return _e
		}
	}
	return nil
}

type Node interface {
	Decode(_fgag *_f.Reader) (int64, error)
	String() string
}

func (_agd *ValueNode) Decode(r *_f.Reader) (int64, error) {
	_gc, _ggf := r.ReadBits(byte(_agd._ead))
	if _ggf != nil {
		return 0, _ggf
	}
	if _agd._edc {
		_gc = -_gc
	}
	return int64(_agd._bcd) + int64(_gc), nil
}
func (_ecg *StandardTable) String() string { return _ecg._bdc.String() + "\u000a" }
func _gda(_ecga []*Code) {
	var _caf int32
	for _, _bgb := range _ecga {
		_caf = _ceb(_caf, _bgb._ada)
	}
	_ccd := make([]int32, _caf+1)
	for _, _agdg := range _ecga {
		_ccd[_agdg._ada]++
	}
	var _fd int32
	_gad := make([]int32, len(_ccd)+1)
	_ccd[0] = 0
	for _cd := int32(1); _cd <= int32(len(_ccd)); _cd++ {
		_gad[_cd] = (_gad[_cd-1] + (_ccd[_cd-1])) << 1
		_fd = _gad[_cd]
		for _, _dab := range _ecga {
			if _dab._ada == _cd {
				_dab._cca = _fd
				_fd++
			}
		}
	}
}

func _bad(_efc, _dg int32) string {
	var _gff int32
	_gba := make([]rune, _dg)
	for _aab := int32(1); _aab <= _dg; _aab++ {
		_gff = _efc >> uint(_dg-_aab) & 1
		if _gff != 0 {
			_gba[_aab-1] = '1'
		} else {
			_gba[_aab-1] = '0'
		}
	}
	return string(_gba)
}

var _ Node = &OutOfBandNode{}

type OutOfBandNode struct{}

func NewFixedSizeTable(codeTable []*Code) (*FixedSizeTable, error) {
	_ed := &FixedSizeTable{_aef: &InternalNode{}}
	if _eg := _ed.InitTree(codeTable); _eg != nil {
		return nil, _eg
	}
	return _ed, nil
}

func (_eb *OutOfBandNode) String() string {
	return _af.Sprintf("\u0025\u0030\u00364\u0062", int64(_ffg.MaxInt64))
}

type ValueNode struct {
	_ead int32
	_bcd int32
	_edc bool
}

func (_ae *EncodedTable) String() string { return _ae._gb.String() + "\u000a" }
func _edec(_ddg int32) *InternalNode     { return &InternalNode{_gbf: _ddg} }
func (_d *EncodedTable) parseTable() error {
	var (
		_fgda          []*Code
		_fga, _ec, _dd int32
		_df            uint64
		_bf            error
	)
	_ad := _d.StreamReader()
	_ba := _d.HtLow()
	for _ba < _d.HtHigh() {
		_df, _bf = _ad.ReadBits(byte(_d.HtPS()))
		if _bf != nil {
			return _bf
		}
		_fga = int32(_df)
		_df, _bf = _ad.ReadBits(byte(_d.HtRS()))
		if _bf != nil {
			return _bf
		}
		_ec = int32(_df)
		_fgda = append(_fgda, NewCode(_fga, _ec, _dd, false))
		_ba += 1 << uint(_ec)
	}
	_df, _bf = _ad.ReadBits(byte(_d.HtPS()))
	if _bf != nil {
		return _bf
	}
	_fga = int32(_df)
	_ec = 32
	_dd = _d.HtLow() - 1
	_fgda = append(_fgda, NewCode(_fga, _ec, _dd, true))
	_df, _bf = _ad.ReadBits(byte(_d.HtPS()))
	if _bf != nil {
		return _bf
	}
	_fga = int32(_df)
	_ec = 32
	_dd = _d.HtHigh()
	_fgda = append(_fgda, NewCode(_fga, _ec, _dd, false))
	if _d.HtOOB() == 1 {
		_df, _bf = _ad.ReadBits(byte(_d.HtPS()))
		if _bf != nil {
			return _bf
		}
		_fga = int32(_df)
		_fgda = append(_fgda, NewCode(_fga, -1, -1, false))
	}
	if _bf = _d.InitTree(_fgda); _bf != nil {
		return _bf
	}
	return nil
}
func (_ef *EncodedTable) RootNode() *InternalNode             { return _ef._gb }
func (_ac *OutOfBandNode) Decode(r *_f.Reader) (int64, error) { return 0, _a.ErrOOB }

var _ Tabler = &EncodedTable{}

func (_aec *StandardTable) Decode(r *_f.Reader) (int64, error) { return _aec._bdc.Decode(r) }

var _ Node = &ValueNode{}

func (_fgge *Code) String() string {
	var _ggg string
	if _fgge._cca != -1 {
		_ggg = _bad(_fgge._cca, _fgge._ada)
	} else {
		_ggg = "\u003f"
	}
	return _af.Sprintf("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _ggg, _fgge._ada, _fgge._ddac, _fgge._gag)
}

func (_bg *InternalNode) append(_bd *Code) (_ce error) {
	if _bd._ada == 0 {
		return nil
	}
	_cag := _bd._ada - 1 - _bg._gbf
	if _cag < 0 {
		return _g.New("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064")
	}
	_ede := (_bd._cca >> uint(_cag)) & 0x1
	if _cag == 0 {
		if _bd._ddac == -1 {
			if _ede == 1 {
				if _bg._bcda != nil {
					return _af.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _bd)
				}
				_bg._bcda = _egg(_bd)
			} else {
				if _bg._dfc != nil {
					return _af.Errorf("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073", _bd)
				}
				_bg._dfc = _egg(_bd)
			}
		} else {
			if _ede == 1 {
				if _bg._bcda != nil {
					return _af.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _bd)
				}
				_bg._bcda = _ca(_bd)
			} else {
				if _bg._dfc != nil {
					return _af.Errorf("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073", _bd)
				}
				_bg._dfc = _ca(_bd)
			}
		}
	} else {
		if _ede == 1 {
			if _bg._bcda == nil {
				_bg._bcda = _edec(_bg._gbf + 1)
			}
			if _ce = _bg._bcda.(*InternalNode).append(_bd); _ce != nil {
				return _ce
			}
		} else {
			if _bg._dfc == nil {
				_bg._dfc = _edec(_bg._gbf + 1)
			}
			if _ce = _bg._dfc.(*InternalNode).append(_bd); _ce != nil {
				return _ce
			}
		}
	}
	return nil
}

type Tabler interface {
	Decode(_dfd *_f.Reader) (int64, error)
	InitTree(_dcg []*Code) error
	String() string
	RootNode() *InternalNode
}

func (_agb *ValueNode) String() string {
	return _af.Sprintf("\u0025\u0064\u002f%\u0064", _agb._ead, _agb._bcd)
}

func _ceb(_bac, _ecb int32) int32 {
	if _bac > _ecb {
		return _bac
	}
	return _ecb
}

func (_ee *InternalNode) String() string {
	_dcd := &_ff.Builder{}
	_dcd.WriteString("\u000a")
	_ee.pad(_dcd)
	_dcd.WriteString("\u0030\u003a\u0020")
	_dcd.WriteString(_ee._dfc.String() + "\u000a")
	_ee.pad(_dcd)
	_dcd.WriteString("\u0031\u003a\u0020")
	_dcd.WriteString(_ee._bcda.String() + "\u000a")
	return _dcd.String()
}

func (_fgdc *InternalNode) Decode(r *_f.Reader) (int64, error) {
	_db, _ga := r.ReadBit()
	if _ga != nil {
		return 0, _ga
	}
	if _db == 0 {
		return _fgdc._dfc.Decode(r)
	}
	return _fgdc._bcda.Decode(r)
}

var _ Node = &InternalNode{}

func (_agf *StandardTable) RootNode() *InternalNode { return _agf._bdc }

type (
	StandardTable struct{ _bdc *InternalNode }
	EncodedTable  struct {
		BasicTabler
		_gb *InternalNode
	}
)

func GetStandardTable(number int) (Tabler, error) {
	if number <= 0 || number > len(_edba) {
		return nil, _g.New("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_gbfg := _edba[number-1]
	if _gbfg == nil {
		var _afb error
		_gbfg, _afb = _gcd(_ffc[number-1])
		if _afb != nil {
			return nil, _afb
		}
		_edba[number-1] = _gbfg
	}
	return _gbfg, nil
}

func _gcd(_bae [][]int32) (*StandardTable, error) {
	var _abg []*Code
	for _fgb := 0; _fgb < len(_bae); _fgb++ {
		_bce := _bae[_fgb][0]
		_edb := _bae[_fgb][1]
		_ade := _bae[_fgb][2]
		var _ddb bool
		if len(_bae[_fgb]) > 3 {
			_ddb = true
		}
		_abg = append(_abg, NewCode(_bce, _edb, _ade, _ddb))
	}
	_bdcg := &StandardTable{_bdc: _edec(0)}
	if _eadb := _bdcg.InitTree(_abg); _eadb != nil {
		return nil, _eadb
	}
	return _bdcg, nil
}

var _edba = make([]Tabler, len(_ffc))

func NewCode(prefixLength, rangeLength, rangeLow int32, isLowerRange bool) *Code {
	return &Code{_ada: prefixLength, _ddac: rangeLength, _gag: rangeLow, _fgg: isLowerRange, _cca: -1}
}
func (_dc *FixedSizeTable) RootNode() *InternalNode            { return _dc._aef }
func (_bc *FixedSizeTable) Decode(r *_f.Reader) (int64, error) { return _bc._aef.Decode(r) }
func (_gg *FixedSizeTable) InitTree(codeTable []*Code) error {
	_gda(codeTable)
	for _, _dda := range codeTable {
		_egd := _gg._aef.append(_dda)
		if _egd != nil {
			return _egd
		}
	}
	return nil
}
func (_fb *FixedSizeTable) String() string { return _fb._aef.String() + "\u000a" }
func _ca(_gd *Code) *ValueNode             { return &ValueNode{_ead: _gd._ddac, _bcd: _gd._gag, _edc: _gd._fgg} }
func (_gfb *InternalNode) pad(_gfc *_ff.Builder) {
	for _de := int32(0); _de < _gfb._gbf; _de++ {
		_gfc.WriteString("\u0020\u0020\u0020")
	}
}
func (_fgd *EncodedTable) Decode(r *_f.Reader) (int64, error) { return _fgd._gb.Decode(r) }
func NewEncodedTable(table BasicTabler) (*EncodedTable, error) {
	_fg := &EncodedTable{_gb: &InternalNode{}, BasicTabler: table}
	if _gbb := _fg.parseTable(); _gbb != nil {
		return nil, _gbb
	}
	return _fg, nil
}

func (_ecc *StandardTable) InitTree(codeTable []*Code) error {
	_gda(codeTable)
	for _, _eca := range codeTable {
		if _bga := _ecc._bdc.append(_eca); _bga != nil {
			return _bga
		}
	}
	return nil
}
func _egg(_ea *Code) *OutOfBandNode { return &OutOfBandNode{} }

type Code struct {
	_ada  int32
	_ddac int32
	_gag  int32
	_fgg  bool
	_cca  int32
}

var _ffc = [][][]int32{{{1, 4, 0}, {2, 8, 16}, {3, 16, 272}, {3, 32, 65808}}, {{1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {6, 32, 75}, {6, -1, 0}}, {{8, 8, -256}, {1, 0, 0}, {2, 0, 1}, {3, 0, 2}, {4, 3, 3}, {5, 6, 11}, {8, 32, -257, 999}, {7, 32, 75}, {6, -1, 0}}, {{1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {5, 32, 76}}, {{7, 8, -255}, {1, 0, 1}, {2, 0, 2}, {3, 0, 3}, {4, 3, 4}, {5, 6, 12}, {7, 32, -256, 999}, {6, 32, 76}}, {{5, 10, -2048}, {4, 9, -1024}, {4, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {2, 7, 0}, {3, 7, 128}, {3, 8, 256}, {4, 9, 512}, {4, 10, 1024}, {6, 32, -2049, 999}, {6, 32, 2048}}, {{4, 9, -1024}, {3, 8, -512}, {4, 7, -256}, {5, 6, -128}, {5, 5, -64}, {4, 5, -32}, {4, 5, 0}, {5, 5, 32}, {5, 6, 64}, {4, 7, 128}, {3, 8, 256}, {3, 9, 512}, {3, 10, 1024}, {5, 32, -1025, 999}, {5, 32, 2048}}, {{8, 3, -15}, {9, 1, -7}, {8, 1, -5}, {9, 0, -3}, {7, 0, -2}, {4, 0, -1}, {2, 1, 0}, {5, 0, 2}, {6, 0, 3}, {3, 4, 4}, {6, 1, 20}, {4, 4, 22}, {4, 5, 38}, {5, 6, 70}, {5, 7, 134}, {6, 7, 262}, {7, 8, 390}, {6, 10, 646}, {9, 32, -16, 999}, {9, 32, 1670}, {2, -1, 0}}, {{8, 4, -31}, {9, 2, -15}, {8, 2, -11}, {9, 1, -7}, {7, 1, -5}, {4, 1, -3}, {3, 1, -1}, {3, 1, 1}, {5, 1, 3}, {6, 1, 5}, {3, 5, 7}, {6, 2, 39}, {4, 5, 43}, {4, 6, 75}, {5, 7, 139}, {5, 8, 267}, {6, 8, 523}, {7, 9, 779}, {6, 11, 1291}, {9, 32, -32, 999}, {9, 32, 3339}, {2, -1, 0}}, {{7, 4, -21}, {8, 0, -5}, {7, 0, -4}, {5, 0, -3}, {2, 2, -2}, {5, 0, 2}, {6, 0, 3}, {7, 0, 4}, {8, 0, 5}, {2, 6, 6}, {5, 5, 70}, {6, 5, 102}, {6, 6, 134}, {6, 7, 198}, {6, 8, 326}, {6, 9, 582}, {6, 10, 1094}, {7, 11, 2118}, {8, 32, -22, 999}, {8, 32, 4166}, {2, -1, 0}}, {{1, 0, 1}, {2, 1, 2}, {4, 0, 4}, {4, 1, 5}, {5, 1, 7}, {5, 2, 9}, {6, 2, 13}, {7, 2, 17}, {7, 3, 21}, {7, 4, 29}, {7, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{1, 0, 1}, {2, 0, 2}, {3, 1, 3}, {5, 0, 5}, {5, 1, 6}, {6, 1, 8}, {7, 0, 10}, {7, 1, 11}, {7, 2, 13}, {7, 3, 17}, {7, 4, 25}, {8, 5, 41}, {8, 32, 73}}, {{1, 0, 1}, {3, 0, 2}, {4, 0, 3}, {5, 0, 4}, {4, 1, 5}, {3, 3, 7}, {6, 1, 15}, {6, 2, 17}, {6, 3, 21}, {6, 4, 29}, {6, 5, 45}, {7, 6, 77}, {7, 32, 141}}, {{3, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {3, 0, 2}}, {{7, 4, -24}, {6, 2, -8}, {5, 1, -4}, {4, 0, -2}, {3, 0, -1}, {1, 0, 0}, {3, 0, 1}, {4, 0, 2}, {5, 1, 3}, {6, 2, 5}, {7, 4, 9}, {7, 32, -25, 999}, {7, 32, 25}}}

type (
	FixedSizeTable struct{ _aef *InternalNode }
	BasicTabler    interface {
		HtHigh() int32
		HtLow() int32
		StreamReader() *_f.Reader
		HtPS() int32
		HtRS() int32
		HtOOB() int32
	}
)
