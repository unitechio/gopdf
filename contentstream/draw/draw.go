package draw

import (
	_c "fmt"
	_ec "math"

	_a "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ee "bitbucket.org/shenghui0779/gopdf/core"
	_d "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ed "bitbucket.org/shenghui0779/gopdf/model"
)

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct {
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	LineColor _ed.PdfColor
	Opacity   float64
	LineWidth float64
	LineStyle LineStyle
	DashArray []int64
	DashPhase int64
}

// NewPoint returns a new point with the coordinates x, y.
func NewPoint(x, y float64) Point { return Point{X: x, Y: y} }

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator(bpath CubicBezierPath, creator *_a.ContentCreator) {
	for _egc, _fgea := range bpath.Curves {
		if _egc == 0 {
			creator.Add_m(_fgea.P0.X, _fgea.P0.Y)
		}
		creator.Add_c(_fgea.P1.X, _fgea.P1.Y, _fgea.P2.X, _fgea.P2.Y, _fgea.P3.X, _fgea.P3.Y)
	}
}

// Draw draws the rectangle. A graphics state can be specified for
// setting additional properties (e.g. opacity). Otherwise pass an empty string
// for the `gsName` parameter. The method returns the content stream as a byte
// array and the bounding box of the shape.
func (_fg Rectangle) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	_bec := _a.NewContentCreator()
	_bec.Add_q()
	if _fg.FillEnabled {
		_bec.SetNonStrokingColor(_fg.FillColor)
	}
	if _fg.BorderEnabled {
		_bec.SetStrokingColor(_fg.BorderColor)
		_bec.Add_w(_fg.BorderWidth)
	}
	if len(gsName) > 1 {
		_bec.Add_gs(_ee.PdfObjectName(gsName))
	}
	var (
		_ce, _ff    = _fg.X, _fg.Y
		_bgc, _gafd = _fg.Width, _fg.Height
		_fb         = _ec.Abs(_fg.BorderRadiusTopLeft)
		_debe       = _ec.Abs(_fg.BorderRadiusTopRight)
		_afe        = _ec.Abs(_fg.BorderRadiusBottomLeft)
		_aae        = _ec.Abs(_fg.BorderRadiusBottomRight)
		_cbd        = 0.4477
	)
	_aea := Path{Points: []Point{{X: _ce + _bgc - _aae, Y: _ff}, {X: _ce + _bgc, Y: _ff + _gafd - _debe}, {X: _ce + _fb, Y: _ff + _gafd}, {X: _ce, Y: _ff + _afe}}}
	_add := [][7]float64{{_aae, _ce + _bgc - _aae*_cbd, _ff, _ce + _bgc, _ff + _aae*_cbd, _ce + _bgc, _ff + _aae}, {_debe, _ce + _bgc, _ff + _gafd - _debe*_cbd, _ce + _bgc - _debe*_cbd, _ff + _gafd, _ce + _bgc - _debe, _ff + _gafd}, {_fb, _ce + _fb*_cbd, _ff + _gafd, _ce, _ff + _gafd - _fb*_cbd, _ce, _ff + _gafd - _fb}, {_afe, _ce, _ff + _afe*_cbd, _ce + _afe*_cbd, _ff, _ce + _afe, _ff}}
	_bec.Add_m(_ce+_afe, _ff)
	for _ddb := 0; _ddb < 4; _ddb++ {
		_fbb := _aea.Points[_ddb]
		_bec.Add_l(_fbb.X, _fbb.Y)
		_abb := _add[_ddb]
		if _ecb := _abb[0]; _ecb != 0 {
			_bec.Add_c(_abb[1], _abb[2], _abb[3], _abb[4], _abb[5], _abb[6])
		}
	}
	_bec.Add_h()
	if _fg.FillEnabled && _fg.BorderEnabled {
		_bec.Add_B()
	} else if _fg.FillEnabled {
		_bec.Add_f()
	} else if _fg.BorderEnabled {
		_bec.Add_S()
	}
	_bec.Add_Q()
	return _bec.Bytes(), _aea.GetBoundingBox().ToPdfRectangle(), nil
}

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween(a Point, b Point) Vector {
	_edd := Vector{}
	_edd.Dx = b.X - a.X
	_edd.Dy = b.Y - a.Y
	return _edd
}

