package pdfutil

import (
	_fa "bitbucket.org/shenghui0779/gopdf/common"
	_e "bitbucket.org/shenghui0779/gopdf/contentstream"
	_b "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_gc "bitbucket.org/shenghui0779/gopdf/core"
	_g "bitbucket.org/shenghui0779/gopdf/model"
)

// NormalizePage performs the following operations on the passed in page:
// - Normalize the page rotation.
//   Rotates the contents of the page according to the Rotate entry, thus
//   flattening the rotation. The Rotate entry of the page is set to nil.
// - Normalize the media box.
//   If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//   the contents of the page are translated to (-Llx, -Lly). After
//   normalization, the media box is updated (Llx and Lly are set to 0 and
//   Urx and Ury are updated accordingly).
// - Normalize the crop box.
//   The crop box of the page is updated based on the previous operations.
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage(page *_g.PdfPage) error {
	_c, _a := page.GetMediaBox()
	if _a != nil {
		return _a
	}
	_gd, _a := page.GetRotate()
	if _a != nil {
		_fa.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _a.Error())
	}
	_fe := _gd%360 != 0 && _gd%90 == 0
	_c.Normalize()
	_bf, _ee, _ef, _d := _c.Llx, _c.Lly, _c.Width(), _c.Height()
	_bg := _bf != 0 || _ee != 0
	if !_fe && !_bg {
		return nil
	}
	_af := func(_bb, _ab, _bd float64) _b.BoundingBox {
		return _b.Path{Points: []_b.Point{_b.NewPoint(0, 0).Rotate(_bd), _b.NewPoint(_bb, 0).Rotate(_bd), _b.NewPoint(0, _ab).Rotate(_bd), _b.NewPoint(_bb, _ab).Rotate(_bd)}}.GetBoundingBox()
	}
	_fc := _e.NewContentCreator()
	var _fg float64
	if _fe {
		_fg = -float64(_gd)
		_aa := _af(_ef, _d, _fg)
		_fc.Translate((_aa.Width-_ef)/2+_ef/2, (_aa.Height-_d)/2+_d/2)
		_fc.RotateDeg(_fg)
		_fc.Translate(-_ef/2, -_d/2)
		_ef, _d = _aa.Width, _aa.Height
	}
	if _bg {
		_fc.Translate(-_bf, -_ee)
	}
	_dc := _fc.Operations()
	_eef, _a := _gc.MakeStream(_dc.Bytes(), _gc.NewFlateEncoder())
	if _a != nil {
		return _a
	}
	_dcf := _gc.MakeArray(_eef)
	_dcf.Append(page.GetContentStreamObjs()...)
	*_c = _g.PdfRectangle{Urx: _ef, Ury: _d}
	if _fd := page.CropBox; _fd != nil {
		_fd.Normalize()
		_gg, _ae, _dd, _ad := _fd.Llx-_bf, _fd.Lly-_ee, _fd.Width(), _fd.Height()
		if _fe {
			_ag := _af(_dd, _ad, _fg)
			_dd, _ad = _ag.Width, _ag.Height
		}
		*_fd = _g.PdfRectangle{Llx: _gg, Lly: _ae, Urx: _gg + _dd, Ury: _ae + _ad}
	}
	_fa.Log.Debug("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066", _fg, _dc, _c)
	page.Contents = _dcf
	page.Rotate = nil
	return nil
}
