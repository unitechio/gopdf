package draw

import (
	_ff "fmt"
	_d "math"

	_fc "bitbucket.org/shenghui0779/gopdf/contentstream"
	_eg "bitbucket.org/shenghui0779/gopdf/core"
	_egg "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_e "bitbucket.org/shenghui0779/gopdf/model"
)

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath() CubicBezierPath {
	_bg := CubicBezierPath{}
	_bg.Curves = []CubicBezierCurve{}
	return _bg
}
func (_dd Point) String() string {
	return _ff.Sprintf("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029", _dd.X, _dd.Y)
}

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_aga Rectangle) ToPdfRectangle() *_e.PdfRectangle {
	return &_e.PdfRectangle{Llx: _aga.X, Lly: _aga.Y, Urx: _aga.X + _aga.Width, Ury: _aga.Y + _aga.Height}
}

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_dge PolyBezierCurve) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	if _dge.BorderColor == nil {
		_dge.BorderColor = _e.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_fdf := NewCubicBezierPath()
	for _, _fcf := range _dge.Curves {
		_fdf = _fdf.AppendCurve(_fcf)
	}
	_aaaga := _fc.NewContentCreator()
	_aaaga.Add_q()
	_dge.FillEnabled = _dge.FillEnabled && _dge.FillColor != nil
	if _dge.FillEnabled {
		_aaaga.SetNonStrokingColor(_dge.FillColor)
	}
	_aaaga.SetStrokingColor(_dge.BorderColor)
	_aaaga.Add_w(_dge.BorderWidth)
	if len(gsName) > 1 {
		_aaaga.Add_gs(_eg.PdfObjectName(gsName))
	}
	for _gc, _ccef := range _fdf.Curves {
		if _gc == 0 {
			_aaaga.Add_m(_ccef.P0.X, _ccef.P0.Y)
		} else {
			_aaaga.Add_l(_ccef.P0.X, _ccef.P0.Y)
		}
		_aaaga.Add_c(_ccef.P1.X, _ccef.P1.Y, _ccef.P2.X, _ccef.P2.Y, _ccef.P3.X, _ccef.P3.Y)
	}
	if _dge.FillEnabled {
		_aaaga.Add_h()
		_aaaga.Add_B()
	} else {
		_aaaga.Add_S()
	}
	_aaaga.Add_Q()
	return _aaaga.Bytes(), _fdf.GetBoundingBox().ToPdfRectangle(), nil
}

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{ Points []Point }

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_agc BasicLine) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	_db := NewPath()
	_db = _db.AppendPoint(NewPoint(_agc.X1, _agc.Y1))
	_db = _db.AppendPoint(NewPoint(_agc.X2, _agc.Y2))
	_baac := _fc.NewContentCreator()
	_baac.Add_q().Add_w(_agc.LineWidth).SetStrokingColor(_agc.LineColor)
	if _agc.LineStyle == LineStyleDashed {
		if _agc.DashArray == nil {
			_agc.DashArray = []int64{1, 1}
		}
		_baac.Add_d(_agc.DashArray, _agc.DashPhase)
	}
	if len(gsName) > 1 {
		_baac.Add_gs(_eg.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_db, _baac)
	_baac.Add_S().Add_Q()
	return _baac.Bytes(), _db.GetBoundingBox().ToPdfRectangle(), nil
}

const (
	LineStyleSolid  LineStyle = 0
	LineStyleDashed LineStyle = 1
)

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve(x0, y0, x1, y1, x2, y2, x3, y3 float64) CubicBezierCurve {
	_a := CubicBezierCurve{}
	_a.P0 = NewPoint(x0, y0)
	_a.P1 = NewPoint(x1, y1)
	_a.P2 = NewPoint(x2, y2)
	_a.P3 = NewPoint(x3, y3)
	return _a
}

