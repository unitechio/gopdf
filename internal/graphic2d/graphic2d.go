package graphic2d

import (
	_a "image/color"
	_g "math"
)

func QuadraticToCubicBezier(startX, startY, x1, y1, x, y float64) (Point, Point) {
	_abdg := Point{X: startX, Y: startY}
	_ded := Point{X: x1, Y: y1}
	_ed := Point{X: x, Y: y}
	_fc := _abdg.Interpolate(_ded, 2.0/3.0)
	_aee := _ed.Interpolate(_ded, 2.0/3.0)
	return _fc, _aee
}

type Point struct{ X, Y float64 }

var Names = []string{"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e", "\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065", "\u0061\u0071\u0075\u0061", "\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065", "\u0061\u007a\u0075r\u0065", "\u0062\u0065\u0069g\u0065", "\u0062\u0069\u0073\u0071\u0075\u0065", "\u0062\u006c\u0061c\u006b", "\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064", "\u0062\u006c\u0075\u0065", "\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074", "\u0062\u0072\u006fw\u006e", "\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd", "\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e", "\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065", "\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e", "\u0063\u006f\u0072a\u006c", "\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065", "\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b", "\u0063r\u0069\u006d\u0073\u006f\u006e", "\u0063\u0079\u0061\u006e", "\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065", "\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e", "\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064", "\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079", "\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n", "\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079", "\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi", "d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061", "\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e", "\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065", "\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064", "\u0064a\u0072\u006b\u0072\u0065\u0064", "\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e", "\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079", "\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074", "\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b", "d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065", "\u0064i\u006d\u0067\u0072\u0061\u0079", "\u0064i\u006d\u0067\u0072\u0065\u0079", "\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065", "\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k", "f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065", "f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e", "\u0066u\u0063\u0068\u0073\u0069\u0061", "\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o", "\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065", "\u0067\u006f\u006c\u0064", "\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd", "\u0067\u0072\u0061\u0079", "\u0067\u0072\u0065e\u006e", "g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077", "\u0067\u0072\u0065\u0079", "\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077", "\u0068o\u0074\u0070\u0069\u006e\u006b", "\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d", "\u0069\u006e\u0064\u0069\u0067\u006f", "\u0069\u0076\u006fr\u0079", "\u006b\u0068\u0061k\u0069", "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072", "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068", "\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n", "\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e", "\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e", "\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c", "\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n", "l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077", "\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y", "\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e", "\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y", "\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek", "l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e", "\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065", "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079", "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079", "\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065", "l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077", "\u006c\u0069\u006d\u0065", "\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n", "\u006c\u0069\u006ee\u006e", "\u006da\u0067\u0065\u006e\u0074\u0061", "\u006d\u0061\u0072\u006f\u006f\u006e", "\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065", "\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064", "\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e", "\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e", "\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064", "\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065", "\u006di\u006e\u0074\u0063\u0072\u0065\u0061m", "\u006di\u0073\u0074\u0079\u0072\u006f\u0073e", "\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e", "n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065", "\u006e\u0061\u0076\u0079", "\u006fl\u0064\u006c\u0061\u0063\u0065", "\u006f\u006c\u0069v\u0065", "\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b", "\u006f\u0072\u0061\u006e\u0067\u0065", "\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d", "\u006f\u0072\u0063\u0068\u0069\u0064", "\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064", "\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n", "\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064", "\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070", "\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f", "\u0070\u0065\u0072\u0075", "\u0070\u0069\u006e\u006b", "\u0070\u006c\u0075\u006d", "\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065", "\u0070\u0075\u0072\u0070\u006c\u0065", "\u0072\u0065\u0064", "\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n", "\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e", "s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e", "\u0073\u0061\u006c\u006d\u006f\u006e", "\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e", "\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c", "\u0073\u0069\u0065\u006e\u006e\u0061", "\u0073\u0069\u006c\u0076\u0065\u0072", "\u0073k\u0079\u0062\u006c\u0075\u0065", "\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e", "\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y", "\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y", "\u0073\u006e\u006f\u0077", "s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e", "\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e", "\u0074\u0061\u006e", "\u0074\u0065\u0061\u006c", "\u0074h\u0069\u0073\u0074\u006c\u0065", "\u0074\u006f\u006d\u0061\u0074\u006f", "\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e", "\u0076\u0069\u006f\u006c\u0065\u0074", "\u0077\u0068\u0065a\u0074", "\u0077\u0068\u0069t\u0065", "\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065", "\u0079\u0065\u006c\u006c\u006f\u0077", "y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e"}

