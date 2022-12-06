package textshaping

import (
	_b "strings"

	_g "github.com/unidoc/garabic"
	_a "golang.org/x/text/unicode/bidi"
)

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape(text string) (string, error) {
	_cf := _a.Paragraph{}
	_cf.SetString(text)
	_ac, _bg := _cf.Order()
	if _bg != nil {
		return "", _bg
	}
	for _gg := 0; _gg < _ac.NumRuns(); _gg++ {
		_d := _ac.Run(_gg)
		_bc := _d.String()
		if _d.Direction() == _a.RightToLeft {
			var (
				_ggg = _g.Shape(_bc)
				_af  = []rune(_ggg)
				_bb  = make([]rune, len(_af))
			)
			_bbf := 0
			for _f := len(_af) - 1; _f >= 0; _f-- {
				_bb[_bbf] = _af[_f]
				_bbf++
			}
			_bc = string(_bb)
			text = _b.Replace(text, _b.TrimSpace(_d.String()), _bc, 1)
		}
	}
	return text, nil
}
