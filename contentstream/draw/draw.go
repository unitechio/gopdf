package draw

import (
	_gf "fmt"
	_g "math"

	_gg "unitechio/gopdf/gopdf/contentstream"
	_c "unitechio/gopdf/gopdf/core"
	_e "unitechio/gopdf/gopdf/internal/transform"
	_gfc "unitechio/gopdf/gopdf/model"
)

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator(path Path, creator *_gg.ContentCreator) {
	for _gfd, _ebb := range path.Points {
		if _gfd == 0 {
			creator.Add_m(_ebb.X, _ebb.Y)
		} else {
			creator.Add_l(_ebb.X, _ebb.Y)
		}
	}
}

// Offset shifts the Bezier path with the specified offsets.
func (_da CubicBezierPath) Offset(offX, offY float64) CubicBezierPath {
	for _gb, _cf := range _da.Curves {
		_da.Curves[_gb] = _cf.AddOffsetXY(offX, offY)
	}
	return _da
}

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     _gfc.PdfColor
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   _gfc.PdfColor
	Opacity       float64
}

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_fda Point) Add(dx, dy float64) Point { _fda.X += dx; _fda.Y += dy; return _fda }

// GetBoundingBox returns the bounding box of the path.
func (_aac Path) GetBoundingBox() BoundingBox {
	_eg := BoundingBox{}
	_ddg := 0.0
	_cef := 0.0
	_fd := 0.0
	_gdc := 0.0
	for _ag, _fca := range _aac.Points {
		if _ag == 0 {
			_ddg = _fca.X
			_cef = _fca.X
			_fd = _fca.Y
			_gdc = _fca.Y
			continue
		}
		if _fca.X < _ddg {
			_ddg = _fca.X
		}
		if _fca.X > _cef {
			_cef = _fca.X
		}
		if _fca.Y < _fd {
			_fd = _fca.Y
		}
		if _fca.Y > _gdc {
			_gdc = _fca.Y
		}
	}
	_eg.X = _ddg
	_eg.Y = _fd
	_eg.Width = _cef - _ddg
	_eg.Height = _gdc - _fd
	return _eg
}

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{ Points []Point }

// AppendPoint adds the specified point to the path.
func (_fc Path) AppendPoint(point Point) Path { _fc.Points = append(_fc.Points, point); return _fc }

// Add adds the specified vector to the current one and returns the result.
func (_cfc Vector) Add(other Vector) Vector { _cfc.Dx += other.Dx; _cfc.Dy += other.Dy; return _cfc }

// AppendCurve appends the specified Bezier curve to the path.
func (_dc CubicBezierPath) AppendCurve(curve CubicBezierCurve) CubicBezierPath {
	_dc.Curves = append(_dc.Curves, curve)
	return _dc
}

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_cc BoundingBox) ToPdfRectangle() *_gfc.PdfRectangle {
	return &_gfc.PdfRectangle{Llx: _cc.X, Lly: _cc.Y, Urx: _cc.X + _cc.Width, Ury: _cc.Y + _cc.Height}
}

// LineStyle refers to how the line will be created.
type LineStyle int

// GetBounds returns the bounding box of the Bezier curve.
func (_ggg CubicBezierCurve) GetBounds() _gfc.PdfRectangle {
	_ee := _ggg.P0.X
	_b := _ggg.P0.X
	_af := _ggg.P0.Y
	_ed := _ggg.P0.Y
	for _bf := 0.0; _bf <= 1.0; _bf += 0.001 {
		Rx := _ggg.P0.X*_g.Pow(1-_bf, 3) + _ggg.P1.X*3*_bf*_g.Pow(1-_bf, 2) + _ggg.P2.X*3*_g.Pow(_bf, 2)*(1-_bf) + _ggg.P3.X*_g.Pow(_bf, 3)
		Ry := _ggg.P0.Y*_g.Pow(1-_bf, 3) + _ggg.P1.Y*3*_bf*_g.Pow(1-_bf, 2) + _ggg.P2.Y*3*_g.Pow(_bf, 2)*(1-_bf) + _ggg.P3.Y*_g.Pow(_bf, 3)
		if Rx < _ee {
			_ee = Rx
		}
		if Rx > _b {
			_b = Rx
		}
		if Ry < _af {
			_af = Ry
		}
		if Ry > _ed {
			_ed = Ry
		}
	}
	_df := _gfc.PdfRectangle{}
	_df.Llx = _ee
	_df.Lly = _af
	_df.Urx = _b
	_df.Ury = _ed
	return _df
}

