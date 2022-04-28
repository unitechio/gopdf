package draw

import (
	_b "fmt"
	_a "math"

	_c "bitbucket.org/shenghui0779/gopdf/contentstream"
	_gb "bitbucket.org/shenghui0779/gopdf/core"
	_ad "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_f "bitbucket.org/shenghui0779/gopdf/model"
)

// FlipY flips the sign of the Dy component of the vector.
func (_ddb Vector) FlipY() Vector { _ddb.Dy = -_ddb.Dy; return _ddb }

// GetBounds returns the bounding box of the Bezier curve.
func (_gbd CubicBezierCurve) GetBounds() _f.PdfRectangle {
	_d := _gbd.P0.X
	_de := _gbd.P0.X
	_fe := _gbd.P0.Y
	_ag := _gbd.P0.Y
	for _bf := 0.0; _bf <= 1.0; _bf += 0.001 {
		Rx := _gbd.P0.X*_a.Pow(1-_bf, 3) + _gbd.P1.X*3*_bf*_a.Pow(1-_bf, 2) + _gbd.P2.X*3*_a.Pow(_bf, 2)*(1-_bf) + _gbd.P3.X*_a.Pow(_bf, 3)
		Ry := _gbd.P0.Y*_a.Pow(1-_bf, 3) + _gbd.P1.Y*3*_bf*_a.Pow(1-_bf, 2) + _gbd.P2.Y*3*_a.Pow(_bf, 2)*(1-_bf) + _gbd.P3.Y*_a.Pow(_bf, 3)
		if Rx < _d {
			_d = Rx
		}
		if Rx > _de {
			_de = Rx
		}
		if Ry < _fe {
			_fe = Ry
		}
		if Ry > _ag {
			_ag = Ry
		}
	}
	_bbf := _f.PdfRectangle{}
	_bbf.Llx = _d
	_bbf.Lly = _fe
	_bbf.Urx = _de
	_bbf.Ury = _ag
	return _bbf
}

// NewPath returns a new empty path.
func NewPath() Path { return Path{} }

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        _f.PdfColor
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 LineEndingStyle
	LineEndingStyle2 LineEndingStyle
	LineStyle        LineStyle
}

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     _f.PdfColor
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   _f.PdfColor
	Opacity       float64
}

// FlipX flips the sign of the Dx component of the vector.
func (_gcda Vector) FlipX() Vector { _gcda.Dx = -_gcda.Dx; return _gcda }

// LineStyle refers to how the line will be created.
type LineStyle int

