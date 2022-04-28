package extractor

import (
	_gb "bytes"
	_g "errors"
	_ca "fmt"
	_af "image/color"
	_a "io"
	_c "math"
	_d "regexp"
	_dc "sort"
	_ad "strings"
	_ag "unicode"
	_ac "unicode/utf8"

	_f "bitbucket.org/shenghui0779/gopdf/common"
	_da "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ff "bitbucket.org/shenghui0779/gopdf/core"
	_bg "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_bc "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ba "bitbucket.org/shenghui0779/gopdf/model"
	_dce "golang.org/x/text/unicode/norm"
	_ce "golang.org/x/xerrors"
)

func (_ggabf *ruling) encloses(_cfaag, _cadcd float64) bool {
	return _ggabf._gcbgg-_gdae <= _cfaag && _cadcd <= _ggabf._cgeee+_gdae
}
func (_bbfc *textObject) showText(_bdf []byte) error { return _bbfc.renderText(_bdf) }

// TextMarkArray is a collection of TextMarks.
type TextMarkArray struct{ _afdf []TextMark }

func _cafce(_ggcf, _eaa _ba.PdfRectangle) _ba.PdfRectangle {
	return _ba.PdfRectangle{Llx: _c.Min(_ggcf.Llx, _eaa.Llx), Lly: _c.Min(_ggcf.Lly, _eaa.Lly), Urx: _c.Max(_ggcf.Urx, _eaa.Urx), Ury: _c.Max(_ggcf.Ury, _eaa.Ury)}
}
func _dcb(_ecd []Font, _fa string) bool {
	for _, _fd := range _ecd {
		if _fd.FontName == _fa {
			return true
		}
	}
	return false
}

type stateStack []*textState

// Font represents the font properties on a PDF page.
type Font struct {
	PdfFont *_ba.PdfFont

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
	FontDescriptor *_ba.PdfFontDescriptor
}

func (_ffbc rulingList) toGrids() []rulingList {
	if _bfg {
		_f.Log.Info("t\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _ffbc)
	}
	_cafac := _ffbc.intersections()
	if _bfg {
		_f.Log.Info("\u0074\u006f\u0047r\u0069\u0064\u0073\u003a \u0076\u0065\u0063\u0073\u003d\u0025\u0064 \u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003d\u0025\u0064\u0020", len(_ffbc), len(_cafac))
		for _, _fcea := range _egcga(_cafac) {
			_ca.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _fcea, _cafac[_fcea])
		}
	}
	_ffde := make(map[int]intSet, len(_ffbc))
	for _abcb := range _ffbc {
		_bdddc := _ffbc.connections(_cafac, _abcb)
		if len(_bdddc) > 0 {
			_ffde[_abcb] = _bdddc
		}
	}
	if _bfg {
		_f.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0063\u006fn\u006e\u0065\u0063\u0074s=\u0025\u0064", len(_ffde))
		for _, _ebfg := range _egcga(_ffde) {
			_ca.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _ebfg, _ffde[_ebfg])
		}
	}
	_bba := _feadg(len(_ffbc), func(_gbdge, _gaaeb int) bool {
		_cbec, _dcac := len(_ffde[_gbdge]), len(_ffde[_gaaeb])
		if _cbec != _dcac {
			return _cbec > _dcac
		}
		return _ffbc.comp(_gbdge, _gaaeb)
	})
	if _bfg {
		_f.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u006f\u0072d\u0065\u0072\u0069\u006eg=\u0025\u0076", _bba)
	}
	_cbeae := [][]int{{_bba[0]}}
_gfdg:
	for _, _aeba := range _bba[1:] {
		for _cgfd, _dbfe := range _cbeae {
			for _, _cbdb := range _dbfe {
				if _ffde[_cbdb].has(_aeba) {
					_cbeae[_cgfd] = append(_dbfe, _aeba)
					continue _gfdg
				}
			}
		}
		_cbeae = append(_cbeae, []int{_aeba})
	}
	if _bfg {
		_f.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020i\u0067\u0072i\u0064\u0073\u003d\u0025\u0076", _cbeae)
	}
	_dc.SliceStable(_cbeae, func(_bgbb, _bfeab int) bool { return len(_cbeae[_bgbb]) > len(_cbeae[_bfeab]) })
	for _, _ffga := range _cbeae {
		_dc.Slice(_ffga, func(_ccge, _egee int) bool { return _ffbc.comp(_ffga[_ccge], _ffga[_egee]) })
	}
	_eaad := make([]rulingList, len(_cbeae))
	for _eabcd, _egafe := range _cbeae {
		_decg := make(rulingList, len(_egafe))
		for _dffg, _bccd := range _egafe {
			_decg[_dffg] = _ffbc[_bccd]
		}
		_eaad[_eabcd] = _decg
	}
	if _bfg {
		_f.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020g\u0072\u0069d\u0073\u003d\u0025\u002b\u0076", _eaad)
	}
	var _cagdb []rulingList
	for _, _gdgd := range _eaad {
		if _agdb, _begde := _gdgd.isActualGrid(); _begde {
			_gdgd = _agdb
			_gdgd = _gdgd.snapToGroups()
			_cagdb = append(_cagdb, _gdgd)
		}
	}
	if _bfg {
		_dcfba("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0061\u0063t\u0075\u0061\u006c\u0047ri\u0064\u0073", _cagdb)
		_f.Log.Info("\u0074\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0067\u0072\u0069\u0064\u0073\u003d%\u0064 \u0061\u0063\u0074\u0075\u0061\u006c\u0047\u0072\u0069\u0064\u0073\u003d\u0025\u0064", len(_eaad), len(_cagdb))
	}
	return _cagdb
}

var (
	_fb = _g.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_bb = _g.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func (_cafc *subpath) add(_afca ..._bc.Point) { _cafc._caf = append(_cafc._caf, _afca...) }
func (_cgga *textObject) getFontDirect(_cgee string) (*_ba.PdfFont, error) {
	_gage, _eegb := _cgga.getFontDict(_cgee)
	if _eegb != nil {
		return nil, _eegb
	}
	_bcae, _eegb := _ba.NewPdfFontFromPdfObject(_gage)
	if _eegb != nil {
		_f.Log.Debug("\u0067\u0065\u0074\u0046\u006f\u006e\u0074\u0044\u0069\u0072\u0065\u0063\u0074\u003a\u0020\u004e\u0065\u0077Pd\u0066F\u006f\u006e\u0074\u0046\u0072\u006f\u006d\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d%\u0023\u0071\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cgee, _eegb)
	}
	return _bcae, _eegb
}
func (_abdd *textLine) bbox() _ba.PdfRectangle { return _abdd.PdfRectangle }
func (_dea *shapesState) devicePoint(_fdeb, _dfd float64) _bc.Point {
	_beg := _dea._gfgg.Mult(_dea._bfc)
	_fdeb, _dfd = _beg.Transform(_fdeb, _dfd)
	return _bc.NewPoint(_fdeb, _dfd)
}

// String returns a string descibing `i`.
func (_cegd gridTile) String() string {
	_fcae := func(_ebfad bool, _edeaa string) string {
		if _ebfad {
			return _edeaa
		}
		return "\u005f"
	}
	return _ca.Sprintf("\u00256\u002e2\u0066\u0020\u0025\u0031\u0073%\u0031\u0073%\u0031\u0073\u0025\u0031\u0073", _cegd.PdfRectangle, _fcae(_cegd._fdge, "\u004c"), _fcae(_cegd._abbbb, "\u0052"), _fcae(_cegd._edcg, "\u0042"), _fcae(_cegd._aebbe, "\u0054"))
}

// String returns a human readable description of `ss`.
func (_ggc *shapesState) String() string {
	return _ca.Sprintf("\u007b\u0025\u0064\u0020su\u0062\u0070\u0061\u0074\u0068\u0073\u0020\u0066\u0072\u0065\u0073\u0068\u003d\u0025t\u007d", len(_ggc._eagg), _ggc._debd)
}
func (_aeg *wordBag) blocked(_fdag *textWord) bool {
	if _fdag.Urx < _aeg.Llx {
		_becg := _cebb(_fdag.PdfRectangle)
		_bgac := _gceeb(_aeg.PdfRectangle)
		if _aeg._edga.blocks(_becg, _bgac) {
			if _ggfd {
				_f.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0078\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _fdag, _aeg)
			}
			return true
		}
	} else if _aeg.Urx < _fdag.Llx {
		_geb := _cebb(_aeg.PdfRectangle)
		_fcad := _gceeb(_fdag.PdfRectangle)
		if _aeg._edga.blocks(_geb, _fcad) {
			if _ggfd {
				_f.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0078\u2192\u0020:\u0020\u0025\u0073\u0020%s", _fdag, _aeg)
			}
			return true
		}
	}
	if _fdag.Ury < _aeg.Lly {
		_fdcb := _efba(_fdag.PdfRectangle)
		_dcge := _gcge(_aeg.PdfRectangle)
		if _aeg._gaef.blocks(_fdcb, _dcge) {
			if _ggfd {
				_f.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0079\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _fdag, _aeg)
			}
			return true
		}
	} else if _aeg.Ury < _fdag.Lly {
		_gbedd := _efba(_aeg.PdfRectangle)
		_bffd := _gcge(_fdag.PdfRectangle)
		if _aeg._gaef.blocks(_gbedd, _bffd) {
			if _ggfd {
				_f.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0079\u2192\u0020:\u0020\u0025\u0073\u0020%s", _fdag, _aeg)
			}
			return true
		}
	}
	return false
}
func (_gedcc paraList) writeText(_ddda _a.Writer) {
	for _aeed, _ccfg := range _gedcc {
		if _ccfg._bgea {
			continue
		}
		_ccfg.writeText(_ddda)
		if _aeed != len(_gedcc)-1 {
			if _dadb(_ccfg, _gedcc[_aeed+1]) {
				_ddda.Write([]byte("\u0020"))
			} else {
				_ddda.Write([]byte("\u000a"))
				_ddda.Write([]byte("\u000a"))
			}
		}
	}
	_ddda.Write([]byte("\u000a"))
	_ddda.Write([]byte("\u000a"))
}
func (_bcdag *subpath) clear() { *_bcdag = subpath{} }

type shapesState struct {
	_bfc  _bc.Matrix
	_gfgg _bc.Matrix
	_eagg []*subpath
	_debd bool
	_ebeg _bc.Point
	_gfad *textObject
}

func (_ffd *wordBag) firstWord(_ecbg int) *textWord { return _ffd._dfec[_ecbg][0] }
func (_bdgf intSet) del(_dbdgb int)                 { delete(_bdgf, _dbdgb) }

// PageFonts represents extracted fonts on a PDF page.
type PageFonts struct{ Fonts []Font }

func (_aaaf *shapesState) fill(_fgec *[]pathSection) {
	_ggbd := pathSection{_cag: _aaaf._eagg, Color: _aaaf._gfad.getFillColor()}
	*_fgec = append(*_fgec, _ggbd)
	if _bfg {
		_aebfa := _ggbd.bbox()
		_ca.Printf("\u0020 \u0020\u0020\u0046\u0049\u004c\u004c\u003a %\u0032\u0064\u0020\u0066\u0069\u006c\u006c\u0073\u0020\u0028\u0025\u0064\u0020\u006ee\u0077\u0029 \u0073\u0073\u003d%\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d\u0025\u0033\u0076\u0020\u0025\u0036\u002e\u0032f\u003d\u00256.\u0032\u0066\u0078%\u0036\u002e\u0032\u0066\u000a", len(*_fgec), len(_ggbd._cag), _aaaf, _ggbd.Color, _aebfa, _aebfa.Width(), _aebfa.Height())
		if _bdda {
			for _fdce, _cege := range _ggbd._cag {
				_ca.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _fdce, _cege)
				if _fdce == 10 {
					break
				}
			}
		}
	}
}
func _ebdga(_eafga []*textWord, _efac *textWord) []*textWord {
	for _bbde, _dfaba := range _eafga {
		if _dfaba == _efac {
			return _ecaa(_eafga, _bbde)
		}
	}
	_f.Log.Error("\u0072\u0065\u006d\u006f\u0076e\u0057\u006f\u0072\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u0073\u0020\u0064o\u0065\u0073\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", _efac)
	return nil
}

type wordBag struct {
	_ba.PdfRectangle
	_egad        float64
	_edga, _gaef rulingList
	_efae        float64
	_dfec        map[int][]*textWord
}

func _cgefb(_ffcfd []_ff.PdfObject) (_bfeef, _beaed float64, _ceecg error) {
	if len(_ffcfd) != 2 {
		return 0, 0, _ca.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0073\u003a \u0025\u0064", len(_ffcfd))
	}
	_ageaa, _ceecg := _ff.GetNumbersAsFloat(_ffcfd)
	if _ceecg != nil {
		return 0, 0, _ceecg
	}
	return _ageaa[0], _ageaa[1], nil
}
func (_ege *stateStack) size() int             { return len(*_ege) }
func _dfee(_caad, _acdea int) uint64           { return uint64(_caad)*0x1000000 + uint64(_acdea) }
func (_gac *stateStack) push(_adaf *textState) { _ccg := *_adaf; *_gac = append(*_gac, &_ccg) }
func (_efgb gridTile) numBorders() int {
	_dgfc := 0
	if _efgb._fdge {
		_dgfc++
	}
	if _efgb._abbbb {
		_dgfc++
	}
	if _efgb._edcg {
		_dgfc++
	}
	if _efgb._aebbe {
		_dgfc++
	}
	return _dgfc
}
func _aged(_bacgd, _gaff bounded) float64 { return _bacgd.bbox().Llx - _gaff.bbox().Urx }
func _ebag(_cfcca, _efdd _ba.PdfRectangle) bool {
	return _cfcca.Llx <= _efdd.Llx && _efdd.Urx <= _cfcca.Urx && _cfcca.Lly <= _efdd.Lly && _efdd.Ury <= _cfcca.Ury
}

type gridTile struct {
	_ba.PdfRectangle
	_aebbe, _fdge, _edcg, _abbbb bool
}

func (_gegg *textMark) inDiacriticArea(_acge *textMark) bool {
	_gbgg := _gegg.Llx - _acge.Llx
	_edbd := _gegg.Urx - _acge.Urx
	_cfgg := _gegg.Lly - _acge.Lly
	return _c.Abs(_gbgg+_edbd) < _gegg.Width()*_egfd && _c.Abs(_cfgg) < _gegg.Height()*_egfd
}
func (_cddbg *textTable) newTablePara() *textPara {
	_bddc := _cddbg.computeBbox()
	_eacca := &textPara{PdfRectangle: _bddc, _dabb: _bddc, _ebcb: _cddbg}
	if _eca {
		_f.Log.Info("\u006e\u0065w\u0054\u0061\u0062l\u0065\u0050\u0061\u0072\u0061\u003a\u0020\u0025\u0073", _eacca)
	}
	return _eacca
}
func (_cefd *wordBag) applyRemovals(_abdc map[int]map[*textWord]struct{}) {
	for _ceb, _eabab := range _abdc {
		if len(_eabab) == 0 {
			continue
		}
		_ceae := _cefd._dfec[_ceb]
		_dfeg := len(_ceae) - len(_eabab)
		if _dfeg == 0 {
			delete(_cefd._dfec, _ceb)
			continue
		}
		_baf := make([]*textWord, _dfeg)
		_dcgc := 0
		for _, _aceec := range _ceae {
			if _, _afcag := _eabab[_aceec]; !_afcag {
				_baf[_dcgc] = _aceec
				_dcgc++
			}
		}
		_cefd._dfec[_ceb] = _baf
	}
}

// ToText returns the page text as a single string.
// Deprecated: This function is deprecated and will be removed in a future major version. Please use
// Text() instead.
func (_gaga PageText) ToText() string { return _gaga.Text() }
func (_edfdb rulingList) tidied(_cagd string) rulingList {
	_eabd := _edfdb.removeDuplicates()
	_eabd.log("\u0075n\u0069\u0071\u0075\u0065\u0073")
	_agcea := _eabd.snapToGroups()
	if _agcea == nil {
		return nil
	}
	_agcea.sort()
	if _bfg {
		_f.Log.Info("\u0074\u0069\u0064i\u0065\u0064\u003a\u0020\u0025\u0071\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u0075\u006e\u0069\u0071\u0075\u0065\u0073\u003d\u0025\u0064\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0064\u003d\u0025\u0064", _cagd, len(_edfdb), len(_eabd), len(_agcea))
	}
	_agcea.log("\u0063o\u0061\u006c\u0065\u0073\u0063\u0065d")
	return _agcea
}

// String returns a description of `l`.
func (_cegg *textLine) String() string {
	return _ca.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _cegg._fcaa, _cegg.PdfRectangle, _cegg._edfec, _cegg.text())
}

// String returns a description of `t`.
func (_fgedb *textTable) String() string {
	return _ca.Sprintf("\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0025\u0074", _fgedb._gcead, _fgedb._ebab, _fgedb._deecf)
}
func (_dfdd rulingList) blocks(_ffae, _aefeg *ruling) bool {
	if _ffae._gcbgg > _aefeg._cgeee || _aefeg._gcbgg > _ffae._cgeee {
		return false
	}
	_egdee := _c.Max(_ffae._gcbgg, _aefeg._gcbgg)
	_afcge := _c.Min(_ffae._cgeee, _aefeg._cgeee)
	if _ffae._ccbf > _aefeg._ccbf {
		_ffae, _aefeg = _aefeg, _ffae
	}
	for _, _cafee := range _dfdd {
		if _ffae._ccbf <= _cafee._ccbf+_gedc && _cafee._ccbf <= _aefeg._ccbf+_gedc && _cafee._gcbgg <= _afcge && _egdee <= _cafee._cgeee {
			return true
		}
	}
	return false
}
func _bfgf(_bbee, _agaee _bc.Point) bool {
	_dafe := _c.Abs(_bbee.X - _agaee.X)
	_dfbd := _c.Abs(_bbee.Y - _agaee.Y)
	return _cded(_dafe, _dfbd)
}
func _egcga(_bgbgbb map[int]intSet) []int {
	_afgf := make([]int, 0, len(_bgbgbb))
	for _daef := range _bgbgbb {
		_afgf = append(_afgf, _daef)
	}
	_dc.Ints(_afgf)
	return _afgf
}

// String returns a string describing `tm`.
func (_dcfd TextMark) String() string {
	_bggdd := _dcfd.BBox
	var _gbdaf string
	if _dcfd.Font != nil {
		_gbdaf = _dcfd.Font.String()
		if len(_gbdaf) > 50 {
			_gbdaf = _gbdaf[:50] + "\u002e\u002e\u002e"
		}
	}
	var _cdee string
	if _dcfd.Meta {
		_cdee = "\u0020\u002a\u004d\u002a"
	}
	return _ca.Sprintf("\u007b\u0054\u0065\u0078t\u004d\u0061\u0072\u006b\u003a\u0020\u0025\u0064\u0020%\u0071\u003d\u0025\u0030\u0032\u0078\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e2\u0066\u0029\u0020\u0028\u00256\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0025\u0073\u007d", _dcfd.Offset, _dcfd.Text, []rune(_dcfd.Text), _bggdd.Llx, _bggdd.Lly, _bggdd.Urx, _bggdd.Ury, _gbdaf, _cdee)
}

const (
	_dedbg = true
	_bgbd  = true
	_edfe  = true
	_cbgf  = false
	_fbfd  = false
	_fcadd = 6
	_fab   = 3.0
	_abg   = 200
	_eege  = true
	_fcda  = true
	_fbbc  = true
	_afbab = true
	_dbbg  = false
)

// ExtractPageImages returns the image contents of the page extractor, including data
// and position, size information for each image.
// A set of options to control page image extraction can be passed in. The options
// parameter can be nil for the default options. By default, inline stencil masks
// are not extracted.
func (_dga *Extractor) ExtractPageImages(options *ImageExtractOptions) (*PageImages, error) {
	_fbef := &imageExtractContext{_caa: options}
	_agc := _fbef.extractContentStreamImages(_dga._ge, _dga._fbe)
	if _agc != nil {
		return nil, _agc
	}
	return &PageImages{Images: _fbef._aeb}, nil
}
func (_ccbg *wordBag) makeRemovals() map[int]map[*textWord]struct{} {
	_aefe := make(map[int]map[*textWord]struct{}, len(_ccbg._dfec))
	for _bacg := range _ccbg._dfec {
		_aefe[_bacg] = make(map[*textWord]struct{})
	}
	return _aefe
}
func (_edfdd intSet) has(_fbbcb int) bool {
	_, _egbcd := _edfdd[_fbbcb]
	return _egbcd
}
func _ffff(_ceedd int, _feeba map[int][]float64) ([]int, int) {
	_cbfd := make([]int, _ceedd)
	_aggc := 0
	for _ccga := 0; _ccga < _ceedd; _ccga++ {
		_cbfd[_ccga] = _aggc
		_aggc += len(_feeba[_ccga]) + 1
	}
	return _cbfd, _aggc
}

type textState struct {
	_aed   float64
	_cef   float64
	_fgdd  float64
	_dgbbf float64
	_bgg   float64
	_adcb  RenderMode
	_aea   float64
	_fdf   *_ba.PdfFont
	_gbff  _ba.PdfRectangle
	_efc   int
	_fcge  int
}

func (_adcgc paraList) inTile(_agfgd gridTile) paraList {
	var _fefa paraList
	for _, _gbeb := range _adcgc {
		if _agfgd.contains(_gbeb.PdfRectangle) {
			_fefa = append(_fefa, _gbeb)
		}
	}
	if _eca {
		_ca.Printf("\u0020 \u0020\u0069\u006e\u0054i\u006c\u0065\u003a\u0020\u0020%\u0073 \u0069n\u0073\u0069\u0064\u0065\u003d\u0025\u0064\n", _agfgd, len(_fefa))
		for _bdcfb, _dbcf := range _fefa {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bdcfb, _dbcf)
		}
		_ca.Println("")
	}
	return _fefa
}
func (_afc *textObject) showTextAdjusted(_ccce *_ff.PdfObjectArray) error {
	_gcg := false
	for _, _fdc := range _ccce.Elements() {
		switch _fdc.(type) {
		case *_ff.PdfObjectFloat, *_ff.PdfObjectInteger:
			_ede, _fbd := _ff.GetNumberAsFloat(_fdc)
			if _fbd != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078t\u0041\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0042\u0061\u0064\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0061\u006c\u0020a\u0072\u0067\u002e\u0020\u006f\u003d\u0025\u0073\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _fdc, _ccce)
				return _fbd
			}
			_gcd, _dbc := -_ede*0.001*_afc._gdcg._bgg, 0.0
			if _gcg {
				_dbc, _gcd = _gcd, _dbc
			}
			_egf := _cfdg(_bc.Point{X: _gcd, Y: _dbc})
			_afc._ddfg.Concat(_egf)
		case *_ff.PdfObjectString:
			_daae, _bgcg := _ff.GetStringBytes(_fdc)
			if !_bgcg {
				_f.Log.Trace("s\u0068\u006f\u0077\u0054\u0065\u0078\u0074\u0041\u0064j\u0075\u0073\u0074\u0065\u0064\u003a\u0020Ba\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u002e\u0020o=\u0025\u0073 \u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _fdc, _ccce)
				return _ff.ErrTypeError
			}
			_afc.renderText(_daae)
		default:
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078\u0074A\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0055\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0028%T\u0029\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _fdc, _ccce)
			return _ff.ErrTypeError
		}
	}
	return nil
}
func _baeb(_ddcb _ba.PdfRectangle, _bffb bounded) float64 { return _ddcb.Ury - _bffb.bbox().Lly }
func (_cbbca rulingList) merge() *ruling {
	_dabc := _cbbca[0]._ccbf
	_aced := _cbbca[0]._gcbgg
	_deccc := _cbbca[0]._cgeee
	for _, _fdfee := range _cbbca[1:] {
		_dabc += _fdfee._ccbf
		if _fdfee._gcbgg < _aced {
			_aced = _fdfee._gcbgg
		}
		if _fdfee._cgeee > _deccc {
			_deccc = _fdfee._cgeee
		}
	}
	_fecab := &ruling{_ecddg: _cbbca[0]._ecddg, _gdee: _cbbca[0]._gdee, Color: _cbbca[0].Color, _ccbf: _dabc / float64(len(_cbbca)), _gcbgg: _aced, _cgeee: _deccc}
	if _dcdb {
		_f.Log.Info("\u006de\u0072g\u0065\u003a\u0020\u0025\u0032d\u0020\u0076e\u0063\u0073\u0020\u0025\u0073", len(_cbbca), _fecab)
		for _gdbg, _cabb := range _cbbca {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _gdbg, _cabb)
		}
	}
	return _fecab
}

// String returns a description of `w`.
func (_fdgd *textWord) String() string {
	return _ca.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _fdgd._addc, _fdgd.PdfRectangle, _fdgd._ggea, _fdgd._eabcc)
}
func (_bcgfb *textTable) markCells() {
	for _dfga := 0; _dfga < _bcgfb._ebab; _dfga++ {
		for _cabfg := 0; _cabfg < _bcgfb._gcead; _cabfg++ {
			_dcdd := _bcgfb.get(_cabfg, _dfga)
			if _dcdd != nil {
				_dcdd._bcgc = true
			}
		}
	}
}
func (_afff *wordBag) allWords() []*textWord {
	var _fdgb []*textWord
	for _, _fcgea := range _afff._dfec {
		_fdgb = append(_fdgb, _fcgea...)
	}
	return _fdgb
}
func (_dabdd compositeCell) parasBBox() (paraList, _ba.PdfRectangle) {
	return _dabdd.paraList, _dabdd.PdfRectangle
}
func (_cbdc *textTable) log(_badc string) {
	if !_eca {
		return
	}
	_f.Log.Info("~\u007e\u007e\u0020\u0025\u0073\u003a \u0025\u0064\u0020\u0078\u0020\u0025d\u0020\u0067\u0072\u0069\u0064\u003d\u0025t\u000a\u0020\u0020\u0020\u0020\u0020\u0020\u0025\u0036\u002e2\u0066", _badc, _cbdc._gcead, _cbdc._ebab, _cbdc._deecf, _cbdc.PdfRectangle)
	for _adeed := 0; _adeed < _cbdc._ebab; _adeed++ {
		for _aageg := 0; _aageg < _cbdc._gcead; _aageg++ {
			_bdfcd := _cbdc.get(_aageg, _adeed)
			if _bdfcd == nil {
				continue
			}
			_ca.Printf("%\u0034\u0064\u0020\u00252d\u003a \u0025\u0036\u002e\u0032\u0066 \u0025\u0071\u0020\u0025\u0064\u000a", _aageg, _adeed, _bdfcd.PdfRectangle, _bbfd(_bdfcd.text(), 50), _ac.RuneCountInString(_bdfcd.text()))
		}
	}
}
func (_cbc *textObject) setTextRenderMode(_bgaf int) {
	if _cbc == nil {
		return
	}
	_cbc._gdcg._adcb = RenderMode(_bgaf)
}
func _cebb(_aeaee _ba.PdfRectangle) *ruling {
	return &ruling{_ecddg: _efcg, _ccbf: _aeaee.Urx, _gcbgg: _aeaee.Lly, _cgeee: _aeaee.Ury}
}
func (_bgga *textTable) depth() float64 {
	_bdaf := 1e10
	for _aeff := 0; _aeff < _bgga._gcead; _aeff++ {
		_bcga := _bgga.get(_aeff, 0)
		if _bcga == nil || _bcga._bgea {
			continue
		}
		_bdaf = _c.Min(_bdaf, _bcga.depth())
	}
	return _bdaf
}

// Text returns the extracted page text.
func (_dega PageText) Text() string { return _dega._gcdc }
func (_caef paraList) llyRange(_aaab []int, _gfag, _agdd float64) []int {
	_fdbc := len(_caef)
	if _agdd < _caef[_aaab[0]].Lly || _gfag > _caef[_aaab[_fdbc-1]].Lly {
		return nil
	}
	_bbfcc := _dc.Search(_fdbc, func(_ccag int) bool { return _caef[_aaab[_ccag]].Lly >= _gfag })
	_effa := _dc.Search(_fdbc, func(_cafab int) bool { return _caef[_aaab[_cafab]].Lly > _agdd })
	return _aaab[_bbfcc:_effa]
}

type textPara struct {
	_ba.PdfRectangle
	_dabb  _ba.PdfRectangle
	_efaf  []*textLine
	_ebcb  *textTable
	_bcgc  bool
	_bgea  bool
	_ceec  *textPara
	_gfeab *textPara
	_ggfg  *textPara
	_gagag *textPara
}

func (_faae paraList) findGridTables(_fedd []gridTiling) []*textTable {
	if _eca {
		_f.Log.Info("\u0066i\u006e\u0064\u0047\u0072\u0069\u0064\u0054\u0061\u0062\u006c\u0065s\u003a\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073", len(_faae))
		for _dddc, _debbb := range _faae {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dddc, _debbb)
		}
	}
	var _egcc []*textTable
	for _dedgg, _adcgb := range _fedd {
		_agcf, _fdggea := _faae.findTableGrid(_adcgb)
		if _agcf != nil {
			_agcf.log(_ca.Sprintf("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0057\u0069\u0074\u0068\u0047\u0072\u0069\u0064\u0073\u003a\u0020%\u0064", _dedgg))
			_egcc = append(_egcc, _agcf)
			_agcf.markCells()
		}
		for _cfcbe := range _fdggea {
			_cfcbe._bcgc = true
		}
	}
	if _eca {
		_f.Log.Info("\u0066i\u006e\u0064\u0047\u0072i\u0064\u0054\u0061\u0062\u006ce\u0073:\u0020%\u0064\u0020\u0074\u0061\u0062\u006c\u0065s", len(_egcc))
	}
	return _egcc
}
func _aeae(_gead string) bool {
	if _ac.RuneCountInString(_gead) < _gcef {
		return false
	}
	_adgf, _ccf := _ac.DecodeLastRuneInString(_gead)
	if _ccf <= 0 || !_ag.Is(_ag.Hyphen, _adgf) {
		return false
	}
	_adgf, _ccf = _ac.DecodeLastRuneInString(_gead[:len(_gead)-_ccf])
	return _ccf > 0 && !_ag.IsSpace(_adgf)
}
func _aafg(_afbfg, _fffg float64) string {
	_abdge := !_bdcfc(_afbfg - _fffg)
	if _abdge {
		return "\u000a"
	}
	return "\u0020"
}
func (_dgbc *textObject) getFont(_fff string) (*_ba.PdfFont, error) {
	if _dgbc._acca._age != nil {
		_agee, _ddb := _dgbc.getFontDict(_fff)
		if _ddb != nil {
			_f.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0067\u0065\u0074\u0046\u006f\u006e\u0074:\u0020n\u0061m\u0065=\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", _fff, _ddb.Error())
			return nil, _ddb
		}
		_dgbc._acca._fg++
		_edgg, _bcd := _dgbc._acca._age[_agee.String()]
		if _bcd {
			_edgg._bbc = _dgbc._acca._fg
			return _edgg._beac, nil
		}
	}
	_bda, _fee := _dgbc.getFontDict(_fff)
	if _fee != nil {
		return nil, _fee
	}
	_cea, _fee := _dgbc.getFontDirect(_fff)
	if _fee != nil {
		return nil, _fee
	}
	if _dgbc._acca._age != nil {
		_gcaf := fontEntry{_cea, _dgbc._acca._fg}
		if len(_dgbc._acca._age) >= _cdde {
			var _ebe []string
			for _cdff := range _dgbc._acca._age {
				_ebe = append(_ebe, _cdff)
			}
			_dc.Slice(_ebe, func(_ebbc, _fcfdc int) bool {
				return _dgbc._acca._age[_ebe[_ebbc]]._bbc < _dgbc._acca._age[_ebe[_fcfdc]]._bbc
			})
			delete(_dgbc._acca._age, _ebe[0])
		}
		_dgbc._acca._age[_bda.String()] = _gcaf
	}
	return _cea, nil
}
func (_bbfg *textTable) compositeColCorridors() map[int][]float64 {
	_ffcbd := make(map[int][]float64, _bbfg._gcead)
	if _eca {
		_f.Log.Info("\u0063\u006f\u006d\u0070o\u0073\u0069\u0074\u0065\u0043\u006f\u006c\u0043\u006f\u0072r\u0069d\u006f\u0072\u0073\u003a\u0020\u0077\u003d%\u0064\u0020", _bbfg._gcead)
	}
	for _abgcgb := 0; _abgcgb < _bbfg._gcead; _abgcgb++ {
		_ffcbd[_abgcgb] = nil
	}
	return _ffcbd
}
func (_gae *subpath) removeDuplicates() {
	if len(_gae._caf) == 0 {
		return
	}
	_aaf := []_bc.Point{_gae._caf[0]}
	for _, _gbfa := range _gae._caf[1:] {
		if !_dadd(_gbfa, _aaf[len(_aaf)-1]) {
			_aaf = append(_aaf, _gbfa)
		}
	}
	_gae._caf = _aaf
}