// Point represents a two-dimensional point.
type Point struct {
	X float64
	Y float64
}

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct {
	Points        [][]Point
	FillEnabled   bool
	FillColor     _ed.PdfColor
	BorderEnabled bool
	BorderColor   _ed.PdfColor
	BorderWidth   float64
}

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct {
	Rings         [][]CubicBezierCurve
	FillEnabled   bool
	FillColor     _ed.PdfColor
	BorderEnabled bool
	BorderColor   _ed.PdfColor
	BorderWidth   float64
}

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_cea Line) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	_def, _bgag := _cea.X1, _cea.X2
	_fge, _fae := _cea.Y1, _cea.Y2
	_cba := _fae - _fge
	_dde := _bgag - _def
	_ggf := _ec.Atan2(_cba, _dde)
	L := _ec.Sqrt(_ec.Pow(_dde, 2.0) + _ec.Pow(_cba, 2.0))
	_bad := _cea.LineWidth
	_ace := _ec.Pi
	_cee := 1.0
	if _dde < 0 {
		_cee *= -1.0
	}
	if _cba < 0 {
		_cee *= -1.0
	}
	VsX := _cee * (-_bad / 2 * _ec.Cos(_ggf+_ace/2))
	VsY := _cee * (-_bad/2*_ec.Sin(_ggf+_ace/2) + _bad*_ec.Sin(_ggf+_ace/2))
	V1X := VsX + _bad/2*_ec.Cos(_ggf+_ace/2)
	V1Y := VsY + _bad/2*_ec.Sin(_ggf+_ace/2)
	V2X := VsX + _bad/2*_ec.Cos(_ggf+_ace/2) + L*_ec.Cos(_ggf)
	V2Y := VsY + _bad/2*_ec.Sin(_ggf+_ace/2) + L*_ec.Sin(_ggf)
	V3X := VsX + _bad/2*_ec.Cos(_ggf+_ace/2) + L*_ec.Cos(_ggf) + _bad*_ec.Cos(_ggf-_ace/2)
	V3Y := VsY + _bad/2*_ec.Sin(_ggf+_ace/2) + L*_ec.Sin(_ggf) + _bad*_ec.Sin(_ggf-_ace/2)
	V4X := VsX + _bad/2*_ec.Cos(_ggf-_ace/2)
	V4Y := VsY + _bad/2*_ec.Sin(_ggf-_ace/2)
	_fff := NewPath()
	_fff = _fff.AppendPoint(NewPoint(V1X, V1Y))
	_fff = _fff.AppendPoint(NewPoint(V2X, V2Y))
	_fff = _fff.AppendPoint(NewPoint(V3X, V3Y))
	_fff = _fff.AppendPoint(NewPoint(V4X, V4Y))
	_gba := _cea.LineEndingStyle1
	_eag := _cea.LineEndingStyle2
	_bbe := 3 * _bad
	_eef := 3 * _bad
	_gcc := (_eef - _bad) / 2
	if _eag == LineEndingStyleArrow {
		_eda := _fff.GetPointNumber(2)
		_dfg := NewVectorPolar(_bbe, _ggf+_ace)
		_bgf := _eda.AddVector(_dfg)
		_fga := NewVectorPolar(_eef/2, _ggf+_ace/2)
		_abg := NewVectorPolar(_bbe, _ggf)
		_bde := NewVectorPolar(_gcc, _ggf+_ace/2)
		_cge := _bgf.AddVector(_bde)
		_ceg := _abg.Add(_fga.Flip())
		_ebb := _cge.AddVector(_ceg)
		_gbg := _fga.Scale(2).Flip().Add(_ceg.Flip())
		_gbe := _ebb.AddVector(_gbg)
		_ggd := _bgf.AddVector(NewVectorPolar(_bad, _ggf-_ace/2))
		_fbf := NewPath()
		_fbf = _fbf.AppendPoint(_fff.GetPointNumber(1))
		_fbf = _fbf.AppendPoint(_bgf)
		_fbf = _fbf.AppendPoint(_cge)
		_fbf = _fbf.AppendPoint(_ebb)
		_fbf = _fbf.AppendPoint(_gbe)
		_fbf = _fbf.AppendPoint(_ggd)
		_fbf = _fbf.AppendPoint(_fff.GetPointNumber(4))
		_fff = _fbf
	}
	if _gba == LineEndingStyleArrow {
		_gdf := _fff.GetPointNumber(1)
		_ffa := _fff.GetPointNumber(_fff.Length())
		_eae := NewVectorPolar(_bad/2, _ggf+_ace+_ace/2)
		_bdc := _gdf.AddVector(_eae)
		_eeg := NewVectorPolar(_bbe, _ggf).Add(NewVectorPolar(_eef/2, _ggf+_ace/2))
		_aced := _bdc.AddVector(_eeg)
		_aef := NewVectorPolar(_gcc, _ggf-_ace/2)
		_bcd := _aced.AddVector(_aef)
		_ege := NewVectorPolar(_bbe, _ggf)
		_cbge := _ffa.AddVector(_ege)
		_agf := NewVectorPolar(_gcc, _ggf+_ace+_ace/2)
		_bbd := _cbge.AddVector(_agf)
		_abc := _bdc
		_gfb := NewPath()
		_gfb = _gfb.AppendPoint(_bdc)
		_gfb = _gfb.AppendPoint(_aced)
		_gfb = _gfb.AppendPoint(_bcd)
		for _, _ebfe := range _fff.Points[1 : len(_fff.Points)-1] {
			_gfb = _gfb.AppendPoint(_ebfe)
		}
		_gfb = _gfb.AppendPoint(_cbge)
		_gfb = _gfb.AppendPoint(_bbd)
		_gfb = _gfb.AppendPoint(_abc)
		_fff = _gfb
	}
	_cef := _a.NewContentCreator()
	_cef.Add_q().SetNonStrokingColor(_cea.LineColor)
	if len(gsName) > 1 {
		_cef.Add_gs(_ee.PdfObjectName(gsName))
	}
	_fff = _fff.Offset(_cea.X1, _cea.Y1)
	_fec := _fff.GetBoundingBox()
	DrawPathWithCreator(_fff, _cef)
	if _cea.LineStyle == LineStyleDashed {
		_cef.Add_d([]int64{1, 1}, 0).Add_S().Add_f().Add_Q()
	} else {
		_cef.Add_f().Add_Q()
	}
	return _cef.Bytes(), _fec.ToPdfRectangle(), nil
}

