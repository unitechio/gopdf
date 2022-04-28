package timeutils

import (
	_dg "errors"
	_e "fmt"
	_ab "regexp"
	_af "strconv"
	_a "time"
)

var _gg = _ab.MustCompile("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f")

func ParsePdfTime(pdfTime string) (_a.Time, error) {
	_bf := _gg.FindAllStringSubmatch(pdfTime, 1)
	if len(_bf) < 1 {
		return _a.Time{}, _e.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029", pdfTime)
	}
	if len(_bf[0]) != 10 {
		return _a.Time{}, _dg.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030")
	}
	_afb, _ := _af.ParseInt(_bf[0][1], 10, 32)
	_fa, _ := _af.ParseInt(_bf[0][2], 10, 32)
	_abg, _ := _af.ParseInt(_bf[0][3], 10, 32)
	_be, _ := _af.ParseInt(_bf[0][4], 10, 32)
	_cb, _ := _af.ParseInt(_bf[0][5], 10, 32)
	_ced, _ := _af.ParseInt(_bf[0][6], 10, 32)
	var (
		_cd byte
		_bg int64
		_eb int64
	)
	if len(_bf[0][7]) > 0 {
		_cd = _bf[0][7][0]
	} else {
		_cd = '+'
	}
	if len(_bf[0][8]) > 0 {
		_bg, _ = _af.ParseInt(_bf[0][8], 10, 32)
	} else {
		_bg = 0
	}
	if len(_bf[0][9]) > 0 {
		_eb, _ = _af.ParseInt(_bf[0][9], 10, 32)
	} else {
		_eb = 0
	}
	_bef := int(_bg*60*60 + _eb*60)
	switch _cd {
	case '-':
		_bef = -_bef
	case 'Z':
		_bef = 0
	}
	_cde := _e.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _cd, _bg, _eb)
	_ebd := _a.FixedZone(_cde, _bef)
	return _a.Date(int(_afb), _a.Month(_fa), int(_abg), int(_be), int(_cb), int(_ced), 0, _ebd), nil
}
func FormatPdfTime(in _a.Time) string {
	_c := in.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_b, _ := _af.ParseInt(_c[1:3], 10, 32)
	_g, _ := _af.ParseInt(_c[4:6], 10, 32)
	_cf := int64(in.Year())
	_ca := int64(in.Month())
	_f := int64(in.Day())
	_ed := int64(in.Hour())
	_fb := int64(in.Minute())
	_fe := int64(in.Second())
	_ce := _c[0]
	return _e.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _cf, _ca, _f, _ed, _fb, _fe, _ce, _b, _g)
}