type textLine struct {
	_ba.PdfRectangle
	_fcaa  float64
	_bdcb  []*textWord
	_edfec float64
}

func (_gabf compositeCell) split(_cabf, _aded []float64) *textTable {
	_eagb := len(_cabf) + 1
	_bdag := len(_aded) + 1
	if _eca {
		_f.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0043\u0065l\u006c\u002e\u0073\u0070l\u0069\u0074\u003a\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a\u0009\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025\u0073\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073=\u0025\u0036\u002e\u0032\u0066\u000a\t\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d%\u0036\u002e\u0032\u0066", _bdag, _eagb, _gabf, _cabf, _aded)
		_ca.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073\u000a", len(_gabf.paraList))
		for _ecggb, _fdff := range _gabf.paraList {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _ecggb, _fdff.String())
		}
		_ca.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", len(_gabf.lines()))
		for _dcgd, _dade := range _gabf.lines() {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dcgd, _dade)
		}
	}
	_cabf = _fbcgb(_cabf, _gabf.Ury, _gabf.Lly)
	_aded = _fbcgb(_aded, _gabf.Llx, _gabf.Urx)
	_dabgf := make(map[uint64]*textPara, _bdag*_eagb)
	_gced := textTable{_gcead: _bdag, _ebab: _eagb, _fecac: _dabgf}
	_efdg := _gabf.paraList
	_dc.Slice(_efdg, func(_cffba, _gbade int) bool {
		_aac, _aab := _efdg[_cffba], _efdg[_gbade]
		_fgcg, _egag := _aac.Lly, _aab.Lly
		if _fgcg != _egag {
			return _fgcg < _egag
		}
		return _aac.Llx < _aab.Llx
	})
	_febe := make(map[uint64]_ba.PdfRectangle, _bdag*_eagb)
	for _fabc, _cfgf := range _cabf[1:] {
		_egae := _cabf[_fabc]
		for _deag, _agafd := range _aded[1:] {
			_dgeg := _aded[_deag]
			_febe[_dfee(_deag, _fabc)] = _ba.PdfRectangle{Llx: _dgeg, Urx: _agafd, Lly: _cfgf, Ury: _egae}
		}
	}
	if _eca {
		_f.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0043\u0065l\u006c\u002e\u0073\u0070\u006c\u0069\u0074\u003a\u0020\u0072e\u0063\u0074\u0073")
		_ca.Printf("\u0020\u0020\u0020\u0020")
		for _gbbag := 0; _gbbag < _bdag; _gbbag++ {
			_ca.Printf("\u0025\u0033\u0030\u0064\u002c\u0020", _gbbag)
		}
		_ca.Println()
		for _geadd := 0; _geadd < _eagb; _geadd++ {
			_ca.Printf("\u0020\u0020\u0025\u0032\u0064\u003a", _geadd)
			for _agad := 0; _agad < _bdag; _agad++ {
				_ca.Printf("\u00256\u002e\u0032\u0066\u002c\u0020", _febe[_dfee(_agad, _geadd)])
			}
			_ca.Println()
		}
	}
	_abgcg := func(_agcbd *textLine) (int, int) {
		for _cceg := 0; _cceg < _eagb; _cceg++ {
			for _affb := 0; _affb < _bdag; _affb++ {
				if _ebag(_febe[_dfee(_affb, _cceg)], _agcbd.PdfRectangle) {
					return _affb, _cceg
				}
			}
		}
		return -1, -1
	}
	_facf := make(map[uint64][]*textLine, _bdag*_eagb)
	for _, _cfbg := range _efdg.lines() {
		_fgfb, _aeedf := _abgcg(_cfbg)
		if _fgfb < 0 {
			continue
		}
		_facf[_dfee(_fgfb, _aeedf)] = append(_facf[_dfee(_fgfb, _aeedf)], _cfbg)
	}
	for _bgeac := 0; _bgeac < len(_cabf)-1; _bgeac++ {
		_cceb := _cabf[_bgeac]
		_gfge := _cabf[_bgeac+1]
		for _bccfc := 0; _bccfc < len(_aded)-1; _bccfc++ {
			_cgdgb := _aded[_bccfc]
			_gecce := _aded[_bccfc+1]
			_fbfb := _ba.PdfRectangle{Llx: _cgdgb, Urx: _gecce, Lly: _gfge, Ury: _cceb}
			_acef := _facf[_dfee(_bccfc, _bgeac)]
			if len(_acef) == 0 {
				continue
			}
			_egeae := _cgab(_fbfb, _acef)
			_gced.put(_bccfc, _bgeac, _egeae)
		}
	}
	return &_gced
}
func (_bbcfb rulingList) splitSec() []rulingList {
	_dc.Slice(_bbcfb, func(_aebgf, _cddae int) bool {
		_abbbf, _cfga := _bbcfb[_aebgf], _bbcfb[_cddae]
		if _abbbf._gcbgg != _cfga._gcbgg {
			return _abbbf._gcbgg < _cfga._gcbgg
		}
		return _abbbf._cgeee < _cfga._cgeee
	})
	_fbcba := make(map[*ruling]struct{}, len(_bbcfb))
	_bcef := func(_gbbae *ruling) rulingList {
		_abae := rulingList{_gbbae}
		_fbcba[_gbbae] = struct{}{}
		for _, _dgfd := range _bbcfb {
			if _, _cdcga := _fbcba[_dgfd]; _cdcga {
				continue
			}
			for _, _cgbdf := range _abae {
				if _dgfd.alignsSec(_cgbdf) {
					_abae = append(_abae, _dgfd)
					_fbcba[_dgfd] = struct{}{}
					break
				}
			}
		}
		return _abae
	}
	_dcef := []rulingList{_bcef(_bbcfb[0])}
	for _, _cebg := range _bbcfb[1:] {
		if _, _ccbbb := _fbcba[_cebg]; _ccbbb {
			continue
		}
		_dcef = append(_dcef, _bcef(_cebg))
	}
	return _dcef
}
func (_egbc *subpath) close() {
	if !_dadd(_egbc._caf[0], _egbc.last()) {
		_egbc.add(_egbc._caf[0])
	}
	_egbc._gcbd = true
	_egbc.removeDuplicates()
}
func (_aafb rulingList) comp(_efbd, _bbag int) bool {
	_fbegd, _aebg := _aafb[_efbd], _aafb[_bbag]
	_gbacg, _dbdgf := _fbegd._ecddg, _aebg._ecddg
	if _gbacg != _dbdgf {
		return _gbacg > _dbdgf
	}
	if _gbacg == _edc {
		return false
	}
	_gaaab := func(_eddg bool) bool {
		if _gbacg == _egd {
			return _eddg
		}
		return !_eddg
	}
	_ffaea, _gfc := _fbegd._ccbf, _aebg._ccbf
	if _ffaea != _gfc {
		return _gaaab(_ffaea > _gfc)
	}
	_ffaea, _gfc = _fbegd._gcbgg, _aebg._gcbgg
	if _ffaea != _gfc {
		return _gaaab(_ffaea < _gfc)
	}
	return _gaaab(_fbegd._cgeee < _aebg._cgeee)
}
func (_fef *stateStack) pop() *textState {
	if _fef.empty() {
		return nil
	}
	_cbeb := *(*_fef)[len(*_fef)-1]
	*_fef = (*_fef)[:len(*_fef)-1]
	return &_cbeb
}
func _ceag(_bfcc *textWord, _gcea float64, _bbcf, _ade rulingList) *wordBag {
	_dedgb := _baa(_bfcc._addc)
	_gfb := []*textWord{_bfcc}
	_eef := wordBag{_dfec: map[int][]*textWord{_dedgb: _gfb}, PdfRectangle: _bfcc.PdfRectangle, _egad: _bfcc._ggea, _efae: _gcea, _edga: _bbcf, _gaef: _ade}
	return &_eef
}

const (
	_ebefa = false
	_badb  = false
	_edea  = false
	_adec  = false
	_dfab  = false
	_gbdd  = false
	_gbgf  = false
	_cacd  = false
	_agce  = false
	_cedf  = _agce && true
	_fggg  = _cedf && false
	_fdfe  = _agce && true
	_eca   = false
	_abab  = _eca && false
	_facg  = _eca && true
	_bfg   = false
	_bdda  = _bfg && false
	_dcdb  = _bfg && false
	_acgbc = _bfg && true
	_eecd  = _bfg && false
	_ggfd  = _bfg && false
)

func _feadg(_ebbfb int, _fccea func(int, int) bool) []int {
	_faeb := make([]int, _ebbfb)
	for _bgdcf := range _faeb {
		_faeb[_bgdcf] = _bgdcf
	}
	_dc.Slice(_faeb, func(_bgcc, _ecbd int) bool { return _fccea(_faeb[_bgcc], _faeb[_ecbd]) })
	return _faeb
}
func (_dfbeb *textPara) fontsize() float64 { return _dfbeb._efaf[0]._edfec }
func (_caag *textLine) toTextMarks(_beecd *int) []TextMark {
	var _bbbd []TextMark
	for _, _agae := range _caag._bdcb {
		if _agae._dfgac {
			_bbbd = _cdae(_bbbd, _beecd, "\u0020")
		}
		_dfae := _agae.toTextMarks(_beecd)
		_bbbd = append(_bbbd, _dfae...)
	}
	return _bbbd
}
func (_dcd *wordBag) depthRange(_gbedf, _bddb int) []int {
	var _edec []int
	for _bggf := range _dcd._dfec {
		if _gbedf <= _bggf && _bggf <= _bddb {
			_edec = append(_edec, _bggf)
		}
	}
	if len(_edec) == 0 {
		return nil
	}
	_dc.Ints(_edec)
	return _edec
}

type lineRuling struct {
	_acgf rulingKind
	_edfd markKind
	_af.Color
	_abeb, _cgae _bc.Point
}

func (_gdfgc gridTile) complete() bool { return _gdfgc.numBorders() == 4 }
func _ggcb(_dgff, _bfdcg int) int {
	if _dgff > _bfdcg {
		return _dgff
	}
	return _bfdcg
}
func (_fgcfa *textPara) toTextMarks(_dfdf *int) []TextMark {
	if _fgcfa._ebcb == nil {
		return _fgcfa.toCellTextMarks(_dfdf)
	}
	var _baea []TextMark
	for _bgbg := 0; _bgbg < _fgcfa._ebcb._ebab; _bgbg++ {
		for _bgbe := 0; _bgbe < _fgcfa._ebcb._gcead; _bgbe++ {
			_gfdec := _fgcfa._ebcb.get(_bgbe, _bgbg)
			if _gfdec == nil {
				_baea = _cdae(_baea, _dfdf, "\u0009")
			} else {
				_gaae := _gfdec.toCellTextMarks(_dfdf)
				_baea = append(_baea, _gaae...)
			}
			_baea = _cdae(_baea, _dfdf, "\u0020")
		}
		if _bgbg < _fgcfa._ebcb._ebab-1 {
			_baea = _cdae(_baea, _dfdf, "\u000a")
		}
	}
	return _baea
}
func (_gged *shapesState) stroke(_ecea *[]pathSection) {
	_cfccc := pathSection{_cag: _gged._eagg, Color: _gged._gfad.getStrokeColor()}
	*_ecea = append(*_ecea, _cfccc)
	if _bfg {
		_ca.Printf("\u0020 \u0020\u0020S\u0054\u0052\u004fK\u0045\u003a\u0020\u0025\u0064\u0020\u0073t\u0072\u006f\u006b\u0065\u0073\u0020s\u0073\u003d\u0025\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d%\u002b\u0076\u0020\u0025\u0036\u002e\u0032\u0066\u000a", len(*_ecea), _gged, _gged._gfad.getStrokeColor(), _cfccc.bbox())
		if _bdda {
			for _bcfa, _aeca := range _gged._eagg {
				_ca.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _bcfa, _aeca)
				if _bcfa == 10 {
					break
				}
			}
		}
	}
}
func (_gbaag *wordBag) arrangeText() *textPara {
	_gbaag.sort()
	if _bgbd {
		_gbaag.removeDuplicates()
	}
	var _gcdd []*textLine
	for _, _gdef := range _gbaag.depthIndexes() {
		for !_gbaag.empty(_gdef) {
			_deec := _gbaag.firstReadingIndex(_gdef)
			_bcbe := _gbaag.firstWord(_deec)
			_bgad := _egc(_gbaag, _deec)
			_gacg := _bcbe._ggea
			_abgf := _bcbe._addc - _ggef*_gacg
			_dcgfd := _bcbe._addc + _ggef*_gacg
			_dfba := _feca * _gacg
			_accga := _cdb * _gacg
		_edbf:
			for {
				var _afaa *textWord
				_dcbd := 0
				for _, _ggfdg := range _gbaag.depthBand(_abgf, _dcgfd) {
					_aage := _gbaag.highestWord(_ggfdg, _abgf, _dcgfd)
					if _aage == nil {
						continue
					}
					_agdg := _aged(_aage, _bgad._bdcb[len(_bgad._bdcb)-1])
					if _agdg < -_accga {
						break _edbf
					}
					if _agdg > _dfba {
						continue
					}
					if _afaa != nil && _dafg(_aage, _afaa) >= 0 {
						continue
					}
					_afaa = _aage
					_dcbd = _ggfdg
				}
				if _afaa == nil {
					break
				}
				_bgad.pullWord(_gbaag, _afaa, _dcbd)
			}
			_bgad.markWordBoundaries()
			_gcdd = append(_gcdd, _bgad)
		}
	}
	if len(_gcdd) == 0 {
		return nil
	}
	_dc.Slice(_gcdd, func(_gddc, _ccdg int) bool { return _agb(_gcdd[_gddc], _gcdd[_ccdg]) < 0 })
	_gcgaf := _cgab(_gbaag.PdfRectangle, _gcdd)
	if _agce {
		_f.Log.Info("\u0061\u0072\u0072an\u0067\u0065\u0054\u0065\u0078\u0074\u0020\u0021\u0021\u0021\u0020\u0070\u0061\u0072\u0061\u003d\u0025\u0073", _gcgaf.String())
		if _cedf {
			for _dcedd, _dafdb := range _gcgaf._efaf {
				_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dcedd, _dafdb.String())
				if _fggg {
					for _eaag, _edad := range _dafdb._bdcb {
						_ca.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _eaag, _edad.String())
						for _gcbc, _eaffa := range _edad._bdedd {
							_ca.Printf("\u00251\u0032\u0064\u003a\u0020\u0025\u0073\n", _gcbc, _eaffa.String())
						}
					}
				}
			}
		}
	}
	return _gcgaf
}
func (_gebd *textLine) text() string {
	var _gaa []string
	for _, _ccad := range _gebd._bdcb {
		if _ccad._dfgac {
			_gaa = append(_gaa, "\u0020")
		}
		_gaa = append(_gaa, _ccad._eabcc)
	}
	return _ad.Join(_gaa, "")
}
func _ccda(_adag []pathSection) rulingList {
	_ccbd(_adag)
	if _bfg {
		_f.Log.Info("\u006da\u006b\u0065\u0046\u0069l\u006c\u0052\u0075\u006c\u0069n\u0067s\u003a \u0025\u0064\u0020\u0066\u0069\u006c\u006cs", len(_adag))
	}
	var _ecbf rulingList
	for _, _ccbb := range _adag {
		for _, _cadb := range _ccbb._cag {
			if !_cadb.isQuadrilateral() {
				if _bfg {
					_f.Log.Error("!\u0069s\u0051\u0075\u0061\u0064\u0072\u0069\u006c\u0061t\u0065\u0072\u0061\u006c: \u0025\u0073", _cadb)
				}
				continue
			}
			if _bbec, _ebea := _cadb.makeRectRuling(_ccbb.Color); _ebea {
				_ecbf = append(_ecbf, _bbec)
			} else {
				if _eecd {
					_f.Log.Error("\u0021\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0073", _cadb)
				}
			}
		}
	}
	if _bfg {
		_f.Log.Info("\u006d\u0061\u006b\u0065Fi\u006c\u006c\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0073", _ecbf.String())
	}
	return _ecbf
}
func (_adb *wordBag) sort() {
	for _, _aagc := range _adb._dfec {
		_dc.Slice(_aagc, func(_ecede, _geccc int) bool { return _dafg(_aagc[_ecede], _aagc[_geccc]) < 0 })
	}
}
func (_bbg paraList) llyOrdering() []int {
	_edaa := make([]int, len(_bbg))
	for _fgfa := range _bbg {
		_edaa[_fgfa] = _fgfa
	}
	_dc.SliceStable(_edaa, func(_dfge, _gdbc int) bool {
		_gege, _abfd := _edaa[_dfge], _edaa[_gdbc]
		return _bbg[_gege].Lly < _bbg[_abfd].Lly
	})
	return _edaa
}
func (_gcbb paraList) merge() *textPara {
	_f.Log.Trace("\u006d\u0065\u0072\u0067\u0065:\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_gcbb))
	if len(_gcbb) == 0 {
		return nil
	}
	_gcbb.sortReadingOrder()
	_adga := _gcbb[0].PdfRectangle
	_agedf := _gcbb[0]._efaf
	for _, _bcdc := range _gcbb[1:] {
		_adga = _cafce(_adga, _bcdc.PdfRectangle)
		_agedf = append(_agedf, _bcdc._efaf...)
	}
	return _cgab(_adga, _agedf)
}

// ApplyArea processes the page text only within the specified area `bbox`.
// Each time ApplyArea is called, it updates the result set in `pt`.
// Can be called multiple times in a row with different bounding boxes.
func (_aeaa *PageText) ApplyArea(bbox _ba.PdfRectangle) {
	_ddee := make([]*textMark, 0, len(_aeaa._faga))
	for _, _ddddf := range _aeaa._faga {
		if _ffdd(_ddddf.bbox(), bbox) {
			_ddee = append(_ddee, _ddddf)
		}
	}
	var _gfd paraList
	_gcec := len(_ddee)
	for _abd := 0; _abd < 360 && _gcec > 0; _abd += 90 {
		_gcba := make([]*textMark, 0, len(_ddee)-_gcec)
		for _, _abde := range _ddee {
			if _abde._bdba == _abd {
				_gcba = append(_gcba, _abde)
			}
		}
		if len(_gcba) > 0 {
			_dbcg := _ccfc(_gcba, _aeaa._gfed, nil, nil)
			_gfd = append(_gfd, _dbcg...)
			_gcec -= len(_gcba)
		}
	}
	_cffa := new(_gb.Buffer)
	_gfd.writeText(_cffa)
	_aeaa._gcdc = _cffa.String()
	_aeaa._cdge = _gfd.toTextMarks()
	_aeaa._defb = _gfd.tables()
}
func (_eabg pathSection) bbox() _ba.PdfRectangle {
	_gbed := _eabg._cag[0]._caf[0]
	_fcb := _ba.PdfRectangle{Llx: _gbed.X, Urx: _gbed.X, Lly: _gbed.Y, Ury: _gbed.Y}
	_cbbf := func(_agaa _bc.Point) {
		if _agaa.X < _fcb.Llx {
			_fcb.Llx = _agaa.X
		} else if _agaa.X > _fcb.Urx {
			_fcb.Urx = _agaa.X
		}
		if _agaa.Y < _fcb.Lly {
			_fcb.Lly = _agaa.Y
		} else if _agaa.Y > _fcb.Ury {
			_fcb.Ury = _agaa.Y
		}
	}
	for _, _dfc := range _eabg._cag[0]._caf[1:] {
		_cbbf(_dfc)
	}
	for _, _dcab := range _eabg._cag[1:] {
		for _, _bfbg := range _dcab._caf {
			_cbbf(_bfbg)
		}
	}
	return _fcb
}
func (_fbgf *textLine) endsInHyphen() bool {
	_egbf := _fbgf._bdcb[len(_fbgf._bdcb)-1]
	_ddea := _egbf._eabcc
	_bfada, _bgee := _ac.DecodeLastRuneInString(_ddea)
	if _bgee <= 0 || !_ag.Is(_ag.Hyphen, _bfada) {
		return false
	}
	if _egbf._dfgac && _aeae(_ddea) {
		return true
	}
	return _aeae(_fbgf.text())
}
func (_bddad paraList) computeEBBoxes() {
	if _ebefa {
		_f.Log.Info("\u0063o\u006dp\u0075\u0074\u0065\u0045\u0042\u0042\u006f\u0078\u0065\u0073\u003a")
	}
	for _, _gfeb := range _bddad {
		_gfeb._dabb = _gfeb.PdfRectangle
	}
	_dcgca := _bddad.yNeighbours(0)
	for _ccfe, _deae := range _bddad {
		_gdfc := _deae._dabb
		_ddgd, _gdfg := -1.0e9, +1.0e9
		for _, _aefec := range _dcgca[_deae] {
			_dcc := _bddad[_aefec]._dabb
			if _dcc.Urx < _gdfc.Llx {
				_ddgd = _c.Max(_ddgd, _dcc.Urx)
			} else if _gdfc.Urx < _dcc.Llx {
				_gdfg = _c.Min(_gdfg, _dcc.Llx)
			}
		}
		for _adfd, _adcg := range _bddad {
			_ffda := _adcg._dabb
			if _ccfe == _adfd || _ffda.Ury > _gdfc.Lly {
				continue
			}
			if _ddgd <= _ffda.Llx && _ffda.Llx < _gdfc.Llx {
				_gdfc.Llx = _ffda.Llx
			} else if _ffda.Urx <= _gdfg && _gdfc.Urx < _ffda.Urx {
				_gdfc.Urx = _ffda.Urx
			}
		}
		if _ebefa {
			_ca.Printf("\u0025\u0034\u0064\u003a %\u0036\u002e\u0032\u0066\u2192\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0071\u000a", _ccfe, _deae._dabb, _gdfc, _bbfd(_deae.text(), 50))
		}
		_deae._dabb = _gdfc
	}
	if _cbgf {
		for _, _gbbcb := range _bddad {
			_gbbcb.PdfRectangle = _gbbcb._dabb
		}
	}
}

// ToTextMark returns the public view of `tm`.
func (_dfbg *textMark) ToTextMark() TextMark {
	return TextMark{Text: _dfbg._abdg, Original: _dfbg._deee, BBox: _dfbg._fdbg, Font: _dfbg._ddbd, FontSize: _dfbg._dfca, FillColor: _dfbg._bbcg, StrokeColor: _dfbg._fbbcf, Orientation: _dfbg._bdba}
}
func _cbcb(_afaae float64) bool { return _c.Abs(_afaae) < _gedc }
func _gegf(_aecc []*textMark, _cbca _ba.PdfRectangle) *textWord {
	_bfag := _aecc[0].PdfRectangle
	_gdafd := _aecc[0]._dfca
	for _, _fgaea := range _aecc[1:] {
		_bfag = _cafce(_bfag, _fgaea.PdfRectangle)
		if _fgaea._dfca > _gdafd {
			_gdafd = _fgaea._dfca
		}
	}
	return &textWord{PdfRectangle: _bfag, _bdedd: _aecc, _addc: _cbca.Ury - _bfag.Lly, _ggea: _gdafd}
}
func (_dbaag *textPara) writeText(_gagee _a.Writer) {
	if _dbaag._ebcb == nil {
		_dbaag.writeCellText(_gagee)
		return
	}
	for _eadd := 0; _eadd < _dbaag._ebcb._ebab; _eadd++ {
		for _dcgeb := 0; _dcgeb < _dbaag._ebcb._gcead; _dcgeb++ {
			_fdgge := _dbaag._ebcb.get(_dcgeb, _eadd)
			if _fdgge == nil {
				_gagee.Write([]byte("\u0009"))
			} else {
				_fdgge.writeCellText(_gagee)
			}
			_gagee.Write([]byte("\u0020"))
		}
		if _eadd < _dbaag._ebcb._ebab-1 {
			_gagee.Write([]byte("\u000a"))
		}
	}
}
func (_eea *textObject) getFillColor() _af.Color {
	return _fafcd(_eea._fdcd.ColorspaceNonStroking, _eea._fdcd.ColorNonStroking)
}
func _cbee(_ecead string) string {
	_gbae := []rune(_ecead)
	return string(_gbae[:len(_gbae)-1])
}
func (_feeb rulingList) primMinMax() (float64, float64) {
	_cdbc, _aaac := _feeb[0]._ccbf, _feeb[0]._ccbf
	for _, _cffg := range _feeb[1:] {
		if _cffg._ccbf < _cdbc {
			_cdbc = _cffg._ccbf
		} else if _cffg._ccbf > _aaac {
			_aaac = _cffg._ccbf
		}
	}
	return _cdbc, _aaac
}
func (_eagd *textTable) bbox() _ba.PdfRectangle { return _eagd.PdfRectangle }
func (_fdgg *wordBag) empty(_decc int) bool     { _, _fffa := _fdgg._dfec[_decc]; return !_fffa }
func (_gaffd rulingList) snapToGroups() rulingList {
	_dcdg, _adfb := _gaffd.vertsHorzs()
	if len(_dcdg) > 0 {
		_dcdg = _dcdg.snapToGroupsDirection()
	}
	if len(_adfb) > 0 {
		_adfb = _adfb.snapToGroupsDirection()
	}
	_cddf := append(_dcdg, _adfb...)
	_cddf.log("\u0073\u006e\u0061p\u0054\u006f\u0047\u0072\u006f\u0075\u0070\u0073")
	return _cddf
}
func (_gcede paraList) addNeighbours() {
	_fddde := func(_adad []int, _gcae *textPara) ([]*textPara, []*textPara) {
		_cbce := make([]*textPara, 0, len(_adad)-1)
		_cbdce := make([]*textPara, 0, len(_adad)-1)
		for _, _gcdde := range _adad {
			_dadec := _gcede[_gcdde]
			if _dadec.Urx <= _gcae.Llx {
				_cbce = append(_cbce, _dadec)
			} else if _dadec.Llx >= _gcae.Urx {
				_cbdce = append(_cbdce, _dadec)
			}
		}
		return _cbce, _cbdce
	}
	_gdfgb := func(_gccd []int, _aega *textPara) ([]*textPara, []*textPara) {
		_efbf := make([]*textPara, 0, len(_gccd)-1)
		_caaggg := make([]*textPara, 0, len(_gccd)-1)
		for _, _fbcd := range _gccd {
			_eafg := _gcede[_fbcd]
			if _eafg.Ury <= _aega.Lly {
				_caaggg = append(_caaggg, _eafg)
			} else if _eafg.Lly >= _aega.Ury {
				_efbf = append(_efbf, _eafg)
			}
		}
		return _efbf, _caaggg
	}
	_dded := _gcede.yNeighbours(_acce)
	for _, _cgebd := range _gcede {
		_ecfe := _dded[_cgebd]
		if len(_ecfe) == 0 {
			continue
		}
		_ggacb, _cccg := _fddde(_ecfe, _cgebd)
		if len(_ggacb) == 0 && len(_cccg) == 0 {
			continue
		}
		if len(_ggacb) > 0 {
			_cbad := _ggacb[0]
			for _, _egeaf := range _ggacb[1:] {
				if _egeaf.Urx >= _cbad.Urx {
					_cbad = _egeaf
				}
			}
			for _, _dgbgf := range _ggacb {
				if _dgbgf != _cbad && _dgbgf.Urx > _cbad.Llx {
					_cbad = nil
					break
				}
			}
			if _cbad != nil && _dabdc(_cgebd.PdfRectangle, _cbad.PdfRectangle) {
				_cgebd._ceec = _cbad
			}
		}
		if len(_cccg) > 0 {
			_dffcg := _cccg[0]
			for _, _gdcgc := range _cccg[1:] {
				if _gdcgc.Llx <= _dffcg.Llx {
					_dffcg = _gdcgc
				}
			}
			for _, _dgcb := range _cccg {
				if _dgcb != _dffcg && _dgcb.Llx < _dffcg.Urx {
					_dffcg = nil
					break
				}
			}
			if _dffcg != nil && _dabdc(_cgebd.PdfRectangle, _dffcg.PdfRectangle) {
				_cgebd._gfeab = _dffcg
			}
		}
	}
	_dded = _gcede.xNeighbours(_adf)
	for _, _bdfe := range _gcede {
		_afeb := _dded[_bdfe]
		if len(_afeb) == 0 {
			continue
		}
		_gfgb, _gfcf := _gdfgb(_afeb, _bdfe)
		if len(_gfgb) == 0 && len(_gfcf) == 0 {
			continue
		}
		if len(_gfcf) > 0 {
			_bgecd := _gfcf[0]
			for _, _aggeb := range _gfcf[1:] {
				if _aggeb.Ury >= _bgecd.Ury {
					_bgecd = _aggeb
				}
			}
			for _, _cdfga := range _gfcf {
				if _cdfga != _bgecd && _cdfga.Ury > _bgecd.Lly {
					_bgecd = nil
					break
				}
			}
			if _bgecd != nil && _ccbc(_bdfe.PdfRectangle, _bgecd.PdfRectangle) {
				_bdfe._gagag = _bgecd
			}
		}
		if len(_gfgb) > 0 {
			_bcbee := _gfgb[0]
			for _, _gcbcc := range _gfgb[1:] {
				if _gcbcc.Lly <= _bcbee.Lly {
					_bcbee = _gcbcc
				}
			}
			for _, _egda := range _gfgb {
				if _egda != _bcbee && _egda.Lly < _bcbee.Ury {
					_bcbee = nil
					break
				}
			}
			if _bcbee != nil && _ccbc(_bdfe.PdfRectangle, _bcbee.PdfRectangle) {
				_bdfe._ggfg = _bcbee
			}
		}
	}
	for _, _caea := range _gcede {
		if _caea._ceec != nil && _caea._ceec._gfeab != _caea {
			_caea._ceec = nil
		}
		if _caea._ggfg != nil && _caea._ggfg._gagag != _caea {
			_caea._ggfg = nil
		}
		if _caea._gfeab != nil && _caea._gfeab._ceec != _caea {
			_caea._gfeab = nil
		}
		if _caea._gagag != nil && _caea._gagag._ggfg != _caea {
			_caea._gagag = nil
		}
	}
}

type imageExtractContext struct {
	_aeb  []ImageMark
	_gf   int
	_afbe int
	_bee  int
	_ecdd map[*_ff.PdfObjectStream]*cachedImage
	_caa  *ImageExtractOptions
}
type rulingKind int

func (_acde *shapesState) quadraticTo(_cbg, _cgaf, _eaf, _acgb float64) {
	if _dfab {
		_f.Log.Info("\u0071\u0075\u0061d\u0072\u0061\u0074\u0069\u0063\u0054\u006f\u003a")
	}
	_acde.addPoint(_eaf, _acgb)
}
func (_cddc *wordBag) highestWord(_beae int, _edf, _bgafa float64) *textWord {
	for _, _dgdd := range _cddc._dfec[_beae] {
		if _edf <= _dgdd._addc && _dgdd._addc <= _bgafa {
			return _dgdd
		}
	}
	return nil
}
func (_dggee rulingList) removeDuplicates() rulingList {
	if len(_dggee) == 0 {
		return nil
	}
	_dggee.sort()
	_edfg := rulingList{_dggee[0]}
	for _, _cfcf := range _dggee[1:] {
		if _cfcf.equals(_edfg[len(_edfg)-1]) {
			continue
		}
		_edfg = append(_edfg, _cfcf)
	}
	return _edfg
}
func (_baebd paraList) log(_fgea string) {
	if !_cacd {
		return
	}
	_f.Log.Info("%\u0038\u0073\u003a\u0020\u0025\u0064 \u0070\u0061\u0072\u0061\u0073\u0020=\u003d\u003d\u003d\u003d\u003d\u003d\u002d-\u002d\u002d\u002d\u002d\u002d\u003d\u003d\u003d\u003d\u003d=\u003d", _fgea, len(_baebd))
	for _gbbf, _ggge := range _baebd {
		if _ggge == nil {
			continue
		}
		_ccbga := _ggge.text()
		_bbcd := "\u0020\u0020"
		if _ggge._ebcb != nil {
			_bbcd = _ca.Sprintf("\u005b%\u0064\u0078\u0025\u0064\u005d", _ggge._ebcb._gcead, _ggge._ebcb._ebab)
		}
		_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0025s\u0020\u0025\u0071\u000a", _gbbf, _ggge.PdfRectangle, _bbcd, _bbfd(_ccbga, 50))
	}
}

// ExtractPageText returns the text contents of `e` (an Extractor for a page) as a PageText.
// TODO(peterwilliams97): The stats complicate this function signature and aren't very useful.
//                        Replace with a function like Extract() (*PageText, error)
func (_fge *Extractor) ExtractPageText() (*PageText, int, int, error) {
	_bfd, _cadd, _fgda, _acf := _fge.extractPageText(_fge._ge, _fge._fbe, _bc.IdentityMatrix(), 0)
	if _acf != nil && _acf != _ba.ErrColorOutOfRange {
		return nil, 0, 0, _acf
	}
	_bfd.computeViews()
	_acf = _badef(_bfd)
	if _acf != nil {
		return nil, 0, 0, _acf
	}
	return _bfd, _cadd, _fgda, nil
}
func (_cbef *textObject) moveLP(_gbac, _gbdac float64) {
	_cbef._bdc.Concat(_bc.NewMatrix(1, 0, 0, 1, _gbac, _gbdac))
	_cbef._ddfg = _cbef._bdc
}

type bounded interface{ bbox() _ba.PdfRectangle }

const (
	RenderModeStroke RenderMode = 1 << iota
	RenderModeFill
	RenderModeClip
)

func (_dfe *textObject) moveTextSetLeading(_cdcg, _add float64) {
	_dfe._gdcg._dgbbf = -_add
	_dfe.moveLP(_cdcg, _add)
}
func (_eeab gridTiling) log(_eegeg string) {
	if !_acgbc {
		return
	}
	_f.Log.Info("\u0074i\u006ci\u006e\u0067\u003a\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0071", len(_eeab._eade), len(_eeab._bcgfa), _eegeg)
	_ca.Printf("\u0020\u0020\u0020l\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a", _eeab._eade)
	_ca.Printf("\u0020\u0020\u0020l\u006c\u0079\u003d\u0025\u002e\u0032\u0066\u000a", _eeab._bcgfa)
	for _deba, _bagf := range _eeab._bcgfa {
		_abggc, _bcbfe := _eeab._bbca[_bagf]
		if !_bcbfe {
			continue
		}
		_ca.Printf("%\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _deba, _bagf)
		for _gdgda, _eccge := range _eeab._eade {
			_cbff, _gbfac := _abggc[_eccge]
			if !_gbfac {
				continue
			}
			_ca.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _gdgda, _cbff.String())
		}
	}
}
func (_badae *textTable) getRight() paraList {
	_fege := make(paraList, _badae._ebab)
	for _deed := 0; _deed < _badae._ebab; _deed++ {
		_cdfec := _badae.get(_badae._gcead-1, _deed)._gfeab
		if _cdfec.taken() {
			return nil
		}
		_fege[_deed] = _cdfec
	}
	for _fccc := 0; _fccc < _badae._ebab-1; _fccc++ {
		if _fege[_fccc]._gagag != _fege[_fccc+1] {
			return nil
		}
	}
	return _fege
}
func (_eggf *subpath) isQuadrilateral() bool {
	if len(_eggf._caf) < 4 || len(_eggf._caf) > 5 {
		return false
	}
	if len(_eggf._caf) == 5 {
		_ccfgb := _eggf._caf[0]
		_bege := _eggf._caf[4]
		if _ccfgb.X != _bege.X || _ccfgb.Y != _bege.Y {
			return false
		}
	}
	return true
}

// PageText represents the layout of text on a device page.
type PageText struct {
	_faga []*textMark
	_gcdc string
	_cdge []TextMark
	_defb []TextTable
	_gfed _ba.PdfRectangle
	_cfde []pathSection
	_fdd  []pathSection
}

func _fbee(_afcbf []compositeCell) []float64 {
	var _fbdf []*textLine
	_aaed := 0
	for _, _cgfe := range _afcbf {
		_aaed += len(_cgfe.paraList)
		_fbdf = append(_fbdf, _cgfe.lines()...)
	}
	_dc.Slice(_fbdf, func(_gfdc, _gceeg int) bool {
		_egdd, _dgddd := _fbdf[_gfdc], _fbdf[_gceeg]
		_fddd, _fefed := _egdd._fcaa, _dgddd._fcaa
		if !_bdcfc(_fddd - _fefed) {
			return _fddd < _fefed
		}
		return _egdd.Llx < _dgddd.Llx
	})
	if _eca {
		_ca.Printf("\u0020\u0020\u0020 r\u006f\u0077\u0042\u006f\u0072\u0064\u0065\u0072\u0073:\u0020%\u0064 \u0070a\u0072\u0061\u0073\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", _aaed, len(_fbdf))
		for _dcbea, _egec := range _fbdf {
			_ca.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _dcbea, _egec)
		}
	}
	var _dffa []float64
	_edag := _fbdf[0]
	var _cgba [][]*textLine
	_eddc := []*textLine{_edag}
	for _cefdd, _abfgg := range _fbdf[1:] {
		if _abfgg.Ury < _edag.Lly {
			_bbaa := 0.5 * (_abfgg.Ury + _edag.Lly)
			if _eca {
				_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u003c\u0020\u0025\u0036.\u0032f\u0020\u0062\u006f\u0072\u0064\u0065\u0072\u003d\u0025\u0036\u002e\u0032\u0066\u000a"+"\u0009\u0020\u0071\u003d\u0025\u0073\u000a\u0009\u0020p\u003d\u0025\u0073\u000a", _cefdd, _abfgg.Ury, _edag.Lly, _bbaa, _edag, _abfgg)
			}
			_dffa = append(_dffa, _bbaa)
			_cgba = append(_cgba, _eddc)
			_eddc = nil
		}
		_eddc = append(_eddc, _abfgg)
		if _abfgg.Lly < _edag.Lly {
			_edag = _abfgg
		}
	}
	if len(_eddc) > 0 {
		_cgba = append(_cgba, _eddc)
	}
	if _eca {
		_ca.Printf(" \u0020\u0020\u0020\u0020\u0020\u0020 \u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073\u003d%\u0036.\u0032\u0066\u000a", _dffa)
	}
	if _eca {
		_f.Log.Info("\u0072\u006f\u0077\u003d\u0025\u0064", len(_afcbf))
		for _cece, _ffbaa := range _afcbf {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cece, _ffbaa)
		}
		_f.Log.Info("\u0067r\u006f\u0075\u0070\u0073\u003d\u0025d", len(_cgba))
		for _fdcbc, _eddb := range _cgba {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0064\u000a", _fdcbc, len(_eddb))
			for _eegg, _ggdb := range _eddb {
				_ca.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _eegg, _ggdb)
			}
		}
	}
	_abcfg := true
	for _eeae, _gdec := range _cgba {
		_fcac := true
		for _cbgc, _dfff := range _afcbf {
			if _eca {
				_ca.Printf("\u0020\u0020\u0020\u007e\u007e\u007e\u0067\u0072\u006f\u0075\u0070\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0063\u0065\u006cl\u0020\u0025\u0064\u0020\u006ff\u0020\u0025d\u0020\u0025\u0073\u000a", _eeae, len(_cgba), _cbgc, len(_afcbf), _dfff)
			}
			if !_dfff.hasLines(_gdec) {
				if _eca {
					_ca.Printf("\u0020\u0020\u0020\u0021\u0021\u0021\u0067\u0072\u006f\u0075\u0070\u0020\u0025d\u0020\u006f\u0066\u0020\u0025\u0064 \u0063\u0065\u006c\u006c\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064 \u004f\u0055\u0054\u000a", _eeae, len(_cgba), _cbgc, len(_afcbf))
				}
				_fcac = false
				break
			}
		}
		if !_fcac {
			_abcfg = false
			break
		}
	}
	if !_abcfg {
		if _eca {
			_f.Log.Info("\u0072\u006f\u0077\u0020\u0063o\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u0020\u0064\u006f\u006e\u0027\u0074 \u0073\u0070\u0061\u006e\u0020\u0061\u006c\u006c\u0020\u0063\u0065\u006c\u006c\u0073\u0020\u0069\u006e\u0020\u0072\u006f\u0077\u002e\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
		}
		_dffa = nil
	}
	if _eca && _dffa != nil {
		_ca.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u002a\u002a*\u0072\u006f\u0077\u0043\u006f\u0072\u0072i\u0064\u006f\u0072\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u000a", _dffa)
	}
	return _dffa
}
func _ccbc(_gcfa, _fefe _ba.PdfRectangle) bool {
	return _fefe.Llx <= _gcfa.Urx && _gcfa.Llx <= _fefe.Urx
}