// AppendCurve appends the specified Bezier curve to the path.
func (_gc CubicBezierPath) AppendCurve(curve CubicBezierCurve) CubicBezierPath {
	_gc.Curves = append(_gc.Curves, curve)
	return _gc
}

// Offset shifts the Bezier path with the specified offsets.
func (_gg CubicBezierPath) Offset(offX, offY float64) CubicBezierPath {
	for _ge, _dc := range _gg.Curves {
		_gg.Curves[_ge] = _dc.AddOffsetXY(offX, offY)
	}
	return _gg
}

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_eb Path) RemovePoint(number int) Path {
	if number < 1 || number > len(_eb.Points) {
		return _eb
	}
	_bd := number - 1
	_eb.Points = append(_eb.Points[:_bd], _eb.Points[_bd+1:]...)
	return _eb
}

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar(length float64, theta float64) Vector {
	_ggeb := Vector{}
	_ggeb.Dx = length * _ec.Cos(theta)
	_ggeb.Dy = length * _ec.Sin(theta)
	return _ggeb
}

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct {
	P0 Point
	P1 Point
	P2 Point
	P3 Point
}

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_aad BasicLine) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	_bede := NewPath()
	_bede = _bede.AppendPoint(NewPoint(_aad.X1, _aad.Y1))
	_bede = _bede.AppendPoint(NewPoint(_aad.X2, _aad.Y2))
	_dce := _a.NewContentCreator()
	_dce.Add_q().Add_w(_aad.LineWidth).SetStrokingColor(_aad.LineColor)
	if _aad.LineStyle == LineStyleDashed {
		if _aad.DashArray == nil {
			_aad.DashArray = []int64{1, 1}
		}
		_dce.Add_d(_aad.DashArray, _aad.DashPhase)
	}
	if len(gsName) > 1 {
		_dce.Add_gs(_ee.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_bede, _dce)
	_dce.Add_S().Add_Q()
	return _dce.Bytes(), _bede.GetBoundingBox().ToPdfRectangle(), nil
}

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_dg Point) Add(dx, dy float64) Point { _dg.X += dx; _dg.Y += dy; return _dg }

