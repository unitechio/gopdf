package extractor

import (
	_gd "bytes"
	_g "errors"
	_gdf "fmt"
	_eda "image/color"
	_ed "io"
	_b "math"
	_c "regexp"
	_f "sort"
	_d "strings"
	_gg "unicode"
	_a "unicode/utf8"

	_ag "bitbucket.org/shenghui0779/gopdf/common"
	_da "bitbucket.org/shenghui0779/gopdf/contentstream"
	_bg "bitbucket.org/shenghui0779/gopdf/core"
	_ab "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_ee "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_bf "bitbucket.org/shenghui0779/gopdf/model"
	_bd "golang.org/x/text/unicode/norm"
	_fb "golang.org/x/xerrors"
)

// GetContentStreamOps returns the contentStreamOps field of `pt`.
func (_gfbc *PageText) GetContentStreamOps() *_da.ContentStreamOperations { return _gfbc._bfg }
func (_dfe *textObject) checkOp(_gadbg *_da.ContentStreamOperation, _gae int, _bgb bool) (_ccec bool, _cag error) {
	if _dfe == nil {
		var _gdc []_bg.PdfObject
		if _gae > 0 {
			_gdc = _gadbg.Params
			if len(_gdc) > _gae {
				_gdc = _gdc[:_gae]
			}
		}
		_ag.Log.Debug("\u0025\u0023q \u006f\u0070\u0065r\u0061\u006e\u0064\u0020out\u0073id\u0065\u0020\u0074\u0065\u0078\u0074\u002e p\u0061\u0072\u0061\u006d\u0073\u003d\u0025+\u0076", _gadbg.Operand, _gdc)
	}
	if _gae >= 0 {
		if len(_gadbg.Params) != _gae {
			if _bgb {
				_cag = _g.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
			}
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _gadbg.Operand, _gae, len(_gadbg.Params), _gadbg.Params)
			return false, _cag
		}
	}
	return true, nil
}

// RenderMode specifies the text rendering mode (Tmode), which determines whether showing text shall cause
// glyph outlines to be  stroked, filled, used as a clipping boundary, or some combination of the three.
// Stroking, filling, and clipping shall have the same effects for a text object as they do for a path object
// (see 8.5.3, "Path-Painting Operators" and 8.5.4, "Clipping Path Operators").
type RenderMode int

func _eagae(_aaafd _bf.PdfRectangle) *ruling {
	return &ruling{_cgef: _beec, _eead: _aaafd.Urx, _bgeb: _aaafd.Lly, _eecc: _aaafd.Ury}
}
func _fafd(_ebab string) string      { _cbad := []rune(_ebab); return string(_cbad[:len(_cbad)-1]) }
func _def(_bce _ee.Point) _ee.Matrix { return _ee.TranslationMatrix(_bce.X, _bce.Y) }
func _cdgf(_gafcf *wordBag, _cdffd int) *textLine {
	_cgcg := _gafcf.firstWord(_cdffd)
	_dbaf := textLine{PdfRectangle: _cgcg.PdfRectangle, _ddef: _cgcg._fbgge, _dcfd: _cgcg._ecgcg}
	_dbaf.pullWord(_gafcf, _cgcg, _cdffd)
	return &_dbaf
}
func (_adbed *textPara) text() string {
	_fefc := new(_gd.Buffer)
	_adbed.writeText(_fefc)
	return _fefc.String()
}
func _gdaf(_fbggc []rulingList) (rulingList, rulingList) {
	var _cbfe rulingList
	for _, _geag := range _fbggc {
		_cbfe = append(_cbfe, _geag...)
	}
	return _cbfe.vertsHorzs()
}

// RangeOffset returns the TextMarks in `ma` that overlap text[start:end] in the extracted text.
// These are tm: `start` <= tm.Offset + len(tm.Text) && tm.Offset < `end` where
// `start` and `end` are offsets in the extracted text.
// NOTE: TextMarks can contain multiple characters. e.g. "ffi" for the ﬃ ligature so the first and
// last elements of the returned TextMarkArray may only partially overlap text[start:end].
func (_gbb *TextMarkArray) RangeOffset(start, end int) (*TextMarkArray, error) {
	if _gbb == nil {
		return nil, _g.New("\u006da\u003d\u003d\u006e\u0069\u006c")
	}
	if end < start {
		return nil, _gdf.Errorf("\u0065\u006e\u0064\u0020\u003c\u0020\u0073\u0074\u0061\u0072\u0074\u002e\u0020\u0052\u0061n\u0067\u0065\u004f\u0066\u0066\u0073\u0065\u0074\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020\u0073\u0074\u0061\u0072t=\u0025\u0064\u0020\u0065\u006e\u0064\u003d\u0025\u0064\u0020", start, end)
	}
	_dbdd := len(_gbb._fceg)
	if _dbdd == 0 {
		return _gbb, nil
	}
	if start < _gbb._fceg[0].Offset {
		start = _gbb._fceg[0].Offset
	}
	if end > _gbb._fceg[_dbdd-1].Offset+1 {
		end = _gbb._fceg[_dbdd-1].Offset + 1
	}
	_bed := _f.Search(_dbdd, func(_cda int) bool { return _gbb._fceg[_cda].Offset+len(_gbb._fceg[_cda].Text)-1 >= start })
	if !(0 <= _bed && _bed < _dbdd) {
		_cadgf := _gdf.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073\u0074\u0061\u0072\u0074\u003d%\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009 \u006c\u0061\u0073\u0074\u003d%\u0076", start, _bed, _dbdd, _gbb._fceg[0], _gbb._fceg[_dbdd-1])
		return nil, _cadgf
	}
	_fdffa := _f.Search(_dbdd, func(_dac int) bool { return _gbb._fceg[_dac].Offset > end-1 })
	if !(0 <= _fdffa && _fdffa < _dbdd) {
		_edb := _gdf.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020r\u0061\u006e\u0067e\u002e\u0020\u0065n\u0064\u003d%\u0064\u0020\u0069\u0045\u006e\u0064=\u0025d \u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d\u0025\u0076", end, _fdffa, _dbdd, _gbb._fceg[0], _gbb._fceg[_dbdd-1])
		return nil, _edb
	}
	if _fdffa <= _bed {
		return nil, _gdf.Errorf("\u0069\u0045\u006e\u0064\u0020\u003c=\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003a\u0020\u0073\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u0065\u006ed\u003d\u0025\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020i\u0045n\u0064\u003d\u0025\u0064", start, end, _bed, _fdffa)
	}
	return &TextMarkArray{_fceg: _gbb._fceg[_bed:_fdffa]}, nil
}
func _acdc(_fefb []TextMark, _fgbc *int) []TextMark {
	_dgecc := _fefb[len(_fefb)-1]
	_caac := []rune(_dgecc.Text)
	if len(_caac) == 1 {
		_fefb = _fefb[:len(_fefb)-1]
		_fcfd := _fefb[len(_fefb)-1]
		*_fgbc = _fcfd.Offset + len(_fcfd.Text)
	} else {
		_eefca := _fafd(_dgecc.Text)
		*_fgbc += len(_eefca) - len(_dgecc.Text)
		_dgecc.Text = _eefca
	}
	return _fefb
}

type event struct {
	_cega  float64
	_fddb  bool
	_dbdfc int
}

func (_befg paraList) xNeighbours(_dcba float64) map[*textPara][]int {
	_agdf := make([]event, 2*len(_befg))
	if _dcba == 0 {
		for _cfadb, _beecb := range _befg {
			_agdf[2*_cfadb] = event{_beecb.Llx, true, _cfadb}
			_agdf[2*_cfadb+1] = event{_beecb.Urx, false, _cfadb}
		}
	} else {
		for _fdbc, _ccbab := range _befg {
			_agdf[2*_fdbc] = event{_ccbab.Llx - _dcba*_ccbab.fontsize(), true, _fdbc}
			_agdf[2*_fdbc+1] = event{_ccbab.Urx + _dcba*_ccbab.fontsize(), false, _fdbc}
		}
	}
	return _befg.eventNeighbours(_agdf)
}
func (_eddg rectRuling) asRuling() (*ruling, bool) {
	_aadg := ruling{_cgef: _eddg._adea, Color: _eddg.Color, _gbgb: _gaea}
	switch _eddg._adea {
	case _beec:
		_aadg._eead = 0.5 * (_eddg.Llx + _eddg.Urx)
		_aadg._bgeb = _eddg.Lly
		_aadg._eecc = _eddg.Ury
		_ebebe, _cfbd := _eddg.checkWidth(_eddg.Llx, _eddg.Urx)
		if !_cfbd {
			if _bded {
				_ag.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067V\u0065\u0072\u0074\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _eddg)
			}
			return nil, false
		}
		_aadg._efgd = _ebebe
	case _ebabd:
		_aadg._eead = 0.5 * (_eddg.Lly + _eddg.Ury)
		_aadg._bgeb = _eddg.Llx
		_aadg._eecc = _eddg.Urx
		_eaade, _gdbcf := _eddg.checkWidth(_eddg.Lly, _eddg.Ury)
		if !_gdbcf {
			if _bded {
				_ag.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067H\u006f\u0072\u007a\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _eddg)
			}
			return nil, false
		}
		_aadg._efgd = _eaade
	default:
		_ag.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _eddg._adea)
		return nil, false
	}
	return &_aadg, true
}
func (_deece *textTable) getComposite(_edbb, _dafaf int) (paraList, _bf.PdfRectangle) {
	_gcga, _bgfdg := _deece._bbfb[_fcbc(_edbb, _dafaf)]
	if _beae {
		_gdf.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0067\u0065\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u002d\u003e\u0025\u0073\u000a", _edbb, _dafaf, _gcga.String())
	}
	if !_bgfdg {
		return nil, _bf.PdfRectangle{}
	}
	return _gcga.parasBBox()
}
func (_ebb *textObject) setTextRenderMode(_cead int) {
	if _ebb == nil {
		return
	}
	_ebb._degf._egea = RenderMode(_cead)
}
func (_ebeb *shapesState) closePath() {
	if _ebeb._ecd {
		_ebeb._cfgd = append(_ebeb._cfgd, _bddc(_ebeb._decg))
		_ebeb._ecd = false
	} else if len(_ebeb._cfgd) == 0 {
		if _bccf {
			_ag.Log.Debug("\u0063\u006c\u006f\u0073eP\u0061\u0074\u0068\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u0020\u0070\u0061t\u0068")
		}
		_ebeb._ecd = false
		return
	}
	_ebeb._cfgd[len(_ebeb._cfgd)-1].close()
	if _bccf {
		_ag.Log.Info("\u0063\u006c\u006f\u0073\u0065\u0050\u0061\u0074\u0068\u003a\u0020\u0025\u0073", _ebeb)
	}
}

// TextMark represents extracted text on a page with information regarding both textual content,
// formatting (font and size) and positioning.
// It is the smallest unit of text on a PDF page, typically a single character.
//
// getBBox() in test_text.go shows how to compute bounding boxes of substrings of extracted text.
// The following code extracts the text on PDF page `page` into `text` then finds the bounding box
// `bbox` of substring `term` in `text`.
//
//     ex, _ := New(page)
//     // handle errors
//     pageText, _, _, err := ex.ExtractPageText()
//     // handle errors
//     text := pageText.Text()
//     textMarks := pageText.Marks()
//
//     	start := strings.Index(text, term)
//      end := start + len(term)
//      spanMarks, err := textMarks.RangeOffset(start, end)
//      // handle errors
//      bbox, ok := spanMarks.BBox()
//      // handle errors
type TextMark struct {

	// Text is the extracted text.
	Text string

	// Original is the text in the PDF. It has not been decoded like `Text`.
	Original string

	// BBox is the bounding box of the text.
	BBox _bf.PdfRectangle

	// Font is the font the text was drawn with.
	Font *_bf.PdfFont

	// FontSize is the font size the text was drawn with.
	FontSize float64

	// Offset is the offset of the start of TextMark.Text in the extracted text. If you do this
	//   text, textMarks := pageText.Text(), pageText.Marks()
	//   marks := textMarks.Elements()
	// then marks[i].Offset is the offset of marks[i].Text in text.
	Offset int

	// Meta is set true for spaces and line breaks that we insert in the extracted text. We insert
	// spaces (line breaks) when we see characters that are over a threshold horizontal (vertical)
	//  distance  apart. See wordJoiner (lineJoiner) in PageText.computeViews().
	Meta bool

	// FillColor is the fill color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	FillColor _eda.Color

	// StrokeColor is the stroke color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	StrokeColor _eda.Color

	// Orientation is the text orientation
	Orientation int

	// DirectObject is the underlying PdfObject (Text Object) that represents the visible texts. This is introduced to get
	// a simple access to the TextObject in case editing or replacment of some text is needed. E.g during redaction.
	DirectObject _bg.PdfObject

	// ObjString is a decoded string operand of a text-showing operator. It has the same value as `Text` attribute except
	// when many glyphs are represented with the same Text Object that contains multiple length string operand in which case
	// ObjString spans more than one character string that falls in different TextMark objects.
	ObjString []string
	Tw        float64
	Th        float64
	Tc        float64
	Index     int
}

// New returns an Extractor instance for extracting content from the input PDF page.
func New(page *_bf.PdfPage) (*Extractor, error) { return NewWithOptions(page, nil) }
func _gcfa(_cgeb *textWord, _aaee float64, _fcfb, _fega rulingList) *wordBag {
	_caed := _feef(_cgeb._ecgcg)
	_agcbc := []*textWord{_cgeb}
	_gddc := wordBag{_bbf: map[int][]*textWord{_caed: _agcbc}, PdfRectangle: _cgeb.PdfRectangle, _gaed: _cgeb._fbgge, _fcgc: _aaee, _fff: _fcfb, _aede: _fega}
	return &_gddc
}
func (_dccf *textWord) absorb(_bdce *textWord) {
	_dccf.PdfRectangle = _effa(_dccf.PdfRectangle, _bdce.PdfRectangle)
	_dccf._gdgg = append(_dccf._gdgg, _bdce._gdgg...)
}
func (_ggffc rulingList) tidied(_aeade string) rulingList {
	_fgcd := _ggffc.removeDuplicates()
	_fgcd.log("\u0075n\u0069\u0071\u0075\u0065\u0073")
	_gdfab := _fgcd.snapToGroups()
	if _gdfab == nil {
		return nil
	}
	_gdfab.sort()
	if _efa {
		_ag.Log.Info("\u0074\u0069\u0064i\u0065\u0064\u003a\u0020\u0025\u0071\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u0075\u006e\u0069\u0071\u0075\u0065\u0073\u003d\u0025\u0064\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0064\u003d\u0025\u0064", _aeade, len(_ggffc), len(_fgcd), len(_gdfab))
	}
	_gdfab.log("\u0063o\u0061\u006c\u0065\u0073\u0063\u0065d")
	return _gdfab
}

var (
	_efedf = map[rune]string{0x0060: "\u0300", 0x02CB: "\u0300", 0x0027: "\u0301", 0x00B4: "\u0301", 0x02B9: "\u0301", 0x02CA: "\u0301", 0x005E: "\u0302", 0x02C6: "\u0302", 0x007E: "\u0303", 0x02DC: "\u0303", 0x00AF: "\u0304", 0x02C9: "\u0304", 0x02D8: "\u0306", 0x02D9: "\u0307", 0x00A8: "\u0308", 0x00B0: "\u030a", 0x02DA: "\u030a", 0x02BA: "\u030b", 0x02DD: "\u030b", 0x02C7: "\u030c", 0x02C8: "\u030d", 0x0022: "\u030e", 0x02BB: "\u0312", 0x02BC: "\u0313", 0x0486: "\u0313", 0x055A: "\u0313", 0x02BD: "\u0314", 0x0485: "\u0314", 0x0559: "\u0314", 0x02D4: "\u031d", 0x02D5: "\u031e", 0x02D6: "\u031f", 0x02D7: "\u0320", 0x02B2: "\u0321", 0x00B8: "\u0327", 0x02CC: "\u0329", 0x02B7: "\u032b", 0x02CD: "\u0331", 0x005F: "\u0332", 0x204E: "\u0359"}
)

func _ffff(_adaa, _ecfa *textPara) bool { return _eagg(_adaa._dcfc, _ecfa._dcfc) }
func (_aced *textObject) setTextLeading(_dcbb float64) {
	if _aced == nil {
		return
	}
	_aced._degf._acfa = _dcbb
}
func (_cdbb *textObject) getFont(_fecf string) (*_bf.PdfFont, error) {
	if _cdbb._ccae._ac != nil {
		_cfca, _gbag := _cdbb.getFontDict(_fecf)
		if _gbag != nil {
			_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0067\u0065\u0074\u0046\u006f\u006e\u0074:\u0020n\u0061m\u0065=\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", _fecf, _gbag.Error())
			return nil, _gbag
		}
		_cdbb._ccae._eb++
		_cbeb, _deac := _cdbb._ccae._ac[_cfca.String()]
		if _deac {
			_cbeb._ffa = _cdbb._ccae._eb
			return _cbeb._bdgeg, nil
		}
	}
	_eab, _ggbdd := _cdbb.getFontDict(_fecf)
	if _ggbdd != nil {
		return nil, _ggbdd
	}
	_aaeg, _ggbdd := _cdbb.getFontDirect(_fecf)
	if _ggbdd != nil {
		return nil, _ggbdd
	}
	if _cdbb._ccae._ac != nil {
		_fgca := fontEntry{_aaeg, _cdbb._ccae._eb}
		if len(_cdbb._ccae._ac) >= _gaeg {
			var _aga []string
			for _baa := range _cdbb._ccae._ac {
				_aga = append(_aga, _baa)
			}
			_f.Slice(_aga, func(_cage, _edd int) bool {
				return _cdbb._ccae._ac[_aga[_cage]]._ffa < _cdbb._ccae._ac[_aga[_edd]]._ffa
			})
			delete(_cdbb._ccae._ac, _aga[0])
		}
		_cdbb._ccae._ac[_eab.String()] = _fgca
	}
	return _aaeg, nil
}

// String returns a description of `k`.
func (_effbd markKind) String() string {
	_ecac, _aegfb := _gecc[_effbd]
	if !_aegfb {
		return _gdf.Sprintf("\u004e\u006f\u0074\u0020\u0061\u0020\u006d\u0061\u0072k\u003a\u0020\u0025\u0064", _effbd)
	}
	return _ecac
}

const (
	_badf  = true
	_bccb  = true
	_bfe   = true
	_gfea  = false
	_effe  = false
	_afce  = 6
	_ecag  = 3.0
	_cacb  = 200
	_bagf  = true
	_bfcda = true
	_ecec  = true
	_dcaed = true
	_cdc   = false
)

func _acgf(_fbcd, _baae bounded) float64 {
	_aafe := _dagb(_fbcd, _baae)
	if !_fbga(_aafe) {
		return _aafe
	}
	return _addb(_fbcd, _baae)
}
func (_fbccg paraList) sortReadingOrder() {
	_ag.Log.Trace("\u0073\u006fr\u0074\u0052\u0065\u0061\u0064i\u006e\u0067\u004f\u0072\u0064e\u0072\u003a\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_fbccg))
	if len(_fbccg) <= 1 {
		return
	}
	_fbccg.computeEBBoxes()
	_f.Slice(_fbccg, func(_adgge, _fegg int) bool { return _acgf(_fbccg[_adgge], _fbccg[_fegg]) <= 0 })
	_ebeg := _fbccg.topoOrder()
	_fbccg.reorder(_ebeg)
}

// String returns a human readable description of `path`.
func (_dbcc *subpath) String() string {
	_cfba := _dbcc._fcdc
	_aaae := len(_cfba)
	if _aaae <= 5 {
		return _gdf.Sprintf("\u0025d\u003a\u0020\u0025\u0036\u002e\u0032f", _aaae, _cfba)
	}
	return _gdf.Sprintf("\u0025d\u003a\u0020\u0025\u0036.\u0032\u0066\u0020\u0025\u0036.\u0032f\u0020.\u002e\u002e\u0020\u0025\u0036\u002e\u0032f", _aaae, _cfba[0], _cfba[1], _cfba[_aaae-1])
}

// String returns a string describing the current state of the textState stack.
func (_efdc *stateStack) String() string {
	_accb := []string{_gdf.Sprintf("\u002d\u002d\u002d\u002d f\u006f\u006e\u0074\u0020\u0073\u0074\u0061\u0063\u006b\u003a\u0020\u0025\u0064", len(*_efdc))}
	for _cedb, _aaaf := range *_efdc {
		_edef := "\u003c\u006e\u0069l\u003e"
		if _aaaf != nil {
			_edef = _aaaf.String()
		}
		_accb = append(_accb, _gdf.Sprintf("\u0009\u0025\u0032\u0064\u003a\u0020\u0025\u0073", _cedb, _edef))
	}
	return _d.Join(_accb, "\u000a")
}
func _aeca(_eegb, _dcef _ee.Point, _caeb _eda.Color) (*ruling, bool) {
	_edgeg := lineRuling{_abec: _eegb, _eacb: _dcef, _aedc: _bega(_eegb, _dcef), Color: _caeb}
	if _edgeg._aedc == _gdbb {
		return nil, false
	}
	return _edgeg.asRuling()
}
func (_fbffe rulingList) augmentGrid() (rulingList, rulingList) {
	_fcgce, _fceb := _fbffe.vertsHorzs()
	if len(_fcgce) == 0 || len(_fceb) == 0 {
		return _fcgce, _fceb
	}
	_cbf, _egeg := _fcgce, _fceb
	_edbce := _fcgce.bbox()
	_aecce := _fceb.bbox()
	if _efa {
		_ag.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0056\u003d\u0025\u0036\u002e\u0032\u0066", _edbce)
		_ag.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0048\u003d\u0025\u0036\u002e\u0032\u0066", _aecce)
	}
	var _ggbf, _cabb, _geeac, _gecee *ruling
	if _aecce.Llx < _edbce.Llx-_deff {
		_ggbf = &ruling{_gbgb: _eddbd, _cgef: _beec, _eead: _aecce.Llx, _bgeb: _edbce.Lly, _eecc: _edbce.Ury}
		_fcgce = append(rulingList{_ggbf}, _fcgce...)
	}
	if _aecce.Urx > _edbce.Urx+_deff {
		_cabb = &ruling{_gbgb: _eddbd, _cgef: _beec, _eead: _aecce.Urx, _bgeb: _edbce.Lly, _eecc: _edbce.Ury}
		_fcgce = append(_fcgce, _cabb)
	}
	if _edbce.Lly < _aecce.Lly-_deff {
		_geeac = &ruling{_gbgb: _eddbd, _cgef: _ebabd, _eead: _edbce.Lly, _bgeb: _aecce.Llx, _eecc: _aecce.Urx}
		_fceb = append(rulingList{_geeac}, _fceb...)
	}
	if _edbce.Ury > _aecce.Ury+_deff {
		_gecee = &ruling{_gbgb: _eddbd, _cgef: _ebabd, _eead: _edbce.Ury, _bgeb: _aecce.Llx, _eecc: _aecce.Urx}
		_fceb = append(_fceb, _gecee)
	}
	if len(_fcgce)+len(_fceb) == len(_fbffe) {
		return _cbf, _egeg
	}
	_aafa := append(_fcgce, _fceb...)
	_fbffe.log("u\u006e\u0061\u0075\u0067\u006d\u0065\u006e\u0074\u0065\u0064")
	_aafa.log("\u0061u\u0067\u006d\u0065\u006e\u0074\u0065d")
	return _fcgce, _fceb
}
func (_aecg paraList) tables() []TextTable {
	var _bcfa []TextTable
	if _beae {
		_ag.Log.Info("\u0070\u0061\u0072\u0061\u0073\u002e\u0074\u0061\u0062\u006c\u0065\u0073\u003a")
	}
	for _, _cagee := range _aecg {
		_gedg := _cagee._bccg
		if _gedg != nil && _gedg.isExportable() {
			_bcfa = append(_bcfa, _gedg.toTextTable())
		}
	}
	return _bcfa
}
func (_gegd rulingList) isActualGrid() (rulingList, bool) {
	_gdbfe, _gdfgb := _gegd.augmentGrid()
	if !(len(_gdbfe) >= _agec+1 && len(_gdfgb) >= _bafa+1) {
		if _efa {
			_ag.Log.Info("\u0069s\u0041\u0063t\u0075\u0061\u006c\u0047r\u0069\u0064\u003a \u004e\u006f\u0074\u0020\u0061\u006c\u0069\u0067\u006eed\u002e\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u003c\u0020\u0025d\u0020\u0078 \u0025\u0064", len(_gdbfe), len(_gdfgb), _agec+1, _bafa+1)
		}
		return nil, false
	}
	if _efa {
		_ag.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075a\u006c\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025t\u0020\u0026\u0020\u0025\u0074\u0020\u2192 \u0025\u0074", _gegd, len(_gdbfe) >= 2, len(_gdfgb) >= 2, len(_gdbfe) >= 2 && len(_gdfgb) >= 2)
		for _ebdc, _cagc := range _gegd {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0076\u000a", _ebdc, _cagc)
		}
	}
	if _cdc {
		_cbfb, _fafa := _gdbfe[0], _gdbfe[len(_gdbfe)-1]
		_gcfb, _fdbee := _gdfgb[0], _gdfgb[len(_gdfgb)-1]
		if !(_geccf(_cbfb._eead-_gcfb._bgeb) && _geccf(_fafa._eead-_gcfb._eecc) && _geccf(_gcfb._eead-_cbfb._eecc) && _geccf(_fdbee._eead-_cbfb._bgeb)) {
			if _efa {
				_ag.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069d\u003a\u0020\u0020N\u006f\u0074 \u0061\u006c\u0069\u0067\u006e\u0065d\u002e\n\t\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0076\u0031\u003d\u0025\u0073\u000a\u0009\u0068\u0030\u003d\u0025\u0073\u000a\u0009\u0068\u0031\u003d\u0025\u0073", _cbfb, _fafa, _gcfb, _fdbee)
			}
			return nil, false
		}
	} else {
		if !_gdbfe.aligned() {
			if _cecg {
				_ag.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0076\u0065\u0072\u0074\u0073\u002e\u0020%\u0064", len(_gdbfe))
			}
			return nil, false
		}
		if !_gdfgb.aligned() {
			if _efa {
				_ag.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0068\u006f\u0072\u007a\u0073\u002e\u0020%\u0064", len(_gdfgb))
			}
			return nil, false
		}
	}
	_ccegg := append(_gdbfe, _gdfgb...)
	return _ccegg, true
}
func _ddce(_ggce []pathSection) rulingList {
	_bcff(_ggce)
	if _efa {
		_ag.Log.Info("\u006da\u006b\u0065\u0046\u0069l\u006c\u0052\u0075\u006c\u0069n\u0067s\u003a \u0025\u0064\u0020\u0066\u0069\u006c\u006cs", len(_ggce))
	}
	var _feded rulingList
	for _, _eadc := range _ggce {
		for _, _efba := range _eadc._gdbf {
			if !_efba.isQuadrilateral() {
				if _efa {
					_ag.Log.Error("!\u0069s\u0051\u0075\u0061\u0064\u0072\u0069\u006c\u0061t\u0065\u0072\u0061\u006c: \u0025\u0073", _efba)
				}
				continue
			}
			if _affb, _efbae := _efba.makeRectRuling(_eadc.Color); _efbae {
				_feded = append(_feded, _affb)
			} else {
				if _bded {
					_ag.Log.Error("\u0021\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0073", _efba)
				}
			}
		}
	}
	if _efa {
		_ag.Log.Info("\u006d\u0061\u006b\u0065Fi\u006c\u006c\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0073", _feded.String())
	}
	return _feded
}
func _afed(_cbdg, _dedg float64) string {
	_cgee := !_fbga(_cbdg - _dedg)
	if _cgee {
		return "\u000a"
	}
	return "\u0020"
}
func (_dgdg *shapesState) newSubPath() {
	_dgdg.clearPath()
	if _bccf {
		_ag.Log.Info("\u006e\u0065\u0077\u0053\u0075\u0062\u0050\u0061\u0074h\u003a\u0020\u0025\u0073", _dgdg)
	}
}

var _gf = false

func _fcbc(_fagegf, _fdbb int) uint64 { return uint64(_fagegf)*0x1000000 + uint64(_fdbb) }

// ToText returns the page text as a single string.
// Deprecated: This function is deprecated and will be removed in a future major version. Please use
// Text() instead.
func (_egd PageText) ToText() string { return _egd.Text() }

// ExtractTextWithStats works like ExtractText but returns the number of characters in the output
// (`numChars`) and the number of characters that were not decoded (`numMisses`).
func (_bbd *Extractor) ExtractTextWithStats() (_fbf string, _fcf int, _cadd int, _ceag error) {
	_ebd, _fcf, _cadd, _ceag := _bbd.ExtractPageText()
	if _ceag != nil {
		return "", _fcf, _cadd, _ceag
	}
	return _ebd.Text(), _fcf, _cadd, nil
}
func (_ecee *wordBag) highestWord(_bff int, _eaaf, _agf float64) *textWord {
	for _, _ged := range _ecee._bbf[_bff] {
		if _eaaf <= _ged._ecgcg && _ged._ecgcg <= _agf {
			return _ged
		}
	}
	return nil
}

type gridTiling struct {
	_bf.PdfRectangle
	_fdeac []float64
	_gccf  []float64
	_ffdd  map[float64]map[float64]gridTile
}

func _gbae(_eacbe string) (string, bool) {
	_ggefb := []rune(_eacbe)
	if len(_ggefb) != 1 {
		return "", false
	}
	_edag, _bfbg := _efedf[_ggefb[0]]
	return _edag, _bfbg
}
func _cacbb(_febe []float64, _fcee, _abaf float64) []float64 {
	_acce, _cdda := _fcee, _abaf
	if _cdda < _acce {
		_acce, _cdda = _cdda, _acce
	}
	_gfcf := make([]float64, 0, len(_febe)+2)
	_gfcf = append(_gfcf, _fcee)
	for _, _gedeb := range _febe {
		if _gedeb <= _acce {
			continue
		} else if _gedeb >= _cdda {
			break
		}
		_gfcf = append(_gfcf, _gedeb)
	}
	_gfcf = append(_gfcf, _abaf)
	return _gfcf
}
func (_gbc *textMark) inDiacriticArea(_agbg *textMark) bool {
	_acd := _gbc.Llx - _agbg.Llx
	_agdg := _gbc.Urx - _agbg.Urx
	_abbde := _gbc.Lly - _agbg.Lly
	return _b.Abs(_acd+_agdg) < _gbc.Width()*_fge && _b.Abs(_abbde) < _gbc.Height()*_fge
}
func _gafce(_cfbba _bf.PdfRectangle, _ededf, _gaec, _fccab, _gbfe *ruling) gridTile {
	_bbgag := _cfbba.Llx
	_cgba := _cfbba.Urx
	_accbf := _cfbba.Lly
	_bcag := _cfbba.Ury
	return gridTile{PdfRectangle: _cfbba, _bceg: _ededf != nil && _ededf.encloses(_accbf, _bcag), _abdfe: _gaec != nil && _gaec.encloses(_accbf, _bcag), _abag: _fccab != nil && _fccab.encloses(_bbgag, _cgba), _gfga: _gbfe != nil && _gbfe.encloses(_bbgag, _cgba)}
}
func _dgcd(_ffgbd, _cdea float64) bool { return _ffgbd/_b.Max(_gbef, _cdea) < _cdfc }

// ExtractPageText returns the text contents of `e` (an Extractor for a page) as a PageText.
// TODO(peterwilliams97): The stats complicate this function signature and aren't very useful.
//                        Replace with a function like Extract() (*PageText, error)
func (_ggba *Extractor) ExtractPageText() (*PageText, int, int, error) {
	_cedc, _dafd, _cae, _adg := _ggba.extractPageText(_ggba._fc, _ggba._ad, _ee.IdentityMatrix(), 0)
	if _adg != nil && _adg != _bf.ErrColorOutOfRange {
		return nil, 0, 0, _adg
	}
	_cedc.computeViews()
	_adg = _eabb(_cedc)
	if _adg != nil {
		return nil, 0, 0, _adg
	}
	if _ggba._bc != nil {
		if _ggba._bc.ApplyCropBox && _ggba._ga != nil {
			_cedc.ApplyArea(*_ggba._ga)
		}
	}
	return _cedc, _dafd, _cae, nil
}
func (_cgfb intSet) has(_egaa int) bool {
	_, _egeb := _cgfb[_egaa]
	return _egeb
}
func _eabb(_cgdaa *PageText) error {
	return nil
}
func (_baba *textObject) newTextMark(_dfa string, _bffc _ee.Matrix, _ddea _ee.Point, _gafa float64, _cbgb *_bf.PdfFont, _bcea float64, _caefc, _ffe _eda.Color, _gddb _bg.PdfObject, _cggcb []string, _bcbcc int) (textMark, bool) {
	_eaga := _bffc.Angle()
	_aada := _fccec(_eaga, _ggdc)
	var _fccc float64
	if _aada%180 != 90 {
		_fccc = _bffc.ScalingFactorY()
	} else {
		_fccc = _bffc.ScalingFactorX()
	}
	_efbd := _ecc(_bffc)
	_acaa := _bf.PdfRectangle{Llx: _efbd.X, Lly: _efbd.Y, Urx: _ddea.X, Ury: _ddea.Y}
	switch _aada % 360 {
	case 90:
		_acaa.Urx -= _fccc
	case 180:
		_acaa.Ury -= _fccc
	case 270:
		_acaa.Urx += _fccc
	case 0:
		_acaa.Ury += _fccc
	default:
		_aada = 0
		_acaa.Ury += _fccc
	}
	if _acaa.Llx > _acaa.Urx {
		_acaa.Llx, _acaa.Urx = _acaa.Urx, _acaa.Llx
	}
	if _acaa.Lly > _acaa.Ury {
		_acaa.Lly, _acaa.Ury = _acaa.Ury, _acaa.Lly
	}
	_afgb := true
	if _baba._ccae._bb.Width() > 0 {
		_dfb, _cebb := _agcbf(_acaa, _baba._ccae._bb)
		if !_cebb {
			_afgb = false
			_ag.Log.Debug("\u0054\u0065\u0078\u0074\u0020m\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0070a\u0067\u0065\u002e\u0020\u0062\u0062\u006f\u0078\u003d\u0025\u0067\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u0067\u0020\u0074\u0065\u0078\u0074\u003d\u0025q", _acaa, _baba._ccae._bb, _dfa)
		}
		_acaa = _dfb
	}
	_dffb := _acaa
	_fcdd := _baba._ccae._bb
	switch _aada % 360 {
	case 90:
		_fcdd.Urx, _fcdd.Ury = _fcdd.Ury, _fcdd.Urx
		_dffb = _bf.PdfRectangle{Llx: _fcdd.Urx - _acaa.Ury, Urx: _fcdd.Urx - _acaa.Lly, Lly: _acaa.Llx, Ury: _acaa.Urx}
	case 180:
		_dffb = _bf.PdfRectangle{Llx: _fcdd.Urx - _acaa.Llx, Urx: _fcdd.Urx - _acaa.Urx, Lly: _fcdd.Ury - _acaa.Lly, Ury: _fcdd.Ury - _acaa.Ury}
	case 270:
		_fcdd.Urx, _fcdd.Ury = _fcdd.Ury, _fcdd.Urx
		_dffb = _bf.PdfRectangle{Llx: _acaa.Ury, Urx: _acaa.Lly, Lly: _fcdd.Ury - _acaa.Llx, Ury: _fcdd.Ury - _acaa.Urx}
	}
	if _dffb.Llx > _dffb.Urx {
		_dffb.Llx, _dffb.Urx = _dffb.Urx, _dffb.Llx
	}
	if _dffb.Lly > _dffb.Ury {
		_dffb.Lly, _dffb.Ury = _dffb.Ury, _dffb.Lly
	}
	_fdffd := textMark{_eeaf: _dfa, PdfRectangle: _dffb, _bbgf: _acaa, _becff: _cbgb, _abba: _fccc, _dbe: _bcea, _ecbd: _bffc, _fcced: _ddea, _bfbc: _aada, _agbe: _caefc, _dggfa: _ffe, _gded: _gddb, _cedf: _cggcb, Th: _baba._degf._gdef, Tw: _baba._degf._gab, _eced: _bcbcc}
	if _eccd {
		_ag.Log.Info("n\u0065\u0077\u0054\u0065\u0078\u0074M\u0061\u0072\u006b\u003a\u0020\u0073t\u0061\u0072\u0074\u003d\u0025\u002e\u0032f\u0020\u0065\u006e\u0064\u003d\u0025\u002e\u0032\u0066\u0020%\u0073", _efbd, _ddea, _fdffd.String())
	}
	return _fdffd, _afgb
}
func (_dfdad *textPara) isAtom() *textTable {
	_abbda := _dfdad
	_dbcb := _dfdad._bfdc
	_agdd := _dfdad._egec
	if _dbcb.taken() || _agdd.taken() {
		return nil
	}
	_agcba := _dbcb._egec
	if _agcba.taken() || _agcba != _agdd._bfdc {
		return nil
	}
	return _fffd(_abbda, _dbcb, _agdd, _agcba)
}
func _dcdd(_fadb, _affbb _ee.Point) bool {
	_aece := _b.Abs(_fadb.X - _affbb.X)
	_bdec := _b.Abs(_fadb.Y - _affbb.Y)
	return _dgcd(_bdec, _aece)
}

var _fbedc = map[rulingKind]string{_gdbb: "\u006e\u006f\u006e\u0065", _ebabd: "\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _beec: "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c"}

func (_acdde gridTiling) complete() bool {
	for _, _aaac := range _acdde._ffdd {
		for _, _dbeb := range _aaac {
			if !_dbeb.complete() {
				return false
			}
		}
	}
	return true
}
func (_bdaf rulingList) mergePrimary() float64 {
	_ggcc := _bdaf[0]._eead
	for _, _fbdg := range _bdaf[1:] {
		_ggcc += _fbdg._eead
	}
	return _ggcc / float64(len(_bdaf))
}
func (_fcdda paraList) toTextMarks() []TextMark {
	_adggb := 0
	var _aabe []TextMark
	for _gbbf, _gddfe := range _fcdda {
		if _gddfe._dgggff {
			continue
		}
		_gabc := _gddfe.toTextMarks(&_adggb)
		_aabe = append(_aabe, _gabc...)
		if _gbbf != len(_fcdda)-1 {
			if _ceabf(_gddfe, _fcdda[_gbbf+1]) {
				_aabe = _abbf(_aabe, &_adggb, "\u0020")
			} else {
				_aabe = _abbf(_aabe, &_adggb, "\u000a")
				_aabe = _abbf(_aabe, &_adggb, "\u000a")
			}
		}
	}
	_aabe = _abbf(_aabe, &_adggb, "\u000a")
	_aabe = _abbf(_aabe, &_adggb, "\u000a")
	return _aabe
}

type rectRuling struct {
	_adea  rulingKind
	_febad markKind
	_eda.Color
	_bf.PdfRectangle
}

func (_bafe *shapesState) addPoint(_bgf, _fgff float64) {
	_eegdc := _bafe.establishSubpath()
	_cfe := _bafe.devicePoint(_bgf, _fgff)
	if _eegdc == nil {
		_bafe._ecd = true
		_bafe._decg = _cfe
	} else {
		_eegdc.add(_cfe)
	}
}
func (_cdad *shapesState) fill(_gdbg *[]pathSection) {
	_dgc := pathSection{_gdbf: _cdad._cfgd, Color: _cdad._gdcg.getFillColor()}
	*_gdbg = append(*_gdbg, _dgc)
	if _efa {
		_ecfd := _dgc.bbox()
		_gdf.Printf("\u0020 \u0020\u0020\u0046\u0049\u004c\u004c\u003a %\u0032\u0064\u0020\u0066\u0069\u006c\u006c\u0073\u0020\u0028\u0025\u0064\u0020\u006ee\u0077\u0029 \u0073\u0073\u003d%\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d\u0025\u0033\u0076\u0020\u0025\u0036\u002e\u0032f\u003d\u00256.\u0032\u0066\u0078%\u0036\u002e\u0032\u0066\u000a", len(*_gdbg), len(_dgc._gdbf), _cdad, _dgc.Color, _ecfd, _ecfd.Width(), _ecfd.Height())
		if _bfga {
			for _egcfb, _cddg := range _dgc._gdbf {
				_gdf.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _egcfb, _cddg)
				if _egcfb == 10 {
					break
				}
			}
		}
	}
}
func (_deb *shapesState) stroke(_fbgg *[]pathSection) {
	_ccd := pathSection{_gdbf: _deb._cfgd, Color: _deb._gdcg.getStrokeColor()}
	*_fbgg = append(*_fbgg, _ccd)
	if _efa {
		_gdf.Printf("\u0020 \u0020\u0020S\u0054\u0052\u004fK\u0045\u003a\u0020\u0025\u0064\u0020\u0073t\u0072\u006f\u006b\u0065\u0073\u0020s\u0073\u003d\u0025\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d%\u002b\u0076\u0020\u0025\u0036\u002e\u0032\u0066\u000a", len(*_fbgg), _deb, _deb._gdcg.getStrokeColor(), _ccd.bbox())
		if _bfga {
			for _eggd, _bec := range _deb._cfgd {
				_gdf.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _eggd, _bec)
				if _eggd == 10 {
					break
				}
			}
		}
	}
}
func _addb(_cceg, _gafb bounded) float64              { return _cceg.bbox().Llx - _gafb.bbox().Llx }
func (_cadg *textObject) moveText(_fgf, _dda float64) { _cadg.moveLP(_fgf, _dda) }