const _cdgg = 20

func (_addf *textObject) getStrokeColor() _af.Color {
	return _fafcd(_addf._fdcd.ColorspaceStroking, _addf._fdcd.ColorStroking)
}
func (_bgdc paraList) extractTables(_baee []gridTiling) paraList {
	if _eca {
		_f.Log.Debug("\u0065\u0078\u0074r\u0061\u0063\u0074\u0054\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_bgdc))
	}
	if len(_bgdc) < _cce {
		return _bgdc
	}
	_fcff := _bgdc.findTables(_baee)
	if _eca {
		_f.Log.Info("c\u006f\u006d\u0062\u0069\u006e\u0065d\u0020\u0074\u0061\u0062\u006c\u0065s\u0020\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_fcff))
		for _dcace, _dgaf := range _fcff {
			_dgaf.log(_ca.Sprintf("c\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u0020\u0025\u0064", _dcace))
		}
	}
	return _bgdc.applyTables(_fcff)
}
func _bdaa(_bdfc, _aeee _ba.PdfRectangle) (_ba.PdfRectangle, bool) {
	if !_ffdd(_bdfc, _aeee) {
		return _ba.PdfRectangle{}, false
	}
	return _ba.PdfRectangle{Llx: _c.Max(_bdfc.Llx, _aeee.Llx), Urx: _c.Min(_bdfc.Urx, _aeee.Urx), Lly: _c.Max(_bdfc.Lly, _aeee.Lly), Ury: _c.Min(_bdfc.Ury, _aeee.Ury)}, true
}

// BBox returns the smallest axis-aligned rectangle that encloses all the TextMarks in `ma`.
func (_bfdc *TextMarkArray) BBox() (_ba.PdfRectangle, bool) {
	var _ebd _ba.PdfRectangle
	_ggg := false
	for _, _ecedf := range _bfdc._afdf {
		if _ecedf.Meta || _cfba(_ecedf.Text) {
			continue
		}
		if _ggg {
			_ebd = _cafce(_ebd, _ecedf.BBox)
		} else {
			_ebd = _ecedf.BBox
			_ggg = true
		}
	}
	return _ebd, _ggg
}
func _cdae(_bcaf []TextMark, _ffcb *int, _fefba string) []TextMark {
	_bcaec := _ffba
	_bcaec.Text = _fefba
	return _bddd(_bcaf, _ffcb, _bcaec)
}
func (_cbf *textLine) pullWord(_agcb *wordBag, _gbgc *textWord, _ebge int) {
	_cbf.appendWord(_gbgc)
	_agcb.removeWord(_gbgc, _ebge)
}
func (_acbb *shapesState) addPoint(_dggb, _fcgc float64) {
	_eada := _acbb.establishSubpath()
	_ddg := _acbb.devicePoint(_dggb, _fcgc)
	if _eada == nil {
		_acbb._debd = true
		_acbb._ebeg = _ddg
	} else {
		_eada.add(_ddg)
	}
}
func (_dgca rulingList) secMinMax() (float64, float64) {
	_bgba, _feabe := _dgca[0]._gcbgg, _dgca[0]._cgeee
	for _, _bdcf := range _dgca[1:] {
		if _bdcf._gcbgg < _bgba {
			_bgba = _bdcf._gcbgg
		}
		if _bdcf._cgeee > _feabe {
			_feabe = _bdcf._cgeee
		}
	}
	return _bgba, _feabe
}
func (_abea *textPara) toCellTextMarks(_ebgef *int) []TextMark {
	var _ecc []TextMark
	for _ffedg, _eccg := range _abea._efaf {
		_efde := _eccg.toTextMarks(_ebgef)
		_egbcb := _dedbg && _eccg.endsInHyphen() && _ffedg != len(_abea._efaf)-1
		if _egbcb {
			_efde = _fecgc(_efde, _ebgef)
		}
		_ecc = append(_ecc, _efde...)
		if !(_egbcb || _ffedg == len(_abea._efaf)-1) {
			_ecc = _cdae(_ecc, _ebgef, _aafg(_eccg._fcaa, _abea._efaf[_ffedg+1]._fcaa))
		}
	}
	return _ecc
}

var (
	_bbfcf = map[rune]string{0x0060: "\u0300", 0x02CB: "\u0300", 0x0027: "\u0301", 0x00B4: "\u0301", 0x02B9: "\u0301", 0x02CA: "\u0301", 0x005E: "\u0302", 0x02C6: "\u0302", 0x007E: "\u0303", 0x02DC: "\u0303", 0x00AF: "\u0304", 0x02C9: "\u0304", 0x02D8: "\u0306", 0x02D9: "\u0307", 0x00A8: "\u0308", 0x00B0: "\u030a", 0x02DA: "\u030a", 0x02BA: "\u030b", 0x02DD: "\u030b", 0x02C7: "\u030c", 0x02C8: "\u030d", 0x0022: "\u030e", 0x02BB: "\u0312", 0x02BC: "\u0313", 0x0486: "\u0313", 0x055A: "\u0313", 0x02BD: "\u0314", 0x0485: "\u0314", 0x0559: "\u0314", 0x02D4: "\u031d", 0x02D5: "\u031e", 0x02D6: "\u031f", 0x02D7: "\u0320", 0x02B2: "\u0321", 0x00B8: "\u0327", 0x02CC: "\u0329", 0x02B7: "\u032b", 0x02CD: "\u0331", 0x005F: "\u0332", 0x204E: "\u0359"}
)

func (_gcfc *shapesState) lastpointEstablished() (_bc.Point, bool) {
	if _gcfc._debd {
		return _gcfc._ebeg, false
	}
	_fgdb := len(_gcfc._eagg)
	if _fgdb > 0 && _gcfc._eagg[_fgdb-1]._gcbd {
		return _gcfc._eagg[_fgdb-1].last(), false
	}
	return _bc.Point{}, true
}
func (_dad *textObject) getFontDict(_cfcc string) (_cfdge _ff.PdfObject, _ccaa error) {
	_fefb := _dad._dgda
	if _fefb == nil {
		_f.Log.Debug("g\u0065\u0074\u0046\u006f\u006e\u0074D\u0069\u0063\u0074\u002e\u0020\u004eo\u0020\u0072\u0065\u0073\u006f\u0075\u0072c\u0065\u0073\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025#\u0071", _cfcc)
		return nil, nil
	}
	_cfdge, _dcbf := _fefb.GetFontByName(_ff.PdfObjectName(_cfcc))
	if !_dcbf {
		_f.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0067\u0065t\u0046\u006f\u006et\u0044\u0069\u0063\u0074\u003a\u0020\u0046\u006f\u006et \u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020\u006e\u0061m\u0065\u003d%\u0023\u0071", _cfcc)
		return nil, _g.New("f\u006f\u006e\u0074\u0020no\u0074 \u0069\u006e\u0020\u0072\u0065s\u006f\u0075\u0072\u0063\u0065\u0073")
	}
	return _cfdge, nil
}
func _baa(_afda float64) int {
	var _ddgb int
	if _afda >= 0 {
		_ddgb = int(_afda / _acag)
	} else {
		_ddgb = int(_afda/_acag) - 1
	}
	return _ddgb
}
func _fcfd(_gfea _bc.Matrix) _bc.Point {
	_fdad, _fadf := _gfea.Translation()
	return _bc.Point{X: _fdad, Y: _fadf}
}
func (_gdd *textObject) setTextRise(_dbgd float64) {
	if _gdd == nil {
		return
	}
	_gdd._gdcg._aea = _dbgd
}
func (_eae *textObject) setFont(_bgce string, _cbde float64) error {
	if _eae == nil {
		return nil
	}
	_eae._gdcg._bgg = _cbde
	_ebga, _gbf := _eae.getFont(_bgce)
	if _gbf != nil {
		return _gbf
	}
	_eae._gdcg._fdf = _ebga
	return nil
}
func (_gafbg *textTable) logComposite(_bfef string) {
	if !_eca {
		return
	}
	_f.Log.Info("\u007e~\u007eP\u0061\u0072\u0061\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _gafbg._gcead, _gafbg._ebab, _bfef)
	_ca.Printf("\u0025\u0035\u0073 \u007c", "")
	for _fafa := 0; _fafa < _gafbg._gcead; _fafa++ {
		_ca.Printf("\u0025\u0033\u0064 \u007c", _fafa)
	}
	_ca.Println("")
	_ca.Printf("\u0025\u0035\u0073 \u002b", "")
	for _dbbfg := 0; _dbbfg < _gafbg._gcead; _dbbfg++ {
		_ca.Printf("\u0025\u0033\u0073 \u002b", "\u002d\u002d\u002d")
	}
	_ca.Println("")
	for _cabc := 0; _cabc < _gafbg._ebab; _cabc++ {
		_ca.Printf("\u0025\u0035\u0064 \u007c", _cabc)
		for _agcc := 0; _agcc < _gafbg._gcead; _agcc++ {
			_cage, _ := _gafbg._gabfc[_dfee(_agcc, _cabc)].parasBBox()
			_ca.Printf("\u0025\u0033\u0064 \u007c", len(_cage))
		}
		_ca.Println("")
	}
	_f.Log.Info("\u007e~\u007eT\u0065\u0078\u0074\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _gafbg._gcead, _gafbg._ebab, _bfef)
	_ca.Printf("\u0025\u0035\u0073 \u007c", "")
	for _gbedb := 0; _gbedb < _gafbg._gcead; _gbedb++ {
		_ca.Printf("\u0025\u0031\u0032\u0064\u0020\u007c", _gbedb)
	}
	_ca.Println("")
	_ca.Printf("\u0025\u0035\u0073 \u002b", "")
	for _fgbef := 0; _fgbef < _gafbg._gcead; _fgbef++ {
		_ca.Print("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d\u002b")
	}
	_ca.Println("")
	for _fgedf := 0; _fgedf < _gafbg._ebab; _fgedf++ {
		_ca.Printf("\u0025\u0035\u0064 \u007c", _fgedf)
		for _dbfb := 0; _dbfb < _gafbg._gcead; _dbfb++ {
			_beedb, _ := _gafbg._gabfc[_dfee(_dbfb, _fgedf)].parasBBox()
			_dcebe := ""
			_dbdgc := _beedb.merge()
			if _dbdgc != nil {
				_dcebe = _dbdgc.text()
			}
			_dcebe = _ca.Sprintf("\u0025\u0071", _bbfd(_dcebe, 12))
			_dcebe = _dcebe[1 : len(_dcebe)-1]
			_ca.Printf("\u0025\u0031\u0032\u0073\u0020\u007c", _dcebe)
		}
		_ca.Println("")
	}
}
func (_cdg *imageExtractContext) extractInlineImage(_cad *_da.ContentStreamInlineImage, _cgf _da.GraphicsState, _dee *_ba.PdfPageResources) error {
	_cb, _bgf := _cad.ToImage(_dee)
	if _bgf != nil {
		return _bgf
	}
	_ddd, _bgf := _cad.GetColorSpace(_dee)
	if _bgf != nil {
		return _bgf
	}
	if _ddd == nil {
		_ddd = _ba.NewPdfColorspaceDeviceGray()
	}
	_agge, _bgf := _ddd.ImageToRGB(*_cb)
	if _bgf != nil {
		return _bgf
	}
	_cf := ImageMark{Image: &_agge, Width: _cgf.CTM.ScalingFactorX(), Height: _cgf.CTM.ScalingFactorY(), Angle: _cgf.CTM.Angle()}
	_cf.X, _cf.Y = _cgf.CTM.Translation()
	_cdg._aeb = append(_cdg._aeb, _cf)
	_cdg._gf++
	return nil
}

const (
	_bcfg  = 1.0e-6
	_ebf   = 1.0e-4
	_aagd  = 10
	_acag  = 6
	_ggef  = 0.5
	_dadg  = 0.12
	_bfge  = 0.19
	_effe  = 0.04
	_cdggc = 0.04
	_cafa  = 1.0
	_bge   = 0.04
	_dfgd  = 0.4
	_bgcef = 0.7
	_cdgee = 1.0
	_caed  = 0.1
	_feca  = 1.4
	_cdb   = 0.46
	_fecg  = 0.02
	_caeb  = 0.2
	_egfd  = 0.5
	_gcef  = 4
	_bbcfg = 4.0
	_cce   = 6
	_bdfg  = 0.3
	_adf   = 0.01
	_acce  = 0.02
	_efb   = 2
	_fffd  = 2
	_agdc  = 500
	_deea  = 4.0
	_gdddc = 4.0
	_dae   = 0.05
	_cbbb  = 0.1
	_gdae  = 2.0
	_gedc  = 2.0
	_effd  = 1.5
	_addag = 3.0
	_cgec  = 0.25
)

func (_afbf *wordBag) maxDepth() float64 { return _afbf._efae - _afbf.Lly }
func (_fbfbg *textWord) absorb(_agdae *textWord) {
	_fbfbg.PdfRectangle = _cafce(_fbfbg.PdfRectangle, _agdae.PdfRectangle)
	_fbfbg._bdedd = append(_fbfbg._bdedd, _agdae._bdedd...)
}
func (_dde *textObject) reset() {
	_dde._ddfg = _bc.IdentityMatrix()
	_dde._bdc = _bc.IdentityMatrix()
	_dde._adda = nil
}
func _cgab(_afea _ba.PdfRectangle, _fecd []*textLine) *textPara {
	return &textPara{PdfRectangle: _afea, _efaf: _fecd}
}
func (_bad *shapesState) drawRectangle(_cfg, _fgcf, _gddd, _eeea float64) {
	if _dfab {
		_egg := _bad.devicePoint(_cfg, _fgcf)
		_gbdg := _bad.devicePoint(_cfg+_gddd, _fgcf+_eeea)
		_bced := _ba.PdfRectangle{Llx: _egg.X, Lly: _egg.Y, Urx: _gbdg.X, Ury: _gbdg.Y}
		_f.Log.Info("d\u0072a\u0077\u0052\u0065\u0063\u0074\u0061\u006e\u0067l\u0065\u003a\u0020\u00256.\u0032\u0066", _bced)
	}
	_bad.newSubPath()
	_bad.moveTo(_cfg, _fgcf)
	_bad.lineTo(_cfg+_gddd, _fgcf)
	_bad.lineTo(_cfg+_gddd, _fgcf+_eeea)
	_bad.lineTo(_cfg, _fgcf+_eeea)
	_bad.closePath()
}
func (_feaa *stateStack) top() *textState {
	if _feaa.empty() {
		return nil
	}
	return (*_feaa)[_feaa.size()-1]
}
func _dbdb(_ggaf []*wordBag) []*wordBag {
	if len(_ggaf) <= 1 {
		return _ggaf
	}
	if _agce {
		_f.Log.Info("\u006d\u0065\u0072\u0067\u0065\u0057\u006f\u0072\u0064B\u0061\u0067\u0073\u003a")
	}
	_dc.Slice(_ggaf, func(_bcde, _fbec int) bool {
		_bcega, _dcbe := _ggaf[_bcde], _ggaf[_fbec]
		_ggbf := _bcega.Width() * _bcega.Height()
		_fafe := _dcbe.Width() * _dcbe.Height()
		if _ggbf != _fafe {
			return _ggbf > _fafe
		}
		if _bcega.Height() != _dcbe.Height() {
			return _bcega.Height() > _dcbe.Height()
		}
		return _bcde < _fbec
	})
	var _fbff []*wordBag
	_bdad := make(intSet)
	for _dead := 0; _dead < len(_ggaf); _dead++ {
		if _bdad.has(_dead) {
			continue
		}
		_beaf := _ggaf[_dead]
		for _edbe := _dead + 1; _edbe < len(_ggaf); _edbe++ {
			if _bdad.has(_dead) {
				continue
			}
			_fdebb := _ggaf[_edbe]
			_debeg := _beaf.PdfRectangle
			_debeg.Llx -= _beaf._egad
			if _ebag(_debeg, _fdebb.PdfRectangle) {
				_beaf.absorb(_fdebb)
				_bdad.add(_edbe)
			}
		}
		_fbff = append(_fbff, _beaf)
	}
	if len(_ggaf) != len(_fbff)+len(_bdad) {
		_f.Log.Error("\u006d\u0065\u0072ge\u0057\u006f\u0072\u0064\u0042\u0061\u0067\u0073\u003a \u0025d\u2192%\u0064 \u0061\u0062\u0073\u006f\u0072\u0062\u0065\u0064\u003d\u0025\u0064", len(_ggaf), len(_fbff), len(_bdad))
	}
	return _fbff
}
func (_ggefg rulingList) findPrimSec(_dddgd, _cdebg float64) *ruling {
	for _, _abbe := range _ggefg {
		if _bdcfc(_abbe._ccbf-_dddgd) && _abbe._gcbgg-_gdae <= _cdebg && _cdebg <= _abbe._cgeee+_gdae {
			return _abbe
		}
	}
	return nil
}

// ImageMark represents an image drawn on a page and its position in device coordinates.
// All coordinates are in device coordinates.
type ImageMark struct {
	Image *_ba.Image

	// Dimensions of the image as displayed in the PDF.
	Width  float64
	Height float64

	// Position of the image in PDF coordinates (lower left corner).
	X float64
	Y float64

	// Angle in degrees, if rotated.
	Angle float64
}

func _adcf(_ebbf map[int][]float64) string {
	_ffec := _dceb(_ebbf)
	_adfa := make([]string, len(_ebbf))
	for _abbf, _bgefa := range _ffec {
		_adfa[_abbf] = _ca.Sprintf("\u0025\u0064\u003a\u0020\u0025\u002e\u0032\u0066", _bgefa, _ebbf[_bgefa])
	}
	return _ca.Sprintf("\u007b\u0025\u0073\u007d", _ad.Join(_adfa, "\u002c\u0020"))
}
func _ffdd(_fdgbg, _ddga _ba.PdfRectangle) bool { return _ccbc(_fdgbg, _ddga) && _dabdc(_fdgbg, _ddga) }
func _bfbd(_cdgc *_da.ContentStreamOperation) (float64, error) {
	if len(_cdgc.Params) != 1 {
		_cff := _g.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
		_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _cdgc.Operand, 1, len(_cdgc.Params), _cdgc.Params)
		return 0.0, _cff
	}
	return _ff.GetNumberAsFloat(_cdgc.Params[0])
}
func (_fdded gridTiling) complete() bool {
	for _, _cgefa := range _fdded._bbca {
		for _, _gbdace := range _cgefa {
			if !_gbdace.complete() {
				return false
			}
		}
	}
	return true
}
func _ccea(_ffdb float64, _adeg int) int {
	if _adeg == 0 {
		_adeg = 1
	}
	_dceg := float64(_adeg)
	return int(_c.Round(_ffdb/_dceg) * _dceg)
}

// ImageExtractOptions contains options for controlling image extraction from
// PDF pages.
type ImageExtractOptions struct{ IncludeInlineStencilMasks bool }

func (_ggebe *wordBag) firstReadingIndex(_faec int) int {
	_cdec := _ggebe.firstWord(_faec)._ggea
	_caga := float64(_faec+1) * _acag
	_begd := _caga + _bbcfg*_cdec
	_ebef := _faec
	for _, _abc := range _ggebe.depthBand(_caga, _begd) {
		if _dafg(_ggebe.firstWord(_abc), _ggebe.firstWord(_ebef)) < 0 {
			_ebef = _abc
		}
	}
	return _ebef
}

// Elements returns the TextMarks in `ma`.
func (_bec *TextMarkArray) Elements() []TextMark { return _bec._afdf }
func _cdfg(_cfda, _aefa bounded) float64         { return _bfddc(_cfda) - _bfddc(_aefa) }

// String returns a human readable description of `path`.
func (_bcfb *subpath) String() string {
	_feeg := _bcfb._caf
	_gde := len(_feeg)
	if _gde <= 5 {
		return _ca.Sprintf("\u0025d\u003a\u0020\u0025\u0036\u002e\u0032f", _gde, _feeg)
	}
	return _ca.Sprintf("\u0025d\u003a\u0020\u0025\u0036.\u0032\u0066\u0020\u0025\u0036.\u0032f\u0020.\u002e\u002e\u0020\u0025\u0036\u002e\u0032f", _gde, _feeg[0], _feeg[1], _feeg[_gde-1])
}
func _adcd(_cdag []*textWord, _ggeb float64, _dfbe, _cgdc rulingList) *wordBag {
	_cbgg := _ceag(_cdag[0], _ggeb, _dfbe, _cgdc)
	for _, _afed := range _cdag[1:] {
		_fece := _baa(_afed._addc)
		_cbgg._dfec[_fece] = append(_cbgg._dfec[_fece], _afed)
		_cbgg.PdfRectangle = _cafce(_cbgg.PdfRectangle, _afed.PdfRectangle)
	}
	_cbgg.sort()
	return _cbgg
}

// String returns a description of `p`.
func (_efff *textPara) String() string {
	if _efff._bgea {
		return _ca.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u005b\u0045\u004d\u0050\u0054\u0059\u005d", _efff.PdfRectangle)
	}
	_bef := ""
	if _efff._ebcb != nil {
		_bef = _ca.Sprintf("\u005b\u0025\u0064\u0078\u0025\u0064\u005d\u0020", _efff._ebcb._gcead, _efff._ebcb._ebab)
	}
	return _ca.Sprintf("\u0025\u0036\u002e\u0032f \u0025\u0073\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u0020\u0025\u0071", _efff.PdfRectangle, _bef, len(_efff._efaf), _bbfd(_efff.text(), 50))
}
func (_dfegb rulingList) sortStrict() {
	_dc.Slice(_dfegb, func(_cgce, _gbgb int) bool {
		_bcgd, _gdaee := _dfegb[_cgce], _dfegb[_gbgb]
		_ccec, _baca := _bcgd._ecddg, _gdaee._ecddg
		if _ccec != _baca {
			return _ccec > _baca
		}
		_eefe, _fcdg := _bcgd._ccbf, _gdaee._ccbf
		if !_bdcfc(_eefe - _fcdg) {
			return _eefe < _fcdg
		}
		_eefe, _fcdg = _bcgd._gcbgg, _gdaee._gcbgg
		if _eefe != _fcdg {
			return _eefe < _fcdg
		}
		return _bcgd._cgeee < _gdaee._cgeee
	})
}

// String returns a description of `state`.
func (_fgde *textState) String() string {
	_ecg := "\u005bN\u004f\u0054\u0020\u0053\u0045\u0054]"
	if _fgde._fdf != nil {
		_ecg = _fgde._fdf.BaseFont()
	}
	return _ca.Sprintf("\u0074\u0063\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0077\u003d\u0025\u002e\u0032\u0066 \u0074f\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071", _fgde._aed, _fgde._cef, _fgde._bgg, _ecg)
}

var _ffba = TextMark{Text: "\u005b\u0058\u005d", Original: "\u0020", Meta: true, FillColor: _af.White, StrokeColor: _af.White}

const (
	_bfa = "\u0045\u0052R\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
	_bca = "\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043a\u006e\u0027\u0074 g\u0065\u0074\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002c\u0020\u0066\u006fn\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064"
	_fba = "\u0045\u0052\u0052O\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
)

func (_bedc *wordBag) depthBand(_dge, _agdf float64) []int {
	if len(_bedc._dfec) == 0 {
		return nil
	}
	return _bedc.depthRange(_bedc.getDepthIdx(_dge), _bedc.getDepthIdx(_agdf))
}
func _dcfba(_fcgg string, _gabe []rulingList) {
	_f.Log.Info("\u0024\u0024 \u0025\u0064\u0020g\u0072\u0069\u0064\u0073\u0020\u002d\u0020\u0025\u0073", len(_gabe), _fcgg)
	for _cdad, _gceb := range _gabe {
		_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cdad, _gceb.String())
	}
}

// New returns an Extractor instance for extracting content from the input PDF page.
func New(page *_ba.PdfPage) (*Extractor, error) {
	_ec, _bf := page.GetAllContentStreams()
	if _bf != nil {
		return nil, _bf
	}
	_df, _bf := page.GetMediaBox()
	if _bf != nil {
		return nil, _ca.Errorf("\u0065\u0078\u0074r\u0061\u0063\u0074\u006fr\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u002e\u0020\u0025\u0076", _bf)
	}
	_acb := &Extractor{_ge: _ec, _fbe: page.Resources, _gg: *_df, _age: map[string]fontEntry{}, _afd: map[string]textResult{}}
	if _acb._gg.Llx > _acb._gg.Urx {
		_f.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0058\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _acb._gg)
		_acb._gg.Llx, _acb._gg.Urx = _acb._gg.Urx, _acb._gg.Llx
	}
	if _acb._gg.Lly > _acb._gg.Ury {
		_f.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0059\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _acb._gg)
		_acb._gg.Lly, _acb._gg.Ury = _acb._gg.Ury, _acb._gg.Lly
	}
	return _acb, nil
}
func (_fafce *textPara) taken() bool { return _fafce == nil || _fafce._bcgc }
func (_gffb *subpath) makeRectRuling(_cfbd _af.Color) (*ruling, bool) {
	if _eecd {
		_f.Log.Info("\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0070\u0061\u0074\u0068\u003d\u0025\u0076", _gffb)
	}
	_gabd := _gffb._caf[:4]
	_ggbfa := make(map[int]rulingKind, len(_gabd))
	for _cdce, _cgeca := range _gabd {
		_fbcc := _gffb._caf[(_cdce+1)%4]
		_ggbfa[_cdce] = _ebca(_cgeca, _fbcc)
		if _eecd {
			_ca.Printf("\u0025\u0034\u0064: \u0025\u0073\u0020\u003d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066", _cdce, _ggbfa[_cdce], _cgeca, _fbcc)
		}
	}
	if _eecd {
		_ca.Printf("\u0020\u0020\u0020\u006b\u0069\u006e\u0064\u0073\u003d\u0025\u002b\u0076\u000a", _ggbfa)
	}
	var _eccd, _eebb []int
	for _gadg, _gbfg := range _ggbfa {
		switch _gbfg {
		case _egd:
			_eebb = append(_eebb, _gadg)
		case _efcg:
			_eccd = append(_eccd, _gadg)
		}
	}
	if _eecd {
		_ca.Printf("\u0020\u0020 \u0068\u006f\u0072z\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_eebb), _eebb)
		_ca.Printf("\u0020\u0020 \u0076\u0065\u0072t\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_eccd), _eccd)
	}
	_dgc := (len(_eebb) == 2 && len(_eccd) == 2) || (len(_eebb) == 2 && len(_eccd) == 0 && _fgdge(_gabd[_eebb[0]], _gabd[_eebb[1]])) || (len(_eccd) == 2 && len(_eebb) == 0 && _bfgf(_gabd[_eccd[0]], _gabd[_eccd[1]]))
	if _eecd {
		_ca.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_eebb), len(_eccd), _dgc)
	}
	if !_dgc {
		if _eecd {
			_f.Log.Error("\u0021!\u006d\u0061\u006b\u0065R\u0065\u0063\u0074\u0052\u0075l\u0069n\u0067:\u0020\u0070\u0061\u0074\u0068\u003d\u0025v", _gffb)
			_ca.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_eebb), len(_eccd), _dgc)
		}
		return &ruling{}, false
	}
	if len(_eccd) == 0 {
		for _gdaa, _eace := range _ggbfa {
			if _eace != _egd {
				_eccd = append(_eccd, _gdaa)
			}
		}
	}
	if len(_eebb) == 0 {
		for _dfea, _cecd := range _ggbfa {
			if _cecd != _efcg {
				_eebb = append(_eebb, _dfea)
			}
		}
	}
	if _eecd {
		_f.Log.Info("\u006da\u006b\u0065R\u0065\u0063\u0074\u0052u\u006c\u0069\u006eg\u003a\u0020\u0068\u006f\u0072\u007a\u0073\u003d\u0025d \u0076\u0065\u0072t\u0073\u003d%\u0064\u0020\u0070\u006f\u0069\u006et\u0073\u003d%\u0064\u000a"+"\u0009\u0020\u0068o\u0072\u007a\u0073\u003d\u0025\u002b\u0076\u000a"+"\u0009\u0020\u0076e\u0072\u0074\u0073\u003d\u0025\u002b\u0076\u000a"+"\t\u0070\u006f\u0069\u006e\u0074\u0073\u003d\u0025\u002b\u0076", len(_eebb), len(_eccd), len(_gabd), _eebb, _eccd, _gabd)
	}
	var _fccge, _ffcf, _deaa, _acda _bc.Point
	if _gabd[_eebb[0]].Y > _gabd[_eebb[1]].Y {
		_deaa, _acda = _gabd[_eebb[0]], _gabd[_eebb[1]]
	} else {
		_deaa, _acda = _gabd[_eebb[1]], _gabd[_eebb[0]]
	}
	if _gabd[_eccd[0]].X > _gabd[_eccd[1]].X {
		_fccge, _ffcf = _gabd[_eccd[0]], _gabd[_eccd[1]]
	} else {
		_fccge, _ffcf = _gabd[_eccd[1]], _gabd[_eccd[0]]
	}
	_cgdcf := _ba.PdfRectangle{Llx: _fccge.X, Urx: _ffcf.X, Lly: _acda.Y, Ury: _deaa.Y}
	if _cgdcf.Llx > _cgdcf.Urx {
		_cgdcf.Llx, _cgdcf.Urx = _cgdcf.Urx, _cgdcf.Llx
	}
	if _cgdcf.Lly > _cgdcf.Ury {
		_cgdcf.Lly, _cgdcf.Ury = _cgdcf.Ury, _cgdcf.Lly
	}
	_bfbfa := rectRuling{PdfRectangle: _cgdcf, _acac: _ffee(_cgdcf), Color: _cfbd}
	if _bfbfa._acac == _edc {
		if _eecd {
			_f.Log.Error("\u006da\u006b\u0065\u0052\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006eg\u003a\u0020\u006b\u0069\u006e\u0064\u003d\u006e\u0069\u006c")
		}
		return nil, false
	}
	_fbebe, _dgge := _bfbfa.asRuling()
	if !_dgge {
		if _eecd {
			_f.Log.Error("\u006da\u006b\u0065\u0052\u0065c\u0074\u0052\u0075\u006c\u0069n\u0067:\u0020!\u0069\u0073\u0052\u0075\u006c\u0069\u006eg")
		}
		return nil, false
	}
	if _bfg {
		_ca.Printf("\u0020\u0020\u0020\u0072\u003d\u0025\u0073\u000a", _fbebe.String())
	}
	return _fbebe, true
}
func (_fffc *wordBag) pullWord(_bggfc *textWord, _caddb int, _fagb map[int]map[*textWord]struct{}) {
	_fffc.PdfRectangle = _cafce(_fffc.PdfRectangle, _bggfc.PdfRectangle)
	if _bggfc._ggea > _fffc._egad {
		_fffc._egad = _bggfc._ggea
	}
	_fffc._dfec[_caddb] = append(_fffc._dfec[_caddb], _bggfc)
	_fagb[_caddb][_bggfc] = struct{}{}
}
func (_dcaf rulingList) snapToGroupsDirection() rulingList {
	_dcaf.sortStrict()
	_gddg := make(map[*ruling]rulingList, len(_dcaf))
	_aceedc := _dcaf[0]
	_deaaf := func(_gcgd *ruling) { _aceedc = _gcgd; _gddg[_aceedc] = rulingList{_gcgd} }
	_deaaf(_dcaf[0])
	for _, _decgg := range _dcaf[1:] {
		if _decgg._ccbf < _aceedc._ccbf-_bcfg {
			_f.Log.Error("\u0073\u006e\u0061\u0070T\u006f\u0047\u0072\u006f\u0075\u0070\u0073\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0057\u0072\u006f\u006e\u0067\u0020\u0070\u0072\u0069\u006da\u0072\u0079\u0020\u006f\u0072d\u0065\u0072\u002e\u000a\u0009\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0020\u0076\u003d\u0025\u0073", _aceedc, _decgg)
		}
		if _decgg._ccbf > _aceedc._ccbf+_gedc {
			_deaaf(_decgg)
		} else {
			_gddg[_aceedc] = append(_gddg[_aceedc], _decgg)
		}
	}
	_fggf := make(map[*ruling]float64, len(_gddg))
	_cfdc := make(map[*ruling]*ruling, len(_dcaf))
	for _ffbca, _dedga := range _gddg {
		_fggf[_ffbca] = _dedga.mergePrimary()
		for _, _bbcaf := range _dedga {
			_cfdc[_bbcaf] = _ffbca
		}
	}
	for _, _egff := range _dcaf {
		_egff._ccbf = _fggf[_cfdc[_egff]]
	}
	_bcffb := make(rulingList, 0, len(_dcaf))
	for _, _fbed := range _gddg {
		_dfabd := _fbed.splitSec()
		for _eacd, _ggcg := range _dfabd {
			_eeec := _ggcg.merge()
			if len(_bcffb) > 0 {
				_eeeg := _bcffb[len(_bcffb)-1]
				if _eeeg.alignsPrimary(_eeec) && _eeeg.alignsSec(_eeec) {
					_f.Log.Error("\u0073\u006e\u0061\u0070\u0054\u006fG\u0072\u006f\u0075\u0070\u0073\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0044\u0075\u0070\u006ci\u0063\u0061\u0074\u0065\u0020\u0069\u003d\u0025\u0064\u000a\u0009\u0077\u003d\u0025s\u000a\t\u0076\u003d\u0025\u0073", _eacd, _eeeg, _eeec)
					continue
				}
			}
			_bcffb = append(_bcffb, _eeec)
		}
	}
	_bcffb.sortStrict()
	return _bcffb
}
func (_fce *textLine) appendWord(_bbbef *textWord) {
	_fce._bdcb = append(_fce._bdcb, _bbbef)
	_fce.PdfRectangle = _cafce(_fce.PdfRectangle, _bbbef.PdfRectangle)
	if _bbbef._ggea > _fce._edfec {
		_fce._edfec = _bbbef._ggea
	}
	if _bbbef._addc > _fce._fcaa {
		_fce._fcaa = _bbbef._addc
	}
}
func (_ggdff *textTable) isExportable() bool {
	if _ggdff._deecf {
		return true
	}
	_cdbcd := func(_aafge int) bool {
		_acecg := _ggdff.get(0, _aafge)
		if _acecg == nil {
			return false
		}
		_cfbfg := _acecg.text()
		_fafg := _ac.RuneCountInString(_cfbfg)
		_bcgdg := _edcgb.MatchString(_cfbfg)
		return _fafg <= 1 || _bcgdg
	}
	for _bcdef := 0; _bcdef < _ggdff._ebab; _bcdef++ {
		if !_cdbcd(_bcdef) {
			return true
		}
	}
	return false
}
func (_fdcc *textObject) getCurrentFont() *_ba.PdfFont {
	_fbebb := _fdcc._gdcg._fdf
	if _fbebb == nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020U\u0073\u0069\u006e\u0067\u0020d\u0065\u0066a\u0075\u006c\u0074\u002e")
		return _ba.DefaultFont()
	}
	return _fbebb
}
func (_eaceg *ruling) alignsPrimary(_ffgg *ruling) bool {
	return _eaceg._ecddg == _ffgg._ecddg && _c.Abs(_eaceg._ccbf-_ffgg._ccbf) < _gedc*0.5
}

