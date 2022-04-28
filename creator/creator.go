package creator

import (
	_fa "bytes"
	"context"
	_f "errors"
	_ad "fmt"
	_fg "image"
	_de "io"
	_dg "math"
	_c "os"
	_e "sort"
	_cf "strconv"
	_db "strings"
	_a "unicode"

	_da "bitbucket.org/shenghui0779/gopdf/common"
	_fc "bitbucket.org/shenghui0779/gopdf/contentstream"
	_bb "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_ec "bitbucket.org/shenghui0779/gopdf/core"
	_gg "bitbucket.org/shenghui0779/gopdf/internal/integrations/unichart"
	_ba "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_g "bitbucket.org/shenghui0779/gopdf/model"
	_b "github.com/unidoc/unichart/render"
)

// TextOverflow determines the behavior of paragraph text which does
// not fit in the available space.
type TextOverflow int

// SetWidthBottom sets border width for bottom.
func (_fed *border) SetWidthBottom(bw float64) { _fed._gfg = bw }

// SetMargins sets the margins for the Image (in relative mode): left, right, top, bottom.
func (_dgga *Image) SetMargins(left, right, top, bottom float64) {
	_dgga._ecbf.Left = left
	_dgga._ecbf.Right = right
	_dgga._ecbf.Top = top
	_dgga._ecbf.Bottom = bottom
}

// Terms returns the terms and conditions section of the invoice as a
// title-content pair.
func (_fgab *Invoice) Terms() (string, string) { return _fgab._fcega[0], _fgab._fcega[1] }
func _dcb(_ceffd, _fgg, _bcge, _faccb float64) *Ellipse {
	_ffff := &Ellipse{}
	_ffff._gffd = _ceffd
	_ffff._cgae = _fgg
	_ffff._afcf = _bcge
	_ffff._edfg = _faccb
	_ffff._gbad = ColorBlack
	_ffff._gca = 1.0
	return _ffff
}

// CellVerticalAlignment defines the table cell's vertical alignment.
type CellVerticalAlignment int

// Context returns the current drawing context.
func (_daea *Creator) Context() DrawContext { return _daea._defb }
func (_fdf *Chapter) headingNumber() string {
	var _cade string
	if _fdf._aec {
		if _fdf._cdb != 0 {
			_cade = _cf.Itoa(_fdf._cdb) + "\u002e"
		}
		if _fdf._dbbg != nil {
			_fffe := _fdf._dbbg.headingNumber()
			if _fffe != "" {
				_cade = _fffe + _cade
			}
		}
	}
	return _cade
}

// SetColumns overwrites any columns in the line items table. This should be
// called before AddLine.
func (_bada *Invoice) SetColumns(cols []*InvoiceCell) { _bada._cfc = cols }

// GetMargins returns the margins of the TOC line: left, right, top, bottom.
func (_deegb *TOCLine) GetMargins() (float64, float64, float64, float64) {
	_bgddf := &_deegb._daedd._cdbab
	return _deegb._egde, _bgddf.Right, _bgddf.Top, _bgddf.Bottom
}
func _dde(_gbdc *Block, _egab *Image, _dcff DrawContext) (DrawContext, error) {
	_egb := _dcff
	_cgba := 1
	_bbda := _ec.PdfObjectName(_ad.Sprintf("\u0049\u006d\u0067%\u0064", _cgba))
	for _gbdc._fga.HasXObjectByName(_bbda) {
		_cgba++
		_bbda = _ec.PdfObjectName(_ad.Sprintf("\u0049\u006d\u0067%\u0064", _cgba))
	}
	_gaee := _gbdc._fga.SetXObjectImageByName(_bbda, _egab._edad)
	if _gaee != nil {
		return _dcff, _gaee
	}
	_dgbff := 0
	_acff := _ec.PdfObjectName(_ad.Sprintf("\u0047\u0053\u0025\u0064", _dgbff))
	for _gbdc._fga.HasExtGState(_acff) {
		_dgbff++
		_acff = _ec.PdfObjectName(_ad.Sprintf("\u0047\u0053\u0025\u0064", _dgbff))
	}
	_bgca := _ec.MakeDict()
	_bgca.Set("\u0042\u004d", _ec.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	if _egab._ebada < 1.0 {
		_bgca.Set("\u0043\u0041", _ec.MakeFloat(_egab._ebada))
		_bgca.Set("\u0063\u0061", _ec.MakeFloat(_egab._ebada))
	}
	_gaee = _gbdc._fga.AddExtGState(_acff, _ec.MakeIndirectObject(_bgca))
	if _gaee != nil {
		return _dcff, _gaee
	}
	_gee := _egab.Width()
	_dfbgb := _egab.Height()
	_, _cbfg := _egab.rotatedSize()
	_bfdg := _dcff.X
	_cdbd := _dcff.PageHeight - _dcff.Y - _dfbgb
	if _egab._cbce.IsRelative() {
		_cdbd -= (_cbfg - _dfbgb) / 2
		switch _egab._eecd {
		case HorizontalAlignmentCenter:
			_bfdg += (_dcff.Width - _gee) / 2
		case HorizontalAlignmentRight:
			_bfdg = _dcff.PageWidth - _dcff.Margins.Right - _egab._ecbf.Right - _gee
		}
	}
	_fcdd := _egab._dfdb
	_ebffd := _fc.NewContentCreator()
	_ebffd.Add_gs(_acff)
	_ebffd.Translate(_bfdg, _cdbd)
	if _fcdd != 0 {
		_ebffd.Translate(_gee/2, _dfbgb/2)
		_ebffd.RotateDeg(_fcdd)
		_ebffd.Translate(-_gee/2, -_dfbgb/2)
	}
	_ebffd.Scale(_gee, _dfbgb).Add_Do(_bbda)
	_cebf := _ebffd.Operations()
	_cebf.WrapIfNeeded()
	_gbdc.addContents(_cebf)
	if _egab._cbce.IsRelative() {
		_dcff.Y += _cbfg
		_dcff.Height -= _cbfg
		return _dcff, nil
	}
	return _egb, nil
}
func (_bdff *Chapter) headingText() string {
	_fddg := _bdff._afaa
	if _cfd := _bdff.headingNumber(); _cfd != "" {
		_fddg = _ad.Sprintf("\u0025\u0073\u0020%\u0073", _cfd, _fddg)
	}
	return _fddg
}

type listItem struct {
	_dege VectorDrawable
	_bebc TextChunk
}

// AddPage adds the specified page to the creator.
// NOTE: If the page has a Rotate flag, the creator will take care of
// transforming the contents to maintain the correct orientation.
func (_cfdc *Creator) AddPage(page *_g.PdfPage) error {
	_ccbb, _dbead := page.GetMediaBox()
	if _dbead != nil {
		_da.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0070\u0061g\u0065\u0020\u006d\u0065\u0064\u0069\u0061\u0062\u006f\u0078\u003a\u0020\u0025\u0076", _dbead)
		return _dbead
	}
	_ccbb.Normalize()
	_dce, _gbgcd := _ccbb.Llx, _ccbb.Lly
	_gdac := _ba.IdentityMatrix()
	_bcab, _dbead := page.GetRotate()
	if _dbead != nil {
		_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _dbead.Error())
	}
	_gdaa := _bcab%360 != 0 && _bcab%90 == 0
	if _gdaa {
		_dcafg := float64((360 + _bcab%360) % 360)
		if _dcafg == 90 {
			_gdac = _gdac.Translate(_ccbb.Width(), 0)
		} else if _dcafg == 180 {
			_gdac = _gdac.Translate(_ccbb.Width(), _ccbb.Height())
		} else if _dcafg == 270 {
			_gdac = _gdac.Translate(0, _ccbb.Height())
		}
		_gdac = _gdac.Mult(_ba.RotationMatrix(_dcafg * _dg.Pi / 180))
		_gdac = _gdac.Round(0.000001)
		_ega := _bcde(_ccbb, _gdac)
		_ccbb = _ega
		_ccbb.Normalize()
	}
	if _dce != 0 || _gbgcd != 0 {
		_gdac = _ba.TranslationMatrix(_dce, _gbgcd).Mult(_gdac)
	}
	if !_gdac.Identity() {
		_gdac = _gdac.Round(0.000001)
		_cfdc._gff[page] = &pageTransformations{_dgd: &_gdac}
	}
	_cfdc._cgb = _ccbb.Width()
	_cfdc._cggd = _ccbb.Height()
	_cfdc.initContext()
	_cfdc._bdb = append(_cfdc._bdb, page)
	_cfdc._defb.Page++
	return nil
}

// NewChart creates a new creator drawable based on the provided
// unichart chart component.
func NewChart(chart _b.ChartRenderable) *Chart { return _bbf(chart) }

// NewPageBreak create a new page break.
func (_cbff *Creator) NewPageBreak() *PageBreak { return _cacf() }

// NewLine creates a new Line with default parameters between (x1,y1) to (x2,y2).
func (_acdb *Creator) NewLine(x1, y1, x2, y2 float64) *Line { return _dbdac(x1, y1, x2, y2) }

// SetMargins sets the Chapter margins: left, right, top, bottom.
// Typically not needed as the creator's page margins are used.
func (_bdg *Chapter) SetMargins(left, right, top, bottom float64) {
	_bdg._acge.Left = left
	_bdg._acge.Right = right
	_bdg._acge.Top = top
	_bdg._acge.Bottom = bottom
}

// NewStyledTOCLine creates a new table of contents line with the provided style.
func (_cbe *Creator) NewStyledTOCLine(number, title, page TextChunk, level uint, style TextStyle) *TOCLine {
	return _ffbgf(number, title, page, level, style)
}

const (
	PositionRelative Positioning = iota
	PositionAbsolute
)

// SetBorderOpacity sets the border opacity.
func (_ccbc *CurvePolygon) SetBorderOpacity(opacity float64) { _ccbc._dfaa = opacity }

// Color interface represents colors in the PDF creator.
type Color interface {
	ToRGB() (float64, float64, float64)
}

// NewTextStyle creates a new text style object which can be used to style
// chunks of text.
// Default attributes:
// Font: Helvetica
// Font size: 10
// Encoding: WinAnsiEncoding
// Text color: black
func (_gcgbd *Creator) NewTextStyle() TextStyle { return _gadb(_gcgbd._bge) }

// SetEnableWrap sets the line wrapping enabled flag.
func (_gcdb *Paragraph) SetEnableWrap(enableWrap bool) {
	_gcdb._acegd = enableWrap
	_gcdb._dbca = false
}

// PageFinalize sets a function to be called for each page before finalization
// (i.e. the last stage of page processing before they get written out).
// The callback function allows final touch-ups for each page, and it
// provides information that might not be known at other stages of designing
// the document (e.g. the total number of pages). Unlike the header/footer
// functions, which are limited to the top/bottom margins of the page, the
// finalize function can be used draw components anywhere on the current page.
func (_ddgf *Creator) PageFinalize(pageFinalizeFunc func(_ebcc PageFinalizeFunctionArgs) error) {
	_ddgf._fbfgd = pageFinalizeFunc
}

// TextChunk represents a chunk of text along with a particular style.
type TextChunk struct {

	// The text that is being rendered in the PDF.
	Text string

	// The style of the text being rendered.
	Style  TextStyle
	_adbcg *_g.PdfAnnotation
	_cegdd bool
}

// VectorDrawable is a Drawable with a specified width and height.
type VectorDrawable interface {
	Drawable

	// Width returns the width of the Drawable.
	Width() float64

	// Height returns the height of the Drawable.
	Height() float64
}

// MoveDown moves the drawing context down by relative displacement dy (negative goes up).
func (_dbad *Creator) MoveDown(dy float64) { _dbad._defb.Y += dy }

// ConvertToBinary converts current image data into binary (Bi-level image) format.
// If provided image is RGB or GrayScale the function converts it into binary image
// using histogram auto threshold method.
func (_dcfde *Image) ConvertToBinary() error { return _dcfde._ggdgd.ConvertToBinary() }

// TextAlignment options for paragraph.
type TextAlignment int

// SetAnnotation sets a annotation on a TextChunk.
func (_dfcba *TextChunk) SetAnnotation(annotation *_g.PdfAnnotation) { _dfcba._adbcg = annotation }

// SetLink makes the line an internal link.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
func (_gcbce *TOCLine) SetLink(page int64, x, y float64) {
	_gcbce._gddbe = x
	_gcbce._cdde = y
	_gcbce._eegc = page
	_eeag := _gcbce._daedd._aacda.Color
	_gcbce.Number.Style.Color = _eeag
	_gcbce.Title.Style.Color = _eeag
	_gcbce.Separator.Style.Color = _eeag
	_gcbce.Page.Style.Color = _eeag
}

// SetShowLinks sets visibility of links for the TOC lines.
func (_fedea *TOC) SetShowLinks(showLinks bool) { _fedea._bfec = showLinks }
func (_bfefa *Invoice) drawAddress(_ggdeg *InvoiceAddress) []*StyledParagraph {
	var _afeg []*StyledParagraph
	if _ggdeg.Heading != "" {
		_cecc := _fbgb(_bfefa._deef)
		_cecc.SetMargins(0, 0, 0, 7)
		_cecc.Append(_ggdeg.Heading)
		_afeg = append(_afeg, _cecc)
	}
	_gdec := _fbgb(_bfefa._ceadb)
	_gdec.SetLineHeight(1.2)
	_cbad := _ggdeg.Separator
	if _cbad == "" {
		_cbad = _bfefa._cecf
	}
	_beba := _ggdeg.City
	if _ggdeg.State != "" {
		if _beba != "" {
			_beba += _cbad
		}
		_beba += _ggdeg.State
	}
	if _ggdeg.Zip != "" {
		if _beba != "" {
			_beba += _cbad
		}
		_beba += _ggdeg.Zip
	}
	if _ggdeg.Name != "" {
		_gdec.Append(_ggdeg.Name + "\u000a")
	}
	if _ggdeg.Street != "" {
		_gdec.Append(_ggdeg.Street + "\u000a")
	}
	if _ggdeg.Street2 != "" {
		_gdec.Append(_ggdeg.Street2 + "\u000a")
	}
	if _beba != "" {
		_gdec.Append(_beba + "\u000a")
	}
	if _ggdeg.Country != "" {
		_gdec.Append(_ggdeg.Country + "\u000a")
	}
	_gbab := _fbgb(_bfefa._ceadb)
	_gbab.SetLineHeight(1.2)
	_gbab.SetMargins(0, 0, 7, 0)
	if _ggdeg.Phone != "" {
		_gbab.Append(_ggdeg.fmtLine(_ggdeg.Phone, "\u0050h\u006f\u006e\u0065\u003a\u0020", _ggdeg.HidePhoneLabel))
	}
	if _ggdeg.Email != "" {
		_gbab.Append(_ggdeg.fmtLine(_ggdeg.Email, "\u0045m\u0061\u0069\u006c\u003a\u0020", _ggdeg.HideEmailLabel))
	}
	_afeg = append(_afeg, _gdec, _gbab)
	return _afeg
}

// SetLineMargins sets the margins for all new lines of the table of contents.
func (_fccbb *TOC) SetLineMargins(left, right, top, bottom float64) {
	_eada := &_fccbb._efbe
	_eada.Left = left
	_eada.Right = right
	_eada.Top = top
	_eada.Bottom = bottom
}

const (
	CellBorderStyleNone CellBorderStyle = iota
	CellBorderStyleSingle
	CellBorderStyleDouble
)

func _dbba(_ddae float64, _dcfdd float64) float64 { return _dg.Round(_ddae/_dcfdd) * _dcfdd }

// StyledParagraph represents text drawn with a specified font and can wrap across lines and pages.
// By default occupies the available width in the drawing context.
type StyledParagraph struct {
	_dgff  []*TextChunk
	_fgegg TextStyle
	_aacda TextStyle
	_abgbf TextAlignment
	_fgce  TextVerticalAlignment
	_bggb  float64
	_cfga  bool
	_bggd  float64
	_deag  bool
	_cfgc  TextOverflow
	_ffgac float64
	_cdbab Margins
	_dfgbc Positioning
	_aage  float64
	_ebcb  float64
	_acfc  float64
	_fdca  float64
	_aageb [][]*TextChunk
	_bcaf  func(_cgec *StyledParagraph, _cfcg DrawContext)
}

// NewPolyBezierCurve creates a new composite Bezier (polybezier) curve.
func (_dcfd *Creator) NewPolyBezierCurve(curves []_bb.CubicBezierCurve) *PolyBezierCurve {
	return _dffaa(curves)
}

// SetShowNumbering sets a flag to indicate whether or not to show chapter numbers as part of title.
func (_daeg *Chapter) SetShowNumbering(show bool) {
	_daeg._aec = show
	_daeg._gacfc.SetText(_daeg.headingText())
}

// InfoLines returns all the rows in the invoice information table as
// description-value cell pairs.
func (_gcfc *Invoice) InfoLines() [][2]*InvoiceCell {
	_bccb := [][2]*InvoiceCell{_gcfc._aafe, _gcfc._dgdd, _gcfc._gdge}
	return append(_bccb, _gcfc._bgfc...)
}

// SetFillColor sets the fill color.
func (_dabfa *PolyBezierCurve) SetFillColor(color Color) { _dabfa._agfc.FillColor = _afag(color) }

// PageSize represents the page size as a 2 element array representing the width and height in PDF document units (points).
type PageSize [2]float64

// RotateDeg rotates the current active page by angle degrees.  An error is returned on failure,
// which can be if there is no currently active page, or the angleDeg is not a multiple of 90 degrees.
func (_cgea *Creator) RotateDeg(angleDeg int64) error {
	_ddd := _cgea.getActivePage()
	if _ddd == nil {
		_da.Log.Debug("F\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0065\u003a\u0020\u006e\u006f\u0020p\u0061\u0067\u0065\u0020\u0063\u0075\u0072\u0072\u0065\u006etl\u0079\u0020\u0061c\u0074i\u0076\u0065")
		return _f.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	if angleDeg%90 != 0 {
		_da.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067e\u0020\u0072\u006f\u0074\u0061\u0074\u0069on\u0020\u0061\u006e\u0067l\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006dul\u0074\u0069p\u006c\u0065\u0020\u006f\u0066\u0020\u0039\u0030")
		return _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	var _ecgb int64
	if _ddd.Rotate != nil {
		_ecgb = *(_ddd.Rotate)
	}
	_ecgb += angleDeg
	_ddd.Rotate = &_ecgb
	return nil
}

// SetLineTitleStyle sets the style for the title part of all new lines
// of the table of contents.
func (_geaec *TOC) SetLineTitleStyle(style TextStyle) { _geaec._dade = style }
func (_bcda *Division) ctxHeight(_gacab float64) float64 {
	var _edcf float64
	for _, _dgdf := range _bcda._eggb {
		_edcf += _cafb(_dgdf, _gacab)
	}
	return _edcf
}

const (
	TextOverflowVisible TextOverflow = iota
	TextOverflowHidden
)

// SetTitleStyle sets the style properties of the invoice title.
func (_acaaf *Invoice) SetTitleStyle(style TextStyle) { _acaaf._bfgc = style }

// InvoiceCellProps holds all style properties for an invoice cell.
type InvoiceCellProps struct {
	TextStyle       TextStyle
	Alignment       CellHorizontalAlignment
	BackgroundColor Color
	BorderColor     Color
	BorderWidth     float64
	BorderSides     []CellBorderSide
}

// Width returns the width of the chart. In relative positioning mode,
// all the available context width is used at render time.
func (_cdfb *Chart) Width() float64 { return float64(_cdfb._gfgd.Width()) }

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_edae *TOC) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cdbfg := ctx
	_ddec, ctx, _efdf := _edae._ebgb.GeneratePageBlocks(ctx)
	if _efdf != nil {
		return _ddec, ctx, _efdf
	}
	for _, _bafee := range _edae._dbcgf {
		_eaaae := _bafee._eegc
		if !_edae._bfec {
			_bafee._eegc = 0
		}
		_edbdg, _daaa, _daaea := _bafee.GeneratePageBlocks(ctx)
		_bafee._eegc = _eaaae
		if _daaea != nil {
			return _ddec, ctx, _daaea
		}
		if len(_edbdg) < 1 {
			continue
		}
		_ddec[len(_ddec)-1].mergeBlocks(_edbdg[0])
		_ddec = append(_ddec, _edbdg[1:]...)
		ctx = _daaa
	}
	if _edae._adeee.IsRelative() {
		ctx.X = _cdbfg.X
	}
	if _edae._adeee.IsAbsolute() {
		return _ddec, _cdbfg, nil
	}
	return _ddec, ctx, nil
}

const (
	TextVerticalAlignmentBaseline TextVerticalAlignment = iota
	TextVerticalAlignmentCenter
)

func (_cgg *Block) addContents(_ef *_fc.ContentStreamOperations) {
	_cgg._cb.WrapIfNeeded()
	_ef.WrapIfNeeded()
	*_cgg._cb = append(*_cgg._cb, *_ef...)
}

// SetInline sets the inline mode of the division.
func (_adeg *Division) SetInline(inline bool) { _adeg._abce = inline }

// Add appends a new item to the list.
// The supported components are: *Paragraph, *StyledParagraph and *List.
// Returns the marker used for the newly added item. The returned marker
// object can be used to change the text and style of the marker for the
// current item.
func (_gfcc *List) Add(item VectorDrawable) (*TextChunk, error) {
	_edbd := &listItem{_dege: item, _bebc: _gfcc._ddcde}
	switch _gbfd := item.(type) {
	case *Paragraph:
	case *StyledParagraph:
	case *List:
		if _gbfd._dcffe {
			_gbfd._fefc = 15
		}
	default:
		return nil, _f.New("\u0074\u0068i\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u006c\u0069\u0073\u0074")
	}
	_gfcc._cbba = append(_gfcc._cbba, _edbd)
	return &_edbd._bebc, nil
}

const (
	DefaultHorizontalScaling = 100
)

// NewImage create a new image from a unidoc image (model.Image).
func (_cedcb *Creator) NewImage(img *_g.Image) (*Image, error) { return _ageb(img) }
func (_ceccg *Invoice) drawInformation() *Table {
	_adga := _dbfc(2)
	_dffc := append([][2]*InvoiceCell{_ceccg._aafe, _ceccg._dgdd, _ceccg._gdge}, _ceccg._bgfc...)
	for _, _efb := range _dffc {
		_caca, _gcbfd := _efb[0], _efb[1]
		if _gcbfd.Value == "" {
			continue
		}
		_ccfe := _adga.NewCell()
		_ccfe.SetBackgroundColor(_caca.BackgroundColor)
		_ceccg.setCellBorder(_ccfe, _caca)
		_bfbe := _fbgb(_caca.TextStyle)
		_bfbe.Append(_caca.Value)
		_bfbe.SetMargins(0, 0, 2, 1)
		_ccfe.SetContent(_bfbe)
		_ccfe = _adga.NewCell()
		_ccfe.SetBackgroundColor(_gcbfd.BackgroundColor)
		_ceccg.setCellBorder(_ccfe, _gcbfd)
		_bfbe = _fbgb(_gcbfd.TextStyle)
		_bfbe.Append(_gcbfd.Value)
		_bfbe.SetMargins(0, 0, 2, 1)
		_ccfe.SetContent(_bfbe)
	}
	return _adga
}

// SetWidthTop sets border width for top.
func (_cac *border) SetWidthTop(bw float64) { _cac._bdf = bw }

const (
	CellHorizontalAlignmentLeft CellHorizontalAlignment = iota
	CellHorizontalAlignmentCenter
	CellHorizontalAlignmentRight
)

// Margins returns the margins of the list: left, right, top, bottom.
func (_ccae *List) Margins() (float64, float64, float64, float64) {
	return _ccae._cgab.Left, _ccae._cgab.Right, _ccae._cgab.Top, _ccae._cgab.Bottom
}

// Rows returns the total number of rows the table has.
func (_fgde *Table) Rows() int { return _fgde._gbcf }

// PageBreak represents a page break for a chapter.
type PageBreak struct{}

// SetStyleRight sets border style for right side.
func (_dfb *border) SetStyleRight(style CellBorderStyle) { _dfb._gga = style }

// SetPos sets the position of the chart to the specified coordinates.
// This method sets the chart to use absolute positioning.
func (_ffca *Chart) SetPos(x, y float64) {
	_ffca._cbfa = PositionAbsolute
	_ffca._ebfbg = x
	_ffca._afe = y
}

// NewChapter creates a new chapter with the specified title as the heading.
func (_acf *Creator) NewChapter(title string) *Chapter {
	_acf._agb++
	_ccbg := _acf.NewTextStyle()
	_ccbg.FontSize = 16
	return _gbg(nil, _acf._fgc, _acf._edge, title, _acf._agb, _ccbg)
}

// SetEnableWrap sets the line wrapping enabled flag.
func (_bfac *StyledParagraph) SetEnableWrap(enableWrap bool) {
	_bfac._cfga = enableWrap
	_bfac._deag = false
}

// SetRowHeight sets the height for a specified row.
func (_aege *Table) SetRowHeight(row int, h float64) error {
	if row < 1 || row > len(_aege._abeg) {
		return _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_aege._abeg[row-1] = h
	return nil
}

// NewInvoice returns an instance of an empty invoice.
func (_bcbg *Creator) NewInvoice() *Invoice {
	_gdaf := _bcbg.NewTextStyle()
	_gdaf.Font = _bcbg._dcaf
	return _bdgb(_bcbg.NewTextStyle(), _gdaf)
}
func _ebccg(_adeda []_bb.Point) *Polyline {
	return &Polyline{_efcf: &_bb.Polyline{Points: _adeda, LineColor: _g.NewPdfColorDeviceRGB(0, 0, 0), LineWidth: 1.0}, _beccb: 1.0}
}
func (_gbdce *StyledParagraph) getTextWidth() float64 {
	var _abad float64
	_bffgd := len(_gbdce._dgff)
	for _aaeee, _cece := range _gbdce._dgff {
		_adfg := &_cece.Style
		_ecfc := len(_cece.Text)
		for _edeb, _ddeaf := range _cece.Text {
			if _ddeaf == '\u000A' {
				continue
			}
			_gfagd, _efbc := _adfg.Font.GetRuneMetrics(_ddeaf)
			if !_efbc {
				_da.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _ddeaf)
				return -1
			}
			_abad += _adfg.FontSize * _gfagd.Wx * _adfg.horizontalScale()
			if _ddeaf != ' ' && (_aaeee != _bffgd-1 || _edeb != _ecfc-1) {
				_abad += _adfg.CharSpacing * 1000.0
			}
		}
	}
	return _abad
}

// NewTOCLine creates a new table of contents line with the default style.
func (_fage *Creator) NewTOCLine(number, title, page string, level uint) *TOCLine {
	return _dgcda(number, title, page, level, _fage.NewTextStyle())
}

// EnableRowWrap controls whether rows are wrapped across pages.
// NOTE: Currently, row wrapping is supported for rows using StyledParagraphs.
func (_agff *Table) EnableRowWrap(enable bool) { _agff._adce = enable }
func _edfgc(_aggfc [][]_bb.Point) *Polygon {
	return &Polygon{_fdfa: &_bb.Polygon{Points: _aggfc}, _aebd: 1.0, _daae: 1.0}
}

// SetStyleTop sets border style for top side.
func (_def *border) SetStyleTop(style CellBorderStyle) { _def._ddc = style }

// FilledCurve represents a closed path of Bezier curves with a border and fill.
type FilledCurve struct {
	_dgfb         []_bb.CubicBezierCurve
	FillEnabled   bool
	_debe         Color
	BorderEnabled bool
	BorderWidth   float64
	_dddbg        Color
}

// SetAngle sets Image rotation angle in degrees.
func (_gag *Image) SetAngle(angle float64) { _gag._dfdb = angle }

// SetMargins sets the Paragraph's margins.
func (_ffgeb *StyledParagraph) SetMargins(left, right, top, bottom float64) {
	_ffgeb._cdbab.Left = left
	_ffgeb._cdbab.Right = right
	_ffgeb._cdbab.Top = top
	_ffgeb._cdbab.Bottom = bottom
}

// Ellipse defines an ellipse with a center at (xc,yc) and a specified width and height.  The ellipse can have a colored
// fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Ellipse struct {
	_gffd float64
	_cgae float64
	_afcf float64
	_edfg float64
	_fcaa Color
	_gbad Color
	_gca  float64
}

// Cols returns the total number of columns the table has.
func (_cggbc *Table) Cols() int { return _cggbc._fabb }

// Scale scales Image by a constant factor, both width and height.
func (_fcgb *Image) Scale(xFactor, yFactor float64) {
	_fcgb._ebb = xFactor * _fcgb._ebb
	_fcgb._gcab = yFactor * _fcgb._gcab
}

// SetBorderColor sets the border color.
func (_ecgcd *PolyBezierCurve) SetBorderColor(color Color) { _ecgcd._agfc.BorderColor = _afag(color) }
func (_fgd *Block) transform(_ee _ba.Matrix) {
	_bga := _fc.NewContentCreator().Add_cm(_ee[0], _ee[1], _ee[3], _ee[4], _ee[6], _ee[7]).Operations()
	*_fgd._cb = append(*_bga, *_fgd._cb...)
	_fgd._cb.WrapIfNeeded()
}

// SetBorderColor sets border color.
func (_gddg *Rectangle) SetBorderColor(col Color) { _gddg._fdegf = col }

// SetFillColor sets the fill color.
func (_cfa *Ellipse) SetFillColor(col Color) { _cfa._fcaa = col }

// SetTitle sets the title of the invoice.
func (_befg *Invoice) SetTitle(title string) { _befg._cbde = title }

// MoveTo moves the drawing context to absolute coordinates (x, y).
func (_bcbd *Creator) MoveTo(x, y float64) { _bcbd._defb.X = x; _bcbd._defb.Y = y }
func (_dgbe *Creator) getActivePage() *_g.PdfPage {
	if _dgbe._abed == nil {
		if len(_dgbe._bdb) == 0 {
			return nil
		}
		return _dgbe._bdb[len(_dgbe._bdb)-1]
	}
	return _dgbe._abed
}
func (_eabaa *Paragraph) getMaxLineWidth() float64 {
	if _eabaa._egge == nil || len(_eabaa._egge) == 0 {
		_eabaa.wrapText()
	}
	var _fedc float64
	for _, _dbeada := range _eabaa._egge {
		_cefg := _eabaa.getTextLineWidth(_dbeada)
		if _cefg > _fedc {
			_fedc = _cefg
		}
	}
	return _fedc
}

// NewPolyline creates a new polyline.
func (_aaeb *Creator) NewPolyline(points []_bb.Point) *Polyline { return _ebccg(points) }

// DrawWithContext draws the Block using the specified drawing context.
func (_bcd *Block) DrawWithContext(d Drawable, ctx DrawContext) error {
	_fdd, _, _bae := d.GeneratePageBlocks(ctx)
	if _bae != nil {
		return _bae
	}
	if len(_fdd) != 1 {
		return _f.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0062\u006c\u006f\u0063k\u0073")
	}
	for _, _cge := range _fdd {
		if _eaf := _bcd.mergeBlocks(_cge); _eaf != nil {
			return _eaf
		}
	}
	return nil
}

// Line defines a line between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Line struct {
	_baba  float64
	_fbb   float64
	_ggadb float64
	_bgbc  float64
	_cbgb  Color
	_bbgc  float64
}

// SetLineWidth sets the line width.
func (_fcddc *Line) SetLineWidth(lw float64) { _fcddc._bbgc = lw }
func _eabf() *FilledCurve {
	_gebfc := FilledCurve{}
	_gebfc._dgfb = []_bb.CubicBezierCurve{}
	return &_gebfc
}
func _aacce(_badd []byte) (*Image, error) {
	_agceg := _fa.NewReader(_badd)
	_ecdc, _dgbf := _g.ImageHandling.Read(_agceg)
	if _dgbf != nil {
		_da.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _dgbf)
		return nil, _dgbf
	}
	return _ageb(_ecdc)
}
func _ageb(_ggaf *_g.Image) (*Image, error) {
	_cbb := float64(_ggaf.Width)
	_bafdf := float64(_ggaf.Height)
	return &Image{_ggdgd: _ggaf, _efeb: _cbb, _aegc: _bafdf, _ebb: _cbb, _gcab: _bafdf, _dfdb: 0, _ebada: 1.0, _cbce: PositionRelative}, nil
}