// Offset shifts the path with the specified offsets.
func (_bfa Path) Offset(offX, offY float64) Path {
	for _cde, _cdb := range _bfa.Points {
		_bfa.Points[_cde] = _cdb.Add(offX, offY)
	}
	return _bfa
}

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        _gfc.PdfColor
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 LineEndingStyle
	LineEndingStyle2 LineEndingStyle
	LineStyle        LineStyle
}

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_fdg Polyline) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	if _fdg.LineColor == nil {
		_fdg.LineColor = _gfc.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_aea := NewPath()
	for _, _fcbd := range _fdg.Points {
		_aea = _aea.AppendPoint(_fcbd)
	}
	_ccgd := _gg.NewContentCreator()
	_ccgd.Add_q().SetStrokingColor(_fdg.LineColor).Add_w(_fdg.LineWidth)
	if len(gsName) > 1 {
		_ccgd.Add_gs(_c.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_aea, _ccgd)
	_ccgd.Add_S()
	_ccgd.Add_Q()
	return _ccgd.Bytes(), _aea.GetBoundingBox().ToPdfRectangle(), nil
}

// Copy returns a clone of the Bezier path.
func (_bb CubicBezierPath) Copy() CubicBezierPath {
	_db := CubicBezierPath{}
	_db.Curves = append(_db.Curves, _bb.Curves...)
	return _db
}

func (_gbd Point) String() string {
	return _gf.Sprintf("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029", _gbd.X, _gbd.Y)
}

// AddVector adds vector to a point.
func (_ac Point) AddVector(v Vector) Point { _ac.X += v.Dx; _ac.Y += v.Dy; return _ac }

// NewPoint returns a new point with the coordinates x, y.
func NewPoint(x, y float64) Point { return Point{X: x, Y: y} }

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar(length float64, theta float64) Vector {
	_ffb := Vector{}
	_ffb.Dx = length * _g.Cos(theta)
	_ffb.Dy = length * _g.Sin(theta)
	return _ffb
}

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_d CubicBezierCurve) AddOffsetXY(offX, offY float64) CubicBezierCurve {
	_d.P0.X += offX
	_d.P1.X += offX
	_d.P2.X += offX
	_d.P3.X += offX
	_d.P0.Y += offY
	_d.P1.Y += offY
	_d.P2.Y += offY
	_d.P3.Y += offY
	return _d
}