type event struct {
	_afgg  float64
	_efea  bool
	_beaeg int
}

func (_gbee *textObject) checkOp(_acccf *_da.ContentStreamOperation, _ffg int, _eda bool) (_fgdg bool, _fcc error) {
	if _gbee == nil {
		var _bfea []_ff.PdfObject
		if _ffg > 0 {
			_bfea = _acccf.Params
			if len(_bfea) > _ffg {
				_bfea = _bfea[:_ffg]
			}
		}
		_f.Log.Debug("\u0025\u0023q \u006f\u0070\u0065r\u0061\u006e\u0064\u0020out\u0073id\u0065\u0020\u0074\u0065\u0078\u0074\u002e p\u0061\u0072\u0061\u006d\u0073\u003d\u0025+\u0076", _acccf.Operand, _bfea)
	}
	if _ffg >= 0 {
		if len(_acccf.Params) != _ffg {
			if _eda {
				_fcc = _g.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
			}
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _acccf.Operand, _ffg, len(_acccf.Params), _acccf.Params)
			return false, _fcc
		}
	}
	return true, nil
}
func (_cfaa *ruling) intersects(_cggag *ruling) bool {
	_ccdgea := (_cfaa._ecddg == _efcg && _cggag._ecddg == _egd) || (_cggag._ecddg == _efcg && _cfaa._ecddg == _egd)
	_cdgd := func(_bbae, _aebb *ruling) bool {
		return _bbae._gcbgg-_gdae <= _aebb._ccbf && _aebb._ccbf <= _bbae._cgeee+_gdae
	}
	_ecdf := _cdgd(_cfaa, _cggag)
	_bcbga := _cdgd(_cggag, _cfaa)
	if _bfg {
		_ca.Printf("\u0020\u0020\u0020\u0020\u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003a\u0020\u0020\u006fr\u0074\u0068\u006f\u0067\u006f\u006e\u0061l\u003d\u0025\u0074\u0020\u006f\u0031\u003d\u0025\u0074\u0020\u006f2\u003d\u0025\u0074\u0020\u2192\u0020\u0025\u0074\u000a"+"\u0020\u0020\u0020 \u0020\u0020\u0020\u0076\u003d\u0025\u0073\u000a"+" \u0020\u0020\u0020\u0020\u0020\u0077\u003d\u0025\u0073\u000a", _ccdgea, _ecdf, _bcbga, _ccdgea && _ecdf && _bcbga, _cfaa, _cggag)
	}
	return _ccdgea && _ecdf && _bcbga
}
func _ebegb(_edcc, _fdcec, _ecab float64) rulingKind {
	if _edcc >= _ecab && _cded(_fdcec, _edcc) {
		return _egd
	}
	if _fdcec >= _ecab && _cded(_edcc, _fdcec) {
		return _efcg
	}
	return _edc
}
func (_ceaf rectRuling) asRuling() (*ruling, bool) {
	_gcc := ruling{_ecddg: _ceaf._acac, Color: _ceaf.Color, _gdee: _ddbe}
	switch _ceaf._acac {
	case _efcg:
		_gcc._ccbf = 0.5 * (_ceaf.Llx + _ceaf.Urx)
		_gcc._gcbgg = _ceaf.Lly
		_gcc._cgeee = _ceaf.Ury
		_eacc, _ccgfg := _ceaf.checkWidth(_ceaf.Llx, _ceaf.Urx)
		if !_ccgfg {
			if _eecd {
				_f.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067V\u0065\u0072\u0074\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _ceaf)
			}
			return nil, false
		}
		_gcc._eeca = _eacc
	case _egd:
		_gcc._ccbf = 0.5 * (_ceaf.Lly + _ceaf.Ury)
		_gcc._gcbgg = _ceaf.Llx
		_gcc._cgeee = _ceaf.Urx
		_gggf, _gdgc := _ceaf.checkWidth(_ceaf.Lly, _ceaf.Ury)
		if !_gdgc {
			if _eecd {
				_f.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067H\u006f\u0072\u007a\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _ceaf)
			}
			return nil, false
		}
		_gcc._eeca = _gggf
	default:
		_f.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _ceaf._acac)
		return nil, false
	}
	return &_gcc, true
}

type paraList []*textPara

func (_febea *textWord) appendMark(_caddd *textMark, _cggcfc _ba.PdfRectangle) {
	_febea._bdedd = append(_febea._bdedd, _caddd)
	_febea.PdfRectangle = _cafce(_febea.PdfRectangle, _caddd.PdfRectangle)
	if _caddd._dfca > _febea._ggea {
		_febea._ggea = _caddd._dfca
	}
	_febea._addc = _cggcfc.Ury - _febea.PdfRectangle.Lly
}
func (_fgd *imageExtractContext) processOperand(_fgfg *_da.ContentStreamOperation, _gabg _da.GraphicsState, _gbg *_ba.PdfPageResources) error {
	if _fgfg.Operand == "\u0042\u0049" && len(_fgfg.Params) == 1 {
		_cc, _ded := _fgfg.Params[0].(*_da.ContentStreamInlineImage)
		if !_ded {
			return nil
		}
		if _ece, _bbd := _ff.GetBoolVal(_cc.ImageMask); _bbd {
			if _ece && !_fgd._caa.IncludeInlineStencilMasks {
				return nil
			}
		}
		return _fgd.extractInlineImage(_cc, _gabg, _gbg)
	} else if _fgfg.Operand == "\u0044\u006f" && len(_fgfg.Params) == 1 {
		_gec, _fac := _ff.GetName(_fgfg.Params[0])
		if !_fac {
			_f.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _fb
		}
		_, _ee := _gbg.GetXObjectByName(*_gec)
		switch _ee {
		case _ba.XObjectTypeImage:
			return _fgd.extractXObjectImage(_gec, _gabg, _gbg)
		case _ba.XObjectTypeForm:
			return _fgd.extractFormImages(_gec, _gabg, _gbg)
		}
	}
	return nil
}
func (_fbge *wordBag) minDepth() float64 { return _fbge._efae - (_fbge.Ury - _fbge._egad) }