// AddVector adds vector to a point.
func (_bbd Point) AddVector(v Vector) Point { _bbd.X += v.Dx; _bbd.Y += v.Dy; return _bbd }

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_cgf Line) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	_dea, _aa := _cgf.X1, _cgf.X2
	_aed, _egc := _cgf.Y1, _cgf.Y2
	_gc := _egc - _aed
	_dfb := _aa - _dea
	_abg := _a.Atan2(_gc, _dfb)
	L := _a.Sqrt(_a.Pow(_dfb, 2.0) + _a.Pow(_gc, 2.0))
	_gf := _cgf.LineWidth
	_dff := _a.Pi
	_bgb := 1.0
	if _dfb < 0 {
		_bgb *= -1.0
	}
	if _gc < 0 {
		_bgb *= -1.0
	}
	VsX := _bgb * (-_gf / 2 * _a.Cos(_abg+_dff/2))
	VsY := _bgb * (-_gf/2*_a.Sin(_abg+_dff/2) + _gf*_a.Sin(_abg+_dff/2))
	V1X := VsX + _gf/2*_a.Cos(_abg+_dff/2)
	V1Y := VsY + _gf/2*_a.Sin(_abg+_dff/2)
	V2X := VsX + _gf/2*_a.Cos(_abg+_dff/2) + L*_a.Cos(_abg)
	V2Y := VsY + _gf/2*_a.Sin(_abg+_dff/2) + L*_a.Sin(_abg)
	V3X := VsX + _gf/2*_a.Cos(_abg+_dff/2) + L*_a.Cos(_abg) + _gf*_a.Cos(_abg-_dff/2)
	V3Y := VsY + _gf/2*_a.Sin(_abg+_dff/2) + L*_a.Sin(_abg) + _gf*_a.Sin(_abg-_dff/2)
	V4X := VsX + _gf/2*_a.Cos(_abg-_dff/2)
	V4Y := VsY + _gf/2*_a.Sin(_abg-_dff/2)
	_aaf := NewPath()
	_aaf = _aaf.AppendPoint(NewPoint(V1X, V1Y))
	_aaf = _aaf.AppendPoint(NewPoint(V2X, V2Y))
	_aaf = _aaf.AppendPoint(NewPoint(V3X, V3Y))
	_aaf = _aaf.AppendPoint(NewPoint(V4X, V4Y))
	_baf := _cgf.LineEndingStyle1
	_cf := _cgf.LineEndingStyle2
	_acgd := 3 * _gf
	_dbf := 3 * _gf
	_gce := (_dbf - _gf) / 2
	if _cf == LineEndingStyleArrow {
		_bbe := _aaf.GetPointNumber(2)
		_bce := NewVectorPolar(_acgd, _abg+_dff)
		_eeg := _bbe.AddVector(_bce)
		_ec := NewVectorPolar(_dbf/2, _abg+_dff/2)
		_dgcb := NewVectorPolar(_acgd, _abg)
		_eca := NewVectorPolar(_gce, _abg+_dff/2)
		_fff := _eeg.AddVector(_eca)
		_cdd := _dgcb.Add(_ec.Flip())
		_gba := _fff.AddVector(_cdd)
		_ddfa := _ec.Scale(2).Flip().Add(_cdd.Flip())
		_feg := _gba.AddVector(_ddfa)
		_fffc := _eeg.AddVector(NewVectorPolar(_gf, _abg-_dff/2))
		_aea := NewPath()
		_aea = _aea.AppendPoint(_aaf.GetPointNumber(1))
		_aea = _aea.AppendPoint(_eeg)
		_aea = _aea.AppendPoint(_fff)
		_aea = _aea.AppendPoint(_gba)
		_aea = _aea.AppendPoint(_feg)
		_aea = _aea.AppendPoint(_fffc)
		_aea = _aea.AppendPoint(_aaf.GetPointNumber(4))
		_aaf = _aea
	}
	if _baf == LineEndingStyleArrow {
		_eec := _aaf.GetPointNumber(1)
		_ace := _aaf.GetPointNumber(_aaf.Length())
		_dad := NewVectorPolar(_gf/2, _abg+_dff+_dff/2)
		_cac := _eec.AddVector(_dad)
		_gd := NewVectorPolar(_acgd, _abg).Add(NewVectorPolar(_dbf/2, _abg+_dff/2))
		_bafa := _cac.AddVector(_gd)
		_egf := NewVectorPolar(_gce, _abg-_dff/2)
		_cfc := _bafa.AddVector(_egf)
		_gbb := NewVectorPolar(_acgd, _abg)
		_fbd := _ace.AddVector(_gbb)
		_gfe := NewVectorPolar(_gce, _abg+_dff+_dff/2)
		_fc := _fbd.AddVector(_gfe)
		_caf := _cac
		_fg := NewPath()
		_fg = _fg.AppendPoint(_cac)
		_fg = _fg.AppendPoint(_bafa)
		_fg = _fg.AppendPoint(_cfc)
		for _, _bfgg := range _aaf.Points[1 : len(_aaf.Points)-1] {
			_fg = _fg.AppendPoint(_bfgg)
		}
		_fg = _fg.AppendPoint(_fbd)
		_fg = _fg.AppendPoint(_fc)
		_fg = _fg.AppendPoint(_caf)
		_aaf = _fg
	}
	_fgf := _c.NewContentCreator()
	_fgf.Add_q().SetNonStrokingColor(_cgf.LineColor)
	if len(gsName) > 1 {
		_fgf.Add_gs(_gb.PdfObjectName(gsName))
	}
	_aaf = _aaf.Offset(_cgf.X1, _cgf.Y1)
	_dbg := _aaf.GetBoundingBox()
	DrawPathWithCreator(_aaf, _fgf)
	if _cgf.LineStyle == LineStyleDashed {
		_fgf.Add_d([]int64{1, 1}, 0).Add_S().Add_f().Add_Q()
	} else {
		_fgf.Add_f().Add_Q()
	}
	return _fgf.Bytes(), _dbg.ToPdfRectangle(), nil
}

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{ Curves []CubicBezierCurve }

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_afcg Polyline) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	if _afcg.LineColor == nil {
		_afcg.LineColor = _f.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_ce := NewPath()
	for _, _bag := range _afcg.Points {
		_ce = _ce.AppendPoint(_bag)
	}
	_deb := _c.NewContentCreator()
	_deb.Add_q().SetStrokingColor(_afcg.LineColor).Add_w(_afcg.LineWidth)
	if len(gsName) > 1 {
		_deb.Add_gs(_gb.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_ce, _deb)
	_deb.Add_S()
	_deb.Add_Q()
	return _deb.Bytes(), _ce.GetBoundingBox().ToPdfRectangle(), nil
}

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_cdf Point) Add(dx, dy float64) Point { _cdf.X += dx; _cdf.Y += dy; return _cdf }