// Copy returns a clone of the path.
func (_dd Path) Copy() Path {
	_ce := Path{}
	_ce.Points = append(_ce.Points, _dd.Points...)
	return _ce
}

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_cdf Path) GetPointNumber(number int) Point {
	if number < 1 || number > len(_cdf.Points) {
		return Point{}
	}
	return _cdf.Points[number-1]
}

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator(bpath CubicBezierPath, creator *_gg.ContentCreator) {
	for _dcg, _bde := range bpath.Curves {
		if _dcg == 0 {
			creator.Add_m(_bde.P0.X, _bde.P0.Y)
		}
		creator.Add_c(_bde.P1.X, _bde.P1.Y, _bde.P2.X, _bde.P2.Y, _bde.P3.X, _bde.P3.Y)
	}
}

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector(dx, dy float64) Vector { _dcbf := Vector{}; _dcbf.Dx = dx; _dcbf.Dy = dy; return _dcbf }

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_aab Point) Rotate(theta float64) Point {
	_ad := _e.NewPoint(_aab.X, _aab.Y).Rotate(theta)
	return NewPoint(_ad.X, _ad.Y)
}

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_ade PolyBezierCurve) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	if _ade.BorderColor == nil {
		_ade.BorderColor = _gfc.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_dfg := NewCubicBezierPath()
	for _, _fe := range _ade.Curves {
		_dfg = _dfg.AppendCurve(_fe)
	}
	_gaf := _gg.NewContentCreator()
	_gaf.Add_q()
	_ade.FillEnabled = _ade.FillEnabled && _ade.FillColor != nil
	if _ade.FillEnabled {
		_gaf.SetNonStrokingColor(_ade.FillColor)
	}
	_gaf.SetStrokingColor(_ade.BorderColor)
	_gaf.Add_w(_ade.BorderWidth)
	if len(gsName) > 1 {
		_gaf.Add_gs(_c.PdfObjectName(gsName))
	}
	for _cg, _bg := range _dfg.Curves {
		if _cg == 0 {
			_gaf.Add_m(_bg.P0.X, _bg.P0.Y)
		} else {
			_gaf.Add_l(_bg.P0.X, _bg.P0.Y)
		}
		_gaf.Add_c(_bg.P1.X, _bg.P1.Y, _bg.P2.X, _bg.P2.Y, _bg.P3.X, _bg.P3.Y)
	}
	if _ade.FillEnabled {
		_gaf.Add_h()
		_gaf.Add_B()
	} else {
		_gaf.Add_S()
	}
	_gaf.Add_Q()
	return _gaf.Bytes(), _dfg.GetBoundingBox().ToPdfRectangle(), nil
}

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_fg CurvePolygon) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	_aba := _gg.NewContentCreator()
	_aba.Add_q()
	_fg.FillEnabled = _fg.FillEnabled && _fg.FillColor != nil
	if _fg.FillEnabled {
		_aba.SetNonStrokingColor(_fg.FillColor)
	}
	_fg.BorderEnabled = _fg.BorderEnabled && _fg.BorderColor != nil
	if _fg.BorderEnabled {
		_aba.SetStrokingColor(_fg.BorderColor)
		_aba.Add_w(_fg.BorderWidth)
	}
	if len(gsName) > 1 {
		_aba.Add_gs(_c.PdfObjectName(gsName))
	}
	_eec := NewCubicBezierPath()
	for _, _gae := range _fg.Rings {
		for _aaf, _gbf := range _gae {
			if _aaf == 0 {
				_aba.Add_m(_gbf.P0.X, _gbf.P0.Y)
			} else {
				_aba.Add_l(_gbf.P0.X, _gbf.P0.Y)
			}
			_aba.Add_c(_gbf.P1.X, _gbf.P1.Y, _gbf.P2.X, _gbf.P2.Y, _gbf.P3.X, _gbf.P3.Y)
			_eec = _eec.AppendCurve(_gbf)
		}
		_aba.Add_h()
	}
	if _fg.FillEnabled && _fg.BorderEnabled {
		_aba.Add_B()
	} else if _fg.FillEnabled {
		_aba.Add_f()
	} else if _fg.BorderEnabled {
		_aba.Add_S()
	}
	_aba.Add_Q()
	return _aba.Bytes(), _eec.GetBoundingBox().ToPdfRectangle(), nil
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
	FillColor   _gfc.PdfColor

	// Border properties.
	BorderEnabled           bool
	BorderColor             _gfc.PdfColor
	BorderWidth             float64
	BorderRadiusTopLeft     float64
	BorderRadiusTopRight    float64
	BorderRadiusBottomLeft  float64
	BorderRadiusBottomRight float64

	// Shape opacity (0-1 interval).
	Opacity float64
}

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{ Curves []CubicBezierCurve }

const (
	LineEndingStyleNone  LineEndingStyle = 0
	LineEndingStyleArrow LineEndingStyle = 1
	LineEndingStyleButt  LineEndingStyle = 2
)