// String returns a description of `k`.
func (_feab markKind) String() string {
	_gbaf, _abfg := _cagg[_feab]
	if !_abfg {
		return _ca.Sprintf("\u004e\u006f\u0074\u0020\u0061\u0020\u006d\u0061\u0072k\u003a\u0020\u0025\u0064", _feab)
	}
	return _gbaf
}
func (_fecb compositeCell) String() string {
	_ddbb := ""
	if len(_fecb.paraList) > 0 {
		_ddbb = _bbfd(_fecb.paraList.merge().text(), 50)
	}
	return _ca.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0070\u0061\u0072a\u0073\u0020\u0025\u0071", _fecb.PdfRectangle, len(_fecb.paraList), _ddbb)
}
func (_cefc paraList) eventNeighbours(_gddde []event) map[*textPara][]int {
	_dc.Slice(_gddde, func(_abdga, _abebg int) bool {
		_cegdc, _ebbee := _gddde[_abdga], _gddde[_abebg]
		_fbdgb, _cccc := _cegdc._afgg, _ebbee._afgg
		if _fbdgb != _cccc {
			return _fbdgb < _cccc
		}
		if _cegdc._efea != _ebbee._efea {
			return _cegdc._efea
		}
		return _abdga < _abebg
	})
	_fdbe := make(map[int]intSet)
	_agde := make(intSet)
	for _, _bfbe := range _gddde {
		if _bfbe._efea {
			_fdbe[_bfbe._beaeg] = make(intSet)
			for _bgdf := range _agde {
				if _bgdf != _bfbe._beaeg {
					_fdbe[_bfbe._beaeg].add(_bgdf)
					_fdbe[_bgdf].add(_bfbe._beaeg)
				}
			}
			_agde.add(_bfbe._beaeg)
		} else {
			_agde.del(_bfbe._beaeg)
		}
	}
	_cefa := map[*textPara][]int{}
	for _gggce, _gceca := range _fdbe {
		_dfbbc := _cefc[_gggce]
		if len(_gceca) == 0 {
			_cefa[_dfbbc] = nil
			continue
		}
		_bgdcg := make([]int, len(_gceca))
		_egeg := 0
		for _cddg := range _gceca {
			_bgdcg[_egeg] = _cddg
			_egeg++
		}
		_cefa[_dfbbc] = _bgdcg
	}
	return _cefa
}
func _bfddc(_fcce bounded) float64 { return -_fcce.bbox().Lly }
func (_bgaeg *textPara) text() string {
	_bdaag := new(_gb.Buffer)
	_bgaeg.writeText(_bdaag)
	return _bdaag.String()
}
func _gafbc(_ggca _ba.PdfRectangle, _dffc, _dbaab, _ddfgc, _gbgag *ruling) gridTile {
	_fgcfd := _ggca.Llx
	_bbcc := _ggca.Urx
	_ecggbd := _ggca.Lly
	_abad := _ggca.Ury
	return gridTile{PdfRectangle: _ggca, _fdge: _dffc != nil && _dffc.encloses(_ecggbd, _abad), _abbbb: _dbaab != nil && _dbaab.encloses(_ecggbd, _abad), _edcg: _ddfgc != nil && _ddfgc.encloses(_fgcfd, _bbcc), _aebbe: _gbgag != nil && _gbgag.encloses(_fgcfd, _bbcc)}
}
func _fcba(_bcb *wordBag, _bafd float64, _aecaa, _dgde rulingList) []*wordBag {
	var _cbda []*wordBag
	for _, _ffcc := range _bcb.depthIndexes() {
		_fgca := false
		for !_bcb.empty(_ffcc) {
			_cdagg := _bcb.firstReadingIndex(_ffcc)
			_debdg := _bcb.firstWord(_cdagg)
			_adac := _ceag(_debdg, _bafd, _aecaa, _dgde)
			_bcb.removeWord(_debdg, _cdagg)
			if _gbgf {
				_f.Log.Info("\u0066\u0069\u0072\u0073\u0074\u0057\u006f\u0072\u0064\u0020\u005e\u005e^\u005e\u0020\u0025\u0073", _debdg.String())
			}
			for _edge := true; _edge; _edge = _fgca {
				_fgca = false
				_abgg := _cdgee * _adac._egad
				_agaf := _dfgd * _adac._egad
				_fedfd := _cafa * _adac._egad
				if _gbgf {
					_f.Log.Info("\u0070a\u0072a\u0057\u006f\u0072\u0064\u0073\u0020\u0064\u0065\u0070\u0074\u0068 \u0025\u002e\u0032\u0066 \u002d\u0020\u0025\u002e\u0032f\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061\u0044\u0065\u0070\u0074\u0068\u0047\u0061\u0070\u003d\u0025\u002e\u0032\u0066\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061R\u0065\u0061\u0064\u0069\u006e\u0067\u0047\u0061p\u003d\u0025\u002e\u0032\u0066", _adac.minDepth(), _adac.maxDepth(), _fedfd, _agaf)
				}
				if _bcb.scanBand("\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c", _adac, _gdddd(_ccdd, 0), _adac.minDepth()-_fedfd, _adac.maxDepth()+_fedfd, _bge, false, false) > 0 {
					_fgca = true
				}
				if _bcb.scanBand("\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _adac, _gdddd(_ccdd, _agaf), _adac.minDepth(), _adac.maxDepth(), _bgcef, false, false) > 0 {
					_fgca = true
				}
				if _fgca {
					continue
				}
				_cdda := _bcb.scanBand("", _adac, _gdddd(_eecb, _abgg), _adac.minDepth(), _adac.maxDepth(), _caed, true, false)
				if _cdda > 0 {
					_gbfc := (_adac.maxDepth() - _adac.minDepth()) / _adac._egad
					if (_cdda > 1 && float64(_cdda) > 0.3*_gbfc) || _cdda <= 10 {
						if _bcb.scanBand("\u006f\u0074\u0068e\u0072", _adac, _gdddd(_eecb, _abgg), _adac.minDepth(), _adac.maxDepth(), _caed, false, true) > 0 {
							_fgca = true
						}
					}
				}
			}
			_cbda = append(_cbda, _adac)
		}
	}
	return _cbda
}
func (_beaba *textTable) subdivide() *textTable {
	_beaba.logComposite("\u0073u\u0062\u0064\u0069\u0076\u0069\u0064e")
	_cdbe := _beaba.compositeRowCorridors()
	_beaa := _beaba.compositeColCorridors()
	if _eca {
		_f.Log.Info("\u0073u\u0062\u0064i\u0076\u0069\u0064\u0065:\u000a\u0009\u0072o\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s=\u0025\u0073\u000a\t\u0063\u006fl\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073=\u0025\u0073", _adcf(_cdbe), _adcf(_beaa))
	}
	if len(_cdbe) == 0 || len(_beaa) == 0 {
		return _beaba
	}
	_eadc(_cdbe)
	_eadc(_beaa)
	if _eca {
		_f.Log.Info("\u0073\u0075\u0062\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0066\u0069\u0078\u0065\u0064\u003a\u000a\u0009r\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073\u003d\u0025\u0073\u000a\u0009\u0063\u006f\u006cC\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d\u0025\u0073", _adcf(_cdbe), _adcf(_beaa))
	}
	_dffgc, _ffbgf := _ffff(_beaba._ebab, _cdbe)
	_aadce, _caaf := _ffff(_beaba._gcead, _beaa)
	_dgdbf := make(map[uint64]*textPara, _caaf*_ffbgf)
	_ggfcc := &textTable{PdfRectangle: _beaba.PdfRectangle, _deecf: _beaba._deecf, _ebab: _ffbgf, _gcead: _caaf, _fecac: _dgdbf}
	if _eca {
		_f.Log.Info("\u0073\u0075b\u0064\u0069\u0076\u0069\u0064\u0065\u003a\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0063\u0065\u006c\u006c\u0073\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0079\u004f\u0066\u0066\u0073\u0065\u0074\u0073=\u0025\u002b\u0076\u000a"+"\u0009\u0078\u004f\u0066\u0066\u0073\u0065\u0074\u0073\u003d\u0025\u002b\u0076", _beaba._gcead, _beaba._ebab, _caaf, _ffbgf, _adcf(_cdbe), _adcf(_beaa), _dffgc, _aadce)
	}
	for _gfga := 0; _gfga < _beaba._ebab; _gfga++ {
		_fdaf := _dffgc[_gfga]
		for _efdb := 0; _efdb < _beaba._gcead; _efdb++ {
			_efag := _aadce[_efdb]
			if _eca {
				_ca.Printf("\u0025\u0036\u0064\u002c %\u0032\u0064\u003a\u0020\u0078\u0030\u003d\u0025\u0064\u0020\u0079\u0030\u003d\u0025d\u000a", _efdb, _gfga, _efag, _fdaf)
			}
			_afdd, _ebbcf := _beaba._gabfc[_dfee(_efdb, _gfga)]
			if !_ebbcf {
				continue
			}
			_bddgf := _afdd.split(_cdbe[_gfga], _beaa[_efdb])
			for _dgaac := 0; _dgaac < _bddgf._ebab; _dgaac++ {
				for _fbdg := 0; _fbdg < _bddgf._gcead; _fbdg++ {
					_bgca := _bddgf.get(_fbdg, _dgaac)
					_ggfcc.put(_efag+_fbdg, _fdaf+_dgaac, _bgca)
					if _eca {
						_ca.Printf("\u0025\u0038\u0064\u002c\u0020\u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _efag+_fbdg, _fdaf+_dgaac, _bgca)
					}
				}
			}
		}
	}
	return _ggfcc
}
func (_fec *textObject) setHorizScaling(_debe float64) {
	if _fec == nil {
		return
	}
	_fec._gdcg._fgdd = _debe
}
func (_fabd paraList) lines() []*textLine {
	var _fbac []*textLine
	for _, _eafc := range _fabd {
		_fbac = append(_fbac, _eafc._efaf...)
	}
	return _fbac
}

// TextTable represents a table.
// Cells are ordered top-to-bottom, left-to-right.
// Cells[y] is the (0-offset) y'th row in the table.
// Cells[y][x] is the (0-offset) x'th column in the table.
type TextTable struct {
	W, H  int
	Cells [][]TableCell
}

func _fafcd(_bcaeec _ba.PdfColorspace, _bcaeg _ba.PdfColor) _af.Color {
	if _bcaeec == nil || _bcaeg == nil {
		return _af.Black
	}
	_fbecg, _edgd := _bcaeec.ColorToRGB(_bcaeg)
	if _edgd != nil {
		_f.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006fu\u006c\u0064\u0020no\u0074\u0020\u0063\u006f\u006e\u0076e\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0025\u0076\u0020\u0028\u0025\u0076)\u0020\u0074\u006f\u0020\u0052\u0047\u0042\u003a \u0025\u0073", _bcaeg, _bcaeec, _edgd)
		return _af.Black
	}
	_ddfcd, _abeef := _fbecg.(*_ba.PdfColorDeviceRGB)
	if !_abeef {
		_f.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0052\u0047\u0042\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0076", _fbecg)
		return _af.Black
	}
	return _af.NRGBA{R: uint8(_ddfcd.R() * 255), G: uint8(_ddfcd.G() * 255), B: uint8(_ddfcd.B() * 255), A: uint8(255)}
}
func (_efcgg *textTable) toTextTable() TextTable {
	if _eca {
		_f.Log.Info("t\u006fT\u0065\u0078\u0074\u0054\u0061\u0062\u006c\u0065:\u0020\u0025\u0064\u0020x \u0025\u0064", _efcgg._gcead, _efcgg._ebab)
	}
	_aabg := make([][]TableCell, _efcgg._ebab)
	for _gcbba := 0; _gcbba < _efcgg._ebab; _gcbba++ {
		_aabg[_gcbba] = make([]TableCell, _efcgg._gcead)
		for _gafff := 0; _gafff < _efcgg._gcead; _gafff++ {
			_geag := _efcgg.get(_gafff, _gcbba)
			if _geag == nil {
				continue
			}
			if _eca {
				_ca.Printf("\u0025\u0034\u0064 \u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _gafff, _gcbba, _geag)
			}
			_aabg[_gcbba][_gafff].Text = _geag.text()
			_bdff := 0
			_aabg[_gcbba][_gafff].Marks._afdf = _geag.toTextMarks(&_bdff)
		}
	}
	return TextTable{W: _efcgg._gcead, H: _efcgg._ebab, Cells: _aabg}
}
func _cded(_eacf, _efbba float64) bool { return _eacf/_c.Max(_cbbb, _efbba) < _dae }
func _ebca(_fddeb, _caeg _bc.Point) rulingKind {
	_ffea := _c.Abs(_fddeb.X - _caeg.X)
	_dgged := _c.Abs(_fddeb.Y - _caeg.Y)
	return _ebegb(_ffea, _dgged, _dae)
}
func (_gdaf *textTable) computeBbox() _ba.PdfRectangle {
	var _afedg _ba.PdfRectangle
	_cfdb := false
	for _dfgeb := 0; _dfgeb < _gdaf._ebab; _dfgeb++ {
		for _fdcff := 0; _fdcff < _gdaf._gcead; _fdcff++ {
			_feea := _gdaf.get(_fdcff, _dfgeb)
			if _feea == nil {
				continue
			}
			if !_cfdb {
				_afedg = _feea.PdfRectangle
				_cfdb = true
			} else {
				_afedg = _cafce(_afedg, _feea.PdfRectangle)
			}
		}
	}
	return _afedg
}
func _bbdg(_faecb float64) float64                       { return _ebf * _c.Round(_faecb/_ebf) }
func (_aadb *textTable) get(_gaaae, _fccf int) *textPara { return _aadb._fecac[_dfee(_gaaae, _fccf)] }
func _bddd(_gfae []TextMark, _dbff *int, _ecega TextMark) []TextMark {
	_ecega.Offset = *_dbff
	_gfae = append(_gfae, _ecega)
	*_dbff += len(_ecega.Text)
	return _gfae
}

// String returns a string describing the current state of the textState stack.
func (_dage *stateStack) String() string {
	_cdeb := []string{_ca.Sprintf("\u002d\u002d\u002d\u002d f\u006f\u006e\u0074\u0020\u0073\u0074\u0061\u0063\u006b\u003a\u0020\u0025\u0064", len(*_dage))}
	for _edg, _dbf := range *_dage {
		_gcee := "\u003c\u006e\u0069l\u003e"
		if _dbf != nil {
			_gcee = _dbf.String()
		}
		_cdeb = append(_cdeb, _ca.Sprintf("\u0009\u0025\u0032\u0064\u003a\u0020\u0025\u0073", _edg, _gcee))
	}
	return _ad.Join(_cdeb, "\u000a")
}
func (_aaad *textMark) bbox() _ba.PdfRectangle { return _aaad.PdfRectangle }
func (_cagc paraList) applyTables(_edda []*textTable) paraList {
	var _eeee paraList
	for _, _dcfe := range _edda {
		_eeee = append(_eeee, _dcfe.newTablePara())
	}
	for _, _gdfe := range _cagc {
		if _gdfe._bcgc {
			continue
		}
		_eeee = append(_eeee, _gdfe)
	}
	return _eeee
}
func (_fdae rulingList) connections(_efafa map[int]intSet, _fdfeb int) intSet {
	_caca := make(intSet)
	_cgef := make(intSet)
	var _becf func(int)
	_becf = func(_dcdf int) {
		if !_cgef.has(_dcdf) {
			_cgef.add(_dcdf)
			for _bcfe := range _fdae {
				if _efafa[_bcfe].has(_dcdf) {
					_caca.add(_bcfe)
				}
			}
			for _ggbe := range _fdae {
				if _caca.has(_ggbe) {
					_becf(_ggbe)
				}
			}
		}
	}
	_becf(_fdfeb)
	return _caca
}
func _gcge(_cbea _ba.PdfRectangle) *ruling {
	return &ruling{_ecddg: _egd, _ccbf: _cbea.Lly, _gcbgg: _cbea.Llx, _cgeee: _cbea.Urx}
}
func (_ebfac rulingList) toTilings() (rulingList, []gridTiling) {
	_ebfac.log("\u0074o\u0054\u0069\u006c\u0069\u006e\u0067s")
	if len(_ebfac) == 0 {
		return nil, nil
	}
	_ebfac = _ebfac.tidied("\u0061\u006c\u006c")
	_ebfac.log("\u0074\u0069\u0064\u0069\u0065\u0064")
	_fded := _ebfac.toGrids()
	_dcbdg := make([]gridTiling, len(_fded))
	for _afae, _gadf := range _fded {
		_dcbdg[_afae] = _gadf.asTiling()
	}
	return _ebfac, _dcbdg
}
func (_ada *PageFonts) extractPageResourcesToFont(_de *_ba.PdfPageResources) error {
	_cee, _cg := _ff.GetDict(_de.Font)
	if !_cg {
		return _g.New(_bfa)
	}
	for _, _ffaa := range _cee.Keys() {
		var (
			_ae  = true
			_ab  []byte
			_bbf string
		)
		_db, _bfad := _de.GetFontByName(_ffaa)
		if !_bfad {
			return _g.New(_bca)
		}
		_ea, _fe := _ba.NewPdfFontFromPdfObject(_db)
		if _fe != nil {
			return _fe
		}
		_dcg := _ea.FontDescriptor()
		_dba := _ea.FontDescriptor().FontName.String()
		_be := _ea.Subtype()
		if _dcb(_ada.Fonts, _dba) {
			continue
		}
		if len(_ea.ToUnicode()) == 0 {
			_ae = false
		}
		if _dcg.FontFile != nil {
			if _bfadf, _deg := _ff.GetStream(_dcg.FontFile); _deg {
				_ab, _fe = _ff.DecodeStream(_bfadf)
				if _fe != nil {
					return _fe
				}
				_bbf = _dba + "\u002e\u0070\u0066\u0062"
			}
		} else if _dcg.FontFile2 != nil {
			if _dd, _bed := _ff.GetStream(_dcg.FontFile2); _bed {
				_ab, _fe = _ff.DecodeStream(_dd)
				if _fe != nil {
					return _fe
				}
				_bbf = _dba + "\u002e\u0074\u0074\u0066"
			}
		} else if _dcg.FontFile3 != nil {
			if _fga, _afb := _ff.GetStream(_dcg.FontFile3); _afb {
				_ab, _fe = _ff.DecodeStream(_fga)
				if _fe != nil {
					return _fe
				}
				_bbf = _dba + "\u002e\u0063\u0066\u0066"
			}
		}
		if len(_bbf) < 1 {
			_f.Log.Debug(_fba)
		}
		_def := Font{FontName: _dba, PdfFont: _ea, IsCID: _ea.IsCID(), IsSimple: _ea.IsSimple(), ToUnicode: _ae, FontType: _be, FontData: _ab, FontFileName: _bbf, FontDescriptor: _dcg}
		_ada.Fonts = append(_ada.Fonts, _def)
	}
	return nil
}
func (_bdd *shapesState) cubicTo(_eaba, _gabb, _debb, _gbbc, _eec, _efg float64) {
	if _dfab {
		_f.Log.Info("\u0063\u0075\u0062\u0069\u0063\u0054\u006f\u003a")
	}
	_bdd.addPoint(_eec, _efg)
}

const (
	_gcfaf markKind = iota
	_deead
	_ddbe
	_ggbdb
)

func (_agbc gridTile) contains(_cedc _ba.PdfRectangle) bool {
	if _agbc.numBorders() < 3 {
		return false
	}
	if _agbc._fdge && _cedc.Llx < _agbc.Llx-_effd {
		return false
	}
	if _agbc._abbbb && _cedc.Urx > _agbc.Urx+_effd {
		return false
	}
	if _agbc._edcg && _cedc.Lly < _agbc.Lly-_effd {
		return false
	}
	if _agbc._aebbe && _cedc.Ury > _agbc.Ury+_effd {
		return false
	}
	return true
}
func (_gea *imageExtractContext) extractContentStreamImages(_deb string, _fc *_ba.PdfPageResources) error {
	_gbb := _da.NewContentStreamParser(_deb)
	_fed, _dfa := _gbb.Parse()
	if _dfa != nil {
		return _dfa
	}
	if _gea._ecdd == nil {
		_gea._ecdd = map[*_ff.PdfObjectStream]*cachedImage{}
	}
	if _gea._caa == nil {
		_gea._caa = &ImageExtractOptions{}
	}
	_cac := _da.NewContentStreamProcessor(*_fed)
	_cac.AddHandler(_da.HandlerConditionEnumAllOperands, "", _gea.processOperand)
	return _cac.Process(_fc)
}
func (_gbef rulingList) log(_dcfce string) {
	if !_bfg {
		return
	}
	_f.Log.Info("\u0023\u0023\u0023\u0020\u0025\u0031\u0030\u0073\u003a\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _dcfce, _gbef.String())
	for _acfc, _edccc := range _gbef {
		_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _acfc, _edccc.String())
	}
}
func _dadd(_gdcd, _bgcgd _bc.Point) bool { return _gdcd.X == _bgcgd.X && _gdcd.Y == _bgcgd.Y }

type textResult struct {
	_fda  PageText
	_ead  int
	_fggb int
}

func (_gd *imageExtractContext) extractFormImages(_cdd *_ff.PdfObjectName, _fbc _da.GraphicsState, _ged *_ba.PdfPageResources) error {
	_afee, _eab := _ged.GetXObjectFormByName(*_cdd)
	if _eab != nil {
		return _eab
	}
	if _afee == nil {
		return nil
	}
	_ggb, _eab := _afee.GetContentStream()
	if _eab != nil {
		return _eab
	}
	_defc := _afee.Resources
	if _defc == nil {
		_defc = _ged
	}
	_eab = _gd.extractContentStreamImages(string(_ggb), _defc)
	if _eab != nil {
		return _eab
	}
	_gd._bee++
	return nil
}

// RenderMode specifies the text rendering mode (Tmode), which determines whether showing text shall cause
// glyph outlines to be  stroked, filled, used as a clipping boundary, or some combination of the three.
// Stroking, filling, and clipping shall have the same effects for a text object as they do for a path object
// (see 8.5.3, "Path-Painting Operators" and 8.5.4, "Clipping Path Operators").
type RenderMode int
type markKind int

func _dceb(_bbdb map[int][]float64) []int {
	_eafb := make([]int, len(_bbdb))
	_eggcc := 0
	for _edeaf := range _bbdb {
		_eafb[_eggcc] = _edeaf
		_eggcc++
	}
	_dc.Ints(_eafb)
	return _eafb
}
func _dabdc(_efaa, _eababe _ba.PdfRectangle) bool {
	return _efaa.Lly <= _eababe.Ury && _eababe.Lly <= _efaa.Ury
}
func _cfdg(_eced _bc.Point) _bc.Matrix { return _bc.TranslationMatrix(_eced.X, _eced.Y) }

var _abeaa = map[rulingKind]string{_edc: "\u006e\u006f\u006e\u0065", _egd: "\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _efcg: "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c"}

func (_ggeda *textTable) emptyCompositeColumn(_aaec int) bool {
	for _dadbe := 0; _dadbe < _ggeda._ebab; _dadbe++ {
		if _abdf, _bfcf := _ggeda._gabfc[_dfee(_aaec, _dadbe)]; _bfcf {
			if len(_abdf.paraList) > 0 {
				return false
			}
		}
	}
	return true
}
func (_ccca *textObject) setWordSpacing(_dggd float64) {
	if _ccca == nil {
		return
	}
	_ccca._gdcg._cef = _dggd
}
func _ecaa(_ecfef []*textWord, _bgfb int) []*textWord {
	_bcedd := len(_ecfef)
	copy(_ecfef[_bgfb:], _ecfef[_bgfb+1:])
	return _ecfef[:_bcedd-1]
}

// PageImages represents extracted images on a PDF page with spatial information:
// display position and size.
type PageImages struct{ Images []ImageMark }

func (_acfe lineRuling) yMean() float64 { return 0.5 * (_acfe._abeb.Y + _acfe._cgae.Y) }
func (_befga *textWord) computeText() string {
	_eecbb := make([]string, len(_befga._bdedd))
	for _cbbdc, _ecegg := range _befga._bdedd {
		_eecbb[_cbbdc] = _ecegg._abdg
	}
	return _ad.Join(_eecbb, "")
}
func _gafa(_bbbe _ba.PdfRectangle) textState {
	return textState{_fgdd: 100, _adcb: RenderModeFill, _gbff: _bbbe}
}
func (_dddb *imageExtractContext) extractXObjectImage(_afe *_ff.PdfObjectName, _cfc _da.GraphicsState, _bfb *_ba.PdfPageResources) error {
	_eg, _ := _bfb.GetXObjectByName(*_afe)
	if _eg == nil {
		return nil
	}
	_dag, _fea := _dddb._ecdd[_eg]
	if !_fea {
		_aef, _cab := _bfb.GetXObjectImageByName(*_afe)
		if _cab != nil {
			return _cab
		}
		if _aef == nil {
			return nil
		}
		_cbe, _cab := _aef.ToImage()
		if _cab != nil {
			return _cab
		}
		_dag = &cachedImage{_feg: _cbe, _gfe: _aef.ColorSpace}
		_dddb._ecdd[_eg] = _dag
	}
	_egb := _dag._feg
	_aee := _dag._gfe
	_fgc, _bce := _aee.ImageToRGB(*_egb)
	if _bce != nil {
		return _bce
	}
	_f.Log.Debug("@\u0044\u006f\u0020\u0043\u0054\u004d\u003a\u0020\u0025\u0073", _cfc.CTM.String())
	_bbb := ImageMark{Image: &_fgc, Width: _cfc.CTM.ScalingFactorX(), Height: _cfc.CTM.ScalingFactorY(), Angle: _cfc.CTM.Angle()}
	_bbb.X, _bbb.Y = _cfc.CTM.Translation()
	_dddb._aeb = append(_dddb._aeb, _bbb)
	_dddb._afbe++
	return nil
}
func _gdddd(_cfcb func(*wordBag, *textWord, float64) bool, _bada float64) func(*wordBag, *textWord) bool {
	return func(_fbeg *wordBag, _cgbb *textWord) bool { return _cfcb(_fbeg, _cgbb, _bada) }
}
func _bdcfc(_ggga float64) bool { return _c.Abs(_ggga) < _bcfg }
func (_eadf *PageText) computeViews() {
	var _ddcf rulingList
	if _fbbc {
		_dcgf := _fcga(_eadf._cfde)
		_ddcf = append(_ddcf, _dcgf...)
	}
	if _afbab {
		_gffg := _ccda(_eadf._fdd)
		_ddcf = append(_ddcf, _gffg...)
	}
	_ddcf, _fdg := _ddcf.toTilings()
	var _cgde paraList
	_cae := len(_eadf._faga)
	for _fdcf := 0; _fdcf < 360 && _cae > 0; _fdcf += 90 {
		_efa := make([]*textMark, 0, len(_eadf._faga)-_cae)
		for _, _aaef := range _eadf._faga {
			if _aaef._bdba == _fdcf {
				_efa = append(_efa, _aaef)
			}
		}
		if len(_efa) > 0 {
			_efcc := _ccfc(_efa, _eadf._gfed, _ddcf, _fdg)
			_cgde = append(_cgde, _efcc...)
			_cae -= len(_efa)
		}
	}
	_gbce := new(_gb.Buffer)
	_cgde.writeText(_gbce)
	_eadf._gcdc = _gbce.String()
	_eadf._cdge = _cgde.toTextMarks()
	_eadf._defb = _cgde.tables()
	if _eca {
		_f.Log.Info("\u0063\u006f\u006dpu\u0074\u0065\u0056\u0069\u0065\u0077\u0073\u003a\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064", len(_eadf._defb))
	}
}
func (_fgff *textTable) getDown() paraList {
	_afcb := make(paraList, _fgff._gcead)
	for _ebbbd := 0; _ebbbd < _fgff._gcead; _ebbbd++ {
		_dcgef := _fgff.get(_ebbbd, _fgff._ebab-1)._gagag
		if _dcgef.taken() {
			return nil
		}
		_afcb[_ebbbd] = _dcgef
	}
	for _ddff := 0; _ddff < _fgff._gcead-1; _ddff++ {
		if _afcb[_ddff]._gfeab != _afcb[_ddff+1] {
			return nil
		}
	}
	return _afcb
}

type intSet map[int]struct{}

// ExtractFonts returns all font information from the page extractor, including
// font name, font type, the raw data of the embedded font file (if embedded), font descriptor and more.
//
// The argument `previousPageFonts` is used when trying to build a complete font catalog for multiple pages or the entire document.
// The entries from `previousPageFonts` are added to the returned result unless already included in the page, i.e. no duplicate entries.
//
// NOTE: If previousPageFonts is nil, all fonts from the page will be returned. Use it when building up a full list of fonts for a document or page range.
func (_ffa *Extractor) ExtractFonts(previousPageFonts *PageFonts) (*PageFonts, error) {
	_dg := PageFonts{}
	_eb := _dg.extractPageResourcesToFont(_ffa._fbe)
	if _eb != nil {
		return nil, _eb
	}
	if previousPageFonts != nil {
		for _, _agg := range previousPageFonts.Fonts {
			if !_dcb(_dg.Fonts, _agg.FontName) {
				_dg.Fonts = append(_dg.Fonts, _agg)
			}
		}
	}
	return &PageFonts{Fonts: _dg.Fonts}, nil
}
func (_dfg *textObject) setCharSpacing(_bffa float64) {
	if _dfg == nil {
		return
	}
	_dfg._gdcg._aed = _bffa
	if _gbdd {
		_f.Log.Info("\u0073\u0065t\u0043\u0068\u0061\u0072\u0053\u0070\u0061\u0063\u0069\u006e\u0067\u003a\u0020\u0025\u002e\u0032\u0066\u0020\u0073\u0074\u0061\u0074e=\u0025\u0073", _bffa, _dfg._gdcg.String())
	}
}
func (_fbbg *ruling) alignsSec(_dfac *ruling) bool {
	const _ceaa = _gedc + 1.0
	return _fbbg._gcbgg-_ceaa <= _dfac._cgeee && _dfac._gcbgg-_ceaa <= _fbbg._cgeee
}

// String returns a string describing `pt`.
func (_ffe PageText) String() string {
	_cfa := _ca.Sprintf("P\u0061\u0067\u0065\u0054ex\u0074:\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073", len(_ffe._faga))
	_cdfe := []string{"\u002d" + _cfa}
	for _, _fgeb := range _ffe._faga {
		_cdfe = append(_cdfe, _fgeb.String())
	}
	_cdfe = append(_cdfe, "\u002b"+_cfa)
	return _ad.Join(_cdfe, "\u000a")
}

// Extractor stores and offers functionality for extracting content from PDF pages.
type Extractor struct {
	_ge  string
	_fbe *_ba.PdfPageResources
	_gg  _ba.PdfRectangle
	_age map[string]fontEntry
	_afd map[string]textResult
	_fg  int64
	_e   int
}

func (_bfcfe paraList) findTableGrid(_gada gridTiling) (*textTable, map[*textPara]struct{}) {
	_bbefa := len(_gada._eade)
	_cagde := len(_gada._bcgfa)
	_cffbg := textTable{_deecf: true, _gcead: _bbefa, _ebab: _cagde, _fecac: make(map[uint64]*textPara, _bbefa*_cagde), _gabfc: make(map[uint64]compositeCell, _bbefa*_cagde)}
	_fdeba := make(map[*textPara]struct{})
	_fgdbd := int((1.0 - _bdfg) * float64(_bbefa*_cagde))
	_gbfb := 0
	if _acgbc {
		_f.Log.Info("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0064\u0020\u0078\u0020%\u0064", _bbefa, _cagde)
	}
	for _aaae, _gbcb := range _gada._bcgfa {
		_bfec, _abdb := _gada._bbca[_gbcb]
		if !_abdb {
			continue
		}
		for _fbffc, _afcd := range _gada._eade {
			_aeabd, _cacg := _bfec[_afcd]
			if !_cacg {
				continue
			}
			_badcg := _bfcfe.inTile(_aeabd)
			if len(_badcg) == 0 {
				_gbfb++
				if _gbfb > _fgdbd {
					if _acgbc {
						_f.Log.Info("\u0021\u006e\u0075m\u0045\u006d\u0070\u0074\u0079\u003d\u0025\u0064", _gbfb)
					}
					return nil, nil
				}
			} else {
				_cffbg.putComposite(_fbffc, _aaae, _badcg, _aeabd.PdfRectangle)
				for _, _dcec := range _badcg {
					_fdeba[_dcec] = struct{}{}
				}
			}
		}
	}
	_cgeea := 0
	for _edca := 0; _edca < _bbefa; _edca++ {
		_bbeac := _cffbg.get(_edca, 0)
		if _bbeac == nil || !_bbeac._bgea {
			_cgeea++
		}
	}
	if _cgeea == 0 {
		if _acgbc {
			_f.Log.Info("\u0021\u006e\u0075m\u0048\u0065\u0061\u0064\u0065\u0072\u003d\u0030")
		}
		return nil, nil
	}
	_ddfff := _cffbg.reduceTiling(_gada, _addag)
	_ddfff = _ddfff.subdivide()
	return _ddfff, _fdeba
}
func (_faaa *wordBag) removeDuplicates() {
	if _fdfe {
		_f.Log.Info("r\u0065m\u006f\u0076\u0065\u0044\u0075\u0070\u006c\u0069c\u0061\u0074\u0065\u0073: \u0025\u0071", _faaa.text())
	}
	for _, _dfcd := range _faaa.depthIndexes() {
		if len(_faaa._dfec[_dfcd]) == 0 {
			continue
		}
		_adbe := _faaa._dfec[_dfcd][0]
		_fgcff := _caeb * _adbe._ggea
		_ddgf := _adbe._addc
		for _, _edeg := range _faaa.depthBand(_ddgf, _ddgf+_fgcff) {
			_eaaa := map[*textWord]struct{}{}
			_cec := _faaa._dfec[_edeg]
			for _, _gbec := range _cec {
				if _, _dfbb := _eaaa[_gbec]; _dfbb {
					continue
				}
				for _, _ggce := range _cec {
					if _, _eeag := _eaaa[_ggce]; _eeag {
						continue
					}
					if _ggce != _gbec && _ggce._eabcc == _gbec._eabcc && _c.Abs(_ggce.Llx-_gbec.Llx) < _fgcff && _c.Abs(_ggce.Urx-_gbec.Urx) < _fgcff && _c.Abs(_ggce.Lly-_gbec.Lly) < _fgcff && _c.Abs(_ggce.Ury-_gbec.Ury) < _fgcff {
						_eaaa[_ggce] = struct{}{}
					}
				}
			}
			if len(_eaaa) > 0 {
				_gace := 0
				for _, _cfdf := range _cec {
					if _, _afad := _eaaa[_cfdf]; !_afad {
						_cec[_gace] = _cfdf
						_gace++
					}
				}
				_faaa._dfec[_edeg] = _cec[:len(_cec)-len(_eaaa)]
				if len(_faaa._dfec[_edeg]) == 0 {
					delete(_faaa._dfec, _edeg)
				}
			}
		}
	}
}
func _ffee(_gbbaa _ba.PdfRectangle) rulingKind {
	_afcg := _gbbaa.Width()
	_dbab := _gbbaa.Height()
	if _afcg > _dbab {
		if _afcg >= _deea {
			return _egd
		}
	} else {
		if _dbab >= _deea {
			return _efcg
		}
	}
	return _edc
}
func _abffd(_aeeb, _ecgd int) int {
	if _aeeb < _ecgd {
		return _aeeb
	}
	return _ecgd
}
func (_afac *wordBag) getDepthIdx(_ecgg float64) int {
	_cgeed := _afac.depthIndexes()
	_cffb := _baa(_ecgg)
	if _cffb < _cgeed[0] {
		return _cgeed[0]
	}
	if _cffb > _cgeed[len(_cgeed)-1] {
		return _cgeed[len(_cgeed)-1]
	}
	return _cffb
}
func (_bfgg lineRuling) asRuling() (*ruling, bool) {
	_bded := ruling{_ecddg: _bfgg._acgf, Color: _bfgg.Color, _gdee: _deead}
	switch _bfgg._acgf {
	case _efcg:
		_bded._ccbf = _bfgg.xMean()
		_bded._gcbgg = _c.Min(_bfgg._abeb.Y, _bfgg._cgae.Y)
		_bded._cgeee = _c.Max(_bfgg._abeb.Y, _bfgg._cgae.Y)
	case _egd:
		_bded._ccbf = _bfgg.yMean()
		_bded._gcbgg = _c.Min(_bfgg._abeb.X, _bfgg._cgae.X)
		_bded._cgeee = _c.Max(_bfgg._abeb.X, _bfgg._cgae.X)
	default:
		_f.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _bfgg._acgf)
		return nil, false
	}
	return &_bded, true
}
func _deeec(_eegge []*textMark, _fcbc _ba.PdfRectangle) []*textWord {
	var _bfeaf []*textWord
	var _gaadf *textWord
	if _badb {
		_f.Log.Info("\u006d\u0061\u006beT\u0065\u0078\u0074\u0057\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0064\u0020\u006d\u0061\u0072\u006b\u0073", len(_eegge))
	}
	_dcca := func() {
		if _gaadf != nil {
			_cecdg := _gaadf.computeText()
			if !_cfba(_cecdg) {
				_gaadf._eabcc = _cecdg
				_bfeaf = append(_bfeaf, _gaadf)
				if _badb {
					_f.Log.Info("\u0061\u0064\u0064Ne\u0077\u0057\u006f\u0072\u0064\u003a\u0020\u0025\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", len(_bfeaf)-1, _gaadf.String())
					for _bbdbg, _ecbb := range _gaadf._bdedd {
						_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bbdbg, _ecbb.String())
					}
				}
			}
			_gaadf = nil
		}
	}
	for _, _dgfa := range _eegge {
		if _edfe && _gaadf != nil && len(_gaadf._bdedd) > 0 {
			_baeeg := _gaadf._bdedd[len(_gaadf._bdedd)-1]
			_ecabe, _ceedb := _debcd(_dgfa._abdg)
			_caage, _cfbfe := _debcd(_baeeg._abdg)
			if _ceedb && !_cfbfe && _baeeg.inDiacriticArea(_dgfa) {
				_gaadf.addDiacritic(_ecabe)
				continue
			}
			if _cfbfe && !_ceedb && _dgfa.inDiacriticArea(_baeeg) {
				_gaadf._bdedd = _gaadf._bdedd[:len(_gaadf._bdedd)-1]
				_gaadf.appendMark(_dgfa, _fcbc)
				_gaadf.addDiacritic(_caage)
				continue
			}
		}
		_gccf := _cfba(_dgfa._abdg)
		if _gccf {
			_dcca()
			continue
		}
		if _gaadf == nil && !_gccf {
			_gaadf = _gegf([]*textMark{_dgfa}, _fcbc)
			continue
		}
		_gedb := _gaadf._ggea
		_accee := _c.Abs(_baeb(_fcbc, _dgfa)-_gaadf._addc) / _gedb
		_afbd := _aged(_dgfa, _gaadf) / _gedb
		if _afbd >= _dadg || !(-_bfge <= _afbd && _accee <= _effe) {
			_dcca()
			_gaadf = _gegf([]*textMark{_dgfa}, _fcbc)
			continue
		}
		_gaadf.appendMark(_dgfa, _fcbc)
	}
	_dcca()
	return _bfeaf
}
func _egc(_bdfd *wordBag, _dcbc int) *textLine {
	_ddacc := _bdfd.firstWord(_dcbc)
	_dbe := textLine{PdfRectangle: _ddacc.PdfRectangle, _edfec: _ddacc._ggea, _fcaa: _ddacc._addc}
	_dbe.pullWord(_bdfd, _ddacc, _dcbc)
	return &_dbe
}
func (_bccf *shapesState) newSubPath() {
	_bccf.clearPath()
	if _dfab {
		_f.Log.Info("\u006e\u0065\u0077\u0053\u0075\u0062\u0050\u0061\u0074h\u003a\u0020\u0025\u0073", _bccf)
	}
}

type gridTiling struct {
	_ba.PdfRectangle
	_eade  []float64
	_bcgfa []float64
	_bbca  map[float64]map[float64]gridTile
}

func (_gabbf rulingList) sort() { _dc.Slice(_gabbf, _gabbf.comp) }
func (_fgac *textTable) emptyCompositeRow(_facc int) bool {
	for _bafdc := 0; _bafdc < _fgac._gcead; _bafdc++ {
		if _gebe, _efdf := _fgac._gabfc[_dfee(_bafdc, _facc)]; _efdf {
			if len(_gebe.paraList) > 0 {
				return false
			}
		}
	}
	return true
}