const (
	_dg  = "\u0045\u0052R\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
	_edf = "\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043a\u006e\u0027\u0074 g\u0065\u0074\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002c\u0020\u0066\u006fn\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064"
	_cgc = "\u0045\u0052\u0052O\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
)

func _bcff(_fece []pathSection) {
	if _cegd < 0.0 {
		return
	}
	if _efa {
		_ag.Log.Info("\u0067\u0072\u0061\u006e\u0075\u006c\u0061\u0072\u0069\u007a\u0065\u003a\u0020\u0025\u0064 \u0073u\u0062\u0070\u0061\u0074\u0068\u0020\u0073\u0065\u0063\u0074\u0069\u006f\u006e\u0073", len(_fece))
	}
	for _cbcgb, _fgef := range _fece {
		for _cgbc, _fgbd := range _fgef._gdbf {
			for _adga, _aaeec := range _fgbd._fcdc {
				_fgbd._fcdc[_adga] = _ee.Point{X: _ccede(_aaeec.X), Y: _ccede(_aaeec.Y)}
				if _efa {
					_fbbd := _fgbd._fcdc[_adga]
					if !_egag(_aaeec, _fbbd) {
						_affae := _ee.Point{X: _fbbd.X - _aaeec.X, Y: _fbbd.Y - _aaeec.Y}
						_gdf.Printf("\u0025\u0034d \u002d\u0020\u00254\u0064\u0020\u002d\u0020%4d\u003a %\u002e\u0032\u0066\u0020\u2192\u0020\u0025.2\u0066\u0020\u0028\u0025\u0067\u0029\u000a", _cbcgb, _cgbc, _adga, _aaeec, _fbbd, _affae)
					}
				}
			}
		}
	}
}
func (_ebe *imageExtractContext) extractXObjectImage(_acf *_bg.PdfObjectName, _eca _da.GraphicsState, _gdfa *_bf.PdfPageResources) error {
	_eba, _ := _gdfa.GetXObjectByName(*_acf)
	if _eba == nil {
		return nil
	}
	_gfg, _deg := _ebe._gda[_eba]
	if !_deg {
		_ega, _aaca := _gdfa.GetXObjectImageByName(*_acf)
		if _aaca != nil {
			return _aaca
		}
		if _ega == nil {
			return nil
		}
		_abf, _aaca := _ega.ToImage()
		if _aaca != nil {
			return _aaca
		}
		_gfg = &cachedImage{_fbc: _abf, _aac: _ega.ColorSpace}
		_ebe._gda[_eba] = _gfg
	}
	_cea := _gfg._fbc
	_egg := _gfg._aac
	_egf, _ggb := _egg.ImageToRGB(*_cea)
	if _ggb != nil {
		return _ggb
	}
	_ag.Log.Debug("@\u0044\u006f\u0020\u0043\u0054\u004d\u003a\u0020\u0025\u0073", _eca.CTM.String())
	_ede := ImageMark{Image: &_egf, Width: _eca.CTM.ScalingFactorX(), Height: _eca.CTM.ScalingFactorY(), Angle: _eca.CTM.Angle()}
	_ede.X, _ede.Y = _eca.CTM.Translation()
	_ebe._dba = append(_ebe._dba, _ede)
	_ebe._fe++
	return nil
}
func (_cedca *shapesState) lastpointEstablished() (_ee.Point, bool) {
	if _cedca._ecd {
		return _cedca._decg, false
	}
	_cacf := len(_cedca._cfgd)
	if _cacf > 0 && _cedca._cfgd[_cacf-1]._fcfe {
		return _cedca._cfgd[_cacf-1].last(), false
	}
	return _ee.Point{}, true
}
func (_gafd paraList) inTile(_cdaf gridTile) paraList {
	var _fdd paraList
	for _, _cbdbc := range _gafd {
		if _cdaf.contains(_cbdbc.PdfRectangle) {
			_fdd = append(_fdd, _cbdbc)
		}
	}
	if _beae {
		_gdf.Printf("\u0020 \u0020\u0069\u006e\u0054i\u006c\u0065\u003a\u0020\u0020%\u0073 \u0069n\u0073\u0069\u0064\u0065\u003d\u0025\u0064\n", _cdaf, len(_fdd))
		for _gdeg, _cfge := range _fdd {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _gdeg, _cfge)
		}
		_gdf.Println("")
	}
	return _fdd
}

type fontEntry struct {
	_bdgeg *_bf.PdfFont
	_ffa   int64
}

// ImageExtractOptions contains options for controlling image extraction from
// PDF pages.
type ImageExtractOptions struct{ IncludeInlineStencilMasks bool }

func (_adff *wordBag) text() string {
	_bedb := _adff.allWords()
	_gbbb := make([]string, len(_bedb))
	for _efdca, _gefe := range _bedb {
		_gbbb[_efdca] = _gefe._bcaa
	}
	return _d.Join(_gbbb, "\u0020")
}

// Text returns the extracted page text.
func (_ceab PageText) Text() string { return _ceab._gcea }
func (_dgce rulingList) blocks(_bagbc, _edefa *ruling) bool {
	if _bagbc._bgeb > _edefa._eecc || _edefa._bgeb > _bagbc._eecc {
		return false
	}
	_dagc := _b.Max(_bagbc._bgeb, _edefa._bgeb)
	_dddg := _b.Min(_bagbc._eecc, _edefa._eecc)
	if _bagbc._eead > _edefa._eead {
		_bagbc, _edefa = _edefa, _bagbc
	}
	for _, _dddae := range _dgce {
		if _bagbc._eead <= _dddae._eead+_dcfe && _dddae._eead <= _edefa._eead+_dcfe && _dddae._bgeb <= _dddg && _dagc <= _dddae._eecc {
			return true
		}
	}
	return false
}
func (_fageg *ruling) encloses(_gegad, _bbab float64) bool {
	return _fageg._bgeb-_deff <= _gegad && _bbab <= _fageg._eecc+_deff
}

// String returns a description of `k`.
func (_acede rulingKind) String() string {
	_acgb, _gefb := _fbedc[_acede]
	if !_gefb {
		return _gdf.Sprintf("\u004e\u006ft\u0020\u0061\u0020r\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _acede)
	}
	return _acgb
}
func _dgecg(_dbdgb, _afbg bounded) float64 { return _dbdgb.bbox().Llx - _afbg.bbox().Urx }
func (_fgfa compositeCell) split(_bddg, _dedd []float64) *textTable {
	_ecba := len(_bddg) + 1
	_beed := len(_dedd) + 1
	if _beae {
		_ag.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0043\u0065l\u006c\u002e\u0073\u0070l\u0069\u0074\u003a\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a\u0009\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025\u0073\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073=\u0025\u0036\u002e\u0032\u0066\u000a\t\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d%\u0036\u002e\u0032\u0066", _beed, _ecba, _fgfa, _bddg, _dedd)
		_gdf.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073\u000a", len(_fgfa.paraList))
		for _bagbe, _cbefc := range _fgfa.paraList {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bagbe, _cbefc.String())
		}
		_gdf.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", len(_fgfa.lines()))
		for _bgaab, _gecab := range _fgfa.lines() {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bgaab, _gecab)
		}
	}
	_bddg = _cacbb(_bddg, _fgfa.Ury, _fgfa.Lly)
	_dedd = _cacbb(_dedd, _fgfa.Llx, _fgfa.Urx)
	_egcdc := make(map[uint64]*textPara, _beed*_ecba)
	_cbebe := textTable{_baee: _beed, _cabfg: _ecba, _aagc: _egcdc}
	_bacg := _fgfa.paraList
	_f.Slice(_bacg, func(_ffafd, _ceafd int) bool {
		_gefeb, _fgcg := _bacg[_ffafd], _bacg[_ceafd]
		_ffbf, _dbda := _gefeb.Lly, _fgcg.Lly
		if _ffbf != _dbda {
			return _ffbf < _dbda
		}
		return _gefeb.Llx < _fgcg.Llx
	})
	_dgf := make(map[uint64]_bf.PdfRectangle, _beed*_ecba)
	for _caaf, _debed := range _bddg[1:] {
		_cagf := _bddg[_caaf]
		for _ecbg, _dbce := range _dedd[1:] {
			_cdba := _dedd[_ecbg]
			_dgf[_fcbc(_ecbg, _caaf)] = _bf.PdfRectangle{Llx: _cdba, Urx: _dbce, Lly: _debed, Ury: _cagf}
		}
	}
	if _beae {
		_ag.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0043\u0065l\u006c\u002e\u0073\u0070\u006c\u0069\u0074\u003a\u0020\u0072e\u0063\u0074\u0073")
		_gdf.Printf("\u0020\u0020\u0020\u0020")
		for _eaad := 0; _eaad < _beed; _eaad++ {
			_gdf.Printf("\u0025\u0033\u0030\u0064\u002c\u0020", _eaad)
		}
		_gdf.Println()
		for _eged := 0; _eged < _ecba; _eged++ {
			_gdf.Printf("\u0020\u0020\u0025\u0032\u0064\u003a", _eged)
			for _beeb := 0; _beeb < _beed; _beeb++ {
				_gdf.Printf("\u00256\u002e\u0032\u0066\u002c\u0020", _dgf[_fcbc(_beeb, _eged)])
			}
			_gdf.Println()
		}
	}
	_cggcc := func(_effb *textLine) (int, int) {
		for _aee := 0; _aee < _ecba; _aee++ {
			for _cbed := 0; _cbed < _beed; _cbed++ {
				if _ceba(_dgf[_fcbc(_cbed, _aee)], _effb.PdfRectangle) {
					return _cbed, _aee
				}
			}
		}
		return -1, -1
	}
	_bcbe := make(map[uint64][]*textLine, _beed*_ecba)
	for _, _ceagg := range _bacg.lines() {
		_feff, _gdbc := _cggcc(_ceagg)
		if _feff < 0 {
			continue
		}
		_bcbe[_fcbc(_feff, _gdbc)] = append(_bcbe[_fcbc(_feff, _gdbc)], _ceagg)
	}
	for _bade := 0; _bade < len(_bddg)-1; _bade++ {
		_fgcb := _bddg[_bade]
		_aedb := _bddg[_bade+1]
		for _edbd := 0; _edbd < len(_dedd)-1; _edbd++ {
			_gdbea := _dedd[_edbd]
			_bcdc := _dedd[_edbd+1]
			_aefg := _bf.PdfRectangle{Llx: _gdbea, Urx: _bcdc, Lly: _aedb, Ury: _fgcb}
			_abgd := _bcbe[_fcbc(_edbd, _bade)]
			if len(_abgd) == 0 {
				continue
			}
			_gecb := _egecc(_aefg, _abgd)
			_cbebe.put(_edbd, _bade, _gecb)
		}
	}
	return &_cbebe
}
func (_fecb *subpath) makeRectRuling(_dggcb _eda.Color) (*ruling, bool) {
	if _bded {
		_ag.Log.Info("\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0070\u0061\u0074\u0068\u003d\u0025\u0076", _fecb)
	}
	_ebed := _fecb._fcdc[:4]
	_ccddb := make(map[int]rulingKind, len(_ebed))
	for _ddda, _dga := range _ebed {
		_fbbf := _fecb._fcdc[(_ddda+1)%4]
		_ccddb[_ddda] = _ddedb(_dga, _fbbf)
		if _bded {
			_gdf.Printf("\u0025\u0034\u0064: \u0025\u0073\u0020\u003d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066", _ddda, _ccddb[_ddda], _dga, _fbbf)
		}
	}
	if _bded {
		_gdf.Printf("\u0020\u0020\u0020\u006b\u0069\u006e\u0064\u0073\u003d\u0025\u002b\u0076\u000a", _ccddb)
	}
	var _ecca, _dfaa []int
	for _bdeed, _efaga := range _ccddb {
		switch _efaga {
		case _ebabd:
			_dfaa = append(_dfaa, _bdeed)
		case _beec:
			_ecca = append(_ecca, _bdeed)
		}
	}
	if _bded {
		_gdf.Printf("\u0020\u0020 \u0068\u006f\u0072z\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_dfaa), _dfaa)
		_gdf.Printf("\u0020\u0020 \u0076\u0065\u0072t\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_ecca), _ecca)
	}
	_afad := (len(_dfaa) == 2 && len(_ecca) == 2) || (len(_dfaa) == 2 && len(_ecca) == 0 && _dcdd(_ebed[_dfaa[0]], _ebed[_dfaa[1]])) || (len(_ecca) == 2 && len(_dfaa) == 0 && _cada(_ebed[_ecca[0]], _ebed[_ecca[1]]))
	if _bded {
		_gdf.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_dfaa), len(_ecca), _afad)
	}
	if !_afad {
		if _bded {
			_ag.Log.Error("\u0021!\u006d\u0061\u006b\u0065R\u0065\u0063\u0074\u0052\u0075l\u0069n\u0067:\u0020\u0070\u0061\u0074\u0068\u003d\u0025v", _fecb)
			_gdf.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_dfaa), len(_ecca), _afad)
		}
		return &ruling{}, false
	}
	if len(_ecca) == 0 {
		for _fbegc, _afdf := range _ccddb {
			if _afdf != _ebabd {
				_ecca = append(_ecca, _fbegc)
			}
		}
	}
	if len(_dfaa) == 0 {
		for _cbbb, _eeeg := range _ccddb {
			if _eeeg != _beec {
				_dfaa = append(_dfaa, _cbbb)
			}
		}
	}
	if _bded {
		_ag.Log.Info("\u006da\u006b\u0065R\u0065\u0063\u0074\u0052u\u006c\u0069\u006eg\u003a\u0020\u0068\u006f\u0072\u007a\u0073\u003d\u0025d \u0076\u0065\u0072t\u0073\u003d%\u0064\u0020\u0070\u006f\u0069\u006et\u0073\u003d%\u0064\u000a"+"\u0009\u0020\u0068o\u0072\u007a\u0073\u003d\u0025\u002b\u0076\u000a"+"\u0009\u0020\u0076e\u0072\u0074\u0073\u003d\u0025\u002b\u0076\u000a"+"\t\u0070\u006f\u0069\u006e\u0074\u0073\u003d\u0025\u002b\u0076", len(_dfaa), len(_ecca), len(_ebed), _dfaa, _ecca, _ebed)
	}
	var _gccb, _fdaf, _fadd, _fdfd _ee.Point
	if _ebed[_dfaa[0]].Y > _ebed[_dfaa[1]].Y {
		_fadd, _fdfd = _ebed[_dfaa[0]], _ebed[_dfaa[1]]
	} else {
		_fadd, _fdfd = _ebed[_dfaa[1]], _ebed[_dfaa[0]]
	}
	if _ebed[_ecca[0]].X > _ebed[_ecca[1]].X {
		_gccb, _fdaf = _ebed[_ecca[0]], _ebed[_ecca[1]]
	} else {
		_gccb, _fdaf = _ebed[_ecca[1]], _ebed[_ecca[0]]
	}
	_abcf := _bf.PdfRectangle{Llx: _gccb.X, Urx: _fdaf.X, Lly: _fdfd.Y, Ury: _fadd.Y}
	if _abcf.Llx > _abcf.Urx {
		_abcf.Llx, _abcf.Urx = _abcf.Urx, _abcf.Llx
	}
	if _abcf.Lly > _abcf.Ury {
		_abcf.Lly, _abcf.Ury = _abcf.Ury, _abcf.Lly
	}
	_ceea := rectRuling{PdfRectangle: _abcf, _adea: _eceb(_abcf), Color: _dggcb}
	if _ceea._adea == _gdbb {
		if _bded {
			_ag.Log.Error("\u006da\u006b\u0065\u0052\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006eg\u003a\u0020\u006b\u0069\u006e\u0064\u003d\u006e\u0069\u006c")
		}
		return nil, false
	}
	_fdgfc, _aeee := _ceea.asRuling()
	if !_aeee {
		if _bded {
			_ag.Log.Error("\u006da\u006b\u0065\u0052\u0065c\u0074\u0052\u0075\u006c\u0069n\u0067:\u0020!\u0069\u0073\u0052\u0075\u006c\u0069\u006eg")
		}
		return nil, false
	}
	if _efa {
		_gdf.Printf("\u0020\u0020\u0020\u0072\u003d\u0025\u0073\u000a", _fdgfc.String())
	}
	return _fdgfc, true
}
func (_deeg *wordBag) makeRemovals() map[int]map[*textWord]struct{} {
	_cggc := make(map[int]map[*textWord]struct{}, len(_deeg._bbf))
	for _bagb := range _deeg._bbf {
		_cggc[_bagb] = make(map[*textWord]struct{})
	}
	return _cggc
}
func _fffd(_febc, _cged, _fbbfc, _gbfa *textPara) *textTable {
	_eacd := &textTable{_baee: 2, _cabfg: 2, _aagc: make(map[uint64]*textPara, 4)}
	_eacd.put(0, 0, _febc)
	_eacd.put(1, 0, _cged)
	_eacd.put(0, 1, _fbbfc)
	_eacd.put(1, 1, _gbfa)
	return _eacd
}

type compositeCell struct {
	_bf.PdfRectangle
	paraList
}

