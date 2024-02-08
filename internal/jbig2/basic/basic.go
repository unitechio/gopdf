package basic

import _be "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"

func (_dbe IntSlice) Size() int { return len(_dbe) }
func (_g NumSlice) GetIntSlice() []int {
	_aa := make([]int, len(_g))
	for _cc, _fc := range _g {
		_aa[_cc] = int(_fc)
	}
	return _aa
}
func (_cfc *NumSlice) AddInt(v int) { *_cfc = append(*_cfc, float32(v)) }
func (_ab *IntSlice) Add(v int) error {
	if _ab == nil {
		return _be.Error("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064", "\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	*_ab = append(*_ab, v)
	return nil
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type IntSlice []int

func NewNumSlice(i int) *NumSlice { _ce := NumSlice(make([]float32, i)); return &_ce }
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Sign(v float32) float32 {
	if v >= 0.0 {
		return 1.0
	}
	return -1.0
}
func (_c IntsMap) Add(key uint64, value int) { _c[key] = append(_c[key], value) }
func Ceil(numerator, denominator int) int {
	if numerator%denominator == 0 {
		return numerator / denominator
	}
	return (numerator / denominator) + 1
}
func (_bf *Stack) top() int                             { return len(_bf.Data) - 1 }
func (_bgba *Stack) Peek() (_ffe interface{}, _df bool) { return _bgba.peek() }
func (_ge *Stack) Pop() (_ee interface{}, _cef bool) {
	_ee, _cef = _ge.peek()
	if !_cef {
		return nil, _cef
	}
	_ge.Data = _ge.Data[:_ge.top()]
	return _ee, true
}
func Abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
func (_e IntsMap) Get(key uint64) (int, bool) {
	_f, _d := _e[key]
	if !_d {
		return 0, false
	}
	if len(_f) == 0 {
		return 0, false
	}
	return _f[0], true
}
func (_dg *Stack) Push(v interface{}) { _dg.Data = append(_dg.Data, v) }
func NewIntSlice(i int) *IntSlice     { _cf := IntSlice(make([]int, i)); return &_cf }
func (_ff *Stack) Len() int           { return len(_ff.Data) }
func (_ea NumSlice) GetInt(i int) (int, error) {
	const _bed = "\u0047\u0065\u0074\u0049\u006e\u0074"
	if i < 0 || i > len(_ea)-1 {
		return 0, _be.Errorf(_bed, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_bgb := _ea[i]
	return int(_bgb + Sign(_bgb)*0.5), nil
}
func (_fd *IntSlice) Copy() *IntSlice {
	_fb := IntSlice(make([]int, len(*_fd)))
	copy(_fb, *_fd)
	return &_fb
}
func (_db IntsMap) Delete(key uint64) { delete(_db, key) }
func (_a IntsMap) GetSlice(key uint64) ([]int, bool) {
	_ef, _bb := _a[key]
	if !_bb {
		return nil, false
	}
	return _ef, true
}
func (_ed NumSlice) Get(i int) (float32, error) {
	if i < 0 || i > len(_ed)-1 {
		return 0, _be.Errorf("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _ed[i], nil
}

type Stack struct {
	Data []interface{}
	Aux  *Stack
}

func (_cff IntSlice) Get(index int) (int, error) {
	if index > len(_cff)-1 {
		return 0, _be.Errorf("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _cff[index], nil
}
func (_gf *Stack) peek() (interface{}, bool) {
	_dfd := _gf.top()
	if _dfd == -1 {
		return nil, false
	}
	return _gf.Data[_dfd], true
}

type NumSlice []float32
type IntsMap map[uint64][]int

func (_bg *NumSlice) Add(v float32) { *_bg = append(*_bg, v) }
