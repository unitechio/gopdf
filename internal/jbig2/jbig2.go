package jbig2

import (
	_c "sort"

	_ae "unitechio/gopdf/gopdf/internal/bitwise"
	_g "unitechio/gopdf/gopdf/internal/jbig2/decoder"
	_e "unitechio/gopdf/gopdf/internal/jbig2/document"
	_aa "unitechio/gopdf/gopdf/internal/jbig2/document/segments"
	_ag "unitechio/gopdf/gopdf/internal/jbig2/errors"
)

type Globals map[int]*_aa.Header

func DecodeGlobals(encoded []byte) (Globals, error) {
	const _b = "\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073"
	_ee := _ae.NewReader(encoded)
	_ab, _f := _e.DecodeDocument(_ee, nil)
	if _f != nil {
		return nil, _ag.Wrap(_f, _b, "")
	}
	if _ab.GlobalSegments == nil || (_ab.GlobalSegments.Segments == nil) {
		return nil, _ag.Error(_b, "\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ef := Globals{}
	for _, _ad := range _ab.GlobalSegments.Segments {
		_ef[int(_ad.SegmentNumber)] = _ad
	}
	return _ef, nil
}

func (_fg Globals) ToDocumentGlobals() *_e.Globals {
	if _fg == nil {
		return nil
	}
	_agc := []*_aa.Header{}
	for _, _ba := range _fg {
		_agc = append(_agc, _ba)
	}
	_c.Slice(_agc, func(_fgf, _fd int) bool { return _agc[_fgf].SegmentNumber < _agc[_fd].SegmentNumber })
	return &_e.Globals{Segments: _agc}
}

func DecodeBytes(encoded []byte, parameters _g.Parameters, globals ...Globals) ([]byte, error) {
	var _d Globals
	if len(globals) > 0 {
		_d = globals[0]
	}
	_eb, _gb := _g.Decode(encoded, parameters, _d.ToDocumentGlobals())
	if _gb != nil {
		return nil, _gb
	}
	return _eb.DecodeNextPage()
}
