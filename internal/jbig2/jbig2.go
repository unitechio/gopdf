package jbig2

import (
	_d "sort"

	_g "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_e "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder"
	_a "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document"
	_c "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document/segments"
	_ca "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func DecodeGlobals(encoded []byte) (Globals, error) {
	const _eb = "\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073"
	_bc := _g.NewReader(encoded)
	_dg, _fe := _a.DecodeDocument(_bc, nil)
	if _fe != nil {
		return nil, _ca.Wrap(_fe, _eb, "")
	}
	if _dg.GlobalSegments == nil || (_dg.GlobalSegments.Segments == nil) {
		return nil, _ca.Error(_eb, "\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_de := Globals{}
	for _, _fd := range _dg.GlobalSegments.Segments {
		_de[int(_fd.SegmentNumber)] = _fd
	}
	return _de, nil
}
func (_fg Globals) ToDocumentGlobals() *_a.Globals {
	if _fg == nil {
		return nil
	}
	_cg := []*_c.Header{}
	for _, _ea := range _fg {
		_cg = append(_cg, _ea)
	}
	_d.Slice(_cg, func(_gd, _cc int) bool { return _cg[_gd].SegmentNumber < _cg[_cc].SegmentNumber })
	return &_a.Globals{Segments: _cg}
}

type Globals map[int]*_c.Header

func DecodeBytes(encoded []byte, parameters _e.Parameters, globals ...Globals) ([]byte, error) {
	var _b Globals
	if len(globals) > 0 {
		_b = globals[0]
	}
	_gc, _gb := _e.Decode(encoded, parameters, _b.ToDocumentGlobals())
	if _gb != nil {
		return nil, _gb
	}
	return _gc.DecodeNextPage()
}