// SetWidthLeft sets border width for left.
func (_aac *border) SetWidthLeft(bw float64) { _aac._age = bw }

// GetCoords returns coordinates of border.
func (_abe *border) GetCoords() (float64, float64) { return _abe._edd, _abe._ccga }

// SetStyleBottom sets border style for bottom side.
func (_dbea *border) SetStyleBottom(style CellBorderStyle) { _dbea._bedg = style }
func _afag(_gfbf Color) _g.PdfColor {
	if _gfbf == nil {
		_gfbf = ColorBlack
	}
	switch _ggcd := _gfbf.(type) {
	case cmykColor:
		return _g.NewPdfColorDeviceCMYK(_ggcd._debb, _ggcd._caad, _ggcd._gfbg, _ggcd._abf)
	}
	return _g.NewPdfColorDeviceRGB(_gfbf.ToRGB())
}

// NewImageFromGoImage creates an Image from a go image.Image data structure.
func (_ebff *Creator) NewImageFromGoImage(goimg _fg.Image) (*Image, error) { return _beebb(goimg) }
func (_dbgc *Paragraph) getTextWidth() float64 {
	_gcedg := 0.0
	for _, _ebag := range _dbgc._bfagd {
		if _ebag == '\u000A' {
			continue
		}
		_dcbe, _eggc := _dbgc._cgbf.GetRuneMetrics(_ebag)
		if !_eggc {
			_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _ebag, _ebag)
			return -1
		}
		_gcedg += _dbgc._ebbae * _dcbe.Wx
	}
	return _gcedg
}

// GetHorizontalAlignment returns the horizontal alignment of the image.
func (_aafb *Image) GetHorizontalAlignment() HorizontalAlignment { return _aafb._eecd }

// SetWidth set the Image's document width to specified w. This does not change the raw image data, i.e.
// no actual scaling of data is performed. That is handled by the PDF viewer.
func (_fbe *Image) SetWidth(w float64) { _fbe._ebb = w }

// SetBorderColor sets the border color.
func (_gafa *Polygon) SetBorderColor(color Color) { _gafa._fdfa.BorderColor = _afag(color) }
func _bcde(_fcgcb *_g.PdfRectangle, _egad _ba.Matrix) *_g.PdfRectangle {
	var _bbfac _g.PdfRectangle
	_bbfac.Llx, _bbfac.Lly = _egad.Transform(_fcgcb.Llx, _fcgcb.Lly)
	_bbfac.Urx, _bbfac.Ury = _egad.Transform(_fcgcb.Urx, _fcgcb.Ury)
	_bbfac.Normalize()
	return &_bbfac
}

// GeneratePageBlocks generates the table page blocks. Multiple blocks are
// generated if the contents wrap over multiple pages.
// Implements the Drawable interface.
func (_afgaf *Table) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bgec := _afgaf
	if _afgaf._adce {
		_bgec = _afgaf.clone()
	}
	return _cgfbc(_bgec, ctx)
}

// Creator is a wrapper around functionality for creating PDF reports and/or adding new
// content onto imported PDF pages, etc.
type Creator struct {

	// Errors keeps error messages that should not interrupt pdf processing and to be checked later.
	Errors []error

	// UnsupportedCharacterReplacement is character that will be used to replace unsupported glyph.
	// The value will be passed to drawing context.
	UnsupportedCharacterReplacement rune
	_bdb                            []*_g.PdfPage
	_fcac                           map[*_g.PdfPage]*Block
	_gff                            map[*_g.PdfPage]*pageTransformations
	_abed                           *_g.PdfPage
	_ebad                           PageSize
	_defb                           DrawContext
	_acda                           Margins
	_cgb, _cggd                     float64
	_agb                            int
	_gace                           func(_ggdb FrontpageFunctionArgs)
	_bbfa                           func(_gbc *TOC) error
	_gce                            func(_bbga *Block, _cbdba HeaderFunctionArgs)
	_bfag                           func(_dagc *Block, _dfbd FooterFunctionArgs)
	_fbfgd                          func(_fagcd PageFinalizeFunctionArgs) error
	_aaee                           func(_gacb *_g.PdfWriter) error
	_ggce                           bool

	// Controls whether a table of contents will be generated.
	AddTOC bool
	_fgc   *TOC

	// Controls whether outlines will be generated.
	AddOutlines bool
	_edge       *_g.Outline
	_cadb       *_g.PdfOutlineTreeNode
	_dfce       *_g.PdfAcroForm
	_bdgf       _ec.PdfObject
	_edgd       _g.Optimizer
	_dbc        []*_g.PdfFont
	_bge        *_g.PdfFont
	_dcaf       *_g.PdfFont
}

func (_cga *Block) translate(_gae, _bab float64) {
	_bca := _fc.NewContentCreator().Translate(_gae, -_bab).Operations()
	*_cga._cb = append(*_bca, *_cga._cb...)
	_cga._cb.WrapIfNeeded()
}

// SetFontSize sets the font size in document units (points).
func (_ccef *Paragraph) SetFontSize(fontSize float64)  { _ccef._ebbae = fontSize }
func (_eeae *Creator) setActivePage(_gebe *_g.PdfPage) { _eeae._abed = _gebe }

// SetColor sets the line color.
// Use ColorRGBFromHex, ColorRGBFrom8bit or ColorRGBFromArithmetic to make the color object.
func (_bddb *Line) SetColor(col Color) { _bddb._cbgb = col }

// DrawHeader sets a function to draw a header on created output pages.
func (_cdff *Creator) DrawHeader(drawHeaderFunc func(_fcb *Block, _dda HeaderFunctionArgs)) {
	_cdff._gce = drawHeaderFunc
}

// Height returns the height of the division, assuming all components are
// stacked on top of each other.
func (_acfac *Division) Height() float64 {
	var _gdafe float64
	for _, _fgca := range _acfac._eggb {
		switch _bbdcc := _fgca.(type) {
		case marginDrawable:
			_, _, _ccgaa, _afbc := _bbdcc.GetMargins()
			_gdafe += _bbdcc.Height() + _ccgaa + _afbc
		default:
			_gdafe += _bbdcc.Height()
		}
	}
	return _gdafe
}

// CreateTableOfContents sets a function to generate table of contents.
func (_edc *Creator) CreateTableOfContents(genTOCFunc func(_gcbf *TOC) error) {
	_edc._bbfa = genTOCFunc
}

// Height returns the total height of all rows.
func (_bfdb *Table) Height() float64 {
	_gbgag := float64(0.0)
	for _, _edefg := range _bfdb._abeg {
		_gbgag += _edefg
	}
	return _gbgag
}

// TitleStyle returns the style properties used to render the invoice title.
func (_edfe *Invoice) TitleStyle() TextStyle { return _edfe._bfgc }

// NewStyledParagraph creates a new styled paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_aded *Creator) NewStyledParagraph() *StyledParagraph { return _fbgb(_aded.NewTextStyle()) }

type rgbColor struct{ _aacc, _beeb, _fgea float64 }

func (_baede *Division) split(_cddb DrawContext) (_bedd, _efgd *Division) {
	var _cdae float64
	var _egda, _bede []VectorDrawable
	for _agbc, _baa := range _baede._eggb {
		_cdae += _cafb(_baa, _cddb.Width)
		if _cdae < _cddb.Height {
			_egda = append(_egda, _baa)
		} else {
			_bede = _baede._eggb[_agbc:]
			break
		}
	}
	if len(_egda) > 0 {
		_bedd = _eded()
		_bedd._eggb = _egda
	}
	if len(_bede) > 0 {
		_efgd = _eded()
		_efgd._eggb = _bede
	}
	return _bedd, _efgd
}

// Height returns Image's document height.
func (_dbfg *Image) Height() float64 { return _dbfg._gcab }
func (_fafbb *StyledParagraph) split(_dcadd DrawContext) (_gecg, _bcad *StyledParagraph, _badc error) {
	if _badc = _fafbb.wrapChunks(false); _badc != nil {
		return nil, nil, _badc
	}
	_afbcb := func(_gcda []*TextChunk, _abdc []*TextChunk) []*TextChunk {
		if len(_abdc) == 0 {
			return _gcda
		}
		_cfff := len(_gcda)
		if _cfff == 0 {
			return append(_gcda, _abdc...)
		}
		_gcda[_cfff-1].Text += _abdc[0].Text
		return append(_gcda, _abdc[1:]...)
	}
	_gcabc := func(_gfdc *StyledParagraph, _egdcb []*TextChunk) *StyledParagraph {
		if len(_egdcb) == 0 {
			return nil
		}
		_bddba := *_gfdc
		_bddba._dgff = _egdcb
		return &_bddba
	}
	var (
		_deaa  float64
		_bcafe []*TextChunk
		_gdfg  []*TextChunk
	)
	for _, _ggeb := range _fafbb._aageb {
		var _bccdc float64
		_aedg := make([]*TextChunk, 0, len(_ggeb))
		for _, _fgge := range _ggeb {
			if _accc := _fgge.Style.FontSize; _accc > _bccdc {
				_bccdc = _accc
			}
			_aedg = append(_aedg, _fgge.clone())
		}
		_bccdc *= _fafbb._bggb
		if _fafbb._dfgbc.IsRelative() {
			if _deaa+_bccdc > _dcadd.Height {
				_gdfg = _afbcb(_gdfg, _aedg)
			} else {
				_bcafe = _afbcb(_bcafe, _aedg)
			}
		}
		_deaa += _bccdc
	}
	_fafbb._aageb = nil
	if len(_gdfg) == 0 {
		return _fafbb, nil, nil
	}
	return _gcabc(_fafbb, _bcafe), _gcabc(_fafbb, _gdfg), nil
}

// IsAbsolute checks if the positioning is absolute.
func (_cdacb Positioning) IsAbsolute() bool { return _cdacb == PositionAbsolute }

// Width returns the width of the Paragraph.
func (_gabba *StyledParagraph) Width() float64 {
	if _gabba._cfga && int(_gabba._bggd) > 0 {
		return _gabba._bggd
	}
	return _gabba.getTextWidth() / 1000.0
}
func _gadb(_dgcb *_g.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 0), Font: _dgcb, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}

// Margins represents page margins or margins around an element.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (_gbgd *Invoice) generateNoteBlocks(_gacbeb DrawContext) ([]*Block, DrawContext, error) {
	_dfcc := _eded()
	_gaba := append([][2]string{_gbgd._bgcb, _gbgd._fcega}, _gbgd._facf...)
	for _, _deaf := range _gaba {
		if _deaf[1] != "" {
			_bgdb := _gbgd.drawSection(_deaf[0], _deaf[1])
			for _, _cdec := range _bgdb {
				_dfcc.Add(_cdec)
			}
			_aaea := _fbgb(_gbgd._fafbeg)
			_aaea.SetMargins(0, 0, 10, 0)
			_dfcc.Add(_aaea)
		}
	}
	return _dfcc.GeneratePageBlocks(_gacbeb)
}

// GeneratePageBlocks generates the page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages. Implements the Drawable interface.
func (_adfc *Paragraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cfce := ctx
	var _fbgf []*Block
	_ggga := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _adfc._fcgd.IsRelative() {
		ctx.X += _adfc._bbdee.Left
		ctx.Y += _adfc._bbdee.Top
		ctx.Width -= _adfc._bbdee.Left + _adfc._bbdee.Right
		ctx.Height -= _adfc._bbdee.Top
		_adfc.SetWidth(ctx.Width)
		if _adfc.Height() > ctx.Height {
			_fbgf = append(_fbgf, _ggga)
			_ggga = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_geaad := ctx
			_geaad.Y = ctx.Margins.Top
			_geaad.X = ctx.Margins.Left + _adfc._bbdee.Left
			_geaad.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
			_geaad.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _adfc._bbdee.Left - _adfc._bbdee.Right
			ctx = _geaad
		}
	} else {
		if int(_adfc._dgee) <= 0 {
			_adfc.SetWidth(_adfc.getTextWidth())
		}
		ctx.X = _adfc._fccc
		ctx.Y = _adfc._ddda
	}
	ctx, _bbdea := _dedcc(_ggga, _adfc, ctx)
	if _bbdea != nil {
		_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bbdea)
		return nil, ctx, _bbdea
	}
	_fbgf = append(_fbgf, _ggga)
	if _adfc._fcgd.IsRelative() {
		ctx.Y += _adfc._bbdee.Bottom
		ctx.Height -= _adfc._bbdee.Bottom
		if !ctx.Inline {
			ctx.X = _cfce.X
			ctx.Width = _cfce.Width
		}
		return _fbgf, ctx, nil
	}
	return _fbgf, _cfce, nil
}

// SetIndent sets the left offset of the list when nested into another list.
func (_fbag *List) SetIndent(indent float64) { _fbag._fefc = indent; _fbag._dcffe = false }

// NewCurve returns new instance of Curve between points (x1,y1) and (x2, y2) with control point (cx,cy).
func (_bddaf *Creator) NewCurve(x1, y1, cx, cy, x2, y2 float64) *Curve {
	return _bggg(x1, y1, cx, cy, x2, y2)
}

// GetMargins returns the Image's margins: left, right, top, bottom.
func (_eece *Image) GetMargins() (float64, float64, float64, float64) {
	return _eece._ecbf.Left, _eece._ecbf.Right, _eece._ecbf.Top, _eece._ecbf.Bottom
}

// SetBorderColor sets the border color.
func (_ggad *CurvePolygon) SetBorderColor(color Color) { _ggad._bafd.BorderColor = _afag(color) }

// NoteHeadingStyle returns the style properties used to render the heading of
// the invoice note sections.
func (_dad *Invoice) NoteHeadingStyle() TextStyle { return _dad._dedb }

// FrontpageFunctionArgs holds the input arguments to a front page drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FrontpageFunctionArgs struct {
	PageNum    int
	TotalPages int
}

func _bdgb(_feea, _cddf TextStyle) *Invoice {
	_bcbf := &Invoice{_cbde: "\u0049N\u0056\u004f\u0049\u0043\u0045", _cecf: "\u002c\u0020", _fafbeg: _feea, _ggae: _cddf}
	_bcbf._aedd = &InvoiceAddress{Separator: _bcbf._cecf}
	_bcbf._dggaa = &InvoiceAddress{Heading: "\u0042i\u006c\u006c\u0020\u0074\u006f", Separator: _bcbf._cecf}
	_afdd := ColorRGBFrom8bit(245, 245, 245)
	_gcbeg := ColorRGBFrom8bit(155, 155, 155)
	_bcbf._bfgc = _cddf
	_bcbf._bfgc.Color = _gcbeg
	_bcbf._bfgc.FontSize = 20
	_bcbf._ceadb = _feea
	_bcbf._deef = _cddf
	_bcbf._efec = _feea
	_bcbf._dedb = _cddf
	_bcbf._dbeg = _bcbf.NewCellProps()
	_bcbf._dbeg.BackgroundColor = _afdd
	_bcbf._dbeg.TextStyle = _cddf
	_bcbf._fbgc = _bcbf.NewCellProps()
	_bcbf._fbgc.TextStyle = _cddf
	_bcbf._fbgc.BackgroundColor = _afdd
	_bcbf._fbgc.BorderColor = _afdd
	_bcbf._cbcee = _bcbf.NewCellProps()
	_bcbf._cbcee.BorderColor = _afdd
	_bcbf._cbcee.BorderSides = []CellBorderSide{CellBorderSideBottom}
	_bcbf._cbcee.Alignment = CellHorizontalAlignmentRight
	_bcbf._ddcd = _bcbf.NewCellProps()
	_bcbf._ddcd.Alignment = CellHorizontalAlignmentRight
	_bcbf._aafe = [2]*InvoiceCell{_bcbf.newCell("\u0049\u006e\u0076\u006f\u0069\u0063\u0065\u0020\u006eu\u006d\u0062\u0065\u0072", _bcbf._dbeg), _bcbf.newCell("", _bcbf._dbeg)}
	_bcbf._dgdd = [2]*InvoiceCell{_bcbf.newCell("\u0044\u0061\u0074\u0065", _bcbf._dbeg), _bcbf.newCell("", _bcbf._dbeg)}
	_bcbf._gdge = [2]*InvoiceCell{_bcbf.newCell("\u0044\u0075\u0065\u0020\u0044\u0061\u0074\u0065", _bcbf._dbeg), _bcbf.newCell("", _bcbf._dbeg)}
	_bcbf._afga = [2]*InvoiceCell{_bcbf.newCell("\u0053\u0075\u0062\u0074\u006f\u0074\u0061\u006c", _bcbf._ddcd), _bcbf.newCell("", _bcbf._ddcd)}
	_accbd := _bcbf._ddcd
	_accbd.TextStyle = _cddf
	_accbd.BackgroundColor = _afdd
	_accbd.BorderColor = _afdd
	_bcbf._cdgb = [2]*InvoiceCell{_bcbf.newCell("\u0054\u006f\u0074a\u006c", _accbd), _bcbf.newCell("", _accbd)}
	_bcbf._bgcb = [2]string{"\u004e\u006f\u0074e\u0073", ""}
	_bcbf._fcega = [2]string{"T\u0065r\u006d\u0073\u0020\u0061\u006e\u0064\u0020\u0063o\u006e\u0064\u0069\u0074io\u006e\u0073", ""}
	_bcbf._cfc = []*InvoiceCell{_bcbf.newColumn("D\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", CellHorizontalAlignmentLeft), _bcbf.newColumn("\u0051\u0075\u0061\u006e\u0074\u0069\u0074\u0079", CellHorizontalAlignmentRight), _bcbf.newColumn("\u0055\u006e\u0069\u0074\u0020\u0070\u0072\u0069\u0063\u0065", CellHorizontalAlignmentRight), _bcbf.newColumn("\u0041\u006d\u006f\u0075\u006e\u0074", CellHorizontalAlignmentRight)}
	return _bcbf
}

// SetFillColor sets the fill color.
func (_fbad *Polygon) SetFillColor(color Color) { _fbad._fdfa.FillColor = _afag(color) }

// SetTotal sets the total of the invoice.
func (_cdfc *Invoice) SetTotal(value string) { _cdfc._cdgb[1].Value = value }

// Notes returns the notes section of the invoice as a title-content pair.
func (_ffbg *Invoice) Notes() (string, string) { return _ffbg._bgcb[0], _ffbg._bgcb[1] }

// GetOptimizer returns current PDF optimizer.
func (_bfd *Creator) GetOptimizer() _g.Optimizer { return _bfd._edgd }

// SetSellerAddress sets the seller address of the invoice.
func (_cgdg *Invoice) SetSellerAddress(address *InvoiceAddress) { _cgdg._aedd = address }

type pageTransformations struct {
	_dgd  *_ba.Matrix
	_cead bool
	_aecf bool
}

// Invoice represents a configurable invoice template.
type Invoice struct {
	_cbde   string
	_abfb   *Image
	_dggaa  *InvoiceAddress
	_aedd   *InvoiceAddress
	_cecf   string
	_aafe   [2]*InvoiceCell
	_dgdd   [2]*InvoiceCell
	_gdge   [2]*InvoiceCell
	_bgfc   [][2]*InvoiceCell
	_cfc    []*InvoiceCell
	_cfg    [][]*InvoiceCell
	_afga   [2]*InvoiceCell
	_cdgb   [2]*InvoiceCell
	_gegc   [][2]*InvoiceCell
	_bgcb   [2]string
	_fcega  [2]string
	_facf   [][2]string
	_fafbeg TextStyle
	_ggae   TextStyle
	_bfgc   TextStyle
	_ceadb  TextStyle
	_deef   TextStyle
	_efec   TextStyle
	_dedb   TextStyle
	_dbeg   InvoiceCellProps
	_fbgc   InvoiceCellProps
	_cbcee  InvoiceCellProps
	_ddcd   InvoiceCellProps
	_fadc   Positioning
}

// AddInfo is used to append a piece of invoice information in the template
// information table.
func (_dccg *Invoice) AddInfo(description, value string) (*InvoiceCell, *InvoiceCell) {
	_fggf := [2]*InvoiceCell{_dccg.newCell(description, _dccg._dbeg), _dccg.newCell(value, _dccg._dbeg)}
	_dccg._bgfc = append(_dccg._bgfc, _fggf)
	return _fggf[0], _fggf[1]
}

// SetLineSeparatorStyle sets the style for the separator part of all new
// lines of the table of contents.
func (_faffa *TOC) SetLineSeparatorStyle(style TextStyle) { _faffa._afca = style }

// SetBorderOpacity sets the border opacity.
func (_dgdge *PolyBezierCurve) SetBorderOpacity(opacity float64) { _dgdge._dgce = opacity }

// NewRectangle creates a new Rectangle with default parameters
// with left corner at (x,y) and width, height as specified.
func (_cde *Creator) NewRectangle(x, y, width, height float64) *Rectangle {
	return _bgdd(x, y, width, height)
}

// Wrap wraps the text of the chunk into lines based on its style and the
// specified width.
func (_degfb *TextChunk) Wrap(width float64) ([]string, error) {
	if int(width) <= 0 {
		return []string{_degfb.Text}, nil
	}
	var _gcafd []string
	var _cbae []rune
	var _aeaf float64
	var _gfadb []float64
	_egff := _degfb.Style
	for _, _dbccg := range _degfb.Text {
		if _dbccg == '\u000A' {
			_gcafd = append(_gcafd, _db.TrimRightFunc(string(_cbae), _a.IsSpace)+string(_dbccg))
			_cbae = nil
			_aeaf = 0
			_gfadb = nil
			continue
		}
		_cabd := _dbccg == ' '
		_efgf, _ddaca := _egff.Font.GetRuneMetrics(_dbccg)
		if !_ddaca {
			_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _dbccg, _dbccg, _egff.Font.BaseFont(), _egff.Font.Subtype())
			_da.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _egff.Font)
			_da.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _egff.Font.Encoder())
			return nil, _f.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
		_gdfd := _egff.FontSize * _efgf.Wx
		_agbcg := _gdfd
		if !_cabd {
			_agbcg = _gdfd + _egff.CharSpacing*1000.0
		}
		if _aeaf+_gdfd > width*1000.0 {
			_abbd := -1
			if !_cabd {
				for _bfdf := len(_cbae) - 1; _bfdf >= 0; _bfdf-- {
					if _cbae[_bfdf] == ' ' {
						_abbd = _bfdf
						break
					}
				}
			}
			_cbgf := string(_cbae)
			if _abbd > 0 {
				_cbgf = string(_cbae[0 : _abbd+1])
				_cbae = append(_cbae[_abbd+1:], _dbccg)
				_gfadb = append(_gfadb[_abbd+1:], _agbcg)
				_aeaf = 0
				for _, _bcbec := range _gfadb {
					_aeaf += _bcbec
				}
			} else {
				if _cabd {
					_cbae = []rune{}
					_gfadb = []float64{}
					_aeaf = 0
				} else {
					_cbae = []rune{_dbccg}
					_gfadb = []float64{_agbcg}
					_aeaf = _agbcg
				}
			}
			_gcafd = append(_gcafd, _db.TrimRightFunc(_cbgf, _a.IsSpace))
		} else {
			_cbae = append(_cbae, _dbccg)
			_aeaf += _agbcg
			_gfadb = append(_gfadb, _agbcg)
		}
	}
	if len(_cbae) > 0 {
		_gcafd = append(_gcafd, string(_cbae))
	}
	return _gcafd, nil
}

// SetFont sets the Paragraph's font.
func (_gabc *Paragraph) SetFont(font *_g.PdfFont) { _gabc._cgbf = font }

// Reset removes all the text chunks the paragraph contains.
func (_ddfb *StyledParagraph) Reset() { _ddfb._dgff = []*TextChunk{} }
func (_edec *pageTransformations) transformPage(_fgeb *_g.PdfPage) error {
	if _edgdc := _edec.applyFlip(_fgeb); _edgdc != nil {
		return _edgdc
	}
	return nil
}

// ScaleToWidth scales the Block to a specified width, maintaining the same aspect ratio.
func (_bf *Block) ScaleToWidth(w float64) { _ddf := w / _bf._deb; _bf.Scale(_ddf, _ddf) }

// GeneratePageBlocks draws the chart onto a block.
func (_fdb *Chart) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gbb := ctx
	_dccf := _fdb._cbfa.IsRelative()
	var _ead []*Block
	if _dccf {
		_dfbg := 1.0
		_gde := _fdb._gedd.Top
		if float64(_fdb._gfgd.Height()) > ctx.Height-_fdb._gedd.Top {
			_ead = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _ffb error
			if _, ctx, _ffb = _cacf().GeneratePageBlocks(ctx); _ffb != nil {
				return nil, ctx, _ffb
			}
			_gde = 0
		}
		ctx.X += _fdb._gedd.Left + _dfbg
		ctx.Y += _gde
		ctx.Width -= _fdb._gedd.Left + _fdb._gedd.Right + 2*_dfbg
		ctx.Height -= _gde
		_fdb._gfgd.SetWidth(int(ctx.Width))
	} else {
		ctx.X = _fdb._ebfbg
		ctx.Y = _fdb._afe
	}
	_ggb := _fc.NewContentCreator()
	_ggb.Translate(0, ctx.PageHeight)
	_ggb.Scale(1, -1)
	_ggb.Translate(ctx.X, ctx.Y)
	_ggc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_fdb._gfgd.Render(_gg.NewRenderer(_ggb, _ggc._fga), nil)
	if _fdaf := _ggc.addContentsByString(_ggb.String()); _fdaf != nil {
		return nil, ctx, _fdaf
	}
	if _dccf {
		_abd := _fdb.Height() + _fdb._gedd.Bottom
		ctx.Y += _abd
		ctx.Height -= _abd
	} else {
		ctx = _gbb
	}
	_ead = append(_ead, _ggc)
	return _ead, ctx, nil
}

// SetHorizontalAlignment sets the cell's horizontal alignment of content.
// Can be one of:
// - CellHorizontalAlignmentLeft
// - CellHorizontalAlignmentCenter
// - CellHorizontalAlignmentRight
func (_cgbff *TableCell) SetHorizontalAlignment(halign CellHorizontalAlignment) {
	_cgbff._dcee = halign
}
func _dffaa(_deba []_bb.CubicBezierCurve) *PolyBezierCurve {
	return &PolyBezierCurve{_agfc: &_bb.PolyBezierCurve{Curves: _deba, BorderColor: _g.NewPdfColorDeviceRGB(0, 0, 0), BorderWidth: 1.0}, _afeb: 1.0, _dgce: 1.0}
}

// Number returns the invoice number description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_caea *Invoice) Number() (*InvoiceCell, *InvoiceCell) { return _caea._aafe[0], _caea._aafe[1] }

// SetMargins sets the Paragraph's margins.
func (_adaa *Paragraph) SetMargins(left, right, top, bottom float64) {
	_adaa._bbdee.Left = left
	_adaa._bbdee.Right = right
	_adaa._bbdee.Top = top
	_adaa._bbdee.Bottom = bottom
}

// FooterFunctionArgs holds the input arguments to a footer drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FooterFunctionArgs struct {
	PageNum    int
	TotalPages int
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

// SetPos sets the absolute position. Changes object positioning to absolute.
func (_gcbfb *Image) SetPos(x, y float64) {
	_gcbfb._cbce = PositionAbsolute
	_gcbfb._edcgf = x
	_gcbfb._baef = y
}

// Write output of creator to io.Writer interface.
func (_bdce *Creator) Write(ws _de.Writer) error {
	if _cbg := _bdce.Finalize(); _cbg != nil {
		return _cbg
	}
	_bafb := _g.NewPdfWriter()
	_bafb.SetOptimizer(_bdce._edgd)
	if _bdce._dfce != nil {
		_gea := _bafb.SetForms(_bdce._dfce)
		if _gea != nil {
			_da.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _gea)
			return _gea
		}
	}
	if _bdce._cadb != nil {
		_bafb.AddOutlineTree(_bdce._cadb)
	} else if _bdce._edge != nil && _bdce.AddOutlines {
		_bafb.AddOutlineTree(&_bdce._edge.ToPdfOutline().PdfOutlineTreeNode)
	}
	if _bdce._bdgf != nil {
		if _fdfg := _bafb.SetPageLabels(_bdce._bdgf); _fdfg != nil {
			_da.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020C\u006f\u0075\u006c\u0064 no\u0074 s\u0065\u0074\u0020\u0070\u0061\u0067\u0065 l\u0061\u0062\u0065\u006c\u0073\u003a\u0020%\u0076", _fdfg)
			return _fdfg
		}
	}
	if _bdce._dbc != nil {
		for _, _gdg := range _bdce._dbc {
			_ggdg := _gdg.SubsetRegistered()
			if _ggdg != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006ft\u0020s\u0075\u0062\u0073\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _ggdg)
				return _ggdg
			}
		}
	}
	if _bdce._aaee != nil {
		_ebce := _bdce._aaee(&_bafb)
		if _ebce != nil {
			_da.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _ebce)
			return _ebce
		}
	}
	for _, _adea := range _bdce._bdb {
		_cab := _bafb.AddPage(_adea)
		if _cab != nil {
			_da.Log.Error("\u0046\u0061\u0069\u006ced\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0050\u0061\u0067\u0065\u003a\u0020%\u0076", _cab)
			return _cab
		}
	}
	_efc := _bafb.Write(ws)
	if _efc != nil {
		return _efc
	}
	return nil
}

// GetCoords returns coordinates of the Rectangle's upper left corner (x,y).
func (_dffe *Rectangle) GetCoords() (float64, float64) { return _dffe._eebca, _dffe._eaaa }

// SetBuyerAddress sets the buyer address of the invoice.
func (_geba *Invoice) SetBuyerAddress(address *InvoiceAddress) { _geba._dggaa = address }

// SetPos sets absolute positioning with specified coordinates.
func (_bbcgf *StyledParagraph) SetPos(x, y float64) {
	_bbcgf._dfgbc = PositionAbsolute
	_bbcgf._aage = x
	_bbcgf._ebcb = y
}

// AddTextItem appends a new item with the specified text to the list.
// The method creates a styled paragraph with the specified text and returns
// it so that the item style can be customized.
// The method also returns the marker used for the newly added item.
// The marker object can be used to change the text and style of the marker
// for the current item.
func (_aeeb *List) AddTextItem(text string) (*StyledParagraph, *TextChunk, error) {
	_dbeed := _fbgb(_aeeb._egbg)
	_dbeed.Append(text)
	_eccg, _befc := _aeeb.Add(_dbeed)
	return _dbeed, _eccg, _befc
}

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_baeda *Paragraph) SetTextAlignment(align TextAlignment) { _baeda._dgadd = align }

// SetSubtotal sets the subtotal of the invoice.
func (_eecc *Invoice) SetSubtotal(value string) { _eecc._afga[1].Value = value }

// LevelOffset returns the amount of space an indentation level occupies.
func (_dgdfa *TOCLine) LevelOffset() float64 { return _dgdfa._dedcg }
func (_gdd *pageTransformations) applyFlip(_aefc *_g.PdfPage) error {
	_dfg, _daba := _gdd._cead, _gdd._aecf
	if !_dfg && !_daba {
		return nil
	}
	if _aefc == nil {
		return _f.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_cgd, _dbfb := _aefc.GetMediaBox()
	if _dbfb != nil {
		return _dbfb
	}
	_bbdc, _aba := _cgd.Width(), _cgd.Height()
	_aab, _dbfb := _aefc.GetRotate()
	if _dbfb != nil {
		_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _dbfb.Error())
	}
	if _gfae := _aab%360 != 0 && _aab%90 == 0; _gfae {
		if _fafd := (360 + _aab%360) % 360; _fafd == 90 || _fafd == 270 {
			_dfg, _daba = _daba, _dfg
		}
	}
	_bcfg, _bebe := 1.0, 0.0
	if _dfg {
		_bcfg, _bebe = -1.0, -_bbdc
	}
	_fdcf, _cgac := 1.0, 0.0
	if _daba {
		_fdcf, _cgac = -1.0, -_aba
	}
	_aaa := _fc.NewContentCreator().Scale(_bcfg, _fdcf).Translate(_bebe, _cgac)
	_cag, _dbfb := _ec.MakeStream(_aaa.Bytes(), _ec.NewFlateEncoder())
	if _dbfb != nil {
		return _dbfb
	}
	_ffea := _ec.MakeArray(_cag)
	_ffea.Append(_aefc.GetContentStreamObjs()...)
	_aefc.Contents = _ffea
	return nil
}

