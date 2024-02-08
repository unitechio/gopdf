package pdfutil

import (
	_fb "bitbucket.org/shenghui0779/gopdf/common"
	_a "bitbucket.org/shenghui0779/gopdf/contentstream"
	_b "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_g "bitbucket.org/shenghui0779/gopdf/core"
	_e "bitbucket.org/shenghui0779/gopdf/model"
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
func NormalizePage(page *_e.PdfPage) error {
	_ec, _ee := page.GetMediaBox()
	if _ee != nil {
		return _ee
	}
	_af, _ee := page.GetRotate()
	if _ee != nil {
		_fb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _ee.Error())
	}
	_ba := _af%360 != 0 && _af%90 == 0
	_ec.Normalize()
	_d, _ga, _bg, _c := _ec.Llx, _ec.Lly, _ec.Width(), _ec.Height()
	_ca := _d != 0 || _ga != 0
	if !_ba && !_ca {
		return nil
	}
	_ae := func(_da, _cd, _bgf float64) _b.BoundingBox {
		return _b.Path{Points: []_b.Point{_b.NewPoint(0, 0).Rotate(_bgf), _b.NewPoint(_da, 0).Rotate(_bgf), _b.NewPoint(0, _cd).Rotate(_bgf), _b.NewPoint(_da, _cd).Rotate(_bgf)}}.GetBoundingBox()
	}
	_gd := _a.NewContentCreator()
	var _bb float64
	if _ba {
		_bb = -float64(_af)
		_gad := _ae(_bg, _c, _bb)
		_gd.Translate((_gad.Width-_bg)/2+_bg/2, (_gad.Height-_c)/2+_c/2)
		_gd.RotateDeg(_bb)
		_gd.Translate(-_bg/2, -_c/2)
		_bg, _c = _gad.Width, _gad.Height
	}
	if _ca {
		_gd.Translate(-_d, -_ga)
	}
	_eca := _gd.Operations()
	_fa, _ee := _g.MakeStream(_eca.Bytes(), _g.NewFlateEncoder())
	if _ee != nil {
		return _ee
	}
	_fd := _g.MakeArray(_fa)
	_fd.Append(page.GetContentStreamObjs()...)
	*_ec = _e.PdfRectangle{Urx: _bg, Ury: _c}
	if _afb := page.CropBox; _afb != nil {
		_afb.Normalize()
		_ac, _gae, _bab, _cg := _afb.Llx-_d, _afb.Lly-_ga, _afb.Width(), _afb.Height()
		if _ba {
			_caf := _ae(_bab, _cg, _bb)
			_bab, _cg = _caf.Width, _caf.Height
		}
		*_afb = _e.PdfRectangle{Llx: _ac, Lly: _gae, Urx: _ac + _bab, Ury: _gae + _cg}
	}
	_fb.Log.Debug("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066", _bb, _eca, _ec)
	page.Contents = _fd
	page.Rotate = nil
	return nil
}