// Offset shifts the path with the specified offsets.
func (_cg Path) Offset(offX, offY float64) Path {
	for _cb, _fec := range _cg.Points {
		_cg.Points[_cb] = _fec.Add(offX, offY)
	}
	return _cg
}

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_b CubicBezierCurve) AddOffsetXY(offX, offY float64) CubicBezierCurve {
	_b.P0.X += offX
	_b.P1.X += offX
	_b.P2.X += offX
	_b.P3.X += offX
	_b.P0.Y += offY
	_b.P1.Y += offY
	_b.P2.Y += offY
	_b.P3.Y += offY
	return _b
}

// FlipY flips the sign of the Dy component of the vector.
func (_cgf Vector) FlipY() Vector { _cgf.Dy = -_cgf.Dy; return _cgf }

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_ce CurvePolygon) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	_efb := _fc.NewContentCreator()
	_efb.Add_q()
	_ce.FillEnabled = _ce.FillEnabled && _ce.FillColor != nil
	if _ce.FillEnabled {
		_efb.SetNonStrokingColor(_ce.FillColor)
	}
	_ce.BorderEnabled = _ce.BorderEnabled && _ce.BorderColor != nil
	if _ce.BorderEnabled {
		_efb.SetStrokingColor(_ce.BorderColor)
		_efb.Add_w(_ce.BorderWidth)
	}
	if len(gsName) > 1 {
		_efb.Add_gs(_eg.PdfObjectName(gsName))
	}
	_agd := NewCubicBezierPath()
	for _, _gdb := range _ce.Rings {
		for _afbd, _ccc := range _gdb {
			if _afbd == 0 {
				_efb.Add_m(_ccc.P0.X, _ccc.P0.Y)
			} else {
				_efb.Add_l(_ccc.P0.X, _ccc.P0.Y)
			}
			_efb.Add_c(_ccc.P1.X, _ccc.P1.Y, _ccc.P2.X, _ccc.P2.Y, _ccc.P3.X, _ccc.P3.Y)
			_agd = _agd.AppendCurve(_ccc)
		}
		_efb.Add_h()
	}
	if _ce.FillEnabled && _ce.BorderEnabled {
		_efb.Add_B()
	} else if _ce.FillEnabled {
		_efb.Add_f()
	} else if _ce.BorderEnabled {
		_efb.Add_S()
	}
	_efb.Add_Q()
	return _efb.Bytes(), _agd.GetBoundingBox().ToPdfRectangle(), nil
}

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_afe Point) Add(dx, dy float64) Point { _afe.X += dx; _afe.Y += dy; return _afe }

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar(length float64, theta float64) Vector {
	_fdd := Vector{}
	_fdd.Dx = length * _d.Cos(theta)
	_fdd.Dy = length * _d.Sin(theta)
	return _fdd
}

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct {
	Points    []Point
	LineColor _e.PdfColor
	LineWidth float64
}

// FlipX flips the sign of the Dx component of the vector.
func (_bcae Vector) FlipX() Vector { _bcae.Dx = -_bcae.Dx; return _bcae }