// GeneratePageBlocks generates the page blocks for the Division component.
// Multiple blocks are generated if the contents wrap over multiple pages.
func (_agf *Division) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var _fbdd []*Block
	_cfbg := _agf._aeae.IsRelative()
	_ada := _agf._abdb.Top
	if _cfbg && !_agf._babf && !_agf._abce {
		_gaag := _agf.ctxHeight(ctx.Width - _agf._abdb.Left - _agf._abdb.Right)
		if _gaag > ctx.Height-_agf._abdb.Top && _gaag <= ctx.PageHeight-ctx.Margins.Top-ctx.Margins.Bottom {
			var _acega error
			if _fbdd, ctx, _acega = _cacf().GeneratePageBlocks(ctx); _acega != nil {
				return nil, ctx, _acega
			}
			_ada = 0
		}
	}
	_eaba := ctx
	if _cfbg {
		ctx.X += _agf._abdb.Left
		ctx.Y += _ada
		ctx.Width -= _agf._abdb.Left + _agf._abdb.Right
		ctx.Height -= _ada + _agf._abdb.Bottom
	}
	ctx.Inline = _agf._abce
	_abgdd := ctx
	_gbga := ctx
	var _cffb float64
	for _, _afd := range _agf._eggb {
		if ctx.Inline {
			if (ctx.X-_abgdd.X)+_afd.Width() <= ctx.Width {
				ctx.Y = _gbga.Y
				ctx.Height = _gbga.Height
			} else {
				ctx.X = _abgdd.X
				ctx.Width = _abgdd.Width
				_gbga.Y += _cffb
				_gbga.Height -= _cffb
				_cffb = 0
			}
		}
		_beccc, _agef, _dbee := _afd.GeneratePageBlocks(ctx)
		if _dbee != nil {
			_da.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006eg\u0020p\u0061\u0067\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073\u003a\u0020\u0025\u0076", _dbee)
			return nil, ctx, _dbee
		}
		if len(_beccc) < 1 {
			continue
		}
		if len(_fbdd) > 0 {
			_fbdd[len(_fbdd)-1].mergeBlocks(_beccc[0])
			_fbdd = append(_fbdd, _beccc[1:]...)
		} else {
			_fbdd = append(_fbdd, _beccc[0:]...)
		}
		if ctx.Inline {
			if ctx.Page != _agef.Page {
				_abgdd.Y = ctx.Margins.Top
				_abgdd.Height = ctx.PageHeight - ctx.Margins.Top
				_gbga.Y = _abgdd.Y
				_gbga.Height = _abgdd.Height
				_cffb = _agef.Height - _abgdd.Height
			} else {
				if _eead := ctx.Height - _agef.Height; _eead > _cffb {
					_cffb = _eead
				}
			}
		} else {
			_agef.X = ctx.X
		}
		ctx = _agef
	}
	ctx.Inline = _eaba.Inline
	if _cfbg {
		ctx.X = _eaba.X
	}
	if _agf._aeae.IsAbsolute() {
		return _fbdd, _eaba, nil
	}
	return _fbdd, ctx, nil
}

// Indent returns the left offset of the list when nested into another list.
func (_gabbf *List) Indent() float64 { return _gabbf._fefc }

// Add adds a new Drawable to the chapter.
func (_efg *Chapter) Add(d Drawable) error {
	if Drawable(_efg) == d {
		_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0061\u0064\u0064\u0020\u0069\u0074\u0073\u0065\u006c\u0066")
		return _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	switch _dbbc := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Block, *Table, *PageBreak, *Chapter:
		_efg._aaga = append(_efg._aaga, d)
	case containerDrawable:
		_adf, _ceec := _dbbc.ContainerComponent(_efg)
		if _ceec != nil {
			return _ceec
		}
		_efg._aaga = append(_efg._aaga, _adf)
	default:
		_da.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u003a\u0020\u0025\u0054", d)
		return _f.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return nil
}

const (
	CellBorderSideLeft CellBorderSide = iota
	CellBorderSideRight
	CellBorderSideTop
	CellBorderSideBottom
	CellBorderSideAll
)

func (_ecf *Block) addContentsByString(_eac string) error {
	_beg := _fc.NewContentStreamParser(_eac)
	_ca, _gfc := _beg.Parse()
	if _gfc != nil {
		return _gfc
	}
	_ecf._cb.WrapIfNeeded()
	_ca.WrapIfNeeded()
	*_ecf._cb = append(*_ecf._cb, *_ca...)
	return nil
}

// GetCoords returns the (x1, y1), (x2, y2) points defining the Line.
func (_cfba *Line) GetCoords() (float64, float64, float64, float64) {
	return _cfba._baba, _cfba._fbb, _cfba._ggadb, _cfba._bgbc
}

// AddSection adds a new content section at the end of the invoice.
func (_bffc *Invoice) AddSection(title, content string) {
	_bffc._facf = append(_bffc._facf, [2]string{title, content})
}

// SetFillColor sets the fill color.
func (_cbbb *Rectangle) SetFillColor(col Color) { _cbbb._edga = col }
func (_cdce *Image) rotatedSize() (float64, float64) {
	_dcgf := _cdce._ebb
	_afg := _cdce._gcab
	_fbdc := _cdce._dfdb
	if _fbdc == 0 {
		return _dcgf, _afg
	}
	_dbda := _bb.Path{Points: []_bb.Point{_bb.NewPoint(0, 0).Rotate(_fbdc), _bb.NewPoint(_dcgf, 0).Rotate(_fbdc), _bb.NewPoint(0, _afg).Rotate(_fbdc), _bb.NewPoint(_dcgf, _afg).Rotate(_fbdc)}}.GetBoundingBox()
	return _dbda.Width, _dbda.Height
}

// EnablePageWrap controls whether the table is wrapped across pages.
// If disabled, the table is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the table is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
func (_dcege *Table) EnablePageWrap(enable bool) { _dcege._ebcca = enable }

// NewBlock creates a new Block with specified width and height.
func NewBlock(width float64, height float64) *Block {
	_dc := &Block{}
	_dc._cb = &_fc.ContentStreamOperations{}
	_dc._fga = _g.NewPdfPageResources()
	_dc._deb = width
	_dc._df = height
	return _dc
}

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
//
func (_gdc *Creator) SetPageSize(size PageSize) {
	_gdc._ebad = size
	_gdc._cgb = size[0]
	_gdc._cggd = size[1]
	_abdd := 0.1 * _gdc._cgb
	_gdc._acda.Left = _abdd
	_gdc._acda.Right = _abdd
	_gdc._acda.Top = _abdd
	_gdc._acda.Bottom = _abdd
}

// TOC returns the table of contents component of the creator.
func (_eae *Creator) TOC() *TOC { return _eae._fgc }

// SetPageMargins sets the page margins: left, right, top, bottom.
// The default page margins are 10% of document width.
func (_dgba *Creator) SetPageMargins(left, right, top, bottom float64) {
	_dgba._acda.Left = left
	_dgba._acda.Right = right
	_dgba._acda.Top = top
	_dgba._acda.Bottom = bottom
}
func (_dgfdf *StyledParagraph) getTextLineWidth(_baacb []*TextChunk) float64 {
	var _cafg float64
	_fcbe := len(_baacb)
	for _fgef, _bfcg := range _baacb {
		_fgaeb := &_bfcg.Style
		_afabd := len(_bfcg.Text)
		for _degb, _agege := range _bfcg.Text {
			if _agege == '\u000A' {
				continue
			}
			_accfc, _ggee := _fgaeb.Font.GetRuneMetrics(_agege)
			if !_ggee {
				_da.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _agege)
				return -1
			}
			_cafg += _fgaeb.FontSize * _accfc.Wx * _fgaeb.horizontalScale()
			if _agege != ' ' && (_fgef != _fcbe-1 || _degb != _afabd-1) {
				_cafg += _fgaeb.CharSpacing * 1000.0
			}
		}
	}
	return _cafg
}

// PolyBezierCurve represents a composite curve that is the result of joining
// multiple cubic Bezier curves.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type PolyBezierCurve struct {
	_agfc *_bb.PolyBezierCurve
	_afeb float64
	_dgce float64
}

// SetLineHeight sets the line height (1.0 default).
func (_bdgag *Paragraph) SetLineHeight(lineheight float64) { _bdgag._dfgg = lineheight }
func _eeab(_edcb, _cceee, _afedb float64) (_degee, _gdfe, _bccdg, _ecae float64) {
	if _afedb == 0 {
		return 0, 0, _edcb, _cceee
	}
	_baad := _bb.Path{Points: []_bb.Point{_bb.NewPoint(0, 0).Rotate(_afedb), _bb.NewPoint(_edcb, 0).Rotate(_afedb), _bb.NewPoint(0, _cceee).Rotate(_afedb), _bb.NewPoint(_edcb, _cceee).Rotate(_afedb)}}.GetBoundingBox()
	return _baad.X, _baad.Y, _baad.Width, _baad.Height
}

// SetWidth sets the the Paragraph width. This is essentially the wrapping width, i.e. the width the
// text can extend to prior to wrapping over to next line.
func (_abaeb *Paragraph) SetWidth(width float64) { _abaeb._dgee = width; _abaeb.wrapText() }

// MultiRowCell makes a new cell with the specified row span and inserts it
// into the table at the current position.
func (_agbeb *Table) MultiRowCell(rowspan int) *TableCell { return _agbeb.MultiCell(rowspan, 1) }

// SetFillOpacity sets the fill opacity.
func (_cbef *CurvePolygon) SetFillOpacity(opacity float64) { _cbef._accd = opacity }

// SetTOC sets the table of content component of the creator.
// This method should be used when building a custom table of contents.
func (_geff *Creator) SetTOC(toc *TOC) {
	if toc == nil {
		return
	}
	_geff._fgc = toc
}

// SetIndent sets the cell's left indent.
func (_cbcg *TableCell) SetIndent(indent float64) { _cbcg._fffcb = indent }

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_fadd *Creator) SetOptimizer(optimizer _g.Optimizer) { _fadd._edgd = optimizer }

// SetOutlineTree adds the specified outline tree to the PDF file generated
// by the creator. Adding an external outline tree disables the automatic
// generation of outlines done by the creator for the relevant components.
func (_acbg *Creator) SetOutlineTree(outlineTree *_g.PdfOutlineTreeNode) { _acbg._cadb = outlineTree }

// SetWidth sets the the Paragraph width. This is essentially the wrapping width,
// i.e. the width the text can extend to prior to wrapping over to next line.
func (_ggdfc *StyledParagraph) SetWidth(width float64) { _ggdfc._bggd = width; _ggdfc.wrapText() }
func _ggef(_fdc, _aedc *_g.PdfPageResources) error {
	_cbdbc, _ := _fdc.GetColorspaces()
	if _cbdbc != nil && len(_cbdbc.Colorspaces) > 0 {
		for _dfa, _cad := range _cbdbc.Colorspaces {
			_dae := *_ec.MakeName(_dfa)
			if _aedc.HasColorspaceByName(_dae) {
				continue
			}
			_fcg := _aedc.SetColorspaceByName(_dae, _cad)
			if _fcg != nil {
				return _fcg
			}
		}
	}
	return nil
}

var PPMM = float64(72 * 1.0 / 25.4)

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_bbb *List) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var _adegf float64
	var _gdcb []*StyledParagraph
	for _, _gdba := range _bbb._cbba {
		_feff := _fbgb(_bbb._egbg)
		_feff.SetEnableWrap(false)
		_feff.SetTextAlignment(TextAlignmentRight)
		_feff.Append(_gdba._bebc.Text).Style = _gdba._bebc.Style
		_gcfeb := _feff.getTextWidth() / 1000.0 / ctx.Width
		if _adegf < _gcfeb {
			_adegf = _gcfeb
		}
		_gdcb = append(_gdcb, _feff)
	}
	_gccf := _dbfc(2)
	_gccf.SetColumnWidths(_adegf, 1-_adegf)
	_gccf.SetMargins(_bbb._fefc, 0, 0, 0)
	for _cgdb, _bdfdc := range _bbb._cbba {
		_gcbfe := _gccf.NewCell()
		_gcbfe.SetIndent(0)
		_gcbfe.SetContent(_gdcb[_cgdb])
		_gcbfe = _gccf.NewCell()
		_gcbfe.SetIndent(0)
		_gcbfe.SetContent(_bdfdc._dege)
	}
	return _gccf.GeneratePageBlocks(ctx)
}

// SetFillColor sets the fill color.
func (_ceda *CurvePolygon) SetFillColor(color Color) { _ceda._bafd.FillColor = _afag(color) }

// Width is not used. Not used as a Division element is designed to fill into available width depending on
// context.  Returns 0.
func (_fcbf *Division) Width() float64 { return 0 }

// AddAnnotation adds an annotation to the current block.
// The annotation will be added to the page the block will be rendered on.
func (_bd *Block) AddAnnotation(annotation *_g.PdfAnnotation) {
	for _, _acb := range _bd._ggg {
		if _acb == annotation {
			return
		}
	}
	_bd._ggg = append(_bd._ggg, annotation)
}
func _dfgb(_cgbc TextStyle) *List {
	return &List{_ddcde: TextChunk{Text: "\u2022\u0020", Style: _cgbc}, _fefc: 0, _dcffe: true, _cffa: PositionRelative, _egbg: _cgbc}
}

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

// NewSubchapter creates a new child chapter with the specified title.
func (_bcf *Chapter) NewSubchapter(title string) *Chapter {
	_bcg := _gadb(_bcf._gacfc._cgbf)
	_bcg.FontSize = 14
	_bcf._fcd++
	_cbf := _gbg(_bcf, _bcf._gecdd, _bcf._gfag, title, _bcf._fcd, _bcg)
	_bcf.Add(_cbf)
	return _cbf
}
func _ffbgf(_ebfef, _cbfd, _ebbd TextChunk, _eefe uint, _dcga TextStyle) *TOCLine {
	_aegf := _fbgb(_dcga)
	_aegf.SetEnableWrap(true)
	_aegf.SetTextAlignment(TextAlignmentLeft)
	_aegf.SetMargins(0, 0, 2, 2)
	_dfeda := &TOCLine{_daedd: _aegf, Number: _ebfef, Title: _cbfd, Page: _ebbd, Separator: TextChunk{Text: "\u002e", Style: _dcga}, _egde: 0, _cefb: _eefe, _dedcg: 10, _cada: PositionRelative}
	_aegf._cdbab.Left = _dfeda._egde + float64(_dfeda._cefb-1)*_dfeda._dedcg
	_aegf._bcaf = _dfeda.prepareParagraph
	return _dfeda
}

// NoteStyle returns the style properties used to render the content of the
// invoice note sections.
func (_gdbcdc *Invoice) NoteStyle() TextStyle { return _gdbcdc._efec }

// Length calculates and returns the line length.
func (_aafg *Line) Length() float64 {
	return _dg.Sqrt(_dg.Pow(_aafg._ggadb-_aafg._baba, 2.0) + _dg.Pow(_aafg._bgbc-_aafg._fbb, 2.0))
}

// ColorRGBFrom8bit creates a Color from 8-bit (0-255) r,g,b values.
// Example:
//   red := ColorRGBFrom8Bit(255, 0, 0)
func ColorRGBFrom8bit(r, g, b byte) Color {
	return rgbColor{_aacc: float64(r) / 255.0, _beeb: float64(g) / 255.0, _fgea: float64(b) / 255.0}
}
func _bbf(_fca _b.ChartRenderable) *Chart {
	return &Chart{_gfgd: _fca, _cbfa: PositionRelative, _gedd: Margins{Top: 10, Bottom: 10}}
}

// SetLineStyle sets the style for all the line components: number, title,
// separator, page. The style is applied only for new lines added to the
// TOC component.
func (_bfdbd *TOC) SetLineStyle(style TextStyle) {
	_bfdbd.SetLineNumberStyle(style)
	_bfdbd.SetLineTitleStyle(style)
	_bfdbd.SetLineSeparatorStyle(style)
	_bfdbd.SetLinePageStyle(style)
}

// Height returns the height of the chart.
func (_bef *Chart) Height() float64 { return float64(_bef._gfgd.Height()) }

// Level returns the indentation level of the TOC line.
func (_gcdcg *TOCLine) Level() uint { return _gcdcg._cefb }

// GeneratePageBlocks draws the polyline on a new block representing the page.
// Implements the Drawable interface.
func (_dddbe *Polyline) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ecff := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_aaeeb, _gbef := _ecff.setOpacity(_dddbe._beccb, _dddbe._beccb)
	if _gbef != nil {
		return nil, ctx, _gbef
	}
	_fgabb := _dddbe._efcf.Points
	for _ebac := range _fgabb {
		_ebfd := &_fgabb[_ebac]
		_ebfd.Y = ctx.PageHeight - _ebfd.Y
	}
	_effg, _, _gbef := _dddbe._efcf.Draw(_aaeeb)
	if _gbef != nil {
		return nil, ctx, _gbef
	}
	if _gbef = _ecff.addContentsByString(string(_effg)); _gbef != nil {
		return nil, ctx, _gbef
	}
	return []*Block{_ecff}, ctx, nil
}

// Date returns the invoice date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_gfe *Invoice) Date() (*InvoiceCell, *InvoiceCell) { return _gfe._dgdd[0], _gfe._dgdd[1] }
func (_cdaec *TextStyle) horizontalScale() float64       { return _cdaec.HorizontalScaling / 100 }
func (_aadb *StyledParagraph) getTextHeight() float64 {
	var _acbfc float64
	for _, _ecaba := range _aadb._dgff {
		_ccge := _ecaba.Style.FontSize * _aadb._bggb
		if _ccge > _acbfc {
			_acbfc = _ccge
		}
	}
	return _acbfc
}

// New creates a new instance of the PDF Creator.
func New() *Creator {
	_dfed := &Creator{}
	_dfed._bdb = []*_g.PdfPage{}
	_dfed._fcac = map[*_g.PdfPage]*Block{}
	_dfed._gff = map[*_g.PdfPage]*pageTransformations{}
	_dfed.SetPageSize(PageSizeLetter)
	_cafa := 0.1 * _dfed._cgb
	_dfed._acda.Left = _cafa
	_dfed._acda.Right = _cafa
	_dfed._acda.Top = _cafa
	_dfed._acda.Bottom = _cafa
	var _fafbe error
	_dfed._bge, _fafbe = _g.NewStandard14Font(_g.HelveticaName)
	if _fafbe != nil {
		_dfed._bge = _g.DefaultFont()
	}
	_dfed._dcaf, _fafbe = _g.NewStandard14Font(_g.HelveticaBoldName)
	if _fafbe != nil {
		_dfed._bge = _g.DefaultFont()
	}
	_dfed._fgc = _dfed.NewTOC("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073")
	_dfed.AddOutlines = true
	_dfed._edge = _g.NewOutline()
	return _dfed
}

// RotatedSize returns the width and height of the rotated block.
func (_fge *Block) RotatedSize() (float64, float64) {
	_, _, _cbd, _aga := _eeab(_fge._deb, _fge._df, _fge._fb)
	return _cbd, _aga
}

// Division is a container component which can wrap across multiple pages (unlike Block).
// It can contain multiple Drawable components (currently supporting Paragraph and Image).
//
// The component stacking behavior is vertical, where the Drawables are drawn on top of each other.
// Also supports horizontal stacking by activating the inline mode.
type Division struct {
	_eggb []VectorDrawable
	_aeae Positioning
	_abdb Margins
	_abce bool
	_babf bool
}

// SetEncoder sets the encoding/compression mechanism for the image.
func (_ggge *Image) SetEncoder(encoder _ec.StreamEncoder) { _ggge._fdgf = encoder }

// GeneratePageBlocks generates the page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages. Implements the Drawable interface.
func (_eafe *StyledParagraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_adgf := ctx
	var _cbea []*Block
	_cfcf := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _eafe._dfgbc.IsRelative() {
		ctx.X += _eafe._cdbab.Left
		ctx.Y += _eafe._cdbab.Top
		ctx.Width -= _eafe._cdbab.Left + _eafe._cdbab.Right
		ctx.Height -= _eafe._cdbab.Top
		_eafe.SetWidth(ctx.Width)
	} else {
		if int(_eafe._bggd) <= 0 {
			_eafe.SetWidth(_eafe.getTextWidth() / 1000.0)
		}
		ctx.X = _eafe._aage
		ctx.Y = _eafe._ebcb
	}
	if _eafe._bcaf != nil {
		_eafe._bcaf(_eafe, ctx)
	}
	if _dafb := _eafe.wrapText(); _dafb != nil {
		return nil, ctx, _dafb
	}
	_edfa := _eafe._aageb
	for {
		_agab, _cdgf, _bbe := _ccec(_cfcf, _eafe, _edfa, ctx)
		if _bbe != nil {
			_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bbe)
			return nil, ctx, _bbe
		}
		ctx = _agab
		_cbea = append(_cbea, _cfcf)
		if _edfa = _cdgf; len(_cdgf) == 0 {
			break
		}
		_cfcf = NewBlock(ctx.PageWidth, ctx.PageHeight)
		ctx.Page++
		_agab = ctx
		_agab.Y = ctx.Margins.Top
		_agab.X = ctx.Margins.Left + _eafe._cdbab.Left
		_agab.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
		_agab.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _eafe._cdbab.Left - _eafe._cdbab.Right
		ctx = _agab
	}
	if _eafe._dfgbc.IsRelative() {
		ctx.Y += _eafe._cdbab.Bottom
		ctx.Height -= _eafe._cdbab.Bottom
		if !ctx.Inline {
			ctx.X = _adgf.X
			ctx.Width = _adgf.Width
		}
		return _cbea, ctx, nil
	}
	return _cbea, _adgf, nil
}