// ExtractTextWithStats works like ExtractText but returns the number of characters in the output
// (`numChars`) and the number of characters that were not decoded (`numMisses`).
func (_daa *Extractor) ExtractTextWithStats() (_dbd string, _dedb int, _ccc int, _agf error) {
	_ddc, _dedb, _ccc, _agf := _daa.ExtractPageText()
	if _agf != nil {
		return "", _dedb, _ccc, _agf
	}
	return _ddc.Text(), _dedb, _ccc, nil
}
func (_abgc paraList) reorder(_dacd []int) {
	_bgbde := make(paraList, len(_abgc))
	for _edggb, _ddeecc := range _dacd {
		_bgbde[_edggb] = _abgc[_ddeecc]
	}
	copy(_abgc, _bgbde)
}

// String returns a string describing `ma`.
func (_cgc TextMarkArray) String() string {
	_fdadg := len(_cgc._afdf)
	if _fdadg == 0 {
		return "\u0045\u004d\u0050T\u0059"
	}
	_gdab := _cgc._afdf[0]
	_gcf := _cgc._afdf[_fdadg-1]
	return _ca.Sprintf("\u007b\u0054\u0045\u0058\u0054\u004d\u0041\u0052K\u0041\u0052\u0052AY\u003a\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025s\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d%\u0073\u007d", _fdadg, _gdab, _gcf)
}

