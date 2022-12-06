package transform

import (
	_g "fmt"
	_a "math"

	_e "bitbucket.org/shenghui0779/gopdf/common"
)

func (_df Matrix) Singular() bool                    { return _a.Abs(_df[0]*_df[4]-_df[1]*_df[3]) < _efg }
func (_cgda Matrix) Translation() (float64, float64) { return _cgda[6], _cgda[7] }
func (_cff Point) Interpolate(b Point, t float64) Point {
	return Point{X: (1-t)*_cff.X + t*b.X, Y: (1-t)*_cff.Y + t*b.Y}
}
func (_eg Matrix) Round(precision float64) Matrix {
	for _c := range _eg {
		_eg[_c] = _a.Round(_eg[_c]/precision) * precision
	}
	return _eg
}
func (_ecd *Point) Transform(a, b, c, d, tx, ty float64) {
	_dcb := NewMatrix(a, b, c, d, tx, ty)
	_ecd.transformByMatrix(_dcb)
}
func (_cg Matrix) Scale(xScale, yScale float64) Matrix { return _cg.Mult(ScaleMatrix(xScale, yScale)) }
func (_d *Matrix) Clone() Matrix                       { return NewMatrix(_d[0], _d[1], _d[3], _d[4], _d[6], _d[7]) }
func (_ege Matrix) String() string {
	_cf, _be, _cc, _bga, _ee, _f := _ege[0], _ege[1], _ege[3], _ege[4], _ege[6], _ege[7]
	return _g.Sprintf("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d", _cf, _be, _cc, _bga, _ee, _f)
}
func NewMatrixFromTransforms(xScale, yScale, theta, tx, ty float64) Matrix {
	return IdentityMatrix().Scale(xScale, yScale).Rotate(theta).Translate(tx, ty)
}
func NewPoint(x, y float64) Point             { return Point{X: x, Y: y} }
func (_dca Point) Displace(delta Point) Point { return Point{_dca.X + delta.X, _dca.Y + delta.Y} }

const _ac = 1.0e-6

func TranslationMatrix(tx, ty float64) Matrix { return NewMatrix(1, 0, 0, 1, tx, ty) }

type Matrix [9]float64

func (_gde Point) String() string {
	return _g.Sprintf("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029", _gde.X, _gde.Y)
}
func (_bgf Matrix) Angle() float64 {
	_ff := _a.Atan2(-_bgf[1], _bgf[0])
	if _ff < 0.0 {
		_ff += 2 * _a.Pi
	}
	return _ff / _a.Pi * 180.0
}

type Point struct {
	X float64
	Y float64
}

func (_af Matrix) Inverse() (Matrix, bool) {
	_fa, _ed := _af[0], _af[1]
	_eb, _gf := _af[3], _af[4]
	_ec, _dff := _af[6], _af[7]
	_dfe := _fa*_gf - _ed*_eb
	if _a.Abs(_dfe) < _ac {
		return Matrix{}, false
	}
	_ba, _fae := _gf/_dfe, -_ed/_dfe
	_fd, _efc := -_eb/_dfe, _fa/_dfe
	_ca := -(_ba*_ec + _fd*_dff)
	_da := -(_fae*_ec + _efc*_dff)
	return NewMatrix(_ba, _fae, _fd, _efc, _ca, _da), true
}

const _efg = 1e-10

func (_bb *Matrix) Shear(x, y float64) { _bb.Concat(ShearMatrix(x, y)) }
func RotationMatrix(angle float64) Matrix {
	_gd := _a.Cos(angle)
	_ef := _a.Sin(angle)
	return NewMatrix(_gd, _ef, -_ef, _gd, 0, 0)
}
func (_fgc Point) Rotate(theta float64) Point {
	_db := _a.Hypot(_fgc.X, _fgc.Y)
	_efb := _a.Atan2(_fgc.Y, _fgc.X)
	_cad, _efd := _a.Sincos(_efb + theta/180.0*_a.Pi)
	return Point{_db * _efd, _db * _cad}
}
func (_bc Matrix) ScalingFactorY() float64      { return _a.Hypot(_bc[3], _bc[4]) }
func (_cgd Matrix) Rotate(theta float64) Matrix { return _cgd.Mult(RotationMatrix(theta)) }
func (_dc Matrix) Transform(x, y float64) (float64, float64) {
	_eea := x*_dc[0] + y*_dc[3] + _dc[6]
	_fb := x*_dc[1] + y*_dc[4] + _dc[7]
	return _eea, _fb
}
func (_ga *Matrix) Concat(b Matrix) {
	*_ga = Matrix{b[0]*_ga[0] + b[1]*_ga[3], b[0]*_ga[1] + b[1]*_ga[4], 0, b[3]*_ga[0] + b[4]*_ga[3], b[3]*_ga[1] + b[4]*_ga[4], 0, b[6]*_ga[0] + b[7]*_ga[3] + _ga[6], b[6]*_ga[1] + b[7]*_ga[4] + _ga[7], 1}
	_ga.clampRange()
}
func (_gc *Matrix) clampRange() {
	for _fe, _gb := range _gc {
		if _gb > _faf {
			_e.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _gb, _faf)
			_gc[_fe] = _faf
		} else if _gb < -_faf {
			_e.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _gb, -_faf)
			_gc[_fe] = -_faf
		}
	}
}
func (_cd Matrix) Translate(tx, ty float64) Matrix { return _cd.Mult(TranslationMatrix(tx, ty)) }
func (_bgb Matrix) ScalingFactorX() float64        { return _a.Hypot(_bgb[0], _bgb[1]) }
func (_aa Matrix) Identity() bool {
	return _aa[0] == 1 && _aa[1] == 0 && _aa[2] == 0 && _aa[3] == 0 && _aa[4] == 1 && _aa[5] == 0 && _aa[6] == 0 && _aa[7] == 0 && _aa[8] == 1
}
func ScaleMatrix(x, y float64) Matrix       { return NewMatrix(x, 0, 0, y, 0, 0) }
func (_cde Point) Distance(b Point) float64 { return _a.Hypot(_cde.X-b.X, _cde.Y-b.Y) }
func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	_bg := Matrix{a, b, 0, c, d, 0, tx, ty, 1}
	_bg.clampRange()
	return _bg
}
func ShearMatrix(x, y float64) Matrix { return NewMatrix(1, y, x, 1, 0, 0) }
func (_cdg *Point) Set(x, y float64)  { _cdg.X, _cdg.Y = x, y }
func (_cga *Matrix) Set(a, b, c, d, tx, ty float64) {
	_cga[0], _cga[1] = a, b
	_cga[3], _cga[4] = c, d
	_cga[6], _cga[7] = tx, ty
	_cga.clampRange()
}
func IdentityMatrix() Matrix                       { return NewMatrix(1, 0, 0, 1, 0, 0) }
func (_eff *Point) transformByMatrix(_fabc Matrix) { _eff.X, _eff.Y = _fabc.Transform(_eff.X, _eff.Y) }
func (_gg Matrix) Mult(b Matrix) Matrix {
	_gg.Concat(b)
	return _gg
}

const _faf = 1e9

func (_bd Matrix) Unrealistic() bool {
	_fg, _dcd, _ag, _ge := _a.Abs(_bd[0]), _a.Abs(_bd[1]), _a.Abs(_bd[3]), _a.Abs(_bd[4])
	_agb := _fg > _ecg && _ge > _ecg
	_fab := _dcd > _ecg && _ag > _ecg
	return !(_agb || _fab)
}

const _ecg = 1e-6