// GetRowHeight returns the height of the specified row.
func (_dcce *Table) GetRowHeight(row int) (float64, error) {
	if row < 1 || row > len(_dcce._abeg) {
		return 0, _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	return _dcce._abeg[row-1], nil
}

// SetVerticalAlignment set the cell's vertical alignment of content.
// Can be one of:
// - CellHorizontalAlignmentTop
// - CellHorizontalAlignmentMiddle
// - CellHorizontalAlignmentBottom
func (_eecfa *TableCell) SetVerticalAlignment(valign CellVerticalAlignment) { _eecfa._dfcfe = valign }
func (_bfc *Creator) newPage() *_g.PdfPage {
	_cfeb := _g.NewPdfPage()
	_abgd := _bfc._ebad[0]
	_gfff := _bfc._ebad[1]
	_dfag := _g.PdfRectangle{Llx: 0, Lly: 0, Urx: _abgd, Ury: _gfff}
	_cfeb.MediaBox = &_dfag
	_bfc._cgb = _abgd
	_bfc._cggd = _gfff
	_bfc.initContext()
	return _cfeb
}
func (_bdaf *Invoice) generateTotalBlocks(_cfgd DrawContext) ([]*Block, DrawContext, error) {
	_gdag := _dbfc(4)
	_gdag.SetMargins(0, 0, 10, 10)
	_edeee := [][2]*InvoiceCell{_bdaf._afga}
	_edeee = append(_edeee, _bdaf._gegc...)
	_edeee = append(_edeee, _bdaf._cdgb)
	for _, _fdgd := range _edeee {
		_bfea, _aagfa := _fdgd[0], _fdgd[1]
		if _aagfa.Value == "" {
			continue
		}
		_gdag.SkipCells(2)
		_baaa := _gdag.NewCell()
		_baaa.SetBackgroundColor(_bfea.BackgroundColor)
		_baaa.SetHorizontalAlignment(_aagfa.Alignment)
		_bdaf.setCellBorder(_baaa, _bfea)
		_eedf := _fbgb(_bfea.TextStyle)
		_eedf.SetMargins(0, 0, 2, 1)
		_eedf.Append(_bfea.Value)
		_baaa.SetContent(_eedf)
		_baaa = _gdag.NewCell()
		_baaa.SetBackgroundColor(_aagfa.BackgroundColor)
		_baaa.SetHorizontalAlignment(_aagfa.Alignment)
		_bdaf.setCellBorder(_baaa, _bfea)
		_eedf = _fbgb(_aagfa.TextStyle)
		_eedf.SetMargins(0, 0, 2, 1)
		_eedf.Append(_aagfa.Value)
		_baaa.SetContent(_eedf)
	}
	return _gdag.GeneratePageBlocks(_cfgd)
}

// TextStyle is a collection of properties that can be assigned to a text chunk.
type TextStyle struct {

	// Color represents the color of the text.
	Color Color

	// OutlineColor represents the color of the text outline.
	OutlineColor Color

	// Font represents the font the text will use.
	Font *_g.PdfFont

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

// GeneratePageBlocks draws the block contents on a template Page block.
// Implements the Drawable interface.
func (_bg *Block) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_eba := _fc.NewContentCreator()
	_fd, _eag := _bg.Width(), _bg.Height()
	if _bg._ff.IsRelative() {
		_eba.Translate(ctx.X, ctx.PageHeight-ctx.Y-_eag)
	} else {
		_eba.Translate(_bg._ce, ctx.PageHeight-_bg._fe-_eag)
	}
	_bc := _eag
	if _bg._fb != 0 {
		_eba.Translate(_fd/2, _eag/2)
		_eba.RotateDeg(_bg._fb)
		_eba.Translate(-_fd/2, -_eag/2)
		_, _bc = _bg.RotatedSize()
	}
	if _bg._ff.IsRelative() {
		ctx.Y += _bc
	}
	_bde := _bg.duplicate()
	_gge := append(*_eba.Operations(), *_bde._cb...)
	_gge.WrapIfNeeded()
	_bde._cb = &_gge
	return []*Block{_bde}, ctx, nil
}

// SetMargins sets the margins of the paragraph.
func (_fdea *List) SetMargins(left, right, top, bottom float64) {
	_fdea._cgab.Left = left
	_fdea._cgab.Right = right
	_fdea._cgab.Top = top
	_fdea._cgab.Bottom = bottom
}

// MoveX moves the drawing context to absolute position x.
func (_fdaa *Creator) MoveX(x float64) { _fdaa._defb.X = x }

// ColorCMYKFrom8bit creates a Color from c,m,y,k values (0-100).
// Example:
//   red := ColorCMYKFrom8Bit(0, 100, 100, 0)
func ColorCMYKFrom8bit(c, m, y, k byte) Color {
	return cmykColor{_debb: _dg.Min(float64(c), 100) / 100.0, _caad: _dg.Min(float64(m), 100) / 100.0, _gfbg: _dg.Min(float64(y), 100) / 100.0, _abf: _dg.Min(float64(k), 100) / 100.0}
}

// NewTOC creates a new table of contents.
func (_afade *Creator) NewTOC(title string) *TOC {
	_ecb := _afade.NewTextStyle()
	_ecb.Font = _afade._dcaf
	return _bedaf(title, _afade.NewTextStyle(), _ecb)
}

// GeneratePageBlocks draws the composite curve polygon on a new block
// representing the page. Implements the Drawable interface.
func (_ddb *CurvePolygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_addd := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_ceff, _eeef := _addd.setOpacity(_ddb._accd, _ddb._dfaa)
	if _eeef != nil {
		return nil, ctx, _eeef
	}
	_cafac := _ddb._bafd
	_cafac.FillEnabled = _cafac.FillColor != nil
	_cafac.BorderEnabled = _cafac.BorderColor != nil && _cafac.BorderWidth > 0
	var (
		_cadc = ctx.PageHeight
		_eaab = _cafac.Rings
		_ffce = make([][]_bb.CubicBezierCurve, 0, len(_cafac.Rings))
	)
	for _, _dea := range _eaab {
		_adcf := make([]_bb.CubicBezierCurve, 0, len(_dea))
		for _, _dbd := range _dea {
			_aaebg := _dbd
			_aaebg.P0.Y = _cadc - _aaebg.P0.Y
			_aaebg.P1.Y = _cadc - _aaebg.P1.Y
			_aaebg.P2.Y = _cadc - _aaebg.P2.Y
			_aaebg.P3.Y = _cadc - _aaebg.P3.Y
			_adcf = append(_adcf, _aaebg)
		}
		_ffce = append(_ffce, _adcf)
	}
	_cafac.Rings = _ffce
	defer func() { _cafac.Rings = _eaab }()
	_dgea, _, _eeef := _cafac.Draw(_ceff)
	if _eeef != nil {
		return nil, ctx, _eeef
	}
	if _eeef = _addd.addContentsByString(string(_dgea)); _eeef != nil {
		return nil, ctx, _eeef
	}
	return []*Block{_addd}, ctx, nil
}
func (_fdgg *pageTransformations) transformBlock(_bffg *Block) {
	if _fdgg._dgd != nil {
		_bffg.transform(*_fdgg._dgd)
	}
}

// Scale block by specified factors in the x and y directions.
func (_adc *Block) Scale(sx, sy float64) {
	_ggd := _fc.NewContentCreator().Scale(sx, sy).Operations()
	*_adc._cb = append(*_ggd, *_adc._cb...)
	_adc._cb.WrapIfNeeded()
	_adc._deb *= sx
	_adc._df *= sy
}

// SetBorder sets the cell's border style.
func (_cdfad *TableCell) SetBorder(side CellBorderSide, style CellBorderStyle, width float64) {
	if style == CellBorderStyleSingle && side == CellBorderSideAll {
		_cdfad._dbcf = CellBorderStyleSingle
		_cdfad._geaef = width
		_cdfad._ebfg = CellBorderStyleSingle
		_cdfad._eecfg = width
		_cdfad._eacc = CellBorderStyleSingle
		_cdfad._edgb = width
		_cdfad._aagb = CellBorderStyleSingle
		_cdfad._eaeg = width
	} else if style == CellBorderStyleDouble && side == CellBorderSideAll {
		_cdfad._dbcf = CellBorderStyleDouble
		_cdfad._geaef = width
		_cdfad._ebfg = CellBorderStyleDouble
		_cdfad._eecfg = width
		_cdfad._eacc = CellBorderStyleDouble
		_cdfad._edgb = width
		_cdfad._aagb = CellBorderStyleDouble
		_cdfad._eaeg = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideLeft {
		_cdfad._dbcf = style
		_cdfad._geaef = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideBottom {
		_cdfad._ebfg = style
		_cdfad._eecfg = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideRight {
		_cdfad._eacc = style
		_cdfad._edgb = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideTop {
		_cdfad._aagb = style
		_cdfad._eaeg = width
	}
}

// SetPos sets absolute positioning with specified coordinates.
func (_ffa *Paragraph) SetPos(x, y float64) {
	_ffa._fcgd = PositionAbsolute
	_ffa._fccc = x
	_ffa._ddda = y
}

// SetHeaderRows turns the selected table rows into headers that are repeated
// for every page the table spans. startRow and endRow are inclusive.
func (_cgbg *Table) SetHeaderRows(startRow, endRow int) error {
	if startRow <= 0 {
		return _f.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020r\u006f\u0077\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if endRow <= 0 {
		return _f.New("\u0068\u0065a\u0064\u0065\u0072\u0020e\u006e\u0064 \u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if startRow > endRow {
		return _f.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0072\u006f\u0077\u0020\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065 \u0065\u006e\u0064\u0020\u0072o\u0077")
	}
	_cgbg._gbcfa = true
	_cgbg._dced = startRow
	_cgbg._dgfbg = endRow
	return nil
}

// Height returns the Block's height.
func (_ag *Block) Height() float64 { return _ag._df }

// AppendColumn appends a column to the line items table.
func (_bdfg *Invoice) AppendColumn(description string) *InvoiceCell {
	_bdga := _bdfg.NewColumn(description)
	_bdfg._cfc = append(_bdfg._cfc, _bdga)
	return _bdga
}

var (
	PageSizeA3     = PageSize{297 * PPMM, 420 * PPMM}
	PageSizeA4     = PageSize{210 * PPMM, 297 * PPMM}
	PageSizeA5     = PageSize{148 * PPMM, 210 * PPMM}
	PageSizeLetter = PageSize{8.5 * PPI, 11 * PPI}
	PageSizeLegal  = PageSize{8.5 * PPI, 14 * PPI}
)

// ScaleToHeight scale Image to a specified height h, maintaining the aspect ratio.
func (_ddaa *Image) ScaleToHeight(h float64) {
	_efdgg := _ddaa._ebb / _ddaa._gcab
	_ddaa._gcab = h
	_ddaa._ebb = h * _efdgg
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

// Fit fits the chunk into the specified bounding box, cropping off the
// remainder in a new chunk, if it exceeds the specified dimensions.
// NOTE: The method assumes a line height of 1.0. In order to account for other
// line height values, the passed in height must be divided by the line height:
// height = height / lineHeight
func (_daed *TextChunk) Fit(width, height float64) (*TextChunk, error) {
	_bfdfg, _eacd := _daed.Wrap(width)
	if _eacd != nil {
		return nil, _eacd
	}
	_ecccg := int(height / _daed.Style.FontSize)
	if _ecccg >= len(_bfdfg) {
		return nil, nil
	}
	_cadd := "\u000a"
	_daed.Text = _db.Replace(_db.Join(_bfdfg[:_ecccg], "\u0020"), _cadd+"\u0020", _cadd, -1)
	_cffc := _db.Replace(_db.Join(_bfdfg[_ecccg:], "\u0020"), _cadd+"\u0020", _cadd, -1)
	return NewTextChunk(_cffc, _daed.Style), nil
}

// SetNoteStyle sets the style properties used to render the content of the
// invoice note sections.
func (_fbdg *Invoice) SetNoteStyle(style TextStyle) { _fbdg._efec = style }

// CurCol returns the currently active cell's column number.
func (_bdcee *Table) CurCol() int { _fcdb := (_bdcee._daab-1)%(_bdcee._fabb) + 1; return _fcdb }

// SetPdfWriterAccessFunc sets a PdfWriter access function/hook.
// Exposes the PdfWriter just prior to writing the PDF.  Can be used to encrypt the output PDF, etc.
//
// Example of encrypting with a user/owner password "password"
// Prior to calling c.WriteFile():
//
// c.SetPdfWriterAccessFunc(func(w *model.PdfWriter) error {
//	userPass := []byte("password")
//	ownerPass := []byte("password")
//	err := w.Encrypt(userPass, ownerPass, nil)
//	return err
// })
//
func (_fbgd *Creator) SetPdfWriterAccessFunc(pdfWriterAccessFunc func(_dfde *_g.PdfWriter) error) {
	_fbgd._aaee = pdfWriterAccessFunc
}

// SetLineHeight sets the line height (1.0 default).
func (_bdee *StyledParagraph) SetLineHeight(lineheight float64) { _bdee._bggb = lineheight }
func _dbdac(_ggf, _abgg, _ege, _dgbc float64) *Line {
	_bgcae := &Line{}
	_bgcae._baba = _ggf
	_bgcae._fbb = _abgg
	_bgcae._ggadb = _ege
	_bgcae._bgbc = _dgbc
	_bgcae._cbgb = ColorBlack
	_bgcae._bbgc = 1.0
	return _bgcae
}

// GeneratePageBlocks draws the rectangle on a new block representing the page.
func (_decf *Ellipse) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_daff := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_ageg := _bb.Circle{X: _decf._gffd - _decf._afcf/2, Y: ctx.PageHeight - _decf._cgae - _decf._edfg/2, Width: _decf._afcf, Height: _decf._edfg, Opacity: 1.0, BorderWidth: _decf._gca}
	if _decf._fcaa != nil {
		_ageg.FillEnabled = true
		_ageg.FillColor = _afag(_decf._fcaa)
	}
	if _decf._gbad != nil {
		_ageg.BorderEnabled = true
		_ageg.BorderColor = _afag(_decf._gbad)
		_ageg.BorderWidth = _decf._gca
	}
	_dcdb, _, _fba := _ageg.Draw("")
	if _fba != nil {
		return nil, ctx, _fba
	}
	_fba = _daff.addContentsByString(string(_dcdb))
	if _fba != nil {
		return nil, ctx, _fba
	}
	return []*Block{_daff}, ctx, nil
}
func _eded() *Division { return &Division{_babf: true} }

// AddSubtable copies the cells of the subtable in the table, starting with the
// specified position. The table row and column indices are 1-based, which
// makes the position of the first cell of the first row of the table 1,1.
// The table is automatically extended if the subtable exceeds its columns.
// This can happen when the subtable has more columns than the table or when
// one or more columns of the subtable starting from the specified position
// exceed the last column of the table.
func (_daffg *Table) AddSubtable(row, col int, subtable *Table) {
	for _, _fbea := range subtable._dcab {
		_faef := &TableCell{}
		*_faef = *_fbea
		_faef._gfdee = _daffg
		_faef._gbdde += col - 1
		if _ededb := _daffg._fabb - (_faef._gbdde - 1); _ededb < _faef._bdbg {
			_daffg._fabb += _faef._bdbg - _ededb
			_daffg.resetColumnWidths()
			_da.Log.Debug("\u0054a\u0062l\u0065\u003a\u0020\u0073\u0075\u0062\u0074\u0061\u0062\u006c\u0065 \u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0061\u0062\u006c\u0065\u002e\u0020\u0045\u0078\u0070\u0061\u006e\u0064\u0069\u006e\u0067\u0020\u0074\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0025\u0064\u0020\u0063\u006fl\u0075\u006d\u006e\u0073\u002e", _daffg._fabb)
		}
		_faef._cbfc += row - 1
		_fdabg := subtable._abeg[_fbea._cbfc-1]
		if _faef._cbfc > _daffg._gbcf {
			for _faef._cbfc > _daffg._gbcf {
				_daffg._gbcf++
				_daffg._abeg = append(_daffg._abeg, _daffg._badaf)
			}
			_daffg._abeg[_faef._cbfc-1] = _fdabg
		} else {
			_daffg._abeg[_faef._cbfc-1] = _dg.Max(_daffg._abeg[_faef._cbfc-1], _fdabg)
		}
		_daffg._dcab = append(_daffg._dcab, _faef)
	}
	_e.Slice(_daffg._dcab, func(_fbcg, _beddb int) bool {
		_gaaec := _daffg._dcab[_fbcg]._cbfc
		_fecag := _daffg._dcab[_beddb]._cbfc
		if _gaaec < _fecag {
			return true
		}
		if _gaaec > _fecag {
			return false
		}
		return _daffg._dcab[_fbcg]._gbdde < _daffg._dcab[_beddb]._gbdde
	})
}
func _bgdd(_cfcb, _dgdc, _gcbd, _cgbe float64) *Rectangle {
	return &Rectangle{_eebca: _cfcb, _eaaa: _dgdc, _ecgg: _gcbd, _eeee: _cgbe, _fdegf: ColorBlack, _abfa: 1.0, _fcae: 1.0, _ccgf: 1.0}
}

// SetWidthRight sets border width for right.
func (_abg *border) SetWidthRight(bw float64) { _abg._cedc = bw }

// GetMargins returns the Block's margins: left, right, top, bottom.
func (_dag *Block) GetMargins() (float64, float64, float64, float64) {
	return _dag._ae.Left, _dag._ae.Right, _dag._ae.Top, _dag._ae.Bottom
}

// InsertColumn inserts a column in the line items table at the specified index.
func (_dcbg *Invoice) InsertColumn(index uint, description string) *InvoiceCell {
	_ggedg := uint(len(_dcbg._cfc))
	if index > _ggedg {
		index = _ggedg
	}
	_ecdd := _dcbg.NewColumn(description)
	_dcbg._cfc = append(_dcbg._cfc[:index], append([]*InvoiceCell{_ecdd}, _dcbg._cfc[index:]...)...)
	return _ecdd
}
func _fbgb(_edfc TextStyle) *StyledParagraph {
	return &StyledParagraph{_dgff: []*TextChunk{}, _fgegg: _edfc, _aacda: _dggd(_edfc.Font), _bggb: 1.0, _abgbf: TextAlignmentLeft, _cfga: true, _deag: true, _ffgac: 0, _acfc: 1, _fdca: 1, _dfgbc: PositionRelative}
}

// SetBorderOpacity sets the border opacity.
func (_bgeb *Rectangle) SetBorderOpacity(opacity float64) { _bgeb._ccgf = opacity }
func _ccec(_cadf *Block, _cbced *StyledParagraph, _afagg [][]*TextChunk, _dcdbd DrawContext) (DrawContext, [][]*TextChunk, error) {
	_adeec := 1
	_bafg := _ec.PdfObjectName(_ad.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _adeec))
	for _cadf._fga.HasFontByName(_bafg) {
		_adeec++
		_bafg = _ec.PdfObjectName(_ad.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _adeec))
	}
	_cddbe := _cadf._fga.SetFontByName(_bafg, _cbced._fgegg.Font.ToPdfObject())
	if _cddbe != nil {
		return _dcdbd, nil, _cddbe
	}
	_adeec++
	_ccab := _bafg
	_gagg := _cbced._fgegg.FontSize
	_cegd := _cbced._dfgbc.IsRelative()
	var _cdga [][]_ec.PdfObjectName
	var _beda [][]*TextChunk
	var _aaeg float64
	for _bbbg, _fafe := range _afagg {
		var _egebb []_ec.PdfObjectName
		var _eecb float64
		if len(_fafe) > 0 {
			_eecb = _fafe[0].Style.FontSize
		}
		for _, _gdda := range _fafe {
			_gcff := _gdda.Style
			if _gdda.Text != "" && _gcff.FontSize > _eecb {
				_eecb = _gcff.FontSize
			}
			_bafg = _ec.PdfObjectName(_ad.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _adeec))
			_efae := _cadf._fga.SetFontByName(_bafg, _gcff.Font.ToPdfObject())
			if _efae != nil {
				return _dcdbd, nil, _efae
			}
			_egebb = append(_egebb, _bafg)
			_adeec++
		}
		_eecb *= _cbced._bggb
		if _cegd && _aaeg+_eecb > _dcdbd.Height {
			_beda = _afagg[_bbbg:]
			_afagg = _afagg[:_bbbg]
			break
		}
		_aaeg += _eecb
		_cdga = append(_cdga, _egebb)
	}
	_fbdcf, _dece, _ebacf := _cbced.getLineMetrics(0)
	_cbac, _eeebg := _fbdcf*_cbced._bggb, _dece*_cbced._bggb
	_fageg := _fc.NewContentCreator()
	_fageg.Add_q()
	_dbdd := _eeebg
	if _cbced._fgce == TextVerticalAlignmentCenter {
		_dbdd = _dece + (_fbdcf+_ebacf-_dece)/2 + (_eeebg-_dece)/2
	}
	_cbfe := _dcdbd.PageHeight - _dcdbd.Y - _dbdd
	_fageg.Translate(_dcdbd.X, _cbfe)
	_dgfa := _cbfe
	if _cbced._ffgac != 0 {
		_fageg.RotateDeg(_cbced._ffgac)
	}
	if _cbced._cfgc == TextOverflowHidden {
		_fageg.Add_re(0, -_aaeg+_cbac+1, _cbced._bggd, _aaeg).Add_W().Add_n()
	}
	_fageg.Add_BT()
	var _fged []*_bb.BasicLine
	for _ggbb, _dbab := range _afagg {
		_dbbd := _dcdbd.X
		var _ddba float64
		if len(_dbab) > 0 {
			_ddba = _dbab[0].Style.FontSize
		}
		for _, _afaf := range _dbab {
			_bfce := &_afaf.Style
			if _afaf.Text != "" && _bfce.FontSize > _ddba {
				_ddba = _bfce.FontSize
			}
		}
		if _ggbb != 0 {
			_fageg.Add_TD(0, -_ddba*_cbced._bggb)
			_dgfa -= _ddba * _cbced._bggb
		}
		_bcebe := _ggbb == len(_afagg)-1
		var (
			_gacgd float64
			_edce  float64
			_fccgf float64
			_cdfa  uint
		)
		var _ccbdf []float64
		for _, _aaega := range _dbab {
			_cbdbb := &_aaega.Style
			if _cbdbb.FontSize > _edce {
				_edce = _cbdbb.FontSize
			}
			_aegcb, _accbe := _cbdbb.Font.GetRuneMetrics(' ')
			if !_accbe {
				return _dcdbd, nil, _f.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
			}
			var _edgdcb uint
			var _ceeb float64
			_agac := len(_aaega.Text)
			for _ffdgd, _decdb := range _aaega.Text {
				if _decdb == ' ' {
					_edgdcb++
					continue
				}
				if _decdb == '\u000A' {
					continue
				}
				_ddeac, _cdffd := _cbdbb.Font.GetRuneMetrics(_decdb)
				if !_cdffd {
					_da.Log.Debug("\u0055\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006ee\u0020%\u0076\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u000a", _decdb)
					return _dcdbd, nil, _f.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
				}
				_ceeb += _cbdbb.FontSize * _ddeac.Wx * _cbdbb.horizontalScale()
				if _ffdgd != _agac-1 {
					_ceeb += _cbdbb.CharSpacing * 1000.0
				}
			}
			_ccbdf = append(_ccbdf, _ceeb)
			_gacgd += _ceeb
			_fccgf += float64(_edgdcb) * _aegcb.Wx * _cbdbb.FontSize * _cbdbb.horizontalScale()
			_cdfa += _edgdcb
		}
		_edce *= _cbced._bggb
		var _ggbd []_ec.PdfObject
		_gbbg := _cbced._bggd * 1000.0
		if _cbced._abgbf == TextAlignmentJustify {
			if _cdfa > 0 && !_bcebe {
				_fccgf = (_gbbg - _gacgd) / float64(_cdfa) / _gagg
			}
		} else if _cbced._abgbf == TextAlignmentCenter {
			_abcb := (_gbbg - _gacgd - _fccgf) / 2
			_fcde := _abcb / _gagg
			_ggbd = append(_ggbd, _ec.MakeFloat(-_fcde))
			_dbbd += _abcb / 1000.0
		} else if _cbced._abgbf == TextAlignmentRight {
			_ffgc := (_gbbg - _gacgd - _fccgf)
			_dbgd := _ffgc / _gagg
			_ggbd = append(_ggbd, _ec.MakeFloat(-_dbgd))
			_dbbd += _ffgc / 1000.0
		}
		if len(_ggbd) > 0 {
			_fageg.Add_Tf(_ccab, _gagg).Add_TL(_gagg * _cbced._bggb).Add_TJ(_ggbd...)
		}
		for _efcfa, _ceag := range _dbab {
			_aadf := &_ceag.Style
			_cdgbf := _ccab
			_decfc := _gagg
			_bccc := _aadf.OutlineColor != nil
			_bgde := _aadf.HorizontalScaling != DefaultHorizontalScaling
			_bbge := _aadf.OutlineSize != 1
			if _bbge {
				_fageg.Add_w(_aadf.OutlineSize)
			}
			_fegga := _aadf.RenderingMode != TextRenderingModeFill
			if _fegga {
				_fageg.Add_Tr(int64(_aadf.RenderingMode))
			}
			_ecbc := _aadf.CharSpacing != 0
			if _ecbc {
				_fageg.Add_Tc(_aadf.CharSpacing)
			}
			_cccgc := _aadf.TextRise != 0
			if _cccgc {
				_fageg.Add_Ts(_aadf.TextRise)
			}
			if _cbced._abgbf != TextAlignmentJustify || _bcebe {
				_gdagf, _eabb := _aadf.Font.GetRuneMetrics(' ')
				if !_eabb {
					return _dcdbd, nil, _f.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
				}
				_cdgbf = _cdga[_ggbb][_efcfa]
				_decfc = _aadf.FontSize
				_fccgf = _gdagf.Wx * _aadf.horizontalScale()
			}
			_dfgc := _aadf.Font.Encoder()
			var _gcbc []byte
			for _, _gdbd := range _ceag.Text {
				if _gdbd == '\u000A' {
					continue
				}
				if _gdbd == ' ' {
					if len(_gcbc) > 0 {
						if _bccc {
							_fageg.SetStrokingColor(_afag(_aadf.OutlineColor))
						}
						if _bgde {
							_fageg.Add_Tz(_aadf.HorizontalScaling)
						}
						_fageg.SetNonStrokingColor(_afag(_aadf.Color)).Add_Tf(_cdga[_ggbb][_efcfa], _aadf.FontSize).Add_TJ([]_ec.PdfObject{_ec.MakeStringFromBytes(_gcbc)}...)
						_gcbc = nil
					}
					if _bgde {
						_fageg.Add_Tz(DefaultHorizontalScaling)
					}
					_fageg.Add_Tf(_cdgbf, _decfc).Add_TJ([]_ec.PdfObject{_ec.MakeFloat(-_fccgf)}...)
					_ccbdf[_efcfa] += _fccgf * _decfc
				} else {
					if _, _bbgg := _dfgc.RuneToCharcode(_gdbd); !_bbgg {
						_cddbe = UnsupportedRuneError{Message: _ad.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _gdbd, _gdbd), Rune: _gdbd}
						_dcdbd._bgdc = append(_dcdbd._bgdc, _cddbe)
						_da.Log.Debug(_cddbe.Error())
						if _dcdbd._bfaf <= 0 {
							continue
						}
						_gdbd = _dcdbd._bfaf
					}
					_gcbc = append(_gcbc, _dfgc.Encode(string(_gdbd))...)
				}
			}
			if len(_gcbc) > 0 {
				if _bccc {
					_fageg.SetStrokingColor(_afag(_aadf.OutlineColor))
				}
				if _bgde {
					_fageg.Add_Tz(_aadf.HorizontalScaling)
				}
				_fageg.SetNonStrokingColor(_afag(_aadf.Color)).Add_Tf(_cdga[_ggbb][_efcfa], _aadf.FontSize).Add_TJ([]_ec.PdfObject{_ec.MakeStringFromBytes(_gcbc)}...)
			}
			_bbgab := _ccbdf[_efcfa] / 1000.0
			if _aadf.Underline {
				_agbcc := _aadf.UnderlineStyle.Color
				if _agbcc == nil {
					_agbcc = _ceag.Style.Color
				}
				_bgaf, _cbgbe, _dbeee := _agbcc.ToRGB()
				_gbadf := _dbbd - _dcdbd.X
				_cgeg := _dgfa - _cbfe + _aadf.TextRise - _aadf.UnderlineStyle.Offset
				_fged = append(_fged, &_bb.BasicLine{X1: _gbadf, Y1: _cgeg, X2: _gbadf + _bbgab, Y2: _cgeg, LineWidth: _ceag.Style.UnderlineStyle.Thickness, LineColor: _g.NewPdfColorDeviceRGB(_bgaf, _cbgbe, _dbeee)})
			}
			if _ceag._adbcg != nil {
				var _afegb *_ec.PdfObjectArray
				if !_ceag._cegdd {
					switch _bdeee := _ceag._adbcg.GetContext().(type) {
					case *_g.PdfAnnotationLink:
						_afegb = _ec.MakeArray()
						_bdeee.Rect = _afegb
						_gcafe, _aeca := _bdeee.Dest.(*_ec.PdfObjectArray)
						if _aeca && _gcafe.Len() == 5 {
							_fdabe, _gbeg := _gcafe.Get(1).(*_ec.PdfObjectName)
							if _gbeg && _fdabe.String() == "\u0058\u0059\u005a" {
								_ffdae, _aaeafd := _ec.GetNumberAsFloat(_gcafe.Get(3))
								if _aaeafd == nil {
									_gcafe.Set(3, _ec.MakeFloat(_dcdbd.PageHeight-_ffdae))
								}
							}
						}
					}
					_ceag._cegdd = true
				}
				if _afegb != nil {
					_aaba := _bb.NewPoint(_dbbd-_dcdbd.X, _dgfa+_aadf.TextRise-_cbfe).Rotate(_cbced._ffgac)
					_aaba.X += _dcdbd.X
					_aaba.Y += _cbfe
					_bagg, _fab, _aggga, _eddf := _eeab(_bbgab, _edce, _cbced._ffgac)
					_aaba.X += _bagg
					_aaba.Y += _fab
					_afegb.Clear()
					_afegb.Append(_ec.MakeFloat(_aaba.X))
					_afegb.Append(_ec.MakeFloat(_aaba.Y))
					_afegb.Append(_ec.MakeFloat(_aaba.X + _aggga))
					_afegb.Append(_ec.MakeFloat(_aaba.Y + _eddf))
				}
				_cadf.AddAnnotation(_ceag._adbcg)
			}
			_dbbd += _bbgab
			if _bbge {
				_fageg.Add_w(1.0)
			}
			if _bccc {
				_fageg.Add_RG(0.0, 0.0, 0.0)
			}
			if _fegga {
				_fageg.Add_Tr(int64(TextRenderingModeFill))
			}
			if _ecbc {
				_fageg.Add_Tc(0)
			}
			if _cccgc {
				_fageg.Add_Ts(0)
			}
			if _bgde {
				_fageg.Add_Tz(DefaultHorizontalScaling)
			}
		}
	}
	_fageg.Add_ET()
	for _, _dddba := range _fged {
		_fageg.SetStrokingColor(_dddba.LineColor).Add_w(_dddba.LineWidth).Add_m(_dddba.X1, _dddba.Y1).Add_l(_dddba.X2, _dddba.Y2).Add_s()
	}
	_fageg.Add_Q()
	_adff := _fageg.Operations()
	_adff.WrapIfNeeded()
	_cadf.addContents(_adff)
	if _cegd {
		_fbfe := _aaeg
		_dcdbd.Y += _fbfe
		_dcdbd.Height -= _fbfe
		if _dcdbd.Inline {
			_dcdbd.X += _cbced.Width() + _cbced._cdbab.Right
		}
	}
	return _dcdbd, _beda, nil
}
func (_aefg *Invoice) newColumn(_ecada string, _gcgcc CellHorizontalAlignment) *InvoiceCell {
	_bbfg := &InvoiceCell{_aefg._fbgc, _ecada}
	_bbfg.Alignment = _gcgcc
	return _bbfg
}

// SetPageLabels adds the specified page labels to the PDF file generated
// by the creator. See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
// NOTE: for existing PDF files, the page label ranges object can be obtained
// using the model.PDFReader's GetPageLabels method.
func (_cgcc *Creator) SetPageLabels(pageLabels _ec.PdfObject) { _cgcc._bdgf = pageLabels }

// SetDueDate sets the due date of the invoice.
func (_gcba *Invoice) SetDueDate(dueDate string) (*InvoiceCell, *InvoiceCell) {
	_gcba._gdge[1].Value = dueDate
	return _gcba._gdge[0], _gcba._gdge[1]
}

// SetLineNumberStyle sets the style for the numbers part of all new lines
// of the table of contents.
func (_affg *TOC) SetLineNumberStyle(style TextStyle) { _affg._ceaf = style }

// SetPos sets the Table's positioning to absolute mode and specifies the upper-left corner
// coordinates as (x,y).
// Note that this is only sensible to use when the table does not wrap over multiple pages.
// TODO: Should be able to set width too (not just based on context/relative positioning mode).
func (_bfbaa *Table) SetPos(x, y float64) {
	_bfbaa._dbedc = PositionAbsolute
	_bfbaa._dfedb = x
	_bfbaa._dccgd = y
}
func (_egae *Invoice) generateInformationBlocks(_gabbg DrawContext) ([]*Block, DrawContext, error) {
	_bece := _fbgb(_egae._fafbeg)
	_bece.SetMargins(0, 0, 0, 20)
	_fece := _egae.drawAddress(_egae._aedd)
	_fece = append(_fece, _bece)
	_fece = append(_fece, _egae.drawAddress(_egae._dggaa)...)
	_bbde := _eded()
	for _, _feac := range _fece {
		_bbde.Add(_feac)
	}
	_dgad := _egae.drawInformation()
	_baddb := _dbfc(2)
	_baddb.SetMargins(0, 0, 25, 0)
	_afbcc := _baddb.NewCell()
	_afbcc.SetIndent(0)
	_afbcc.SetContent(_bbde)
	_afbcc = _baddb.NewCell()
	_afbcc.SetContent(_dgad)
	return _baddb.GeneratePageBlocks(_gabbg)
}

// Polygon represents a polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Polygon struct {
	_fdfa *_bb.Polygon
	_aebd float64
	_daae float64
}

// SetColor sets the color of the Paragraph text.
//
// Example:
// 1.   p := NewParagraph("Red paragraph")
//      // Set to red color with a hex code:
//      p.SetColor(creator.ColorRGBFromHex("#ff0000"))
//
// 2. Make Paragraph green with 8-bit rgb values (0-255 each component)
//      p.SetColor(creator.ColorRGBFrom8bit(0, 255, 0)
//
// 3. Make Paragraph blue with arithmetic (0-1) rgb components.
//      p.SetColor(creator.ColorRGBFromArithmetic(0, 0, 1.0)
//
func (_dabf *Paragraph) SetColor(col Color) { _dabf._gade = col }

// NewParagraph creates a new text paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_fbfa *Creator) NewParagraph(text string) *Paragraph { return _ebfbc(text, _fbfa.NewTextStyle()) }
func _fefe(_ffaf *_g.PdfAnnotationLink) *_g.PdfAnnotationLink {
	if _ffaf == nil {
		return nil
	}
	_ebbg := _g.NewPdfAnnotationLink()
	_ebbg.BS = _ffaf.BS
	_ebbg.A = _ffaf.A
	if _bbeg, _ebace := _ffaf.GetAction(); _ebace == nil && _bbeg != nil {
		_ebbg.SetAction(_bbeg)
	}
	if _gdbda, _bbfef := _ffaf.Dest.(*_ec.PdfObjectArray); _bbfef {
		_ebbg.Dest = _ec.MakeArray(_gdbda.Elements()...)
	}
	return _ebbg
}
func (_cdca *List) tableHeight(_aagcb float64) float64 {
	var _gcce float64
	for _, _aaae := range _cdca._cbba {
		switch _ccd := _aaae._dege.(type) {
		case *Paragraph:
			_abae := _ccd
			if _abae._acegd {
				_abae.SetWidth(_aagcb)
			}
			_gcce += _abae.Height() + _abae._bbdee.Bottom + _abae._bbdee.Bottom
			_gcce += 0.5 * _abae._ebbae * _abae._dfgg
		case *StyledParagraph:
			_bceb := _ccd
			if _bceb._cfga {
				_bceb.SetWidth(_aagcb)
			}
			_gcce += _bceb.Height() + _bceb._cdbab.Top + _bceb._cdbab.Bottom
			_gcce += 0.5 * _bceb.getTextHeight()
		default:
			_gcce += _aaae._dege.Height()
		}
	}
	return _gcce
}

// SetColorLeft sets border color for left.
func (_bgb *border) SetColorLeft(col Color) { _bgb._gcbe = col }

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_gcee *Invoice) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gcdc := ctx
	_dbff := []func(_bacf DrawContext) ([]*Block, DrawContext, error){_gcee.generateHeaderBlocks, _gcee.generateInformationBlocks, _gcee.generateLineBlocks, _gcee.generateTotalBlocks, _gcee.generateNoteBlocks}
	var _cefd []*Block
	for _, _abge := range _dbff {
		_dadg, _adcga, _adcb := _abge(ctx)
		if _adcb != nil {
			return _cefd, ctx, _adcb
		}
		if len(_cefd) == 0 {
			_cefd = _dadg
		} else if len(_dadg) > 0 {
			_cefd[len(_cefd)-1].mergeBlocks(_dadg[0])
			_cefd = append(_cefd, _dadg[1:]...)
		}
		ctx = _adcga
	}
	if _gcee._fadc.IsRelative() {
		ctx.X = _gcdc.X
	}
	if _gcee._fadc.IsAbsolute() {
		return _cefd, _gcdc, nil
	}
	return _cefd, ctx, nil
}

// TableCell defines a table cell which can contain a Drawable as content.
type TableCell struct {
	_dcbgd        Color
	_egba         _bb.LineStyle
	_dbcf         CellBorderStyle
	_faaea        Color
	_geaef        float64
	_ebfg         CellBorderStyle
	_becd         Color
	_eecfg        float64
	_eacc         CellBorderStyle
	_cedab        Color
	_edgb         float64
	_aagb         CellBorderStyle
	_bdcf         Color
	_eaeg         float64
	_cbfc, _gbdde int
	_cdcd         int
	_bdbg         int
	_adeed        VectorDrawable
	_dcee         CellHorizontalAlignment
	_dfcfe        CellVerticalAlignment
	_fffcb        float64
	_gfdee        *Table
}

// SetLineLevelOffset sets the amount of space an indentation level occupies
// for all new lines of the table of contents.
func (_dbge *TOC) SetLineLevelOffset(levelOffset float64) { _dbge._eecca = levelOffset }
func _cdbb(_dedc [][]_bb.CubicBezierCurve) *CurvePolygon {
	return &CurvePolygon{_bafd: &_bb.CurvePolygon{Rings: _dedc}, _accd: 1.0, _dfaa: 1.0}
}

// Add adds a VectorDrawable to the Division container.
// Currently supported VectorDrawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Chart
func (_cgde *Division) Add(d VectorDrawable) error {
	switch _accbc := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart:
	case containerDrawable:
		_dbfe, _dgdgd := _accbc.ContainerComponent(_cgde)
		if _dgdgd != nil {
			return _dgdgd
		}
		_becc, _feae := _dbfe.(VectorDrawable)
		if !_feae {
			return _ad.Errorf("\u0072\u0065\u0073\u0075\u006ct\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u002d\u0020\u0025\u0054\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0056\u0065c\u0074\u006f\u0072\u0044\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020i\u006e\u0074\u0065\u0072\u0066\u0061c\u0065", _dbfe)
		}
		d = _becc
	default:
		return _f.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0044\u0069\u0076\u0069\u0073i\u006f\u006e")
	}
	_cgde._eggb = append(_cgde._eggb, d)
	return nil
}

type cmykColor struct{ _debb, _caad, _gfbg, _abf float64 }

// Columns returns all the columns in the invoice line items table.
func (_bgag *Invoice) Columns() []*InvoiceCell { return _bgag._cfc }

// SetAngle sets the rotation angle in degrees.
func (_eb *Block) SetAngle(angleDeg float64) { _eb._fb = angleDeg }

// ScaleToHeight scales the Block to a specified height, maintaining the same aspect ratio.
func (_bfe *Block) ScaleToHeight(h float64) { _cd := h / _bfe._df; _bfe.Scale(_cd, _cd) }
func (_ccfd *FilledCurve) draw(_fgfc string) ([]byte, *_g.PdfRectangle, error) {
	_efca := _bb.NewCubicBezierPath()
	for _, _geaa := range _ccfd._dgfb {
		_efca = _efca.AppendCurve(_geaa)
	}
	creator := _fc.NewContentCreator()
	creator.Add_q()
	if _ccfd.FillEnabled && _ccfd._debe != nil {
		creator.SetNonStrokingColor(_afag(_ccfd._debe))
	}
	if _ccfd.BorderEnabled {
		if _ccfd._dddbg != nil {
			creator.SetStrokingColor(_afag(_ccfd._dddbg))
		}
		creator.Add_w(_ccfd.BorderWidth)
	}
	if len(_fgfc) > 1 {
		creator.Add_gs(_ec.PdfObjectName(_fgfc))
	}
	_bb.DrawBezierPathWithCreator(_efca, creator)
	creator.Add_h()
	if _ccfd.FillEnabled && _ccfd.BorderEnabled {
		creator.Add_B()
	} else if _ccfd.FillEnabled {
		creator.Add_f()
	} else if _ccfd.BorderEnabled {
		creator.Add_S()
	}
	creator.Add_Q()
	_dfdd := _efca.GetBoundingBox()
	if _ccfd.BorderEnabled {
		_dfdd.Height += _ccfd.BorderWidth
		_dfdd.Width += _ccfd.BorderWidth
		_dfdd.X -= _ccfd.BorderWidth / 2
		_dfdd.Y -= _ccfd.BorderWidth / 2
	}
	_egdfa := &_g.PdfRectangle{}
	_egdfa.Llx = _dfdd.X
	_egdfa.Lly = _dfdd.Y
	_egdfa.Urx = _dfdd.X + _dfdd.Width
	_egdfa.Ury = _dfdd.Y + _dfdd.Height
	return creator.Bytes(), _egdfa, nil
}

// NewPolygon creates a new polygon.
func (_dfcb *Creator) NewPolygon(points [][]_bb.Point) *Polygon { return _edfgc(points) }
func _cgccc(_aadc string) *_g.PdfAnnotation {
	_ddcbc := _g.NewPdfAnnotationLink()
	_abac := _g.NewBorderStyle()
	_abac.SetBorderWidth(0)
	_ddcbc.BS = _abac.ToPdfObject()
	_egga := _g.NewPdfActionURI()
	_egga.URI = _ec.MakeString(_aadc)
	_ddcbc.SetAction(_egga.PdfAction)
	return _ddcbc.PdfAnnotation
}