// Draw draws the rectangle. A graphics state can be specified for
// setting additional properties (e.g. opacity). Otherwise pass an empty string
// for the `gsName` parameter. The method returns the content stream as a byte
// array and the bounding box of the shape.
func (_bag Rectangle) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	_afba := _fc.NewContentCreator()
	_afba.Add_q()
	if _bag.FillEnabled {
		_afba.SetNonStrokingColor(_bag.FillColor)
	}
	if _bag.BorderEnabled {
		_afba.SetStrokingColor(_bag.BorderColor)
		_afba.Add_w(_bag.BorderWidth)
	}
	if len(gsName) > 1 {
		_afba.Add_gs(_eg.PdfObjectName(gsName))
	}
	var (
		_ccca, _daee = _bag.X, _bag.Y
		_ddf, _bec   = _bag.Width, _bag.Height
		_cdg         = _d.Abs(_bag.BorderRadiusTopLeft)
		_eggd        = _d.Abs(_bag.BorderRadiusTopRight)
		_afg         = _d.Abs(_bag.BorderRadiusBottomLeft)
		_bf          = _d.Abs(_bag.BorderRadiusBottomRight)
		_daf         = 0.4477
	)
	_deg := Path{Points: []Point{{X: _ccca + _ddf - _bf, Y: _daee}, {X: _ccca + _ddf, Y: _daee + _bec - _eggd}, {X: _ccca + _cdg, Y: _daee + _bec}, {X: _ccca, Y: _daee + _afg}}}
	_gf := [][7]float64{{_bf, _ccca + _ddf - _bf*_daf, _daee, _ccca + _ddf, _daee + _bf*_daf, _ccca + _ddf, _daee + _bf}, {_eggd, _ccca + _ddf, _daee + _bec - _eggd*_daf, _ccca + _ddf - _eggd*_daf, _daee + _bec, _ccca + _ddf - _eggd, _daee + _bec}, {_cdg, _ccca + _cdg*_daf, _daee + _bec, _ccca, _daee + _bec - _cdg*_daf, _ccca, _daee + _bec - _cdg}, {_afg, _ccca, _daee + _afg*_daf, _ccca + _afg*_daf, _daee, _ccca + _afg, _daee}}
	_afba.Add_m(_ccca+_afg, _daee)
	for _fcaa := 0; _fcaa < 4; _fcaa++ {
		_cfg := _deg.Points[_fcaa]
		_afba.Add_l(_cfg.X, _cfg.Y)
		_cfa := _gf[_fcaa]
		if _bedc := _cfa[0]; _bedc != 0 {
			_afba.Add_c(_cfa[1], _cfa[2], _cfa[3], _cfa[4], _cfa[5], _cfa[6])
		}
	}
	_afba.Add_h()
	if _bag.FillEnabled && _bag.BorderEnabled {
		_afba.Add_B()
	} else if _bag.FillEnabled {
		_afba.Add_f()
	} else if _bag.BorderEnabled {
		_afba.Add_S()
	}
	_afba.Add_Q()
	return _afba.Bytes(), _deg.GetBoundingBox().ToPdfRectangle(), nil
}

// Scale scales the vector by the specified factor.
func (_efba Vector) Scale(factor float64) Vector {
	_gaa := _efba.Magnitude()
	_eeg := _efba.GetPolarAngle()
	_efba.Dx = factor * _gaa * _d.Cos(_eeg)
	_efba.Dy = factor * _gaa * _d.Sin(_eeg)
	return _efba
}

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_ddg Circle) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	_cf := _ddg.Width / 2
	_fcd := _ddg.Height / 2
	if _ddg.BorderEnabled {
		_cf -= _ddg.BorderWidth / 2
		_fcd -= _ddg.BorderWidth / 2
	}
	_cd := 0.551784
	_dae := _cf * _cd
	_gg := _fcd * _cd
	_dgb := NewCubicBezierPath()
	_dgb = _dgb.AppendCurve(NewCubicBezierCurve(-_cf, 0, -_cf, _gg, -_dae, _fcd, 0, _fcd))
	_dgb = _dgb.AppendCurve(NewCubicBezierCurve(0, _fcd, _dae, _fcd, _cf, _gg, _cf, 0))
	_dgb = _dgb.AppendCurve(NewCubicBezierCurve(_cf, 0, _cf, -_gg, _dae, -_fcd, 0, -_fcd))
	_dgb = _dgb.AppendCurve(NewCubicBezierCurve(0, -_fcd, -_dae, -_fcd, -_cf, -_gg, -_cf, 0))
	_dgb = _dgb.Offset(_cf, _fcd)
	if _ddg.BorderEnabled {
		_dgb = _dgb.Offset(_ddg.BorderWidth/2, _ddg.BorderWidth/2)
	}
	if _ddg.X != 0 || _ddg.Y != 0 {
		_dgb = _dgb.Offset(_ddg.X, _ddg.Y)
	}
	_ffa := _fc.NewContentCreator()
	_ffa.Add_q()
	if _ddg.FillEnabled {
		_ffa.SetNonStrokingColor(_ddg.FillColor)
	}
	if _ddg.BorderEnabled {
		_ffa.SetStrokingColor(_ddg.BorderColor)
		_ffa.Add_w(_ddg.BorderWidth)
	}
	if len(gsName) > 1 {
		_ffa.Add_gs(_eg.PdfObjectName(gsName))
	}
	DrawBezierPathWithCreator(_dgb, _ffa)
	_ffa.Add_h()
	if _ddg.FillEnabled && _ddg.BorderEnabled {
		_ffa.Add_B()
	} else if _ddg.FillEnabled {
		_ffa.Add_f()
	} else if _ddg.BorderEnabled {
		_ffa.Add_S()
	}
	_ffa.Add_Q()
	_ae := _dgb.GetBoundingBox()
	if _ddg.BorderEnabled {
		_ae.Height += _ddg.BorderWidth
		_ae.Width += _ddg.BorderWidth
		_ae.X -= _ddg.BorderWidth / 2
		_ae.Y -= _ddg.BorderWidth / 2
	}
	return _ffa.Bytes(), _ae.ToPdfRectangle(), nil
}