// Length returns the number of points in the path.
func (_eeg Path) Length() int { return len(_eeg.Points) }

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_gga BasicLine) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	_gbg := NewPath()
	_gbg = _gbg.AppendPoint(NewPoint(_gga.X1, _gga.Y1))
	_gbg = _gbg.AppendPoint(NewPoint(_gga.X2, _gga.Y2))
	_dda := _gg.NewContentCreator()
	_dda.Add_q().Add_w(_gga.LineWidth).SetStrokingColor(_gga.LineColor)
	if _gga.LineStyle == LineStyleDashed {
		if _gga.DashArray == nil {
			_gga.DashArray = []int64{1, 1}
		}
		_dda.Add_d(_gga.DashArray, _gga.DashPhase)
	}
	if len(gsName) > 1 {
		_dda.Add_gs(_c.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_gbg, _dda)
	_dda.Add_S().Add_Q()
	return _dda.Bytes(), _gbg.GetBoundingBox().ToPdfRectangle(), nil
}

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_bfg Line) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	_ebc, _gag := _bfg.X1, _bfg.X2
	_adef, _cb := _bfg.Y1, _bfg.Y2
	_fbg := _cb - _adef
	_gfb := _gag - _ebc
	_feb := _g.Atan2(_fbg, _gfb)
	L := _g.Sqrt(_g.Pow(_gfb, 2.0) + _g.Pow(_fbg, 2.0))
	_gee := _bfg.LineWidth
	_ebf := _g.Pi
	_beg := 1.0
	if _gfb < 0 {
		_beg *= -1.0
	}
	if _fbg < 0 {
		_beg *= -1.0
	}
	VsX := _beg * (-_gee / 2 * _g.Cos(_feb+_ebf/2))
	VsY := _beg * (-_gee/2*_g.Sin(_feb+_ebf/2) + _gee*_g.Sin(_feb+_ebf/2))
	V1X := VsX + _gee/2*_g.Cos(_feb+_ebf/2)
	V1Y := VsY + _gee/2*_g.Sin(_feb+_ebf/2)
	V2X := VsX + _gee/2*_g.Cos(_feb+_ebf/2) + L*_g.Cos(_feb)
	V2Y := VsY + _gee/2*_g.Sin(_feb+_ebf/2) + L*_g.Sin(_feb)
	V3X := VsX + _gee/2*_g.Cos(_feb+_ebf/2) + L*_g.Cos(_feb) + _gee*_g.Cos(_feb-_ebf/2)
	V3Y := VsY + _gee/2*_g.Sin(_feb+_ebf/2) + L*_g.Sin(_feb) + _gee*_g.Sin(_feb-_ebf/2)
	V4X := VsX + _gee/2*_g.Cos(_feb-_ebf/2)
	V4Y := VsY + _gee/2*_g.Sin(_feb-_ebf/2)
	_ea := NewPath()
	_ea = _ea.AppendPoint(NewPoint(V1X, V1Y))
	_ea = _ea.AppendPoint(NewPoint(V2X, V2Y))
	_ea = _ea.AppendPoint(NewPoint(V3X, V3Y))
	_ea = _ea.AppendPoint(NewPoint(V4X, V4Y))
	_cgf := _bfg.LineEndingStyle1
	_bdga := _bfg.LineEndingStyle2
	_gaef := 3 * _gee
	_fgb := 3 * _gee
	_baa := (_fgb - _gee) / 2
	if _bdga == LineEndingStyleArrow {
		_aff := _ea.GetPointNumber(2)
		_ff := NewVectorPolar(_gaef, _feb+_ebf)
		_cdc := _aff.AddVector(_ff)
		_bfb := NewVectorPolar(_fgb/2, _feb+_ebf/2)
		_gggd := NewVectorPolar(_gaef, _feb)
		_eee := NewVectorPolar(_baa, _feb+_ebf/2)
		_dcf := _cdc.AddVector(_eee)
		_abae := _gggd.Add(_bfb.Flip())
		_aabd := _dcf.AddVector(_abae)
		_bgg := _bfb.Scale(2).Flip().Add(_abae.Flip())
		_adb := _aabd.AddVector(_bgg)
		_dad := _cdc.AddVector(NewVectorPolar(_gee, _feb-_ebf/2))
		_dcc := NewPath()
		_dcc = _dcc.AppendPoint(_ea.GetPointNumber(1))
		_dcc = _dcc.AppendPoint(_cdc)
		_dcc = _dcc.AppendPoint(_dcf)
		_dcc = _dcc.AppendPoint(_aabd)
		_dcc = _dcc.AppendPoint(_adb)
		_dcc = _dcc.AppendPoint(_dad)
		_dcc = _dcc.AppendPoint(_ea.GetPointNumber(4))
		_ea = _dcc
	}
	if _cgf == LineEndingStyleArrow {
		_efe := _ea.GetPointNumber(1)
		_cbf := _ea.GetPointNumber(_ea.Length())
		_gbfb := NewVectorPolar(_gee/2, _feb+_ebf+_ebf/2)
		_deg := _efe.AddVector(_gbfb)
		_bag := NewVectorPolar(_gaef, _feb).Add(NewVectorPolar(_fgb/2, _feb+_ebf/2))
		_cge := _deg.AddVector(_bag)
		_fdf := NewVectorPolar(_baa, _feb-_ebf/2)
		_dfa := _cge.AddVector(_fdf)
		_aafd := NewVectorPolar(_gaef, _feb)
		_fbb := _cbf.AddVector(_aafd)
		_fea := NewVectorPolar(_baa, _feb+_ebf+_ebf/2)
		_fga := _fbb.AddVector(_fea)
		_dbbc := _deg
		_ccg := NewPath()
		_ccg = _ccg.AppendPoint(_deg)
		_ccg = _ccg.AppendPoint(_cge)
		_ccg = _ccg.AppendPoint(_dfa)
		for _, _cgfe := range _ea.Points[1 : len(_ea.Points)-1] {
			_ccg = _ccg.AppendPoint(_cgfe)
		}
		_ccg = _ccg.AppendPoint(_fbb)
		_ccg = _ccg.AppendPoint(_fga)
		_ccg = _ccg.AppendPoint(_dbbc)
		_ea = _ccg
	}
	_fdag := _gg.NewContentCreator()
	_fdag.Add_q().SetNonStrokingColor(_bfg.LineColor)
	if len(gsName) > 1 {
		_fdag.Add_gs(_c.PdfObjectName(gsName))
	}
	_ea = _ea.Offset(_bfg.X1, _bfg.Y1)
	_dga := _ea.GetBoundingBox()
	DrawPathWithCreator(_ea, _fdag)
	if _bfg.LineStyle == LineStyleDashed {
		_fdag.Add_d([]int64{1, 1}, 0).Add_S().Add_f().Add_Q()
	} else {
		_fdag.Add_f().Add_Q()
	}
	return _fdag.Bytes(), _dga.ToPdfRectangle(), nil
}