// String returns a description of `k`.
func (_bbgee rulingKind) String() string {
	_gcde, _agfa := _abeaa[_bbgee]
	if !_agfa {
		return _ca.Sprintf("\u004e\u006ft\u0020\u0061\u0020r\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _bbgee)
	}
	return _gcde
}

type textWord struct {
	_ba.PdfRectangle
	_addc  float64
	_eabcc string
	_bdedd []*textMark
	_ggea  float64
	_dfgac bool
}

func _bbcge(_gaad, _deeed, _eggce, _adgb *textPara) *textTable {
	_bfab := &textTable{_gcead: 2, _ebab: 2, _fecac: make(map[uint64]*textPara, 4)}
	_bfab.put(0, 0, _gaad)
	_bfab.put(1, 0, _deeed)
	_bfab.put(0, 1, _eggce)
	_bfab.put(1, 1, _adgb)
	return _bfab
}
func (_bcge *textPara) writeCellText(_gddf _a.Writer) {
	for _aggg, _bcgf := range _bcge._efaf {
		_ffdbg := _bcgf.text()
		_dfdg := _dedbg && _bcgf.endsInHyphen() && _aggg != len(_bcge._efaf)-1
		if _dfdg {
			_ffdbg = _cbee(_ffdbg)
		}
		_gddf.Write([]byte(_ffdbg))
		if !(_dfdg || _aggg == len(_bcge._efaf)-1) {
			_gddf.Write([]byte(_aafg(_bcgf._fcaa, _bcge._efaf[_aggg+1]._fcaa)))
		}
	}
}
func (_eceb *wordBag) absorb(_dcfc *wordBag) {
	_aagb := _dcfc.makeRemovals()
	for _gcbf, _fgbe := range _dcfc._dfec {
		for _, _dced := range _fgbe {
			_eceb.pullWord(_dced, _gcbf, _aagb)
		}
	}
	_dcfc.applyRemovals(_aagb)
}

// ExtractText processes and extracts all text data in content streams and returns as a string.
// It takes into account character encodings in the PDF file, which are decoded by
// CharcodeBytesToUnicode.
// Characters that can't be decoded are replaced with MissingCodeRune ('\ufffd' = �).
func (_cde *Extractor) ExtractText() (string, error) {
	_fcg, _, _, _aa := _cde.ExtractTextWithStats()
	return _fcg, _aa
}

// String returns a description of `tm`.
func (_gcdg *textMark) String() string {
	return _ca.Sprintf("\u0025\u002e\u0032f \u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\u0022\u0025\u0073\u0022", _gcdg.PdfRectangle, _gcdg._dfca, _gcdg._abdg)
}

const (
	_edc rulingKind = iota
	_egd
	_efcg
)

func (_cdfgg *textTable) reduceTiling(_ffcbf gridTiling, _edccg float64) *textTable {
	_baba := make([]int, 0, _cdfgg._ebab)
	_aeef := make([]int, 0, _cdfgg._gcead)
	_dfabde := _ffcbf._eade
	_bffde := _ffcbf._bcgfa
	for _cfbb := 0; _cfbb < _cdfgg._ebab; _cfbb++ {
		_dcba := _cfbb > 0 && _c.Abs(_bffde[_cfbb-1]-_bffde[_cfbb]) < _edccg && _cdfgg.emptyCompositeRow(_cfbb)
		if !_dcba {
			_baba = append(_baba, _cfbb)
		}
	}
	for _gggc := 0; _gggc < _cdfgg._gcead; _gggc++ {
		_cdef := _gggc < _cdfgg._gcead-1 && _c.Abs(_dfabde[_gggc+1]-_dfabde[_gggc]) < _edccg && _cdfgg.emptyCompositeColumn(_gggc)
		if !_cdef {
			_aeef = append(_aeef, _gggc)
		}
	}
	if len(_baba) == _cdfgg._ebab && len(_aeef) == _cdfgg._gcead {
		return _cdfgg
	}
	_edbde := textTable{_deecf: _cdfgg._deecf, _gcead: len(_aeef), _ebab: len(_baba), _gabfc: make(map[uint64]compositeCell, len(_aeef)*len(_baba))}
	if _eca {
		_f.Log.Info("\u0072\u0065\u0064\u0075c\u0065\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0025d\u0078%\u0064\u0020\u002d\u003e\u0020\u0025\u0064x\u0025\u0064", _cdfgg._gcead, _cdfgg._ebab, len(_aeef), len(_baba))
		_f.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _aeef)
		_f.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _baba)
	}
	for _debae, _cgfaa := range _baba {
		for _bbef, _begdeg := range _aeef {
			_affe, _defg := _cdfgg.getComposite(_begdeg, _cgfaa)
			if len(_affe) == 0 {
				continue
			}
			if _eca {
				_ca.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _bbef, _debae, _begdeg, _cgfaa, _bbfd(_affe.merge().text(), 50))
			}
			_edbde.putComposite(_bbef, _debae, _affe, _defg)
		}
	}
	return &_edbde
}
func _dadb(_bage, _fgdab *textPara) bool {
	if _bage._bgea || _fgdab._bgea {
		return true
	}
	return _bdcfc(_bage.depth() - _fgdab.depth())
}
func _ddbba(_dgbbd map[float64]map[float64]gridTile) []float64 {
	_gdde := make([]float64, 0, len(_dgbbd))
	for _aacgg := range _dgbbd {
		_gdde = append(_gdde, _aacgg)
	}
	_dc.Float64s(_gdde)
	_gffee := len(_gdde)
	for _ccdde := 0; _ccdde < _gffee/2; _ccdde++ {
		_gdde[_ccdde], _gdde[_gffee-1-_ccdde] = _gdde[_gffee-1-_ccdde], _gdde[_ccdde]
	}
	return _gdde
}
func (_gfgcf *textTable) putComposite(_cbafc, _ddeg int, _ffef paraList, _feag _ba.PdfRectangle) {
	if len(_ffef) == 0 {
		_f.Log.Error("\u0074\u0065xt\u0054\u0061\u0062l\u0065\u0029\u0020\u0070utC\u006fmp\u006f\u0073\u0069\u0074\u0065\u003a\u0020em\u0070\u0074\u0079\u0020\u0070\u0061\u0072a\u0073")
		return
	}
	_bgbgb := compositeCell{PdfRectangle: _feag, paraList: _ffef}
	if _eca {
		_ca.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0070\u0075\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u003c\u002d\u0025\u0073\u000a", _cbafc, _ddeg, _bgbgb.String())
	}
	_bgbgb.updateBBox()
	_gfgcf._gabfc[_dfee(_cbafc, _ddeg)] = _bgbgb
}
func _badef(_ccab *PageText) error {
	return nil
}
func (_ddbeg rulingList) asTiling() gridTiling {
	if _acgbc {
		_f.Log.Info("r\u0075\u006ci\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0063s\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u002b\u002b\u002b\u0020\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_ddbeg))
	}
	for _dfcc, _bbea := range _ddbeg[1:] {
		_aegd := _ddbeg[_dfcc]
		if _aegd.alignsPrimary(_bbea) && _aegd.alignsSec(_bbea) {
			_f.Log.Error("a\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0044\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065 \u0072\u0075\u006c\u0069\u006e\u0067\u0073\u002e\u000a\u0009v=\u0025\u0073\u000a\t\u0077=\u0025\u0073", _bbea, _aegd)
		}
	}
	_ddbeg.sortStrict()
	_ddbeg.log("\u0073n\u0061\u0070\u0070\u0065\u0064")
	_efca, _defcd := _ddbeg.vertsHorzs()
	_accd := _efca.primaries()
	_cgcbd := _defcd.primaries()
	_abddf := len(_accd) - 1
	_egdg := len(_cgcbd) - 1
	if _abddf == 0 || _egdg == 0 {
		return gridTiling{}
	}
	_cfe := _ba.PdfRectangle{Llx: _accd[0], Urx: _accd[_abddf], Lly: _cgcbd[0], Ury: _cgcbd[_egdg]}
	if _acgbc {
		_f.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0072\u0074s=\u0025\u0064", len(_efca))
		for _aece, _ggfc := range _efca {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _aece, _ggfc)
		}
		_f.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0068\u006f\u0072\u007as=\u0025\u0064", len(_defcd))
		for _bdebd, _fffag := range _defcd {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bdebd, _fffag)
		}
		_f.Log.Info("\u0072\u0075\u006c\u0069\u006eg\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067:\u0020\u0020\u0077\u0078\u0068\u003d\u0025\u0064\u0078\u0025\u0064\u000a\u0009\u006c\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a\u0009\u006c\u006c\u0079\u003d\u0025\u002e\u0032f", _abddf, _egdg, _accd, _cgcbd)
	}
	_adef := make([]gridTile, _abddf*_egdg)
	for _gdfgd := _egdg - 1; _gdfgd >= 0; _gdfgd-- {
		_cdfag := _cgcbd[_gdfgd]
		_acgec := _cgcbd[_gdfgd+1]
		for _ebgb := 0; _ebgb < _abddf; _ebgb++ {
			_ceaed := _accd[_ebgb]
			_cbbe := _accd[_ebgb+1]
			_bffbf := _efca.findPrimSec(_ceaed, _cdfag)
			_dbbgf := _efca.findPrimSec(_cbbe, _cdfag)
			_fgae := _defcd.findPrimSec(_cdfag, _ceaed)
			_gdbb := _defcd.findPrimSec(_acgec, _ceaed)
			_bcaee := _ba.PdfRectangle{Llx: _ceaed, Urx: _cbbe, Lly: _cdfag, Ury: _acgec}
			_cbga := _gafbc(_bcaee, _bffbf, _dbbgf, _fgae, _gdbb)
			_adef[_gdfgd*_abddf+_ebgb] = _cbga
			if _acgbc {
				_ca.Printf("\u0020\u0020\u0078\u003d\u0025\u0032\u0064\u0020\u0079\u003d\u0025\u0032\u0064\u003a\u0020%\u0073 \u0025\u0036\u002e\u0032\u0066\u0020\u0078\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _ebgb, _gdfgd, _cbga.String(), _cbga.Width(), _cbga.Height())
			}
		}
	}
	if _acgbc {
		_f.Log.Info("r\u0075\u006c\u0069\u006e\u0067\u004c\u0069\u0073\u0074.\u0061\u0073\u0054\u0069\u006c\u0069\u006eg:\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0020\u0068\u006f\u0072\u0069zo\u006e\u0074a\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _cfe)
	}
	_ggebc := make([]map[float64]gridTile, _egdg)
	for _bfbc := _egdg - 1; _bfbc >= 0; _bfbc-- {
		if _acgbc {
			_ca.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _bfbc)
		}
		_ggebc[_bfbc] = make(map[float64]gridTile, _abddf)
		for _efe := 0; _efe < _abddf; _efe++ {
			_ccgc := _adef[_bfbc*_abddf+_efe]
			if _acgbc {
				_ca.Printf("\u0020\u0020\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _efe, _ccgc)
			}
			if !_ccgc._fdge {
				continue
			}
			_adgaf := _efe
			for _ecfb := _efe + 1; !_ccgc._abbbb && _ecfb < _abddf; _ecfb++ {
				_eeef := _adef[_bfbc*_abddf+_ecfb]
				_ccgc.Urx = _eeef.Urx
				_ccgc._aebbe = _ccgc._aebbe || _eeef._aebbe
				_ccgc._edcg = _ccgc._edcg || _eeef._edcg
				_ccgc._abbbb = _eeef._abbbb
				if _acgbc {
					_ca.Printf("\u0020 \u0020%\u0034\u0064\u003a\u0020\u0025s\u0020\u2192 \u0025\u0073\u000a", _ecfb, _eeef, _ccgc)
				}
				_adgaf = _ecfb
			}
			if _acgbc {
				_ca.Printf(" \u0020 \u0025\u0032\u0064\u0020\u002d\u0020\u0025\u0032d\u0020\u2192\u0020\u0025s\n", _efe, _adgaf, _ccgc)
			}
			_efe = _adgaf
			_ggebc[_bfbc][_ccgc.Llx] = _ccgc
		}
	}
	_cegc := make(map[float64]map[float64]gridTile, _egdg)
	_bfggf := make(map[float64]map[float64]struct{}, _egdg)
	for _dgcf := _egdg - 1; _dgcf >= 0; _dgcf-- {
		_agcd := _adef[_dgcf*_abddf].Lly
		_cegc[_agcd] = make(map[float64]gridTile, _abddf)
		_bfggf[_agcd] = make(map[float64]struct{}, _abddf)
	}
	if _acgbc {
		_f.Log.Info("\u0072u\u006c\u0069n\u0067\u004c\u0069s\u0074\u002e\u0061\u0073\u0054\u0069\u006ci\u006e\u0067\u003a\u0020\u0063\u006fa\u006c\u0065\u0073\u0063\u0065\u0020\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _cfe)
	}
	for _dcbcc := _egdg - 1; _dcbcc >= 0; _dcbcc-- {
		_ebbd := _adef[_dcbcc*_abddf].Lly
		_dgaa := _ggebc[_dcbcc]
		if _acgbc {
			_ca.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _dcbcc)
		}
		for _, _ccdb := range _acff(_dgaa) {
			if _, _gdefa := _bfggf[_ebbd][_ccdb]; _gdefa {
				continue
			}
			_aaada := _dgaa[_ccdb]
			if _acgbc {
				_ca.Printf(" \u0020\u0020\u0020\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _aaada.String())
			}
			for _bcgg := _dcbcc - 1; _bcgg >= 0; _bcgg-- {
				if _aaada._edcg {
					break
				}
				_decca := _ggebc[_bcgg]
				_acbec, _gffe := _decca[_ccdb]
				if !_gffe {
					break
				}
				if _acbec.Urx != _aaada.Urx {
					break
				}
				_aaada._edcg = _acbec._edcg
				_aaada.Lly = _acbec.Lly
				if _acgbc {
					_ca.Printf("\u0020\u0020\u0020\u0020  \u0020\u0020\u0076\u003d\u0025\u0073\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _acbec.String(), _aaada.String())
				}
				_bfggf[_acbec.Lly][_acbec.Llx] = struct{}{}
			}
			if _dcbcc == 0 {
				_aaada._edcg = true
			}
			if _aaada.complete() {
				_cegc[_ebbd][_ccdb] = _aaada
			}
		}
	}
	_deac := gridTiling{PdfRectangle: _cfe, _eade: _cdbce(_cegc), _bcgfa: _ddbba(_cegc), _bbca: _cegc}
	_deac.log("\u0043r\u0065\u0061\u0074\u0065\u0064")
	return _deac
}

var _cd = false

func (_cfea *textWord) addDiacritic(_ddca string) {
	_gbaaa := _cfea._bdedd[len(_cfea._bdedd)-1]
	_gbaaa._abdg += _ddca
	_gbaaa._abdg = _dce.NFKC.String(_gbaaa._abdg)
}
func _bbfd(_gcegg string, _gaag int) string {
	if len(_gcegg) < _gaag {
		return _gcegg
	}
	return _gcegg[:_gaag]
}
func _debcd(_adeff string) (string, bool) {
	_adeb := []rune(_adeff)
	if len(_adeb) != 1 {
		return "", false
	}
	_efdfe, _agcef := _bbfcf[_adeb[0]]
	return _efdfe, _agcef
}
func (_dgad rectRuling) checkWidth(_dcfdf, _bcbg float64) (float64, bool) {
	_cdecd := _bcbg - _dcfdf
	_cfab := _cdecd <= _gedc
	return _cdecd, _cfab
}
func (_cedcf paraList) yNeighbours(_eggd float64) map[*textPara][]int {
	_ebff := make([]event, 2*len(_cedcf))
	if _eggd == 0 {
		for _ffaef, _bbac := range _cedcf {
			_ebff[2*_ffaef] = event{_bbac.Lly, true, _ffaef}
			_ebff[2*_ffaef+1] = event{_bbac.Ury, false, _ffaef}
		}
	} else {
		for _cdcdg, _cabe := range _cedcf {
			_ebff[2*_cdcdg] = event{_cabe.Lly - _eggd*_cabe.fontsize(), true, _cdcdg}
			_ebff[2*_cdcdg+1] = event{_cabe.Ury + _eggd*_cabe.fontsize(), false, _cdcdg}
		}
	}
	return _cedcf.eventNeighbours(_ebff)
}
func (_ebfcde paraList) findTextTables() []*textTable {
	var _beedd []*textTable
	for _, _ceagd := range _ebfcde {
		if _ceagd.taken() || _ceagd.Width() == 0 {
			continue
		}
		_cefb := _ceagd.isAtom()
		if _cefb == nil {
			continue
		}
		_cefb.growTable()
		if _cefb._gcead*_cefb._ebab < _cce {
			continue
		}
		_cefb.markCells()
		_cefb.log("\u0067\u0072\u006fw\u006e")
		_beedd = append(_beedd, _cefb)
	}
	return _beedd
}
func (_cdfa paraList) toTextMarks() []TextMark {
	_bffc := 0
	var _egea []TextMark
	for _fdcfe, _cgbd := range _cdfa {
		if _cgbd._bgea {
			continue
		}
		_dac := _cgbd.toTextMarks(&_bffc)
		_egea = append(_egea, _dac...)
		if _fdcfe != len(_cdfa)-1 {
			if _dadb(_cgbd, _cdfa[_fdcfe+1]) {
				_egea = _cdae(_egea, &_bffc, "\u0020")
			} else {
				_egea = _cdae(_egea, &_bffc, "\u000a")
				_egea = _cdae(_egea, &_bffc, "\u000a")
			}
		}
	}
	_egea = _cdae(_egea, &_bffc, "\u000a")
	_egea = _cdae(_egea, &_bffc, "\u000a")
	return _egea
}
func _fbcgb(_cagf []float64, _fecbg, _fedff float64) []float64 {
	_bafb, _fcee := _fecbg, _fedff
	if _fcee < _bafb {
		_bafb, _fcee = _fcee, _bafb
	}
	_acgfe := make([]float64, 0, len(_cagf)+2)
	_acgfe = append(_acgfe, _fecbg)
	for _, _daad := range _cagf {
		if _daad <= _bafb {
			continue
		} else if _daad >= _fcee {
			break
		}
		_acgfe = append(_acgfe, _daad)
	}
	_acgfe = append(_acgfe, _fedff)
	return _acgfe
}
func (_ggcc rulingList) bbox() _ba.PdfRectangle {
	var _badba _ba.PdfRectangle
	if len(_ggcc) == 0 {
		_f.Log.Error("r\u0075\u006c\u0069\u006e\u0067\u004ci\u0073\u0074\u002e\u0062\u0062\u006f\u0078\u003a\u0020n\u006f\u0020\u0072u\u006ci\u006e\u0067\u0073")
		return _ba.PdfRectangle{}
	}
	if _ggcc[0]._ecddg == _egd {
		_badba.Llx, _badba.Urx = _ggcc.secMinMax()
		_badba.Lly, _badba.Ury = _ggcc.primMinMax()
	} else {
		_badba.Llx, _badba.Urx = _ggcc.primMinMax()
		_badba.Lly, _badba.Ury = _ggcc.secMinMax()
	}
	return _badba
}
func (_acae *shapesState) closePath() {
	if _acae._debd {
		_acae._eagg = append(_acae._eagg, _bcecc(_acae._ebeg))
		_acae._debd = false
	} else if len(_acae._eagg) == 0 {
		if _dfab {
			_f.Log.Debug("\u0063\u006c\u006f\u0073eP\u0061\u0074\u0068\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u0020\u0070\u0061t\u0068")
		}
		_acae._debd = false
		return
	}
	_acae._eagg[len(_acae._eagg)-1].close()
	if _dfab {
		_f.Log.Info("\u0063\u006c\u006f\u0073\u0065\u0050\u0061\u0074\u0068\u003a\u0020\u0025\u0073", _acae)
	}
}
func (_bbfb paraList) readBefore(_dddg []int, _bdbg, _fdde int) bool {
	_aafa, _bgae := _bbfb[_bdbg], _bbfb[_fdde]
	if _edbeb(_aafa, _bgae) && _aafa.Lly > _bgae.Lly {
		return true
	}
	if !(_aafa._dabb.Urx < _bgae._dabb.Llx) {
		return false
	}
	_aggb, _eac := _aafa.Lly, _bgae.Lly
	if _aggb > _eac {
		_eac, _aggb = _aggb, _eac
	}
	_gecb := _c.Max(_aafa._dabb.Llx, _bgae._dabb.Llx)
	_dbgf := _c.Min(_aafa._dabb.Urx, _bgae._dabb.Urx)
	_ddce := _bbfb.llyRange(_dddg, _aggb, _eac)
	for _, _aceed := range _ddce {
		if _aceed == _bdbg || _aceed == _fdde {
			continue
		}
		_acfd := _bbfb[_aceed]
		if _acfd._dabb.Llx <= _dbgf && _gecb <= _acfd._dabb.Urx {
			return false
		}
	}
	return true
}

var _cagg = map[markKind]string{_deead: "\u0073\u0074\u0072\u006f\u006b\u0065", _ddbe: "\u0066\u0069\u006c\u006c", _ggbdb: "\u0061u\u0067\u006d\u0065\u006e\u0074"}

func _gafb(_aadc, _caagg _bc.Point) rulingKind {
	_agbd := _c.Abs(_aadc.X - _caagg.X)
	_fead := _c.Abs(_aadc.Y - _caagg.Y)
	return _ebegb(_agbd, _fead, _deea)
}
func _ccdd(_gaeb *wordBag, _gdb *textWord, _cgdg float64) bool {
	return _gdb.Llx < _gaeb.Urx+_cgdg && _gaeb.Llx-_cgdg < _gdb.Urx
}
func _ccbd(_beeb []pathSection) {
	if _ebf < 0.0 {
		return
	}
	if _bfg {
		_f.Log.Info("\u0067\u0072\u0061\u006e\u0075\u006c\u0061\u0072\u0069\u007a\u0065\u003a\u0020\u0025\u0064 \u0073u\u0062\u0070\u0061\u0074\u0068\u0020\u0073\u0065\u0063\u0074\u0069\u006f\u006e\u0073", len(_beeb))
	}
	for _eeeee, _gdff := range _beeb {
		for _abdec, _ebee := range _gdff._cag {
			for _gcfb, _gaba := range _ebee._caf {
				_ebee._caf[_gcfb] = _bc.Point{X: _bbdg(_gaba.X), Y: _bbdg(_gaba.Y)}
				if _bfg {
					_efdgg := _ebee._caf[_gcfb]
					if !_dadd(_gaba, _efdgg) {
						_ebgd := _bc.Point{X: _efdgg.X - _gaba.X, Y: _efdgg.Y - _gaba.Y}
						_ca.Printf("\u0025\u0034d \u002d\u0020\u00254\u0064\u0020\u002d\u0020%4d\u003a %\u002e\u0032\u0066\u0020\u2192\u0020\u0025.2\u0066\u0020\u0028\u0025\u0067\u0029\u000a", _eeeee, _abdec, _gcfb, _gaba, _efdgg, _ebgd)
					}
				}
			}
		}
	}
}
func (_cbbc *wordBag) removeWord(_edb *textWord, _ddac int) {
	_fcca := _cbbc._dfec[_ddac]
	_fcca = _ebdga(_fcca, _edb)
	if len(_fcca) == 0 {
		delete(_cbbc._dfec, _ddac)
	} else {
		_cbbc._dfec[_ddac] = _fcca
	}
}

const _abee = 1.0 / 1000.0

var _edcgb = _d.MustCompile("\u005e\u005c\u0073\u002a\u0028\u005c\u0064\u002b\u005c\u002e\u003f|\u005b\u0049\u0069\u0076\u005d\u002b\u0029\u005c\u0073\u002a\\\u0029\u003f\u0024")

func (_bfff *subpath) last() _bc.Point         { return _bfff._caf[len(_bfff._caf)-1] }
func (_feed *textPara) bbox() _ba.PdfRectangle { return _feed.PdfRectangle }
func (_ddacd *compositeCell) updateBBox() {
	for _, _egfe := range _ddacd.paraList {
		_ddacd.PdfRectangle = _cafce(_ddacd.PdfRectangle, _egfe.PdfRectangle)
	}
}
func (_ecggg rulingList) augmentGrid() (rulingList, rulingList) {
	_agfe, _fgbc := _ecggg.vertsHorzs()
	if len(_agfe) == 0 || len(_fgbc) == 0 {
		return _agfe, _fgbc
	}
	_ceed, _dbce := _agfe, _fgbc
	_ffabe := _agfe.bbox()
	_adaga := _fgbc.bbox()
	if _bfg {
		_f.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0056\u003d\u0025\u0036\u002e\u0032\u0066", _ffabe)
		_f.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0048\u003d\u0025\u0036\u002e\u0032\u0066", _adaga)
	}
	var _ccdge, _acbff, _cbgd, _dfbc *ruling
	if _adaga.Llx < _ffabe.Llx-_gdae {
		_ccdge = &ruling{_gdee: _ggbdb, _ecddg: _efcg, _ccbf: _adaga.Llx, _gcbgg: _ffabe.Lly, _cgeee: _ffabe.Ury}
		_agfe = append(rulingList{_ccdge}, _agfe...)
	}
	if _adaga.Urx > _ffabe.Urx+_gdae {
		_acbff = &ruling{_gdee: _ggbdb, _ecddg: _efcg, _ccbf: _adaga.Urx, _gcbgg: _ffabe.Lly, _cgeee: _ffabe.Ury}
		_agfe = append(_agfe, _acbff)
	}
	if _ffabe.Lly < _adaga.Lly-_gdae {
		_cbgd = &ruling{_gdee: _ggbdb, _ecddg: _egd, _ccbf: _ffabe.Lly, _gcbgg: _adaga.Llx, _cgeee: _adaga.Urx}
		_fgbc = append(rulingList{_cbgd}, _fgbc...)
	}
	if _ffabe.Ury > _adaga.Ury+_gdae {
		_dfbc = &ruling{_gdee: _ggbdb, _ecddg: _egd, _ccbf: _ffabe.Ury, _gcbgg: _adaga.Llx, _cgeee: _adaga.Urx}
		_fgbc = append(_fgbc, _dfbc)
	}
	if len(_agfe)+len(_fgbc) == len(_ecggg) {
		return _ceed, _dbce
	}
	_gdcf := append(_agfe, _fgbc...)
	_ecggg.log("u\u006e\u0061\u0075\u0067\u006d\u0065\u006e\u0074\u0065\u0064")
	_gdcf.log("\u0061u\u0067\u006d\u0065\u006e\u0074\u0065d")
	return _agfe, _fgbc
}
func (_aec *shapesState) moveTo(_fdfd, _bcc float64) {
	_aec._debd = true
	_aec._ebeg = _aec.devicePoint(_fdfd, _bcc)
	if _dfab {
		_f.Log.Info("\u006d\u006fv\u0065\u0054\u006f\u003a\u0020\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u003d%.\u0032\u0066", _fdfd, _bcc, _aec._ebeg)
	}
}
func (_eceg *textObject) nextLine() { _eceg.moveLP(0, -_eceg._gdcg._dgbbf) }
func _fcga(_gbbg []pathSection) rulingList {
	_ccbd(_gbbg)
	if _bfg {
		_f.Log.Info("\u006d\u0061k\u0065\u0053\u0074\u0072\u006f\u006b\u0065\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0064\u0020\u0073\u0074\u0072ok\u0065\u0073", len(_gbbg))
	}
	var _dggdb rulingList
	for _, _egcf := range _gbbg {
		for _, _cdcbc := range _egcf._cag {
			if len(_cdcbc._caf) < 2 {
				continue
			}
			_cffd := _cdcbc._caf[0]
			for _, _ffcba := range _cdcbc._caf[1:] {
				if _cgcb, _caddba := _befg(_cffd, _ffcba, _egcf.Color); _caddba {
					_dggdb = append(_dggdb, _cgcb)
				}
				_cffd = _ffcba
			}
		}
	}
	if _bfg {
		_f.Log.Info("m\u0061\u006b\u0065\u0053tr\u006fk\u0065\u0052\u0075\u006c\u0069n\u0067\u0073\u003a\u0020\u0025\u0073", _dggdb)
	}
	return _dggdb
}

type textMark struct {
	_ba.PdfRectangle
	_bdba  int
	_abdg  string
	_deee  string
	_ddbd  *_ba.PdfFont
	_dfca  float64
	_geca  float64
	_gbaa  _bc.Matrix
	_ffad  _bc.Point
	_fdbg  _ba.PdfRectangle
	_bbcg  _af.Color
	_fbbcf _af.Color
}

func (_edeb rulingList) primaries() []float64 {
	_eabda := make(map[float64]struct{}, len(_edeb))
	for _, _ecabb := range _edeb {
		_eabda[_ecabb._ccbf] = struct{}{}
	}
	_bdfgb := make([]float64, len(_eabda))
	_afga := 0
	for _gbffg := range _eabda {
		_bdfgb[_afga] = _gbffg
		_afga++
	}
	_dc.Float64s(_bdfgb)
	return _bdfgb
}

type textObject struct {
	_acca *Extractor
	_dgda *_ba.PdfPageResources
	_fdcd _da.GraphicsState
	_gdcg *textState
	_gag  *stateStack
	_ddfg _bc.Matrix
	_bdc  _bc.Matrix
	_adda []*textMark
	_fged bool
}

func (_feaf *textObject) moveText(_aae, _gba float64) { _feaf.moveLP(_aae, _gba) }
func _egcb(_dfcad []int) []int {
	_cbacd := make([]int, len(_dfcad))
	for _ebce, _cdcd := range _dfcad {
		_cbacd[len(_dfcad)-1-_ebce] = _cdcd
	}
	return _cbacd
}

type pathSection struct {
	_cag []*subpath
	_af.Color
}

func (_abeg compositeCell) hasLines(_ccgf []*textLine) bool {
	for _bcbf, _fccg := range _ccgf {
		_fbgfe := _ffdd(_abeg.PdfRectangle, _fccg.PdfRectangle)
		if _eca {
			_ca.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u005e\u005e\u005e\u0069\u006e\u0074\u0065\u0072\u0073e\u0063t\u0073\u003d\u0025\u0074\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u000a", _fbgfe, _bcbf, len(_ccgf))
			_ca.Printf("\u0020\u0020\u0020\u0020  \u005e\u005e\u005e\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025s\u000a", _abeg)
			_ca.Printf("\u0020 \u0020 \u0020\u0020\u0020\u006c\u0069\u006e\u0065\u003d\u0025\u0073\u000a", _fccg)
		}
		if _fbgfe {
			return true
		}
	}
	return false
}
func (_baga *textTable) compositeRowCorridors() map[int][]float64 {
	_cdedf := make(map[int][]float64, _baga._ebab)
	if _eca {
		_f.Log.Info("c\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0052\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073:\u0020h\u003d\u0025\u0064", _baga._ebab)
	}
	for _ccbbc := 1; _ccbbc < _baga._ebab; _ccbbc++ {
		var _eabf []compositeCell
		for _abcf := 0; _abcf < _baga._gcead; _abcf++ {
			if _feadc, _efed := _baga._gabfc[_dfee(_abcf, _ccbbc)]; _efed {
				_eabf = append(_eabf, _feadc)
			}
		}
		if len(_eabf) == 0 {
			continue
		}
		_gcbe := _fbee(_eabf)
		_cdedf[_ccbbc] = _gcbe
		if _eca {
			_ca.Printf("\u0020\u0020\u0020\u0025\u0032\u0064\u003a\u0020\u00256\u002e\u0032\u0066\u000a", _ccbbc, _gcbe)
		}
	}
	return _cdedf
}
func (_faed lineRuling) xMean() float64 { return 0.5 * (_faed._abeb.X + _faed._cgae.X) }

// Tables returns the tables extracted from the page.
func (_bea PageText) Tables() []TextTable {
	if _eca {
		_f.Log.Info("\u0054\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_bea._defb))
	}
	return _bea._defb
}
func (_acd *textObject) setTextMatrix(_feac []float64) {
	if len(_feac) != 6 {
		_f.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006c\u0065\u006e\u0028\u0066\u0029\u0020\u0021\u003d\u0020\u0036\u0020\u0028\u0025\u0064\u0029", len(_feac))
		return
	}
	_bff, _accc, _gbd, _ddf, _ebb, _eff := _feac[0], _feac[1], _feac[2], _feac[3], _feac[4], _feac[5]
	_acd._ddfg = _bc.NewMatrix(_bff, _accc, _gbd, _ddf, _ebb, _eff)
	_acd._bdc = _acd._ddfg
}

// String returns a description of `v`.
func (_bcaff *ruling) String() string {
	if _bcaff._ecddg == _edc {
		return "\u004e\u004f\u0054\u0020\u0052\u0055\u004c\u0049\u004e\u0047"
	}
	_ddae, _bebf := "\u0078", "\u0079"
	if _bcaff._ecddg == _egd {
		_ddae, _bebf = "\u0079", "\u0078"
	}
	_gbgcc := ""
	if _bcaff._eeca != 0.0 {
		_gbgcc = _ca.Sprintf(" \u0077\u0069\u0064\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _bcaff._eeca)
	}
	return _ca.Sprintf("\u0025\u00310\u0073\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0020\u0025\u0076\u0025\u0073", _bcaff._ecddg, _ddae, _bcaff._ccbf, _bebf, _bcaff._gcbgg, _bcaff._cgeee, _bcaff._cgeee-_bcaff._gcbgg, _bcaff._gdee, _bcaff.Color, _gbgcc)
}
func (_dbebf *textTable) reduce() *textTable {
	_bcea := make([]int, 0, _dbebf._ebab)
	_bbbdd := make([]int, 0, _dbebf._gcead)
	for _gcfe := 0; _gcfe < _dbebf._ebab; _gcfe++ {
		if !_dbebf.emptyCompositeRow(_gcfe) {
			_bcea = append(_bcea, _gcfe)
		}
	}
	for _gdfa := 0; _gdfa < _dbebf._gcead; _gdfa++ {
		if !_dbebf.emptyCompositeColumn(_gdfa) {
			_bbbdd = append(_bbbdd, _gdfa)
		}
	}
	if len(_bcea) == _dbebf._ebab && len(_bbbdd) == _dbebf._gcead {
		return _dbebf
	}
	_bcdeg := textTable{_deecf: _dbebf._deecf, _gcead: len(_bbbdd), _ebab: len(_bcea), _fecac: make(map[uint64]*textPara, len(_bbbdd)*len(_bcea))}
	if _eca {
		_f.Log.Info("\u0072\u0065\u0064\u0075ce\u003a\u0020\u0025\u0064\u0078\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064\u0078%\u0064", _dbebf._gcead, _dbebf._ebab, len(_bbbdd), len(_bcea))
		_f.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _bbbdd)
		_f.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _bcea)
	}
	for _fegd, _fdfg := range _bcea {
		for _agea, _gafgc := range _bbbdd {
			_deef, _eaaad := _dbebf.getComposite(_gafgc, _fdfg)
			if _deef == nil {
				continue
			}
			if _eca {
				_ca.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _agea, _fegd, _gafgc, _fdfg, _bbfd(_deef.merge().text(), 50))
			}
			_bcdeg.putComposite(_agea, _fegd, _deef, _eaaad)
		}
	}
	return &_bcdeg
}
func (_addd *textTable) growTable() {
	_ebbca := func(_edbea paraList) {
		_addd._ebab++
		for _bgbga := 0; _bgbga < _addd._gcead; _bgbga++ {
			_bdg := _edbea[_bgbga]
			_addd.put(_bgbga, _addd._ebab-1, _bdg)
		}
	}
	_dbbgb := func(_bcce paraList) {
		_addd._gcead++
		for _dfbba := 0; _dfbba < _addd._ebab; _dfbba++ {
			_fgdf := _bcce[_dfbba]
			_addd.put(_addd._gcead-1, _dfbba, _fgdf)
		}
	}
	if _abab {
		_addd.log("\u0067r\u006f\u0077\u0054\u0061\u0062\u006ce")
	}
	for _ddace := 0; ; _ddace++ {
		_daace := false
		_cedea := _addd.getDown()
		_ffcd := _addd.getRight()
		if _abab {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _ddace, _addd)
			_ca.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0020\u0064\u006f\u0077\u006e\u003d\u0025\u0073\u000a", _cedea)
			_ca.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0072\u0069\u0067\u0068\u0074\u003d\u0025\u0073\u000a", _ffcd)
		}
		if _cedea != nil && _ffcd != nil {
			_bgec := _cedea[len(_cedea)-1]
			if !_bgec.taken() && _bgec == _ffcd[len(_ffcd)-1] {
				_ebbca(_cedea)
				if _ffcd = _addd.getRight(); _ffcd != nil {
					_dbbgb(_ffcd)
					_addd.put(_addd._gcead-1, _addd._ebab-1, _bgec)
				}
				_daace = true
			}
		}
		if !_daace && _cedea != nil {
			_ebbca(_cedea)
			_daace = true
		}
		if !_daace && _ffcd != nil {
			_dbbgb(_ffcd)
			_daace = true
		}
		if !_daace {
			break
		}
	}
}
func _ccfc(_gcefg []*textMark, _bbe _ba.PdfRectangle, _agdce rulingList, _fggba []gridTiling) paraList {
	_f.Log.Trace("\u006d\u0061\u006b\u0065\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u003a \u0025\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073\u0020\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065=\u0025\u002e\u0032\u0066", len(_gcefg), _bbe)
	if len(_gcefg) == 0 {
		return nil
	}
	_cebe := _deeec(_gcefg, _bbe)
	if len(_cebe) == 0 {
		return nil
	}
	_agdce.log("\u006d\u0061\u006be\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065")
	_dfcb, _ebfd := _agdce.vertsHorzs()
	_aefac := _adcd(_cebe, _bbe.Ury, _dfcb, _ebfd)
	_dagd := _fcba(_aefac, _bbe.Ury, _dfcb, _ebfd)
	_dagd = _dbdb(_dagd)
	_bddg := make(paraList, 0, len(_dagd))
	for _, _gcaag := range _dagd {
		_ggec := _gcaag.arrangeText()
		if _ggec != nil {
			_bddg = append(_bddg, _ggec)
		}
	}
	if len(_bddg) >= _cce {
		_bddg = _bddg.extractTables(_fggba)
	}
	_bddg.sortReadingOrder()
	_bddg.log("\u0073\u006f\u0072te\u0064\u0020\u0069\u006e\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0064\u0065\u0072")
	return _bddg
}
func (_dddda rulingList) isActualGrid() (rulingList, bool) {
	_fdef, _ccee := _dddda.augmentGrid()
	if !(len(_fdef) >= _efb+1 && len(_ccee) >= _fffd+1) {
		if _bfg {
			_f.Log.Info("\u0069s\u0041\u0063t\u0075\u0061\u006c\u0047r\u0069\u0064\u003a \u004e\u006f\u0074\u0020\u0061\u006c\u0069\u0067\u006eed\u002e\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u003c\u0020\u0025d\u0020\u0078 \u0025\u0064", len(_fdef), len(_ccee), _efb+1, _fffd+1)
		}
		return nil, false
	}
	if _bfg {
		_f.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075a\u006c\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025t\u0020\u0026\u0020\u0025\u0074\u0020\u2192 \u0025\u0074", _dddda, len(_fdef) >= 2, len(_ccee) >= 2, len(_fdef) >= 2 && len(_ccee) >= 2)
		for _fbfe, _cede := range _dddda {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0076\u000a", _fbfe, _cede)
		}
	}
	if _dbbg {
		_afg, _adagb := _fdef[0], _fdef[len(_fdef)-1]
		_bcff, _beee := _ccee[0], _ccee[len(_ccee)-1]
		if !(_cbcb(_afg._ccbf-_bcff._gcbgg) && _cbcb(_adagb._ccbf-_bcff._cgeee) && _cbcb(_bcff._ccbf-_afg._cgeee) && _cbcb(_beee._ccbf-_afg._gcbgg)) {
			if _bfg {
				_f.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069d\u003a\u0020\u0020N\u006f\u0074 \u0061\u006c\u0069\u0067\u006e\u0065d\u002e\n\t\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0076\u0031\u003d\u0025\u0073\u000a\u0009\u0068\u0030\u003d\u0025\u0073\u000a\u0009\u0068\u0031\u003d\u0025\u0073", _afg, _adagb, _bcff, _beee)
			}
			return nil, false
		}
	} else {
		if !_fdef.aligned() {
			if _dcdb {
				_f.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0076\u0065\u0072\u0074\u0073\u002e\u0020%\u0064", len(_fdef))
			}
			return nil, false
		}
		if !_ccee.aligned() {
			if _bfg {
				_f.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0068\u006f\u0072\u007a\u0073\u002e\u0020%\u0064", len(_ccee))
			}
			return nil, false
		}
	}
	_becgf := append(_fdef, _ccee...)
	return _becgf, true
}
func (_abbd *shapesState) lineTo(_bcda, _gfee float64) {
	if _dfab {
		_f.Log.Info("\u006c\u0069\u006eeT\u006f\u0028\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0070\u003d\u0025\u002e\u0032\u0066", _bcda, _gfee, _abbd.devicePoint(_bcda, _gfee))
	}
	_abbd.addPoint(_bcda, _gfee)
}

type cachedImage struct {
	_feg *_ba.Image
	_gfe _ba.PdfColorspace
}

func (_eeg *stateStack) empty() bool { return len(*_eeg) == 0 }
func (_bdeb *wordBag) text() string {
	_cgb := _bdeb.allWords()
	_aba := make([]string, len(_cgb))
	for _ebbe, _cgdeb := range _cgb {
		_aba[_ebbe] = _cgdeb._eabcc
	}
	return _ad.Join(_aba, "\u0020")
}
func (_ccb *shapesState) clearPath() {
	_ccb._eagg = nil
	_ccb._debd = false
	if _dfab {
		_f.Log.Info("\u0043\u004c\u0045A\u0052\u003a\u0020\u0073\u0073\u003d\u0025\u0073", _ccb)
	}
}

// Append appends `mark` to the mark array.
func (_degf *TextMarkArray) Append(mark TextMark) { _degf._afdf = append(_degf._afdf, mark) }
func _edbeb(_fegg, _bade *textPara) bool          { return _ccbc(_fegg._dabb, _bade._dabb) }

// String returns a human readable description of `vecs`.
func (_bebg rulingList) String() string {
	if len(_bebg) == 0 {
		return "\u007b \u0045\u004d\u0050\u0054\u0059\u0020}"
	}
	_cbfb, _dgdea := _bebg.vertsHorzs()
	_faff := len(_cbfb)
	_ebefb := len(_dgdea)
	if _faff == 0 || _ebefb == 0 {
		return _ca.Sprintf("\u007b%\u0064\u0020\u0078\u0020\u0025\u0064}", _faff, _ebefb)
	}
	_cfbdg := _ba.PdfRectangle{Llx: _cbfb[0]._ccbf, Urx: _cbfb[_faff-1]._ccbf, Lly: _dgdea[_ebefb-1]._ccbf, Ury: _dgdea[0]._ccbf}
	return _ca.Sprintf("\u007b\u0025d\u0020\u0078\u0020%\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u007d", _faff, _ebefb, _cfbdg)
}
func (_efad *textWord) toTextMarks(_fdgeg *int) []TextMark {
	var _fdggeg []TextMark
	for _, _dacgd := range _efad._bdedd {
		_fdggeg = _bddd(_fdggeg, _fdgeg, _dacgd.ToTextMark())
	}
	return _fdggeg
}
func (_bcgeg paraList) findTables(_cbgfe []gridTiling) []*textTable {
	_bcgeg.addNeighbours()
	_dc.Slice(_bcgeg, func(_dcfca, _dgcd int) bool { return _dgab(_bcgeg[_dcfca], _bcgeg[_dgcd]) < 0 })
	var _geac []*textTable
	if _eege {
		_dbgda := _bcgeg.findGridTables(_cbgfe)
		_geac = append(_geac, _dbgda...)
	}
	if _fcda {
		_fgcd := _bcgeg.findTextTables()
		_geac = append(_geac, _fgcd...)
	}
	return _geac
}
func (_cfge rulingList) intersections() map[int]intSet {
	var _ebfcd, _dfcac []int
	for _gacd, _gcdga := range _cfge {
		switch _gcdga._ecddg {
		case _efcg:
			_ebfcd = append(_ebfcd, _gacd)
		case _egd:
			_dfcac = append(_dfcac, _gacd)
		}
	}
	if len(_ebfcd) < _efb+1 || len(_dfcac) < _fffd+1 {
		return nil
	}
	if len(_ebfcd)+len(_dfcac) > _agdc {
		_f.Log.Debug("\u0069\u006e\u0074\u0065\u0072\u0073e\u0063\u0074\u0069\u006f\u006e\u0073\u003a\u0020\u0054\u004f\u004f\u0020\u004d\u0041\u004e\u0059\u0020\u0072\u0075\u006ci\u006e\u0067\u0073\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u003d\u0020%\u0064 \u0078\u0020\u0025\u0064", len(_cfge), len(_ebfcd), len(_dfcac))
		return nil
	}
	_gceg := make(map[int]intSet, len(_ebfcd)+len(_dfcac))
	for _, _aacg := range _ebfcd {
		for _, _ggdf := range _dfcac {
			if _cfge[_aacg].intersects(_cfge[_ggdf]) {
				if _, _eabae := _gceg[_aacg]; !_eabae {
					_gceg[_aacg] = make(intSet)
				}
				if _, _dgbg := _gceg[_ggdf]; !_dgbg {
					_gceg[_ggdf] = make(intSet)
				}
				_gceg[_aacg].add(_ggdf)
				_gceg[_ggdf].add(_aacg)
			}
		}
	}
	return _gceg
}
func (_daec rulingList) aligned() bool {
	if len(_daec) < 2 {
		return false
	}
	_dacg := make(map[*ruling]int)
	_dacg[_daec[0]] = 0
	for _, _cdcc := range _daec[1:] {
		_bfddcg := false
		for _ebdg := range _dacg {
			if _cdcc.gridIntersecting(_ebdg) {
				_dacg[_ebdg]++
				_bfddcg = true
				break
			}
		}
		if !_bfddcg {
			_dacg[_cdcc] = 0
		}
	}
	_edde := 0
	for _, _eaadf := range _dacg {
		if _eaadf == 0 {
			_edde++
		}
	}
	_fdga := float64(_edde) / float64(len(_daec))
	_eed := _fdga <= 1.0-_cgec
	if _bfg {
		_f.Log.Info("\u0061\u006c\u0069\u0067\u006e\u0065\u0064\u003d\u0025\u0074\u0020\u0075\u006em\u0061\u0074\u0063\u0068\u0065\u0064=\u0025\u002e\u0032\u0066\u003d\u0025\u0064\u002f\u0025\u0064\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _eed, _fdga, _edde, len(_daec), _daec.String())
	}
	return _eed
}

type rulingList []*ruling

func (_gede paraList) tables() []TextTable {
	var _gedfg []TextTable
	if _eca {
		_f.Log.Info("\u0070\u0061\u0072\u0061\u0073\u002e\u0074\u0061\u0062\u006c\u0065\u0073\u003a")
	}
	for _, _agdaa := range _gede {
		_daba := _agdaa._ebcb
		if _daba != nil && _daba.isExportable() {
			_gedfg = append(_gedfg, _daba.toTextTable())
		}
	}
	return _gedfg
}
func (_cecc intSet) add(_acagd int) { _cecc[_acagd] = struct{}{} }

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
	BBox _ba.PdfRectangle

	// Font is the font the text was drawn with.
	Font *_ba.PdfFont

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
	FillColor _af.Color

	// StrokeColor is the stroke color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	StrokeColor _af.Color

	// Orientation is the text orientation
	Orientation int
}

func (_egfa *textTable) put(_dfbed, _cceaa int, _efab *textPara) {
	_egfa._fecac[_dfee(_dfbed, _cceaa)] = _efab
}
func _befg(_abbb, _dfgg _bc.Point, _ffdf _af.Color) (*ruling, bool) {
	_bgadd := lineRuling{_abeb: _abbb, _cgae: _dfgg, _acgf: _gafb(_abbb, _dfgg), Color: _ffdf}
	if _bgadd._acgf == _edc {
		return nil, false
	}
	return _bgadd.asRuling()
}

type compositeCell struct {
	_ba.PdfRectangle
	paraList
}

func (_dabg *shapesState) establishSubpath() *subpath {
	_ecf, _afa := _dabg.lastpointEstablished()
	if !_afa {
		_dabg._eagg = append(_dabg._eagg, _bcecc(_ecf))
	}
	if len(_dabg._eagg) == 0 {
		return nil
	}
	_dabg._debd = false
	return _dabg._eagg[len(_dabg._eagg)-1]
}
func (_gece *textObject) newTextMark(_bdbe string, _cfbf _bc.Matrix, _ebfa _bc.Point, _afef float64, _gdea *_ba.PdfFont, _gggg float64, _ffede, _cbac _af.Color) (textMark, bool) {
	_gfade := _cfbf.Angle()
	_dagcc := _ccea(_gfade, _aagd)
	var _begg float64
	if _dagcc%180 != 90 {
		_begg = _cfbf.ScalingFactorY()
	} else {
		_begg = _cfbf.ScalingFactorX()
	}
	_feb := _fcfd(_cfbf)
	_aedc := _ba.PdfRectangle{Llx: _feb.X, Lly: _feb.Y, Urx: _ebfa.X, Ury: _ebfa.Y}
	switch _dagcc % 360 {
	case 90:
		_aedc.Urx -= _begg
	case 180:
		_aedc.Ury -= _begg
	case 270:
		_aedc.Urx += _begg
	case 0:
		_aedc.Ury += _begg
	default:
		_dagcc = 0
		_aedc.Ury += _begg
	}
	if _aedc.Llx > _aedc.Urx {
		_aedc.Llx, _aedc.Urx = _aedc.Urx, _aedc.Llx
	}
	if _aedc.Lly > _aedc.Ury {
		_aedc.Lly, _aedc.Ury = _aedc.Ury, _aedc.Lly
	}
	_gbgfa := true
	if _gece._acca._gg.Width() > 0 {
		_dfad, _dfegc := _bdaa(_aedc, _gece._acca._gg)
		if !_dfegc {
			_gbgfa = false
			_f.Log.Debug("\u0054\u0065\u0078\u0074\u0020m\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0070a\u0067\u0065\u002e\u0020\u0062\u0062\u006f\u0078\u003d\u0025\u0067\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u0067\u0020\u0074\u0065\u0078\u0074\u003d\u0025q", _aedc, _gece._acca._gg, _bdbe)
		}
		_aedc = _dfad
	}
	_gaed := _aedc
	_fede := _gece._acca._gg
	switch _dagcc % 360 {
	case 90:
		_fede.Urx, _fede.Ury = _fede.Ury, _fede.Urx
		_gaed = _ba.PdfRectangle{Llx: _fede.Urx - _aedc.Ury, Urx: _fede.Urx - _aedc.Lly, Lly: _aedc.Llx, Ury: _aedc.Urx}
	case 180:
		_gaed = _ba.PdfRectangle{Llx: _fede.Urx - _aedc.Llx, Urx: _fede.Urx - _aedc.Urx, Lly: _fede.Ury - _aedc.Lly, Ury: _fede.Ury - _aedc.Ury}
	case 270:
		_fede.Urx, _fede.Ury = _fede.Ury, _fede.Urx
		_gaed = _ba.PdfRectangle{Llx: _aedc.Ury, Urx: _aedc.Lly, Lly: _fede.Ury - _aedc.Llx, Ury: _fede.Ury - _aedc.Urx}
	}
	if _gaed.Llx > _gaed.Urx {
		_gaed.Llx, _gaed.Urx = _gaed.Urx, _gaed.Llx
	}
	if _gaed.Lly > _gaed.Ury {
		_gaed.Lly, _gaed.Ury = _gaed.Ury, _gaed.Lly
	}
	_dece := textMark{_abdg: _bdbe, PdfRectangle: _gaed, _fdbg: _aedc, _ddbd: _gdea, _dfca: _begg, _geca: _gggg, _gbaa: _cfbf, _ffad: _ebfa, _bdba: _dagcc, _bbcg: _ffede, _fbbcf: _cbac}
	if _badb {
		_f.Log.Info("n\u0065\u0077\u0054\u0065\u0078\u0074M\u0061\u0072\u006b\u003a\u0020\u0073t\u0061\u0072\u0074\u003d\u0025\u002e\u0032f\u0020\u0065\u006e\u0064\u003d\u0025\u002e\u0032\u0066\u0020%\u0073", _feb, _ebfa, _dece.String())
	}
	return _dece, _gbgfa
}
func _fgdge(_fcggd, _dgdab _bc.Point) bool {
	_cefdf := _c.Abs(_fcggd.X - _dgdab.X)
	_ggfe := _c.Abs(_fcggd.Y - _dgdab.Y)
	return _cded(_ggfe, _cefdf)
}
func (_fbcgd *textObject) setTextLeading(_gedf float64) {
	if _fbcgd == nil {
		return
	}
	_fbcgd._gdcg._dgbbf = _gedf
}
func (_ddbbd *textPara) isAtom() *textTable {
	_egcg := _ddbbd
	_gcff := _ddbbd._gfeab
	_dgga := _ddbbd._gagag
	if _gcff.taken() || _dgga.taken() {
		return nil
	}
	_fdab := _gcff._gagag
	if _fdab.taken() || _fdab != _dgga._gfeab {
		return nil
	}
	return _bbcge(_egcg, _gcff, _dgga, _fdab)
}
func (_dbdg paraList) sortReadingOrder() {
	_f.Log.Trace("\u0073\u006fr\u0074\u0052\u0065\u0061\u0064i\u006e\u0067\u004f\u0072\u0064e\u0072\u003a\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_dbdg))
	if len(_dbdg) <= 1 {
		return
	}
	_dbdg.computeEBBoxes()
	_dc.Slice(_dbdg, func(_aaeg, _gafd int) bool { return _agb(_dbdg[_aaeg], _dbdg[_gafd]) <= 0 })
	_beab := _dbdg.topoOrder()
	_dbdg.reorder(_beab)
}

// String returns a human readable description of `s`.
func (_cebea intSet) String() string {
	var _afggf []int
	for _dgcdf := range _cebea {
		if _cebea.has(_dgcdf) {
			_afggf = append(_afggf, _dgcdf)
		}
	}
	_dc.Ints(_afggf)
	return _ca.Sprintf("\u0025\u002b\u0076", _afggf)
}
func _eadc(_dcbcd map[int][]float64) {
	if len(_dcbcd) <= 1 {
		return
	}
	_bgdb := _dceb(_dcbcd)
	if _eca {
		_f.Log.Info("\u0066i\u0078C\u0065\u006c\u006c\u0073\u003a \u006b\u0065y\u0073\u003d\u0025\u002b\u0076", _bgdb)
	}
	var _edce, _bcdd int
	for _edce, _bcdd = range _bgdb {
		if _dcbcd[_bcdd] != nil {
			break
		}
	}
	for _bagb, _decd := range _bgdb[_edce:] {
		_gffd := _dcbcd[_decd]
		if _gffd == nil {
			continue
		}
		if _eca {
			_ca.Printf("\u0025\u0034\u0064\u003a\u0020\u006b\u0030\u003d\u0025\u0064\u0020\u006b1\u003d\u0025\u0064\u000a", _edce+_bagb, _bcdd, _decd)
		}
		_abce := _dcbcd[_decd]
		if _abce[len(_abce)-1] > _gffd[0] {
			_abce[len(_abce)-1] = _gffd[0]
			_dcbcd[_bcdd] = _abce
		}
		_bcdd = _decd
	}
}
func (_efbb rulingList) vertsHorzs() (rulingList, rulingList) {
	var _bbab, _ffbg rulingList
	for _, _ecgf := range _efbb {
		switch _ecgf._ecddg {
		case _efcg:
			_bbab = append(_bbab, _ecgf)
		case _egd:
			_ffbg = append(_ffbg, _ecgf)
		}
	}
	return _bbab, _ffbg
}
func _dgab(_dff, _eggc bounded) float64 {
	_gcga := _dafg(_dff, _eggc)
	if !_bdcfc(_gcga) {
		return _gcga
	}
	return _cdfg(_dff, _eggc)
}
func _cfba(_degb string) bool {
	for _, _bfefa := range _degb {
		if !_ag.IsSpace(_bfefa) {
			return false
		}
	}
	return true
}
func (_faa *wordBag) scanBand(_egaf string, _fgdbe *wordBag, _aebc func(_bcac *wordBag, _fdcfd *textWord) bool, _gabgb, _fbb, _accg float64, _dggda, _ecde bool) int {
	_ddeec := _fgdbe._egad
	var _degd map[int]map[*textWord]struct{}
	if !_dggda {
		_degd = _faa.makeRemovals()
	}
	_abbdg := _ggef * _ddeec
	_faag := 0
	for _, _gbad := range _faa.depthBand(_gabgb-_abbdg, _fbb+_abbdg) {
		if len(_faa._dfec[_gbad]) == 0 {
			continue
		}
		for _, _faf := range _faa._dfec[_gbad] {
			if !(_gabgb-_abbdg <= _faf._addc && _faf._addc <= _fbb+_abbdg) {
				continue
			}
			if !_aebc(_fgdbe, _faf) {
				continue
			}
			_cggcf := 2.0 * _c.Abs(_faf._ggea-_fgdbe._egad) / (_faf._ggea + _fgdbe._egad)
			_bceg := _c.Max(_faf._ggea/_fgdbe._egad, _fgdbe._egad/_faf._ggea)
			_bbdd := _c.Min(_cggcf, _bceg)
			if _accg > 0 && _bbdd > _accg {
				continue
			}
			if _fgdbe.blocked(_faf) {
				continue
			}
			if !_dggda {
				_fgdbe.pullWord(_faf, _gbad, _degd)
			}
			_faag++
			if !_ecde {
				if _faf._addc < _gabgb {
					_gabgb = _faf._addc
				}
				if _faf._addc > _fbb {
					_fbb = _faf._addc
				}
			}
			if _dggda {
				break
			}
		}
	}
	if !_dggda {
		_faa.applyRemovals(_degd)
	}
	return _faag
}
func _gffgc(_affg, _dadeg float64) bool { return _c.Abs(_affg-_dadeg) <= _gdae }

// String returns a description of `b`.
func (_bdcc *wordBag) String() string {
	var _bdcg []string
	for _, _dec := range _bdcc.depthIndexes() {
		_cfca := _bdcc._dfec[_dec]
		for _, _cafe := range _cfca {
			_bdcg = append(_bdcg, _cafe._eabcc)
		}
	}
	return _ca.Sprintf("\u0025.\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065=\u0025\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0025\u0071", _bdcc.PdfRectangle, _bdcc._egad, len(_bdcg), _bdcg)
}
func _eecb(_dbaa *wordBag, _cdcb *textWord, _ffcgc float64) bool {
	return _dbaa.Urx <= _cdcb.Llx && _cdcb.Llx < _dbaa.Urx+_ffcgc
}
func _bcecc(_ega _bc.Point) *subpath { return &subpath{_caf: []_bc.Point{_ega}} }
func _fecgc(_ebbb []TextMark, _eafd *int) []TextMark {
	_ggab := _ebbb[len(_ebbb)-1]
	_bbcb := []rune(_ggab.Text)
	if len(_bbcb) == 1 {
		_ebbb = _ebbb[:len(_ebbb)-1]
		_fdccc := _ebbb[len(_ebbb)-1]
		*_eafd = _fdccc.Offset + len(_fdccc.Text)
	} else {
		_fagc := _cbee(_ggab.Text)
		*_eafd += len(_fagc) - len(_ggab.Text)
		_ggab.Text = _fagc
	}
	return _ebbb
}
func _dafg(_acec, _gdf bounded) float64 { return _acec.bbox().Llx - _gdf.bbox().Llx }

// RangeOffset returns the TextMarks in `ma` that overlap text[start:end] in the extracted text.
// These are tm: `start` <= tm.Offset + len(tm.Text) && tm.Offset < `end` where
// `start` and `end` are offsets in the extracted text.
// NOTE: TextMarks can contain multiple characters. e.g. "ffi" for the ﬃ ligature so the first and
// last elements of the returned TextMarkArray may only partially overlap text[start:end].
func (_cdfd *TextMarkArray) RangeOffset(start, end int) (*TextMarkArray, error) {
	if _cdfd == nil {
		return nil, _g.New("\u006da\u003d\u003d\u006e\u0069\u006c")
	}
	if end < start {
		return nil, _ca.Errorf("\u0065\u006e\u0064\u0020\u003c\u0020\u0073\u0074\u0061\u0072\u0074\u002e\u0020\u0052\u0061n\u0067\u0065\u004f\u0066\u0066\u0073\u0065\u0074\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020\u0073\u0074\u0061\u0072t=\u0025\u0064\u0020\u0065\u006e\u0064\u003d\u0025\u0064\u0020", start, end)
	}
	_ffc := len(_cdfd._afdf)
	if _ffc == 0 {
		return _cdfd, nil
	}
	if start < _cdfd._afdf[0].Offset {
		start = _cdfd._afdf[0].Offset
	}
	if end > _cdfd._afdf[_ffc-1].Offset+1 {
		end = _cdfd._afdf[_ffc-1].Offset + 1
	}
	_fgbg := _dc.Search(_ffc, func(_fcd int) bool { return _cdfd._afdf[_fcd].Offset+len(_cdfd._afdf[_fcd].Text)-1 >= start })
	if !(0 <= _fgbg && _fgbg < _ffc) {
		_eagf := _ca.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073\u0074\u0061\u0072\u0074\u003d%\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009 \u006c\u0061\u0073\u0074\u003d%\u0076", start, _fgbg, _ffc, _cdfd._afdf[0], _cdfd._afdf[_ffc-1])
		return nil, _eagf
	}
	_dagc := _dc.Search(_ffc, func(_gafg int) bool { return _cdfd._afdf[_gafg].Offset > end-1 })
	if !(0 <= _dagc && _dagc < _ffc) {
		_fbeb := _ca.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020r\u0061\u006e\u0067e\u002e\u0020\u0065n\u0064\u003d%\u0064\u0020\u0069\u0045\u006e\u0064=\u0025d \u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d\u0025\u0076", end, _dagc, _ffc, _cdfd._afdf[0], _cdfd._afdf[_ffc-1])
		return nil, _fbeb
	}
	if _dagc <= _fgbg {
		return nil, _ca.Errorf("\u0069\u0045\u006e\u0064\u0020\u003c=\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003a\u0020\u0073\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u0065\u006ed\u003d\u0025\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020i\u0045n\u0064\u003d\u0025\u0064", start, end, _fgbg, _dagc)
	}
	return &TextMarkArray{_afdf: _cdfd._afdf[_fgbg:_dagc]}, nil
}
func (_dedg *textObject) renderText(_fbgge []byte) error {
	if _dedg._fged {
		_f.Log.Debug("\u0072\u0065\u006e\u0064\u0065r\u0054\u0065\u0078\u0074\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0066\u006f\u006e\u0074\u002e\u0020\u004e\u006f\u0074\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u002e")
		return nil
	}
	_acee := _dedg.getCurrentFont()
	_dcf := _acee.BytesToCharcodes(_fbgge)
	_bfdd, _geaf, _gagd := _acee.CharcodesToStrings(_dcf)
	if _gagd > 0 {
		_f.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _geaf, _gagd)
	}
	_dedg._gdcg._efc += _geaf
	_dedg._gdcg._fcge += _gagd
	_bggd := _dedg._gdcg
	_adcbc := _bggd._bgg
	_ffab := _bggd._fgdd / 100.0
	_gda := _abee
	if _acee.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_gda = 1
	}
	_bde, _abff := _acee.GetRuneMetrics(' ')
	if !_abff {
		_bde, _abff = _acee.GetCharMetrics(32)
	}
	if !_abff {
		_bde, _ = _ba.DefaultFont().GetRuneMetrics(' ')
	}
	_cgac := _bde.Wx * _gda
	_f.Log.Trace("\u0073p\u0061\u0063e\u0057\u0069\u0064t\u0068\u003d\u0025\u002e\u0032\u0066\u0020t\u0065\u0078\u0074\u003d\u0025\u0071 \u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0066\u006f\u006et\u0053\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066", _cgac, _bfdd, _acee, _adcbc)
	_gcbg := _bc.NewMatrix(_adcbc*_ffab, 0, 0, _adcbc, 0, _bggd._aea)
	if _gbdd {
		_f.Log.Info("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0074\u0065\u0078t\u0073\u003d\u0025\u0071", len(_dcf), _dcf, _bfdd)
	}
	_f.Log.Trace("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0072\u0075\u006ee\u0073\u003d\u0025\u0071", len(_dcf), _dcf, len(_bfdd))
	_bdb := _dedg.getFillColor()
	_bcg := _dedg.getStrokeColor()
	for _ceg, _bcec := range _bfdd {
		_eba := []rune(_bcec)
		if len(_eba) == 1 && _eba[0] == '\x00' {
			continue
		}
		_ggdd := _dcf[_ceg]
		_ggff := _dedg._fdcd.CTM.Mult(_dedg._ddfg).Mult(_gcbg)
		_aca := 0.0
		if len(_eba) == 1 && _eba[0] == 32 {
			_aca = _bggd._cef
		}
		_gedg, _gfg := _acee.GetCharMetrics(_ggdd)
		if !_gfg {
			_f.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u004e\u006f \u006d\u0065\u0074r\u0069\u0063\u0020\u0066\u006f\u0072\u0020\u0063\u006fde\u003d\u0025\u0064 \u0072\u003d0\u0078\u0025\u0030\u0034\u0078\u003d%\u002b\u0071 \u0025\u0073", _ggdd, _eba, _eba, _acee)
			return _ca.Errorf("\u006e\u006f\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073:\u0020f\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0063\u006f\u0064\u0065\u003d\u0025\u0064", _acee.String(), _ggdd)
		}
		_gge := _bc.Point{X: _gedg.Wx * _gda, Y: _gedg.Wy * _gda}
		_gbdb := _bc.Point{X: (_gge.X*_adcbc + _aca) * _ffab}
		_gef := _bc.Point{X: (_gge.X*_adcbc + _bggd._aed + _aca) * _ffab}
		if _gbdd {
			_f.Log.Info("\u0074\u0066\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0063\u003d\u0025\u002e\u0032f\u0020t\u0077\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _adcbc, _bggd._aed, _bggd._cef, _ffab)
			_f.Log.Info("\u0064x\u002c\u0064\u0079\u003d%\u002e\u0033\u0066\u0020\u00740\u003d%\u002e3\u0066\u0020\u0074\u003d\u0025\u002e\u0033f", _gge, _gbdb, _gef)
		}
		_aad := _cfdg(_gbdb)
		_aga := _cfdg(_gef)
		_cedb := _dedg._fdcd.CTM.Mult(_dedg._ddfg).Mult(_aad)
		if _adec {
			_f.Log.Info("e\u006e\u0064\u003a\u000a\tC\u0054M\u003d\u0025\u0073\u000a\u0009 \u0074\u006d\u003d\u0025\u0073\u000a"+"\u0009\u0020t\u0064\u003d\u0025s\u0020\u0078\u006c\u0061\u0074\u003d\u0025\u0073\u000a"+"\u0009t\u0064\u0030\u003d\u0025s\u000a\u0009\u0020\u0020\u2192 \u0025s\u0020x\u006c\u0061\u0074\u003d\u0025\u0073", _dedg._fdcd.CTM, _dedg._ddfg, _aga, _fcfd(_dedg._fdcd.CTM.Mult(_dedg._ddfg).Mult(_aga)), _aad, _cedb, _fcfd(_cedb))
		}
		_adg, _gff := _dedg.newTextMark(_bg.ExpandLigatures(_eba), _ggff, _fcfd(_cedb), _c.Abs(_cgac*_ggff.ScalingFactorX()), _acee, _dedg._gdcg._aed, _bdb, _bcg)
		if !_gff {
			_f.Log.Debug("\u0054\u0065\u0078\u0074\u0020\u006d\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069d\u0065 \u0070\u0061\u0067\u0065\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		if _acee == nil {
			_f.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u002e")
		} else if _acee.Encoder() == nil {
			_f.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020N\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u002e\u0020\u0066o\u006et\u003d\u0025\u0073", _acee)
		} else {
			if _cda, _fdb := _acee.Encoder().CharcodeToRune(_ggdd); _fdb {
				_adg._deee = string(_cda)
			}
		}
		_f.Log.Trace("i\u003d\u0025\u0064\u0020\u0063\u006fd\u0065\u003d\u0025\u0064\u0020\u006d\u0061\u0072\u006b=\u0025\u0073\u0020t\u0072m\u003d\u0025\u0073", _ceg, _ggdd, _adg, _ggff)
		_dedg._adda = append(_dedg._adda, &_adg)
		_dedg._ddfg.Concat(_aga)
	}
	return nil
}
func (_eefc *textPara) depth() float64 {
	if _eefc._bgea {
		return -1.0
	}
	if len(_eefc._efaf) > 0 {
		return _eefc._efaf[0]._fcaa
	}
	return _eefc._ebcb.depth()
}

type fontEntry struct {
	_beac *_ba.PdfFont
	_bbc  int64
}

func _gceeb(_adegc _ba.PdfRectangle) *ruling {
	return &ruling{_ecddg: _efcg, _ccbf: _adegc.Llx, _gcbgg: _adegc.Lly, _cgeee: _adegc.Ury}
}

// Len returns the number of TextMarks in `ma`.
func (_efd *TextMarkArray) Len() int {
	if _efd == nil {
		return 0
	}
	return len(_efd._afdf)
}
func (_bbddg *textTable) getComposite(_gagaa, _fggeb int) (paraList, _ba.PdfRectangle) {
	_bgaca, _acbed := _bbddg._gabfc[_dfee(_gagaa, _fggeb)]
	if _eca {
		_ca.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0067\u0065\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u002d\u003e\u0025\u0073\u000a", _gagaa, _fggeb, _bgaca.String())
	}
	if !_acbed {
		return nil, _ba.PdfRectangle{}
	}
	return _bgaca.parasBBox()
}

const _cdde = 10

func (_ecebf paraList) xNeighbours(_dgbcf float64) map[*textPara][]int {
	_dggf := make([]event, 2*len(_ecebf))
	if _dgbcf == 0 {
		for _aede, _dfbeg := range _ecebf {
			_dggf[2*_aede] = event{_dfbeg.Llx, true, _aede}
			_dggf[2*_aede+1] = event{_dfbeg.Urx, false, _aede}
		}
	} else {
		for _ageb, _cefg := range _ecebf {
			_dggf[2*_ageb] = event{_cefg.Llx - _dgbcf*_cefg.fontsize(), true, _ageb}
			_dggf[2*_ageb+1] = event{_cefg.Urx + _dgbcf*_cefg.fontsize(), false, _ageb}
		}
	}
	return _ecebf.eventNeighbours(_dggf)
}
func (_abdgeg *ruling) gridIntersecting(_cgea *ruling) bool {
	return _gffgc(_abdgeg._gcbgg, _cgea._gcbgg) && _gffgc(_abdgeg._cgeee, _cgea._cgeee)
}
func (_febec *textWord) bbox() _ba.PdfRectangle { return _febec.PdfRectangle }
func (_gbeea *textLine) markWordBoundaries() {
	_dgbf := _fecg * _gbeea._edfec
	for _gaaa, _dbeb := range _gbeea._bdcb[1:] {
		if _aged(_dbeb, _gbeea._bdcb[_gaaa]) >= _dgbf {
			_dbeb._dfgac = true
		}
	}
}
func _acff(_fecde map[float64]gridTile) []float64 {
	_ebaf := make([]float64, 0, len(_fecde))
	for _gcaac := range _fecde {
		_ebaf = append(_ebaf, _gcaac)
	}
	_dc.Float64s(_ebaf)
	return _ebaf
}

// TableCell is a cell in a TextTable.
type TableCell struct {

	// Text is the extracted text.
	Text string

	// Marks returns the TextMarks corresponding to the text in Text.
	Marks TextMarkArray
}

func (_ffbab *wordBag) depthIndexes() []int {
	if len(_ffbab._dfec) == 0 {
		return nil
	}
	_ggac := make([]int, len(_ffbab._dfec))
	_eegfc := 0
	for _dbbf := range _ffbab._dfec {
		_ggac[_eegfc] = _dbbf
		_eegfc++
	}
	_dc.Ints(_ggac)
	return _ggac
}
func _agb(_gcaa, _gbbcg bounded) float64 {
	_bab := _cdfg(_gcaa, _gbbcg)
	if !_bdcfc(_bab) {
		return _bab
	}
	return _dafg(_gcaa, _gbbcg)
}

type ruling struct {
	_ecddg rulingKind
	_gdee  markKind
	_af.Color
	_ccbf  float64
	_gcbgg float64
	_cgeee float64
	_eeca  float64
}

func (_gga *Extractor) extractPageText(_gbc string, _cca *_ba.PdfPageResources, _cbd _bc.Matrix, _cge int) (*PageText, int, int, error) {
	_f.Log.Trace("\u0065x\u0074\u0072\u0061\u0063t\u0050\u0061\u0067\u0065\u0054e\u0078t\u003a \u006c\u0065\u0076\u0065\u006c\u003d\u0025d", _cge)
	_fbcg := &PageText{_gfed: _gga._gg}
	_dbb := _gafa(_gga._gg)
	var _fad stateStack
	_aag := _fag(_gga, _cca, _da.GraphicsState{}, &_dbb, &_fad)
	_ef := shapesState{_gfgg: _cbd, _bfc: _bc.IdentityMatrix(), _gfad: _aag}
	var _cfb bool
	if _cge > _cdgg {
		_acbe := _g.New("\u0066\u006f\u0072\u006d s\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077")
		_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0065\u0078\u0074\u0072\u0061\u0063\u0074\u0050\u0061\u0067\u0065\u0054\u0065\u0078\u0074\u002e\u0020\u0072\u0065\u0063u\u0072\u0073\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u003d\u0025\u0064 \u0065r\u0072\u003d\u0025\u0076", _cge, _acbe)
		return _fbcg, _dbb._efc, _dbb._fcge, _acbe
	}
	_bgd := _da.NewContentStreamParser(_gbc)
	_ace, _gbe := _bgd.Parse()
	if _gbe != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020e\u0078\u0074\u0072a\u0063\u0074\u0050\u0061g\u0065\u0054\u0065\u0078\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gbe)
		return _fbcg, _dbb._efc, _dbb._fcge, _gbe
	}
	_bae := _da.NewContentStreamProcessor(*_ace)
	_bae.AddHandler(_da.HandlerConditionEnumAllOperands, "", func(_dcea *_da.ContentStreamOperation, _fcf _da.GraphicsState, _abe *_ba.PdfPageResources) error {
		_cgeb := _dcea.Operand
		if _edea {
			_f.Log.Info("\u0026&\u0026\u0020\u006f\u0070\u003d\u0025s", _dcea)
		}
		switch _cgeb {
		case "\u0071":
			if _dfab {
				_f.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ef._bfc)
			}
			_fad.push(&_dbb)
		case "\u0051":
			if !_fad.empty() {
				_dbb = *_fad.pop()
			}
			_ef._bfc = _fcf.CTM
			if _dfab {
				_f.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ef._bfc)
			}
		case "\u0042\u0054":
			if _cfb {
				_f.Log.Debug("\u0042\u0054\u0020\u0063\u0061\u006c\u006c\u0065\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0069n\u0020\u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
				_fbcg._faga = append(_fbcg._faga, _aag._adda...)
			}
			_cfb = true
			_ccd := _fcf
			_ccd.CTM = _cbd.Mult(_ccd.CTM)
			_aag = _fag(_gga, _abe, _ccd, &_dbb, &_fad)
			_ef._gfad = _aag
		case "\u0045\u0054":
			if !_cfb {
				_f.Log.Debug("\u0045\u0054\u0020ca\u006c\u006c\u0065\u0064\u0020\u006f\u0075\u0074\u0073i\u0064e\u0020o\u0066 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_cfb = false
			_fbcg._faga = append(_fbcg._faga, _aag._adda...)
			_aag.reset()
		case "\u0054\u002a":
			_aag.nextLine()
		case "\u0054\u0064":
			if _afba, _bcab := _aag.checkOp(_dcea, 2, true); !_afba {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bcab)
				return _bcab
			}
			_fbg, _bgc, _cba := _cgefb(_dcea.Params)
			if _cba != nil {
				return _cba
			}
			_aag.moveText(_fbg, _bgc)
		case "\u0054\u0044":
			if _fae, _cdf := _aag.checkOp(_dcea, 2, true); !_fae {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cdf)
				return _cdf
			}
			_dgbb, _bfe, _bd := _cgefb(_dcea.Params)
			if _bd != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bd)
				return _bd
			}
			_aag.moveTextSetLeading(_dgbb, _bfe)
		case "\u0054\u006a":
			if _bfbf, _cadc := _aag.checkOp(_dcea, 1, true); !_bfbf {
				_f.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0054\u006a\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d%\u0076", _dcea, _cadc)
				return _cadc
			}
			_cdc, _dab := _ff.GetStringBytes(_dcea.Params[0])
			if !_dab {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020T\u006a\u0020o\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074S\u0074\u0072\u0069\u006e\u0067\u0042\u0079\u0074\u0065\u0073\u0020\u0066a\u0069\u006c\u0065\u0064", _dcea)
				return _ff.ErrTypeError
			}
			return _aag.showText(_cdc)
		case "\u0054\u004a":
			if _bag, _bga := _aag.checkOp(_dcea, 1, true); !_bag {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bga)
				return _bga
			}
			_dfb, _ed := _ff.GetArray(_dcea.Params[0])
			if !_ed {
				_f.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _dcea)
				return _gbe
			}
			return _aag.showTextAdjusted(_dfb)
		case "\u0027":
			if _gfeg, _bgb := _aag.checkOp(_dcea, 1, true); !_gfeg {
				_f.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0027\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bgb)
				return _bgb
			}
			_dddd, _cbae := _ff.GetStringBytes(_dcea.Params[0])
			if !_cbae {
				_f.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020'\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _dcea)
				return _ff.ErrTypeError
			}
			_aag.nextLine()
			return _aag.showText(_dddd)
		case "\u0022":
			if _dbg, _gc := _aag.checkOp(_dcea, 3, true); !_dbg {
				_f.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0022\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gc)
				return _gc
			}
			_cbb, _fadg, _cdgb := _cgefb(_dcea.Params[:2])
			if _cdgb != nil {
				return _cdgb
			}
			_cgg, _dgd := _ff.GetStringBytes(_dcea.Params[2])
			if !_dgd {
				_f.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020\"\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _dcea)
				return _ff.ErrTypeError
			}
			_aag.setCharSpacing(_cbb)
			_aag.setWordSpacing(_fadg)
			_aag.nextLine()
			return _aag.showText(_cgg)
		case "\u0054\u004c":
			_fgg, _abb := _bfbd(_dcea)
			if _abb != nil {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004c\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _abb)
				return _abb
			}
			_aag.setTextLeading(_fgg)
		case "\u0054\u0063":
			_geg, _fgb := _bfbd(_dcea)
			if _fgb != nil {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0063\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fgb)
				return _fgb
			}
			_aag.setCharSpacing(_geg)
		case "\u0054\u0066":
			if _dgg, _edd := _aag.checkOp(_dcea, 2, true); !_dgg {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0066\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _edd)
				return _edd
			}
			_gcb, _ffaf := _ff.GetNameVal(_dcea.Params[0])
			if !_ffaf {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _dcea)
				return _ff.ErrTypeError
			}
			_ffbd, _eag := _ff.GetNumberAsFloat(_dcea.Params[1])
			if !_ffaf {
				_f.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u0046\u006c\u006f\u0061\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065d\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dcea, _eag)
				return _eag
			}
			_eag = _aag.setFont(_gcb, _ffbd)
			_aag._fged = _ce.Is(_eag, _ff.ErrNotSupported)
			if _eag != nil && !_aag._fged {
				return _eag
			}
		case "\u0054\u006d":
			if _aaa, _egbd := _aag.checkOp(_dcea, 6, true); !_aaa {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u006d\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _egbd)
				return _egbd
			}
			_gdc, _dgdb := _ff.GetNumbersAsFloat(_dcea.Params)
			if _dgdb != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dgdb)
				return _dgdb
			}
			_aag.setTextMatrix(_gdc)
		case "\u0054\u0072":
			if _gfa, _daf := _aag.checkOp(_dcea, 1, true); !_gfa {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0072\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _daf)
				return _daf
			}
			_eabe, _agd := _ff.GetIntVal(_dcea.Params[0])
			if !_agd {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0020\u006f\u0070\u003d\u0025\u0073 \u0047e\u0074\u0049\u006e\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _dcea)
				return _ff.ErrTypeError
			}
			_aag.setTextRenderMode(_eabe)
		case "\u0054\u0073":
			if _ggd, _cga := _aag.checkOp(_dcea, 1, true); !_ggd {
				_f.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cga)
				return _cga
			}
			_gca, _cgd := _ff.GetNumberAsFloat(_dcea.Params[0])
			if _cgd != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cgd)
				return _cgd
			}
			_aag.setTextRise(_gca)
		case "\u0054\u0077":
			if _fbgg, _beb := _aag.checkOp(_dcea, 1, true); !_fbgg {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _beb)
				return _beb
			}
			_ageg, _bac := _ff.GetNumberAsFloat(_dcea.Params[0])
			if _bac != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bac)
				return _bac
			}
			_aag.setWordSpacing(_ageg)
		case "\u0054\u007a":
			if _bcf, _fbf := _aag.checkOp(_dcea, 1, true); !_bcf {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fbf)
				return _fbf
			}
			_acc, _fedf := _ff.GetNumberAsFloat(_dcea.Params[0])
			if _fedf != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fedf)
				return _fedf
			}
			_aag.setHorizScaling(_acc)
		case "\u0063\u006d":
			_ef._bfc = _fcf.CTM
			if _ef._bfc.Singular() {
				_daac := _bc.IdentityMatrix().Translate(_ef._bfc.Translation())
				_f.Log.Debug("S\u0069n\u0067\u0075\u006c\u0061\u0072\u0020\u0063\u0074m\u003d\u0025\u0073\u2192%s", _ef._bfc, _daac)
				_ef._bfc = _daac
			}
			if _dfab {
				_f.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ef._bfc)
			}
		case "\u006d":
			if len(_dcea.Params) != 2 {
				_f.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _bb)
				return nil
			}
			_gaf, _eee := _ff.GetNumbersAsFloat(_dcea.Params)
			if _eee != nil {
				return _eee
			}
			_ef.moveTo(_gaf[0], _gaf[1])
		case "\u006c":
			if len(_dcea.Params) != 2 {
				_f.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _bb)
				return nil
			}
			_fbea, _ebg := _ff.GetNumbersAsFloat(_dcea.Params)
			if _ebg != nil {
				return _ebg
			}
			_ef.lineTo(_fbea[0], _fbea[1])
		case "\u0063":
			if len(_dcea.Params) != 6 {
				return _bb
			}
			_dafb, _gbga := _ff.GetNumbersAsFloat(_dcea.Params)
			if _gbga != nil {
				return _gbga
			}
			_f.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _dafb)
			_ef.cubicTo(_dafb[0], _dafb[1], _dafb[2], _dafb[3], _dafb[4], _dafb[5])
		case "\u0076", "\u0079":
			if len(_dcea.Params) != 4 {
				return _bb
			}
			_ced, _abf := _ff.GetNumbersAsFloat(_dcea.Params)
			if _abf != nil {
				return _abf
			}
			_f.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _ced)
			_ef.quadraticTo(_ced[0], _ced[1], _ced[2], _ced[3])
		case "\u0068":
			_ef.closePath()
		case "\u0072\u0065":
			if len(_dcea.Params) != 4 {
				return _bb
			}
			_debc, _cbaeb := _ff.GetNumbersAsFloat(_dcea.Params)
			if _cbaeb != nil {
				return _cbaeb
			}
			_ef.drawRectangle(_debc[0], _debc[1], _debc[2], _debc[3])
			_ef.closePath()
		case "\u0053":
			_ef.stroke(&_fbcg._cfde)
			_ef.clearPath()
		case "\u0073":
			_ef.closePath()
			_ef.stroke(&_fbcg._cfde)
			_ef.clearPath()
		case "\u0046":
			_ef.fill(&_fbcg._fdd)
			_ef.clearPath()
		case "\u0066", "\u0066\u002a":
			_ef.closePath()
			_ef.fill(&_fbcg._fdd)
			_ef.clearPath()
		case "\u0042", "\u0042\u002a":
			_ef.fill(&_fbcg._fdd)
			_ef.stroke(&_fbcg._cfde)
			_ef.clearPath()
		case "\u0062", "\u0062\u002a":
			_ef.closePath()
			_ef.fill(&_fbcg._fdd)
			_ef.stroke(&_fbcg._cfde)
			_ef.clearPath()
		case "\u006e":
			_ef.clearPath()
		case "\u0044\u006f":
			if len(_dcea.Params) == 0 {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0058\u004fbj\u0065c\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0070\u0065\u0072\u0061n\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u006f\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072.\u0020\u0047\u006f\u0074\u0020\u0025\u002b\u0076\u002e", _dcea.Params)
				return _ff.ErrRangeError
			}
			_ggf, _aebf := _ff.GetName(_dcea.Params[0])
			if !_aebf {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u006f\u0020\u006f\u0070e\u0072a\u0074\u006f\u0072\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006fp\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u002b\u0076\u002e", _dcea.Params[0])
				return _ff.ErrTypeError
			}
			_, _dgf := _abe.GetXObjectByName(*_ggf)
			if _dgf != _ba.XObjectTypeForm {
				break
			}
			_gad, _aebf := _gga._afd[_ggf.String()]
			if !_aebf {
				_adc, _cggc := _abe.GetXObjectFormByName(*_ggf)
				if _cggc != nil {
					_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cggc)
					return _cggc
				}
				_fde, _cggc := _adc.GetContentStream()
				if _cggc != nil {
					_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cggc)
					return _cggc
				}
				_gbba := _adc.Resources
				if _gbba == nil {
					_gbba = _abe
				}
				_agda, _gdg, _bedg, _cggc := _gga.extractPageText(string(_fde), _gbba, _cbd.Mult(_fcf.CTM), _cge+1)
				if _cggc != nil {
					_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cggc)
					return _cggc
				}
				_gad = textResult{*_agda, _gdg, _bedg}
				_gga._afd[_ggf.String()] = _gad
			}
			_ef._bfc = _fcf.CTM
			if _dfab {
				_f.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ef._bfc)
			}
			_fbcg._faga = append(_fbcg._faga, _gad._fda._faga...)
			_fbcg._cfde = append(_fbcg._cfde, _gad._fda._cfde...)
			_fbcg._fdd = append(_fbcg._fdd, _gad._fda._fdd...)
			_dbb._efc += _gad._ead
			_dbb._fcge += _gad._fggb
		case "\u0072\u0067", "\u0067", "\u006b", "\u0063\u0073", "\u0073\u0063", "\u0073\u0063\u006e":
			_aag._fdcd.ColorspaceNonStroking = _fcf.ColorspaceNonStroking
			_aag._fdcd.ColorNonStroking = _fcf.ColorNonStroking
		case "\u0052\u0047", "\u0047", "\u004b", "\u0043\u0053", "\u0053\u0043", "\u0053\u0043\u004e":
			_aag._fdcd.ColorspaceStroking = _fcf.ColorspaceStroking
			_aag._fdcd.ColorStroking = _fcf.ColorStroking
		}
		return nil
	})
	_gbe = _bae.Process(_cca)
	return _fbcg, _dbb._efc, _dbb._fcge, _gbe
}