// TextVerticalAlignment controls the vertical position of the text
// in a styled paragraph.
type TextVerticalAlignment int
type border struct {
	_edd      float64
	_ccga     float64
	_debf     float64
	_cfb      float64
	_edee     Color
	_gcbe     Color
	_age      float64
	_aad      Color
	_gfg      float64
	_edf      Color
	_cedc     float64
	_acg      Color
	_bdf      float64
	LineStyle _bb.LineStyle
	_gda      CellBorderStyle
	_gga      CellBorderStyle
	_ddc      CellBorderStyle
	_bedg     CellBorderStyle
}

func (_dfef *TOCLine) prepareParagraph(_eeagg *StyledParagraph, _dbec DrawContext) {
	_aeeff := _dfef.Title.Text
	if _dfef.Number.Text != "" {
		_aeeff = "\u0020" + _aeeff
	}
	_aeeff += "\u0020"
	_efdc := _dfef.Page.Text
	if _efdc != "" {
		_efdc = "\u0020" + _efdc
	}
	_eeagg._dgff = []*TextChunk{{Text: _dfef.Number.Text, Style: _dfef.Number.Style, _adbcg: _dfef.getLineLink()}, {Text: _aeeff, Style: _dfef.Title.Style, _adbcg: _dfef.getLineLink()}, {Text: _efdc, Style: _dfef.Page.Style, _adbcg: _dfef.getLineLink()}}
	_eeagg.wrapText()
	_gfef := len(_eeagg._aageb)
	if _gfef == 0 {
		return
	}
	_fbddd := _dbec.Width*1000 - _eeagg.getTextLineWidth(_eeagg._aageb[_gfef-1])
	_aaeebg := _eeagg.getTextLineWidth([]*TextChunk{&_dfef.Separator})
	_dgfaa := int(_fbddd / _aaeebg)
	_fefa := _db.Repeat(_dfef.Separator.Text, _dgfaa)
	_gfggba := _dfef.Separator.Style
	_bcbef := _eeagg.Insert(2, _fefa)
	_bcbef.Style = _gfggba
	_bcbef._adbcg = _dfef.getLineLink()
	_fbddd = _fbddd - float64(_dgfaa)*_aaeebg
	if _fbddd > 500 {
		_gfee, _bbgfd := _gfggba.Font.GetRuneMetrics(' ')
		if _bbgfd && _fbddd > _gfee.Wx {
			_eedc := int(_fbddd / _gfee.Wx)
			if _eedc > 0 {
				_cbbf := _gfggba
				_cbbf.FontSize = 1
				_bcbef = _eeagg.Insert(2, _db.Repeat("\u0020", _eedc))
				_bcbef.Style = _cbbf
				_bcbef._adbcg = _dfef.getLineLink()
			}
		}
	}
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
	_bfaf  rune
	_bgdc  []error
}

// Height returns the height of the Paragraph. The height is calculated based on the input text and
// how it is wrapped within the container. Does not include Margins.
func (_bfdde *Paragraph) Height() float64 {
	_bfdde.wrapText()
	return float64(len(_bfdde._egge)) * _bfdde._dfgg * _bfdde._ebbae
}

// Inline returns whether the inline mode of the division is active.
func (_bfbb *Division) Inline() bool { return _bfbb._abce }

// Sections returns the custom content sections of the invoice as
// title-content pairs.
func (_bedga *Invoice) Sections() [][2]string { return _bedga._facf }

// NewCurvePolygon creates a new curve polygon.
func (_egdf *Creator) NewCurvePolygon(rings [][]_bb.CubicBezierCurve) *CurvePolygon {
	return _cdbb(rings)
}

// Block contains a portion of PDF Page contents. It has a width and a position and can
// be placed anywhere on a Page.  It can even contain a whole Page, and is used in the creator
// where each Drawable object can output one or more blocks, each representing content for separate pages
// (typically needed when Page breaks occur).
type Block struct {
	_cb      *_fc.ContentStreamOperations
	_fga     *_g.PdfPageResources
	_ff      Positioning
	_ce, _fe float64
	_deb     float64
	_df      float64
	_fb      float64
	_ae      Margins
	_ggg     []*_g.PdfAnnotation
}

// HeaderFunctionArgs holds the input arguments to a header drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type HeaderFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// Insert adds a new text chunk at the specified position in the paragraph.
func (_bdge *StyledParagraph) Insert(index uint, text string) *TextChunk {
	_ffdaa := uint(len(_bdge._dgff))
	if index > _ffdaa {
		index = _ffdaa
	}
	_acfd := NewTextChunk(text, _bdge._fgegg)
	_bdge._dgff = append(_bdge._dgff[:index], append([]*TextChunk{_acfd}, _bdge._dgff[index:]...)...)
	_bdge.wrapText()
	return _acfd
}

// SetNotes sets the notes section of the invoice.
func (_adee *Invoice) SetNotes(title, content string) { _adee._bgcb = [2]string{title, content} }

// ColorCMYKFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//   green := ColorCMYKFromArithmetic(1.0, 0.0, 1.0, 0.0)
func ColorCMYKFromArithmetic(c, m, y, k float64) Color {
	return cmykColor{_debb: _dg.Max(_dg.Min(c, 1.0), 0.0), _caad: _dg.Max(_dg.Min(m, 1.0), 0.0), _gfbg: _dg.Max(_dg.Min(y, 1.0), 0.0), _abf: _dg.Max(_dg.Min(k, 1.0), 0.0)}
}
func (_agfb *Paragraph) getTextLineWidth(_ffdca string) float64 {
	var _dedf float64
	for _, _ffga := range _ffdca {
		if _ffga == '\u000A' {
			continue
		}
		_bcgc, _cfbgf := _agfb._cgbf.GetRuneMetrics(_ffga)
		if !_cfbgf {
			_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _ffga, _ffga)
			return -1
		}
		_dedf += _agfb._ebbae * _bcgc.Wx
	}
	return _dedf
}

// GeneratePageBlocks draws the composite Bezier curve on a new block
// representing the page. Implements the Drawable interface.
func (_fcfdg *PolyBezierCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ddac := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_cdaf, _dcgd := _ddac.setOpacity(_fcfdg._afeb, _fcfdg._dgce)
	if _dcgd != nil {
		return nil, ctx, _dcgd
	}
	_cgfd := _fcfdg._agfc
	_cgfd.FillEnabled = _cgfd.FillColor != nil
	var (
		_fbgce = ctx.PageHeight
		_gdga  = _cgfd.Curves
		_bcbfg = make([]_bb.CubicBezierCurve, 0, len(_cgfd.Curves))
	)
	for _dcba := range _cgfd.Curves {
		_fgeba := _gdga[_dcba]
		_fgeba.P0.Y = _fbgce - _fgeba.P0.Y
		_fgeba.P1.Y = _fbgce - _fgeba.P1.Y
		_fgeba.P2.Y = _fbgce - _fgeba.P2.Y
		_fgeba.P3.Y = _fbgce - _fgeba.P3.Y
		_bcbfg = append(_bcbfg, _fgeba)
	}
	_cgfd.Curves = _bcbfg
	defer func() { _cgfd.Curves = _gdga }()
	_fbeg, _, _dcgd := _cgfd.Draw(_cdaf)
	if _dcgd != nil {
		return nil, ctx, _dcgd
	}
	if _dcgd = _ddac.addContentsByString(string(_fbeg)); _dcgd != nil {
		return nil, ctx, _dcgd
	}
	return []*Block{_ddac}, ctx, nil
}
func _cacf() *PageBreak { return &PageBreak{} }

// AddInternalLink adds a new internal link to the paragraph.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
// The zoom of the destination page is controlled with the zoom
// parameter. Pass in 0 to keep the current zoom value.
func (_bdca *StyledParagraph) AddInternalLink(text string, page int64, x, y, zoom float64) *TextChunk {
	_ddaag := NewTextChunk(text, _bdca._aacda)
	_ddaag._adbcg = _fade(page-1, x, y, zoom)
	return _bdca.appendChunk(_ddaag)
}

// GeneratePageBlocks draws the curve onto page blocks.
func (_cgfga *Curve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bddf := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_cgad := _fc.NewContentCreator()
	_cgad.Add_q().Add_w(_cgfga._faca).SetStrokingColor(_afag(_cgfga._eeb)).Add_m(_cgfga._aacg, ctx.PageHeight-_cgfga._eagd).Add_v(_cgfga._aade, ctx.PageHeight-_cgfga._cbec, _cgfga._cceec, ctx.PageHeight-_cgfga._cdbf).Add_S().Add_Q()
	_dff := _bddf.addContentsByString(_cgad.String())
	if _dff != nil {
		return nil, ctx, _dff
	}
	return []*Block{_bddf}, ctx, nil
}

// SkipRows skips over a specified number of rows in the table.
func (_bdbc *Table) SkipRows(num int) {
	_cgecb := num*_bdbc._fabb - 1
	if _cgecb < 0 {
		_da.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	_bdbc._daab += _cgecb
}
func (_aeff *Invoice) generateHeaderBlocks(_acgc DrawContext) ([]*Block, DrawContext, error) {
	_ggeg := _fbgb(_aeff._bfgc)
	_ggeg.SetEnableWrap(true)
	_ggeg.Append(_aeff._cbde)
	_fbdfd := _dbfc(2)
	if _aeff._abfb != nil {
		_fbfd := _fbdfd.NewCell()
		_fbfd.SetHorizontalAlignment(CellHorizontalAlignmentLeft)
		_fbfd.SetVerticalAlignment(CellVerticalAlignmentMiddle)
		_fbfd.SetIndent(0)
		_fbfd.SetContent(_aeff._abfb)
		_aeff._abfb.ScaleToHeight(_ggeg.Height() + 20)
	} else {
		_fbdfd.SkipCells(1)
	}
	_gffad := _fbdfd.NewCell()
	_gffad.SetHorizontalAlignment(CellHorizontalAlignmentRight)
	_gffad.SetVerticalAlignment(CellVerticalAlignmentMiddle)
	_gffad.SetContent(_ggeg)
	return _fbdfd.GeneratePageBlocks(_acgc)
}

// SetBorderColor sets the cell's border color.
func (_eafda *TableCell) SetBorderColor(col Color) {
	_eafda._faaea = col
	_eafda._becd = col
	_eafda._cedab = col
	_eafda._bdcf = col
}
func _fffb(_dggag *_g.PdfAnnotation) *_g.PdfAnnotation {
	if _dggag == nil {
		return nil
	}
	var _fafa *_g.PdfAnnotation
	switch _edead := _dggag.GetContext().(type) {
	case *_g.PdfAnnotationLink:
		if _feed := _fefe(_edead); _feed != nil {
			_fafa = _feed.PdfAnnotation
		}
	}
	return _fafa
}

// SetMargins sets the margins TOC line.
func (_gaaeca *TOCLine) SetMargins(left, right, top, bottom float64) {
	_gaaeca._egde = left
	_cadbf := &_gaaeca._daedd._cdbab
	_cadbf.Left = _gaaeca._egde + float64(_gaaeca._cefb-1)*_gaaeca._dedcg
	_cadbf.Right = right
	_cadbf.Top = top
	_cadbf.Bottom = bottom
}
func (_fgedb *TextChunk) clone() *TextChunk {
	_cfbc := *_fgedb
	_cfbc._adbcg = _fffb(_fgedb._adbcg)
	return &_cfbc
}
func (_adb *Block) setOpacity(_gf float64, _be float64) (string, error) {
	if (_gf < 0 || _gf >= 1.0) && (_be < 0 || _be >= 1.0) {
		return "", nil
	}
	_dbf := 0
	_fee := _ad.Sprintf("\u0047\u0053\u0025\u0064", _dbf)
	for _adb._fga.HasExtGState(_ec.PdfObjectName(_fee)) {
		_dbf++
		_fee = _ad.Sprintf("\u0047\u0053\u0025\u0064", _dbf)
	}
	_dcd := _ec.MakeDict()
	if _gf >= 0 && _gf < 1.0 {
		_dcd.Set("\u0063\u0061", _ec.MakeFloat(_gf))
	}
	if _be >= 0 && _be < 1.0 {
		_dcd.Set("\u0043\u0041", _ec.MakeFloat(_be))
	}
	_cee := _adb._fga.AddExtGState(_ec.PdfObjectName(_fee), _dcd)
	if _cee != nil {
		return "", _cee
	}
	return _fee, nil
}

// NewTextChunk returns a new text chunk instance.
func NewTextChunk(text string, style TextStyle) *TextChunk {
	return &TextChunk{Text: text, Style: style}
}

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_ffead *StyledParagraph) GetMargins() (float64, float64, float64, float64) {
	return _ffead._cdbab.Left, _ffead._cdbab.Right, _ffead._cdbab.Top, _ffead._cdbab.Bottom
}

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_eafg *Paragraph) GetMargins() (float64, float64, float64, float64) {
	return _eafg._bbdee.Left, _eafg._bbdee.Right, _eafg._bbdee.Top, _eafg._bbdee.Bottom
}

// SetSideBorderColor sets the cell's side border color.
func (_dage *TableCell) SetSideBorderColor(side CellBorderSide, col Color) {
	switch side {
	case CellBorderSideTop:
		_dage._bdcf = col
	case CellBorderSideBottom:
		_dage._becd = col
	case CellBorderSideLeft:
		_dage._faaea = col
	case CellBorderSideRight:
		_dage._cedab = col
	}
}

// Width is not used. Not used as a Table element is designed to fill into
// available width depending on the context. Returns 0.
func (_eaeed *Table) Width() float64 { return 0 }

type marginDrawable interface {
	VectorDrawable
	GetMargins() (float64, float64, float64, float64)
}

// Logo returns the logo of the invoice.
func (_eebc *Invoice) Logo() *Image { return _eebc._abfb }

var (
	ColorBlack  = ColorRGBFromArithmetic(0, 0, 0)
	ColorWhite  = ColorRGBFromArithmetic(1, 1, 1)
	ColorRed    = ColorRGBFromArithmetic(1, 0, 0)
	ColorGreen  = ColorRGBFromArithmetic(0, 1, 0)
	ColorBlue   = ColorRGBFromArithmetic(0, 0, 1)
	ColorYellow = ColorRGBFromArithmetic(1, 1, 0)
)

// Append adds a new text chunk to the paragraph.
func (_fgae *StyledParagraph) Append(text string) *TextChunk {
	_ecgag := NewTextChunk(text, _fgae._fgegg)
	return _fgae.appendChunk(_ecgag)
}

// SetAddressStyle sets the style properties used to render the content of
// the invoice address sections.
func (_gafb *Invoice) SetAddressStyle(style TextStyle) { _gafb._ceadb = style }
func (_eccb *Table) clone() *Table {
	_ccgb := *_eccb
	_ccgb._abeg = make([]float64, len(_eccb._abeg))
	copy(_ccgb._abeg, _eccb._abeg)
	_ccgb._dcdg = make([]float64, len(_eccb._dcdg))
	copy(_ccgb._dcdg, _eccb._dcdg)
	_ccgb._dcab = make([]*TableCell, 0, len(_eccb._dcab))
	for _, _dcbc := range _eccb._dcab {
		_caga := *_dcbc
		_caga._gfdee = &_ccgb
		_ccgb._dcab = append(_ccgb._dcab, &_caga)
	}
	return &_ccgb
}
func _cgfbc(_cded *Table, _gaff DrawContext) ([]*Block, DrawContext, error) {
	var _ebgg []*Block
	_fbcac := NewBlock(_gaff.PageWidth, _gaff.PageHeight)
	_cded.updateRowHeights(_gaff.Width - _cded._gcaee.Left - _cded._gcaee.Right)
	_ecfe := _cded._gcaee.Top
	if _cded._dbedc.IsRelative() && !_cded._ebcca {
		_eafa := _cded.Height()
		if _eafa > _gaff.Height-_cded._gcaee.Top && _eafa <= _gaff.PageHeight-_gaff.Margins.Top-_gaff.Margins.Bottom {
			_ebgg = []*Block{NewBlock(_gaff.PageWidth, _gaff.PageHeight-_gaff.Y)}
			var _degf error
			if _, _gaff, _degf = _cacf().GeneratePageBlocks(_gaff); _degf != nil {
				return nil, _gaff, _degf
			}
			_ecfe = 0
		}
	}
	_fdce := _gaff
	if _cded._dbedc.IsAbsolute() {
		_gaff.X = _cded._dfedb
		_gaff.Y = _cded._dccgd
	} else {
		_gaff.X += _cded._gcaee.Left
		_gaff.Y += _ecfe
		_gaff.Width -= _cded._gcaee.Left + _cded._gcaee.Right
		_gaff.Height -= _ecfe
	}
	_edde := _gaff.Width
	_faeb := _gaff.X
	_ffcd := _gaff.Y
	_eecf := _gaff.Height
	_dbcc := 0
	_bgcaa, _cacb := -1, -1
	if _cded._gbcfa {
		for _dgeeb, _fbada := range _cded._dcab {
			if _fbada._cbfc < _cded._dced {
				continue
			}
			if _fbada._cbfc > _cded._dgfbg {
				break
			}
			if _bgcaa < 0 {
				_bgcaa = _dgeeb
			}
			_cacb = _dgeeb
		}
	}
	var (
		_gffg  bool
		_gcga  int
		_gaeea int
		_ffbeg bool
		_eeea  int
		_ecfge error
	)
	for _agee := 0; _agee < len(_cded._dcab); _agee++ {
		_fffee := _cded._dcab[_agee]
		_addg := _fffee.width(_cded._dcdg, _edde)
		_dbcge := float64(0.0)
		for _fgdcf := 0; _fgdcf < _fffee._gbdde-1; _fgdcf++ {
			_dbcge += _cded._dcdg[_fgdcf] * _edde
		}
		_bcegg := float64(0.0)
		for _bbbf := _dbcc; _bbbf < _fffee._cbfc-1; _bbbf++ {
			_bcegg += _cded._abeg[_bbbf]
		}
		_gaff.Height = _eecf - _bcegg
		_fgbg := float64(0.0)
		for _eaeb := 0; _eaeb < _fffee._cdcd; _eaeb++ {
			_fgbg += _cded._abeg[_fffee._cbfc+_eaeb-1]
		}
		_fdcd := _ffbeg && _fffee._cbfc != _eeea
		_eeea = _fffee._cbfc
		if _fdcd || _fgbg > _gaff.Height {
			if _cded._adce && !_ffbeg {
				_ffbeg, _ecfge = _cded.wrapRow(_agee, _gaff, _edde)
				if _ecfge != nil {
					return nil, _gaff, _ecfge
				}
				if _ffbeg {
					_agee--
					continue
				}
			}
			_ebgg = append(_ebgg, _fbcac)
			_fbcac = NewBlock(_gaff.PageWidth, _gaff.PageHeight)
			_faeb = _gaff.Margins.Left + _cded._gcaee.Left
			_ffcd = _gaff.Margins.Top
			_gaff.Height = _gaff.PageHeight - _gaff.Margins.Top - _gaff.Margins.Bottom
			_gaff.Page++
			_eecf = _gaff.Height
			_dbcc = _fffee._cbfc - 1
			_bcegg = 0
			_ffbeg = false
			if _cded._gbcfa && _bgcaa >= 0 {
				_gcga = _agee
				_agee = _bgcaa - 1
				_gaeea = _dbcc
				_dbcc = _cded._dced - 1
				_gffg = true
				continue
			}
			if _fdcd {
				_agee--
				continue
			}
		}
		_gaff.Width = _addg
		_gaff.X = _faeb + _dbcge
		_gaff.Y = _ffcd + _bcegg
		_bcfc := _cda(_gaff.X, _gaff.Y, _addg, _fgbg)
		if _fffee._dcbgd != nil {
			_bcfc.SetFillColor(_fffee._dcbgd)
		}
		_bcfc.LineStyle = _fffee._egba
		_bcfc._gda = _fffee._dbcf
		_bcfc._gga = _fffee._eacc
		_bcfc._ddc = _fffee._aagb
		_bcfc._bedg = _fffee._ebfg
		if _fffee._faaea != nil {
			_bcfc.SetColorLeft(_fffee._faaea)
		}
		if _fffee._becd != nil {
			_bcfc.SetColorBottom(_fffee._becd)
		}
		if _fffee._cedab != nil {
			_bcfc.SetColorRight(_fffee._cedab)
		}
		if _fffee._bdcf != nil {
			_bcfc.SetColorTop(_fffee._bdcf)
		}
		_bcfc.SetWidthBottom(_fffee._eecfg)
		_bcfc.SetWidthLeft(_fffee._geaef)
		_bcfc.SetWidthRight(_fffee._edgb)
		_bcfc.SetWidthTop(_fffee._eaeg)
		_efbg := _fbcac.Draw(_bcfc)
		if _efbg != nil {
			_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _efbg)
		}
		if _fffee._adeed != nil {
			_acdd := _fffee._adeed.Width()
			_deea := _fffee._adeed.Height()
			_abcd := 0.0
			switch _dgaddb := _fffee._adeed.(type) {
			case *Paragraph:
				if _dgaddb._acegd {
					_acdd = _dgaddb.getMaxLineWidth() / 1000.0
				}
				_acdd += _dgaddb._bbdee.Left + _dgaddb._bbdee.Right
				_deea += _dgaddb._bbdee.Top + _dgaddb._bbdee.Bottom
			case *StyledParagraph:
				if _dgaddb._cfga {
					_acdd = _dgaddb.getMaxLineWidth() / 1000.0
				}
				_bcbe, _cdaff, _ceee := _dgaddb.getLineMetrics(0)
				_abdbc, _fbga := _bcbe*_dgaddb._bggb, _cdaff*_dgaddb._bggb
				if _dgaddb._fgce == TextVerticalAlignmentCenter {
					_abcd = _fbga - (_cdaff + (_bcbe+_ceee-_cdaff)/2 + (_fbga-_cdaff)/2)
				}
				if len(_dgaddb._aageb) == 1 {
					_deea = _abdbc
				} else {
					_deea = _deea - _fbga + _abdbc
				}
				_abcd += _abdbc - _fbga
				switch _fffee._dfcfe {
				case CellVerticalAlignmentTop:
					_abcd += _abdbc * 0.5
				case CellVerticalAlignmentBottom:
					_abcd -= _abdbc * 0.5
				}
				_acdd += _dgaddb._cdbab.Left + _dgaddb._cdbab.Right
				_deea += _dgaddb._cdbab.Top + _dgaddb._cdbab.Bottom
			case *Table:
				_acdd = _addg
			case *List:
				_acdd = _addg
			case *Division:
				_acdd = _addg
			case *Chart:
				_acdd = _addg
			}
			switch _fffee._dcee {
			case CellHorizontalAlignmentLeft:
				_gaff.X += _fffee._fffcb
				_gaff.Width -= _fffee._fffcb
			case CellHorizontalAlignmentCenter:
				if _agcb := _addg - _acdd; _agcb > 0 {
					_gaff.X += _agcb / 2
					_gaff.Width -= _agcb / 2
				}
			case CellHorizontalAlignmentRight:
				if _addg > _acdd {
					_gaff.X = _gaff.X + _addg - _acdd - _fffee._fffcb
					_gaff.Width -= _fffee._fffcb
				}
			}
			_gaff.Y += _abcd
			switch _fffee._dfcfe {
			case CellVerticalAlignmentTop:
			case CellVerticalAlignmentMiddle:
				if _fdbc := _fgbg - _deea; _fdbc > 0 {
					_gaff.Y += _fdbc / 2
					_gaff.Height -= _fdbc / 2
				}
			case CellVerticalAlignmentBottom:
				if _fgbg > _deea {
					_gaff.Y = _gaff.Y + _fgbg - _deea
					_gaff.Height = _fgbg
				}
			}
			_faccd := _fbcac.DrawWithContext(_fffee._adeed, _gaff)
			if _faccd != nil {
				_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _faccd)
			}
			_gaff.Y -= _abcd
		}
		_gaff.Y += _fgbg
		_gaff.Height -= _fgbg
		if _gffg && _agee+1 > _cacb {
			_ffcd += _bcegg + _fgbg
			_eecf -= _fgbg + _bcegg
			_dbcc = _gaeea
			_agee = _gcga - 1
			_gffg = false
		}
	}
	_ebgg = append(_ebgg, _fbcac)
	if _cded._dbedc.IsAbsolute() {
		return _ebgg, _fdce, nil
	}
	_gaff.X = _fdce.X
	_gaff.Width = _fdce.Width
	_gaff.Y += _cded._gcaee.Bottom
	_gaff.Height -= _cded._gcaee.Bottom
	return _ebgg, _gaff, nil
}

// CellBorderSide defines the table cell's border side.
type CellBorderSide int

// Positioning represents the positioning type for drawing creator components (relative/absolute).
type Positioning int

// SetAngle sets the rotation angle of the text.
func (_fedf *StyledParagraph) SetAngle(angle float64) { _fedf._ffgac = angle }
func _dbfc(_gffc int) *Table {
	_debbc := &Table{_fabb: _gffc, _badaf: 10.0, _dcdg: []float64{}, _abeg: []float64{}, _dcab: []*TableCell{}, _aegce: make([]int, _gffc), _ebcca: true}
	_debbc.resetColumnWidths()
	return _debbc
}

// The Image type is used to draw an image onto PDF.
type Image struct {
	_edad         *_g.XObjectImage
	_ggdgd        *_g.Image
	_dfdb         float64
	_ebb, _gcab   float64
	_efeb, _aegc  float64
	_cbce         Positioning
	_eecd         HorizontalAlignment
	_edcgf        float64
	_baef         float64
	_ebada        float64
	_ecbf         Margins
	_caadd, _aagf float64
	_fdgf         _ec.StreamEncoder
}

func _cda(_beca, _cfbf, _agg, _abb float64) *border {
	_gecd := &border{}
	_gecd._edd = _beca
	_gecd._ccga = _cfbf
	_gecd._debf = _agg
	_gecd._cfb = _abb
	_gecd._acg = ColorBlack
	_gecd._aad = ColorBlack
	_gecd._gcbe = ColorBlack
	_gecd._edf = ColorBlack
	_gecd._bdf = 0
	_gecd._gfg = 0
	_gecd._age = 0
	_gecd._cedc = 0
	_gecd.LineStyle = _bb.LineStyleSolid
	return _gecd
}