// GetBoundingBox returns the bounding box of the Bezier path.
func (_b CubicBezierPath) GetBoundingBox() Rectangle {
	_cc := Rectangle{}
	_ccc := 0.0
	_bc := 0.0
	_de := 0.0
	_df := 0.0
	for _ba, _aa := range _b.Curves {
		_fcg := _aa.GetBounds()
		if _ba == 0 {
			_ccc = _fcg.Llx
			_bc = _fcg.Urx
			_de = _fcg.Lly
			_df = _fcg.Ury
			continue
		}
		if _fcg.Llx < _ccc {
			_ccc = _fcg.Llx
		}
		if _fcg.Urx > _bc {
			_bc = _fcg.Urx
		}
		if _fcg.Lly < _de {
			_de = _fcg.Lly
		}
		if _fcg.Ury > _df {
			_df = _fcg.Ury
		}
	}
	_cc.X = _ccc
	_cc.Y = _de
	_cc.Width = _bc - _ccc
	_cc.Height = _df - _de
	return _cc
}

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{ Points []Point }

// GetBounds returns the bounding box of the Bezier curve.
func (_g CubicBezierCurve) GetBounds() _ed.PdfRectangle {
	_gb := _g.P0.X
	_ca := _g.P0.X
	_ag := _g.P0.Y
	_ae := _g.P0.Y
	for _ecc := 0.0; _ecc <= 1.0; _ecc += 0.001 {
		Rx := _g.P0.X*_ec.Pow(1-_ecc, 3) + _g.P1.X*3*_ecc*_ec.Pow(1-_ecc, 2) + _g.P2.X*3*_ec.Pow(_ecc, 2)*(1-_ecc) + _g.P3.X*_ec.Pow(_ecc, 3)
		Ry := _g.P0.Y*_ec.Pow(1-_ecc, 3) + _g.P1.Y*3*_ecc*_ec.Pow(1-_ecc, 2) + _g.P2.Y*3*_ec.Pow(_ecc, 2)*(1-_ecc) + _g.P3.Y*_ec.Pow(_ecc, 3)
		if Rx < _gb {
			_gb = Rx
		}
		if Rx > _ca {
			_ca = Rx
		}
		if Ry < _ag {
			_ag = Ry
		}
		if Ry > _ae {
			_ae = Ry
		}
	}
	_ab := _ed.PdfRectangle{}
	_ab.Llx = _gb
	_ab.Lly = _ag
	_ab.Urx = _ca
	_ab.Ury = _ae
	return _ab
}

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_cbed Vector) GetPolarAngle() float64 { return _ec.Atan2(_cbed.Dy, _cbed.Dx) }

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_fgd Polyline) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	if _fgd.LineColor == nil {
		_fgd.LineColor = _ed.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_aca := NewPath()
	for _, _fcd := range _fgd.Points {
		_aca = _aca.AppendPoint(_fcd)
	}
	_bbb := _a.NewContentCreator()
	_bbb.Add_q().SetStrokingColor(_fgd.LineColor).Add_w(_fgd.LineWidth)
	if len(gsName) > 1 {
		_bbb.Add_gs(_ee.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_aca, _bbb)
	_bbb.Add_S()
	_bbb.Add_Q()
	return _bbb.Bytes(), _aca.GetBoundingBox().ToPdfRectangle(), nil
}

// NewPath returns a new empty path.
func NewPath() Path { return Path{} }

// Vector represents a two-dimensional vector.
type Vector struct {
	Dx float64
	Dy float64
}

// FlipX flips the sign of the Dx component of the vector.
func (_bfc Vector) FlipX() Vector { _bfc.Dx = -_bfc.Dx; return _bfc }

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_bg Path) GetPointNumber(number int) Point {
	if number < 1 || number > len(_bg.Points) {
		return Point{}
	}
	return _bg.Points[number-1]
}

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_deb Circle) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	_abd := _deb.Width / 2
	_cb := _deb.Height / 2
	if _deb.BorderEnabled {
		_abd -= _deb.BorderWidth / 2
		_cb -= _deb.BorderWidth / 2
	}
	_fca := 0.551784
	_ccf := _abd * _fca
	_afb := _cb * _fca
	_gd := NewCubicBezierPath()
	_gd = _gd.AppendCurve(NewCubicBezierCurve(-_abd, 0, -_abd, _afb, -_ccf, _cb, 0, _cb))
	_gd = _gd.AppendCurve(NewCubicBezierCurve(0, _cb, _ccf, _cb, _abd, _afb, _abd, 0))
	_gd = _gd.AppendCurve(NewCubicBezierCurve(_abd, 0, _abd, -_afb, _ccf, -_cb, 0, -_cb))
	_gd = _gd.AppendCurve(NewCubicBezierCurve(0, -_cb, -_ccf, -_cb, -_abd, -_afb, -_abd, 0))
	_gd = _gd.Offset(_abd, _cb)
	if _deb.BorderEnabled {
		_gd = _gd.Offset(_deb.BorderWidth/2, _deb.BorderWidth/2)
	}
	if _deb.X != 0 || _deb.Y != 0 {
		_gd = _gd.Offset(_deb.X, _deb.Y)
	}
	_bga := _a.NewContentCreator()
	_bga.Add_q()
	if _deb.FillEnabled {
		_bga.SetNonStrokingColor(_deb.FillColor)
	}
	if _deb.BorderEnabled {
		_bga.SetStrokingColor(_deb.BorderColor)
		_bga.Add_w(_deb.BorderWidth)
	}
	if len(gsName) > 1 {
		_bga.Add_gs(_ee.PdfObjectName(gsName))
	}
	DrawBezierPathWithCreator(_gd, _bga)
	_bga.Add_h()
	if _deb.FillEnabled && _deb.BorderEnabled {
		_bga.Add_B()
	} else if _deb.FillEnabled {
		_bga.Add_f()
	} else if _deb.BorderEnabled {
		_bga.Add_S()
	}
	_bga.Add_Q()
	_gf := _gd.GetBoundingBox()
	if _deb.BorderEnabled {
		_gf.Height += _deb.BorderWidth
		_gf.Width += _deb.BorderWidth
		_gf.X -= _deb.BorderWidth / 2
		_gf.Y -= _deb.BorderWidth / 2
	}
	return _bga.Bytes(), _gf.ToPdfRectangle(), nil
}
func (_ggc Point) String() string {
	return _c.Sprintf("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029", _ggc.X, _ggc.Y)
}

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector(dx, dy float64) Vector { _fdc := Vector{}; _fdc.Dx = dx; _fdc.Dy = dy; return _fdc }

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct {
	Curves      []CubicBezierCurve
	BorderWidth float64
	BorderColor _ed.PdfColor
	FillEnabled bool
	FillColor   _ed.PdfColor
}

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     _ed.PdfColor
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   _ed.PdfColor
	Opacity       float64
}

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_aee Point) Rotate(theta float64) Point {
	_aaf := _d.NewPoint(_aee.X, _aee.Y).Rotate(theta)
	return NewPoint(_aaf.X, _aaf.Y)
}