// Rotate rotates the vector by the specified angle.
func (_dca Vector) Rotate(phi float64) Vector {
	_fcbdf := _dca.Magnitude()
	_cefe := _dca.GetPolarAngle()
	return NewVectorPolar(_fcbdf, _cefe+phi)
}

// FlipY flips the sign of the Dy component of the vector.
func (_dbe Vector) FlipY() Vector { _dbe.Dy = -_dbe.Dy; return _dbe }

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_aad Polygon) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	_ef := _gg.NewContentCreator()
	_ef.Add_q()
	_aad.FillEnabled = _aad.FillEnabled && _aad.FillColor != nil
	if _aad.FillEnabled {
		_ef.SetNonStrokingColor(_aad.FillColor)
	}
	_aad.BorderEnabled = _aad.BorderEnabled && _aad.BorderColor != nil
	if _aad.BorderEnabled {
		_ef.SetStrokingColor(_aad.BorderColor)
		_ef.Add_w(_aad.BorderWidth)
	}
	if len(gsName) > 1 {
		_ef.Add_gs(_c.PdfObjectName(gsName))
	}
	_fcf := NewPath()
	for _, _bd := range _aad.Points {
		for _eb, _fbd := range _bd {
			_fcf = _fcf.AppendPoint(_fbd)
			if _eb == 0 {
				_ef.Add_m(_fbd.X, _fbd.Y)
			} else {
				_ef.Add_l(_fbd.X, _fbd.Y)
			}
		}
		_ef.Add_h()
	}
	if _aad.FillEnabled && _aad.BorderEnabled {
		_ef.Add_B()
	} else if _aad.FillEnabled {
		_ef.Add_f()
	} else if _aad.BorderEnabled {
		_ef.Add_S()
	}
	_ef.Add_Q()
	return _ef.Bytes(), _fcf.GetBoundingBox().ToPdfRectangle(), nil
}

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve(x0, y0, x1, y1, x2, y2, x3, y3 float64) CubicBezierCurve {
	_aa := CubicBezierCurve{}
	_aa.P0 = NewPoint(x0, y0)
	_aa.P1 = NewPoint(x1, y1)
	_aa.P2 = NewPoint(x2, y2)
	_aa.P3 = NewPoint(x3, y3)
	return _aa
}

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct {
	Points    []Point
	LineColor _gfc.PdfColor
	LineWidth float64
}

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Scale scales the vector by the specified factor.
func (_bdgg Vector) Scale(factor float64) Vector {
	_acc := _bdgg.Magnitude()
	_edb := _bdgg.GetPolarAngle()
	_bdgg.Dx = factor * _acc * _g.Cos(_edb)
	_bdgg.Dy = factor * _acc * _g.Sin(_edb)
	return _bdgg
}

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_ae Path) RemovePoint(number int) Path {
	if number < 1 || number > len(_ae.Points) {
		return _ae
	}
	_ga := number - 1
	_ae.Points = append(_ae.Points[:_ga], _ae.Points[_ga+1:]...)
	return _ae
}