func (_bdbd *shapesState) moveTo(_aabf, _ada float64) {
	_bdbd._ecd = true
	_bdbd._decg = _bdbd.devicePoint(_aabf, _ada)
	if _bccf {
		_ag.Log.Info("\u006d\u006fv\u0065\u0054\u006f\u003a\u0020\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u003d%.\u0032\u0066", _aabf, _ada, _bdbd._decg)
	}
}
func _babd(_cabgd string, _adafg int) string {
	if len(_cabgd) < _adafg {
		return _cabgd
	}
	return _cabgd[:_adafg]
}
func (_cffce rulingList) asTiling() gridTiling {
	if _efbc {
		_ag.Log.Info("r\u0075\u006ci\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0063s\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u002b\u002b\u002b\u0020\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_cffce))
	}
	for _ecgc, _dcgf := range _cffce[1:] {
		_cbcd := _cffce[_ecgc]
		if _cbcd.alignsPrimary(_dcgf) && _cbcd.alignsSec(_dcgf) {
			_ag.Log.Error("a\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0044\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065 \u0072\u0075\u006c\u0069\u006e\u0067\u0073\u002e\u000a\u0009v=\u0025\u0073\u000a\t\u0077=\u0025\u0073", _dcgf, _cbcd)
		}
	}
	_cffce.sortStrict()
	_cffce.log("\u0073n\u0061\u0070\u0070\u0065\u0064")
	_eggbf, _ebfc := _cffce.vertsHorzs()
	_gfeaa := _eggbf.primaries()
	_ecggc := _ebfc.primaries()
	_edaaf := len(_gfeaa) - 1
	_fefg := len(_ecggc) - 1
	if _edaaf == 0 || _fefg == 0 {
		return gridTiling{}
	}
	_deec := _bf.PdfRectangle{Llx: _gfeaa[0], Urx: _gfeaa[_edaaf], Lly: _ecggc[0], Ury: _ecggc[_fefg]}
	if _efbc {
		_ag.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0072\u0074s=\u0025\u0064", len(_eggbf))
		for _acff, _feaa := range _eggbf {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _acff, _feaa)
		}
		_ag.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0068\u006f\u0072\u007as=\u0025\u0064", len(_ebfc))
		for _bfad, _ccee := range _ebfc {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bfad, _ccee)
		}
		_ag.Log.Info("\u0072\u0075\u006c\u0069\u006eg\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067:\u0020\u0020\u0077\u0078\u0068\u003d\u0025\u0064\u0078\u0025\u0064\u000a\u0009\u006c\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a\u0009\u006c\u006c\u0079\u003d\u0025\u002e\u0032f", _edaaf, _fefg, _gfeaa, _ecggc)
	}
	_fcaf := make([]gridTile, _edaaf*_fefg)
	for _afgf := _fefg - 1; _afgf >= 0; _afgf-- {
		_fefbb := _ecggc[_afgf]
		_dbbde := _ecggc[_afgf+1]
		for _edeb := 0; _edeb < _edaaf; _edeb++ {
			_fdeaa := _gfeaa[_edeb]
			_ddde := _gfeaa[_edeb+1]
			_cddf := _eggbf.findPrimSec(_fdeaa, _fefbb)
			_fcb := _eggbf.findPrimSec(_ddde, _fefbb)
			_dgge := _ebfc.findPrimSec(_fefbb, _fdeaa)
			_caff := _ebfc.findPrimSec(_dbbde, _fdeaa)
			_afcea := _bf.PdfRectangle{Llx: _fdeaa, Urx: _ddde, Lly: _fefbb, Ury: _dbbde}
			_efagg := _gafce(_afcea, _cddf, _fcb, _dgge, _caff)
			_fcaf[_afgf*_edaaf+_edeb] = _efagg
			if _efbc {
				_gdf.Printf("\u0020\u0020\u0078\u003d\u0025\u0032\u0064\u0020\u0079\u003d\u0025\u0032\u0064\u003a\u0020%\u0073 \u0025\u0036\u002e\u0032\u0066\u0020\u0078\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _edeb, _afgf, _efagg.String(), _efagg.Width(), _efagg.Height())
			}
		}
	}
	if _efbc {
		_ag.Log.Info("r\u0075\u006c\u0069\u006e\u0067\u004c\u0069\u0073\u0074.\u0061\u0073\u0054\u0069\u006c\u0069\u006eg:\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0020\u0068\u006f\u0072\u0069zo\u006e\u0074a\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _deec)
	}
	_feefe := make([]map[float64]gridTile, _fefg)
	for _dbbc := _fefg - 1; _dbbc >= 0; _dbbc-- {
		if _efbc {
			_gdf.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _dbbc)
		}
		_feefe[_dbbc] = make(map[float64]gridTile, _edaaf)
		for _ccca := 0; _ccca < _edaaf; _ccca++ {
			_gggbg := _fcaf[_dbbc*_edaaf+_ccca]
			if _efbc {
				_gdf.Printf("\u0020\u0020\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _ccca, _gggbg)
			}
			if !_gggbg._bceg {
				continue
			}
			_fefe := _ccca
			for _dbfa := _ccca + 1; !_gggbg._abdfe && _dbfa < _edaaf; _dbfa++ {
				_cffgc := _fcaf[_dbbc*_edaaf+_dbfa]
				_gggbg.Urx = _cffgc.Urx
				_gggbg._gfga = _gggbg._gfga || _cffgc._gfga
				_gggbg._abag = _gggbg._abag || _cffgc._abag
				_gggbg._abdfe = _cffgc._abdfe
				if _efbc {
					_gdf.Printf("\u0020 \u0020%\u0034\u0064\u003a\u0020\u0025s\u0020\u2192 \u0025\u0073\u000a", _dbfa, _cffgc, _gggbg)
				}
				_fefe = _dbfa
			}
			if _efbc {
				_gdf.Printf(" \u0020 \u0025\u0032\u0064\u0020\u002d\u0020\u0025\u0032d\u0020\u2192\u0020\u0025s\n", _ccca, _fefe, _gggbg)
			}
			_ccca = _fefe
			_feefe[_dbbc][_gggbg.Llx] = _gggbg
		}
	}
	_cebbd := make(map[float64]map[float64]gridTile, _fefg)
	_bbag := make(map[float64]map[float64]struct{}, _fefg)
	for _cgec := _fefg - 1; _cgec >= 0; _cgec-- {
		_cface := _fcaf[_cgec*_edaaf].Lly
		_cebbd[_cface] = make(map[float64]gridTile, _edaaf)
		_bbag[_cface] = make(map[float64]struct{}, _edaaf)
	}
	if _efbc {
		_ag.Log.Info("\u0072u\u006c\u0069n\u0067\u004c\u0069s\u0074\u002e\u0061\u0073\u0054\u0069\u006ci\u006e\u0067\u003a\u0020\u0063\u006fa\u006c\u0065\u0073\u0063\u0065\u0020\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _deec)
	}
	for _fbfc := _fefg - 1; _fbfc >= 0; _fbfc-- {
		_begc := _fcaf[_fbfc*_edaaf].Lly
		_ebce := _feefe[_fbfc]
		if _efbc {
			_gdf.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _fbfc)
		}
		for _, _ggfa := range _gfbg(_ebce) {
			if _, _egbge := _bbag[_begc][_ggfa]; _egbge {
				continue
			}
			_eadfd := _ebce[_ggfa]
			if _efbc {
				_gdf.Printf(" \u0020\u0020\u0020\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _eadfd.String())
			}
			for _abdd := _fbfc - 1; _abdd >= 0; _abdd-- {
				if _eadfd._abag {
					break
				}
				_cgfga := _feefe[_abdd]
				_gbfg, _edee := _cgfga[_ggfa]
				if !_edee {
					break
				}
				if _gbfg.Urx != _eadfd.Urx {
					break
				}
				_eadfd._abag = _gbfg._abag
				_eadfd.Lly = _gbfg.Lly
				if _efbc {
					_gdf.Printf("\u0020\u0020\u0020\u0020  \u0020\u0020\u0076\u003d\u0025\u0073\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _gbfg.String(), _eadfd.String())
				}
				_bbag[_gbfg.Lly][_gbfg.Llx] = struct{}{}
			}
			if _fbfc == 0 {
				_eadfd._abag = true
			}
			if _eadfd.complete() {
				_cebbd[_begc][_ggfa] = _eadfd
			}
		}
	}
	_bddb := gridTiling{PdfRectangle: _deec, _fdeac: _cffd(_cebbd), _gccf: _badbf(_cebbd), _ffdd: _cebbd}
	_bddb.log("\u0043r\u0065\u0061\u0074\u0065\u0064")
	return _bddb
}

// Extractor stores and offers functionality for extracting content from PDF pages.
type Extractor struct {
	_fc  string
	_ad  *_bf.PdfPageResources
	_bb  _bf.PdfRectangle
	_ga  *_bf.PdfRectangle
	_ac  map[string]fontEntry
	_fd  map[string]textResult
	_eb  int64
	_cgf int
	_bc  *Options
}

// ExtractFonts returns all font information from the page extractor, including
// font name, font type, the raw data of the embedded font file (if embedded), font descriptor and more.
//
// The argument `previousPageFonts` is used when trying to build a complete font catalog for multiple pages or the entire document.
// The entries from `previousPageFonts` are added to the returned result unless already included in the page, i.e. no duplicate entries.
//
// NOTE: If previousPageFonts is nil, all fonts from the page will be returned. Use it when building up a full list of fonts for a document or page range.
func (_ca *Extractor) ExtractFonts(previousPageFonts *PageFonts) (*PageFonts, error) {
	_caf := PageFonts{}
	_daf := _caf.extractPageResourcesToFont(_ca._ad)
	if _daf != nil {
		return nil, _daf
	}
	if previousPageFonts != nil {
		for _, _aff := range previousPageFonts.Fonts {
			if !_ea(_caf.Fonts, _aff.FontName) {
				_caf.Fonts = append(_caf.Fonts, _aff)
			}
		}
	}
	return &PageFonts{Fonts: _caf.Fonts}, nil
}

var (
	_cg = _g.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_fa = _g.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func _fbga(_acgfg float64) bool { return _b.Abs(_acgfg) < _ffba }
func (_cdede paraList) findTableGrid(_bbedg gridTiling) (*textTable, map[*textPara]struct{}) {
	_fgafd := len(_bbedg._fdeac)
	_eafd := len(_bbedg._gccf)
	_gffb := textTable{_acgbd: true, _baee: _fgafd, _cabfg: _eafd, _aagc: make(map[uint64]*textPara, _fgafd*_eafd), _bbfb: make(map[uint64]compositeCell, _fgafd*_eafd)}
	_ecdcd := make(map[*textPara]struct{})
	_eccb := int((1.0 - _bbgg) * float64(_fgafd*_eafd))
	_ddcb := 0
	if _efbc {
		_ag.Log.Info("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0064\u0020\u0078\u0020%\u0064", _fgafd, _eafd)
	}
	for _fagg, _eacfe := range _bbedg._gccf {
		_gdbd, _cbaa := _bbedg._ffdd[_eacfe]
		if !_cbaa {
			continue
		}
		for _dccd, _acbgc := range _bbedg._fdeac {
			_ggdca, _eeed := _gdbd[_acbgc]
			if !_eeed {
				continue
			}
			_febbc := _cdede.inTile(_ggdca)
			if len(_febbc) == 0 {
				_ddcb++
				if _ddcb > _eccb {
					if _efbc {
						_ag.Log.Info("\u0021\u006e\u0075m\u0045\u006d\u0070\u0074\u0079\u003d\u0025\u0064", _ddcb)
					}
					return nil, nil
				}
			} else {
				_gffb.putComposite(_dccd, _fagg, _febbc, _ggdca.PdfRectangle)
				for _, _addgg := range _febbc {
					_ecdcd[_addgg] = struct{}{}
				}
			}
		}
	}
	_cggb := 0
	for _cgcbc := 0; _cgcbc < _fgafd; _cgcbc++ {
		_adec := _gffb.get(_cgcbc, 0)
		if _adec == nil || !_adec._dgggff {
			_cggb++
		}
	}
	if _cggb == 0 {
		if _efbc {
			_ag.Log.Info("\u0021\u006e\u0075m\u0048\u0065\u0061\u0064\u0065\u0072\u003d\u0030")
		}
		return nil, nil
	}
	_eecgc := _gffb.reduceTiling(_bbedg, _dded)
	_eecgc = _eecgc.subdivide()
	return _eecgc, _ecdcd
}
func (_aceb *textObject) setTextMatrix(_bdeea []float64) {
	if len(_bdeea) != 6 {
		_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006c\u0065\u006e\u0028\u0066\u0029\u0020\u0021\u003d\u0020\u0036\u0020\u0028\u0025\u0064\u0029", len(_bdeea))
		return
	}
	_dec, _abbg, _agb, _aacg, _dbbg, _aec := _bdeea[0], _bdeea[1], _bdeea[2], _bdeea[3], _bdeea[4], _bdeea[5]
	_aceb._gcb = _ee.NewMatrix(_dec, _abbg, _agb, _aacg, _dbbg, _aec)
	_aceb._baff = _aceb._gcb
}
func (_cgfa *PageText) computeViews() {
	var _ddff rulingList
	if _ecec {
		_caab := _gabdb(_cgfa._fcdb)
		_ddff = append(_ddff, _caab...)
	}
	if _dcaed {
		_eggb := _ddce(_cgfa._cddd)
		_ddff = append(_ddff, _eggb...)
	}
	_ddff, _fcce := _ddff.toTilings()
	var _acef paraList
	_dbdf := len(_cgfa._cac)
	for _cfbb := 0; _cfbb < 360 && _dbdf > 0; _cfbb += 90 {
		_cbg := make([]*textMark, 0, len(_cgfa._cac)-_dbdf)
		for _, _fec := range _cgfa._cac {
			if _fec._bfbc == _cfbb {
				_cbg = append(_cbg, _fec)
			}
		}
		if len(_cbg) > 0 {
			_gfef := _aadf(_cbg, _cgfa._fabf, _ddff, _fcce)
			_acef = append(_acef, _gfef...)
			_dbdf -= len(_cbg)
		}
	}
	_dggc := new(_gd.Buffer)
	_acef.writeText(_dggc)
	_cgfa._gcea = _dggc.String()
	_cgfa._cagb = _acef.toTextMarks()
	_cgfa._ece = _acef.tables()
	if _beae {
		_ag.Log.Info("\u0063\u006f\u006dpu\u0074\u0065\u0056\u0069\u0065\u0077\u0073\u003a\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064", len(_cgfa._ece))
	}
}
func (_feedf paraList) computeEBBoxes() {
	if _ceabd {
		_ag.Log.Info("\u0063o\u006dp\u0075\u0074\u0065\u0045\u0042\u0042\u006f\u0078\u0065\u0073\u003a")
	}
	for _, _edgd := range _feedf {
		_edgd._dcfc = _edgd.PdfRectangle
	}
	_aafc := _feedf.yNeighbours(0)
	for _aacc, _dbbd := range _feedf {
		_aebe := _dbbd._dcfc
		_effc, _ceef := -1.0e9, +1.0e9
		for _, _adgc := range _aafc[_dbbd] {
			_babf := _feedf[_adgc]._dcfc
			if _babf.Urx < _aebe.Llx {
				_effc = _b.Max(_effc, _babf.Urx)
			} else if _aebe.Urx < _babf.Llx {
				_ceef = _b.Min(_ceef, _babf.Llx)
			}
		}
		for _aaaef, _egaf := range _feedf {
			_gabd := _egaf._dcfc
			if _aacc == _aaaef || _gabd.Ury > _aebe.Lly {
				continue
			}
			if _effc <= _gabd.Llx && _gabd.Llx < _aebe.Llx {
				_aebe.Llx = _gabd.Llx
			} else if _gabd.Urx <= _ceef && _aebe.Urx < _gabd.Urx {
				_aebe.Urx = _gabd.Urx
			}
		}
		if _ceabd {
			_gdf.Printf("\u0025\u0034\u0064\u003a %\u0036\u002e\u0032\u0066\u2192\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0071\u000a", _aacc, _dbbd._dcfc, _aebe, _babd(_dbbd.text(), 50))
		}
		_dbbd._dcfc = _aebe
	}
	if _gfea {
		for _, _bdffg := range _feedf {
			_bdffg.PdfRectangle = _bdffg._dcfc
		}
	}
}

// String returns a human readable description of `s`.
func (_gbbbb intSet) String() string {
	var _bbaa []int
	for _cfgda := range _gbbbb {
		if _gbbbb.has(_cfgda) {
			_bbaa = append(_bbaa, _cfgda)
		}
	}
	_f.Ints(_bbaa)
	return _gdf.Sprintf("\u0025\u002b\u0076", _bbaa)
}
func _cada(_gcbed, _abac _ee.Point) bool {
	_bdfaa := _b.Abs(_gcbed.X - _abac.X)
	_gfcbb := _b.Abs(_gcbed.Y - _abac.Y)
	return _dgcd(_bdfaa, _gfcbb)
}
func (_acfbf paraList) merge() *textPara {
	_ag.Log.Trace("\u006d\u0065\u0072\u0067\u0065:\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_acfbf))
	if len(_acfbf) == 0 {
		return nil
	}
	_acfbf.sortReadingOrder()
	_bdfe := _acfbf[0].PdfRectangle
	_ccdf := _acfbf[0]._ecbdg
	for _, _acea := range _acfbf[1:] {
		_bdfe = _effa(_bdfe, _acea.PdfRectangle)
		_ccdf = append(_ccdf, _acea._ecbdg...)
	}
	return _egecc(_bdfe, _ccdf)
}

// String returns a string describing `ma`.
func (_fed TextMarkArray) String() string {
	_feed := len(_fed._fceg)
	if _feed == 0 {
		return "\u0045\u004d\u0050T\u0059"
	}
	_aag := _fed._fceg[0]
	_gfgf := _fed._fceg[_feed-1]
	return _gdf.Sprintf("\u007b\u0054\u0045\u0058\u0054\u004d\u0041\u0052K\u0041\u0052\u0052AY\u003a\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025s\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d%\u0073\u007d", _feed, _aag, _gfgf)
}
func (_cfgf rulingList) splitSec() []rulingList {
	_f.Slice(_cfgf, func(_aecac, _gfbge int) bool {
		_bbef, _ccebdb := _cfgf[_aecac], _cfgf[_gfbge]
		if _bbef._bgeb != _ccebdb._bgeb {
			return _bbef._bgeb < _ccebdb._bgeb
		}
		return _bbef._eecc < _ccebdb._eecc
	})
	_dgecf := make(map[*ruling]struct{}, len(_cfgf))
	_aecab := func(_ebfe *ruling) rulingList {
		_ffbab := rulingList{_ebfe}
		_dgecf[_ebfe] = struct{}{}
		for _, _bcagc := range _cfgf {
			if _, _gdcd := _dgecf[_bcagc]; _gdcd {
				continue
			}
			for _, _cbcg := range _ffbab {
				if _bcagc.alignsSec(_cbcg) {
					_ffbab = append(_ffbab, _bcagc)
					_dgecf[_bcagc] = struct{}{}
					break
				}
			}
		}
		return _ffbab
	}
	_daccf := []rulingList{_aecab(_cfgf[0])}
	for _, _eaac := range _cfgf[1:] {
		if _, _ggdg := _dgecf[_eaac]; _ggdg {
			continue
		}
		_daccf = append(_daccf, _aecab(_eaac))
	}
	return _daccf
}
func (_ceb *imageExtractContext) processOperand(_dcb *_da.ContentStreamOperation, _cce _da.GraphicsState, _ccb *_bf.PdfPageResources) error {
	if _dcb.Operand == "\u0042\u0049" && len(_dcb.Params) == 1 {
		_fcc, _gcf := _dcb.Params[0].(*_da.ContentStreamInlineImage)
		if !_gcf {
			return nil
		}
		if _dbf, _ba := _bg.GetBoolVal(_fcc.ImageMask); _ba {
			if _dbf && !_ceb._gfa.IncludeInlineStencilMasks {
				return nil
			}
		}
		return _ceb.extractInlineImage(_fcc, _cce, _ccb)
	} else if _dcb.Operand == "\u0044\u006f" && len(_dcb.Params) == 1 {
		_dca, _cfg := _bg.GetName(_dcb.Params[0])
		if !_cfg {
			_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _cg
		}
		_, _ccg := _ccb.GetXObjectByName(*_dca)
		switch _ccg {
		case _bf.XObjectTypeImage:
			return _ceb.extractXObjectImage(_dca, _cce, _ccb)
		case _bf.XObjectTypeForm:
			return _ceb.extractFormImages(_dca, _cce, _ccb)
		}
	}
	return nil
}
func (_afedg paraList) log(_bebe string) {
	if !_dada {
		return
	}
	_ag.Log.Info("%\u0038\u0073\u003a\u0020\u0025\u0064 \u0070\u0061\u0072\u0061\u0073\u0020=\u003d\u003d\u003d\u003d\u003d\u003d\u002d-\u002d\u002d\u002d\u002d\u002d\u003d\u003d\u003d\u003d\u003d=\u003d", _bebe, len(_afedg))
	for _egdg, _ecfca := range _afedg {
		if _ecfca == nil {
			continue
		}
		_bcbcd := _ecfca.text()
		_ggaf := "\u0020\u0020"
		if _ecfca._bccg != nil {
			_ggaf = _gdf.Sprintf("\u005b%\u0064\u0078\u0025\u0064\u005d", _ecfca._bccg._baee, _ecfca._bccg._cabfg)
		}
		_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0025s\u0020\u0025\u0071\u000a", _egdg, _ecfca.PdfRectangle, _ggaf, _babd(_bcbcd, 50))
	}
}
func (_fdfg *ruling) equals(_fdec *ruling) bool {
	return _fdfg._cgef == _fdec._cgef && _cgcd(_fdfg._eead, _fdec._eead) && _cgcd(_fdfg._bgeb, _fdec._bgeb) && _cgcd(_fdfg._eecc, _fdec._eecc)
}
func (_ccaa *textObject) reset() {
	_ccaa._gcb = _ee.IdentityMatrix()
	_ccaa._baff = _ee.IdentityMatrix()
	_ccaa._dbdg = nil
}

// ToTextMark returns the public view of `tm`.
func (_gdgf *textMark) ToTextMark() TextMark {
	return TextMark{Text: _gdgf._eeaf, Original: _gdgf._abcec, BBox: _gdgf._bbgf, Font: _gdgf._becff, FontSize: _gdgf._abba, FillColor: _gdgf._agbe, StrokeColor: _gdgf._dggfa, Orientation: _gdgf._bfbc, DirectObject: _gdgf._gded, ObjString: _gdgf._cedf, Tw: _gdgf.Tw, Th: _gdgf.Th, Tc: _gdgf._dbe, Index: _gdgf._eced}
}

// String returns a string descibing `i`.
func (_aaaca gridTile) String() string {
	_geacf := func(_abfdf bool, _bdbfg string) string {
		if _abfdf {
			return _bdbfg
		}
		return "\u005f"
	}
	return _gdf.Sprintf("\u00256\u002e2\u0066\u0020\u0025\u0031\u0073%\u0031\u0073%\u0031\u0073\u0025\u0031\u0073", _aaaca.PdfRectangle, _geacf(_aaaca._bceg, "\u004c"), _geacf(_aaaca._abdfe, "\u0052"), _geacf(_aaaca._abag, "\u0042"), _geacf(_aaaca._gfga, "\u0054"))
}
func _aecc(_ffad func(*wordBag, *textWord, float64) bool, _bccd float64) func(*wordBag, *textWord) bool {
	return func(_bae *wordBag, _egcd *textWord) bool { return _ffad(_bae, _egcd, _bccd) }
}
func (_cabg *textPara) taken() bool { return _cabg == nil || _cabg._bggb }

type pathSection struct {
	_gdbf []*subpath
	_eda.Color
}

func _afa(_bbfg, _dfcca bounded) float64 {
	_fbcfdb := _addb(_bbfg, _dfcca)
	if !_fbga(_fbcfdb) {
		return _fbcfdb
	}
	return _dagb(_bbfg, _dfcca)
}
func (_agde *compositeCell) updateBBox() {
	for _, _bdeb := range _agde.paraList {
		_agde.PdfRectangle = _effa(_agde.PdfRectangle, _bdeb.PdfRectangle)
	}
}
func (_cffb *textPara) bbox() _bf.PdfRectangle { return _cffb.PdfRectangle }
func _ceabf(_efff, _gfaf *textPara) bool {
	if _efff._dgggff || _gfaf._dgggff {
		return true
	}
	return _fbga(_efff.depth() - _gfaf.depth())
}

type rulingKind int

func _ceba(_fggb, _fede _bf.PdfRectangle) bool {
	return _fggb.Llx <= _fede.Llx && _fede.Urx <= _fggb.Urx && _fggb.Lly <= _fede.Lly && _fede.Ury <= _fggb.Ury
}

// Options extractor options.
type Options struct {

	// ApplyCropBox will extract page text based on page cropbox if set to `true`.
	ApplyCropBox bool
}

func _ccagd(_dfdd map[int]intSet) []int {
	_ecbaf := make([]int, 0, len(_dfdd))
	for _dccdf := range _dfdd {
		_ecbaf = append(_ecbaf, _dccdf)
	}
	_f.Ints(_ecbaf)
	return _ecbaf
}

type paraList []*textPara

func (_ffeg *textTable) compositeColCorridors() map[int][]float64 {
	_ecbdd := make(map[int][]float64, _ffeg._baee)
	if _beae {
		_ag.Log.Info("\u0063\u006f\u006d\u0070o\u0073\u0069\u0074\u0065\u0043\u006f\u006c\u0043\u006f\u0072r\u0069d\u006f\u0072\u0073\u003a\u0020\u0077\u003d%\u0064\u0020", _ffeg._baee)
	}
	for _fggc := 0; _fggc < _ffeg._baee; _fggc++ {
		_ecbdd[_fggc] = nil
	}
	return _ecbdd
}
func (_ecb *shapesState) drawRectangle(_gfag, _bdaa, _afca, _cceb float64) {
	if _bccf {
		_gacg := _ecb.devicePoint(_gfag, _bdaa)
		_cdeg := _ecb.devicePoint(_gfag+_afca, _bdaa+_cceb)
		_gdba := _bf.PdfRectangle{Llx: _gacg.X, Lly: _gacg.Y, Urx: _cdeg.X, Ury: _cdeg.Y}
		_ag.Log.Info("d\u0072a\u0077\u0052\u0065\u0063\u0074\u0061\u006e\u0067l\u0065\u003a\u0020\u00256.\u0032\u0066", _gdba)
	}
	_ecb.newSubPath()
	_ecb.moveTo(_gfag, _bdaa)
	_ecb.lineTo(_gfag+_afca, _bdaa)
	_ecb.lineTo(_gfag+_afca, _bdaa+_cceb)
	_ecb.lineTo(_gfag, _bdaa+_cceb)
	_ecb.closePath()
}
func (_gbee *wordBag) maxDepth() float64 { return _gbee._fcgc - _gbee.Lly }
func (_fcdgd *wordBag) removeWord(_ffaf *textWord, _fag int) {
	_abbd := _fcdgd._bbf[_fag]
	_abbd = _fbcac(_abbd, _ffaf)
	if len(_abbd) == 0 {
		delete(_fcdgd._bbf, _fag)
	} else {
		_fcdgd._bbf[_fag] = _abbd
	}
}
func (_cbgfa *textTable) reduce() *textTable {
	_eacce := make([]int, 0, _cbgfa._cabfg)
	_debfc := make([]int, 0, _cbgfa._baee)
	for _dddef := 0; _dddef < _cbgfa._cabfg; _dddef++ {
		if !_cbgfa.emptyCompositeRow(_dddef) {
			_eacce = append(_eacce, _dddef)
		}
	}
	for _efbf := 0; _efbf < _cbgfa._baee; _efbf++ {
		if !_cbgfa.emptyCompositeColumn(_efbf) {
			_debfc = append(_debfc, _efbf)
		}
	}
	if len(_eacce) == _cbgfa._cabfg && len(_debfc) == _cbgfa._baee {
		return _cbgfa
	}
	_ggbfc := textTable{_acgbd: _cbgfa._acgbd, _baee: len(_debfc), _cabfg: len(_eacce), _aagc: make(map[uint64]*textPara, len(_debfc)*len(_eacce))}
	if _beae {
		_ag.Log.Info("\u0072\u0065\u0064\u0075ce\u003a\u0020\u0025\u0064\u0078\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064\u0078%\u0064", _cbgfa._baee, _cbgfa._cabfg, len(_debfc), len(_eacce))
		_ag.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _debfc)
		_ag.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _eacce)
	}
	for _dfae, _fedd := range _eacce {
		for _acgg, _effef := range _debfc {
			_gged, _accbc := _cbgfa.getComposite(_effef, _fedd)
			if _gged == nil {
				continue
			}
			if _beae {
				_gdf.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _acgg, _dfae, _effef, _fedd, _babd(_gged.merge().text(), 50))
			}
			_ggbfc.putComposite(_acgg, _dfae, _gged, _accbc)
		}
	}
	return &_ggbfc
}
func _ddedb(_aecf, _ddfeg _ee.Point) rulingKind {
	_cgcgb := _b.Abs(_aecf.X - _ddfeg.X)
	_gfcbg := _b.Abs(_aecf.Y - _ddfeg.Y)
	return _bfcbe(_cgcgb, _gfcbg, _cdfc)
}
func (_ddc *Extractor) extractPageText(_age string, _ded *_bf.PdfPageResources, _ccge _ee.Matrix, _fgg int) (*PageText, int, int, error) {
	_ag.Log.Trace("\u0065x\u0074\u0072\u0061\u0063t\u0050\u0061\u0067\u0065\u0054e\u0078t\u003a \u006c\u0065\u0076\u0065\u006c\u003d\u0025d", _fgg)
	_cdb := &PageText{_fabf: _ddc._bb}
	_ffb := _cde(_ddc._bb)
	var _gca stateStack
	_dbb := _eag(_ddc, _ded, _da.GraphicsState{}, &_ffb, &_gca)
	_ge := shapesState{_efgg: _ccge, _egb: _ee.IdentityMatrix(), _gdcg: _dbb}
	var _gge bool
	if _fgg > _egc {
		_gga := _g.New("\u0066\u006f\u0072\u006d s\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077")
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0065\u0078\u0074\u0072\u0061\u0063\u0074\u0050\u0061\u0067\u0065\u0054\u0065\u0078\u0074\u002e\u0020\u0072\u0065\u0063u\u0072\u0073\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u003d\u0025\u0064 \u0065r\u0072\u003d\u0025\u0076", _fgg, _gga)
		return _cdb, _ffb._fce, _ffb._ddga, _gga
	}
	_feb := _da.NewContentStreamParser(_age)
	_facf, _gad := _feb.Parse()
	if _gad != nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020e\u0078\u0074\u0072a\u0063\u0074\u0050\u0061g\u0065\u0054\u0065\u0078\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gad)
		return _cdb, _ffb._fce, _ffb._ddga, _gad
	}
	_cdb._bfg = _facf
	_gde := _da.NewContentStreamProcessor(*_facf)
	_gde.AddHandler(_da.HandlerConditionEnumAllOperands, "", func(_bac *_da.ContentStreamOperation, _aea _da.GraphicsState, _aae *_bf.PdfPageResources) error {
		_ffce := _bac.Operand
		if _ddb {
			_ag.Log.Info("\u0026&\u0026\u0020\u006f\u0070\u003d\u0025s", _bac)
		}
		switch _ffce {
		case "\u0071":
			if _bccf {
				_ag.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ge._egb)
			}
			_gca.push(&_ffb)
		case "\u0051":
			if !_gca.empty() {
				_ffb = *_gca.pop()
			}
			_ge._egb = _aea.CTM
			if _bccf {
				_ag.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ge._egb)
			}
		case "\u0042\u0054":
			if _gge {
				_ag.Log.Debug("\u0042\u0054\u0020\u0063\u0061\u006c\u006c\u0065\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0069n\u0020\u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
				_cdb._cac = append(_cdb._cac, _dbb._dbdg...)
			}
			_gge = true
			_abcc := _aea
			_abcc.CTM = _ccge.Mult(_abcc.CTM)
			_dbb = _eag(_ddc, _aae, _abcc, &_ffb, &_gca)
			_ge._gdcg = _dbb
		case "\u0045\u0054":
			if !_gge {
				_ag.Log.Debug("\u0045\u0054\u0020ca\u006c\u006c\u0065\u0064\u0020\u006f\u0075\u0074\u0073i\u0064e\u0020o\u0066 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_gge = false
			_cdb._cac = append(_cdb._cac, _dbb._dbdg...)
			_dbb.reset()
		case "\u0054\u002a":
			_dbb.nextLine()
		case "\u0054\u0064":
			if _gdg, _adb := _dbb.checkOp(_bac, 2, true); !_gdg {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _adb)
				return _adb
			}
			_egae, _egga, _ade := _ggfdc(_bac.Params)
			if _ade != nil {
				return _ade
			}
			_dbb.moveText(_egae, _egga)
		case "\u0054\u0044":
			if _bdcb, _gfb := _dbb.checkOp(_bac, 2, true); !_bdcb {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gfb)
				return _gfb
			}
			_cee, _cdd, _gff := _ggfdc(_bac.Params)
			if _gff != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gff)
				return _gff
			}
			_dbb.moveTextSetLeading(_cee, _cdd)
		case "\u0054\u006a":
			if _aad, _dbfb := _dbb.checkOp(_bac, 1, true); !_aad {
				_ag.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0054\u006a\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d%\u0076", _bac, _dbfb)
				return _dbfb
			}
			_bcb := _bg.TraceToDirectObject(_bac.Params[0])
			_bab, _dgea := _bg.GetStringBytes(_bcb)
			if !_dgea {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020T\u006a\u0020o\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074S\u0074\u0072\u0069\u006e\u0067\u0042\u0079\u0074\u0065\u0073\u0020\u0066a\u0069\u006c\u0065\u0064", _bac)
				return _bg.ErrTypeError
			}
			return _dbb.showText(_bcb, _bab)
		case "\u0054\u004a":
			if _fba, _ace := _dbb.checkOp(_bac, 1, true); !_fba {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ace)
				return _ace
			}
			_ead, _acc := _bg.GetArray(_bac.Params[0])
			if !_acc {
				_ag.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _bac)
				return _gad
			}
			return _dbb.showTextAdjusted(_ead)
		case "\u0027":
			if _fab, _aade := _dbb.checkOp(_bac, 1, true); !_fab {
				_ag.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0027\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _aade)
				return _aade
			}
			_fca := _bg.TraceToDirectObject(_bac.Params[0])
			_gcg, _bcd := _bg.GetStringBytes(_fca)
			if !_bcd {
				_ag.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020'\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _bac)
				return _bg.ErrTypeError
			}
			_dbb.nextLine()
			return _dbb.showText(_fca, _gcg)
		case "\u0022":
			if _ccbb, _aab := _dbb.checkOp(_bac, 3, true); !_ccbb {
				_ag.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0022\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _aab)
				return _aab
			}
			_ggbd, _eaa, _bdee := _ggfdc(_bac.Params[:2])
			if _bdee != nil {
				return _bdee
			}
			_efc := _bg.TraceToDirectObject(_bac.Params[2])
			_bef, _afe := _bg.GetStringBytes(_efc)
			if !_afe {
				_ag.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020\"\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _bac)
				return _bg.ErrTypeError
			}
			_dbb.setCharSpacing(_ggbd)
			_dbb.setWordSpacing(_eaa)
			_dbb.nextLine()
			return _dbb.showText(_efc, _bef)
		case "\u0054\u004c":
			_bdg, _adf := _efd(_bac)
			if _adf != nil {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004c\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _adf)
				return _adf
			}
			_dbb.setTextLeading(_bdg)
		case "\u0054\u0063":
			_fgc, _cbc := _efd(_bac)
			if _cbc != nil {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0063\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cbc)
				return _cbc
			}
			_dbb.setCharSpacing(_fgc)
		case "\u0054\u0066":
			if _ffg, _bdcd := _dbb.checkOp(_bac, 2, true); !_ffg {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0066\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bdcd)
				return _bdcd
			}
			_gddg, _dbc := _bg.GetNameVal(_bac.Params[0])
			if !_dbc {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _bac)
				return _bg.ErrTypeError
			}
			_cgd, _acbe := _bg.GetNumberAsFloat(_bac.Params[1])
			if !_dbc {
				_ag.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u0046\u006c\u006f\u0061\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065d\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bac, _acbe)
				return _acbe
			}
			_acbe = _dbb.setFont(_gddg, _cgd)
			_dbb._cdbc = _fb.Is(_acbe, _bg.ErrNotSupported)
			if _acbe != nil && !_dbb._cdbc {
				return _acbe
			}
		case "\u0054\u006d":
			if _ecg, _gadb := _dbb.checkOp(_bac, 6, true); !_ecg {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u006d\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gadb)
				return _gadb
			}
			_dgd, _bdb := _bg.GetNumbersAsFloat(_bac.Params)
			if _bdb != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bdb)
				return _bdb
			}
			_dbb.setTextMatrix(_dgd)
		case "\u0054\u0072":
			if _abe, _acec := _dbb.checkOp(_bac, 1, true); !_abe {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0072\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _acec)
				return _acec
			}
			_eea, _gced := _bg.GetIntVal(_bac.Params[0])
			if !_gced {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0020\u006f\u0070\u003d\u0025\u0073 \u0047e\u0074\u0049\u006e\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _bac)
				return _bg.ErrTypeError
			}
			_dbb.setTextRenderMode(_eea)
		case "\u0054\u0073":
			if _aeb, _dag := _dbb.checkOp(_bac, 1, true); !_aeb {
				_ag.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dag)
				return _dag
			}
			_efg, _ccc := _bg.GetNumberAsFloat(_bac.Params[0])
			if _ccc != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ccc)
				return _ccc
			}
			_dbb.setTextRise(_efg)
		case "\u0054\u0077":
			if _gag, _fdgf := _dbb.checkOp(_bac, 1, true); !_gag {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fdgf)
				return _fdgf
			}
			_ebaa, _bda := _bg.GetNumberAsFloat(_bac.Params[0])
			if _bda != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bda)
				return _bda
			}
			_dbb.setWordSpacing(_ebaa)
		case "\u0054\u007a":
			if _dafg, _bcbc := _dbb.checkOp(_bac, 1, true); !_dafg {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bcbc)
				return _bcbc
			}
			_cfc, _eae := _bg.GetNumberAsFloat(_bac.Params[0])
			if _eae != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eae)
				return _eae
			}
			_dbb.setHorizScaling(_cfc)
		case "\u0063\u006d":
			_ge._egb = _aea.CTM
			if _ge._egb.Singular() {
				_dgee := _ee.IdentityMatrix().Translate(_ge._egb.Translation())
				_ag.Log.Debug("S\u0069n\u0067\u0075\u006c\u0061\u0072\u0020\u0063\u0074m\u003d\u0025\u0073\u2192%s", _ge._egb, _dgee)
				_ge._egb = _dgee
			}
			if _bccf {
				_ag.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ge._egb)
			}
		case "\u006d":
			if len(_bac.Params) != 2 {
				_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _fa)
				return nil
			}
			_dbaa, _accg := _bg.GetNumbersAsFloat(_bac.Params)
			if _accg != nil {
				return _accg
			}
			_ge.moveTo(_dbaa[0], _dbaa[1])
		case "\u006c":
			if len(_bac.Params) != 2 {
				_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _fa)
				return nil
			}
			_cfb, _bdgg := _bg.GetNumbersAsFloat(_bac.Params)
			if _bdgg != nil {
				return _bdgg
			}
			_ge.lineTo(_cfb[0], _cfb[1])
		case "\u0063":
			if len(_bac.Params) != 6 {
				return _fa
			}
			_accgg, _acga := _bg.GetNumbersAsFloat(_bac.Params)
			if _acga != nil {
				return _acga
			}
			_ag.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _accgg)
			_ge.cubicTo(_accgg[0], _accgg[1], _accgg[2], _accgg[3], _accgg[4], _accgg[5])
		case "\u0076", "\u0079":
			if len(_bac.Params) != 4 {
				return _fa
			}
			_cge, _dfc := _bg.GetNumbersAsFloat(_bac.Params)
			if _dfc != nil {
				return _dfc
			}
			_ag.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _cge)
			_ge.quadraticTo(_cge[0], _cge[1], _cge[2], _cge[3])
		case "\u0068":
			_ge.closePath()
		case "\u0072\u0065":
			if len(_bac.Params) != 4 {
				return _fa
			}
			_adgf, _aadee := _bg.GetNumbersAsFloat(_bac.Params)
			if _aadee != nil {
				return _aadee
			}
			_ge.drawRectangle(_adgf[0], _adgf[1], _adgf[2], _adgf[3])
			_ge.closePath()
		case "\u0053":
			_ge.stroke(&_cdb._fcdb)
			_ge.clearPath()
		case "\u0073":
			_ge.closePath()
			_ge.stroke(&_cdb._fcdb)
			_ge.clearPath()
		case "\u0046":
			_ge.fill(&_cdb._cddd)
			_ge.clearPath()
		case "\u0066", "\u0066\u002a":
			_ge.closePath()
			_ge.fill(&_cdb._cddd)
			_ge.clearPath()
		case "\u0042", "\u0042\u002a":
			_ge.fill(&_cdb._cddd)
			_ge.stroke(&_cdb._fcdb)
			_ge.clearPath()
		case "\u0062", "\u0062\u002a":
			_ge.closePath()
			_ge.fill(&_cdb._cddd)
			_ge.stroke(&_cdb._fcdb)
			_ge.clearPath()
		case "\u006e":
			_ge.clearPath()
		case "\u0044\u006f":
			if len(_bac.Params) == 0 {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0058\u004fbj\u0065c\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0070\u0065\u0072\u0061n\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u006f\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072.\u0020\u0047\u006f\u0074\u0020\u0025\u002b\u0076\u002e", _bac.Params)
				return _bg.ErrRangeError
			}
			_afc, _gee := _bg.GetName(_bac.Params[0])
			if !_gee {
				_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u006f\u0020\u006f\u0070e\u0072a\u0074\u006f\u0072\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006fp\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u002b\u0076\u002e", _bac.Params[0])
				return _bg.ErrTypeError
			}
			_, _cgce := _aae.GetXObjectByName(*_afc)
			if _cgce != _bf.XObjectTypeForm {
				break
			}
			_ege, _gee := _ddc._fd[_afc.String()]
			if !_gee {
				_gaf, _bdd := _aae.GetXObjectFormByName(*_afc)
				if _bdd != nil {
					_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bdd)
					return _bdd
				}
				_bdf, _bdd := _gaf.GetContentStream()
				if _bdd != nil {
					_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bdd)
					return _bdd
				}
				_dcc := _gaf.Resources
				if _dcc == nil {
					_dcc = _aae
				}
				_aed := _aea.CTM
				if _gbf, _efcd := _bg.GetArray(_gaf.Matrix); _efcd {
					_bag, _fea := _gbf.GetAsFloat64Slice()
					if _fea != nil {
						return _fea
					}
					if len(_bag) != 6 {
						return _fa
					}
					_cfbe := _ee.NewMatrix(_bag[0], _bag[1], _bag[2], _bag[3], _bag[4], _bag[5])
					_aed = _aea.CTM.Mult(_cfbe)
				}
				_cebf, _ceg, _bagg, _bdd := _ddc.extractPageText(string(_bdf), _dcc, _ccge.Mult(_aed), _fgg+1)
				if _bdd != nil {
					_ag.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bdd)
					return _bdd
				}
				_ege = textResult{*_cebf, _ceg, _bagg}
				_ddc._fd[_afc.String()] = _ege
			}
			_ge._egb = _aea.CTM
			if _bccf {
				_ag.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ge._egb)
			}
			_cdb._cac = append(_cdb._cac, _ege._bafb._cac...)
			_cdb._fcdb = append(_cdb._fcdb, _ege._bafb._fcdb...)
			_cdb._cddd = append(_cdb._cddd, _ege._bafb._cddd...)
			_ffb._fce += _ege._dfcc
			_ffb._ddga += _ege._bcbb
		case "\u0072\u0067", "\u0067", "\u006b", "\u0063\u0073", "\u0073\u0063", "\u0073\u0063\u006e":
			_dbb._fcdg.ColorspaceNonStroking = _aea.ColorspaceNonStroking
			_dbb._fcdg.ColorNonStroking = _aea.ColorNonStroking
		case "\u0052\u0047", "\u0047", "\u004b", "\u0043\u0053", "\u0053\u0043", "\u0053\u0043\u004e":
			_dbb._fcdg.ColorspaceStroking = _aea.ColorspaceStroking
			_dbb._fcdg.ColorStroking = _aea.ColorStroking
		}
		return nil
	})
	_gad = _gde.Process(_ded)
	return _cdb, _ffb._fce, _ffb._ddga, _gad
}
func _agcbf(_cbge, _egcb _bf.PdfRectangle) (_bf.PdfRectangle, bool) {
	if !_gcgg(_cbge, _egcb) {
		return _bf.PdfRectangle{}, false
	}
	return _bf.PdfRectangle{Llx: _b.Max(_cbge.Llx, _egcb.Llx), Urx: _b.Min(_cbge.Urx, _egcb.Urx), Lly: _b.Max(_cbge.Lly, _egcb.Lly), Ury: _b.Min(_cbge.Ury, _egcb.Ury)}, true
}
func (_afeb *textTable) emptyCompositeColumn(_cfbbde int) bool {
	for _eege := 0; _eege < _afeb._cabfg; _eege++ {
		if _caee, _bgfda := _afeb._bbfb[_fcbc(_cfbbde, _eege)]; _bgfda {
			if len(_caee.paraList) > 0 {
				return false
			}
		}
	}
	return true
}
func (_abc *imageExtractContext) extractFormImages(_aaf *_bg.PdfObjectName, _ddg _da.GraphicsState, _dab *_bf.PdfPageResources) error {
	_cbd, _fdg := _dab.GetXObjectFormByName(*_aaf)
	if _fdg != nil {
		return _fdg
	}
	if _cbd == nil {
		return nil
	}
	_bfdb, _fdg := _cbd.GetContentStream()
	if _fdg != nil {
		return _fdg
	}
	_efe := _cbd.Resources
	if _efe == nil {
		_efe = _dab
	}
	_fdg = _abc.extractContentStreamImages(string(_bfdb), _efe)
	if _fdg != nil {
		return _fdg
	}
	_abc._cc++
	return nil
}
func (_cfec rulingList) sort() { _f.Slice(_cfec, _cfec.comp) }
func (_dafgf paraList) topoOrder() []int {
	if _dada {
		_ag.Log.Info("\u0074\u006f\u0070\u006f\u004f\u0072\u0064\u0065\u0072\u003a")
	}
	_dafa := len(_dafgf)
	_dfdb := make([]bool, _dafa)
	_cdcg := make([]int, 0, _dafa)
	_egfe := _dafgf.llyOrdering()
	var _adbe func(_eecgf int)
	_adbe = func(_ddaac int) {
		_dfdb[_ddaac] = true
		for _cfbbd := 0; _cfbbd < _dafa; _cfbbd++ {
			if !_dfdb[_cfbbd] {
				if _dafgf.readBefore(_egfe, _ddaac, _cfbbd) {
					_adbe(_cfbbd)
				}
			}
		}
		_cdcg = append(_cdcg, _ddaac)
	}
	for _bede := 0; _bede < _dafa; _bede++ {
		if !_dfdb[_bede] {
			_adbe(_bede)
		}
	}
	return _bceab(_cdcg)
}
func (_babfg *ruling) alignsSec(_cadea *ruling) bool {
	const _gffa = _dcfe + 1.0
	return _babfg._bgeb-_gffa <= _cadea._eecc && _cadea._bgeb-_gffa <= _babfg._eecc
}