// Copy returns a clone of the path.
func (_cd Path) Copy() Path {
	_bc := Path{}
	_bc.Points = append(_bc.Points, _cd.Points...)
	return _bc
}

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct {
	Points        [][]Point
	FillEnabled   bool
	FillColor     _f.PdfColor
	BorderEnabled bool
	BorderColor   _f.PdfColor
	BorderWidth   float64
}

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_agbb Vector) GetPolarAngle() float64 { return _a.Atan2(_agbb.Dy, _agbb.Dx) }

// Offset shifts the Bezier path with the specified offsets.
func (_gg CubicBezierPath) Offset(offX, offY float64) CubicBezierPath {
	for _ff, _bg := range _gg.Curves {
		_gg.Curves[_ff] = _bg.AddOffsetXY(offX, offY)
	}
	return _gg
}

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_ffc BasicLine) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	_fgc := _ffc.LineWidth
	_be := NewPath()
	_be = _be.AppendPoint(NewPoint(_ffc.X1, _ffc.Y1))
	_be = _be.AppendPoint(NewPoint(_ffc.X2, _ffc.Y2))
	_adg := _c.NewContentCreator()
	_abd := _be.GetBoundingBox()
	DrawPathWithCreator(_be, _adg)
	if _ffc.LineStyle == LineStyleDashed {
		_adg.Add_d([]int64{1, 1}, 0)
	}
	_adg.SetStrokingColor(_ffc.LineColor).Add_w(_fgc).Add_S().Add_Q()
	return _adg.Bytes(), _abd.ToPdfRectangle(), nil
}

// Magnitude returns the magnitude of the vector.
func (_dcd Vector) Magnitude() float64 { return _a.Sqrt(_a.Pow(_dcd.Dx, 2.0) + _a.Pow(_dcd.Dy, 2.0)) }

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_fec Path) RemovePoint(number int) Path {
	if number < 1 || number > len(_fec.Points) {
		return _fec
	}
	_dc := number - 1
	_fec.Points = append(_fec.Points[:_dc], _fec.Points[_dc+1:]...)
	return _fec
}

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct {
	Points    []Point
	LineColor _f.PdfColor
	LineWidth float64
}

func (_afg Point) String() string {
	return _b.Sprintf("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029", _afg.X, _afg.Y)
}

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator(path Path, creator *_c.ContentCreator) {
	for _cba, _bgd := range path.Points {
		if _cba == 0 {
			creator.Add_m(_bgd.X, _bgd.Y)
		} else {
			creator.Add_l(_bgd.X, _bgd.Y)
		}
	}
}