const (
	LineStyleSolid  LineStyle = 0
	LineStyleDashed LineStyle = 1
)

// NewPath returns a new empty path.
func NewPath() Path { return Path{} }

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween(a Point, b Point) Vector {
	_dbbcb := Vector{}
	_dbbcb.Dx = b.X - a.X
	_dbbcb.Dy = b.Y - a.Y
	return _dbbcb
}

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_ggaa Vector) GetPolarAngle() float64 { return _g.Atan2(_ggaa.Dy, _ggaa.Dx) }

// GetBoundingBox returns the bounding box of the Bezier path.
func (_be CubicBezierPath) GetBoundingBox() Rectangle {
	_ede := Rectangle{}
	_dac := 0.0
	_gd := 0.0
	_gdf := 0.0
	_f := 0.0
	for _cd, _ab := range _be.Curves {
		_dg := _ab.GetBounds()
		if _cd == 0 {
			_dac = _dg.Llx
			_gd = _dg.Urx
			_gdf = _dg.Lly
			_f = _dg.Ury
			continue
		}
		if _dg.Llx < _dac {
			_dac = _dg.Llx
		}
		if _dg.Urx > _gd {
			_gd = _dg.Urx
		}
		if _dg.Lly < _gdf {
			_gdf = _dg.Lly
		}
		if _dg.Ury > _f {
			_f = _dg.Ury
		}
	}
	_ede.X = _dac
	_ede.Y = _gdf
	_ede.Width = _gd - _dac
	_ede.Height = _f - _gdf
	return _ede
}

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct {
	Rings         [][]CubicBezierCurve
	FillEnabled   bool
	FillColor     _gfc.PdfColor
	BorderEnabled bool
	BorderColor   _gfc.PdfColor
	BorderWidth   float64
}

