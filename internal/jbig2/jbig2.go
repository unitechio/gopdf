package jbig2

import (
	_a "sort"

	_c "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_b "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder"
	_fa "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_aa "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document/segments"
	_g "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

type Globals map[int]*_aa.Header

func DecodeBytes(encoded []byte, parameters _b.Parameters, globals ...Globals) ([]byte, error) {
	var _fb Globals
	if len(globals) > 0 {
		_fb = globals[0]
	}
	_gf, _ba := _b.Decode(encoded, parameters, _fb.ToDocumentGlobals())
	if _ba != nil {
		return nil, _ba
	}
	return _gf.DecodeNextPage()
}
func (_gc Globals) ToDocumentGlobals() *_fa.Globals {
	if _gc == nil {
		return nil
	}
	_d := []*_aa.Header{}
	for _, _ab := range _gc {
		_d = append(_d, _ab)
	}
	_a.Slice(_d, func(_cf, _bc int) bool { return _d[_cf].SegmentNumber < _d[_bc].SegmentNumber })
	return &_fa.Globals{Segments: _d}
}
func DecodeGlobals(encoded []byte) (Globals, error) {
	const _fg = "\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073"
	_e := _c.NewReader(encoded)
	_bb, _fe := _fa.DecodeDocument(_e, nil)
	if _fe != nil {
		return nil, _g.Wrap(_fe, _fg, "")
	}
	if _bb.GlobalSegments == nil || (_bb.GlobalSegments.Segments == nil) {
		return nil, _g.Error(_fg, "\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ca := Globals{}
	for _, _ga := range _bb.GlobalSegments.Segments {
		_ca[int(_ga.SegmentNumber)] = _ga
	}
	return _ca, nil
}
