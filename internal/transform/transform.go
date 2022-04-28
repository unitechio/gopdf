package transform

import (
	_a "fmt"
	_c "math"

	_e "bitbucket.org/shenghui0779/gopdf/common"
)

const _da = 1e-6

func IdentityMatrix() Matrix                       { return NewMatrix(1, 0, 0, 1, 0, 0) }
func TranslationMatrix(tx, ty float64) Matrix      { return NewMatrix(1, 0, 0, 1, tx, ty) }
func (_dd Matrix) Translate(tx, ty float64) Matrix { return _dd.Mult(TranslationMatrix(tx, ty)) }
func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	_d := Matrix{a, b, 0, c, d, 0, tx, ty, 1}
	_d.clampRange()
	return _d
}
func (_bdg *Matrix) clampRange() {
	for _dff, _gdb := range _bdg {
		if _gdb > _gb {
			_e.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _gdb, _gb)
			_bdg[_dff] = _gb
		} else if _gdb < -_gb {
			_e.Log.Debug("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067", _gdb, -_gb)
			_bdg[_dff] = -_gb
		}
	}
}

const _eg = 1e-10

func (_bg *Matrix) Concat(b Matrix) {
	*_bg = Matrix{b[0]*_bg[0] + b[1]*_bg[3], b[0]*_bg[1] + b[1]*_bg[4], 0, b[3]*_bg[0] + b[4]*_bg[3], b[3]*_bg[1] + b[4]*_bg[4], 0, b[6]*_bg[0] + b[7]*_bg[3] + _bg[6], b[6]*_bg[1] + b[7]*_bg[4] + _bg[7], 1}
	_bg.clampRange()
}
func (_fbe Point) Interpolate(b Point, t float64) Point {
	return Point{X: (1-t)*_fbe.X + t*b.X, Y: (1-t)*_fbe.Y + t*b.Y}
}
func (_dc *Matrix) Shear(x, y float64) { _dc.Concat(ShearMatrix(x, y)) }

type Point struct {
	X float64
	Y float64
}

func ShearMatrix(x, y float64) Matrix      { return NewMatrix(1, y, x, 1, 0, 0) }
func (_gf Matrix) ScalingFactorX() float64 { return _c.Hypot(_gf[0], _gf[1]) }
func (_fd Matrix) Transform(x, y float64) (float64, float64) {
	_ebc := x*_fd[0] + y*_fd[3] + _fd[6]
	_df := x*_fd[1] + y*_fd[4] + _fd[7]
	return _ebc, _df
}
func (_eb Matrix) Identity() bool {
	return _eb[0] == 1 && _eb[1] == 0 && _eb[2] == 0 && _eb[3] == 0 && _eb[4] == 1 && _eb[5] == 0 && _eb[6] == 0 && _eb[7] == 0 && _eb[8] == 1
}
func RotationMatrix(angle float64) Matrix {
	_bd := _c.Cos(angle)
	_ebd := _c.Sin(angle)
	return NewMatrix(_bd, _ebd, -_ebd, _bd, 0, 0)
}

const _edb = 1.0e-6

func (_bb Matrix) Scale(xScale, yScale float64) Matrix { return _bb.Mult(ScaleMatrix(xScale, yScale)) }

const _gb = 1e9