// Vector represents a two-dimensional vector.
type Vector struct {
	Dx float64
	Dy float64
}

// NewPoint returns a new point with the coordinates x, y.
func NewPoint(x, y float64) Point { return Point{X: x, Y: y} }

// LineStyle refers to how the line will be created.
type LineStyle int

// AppendPoint adds the specified point to the path.
func (_cce Path) AppendPoint(point Point) Path { _cce.Points = append(_cce.Points, point); return _cce }

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_bgf Vector) GetPolarAngle() float64 { return _d.Atan2(_bgf.Dy, _bgf.Dx) }

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator(path Path, creator *_fc.ContentCreator) {
	for _dddf, _cbf := range path.Points {
		if _dddf == 0 {
			creator.Add_m(_cbf.X, _cbf.Y)
		} else {
			creator.Add_l(_cbf.X, _cbf.Y)
		}
	}
}

// Copy returns a clone of the path.
func (_fce Path) Copy() Path {
	_bc := Path{}
	_bc.Points = append(_bc.Points, _fce.Points...)
	return _bc
}

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{ Curves []CubicBezierCurve }

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_ged Path) RemovePoint(number int) Path {
	if number < 1 || number > len(_ged.Points) {
		return _ged
	}
	_fg := number - 1
	_ged.Points = append(_ged.Points[:_fg], _ged.Points[_fg+1:]...)
	return _ged
}

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct {
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	LineColor _e.PdfColor
	Opacity   float64
	LineWidth float64
	LineStyle LineStyle
	DashArray []int64
	DashPhase int64
}

// AddVector adds vector to a point.
func (_fd Point) AddVector(v Vector) Point { _fd.X += v.Dx; _fd.Y += v.Dy; return _fd }

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween(a Point, b Point) Vector {
	_ceg := Vector{}
	_ceg.Dx = b.X - a.X
	_ceg.Dy = b.Y - a.Y
	return _ceg
}

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct {
	Curves      []CubicBezierCurve
	BorderWidth float64
	BorderColor _e.PdfColor
	FillEnabled bool
	FillColor   _e.PdfColor
}

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector(dx, dy float64) Vector { _ffdf := Vector{}; _ffdf.Dx = dx; _ffdf.Dy = dy; return _ffdf }

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator(bpath CubicBezierPath, creator *_fc.ContentCreator) {
	for _fced, _dgf := range bpath.Curves {
		if _fced == 0 {
			creator.Add_m(_dgf.P0.X, _dgf.P0.Y)
		}
		creator.Add_c(_dgf.P1.X, _dgf.P1.Y, _dgf.P2.X, _dgf.P2.Y, _dgf.P3.X, _dgf.P3.Y)
	}
}

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        _e.PdfColor
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 LineEndingStyle
	LineEndingStyle2 LineEndingStyle
	LineStyle        LineStyle
}

