package textshaping

import (
	_c "strings"

	_e "github.com/unidoc/garabic"
	_b "golang.org/x/text/unicode/bidi"
)

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape(text string) (string, error) {
	_eg := _b.Paragraph{}
	_eg.SetString(text)
	_bb, _bd := _eg.Order()
	if _bd != nil {
		return "", _bd
	}
	for _aa := 0; _aa < _bb.NumRuns(); _aa++ {
		_g := _bb.Run(_aa)
		_d := _g.String()
		if _g.Direction() == _b.RightToLeft {
			var (
				_gg = _e.Shape(_d)
				_ga = []rune(_gg)
				_de = make([]rune, len(_ga))
			)
			_ed := 0
			for _cc := len(_ga) - 1; _cc >= 0; _cc-- {
				_de[_ed] = _ga[_cc]
				_ed++
			}
			_d = string(_de)
			text = _c.Replace(text, _c.TrimSpace(_g.String()), _d, 1)
		}
	}
	return text, nil
}
