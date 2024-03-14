package pdfutil

import (
	_c "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/contentstream"
	_a "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_ce "bitbucket.org/shenghui0779/gopdf/core"
	_ge "bitbucket.org/shenghui0779/gopdf/model"
)

// NormalizePage performs the following operations on the passed in page:
//   - Normalize the page rotation.
//     Rotates the contents of the page according to the Rotate entry, thus
//     flattening the rotation. The Rotate entry of the page is set to nil.
//   - Normalize the media box.
//     If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//     the contents of the page are translated to (-Llx, -Lly). After
//     normalization, the media box is updated (Llx and Lly are set to 0 and
//     Urx and Ury are updated accordingly).
//   - Normalize the crop box.
//     The crop box of the page is updated based on the previous operations.
//
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage(page *_ge.PdfPage) error {
	_gee, _db := page.GetMediaBox()
	if _db != nil {
		return _db
	}
	_ad, _db := page.GetRotate()
	if _db != nil {
		_c.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _db.Error())
	}
	_e := _ad%360 != 0 && _ad%90 == 0
	_gee.Normalize()
	_gg, _ed, _b, _ae := _gee.Llx, _gee.Lly, _gee.Width(), _gee.Height()
	_ga := _gg != 0 || _ed != 0
	if !_e && !_ga {
		return nil
	}
	_f := func(_dg, _gb, _dbb float64) _a.BoundingBox {
		return _a.Path{Points: []_a.Point{_a.NewPoint(0, 0).Rotate(_dbb), _a.NewPoint(_dg, 0).Rotate(_dbb), _a.NewPoint(0, _gb).Rotate(_dbb), _a.NewPoint(_dg, _gb).Rotate(_dbb)}}.GetBoundingBox()
	}
	_ba := _d.NewContentCreator()
	var _bg float64
	if _e {
		_bg = -float64(_ad)
		_fb := _f(_b, _ae, _bg)
		_ba.Translate((_fb.Width-_b)/2+_b/2, (_fb.Height-_ae)/2+_ae/2)
		_ba.RotateDeg(_bg)
		_ba.Translate(-_b/2, -_ae/2)
		_b, _ae = _fb.Width, _fb.Height
	}
	if _ga {
		_ba.Translate(-_gg, -_ed)
	}
	_df := _ba.Operations()
	_bad, _db := _ce.MakeStream(_df.Bytes(), _ce.NewFlateEncoder())
	if _db != nil {
		return _db
	}
	_fd := _ce.MakeArray(_bad)
	_fd.Append(page.GetContentStreamObjs()...)
	*_gee = _ge.PdfRectangle{Urx: _b, Ury: _ae}
	if _gf := page.CropBox; _gf != nil {
		_gf.Normalize()
		_dd, _cd, _bd, _gac := _gf.Llx-_gg, _gf.Lly-_ed, _gf.Width(), _gf.Height()
		if _e {
			_bc := _f(_bd, _gac, _bg)
			_bd, _gac = _bc.Width, _bc.Height
		}
		*_gf = _ge.PdfRectangle{Llx: _dd, Lly: _cd, Urx: _dd + _bd, Ury: _cd + _gac}
	}
	_c.Log.Debug("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066", _bg, _df, _gee)
	page.Contents = _fd
	page.Rotate = nil
	return nil
}