// Offset shifts the Bezier path with the specified offsets.
func (_ef CubicBezierPath) Offset(offX, offY float64) CubicBezierPath {
	for _feg, _cc := range _ef.Curves {
		_ef.Curves[_feg] = _cc.AddOffsetXY(offX, offY)
	}
	return _ef
}

// Point represents a two-dimensional point.
type Point struct {
	X float64
	Y float64
}

// AppendCurve appends the specified Bezier curve to the path.
func (_aaa CubicBezierPath) AppendCurve(curve CubicBezierCurve) CubicBezierPath {
	_aaa.Curves = append(_aaa.Curves, curve)
	return _aaa
}

const (
	LineEndingStyleNone  LineEndingStyle = 0
	LineEndingStyleArrow LineEndingStyle = 1
	LineEndingStyleButt  LineEndingStyle = 2
)

// Flip changes the sign of the vector: -vector.
func (_fbe Vector) Flip() Vector {
	_ffe := _fbe.Magnitude()
	_afbb := _fbe.GetPolarAngle()
	_fbe.Dx = _ffe * _d.Cos(_afbb+_d.Pi)
	_fbe.Dy = _ffe * _d.Sin(_afbb+_d.Pi)
	return _fbe
}

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     _e.PdfColor
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   _e.PdfColor
	Opacity       float64
}

// Add adds the specified vector to the current one and returns the result.
func (_aae Vector) Add(other Vector) Vector { _aae.Dx += other.Dx; _aae.Dy += other.Dy; return _aae }

// GetBoundingBox returns the bounding box of the path.
func (_ed Path) GetBoundingBox() BoundingBox {
	_dab := BoundingBox{}
	_fgf := 0.0
	_abf := 0.0
	_aaag := 0.0
	_bca := 0.0
	for _ead, _af := range _ed.Points {
		if _ead == 0 {
			_fgf = _af.X
			_abf = _af.X
			_aaag = _af.Y
			_bca = _af.Y
			continue
		}
		if _af.X < _fgf {
			_fgf = _af.X
		}
		if _af.X > _abf {
			_abf = _af.X
		}
		if _af.Y < _aaag {
			_aaag = _af.Y
		}
		if _af.Y > _bca {
			_bca = _af.Y
		}
	}
	_dab.X = _fgf
	_dab.Y = _aaag
	_dab.Width = _abf - _fgf
	_dab.Height = _bca - _aaag
	return _dab
}

// Magnitude returns the magnitude of the vector.
func (_gegc Vector) Magnitude() float64 {
	return _d.Sqrt(_d.Pow(_gegc.Dx, 2.0) + _d.Pow(_gegc.Dy, 2.0))
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
	FillColor   _e.PdfColor

	// Border properties.
	BorderEnabled           bool
	BorderColor             _e.PdfColor
	BorderWidth             float64
	BorderRadiusTopLeft     float64
	BorderRadiusTopRight    float64
	BorderRadiusBottomLeft  float64
	BorderRadiusBottomRight float64

	// Shape opacity (0-1 interval).
	Opacity float64
}

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct {
	Rings         [][]CubicBezierCurve
	FillEnabled   bool
	FillColor     _e.PdfColor
	BorderEnabled bool
	BorderColor   _e.PdfColor
	BorderWidth   float64
}

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_ac Path) GetPointNumber(number int) Point {
	if number < 1 || number > len(_ac.Points) {
		return Point{}
	}
	return _ac.Points[number-1]
}

// Rotate rotates the vector by the specified angle.
func (_fbc Vector) Rotate(phi float64) Vector {
	_bbb := _fbc.Magnitude()
	_egca := _fbc.GetPolarAngle()
	return NewVectorPolar(_bbb, _egca+phi)
}

