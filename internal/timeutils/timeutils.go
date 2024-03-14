package timeutils

import (
	_c "errors"
	_cf "fmt"
	_gf "regexp"
	_d "strconv"
	_ca "time"
)

func FormatPdfTime(in _ca.Time) string {
	_e := in.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_ea, _ := _d.ParseInt(_e[1:3], 10, 32)
	_a, _ := _d.ParseInt(_e[4:6], 10, 32)
	_ac := int64(in.Year())
	_b := int64(in.Month())
	_ef := int64(in.Day())
	_dg := int64(in.Hour())
	_cfb := int64(in.Minute())
	_bg := int64(in.Second())
	_eff := _e[0]
	return _cf.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _ac, _b, _ef, _dg, _cfb, _bg, _eff, _ea, _a)
}

func ParsePdfTime(pdfTime string) (_ca.Time, error) {
	_bf := _ag.FindAllStringSubmatch(pdfTime, 1)
	if len(_bf) < 1 {
		if len(pdfTime) > 0 && pdfTime[0] != 'D' {
			pdfTime = _cf.Sprintf("\u0044\u003a\u0025\u0073", pdfTime)
			return ParsePdfTime(pdfTime)
		}
		return _ca.Time{}, _cf.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029", pdfTime)
	}
	if len(_bf[0]) != 10 {
		return _ca.Time{}, _c.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030")
	}
	_ba, _ := _d.ParseInt(_bf[0][1], 10, 32)
	_gg, _ := _d.ParseInt(_bf[0][2], 10, 32)
	_da, _ := _d.ParseInt(_bf[0][3], 10, 32)
	_gb, _ := _d.ParseInt(_bf[0][4], 10, 32)
	_db, _ := _d.ParseInt(_bf[0][5], 10, 32)
	_ed, _ := _d.ParseInt(_bf[0][6], 10, 32)
	var (
		_cc byte
		_fa int64
		_ce int64
	)
	if len(_bf[0][7]) > 0 {
		_cc = _bf[0][7][0]
	} else {
		_cc = '+'
	}
	if len(_bf[0][8]) > 0 {
		_fa, _ = _d.ParseInt(_bf[0][8], 10, 32)
	} else {
		_fa = 0
	}
	if len(_bf[0][9]) > 0 {
		_ce, _ = _d.ParseInt(_bf[0][9], 10, 32)
	} else {
		_ce = 0
	}
	_eba := int(_fa*60*60 + _ce*60)
	switch _cc {
	case '-':
		_eba = -_eba
	case 'Z':
		_eba = 0
	}
	_ff := _cf.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _cc, _fa, _ce)
	_fb := _ca.FixedZone(_ff, _eba)
	return _ca.Date(int(_ba), _ca.Month(_gg), int(_da), int(_gb), int(_db), int(_ed), 0, _fb), nil
}

var _ag = _gf.MustCompile("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f")