// Rotate rotates the vector by the specified angle.
func (_cafc Vector) Rotate(phi float64) Vector {
	_fae := _cafc.Magnitude()
	_eb := _cafc.GetPolarAngle()
	return NewVectorPolar(_fae, _eb+phi)
}

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_efa PolyBezierCurve) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	if _efa.BorderColor == nil {
		_efa.BorderColor = _f.NewPdfColorDeviceRGB(0, 0, 0)
	}
	_cc := NewCubicBezierPath()
	for _, _bfc := range _efa.Curves {
		_cc = _cc.AppendCurve(_bfc)
	}
	_bd := _c.NewContentCreator()
	_bd.Add_q()
	_efa.FillEnabled = _efa.FillEnabled && _efa.FillColor != nil
	if _efa.FillEnabled {
		_bd.SetNonStrokingColor(_efa.FillColor)
	}
	_bd.SetStrokingColor(_efa.BorderColor)
	_bd.Add_w(_efa.BorderWidth)
	if len(gsName) > 1 {
		_bd.Add_gs(_gb.PdfObjectName(gsName))
	}
	for _da, _ca := range _cc.Curves {
		if _da == 0 {
			_bd.Add_m(_ca.P0.X, _ca.P0.Y)
		} else {
			_bd.Add_l(_ca.P0.X, _ca.P0.Y)
		}
		_bd.Add_c(_ca.P1.X, _ca.P1.Y, _ca.P2.X, _ca.P2.Y, _ca.P3.X, _ca.P3.Y)
	}
	if _efa.FillEnabled {
		_bd.Add_h()
		_bd.Add_B()
	} else {
		_bd.Add_S()
	}
	_bd.Add_Q()
	return _bd.Bytes(), _cc.GetBoundingBox().ToPdfRectangle(), nil
}

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_fbc Rectangle) ToPdfRectangle() *_f.PdfRectangle {
	return &_f.PdfRectangle{Llx: _fbc.X, Lly: _fbc.Y, Urx: _fbc.X + _fbc.Width, Ury: _fbc.Y + _fbc.Height}
}

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar(length float64, theta float64) Vector {
	_dcc := Vector{}
	_dcc.Dx = length * _a.Cos(theta)
	_dcc.Dy = length * _a.Sin(theta)
	return _dcc
}

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_fb Point) Rotate(theta float64) Point {
	_ggf := _ad.NewPoint(_fb.X, _fb.Y).Rotate(theta)
	return NewPoint(_ggf.X, _ggf.Y)
}

const (
	LineStyleSolid  LineStyle = 0
	LineStyleDashed LineStyle = 1
)

// Point represents a two-dimensional point.
type Point struct {
	X float64
	Y float64
}

// Draw draws the rectangle. Can specify a graphics state (gsName) for setting opacity etc.
// Otherwise leave empty (""). Returns the content stream as a byte array, bounding box and an error on failure.
func (_ffb Rectangle) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	_eee := NewPath()
	_eee = _eee.AppendPoint(NewPoint(0, 0))
	_eee = _eee.AppendPoint(NewPoint(0, _ffb.Height))
	_eee = _eee.AppendPoint(NewPoint(_ffb.Width, _ffb.Height))
	_eee = _eee.AppendPoint(NewPoint(_ffb.Width, 0))
	_eee = _eee.AppendPoint(NewPoint(0, 0))
	if _ffb.X != 0 || _ffb.Y != 0 {
		_eee = _eee.Offset(_ffb.X, _ffb.Y)
	}
	_ggd := _c.NewContentCreator()
	_ggd.Add_q()
	if _ffb.FillEnabled {
		_ggd.SetNonStrokingColor(_ffb.FillColor)
	}
	if _ffb.BorderEnabled {
		_ggd.SetStrokingColor(_ffb.BorderColor)
		_ggd.Add_w(_ffb.BorderWidth)
	}
	if len(gsName) > 1 {
		_ggd.Add_gs(_gb.PdfObjectName(gsName))
	}
	DrawPathWithCreator(_eee, _ggd)
	_ggd.Add_h()
	if _ffb.FillEnabled && _ffb.BorderEnabled {
		_ggd.Add_B()
	} else if _ffb.FillEnabled {
		_ggd.Add_f()
	} else if _ffb.BorderEnabled {
		_ggd.Add_S()
	}
	_ggd.Add_Q()
	return _ggd.Bytes(), _eee.GetBoundingBox().ToPdfRectangle(), nil
}

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_ddf Polygon) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	_fbg := _c.NewContentCreator()
	_fbg.Add_q()
	_ddf.FillEnabled = _ddf.FillEnabled && _ddf.FillColor != nil
	if _ddf.FillEnabled {
		_fbg.SetNonStrokingColor(_ddf.FillColor)
	}
	_ddf.BorderEnabled = _ddf.BorderEnabled && _ddf.BorderColor != nil
	if _ddf.BorderEnabled {
		_fbg.SetStrokingColor(_ddf.BorderColor)
		_fbg.Add_w(_ddf.BorderWidth)
	}
	if len(gsName) > 1 {
		_fbg.Add_gs(_gb.PdfObjectName(gsName))
	}
	_cbb := NewPath()
	for _, _feb := range _ddf.Points {
		for _ac, _dg := range _feb {
			_cbb = _cbb.AppendPoint(_dg)
			if _ac == 0 {
				_fbg.Add_m(_dg.X, _dg.Y)
			} else {
				_fbg.Add_l(_dg.X, _dg.Y)
			}
		}
		_fbg.Add_h()
	}
	if _ddf.FillEnabled && _ddf.BorderEnabled {
		_fbg.Add_B()
	} else if _ddf.FillEnabled {
		_fbg.Add_f()
	} else if _ddf.BorderEnabled {
		_fbg.Add_S()
	}
	_fbg.Add_Q()
	return _fbg.Bytes(), _cbb.GetBoundingBox().ToPdfRectangle(), nil
}

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_afc BoundingBox) ToPdfRectangle() *_f.PdfRectangle {
	return &_f.PdfRectangle{Llx: _afc.X, Lly: _afc.Y, Urx: _afc.X + _afc.Width, Ury: _afc.Y + _afc.Height}
}

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve(x0, y0, x1, y1, x2, y2, x3, y3 float64) CubicBezierCurve {
	_bb := CubicBezierCurve{}
	_bb.P0 = NewPoint(x0, y0)
	_bb.P1 = NewPoint(x1, y1)
	_bb.P2 = NewPoint(x2, y2)
	_bb.P3 = NewPoint(x3, y3)
	return _bb
}

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct {
	Curves      []CubicBezierCurve
	BorderWidth float64
	BorderColor _f.PdfColor
	FillEnabled bool
	FillColor   _f.PdfColor
}