// GetBoundingBox returns the bounding box of the Bezier path.
func (_da CubicBezierPath) GetBoundingBox() Rectangle {
	_gb := Rectangle{}
	_dg := 0.0
	_efa := 0.0
	_ea := 0.0
	_fb := 0.0
	for _ge, _gd := range _da.Curves {
		_ec := _gd.GetBounds()
		if _ge == 0 {
			_dg = _ec.Llx
			_efa = _ec.Urx
			_ea = _ec.Lly
			_fb = _ec.Ury
			continue
		}
		if _ec.Llx < _dg {
			_dg = _ec.Llx
		}
		if _ec.Urx > _efa {
			_efa = _ec.Urx
		}
		if _ec.Lly < _ea {
			_ea = _ec.Lly
		}
		if _ec.Ury > _fb {
			_fb = _ec.Ury
		}
	}
	_gb.X = _dg
	_gb.Y = _ea
	_gb.Width = _efa - _dg
	_gb.Height = _fb - _ea
	return _gb
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

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_dcb Line) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	_df, _dce := _dcb.X1, _dcb.X2
	_gce, _aef := _dcb.Y1, _dcb.Y2
	_egc := _aef - _gce
	_gfc := _dce - _df
	_cgb := _d.Atan2(_egc, _gfc)
	L := _d.Sqrt(_d.Pow(_gfc, 2.0) + _d.Pow(_egc, 2.0))
	_ddd := _dcb.LineWidth
	_geg := _d.Pi
	_ddc := 1.0
	if _gfc < 0 {
		_ddc *= -1.0
	}
	if _egc < 0 {
		_ddc *= -1.0
	}
	VsX := _ddc * (-_ddd / 2 * _d.Cos(_cgb+_geg/2))
	VsY := _ddc * (-_ddd/2*_d.Sin(_cgb+_geg/2) + _ddd*_d.Sin(_cgb+_geg/2))
	V1X := VsX + _ddd/2*_d.Cos(_cgb+_geg/2)
	V1Y := VsY + _ddd/2*_d.Sin(_cgb+_geg/2)
	V2X := VsX + _ddd/2*_d.Cos(_cgb+_geg/2) + L*_d.Cos(_cgb)
	V2Y := VsY + _ddd/2*_d.Sin(_cgb+_geg/2) + L*_d.Sin(_cgb)
	V3X := VsX + _ddd/2*_d.Cos(_cgb+_geg/2) + L*_d.Cos(_cgb) + _ddd*_d.Cos(_cgb-_geg/2)
	V3Y := VsY + _ddd/2*_d.Sin(_cgb+_geg/2) + L*_d.Sin(_cgb) + _ddd*_d.Sin(_cgb-_geg/2)
	V4X := VsX + _ddd/2*_d.Cos(_cgb-_geg/2)
	V4Y := VsY + _ddd/2*_d.Sin(_cgb-_geg/2)
	_bcf := NewPath()
	_bcf = _bcf.AppendPoint(NewPoint(V1X, V1Y))
	_bcf = _bcf.AppendPoint(NewPoint(V2X, V2Y))
	_bcf = _bcf.AppendPoint(NewPoint(V3X, V3Y))
	_bcf = _bcf.AppendPoint(NewPoint(V4X, V4Y))
	_afgg := _dcb.LineEndingStyle1
	_gec := _dcb.LineEndingStyle2
	_acf := 3 * _ddd
	_afa := 3 * _ddd
	_edc := (_afa - _ddd) / 2
	if _gec == LineEndingStyleArrow {
		_acb := _bcf.GetPointNumber(2)
		_gbd := NewVectorPolar(_acf, _cgb+_geg)
		_abd := _acb.AddVector(_gbd)
		_fbg := NewVectorPolar(_afa/2, _cgb+_geg/2)
		_cec := NewVectorPolar(_acf, _cgb)
		_gdg := NewVectorPolar(_edc, _cgb+_geg/2)
		_geef := _abd.AddVector(_gdg)
		_baa := _cec.Add(_fbg.Flip())
		_ca := _geef.AddVector(_baa)
		_daed := _fbg.Scale(2).Flip().Add(_baa.Flip())
		_eff := _ca.AddVector(_daed)
		_fac := _abd.AddVector(NewVectorPolar(_ddd, _cgb-_geg/2))
		_ggd := NewPath()
		_ggd = _ggd.AppendPoint(_bcf.GetPointNumber(1))
		_ggd = _ggd.AppendPoint(_abd)
		_ggd = _ggd.AppendPoint(_geef)
		_ggd = _ggd.AppendPoint(_ca)
		_ggd = _ggd.AppendPoint(_eff)
		_ggd = _ggd.AppendPoint(_fac)
		_ggd = _ggd.AppendPoint(_bcf.GetPointNumber(4))
		_bcf = _ggd
	}
	if _afgg == LineEndingStyleArrow {
		_ffd := _bcf.GetPointNumber(1)
		_eb := _bcf.GetPointNumber(_bcf.Length())
		_cccaa := NewVectorPolar(_ddd/2, _cgb+_geg+_geg/2)
		_def := _ffd.AddVector(_cccaa)
		_cda := NewVectorPolar(_acf, _cgb).Add(NewVectorPolar(_afa/2, _cgb+_geg/2))
		_bbf := _def.AddVector(_cda)
		_ad := NewVectorPolar(_edc, _cgb-_geg/2)
		_fbb := _bbf.AddVector(_ad)
		_cccc := NewVectorPolar(_acf, _cgb)
		_efg := _eb.AddVector(_cccc)
		_ga := NewVectorPolar(_edc, _cgb+_geg+_geg/2)
		_ebb := _efg.AddVector(_ga)
		_cdag := _def
		_fgc := NewPath()
		_fgc = _fgc.AppendPoint(_def)
		_fgc = _fgc.AppendPoint(_bbf)
		_fgc = _fgc.AppendPoint(_fbb)
		for _, _dfe := range _bcf.Points[1 : len(_bcf.Points)-1] {
			_fgc = _fgc.AppendPoint(_dfe)
		}
		_fgc = _fgc.AppendPoint(_efg)
		_fgc = _fgc.AppendPoint(_ebb)
		_fgc = _fgc.AppendPoint(_cdag)
		_bcf = _fgc
	}
	_ecc := _fc.NewContentCreator()
	_ecc.Add_q().SetNonStrokingColor(_dcb.LineColor)
	if len(gsName) > 1 {
		_ecc.Add_gs(_eg.PdfObjectName(gsName))
	}
	_bcf = _bcf.Offset(_dcb.X1, _dcb.Y1)
	_eae := _bcf.GetBoundingBox()
	DrawPathWithCreator(_bcf, _ecc)
	if _dcb.LineStyle == LineStyleDashed {
		_ecc.Add_d([]int64{1, 1}, 0).Add_S().Add_f().Add_Q()
	} else {
		_ecc.Add_f().Add_Q()
	}
	return _ecc.Bytes(), _eae.ToPdfRectangle(), nil
}

