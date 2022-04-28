package basic

import _b "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"

func NewNumSlice(i int) *NumSlice { _ab := NumSlice(make([]float32, i)); return &_ab }
func (_be IntsMap) GetSlice(key uint64) ([]int, bool) {
	_a, _ff := _be[key]
	if !_ff {
		return nil, false
	}
	return _a, true
}

type IntSlice []int

func (_da *Stack) Peek() (_ec interface{}, _fcb bool) { return _da.peek() }
func Sign(v float32) float32 {
	if v >= 0.0 {
		return 1.0
	}
	return -1.0
}
func Ceil(numerator, denominator int) int {
	if numerator%denominator == 0 {
		return numerator / denominator
	}
	return (numerator / denominator) + 1
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func (_bef *NumSlice) AddInt(v int) { *_bef = append(*_bef, float32(v)) }
func (_ee *Stack) Pop() (_bb interface{}, _cc bool) {
	_bb, _cc = _ee.peek()
	if !_cc {
		return nil, _cc
	}
	_ee.Data = _ee.Data[:_ee.top()]
	return _bb, true
}

type Stack struct {
	Data []interface{}
	Aux  *Stack
}

func (_ad *NumSlice) Add(v float32) { *_ad = append(*_ad, v) }
func Abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
func (_gg IntSlice) Get(index int) (int, error) {
	if index > len(_gg)-1 {
		return 0, _b.Errorf("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _gg[index], nil
}
func (_dg *Stack) peek() (interface{}, bool) {
	_gcf := _dg.top()
	if _gcf == -1 {
		return nil, false
	}
	return _dg.Data[_gcf], true
}
func NewIntSlice(i int) *IntSlice      { _ac := IntSlice(make([]int, i)); return &_ac }
func (_gc *Stack) Push(v interface{})  { _gc.Data = append(_gc.Data, v) }
func (_cef IntsMap) Delete(key uint64) { delete(_cef, key) }
func (_g *IntSlice) Copy() *IntSlice {
	_bfc := IntSlice(make([]int, len(*_g)))
	copy(_bfc, *_g)
	return &_bfc
}
func (_e NumSlice) Get(i int) (float32, error) {
	if i < 0 || i > len(_e)-1 {
		return 0, _b.Errorf("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _e[i], nil
}

type IntsMap map[uint64][]int
type NumSlice []float32

func (_f IntsMap) Add(key uint64, value int) { _f[key] = append(_f[key], value) }
func (_bd NumSlice) GetInt(i int) (int, error) {
	const _bfe = "\u0047\u0065\u0074\u0049\u006e\u0074"
	if i < 0 || i > len(_bd)-1 {
		return 0, _b.Errorf(_bfe, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_cd := _bd[i]
	return int(_cd + Sign(_cd)*0.5), nil
}
func (_ae IntSlice) Size() int { return len(_ae) }
func (_fc NumSlice) GetIntSlice() []int {
	_cb := make([]int, len(_fc))
	for _aef, _eb := range _fc {
		_cb[_aef] = int(_eb)
	}
	return _cb
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func (_abc *Stack) top() int { return len(_abc.Data) - 1 }
func (_df IntsMap) Get(key uint64) (int, bool) {
	_c, _ce := _df[key]
	if !_ce {
		return 0, false
	}
	if len(_c) == 0 {
		return 0, false
	}
	return _c[0], true
}
func (_gf *Stack) Len() int { return len(_gf.Data) }
func (_bf *IntSlice) Add(v int) error {
	if _bf == nil {
		return _b.Error("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064", "\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	*_bf = append(*_bf, v)
	return nil
}
