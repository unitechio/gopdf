package transform

import (
	_a "fmt"
	_f "math"

	_cd "bitbucket.org/shenghui0779/gopdf/common"
)

const _db = 1e-10

func (_agb *Matrix) Set(a, b, c, d, tx, ty float64) {
	_agb[0], _agb[1] = a, b
	_agb[3], _agb[4] = c, d
	_agb[6], _agb[7] = tx, ty
	_agb.clampRange()
}

const _ecf = 1e9

func (_fea Matrix) Scale(xScale, yScale float64) Matrix {
	return _fea.Mult(ScaleMatrix(xScale, yScale))
}
func (_gf *Matrix) clampRange() {
	for _de, _fa := range _gf {
		if _fa > _ecf {
			_cd.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _fa, _ecf)
			_gf[_de] = _ecf
		} else if _fa < -_ecf {
			_cd.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _fa, -_ecf)
			_gf[_de] = -_ecf
		}
	}
}

type Matrix [9]float64

func (_dcfb *Point) Set(x, y float64) { _dcfb.X, _dcfb.Y = x, y }
func (_e Matrix) Round(precision float64) Matrix {
	for _ag := range _e {
		_e[_ag] = _f.Round(_e[_ag]/precision) * precision
	}
	return _e
}
func (_cc Matrix) String() string {
	_ad, _b, _bf, _fe, _fee, _eb := _cc[0], _cc[1], _cc[3], _cc[4], _cc[6], _cc[7]
	return _a.Sprintf("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d", _ad, _b, _bf, _fe, _fee, _eb)
}
func (_bd *Point) Transform(a, b, c, d, tx, ty float64) {
	_eba := NewMatrix(a, b, c, d, tx, ty)
	_bd.transformByMatrix(_eba)
}
func (_dc Matrix) ScalingFactorY() float64 { return _f.Hypot(_dc[3], _dc[4]) }
func (_cef Point) String() string {
	return _a.Sprintf("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029", _cef.X, _cef.Y)
}
func (_ef Matrix) Identity() bool {
	return _ef[0] == 1 && _ef[1] == 0 && _ef[2] == 0 && _ef[3] == 0 && _ef[4] == 1 && _ef[5] == 0 && _ef[6] == 0 && _ef[7] == 0 && _ef[8] == 1
}
func (_fb Matrix) Mult(b Matrix) Matrix { _fb.Concat(b); return _fb }
func (_dba Matrix) Unrealistic() bool {
	_agd, _bga, _dee, _gag := _f.Abs(_dba[0]), _f.Abs(_dba[1]), _f.Abs(_dba[3]), _f.Abs(_dba[4])
	_dg := _agd > _bgf && _gag > _bgf
	_be := _bga > _bgf && _dee > _bgf
	return !(_dg || _be)
}
func (_efb Matrix) Translation() (float64, float64) { return _efb[6], _efb[7] }
func ScaleMatrix(x, y float64) Matrix               { return NewMatrix(x, 0, 0, y, 0, 0) }
func (_ebg Matrix) Singular() bool                  { return _f.Abs(_ebg[0]*_ebg[4]-_ebg[1]*_ebg[3]) < _db }
func (_abf *Matrix) Clone() Matrix {
	return NewMatrix(_abf[0], _abf[1], _abf[3], _abf[4], _abf[6], _abf[7])
}

const _bgf = 1e-6

func (_ab Matrix) Rotate(theta float64) Matrix { return _ab.Mult(RotationMatrix(theta)) }
func (_bg Matrix) ScalingFactorX() float64     { return _f.Hypot(_bg[0], _bg[1]) }
func (_ge Point) Rotate(theta float64) Point {
	_efc := _f.Hypot(_ge.X, _ge.Y)
	_da := _f.Atan2(_ge.Y, _ge.X)
	_dec, _cde := _f.Sincos(_da + theta/180.0*_f.Pi)
	return Point{_efc * _cde, _efc * _dec}
}
func (_ee Point) Displace(delta Point) Point      { return Point{_ee.X + delta.X, _ee.Y + delta.Y} }
func TranslationMatrix(tx, ty float64) Matrix     { return NewMatrix(1, 0, 0, 1, tx, ty) }
func (_adf *Point) transformByMatrix(_bdg Matrix) { _adf.X, _adf.Y = _bdg.Transform(_adf.X, _adf.Y) }
func ShearMatrix(x, y float64) Matrix             { return NewMatrix(1, y, x, 1, 0, 0) }
func (_ea *Matrix) Shear(x, y float64)            { _ea.Concat(ShearMatrix(x, y)) }
func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	_g := Matrix{a, b, 0, c, d, 0, tx, ty, 1}
	_g.clampRange()
	return _g
}
func (_ga Matrix) Transform(x, y float64) (float64, float64) {
	_df := x*_ga[0] + y*_ga[3] + _ga[6]
	_ada := x*_ga[1] + y*_ga[4] + _ga[7]
	return _df, _ada
}
func (_aegd Point) Distance(b Point) float64 { return _f.Hypot(_aegd.X-b.X, _aegd.Y-b.Y) }
func NewMatrixFromTransforms(xScale, yScale, theta, tx, ty float64) Matrix {
	return IdentityMatrix().Scale(xScale, yScale).Rotate(theta).Translate(tx, ty)
}
func (_dcf Matrix) Inverse() (Matrix, bool) {
	_ec, _gc := _dcf[0], _dcf[1]
	_dd, _dbc := _dcf[3], _dcf[4]
	_fba, _ae := _dcf[6], _dcf[7]
	_bge := _ec*_dbc - _gc*_dd
	if _f.Abs(_bge) < _agc {
		return Matrix{}, false
	}
	_bff, _ff := _dbc/_bge, -_gc/_bge
	_ecc, _cg := -_dd/_bge, _ec/_bge
	_aeg := -(_bff*_fba + _ecc*_ae)
	_aa := -(_ff*_fba + _cg*_ae)
	return NewMatrix(_bff, _ff, _ecc, _cg, _aeg, _aa), true
}
func RotationMatrix(angle float64) Matrix {
	_d := _f.Cos(angle)
	_fc := _f.Sin(angle)
	return NewMatrix(_d, _fc, -_fc, _d, 0, 0)
}
func IdentityMatrix() Matrix { return NewMatrix(1, 0, 0, 1, 0, 0) }

type Point struct {
	X float64
	Y float64
}

func (_ce *Matrix) Concat(b Matrix) {
	*_ce = Matrix{b[0]*_ce[0] + b[1]*_ce[3], b[0]*_ce[1] + b[1]*_ce[4], 0, b[3]*_ce[0] + b[4]*_ce[3], b[3]*_ce[1] + b[4]*_ce[4], 0, b[6]*_ce[0] + b[7]*_ce[3] + _ce[6], b[6]*_ce[1] + b[7]*_ce[4] + _ce[7], 1}
	_ce.clampRange()
}
func (_eg Matrix) Angle() float64 {
	_bc := _f.Atan2(-_eg[1], _eg[0])
	if _bc < 0.0 {
		_bc += 2 * _f.Pi
	}
	return _bc / _f.Pi * 180.0
}
func NewPoint(x, y float64) Point { return Point{X: x, Y: y} }
func (_ffc Point) Interpolate(b Point, t float64) Point {
	return Point{X: (1-t)*_ffc.X + t*b.X, Y: (1-t)*_ffc.Y + t*b.Y}
}
func (_bfa Matrix) Translate(tx, ty float64) Matrix { return _bfa.Mult(TranslationMatrix(tx, ty)) }

const _agc = 1.0e-6