func (_ef Matrix) Inverse() (Matrix, bool) {
	_ee, _dbe := _ef[0], _ef[1]
	_egd, _aac := _ef[3], _ef[4]
	_af, _fe := _ef[6], _ef[7]
	_afd := _ee*_aac - _dbe*_egd
	if _c.Abs(_afd) < _edb {
		return Matrix{}, false
	}
	_ca, _ed := _aac/_afd, -_dbe/_afd
	_eee, _dbb := -_egd/_afd, _ee/_afd
	_fg := -(_ca*_af + _eee*_fe)
	_bab := -(_ed*_af + _dbb*_fe)
	return NewMatrix(_ca, _ed, _eee, _dbb, _fg, _bab), true
}
func (_bba *Matrix) Set(a, b, c, d, tx, ty float64) {
	_bba[0], _bba[1] = a, b
	_bba[3], _bba[4] = c, d
	_bba[6], _bba[7] = tx, ty
	_bba.clampRange()
}
func (_bgg Matrix) Mult(b Matrix) Matrix {
	_bgg.Concat(b)
	return _bgg
}
func (_ecb Point) Displace(delta Point) Point { return Point{_ecb.X + delta.X, _ecb.Y + delta.Y} }
func (_b Matrix) Round(precision float64) Matrix {
	for _gc := range _b {
		_b[_gc] = _c.Round(_b[_gc]/precision) * precision
	}
	return _b
}
func (_ddg *Point) transformByMatrix(_ac Matrix) { _ddg.X, _ddg.Y = _ac.Transform(_ddg.X, _ddg.Y) }
func (_ea Matrix) Angle() float64 {
	_be := _c.Atan2(-_ea[1], _ea[0])
	if _be < 0.0 {
		_be += 2 * _c.Pi
	}
	return _be / _c.Pi * 180.0
}
func (_ec *Point) Transform(a, b, c, d, tx, ty float64) {
	_ge := NewMatrix(a, b, c, d, tx, ty)
	_ec.transformByMatrix(_ge)
}
func (_cb Matrix) ScalingFactorY() float64 { return _c.Hypot(_cb[3], _cb[4]) }

type Matrix [9]float64

func (_cd *Matrix) Clone() Matrix              { return NewMatrix(_cd[0], _cd[1], _cd[3], _cd[4], _cd[6], _cd[7]) }
func (_db Matrix) Singular() bool              { return _c.Abs(_db[0]*_db[4]-_db[1]*_db[3]) < _eg }
func (_ad *Point) Set(x, y float64)            { _ad.X, _ad.Y = x, y }
func ScaleMatrix(x, y float64) Matrix          { return NewMatrix(x, 0, 0, y, 0, 0) }
func (_ba Matrix) Rotate(theta float64) Matrix { return _ba.Mult(RotationMatrix(theta)) }
func NewPoint(x, y float64) Point              { return Point{X: x, Y: y} }
func NewMatrixFromTransforms(xScale, yScale, theta, tx, ty float64) Matrix {
	return IdentityMatrix().Scale(xScale, yScale).Rotate(theta).Translate(tx, ty)
}
func (_bbg Point) Distance(b Point) float64        { return _c.Hypot(_bbg.X-b.X, _bbg.Y-b.Y) }
func (_gg Matrix) Translation() (float64, float64) { return _gg[6], _gg[7] }
func (_eaa Matrix) Unrealistic() bool {
	_ab, _gcc, _gdc, _fb := _c.Abs(_eaa[0]), _c.Abs(_eaa[1]), _c.Abs(_eaa[3]), _c.Abs(_eaa[4])
	_bac := _ab > _da && _fb > _da
	_eeg := _gcc > _da && _gdc > _da
	return !(_bac || _eeg)
}
func (_aa Matrix) String() string {
	_ce, _ebde, _ga, _gd, _f, _cc := _aa[0], _aa[1], _aa[3], _aa[4], _aa[6], _aa[7]
	return _a.Sprintf("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d", _ce, _ebde, _ga, _gd, _f, _cc)
}
func (_ag Point) Rotate(theta float64) Point {
	_ced := _c.Hypot(_ag.X, _ag.Y)
	_dg := _c.Atan2(_ag.Y, _ag.X)
	_bgd, _de := _c.Sincos(_dg + theta/180.0*_c.Pi)
	return Point{_ced * _de, _ced * _bgd}
}
func (_dcg Point) String() string {
	return _a.Sprintf("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029", _dcg.X, _dcg.Y)
}
