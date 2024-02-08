package timeutils

import (
	_a "errors"
	_ce "fmt"
	_c "regexp"
	_e "strconv"
	_d "time"
)

func FormatPdfTime(in _d.Time) string {
	_ae := in.Format("\u002d\u0030\u0037\u003a\u0030\u0030")
	_de, _ := _e.ParseInt(_ae[1:3], 10, 32)
	_eb, _ := _e.ParseInt(_ae[4:6], 10, 32)
	_aeb := int64(in.Year())
	_cc := int64(in.Month())
	_fg := int64(in.Day())
	_eee := int64(in.Hour())
	_fgg := int64(in.Minute())
	_fga := int64(in.Second())
	_b := _ae[0]
	return _ce.Sprintf("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027", _aeb, _cc, _fg, _eee, _fgg, _fga, _b, _de, _eb)
}
func ParsePdfTime(pdfTime string) (_d.Time, error) {
	_dc := _bc.FindAllStringSubmatch(pdfTime, 1)
	if len(_dc) < 1 {
		if len(pdfTime) > 0 && pdfTime[0] != 'D' {
			pdfTime = _ce.Sprintf("\u0044\u003a\u0025\u0073", pdfTime)
			return ParsePdfTime(pdfTime)
		}
		return _d.Time{}, _ce.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029", pdfTime)
	}
	if len(_dc[0]) != 10 {
		return _d.Time{}, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030")
	}
	_bg, _ := _e.ParseInt(_dc[0][1], 10, 32)
	_g, _ := _e.ParseInt(_dc[0][2], 10, 32)
	_ccb, _ := _e.ParseInt(_dc[0][3], 10, 32)
	_fe, _ := _e.ParseInt(_dc[0][4], 10, 32)
	_dbb, _ := _e.ParseInt(_dc[0][5], 10, 32)
	_bb, _ := _e.ParseInt(_dc[0][6], 10, 32)
	var (
		_aa  byte
		_fee int64
		_gg  int64
	)
	if len(_dc[0][7]) > 0 {
		_aa = _dc[0][7][0]
	} else {
		_aa = '+'
	}
	if len(_dc[0][8]) > 0 {
		_fee, _ = _e.ParseInt(_dc[0][8], 10, 32)
	} else {
		_fee = 0
	}
	if len(_dc[0][9]) > 0 {
		_gg, _ = _e.ParseInt(_dc[0][9], 10, 32)
	} else {
		_gg = 0
	}
	_eg := int(_fee*60*60 + _gg*60)
	switch _aa {
	case '-':
		_eg = -_eg
	case 'Z':
		_eg = 0
	}
	_gc := _ce.Sprintf("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064", _aa, _fee, _gg)
	_eeb := _d.FixedZone(_gc, _eg)
	return _d.Date(int(_bg), _d.Month(_g), int(_ccb), int(_fe), int(_dbb), int(_bb), 0, _eeb), nil
}

var _bc = _c.MustCompile("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f")