// NewPath returns a new empty path.
func NewPath() Path { return Path{} }

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct {
	Points        [][]Point
	FillEnabled   bool
	FillColor     _e.PdfColor
	BorderEnabled bool
	BorderColor   _e.PdfColor
	BorderWidth   float64
}

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int

// Copy returns a clone of the Bezier path.
func (_g CubicBezierPath) Copy() CubicBezierPath {
	_fe := CubicBezierPath{}
	_fe.Curves = append(_fe.Curves, _g.Curves...)
	return _fe
}

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_de Polygon) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	_gee := _fc.NewContentCreator()
	_gee.Add_q()
	_de.FillEnabled = _de.FillEnabled && _de.FillColor != nil
	if _de.FillEnabled {
		_gee.SetNonStrokingColor(_de.FillColor)
	}
	_de.BorderEnabled = _de.BorderEnabled && _de.BorderColor != nil
	if _de.BorderEnabled {
		_gee.SetStrokingColor(_de.BorderColor)
		_gee.Add_w(_de.BorderWidth)
	}
	if len(gsName) > 1 {
		_gee.Add_gs(_eg.PdfObjectName(gsName))
	}
	_fca := NewPath()
	for _, _dad := range _de.Points {
		for _be, _bed := range _dad {
			_fca = _fca.AppendPoint(_bed)
			if _be == 0 {
				_gee.Add_m(_bed.X, _bed.Y)
			} else {
				_gee.Add_l(_bed.X, _bed.Y)
			}
		}
		_gee.Add_h()
	}
	if _de.FillEnabled && _de.BorderEnabled {
		_gee.Add_B()
	} else if _de.FillEnabled {
		_gee.Add_f()
	} else if _de.BorderEnabled {
		_gee.Add_S()
	}
	_gee.Add_Q()
	return _gee.Bytes(), _fca.GetBoundingBox().ToPdfRectangle(), nil
}

