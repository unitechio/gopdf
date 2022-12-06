package timeutils

import (
	_d "errors"
	_a "fmt"
	_cf "regexp"
	_e "strconv"
	_db "time"
)

var _bgf = _cf.MustCompile("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f")

func ParsePdfTime(pdfTime string) (_db.Time, error) {
	_ed := _bgf.FindAllStringSubmatch(pdfTime, 1)
	if len(_ed) < 1 {
		return _db.Time{}, _a.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029", pdfTime)
	}
	if len(_ed[0]) != 10 {
		return _db.Time{}, _d.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030")
	}
	_ebd, _ := _e.ParseInt(_ed[0][1], 10, 32)
	_fa, _ := _e.ParseInt(_ed[0][2], 10, 32)
	_bca, _ := _e.ParseInt(_ed[0][3], 10, 32)
	_fda, _ := _e.ParseInt(_ed[0][4], 10, 32)
	_gc, _ := _e.ParseInt(_ed[0][5], 10, 32)
	_bg, _ := _e.ParseInt(_ed[0][6], 10, 32)
	var (
		_ca  byte
		_fdf int64
		_ec  int64
	)
	if len(_ed[0][7]) > 0 {
		_ca = _ed[0][7][0]
	} else {
		_ca = '+'
	}
	if len(_ed[0][8]) > 0 {
		_fdf, _ = _e.ParseInt(_ed[0][8], 10, 32)
	} else {
		_fdf = 0
	}
	if len(_ed[0][9]) > 0 {
		_ec, _ = _e.ParseInt(_ed[0][9], 10, 32)
	} else {
		_ec = 0
	}
	_cc := int(_fdf*60*60 + _ec*60)
	switch _ca {
	case '-':
		_cc = -_cc
	case 'Z':
		_cc = 0
	}
	_gcf := _a.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _ca, _fdf, _ec)
	_bcb := _db.FixedZone(_gcf, _cc)
	return _db.Date(int(_ebd), _db.Month(_fa), int(_bca), int(_fda), int(_gc), int(_bg), 0, _bcb), nil
}
func FormatPdfTime(in _db.Time) string {
	_f := in.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_fd, _ := _e.ParseInt(_f[1:3], 10, 32)
	_g, _ := _e.ParseInt(_f[4:6], 10, 32)
	_ce := int64(in.Year())
	_bc := int64(in.Month())
	_ff := int64(in.Day())
	_dc := int64(in.Hour())
	_gd := int64(in.Minute())
	_eb := int64(in.Second())
	_ad := _f[0]
	return _a.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _ce, _bc, _ff, _dc, _gd, _eb, _ad, _fd, _g)
}