// Scale scales the vector by the specified factor.
func (_dbe Vector) Scale(factor float64) Vector {
	_ed := _dbe.Magnitude()
	_feba := _dbe.GetPolarAngle()
	_dbe.Dx = factor * _ed * _a.Cos(_feba)
	_dbe.Dy = factor * _ed * _a.Sin(_feba)
	return _dbe
}

// Flip changes the sign of the vector: -vector.
func (_bec Vector) Flip() Vector {
	_cad := _bec.Magnitude()
	_gdd := _bec.GetPolarAngle()
	_bec.Dx = _cad * _a.Cos(_gdd+_a.Pi)
	_bec.Dy = _cad * _a.Sin(_gdd+_a.Pi)
	return _bec
}

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_ae Circle) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	_ee := _ae.Width / 2
	_ege := _ae.Height / 2
	if _ae.BorderEnabled {
		_ee -= _ae.BorderWidth / 2
		_ege -= _ae.BorderWidth / 2
	}
	_bcg := 0.551784
	_gbc := _ee * _bcg
	_cg := _ege * _bcg
	_ef := NewCubicBezierPath()
	_ef = _ef.AppendCurve(NewCubicBezierCurve(-_ee, 0, -_ee, _cg, -_gbc, _ege, 0, _ege))
	_ef = _ef.AppendCurve(NewCubicBezierCurve(0, _ege, _gbc, _ege, _ee, _cg, _ee, 0))
	_ef = _ef.AppendCurve(NewCubicBezierCurve(_ee, 0, _ee, -_cg, _gbc, -_ege, 0, -_ege))
	_ef = _ef.AppendCurve(NewCubicBezierCurve(0, -_ege, -_gbc, -_ege, -_ee, -_cg, -_ee, 0))
	_ef = _ef.Offset(_ee, _ege)
	if _ae.BorderEnabled {
		_ef = _ef.Offset(_ae.BorderWidth/2, _ae.BorderWidth/2)
	}
	if _ae.X != 0 || _ae.Y != 0 {
		_ef = _ef.Offset(_ae.X, _ae.Y)
	}
	_db := _c.NewContentCreator()
	_db.Add_q()
	if _ae.FillEnabled {
		_db.SetNonStrokingColor(_ae.FillColor)
	}
	if _ae.BorderEnabled {
		_db.SetStrokingColor(_ae.BorderColor)
		_db.Add_w(_ae.BorderWidth)
	}
	if len(gsName) > 1 {
		_db.Add_gs(_gb.PdfObjectName(gsName))
	}
	DrawBezierPathWithCreator(_ef, _db)
	_db.Add_h()
	if _ae.FillEnabled && _ae.BorderEnabled {
		_db.Add_B()
	} else if _ae.FillEnabled {
		_db.Add_f()
	} else if _ae.BorderEnabled {
		_db.Add_S()
	}
	_db.Add_Q()
	_bcd := _ef.GetBoundingBox()
	if _ae.BorderEnabled {
		_bcd.Height += _ae.BorderWidth
		_bcd.Width += _ae.BorderWidth
		_bcd.X -= _ae.BorderWidth / 2
		_bcd.Y -= _ae.BorderWidth / 2
	}
	return _db.Bytes(), _bcd.ToPdfRectangle(), nil
}