func (_df Point) Sub(q Point) Point { return Point{_df.X - q.X, _df.Y - q.Y} }
func _def(_dg, _gd, _bf, _eca, _fa, _afc float64) (float64, float64) {
	_defe, _cfa := _g.Sincos(_afc)
	_dbg, _dbe := _g.Sincos(_bf)
	_be := _eca + _dg*_cfa*_dbe - _gd*_defe*_dbg
	_ca := _fa + _dg*_cfa*_dbg + _gd*_defe*_dbe
	return _be, _ca
}
func _abb(_abc float64) float64 {
	_abc = _g.Mod(_abc, 2.0*_g.Pi)
	if _abc < 0.0 {
		_abc += 2.0 * _g.Pi
	}
	return _abc
}
func (_egd Point) Add(q Point) Point { return Point{_egd.X + q.X, _egd.Y + q.Y} }

var (
	Aliceblue            = _a.RGBA{0xf0, 0xf8, 0xff, 0xff}
	Antiquewhite         = _a.RGBA{0xfa, 0xeb, 0xd7, 0xff}
	Aqua                 = _a.RGBA{0x00, 0xff, 0xff, 0xff}
	Aquamarine           = _a.RGBA{0x7f, 0xff, 0xd4, 0xff}
	Azure                = _a.RGBA{0xf0, 0xff, 0xff, 0xff}
	Beige                = _a.RGBA{0xf5, 0xf5, 0xdc, 0xff}
	Bisque               = _a.RGBA{0xff, 0xe4, 0xc4, 0xff}
	Black                = _a.RGBA{0x00, 0x00, 0x00, 0xff}
	Blanchedalmond       = _a.RGBA{0xff, 0xeb, 0xcd, 0xff}
	Blue                 = _a.RGBA{0x00, 0x00, 0xff, 0xff}
	Blueviolet           = _a.RGBA{0x8a, 0x2b, 0xe2, 0xff}
	Brown                = _a.RGBA{0xa5, 0x2a, 0x2a, 0xff}
	Burlywood            = _a.RGBA{0xde, 0xb8, 0x87, 0xff}
	Cadetblue            = _a.RGBA{0x5f, 0x9e, 0xa0, 0xff}
	Chartreuse           = _a.RGBA{0x7f, 0xff, 0x00, 0xff}
	Chocolate            = _a.RGBA{0xd2, 0x69, 0x1e, 0xff}
	Coral                = _a.RGBA{0xff, 0x7f, 0x50, 0xff}
	Cornflowerblue       = _a.RGBA{0x64, 0x95, 0xed, 0xff}
	Cornsilk             = _a.RGBA{0xff, 0xf8, 0xdc, 0xff}
	Crimson              = _a.RGBA{0xdc, 0x14, 0x3c, 0xff}
	Cyan                 = _a.RGBA{0x00, 0xff, 0xff, 0xff}
	Darkblue             = _a.RGBA{0x00, 0x00, 0x8b, 0xff}
	Darkcyan             = _a.RGBA{0x00, 0x8b, 0x8b, 0xff}
	Darkgoldenrod        = _a.RGBA{0xb8, 0x86, 0x0b, 0xff}
	Darkgray             = _a.RGBA{0xa9, 0xa9, 0xa9, 0xff}
	Darkgreen            = _a.RGBA{0x00, 0x64, 0x00, 0xff}
	Darkgrey             = _a.RGBA{0xa9, 0xa9, 0xa9, 0xff}
	Darkkhaki            = _a.RGBA{0xbd, 0xb7, 0x6b, 0xff}
	Darkmagenta          = _a.RGBA{0x8b, 0x00, 0x8b, 0xff}
	Darkolivegreen       = _a.RGBA{0x55, 0x6b, 0x2f, 0xff}
	Darkorange           = _a.RGBA{0xff, 0x8c, 0x00, 0xff}
	Darkorchid           = _a.RGBA{0x99, 0x32, 0xcc, 0xff}
	Darkred              = _a.RGBA{0x8b, 0x00, 0x00, 0xff}
	Darksalmon           = _a.RGBA{0xe9, 0x96, 0x7a, 0xff}
	Darkseagreen         = _a.RGBA{0x8f, 0xbc, 0x8f, 0xff}
	Darkslateblue        = _a.RGBA{0x48, 0x3d, 0x8b, 0xff}
	Darkslategray        = _a.RGBA{0x2f, 0x4f, 0x4f, 0xff}
	Darkslategrey        = _a.RGBA{0x2f, 0x4f, 0x4f, 0xff}
	Darkturquoise        = _a.RGBA{0x00, 0xce, 0xd1, 0xff}
	Darkviolet           = _a.RGBA{0x94, 0x00, 0xd3, 0xff}
	Deeppink             = _a.RGBA{0xff, 0x14, 0x93, 0xff}
	Deepskyblue          = _a.RGBA{0x00, 0xbf, 0xff, 0xff}
	Dimgray              = _a.RGBA{0x69, 0x69, 0x69, 0xff}
	Dimgrey              = _a.RGBA{0x69, 0x69, 0x69, 0xff}
	Dodgerblue           = _a.RGBA{0x1e, 0x90, 0xff, 0xff}
	Firebrick            = _a.RGBA{0xb2, 0x22, 0x22, 0xff}
	Floralwhite          = _a.RGBA{0xff, 0xfa, 0xf0, 0xff}
	Forestgreen          = _a.RGBA{0x22, 0x8b, 0x22, 0xff}
	Fuchsia              = _a.RGBA{0xff, 0x00, 0xff, 0xff}
	Gainsboro            = _a.RGBA{0xdc, 0xdc, 0xdc, 0xff}
	Ghostwhite           = _a.RGBA{0xf8, 0xf8, 0xff, 0xff}
	Gold                 = _a.RGBA{0xff, 0xd7, 0x00, 0xff}
	Goldenrod            = _a.RGBA{0xda, 0xa5, 0x20, 0xff}
	Gray                 = _a.RGBA{0x80, 0x80, 0x80, 0xff}
	Green                = _a.RGBA{0x00, 0x80, 0x00, 0xff}
	Greenyellow          = _a.RGBA{0xad, 0xff, 0x2f, 0xff}
	Grey                 = _a.RGBA{0x80, 0x80, 0x80, 0xff}
	Honeydew             = _a.RGBA{0xf0, 0xff, 0xf0, 0xff}
	Hotpink              = _a.RGBA{0xff, 0x69, 0xb4, 0xff}
	Indianred            = _a.RGBA{0xcd, 0x5c, 0x5c, 0xff}
	Indigo               = _a.RGBA{0x4b, 0x00, 0x82, 0xff}
	Ivory                = _a.RGBA{0xff, 0xff, 0xf0, 0xff}
	Khaki                = _a.RGBA{0xf0, 0xe6, 0x8c, 0xff}
	Lavender             = _a.RGBA{0xe6, 0xe6, 0xfa, 0xff}
	Lavenderblush        = _a.RGBA{0xff, 0xf0, 0xf5, 0xff}
	Lawngreen            = _a.RGBA{0x7c, 0xfc, 0x00, 0xff}
	Lemonchiffon         = _a.RGBA{0xff, 0xfa, 0xcd, 0xff}
	Lightblue            = _a.RGBA{0xad, 0xd8, 0xe6, 0xff}
	Lightcoral           = _a.RGBA{0xf0, 0x80, 0x80, 0xff}
	Lightcyan            = _a.RGBA{0xe0, 0xff, 0xff, 0xff}
	Lightgoldenrodyellow = _a.RGBA{0xfa, 0xfa, 0xd2, 0xff}
	Lightgray            = _a.RGBA{0xd3, 0xd3, 0xd3, 0xff}
	Lightgreen           = _a.RGBA{0x90, 0xee, 0x90, 0xff}
	Lightgrey            = _a.RGBA{0xd3, 0xd3, 0xd3, 0xff}
	Lightpink            = _a.RGBA{0xff, 0xb6, 0xc1, 0xff}
	Lightsalmon          = _a.RGBA{0xff, 0xa0, 0x7a, 0xff}
	Lightseagreen        = _a.RGBA{0x20, 0xb2, 0xaa, 0xff}
	Lightskyblue         = _a.RGBA{0x87, 0xce, 0xfa, 0xff}
	Lightslategray       = _a.RGBA{0x77, 0x88, 0x99, 0xff}
	Lightslategrey       = _a.RGBA{0x77, 0x88, 0x99, 0xff}
	Lightsteelblue       = _a.RGBA{0xb0, 0xc4, 0xde, 0xff}
	Lightyellow          = _a.RGBA{0xff, 0xff, 0xe0, 0xff}
	Lime                 = _a.RGBA{0x00, 0xff, 0x00, 0xff}
	Limegreen            = _a.RGBA{0x32, 0xcd, 0x32, 0xff}
	Linen                = _a.RGBA{0xfa, 0xf0, 0xe6, 0xff}
	Magenta              = _a.RGBA{0xff, 0x00, 0xff, 0xff}
	Maroon               = _a.RGBA{0x80, 0x00, 0x00, 0xff}
	Mediumaquamarine     = _a.RGBA{0x66, 0xcd, 0xaa, 0xff}
	Mediumblue           = _a.RGBA{0x00, 0x00, 0xcd, 0xff}
	Mediumorchid         = _a.RGBA{0xba, 0x55, 0xd3, 0xff}
	Mediumpurple         = _a.RGBA{0x93, 0x70, 0xdb, 0xff}
	Mediumseagreen       = _a.RGBA{0x3c, 0xb3, 0x71, 0xff}
	Mediumslateblue      = _a.RGBA{0x7b, 0x68, 0xee, 0xff}
	Mediumspringgreen    = _a.RGBA{0x00, 0xfa, 0x9a, 0xff}
	Mediumturquoise      = _a.RGBA{0x48, 0xd1, 0xcc, 0xff}
	Mediumvioletred      = _a.RGBA{0xc7, 0x15, 0x85, 0xff}
	Midnightblue         = _a.RGBA{0x19, 0x19, 0x70, 0xff}
	Mintcream            = _a.RGBA{0xf5, 0xff, 0xfa, 0xff}
	Mistyrose            = _a.RGBA{0xff, 0xe4, 0xe1, 0xff}
	Moccasin             = _a.RGBA{0xff, 0xe4, 0xb5, 0xff}
	Navajowhite          = _a.RGBA{0xff, 0xde, 0xad, 0xff}
	Navy                 = _a.RGBA{0x00, 0x00, 0x80, 0xff}
	Oldlace              = _a.RGBA{0xfd, 0xf5, 0xe6, 0xff}
	Olive                = _a.RGBA{0x80, 0x80, 0x00, 0xff}
	Olivedrab            = _a.RGBA{0x6b, 0x8e, 0x23, 0xff}
	Orange               = _a.RGBA{0xff, 0xa5, 0x00, 0xff}
	Orangered            = _a.RGBA{0xff, 0x45, 0x00, 0xff}
	Orchid               = _a.RGBA{0xda, 0x70, 0xd6, 0xff}
	Palegoldenrod        = _a.RGBA{0xee, 0xe8, 0xaa, 0xff}
	Palegreen            = _a.RGBA{0x98, 0xfb, 0x98, 0xff}
	Paleturquoise        = _a.RGBA{0xaf, 0xee, 0xee, 0xff}
	Palevioletred        = _a.RGBA{0xdb, 0x70, 0x93, 0xff}
	Papayawhip           = _a.RGBA{0xff, 0xef, 0xd5, 0xff}
	Peachpuff            = _a.RGBA{0xff, 0xda, 0xb9, 0xff}
	Peru                 = _a.RGBA{0xcd, 0x85, 0x3f, 0xff}
	Pink                 = _a.RGBA{0xff, 0xc0, 0xcb, 0xff}
	Plum                 = _a.RGBA{0xdd, 0xa0, 0xdd, 0xff}
	Powderblue           = _a.RGBA{0xb0, 0xe0, 0xe6, 0xff}
	Purple               = _a.RGBA{0x80, 0x00, 0x80, 0xff}
	Red                  = _a.RGBA{0xff, 0x00, 0x00, 0xff}
	Rosybrown            = _a.RGBA{0xbc, 0x8f, 0x8f, 0xff}
	Royalblue            = _a.RGBA{0x41, 0x69, 0xe1, 0xff}
	Saddlebrown          = _a.RGBA{0x8b, 0x45, 0x13, 0xff}
	Salmon               = _a.RGBA{0xfa, 0x80, 0x72, 0xff}
	Sandybrown           = _a.RGBA{0xf4, 0xa4, 0x60, 0xff}
	Seagreen             = _a.RGBA{0x2e, 0x8b, 0x57, 0xff}
	Seashell             = _a.RGBA{0xff, 0xf5, 0xee, 0xff}
	Sienna               = _a.RGBA{0xa0, 0x52, 0x2d, 0xff}
	Silver               = _a.RGBA{0xc0, 0xc0, 0xc0, 0xff}
	Skyblue              = _a.RGBA{0x87, 0xce, 0xeb, 0xff}
	Slateblue            = _a.RGBA{0x6a, 0x5a, 0xcd, 0xff}
	Slategray            = _a.RGBA{0x70, 0x80, 0x90, 0xff}
	Slategrey            = _a.RGBA{0x70, 0x80, 0x90, 0xff}
	Snow                 = _a.RGBA{0xff, 0xfa, 0xfa, 0xff}
	Springgreen          = _a.RGBA{0x00, 0xff, 0x7f, 0xff}
	Steelblue            = _a.RGBA{0x46, 0x82, 0xb4, 0xff}
	Tan                  = _a.RGBA{0xd2, 0xb4, 0x8c, 0xff}
	Teal                 = _a.RGBA{0x00, 0x80, 0x80, 0xff}
	Thistle              = _a.RGBA{0xd8, 0xbf, 0xd8, 0xff}
	Tomato               = _a.RGBA{0xff, 0x63, 0x47, 0xff}
	Turquoise            = _a.RGBA{0x40, 0xe0, 0xd0, 0xff}
	Violet               = _a.RGBA{0xee, 0x82, 0xee, 0xff}
	Wheat                = _a.RGBA{0xf5, 0xde, 0xb3, 0xff}
	White                = _a.RGBA{0xff, 0xff, 0xff, 0xff}
	Whitesmoke           = _a.RGBA{0xf5, 0xf5, 0xf5, 0xff}
	Yellow               = _a.RGBA{0xff, 0xff, 0x00, 0xff}
	Yellowgreen          = _a.RGBA{0x9a, 0xcd, 0x32, 0xff}
)