type gridTile struct {
	_bf.PdfRectangle
	_gfga, _bceg, _abag, _abdfe bool
}

func (_ggeec *textMark) bbox() _bf.PdfRectangle { return _ggeec.PdfRectangle }
func (_debe compositeCell) hasLines(_ceabfb []*textLine) bool {
	for _bdgb, _bgfdd := range _ceabfb {
		_becag := _gcgg(_debe.PdfRectangle, _bgfdd.PdfRectangle)
		if _beae {
			_gdf.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u005e\u005e\u005e\u0069\u006e\u0074\u0065\u0072\u0073e\u0063t\u0073\u003d\u0025\u0074\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u000a", _becag, _bdgb, len(_ceabfb))
			_gdf.Printf("\u0020\u0020\u0020\u0020  \u005e\u005e\u005e\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025s\u000a", _debe)
			_gdf.Printf("\u0020 \u0020 \u0020\u0020\u0020\u006c\u0069\u006e\u0065\u003d\u0025\u0073\u000a", _bgfdd)
		}
		if _becag {
			return true
		}
	}
	return false
}

// Tables returns the tables extracted from the page.
func (_dadf PageText) Tables() []TextTable {
	if _beae {
		_ag.Log.Info("\u0054\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_dadf._ece))
	}
	return _dadf._ece
}
func (_bfc *imageExtractContext) extractContentStreamImages(_eec string, _cad *_bf.PdfPageResources) error {
	_cfa := _da.NewContentStreamParser(_eec)
	_fcd, _cade := _cfa.Parse()
	if _cade != nil {
		return _cade
	}
	if _bfc._gda == nil {
		_bfc._gda = map[*_bg.PdfObjectStream]*cachedImage{}
	}
	if _bfc._gfa == nil {
		_bfc._gfa = &ImageExtractOptions{}
	}
	_bdc := _da.NewContentStreamProcessor(*_fcd)
	_bdc.AddHandler(_da.HandlerConditionEnumAllOperands, "", _bfc.processOperand)
	return _bdc.Process(_cad)
}
func (_bgdcd *textTable) isExportable() bool {
	if _bgdcd._acgbd {
		return true
	}
	_fadca := func(_fcdbd int) bool {
		_ffee := _bgdcd.get(0, _fcdbd)
		if _ffee == nil {
			return false
		}
		_aebc := _ffee.text()
		_fffg := _a.RuneCountInString(_aebc)
		_dead := _ebbb.MatchString(_aebc)
		return _fffg <= 1 || _dead
	}
	for _cbfaa := 0; _cbfaa < _bgdcd._cabfg; _cbfaa++ {
		if !_fadca(_cbfaa) {
			return true
		}
	}
	return false
}

type textMark struct {
	_bf.PdfRectangle
	_bfbc  int
	_eeaf  string
	_abcec string
	_becff *_bf.PdfFont
	_abba  float64
	_dbe   float64
	_ecbd  _ee.Matrix
	_fcced _ee.Point
	_bbgf  _bf.PdfRectangle
	_agbe  _eda.Color
	_dggfa _eda.Color
	_gded  _bg.PdfObject
	_cedf  []string
	Tw     float64
	Th     float64
	_eced  int
}

func (_aggf *wordBag) depthBand(_efea, _acfe float64) []int {
	if len(_aggf._bbf) == 0 {
		return nil
	}
	return _aggf.depthRange(_aggf.getDepthIdx(_efea), _aggf.getDepthIdx(_acfe))
}

// ImageMark represents an image drawn on a page and its position in device coordinates.
// All coordinates are in device coordinates.
type ImageMark struct {
	Image *_bf.Image

	// Dimensions of the image as displayed in the PDF.
	Width  float64
	Height float64

	// Position of the image in PDF coordinates (lower left corner).
	X float64
	Y float64

	// Angle in degrees, if rotated.
	Angle float64
}

var _ebf = TextMark{Text: "\u005b\u0058\u005d", Original: "\u0020", Meta: true, FillColor: _eda.White, StrokeColor: _eda.White}

func (_gagb *stateStack) empty() bool { return len(*_gagb) == 0 }

// String returns a description of `w`.
func (_eaee *textWord) String() string {
	return _gdf.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _eaee._ecgcg, _eaee.PdfRectangle, _eaee._fbgge, _eaee._bcaa)
}

const _gaeg = 10

func _egecc(_bccfc _bf.PdfRectangle, _gacd []*textLine) *textPara {
	return &textPara{PdfRectangle: _bccfc, _ecbdg: _gacd}
}

// String returns a description of `p`.
func (_gfdb *textPara) String() string {
	if _gfdb._dgggff {
		return _gdf.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u005b\u0045\u004d\u0050\u0054\u0059\u005d", _gfdb.PdfRectangle)
	}
	_efede := ""
	if _gfdb._bccg != nil {
		_efede = _gdf.Sprintf("\u005b\u0025\u0064\u0078\u0025\u0064\u005d\u0020", _gfdb._bccg._baee, _gfdb._bccg._cabfg)
	}
	return _gdf.Sprintf("\u0025\u0036\u002e\u0032f \u0025\u0073\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u0020\u0025\u0071", _gfdb.PdfRectangle, _efede, len(_gfdb._ecbdg), _babd(_gfdb.text(), 50))
}
func (_bdfdg rulingList) intersections() map[int]intSet {
	var _ebcd, _eeef []int
	for _gfefb, _afge := range _bdfdg {
		switch _afge._cgef {
		case _beec:
			_ebcd = append(_ebcd, _gfefb)
		case _ebabd:
			_eeef = append(_eeef, _gfefb)
		}
	}
	if len(_ebcd) < _agec+1 || len(_eeef) < _bafa+1 {
		return nil
	}
	if len(_ebcd)+len(_eeef) > _aedf {
		_ag.Log.Debug("\u0069\u006e\u0074\u0065\u0072\u0073e\u0063\u0074\u0069\u006f\u006e\u0073\u003a\u0020\u0054\u004f\u004f\u0020\u004d\u0041\u004e\u0059\u0020\u0072\u0075\u006ci\u006e\u0067\u0073\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u003d\u0020%\u0064 \u0078\u0020\u0025\u0064", len(_bdfdg), len(_ebcd), len(_eeef))
		return nil
	}
	_efcb := make(map[int]intSet, len(_ebcd)+len(_eeef))
	for _, _cgda := range _ebcd {
		for _, _cdgc := range _eeef {
			if _bdfdg[_cgda].intersects(_bdfdg[_cdgc]) {
				if _, _acfaa := _efcb[_cgda]; !_acfaa {
					_efcb[_cgda] = make(intSet)
				}
				if _, _gcgc := _efcb[_cdgc]; !_gcgc {
					_efcb[_cdgc] = make(intSet)
				}
				_efcb[_cgda].add(_cdgc)
				_efcb[_cdgc].add(_cgda)
			}
		}
	}
	return _efcb
}
func (_fcgb lineRuling) yMean() float64 { return 0.5 * (_fcgb._abec.Y + _fcgb._eacb.Y) }
func _gecbc(_eecca, _afegcc int) int {
	if _eecca > _afegcc {
		return _eecca
	}
	return _afegcc
}
func (_fadf *textPara) toTextMarks(_fegbf *int) []TextMark {
	if _fadf._bccg == nil {
		return _fadf.toCellTextMarks(_fegbf)
	}
	var _efag []TextMark
	for _bcfe := 0; _bcfe < _fadf._bccg._cabfg; _bcfe++ {
		for _cbee := 0; _cbee < _fadf._bccg._baee; _cbee++ {
			_bebf := _fadf._bccg.get(_cbee, _bcfe)
			if _bebf == nil {
				_efag = _abbf(_efag, _fegbf, "\u0009")
			} else {
				_edga := _bebf.toCellTextMarks(_fegbf)
				_efag = append(_efag, _edga...)
			}
			_efag = _abbf(_efag, _fegbf, "\u0020")
		}
		if _bcfe < _fadf._bccg._cabfg-1 {
			_efag = _abbf(_efag, _fegbf, "\u000a")
		}
	}
	return _efag
}
func (_efgee *textPara) fontsize() float64 { return _efgee._ecbdg[0]._ddef }
func (_bdfa *textObject) getStrokeColor() _eda.Color {
	return _gfdbb(_bdfa._fcdg.ColorspaceStroking, _bdfa._fcdg.ColorStroking)
}
func (_bea *textObject) setCharSpacing(_fbg float64) {
	if _bea == nil {
		return
	}
	_bea._degf._bfcb = _fbg
	if _dbde {
		_ag.Log.Info("\u0073\u0065t\u0043\u0068\u0061\u0072\u0053\u0070\u0061\u0063\u0069\u006e\u0067\u003a\u0020\u0025\u002e\u0032\u0066\u0020\u0073\u0074\u0061\u0074e=\u0025\u0073", _fbg, _bea._degf.String())
	}
}
func (_ddec rulingList) primaries() []float64 {
	_degg := make(map[float64]struct{}, len(_ddec))
	for _, _ccebd := range _ddec {
		_degg[_ccebd._eead] = struct{}{}
	}
	_fagb := make([]float64, len(_degg))
	_badb := 0
	for _adaf := range _degg {
		_fagb[_badb] = _adaf
		_badb++
	}
	_f.Float64s(_fagb)
	return _fagb
}
func (_abfd *textObject) setWordSpacing(_gfd float64) {
	if _abfd == nil {
		return
	}
	_abfd._degf._gab = _gfd
}
func _ccdef(_bdad int, _fgae func(int, int) bool) []int {
	_cbcb := make([]int, _bdad)
	for _fbca := range _cbcb {
		_cbcb[_fbca] = _fbca
	}
	_f.Slice(_cbcb, func(_bffd, _feecc int) bool { return _fgae(_cbcb[_bffd], _cbcb[_feecc]) })
	return _cbcb
}
func _eagg(_becf, _bfa _bf.PdfRectangle) bool { return _bfa.Llx <= _becf.Urx && _becf.Llx <= _bfa.Urx }
func (_ggafe lineRuling) asRuling() (*ruling, bool) {
	_gbad := ruling{_cgef: _ggafe._aedc, Color: _ggafe.Color, _gbgb: _caade}
	switch _ggafe._aedc {
	case _beec:
		_gbad._eead = _ggafe.xMean()
		_gbad._bgeb = _b.Min(_ggafe._abec.Y, _ggafe._eacb.Y)
		_gbad._eecc = _b.Max(_ggafe._abec.Y, _ggafe._eacb.Y)
	case _ebabd:
		_gbad._eead = _ggafe.yMean()
		_gbad._bgeb = _b.Min(_ggafe._abec.X, _ggafe._eacb.X)
		_gbad._eecc = _b.Max(_ggafe._abec.X, _ggafe._eacb.X)
	default:
		_ag.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _ggafe._aedc)
		return nil, false
	}
	return &_gbad, true
}
func _gcgg(_adcd, _bbdb _bf.PdfRectangle) bool { return _eagg(_adcd, _bbdb) && _efeg(_adcd, _bbdb) }
func (_fcea *textLine) markWordBoundaries() {
	_dgggf := _gaca * _fcea._ddef
	for _beca, _cgfgd := range _fcea._egad[1:] {
		if _dgecg(_cgfgd, _fcea._egad[_beca]) >= _dgggf {
			_cgfgd._egcbd = true
		}
	}
}
func _cdg(_cga []*textWord, _cacg float64, _bgc, _dbfce rulingList) *wordBag {
	_ggca := _gcfa(_cga[0], _cacg, _bgc, _dbfce)
	for _, _ddd := range _cga[1:] {
		_geeg := _feef(_ddd._ecgcg)
		_ggca._bbf[_geeg] = append(_ggca._bbf[_geeg], _ddd)
		_ggca.PdfRectangle = _effa(_ggca.PdfRectangle, _ddd.PdfRectangle)
	}
	_ggca.sort()
	return _ggca
}
func (_abfa *textObject) renderText(_dea _bg.PdfObject, _gffg []byte) error {
	if _abfa._cdbc {
		_ag.Log.Debug("\u0072\u0065\u006e\u0064\u0065r\u0054\u0065\u0078\u0074\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0066\u006f\u006e\u0074\u002e\u0020\u004e\u006f\u0074\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u002e")
		return nil
	}
	_dgb := _abfa.getCurrentFont()
	_aedd := _dgb.BytesToCharcodes(_gffg)
	_cfad, _eafg, _dggf := _dgb.CharcodesToStrings(_aedd)
	if _dggf > 0 {
		_ag.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _eafg, _dggf)
	}
	_abfa._degf._fce += _eafg
	_abfa._degf._ddga += _dggf
	_gadf := _abfa._degf
	_eadf := _gadf._adfb
	_dgdb := _gadf._gdef / 100.0
	_ecgg := _edec
	if _dgb.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_ecgg = 1
	}
	_dce, _eadb := _dgb.GetRuneMetrics(' ')
	if !_eadb {
		_dce, _eadb = _dgb.GetCharMetrics(32)
	}
	if !_eadb {
		_dce, _ = _bf.DefaultFont().GetRuneMetrics(' ')
	}
	_cadc := _dce.Wx * _ecgg
	_ag.Log.Trace("\u0073p\u0061\u0063e\u0057\u0069\u0064t\u0068\u003d\u0025\u002e\u0032\u0066\u0020t\u0065\u0078\u0074\u003d\u0025\u0071 \u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0066\u006f\u006et\u0053\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066", _cadc, _cfad, _dgb, _eadf)
	_fgd := _ee.NewMatrix(_eadf*_dgdb, 0, 0, _eadf, 0, _gadf._bacc)
	if _dbde {
		_ag.Log.Info("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0074\u0065\u0078t\u0073\u003d\u0025\u0071", len(_aedd), _aedd, _cfad)
	}
	_ag.Log.Trace("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0072\u0075\u006ee\u0073\u003d\u0025\u0071", len(_aedd), _aedd, len(_cfad))
	_eee := _abfa.getFillColor()
	_cfcf := _abfa.getStrokeColor()
	for _bbda, _edfb := range _cfad {
		_efge := []rune(_edfb)
		if len(_efge) == 1 && _efge[0] == '\x00' {
			continue
		}
		_fbff := _aedd[_bbda]
		_dccg := _abfa._fcdg.CTM.Mult(_abfa._gcb).Mult(_fgd)
		_gbe := 0.0
		if len(_efge) == 1 && _efge[0] == 32 {
			_gbe = _gadf._gab
		}
		_dagg, _eaab := _dgb.GetCharMetrics(_fbff)
		if !_eaab {
			_ag.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u004e\u006f \u006d\u0065\u0074r\u0069\u0063\u0020\u0066\u006f\u0072\u0020\u0063\u006fde\u003d\u0025\u0064 \u0072\u003d0\u0078\u0025\u0030\u0034\u0078\u003d%\u002b\u0071 \u0025\u0073", _fbff, _efge, _efge, _dgb)
			return _gdf.Errorf("\u006e\u006f\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073:\u0020f\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0063\u006f\u0064\u0065\u003d\u0025\u0064", _dgb.String(), _fbff)
		}
		_aeab := _ee.Point{X: _dagg.Wx * _ecgg, Y: _dagg.Wy * _ecgg}
		_dfd := _ee.Point{X: (_aeab.X*_eadf + _gbe) * _dgdb}
		_fbed := _ee.Point{X: (_aeab.X*_eadf + _gadf._bfcb + _gbe) * _dgdb}
		if _dbde {
			_ag.Log.Info("\u0074\u0066\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0063\u003d\u0025\u002e\u0032f\u0020t\u0077\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _eadf, _gadf._bfcb, _gadf._gab, _dgdb)
			_ag.Log.Info("\u0064x\u002c\u0064\u0079\u003d%\u002e\u0033\u0066\u0020\u00740\u003d%\u002e3\u0066\u0020\u0074\u003d\u0025\u002e\u0033f", _aeab, _dfd, _fbed)
		}
		_gdcc := _def(_dfd)
		_feec := _def(_fbed)
		_bfcd := _abfa._fcdg.CTM.Mult(_abfa._gcb).Mult(_gdcc)
		if _dbcf {
			_ag.Log.Info("e\u006e\u0064\u003a\u000a\tC\u0054M\u003d\u0025\u0073\u000a\u0009 \u0074\u006d\u003d\u0025\u0073\u000a"+"\u0009\u0020t\u0064\u003d\u0025s\u0020\u0078\u006c\u0061\u0074\u003d\u0025\u0073\u000a"+"\u0009t\u0064\u0030\u003d\u0025s\u000a\u0009\u0020\u0020\u2192 \u0025s\u0020x\u006c\u0061\u0074\u003d\u0025\u0073", _abfa._fcdg.CTM, _abfa._gcb, _feec, _ecc(_abfa._fcdg.CTM.Mult(_abfa._gcb).Mult(_feec)), _gdcc, _bfcd, _ecc(_bfcd))
		}
		_fcg, _dcab := _abfa.newTextMark(_ab.ExpandLigatures(_efge), _dccg, _ecc(_bfcd), _b.Abs(_cadc*_dccg.ScalingFactorX()), _dgb, _abfa._degf._bfcb, _eee, _cfcf, _dea, _cfad, _bbda)
		if !_dcab {
			_ag.Log.Debug("\u0054\u0065\u0078\u0074\u0020\u006d\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069d\u0065 \u0070\u0061\u0067\u0065\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		if _dgb == nil {
			_ag.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u002e")
		} else if _dgb.Encoder() == nil {
			_ag.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020N\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u002e\u0020\u0066o\u006et\u003d\u0025\u0073", _dgb)
		} else {
			if _ddf, _caea := _dgb.Encoder().CharcodeToRune(_fbff); _caea {
				_fcg._abcec = string(_ddf)
			}
		}
		_ag.Log.Trace("i\u003d\u0025\u0064\u0020\u0063\u006fd\u0065\u003d\u0025\u0064\u0020\u006d\u0061\u0072\u006b=\u0025\u0073\u0020t\u0072m\u003d\u0025\u0073", _bbda, _fbff, _fcg, _dccg)
		_abfa._dbdg = append(_abfa._dbdg, &_fcg)
		_abfa._gcb.Concat(_feec)
	}
	return nil
}
func _ecc(_efed _ee.Matrix) _ee.Point {
	_fdff, _efda := _efed.Translation()
	return _ee.Point{X: _fdff, Y: _efda}
}

type rulingList []*ruling
type cachedImage struct {
	_fbc *_bf.Image
	_aac _bf.PdfColorspace
}
type shapesState struct {
	_egb  _ee.Matrix
	_efgg _ee.Matrix
	_cfgd []*subpath
	_ecd  bool
	_decg _ee.Point
	_gdcg *textObject
}

func (_abd *shapesState) quadraticTo(_ddfe, _cbca, _cbb, _bggg float64) {
	if _bccf {
		_ag.Log.Info("\u0071\u0075\u0061d\u0072\u0061\u0074\u0069\u0063\u0054\u006f\u003a")
	}
	_abd.addPoint(_cbb, _bggg)
}
func (_deab paraList) writeText(_efeag _ed.Writer) {
	for _dbgc, _abdg := range _deab {
		if _abdg._dgggff {
			continue
		}
		_abdg.writeText(_efeag)
		if _dbgc != len(_deab)-1 {
			if _ceabf(_abdg, _deab[_dbgc+1]) {
				_efeag.Write([]byte("\u0020"))
			} else {
				_efeag.Write([]byte("\u000a"))
				_efeag.Write([]byte("\u000a"))
			}
		}
	}
	_efeag.Write([]byte("\u000a"))
	_efeag.Write([]byte("\u000a"))
}
func _efd(_bgdb *_da.ContentStreamOperation) (float64, error) {
	if len(_bgdb.Params) != 1 {
		_agc := _g.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _bgdb.Operand, 1, len(_bgdb.Params), _bgdb.Params)
		return 0.0, _agc
	}
	return _bg.GetNumberAsFloat(_bgdb.Params[0])
}
func _gfbg(_dbcge map[float64]gridTile) []float64 {
	_begd := make([]float64, 0, len(_dbcge))
	for _eecf := range _dbcge {
		_begd = append(_begd, _eecf)
	}
	_f.Float64s(_begd)
	return _begd
}
func _bcdg(_bbe *wordBag, _gbgdb float64, _ffag, _ebc rulingList) []*wordBag {
	var _dbbgc []*wordBag
	for _, _bacfg := range _bbe.depthIndexes() {
		_dfab := false
		for !_bbe.empty(_bacfg) {
			_eggdf := _bbe.firstReadingIndex(_bacfg)
			_gdbe := _bbe.firstWord(_eggdf)
			_bdcfd := _gcfa(_gdbe, _gbgdb, _ffag, _ebc)
			_bbe.removeWord(_gdbe, _eggdf)
			if _gfcbc {
				_ag.Log.Info("\u0066\u0069\u0072\u0073\u0074\u0057\u006f\u0072\u0064\u0020\u005e\u005e^\u005e\u0020\u0025\u0073", _gdbe.String())
			}
			for _ecaf := true; _ecaf; _ecaf = _dfab {
				_dfab = false
				_ccdd := _efef * _bdcfd._gaed
				_aegf := _cfede * _bdcfd._gaed
				_efce := _cccf * _bdcfd._gaed
				if _gfcbc {
					_ag.Log.Info("\u0070a\u0072a\u0057\u006f\u0072\u0064\u0073\u0020\u0064\u0065\u0070\u0074\u0068 \u0025\u002e\u0032\u0066 \u002d\u0020\u0025\u002e\u0032f\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061\u0044\u0065\u0070\u0074\u0068\u0047\u0061\u0070\u003d\u0025\u002e\u0032\u0066\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061R\u0065\u0061\u0064\u0069\u006e\u0067\u0047\u0061p\u003d\u0025\u002e\u0032\u0066", _bdcfd.minDepth(), _bdcfd.maxDepth(), _efce, _aegf)
				}
				if _bbe.scanBand("\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c", _bdcfd, _aecc(_cdag, 0), _bdcfd.minDepth()-_efce, _bdcfd.maxDepth()+_efce, _aaff, false, false) > 0 {
					_dfab = true
				}
				if _bbe.scanBand("\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _bdcfd, _aecc(_cdag, _aegf), _bdcfd.minDepth(), _bdcfd.maxDepth(), _bdea, false, false) > 0 {
					_dfab = true
				}
				if _dfab {
					continue
				}
				_cbef := _bbe.scanBand("", _bdcfd, _aecc(_ggdb, _ccdd), _bdcfd.minDepth(), _bdcfd.maxDepth(), _bee, true, false)
				if _cbef > 0 {
					_cfae := (_bdcfd.maxDepth() - _bdcfd.minDepth()) / _bdcfd._gaed
					if (_cbef > 1 && float64(_cbef) > 0.3*_cfae) || _cbef <= 10 {
						if _bbe.scanBand("\u006f\u0074\u0068e\u0072", _bdcfd, _aecc(_ggdb, _ccdd), _bdcfd.minDepth(), _bdcfd.maxDepth(), _bee, false, true) > 0 {
							_dfab = true
						}
					}
				}
			}
			_dbbgc = append(_dbbgc, _bdcfd)
		}
	}
	return _dbbgc
}

