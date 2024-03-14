package graphic2d

import (
	_f "image/color"
	_g "math"
)

func _gd(_bba, _gcf, _ab, _ae, _ga, _ff float64) (float64, float64) {
	_cc, _d := _g.Sincos(_ff)
	_gf, _ffc := _g.Sincos(_ab)
	_ag := _ae + _bba*_d*_ffc - _gcf*_cc*_gf
	_eca := _ga + _bba*_d*_gf + _gcf*_cc*_ffc
	return _ag, _eca
}
func (_cb Point) Sub(q Point) Point { return Point{_cb.X - q.X, _cb.Y - q.Y} }

var ColorMap = map[string]_f.RGBA{"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e": {0xf0, 0xf8, 0xff, 0xff}, "\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065": {0xfa, 0xeb, 0xd7, 0xff}, "\u0061\u0071\u0075\u0061": {0x00, 0xff, 0xff, 0xff}, "\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065": {0x7f, 0xff, 0xd4, 0xff}, "\u0061\u007a\u0075r\u0065": {0xf0, 0xff, 0xff, 0xff}, "\u0062\u0065\u0069g\u0065": {0xf5, 0xf5, 0xdc, 0xff}, "\u0062\u0069\u0073\u0071\u0075\u0065": {0xff, 0xe4, 0xc4, 0xff}, "\u0062\u006c\u0061c\u006b": {0x00, 0x00, 0x00, 0xff}, "\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064": {0xff, 0xeb, 0xcd, 0xff}, "\u0062\u006c\u0075\u0065": {0x00, 0x00, 0xff, 0xff}, "\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074": {0x8a, 0x2b, 0xe2, 0xff}, "\u0062\u0072\u006fw\u006e": {0xa5, 0x2a, 0x2a, 0xff}, "\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd": {0xde, 0xb8, 0x87, 0xff}, "\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e": {0x5f, 0x9e, 0xa0, 0xff}, "\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065": {0x7f, 0xff, 0x00, 0xff}, "\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e": {0xd2, 0x69, 0x1e, 0xff}, "\u0063\u006f\u0072a\u006c": {0xff, 0x7f, 0x50, 0xff}, "\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065": {0x64, 0x95, 0xed, 0xff}, "\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b": {0xff, 0xf8, 0xdc, 0xff}, "\u0063r\u0069\u006d\u0073\u006f\u006e": {0xdc, 0x14, 0x3c, 0xff}, "\u0063\u0079\u0061\u006e": {0x00, 0xff, 0xff, 0xff}, "\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065": {0x00, 0x00, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e": {0x00, 0x8b, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064": {0xb8, 0x86, 0x0b, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079": {0xa9, 0xa9, 0xa9, 0xff}, "\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n": {0x00, 0x64, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079": {0xa9, 0xa9, 0xa9, 0xff}, "\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi": {0xbd, 0xb7, 0x6b, 0xff}, "d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061": {0x8b, 0x00, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e": {0x55, 0x6b, 0x2f, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065": {0xff, 0x8c, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064": {0x99, 0x32, 0xcc, 0xff}, "\u0064a\u0072\u006b\u0072\u0065\u0064": {0x8b, 0x00, 0x00, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e": {0xe9, 0x96, 0x7a, 0xff}, "\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": {0x8f, 0xbc, 0x8f, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065": {0x48, 0x3d, 0x8b, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079": {0x2f, 0x4f, 0x4f, 0xff}, "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079": {0x2f, 0x4f, 0x4f, 0xff}, "\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": {0x00, 0xce, 0xd1, 0xff}, "\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074": {0x94, 0x00, 0xd3, 0xff}, "\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b": {0xff, 0x14, 0x93, 0xff}, "d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065": {0x00, 0xbf, 0xff, 0xff}, "\u0064i\u006d\u0067\u0072\u0061\u0079": {0x69, 0x69, 0x69, 0xff}, "\u0064i\u006d\u0067\u0072\u0065\u0079": {0x69, 0x69, 0x69, 0xff}, "\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065": {0x1e, 0x90, 0xff, 0xff}, "\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k": {0xb2, 0x22, 0x22, 0xff}, "f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065": {0xff, 0xfa, 0xf0, 0xff}, "f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e": {0x22, 0x8b, 0x22, 0xff}, "\u0066u\u0063\u0068\u0073\u0069\u0061": {0xff, 0x00, 0xff, 0xff}, "\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o": {0xdc, 0xdc, 0xdc, 0xff}, "\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065": {0xf8, 0xf8, 0xff, 0xff}, "\u0067\u006f\u006c\u0064": {0xff, 0xd7, 0x00, 0xff}, "\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd": {0xda, 0xa5, 0x20, 0xff}, "\u0067\u0072\u0061\u0079": {0x80, 0x80, 0x80, 0xff}, "\u0067\u0072\u0065e\u006e": {0x00, 0x80, 0x00, 0xff}, "g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077": {0xad, 0xff, 0x2f, 0xff}, "\u0067\u0072\u0065\u0079": {0x80, 0x80, 0x80, 0xff}, "\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077": {0xf0, 0xff, 0xf0, 0xff}, "\u0068o\u0074\u0070\u0069\u006e\u006b": {0xff, 0x69, 0xb4, 0xff}, "\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d": {0xcd, 0x5c, 0x5c, 0xff}, "\u0069\u006e\u0064\u0069\u0067\u006f": {0x4b, 0x00, 0x82, 0xff}, "\u0069\u0076\u006fr\u0079": {0xff, 0xff, 0xf0, 0xff}, "\u006b\u0068\u0061k\u0069": {0xf0, 0xe6, 0x8c, 0xff}, "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072": {0xe6, 0xe6, 0xfa, 0xff}, "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068": {0xff, 0xf0, 0xf5, 0xff}, "\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n": {0x7c, 0xfc, 0x00, 0xff}, "\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e": {0xff, 0xfa, 0xcd, 0xff}, "\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e": {0xad, 0xd8, 0xe6, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c": {0xf0, 0x80, 0x80, 0xff}, "\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n": {0xe0, 0xff, 0xff, 0xff}, "l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077": {0xfa, 0xfa, 0xd2, 0xff}, "\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y": {0xd3, 0xd3, 0xd3, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e": {0x90, 0xee, 0x90, 0xff}, "\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y": {0xd3, 0xd3, 0xd3, 0xff}, "\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek": {0xff, 0xb6, 0xc1, 0xff}, "l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e": {0xff, 0xa0, 0x7a, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": {0x20, 0xb2, 0xaa, 0xff}, "\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065": {0x87, 0xce, 0xfa, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079": {0x77, 0x88, 0x99, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079": {0x77, 0x88, 0x99, 0xff}, "\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065": {0xb0, 0xc4, 0xde, 0xff}, "l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077": {0xff, 0xff, 0xe0, 0xff}, "\u006c\u0069\u006d\u0065": {0x00, 0xff, 0x00, 0xff}, "\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n": {0x32, 0xcd, 0x32, 0xff}, "\u006c\u0069\u006ee\u006e": {0xfa, 0xf0, 0xe6, 0xff}, "\u006da\u0067\u0065\u006e\u0074\u0061": {0xff, 0x00, 0xff, 0xff}, "\u006d\u0061\u0072\u006f\u006f\u006e": {0x80, 0x00, 0x00, 0xff}, "\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065": {0x66, 0xcd, 0xaa, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065": {0x00, 0x00, 0xcd, 0xff}, "\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064": {0xba, 0x55, 0xd3, 0xff}, "\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065": {0x93, 0x70, 0xdb, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e": {0x3c, 0xb3, 0x71, 0xff}, "\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065": {0x7b, 0x68, 0xee, 0xff}, "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e": {0x00, 0xfa, 0x9a, 0xff}, "\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": {0x48, 0xd1, 0xcc, 0xff}, "\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064": {0xc7, 0x15, 0x85, 0xff}, "\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065": {0x19, 0x19, 0x70, 0xff}, "\u006di\u006e\u0074\u0063\u0072\u0065\u0061m": {0xf5, 0xff, 0xfa, 0xff}, "\u006di\u0073\u0074\u0079\u0072\u006f\u0073e": {0xff, 0xe4, 0xe1, 0xff}, "\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e": {0xff, 0xe4, 0xb5, 0xff}, "n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065": {0xff, 0xde, 0xad, 0xff}, "\u006e\u0061\u0076\u0079": {0x00, 0x00, 0x80, 0xff}, "\u006fl\u0064\u006c\u0061\u0063\u0065": {0xfd, 0xf5, 0xe6, 0xff}, "\u006f\u006c\u0069v\u0065": {0x80, 0x80, 0x00, 0xff}, "\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b": {0x6b, 0x8e, 0x23, 0xff}, "\u006f\u0072\u0061\u006e\u0067\u0065": {0xff, 0xa5, 0x00, 0xff}, "\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d": {0xff, 0x45, 0x00, 0xff}, "\u006f\u0072\u0063\u0068\u0069\u0064": {0xda, 0x70, 0xd6, 0xff}, "\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064": {0xee, 0xe8, 0xaa, 0xff}, "\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n": {0x98, 0xfb, 0x98, 0xff}, "\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065": {0xaf, 0xee, 0xee, 0xff}, "\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064": {0xdb, 0x70, 0x93, 0xff}, "\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070": {0xff, 0xef, 0xd5, 0xff}, "\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f": {0xff, 0xda, 0xb9, 0xff}, "\u0070\u0065\u0072\u0075": {0xcd, 0x85, 0x3f, 0xff}, "\u0070\u0069\u006e\u006b": {0xff, 0xc0, 0xcb, 0xff}, "\u0070\u006c\u0075\u006d": {0xdd, 0xa0, 0xdd, 0xff}, "\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065": {0xb0, 0xe0, 0xe6, 0xff}, "\u0070\u0075\u0072\u0070\u006c\u0065": {0x80, 0x00, 0x80, 0xff}, "\u0072\u0065\u0064": {0xff, 0x00, 0x00, 0xff}, "\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n": {0xbc, 0x8f, 0x8f, 0xff}, "\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e": {0x41, 0x69, 0xe1, 0xff}, "s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e": {0x8b, 0x45, 0x13, 0xff}, "\u0073\u0061\u006c\u006d\u006f\u006e": {0xfa, 0x80, 0x72, 0xff}, "\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e": {0xf4, 0xa4, 0x60, 0xff}, "\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e": {0x2e, 0x8b, 0x57, 0xff}, "\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c": {0xff, 0xf5, 0xee, 0xff}, "\u0073\u0069\u0065\u006e\u006e\u0061": {0xa0, 0x52, 0x2d, 0xff}, "\u0073\u0069\u006c\u0076\u0065\u0072": {0xc0, 0xc0, 0xc0, 0xff}, "\u0073k\u0079\u0062\u006c\u0075\u0065": {0x87, 0xce, 0xeb, 0xff}, "\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e": {0x6a, 0x5a, 0xcd, 0xff}, "\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y": {0x70, 0x80, 0x90, 0xff}, "\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y": {0x70, 0x80, 0x90, 0xff}, "\u0073\u006e\u006f\u0077": {0xff, 0xfa, 0xfa, 0xff}, "s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e": {0x00, 0xff, 0x7f, 0xff}, "\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e": {0x46, 0x82, 0xb4, 0xff}, "\u0074\u0061\u006e": {0xd2, 0xb4, 0x8c, 0xff}, "\u0074\u0065\u0061\u006c": {0x00, 0x80, 0x80, 0xff}, "\u0074h\u0069\u0073\u0074\u006c\u0065": {0xd8, 0xbf, 0xd8, 0xff}, "\u0074\u006f\u006d\u0061\u0074\u006f": {0xff, 0x63, 0x47, 0xff}, "\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e": {0x40, 0xe0, 0xd0, 0xff}, "\u0076\u0069\u006f\u006c\u0065\u0074": {0xee, 0x82, 0xee, 0xff}, "\u0077\u0068\u0065a\u0074": {0xf5, 0xde, 0xb3, 0xff}, "\u0077\u0068\u0069t\u0065": {0xff, 0xff, 0xff, 0xff}, "\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065": {0xf5, 0xf5, 0xf5, 0xff}, "\u0079\u0065\u006c\u006c\u006f\u0077": {0xff, 0xff, 0x00, 0xff}, "y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e": {0x9a, 0xcd, 0x32, 0xff}}

func _bg(_aaef, _ee, _adf float64, _ccc bool, _bbg float64) (float64, float64) {
	_bgb, _be := _g.Sincos(_bbg)
	_gdc, _dc := _g.Sincos(_adf)
	_addd := -_aaef*_bgb*_dc - _ee*_be*_gdc
	_aag := -_aaef*_bgb*_gdc + _ee*_be*_dc
	if !_ccc {
		return -_addd, -_aag
	}
	return _addd, _aag
}

var Names = []string{"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e", "\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065", "\u0061\u0071\u0075\u0061", "\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065", "\u0061\u007a\u0075r\u0065", "\u0062\u0065\u0069g\u0065", "\u0062\u0069\u0073\u0071\u0075\u0065", "\u0062\u006c\u0061c\u006b", "\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064", "\u0062\u006c\u0075\u0065", "\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074", "\u0062\u0072\u006fw\u006e", "\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd", "\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e", "\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065", "\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e", "\u0063\u006f\u0072a\u006c", "\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065", "\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b", "\u0063r\u0069\u006d\u0073\u006f\u006e", "\u0063\u0079\u0061\u006e", "\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065", "\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e", "\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064", "\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079", "\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n", "\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079", "\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi", "d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061", "\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e", "\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065", "\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064", "\u0064a\u0072\u006b\u0072\u0065\u0064", "\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e", "\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079", "\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079", "\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074", "\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b", "d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065", "\u0064i\u006d\u0067\u0072\u0061\u0079", "\u0064i\u006d\u0067\u0072\u0065\u0079", "\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065", "\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k", "f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065", "f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e", "\u0066u\u0063\u0068\u0073\u0069\u0061", "\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o", "\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065", "\u0067\u006f\u006c\u0064", "\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd", "\u0067\u0072\u0061\u0079", "\u0067\u0072\u0065e\u006e", "g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077", "\u0067\u0072\u0065\u0079", "\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077", "\u0068o\u0074\u0070\u0069\u006e\u006b", "\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d", "\u0069\u006e\u0064\u0069\u0067\u006f", "\u0069\u0076\u006fr\u0079", "\u006b\u0068\u0061k\u0069", "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072", "\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068", "\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n", "\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e", "\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e", "\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c", "\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n", "l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077", "\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y", "\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e", "\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y", "\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek", "l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e", "\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065", "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079", "\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079", "\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065", "l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077", "\u006c\u0069\u006d\u0065", "\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n", "\u006c\u0069\u006ee\u006e", "\u006da\u0067\u0065\u006e\u0074\u0061", "\u006d\u0061\u0072\u006f\u006f\u006e", "\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065", "\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064", "\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e", "\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065", "\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e", "\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064", "\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065", "\u006di\u006e\u0074\u0063\u0072\u0065\u0061m", "\u006di\u0073\u0074\u0079\u0072\u006f\u0073e", "\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e", "n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065", "\u006e\u0061\u0076\u0079", "\u006fl\u0064\u006c\u0061\u0063\u0065", "\u006f\u006c\u0069v\u0065", "\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b", "\u006f\u0072\u0061\u006e\u0067\u0065", "\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d", "\u006f\u0072\u0063\u0068\u0069\u0064", "\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064", "\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n", "\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065", "\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064", "\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070", "\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f", "\u0070\u0065\u0072\u0075", "\u0070\u0069\u006e\u006b", "\u0070\u006c\u0075\u006d", "\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065", "\u0070\u0075\u0072\u0070\u006c\u0065", "\u0072\u0065\u0064", "\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n", "\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e", "s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e", "\u0073\u0061\u006c\u006d\u006f\u006e", "\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e", "\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e", "\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c", "\u0073\u0069\u0065\u006e\u006e\u0061", "\u0073\u0069\u006c\u0076\u0065\u0072", "\u0073k\u0079\u0062\u006c\u0075\u0065", "\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e", "\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y", "\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y", "\u0073\u006e\u006f\u0077", "s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e", "\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e", "\u0074\u0061\u006e", "\u0074\u0065\u0061\u006c", "\u0074h\u0069\u0073\u0074\u006c\u0065", "\u0074\u006f\u006d\u0061\u0074\u006f", "\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e", "\u0076\u0069\u006f\u006c\u0065\u0074", "\u0077\u0068\u0065a\u0074", "\u0077\u0068\u0069t\u0065", "\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065", "\u0079\u0065\u006c\u006c\u006f\u0077", "y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e"}

func (_bbf Point) Interpolate(q Point, t float64) Point {
	return Point{(1-t)*_bbf.X + t*q.X, (1-t)*_bbf.Y + t*q.Y}
}

var (
	Aliceblue            = _f.RGBA{0xf0, 0xf8, 0xff, 0xff}
	Antiquewhite         = _f.RGBA{0xfa, 0xeb, 0xd7, 0xff}
	Aqua                 = _f.RGBA{0x00, 0xff, 0xff, 0xff}
	Aquamarine           = _f.RGBA{0x7f, 0xff, 0xd4, 0xff}
	Azure                = _f.RGBA{0xf0, 0xff, 0xff, 0xff}
	Beige                = _f.RGBA{0xf5, 0xf5, 0xdc, 0xff}
	Bisque               = _f.RGBA{0xff, 0xe4, 0xc4, 0xff}
	Black                = _f.RGBA{0x00, 0x00, 0x00, 0xff}
	Blanchedalmond       = _f.RGBA{0xff, 0xeb, 0xcd, 0xff}
	Blue                 = _f.RGBA{0x00, 0x00, 0xff, 0xff}
	Blueviolet           = _f.RGBA{0x8a, 0x2b, 0xe2, 0xff}
	Brown                = _f.RGBA{0xa5, 0x2a, 0x2a, 0xff}
	Burlywood            = _f.RGBA{0xde, 0xb8, 0x87, 0xff}
	Cadetblue            = _f.RGBA{0x5f, 0x9e, 0xa0, 0xff}
	Chartreuse           = _f.RGBA{0x7f, 0xff, 0x00, 0xff}
	Chocolate            = _f.RGBA{0xd2, 0x69, 0x1e, 0xff}
	Coral                = _f.RGBA{0xff, 0x7f, 0x50, 0xff}
	Cornflowerblue       = _f.RGBA{0x64, 0x95, 0xed, 0xff}
	Cornsilk             = _f.RGBA{0xff, 0xf8, 0xdc, 0xff}
	Crimson              = _f.RGBA{0xdc, 0x14, 0x3c, 0xff}
	Cyan                 = _f.RGBA{0x00, 0xff, 0xff, 0xff}
	Darkblue             = _f.RGBA{0x00, 0x00, 0x8b, 0xff}
	Darkcyan             = _f.RGBA{0x00, 0x8b, 0x8b, 0xff}
	Darkgoldenrod        = _f.RGBA{0xb8, 0x86, 0x0b, 0xff}
	Darkgray             = _f.RGBA{0xa9, 0xa9, 0xa9, 0xff}
	Darkgreen            = _f.RGBA{0x00, 0x64, 0x00, 0xff}
	Darkgrey             = _f.RGBA{0xa9, 0xa9, 0xa9, 0xff}
	Darkkhaki            = _f.RGBA{0xbd, 0xb7, 0x6b, 0xff}
	Darkmagenta          = _f.RGBA{0x8b, 0x00, 0x8b, 0xff}
	Darkolivegreen       = _f.RGBA{0x55, 0x6b, 0x2f, 0xff}
	Darkorange           = _f.RGBA{0xff, 0x8c, 0x00, 0xff}
	Darkorchid           = _f.RGBA{0x99, 0x32, 0xcc, 0xff}
	Darkred              = _f.RGBA{0x8b, 0x00, 0x00, 0xff}
	Darksalmon           = _f.RGBA{0xe9, 0x96, 0x7a, 0xff}
	Darkseagreen         = _f.RGBA{0x8f, 0xbc, 0x8f, 0xff}
	Darkslateblue        = _f.RGBA{0x48, 0x3d, 0x8b, 0xff}
	Darkslategray        = _f.RGBA{0x2f, 0x4f, 0x4f, 0xff}
	Darkslategrey        = _f.RGBA{0x2f, 0x4f, 0x4f, 0xff}
	Darkturquoise        = _f.RGBA{0x00, 0xce, 0xd1, 0xff}
	Darkviolet           = _f.RGBA{0x94, 0x00, 0xd3, 0xff}
	Deeppink             = _f.RGBA{0xff, 0x14, 0x93, 0xff}
	Deepskyblue          = _f.RGBA{0x00, 0xbf, 0xff, 0xff}
	Dimgray              = _f.RGBA{0x69, 0x69, 0x69, 0xff}
	Dimgrey              = _f.RGBA{0x69, 0x69, 0x69, 0xff}
	Dodgerblue           = _f.RGBA{0x1e, 0x90, 0xff, 0xff}
	Firebrick            = _f.RGBA{0xb2, 0x22, 0x22, 0xff}
	Floralwhite          = _f.RGBA{0xff, 0xfa, 0xf0, 0xff}
	Forestgreen          = _f.RGBA{0x22, 0x8b, 0x22, 0xff}
	Fuchsia              = _f.RGBA{0xff, 0x00, 0xff, 0xff}
	Gainsboro            = _f.RGBA{0xdc, 0xdc, 0xdc, 0xff}
	Ghostwhite           = _f.RGBA{0xf8, 0xf8, 0xff, 0xff}
	Gold                 = _f.RGBA{0xff, 0xd7, 0x00, 0xff}
	Goldenrod            = _f.RGBA{0xda, 0xa5, 0x20, 0xff}
	Gray                 = _f.RGBA{0x80, 0x80, 0x80, 0xff}
	Green                = _f.RGBA{0x00, 0x80, 0x00, 0xff}
	Greenyellow          = _f.RGBA{0xad, 0xff, 0x2f, 0xff}
	Grey                 = _f.RGBA{0x80, 0x80, 0x80, 0xff}
	Honeydew             = _f.RGBA{0xf0, 0xff, 0xf0, 0xff}
	Hotpink              = _f.RGBA{0xff, 0x69, 0xb4, 0xff}
	Indianred            = _f.RGBA{0xcd, 0x5c, 0x5c, 0xff}
	Indigo               = _f.RGBA{0x4b, 0x00, 0x82, 0xff}
	Ivory                = _f.RGBA{0xff, 0xff, 0xf0, 0xff}
	Khaki                = _f.RGBA{0xf0, 0xe6, 0x8c, 0xff}
	Lavender             = _f.RGBA{0xe6, 0xe6, 0xfa, 0xff}
	Lavenderblush        = _f.RGBA{0xff, 0xf0, 0xf5, 0xff}
	Lawngreen            = _f.RGBA{0x7c, 0xfc, 0x00, 0xff}
	Lemonchiffon         = _f.RGBA{0xff, 0xfa, 0xcd, 0xff}
	Lightblue            = _f.RGBA{0xad, 0xd8, 0xe6, 0xff}
	Lightcoral           = _f.RGBA{0xf0, 0x80, 0x80, 0xff}
	Lightcyan            = _f.RGBA{0xe0, 0xff, 0xff, 0xff}
	Lightgoldenrodyellow = _f.RGBA{0xfa, 0xfa, 0xd2, 0xff}
	Lightgray            = _f.RGBA{0xd3, 0xd3, 0xd3, 0xff}
	Lightgreen           = _f.RGBA{0x90, 0xee, 0x90, 0xff}
	Lightgrey            = _f.RGBA{0xd3, 0xd3, 0xd3, 0xff}
	Lightpink            = _f.RGBA{0xff, 0xb6, 0xc1, 0xff}
	Lightsalmon          = _f.RGBA{0xff, 0xa0, 0x7a, 0xff}
	Lightseagreen        = _f.RGBA{0x20, 0xb2, 0xaa, 0xff}
	Lightskyblue         = _f.RGBA{0x87, 0xce, 0xfa, 0xff}
	Lightslategray       = _f.RGBA{0x77, 0x88, 0x99, 0xff}
	Lightslategrey       = _f.RGBA{0x77, 0x88, 0x99, 0xff}
	Lightsteelblue       = _f.RGBA{0xb0, 0xc4, 0xde, 0xff}
	Lightyellow          = _f.RGBA{0xff, 0xff, 0xe0, 0xff}
	Lime                 = _f.RGBA{0x00, 0xff, 0x00, 0xff}
	Limegreen            = _f.RGBA{0x32, 0xcd, 0x32, 0xff}
	Linen                = _f.RGBA{0xfa, 0xf0, 0xe6, 0xff}
	Magenta              = _f.RGBA{0xff, 0x00, 0xff, 0xff}
	Maroon               = _f.RGBA{0x80, 0x00, 0x00, 0xff}
	Mediumaquamarine     = _f.RGBA{0x66, 0xcd, 0xaa, 0xff}
	Mediumblue           = _f.RGBA{0x00, 0x00, 0xcd, 0xff}
	Mediumorchid         = _f.RGBA{0xba, 0x55, 0xd3, 0xff}
	Mediumpurple         = _f.RGBA{0x93, 0x70, 0xdb, 0xff}
	Mediumseagreen       = _f.RGBA{0x3c, 0xb3, 0x71, 0xff}
	Mediumslateblue      = _f.RGBA{0x7b, 0x68, 0xee, 0xff}
	Mediumspringgreen    = _f.RGBA{0x00, 0xfa, 0x9a, 0xff}
	Mediumturquoise      = _f.RGBA{0x48, 0xd1, 0xcc, 0xff}
	Mediumvioletred      = _f.RGBA{0xc7, 0x15, 0x85, 0xff}
	Midnightblue         = _f.RGBA{0x19, 0x19, 0x70, 0xff}
	Mintcream            = _f.RGBA{0xf5, 0xff, 0xfa, 0xff}
	Mistyrose            = _f.RGBA{0xff, 0xe4, 0xe1, 0xff}
	Moccasin             = _f.RGBA{0xff, 0xe4, 0xb5, 0xff}
	Navajowhite          = _f.RGBA{0xff, 0xde, 0xad, 0xff}
	Navy                 = _f.RGBA{0x00, 0x00, 0x80, 0xff}
	Oldlace              = _f.RGBA{0xfd, 0xf5, 0xe6, 0xff}
	Olive                = _f.RGBA{0x80, 0x80, 0x00, 0xff}
	Olivedrab            = _f.RGBA{0x6b, 0x8e, 0x23, 0xff}
	Orange               = _f.RGBA{0xff, 0xa5, 0x00, 0xff}
	Orangered            = _f.RGBA{0xff, 0x45, 0x00, 0xff}
	Orchid               = _f.RGBA{0xda, 0x70, 0xd6, 0xff}
	Palegoldenrod        = _f.RGBA{0xee, 0xe8, 0xaa, 0xff}
	Palegreen            = _f.RGBA{0x98, 0xfb, 0x98, 0xff}
	Paleturquoise        = _f.RGBA{0xaf, 0xee, 0xee, 0xff}
	Palevioletred        = _f.RGBA{0xdb, 0x70, 0x93, 0xff}
	Papayawhip           = _f.RGBA{0xff, 0xef, 0xd5, 0xff}
	Peachpuff            = _f.RGBA{0xff, 0xda, 0xb9, 0xff}
	Peru                 = _f.RGBA{0xcd, 0x85, 0x3f, 0xff}
	Pink                 = _f.RGBA{0xff, 0xc0, 0xcb, 0xff}
	Plum                 = _f.RGBA{0xdd, 0xa0, 0xdd, 0xff}
	Powderblue           = _f.RGBA{0xb0, 0xe0, 0xe6, 0xff}
	Purple               = _f.RGBA{0x80, 0x00, 0x80, 0xff}
	Red                  = _f.RGBA{0xff, 0x00, 0x00, 0xff}
	Rosybrown            = _f.RGBA{0xbc, 0x8f, 0x8f, 0xff}
	Royalblue            = _f.RGBA{0x41, 0x69, 0xe1, 0xff}
	Saddlebrown          = _f.RGBA{0x8b, 0x45, 0x13, 0xff}
	Salmon               = _f.RGBA{0xfa, 0x80, 0x72, 0xff}
	Sandybrown           = _f.RGBA{0xf4, 0xa4, 0x60, 0xff}
	Seagreen             = _f.RGBA{0x2e, 0x8b, 0x57, 0xff}
	Seashell             = _f.RGBA{0xff, 0xf5, 0xee, 0xff}
	Sienna               = _f.RGBA{0xa0, 0x52, 0x2d, 0xff}
	Silver               = _f.RGBA{0xc0, 0xc0, 0xc0, 0xff}
	Skyblue              = _f.RGBA{0x87, 0xce, 0xeb, 0xff}
	Slateblue            = _f.RGBA{0x6a, 0x5a, 0xcd, 0xff}
	Slategray            = _f.RGBA{0x70, 0x80, 0x90, 0xff}
	Slategrey            = _f.RGBA{0x70, 0x80, 0x90, 0xff}
	Snow                 = _f.RGBA{0xff, 0xfa, 0xfa, 0xff}
	Springgreen          = _f.RGBA{0x00, 0xff, 0x7f, 0xff}
	Steelblue            = _f.RGBA{0x46, 0x82, 0xb4, 0xff}
	Tan                  = _f.RGBA{0xd2, 0xb4, 0x8c, 0xff}
	Teal                 = _f.RGBA{0x00, 0x80, 0x80, 0xff}
	Thistle              = _f.RGBA{0xd8, 0xbf, 0xd8, 0xff}
	Tomato               = _f.RGBA{0xff, 0x63, 0x47, 0xff}
	Turquoise            = _f.RGBA{0x40, 0xe0, 0xd0, 0xff}
	Violet               = _f.RGBA{0xee, 0x82, 0xee, 0xff}
	Wheat                = _f.RGBA{0xf5, 0xde, 0xb3, 0xff}
	White                = _f.RGBA{0xff, 0xff, 0xff, 0xff}
	Whitesmoke           = _f.RGBA{0xf5, 0xf5, 0xf5, 0xff}
	Yellow               = _f.RGBA{0xff, 0xff, 0x00, 0xff}
	Yellowgreen          = _f.RGBA{0x9a, 0xcd, 0x32, 0xff}
)

const _bae = 1e-10

func QuadraticToCubicBezier(startX, startY, x1, y1, x, y float64) (Point, Point) {
	_beg := Point{X: startX, Y: startY}
	_abga := Point{X: x1, Y: y1}
	_gde := Point{X: x, Y: y}
	_baf := _beg.Interpolate(_abga, 2.0/3.0)
	_gcb := _gde.Interpolate(_abga, 2.0/3.0)
	return _baf, _gcb
}

type Point struct{ X, Y float64 }

func _ce(_cfe float64) float64 {
	_cfe = _g.Mod(_cfe, 2.0*_g.Pi)
	if _cfe < 0.0 {
		_cfe += 2.0 * _g.Pi
	}
	return _cfe
}
func (_acf Point) Mul(f float64) Point { return Point{f * _acf.X, f * _acf.Y} }
func _bf(_cd, _gge, _agg, _faf, _gdb float64, _aae, _ffe bool, _gae, _bfc float64) (float64, float64, float64, float64) {
	if _cde(_cd, _gae) && _cde(_gge, _bfc) {
		return _cd, _gge, 0.0, 0.0
	}
	_feb, _abg := _g.Sincos(_gdb)
	_gdd := _abg*(_cd-_gae)/2.0 + _feb*(_gge-_bfc)/2.0
	_bfg := -_feb*(_cd-_gae)/2.0 + _abg*(_gge-_bfc)/2.0
	_gff := _gdd*_gdd/_agg/_agg + _bfg*_bfg/_faf/_faf
	if _gff > 1.0 {
		_agg *= _g.Sqrt(_gff)
		_faf *= _g.Sqrt(_gff)
	}
	_gdg := (_agg*_agg*_faf*_faf - _agg*_agg*_bfg*_bfg - _faf*_faf*_gdd*_gdd) / (_agg*_agg*_bfg*_bfg + _faf*_faf*_gdd*_gdd)
	if _gdg < 0.0 {
		_gdg = 0.0
	}
	_gb := _g.Sqrt(_gdg)
	if _aae == _ffe {
		_gb = -_gb
	}
	_aaeb := _gb * _agg * _bfg / _faf
	_gcc := _gb * -_faf * _gdd / _agg
	_feg := _abg*_aaeb - _feb*_gcc + (_cd+_gae)/2.0
	_cdg := _feb*_aaeb + _abg*_gcc + (_gge+_bfc)/2.0
	_eb := (_gdd - _aaeb) / _agg
	_geg := (_bfg - _gcc) / _faf
	_bab := -(_gdd + _aaeb) / _agg
	_ca := -(_bfg + _gcc) / _faf
	_ac := _g.Acos(_eb / _g.Sqrt(_eb*_eb+_geg*_geg))
	if _geg < 0.0 {
		_ac = -_ac
	}
	_ac = _ce(_ac)
	_aabd := (_eb*_bab + _geg*_ca) / _g.Sqrt((_eb*_eb+_geg*_geg)*(_bab*_bab+_ca*_ca))
	_aabd = _g.Min(1.0, _g.Max(-1.0, _aabd))
	_fb := _g.Acos(_aabd)
	if _eb*_ca-_geg*_bab < 0.0 {
		_fb = -_fb
	}
	if !_ffe && _fb > 0.0 {
		_fb -= 2.0 * _g.Pi
	} else if _ffe && _fb < 0.0 {
		_fb += 2.0 * _g.Pi
	}
	return _feg, _cdg, _ac, _ac + _fb
}
func _cde(_gaea, _baa float64) bool { return _g.Abs(_gaea-_baa) <= _bae }
func EllipseToCubicBeziers(startX, startY, rx, ry, rot float64, large, sweep bool, endX, endY float64) [][4]Point {
	rx = _g.Abs(rx)
	ry = _g.Abs(ry)
	if rx < ry {
		rx, ry = ry, rx
		rot += 90.0
	}
	_fc := _ce(rot * _g.Pi / 180.0)
	if _g.Pi <= _fc {
		_fc -= _g.Pi
	}
	_e, _b, _gg, _fg := _bf(startX, startY, rx, ry, _fc, large, sweep, endX, endY)
	_eg := _g.Pi / 2.0
	_c := int(_g.Ceil(_g.Abs(_fg-_gg) / _eg))
	_eg = _g.Abs(_fg-_gg) / float64(_c)
	_aa := _g.Sin(_eg) * (_g.Sqrt(4.0+3.0*_g.Pow(_g.Tan(_eg/2.0), 2.0)) - 1.0) / 3.0
	if !sweep {
		_eg = -_eg
	}
	_bb := Point{X: startX, Y: startY}
	_bd, _aab := _bg(rx, ry, _fc, sweep, _gg)
	_ec := Point{X: _bd, Y: _aab}
	_bc := [][4]Point{}
	for _ge := 1; _ge < _c+1; _ge++ {
		_ecc := _gg + float64(_ge)*_eg
		_ad, _ba := _gd(rx, ry, _fc, _e, _b, _ecc)
		_cf := Point{X: _ad, Y: _ba}
		_fe, _fa := _bg(rx, ry, _fc, sweep, _ecc)
		_fge := Point{X: _fe, Y: _fa}
		_add := _bb.Add(_ec.Mul(_aa))
		_gc := _cf.Sub(_fge.Mul(_aa))
		_bc = append(_bc, [4]Point{_bb, _add, _gc, _cf})
		_ec = _fge
		_bb = _cf
	}
	return _bc
}
func (_bdf Point) Add(q Point) Point { return Point{_bdf.X + q.X, _bdf.Y + q.Y} }