type rectRuling struct {
	_acac  rulingKind
	_gbddg markKind
	_af.Color
	_ba.PdfRectangle
}
type subpath struct {
	_caf  []_bc.Point
	_gcbd bool
}

func (_gfcc rulingList) mergePrimary() float64 {
	_egfeg := _gfcc[0]._ccbf
	for _, _eeead := range _gfcc[1:] {
		_egfeg += _eeead._ccbf
	}
	return _egfeg / float64(len(_gfcc))
}

// Marks returns the TextMark collection for a page. It represents all the text on the page.
func (_fca PageText) Marks() *TextMarkArray { return &TextMarkArray{_afdf: _fca._cdge} }
func _cdbce(_gffgb map[float64]map[float64]gridTile) []float64 {
	_cccb := make([]float64, 0, len(_gffgb))
	_gbbfc := make(map[float64]struct{}, len(_gffgb))
	for _, _cfgfb := range _gffgb {
		for _fbgfa := range _cfgfb {
			if _, _fgcab := _gbbfc[_fbgfa]; _fgcab {
				continue
			}
			_cccb = append(_cccb, _fbgfa)
			_gbbfc[_fbgfa] = struct{}{}
		}
	}
	_dc.Float64s(_cccb)
	return _cccb
}
func (_cdga *ruling) equals(_beed *ruling) bool {
	return _cdga._ecddg == _beed._ecddg && _gffgc(_cdga._ccbf, _beed._ccbf) && _gffgc(_cdga._gcbgg, _beed._gcbgg) && _gffgc(_cdga._cgeee, _beed._cgeee)
}

// NewFromContents creates a new extractor from contents and page resources.
func NewFromContents(contents string, resources *_ba.PdfPageResources) (*Extractor, error) {
	_gab := &Extractor{_ge: contents, _fbe: resources, _age: map[string]fontEntry{}, _afd: map[string]textResult{}}
	return _gab, nil
}
func _fag(_gdgb *Extractor, _cbeg *_ba.PdfPageResources, _gbda _da.GraphicsState, _geae *textState, _cfd *stateStack) *textObject {
	return &textObject{_acca: _gdgb, _dgda: _cbeg, _fdcd: _gbda, _gag: _cfd, _gdcg: _geae, _ddfg: _bc.IdentityMatrix(), _bdc: _bc.IdentityMatrix()}
}
func (_ebfc paraList) topoOrder() []int {
	if _cacd {
		_f.Log.Info("\u0074\u006f\u0070\u006f\u004f\u0072\u0064\u0065\u0072\u003a")
	}
	_eecdf := len(_ebfc)
	_ebeb := make([]bool, _eecdf)
	_ecga := make([]int, 0, _eecdf)
	_agfg := _ebfc.llyOrdering()
	var _gfda func(_eadg int)
	_gfda = func(_dceab int) {
		_ebeb[_dceab] = true
		for _eaff := 0; _eaff < _eecdf; _eaff++ {
			if !_ebeb[_eaff] {
				if _ebfc.readBefore(_agfg, _dceab, _eaff) {
					_gfda(_eaff)
				}
			}
		}
		_ecga = append(_ecga, _dceab)
	}
	for _caee := 0; _caee < _eecdf; _caee++ {
		if !_ebeb[_caee] {
			_gfda(_caee)
		}
	}
	return _egcb(_ecga)
}
func _efba(_bfga _ba.PdfRectangle) *ruling {
	return &ruling{_ecddg: _egd, _ccbf: _bfga.Ury, _gcbgg: _bfga.Llx, _cgeee: _bfga.Urx}
}
func _decb(_egde []rulingList) (rulingList, rulingList) {
	var _fbcb rulingList
	for _, _bcgfc := range _egde {
		_fbcb = append(_fbcb, _bcgfc...)
	}
	return _fbcb.vertsHorzs()
}

type textTable struct {
	_ba.PdfRectangle
	_gcead, _ebab int
	_deecf        bool
	_fecac        map[uint64]*textPara
	_gabfc        map[uint64]compositeCell
}
