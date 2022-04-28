package decoder

import (
	_c "image"

	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_g "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_eg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_ca *Decoder) decodePageImage(_ec int) (_c.Image, error) {
	const _bdd = "\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	if _ec < 0 {
		return nil, _d.Errorf(_bdd, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _ec)
	}
	if _ec > int(_ca._cg.NumberOfPages) {
		return nil, _d.Errorf(_bdd, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _ec)
	}
	_cf, _cge := _ca._cg.GetPage(_ec)
	if _cge != nil {
		return nil, _d.Wrap(_cge, _bdd, "")
	}
	_be, _cge := _cf.GetBitmap()
	if _cge != nil {
		return nil, _d.Wrap(_cge, _bdd, "")
	}
	_be.InverseData()
	return _be.ToImage(), nil
}
func (_bda *Decoder) DecodePageImage(pageNumber int) (_c.Image, error) {
	const _dd = "\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	_ddf, _aaa := _bda.decodePageImage(pageNumber)
	if _aaa != nil {
		return nil, _d.Wrap(_aaa, _dd, "")
	}
	return _ddf, nil
}
func (_aa *Decoder) DecodePage(pageNumber int) ([]byte, error) { return _aa.decodePage(pageNumber) }
func (_gb *Decoder) decodePage(_bc int) ([]byte, error) {
	const _de = "\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065"
	if _bc < 0 {
		return nil, _d.Errorf(_de, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _bc)
	}
	if _bc > int(_gb._cg.NumberOfPages) {
		return nil, _d.Errorf(_de, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _bc)
	}
	_gf, _ga := _gb._cg.GetPage(_bc)
	if _ga != nil {
		return nil, _d.Wrap(_ga, _de, "")
	}
	_ac, _ga := _gf.GetBitmap()
	if _ga != nil {
		return nil, _d.Wrap(_ga, _de, "")
	}
	_ac.InverseData()
	if !_gb._bd.UnpaddedData {
		return _ac.Data, nil
	}
	return _ac.GetUnpaddedData()
}
func Decode(input []byte, parameters Parameters, globals *_eg.Globals) (*Decoder, error) {
	_dc := _f.NewReader(input)
	_cfd, _aab := _eg.DecodeDocument(_dc, globals)
	if _aab != nil {
		return nil, _aab
	}
	return &Decoder{_a: _dc, _cg: _cfd, _bd: parameters}, nil
}
func (_bg *Decoder) DecodeNextPage() ([]byte, error) {
	_bg._b++
	_fg := _bg._b
	return _bg.decodePage(_fg)
}
func (_gd *Decoder) PageNumber() (int, error) {
	const _db = "\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072"
	if _gd._cg == nil {
		return 0, _d.Error(_db, "d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074")
	}
	return int(_gd._cg.NumberOfPages), nil
}

type Decoder struct {
	_a  _f.StreamReader
	_cg *_eg.Document
	_b  int
	_bd Parameters
}
type Parameters struct {
	UnpaddedData bool
	Color        _g.Color
}