const (
	LineStyleSolid  LineStyle = 0
	LineStyleDashed LineStyle = 1
)

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath() CubicBezierPath {
	_ea := CubicBezierPath{}
	_ea.Curves = []CubicBezierCurve{}
	return _ea
}

const (
	LineEndingStyleNone  LineEndingStyle = 0
	LineEndingStyleArrow LineEndingStyle = 1
	LineEndingStyleButt  LineEndingStyle = 2
)

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_f CubicBezierCurve) AddOffsetXY(offX, offY float64) CubicBezierCurve {
	_f.P0.X += offX
	_f.P1.X += offX
	_f.P2.X += offX
	_f.P3.X += offX
	_f.P0.Y += offY
	_f.P1.Y += offY
	_f.P2.Y += offY
	_f.P3.Y += offY
	return _f
}

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Copy returns a clone of the Bezier path.
func (_af CubicBezierPath) Copy() CubicBezierPath {
	_fc := CubicBezierPath{}
	_fc.Curves = append(_fc.Curves, _af.Curves...)
	return _fc
}

// LineStyle refers to how the line will be created.
type LineStyle int

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve(x0, y0, x1, y1, x2, y2, x3, y3 float64) CubicBezierCurve {
	_ac := CubicBezierCurve{}
	_ac.P0 = NewPoint(x0, y0)
	_ac.P1 = NewPoint(x1, y1)
	_ac.P2 = NewPoint(x2, y2)
	_ac.P3 = NewPoint(x3, y3)
	return _ac
}