// CurRow returns the currently active cell's row number.
func (_efdd *Table) CurRow() int { _fbda := (_efdd._daab-1)/_efdd._fabb + 1; return _fbda }
func _cfe(_bcc *_fc.ContentStreamOperations, _caa *_g.PdfPageResources, _eef *_fc.ContentStreamOperations, _cgag *_g.PdfPageResources) error {
	_gacg := map[_ec.PdfObjectName]_ec.PdfObjectName{}
	_gc := map[_ec.PdfObjectName]_ec.PdfObjectName{}
	_eec := map[_ec.PdfObjectName]_ec.PdfObjectName{}
	_acc := map[_ec.PdfObjectName]_ec.PdfObjectName{}
	_gaca := map[_ec.PdfObjectName]_ec.PdfObjectName{}
	_fbd := map[_ec.PdfObjectName]_ec.PdfObjectName{}
	for _, _baf := range *_eef {
		switch _baf.Operand {
		case "\u0044\u006f":
			if len(_baf.Params) == 1 {
				if _bdc, _ab := _baf.Params[0].(*_ec.PdfObjectName); _ab {
					if _, _bgf := _gacg[*_bdc]; !_bgf {
						var _cdc _ec.PdfObjectName
						_gfba, _ := _cgag.GetXObjectByName(*_bdc)
						if _gfba != nil {
							_cdc = *_bdc
							for {
								_bba, _ := _caa.GetXObjectByName(_cdc)
								if _bba == nil || _bba == _gfba {
									break
								}
								_cdc = _cdc + "\u0030"
							}
						}
						_caa.SetXObjectByName(_cdc, _gfba)
						_gacg[*_bdc] = _cdc
					}
					_ged := _gacg[*_bdc]
					_baf.Params[0] = &_ged
				}
			}
		case "\u0054\u0066":
			if len(_baf.Params) == 2 {
				if _ffc, _gd := _baf.Params[0].(*_ec.PdfObjectName); _gd {
					if _, _bag := _gc[*_ffc]; !_bag {
						_gad, _ccg := _cgag.GetFontByName(*_ffc)
						_dgg := *_ffc
						if _ccg && _gad != nil {
							_dgg = _afb(_ffc.String(), _gad, _caa)
						}
						_caa.SetFontByName(_dgg, _gad)
						_gc[*_ffc] = _dgg
					}
					_gaf := _gc[*_ffc]
					_baf.Params[0] = &_gaf
				}
			}
		case "\u0043\u0053", "\u0063\u0073":
			if len(_baf.Params) == 1 {
				if _accb, _eeg := _baf.Params[0].(*_ec.PdfObjectName); _eeg {
					if _, _cgc := _eec[*_accb]; !_cgc {
						var _afa _ec.PdfObjectName
						_bec, _gcb := _cgag.GetColorspaceByName(*_accb)
						if _gcb {
							_afa = *_accb
							for {
								_gcg, _bgg := _caa.GetColorspaceByName(_afa)
								if !_bgg || _bec == _gcg {
									break
								}
								_afa = _afa + "\u0030"
							}
							_caa.SetColorspaceByName(_afa, _bec)
							_eec[*_accb] = _afa
						} else {
							_da.Log.Debug("C\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064")
						}
					}
					if _ffd, _dee := _eec[*_accb]; _dee {
						_baf.Params[0] = &_ffd
					} else {
						_da.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", *_accb)
					}
				}
			}
		case "\u0053\u0043\u004e", "\u0073\u0063\u006e":
			if len(_baf.Params) == 1 {
				if _ccb, _cgf := _baf.Params[0].(*_ec.PdfObjectName); _cgf {
					if _, _bagf := _acc[*_ccb]; !_bagf {
						var _feg _ec.PdfObjectName
						_aa, _aae := _cgag.GetPatternByName(*_ccb)
						if _aae {
							_feg = *_ccb
							for {
								_fgeg, _eg := _caa.GetPatternByName(_feg)
								if !_eg || _fgeg == _aa {
									break
								}
								_feg = _feg + "\u0030"
							}
							_fae := _caa.SetPatternByName(_feg, _aa.ToPdfObject())
							if _fae != nil {
								return _fae
							}
							_acc[*_ccb] = _feg
						}
					}
					if _bbag, _aed := _acc[*_ccb]; _aed {
						_baf.Params[0] = &_bbag
					}
				}
			}
		case "\u0073\u0068":
			if len(_baf.Params) == 1 {
				if _feb, _fda := _baf.Params[0].(*_ec.PdfObjectName); _fda {
					if _, _gcf := _gaca[*_feb]; !_gcf {
						var _bed _ec.PdfObjectName
						_eed, _bgac := _cgag.GetShadingByName(*_feb)
						if _bgac {
							_bed = *_feb
							for {
								_afad, _bff := _caa.GetShadingByName(_bed)
								if !_bff || _eed == _afad {
									break
								}
								_bed = _bed + "\u0030"
							}
							_ced := _caa.SetShadingByName(_bed, _eed.ToPdfObject())
							if _ced != nil {
								_da.Log.Debug("E\u0052\u0052\u004f\u0052 S\u0065t\u0020\u0073\u0068\u0061\u0064i\u006e\u0067\u003a\u0020\u0025\u0076", _ced)
								return _ced
							}
							_gaca[*_feb] = _bed
						} else {
							_da.Log.Debug("\u0053\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
						}
					}
					if _dagf, _ecga := _gaca[*_feb]; _ecga {
						_baf.Params[0] = &_dagf
					} else {
						_da.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020S\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0073 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064", *_feb)
					}
				}
			}
		case "\u0067\u0073":
			if len(_baf.Params) == 1 {
				if _ede, _ddg := _baf.Params[0].(*_ec.PdfObjectName); _ddg {
					if _, _fccg := _fbd[*_ede]; !_fccg {
						var _deeb _ec.PdfObjectName
						_cbag, _gab := _cgag.GetExtGState(*_ede)
						if _gab {
							_deeb = *_ede
							_ffe := 1
							for {
								_gec, _dfc := _caa.GetExtGState(_deeb)
								if !_dfc || _cbag == _gec {
									break
								}
								_deeb = _ec.PdfObjectName(_ad.Sprintf("\u0047\u0053\u0025\u0064", _ffe))
								_ffe++
							}
						}
						_caa.AddExtGState(_deeb, _cbag)
						_fbd[*_ede] = _deeb
					}
					_dabd := _fbd[*_ede]
					_baf.Params[0] = &_dabd
				}
			}
		}
		*_bcc = append(*_bcc, _baf)
	}
	return nil
}
func (_fcbc *StyledParagraph) getMaxLineWidth() float64 {
	if _fcbc._aageb == nil || len(_fcbc._aageb) == 0 {
		_fcbc.wrapText()
	}
	var _deggb float64
	for _, _gddc := range _fcbc._aageb {
		_gbgg := _fcbc.getTextLineWidth(_gddc)
		if _gbgg > _deggb {
			_deggb = _gbgg
		}
	}
	return _deggb
}

// HorizontalAlignment represents the horizontal alignment of components
// within a page.
type HorizontalAlignment int

// SetBackgroundColor sets the cell's background color.
func (_gcbgg *TableCell) SetBackgroundColor(col Color) { _gcbgg._dcbgd = col }

// UnsupportedRuneError is an error that occurs when there is unsupported glyph being used.
type UnsupportedRuneError struct {
	Message string
	Rune    rune
}

// SetColorTop sets border color for top.
func (_gef *border) SetColorTop(col Color) { _gef._acg = col }
func (_bdfgc *TOCLine) getLineLink() *_g.PdfAnnotation {
	if _bdfgc._eegc <= 0 {
		return nil
	}
	return _fade(_bdfgc._eegc-1, _bdfgc._gddbe, _bdfgc._cdde, 0)
}

// SetBorderWidth sets the border width.
func (_afcb *CurvePolygon) SetBorderWidth(borderWidth float64) { _afcb._bafd.BorderWidth = borderWidth }

// Drawable is a widget that can be used to draw with the Creator.
type Drawable interface {

	// GeneratePageBlocks draw onto blocks representing Page contents. As the content can wrap over many pages, multiple
	// templates are returned, one per Page.  The function also takes a draw context containing information
	// where to draw (if relative positioning) and the available height to draw on accounting for Margins etc.
	GeneratePageBlocks(_accf DrawContext) ([]*Block, DrawContext, error)
}

// SetMargins sets the Block's left, right, top, bottom, margins.
func (_faf *Block) SetMargins(left, right, top, bottom float64) {
	_faf._ae.Left = left
	_faf._ae.Right = right
	_faf._ae.Top = top
	_faf._ae.Bottom = bottom
}
func (_ecddf *Invoice) setCellBorder(_aeb *TableCell, _dffa *InvoiceCell) {
	for _, _dggf := range _dffa.BorderSides {
		_aeb.SetBorder(_dggf, CellBorderStyleSingle, _dffa.BorderWidth)
	}
	_aeb.SetBorderColor(_dffa.BorderColor)
}

// SetBorderOpacity sets the border opacity.
func (_dbdf *Polygon) SetBorderOpacity(opacity float64) { _dbdf._daae = opacity }

// Curve represents a cubic Bezier curve with a control point.
type Curve struct {
	_aacg  float64
	_eagd  float64
	_aade  float64
	_cbec  float64
	_cceec float64
	_cdbf  float64
	_eeb   Color
	_faca  float64
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
func (_edef *Creator) Finalize() error {
	if _edef._ggce {
		return nil
	}
	_gdbb := len(_edef._bdb)
	_ccea := 0
	if _edef._gace != nil {
		_dge := *_edef
		_edef._bdb = nil
		_edef._abed = nil
		_edef.initContext()
		_gceg := FrontpageFunctionArgs{PageNum: 1, TotalPages: _gdbb}
		_edef._gace(_gceg)
		_ccea += len(_edef._bdb)
		_edef._bdb = _dge._bdb
		_edef._abed = _dge._abed
	}
	if _edef.AddTOC {
		_edef.initContext()
		_edef._defb.Page = _ccea + 1
		if _edef._bbfa != nil {
			if _fde := _edef._bbfa(_edef._fgc); _fde != nil {
				return _fde
			}
		}
		_abag, _, _bffd := _edef._fgc.GeneratePageBlocks(_edef._defb)
		if _bffd != nil {
			_da.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073: \u0025\u0076", _bffd)
			return _bffd
		}
		_ccea += len(_abag)
		_edgf := _edef._fgc.Lines()
		for _, _gcfe := range _edgf {
			_eafd, _ece := _cf.Atoi(_gcfe.Page.Text)
			if _ece != nil {
				continue
			}
			_gcfe.Page.Text = _cf.Itoa(_eafd + _ccea)
		}
	}
	_acegg := false
	var _feaa []*_g.PdfPage
	if _edef._gace != nil {
		_eda := *_edef
		_edef._bdb = nil
		_edef._abed = nil
		_cdg := FrontpageFunctionArgs{PageNum: 1, TotalPages: _gdbb}
		_edef._gace(_cdg)
		_gdbb += len(_edef._bdb)
		_feaa = _edef._bdb
		_edef._bdb = append(_edef._bdb, _eda._bdb...)
		_edef._abed = _eda._abed
		_acegg = true
	}
	var _gba []*_g.PdfPage
	if _edef.AddTOC {
		_edef.initContext()
		if _edef._bbfa != nil {
			if _abc := _edef._bbfa(_edef._fgc); _abc != nil {
				_da.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _abc)
				return _abc
			}
		}
		_bbfc := _edef._fgc.Lines()
		for _, _ceb := range _bbfc {
			_ceb._eegc += int64(_ccea)
		}
		_eefga, _, _ := _edef._fgc.GeneratePageBlocks(_edef._defb)
		for _, _gdbc := range _eefga {
			_gdbc.SetPos(0, 0)
			_gdbb++
			_badb := _edef.newPage()
			_gba = append(_gba, _badb)
			_edef.setActivePage(_badb)
			_edef.Draw(_gdbc)
		}
		if _acegg {
			_cgfb := _feaa
			_ded := _edef._bdb[len(_feaa):]
			_edef._bdb = append([]*_g.PdfPage{}, _cgfb...)
			_edef._bdb = append(_edef._bdb, _gba...)
			_edef._bdb = append(_edef._bdb, _ded...)
		} else {
			_edef._bdb = append(_gba, _edef._bdb...)
		}
	}
	if _edef._edge != nil && _edef.AddOutlines {
		var _fgb func(_aeac *_g.OutlineItem)
		_fgb = func(_eaee *_g.OutlineItem) {
			_eaee.Dest.Page += int64(_ccea)
			if _dbcg := int(_eaee.Dest.Page); _dbcg >= 0 && _dbcg < len(_edef._bdb) {
				_eaee.Dest.PageObj = _edef._bdb[_dbcg].GetPageAsIndirectObject()
			} else {
				_da.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _dbcg)
			}
			_eaee.Dest.Y = _edef._cggd - _eaee.Dest.Y
			_feaf := _eaee.Items()
			for _, _dcae := range _feaf {
				_fgb(_dcae)
			}
		}
		_cbaf := _edef._edge.Items()
		for _, _decd := range _cbaf {
			_fgb(_decd)
		}
		if _edef.AddTOC {
			var _eggd int
			if _acegg {
				_eggd = len(_feaa)
			}
			_aeg := _g.NewOutlineDest(int64(_eggd), 0, _edef._cggd)
			if _eggd >= 0 && _eggd < len(_edef._bdb) {
				_aeg.PageObj = _edef._bdb[_eggd].GetPageAsIndirectObject()
			} else {
				_da.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _eggd)
			}
			_edef._edge.Insert(0, _g.NewOutlineItem("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073", _aeg))
		}
	}
	for _gged, _gebf := range _edef._bdb {
		_edef.setActivePage(_gebf)
		if _edef._fbfgd != nil {
			_gbag, _ccf, _fef := _gebf.Size()
			if _fef != nil {
				return _fef
			}
			_ecd := PageFinalizeFunctionArgs{PageNum: _gged + 1, PageWidth: _gbag, PageHeight: _ccf, TOCPages: len(_gba), TotalPages: _gdbb}
			if _ecfg := _edef._fbfgd(_ecd); _ecfg != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0070\u0061\u0067\u0065\u0020\u0066\u0069\u006e\u0061\u006c\u0069\u007a\u0065 \u0063\u0061\u006c\u006c\u0062\u0061\u0063k\u003a\u0020\u0025\u0076", _ecfg)
				return _ecfg
			}
		}
		if _edef._gce != nil {
			_bfedc := NewBlock(_edef._cgb, _edef._acda.Top)
			_acegc := HeaderFunctionArgs{PageNum: _gged + 1, TotalPages: _gdbb}
			_edef._gce(_bfedc, _acegc)
			_bfedc.SetPos(0, 0)
			if _dfcd := _edef.Draw(_bfedc); _dfcd != nil {
				_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0068e\u0061\u0064\u0065\u0072\u003a\u0020\u0025v", _dfcd)
				return _dfcd
			}
		}
		if _edef._bfag != nil {
			_ffeb := NewBlock(_edef._cgb, _edef._acda.Bottom)
			_fccb := FooterFunctionArgs{PageNum: _gged + 1, TotalPages: _gdbb}
			_edef._bfag(_ffeb, _fccb)
			_ffeb.SetPos(0, _edef._cggd-_ffeb._df)
			if _acbf := _edef.Draw(_ffeb); _acbf != nil {
				_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0066o\u006f\u0074\u0065\u0072\u003a\u0020\u0025v", _acbf)
				return _acbf
			}
		}
		_cdda, _gffa := _edef._gff[_gebf]
		if _dfbc, _bcabe := _edef._fcac[_gebf]; _bcabe {
			if _gffa {
				_cdda.transformBlock(_dfbc)
			}
			if _eab := _dfbc.drawToPage(_gebf); _eab != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020%\u0064\u0020\u0062\u006c\u006f\u0063\u006bs\u003a\u0020\u0025\u0076", _gged+1, _eab)
				return _eab
			}
		}
		if _gffa {
			if _aegb := _cdda.transformPage(_gebf); _aegb != nil {
				_da.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0074\u0072\u0061\u006e\u0073f\u006f\u0072\u006d\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0076", _aegb)
				return _aegb
			}
		}
	}
	_edef._ggce = true
	return nil
}

// SetTerms sets the terms and conditions section of the invoice.
func (_cggb *Invoice) SetTerms(title, content string) { _cggb._fcega = [2]string{title, content} }

// SetAngle sets the rotation angle of the text.
func (_ecea *Paragraph) SetAngle(angle float64) { _ecea._cbccg = angle }
func (_aagbg *TableCell) height(_dbfba float64) float64 {
	var _dfaf float64
	switch _affe := _aagbg._adeed.(type) {
	case *Paragraph:
		if _affe._acegd {
			_affe.SetWidth(_dbfba - _aagbg._fffcb - _affe._bbdee.Left - _affe._bbdee.Right)
		}
		_dfaf = _affe.Height() + _affe._bbdee.Top + _affe._bbdee.Bottom + 0.5*_affe._ebbae*_affe._dfgg
	case *StyledParagraph:
		if _affe._cfga {
			_affe.SetWidth(_dbfba - _aagbg._fffcb - _affe._cdbab.Left - _affe._cdbab.Right)
		}
		_dfaf = _affe.Height() + _affe._cdbab.Top + _affe._cdbab.Bottom + 0.5*_affe.getTextHeight()
	case *Image:
		_dfaf = _affe.Height() + _affe._ecbf.Top + _affe._ecbf.Bottom
	case *Table:
		_affe.updateRowHeights(_dbfba - _aagbg._fffcb - _affe._gcaee.Left - _affe._gcaee.Right)
		_dfaf = _affe.Height() + _affe._gcaee.Top + _affe._gcaee.Bottom
	case *List:
		_dfaf = _affe.tableHeight(_dbfba-_aagbg._fffcb) + _affe._cgab.Top + _affe._cgab.Bottom
	case *Division:
		_dfaf = _affe.ctxHeight(_dbfba-_aagbg._fffcb) + _affe._abdb.Top + _affe._abdb.Bottom
	case *Chart:
		_dfaf = _affe.Height() + _affe._gedd.Top + _affe._gedd.Bottom
	}
	return _dfaf
}

// IsRelative checks if the positioning is relative.
func (_dcf Positioning) IsRelative() bool { return _dcf == PositionRelative }

// Table allows organizing content in an rows X columns matrix, which can spawn across multiple pages.
type Table struct {
	_gbcf          int
	_fabb          int
	_daab          int
	_dcdg          []float64
	_abeg          []float64
	_badaf         float64
	_dcab          []*TableCell
	_aegce         []int
	_dbedc         Positioning
	_dfedb, _dccgd float64
	_gcaee         Margins
	_gbcfa         bool
	_dced          int
	_dgfbg         int
	_adce          bool
	_ebcca         bool
}

// GeneratePageBlocks generate the Page blocks. Draws the Image on a block, implementing the Drawable interface.
func (_begc *Image) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	if _begc._edad == nil {
		if _eedg := _begc.makeXObject(); _eedg != nil {
			return nil, ctx, _eedg
		}
	}
	var _eade []*Block
	_bcba := ctx
	_ebfe := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _begc._cbce.IsRelative() {
		ctx.X += _begc._ecbf.Left
		ctx.Y += _begc._ecbf.Top
		ctx.Width -= _begc._ecbf.Left + _begc._ecbf.Right
		ctx.Height -= _begc._ecbf.Top + _begc._ecbf.Bottom
		if _begc._gcab > ctx.Height {
			_eade = append(_eade, _ebfe)
			_ebfe = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_fgfb := ctx
			_fgfb.Y = ctx.Margins.Top + _begc._ecbf.Top
			_fgfb.X = ctx.Margins.Left + _begc._ecbf.Left
			_fgfb.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _begc._ecbf.Top - _begc._ecbf.Bottom
			_fgfb.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _begc._ecbf.Left - _begc._ecbf.Right
			ctx = _fgfb
			_bcba.X = ctx.Margins.Left
			_bcba.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right
		}
	} else {
		ctx.X = _begc._edcgf
		ctx.Y = _begc._baef
	}
	ctx, _fdeg := _dde(_ebfe, _begc, ctx)
	if _fdeg != nil {
		return nil, ctx, _fdeg
	}
	_eade = append(_eade, _ebfe)
	if _begc._cbce.IsAbsolute() {
		ctx = _bcba
	} else {
		ctx.X = _bcba.X
		ctx.Y += _begc._ecbf.Bottom
		ctx.Width = _bcba.Width
	}
	return _eade, ctx, nil
}

// SetOpacity sets opacity for Image.
func (_aced *Image) SetOpacity(opacity float64) { _aced._ebada = opacity }

// SetFillOpacity sets the fill opacity.
func (_bccbb *Polygon) SetFillOpacity(opacity float64) { _bccbb._aebd = opacity }

// Height returns the height of the list.
func (_cafd *List) Height() float64 {
	var _bbcfg float64
	for _, _ebba := range _cafd._cbba {
		_bbcfg += _ebba._dege.Height()
	}
	return _bbcfg
}

// TextRenderingMode determines whether showing text shall cause glyph
// outlines to be stroked, filled, used as a clipping boundary, or some
// combination of the three.
// See section 9.3 "Text State Parameters and Operators" and
// Table 106 (pp. 254-255 PDF32000_2008).
type TextRenderingMode int

// SetFillOpacity sets the fill opacity.
func (_ecfb *Rectangle) SetFillOpacity(opacity float64) { _ecfb._fcae = opacity }
func _dedcc(_eeca *Block, _gaeeb *Paragraph, _dceg DrawContext) (DrawContext, error) {
	_dgfd := 1
	_bgda := _ec.PdfObjectName("\u0046\u006f\u006e\u0074" + _cf.Itoa(_dgfd))
	for _eeca._fga.HasFontByName(_bgda) {
		_dgfd++
		_bgda = _ec.PdfObjectName("\u0046\u006f\u006e\u0074" + _cf.Itoa(_dgfd))
	}
	_ffef := _eeca._fga.SetFontByName(_bgda, _gaeeb._cgbf.ToPdfObject())
	if _ffef != nil {
		return _dceg, _ffef
	}
	_gaeeb.wrapText()
	_eedba := _fc.NewContentCreator()
	_eedba.Add_q()
	_bffgb := _dceg.PageHeight - _dceg.Y - _gaeeb._ebbae*_gaeeb._dfgg
	_eedba.Translate(_dceg.X, _bffgb)
	if _gaeeb._cbccg != 0 {
		_eedba.RotateDeg(_gaeeb._cbccg)
	}
	_eedba.Add_BT().SetNonStrokingColor(_afag(_gaeeb._gade)).Add_Tf(_bgda, _gaeeb._ebbae).Add_TL(_gaeeb._ebbae * _gaeeb._dfgg)
	for _aeebb, _gfggf := range _gaeeb._egge {
		if _aeebb != 0 {
			_eedba.Add_Tstar()
		}
		_gcae := []rune(_gfggf)
		_dffab := 0.0
		_ggcee := 0
		for _adfa, _eeac := range _gcae {
			if _eeac == ' ' {
				_ggcee++
				continue
			}
			if _eeac == '\u000A' {
				continue
			}
			_dafa, _fgfd := _gaeeb._cgbf.GetRuneMetrics(_eeac)
			if !_fgfd {
				_da.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006e\u0065\u0020\u0069=\u0025\u0064\u0020\u0072\u0075\u006e\u0065=\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0069n\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0025\u0073", _adfa, _eeac, _eeac, _gaeeb._cgbf.BaseFont(), _gaeeb._cgbf.Subtype())
				return _dceg, _f.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
			}
			_dffab += _gaeeb._ebbae * _dafa.Wx
		}
		var _dfca []_ec.PdfObject
		_ggdd, _faga := _gaeeb._cgbf.GetRuneMetrics(' ')
		if !_faga {
			return _dceg, _f.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
		}
		_abfbb := _ggdd.Wx
		switch _gaeeb._dgadd {
		case TextAlignmentJustify:
			if _ggcee > 0 && _aeebb < len(_gaeeb._egge)-1 {
				_abfbb = (_gaeeb._dgee*1000.0 - _dffab) / float64(_ggcee) / _gaeeb._ebbae
			}
		case TextAlignmentCenter:
			_cbed := _dffab + float64(_ggcee)*_abfbb*_gaeeb._ebbae
			_gefa := (_gaeeb._dgee*1000.0 - _cbed) / 2 / _gaeeb._ebbae
			_dfca = append(_dfca, _ec.MakeFloat(-_gefa))
		case TextAlignmentRight:
			_agga := _dffab + float64(_ggcee)*_abfbb*_gaeeb._ebbae
			_ffdg := (_gaeeb._dgee*1000.0 - _agga) / _gaeeb._ebbae
			_dfca = append(_dfca, _ec.MakeFloat(-_ffdg))
		}
		_gfed := _gaeeb._cgbf.Encoder()
		var _dega []byte
		for _, _gbbd := range _gcae {
			if _gbbd == '\u000A' {
				continue
			}
			if _gbbd == ' ' {
				if len(_dega) > 0 {
					_dfca = append(_dfca, _ec.MakeStringFromBytes(_dega))
					_dega = nil
				}
				_dfca = append(_dfca, _ec.MakeFloat(-_abfbb))
			} else {
				if _, _bcce := _gfed.RuneToCharcode(_gbbd); !_bcce {
					_ffef = UnsupportedRuneError{Message: _ad.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _gbbd, _gbbd), Rune: _gbbd}
					_dceg._bgdc = append(_dceg._bgdc, _ffef)
					_da.Log.Debug(_ffef.Error())
					if _dceg._bfaf <= 0 {
						continue
					}
					_gbbd = _dceg._bfaf
				}
				_dega = append(_dega, _gfed.Encode(string(_gbbd))...)
			}
		}
		if len(_dega) > 0 {
			_dfca = append(_dfca, _ec.MakeStringFromBytes(_dega))
		}
		_eedba.Add_TJ(_dfca...)
	}
	_eedba.Add_ET()
	_eedba.Add_Q()
	_ddaf := _eedba.Operations()
	_ddaf.WrapIfNeeded()
	_eeca.addContents(_ddaf)
	if _gaeeb._fcgd.IsRelative() {
		_fdab := _gaeeb.Height()
		_dceg.Y += _fdab
		_dceg.Height -= _fdab
		if _dceg.Inline {
			_dceg.X += _gaeeb.Width() + _gaeeb._bbdee.Right
		}
	}
	return _dceg, nil
}

// SetBorderWidth sets the border width.
func (_bgfe *Ellipse) SetBorderWidth(bw float64) { _bgfe._gca = bw }

// Polyline represents a slice of points that are connected as straight lines.
// Implements the Drawable interface and can be rendered using the Creator.
type Polyline struct {
	_efcf  *_bb.Polyline
	_beccb float64
}

// AddressHeadingStyle returns the style properties used to render the
// heading of the invoice address sections.
func (_fcage *Invoice) AddressHeadingStyle() TextStyle { return _fcage._ggae }

// CellHorizontalAlignment defines the table cell's horizontal alignment.
type CellHorizontalAlignment int

func _gedg(_efgc string) (*Image, error) {
	_ebe, _bac := _c.Open(_efgc)
	if _bac != nil {
		return nil, _bac
	}
	defer _ebe.Close()
	_eegg, _bac := _g.ImageHandling.Read(_ebe)
	if _bac != nil {
		_da.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bac)
		return nil, _bac
	}
	return _ageb(_eegg)
}

// Width returns the Block's width.
func (_gfb *Block) Width() float64 { return _gfb._deb }

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

// MultiColCell makes a new cell with the specified column span and inserts it
// into the table at the current position.
func (_ecdb *Table) MultiColCell(colspan int) *TableCell { return _ecdb.MultiCell(1, colspan) }

// Height returns the height of the Paragraph. The height is calculated based on the input text and how it is wrapped
// within the container. Does not include Margins.
func (_efba *StyledParagraph) Height() float64 {
	_efba.wrapText()
	var _adeb float64
	for _, _fcgc := range _efba._aageb {
		var _ffed float64
		for _, _ffdec := range _fcgc {
			_agebf := _efba._bggb * _ffdec.Style.FontSize
			if _agebf > _ffed {
				_ffed = _agebf
			}
		}
		_adeb += _ffed
	}
	return _adeb
}
func (_fbcd *Invoice) generateLineBlocks(_ffbad DrawContext) ([]*Block, DrawContext, error) {
	_badae := _dbfc(len(_fbcd._cfc))
	_badae.SetMargins(0, 0, 25, 0)
	for _, _edac := range _fbcd._cfc {
		_bgcd := _fbgb(_edac.TextStyle)
		_bgcd.SetMargins(0, 0, 1, 0)
		_bgcd.Append(_edac.Value)
		_bfgb := _badae.NewCell()
		_bfgb.SetHorizontalAlignment(_edac.Alignment)
		_bfgb.SetBackgroundColor(_edac.BackgroundColor)
		_fbcd.setCellBorder(_bfgb, _edac)
		_bfgb.SetContent(_bgcd)
	}
	for _, _cbcc := range _fbcd._cfg {
		for _, _cacgb := range _cbcc {
			_dbg := _fbgb(_cacgb.TextStyle)
			_dbg.SetMargins(0, 0, 3, 2)
			_dbg.Append(_cacgb.Value)
			_gbfe := _badae.NewCell()
			_gbfe.SetHorizontalAlignment(_cacgb.Alignment)
			_gbfe.SetBackgroundColor(_cacgb.BackgroundColor)
			_fbcd.setCellBorder(_gbfe, _cacgb)
			_gbfe.SetContent(_dbg)
		}
	}
	return _badae.GeneratePageBlocks(_ffbad)
}

// Paragraph represents text drawn with a specified font and can wrap across lines and pages.
// By default it occupies the available width in the drawing context.
type Paragraph struct {
	_bfagd         string
	_cgbf          *_g.PdfFont
	_ebbae         float64
	_dfgg          float64
	_gade          Color
	_dgadd         TextAlignment
	_acegd         bool
	_dgee          float64
	_gefg          int
	_dbca          bool
	_cbccg         float64
	_bbdee         Margins
	_fcgd          Positioning
	_fccc          float64
	_ddda          float64
	_afedc, _fcdgg float64
	_egge          []string
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

// MoveRight moves the drawing context right by relative displacement dx (negative goes left).
func (_ceba *Creator) MoveRight(dx float64) { _ceba._defb.X += dx }

// Width returns the current page width.
func (_bagd *Creator) Width() float64 { return _bagd._cgb }

// NewImageFromFile creates an Image from a file.
func (_fadb *Creator) NewImageFromFile(path string) (*Image, error) { return _gedg(path) }

// NewImageFromURL creates an Image from a url.
func (_fadb *Creator) NewImageFromURL(url string) (*Image, error) {
	resp, err := _da.HTTPGet(context.Background(), url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_eegg, _bac := _g.ImageHandling.Read(resp.Body)
	if _bac != nil {
		_da.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bac)
		return nil, _bac
	}
	return _ageb(_eegg)
}

// Width is not used. The list component is designed to fill into the available
// width depending on the context. Returns 0.
func (_cgeaf *List) Width() float64 { return 0 }
func (_edff *Table) updateRowHeights(_acgca float64) {
	for _, _eeaef := range _edff._dcab {
		_cgcg := _eeaef.width(_edff._dcdg, _acgca)
		_eege := _edff._abeg[_eeaef._cbfc+_eeaef._cdcd-2]
		if _ddafb := _eeaef.height(_cgcg); _ddafb > _eege {
			_adab := _ddafb / float64(_eeaef._cdcd)
			for _ddcb := 1; _ddcb <= _eeaef._cdcd; _ddcb++ {
				if _adab > _edff._abeg[_eeaef._cbfc+_ddcb-2] {
					_edff._abeg[_eeaef._cbfc+_ddcb-2] = _adab
				}
			}
		}
	}
}

// GeneratePageBlocks draws the polygon on a new block representing the page.
// Implements the Drawable interface.
func (_aegg *Polygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_aagg := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_bcfgb, _acdbg := _aagg.setOpacity(_aegg._aebd, _aegg._daae)
	if _acdbg != nil {
		return nil, ctx, _acdbg
	}
	_cddfe := _aegg._fdfa
	_cddfe.FillEnabled = _cddfe.FillColor != nil
	_cddfe.BorderEnabled = _cddfe.BorderColor != nil && _cddfe.BorderWidth > 0
	_fdaef := _cddfe.Points
	for _egc := range _fdaef {
		for _dfccd := range _fdaef[_egc] {
			_ggdf := &_fdaef[_egc][_dfccd]
			_ggdf.Y = ctx.PageHeight - _ggdf.Y
		}
	}
	_ccbbf, _, _acdbg := _cddfe.Draw(_bcfgb)
	if _acdbg != nil {
		return nil, ctx, _acdbg
	}
	if _acdbg = _aagg.addContentsByString(string(_ccbbf)); _acdbg != nil {
		return nil, ctx, _acdbg
	}
	return []*Block{_aagg}, ctx, nil
}

// GeneratePageBlocks draws the line on a new block representing the page. Implements the Drawable interface.
func (_fede *Line) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fcdg := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_fffc := _bb.Line{LineWidth: _fede._bbgc, Opacity: 1.0, LineColor: _afag(_fede._cbgb), LineEndingStyle1: _bb.LineEndingStyleNone, LineEndingStyle2: _bb.LineEndingStyleNone, X1: _fede._baba, Y1: ctx.PageHeight - _fede._fbb, X2: _fede._ggadb, Y2: ctx.PageHeight - _fede._bgbc}
	_cdfg, _, _bbgf := _fffc.Draw("")
	if _bbgf != nil {
		return nil, ctx, _bbgf
	}
	_bbgf = _fcdg.addContentsByString(string(_cdfg))
	if _bbgf != nil {
		return nil, ctx, _bbgf
	}
	return []*Block{_fcdg}, ctx, nil
}

// SetContent sets the cell's content.  The content is a VectorDrawable, i.e.
// a Drawable with a known height and width.
// Currently supported VectorDrawables:
// - *Paragraph
// - *StyledParagraph
// - *Image
// - *Table
// - *List
// - *Division
// - *Chart
func (_cccf *TableCell) SetContent(vd VectorDrawable) error {
	switch _fgec := vd.(type) {
	case *Paragraph:
		if _fgec._dbca {
			_fgec._acegd = true
		}
		_cccf._adeed = vd
	case *StyledParagraph:
		if _fgec._deag {
			_fgec._cfga = true
		}
		_cccf._adeed = vd
	case *Image:
		_cccf._adeed = vd
	case *Table:
		_cccf._adeed = vd
	case *List:
		_cccf._adeed = vd
	case *Division:
		_cccf._adeed = vd
	case *Chart:
		_cccf._adeed = vd
	default:
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0079p\u0065\u0020\u0025\u0054", vd)
		return _ec.ErrTypeError
	}
	return nil
}

// WriteToFile writes the Creator output to file specified by path.
func (_gaaf *Creator) WriteToFile(outputPath string) error {
	_bgc, _cagd := _c.Create(outputPath)
	if _cagd != nil {
		return _cagd
	}
	defer _bgc.Close()
	return _gaaf.Write(_bgc)
}

// Width returns the width of the Paragraph.
func (_cbfb *Paragraph) Width() float64 {
	if _cbfb._acegd && int(_cbfb._dgee) > 0 {
		return _cbfb._dgee
	}
	return _cbfb.getTextWidth() / 1000.0
}

// NewCell returns a new invoice table cell.
func (_gbf *Invoice) NewCell(value string) *InvoiceCell {
	return _gbf.newCell(value, _gbf.NewCellProps())
}
func (_aefa *StyledParagraph) wrapText() error { return _aefa.wrapChunks(true) }

// NewBlockFromPage creates a Block from a PDF Page.  Useful for loading template pages as blocks
// from a PDF document and additional content with the creator.
func NewBlockFromPage(page *_g.PdfPage) (*Block, error) {
	_bbd := &Block{}
	_fcc, _fea := page.GetAllContentStreams()
	if _fea != nil {
		return nil, _fea
	}
	_dd := _fc.NewContentStreamParser(_fcc)
	_dcc, _fea := _dd.Parse()
	if _fea != nil {
		return nil, _fea
	}
	_dcc.WrapIfNeeded()
	_bbd._cb = _dcc
	if page.Resources != nil {
		_bbd._fga = page.Resources
	} else {
		_bbd._fga = _g.NewPdfPageResources()
	}
	_ea, _fea := page.GetMediaBox()
	if _fea != nil {
		return nil, _fea
	}
	if _ea.Llx != 0 || _ea.Lly != 0 {
		_bbd.translate(-_ea.Llx, _ea.Lly)
	}
	_bbd._deb = _ea.Urx - _ea.Llx
	_bbd._df = _ea.Ury - _ea.Lly
	if page.Rotate != nil {
		_bbd._fb = -float64(*page.Rotate)
	}
	return _bbd, nil
}