// String returns a string describing `pt`.
func (_gafc PageText) String() string {
	_ggee := _gdf.Sprintf("P\u0061\u0067\u0065\u0054ex\u0074:\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073", len(_gafc._cac))
	_aeg := []string{"\u002d" + _ggee}
	for _, _ccgf := range _gafc._cac {
		_aeg = append(_aeg, _ccgf.String())
	}
	_aeg = append(_aeg, "\u002b"+_ggee)
	return _d.Join(_aeg, "\u000a")
}

type textState struct {
	_bfcb float64
	_gab  float64
	_gdef float64
	_acfa float64
	_adfb float64
	_egea RenderMode
	_bacc float64
	_dde  *_bf.PdfFont
	_dadc _bf.PdfRectangle
	_fce  int
	_ddga int
}

func _affac(_afcd []TextMark, _fedf *int, _bcgaa TextMark) []TextMark {
	_bcgaa.Offset = *_fedf
	_afcd = append(_afcd, _bcgaa)
	*_fedf += len(_bcgaa.Text)
	return _afcd
}
func (_efbb paraList) reorder(_gdgdf []int) {
	_afeg := make(paraList, len(_efbb))
	for _bbdc, _fgb := range _gdgdf {
		_afeg[_bbdc] = _efbb[_fgb]
	}
	copy(_efbb, _afeg)
}
func (_daee *textTable) computeBbox() _bf.PdfRectangle {
	var _cffbe _bf.PdfRectangle
	_gcbd := false
	for _bbgc := 0; _bbgc < _daee._cabfg; _bbgc++ {
		for _dadcf := 0; _dadcf < _daee._baee; _dadcf++ {
			_bafeb := _daee.get(_dadcf, _bbgc)
			if _bafeb == nil {
				continue
			}
			if !_gcbd {
				_cffbe = _bafeb.PdfRectangle
				_gcbd = true
			} else {
				_cffbe = _effa(_cffbe, _bafeb.PdfRectangle)
			}
		}
	}
	return _cffbe
}

type imageExtractContext struct {
	_dba []ImageMark
	_gcc int
	_fe  int
	_cc  int
	_gda map[*_bg.PdfObjectStream]*cachedImage
	_gfa *ImageExtractOptions
}
type intSet map[int]struct{}

func (_fdbg paraList) addNeighbours() {
	_ggcaf := func(_ggcab []int, _decc *textPara) ([]*textPara, []*textPara) {
		_baaeg := make([]*textPara, 0, len(_ggcab)-1)
		_cddfc := make([]*textPara, 0, len(_ggcab)-1)
		for _, _ffcca := range _ggcab {
			_egeab := _fdbg[_ffcca]
			if _egeab.Urx <= _decc.Llx {
				_baaeg = append(_baaeg, _egeab)
			} else if _egeab.Llx >= _decc.Urx {
				_cddfc = append(_cddfc, _egeab)
			}
		}
		return _baaeg, _cddfc
	}
	_fedg := func(_fdda []int, _cfde *textPara) ([]*textPara, []*textPara) {
		_dddf := make([]*textPara, 0, len(_fdda)-1)
		_cfeb := make([]*textPara, 0, len(_fdda)-1)
		for _, _dadfdg := range _fdda {
			_fefcd := _fdbg[_dadfdg]
			if _fefcd.Ury <= _cfde.Lly {
				_cfeb = append(_cfeb, _fefcd)
			} else if _fefcd.Lly >= _cfde.Ury {
				_dddf = append(_dddf, _fefcd)
			}
		}
		return _dddf, _cfeb
	}
	_gdbfa := _fdbg.yNeighbours(_cfeg)
	for _, _cabc := range _fdbg {
		_aeegg := _gdbfa[_cabc]
		if len(_aeegg) == 0 {
			continue
		}
		_dgbf, _ddgc := _ggcaf(_aeegg, _cabc)
		if len(_dgbf) == 0 && len(_ddgc) == 0 {
			continue
		}
		if len(_dgbf) > 0 {
			_cgcbcc := _dgbf[0]
			for _, _fabb := range _dgbf[1:] {
				if _fabb.Urx >= _cgcbcc.Urx {
					_cgcbcc = _fabb
				}
			}
			for _, _cedce := range _dgbf {
				if _cedce != _cgcbcc && _cedce.Urx > _cgcbcc.Llx {
					_cgcbcc = nil
					break
				}
			}
			if _cgcbcc != nil && _efeg(_cabc.PdfRectangle, _cgcbcc.PdfRectangle) {
				_cabc._fgeg = _cgcbcc
			}
		}
		if len(_ddgc) > 0 {
			_fffbb := _ddgc[0]
			for _, _fbaf := range _ddgc[1:] {
				if _fbaf.Llx <= _fffbb.Llx {
					_fffbb = _fbaf
				}
			}
			for _, _bdfdga := range _ddgc {
				if _bdfdga != _fffbb && _bdfdga.Llx < _fffbb.Urx {
					_fffbb = nil
					break
				}
			}
			if _fffbb != nil && _efeg(_cabc.PdfRectangle, _fffbb.PdfRectangle) {
				_cabc._bfdc = _fffbb
			}
		}
	}
	_gdbfa = _fdbg.xNeighbours(_fbcc)
	for _, _dbdac := range _fdbg {
		_bcad := _gdbfa[_dbdac]
		if len(_bcad) == 0 {
			continue
		}
		_gaegg, _dadg := _fedg(_bcad, _dbdac)
		if len(_gaegg) == 0 && len(_dadg) == 0 {
			continue
		}
		if len(_dadg) > 0 {
			_agae := _dadg[0]
			for _, _egce := range _dadg[1:] {
				if _egce.Ury >= _agae.Ury {
					_agae = _egce
				}
			}
			for _, _acaf := range _dadg {
				if _acaf != _agae && _acaf.Ury > _agae.Lly {
					_agae = nil
					break
				}
			}
			if _agae != nil && _eagg(_dbdac.PdfRectangle, _agae.PdfRectangle) {
				_dbdac._egec = _agae
			}
		}
		if len(_gaegg) > 0 {
			_aace := _gaegg[0]
			for _, _ebbg := range _gaegg[1:] {
				if _ebbg.Lly <= _aace.Lly {
					_aace = _ebbg
				}
			}
			for _, _fgefa := range _gaegg {
				if _fgefa != _aace && _fgefa.Lly < _aace.Ury {
					_aace = nil
					break
				}
			}
			if _aace != nil && _eagg(_dbdac.PdfRectangle, _aace.PdfRectangle) {
				_dbdac._adad = _aace
			}
		}
	}
	for _, _bbggg := range _fdbg {
		if _bbggg._fgeg != nil && _bbggg._fgeg._bfdc != _bbggg {
			_bbggg._fgeg = nil
		}
		if _bbggg._adad != nil && _bbggg._adad._egec != _bbggg {
			_bbggg._adad = nil
		}
		if _bbggg._bfdc != nil && _bbggg._bfdc._fgeg != _bbggg {
			_bbggg._bfdc = nil
		}
		if _bbggg._egec != nil && _bbggg._egec._adad != _bbggg {
			_bbggg._egec = nil
		}
	}
}

// PageImages represents extracted images on a PDF page with spatial information:
// display position and size.
type PageImages struct{ Images []ImageMark }

func _fccec(_fgdcd float64, _edc int) int {
	if _edc == 0 {
		_edc = 1
	}
	_ccag := float64(_edc)
	return int(_b.Round(_fgdcd/_ccag) * _ccag)
}
func (_ecdd *wordBag) sort() {
	for _, _adc := range _ecdd._bbf {
		_f.Slice(_adc, func(_fgdc, _bbcf int) bool { return _addb(_adc[_fgdc], _adc[_bbcf]) < 0 })
	}
}
func (_ffd *textObject) getCurrentFont() *_bf.PdfFont {
	_ggf := _ffd._degf._dde
	if _ggf == nil {
		_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020U\u0073\u0069\u006e\u0067\u0020d\u0065\u0066a\u0075\u006c\u0074\u002e")
		return _bf.DefaultFont()
	}
	return _ggf
}
func (_cgcb rulingList) log(_fadcg string) {
	if !_efa {
		return
	}
	_ag.Log.Info("\u0023\u0023\u0023\u0020\u0025\u0031\u0030\u0073\u003a\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _fadcg, _cgcb.String())
	for _cdbgc, _gfgc := range _cgcb {
		_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cdbgc, _gfgc.String())
	}
}

// PageFonts represents extracted fonts on a PDF page.
type PageFonts struct{ Fonts []Font }

func (_feece *textWord) bbox() _bf.PdfRectangle { return _feece.PdfRectangle }
func (_bdfc lineRuling) xMean() float64         { return 0.5 * (_bdfc._abec.X + _bdfc._eacb.X) }
func _fbcac(_fdfc []*textWord, _eceda *textWord) []*textWord {
	for _ggbgg, _gcfe := range _fdfc {
		if _gcfe == _eceda {
			return _bbcg(_fdfc, _ggbgg)
		}
	}
	_ag.Log.Error("\u0072\u0065\u006d\u006f\u0076e\u0057\u006f\u0072\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u0073\u0020\u0064o\u0065\u0073\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", _eceda)
	return nil
}
func (_ddcgf intSet) add(_cgfc int) { _ddcgf[_cgfc] = struct{}{} }
func (_cgbg *stateStack) size() int { return len(*_cgbg) }

// PageText represents the layout of text on a device page.
type PageText struct {
	_cac  []*textMark
	_gcea string
	_cagb []TextMark
	_ece  []TextTable
	_fabf _bf.PdfRectangle
	_fcdb []pathSection
	_cddd []pathSection
	_bfg  *_da.ContentStreamOperations
}

func (_gaef pathSection) bbox() _bf.PdfRectangle {
	_dgbe := _gaef._gdbf[0]._fcdc[0]
	_bbc := _bf.PdfRectangle{Llx: _dgbe.X, Urx: _dgbe.X, Lly: _dgbe.Y, Ury: _dgbe.Y}
	_aca := func(_fgac _ee.Point) {
		if _fgac.X < _bbc.Llx {
			_bbc.Llx = _fgac.X
		} else if _fgac.X > _bbc.Urx {
			_bbc.Urx = _fgac.X
		}
		if _fgac.Y < _bbc.Lly {
			_bbc.Lly = _fgac.Y
		} else if _fgac.Y > _bbc.Ury {
			_bbc.Ury = _fgac.Y
		}
	}
	for _, _bdfd := range _gaef._gdbf[0]._fcdc[1:] {
		_aca(_bdfd)
	}
	for _, _cbgd := range _gaef._gdbf[1:] {
		for _, _accge := range _cbgd._fcdc {
			_aca(_accge)
		}
	}
	return _bbc
}
func (_gdcce compositeCell) String() string {
	_gaab := ""
	if len(_gdcce.paraList) > 0 {
		_gaab = _babd(_gdcce.paraList.merge().text(), 50)
	}
	return _gdf.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0070\u0061\u0072a\u0073\u0020\u0025\u0071", _gdcce.PdfRectangle, len(_gdcce.paraList), _gaab)
}
func (_cgfab *wordBag) minDepth() float64 { return _cgfab._fcgc - (_cgfab.Ury - _cgfab._gaed) }
func (_aeeg *textTable) emptyCompositeRow(_baccb int) bool {
	for _cfea := 0; _cfea < _aeeg._baee; _cfea++ {
		if _eddbb, _bfbbb := _aeeg._bbfb[_fcbc(_cfea, _baccb)]; _bfbbb {
			if len(_eddbb.paraList) > 0 {
				return false
			}
		}
	}
	return true
}
func (_fggg rulingList) connections(_gcaf map[int]intSet, _cdadf int) intSet {
	_fbd := make(intSet)
	_bfbd := make(intSet)
	var _ecbe func(int)
	_ecbe = func(_deag int) {
		if !_bfbd.has(_deag) {
			_bfbd.add(_deag)
			for _ggeee := range _fggg {
				if _gcaf[_ggeee].has(_deag) {
					_fbd.add(_ggeee)
				}
			}
			for _bdef := range _fggg {
				if _fbd.has(_bdef) {
					_ecbe(_bdef)
				}
			}
		}
	}
	_ecbe(_cdadf)
	return _fbd
}
func (_cced *textPara) writeCellText(_bdba _ed.Writer) {
	for _debg, _gbge := range _cced._ecbdg {
		_fgfe := _gbge.text()
		_bcfeg := _badf && _gbge.endsInHyphen() && _debg != len(_cced._ecbdg)-1
		if _bcfeg {
			_fgfe = _fafd(_fgfe)
		}
		_bdba.Write([]byte(_fgfe))
		if !(_bcfeg || _debg == len(_cced._ecbdg)-1) {
			_bdba.Write([]byte(_afed(_gbge._dcfd, _cced._ecbdg[_debg+1]._dcfd)))
		}
	}
}
func (_addg *textTable) reduceTiling(_bfge gridTiling, _eebc float64) *textTable {
	_eage := make([]int, 0, _addg._cabfg)
	_ebadc := make([]int, 0, _addg._baee)
	_efec := _bfge._fdeac
	_dcecd := _bfge._gccf
	for _dedcg := 0; _dedcg < _addg._cabfg; _dedcg++ {
		_gccdb := _dedcg > 0 && _b.Abs(_dcecd[_dedcg-1]-_dcecd[_dedcg]) < _eebc && _addg.emptyCompositeRow(_dedcg)
		if !_gccdb {
			_eage = append(_eage, _dedcg)
		}
	}
	for _cfcb := 0; _cfcb < _addg._baee; _cfcb++ {
		_fgbbb := _cfcb < _addg._baee-1 && _b.Abs(_efec[_cfcb+1]-_efec[_cfcb]) < _eebc && _addg.emptyCompositeColumn(_cfcb)
		if !_fgbbb {
			_ebadc = append(_ebadc, _cfcb)
		}
	}
	if len(_eage) == _addg._cabfg && len(_ebadc) == _addg._baee {
		return _addg
	}
	_ddfea := textTable{_acgbd: _addg._acgbd, _baee: len(_ebadc), _cabfg: len(_eage), _bbfb: make(map[uint64]compositeCell, len(_ebadc)*len(_eage))}
	if _beae {
		_ag.Log.Info("\u0072\u0065\u0064\u0075c\u0065\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0025d\u0078%\u0064\u0020\u002d\u003e\u0020\u0025\u0064x\u0025\u0064", _addg._baee, _addg._cabfg, len(_ebadc), len(_eage))
		_ag.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _ebadc)
		_ag.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _eage)
	}
	for _dccb, _ecgge := range _eage {
		for _bbcb, _fdeaae := range _ebadc {
			_ccacb, _ffabe := _addg.getComposite(_fdeaae, _ecgge)
			if len(_ccacb) == 0 {
				continue
			}
			if _beae {
				_gdf.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _bbcb, _dccb, _fdeaae, _ecgge, _babd(_ccacb.merge().text(), 50))
			}
			_ddfea.putComposite(_bbcb, _dccb, _ccacb, _ffabe)
		}
	}
	return &_ddfea
}
func _bbcg(_ggefba []*textWord, _cfdd int) []*textWord {
	_dfcdg := len(_ggefba)
	copy(_ggefba[_cfdd:], _ggefba[_cfdd+1:])
	return _ggefba[:_dfcdg-1]
}
func (_ffca *textTable) bbox() _bf.PdfRectangle { return _ffca.PdfRectangle }
func (_dfccd *textObject) setHorizScaling(_ffcg float64) {
	if _dfccd == nil {
		return
	}
	_dfccd._degf._gdef = _ffcg
}
func _bceab(_daca []int) []int {
	_egafe := make([]int, len(_daca))
	for _eddf, _fbfe := range _daca {
		_egafe[len(_daca)-1-_eddf] = _fbfe
	}
	return _egafe
}
func (_efdcf paraList) lines() []*textLine {
	var _fbab []*textLine
	for _, _aedee := range _efdcf {
		_fbab = append(_fbab, _aedee._ecbdg...)
	}
	return _fbab
}
func _ggfdc(_gdea []_bg.PdfObject) (_fgdce, _abfg float64, _dadcfd error) {
	if len(_gdea) != 2 {
		return 0, 0, _gdf.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0073\u003a \u0025\u0064", len(_gdea))
	}
	_aefc, _dadcfd := _bg.GetNumbersAsFloat(_gdea)
	if _dadcfd != nil {
		return 0, 0, _dadcfd
	}
	return _aefc[0], _aefc[1], nil
}
func (_efdcb *textObject) getFontDirect(_aafd string) (*_bf.PdfFont, error) {
	_eecg, _dbg := _efdcb.getFontDict(_aafd)
	if _dbg != nil {
		return nil, _dbg
	}
	_ggaa, _dbg := _bf.NewPdfFontFromPdfObject(_eecg)
	if _dbg != nil {
		_ag.Log.Debug("\u0067\u0065\u0074\u0046\u006f\u006e\u0074\u0044\u0069\u0072\u0065\u0063\u0074\u003a\u0020\u004e\u0065\u0077Pd\u0066F\u006f\u006e\u0074\u0046\u0072\u006f\u006d\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d%\u0023\u0071\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _aafd, _dbg)
	}
	return _ggaa, _dbg
}
func (_egbb *subpath) add(_geeb ..._ee.Point) { _egbb._fcdc = append(_egbb._fcdc, _geeb...) }

// String returns a description of `state`.
func (_gbg *textState) String() string {
	_eef := "\u005bN\u004f\u0054\u0020\u0053\u0045\u0054]"
	if _gbg._dde != nil {
		_eef = _gbg._dde.BaseFont()
	}
	return _gdf.Sprintf("\u0074\u0063\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0077\u003d\u0025\u002e\u0032\u0066 \u0074f\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071", _gbg._bfcb, _gbg._gab, _gbg._adfb, _eef)
}
func (_ccaad gridTile) complete() bool { return _ccaad.numBorders() == 4 }
func (_dfdg rulingList) removeDuplicates() rulingList {
	if len(_dfdg) == 0 {
		return nil
	}
	_dfdg.sort()
	_edbc := rulingList{_dfdg[0]}
	for _, _fagf := range _dfdg[1:] {
		if _fagf.equals(_edbc[len(_edbc)-1]) {
			continue
		}
		_edbc = append(_edbc, _fagf)
	}
	return _edbc
}
func (_baac paraList) llyOrdering() []int {
	_ccddd := make([]int, len(_baac))
	for _bcfd := range _baac {
		_ccddd[_bcfd] = _bcfd
	}
	_f.SliceStable(_ccddd, func(_abgb, _fdbe int) bool {
		_bba, _eddb := _ccddd[_abgb], _ccddd[_fdbe]
		return _baac[_bba].Lly < _baac[_eddb].Lly
	})
	return _ccddd
}
func (_babb *wordBag) allWords() []*textWord {
	var _gdeec []*textWord
	for _, _aeba := range _babb._bbf {
		_gdeec = append(_gdeec, _aeba...)
	}
	return _gdeec
}
func _bddc(_bcdb _ee.Point) *subpath { return &subpath{_fcdc: []_ee.Point{_bcdb}} }
func (_ebgd *wordBag) absorb(_dcae *wordBag) {
	_dcgd := _dcae.makeRemovals()
	for _ffgb, _fgdf := range _dcae._bbf {
		for _, _gcd := range _fgdf {
			_ebgd.pullWord(_gcd, _ffgb, _dcgd)
		}
	}
	_dcae.applyRemovals(_dcgd)
}
func _bega(_ebcf, _bdfef _ee.Point) rulingKind {
	_affd := _b.Abs(_ebcf.X - _bdfef.X)
	_fggbg := _b.Abs(_ebcf.Y - _bdfef.Y)
	return _bfcbe(_affd, _fggbg, _ffdeb)
}
func (_acfdb rectRuling) checkWidth(_cegg, _ceggc float64) (float64, bool) {
	_ggdba := _ceggc - _cegg
	_dbcga := _ggdba <= _dcfe
	return _ggdba, _dbcga
}
func (_bfcg *imageExtractContext) extractInlineImage(_eaf *_da.ContentStreamInlineImage, _cb _da.GraphicsState, _acb *_bf.PdfPageResources) error {
	_eeg, _dad := _eaf.ToImage(_acb)
	if _dad != nil {
		return _dad
	}
	_baf, _dad := _eaf.GetColorSpace(_acb)
	if _dad != nil {
		return _dad
	}
	if _baf == nil {
		_baf = _bf.NewPdfColorspaceDeviceGray()
	}
	_ffc, _dad := _baf.ImageToRGB(*_eeg)
	if _dad != nil {
		return _dad
	}
	_ae := ImageMark{Image: &_ffc, Width: _cb.CTM.ScalingFactorX(), Height: _cb.CTM.ScalingFactorY(), Angle: _cb.CTM.Angle()}
	_ae.X, _ae.Y = _cb.CTM.Translation()
	_bfcg._dba = append(_bfcg._dba, _ae)
	_bfcg._gcc++
	return nil
}

type textObject struct {
	_ccae *Extractor
	_caa  *_bf.PdfPageResources
	_fcdg _da.GraphicsState
	_degf *textState
	_afg  *stateStack
	_gcb  _ee.Matrix
	_baff _ee.Matrix
	_dbdg []*textMark
	_cdbc bool
}

func (_ddfg *ruling) alignsPrimary(_feede *ruling) bool {
	return _ddfg._cgef == _feede._cgef && _b.Abs(_ddfg._eead-_feede._eead) < _dcfe*0.5
}
func (_gcgbf *textTable) get(_cabad, _facd int) *textPara { return _gcgbf._aagc[_fcbc(_cabad, _facd)] }

// ExtractPageImages returns the image contents of the page extractor, including data
// and position, size information for each image.
// A set of options to control page image extraction can be passed in. The options
// parameter can be nil for the default options. By default, inline stencil masks
// are not extracted.
func (_adde *Extractor) ExtractPageImages(options *ImageExtractOptions) (*PageImages, error) {
	_aaa := &imageExtractContext{_gfa: options}
	_ced := _aaa.extractContentStreamImages(_adde._fc, _adde._ad)
	if _ced != nil {
		return nil, _ced
	}
	return &PageImages{Images: _aaa._dba}, nil
}
func _eagab(_cebgb []compositeCell) []float64 {
	var _acgad []*textLine
	_ebgc := 0
	for _, _acfac := range _cebgb {
		_ebgc += len(_acfac.paraList)
		_acgad = append(_acgad, _acfac.lines()...)
	}
	_f.Slice(_acgad, func(_gcde, _edgc int) bool {
		_dfca, _abee := _acgad[_gcde], _acgad[_edgc]
		_agbd, _faca := _dfca._dcfd, _abee._dcfd
		if !_fbga(_agbd - _faca) {
			return _agbd < _faca
		}
		return _dfca.Llx < _abee.Llx
	})
	if _beae {
		_gdf.Printf("\u0020\u0020\u0020 r\u006f\u0077\u0042\u006f\u0072\u0064\u0065\u0072\u0073:\u0020%\u0064 \u0070a\u0072\u0061\u0073\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", _ebgc, len(_acgad))
		for _eeaa, _dadca := range _acgad {
			_gdf.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _eeaa, _dadca)
		}
	}
	var _bbad []float64
	_cadad := _acgad[0]
	var _gfdf [][]*textLine
	_bdde := []*textLine{_cadad}
	for _bgdbd, _eabf := range _acgad[1:] {
		if _eabf.Ury < _cadad.Lly {
			_agecd := 0.5 * (_eabf.Ury + _cadad.Lly)
			if _beae {
				_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u003c\u0020\u0025\u0036.\u0032f\u0020\u0062\u006f\u0072\u0064\u0065\u0072\u003d\u0025\u0036\u002e\u0032\u0066\u000a"+"\u0009\u0020\u0071\u003d\u0025\u0073\u000a\u0009\u0020p\u003d\u0025\u0073\u000a", _bgdbd, _eabf.Ury, _cadad.Lly, _agecd, _cadad, _eabf)
			}
			_bbad = append(_bbad, _agecd)
			_gfdf = append(_gfdf, _bdde)
			_bdde = nil
		}
		_bdde = append(_bdde, _eabf)
		if _eabf.Lly < _cadad.Lly {
			_cadad = _eabf
		}
	}
	if len(_bdde) > 0 {
		_gfdf = append(_gfdf, _bdde)
	}
	if _beae {
		_gdf.Printf(" \u0020\u0020\u0020\u0020\u0020\u0020 \u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073\u003d%\u0036.\u0032\u0066\u000a", _bbad)
	}
	if _beae {
		_ag.Log.Info("\u0072\u006f\u0077\u003d\u0025\u0064", len(_cebgb))
		for _eddge, _cdbda := range _cebgb {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _eddge, _cdbda)
		}
		_ag.Log.Info("\u0067r\u006f\u0075\u0070\u0073\u003d\u0025d", len(_gfdf))
		for _ddgd, _gdfad := range _gfdf {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0064\u000a", _ddgd, len(_gdfad))
			for _faa, _gaebf := range _gdfad {
				_gdf.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _faa, _gaebf)
			}
		}
	}
	_ecdc := true
	for _eed, _ddca := range _gfdf {
		_cbfae := true
		for _gcfg, _adfcb := range _cebgb {
			if _beae {
				_gdf.Printf("\u0020\u0020\u0020\u007e\u007e\u007e\u0067\u0072\u006f\u0075\u0070\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0063\u0065\u006cl\u0020\u0025\u0064\u0020\u006ff\u0020\u0025d\u0020\u0025\u0073\u000a", _eed, len(_gfdf), _gcfg, len(_cebgb), _adfcb)
			}
			if !_adfcb.hasLines(_ddca) {
				if _beae {
					_gdf.Printf("\u0020\u0020\u0020\u0021\u0021\u0021\u0067\u0072\u006f\u0075\u0070\u0020\u0025d\u0020\u006f\u0066\u0020\u0025\u0064 \u0063\u0065\u006c\u006c\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064 \u004f\u0055\u0054\u000a", _eed, len(_gfdf), _gcfg, len(_cebgb))
				}
				_cbfae = false
				break
			}
		}
		if !_cbfae {
			_ecdc = false
			break
		}
	}
	if !_ecdc {
		if _beae {
			_ag.Log.Info("\u0072\u006f\u0077\u0020\u0063o\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u0020\u0064\u006f\u006e\u0027\u0074 \u0073\u0070\u0061\u006e\u0020\u0061\u006c\u006c\u0020\u0063\u0065\u006c\u006c\u0073\u0020\u0069\u006e\u0020\u0072\u006f\u0077\u002e\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
		}
		_bbad = nil
	}
	if _beae && _bbad != nil {
		_gdf.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u002a\u002a*\u0072\u006f\u0077\u0043\u006f\u0072\u0072i\u0064\u006f\u0072\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u000a", _bbad)
	}
	return _bbad
}
func (_cca *textObject) setTextRise(_cgg float64) {
	if _cca == nil {
		return
	}
	_cca._degf._bacc = _cgg
}

// NewWithOptions an Extractor instance for extracting content from the input PDF page with options.
func NewWithOptions(page *_bf.PdfPage, options *Options) (*Extractor, error) {
	_fbe, _bgg := page.GetAllContentStreams()
	if _bgg != nil {
		return nil, _bgg
	}
	_ce, _bgg := page.GetMediaBox()
	if _bgg != nil {
		return nil, _gdf.Errorf("\u0065\u0078\u0074r\u0061\u0063\u0074\u006fr\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u002e\u0020\u0025\u0076", _bgg)
	}
	_af := &Extractor{_fc: _fbe, _ad: page.Resources, _bb: *_ce, _ga: page.CropBox, _ac: map[string]fontEntry{}, _fd: map[string]textResult{}, _bc: options}
	if _af._bb.Llx > _af._bb.Urx {
		_ag.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0058\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _af._bb)
		_af._bb.Llx, _af._bb.Urx = _af._bb.Urx, _af._bb.Llx
	}
	if _af._bb.Lly > _af._bb.Ury {
		_ag.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0059\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _af._bb)
		_af._bb.Lly, _af._bb.Ury = _af._bb.Ury, _af._bb.Lly
	}
	return _af, nil
}
func (_dbggd *textLine) endsInHyphen() bool {
	_eddec := _dbggd._egad[len(_dbggd._egad)-1]
	_geb := _eddec._bcaa
	_bfec, _ccda := _a.DecodeLastRuneInString(_geb)
	if _ccda <= 0 || !_gg.Is(_gg.Hyphen, _bfec) {
		return false
	}
	if _eddec._egcbd && _fef(_geb) {
		return true
	}
	return _fef(_dbggd.text())
}
func (_gadc intSet) del(_ggag int) { delete(_gadc, _ggag) }

const (
	_dcbf markKind = iota
	_caade
	_gaea
	_eddbd
)