// Flip changes the sign of the vector: -vector.
func (_beg Vector) Flip() Vector {
	_dbb := _beg.Magnitude()
	_fgb := _beg.GetPolarAngle()
	_beg.Dx = _dbb * _ec.Cos(_fgb+_ec.Pi)
	_beg.Dy = _dbb * _ec.Sin(_fgb+_ec.Pi)
	return _beg
}

// Offset shifts the path with the specified offsets.
func (_dea Path) Offset(offX, offY float64) Path {
	for _cce, _be := range _dea.Points {
		_dea.Points[_cce] = _be.Add(offX, offY)
	}
	return _dea
}

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        _ed.PdfColor
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 LineEndingStyle
	LineEndingStyle2 LineEndingStyle
	LineStyle        LineStyle
}

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_cbf Rectangle) ToPdfRectangle() *_ed.PdfRectangle {
	return &_ed.PdfRectangle{Llx: _cbf.X, Lly: _cbf.Y, Urx: _cbf.X + _cbf.Width, Ury: _cbf.Y + _cbf.Height}
}

// AddVector adds vector to a point.
func (_ega Point) AddVector(v Vector) Point { _ega.X += v.Dx; _ega.Y += v.Dy; return _ega }

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_dcc CurvePolygon) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	_fac := _a.NewContentCreator()
	_fac.Add_q()
	_dcc.FillEnabled = _dcc.FillEnabled && _dcc.FillColor != nil
	if _dcc.FillEnabled {
		_fac.SetNonStrokingColor(_dcc.FillColor)
	}
	_dcc.BorderEnabled = _dcc.BorderEnabled && _dcc.BorderColor != nil
	if _dcc.BorderEnabled {
		_fac.SetStrokingColor(_dcc.BorderColor)
		_fac.Add_w(_dcc.BorderWidth)
	}
	if len(gsName) > 1 {
		_fac.Add_gs(_ee.PdfObjectName(gsName))
	}
	_ebf := NewCubicBezierPath()
	for _, _ef := range _dcc.Rings {
		for _fcgc, _bgb := range _ef {
			if _fcgc == 0 {
				_fac.Add_m(_bgb.P0.X, _bgb.P0.Y)
			} else {
				_fac.Add_l(_bgb.P0.X, _bgb.P0.Y)
			}
			_fac.Add_c(_bgb.P1.X, _bgb.P1.Y, _bgb.P2.X, _bgb.P2.Y, _bgb.P3.X, _bgb.P3.Y)
			_ebf = _ebf.AppendCurve(_bgb)
		}
		_fac.Add_h()
	}
	if _dcc.FillEnabled && _dcc.BorderEnabled {
		_fac.Add_B()
	} else if _dcc.FillEnabled {
		_fac.Add_f()
	} else if _dcc.BorderEnabled {
		_fac.Add_S()
	}
	_fac.Add_Q()
	return _fac.Bytes(), _ebf.GetBoundingBox().ToPdfRectangle(), nil
}

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct {

	// Position and size properties.
	X      float64
	Y      float64
	Width  float64
	Height float64

	// Fill properties.
	FillEnabled bool
	FillColor   _ed.PdfColor

	// Border properties.
	BorderEnabled           bool
	BorderColor             _ed.PdfColor
	BorderWidth             float64
	BorderRadiusTopLeft     float64
	BorderRadiusTopRight    float64
	BorderRadiusBottomLeft  float64
	BorderRadiusBottomRight float64

	// Shape opacity (0-1 interval).
	Opacity float64
}