// Draw processes the specified Drawable widget and generates blocks that can
// be rendered to the output document. The generated blocks can span over one
// or more pages. Additional pages are added if the contents go over the current
// page. Each generated block is assigned to the creator page it will be
// rendered to. In order to render the generated blocks to the creator pages,
// call Finalize, Write or WriteToFile.
func (_afab *Creator) Draw(d Drawable) error {
	if _afab.getActivePage() == nil {
		_afab.NewPage()
	}
	_gggba, _cca, _eaa := d.GeneratePageBlocks(_afab._defb)
	if _eaa != nil {
		return _eaa
	}
	if len(_cca._bgdc) > 0 {
		_afab.Errors = append(_afab.Errors, _cca._bgdc...)
	}
	for _gfgg, _gcgc := range _gggba {
		if _gfgg > 0 {
			_afab.NewPage()
		}
		_gfcd := _afab.getActivePage()
		if _baga, _ccee := _afab._fcac[_gfcd]; _ccee {
			if _fcbd := _baga.mergeBlocks(_gcgc); _fcbd != nil {
				return _fcbd
			}
			if _febb := _ggef(_gcgc._fga, _baga._fga); _febb != nil {
				return _febb
			}
		} else {
			_afab._fcac[_gfcd] = _gcgc
		}
	}
	_afab._defb.X = _cca.X
	_afab._defb.Y = _cca.Y
	_afab._defb.Height = _cca.PageHeight - _cca.Y - _cca.Margins.Bottom
	return nil
}
func _cafb(_debd VectorDrawable, _fdafd float64) float64 {
	switch _edefa := _debd.(type) {
	case *Paragraph:
		if _edefa._acegd {
			_edefa.SetWidth(_fdafd)
		}
		return _edefa.Height() + _edefa._bbdee.Top + _edefa._bbdee.Bottom
	case *StyledParagraph:
		if _edefa._cfga {
			_edefa.SetWidth(_fdafd)
		}
		return _edefa.Height() + _edefa._cdbab.Top + _edefa._cdbab.Bottom
	case marginDrawable:
		_, _, _accbb, _bbcf := _edefa.GetMargins()
		return _edefa.Height() + _accbb + _bbcf
	default:
		return _edefa.Height()
	}
}
func (_abca *StyledParagraph) getLineMetrics(_bcafb int) (_bgba, _acfff, _ccbd float64) {
	if _abca._aageb == nil || len(_abca._aageb) == 0 {
		_abca.wrapText()
	}
	if _bcafb < 0 || _bcafb > len(_abca._aageb)-1 {
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020p\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u006c\u0069\u006e\u0065 \u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002e\u0020\u0052\u0065tu\u0072\u006e\u0069\u006e\u0067\u0020\u0030\u002c\u0020\u0030", _bcafb)
		return 0, 0, 0
	}
	_cbcf := _abca._aageb[_bcafb]
	for _, _fcdf := range _cbcf {
		_edea, _gdacd := _fcdf.Style.Font.GetFontDescriptor()
		if _gdacd != nil {
			_da.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
		}
		var _dgffa, _cbcce float64
		if _edea != nil {
			if _dgffa, _gdacd = _edea.GetCapHeight(); _gdacd != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _gdacd)
			}
			if _cbcce, _gdacd = _edea.GetDescent(); _gdacd != nil {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020U\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0067\u0065t\u0020\u0066\u006f\u006e\u0074\u0020\u0044\u0065\u0073\u0063\u0065\u006et\u003a\u0020\u0025\u0076", _gdacd)
			}
		}
		if int(_dgffa) <= 0 {
			_da.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
			_dgffa = 1000
		}
		if _baaac := _dgffa / 1000.0 * _fcdf.Style.FontSize; _baaac > _bgba {
			_bgba = _baaac
		}
		if _gcca := _fcdf.Style.FontSize; _gcca > _acfff {
			_acfff = _gcca
		}
		if _gdcd := _cbcce / 1000.0 * _fcdf.Style.FontSize; _gdcd < _ccbd {
			_ccbd = _gdcd
		}
	}
	return _bgba, _acfff, _ccbd
}
func (_bgff cmykColor) ToRGB() (float64, float64, float64) {
	_aef := _bgff._abf
	return 1 - (_bgff._debb*(1-_aef) + _aef), 1 - (_bgff._caad*(1-_aef) + _aef), 1 - (_bgff._gfbg*(1-_aef) + _aef)
}

var PPI float64 = 72

// AppendCurve appends a Bezier curve to the filled curve.
func (_gabg *FilledCurve) AppendCurve(curve _bb.CubicBezierCurve) *FilledCurve {
	_gabg._dgfb = append(_gabg._dgfb, curve)
	return _gabg
}
func (_acdbgc *TableCell) width(_bfddd []float64, _gbca float64) float64 {
	_aeef := float64(0.0)
	for _bdad := 0; _bdad < _acdbgc._bdbg; _bdad++ {
		_aeef += _bfddd[_acdbgc._gbdde+_bdad-1]
	}
	return _aeef * _gbca
}

// CellBorderStyle defines the table cell's border style.
type CellBorderStyle int

// NewColumn returns a new column for the line items invoice table.
func (_daag *Invoice) NewColumn(description string) *InvoiceCell {
	return _daag.newColumn(description, CellHorizontalAlignmentLeft)
}

// AddLine appends a new line to the invoice line items table.
func (_decfb *Invoice) AddLine(values ...string) []*InvoiceCell {
	_eceg := len(_decfb._cfc)
	var _abced []*InvoiceCell
	for _eccc, _gaae := range values {
		_baac := _decfb.newCell(_gaae, _decfb._cbcee)
		if _eccc < _eceg {
			_baac.Alignment = _decfb._cfc[_eccc].Alignment
		}
		_abced = append(_abced, _baac)
	}
	_decfb._cfg = append(_decfb._cfg, _abced)
	return _abced
}

// SetAddressHeadingStyle sets the style properties used to render the
// heading of the invoice address sections.
func (_cacg *Invoice) SetAddressHeadingStyle(style TextStyle) { _cacg._deef = style }

// Lines returns all the rows of the invoice line items table.
func (_dcfb *Invoice) Lines() [][]*InvoiceCell { return _dcfb._cfg }

// SetBorderColor sets the border color.
func (_ffba *Ellipse) SetBorderColor(col Color) { _ffba._gbad = col }

// NewFilledCurve returns a instance of filled curve.
func (_eggdd *Creator) NewFilledCurve() *FilledCurve { return _eabf() }

// Title returns the title of the invoice.
func (_cbeca *Invoice) Title() string { return _cbeca._cbde }

// Chapter is used to arrange multiple drawables (paragraphs, images, etc) into a single section.
// The concept is the same as a book or a report chapter.
type Chapter struct {
	_cdb         int
	_afaa        string
	_gacfc       *Paragraph
	_aaga        []Drawable
	_fcd         int
	_aec         bool
	_cdd         bool
	_bgd         Positioning
	_gacad, _caf float64
	_acge        Margins
	_dbbg        *Chapter
	_gecdd       *TOC
	_gfag        *_g.Outline
	_bee         *_g.OutlineItem
	_cef         uint
}

// SetNoteHeadingStyle sets the style properties used to render the heading
// of the invoice note sections.
func (_gabb *Invoice) SetNoteHeadingStyle(style TextStyle) { _gabb._dedb = style }

// SetTextOverflow controls the behavior of paragraph text which
// does not fit in the available space.
func (_abbb *StyledParagraph) SetTextOverflow(textOverflow TextOverflow) { _abbb._cfgc = textOverflow }
func (_bceg *Invoice) drawSection(_dabc, _aggf string) []*StyledParagraph {
	var _fbed []*StyledParagraph
	if _dabc != "" {
		_cega := _fbgb(_bceg._dedb)
		_cega.SetMargins(0, 0, 0, 5)
		_cega.Append(_dabc)
		_fbed = append(_fbed, _cega)
	}
	if _aggf != "" {
		_feca := _fbgb(_bceg._efec)
		_feca.Append(_aggf)
		_fbed = append(_fbed, _feca)
	}
	return _fbed
}

// SetHeading sets the text and the style of the heading of the TOC component.
func (_ffbf *TOC) SetHeading(text string, style TextStyle) {
	_caegc := _ffbf.Heading()
	_caegc.Reset()
	_fbagg := _caegc.Append(text)
	_fbagg.Style = style
}

// TOCLine represents a line in a table of contents.
// The component can be used both in the context of a
// table of contents component and as a standalone component.
// The representation of a table of contents line is as follows:
//       [number] [title]      [separator] [page]
// e.g.: Chapter1 Introduction ........... 1
type TOCLine struct {
	_daedd *StyledParagraph

	// Holds the text and style of the number part of the TOC line.
	Number TextChunk

	// Holds the text and style of the title part of the TOC line.
	Title TextChunk

	// Holds the text and style of the separator part of the TOC line.
	Separator TextChunk

	// Holds the text and style of the page part of the TOC line.
	Page   TextChunk
	_egde  float64
	_cefb  uint
	_dedcg float64
	_cada  Positioning
	_gddbe float64
	_cdde  float64
	_eegc  int64
}

// SetLineWidth sets the line width.
func (_ceef *Polyline) SetLineWidth(lineWidth float64) { _ceef._efcf.LineWidth = lineWidth }

// SetText replaces all the text of the paragraph with the specified one.
func (_fgdc *StyledParagraph) SetText(text string) *TextChunk {
	_fgdc.Reset()
	return _fgdc.Append(text)
}

// ScaleToWidth scale Image to a specified width w, maintaining the aspect ratio.
func (_cged *Image) ScaleToWidth(w float64) {
	_cgaga := _cged._gcab / _cged._ebb
	_cged._ebb = w
	_cged._gcab = w * _cgaga
}

const (
	TextAlignmentLeft TextAlignment = iota
	TextAlignmentRight
	TextAlignmentCenter
	TextAlignmentJustify
)

// SetText sets the text content of the Paragraph.
func (_cadga *Paragraph) SetText(text string)            { _cadga._bfagd = text }
func (_ceg rgbColor) ToRGB() (float64, float64, float64) { return _ceg._aacc, _ceg._beeb, _ceg._fgea }

// NewList creates a new list.
func (_fgf *Creator) NewList() *List { return _dfgb(_fgf.NewTextStyle()) }
func _dgcda(_deaff, _fedb, _bcfe string, _gfffd uint, _bced TextStyle) *TOCLine {
	return _ffbgf(TextChunk{Text: _deaff, Style: _bced}, TextChunk{Text: _fedb, Style: _bced}, TextChunk{Text: _bcfe, Style: _bced}, _gfffd, _bced)
}
func (_eadb *StyledParagraph) appendChunk(_dfdf *TextChunk) *TextChunk {
	_eadb._dgff = append(_eadb._dgff, _dfdf)
	_eadb.wrapText()
	return _dfdf
}

// SetNumber sets the number of the invoice.
func (_ebeb *Invoice) SetNumber(number string) (*InvoiceCell, *InvoiceCell) {
	_ebeb._aafe[1].Value = number
	return _ebeb._aafe[0], _ebeb._aafe[1]
}

// String implements error interface.
func (_afed UnsupportedRuneError) Error() string { return _afed.Message }
func (_daac *Table) moveToNextAvailableCell() int {
	_cebg := (_daac._daab-1)%(_daac._fabb) + 1
	for {
		if _cebg-1 >= len(_daac._aegce) {
			return _cebg
		} else if _daac._aegce[_cebg-1] == 0 {
			return _cebg
		} else {
			_daac._daab++
			_daac._aegce[_cebg-1]--
		}
		_cebg++
	}
}

// TOC represents a table of contents component.
// It consists of a paragraph heading and a collection of
// table of contents lines.
// The representation of a table of contents line is as follows:
//       [number] [title]      [separator] [page]
// e.g.: Chapter1 Introduction ........... 1
type TOC struct {
	_ebgb  *StyledParagraph
	_dbcgf []*TOCLine
	_ceaf  TextStyle
	_dade  TextStyle
	_afca  TextStyle
	_feaff TextStyle
	_acfe  string
	_eecca float64
	_efbe  Margins
	_adeee Positioning
	_dcffa TextStyle
	_bfec  bool
}

// GeneratePageBlocks generates a page break block.
func (_fcca *PageBreak) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_egf := []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y), NewBlock(ctx.PageWidth, ctx.PageHeight)}
	ctx.Page++
	_cddg := ctx
	_cddg.Y = ctx.Margins.Top
	_cddg.X = ctx.Margins.Left
	_cddg.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
	_cddg.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right
	ctx = _cddg
	return _egf, ctx, nil
}
func _bggg(_fceg, _daa, _caeg, _acfa, _faff, _ecca float64) *Curve {
	_adcg := &Curve{}
	_adcg._aacg = _fceg
	_adcg._eagd = _daa
	_adcg._aade = _caeg
	_adcg._cbec = _acfa
	_adcg._cceec = _faff
	_adcg._cdbf = _ecca
	_adcg._eeb = ColorBlack
	_adcg._faca = 1.0
	return _adcg
}

// CreateFrontPage sets a function to generate a front Page.
func (_acbc *Creator) CreateFrontPage(genFrontPageFunc func(_agbd FrontpageFunctionArgs)) {
	_acbc._gace = genFrontPageFunc
}
func (_edgde *Table) resetColumnWidths() {
	_edgde._dcdg = []float64{}
	_dfgbd := float64(1.0) / float64(_edgde._fabb)
	for _aff := 0; _aff < _edgde._fabb; _aff++ {
		_edgde._dcdg = append(_edgde._dcdg, _dfgbd)
	}
}

// NewImageFromData creates an Image from image data.
func (_fegg *Creator) NewImageFromData(data []byte) (*Image, error) { return _aacce(data) }
func _ebfbc(_gced string, _agd TextStyle) *Paragraph {
	_ebd := &Paragraph{_bfagd: _gced, _cgbf: _agd.Font, _ebbae: _agd.FontSize, _dfgg: 1.0, _acegd: true, _dbca: true, _dgadd: TextAlignmentLeft, _cbccg: 0, _afedc: 1, _fcdgg: 1, _fcgd: PositionRelative}
	_ebd.SetColor(_agd.Color)
	return _ebd
}

// AddTotalLine adds a new line in the invoice totals table.
func (_adege *Invoice) AddTotalLine(desc, value string) (*InvoiceCell, *InvoiceCell) {
	_fgdf := &InvoiceCell{_adege._ddcd, desc}
	_cdbg := &InvoiceCell{_adege._ddcd, value}
	_adege._gegc = append(_adege._gegc, [2]*InvoiceCell{_fgdf, _cdbg})
	return _fgdf, _cdbg
}

// Add adds a new line with the default style to the table of contents.
func (_baae *TOC) Add(number, title, page string, level uint) *TOCLine {
	_adfag := _baae.AddLine(_ffbgf(TextChunk{Text: number, Style: _baae._ceaf}, TextChunk{Text: title, Style: _baae._dade}, TextChunk{Text: page, Style: _baae._feaff}, level, _baae._dcffa))
	if _adfag == nil {
		return nil
	}
	_bcbab := &_baae._efbe
	_adfag.SetMargins(_bcbab.Left, _bcbab.Right, _bcbab.Top, _bcbab.Bottom)
	_adfag.SetLevelOffset(_baae._eecca)
	_adfag.Separator.Text = _baae._acfe
	_adfag.Separator.Style = _baae._afca
	return _adfag
}
func _gbg(_eaff *Chapter, _facc *TOC, _bbg *_g.Outline, _ffda string, _ebcf int, _ffde TextStyle) *Chapter {
	var _bad uint = 1
	if _eaff != nil {
		_bad = _eaff._cef + 1
	}
	_cdad := &Chapter{_cdb: _ebcf, _afaa: _ffda, _aec: true, _cdd: true, _dbbg: _eaff, _gecdd: _facc, _gfag: _bbg, _aaga: []Drawable{}, _cef: _bad}
	_ecab := _ebfbc(_cdad.headingText(), _ffde)
	_ecab.SetFont(_ffde.Font)
	_ecab.SetFontSize(_ffde.FontSize)
	_cdad._gacfc = _ecab
	return _cdad
}

// GeneratePageBlocks implements drawable interface.
func (_dgcd *border) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cce := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_dgb := _dgcd._edd
	_fdg := ctx.PageHeight - _dgcd._ccga
	if _dgcd._edee != nil {
		_eca := _bb.Rectangle{Opacity: 1.0, X: _dgcd._edd, Y: ctx.PageHeight - _dgcd._ccga - _dgcd._cfb, Height: _dgcd._cfb, Width: _dgcd._debf}
		_eca.FillEnabled = true
		_eca.FillColor = _afag(_dgcd._edee)
		_eca.BorderEnabled = false
		_bfb, _, _gfa := _eca.Draw("")
		if _gfa != nil {
			return nil, ctx, _gfa
		}
		_gfa = _cce.addContentsByString(string(_bfb))
		if _gfa != nil {
			return nil, ctx, _gfa
		}
	}
	_ebf := _dgcd._bdf
	_baed := _dgcd._gfg
	_geb := _dgcd._age
	_aacd := _dgcd._cedc
	_gfce := _dgcd._bdf
	if _dgcd._ddc == CellBorderStyleDouble {
		_gfce += 2 * _ebf
	}
	_bbc := _dgcd._gfg
	if _dgcd._bedg == CellBorderStyleDouble {
		_bbc += 2 * _baed
	}
	_ccgc := _dgcd._age
	if _dgcd._gda == CellBorderStyleDouble {
		_ccgc += 2 * _geb
	}
	_bdfc := _dgcd._cedc
	if _dgcd._gga == CellBorderStyleDouble {
		_bdfc += 2 * _aacd
	}
	_fag := (_gfce - _ccgc) / 2
	_fagd := (_gfce - _bdfc) / 2
	_gfd := (_bbc - _ccgc) / 2
	_dbb := (_bbc - _bdfc) / 2
	if _dgcd._bdf != 0 {
		_fdad := _dgb
		_aag := _fdg
		if _dgcd._ddc == CellBorderStyleDouble {
			_aag -= _ebf
			_afc := _bb.BasicLine{LineColor: _afag(_dgcd._acg), Opacity: 1.0, LineWidth: _dgcd._bdf, LineStyle: _dgcd.LineStyle, X1: _fdad - _gfce/2 + _fag, Y1: _aag + 2*_ebf, X2: _fdad + _gfce/2 - _fagd + _dgcd._debf, Y2: _aag + 2*_ebf}
			_bbcg, _, _aacb := _afc.Draw("")
			if _aacb != nil {
				return nil, ctx, _aacb
			}
			_aacb = _cce.addContentsByString(string(_bbcg))
			if _aacb != nil {
				return nil, ctx, _aacb
			}
		}
		_dfd := _bb.BasicLine{LineWidth: _dgcd._bdf, Opacity: 1.0, LineColor: _afag(_dgcd._acg), LineStyle: _dgcd.LineStyle, X1: _fdad - _gfce/2 + _fag + (_ccgc - _dgcd._age), Y1: _aag, X2: _fdad + _gfce/2 - _fagd + _dgcd._debf - (_bdfc - _dgcd._cedc), Y2: _aag}
		_febf, _, _dbed := _dfd.Draw("")
		if _dbed != nil {
			return nil, ctx, _dbed
		}
		_dbed = _cce.addContentsByString(string(_febf))
		if _dbed != nil {
			return nil, ctx, _dbed
		}
	}
	if _dgcd._gfg != 0 {
		_acbd := _dgb
		_gfad := _fdg - _dgcd._cfb
		if _dgcd._bedg == CellBorderStyleDouble {
			_gfad += _baed
			_bcb := _bb.BasicLine{LineWidth: _dgcd._gfg, Opacity: 1.0, LineColor: _afag(_dgcd._aad), LineStyle: _dgcd.LineStyle, X1: _acbd - _bbc/2 + _gfd, Y1: _gfad - 2*_baed, X2: _acbd + _bbc/2 - _dbb + _dgcd._debf, Y2: _gfad - 2*_baed}
			_dagg, _, _fdae := _bcb.Draw("")
			if _fdae != nil {
				return nil, ctx, _fdae
			}
			_fdae = _cce.addContentsByString(string(_dagg))
			if _fdae != nil {
				return nil, ctx, _fdae
			}
		}
		_adcd := _bb.BasicLine{LineWidth: _dgcd._gfg, Opacity: 1.0, LineColor: _afag(_dgcd._aad), LineStyle: _dgcd.LineStyle, X1: _acbd - _bbc/2 + _gfd + (_ccgc - _dgcd._age), Y1: _gfad, X2: _acbd + _bbc/2 - _dbb + _dgcd._debf - (_bdfc - _dgcd._cedc), Y2: _gfad}
		_acd, _, _fbfg := _adcd.Draw("")
		if _fbfg != nil {
			return nil, ctx, _fbfg
		}
		_fbfg = _cce.addContentsByString(string(_acd))
		if _fbfg != nil {
			return nil, ctx, _fbfg
		}
	}
	if _dgcd._age != 0 {
		_bbdb := _dgb
		_eedb := _fdg
		if _dgcd._gda == CellBorderStyleDouble {
			_bbdb += _geb
			_ebfb := _bb.BasicLine{LineWidth: _dgcd._age, Opacity: 1.0, LineColor: _afag(_dgcd._gcbe), LineStyle: _dgcd.LineStyle, X1: _bbdb - 2*_geb, Y1: _eedb + _ccgc/2 + _fag, X2: _bbdb - 2*_geb, Y2: _eedb - _ccgc/2 - _gfd - _dgcd._cfb}
			_aaed, _, _aggg := _ebfb.Draw("")
			if _aggg != nil {
				return nil, ctx, _aggg
			}
			_aggg = _cce.addContentsByString(string(_aaed))
			if _aggg != nil {
				return nil, ctx, _aggg
			}
		}
		_bea := _bb.BasicLine{LineWidth: _dgcd._age, Opacity: 1.0, LineColor: _afag(_dgcd._gcbe), LineStyle: _dgcd.LineStyle, X1: _bbdb, Y1: _eedb + _ccgc/2 + _fag - (_gfce - _dgcd._bdf), X2: _bbdb, Y2: _eedb - _ccgc/2 - _gfd - _dgcd._cfb + (_bbc - _dgcd._gfg)}
		_edg, _, _gdf := _bea.Draw("")
		if _gdf != nil {
			return nil, ctx, _gdf
		}
		_gdf = _cce.addContentsByString(string(_edg))
		if _gdf != nil {
			return nil, ctx, _gdf
		}
	}
	if _dgcd._cedc != 0 {
		_cec := _dgb + _dgcd._debf
		_fad := _fdg
		if _dgcd._gga == CellBorderStyleDouble {
			_cec -= _aacd
			_gcbg := _bb.BasicLine{LineWidth: _dgcd._cedc, Opacity: 1.0, LineColor: _afag(_dgcd._edf), LineStyle: _dgcd.LineStyle, X1: _cec + 2*_aacd, Y1: _fad + _bdfc/2 + _fagd, X2: _cec + 2*_aacd, Y2: _fad - _bdfc/2 - _dbb - _dgcd._cfb}
			_aceg, _, _fbg := _gcbg.Draw("")
			if _fbg != nil {
				return nil, ctx, _fbg
			}
			_fbg = _cce.addContentsByString(string(_aceg))
			if _fbg != nil {
				return nil, ctx, _fbg
			}
		}
		_dagb := _bb.BasicLine{LineWidth: _dgcd._cedc, Opacity: 1.0, LineColor: _afag(_dgcd._edf), LineStyle: _dgcd.LineStyle, X1: _cec, Y1: _fad + _bdfc/2 + _fagd - (_gfce - _dgcd._bdf), X2: _cec, Y2: _fad - _bdfc/2 - _dbb - _dgcd._cfb + (_bbc - _dgcd._gfg)}
		_fac, _, _gggb := _dagb.Draw("")
		if _gggb != nil {
			return nil, ctx, _gggb
		}
		_gggb = _cce.addContentsByString(string(_fac))
		if _gggb != nil {
			return nil, ctx, _gggb
		}
	}
	return []*Block{_cce}, ctx, nil
}

// SetLineColor sets the line color.
func (_bfdc *Polyline) SetLineColor(color Color) { _bfdc._efcf.LineColor = _afag(color) }

// GeneratePageBlocks generate the Page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages.
func (_fdgc *Chapter) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_egg := ctx
	if _fdgc._bgd.IsRelative() {
		ctx.X += _fdgc._acge.Left
		ctx.Y += _fdgc._acge.Top
		ctx.Width -= _fdgc._acge.Left + _fdgc._acge.Right
		ctx.Height -= _fdgc._acge.Top
	}
	_bdgc, _fcfd, _abgb := _fdgc._gacfc.GeneratePageBlocks(ctx)
	if _abgb != nil {
		return _bdgc, ctx, _abgb
	}
	ctx = _fcfd
	_ade := ctx.X
	_gaef := ctx.Y - _fdgc._gacfc.Height()
	_aaf := int64(ctx.Page)
	_deeg := _fdgc.headingNumber()
	_fbc := _fdgc.headingText()
	if _fdgc._cdd {
		_ecgc := _fdgc._gecdd.Add(_deeg, _fdgc._afaa, _cf.FormatInt(_aaf, 10), _fdgc._cef)
		if _fdgc._gecdd._bfec {
			_ecgc.SetLink(_aaf, _ade, _gaef)
		}
	}
	if _fdgc._bee == nil {
		_fdgc._bee = _g.NewOutlineItem(_fbc, _g.NewOutlineDest(_aaf-1, _ade, _gaef))
		if _fdgc._dbbg != nil {
			_fdgc._dbbg._bee.Add(_fdgc._bee)
		} else {
			_fdgc._gfag.Add(_fdgc._bee)
		}
	} else {
		_edb := &_fdgc._bee.Dest
		_edb.Page = _aaf - 1
		_edb.X = _ade
		_edb.Y = _gaef
	}
	for _, _add := range _fdgc._aaga {
		_bcgb, _efa, _dca := _add.GeneratePageBlocks(ctx)
		if _dca != nil {
			return _bdgc, ctx, _dca
		}
		if len(_bcgb) < 1 {
			continue
		}
		_bdgc[len(_bdgc)-1].mergeBlocks(_bcgb[0])
		_bdgc = append(_bdgc, _bcgb[1:]...)
		ctx = _efa
	}
	if _fdgc._bgd.IsRelative() {
		ctx.X = _egg.X
	}
	if _fdgc._bgd.IsAbsolute() {
		return _bdgc, _egg, nil
	}
	return _bdgc, ctx, nil
}
func _dfcea(_bdebf *_c.File) ([]*_g.PdfPage, error) {
	_debg, _geac := _g.NewPdfReader(_bdebf)
	if _geac != nil {
		return nil, _geac
	}
	_fddc, _geac := _debg.GetNumPages()
	if _geac != nil {
		return nil, _geac
	}
	var _gbgaf []*_g.PdfPage
	for _bfdgf := 0; _bfdgf < _fddc; _bfdgf++ {
		_egdb, _cgbfg := _debg.GetPage(_bfdgf + 1)
		if _cgbfg != nil {
			return nil, _cgbfg
		}
		_gbgaf = append(_gbgaf, _egdb)
	}
	return _gbgaf, nil
}
func _beebb(_cfdb _fg.Image) (*Image, error) {
	_deda, _gcc := _g.ImageHandling.NewImageFromGoImage(_cfdb)
	if _gcc != nil {
		return nil, _gcc
	}
	return _ageb(_deda)
}

// CurvePolygon represents a curve polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type CurvePolygon struct {
	_bafd *_bb.CurvePolygon
	_accd float64
	_dfaa float64
}

// DueDate returns the invoice due date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_ffge *Invoice) DueDate() (*InvoiceCell, *InvoiceCell) { return _ffge._gdge[0], _ffge._gdge[1] }

// GetCoords returns the coordinates of the Ellipse's center (xc,yc).
func (_eeff *Ellipse) GetCoords() (float64, float64) { return _eeff._gffd, _eeff._cgae }
func (_agbb *StyledParagraph) wrapChunks(_ffcec bool) error {
	if !_agbb._cfga || int(_agbb._bggd) <= 0 {
		_agbb._aageb = [][]*TextChunk{_agbb._dgff}
		return nil
	}
	_agbb._aageb = [][]*TextChunk{}
	var _dfbde []*TextChunk
	var _abda float64
	_cbaa := _a.IsSpace
	if !_ffcec {
		_cbaa = func(rune) bool { return false }
	}
	_aace := _dbba(_agbb._bggd*1000.0, 0.000001)
	for _, _ceae := range _agbb._dgff {
		_babe := _ceae.Style
		_bgfd := _ceae._adbcg
		var (
			_gggc  []rune
			_aaccf []float64
		)
		for _, _fbab := range _ceae.Text {
			if _fbab == '\u000A' {
				if !_ffcec {
					_gggc = append(_gggc, _fbab)
				}
				_dfbde = append(_dfbde, &TextChunk{Text: _db.TrimRightFunc(string(_gggc), _cbaa), Style: _babe, _adbcg: _fffb(_bgfd)})
				_agbb._aageb = append(_agbb._aageb, _dfbde)
				_dfbde = nil
				_abda = 0
				_gggc = nil
				_aaccf = nil
				continue
			}
			_aaeaf := _fbab == ' '
			_bdaa, _gcaf := _babe.Font.GetRuneMetrics(_fbab)
			if !_gcaf {
				_da.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _fbab)
				return _f.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
			}
			_ffcg := _babe.FontSize * _bdaa.Wx * _babe.horizontalScale()
			_gddb := _ffcg
			if !_aaeaf {
				_gddb = _ffcg + _babe.CharSpacing*1000.0
			}
			if _abda+_ffcg > _aace {
				_bbce := -1
				if !_aaeaf {
					for _faae := len(_gggc) - 1; _faae >= 0; _faae-- {
						if _gggc[_faae] == ' ' {
							_bbce = _faae
							break
						}
					}
				}
				_egdc := string(_gggc)
				if _bbce >= 0 {
					_egdc = string(_gggc[0 : _bbce+1])
					_gggc = _gggc[_bbce+1:]
					_gggc = append(_gggc, _fbab)
					_aaccf = _aaccf[_bbce+1:]
					_aaccf = append(_aaccf, _gddb)
					_abda = 0
					for _, _abdbb := range _aaccf {
						_abda += _abdbb
					}
				} else {
					if _aaeaf {
						_abda = 0
						_gggc = []rune{}
						_aaccf = []float64{}
					} else {
						_abda = _gddb
						_gggc = []rune{_fbab}
						_aaccf = []float64{_gddb}
					}
				}
				if !_ffcec && _aaeaf {
					_egdc += "\u0020"
				}
				_dfbde = append(_dfbde, &TextChunk{Text: _db.TrimRightFunc(_egdc, _cbaa), Style: _babe, _adbcg: _fffb(_bgfd)})
				_agbb._aageb = append(_agbb._aageb, _dfbde)
				_dfbde = []*TextChunk{}
			} else {
				_abda += _gddb
				_gggc = append(_gggc, _fbab)
				_aaccf = append(_aaccf, _gddb)
			}
		}
		if len(_gggc) > 0 {
			_dfbde = append(_dfbde, &TextChunk{Text: string(_gggc), Style: _babe, _adbcg: _fffb(_bgfd)})
		}
	}
	if len(_dfbde) > 0 {
		_agbb._aageb = append(_agbb._aageb, _dfbde)
	}
	return nil
}
func (_dcda *Table) wrapRow(_cggdf int, _begce DrawContext, _cagf float64) (bool, error) {
	if !_dcda._adce {
		return false, nil
	}
	var (
		_fefce = _dcda._dcab[_cggdf]
		_dfagg = -1
		_cdgad []*TableCell
		_ecbfc float64
		_baacg bool
		_gagf  = make([]float64, 0, len(_dcda._dcdg))
	)
	_ccaeb := func(_aebde *TableCell, _fdbg VectorDrawable, _bggbb bool) *TableCell {
		_gcec := *_aebde
		_gcec._adeed = _fdbg
		if _bggbb {
			_gcec._cbfc++
		}
		return &_gcec
	}
	_beef := func(_agdc int, _abff VectorDrawable) {
		var _edgda float64 = -1
		if _abff == nil {
			if _fdfe := _gagf[_agdc-_cggdf]; _fdfe > _begce.Height {
				_abff = _dcda._dcab[_agdc]._adeed
				_dcda._dcab[_agdc]._adeed = nil
				_gagf[_agdc-_cggdf] = 0
				_edgda = _fdfe
			}
		}
		_bgfg := _ccaeb(_dcda._dcab[_agdc], _abff, true)
		_cdgad = append(_cdgad, _bgfg)
		if _edgda < 0 {
			_edgda = _bgfg.height(_begce.Width)
		}
		if _edgda > _ecbfc {
			_ecbfc = _edgda
		}
	}
	for _acea := _cggdf; _acea < len(_dcda._dcab); _acea++ {
		_ffgec := _dcda._dcab[_acea]
		if _fefce._cbfc != _ffgec._cbfc {
			_dfagg = _acea
			break
		}
		_begce.Width = _ffgec.width(_dcda._dcdg, _cagf)
		var _dbcb VectorDrawable
		switch _fffd := _ffgec._adeed.(type) {
		case *StyledParagraph:
			if _ffeg := _ffgec.height(_begce.Width); _ffeg > _begce.Height {
				_dcde := _begce
				_dcde.Height = _dg.Floor(_begce.Height - _fffd._cdbab.Top - _fffd._cdbab.Bottom - 0.5*_fffd.getTextHeight())
				_afbe, _fbcgc, _fbgag := _fffd.split(_dcde)
				if _fbgag != nil {
					return false, _fbgag
				}
				if _afbe != nil && _fbcgc != nil {
					_fffd = _afbe
					_ffgec = _ccaeb(_ffgec, _afbe, false)
					_dcda._dcab[_acea] = _ffgec
					_dbcb = _fbcgc
					_baacg = true
				}
			}
		case *Division:
			if _cefc := _ffgec.height(_begce.Width); _cefc > _begce.Height {
				_bfddf := _begce
				_bfddf.Height = _dg.Floor(_begce.Height - _fffd._abdb.Top - _fffd._abdb.Bottom)
				_bege, _beae := _fffd.split(_bfddf)
				if _bege != nil && _beae != nil {
					_fffd = _bege
					_ffgec = _ccaeb(_ffgec, _bege, false)
					_dcda._dcab[_acea] = _ffgec
					_dbcb = _beae
					_baacg = true
				}
			}
		}
		_gagf = append(_gagf, _ffgec.height(_begce.Width))
		if _baacg {
			if _cdgad == nil {
				_cdgad = make([]*TableCell, 0, len(_dcda._dcdg))
				for _dcbef := _cggdf; _dcbef < _acea; _dcbef++ {
					_beef(_dcbef, nil)
				}
			}
			_beef(_acea, _dbcb)
		}
	}
	var _fgcc float64
	for _, _gfga := range _gagf {
		if _gfga > _fgcc {
			_fgcc = _gfga
		}
	}
	if _baacg && _fgcc < _begce.Height {
		if _dfagg < 0 {
			_dfagg = len(_dcda._dcab)
		}
		_cdeg := _dcda._dcab[_dfagg-1]._cbfc + _dcda._dcab[_dfagg-1]._cdcd - 1
		for _cafea := _dfagg; _cafea < len(_dcda._dcab); _cafea++ {
			_dcda._dcab[_cafea]._cbfc++
		}
		_dcda._dcab = append(_dcda._dcab[:_dfagg], append(_cdgad, _dcda._dcab[_dfagg:]...)...)
		_dcda._abeg = append(_dcda._abeg[:_cdeg], append([]float64{_ecbfc}, _dcda._abeg[_cdeg:]...)...)
		_dcda._abeg[_fefce._cbfc+_fefce._cdcd-2] = _fgcc
	}
	return _baacg, nil
}

