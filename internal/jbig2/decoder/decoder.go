package decoder

import (
	_b "image"

	_a "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_gd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_g "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

type Parameters struct {
	UnpaddedData bool
	Color        _d.Color
}

func (_cg *Decoder) PageNumber() (int, error) {
	const _ea = "\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072"
	if _cg._gda == nil {
		return 0, _g.Error(_ea, "d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074")
	}
	return int(_cg._gda.NumberOfPages), nil
}

func (_bag *Decoder) DecodePageImage(pageNumber int) (_b.Image, error) {
	const _ad = "\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	_aa, _bf := _bag.decodePageImage(pageNumber)
	if _bf != nil {
		return nil, _g.Wrap(_bf, _ad, "")
	}
	return _aa, nil
}

func (_aaf *Decoder) decodePage(_af int) ([]byte, error) {
	const _bd = "\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065"
	if _af < 0 {
		return nil, _g.Errorf(_bd, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _af)
	}
	if _af > int(_aaf._gda.NumberOfPages) {
		return nil, _g.Errorf(_bd, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _af)
	}
	_fb, _gaa := _aaf._gda.GetPage(_af)
	if _gaa != nil {
		return nil, _g.Wrap(_gaa, _bd, "")
	}
	_bb, _gaa := _fb.GetBitmap()
	if _gaa != nil {
		return nil, _g.Wrap(_gaa, _bd, "")
	}
	_bb.InverseData()
	if !_aaf._bad.UnpaddedData {
		return _bb.Data, nil
	}
	return _bb.GetUnpaddedData()
}

type Decoder struct {
	_ba  *_a.Reader
	_gda *_gd.Document
	_c   int
	_bad Parameters
}

func (_df *Decoder) DecodeNextPage() ([]byte, error) {
	_df._c++
	_ga := _df._c
	return _df.decodePage(_ga)
}
func (_e *Decoder) DecodePage(pageNumber int) ([]byte, error) { return _e.decodePage(pageNumber) }
func (_ae *Decoder) decodePageImage(_ce int) (_b.Image, error) {
	const _bg = "\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	if _ce < 0 {
		return nil, _g.Errorf(_bg, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _ce)
	}
	if _ce > int(_ae._gda.NumberOfPages) {
		return nil, _g.Errorf(_bg, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _ce)
	}
	_be, _bbc := _ae._gda.GetPage(_ce)
	if _bbc != nil {
		return nil, _g.Wrap(_bbc, _bg, "")
	}
	_add, _bbc := _be.GetBitmap()
	if _bbc != nil {
		return nil, _g.Wrap(_bbc, _bg, "")
	}
	_add.InverseData()
	return _add.ToImage(), nil
}

func Decode(input []byte, parameters Parameters, globals *_gd.Globals) (*Decoder, error) {
	_gac := _a.NewReader(input)
	_cc, _cd := _gd.DecodeDocument(_gac, globals)
	if _cd != nil {
		return nil, _cd
	}
	return &Decoder{_ba: _gac, _gda: _cc, _bad: parameters}, nil
}