// Scale scales the vector by the specified factor.
func (_aab Vector) Scale(factor float64) Vector {
	_feb := _aab.Magnitude()
	_gff := _aab.GetPolarAngle()
	_aab.Dx = factor * _feb * _ec.Cos(_gff)
	_aab.Dy = factor * _feb * _ec.Sin(_gff)
	return _aab
}

// Copy returns a clone of the path.
func (_gge Path) Copy() Path {
	_dd := Path{}
	_dd.Points = append(_dd.Points, _gge.Points...)
	return _dd
}

// AppendPoint adds the specified point to the path.
func (_fa Path) AppendPoint(point Point) Path { _fa.Points = append(_fa.Points, point); return _fa }

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_cgg Polygon) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	_eab := _a.NewContentCreator()
	_eab.Add_q()
	_cgg.FillEnabled = _cgg.FillEnabled && _cgg.FillColor != nil
	if _cgg.FillEnabled {
		_eab.SetNonStrokingColor(_cgg.FillColor)
	}
	_cgg.BorderEnabled = _cgg.BorderEnabled && _cgg.BorderColor != nil
	if _cgg.BorderEnabled {
		_eab.SetStrokingColor(_cgg.BorderColor)
		_eab.Add_w(_cgg.BorderWidth)
	}
	if len(gsName) > 1 {
		_eab.Add_gs(_ee.PdfObjectName(gsName))
	}
	_cac := NewPath()
	for _, _cbg := range _cgg.Points {
		for _gdg, _cbe := range _cbg {
			_cac = _cac.AppendPoint(_cbe)
			if _gdg == 0 {
				_eab.Add_m(_cbe.X, _cbe.Y)
			} else {
				_eab.Add_l(_cbe.X, _cbe.Y)
			}
		}
		_eab.Add_h()
	}
	if _cgg.FillEnabled && _cgg.BorderEnabled {
		_eab.Add_B()
	} else if _cgg.FillEnabled {
		_eab.Add_f()
	} else if _cgg.BorderEnabled {
		_eab.Add_S()
	}
	_eab.Add_Q()
	return _eab.Bytes(), _cac.GetBoundingBox().ToPdfRectangle(), nil
}

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_eg BoundingBox) ToPdfRectangle() *_ed.PdfRectangle {
	return &_ed.PdfRectangle{Llx: _eg.X, Lly: _eg.Y, Urx: _eg.X + _eg.Width, Ury: _eg.Y + _eg.Height}
}

