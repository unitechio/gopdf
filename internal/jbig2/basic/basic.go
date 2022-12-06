package basic

import _ef "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"

func (_ee *IntSlice) Add(v int) error {
	if _ee == nil {
		return _ef.Error("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064", "\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	*_ee = append(*_ee, v)
	return nil
}
func Abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
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

type IntSlice []int

func (_eg NumSlice) Get(i int) (float32, error) {
	if i < 0 || i > len(_eg)-1 {
		return 0, _ef.Errorf("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _eg[i], nil
}
func (_f IntsMap) Add(key uint64, value int) { _f[key] = append(_f[key], value) }
func (_db *Stack) Len() int                  { return len(_db.Data) }
func Ceil(numerator, denominator int) int {
	if numerator%denominator == 0 {
		return numerator / denominator
	}
	return (numerator / denominator) + 1
}
func (_fe IntsMap) Delete(key uint64) { delete(_fe, key) }
func (_gb NumSlice) GetIntSlice() []int {
	_gga := make([]int, len(_gb))
	for _af, _ae := range _gb {
		_gga[_af] = int(_ae)
	}
	return _gga
}
func (_ga *Stack) Pop() (_cc interface{}, _dg bool) {
	_cc, _dg = _ga.peek()
	if !_dg {
		return nil, _dg
	}
	_ga.Data = _ga.Data[:_ga.top()]
	return _cc, true
}
func (_ccb *Stack) top() int         { return len(_ccb.Data) - 1 }
func (_fgf *NumSlice) Add(v float32) { *_fgf = append(*_fgf, v) }
func (_fga IntSlice) Size() int      { return len(_fga) }
func NewNumSlice(i int) *NumSlice    { _gg := NumSlice(make([]float32, i)); return &_gg }

type NumSlice []float32
type IntsMap map[uint64][]int

func (_fb *NumSlice) AddInt(v int)                   { *_fb = append(*_fb, float32(v)) }
func (_gf *Stack) Push(v interface{})                { _gf.Data = append(_gf.Data, v) }
func (_cd *Stack) Peek() (_bg interface{}, _df bool) { return _cd.peek() }
func (_c IntsMap) Get(key uint64) (int, bool) {
	_a, _fd := _c[key]
	if !_fd {
		return 0, false
	}
	if len(_a) == 0 {
		return 0, false
	}
	return _a[0], true
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Stack struct {
	Data []interface{}
	Aux  *Stack
}

func NewIntSlice(i int) *IntSlice { _fc := IntSlice(make([]int, i)); return &_fc }
func (_ede *Stack) peek() (interface{}, bool) {
	_ge := _ede.top()
	if _ge == -1 {
		return nil, false
	}
	return _ede.Data[_ge], true
}
func (_d IntsMap) GetSlice(key uint64) ([]int, bool) {
	_fg, _cg := _d[key]
	if !_cg {
		return nil, false
	}
	return _fg, true
}
func (_bd NumSlice) GetInt(i int) (int, error) {
	const _bb = "\u0047\u0065\u0074\u0049\u006e\u0074"
	if i < 0 || i > len(_bd)-1 {
		return 0, _ef.Errorf(_bb, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_ed := _bd[i]
	return int(_ed + Sign(_ed)*0.5), nil
}
func (_g *IntSlice) Copy() *IntSlice {
	_ec := IntSlice(make([]int, len(*_g)))
	copy(_ec, *_g)
	return &_ec
}
func (_b IntSlice) Get(index int) (int, error) {
	if index > len(_b)-1 {
		return 0, _ef.Errorf("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _b[index], nil
}