func (_cea Point) Interpolate(q Point, t float64) Point {
	return Point{(1-t)*_cea.X + t*q.X, (1-t)*_cea.Y + t*q.Y}
}
func EllipseToCubicBeziers(startX, startY, rx, ry, rot float64, large, sweep bool, endX, endY float64) [][4]Point {
	rx = _g.Abs(rx)
	ry = _g.Abs(ry)
	if rx < ry {
		rx, ry = ry, rx
		rot += 90.0
	}
	_af := _abb(rot * _g.Pi / 180.0)
	if _g.Pi <= _af {
		_af -= _g.Pi
	}
	_ba, _ab, _d, _f := _bad(startX, startY, rx, ry, _af, large, sweep, endX, endY)
	_ag := _g.Pi / 2.0
	_fe := int(_g.Ceil(_g.Abs(_f-_d) / _ag))
	_ag = _g.Abs(_f-_d) / float64(_fe)
	_e := _g.Sin(_ag) * (_g.Sqrt(4.0+3.0*_g.Pow(_g.Tan(_ag/2.0), 2.0)) - 1.0) / 3.0
	if !sweep {
		_ag = -_ag
	}
	_gg := Point{X: startX, Y: startY}
	_c, _ec := _acf(rx, ry, _af, sweep, _d)
	_cd := Point{X: _c, Y: _ec}
	_eg := [][4]Point{}
	for _ff := 1; _ff < _fe+1; _ff++ {
		_cb := _d + float64(_ff)*_ag
		_cf, _age := _def(rx, ry, _af, _ba, _ab, _cb)
		_ece := Point{X: _cf, Y: _age}
		_db, _fed := _acf(rx, ry, _af, sweep, _cb)
		_cdg := Point{X: _db, Y: _fed}
		_abd := _gg.Add(_cd.Mul(_e))
		_de := _ece.Sub(_cdg.Mul(_e))
		_eg = append(_eg, [4]Point{_gg, _abd, _de, _ece})
		_cd = _cdg
		_gg = _ece
	}
	return _eg
}
func (_bcb Point) Mul(f float64) Point { return Point{f * _bcb.X, f * _bcb.Y} }
func _acf(_cc, _affb, _agd float64, _dc bool, _ceb float64) (float64, float64) {
	_ae, _cfab := _g.Sincos(_ceb)
	_gdd, _afg := _g.Sincos(_agd)
	_fae := -_cc*_ae*_afg - _affb*_cfab*_gdd
	_egb := -_cc*_ae*_gdd + _affb*_cfab*_afg
	if !_dc {
		return -_fae, -_egb
	}
	return _fae, _egb
}
func _bad(_aff, _cda, _da, _ce, _egg float64, _dd, _bfc bool, _bg, _ffb float64) (float64, float64, float64, float64) {
	if _bd(_aff, _bg) && _bd(_cda, _ffb) {
		return _aff, _cda, 0.0, 0.0
	}
	_bgb, _ea := _g.Sincos(_egg)
	_dgg := _ea*(_aff-_bg)/2.0 + _bgb*(_cda-_ffb)/2.0
	_eae := -_bgb*(_aff-_bg)/2.0 + _ea*(_cda-_ffb)/2.0
	_dgd := _dgg*_dgg/_da/_da + _eae*_eae/_ce/_ce
	if _dgd > 1.0 {
		_da *= _g.Sqrt(_dgd)
		_ce *= _g.Sqrt(_dgd)
	}
	_cba := (_da*_da*_ce*_ce - _da*_da*_eae*_eae - _ce*_ce*_dgg*_dgg) / (_da*_da*_eae*_eae + _ce*_ce*_dgg*_dgg)
	if _cba < 0.0 {
		_cba = 0.0
	}
	_fee := _g.Sqrt(_cba)
	if _dd == _bfc {
		_fee = -_fee
	}
	_ac := _fee * _da * _eae / _ce
	_dee := _fee * -_ce * _dgg / _da
	_bc := _ea*_ac - _bgb*_dee + (_aff+_bg)/2.0
	_dec := _bgb*_ac + _ea*_dee + (_cda+_ffb)/2.0
	_dad := (_dgg - _ac) / _da
	_dgf := (_eae - _dee) / _ce
	_cbb := -(_dgg + _ac) / _da
	_baf := -(_eae + _dee) / _ce
	_acd := _g.Acos(_dad / _g.Sqrt(_dad*_dad+_dgf*_dgf))
	if _dgf < 0.0 {
		_acd = -_acd
	}
	_acd = _abb(_acd)
	_eaf := (_dad*_cbb + _dgf*_baf) / _g.Sqrt((_dad*_dad+_dgf*_dgf)*(_cbb*_cbb+_baf*_baf))
	_eaf = _g.Min(1.0, _g.Max(-1.0, _eaf))
	_ggf := _g.Acos(_eaf)
	if _dad*_baf-_dgf*_cbb < 0.0 {
		_ggf = -_ggf
	}
	if !_bfc && _ggf > 0.0 {
		_ggf -= 2.0 * _g.Pi
	} else if _bfc && _ggf < 0.0 {
		_ggf += 2.0 * _g.Pi
	}
	return _bc, _dec, _acd, _acd + _ggf
}

