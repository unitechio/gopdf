package basic

import _d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"

func (_bgc *NumSlice) Add(v float32) { *_bgc = append(*_bgc, v) }
func (_a IntsMap) Get(key uint64) (int, bool) {
	_eg, _ege := _a[key]
	if !_ege {
		return 0, false
	}
	if len(_eg) == 0 {
		return 0, false
	}
	return _eg[0], true
}

func (_be IntsMap) GetSlice(key uint64) ([]int, bool) {
	_g, _f := _be[key]
	if !_f {
		return nil, false
	}
	return _g, true
}
func (_gb *Stack) Len() int    { return len(_gb.Data) }
func (_bb IntSlice) Size() int { return len(_bb) }
func (_ae *IntSlice) Add(v int) error {
	if _ae == nil {
		return _d.Error("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064", "\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	*_ae = append(*_ae, v)
	return nil
}

func (_bg *IntSlice) Copy() *IntSlice {
	_ge := IntSlice(make([]int, len(*_bg)))
	copy(_ge, *_bg)
	return &_ge
}
func (_cb *Stack) top() int                           { return len(_cb.Data) - 1 }
func (_ca *Stack) Push(v interface{})                 { _ca.Data = append(_ca.Data, v) }
func (_aa IntsMap) Delete(key uint64)                 { delete(_aa, key) }
func (_e IntsMap) Add(key uint64, value int)          { _e[key] = append(_e[key], value) }
func (_fbe *Stack) Peek() (_fg interface{}, _da bool) { return _fbe.peek() }
func (_bba *Stack) Pop() (_de interface{}, _ad bool) {
	_de, _ad = _bba.peek()
	if !_ad {
		return nil, _ad
	}
	_bba.Data = _bba.Data[:_bba.top()]
	return _de, true
}
func NewIntSlice(i int) *IntSlice  { _ea := IntSlice(make([]int, i)); return &_ea }
func NewNumSlice(i int) *NumSlice  { _cg := NumSlice(make([]float32, i)); return &_cg }
func (_fe *NumSlice) AddInt(v int) { *_fe = append(*_fe, float32(v)) }
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type IntsMap map[uint64][]int

func (_cf NumSlice) GetIntSlice() []int {
	_aeb := make([]int, len(_cf))
	for _ef, _gg := range _cf {
		_aeb[_ef] = int(_gg)
	}
	return _aeb
}

func (_dc NumSlice) Get(i int) (float32, error) {
	if i < 0 || i > len(_dc)-1 {
		return 0, _d.Errorf("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _dc[i], nil
}

type NumSlice []float32

func Abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

type IntSlice []int

func (_fb NumSlice) GetInt(i int) (int, error) {
	const _ec = "\u0047\u0065\u0074\u0049\u006e\u0074"
	if i < 0 || i > len(_fb)-1 {
		return 0, _d.Errorf(_ec, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_ff := _fb[i]
	return int(_ff + Sign(_ff)*0.5), nil
}

func Sign(v float32) float32 {
	if v >= 0.0 {
		return 1.0
	}
	return -1.0
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (_c IntSlice) Get(index int) (int, error) {
	if index > len(_c)-1 {
		return 0, _d.Errorf("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _c[index], nil
}

type Stack struct {
	Data []interface{}
	Aux  *Stack
}

func Ceil(numerator, denominator int) int {
	if numerator%denominator == 0 {
		return numerator / denominator
	}
	return (numerator / denominator) + 1
}

func (_gd *Stack) peek() (interface{}, bool) {
	_dae := _gd.top()
	if _dae == -1 {
		return nil, false
	}
	return _gd.Data[_dae], true
}