// String returns a human readable description of `vecs`.
func (_gedfd rulingList) String() string {
	if len(_gedfd) == 0 {
		return "\u007b \u0045\u004d\u0050\u0054\u0059\u0020}"
	}
	_fead, _edce := _gedfd.vertsHorzs()
	_cggfe := len(_fead)
	_gfdd := len(_edce)
	if _cggfe == 0 || _gfdd == 0 {
		return _gdf.Sprintf("\u007b%\u0064\u0020\u0078\u0020\u0025\u0064}", _cggfe, _gfdd)
	}
	_dcgde := _bf.PdfRectangle{Llx: _fead[0]._eead, Urx: _fead[_cggfe-1]._eead, Lly: _edce[_gfdd-1]._eead, Ury: _edce[0]._eead}
	return _gdf.Sprintf("\u007b\u0025d\u0020\u0078\u0020%\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u007d", _cggfe, _gfdd, _dcgde)
}
func (_gede rulingList) snapToGroupsDirection() rulingList {
	_gede.sortStrict()
	_ggbg := make(map[*ruling]rulingList, len(_gede))
	_bfbb := _gede[0]
	_eacc := func(_cbbe *ruling) { _bfbb = _cbbe; _ggbg[_bfbb] = rulingList{_cbbe} }
	_eacc(_gede[0])
	for _, _dbbbd := range _gede[1:] {
		if _dbbbd._eead < _bfbb._eead-_ffba {
			_ag.Log.Error("\u0073\u006e\u0061\u0070T\u006f\u0047\u0072\u006f\u0075\u0070\u0073\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0057\u0072\u006f\u006e\u0067\u0020\u0070\u0072\u0069\u006da\u0072\u0079\u0020\u006f\u0072d\u0065\u0072\u002e\u000a\u0009\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0020\u0076\u003d\u0025\u0073", _bfbb, _dbbbd)
		}
		if _dbbbd._eead > _bfbb._eead+_dcfe {
			_eacc(_dbbbd)
		} else {
			_ggbg[_bfbb] = append(_ggbg[_bfbb], _dbbbd)
		}
	}
	_bfed := make(map[*ruling]float64, len(_ggbg))
	_agfc := make(map[*ruling]*ruling, len(_gede))
	for _fgadf, _eceef := range _ggbg {
		_bfed[_fgadf] = _eceef.mergePrimary()
		for _, _dbeg := range _eceef {
			_agfc[_dbeg] = _fgadf
		}
	}
	for _, _ebbc := range _gede {
		_ebbc._eead = _bfed[_agfc[_ebbc]]
	}
	_eafe := make(rulingList, 0, len(_gede))
	for _, _aefd := range _ggbg {
		_bdag := _aefd.splitSec()
		for _cfff, _ffgd := range _bdag {
			_eagc := _ffgd.merge()
			if len(_eafe) > 0 {
				_ebff := _eafe[len(_eafe)-1]
				if _ebff.alignsPrimary(_eagc) && _ebff.alignsSec(_eagc) {
					_ag.Log.Error("\u0073\u006e\u0061\u0070\u0054\u006fG\u0072\u006f\u0075\u0070\u0073\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0044\u0075\u0070\u006ci\u0063\u0061\u0074\u0065\u0020\u0069\u003d\u0025\u0064\u000a\u0009\u0077\u003d\u0025s\u000a\t\u0076\u003d\u0025\u0073", _cfff, _ebff, _eagc)
					continue
				}
			}
			_eafe = append(_eafe, _eagc)
		}
	}
	_eafe.sortStrict()
	return _eafe
}
func (_gfdg *wordBag) depthRange(_dbgd, _bcbd int) []int {
	var _ddgg []int
	for _agd := range _gfdg._bbf {
		if _dbgd <= _agd && _agd <= _bcbd {
			_ddgg = append(_ddgg, _agd)
		}
	}
	if len(_ddgg) == 0 {
		return nil
	}
	_f.Ints(_ddgg)
	return _ddgg
}
func (_caad *shapesState) establishSubpath() *subpath {
	_ecfc, _ecdb := _caad.lastpointEstablished()
	if !_ecdb {
		_caad._cfgd = append(_caad._cfgd, _bddc(_ecfc))
	}
	if len(_caad._cfgd) == 0 {
		return nil
	}
	_caad._ecd = false
	return _caad._cfgd[len(_caad._cfgd)-1]
}
func (_ffga *textPara) depth() float64 {
	if _ffga._dgggff {
		return -1.0
	}
	if len(_ffga._ecbdg) > 0 {
		return _ffga._ecbdg[0]._dcfd
	}
	return _ffga._bccg.depth()
}
func _gfdbb(_ggdcb _bf.PdfColorspace, _badeg _bf.PdfColor) _eda.Color {
	if _ggdcb == nil || _badeg == nil {
		return _eda.Black
	}
	_fdca, _efeeg := _ggdcb.ColorToRGB(_badeg)
	if _efeeg != nil {
		_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006fu\u006c\u0064\u0020no\u0074\u0020\u0063\u006f\u006e\u0076e\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0025\u0076\u0020\u0028\u0025\u0076)\u0020\u0074\u006f\u0020\u0052\u0047\u0042\u003a \u0025\u0073", _badeg, _ggdcb, _efeeg)
		return _eda.Black
	}
	_dcaf, _fbbdd := _fdca.(*_bf.PdfColorDeviceRGB)
	if !_fbbdd {
		_ag.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0052\u0047\u0042\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0076", _fdca)
		return _eda.Black
	}
	return _eda.NRGBA{R: uint8(_dcaf.R() * 255), G: uint8(_dcaf.G() * 255), B: uint8(_dcaf.B() * 255), A: uint8(255)}
}
func (_ebgg *textObject) getFontDict(_ecf string) (_cfd _bg.PdfObject, _gec error) {
	_ceac := _ebgg._caa
	if _ceac == nil {
		_ag.Log.Debug("g\u0065\u0074\u0046\u006f\u006e\u0074D\u0069\u0063\u0074\u002e\u0020\u004eo\u0020\u0072\u0065\u0073\u006f\u0075\u0072c\u0065\u0073\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025#\u0071", _ecf)
		return nil, nil
	}
	_cfd, _cfdb := _ceac.GetFontByName(_bg.PdfObjectName(_ecf))
	if !_cfdb {
		_ag.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0067\u0065t\u0046\u006f\u006et\u0044\u0069\u0063\u0074\u003a\u0020\u0046\u006f\u006et \u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020\u006e\u0061m\u0065\u003d%\u0023\u0071", _ecf)
		return nil, _g.New("f\u006f\u006e\u0074\u0020no\u0074 \u0069\u006e\u0020\u0072\u0065s\u006f\u0075\u0072\u0063\u0065\u0073")
	}
	return _cfd, nil
}

// Append appends `mark` to the mark array.
func (_cbe *TextMarkArray) Append(mark TextMark) { _cbe._fceg = append(_cbe._fceg, mark) }
func (_ebbcf paraList) eventNeighbours(_ffbag []event) map[*textPara][]int {
	_f.Slice(_ffbag, func(_caeff, _aadeb int) bool {
		_eaba, _febgc := _ffbag[_caeff], _ffbag[_aadeb]
		_ededb, _adab := _eaba._cega, _febgc._cega
		if _ededb != _adab {
			return _ededb < _adab
		}
		if _eaba._fddb != _febgc._fddb {
			return _eaba._fddb
		}
		return _caeff < _aadeb
	})
	_gbed := make(map[int]intSet)
	_addf := make(intSet)
	for _, _dfba := range _ffbag {
		if _dfba._fddb {
			_gbed[_dfba._dbdfc] = make(intSet)
			for _fddg := range _addf {
				if _fddg != _dfba._dbdfc {
					_gbed[_dfba._dbdfc].add(_fddg)
					_gbed[_fddg].add(_dfba._dbdfc)
				}
			}
			_addf.add(_dfba._dbdfc)
		} else {
			_addf.del(_dfba._dbdfc)
		}
	}
	_bffcc := map[*textPara][]int{}
	for _fbdb, _feffa := range _gbed {
		_ggdd := _ebbcf[_fbdb]
		if len(_feffa) == 0 {
			_bffcc[_ggdd] = nil
			continue
		}
		_fbeec := make([]int, len(_feffa))
		_addc := 0
		for _gdcb := range _feffa {
			_fbeec[_addc] = _gdcb
			_addc++
		}
		_bffcc[_ggdd] = _fbeec
	}
	return _bffcc
}

// String returns a description of `v`.
func (_ffffg *ruling) String() string {
	if _ffffg._cgef == _gdbb {
		return "\u004e\u004f\u0054\u0020\u0052\u0055\u004c\u0049\u004e\u0047"
	}
	_dgac, _cacgd := "\u0078", "\u0079"
	if _ffffg._cgef == _ebabd {
		_dgac, _cacgd = "\u0079", "\u0078"
	}
	_gefff := ""
	if _ffffg._efgd != 0.0 {
		_gefff = _gdf.Sprintf(" \u0077\u0069\u0064\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _ffffg._efgd)
	}
	return _gdf.Sprintf("\u0025\u00310\u0073\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0020\u0025\u0076\u0025\u0073", _ffffg._cgef, _dgac, _ffffg._eead, _cacgd, _ffffg._bgeb, _ffffg._eecc, _ffffg._eecc-_ffffg._bgeb, _ffffg._gbgb, _ffffg.Color, _gefff)
}
func (_bbgaa *subpath) removeDuplicates() {
	if len(_bbgaa._fcdc) == 0 {
		return
	}
	_agg := []_ee.Point{_bbgaa._fcdc[0]}
	for _, _bfb := range _bbgaa._fcdc[1:] {
		if !_egag(_bfb, _agg[len(_agg)-1]) {
			_agg = append(_agg, _bfb)
		}
	}
	_bbgaa._fcdc = _agg
}
func (_efaf *textTable) getRight() paraList {
	_abcg := make(paraList, _efaf._cabfg)
	for _afcdd := 0; _afcdd < _efaf._cabfg; _afcdd++ {
		_gefc := _efaf.get(_efaf._baee-1, _afcdd)._bfdc
		if _gefc.taken() {
			return nil
		}
		_abcg[_afcdd] = _gefc
	}
	for _febfg := 0; _febfg < _efaf._cabfg-1; _febfg++ {
		if _abcg[_febfg]._egec != _abcg[_febfg+1] {
			return nil
		}
	}
	return _abcg
}
func (_feaf paraList) findTextTables() []*textTable {
	var _dfdf []*textTable
	for _, _edbeb := range _feaf {
		if _edbeb.taken() || _edbeb.Width() == 0 {
			continue
		}
		_acade := _edbeb.isAtom()
		if _acade == nil {
			continue
		}
		_acade.growTable()
		if _acade._baee*_acade._cabfg < _gdfd {
			continue
		}
		_acade.markCells()
		_acade.log("\u0067\u0072\u006fw\u006e")
		_dfdf = append(_dfdf, _acade)
	}
	return _dfdf
}
func (_dage gridTile) contains(_gagag _bf.PdfRectangle) bool {
	if _dage.numBorders() < 3 {
		return false
	}
	if _dage._bceg && _gagag.Llx < _dage.Llx-_egfc {
		return false
	}
	if _dage._abdfe && _gagag.Urx > _dage.Urx+_egfc {
		return false
	}
	if _dage._abag && _gagag.Lly < _dage.Lly-_egfc {
		return false
	}
	if _dage._gfga && _gagag.Ury > _dage.Ury+_egfc {
		return false
	}
	return true
}

// ApplyArea processes the page text only within the specified area `bbox`.
// Each time ApplyArea is called, it updates the result set in `pt`.
// Can be called multiple times in a row with different bounding boxes.
func (_dgeb *PageText) ApplyArea(bbox _bf.PdfRectangle) {
	_agcb := make([]*textMark, 0, len(_dgeb._cac))
	for _, _dee := range _dgeb._cac {
		if _gcgg(_dee.bbox(), bbox) {
			_agcb = append(_agcb, _dee)
		}
	}
	var _cggf paraList
	_gbeb := len(_agcb)
	for _cedcf := 0; _cedcf < 360 && _gbeb > 0; _cedcf += 90 {
		_gba := make([]*textMark, 0, len(_agcb)-_gbeb)
		for _, _ggcd := range _agcb {
			if _ggcd._bfbc == _cedcf {
				_gba = append(_gba, _ggcd)
			}
		}
		if len(_gba) > 0 {
			_dcf := _aadf(_gba, _dgeb._fabf, nil, nil)
			_cggf = append(_cggf, _dcf...)
			_gbeb -= len(_gba)
		}
	}
	_afb := new(_gd.Buffer)
	_cggf.writeText(_afb)
	_dgeb._gcea = _afb.String()
	_dgeb._cagb = _cggf.toTextMarks()
	_dgeb._ece = _cggf.tables()
}
func (_gcgb *textLine) text() string {
	var _febb []string
	for _, _ebda := range _gcgb._egad {
		if _ebda._egcbd {
			_febb = append(_febb, "\u0020")
		}
		_febb = append(_febb, _ebda._bcaa)
	}
	return _d.Join(_febb, "")
}
func (_ggcda *shapesState) clearPath() {
	_ggcda._cfgd = nil
	_ggcda._ecd = false
	if _bccf {
		_ag.Log.Info("\u0043\u004c\u0045A\u0052\u003a\u0020\u0073\u0073\u003d\u0025\u0073", _ggcda)
	}
}
func (_adfd *shapesState) cubicTo(_cgbb, _egbg, _eadg, _gdga, _bcbab, _bca float64) {
	if _bccf {
		_ag.Log.Info("\u0063\u0075\u0062\u0069\u0063\u0054\u006f\u003a")
	}
	_adfd.addPoint(_bcbab, _bca)
}

type textPara struct {
	_bf.PdfRectangle
	_dcfc   _bf.PdfRectangle
	_ecbdg  []*textLine
	_bccg   *textTable
	_bggb   bool
	_dgggff bool
	_fgeg   *textPara
	_bfdc   *textPara
	_adad   *textPara
	_egec   *textPara
}

func _gabdb(_bdda []pathSection) rulingList {
	_bcff(_bdda)
	if _efa {
		_ag.Log.Info("\u006d\u0061k\u0065\u0053\u0074\u0072\u006f\u006b\u0065\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0064\u0020\u0073\u0074\u0072ok\u0065\u0073", len(_bdda))
	}
	var _caba rulingList
	for _, _gcgf := range _bdda {
		for _, _cbgf := range _gcgf._gdbf {
			if len(_cbgf._fcdc) < 2 {
				continue
			}
			_cgfd := _cbgf._fcdc[0]
			for _, _cefa := range _cbgf._fcdc[1:] {
				if _gadd, _ebcc := _aeca(_cgfd, _cefa, _gcgf.Color); _ebcc {
					_caba = append(_caba, _gadd)
				}
				_cgfd = _cefa
			}
		}
	}
	if _efa {
		_ag.Log.Info("m\u0061\u006b\u0065\u0053tr\u006fk\u0065\u0052\u0075\u006c\u0069n\u0067\u0073\u003a\u0020\u0025\u0073", _caba)
	}
	return _caba
}

const _edec = 1.0 / 1000.0

func _fef(_bcgg string) bool {
	if _a.RuneCountInString(_bcgg) < _feba {
		return false
	}
	_ffab, _egbf := _a.DecodeLastRuneInString(_bcgg)
	if _egbf <= 0 || !_gg.Is(_gg.Hyphen, _ffab) {
		return false
	}
	_ffab, _egbf = _a.DecodeLastRuneInString(_bcgg[:len(_bcgg)-_egbf])
	return _egbf > 0 && !_gg.IsSpace(_ffab)
}
func (_gegg *textObject) setFont(_cgfg string, _bdcdd float64) error {
	if _gegg == nil {
		return nil
	}
	_gegg._degf._adfb = _bdcdd
	_dedc, _bcc := _gegg.getFont(_cgfg)
	if _bcc != nil {
		return _bcc
	}
	_gegg._degf._dde = _dedc
	return nil
}
func (_fddd *textWord) appendMark(_dcfdg *textMark, _agfb _bf.PdfRectangle) {
	_fddd._gdgg = append(_fddd._gdgg, _dcfdg)
	_fddd.PdfRectangle = _effa(_fddd.PdfRectangle, _dcfdg.PdfRectangle)
	if _dcfdg._abba > _fddd._fbgge {
		_fddd._fbgge = _dcfdg._abba
	}
	_fddd._ecgcg = _agfb.Ury - _fddd.PdfRectangle.Lly
}
func _eceb(_bgcca _bf.PdfRectangle) rulingKind {
	_aefb := _bgcca.Width()
	_bfaa := _bgcca.Height()
	if _aefb > _bfaa {
		if _aefb >= _ffdeb {
			return _ebabd
		}
	} else {
		if _bfaa >= _ffdeb {
			return _beec
		}
	}
	return _gdbb
}
func (_eegd *textObject) showText(_beg _bg.PdfObject, _bad []byte) error {
	return _eegd.renderText(_beg, _bad)
}
func (_bgbg *subpath) isQuadrilateral() bool {
	if len(_bgbg._fcdc) < 4 || len(_bgbg._fcdc) > 5 {
		return false
	}
	if len(_bgbg._fcdc) == 5 {
		_baaf := _bgbg._fcdc[0]
		_cbfg := _bgbg._fcdc[4]
		if _baaf.X != _cbfg.X || _baaf.Y != _cbfg.Y {
			return false
		}
	}
	return true
}
func (_gebc *textTable) putComposite(_cccd, _daaeb int, _cgae paraList, _egcfc _bf.PdfRectangle) {
	if len(_cgae) == 0 {
		_ag.Log.Error("\u0074\u0065xt\u0054\u0061\u0062l\u0065\u0029\u0020\u0070utC\u006fmp\u006f\u0073\u0069\u0074\u0065\u003a\u0020em\u0070\u0074\u0079\u0020\u0070\u0061\u0072a\u0073")
		return
	}
	_cfab := compositeCell{PdfRectangle: _egcfc, paraList: _cgae}
	if _beae {
		_gdf.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0070\u0075\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u003c\u002d\u0025\u0073\u000a", _cccd, _daaeb, _cfab.String())
	}
	_cfab.updateBBox()
	_gebc._bbfb[_fcbc(_cccd, _daaeb)] = _cfab
}
func (_cadf compositeCell) parasBBox() (paraList, _bf.PdfRectangle) {
	return _cadf.paraList, _cadf.PdfRectangle
}
func _dfdae(_egdc, _cedaf int) int {
	if _egdc < _cedaf {
		return _egdc
	}
	return _cedaf
}

type markKind int

func (_fbcdc paraList) applyTables(_baed []*textTable) paraList {
	var _cede paraList
	for _, _feecb := range _baed {
		_cede = append(_cede, _feecb.newTablePara())
	}
	for _, _ecfga := range _fbcdc {
		if _ecfga._bggb {
			continue
		}
		_cede = append(_cede, _ecfga)
	}
	return _cede
}
func (_abg *wordBag) pullWord(_dgebb *textWord, _agad int, _gagbb map[int]map[*textWord]struct{}) {
	_abg.PdfRectangle = _effa(_abg.PdfRectangle, _dgebb.PdfRectangle)
	if _dgebb._fbgge > _abg._gaed {
		_abg._gaed = _dgebb._fbgge
	}
	_abg._bbf[_agad] = append(_abg._bbf[_agad], _dgebb)
	_gagbb[_agad][_dgebb] = struct{}{}
}
func _ea(_gddf []Font, _aa string) bool {
	for _, _dafe := range _gddf {
		if _dafe.FontName == _aa {
			return true
		}
	}
	return false
}
func _dagb(_acegd, _geff bounded) float64 { return _acbd(_acegd) - _acbd(_geff) }
func _ggdb(_acbg *wordBag, _fbfb *textWord, _ccfa float64) bool {
	return _acbg.Urx <= _fbfb.Llx && _fbfb.Llx < _acbg.Urx+_ccfa
}

// TableCell is a cell in a TextTable.
type TableCell struct {

	// Text is the extracted text.
	Text string

	// Marks returns the TextMarks corresponding to the text in Text.
	Marks TextMarkArray
}

func _aadf(_afae []*textMark, _fege _bf.PdfRectangle, _ecfb rulingList, _gada []gridTiling) paraList {
	_ag.Log.Trace("\u006d\u0061\u006b\u0065\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u003a \u0025\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073\u0020\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065=\u0025\u002e\u0032\u0066", len(_afae), _fege)
	if len(_afae) == 0 {
		return nil
	}
	_fccf := _cdbag(_afae, _fege)
	if len(_fccf) == 0 {
		return nil
	}
	_ecfb.log("\u006d\u0061\u006be\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065")
	_badc, _cgeg := _ecfb.vertsHorzs()
	_debdb := _cdg(_fccf, _fege.Ury, _badc, _cgeg)
	_afgbb := _bcdg(_debdb, _fege.Ury, _badc, _cgeg)
	_afgbb = _eceed(_afgbb)
	_bdcf := make(paraList, 0, len(_afgbb))
	for _, _gcae := range _afgbb {
		_gcce := _gcae.arrangeText()
		if _gcce != nil {
			_bdcf = append(_bdcf, _gcce)
		}
	}
	if len(_bdcf) >= _gdfd {
		_bdcf = _bdcf.extractTables(_gada)
	}
	_bdcf.sortReadingOrder()
	_bdcf.log("\u0073\u006f\u0072te\u0064\u0020\u0069\u006e\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0064\u0065\u0072")
	return _bdcf
}
func _beda(_adfbc map[int][]float64) {
	if len(_adfbc) <= 1 {
		return
	}
	_acde := _bacfa(_adfbc)
	if _beae {
		_ag.Log.Info("\u0066i\u0078C\u0065\u006c\u006c\u0073\u003a \u006b\u0065y\u0073\u003d\u0025\u002b\u0076", _acde)
	}
	var _aadec, _dgacg int
	for _aadec, _dgacg = range _acde {
		if _adfbc[_dgacg] != nil {
			break
		}
	}
	for _ebea, _afee := range _acde[_aadec:] {
		_gfee := _adfbc[_afee]
		if _gfee == nil {
			continue
		}
		if _beae {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u006b\u0030\u003d\u0025\u0064\u0020\u006b1\u003d\u0025\u0064\u000a", _aadec+_ebea, _dgacg, _afee)
		}
		_ecfbg := _adfbc[_afee]
		if _ecfbg[len(_ecfbg)-1] > _gfee[0] {
			_ecfbg[len(_ecfbg)-1] = _gfee[0]
			_adfbc[_dgacg] = _ecfbg
		}
		_dgacg = _afee
	}
}

// String returns a description of `l`.
func (_eegf *textLine) String() string {
	return _gdf.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _eegf._dcfd, _eegf.PdfRectangle, _eegf._ddef, _eegf.text())
}
func (_ecfg paraList) extractTables(_aafcb []gridTiling) paraList {
	if _beae {
		_ag.Log.Debug("\u0065\u0078\u0074r\u0061\u0063\u0074\u0054\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_ecfg))
	}
	if len(_ecfg) < _gdfd {
		return _ecfg
	}
	_aeed := _ecfg.findTables(_aafcb)
	if _beae {
		_ag.Log.Info("c\u006f\u006d\u0062\u0069\u006e\u0065d\u0020\u0074\u0061\u0062\u006c\u0065s\u0020\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_aeed))
		for _dbcfg, _fadcb := range _aeed {
			_fadcb.log(_gdf.Sprintf("c\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u0020\u0025\u0064", _dbcfg))
		}
	}
	return _ecfg.applyTables(_aeed)
}
func (_bcaf rulingList) sortStrict() {
	_f.Slice(_bcaf, func(_efcee, _agcbcg int) bool {
		_cbfed, _cgeba := _bcaf[_efcee], _bcaf[_agcbcg]
		_gcdd, _gegga := _cbfed._cgef, _cgeba._cgef
		if _gcdd != _gegga {
			return _gcdd > _gegga
		}
		_aeae, _dacc := _cbfed._eead, _cgeba._eead
		if !_fbga(_aeae - _dacc) {
			return _aeae < _dacc
		}
		_aeae, _dacc = _cbfed._bgeb, _cgeba._bgeb
		if _aeae != _dacc {
			return _aeae < _dacc
		}
		return _cbfed._eecc < _cgeba._eecc
	})
}
func _cdbag(_efbg []*textMark, _bdcdg _bf.PdfRectangle) []*textWord {
	var _cdefb []*textWord
	var _eeedg *textWord
	if _eccd {
		_ag.Log.Info("\u006d\u0061\u006beT\u0065\u0078\u0074\u0057\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0064\u0020\u006d\u0061\u0072\u006b\u0073", len(_efbg))
	}
	_bacca := func() {
		if _eeedg != nil {
			_befcd := _eeedg.computeText()
			if !_cgcdb(_befcd) {
				_eeedg._bcaa = _befcd
				_cdefb = append(_cdefb, _eeedg)
				if _eccd {
					_ag.Log.Info("\u0061\u0064\u0064Ne\u0077\u0057\u006f\u0072\u0064\u003a\u0020\u0025\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", len(_cdefb)-1, _eeedg.String())
					for _effede, _adae := range _eeedg._gdgg {
						_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _effede, _adae.String())
					}
				}
			}
			_eeedg = nil
		}
	}
	for _, _bgff := range _efbg {
		if _bfe && _eeedg != nil && len(_eeedg._gdgg) > 0 {
			_gbedd := _eeedg._gdgg[len(_eeedg._gdgg)-1]
			_cfebf, _acbbd := _gbae(_bgff._eeaf)
			_gggf, _abbc := _gbae(_gbedd._eeaf)
			if _acbbd && !_abbc && _gbedd.inDiacriticArea(_bgff) {
				_eeedg.addDiacritic(_cfebf)
				continue
			}
			if _abbc && !_acbbd && _bgff.inDiacriticArea(_gbedd) {
				_eeedg._gdgg = _eeedg._gdgg[:len(_eeedg._gdgg)-1]
				_eeedg.appendMark(_bgff, _bdcdg)
				_eeedg.addDiacritic(_gggf)
				continue
			}
		}
		_bgdcb := _cgcdb(_bgff._eeaf)
		if _bgdcb {
			_bacca()
			continue
		}
		if _eeedg == nil && !_bgdcb {
			_eeedg = _cedeb([]*textMark{_bgff}, _bdcdg)
			continue
		}
		_gabda := _eeedg._fbgge
		_fecec := _b.Abs(_aaegg(_bdcdg, _bgff)-_eeedg._ecgcg) / _gabda
		_begdf := _dgecg(_bgff, _eeedg) / _gabda
		if _begdf >= _gfdad || !(-_edge <= _begdf && _fecec <= _cbgg) {
			_bacca()
			_eeedg = _cedeb([]*textMark{_bgff}, _bdcdg)
			continue
		}
		_eeedg.appendMark(_bgff, _bdcdg)
	}
	_bacca()
	return _cdefb
}
func (_dcecg *textWord) addDiacritic(_eafda string) {
	_bbeg := _dcecg._gdgg[len(_dcecg._gdgg)-1]
	_bbeg._eeaf += _eafda
	_bbeg._eeaf = _bd.NFKC.String(_bbeg._eeaf)
}
func _cffd(_gbagg map[float64]map[float64]gridTile) []float64 {
	_cbec := make([]float64, 0, len(_gbagg))
	_defe := make(map[float64]struct{}, len(_gbagg))
	for _, _cfdca := range _gbagg {
		for _gebb := range _cfdca {
			if _, _ffcgf := _defe[_gebb]; _ffcgf {
				continue
			}
			_cbec = append(_cbec, _gebb)
			_defe[_gebb] = struct{}{}
		}
	}
	_f.Float64s(_cbec)
	return _cbec
}
func (_ddad *textObject) getFillColor() _eda.Color {
	return _gfdbb(_ddad._fcdg.ColorspaceNonStroking, _ddad._fcdg.ColorNonStroking)
}
func (_affa *wordBag) applyRemovals(_abce map[int]map[*textWord]struct{}) {
	for _gfcb, _cfac := range _abce {
		if len(_cfac) == 0 {
			continue
		}
		_cdf := _affa._bbf[_gfcb]
		_dafbe := len(_cdf) - len(_cfac)
		if _dafbe == 0 {
			delete(_affa._bbf, _gfcb)
			continue
		}
		_fde := make([]*textWord, _dafbe)
		_gea := 0
		for _, _bgfd := range _cdf {
			if _, _geea := _cfac[_bgfd]; !_geea {
				_fde[_gea] = _bgfd
				_gea++
			}
		}
		_affa._bbf[_gfcb] = _fde
	}
}

var _ebbb = _c.MustCompile("\u005e\u005c\u0073\u002a\u0028\u005c\u0064\u002b\u005c\u002e\u003f|\u005b\u0049\u0069\u0076\u005d\u002b\u0029\u005c\u0073\u002a\\\u0029\u003f\u0024")

func (_bgeg rulingList) vertsHorzs() (rulingList, rulingList) {
	var _agbf, _fdga rulingList
	for _, _daab := range _bgeg {
		switch _daab._cgef {
		case _beec:
			_agbf = append(_agbf, _daab)
		case _ebabd:
			_fdga = append(_fdga, _daab)
		}
	}
	return _agbf, _fdga
}
func (_ddfc *shapesState) lineTo(_bcba, _fbcfd float64) {
	if _bccf {
		_ag.Log.Info("\u006c\u0069\u006eeT\u006f\u0028\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0070\u003d\u0025\u002e\u0032\u0066", _bcba, _fbcfd, _ddfc.devicePoint(_bcba, _fbcfd))
	}
	_ddfc.addPoint(_bcba, _fbcfd)
}
func (_adgg *wordBag) firstWord(_ccde int) *textWord { return _adgg._bbf[_ccde][0] }
func (_efggd *ruling) intersects(_cedcd *ruling) bool {
	_fgec := (_efggd._cgef == _beec && _cedcd._cgef == _ebabd) || (_cedcd._cgef == _beec && _efggd._cgef == _ebabd)
	_cdfd := func(_efac, _fcfa *ruling) bool {
		return _efac._bgeb-_deff <= _fcfa._eead && _fcfa._eead <= _efac._eecc+_deff
	}
	_aadea := _cdfd(_efggd, _cedcd)
	_fgfb := _cdfd(_cedcd, _efggd)
	if _efa {
		_gdf.Printf("\u0020\u0020\u0020\u0020\u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003a\u0020\u0020\u006fr\u0074\u0068\u006f\u0067\u006f\u006e\u0061l\u003d\u0025\u0074\u0020\u006f\u0031\u003d\u0025\u0074\u0020\u006f2\u003d\u0025\u0074\u0020\u2192\u0020\u0025\u0074\u000a"+"\u0020\u0020\u0020 \u0020\u0020\u0020\u0076\u003d\u0025\u0073\u000a"+" \u0020\u0020\u0020\u0020\u0020\u0077\u003d\u0025\u0073\u000a", _fgec, _aadea, _fgfb, _fgec && _aadea && _fgfb, _efggd, _cedcd)
	}
	return _fgec && _aadea && _fgfb
}

const (
	_gdbb rulingKind = iota
	_ebabd
	_beec
)

func (_gdbab gridTile) numBorders() int {
	_fafb := 0
	if _gdbab._bceg {
		_fafb++
	}
	if _gdbab._abdfe {
		_fafb++
	}
	if _gdbab._abag {
		_fafb++
	}
	if _gdbab._gfga {
		_fafb++
	}
	return _fafb
}
func _aadc(_dcea string, _gdgaa []rulingList) {
	_ag.Log.Info("\u0024\u0024 \u0025\u0064\u0020g\u0072\u0069\u0064\u0073\u0020\u002d\u0020\u0025\u0073", len(_gdgaa), _dcea)
	for _ddee, _fgea := range _gdgaa {
		_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _ddee, _fgea.String())
	}
}
func (_gead rulingList) comp(_cbfa, _dgeeg int) bool {
	_gfec, _gcec := _gead[_cbfa], _gead[_dgeeg]
	_cffga, _baab := _gfec._cgef, _gcec._cgef
	if _cffga != _baab {
		return _cffga > _baab
	}
	if _cffga == _gdbb {
		return false
	}
	_cedcc := func(_baded bool) bool {
		if _cffga == _ebabd {
			return _baded
		}
		return !_baded
	}
	_beab, _afdgd := _gfec._eead, _gcec._eead
	if _beab != _afdgd {
		return _cedcc(_beab > _afdgd)
	}
	_beab, _afdgd = _gfec._bgeb, _gcec._bgeb
	if _beab != _afdgd {
		return _cedcc(_beab < _afdgd)
	}
	return _cedcc(_gfec._eecc < _gcec._eecc)
}
func (_dbgg *wordBag) depthIndexes() []int {
	if len(_dbgg._bbf) == 0 {
		return nil
	}
	_gggcd := make([]int, len(_dbgg._bbf))
	_edbe := 0
	for _gdab := range _dbgg._bbf {
		_gggcd[_edbe] = _gdab
		_edbe++
	}
	_f.Ints(_gggcd)
	return _gggcd
}
func (_dggfg rulingList) merge() *ruling {
	_cegb := _dggfg[0]._eead
	_eeea := _dggfg[0]._bgeb
	_ebced := _dggfg[0]._eecc
	for _, _fegdf := range _dggfg[1:] {
		_cegb += _fegdf._eead
		if _fegdf._bgeb < _eeea {
			_eeea = _fegdf._bgeb
		}
		if _fegdf._eecc > _ebced {
			_ebced = _fegdf._eecc
		}
	}
	_efgf := &ruling{_cgef: _dggfg[0]._cgef, _gbgb: _dggfg[0]._gbgb, Color: _dggfg[0].Color, _eead: _cegb / float64(len(_dggfg)), _bgeb: _eeea, _eecc: _ebced}
	if _cecg {
		_ag.Log.Info("\u006de\u0072g\u0065\u003a\u0020\u0025\u0032d\u0020\u0076e\u0063\u0073\u0020\u0025\u0073", len(_dggfg), _efgf)
		for _cdfca, _cedae := range _dggfg {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cdfca, _cedae)
		}
	}
	return _efgf
}