const _gb = 1e-10

func _bd(_ffc, _gf float64) bool { return _g.Abs(_ffc-_gf) <= _gb }

var ColorMap = map[string]_a.RGBA{"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e": _a.RGBA{0xf0, 0xf8, 0xff, 0xff}, "\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065": _a.RGBA{0xfa, 0xeb, 0xd7, 0xff}, "\u0061\u0071\u0075\u0061": _a.RGBA{0x00, 0xff, 0xff, 0xff}, "\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065": _a.RGBA{0x7f, 0xff, 0xd4, 0xff}, "\u0061\u007a\u0075r\u0065": _a.RGBA{0xf0, 0xff, 0xff, 0xff}, "\u0062\u0065\u0069g\u0065": _a.RGBA{0xf5, 0xf5, 0xdc, 0xff}, "\u0062\u0069\u0073\u0071\u0075\u0065": _a.RGBA{0xff, 0xe4, 0xc4, 0xff}, "\u0062\u006c\u0061c\u006b": _a.RGBA{0x00, 0x00, 0x00, 0xff}, "\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064": _a.RGBA{0xff, 0xeb, 0xcd, 0xff}, "\u0062\u006c\u0075\u0065": _a.RGBA{0x00, 0x00, 0xff, 0xff}, "\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074": _a.RGBA{0x8a, 0x2b, 0xe2, 0xff}, "\u0062\u0072\u006fw\u006e": _a.RGBA{0xa5, 0x2a, 0x2a, 0xff}, "\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd": _a.RGBA{0xde, 0xb8, 0x87, 0xff}, "\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e": _a.RGBA{0x5f, 0x9e, 0xa0, 0xff}, "\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065": _a.RGBA{0x7f, 0xff, 0x00, 0xff}, "\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e": _a.RGBA{0xd2, 0x69, 0x1e, 0xff}, "\u0063\u006f\u0072a\u006c": _a.RGBA{0xff, 0x7f, 0x50, 0xff}, "\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065": _a.RGBA{0x64, 0x95, 0xed, 0xff}, "\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b": _a.RGBA{0xff, 0xf8, 0xdc, 0xff}, "\u0063r\u0069\u006d\u0073\u006f\u006e": _a.RGBA{0xdc, 0x14, 0x3c, 0xff}, "\u0063\u0079\u0061\u006e": _a.RGBA{0x00, 0xff, 0xff, 0xff}, "\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065": _a.RGBA{0x00, 0x00, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e": _a.RGBA{0x00, 0x8b, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064": _a.RGBA{0xb8, 0x86, 0x0b, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079": _a.RGBA{0xa9, 0xa9, 0xa9, 0xff}, "\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n": _a.RGBA{0x00, 0x64, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079": _a.RGBA{0xa9, 0xa9, 0xa9, 0xff}, "\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi": _a.RGBA{0xbd, 0xb7, 0x6b, 0xff}, "d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061": _a.RGBA{0x8b, 0x00, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e": _a.RGBA{0x55, 0x6b, 0x2f, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065": _a.RGBA{0xff, 0x8c, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064": _a.RGBA{0x99, 0x32, 0xcc, 0xff}, "\u0064a\u0072\u006b\u0072\u0065\u0064": _a.RGBA{0x8b, 0x00, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e": _a.RGBA{0xe9, 0x96, 0x7a, 0xff}, "\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": _a.RGBA{0x8f, 0xbc, 0x8f, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065": _a.RGBA{0x48, 0x3d, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079": _a.RGBA{0x2f, 0x4f, 0x4f, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079": _a.RGBA{0x2f, 0x4f, 0x4f, 0xff}, "\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": _a.RGBA{0x00, 0xce, 0xd1, 0xff}, "\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074": _a.RGBA{0x94, 0x00, 0xd3, 0xff}, "\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b": _a.RGBA{0xff, 0x14, 0x93, 0xff}, "d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065": _a.RGBA{0x00, 0xbf, 0xff, 0xff}, "\u0064i\u006d\u0067\u0072\u0061\u0079": _a.RGBA{0x69, 0x69, 0x69, 0xff}, "\u0064i\u006d\u0067\u0072\u0065\u0079": _a.RGBA{0x69, 0x69, 0x69, 0xff}, "\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065": _a.RGBA{0x1e, 0x90, 0xff, 0xff}, "\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k": _a.RGBA{0xb2, 0x22, 0x22, 0xff}, "f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065": _a.RGBA{0xff, 0xfa, 0xf0, 0xff}, "f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e": _a.RGBA{0x22, 0x8b, 0x22, 0xff}, "\u0066u\u0063\u0068\u0073\u0069\u0061": _a.RGBA{0xff, 0x00, 0xff, 0xff}, "\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o": _a.RGBA{0xdc, 0xdc, 0xdc, 0xff}, "\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065": _a.RGBA{0xf8, 0xf8, 0xff, 0xff}, "\u0067\u006f\u006c\u0064": _a.RGBA{0xff, 0xd7, 0x00, 0xff}, "\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd": _a.RGBA{0xda, 0xa5, 0x20, 0xff}, "\u0067\u0072\u0061\u0079": _a.RGBA{0x80, 0x80, 0x80, 0xff}, "\u0067\u0072\u0065e\u006e": _a.RGBA{0x00, 0x80, 0x00, 0xff}, "g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077": _a.RGBA{0xad, 0xff, 0x2f, 0xff}, "\u0067\u0072\u0065\u0079": _a.RGBA{0x80, 0x80, 0x80, 0xff}, "\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077": _a.RGBA{0xf0, 0xff, 0xf0, 0xff}, "\u0068o\u0074\u0070\u0069\u006e\u006b": _a.RGBA{0xff, 0x69, 0xb4, 0xff}, "\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d": _a.RGBA{0xcd, 0x5c, 0x5c, 0xff}, "\u0069\u006e\u0064\u0069\u0067\u006f": _a.RGBA{0x4b, 0x00, 0x82, 0xff}, "\u0069\u0076\u006fr\u0079": _a.RGBA{0xff, 0xff, 0xf0, 0xff}, "\u006b\u0068\u0061k\u0069": _a.RGBA{0xf0, 0xe6, 0x8c, 0xff}, "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072": _a.RGBA{0xe6, 0xe6, 0xfa, 0xff}, "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068": _a.RGBA{0xff, 0xf0, 0xf5, 0xff}, "\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n": _a.RGBA{0x7c, 0xfc, 0x00, 0xff}, "\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e": _a.RGBA{0xff, 0xfa, 0xcd, 0xff}, "\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e": _a.RGBA{0xad, 0xd8, 0xe6, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c": _a.RGBA{0xf0, 0x80, 0x80, 0xff}, "\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n": _a.RGBA{0xe0, 0xff, 0xff, 0xff}, "l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077": _a.RGBA{0xfa, 0xfa, 0xd2, 0xff}, "\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y": _a.RGBA{0xd3, 0xd3, 0xd3, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e": _a.RGBA{0x90, 0xee, 0x90, 0xff}, "\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y": _a.RGBA{0xd3, 0xd3, 0xd3, 0xff}, "\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek": _a.RGBA{0xff, 0xb6, 0xc1, 0xff}, "l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e": _a.RGBA{0xff, 0xa0, 0x7a, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": _a.RGBA{0x20, 0xb2, 0xaa, 0xff}, "\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065": _a.RGBA{0x87, 0xce, 0xfa, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079": _a.RGBA{0x77, 0x88, 0x99, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079": _a.RGBA{0x77, 0x88, 0x99, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065": _a.RGBA{0xb0, 0xc4, 0xde, 0xff}, "l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077": _a.RGBA{0xff, 0xff, 0xe0, 0xff}, "\u006c\u0069\u006d\u0065": _a.RGBA{0x00, 0xff, 0x00, 0xff}, "\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n": _a.RGBA{0x32, 0xcd, 0x32, 0xff}, "\u006c\u0069\u006ee\u006e": _a.RGBA{0xfa, 0xf0, 0xe6, 0xff}, "\u006da\u0067\u0065\u006e\u0074\u0061": _a.RGBA{0xff, 0x00, 0xff, 0xff}, "\u006d\u0061\u0072\u006f\u006f\u006e": _a.RGBA{0x80, 0x00, 0x00, 0xff}, "\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065": _a.RGBA{0x66, 0xcd, 0xaa, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065": _a.RGBA{0x00, 0x00, 0xcd, 0xff}, "\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064": _a.RGBA{0xba, 0x55, 0xd3, 0xff}, "\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065": _a.RGBA{0x93, 0x70, 0xdb, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e": _a.RGBA{0x3c, 0xb3, 0x71, 0xff}, "\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065": _a.RGBA{0x7b, 0x68, 0xee, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e": _a.RGBA{0x00, 0xfa, 0x9a, 0xff}, "\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": _a.RGBA{0x48, 0xd1, 0xcc, 0xff}, "\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064": _a.RGBA{0xc7, 0x15, 0x85, 0xff}, "\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065": _a.RGBA{0x19, 0x19, 0x70, 0xff}, "\u006di\u006e\u0074\u0063\u0072\u0065\u0061m": _a.RGBA{0xf5, 0xff, 0xfa, 0xff}, "\u006di\u0073\u0074\u0079\u0072\u006f\u0073e": _a.RGBA{0xff, 0xe4, 0xe1, 0xff}, "\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e": _a.RGBA{0xff, 0xe4, 0xb5, 0xff}, "n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065": _a.RGBA{0xff, 0xde, 0xad, 0xff}, "\u006e\u0061\u0076\u0079": _a.RGBA{0x00, 0x00, 0x80, 0xff}, "\u006fl\u0064\u006c\u0061\u0063\u0065": _a.RGBA{0xfd, 0xf5, 0xe6, 0xff}, "\u006f\u006c\u0069v\u0065": _a.RGBA{0x80, 0x80, 0x00, 0xff}, "\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b": _a.RGBA{0x6b, 0x8e, 0x23, 0xff}, "\u006f\u0072\u0061\u006e\u0067\u0065": _a.RGBA{0xff, 0xa5, 0x00, 0xff}, "\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d": _a.RGBA{0xff, 0x45, 0x00, 0xff}, "\u006f\u0072\u0063\u0068\u0069\u0064": _a.RGBA{0xda, 0x70, 0xd6, 0xff}, "\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064": _a.RGBA{0xee, 0xe8, 0xaa, 0xff}, "\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n": _a.RGBA{0x98, 0xfb, 0x98, 0xff}, "\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": _a.RGBA{0xaf, 0xee, 0xee, 0xff}, "\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064": _a.RGBA{0xdb, 0x70, 0x93, 0xff}, "\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070": _a.RGBA{0xff, 0xef, 0xd5, 0xff}, "\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f": _a.RGBA{0xff, 0xda, 0xb9, 0xff}, "\u0070\u0065\u0072\u0075": _a.RGBA{0xcd, 0x85, 0x3f, 0xff}, "\u0070\u0069\u006e\u006b": _a.RGBA{0xff, 0xc0, 0xcb, 0xff}, "\u0070\u006c\u0075\u006d": _a.RGBA{0xdd, 0xa0, 0xdd, 0xff}, "\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065": _a.RGBA{0xb0, 0xe0, 0xe6, 0xff}, "\u0070\u0075\u0072\u0070\u006c\u0065": _a.RGBA{0x80, 0x00, 0x80, 0xff}, "\u0072\u0065\u0064": _a.RGBA{0xff, 0x00, 0x00, 0xff}, "\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n": _a.RGBA{0xbc, 0x8f, 0x8f, 0xff}, "\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e": _a.RGBA{0x41, 0x69, 0xe1, 0xff}, "s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e": _a.RGBA{0x8b, 0x45, 0x13, 0xff}, "\u0073\u0061\u006c\u006d\u006f\u006e": _a.RGBA{0xfa, 0x80, 0x72, 0xff}, "\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e": _a.RGBA{0xf4, 0xa4, 0x60, 0xff}, "\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": _a.RGBA{0x2e, 0x8b, 0x57, 0xff}, "\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c": _a.RGBA{0xff, 0xf5, 0xee, 0xff}, "\u0073\u0069\u0065\u006e\u006e\u0061": _a.RGBA{0xa0, 0x52, 0x2d, 0xff}, "\u0073\u0069\u006c\u0076\u0065\u0072": _a.RGBA{0xc0, 0xc0, 0xc0, 0xff}, "\u0073k\u0079\u0062\u006c\u0075\u0065": _a.RGBA{0x87, 0xce, 0xeb, 0xff}, "\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e": _a.RGBA{0x6a, 0x5a, 0xcd, 0xff}, "\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y": _a.RGBA{0x70, 0x80, 0x90, 0xff}, "\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y": _a.RGBA{0x70, 0x80, 0x90, 0xff}, "\u0073\u006e\u006f\u0077": _a.RGBA{0xff, 0xfa, 0xfa, 0xff}, "s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e": _a.RGBA{0x00, 0xff, 0x7f, 0xff}, "\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e": _a.RGBA{0x46, 0x82, 0xb4, 0xff}, "\u0074\u0061\u006e": _a.RGBA{0xd2, 0xb4, 0x8c, 0xff}, "\u0074\u0065\u0061\u006c": _a.RGBA{0x00, 0x80, 0x80, 0xff}, "\u0074h\u0069\u0073\u0074\u006c\u0065": _a.RGBA{0xd8, 0xbf, 0xd8, 0xff}, "\u0074\u006f\u006d\u0061\u0074\u006f": _a.RGBA{0xff, 0x63, 0x47, 0xff}, "\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e": _a.RGBA{0x40, 0xe0, 0xd0, 0xff}, "\u0076\u0069\u006f\u006c\u0065\u0074": _a.RGBA{0xee, 0x82, 0xee, 0xff}, "\u0077\u0068\u0065a\u0074": _a.RGBA{0xf5, 0xde, 0xb3, 0xff}, "\u0077\u0068\u0069t\u0065": _a.RGBA{0xff, 0xff, 0xff, 0xff}, "\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065": _a.RGBA{0xf5, 0xf5, 0xf5, 0xff}, "\u0079\u0065\u006c\u006c\u006f\u0077": _a.RGBA{0xff, 0xff, 0x00, 0xff}, "y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e": _a.RGBA{0x9a, 0xcd, 0x32, 0xff}}
