package creator

import (
	_b "bytes"
	_dda "encoding/xml"
	_c "errors"
	_g "fmt"
	_ba "image"
	_eb "io"
	_df "log"
	_cd "math"
	_e "os"
	_p "path"
	_pf "path/filepath"
	_ab "regexp"
	_f "sort"
	_aa "strconv"
	_a "strings"
	_dd "text/template"
	_bc "unicode"

	_bcd "bitbucket.org/shenghui0779/gopdf/common"
	_da "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ff "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_cc "bitbucket.org/shenghui0779/gopdf/core"
	_ca "bitbucket.org/shenghui0779/gopdf/internal/graphic2d/svg"
	_cad "bitbucket.org/shenghui0779/gopdf/internal/integrations/unichart"
	_gb "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ga "bitbucket.org/shenghui0779/gopdf/model"
	_dc "github.com/gorilla/i18n/linebreak"
	_dce "github.com/unidoc/unichart/render"
	_gf "golang.org/x/text/unicode/bidi"
)

// SetCompactMode sets the compact mode flag for this table.
//
// By enabling compact mode, table cell that contains Paragraph/StyleParagraph
// would not add extra height when calculating it's height.
//
// The default value is false.
func (_egeba *Table) SetCompactMode(enable bool) { _egeba._adbfb = enable }

// NewTOCLine creates a new table of contents line with the default style.
func (_ebee *Creator) NewTOCLine(number, title, page string, level uint) *TOCLine {
	return _gbgfg(number, title, page, level, _ebee.NewTextStyle())
}

// BorderOpacity returns the border opacity of the rectangle (0-1).
func (_acca *Rectangle) BorderOpacity() float64 { return _acca._gffa }

// SetTextVerticalAlignment sets the vertical alignment of the text within the
// bounds of the styled paragraph.
//
// Note: Currently Styled Paragraph doesn't support TextVerticalAlignmentBottom
// as that option only used for aligning text chunks.
//
// In order to change the vertical alignment of individual text chunks, use TextChunk.VerticalAlignment.
func (_faac *StyledParagraph) SetTextVerticalAlignment(align TextVerticalAlignment) {
	_faac._beda = align
}

// NewChart creates a new creator drawable based on the provided
// unichart chart component.
func NewChart(chart _dce.ChartRenderable) *Chart { return _bbgb(chart) }

func (_acg *Invoice) generateLineBlocks(_edb DrawContext) ([]*Block, DrawContext, error) {
	_deec := _cbgg(len(_acg._cdfg))
	_deec.SetMargins(0, 0, 25, 0)
	for _, _gcb := range _acg._cdfg {
		_dggg := _cgfa(_gcb.TextStyle)
		_dggg.SetMargins(0, 0, 1, 0)
		_dggg.Append(_gcb.Value)
		_agac := _deec.NewCell()
		_agac.SetHorizontalAlignment(_gcb.Alignment)
		_agac.SetBackgroundColor(_gcb.BackgroundColor)
		_acg.setCellBorder(_agac, _gcb)
		_agac.SetContent(_dggg)
	}
	for _, _cfba := range _acg._gdcac {
		for _, _ffbca := range _cfba {
			_fcef := _cgfa(_ffbca.TextStyle)
			_fcef.SetMargins(0, 0, 3, 2)
			_fcef.Append(_ffbca.Value)
			_feaa := _deec.NewCell()
			_feaa.SetHorizontalAlignment(_ffbca.Alignment)
			_feaa.SetBackgroundColor(_ffbca.BackgroundColor)
			_acg.setCellBorder(_feaa, _ffbca)
			_feaa.SetContent(_fcef)
		}
	}
	return _deec.GeneratePageBlocks(_edb)
}

// TextVerticalAlignment controls the vertical position of the text
// in a styled paragraph.
type TextVerticalAlignment int

// Add adds a new line with the default style to the table of contents.
func (_egbbd *TOC) Add(number, title, page string, level uint) *TOCLine {
	_gfbcg := _egbbd.AddLine(_ecbaa(TextChunk{Text: number, Style: _egbbd._daaaa}, TextChunk{Text: title, Style: _egbbd._gdeab}, TextChunk{Text: page, Style: _egbbd._afccg}, level, _egbbd._aced))
	if _gfbcg == nil {
		return nil
	}
	_beegd := &_egbbd._ffcde
	_gfbcg.SetMargins(_beegd.Left, _beegd.Right, _beegd.Top, _beegd.Bottom)
	_gfbcg.SetLevelOffset(_egbbd._cbbeg)
	_gfbcg.Separator.Text = _egbbd._fbag
	_gfbcg.Separator.Style = _egbbd._dbdd
	return _gfbcg
}

func (_cbbf *Invoice) drawSection(_ccee, _dafg string) []*StyledParagraph {
	var _ddaac []*StyledParagraph
	if _ccee != "" {
		_acea := _cgfa(_cbbf._dede)
		_acea.SetMargins(0, 0, 0, 5)
		_acea.Append(_ccee)
		_ddaac = append(_ddaac, _acea)
	}
	if _dafg != "" {
		_bfdg := _cgfa(_cbbf._fdbgg)
		_bfdg.Append(_dafg)
		_ddaac = append(_ddaac, _bfdg)
	}
	return _ddaac
}

func _aceb(_badbc *Table, _bfaf DrawContext) ([]*Block, DrawContext, error) {
	var _gdaef []*Block
	_ecgd := NewBlock(_bfaf.PageWidth, _bfaf.PageHeight)
	_badbc.updateRowHeights(_bfaf.Width - _badbc._bebd.Left - _badbc._bebd.Right)
	_aebe := _badbc._bebd.Top
	if _badbc._efce.IsRelative() && !_badbc._gbdb {
		_bgea := _badbc.Height()
		if _bgea > _bfaf.Height-_badbc._bebd.Top && _bgea <= _bfaf.PageHeight-_bfaf.Margins.Top-_bfaf.Margins.Bottom {
			_gdaef = []*Block{NewBlock(_bfaf.PageWidth, _bfaf.PageHeight-_bfaf.Y)}
			var _cece error
			if _, _bfaf, _cece = _dbae().GeneratePageBlocks(_bfaf); _cece != nil {
				return nil, _bfaf, _cece
			}
			_aebe = 0
		}
	}
	_cgfdg := _bfaf
	if _badbc._efce.IsAbsolute() {
		_bfaf.X = _badbc._aabe
		_bfaf.Y = _badbc._bfbac
	} else {
		_bfaf.X += _badbc._bebd.Left
		_bfaf.Y += _aebe
		_bfaf.Width -= _badbc._bebd.Left + _badbc._bebd.Right
		_bfaf.Height -= _aebe
	}
	_caad := _bfaf.Width
	_cffd := _bfaf.X
	_gbadb := _bfaf.Y
	_cdaef := _bfaf.Height
	_fbga := 0
	_ccbaca, _gggg := -1, -1
	if _badbc._gdaae {
		for _dgfa, _egcgc := range _badbc._gebec {
			if _egcgc._bgdca < _badbc._gcdff {
				continue
			}
			if _egcgc._bgdca > _badbc._bgec {
				break
			}
			if _ccbaca < 0 {
				_ccbaca = _dgfa
			}
			_gggg = _dgfa
		}
	}
	if _caaca := _badbc.wrapContent(_bfaf); _caaca != nil {
		return nil, _bfaf, _caaca
	}
	_badbc.updateRowHeights(_bfaf.Width - _badbc._bebd.Left - _badbc._bebd.Right)
	var (
		_cega  bool
		_gegbd int
		_cgebg int
		_fbacg bool
		_adec  int
		_fffed error
	)
	for _fcbgb := 0; _fcbgb < len(_badbc._gebec); _fcbgb++ {
		_ebgg := _badbc._gebec[_fcbgb]
		if _fbgg, _ebgag := _badbc.getLastCellFromCol(_ebgg._fcbe); _fbgg == _fcbgb {
			if (_ebgag._bgdca + _ebgag._fddc - 1) < _badbc._dbbf {
				for _cdbdd := _ebgg._bgdca; _cdbdd < _badbc._dbbf; _cdbdd++ {
					_eagd := &TableCell{}
					_eagd._bgdca = _cdbdd + 1
					_eagd._fddc = 1
					_eagd._fcbe = _ebgg._fcbe
					_badbc._gebec = append(_badbc._gebec, _eagd)
				}
			}
		}
		_dgfgc := _ebgg.width(_badbc._eeggd, _caad)
		_fdgga := float64(0.0)
		for _eecg := 0; _eecg < _ebgg._fcbe-1; _eecg++ {
			_fdgga += _badbc._eeggd[_eecg] * _caad
		}
		_cecc := float64(0.0)
		for _bggcb := _fbga; _bggcb < _ebgg._bgdca-1; _bggcb++ {
			_cecc += _badbc._ecab[_bggcb]
		}
		_bfaf.Height = _cdaef - _cecc
		_eadb := float64(0.0)
		for _gfdba := 0; _gfdba < _ebgg._fddc; _gfdba++ {
			_eadb += _badbc._ecab[_ebgg._bgdca+_gfdba-1]
		}
		_aaebf := _fbacg && _ebgg._bgdca != _adec
		_adec = _ebgg._bgdca
		if _aaebf || _eadb > _bfaf.Height {
			if _badbc._ggea && !_fbacg {
				_fbacg, _fffed = _badbc.wrapRow(_fcbgb, _bfaf, _caad)
				if _fffed != nil {
					return nil, _bfaf, _fffed
				}
				if _fbacg {
					_fcbgb--
					continue
				}
			}
			_gdaef = append(_gdaef, _ecgd)
			_ecgd = NewBlock(_bfaf.PageWidth, _bfaf.PageHeight)
			_cffd = _bfaf.Margins.Left + _badbc._bebd.Left
			_gbadb = _bfaf.Margins.Top
			_bfaf.Height = _bfaf.PageHeight - _bfaf.Margins.Top - _bfaf.Margins.Bottom
			_bfaf.Page++
			_cdaef = _bfaf.Height
			_fbga = _ebgg._bgdca - 1
			_cecc = 0
			_fbacg = false
			if _badbc._gdaae && _ccbaca >= 0 {
				_gegbd = _fcbgb
				_fcbgb = _ccbaca - 1
				_cgebg = _fbga
				_fbga = _badbc._gcdff - 1
				_cega = true
				if _ebgg._fddc > (_badbc._dbbf-_adec) || (_ebgg._fddc > 1 && _fcbgb < 0) {
					_bcd.Log.Debug("\u0054a\u0062\u006ce\u0020\u0068\u0065a\u0064\u0065\u0072\u0020\u0072\u006f\u0077s\u0070\u0061\u006e\u0020\u0065\u0078c\u0065\u0065\u0064\u0073\u0020\u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u002e")
					_cega = false
					_ccbaca, _gggg = -1, -1
				}
				continue
			}
			if _aaebf {
				_fcbgb--
				continue
			}
		}
		_bfaf.Width = _dgfgc
		_bfaf.X = _cffd + _fdgga
		_bfaf.Y = _gbadb + _cecc
		if _eadb > _bfaf.PageHeight-_bfaf.Margins.Top-_bfaf.Margins.Bottom {
			_eadb = _bfaf.PageHeight - _bfaf.Margins.Top - _bfaf.Margins.Bottom
		}
		_afff := _ada(_bfaf.X, _bfaf.Y, _dgfgc, _eadb)
		if _ebgg._dedd != nil {
			_afff.SetFillColor(_ebgg._dedd)
		}
		_afff.LineStyle = _ebgg._acafb
		_afff._dad = _ebgg._ebce
		_afff._fbae = _ebgg._bage
		_afff._acef = _ebgg._efdbd
		_afff._aacc = _ebgg._ggce
		if _ebgg._efcg != nil {
			_afff.SetColorLeft(_ebgg._efcg)
		}
		if _ebgg._ebgec != nil {
			_afff.SetColorBottom(_ebgg._ebgec)
		}
		if _ebgg._fcec != nil {
			_afff.SetColorRight(_ebgg._fcec)
		}
		if _ebgg._afddf != nil {
			_afff.SetColorTop(_ebgg._afddf)
		}
		_afff.SetWidthBottom(_ebgg._aeba)
		_afff.SetWidthLeft(_ebgg._aabc)
		_afff.SetWidthRight(_ebgg._edag)
		_afff.SetWidthTop(_ebgg._fgbd)
		_dccc := NewBlock(_ecgd._fd, _ecgd._cce)
		_dfef := _ecgd.Draw(_afff)
		if _dfef != nil {
			_bcd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dfef)
		}
		if _ebgg._bfef != nil {
			_afdab := _ebgg._bfef.Width()
			_bbda := _ebgg._bfef.Height()
			_fabgb := 0.0
			switch _facaf := _ebgg._bfef.(type) {
			case *Paragraph:
				if _facaf._dbfb {
					_afdab = _facaf.getMaxLineWidth() / 1000.0
				}
				_edbec, _adaf, _ := _facaf.getTextMetrics()
				_gfcdb, _bece := _edbec*_facaf._agdb, _adaf*_facaf._agdb
				_bbda = _bbda - _bece + _gfcdb
				_fabgb += _gfcdb - _bece
				_cbfg := 0.5
				if _badbc._adbfb {
					_cbfg = 0.3
				}
				switch _ebgg._ddfbf {
				case CellVerticalAlignmentTop:
					_fabgb += _gfcdb * _cbfg
				case CellVerticalAlignmentBottom:
					_fabgb -= _gfcdb * _cbfg
				}
				_afdab += _facaf._bcggf.Left + _facaf._bcggf.Right
				_bbda += _facaf._bcggf.Top + _facaf._bcggf.Bottom
			case *StyledParagraph:
				if _facaf._bfcgb {
					_afdab = _facaf.getMaxLineWidth() / 1000.0
				}
				_edcaa, _bcffb, _bcag := _facaf.getLineMetrics(0)
				_bcadc, _baefe := _edcaa*_facaf._bebe, _bcffb*_facaf._bebe
				if _facaf._beda == TextVerticalAlignmentCenter {
					_fabgb = _baefe - (_bcffb + (_edcaa+_bcag-_bcffb)/2 + (_baefe-_bcffb)/2)
				}
				if len(_facaf._bddb) == 1 {
					_bbda = _bcadc
				} else {
					_bbda = _bbda - _baefe + _bcadc
				}
				_fabgb += _bcadc - _baefe
				switch _ebgg._ddfbf {
				case CellVerticalAlignmentTop:
					_fabgb += _bcadc * 0.5
				case CellVerticalAlignmentBottom:
					_fabgb -= _bcadc * 0.5
				}
				_afdab += _facaf._fadcg.Left + _facaf._fadcg.Right
				_bbda += _facaf._fadcg.Top + _facaf._fadcg.Bottom
			case *Table:
				_afdab = _dgfgc
			case *List:
				_afdab = _dgfgc
			case *Division:
				_afdab = _dgfgc
			case *Chart:
				_afdab = _dgfgc
			case *Line:
				_bbda += _facaf._gabb.Top + _facaf._gabb.Bottom
				_fabgb -= _facaf.Height() / 2
			case *Image:
				_afdab += _facaf._eaac.Left + _facaf._eaac.Right
				_bbda += _facaf._eaac.Top + _facaf._eaac.Bottom
			}
			switch _ebgg._dbdc {
			case CellHorizontalAlignmentLeft:
				_bfaf.X += _ebgg._fccd
				_bfaf.Width -= _ebgg._fccd
			case CellHorizontalAlignmentCenter:
				if _gabdae := _dgfgc - _afdab; _gabdae > 0 {
					_bfaf.X += _gabdae / 2
					_bfaf.Width -= _gabdae / 2
				}
			case CellHorizontalAlignmentRight:
				if _dgfgc > _afdab {
					_bfaf.X = _bfaf.X + _dgfgc - _afdab - _ebgg._fccd
					_bfaf.Width -= _ebgg._fccd
				}
			}
			_fcbbg := _bfaf.Y
			_bbfc := _bfaf.Height
			_bfaf.Y += _fabgb
			switch _ebgg._ddfbf {
			case CellVerticalAlignmentTop:
			case CellVerticalAlignmentMiddle:
				if _egad := _eadb - _bbda; _egad > 0 {
					_bfaf.Y += _egad / 2
					_bfaf.Height -= _egad / 2
				}
			case CellVerticalAlignmentBottom:
				if _eadb > _bbda {
					_bfaf.Y = _bfaf.Y + _eadb - _bbda
					_bfaf.Height = _eadb
				}
			}
			_ccbc := _ecgd.DrawWithContext(_ebgg._bfef, _bfaf)
			if _ccbc != nil {
				if _c.Is(_ccbc, ErrContentNotFit) && !_aaebf {
					_ecgd = _dccc
					_aaebf = true
					_fcbgb--
					continue
				}
				_bcd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ccbc)
			}
			_bfaf.Y = _fcbbg
			_bfaf.Height = _bbfc
		}
		_bfaf.Y += _eadb
		_bfaf.Height -= _eadb
		if _cega && _fcbgb+1 > _gggg {
			_gbadb += _cecc + _eadb
			_cdaef -= _eadb + _cecc
			_fbga = _cgebg
			_fcbgb = _gegbd - 1
			_cega = false
		}
	}
	_gdaef = append(_gdaef, _ecgd)
	if _badbc._efce.IsAbsolute() {
		return _gdaef, _cgfdg, nil
	}
	_bfaf.X = _cgfdg.X
	_bfaf.Width = _cgfdg.Width
	_bfaf.Y += _badbc._bebd.Bottom
	_bfaf.Height -= _badbc._bebd.Bottom
	return _gdaef, _bfaf, nil
}

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_dbcb *Creator) DrawTemplate(r _eb.Reader, data interface{}, options *TemplateOptions) error {
	return _gebee(_dbcb, r, data, options, _dbcb)
}

// AddPatternResource adds pattern dictionary inside the resources dictionary.
func (_eebca *RadialShading) AddPatternResource(block *Block) (_cdde _cc.PdfObjectName, _ccgf error) {
	_gecc := 1
	_ffaca := _cc.PdfObjectName("\u0050" + _aa.Itoa(_gecc))
	for block._ccf.HasPatternByName(_ffaca) {
		_gecc++
		_ffaca = _cc.PdfObjectName("\u0050" + _aa.Itoa(_gecc))
	}
	if _bcgfg := block._ccf.SetPatternByName(_ffaca, _eebca.ToPdfShadingPattern().ToPdfObject()); _bcgfg != nil {
		return "", _bcgfg
	}
	return _ffaca, nil
}

// Height returns the height of the Paragraph. The height is calculated based on the input text and
// how it is wrapped within the container. Does not include Margins.
func (_fgbb *Paragraph) Height() float64 {
	_fgbb.wrapText()
	return float64(len(_fgbb._aggb)) * _fgbb._agdb * _fgbb._dbad
}

func (_ccecg *templateProcessor) parseTextVerticalAlignmentAttr(_aacafc, _bgcbg string) TextVerticalAlignment {
	_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _aacafc, _bgcbg)
	_fecbf := map[string]TextVerticalAlignment{"\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065": TextVerticalAlignmentBaseline, "\u0063\u0065\u006e\u0074\u0065\u0072": TextVerticalAlignmentCenter}[_bgcbg]
	return _fecbf
}

// NewRadialGradientColor creates a radial gradient color that could act as a color in other componenents.
// Note: The innerRadius must be smaller than outerRadius for the circle to render properly.
func (_gce *Creator) NewRadialGradientColor(x float64, y float64, innerRadius float64, outerRadius float64, colorPoints []*ColorPoint) *RadialShading {
	return _eebb(x, y, innerRadius, outerRadius, colorPoints)
}

// NewTextStyle creates a new text style object which can be used to style
// chunks of text.
// Default attributes:
// Font: Helvetica
// Font size: 10
// Encoding: WinAnsiEncoding
// Text color: black
func (_faaf *Creator) NewTextStyle() TextStyle { return _gbdbf(_faaf._ffcg) }

// NewFilledCurve returns a instance of filled curve.
func (_fbeg *Creator) NewFilledCurve() *FilledCurve { return _bcbd() }

// FrontpageFunctionArgs holds the input arguments to a front page drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FrontpageFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// SetAddressHeadingStyle sets the style properties used to render the
// heading of the invoice address sections.
func (_gade *Invoice) SetAddressHeadingStyle(style TextStyle) { _gade._ffgb = style }

func _dgga(_acda []byte) (*Image, error) {
	_gbgd := _b.NewReader(_acda)
	_beab, _dgfb := _ga.ImageHandling.Read(_gbgd)
	if _dgfb != nil {
		_bcd.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _dgfb)
		return nil, _dgfb
	}
	return _cfcc(_beab)
}

// SetIndent sets the left offset of the list when nested into another list.
func (_efdd *List) SetIndent(indent float64) { _efdd._addd = indent; _efdd._ecee = false }

var PPI float64 = 72

func (_gfge *templateProcessor) parsePositioningAttr(_cdbb, _adga string) Positioning {
	_bcd.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0069\u006e\u0067\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _cdbb, _adga)
	_dbaea := map[string]Positioning{"\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065": PositionRelative, "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065": PositionAbsolute}[_adga]
	return _dbaea
}

// ColorCMYKFrom8bit creates a Color from c,m,y,k values (0-100).
// Example:
//
//	red := ColorCMYKFrom8Bit(0, 100, 100, 0)
func ColorCMYKFrom8bit(c, m, y, k byte) Color {
	return cmykColor{_bdge: _cd.Min(float64(c), 100) / 100.0, _eegd: _cd.Min(float64(m), 100) / 100.0, _ebe: _cd.Min(float64(y), 100) / 100.0, _agg: _cd.Min(float64(k), 100) / 100.0}
}

// InsertColumn inserts a column in the line items table at the specified index.
func (_ecge *Invoice) InsertColumn(index uint, description string) *InvoiceCell {
	_caef := uint(len(_ecge._cdfg))
	if index > _caef {
		index = _caef
	}
	_aggga := _ecge.NewColumn(description)
	_ecge._cdfg = append(_ecge._cdfg[:index], append([]*InvoiceCell{_aggga}, _ecge._cdfg[index:]...)...)
	return _aggga
}

// MoveTo moves the drawing context to absolute coordinates (x, y).
func (_dgfg *Creator) MoveTo(x, y float64) { _dgfg._ggab.X = x; _dgfg._ggab.Y = y }

// Fit fits the chunk into the specified bounding box, cropping off the
// remainder in a new chunk, if it exceeds the specified dimensions.
// NOTE: The method assumes a line height of 1.0. In order to account for other
// line height values, the passed in height must be divided by the line height:
// height = height / lineHeight
func (_abdfg *TextChunk) Fit(width, height float64) (*TextChunk, error) {
	_cdcgc, _fadf := _abdfg.Wrap(width)
	if _fadf != nil {
		return nil, _fadf
	}
	_bede := int(height / _abdfg.Style.FontSize)
	if _bede >= len(_cdcgc) {
		return nil, nil
	}
	_cgca := "\u000a"
	_abdfg.Text = _a.Replace(_a.Join(_cdcgc[:_bede], "\u0020"), _cgca+"\u0020", _cgca, -1)
	_abdcc := _a.Replace(_a.Join(_cdcgc[_bede:], "\u0020"), _cgca+"\u0020", _cgca, -1)
	return NewTextChunk(_abdcc, _abdfg.Style), nil
}

func _cbba(_egac *_ca.GraphicSVG) (*GraphicSVG, error) {
	return &GraphicSVG{_debae: _egac, _dabg: PositionRelative, _fgfa: Margins{Top: 10, Bottom: 10}}, nil
}

// SetHeight sets the Image's document height to specified h.
func (_dfdf *Image) SetHeight(h float64) { _dfdf._cgdb = h }

// SetMargins sets the Table's left, right, top, bottom margins.
func (_eafbf *Table) SetMargins(left, right, top, bottom float64) {
	_eafbf._bebd.Left = left
	_eafbf._bebd.Right = right
	_eafbf._bebd.Top = top
	_eafbf._bebd.Bottom = bottom
}

// Color interface represents colors in the PDF creator.
type Color interface {
	ToRGB() (float64, float64, float64)
}

func (_cdfgc *Paragraph) getTextMetrics() (_beae, _dfbf, _ffce float64) {
	_baage := _abff(_cdfgc._bfbd, _cdfgc._dbad)
	if _baage._agaa > _beae {
		_beae = _baage._agaa
	}
	if _baage._ceddg < _ffce {
		_ffce = _baage._ceddg
	}
	if _gace := _cdfgc._dbad; _gace > _dfbf {
		_dfbf = _gace
	}
	return _beae, _dfbf, _ffce
}

// GetMargins returns the margins of the graphic svg (left, right, top, bottom).
func (_cfbb *GraphicSVG) GetMargins() (float64, float64, float64, float64) {
	return _cfbb._fgfa.Left, _cfbb._fgfa.Right, _cfbb._fgfa.Top, _cfbb._fgfa.Bottom
}

func _afaef(_fccbf []_ff.CubicBezierCurve) *PolyBezierCurve {
	return &PolyBezierCurve{_ecgf: &_ff.PolyBezierCurve{Curves: _fccbf, BorderColor: _ga.NewPdfColorDeviceRGB(0, 0, 0), BorderWidth: 1.0}, _gbbfd: 1.0, _dagd: 1.0}
}

// SetPageLabels adds the specified page labels to the PDF file generated
// by the creator. See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
// NOTE: for existing PDF files, the page label ranges object can be obtained
// using the model.PDFReader's GetPageLabels method.
func (_ecdc *Creator) SetPageLabels(pageLabels _cc.PdfObject) { _ecdc._fdff = pageLabels }

// Columns returns all the columns in the invoice line items table.
func (_efec *Invoice) Columns() []*InvoiceCell { return _efec._cdfg }

func (_eggdc *templateProcessor) parseChapterHeading(_gddfc *templateNode) (interface{}, error) {
	if _gddfc._gbdf == nil {
		_eggdc.nodeLogError(_gddfc, "\u0043\u0068a\u0070\u0074\u0065\u0072 \u0068\u0065a\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0072e\u006e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c\u002e")
		return nil, _bfegb
	}
	_bfdga, _cfgfa := _gddfc._gbdf._defae.(*Chapter)
	if !_cfgfa {
		_eggdc.nodeLogError(_gddfc, "\u0043h\u0061\u0070t\u0065\u0072\u0020h\u0065\u0061\u0064\u0069\u006e\u0067\u0020p\u0061\u0072\u0065\u006e\u0074\u0020(\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u002e", _gddfc._gbdf._defae)
		return nil, _bfegb
	}
	_bfbgb := _bfdga.GetHeading()
	if _, _gbee := _eggdc.parseParagraph(_gddfc, _bfbgb); _gbee != nil {
		return nil, _gbee
	}
	return _bfbgb, nil
}

func (_ddb *Block) drawToPage(_bcgd *_ga.PdfPage) error {
	_caa := &_da.ContentStreamOperations{}
	if _bcgd.Resources == nil {
		_bcgd.Resources = _ga.NewPdfPageResources()
	}
	_dfgc := _cgd(_caa, _bcgd.Resources, _ddb._fa, _ddb._ccf)
	if _dfgc != nil {
		return _dfgc
	}
	if _dfgc = _cdae(_ddb._ccf, _bcgd.Resources); _dfgc != nil {
		return _dfgc
	}
	if _dfgc = _bcgd.AppendContentBytes(_caa.Bytes(), true); _dfgc != nil {
		return _dfgc
	}
	for _, _ge := range _ddb._fdd {
		_bcgd.AddAnnotation(_ge)
	}
	return nil
}

// SetWidthLeft sets border width for left.
func (_fea *border) SetWidthLeft(bw float64) { _fea._fab = bw }

// SetPageSize sets the Creator's page size.  Pages that are added after this will be created with
// this Page size.
// Does not affect pages already created.
//
// Common page sizes are defined as constants.
// Examples:
// 1. c.SetPageSize(creator.PageSizeA4)
// 2. c.SetPageSize(creator.PageSizeA3)
// 3. c.SetPageSize(creator.PageSizeLegal)
// 4. c.SetPageSize(creator.PageSizeLetter)
//
// For custom sizes: Use the PPMM (points per mm) and PPI (points per inch) when defining those based on
// physical page sizes:
//
// Examples:
// 1. 10x15 sq. mm: SetPageSize(PageSize{10*creator.PPMM, 15*creator.PPMM}) where PPMM is points per mm.
// 2. 3x2 sq. inches: SetPageSize(PageSize{3*creator.PPI, 2*creator.PPI}) where PPI is points per inch.
func (_ebgb *Creator) SetPageSize(size PageSize) {
	_ebgb._aad = size
	_ebgb._bagf = size[0]
	_ebgb._aafc = size[1]
	_fcad := 0.1 * _ebgb._bagf
	_ebgb._eaf.Left = _fcad
	_ebgb._eaf.Right = _fcad
	_ebgb._eaf.Top = _fcad
	_ebgb._eaf.Bottom = _fcad
}

// Vertical returns total vertical (top + bottom) margin.
func (_ega *Margins) Vertical() float64 { return _ega.Bottom + _ega.Top }

var PPMM = float64(72 * 1.0 / 25.4)

// AddInternalLink adds a new internal link to the paragraph.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
// The zoom of the destination page is controlled with the zoom
// parameter. Pass in 0 to keep the current zoom value.
func (_dcfda *StyledParagraph) AddInternalLink(text string, page int64, x, y, zoom float64) *TextChunk {
	_adgd := NewTextChunk(text, _dcfda._dbgf)
	_adgd._cgfaa = _bgfbc(page-1, x, y, zoom)
	return _dcfda.appendChunk(_adgd)
}

func (_aag *Block) translate(_aaf, _fc float64) {
	_cdb := _da.NewContentCreator().Translate(_aaf, -_fc).Operations()
	*_aag._fa = append(*_cdb, *_aag._fa...)
	_aag._fa.WrapIfNeeded()
}

// Subtotal returns the invoice subtotal description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_aadc *Invoice) Subtotal() (*InvoiceCell, *InvoiceCell) { return _aadc._gfag[0], _aadc._gfag[1] }

// TableCell defines a table cell which can contain a Drawable as content.
type TableCell struct {
	_dedd         Color
	_acafb        _ff.LineStyle
	_ebce         CellBorderStyle
	_efcg         Color
	_aabc         float64
	_ggce         CellBorderStyle
	_ebgec        Color
	_aeba         float64
	_bage         CellBorderStyle
	_fcec         Color
	_edag         float64
	_efdbd        CellBorderStyle
	_afddf        Color
	_fgbd         float64
	_bgdca, _fcbe int
	_fddc         int
	_daec         int
	_bfef         VectorDrawable
	_dbdc         CellHorizontalAlignment
	_ddfbf        CellVerticalAlignment
	_fccd         float64
	_ccbf         *Table
}

// DashPattern returns the dash pattern of the line.
func (_bgbe *Line) DashPattern() (_cecfb []int64, _fgeb int64) { return _bgbe._fedd, _bgbe._ffgdf }

// SetCoords sets the upper left corner coordinates of the rectangle.
func (_ggadg *Rectangle) SetCoords(x, y float64) { _ggadg._fcgf = x; _ggadg._cbaa = y }

func _fccae(_fae, _fafe, _dfb, _adfa float64) *Ellipse {
	return &Ellipse{_gdbee: _fae, _bgfb: _fafe, _eded: _dfb, _bdcce: _adfa, _fabd: PositionAbsolute, _bbdc: 1.0, _bacf: ColorBlack, _eeaa: 1.0, _eade: 1.0}
}

// SetDate sets the date of the invoice.
func (_egacf *Invoice) SetDate(date string) (*InvoiceCell, *InvoiceCell) {
	_egacf._afdb[1].Value = date
	return _egacf._afdb[0], _egacf._afdb[1]
}

// AddExternalLink adds a new external link to the paragraph.
// The text parameter represents the text that is displayed and the url
// parameter sets the destionation of the link.
func (_gabce *StyledParagraph) AddExternalLink(text, url string) *TextChunk {
	_bddg := NewTextChunk(text, _gabce._dbgf)
	_bddg._cgfaa = _eadc(url)
	return _gabce.appendChunk(_bddg)
}

// Width returns the width of the chart. In relative positioning mode,
// all the available context width is used at render time.
func (_adbf *Chart) Width() float64 { return float64(_adbf._dece.Width()) }

// StyledParagraph represents text drawn with a specified font and can wrap across lines and pages.
// By default occupies the available width in the drawing context.
type StyledParagraph struct {
	_eddc  []*TextChunk
	_afaf  TextStyle
	_dbgf  TextStyle
	_dcgfc TextAlignment
	_beda  TextVerticalAlignment
	_bebe  float64
	_bfcgb bool
	_cccde float64
	_ceba  bool
	_afeg  bool
	_egead TextOverflow
	_ecdcf float64
	_fadcg Margins
	_bedd  Positioning
	_fcadg float64
	_ffdc  float64
	_adbb  float64
	_cgee  float64
	_bddb  [][]*TextChunk
	_egdcb func(_debfb *StyledParagraph, _fbbba DrawContext)
}

func (_febg *templateProcessor) parseLinkAttr(_cfgc, _dcbf string) *_ga.PdfAnnotation {
	_dcbf = _a.TrimSpace(_dcbf)
	if _a.HasPrefix(_dcbf, "\u0075\u0072\u006c(\u0027") && _a.HasSuffix(_dcbf, "\u0027\u0029") && len(_dcbf) > 7 {
		return _eadc(_dcbf[5 : len(_dcbf)-2])
	}
	if _a.HasPrefix(_dcbf, "\u0070\u0061\u0067e\u0028") && _a.HasSuffix(_dcbf, "\u0029") && len(_dcbf) > 6 {
		var (
			_daaffe error
			_dgaa   int64
			_dbaeg  float64
			_dbfa   float64
			_eafge  = 1.0
			_dbada  = _a.Split(_dcbf[5:len(_dcbf)-1], "\u002c")
		)
		_dgaa, _daaffe = _aa.ParseInt(_a.TrimSpace(_dbada[0]), 10, 64)
		if _daaffe != nil {
			_bcd.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _daaffe)
			return nil
		}
		if len(_dbada) >= 2 {
			_dbaeg, _daaffe = _aa.ParseFloat(_a.TrimSpace(_dbada[1]), 64)
			if _daaffe != nil {
				_bcd.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0058\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _daaffe)
				return nil
			}
		}
		if len(_dbada) >= 3 {
			_dbfa, _daaffe = _aa.ParseFloat(_a.TrimSpace(_dbada[2]), 64)
			if _daaffe != nil {
				_bcd.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0059\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _daaffe)
				return nil
			}
		}
		if len(_dbada) >= 4 {
			_eafge, _daaffe = _aa.ParseFloat(_a.TrimSpace(_dbada[3]), 64)
			if _daaffe != nil {
				_bcd.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u007a\u006f\u006f\u006d\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _daaffe)
				return nil
			}
		}
		return _bgfbc(_dgaa-1, _dbaeg, _dbfa, _eafge)
	}
	return nil
}

// SetCoords sets the center coordinates of the ellipse.
func (_fagg *Ellipse) SetCoords(xc, yc float64) { _fagg._gdbee = xc; _fagg._bgfb = yc }

const (
	AnchorBottomLeft AnchorPoint = iota
	AnchorBottomRight
	AnchorTopLeft
	AnchorTopRight
	AnchorCenter
	AnchorLeft
	AnchorRight
	AnchorTop
	AnchorBottom
)

func _abff(_accbc *_ga.PdfFont, _gddec float64) *fontMetrics {
	_dgae := &fontMetrics{}
	if _accbc == nil {
		_bcd.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0069s\u0020\u006e\u0069\u006c")
		return _dgae
	}
	_ccbe, _aagcf := _accbc.GetFontDescriptor()
	if _aagcf != nil {
		_bcd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0067\u0065t\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063ri\u0070\u0074\u006fr\u003a \u0025\u0076", _aagcf)
		return _dgae
	}
	if _dgae._agaa, _aagcf = _ccbe.GetCapHeight(); _aagcf != nil {
		_bcd.Log.Trace("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _aagcf)
	}
	if int(_dgae._agaa) <= 0 {
		_bcd.Log.Trace("\u0057\u0041\u0052\u004e\u003a\u0020\u0043\u0061p\u0020\u0048\u0065ig\u0068\u0074\u0020\u006e\u006f\u0074 \u0061\u0076\u0061\u0069\u006c\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065\u0074t\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u00310\u0030\u0030")
		_dgae._agaa = 1000
	}
	_dgae._agaa *= _gddec / 1000.0
	if _dgae._bgggg, _aagcf = _ccbe.GetXHeight(); _aagcf != nil {
		_bcd.Log.Trace("\u0057\u0041R\u004e\u003a\u0020\u0055n\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0078\u002d\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _aagcf)
	}
	_dgae._bgggg *= _gddec / 1000.0
	if _dgae._agdbg, _aagcf = _ccbe.GetAscent(); _aagcf != nil {
		_bcd.Log.Trace("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0061\u0073\u0063\u0065\u006e\u0074\u003a\u0020%\u0076", _aagcf)
	}
	_dgae._agdbg *= _gddec / 1000.0
	if _dgae._ceddg, _aagcf = _ccbe.GetDescent(); _aagcf != nil {
		_bcd.Log.Trace("\u0057\u0041RN\u003a\u0020\u0055n\u0061\u0062\u006c\u0065 to\u0020ge\u0074\u0020\u0066\u006f\u006e\u0074\u0020de\u0073\u0063\u0065\u006e\u0074\u003a\u0020%\u0076", _aagcf)
	}
	_dgae._ceddg *= _gddec / 1000.0
	return _dgae
}

// SetBorderOpacity sets the border opacity.
func (_bcae *CurvePolygon) SetBorderOpacity(opacity float64) { _bcae._adad = opacity }

type border struct {
	_cbd      float64
	_bcgdb    float64
	_begd     float64
	_cfcd     float64
	_fga      Color
	_efb      Color
	_fab      float64
	_adc      Color
	_gab      float64
	_gdc      Color
	_dggc     float64
	_beea     Color
	_dca      float64
	LineStyle _ff.LineStyle
	_dad      CellBorderStyle
	_fbae     CellBorderStyle
	_acef     CellBorderStyle
	_aacc     CellBorderStyle
}

// SetBorderWidth sets the border width.
func (_bggg *PolyBezierCurve) SetBorderWidth(borderWidth float64) {
	_bggg._ecgf.BorderWidth = borderWidth
}

// GeneratePageBlocks draws the curve onto page blocks.
func (_ffcga *Curve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_egfc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_acag := _da.NewContentCreator()
	_acag.Add_q().Add_w(_ffcga._cdaf).SetStrokingColor(_fce(_ffcga._gag)).Add_m(_ffcga._caec, ctx.PageHeight-_ffcga._fcdb).Add_v(_ffcga._fcg, ctx.PageHeight-_ffcga._cfdd, _ffcga._fcfe, ctx.PageHeight-_ffcga._cbf).Add_S().Add_Q()
	_dgef := _egfc.addContentsByString(_acag.String())
	if _dgef != nil {
		return nil, ctx, _dgef
	}
	return []*Block{_egfc}, ctx, nil
}

// SetBorderColor sets the cell's border color.
func (_ceef *TableCell) SetBorderColor(col Color) {
	_ceef._efcg = col
	_ceef._ebgec = col
	_ceef._fcec = col
	_ceef._afddf = col
}

// Height returns the height of the line.
func (_gaageg *Line) Height() float64 {
	_effb := _gaageg._bfe
	if _gaageg._ccg == _gaageg._dfcae {
		_effb /= 2
	}
	return _cd.Abs(_gaageg._feafb-_gaageg._eggf) + _effb
}

// VectorDrawable is a Drawable with a specified width and height.
type VectorDrawable interface {
	Drawable

	// Width returns the width of the Drawable.
	Width() float64

	// Height returns the height of the Drawable.
	Height() float64
}

// SetForms adds an Acroform to a PDF file.  Sets the specified form for writing.
func (_bde *Creator) SetForms(form *_ga.PdfAcroForm) error { _bde._caff = form; return nil }

// GeneratePageBlocks draws the filled curve on page blocks.
func (_bdgf *FilledCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_defde := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_gfef, _, _ffgdd := _bdgf.draw(_defde, "")
	if _ffgdd != nil {
		return nil, ctx, _ffgdd
	}
	_ffgdd = _defde.addContentsByString(string(_gfef))
	if _ffgdd != nil {
		return nil, ctx, _ffgdd
	}
	return []*Block{_defde}, ctx, nil
}

// BorderWidth returns the border width of the ellipse.
func (_bagad *Ellipse) BorderWidth() float64 { return _bagad._eeaa }

// EnablePageWrap controls whether the division is wrapped across pages.
// If disabled, the division is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the division is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
// Currently, page wrapping can only be disabled for vertical divisions.
func (_bdfc *Division) EnablePageWrap(enable bool) { _bdfc._dccg = enable }

type cmykColor struct{ _bdge, _eegd, _ebe, _agg float64 }

// GraphicSVG represents a drawable graphic SVG.
// It is used to render the graphic SVG components using a creator instance.
type GraphicSVG struct {
	_debae *_ca.GraphicSVG
	_dabg  Positioning
	_dfgae float64
	_egbf  float64
	_fgfa  Margins
}

// SellerAddress returns the seller address used in the invoice template.
func (_fecf *Invoice) SellerAddress() *InvoiceAddress { return _fecf._dceeg }

// SetTitleStyle sets the style properties of the invoice title.
func (_egefa *Invoice) SetTitleStyle(style TextStyle) { _egefa._adgf = style }

func (_abad *Image) makeXObject() error {
	_cbde, _effd := _ga.NewXObjectImageFromImage(_abad._dfad, nil, _abad._cbce)
	if _effd != nil {
		_bcd.Log.Error("\u0046\u0061\u0069le\u0064\u0020\u0074\u006f\u0020\u0063\u0072\u0065\u0061t\u0065 \u0078o\u0062j\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _effd)
		return _effd
	}
	_abad._ebfc = _cbde
	return nil
}

func (_bcccc *templateProcessor) parseFitModeAttr(_egbba, _cccg string) FitMode {
	_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u0069\u0074\u0020\u006do\u0064\u0065\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _egbba, _cccg)
	_cecbf := map[string]FitMode{"\u006e\u006f\u006e\u0065": FitModeNone, "\u0066\u0069\u006c\u006c\u002d\u0077\u0069\u0064\u0074\u0068": FitModeFillWidth}[_cccg]
	return _cecbf
}

// FilledCurve represents a closed path of Bezier curves with a border and fill.
type FilledCurve struct {
	_gefc         []_ff.CubicBezierCurve
	FillEnabled   bool
	_fbgc         Color
	BorderEnabled bool
	BorderWidth   float64
	_afe          Color
}

// SetPos sets the position of the graphic svg to the specified coordinates.
// This method sets the graphic svg to use absolute positioning.
func (_cgde *GraphicSVG) SetPos(x, y float64) {
	_cgde._dabg = PositionAbsolute
	_cgde._dfgae = x
	_cgde._egbf = y
}

// Height returns the height of the chart.
func (_fca *Chart) Height() float64 { return float64(_fca._dece.Height()) }

func _caaf(_fgaag string) (*Image, error) {
	_bdfe, _afc := _e.Open(_fgaag)
	if _afc != nil {
		return nil, _afc
	}
	defer _bdfe.Close()
	_aggd, _afc := _ga.ImageHandling.Read(_bdfe)
	if _afc != nil {
		_bcd.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _afc)
		return nil, _afc
	}
	return _cfcc(_aggd)
}

// Date returns the invoice date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_gfgd *Invoice) Date() (*InvoiceCell, *InvoiceCell) { return _gfgd._afdb[0], _gfgd._afdb[1] }

func (_efcf *templateProcessor) nodeError(_ffag *templateNode, _bcgdd string, _cdfb ...interface{}) error {
	return _g.Errorf(_efcf.getNodeErrorLocation(_ffag, _bcgdd, _cdfb...))
}

// Lines returns all the rows of the invoice line items table.
func (_ceggf *Invoice) Lines() [][]*InvoiceCell { return _ceggf._gdcac }

// SetEncoder sets the encoding/compression mechanism for the image.
func (_ecaa *Image) SetEncoder(encoder _cc.StreamEncoder) { _ecaa._cbce = encoder }

// NewPageBreak create a new page break.
func (_dbe *Creator) NewPageBreak() *PageBreak { return _dbae() }

func (_cfad *Table) clone() *Table {
	_gecff := *_cfad
	_gecff._ecab = make([]float64, len(_cfad._ecab))
	copy(_gecff._ecab, _cfad._ecab)
	_gecff._eeggd = make([]float64, len(_cfad._eeggd))
	copy(_gecff._eeggd, _cfad._eeggd)
	_gecff._gebec = make([]*TableCell, 0, len(_cfad._gebec))
	for _, _eafg := range _cfad._gebec {
		_ggbdb := *_eafg
		_ggbdb._ccbf = &_gecff
		_gecff._gebec = append(_gecff._gebec, &_ggbdb)
	}
	return &_gecff
}

// SetLink makes the line an internal link.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
func (_cbaag *TOCLine) SetLink(page int64, x, y float64) {
	_cbaag._dfbbb = x
	_cbaag._gcab = y
	_cbaag._dfbfg = page
	_fgcfc := _cbaag._eedce._dbgf.Color
	_cbaag.Number.Style.Color = _fgcfc
	_cbaag.Title.Style.Color = _fgcfc
	_cbaag.Separator.Style.Color = _fgcfc
	_cbaag.Page.Style.Color = _fgcfc
}

// FillColor returns the fill color of the rectangle.
func (_dcgg *Rectangle) FillColor() Color { return _dcgg._gbfa }

// GetOptimizer returns current PDF optimizer.
func (_cbga *Creator) GetOptimizer() _ga.Optimizer { return _cbga._adca }

// SetFitMode sets the fit mode of the rectangle.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_aceaf *Rectangle) SetFitMode(fitMode FitMode) { _aceaf._efbb = fitMode }

func (_cbfgg *templateProcessor) parseInt64Attr(_beef, _fdae string) int64 {
	_bcd.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0069\u006e\u0074\u00364 a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _beef, _fdae)
	_aded, _ := _aa.ParseInt(_fdae, 10, 64)
	return _aded
}

// SetPos sets the absolute position. Changes object positioning to absolute.
func (_daab *Image) SetPos(x, y float64) {
	_daab._geca = PositionAbsolute
	_daab._fdab = x
	_daab._egca = y
}

func _eega(_agffd []*ColorPoint) *LinearShading {
	return &LinearShading{_cfbfe: &shading{_cfec: ColorWhite, _cebg: false, _fedf: []bool{false, false}, _ccaf: _agffd}, _ebfea: &_ga.PdfRectangle{}}
}

// Link returns link information for this line.
func (_fgde *TOCLine) Link() (_bgce int64, _eaaab, _egfb float64) {
	return _fgde._dfbfg, _fgde._dfbbb, _fgde._gcab
}

const (
	FitModeNone FitMode = iota
	FitModeFillWidth
)

// Width returns the width of the ellipse.
func (_egdgb *Ellipse) Width() float64 { return _egdgb._eded }

var (
	PageSizeA3     = PageSize{297 * PPMM, 420 * PPMM}
	PageSizeA4     = PageSize{210 * PPMM, 297 * PPMM}
	PageSizeA5     = PageSize{148 * PPMM, 210 * PPMM}
	PageSizeLetter = PageSize{8.5 * PPI, 11 * PPI}
	PageSizeLegal  = PageSize{8.5 * PPI, 14 * PPI}
)

// AddSection adds a new content section at the end of the invoice.
func (_fdfg *Invoice) AddSection(title, content string) {
	_fdfg._eefc = append(_fdfg._eefc, [2]string{title, content})
}

// InvoiceAddress contains contact information that can be displayed
// in an invoice. It is used for the seller and buyer information in the
// invoice template.
type InvoiceAddress struct {
	Heading string
	Name    string
	Street  string
	Street2 string
	Zip     string
	City    string
	State   string
	Country string
	Phone   string
	Email   string

	// Separator defines the separator between different address components,
	// such as the city, state and zip code. It defaults to ", " when the
	// field is an empty string.
	Separator string

	// If enabled, the Phone field label (`Phone: `) is not displayed.
	HidePhoneLabel bool

	// If enabled, the Email field label (`Email: `) is not displayed.
	HideEmailLabel bool
}

const (
	CellBorderSideLeft CellBorderSide = iota
	CellBorderSideRight
	CellBorderSideTop
	CellBorderSideBottom
	CellBorderSideAll
)

// Lines returns all the lines the table of contents has.
func (_gbefd *TOC) Lines() []*TOCLine { return _gbefd._gcgf }

// String implements error interface.
func (_ddfg UnsupportedRuneError) Error() string { return _ddfg.Message }

func _fecea(_abf VectorDrawable, _agge float64) float64 {
	switch _aggec := _abf.(type) {
	case *Paragraph:
		if _aggec._dbfb {
			_aggec.SetWidth(_agge - _aggec._bcggf.Left - _aggec._bcggf.Right)
		}
		return _aggec.Height() + _aggec._bcggf.Top + _aggec._bcggf.Bottom
	case *StyledParagraph:
		if _aggec._bfcgb {
			_aggec.SetWidth(_agge - _aggec._fadcg.Left - _aggec._fadcg.Right)
		}
		return _aggec.Height() + _aggec._fadcg.Top + _aggec._fadcg.Bottom
	case *Image:
		_aggec.applyFitMode(_agge)
		return _aggec.Height() + _aggec._eaac.Top + _aggec._eaac.Bottom
	case *Rectangle:
		_aggec.applyFitMode(_agge)
		return _aggec.Height() + _aggec._eaaa.Top + _aggec._eaaa.Bottom + _aggec._bffc
	case *Ellipse:
		_aggec.applyFitMode(_agge)
		return _aggec.Height() + _aggec._abgc.Top + _aggec._abgc.Bottom
	case *Division:
		return _aggec.ctxHeight(_agge) + _aggec._acaa.Top + _aggec._acaa.Bottom + _aggec._cgbb.Top + _aggec._cgbb.Bottom
	case *Table:
		_aggec.updateRowHeights(_agge - _aggec._bebd.Left - _aggec._bebd.Right)
		return _aggec.Height() + _aggec._bebd.Top + _aggec._bebd.Bottom
	case *List:
		return _aggec.ctxHeight(_agge) + _aggec._fcag.Top + _aggec._fcag.Bottom
	case marginDrawable:
		_, _, _baga, _fbaag := _aggec.GetMargins()
		return _aggec.Height() + _baga + _fbaag
	default:
		return _aggec.Height()
	}
}

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_adebc *Paragraph) SetTextAlignment(align TextAlignment) { _adebc._caea = align }

// Scale scales Image by a constant factor, both width and height.
func (_ceda *Image) Scale(xFactor, yFactor float64) {
	_ceda._ggda = xFactor * _ceda._ggda
	_ceda._cgdb = yFactor * _ceda._cgdb
}

// Finalize renders all blocks to the creator pages. In addition, it takes care
// of adding headers and footers, as well as generating the front page,
// table of contents and outlines.
// Finalize is automatically called before writing the document out. Calling the
// method manually can be useful when adding external pages to the creator,
// using the AddPage method, as it renders all creator blocks to the added
// pages, without having to write the document out.
// NOTE: TOC and outlines are generated only if the AddTOC and AddOutlines
// fields of the creator are set to true (enabled by default). Furthermore, TOCs
// and outlines without content are skipped. TOC and outline content is
// added automatically when using the chapter component. TOCs and outlines can
// also be set externally, using the SetTOC and SetOutlineTree methods.
// Finalize should only be called once, after all draw calls have taken place,
// as it will return immediately if the creator instance has been finalized.
func (_cdbe *Creator) Finalize() error {
	if _cdbe._ccfd {
		return nil
	}
	_ebab := len(_cdbe._agfg)
	_aegb := 0
	if _cdbe._cge != nil {
		_gbc := *_cdbe
		_cdbe._agfg = nil
		_cdbe._abd = nil
		_cdbe.initContext()
		_bebc := FrontpageFunctionArgs{PageNum: 1, TotalPages: _ebab}
		_cdbe._cge(_bebc)
		_aegb += len(_cdbe._agfg)
		_cdbe._agfg = _gbc._agfg
		_cdbe._abd = _gbc._abd
	}
	if _cdbe.AddTOC {
		_cdbe.initContext()
		_cdbe._ggab.Page = _aegb + 1
		if _cdbe.CustomTOC && _cdbe._caf != nil {
			_dcgf := *_cdbe
			_cdbe._agfg = nil
			_cdbe._abd = nil
			if _eafc := _cdbe._caf(_cdbe._edad); _eafc != nil {
				return _eafc
			}
			_aegb += len(_cdbe._agfg)
			_cdbe._agfg = _dcgf._agfg
			_cdbe._abd = _dcgf._abd
		} else {
			if _cdbe._caf != nil {
				if _cagd := _cdbe._caf(_cdbe._edad); _cagd != nil {
					return _cagd
				}
			}
			_eage, _, _bbgc := _cdbe._edad.GeneratePageBlocks(_cdbe._ggab)
			if _bbgc != nil {
				_bcd.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073: \u0025\u0076", _bbgc)
				return _bbgc
			}
			_aegb += len(_eage)
		}
		_fagef := _cdbe._edad.Lines()
		for _, _gefg := range _fagef {
			_bdgg, _dfc := _aa.Atoi(_gefg.Page.Text)
			if _dfc != nil {
				continue
			}
			_gefg.Page.Text = _aa.Itoa(_bdgg + _aegb)
			_gefg._dfbfg += int64(_aegb)
		}
	}
	_aedf := false
	var _ffgdc []*_ga.PdfPage
	if _cdbe._cge != nil {
		_gffd := *_cdbe
		_cdbe._agfg = nil
		_cdbe._abd = nil
		_efbc := FrontpageFunctionArgs{PageNum: 1, TotalPages: _ebab}
		_cdbe._cge(_efbc)
		_ebab += len(_cdbe._agfg)
		_ffgdc = _cdbe._agfg
		_cdbe._agfg = append(_cdbe._agfg, _gffd._agfg...)
		_cdbe._abd = _gffd._abd
		_aedf = true
	}
	var _bbe []*_ga.PdfPage
	if _cdbe.AddTOC {
		_cdbe.initContext()
		if _cdbe.CustomTOC && _cdbe._caf != nil {
			_baagf := *_cdbe
			_cdbe._agfg = nil
			_cdbe._abd = nil
			if _egdg := _cdbe._caf(_cdbe._edad); _egdg != nil {
				_bcd.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _egdg)
				return _egdg
			}
			_bbe = _cdbe._agfg
			_ebab += len(_bbe)
			_cdbe._agfg = _baagf._agfg
			_cdbe._abd = _baagf._abd
		} else {
			if _cdbe._caf != nil {
				if _ead := _cdbe._caf(_cdbe._edad); _ead != nil {
					_bcd.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _ead)
					return _ead
				}
			}
			_gad, _, _ := _cdbe._edad.GeneratePageBlocks(_cdbe._ggab)
			for _, _baca := range _gad {
				_baca.SetPos(0, 0)
				_ebab++
				_dgge := _cdbe.newPage()
				_bbe = append(_bbe, _dgge)
				_cdbe.setActivePage(_dgge)
				_cdbe.Draw(_baca)
			}
		}
		if _aedf {
			_eecb := _ffgdc
			_effg := _cdbe._agfg[len(_ffgdc):]
			_cdbe._agfg = append([]*_ga.PdfPage{}, _eecb...)
			_cdbe._agfg = append(_cdbe._agfg, _bbe...)
			_cdbe._agfg = append(_cdbe._agfg, _effg...)
		} else {
			_cdbe._agfg = append(_bbe, _cdbe._agfg...)
		}
	}
	if _cdbe._bcb != nil && _cdbe.AddOutlines {
		var _accd func(_eegee *_ga.OutlineItem)
		_accd = func(_fccg *_ga.OutlineItem) {
			_fccg.Dest.Page += int64(_aegb)
			if _eabb := int(_fccg.Dest.Page); _eabb >= 0 && _eabb < len(_cdbe._agfg) {
				_fccg.Dest.PageObj = _cdbe._agfg[_eabb].GetPageAsIndirectObject()
			} else {
				_bcd.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _eabb)
			}
			_fccg.Dest.Y = _cdbe._aafc - _fccg.Dest.Y
			_gac := _fccg.Items()
			for _, _cgaa := range _gac {
				_accd(_cgaa)
			}
		}
		_agbf := _cdbe._bcb.Items()
		for _, _gbca := range _agbf {
			_accd(_gbca)
		}
		if _cdbe.AddTOC {
			var _fff int
			if _aedf {
				_fff = len(_ffgdc)
			}
			_gdb := _ga.NewOutlineDest(int64(_fff), 0, _cdbe._aafc)
			if _fff >= 0 && _fff < len(_cdbe._agfg) {
				_gdb.PageObj = _cdbe._agfg[_fff].GetPageAsIndirectObject()
			} else {
				_bcd.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _fff)
			}
			_cdbe._bcb.Insert(0, _ga.NewOutlineItem("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073", _gdb))
		}
	}
	for _gbcc, _fafa := range _cdbe._agfg {
		_cdbe.setActivePage(_fafa)
		if _cdbe._aagf != nil {
			_dcaa, _bcc, _eaad := _fafa.Size()
			if _eaad != nil {
				return _eaad
			}
			_gfcd := PageFinalizeFunctionArgs{PageNum: _gbcc + 1, PageWidth: _dcaa, PageHeight: _bcc, TOCPages: len(_bbe), TotalPages: _ebab}
			if _cfcb := _cdbe._aagf(_gfcd); _cfcb != nil {
				_bcd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0070\u0061\u0067\u0065\u0020\u0066\u0069\u006e\u0061\u006c\u0069\u007a\u0065 \u0063\u0061\u006c\u006c\u0062\u0061\u0063k\u003a\u0020\u0025\u0076", _cfcb)
				return _cfcb
			}
		}
		if _cdbe._bebfg != nil {
			_bdcc := NewBlock(_cdbe._bagf, _cdbe._eaf.Top)
			_bggd := HeaderFunctionArgs{PageNum: _gbcc + 1, TotalPages: _ebab}
			_cdbe._bebfg(_bdcc, _bggd)
			_bdcc.SetPos(0, 0)
			if _fgab := _cdbe.Draw(_bdcc); _fgab != nil {
				_bcd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0068e\u0061\u0064\u0065\u0072\u003a\u0020\u0025v", _fgab)
				return _fgab
			}
		}
		if _cdbe._acfe != nil {
			_cbec := NewBlock(_cdbe._bagf, _cdbe._eaf.Bottom)
			_ebdb := FooterFunctionArgs{PageNum: _gbcc + 1, TotalPages: _ebab}
			_cdbe._acfe(_cbec, _ebdb)
			_cbec.SetPos(0, _cdbe._aafc-_cbec._cce)
			if _fdbga := _cdbe.Draw(_cbec); _fdbga != nil {
				_bcd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0066o\u006f\u0074\u0065\u0072\u003a\u0020\u0025v", _fdbga)
				return _fdbga
			}
		}
		_cfac, _cgba := _cdbe._bbf[_fafa]
		if _bcdc, _ebag := _cdbe._fgd[_fafa]; _ebag {
			if _cgba {
				_cfac.transformBlock(_bcdc)
			}
			if _cgbg := _bcdc.drawToPage(_fafa); _cgbg != nil {
				_bcd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020%\u0064\u0020\u0062\u006c\u006f\u0063\u006bs\u003a\u0020\u0025\u0076", _gbcc+1, _cgbg)
				return _cgbg
			}
		}
		if _cgba {
			if _aca := _cfac.transformPage(_fafa); _aca != nil {
				_bcd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0074\u0072\u0061\u006e\u0073f\u006f\u0072\u006d\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0076", _aca)
				return _aca
			}
		}
	}
	_cdbe._ccfd = true
	return nil
}

// TextAlignment options for paragraph.
type TextAlignment int

// DueDate returns the invoice due date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_cedd *Invoice) DueDate() (*InvoiceCell, *InvoiceCell) { return _cedd._aagc[0], _cedd._aagc[1] }
func (_fdad *StyledParagraph) wrapText() error               { return _fdad.wrapChunks(true) }

// SetHorizontalAlignment sets the cell's horizontal alignment of content.
// Can be one of:
// - CellHorizontalAlignmentLeft
// - CellHorizontalAlignmentCenter
// - CellHorizontalAlignmentRight
func (_bbfd *TableCell) SetHorizontalAlignment(halign CellHorizontalAlignment) { _bbfd._dbdc = halign }

// Width returns the width of the line.
// NOTE: Depending on the fit mode the line is set to use, its width may be
// calculated at runtime (e.g. when using FitModeFillWidth).
func (_egee *Line) Width() float64 { return _cd.Abs(_egee._dfcae - _egee._ccg) }

// TOC returns the table of contents component of the creator.
func (_dbg *Creator) TOC() *TOC { return _dbg._edad }

func (_ecff *listItem) ctxHeight(_bceb float64) float64 {
	var _eaed float64
	switch _fdca := _ecff._cecd.(type) {
	case *Paragraph:
		if _fdca._dbfb {
			_fdca.SetWidth(_bceb - _fdca._bcggf.Horizontal())
		}
		_eaed = _fdca.Height() + _fdca._bcggf.Vertical()
		_eaed += 0.5 * _fdca._dbad * _fdca._agdb
	case *StyledParagraph:
		if _fdca._bfcgb {
			_fdca.SetWidth(_bceb - _fdca._fadcg.Horizontal())
		}
		_eaed = _fdca.Height() + _fdca._fadcg.Vertical()
		_eaed += 0.5 * _fdca.getTextHeight()
	case *List:
		_dbcd := _bceb - _ecff._feced.Width() - _fdca._fcag.Horizontal() - _fdca._addd
		_eaed = _fdca.ctxHeight(_dbcd) + _fdca._fcag.Vertical()
	case *Image:
		_eaed = _fdca.Height() + _fdca._eaac.Vertical()
	case *Division:
		_eegb := _bceb - _ecff._feced.Width() - _fdca._acaa.Horizontal()
		_eaed = _fdca.ctxHeight(_eegb) + _fdca._acaa.Vertical()
	case *Table:
		_edfc := _bceb - _ecff._feced.Width() - _fdca._bebd.Horizontal()
		_fdca.updateRowHeights(_edfc)
		_eaed = _fdca.Height() + _fdca._bebd.Vertical()
	default:
		_eaed = _ecff._cecd.Height()
	}
	return _eaed
}

// MultiColCell makes a new cell with the specified column span and inserts it
// into the table at the current position.
func (_bcga *Table) MultiColCell(colspan int) *TableCell { return _bcga.MultiCell(1, colspan) }

// SetText sets the text content of the Paragraph.
func (_dacb *Paragraph) SetText(text string) { _dacb._bfgg = text }

func _cdae(_cae, _edcba *_ga.PdfPageResources) error {
	_bba, _ := _cae.GetColorspaces()
	if _bba != nil && len(_bba.Colorspaces) > 0 {
		for _beg, _aef := range _bba.Colorspaces {
			_bgg := *_cc.MakeName(_beg)
			if _edcba.HasColorspaceByName(_bgg) {
				continue
			}
			_dga := _edcba.SetColorspaceByName(_bgg, _aef)
			if _dga != nil {
				return _dga
			}
		}
	}
	return nil
}

func (_agbef *Division) drawBackground(_afg []*Block, _fgcg, _bgaac DrawContext, _debd bool) ([]*Block, error) {
	_dafd := len(_afg)
	if _dafd == 0 || _agbef._eeaf == nil {
		return _afg, nil
	}
	_eaea := make([]*Block, 0, len(_afg))
	for _cbeca, _fcade := range _afg {
		var (
			_efd  = _agbef._eeaf.BorderRadiusTopLeft
			_aecb = _agbef._eeaf.BorderRadiusTopRight
			_aacf = _agbef._eeaf.BorderRadiusBottomLeft
			_gdeg = _agbef._eeaf.BorderRadiusBottomRight
		)
		_bfdd := _fgcg
		_bfdd.Page += _cbeca
		if _cbeca == 0 {
			if _debd {
				_eaea = append(_eaea, _fcade)
				continue
			}
			if _dafd == 1 {
				_bfdd.Height = _bgaac.Y - _fgcg.Y
			}
		} else {
			_bfdd.X = _bfdd.Margins.Left + _agbef._acaa.Left
			_bfdd.Y = _bfdd.Margins.Top
			_bfdd.Width = _bfdd.PageWidth - _bfdd.Margins.Left - _bfdd.Margins.Right - _agbef._acaa.Left - _agbef._acaa.Right
			if _cbeca == _dafd-1 {
				_bfdd.Height = _bgaac.Y - _bfdd.Margins.Top - _agbef._acaa.Top
			} else {
				_bfdd.Height = _bfdd.PageHeight - _bfdd.Margins.Top - _bfdd.Margins.Bottom
			}
			if !_debd {
				_efd = 0
				_aecb = 0
			}
		}
		if _dafd > 1 && _cbeca != _dafd-1 {
			_aacf = 0
			_gdeg = 0
		}
		_gec := _cfeg(_bfdd.X, _bfdd.Y, _bfdd.Width, _bfdd.Height)
		_gec.SetFillColor(_agbef._eeaf.FillColor)
		_gec.SetBorderColor(_agbef._eeaf.BorderColor)
		_gec.SetBorderWidth(_agbef._eeaf.BorderSize)
		_gec.SetBorderRadius(_efd, _aecb, _aacf, _gdeg)
		_cbef, _, _fgbe := _gec.GeneratePageBlocks(_bfdd)
		if _fgbe != nil {
			return nil, _fgbe
		}
		if len(_cbef) == 0 {
			continue
		}
		_gage := _cbef[0]
		if _fgbe = _gage.mergeBlocks(_fcade); _fgbe != nil {
			return nil, _fgbe
		}
		_eaea = append(_eaea, _gage)
	}
	return _eaea, nil
}

func (_adba *TableCell) cloneProps(_ffff VectorDrawable) *TableCell {
	_feeeb := *_adba
	_feeeb._bfef = _ffff
	return &_feeeb
}

func _faacc(_gaab *templateProcessor, _ccfa *templateNode) (interface{}, error) {
	return _gaab.parseChart(_ccfa)
}

func _gcd(_ddc string) string {
	_egd := _adg.FindAllString(_ddc, -1)
	if len(_egd) == 0 {
		_ddc = _ddc + "\u0030"
	} else {
		_gffe, _abe := _aa.Atoi(_egd[len(_egd)-1])
		if _abe != nil {
			_bcd.Log.Debug("\u0045r\u0072\u006f\u0072 \u0063\u006f\u006ev\u0065rt\u0069\u006e\u0067\u0020\u0064\u0069\u0067i\u0074\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006e\u0061\u006de,\u0020f\u0061\u006c\u006c\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0062a\u0073\u0069\u0063\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a \u0025\u0076", _abe)
			_ddc = _ddc + "\u0030"
		} else {
			_gffe++
			_fbf := _a.LastIndex(_ddc, _egd[len(_egd)-1])
			if _fbf == -1 {
				_ddc = _g.Sprintf("\u0025\u0073\u0025\u0064", _ddc[:len(_ddc)-1], _gffe)
			} else {
				_ddc = _ddc[:_fbf] + _aa.Itoa(_gffe)
			}
		}
	}
	return _ddc
}

// AppendCurve appends a Bezier curve to the filled curve.
func (_cacd *FilledCurve) AppendCurve(curve _ff.CubicBezierCurve) *FilledCurve {
	_cacd._gefc = append(_cacd._gefc, curve)
	return _cacd
}

func _ebege(_gdegg *templateProcessor, _cbdbd *templateNode) (interface{}, error) {
	return _gdegg.parseStyledParagraph(_cbdbd)
}

func _ebda(_fabg *Chapter, _bff *TOC, _ddfa *_ga.Outline, _gcc string, _faa int, _dbd TextStyle) *Chapter {
	var _baag uint = 1
	if _fabg != nil {
		_baag = _fabg._fge + 1
	}
	_gea := &Chapter{_baef: _faa, _gbb: _gcc, _ebgde: true, _gge: true, _eagg: _fabg, _aaccb: _bff, _eebd: _ddfa, _gefd: []Drawable{}, _fge: _baag}
	_gdgg := _agcdg(_gea.headingText(), _dbd)
	_gdgg.SetFont(_dbd.Font)
	_gdgg.SetFontSize(_dbd.FontSize)
	_gea._bgfc = _gdgg
	return _gea
}

func (_cacdg *Paragraph) getTextWidth() float64 {
	_ddegg := 0.0
	for _, _gdde := range _cacdg._bfgg {
		if _gdde == '\u000A' {
			continue
		}
		_ggeg, _adffb := _cacdg._bfbd.GetRuneMetrics(_gdde)
		if !_adffb {
			_bcd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _gdde, _gdde)
			return -1
		}
		_ddegg += _cacdg._dbad * _ggeg.Wx
	}
	return _ddegg
}

func (_abdd *templateProcessor) run() error {
	_ccga := _dda.NewDecoder(_b.NewReader(_abdd._cebag))
	var _bgacd *templateNode
	for {
		_ccgcf, _acae := _ccga.Token()
		if _acae != nil {
			if _acae == _eb.EOF {
				return nil
			}
			return _acae
		}
		if _ccgcf == nil {
			break
		}
		_eeec, _abcab := _cddb(_ccga)
		_babb := _ccga.InputOffset()
		switch _dagag := _ccgcf.(type) {
		case _dda.StartElement:
			_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006eg\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0074\u0061r\u0074\u0020\u0074\u0061\u0067\u003a\u0020`\u0025\u0073\u0060\u002e", _dagag.Name.Local)
			_gfcca, _dggddb := _cabdda[_dagag.Name.Local]
			if !_dggddb {
				if _abdd._bfgd == "" {
					if _eeec != 0 {
						_bcd.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 \u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0064\u003a\u0025\u0064\u005d", _dagag.Name.Local, _eeec, _abcab)
					} else {
						_bcd.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070i\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u002e\u0020\u005b%\u0064\u005d", _dagag.Name.Local, _babb)
					}
				} else {
					if _eeec != 0 {
						_bcd.Log.Debug("\u0055\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0074e\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0073\u003a\u0025\u0064\u003a\u0025d\u005d", _dagag.Name.Local, _abdd._bfgd, _eeec, _abcab)
					} else {
						_bcd.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 \u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0073\u003a\u0025\u0064\u005d", _dagag.Name.Local, _abdd._bfgd, _babb)
					}
				}
				continue
			}
			_bgacd = &templateNode{_acdf: _dagag, _gbdf: _bgacd, _ffgg: _eeec, _gcdae: _abcab, _cgcf: _babb}
			if _dbfgc := _gfcca._aecbab; _dbfgc != nil {
				_bgacd._defae, _acae = _dbfgc(_abdd, _bgacd)
				if _acae != nil {
					return _acae
				}
			}
		case _dda.EndElement:
			_bcd.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0065\u006e\u0064\u0020\u0074\u0061\u0067\u003a\u0020\u0060\u0025\u0073\u0060\u002e", _dagag.Name.Local)
			if _bgacd != nil {
				if _bgacd._defae != nil {
					if _cccda := _abdd.renderNode(_bgacd); _cccda != nil {
						return _cccda
					}
				}
				_bgacd = _bgacd._gbdf
			}
		case _dda.CharData:
			if _bgacd != nil && _bgacd._defae != nil {
				if _cfccg := _abdd.addNodeText(_bgacd, string(_dagag)); _cfccg != nil {
					return _cfccg
				}
			}
		case _dda.Comment:
			_bcd.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0063\u006f\u006d\u006d\u0065\u006e\u0074\u003a\u0020\u0060\u0025\u0073\u0060\u002e", string(_dagag))
		}
	}
	return nil
}

func (_bfbee *Table) wrapRow(_ffge int, _aebef DrawContext, _bafeb float64) (bool, error) {
	if !_bfbee._ggea {
		return false, nil
	}
	var (
		_edbea = _bfbee._gebec[_ffge]
		_adbfg = -1
		_abbd  []*TableCell
		_decg  float64
		_bedb  bool
		_gecee = make([]float64, 0, len(_bfbee._eeggd))
	)
	_cebgg := func(_cfed *TableCell, _agbec VectorDrawable, _cegf bool) *TableCell {
		_gcgb := *_cfed
		_gcgb._bfef = _agbec
		if _cegf {
			_gcgb._bgdca++
		}
		return &_gcgb
	}
	_abcd := func(_fccbfa int, _effab VectorDrawable) {
		var _bgdf float64 = -1
		if _effab == nil {
			if _efgde := _gecee[_fccbfa-_ffge]; _efgde > _aebef.Height {
				_effab = _bfbee._gebec[_fccbfa]._bfef
				_bfbee._gebec[_fccbfa]._bfef = nil
				_gecee[_fccbfa-_ffge] = 0
				_bgdf = _efgde
			}
		}
		_dcabc := _cebgg(_bfbee._gebec[_fccbfa], _effab, true)
		_abbd = append(_abbd, _dcabc)
		if _bgdf < 0 {
			_bgdf = _dcabc.height(_aebef.Width)
		}
		if _bgdf > _decg {
			_decg = _bgdf
		}
	}
	for _dfea := _ffge; _dfea < len(_bfbee._gebec); _dfea++ {
		_efaea := _bfbee._gebec[_dfea]
		if _edbea._bgdca != _efaea._bgdca {
			_adbfg = _dfea
			break
		}
		_aebef.Width = _efaea.width(_bfbee._eeggd, _bafeb)
		_gcgbf := _efaea.height(_aebef.Width)
		var _bfbaa VectorDrawable
		switch _gadae := _efaea._bfef.(type) {
		case *StyledParagraph:
			if _gcgbf > _aebef.Height {
				_aade := _aebef
				_aade.Height = _cd.Floor(_aebef.Height - _gadae._fadcg.Top - _gadae._fadcg.Bottom - 0.5*_gadae.getTextHeight())
				_fgdd, _degga, _ddggb := _gadae.split(_aade)
				if _ddggb != nil {
					return false, _ddggb
				}
				if _fgdd != nil && _degga != nil {
					_gadae = _fgdd
					_efaea = _cebgg(_efaea, _fgdd, false)
					_bfbee._gebec[_dfea] = _efaea
					_bfbaa = _degga
					_bedb = true
				}
				_gcgbf = _efaea.height(_aebef.Width)
			}
		case *Division:
			if _gcgbf > _aebef.Height {
				_bfdgg := _aebef
				_bfdgg.Height = _cd.Floor(_aebef.Height - _gadae._acaa.Top - _gadae._acaa.Bottom)
				_gfeg, _gfbg := _gadae.split(_bfdgg)
				if _gfeg != nil && _gfbg != nil {
					_gadae = _gfeg
					_efaea = _cebgg(_efaea, _gfeg, false)
					_bfbee._gebec[_dfea] = _efaea
					_bfbaa = _gfbg
					_bedb = true
					if _gfeg._eeaf != nil {
						_gfeg._eeaf.BorderRadiusBottomLeft = 0
						_gfeg._eeaf.BorderRadiusBottomRight = 0
					}
					if _gfbg._eeaf != nil {
						_gfbg._eeaf.BorderRadiusTopLeft = 0
						_gfbg._eeaf.BorderRadiusTopRight = 0
					}
					_gcgbf = _efaea.height(_aebef.Width)
				}
			}
		case *List:
			if _gcgbf > _aebef.Height {
				_gacf := _aebef
				_gacf.Height = _cd.Floor(_aebef.Height - _gadae._fcag.Vertical())
				_cffa, _dfcgc := _gadae.split(_gacf)
				if _cffa != nil {
					_gadae = _cffa
					_efaea = _cebgg(_efaea, _cffa, false)
					_bfbee._gebec[_dfea] = _efaea
				}
				if _dfcgc != nil {
					_bfbaa = _dfcgc
					_bedb = true
				}
				_gcgbf = _efaea.height(_aebef.Width)
			}
		}
		_gecee = append(_gecee, _gcgbf)
		if _bedb {
			if _abbd == nil {
				_abbd = make([]*TableCell, 0, len(_bfbee._eeggd))
				for _ggde := _ffge; _ggde < _dfea; _ggde++ {
					_abcd(_ggde, nil)
				}
			}
			_abcd(_dfea, _bfbaa)
		}
	}
	var _gggbf float64
	for _, _fabba := range _gecee {
		if _fabba > _gggbf {
			_gggbf = _fabba
		}
	}
	if _bedb && _gggbf < _aebef.Height {
		if _adbfg < 0 {
			_adbfg = len(_bfbee._gebec)
		}
		_cbbb := _bfbee._gebec[_adbfg-1]._bgdca + _bfbee._gebec[_adbfg-1]._fddc - 1
		for _fgfaa := _adbfg; _fgfaa < len(_bfbee._gebec); _fgfaa++ {
			_bfbee._gebec[_fgfaa]._bgdca++
		}
		_bfbee._gebec = append(_bfbee._gebec[:_adbfg], append(_abbd, _bfbee._gebec[_adbfg:]...)...)
		_bfbee._ecab = append(_bfbee._ecab[:_cbbb], append([]float64{_decg}, _bfbee._ecab[_cbbb:]...)...)
		_bfbee._ecab[_edbea._bgdca+_edbea._fddc-2] = _gggbf
	}
	return _bedb, nil
}

func (_eebg *Line) computeCoords(_gbbf DrawContext) (_bbbb, _fggg, _agdeg, _dbgg float64) {
	_bbbb = _gbbf.X
	_agdeg = _bbbb + _eebg._dfcae - _eebg._ccg
	_affa := _eebg._bfe
	if _eebg._ccg == _eebg._dfcae {
		_affa /= 2
	}
	if _eebg._eggf < _eebg._feafb {
		_fggg = _gbbf.PageHeight - _gbbf.Y - _affa
		_dbgg = _fggg - _eebg._feafb + _eebg._eggf
	} else {
		_dbgg = _gbbf.PageHeight - _gbbf.Y - _affa
		_fggg = _dbgg - _eebg._eggf + _eebg._feafb
	}
	switch _eebg._cbab {
	case FitModeFillWidth:
		_agdeg = _bbbb + _gbbf.Width
	}
	return _bbbb, _fggg, _agdeg, _dbgg
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_eabdd *TOC) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ababc := ctx
	_efgg, ctx, _fegg := _eabdd._aeeb.GeneratePageBlocks(ctx)
	if _fegg != nil {
		return _efgg, ctx, _fegg
	}
	for _, _fcagc := range _eabdd._gcgf {
		_fdda := _fcagc._dfbfg
		if !_eabdd._bbee {
			_fcagc._dfbfg = 0
		}
		_eggfc, _bfccd, _ccafg := _fcagc.GeneratePageBlocks(ctx)
		_fcagc._dfbfg = _fdda
		if _ccafg != nil {
			return _efgg, ctx, _ccafg
		}
		if len(_eggfc) < 1 {
			continue
		}
		_efgg[len(_efgg)-1].mergeBlocks(_eggfc[0])
		_efgg = append(_efgg, _eggfc[1:]...)
		ctx = _bfccd
	}
	if _eabdd._dfdc.IsRelative() {
		ctx.X = _ababc.X
	}
	if _eabdd._dfdc.IsAbsolute() {
		return _efgg, _ababc, nil
	}
	return _efgg, ctx, nil
}

func (_badba *TOCLine) prepareParagraph(_fdfcf *StyledParagraph, _badac DrawContext) {
	_gedga := _badba.Title.Text
	if _badba.Number.Text != "" {
		_gedga = "\u0020" + _gedga
	}
	_gedga += "\u0020"
	_adcbd := _badba.Page.Text
	if _adcbd != "" {
		_adcbd = "\u0020" + _adcbd
	}
	_fdfcf._eddc = []*TextChunk{{Text: _badba.Number.Text, Style: _badba.Number.Style, _cgfaa: _badba.getLineLink()}, {Text: _gedga, Style: _badba.Title.Style, _cgfaa: _badba.getLineLink()}, {Text: _adcbd, Style: _badba.Page.Style, _cgfaa: _badba.getLineLink()}}
	_fdfcf.wrapText()
	_dcgac := len(_fdfcf._bddb)
	if _dcgac == 0 {
		return
	}
	_eggdd := _badac.Width*1000 - _fdfcf.getTextLineWidth(_fdfcf._bddb[_dcgac-1])
	_efaca := _fdfcf.getTextLineWidth([]*TextChunk{&_badba.Separator})
	_degcb := int(_eggdd / _efaca)
	_gdaaf := _a.Repeat(_badba.Separator.Text, _degcb)
	_gdcc := _badba.Separator.Style
	_gbbe := _fdfcf.Insert(2, _gdaaf)
	_gbbe.Style = _gdcc
	_gbbe._cgfaa = _badba.getLineLink()
	_eggdd = _eggdd - float64(_degcb)*_efaca
	if _eggdd > 500 {
		_eeea, _bgfae := _gdcc.Font.GetRuneMetrics(' ')
		if _bgfae && _eggdd > _eeea.Wx {
			_gcded := int(_eggdd / _eeea.Wx)
			if _gcded > 0 {
				_cegag := _gdcc
				_cegag.FontSize = 1
				_gbbe = _fdfcf.Insert(2, _a.Repeat("\u0020", _gcded))
				_gbbe.Style = _cegag
				_gbbe._cgfaa = _badba.getLineLink()
			}
		}
	}
}

func (_faad *templateProcessor) parseHorizontalAlignmentAttr(_fcbab, _cfacb string) HorizontalAlignment {
	_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020a\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029.", _fcbab, _cfacb)
	_bdag := map[string]HorizontalAlignment{"\u006c\u0065\u0066\u0074": HorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": HorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": HorizontalAlignmentRight}[_cfacb]
	return _bdag
}

// NoteHeadingStyle returns the style properties used to render the heading of
// the invoice note sections.
func (_fggdg *Invoice) NoteHeadingStyle() TextStyle { return _fggdg._dede }

func (_agcbc *templateProcessor) parseChapter(_cdgdc *templateNode) (interface{}, error) {
	_baaf := _agcbc.creator.NewChapter
	if _cdgdc._gbdf != nil {
		if _adeg, _ggefd := _cdgdc._gbdf._defae.(*Chapter); _ggefd {
			_baaf = _adeg.NewSubchapter
		}
	}
	_fbfg := _baaf("")
	for _, _adaad := range _cdgdc._acdf.Attr {
		_ebfg := _adaad.Value
		switch _becg := _adaad.Name.Local; _becg {
		case "\u0073\u0068\u006f\u0077\u002d\u006e\u0075\u006d\u0062e\u0072\u0069\u006e\u0067":
			_fbfg.SetShowNumbering(_agcbc.parseBoolAttr(_becg, _ebfg))
		case "\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u002d\u0069n\u002d\u0074\u006f\u0063":
			_fbfg.SetIncludeInTOC(_agcbc.parseBoolAttr(_becg, _ebfg))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_bdbb := _agcbc.parseMarginAttr(_becg, _ebfg)
			_fbfg.SetMargins(_bdbb.Left, _bdbb.Right, _bdbb.Top, _bdbb.Bottom)
		default:
			_agcbc.nodeLogDebug(_cdgdc, "\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _becg)
		}
	}
	return _fbfg, nil
}

// Creator is a wrapper around functionality for creating PDF reports and/or adding new
// content onto imported PDF pages, etc.
type Creator struct {
	// Errors keeps error messages that should not interrupt pdf processing and to be checked later.
	Errors []error

	// UnsupportedCharacterReplacement is character that will be used to replace unsupported glyph.
	// The value will be passed to drawing context.
	UnsupportedCharacterReplacement rune
	_agfg                           []*_ga.PdfPage
	_fgd                            map[*_ga.PdfPage]*Block
	_bbf                            map[*_ga.PdfPage]*pageTransformations
	_abd                            *_ga.PdfPage
	_aad                            PageSize
	_ggab                           DrawContext
	_eaf                            Margins
	_bagf, _aafc                    float64
	_age                            int
	_cge                            func(_fcc FrontpageFunctionArgs)
	_caf                            func(_bea *TOC) error
	_bebfg                          func(_fcfa *Block, _ageg HeaderFunctionArgs)
	_acfe                           func(_aaed *Block, _bfd FooterFunctionArgs)
	_aagf                           func(_bddf PageFinalizeFunctionArgs) error
	_cefc                           func(_gaba *_ga.PdfWriter) error
	_ccfd                           bool

	// Controls whether a table of contents will be generated.
	AddTOC bool

	// CustomTOC specifies if the TOC is rendered by the user.
	// When the `CustomTOC` field is set to `true`, the default TOC component is not rendered.
	// Instead the TOC is drawn by the user, in the callback provided to
	// the `Creator.CreateTableOfContents` method.
	// If `CustomTOC` is set to `false`, the callback provided to
	// `Creator.CreateTableOfContents` customizes the style of the automatically generated TOC component.
	CustomTOC bool
	_edad     *TOC

	// Controls whether outlines will be generated.
	AddOutlines bool
	_bcb        *_ga.Outline
	_afag       *_ga.PdfOutlineTreeNode
	_caff       *_ga.PdfAcroForm
	_fdff       _cc.PdfObject
	_adca       _ga.Optimizer
	_abdf       []*_ga.PdfFont
	_ffcg       *_ga.PdfFont
	_cbag       *_ga.PdfFont
}

// NewList creates a new list.
func (_ddfac *Creator) NewList() *List { return _gdcb(_ddfac.NewTextStyle()) }

func (_aefb *templateProcessor) nodeLogError(_gffed *templateNode, _fcaf string, _afcf ...interface{}) {
	_bcd.Log.Error(_aefb.getNodeErrorLocation(_gffed, _fcaf, _afcf...))
}

// Block contains a portion of PDF Page contents. It has a width and a position and can
// be placed anywhere on a Page.  It can even contain a whole Page, and is used in the creator
// where each Drawable object can output one or more blocks, each representing content for separate pages
// (typically needed when Page breaks occur).
type Block struct {
	_fa        *_da.ContentStreamOperations
	_ccf       *_ga.PdfPageResources
	_gfb       Positioning
	_bcde, _gc float64
	_fd        float64
	_cce       float64
	_be        float64
	_bf        Margins
	_fdd       []*_ga.PdfAnnotation
}

func _dcgb(_gafgd *_ga.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 238), Font: _gafgd, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}

// SetWidthTop sets border width for top.
func (_eeg *border) SetWidthTop(bw float64) { _eeg._dca = bw }

// SkipCells skips over a specified number of cells in the table.
func (_dcfde *Table) SkipCells(num int) {
	if num < 0 {
		_bcd.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _cffdd := 0; _cffdd < num; _cffdd++ {
		_dcfde.NewCell()
	}
}

// RotatedSize returns the width and height of the rotated block.
func (_ea *Block) RotatedSize() (float64, float64) {
	_, _, _bacb, _ddd := _aecbg(_ea._fd, _ea._cce, _ea._be)
	return _bacb, _ddd
}

type listItem struct {
	_cecd  VectorDrawable
	_feced TextChunk
}

// Write output of creator to io.Writer interface.
func (_bbd *Creator) Write(ws _eb.Writer) error {
	if _cfff := _bbd.Finalize(); _cfff != nil {
		return _cfff
	}
	_ffad := _ga.NewPdfWriter()
	_ffad.SetOptimizer(_bbd._adca)
	if _bbd._caff != nil {
		_dagb := _ffad.SetForms(_bbd._caff)
		if _dagb != nil {
			_bcd.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _dagb)
			return _dagb
		}
	}
	if _bbd._afag != nil {
		_ffad.AddOutlineTree(_bbd._afag)
	} else if _bbd._bcb != nil && _bbd.AddOutlines {
		_ffad.AddOutlineTree(&_bbd._bcb.ToPdfOutline().PdfOutlineTreeNode)
	}
	if _bbd._fdff != nil {
		if _gefb := _ffad.SetPageLabels(_bbd._fdff); _gefb != nil {
			_bcd.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020C\u006f\u0075\u006c\u0064 no\u0074 s\u0065\u0074\u0020\u0070\u0061\u0067\u0065 l\u0061\u0062\u0065\u006c\u0073\u003a\u0020%\u0076", _gefb)
			return _gefb
		}
	}
	if _bbd._abdf != nil {
		for _, _acfd := range _bbd._abdf {
			_dcad := _acfd.SubsetRegistered()
			if _dcad != nil {
				_bcd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006ft\u0020s\u0075\u0062\u0073\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _dcad)
				return _dcad
			}
		}
	}
	if _bbd._cefc != nil {
		_caee := _bbd._cefc(&_ffad)
		if _caee != nil {
			_bcd.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _caee)
			return _caee
		}
	}
	for _, _efe := range _bbd._agfg {
		_facdf := _ffad.AddPage(_efe)
		if _facdf != nil {
			_bcd.Log.Error("\u0046\u0061\u0069\u006ced\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0050\u0061\u0067\u0065\u003a\u0020%\u0076", _facdf)
			return _facdf
		}
	}
	_cfgf := _ffad.Write(ws)
	if _cfgf != nil {
		return _cfgf
	}
	return nil
}

func (_feb *Chapter) headingText() string {
	_eaae := _feb._gbb
	if _gfee := _feb.headingNumber(); _gfee != "" {
		_eaae = _g.Sprintf("\u0025\u0073\u0020%\u0073", _gfee, _eaae)
	}
	return _eaae
}

// AddTotalLine adds a new line in the invoice totals table.
func (_cfab *Invoice) AddTotalLine(desc, value string) (*InvoiceCell, *InvoiceCell) {
	_fafef := &InvoiceCell{_cfab._cbefc, desc}
	_eacg := &InvoiceCell{_cfab._cbefc, value}
	_cfab._gbece = append(_cfab._gbece, [2]*InvoiceCell{_fafef, _eacg})
	return _fafef, _eacg
}

// SetLineSeparatorStyle sets the style for the separator part of all new
// lines of the table of contents.
func (_deac *TOC) SetLineSeparatorStyle(style TextStyle) { _deac._dbdd = style }

// NewParagraph creates a new text paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_bacad *Creator) NewParagraph(text string) *Paragraph {
	return _agcdg(text, _bacad.NewTextStyle())
}

const (
	DefaultHorizontalScaling = 100
)

func _gbgfg(_gdef, _eaab, _gbcg string, _gdgfa uint, _gffg TextStyle) *TOCLine {
	return _ecbaa(TextChunk{Text: _gdef, Style: _gffg}, TextChunk{Text: _eaab, Style: _gffg}, TextChunk{Text: _gbcg, Style: _gffg}, _gdgfa, _gffg)
}

var _cabdda = map[string]*templateTag{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _aecbab: _ebege}, "\u0074\u0065\u0078\u0074\u002d\u0063\u0068\u0075\u006e\u006b": {_cdcdf: map[string]struct{}{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": {}}, _aecbab: _dbdgc}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _aecbab: _ffgbb}, "\u0074\u0061\u0062l\u0065": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _aecbab: _fgeeg}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {_cdcdf: map[string]struct{}{"\u0074\u0061\u0062l\u0065": {}}, _aecbab: _bbff}, "\u006c\u0069\u006e\u0065": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _aecbab: _bafebe}, "\u0072e\u0063\u0074\u0061\u006e\u0067\u006ce": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _aecbab: _acfdb}, "\u0065l\u006c\u0069\u0070\u0073\u0065": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _aecbab: _eedg}, "\u0069\u006d\u0061g\u0065": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _aecbab: _befgb}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _aecbab: _ccgag}, "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067": {_cdcdf: map[string]struct{}{"\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _aecbab: _caceg}, "\u0063\u0068\u0061r\u0074": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _aecbab: _faacc}, "\u0070\u0061\u0067\u0065\u002d\u0062\u0072\u0065\u0061\u006b": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _aecbab: _eagdc}, "\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064": {_cdcdf: map[string]struct{}{"\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}}, _aecbab: _ffffg}, "\u006c\u0069\u0073\u0074": {_cdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _aecbab: _gfgbg}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {_cdcdf: map[string]struct{}{"\u006c\u0069\u0073\u0074": {}}, _aecbab: _ffcbg}, "l\u0069\u0073\u0074\u002d\u006d\u0061\u0072\u006b\u0065\u0072": {_cdcdf: map[string]struct{}{"\u006c\u0069\u0073\u0074": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _aecbab: _adgc}}

// SetBuyerAddress sets the buyer address of the invoice.
func (_gaea *Invoice) SetBuyerAddress(address *InvoiceAddress) { _gaea._cfccb = address }

// SetColorLeft sets border color for left.
func (_eefe *border) SetColorLeft(col Color) { _eefe._efb = col }

// NewDivision returns a new Division container component.
func (_efbf *Creator) NewDivision() *Division { return _dfdab() }

// SetLineLevelOffset sets the amount of space an indentation level occupies
// for all new lines of the table of contents.
func (_feaag *TOC) SetLineLevelOffset(levelOffset float64) { _feaag._cbbeg = levelOffset }

// SetMargins sets the margins of the component. The margins are applied
// around the division.
func (_accg *Division) SetMargins(left, right, top, bottom float64) {
	_accg._acaa.Left = left
	_accg._acaa.Right = right
	_accg._acaa.Top = top
	_accg._acaa.Bottom = bottom
}

// GetMargins returns the margins of the line: left, right, top, bottom.
func (_cffg *Line) GetMargins() (float64, float64, float64, float64) {
	return _cffg._gabb.Left, _cffg._gabb.Right, _cffg._gabb.Top, _cffg._gabb.Bottom
}

// Padding returns the padding of the component.
func (_eecf *Division) Padding() (_fdeg, _dcga, _aaec, _eacb float64) {
	return _eecf._cgbb.Left, _eecf._cgbb.Right, _eecf._cgbb.Top, _eecf._cgbb.Bottom
}

// AddPatternResource adds pattern dictionary inside the resources dictionary.
func (_bgca *LinearShading) AddPatternResource(block *Block) (_gbgc _cc.PdfObjectName, _ddagf error) {
	_bfcbd := 1
	_eefbf := _cc.PdfObjectName("\u0050" + _aa.Itoa(_bfcbd))
	for block._ccf.HasPatternByName(_eefbf) {
		_bfcbd++
		_eefbf = _cc.PdfObjectName("\u0050" + _aa.Itoa(_bfcbd))
	}
	if _fdec := block._ccf.SetPatternByName(_eefbf, _bgca.ToPdfShadingPattern().ToPdfObject()); _fdec != nil {
		return "", _fdec
	}
	return _eefbf, nil
}

func (_agde *Invoice) newCell(_ggfe string, _cfbfb InvoiceCellProps) *InvoiceCell {
	return &InvoiceCell{_cfbfb, _ggfe}
}

// Curve represents a cubic Bezier curve with a control point.
type Curve struct {
	_caec float64
	_fcdb float64
	_fcg  float64
	_cfdd float64
	_fcfe float64
	_cbf  float64
	_gag  Color
	_cdaf float64
}

// SetColor sets the line color.
func (_gege *Curve) SetColor(col Color) { _gege._gag = col }

// SetMargins sets the Paragraph's margins.
func (_dacg *Paragraph) SetMargins(left, right, top, bottom float64) {
	_dacg._bcggf.Left = left
	_dacg._bcggf.Right = right
	_dacg._bcggf.Top = top
	_dacg._bcggf.Bottom = bottom
}

// ScaleToHeight scales the Block to a specified height, maintaining the same aspect ratio.
func (_dgg *Block) ScaleToHeight(h float64) { _eaef := h / _dgg._cce; _dgg.Scale(_eaef, _eaef) }

// ColorPoint is a pair of Color and a relative point where the color
// would be rendered.
type ColorPoint struct {
	_egea Color
	_ecbg float64
}

// NewGraphicSVGFromString creates a graphic SVG from a SVG string.
func NewGraphicSVGFromString(svgStr string) (*GraphicSVG, error) { return _agba(svgStr) }

// GetIndent get the cell's left indent.
func (_dgceb *TableCell) GetIndent() float64 { return _dgceb._fccd }

func _gfaf(_edede *_ga.PdfRectangle, _dcbc _gb.Matrix) *_ga.PdfRectangle {
	var _dddb _ga.PdfRectangle
	_dddb.Llx, _dddb.Lly = _dcbc.Transform(_edede.Llx, _edede.Lly)
	_dddb.Urx, _dddb.Ury = _dcbc.Transform(_edede.Urx, _edede.Ury)
	_dddb.Normalize()
	return &_dddb
}

// SetStyle sets the style for all the line components: number, title,
// separator, page.
func (_gdcg *TOCLine) SetStyle(style TextStyle) {
	_gdcg.Number.Style = style
	_gdcg.Title.Style = style
	_gdcg.Separator.Style = style
	_gdcg.Page.Style = style
}

// SetLineColor sets the line color.
func (_ecefc *Polyline) SetLineColor(color Color) { _ecefc._agbgf.LineColor = _fce(color) }

// CurRow returns the currently active cell's row number.
func (_ddfeg *Table) CurRow() int { _edaf := (_ddfeg._ebad-1)/_ddfeg._edfe + 1; return _edaf }

// InfoLines returns all the rows in the invoice information table as
// description-value cell pairs.
func (_badc *Invoice) InfoLines() [][2]*InvoiceCell {
	_efbeb := [][2]*InvoiceCell{_badc._badb, _badc._afdb, _badc._aagc}
	return append(_efbeb, _badc._gfda...)
}

func _fce(_eabg Color) _ga.PdfColor {
	if _eabg == nil {
		_eabg = ColorBlack
	}
	switch _cdea := _eabg.(type) {
	case cmykColor:
		return _ga.NewPdfColorDeviceCMYK(_cdea._bdge, _cdea._eegd, _cdea._ebe, _cdea._agg)
	case *LinearShading:
		return _ga.NewPdfColorPatternType2()
	case *RadialShading:
		return _ga.NewPdfColorPatternType3()
	}
	return _ga.NewPdfColorDeviceRGB(_eabg.ToRGB())
}

// Width returns the current page width.
func (_ecec *Creator) Width() float64 { return _ecec._bagf }

// NewImageFromFile creates an Image from a file.
func (_deba *Creator) NewImageFromFile(path string) (*Image, error) { return _caaf(path) }

func (_eagf *Ellipse) applyFitMode(_efg float64) {
	_efg -= _eagf._abgc.Left + _eagf._abgc.Right
	switch _eagf._ceeab {
	case FitModeFillWidth:
		_eagf.ScaleToWidth(_efg)
	}
}

// GeneratePageBlocks generates the page blocks for the Division component.
// Multiple blocks are generated if the contents wrap over multiple pages.
func (_fbbf *Division) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_aaca   []*Block
		_abdb   bool
		_egef   error
		_facdff = _fbbf._acfc.IsRelative()
		_acb    = _fbbf._acaa.Top
	)
	if _facdff && !_fbbf._dccg && !_fbbf._cdgb {
		_eaaf := _fbbf.ctxHeight(ctx.Width)
		if _eaaf > ctx.Height-_fbbf._acaa.Top && _eaaf <= ctx.PageHeight-ctx.Margins.Top-ctx.Margins.Bottom {
			if _aaca, ctx, _egef = _dbae().GeneratePageBlocks(ctx); _egef != nil {
				return nil, ctx, _egef
			}
			_abdb = true
			_acb = 0
		}
	}
	_gabg := ctx
	_aaae := ctx
	if _facdff {
		ctx.X += _fbbf._acaa.Left
		ctx.Y += _acb
		ctx.Width -= _fbbf._acaa.Left + _fbbf._acaa.Right
		ctx.Height -= _acb
		_aaae = ctx
		ctx.X += _fbbf._cgbb.Left
		ctx.Y += _fbbf._cgbb.Top
		ctx.Width -= _fbbf._cgbb.Left + _fbbf._cgbb.Right
		ctx.Height -= _fbbf._cgbb.Top
		ctx.Margins.Top += _fbbf._cgbb.Top
		ctx.Margins.Bottom += _fbbf._cgbb.Bottom
		ctx.Margins.Left += _fbbf._acaa.Left + _fbbf._cgbb.Left
		ctx.Margins.Right += _fbbf._acaa.Right + _fbbf._cgbb.Right
	}
	ctx.Inline = _fbbf._cdgb
	_egg := ctx
	_bbec := ctx
	var _aegf float64
	for _, _acaf := range _fbbf._dccd {
		if ctx.Inline {
			if (ctx.X-_egg.X)+_acaf.Width() <= ctx.Width {
				ctx.Y = _bbec.Y
				ctx.Height = _bbec.Height
			} else {
				ctx.X = _egg.X
				ctx.Width = _egg.Width
				_bbec.Y += _aegf
				_bbec.Height -= _aegf
				_aegf = 0
			}
		}
		_dgce, _ebcb, _dgcd := _acaf.GeneratePageBlocks(ctx)
		if _dgcd != nil {
			_bcd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006eg\u0020p\u0061\u0067\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073\u003a\u0020\u0025\u0076", _dgcd)
			return nil, ctx, _dgcd
		}
		if len(_dgce) < 1 {
			continue
		}
		if len(_aaca) > 0 {
			_aaca[len(_aaca)-1].mergeBlocks(_dgce[0])
			_aaca = append(_aaca, _dgce[1:]...)
		} else {
			if _bbga := _dgce[0]._fa; _bbga == nil || len(*_bbga) == 0 {
				_abdb = true
			}
			_aaca = append(_aaca, _dgce[0:]...)
		}
		if ctx.Inline {
			if ctx.Page != _ebcb.Page {
				_egg.Y = ctx.Margins.Top
				_egg.Height = ctx.PageHeight - ctx.Margins.Top
				_bbec.Y = _egg.Y
				_bbec.Height = _egg.Height
				_aegf = _ebcb.Height - _egg.Height
			} else {
				if _edefc := ctx.Height - _ebcb.Height; _edefc > _aegf {
					_aegf = _edefc
				}
			}
		} else {
			_ebcb.X = ctx.X
		}
		ctx = _ebcb
	}
	ctx.Inline = _gabg.Inline
	ctx.Margins = _gabg.Margins
	if _facdff {
		ctx.X = _gabg.X
		ctx.Width = _gabg.Width
		ctx.Y += _fbbf._cgbb.Bottom
		ctx.Height -= _fbbf._cgbb.Bottom
	}
	if _fbbf._eeaf != nil {
		_aaca, _egef = _fbbf.drawBackground(_aaca, _aaae, ctx, _abdb)
		if _egef != nil {
			return nil, ctx, _egef
		}
	}
	if _fbbf._acfc.IsAbsolute() {
		return _aaca, _gabg, nil
	}
	ctx.Y += _fbbf._acaa.Bottom
	ctx.Height -= _fbbf._acaa.Bottom
	return _aaca, ctx, nil
}
func (_cgfed *TextStyle) horizontalScale() float64 { return _cgfed.HorizontalScaling / 100 }

// LinearShading holds data for rendering a linear shading gradient.
type LinearShading struct {
	_cfbfe *shading
	_ebfea *_ga.PdfRectangle
	_dgdad float64
}

func (_gebd *Creator) newPage() *_ga.PdfPage {
	_gfbc := _ga.NewPdfPage()
	_adgg := _gebd._aad[0]
	_aeg := _gebd._aad[1]
	_ggb := _ga.PdfRectangle{Llx: 0, Lly: 0, Urx: _adgg, Ury: _aeg}
	_gfbc.MediaBox = &_ggb
	_gebd._bagf = _adgg
	_gebd._aafc = _aeg
	_gebd.initContext()
	return _gfbc
}

func (_dcdd *Invoice) drawAddress(_afdg *InvoiceAddress) []*StyledParagraph {
	var _fcadc []*StyledParagraph
	if _afdg.Heading != "" {
		_acfg := _cgfa(_dcdd._ffgb)
		_acfg.SetMargins(0, 0, 0, 7)
		_acfg.Append(_afdg.Heading)
		_fcadc = append(_fcadc, _acfg)
	}
	_afgb := _cgfa(_dcdd._gaae)
	_afgb.SetLineHeight(1.2)
	_afaa := _afdg.Separator
	if _afaa == "" {
		_afaa = _dcdd._ccfe
	}
	_dgee := _afdg.City
	if _afdg.State != "" {
		if _dgee != "" {
			_dgee += _afaa
		}
		_dgee += _afdg.State
	}
	if _afdg.Zip != "" {
		if _dgee != "" {
			_dgee += _afaa
		}
		_dgee += _afdg.Zip
	}
	if _afdg.Name != "" {
		_afgb.Append(_afdg.Name + "\u000a")
	}
	if _afdg.Street != "" {
		_afgb.Append(_afdg.Street + "\u000a")
	}
	if _afdg.Street2 != "" {
		_afgb.Append(_afdg.Street2 + "\u000a")
	}
	if _dgee != "" {
		_afgb.Append(_dgee + "\u000a")
	}
	if _afdg.Country != "" {
		_afgb.Append(_afdg.Country + "\u000a")
	}
	_fcdcd := _cgfa(_dcdd._gaae)
	_fcdcd.SetLineHeight(1.2)
	_fcdcd.SetMargins(0, 0, 7, 0)
	if _afdg.Phone != "" {
		_fcdcd.Append(_afdg.fmtLine(_afdg.Phone, "\u0050h\u006f\u006e\u0065\u003a\u0020", _afdg.HidePhoneLabel))
	}
	if _afdg.Email != "" {
		_fcdcd.Append(_afdg.fmtLine(_afdg.Email, "\u0045m\u0061\u0069\u006c\u003a\u0020", _afdg.HideEmailLabel))
	}
	_fcadc = append(_fcadc, _afgb, _fcdcd)
	return _fcadc
}

// SetLineHeight sets the line height (1.0 default).
func (_gafb *StyledParagraph) SetLineHeight(lineheight float64) { _gafb._bebe = lineheight }

// LineWidth returns the width of the line.
func (_eebc *Line) LineWidth() float64 { return _eebc._bfe }

// SetAngle would set the angle at which the gradient is rendered.
//
// The default angle would be 0 where the gradient would be rendered from left to right side.
func (_eaagf *LinearShading) SetAngle(angle float64) { _eaagf._dgdad = angle }

func (_gfff *Chapter) headingNumber() string {
	var _ffa string
	if _gfff._ebgde {
		if _gfff._baef != 0 {
			_ffa = _aa.Itoa(_gfff._baef) + "\u002e"
		}
		if _gfff._eagg != nil {
			_dadc := _gfff._eagg.headingNumber()
			if _dadc != "" {
				_ffa = _dadc + _ffa
			}
		}
	}
	return _ffa
}

// EnablePageWrap controls whether the table is wrapped across pages.
// If disabled, the table is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the table is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
func (_acee *Table) EnablePageWrap(enable bool) { _acee._gbdb = enable }

// SetFillOpacity sets the fill opacity of the rectangle.
func (_fagc *Rectangle) SetFillOpacity(opacity float64) { _fagc._gfdd = opacity }

// AddLine appends a new line to the invoice line items table.
func (_dgfea *Invoice) AddLine(values ...string) []*InvoiceCell {
	_dcgc := len(_dgfea._cdfg)
	var _fed []*InvoiceCell
	for _gaeag, _dagg := range values {
		_degb := _dgfea.newCell(_dagg, _dgfea._bgaab)
		if _gaeag < _dcgc {
			_degb.Alignment = _dgfea._cdfg[_gaeag].Alignment
		}
		_fed = append(_fed, _degb)
	}
	_dgfea._gdcac = append(_dgfea._gdcac, _fed)
	return _fed
}

func _gdcb(_ccba TextStyle) *List {
	return &List{_feeb: TextChunk{Text: "\u2022\u0020", Style: _ccba}, _addd: 0, _ecee: true, _ffeef: PositionRelative, _edda: _ccba}
}

// MultiRowCell makes a new cell with the specified row span and inserts it
// into the table at the current position.
func (_fceb *Table) MultiRowCell(rowspan int) *TableCell { return _fceb.MultiCell(rowspan, 1) }

// Context returns the current drawing context.
func (_dcab *Creator) Context() DrawContext { return _dcab._ggab }

func (_cgfba *templateProcessor) parseRectangle(_efbbc *templateNode) (interface{}, error) {
	_geec := _cgfba.creator.NewRectangle(0, 0, 0, 0)
	for _, _gaad := range _efbbc._acdf.Attr {
		_affc := _gaad.Value
		switch _agaf := _gaad.Name.Local; _agaf {
		case "\u0078":
			_geec._fcgf = _cgfba.parseFloatAttr(_agaf, _affc)
		case "\u0079":
			_geec._cbaa = _cgfba.parseFloatAttr(_agaf, _affc)
		case "\u0077\u0069\u0064t\u0068":
			_geec.SetWidth(_cgfba.parseFloatAttr(_agaf, _affc))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_geec.SetHeight(_cgfba.parseFloatAttr(_agaf, _affc))
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_geec.SetFillColor(_cgfba.parseColorAttr(_agaf, _affc))
		case "\u0066\u0069\u006cl\u002d\u006f\u0070\u0061\u0063\u0069\u0074\u0079":
			_geec.SetFillOpacity(_cgfba.parseFloatAttr(_agaf, _affc))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_geec.SetBorderColor(_cgfba.parseColorAttr(_agaf, _affc))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006f\u0070a\u0063\u0069\u0074\u0079":
			_geec.SetBorderOpacity(_cgfba.parseFloatAttr(_agaf, _affc))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_geec.SetBorderWidth(_cgfba.parseFloatAttr(_agaf, _affc))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_eccg, _gffad, _aacg, _adcgg := _cgfba.parseBorderRadiusAttr(_agaf, _affc)
			_geec.SetBorderRadius(_eccg, _gffad, _adcgg, _aacg)
		case "\u0062\u006f\u0072\u0064er\u002d\u0074\u006f\u0070\u002d\u006c\u0065\u0066\u0074\u002d\u0072\u0061\u0064\u0069u\u0073":
			_geec._acgb = _cgfba.parseFloatAttr(_agaf, _affc)
		case "\u0062\u006f\u0072de\u0072\u002d\u0074\u006f\u0070\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_geec._cgea = _cgfba.parseFloatAttr(_agaf, _affc)
		case "\u0062o\u0072\u0064\u0065\u0072-\u0062\u006f\u0074\u0074\u006fm\u002dl\u0065f\u0074\u002d\u0072\u0061\u0064\u0069\u0075s":
			_geec._bfag = _cgfba.parseFloatAttr(_agaf, _affc)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0062\u006f\u0074\u0074o\u006d\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061d\u0069\u0075\u0073":
			_geec._ebbe = _cgfba.parseFloatAttr(_agaf, _affc)
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_geec.SetPositioning(_cgfba.parsePositioningAttr(_agaf, _affc))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_geec.SetFitMode(_cgfba.parseFitModeAttr(_agaf, _affc))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_acdfc := _cgfba.parseMarginAttr(_agaf, _affc)
			_geec.SetMargins(_acdfc.Left, _acdfc.Right, _acdfc.Top, _acdfc.Bottom)
		default:
			_cgfba.nodeLogDebug(_efbbc, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020re\u0063\u0074\u0061\u006e\u0067\u006ce\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _agaf)
		}
	}
	return _geec, nil
}

func _gcffd(_fcge string) ([]string, error) {
	var (
		_febgd []string
		_afagd []rune
	)
	for _, _egdaf := range _fcge {
		if _egdaf == '\u000A' {
			if len(_afagd) > 0 {
				_febgd = append(_febgd, string(_afagd))
			}
			_febgd = append(_febgd, string(_egdaf))
			_afagd = nil
			continue
		}
		_afagd = append(_afagd, _egdaf)
	}
	if len(_afagd) > 0 {
		_febgd = append(_febgd, string(_afagd))
	}
	var _becab []string
	for _, _aedgc := range _febgd {
		_cdfcd := []rune(_aedgc)
		_abefa := _dc.NewScanner(_cdfcd)
		var _abfec []rune
		for _ccfg := 0; _ccfg < len(_cdfcd); _ccfg++ {
			_, _gagg, _befb := _abefa.Next()
			if _befb != nil {
				return nil, _befb
			}
			if _gagg == _dc.BreakProhibited || _bc.IsSpace(_cdfcd[_ccfg]) {
				_abfec = append(_abfec, _cdfcd[_ccfg])
				if _bc.IsSpace(_cdfcd[_ccfg]) {
					_becab = append(_becab, string(_abfec))
					_abfec = []rune{}
				}
				continue
			} else {
				if len(_abfec) > 0 {
					_becab = append(_becab, string(_abfec))
				}
				_abfec = []rune{_cdfcd[_ccfg]}
			}
		}
		if len(_abfec) > 0 {
			_becab = append(_becab, string(_abfec))
		}
	}
	return _becab, nil
}

// Opacity returns the opacity of the line.
func (_gddb *Line) Opacity() float64 { return _gddb._gbcb }

// Level returns the indentation level of the TOC line.
func (_agefc *TOCLine) Level() uint { return _agefc._dffba }

// SetPositioning sets the positioning of the line (absolute or relative).
func (_gfdc *Line) SetPositioning(positioning Positioning) { _gfdc._fcdd = positioning }

func _ada(_cbdc, _dbc, _dgfe, _eefb float64) *border {
	_fad := &border{}
	_fad._cbd = _cbdc
	_fad._bcgdb = _dbc
	_fad._begd = _dgfe
	_fad._cfcd = _eefb
	_fad._beea = ColorBlack
	_fad._adc = ColorBlack
	_fad._efb = ColorBlack
	_fad._gdc = ColorBlack
	_fad._dca = 0
	_fad._gab = 0
	_fad._fab = 0
	_fad._dggc = 0
	_fad.LineStyle = _ff.LineStyleSolid
	return _fad
}

// SetMargins sets the margins of the ellipse.
// NOTE: ellipse margins are only applied if relative positioning is used.
func (_fbcc *Ellipse) SetMargins(left, right, top, bottom float64) {
	_fbcc._abgc.Left = left
	_fbcc._abgc.Right = right
	_fbcc._abgc.Top = top
	_fbcc._abgc.Bottom = bottom
}

func _dbdgc(_efgded *templateProcessor, _ceceg *templateNode) (interface{}, error) {
	return _efgded.parseTextChunk(_ceceg, nil)
}

// UnsupportedRuneError is an error that occurs when there is unsupported glyph being used.
type UnsupportedRuneError struct {
	Message string
	Rune    rune
}

// SetBorderWidth sets the border width of the ellipse.
func (_bgbd *Ellipse) SetBorderWidth(bw float64) { _bgbd._eeaa = bw }

// PageSize represents the page size as a 2 element array representing the width and height in PDF document units (points).
type PageSize [2]float64

// SetTextOverflow controls the behavior of paragraph text which
// does not fit in the available space.
func (_ggaa *StyledParagraph) SetTextOverflow(textOverflow TextOverflow) { _ggaa._egead = textOverflow }

// Invoice represents a configurable invoice template.
type Invoice struct {
	_efdb  string
	_bagfe *Image
	_cfccb *InvoiceAddress
	_dceeg *InvoiceAddress
	_ccfe  string
	_badb  [2]*InvoiceCell
	_afdb  [2]*InvoiceCell
	_aagc  [2]*InvoiceCell
	_gfda  [][2]*InvoiceCell
	_cdfg  []*InvoiceCell
	_gdcac [][]*InvoiceCell
	_gfag  [2]*InvoiceCell
	_eacf  [2]*InvoiceCell
	_gbece [][2]*InvoiceCell
	_aeee  [2]string
	_ggbc  [2]string
	_eefc  [][2]string
	_dcefb TextStyle
	_bfcc  TextStyle
	_adgf  TextStyle
	_gaae  TextStyle
	_ffgb  TextStyle
	_fdbgg TextStyle
	_dede  TextStyle
	_eeab  InvoiceCellProps
	_bce   InvoiceCellProps
	_bgaab InvoiceCellProps
	_cbefc InvoiceCellProps
	_aaede Positioning
}

// SetWidth set the Image's document width to specified w. This does not change the raw image data, i.e.
// no actual scaling of data is performed. That is handled by the PDF viewer.
func (_gfeea *Image) SetWidth(w float64) { _gfeea._ggda = w }

func (_cbefb *templateProcessor) parseTextRenderingModeAttr(_efead, _aebb string) TextRenderingMode {
	_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u006d\u006f\u0064e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _efead, _aebb)
	_eace := map[string]TextRenderingMode{"\u0066\u0069\u006c\u006c": TextRenderingModeFill, "\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeStroke, "f\u0069\u006c\u006c\u002d\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeFillStroke, "\u0069n\u0076\u0069\u0073\u0069\u0062\u006ce": TextRenderingModeInvisible, "\u0066i\u006c\u006c\u002d\u0063\u006c\u0069p": TextRenderingModeFillClip, "s\u0074\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeStrokeClip, "\u0066\u0069l\u006c\u002d\u0073t\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeFillStrokeClip, "\u0063\u006c\u0069\u0070": TextRenderingModeClip}[_aebb]
	return _eace
}

// FitMode defines resizing options of an object inside a container.
type FitMode int

var (
	ColorBlack  = ColorRGBFromArithmetic(0, 0, 0)
	ColorWhite  = ColorRGBFromArithmetic(1, 1, 1)
	ColorRed    = ColorRGBFromArithmetic(1, 0, 0)
	ColorGreen  = ColorRGBFromArithmetic(0, 1, 0)
	ColorBlue   = ColorRGBFromArithmetic(0, 0, 1)
	ColorYellow = ColorRGBFromArithmetic(1, 1, 0)
)

// SetStyleBottom sets border style for bottom side.
func (_bca *border) SetStyleBottom(style CellBorderStyle) { _bca._aacc = style }

// MoveY moves the drawing context to absolute position y.
func (_defd *Creator) MoveY(y float64) { _defd._ggab.Y = y }

// SetVerticalAlignment set the cell's vertical alignment of content.
// Can be one of:
// - CellHorizontalAlignmentTop
// - CellHorizontalAlignmentMiddle
// - CellHorizontalAlignmentBottom
func (_bdcda *TableCell) SetVerticalAlignment(valign CellVerticalAlignment) { _bdcda._ddfbf = valign }

// SetWidth sets the width of the rectangle.
func (_edbe *Rectangle) SetWidth(width float64) { _edbe._bcffg = width }

// Flip flips the active page on the specified axes.
// If `flipH` is true, the page is flipped horizontally. Similarly, if `flipV`
// is true, the page is flipped vertically. If both are true, the page is
// flipped both horizontally and vertically.
// NOTE: the flip transformations are applied when the creator is finalized,
// which is at write time in most cases.
func (_aga *Creator) Flip(flipH, flipV bool) error {
	_eebda := _aga.getActivePage()
	if _eebda == nil {
		return _c.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_bedc, _acdc := _aga._bbf[_eebda]
	if !_acdc {
		_bedc = &pageTransformations{}
		_aga._bbf[_eebda] = _bedc
	}
	_bedc._cca = flipH
	_bedc._eea = flipV
	return nil
}

// SetColor sets the line color. Use ColorRGBFromHex, ColorRGBFrom8bit or
// ColorRGBFromArithmetic to create the color object.
func (_agcfa *Line) SetColor(color Color) { _agcfa._efdf = color }

// SetBorderWidth sets the border width of the rectangle.
func (_ebfe *Rectangle) SetBorderWidth(bw float64) { _ebfe._bffc = bw }

func (_ccd *pageTransformations) applyFlip(_aeb *_ga.PdfPage) error {
	_cfbe, _ccc := _ccd._cca, _ccd._eea
	if !_cfbe && !_ccc {
		return nil
	}
	if _aeb == nil {
		return _c.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_dcea, _ceac := _aeb.GetMediaBox()
	if _ceac != nil {
		return _ceac
	}
	_dfa, _afd := _dcea.Width(), _dcea.Height()
	_bfg, _ceac := _aeb.GetRotate()
	if _ceac != nil {
		_bcd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _ceac.Error())
	}
	if _bbgbe := _bfg%360 != 0 && _bfg%90 == 0; _bbgbe {
		if _ffaa := (360 + _bfg%360) % 360; _ffaa == 90 || _ffaa == 270 {
			_cfbe, _ccc = _ccc, _cfbe
		}
	}
	_ebcg, _daae := 1.0, 0.0
	if _cfbe {
		_ebcg, _daae = -1.0, -_dfa
	}
	_affe, _bdbdd := 1.0, 0.0
	if _ccc {
		_affe, _bdbdd = -1.0, -_afd
	}
	_aaeb := _da.NewContentCreator().Scale(_ebcg, _affe).Translate(_daae, _bdbdd)
	_cecf, _ceac := _cc.MakeStream(_aaeb.Bytes(), _cc.NewFlateEncoder())
	if _ceac != nil {
		return _ceac
	}
	_aafb := _cc.MakeArray(_cecf)
	_aafb.Append(_aeb.GetContentStreamObjs()...)
	_aeb.Contents = _aafb
	return nil
}

func (_ebaga *templateProcessor) parseLine(_ebbb *templateNode) (interface{}, error) {
	_gdbb := _ebaga.creator.NewLine(0, 0, 0, 0)
	for _, _gcgdc := range _ebbb._acdf.Attr {
		_cbgdd := _gcgdc.Value
		switch _geaf := _gcgdc.Name.Local; _geaf {
		case "\u0078\u0031":
			_gdbb._ccg = _ebaga.parseFloatAttr(_geaf, _cbgdd)
		case "\u0079\u0031":
			_gdbb._eggf = _ebaga.parseFloatAttr(_geaf, _cbgdd)
		case "\u0078\u0032":
			_gdbb._dfcae = _ebaga.parseFloatAttr(_geaf, _cbgdd)
		case "\u0079\u0032":
			_gdbb._feafb = _ebaga.parseFloatAttr(_geaf, _cbgdd)
		case "\u0074h\u0069\u0063\u006b\u006e\u0065\u0073s":
			_gdbb.SetLineWidth(_ebaga.parseFloatAttr(_geaf, _cbgdd))
		case "\u0063\u006f\u006co\u0072":
			_gdbb.SetColor(_ebaga.parseColorAttr(_geaf, _cbgdd))
		case "\u0073\u0074\u0079l\u0065":
			_gdbb.SetStyle(_ebaga.parseLineStyleAttr(_geaf, _cbgdd))
		case "\u0064\u0061\u0073\u0068\u002d\u0061\u0072\u0072\u0061\u0079":
			_gdbb.SetDashPattern(_ebaga.parseInt64Array(_geaf, _cbgdd), _gdbb._ffgdf)
		case "\u0064\u0061\u0073\u0068\u002d\u0070\u0068\u0061\u0073\u0065":
			_gdbb.SetDashPattern(_gdbb._fedd, _ebaga.parseInt64Attr(_geaf, _cbgdd))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_gdbb.SetOpacity(_ebaga.parseFloatAttr(_geaf, _cbgdd))
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_gdbb.SetPositioning(_ebaga.parsePositioningAttr(_geaf, _cbgdd))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_gdbb.SetFitMode(_ebaga.parseFitModeAttr(_geaf, _cbgdd))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_cbed := _ebaga.parseMarginAttr(_geaf, _cbgdd)
			_gdbb.SetMargins(_cbed.Left, _cbed.Right, _cbed.Top, _cbed.Bottom)
		default:
			_ebaga.nodeLogDebug(_ebbb, "\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u006c\u0069\u006e\u0065 \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _geaf)
		}
	}
	return _gdbb, nil
}
func _gebe(_bacbd Color, _aage float64) *ColorPoint { return &ColorPoint{_egea: _bacbd, _ecbg: _aage} }

// Margins returns the margins of the component.
func (_ggdb *Division) Margins() (_bcgfa, _dgedb, _cdfc, _daag float64) {
	return _ggdb._acaa.Left, _ggdb._acaa.Right, _ggdb._acaa.Top, _ggdb._acaa.Bottom
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_fgee *List) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var _ebcc float64
	var _dbeb []*StyledParagraph
	for _, _debfg := range _fgee._bdccg {
		_fefe := _cgfa(_fgee._edda)
		_fefe.SetEnableWrap(false)
		_fefe.SetTextAlignment(TextAlignmentRight)
		_fefe.Append(_debfg._feced.Text).Style = _debfg._feced.Style
		_ggaf := _fefe.getTextWidth() / 1000.0 / ctx.Width
		if _ebcc < _ggaf {
			_ebcc = _ggaf
		}
		_dbeb = append(_dbeb, _fefe)
	}
	_ccdf := _cbgg(2)
	_ccdf.SetColumnWidths(_ebcc, 1-_ebcc)
	_ccdf.SetMargins(_fgee._fcag.Left+_fgee._addd, _fgee._fcag.Right, _fgee._fcag.Top, _fgee._fcag.Bottom)
	_ccdf.EnableRowWrap(true)
	for _eefed, _eefa := range _fgee._bdccg {
		_ffga := _ccdf.NewCell()
		_ffga.SetIndent(0)
		_ffga.SetContent(_dbeb[_eefed])
		_ffga = _ccdf.NewCell()
		_ffga.SetIndent(0)
		_ffga.SetContent(_eefa._cecd)
	}
	return _ccdf.GeneratePageBlocks(ctx)
}

func _aaagb(_agag ...interface{}) (map[string]interface{}, error) {
	_abdg := len(_agag)
	if _abdg%2 != 0 {
		_bcd.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074i\u006e\u0067\u0020\u006d\u0061\u0070\u003a\u0020\u0025\u0064\u002e", _abdg)
		return nil, _cc.ErrRangeError
	}
	_dgggfb := map[string]interface{}{}
	for _ecabb := 0; _ecabb < _abdg; _ecabb += 2 {
		_fafb, _ccdd := _agag[_ecabb].(string)
		if !_ccdd {
			_bcd.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006d\u0061\u0070 \u006b\u0065\u0079\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u002e\u0020\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u002e", _agag[_ecabb])
			return nil, _cc.ErrTypeError
		}
		_dgggfb[_fafb] = _agag[_ecabb+1]
	}
	return _dgggfb, nil
}

// WriteToFile writes the Creator output to file specified by path.
func (_ggd *Creator) WriteToFile(outputPath string) error {
	abspath, _cagd := _pf.Abs(outputPath)
	if _cagd != nil {
		return _cagd
	}
	if _cagd = _e.MkdirAll(_p.Dir(abspath), 0o775); _cagd != nil {
		return _cagd
	}
	_dgca, _cagd := _e.OpenFile(abspath, _e.O_RDWR|_e.O_CREATE|_e.O_TRUNC, 0o775)
	if _cagd != nil {
		return _cagd
	}
	defer _dgca.Close()
	return _ggd.Write(_dgca)
}

// SetAddressStyle sets the style properties used to render the content of
// the invoice address sections.
func (_cedb *Invoice) SetAddressStyle(style TextStyle) { _cedb._gaae = style }

// Inline returns whether the inline mode of the division is active.
func (_eac *Division) Inline() bool { return _eac._cdgb }

// SetPos sets absolute positioning with specified coordinates.
func (_dccge *Paragraph) SetPos(x, y float64) {
	_dccge._gece = PositionAbsolute
	_dccge._acab = x
	_dccge._bge = y
}

// GeneratePageBlocks draws the rectangle on a new block representing the page. Implements the Drawable interface.
func (_dgeed *Rectangle) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_bedf  []*Block
		_bbefb = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_fgac  = ctx
		_aaad  = _dgeed._bffc / 2
	)
	_bccd := _dgeed._dgde.IsRelative()
	if _bccd {
		_dgeed.applyFitMode(ctx.Width)
		ctx.X += _dgeed._eaaa.Left + _aaad
		ctx.Y += _dgeed._eaaa.Top + _aaad
		ctx.Width -= _dgeed._eaaa.Left + _dgeed._eaaa.Right
		ctx.Height -= _dgeed._eaaa.Top + _dgeed._eaaa.Bottom
		if _dgeed._ebegg > ctx.Height {
			_bedf = append(_bedf, _bbefb)
			_bbefb = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_cegda := ctx
			_cegda.Y = ctx.Margins.Top + _dgeed._eaaa.Top + _aaad
			_cegda.X = ctx.Margins.Left + _dgeed._eaaa.Left + _aaad
			_cegda.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _dgeed._eaaa.Top - _dgeed._eaaa.Bottom
			_cegda.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _dgeed._eaaa.Left - _dgeed._eaaa.Right
			ctx = _cegda
		}
	} else {
		ctx.X = _dgeed._fcgf
		ctx.Y = _dgeed._cbaa
	}
	_facbf := _ff.Rectangle{X: ctx.X, Y: ctx.PageHeight - ctx.Y - _dgeed._ebegg, Width: _dgeed._bcffg, Height: _dgeed._ebegg, BorderRadiusTopLeft: _dgeed._acgb, BorderRadiusTopRight: _dgeed._cgea, BorderRadiusBottomLeft: _dgeed._bfag, BorderRadiusBottomRight: _dgeed._ebbe, Opacity: 1.0}
	if _dgeed._gbfa != nil {
		_facbf.FillEnabled = true
		_ecegb := _fce(_dgeed._gbfa)
		_afbb := _acdb(_bbefb, _ecegb, _dgeed._gbfa, func() Rectangle {
			return Rectangle{_fcgf: _facbf.X, _cbaa: _facbf.Y, _bcffg: _facbf.Width, _ebegg: _facbf.Height}
		})
		if _afbb != nil {
			return nil, ctx, _afbb
		}
		_facbf.FillColor = _ecegb
	}
	if _dgeed._gggb != nil && _dgeed._bffc > 0 {
		_facbf.BorderEnabled = true
		_facbf.BorderColor = _fce(_dgeed._gggb)
		_facbf.BorderWidth = _dgeed._bffc
	}
	_bbgf, _cdbd := _bbefb.setOpacity(_dgeed._gfdd, _dgeed._gffa)
	if _cdbd != nil {
		return nil, ctx, _cdbd
	}
	_faee, _, _cdbd := _facbf.Draw(_bbgf)
	if _cdbd != nil {
		return nil, ctx, _cdbd
	}
	if _cdbd = _bbefb.addContentsByString(string(_faee)); _cdbd != nil {
		return nil, ctx, _cdbd
	}
	if _bccd {
		ctx.X = _fgac.X
		ctx.Width = _fgac.Width
		_gfgb := _dgeed._ebegg + _aaad
		ctx.Y += _gfgb + _dgeed._eaaa.Bottom
		ctx.Height -= _gfgb
	} else {
		ctx = _fgac
	}
	_bedf = append(_bedf, _bbefb)
	return _bedf, ctx, nil
}

// SetBorderLineStyle sets border style (currently dashed or plain).
func (_dbade *TableCell) SetBorderLineStyle(style _ff.LineStyle) { _dbade._acafb = style }

// Margins returns the margins of the list: left, right, top, bottom.
func (_gdbf *List) Margins() (float64, float64, float64, float64) {
	return _gdbf._fcag.Left, _gdbf._fcag.Right, _gdbf._fcag.Top, _gdbf._fcag.Bottom
}

// Width returns Image's document width.
func (_gdbg *Image) Width() float64 { return _gdbg._ggda }

// SetAngle sets the rotation angle of the text.
func (_faca *Paragraph) SetAngle(angle float64) { _faca._cgaaf = angle }

func (_gdaec *Creator) wrapPageIfNeeded(_ged *_ga.PdfPage) (*_ga.PdfPage, error) {
	_eaeg, _egda := _ged.GetAllContentStreams()
	if _egda != nil {
		return nil, _egda
	}
	_aegc := _da.NewContentStreamParser(_eaeg)
	_fadd, _egda := _aegc.Parse()
	if _egda != nil {
		return nil, _egda
	}
	if !_fadd.HasUnclosedQ() {
		return nil, nil
	}
	_fadd.WrapIfNeeded()
	_dag, _egda := _cc.MakeStream(_fadd.Bytes(), _cc.NewFlateEncoder())
	if _egda != nil {
		return nil, _egda
	}
	_ged.Contents = _cc.MakeArray(_dag)
	return _ged, nil
}

// ColorRGBFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//
//	green := ColorRGBFromArithmetic(0.0, 1.0, 0.0)
func ColorRGBFromArithmetic(r, g, b float64) Color {
	return rgbColor{_ddae: _cd.Max(_cd.Min(r, 1.0), 0.0), _cede: _cd.Max(_cd.Min(g, 1.0), 0.0), _fgef: _cd.Max(_cd.Min(b, 1.0), 0.0)}
}

func _befgb(_feddc *templateProcessor, _cabcc *templateNode) (interface{}, error) {
	return _feddc.parseImage(_cabcc)
}

// SetShowLinks sets visibility of links for the TOC lines.
func (_acgc *TOC) SetShowLinks(showLinks bool) { _acgc._bbee = showLinks }

// SetTerms sets the terms and conditions section of the invoice.
func (_dcfge *Invoice) SetTerms(title, content string) { _dcfge._ggbc = [2]string{title, content} }

// SetFillColor sets the fill color.
func (_ebbg *PolyBezierCurve) SetFillColor(color Color) {
	_ebbg._ggag = color
	_ebbg._ecgf.FillColor = _fce(color)
}

func (_dcb *Image) applyFitMode(_gaegg float64) {
	_gaegg -= _dcb._eaac.Left + _dcb._eaac.Right
	switch _dcb._dfdb {
	case FitModeFillWidth:
		_dcb.ScaleToWidth(_gaegg)
	}
}

// Scale scales the rectangle dimensions by the specified factors.
func (_dgdc *Rectangle) Scale(xFactor, yFactor float64) {
	_dgdc._bcffg = xFactor * _dgdc._bcffg
	_dgdc._ebegg = yFactor * _dgdc._ebegg
}

// SetFillColor sets the fill color of the ellipse.
func (_egfa *Ellipse) SetFillColor(col Color) { _egfa._fecae = col }

// SetHeight sets the height of the rectangle.
func (_fgda *Rectangle) SetHeight(height float64) { _fgda._ebegg = height }

func _ccgag(_defdgb *templateProcessor, _ebegef *templateNode) (interface{}, error) {
	return _defdgb.parseChapter(_ebegef)
}

// Drawable is a widget that can be used to draw with the Creator.
type Drawable interface {
	// GeneratePageBlocks draw onto blocks representing Page contents. As the content can wrap over many pages, multiple
	// templates are returned, one per Page.  The function also takes a draw context containing information
	// where to draw (if relative positioning) and the available height to draw on accounting for Margins etc.
	GeneratePageBlocks(_bege DrawContext) ([]*Block, DrawContext, error)
}

// SetAngle sets the rotation angle of the text.
func (_bdgee *StyledParagraph) SetAngle(angle float64) { _bdgee._ecdcf = angle }

// SetPdfWriterAccessFunc sets a PdfWriter access function/hook.
// Exposes the PdfWriter just prior to writing the PDF.  Can be used to encrypt the output PDF, etc.
//
// Example of encrypting with a user/owner password "password"
// Prior to calling c.WriteFile():
//
//	c.SetPdfWriterAccessFunc(func(w *model.PdfWriter) error {
//		userPass := []byte("password")
//		ownerPass := []byte("password")
//		err := w.Encrypt(userPass, ownerPass, nil)
//		return err
//	})
func (_fbaec *Creator) SetPdfWriterAccessFunc(pdfWriterAccessFunc func(_dggeb *_ga.PdfWriter) error) {
	_fbaec._cefc = pdfWriterAccessFunc
}

// Style returns the style of the line.
func (_eggg *Line) Style() _ff.LineStyle { return _eggg._ffgbg }

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_ebeggb *StyledParagraph) GetMargins() (float64, float64, float64, float64) {
	return _ebeggb._fadcg.Left, _ebeggb._fadcg.Right, _ebeggb._fadcg.Top, _ebeggb._fadcg.Bottom
}

// SetColor sets the color of the Paragraph text.
//
// Example:
//
//  1. p := NewParagraph("Red paragraph")
//     // Set to red color with a hex code:
//     p.SetColor(creator.ColorRGBFromHex("#ff0000"))
//
//  2. Make Paragraph green with 8-bit rgb values (0-255 each component)
//     p.SetColor(creator.ColorRGBFrom8bit(0, 255, 0)
//
//  3. Make Paragraph blue with arithmetic (0-1) rgb components.
//     p.SetColor(creator.ColorRGBFromArithmetic(0, 0, 1.0)
func (_dbadf *Paragraph) SetColor(col Color) { _dbadf._ggbe = col }

// Heading returns the heading component of the table of contents.
func (_bffg *TOC) Heading() *StyledParagraph { return _bffg._aeeb }

// SetBorderOpacity sets the border opacity of the rectangle.
func (_ecga *Rectangle) SetBorderOpacity(opacity float64) { _ecga._gffa = opacity }

// NewPolygon creates a new polygon.
func (_gdbe *Creator) NewPolygon(points [][]_ff.Point) *Polygon { return _bdda(points) }

func _aafcf(_afca _ba.Image) (*Image, error) {
	_fdgg, _cbda := _ga.ImageHandling.NewImageFromGoImage(_afca)
	if _cbda != nil {
		return nil, _cbda
	}
	return _cfcc(_fdgg)
}

func _eebb(_ecdb float64, _cgaec float64, _acfcg float64, _agfgg float64, _dfade []*ColorPoint) *RadialShading {
	return &RadialShading{_acdae: &shading{_cfec: ColorWhite, _cebg: false, _fedf: []bool{false, false}, _ccaf: _dfade}, _dfgdc: _ecdb, _acfgc: _cgaec, _fecb: _acfcg, _dadbe: _agfgg, _bafe: AnchorCenter}
}

func (_cbggg *TableCell) width(_ddbg []float64, _afeea float64) float64 {
	_dbdcc := float64(0.0)
	for _bcdf := 0; _bcdf < _cbggg._daec; _bcdf++ {
		_dbdcc += _ddbg[_cbggg._fcbe+_bcdf-1]
	}
	return _dbdcc * _afeea
}

// Ellipse defines an ellipse with a center at (xc,yc) and a specified width and height.  The ellipse can have a colored
// fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Ellipse struct {
	_gdbee float64
	_bgfb  float64
	_eded  float64
	_bdcce float64
	_fabd  Positioning
	_fecae Color
	_bbdc  float64
	_bacf  Color
	_eeaa  float64
	_eade  float64
	_abgc  Margins
	_ceeab FitMode
}

// NewCell returns a new invoice table cell.
func (_dfdgg *Invoice) NewCell(value string) *InvoiceCell {
	return _dfdgg.newCell(value, _dfdgg.NewCellProps())
}

func _eadc(_eebdf string) *_ga.PdfAnnotation {
	_fbdef := _ga.NewPdfAnnotationLink()
	_caefc := _ga.NewBorderStyle()
	_caefc.SetBorderWidth(0)
	_fbdef.BS = _caefc.ToPdfObject()
	_bcebg := _ga.NewPdfActionURI()
	_bcebg.URI = _cc.MakeString(_eebdf)
	_fbdef.SetAction(_bcebg.PdfAction)
	return _fbdef.PdfAnnotation
}

// IsRelative checks if the positioning is relative.
func (_eec Positioning) IsRelative() bool { return _eec == PositionRelative }

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_fbaa *Creator) SetOptimizer(optimizer _ga.Optimizer) { _fbaa._adca = optimizer }

// NewPage adds a new Page to the Creator and sets as the active Page.
func (_egf *Creator) NewPage() *_ga.PdfPage {
	_daf := _egf.newPage()
	_egf._agfg = append(_egf._agfg, _daf)
	_egf._ggab.Page++
	return _daf
}

// MoveRight moves the drawing context right by relative displacement dx (negative goes left).
func (_feae *Creator) MoveRight(dx float64) { _feae._ggab.X += dx }

// NewSubchapter creates a new child chapter with the specified title.
func (_cdeb *Chapter) NewSubchapter(title string) *Chapter {
	_bfbe := _gbdbf(_cdeb._bgfc._bfbd)
	_bfbe.FontSize = 14
	_cdeb._gcf++
	_befg := _ebda(_cdeb, _cdeb._aaccb, _cdeb._eebd, title, _cdeb._gcf, _bfbe)
	_cdeb.Add(_befg)
	return _befg
}

func (_beffd *templateProcessor) parseImage(_cceb *templateNode) (interface{}, error) {
	var _aafbe string
	for _, _abfe := range _cceb._acdf.Attr {
		_fgaab := _abfe.Value
		switch _ccge := _abfe.Name.Local; _ccge {
		case "\u0073\u0072\u0063":
			_aafbe = _fgaab
		}
	}
	_cdcgg, _bgge := _beffd.loadImageFromSrc(_aafbe)
	if _bgge != nil {
		return nil, _bgge
	}
	for _, _adfgg := range _cceb._acdf.Attr {
		_dgdg := _adfgg.Value
		switch _ceff := _adfgg.Name.Local; _ceff {
		case "\u0061\u006c\u0069g\u006e":
			_cdcgg.SetHorizontalAlignment(_beffd.parseHorizontalAlignmentAttr(_ceff, _dgdg))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_cdcgg.SetOpacity(_beffd.parseFloatAttr(_ceff, _dgdg))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_efaee := _beffd.parseMarginAttr(_ceff, _dgdg)
			_cdcgg.SetMargins(_efaee.Left, _efaee.Right, _efaee.Top, _efaee.Bottom)
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_cdcgg.SetFitMode(_beffd.parseFitModeAttr(_ceff, _dgdg))
		case "\u0078":
			_cdcgg.SetPos(_beffd.parseFloatAttr(_ceff, _dgdg), _cdcgg._egca)
		case "\u0079":
			_cdcgg.SetPos(_cdcgg._fdab, _beffd.parseFloatAttr(_ceff, _dgdg))
		case "\u0077\u0069\u0064t\u0068":
			_cdcgg.SetWidth(_beffd.parseFloatAttr(_ceff, _dgdg))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_cdcgg.SetHeight(_beffd.parseFloatAttr(_ceff, _dgdg))
		case "\u0061\u006e\u0067l\u0065":
			_cdcgg.SetAngle(_beffd.parseFloatAttr(_ceff, _dgdg))
		case "\u0073\u0072\u0063":
			break
		default:
			_beffd.nodeLogDebug(_cceb, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0069m\u0061\u0067\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _ceff)
		}
	}
	return _cdcgg, nil
}

func (_dbdb *templateProcessor) parseFloatArray(_cdff, _gbfd string) []float64 {
	_bcd.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0066\u006c\u006f\u0061\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _cdff, _gbfd)
	_abdc := _a.Fields(_gbfd)
	_gcaed := make([]float64, 0, len(_abdc))
	for _, _acgg := range _abdc {
		_ddcg, _ := _aa.ParseFloat(_acgg, 64)
		_gcaed = append(_gcaed, _ddcg)
	}
	return _gcaed
}

func _agba(_gdgc string) (*GraphicSVG, error) {
	_egeb, _baaec := _ca.ParseFromString(_gdgc)
	if _baaec != nil {
		return nil, _baaec
	}
	return _cbba(_egeb)
}

// SetTotal sets the total of the invoice.
func (_aebg *Invoice) SetTotal(value string) { _aebg._eacf[1].Value = value }

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_cea *Block) DrawTemplate(c *Creator, r _eb.Reader, data interface{}, options *TemplateOptions) error {
	return _gebee(c, r, data, options, _cea)
}

// NewRectangle creates a new rectangle with the left corner at (`x`, `y`),
// having the specified width and height.
// NOTE: In relative positioning mode, `x` and `y` are calculated using the
// current context. Furthermore, when the fit mode is set to fill the available
// space, the rectangle is scaled so that it occupies the entire context width
// while maintaining the original aspect ratio.
func (_fdbf *Creator) NewRectangle(x, y, width, height float64) *Rectangle {
	return _cfeg(x, y, width, height)
}

// SetBackground sets the background properties of the component.
func (_gaeg *Division) SetBackground(background *Background) { _gaeg._eeaf = background }

// Add appends a new item to the list.
// The supported components are: *Paragraph, *StyledParagraph, *Division, *Image, *Table, and *List.
// Returns the marker used for the newly added item. The returned marker
// object can be used to change the text and style of the marker for the
// current item.
func (_bcad *List) Add(item VectorDrawable) (*TextChunk, error) {
	_gaffc := &listItem{_cecd: item, _feced: _bcad._feeb}
	switch _cgae := item.(type) {
	case *Paragraph:
	case *StyledParagraph:
	case *List:
		if _cgae._ecee {
			_cgae._addd = 15
		}
	case *Division:
	case *Image:
	case *Table:
	default:
		return nil, _c.New("\u0074\u0068i\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u006c\u0069\u0073\u0074")
	}
	_bcad._bdccg = append(_bcad._bdccg, _gaffc)
	return &_gaffc._feced, nil
}

func (_ebae *templateProcessor) parseInt64Array(_agae, _dcccc string) []int64 {
	_bcd.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0069\u006e\u0074\u0036\u0034\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _agae, _dcccc)
	_acfdc := _a.Fields(_dcccc)
	_dfeg := make([]int64, 0, len(_acfdc))
	for _, _bfec := range _acfdc {
		_dfae, _ := _aa.ParseInt(_bfec, 10, 64)
		_dfeg = append(_dfeg, _dfae)
	}
	return _dfeg
}

// ScaleToHeight sets the graphic svg scaling factor with the given height.
func (_fddbc *GraphicSVG) ScaleToHeight(h float64) {
	_dcgfg := _fddbc._debae.Width / _fddbc._debae.Height
	_fddbc._debae.Height = h
	_fddbc._debae.Width = h * _dcgfg
	_fddbc._debae.SetScaling(_dcgfg, _dcgfg)
}

func (_abeb *templateProcessor) parseTextAlignmentAttr(_aabaf, _bcggff string) TextAlignment {
	_bcd.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020t\u0065\u0078\u0074\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _aabaf, _bcggff)
	_fgebe := map[string]TextAlignment{"\u006c\u0065\u0066\u0074": TextAlignmentLeft, "\u0072\u0069\u0067h\u0074": TextAlignmentRight, "\u0063\u0065\u006e\u0074\u0065\u0072": TextAlignmentCenter, "\u006au\u0073\u0074\u0069\u0066\u0079": TextAlignmentJustify}[_bcggff]
	return _fgebe
}

func (_gccc *Invoice) setCellBorder(_dceed *TableCell, _fdfd *InvoiceCell) {
	for _, _daaa := range _fdfd.BorderSides {
		_dceed.SetBorder(_daaa, CellBorderStyleSingle, _fdfd.BorderWidth)
	}
	_dceed.SetBorderColor(_fdfd.BorderColor)
}

// Width is not used. Not used as a Table element is designed to fill into
// available width depending on the context. Returns 0.
func (_ccdff *Table) Width() float64 { return 0 }

// ToRGB implements interface Color.
// Note: It's not directly used since shading color works differently than regular color.
func (_fbgcb *RadialShading) ToRGB() (float64, float64, float64) { return 0, 0, 0 }

// Height returns the Block's height.
func (_ed *Block) Height() float64 { return _ed._cce }

// TotalLines returns all the rows in the invoice totals table as
// description-value cell pairs.
func (_bcdb *Invoice) TotalLines() [][2]*InvoiceCell {
	_daagf := [][2]*InvoiceCell{_bcdb._gfag}
	_daagf = append(_daagf, _bcdb._gbece...)
	return append(_daagf, _bcdb._eacf)
}

// Draw processes the specified Drawable widget and generates blocks that can
// be rendered to the output document. The generated blocks can span over one
// or more pages. Additional pages are added if the contents go over the current
// page. Each generated block is assigned to the creator page it will be
// rendered to. In order to render the generated blocks to the creator pages,
// call Finalize, Write or WriteToFile.
func (_fec *Creator) Draw(d Drawable) error {
	if _fec.getActivePage() == nil {
		_fec.NewPage()
	}
	_ecca, _ccdg, _afae := d.GeneratePageBlocks(_fec._ggab)
	if _afae != nil {
		return _afae
	}
	if len(_ccdg._ddcc) > 0 {
		_fec.Errors = append(_fec.Errors, _ccdg._ddcc...)
	}
	for _aege, _bdbdc := range _ecca {
		if _aege > 0 {
			_fec.NewPage()
		}
		_dged := _fec.getActivePage()
		if _ffb, _feca := _fec._fgd[_dged]; _feca {
			if _dfga := _ffb.mergeBlocks(_bdbdc); _dfga != nil {
				return _dfga
			}
			if _agbe := _cdae(_bdbdc._ccf, _ffb._ccf); _agbe != nil {
				return _agbe
			}
		} else {
			_fec._fgd[_dged] = _bdbdc
		}
	}
	_fec._ggab.X = _ccdg.X
	_fec._ggab.Y = _ccdg.Y
	_fec._ggab.Height = _ccdg.PageHeight - _ccdg.Y - _ccdg.Margins.Bottom
	return nil
}

// SetAngle sets Image rotation angle in degrees.
func (_cbcc *Image) SetAngle(angle float64) { _cbcc._egacc = angle }

// Logo returns the logo of the invoice.
func (_fgaad *Invoice) Logo() *Image { return _fgaad._bagfe }

func _cffaa(_bgef string, _dfbd, _ggfb TextStyle) *TOC {
	_cfgg := _ggfb
	_cfgg.FontSize = 14
	_fbce := _cgfa(_cfgg)
	_fbce.SetEnableWrap(true)
	_fbce.SetTextAlignment(TextAlignmentLeft)
	_fbce.SetMargins(0, 0, 0, 5)
	_afbeb := _fbce.Append(_bgef)
	_afbeb.Style = _cfgg
	return &TOC{_aeeb: _fbce, _gcgf: []*TOCLine{}, _daaaa: _dfbd, _gdeab: _dfbd, _dbdd: _dfbd, _afccg: _dfbd, _fbag: "\u002e", _cbbeg: 10, _ffcde: Margins{0, 0, 2, 2}, _dfdc: PositionRelative, _aced: _dfbd, _bbee: true}
}

// Width is not used. The list component is designed to fill into the available
// width depending on the context. Returns 0.
func (_gfca *List) Width() float64 { return 0 }

func (_bebf *Block) mergeBlocks(_cdg *Block) error {
	_fg := _cgd(_bebf._fa, _bebf._ccf, _cdg._fa, _cdg._ccf)
	if _fg != nil {
		return _fg
	}
	for _, _ebde := range _cdg._fdd {
		_bebf.AddAnnotation(_ebde)
	}
	return nil
}

func (_aadcd *Invoice) generateTotalBlocks(_aabf DrawContext) ([]*Block, DrawContext, error) {
	_affg := _cbgg(4)
	_affg.SetMargins(0, 0, 10, 10)
	_gfdb := [][2]*InvoiceCell{_aadcd._gfag}
	_gfdb = append(_gfdb, _aadcd._gbece...)
	_gfdb = append(_gfdb, _aadcd._eacf)
	for _, _cfcde := range _gfdb {
		_cabbc, _dcfa := _cfcde[0], _cfcde[1]
		if _dcfa.Value == "" {
			continue
		}
		_affg.SkipCells(2)
		_egfcg := _affg.NewCell()
		_egfcg.SetBackgroundColor(_cabbc.BackgroundColor)
		_egfcg.SetHorizontalAlignment(_dcfa.Alignment)
		_aadcd.setCellBorder(_egfcg, _cabbc)
		_fcgb := _cgfa(_cabbc.TextStyle)
		_fcgb.SetMargins(0, 0, 2, 1)
		_fcgb.Append(_cabbc.Value)
		_egfcg.SetContent(_fcgb)
		_egfcg = _affg.NewCell()
		_egfcg.SetBackgroundColor(_dcfa.BackgroundColor)
		_egfcg.SetHorizontalAlignment(_dcfa.Alignment)
		_aadcd.setCellBorder(_egfcg, _cabbc)
		_fcgb = _cgfa(_dcfa.TextStyle)
		_fcgb.SetMargins(0, 0, 2, 1)
		_fcgb.Append(_dcfa.Value)
		_egfcg.SetContent(_fcgb)
	}
	return _affg.GeneratePageBlocks(_aabf)
}

// Height returns the height of the graphic svg.
func (_gfg *GraphicSVG) Height() float64 { return _gfg._debae.Height }

func (_gadgd *StyledParagraph) wrapWordChunks() {
	if !_gadgd._ceba {
		return
	}
	var (
		_bbcd   []*TextChunk
		_ededbg *_ga.PdfFont
	)
	for _, _aeca := range _gadgd._eddc {
		_fcfd := []rune(_aeca.Text)
		if _ededbg == nil {
			_ededbg = _aeca.Style.Font
		}
		_dfccd := _aeca._cgfaa
		_afbc := _aeca.VerticalAlignment
		if len(_bbcd) > 0 {
			if len(_fcfd) == 1 && _bc.IsPunct(_fcfd[0]) && _aeca.Style.Font == _ededbg {
				_cccfe := []rune(_bbcd[len(_bbcd)-1].Text)
				_bbcd[len(_bbcd)-1].Text = string(append(_cccfe, _fcfd[0]))
				continue
			} else {
				_, _dgbb := _aa.Atoi(_aeca.Text)
				if _dgbb == nil {
					_fbdga := []rune(_bbcd[len(_bbcd)-1].Text)
					_ffaad := len(_fbdga)
					if _ffaad >= 2 {
						_, _gbggf := _aa.Atoi(string(_fbdga[_ffaad-2]))
						if _gbggf == nil && _bc.IsPunct(_fbdga[_ffaad-1]) {
							_bbcd[len(_bbcd)-1].Text = string(append(_fbdga, _fcfd...))
							continue
						}
					}
				}
			}
		}
		_cedc, _edfd := _gcffd(_aeca.Text)
		if _edfd != nil {
			_bcd.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0062\u0072\u0065\u0061\u006b\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020w\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0076", _edfd)
			_cedc = []string{_aeca.Text}
		}
		for _, _ddaaf := range _cedc {
			_gdeb := NewTextChunk(_ddaaf, _aeca.Style)
			_gdeb._cgfaa = _dffg(_dfccd)
			_gdeb.VerticalAlignment = _afbc
			_bbcd = append(_bbcd, _gdeb)
		}
		_ededbg = _aeca.Style.Font
	}
	if len(_bbcd) > 0 {
		_gadgd._eddc = _bbcd
	}
}

// SetPos sets the Block's positioning to absolute mode with the specified coordinates.
func (_dcf *Block) SetPos(x, y float64) { _dcf._gfb = PositionAbsolute; _dcf._bcde = x; _dcf._gc = y }

// Positioning returns the type of positioning the rectangle is set to use.
func (_afdef *Rectangle) Positioning() Positioning { return _afdef._dgde }

func _cbea(_afeef string) bool {
	_bfbgd := func(_gabf rune) bool { return _gabf == '\u000A' }
	_dfec := _a.TrimFunc(_afeef, _bfbgd)
	_fffg := _gf.Paragraph{}
	_, _cdaae := _fffg.SetString(_dfec)
	if _cdaae != nil {
		return true
	}
	_deacc, _cdaae := _fffg.Order()
	if _cdaae != nil {
		return true
	}
	if _deacc.NumRuns() < 1 {
		return true
	}
	return _fffg.IsLeftToRight()
}

func _agcdg(_dbgad string, _adeb TextStyle) *Paragraph {
	_cgdd := &Paragraph{_bfgg: _dbgad, _bfbd: _adeb.Font, _dbad: _adeb.FontSize, _agdb: 1.0, _dbfb: true, _deaf: true, _caea: TextAlignmentLeft, _cgaaf: 0, _ddeg: 1, _fabcb: 1, _gece: PositionRelative}
	_cgdd.SetColor(_adeb.Color)
	return _cgdd
}

// GetCoords returns the (x1, y1), (x2, y2) points defining the Line.
func (_aacac *Line) GetCoords() (float64, float64, float64, float64) {
	return _aacac._ccg, _aacac._eggf, _aacac._dfcae, _aacac._feafb
}

func (_ecdcfe *templateProcessor) parseTextOverflowAttr(_fecbg, _gcdde string) TextOverflow {
	_bcd.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0074e\u0078\u0074\u0020o\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _fecbg, _gcdde)
	_bdeb := map[string]TextOverflow{"\u0076i\u0073\u0069\u0062\u006c\u0065": TextOverflowVisible, "\u0068\u0069\u0064\u0064\u0065\u006e": TextOverflowHidden}[_gcdde]
	return _bdeb
}

func (_decbd *StyledParagraph) getTextWidth() float64 {
	var _abada float64
	_eggd := len(_decbd._eddc)
	for _feeeg, _cbgd := range _decbd._eddc {
		_abgef := &_cbgd.Style
		_gega := len(_cbgd.Text)
		for _aadf, _gbd := range _cbgd.Text {
			if _gbd == '\u000A' {
				continue
			}
			_gbgdd, _fbcg := _abgef.Font.GetRuneMetrics(_gbd)
			if !_fbcg {
				_bcd.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _gbd)
				return -1
			}
			_abada += _abgef.FontSize * _gbgdd.Wx * _abgef.horizontalScale()
			if _gbd != ' ' && (_feeeg != _eggd-1 || _aadf != _gega-1) {
				_abada += _abgef.CharSpacing * 1000.0
			}
		}
	}
	return _abada
}

// SetFillColor sets background color for border.
func (_egb *border) SetFillColor(col Color) { _egb._fga = col }

func _cgfa(_ccgc TextStyle) *StyledParagraph {
	return &StyledParagraph{_eddc: []*TextChunk{}, _afaf: _ccgc, _dbgf: _dcgb(_ccgc.Font), _bebe: 1.0, _dcgfc: TextAlignmentLeft, _bfcgb: true, _afeg: true, _ceba: false, _ecdcf: 0, _adbb: 1, _cgee: 1, _bedd: PositionRelative}
}

type marginDrawable interface {
	VectorDrawable
	GetMargins() (float64, float64, float64, float64)
}

// MoveDown moves the drawing context down by relative displacement dy (negative goes up).
func (_cade *Creator) MoveDown(dy float64) { _cade._ggab.Y += dy }

// SetTitle sets the title of the invoice.
func (_ebgcf *Invoice) SetTitle(title string) { _ebgcf._efdb = title }

func (_baee *templateProcessor) parseListItem(_cdbg *templateNode) (interface{}, error) {
	if _cdbg._gbdf == nil {
		_baee.nodeLogError(_cdbg, "\u004c\u0069\u0073t\u0020\u0069\u0074\u0065m\u0020\u0070\u0061\u0072\u0065\u006e\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _bfegb
	}
	_cgac, _ecdf := _cdbg._gbdf._defae.(*List)
	if !_ecdf {
		_baee.nodeLogError(_cdbg, "\u004c\u0069s\u0074\u0020\u0069\u0074\u0065\u006d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u004cis\u0074\u002e")
		return nil, _bfegb
	}
	_ggfg := _egafb()
	_ggfg._feced = _cgac._feeb
	return _ggfg, nil
}

// SetPositioning sets the positioning of the ellipse (absolute or relative).
func (_gdgge *Ellipse) SetPositioning(position Positioning) { _gdgge._fabd = position }

// SetWidthBottom sets border width for bottom.
func (_bdbd *border) SetWidthBottom(bw float64) { _bdbd._gab = bw }

// NewCellProps returns the default properties of an invoice cell.
func (_bagc *Invoice) NewCellProps() InvoiceCellProps {
	_ggdf := ColorRGBFrom8bit(255, 255, 255)
	return InvoiceCellProps{TextStyle: _bagc._dcefb, Alignment: CellHorizontalAlignmentLeft, BackgroundColor: _ggdf, BorderColor: _ggdf, BorderWidth: 1, BorderSides: []CellBorderSide{CellBorderSideAll}}
}

// SetFitMode sets the fit mode of the ellipse.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_dbga *Ellipse) SetFitMode(fitMode FitMode) { _dbga._ceeab = fitMode }

// NewStyledTOCLine creates a new table of contents line with the provided style.
func (_fbdg *Creator) NewStyledTOCLine(number, title, page TextChunk, level uint, style TextStyle) *TOCLine {
	return _ecbaa(number, title, page, level, style)
}

// HeaderFunctionArgs holds the input arguments to a header drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type HeaderFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// EnableWordWrap sets the paragraph word wrap flag.
func (_cgga *StyledParagraph) EnableWordWrap(val bool) { _cgga._ceba = val }

// SetBorderWidth sets the border width.
func (_dcfbe *Polygon) SetBorderWidth(borderWidth float64) { _dcfbe._acaag.BorderWidth = borderWidth }

// SetFontSize sets the font size in document units (points).
func (_bffa *Paragraph) SetFontSize(fontSize float64) { _bffa._dbad = fontSize }

func (_bbfce *templateProcessor) parsePageBreak(_egfae *templateNode) (interface{}, error) {
	return _dbae(), nil
}

func (_ggbcf *templateProcessor) parseDivision(_deaa *templateNode) (interface{}, error) {
	_daeag := _ggbcf.creator.NewDivision()
	for _, _dbdff := range _deaa._acdf.Attr {
		_fbda := _dbdff.Value
		switch _bbadd := _dbdff.Name.Local; _bbadd {
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_daeag.EnablePageWrap(_ggbcf.parseBoolAttr(_bbadd, _fbda))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_dagbdc := _ggbcf.parseMarginAttr(_bbadd, _fbda)
			_daeag.SetMargins(_dagbdc.Left, _dagbdc.Right, _dagbdc.Top, _dagbdc.Bottom)
		case "\u0070a\u0064\u0064\u0069\u006e\u0067":
			_cgda := _ggbcf.parseMarginAttr(_bbadd, _fbda)
			_daeag.SetPadding(_cgda.Left, _cgda.Right, _cgda.Top, _cgda.Bottom)
		default:
			_ggbcf.nodeLogDebug(_deaa, "U\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0064\u0069\u0076\u0069\u0073\u0069on\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025s`\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _bbadd)
		}
	}
	return _daeag, nil
}

// Rectangle defines a rectangle with upper left corner at (x,y) and a specified width and height.  The rectangle
// can have a colored fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Rectangle struct {
	_fcgf  float64
	_cbaa  float64
	_bcffg float64
	_ebegg float64
	_dgde  Positioning
	_gbfa  Color
	_gfdd  float64
	_gggb  Color
	_bffc  float64
	_gffa  float64
	_acgb  float64
	_cgea  float64
	_bfag  float64
	_ebbe  float64
	_eaaa  Margins
	_efbb  FitMode
}

// SetExtends specifies whether ot extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_cgcb *LinearShading) SetExtends(start bool, end bool) { _cgcb._cfbfe.SetExtends(start, end) }

type shading struct {
	_cfec Color
	_cebg bool
	_fedf []bool
	_ccaf []*ColorPoint
}

func (_bbgag *Invoice) drawInformation() *Table {
	_edaa := _cbgg(2)
	_bcbc := append([][2]*InvoiceCell{_bbgag._badb, _bbgag._afdb, _bbgag._aagc}, _bbgag._gfda...)
	for _, _gggf := range _bcbc {
		_cgc, _adfg := _gggf[0], _gggf[1]
		if _adfg.Value == "" {
			continue
		}
		_cdcda := _edaa.NewCell()
		_cdcda.SetBackgroundColor(_cgc.BackgroundColor)
		_bbgag.setCellBorder(_cdcda, _cgc)
		_accea := _cgfa(_cgc.TextStyle)
		_accea.Append(_cgc.Value)
		_accea.SetMargins(0, 0, 2, 1)
		_cdcda.SetContent(_accea)
		_cdcda = _edaa.NewCell()
		_cdcda.SetBackgroundColor(_adfg.BackgroundColor)
		_bbgag.setCellBorder(_cdcda, _adfg)
		_accea = _cgfa(_adfg.TextStyle)
		_accea.Append(_adfg.Value)
		_accea.SetMargins(0, 0, 2, 1)
		_cdcda.SetContent(_accea)
	}
	return _edaa
}

// GeneratePageBlocks generates the table page blocks. Multiple blocks are
// generated if the contents wrap over multiple pages.
// Implements the Drawable interface.
func (_cfcec *Table) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ddbe := _cfcec
	if _cfcec._ggea {
		_ddbe = _cfcec.clone()
	}
	return _aceb(_ddbe, ctx)
}

// Height returns the height of the ellipse.
func (_edcbb *Ellipse) Height() float64 { return _edcbb._bdcce }

// SetHeight sets the height of the ellipse.
func (_cfda *Ellipse) SetHeight(height float64) { _cfda._bdcce = height }

// SetDueDate sets the due date of the invoice.
func (_ddfc *Invoice) SetDueDate(dueDate string) (*InvoiceCell, *InvoiceCell) {
	_ddfc._aagc[1].Value = dueDate
	return _ddfc._aagc[0], _ddfc._aagc[1]
}

// SetFillOpacity sets the fill opacity.
func (_gfcgf *CurvePolygon) SetFillOpacity(opacity float64) { _gfcgf._gfbf = opacity }

// GetHeading returns the chapter heading paragraph. Used to give access to address style: font, sizing etc.
func (_acd *Chapter) GetHeading() *Paragraph { return _acd._bgfc }

// CurCol returns the currently active cell's column number.
func (_dceeb *Table) CurCol() int { _fedag := (_dceeb._ebad-1)%(_dceeb._edfe) + 1; return _fedag }

// SetOutlineTree adds the specified outline tree to the PDF file generated
// by the creator. Adding an external outline tree disables the automatic
// generation of outlines done by the creator for the relevant components.
func (_cabg *Creator) SetOutlineTree(outlineTree *_ga.PdfOutlineTreeNode) { _cabg._afag = outlineTree }

// SetSideBorderStyle sets the cell's side border style.
func (_aecba *TableCell) SetSideBorderStyle(side CellBorderSide, style CellBorderStyle) {
	switch side {
	case CellBorderSideAll:
		_aecba._efdbd = style
		_aecba._ggce = style
		_aecba._ebce = style
		_aecba._bage = style
	case CellBorderSideTop:
		_aecba._efdbd = style
	case CellBorderSideBottom:
		_aecba._ggce = style
	case CellBorderSideLeft:
		_aecba._ebce = style
	case CellBorderSideRight:
		_aecba._bage = style
	}
}

func (_aece *Invoice) generateHeaderBlocks(_faag DrawContext) ([]*Block, DrawContext, error) {
	_cggc := _cgfa(_aece._adgf)
	_cggc.SetEnableWrap(true)
	_cggc.Append(_aece._efdb)
	_fgbc := _cbgg(2)
	if _aece._bagfe != nil {
		_abgg := _fgbc.NewCell()
		_abgg.SetHorizontalAlignment(CellHorizontalAlignmentLeft)
		_abgg.SetVerticalAlignment(CellVerticalAlignmentMiddle)
		_abgg.SetIndent(0)
		_abgg.SetContent(_aece._bagfe)
		_aece._bagfe.ScaleToHeight(_cggc.Height() + 20)
	} else {
		_fgbc.SkipCells(1)
	}
	_feafg := _fgbc.NewCell()
	_feafg.SetHorizontalAlignment(CellHorizontalAlignmentRight)
	_feafg.SetVerticalAlignment(CellVerticalAlignmentMiddle)
	_feafg.SetContent(_cggc)
	return _fgbc.GeneratePageBlocks(_faag)
}

// SetBorderOpacity sets the border opacity.
func (_cbfa *Polygon) SetBorderOpacity(opacity float64) { _cbfa._bgac = opacity }

// InvoiceCell represents any cell belonging to a table from the invoice
// template. The main tables are the invoice information table, the line
// items table and totals table. Contains the text value of the cell and
// the style properties of the cell.
type InvoiceCell struct {
	InvoiceCellProps
	Value string
}

func _edae(_defdc *Creator, _dbdf string, _edga []byte, _gdbec *TemplateOptions, _dege componentRenderer) *templateProcessor {
	if _gdbec == nil {
		_gdbec = &TemplateOptions{}
	}
	_gdbec.init()
	if _dege == nil {
		_dege = _defdc
	}
	return &templateProcessor{creator: _defdc, _cebag: _edga, _agcb: _gdbec, _ebdg: _dege, _bfgd: _dbdf}
}

// FitMode returns the fit mode of the ellipse.
func (_eed *Ellipse) FitMode() FitMode { return _eed._ceeab }

// GeneratePageBlocks draws the polygon on a new block representing the page.
// Implements the Drawable interface.
func (_cbeg *Polygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cdfga := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_egcg, _bfaa := _cdfga.setOpacity(_cbeg._afdbc, _cbeg._bgac)
	if _bfaa != nil {
		return nil, ctx, _bfaa
	}
	_fccf := _cbeg._acaag
	_fccf.FillEnabled = _fccf.FillColor != nil
	_fccf.BorderEnabled = _fccf.BorderColor != nil && _fccf.BorderWidth > 0
	_gbga := _fccf.Points
	_gcda := _ga.PdfRectangle{}
	_daac := false
	for _aaef := range _gbga {
		for _adcfg := range _gbga[_aaef] {
			_gbcbd := &_gbga[_aaef][_adcfg]
			_gbcbd.Y = ctx.PageHeight - _gbcbd.Y
			if !_daac {
				_gcda.Llx = _gbcbd.X
				_gcda.Lly = _gbcbd.Y
				_gcda.Urx = _gbcbd.X
				_gcda.Ury = _gbcbd.Y
				_daac = true
			} else {
				_gcda.Llx = _cd.Min(_gcda.Llx, _gbcbd.X)
				_gcda.Lly = _cd.Min(_gcda.Lly, _gbcbd.Y)
				_gcda.Urx = _cd.Max(_gcda.Urx, _gbcbd.X)
				_gcda.Ury = _cd.Max(_gcda.Ury, _gbcbd.Y)
			}
		}
	}
	if _fccf.FillEnabled {
		_dadg := _acdb(_cdfga, _cbeg._acaag.FillColor, _cbeg._ddad, func() Rectangle {
			return Rectangle{_fcgf: _gcda.Llx, _cbaa: _gcda.Lly, _bcffg: _gcda.Width(), _ebegg: _gcda.Height()}
		})
		if _dadg != nil {
			return nil, ctx, _dadg
		}
	}
	_defbe, _, _bfaa := _fccf.Draw(_egcg)
	if _bfaa != nil {
		return nil, ctx, _bfaa
	}
	if _bfaa = _cdfga.addContentsByString(string(_defbe)); _bfaa != nil {
		return nil, ctx, _bfaa
	}
	return []*Block{_cdfga}, ctx, nil
}

// DrawWithContext draws the Block using the specified drawing context.
func (_edf *Block) DrawWithContext(d Drawable, ctx DrawContext) error {
	_agf, _, _dea := d.GeneratePageBlocks(ctx)
	if _dea != nil {
		return _dea
	}
	if len(_agf) != 1 {
		return ErrContentNotFit
	}
	for _, _daa := range _agf {
		if _cba := _edf.mergeBlocks(_daa); _cba != nil {
			return _cba
		}
	}
	return nil
}
func _cddb(_bagadd *_dda.Decoder) (int, int) { return 0, 0 }
func _gfgbg(_fdbgaf *templateProcessor, _bffcf *templateNode) (interface{}, error) {
	return _fdbgaf.parseList(_bffcf)
}

// AddShadingResource adds shading dictionary inside the resources dictionary.
func (_gcbg *RadialShading) AddShadingResource(block *Block) (_gfbfb _cc.PdfObjectName, _gded error) {
	_dbff := 1
	_gfbfb = _cc.PdfObjectName("\u0053\u0068" + _aa.Itoa(_dbff))
	for block._ccf.HasShadingByName(_gfbfb) {
		_dbff++
		_gfbfb = _cc.PdfObjectName("\u0053\u0068" + _aa.Itoa(_dbff))
	}
	if _gfea := block._ccf.SetShadingByName(_gfbfb, _gcbg.shadingModel().ToPdfObject()); _gfea != nil {
		return "", _gfea
	}
	return _gfbfb, nil
}

const (
	TextOverflowVisible TextOverflow = iota
	TextOverflowHidden
)

// Background contains properties related to the background of a component.
type Background struct {
	FillColor               Color
	BorderColor             Color
	BorderSize              float64
	BorderRadiusTopLeft     float64
	BorderRadiusTopRight    float64
	BorderRadiusBottomLeft  float64
	BorderRadiusBottomRight float64
}

// SetOpacity sets the opacity of the line (0-1).
func (_fcba *Line) SetOpacity(opacity float64) { _fcba._gbcb = opacity }

func (_ddf *Block) duplicate() *Block {
	_ebg := &Block{}
	*_ebg = *_ddf
	_bee := _da.ContentStreamOperations{}
	_bee = append(_bee, *_ddf._fa...)
	_ebg._fa = &_bee
	return _ebg
}

func _bbb(_dffd, _dgac TextStyle) *Invoice {
	_ecgc := &Invoice{_efdb: "\u0049N\u0056\u004f\u0049\u0043\u0045", _ccfe: "\u002c\u0020", _dcefb: _dffd, _bfcc: _dgac}
	_ecgc._dceeg = &InvoiceAddress{Separator: _ecgc._ccfe}
	_ecgc._cfccb = &InvoiceAddress{Heading: "\u0042i\u006c\u006c\u0020\u0074\u006f", Separator: _ecgc._ccfe}
	_dgcc := ColorRGBFrom8bit(245, 245, 245)
	_fggc := ColorRGBFrom8bit(155, 155, 155)
	_ecgc._adgf = _dgac
	_ecgc._adgf.Color = _fggc
	_ecgc._adgf.FontSize = 20
	_ecgc._gaae = _dffd
	_ecgc._ffgb = _dgac
	_ecgc._fdbgg = _dffd
	_ecgc._dede = _dgac
	_ecgc._eeab = _ecgc.NewCellProps()
	_ecgc._eeab.BackgroundColor = _dgcc
	_ecgc._eeab.TextStyle = _dgac
	_ecgc._bce = _ecgc.NewCellProps()
	_ecgc._bce.TextStyle = _dgac
	_ecgc._bce.BackgroundColor = _dgcc
	_ecgc._bce.BorderColor = _dgcc
	_ecgc._bgaab = _ecgc.NewCellProps()
	_ecgc._bgaab.BorderColor = _dgcc
	_ecgc._bgaab.BorderSides = []CellBorderSide{CellBorderSideBottom}
	_ecgc._bgaab.Alignment = CellHorizontalAlignmentRight
	_ecgc._cbefc = _ecgc.NewCellProps()
	_ecgc._cbefc.Alignment = CellHorizontalAlignmentRight
	_ecgc._badb = [2]*InvoiceCell{_ecgc.newCell("\u0049\u006e\u0076\u006f\u0069\u0063\u0065\u0020\u006eu\u006d\u0062\u0065\u0072", _ecgc._eeab), _ecgc.newCell("", _ecgc._eeab)}
	_ecgc._afdb = [2]*InvoiceCell{_ecgc.newCell("\u0044\u0061\u0074\u0065", _ecgc._eeab), _ecgc.newCell("", _ecgc._eeab)}
	_ecgc._aagc = [2]*InvoiceCell{_ecgc.newCell("\u0044\u0075\u0065\u0020\u0044\u0061\u0074\u0065", _ecgc._eeab), _ecgc.newCell("", _ecgc._eeab)}
	_ecgc._gfag = [2]*InvoiceCell{_ecgc.newCell("\u0053\u0075\u0062\u0074\u006f\u0074\u0061\u006c", _ecgc._cbefc), _ecgc.newCell("", _ecgc._cbefc)}
	_decb := _ecgc._cbefc
	_decb.TextStyle = _dgac
	_decb.BackgroundColor = _dgcc
	_decb.BorderColor = _dgcc
	_ecgc._eacf = [2]*InvoiceCell{_ecgc.newCell("\u0054\u006f\u0074a\u006c", _decb), _ecgc.newCell("", _decb)}
	_ecgc._aeee = [2]string{"\u004e\u006f\u0074e\u0073", ""}
	_ecgc._ggbc = [2]string{"T\u0065r\u006d\u0073\u0020\u0061\u006e\u0064\u0020\u0063o\u006e\u0064\u0069\u0074io\u006e\u0073", ""}
	_ecgc._cdfg = []*InvoiceCell{_ecgc.newColumn("D\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", CellHorizontalAlignmentLeft), _ecgc.newColumn("\u0051\u0075\u0061\u006e\u0074\u0069\u0074\u0079", CellHorizontalAlignmentRight), _ecgc.newColumn("\u0055\u006e\u0069\u0074\u0020\u0070\u0072\u0069\u0063\u0065", CellHorizontalAlignmentRight), _ecgc.newColumn("\u0041\u006d\u006f\u0075\u006e\u0074", CellHorizontalAlignmentRight)}
	return _ecgc
}

// SetFillColor sets the fill color of the rectangle.
func (_fdef *Rectangle) SetFillColor(col Color) { _fdef._gbfa = col }

// SetFillColor sets the fill color.
func (_degd *CurvePolygon) SetFillColor(color Color) {
	_degd._aged = color
	_degd._efbeg.FillColor = _fce(color)
}

func (_dgfgcc *templateProcessor) parseList(_ggdd *templateNode) (interface{}, error) {
	_efbff := _dgfgcc.creator.NewList()
	for _, _fecca := range _ggdd._acdf.Attr {
		_ecdg := _fecca.Value
		switch _egce := _fecca.Name.Local; _egce {
		case "\u0069\u006e\u0064\u0065\u006e\u0074":
			_efbff.SetIndent(_dgfgcc.parseFloatAttr(_egce, _ecdg))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_daggd := _dgfgcc.parseMarginAttr(_egce, _ecdg)
			_efbff.SetMargins(_daggd.Left, _daggd.Right, _daggd.Top, _daggd.Bottom)
		default:
			_dgfgcc.nodeLogDebug(_ggdd, "\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u006c\u0069\u0073\u0074 \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _egce)
		}
	}
	return _efbff, nil
}

// SetSellerAddress sets the seller address of the invoice.
func (_bgfcd *Invoice) SetSellerAddress(address *InvoiceAddress) { _bgfcd._dceeg = address }

// Table allows organizing content in an rows X columns matrix, which can spawn across multiple pages.
type Table struct {
	_dbbf         int
	_edfe         int
	_ebad         int
	_eeggd        []float64
	_ecab         []float64
	_bebbc        float64
	_gebec        []*TableCell
	_bacd         []int
	_efce         Positioning
	_aabe, _bfbac float64
	_bebd         Margins
	_gdaae        bool
	_gcdff        int
	_bgec         int
	_ggea         bool
	_gbdb         bool
	_adbfb        bool
}

// NewBlockFromPage creates a Block from a PDF Page.  Useful for loading template pages as blocks
// from a PDF document and additional content with the creator.
func NewBlockFromPage(page *_ga.PdfPage) (*Block, error) {
	_bcg := &Block{}
	_dfd, _ddg := page.GetAllContentStreams()
	if _ddg != nil {
		return nil, _ddg
	}
	_eba := _da.NewContentStreamParser(_dfd)
	_ee, _ddg := _eba.Parse()
	if _ddg != nil {
		return nil, _ddg
	}
	_ee.WrapIfNeeded()
	_bcg._fa = _ee
	if page.Resources != nil {
		_bcg._ccf = page.Resources
	} else {
		_bcg._ccf = _ga.NewPdfPageResources()
	}
	_gff, _ddg := page.GetMediaBox()
	if _ddg != nil {
		return nil, _ddg
	}
	if _gff.Llx != 0 || _gff.Lly != 0 {
		_bcg.translate(-_gff.Llx, _gff.Lly)
	}
	_bcg._fd = _gff.Urx - _gff.Llx
	_bcg._cce = _gff.Ury - _gff.Lly
	if page.Rotate != nil {
		_bcg._be = -float64(*page.Rotate)
	}
	return _bcg, nil
}

// SetAnchor set gradient position anchor.
// Default to center.
func (_bgbeb *RadialShading) SetAnchor(anchor AnchorPoint) { _bgbeb._bafe = anchor }

func (_edca *StyledParagraph) appendChunk(_fgcaf *TextChunk) *TextChunk {
	_edca._eddc = append(_edca._eddc, _fgcaf)
	_edca.wrapText()
	return _fgcaf
}

func (_cab *Block) addContentsByString(_eae string) error {
	_cfd := _da.NewContentStreamParser(_eae)
	_cfdb, _ebd := _cfd.Parse()
	if _ebd != nil {
		return _ebd
	}
	_cab._fa.WrapIfNeeded()
	_cfdb.WrapIfNeeded()
	*_cab._fa = append(*_cab._fa, *_cfdb...)
	return nil
}

// SetColorBottom sets border color for bottom.
func (_dcac *border) SetColorBottom(col Color) { _dcac._adc = col }

// TextRenderingMode determines whether showing text shall cause glyph
// outlines to be stroked, filled, used as a clipping boundary, or some
// combination of the three.
// See section 9.3 "Text State Parameters and Operators" and
// Table 106 (pp. 254-255 PDF32000_2008).
type TextRenderingMode int

func (_befgg *templateProcessor) parseCellBorderStyleAttr(_cbccc, _edeac string) CellBorderStyle {
	_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020c\u0065\u006c\u006c b\u006f\u0072\u0064\u0065\u0072\u0020s\u0074\u0079\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025s\u0029\u002e", _cbccc, _edeac)
	_geccc := map[string]CellBorderStyle{"\u006e\u006f\u006e\u0065": CellBorderStyleNone, "\u0073\u0069\u006e\u0067\u006c\u0065": CellBorderStyleSingle, "\u0064\u006f\u0075\u0062\u006c\u0065": CellBorderStyleDouble}[_edeac]
	return _geccc
}
func _baeaf(_afeee float64, _bdgde float64) float64 { return _cd.Round(_afeee/_bdgde) * _bdgde }

// SetIndent sets the cell's left indent.
func (_ffcge *TableCell) SetIndent(indent float64) { _ffcge._fccd = indent }

func (_fgaa *FilledCurve) draw(_aace *Block, _cbb string) ([]byte, *_ga.PdfRectangle, error) {
	_bgbdg := _ff.NewCubicBezierPath()
	for _, _abgdf := range _fgaa._gefc {
		_bgbdg = _bgbdg.AppendCurve(_abgdf)
	}
	creator := _da.NewContentCreator()
	creator.Add_q()
	if _fgaa.FillEnabled && _fgaa._fbgc != nil {
		_efac := _fce(_fgaa._fbgc)
		_dgcea := _acdb(_aace, _efac, _fgaa._fbgc, func() Rectangle {
			_febf := _ff.NewCubicBezierPath()
			for _, _debf := range _fgaa._gefc {
				_febf = _febf.AppendCurve(_debf)
			}
			_cfacc := _febf.GetBoundingBox()
			if _fgaa.BorderEnabled {
				_cfacc.Height += _fgaa.BorderWidth
				_cfacc.Width += _fgaa.BorderWidth
				_cfacc.X -= _fgaa.BorderWidth / 2
				_cfacc.Y -= _fgaa.BorderWidth / 2
			}
			return Rectangle{_fcgf: _cfacc.X, _cbaa: _cfacc.Y, _bcffg: _cfacc.Width, _ebegg: _cfacc.Height}
		})
		if _dgcea != nil {
			return nil, nil, _dgcea
		}
		creator.SetNonStrokingColor(_efac)
	}
	if _fgaa.BorderEnabled {
		if _fgaa._afe != nil {
			creator.SetStrokingColor(_fce(_fgaa._afe))
		}
		creator.Add_w(_fgaa.BorderWidth)
	}
	if len(_cbb) > 1 {
		creator.Add_gs(_cc.PdfObjectName(_cbb))
	}
	_ff.DrawBezierPathWithCreator(_bgbdg, creator)
	creator.Add_h()
	if _fgaa.FillEnabled && _fgaa.BorderEnabled {
		creator.Add_B()
	} else if _fgaa.FillEnabled {
		creator.Add_f()
	} else if _fgaa.BorderEnabled {
		creator.Add_S()
	}
	creator.Add_Q()
	_cage := _bgbdg.GetBoundingBox()
	if _fgaa.BorderEnabled {
		_cage.Height += _fgaa.BorderWidth
		_cage.Width += _fgaa.BorderWidth
		_cage.X -= _fgaa.BorderWidth / 2
		_cage.Y -= _fgaa.BorderWidth / 2
	}
	_dgeb := &_ga.PdfRectangle{}
	_dgeb.Llx = _cage.X
	_dgeb.Lly = _cage.Y
	_dgeb.Urx = _cage.X + _cage.Width
	_dgeb.Ury = _cage.Y + _cage.Height
	return creator.Bytes(), _dgeb, nil
}

type templateTag struct {
	_cdcdf  map[string]struct{}
	_aecbab func(*templateProcessor, *templateNode) (interface{}, error)
}

// Indent returns the left offset of the list when nested into another list.
func (_cdag *List) Indent() float64 { return _cdag._addd }

// InvoiceCellProps holds all style properties for an invoice cell.
type InvoiceCellProps struct {
	TextStyle       TextStyle
	Alignment       CellHorizontalAlignment
	BackgroundColor Color
	BorderColor     Color
	BorderWidth     float64
	BorderSides     []CellBorderSide
}

// EnableFontSubsetting enables font subsetting for `font` when the creator output is written to file.
// Embeds only the subset of the runes/glyphs that are actually used to display the file.
// Subsetting can reduce the size of fonts significantly.
func (_dgff *Creator) EnableFontSubsetting(font *_ga.PdfFont) {
	_dgff._abdf = append(_dgff._abdf, font)
}

// Width returns the width of the graphic svg.
func (_ceeg *GraphicSVG) Width() float64 { return _ceeg._debae.Width }

// GetCoords returns the upper left corner coordinates of the rectangle (`x`, `y`).
func (_gaegf *Rectangle) GetCoords() (float64, float64) { return _gaegf._fcgf, _gaegf._cbaa }

func (_geea *StyledParagraph) getTextHeight() float64 {
	var _aaefc float64
	for _, _ffcca := range _geea._eddc {
		_cgeea := _ffcca.Style.FontSize * _geea._bebe
		if _cgeea > _aaefc {
			_aaefc = _cgeea
		}
	}
	return _aaefc
}

// PageFinalizeFunctionArgs holds the input arguments provided to the page
// finalize callback function which can be set using Creator.PageFinalize.
type PageFinalizeFunctionArgs struct {
	PageNum    int
	PageWidth  float64
	PageHeight float64
	TOCPages   int
	TotalPages int
}

func (_gcec *Division) split(_ebagf DrawContext) (_bggde, _eeed *Division) {
	var (
		_bbef        float64
		_cdee, _bcff []VectorDrawable
	)
	_degf := _ebagf.Width - _gcec._acaa.Left - _gcec._acaa.Right - _gcec._cgbb.Left - _gcec._cgbb.Right
	for _dbda, _aedfc := range _gcec._dccd {
		_bbef += _fecea(_aedfc, _degf)
		if _bbef < _ebagf.Height {
			_cdee = append(_cdee, _aedfc)
		} else {
			_bcff = _gcec._dccd[_dbda:]
			break
		}
	}
	if len(_cdee) > 0 {
		_bggde = _dfdab()
		*_bggde = *_gcec
		_bggde._dccd = _cdee
		if _gcec._eeaf != nil {
			_bggde._eeaf = &Background{}
			*_bggde._eeaf = *_gcec._eeaf
		}
	}
	if len(_bcff) > 0 {
		_eeed = _dfdab()
		*_eeed = *_gcec
		_eeed._dccd = _bcff
		if _gcec._eeaf != nil {
			_eeed._eeaf = &Background{}
			*_eeed._eeaf = *_gcec._eeaf
		}
	}
	return _bggde, _eeed
}

// PageBreak represents a page break for a chapter.
type PageBreak struct{}

// GeneratePageBlocks generates the page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages. Implements the Drawable interface.
func (_eccae *StyledParagraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_dfcgd := ctx
	var _daga []*Block
	_fecd := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _eccae._bedd.IsRelative() {
		ctx.X += _eccae._fadcg.Left
		ctx.Y += _eccae._fadcg.Top
		ctx.Width -= _eccae._fadcg.Left + _eccae._fadcg.Right
		ctx.Height -= _eccae._fadcg.Top
		_eccae.SetWidth(ctx.Width)
	} else {
		if int(_eccae._cccde) <= 0 {
			_eccae.SetWidth(_eccae.getTextWidth() / 1000.0)
		}
		ctx.X = _eccae._fcadg
		ctx.Y = _eccae._ffdc
	}
	if _eccae._egdcb != nil {
		_eccae._egdcb(_eccae, ctx)
	}
	if _bbdg := _eccae.wrapText(); _bbdg != nil {
		return nil, ctx, _bbdg
	}
	_gceef := _eccae._bddb
	for {
		_affeg, _eaefd, _cecff := _bceg(_fecd, _eccae, _gceef, ctx)
		if _cecff != nil {
			_bcd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cecff)
			return nil, ctx, _cecff
		}
		ctx = _affeg
		_daga = append(_daga, _fecd)
		if _gceef = _eaefd; len(_eaefd) == 0 {
			break
		}
		_fecd = NewBlock(ctx.PageWidth, ctx.PageHeight)
		ctx.Page++
		_affeg = ctx
		_affeg.Y = ctx.Margins.Top
		_affeg.X = ctx.Margins.Left + _eccae._fadcg.Left
		_affeg.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
		_affeg.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _eccae._fadcg.Left - _eccae._fadcg.Right
		ctx = _affeg
	}
	if _eccae._bedd.IsRelative() {
		ctx.Y += _eccae._fadcg.Bottom
		ctx.Height -= _eccae._fadcg.Bottom
		if !ctx.Inline {
			ctx.X = _dfcgd.X
			ctx.Width = _dfcgd.Width
		}
		return _daga, ctx, nil
	}
	return _daga, _dfcgd, nil
}

func (_abde *Division) ctxHeight(_cccf float64) float64 {
	_cccf -= _abde._acaa.Left + _abde._acaa.Right + _abde._cgbb.Left + _abde._cgbb.Right
	var _gcdd float64
	for _, _agab := range _abde._dccd {
		_gcdd += _fecea(_agab, _cccf)
	}
	return _gcdd
}

// SetMaxLines sets the maximum number of lines before the paragraph
// text is truncated.
func (_egbc *Paragraph) SetMaxLines(maxLines int) { _egbc._gfbef = maxLines; _egbc.wrapText() }

// Scale scales the ellipse dimensions by the specified factors.
func (_ebcdc *Ellipse) Scale(xFactor, yFactor float64) {
	_ebcdc._eded = xFactor * _ebcdc._eded
	_ebcdc._bdcce = yFactor * _ebcdc._bdcce
}

func (_ffgdfd *templateProcessor) parseColorAttr(_cfea, _ggfa string) Color {
	_bcd.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0063\u006f\u006c\u006fr a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _cfea, _ggfa)
	_ggfa = _a.TrimSpace(_ggfa)
	if _a.HasPrefix(_ggfa, "\u006c\u0069n\u0065\u0061\u0072-\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0028") && _a.HasSuffix(_ggfa, "\u0029") && len(_ggfa) > 17 {
		return _ffgdfd.parseLinearGradientAttr(_ffgdfd.creator, _ggfa)
	}
	if _a.HasPrefix(_ggfa, "\u0072\u0061d\u0069\u0061\u006c-\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0028") && _a.HasSuffix(_ggfa, "\u0029") && len(_ggfa) > 17 {
		return _ffgdfd.parseRadialGradientAttr(_ffgdfd.creator, _ggfa)
	}
	if _feeeba := _ffgdfd.parseColor(_ggfa); _feeeba != nil {
		return _feeeba
	}
	return ColorBlack
}

// SetTOC sets the table of content component of the creator.
// This method should be used when building a custom table of contents.
func (_eege *Creator) SetTOC(toc *TOC) {
	if toc == nil {
		return
	}
	_eege._edad = toc
}

// SetLevelOffset sets the amount of space an indentation level occupies.
func (_dedea *TOCLine) SetLevelOffset(levelOffset float64) {
	_dedea._dgefgc = levelOffset
	_dedea._eedce._fadcg.Left = _dedea._dbdgg + float64(_dedea._dffba-1)*_dedea._dgefgc
}

// ToPdfShadingPattern generates a new model.PdfShadingPatternType3 object.
func (_cbgcg *RadialShading) ToPdfShadingPattern() *_ga.PdfShadingPatternType3 {
	_edcg, _cbee, _ebga := _cbgcg._acdae._cfec.ToRGB()
	_bbcc := _cbgcg.shadingModel()
	_bbcc.PdfShading.Background = _cc.MakeArrayFromFloats([]float64{_edcg, _cbee, _ebga})
	_bgdd := _ga.NewPdfShadingPatternType3()
	_bgdd.Shading = _bbcc
	return _bgdd
}

// FillOpacity returns the fill opacity of the ellipse (0-1).
func (_ddabc *Ellipse) FillOpacity() float64 { return _ddabc._bbdc }

// NewColorPoint creates a new color and point object for use in the gradient rendering process.
func NewColorPoint(color Color, point float64) *ColorPoint { return _gebe(color, point) }

func _bgfbc(_daff int64, _gfde, _befea, _fdgfg float64) *_ga.PdfAnnotation {
	_bbgfe := _ga.NewPdfAnnotationLink()
	_cccea := _ga.NewBorderStyle()
	_cccea.SetBorderWidth(0)
	_bbgfe.BS = _cccea.ToPdfObject()
	if _daff < 0 {
		_daff = 0
	}
	_bbgfe.Dest = _cc.MakeArray(_cc.MakeInteger(_daff), _cc.MakeName("\u0058\u0059\u005a"), _cc.MakeFloat(_gfde), _cc.MakeFloat(_befea), _cc.MakeFloat(_fdgfg))
	return _bbgfe.PdfAnnotation
}

// Terms returns the terms and conditions section of the invoice as a
// title-content pair.
func (_gdbgb *Invoice) Terms() (string, string) { return _gdbgb._ggbc[0], _gdbgb._ggbc[1] }

func (_ggge *List) ctxHeight(_agabd float64) float64 {
	_agabd -= _ggge._addd
	var _beffa float64
	for _, _eedc := range _ggge._bdccg {
		_beffa += _eedc.ctxHeight(_agabd)
	}
	return _beffa
}

// SetStyleLeft sets border style for left side.
func (_facf *border) SetStyleLeft(style CellBorderStyle) { _facf._dad = style }

// BorderWidth returns the border width of the rectangle.
func (_ebbf *Rectangle) BorderWidth() float64 { return _ebbf._bffc }

// Height returns the height of the Paragraph. The height is calculated based on the input text and how it is wrapped
// within the container. Does not include Margins.
func (_fbed *StyledParagraph) Height() float64 {
	_fbed.wrapText()
	var _dadd float64
	for _, _bcfc := range _fbed._bddb {
		var _ggcc float64
		for _, _abfd := range _bcfc {
			_aggbd := _fbed._bebe * _abfd.Style.FontSize
			if _aggbd > _ggcc {
				_ggcc = _aggbd
			}
		}
		_dadd += _ggcc
	}
	return _dadd
}

// SetPadding sets the padding of the component. The padding represents
// inner margins which are applied around the contents of the division.
// The background of the component is not affected by its padding.
func (_cegg *Division) SetPadding(left, right, top, bottom float64) {
	_cegg._cgbb.Left = left
	_cegg._cgbb.Right = right
	_cegg._cgbb.Top = top
	_cegg._cgbb.Bottom = bottom
}

func _gbdbf(_ggeb *_ga.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 0), Font: _ggeb, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}

func (_efdbb *templateProcessor) parseTextChunk(_afcc *templateNode, _gddbg *TextChunk) (interface{}, error) {
	if _afcc._gbdf == nil {
		_efdbb.nodeLogError(_afcc, "\u0054\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u0020\u0070\u0061\u0072\u0065n\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _bfegb
	}
	var (
		_ccgd  = _efdbb.creator.NewTextStyle()
		_ffaaa bool
	)
	for _, _dfeed := range _afcc._acdf.Attr {
		if _dfeed.Name.Local == "\u006c\u0069\u006e\u006b" {
			_egeg, _cdeffg := _afcc._gbdf._defae.(*StyledParagraph)
			if !_cdeffg {
				_efdbb.nodeLogError(_afcc, "\u004c\u0069\u006e\u006b \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u006f\u006e\u006c\u0079\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0070\u0061\u0072\u0061\u0067r\u0061\u0070\u0068\u0027\u0073\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u002e")
				_ffaaa = true
			} else {
				_ccgd = _egeg._dbgf
			}
			break
		}
	}
	if _gddbg == nil {
		_gddbg = NewTextChunk("", _ccgd)
	}
	for _, _ebcbg := range _afcc._acdf.Attr {
		_feec := _ebcbg.Value
		switch _acegb := _ebcbg.Name.Local; _acegb {
		case "\u0063\u006f\u006co\u0072":
			_gddbg.Style.Color = _efdbb.parseColorAttr(_acegb, _feec)
		case "\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_gddbg.Style.OutlineColor = _efdbb.parseColorAttr(_acegb, _feec)
		case "\u0066\u006f\u006e\u0074":
			_gddbg.Style.Font = _efdbb.parseFontAttr(_acegb, _feec)
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_gddbg.Style.FontSize = _efdbb.parseFloatAttr(_acegb, _feec)
		case "\u006f\u0075\u0074l\u0069\u006e\u0065\u002d\u0073\u0069\u007a\u0065":
			_gddbg.Style.OutlineSize = _efdbb.parseFloatAttr(_acegb, _feec)
		case "\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u002d\u0073\u0070a\u0063\u0069\u006e\u0067":
			_gddbg.Style.CharSpacing = _efdbb.parseFloatAttr(_acegb, _feec)
		case "\u0068o\u0072i\u007a\u006f\u006e\u0074\u0061l\u002d\u0073c\u0061\u006c\u0069\u006e\u0067":
			_gddbg.Style.HorizontalScaling = _efdbb.parseFloatAttr(_acegb, _feec)
		case "\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067-\u006d\u006f\u0064\u0065":
			_gddbg.Style.RenderingMode = _efdbb.parseTextRenderingModeAttr(_acegb, _feec)
		case "\u0075n\u0064\u0065\u0072\u006c\u0069\u006ee":
			_gddbg.Style.Underline = _efdbb.parseBoolAttr(_acegb, _feec)
		case "\u0075n\u0064e\u0072\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_gddbg.Style.UnderlineStyle.Color = _efdbb.parseColorAttr(_acegb, _feec)
		case "\u0075\u006ed\u0065\u0072\u006ci\u006e\u0065\u002d\u006f\u0066\u0066\u0073\u0065\u0074":
			_gddbg.Style.UnderlineStyle.Offset = _efdbb.parseFloatAttr(_acegb, _feec)
		case "\u0075\u006e\u0064\u0065rl\u0069\u006e\u0065\u002d\u0074\u0068\u0069\u0063\u006b\u006e\u0065\u0073\u0073":
			_gddbg.Style.UnderlineStyle.Thickness = _efdbb.parseFloatAttr(_acegb, _feec)
		case "\u006c\u0069\u006e\u006b":
			if !_ffaaa {
				_gddbg._cgfaa = _efdbb.parseLinkAttr(_acegb, _feec)
			}
		case "\u0074e\u0078\u0074\u002d\u0072\u0069\u0073e":
			_gddbg.Style.TextRise = _efdbb.parseFloatAttr(_acegb, _feec)
		default:
			_efdbb.nodeLogDebug(_afcc, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _acegb)
		}
	}
	return _gddbg, nil
}

// GeneratePageBlocks draws the line on a new block representing the page.
// Implements the Drawable interface.
func (_aaecf *Line) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_eefcc        []*Block
		_bbaa         = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_ddag         = ctx
		_dcdb, _efacd = _aaecf._ccg, ctx.PageHeight - _aaecf._eggf
		_gada, _ggfee = _aaecf._dfcae, ctx.PageHeight - _aaecf._feafb
	)
	_ebeef := _aaecf._fcdd.IsRelative()
	if _ebeef {
		ctx.X += _aaecf._gabb.Left
		ctx.Y += _aaecf._gabb.Top
		ctx.Width -= _aaecf._gabb.Left + _aaecf._gabb.Right
		ctx.Height -= _aaecf._gabb.Top + _aaecf._gabb.Bottom
		_dcdb, _efacd, _gada, _ggfee = _aaecf.computeCoords(ctx)
		if _aaecf.Height() > ctx.Height {
			_eefcc = append(_eefcc, _bbaa)
			_bbaa = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_gdfg := ctx
			_gdfg.Y = ctx.Margins.Top + _aaecf._gabb.Top
			_gdfg.X = ctx.Margins.Left + _aaecf._gabb.Left
			_gdfg.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _aaecf._gabb.Top - _aaecf._gabb.Bottom
			_gdfg.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _aaecf._gabb.Left - _aaecf._gabb.Right
			ctx = _gdfg
			_dcdb, _efacd, _gada, _ggfee = _aaecf.computeCoords(ctx)
		}
	}
	_cfcea := _ff.BasicLine{X1: _dcdb, Y1: _efacd, X2: _gada, Y2: _ggfee, LineColor: _fce(_aaecf._efdf), Opacity: _aaecf._gbcb, LineWidth: _aaecf._bfe, LineStyle: _aaecf._ffgbg, DashArray: _aaecf._fedd, DashPhase: _aaecf._ffgdf}
	_ddgb, _baaee := _bbaa.setOpacity(1.0, _aaecf._gbcb)
	if _baaee != nil {
		return nil, ctx, _baaee
	}
	_ccb, _, _baaee := _cfcea.Draw(_ddgb)
	if _baaee != nil {
		return nil, ctx, _baaee
	}
	if _baaee = _bbaa.addContentsByString(string(_ccb)); _baaee != nil {
		return nil, ctx, _baaee
	}
	if _ebeef {
		ctx.X = _ddag.X
		ctx.Width = _ddag.Width
		_bbeb := _aaecf.Height()
		ctx.Y += _bbeb + _aaecf._gabb.Bottom
		ctx.Height -= _bbeb
	} else {
		ctx = _ddag
	}
	_eefcc = append(_eefcc, _bbaa)
	return _eefcc, ctx, nil
}

func (_fced *Paragraph) getTextLineWidth(_aegg string) float64 {
	var _cbgb float64
	for _, _fbbe := range _aegg {
		if _fbbe == '\u000A' {
			continue
		}
		_bedcc, _dccf := _fced._bfbd.GetRuneMetrics(_fbbe)
		if !_dccf {
			_bcd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _fbbe, _fbbe)
			return -1
		}
		_cbgb += _fced._dbad * _bedcc.Wx
	}
	return _cbgb
}

func (_dabc *templateProcessor) parseParagraph(_aefe *templateNode, _dfbbg *Paragraph) (interface{}, error) {
	if _dfbbg == nil {
		_dfbbg = _dabc.creator.NewParagraph("")
	}
	for _, _cedcg := range _aefe._acdf.Attr {
		_gged := _cedcg.Value
		switch _dfcgg := _cedcg.Name.Local; _dfcgg {
		case "\u0066\u006f\u006e\u0074":
			_dfbbg.SetFont(_dabc.parseFontAttr(_dfcgg, _gged))
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_dfbbg.SetFontSize(_dabc.parseFloatAttr(_dfcgg, _gged))
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_dfbbg.SetTextAlignment(_dabc.parseTextAlignmentAttr(_dfcgg, _gged))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_dfbbg.SetLineHeight(_dabc.parseFloatAttr(_dfcgg, _gged))
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_dfbbg.SetEnableWrap(_dabc.parseBoolAttr(_dfcgg, _gged))
		case "\u0063\u006f\u006co\u0072":
			_dfbbg.SetColor(_dabc.parseColorAttr(_dfcgg, _gged))
		case "\u0078":
			_dfbbg.SetPos(_dabc.parseFloatAttr(_dfcgg, _gged), _dfbbg._bge)
		case "\u0079":
			_dfbbg.SetPos(_dfbbg._acab, _dabc.parseFloatAttr(_dfcgg, _gged))
		case "\u0061\u006e\u0067l\u0065":
			_dfbbg.SetAngle(_dabc.parseFloatAttr(_dfcgg, _gged))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_faddf := _dabc.parseMarginAttr(_dfcgg, _gged)
			_dfbbg.SetMargins(_faddf.Left, _faddf.Right, _faddf.Top, _faddf.Bottom)
		case "\u006da\u0078\u002d\u006c\u0069\u006e\u0065s":
			_dfbbg.SetMaxLines(int(_dabc.parseInt64Attr(_dfcgg, _gged)))
		default:
			_dabc.nodeLogDebug(_aefe, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020pa\u0072\u0061\u0067\u0072\u0061\u0070h\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _dfcgg)
		}
	}
	return _dfbbg, nil
}

func (_egfcgb *templateProcessor) parseCellAlignmentAttr(_efeb, _gdgfc string) CellHorizontalAlignment {
	_bcd.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020c\u0065\u006c\u006c\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _efeb, _gdgfc)
	_eaagg := map[string]CellHorizontalAlignment{"\u006c\u0065\u0066\u0074": CellHorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": CellHorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": CellHorizontalAlignmentRight}[_gdgfc]
	return _eaagg
}

// SetContent sets the cell's content.  The content is a VectorDrawable, i.e.
// a Drawable with a known height and width.
// Currently supported VectorDrawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
// - *Table
// - *Division
// - *List
// - *Rectangle
// - *Ellipse
// - *Line
func (_dgdf *TableCell) SetContent(vd VectorDrawable) error {
	switch _aefd := vd.(type) {
	case *Paragraph:
		if _aefd._deaf {
			_aefd._dbfb = true
		}
		_dgdf._bfef = vd
	case *StyledParagraph:
		if _aefd._afeg {
			_aefd._bfcgb = true
		}
		_dgdf._bfef = vd
	case *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line:
		_dgdf._bfef = vd
	default:
		_bcd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0079p\u0065\u0020\u0025\u0054", vd)
		return _cc.ErrTypeError
	}
	return nil
}

// TOCLine represents a line in a table of contents.
// The component can be used both in the context of a
// table of contents component and as a standalone component.
// The representation of a table of contents line is as follows:
/*
         [number] [title]      [separator] [page]
   e.g.: Chapter1 Introduction ........... 1
*/
type TOCLine struct {
	_eedce *StyledParagraph

	// Holds the text and style of the number part of the TOC line.
	Number TextChunk

	// Holds the text and style of the title part of the TOC line.
	Title TextChunk

	// Holds the text and style of the separator part of the TOC line.
	Separator TextChunk

	// Holds the text and style of the page part of the TOC line.
	Page    TextChunk
	_dbdgg  float64
	_dffba  uint
	_dgefgc float64
	_cdfbe  Positioning
	_dfbbb  float64
	_gcab   float64
	_dfbfg  int64
}

func _bcbd() *FilledCurve {
	_adbd := FilledCurve{}
	_adbd._gefc = []_ff.CubicBezierCurve{}
	return &_adbd
}

// SetNoteHeadingStyle sets the style properties used to render the heading
// of the invoice note sections.
func (_bgab *Invoice) SetNoteHeadingStyle(style TextStyle) { _bgab._dede = style }

func _gbfe(_bced *_ga.PdfAnnotationLink) *_ga.PdfAnnotationLink {
	if _bced == nil {
		return nil
	}
	_fbfe := _ga.NewPdfAnnotationLink()
	_fbfe.BS = _bced.BS
	_fbfe.A = _bced.A
	if _bdffc, _aecee := _bced.GetAction(); _aecee == nil && _bdffc != nil {
		_fbfe.SetAction(_bdffc)
	}
	if _ecde, _cfadd := _bced.Dest.(*_cc.PdfObjectArray); _cfadd {
		_fbfe.Dest = _cc.MakeArray(_ecde.Elements()...)
	}
	return _fbfe
}

// SetLineStyle sets the style for all the line components: number, title,
// separator, page. The style is applied only for new lines added to the
// TOC component.
func (_fafc *TOC) SetLineStyle(style TextStyle) {
	_fafc.SetLineNumberStyle(style)
	_fafc.SetLineTitleStyle(style)
	_fafc.SetLineSeparatorStyle(style)
	_fafc.SetLinePageStyle(style)
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_bbbe *Invoice) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ebfb := ctx
	_dcca := []func(_gcaea DrawContext) ([]*Block, DrawContext, error){_bbbe.generateHeaderBlocks, _bbbe.generateInformationBlocks, _bbbe.generateLineBlocks, _bbbe.generateTotalBlocks, _bbbe.generateNoteBlocks}
	var _aacaf []*Block
	for _, _fgbf := range _dcca {
		_efbfe, _acad, _aagcc := _fgbf(ctx)
		if _aagcc != nil {
			return _aacaf, ctx, _aagcc
		}
		if len(_aacaf) == 0 {
			_aacaf = _efbfe
		} else if len(_efbfe) > 0 {
			_aacaf[len(_aacaf)-1].mergeBlocks(_efbfe[0])
			_aacaf = append(_aacaf, _efbfe[1:]...)
		}
		ctx = _acad
	}
	if _bbbe._aaede.IsRelative() {
		ctx.X = _ebfb.X
	}
	if _bbbe._aaede.IsAbsolute() {
		return _aacaf, _ebfb, nil
	}
	return _aacaf, ctx, nil
}

// SetBorder sets the cell's border style.
func (_fcdf *TableCell) SetBorder(side CellBorderSide, style CellBorderStyle, width float64) {
	if style == CellBorderStyleSingle && side == CellBorderSideAll {
		_fcdf._ebce = CellBorderStyleSingle
		_fcdf._aabc = width
		_fcdf._ggce = CellBorderStyleSingle
		_fcdf._aeba = width
		_fcdf._bage = CellBorderStyleSingle
		_fcdf._edag = width
		_fcdf._efdbd = CellBorderStyleSingle
		_fcdf._fgbd = width
	} else if style == CellBorderStyleDouble && side == CellBorderSideAll {
		_fcdf._ebce = CellBorderStyleDouble
		_fcdf._aabc = width
		_fcdf._ggce = CellBorderStyleDouble
		_fcdf._aeba = width
		_fcdf._bage = CellBorderStyleDouble
		_fcdf._edag = width
		_fcdf._efdbd = CellBorderStyleDouble
		_fcdf._fgbd = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideLeft {
		_fcdf._ebce = style
		_fcdf._aabc = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideBottom {
		_fcdf._ggce = style
		_fcdf._aeba = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideRight {
		_fcdf._bage = style
		_fcdf._edag = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideTop {
		_fcdf._efdbd = style
		_fcdf._fgbd = width
	}
}

// GeneratePageBlocks generates the page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages. Implements the Drawable interface.
func (_ddcf *Paragraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_abfb := ctx
	var _geaa []*Block
	_dbcde := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _ddcf._gece.IsRelative() {
		ctx.X += _ddcf._bcggf.Left
		ctx.Y += _ddcf._bcggf.Top
		ctx.Width -= _ddcf._bcggf.Left + _ddcf._bcggf.Right
		ctx.Height -= _ddcf._bcggf.Top
		_ddcf.SetWidth(ctx.Width)
		if _ddcf.Height() > ctx.Height {
			_geaa = append(_geaa, _dbcde)
			_dbcde = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_bgaf := ctx
			_bgaf.Y = ctx.Margins.Top
			_bgaf.X = ctx.Margins.Left + _ddcf._bcggf.Left
			_bgaf.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
			_bgaf.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _ddcf._bcggf.Left - _ddcf._bcggf.Right
			ctx = _bgaf
		}
	} else {
		if int(_ddcf._adcfe) <= 0 {
			_ddcf.SetWidth(_ddcf.getTextWidth())
		}
		ctx.X = _ddcf._acab
		ctx.Y = _ddcf._bge
	}
	ctx, _ccdga := _efff(_dbcde, _ddcf, ctx)
	if _ccdga != nil {
		_bcd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ccdga)
		return nil, ctx, _ccdga
	}
	_geaa = append(_geaa, _dbcde)
	if _ddcf._gece.IsRelative() {
		ctx.Y += _ddcf._bcggf.Bottom
		ctx.Height -= _ddcf._bcggf.Bottom
		if !ctx.Inline {
			ctx.X = _abfb.X
			ctx.Width = _abfb.Width
		}
		return _geaa, ctx, nil
	}
	return _geaa, _abfb, nil
}

// Angle returns the block rotation angle in degrees.
func (_cg *Block) Angle() float64 { return _cg._be }

// AddTextItem appends a new item with the specified text to the list.
// The method creates a styled paragraph with the specified text and returns
// it so that the item style can be customized.
// The method also returns the marker used for the newly added item.
// The marker object can be used to change the text and style of the marker
// for the current item.
func (_bcab *List) AddTextItem(text string) (*StyledParagraph, *TextChunk, error) {
	_daaf := _cgfa(_bcab._edda)
	_daaf.Append(text)
	_fcgc, _acge := _bcab.Add(_daaf)
	return _daaf, _fcgc, _acge
}

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_cgag *LinearShading) SetAntiAlias(enable bool)   { _cgag._cfbfe.SetAntiAlias(enable) }
func (_efbe *Creator) setActivePage(_fbfb *_ga.PdfPage) { _efbe._abd = _fbfb }

const (
	TextVerticalAlignmentBaseline TextVerticalAlignment = iota
	TextVerticalAlignmentCenter
	TextVerticalAlignmentBottom
	TextVerticalAlignmentTop
)

func _bbff(_deeae *templateProcessor, _egdf *templateNode) (interface{}, error) {
	return _deeae.parseTableCell(_egdf)
}

func (_bebbg *Invoice) generateNoteBlocks(_dcegg DrawContext) ([]*Block, DrawContext, error) {
	_aebgb := _dfdab()
	_dbfea := append([][2]string{_bebbg._aeee, _bebbg._ggbc}, _bebbg._eefc...)
	for _, _dbag := range _dbfea {
		if _dbag[1] != "" {
			_dbgae := _bebbg.drawSection(_dbag[0], _dbag[1])
			for _, _gdec := range _dbgae {
				_aebgb.Add(_gdec)
			}
			_dgbe := _cgfa(_bebbg._dcefb)
			_dgbe.SetMargins(0, 0, 10, 0)
			_aebgb.Add(_dgbe)
		}
	}
	return _aebgb.GeneratePageBlocks(_dcegg)
}

// NewChapter creates a new chapter with the specified title as the heading.
func (_ebcd *Creator) NewChapter(title string) *Chapter {
	_ebcd._age++
	_eaffe := _ebcd.NewTextStyle()
	_eaffe.FontSize = 16
	return _ebda(nil, _ebcd._edad, _ebcd._bcb, title, _ebcd._age, _eaffe)
}

// MoveX moves the drawing context to absolute position x.
func (_feabb *Creator) MoveX(x float64) { _feabb._ggab.X = x }

// SetMargins sets the margins of the graphic svg component.
func (_gabd *GraphicSVG) SetMargins(left, right, top, bottom float64) {
	_gabd._fgfa.Left = left
	_gabd._fgfa.Right = right
	_gabd._fgfa.Top = top
	_gabd._fgfa.Bottom = bottom
}

// Number returns the invoice number description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_bgbb *Invoice) Number() (*InvoiceCell, *InvoiceCell) { return _bgbb._badb[0], _bgbb._badb[1] }

// SetEnableWrap sets the line wrapping enabled flag.
func (_bacfe *StyledParagraph) SetEnableWrap(enableWrap bool) {
	_bacfe._bfcgb = enableWrap
	_bacfe._afeg = false
}

// AddressHeadingStyle returns the style properties used to render the
// heading of the invoice address sections.
func (_bcfg *Invoice) AddressHeadingStyle() TextStyle { return _bcfg._bfcc }

func _eagdc(_acfdca *templateProcessor, _dbffc *templateNode) (interface{}, error) {
	return _acfdca.parsePageBreak(_dbffc)
}

// Margins represents page margins or margins around an element.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (_cbfb *templateProcessor) parseStyledParagraph(_ffggd *templateNode) (interface{}, error) {
	_edde := _cbfb.creator.NewStyledParagraph()
	for _, _cdcg := range _ffggd._acdf.Attr {
		_eaead := _cdcg.Value
		switch _gegf := _cdcg.Name.Local; _gegf {
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_edde.SetTextAlignment(_cbfb.parseTextAlignmentAttr(_gegf, _eaead))
		case "\u0076\u0065\u0072\u0074ic\u0061\u006c\u002d\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_edde.SetTextVerticalAlignment(_cbfb.parseTextVerticalAlignmentAttr(_gegf, _eaead))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_edde.SetLineHeight(_cbfb.parseFloatAttr(_gegf, _eaead))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_bfdbc := _cbfb.parseMarginAttr(_gegf, _eaead)
			_edde.SetMargins(_bfdbc.Left, _bfdbc.Right, _bfdbc.Top, _bfdbc.Bottom)
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_edde.SetEnableWrap(_cbfb.parseBoolAttr(_gegf, _eaead))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0077\u006f\u0072\u0064\u002d\u0077\u0072\u0061\u0070":
			_edde.EnableWordWrap(_cbfb.parseBoolAttr(_gegf, _eaead))
		case "\u0074\u0065\u0078\u0074\u002d\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077":
			_edde.SetTextOverflow(_cbfb.parseTextOverflowAttr(_gegf, _eaead))
		case "\u0078":
			_edde.SetPos(_cbfb.parseFloatAttr(_gegf, _eaead), _edde._ffdc)
		case "\u0079":
			_edde.SetPos(_edde._fcadg, _cbfb.parseFloatAttr(_gegf, _eaead))
		case "\u0061\u006e\u0067l\u0065":
			_edde.SetAngle(_cbfb.parseFloatAttr(_gegf, _eaead))
		default:
			_cbfb.nodeLogDebug(_ffggd, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0073\u0074\u0079l\u0065\u0064\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0060\u0025\u0073`.\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _gegf)
		}
	}
	return _edde, nil
}

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_cffea *shading) SetAntiAlias(enable bool) { _cffea._cebg = enable }

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_bcbe *Paragraph) GetMargins() (float64, float64, float64, float64) {
	return _bcbe._bcggf.Left, _bcbe._bcggf.Right, _bcbe._bcggf.Top, _bcbe._bcggf.Bottom
}

// CellVerticalAlignment defines the table cell's vertical alignment.
type CellVerticalAlignment int

// SetBorderColor sets the border color of the ellipse.
func (_beaf *Ellipse) SetBorderColor(col Color) { _beaf._bacf = col }

// GeneratePageBlocks generate the Page blocks. Draws the Image on a block, implementing the Drawable interface.
func (_aceg *Image) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	if _aceg._ebfc == nil {
		if _cbaf := _aceg.makeXObject(); _cbaf != nil {
			return nil, ctx, _cbaf
		}
	}
	var _gcgg []*Block
	_dceee := ctx
	_fdcf := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _aceg._geca.IsRelative() {
		_aceg.applyFitMode(ctx.Width)
		ctx.X += _aceg._eaac.Left
		ctx.Y += _aceg._eaac.Top
		ctx.Width -= _aceg._eaac.Left + _aceg._eaac.Right
		ctx.Height -= _aceg._eaac.Top + _aceg._eaac.Bottom
		if _aceg._cgdb > ctx.Height {
			_gcgg = append(_gcgg, _fdcf)
			_fdcf = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_dcdg := ctx
			_dcdg.Y = ctx.Margins.Top + _aceg._eaac.Top
			_dcdg.X = ctx.Margins.Left + _aceg._eaac.Left
			_dcdg.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _aceg._eaac.Top - _aceg._eaac.Bottom
			_dcdg.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _aceg._eaac.Left - _aceg._eaac.Right
			ctx = _dcdg
		}
	} else {
		ctx.X = _aceg._fdab
		ctx.Y = _aceg._egca
	}
	ctx, _dfgd := _dcabe(_fdcf, _aceg, ctx)
	if _dfgd != nil {
		return nil, ctx, _dfgd
	}
	_gcgg = append(_gcgg, _fdcf)
	if _aceg._geca.IsAbsolute() {
		ctx = _dceee
	} else {
		ctx.X = _dceee.X
		ctx.Width = _dceee.Width
		ctx.Y += _aceg._eaac.Bottom
	}
	return _gcgg, ctx, nil
}

type rgbColor struct{ _ddae, _cede, _fgef float64 }

func _dfdab() *Division { return &Division{_dccg: true} }

// Scale sets the scale ratio with `X` factor and `Y` factor for the graphic svg.
func (_bdedg *GraphicSVG) Scale(xFactor, yFactor float64) {
	_bdedg._debae.Width = xFactor * _bdedg._debae.Width
	_bdedg._debae.Height = yFactor * _bdedg._debae.Height
	_bdedg._debae.SetScaling(xFactor, yFactor)
}

// SetWidth sets the the Paragraph width. This is essentially the wrapping width,
// i.e. the width the text can extend to prior to wrapping over to next line.
func (_febe *StyledParagraph) SetWidth(width float64) { _febe._cccde = width; _febe.wrapText() }

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_bgcc *RadialShading) SetBackgroundColor(backgroundColor Color) {
	_bgcc._acdae.SetBackgroundColor(backgroundColor)
}

// SetLogo sets the logo of the invoice.
func (_fggd *Invoice) SetLogo(logo *Image) { _fggd._bagfe = logo }

func (_dgdaee *templateProcessor) nodeLogDebug(_edgg *templateNode, _fcgcg string, _bbdfg ...interface{}) {
	_bcd.Log.Debug(_dgdaee.getNodeErrorLocation(_edgg, _fcgcg, _bbdfg...))
}

func (_ade *List) split(_eead DrawContext) (_gadb, _dafb *List) {
	var (
		_ccgb        float64
		_afed, _gcee []*listItem
	)
	_ecbe := _eead.Width - _ade._fcag.Horizontal() - _ade._addd - _ade.markerWidth()
	_ebeg := _ade.markerWidth()
	for _addb, _eecc := range _ade._bdccg {
		_fbaea := _eecc.ctxHeight(_ecbe)
		_ccgb += _fbaea
		if _ccgb <= _eead.Height {
			_afed = append(_afed, _eecc)
		} else {
			switch _fccba := _eecc._cecd.(type) {
			case *List:
				_baff := _eead
				_baff.Height = _cd.Floor(_fbaea - (_ccgb - _eead.Height))
				_bfga, _gaeb := _fccba.split(_baff)
				if _bfga != nil {
					_faaab := _egafb()
					_faaab._feced = _eecc._feced
					_faaab._cecd = _bfga
					_afed = append(_afed, _faaab)
				}
				if _gaeb != nil {
					_bcce := _fccba._feeb.Style.FontSize
					_fgge, _cdef := _fccba._feeb.Style.Font.GetRuneMetrics(' ')
					if _cdef {
						_bcce = _fccba._feeb.Style.FontSize * _fgge.Wx * _fccba._feeb.Style.horizontalScale() / 1000.0
					}
					_agff := _a.Repeat("\u0020", int(_ebeg/_bcce))
					_gede := _egafb()
					_gede._feced = *NewTextChunk(_agff, _fccba._feeb.Style)
					_gede._cecd = _gaeb
					_gcee = append(_gcee, _gede)
					_gcee = append(_gcee, _ade._bdccg[_addb+1:]...)
				}
			default:
				_gcee = _ade._bdccg[_addb:]
			}
			if len(_gcee) > 0 {
				break
			}
		}
	}
	if len(_afed) > 0 {
		_gadb = _gdcb(_ade._edda)
		*_gadb = *_ade
		_gadb._bdccg = _afed
	}
	if len(_gcee) > 0 {
		_dafb = _gdcb(_ade._edda)
		*_dafb = *_ade
		_dafb._bdccg = _gcee
	}
	return _gadb, _dafb
}

// SetMargins sets the Chapter margins: left, right, top, bottom.
// Typically not needed as the creator's page margins are used.
func (_ddbd *Chapter) SetMargins(left, right, top, bottom float64) {
	_ddbd._ecd.Left = left
	_ddbd._ecd.Right = right
	_ddbd._ecd.Top = top
	_ddbd._ecd.Bottom = bottom
}

// ToPdfShadingPattern generates a new model.PdfShadingPatternType2 object.
func (_aeacf *LinearShading) ToPdfShadingPattern() *_ga.PdfShadingPatternType2 {
	_eddac, _edbb, _cgebc := _aeacf._cfbfe._cfec.ToRGB()
	_gadbg := _aeacf.shadingModel()
	_gadbg.PdfShading.Background = _cc.MakeArrayFromFloats([]float64{_eddac, _edbb, _cgebc})
	_aegcd := _ga.NewPdfShadingPatternType2()
	_aegcd.Shading = _gadbg
	return _aegcd
}

// GeneratePageBlocks draws the composite Bezier curve on a new block
// representing the page. Implements the Drawable interface.
func (_gbac *PolyBezierCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_deae := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_fbege, _dagdf := _deae.setOpacity(_gbac._gbbfd, _gbac._dagd)
	if _dagdf != nil {
		return nil, ctx, _dagdf
	}
	_bdbe := _gbac._ecgf
	_bdbe.FillEnabled = _bdbe.FillColor != nil
	var (
		_gabda = ctx.PageHeight
		_ecce  = _bdbe.Curves
		_gecf  = make([]_ff.CubicBezierCurve, 0, len(_bdbe.Curves))
	)
	_ggbd := _ga.PdfRectangle{}
	for _fbad := range _bdbe.Curves {
		_eadd := _ecce[_fbad]
		_eadd.P0.Y = _gabda - _eadd.P0.Y
		_eadd.P1.Y = _gabda - _eadd.P1.Y
		_eadd.P2.Y = _gabda - _eadd.P2.Y
		_eadd.P3.Y = _gabda - _eadd.P3.Y
		_gecf = append(_gecf, _eadd)
		_dgcb := _eadd.GetBounds()
		if _fbad == 0 {
			_ggbd = _dgcb
		} else {
			_ggbd.Llx = _cd.Min(_ggbd.Llx, _dgcb.Llx)
			_ggbd.Lly = _cd.Min(_ggbd.Lly, _dgcb.Lly)
			_ggbd.Urx = _cd.Max(_ggbd.Urx, _dgcb.Urx)
			_ggbd.Ury = _cd.Max(_ggbd.Ury, _dgcb.Ury)
		}
	}
	_bdbe.Curves = _gecf
	defer func() { _bdbe.Curves = _ecce }()
	if _bdbe.FillEnabled {
		_ddfgg := _acdb(_deae, _gbac._ecgf.FillColor, _gbac._ggag, func() Rectangle {
			return Rectangle{_fcgf: _ggbd.Llx, _cbaa: _ggbd.Lly, _bcffg: _ggbd.Width(), _ebegg: _ggbd.Height()}
		})
		if _ddfgg != nil {
			return nil, ctx, _ddfgg
		}
	}
	_cffgg, _, _dagdf := _bdbe.Draw(_fbege)
	if _dagdf != nil {
		return nil, ctx, _dagdf
	}
	if _dagdf = _deae.addContentsByString(string(_cffgg)); _dagdf != nil {
		return nil, ctx, _dagdf
	}
	return []*Block{_deae}, ctx, nil
}

func _bafebe(_beee *templateProcessor, _cacab *templateNode) (interface{}, error) {
	return _beee.parseLine(_cacab)
}

// FillOpacity returns the fill opacity of the rectangle (0-1).
func (_bgfd *Rectangle) FillOpacity() float64 { return _bgfd._gfdd }

// GetMargins returns the margins of the ellipse: left, right, top, bottom.
func (_bedg *Ellipse) GetMargins() (float64, float64, float64, float64) {
	return _bedg._abgc.Left, _bedg._abgc.Right, _bedg._abgc.Top, _bedg._abgc.Bottom
}

// SetPageMargins sets the page margins: left, right, top, bottom.
// The default page margins are 10% of document width.
func (_fgdb *Creator) SetPageMargins(left, right, top, bottom float64) {
	_fgdb._eaf.Left = left
	_fgdb._eaf.Right = right
	_fgdb._eaf.Top = top
	_fgdb._eaf.Bottom = bottom
}

func (_dedfd *templateProcessor) loadImageFromSrc(_ggbcfd string) (*Image, error) {
	if _ggbcfd == "" {
		_bcd.Log.Error("\u0049\u006d\u0061\u0067\u0065\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _efbef
	}
	_eeef := _a.Split(_ggbcfd, "\u002c")
	for _, _gabe := range _eeef {
		_gabe = _a.TrimSpace(_gabe)
		if _gabe == "" {
			continue
		}
		_bfegf, _bgdab := _dedfd._agcb.ImageMap[_gabe]
		if _bgdab {
			return _cfcc(_bfegf)
		}
		if _bfdgc := _dedfd.parseAttrPropList(_gabe); len(_bfdgc) > 0 {
			if _bgaba, _fdccf := _bfdgc["\u0070\u0061\u0074\u0068"]; _fdccf {
				if _cdaa, _agef := _caaf(_bgaba); _agef != nil {
					_bcd.Log.Debug("\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0060\u0025\u0073\u0060\u003a\u0020\u0025\u0076\u002e", _bgaba, _agef)
				} else {
					return _cdaa, nil
				}
			}
		}
	}
	_bcd.Log.Error("\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0069ma\u0067\u0065\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _ggbcfd)
	return nil, _efbef
}

type componentRenderer interface {
	Draw(_begf Drawable) error
}

func (_dedf *templateProcessor) getNodeErrorLocation(_gaegd *templateNode, _ffaee string, _cace ...interface{}) string {
	_deeaf := _g.Sprintf(_ffaee, _cace...)
	_gaaa := _g.Sprintf("\u0025\u0064", _gaegd._cgcf)
	if _gaegd._ffgg != 0 {
		_gaaa = _g.Sprintf("\u0025\u0064\u003a%\u0064", _gaegd._ffgg, _gaegd._gcdae)
	}
	if _dedf._bfgd != "" {
		return _g.Sprintf("\u0025\u0073\u0020\u005b\u0025\u0073\u003a\u0025\u0073\u005d", _deeaf, _dedf._bfgd, _gaaa)
	}
	return _g.Sprintf("\u0025s\u0020\u005b\u0025\u0073\u005d", _deeaf, _gaaa)
}

// Rows returns the total number of rows the table has.
func (_gaedf *Table) Rows() int { return _gaedf._dbbf }

func _aecbg(_geda, _acccc, _ecabbb float64) (_bfbga, _daeaf, _beffde, _cbegd float64) {
	if _ecabbb == 0 {
		return 0, 0, _geda, _acccc
	}
	_ggfeg := _ff.Path{Points: []_ff.Point{_ff.NewPoint(0, 0).Rotate(_ecabbb), _ff.NewPoint(_geda, 0).Rotate(_ecabbb), _ff.NewPoint(0, _acccc).Rotate(_ecabbb), _ff.NewPoint(_geda, _acccc).Rotate(_ecabbb)}}.GetBoundingBox()
	return _ggfeg.X, _ggfeg.Y, _ggfeg.Width, _ggfeg.Height
}

// SetBorderWidth sets the border width.
func (_fcca *CurvePolygon) SetBorderWidth(borderWidth float64) {
	_fcca._efbeg.BorderWidth = borderWidth
}

// EnableRowWrap controls whether rows are wrapped across pages.
// NOTE: Currently, row wrapping is supported for rows using StyledParagraphs.
func (_fecbe *Table) EnableRowWrap(enable bool) { _fecbe._ggea = enable }

// GetRowHeight returns the height of the specified row.
func (_dcdcb *Table) GetRowHeight(row int) (float64, error) {
	if row < 1 || row > len(_dcdcb._ecab) {
		return 0, _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	return _dcdcb._ecab[row-1], nil
}

// SkipRows skips over a specified number of rows in the table.
func (_gbfg *Table) SkipRows(num int) {
	_bgeda := num*_gbfg._edfe - 1
	if _bgeda < 0 {
		_bcd.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _cgdg := 0; _cgdg < _bgeda; _cgdg++ {
		_gbfg.NewCell()
	}
}

const (
	TextAlignmentLeft TextAlignment = iota
	TextAlignmentRight
	TextAlignmentCenter
	TextAlignmentJustify
)

// SetMargins sets the margins TOC line.
func (_cbdg *TOCLine) SetMargins(left, right, top, bottom float64) {
	_cbdg._dbdgg = left
	_aegefg := &_cbdg._eedce._fadcg
	_aegefg.Left = _cbdg._dbdgg + float64(_cbdg._dffba-1)*_cbdg._dgefgc
	_aegefg.Right = right
	_aegefg.Top = top
	_aegefg.Bottom = bottom
}

// AddColorStop add color stop info for rendering gradient color.
func (_beba *RadialShading) AddColorStop(color Color, point float64) {
	_beba._acdae.AddColorStop(color, point)
}

func (_bfde *templateProcessor) parseLinearGradientAttr(creator *Creator, _geagf string) Color {
	_fgebf := ColorBlack
	if _geagf == "" {
		return _fgebf
	}
	_fbdd := creator.NewLinearGradientColor([]*ColorPoint{})
	_fbdd.SetExtends(true, true)
	var (
		_bccee = _a.Split(_geagf[16:len(_geagf)-1], "\u002c")
		_cbabf = _a.TrimSpace(_bccee[0])
	)
	if _a.HasSuffix(_cbabf, "\u0064\u0065\u0067") {
		_beacf, _ebcgb := _aa.ParseFloat(_cbabf[:len(_cbabf)-3], 64)
		if _ebcgb != nil {
			_bcd.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0067\u0072\u0061\u0064\u0069e\u006e\u0074\u0020\u0061\u006e\u0067\u006ce\u003a\u0020\u0025\u0076", _ebcgb)
		} else {
			_fbdd.SetAngle(_beacf)
		}
		_bccee = _bccee[1:]
	}
	_fcebe, _agcg := _bfde.processGradientColorPair(_bccee)
	if _fcebe == nil || _agcg == nil {
		return _fgebf
	}
	for _dfgf := 0; _dfgf < len(_fcebe); _dfgf++ {
		_fbdd.AddColorStop(_fcebe[_dfgf], _agcg[_dfgf])
	}
	return _fbdd
}

// NewCurvePolygon creates a new curve polygon.
func (_agbg *Creator) NewCurvePolygon(rings [][]_ff.CubicBezierCurve) *CurvePolygon {
	return _acce(rings)
}

var ErrContentNotFit = _c.New("\u0043\u0061\u006e\u006e\u006ft\u0020\u0066\u0069\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020i\u006e\u0074\u006f\u0020\u0061\u006e\u0020\u0065\u0078\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0073\u0070\u0061\u0063\u0065")

// SetLineWidth sets the line width.
func (_gdea *Line) SetLineWidth(width float64) { _gdea._bfe = width }

// PageFinalize sets a function to be called for each page before finalization
// (i.e. the last stage of page processing before they get written out).
// The callback function allows final touch-ups for each page, and it
// provides information that might not be known at other stages of designing
// the document (e.g. the total number of pages). Unlike the header/footer
// functions, which are limited to the top/bottom margins of the page, the
// finalize function can be used draw components anywhere on the current page.
func (_cfcaa *Creator) PageFinalize(pageFinalizeFunc func(_bbad PageFinalizeFunctionArgs) error) {
	_cfcaa._aagf = pageFinalizeFunc
}

// SetBorderColor sets the border color.
func (_cgce *Polygon) SetBorderColor(color Color) { _cgce._acaag.BorderColor = _fce(color) }

func _gadcc(_bgggb string, _cagg bool) string {
	_ggff := _bgggb
	if _ggff == "" {
		return ""
	}
	_eaaae := _gf.Paragraph{}
	_, _eadbd := _eaaae.SetString(_bgggb)
	if _eadbd != nil {
		return _ggff
	}
	_cfde, _eadbd := _eaaae.Order()
	if _eadbd != nil {
		return _ggff
	}
	_cggf := _cfde.NumRuns()
	_cfbbbb := make([]string, _cggf)
	for _dbeed := 0; _dbeed < _cfde.NumRuns(); _dbeed++ {
		_ebdbf := _cfde.Run(_dbeed)
		_ebgf := _ebdbf.String()
		if _ebdbf.Direction() == _gf.RightToLeft {
			_ebgf = _gf.ReverseString(_ebgf)
		}
		if _cagg {
			_cfbbbb[_dbeed] = _ebgf
		} else {
			_cfbbbb[_cggf-1] = _ebgf
		}
		_cggf--
	}
	if len(_cfbbbb) != _cfde.NumRuns() {
		return _bgggb
	}
	_ggff = _a.Join(_cfbbbb, "")
	return _ggff
}

// Chart represents a chart drawable.
// It is used to render unichart chart components using a creator instance.
type Chart struct {
	_dece _dce.ChartRenderable
	_bdcb Positioning
	_cdd  float64
	_dcd  float64
	_baae Margins
}

// SetLinePageStyle sets the style for the page part of all new lines
// of the table of contents.
func (_bfdbf *TOC) SetLinePageStyle(style TextStyle) { _bfdbf._afccg = style }

// BorderColor returns the border color of the ellipse.
func (_cbdca *Ellipse) BorderColor() Color { return _cbdca._bacf }

// Height returns the total height of all rows.
func (_bfcbg *Table) Height() float64 {
	_ccff := float64(0.0)
	for _, _decbdf := range _bfcbg._ecab {
		_ccff += _decbdf
	}
	return _ccff
}

// SetExtends specifies whether to extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_eggfb *shading) SetExtends(start bool, end bool) { _eggfb._fedf = []bool{start, end} }

// NewStyledParagraph creates a new styled paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_cbgac *Creator) NewStyledParagraph() *StyledParagraph { return _cgfa(_cbgac.NewTextStyle()) }

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_ecdbg *TOCLine) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ggbff := ctx
	_gfaaf, ctx, _cbedg := _ecdbg._eedce.GeneratePageBlocks(ctx)
	if _cbedg != nil {
		return _gfaaf, ctx, _cbedg
	}
	if _ecdbg._cdfbe.IsRelative() {
		ctx.X = _ggbff.X
	}
	if _ecdbg._cdfbe.IsAbsolute() {
		return _gfaaf, _ggbff, nil
	}
	return _gfaaf, ctx, nil
}

// GetMargins returns the Block's margins: left, right, top, bottom.
func (_gbg *Block) GetMargins() (float64, float64, float64, float64) {
	return _gbg._bf.Left, _gbg._bf.Right, _gbg._bf.Top, _gbg._bf.Bottom
}

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_beebe *RadialShading) SetAntiAlias(enable bool) { _beebe._acdae.SetAntiAlias(enable) }

func (_cdcf *Table) updateRowHeights(_cebab float64) {
	for _, _fbcd := range _cdcf._gebec {
		_cfbbb := _fbcd.width(_cdcf._eeggd, _cebab)
		_gabdd := _fbcd.height(_cfbbb)
		_eecca := _cdcf._ecab[_fbcd._bgdca+_fbcd._fddc-2]
		if _fbcd._fddc > 1 {
			_baab := 0.0
			_beead := _cdcf._ecab[_fbcd._bgdca-1 : (_fbcd._bgdca + _fbcd._fddc - 1)]
			for _, _fgcd := range _beead {
				_baab += _fgcd
			}
			if _gabdd <= _baab {
				continue
			}
		}
		if _gabdd > _eecca {
			_cgef := _gabdd / float64(_fbcd._fddc)
			if _cgef > _eecca {
				for _adcca := 1; _adcca <= _fbcd._fddc; _adcca++ {
					if _cgef > _cdcf._ecab[_fbcd._bgdca+_adcca-2] {
						_cdcf._ecab[_fbcd._bgdca+_adcca-2] = _cgef
					}
				}
			}
		}
	}
}

// ScaleToHeight scale Image to a specified height h, maintaining the aspect ratio.
func (_dced *Image) ScaleToHeight(h float64) {
	_dcedd := _dced._ggda / _dced._cgdb
	_dced._cgdb = h
	_dced._ggda = h * _dcedd
}

// AddPage adds the specified page to the creator.
// NOTE: If the page has a Rotate flag, the creator will take care of
// transforming the contents to maintain the correct orientation.
func (_adaec *Creator) AddPage(page *_ga.PdfPage) error {
	_agc, _abgd := _adaec.wrapPageIfNeeded(page)
	if _abgd != nil {
		return _abgd
	}
	if _agc != nil {
		page = _agc
	}
	_abab, _abgd := page.GetMediaBox()
	if _abgd != nil {
		_bcd.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0070\u0061g\u0065\u0020\u006d\u0065\u0064\u0069\u0061\u0062\u006f\u0078\u003a\u0020\u0025\u0076", _abgd)
		return _abgd
	}
	_abab.Normalize()
	_dabd, _ceea := _abab.Llx, _abab.Lly
	_feg := _abab
	if _afde := page.CropBox; _afde != nil && *_afde != *_abab {
		_afde.Normalize()
		_dabd, _ceea = _afde.Llx, _afde.Lly
		_feg = _afde
	}
	_beff := _gb.IdentityMatrix()
	_gfec, _abgd := page.GetRotate()
	if _abgd != nil {
		_bcd.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _abgd.Error())
	}
	_dffc := _gfec%360 != 0 && _gfec%90 == 0
	if _dffc {
		_ffe := float64((360 + _gfec%360) % 360)
		if _ffe == 90 {
			_beff = _beff.Translate(_feg.Width(), 0)
		} else if _ffe == 180 {
			_beff = _beff.Translate(_feg.Width(), _feg.Height())
		} else if _ffe == 270 {
			_beff = _beff.Translate(0, _feg.Height())
		}
		_beff = _beff.Mult(_gb.RotationMatrix(_ffe * _cd.Pi / 180))
		_beff = _beff.Round(0.000001)
		_egaf := _gfaf(_feg, _beff)
		_feg = _egaf
		_feg.Normalize()
	}
	if _dabd != 0 || _ceea != 0 {
		_beff = _gb.TranslationMatrix(_dabd, _ceea).Mult(_beff)
	}
	if !_beff.Identity() {
		_beff = _beff.Round(0.000001)
		_adaec._bbf[page] = &pageTransformations{_aefc: &_beff}
	}
	_adaec._bagf = _feg.Width()
	_adaec._aafc = _feg.Height()
	_adaec.initContext()
	_adaec._agfg = append(_adaec._agfg, page)
	_adaec._ggab.Page++
	return nil
}

func (_cgfg *TOCLine) getLineLink() *_ga.PdfAnnotation {
	if _cgfg._dfbfg <= 0 {
		return nil
	}
	return _bgfbc(_cgfg._dfbfg-1, _cgfg._dfbbb, _cgfg._gcab, 0)
}

// FitMode returns the fit mode of the rectangle.
func (_ddff *Rectangle) FitMode() FitMode { return _ddff._efbb }

func _adgc(_fegca *templateProcessor, _bbabb *templateNode) (interface{}, error) {
	return _fegca.parseListMarker(_bbabb)
}

// Width returns the width of the Paragraph.
func (_afaea *StyledParagraph) Width() float64 {
	if _afaea._bfcgb && int(_afaea._cccde) > 0 {
		return _afaea._cccde
	}
	return _afaea.getTextWidth() / 1000.0
}

// NewImage create a new image from a unidoc image (model.Image).
func (_ffcc *Creator) NewImage(img *_ga.Image) (*Image, error) { return _cfcc(img) }

// GeneratePageBlocks generate the Page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages.
func (_cfg *Chapter) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fcd := ctx
	if _cfg._fdf.IsRelative() {
		ctx.X += _cfg._ecd.Left
		ctx.Y += _cfg._ecd.Top
		ctx.Width -= _cfg._ecd.Left + _cfg._ecd.Right
		ctx.Height -= _cfg._ecd.Top
	}
	_bbg, _degc, _ddfaa := _cfg._bgfc.GeneratePageBlocks(ctx)
	if _ddfaa != nil {
		return _bbg, ctx, _ddfaa
	}
	ctx = _degc
	_gdcf := ctx.X
	_afab := ctx.Y - _cfg._bgfc.Height()
	_agd := int64(ctx.Page)
	_aeae := _cfg.headingNumber()
	_gdaa := _cfg.headingText()
	if _cfg._gge {
		_cegd := _cfg._aaccb.Add(_aeae, _cfg._gbb, _aa.FormatInt(_agd, 10), _cfg._fge)
		if _cfg._aaccb._bbee {
			_cegd.SetLink(_agd, _gdcf, _afab)
		}
	}
	if _cfg._fddf == nil {
		_cfg._fddf = _ga.NewOutlineItem(_gdaa, _ga.NewOutlineDest(_agd-1, _gdcf, _afab))
		if _cfg._eagg != nil {
			_cfg._eagg._fddf.Add(_cfg._fddf)
		} else {
			_cfg._eebd.Add(_cfg._fddf)
		}
	} else {
		_ebf := &_cfg._fddf.Dest
		_ebf.Page = _agd - 1
		_ebf.X = _gdcf
		_ebf.Y = _afab
	}
	for _, _dddf := range _cfg._gefd {
		_feab, _dcc, _fee := _dddf.GeneratePageBlocks(ctx)
		if _fee != nil {
			return _bbg, ctx, _fee
		}
		if len(_feab) < 1 {
			continue
		}
		_bbg[len(_bbg)-1].mergeBlocks(_feab[0])
		_bbg = append(_bbg, _feab[1:]...)
		ctx = _dcc
	}
	if _cfg._fdf.IsRelative() {
		ctx.X = _fcd.X
	}
	if _cfg._fdf.IsAbsolute() {
		return _bbg, _fcd, nil
	}
	return _bbg, ctx, nil
}

// SetWidth sets the width of the ellipse.
func (_eced *Ellipse) SetWidth(width float64) { _eced._eded = width }

// NewColumn returns a new column for the line items invoice table.
func (_cacb *Invoice) NewColumn(description string) *InvoiceCell {
	return _cacb.newColumn(description, CellHorizontalAlignmentLeft)
}

// Title returns the title of the invoice.
func (_eagc *Invoice) Title() string { return _eagc._efdb }

// New creates a new instance of the PDF Creator.
func New() *Creator {
	const _aaa = "c\u0072\u0065\u0061\u0074\u006f\u0072\u002e\u004e\u0065\u0077"
	_eaff := &Creator{}
	_eaff._agfg = []*_ga.PdfPage{}
	_eaff._fgd = map[*_ga.PdfPage]*Block{}
	_eaff._bbf = map[*_ga.PdfPage]*pageTransformations{}
	_eaff.SetPageSize(PageSizeLetter)
	_abge := 0.1 * _eaff._bagf
	_eaff._eaf.Left = _abge
	_eaff._eaf.Right = _abge
	_eaff._eaf.Top = _abge
	_eaff._eaf.Bottom = _abge
	var _bded error
	_eaff._ffcg, _bded = _ga.NewStandard14Font(_ga.HelveticaName)
	if _bded != nil {
		_eaff._ffcg = _ga.DefaultFont()
	}
	_eaff._cbag, _bded = _ga.NewStandard14Font(_ga.HelveticaBoldName)
	if _bded != nil {
		_eaff._ffcg = _ga.DefaultFont()
	}
	_eaff._edad = _eaff.NewTOC("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073")
	_eaff.AddOutlines = true
	_eaff._bcb = _ga.NewOutline()
	return _eaff
}

// GeneratePageBlocks draws the composite curve polygon on a new block
// representing the page. Implements the Drawable interface.
func (_dae *CurvePolygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fece := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_ddfb, _feee := _fece.setOpacity(_dae._gfbf, _dae._adad)
	if _feee != nil {
		return nil, ctx, _feee
	}
	_dfca := _dae._efbeg
	_dfca.FillEnabled = _dfca.FillColor != nil
	_dfca.BorderEnabled = _dfca.BorderColor != nil && _dfca.BorderWidth > 0
	var (
		_deed = ctx.PageHeight
		_fdaa = _dfca.Rings
		_daee = make([][]_ff.CubicBezierCurve, 0, len(_dfca.Rings))
	)
	_cgeb := _ga.PdfRectangle{}
	if len(_fdaa) > 0 && len(_fdaa[0]) > 0 {
		_adcd := _fdaa[0][0]
		_adcd.P0.Y = _deed - _adcd.P0.Y
		_adcd.P1.Y = _deed - _adcd.P1.Y
		_adcd.P2.Y = _deed - _adcd.P2.Y
		_adcd.P3.Y = _deed - _adcd.P3.Y
		_cgeb = _adcd.GetBounds()
	}
	for _, _fbb := range _fdaa {
		_fdc := make([]_ff.CubicBezierCurve, 0, len(_fbb))
		for _, _bgaa := range _fbb {
			_adaa := _bgaa
			_adaa.P0.Y = _deed - _adaa.P0.Y
			_adaa.P1.Y = _deed - _adaa.P1.Y
			_adaa.P2.Y = _deed - _adaa.P2.Y
			_adaa.P3.Y = _deed - _adaa.P3.Y
			_fdc = append(_fdc, _adaa)
			_fffe := _adaa.GetBounds()
			_cgeb.Llx = _cd.Min(_cgeb.Llx, _fffe.Llx)
			_cgeb.Lly = _cd.Min(_cgeb.Lly, _fffe.Lly)
			_cgeb.Urx = _cd.Max(_cgeb.Urx, _fffe.Urx)
			_cgeb.Ury = _cd.Max(_cgeb.Ury, _fffe.Ury)
		}
		_daee = append(_daee, _fdc)
	}
	_dfca.Rings = _daee
	defer func() { _dfca.Rings = _fdaa }()
	if _dfca.FillEnabled {
		_bdgge := _acdb(_fece, _dae._efbeg.FillColor, _dae._aged, func() Rectangle {
			return Rectangle{_fcgf: _cgeb.Llx, _cbaa: _cgeb.Lly, _bcffg: _cgeb.Width(), _ebegg: _cgeb.Height()}
		})
		if _bdgge != nil {
			return nil, ctx, _bdgge
		}
	}
	_bdf, _, _feee := _dfca.Draw(_ddfb)
	if _feee != nil {
		return nil, ctx, _feee
	}
	if _feee = _fece.addContentsByString(string(_bdf)); _feee != nil {
		return nil, ctx, _feee
	}
	return []*Block{_fece}, ctx, nil
}

// NewCurve returns new instance of Curve between points (x1,y1) and (x2, y2) with control point (cx,cy).
func (_bdcbf *Creator) NewCurve(x1, y1, cx, cy, x2, y2 float64) *Curve {
	return _dcabd(x1, y1, cx, cy, x2, y2)
}

// ColorCMYKFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//
//	green := ColorCMYKFromArithmetic(1.0, 0.0, 1.0, 0.0)
func ColorCMYKFromArithmetic(c, m, y, k float64) Color {
	return cmykColor{_bdge: _cd.Max(_cd.Min(c, 1.0), 0.0), _eegd: _cd.Max(_cd.Min(m, 1.0), 0.0), _ebe: _cd.Max(_cd.Min(y, 1.0), 0.0), _agg: _cd.Max(_cd.Min(k, 1.0), 0.0)}
}

type containerDrawable interface {
	Drawable

	// ContainerComponent checks if the component is allowed to be added into provided 'container' and returns
	// preprocessed copy of itself. If the component is not changed it is allowed to return itself in a callback way.
	// If the component is not compatible with provided container this method should return an error.
	ContainerComponent(_egcf Drawable) (Drawable, error)
}

// NewInvoice returns an instance of an empty invoice.
func (_aggg *Creator) NewInvoice() *Invoice {
	_abae := _aggg.NewTextStyle()
	_abae.Font = _aggg._cbag
	return _bbb(_aggg.NewTextStyle(), _abae)
}

// AddAnnotation adds an annotation to the current block.
// The annotation will be added to the page the block will be rendered on.
func (_dfg *Block) AddAnnotation(annotation *_ga.PdfAnnotation) {
	for _, _gca := range _dfg._fdd {
		if _gca == annotation {
			return
		}
	}
	_dfg._fdd = append(_dfg._fdd, annotation)
}

// Color returns the color of the line.
func (_afbe *Line) Color() Color { return _afbe._efdf }

// SetOpacity sets opacity for Image.
func (_egbe *Image) SetOpacity(opacity float64) { _egbe._gafe = opacity }

func _caceg(_dbcbde *templateProcessor, _fdgaf *templateNode) (interface{}, error) {
	return _dbcbde.parseChapterHeading(_fdgaf)
}

func _acce(_efae [][]_ff.CubicBezierCurve) *CurvePolygon {
	return &CurvePolygon{_efbeg: &_ff.CurvePolygon{Rings: _efae}, _gfbf: 1.0, _adad: 1.0}
}

// CellBorderStyle defines the table cell's border style.
type CellBorderStyle int

func (_fdgc *templateProcessor) parseFloatAttr(_fcaeb, _aede string) float64 {
	_bcd.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0066\u006c\u006f\u0061t a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _fcaeb, _aede)
	_cabba, _ := _aa.ParseFloat(_aede, 64)
	return _cabba
}

// GetMargins returns the margins of the TOC line: left, right, top, bottom.
func (_fdgbd *TOCLine) GetMargins() (float64, float64, float64, float64) {
	_bgee := &_fdgbd._eedce._fadcg
	return _fdgbd._dbdgg, _bgee.Right, _bgee.Top, _bgee.Bottom
}

// FitMode returns the fit mode of the line.
func (_dfbg *Line) FitMode() FitMode { return _dfbg._cbab }

// AddColorStop add color stop information for rendering gradient.
func (_gfcec *shading) AddColorStop(color Color, point float64) {
	_gfcec._ccaf = append(_gfcec._ccaf, _gebe(color, point))
}

// DrawFooter sets a function to draw a footer on created output pages.
func (_faaa *Creator) DrawFooter(drawFooterFunc func(_adcg *Block, _egc FooterFunctionArgs)) {
	_faaa._acfe = drawFooterFunc
}
func _dbae() *PageBreak { return &PageBreak{} }
func (_ceafa *StyledParagraph) getTextLineWidth(_aggda []*TextChunk) float64 {
	var _gdede float64
	_bgcf := len(_aggda)
	for _fdgb, _agabc := range _aggda {
		_agce := &_agabc.Style
		_ccac := len(_agabc.Text)
		for _fbac, _bdef := range _agabc.Text {
			if _bdef == '\u000A' {
				continue
			}
			_afdag, _defa := _agce.Font.GetRuneMetrics(_bdef)
			if !_defa {
				_bcd.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _bdef)
				return -1
			}
			_gdede += _agce.FontSize * _afdag.Wx * _agce.horizontalScale()
			if _bdef != ' ' && (_fdgb != _bgcf-1 || _fbac != _ccac-1) {
				_gdede += _agce.CharSpacing * 1000.0
			}
		}
	}
	return _gdede
}

func (_cdccf *templateProcessor) parseLineStyleAttr(_gced, _aaaag string) _ff.LineStyle {
	_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _gced, _aaaag)
	_ggcg := map[string]_ff.LineStyle{"\u0073\u006f\u006ci\u0064": _ff.LineStyleSolid, "\u0064\u0061\u0073\u0068\u0065\u0064": _ff.LineStyleDashed}[_aaaag]
	return _ggcg
}

// TextOverflow determines the behavior of paragraph text which does
// not fit in the available space.
type TextOverflow int

func (_cga rgbColor) ToRGB() (float64, float64, float64) { return _cga._ddae, _cga._cede, _cga._fgef }

// ScaleToWidth scales the rectangle to the specified width. The height of
// the rectangle is scaled so that the aspect ratio is maintained.
func (_fdde *Rectangle) ScaleToWidth(w float64) {
	_eefee := _fdde._ebegg / _fdde._bcffg
	_fdde._bcffg = w
	_fdde._ebegg = w * _eefee
}

// Scale block by specified factors in the x and y directions.
func (_fdbd *Block) Scale(sx, sy float64) {
	_adf := _da.NewContentCreator().Scale(sx, sy).Operations()
	*_fdbd._fa = append(*_adf, *_fdbd._fa...)
	_fdbd._fa.WrapIfNeeded()
	_fdbd._fd *= sx
	_fdbd._cce *= sy
}

func (_debfd *Paragraph) wrapText() error {
	if !_debfd._dbfb || int(_debfd._adcfe) <= 0 {
		_debfd._aggb = []string{_debfd._bfgg}
		return nil
	}
	_bgdc := NewTextChunk(_debfd._bfgg, TextStyle{Font: _debfd._bfbd, FontSize: _debfd._dbad})
	_ffeefg, _ecaf := _bgdc.Wrap(_debfd._adcfe)
	if _ecaf != nil {
		return _ecaf
	}
	if _debfd._gfbef > 0 && len(_ffeefg) > _debfd._gfbef {
		_ffeefg = _ffeefg[:_debfd._gfbef]
	}
	_debfd._aggb = _ffeefg
	return nil
}

// AnchorPoint defines anchor point where the center position of the radial gradient would be calculated.
type AnchorPoint int

func _ecbaa(_bcbea, _acgbb, _ffdec TextChunk, _cbeeb uint, _gbecc TextStyle) *TOCLine {
	_gdcfba := _cgfa(_gbecc)
	_gdcfba.SetEnableWrap(true)
	_gdcfba.SetTextAlignment(TextAlignmentLeft)
	_gdcfba.SetMargins(0, 0, 2, 2)
	_baeb := &TOCLine{_eedce: _gdcfba, Number: _bcbea, Title: _acgbb, Page: _ffdec, Separator: TextChunk{Text: "\u002e", Style: _gbecc}, _dbdgg: 0, _dffba: _cbeeb, _dgefgc: 10, _cdfbe: PositionRelative}
	_gdcfba._fadcg.Left = _baeb._dbdgg + float64(_baeb._dffba-1)*_baeb._dgefgc
	_gdcfba._egdcb = _baeb.prepareParagraph
	return _baeb
}

// Height returns the current page height.
func (_edef *Creator) Height() float64 { return _edef._aafc }

// SetFont sets the Paragraph's font.
func (_gfbb *Paragraph) SetFont(font *_ga.PdfFont) { _gfbb._bfbd = font }

// RotateDeg rotates the current active page by angle degrees.  An error is returned on failure,
// which can be if there is no currently active page, or the angleDeg is not a multiple of 90 degrees.
func (_ggac *Creator) RotateDeg(angleDeg int64) error {
	_ffgd := _ggac.getActivePage()
	if _ffgd == nil {
		_bcd.Log.Debug("F\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0065\u003a\u0020\u006e\u006f\u0020p\u0061\u0067\u0065\u0020\u0063\u0075\u0072\u0072\u0065\u006etl\u0079\u0020\u0061c\u0074i\u0076\u0065")
		return _c.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	if angleDeg%90 != 0 {
		_bcd.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067e\u0020\u0072\u006f\u0074\u0061\u0074\u0069on\u0020\u0061\u006e\u0067l\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006dul\u0074\u0069p\u006c\u0065\u0020\u006f\u0066\u0020\u0039\u0030")
		return _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	var _dfged int64
	if _ffgd.Rotate != nil {
		_dfged = *(_ffgd.Rotate)
	}
	_dfged += angleDeg
	_ffgd.Rotate = &_dfged
	return nil
}

// SetMargins sets the margins of the paragraph.
func (_ccec *List) SetMargins(left, right, top, bottom float64) {
	_ccec._fcag.Left = left
	_ccec._fcag.Right = right
	_ccec._fcag.Top = top
	_ccec._fcag.Bottom = bottom
}

func (_eeafa *Table) moveToNextAvailableCell() int {
	_aaaa := (_eeafa._ebad-1)%(_eeafa._edfe) + 1
	for {
		if _aaaa-1 >= len(_eeafa._bacd) {
			if _eeafa._bacd[0] == 0 {
				return _aaaa
			}
			_aaaa = 1
		} else if _eeafa._bacd[_aaaa-1] == 0 {
			return _aaaa
		}
		_eeafa._ebad++
		_eeafa._bacd[_aaaa-1]--
		_aaaa++
	}
}

func _cgf(_eceg string, _fbg _cc.PdfObject, _aed *_ga.PdfPageResources) _cc.PdfObjectName {
	_accc := _a.TrimRightFunc(_a.TrimSpace(_eceg), func(_fbgb rune) bool { return _bc.IsNumber(_fbgb) })
	if _accc == "" {
		_accc = "\u0046\u006f\u006e\u0074"
	}
	_ffg := 0
	_fbde := _cc.PdfObjectName(_eceg)
	for {
		_ebb, _fdg := _aed.GetFontByName(_fbde)
		if !_fdg || _ebb == _fbg {
			break
		}
		_ffg++
		_fbde = _cc.PdfObjectName(_g.Sprintf("\u0025\u0073\u0025\u0064", _accc, _ffg))
	}
	return _fbde
}

func (_fdgbf *templateProcessor) parseBackground(_cabc *templateNode) (interface{}, error) {
	_gbgac := &Background{}
	for _, _ccbg := range _cabc._acdf.Attr {
		_gafcg := _ccbg.Value
		switch _gbag := _ccbg.Name.Local; _gbag {
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_gbgac.FillColor = _fdgbf.parseColorAttr(_gbag, _gafcg)
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_gbgac.BorderColor = _fdgbf.parseColorAttr(_gbag, _gafcg)
		case "b\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0069\u007a\u0065":
			_gbgac.BorderSize = _fdgbf.parseFloatAttr(_gbag, _gafcg)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_facbe, _baegg, _ffcdc, _degba := _fdgbf.parseBorderRadiusAttr(_gbag, _gafcg)
			_gbgac.SetBorderRadius(_facbe, _baegg, _degba, _ffcdc)
		case "\u0062\u006f\u0072\u0064er\u002d\u0074\u006f\u0070\u002d\u006c\u0065\u0066\u0074\u002d\u0072\u0061\u0064\u0069u\u0073":
			_gbgac.BorderRadiusTopLeft = _fdgbf.parseFloatAttr(_gbag, _gafcg)
		case "\u0062\u006f\u0072de\u0072\u002d\u0074\u006f\u0070\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_gbgac.BorderRadiusTopRight = _fdgbf.parseFloatAttr(_gbag, _gafcg)
		case "\u0062o\u0072\u0064\u0065\u0072-\u0062\u006f\u0074\u0074\u006fm\u002dl\u0065f\u0074\u002d\u0072\u0061\u0064\u0069\u0075s":
			_gbgac.BorderRadiusBottomLeft = _fdgbf.parseFloatAttr(_gbag, _gafcg)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0062\u006f\u0074\u0074o\u006d\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061d\u0069\u0075\u0073":
			_gbgac.BorderRadiusBottomRight = _fdgbf.parseFloatAttr(_gbag, _gafcg)
		default:
			_fdgbf.nodeLogDebug(_cabc, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _gbag)
		}
	}
	return _gbgac, nil
}

// ColorRGBFromHex converts color hex code to rgb color for using with creator.
// NOTE: If there is a problem interpreting the string, then will use black color and log a debug message.
// Example hex code: #ffffff -> (1,1,1) white.
func ColorRGBFromHex(hexStr string) Color {
	_ddca := rgbColor{}
	if (len(hexStr) != 4 && len(hexStr) != 7) || hexStr[0] != '#' {
		_bcd.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
		return _ddca
	}
	var _gfab, _gbad, _gebb int
	if len(hexStr) == 4 {
		var _gaa, _gfcc, _eccd int
		_bada, _eff := _g.Sscanf(hexStr, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_gaa, &_gfcc, &_eccd)
		if _eff != nil {
			_bcd.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", hexStr, _eff)
			return _ddca
		}
		if _bada != 3 {
			_bcd.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _ddca
		}
		_gfab = _gaa*16 + _gaa
		_gbad = _gfcc*16 + _gfcc
		_gebb = _eccd*16 + _eccd
	} else {
		_ebfa, _ddaaa := _g.Sscanf(hexStr, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_gfab, &_gbad, &_gebb)
		if _ddaaa != nil {
			_bcd.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _ddca
		}
		if _ebfa != 3 {
			_bcd.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", hexStr, _ebfa)
			return _ddca
		}
	}
	_bdd := float64(_gfab) / 255.0
	_fdbdf := float64(_gbad) / 255.0
	_dfdg := float64(_gebb) / 255.0
	_ddca._ddae = _bdd
	_ddca._cede = _fdbdf
	_ddca._fgef = _dfdg
	return _ddca
}

// ScaleToHeight scales the ellipse to the specified height. The width of
// the ellipse is scaled so that the aspect ratio is maintained.
func (_ggf *Ellipse) ScaleToHeight(h float64) {
	_bfcg := _ggf._eded / _ggf._bdcce
	_ggf._bdcce = h
	_ggf._eded = h * _bfcg
}

// SetMargins sets the Block's left, right, top, bottom, margins.
func (_eab *Block) SetMargins(left, right, top, bottom float64) {
	_eab._bf.Left = left
	_eab._bf.Right = right
	_eab._bf.Top = top
	_eab._bf.Bottom = bottom
}

// SetBorderColor sets border color of the rectangle.
func (_fcgaa *Rectangle) SetBorderColor(col Color) { _fcgaa._gggb = col }

func (_ccda *templateProcessor) parseTable(_cbefa *templateNode) (interface{}, error) {
	var _gbbcd int64
	for _, _dgad := range _cbefa._acdf.Attr {
		_fecgf := _dgad.Value
		switch _edfb := _dgad.Name.Local; _edfb {
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
			_gbbcd = _ccda.parseInt64Attr(_edfb, _fecgf)
		}
	}
	if _gbbcd <= 0 {
		_ccda.nodeLogDebug(_cbefa, "\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006eu\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0074\u0061\u0062\u006ce\u0020\u0063\u006f\u006cu\u006d\u006e\u0073\u003a\u0020\u0025\u0064\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0031\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020m\u0061\u0079\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _gbbcd)
		_gbbcd = 1
	}
	_efbbe := _ccda.creator.NewTable(int(_gbbcd))
	for _, _bbde := range _cbefa._acdf.Attr {
		_eadbf := _bbde.Value
		switch _ddea := _bbde.Name.Local; _ddea {
		case "\u0063\u006f\u006c\u0075\u006d\u006e\u002d\u0077\u0069\u0064\u0074\u0068\u0073":
			_efbbe.SetColumnWidths(_ccda.parseFloatArray(_ddea, _eadbf)...)
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_abef := _ccda.parseMarginAttr(_ddea, _eadbf)
			_efbbe.SetMargins(_abef.Left, _abef.Right, _abef.Top, _abef.Bottom)
		case "\u0078":
			_efbbe.SetPos(_ccda.parseFloatAttr(_ddea, _eadbf), _efbbe._bfbac)
		case "\u0079":
			_efbbe.SetPos(_efbbe._aabe, _ccda.parseFloatAttr(_ddea, _eadbf))
		case "\u0068\u0065a\u0064\u0065\u0072-\u0073\u0074\u0061\u0072\u0074\u002d\u0072\u006f\u0077":
			_efbbe._gcdff = int(_ccda.parseInt64Attr(_ddea, _eadbf))
		case "\u0068\u0065\u0061\u0064\u0065\u0072\u002d\u0065\u006ed\u002d\u0072\u006f\u0077":
			_efbbe._bgec = int(_ccda.parseInt64Attr(_ddea, _eadbf))
		case "\u0065n\u0061b\u006c\u0065\u002d\u0072\u006f\u0077\u002d\u0077\u0072\u0061\u0070":
			_efbbe.EnableRowWrap(_ccda.parseBoolAttr(_ddea, _eadbf))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_efbbe.EnablePageWrap(_ccda.parseBoolAttr(_ddea, _eadbf))
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
			break
		default:
			_ccda.nodeLogDebug(_cbefa, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0074a\u0062\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _ddea)
		}
	}
	if _efbbe._gcdff != 0 && _efbbe._bgec != 0 {
		_bda := _efbbe.SetHeaderRows(_efbbe._gcdff, _efbbe._bgec)
		if _bda != nil {
			_ccda.nodeLogDebug(_cbefa, "\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0073\u0065t t\u0061bl\u0065\u0020\u0068\u0065\u0061\u0064\u0065r \u0072\u006f\u0077\u0073\u003a\u0020\u0025v\u002e", _bda)
		}
	} else {
		_efbbe._gcdff = 0
		_efbbe._bgec = 0
	}
	return _efbbe, nil
}

// Line defines a line between point 1 (X1, Y1) and point 2 (X2, Y2).
// The line width, color, style (solid or dashed) and opacity can be
// configured. Implements the Drawable interface.
type Line struct {
	_ccg   float64
	_eggf  float64
	_dfcae float64
	_feafb float64
	_efdf  Color
	_ffgbg _ff.LineStyle
	_gbcb  float64
	_fedd  []int64
	_ffgdf int64
	_bfe   float64
	_fcdd  Positioning
	_cbab  FitMode
	_gabb  Margins
}

// SetPos sets absolute positioning with specified coordinates.
func (_ecdd *StyledParagraph) SetPos(x, y float64) {
	_ecdd._bedd = PositionAbsolute
	_ecdd._fcadg = x
	_ecdd._ffdc = y
}

// SetHeading sets the text and the style of the heading of the TOC component.
func (_gedd *TOC) SetHeading(text string, style TextStyle) {
	_cbcee := _gedd.Heading()
	_cbcee.Reset()
	_cced := _cbcee.Append(text)
	_cced.Style = style
}

// HorizontalAlignment represents the horizontal alignment of components
// within a page.
type HorizontalAlignment int

// GetMargins returns the margins of the chart (left, right, top, bottom).
func (_cfca *Chart) GetMargins() (float64, float64, float64, float64) {
	return _cfca._baae.Left, _cfca._baae.Right, _cfca._baae.Top, _cfca._baae.Bottom
}

const (
	CellBorderStyleNone CellBorderStyle = iota
	CellBorderStyleSingle
	CellBorderStyleDouble
)

// SetShowNumbering sets a flag to indicate whether or not to show chapter numbers as part of title.
func (_afb *Chapter) SetShowNumbering(show bool) {
	_afb._ebgde = show
	_afb._bgfc.SetText(_afb.headingText())
}

// GetHorizontalAlignment returns the horizontal alignment of the image.
func (_cccb *Image) GetHorizontalAlignment() HorizontalAlignment { return _cccb._cccd }

func (_eegg cmykColor) ToRGB() (float64, float64, float64) {
	_feaf := _eegg._agg
	return 1 - (_eegg._bdge*(1-_feaf) + _feaf), 1 - (_eegg._eegd*(1-_feaf) + _feaf), 1 - (_eegg._ebe*(1-_feaf) + _feaf)
}

func (_beeg *templateProcessor) parseEllipse(_afbec *templateNode) (interface{}, error) {
	_geced := _beeg.creator.NewEllipse(0, 0, 0, 0)
	for _, _edabb := range _afbec._acdf.Attr {
		_ddgc := _edabb.Value
		switch _cgedc := _edabb.Name.Local; _cgedc {
		case "\u0063\u0078":
			_geced._gdbee = _beeg.parseFloatAttr(_cgedc, _ddgc)
		case "\u0063\u0079":
			_geced._bgfb = _beeg.parseFloatAttr(_cgedc, _ddgc)
		case "\u0077\u0069\u0064t\u0068":
			_geced.SetWidth(_beeg.parseFloatAttr(_cgedc, _ddgc))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_geced.SetHeight(_beeg.parseFloatAttr(_cgedc, _ddgc))
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_geced.SetFillColor(_beeg.parseColorAttr(_cgedc, _ddgc))
		case "\u0066\u0069\u006cl\u002d\u006f\u0070\u0061\u0063\u0069\u0074\u0079":
			_geced.SetFillOpacity(_beeg.parseFloatAttr(_cgedc, _ddgc))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_geced.SetBorderColor(_beeg.parseColorAttr(_cgedc, _ddgc))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006f\u0070a\u0063\u0069\u0074\u0079":
			_geced.SetBorderOpacity(_beeg.parseFloatAttr(_cgedc, _ddgc))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_geced.SetBorderWidth(_beeg.parseFloatAttr(_cgedc, _ddgc))
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_geced.SetPositioning(_beeg.parsePositioningAttr(_cgedc, _ddgc))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_geced.SetFitMode(_beeg.parseFitModeAttr(_cgedc, _ddgc))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_cgfbe := _beeg.parseMarginAttr(_cgedc, _ddgc)
			_geced.SetMargins(_cgfbe.Left, _cgfbe.Right, _cgfbe.Top, _cgfbe.Bottom)
		default:
			_beeg.nodeLogDebug(_afbec, "\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u006c\u0069\u0070\u0073\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _cgedc)
		}
	}
	return _geced, nil
}

// Text sets the text content of the Paragraph.
func (_fffd *Paragraph) Text() string { return _fffd._bfgg }

const (
	CellVerticalAlignmentTop CellVerticalAlignment = iota
	CellVerticalAlignmentMiddle
	CellVerticalAlignmentBottom
)

// Polygon represents a polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Polygon struct {
	_acaag *_ff.Polygon
	_afdbc float64
	_bgac  float64
	_ddad  Color
}

// GetMargins returns the Chapter's margin: left, right, top, bottom.
func (_deg *Chapter) GetMargins() (float64, float64, float64, float64) {
	return _deg._ecd.Left, _deg._ecd.Right, _deg._ecd.Top, _deg._ecd.Bottom
}

// Width returns the width of the rectangle.
// NOTE: the returned value does not include the border width of the rectangle.
func (_gfefc *Rectangle) Width() float64 { return _gfefc._bcffg }

// NewEllipse creates a new ellipse with the center at (`xc`, `yc`),
// having the specified width and height.
// NOTE: In relative positioning mode, `xc` and `yc` are calculated using the
// current context. Furthermore, when the fit mode is set to fill the available
// space, the ellipse is scaled so that it occupies the entire context width
// while maintaining the original aspect ratio.
func (_dba *Creator) NewEllipse(xc, yc, width, height float64) *Ellipse {
	return _fccae(xc, yc, width, height)
}

// TemplateOptions contains options and resources to use when rendering
// a template with a Creator instance.
// All the resources in the map fields can be referenced by their
// name/key in the template which is rendered using the options instance.
type TemplateOptions struct {
	// HelperFuncMap is used to define functions which can be accessed
	// inside the rendered templates by their assigned names.
	HelperFuncMap _dd.FuncMap

	// SubtemplateMap contains templates which can be rendered alongside
	// the main template. They can be accessed using their assigned names
	// in the main template or in the other subtemplates.
	// Subtemplates defined inside the subtemplates specified in the map
	// can be accessed directly.
	// All resources available to the main template are also available
	// to the subtemplates.
	SubtemplateMap map[string]_eb.Reader

	// FontMap contains pre-loaded fonts which can be accessed
	// inside the rendered templates by their assigned names.
	FontMap map[string]*_ga.PdfFont

	// ImageMap contains pre-loaded images which can be accessed
	// inside the rendered templates by their assigned names.
	ImageMap map[string]*_ga.Image

	// ColorMap contains colors which can be accessed
	// inside the rendered templates by their assigned names.
	ColorMap map[string]Color

	// ChartMap contains charts which can be accessed
	// inside the rendered templates by their assigned names.
	ChartMap map[string]_dce.ChartRenderable
}

func _bbgb(_fabc _dce.ChartRenderable) *Chart {
	return &Chart{_dece: _fabc, _bdcb: PositionRelative, _baae: Margins{Top: 10, Bottom: 10}}
}

// List represents a list of items.
// The representation of a list item is as follows:
//
//	[marker] [content]
//
// e.g.:        • This is the content of the item.
// The supported components to add content to list items are:
// - Paragraph
// - StyledParagraph
// - List
type List struct {
	_bdccg []*listItem
	_fcag  Margins
	_feeb  TextChunk
	_addd  float64
	_ecee  bool
	_ffeef Positioning
	_edda  TextStyle
}

// Reset removes all the text chunks the paragraph contains.
func (_abaad *StyledParagraph) Reset() { _abaad._eddc = []*TextChunk{} }

const (
	TextRenderingModeFill TextRenderingMode = iota
	TextRenderingModeStroke
	TextRenderingModeFillStroke
	TextRenderingModeInvisible
	TextRenderingModeFillClip
	TextRenderingModeStrokeClip
	TextRenderingModeFillStrokeClip
	TextRenderingModeClip
)

// GeneratePageBlocks implements drawable interface.
func (_adcf *border) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_eag := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_gda := _adcf._cbd
	_bga := ctx.PageHeight - _adcf._bcgdb
	if _adcf._fga != nil {
		_edd := _ff.Rectangle{Opacity: 1.0, X: _adcf._cbd, Y: ctx.PageHeight - _adcf._bcgdb - _adcf._cfcd, Height: _adcf._cfcd, Width: _adcf._begd}
		_edd.FillEnabled = true
		_cdcc := _fce(_adcf._fga)
		_gdae := _acdb(_eag, _cdcc, _adcf._fga, func() Rectangle {
			return Rectangle{_fcgf: _edd.X, _cbaa: _edd.Y, _bcffg: _edd.Width, _ebegg: _edd.Height}
		})
		if _gdae != nil {
			return nil, ctx, _gdae
		}
		_edd.FillColor = _cdcc
		_edd.BorderEnabled = false
		_fgg, _, _gdae := _edd.Draw("")
		if _gdae != nil {
			return nil, ctx, _gdae
		}
		_gdae = _eag.addContentsByString(string(_fgg))
		if _gdae != nil {
			return nil, ctx, _gdae
		}
	}
	_cgfd := _adcf._dca
	_fadc := _adcf._gab
	_ebac := _adcf._fab
	_cgfe := _adcf._dggc
	_bbab := _adcf._dca
	if _adcf._acef == CellBorderStyleDouble {
		_bbab += 2 * _cgfd
	}
	_aafg := _adcf._gab
	if _adcf._aacc == CellBorderStyleDouble {
		_aafg += 2 * _fadc
	}
	_beeae := _adcf._fab
	if _adcf._dad == CellBorderStyleDouble {
		_beeae += 2 * _ebac
	}
	_gdd := _adcf._dggc
	if _adcf._fbae == CellBorderStyleDouble {
		_gdd += 2 * _cgfe
	}
	_gae := (_bbab - _beeae) / 2
	_begg := (_bbab - _gdd) / 2
	_ddaa := (_aafg - _beeae) / 2
	_fage := (_aafg - _gdd) / 2
	if _adcf._dca != 0 {
		_cgb := _gda
		_bfa := _bga
		if _adcf._acef == CellBorderStyleDouble {
			_bfa -= _cgfd
			_bab := _ff.BasicLine{LineColor: _fce(_adcf._beea), Opacity: 1.0, LineWidth: _adcf._dca, LineStyle: _adcf.LineStyle, X1: _cgb - _bbab/2 + _gae, Y1: _bfa + 2*_cgfd, X2: _cgb + _bbab/2 - _begg + _adcf._begd, Y2: _bfa + 2*_cgfd}
			_dee, _, _bed := _bab.Draw("")
			if _bed != nil {
				return nil, ctx, _bed
			}
			_bed = _eag.addContentsByString(string(_dee))
			if _bed != nil {
				return nil, ctx, _bed
			}
		}
		_bafa := _ff.BasicLine{LineWidth: _adcf._dca, Opacity: 1.0, LineColor: _fce(_adcf._beea), LineStyle: _adcf.LineStyle, X1: _cgb - _bbab/2 + _gae + (_beeae - _adcf._fab), Y1: _bfa, X2: _cgb + _bbab/2 - _begg + _adcf._begd - (_gdd - _adcf._dggc), Y2: _bfa}
		_faf, _, _abeg := _bafa.Draw("")
		if _abeg != nil {
			return nil, ctx, _abeg
		}
		_abeg = _eag.addContentsByString(string(_faf))
		if _abeg != nil {
			return nil, ctx, _abeg
		}
	}
	if _adcf._gab != 0 {
		_ecc := _gda
		_fgba := _bga - _adcf._cfcd
		if _adcf._aacc == CellBorderStyleDouble {
			_fgba += _fadc
			_facd := _ff.BasicLine{LineWidth: _adcf._gab, Opacity: 1.0, LineColor: _fce(_adcf._adc), LineStyle: _adcf.LineStyle, X1: _ecc - _aafg/2 + _ddaa, Y1: _fgba - 2*_fadc, X2: _ecc + _aafg/2 - _fage + _adcf._begd, Y2: _fgba - 2*_fadc}
			_adae, _, _deb := _facd.Draw("")
			if _deb != nil {
				return nil, ctx, _deb
			}
			_deb = _eag.addContentsByString(string(_adae))
			if _deb != nil {
				return nil, ctx, _deb
			}
		}
		_cgfc := _ff.BasicLine{LineWidth: _adcf._gab, Opacity: 1.0, LineColor: _fce(_adcf._adc), LineStyle: _adcf.LineStyle, X1: _ecc - _aafg/2 + _ddaa + (_beeae - _adcf._fab), Y1: _fgba, X2: _ecc + _aafg/2 - _fage + _adcf._begd - (_gdd - _adcf._dggc), Y2: _fgba}
		_dgcf, _, _gfc := _cgfc.Draw("")
		if _gfc != nil {
			return nil, ctx, _gfc
		}
		_gfc = _eag.addContentsByString(string(_dgcf))
		if _gfc != nil {
			return nil, ctx, _gfc
		}
	}
	if _adcf._fab != 0 {
		_ecef := _gda
		_cadc := _bga
		if _adcf._dad == CellBorderStyleDouble {
			_ecef += _ebac
			_dff := _ff.BasicLine{LineWidth: _adcf._fab, Opacity: 1.0, LineColor: _fce(_adcf._efb), LineStyle: _adcf.LineStyle, X1: _ecef - 2*_ebac, Y1: _cadc + _beeae/2 + _gae, X2: _ecef - 2*_ebac, Y2: _cadc - _beeae/2 - _ddaa - _adcf._cfcd}
			_gdg, _, _gfcg := _dff.Draw("")
			if _gfcg != nil {
				return nil, ctx, _gfcg
			}
			_gfcg = _eag.addContentsByString(string(_gdg))
			if _gfcg != nil {
				return nil, ctx, _gfcg
			}
		}
		_bfbg := _ff.BasicLine{LineWidth: _adcf._fab, Opacity: 1.0, LineColor: _fce(_adcf._efb), LineStyle: _adcf.LineStyle, X1: _ecef, Y1: _cadc + _beeae/2 + _gae - (_bbab - _adcf._dca), X2: _ecef, Y2: _cadc - _beeae/2 - _ddaa - _adcf._cfcd + (_aafg - _adcf._gab)}
		_dddg, _, _cff := _bfbg.Draw("")
		if _cff != nil {
			return nil, ctx, _cff
		}
		_cff = _eag.addContentsByString(string(_dddg))
		if _cff != nil {
			return nil, ctx, _cff
		}
	}
	if _adcf._dggc != 0 {
		_aacb := _gda + _adcf._begd
		_bdg := _bga
		if _adcf._fbae == CellBorderStyleDouble {
			_aacb -= _cgfe
			_dgb := _ff.BasicLine{LineWidth: _adcf._dggc, Opacity: 1.0, LineColor: _fce(_adcf._gdc), LineStyle: _adcf.LineStyle, X1: _aacb + 2*_cgfe, Y1: _bdg + _gdd/2 + _begg, X2: _aacb + 2*_cgfe, Y2: _bdg - _gdd/2 - _fage - _adcf._cfcd}
			_gcae, _, _egbb := _dgb.Draw("")
			if _egbb != nil {
				return nil, ctx, _egbb
			}
			_egbb = _eag.addContentsByString(string(_gcae))
			if _egbb != nil {
				return nil, ctx, _egbb
			}
		}
		_baac := _ff.BasicLine{LineWidth: _adcf._dggc, Opacity: 1.0, LineColor: _fce(_adcf._gdc), LineStyle: _adcf.LineStyle, X1: _aacb, Y1: _bdg + _gdd/2 + _begg - (_bbab - _adcf._dca), X2: _aacb, Y2: _bdg - _gdd/2 - _fage - _adcf._cfcd + (_aafg - _adcf._gab)}
		_def, _, _fde := _baac.Draw("")
		if _fde != nil {
			return nil, ctx, _fde
		}
		_fde = _eag.addContentsByString(string(_def))
		if _fde != nil {
			return nil, ctx, _fde
		}
	}
	return []*Block{_eag}, ctx, nil
}

// SetExtends specifies whether ot extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_ebbee *RadialShading) SetExtends(start bool, end bool) { _ebbee._acdae.SetExtends(start, end) }

// ScaleToWidth scales the ellipse to the specified width. The height of
// the ellipse is scaled so that the aspect ratio is maintained.
func (_bcbg *Ellipse) ScaleToWidth(w float64) {
	_agcd := _bcbg._bdcce / _bcbg._eded
	_bcbg._eded = w
	_bcbg._bdcce = w * _agcd
}

// TextDecorationLineStyle represents the style of lines used to decorate
// a text chunk (e.g. underline).
type TextDecorationLineStyle struct {
	// Color represents the color of the line (default: the color of the text).
	Color Color

	// Offset represents the vertical offset of the line (default: 1).
	Offset float64

	// Thickness represents the thickness of the line (default: 1).
	Thickness float64
}

// SetPos sets the position of the chart to the specified coordinates.
// This method sets the chart to use absolute positioning.
func (_ecg *Chart) SetPos(x, y float64) { _ecg._bdcb = PositionAbsolute; _ecg._cdd = x; _ecg._dcd = y }

// AddLine adds a new line with the provided style to the table of contents.
func (_dgceac *TOC) AddLine(line *TOCLine) *TOCLine {
	if line == nil {
		return nil
	}
	_dgceac._gcgf = append(_dgceac._gcgf, line)
	return line
}

// SetMargins sets the margins of the chart component.
func (_dcg *Chart) SetMargins(left, right, top, bottom float64) {
	_dcg._baae.Left = left
	_dcg._baae.Right = right
	_dcg._baae.Top = top
	_dcg._baae.Bottom = bottom
}

func (_fbccd *Invoice) generateInformationBlocks(_efeca DrawContext) ([]*Block, DrawContext, error) {
	_eaec := _cgfa(_fbccd._dcefb)
	_eaec.SetMargins(0, 0, 0, 20)
	_dgefa := _fbccd.drawAddress(_fbccd._dceeg)
	_dgefa = append(_dgefa, _eaec)
	_dgefa = append(_dgefa, _fbccd.drawAddress(_fbccd._cfccb)...)
	_cfe := _dfdab()
	for _, _dddd := range _dgefa {
		_cfe.Add(_dddd)
	}
	_abc := _fbccd.drawInformation()
	_efed := _cbgg(2)
	_efed.SetMargins(0, 0, 25, 0)
	_gegb := _efed.NewCell()
	_gegb.SetIndent(0)
	_gegb.SetContent(_cfe)
	_gegb = _efed.NewCell()
	_gegb.SetContent(_abc)
	return _efed.GeneratePageBlocks(_efeca)
}

// AddShadingResource adds shading dictionary inside the resources dictionary.
func (_cafg *LinearShading) AddShadingResource(block *Block) (_egba _cc.PdfObjectName, _edcca error) {
	_feda := 1
	_egba = _cc.PdfObjectName("\u0053\u0068" + _aa.Itoa(_feda))
	for block._ccf.HasShadingByName(_egba) {
		_feda++
		_egba = _cc.PdfObjectName("\u0053\u0068" + _aa.Itoa(_feda))
	}
	if _fcbce := block._ccf.SetShadingByName(_egba, _cafg.shadingModel().ToPdfObject()); _fcbce != nil {
		return "", _fcbce
	}
	return _egba, nil
}

type templateNode struct {
	_defae interface{}
	_acdf  _dda.StartElement
	_gbdf  *templateNode
	_ffgg  int
	_gcdae int
	_cgcf  int64
}

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_fccbb *shading) SetBackgroundColor(backgroundColor Color) { _fccbb._cfec = backgroundColor }

// SetLineTitleStyle sets the style for the title part of all new lines
// of the table of contents.
func (_bbfa *TOC) SetLineTitleStyle(style TextStyle) { _bbfa._gdeab = style }

// SetFitMode sets the fit mode of the line.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_ffee *Line) SetFitMode(fitMode FitMode) { _ffee._cbab = fitMode }

// Height returns the height of the division, assuming all components are
// stacked on top of each other.
func (_edac *Division) Height() float64 {
	var _bcgc float64
	for _, _eebe := range _edac._dccd {
		switch _dfaf := _eebe.(type) {
		case marginDrawable:
			_, _, _cbeb, _baacf := _dfaf.GetMargins()
			_bcgc += _dfaf.Height() + _cbeb + _baacf
		default:
			_bcgc += _dfaf.Height()
		}
	}
	return _bcgc
}

// SetNoteStyle sets the style properties used to render the content of the
// invoice note sections.
func (_cgfb *Invoice) SetNoteStyle(style TextStyle) { _cgfb._fdbgg = style }

// NoteStyle returns the style properties used to render the content of the
// invoice note sections.
func (_ceaea *Invoice) NoteStyle() TextStyle { return _ceaea._fdbgg }

func _bceg(_ecafd *Block, _cgdf *StyledParagraph, _cabdd [][]*TextChunk, _ffdce DrawContext) (DrawContext, [][]*TextChunk, error) {
	_cegge := 1
	_edec := _cc.PdfObjectName(_g.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cegge))
	for _ecafd._ccf.HasFontByName(_edec) {
		_cegge++
		_edec = _cc.PdfObjectName(_g.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cegge))
	}
	_eegc := _ecafd._ccf.SetFontByName(_edec, _cgdf._afaf.Font.ToPdfObject())
	if _eegc != nil {
		return _ffdce, nil, _eegc
	}
	_cegge++
	_gbdd := _edec
	_cefa := _cgdf._afaf.FontSize
	_gfdg := _cgdf._bedd.IsRelative()
	var _gdfge [][]_cc.PdfObjectName
	var _cdggb [][]*TextChunk
	var _gegc float64
	for _fedg, _bged := range _cabdd {
		var _egbg []_cc.PdfObjectName
		var _aebc float64
		if len(_bged) > 0 {
			_aebc = _bged[0].Style.FontSize
		}
		for _, _bfff := range _bged {
			_gcad := _bfff.Style
			if _bfff.Text != "" && _gcad.FontSize > _aebc {
				_aebc = _gcad.FontSize
			}
			if _aebc > _ffdce.PageHeight {
				return _ffdce, nil, _c.New("\u0050\u0061\u0072\u0061\u0067\u0072a\u0070\u0068\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020\u006ca\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0070\u0061\u0067\u0065 \u0068e\u0069\u0067\u0068\u0074")
			}
			_edec = _cc.PdfObjectName(_g.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cegge))
			_eafbg := _ecafd._ccf.SetFontByName(_edec, _gcad.Font.ToPdfObject())
			if _eafbg != nil {
				return _ffdce, nil, _eafbg
			}
			_egbg = append(_egbg, _edec)
			_cegge++
		}
		_aebc *= _cgdf._bebe
		if _gfdg && _gegc+_aebc > _ffdce.Height {
			_cdggb = _cabdd[_fedg:]
			_cabdd = _cabdd[:_fedg]
			break
		}
		_gegc += _aebc
		_gdfge = append(_gdfge, _egbg)
	}
	_ffbed, _gdgf, _ccab := _cgdf.getLineMetrics(0)
	_adfaf, _baea := _ffbed*_cgdf._bebe, _gdgf*_cgdf._bebe
	if len(_cabdd) == 0 {
		return _ffdce, _cdggb, nil
	}
	_bcba := _da.NewContentCreator()
	_bcba.Add_q()
	_gbdc := _baea
	if _cgdf._beda == TextVerticalAlignmentCenter {
		_gbdc = _gdgf + (_ffbed+_ccab-_gdgf)/2 + (_baea-_gdgf)/2
	}
	_edfdb := _ffdce.PageHeight - _ffdce.Y - _gbdc
	_bcba.Translate(_ffdce.X, _edfdb)
	_cabae := _edfdb
	if _cgdf._ecdcf != 0 {
		_bcba.RotateDeg(_cgdf._ecdcf)
	}
	if _cgdf._egead == TextOverflowHidden {
		_bcba.Add_re(0, -_gegc+_adfaf+1, _cgdf._cccde, _gegc).Add_W().Add_n()
	}
	_bcba.Add_BT()
	var _eaged []*_ff.BasicLine
	for _fbcga, _edcac := range _cabdd {
		_agbfe := _ffdce.X
		var _dbagc float64
		if len(_edcac) > 0 {
			_dbagc = _edcac[0].Style.FontSize
		}
		_ffbed, _, _ccab = _cgdf.getLineMetrics(_fbcga)
		_baea = (_ffbed + _ccab)
		for _, _baeg := range _edcac {
			_bace := &_baeg.Style
			if _baeg.Text != "" && _bace.FontSize > _dbagc {
				_dbagc = _bace.FontSize
			}
			if _baea > _dbagc {
				_dbagc = _baea
			}
		}
		if _fbcga != 0 {
			_bcba.Add_TD(0, -_dbagc*_cgdf._bebe)
			_cabae -= _dbagc * _cgdf._bebe
		}
		_dffa := _fbcga == len(_cabdd)-1
		var (
			_aabdd float64
			_edgef float64
			_ffea  *fontMetrics
			_gaffb float64
			_ecgg  uint
		)
		var _dacbc []float64
		for _, _bgdac := range _edcac {
			_deaea := &_bgdac.Style
			if _deaea.FontSize > _edgef {
				_edgef = _deaea.FontSize
				_ffea = _abff(_bgdac.Style.Font, _deaea.FontSize)
			}
			if _baea > _edgef {
				_edgef = _baea
			}
			_gbaa, _gabba := _deaea.Font.GetRuneMetrics(' ')
			if !_gabba {
				return _ffdce, nil, _c.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
			}
			var _fbgcc uint
			var _bgfg float64
			_faga := len(_bgdac.Text)
			for _bcebd, _fgga := range _bgdac.Text {
				if _fgga == ' ' {
					_fbgcc++
					continue
				}
				if _fgga == '\u000A' {
					continue
				}
				_fbba, _febc := _deaea.Font.GetRuneMetrics(_fgga)
				if !_febc {
					_bcd.Log.Debug("\u0055\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006ee\u0020%\u0076\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u000a", _fgga)
					return _ffdce, nil, _c.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
				}
				_bgfg += _deaea.FontSize * _fbba.Wx * _deaea.horizontalScale()
				if _bcebd != _faga-1 {
					_bgfg += _deaea.CharSpacing * 1000.0
				}
			}
			_dacbc = append(_dacbc, _bgfg)
			_aabdd += _bgfg
			_gaffb += float64(_fbgcc) * _gbaa.Wx * _deaea.FontSize * _deaea.horizontalScale()
			_ecgg += _fbgcc
		}
		_edgef *= _cgdf._bebe
		var _cagc []_cc.PdfObject
		_aagef := _cgdf._cccde * 1000.0
		if _cgdf._dcgfc == TextAlignmentJustify {
			if _ecgg > 0 && !_dffa {
				_gaffb = (_aagef - _aabdd) / float64(_ecgg) / _cefa
			}
		} else if _cgdf._dcgfc == TextAlignmentCenter {
			_egbfa := (_aagef - _aabdd - _gaffb) / 2
			_cbge := _egbfa / _cefa
			_cagc = append(_cagc, _cc.MakeFloat(-_cbge))
			_agbfe += _egbfa / 1000.0
		} else if _cgdf._dcgfc == TextAlignmentRight {
			_baagg := (_aagef - _aabdd - _gaffb)
			_eece := _baagg / _cefa
			_cagc = append(_cagc, _cc.MakeFloat(-_eece))
			_agbfe += _baagg / 1000.0
		}
		if len(_cagc) > 0 {
			_bcba.Add_Tf(_gbdd, _cefa).Add_TL(_cefa * _cgdf._bebe).Add_TJ(_cagc...)
		}
		_feef := 0.0
		for _gaed, _fgdc := range _edcac {
			_dcdc := &_fgdc.Style
			_ecba := _gbdd
			_cbgab := _cefa
			_bgad := _dcdc.OutlineColor != nil
			_fgdge := _dcdc.HorizontalScaling != DefaultHorizontalScaling
			_eebf := _dcdc.OutlineSize != 1
			if _eebf {
				_bcba.Add_w(_dcdc.OutlineSize)
			}
			_dbfd := _dcdc.RenderingMode != TextRenderingModeFill
			if _dbfd {
				_bcba.Add_Tr(int64(_dcdc.RenderingMode))
			}
			_aeeg := _dcdc.CharSpacing != 0
			if _aeeg {
				_bcba.Add_Tc(_dcdc.CharSpacing)
			}
			_cfbea := _dcdc.TextRise != 0
			if _cfbea {
				_bcba.Add_Ts(_dcdc.TextRise)
			}
			if _fgdc.VerticalAlignment != TextVerticalAlignmentBaseline {
				_gbdg := _abff(_fgdc.Style.Font, _dcdc.FontSize)
				switch _fgdc.VerticalAlignment {
				case TextVerticalAlignmentCenter:
					_feef = _ffea._bgggg/2 - _gbdg._bgggg/2
				case TextVerticalAlignmentBottom:
					_feef = _ffea._ceddg - _gbdg._ceddg
				case TextVerticalAlignmentTop:
					_feef = _gdgf - _dcdc.FontSize
				}
				if _feef != 0.0 {
					_bcba.Translate(0, _feef)
				}
			}
			if _cgdf._dcgfc != TextAlignmentJustify || _dffa {
				_cacbd, _bfcge := _dcdc.Font.GetRuneMetrics(' ')
				if !_bfcge {
					return _ffdce, nil, _c.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
				}
				_ecba = _gdfge[_fbcga][_gaed]
				_cbgab = _dcdc.FontSize
				_gaffb = _cacbd.Wx * _dcdc.horizontalScale()
			}
			_bbada := _dcdc.Font.Encoder()
			var _dagbb []byte
			for _, _ebecf := range _fgdc.Text {
				if _ebecf == '\u000A' {
					continue
				}
				if _ebecf == ' ' {
					if len(_dagbb) > 0 {
						if _bgad {
							_bcba.SetStrokingColor(_fce(_dcdc.OutlineColor))
						}
						if _fgdge {
							_bcba.Add_Tz(_dcdc.HorizontalScaling)
						}
						_bcba.SetNonStrokingColor(_fce(_dcdc.Color)).Add_Tf(_gdfge[_fbcga][_gaed], _dcdc.FontSize).Add_TJ([]_cc.PdfObject{_cc.MakeStringFromBytes(_dagbb)}...)
						_dagbb = nil
					}
					if _fgdge {
						_bcba.Add_Tz(DefaultHorizontalScaling)
					}
					_bcba.Add_Tf(_ecba, _cbgab).Add_TJ([]_cc.PdfObject{_cc.MakeFloat(-_gaffb)}...)
					_dacbc[_gaed] += _gaffb * _cbgab
				} else {
					if _, _edgc := _bbada.RuneToCharcode(_ebecf); !_edgc {
						_eegc = UnsupportedRuneError{Message: _g.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _ebecf, _ebecf), Rune: _ebecf}
						_ffdce._ddcc = append(_ffdce._ddcc, _eegc)
						_bcd.Log.Debug(_eegc.Error())
						if _ffdce._dfe <= 0 {
							continue
						}
						_ebecf = _ffdce._dfe
					}
					_dagbb = append(_dagbb, _bbada.Encode(string(_ebecf))...)
				}
			}
			if len(_dagbb) > 0 {
				if _bgad {
					_bcba.SetStrokingColor(_fce(_dcdc.OutlineColor))
				}
				if _fgdge {
					_bcba.Add_Tz(_dcdc.HorizontalScaling)
				}
				_bcba.SetNonStrokingColor(_fce(_dcdc.Color)).Add_Tf(_gdfge[_fbcga][_gaed], _dcdc.FontSize).Add_TJ([]_cc.PdfObject{_cc.MakeStringFromBytes(_dagbb)}...)
			}
			_dfgcb := _dacbc[_gaed] / 1000.0
			if _dcdc.Underline {
				_ffae := _dcdc.UnderlineStyle.Color
				if _ffae == nil {
					_ffae = _fgdc.Style.Color
				}
				_ffde, _gbgge, _daaeg := _ffae.ToRGB()
				_fgcf := _agbfe - _ffdce.X
				_bcca := _cabae - _edfdb + _dcdc.TextRise - _dcdc.UnderlineStyle.Offset
				_eaged = append(_eaged, &_ff.BasicLine{X1: _fgcf, Y1: _bcca, X2: _fgcf + _dfgcb, Y2: _bcca, LineWidth: _fgdc.Style.UnderlineStyle.Thickness, LineColor: _ga.NewPdfColorDeviceRGB(_ffde, _gbgge, _daaeg)})
			}
			if _fgdc._cgfaa != nil {
				var _efcb *_cc.PdfObjectArray
				if !_fgdc._gedb {
					switch _cbgfe := _fgdc._cgfaa.GetContext().(type) {
					case *_ga.PdfAnnotationLink:
						_efcb = _cc.MakeArray()
						_cbgfe.Rect = _efcb
						_cbdcb, _eedd := _cbgfe.Dest.(*_cc.PdfObjectArray)
						if _eedd && _cbdcb.Len() == 5 {
							_adcb, _cbdd := _cbdcb.Get(1).(*_cc.PdfObjectName)
							if _cbdd && _adcb.String() == "\u0058\u0059\u005a" {
								_ebdef, _gcac := _cc.GetNumberAsFloat(_cbdcb.Get(3))
								if _gcac == nil {
									_cbdcb.Set(3, _cc.MakeFloat(_ffdce.PageHeight-_ebdef))
								}
							}
						}
					}
					_fgdc._gedb = true
				}
				if _efcb != nil {
					_bdbgg := _ff.NewPoint(_agbfe-_ffdce.X, _cabae+_dcdc.TextRise-_edfdb).Rotate(_cgdf._ecdcf)
					_bdbgg.X += _ffdce.X
					_bdbgg.Y += _edfdb
					_cbabd, _gcge, _bdba, _befce := _aecbg(_dfgcb, _edgef, _cgdf._ecdcf)
					_bdbgg.X += _cbabd
					_bdbgg.Y += _gcge
					_efcb.Clear()
					_efcb.Append(_cc.MakeFloat(_bdbgg.X))
					_efcb.Append(_cc.MakeFloat(_bdbgg.Y))
					_efcb.Append(_cc.MakeFloat(_bdbgg.X + _bdba))
					_efcb.Append(_cc.MakeFloat(_bdbgg.Y + _befce))
				}
				_ecafd.AddAnnotation(_fgdc._cgfaa)
			}
			_agbfe += _dfgcb
			if _eebf {
				_bcba.Add_w(1.0)
			}
			if _bgad {
				_bcba.Add_RG(0.0, 0.0, 0.0)
			}
			if _dbfd {
				_bcba.Add_Tr(int64(TextRenderingModeFill))
			}
			if _aeeg {
				_bcba.Add_Tc(0)
			}
			if _cfbea {
				_bcba.Add_Ts(0)
			}
			if _fgdge {
				_bcba.Add_Tz(DefaultHorizontalScaling)
			}
			if _feef != 0.0 {
				_bcba.Translate(0, -_feef)
				_feef = 0.0
			}
		}
	}
	_bcba.Add_ET()
	for _, _gfbdf := range _eaged {
		_bcba.SetStrokingColor(_gfbdf.LineColor).Add_w(_gfbdf.LineWidth).Add_m(_gfbdf.X1, _gfbdf.Y1).Add_l(_gfbdf.X2, _gfbdf.Y2).Add_s()
	}
	_bcba.Add_Q()
	_ebcdf := _bcba.Operations()
	_ebcdf.WrapIfNeeded()
	_ecafd.addContents(_ebcdf)
	if _gfdg {
		_bbcbd := _gegc
		_ffdce.Y += _bbcbd
		_ffdce.Height -= _bbcbd
		if _ffdce.Inline {
			_ffdce.X += _cgdf.Width() + _cgdf._fadcg.Right
		}
	}
	return _ffdce, _cdggb, nil
}

// TextChunk represents a chunk of text along with a particular style.
type TextChunk struct {
	// The text that is being rendered in the PDF.
	Text string

	// The style of the text being rendered.
	Style  TextStyle
	_cgfaa *_ga.PdfAnnotation
	_gedb  bool

	// The vertical alignment of the text chunk.
	VerticalAlignment TextVerticalAlignment
}

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

const (
	CellHorizontalAlignmentLeft CellHorizontalAlignment = iota
	CellHorizontalAlignmentCenter
	CellHorizontalAlignmentRight
)

// ScaleToWidth scale Image to a specified width w, maintaining the aspect ratio.
func (_agfb *Image) ScaleToWidth(w float64) {
	_cgedg := _agfb._cgdb / _agfb._ggda
	_agfb._ggda = w
	_agfb._cgdb = w * _cgedg
}

type pageTransformations struct {
	_aefc *_gb.Matrix
	_cca  bool
	_eea  bool
}

// Paragraph represents text drawn with a specified font and can wrap across lines and pages.
// By default it occupies the available width in the drawing context.
type Paragraph struct {
	_bfgg         string
	_bfbd         *_ga.PdfFont
	_dbad         float64
	_agdb         float64
	_ggbe         Color
	_caea         TextAlignment
	_dbfb         bool
	_adcfe        float64
	_gfbef        int
	_deaf         bool
	_cgaaf        float64
	_bcggf        Margins
	_gece         Positioning
	_acab         float64
	_bge          float64
	_ddeg, _fabcb float64
	_aggb         []string
}

// SetFillOpacity sets the fill opacity of the ellipse.
func (_fgae *Ellipse) SetFillOpacity(opacity float64) { _fgae._bbdc = opacity }

func (_bgfee *shading) generatePdfFunctions() []_ga.PdfFunction {
	if len(_bgfee._ccaf) == 0 {
		return nil
	} else if len(_bgfee._ccaf) <= 2 {
		_dffe, _facfa, _ddbb := _bgfee._ccaf[0]._egea.ToRGB()
		_dabgc, _efddd, _gfcab := _bgfee._ccaf[len(_bgfee._ccaf)-1]._egea.ToRGB()
		return []_ga.PdfFunction{&_ga.PdfFunctionType2{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, N: 1, C0: []float64{_dffe, _facfa, _ddbb}, C1: []float64{_dabgc, _efddd, _gfcab}}}
	} else {
		_beca := []_ga.PdfFunction{}
		_gfaba := []float64{}
		for _bfab := 0; _bfab < len(_bgfee._ccaf)-1; _bfab++ {
			_afdc, _abac, _afgbf := _bgfee._ccaf[_bfab]._egea.ToRGB()
			_gedc, _dbee, _gbfc := _bgfee._ccaf[_bfab+1]._egea.ToRGB()
			_bbeg := &_ga.PdfFunctionType2{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, N: 1, C0: []float64{_afdc, _abac, _afgbf}, C1: []float64{_gedc, _dbee, _gbfc}}
			_beca = append(_beca, _bbeg)
			if _bfab > 0 {
				_gfaba = append(_gfaba, _bgfee._ccaf[_bfab]._ecbg)
			}
		}
		_bbge := []float64{}
		for range _beca {
			_bbge = append(_bbge, []float64{0.0, 1.0}...)
		}
		return []_ga.PdfFunction{&_ga.PdfFunctionType3{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, Functions: _beca, Bounds: _gfaba, Encode: _bbge}}
	}
}

// SetIncludeInTOC sets a flag to indicate whether or not to include in tOC.
func (_adff *Chapter) SetIncludeInTOC(includeInTOC bool) { _adff._gge = includeInTOC }

// CellHorizontalAlignment defines the table cell's horizontal alignment.
type CellHorizontalAlignment int

func (_aceag *templateProcessor) parseListMarker(_fbfgg *templateNode) (interface{}, error) {
	if _fbfgg._gbdf == nil {
		_aceag.nodeLogError(_fbfgg, "\u004c\u0069\u0073\u0074\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0063a\u006e\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c\u002e")
		return nil, _bfegb
	}
	var _ebef *TextChunk
	switch _bbadg := _fbfgg._gbdf._defae.(type) {
	case *List:
		_ebef = &_bbadg._feeb
	case *listItem:
		_ebef = &_bbadg._feced
	default:
		_aceag.nodeLogError(_fbfgg, "\u0025\u0076 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0066\u006f\u0072\u0020\u006c\u0069\u0073\u0074\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e", _bbadg)
		return nil, _bfegb
	}
	if _, _faef := _aceag.parseTextChunk(_fbfgg, _ebef); _faef != nil {
		_aceag.nodeLogError(_fbfgg, "\u0043\u006f\u0075ld\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072\u0073e\u0020l\u0069s\u0074 \u006d\u0061\u0072\u006b\u0065\u0072\u003a\u0020\u0060\u0025\u0076\u0060\u002e", _faef)
		return nil, nil
	}
	return _ebef, nil
}

// LevelOffset returns the amount of space an indentation level occupies.
func (_adedd *TOCLine) LevelOffset() float64 { return _adedd._dgefgc }

func _bfba(_fcga []_ff.Point) *Polyline {
	return &Polyline{_agbgf: &_ff.Polyline{Points: _fcga, LineColor: _ga.NewPdfColorDeviceRGB(0, 0, 0), LineWidth: 1.0}, _cbbd: 1.0}
}

// Draw draws the drawable d on the block.
// Note that the drawable must not wrap, i.e. only return one block. Otherwise an error is returned.
func (_gaf *Block) Draw(d Drawable) error {
	_dcfb := DrawContext{}
	_dcfb.Width = _gaf._fd
	_dcfb.Height = _gaf._cce
	_dcfb.PageWidth = _gaf._fd
	_dcfb.PageHeight = _gaf._cce
	_dcfb.X = 0
	_dcfb.Y = 0
	_ebgd, _, _edc := d.GeneratePageBlocks(_dcfb)
	if _edc != nil {
		return _edc
	}
	if len(_ebgd) != 1 {
		return ErrContentNotFit
	}
	for _, _ffc := range _ebgd {
		if _ae := _gaf.mergeBlocks(_ffc); _ae != nil {
			return _ae
		}
	}
	return nil
}

func _acdg(_bggce *_e.File) ([]*_ga.PdfPage, error) {
	_cbfggd, _bcfbf := _ga.NewPdfReader(_bggce)
	if _bcfbf != nil {
		return nil, _bcfbf
	}
	_eeebc, _bcfbf := _cbfggd.GetNumPages()
	if _bcfbf != nil {
		return nil, _bcfbf
	}
	var _fagcd []*_ga.PdfPage
	for _gebf := 0; _gebf < _eeebc; _gebf++ {
		_fdfb, _gefca := _cbfggd.GetPage(_gebf + 1)
		if _gefca != nil {
			return nil, _gefca
		}
		_fagcd = append(_fagcd, _fdfb)
	}
	return _fagcd, nil
}

func (_cedda *templateProcessor) parseRadialGradientAttr(creator *Creator, _aafgd string) Color {
	_gbcf := ColorBlack
	if _aafgd == "" {
		return _gbcf
	}
	var (
		_afbd   error
		_gdacc  = 0.0
		_cbccfc = 0.0
		_adbc   = -1.0
		_fdbfa  = _a.Split(_aafgd[16:len(_aafgd)-1], "\u002c")
	)
	_bcebf := _a.Fields(_fdbfa[0])
	if len(_bcebf) == 2 && _a.TrimSpace(_bcebf[0])[0] != '#' {
		_gdacc, _afbd = _aa.ParseFloat(_bcebf[0], 64)
		if _afbd != nil {
			_bcd.Log.Debug("\u0046a\u0069\u006ce\u0064\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0072a\u0064\u0069\u0061\u006c\u0020\u0067r\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0058\u0020\u0070\u006fs\u0069\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076", _afbd)
		}
		_cbccfc, _afbd = _aa.ParseFloat(_bcebf[1], 64)
		if _afbd != nil {
			_bcd.Log.Debug("\u0046a\u0069\u006ce\u0064\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0072a\u0064\u0069\u0061\u006c\u0020\u0067r\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0059\u0020\u0070\u006fs\u0069\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076", _afbd)
		}
		_fdbfa = _fdbfa[1:]
	}
	_aacace := _a.TrimSpace(_fdbfa[0])
	if _aacace[0] != '#' {
		_adbc, _afbd = _aa.ParseFloat(_aacace, 64)
		if _afbd != nil {
			_bcd.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0072\u0061\u0064\u0069\u0061l\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0073\u0069\u007ae\u003a\u0020\u0025\u0076", _afbd)
		}
		_fdbfa = _fdbfa[1:]
	}
	_fdgef, _cgede := _cedda.processGradientColorPair(_fdbfa)
	if _fdgef == nil || _cgede == nil {
		return _gbcf
	}
	_bgede := creator.NewRadialGradientColor(_gdacc, _cbccfc, 0, _adbc, []*ColorPoint{})
	for _ecfa := 0; _ecfa < len(_fdgef); _ecfa++ {
		_bgede.AddColorStop(_fdgef[_ecfa], _cgede[_ecfa])
	}
	return _bgede
}

// SetAnnotation sets a annotation on a TextChunk.
func (_befe *TextChunk) SetAnnotation(annotation *_ga.PdfAnnotation) { _befe._cgfaa = annotation }

func _dffg(_befa *_ga.PdfAnnotation) *_ga.PdfAnnotation {
	if _befa == nil {
		return nil
	}
	var _cebe *_ga.PdfAnnotation
	switch _ggaag := _befa.GetContext().(type) {
	case *_ga.PdfAnnotationLink:
		if _eeffe := _gbfe(_ggaag); _eeffe != nil {
			_cebe = _eeffe.PdfAnnotation
		}
	}
	return _cebe
}

// ToRGB implements interface Color.
// Note: It's not directly used since shading color works differently than regular color.
func (_acff *LinearShading) ToRGB() (float64, float64, float64) { return 0, 0, 0 }

// Height returns Image's document height.
func (_afeb *Image) Height() float64 { return _afeb._cgdb }

func (_fecaec *TableCell) height(_cgfca float64) float64 {
	var _agdf float64
	switch _dgdfb := _fecaec._bfef.(type) {
	case *Paragraph:
		if _dgdfb._dbfb {
			_dgdfb.SetWidth(_cgfca - _fecaec._fccd - _dgdfb._bcggf.Left - _dgdfb._bcggf.Right)
		}
		_agdf = _dgdfb.Height() + _dgdfb._bcggf.Top + _dgdfb._bcggf.Bottom
		if !_fecaec._ccbf._adbfb {
			_agdf += (0.5 * _dgdfb._dbad * _dgdfb._agdb)
		}
	case *StyledParagraph:
		if _dgdfb._bfcgb {
			_dgdfb.SetWidth(_cgfca - _fecaec._fccd - _dgdfb._fadcg.Left - _dgdfb._fadcg.Right)
		}
		_agdf = _dgdfb.Height() + _dgdfb._fadcg.Top + _dgdfb._fadcg.Bottom
		if !_fecaec._ccbf._adbfb {
			_agdf += (0.5 * _dgdfb.getTextHeight())
		}
	case *Image:
		_dgdfb.applyFitMode(_cgfca - _fecaec._fccd)
		_agdf = _dgdfb.Height() + _dgdfb._eaac.Top + _dgdfb._eaac.Bottom
	case *Table:
		_dgdfb.updateRowHeights(_cgfca - _fecaec._fccd - _dgdfb._bebd.Left - _dgdfb._bebd.Right)
		_agdf = _dgdfb.Height() + _dgdfb._bebd.Top + _dgdfb._bebd.Bottom
	case *List:
		_agdf = _dgdfb.ctxHeight(_cgfca-_fecaec._fccd) + _dgdfb._fcag.Top + _dgdfb._fcag.Bottom
	case *Division:
		_agdf = _dgdfb.ctxHeight(_cgfca-_fecaec._fccd) + _dgdfb._acaa.Top + _dgdfb._acaa.Bottom + _dgdfb._cgbb.Top + _dgdfb._cgbb.Bottom
	case *Chart:
		_agdf = _dgdfb.Height() + _dgdfb._baae.Top + _dgdfb._baae.Bottom
	case *Rectangle:
		_dgdfb.applyFitMode(_cgfca - _fecaec._fccd)
		_agdf = _dgdfb.Height() + _dgdfb._eaaa.Top + _dgdfb._eaaa.Bottom + _dgdfb._bffc
	case *Ellipse:
		_dgdfb.applyFitMode(_cgfca - _fecaec._fccd)
		_agdf = _dgdfb.Height() + _dgdfb._abgc.Top + _dgdfb._abgc.Bottom
	case *Line:
		_agdf = _dgdfb.Height() + _dgdfb._gabb.Top + _dgdfb._gabb.Bottom
	}
	return _agdf
}

// CellBorderSide defines the table cell's border side.
type CellBorderSide int

func (_abca *Table) sortCells() {
	_f.Slice(_abca._gebec, func(_gdbeb, _gaac int) bool {
		_bfbc := _abca._gebec[_gdbeb]._bgdca
		_defba := _abca._gebec[_gaac]._bgdca
		if _bfbc < _defba {
			return true
		}
		if _bfbc > _defba {
			return false
		}
		return _abca._gebec[_gdbeb]._fcbe < _abca._gebec[_gaac]._fcbe
	})
}

// SetColPosition sets cell column position.
func (_efbfg *TableCell) SetColPosition(col int) { _efbfg._fcbe = col }

// GetMargins returns the left, right, top, bottom Margins.
func (_fcedg *Table) GetMargins() (float64, float64, float64, float64) {
	return _fcedg._bebd.Left, _fcedg._bebd.Right, _fcedg._bebd.Top, _fcedg._bebd.Bottom
}

// SetRowPosition sets cell row position.
func (_bbag *TableCell) SetRowPosition(row int) { _bbag._bgdca = row }

// NewBlock creates a new Block with specified width and height.
func NewBlock(width float64, height float64) *Block {
	_fb := &Block{}
	_fb._fa = &_da.ContentStreamOperations{}
	_fb._ccf = _ga.NewPdfPageResources()
	_fb._fd = width
	_fb._cce = height
	return _fb
}

// CurvePolygon represents a curve polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type CurvePolygon struct {
	_efbeg *_ff.CurvePolygon
	_gfbf  float64
	_adad  float64
	_aged  Color
}

func (_bbegb *Table) resetColumnWidths() {
	_bbegb._eeggd = []float64{}
	_eabd := float64(1.0) / float64(_bbegb._edfe)
	for _aaac := 0; _aaac < _bbegb._edfe; _aaac++ {
		_bbegb._eeggd = append(_bbegb._eeggd, _eabd)
	}
}

// GeneratePageBlocks draws the block contents on a template Page block.
// Implements the Drawable interface.
func (_cde *Block) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gfa := _gb.IdentityMatrix()
	_ef, _ac := _cde.Width(), _cde.Height()
	if _cde._gfb.IsRelative() {
		_gfa = _gfa.Translate(ctx.X, ctx.PageHeight-ctx.Y-_ac)
	} else {
		_gfa = _gfa.Translate(_cde._bcde, ctx.PageHeight-_cde._gc-_ac)
	}
	_ceb := _ac
	if _cde._be != 0 {
		_gfa = _gfa.Translate(_ef/2, _ac/2).Rotate(_cde._be*_cd.Pi/180.0).Translate(-_ef/2, -_ac/2)
		_, _ceb = _cde.RotatedSize()
	}
	if _cde._gfb.IsRelative() {
		ctx.Y += _ceb
	}
	_bfc := _da.NewContentCreator()
	_bfc.Add_cm(_gfa[0], _gfa[1], _gfa[3], _gfa[4], _gfa[6], _gfa[7])
	_gfbe := _cde.duplicate()
	_beeb := append(*_bfc.Operations(), *_gfbe._fa...)
	_beeb.WrapIfNeeded()
	_gfbe._fa = &_beeb
	for _, _fac := range _cde._fdd {
		_cdeg, _beb := _cc.GetArray(_fac.Rect)
		if !_beb || _cdeg.Len() != 4 {
			_bcd.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0069\u006e\u0076\u0061\u006ci\u0064 \u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0052\u0065\u0063\u0074\u0020\u0066\u0069\u0065l\u0064\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _fac.Rect)
			continue
		}
		_fdb, _cda := _ga.NewPdfRectangle(*_cdeg)
		if _cda != nil {
			_bcd.Log.Debug("\u0057A\u0052N\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0070\u0061\u0072\u0073e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0052\u0065\u0063\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061y\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002e", _cda)
			continue
		}
		_fdb.Transform(_gfa)
		_fac.Rect = _fdb.ToPdfObject()
	}
	return []*Block{_gfbe}, ctx, nil
}

// Width is not used as the division component is designed to fill all the
// available space, depending on the context. Returns 0.
func (_aaee *Division) Width() float64 { return 0 }

func (_bdfb *templateProcessor) parseColor(_eabgb string) Color {
	if _eabgb == "" {
		return nil
	}
	_babd, _deab := _bdfb._agcb.ColorMap[_eabgb]
	if _deab {
		return _babd
	}
	if _eabgb[0] == '#' {
		return ColorRGBFromHex(_eabgb)
	}
	return nil
}

var _adg = _ab.MustCompile("\u005c\u0064\u002b")

// SetFillOpacity sets the fill opacity.
func (_fcbb *PolyBezierCurve) SetFillOpacity(opacity float64) { _fcbb._gbbfd = opacity }

// SetInline sets the inline mode of the division.
func (_cabd *Division) SetInline(inline bool) { _cabd._cdgb = inline }

// IsAbsolute checks if the positioning is absolute.
func (_dcacb Positioning) IsAbsolute() bool { return _dcacb == PositionAbsolute }

// The Image type is used to draw an image onto PDF.
type Image struct {
	_ebfc         *_ga.XObjectImage
	_dfad         *_ga.Image
	_egacc        float64
	_ggda, _cgdb  float64
	_gcde, _eebef float64
	_geca         Positioning
	_cccd         HorizontalAlignment
	_fdab         float64
	_egca         float64
	_gafe         float64
	_eaac         Margins
	_daeg, _dbfe  float64
	_cbce         _cc.StreamEncoder
	_dfdb         FitMode
}

func (_gecd *templateProcessor) renderNode(_dgab *templateNode) error {
	_babcg := _dgab._defae
	if _babcg == nil {
		return nil
	}
	_gedg := _dgab._acdf.Name.Local
	_dfeeda, _gfgf := _cabdda[_gedg]
	if !_gfgf {
		_gecd.nodeLogDebug(_dgab, "I\u006e\u0076\u0061\u006c\u0069\u0064 \u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e \u0053\u006b\u0069p\u0070i\u006e\u0067\u002e", _gedg)
		return nil
	}
	var _debe interface{}
	if _dgab._gbdf != nil && _dgab._gbdf._defae != nil {
		_gdfe := _dgab._gbdf._acdf.Name.Local
		if _, _gfgf = _dfeeda._cdcdf[_gdfe]; !_gfgf {
			_gecd.nodeLogDebug(_dgab, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e \u0069\u0073\u0020no\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u003c\u0025\u0073\u003e\u0020\u0074a\u0067\u002e", _gdfe, _gedg)
			return _bfegb
		}
		_debe = _dgab._gbdf._defae
	} else {
		_bgag := "\u0063r\u0065\u0061\u0074\u006f\u0072"
		switch _gecd._ebdg.(type) {
		case *Block:
			_bgag = "\u0062\u006c\u006fc\u006b"
		}
		if _, _gfgf = _dfeeda._cdcdf[_bgag]; !_gfgf {
			_gecd.nodeLogDebug(_dgab, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e \u0069\u0073\u0020no\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u003c\u0025\u0073\u003e\u0020\u0074a\u0067\u002e", _bgag, _gedg)
			return _bfegb
		}
		_debe = _gecd._ebdg
	}
	switch _fagad := _debe.(type) {
	case componentRenderer:
		_dggea, _adfbb := _babcg.(Drawable)
		if !_adfbb {
			_gecd.nodeLogError(_dgab, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020\u0069s\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u002e", _gedg, _babcg)
			return _bgdcf
		}
		_dbfgf := _fagad.Draw(_dggea)
		if _dbfgf != nil {
			return _gecd.nodeError(_dgab, "\u0043\u0061\u006en\u006f\u0074\u0020\u0064r\u0061\u0077\u0073\u0020\u0074\u0061\u0067 \u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u0073\u002e", _gedg, _babcg, _dbfgf)
		}
	case *Division:
		switch _fdbcff := _babcg.(type) {
		case *Background:
			_fagad.SetBackground(_fdbcff)
		case VectorDrawable:
			_affeb := _fagad.Add(_fdbcff)
			if _affeb != nil {
				return _gecd.nodeError(_dgab, "\u0043a\u006e\u006eo\u0074\u0020\u0061d\u0064\u0020\u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u0020\u0028\u0025\u0054)\u0020\u0069\u006e\u0074\u006f\u0020\u0061\u0020\u0044\u0069\u0076i\u0073\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u002e", _gedg, _babcg, _affeb)
			}
		}
	case *TableCell:
		_dggddc, _abgcf := _babcg.(VectorDrawable)
		if !_abgcf {
			_gecd.nodeLogError(_dgab, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029 \u0069\u0073\u0020\u006e\u006f\u0074 \u0061\u0020\u0076\u0065\u0063\u0074\u006f\u0072\u0020\u0064\u0072\u0061\u0077a\u0062\u006c\u0065\u002e", _gedg, _babcg)
			return _bgdcf
		}
		_affb := _fagad.SetContent(_dggddc)
		if _affb != nil {
			return _gecd.nodeError(_dgab, "C\u0061\u006e\u006e\u006f\u0074\u0020\u0061\u0064\u0064 \u0074\u0061\u0067\u0020\u003c\u0025\u0073> \u0028\u0025\u0054\u0029 \u0069\u006e\u0074\u006f\u0020\u0061\u0020\u0074\u0061bl\u0065\u0020c\u0065\u006c\u006c\u003a\u0020\u0025\u0073\u002e", _gedg, _babcg, _affb)
		}
	case *StyledParagraph:
		_aedgg, _ebfaa := _babcg.(*TextChunk)
		if !_ebfaa {
			_gecd.nodeLogError(_dgab, "\u0054\u0061\u0067 <\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020i\u0073 \u006eo\u0074 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u002e", _gedg, _babcg)
			return _bgdcf
		}
		_fagad.appendChunk(_aedgg)
	case *Chapter:
		switch _dgfc := _babcg.(type) {
		case *Chapter:
			return nil
		case *Paragraph:
			if _dgab._acdf.Name.Local == "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067" {
				return nil
			}
			_dedb := _fagad.Add(_dgfc)
			if _dedb != nil {
				return _gecd.nodeError(_dgab, "\u0043a\u006e\u006eo\u0074\u0020\u0061\u0064d\u0020\u0074\u0061g\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054) \u0069\u006e\u0074o\u0020\u0061 \u0043\u0068\u0061\u0070\u0074\u0065r\u003a\u0020%\u0073\u002e", _gedg, _babcg, _dedb)
			}
		case Drawable:
			_abcad := _fagad.Add(_dgfc)
			if _abcad != nil {
				return _gecd.nodeError(_dgab, "\u0043a\u006e\u006eo\u0074\u0020\u0061\u0064d\u0020\u0074\u0061g\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054) \u0069\u006e\u0074o\u0020\u0061 \u0043\u0068\u0061\u0070\u0074\u0065r\u003a\u0020%\u0073\u002e", _gedg, _babcg, _abcad)
			}
		}
	case *List:
		switch _caadf := _babcg.(type) {
		case *TextChunk:
		case *listItem:
			_fagad._bdccg = append(_fagad._bdccg, _caadf)
		default:
			_gecd.nodeLogError(_dgab, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073>\u0020\u0028\u0025T\u0029\u0020\u0069\u0073 \u006e\u006f\u0074\u0020\u0061\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u0074\u0065\u006d\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _gedg, _babcg)
		}
	case *listItem:
		switch _egde := _babcg.(type) {
		case *TextChunk:
		case *StyledParagraph:
			_fagad._cecd = _egde
		case *List:
			if _egde._ecee {
				_egde._addd = 15
			}
			_fagad._cecd = _egde
		case *Image:
			_fagad._cecd = _egde
		case *Division:
			_fagad._cecd = _egde
		case *Table:
			_fagad._cecd = _egde
		default:
			_gecd.nodeLogError(_dgab, "\u0054\u0061\u0067\u0020\u003c%\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u006c\u0069\u0073\u0074\u002e", _gedg, _babcg)
			return _bgdcf
		}
	}
	return nil
}

// ScaleToWidth sets the graphic svg scaling factor with the given width.
func (_gacd *GraphicSVG) ScaleToWidth(w float64) {
	_fecg := _gacd._debae.Height / _gacd._debae.Width
	_gacd._debae.Width = w
	_gacd._debae.Height = w * _fecg
	_gacd._debae.SetScaling(_fecg, _fecg)
}

// SetText replaces all the text of the paragraph with the specified one.
func (_aeacb *StyledParagraph) SetText(text string) *TextChunk {
	_aeacb.Reset()
	return _aeacb.Append(text)
}

type fontMetrics struct {
	_agaa  float64
	_bgggg float64
	_agdbg float64
	_ceddg float64
}

// SetBackgroundColor sets the cell's background color.
func (_aedg *TableCell) SetBackgroundColor(col Color) { _aedg._dedd = col }

// GetCoords returns coordinates of border.
func (_ceg *border) GetCoords() (float64, float64) { return _ceg._cbd, _ceg._bcgdb }

// SetWidthRight sets border width for right.
func (_baa *border) SetWidthRight(bw float64) { _baa._dggc = bw }

// CreateFrontPage sets a function to generate a front Page.
func (_eaee *Creator) CreateFrontPage(genFrontPageFunc func(_dfda FrontpageFunctionArgs)) {
	_eaee._cge = genFrontPageFunc
}

// FillColor returns the fill color of the ellipse.
func (_gfdf *Ellipse) FillColor() Color { return _gfdf._fecae }

// SetColumns overwrites any columns in the line items table. This should be
// called before AddLine.
func (_acac *Invoice) SetColumns(cols []*InvoiceCell) { _acac._cdfg = cols }

// FooterFunctionArgs holds the input arguments to a footer drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FooterFunctionArgs struct {
	PageNum    int
	TotalPages int
}

func _acfdb(_accb *templateProcessor, _gcfd *templateNode) (interface{}, error) {
	return _accb.parseRectangle(_gcfd)
}

// Add adds a VectorDrawable to the Division container.
// Currently supported VectorDrawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
// - *Rectangle
// - *Ellipse
// - *Line
// - *Table
// - *Division
// - *List
func (_dgefg *Division) Add(d VectorDrawable) error {
	switch _edfae := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Rectangle, *Ellipse, *Line, *Table, *Division, *List:
	case containerDrawable:
		_ddgd, _gdca := _edfae.ContainerComponent(_dgefg)
		if _gdca != nil {
			return _gdca
		}
		_fccb, _edab := _ddgd.(VectorDrawable)
		if !_edab {
			return _g.Errorf("\u0072\u0065\u0073\u0075\u006ct\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u002d\u0020\u0025\u0054\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0056\u0065c\u0074\u006f\u0072\u0044\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020i\u006e\u0074\u0065\u0072\u0066\u0061c\u0065", _ddgd)
		}
		d = _fccb
	default:
		return _c.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0044\u0069\u0076\u0069\u0073i\u006f\u006e")
	}
	_dgefg._dccd = append(_dgefg._dccd, d)
	return nil
}

// SetBorderColor sets the border color.
func (_bdbf *CurvePolygon) SetBorderColor(color Color) { _bdbf._efbeg.BorderColor = _fce(color) }

func (_aefcf *Image) rotatedSize() (float64, float64) {
	_fgea := _aefcf._ggda
	_dcgd := _aefcf._cgdb
	_bdcbb := _aefcf._egacc
	if _bdcbb == 0 {
		return _fgea, _dcgd
	}
	_faggc := _ff.Path{Points: []_ff.Point{_ff.NewPoint(0, 0).Rotate(_bdcbb), _ff.NewPoint(_fgea, 0).Rotate(_bdcbb), _ff.NewPoint(0, _dcgd).Rotate(_bdcbb), _ff.NewPoint(_fgea, _dcgd).Rotate(_bdcbb)}}.GetBoundingBox()
	return _faggc.Width, _faggc.Height
}

type templateProcessor struct {
	creator *Creator
	_cebag  []byte
	_agcb   *TemplateOptions
	_ebdg   componentRenderer
	_bfgd   string
}

// Length calculates and returns the length of the line.
func (_bdcba *Line) Length() float64 {
	return _cd.Sqrt(_cd.Pow(_bdcba._dfcae-_bdcba._ccg, 2.0) + _cd.Pow(_bdcba._feafb-_bdcba._eggf, 2.0))
}

// SetEnableWrap sets the line wrapping enabled flag.
func (_dcegd *Paragraph) SetEnableWrap(enableWrap bool) {
	_dcegd._dbfb = enableWrap
	_dcegd._deaf = false
}

func (_dcfe *templateProcessor) parseChart(_dcaf *templateNode) (interface{}, error) {
	var _bbac string
	for _, _cdegf := range _dcaf._acdf.Attr {
		_geag := _cdegf.Value
		switch _dcfdc := _cdegf.Name.Local; _dcfdc {
		case "\u0073\u0072\u0063":
			_bbac = _geag
		}
	}
	if _bbac == "" {
		_dcfe.nodeLogError(_dcaf, "\u0043\u0068\u0061\u0072\u0074\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _efbef
	}
	_bdff, _bffaa := _dcfe._agcb.ChartMap[_bbac]
	if !_bffaa {
		_dcfe.nodeLogError(_dcaf, "\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0063ha\u0072\u0074\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _bbac)
		return nil, _efbef
	}
	_dfebg := NewChart(_bdff)
	for _, _ggdfc := range _dcaf._acdf.Attr {
		_eeccd := _ggdfc.Value
		switch _effcc := _ggdfc.Name.Local; _effcc {
		case "\u0078":
			_dfebg.SetPos(_dcfe.parseFloatAttr(_effcc, _eeccd), _dfebg._dcd)
		case "\u0079":
			_dfebg.SetPos(_dfebg._cdd, _dcfe.parseFloatAttr(_effcc, _eeccd))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_dgag := _dcfe.parseMarginAttr(_effcc, _eeccd)
			_dfebg.SetMargins(_dgag.Left, _dgag.Right, _dgag.Top, _dgag.Bottom)
		case "\u0077\u0069\u0064t\u0068":
			_dfebg._dece.SetWidth(int(_dcfe.parseFloatAttr(_effcc, _eeccd)))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_dfebg._dece.SetHeight(int(_dcfe.parseFloatAttr(_effcc, _eeccd)))
		case "\u0073\u0072\u0063":
			break
		default:
			_dcfe.nodeLogDebug(_dcaf, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0063h\u0061\u0072\u0074\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _effcc)
		}
	}
	return _dfebg, nil
}

// GetMargins returns the margins of the rectangle: left, right, top, bottom.
func (_bbgbf *Rectangle) GetMargins() (float64, float64, float64, float64) {
	return _bbgbf._eaaa.Left, _bbgbf._eaaa.Right, _bbgbf._eaaa.Top, _bbgbf._eaaa.Bottom
}

func _ffgbb(_gdad *templateProcessor, _ageb *templateNode) (interface{}, error) {
	return _gdad.parseDivision(_ageb)
}

// AppendColumn appends a column to the line items table.
func (_bfgf *Invoice) AppendColumn(description string) *InvoiceCell {
	_fgag := _bfgf.NewColumn(description)
	_bfgf._cdfg = append(_bfgf._cdfg, _fgag)
	return _fgag
}

func _ffcbg(_abbg *templateProcessor, _bcfa *templateNode) (interface{}, error) {
	return _abbg.parseListItem(_bcfa)
}

// SetFillOpacity sets the fill opacity.
func (_ffeb *Polygon) SetFillOpacity(opacity float64) { _ffeb._afdbc = opacity }

// SetLineWidth sets the line width.
func (_facb *Polyline) SetLineWidth(lineWidth float64) { _facb._agbgf.LineWidth = lineWidth }

// Add adds a new Drawable to the chapter.
// Currently supported Drawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
// - *Table
// - *Division
// - *List
// - *Rectangle
// - *Ellipse
// - *Line
// - *Block,
// - *PageBreak
// - *Chapter
func (_bbcf *Chapter) Add(d Drawable) error {
	if Drawable(_bbcf) == d {
		_bcd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0061\u0064\u0064\u0020\u0069\u0074\u0073\u0065\u006c\u0066")
		return _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	switch _dab := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line, *Block, *PageBreak, *Chapter:
		_bbcf._gefd = append(_bbcf._gefd, d)
	case containerDrawable:
		_bdc, _fabb := _dab.ContainerComponent(_bbcf)
		if _fabb != nil {
			return _fabb
		}
		_bbcf._gefd = append(_bbcf._gefd, _bdc)
	default:
		_bcd.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u003a\u0020\u0025\u0054", d)
		return _c.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return nil
}

func (_cabaa *List) markerWidth() float64 {
	var _efeae float64
	for _, _aead := range _cabaa._bdccg {
		_abbb := _cgfa(_cabaa._edda)
		_abbb.SetEnableWrap(false)
		_abbb.SetTextAlignment(TextAlignmentRight)
		_abbb.Append(_aead._feced.Text).Style = _aead._feced.Style
		_bgaaba := _abbb.getTextWidth() / 1000.0
		if _efeae < _bgaaba {
			_efeae = _bgaaba
		}
	}
	return _efeae
}

// NewPolyline creates a new polyline.
func (_dgfff *Creator) NewPolyline(points []_ff.Point) *Polyline { return _bfba(points) }

func _dcabd(_dcag, _dbed, _dcfd, _begc, _aeed, _fcfb float64) *Curve {
	_bgb := &Curve{}
	_bgb._caec = _dcag
	_bgb._fcdb = _dbed
	_bgb._fcg = _dcfd
	_bgb._cfdd = _begc
	_bgb._fcfe = _aeed
	_bgb._cbf = _fcfb
	_bgb._gag = ColorBlack
	_bgb._cdaf = 1.0
	return _bgb
}

// ConvertToBinary converts current image data into binary (Bi-level image) format.
// If provided image is RGB or GrayScale the function converts it into binary image
// using histogram auto threshold method.
func (_eaca *Image) ConvertToBinary() error { return _eaca._dfad.ConvertToBinary() }

func (_geeb *Creator) getActivePage() *_ga.PdfPage {
	if _geeb._abd == nil {
		if len(_geeb._agfg) == 0 {
			return nil
		}
		return _geeb._agfg[len(_geeb._agfg)-1]
	}
	return _geeb._abd
}

// Width returns the width of the specified text chunk.
func (_bccf *TextChunk) Width() float64 {
	var (
		_ggdc float64
		_cafe = _bccf.Style
	)
	for _, _bdfbd := range _bccf.Text {
		_cdbef, _fadcb := _cafe.Font.GetRuneMetrics(_bdfbd)
		if !_fadcb {
			_bcd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _bdfbd, _bdfbd, _cafe.Font.BaseFont(), _cafe.Font.Subtype())
			_bcd.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _cafe.Font)
			_bcd.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _cafe.Font.Encoder())
		}
		_cbdbe := _cafe.FontSize * _cdbef.Wx
		_dacd := _cbdbe
		if _bdfbd != ' ' {
			_dacd = _cbdbe + _cafe.CharSpacing*1000.0
		}
		_ggdc += _dacd
	}
	return _ggdc / 1000.0
}

// SetLineHeight sets the line height (1.0 default).
func (_fcdcg *Paragraph) SetLineHeight(lineheight float64) { _fcdcg._agdb = lineheight }

// AddressStyle returns the style properties used to render the content of
// the invoice address sections.
func (_afda *Invoice) AddressStyle() TextStyle { return _afda._gaae }

func (_adcda *InvoiceAddress) fmtLine(_effa, _egfe string, _bebbe bool) string {
	if _bebbe {
		_egfe = ""
	}
	return _g.Sprintf("\u0025\u0073\u0025s\u000a", _egfe, _effa)
}

func _fgeeg(_bdfd *templateProcessor, _cfadf *templateNode) (interface{}, error) {
	return _bdfd.parseTable(_cfadf)
}

// Cols returns the total number of columns the table has.
func (_fcgff *Table) Cols() int { return _fcgff._edfe }

// Polyline represents a slice of points that are connected as straight lines.
// Implements the Drawable interface and can be rendered using the Creator.
type Polyline struct {
	_agbgf *_ff.Polyline
	_cbbd  float64
}

// SetMargins sets the Paragraph's margins.
func (_gdac *StyledParagraph) SetMargins(left, right, top, bottom float64) {
	_gdac._fadcg.Left = left
	_gdac._fadcg.Right = right
	_gdac._fadcg.Top = top
	_gdac._fadcg.Bottom = bottom
}

func _cgd(_bfb *_da.ContentStreamOperations, _cee *_ga.PdfPageResources, _gfaa *_da.ContentStreamOperations, _dec *_ga.PdfPageResources) error {
	_acc := map[_cc.PdfObjectName]_cc.PdfObjectName{}
	_ece := map[_cc.PdfObjectName]_cc.PdfObjectName{}
	_addg := map[_cc.PdfObjectName]_cc.PdfObjectName{}
	_aee := map[_cc.PdfObjectName]_cc.PdfObjectName{}
	_cdf := map[_cc.PdfObjectName]_cc.PdfObjectName{}
	_cfc := map[_cc.PdfObjectName]_cc.PdfObjectName{}
	for _, _abg := range *_gfaa {
		switch _abg.Operand {
		case "\u0044\u006f":
			if len(_abg.Params) == 1 {
				if _fe, _bb := _abg.Params[0].(*_cc.PdfObjectName); _bb {
					if _, _aecg := _acc[*_fe]; !_aecg {
						var _gba _cc.PdfObjectName
						_ddab, _ := _dec.GetXObjectByName(*_fe)
						if _ddab != nil {
							_gba = *_fe
							for {
								_eee, _ := _cee.GetXObjectByName(_gba)
								if _eee == nil || _eee == _ddab {
									break
								}
								_gba = *_cc.MakeName(_gcd(_gba.String()))
							}
						}
						_cee.SetXObjectByName(_gba, _ddab)
						_acc[*_fe] = _gba
					}
					_acf := _acc[*_fe]
					_abg.Params[0] = &_acf
				}
			}
		case "\u0054\u0066":
			if len(_abg.Params) == 2 {
				if _ebgc, _bef := _abg.Params[0].(*_cc.PdfObjectName); _bef {
					if _, _deca := _ece[*_ebgc]; !_deca {
						_dge, _cac := _dec.GetFontByName(*_ebgc)
						_gd := *_ebgc
						if _cac && _dge != nil {
							_gd = _cgf(_ebgc.String(), _dge, _cee)
						}
						_cee.SetFontByName(_gd, _dge)
						_ece[*_ebgc] = _gd
					}
					_bgdg := _ece[*_ebgc]
					_abg.Params[0] = &_bgdg
				}
			}
		case "\u0043\u0053", "\u0063\u0073":
			if len(_abg.Params) == 1 {
				if _ecb, _beed := _abg.Params[0].(*_cc.PdfObjectName); _beed {
					if _, _bd := _addg[*_ecb]; !_bd {
						var _dfge _cc.PdfObjectName
						_aae, _agb := _dec.GetColorspaceByName(*_ecb)
						if _agb {
							_dfge = *_ecb
							for {
								_ecf, _efa := _cee.GetColorspaceByName(_dfge)
								if !_efa || _aae == _ecf {
									break
								}
								_dfge = *_cc.MakeName(_gcd(_dfge.String()))
							}
							_cee.SetColorspaceByName(_dfge, _aae)
							_addg[*_ecb] = _dfge
						} else {
							_bcd.Log.Debug("C\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064")
						}
					}
					if _cabb, _ege := _addg[*_ecb]; _ege {
						_abg.Params[0] = &_cabb
					} else {
						_bcd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", *_ecb)
					}
				}
			}
		case "\u0053\u0043\u004e", "\u0073\u0063\u006e":
			if len(_abg.Params) == 1 {
				if _dgf, _gde := _abg.Params[0].(*_cc.PdfObjectName); _gde {
					if _, _bdb := _aee[*_dgf]; !_bdb {
						var _fba _cc.PdfObjectName
						_fgc, _fbe := _dec.GetPatternByName(*_dgf)
						if _fbe {
							_fba = *_dgf
							for {
								_aecd, _dgc := _cee.GetPatternByName(_fba)
								if !_dgc || _aecd == _fgc {
									break
								}
								_fba = *_cc.MakeName(_gcd(_fba.String()))
							}
							_edcb := _cee.SetPatternByName(_fba, _fgc.ToPdfObject())
							if _edcb != nil {
								return _edcb
							}
							_aee[*_dgf] = _fba
						}
					}
					if _dcfg, _dggd := _aee[*_dgf]; _dggd {
						_abg.Params[0] = &_dcfg
					}
				}
			}
		case "\u0073\u0068":
			if len(_abg.Params) == 1 {
				if _edfa, _bad := _abg.Params[0].(*_cc.PdfObjectName); _bad {
					if _, _cdc := _cdf[*_edfa]; !_cdc {
						var _cfb _cc.PdfObjectName
						_gef, _afa := _dec.GetShadingByName(*_edfa)
						if _afa {
							_cfb = *_edfa
							for {
								_bgf, _geg := _cee.GetShadingByName(_cfb)
								if !_geg || _gef == _bgf {
									break
								}
								_cfb = *_cc.MakeName(_gcd(_cfb.String()))
							}
							_bbc := _cee.SetShadingByName(_cfb, _gef.ToPdfObject())
							if _bbc != nil {
								_bcd.Log.Debug("E\u0052\u0052\u004f\u0052 S\u0065t\u0020\u0073\u0068\u0061\u0064i\u006e\u0067\u003a\u0020\u0025\u0076", _bbc)
								return _bbc
							}
							_cdf[*_edfa] = _cfb
						} else {
							_bcd.Log.Debug("\u0053\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
						}
					}
					if _bbcb, _fbd := _cdf[*_edfa]; _fbd {
						_abg.Params[0] = &_bbcb
					} else {
						_bcd.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020S\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0073 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064", *_edfa)
					}
				}
			}
		case "\u0067\u0073":
			if len(_abg.Params) == 1 {
				if _ced, _eeb := _abg.Params[0].(*_cc.PdfObjectName); _eeb {
					if _, _ded := _cfc[*_ced]; !_ded {
						var _eca _cc.PdfObjectName
						_ecad, _cdab := _dec.GetExtGState(*_ced)
						if _cdab {
							_eca = *_ced
							for {
								_cded, _aea := _cee.GetExtGState(_eca)
								if !_aea || _ecad == _cded {
									break
								}
								_eca = *_cc.MakeName(_gcd(_eca.String()))
							}
						}
						_cee.AddExtGState(_eca, _ecad)
						_cfc[*_ced] = _eca
					}
					_aac := _cfc[*_ced]
					_abg.Params[0] = &_aac
				}
			}
		}
		*_bfb = append(*_bfb, _abg)
	}
	return nil
}

func _agbcg(_ggfc, _cddc, _bcge, _bfcd float64) *Line {
	return &Line{_ccg: _ggfc, _eggf: _cddc, _dfcae: _bcge, _feafb: _bfcd, _efdf: ColorBlack, _gbcb: 1.0, _bfe: 1.0, _fedd: []int64{1, 1}, _fcdd: PositionAbsolute}
}

// SetMargins sets the margins of the line.
// NOTE: line margins are only applied if relative positioning is used.
func (_gadec *Line) SetMargins(left, right, top, bottom float64) {
	_gadec._gabb.Left = left
	_gadec._gabb.Right = right
	_gadec._gabb.Top = top
	_gadec._gabb.Bottom = bottom
}

var (
	_cddbe = _ab.MustCompile("\u0028[\u005cw\u002d\u005d\u002b\u0029\u005c(\u0027\u0028.\u002b\u0029\u0027\u005c\u0029")
	_caca  = _c.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0063\u0072\u0065a\u0074\u006f\u0072\u0020\u0069\u006e\u0073t\u0061\u006e\u0063\u0065")
	_bfegb = _c.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074e\u0020p\u0061\u0072\u0065\u006e\u0074\u0020\u006eo\u0064\u0065")
	_bgdcf = _c.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020c\u0068\u0069\u006cd\u0020n\u006f\u0064\u0065")
	_efbef = _c.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070l\u0061t\u0065 \u0072\u0065\u0073\u006f\u0075\u0072\u0063e")
)

// Horizontal returns total horizontal (left + right) margin.
func (_bdcd *Margins) Horizontal() float64 { return _bdcd.Left + _bdcd.Right }

func (_edfeb *TemplateOptions) init() {
	if _edfeb.SubtemplateMap == nil {
		_edfeb.SubtemplateMap = map[string]_eb.Reader{}
	}
	if _edfeb.FontMap == nil {
		_edfeb.FontMap = map[string]*_ga.PdfFont{}
	}
	if _edfeb.ImageMap == nil {
		_edfeb.ImageMap = map[string]*_ga.Image{}
	}
	if _edfeb.ColorMap == nil {
		_edfeb.ColorMap = map[string]Color{}
	}
	if _edfeb.ChartMap == nil {
		_edfeb.ChartMap = map[string]_dce.ChartRenderable{}
	}
}

// SetSideBorderColor sets the cell's side border color.
func (_fagda *TableCell) SetSideBorderColor(side CellBorderSide, col Color) {
	switch side {
	case CellBorderSideAll:
		_fagda._afddf = col
		_fagda._ebgec = col
		_fagda._efcg = col
		_fagda._fcec = col
	case CellBorderSideTop:
		_fagda._afddf = col
	case CellBorderSideBottom:
		_fagda._ebgec = col
	case CellBorderSideLeft:
		_fagda._efcg = col
	case CellBorderSideRight:
		_fagda._fcec = col
	}
}

// SetPos sets the Table's positioning to absolute mode and specifies the upper-left corner
// coordinates as (x,y).
// Note that this is only sensible to use when the table does not wrap over multiple pages.
// TODO: Should be able to set width too (not just based on context/relative positioning mode).
func (_fefg *Table) SetPos(x, y float64) {
	_fefg._efce = PositionAbsolute
	_fefg._aabe = x
	_fefg._bfbac = y
}

// NewLinearGradientColor creates a linear gradient color that could act as a color in other components.
func (_bcgf *Creator) NewLinearGradientColor(colorPoints []*ColorPoint) *LinearShading {
	return _eega(colorPoints)
}

// SetPositioning sets the positioning of the rectangle (absolute or relative).
func (_acfbf *Rectangle) SetPositioning(position Positioning) { _acfbf._dgde = position }

func (_caac *pageTransformations) transformBlock(_cfce *Block) {
	if _caac._aefc != nil {
		_cfce.transform(*_caac._aefc)
	}
}

// SetLineOpacity sets the line opacity.
func (_fbeea *Polyline) SetLineOpacity(opacity float64) { _fbeea._cbbd = opacity }

// SkipOver skips over a specified number of rows and cols.
func (_feac *Table) SkipOver(rows, cols int) {
	_faedb := rows*_feac._edfe + cols - 1
	if _faedb < 0 {
		_bcd.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _fabac := 0; _fabac < _faedb; _fabac++ {
		_feac.NewCell()
	}
}

// Chapter is used to arrange multiple drawables (paragraphs, images, etc) into a single section.
// The concept is the same as a book or a report chapter.
type Chapter struct {
	_baef        int
	_gbb         string
	_bgfc        *Paragraph
	_gefd        []Drawable
	_gcf         int
	_ebgde       bool
	_gge         bool
	_fdf         Positioning
	_gabc, _dcee float64
	_ecd         Margins
	_eagg        *Chapter
	_aaccb       *TOC
	_eebd        *_ga.Outline
	_fddf        *_ga.OutlineItem
	_fge         uint
}

// Total returns the invoice total description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_dfgad *Invoice) Total() (*InvoiceCell, *InvoiceCell) { return _dfgad._eacf[0], _dfgad._eacf[1] }

// Width returns the Block's width.
func (_cb *Block) Width() float64 { return _cb._fd }

func (_fabgc *pageTransformations) transformPage(_faff *_ga.PdfPage) error {
	if _fagec := _fabgc.applyFlip(_faff); _fagec != nil {
		return _fagec
	}
	return nil
}

// DrawContext defines the drawing context. The DrawContext is continuously used and updated when
// drawing the page contents in relative mode.  Keeps track of current X, Y position, available
// height as well as other page parameters such as margins and dimensions.
type DrawContext struct {
	// Current page number.
	Page int

	// Current position.  In a relative positioning mode, a drawable will be placed at these coordinates.
	X, Y float64

	// Context dimensions.  Available width and height (on current page).
	Width, Height float64

	// Page Margins.
	Margins Margins

	// Absolute Page size, widths and height.
	PageWidth  float64
	PageHeight float64

	// Controls whether the components are stacked horizontally
	Inline bool
	_dfe   rune
	_ddcc  []error
}

// NewImageFromGoImage creates an Image from a go image.Image data structure.
func (_agfd *Creator) NewImageFromGoImage(goimg _ba.Image) (*Image, error) { return _aafcf(goimg) }

func (_afcg *LinearShading) shadingModel() *_ga.PdfShadingType2 {
	_bdbc := _ff.NewPoint(_afcg._ebfea.Llx+_afcg._ebfea.Width()/2, _afcg._ebfea.Lly+_afcg._ebfea.Height()/2)
	_ggfcg := _ff.NewPoint(_afcg._ebfea.Llx, _afcg._ebfea.Lly+_afcg._ebfea.Height()/2).Add(-_bdbc.X, -_bdbc.Y).Rotate(_afcg._dgdad).Add(_bdbc.X, _bdbc.Y)
	_ggfcg = _ff.NewPoint(_cd.Max(_cd.Min(_ggfcg.X, _afcg._ebfea.Urx), _afcg._ebfea.Llx), _cd.Max(_cd.Min(_ggfcg.Y, _afcg._ebfea.Ury), _afcg._ebfea.Lly))
	_cbff := _ff.NewPoint(_afcg._ebfea.Urx, _afcg._ebfea.Lly+_afcg._ebfea.Height()/2).Add(-_bdbc.X, -_bdbc.Y).Rotate(_afcg._dgdad).Add(_bdbc.X, _bdbc.Y)
	_cbff = _ff.NewPoint(_cd.Min(_cd.Max(_cbff.X, _afcg._ebfea.Llx), _afcg._ebfea.Urx), _cd.Min(_cd.Max(_cbff.Y, _afcg._ebfea.Lly), _afcg._ebfea.Ury))
	_fegb := _ga.NewPdfShadingType2()
	_fegb.PdfShading.ShadingType = _cc.MakeInteger(2)
	_fegb.PdfShading.ColorSpace = _ga.NewPdfColorspaceDeviceRGB()
	_fegb.PdfShading.AntiAlias = _cc.MakeBool(_afcg._cfbfe._cebg)
	_fegb.Coords = _cc.MakeArrayFromFloats([]float64{_ggfcg.X, _ggfcg.Y, _cbff.X, _cbff.Y})
	_fegb.Extend = _cc.MakeArray(_cc.MakeBool(_afcg._cfbfe._fedf[0]), _cc.MakeBool(_afcg._cfbfe._fedf[1]))
	_fegb.Function = _afcg._cfbfe.generatePdfFunctions()
	return _fegb
}

// FitMode returns the fit mode of the image.
func (_fcdc *Image) FitMode() FitMode { return _fcdc._dfdb }

// SetHeaderRows turns the selected table rows into headers that are repeated
// for every page the table spans. startRow and endRow are inclusive.
func (_gcace *Table) SetHeaderRows(startRow, endRow int) error {
	if startRow <= 0 {
		return _c.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020r\u006f\u0077\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if endRow <= 0 {
		return _c.New("\u0068\u0065a\u0064\u0065\u0072\u0020e\u006e\u0064 \u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if startRow > endRow {
		return _c.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0072\u006f\u0077\u0020\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065 \u0065\u006e\u0064\u0020\u0072o\u0077")
	}
	_gcace._gdaae = true
	_gcace._gcdff = startRow
	_gcace._bgec = endRow
	return nil
}

func _eedg(_acaae *templateProcessor, _cgdec *templateNode) (interface{}, error) {
	return _acaae.parseEllipse(_cgdec)
}

func (_ffggf *templateProcessor) parseAttrPropList(_ecda string) map[string]string {
	_aceeg := _a.Fields(_ecda)
	if len(_aceeg) == 0 {
		return nil
	}
	_cfgfgg := map[string]string{}
	for _, _aebeg := range _aceeg {
		_ccbd := _cddbe.FindStringSubmatch(_aebeg)
		if len(_ccbd) < 3 {
			continue
		}
		_ffef, _gcag := _a.TrimSpace(_ccbd[1]), _ccbd[2]
		if _ffef == "" {
			continue
		}
		_cfgfgg[_ffef] = _gcag
	}
	return _cfgfgg
}

// SetFitMode sets the fit mode of the image.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_cfbf *Image) SetFitMode(fitMode FitMode) { _cfbf._dfdb = fitMode }

func _cfcc(_ddccf *_ga.Image) (*Image, error) {
	_ggef := float64(_ddccf.Width)
	_dfab := float64(_ddccf.Height)
	return &Image{_dfad: _ddccf, _gcde: _ggef, _eebef: _dfab, _ggda: _ggef, _cgdb: _dfab, _egacc: 0, _gafe: 1.0, _geca: PositionRelative}, nil
}

func (_dagc *StyledParagraph) getLineMetrics(_cfgfg int) (_bggc, _ddcfg, _bcfga float64) {
	if _dagc._bddb == nil || len(_dagc._bddb) == 0 {
		_dagc.wrapText()
	}
	if _cfgfg < 0 || _cfgfg > len(_dagc._bddb)-1 {
		_bcd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020p\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u006c\u0069\u006e\u0065 \u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002e\u0020\u0052\u0065tu\u0072\u006e\u0069\u006e\u0067\u0020\u0030\u002c\u0020\u0030", _cfgfg)
		return 0, 0, 0
	}
	_accae := _dagc._bddb[_cfgfg]
	for _, _gagd := range _accae {
		_dffb := _abff(_gagd.Style.Font, _gagd.Style.FontSize)
		if _dffb._agaa > _bggc {
			_bggc = _dffb._agaa
		}
		if _dffb._ceddg < _bcfga {
			_bcfga = _dffb._ceddg
		}
		if _ceee := _gagd.Style.FontSize; _ceee > _ddcfg {
			_ddcfg = _ceee
		}
	}
	return _bggc, _ddcfg, _bcfga
}

func (_dcef *Block) setOpacity(_ce float64, _fda float64) (string, error) {
	if (_ce < 0 || _ce >= 1.0) && (_fda < 0 || _fda >= 1.0) {
		return "", nil
	}
	_ec := 0
	_af := _g.Sprintf("\u0047\u0053\u0025\u0064", _ec)
	for _dcef._ccf.HasExtGState(_cc.PdfObjectName(_af)) {
		_ec++
		_af = _g.Sprintf("\u0047\u0053\u0025\u0064", _ec)
	}
	_ag := _cc.MakeDict()
	if _ce >= 0 && _ce < 1.0 {
		_ag.Set("\u0063\u0061", _cc.MakeFloat(_ce))
	}
	if _fda >= 0 && _fda < 1.0 {
		_ag.Set("\u0043\u0041", _cc.MakeFloat(_fda))
	}
	_bae := _dcef._ccf.AddExtGState(_cc.PdfObjectName(_af), _ag)
	if _bae != nil {
		return "", _bae
	}
	return _af, nil
}

// SetBorderRadius sets the radius of the rectangle corners.
func (_faab *Rectangle) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_faab._acgb = topLeft
	_faab._cgea = topRight
	_faab._bfag = bottomLeft
	_faab._ebbe = bottomRight
}

// NewLine creates a new line between (x1, y1) to (x2, y2),
// using default attributes.
// NOTE: In relative positioning mode, `x1` and `y1` are calculated using the
// current context and `x2`, `y2` are used only to calculate the position of
// the second point in relation to the first one (used just as a measurement
// of size). Furthermore, when the fit mode is set to fill the context width,
// `x2` is set to the right edge coordinate of the context.
func (_dgaf *Creator) NewLine(x1, y1, x2, y2 float64) *Line { return _agbcg(x1, y1, x2, y2) }

// Height returns the height of the list.
func (_dbge *List) Height() float64 {
	var _dgda float64
	for _, _adcc := range _dbge._bdccg {
		_dgda += _adcc.ctxHeight(_dbge.Width())
	}
	return _dgda
}

// SetBoundingBox set gradient color bounding box where the gradient would be rendered.
func (_fffc *RadialShading) SetBoundingBox(x, y, width, height float64) {
	_fffc._bcgga = &_ga.PdfRectangle{Llx: x, Lly: y, Urx: x + width, Ury: y + height}
}

func (_fddd *Creator) initContext() {
	_fddd._ggab.X = _fddd._eaf.Left
	_fddd._ggab.Y = _fddd._eaf.Top
	_fddd._ggab.Width = _fddd._bagf - _fddd._eaf.Right - _fddd._eaf.Left
	_fddd._ggab.Height = _fddd._aafc - _fddd._eaf.Bottom - _fddd._eaf.Top
	_fddd._ggab.PageHeight = _fddd._aafc
	_fddd._ggab.PageWidth = _fddd._bagf
	_fddd._ggab.Margins = _fddd._eaf
	_fddd._ggab._dfe = _fddd.UnsupportedCharacterReplacement
}

func (_daabg *templateProcessor) processGradientColorPair(_bgbdc []string) (_gefga []Color, _fgcdb []float64) {
	for _, _bfbb := range _bgbdc {
		var (
			_ebdcf = _a.Fields(_bfbb)
			_eabdc = len(_ebdcf)
		)
		if _eabdc == 0 {
			continue
		}
		_cgded := ""
		if _eabdc > 1 {
			_cgded = _a.TrimSpace(_ebdcf[1])
		}
		_dfba := -1.0
		if _a.HasSuffix(_cgded, "\u0025") {
			_caadd, _cdaec := _aa.ParseFloat(_cgded[:len(_cgded)-1], 64)
			if _cdaec != nil {
				_bcd.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0070\u006f\u0069n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _cdaec)
			}
			_dfba = _caadd / 100.0
		}
		_ddeaa := _daabg.parseColor(_a.TrimSpace(_ebdcf[0]))
		if _ddeaa != nil {
			_gefga = append(_gefga, _ddeaa)
			_fgcdb = append(_fgcdb, _dfba)
		}
	}
	if len(_gefga) != len(_fgcdb) {
		_bcd.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u006c\u0069\u006e\u0065\u0061\u0072\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0064\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0021")
		return nil, nil
	}
	_eeggc := -1
	_gcfff := 0.0
	for _dcba, _bagca := range _fgcdb {
		if _bagca == -1.0 {
			if _dcba == 0 {
				_bagca = 0.0
				_fgcdb[_dcba] = 0.0
				continue
			}
			_eeggc++
			if _dcba < len(_fgcdb)-1 {
				continue
			} else {
				_bagca = 1.0
				_fgcdb[_dcba] = 1.0
			}
		}
		_dggcg := _eeggc + 1
		for _bdgd := _dcba - _eeggc; _bdgd < _dcba; _bdgd++ {
			_fgcdb[_bdgd] = _gcfff + (float64(_bdgd) * (_bagca - _gcfff) / float64(_dggcg))
		}
		_gcfff = _bagca
		_eeggc = -1
	}
	return _gefga, _fgcdb
}

// SetNumber sets the number of the invoice.
func (_efgc *Invoice) SetNumber(number string) (*InvoiceCell, *InvoiceCell) {
	_efgc._badb[1].Value = number
	return _efgc._badb[0], _efgc._badb[1]
}

// AddInfo is used to append a piece of invoice information in the template
// information table.
func (_edefg *Invoice) AddInfo(description, value string) (*InvoiceCell, *InvoiceCell) {
	_defb := [2]*InvoiceCell{_edefg.newCell(description, _edefg._eeab), _edefg.newCell(value, _edefg._eeab)}
	_edefg._gfda = append(_edefg._gfda, _defb)
	return _defb[0], _defb[1]
}

func (_dceg *Block) transform(_ccfc _gb.Matrix) {
	_edg := _da.NewContentCreator().Add_cm(_ccfc[0], _ccfc[1], _ccfc[3], _ccfc[4], _ccfc[6], _ccfc[7]).Operations()
	*_dceg._fa = append(*_edg, *_dceg._fa...)
	_dceg._fa.WrapIfNeeded()
}

// SetLineMargins sets the margins for all new lines of the table of contents.
func (_aaea *TOC) SetLineMargins(left, right, top, bottom float64) {
	_cacf := &_aaea._ffcde
	_cacf.Left = left
	_cacf.Right = right
	_cacf.Top = top
	_cacf.Bottom = bottom
}

// BorderColor returns the border color of the rectangle.
func (_ffbe *Rectangle) BorderColor() Color { return _ffbe._gggb }

// SetAngle sets the rotation angle in degrees.
func (_cadg *Block) SetAngle(angleDeg float64) { _cadg._be = angleDeg }

// GeneratePageBlocks draws the ellipse on a new block representing the page.
func (_acdcd *Ellipse) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_gdag []*Block
		_ddeb = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_agfc = ctx
	)
	_gadc := _acdcd._fabd.IsRelative()
	if _gadc {
		_acdcd.applyFitMode(ctx.Width)
		ctx.X += _acdcd._abgc.Left
		ctx.Y += _acdcd._abgc.Top
		ctx.Width -= _acdcd._abgc.Left + _acdcd._abgc.Right
		ctx.Height -= _acdcd._abgc.Top + _acdcd._abgc.Bottom
		if _acdcd._bdcce > ctx.Height {
			_gdag = append(_gdag, _ddeb)
			_ddeb = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_fbbb := ctx
			_fbbb.Y = ctx.Margins.Top + _acdcd._abgc.Top
			_fbbb.X = ctx.Margins.Left + _acdcd._abgc.Left
			_fbbb.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _acdcd._abgc.Top - _acdcd._abgc.Bottom
			_fbbb.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _acdcd._abgc.Left - _acdcd._abgc.Right
			ctx = _fbbb
		}
	} else {
		ctx.X = _acdcd._gdbee - _acdcd._eded/2
		ctx.Y = _acdcd._bgfb - _acdcd._bdcce/2
	}
	_bggf := _ff.Circle{X: ctx.X, Y: ctx.PageHeight - ctx.Y - _acdcd._bdcce, Width: _acdcd._eded, Height: _acdcd._bdcce, BorderWidth: _acdcd._eeaa, Opacity: 1.0}
	if _acdcd._fecae != nil {
		_bggf.FillEnabled = true
		_agcf := _fce(_acdcd._fecae)
		_dfff := _acdb(_ddeb, _agcf, _acdcd._fecae, func() Rectangle {
			return Rectangle{_fcgf: _bggf.X, _cbaa: _bggf.Y, _bcffg: _bggf.Width, _ebegg: _bggf.Height}
		})
		if _dfff != nil {
			return nil, ctx, _dfff
		}
		_bggf.FillColor = _agcf
	}
	if _acdcd._bacf != nil {
		_bggf.BorderEnabled = false
		if _acdcd._eeaa > 0 {
			_bggf.BorderEnabled = true
		}
		_bggf.BorderColor = _fce(_acdcd._bacf)
		_bggf.BorderWidth = _acdcd._eeaa
	}
	_cdcd, _fbfd := _ddeb.setOpacity(_acdcd._bbdc, _acdcd._eade)
	if _fbfd != nil {
		return nil, ctx, _fbfd
	}
	_cffe, _, _fbfd := _bggf.Draw(_cdcd)
	if _fbfd != nil {
		return nil, ctx, _fbfd
	}
	_fbfd = _ddeb.addContentsByString(string(_cffe))
	if _fbfd != nil {
		return nil, ctx, _fbfd
	}
	if _gadc {
		ctx.X = _agfc.X
		ctx.Width = _agfc.Width
		ctx.Y += _acdcd._bdcce + _acdcd._abgc.Bottom
		ctx.Height -= _acdcd._bdcce
	} else {
		ctx = _agfc
	}
	_gdag = append(_gdag, _ddeb)
	return _gdag, ctx, nil
}

// Insert adds a new text chunk at the specified position in the paragraph.
func (_ecgab *StyledParagraph) Insert(index uint, text string) *TextChunk {
	_bcfd := uint(len(_ecgab._eddc))
	if index > _bcfd {
		index = _bcfd
	}
	_gcgd := NewTextChunk(text, _ecgab._afaf)
	_ecgab._eddc = append(_ecgab._eddc[:index], append([]*TextChunk{_gcgd}, _ecgab._eddc[index:]...)...)
	_ecgab.wrapText()
	return _gcgd
}

// Width returns the cell's width based on the input draw context.
func (_ffbb *TableCell) Width(ctx DrawContext) float64 {
	_fcfbb := float64(0.0)
	for _cbggc := 0; _cbggc < _ffbb._daec; _cbggc++ {
		_fcfbb += _ffbb._ccbf._eeggd[_ffbb._fcbe+_cbggc-1]
	}
	_eefd := ctx.Width * _fcfbb
	return _eefd
}

// BuyerAddress returns the buyer address used in the invoice template.
func (_dfcc *Invoice) BuyerAddress() *InvoiceAddress { return _dfcc._cfccb }

// SetDashPattern sets the dash pattern of the line.
// NOTE: the dash pattern is taken into account only if the style of the
// line is set to dashed.
func (_gfac *Line) SetDashPattern(dashArray []int64, dashPhase int64) {
	_gfac._fedd = dashArray
	_gfac._ffgdf = dashPhase
}

// Width returns the width of the Paragraph.
func (_agfda *Paragraph) Width() float64 {
	if _agfda._dbfb && int(_agfda._adcfe) > 0 {
		return _agfda._adcfe
	}
	return _agfda.getTextWidth() / 1000.0
}

// Positioning returns the type of positioning the line is set to use.
func (_fgdf *Line) Positioning() Positioning { return _fgdf._fcdd }

// MultiCell makes a new cell with the specified row span and col span
// and inserts it into the table at the current position.
func (_edbge *Table) MultiCell(rowspan, colspan int) *TableCell {
	_edbge._ebad++
	_gecg := (_edbge.moveToNextAvailableCell()-1)%(_edbge._edfe) + 1
	_ebca := (_edbge._ebad-1)/_edbge._edfe + 1
	for _ebca > _edbge._dbbf {
		_edbge._dbbf++
		_edbge._ecab = append(_edbge._ecab, _edbge._bebbc)
	}
	_fabf := &TableCell{}
	_fabf._bgdca = _ebca
	_fabf._fcbe = _gecg
	_fabf._fccd = 5
	_fabf._ebce = CellBorderStyleNone
	_fabf._acafb = _ff.LineStyleSolid
	_fabf._dbdc = CellHorizontalAlignmentLeft
	_fabf._ddfbf = CellVerticalAlignmentTop
	_fabf._aabc = 0
	_fabf._aeba = 0
	_fabf._edag = 0
	_fabf._fgbd = 0
	_edecg := ColorBlack
	_fabf._efcg = _edecg
	_fabf._ebgec = _edecg
	_fabf._fcec = _edecg
	_fabf._afddf = _edecg
	if rowspan < 1 {
		_bcd.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061t\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0072\u006f\u0077s\u0070\u0061n\u0020\u0074o\u00201\u002e", rowspan)
		rowspan = 1
	}
	_eebbfe := _edbge._dbbf - (_fabf._bgdca - 1)
	if rowspan > _eebbfe {
		_bcd.Log.Debug("\u0054\u0061b\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070\u0061\u006e\u0020\u0028\u0025d\u0029\u0020\u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072o\u0077\u0073 \u0028\u0025\u0064\u0029.\u0020\u0041\u0064\u0064\u0069n\u0067\u0020\u0072\u006f\u0077\u0073\u002e", rowspan, _eebbfe)
		_edbge._dbbf += rowspan - 1
		for _gfacd := 0; _gfacd <= rowspan-_eebbfe; _gfacd++ {
			_edbge._ecab = append(_edbge._ecab, _edbge._bebbc)
		}
	}
	for _ceaed := 0; _ceaed < colspan && _gecg+_ceaed-1 < len(_edbge._bacd); _ceaed++ {
		_edbge._bacd[_gecg+_ceaed-1] = rowspan - 1
	}
	_fabf._fddc = rowspan
	if colspan < 1 {
		_bcd.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0063\u006f\u006c\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061n\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006cs\u0070\u0061n\u0020\u0074o\u00201\u002e", colspan)
		colspan = 1
	}
	_dfeb := _edbge._edfe - (_fabf._fcbe - 1)
	if colspan > _dfeb {
		_bcd.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0065\u006c\u006c\u0020\u0063o\u006c\u0073\u0070\u0061\u006e\u0020\u0028\u0025\u0064\u0029\u0020\u0065\u0078\u0063\u0065\u0065\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072\u006f\u0077\u0020\u0063\u006f\u006c\u0073\u0020\u0028\u0025d\u0029\u002e\u0020\u0041\u0064\u006a\u0075\u0073\u0074\u0069\u006e\u0067 \u0063\u006f\u006c\u0073\u0070\u0061n\u002e", colspan, _dfeb)
		colspan = _dfeb
	}
	_fabf._daec = colspan
	_edbge._ebad += colspan - 1
	_edbge._gebec = append(_edbge._gebec, _fabf)
	_fabf._ccbf = _edbge
	return _fabf
}

func _efff(_aebf *Block, _adfb *Paragraph, _gggc DrawContext) (DrawContext, error) {
	_bbabg := 1
	_efda := _cc.PdfObjectName("\u0046\u006f\u006e\u0074" + _aa.Itoa(_bbabg))
	for _aebf._ccf.HasFontByName(_efda) {
		_bbabg++
		_efda = _cc.PdfObjectName("\u0046\u006f\u006e\u0074" + _aa.Itoa(_bbabg))
	}
	_abbbe := _aebf._ccf.SetFontByName(_efda, _adfb._bfbd.ToPdfObject())
	if _abbbe != nil {
		return _gggc, _abbbe
	}
	_adfb.wrapText()
	_gcdf := _da.NewContentCreator()
	_gcdf.Add_q()
	_efag := _gggc.PageHeight - _gggc.Y - _adfb._dbad*_adfb._agdb
	_gcdf.Translate(_gggc.X, _efag)
	if _adfb._cgaaf != 0 {
		_gcdf.RotateDeg(_adfb._cgaaf)
	}
	_caeff := _fce(_adfb._ggbe)
	_abbbe = _acdb(_aebf, _caeff, _adfb._ggbe, func() Rectangle {
		return Rectangle{_fcgf: _gggc.X, _cbaa: _efag, _bcffg: _adfb.getMaxLineWidth() / 1000.0, _ebegg: _adfb.Height()}
	})
	if _abbbe != nil {
		return _gggc, _abbbe
	}
	_gcdf.Add_BT().SetNonStrokingColor(_caeff).Add_Tf(_efda, _adfb._dbad).Add_TL(_adfb._dbad * _adfb._agdb)
	for _fbee, _bfcb := range _adfb._aggb {
		if _fbee != 0 {
			_gcdf.Add_Tstar()
		}
		_agcdf := []rune(_bfcb)
		_ededb := 0.0
		_dbfg := 0
		for _bcdcf, _eceed := range _agcdf {
			if _eceed == ' ' {
				_dbfg++
				continue
			}
			if _eceed == '\u000A' {
				continue
			}
			_gfgc, _bgbf := _adfb._bfbd.GetRuneMetrics(_eceed)
			if !_bgbf {
				_bcd.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006e\u0065\u0020\u0069=\u0025\u0064\u0020\u0072\u0075\u006e\u0065=\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0069n\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0025\u0073", _bcdcf, _eceed, _eceed, _adfb._bfbd.BaseFont(), _adfb._bfbd.Subtype())
				return _gggc, _c.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
			}
			_ededb += _adfb._dbad * _gfgc.Wx
		}
		var _ggc []_cc.PdfObject
		_fcaded, _eaag := _adfb._bfbd.GetRuneMetrics(' ')
		if !_eaag {
			return _gggc, _c.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
		}
		_bgbec := _fcaded.Wx
		switch _adfb._caea {
		case TextAlignmentJustify:
			if _dbfg > 0 && _fbee < len(_adfb._aggb)-1 {
				_bgbec = (_adfb._adcfe*1000.0 - _ededb) / float64(_dbfg) / _adfb._dbad
			}
		case TextAlignmentCenter:
			_gccce := _ededb + float64(_dbfg)*_bgbec*_adfb._dbad
			_bead := (_adfb._adcfe*1000.0 - _gccce) / 2 / _adfb._dbad
			_ggc = append(_ggc, _cc.MakeFloat(-_bead))
		case TextAlignmentRight:
			_ceed := _ededb + float64(_dbfg)*_bgbec*_adfb._dbad
			_cdgg := (_adfb._adcfe*1000.0 - _ceed) / _adfb._dbad
			_ggc = append(_ggc, _cc.MakeFloat(-_cdgg))
		}
		_gdcd := _adfb._bfbd.Encoder()
		var _dade []byte
		for _, _gebda := range _agcdf {
			if _gebda == '\u000A' {
				continue
			}
			if _gebda == ' ' {
				if len(_dade) > 0 {
					_ggc = append(_ggc, _cc.MakeStringFromBytes(_dade))
					_dade = nil
				}
				_ggc = append(_ggc, _cc.MakeFloat(-_bgbec))
			} else {
				if _, _bacc := _gdcd.RuneToCharcode(_gebda); !_bacc {
					_abbbe = UnsupportedRuneError{Message: _g.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _gebda, _gebda), Rune: _gebda}
					_gggc._ddcc = append(_gggc._ddcc, _abbbe)
					_bcd.Log.Debug(_abbbe.Error())
					if _gggc._dfe <= 0 {
						continue
					}
					_gebda = _gggc._dfe
				}
				_dade = append(_dade, _gdcd.Encode(string(_gebda))...)
			}
		}
		if len(_dade) > 0 {
			_ggc = append(_ggc, _cc.MakeStringFromBytes(_dade))
		}
		_gcdf.Add_TJ(_ggc...)
	}
	_gcdf.Add_ET()
	_gcdf.Add_Q()
	_ffcb := _gcdf.Operations()
	_ffcb.WrapIfNeeded()
	_aebf.addContents(_ffcb)
	if _adfb._gece.IsRelative() {
		_effc := _adfb.Height()
		_gggc.Y += _effc
		_gggc.Height -= _effc
		if _gggc.Inline {
			_gggc.X += _adfb.Width() + _adfb._bcggf.Right
		}
	}
	return _gggc, nil
}

// Wrap wraps the text of the chunk into lines based on its style and the
// specified width.
func (_egge *TextChunk) Wrap(width float64) ([]string, error) {
	if int(width) <= 0 {
		return []string{_egge.Text}, nil
	}
	var _bgeb []string
	var _dcddb []rune
	var _ceag float64
	var _abded []float64
	_cdafa := _egge.Style
	_ebcad := _cbea(_egge.Text)
	for _, _fegaf := range _egge.Text {
		if _fegaf == '\u000A' {
			_acgd := _gadcc(string(_dcddb), _ebcad)
			_bgeb = append(_bgeb, _a.TrimRightFunc(_acgd, _bc.IsSpace)+string(_fegaf))
			_dcddb = nil
			_ceag = 0
			_abded = nil
			continue
		}
		_afcag := _fegaf == ' '
		_bdeca, _gdadd := _cdafa.Font.GetRuneMetrics(_fegaf)
		if !_gdadd {
			_bcd.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _fegaf, _fegaf, _cdafa.Font.BaseFont(), _cdafa.Font.Subtype())
			_bcd.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _cdafa.Font)
			_bcd.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _cdafa.Font.Encoder())
			return nil, _c.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
		_dgfde := _cdafa.FontSize * _bdeca.Wx
		_dfadg := _dgfde
		if !_afcag {
			_dfadg = _dgfde + _cdafa.CharSpacing*1000.0
		}
		if _ceag+_dgfde > width*1000.0 {
			_deabf := -1
			if !_afcag {
				for _dacbg := len(_dcddb) - 1; _dacbg >= 0; _dacbg-- {
					if _dcddb[_dacbg] == ' ' {
						_deabf = _dacbg
						break
					}
				}
			}
			_gfecgf := string(_dcddb)
			if _deabf > 0 {
				_gfecgf = string(_dcddb[0 : _deabf+1])
				_dcddb = append(_dcddb[_deabf+1:], _fegaf)
				_abded = append(_abded[_deabf+1:], _dfadg)
				_ceag = 0
				for _, _gedbg := range _abded {
					_ceag += _gedbg
				}
			} else {
				if _afcag {
					_dcddb = []rune{}
					_abded = []float64{}
					_ceag = 0
				} else {
					_dcddb = []rune{_fegaf}
					_abded = []float64{_dfadg}
					_ceag = _dfadg
				}
			}
			_gfecgf = _gadcc(_gfecgf, _ebcad)
			_bgeb = append(_bgeb, _a.TrimRightFunc(_gfecgf, _bc.IsSpace))
		} else {
			_dcddb = append(_dcddb, _fegaf)
			_ceag += _dfadg
			_abded = append(_abded, _dfadg)
		}
	}
	if len(_dcddb) > 0 {
		_eccad := string(_dcddb)
		_eccad = _gadcc(_eccad, _ebcad)
		_bgeb = append(_bgeb, _eccad)
	}
	return _bgeb, nil
}

// GetCoords returns the center coordinates of ellipse (`xc`, `yc`).
func (_dgged *Ellipse) GetCoords() (float64, float64) { return _dgged._gdbee, _dgged._bgfb }

// GeneratePageBlocks draw graphic svg into block.
func (_ceaf *GraphicSVG) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gdf := ctx
	_cbc := _ceaf._dabg.IsRelative()
	var _cafc []*Block
	if _cbc {
		_bebb := 1.0
		_dfcg := _ceaf._fgfa.Top
		if _ceaf._debae.Height > ctx.Height-_ceaf._fgfa.Top {
			_cafc = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _ggbg error
			if _, ctx, _ggbg = _dbae().GeneratePageBlocks(ctx); _ggbg != nil {
				return nil, ctx, _ggbg
			}
			_dfcg = 0
		}
		ctx.X += _ceaf._fgfa.Left + _bebb
		ctx.Y += _dfcg
		ctx.Width -= _ceaf._fgfa.Left + _ceaf._fgfa.Right + 2*_bebb
		ctx.Height -= _dfcg
	} else {
		ctx.X = _ceaf._dfgae
		ctx.Y = _ceaf._egbf
	}
	_dbf := _da.NewContentCreator()
	_dbf.Translate(0, ctx.PageHeight)
	_dbf.Scale(1, -1)
	_dbf.Translate(ctx.X, ctx.Y)
	_ffbc := _ceaf._debae.Width / _ceaf._debae.ViewBox.W
	_fdfab := _ceaf._debae.Height / _ceaf._debae.ViewBox.H
	_gbec := 0.0
	_ddgg := 0.0
	if _cbc {
		_gbec = _ceaf._dfgae - (_ceaf._debae.ViewBox.X * _cd.Max(_ffbc, _fdfab))
		_ddgg = _ceaf._egbf - (_ceaf._debae.ViewBox.Y * _cd.Max(_ffbc, _fdfab))
	}
	_ceaf._debae.ToContentCreator(_dbf, _ffbc, _fdfab, _gbec, _ddgg)
	_bccc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _gbea := _bccc.addContentsByString(_dbf.String()); _gbea != nil {
		return nil, ctx, _gbea
	}
	if _cbc {
		_efea := _ceaf.Height() + _ceaf._fgfa.Bottom
		ctx.Y += _efea
		ctx.Height -= _efea
	} else {
		ctx = _gdf
	}
	_cafc = append(_cafc, _bccc)
	return _cafc, ctx, nil
}

// BorderOpacity returns the border opacity of the ellipse (0-1).
func (_agdg *Ellipse) BorderOpacity() float64 { return _agdg._eade }

func (_afaefd *StyledParagraph) getMaxLineWidth() float64 {
	if _afaefd._bddb == nil || len(_afaefd._bddb) == 0 {
		_afaefd.wrapText()
	}
	var _abaag float64
	for _, _cbdad := range _afaefd._bddb {
		_edbg := _afaefd.getTextLineWidth(_cbdad)
		if _edbg > _abaag {
			_abaag = _edbg
		}
	}
	return _abaag
}

// Sections returns the custom content sections of the invoice as
// title-content pairs.
func (_fdbdg *Invoice) Sections() [][2]string { return _fdbdg._eefc }

func _dcabe(_bfca *Block, _gbef *Image, _dcce DrawContext) (DrawContext, error) {
	_gaff := _dcce
	_aacbb := 1
	_bbfe := _cc.PdfObjectName(_g.Sprintf("\u0049\u006d\u0067%\u0064", _aacbb))
	for _bfca._ccf.HasXObjectByName(_bbfe) {
		_aacbb++
		_bbfe = _cc.PdfObjectName(_g.Sprintf("\u0049\u006d\u0067%\u0064", _aacbb))
	}
	_fade := _bfca._ccf.SetXObjectImageByName(_bbfe, _gbef._ebfc)
	if _fade != nil {
		return _dcce, _fade
	}
	_adfac := 0
	_eagga := _cc.PdfObjectName(_g.Sprintf("\u0047\u0053\u0025\u0064", _adfac))
	for _bfca._ccf.HasExtGState(_eagga) {
		_adfac++
		_eagga = _cc.PdfObjectName(_g.Sprintf("\u0047\u0053\u0025\u0064", _adfac))
	}
	_dggdc := _cc.MakeDict()
	_dggdc.Set("\u0042\u004d", _cc.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	if _gbef._gafe < 1.0 {
		_dggdc.Set("\u0043\u0041", _cc.MakeFloat(_gbef._gafe))
		_dggdc.Set("\u0063\u0061", _cc.MakeFloat(_gbef._gafe))
	}
	_fade = _bfca._ccf.AddExtGState(_eagga, _cc.MakeIndirectObject(_dggdc))
	if _fade != nil {
		return _dcce, _fade
	}
	_dac := _gbef.Width()
	_decc := _gbef.Height()
	_, _fega := _gbef.rotatedSize()
	_ddfgd := _dcce.X
	_dgd := _dcce.PageHeight - _dcce.Y - _decc
	if _gbef._geca.IsRelative() {
		_dgd -= (_fega - _decc) / 2
		switch _gbef._cccd {
		case HorizontalAlignmentCenter:
			_ddfgd += (_dcce.Width - _dac) / 2
		case HorizontalAlignmentRight:
			_ddfgd = _dcce.PageWidth - _dcce.Margins.Right - _gbef._eaac.Right - _dac
		}
	}
	_gafc := _gbef._egacc
	_gbeaf := _da.NewContentCreator()
	_gbeaf.Add_gs(_eagga)
	_gbeaf.Translate(_ddfgd, _dgd)
	if _gafc != 0 {
		_gbeaf.Translate(_dac/2, _decc/2)
		_gbeaf.RotateDeg(_gafc)
		_gbeaf.Translate(-_dac/2, -_decc/2)
	}
	_gbeaf.Scale(_dac, _decc).Add_Do(_bbfe)
	_fgdbb := _gbeaf.Operations()
	_fgdbb.WrapIfNeeded()
	_bfca.addContents(_fgdbb)
	if _gbef._geca.IsRelative() {
		_dcce.Y += _fega
		_dcce.Height -= _fega
		return _dcce, nil
	}
	return _gaff, nil
}

// SetFillColor sets the fill color.
func (_dagde *Polygon) SetFillColor(color Color) {
	_dagde._ddad = color
	_dagde._acaag.FillColor = _fce(color)
}

func (_badfg *Rectangle) applyFitMode(_fbgf float64) {
	_fbgf -= _badfg._eaaa.Left + _badfg._eaaa.Right + _badfg._bffc
	switch _badfg._efbb {
	case FitModeFillWidth:
		_badfg.ScaleToWidth(_fbgf)
	}
}

func _fdfa(_aaga string) (*GraphicSVG, error) {
	_gaag, _fdfc := _ca.ParseFromFile(_aaga)
	if _fdfc != nil {
		return nil, _fdfc
	}
	return _cbba(_gaag)
}

// Marker returns the marker used for the list items.
// The marker instance can be used the change the text and the style
// of newly added list items.
func (_cabab *List) Marker() *TextChunk { return &_cabab._feeb }

// Division is a container component which can wrap across multiple pages.
// Currently supported drawable components:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
//
// The component stacking behavior is vertical, where the drawables are drawn
// on top of each other.
type Division struct {
	_dccd []VectorDrawable
	_acfc Positioning
	_acaa Margins
	_cgbb Margins
	_cdgb bool
	_dccg bool
	_eeaf *Background
}

// SetNotes sets the notes section of the invoice.
func (_debc *Invoice) SetNotes(title, content string) { _debc._aeee = [2]string{title, content} }

func _gebee(_ffebc *Creator, _afedg _eb.Reader, _gbbc interface{}, _aeacg *TemplateOptions, _dadcd componentRenderer) error {
	if _ffebc == nil {
		_bcd.Log.Error("\u0043\u0072\u0065a\u0074\u006f\u0072\u0020i\u006e\u0073\u0074\u0061\u006e\u0063\u0065 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return _caca
	}
	_feeec := ""
	if _gdfgeb, _aagaf := _afedg.(*_e.File); _aagaf {
		_feeec = _gdfgeb.Name()
	}
	_dfee := _b.NewBuffer(nil)
	if _, _agbac := _eb.Copy(_dfee, _afedg); _agbac != nil {
		return _agbac
	}
	_ffbef := _dd.FuncMap{"\u0064\u0069\u0063\u0074": _aaagb}
	if _aeacg != nil && _aeacg.HelperFuncMap != nil {
		for _gfdgf, _dfbb := range _aeacg.HelperFuncMap {
			if _, _cgab := _ffbef[_gfdgf]; _cgab {
				_bcd.Log.Debug("\u0043\u0061\u006e\u006e\u006f\u0074 \u006f\u0076\u0065r\u0072\u0069\u0064e\u0020\u0062\u0075\u0069\u006c\u0074\u002d\u0069\u006e\u0020`\u0025\u0073\u0060\u0020\u0068el\u0070\u0065\u0072\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _gfdgf)
				continue
			}
			_ffbef[_gfdgf] = _dfbb
		}
	}
	_cbcf, _fgfg := _dd.New("").Funcs(_ffbef).Parse(_dfee.String())
	if _fgfg != nil {
		return _fgfg
	}
	if _aeacg != nil && _aeacg.SubtemplateMap != nil {
		for _fecc, _eaddf := range _aeacg.SubtemplateMap {
			if _fecc == "" {
				_bcd.Log.Debug("\u0053\u0075\u0062\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067.\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
				continue
			}
			if _eaddf == nil {
				_bcd.Log.Debug("S\u0075\u0062t\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u0061\u006e\u006eo\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079 \u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e")
				continue
			}
			_cface := _b.NewBuffer(nil)
			if _, _gcff := _eb.Copy(_cface, _eaddf); _gcff != nil {
				return _gcff
			}
			if _, _agec := _cbcf.New(_fecc).Parse(_cface.String()); _agec != nil {
				return _agec
			}
		}
	}
	_dfee.Reset()
	if _ceccg := _cbcf.Execute(_dfee, _gbbc); _ceccg != nil {
		return _ceccg
	}
	return _edae(_ffebc, _feeec, _dfee.Bytes(), _aeacg, _dadcd).run()
}

func (_fgbbg *RadialShading) shadingModel() *_ga.PdfShadingType3 {
	_gdceg, _gcca, _gdbfe := _fgbbg._acdae._cfec.ToRGB()
	var _cgcc _ff.Point
	switch _fgbbg._bafe {
	case AnchorBottomLeft:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Llx, Y: _fgbbg._bcgga.Lly}
	case AnchorBottomRight:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Urx, Y: _fgbbg._bcgga.Ury - _fgbbg._bcgga.Height()}
	case AnchorTopLeft:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Llx, Y: _fgbbg._bcgga.Lly + _fgbbg._bcgga.Height()}
	case AnchorTopRight:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Urx, Y: _fgbbg._bcgga.Ury}
	case AnchorLeft:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Llx, Y: _fgbbg._bcgga.Lly + _fgbbg._bcgga.Height()/2}
	case AnchorTop:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Llx + _fgbbg._bcgga.Width()/2, Y: _fgbbg._bcgga.Ury}
	case AnchorRight:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Urx, Y: _fgbbg._bcgga.Lly + _fgbbg._bcgga.Height()/2}
	case AnchorBottom:
		_cgcc = _ff.Point{X: _fgbbg._bcgga.Urx + _fgbbg._bcgga.Width()/2, Y: _fgbbg._bcgga.Lly}
	default:
		_cgcc = _ff.NewPoint(_fgbbg._bcgga.Llx+_fgbbg._bcgga.Width()/2, _fgbbg._bcgga.Lly+_fgbbg._bcgga.Height()/2)
	}
	_decca := _fgbbg._fecb
	_gddf := _fgbbg._dadbe
	_gefa := _cgcc.X + _fgbbg._dfgdc
	_cege := _cgcc.Y + _fgbbg._acfgc
	if _decca == -1.0 {
		_decca = 0.0
	}
	if _gddf == -1.0 {
		var _gdab []float64
		_bdcde := _cd.Pow(_gefa-_fgbbg._bcgga.Llx, 2) + _cd.Pow(_cege-_fgbbg._bcgga.Lly, 2)
		_gdab = append(_gdab, _cd.Abs(_bdcde))
		_agbefg := _cd.Pow(_gefa-_fgbbg._bcgga.Llx, 2) + _cd.Pow(_fgbbg._bcgga.Lly+_fgbbg._bcgga.Height()-_cege, 2)
		_gdab = append(_gdab, _cd.Abs(_agbefg))
		_gdbge := _cd.Pow(_fgbbg._bcgga.Urx-_gefa, 2) + _cd.Pow(_cege-_fgbbg._bcgga.Ury-_fgbbg._bcgga.Height(), 2)
		_gdab = append(_gdab, _cd.Abs(_gdbge))
		_ffac := _cd.Pow(_fgbbg._bcgga.Urx-_gefa, 2) + _cd.Pow(_fgbbg._bcgga.Ury-_cege, 2)
		_gdab = append(_gdab, _cd.Abs(_ffac))
		_f.Slice(_gdab, func(_gbgg, _ccbac int) bool { return _gbgg > _ccbac })
		_gddf = _cd.Sqrt(_gdab[0])
	}
	_fafefd := &_ga.PdfRectangle{Llx: _gefa - _gddf, Lly: _cege - _gddf, Urx: _gefa + _gddf, Ury: _cege + _gddf}
	_deea := _ga.NewPdfShadingType3()
	_deea.PdfShading.ShadingType = _cc.MakeInteger(3)
	_deea.PdfShading.ColorSpace = _ga.NewPdfColorspaceDeviceRGB()
	_deea.PdfShading.Background = _cc.MakeArrayFromFloats([]float64{_gdceg, _gcca, _gdbfe})
	_deea.PdfShading.BBox = _fafefd
	_deea.PdfShading.AntiAlias = _cc.MakeBool(_fgbbg._acdae._cebg)
	_deea.Coords = _cc.MakeArrayFromFloats([]float64{_gefa, _cege, _decca, _gefa, _cege, _gddf})
	_deea.Domain = _cc.MakeArrayFromFloats([]float64{0.0, 1.0})
	_deea.Extend = _cc.MakeArray(_cc.MakeBool(_fgbbg._acdae._fedf[0]), _cc.MakeBool(_fgbbg._acdae._fedf[1]))
	_deea.Function = _fgbbg._acdae.generatePdfFunctions()
	return _deea
}

// GeneratePageBlocks draws the chart onto a block.
func (_fcf *Chart) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gee := ctx
	_gga := _fcf._bdcb.IsRelative()
	var _cbe []*Block
	if _gga {
		_gbe := 1.0
		_egdc := _fcf._baae.Top
		if float64(_fcf._dece.Height()) > ctx.Height-_fcf._baae.Top {
			_cbe = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _fdbg error
			if _, ctx, _fdbg = _dbae().GeneratePageBlocks(ctx); _fdbg != nil {
				return nil, ctx, _fdbg
			}
			_egdc = 0
		}
		ctx.X += _fcf._baae.Left + _gbe
		ctx.Y += _egdc
		ctx.Width -= _fcf._baae.Left + _fcf._baae.Right + 2*_gbe
		ctx.Height -= _egdc
		_fcf._dece.SetWidth(int(ctx.Width))
	} else {
		ctx.X = _fcf._cdd
		ctx.Y = _fcf._dcd
	}
	_dgfd := _da.NewContentCreator()
	_dgfd.Translate(0, ctx.PageHeight)
	_dgfd.Scale(1, -1)
	_dgfd.Translate(ctx.X, ctx.Y)
	_gcg := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_fcf._dece.Render(_cad.NewRenderer(_dgfd, _gcg._ccf), nil)
	if _cbg := _gcg.addContentsByString(_dgfd.String()); _cbg != nil {
		return nil, ctx, _cbg
	}
	if _gga {
		_cdaeb := _fcf.Height() + _fcf._baae.Bottom
		ctx.Y += _cdaeb
		ctx.Height -= _cdaeb
	} else {
		ctx = _gee
	}
	_cbe = append(_cbe, _gcg)
	return _cbe, ctx, nil
}

func _cfeg(_baba, _dgbeb, _cfdbe, _gadg float64) *Rectangle {
	return &Rectangle{_fcgf: _baba, _cbaa: _dgbeb, _bcffg: _cfdbe, _ebegg: _gadg, _dgde: PositionAbsolute, _gfdd: 1.0, _gggb: ColorBlack, _bffc: 1.0, _gffa: 1.0}
}

func (_bcgfac *templateProcessor) parseCellVerticalAlignmentAttr(_eagcb, _gebde string) CellVerticalAlignment {
	_bcd.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u0065\u006c\u006c\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _eagcb, _gebde)
	_bdgdc := map[string]CellVerticalAlignment{"\u0074\u006f\u0070": CellVerticalAlignmentTop, "\u006d\u0069\u0064\u0064\u006c\u0065": CellVerticalAlignmentMiddle, "\u0062\u006f\u0074\u0074\u006f\u006d": CellVerticalAlignmentBottom}[_gebde]
	return _bdgdc
}

// CreateTableOfContents sets a function to generate table of contents.
func (_fef *Creator) CreateTableOfContents(genTOCFunc func(_cfa *TOC) error) { _fef._caf = genTOCFunc }

func (_bcbdg *Table) getLastCellFromCol(_effba int) (int, *TableCell) {
	for _bfge := len(_bcbdg._gebec) - 1; _bfge >= 0; _bfge-- {
		if _bcbdg._gebec[_bfge]._fcbe == _effba {
			return _bfge, _bcbdg._gebec[_bfge]
		}
	}
	return 0, nil
}

// SetSideBorderWidth sets the cell's side border width.
func (_bgecf *TableCell) SetSideBorderWidth(side CellBorderSide, width float64) {
	switch side {
	case CellBorderSideAll:
		_bgecf._fgbd = width
		_bgecf._aeba = width
		_bgecf._aabc = width
		_bgecf._edag = width
	case CellBorderSideTop:
		_bgecf._fgbd = width
	case CellBorderSideBottom:
		_bgecf._aeba = width
	case CellBorderSideLeft:
		_bgecf._aabc = width
	case CellBorderSideRight:
		_bgecf._edag = width
	}
}

// GeneratePageBlocks generates a page break block.
func (_gggd *PageBreak) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bdce := []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y), NewBlock(ctx.PageWidth, ctx.PageHeight)}
	ctx.Page++
	_cffc := ctx
	_cffc.Y = ctx.Margins.Top
	_cffc.X = ctx.Margins.Left
	_cffc.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
	_cffc.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right
	ctx = _cffc
	return _bdce, ctx, nil
}

// AddColorStop add color stop info for rendering gradient color.
func (_bgcb *LinearShading) AddColorStop(color Color, point float64) {
	_bgcb._cfbfe.AddColorStop(color, point)
}

// SetLineNumberStyle sets the style for the numbers part of all new lines
// of the table of contents.
func (_dfbe *TOC) SetLineNumberStyle(style TextStyle) { _dfbe._daaaa = style }

func _acdb(_fcbc *Block, _cebd _ga.PdfColor, _bgbef Color, _dafc func() Rectangle) error {
	switch _cedbd := _cebd.(type) {
	case *_ga.PdfColorPatternType2:
		_edcc, _ffd := _bgbef.(*LinearShading)
		if !_ffd {
			return _g.Errorf("\u0043\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u004c\u0069\u006e\u0065\u0061\u0072\u0053\u0068\u0061d\u0069\u006e\u0067")
		}
		_fcbg := _dafc()
		_edcc.SetBoundingBox(_fcbg._fcgf, _fcbg._cbaa, _fcbg._bcffg, _fcbg._ebegg)
		_fdaag, _gggdd := _edcc.AddPatternResource(_fcbc)
		if _gggdd != nil {
			return _g.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0074\u006f \u0072\u0065\u0073\u006f\u0075r\u0063\u0065s\u003a\u0020\u0025\u0076", _gggdd)
		}
		_cedbd.PatternName = _fdaag
	case *_ga.PdfColorPatternType3:
		_beac, _dedc := _bgbef.(*RadialShading)
		if !_dedc {
			return _g.Errorf("\u0043\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0052\u0061\u0064\u0069\u0061\u006c\u0053\u0068\u0061d\u0069\u006e\u0067")
		}
		_fbff := _dafc()
		_beac.SetBoundingBox(_fbff._fcgf, _fbff._cbaa, _fbff._bcffg, _fbff._ebegg)
		_effgg, _dcgcd := _beac.AddPatternResource(_fcbc)
		if _dcgcd != nil {
			return _g.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0074\u006f \u0072\u0065\u0073\u006f\u0075r\u0063\u0065s\u003a\u0020\u0025\u0076", _dcgcd)
		}
		_cedbd.PatternName = _effgg
	}
	return nil
}

func (_cegb *templateProcessor) parseFontAttr(_baagb, _ggbf string) *_ga.PdfFont {
	_bcd.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _baagb, _ggbf)
	_cbcef := _cegb.creator._ffcg
	if _ggbf == "" {
		return _cbcef
	}
	_adddc := _a.Split(_ggbf, "\u002c")
	for _, _acba := range _adddc {
		_acba = _a.TrimSpace(_acba)
		if _acba == "" {
			continue
		}
		_gdgd, _ffbd := _cegb._agcb.FontMap[_ggbf]
		if _ffbd {
			return _gdgd
		}
		_fegc, _ffbd := map[string]_ga.StdFontName{"\u0063o\u0075\u0072\u0069\u0065\u0072": _ga.CourierName, "\u0063\u006f\u0075r\u0069\u0065\u0072\u002d\u0062\u006f\u006c\u0064": _ga.CourierBoldName, "\u0063o\u0075r\u0069\u0065\u0072\u002d\u006f\u0062\u006c\u0069\u0071\u0075\u0065": _ga.CourierObliqueName, "c\u006fu\u0072\u0069\u0065\u0072\u002d\u0062\u006f\u006cd\u002d\u006f\u0062\u006ciq\u0075\u0065": _ga.CourierBoldObliqueName, "\u0068e\u006c\u0076\u0065\u0074\u0069\u0063a": _ga.HelveticaName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0062\u006f\u006c\u0064": _ga.HelveticaBoldName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u006f\u0062l\u0069\u0071\u0075\u0065": _ga.HelveticaObliqueName, "\u0068\u0065\u006c\u0076et\u0069\u0063\u0061\u002d\u0062\u006f\u006c\u0064\u002d\u006f\u0062\u006c\u0069\u0071u\u0065": _ga.HelveticaBoldObliqueName, "\u0073\u0079\u006d\u0062\u006f\u006c": _ga.SymbolName, "\u007a\u0061\u0070\u0066\u002d\u0064\u0069\u006e\u0067\u0062\u0061\u0074\u0073": _ga.ZapfDingbatsName, "\u0074\u0069\u006de\u0073": _ga.TimesRomanName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064": _ga.TimesBoldName, "\u0074\u0069\u006de\u0073\u002d\u0069\u0074\u0061\u006c\u0069\u0063": _ga.TimesItalicName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064\u002d\u0069t\u0061\u006c\u0069\u0063": _ga.TimesBoldItalicName}[_ggbf]
		if _ffbd {
			if _gbfca, _fegacg := _ga.NewStandard14Font(_fegc); _fegacg == nil {
				return _gbfca
			}
		}
		if _dfce := _cegb.parseAttrPropList(_acba); len(_dfce) > 0 {
			if _adeba, _dcafg := _dfce["\u0070\u0061\u0074\u0068"]; _dcafg {
				_abbbd := _ga.NewPdfFontFromTTFFile
				if _eabf, _aaaee := _dfce["\u0074\u0079\u0070\u0065"]; _aaaee && _eabf == "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e" {
					_abbbd = _ga.NewCompositePdfFontFromTTFFile
				}
				if _gefaf, _febba := _abbbd(_adeba); _febba != nil {
					_bcd.Log.Debug("\u0043\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0060\u0025\u0073\u0060\u003a %\u0076\u002e", _adeba, _febba)
				} else {
					return _gefaf
				}
			}
		}
	}
	return _cbcef
}

// SetStyleTop sets border style for top side.
func (_fbc *border) SetStyleTop(style CellBorderStyle) { _fbc._acef = style }

// SetBoundingBox set gradient color bounding box where the gradient would be rendered.
func (_cfae *LinearShading) SetBoundingBox(x, y, width, height float64) {
	_cfae._ebfea = &_ga.PdfRectangle{Llx: x, Lly: y, Urx: x + width, Ury: y + height}
}

// SetBorderColor sets the border color for the path.
func (_bade *FilledCurve) SetBorderColor(color Color) { _bade._afe = color }

func (_eeeb *templateProcessor) parseTableCell(_eccf *templateNode) (interface{}, error) {
	if _eccf._gbdf == nil {
		_eeeb.nodeLogError(_eccf, "\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065n\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _bfegb
	}
	_gabaf, _begec := _eccf._gbdf._defae.(*Table)
	if !_begec {
		_eeeb.nodeLogError(_eccf, "\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0028\u0025\u0054\u0029\u0020\u0069s\u0020\u006e\u006f\u0074\u0020a\u0020\u0074a\u0062\u006c\u0065\u002e", _eccf._gbdf._defae)
		return nil, _bfegb
	}
	var _fdge, _aefba int64
	for _, _eccda := range _eccf._acdf.Attr {
		_dfdfa := _eccda.Value
		switch _gabbd := _eccda.Name.Local; _gabbd {
		case "\u0063o\u006c\u0073\u0070\u0061\u006e":
			_fdge = _eeeb.parseInt64Attr(_gabbd, _dfdfa)
		case "\u0072o\u0077\u0073\u0070\u0061\u006e":
			_aefba = _eeeb.parseInt64Attr(_gabbd, _dfdfa)
		}
	}
	if _fdge <= 0 {
		_fdge = 1
	}
	if _aefba <= 0 {
		_aefba = 1
	}
	_aggaa := _gabaf.MultiCell(int(_aefba), int(_fdge))
	for _, _cbac := range _eccf._acdf.Attr {
		_fgfga := _cbac.Value
		switch _fdeb := _cbac.Name.Local; _fdeb {
		case "\u0069\u006e\u0064\u0065\u006e\u0074":
			_aggaa.SetIndent(_eeeb.parseFloatAttr(_fdeb, _fgfga))
		case "\u0061\u006c\u0069g\u006e":
			_aggaa.SetHorizontalAlignment(_eeeb.parseCellAlignmentAttr(_fdeb, _fgfga))
		case "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c\u002da\u006c\u0069\u0067\u006e":
			_aggaa.SetVerticalAlignment(_eeeb.parseCellVerticalAlignmentAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065":
			_aggaa.SetSideBorderStyle(CellBorderSideAll, _eeeb.parseCellBorderStyleAttr(_fdeb, _fgfga))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0073\u0074\u0079\u006c\u0065\u002d\u0074\u006f\u0070":
			_aggaa.SetSideBorderStyle(CellBorderSideTop, _eeeb.parseCellBorderStyleAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072\u0064er\u002d\u0073\u0074\u0079\u006c\u0065\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_aggaa.SetSideBorderStyle(CellBorderSideBottom, _eeeb.parseCellBorderStyleAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065-\u006c\u0065\u0066\u0074":
			_aggaa.SetSideBorderStyle(CellBorderSideLeft, _eeeb.parseCellBorderStyleAttr(_fdeb, _fgfga))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0073\u0074y\u006c\u0065-\u0072\u0069\u0067\u0068\u0074":
			_aggaa.SetSideBorderStyle(CellBorderSideRight, _eeeb.parseCellBorderStyleAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_aggaa.SetSideBorderWidth(CellBorderSideAll, _eeeb.parseFloatAttr(_fdeb, _fgfga))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0077\u0069\u0064\u0074\u0068\u002d\u0074\u006f\u0070":
			_aggaa.SetSideBorderWidth(CellBorderSideTop, _eeeb.parseFloatAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072\u0064er\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_aggaa.SetSideBorderWidth(CellBorderSideBottom, _eeeb.parseFloatAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068-\u006c\u0065\u0066\u0074":
			_aggaa.SetSideBorderWidth(CellBorderSideLeft, _eeeb.parseFloatAttr(_fdeb, _fgfga))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0077\u0069d\u0074\u0068-\u0072\u0069\u0067\u0068\u0074":
			_aggaa.SetSideBorderWidth(CellBorderSideRight, _eeeb.parseFloatAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_aggaa.SetSideBorderColor(CellBorderSideAll, _eeeb.parseColorAttr(_fdeb, _fgfga))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0063\u006f\u006c\u006f\u0072\u002d\u0074\u006f\u0070":
			_aggaa.SetSideBorderColor(CellBorderSideTop, _eeeb.parseColorAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072\u0064er\u002d\u0063\u006f\u006c\u006f\u0072\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_aggaa.SetSideBorderColor(CellBorderSideBottom, _eeeb.parseColorAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072-\u006c\u0065\u0066\u0074":
			_aggaa.SetSideBorderColor(CellBorderSideLeft, _eeeb.parseColorAttr(_fdeb, _fgfga))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0063\u006fl\u006f\u0072-\u0072\u0069\u0067\u0068\u0074":
			_aggaa.SetSideBorderColor(CellBorderSideRight, _eeeb.parseColorAttr(_fdeb, _fgfga))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006c\u0069\u006e\u0065\u002ds\u0074\u0079\u006c\u0065":
			_aggaa.SetBorderLineStyle(_eeeb.parseLineStyleAttr(_fdeb, _fgfga))
		case "\u0062\u0061c\u006b\u0067\u0072o\u0075\u006e\u0064\u002d\u0063\u006f\u006c\u006f\u0072":
			_aggaa.SetBackgroundColor(_eeeb.parseColorAttr(_fdeb, _fgfga))
		case "\u0063o\u006c\u0073\u0070\u0061\u006e", "\u0072o\u0077\u0073\u0070\u0061\u006e":
			break
		default:
			_eeeb.nodeLogDebug(_eccf, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _fdeb)
		}
	}
	return _aggaa, nil
}

func (_gbddd *templateProcessor) addNodeText(_gaedfd *templateNode, _ccdge string) error {
	_agcbg := _gaedfd._defae
	if _agcbg == nil {
		return nil
	}
	switch _gdfdg := _agcbg.(type) {
	case *TextChunk:
		_gdfdg.Text = _ccdge
	case *Paragraph:
		switch _gaedfd._acdf.Name.Local {
		case "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067":
			if _gaedfd._gbdf != nil {
				if _fbccg, _fcde := _gaedfd._gbdf._defae.(*Chapter); _fcde {
					_fbccg._gbb = _ccdge
					_gdfdg.SetText(_fbccg.headingText())
				}
			}
		default:
			_gdfdg.SetText(_ccdge)
		}
	}
	return nil
}

// SetStyleRight sets border style for right side.
func (_adb *border) SetStyleRight(style CellBorderStyle) { _adb._fbae = style }

// TOC represents a table of contents component.
// It consists of a paragraph heading and a collection of
// table of contents lines.
// The representation of a table of contents line is as follows:
//
//	[number] [title]      [separator] [page]
//
// e.g.: Chapter1 Introduction ........... 1
type TOC struct {
	_aeeb  *StyledParagraph
	_gcgf  []*TOCLine
	_daaaa TextStyle
	_gdeab TextStyle
	_dbdd  TextStyle
	_afccg TextStyle
	_fbag  string
	_cbbeg float64
	_ffcde Margins
	_dfdc  Positioning
	_aced  TextStyle
	_bbee  bool
}

// NewPolyBezierCurve creates a new composite Bezier (polybezier) curve.
func (_ebge *Creator) NewPolyBezierCurve(curves []_ff.CubicBezierCurve) *PolyBezierCurve {
	return _afaef(curves)
}

// PolyBezierCurve represents a composite curve that is the result of joining
// multiple cubic Bezier curves.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type PolyBezierCurve struct {
	_ecgf  *_ff.PolyBezierCurve
	_gbbfd float64
	_dagd  float64
	_ggag  Color
}

// Notes returns the notes section of the invoice as a title-content pair.
func (_degg *Invoice) Notes() (string, string) { return _degg._aeee[0], _degg._aeee[1] }

// DrawHeader sets a function to draw a header on created output pages.
func (_cag *Creator) DrawHeader(drawHeaderFunc func(_gfce *Block, _aba HeaderFunctionArgs)) {
	_cag._bebfg = drawHeaderFunc
}

// ColorRGBFrom8bit creates a Color from 8-bit (0-255) r,g,b values.
// Example:
//
//	red := ColorRGBFrom8Bit(255, 0, 0)
func ColorRGBFrom8bit(r, g, b byte) Color {
	return rgbColor{_ddae: float64(r) / 255.0, _cede: float64(g) / 255.0, _fgef: float64(b) / 255.0}
}

func (_ebfgb *templateProcessor) parseBorderRadiusAttr(_cfge, _bdbcg string) (_gacg, _cggg, _fgad, _gbagf float64) {
	_bcd.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0062o\u0072\u0064\u0065r\u0020\u0072\u0061\u0064\u0069\u0075\u0073\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _cfge, _bdbcg)
	switch _fggdd := _a.Fields(_bdbcg); len(_fggdd) {
	case 1:
		_gacg, _ = _aa.ParseFloat(_fggdd[0], 64)
		_cggg = _gacg
		_fgad = _gacg
		_gbagf = _gacg
	case 2:
		_gacg, _ = _aa.ParseFloat(_fggdd[0], 64)
		_fgad = _gacg
		_cggg, _ = _aa.ParseFloat(_fggdd[1], 64)
		_gbagf = _cggg
	case 3:
		_gacg, _ = _aa.ParseFloat(_fggdd[0], 64)
		_cggg, _ = _aa.ParseFloat(_fggdd[1], 64)
		_gbagf = _cggg
		_fgad, _ = _aa.ParseFloat(_fggdd[2], 64)
	case 4:
		_gacg, _ = _aa.ParseFloat(_fggdd[0], 64)
		_cggg, _ = _aa.ParseFloat(_fggdd[1], 64)
		_fgad, _ = _aa.ParseFloat(_fggdd[2], 64)
		_gbagf, _ = _aa.ParseFloat(_fggdd[3], 64)
	}
	return _gacg, _cggg, _fgad, _gbagf
}

func _bdda(_bfeg [][]_ff.Point) *Polygon {
	return &Polygon{_acaag: &_ff.Polygon{Points: _bfeg}, _afdbc: 1.0, _bgac: 1.0}
}

// RadialShading holds information that will be used to render a radial shading.
type RadialShading struct {
	_acdae *shading
	_bcgga *_ga.PdfRectangle
	_bafe  AnchorPoint
	_dfgdc float64
	_acfgc float64
	_fecb  float64
	_dadbe float64
}

// SetBorderOpacity sets the border opacity.
func (_faafa *PolyBezierCurve) SetBorderOpacity(opacity float64) { _faafa._dagd = opacity }

// SetWidth sets line width.
func (_badf *Curve) SetWidth(width float64) { _badf._cdaf = width }

// SetColorRight sets border color for right.
func (_fgb *border) SetColorRight(col Color) { _fgb._gdc = col }

func (_bcfgg *StyledParagraph) split(_befc DrawContext) (_ddbdg, _cgaee *StyledParagraph, _egag error) {
	if _egag = _bcfgg.wrapChunks(false); _egag != nil {
		return nil, nil, _egag
	}
	if len(_bcfgg._bddb) == 1 && _bcfgg._bebe > _befc.Height {
		return _bcfgg, nil, nil
	}
	_cbgf := func(_bbdf []*TextChunk, _gfbd []*TextChunk) []*TextChunk {
		if len(_gfbd) == 0 {
			return _bbdf
		}
		_decec := len(_bbdf)
		if _decec == 0 {
			return append(_bbdf, _gfbd...)
		}
		if _bbdf[_decec-1].Style == _gfbd[0].Style {
			_bbdf[_decec-1].Text += _gfbd[0].Text
		} else {
			_bbdf = append(_bbdf, _gfbd[0])
		}
		return append(_bbdf, _gfbd[1:]...)
	}
	_bddd := func(_aebd *StyledParagraph, _efc []*TextChunk) *StyledParagraph {
		if len(_efc) == 0 {
			return nil
		}
		_cbbe := *_aebd
		_cbbe._eddc = _efc
		return &_cbbe
	}
	var (
		_gdaf  float64
		_fgcgg []*TextChunk
		_gdbd  []*TextChunk
	)
	for _, _gdba := range _bcfgg._bddb {
		var _gdafa float64
		_dagbd := make([]*TextChunk, 0, len(_gdba))
		for _, _dedef := range _gdba {
			if _eefbff := _dedef.Style.FontSize; _eefbff > _gdafa {
				_gdafa = _eefbff
			}
			_dagbd = append(_dagbd, _dedef.clone())
		}
		_gdafa *= _bcfgg._bebe
		if _bcfgg._bedd.IsRelative() {
			if _gdaf+_gdafa > _befc.Height {
				_gdbd = _cbgf(_gdbd, _dagbd)
			} else {
				_fgcgg = _cbgf(_fgcgg, _dagbd)
			}
		}
		_gdaf += _gdafa
	}
	_bcfgg._bddb = nil
	if len(_gdbd) == 0 {
		return _bcfgg, nil, nil
	}
	return _bddd(_bcfgg, _fgcgg), _bddd(_bcfgg, _gdbd), nil
}

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_debb *LinearShading) SetBackgroundColor(backgroundColor Color) {
	_debb._cfbfe.SetBackgroundColor(backgroundColor)
}

// SetBorderRadius sets the radius of the background corners.
func (_cf *Background) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_cf.BorderRadiusTopLeft = topLeft
	_cf.BorderRadiusTopRight = topRight
	_cf.BorderRadiusBottomLeft = bottomLeft
	_cf.BorderRadiusBottomRight = bottomRight
}

// SetStyle sets the style of the line (solid or dashed).
func (_dgggf *Line) SetStyle(style _ff.LineStyle) { _dgggf._ffgbg = style }
func _egafb() *listItem                           { return &listItem{} }
func (_cbgc *Invoice) newColumn(_gaage string, _gdce CellHorizontalAlignment) *InvoiceCell {
	_gefba := &InvoiceCell{_cbgc._bce, _gaage}
	_gefba.Alignment = _gdce
	return _gefba
}

// NewTOC creates a new table of contents.
func (_aaag *Creator) NewTOC(title string) *TOC {
	_dggf := _aaag.NewTextStyle()
	_dggf.Font = _aaag._cbag
	return _cffaa(title, _aaag.NewTextStyle(), _dggf)
}

// GetMargins returns the Image's margins: left, right, top, bottom.
func (_beag *Image) GetMargins() (float64, float64, float64, float64) {
	return _beag._eaac.Left, _beag._eaac.Right, _beag._eaac.Top, _beag._eaac.Bottom
}

// NewTable create a new Table with a specified number of columns.
func (_cccc *Creator) NewTable(cols int) *Table { return _cbgg(cols) }

// SetSubtotal sets the subtotal of the invoice.
func (_caba *Invoice) SetSubtotal(value string) { _caba._gfag[1].Value = value }

func _cbgg(_cfgdg int) *Table {
	_bgcfa := &Table{_edfe: _cfgdg, _bebbc: 10.0, _eeggd: []float64{}, _ecab: []float64{}, _gebec: []*TableCell{}, _bacd: make([]int, _cfgdg), _gbdb: true}
	_bgcfa.resetColumnWidths()
	return _bgcfa
}

// SetLineSeparator sets the separator for all new lines of the table of contents.
func (_efge *TOC) SetLineSeparator(separator string) { _efge._fbag = separator }

// Crop crops the Image to the specified bounds.
func (_gffde *Image) Crop(x0, y0, x1, y1 int) {
	_abb, _fdbc := _gffde._dfad.ToGoImage()
	if _fdbc != nil {
		_df.Fatalf("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074o\u0020\u0047\u006f\u0020\u0049m\u0061\u0067e\u003a\u0020\u0025\u0076", _fdbc)
	}
	var _ffcd _ba.Image
	_ceae := _ba.Rect(x0, y0, x1, y1)
	if _fcgg := _ceae.Intersect(_abb.Bounds()); !_ceae.Empty() {
		_gaaf := _ba.NewRGBA(_ba.Rect(0, 0, _ceae.Dx(), _ceae.Dy()))
		for _gbce := _fcgg.Min.Y; _gbce < _fcgg.Max.Y; _gbce++ {
			for _ddfba := _fcgg.Min.X; _ddfba < _fcgg.Max.X; _ddfba++ {
				_gaaf.Set(_ddfba-_fcgg.Min.X, _gbce-_fcgg.Min.Y, _abb.At(_ddfba, _gbce))
			}
		}
		_ffcd = _gaaf
	} else {
		_ffcd = &_ba.RGBA{}
	}
	_cfdg, _fdbc := _ga.ImageHandling.NewImageFromGoImage(_ffcd)
	if _fdbc != nil {
		_df.Fatalf("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0066\u0072\u006fm\u0020\u0047\u006f\u0020\u0049m\u0061\u0067e\u003a\u0020\u0025\u0076", _fdbc)
	}
	_feaff := float64(_cfdg.Width)
	_dadb := float64(_cfdg.Height)
	_gffde._dfad = _cfdg
	_gffde._gcde = _feaff
	_gffde._eebef = _dadb
	_gffde._ggda = _feaff
	_gffde._cgdb = _dadb
}

// SetFillColor sets the fill color for the path.
func (_acfb *FilledCurve) SetFillColor(color Color) { _acfb._fbgc = color }

// TitleStyle returns the style properties used to render the invoice title.
func (_cceg *Invoice) TitleStyle() TextStyle { return _cceg._adgf }

// AddSubtable copies the cells of the subtable in the table, starting with the
// specified position. The table row and column indices are 1-based, which
// makes the position of the first cell of the first row of the table 1,1.
// The table is automatically extended if the subtable exceeds its columns.
// This can happen when the subtable has more columns than the table or when
// one or more columns of the subtable starting from the specified position
// exceed the last column of the table.
func (_fagd *Table) AddSubtable(row, col int, subtable *Table) {
	for _, _dgfdf := range subtable._gebec {
		_dadga := &TableCell{}
		*_dadga = *_dgfdf
		_dadga._ccbf = _fagd
		_dadga._fcbe += col - 1
		if _gefac := _fagd._edfe - (_dadga._fcbe - 1); _gefac < _dadga._daec {
			_fagd._edfe += _dadga._daec - _gefac
			_fagd.resetColumnWidths()
			_bcd.Log.Debug("\u0054a\u0062l\u0065\u003a\u0020\u0073\u0075\u0062\u0074\u0061\u0062\u006c\u0065 \u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0061\u0062\u006c\u0065\u002e\u0020\u0045\u0078\u0070\u0061\u006e\u0064\u0069\u006e\u0067\u0020\u0074\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0025\u0064\u0020\u0063\u006fl\u0075\u006d\u006e\u0073\u002e", _fagd._edfe)
		}
		_dadga._bgdca += row - 1
		_gggfa := subtable._ecab[_dgfdf._bgdca-1]
		if _dadga._bgdca > _fagd._dbbf {
			for _dadga._bgdca > _fagd._dbbf {
				_fagd._dbbf++
				_fagd._ecab = append(_fagd._ecab, _fagd._bebbc)
			}
			_fagd._ecab[_dadga._bgdca-1] = _gggfa
		} else {
			_fagd._ecab[_dadga._bgdca-1] = _cd.Max(_fagd._ecab[_dadga._bgdca-1], _gggfa)
		}
		_fagd._gebec = append(_fagd._gebec, _dadga)
	}
	_fagd.sortCells()
}

func (_eadf *Paragraph) getMaxLineWidth() float64 {
	if _eadf._aggb == nil || len(_eadf._aggb) == 0 {
		_eadf.wrapText()
	}
	var _cgdea float64
	for _, _abggd := range _eadf._aggb {
		_babc := _eadf.getTextLineWidth(_abggd)
		if _babc > _cgdea {
			_cgdea = _babc
		}
	}
	return _cgdea
}

func (_bfdb *StyledParagraph) wrapChunks(_bfddb bool) error {
	if !_bfdb._bfcgb || int(_bfdb._cccde) <= 0 {
		_bfdb._bddb = [][]*TextChunk{_bfdb._eddc}
		return nil
	}
	if _bfdb._ceba {
		_bfdb.wrapWordChunks()
	}
	_bfdb._bddb = [][]*TextChunk{}
	var _gafg []*TextChunk
	var _ecbf float64
	_ffcf := _bc.IsSpace
	if !_bfddb {
		_ffcf = func(rune) bool { return false }
	}
	_bcee := _baeaf(_bfdb._cccde*1000.0, 0.000001)
	for _, _fdgf := range _bfdb._eddc {
		_fgdg := _fdgf.Style
		_fggb := _fdgf._cgfaa
		_defdg := _fdgf.VerticalAlignment
		var (
			_ceab  []rune
			_cfbcb []float64
		)
		_cddee := _cbea(_fdgf.Text)
		for _, _ebgdd := range _fdgf.Text {
			if _ebgdd == '\u000A' {
				if !_bfddb {
					_ceab = append(_ceab, _ebgdd)
				}
				_gafg = append(_gafg, &TextChunk{Text: _a.TrimRightFunc(string(_ceab), _ffcf), Style: _fgdg, _cgfaa: _dffg(_fggb), VerticalAlignment: _defdg})
				_bfdb._bddb = append(_bfdb._bddb, _gafg)
				_gafg = nil
				_ecbf = 0
				_ceab = nil
				_cfbcb = nil
				continue
			}
			_gadgc := _ebgdd == ' '
			_eafb, _dgbd := _fgdg.Font.GetRuneMetrics(_ebgdd)
			if !_dgbd {
				_bcd.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _ebgdd)
				return _c.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
			}
			_dcbd := _fgdg.FontSize * _eafb.Wx * _fgdg.horizontalScale()
			_afce := _dcbd
			if !_gadgc {
				_afce = _dcbd + _fgdg.CharSpacing*1000.0
			}
			if _ecbf+_dcbd > _bcee {
				_eebbf := -1
				if !_gadgc {
					for _dgbc := len(_ceab) - 1; _dgbc >= 0; _dgbc-- {
						if _ceab[_dgbc] == ' ' {
							_eebbf = _dgbc
							break
						}
					}
				}
				if _bfdb._ceba {
					_gbaf := len(_gafg)
					if _gbaf > 0 {
						_gafg[_gbaf-1].Text = _a.TrimRightFunc(_gafg[_gbaf-1].Text, _ffcf)
						_bfdb._bddb = append(_bfdb._bddb, _gafg)
						_gafg = []*TextChunk{}
					}
					_ceab = append(_ceab, _ebgdd)
					_cfbcb = append(_cfbcb, _afce)
					if _eebbf >= 0 {
						_ceab = _ceab[_eebbf+1:]
						_cfbcb = _cfbcb[_eebbf+1:]
					}
					_ecbf = 0
					for _, _ggbb := range _cfbcb {
						_ecbf += _ggbb
					}
					if _ecbf > _bcee {
						_deag := string(_ceab[:len(_ceab)-1])
						_deag = _gadcc(_deag, _cddee)
						if !_bfddb && _gadgc {
							_deag += "\u0020"
						}
						_gafg = append(_gafg, &TextChunk{Text: _a.TrimRightFunc(_deag, _ffcf), Style: _fgdg, _cgfaa: _dffg(_fggb), VerticalAlignment: _defdg})
						_bfdb._bddb = append(_bfdb._bddb, _gafg)
						_gafg = []*TextChunk{}
						_ceab = []rune{_ebgdd}
						_cfbcb = []float64{_afce}
						_ecbf = _afce
					}
					continue
				}
				_cadcf := string(_ceab)
				if _eebbf >= 0 {
					_cadcf = string(_ceab[0 : _eebbf+1])
					_ceab = _ceab[_eebbf+1:]
					_ceab = append(_ceab, _ebgdd)
					_cfbcb = _cfbcb[_eebbf+1:]
					_cfbcb = append(_cfbcb, _afce)
					_ecbf = 0
					for _, _abfg := range _cfbcb {
						_ecbf += _abfg
					}
				} else {
					if _gadgc {
						_ecbf = 0
						_ceab = []rune{}
						_cfbcb = []float64{}
					} else {
						_ecbf = _afce
						_ceab = []rune{_ebgdd}
						_cfbcb = []float64{_afce}
					}
				}
				_cadcf = _gadcc(_cadcf, _cddee)
				if !_bfddb && _gadgc {
					_cadcf += "\u0020"
				}
				_gafg = append(_gafg, &TextChunk{Text: _a.TrimRightFunc(_cadcf, _ffcf), Style: _fgdg, _cgfaa: _dffg(_fggb), VerticalAlignment: _defdg})
				_bfdb._bddb = append(_bfdb._bddb, _gafg)
				_gafg = []*TextChunk{}
			} else {
				_ecbf += _afce
				_ceab = append(_ceab, _ebgdd)
				_cfbcb = append(_cfbcb, _afce)
			}
		}
		if len(_ceab) > 0 {
			_daea := _gadcc(string(_ceab), _cddee)
			_gafg = append(_gafg, &TextChunk{Text: _daea, Style: _fgdg, _cgfaa: _dffg(_fggb), VerticalAlignment: _defdg})
		}
	}
	if len(_gafg) > 0 {
		_bfdb._bddb = append(_bfdb._bddb, _gafg)
	}
	return nil
}

const (
	PositionRelative Positioning = iota
	PositionAbsolute
)

// SetMargins sets the margins of the rectangle.
// NOTE: rectangle margins are only applied if relative positioning is used.
func (_aeac *Rectangle) SetMargins(left, right, top, bottom float64) {
	_aeac._eaaa.Left = left
	_aeac._eaaa.Right = right
	_aeac._eaaa.Top = top
	_aeac._eaaa.Bottom = bottom
}

// SetLevel sets the indentation level of the TOC line.
func (_eaecb *TOCLine) SetLevel(level uint) {
	_eaecb._dffba = level
	_eaecb._eedce._fadcg.Left = _eaecb._dbdgg + float64(_eaecb._dffba-1)*_eaecb._dgefgc
}

// SetHorizontalAlignment sets the horizontal alignment of the image.
func (_gaagb *Image) SetHorizontalAlignment(alignment HorizontalAlignment) { _gaagb._cccd = alignment }

// Positioning represents the positioning type for drawing creator components (relative/absolute).
type Positioning int

// SetRowHeight sets the height for a specified row.
func (_agga *Table) SetRowHeight(row int, h float64) error {
	if row < 1 || row > len(_agga._ecab) {
		return _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_agga._ecab[row-1] = h
	return nil
}

func (_afagf *templateProcessor) parseBoolAttr(_bfaae, _acccf string) bool {
	_bcd.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0062\u006f\u006f\u006c\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _bfaae, _acccf)
	_eabe, _ := _aa.ParseBool(_acccf)
	return _acccf == "" || _eabe
}

func (_dfcf *Table) wrapContent(_gdagf DrawContext) error {
	if _dfcf._ggea {
		return nil
	}
	_dfcf.sortCells()
	_fbadg := func(_cbcg *TableCell, _ccce int, _fcage int, _bgde int) (_dfaba int) {
		if _bgde < 1 {
			return -1
		}
		_bbfdg := 0
		for _gcdag := _fcage + 1; _gcdag < len(_dfcf._gebec)-1; _gcdag++ {
			_affd := _dfcf._gebec[_gcdag]
			if _affd._bgdca == _bgde && _bbfdg != _fcage {
				_bbfdg = _gcdag
				if (_affd._fcbe < _cbcg._fcbe && _dfcf._edfe > _affd._fcbe) || _cbcg._fcbe < _dfcf._edfe {
					continue
				}
				break
			}
		}
		_fdga := float64(0.0)
		for _fdfgc := 0; _fdfgc < _cbcg._fddc; _fdfgc++ {
			_fdga += _dfcf._ecab[_cbcg._bgdca+_fdfgc-1]
		}
		_bdec := _cbcg.width(_dfcf._eeggd, _gdagf.Width)
		var (
			_fdfaf VectorDrawable
			_bcfb  = false
		)
		switch _efdfb := _cbcg._bfef.(type) {
		case *StyledParagraph:
			_cgbba := _gdagf
			_cgbba.Height = _cd.Floor(_fdga - _efdfb._fadcg.Top - _efdfb._fadcg.Bottom - 0.5*_efdfb.getTextHeight())
			_cgbba.Width = _bdec
			_ggaaa, _ebbc, _dbcbd := _efdfb.split(_cgbba)
			if _dbcbd != nil {
				_bcd.Log.Error("\u0045\u0072\u0072o\u0072\u0020\u0077\u0072a\u0070\u0020\u0073\u0074\u0079\u006c\u0065d\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u003a\u0020\u0025\u0076", _dbcbd.Error())
			}
			if _ggaaa != nil && _ebbc != nil {
				_dfcf._gebec[_fcage]._bfef = _ggaaa
				_fdfaf = _ebbc
				_bcfb = true
			}
		}
		_dfcf._gebec[_fcage]._fddc = _cbcg._fddc
		_gdagf.Height = _gdagf.PageHeight - _gdagf.Margins.Top - _gdagf.Margins.Bottom
		_decgb := _cbcg.cloneProps(nil)
		if _bcfb {
			_decgb._bfef = _fdfaf
		}
		_decgb._fddc = _ccce
		_decgb._bgdca = _bgde + 1
		_decgb._fcbe = _cbcg._fcbe
		if _decgb._bgdca+_decgb._fddc-1 > _dfcf._dbbf {
			for _febb := _dfcf._dbbf; _febb < _decgb._bgdca+_decgb._fddc-1; _febb++ {
				_dfcf._dbbf++
				_dfcf._ecab = append(_dfcf._ecab, _dfcf._bebbc)
			}
		}
		_dfcf._gebec = append(_dfcf._gebec[:_bbfdg+1], append([]*TableCell{_decgb}, _dfcf._gebec[_bbfdg+1:]...)...)
		return _bbfdg + 1
	}
	_aaeg := func(_fdadg *TableCell, _acbd int, _ffcbf int, _edbbf float64) (_aggf int) {
		_gceg := _fdadg.width(_dfcf._eeggd, _gdagf.Width)
		_daaff := _edbbf
		_efef := 1
		_degcc := _gdagf.Height
		if _degcc > 0 {
			for _daaff > _degcc {
				_daaff -= _gdagf.Height
				_degcc = _gdagf.PageHeight - _gdagf.Margins.Top - _gdagf.Margins.Bottom
				_efef++
			}
		}
		var (
			_fbfba VectorDrawable
			_baed  = false
		)
		switch _cddcd := _fdadg._bfef.(type) {
		case *StyledParagraph:
			_aaba := _gdagf
			_aaba.Height = _cd.Floor(_gdagf.Height - _cddcd._fadcg.Top - _cddcd._fadcg.Bottom - 0.5*_cddcd.getTextHeight())
			_aaba.Width = _gceg
			_dfed, _cfafa, _aadd := _cddcd.split(_aaba)
			if _aadd != nil {
				_bcd.Log.Error("\u0045\u0072\u0072o\u0072\u0020\u0077\u0072a\u0070\u0020\u0073\u0074\u0079\u006c\u0065d\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u003a\u0020\u0025\u0076", _aadd.Error())
			}
			if _dfed != nil && _cfafa != nil {
				_dfcf._gebec[_acbd]._bfef = _dfed
				_fbfba = _cfafa
				_baed = true
			}
		}
		if _efef < 2 {
			return -1
		}
		if _dfcf._gebec[_acbd]._bgdca+_efef-1 > _dfcf._dbbf {
			for _gageb := 0; _gageb < _efef; _gageb++ {
				_dfcf._dbbf++
				_dfcf._ecab = append(_dfcf._ecab, _dfcf._bebbc)
			}
		}
		_fccbfe := _edbbf / float64(_efef)
		for _dcgcf := 0; _dcgcf < _efef; _dcgcf++ {
			_dfcf._ecab[_ffcbf+_dcgcf-1] = _fccbfe
		}
		_gdagf.Height = _gdagf.PageHeight - _gdagf.Margins.Top - _gdagf.Margins.Bottom
		_adccf := _fdadg.cloneProps(nil)
		if _baed {
			_adccf._bfef = _fbfba
		}
		_adccf._fddc = 1
		_adccf._bgdca = _ffcbf + _efef - 1
		_adccf._fcbe = _fdadg._fcbe
		_dfcf._gebec = append(_dfcf._gebec, _adccf)
		return len(_dfcf._gebec)
	}
	_bbce := 1
	_eeacb := -1
	for _bgbda := 0; _bgbda < len(_dfcf._gebec); _bgbda++ {
		_afbf := _dfcf._gebec[_bgbda]
		if _eeacb == _bgbda {
			_bbce = _afbf._bgdca
		}
		if _afbf._fddc < 2 {
			if _cdeff := _dfcf._ecab[_afbf._bgdca-1]; _cdeff > _gdagf.Height {
				_eeacb = _aaeg(_afbf, _bgbda, _afbf._bgdca, _cdeff)
				continue
			}
			continue
		}
		_dgfgd := float64(0)
		for _fbge := 0; _fbge < _afbf._fddc; _fbge++ {
			_dgfgd += _dfcf._ecab[_afbf._bgdca+_fbge-1]
		}
		_eggb := float64(0)
		for _dgba := _bbce - 1; _dgba < _afbf._bgdca-1; _dgba++ {
			_eggb += _dfcf._ecab[_dgba]
		}
		if _dgfgd <= (_gdagf.Height - _eggb) {
			continue
		}
		_dfcgf := float64(0.0)
		_ecdcb := _afbf._fddc
		_edbbc := -1
		_gdceb := 1
		for _gcaeb := 1; _gcaeb <= _afbf._fddc; _gcaeb++ {
			if (_dfcgf + _dfcf._ecab[_afbf._bgdca+_gcaeb-2]) > (_gdagf.Height - _eggb) {
				_gdceb--
				break
			}
			_edbbc = _afbf._bgdca + _gcaeb - 1
			_ecdcb = _afbf._fddc - _gcaeb
			_dfcgf += _dfcf._ecab[_afbf._bgdca+_gcaeb-2]
			_gdceb++
		}
		if _afbf._fddc == _ecdcb {
			_gdagf.Height = _gdagf.PageHeight - _gdagf.Margins.Top - _gdagf.Margins.Bottom
			_bbce = _afbf._bgdca
			_bgbda--
			continue
		}
		if _ecdcb > 0 && _afbf._fddc > _gdceb {
			_afbf._fddc = _gdceb
			_eeacb = _fbadg(_afbf, _ecdcb, _bgbda, _edbbc)
			if _bgbda+1 == _eeacb {
				_bgbda--
			}
		}
		_bbce = _afbf._bgdca
	}
	_dfcf.sortCells()
	return nil
}

func (_ad *Block) addContents(_eef *_da.ContentStreamOperations) {
	_ad._fa.WrapIfNeeded()
	_eef.WrapIfNeeded()
	*_ad._fa = append(*_ad._fa, *_eef...)
}

// SetBorderOpacity sets the border opacity of the ellipse.
func (_cada *Ellipse) SetBorderOpacity(opacity float64) { _cada._eade = opacity }

// ScaleToHeight scales the rectangle to the specified height. The width of
// the rectangle is scaled so that the aspect ratio is maintained.
func (_fdefe *Rectangle) ScaleToHeight(h float64) {
	_bgda := _fdefe._bcffg / _fdefe._ebegg
	_fdefe._ebegg = h
	_fdefe._bcffg = h * _bgda
}

// NewGraphicSVGFromFile creates a graphic SVG from a file.
func NewGraphicSVGFromFile(path string) (*GraphicSVG, error) { return _fdfa(path) }

func (_eefff *TextChunk) clone() *TextChunk {
	_agfa := *_eefff
	_agfa._cgfaa = _dffg(_eefff._cgfaa)
	return &_agfa
}

// ScaleToWidth scales the Block to a specified width, maintaining the same aspect ratio.
func (_eg *Block) ScaleToWidth(w float64) { _fag := w / _eg._fd; _eg.Scale(_fag, _fag) }

// NewTextChunk returns a new text chunk instance.
func NewTextChunk(text string, style TextStyle) *TextChunk {
	return &TextChunk{Text: text, Style: style, VerticalAlignment: TextVerticalAlignmentBaseline}
}

// Append adds a new text chunk to the paragraph.
func (_aefa *StyledParagraph) Append(text string) *TextChunk {
	_abaa := NewTextChunk(text, _aefa._afaf)
	return _aefa.appendChunk(_abaa)
}

// SetWidth sets the the Paragraph width. This is essentially the wrapping width, i.e. the width the
// text can extend to prior to wrapping over to next line.
func (_aegef *Paragraph) SetWidth(width float64) { _aegef._adcfe = width; _aegef.wrapText() }

// Height returns the height of the rectangle.
// NOTE: the returned value does not include the border width of the rectangle.
func (_aafa *Rectangle) Height() float64 { return _aafa._ebegg }

// SetMargins sets the margins for the Image (in relative mode): left, right, top, bottom.
func (_faed *Image) SetMargins(left, right, top, bottom float64) {
	_faed._eaac.Left = left
	_faed._eaac.Right = right
	_faed._eaac.Top = top
	_faed._eaac.Bottom = bottom
}

// SetBorderColor sets the border color.
func (_ecbb *PolyBezierCurve) SetBorderColor(color Color) { _ecbb._ecgf.BorderColor = _fce(color) }

// TextStyle is a collection of properties that can be assigned to a text chunk.
type TextStyle struct {
	// Color represents the color of the text.
	Color Color

	// OutlineColor represents the color of the text outline.
	OutlineColor Color

	// Font represents the font the text will use.
	Font *_ga.PdfFont

	// FontSize represents the size of the font.
	FontSize float64

	// OutlineSize represents the thickness of the text outline.
	OutlineSize float64

	// CharSpacing represents the character spacing.
	CharSpacing float64

	// HorizontalScaling represents the percentage to horizontally scale
	// characters by (default: 100). Values less than 100 will result in
	// narrower text while values greater than 100 will result in wider text.
	HorizontalScaling float64

	// RenderingMode represents the rendering mode.
	RenderingMode TextRenderingMode

	// Underline specifies if the text chunk is underlined.
	Underline bool

	// UnderlineStyle represents the style of the line used to underline text.
	UnderlineStyle TextDecorationLineStyle

	// TextRise specifies a vertical adjustment for text. It is useful for
	// drawing subscripts/superscripts. A positive text rise value will
	// produce superscript text, while a negative one will result in
	// subscript text.
	TextRise float64
}

// SetColorTop sets border color for top.
func (_bag *border) SetColorTop(col Color) { _bag._beea = col }

func (_fgbeg *templateProcessor) parseMarginAttr(_gfcdg, _acgef string) Margins {
	_bcd.Log.Debug("\u0050\u0061r\u0073\u0069\u006e\u0067 \u006d\u0061r\u0067\u0069\u006e\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _gfcdg, _acgef)
	_ccbb := Margins{}
	switch _gdfc := _a.Fields(_acgef); len(_gdfc) {
	case 1:
		_ccbb.Top, _ = _aa.ParseFloat(_gdfc[0], 64)
		_ccbb.Bottom = _ccbb.Top
		_ccbb.Left = _ccbb.Top
		_ccbb.Right = _ccbb.Top
	case 2:
		_ccbb.Top, _ = _aa.ParseFloat(_gdfc[0], 64)
		_ccbb.Bottom = _ccbb.Top
		_ccbb.Left, _ = _aa.ParseFloat(_gdfc[1], 64)
		_ccbb.Right = _ccbb.Left
	case 3:
		_ccbb.Top, _ = _aa.ParseFloat(_gdfc[0], 64)
		_ccbb.Left, _ = _aa.ParseFloat(_gdfc[1], 64)
		_ccbb.Right = _ccbb.Left
		_ccbb.Bottom, _ = _aa.ParseFloat(_gdfc[2], 64)
	case 4:
		_ccbb.Top, _ = _aa.ParseFloat(_gdfc[0], 64)
		_ccbb.Right, _ = _aa.ParseFloat(_gdfc[1], 64)
		_ccbb.Bottom, _ = _aa.ParseFloat(_gdfc[2], 64)
		_ccbb.Left, _ = _aa.ParseFloat(_gdfc[3], 64)
	}
	return _ccbb
}

// NewCell makes a new cell and inserts it into the table at the current position.
func (_aefg *Table) NewCell() *TableCell { return _aefg.MultiCell(1, 1) }

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_edfcc *StyledParagraph) SetTextAlignment(align TextAlignment) { _edfcc._dcgfc = align }

// NewImageFromData creates an Image from image data.
func (_gbgf *Creator) NewImageFromData(data []byte) (*Image, error) { return _dgga(data) }

func _ffffg(_dgced *templateProcessor, _geee *templateNode) (interface{}, error) {
	return _dgced.parseBackground(_geee)
}

// SetColumnWidths sets the fractional column widths.
// Each width should be in the range 0-1 and is a fraction of the table width.
// The number of width inputs must match number of columns, otherwise an error is returned.
func (_daege *Table) SetColumnWidths(widths ...float64) error {
	if len(widths) != _daege._edfe {
		_bcd.Log.Debug("M\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0077\u0069\u0064\u0074\u0068\u0073\u0020\u0061nd\u0020\u0063\u006fl\u0075m\u006e\u0073")
		return _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_daege._eeggd = widths
	return nil
}

// GeneratePageBlocks draws the polyline on a new block representing the page.
// Implements the Drawable interface.
func (_ddegb *Polyline) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_efgd := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_ccegg, _cecb := _efgd.setOpacity(_ddegb._cbbd, _ddegb._cbbd)
	if _cecb != nil {
		return nil, ctx, _cecb
	}
	_bdbg := _ddegb._agbgf.Points
	for _eeac := range _bdbg {
		_cfdbc := &_bdbg[_eeac]
		_cfdbc.Y = ctx.PageHeight - _cfdbc.Y
	}
	_afee, _, _cecb := _ddegb._agbgf.Draw(_ccegg)
	if _cecb != nil {
		return nil, ctx, _cecb
	}
	if _cecb = _efgd.addContentsByString(string(_afee)); _cecb != nil {
		return nil, ctx, _cecb
	}
	return []*Block{_efgd}, ctx, nil
}

// Positioning returns the type of positioning the ellipse is set to use.
func (_cged *Ellipse) Positioning() Positioning { return _cged._fabd }