type ruling struct {
	_cgef rulingKind
	_gbgb markKind
	_eda.Color
	_eead float64
	_bgeb float64
	_eecc float64
	_efgd float64
}
type wordBag struct {
	_bf.PdfRectangle
	_gaed       float64
	_fff, _aede rulingList
	_fcgc       float64
	_bbf        map[int][]*textWord
}

func _abbf(_cffca []TextMark, _feea *int, _beb string) []TextMark {
	_debd := _ebf
	_debd.Text = _beb
	return _affac(_cffca, _feea, _debd)
}
func (_ffgac *textTable) depth() float64 {
	_eaadc := 1e10
	for _ccfe := 0; _ccfe < _ffgac._baee; _ccfe++ {
		_cgced := _ffgac.get(_ccfe, 0)
		if _cgced == nil || _cgced._dgggff {
			continue
		}
		_eaadc = _b.Min(_eaadc, _cgced.depth())
	}
	return _eaadc
}
func (_bbed rulingList) primMinMax() (float64, float64) {
	_gdddf, _acbb := _bbed[0]._eead, _bbed[0]._eead
	for _, _bccec := range _bbed[1:] {
		if _bccec._eead < _gdddf {
			_gdddf = _bccec._eead
		} else if _bccec._eead > _acbb {
			_acbb = _bccec._eead
		}
	}
	return _gdddf, _acbb
}
func (_eff *subpath) close() {
	if !_egag(_eff._fcdc[0], _eff.last()) {
		_eff.add(_eff._fcdc[0])
	}
	_eff._fcfe = true
	_eff.removeDuplicates()
}

const (
	RenderModeStroke RenderMode = 1 << iota
	RenderModeFill
	RenderModeClip
)

func _bbgcb(_acbf int, _bgae map[int][]float64) ([]int, int) {
	_agge := make([]int, _acbf)
	_ecbeg := 0
	for _abgfd := 0; _abgfd < _acbf; _abgfd++ {
		_agge[_abgfd] = _ecbeg
		_ecbeg += len(_bgae[_abgfd]) + 1
	}
	return _agge, _ecbeg
}
func _aaegg(_decb _bf.PdfRectangle, _ceaf bounded) float64 { return _decb.Ury - _ceaf.bbox().Lly }
func (_dfda *subpath) last() _ee.Point                     { return _dfda._fcdc[len(_dfda._fcdc)-1] }

var _gecc = map[markKind]string{_caade: "\u0073\u0074\u0072\u006f\u006b\u0065", _gaea: "\u0066\u0069\u006c\u006c", _eddbd: "\u0061u\u0067\u006d\u0065\u006e\u0074"}

func (_ec *PageFonts) extractPageResourcesToFont(_de *_bf.PdfPageResources) error {
	_cec, _gb := _bg.GetDict(_de.Font)
	if !_gb {
		return _g.New(_dg)
	}
	for _, _dafb := range _cec.Keys() {
		var (
			_gdd = true
			_eg  []byte
			_add string
		)
		_ef, _gac := _de.GetFontByName(_dafb)
		if !_gac {
			return _g.New(_edf)
		}
		_cf, _bgd := _bf.NewPdfFontFromPdfObject(_ef)
		if _bgd != nil {
			return _bgd
		}
		_ff := _cf.FontDescriptor()
		_df := _cf.FontDescriptor().FontName.String()
		_bbg := _cf.Subtype()
		if _ea(_ec.Fonts, _df) {
			continue
		}
		if len(_cf.ToUnicode()) == 0 {
			_gdd = false
		}
		if _ff.FontFile != nil {
			if _bge, _db := _bg.GetStream(_ff.FontFile); _db {
				_eg, _bgd = _bg.DecodeStream(_bge)
				if _bgd != nil {
					return _bgd
				}
				_add = _df + "\u002e\u0070\u0066\u0062"
			}
		} else if _ff.FontFile2 != nil {
			if _ggd, _gce := _bg.GetStream(_ff.FontFile2); _gce {
				_eg, _bgd = _bg.DecodeStream(_ggd)
				if _bgd != nil {
					return _bgd
				}
				_add = _df + "\u002e\u0074\u0074\u0066"
			}
		} else if _ff.FontFile3 != nil {
			if _fg, _cgb := _bg.GetStream(_ff.FontFile3); _cgb {
				_eg, _bgd = _bg.DecodeStream(_fg)
				if _bgd != nil {
					return _bgd
				}
				_add = _df + "\u002e\u0063\u0066\u0066"
			}
		}
		if len(_add) < 1 {
			_ag.Log.Debug(_cgc)
		}
		_dc := Font{FontName: _df, PdfFont: _cf, IsCID: _cf.IsCID(), IsSimple: _cf.IsSimple(), ToUnicode: _gdd, FontType: _bbg, FontData: _eg, FontFileName: _add, FontDescriptor: _ff}
		_ec.Fonts = append(_ec.Fonts, _dc)
	}
	return nil
}
func (_bdfad *textLine) toTextMarks(_bfdd *int) []TextMark {
	var _dcd []TextMark
	for _, _ggbag := range _bdfad._egad {
		if _ggbag._egcbd {
			_dcd = _abbf(_dcd, _bfdd, "\u0020")
		}
		_fgdb := _ggbag.toTextMarks(_bfdd)
		_dcd = append(_dcd, _fgdb...)
	}
	return _dcd
}

type textTable struct {
	_bf.PdfRectangle
	_baee, _cabfg int
	_acgbd        bool
	_aagc         map[uint64]*textPara
	_bbfb         map[uint64]compositeCell
}

func _bfcbe(_fcca, _faec, _afadd float64) rulingKind {
	if _fcca >= _afadd && _dgcd(_faec, _fcca) {
		return _ebabd
	}
	if _faec >= _afadd && _dgcd(_fcca, _faec) {
		return _beec
	}
	return _gdbb
}
func (_gedf *wordBag) arrangeText() *textPara {
	_gedf.sort()
	if _bccb {
		_gedf.removeDuplicates()
	}
	var _gdgb []*textLine
	for _, _gccea := range _gedf.depthIndexes() {
		for !_gedf.empty(_gccea) {
			_bfea := _gedf.firstReadingIndex(_gccea)
			_aafg := _gedf.firstWord(_bfea)
			_ecaa := _cdgf(_gedf, _bfea)
			_ecaff := _aafg._fbgge
			_bacd := _aafg._ecgcg - _bgcc*_ecaff
			_addef := _aafg._ecgcg + _bgcc*_ecaff
			_fbee := _dcca * _ecaff
			_cbgec := _ceda * _ecaff
		_aba:
			for {
				var _geeba *textWord
				_daba := 0
				for _, _fbcdb := range _gedf.depthBand(_bacd, _addef) {
					_ecdf := _gedf.highestWord(_fbcdb, _bacd, _addef)
					if _ecdf == nil {
						continue
					}
					_bgaa := _dgecg(_ecdf, _ecaa._egad[len(_ecaa._egad)-1])
					if _bgaa < -_cbgec {
						break _aba
					}
					if _bgaa > _fbee {
						continue
					}
					if _geeba != nil && _addb(_ecdf, _geeba) >= 0 {
						continue
					}
					_geeba = _ecdf
					_daba = _fbcdb
				}
				if _geeba == nil {
					break
				}
				_ecaa.pullWord(_gedf, _geeba, _daba)
			}
			_ecaa.markWordBoundaries()
			_gdgb = append(_gdgb, _ecaa)
		}
	}
	if len(_gdgb) == 0 {
		return nil
	}
	_f.Slice(_gdgb, func(_bbb, _bfdf int) bool { return _acgf(_gdgb[_bbb], _gdgb[_bfdf]) < 0 })
	_gaaa := _egecc(_gedf.PdfRectangle, _gdgb)
	if _cfed {
		_ag.Log.Info("\u0061\u0072\u0072an\u0067\u0065\u0054\u0065\u0078\u0074\u0020\u0021\u0021\u0021\u0020\u0070\u0061\u0072\u0061\u003d\u0025\u0073", _gaaa.String())
		if _bcf {
			for _dfeg, _eceeg := range _gaaa._ecbdg {
				_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dfeg, _eceeg.String())
				if _aebf {
					for _bbggd, _deeb := range _eceeg._egad {
						_gdf.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _bbggd, _deeb.String())
						for _fcdcd, _baca := range _deeb._gdgg {
							_gdf.Printf("\u00251\u0032\u0064\u003a\u0020\u0025\u0073\n", _fcdcd, _baca.String())
						}
					}
				}
			}
		}
	}
	return _gaaa
}

// Marks returns the TextMark collection for a page. It represents all the text on the page.
func (_caef PageText) Marks() *TextMarkArray { return &TextMarkArray{_fceg: _caef._cagb} }

type subpath struct {
	_fcdc []_ee.Point
	_fcfe bool
}

// Font represents the font properties on a PDF page.
type Font struct {
	PdfFont *_bf.PdfFont

	// FontName represents Font Name from font properties.
	FontName string

	// FontType represents Font Subtype entry in the font dictionary inside page resources.
	// Examples : type0, Type1, MMType1, Type3, TrueType, CIDFont.
	FontType string

	// ToUnicode is true if font provides a `ToUnicode` mapping.
	ToUnicode bool

	// IsCID is true if underlying font is a composite font.
	// Composite font is represented by a font dictionary whose Subtype is `Type0`
	IsCID bool

	// IsSimple is true if font is simple font.
	// A simple font is limited to only 8 bit (255) character codes.
	IsSimple bool

	// FontData represents the raw data of the embedded font file.
	// It can have format TrueType (TTF), PostScript Font (PFB) or Compact Font Format (CCF).
	// FontData value can be indicates from `FontFile`, `FontFile2` or `FontFile3` inside Font Descriptor.
	// At most, only one of `FontFile`, `FontFile2` or `FontFile3` will be FontData value.
	FontData []byte

	// FontFileName is a name representing the font. it has format:
	// (Font Name) + (Font Type Extension), example: helvetica.ttf.
	FontFileName string

	// FontDescriptor represents metrics and other attributes inside font properties from PDF Structure (Font Descriptor).
	FontDescriptor *_bf.PdfFontDescriptor
}

func _gfaeb(_gcda _bf.PdfRectangle) *ruling {
	return &ruling{_cgef: _ebabd, _eead: _gcda.Lly, _bgeb: _gcda.Llx, _eecc: _gcda.Urx}
}

type lineRuling struct {
	_aedc rulingKind
	_daff markKind
	_eda.Color
	_abec, _eacb _ee.Point
}

func (_ggff *textLine) appendWord(_eaff *textWord) {
	_ggff._egad = append(_ggff._egad, _eaff)
	_ggff.PdfRectangle = _effa(_ggff.PdfRectangle, _eaff.PdfRectangle)
	if _eaff._fbgge > _ggff._ddef {
		_ggff._ddef = _eaff._fbgge
	}
	if _eaff._ecgcg > _ggff._dcfd {
		_ggff._dcfd = _eaff._ecgcg
	}
}

// TextMarkArray is a collection of TextMarks.
type TextMarkArray struct{ _fceg []TextMark }

func (_cabf paraList) llyRange(_aceba []int, _egbe, _gddd float64) []int {
	_eefc := len(_cabf)
	if _gddd < _cabf[_aceba[0]].Lly || _egbe > _cabf[_aceba[_eefc-1]].Lly {
		return nil
	}
	_bdbea := _f.Search(_eefc, func(_bgdd int) bool { return _cabf[_aceba[_bgdd]].Lly >= _egbe })
	_gece := _f.Search(_eefc, func(_ddadf int) bool { return _cabf[_aceba[_ddadf]].Lly > _gddd })
	return _aceba[_bdbea:_gece]
}
func (_ggea *wordBag) getDepthIdx(_adbc float64) int {
	_beaa := _ggea.depthIndexes()
	_gcbg := _feef(_adbc)
	if _gcbg < _beaa[0] {
		return _beaa[0]
	}
	if _gcbg > _beaa[len(_beaa)-1] {
		return _beaa[len(_beaa)-1]
	}
	return _gcbg
}
func (_acgaf *wordBag) scanBand(_eafc string, _eafa *wordBag, _bgdc func(_bbff *wordBag, _efb *textWord) bool, _dbbb, _ceed, _dabe float64, _bdbf, _eac bool) int {
	_bcga := _eafa._gaed
	var _fcgf map[int]map[*textWord]struct{}
	if !_bdbf {
		_fcgf = _acgaf.makeRemovals()
	}
	_cdef := _bgcc * _bcga
	_afde := 0
	for _, _cdae := range _acgaf.depthBand(_dbbb-_cdef, _ceed+_cdef) {
		if len(_acgaf._bbf[_cdae]) == 0 {
			continue
		}
		for _, _cffg := range _acgaf._bbf[_cdae] {
			if !(_dbbb-_cdef <= _cffg._ecgcg && _cffg._ecgcg <= _ceed+_cdef) {
				continue
			}
			if !_bgdc(_eafa, _cffg) {
				continue
			}
			_aeac := 2.0 * _b.Abs(_cffg._fbgge-_eafa._gaed) / (_cffg._fbgge + _eafa._gaed)
			_fegd := _b.Max(_cffg._fbgge/_eafa._gaed, _eafa._gaed/_cffg._fbgge)
			_cgbgf := _b.Min(_aeac, _fegd)
			if _dabe > 0 && _cgbgf > _dabe {
				continue
			}
			if _eafa.blocked(_cffg) {
				continue
			}
			if !_bdbf {
				_eafa.pullWord(_cffg, _cdae, _fcgf)
			}
			_afde++
			if !_eac {
				if _cffg._ecgcg < _dbbb {
					_dbbb = _cffg._ecgcg
				}
				if _cffg._ecgcg > _ceed {
					_ceed = _cffg._ecgcg
				}
			}
			if _bdbf {
				break
			}
		}
	}
	if !_bdbf {
		_acgaf.applyRemovals(_fcgf)
	}
	return _afde
}
func _feef(_ddcg float64) int {
	var _cdff int
	if _ddcg >= 0 {
		_cdff = int(_ddcg / _ffde)
	} else {
		_cdff = int(_ddcg/_ffde) - 1
	}
	return _cdff
}
func (_accad paraList) findTables(_adbeb []gridTiling) []*textTable {
	_accad.addNeighbours()
	_f.Slice(_accad, func(_bgbc, _cacfa int) bool { return _afa(_accad[_bgbc], _accad[_cacfa]) < 0 })
	var _cgage []*textTable
	if _bagf {
		_gfaa := _accad.findGridTables(_adbeb)
		_cgage = append(_cgage, _gfaa...)
	}
	if _bfcda {
		_cgdb := _accad.findTextTables()
		_cgage = append(_cgage, _cgdb...)
	}
	return _cgage
}
func (_eaeb *wordBag) blocked(_eded *textWord) bool {
	if _eded.Urx < _eaeb.Llx {
		_acfb := _eagae(_eded.PdfRectangle)
		_cfce := _debdg(_eaeb.PdfRectangle)
		if _eaeb._fff.blocks(_acfb, _cfce) {
			if _fgad {
				_ag.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0078\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _eded, _eaeb)
			}
			return true
		}
	} else if _eaeb.Urx < _eded.Llx {
		_aceg := _eagae(_eaeb.PdfRectangle)
		_ecab := _debdg(_eded.PdfRectangle)
		if _eaeb._fff.blocks(_aceg, _ecab) {
			if _fgad {
				_ag.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0078\u2192\u0020:\u0020\u0025\u0073\u0020%s", _eded, _eaeb)
			}
			return true
		}
	}
	if _eded.Ury < _eaeb.Lly {
		_fgfg := _eeee(_eded.PdfRectangle)
		_cecd := _gfaeb(_eaeb.PdfRectangle)
		if _eaeb._aede.blocks(_fgfg, _cecd) {
			if _fgad {
				_ag.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0079\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _eded, _eaeb)
			}
			return true
		}
	} else if _eaeb.Ury < _eded.Lly {
		_gcfd := _eeee(_eaeb.PdfRectangle)
		_dcff := _gfaeb(_eded.PdfRectangle)
		if _eaeb._aede.blocks(_gcfd, _dcff) {
			if _fgad {
				_ag.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0079\u2192\u0020:\u0020\u0025\u0073\u0020%s", _eded, _eaeb)
			}
			return true
		}
	}
	return false
}
func (_fffc *wordBag) firstReadingIndex(_fadc int) int {
	_bcbdc := _fffc.firstWord(_fadc)._fbgge
	_cgcef := float64(_fadc+1) * _ffde
	_ceee := _cgcef + _cdge*_bcbdc
	_decf := _fadc
	for _, _gdfc := range _fffc.depthBand(_cgcef, _ceee) {
		if _addb(_fffc.firstWord(_gdfc), _fffc.firstWord(_decf)) < 0 {
			_decf = _gdfc
		}
	}
	return _decf
}
func (_bbdd *textWord) computeText() string {
	_bddbb := make([]string, len(_bbdd._gdgg))
	for _feda, _fedca := range _bbdd._gdgg {
		_bddbb[_feda] = _fedca._eeaf
	}
	return _d.Join(_bddbb, "")
}
func (_dgg *textObject) nextLine() { _dgg.moveLP(0, -_dgg._degf._acfa) }
func (_eaed rulingList) secMinMax() (float64, float64) {
	_aaad, _dafbd := _eaed[0]._bgeb, _eaed[0]._eecc
	for _, _adbcb := range _eaed[1:] {
		if _adbcb._bgeb < _aaad {
			_aaad = _adbcb._bgeb
		}
		if _adbcb._eecc > _dafbd {
			_dafbd = _adbcb._eecc
		}
	}
	return _aaad, _dafbd
}
func (_fffb *textTable) newTablePara() *textPara {
	_ecce := _fffb.computeBbox()
	_cfbf := &textPara{PdfRectangle: _ecce, _dcfc: _ecce, _bccg: _fffb}
	if _beae {
		_ag.Log.Info("\u006e\u0065w\u0054\u0061\u0062l\u0065\u0050\u0061\u0072\u0061\u003a\u0020\u0025\u0073", _cfbf)
	}
	return _cfbf
}
func _ccede(_efbdf float64) float64 { return _cegd * _b.Round(_efbdf/_cegd) }

// String returns a description of `b`.
func (_fgaf *wordBag) String() string {
	var _caefb []string
	for _, _bdbg := range _fgaf.depthIndexes() {
		_cgfag := _fgaf._bbf[_bdbg]
		for _, _cff := range _cgfag {
			_caefb = append(_caefb, _cff._bcaa)
		}
	}
	return _gdf.Sprintf("\u0025.\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065=\u0025\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0025\u0071", _fgaf.PdfRectangle, _fgaf._gaed, len(_caefb), _caefb)
}

// Elements returns the TextMarks in `ma`.
func (_fae *TextMarkArray) Elements() []TextMark { return _fae._fceg }
func (_dbd *stateStack) top() *textState {
	if _dbd.empty() {
		return nil
	}
	return (*_dbd)[_dbd.size()-1]
}
func _effa(_fegb, _gcfc _bf.PdfRectangle) _bf.PdfRectangle {
	return _bf.PdfRectangle{Llx: _b.Min(_fegb.Llx, _gcfc.Llx), Lly: _b.Min(_fegb.Lly, _gcfc.Lly), Urx: _b.Max(_fegb.Urx, _gcfc.Urx), Ury: _b.Max(_fegb.Ury, _gcfc.Ury)}
}
func _eceed(_fegf []*wordBag) []*wordBag {
	if len(_fegf) <= 1 {
		return _fegf
	}
	if _cfed {
		_ag.Log.Info("\u006d\u0065\u0072\u0067\u0065\u0057\u006f\u0072\u0064B\u0061\u0067\u0073\u003a")
	}
	_f.Slice(_fegf, func(_eeff, _debc int) bool {
		_abdf, _geggb := _fegf[_eeff], _fegf[_debc]
		_fbad := _abdf.Width() * _abdf.Height()
		_becc := _geggb.Width() * _geggb.Height()
		if _fbad != _becc {
			return _fbad > _becc
		}
		if _abdf.Height() != _geggb.Height() {
			return _abdf.Height() > _geggb.Height()
		}
		return _eeff < _debc
	})
	var _beag []*wordBag
	_gbfb := make(intSet)
	for _gaa := 0; _gaa < len(_fegf); _gaa++ {
		if _gbfb.has(_gaa) {
			continue
		}
		_cgad := _fegf[_gaa]
		for _dgec := _gaa + 1; _dgec < len(_fegf); _dgec++ {
			if _gbfb.has(_gaa) {
				continue
			}
			_gaae := _fegf[_dgec]
			_gggc := _cgad.PdfRectangle
			_gggc.Llx -= _cgad._gaed
			if _ceba(_gggc, _gaae.PdfRectangle) {
				_cgad.absorb(_gaae)
				_gbfb.add(_dgec)
			}
		}
		_beag = append(_beag, _cgad)
	}
	if len(_fegf) != len(_beag)+len(_gbfb) {
		_ag.Log.Error("\u006d\u0065\u0072ge\u0057\u006f\u0072\u0064\u0042\u0061\u0067\u0073\u003a \u0025d\u2192%\u0064 \u0061\u0062\u0073\u006f\u0072\u0062\u0065\u0064\u003d\u0025\u0064", len(_fegf), len(_beag), len(_gbfb))
	}
	return _beag
}

// String returns a human readable description of `ss`.
func (_bcce *shapesState) String() string {
	return _gdf.Sprintf("\u007b\u0025\u0064\u0020su\u0062\u0070\u0061\u0074\u0068\u0073\u0020\u0066\u0072\u0065\u0073\u0068\u003d\u0025t\u007d", len(_bcce._cfgd), _bcce._ecd)
}
func _badbf(_fagba map[float64]map[float64]gridTile) []float64 {
	_bbaf := make([]float64, 0, len(_fagba))
	for _aadb := range _fagba {
		_bbaf = append(_bbaf, _aadb)
	}
	_f.Float64s(_bbaf)
	_ceaa := len(_bbaf)
	for _gaeef := 0; _gaeef < _ceaa/2; _gaeef++ {
		_bbaf[_gaeef], _bbaf[_ceaa-1-_gaeef] = _bbaf[_ceaa-1-_gaeef], _bbaf[_gaeef]
	}
	return _bbaf
}
func _cde(_bcg _bf.PdfRectangle) textState {
	return textState{_gdef: 100, _egea: RenderModeFill, _dadc: _bcg}
}
func (_egbc gridTiling) log(_acca string) {
	if !_efbc {
		return
	}
	_ag.Log.Info("\u0074i\u006ci\u006e\u0067\u003a\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0071", len(_egbc._fdeac), len(_egbc._gccf), _acca)
	_gdf.Printf("\u0020\u0020\u0020l\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a", _egbc._fdeac)
	_gdf.Printf("\u0020\u0020\u0020l\u006c\u0079\u003d\u0025\u002e\u0032\u0066\u000a", _egbc._gccf)
	for _gdge, _dcdf := range _egbc._gccf {
		_efdf, _efcec := _egbc._ffdd[_dcdf]
		if !_efcec {
			continue
		}
		_gdf.Printf("%\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _gdge, _dcdf)
		for _fdag, _aage := range _egbc._fdeac {
			_ggbe, _cebfg := _efdf[_aage]
			if !_cebfg {
				continue
			}
			_gdf.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _fdag, _ggbe.String())
		}
	}
}
func (_dbcg *textObject) moveLP(_gacc, _gdce float64) {
	_dbcg._baff.Concat(_ee.NewMatrix(1, 0, 0, 1, _gacc, _gdce))
	_dbcg._gcb = _dbcg._baff
}
func (_ffagb *textTable) getDown() paraList {
	_ccff := make(paraList, _ffagb._baee)
	for _abda := 0; _abda < _ffagb._baee; _abda++ {
		_fafe := _ffagb.get(_abda, _ffagb._cabfg-1)._egec
		if _fafe.taken() {
			return nil
		}
		_ccff[_abda] = _fafe
	}
	for _bffg := 0; _bffg < _ffagb._baee-1; _bffg++ {
		if _ccff[_bffg]._bfdc != _ccff[_bffg+1] {
			return nil
		}
	}
	return _ccff
}
func (_bdddd *ruling) gridIntersecting(_cgfdc *ruling) bool {
	return _cgcd(_bdddd._bgeb, _cgfdc._bgeb) && _cgcd(_bdddd._eecc, _cgfdc._eecc)
}
func _cgcd(_dacg, _gecbf float64) bool { return _b.Abs(_dacg-_gecbf) <= _deff }

type textWord struct {
	_bf.PdfRectangle
	_ecgcg float64
	_bcaa  string
	_gdgg  []*textMark
	_fbgge float64
	_egcbd bool
}

// String returns a string describing `tm`.
func (_dcg TextMark) String() string {
	_befb := _dcg.BBox
	var _bgde string
	if _dcg.Font != nil {
		_bgde = _dcg.Font.String()
		if len(_bgde) > 50 {
			_bgde = _bgde[:50] + "\u002e\u002e\u002e"
		}
	}
	var _gccd string
	if _dcg.Meta {
		_gccd = "\u0020\u002a\u004d\u002a"
	}
	return _gdf.Sprintf("\u007b\u0054\u0065\u0078t\u004d\u0061\u0072\u006b\u003a\u0020\u0025\u0064\u0020%\u0071\u003d\u0025\u0030\u0032\u0078\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e2\u0066\u0029\u0020\u0028\u00256\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0025\u0073\u007d", _dcg.Offset, _dcg.Text, []rune(_dcg.Text), _befb.Llx, _befb.Lly, _befb.Urx, _befb.Ury, _bgde, _gccd)
}

// BBox returns the smallest axis-aligned rectangle that encloses all the TextMarks in `ma`.
func (_fad *TextMarkArray) BBox() (_bf.PdfRectangle, bool) {
	var _gfc _bf.PdfRectangle
	_fcga := false
	for _, _fbeg := range _fad._fceg {
		if _fbeg.Meta || _cgcdb(_fbeg.Text) {
			continue
		}
		if _fcga {
			_gfc = _effa(_gfc, _fbeg.BBox)
		} else {
			_gfc = _fbeg.BBox
			_fcga = true
		}
	}
	return _gfc, _fcga
}

// TextTable represents a table.
// Cells are ordered top-to-bottom, left-to-right.
// Cells[y] is the (0-offset) y'th row in the table.
// Cells[y][x] is the (0-offset) x'th column in the table.
type TextTable struct {
	W, H  int
	Cells [][]TableCell
}

func (_cdcb paraList) readBefore(_gdgd []int, _bgccb, _edfa int) bool {
	_bedf, _geac := _cdcb[_bgccb], _cdcb[_edfa]
	if _ffff(_bedf, _geac) && _bedf.Lly > _geac.Lly {
		return true
	}
	if !(_bedf._dcfc.Urx < _geac._dcfc.Llx) {
		return false
	}
	_affe, _ggfd := _bedf.Lly, _geac.Lly
	if _affe > _ggfd {
		_ggfd, _affe = _affe, _ggfd
	}
	_acfd := _b.Max(_bedf._dcfc.Llx, _geac._dcfc.Llx)
	_fdea := _b.Min(_bedf._dcfc.Urx, _geac._dcfc.Urx)
	_dbdcc := _cdcb.llyRange(_gdgd, _affe, _ggfd)
	for _, _abfc := range _dbdcc {
		if _abfc == _bgccb || _abfc == _edfa {
			continue
		}
		_cab := _cdcb[_abfc]
		if _cab._dcfc.Llx <= _fdea && _acfd <= _cab._dcfc.Urx {
			return false
		}
	}
	return true
}
func (_gef *textObject) showTextAdjusted(_fdb *_bg.PdfObjectArray) error {
	_dae := false
	for _, _bdbe := range _fdb.Elements() {
		switch _bdbe.(type) {
		case *_bg.PdfObjectFloat, *_bg.PdfObjectInteger:
			_fga, _abca := _bg.GetNumberAsFloat(_bdbe)
			if _abca != nil {
				_ag.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078t\u0041\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0042\u0061\u0064\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0061\u006c\u0020a\u0072\u0067\u002e\u0020\u006f\u003d\u0025\u0073\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _bdbe, _fdb)
				return _abca
			}
			_dbfc, _geg := -_fga*0.001*_gef._degf._adfb, 0.0
			if _dae {
				_geg, _dbfc = _dbfc, _geg
			}
			_cebd := _def(_ee.Point{X: _dbfc, Y: _geg})
			_gef._gcb.Concat(_cebd)
		case *_bg.PdfObjectString:
			_bdga := _bg.TraceToDirectObject(_bdbe)
			_ddaa, _cccb := _bg.GetStringBytes(_bdga)
			if !_cccb {
				_ag.Log.Trace("s\u0068\u006f\u0077\u0054\u0065\u0078\u0074\u0041\u0064j\u0075\u0073\u0074\u0065\u0064\u003a\u0020Ba\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u002e\u0020o=\u0025\u0073 \u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _bdbe, _fdb)
				return _bg.ErrTypeError
			}
			_gef.renderText(_bdga, _ddaa)
		default:
			_ag.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078\u0074A\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0055\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0028%T\u0029\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _bdbe, _fdb)
			return _bg.ErrTypeError
		}
	}
	return nil
}
func _acbd(_ccea bounded) float64 { return -_ccea.bbox().Lly }
func (_gega *textLine) pullWord(_cebg *wordBag, _edde *textWord, _gbgd int) {
	_gega.appendWord(_edde)
	_cebg.removeWord(_edde, _gbgd)
}
func (_bcbcg rulingList) toGrids() []rulingList {
	if _efa {
		_ag.Log.Info("t\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _bcbcg)
	}
	_dgfc := _bcbcg.intersections()
	if _efa {
		_ag.Log.Info("\u0074\u006f\u0047r\u0069\u0064\u0073\u003a \u0076\u0065\u0063\u0073\u003d\u0025\u0064 \u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003d\u0025\u0064\u0020", len(_bcbcg), len(_dgfc))
		for _, _accd := range _ccagd(_dgfc) {
			_gdf.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _accd, _dgfc[_accd])
		}
	}
	_fbgc := make(map[int]intSet, len(_bcbcg))
	for _abad := range _bcbcg {
		_dbef := _bcbcg.connections(_dgfc, _abad)
		if len(_dbef) > 0 {
			_fbgc[_abad] = _dbef
		}
	}
	if _efa {
		_ag.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0063\u006fn\u006e\u0065\u0063\u0074s=\u0025\u0064", len(_fbgc))
		for _, _gfca := range _ccagd(_fbgc) {
			_gdf.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _gfca, _fbgc[_gfca])
		}
	}
	_aggg := _ccdef(len(_bcbcg), func(_fgde, _cffbg int) bool {
		_adbcf, _ffbb := len(_fbgc[_fgde]), len(_fbgc[_cffbg])
		if _adbcf != _ffbb {
			return _adbcf > _ffbb
		}
		return _bcbcg.comp(_fgde, _cffbg)
	})
	if _efa {
		_ag.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u006f\u0072d\u0065\u0072\u0069\u006eg=\u0025\u0076", _aggg)
	}
	_dfed := [][]int{{_aggg[0]}}
