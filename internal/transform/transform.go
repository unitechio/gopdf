package transform

import (
	_g "fmt"
	_a "math"

	_f "bitbucket.org/shenghui0779/gopdf/common"
)

func IdentityMatrix() Matrix { return NewMatrix(1, 0, 0, 1, 0, 0) }

const _gd = 1e-10

func NewMatrixFromTransforms(xScale, yScale, theta, tx, ty float64) Matrix {
	return IdentityMatrix().Scale(xScale, yScale).Rotate(theta).Translate(tx, ty)
}
func ScaleMatrix(x, y float64) Matrix              { return NewMatrix(x, 0, 0, y, 0, 0) }
func (_da Matrix) Translate(tx, ty float64) Matrix { return _da.Mult(TranslationMatrix(tx, ty)) }

type Matrix [9]float64

func (_ea Matrix) Round(precision float64) Matrix {
	for _d := range _ea {
		_ea[_d] = _a.Round(_ea[_d]/precision) * precision
	}
	return _ea
}

func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	_bf := Matrix{a, b, 0, c, d, 0, tx, ty, 1}
	_bf.clampRange()
	return _bf
}

type Point struct {
	X float64
	Y float64
}

func (_gbe Matrix) Angle() float64 {
	_fa := _a.Atan2(-_gbe[1], _gbe[0])
	if _fa < 0.0 {
		_fa += 2 * _a.Pi
	}
	return _fa / _a.Pi * 180.0
}

func (_ag *Matrix) Set(a, b, c, d, tx, ty float64) {
	_ag[0], _ag[1] = a, b
	_ag[3], _ag[4] = c, d
	_ag[6], _ag[7] = tx, ty
	_ag.clampRange()
}

func (_cd Point) Interpolate(b Point, t float64) Point {
	return Point{X: (1-t)*_cd.X + t*b.X, Y: (1-t)*_cd.Y + t*b.Y}
}
func (_dg Point) Displace(delta Point) Point        { return Point{_dg.X + delta.X, _dg.Y + delta.Y} }
func (_caa Matrix) Translation() (float64, float64) { return _caa[6], _caa[7] }
func (_ga Matrix) ScalingFactorY() float64          { return _a.Hypot(_ga[3], _ga[4]) }
func (_fbeg *Point) Transform(a, b, c, d, tx, ty float64) {
	_cb := NewMatrix(a, b, c, d, tx, ty)
	_fbeg.transformByMatrix(_cb)
}
func (_ee Matrix) Singular() bool          { return _a.Abs(_ee[0]*_ee[4]-_ee[1]*_ee[3]) < _gd }
func (_bg Matrix) ScalingFactorX() float64 { return _a.Hypot(_bg[0], _bg[1]) }
func (_fc *Matrix) clampRange() {
	for _gfb, _cge := range _fc {
		if _cge > _dab {
			_f.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _cge, _dab)
			_fc[_gfb] = _dab
		} else if _cge < -_dab {
			_f.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _cge, -_dab)
			_fc[_gfb] = -_dab
		}
	}
}

const _dab = 1e9

func (_fbe *Matrix) Clone() Matrix {
	return NewMatrix(_fbe[0], _fbe[1], _fbe[3], _fbe[4], _fbe[6], _fbe[7])
}
func (_daf *Matrix) Shear(x, y float64) { _daf.Concat(ShearMatrix(x, y)) }
func (_ca Matrix) Mult(b Matrix) Matrix {
	_ca.Concat(b)
	return _ca
}
func (_bff *Point) transformByMatrix(_fgd Matrix) { _bff.X, _bff.Y = _fgd.Transform(_bff.X, _bff.Y) }
func (_dcb Matrix) Rotate(theta float64) Matrix   { return _dcb.Mult(RotationMatrix(theta)) }
func (_c *Matrix) Concat(b Matrix) {
	*_c = Matrix{b[0]*_c[0] + b[1]*_c[3], b[0]*_c[1] + b[1]*_c[4], 0, b[3]*_c[0] + b[4]*_c[3], b[3]*_c[1] + b[4]*_c[4], 0, b[6]*_c[0] + b[7]*_c[3] + _c[6], b[6]*_c[1] + b[7]*_c[4] + _c[7], 1}
	_c.clampRange()
}

const _be = 1e-6

func (_fg Matrix) Transform(x, y float64) (float64, float64) {
	_cg := x*_fg[0] + y*_fg[3] + _fg[6]
	_bb := x*_fg[1] + y*_fg[4] + _fg[7]
	return _cg, _bb
}
func (_ef Matrix) Scale(xScale, yScale float64) Matrix { return _ef.Mult(ScaleMatrix(xScale, yScale)) }
func (_fga Point) Distance(b Point) float64            { return _a.Hypot(_fga.X-b.X, _fga.Y-b.Y) }
func (_fb Matrix) String() string {
	_fd, _dc, _eb, _eaa, _gb, _eab := _fb[0], _fb[1], _fb[3], _fb[4], _fb[6], _fb[7]
	return _g.Sprintf("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d", _fd, _dc, _eb, _eaa, _gb, _eab)
}

func (_ce Point) Rotate(theta float64) Point {
	_fdb := _a.Hypot(_ce.X, _ce.Y)
	_af := _a.Atan2(_ce.Y, _ce.X)
	_aaf, _eaac := _a.Sincos(_af + theta/180.0*_a.Pi)
	return Point{_fdb * _eaac, _fdb * _aaf}
}

const _egc = 1.0e-6

func (_b Matrix) Identity() bool {
	return _b[0] == 1 && _b[1] == 0 && _b[2] == 0 && _b[3] == 0 && _b[4] == 1 && _b[5] == 0 && _b[6] == 0 && _b[7] == 0 && _b[8] == 1
}

func (_bgf Matrix) Inverse() (Matrix, bool) {
	_ac, _ead := _bgf[0], _bgf[1]
	_cc, _aa := _bgf[3], _bgf[4]
	_dd, _ebf := _bgf[6], _bgf[7]
	_ae := _ac*_aa - _ead*_cc
	if _a.Abs(_ae) < _egc {
		return Matrix{}, false
	}
	_gf, _daa := _aa/_ae, -_ead/_ae
	_ebe, _ccg := -_cc/_ae, _ac/_ae
	_fdc := -(_gf*_dd + _ebe*_ebf)
	_gc := -(_daa*_dd + _ccg*_ebf)
	return NewMatrix(_gf, _daa, _ebe, _ccg, _fdc, _gc), true
}
func ShearMatrix(x, y float64) Matrix { return NewMatrix(1, y, x, 1, 0, 0) }
func RotationMatrix(angle float64) Matrix {
	_gg := _a.Cos(angle)
	_db := _a.Sin(angle)
	return NewMatrix(_gg, _db, -_db, _gg, 0, 0)
}

func (_ggb Matrix) Unrealistic() bool {
	_ba, _bac, _cgc, _gcg := _a.Abs(_ggb[0]), _a.Abs(_ggb[1]), _a.Abs(_ggb[3]), _a.Abs(_ggb[4])
	_eg := _ba > _be && _gcg > _be
	_bbb := _bac > _be && _cgc > _be
	return !(_eg || _bbb)
}

func (_fca Point) String() string {
	return _g.Sprintf("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029", _fca.X, _fca.Y)
}
func NewPoint(x, y float64) Point             { return Point{X: x, Y: y} }
func TranslationMatrix(tx, ty float64) Matrix { return NewMatrix(1, 0, 0, 1, tx, ty) }
func (_eag *Point) Set(x, y float64)          { _eag.X, _eag.Y = x, y }