// FlipY flips the sign of the Dy component of the vector.
func (_da Vector) FlipY() Vector { _da.Dy = -_da.Dy; return _da }

// Length returns the number of points in the path.
func (_cg Path) Length() int { return len(_cg.Points) }

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_fd PolyBezierCurve) Draw(gsName string) ([]byte, *_ed.PdfRectangle, error) {
	if _fd.BorderColor == nil {
		_fd.BorderColor = _ed.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_db := NewCubicBezierPath()
	for _, _baf := range _fd.Curves {
		_db = _db.AppendCurve(_baf)
	}
	_ad := _a.NewContentCreator()
	_ad.Add_q()
	_fd.FillEnabled = _fd.FillEnabled && _fd.FillColor != nil
	if _fd.FillEnabled {
		_ad.SetNonStrokingColor(_fd.FillColor)
	}
	_ad.SetStrokingColor(_fd.BorderColor)
	_ad.Add_w(_fd.BorderWidth)
	if len(gsName) > 1 {
		_ad.Add_gs(_ee.PdfObjectName(gsName))
	}
	for _acf, _bed := range _db.Curves {
		if _acf == 0 {
			_ad.Add_m(_bed.P0.X, _bed.P0.Y)
		} else {
			_ad.Add_l(_bed.P0.X, _bed.P0.Y)
		}
		_ad.Add_c(_bed.P1.X, _bed.P1.Y, _bed.P2.X, _bed.P2.Y, _bed.P3.X, _bed.P3.Y)
	}
	if _fd.FillEnabled {
		_ad.Add_h()
		_ad.Add_B()
	} else {
		_ad.Add_S()
	}
	_ad.Add_Q()
	return _ad.Bytes(), _db.GetBoundingBox().ToPdfRectangle(), nil
}

// Add adds the specified vector to the current one and returns the result.
func (_fad Vector) Add(other Vector) Vector { _fad.Dx += other.Dx; _fad.Dy += other.Dy; return _fad }

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct {
	Points    []Point
	LineColor _ed.PdfColor
	LineWidth float64
}

// GetBoundingBox returns the bounding box of the path.
func (_bf Path) GetBoundingBox() BoundingBox {
	_aba := BoundingBox{}
	_dcb := 0.0
	_bee := 0.0
	_fe := 0.0
	_gbc := 0.0
	for _bb, _ga := range _bf.Points {
		if _bb == 0 {
			_dcb = _ga.X
			_bee = _ga.X
			_fe = _ga.Y
			_gbc = _ga.Y
			continue
		}
		if _ga.X < _dcb {
			_dcb = _ga.X
		}
		if _ga.X > _bee {
			_bee = _ga.X
		}
		if _ga.Y < _fe {
			_fe = _ga.Y
		}
		if _ga.Y > _gbc {
			_gbc = _ga.Y
		}
	}
	_aba.X = _dcb
	_aba.Y = _fe
	_aba.Width = _bee - _dcb
	_aba.Height = _gbc - _fe
	return _aba
}

// Magnitude returns the magnitude of the vector.
func (_eec Vector) Magnitude() float64 {
	return _ec.Sqrt(_ec.Pow(_eec.Dx, 2.0) + _ec.Pow(_eec.Dy, 2.0))
}

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator(path Path, creator *_a.ContentCreator) {
	for _eaeb, _edae := range path.Points {
		if _eaeb == 0 {
			creator.Add_m(_edae.X, _edae.Y)
		} else {
			creator.Add_l(_edae.X, _edae.Y)
		}
	}
}

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{ Curves []CubicBezierCurve }

// Rotate rotates the vector by the specified angle.
func (_bab Vector) Rotate(phi float64) Vector {
	_ggcc := _bab.Magnitude()
	_bda := _bab.GetPolarAngle()
	return NewVectorPolar(_ggcc, _bda+phi)
}