// AppendCurve appends the specified Bezier curve to the path.
func (_dee CubicBezierPath) AppendCurve(curve CubicBezierCurve) CubicBezierPath {
	_dee.Curves = append(_dee.Curves, curve)
	return _dee
}

// Length returns the number of points in the path.
func (_ba Path) Length() int { return len(_ba.Points) }

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Offset shifts the path with the specified offsets.
func (_fd Path) Offset(offX, offY float64) Path {
	for _fdb, _fa := range _fd.Points {
		_fd.Points[_fdb] = _fa.Add(offX, offY)
	}
	return _fd
}

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     _f.PdfColor
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   _f.PdfColor
	Opacity       float64
}

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_e CubicBezierCurve) AddOffsetXY(offX, offY float64) CubicBezierCurve {
	_e.P0.X += offX
	_e.P1.X += offX
	_e.P2.X += offX
	_e.P3.X += offX
	_e.P0.Y += offY
	_e.P1.Y += offY
	_e.P2.Y += offY
	_e.P3.Y += offY
	return _e
}

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_df CurvePolygon) Draw(gsName string) ([]byte, *_f.PdfRectangle, error) {
	_dgc := _c.NewContentCreator()
	_dgc.Add_q()
	_df.FillEnabled = _df.FillEnabled && _df.FillColor != nil
	if _df.FillEnabled {
		_dgc.SetNonStrokingColor(_df.FillColor)
	}
	_df.BorderEnabled = _df.BorderEnabled && _df.BorderColor != nil
	if _df.BorderEnabled {
		_dgc.SetStrokingColor(_df.BorderColor)
		_dgc.Add_w(_df.BorderWidth)
	}
	if len(gsName) > 1 {
		_dgc.Add_gs(_gb.PdfObjectName(gsName))
	}
	_bfaf := NewCubicBezierPath()
	for _, _eea := range _df.Rings {
		for _gbe, _acg := range _eea {
			if _gbe == 0 {
				_dgc.Add_m(_acg.P0.X, _acg.P0.Y)
			} else {
				_dgc.Add_l(_acg.P0.X, _acg.P0.Y)
			}
			_dgc.Add_c(_acg.P1.X, _acg.P1.Y, _acg.P2.X, _acg.P2.Y, _acg.P3.X, _acg.P3.Y)
			_bfaf = _bfaf.AppendCurve(_acg)
		}
		_dgc.Add_h()
	}
	if _df.FillEnabled && _df.BorderEnabled {
		_dgc.Add_B()
	} else if _df.FillEnabled {
		_dgc.Add_f()
	} else if _df.BorderEnabled {
		_dgc.Add_S()
	}
	_dgc.Add_Q()
	return _dgc.Bytes(), _bfaf.GetBoundingBox().ToPdfRectangle(), nil
}

