package decoder

import (
	_c "image"

	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_b "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_fb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_bg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_bc *Decoder) DecodePage(pageNumber int) ([]byte, error) { return _bc.decodePage(pageNumber) }
func (_fbe *Decoder) decodePageImage(_cf int) (_c.Image, error) {
	const _af = "\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	if _cf < 0 {
		return nil, _bg.Errorf(_af, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _cf)
	}
	if _cf > int(_fbe._eb.NumberOfPages) {
		return nil, _bg.Errorf(_af, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _cf)
	}
	_baa, _ee := _fbe._eb.GetPage(_cf)
	if _ee != nil {
		return nil, _bg.Wrap(_ee, _af, "")
	}
	_gc, _ee := _baa.GetBitmap()
	if _ee != nil {
		return nil, _bg.Wrap(_ee, _af, "")
	}
	_gc.InverseData()
	return _gc.ToImage(), nil
}
func (_d *Decoder) DecodePageImage(pageNumber int) (_c.Image, error) {
	const _ba = "\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065"
	_cb, _ff := _d.decodePageImage(pageNumber)
	if _ff != nil {
		return nil, _bg.Wrap(_ff, _ba, "")
	}
	return _cb, nil
}
func (_bca *Decoder) PageNumber() (int, error) {
	const _dd = "\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072"
	if _bca._eb == nil {
		return 0, _bg.Error(_dd, "d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074")
	}
	return int(_bca._eb.NumberOfPages), nil
}
func (_gd *Decoder) DecodeNextPage() ([]byte, error) {
	_gd._ga++
	_db := _gd._ga
	return _gd.decodePage(_db)
}
func (_dg *Decoder) decodePage(_a int) ([]byte, error) {
	const _bd = "\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065"
	if _a < 0 {
		return nil, _bg.Errorf(_bd, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'", _a)
	}
	if _a > int(_dg._eb.NumberOfPages) {
		return nil, _bg.Errorf(_bd, "p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072", _a)
	}
	_fa, _ac := _dg._eb.GetPage(_a)
	if _ac != nil {
		return nil, _bg.Wrap(_ac, _bd, "")
	}
	_bf, _ac := _fa.GetBitmap()
	if _ac != nil {
		return nil, _bg.Wrap(_ac, _bd, "")
	}
	_bf.InverseData()
	if !_dg._fd.UnpaddedData {
		return _bf.Data, nil
	}
	return _bf.GetUnpaddedData()
}
func Decode(input []byte, parameters Parameters, globals *_fb.Globals) (*Decoder, error) {
	_aff := _f.NewReader(input)
	_ae, _cc := _fb.DecodeDocument(_aff, globals)
	if _cc != nil {
		return nil, _cc
	}
	return &Decoder{_g: _aff, _eb: _ae, _fd: parameters}, nil
}

type Parameters struct {
	UnpaddedData bool
	Color        _b.Color
}
type Decoder struct {
	_g  *_f.Reader
	_eb *_fb.Document
	_ga int
	_fd Parameters
}