// SetFillColor sets background color for border.
func (_dagd *border) SetFillColor(col Color) { _dagd._edee = col }

// Width returns Image's document width.
func (_fcag *Image) Width() float64 { return _fcag._ebb }

// DrawFooter sets a function to draw a footer on created output pages.
func (_ebg *Creator) DrawFooter(drawFooterFunc func(_agce *Block, _dga FooterFunctionArgs)) {
	_ebg._bfag = drawFooterFunc
}
func _afb(_ddfc string, _feeb _ec.PdfObject, _bdeb *_g.PdfPageResources) _ec.PdfObjectName {
	_fgad := _db.TrimRightFunc(_db.TrimSpace(_ddfc), func(_bfef rune) bool { return _a.IsNumber(_bfef) })
	if _fgad == "" {
		_fgad = "\u0046\u006f\u006e\u0074"
	}
	_fff := 0
	_bdae := _ec.PdfObjectName(_ddfc)
	for {
		_fbf, _ecge := _bdeb.GetFontByName(_bdae)
		if !_ecge || _fbf == _feeb {
			break
		}
		_fff++
		_bdae = _ec.PdfObjectName(_ad.Sprintf("\u0025\u0073\u0025\u0064", _fgad, _fff))
	}
	return _bdae
}

// Lines returns all the lines the table of contents has.
func (_baaag *TOC) Lines() []*TOCLine { return _baaag._dbcgf }

// InvoiceCell represents any cell belonging to a table from the invoice
// template. The main tables are the invoice information table, the line
// items table and totals table. Contains the text value of the cell and
// the style properties of the cell.
type InvoiceCell struct {
	InvoiceCellProps
	Value string
}

// BuyerAddress returns the buyer address used in the invoice template.
func (_cdfbb *Invoice) BuyerAddress() *InvoiceAddress { return _cdfbb._dggaa }
func (_gac *Block) drawToPage(_fcee *_g.PdfPage) error {
	_gaea := &_fc.ContentStreamOperations{}
	if _fcee.Resources == nil {
		_fcee.Resources = _g.NewPdfPageResources()
	}
	_fcf := _cfe(_gaea, _fcee.Resources, _gac._cb, _gac._fga)
	if _fcf != nil {
		return _fcf
	}
	if _fcf = _ggef(_gac._fga, _fcee.Resources); _fcf != nil {
		return _fcf
	}
	if _fcf = _fcee.AppendContentBytes(_gaea.Bytes(), true); _fcf != nil {
		return _fcf
	}
	for _, _aee := range _gac._ggg {
		_fcee.AddAnnotation(_aee)
	}
	return nil
}

// GetMargins returns the Chapter's margin: left, right, top, bottom.
func (_eea *Chapter) GetMargins() (float64, float64, float64, float64) {
	return _eea._acge.Left, _eea._acge.Right, _eea._acge.Top, _eea._acge.Bottom
}

type containerDrawable interface {
	Drawable

	// ContainerComponent checks if the component is allowed to be added into provided 'container' and returns
	// preprocessed copy of itself. If the component is not changed it is allowed to return itself in a callback way.
	// If the component is not compatible with provided container this method should return an error.
	ContainerComponent(_acaa Drawable) (Drawable, error)
}

// SetBorderLineStyle sets border style (currently dashed or plain).
func (_cege *TableCell) SetBorderLineStyle(style _bb.LineStyle) { _cege._egba = style }

// SetLineSeparator sets the separator for all new lines of the table of contents.
func (_fcbcc *TOC) SetLineSeparator(separator string) { _fcbcc._acfe = separator }
func (_cdf *Block) mergeBlocks(_bfa *Block) error {
	_efd := _cfe(_cdf._cb, _cdf._fga, _bfa._cb, _bfa._fga)
	if _efd != nil {
		return _efd
	}
	for _, _eage := range _bfa._ggg {
		_cdf.AddAnnotation(_eage)
	}
	return nil
}

// SkipOver skips over a specified number of rows and cols.
func (_abadg *Table) SkipOver(rows, cols int) {
	_egca := rows*_abadg._fabb + cols - 1
	if _egca < 0 {
		_da.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	_abadg._daab += _egca
}
func _bedaf(_ddgg string, _aabd, _gcfg TextStyle) *TOC {
	_edcfe := _gcfg
	_edcfe.FontSize = 14
	_bedae := _fbgb(_edcfe)
	_bedae.SetEnableWrap(true)
	_bedae.SetTextAlignment(TextAlignmentLeft)
	_bedae.SetMargins(0, 0, 0, 5)
	_fafde := _bedae.Append(_ddgg)
	_fafde.Style = _edcfe
	return &TOC{_ebgb: _bedae, _dbcgf: []*TOCLine{}, _ceaf: _aabd, _dade: _aabd, _afca: _aabd, _feaff: _aabd, _acfe: "\u002e", _eecca: 10, _efbe: Margins{0, 0, 2, 2}, _adeee: PositionRelative, _dcffa: _aabd, _bfec: true}
}
func (_faddd *Paragraph) wrapText() error {
	if !_faddd._acegd || int(_faddd._dgee) <= 0 {
		_faddd._egge = []string{_faddd._bfagd}
		return nil
	}
	_ecaf := NewTextChunk(_faddd._bfagd, TextStyle{Font: _faddd._cgbf, FontSize: _faddd._ebbae})
	_ffbe, _fdeb := _ecaf.Wrap(_faddd._dgee)
	if _fdeb != nil {
		return _fdeb
	}
	if _faddd._gefg > 0 && len(_ffbe) > _faddd._gefg {
		_ffbe = _ffbe[:_faddd._gefg]
	}
	_faddd._egge = _ffbe
	return nil
}

// SetStyle sets the style for all the line components: number, title,
// separator, page.
func (_gdecg *TOCLine) SetStyle(style TextStyle) {
	_gdecg.Number.Style = style
	_gdecg.Title.Style = style
	_gdecg.Separator.Style = style
	_gdecg.Page.Style = style
}

// SetColorBottom sets border color for bottom.
func (_eefg *border) SetColorBottom(col Color) { _eefg._aad = col }

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_ebffe *StyledParagraph) SetTextAlignment(align TextAlignment) { _ebffe._abgbf = align }

// SetPos sets the Block's positioning to absolute mode with the specified coordinates.
func (_ga *Block) SetPos(x, y float64) { _ga._ff = PositionAbsolute; _ga._ce = x; _ga._fe = y }

// NewDivision returns a new Division container component.
func (_gefb *Creator) NewDivision() *Division { return _eded() }

// SetBorderColor sets the border color for the path.
func (_afef *FilledCurve) SetBorderColor(color Color) { _afef._dddbg = color }

// NewEllipse creates a new ellipse centered at (xc,yc) with a width and height specified.
func (_aca *Creator) NewEllipse(xc, yc, width, height float64) *Ellipse {
	return _dcb(xc, yc, width, height)
}

// GetMargins returns the left, right, top, bottom Margins.
func (_facb *Table) GetMargins() (float64, float64, float64, float64) {
	return _facb._gcaee.Left, _facb._gcaee.Right, _facb._gcaee.Top, _facb._gcaee.Bottom
}

// SetIncludeInTOC sets a flag to indicate whether or not to include in tOC.
func (_baeg *Chapter) SetIncludeInTOC(includeInTOC bool) { _baeg._cdd = includeInTOC }

// SetBorderWidth sets the border width.
func (_egeb *PolyBezierCurve) SetBorderWidth(borderWidth float64) {
	_egeb._agfc.BorderWidth = borderWidth
}

// MultiCell makes a new cell with the specified row span and col span
// and inserts it into the table at the current position.
func (_ecfcb *Table) MultiCell(rowspan, colspan int) *TableCell {
	_ecfcb._daab++
	_bfeb := (_ecfcb.moveToNextAvailableCell()-1)%(_ecfcb._fabb) + 1
	_edcga := (_ecfcb._daab-1)/_ecfcb._fabb + 1
	for _edcga > _ecfcb._gbcf {
		_ecfcb._gbcf++
		_ecfcb._abeg = append(_ecfcb._abeg, _ecfcb._badaf)
	}
	_eeaeg := &TableCell{}
	_eeaeg._cbfc = _edcga
	_eeaeg._gbdde = _bfeb
	_eeaeg._fffcb = 5
	_eeaeg._dbcf = CellBorderStyleNone
	_eeaeg._egba = _bb.LineStyleSolid
	_eeaeg._dcee = CellHorizontalAlignmentLeft
	_eeaeg._dfcfe = CellVerticalAlignmentTop
	_eeaeg._geaef = 0
	_eeaeg._eecfg = 0
	_eeaeg._edgb = 0
	_eeaeg._eaeg = 0
	_afedg := ColorBlack
	_eeaeg._faaea = _afedg
	_eeaeg._becd = _afedg
	_eeaeg._cedab = _afedg
	_eeaeg._bdcf = _afedg
	if rowspan < 1 {
		_da.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061t\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0072\u006f\u0077s\u0070\u0061n\u0020\u0074o\u00201\u002e", rowspan)
		rowspan = 1
	}
	_gbbf := _ecfcb._gbcf - (_eeaeg._cbfc - 1)
	if rowspan > _gbbf {
		_da.Log.Debug("\u0054\u0061b\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070\u0061\u006e\u0020\u0028\u0025d\u0029\u0020\u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072o\u0077\u0073 \u0028\u0025\u0064\u0029.\u0020\u0041\u0064\u0064\u0069n\u0067\u0020\u0072\u006f\u0077\u0073\u002e", rowspan, _gbbf)
		_ecfcb._gbcf += rowspan - 1
		for _gcedgg := 0; _gcedgg <= rowspan-_gbbf; _gcedgg++ {
			_ecfcb._abeg = append(_ecfcb._abeg, _ecfcb._badaf)
		}
	}
	for _ffee := 0; _ffee < colspan && _bfeb+_ffee-1 < len(_ecfcb._aegce); _ffee++ {
		_ecfcb._aegce[_bfeb+_ffee-1] = rowspan - 1
	}
	_eeaeg._cdcd = rowspan
	if colspan < 1 {
		_da.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0063\u006f\u006c\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061n\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006cs\u0070\u0061n\u0020\u0074o\u00201\u002e", colspan)
		colspan = 1
	}
	_gfaf := _ecfcb._fabb - (_eeaeg._gbdde - 1)
	if colspan > _gfaf {
		_da.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0065\u006c\u006c\u0020\u0063o\u006c\u0073\u0070\u0061\u006e\u0020\u0028\u0025\u0064\u0029\u0020\u0065\u0078\u0063\u0065\u0065\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072\u006f\u0077\u0020\u0063\u006f\u006c\u0073\u0020\u0028\u0025d\u0029\u002e\u0020\u0041\u0064\u006a\u0075\u0073\u0074\u0069\u006e\u0067 \u0063\u006f\u006c\u0073\u0070\u0061n\u002e", colspan, _gfaf)
		colspan = _gfaf
	}
	_eeaeg._bdbg = colspan
	_ecfcb._daab += colspan - 1
	_ecfcb._dcab = append(_ecfcb._dcab, _eeaeg)
	_eeaeg._gfdee = _ecfcb
	return _eeaeg
}

// SetFillOpacity sets the fill opacity.
func (_egdad *PolyBezierCurve) SetFillOpacity(opacity float64) { _egdad._afeb = opacity }

// SetStyleLeft sets border style for left side.
func (_dgc *border) SetStyleLeft(style CellBorderStyle) { _dgc._gda = style }

// SetMargins sets the margins of the chart component.
func (_becg *Chart) SetMargins(left, right, top, bottom float64) {
	_becg._gedd.Left = left
	_becg._gedd.Right = right
	_becg._gedd.Top = top
	_becg._gedd.Bottom = bottom
}
func _dggd(_bfge *_g.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 238), Font: _bfge, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}
func (_ecc *Block) duplicate() *Block {
	_dbe := &Block{}
	*_dbe = *_ecc
	_aea := _fc.ContentStreamOperations{}
	_aea = append(_aea, *_ecc._cb...)
	_dbe._cb = &_aea
	return _dbe
}

// SetHorizontalAlignment sets the horizontal alignment of the image.
func (_eadc *Image) SetHorizontalAlignment(alignment HorizontalAlignment) { _eadc._eecd = alignment }

// SetMargins sets the Table's left, right, top, bottom margins.
func (_dfcf *Table) SetMargins(left, right, top, bottom float64) {
	_dfcf._gcaee.Left = left
	_dfcf._gcaee.Right = right
	_dfcf._gcaee.Top = top
	_dfcf._gcaee.Bottom = bottom
}

// Angle returns the block rotation angle in degrees.
func (_cff *Block) Angle() float64 { return _cff._fb }

// SetMaxLines sets the maximum number of lines before the paragraph
// text is truncated.
func (_agcf *Paragraph) SetMaxLines(maxLines int) { _agcf._gefg = maxLines; _agcf.wrapText() }

// SetLinePageStyle sets the style for the page part of all new lines
// of the table of contents.
func (_gbcc *TOC) SetLinePageStyle(style TextStyle) { _gbcc._feaff = style }

// ColorRGBFromHex converts color hex code to rgb color for using with creator.
// NOTE: If there is a problem interpreting the string, then will use black color and log a debug message.
// Example hex code: #ffffff -> (1,1,1) white.
func ColorRGBFromHex(hexStr string) Color {
	_fagc := rgbColor{}
	if (len(hexStr) != 4 && len(hexStr) != 7) || hexStr[0] != '#' {
		_da.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
		return _fagc
	}
	var _gcgb, _gdb, _gbgc int
	if len(hexStr) == 4 {
		var _afadb, _gcd, _ccgg int
		_dec, _caab := _ad.Sscanf(hexStr, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_afadb, &_gcd, &_ccgg)
		if _caab != nil {
			_da.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", hexStr, _caab)
			return _fagc
		}
		if _dec != 3 {
			_da.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _fagc
		}
		_gcgb = _afadb*16 + _afadb
		_gdb = _gcd*16 + _gcd
		_gbgc = _ccgg*16 + _ccgg
	} else {
		_cdac, _ffdc := _ad.Sscanf(hexStr, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_gcgb, &_gdb, &_gbgc)
		if _ffdc != nil {
			_da.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _fagc
		}
		if _cdac != 3 {
			_da.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", hexStr, _cdac)
			return _fagc
		}
	}
	_deg := float64(_gcgb) / 255.0
	_dfe := float64(_gdb) / 255.0
	_fcda := float64(_gbgc) / 255.0
	_fagc._aacc = _deg
	_fagc._beeb = _dfe
	_fagc._fgea = _fcda
	return _fagc
}
func _fade(_eggca int64, _aedb, _abec, _bdebb float64) *_g.PdfAnnotation {
	_eace := _g.NewPdfAnnotationLink()
	_ddgd := _g.NewBorderStyle()
	_ddgd.SetBorderWidth(0)
	_eace.BS = _ddgd.ToPdfObject()
	if _eggca < 0 {
		_eggca = 0
	}
	_eace.Dest = _ec.MakeArray(_ec.MakeInteger(_eggca), _ec.MakeName("\u0058\u0059\u005a"), _ec.MakeFloat(_aedb), _ec.MakeFloat(_abec), _ec.MakeFloat(_bdebb))
	return _eace.PdfAnnotation
}
func (_dac *Creator) initContext() {
	_dac._defb.X = _dac._acda.Left
	_dac._defb.Y = _dac._acda.Top
	_dac._defb.Width = _dac._cgb - _dac._acda.Right - _dac._acda.Left
	_dac._defb.Height = _dac._cggd - _dac._acda.Bottom - _dac._acda.Top
	_dac._defb.PageHeight = _dac._cggd
	_dac._defb.PageWidth = _dac._cgb
	_dac._defb.Margins = _dac._acda
	_dac._defb._bfaf = _dac.UnsupportedCharacterReplacement
}

// SetColorRight sets border color for right.
func (_bccf *border) SetColorRight(col Color) { _bccf._edf = col }

// GetHeading returns the chapter heading paragraph. Used to give access to address style: font, sizing etc.
func (_geg *Chapter) GetHeading() *Paragraph { return _geg._gacfc }

// Draw draws the drawable d on the block.
// Note that the drawable must not wrap, i.e. only return one block. Otherwise an error is returned.
func (_cba *Block) Draw(d Drawable) error {
	_fafb := DrawContext{}
	_fafb.Width = _cba._deb
	_fafb.Height = _cba._df
	_fafb.PageWidth = _cba._deb
	_fafb.PageHeight = _cba._df
	_fafb.X = 0
	_fafb.Y = 0
	_agc, _, _cbdb := d.GeneratePageBlocks(_fafb)
	if _cbdb != nil {
		return _cbdb
	}
	if len(_agc) != 1 {
		return _f.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0062\u006c\u006f\u0063k\u0073")
	}
	for _, _ffg := range _agc {
		if _bfed := _cba.mergeBlocks(_ffg); _bfed != nil {
			return _bfed
		}
	}
	return nil
}
func (_aeed *InvoiceAddress) fmtLine(_eageb, _bce string, _dacb bool) string {
	if _dacb {
		_bce = ""
	}
	return _ad.Sprintf("\u0025\u0073\u0025s\u000a", _bce, _eageb)
}

// ColorRGBFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//   green := ColorRGBFromArithmetic(0.0, 1.0, 0.0)
func ColorRGBFromArithmetic(r, g, b float64) Color {
	return rgbColor{_aacc: _dg.Max(_dg.Min(r, 1.0), 0.0), _beeb: _dg.Max(_dg.Min(g, 1.0), 0.0), _fgea: _dg.Max(_dg.Min(b, 1.0), 0.0)}
}

// GetMargins returns the margins of the chart (left, right, top, bottom).
func (_fegd *Chart) GetMargins() (float64, float64, float64, float64) {
	return _fegd._gedd.Left, _fegd._gedd.Right, _fegd._gedd.Top, _fegd._gedd.Bottom
}

// SetColumnWidths sets the fractional column widths.
// Each width should be in the range 0-1 and is a fraction of the table width.
// The number of width inputs must match number of columns, otherwise an error is returned.
func (_bcgbe *Table) SetColumnWidths(widths ...float64) error {
	if len(widths) != _bcgbe._fabb {
		_da.Log.Debug("M\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0077\u0069\u0064\u0074\u0068\u0073\u0020\u0061nd\u0020\u0063\u006fl\u0075m\u006e\u0073")
		return _f.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_bcgbe._dcdg = widths
	return nil
}

// SetForms adds an Acroform to a PDF file.  Sets the specified form for writing.
func (_aagd *Creator) SetForms(form *_g.PdfAcroForm) error { _aagd._dfce = form; return nil }

// Chart represents a chart drawable.
// It is used to render unichart chart components using a creator instance.
type Chart struct {
	_gfgd  _b.ChartRenderable
	_cbfa  Positioning
	_ebfbg float64
	_afe   float64
	_gedd  Margins
}

// SetHeight sets the Image's document height to specified h.
func (_eeeb *Image) SetHeight(h float64) { _eeeb._gcab = h }

// NewCell makes a new cell and inserts it into the table at the current position.
func (_cafdc *Table) NewCell() *TableCell { return _cafdc.MultiCell(1, 1) }

// SetDate sets the date of the invoice.
func (_gfggb *Invoice) SetDate(date string) (*InvoiceCell, *InvoiceCell) {
	_gfggb._dgdd[1].Value = date
	return _gfggb._dgdd[0], _gfggb._dgdd[1]
}

// Subtotal returns the invoice subtotal description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_bdgcg *Invoice) Subtotal() (*InvoiceCell, *InvoiceCell) {
	return _bdgcg._afga[0], _bdgcg._afga[1]
}

// SetColor sets the line color.
func (_dcaee *Curve) SetColor(col Color) { _dcaee._eeb = col }

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_fbcdc *TOCLine) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_egcg := ctx
	_dcag, ctx, _fgfdca := _fbcdc._daedd.GeneratePageBlocks(ctx)
	if _fgfdca != nil {
		return _dcag, ctx, _fgfdca
	}
	if _fbcdc._cada.IsRelative() {
		ctx.X = _egcg.X
	}
	if _fbcdc._cada.IsAbsolute() {
		return _dcag, _egcg, nil
	}
	return _dcag, ctx, nil
}

// Text sets the text content of the Paragraph.
func (_dcca *Paragraph) Text() string { return _dcca._bfagd }

// SetLevel sets the indentation level of the TOC line.
func (_begb *TOCLine) SetLevel(level uint) {
	_begb._cefb = level
	_begb._daedd._cdbab.Left = _begb._egde + float64(_begb._cefb-1)*_begb._dedcg
}
func (_fdbb *Image) makeXObject() error {
	_ddag := _fdbb._fdgf
	if _ddag == nil {
		_ddag = _ec.NewFlateEncoder()
	}
	_bdfd, _ccc := _g.NewXObjectImageFromImage(_fdbb._ggdgd, nil, _ddag)
	if _ccc != nil {
		_da.Log.Error("\u0046\u0061\u0069le\u0064\u0020\u0074\u006f\u0020\u0063\u0072\u0065\u0061t\u0065 \u0078o\u0062j\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _ccc)
		return _ccc
	}
	_fdbb._edad = _bdfd
	return nil
}

// SetWidth sets line width.
func (_fegb *Curve) SetWidth(width float64) { _fegb._faca = width }

// Flip flips the active page on the specified axes.
// If `flipH` is true, the page is flipped horizontally. Similarly, if `flipV`
// is true, the page is flipped vertically. If both are true, the page is
// flipped both horizontally and vertically.
// NOTE: the flip transformations are applied when the creator is finalized,
// which is at write time in most cases.
func (_adda *Creator) Flip(flipH, flipV bool) error {
	_febg := _adda.getActivePage()
	if _febg == nil {
		return _f.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_dgf, _cgfg := _adda._gff[_febg]
	if !_cgfg {
		_dgf = &pageTransformations{}
		_adda._gff[_febg] = _dgf
	}
	_dgf._cead = flipH
	_dgf._aecf = flipV
	return nil
}

// AddLine adds a new line with the provided style to the table of contents.
func (_bgace *TOC) AddLine(line *TOCLine) *TOCLine {
	if line == nil {
		return nil
	}
	_bgace._dbcgf = append(_bgace._dbcgf, line)
	return line
}

// Height returns the current page height.
func (_bffb *Creator) Height() float64 { return _bffb._cggd }

// GeneratePageBlocks draws the rectangle on a new block representing the page. Implements the Drawable interface.
func (_ebcd *Rectangle) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cddga := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_gfde := _bb.Rectangle{Opacity: 1.0, X: _ebcd._eebca, Y: ctx.PageHeight - _ebcd._eaaa - _ebcd._eeee, Height: _ebcd._eeee, Width: _ebcd._ecgg}
	if _ebcd._edga != nil {
		_gfde.FillEnabled = true
		_gfde.FillColor = _afag(_ebcd._edga)
	}
	if _ebcd._fdegf != nil && _ebcd._abfa > 0 {
		_gfde.BorderEnabled = true
		_gfde.BorderColor = _afag(_ebcd._fdegf)
		_gfde.BorderWidth = _ebcd._abfa
	}
	_cgfdg, _bfba := _cddga.setOpacity(_ebcd._fcae, _ebcd._ccgf)
	if _bfba != nil {
		return nil, ctx, _bfba
	}
	_ebfa, _, _bfba := _gfde.Draw(_cgfdg)
	if _bfba != nil {
		return nil, ctx, _bfba
	}
	if _bfba = _cddga.addContentsByString(string(_ebfa)); _bfba != nil {
		return nil, ctx, _bfba
	}
	return []*Block{_cddga}, ctx, nil
}

// Total returns the invoice total description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_cafe *Invoice) Total() (*InvoiceCell, *InvoiceCell) { return _cafe._cdgb[0], _cafe._cdgb[1] }

// SetBorderWidth sets the border width.
func (_cedf *Polygon) SetBorderWidth(borderWidth float64) { _cedf._fdfa.BorderWidth = borderWidth }

// EnablePageWrap controls whether the division is wrapped across pages.
// If disabled, the division is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the division is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
// Currently, page wrapping can only be disabled for vertical divisions.
func (_defe *Division) EnablePageWrap(enable bool) { _defe._babf = enable }

// NewCellProps returns the default properties of an invoice cell.
func (_bafdb *Invoice) NewCellProps() InvoiceCellProps {
	_aeaa := ColorRGBFrom8bit(255, 255, 255)
	return InvoiceCellProps{TextStyle: _bafdb._fafbeg, Alignment: CellHorizontalAlignmentLeft, BackgroundColor: _aeaa, BorderColor: _aeaa, BorderWidth: 1, BorderSides: []CellBorderSide{CellBorderSideAll}}
}

// MoveY moves the drawing context to absolute position y.
func (_dcg *Creator) MoveY(y float64) { _dcg._defb.Y = y }

// SetBorderWidth sets the border width.
func (_gfcdf *Rectangle) SetBorderWidth(bw float64) { _gfcdf._abfa = bw }

// AddressStyle returns the style properties used to render the content of
// the invoice address sections.
func (_adac *Invoice) AddressStyle() TextStyle { return _adac._ceadb }
func (_adg *Invoice) newCell(_cadg string, _ggde InvoiceCellProps) *InvoiceCell {
	return &InvoiceCell{_ggde, _cadg}
}

// Width returns the cell's width based on the input draw context.
func (_ddfbb *TableCell) Width(ctx DrawContext) float64 {
	_bebg := float64(0.0)
	for _fbagc := 0; _fbagc < _ddfbb._bdbg; _fbagc++ {
		_bebg += _ddfbb._gfdee._dcdg[_ddfbb._gbdde+_fbagc-1]
	}
	_ccad := ctx.Width * _bebg
	return _ccad
}

// Heading returns the heading component of the table of contents.
func (_afae *TOC) Heading() *StyledParagraph { return _afae._ebgb }

// SetFillColor sets the fill color for the path.
func (_bccd *FilledCurve) SetFillColor(color Color) { _bccd._debe = color }

// SkipCells skips over a specified number of cells in the table.
func (_gcbff *Table) SkipCells(num int) {
	if num < 0 {
		_da.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	_gcbff._daab += num
}

// AddExternalLink adds a new external link to the paragraph.
// The text parameter represents the text that is displayed and the url
// parameter sets the destionation of the link.
func (_cddd *StyledParagraph) AddExternalLink(text, url string) *TextChunk {
	_fcfb := NewTextChunk(text, _cddd._aacda)
	_fcfb._adbcg = _cgccc(url)
	return _cddd.appendChunk(_fcfb)
}

// SetLevelOffset sets the amount of space an indentation level occupies.
func (_edgfc *TOCLine) SetLevelOffset(levelOffset float64) {
	_edgfc._dedcg = levelOffset
	_edgfc._daedd._cdbab.Left = _edgfc._egde + float64(_edgfc._cefb-1)*_edgfc._dedcg
}

// Marker returns the marker used for the list items.
// The marker instance can be used the change the text and the style
// of newly added list items.
func (_fgfg *List) Marker() *TextChunk { return &_fgfg._ddcde }

// SellerAddress returns the seller address used in the invoice template.
func (_gacbe *Invoice) SellerAddress() *InvoiceAddress { return _gacbe._aedd }

const (
	CellVerticalAlignmentTop CellVerticalAlignment = iota
	CellVerticalAlignmentMiddle
	CellVerticalAlignmentBottom
)

// TotalLines returns all the rows in the invoice totals table as
// description-value cell pairs.
func (_fdfc *Invoice) TotalLines() [][2]*InvoiceCell {
	_bfedf := [][2]*InvoiceCell{_fdfc._afga}
	_bfedf = append(_bfedf, _fdfc._gegc...)
	return append(_bfedf, _fdfc._cdgb)
}

// SetLogo sets the logo of the invoice.
func (_bdgcd *Invoice) SetLogo(logo *Image) { _bdgcd._abfb = logo }

// Rectangle defines a rectangle with upper left corner at (x,y) and a specified width and height.  The rectangle
// can have a colored fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Rectangle struct {
	_eebca float64
	_eaaa  float64
	_ecgg  float64
	_eeee  float64
	_edga  Color
	_fcae  float64
	_fdegf Color
	_abfa  float64
	_ccgf  float64
}

// SetTextVerticalAlignment sets the vertical alignment of the text within the
// bounds of the styled paragraph.
func (_eggf *StyledParagraph) SetTextVerticalAlignment(align TextVerticalAlignment) {
	_eggf._fgce = align
}

// SetLineOpacity sets the line opacity.
func (_egcb *Polyline) SetLineOpacity(opacity float64) { _egcb._beccb = opacity }

// NewTable create a new Table with a specified number of columns.
func (_bedgd *Creator) NewTable(cols int) *Table { return _dbfc(cols) }

// List represents a list of items.
// The representation of a list item is as follows:
//       [marker] [content]
// e.g.:        • This is the content of the item.
// The supported components to add content to list items are:
// - Paragraph
// - StyledParagraph
// - List
type List struct {
	_cbba  []*listItem
	_cgab  Margins
	_ddcde TextChunk
	_fefc  float64
	_dcffe bool
	_cffa  Positioning
	_egbg  TextStyle
}

// NewPage adds a new Page to the Creator and sets as the active Page.
func (_egd *Creator) NewPage() *_g.PdfPage {
	_eee := _egd.newPage()
	_egd._bdb = append(_egd._bdb, _eee)
	_egd._defb.Page++
	return _eee
}

// EnableFontSubsetting enables font subsetting for `font` when the creator output is written to file.
// Embeds only the subset of the runes/glyphs that are actually used to display the file.
// Subsetting can reduce the size of fonts significantly.
func (_edcg *Creator) EnableFontSubsetting(font *_g.PdfFont) { _edcg._dbc = append(_edcg._dbc, font) }

// GeneratePageBlocks draws the filled curve on page blocks.
func (_eced *FilledCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bdcc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_degg, _, _bfg := _eced.draw("")
	if _bfg != nil {
		return nil, ctx, _bfg
	}
	_bfg = _bdcc.addContentsByString(string(_degg))
	if _bfg != nil {
		return nil, ctx, _bfg
	}
	return []*Block{_bdcc}, ctx, nil
}