// GetBoundingBox returns the bounding box of the Bezier path.
func (_cb CubicBezierPath) GetBoundingBox() Rectangle {
	_bga := Rectangle{}
	_dd := 0.0
	_af := 0.0
	_ggg := 0.0
	_eg := 0.0
	for _agb, _deef := range _cb.Curves {
		_bfa := _deef.GetBounds()
		if _agb == 0 {
			_dd = _bfa.Llx
			_af = _bfa.Urx
			_ggg = _bfa.Lly
			_eg = _bfa.Ury
			continue
		}
		if _bfa.Llx < _dd {
			_dd = _bfa.Llx
		}
		if _bfa.Urx > _af {
			_af = _bfa.Urx
		}
		if _bfa.Lly < _ggg {
			_ggg = _bfa.Lly
		}
		if _bfa.Ury > _eg {
			_eg = _bfa.Ury
		}
	}
	_bga.X = _dd
	_bga.Y = _ggg
	_bga.Width = _af - _dd
	_bga.Height = _eg - _ggg
	return _bga
}

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct {
	Rings         [][]CubicBezierCurve
	FillEnabled   bool
	FillColor     _f.PdfColor
	BorderEnabled bool
	BorderColor   _f.PdfColor
	BorderWidth   float64
}

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath() CubicBezierPath {
	_bfg := CubicBezierPath{}
	_bfg.Curves = []CubicBezierCurve{}
	return _bfg
}

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween(a Point, b Point) Vector {
	_agge := Vector{}
	_agge.Dx = b.X - a.X
	_agge.Dy = b.Y - a.Y
	return _agge
}

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector(dx, dy float64) Vector { _fdc := Vector{}; _fdc.Dx = dx; _fdc.Dy = dy; return _fdc }

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_ffd Path) GetPointNumber(number int) Point {
	if number < 1 || number > len(_ffd.Points) {
		return Point{}
	}
	return _ffd.Points[number-1]
}

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{ Points []Point }

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct {
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	LineColor _f.PdfColor
	Opacity   float64
	LineWidth float64
	LineStyle LineStyle
}

// Add adds the specified vector to the current one and returns the result.
func (_gcd Vector) Add(other Vector) Vector { _gcd.Dx += other.Dx; _gcd.Dy += other.Dy; return _gcd }

// GetBoundingBox returns the bounding box of the path.
func (_ab Path) GetBoundingBox() BoundingBox {
	_deg := BoundingBox{}
	_def := 0.0
	_agg := 0.0
	_ea := 0.0
	_afd := 0.0
	for _afe, _gag := range _ab.Points {
		if _afe == 0 {
			_def = _gag.X
			_agg = _gag.X
			_ea = _gag.Y
			_afd = _gag.Y
			continue
		}
		if _gag.X < _def {
			_def = _gag.X
		}
		if _gag.X > _agg {
			_agg = _gag.X
		}
		if _gag.Y < _ea {
			_ea = _gag.Y
		}
		if _gag.Y > _afd {
			_afd = _gag.Y
		}
	}
	_deg.X = _def
	_deg.Y = _ea
	_deg.Width = _agg - _def
	_deg.Height = _afd - _ea
	return _deg
}

const (
	LineEndingStyleNone  LineEndingStyle = 0
	LineEndingStyleArrow LineEndingStyle = 1
	LineEndingStyleButt  LineEndingStyle = 2
)

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator(bpath CubicBezierPath, creator *_c.ContentCreator) {
	for _gef, _ccd := range bpath.Curves {
		if _gef == 0 {
			creator.Add_m(_ccd.P0.X, _ccd.P0.Y)
		}
		creator.Add_c(_ccd.P1.X, _ccd.P1.Y, _ccd.P2.X, _ccd.P2.Y, _ccd.P3.X, _ccd.P3.Y)
	}
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

// Copy returns a clone of the Bezier path.
func (_ge CubicBezierPath) Copy() CubicBezierPath {
	_ga := CubicBezierPath{}
	_ga.Curves = append(_ga.Curves, _ge.Curves...)
	return _ga
}

// Vector represents a two-dimensional vector.
type Vector struct {
	Dx float64
	Dy float64
}

// NewPoint returns a new point with the coordinates x, y.
func NewPoint(x, y float64) Point { return Point{X: x, Y: y} }

// AppendPoint adds the specified point to the path.
func (_cbe Path) AppendPoint(point Point) Path { _cbe.Points = append(_cbe.Points, point); return _cbe }
