package pdfutil

import (
	_b "bitbucket.org/shenghui0779/gopdf/common"
	_a "bitbucket.org/shenghui0779/gopdf/contentstream"
	_fe "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_f "bitbucket.org/shenghui0779/gopdf/core"
	_eg "bitbucket.org/shenghui0779/gopdf/model"
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
func NormalizePage(page *_eg.PdfPage) error {
	_g, _ge := page.GetMediaBox()
	if _ge != nil {
		return _ge
	}
	_af, _ge := page.GetRotate()
	if _ge != nil {
		_b.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _ge.Error())
	}
	_ag := _af%360 != 0 && _af%90 == 0
	_g.Normalize()
	_gb, _fed, _be, _fb := _g.Llx, _g.Lly, _g.Width(), _g.Height()
	_agf := _gb != 0 || _fed != 0
	if !_ag && !_agf {
		return nil
	}
	_fbf := func(_d, _dg, _ege float64) _fe.BoundingBox {
		return _fe.Path{Points: []_fe.Point{_fe.NewPoint(0, 0).Rotate(_ege), _fe.NewPoint(_d, 0).Rotate(_ege), _fe.NewPoint(0, _dg).Rotate(_ege), _fe.NewPoint(_d, _dg).Rotate(_ege)}}.GetBoundingBox()
	}
	_aa := _a.NewContentCreator()
	var _gc float64
	if _ag {
		_gc = -float64(_af)
		_ac := _fbf(_be, _fb, _gc)
		_aa.Translate((_ac.Width-_be)/2+_be/2, (_ac.Height-_fb)/2+_fb/2)
		_aa.RotateDeg(_gc)
		_aa.Translate(-_be/2, -_fb/2)
		_be, _fb = _ac.Width, _ac.Height
	}
	if _agf {
		_aa.Translate(-_gb, -_fed)
	}
	_gbe := _aa.Operations()
	_fg, _ge := _f.MakeStream(_gbe.Bytes(), _f.NewFlateEncoder())
	if _ge != nil {
		return _ge
	}
	_ef := _f.MakeArray(_fg)
	_ef.Append(page.GetContentStreamObjs()...)
	*_g = _eg.PdfRectangle{Urx: _be, Ury: _fb}
	if _bed := page.CropBox; _bed != nil {
		_bed.Normalize()
		_c, _bg, _afa, _ce := _bed.Llx-_gb, _bed.Lly-_fed, _bed.Width(), _bed.Height()
		if _ag {
			_cb := _fbf(_afa, _ce, _gc)
			_afa, _ce = _cb.Width, _cb.Height
		}
		*_bed = _eg.PdfRectangle{Llx: _c, Lly: _bg, Urx: _c + _afa, Ury: _bg + _ce}
	}
	_b.Log.Debug("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066", _gc, _gbe, _g)
	page.Contents = _ef
	page.Rotate = nil
	return nil
}