_dabaf:
	for _, _cbgc := range _aggg[1:] {
		for _abeb, _afegc := range _dfed {
			for _, _aaccb := range _afegc {
				if _fbgc[_aaccb].has(_cbgc) {
					_dfed[_abeb] = append(_afegc, _cbgc)
					continue _dabaf
				}
			}
		}
		_dfed = append(_dfed, []int{_cbgc})
	}
	if _efa {
		_ag.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020i\u0067\u0072i\u0064\u0073\u003d\u0025\u0076", _dfed)
	}
	_f.SliceStable(_dfed, func(_bagfc, _ceec int) bool { return len(_dfed[_bagfc]) > len(_dfed[_ceec]) })
	for _, _abfcg := range _dfed {
		_f.Slice(_abfcg, func(_cedg, _ddfd int) bool { return _bcbcg.comp(_abfcg[_cedg], _abfcg[_ddfd]) })
	}
	_bcgaf := make([]rulingList, len(_dfed))
	for _gedgf, _bgfg := range _dfed {
		_fafg := make(rulingList, len(_bgfg))
		for _bgga, _ccba := range _bgfg {
			_fafg[_bgga] = _bcbcg[_ccba]
		}
		_bcgaf[_gedgf] = _fafg
	}
	if _efa {
		_ag.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020g\u0072\u0069d\u0073\u003d\u0025\u002b\u0076", _bcgaf)
	}
	var _dcec []rulingList
	for _, _bebfe := range _bcgaf {
		if _geacb, _fage := _bebfe.isActualGrid(); _fage {
			_bebfe = _geacb
			_bebfe = _bebfe.snapToGroups()
			_dcec = append(_dcec, _bebfe)
		}
	}
	if _efa {
		_aadc("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0061\u0063t\u0075\u0061\u006c\u0047ri\u0064\u0073", _dcec)
		_ag.Log.Info("\u0074\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0067\u0072\u0069\u0064\u0073\u003d%\u0064 \u0061\u0063\u0074\u0075\u0061\u006c\u0047\u0072\u0069\u0064\u0073\u003d\u0025\u0064", len(_bcgaf), len(_dcec))
	}
	return _dcec
}
func (_bdeec paraList) yNeighbours(_cbddd float64) map[*textPara][]int {
	_effed := make([]event, 2*len(_bdeec))
	if _cbddd == 0 {
		for _faecd, _gdbdf := range _bdeec {
			_effed[2*_faecd] = event{_gdbdf.Lly, true, _faecd}
			_effed[2*_faecd+1] = event{_gdbdf.Ury, false, _faecd}
		}
	} else {
		for _edeeg, _bcffb := range _bdeec {
			_effed[2*_edeeg] = event{_bcffb.Lly - _cbddd*_bcffb.fontsize(), true, _edeeg}
			_effed[2*_edeeg+1] = event{_bcffb.Ury + _cbddd*_bcffb.fontsize(), false, _edeeg}
		}
	}
	return _bdeec.eventNeighbours(_effed)
}
func (_edcg *textTable) markCells() {
	for _bbdcb := 0; _bbdcb < _edcg._cabfg; _bbdcb++ {
		for _bagc := 0; _bagc < _edcg._baee; _bagc++ {
			_ceaag := _edcg.get(_bagc, _bbdcb)
			if _ceaag != nil {
				_ceaag._bggb = true
			}
		}
	}
}

type textLine struct {
	_bf.PdfRectangle
	_dcfd float64
	_egad []*textWord
	_ddef float64
}

func _egag(_dcbd, _bdcda _ee.Point) bool { return _dcbd.X == _bdcda.X && _dcbd.Y == _bdcda.Y }
func (_aadcd *textTable) growTable() {
	_eaefg := func(_bfbce paraList) {
		_aadcd._cabfg++
		for _ccfg := 0; _ccfg < _aadcd._baee; _ccfg++ {
			_dfg := _bfbce[_ccfg]
			_aadcd.put(_ccfg, _aadcd._cabfg-1, _dfg)
		}
	}
	_badcc := func(_ggge paraList) {
		_aadcd._baee++
		for _bccdd := 0; _bccdd < _aadcd._cabfg; _bccdd++ {
			_debea := _ggge[_bccdd]
			_aadcd.put(_aadcd._baee-1, _bccdd, _debea)
		}
	}
	if _ddgf {
		_aadcd.log("\u0067r\u006f\u0077\u0054\u0061\u0062\u006ce")
	}
	for _bfeag := 0; ; _bfeag++ {
		_dddgg := false
		_bead := _aadcd.getDown()
		_aaege := _aadcd.getRight()
		if _ddgf {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bfeag, _aadcd)
			_gdf.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0020\u0064\u006f\u0077\u006e\u003d\u0025\u0073\u000a", _bead)
			_gdf.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0072\u0069\u0067\u0068\u0074\u003d\u0025\u0073\u000a", _aaege)
		}
		if _bead != nil && _aaege != nil {
			_febf := _bead[len(_bead)-1]
			if !_febf.taken() && _febf == _aaege[len(_aaege)-1] {
				_eaefg(_bead)
				if _aaege = _aadcd.getRight(); _aaege != nil {
					_badcc(_aaege)
					_aadcd.put(_aadcd._baee-1, _aadcd._cabfg-1, _febf)
				}
				_dddgg = true
			}
		}
		if !_dddgg && _bead != nil {
			_eaefg(_bead)
			_dddgg = true
		}
		if !_dddgg && _aaege != nil {
			_badcc(_aaege)
			_dddgg = true
		}
		if !_dddgg {
			break
		}
	}
}
func _bacfa(_ebba map[int][]float64) []int {
	_edfd := make([]int, len(_ebba))
	_dfaaf := 0
	for _bbbc := range _ebba {
		_edfd[_dfaaf] = _bbbc
		_dfaaf++
	}
	_f.Ints(_edfd)
	return _edfd
}
func (_baeb rulingList) aligned() bool {
	if len(_baeb) < 2 {
		return false
	}
	_dged := make(map[*ruling]int)
	_dged[_baeb[0]] = 0
	for _, _eccdb := range _baeb[1:] {
		_dceb := false
		for _fadg := range _dged {
			if _eccdb.gridIntersecting(_fadg) {
				_dged[_fadg]++
				_dceb = true
				break
			}
		}
		if !_dceb {
			_dged[_eccdb] = 0
		}
	}
	_dfaf := 0
	for _, _cbdc := range _dged {
		if _cbdc == 0 {
			_dfaf++
		}
	}
	_cbdb := float64(_dfaf) / float64(len(_baeb))
	_fadfb := _cbdb <= 1.0-_cagg
	if _efa {
		_ag.Log.Info("\u0061\u006c\u0069\u0067\u006e\u0065\u0064\u003d\u0025\u0074\u0020\u0075\u006em\u0061\u0074\u0063\u0068\u0065\u0064=\u0025\u002e\u0032\u0066\u003d\u0025\u0064\u002f\u0025\u0064\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _fadfb, _cbdb, _dfaf, len(_baeb), _baeb.String())
	}
	return _fadfb
}
func (_gdfag *textTable) subdivide() *textTable {
	_gdfag.logComposite("\u0073u\u0062\u0064\u0069\u0076\u0069\u0064e")
	_cecc := _gdfag.compositeRowCorridors()
	_ecbgd := _gdfag.compositeColCorridors()
	if _beae {
		_ag.Log.Info("\u0073u\u0062\u0064i\u0076\u0069\u0064\u0065:\u000a\u0009\u0072o\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s=\u0025\u0073\u000a\t\u0063\u006fl\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073=\u0025\u0073", _cgbbf(_cecc), _cgbbf(_ecbgd))
	}
	if len(_cecc) == 0 || len(_ecbgd) == 0 {
		return _gdfag
	}
	_beda(_cecc)
	_beda(_ecbgd)
	if _beae {
		_ag.Log.Info("\u0073\u0075\u0062\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0066\u0069\u0078\u0065\u0064\u003a\u000a\u0009r\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073\u003d\u0025\u0073\u000a\u0009\u0063\u006f\u006cC\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d\u0025\u0073", _cgbbf(_cecc), _cgbbf(_ecbgd))
	}
	_cgge, _abcd := _bbgcb(_gdfag._cabfg, _cecc)
	_bgac, _abfca := _bbgcb(_gdfag._baee, _ecbgd)
	_fdba := make(map[uint64]*textPara, _abfca*_abcd)
	_gdgde := &textTable{PdfRectangle: _gdfag.PdfRectangle, _acgbd: _gdfag._acgbd, _cabfg: _abcd, _baee: _abfca, _aagc: _fdba}
	if _beae {
		_ag.Log.Info("\u0073\u0075b\u0064\u0069\u0076\u0069\u0064\u0065\u003a\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0063\u0065\u006c\u006c\u0073\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0079\u004f\u0066\u0066\u0073\u0065\u0074\u0073=\u0025\u002b\u0076\u000a"+"\u0009\u0078\u004f\u0066\u0066\u0073\u0065\u0074\u0073\u003d\u0025\u002b\u0076", _gdfag._baee, _gdfag._cabfg, _abfca, _abcd, _cgbbf(_cecc), _cgbbf(_ecbgd), _cgge, _bgac)
	}
	for _aaacb := 0; _aaacb < _gdfag._cabfg; _aaacb++ {
		_ddece := _cgge[_aaacb]
		for _addeb := 0; _addeb < _gdfag._baee; _addeb++ {
			_bfede := _bgac[_addeb]
			if _beae {
				_gdf.Printf("\u0025\u0036\u0064\u002c %\u0032\u0064\u003a\u0020\u0078\u0030\u003d\u0025\u0064\u0020\u0079\u0030\u003d\u0025d\u000a", _addeb, _aaacb, _bfede, _ddece)
			}
			_gage, _egdbc := _gdfag._bbfb[_fcbc(_addeb, _aaacb)]
			if !_egdbc {
				continue
			}
			_aebg := _gage.split(_cecc[_aaacb], _ecbgd[_addeb])
			for _adbg := 0; _adbg < _aebg._cabfg; _adbg++ {
				for _fcad := 0; _fcad < _aebg._baee; _fcad++ {
					_debf := _aebg.get(_fcad, _adbg)
					_gdgde.put(_bfede+_fcad, _ddece+_adbg, _debf)
					if _beae {
						_gdf.Printf("\u0025\u0038\u0064\u002c\u0020\u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _bfede+_fcad, _ddece+_adbg, _debf)
					}
				}
			}
		}
	}
	return _gdgde
}

const (
	_ffba  = 1.0e-6
	_cegd  = 1.0e-4
	_ggdc  = 10
	_ffde  = 6
	_bgcc  = 0.5
	_gfdad = 0.12
	_edge  = 0.19
	_cbgg  = 0.04
	_ggcf  = 0.04
	_cccf  = 1.0
	_aaff  = 0.04
	_cfede = 0.4
	_bdea  = 0.7
	_efef  = 1.0
	_bee   = 0.1
	_dcca  = 1.4
	_ceda  = 0.46
	_gaca  = 0.02
	_gefd  = 0.2
	_fge   = 0.5
	_feba  = 4
	_cdge  = 4.0
	_gdfd  = 6
	_bbgg  = 0.3
	_fbcc  = 0.01
	_cfeg  = 0.02
	_agec  = 2
	_bafa  = 2
	_aedf  = 500
	_ffdeb = 4.0
	_dbgdc = 4.0
	_cdfc  = 0.05
	_gbef  = 0.1
	_deff  = 2.0
	_dcfe  = 2.0
	_egfc  = 1.5
	_dded  = 3.0
	_cagg  = 0.25
)
const _egc = 20

func _debdg(_gcca _bf.PdfRectangle) *ruling {
	return &ruling{_cgef: _beec, _eead: _gcca.Llx, _bgeb: _gcca.Lly, _eecc: _gcca.Ury}
}
func _cgbbf(_gfeb map[int][]float64) string {
	_agdb := _bacfa(_gfeb)
	_fcaeb := make([]string, len(_gfeb))
	for _aedbg, _ffgf := range _agdb {
		_fcaeb[_aedbg] = _gdf.Sprintf("\u0025\u0064\u003a\u0020\u0025\u002e\u0032\u0066", _ffgf, _gfeb[_ffgf])
	}
	return _gdf.Sprintf("\u007b\u0025\u0073\u007d", _d.Join(_fcaeb, "\u002c\u0020"))
}
func (_bacf *subpath) clear() { *_bacf = subpath{} }

// ExtractText processes and extracts all text data in content streams and returns as a string.
// It takes into account character encodings in the PDF file, which are decoded by
// CharcodeBytesToUnicode.
// Characters that can't be decoded are replaced with MissingCodeRune ('\ufffd' = �).
func (_be *Extractor) ExtractText() (string, error) {
	_fac, _, _, _dge := _be.ExtractTextWithStats()
	return _fac, _dge
}
func _cgcdb(_bfddc string) bool {
	for _, _cceeg := range _bfddc {
		if !_gg.IsSpace(_cceeg) {
			return false
		}
	}
	return true
}
func (_eeeb *textPara) toCellTextMarks(_bga *int) []TextMark {
	var _gbd []TextMark
	for _cfdc, _gfeab := range _eeeb._ecbdg {
		_cfef := _gfeab.toTextMarks(_bga)
		_ceede := _badf && _gfeab.endsInHyphen() && _cfdc != len(_eeeb._ecbdg)-1
		if _ceede {
			_cfef = _acdc(_cfef, _bga)
		}
		_gbd = append(_gbd, _cfef...)
		if !(_ceede || _cfdc == len(_eeeb._ecbdg)-1) {
			_gbd = _abbf(_gbd, _bga, _afed(_gfeab._dcfd, _eeeb._ecbdg[_cfdc+1]._dcfd))
		}
	}
	return _gbd
}
func _eag(_gbgc *Extractor, _fccd *_bf.PdfPageResources, _fbcf _da.GraphicsState, _ebg *textState, _daa *stateStack) *textObject {
	return &textObject{_ccae: _gbgc, _caa: _fccd, _fcdg: _fbcf, _afg: _daa, _degf: _ebg, _gcb: _ee.IdentityMatrix(), _baff: _ee.IdentityMatrix()}
}

// String returns a description of `t`.
func (_eeege *textTable) String() string {
	return _gdf.Sprintf("\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0025\u0074", _eeege._baee, _eeege._cabfg, _eeege._acgbd)
}

// String returns a description of `tm`.
func (_gfae *textMark) String() string {
	return _gdf.Sprintf("\u0025\u002e\u0032f \u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\u0022\u0025\u0073\u0022", _gfae.PdfRectangle, _gfae._abba, _gfae._eeaf)
}
func (_acad rulingList) snapToGroups() rulingList {
	_dfec, _fedc := _acad.vertsHorzs()
	if len(_dfec) > 0 {
		_dfec = _dfec.snapToGroupsDirection()
	}
	if len(_fedc) > 0 {
		_fedc = _fedc.snapToGroupsDirection()
	}
	_dggec := append(_dfec, _fedc...)
	_dggec.log("\u0073\u006e\u0061p\u0054\u006f\u0047\u0072\u006f\u0075\u0070\u0073")
	return _dggec
}
func (_befc *textTable) log(_dfcd string) {
	if !_beae {
		return
	}
	_ag.Log.Info("~\u007e\u007e\u0020\u0025\u0073\u003a \u0025\u0064\u0020\u0078\u0020\u0025d\u0020\u0067\u0072\u0069\u0064\u003d\u0025t\u000a\u0020\u0020\u0020\u0020\u0020\u0020\u0025\u0036\u002e2\u0066", _dfcd, _befc._baee, _befc._cabfg, _befc._acgbd, _befc.PdfRectangle)
	for _eebb := 0; _eebb < _befc._cabfg; _eebb++ {
		for _faaf := 0; _faaf < _befc._baee; _faaf++ {
			_bdefg := _befc.get(_faaf, _eebb)
			if _bdefg == nil {
				continue
			}
			_gdf.Printf("%\u0034\u0064\u0020\u00252d\u003a \u0025\u0036\u002e\u0032\u0066 \u0025\u0071\u0020\u0025\u0064\u000a", _faaf, _eebb, _bdefg.PdfRectangle, _babd(_bdefg.text(), 50), _a.RuneCountInString(_bdefg.text()))
		}
	}
}
func (_cfdbd paraList) findGridTables(_cdfb []gridTiling) []*textTable {
	if _beae {
		_ag.Log.Info("\u0066i\u006e\u0064\u0047\u0072\u0069\u0064\u0054\u0061\u0062\u006c\u0065s\u003a\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073", len(_cfdbd))
		for _aedba, _eebf := range _cfdbd {
			_gdf.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _aedba, _eebf)
		}
	}
	var _cebe []*textTable
	for _ccad, _gefaf := range _cdfb {
		_cbfc, _debgb := _cfdbd.findTableGrid(_gefaf)
		if _cbfc != nil {
			_cbfc.log(_gdf.Sprintf("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0057\u0069\u0074\u0068\u0047\u0072\u0069\u0064\u0073\u003a\u0020%\u0064", _ccad))
			_cebe = append(_cebe, _cbfc)
			_cbfc.markCells()
		}
		for _ebfea := range _debgb {
			_ebfea._bggb = true
		}
	}
	if _beae {
		_ag.Log.Info("\u0066i\u006e\u0064\u0047\u0072i\u0064\u0054\u0061\u0062\u006ce\u0073:\u0020%\u0064\u0020\u0074\u0061\u0062\u006c\u0065s", len(_cebe))
	}
	return _cebe
}

type stateStack []*textState

func (_gagf *textWord) toTextMarks(_cfdg *int) []TextMark {
	var _fbgae []TextMark
	for _, _gfdaf := range _gagf._gdgg {
		_fbgae = _affac(_fbgae, _cfdg, _gfdaf.ToTextMark())
	}
	return _fbgae
}
func (_gaee *stateStack) pop() *textState {
	if _gaee.empty() {
		return nil
	}
	_fbb := *(*_gaee)[len(*_gaee)-1]
	*_gaee = (*_gaee)[:len(*_gaee)-1]
	return &_fbb
}
func (_eagd *textTable) put(_cagd, _bgccbg int, _efcdc *textPara) {
	_eagd._aagc[_fcbc(_cagd, _bgccbg)] = _efcdc
}
func (_dff *wordBag) empty(_gggb int) bool {
	_, _accgge := _dff._bbf[_gggb]
	return !_accgge
}
func (_fgbb *textTable) compositeRowCorridors() map[int][]float64 {
	_gdbcfg := make(map[int][]float64, _fgbb._cabfg)
	if _beae {
		_ag.Log.Info("c\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0052\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073:\u0020h\u003d\u0025\u0064", _fgbb._cabfg)
	}
	for _ebad := 1; _ebad < _fgbb._cabfg; _ebad++ {
		var _fbde []compositeCell
		for _efgb := 0; _efgb < _fgbb._baee; _efgb++ {
			if _gaeff, _ggef := _fgbb._bbfb[_fcbc(_efgb, _ebad)]; _ggef {
				_fbde = append(_fbde, _gaeff)
			}
		}
		if len(_fbde) == 0 {
			continue
		}
		_bddag := _eagab(_fbde)
		_gdbcfg[_ebad] = _bddag
		if _beae {
			_gdf.Printf("\u0020\u0020\u0020\u0025\u0032\u0064\u003a\u0020\u00256\u002e\u0032\u0066\u000a", _ebad, _bddag)
		}
	}
	return _gdbcfg
}

type bounded interface{ bbox() _bf.PdfRectangle }

func (_acdd *textPara) writeText(_fdbf _ed.Writer) {
	if _acdd._bccg == nil {
		_acdd.writeCellText(_fdbf)
		return
	}
	for _cgag := 0; _cgag < _acdd._bccg._cabfg; _cgag++ {
		for _bfee := 0; _bfee < _acdd._bccg._baee; _bfee++ {
			_gacgc := _acdd._bccg.get(_bfee, _cgag)
			if _gacgc == nil {
				_fdbf.Write([]byte("\u0009"))
			} else {
				_gacgc.writeCellText(_fdbf)
			}
			_fdbf.Write([]byte("\u0020"))
		}
		if _cgag < _acdd._bccg._cabfg-1 {
			_fdbf.Write([]byte("\u000a"))
		}
	}
}
func (_cdaeb *textTable) logComposite(_ccfaf string) {
	if !_beae {
		return
	}
	_ag.Log.Info("\u007e~\u007eP\u0061\u0072\u0061\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _cdaeb._baee, _cdaeb._cabfg, _ccfaf)
	_gdf.Printf("\u0025\u0035\u0073 \u007c", "")
	for _deebe := 0; _deebe < _cdaeb._baee; _deebe++ {
		_gdf.Printf("\u0025\u0033\u0064 \u007c", _deebe)
	}
	_gdf.Println("")
	_gdf.Printf("\u0025\u0035\u0073 \u002b", "")
	for _eeab := 0; _eeab < _cdaeb._baee; _eeab++ {
		_gdf.Printf("\u0025\u0033\u0073 \u002b", "\u002d\u002d\u002d")
	}
	_gdf.Println("")
	for _ggeg := 0; _ggeg < _cdaeb._cabfg; _ggeg++ {
		_gdf.Printf("\u0025\u0035\u0064 \u007c", _ggeg)
		for _ccbg := 0; _ccbg < _cdaeb._baee; _ccbg++ {
			_beee, _ := _cdaeb._bbfb[_fcbc(_ccbg, _ggeg)].parasBBox()
			_gdf.Printf("\u0025\u0033\u0064 \u007c", len(_beee))
		}
		_gdf.Println("")
	}
	_ag.Log.Info("\u007e~\u007eT\u0065\u0078\u0074\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _cdaeb._baee, _cdaeb._cabfg, _ccfaf)
	_gdf.Printf("\u0025\u0035\u0073 \u007c", "")
	for _fgdbd := 0; _fgdbd < _cdaeb._baee; _fgdbd++ {
		_gdf.Printf("\u0025\u0031\u0032\u0064\u0020\u007c", _fgdbd)
	}
	_gdf.Println("")
	_gdf.Printf("\u0025\u0035\u0073 \u002b", "")
	for _cageeg := 0; _cageeg < _cdaeb._baee; _cageeg++ {
		_gdf.Print("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d\u002b")
	}
	_gdf.Println("")
	for _bcceg := 0; _bcceg < _cdaeb._cabfg; _bcceg++ {
		_gdf.Printf("\u0025\u0035\u0064 \u007c", _bcceg)
		for _fdef := 0; _fdef < _cdaeb._baee; _fdef++ {
			_ddbf, _ := _cdaeb._bbfb[_fcbc(_fdef, _bcceg)].parasBBox()
			_aeaef := ""
			_cafcg := _ddbf.merge()
			if _cafcg != nil {
				_aeaef = _cafcg.text()
			}
			_aeaef = _gdf.Sprintf("\u0025\u0071", _babd(_aeaef, 12))
			_aeaef = _aeaef[1 : len(_aeaef)-1]
			_gdf.Printf("\u0025\u0031\u0032\u0073\u0020\u007c", _aeaef)
		}
		_gdf.Println("")
	}
}

const (
	_ceabd = false
	_eccd  = false
	_ddb   = false
	_dbcf  = false
	_bccf  = false
	_dbde  = false
	_gfcbc = false
	_dada  = false
	_cfed  = false
	_bcf   = _cfed && true
	_aebf  = _bcf && false
	_dbbge = _cfed && true
	_beae  = false
	_ddgf  = _beae && false
	_fcff  = _beae && true
	_efa   = false
	_bfga  = _efa && false
	_cecg  = _efa && false
	_efbc  = _efa && true
	_bded  = _efa && false
	_fgad  = _efa && false
)

func _geccf(_facc float64) bool { return _b.Abs(_facc) < _dcfe }

// Len returns the number of TextMarks in `ma`.
func (_bdge *TextMarkArray) Len() int {
	if _bdge == nil {
		return 0
	}
	return len(_bdge._fceg)
}

// NewFromContents creates a new extractor from contents and page resources.
func NewFromContents(contents string, resources *_bf.PdfPageResources) (*Extractor, error) {
	_acg := &Extractor{_fc: contents, _ad: resources, _ac: map[string]fontEntry{}, _fd: map[string]textResult{}}
	return _acg, nil
}
func (_egcf *stateStack) push(_fee *textState) { _cbdd := *_fee; *_egcf = append(*_egcf, &_cbdd) }
func (_ffbd *textObject) moveTextSetLeading(_cba, _cef float64) {
	_ffbd._degf._acfa = -_cef
	_ffbd.moveLP(_cba, _cef)
}
func _cedeb(_egbce []*textMark, _eebe _bf.PdfRectangle) *textWord {
	_gegc := _egbce[0].PdfRectangle
	_fbcfg := _egbce[0]._abba
	for _, _agga := range _egbce[1:] {
		_gegc = _effa(_gegc, _agga.PdfRectangle)
		if _agga._abba > _fbcfg {
			_fbcfg = _agga._abba
		}
	}
	return &textWord{PdfRectangle: _gegc, _gdgg: _egbce, _ecgcg: _eebe.Ury - _gegc.Lly, _fbgge: _fbcfg}
}
func _eeee(_febg _bf.PdfRectangle) *ruling {
	return &ruling{_cgef: _ebabd, _eead: _febg.Ury, _bgeb: _febg.Llx, _eecc: _febg.Urx}
}
func (_geca *textLine) bbox() _bf.PdfRectangle { return _geca.PdfRectangle }
func (_gdfg rulingList) toTilings() (rulingList, []gridTiling) {
	_gdfg.log("\u0074o\u0054\u0069\u006c\u0069\u006e\u0067s")
	if len(_gdfg) == 0 {
		return nil, nil
	}
	_gdfg = _gdfg.tidied("\u0061\u006c\u006c")
	_gdfg.log("\u0074\u0069\u0064\u0069\u0065\u0064")
	_abef := _gdfg.toGrids()
	_cafd := make([]gridTiling, len(_abef))
	for _acgff, _bfde := range _abef {
		_cafd[_acgff] = _bfde.asTiling()
	}
	return _gdfg, _cafd
}
func _cdag(_fbegd *wordBag, _cffc *textWord, _cdbg float64) bool {
	return _cffc.Llx < _fbegd.Urx+_cdbg && _fbegd.Llx-_cdbg < _cffc.Urx
}
func (_cfefe *textTable) toTextTable() TextTable {
	if _beae {
		_ag.Log.Info("t\u006fT\u0065\u0078\u0074\u0054\u0061\u0062\u006c\u0065:\u0020\u0025\u0064\u0020x \u0025\u0064", _cfefe._baee, _cfefe._cabfg)
	}
	_adcb := make([][]TableCell, _cfefe._cabfg)
	for _abebf := 0; _abebf < _cfefe._cabfg; _abebf++ {
		_adcb[_abebf] = make([]TableCell, _cfefe._baee)
		for _dafbg := 0; _dafbg < _cfefe._baee; _dafbg++ {
			_eafb := _cfefe.get(_dafbg, _abebf)
			if _eafb == nil {
				continue
			}
			if _beae {
				_gdf.Printf("\u0025\u0034\u0064 \u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _dafbg, _abebf, _eafb)
			}
			_adcb[_abebf][_dafbg].Text = _eafb.text()
			_fcceb := 0
			_adcb[_abebf][_dafbg].Marks._fceg = _eafb.toTextMarks(&_fcceb)
		}
	}
	return TextTable{W: _cfefe._baee, H: _cfefe._cabfg, Cells: _adcb}
}
func _efeg(_faf, _dggg _bf.PdfRectangle) bool { return _faf.Lly <= _dggg.Ury && _dggg.Lly <= _faf.Ury }
func (_fbcg rulingList) bbox() _bf.PdfRectangle {
	var _gfcab _bf.PdfRectangle
	if len(_fbcg) == 0 {
		_ag.Log.Error("r\u0075\u006c\u0069\u006e\u0067\u004ci\u0073\u0074\u002e\u0062\u0062\u006f\u0078\u003a\u0020n\u006f\u0020\u0072u\u006ci\u006e\u0067\u0073")
		return _bf.PdfRectangle{}
	}
	if _fbcg[0]._cgef == _ebabd {
		_gfcab.Llx, _gfcab.Urx = _fbcg.secMinMax()
		_gfcab.Lly, _gfcab.Ury = _fbcg.primMinMax()
	} else {
		_gfcab.Llx, _gfcab.Urx = _fbcg.primMinMax()
		_gfcab.Lly, _gfcab.Ury = _fbcg.secMinMax()
	}
	return _gfcab
}

type textResult struct {
	_bafb PageText
	_dfcc int
	_bcbb int
}

func (_eecd *shapesState) devicePoint(_eaaa, _dbdc float64) _ee.Point {
	_fbfd := _eecd._efgg.Mult(_eecd._egb)
	_eaaa, _dbdc = _fbfd.Transform(_eaaa, _dbdc)
	return _ee.NewPoint(_eaaa, _dbdc)
}
func (_eabe *wordBag) removeDuplicates() {
	if _dbbge {
		_ag.Log.Info("r\u0065m\u006f\u0076\u0065\u0044\u0075\u0070\u006c\u0069c\u0061\u0074\u0065\u0073: \u0025\u0071", _eabe.text())
	}
	for _, _bfdg := range _eabe.depthIndexes() {
		if len(_eabe._bbf[_bfdg]) == 0 {
			continue
		}
		_edbec := _eabe._bbf[_bfdg][0]
		_cded := _gefd * _edbec._fbgge
		_dgeef := _edbec._ecgcg
		for _, _fdbec := range _eabe.depthBand(_dgeef, _dgeef+_cded) {
			_fgffd := map[*textWord]struct{}{}
			_fda := _eabe._bbf[_fdbec]
			for _, _cecgg := range _fda {
				if _, _edgdd := _fgffd[_cecgg]; _edgdd {
					continue
				}
				for _, _cdfg := range _fda {
					if _, _cgaga := _fgffd[_cdfg]; _cgaga {
						continue
					}
					if _cdfg != _cecgg && _cdfg._bcaa == _cecgg._bcaa && _b.Abs(_cdfg.Llx-_cecgg.Llx) < _cded && _b.Abs(_cdfg.Urx-_cecgg.Urx) < _cded && _b.Abs(_cdfg.Lly-_cecgg.Lly) < _cded && _b.Abs(_cdfg.Ury-_cecgg.Ury) < _cded {
						_fgffd[_cdfg] = struct{}{}
					}
				}
			}
			if len(_fgffd) > 0 {
				_eadd := 0
				for _, _ebabg := range _fda {
					if _, _affg := _fgffd[_ebabg]; !_affg {
						_fda[_eadd] = _ebabg
						_eadd++
					}
				}
				_eabe._bbf[_fdbec] = _fda[:len(_fda)-len(_fgffd)]
				if len(_eabe._bbf[_fdbec]) == 0 {
					delete(_eabe._bbf, _fdbec)
				}
			}
		}
	}
}
func (_eeac rulingList) findPrimSec(_gcbe, _efee float64) *ruling {
	for _, _gdbbd := range _eeac {
		if _fbga(_gdbbd._eead-_gcbe) && _gdbbd._bgeb-_deff <= _efee && _efee <= _gdbbd._eecc+_deff {
			return _gdbbd
		}
	}
	return nil
}