// Length returns the number of points in the path.
func (_bd Path) Length() int { return len(_bd.Points) }

// GetBounds returns the bounding box of the Bezier curve.
func (_ba CubicBezierCurve) GetBounds() _e.PdfRectangle {
	_aa := _ba.P0.X
	_dc := _ba.P0.X
	_ee := _ba.P0.Y
	_aac := _ba.P0.Y
	for _c := 0.0; _c <= 1.0; _c += 0.001 {
		Rx := _ba.P0.X*_d.Pow(1-_c, 3) + _ba.P1.X*3*_c*_d.Pow(1-_c, 2) + _ba.P2.X*3*_d.Pow(_c, 2)*(1-_c) + _ba.P3.X*_d.Pow(_c, 3)
		Ry := _ba.P0.Y*_d.Pow(1-_c, 3) + _ba.P1.Y*3*_c*_d.Pow(1-_c, 2) + _ba.P2.Y*3*_d.Pow(_c, 2)*(1-_c) + _ba.P3.Y*_d.Pow(_c, 3)
		if Rx < _aa {
			_aa = Rx
		}
		if Rx > _dc {
			_dc = Rx
		}
		if Ry < _ee {
			_ee = Ry
		}
		if Ry > _aac {
			_aac = Ry
		}
	}
	_ab := _e.PdfRectangle{}
	_ab.Llx = _aa
	_ab.Lly = _ee
	_ab.Urx = _dc
	_ab.Ury = _aac
	return _ab
}

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_ecd BoundingBox) ToPdfRectangle() *_e.PdfRectangle {
	return &_e.PdfRectangle{Llx: _ecd.X, Lly: _ecd.Y, Urx: _ecd.X + _ecd.Width, Ury: _ecd.Y + _ecd.Height}
}

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_bedd Polyline) Draw(gsName string) ([]byte, *_e.PdfRectangle, error) {
	if _bedd.LineColor == nil {
		_bedd.LineColor = _e.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_gdc := NewPath()
	for _, _dadf := range _bedd.Points {
		_gdc = _gdc.AppendPoint(_dadf)
	}
	_gff := _fc.NewContentCreator()
	_gff.Add_q().SetStrokingColor(_bedd.LineColor).Add_w(_bedd.LineWidth)
	if len(gsName) > 1 {
		_gff.Add_gs(_eg.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_gdc, _gff)
	_gff.Add_S()
	_gff.Add_Q()
	return _gff.Bytes(), _gdc.GetBoundingBox().ToPdfRectangle(), nil
}

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_bb Point) Rotate(theta float64) Point {
	_ag := _egg.NewPoint(_bb.X, _bb.Y).Rotate(theta)
	return NewPoint(_ag.X, _ag.Y)
}
