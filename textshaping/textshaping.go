package textshaping

import (
	_b "strings"

	_c "github.com/unidoc/garabic"
	_f "golang.org/x/text/unicode/bidi"
)

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape(text string) (string, error) {
	_gc := _f.Paragraph{}
	_gc.SetString(text)
	_e, _d := _gc.Order()
	if _d != nil {
		return "", _d
	}
	for _ca := 0; _ca < _e.NumRuns(); _ca++ {
		_ee := _e.Run(_ca)
		_bc := _ee.String()
		if _ee.Direction() == _f.RightToLeft {
			var (
				_cg = _c.Shape(_bc)
				_fg = []rune(_cg)
				_ba = make([]rune, len(_fg))
			)
			_cc := 0
			for _bf := len(_fg) - 1; _bf >= 0; _bf-- {
				_ba[_cc] = _fg[_bf]
				_cc++
			}
			_bc = string(_ba)
			text = _b.Replace(text, _b.TrimSpace(_ee.String()), _bc, 1)
		}
	}
	return text, nil
}