// Point represents a two-dimensional point.
type Point struct {
	X float64
	Y float64
}

// FlipX flips the sign of the Dx component of the vector.
func (_gcf Vector) FlipX() Vector { _gcf.Dx = -_gcf.Dx; return _gcf }

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct {
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	LineColor _gfc.PdfColor
	Opacity   float64
	LineWidth float64
	LineStyle LineStyle
	DashArray []int64
	DashPhase int64
}

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct {
	Curves      []CubicBezierCurve
	BorderWidth float64
	BorderColor _gfc.PdfColor
	FillEnabled bool
	FillColor   _gfc.PdfColor
}

// Draw draws the rectangle. A graphics state can be specified for
// setting additional properties (e.g. opacity). Otherwise pass an empty string
// for the `gsName` parameter. The method returns the content stream as a byte
// array and the bounding box of the shape.
func (_gc Rectangle) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	_abe := _gg.NewContentCreator()
	_abe.Add_q()
	if _gc.FillEnabled {
		_abe.SetNonStrokingColor(_gc.FillColor)
	}
	if _gc.BorderEnabled {
		_abe.SetStrokingColor(_gc.BorderColor)
		_abe.Add_w(_gc.BorderWidth)
	}
	if len(gsName) > 1 {
		_abe.Add_gs(_c.PdfObjectName(gsName))
	}
	var (
		_dbbd, _bda = _gc.X, _gc.Y
		_fcag, _bdg = _gc.Width, _gc.Height
		_aae        = _g.Abs(_gc.BorderRadiusTopLeft)
		_eef        = _g.Abs(_gc.BorderRadiusTopRight)
		_cfe        = _g.Abs(_gc.BorderRadiusBottomLeft)
		_acg        = _g.Abs(_gc.BorderRadiusBottomRight)
		_ddf        = 0.4477
	)
	_aabf := Path{Points: []Point{{X: _dbbd + _fcag - _acg, Y: _bda}, {X: _dbbd + _fcag, Y: _bda + _bdg - _eef}, {X: _dbbd + _aae, Y: _bda + _bdg}, {X: _dbbd, Y: _bda + _cfe}}}
	_acb := [][7]float64{{_acg, _dbbd + _fcag - _acg*_ddf, _bda, _dbbd + _fcag, _bda + _acg*_ddf, _dbbd + _fcag, _bda + _acg}, {_eef, _dbbd + _fcag, _bda + _bdg - _eef*_ddf, _dbbd + _fcag - _eef*_ddf, _bda + _bdg, _dbbd + _fcag - _eef, _bda + _bdg}, {_aae, _dbbd + _aae*_ddf, _bda + _bdg, _dbbd, _bda + _bdg - _aae*_ddf, _dbbd, _bda + _bdg - _aae}, {_cfe, _dbbd, _bda + _cfe*_ddf, _dbbd + _cfe*_ddf, _bda, _dbbd + _cfe, _bda}}
	_abe.Add_m(_dbbd+_cfe, _bda)
	for _dcb := 0; _dcb < 4; _dcb++ {
		_ge := _aabf.Points[_dcb]
		_abe.Add_l(_ge.X, _ge.Y)
		_bc := _acb[_dcb]
		if _ba := _bc[0]; _ba != 0 {
			_abe.Add_c(_bc[1], _bc[2], _bc[3], _bc[4], _bc[5], _bc[6])
		}
	}
	_abe.Add_h()
	if _gc.FillEnabled && _gc.BorderEnabled {
		_abe.Add_B()
	} else if _gc.FillEnabled {
		_abe.Add_f()
	} else if _gc.BorderEnabled {
		_abe.Add_S()
	}
	_abe.Add_Q()
	return _abe.Bytes(), _aabf.GetBoundingBox().ToPdfRectangle(), nil
}

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct {
	Points        [][]Point
	FillEnabled   bool
	FillColor     _gfc.PdfColor
	BorderEnabled bool
	BorderColor   _gfc.PdfColor
	BorderWidth   float64
}

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_dfc Circle) Draw(gsName string) ([]byte, *_gfc.PdfRectangle, error) {
	_bfd := _dfc.Width / 2
	_edg := _dfc.Height / 2
	if _dfc.BorderEnabled {
		_bfd -= _dfc.BorderWidth / 2
		_edg -= _dfc.BorderWidth / 2
	}
	_fce := 0.551784
	_cee := _bfd * _fce
	_adf := _edg * _fce
	_cfd := NewCubicBezierPath()
	_cfd = _cfd.AppendCurve(NewCubicBezierCurve(-_bfd, 0, -_bfd, _adf, -_cee, _edg, 0, _edg))
	_cfd = _cfd.AppendCurve(NewCubicBezierCurve(0, _edg, _cee, _edg, _bfd, _adf, _bfd, 0))
	_cfd = _cfd.AppendCurve(NewCubicBezierCurve(_bfd, 0, _bfd, -_adf, _cee, -_edg, 0, -_edg))
	_cfd = _cfd.AppendCurve(NewCubicBezierCurve(0, -_edg, -_cee, -_edg, -_bfd, -_adf, -_bfd, 0))
	_cfd = _cfd.Offset(_bfd, _edg)
	if _dfc.BorderEnabled {
		_cfd = _cfd.Offset(_dfc.BorderWidth/2, _dfc.BorderWidth/2)
	}
	if _dfc.X != 0 || _dfc.Y != 0 {
		_cfd = _cfd.Offset(_dfc.X, _dfc.Y)
	}
	_abc := _gg.NewContentCreator()
	_abc.Add_q()
	if _dfc.FillEnabled {
		_abc.SetNonStrokingColor(_dfc.FillColor)
	}
	if _dfc.BorderEnabled {
		_abc.SetStrokingColor(_dfc.BorderColor)
		_abc.Add_w(_dfc.BorderWidth)
	}
	if len(gsName) > 1 {
		_abc.Add_gs(_c.PdfObjectName(gsName))
	}
	DrawBezierPathWithCreator(_cfd, _abc)
	_abc.Add_h()
	if _dfc.FillEnabled && _dfc.BorderEnabled {
		_abc.Add_B()
	} else if _dfc.FillEnabled {
		_abc.Add_f()
	} else if _dfc.BorderEnabled {
		_abc.Add_S()
	}
	_abc.Add_Q()
	_dbb := _cfd.GetBoundingBox()
	if _dfc.BorderEnabled {
		_dbb.Height += _dfc.BorderWidth
		_dbb.Width += _dfc.BorderWidth
		_dbb.X -= _dfc.BorderWidth / 2
		_dbb.Y -= _dfc.BorderWidth / 2
	}
	return _abc.Bytes(), _dbb.ToPdfRectangle(), nil
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

// Vector represents a two-dimensional vector.
type Vector struct {
	Dx float64
	Dy float64
}

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_de Rectangle) ToPdfRectangle() *_gfc.PdfRectangle {
	return &_gfc.PdfRectangle{Llx: _de.X, Lly: _de.Y, Urx: _de.X + _de.Width, Ury: _de.Y + _de.Height}
}

// Flip changes the sign of the vector: -vector.
func (_ddb Vector) Flip() Vector {
	_gaa := _ddb.Magnitude()
	_cce := _ddb.GetPolarAngle()
	_ddb.Dx = _gaa * _g.Cos(_cce+_g.Pi)
	_ddb.Dy = _gaa * _g.Sin(_cce+_g.Pi)
	return _ddb
}

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath() CubicBezierPath {
	_ca := CubicBezierPath{}
	_ca.Curves = []CubicBezierCurve{}
	return _ca
}

// Magnitude returns the magnitude of the vector.
func (_abcf Vector) Magnitude() float64 {
	return _g.Sqrt(_g.Pow(_abcf.Dx, 2.0) + _g.Pow(_abcf.Dy, 2.0))
}
